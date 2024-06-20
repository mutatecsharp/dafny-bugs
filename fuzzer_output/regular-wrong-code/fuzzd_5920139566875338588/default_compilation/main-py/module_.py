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
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjcdhou")))])).Elements:
                d_1_v0_: int = compr_0_
                if (d_1_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjcdhou")))])):
                    coll0_[(d_1_v0_) - (len(_dafny.Map({_dafny.CodePoint('h'): True})))] = not(True)
            return _dafny.Map(coll0_)
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "clrpyqdpw"))): False})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-165, len(_dafny.Set({False, False})), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): False})), 590])) for d_0_i0_ in range(default__.abs(898))])): True}))) | (iife0_()
        )

    @staticmethod
    def fm1(p0, globalState):
        if False:
            return (0) - ((len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(D16_DC45((0) - (-80), _dafny.CodePoint('o'), 808, True)).cf79, _dafny.CodePoint('k'), _dafny.CodePoint('o'), _dafny.CodePoint('j')])): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2_i0_ in range(default__.abs(865))]))}))) * (len(_dafny.SeqWithoutIsStrInference([False, not(True)]))))
        elif True:
            return (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_3_i1_ in range(default__.abs(781))]))) - (56)

    @staticmethod
    def fm2(globalState):
        return not(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "grkmmpyo"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "holvusflx")))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gl"))))

    @staticmethod
    def fm12(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvso"))): len(_dafny.SeqWithoutIsStrInference([True]))})) for d_4_i0_ in range(default__.abs(-895))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_5_i1_ in range(default__.abs(251))])), -624]))) + (_dafny.SeqWithoutIsStrInference([789, 253]))

    @staticmethod
    def fm13(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True, not(True)])) + (_dafny.SeqWithoutIsStrInference([True]))) + ((_dafny.SeqWithoutIsStrInference([True, False])) + (_dafny.SeqWithoutIsStrInference([not(True)])))

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        return (_dafny.Map({False: 556})) | ((_dafny.Map({False: len(_dafny.Set({False}))})))

    @staticmethod
    def fm18(p0, p1, globalState):
        return D0_DC0(not(not(True)))

    @staticmethod
    def fm19(globalState):
        return ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False]))) + ((_dafny.SeqWithoutIsStrInference([True, not(True)])) + (_dafny.SeqWithoutIsStrInference([False, True, True])))

    @staticmethod
    def fm20(p0, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: D7
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([D7_DC16(True, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_6_i0_ in range(default__.abs(409))]))), not(not(False))), D7_DC16(True, -288, False)])).Elements:
                d_7_v0_: D7 = compr_1_
                if (d_7_v0_) in (_dafny.SeqWithoutIsStrInference([D7_DC16(True, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_6_i0_ in range(default__.abs(409))]))), not(not(False))), D7_DC16(True, -288, False)])):
                    coll1_[d_7_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsvfptg"))
            return _dafny.Map(coll1_)
        source0_ = D13_DC36(iife1_()
)
        if source0_.is_DC37:
            d_8___mcc_h0_ = source0_.cf65
            d_9___mcc_h1_ = source0_.cf66
            d_10_cf66_ = d_9___mcc_h1_
            d_11_cf65_ = d_8___mcc_h0_
            return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_10_cf66_, d_10_cf66_]), _dafny.SeqWithoutIsStrInference([d_10_cf66_, d_11_cf65_, d_10_cf66_]), _dafny.SeqWithoutIsStrInference([d_11_cf65_, d_11_cf65_, d_11_cf65_, d_10_cf66_]), _dafny.SeqWithoutIsStrInference([d_11_cf65_])])
        elif True:
            d_12___mcc_h2_ = source0_.cf64
            d_13_cf64_ = d_12___mcc_h2_
            return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_14_i1_ in range(default__.abs(-526))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(True), not(True)])])))

    @staticmethod
    def fm21(globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqwtxsk"))

    @staticmethod
    def fm22(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqfqwylgv"))) for d_15_i0_ in range(default__.abs(-625))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: False})), (_dafny.MultiSet([-177, len(_dafny.SeqWithoutIsStrInference([-425]))])).cardinality]), _dafny.SeqWithoutIsStrInference([-684])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r')]))]) for d_16_i1_ in range(default__.abs(-76))]))

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        if (False) and (True):
            return _dafny.SeqWithoutIsStrInference([-383, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfg"))): _dafny.CodePoint('a')}))])
        elif True:
            return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oi"))), (0) - ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Set({True, not(False)})]))))), -560, 92, 248])

    @staticmethod
    def fm26(p0, p1, p2, p3, globalState):
        return (_dafny.Map({len(_dafny.Map({True: len(_dafny.Map({False: True}))})): _dafny.Set({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqcb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qswfrnph"))])): -964}))})})) | (_dafny.Map({515: _dafny.Set({-634})}))

    @staticmethod
    def fm27(p0, p1, p2, p3, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(361, 313):
                d_17_v0_: int = compr_2_
                if ((361) <= (d_17_v0_)) and ((d_17_v0_) < (313)):
                    coll2_[(d_17_v0_) * (470)] = _dafny.Set({(0) - (-602), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbheb")))})
            return _dafny.Map(coll2_)
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(699, 68):
                d_18_v1_: int = compr_3_
                if ((699) <= (d_18_v1_)) and ((d_18_v1_) < (68)):
                    coll3_ = coll3_.union(_dafny.Set([(d_18_v1_) + (len(_dafny.Set({670, 565})))]))
            return _dafny.Set(coll3_)
        return ((_dafny.Map({492: _dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))})})) | (iife2_()
        )) | ((_dafny.Map({(_dafny.MultiSet([False])).cardinality: _dafny.Set({len(_dafny.Map({True: 772}))})}) if True else _dafny.Map({282: iife3_()
        })))

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        if (False if False else not(not(False))):
            def iife4_():
                coll4_ = _dafny.Set()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(768, 29):
                    d_19_v0_: int = compr_4_
                    if ((768) <= (d_19_v0_)) and ((d_19_v0_) < (29)):
                        coll4_ = coll4_.union(_dafny.Set([default__.safeDivisionInt(d_19_v0_, 801)]))
                return _dafny.Set(coll4_)
            return iife4_()
            
        elif True:
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(678, -66):
                    d_20_v1_: int = compr_5_
                    if ((678) <= (d_20_v1_)) and ((d_20_v1_) < (-66)):
                        coll5_[(d_20_v1_) - (325)] = True
                return _dafny.Map(coll5_)
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(75, 695):
                    d_22_v2_: int = compr_6_
                    if ((75) <= (d_22_v2_)) and ((d_22_v2_) < (695)):
                        coll6_[(d_22_v2_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_23_i1_ in range(default__.abs(654))])))] = len(_dafny.Map({False: True}))
                return _dafny.Map(coll6_)
            return (_dafny.Set({(0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([115]))).cardinality)})).intersection(_dafny.Set({785, len(_dafny.Set({378, len(iife5_()
            ), len(_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nb"))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kokftkjn"))), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_21_i0_ in range(default__.abs(784))])): _dafny.Map({D12_DC34(iife6_()
): True})}))})), 709, 124, len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmilbnjok")))}))}))

    @staticmethod
    def fm29(p0, p1, p2, p3, globalState):
        def iife7_():
            coll7_ = _dafny.Set()
            compr_7_: int
            for compr_7_ in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mphpqtpee"))), 588})).Elements:
                d_24_v0_: int = compr_7_
                if (d_24_v0_) in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mphpqtpee"))), 588})):
                    coll7_ = coll7_.union(_dafny.Set([(d_24_v0_) * ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymib")))))]))
            return _dafny.Set(coll7_)
        source1_ = D9_DC24(_dafny.Map({794: _dafny.Map({304: iife7_()
})}), _dafny.SeqWithoutIsStrInference([True]))
        if source1_.is_DC23:
            d_25___mcc_h0_ = source1_.cf37
            d_26___mcc_h1_ = source1_.cf38
            d_27___mcc_h2_ = source1_.cf39
            d_28___mcc_h3_ = source1_.cf40
            d_29___mcc_h4_ = source1_.cf41
            d_30_cf41_ = d_29___mcc_h4_
            d_31_cf40_ = d_28___mcc_h3_
            d_32_cf39_ = d_27___mcc_h2_
            d_33_cf38_ = d_26___mcc_h1_
            d_34_cf37_ = d_25___mcc_h0_
            return d_30_cf41_
        elif source1_.is_DC24:
            d_35___mcc_h5_ = source1_.cf42
            d_36___mcc_h6_ = source1_.cf43
            d_37_cf43_ = d_36___mcc_h6_
            d_38_cf42_ = d_35___mcc_h5_
            return _dafny.CodePoint('t')
        elif source1_.is_DC22:
            d_39___mcc_h7_ = source1_.cf36
            d_40_cf36_ = d_39___mcc_h7_
            if True:
                return _dafny.CodePoint('w')
            elif True:
                return _dafny.CodePoint('d')
        elif True:
            d_41___mcc_h8_ = source1_.cf44
            d_42_cf44_ = d_41___mcc_h8_
            return _dafny.CodePoint('a')

    @staticmethod
    def fm30(p0, p1, globalState):
        return ((_dafny.MultiSet([_dafny.MultiSet([True, False])])) - (_dafny.MultiSet([_dafny.MultiSet([True, False, False])]))).intersection(_dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True]))]))

    @staticmethod
    def fm31(p0, p1, globalState):
        return _dafny.Set({not(True), (not(True) if True else True), not (False) or (True), (494) == ((0) - (-851)), not (False) or (not(False))})

    @staticmethod
    def fm32(p0, p1, p2, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(258, 368):
                d_43_v0_: int = compr_8_
                if ((258) <= (d_43_v0_)) and ((d_43_v0_) < (368)):
                    coll8_[(d_43_v0_) - (81)] = False
            return _dafny.Map(coll8_)
        return _dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: _dafny.Map({43: False})})), (-552) * (len(iife8_()
        )), default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kchk"))), 731)])

    @staticmethod
    def fm33(p0, p1, p2, globalState):
        return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yoy")))])): 229})) | ((_dafny.Map({223: 388})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True, False])): 642})))

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({True: not(not(not(not(False))))}))) | ((_dafny.Map({True: True})) | (_dafny.Map({False: False})))

    @staticmethod
    def fm35(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([D0_DC0(not(False))])) + (_dafny.SeqWithoutIsStrInference([D0_DC0(False), D0_DC0(False), D0_DC0(True)]))

    @staticmethod
    def fm36(p0, p1, globalState):
        return (_dafny.Map({not(False): len(_dafny.Map({True: -760}))})) | ((_dafny.Map({True: len(_dafny.Map({True: False}))})) | (_dafny.Map({not(not(False)): -685})))

    @staticmethod
    def fm37(p0, p1, globalState):
        return D10_DC28((-150 if True else 41))

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        def iife9_():
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: D7
                for compr_10_ in (_dafny.SeqWithoutIsStrInference([D7_DC16(False, -159, True), D7_DC16(False, -950, False), D7_DC16(not(False), 953, False), D7_DC16(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbwndisey"))), True), D7_DC16(False, 344, True)])).Elements:
                    d_45_v1_: D7 = compr_10_
                    if (d_45_v1_) in (_dafny.SeqWithoutIsStrInference([D7_DC16(False, -159, True), D7_DC16(False, -950, False), D7_DC16(not(False), 953, False), D7_DC16(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbwndisey"))), True), D7_DC16(False, 344, True)])):
                        coll10_[d_45_v1_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmge"))
                return _dafny.Map(coll10_)
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in (_dafny.SeqWithoutIsStrInference([757, 853, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krtxro"))))])).Elements:
                d_44_v0_: int = compr_9_
                if (d_44_v0_) in (_dafny.SeqWithoutIsStrInference([757, 853, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krtxro"))))])):
                    coll9_[default__.safeModuloInt(d_44_v0_, (_dafny.MultiSet([(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([D13_DC36(_dafny.Map({D7_DC16(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsiiwqtrg"))), False): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))})), D13_DC36(iife10_()
)]))))])).cardinality)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubeopiw"))
            return _dafny.Map(coll9_)
        return (_dafny.Map({826: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fghwqouji"))})) | (iife9_()
        )

    @staticmethod
    def fm39(p0, globalState):
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_46_i0_ in range(default__.abs(-847))]))]))])).Elements:
                d_47_v0_: int = compr_11_
                if (d_47_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_46_i0_ in range(default__.abs(-847))]))]))])):
                    coll11_[(d_47_v0_) * (954)] = False
            return _dafny.Map(coll11_)
        return (_dafny.Map({_dafny.Map({(_dafny.MultiSet([len(_dafny.Set({not(True), True, False}))])).cardinality: True}): False})) | ((_dafny.Map({_dafny.Map({144: False}): False})) | (_dafny.Map({iife11_()
        : False})))

    @staticmethod
    def fm40(p0, globalState):
        def iife12_():
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(862, 629):
                    d_49_v1_: int = compr_14_
                    if ((862) <= (d_49_v1_)) and ((d_49_v1_) < (629)):
                        coll14_[default__.safeModuloInt(d_49_v1_, -270)] = _dafny.Map({756: _dafny.Set({len(_dafny.Map({-241: -358}))})})
                return _dafny.Map(coll14_)
            coll12_ = _dafny.Map()
            def iife13_():
                coll13_ = _dafny.Map()
                compr_13_: int
                for compr_13_ in _dafny.IntegerRange(862, 629):
                    d_49_v1_: int = compr_13_
                    if ((862) <= (d_49_v1_)) and ((d_49_v1_) < (629)):
                        coll13_[default__.safeModuloInt(d_49_v1_, -270)] = _dafny.Map({756: _dafny.Set({len(_dafny.Map({-241: -358}))})})
                return _dafny.Map(coll13_)
            compr_12_: _dafny.Seq
            for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D9_DC24(iife13_()
, _dafny.SeqWithoutIsStrInference([True])) for d_50_i1_ in range(default__.abs(687))])])).Elements:
                d_51_v0_: _dafny.Seq = compr_12_
                if (d_51_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D9_DC24(iife14_()
, _dafny.SeqWithoutIsStrInference([True])) for d_50_i1_ in range(default__.abs(687))])])):
                    coll12_[d_51_v0_] = 520
            return _dafny.Map(coll12_)
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([D9_DC24(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, False])): _dafny.Map({334: _dafny.Set({345})})}), _dafny.SeqWithoutIsStrInference([True])), D9_DC24(_dafny.Map({80: _dafny.Map({len(_dafny.Set({False})): _dafny.Set({455, 528, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_48_i0_ in range(default__.abs(747))])), len(_dafny.Set({563}))})})}), _dafny.SeqWithoutIsStrInference([False, True]))]): 488})) | (iife12_()
        )

    @staticmethod
    def fm41(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (-901), -381, len(_dafny.SeqWithoutIsStrInference([True]))]))) - (_dafny.MultiSet([684, 946])), _dafny.MultiSet([len((D8_DC21(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_52_i0_ in range(default__.abs(-667))]), _dafny.Set({False, True}))).cf34), (D11_DC32(False, 392, True)).cf60, len(_dafny.Set({True, True})), -584, len(_dafny.SeqWithoutIsStrInference([not(True)]))]), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_53_i1_ in range(default__.abs(112))]))])) if False else _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkidualy")))]))])

    @staticmethod
    def fm42(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([866, (_dafny.MultiSet([572])).cardinality])), 556, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ls"))), 980, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))))]))).intersection(_dafny.MultiSet([-141]))

    @staticmethod
    def fm43(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([_dafny.MultiSet([len(_dafny.Map({False: not(not(True))}))])])) | (_dafny.MultiSet([_dafny.MultiSet([699]), _dafny.MultiSet([546])]))) | (_dafny.MultiSet([_dafny.MultiSet([-645, -185]), _dafny.MultiSet([-60, 334, 918]), _dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True, not(True)])): False})), 791])]))

    @staticmethod
    def fm44(globalState):
        return D13_DC37(not(False), False)

    @staticmethod
    def fm45(p0, globalState):
        def iife15_():
            coll15_ = _dafny.Map()
            compr_15_: int
            for compr_15_ in _dafny.IntegerRange(809, 705):
                d_54_v0_: int = compr_15_
                if ((809) <= (d_54_v0_)) and ((d_54_v0_) < (705)):
                    coll15_[default__.safeDivisionInt(d_54_v0_, len(_dafny.Map({False: False})))] = True
            return _dafny.Map(coll15_)
        return ((_dafny.SeqWithoutIsStrInference([iife15_()
        ])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({-192: not(False)}), _dafny.Map({745: False})]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "afelo"))): True}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))): True}), _dafny.Map({len(_dafny.Set({661})): False})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({769: True}), _dafny.Map({639: False})])))

    @staticmethod
    def fm46(p0, p1, p2, p3, globalState):
        source2_ = D8_DC19(_dafny.Set({True, False}))
        if source2_.is_DC20:
            d_55___mcc_h0_ = source2_.cf31
            d_56___mcc_h1_ = source2_.cf32
            d_57___mcc_h2_ = source2_.cf33
            d_58_cf33_ = d_57___mcc_h2_
            d_59_cf32_ = d_56___mcc_h1_
            d_60_cf31_ = d_55___mcc_h0_
            return _dafny.MultiSet([_dafny.CodePoint('k')])
        elif source2_.is_DC21:
            d_61___mcc_h3_ = source2_.cf34
            d_62___mcc_h4_ = source2_.cf35
            d_63_cf35_ = d_62___mcc_h4_
            d_64_cf34_ = d_61___mcc_h3_
            return (_dafny.MultiSet([_dafny.CodePoint('g')])) | (_dafny.MultiSet([(d_64_cf34_)[default__.safeIndex(856, len(d_64_cf34_))]]))
        elif True:
            d_65___mcc_h5_ = source2_.cf30
            d_66_cf30_ = d_65___mcc_h5_
            return _dafny.MultiSet([_dafny.CodePoint('d')])

    @staticmethod
    def fm47(p0, p1, p2, p3, globalState):
        return _dafny.Map({(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([811 for d_67_i0_ in range(default__.abs(935))])): False})) | (_dafny.Map({-766: not(False)})): 226})

    @staticmethod
    def fm48(p0, globalState):
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: _dafny.Map
            for compr_16_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Map({len(_dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality})): not(True)})): True}), _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yma")))): True})])).Elements:
                d_68_v0_: _dafny.Map = compr_16_
                if (d_68_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Map({len(_dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality})): not(True)})): True}), _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yma")))): True})])):
                    coll16_[d_68_v0_] = not(False)
            return _dafny.Map(coll16_)
        def iife17_():
            def iife19_():
                coll19_ = _dafny.Map()
                compr_19_: int
                for compr_19_ in _dafny.IntegerRange(728, -532):
                    d_69_v2_: int = compr_19_
                    if ((728) <= (d_69_v2_)) and ((d_69_v2_) < (-532)):
                        coll19_[(d_69_v2_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v')])))] = False
                return _dafny.Map(coll19_)
            coll17_ = _dafny.Map()
            def iife18_():
                coll18_ = _dafny.Map()
                compr_18_: int
                for compr_18_ in _dafny.IntegerRange(728, -532):
                    d_69_v2_: int = compr_18_
                    if ((728) <= (d_69_v2_)) and ((d_69_v2_) < (-532)):
                        coll18_[(d_69_v2_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v')])))] = False
                return _dafny.Map(coll18_)
            compr_17_: _dafny.Map
            for compr_17_ in (_dafny.MultiSet([iife18_()
            ])).Elements:
                d_70_v1_: _dafny.Map = compr_17_
                if (d_70_v1_) in (_dafny.MultiSet([iife19_()
                ])):
                    coll17_[d_70_v1_] = True
            return _dafny.Map(coll17_)
        return D3_DC7(len((iife16_()
) | (iife17_()
)))

    @staticmethod
    def fm49(p0, globalState):
        return (_dafny.Map({True: -586})) | (_dafny.Map({True: 492}))

    @staticmethod
    def fm50(p0, p1, globalState):
        if not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqld"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "futiieesj")))):
            return (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: True}), _dafny.Map({True: not(True)}), _dafny.Map({False: False}), _dafny.Map({False: True})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: True})]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.Map({True: True}) for d_71_i0_ in range(default__.abs(387))])

    @staticmethod
    def fm51(p0, globalState):
        return ((_dafny.Map({611: _dafny.MultiSet([False])})) | (_dafny.Map({len(_dafny.Map({_dafny.Map({not(True): _dafny.CodePoint('a')}): 187})): _dafny.MultiSet([False, True])}))) | (_dafny.Map({len(_dafny.Set({False})): _dafny.MultiSet([not(True)])}))

    @staticmethod
    def fm52(p0, p1, globalState):
        return D3_DC8(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfdjxwbg"))): len(_dafny.SeqWithoutIsStrInference([not(False), not(False)]))})): True}): _dafny.SeqWithoutIsStrInference([370])})), 181]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thgik")), (_dafny.MultiSet([True])).isdisjoint(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))), _dafny.CodePoint('u'))

    @staticmethod
    def fm53(p0, p1, globalState):
        return D15_DC43(not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umnb"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgs")))), (False) and (False), True)

    @staticmethod
    def fm54(p0, globalState):
        return ((_dafny.Map({369: _dafny.CodePoint('w')})) | (_dafny.Map({918: _dafny.CodePoint('l')}))) | (_dafny.Map({598: _dafny.CodePoint('s')}))

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        return _dafny.Set({D11_DC31(_dafny.SeqWithoutIsStrInference([True]), 975, (_dafny.MultiSet([D8_DC21(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmwvibhcx")), _dafny.Set({False}))])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dddghnql"))), len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([not(False)]))), 760])))})

    @staticmethod
    def fm56(p0, p1, globalState):
        if False:
            return D0_DC2(D0_DC0(False))
        elif True:
            return D0_DC2(D0_DC2(D0_DC0(not(True))))

    @staticmethod
    def fm57(p0, p1, globalState):
        return D14_DC40((64) == (len(_dafny.SeqWithoutIsStrInference([849]))))

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        d_72_v0_: _dafny.Array
        def lambda0_(d_73_i0_):
            return True

        init0_ = lambda0_
        nw0_ = _dafny.Array(None, 23)
        for i0_0_ in range(nw0_.length(0)):
            nw0_[i0_0_] = init0_(i0_0_)
        d_72_v0_ = nw0_
        d_74_v1_: _dafny.Set
        d_74_v1_ = _dafny.Set({d_72_v0_})
        d_75_v2_: D3
        d_75_v2_ = D3_DC6(p2)
        d_76_v3_: _dafny.Seq
        d_76_v3_ = _dafny.SeqWithoutIsStrInference([d_75_v2_])
        d_77_v4_: bool
        d_77_v4_ = True
        d_78_v5_: C0
        nw1_ = C0()
        nw1_.ctor__(d_76_v3_, d_77_v4_)
        d_78_v5_ = nw1_
        d_79_v6_: _dafny.Array
        nw2_ = _dafny.Array(None, 15)
        nw2_[int(0)] = d_78_v5_
        nw2_[int(1)] = d_78_v5_
        nw2_[int(2)] = d_78_v5_
        nw2_[int(3)] = d_78_v5_
        nw2_[int(4)] = d_78_v5_
        nw2_[int(5)] = d_78_v5_
        nw2_[int(6)] = d_78_v5_
        nw2_[int(7)] = d_78_v5_
        nw2_[int(8)] = d_78_v5_
        nw2_[int(9)] = d_78_v5_
        nw2_[int(10)] = d_78_v5_
        nw2_[int(11)] = d_78_v5_
        nw2_[int(12)] = d_78_v5_
        nw2_[int(13)] = d_78_v5_
        nw2_[int(14)] = d_78_v5_
        d_79_v6_ = nw2_
        d_80_v7_: _dafny.Seq
        d_80_v7_ = _dafny.SeqWithoutIsStrInference([d_79_v6_, d_79_v6_, d_79_v6_])
        d_81_v8_: D27
        d_81_v8_ = D27_DC61(d_80_v7_)
        pat_let_tv0_ = d_80_v7_
        def iife20_(_pat_let0_0):
            def iife21_(d_82_dt__update__tmp_h0_):
                def iife22_(_pat_let1_0):
                    def iife23_(d_83_dt__update_hcf101_h0_):
                        return D27_DC61(d_83_dt__update_hcf101_h0_)
                    return iife23_(_pat_let1_0)
                return iife22_(pat_let_tv0_)
            return iife21_(_pat_let0_0)
        source3_ = (iife20_(D27_DC61(d_80_v7_)) if (d_74_v1_).issubset(d_74_v1_) else d_81_v8_)
        if source3_.is_DC62:
            d_84___mcc_h0_ = source3_.cf102
            d_85___mcc_h1_ = source3_.cf103
            d_86___mcc_h2_ = source3_.cf104
            d_87___mcc_h3_ = source3_.cf105
            d_88_cf105_ = d_87___mcc_h3_
            d_89_cf104_ = d_86___mcc_h2_
            d_90_cf103_ = d_85___mcc_h1_
            d_91_cf102_ = d_84___mcc_h0_
            d_92_v9_: int
            d_92_v9_ = 398
            rhs0_ = d_92_v9_
            rhs1_ = d_90_cf103_
            rhs2_ = (d_78_v5_).f23
            rhs3_ = (0) - (d_92_v9_)
            lhs0_ = globalState
            lhs1_ = globalState
            lhs0_.f6 = rhs0_
            d_89_cf104_ = rhs1_
            d_91_cf102_ = rhs2_
            lhs1_.f6 = rhs3_
            d_93_v10_: _dafny.Seq
            d_93_v10_ = _dafny.SeqWithoutIsStrInference([d_92_v9_])
            d_94_v11_: _dafny.Array
            nw3_ = _dafny.Array(None, 22)
            nw3_[int(0)] = d_93_v10_
            nw3_[int(1)] = _dafny.SeqWithoutIsStrInference([d_92_v9_])
            nw3_[int(2)] = d_93_v10_
            nw3_[int(3)] = d_93_v10_
            nw3_[int(4)] = d_93_v10_
            nw3_[int(5)] = d_93_v10_
            nw3_[int(6)] = _dafny.SeqWithoutIsStrInference([d_92_v9_, d_92_v9_, d_92_v9_, (0) - (d_92_v9_), d_92_v9_])
            nw3_[int(7)] = d_93_v10_
            nw3_[int(8)] = d_93_v10_
            nw3_[int(9)] = d_93_v10_
            nw3_[int(10)] = _dafny.SeqWithoutIsStrInference([d_92_v9_])
            nw3_[int(11)] = _dafny.SeqWithoutIsStrInference([d_92_v9_ for d_95_i1_ in range(default__.abs(-471))])
            nw3_[int(12)] = (d_93_v10_).set(default__.safeIndex(d_92_v9_, len(d_93_v10_)), d_92_v9_)
            nw3_[int(13)] = d_93_v10_
            nw3_[int(14)] = d_93_v10_
            nw3_[int(15)] = d_93_v10_
            nw3_[int(16)] = d_93_v10_
            nw3_[int(17)] = d_93_v10_
            nw3_[int(18)] = d_93_v10_
            nw3_[int(19)] = d_93_v10_
            nw3_[int(20)] = d_93_v10_
            nw3_[int(21)] = d_93_v10_
            d_94_v11_ = nw3_
            d_96_v12_: D14
            d_96_v12_ = D14_DC38(d_94_v11_)
            d_97_v13_: _dafny.Seq
            d_97_v13_ = _dafny.SeqWithoutIsStrInference([D14_DC41(d_96_v12_)])
            d_92_v9_ = len(d_97_v13_)
            d_98_v14_: _dafny.Set
            d_98_v14_ = _dafny.Set({p3})
            d_99_v15_: str
            d_99_v15_ = _dafny.CodePoint('j')
            d_100_v16_: _dafny.Map
            d_100_v16_ = _dafny.Map({(p3).set(default__.safeIndex(d_92_v9_, len(p3)), d_99_v15_): d_92_v9_})
            d_101_v18_: _dafny.Set
            d_101_v18_ = _dafny.Set({d_88_cf105_})
            d_102_v19_: _dafny.Map
            d_102_v19_ = _dafny.Map({len(d_101_v18_): p3})
            d_103_v20_: _dafny.Seq
            d_103_v20_ = _dafny.SeqWithoutIsStrInference([d_88_cf105_])
            nw4_ = _dafny.Array(None, 11)
            nw4_[int(0)] = (d_78_v5_).f23
            nw4_[int(1)] = d_89_cf104_
            nw4_[int(2)] = (d_78_v5_).f23
            nw4_[int(3)] = d_88_cf105_
            nw4_[int(4)] = (d_78_v5_).f23
            nw4_[int(5)] = d_88_cf105_
            def iife24_():
                coll20_ = _dafny.Set()
                compr_20_: _dafny.Seq
                for compr_20_ in (d_100_v16_).keys.Elements:
                    d_104_v17_: _dafny.Seq = compr_20_
                    if (d_104_v17_) in (d_100_v16_):
                        coll20_ = coll20_.union(_dafny.Set([d_104_v17_]))
                return _dafny.Set(coll20_)
            nw4_[int(6)] = not(not((iife24_()
            ).issubset(d_98_v14_)))
            nw4_[int(7)] = (p3) < (p3)
            nw4_[int(8)] = (d_98_v14_).issubset(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctqonb")), ((d_102_v19_)[831] if (831) in (d_102_v19_) else p3), p3, p3, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqkhgippl"))}))
            nw4_[int(9)] = d_77_v4_
            nw4_[int(10)] = (d_78_v5_).fm17(len(p3), (D11_DC32((d_103_v20_)[default__.safeIndex(592, len(d_103_v20_))], d_92_v9_, True)).cf60, d_91_cf102_, globalState)
            d_72_v0_ = nw4_
            index0_ = default__.safeIndex(618, (p2).length(0))
            (p2)[index0_] = d_92_v9_
            d_105_v22_: _dafny.Map
            d_105_v22_ = _dafny.Map({default__.fm1((d_78_v5_).f23, globalState): not(d_88_cf105_)})
            d_106_v23_: _dafny.Seq
            d_106_v23_ = _dafny.SeqWithoutIsStrInference([d_105_v22_])
            d_107_v24_: C2
            nw5_ = C2()
            def iife25_():
                coll21_ = _dafny.Map()
                compr_21_: _dafny.Map
                for compr_21_ in (d_106_v23_).Elements:
                    d_108_v21_: _dafny.Map = compr_21_
                    if (d_108_v21_) in (d_106_v23_):
                        coll21_[d_108_v21_] = d_92_v9_
                return _dafny.Map(coll21_)
            nw5_.ctor__(d_92_v9_, d_92_v9_, iife25_()
            )
            d_107_v24_ = nw5_
            d_109_v25_: _dafny.MultiSet
            d_109_v25_ = _dafny.MultiSet([d_107_v24_, d_107_v24_, d_107_v24_])
            index1_ = default__.safeIndex(618, (p2).length(0))
            rhs4_ = d_99_v15_
            rhs5_ = default__.fm1((d_109_v25_).ispropersubset(d_109_v25_), globalState)
            lhs2_ = p2
            lhs3_ = default__.safeIndex(618, (p2).length(0))
            d_99_v15_ = rhs4_
            lhs2_[lhs3_] = rhs5_
        elif source3_.is_DC63:
            d_110___mcc_h4_ = source3_.cf106
            d_111___mcc_h5_ = source3_.cf107
            d_112_cf107_ = d_111___mcc_h5_
            d_113_cf106_ = d_110___mcc_h4_
            d_112_cf107_ = d_112_cf107_
            d_114_v26_: _dafny.Set
            d_114_v26_ = _dafny.Set({(d_78_v5_).f23})
            index2_ = default__.safeIndex(476, (d_72_v0_).length(0))
            (d_72_v0_)[index2_] = (d_114_v26_).ispropersubset(d_114_v26_)
            d_115_v27_: D8
            d_115_v27_ = D8_DC20((0) - (d_112_cf107_), d_112_cf107_, d_112_cf107_)
            index3_ = default__.safeIndex(476, (d_72_v0_).length(0))
            (d_72_v0_)[index3_] = (default__.fm52(d_115_v27_, d_112_cf107_, globalState)).cf10
            index4_ = default__.safeIndex(30, (d_79_v6_).length(0))
            (d_79_v6_)[index4_] = d_78_v5_
            index5_ = default__.safeIndex(30, (d_79_v6_).length(0))
            (d_79_v6_)[index5_] = d_78_v5_
            d_116_v28_: C0
            nw6_ = C0()
            nw6_.ctor__((d_78_v5_).f22, not((True if (d_78_v5_).f23 else not(d_77_v4_))))
            d_116_v28_ = nw6_
        elif True:
            d_117___mcc_h6_ = source3_.cf101
            d_118_cf101_ = d_117___mcc_h6_
            (globalState).f1 = d_77_v4_
            d_119_v29_: C4
            nw7_ = C4()
            nw7_.ctor__()
            d_119_v29_ = nw7_
            d_120_v30_: _dafny.Array
            nw8_ = _dafny.Array(_dafny.CodePoint('D'), 21)
            d_120_v30_ = nw8_
            d_121_v31_: str
            d_121_v31_ = _dafny.CodePoint('t')
            index6_ = default__.safeIndex(639, (d_120_v30_).length(0))
            (d_120_v30_)[index6_] = d_121_v31_
            d_122_v32_: _dafny.Seq
            d_122_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nctipsdh"))
            d_123_v33_: int
            d_123_v33_ = 177
            d_124_v34_: T0
            nw9_ = C5()
            nw9_.ctor__(d_123_v33_)
            d_124_v34_ = nw9_
            index7_ = default__.safeIndex(212, (d_79_v6_).length(0))
            (d_79_v6_)[index7_] = d_78_v5_
            index8_ = default__.safeIndex(639, (d_120_v30_).length(0))
            index9_ = default__.safeIndex(212, (d_79_v6_).length(0))
            rhs6_ = d_121_v31_
            rhs7_ = p3
            rhs8_ = d_124_v34_
            rhs9_ = d_78_v5_
            lhs4_ = d_120_v30_
            lhs5_ = default__.safeIndex(639, (d_120_v30_).length(0))
            lhs6_ = d_79_v6_
            lhs7_ = default__.safeIndex(212, (d_79_v6_).length(0))
            lhs4_[lhs5_] = rhs6_
            d_122_v32_ = rhs7_
            d_124_v34_ = rhs8_
            lhs6_[lhs7_] = rhs9_
            d_125_v35_: _dafny.Seq
            d_125_v35_ = _dafny.SeqWithoutIsStrInference([-553])
            d_126_v36_: _dafny.Seq
            d_126_v36_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_125_v35_, d_125_v35_])])
            (globalState).f6 = default__.safeModuloInt(len((d_126_v36_)[default__.safeIndex((D7_DC16(True, d_123_v33_, (d_78_v5_).f23)).cf22, len(d_126_v36_))]), d_123_v33_)
        d_127_v37_: _dafny.Array
        nw10_ = _dafny.Array(int(0), 10)
        d_127_v37_ = nw10_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_127_v37_).length(0)):
            d_128_i2_: int = guard_loop_0_
            if (True) and (((0) <= (d_128_i2_)) and ((d_128_i2_) < ((d_127_v37_).length(0)))):
                (d_127_v37_)[(d_128_i2_)] = (d_128_i2_) + (len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([d_77_v4_]))): -253})))
        d_129_v38_: int
        d_129_v38_ = 117
        (globalState).f1 = (d_78_v5_).fm17(d_129_v38_, d_129_v38_, d_77_v4_, globalState)
        d_130_i3_: int
        d_130_i3_ = 0
        with _dafny.label("0"):
            while (p3) < (p3):
                with _dafny.c_label("0"):
                    if (d_130_i3_) >= (100):
                        raise _dafny.Break("0")
                    d_130_i3_ = (d_130_i3_) + (1)
                    d_131_v39_: _dafny.Map
                    d_131_v39_ = _dafny.Map({(d_78_v5_).f23: (d_78_v5_).f23})
                    d_132_v40_: D1
                    d_132_v40_ = D1_DC3((d_131_v39_).set(d_77_v4_, d_77_v4_))
                    pat_let_tv1_ = d_131_v39_
                    d_133_v41_: _dafny.Set
                    def iife26_(_pat_let2_0):
                        def iife27_(d_134_dt__update__tmp_h1_):
                            def iife28_(_pat_let3_0):
                                def iife29_(d_135_dt__update_hcf3_h0_):
                                    return D1_DC3(d_135_dt__update_hcf3_h0_)
                                return iife29_(_pat_let3_0)
                            return iife28_(pat_let_tv1_)
                        return iife27_(_pat_let2_0)
                    d_133_v41_ = _dafny.Set({iife26_(d_132_v40_), d_132_v40_})
                    d_136_v42_: _dafny.Map
                    d_136_v42_ = _dafny.Map({len(d_133_v41_): d_129_v38_})
                    d_136_v42_ = (d_136_v42_).set(d_129_v38_, d_129_v38_)
                    d_137_v43_: _dafny.Seq
                    d_137_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dicognbwi"))
                    d_138_v44_: str
                    d_138_v44_ = _dafny.CodePoint('j')
                    d_137_v43_ = (p3) + (((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_139_i4_ in range(default__.abs(224))])) + (p3)).set(default__.safeIndex(d_129_v38_, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_140_i4_ in range(default__.abs(224))])) + (p3))), d_138_v44_))
                    if default__.fm2(globalState):
                        index10_ = default__.safeIndex(534, (d_127_v37_).length(0))
                        (d_127_v37_)[index10_] = default__.safeDivisionInt(len(d_136_v42_), d_129_v38_)
                        d_141_v45_: _dafny.Map
                        d_141_v45_ = _dafny.Map({d_129_v38_: p3})
                        index11_ = default__.safeIndex(534, (d_127_v37_).length(0))
                        (d_127_v37_)[index11_] = (len(((d_141_v45_)[d_129_v38_] if (d_129_v38_) in (d_141_v45_) else (p3).set(default__.safeIndex(d_129_v38_, len(p3)), _dafny.CodePoint('m'))))) - (d_129_v38_)
                        (globalState).f6 = -970
                        d_142_v46_: D20
                        d_142_v46_ = D20_DC50((d_78_v5_).f23, d_129_v38_, d_129_v38_, True, False)
                        (globalState).f1 = ((d_142_v46_).cf87) > (((d_127_v37_)[default__.safeIndex(534, (d_127_v37_).length(0))]) + (d_129_v38_))
                        index12_ = default__.safeIndex(216, (d_72_v0_).length(0))
                        (d_72_v0_)[index12_] = (d_129_v38_) <= (len(p3))
                        index13_ = default__.safeIndex(534, (d_127_v37_).length(0))
                        index14_ = default__.safeIndex(216, (d_72_v0_).length(0))
                        rhs10_ = -76
                        rhs11_ = (p0).ispropersubset(p0)
                        rhs12_ = 375
                        lhs8_ = d_127_v37_
                        lhs9_ = default__.safeIndex(534, (d_127_v37_).length(0))
                        lhs10_ = d_72_v0_
                        lhs11_ = default__.safeIndex(216, (d_72_v0_).length(0))
                        lhs12_ = globalState
                        lhs8_[lhs9_] = rhs10_
                        lhs10_[lhs11_] = rhs11_
                        lhs12_.f6 = rhs12_
                        (globalState).f6 = (443) - ((798) * (54))
                    elif True:
                        (globalState).f6 = d_129_v38_
                        d_143_v47_: _dafny.Map
                        d_143_v47_ = _dafny.Map({d_138_v44_: not((d_78_v5_).f23)})
                        index15_ = default__.safeIndex(28, (d_72_v0_).length(0))
                        (d_72_v0_)[index15_] = ((d_143_v47_)[d_138_v44_] if (d_138_v44_) in (d_143_v47_) else (d_78_v5_).f23)
                        index16_ = default__.safeIndex(28, (d_72_v0_).length(0))
                        (d_72_v0_)[index16_] = (d_137_v43_) != (d_137_v43_)
                        d_144_v48_: _dafny.Array
                        nw11_ = _dafny.Array(_dafny.Array(None, 0), 11)
                        d_144_v48_ = nw11_
                        d_145_v49_: _dafny.Array
                        nw12_ = _dafny.Array(D10.default()(), 27)
                        d_145_v49_ = nw12_
                        index17_ = default__.safeIndex(450, (d_144_v48_).length(0))
                        (d_144_v48_)[index17_] = d_145_v49_
                        index18_ = default__.safeIndex(450, (d_144_v48_).length(0))
                        (d_144_v48_)[index18_] = d_145_v49_
                        index19_ = default__.safeIndex(581, (d_127_v37_).length(0))
                        (d_127_v37_)[index19_] = (d_129_v38_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_77_v4_])])))
                        d_146_v50_: _dafny.Map
                        d_146_v50_ = _dafny.Map({d_129_v38_: (d_78_v5_).f23})
                        d_147_v51_: _dafny.Seq
                        d_147_v51_ = _dafny.SeqWithoutIsStrInference([len(d_137_v43_), d_129_v38_, -261, d_129_v38_])
                        index20_ = default__.safeIndex(581, (d_127_v37_).length(0))
                        rhs13_ = (default__.safeModuloInt(len(d_137_v43_), d_129_v38_)) == (d_129_v38_)
                        rhs14_ = (len(d_146_v50_)) + (default__.safeDivisionInt(d_129_v38_, (0) - ((d_147_v51_)[default__.safeIndex(d_129_v38_, len(d_147_v51_))])))
                        lhs13_ = globalState
                        lhs14_ = d_127_v37_
                        lhs15_ = default__.safeIndex(581, (d_127_v37_).length(0))
                        lhs13_.f1 = rhs13_
                        lhs14_[lhs15_] = rhs14_
                        d_148_v52_: _dafny.Array
                        nw13_ = _dafny.Array(D3.default()(), 20)
                        d_148_v52_ = nw13_
                        d_149_v53_: _dafny.Set
                        d_149_v53_ = _dafny.Set({d_148_v52_, d_148_v52_})
                        d_150_v54_: D9
                        d_150_v54_ = D9_DC23(d_149_v53_, p3, (d_78_v5_).f23, p2, _dafny.CodePoint('i'))
                        d_150_v54_ = D9_DC23(_dafny.Set({d_148_v52_, d_148_v52_, d_148_v52_, d_148_v52_}), (p3).set(default__.safeIndex(d_129_v38_, len(p3)), _dafny.CodePoint('a')), (d_72_v0_)[default__.safeIndex(28, (d_72_v0_).length(0))], p2, d_138_v44_)
                    (globalState).f1 = not(d_77_v4_)
                    pass
            pass
        hi0_ = d_129_v38_
        for d_151_i5_ in range(d_129_v38_, hi0_):
            d_152_v55_: _dafny.Seq
            d_152_v55_ = _dafny.SeqWithoutIsStrInference([d_129_v38_, d_151_i5_, d_151_i5_])
            (globalState).f6 = (default__.safeModuloInt(len(d_152_v55_), d_151_i5_)) - (default__.safeModuloInt(d_151_i5_, len(default__.fm21(globalState))))
            (globalState).f1 = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_153_i6_ in range(default__.abs(742))])) <= (p3)
            index21_ = default__.safeIndex(850, (d_72_v0_).length(0))
            (d_72_v0_)[index21_] = (d_78_v5_).f23
            index22_ = default__.safeIndex(497, (p2).length(0))
            (p2)[index22_] = d_129_v38_
            d_154_v56_: _dafny.Seq
            d_154_v56_ = _dafny.SeqWithoutIsStrInference([d_78_v5_, d_78_v5_])
            d_155_v57_: D3
            d_155_v57_ = D3_DC7(d_151_i5_)
            index23_ = default__.safeIndex(850, (d_72_v0_).length(0))
            index24_ = default__.safeIndex(497, (p2).length(0))
            rhs15_ = ((d_129_v38_) + (d_129_v38_)) != (520)
            rhs16_ = (d_129_v38_) * (len(d_154_v56_))
            rhs17_ = (d_155_v57_).cf7
            lhs16_ = d_72_v0_
            lhs17_ = default__.safeIndex(850, (d_72_v0_).length(0))
            lhs18_ = p2
            lhs19_ = default__.safeIndex(497, (p2).length(0))
            lhs20_ = globalState
            lhs16_[lhs17_] = rhs15_
            lhs18_[lhs19_] = rhs16_
            lhs20_.f6 = rhs17_
            index25_ = default__.safeIndex(497, (p2).length(0))
            (p2)[index25_] = d_129_v38_
        d_127_v37_ = p2

    @staticmethod
    def Main(noArgsParameter__):
        d_156_v0_: bool
        d_156_v0_ = True
        d_157_v1_: _dafny.MultiSet
        d_157_v1_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_156_v0_]))])
        d_158_v2_: _dafny.Map
        d_158_v2_ = _dafny.Map({d_157_v1_: False})
        d_159_v3_: int
        d_159_v3_ = 781
        d_160_v4_: _dafny.Map
        d_160_v4_ = _dafny.Map({d_159_v3_: d_156_v0_})
        d_161_v5_: _dafny.MultiSet
        d_161_v5_ = _dafny.MultiSet([_dafny.Map({d_159_v3_: d_156_v0_}), d_160_v4_, d_160_v4_, d_160_v4_])
        d_162_v6_: _dafny.Seq
        d_162_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kajndjn"))
        d_163_v7_: _dafny.Seq
        d_163_v7_ = _dafny.SeqWithoutIsStrInference([d_162_v6_])
        d_164_v8_: _dafny.Array
        def lambda1_(d_165_v6_):
            def lambda2_(d_166_i0_):
                return (d_166_i0_) * (len(d_165_v6_))

            return lambda2_

        init1_ = lambda1_(d_162_v6_)
        nw14_ = _dafny.Array(None, 4)
        for i0_1_ in range(nw14_.length(0)):
            nw14_[i0_1_] = init1_(i0_1_)
        d_164_v8_ = nw14_
        d_167_v9_: _dafny.Set
        d_167_v9_ = _dafny.Set({(0) - (d_159_v3_), d_159_v3_})
        d_168_v10_: _dafny.Array
        def lambda3_(d_169_v0_):
            def lambda4_(d_170_i1_):
                return d_169_v0_

            return lambda4_

        init2_ = lambda3_(d_156_v0_)
        nw15_ = _dafny.Array(None, 9)
        for i0_2_ in range(nw15_.length(0)):
            nw15_[i0_2_] = init2_(i0_2_)
        d_168_v10_ = nw15_
        d_171_globalState_: GlobalState
        nw16_ = GlobalState()
        nw16_.ctor__(-19, False, True, 989, d_158_v2_, (d_161_v5_) - (_dafny.MultiSet([d_160_v4_, d_160_v4_])), 523, False, (d_163_v7_) + (d_163_v7_), d_164_v8_, False, d_167_v9_, True, 565, d_168_v10_)
        d_171_globalState_ = nw16_
        d_172_v11_: str
        d_172_v11_ = _dafny.CodePoint('w')
        d_172_v11_ = d_172_v11_
        d_173_v12_: _dafny.Seq
        d_173_v12_ = _dafny.SeqWithoutIsStrInference([True])
        d_159_v3_ = (0) - (default__.safeDivisionInt(len(d_173_v12_), (0) - (((0) - (len(d_162_v6_))) * (d_159_v3_))))
        d_174_v13_: _dafny.Array
        nw17_ = _dafny.Array(None, 3)
        nw17_[int(0)] = d_164_v8_
        nw17_[int(1)] = d_164_v8_
        nw17_[int(2)] = d_164_v8_
        d_174_v13_ = nw17_
        nw18_ = _dafny.Array(None, 1)
        nw18_[int(0)] = d_164_v8_
        d_174_v13_ = nw18_
        d_175_v14_: _dafny.Array
        def lambda5_(d_176_v4_):
            def lambda6_(d_177_i2_):
                return d_176_v4_

            return lambda6_

        init3_ = lambda5_(d_160_v4_)
        nw19_ = _dafny.Array(None, 9)
        for i0_3_ in range(nw19_.length(0)):
            nw19_[i0_3_] = init3_(i0_3_)
        d_175_v14_ = nw19_
        index26_ = default__.safeIndex(696, (d_175_v14_).length(0))
        (d_175_v14_)[index26_] = _dafny.Map({d_159_v3_: d_156_v0_})
        index27_ = default__.safeIndex(696, (d_175_v14_).length(0))
        (d_175_v14_)[index27_] = ((d_160_v4_) | (_dafny.Map({d_159_v3_: d_156_v0_}))) | (default__.fm0(d_156_v0_, False, d_159_v3_, d_171_globalState_))
        hi1_ = d_159_v3_
        for d_178_i3_ in range(len(d_173_v12_), hi1_):
            index28_ = default__.safeIndex(390, (d_168_v10_).length(0))
            (d_168_v10_)[index28_] = d_156_v0_
            index29_ = default__.safeIndex(390, (d_168_v10_).length(0))
            (d_168_v10_)[index29_] = False
            index30_ = default__.safeIndex(390, (d_168_v10_).length(0))
            (d_168_v10_)[index30_] = not((d_168_v10_)[default__.safeIndex(390, (d_168_v10_).length(0))])
            d_179_v15_: _dafny.MultiSet
            d_179_v15_ = _dafny.MultiSet([d_164_v8_, d_164_v8_, d_164_v8_])
            d_179_v15_ = (d_179_v15_) - ((d_179_v15_).intersection(_dafny.MultiSet([d_164_v8_])))
            d_180_v16_: _dafny.Map
            d_180_v16_ = _dafny.Map({d_178_i3_: d_178_i3_})
            d_181_v17_: _dafny.Map
            d_181_v17_ = _dafny.Map({d_173_v12_: (len(d_180_v16_) if d_156_v0_ else d_178_i3_)})
            d_182_v18_: D0
            d_182_v18_ = D0_DC0(d_156_v0_)
            d_181_v17_ = (d_181_v17_).set(d_173_v12_, default__.fm1((d_182_v18_).cf0, d_171_globalState_))
        (d_171_globalState_).f9 = d_164_v8_
        d_183_v19_: _dafny.Map
        d_183_v19_ = _dafny.Map({d_164_v8_: d_168_v10_})
        d_183_v19_ = _dafny.Map({d_164_v8_: d_168_v10_})
        if (d_173_v12_)[default__.safeIndex((len(d_162_v6_)) + (d_159_v3_), len(d_173_v12_))]:
            d_184_v20_: _dafny.Map
            d_184_v20_ = _dafny.Map({d_156_v0_: d_159_v3_})
            d_185_v21_: D0
            d_185_v21_ = D0_DC0(d_156_v0_)
            d_186_v22_: D0
            d_186_v22_ = D0_DC2(d_185_v21_)
            d_187_v23_: D0
            d_187_v23_ = D0_DC2(d_186_v22_)
            pat_let_tv2_ = d_186_v22_
            d_188_v24_: _dafny.Array
            nw20_ = _dafny.Array(None, 13)
            nw20_[int(0)] = d_187_v23_
            nw20_[int(1)] = d_187_v23_
            nw20_[int(2)] = D0_DC2(d_186_v22_)
            nw20_[int(3)] = d_187_v23_
            def iife30_(_pat_let4_0):
                def iife31_(d_189_dt__update__tmp_h0_):
                    def iife32_(_pat_let5_0):
                        def iife33_(d_190_dt__update_hcf2_h0_):
                            return D0_DC2(d_190_dt__update_hcf2_h0_)
                        return iife33_(_pat_let5_0)
                    return iife32_(pat_let_tv2_)
                return iife31_(_pat_let4_0)
            nw20_[int(4)] = iife30_(d_187_v23_)
            nw20_[int(5)] = D0_DC2(d_185_v21_)
            nw20_[int(6)] = d_187_v23_
            nw20_[int(7)] = d_187_v23_
            nw20_[int(8)] = d_187_v23_
            nw20_[int(9)] = d_187_v23_
            nw20_[int(10)] = d_187_v23_
            nw20_[int(11)] = d_187_v23_
            nw20_[int(12)] = d_187_v23_
            d_188_v24_ = nw20_
            default__.m0(_dafny.Set({len(d_184_v20_), d_159_v3_, default__.fm1(d_156_v0_, d_171_globalState_)}), d_188_v24_, d_164_v8_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvsasg")), d_171_globalState_)
            d_191_v25_: _dafny.Array
            nw21_ = _dafny.Array(_dafny.CodePoint('D'), 16)
            d_191_v25_ = nw21_
            d_192_v26_: _dafny.MultiSet
            d_192_v26_ = _dafny.MultiSet([d_191_v25_])
            d_192_v26_ = _dafny.MultiSet([(d_191_v25_ if d_156_v0_ else d_191_v25_)])
            (d_171_globalState_).f1 = d_156_v0_
            index31_ = default__.safeIndex(992, (d_168_v10_).length(0))
            (d_168_v10_)[index31_] = d_156_v0_
            index32_ = default__.safeIndex(992, (d_168_v10_).length(0))
            (d_168_v10_)[index32_] = True
            default__.m0(d_167_v9_, d_188_v24_, d_164_v8_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwejgvdb")), d_171_globalState_)
        elif True:
            if d_156_v0_:
                (d_171_globalState_).f1 = d_156_v0_
                d_193_v27_: _dafny.MultiSet
                d_193_v27_ = _dafny.MultiSet([d_156_v0_, d_156_v0_])
                d_194_v28_: _dafny.Map
                d_194_v28_ = _dafny.Map({d_156_v0_: d_159_v3_})
                d_195_v29_: _dafny.Seq
                d_195_v29_ = _dafny.SeqWithoutIsStrInference([d_159_v3_])
                d_196_v30_: _dafny.Seq
                d_196_v30_ = _dafny.SeqWithoutIsStrInference([len(d_194_v28_), len(d_167_v9_), len(d_195_v29_)])
                d_197_v31_: _dafny.Map
                d_197_v31_ = _dafny.Map({d_196_v30_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xngsnlvv"))})
                d_162_v6_ = ((d_162_v6_).set(default__.safeIndex((0) - (((d_193_v27_)[d_156_v0_] if (d_156_v0_) in (d_193_v27_) else d_159_v3_)), len(d_162_v6_)), d_172_v11_)) + (((d_197_v31_)[d_195_v29_] if (d_195_v29_) in (d_197_v31_) else d_162_v6_))
                rhs18_ = ((621) * (d_159_v3_)) * (d_159_v3_)
                rhs19_ = (d_159_v3_) + (398)
                rhs20_ = d_156_v0_
                rhs21_ = d_156_v0_
                lhs21_ = d_171_globalState_
                lhs22_ = d_171_globalState_
                lhs23_ = d_171_globalState_
                lhs21_.f6 = rhs18_
                lhs22_.f6 = rhs19_
                lhs23_.f1 = rhs20_
                d_156_v0_ = rhs21_
                (d_171_globalState_).f1 = d_156_v0_
                d_198_v32_: D0
                d_198_v32_ = D0_DC1(default__.fm1(d_156_v0_, d_171_globalState_))
                rhs22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "el"))
                rhs23_ = D0_DC1(d_159_v3_)
                d_162_v6_ = rhs22_
                d_198_v32_ = rhs23_
            elif True:
                d_162_v6_ = d_162_v6_
                d_159_v3_ = d_159_v3_
                d_199_v34_: _dafny.Map
                d_199_v34_ = _dafny.Map({(d_157_v1_).cardinality: d_172_v11_})
                d_200_v35_: _dafny.Map
                def iife34_():
                    coll22_ = _dafny.Map()
                    compr_22_: int
                    for compr_22_ in (d_199_v34_).keys.Elements:
                        d_201_v33_: int = compr_22_
                        if (d_201_v33_) in (d_199_v34_):
                            coll22_[(d_201_v33_) - (d_159_v3_)] = d_159_v3_
                    return _dafny.Map(coll22_)
                d_200_v35_ = _dafny.Map({False: iife34_()
                })
                d_202_v36_: _dafny.Seq
                d_202_v36_ = _dafny.SeqWithoutIsStrInference([d_159_v3_])
                d_203_v37_: _dafny.Seq
                d_203_v37_ = _dafny.SeqWithoutIsStrInference([len(d_173_v12_), len(d_202_v36_)])
                d_204_v39_: _dafny.Map
                d_204_v39_ = _dafny.Map({d_172_v11_: d_156_v0_})
                d_205_v40_: _dafny.Map
                def iife35_():
                    coll23_ = _dafny.Map()
                    compr_23_: str
                    for compr_23_ in (d_204_v39_).keys.Elements:
                        d_206_v38_: str = compr_23_
                        if (d_206_v38_) in (d_204_v39_):
                            coll23_[d_206_v38_] = d_156_v0_
                    return _dafny.Map(coll23_)
                d_205_v40_ = _dafny.Map({(d_203_v37_)[default__.safeIndex((d_203_v37_)[default__.safeIndex(len(iife35_()
                ), len(d_203_v37_))], len(d_203_v37_))]: d_159_v3_})
                d_200_v35_ = (d_200_v35_).set(d_156_v0_, (d_205_v40_) | (d_205_v40_))
                index33_ = default__.safeIndex(503, (d_164_v8_).length(0))
                (d_164_v8_)[index33_] = len(_dafny.Map({default__.fm2(d_171_globalState_): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxdfyxg"))}))
                index34_ = default__.safeIndex(503, (d_164_v8_).length(0))
                (d_164_v8_)[index34_] = (d_159_v3_) + (d_159_v3_)
                d_207_v41_: T0
                nw22_ = C6()
                nw22_.ctor__()
                d_207_v41_ = nw22_
                d_208_v42_: _dafny.MultiSet
                d_208_v42_ = _dafny.MultiSet([d_207_v41_])
                d_209_v43_: C9
                nw23_ = C9()
                nw23_.ctor__(((d_164_v8_)[default__.safeIndex(503, (d_164_v8_).length(0))] if d_156_v0_ else (0) - ((d_208_v42_).cardinality)))
                d_209_v43_ = nw23_
            index35_ = default__.safeIndex(669, (d_168_v10_).length(0))
            (d_168_v10_)[index35_] = False
            index36_ = default__.safeIndex(669, (d_168_v10_).length(0))
            (d_168_v10_)[index36_] = (d_156_v0_) and (d_156_v0_)
            d_210_v44_: _dafny.Seq
            d_210_v44_ = _dafny.SeqWithoutIsStrInference([d_159_v3_])
            d_210_v44_ = (_dafny.SeqWithoutIsStrInference([(0) - (d_159_v3_) for d_211_i4_ in range(default__.abs(495))]) if (-916) > (d_159_v3_) else d_210_v44_)
            d_212_v45_: _dafny.Map
            d_212_v45_ = _dafny.Map({d_159_v3_: d_159_v3_})
            d_213_v47_: _dafny.Map
            def iife36_():
                coll24_ = _dafny.Set()
                compr_24_: int
                for compr_24_ in (d_212_v45_).keys.Elements:
                    d_214_v46_: int = compr_24_
                    if (d_214_v46_) in (d_212_v45_):
                        coll24_ = coll24_.union(_dafny.Set([default__.safeModuloInt(d_214_v46_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "co"))))]))
                return _dafny.Set(coll24_)
            d_213_v47_ = _dafny.Map({d_156_v0_: iife36_()
            })
            d_215_v48_: _dafny.Array
            nw24_ = _dafny.Array(None, 13)
            nw24_[int(0)] = (d_213_v47_) | (d_213_v47_)
            nw24_[int(1)] = d_213_v47_
            nw24_[int(2)] = d_213_v47_
            nw24_[int(3)] = (d_213_v47_) | (d_213_v47_)
            nw24_[int(4)] = ((d_213_v47_).set(d_156_v0_, _dafny.Set({d_159_v3_, 576}))) | (d_213_v47_)
            nw24_[int(5)] = d_213_v47_
            nw24_[int(6)] = d_213_v47_
            nw24_[int(7)] = (d_213_v47_) | (d_213_v47_)
            nw24_[int(8)] = _dafny.Map({not(default__.fm2(d_171_globalState_)): d_167_v9_})
            nw24_[int(9)] = d_213_v47_
            nw24_[int(10)] = d_213_v47_
            nw24_[int(11)] = _dafny.Map({(d_168_v10_)[default__.safeIndex(669, (d_168_v10_).length(0))]: d_167_v9_})
            nw24_[int(12)] = (d_213_v47_) | (d_213_v47_)
            d_215_v48_ = nw24_
            index37_ = default__.safeIndex(891, (d_215_v48_).length(0))
            (d_215_v48_)[index37_] = (d_213_v47_) | (d_213_v47_)
            d_216_v49_: _dafny.Array
            nw25_ = _dafny.Array(False, 12)
            d_216_v49_ = nw25_
            d_217_v50_: _dafny.Map
            d_217_v50_ = _dafny.Map({(d_168_v10_)[default__.safeIndex(669, (d_168_v10_).length(0))]: d_216_v49_})
            d_218_v51_: _dafny.Set
            d_218_v51_ = _dafny.Set({(d_168_v10_)[default__.safeIndex(669, (d_168_v10_).length(0))]})
            index38_ = default__.safeIndex(891, (d_215_v48_).length(0))
            rhs24_ = d_159_v3_
            rhs25_ = (d_168_v10_)[default__.safeIndex(669, (d_168_v10_).length(0))]
            rhs26_ = ((d_217_v50_)[((d_168_v10_)[default__.safeIndex(669, (d_168_v10_).length(0))]) in (d_218_v51_)] if (((d_168_v10_)[default__.safeIndex(669, (d_168_v10_).length(0))]) in (d_218_v51_)) in (d_217_v50_) else d_216_v49_)
            rhs27_ = (_dafny.Map({(d_168_v10_)[default__.safeIndex(669, (d_168_v10_).length(0))]: d_167_v9_})) | (d_213_v47_)
            lhs24_ = d_171_globalState_
            lhs25_ = d_171_globalState_
            lhs26_ = d_171_globalState_
            lhs27_ = d_215_v48_
            lhs28_ = default__.safeIndex(891, (d_215_v48_).length(0))
            lhs24_.f6 = rhs24_
            lhs25_.f1 = rhs25_
            lhs26_.f14 = rhs26_
            lhs27_[lhs28_] = rhs27_
            if d_156_v0_:
                index39_ = default__.safeIndex(449, (d_164_v8_).length(0))
                (d_164_v8_)[index39_] = d_159_v3_
                index40_ = default__.safeIndex(669, (d_168_v10_).length(0))
                index41_ = default__.safeIndex(449, (d_164_v8_).length(0))
                rhs28_ = (d_168_v10_)[default__.safeIndex(669, (d_168_v10_).length(0))]
                rhs29_ = (d_159_v3_) + ((0) - ((0) - (d_159_v3_)))
                lhs29_ = d_168_v10_
                lhs30_ = default__.safeIndex(669, (d_168_v10_).length(0))
                lhs31_ = d_164_v8_
                lhs32_ = default__.safeIndex(449, (d_164_v8_).length(0))
                lhs29_[lhs30_] = rhs28_
                lhs31_[lhs32_] = rhs29_
                index42_ = default__.safeIndex(449, (d_164_v8_).length(0))
                (d_164_v8_)[index42_] = d_159_v3_
                d_219_v52_: _dafny.Map
                d_219_v52_ = _dafny.Map({(d_175_v14_)[default__.safeIndex(696, (d_175_v14_).length(0))]: d_159_v3_})
                d_220_v53_: T1
                nw26_ = C2()
                nw26_.ctor__((d_164_v8_)[default__.safeIndex(449, (d_164_v8_).length(0))], d_159_v3_, d_219_v52_)
                d_220_v53_ = nw26_
                d_162_v6_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flwge"))).set(default__.safeIndex((_dafny.MultiSet([d_220_v53_, d_220_v53_, d_220_v53_])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flwge")))), d_172_v11_)
                d_218_v51_ = (d_218_v51_).intersection((d_218_v51_) | (d_218_v51_))
                d_221_v54_: _dafny.Array
                nw27_ = _dafny.Array(D0.default()(), 17)
                d_221_v54_ = nw27_
                d_222_v55_: _dafny.Array
                nw28_ = _dafny.Array(int(0), 4)
                d_222_v55_ = nw28_
                default__.m0(d_167_v9_, d_221_v54_, d_222_v55_, d_162_v6_, d_171_globalState_)
            elif True:
                d_223_v56_: _dafny.Map
                d_223_v56_ = _dafny.Map({(d_175_v14_)[default__.safeIndex(696, (d_175_v14_).length(0))]: d_159_v3_})
                d_224_v57_: C2
                nw29_ = C2()
                nw29_.ctor__(d_159_v3_, d_159_v3_, (d_223_v56_) | (d_223_v56_))
                d_224_v57_ = nw29_
                (d_171_globalState_).f1 = (d_224_v57_.f21) == (d_159_v3_)
                d_225_v58_: _dafny.Array
                nw30_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
                d_225_v58_ = nw30_
                d_225_v58_ = d_225_v58_
                d_226_v59_: _dafny.Array
                nw31_ = _dafny.Array(_dafny.Set({}), 25)
                d_226_v59_ = nw31_
                index43_ = default__.safeIndex(909, (d_226_v59_).length(0))
                (d_226_v59_)[index43_] = _dafny.Set({d_162_v6_})
                index44_ = default__.safeIndex(909, (d_226_v59_).length(0))
                (d_226_v59_)[index44_] = _dafny.Set({(d_162_v6_ if (d_168_v10_)[default__.safeIndex(669, (d_168_v10_).length(0))] else default__.fm21(d_171_globalState_)), default__.fm21(d_171_globalState_)})
                d_227_v60_: _dafny.Array
                nw32_ = _dafny.Array(_dafny.Set({}), 17)
                d_227_v60_ = nw32_
                d_228_v61_: D11
                d_228_v61_ = D11_DC31(d_173_v12_, len(d_162_v6_), (d_224_v57_).f20, (d_224_v57_).f20, d_159_v3_)
                d_229_v62_: _dafny.Set
                d_229_v62_ = _dafny.Set({d_228_v61_})
                index45_ = default__.safeIndex(747, (d_227_v60_).length(0))
                (d_227_v60_)[index45_] = (d_229_v62_) - (_dafny.Set({d_228_v61_}))
                d_230_v63_: D3
                d_230_v63_ = D3_DC7(len(d_212_v45_))
                d_231_v64_: D10
                d_231_v64_ = D10_DC27(d_224_v57_.f21, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "upuot")), (d_224_v57_).f20, d_230_v63_, _dafny.Map({d_216_v49_: (d_168_v10_)[default__.safeIndex(669, (d_168_v10_).length(0))]}))
                index46_ = default__.safeIndex(747, (d_227_v60_).length(0))
                (d_227_v60_)[index46_] = default__.fm55((0) - (d_159_v3_), (d_231_v64_).cf48, d_159_v3_, d_171_globalState_)
        d_172_v11_ = d_172_v11_
        if not (d_156_v0_) or (d_156_v0_):
            d_232_v65_: _dafny.Set
            d_232_v65_ = _dafny.Set({d_172_v11_, d_172_v11_, _dafny.CodePoint('p')})
            d_232_v65_ = d_232_v65_
            d_233_v66_: _dafny.Map
            d_233_v66_ = _dafny.Map({(d_175_v14_)[default__.safeIndex(696, (d_175_v14_).length(0))]: d_159_v3_})
            d_234_v67_: T1
            nw33_ = C3()
            nw33_.ctor__(d_156_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "upii"))), d_233_v66_)
            d_234_v67_ = nw33_
            d_235_v68_: _dafny.Set
            d_235_v68_ = _dafny.Set({d_156_v0_})
            d_236_v69_: D11
            d_236_v69_ = D11_DC31(d_173_v12_, d_159_v3_, d_159_v3_, len(d_235_v68_), d_159_v3_)
            d_237_v70_: _dafny.Array
            nw34_ = _dafny.Array(None, 19)
            nw34_[int(0)] = d_173_v12_
            nw34_[int(1)] = d_173_v12_
            nw34_[int(2)] = d_173_v12_
            nw34_[int(3)] = d_173_v12_
            nw34_[int(4)] = _dafny.SeqWithoutIsStrInference([d_156_v0_, not(not(d_156_v0_)), d_156_v0_, d_156_v0_, d_156_v0_])
            nw34_[int(5)] = _dafny.SeqWithoutIsStrInference([d_156_v0_])
            nw34_[int(6)] = d_173_v12_
            nw34_[int(7)] = (d_173_v12_) + (d_173_v12_)
            nw34_[int(8)] = (d_173_v12_) + (d_173_v12_)
            nw34_[int(9)] = d_173_v12_
            nw34_[int(10)] = d_173_v12_
            nw34_[int(11)] = d_173_v12_
            nw34_[int(12)] = (d_173_v12_).set(default__.safeIndex(default__.fm1(d_156_v0_, d_171_globalState_), len(d_173_v12_)), d_156_v0_)
            nw34_[int(13)] = ((d_173_v12_) + ((d_236_v69_).cf54)).set(default__.safeIndex(d_159_v3_, len((d_173_v12_) + ((d_236_v69_).cf54))), d_156_v0_)
            nw34_[int(14)] = default__.fm19(d_171_globalState_)
            nw34_[int(15)] = d_173_v12_
            nw34_[int(16)] = d_173_v12_
            nw34_[int(17)] = d_173_v12_
            nw34_[int(18)] = (_dafny.SeqWithoutIsStrInference([d_156_v0_])).set(default__.safeIndex(309, len(_dafny.SeqWithoutIsStrInference([d_156_v0_]))), d_156_v0_)
            d_237_v70_ = nw34_
            d_238_v71_: _dafny.Seq
            d_238_v71_ = _dafny.SeqWithoutIsStrInference([d_173_v12_, d_173_v12_, _dafny.SeqWithoutIsStrInference([d_156_v0_])])
            index47_ = default__.safeIndex(883, (d_237_v70_).length(0))
            (d_237_v70_)[index47_] = (d_238_v71_)[default__.safeIndex(d_159_v3_, len(d_238_v71_))]
            d_239_v72_: _dafny.Seq
            d_239_v72_ = _dafny.SeqWithoutIsStrInference([d_168_v10_, d_168_v10_])
            index48_ = default__.safeIndex(883, (d_237_v70_).length(0))
            rhs30_ = d_156_v0_
            rhs31_ = len(d_239_v72_)
            rhs32_ = default__.fm2(d_171_globalState_)
            rhs33_ = d_234_v67_
            rhs34_ = d_173_v12_
            lhs33_ = d_171_globalState_
            lhs34_ = d_171_globalState_
            lhs35_ = d_171_globalState_
            lhs36_ = d_237_v70_
            lhs37_ = default__.safeIndex(883, (d_237_v70_).length(0))
            lhs33_.f1 = rhs30_
            lhs34_.f6 = rhs31_
            lhs35_.f1 = rhs32_
            d_234_v67_ = rhs33_
            lhs36_[lhs37_] = rhs34_
            d_240_v73_: _dafny.Seq
            d_240_v73_ = _dafny.SeqWithoutIsStrInference([(0) - (-359)])
            d_240_v73_ = d_240_v73_
            (d_171_globalState_).f1 = default__.fm2(d_171_globalState_)
            if d_156_v0_:
                d_241_v74_: _dafny.Map
                d_241_v74_ = _dafny.Map({d_156_v0_: d_156_v0_})
                d_242_v75_: str
                out0_: str
                out0_ = (d_234_v67_).m11(_dafny.SeqWithoutIsStrInference([d_172_v11_ for d_243_i5_ in range(default__.abs(44))]), len(d_241_v74_), d_171_globalState_)
                d_242_v75_ = out0_
                d_244_v76_: C8
                nw35_ = C8()
                nw35_.ctor__()
                d_244_v76_ = nw35_
                d_244_v76_ = d_244_v76_
                d_245_v77_: _dafny.Seq
                d_245_v77_ = _dafny.SeqWithoutIsStrInference([d_234_v67_])
                d_246_v78_: _dafny.Map
                d_246_v78_ = _dafny.Map({d_245_v77_: True})
                d_247_v79_: _dafny.Array
                nw36_ = _dafny.Array(None, 15)
                nw36_[int(0)] = (d_246_v78_) | (d_246_v78_)
                nw36_[int(1)] = d_246_v78_
                nw36_[int(2)] = d_246_v78_
                nw36_[int(3)] = _dafny.Map({d_245_v77_: d_156_v0_})
                nw36_[int(4)] = d_246_v78_
                nw36_[int(5)] = d_246_v78_
                nw36_[int(6)] = d_246_v78_
                nw36_[int(7)] = (d_246_v78_) | (d_246_v78_)
                nw36_[int(8)] = d_246_v78_
                nw36_[int(9)] = d_246_v78_
                nw36_[int(10)] = d_246_v78_
                nw36_[int(11)] = (d_246_v78_).set(d_245_v77_, d_156_v0_)
                nw36_[int(12)] = d_246_v78_
                nw36_[int(13)] = d_246_v78_
                nw36_[int(14)] = d_246_v78_
                d_247_v79_ = nw36_
                index49_ = default__.safeIndex(701, (d_247_v79_).length(0))
                (d_247_v79_)[index49_] = d_246_v78_
                index50_ = default__.safeIndex(701, (d_247_v79_).length(0))
                (d_247_v79_)[index50_] = (d_246_v78_).set(d_245_v77_, (d_156_v0_ if d_156_v0_ else d_156_v0_))
                d_248_v80_: _dafny.Map
                d_248_v80_ = _dafny.Map({d_164_v8_: d_156_v0_})
                d_248_v80_ = d_248_v80_
                d_249_v81_: _dafny.Array
                nw37_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_249_v81_ = nw37_
                d_250_v82_: D5
                d_250_v82_ = D5_DC10(d_249_v81_)
                rhs35_ = (default__.safeDivisionInt(d_159_v3_, len(d_162_v6_))) - ((972) * (len(d_162_v6_)))
                rhs36_ = d_159_v3_
                rhs37_ = d_250_v82_
                lhs38_ = d_171_globalState_
                lhs39_ = d_171_globalState_
                lhs38_.f6 = rhs35_
                lhs39_.f6 = rhs36_
                d_250_v82_ = rhs37_
            elif True:
                (d_171_globalState_).f6 = d_159_v3_
                (d_171_globalState_).f1 = default__.fm2(d_171_globalState_)
                (d_171_globalState_).f1 = d_156_v0_
                index51_ = default__.safeIndex(252, (d_168_v10_).length(0))
                (d_168_v10_)[index51_] = d_156_v0_
                index52_ = default__.safeIndex(252, (d_168_v10_).length(0))
                (d_168_v10_)[index52_] = (default__.fm2(d_171_globalState_) if d_156_v0_ else (d_156_v0_ if d_156_v0_ else d_156_v0_))
                d_251_v83_: C4
                nw38_ = C4()
                nw38_.ctor__()
                d_251_v83_ = nw38_
                d_252_v84_: _dafny.Array
                nw39_ = _dafny.Array(D10.default()(), 17)
                d_252_v84_ = nw39_
                d_253_v85_: _dafny.Array
                nw40_ = _dafny.Array(None, 24)
                d_253_v85_ = nw40_
                d_254_v86_: _dafny.Seq
                d_254_v86_ = _dafny.SeqWithoutIsStrInference([d_253_v85_, d_253_v85_])
                d_255_v87_: _dafny.Map
                d_255_v87_ = _dafny.Map({(d_168_v10_)[default__.safeIndex(252, (d_168_v10_).length(0))]: default__.fm1((d_168_v10_)[default__.safeIndex(252, (d_168_v10_).length(0))], d_171_globalState_)})
                d_256_v88_: _dafny.Array
                nw41_ = _dafny.Array(None, 3)
                nw41_[int(0)] = (d_168_v10_)[default__.safeIndex(252, (d_168_v10_).length(0))]
                nw41_[int(1)] = d_156_v0_
                nw41_[int(2)] = False
                d_256_v88_ = nw41_
                d_257_v89_: D27
                d_257_v89_ = D27_DC61(d_254_v86_)
                index53_ = default__.safeIndex(252, (d_168_v10_).length(0))
                rhs38_ = ((d_239_v72_) + ((d_239_v72_).set(default__.safeIndex(((d_255_v87_)[d_156_v0_] if (d_156_v0_) in (d_255_v87_) else ((d_255_v87_)[d_156_v0_] if (d_156_v0_) in (d_255_v87_) else default__.fm1(False, d_171_globalState_))), len(d_239_v72_)), d_256_v88_))) != ((d_239_v72_) + (_dafny.SeqWithoutIsStrInference([d_256_v88_, d_256_v88_])))
                rhs39_ = d_251_v83_
                rhs40_ = d_252_v84_
                rhs41_ = d_256_v88_
                rhs42_ = (d_257_v89_).cf101
                lhs40_ = d_168_v10_
                lhs41_ = default__.safeIndex(252, (d_168_v10_).length(0))
                lhs42_ = d_171_globalState_
                lhs40_[lhs41_] = rhs38_
                d_251_v83_ = rhs39_
                d_252_v84_ = rhs40_
                lhs42_.f14 = rhs41_
                d_254_v86_ = rhs42_
        elif True:
            d_258_v90_: _dafny.Array
            nw42_ = _dafny.Array(D0.default()(), 19)
            d_258_v90_ = nw42_
            default__.m0(d_167_v9_, d_258_v90_, d_164_v8_, (d_162_v6_) + (d_162_v6_), d_171_globalState_)
            d_164_v8_ = (d_164_v8_ if True else d_164_v8_)
            (d_171_globalState_).f6 = (0) - (default__.safeModuloInt((0) - (d_159_v3_), d_159_v3_))
            d_259_v91_: _dafny.MultiSet
            d_259_v91_ = _dafny.MultiSet([d_172_v11_])
            d_259_v91_ = (d_259_v91_).intersection(d_259_v91_)
            d_260_v92_: D3
            d_260_v92_ = D3_DC6(d_164_v8_)
            d_261_v93_: _dafny.Seq
            d_261_v93_ = _dafny.SeqWithoutIsStrInference([d_260_v92_])
            d_262_v94_: C0
            nw43_ = C0()
            nw43_.ctor__(d_261_v93_, d_156_v0_)
            d_262_v94_ = nw43_
        (d_171_globalState_).f1 = default__.fm2(d_171_globalState_)
        d_263_v95_: _dafny.Map
        d_263_v95_ = _dafny.Map({True: d_159_v3_})
        d_264_v96_: _dafny.Seq
        d_264_v96_ = _dafny.SeqWithoutIsStrInference([d_263_v95_, d_263_v95_])
        d_265_v97_: _dafny.Array
        nw44_ = _dafny.Array(None, 10)
        nw44_[int(0)] = d_263_v95_
        nw44_[int(1)] = (d_263_v95_) | (d_263_v95_)
        nw44_[int(2)] = d_263_v95_
        nw44_[int(3)] = (d_263_v95_) | (d_263_v95_)
        nw44_[int(4)] = d_263_v95_
        nw44_[int(5)] = (d_264_v96_)[default__.safeIndex(d_159_v3_, len(d_264_v96_))]
        nw44_[int(6)] = _dafny.Map({d_156_v0_: d_159_v3_})
        nw44_[int(7)] = d_263_v95_
        nw44_[int(8)] = (_dafny.Map({True: d_159_v3_}) if not(d_156_v0_) else d_263_v95_)
        nw44_[int(9)] = d_263_v95_
        d_265_v97_ = nw44_
        index54_ = default__.safeIndex(31, (d_265_v97_).length(0))
        (d_265_v97_)[index54_] = (d_263_v95_) | (d_263_v95_)
        d_266_v98_: _dafny.Map
        d_266_v98_ = _dafny.Map({647: len(d_162_v6_)})
        d_267_v99_: D12
        d_267_v99_ = D12_DC34(d_266_v98_)
        index55_ = default__.safeIndex(31, (d_265_v97_).length(0))
        (d_265_v97_)[index55_] = default__.fm36(d_172_v11_, d_267_v99_, d_171_globalState_)
        d_268_v100_: _dafny.Array
        nw45_ = _dafny.Array(_dafny.Array(None, 0), 29)
        d_268_v100_ = nw45_
        source4_ = D5_DC10(d_268_v100_)
        if source4_.is_DC11:
            d_269___mcc_h0_ = source4_.cf14
            d_270_cf14_ = d_269___mcc_h0_
            d_271_v101_: _dafny.Seq
            d_271_v101_ = _dafny.SeqWithoutIsStrInference([default__.fm1(d_156_v0_, d_171_globalState_)])
            d_272_v102_: D0
            d_272_v102_ = D0_DC1(len(d_271_v101_))
            d_273_v103_: D0
            d_273_v103_ = D0_DC2(d_272_v102_)
            pat_let_tv3_ = d_272_v102_
            d_274_v104_: _dafny.Array
            nw46_ = _dafny.Array(None, 17)
            nw46_[int(0)] = D0_DC2(d_272_v102_)
            nw46_[int(1)] = d_273_v103_
            nw46_[int(2)] = d_273_v103_
            nw46_[int(3)] = d_273_v103_
            nw46_[int(4)] = d_273_v103_
            nw46_[int(5)] = d_273_v103_
            def iife37_(_pat_let6_0):
                def iife38_(d_275_dt__update__tmp_h1_):
                    def iife39_(_pat_let7_0):
                        def iife40_(d_276_dt__update_hcf2_h1_):
                            return D0_DC2(d_276_dt__update_hcf2_h1_)
                        return iife40_(_pat_let7_0)
                    return iife39_(pat_let_tv3_)
                return iife38_(_pat_let6_0)
            nw46_[int(6)] = iife37_(d_273_v103_)
            nw46_[int(7)] = default__.fm56(d_270_cf14_, d_172_v11_, d_171_globalState_)
            nw46_[int(8)] = d_273_v103_
            nw46_[int(9)] = d_273_v103_
            nw46_[int(10)] = d_273_v103_
            nw46_[int(11)] = d_273_v103_
            nw46_[int(12)] = d_273_v103_
            nw46_[int(13)] = D0_DC2(d_272_v102_)
            nw46_[int(14)] = d_273_v103_
            nw46_[int(15)] = D0_DC2(d_272_v102_)
            nw46_[int(16)] = D0_DC2(d_272_v102_)
            d_274_v104_ = nw46_
            default__.m0(d_167_v9_, d_274_v104_, (d_164_v8_ if d_156_v0_ else d_164_v8_), (d_162_v6_) + (d_162_v6_), d_171_globalState_)
            d_277_v105_: C8
            nw47_ = C8()
            nw47_.ctor__()
            d_277_v105_ = nw47_
            d_278_v106_: C5
            nw48_ = C5()
            nw48_.ctor__(d_270_cf14_)
            d_278_v106_ = nw48_
            (d_171_globalState_).f6 = d_159_v3_
        elif source4_.is_DC12:
            d_279___mcc_h1_ = source4_.cf15
            d_280___mcc_h2_ = source4_.cf16
            d_281___mcc_h3_ = source4_.cf17
            d_282_cf17_ = d_281___mcc_h3_
            d_283_cf16_ = d_280___mcc_h2_
            d_284_cf15_ = d_279___mcc_h1_
            d_285_v107_: _dafny.Map
            d_285_v107_ = _dafny.Map({(932) - (default__.fm1(d_156_v0_, d_171_globalState_)): d_282_cf17_})
            def iife41_():
                coll25_ = _dafny.Set()
                compr_25_: int
                for compr_25_ in _dafny.IntegerRange(-232, 590):
                    d_286_v108_: int = compr_25_
                    if ((-232) <= (d_286_v108_)) and ((d_286_v108_) < (590)):
                        coll25_ = coll25_.union(_dafny.Set([default__.safeDivisionInt(d_286_v108_, d_159_v3_)]))
                return _dafny.Set(coll25_)
            d_285_v107_ = (d_285_v107_).set(len(iife41_()
            ), d_172_v11_)
            (d_171_globalState_).f6 = d_159_v3_
            if True:
                index56_ = default__.safeIndex(974, (d_164_v8_).length(0))
                (d_164_v8_)[index56_] = d_284_cf15_
                d_287_v109_: _dafny.Set
                d_287_v109_ = _dafny.Set({d_283_cf16_, not(d_283_cf16_)})
                d_288_v110_: _dafny.Seq
                d_288_v110_ = _dafny.SeqWithoutIsStrInference([d_159_v3_, len(d_287_v109_)])
                d_289_v111_: _dafny.Seq
                d_289_v111_ = _dafny.SeqWithoutIsStrInference([d_288_v110_, d_288_v110_])
                index57_ = default__.safeIndex(974, (d_164_v8_).length(0))
                (d_164_v8_)[index57_] = len((d_289_v111_)[default__.safeIndex(d_159_v3_, len(d_289_v111_))])
                d_159_v3_ = (d_159_v3_) * (d_284_cf15_)
                index58_ = default__.safeIndex(974, (d_164_v8_).length(0))
                (d_164_v8_)[index58_] = d_284_cf15_
                d_290_v112_: _dafny.Map
                d_290_v112_ = _dafny.Map({((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdakhap"))) + (d_162_v6_)).set(default__.safeIndex(d_159_v3_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdakhap"))) + (d_162_v6_))), d_172_v11_): d_157_v1_})
                d_290_v112_ = (d_290_v112_).set(d_162_v6_, _dafny.MultiSet(d_288_v110_))
                d_291_v113_: D0
                d_291_v113_ = D0_DC1((d_164_v8_)[default__.safeIndex(974, (d_164_v8_).length(0))])
                d_292_v114_: D0
                d_292_v114_ = D0_DC2(d_291_v113_)
                pat_let_tv4_ = d_156_v0_
                pat_let_tv5_ = d_291_v113_
                d_293_v115_: _dafny.Array
                nw49_ = _dafny.Array(None, 14)
                nw49_[int(0)] = d_292_v114_
                nw49_[int(1)] = d_292_v114_
                nw49_[int(2)] = d_292_v114_
                nw49_[int(3)] = d_292_v114_
                nw49_[int(4)] = d_292_v114_
                nw49_[int(5)] = D0_DC2(D0_DC2(D0_DC2(d_291_v113_)))
                nw49_[int(6)] = d_292_v114_
                def iife42_(_pat_let8_0):
                    def iife43_(d_294_dt__update__tmp_h2_):
                        def iife44_(_pat_let9_0):
                            def iife45_(d_295_dt__update_hcf2_h2_):
                                return D0_DC2(d_295_dt__update_hcf2_h2_)
                            return iife45_(_pat_let9_0)
                        return iife44_(D0_DC0(pat_let_tv4_))
                    return iife43_(_pat_let8_0)
                nw49_[int(7)] = iife42_(d_292_v114_)
                nw49_[int(8)] = D0_DC2(d_291_v113_)
                nw49_[int(9)] = d_292_v114_
                nw49_[int(10)] = d_292_v114_
                nw49_[int(11)] = d_292_v114_
                def iife46_(_pat_let10_0):
                    def iife47_(d_296_dt__update__tmp_h3_):
                        def iife48_(_pat_let11_0):
                            def iife49_(d_297_dt__update_hcf2_h3_):
                                return D0_DC2(d_297_dt__update_hcf2_h3_)
                            return iife49_(_pat_let11_0)
                        return iife48_(pat_let_tv5_)
                    return iife47_(_pat_let10_0)
                nw49_[int(12)] = iife46_(d_292_v114_)
                nw49_[int(13)] = d_292_v114_
                d_293_v115_ = nw49_
                d_298_v116_: _dafny.Array
                nw50_ = _dafny.Array(int(0), 28)
                d_298_v116_ = nw50_
                default__.m0((d_167_v9_) | (d_167_v9_), d_293_v115_, d_298_v116_, d_162_v6_, d_171_globalState_)
            elif True:
                d_299_v117_: _dafny.Array
                def lambda7_(d_300_v0_):
                    def lambda8_(d_301_i6_):
                        return D0_DC2(D0_DC0(d_300_v0_))

                    return lambda8_

                init4_ = lambda7_(d_156_v0_)
                nw51_ = _dafny.Array(None, 27)
                for i0_4_ in range(nw51_.length(0)):
                    nw51_[i0_4_] = init4_(i0_4_)
                d_299_v117_ = nw51_
                default__.m0((d_167_v9_) | (d_167_v9_), d_299_v117_, d_164_v8_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")), d_171_globalState_)
                index59_ = default__.safeIndex(206, (d_164_v8_).length(0))
                (d_164_v8_)[index59_] = d_284_cf15_
                index60_ = default__.safeIndex(206, (d_164_v8_).length(0))
                (d_164_v8_)[index60_] = d_159_v3_
                d_283_cf16_ = not(d_156_v0_)
                d_302_v118_: _dafny.Seq
                d_302_v118_ = _dafny.SeqWithoutIsStrInference([-465, (d_164_v8_)[default__.safeIndex(206, (d_164_v8_).length(0))], 233])
                (d_171_globalState_).f1 = ((d_164_v8_)[default__.safeIndex(206, (d_164_v8_).length(0))]) in (d_302_v118_)
                index61_ = default__.safeIndex(649, (d_168_v10_).length(0))
                (d_168_v10_)[index61_] = d_156_v0_
                d_303_v119_: _dafny.Map
                d_303_v119_ = _dafny.Map({d_156_v0_: d_156_v0_})
                d_304_v120_: D20
                d_304_v120_ = D20_DC50(d_156_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "korfyx"))), d_284_cf15_, d_156_v0_, d_156_v0_)
                d_305_v121_: _dafny.Set
                d_305_v121_ = _dafny.Set({(d_304_v120_).cf90})
                index62_ = default__.safeIndex(649, (d_168_v10_).length(0))
                (d_168_v10_)[index62_] = (((d_302_v118_)[default__.safeIndex(d_159_v3_, len(d_302_v118_))]) > (d_159_v3_)) or ((len(d_303_v119_)) >= (len(d_305_v121_)))
            index63_ = default__.safeIndex(512, (d_164_v8_).length(0))
            (d_164_v8_)[index63_] = d_159_v3_
            index64_ = default__.safeIndex(512, (d_164_v8_).length(0))
            (d_164_v8_)[index64_] = d_284_cf15_
        elif source4_.is_DC10:
            d_306___mcc_h4_ = source4_.cf13
            d_307_cf13_ = d_306___mcc_h4_
            d_159_v3_ = d_159_v3_
            (d_171_globalState_).f1 = False
            d_308_v123_: _dafny.Array
            nw52_ = _dafny.Array(D0.default()(), 20)
            d_308_v123_ = nw52_
            def iife50_():
                coll26_ = _dafny.Set()
                compr_26_: int
                for compr_26_ in _dafny.IntegerRange(29, 167):
                    d_309_v122_: int = compr_26_
                    if ((29) <= (d_309_v122_)) and ((d_309_v122_) < (167)):
                        coll26_ = coll26_.union(_dafny.Set([(d_309_v122_) * (d_159_v3_)]))
                return _dafny.Set(coll26_)
            default__.m0(iife50_()
            , d_308_v123_, d_164_v8_, d_162_v6_, d_171_globalState_)
            d_310_v124_: _dafny.Map
            d_310_v124_ = _dafny.Map({d_156_v0_: d_172_v11_})
            d_311_v125_: _dafny.Map
            d_311_v125_ = d_310_v124_
            d_312_v126_: _dafny.MultiSet
            d_312_v126_ = _dafny.MultiSet([d_311_v125_, d_310_v124_])
            (d_171_globalState_).f6 = (((d_312_v126_ if True else d_312_v126_)) - (d_312_v126_)).cardinality
        elif True:
            d_313___mcc_h5_ = source4_.cf18
            d_314_cf18_ = d_313___mcc_h5_
            d_315_v127_: C4
            nw53_ = C4()
            nw53_.ctor__()
            d_315_v127_ = nw53_
            d_315_v127_ = d_315_v127_
            d_316_v128_: int
            d_317_v129_: bool
            out1_: int
            out2_: bool
            out1_, out2_ = (d_315_v127_).m1(d_171_globalState_)
            d_316_v128_ = out1_
            d_317_v129_ = out2_
            if True:
                d_159_v3_ = d_159_v3_
                (d_171_globalState_).f6 = d_316_v128_
                (d_171_globalState_).f6 = (0) - (default__.safeModuloInt(d_159_v3_, -943))
                d_318_v130_: _dafny.MultiSet
                d_318_v130_ = _dafny.MultiSet([d_156_v0_])
                d_319_v131_: _dafny.Map
                d_319_v131_ = _dafny.Map({(d_175_v14_)[default__.safeIndex(696, (d_175_v14_).length(0))]: 177})
                d_320_v132_: C2
                nw54_ = C2()
                nw54_.ctor__(d_159_v3_, (d_318_v130_).cardinality, d_319_v131_)
                d_320_v132_ = nw54_
                (d_171_globalState_).f1 = d_156_v0_
            elif True:
                d_321_v133_: _dafny.Array
                def lambda9_(d_322_v11_):
                    def lambda10_(d_323_i7_):
                        return d_322_v11_

                    return lambda10_

                init5_ = lambda9_(d_172_v11_)
                nw55_ = _dafny.Array(None, 20)
                for i0_5_ in range(nw55_.length(0)):
                    nw55_[i0_5_] = init5_(i0_5_)
                d_321_v133_ = nw55_
                d_321_v133_ = d_321_v133_
                d_324_v134_: C8
                nw56_ = C8()
                nw56_.ctor__()
                d_324_v134_ = nw56_
                index65_ = default__.safeIndex(250, (d_164_v8_).length(0))
                (d_164_v8_)[index65_] = default__.safeModuloInt(-570, -137)
                index66_ = default__.safeIndex(250, (d_164_v8_).length(0))
                (d_164_v8_)[index66_] = d_316_v128_
                d_317_v129_ = (d_316_v128_) > ((d_164_v8_)[default__.safeIndex(250, (d_164_v8_).length(0))])
                index67_ = default__.safeIndex(40, (d_168_v10_).length(0))
                (d_168_v10_)[index67_] = default__.fm2(d_171_globalState_)
                index68_ = default__.safeIndex(40, (d_168_v10_).length(0))
                (d_168_v10_)[index68_] = (not(not(True)) if d_156_v0_ else d_317_v129_)
            d_325_v135_: D5
            d_325_v135_ = D5_DC12(d_316_v128_, d_156_v0_, d_172_v11_)
            d_317_v129_ = (d_325_v135_).cf16
        d_326_v136_: _dafny.Seq
        d_326_v136_ = _dafny.SeqWithoutIsStrInference([d_159_v3_])
        d_327_v138_: D0
        d_327_v138_ = D0_DC0(d_156_v0_)
        d_328_v139_: D0
        d_328_v139_ = D0_DC2(d_327_v138_)
        d_329_v140_: D0
        d_329_v140_ = D0_DC2(d_327_v138_)
        d_330_v141_: _dafny.Array
        nw57_ = _dafny.Array(None, 1)
        nw57_[int(0)] = d_329_v140_
        d_330_v141_ = nw57_
        def iife51_():
            coll27_ = _dafny.Set()
            compr_27_: int
            for compr_27_ in (d_326_v136_).Elements:
                d_331_v137_: int = compr_27_
                if (d_331_v137_) in (d_326_v136_):
                    coll27_ = coll27_.union(_dafny.Set([default__.safeModuloInt(d_331_v137_, 634)]))
            return _dafny.Set(coll27_)
        default__.m0(iife51_()
        , d_330_v141_, d_164_v8_, (d_162_v6_) + (d_162_v6_), d_171_globalState_)
        if default__.fm2(d_171_globalState_):
            (d_171_globalState_).f6 = 286
            d_332_v142_: _dafny.Map
            d_332_v142_ = _dafny.Map({d_156_v0_: d_156_v0_})
            d_333_v143_: _dafny.Map
            d_333_v143_ = _dafny.Map({d_160_v4_: d_159_v3_})
            d_334_v144_: T1
            nw58_ = C3()
            nw58_.ctor__(((D14_DC39(d_326_v136_, d_332_v142_, False)).cf68) < (d_326_v136_), len(d_332_v142_), (d_333_v143_) | (d_333_v143_))
            d_334_v144_ = nw58_
            d_335_v145_: _dafny.Seq
            d_335_v145_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_173_v12_)])
            rhs43_ = d_334_v144_
            rhs44_ = True
            rhs45_ = (d_156_v0_) == (d_156_v0_)
            rhs46_ = d_156_v0_
            rhs47_ = ((((d_162_v6_) + (d_162_v6_)) + (d_162_v6_)).set(default__.safeIndex(len((d_162_v6_).set(default__.safeIndex(len(d_167_v9_), len(d_162_v6_)), d_172_v11_)), len(((d_162_v6_) + (d_162_v6_)) + (d_162_v6_))), d_172_v11_)).set(default__.safeIndex((0) - (((d_335_v145_)[default__.safeIndex(d_159_v3_, len(d_335_v145_))]).cardinality), len((((d_162_v6_) + (d_162_v6_)) + (d_162_v6_)).set(default__.safeIndex(len((d_162_v6_).set(default__.safeIndex(len(d_167_v9_), len(d_162_v6_)), d_172_v11_)), len(((d_162_v6_) + (d_162_v6_)) + (d_162_v6_))), d_172_v11_))), d_172_v11_)
            lhs43_ = d_171_globalState_
            d_334_v144_ = rhs43_
            d_156_v0_ = rhs44_
            d_156_v0_ = rhs45_
            lhs43_.f1 = rhs46_
            d_162_v6_ = rhs47_
            d_336_v146_: _dafny.MultiSet
            d_336_v146_ = _dafny.MultiSet([d_156_v0_, d_156_v0_])
            d_337_v147_: C2
            nw59_ = C2()
            nw59_.ctor__(d_159_v3_, (d_336_v146_).cardinality, _dafny.Map({(d_175_v14_)[default__.safeIndex(696, (d_175_v14_).length(0))]: d_159_v3_}))
            d_337_v147_ = nw59_
            (d_171_globalState_).f6 = ((d_336_v146_)[d_156_v0_] if (d_156_v0_) in (d_336_v146_) else d_337_v147_.f21)
            if True:
                d_338_v148_: _dafny.Map
                d_338_v148_ = _dafny.Map({d_156_v0_: d_162_v6_})
                d_338_v148_ = (d_338_v148_).set(d_156_v0_, d_162_v6_)
                d_335_v145_ = d_335_v145_
                (d_171_globalState_).f1 = (d_173_v12_)[default__.safeIndex(d_337_v147_.f21, len(d_173_v12_))]
                d_164_v8_ = d_164_v8_
                d_339_v149_: str
                out3_: str
                out3_ = (d_334_v144_).m11(d_162_v6_, (d_337_v147_).f20, d_171_globalState_)
                d_339_v149_ = out3_
            elif True:
                (d_171_globalState_).f6 = (d_159_v3_) * ((d_337_v147_).f20)
                d_340_v150_: _dafny.Map
                d_340_v150_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ep")): d_334_v144_})
                d_341_v151_: _dafny.Map
                d_341_v151_ = _dafny.Map({(0) - (d_337_v147_.f21): d_334_v144_})
                d_340_v150_ = (d_340_v150_).set(d_162_v6_, ((d_341_v151_)[(d_337_v147_).f20] if ((d_337_v147_).f20) in (d_341_v151_) else d_334_v144_))
                d_342_v152_: D0
                d_342_v152_ = D0_DC0(d_156_v0_)
                (d_171_globalState_).f1 = (d_337_v147_).fm16(d_342_v152_, True, d_171_globalState_)
                d_343_v153_: bool
                d_344_v154_: _dafny.Seq
                out4_: bool
                out5_: _dafny.Seq
                out4_, out5_ = (d_337_v147_).m12((d_337_v147_).f20, d_337_v147_.f21, (d_162_v6_) + (d_162_v6_), (d_163_v7_)[default__.safeIndex(d_337_v147_.f21, len(d_163_v7_))], d_171_globalState_)
                d_343_v153_ = out4_
                d_344_v154_ = out5_
                d_345_v155_: _dafny.Set
                d_345_v155_ = _dafny.Set({d_343_v153_, default__.fm2(d_171_globalState_)})
                d_346_v156_: D8
                d_346_v156_ = D8_DC21(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahhjndpbh")), d_345_v155_)
                d_347_v157_: str
                out6_: str
                out6_ = (d_337_v147_).m11((((d_346_v156_).cf34).set(default__.safeIndex((d_337_v147_).f20, len((d_346_v156_).cf34)), d_172_v11_)) + (d_344_v154_), d_337_v147_.f21, d_171_globalState_)
                d_347_v157_ = out6_
        elif True:
            d_348_v158_: _dafny.Set
            d_348_v158_ = _dafny.Set({(d_162_v6_)[default__.safeIndex(d_159_v3_, len(d_162_v6_))]})
            (d_171_globalState_).f6 = (0) - (((d_159_v3_) * (d_159_v3_)) * ((d_159_v3_ if d_156_v0_ else len(d_348_v158_))))
            d_159_v3_ = d_159_v3_
            d_349_v159_: _dafny.Map
            d_349_v159_ = _dafny.Map({d_156_v0_: d_156_v0_})
            d_350_v160_: _dafny.Map
            d_350_v160_ = _dafny.Map({len(d_349_v159_): d_157_v1_})
            if (((d_350_v160_)[d_159_v3_] if (d_159_v3_) in (d_350_v160_) else default__.fm42(default__.fm1(not(d_156_v0_), d_171_globalState_), d_172_v11_, d_159_v3_, True, d_171_globalState_))).issubset((d_157_v1_) | (d_157_v1_)):
                d_156_v0_ = d_156_v0_
                d_351_v161_: C8
                nw60_ = C8()
                nw60_.ctor__()
                d_351_v161_ = nw60_
                d_352_v162_: _dafny.Map
                d_352_v162_ = _dafny.Map({d_351_v161_: 928})
                d_353_v163_: C9
                nw61_ = C9()
                nw61_.ctor__(len(d_162_v6_))
                d_353_v163_ = nw61_
                d_354_v164_: _dafny.Seq
                d_354_v164_ = _dafny.SeqWithoutIsStrInference([d_353_v163_, d_353_v163_])
                d_266_v98_ = (d_266_v98_).set(((d_352_v162_)[d_351_v161_] if (d_351_v161_) in (d_352_v162_) else len((_dafny.Map({len(d_354_v164_): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "glsxubw")))})).set(d_159_v3_, (d_353_v163_).f15))), (0) - (len(d_162_v6_)))
                d_355_v165_: C4
                nw62_ = C4()
                nw62_.ctor__()
                d_355_v165_ = nw62_
                d_356_v166_: C5
                nw63_ = C5()
                nw63_.ctor__(d_159_v3_)
                d_356_v166_ = nw63_
                d_357_v167_: _dafny.Map
                d_357_v167_ = _dafny.Map({d_159_v3_: d_356_v166_})
                d_358_v168_: T1
                nw64_ = C2()
                nw64_.ctor__(len(d_357_v167_), (d_353_v163_).f15, _dafny.Map({(d_175_v14_)[default__.safeIndex(696, (d_175_v14_).length(0))]: (d_356_v166_).f16}))
                d_358_v168_ = nw64_
                d_359_v169_: _dafny.Array
                nw65_ = _dafny.Array(_dafny.Map({}), 22)
                d_359_v169_ = nw65_
                d_360_v170_: D20
                d_360_v170_ = D20_DC49(d_359_v169_)
                d_361_v171_: _dafny.Map
                d_361_v171_ = _dafny.Map({d_358_v168_: d_360_v170_})
                d_362_v172_: _dafny.Map
                d_362_v172_ = _dafny.Map({d_361_v171_: default__.fm2(d_171_globalState_)})
                d_363_v173_: C5
                nw66_ = C5()
                nw66_.ctor__((0) - (default__.safeDivisionInt(d_159_v3_, len(d_362_v172_))))
                d_363_v173_ = nw66_
                d_364_v174_: bool
                d_365_v175_: _dafny.Seq
                out7_: bool
                out8_: _dafny.Seq
                out7_, out8_ = (d_358_v168_).m12(len((d_162_v6_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "swd")))), d_159_v3_, (d_162_v6_) + (d_162_v6_), ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yeigb"))).set(default__.safeIndex((d_353_v163_).f15, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yeigb")))), d_172_v11_)) + (d_162_v6_), d_171_globalState_)
                d_364_v174_ = out7_
                d_365_v175_ = out8_
            elif True:
                d_366_v176_: D14
                d_366_v176_ = D14_DC40(not(d_156_v0_))
                pat_let_tv6_ = d_156_v0_
                pat_let_tv7_ = d_156_v0_
                d_367_v177_: _dafny.Array
                nw67_ = _dafny.Array(None, 26)
                nw67_[int(0)] = d_366_v176_
                nw67_[int(1)] = d_366_v176_
                nw67_[int(2)] = d_366_v176_
                nw67_[int(3)] = d_366_v176_
                nw67_[int(4)] = d_366_v176_
                nw67_[int(5)] = default__.fm57(d_163_v7_, d_173_v12_, d_171_globalState_)
                nw67_[int(6)] = d_366_v176_
                nw67_[int(7)] = D14_DC40(d_156_v0_)
                nw67_[int(8)] = d_366_v176_
                nw67_[int(9)] = d_366_v176_
                nw67_[int(10)] = d_366_v176_
                nw67_[int(11)] = d_366_v176_
                nw67_[int(12)] = d_366_v176_
                nw67_[int(13)] = d_366_v176_
                nw67_[int(14)] = d_366_v176_
                nw67_[int(15)] = d_366_v176_
                nw67_[int(16)] = D14_DC40(d_156_v0_)
                nw67_[int(17)] = D14_DC40(d_156_v0_)
                nw67_[int(18)] = d_366_v176_
                nw67_[int(19)] = D14_DC40(d_156_v0_)
                nw67_[int(20)] = d_366_v176_
                def iife53_(_pat_let13_0):
                    def iife54_(d_368_dt__update__tmp_h4_):
                        def iife55_(_pat_let14_0):
                            def iife56_(d_369_dt__update_hcf71_h0_):
                                return D14_DC40(d_369_dt__update_hcf71_h0_)
                            return iife56_(_pat_let14_0)
                        return iife55_(pat_let_tv6_)
                    return iife54_(_pat_let13_0)
                def iife52_(_pat_let12_0):
                    def iife57_(d_370_dt__update__tmp_h5_):
                        def iife58_(_pat_let15_0):
                            def iife59_(d_371_dt__update_hcf71_h1_):
                                return D14_DC40(d_371_dt__update_hcf71_h1_)
                            return iife59_(_pat_let15_0)
                        return iife58_(pat_let_tv7_)
                    return iife57_(_pat_let12_0)
                nw67_[int(21)] = iife52_(iife53_(d_366_v176_))
                nw67_[int(22)] = D14_DC40(True)
                nw67_[int(23)] = d_366_v176_
                nw67_[int(24)] = d_366_v176_
                nw67_[int(25)] = d_366_v176_
                d_367_v177_ = nw67_
                index69_ = default__.safeIndex(799, (d_367_v177_).length(0))
                (d_367_v177_)[index69_] = d_366_v176_
                index70_ = default__.safeIndex(799, (d_367_v177_).length(0))
                (d_367_v177_)[index70_] = D14_DC40(d_156_v0_)
                d_372_v178_: _dafny.Array
                def lambda11_(d_373_v136_):
                    def lambda12_(d_374_i8_):
                        return d_373_v136_

                    return lambda12_

                init6_ = lambda11_(d_326_v136_)
                nw68_ = _dafny.Array(None, 5)
                for i0_6_ in range(nw68_.length(0)):
                    nw68_[i0_6_] = init6_(i0_6_)
                d_372_v178_ = nw68_
                d_375_v179_: _dafny.Seq
                d_375_v179_ = d_326_v136_
                d_376_v180_: _dafny.Set
                d_376_v180_ = _dafny.Set({d_266_v98_, d_266_v98_})
                index71_ = default__.safeIndex(434, (d_372_v178_).length(0))
                (d_372_v178_)[index71_] = (d_326_v136_) + (((d_375_v179_)).set(default__.safeIndex(len(d_376_v180_), len((d_375_v179_))), d_159_v3_))
                d_377_v181_: _dafny.Map
                d_377_v181_ = _dafny.Map({d_160_v4_: (0) - (d_159_v3_)})
                d_378_v182_: T1
                nw69_ = C3()
                nw69_.ctor__(d_156_v0_, default__.fm1(d_156_v0_, d_171_globalState_), d_377_v181_)
                d_378_v182_ = nw69_
                index72_ = default__.safeIndex(434, (d_372_v178_).length(0))
                rhs48_ = d_159_v3_
                rhs49_ = d_172_v11_
                rhs50_ = (_dafny.SeqWithoutIsStrInference([len(default__.fm28(d_159_v3_, True, d_156_v0_, d_171_globalState_)), d_159_v3_, d_159_v3_])) + (d_326_v136_)
                rhs51_ = d_378_v182_
                lhs44_ = d_372_v178_
                lhs45_ = default__.safeIndex(434, (d_372_v178_).length(0))
                d_159_v3_ = rhs48_
                d_172_v11_ = rhs49_
                lhs44_[lhs45_] = rhs50_
                d_378_v182_ = rhs51_
                d_379_v183_: _dafny.Map
                d_379_v183_ = _dafny.Map({d_159_v3_: d_378_v182_})
                d_380_v184_: _dafny.Map
                d_380_v184_ = _dafny.Map({d_156_v0_: d_379_v183_})
                (d_171_globalState_).f6 = ((d_372_v178_)[default__.safeIndex(434, (d_372_v178_).length(0))])[default__.safeIndex(len((d_380_v184_) | (d_380_v184_)), len((d_372_v178_)[default__.safeIndex(434, (d_372_v178_).length(0))]))]
                d_381_v185_: _dafny.Map
                d_381_v185_ = d_377_v181_
                d_382_v186_: C3
                nw70_ = C3()
                nw70_.ctor__(d_156_v0_, d_159_v3_, (d_381_v185_))
                d_382_v186_ = nw70_
                d_383_v187_: _dafny.Map
                d_383_v187_ = _dafny.Map({(d_382_v186_).f19: d_382_v186_})
                d_384_v188_: _dafny.Array
                nw71_ = _dafny.Array(None, 27)
                nw71_[int(0)] = d_382_v186_
                nw71_[int(1)] = d_382_v186_
                nw71_[int(2)] = d_382_v186_
                nw71_[int(3)] = d_382_v186_
                nw71_[int(4)] = d_382_v186_
                nw71_[int(5)] = d_382_v186_
                nw71_[int(6)] = d_382_v186_
                nw71_[int(7)] = d_382_v186_
                nw71_[int(8)] = d_382_v186_
                nw71_[int(9)] = d_382_v186_
                nw71_[int(10)] = d_382_v186_
                nw71_[int(11)] = d_382_v186_
                nw71_[int(12)] = d_382_v186_
                nw71_[int(13)] = d_382_v186_
                nw71_[int(14)] = d_382_v186_
                nw71_[int(15)] = d_382_v186_
                nw71_[int(16)] = d_382_v186_
                nw71_[int(17)] = d_382_v186_
                nw71_[int(18)] = d_382_v186_
                nw71_[int(19)] = d_382_v186_
                nw71_[int(20)] = d_382_v186_
                nw71_[int(21)] = d_382_v186_
                nw71_[int(22)] = d_382_v186_
                nw71_[int(23)] = d_382_v186_
                nw71_[int(24)] = d_382_v186_
                nw71_[int(25)] = d_382_v186_
                nw71_[int(26)] = ((d_383_v187_)[(d_382_v186_).f19] if ((d_382_v186_).f19) in (d_383_v187_) else d_382_v186_)
                d_384_v188_ = nw71_
                index73_ = default__.safeIndex(289, (d_384_v188_).length(0))
                (d_384_v188_)[index73_] = d_382_v186_
                index74_ = default__.safeIndex(289, (d_384_v188_).length(0))
                (d_384_v188_)[index74_] = d_382_v186_
                index75_ = default__.safeIndex(617, (d_174_v13_).length(0))
                (d_174_v13_)[index75_] = d_164_v8_
                d_385_v189_: _dafny.Map
                d_385_v189_ = _dafny.Map({d_266_v98_: (d_382_v186_).f19})
                index76_ = default__.safeIndex(617, (d_174_v13_).length(0))
                nw72_ = _dafny.Array(None, 20)
                nw72_[int(0)] = d_159_v3_
                nw72_[int(1)] = ((d_382_v186_).f19) + ((d_382_v186_).f19)
                nw72_[int(2)] = default__.safeModuloInt(len(d_173_v12_), len(_dafny.SeqWithoutIsStrInference([(d_382_v186_).f19])))
                nw72_[int(3)] = d_159_v3_
                nw72_[int(4)] = (d_159_v3_ if not((d_382_v186_).f18) else 916)
                nw72_[int(5)] = d_159_v3_
                nw72_[int(6)] = d_159_v3_
                nw72_[int(7)] = d_159_v3_
                nw72_[int(8)] = 106
                nw72_[int(9)] = default__.safeDivisionInt((d_382_v186_).f19, len(d_162_v6_))
                nw72_[int(10)] = d_159_v3_
                nw72_[int(11)] = -541
                nw72_[int(12)] = len(d_162_v6_)
                nw72_[int(13)] = -581
                nw72_[int(14)] = (0) - ((0) - ((d_382_v186_).f19))
                nw72_[int(15)] = d_159_v3_
                nw72_[int(16)] = default__.fm1((d_382_v186_).f18, d_171_globalState_)
                nw72_[int(17)] = (942) * (d_159_v3_)
                nw72_[int(18)] = len(d_385_v189_)
                nw72_[int(19)] = default__.safeDivisionInt(-559, d_159_v3_)
                (d_174_v13_)[index76_] = nw72_
            d_160_v4_ = (d_160_v4_).set(d_159_v3_, d_156_v0_)
            rhs52_ = d_164_v8_
            rhs53_ = default__.fm2(d_171_globalState_)
            rhs54_ = (d_326_v136_)[default__.safeIndex(d_159_v3_, len(d_326_v136_))]
            lhs46_ = d_171_globalState_
            lhs47_ = d_171_globalState_
            lhs48_ = d_171_globalState_
            lhs46_.f9 = rhs52_
            lhs47_.f1 = rhs53_
            lhs48_.f6 = rhs54_
        (d_171_globalState_).f6 = (d_159_v3_) - (168)
        _dafny.print(_dafny.string_of(d_156_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v1_) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v2_) == (_dafny.Map({_dafny.MultiSet([1]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_159_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v4_) == (_dafny.Map({781: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_v5_) == (_dafny.MultiSet([_dafny.Map({781: True}), _dafny.Map({781: True}), _dafny.Map({781: True}), _dafny.Map({781: True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_162_v6_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_163_v7_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kajndjn"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v8_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v8_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v8_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v8_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_167_v9_) == (_dafny.Set({-781, 781}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v10_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v10_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v10_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v10_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v10_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v10_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v10_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v10_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v10_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_171_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_171_globalState_).f4) == (_dafny.Map({_dafny.MultiSet([1]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_171_globalState_).f5) == (_dafny.MultiSet([_dafny.Map({781: True}), _dafny.Map({781: True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_171_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_171_globalState_).f8) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kajndjn")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kajndjn"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f9)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f9)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f9)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f9)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f9)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f9)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f9)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f9)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f9)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f9)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f9)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_171_globalState_).f11) == (_dafny.Set({-781, 781}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f14)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f14)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f14)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f14)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f14)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f14)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f14)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f14)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_globalState_.f14)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_172_v11_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_173_v12_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_174_v13_)[0])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_174_v13_)[0])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_174_v13_)[0])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_174_v13_)[0])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v14_)[0]) == (_dafny.Map({781: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v14_)[1]) == (_dafny.Map({781: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v14_)[2]) == (_dafny.Map({781: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v14_)[3]) == (_dafny.Map({781: True, 0: True, 9: False, 898: True, 6: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v14_)[4]) == (_dafny.Map({781: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v14_)[5]) == (_dafny.Map({781: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v14_)[6]) == (_dafny.Map({781: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v14_)[7]) == (_dafny.Map({781: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v14_)[8]) == (_dafny.Map({781: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_183_v19_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v95_) == (_dafny.Map({True: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v96_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: 0}), _dafny.Map({True: 0})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v97_)[0]) == (_dafny.Map({True: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v97_)[1]) == (_dafny.Map({True: 1, False: -685}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v97_)[2]) == (_dafny.Map({True: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v97_)[3]) == (_dafny.Map({True: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v97_)[4]) == (_dafny.Map({True: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v97_)[5]) == (_dafny.Map({True: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v97_)[6]) == (_dafny.Map({True: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v97_)[7]) == (_dafny.Map({True: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v97_)[8]) == (_dafny.Map({True: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v97_)[9]) == (_dafny.Map({True: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_266_v98_) == (_dafny.Map({647: 7}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_267_v99_).cf63) == (_dafny.Map({647: 7}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_326_v136_) == (_dafny.SeqWithoutIsStrInference([0]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v138_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_328_v139_).cf2).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_329_v140_).cf2).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_330_v141_)[0]).cf2).cf0))
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
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)

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

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC5(D2, NamedTuple('DC5', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC7(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)

class D3_DC7(D3, NamedTuple('DC7', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf8)}, {self.cf9.VerbatimString(True)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC6(D3, NamedTuple('DC6', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)

class D4_DC9(D4, NamedTuple('DC9', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC11(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D5_DC10)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)

class D5_DC11(D5, NamedTuple('DC11', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC10(D5, NamedTuple('DC10', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC10({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC10) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)

class D6_DC14(D6, NamedTuple('DC14', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC16(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)

class D7_DC16(D7, NamedTuple('DC16', [('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC15(D7, NamedTuple('DC15', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC20(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)

class D8_DC20(D8, NamedTuple('DC20', [('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({self.cf34.VerbatimString(True)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC19(D8, NamedTuple('DC19', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC23(_dafny.Set({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, _dafny.Array(None, 0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)

class D9_DC23(D9, NamedTuple('DC23', [('cf37', Any), ('cf38', Any), ('cf39', Any), ('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf37)}, {self.cf38.VerbatimString(True)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf42', Any), ('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC22(D9, NamedTuple('DC22', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC27(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), D3.default()(), _dafny.Map({}))
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

class D10_DC27(D10, NamedTuple('DC27', [('cf46', Any), ('cf47', Any), ('cf48', Any), ('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf46)}, {self.cf47.VerbatimString(True)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC26(D10, NamedTuple('DC26', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC29(D10, NamedTuple('DC29', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC31(_dafny.Seq({}), int(0), int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)

class D11_DC31(D11, NamedTuple('DC31', [('cf54', Any), ('cf55', Any), ('cf56', Any), ('cf57', Any), ('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC32(D11, NamedTuple('DC32', [('cf59', Any), ('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC33(D11, NamedTuple('DC33', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC35()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D12_DC35)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)

class D12_DC35(D12, NamedTuple('DC35', [])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35)
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC37(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D13_DC37)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D13_DC36)

class D13_DC37(D13, NamedTuple('DC37', [('cf65', Any), ('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC37({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC37) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC36(D13, NamedTuple('DC36', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC36({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC36) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC39(_dafny.Seq({}), _dafny.Map({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D14_DC39)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D14_DC40)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D14_DC38)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D14_DC41)

class D14_DC39(D14, NamedTuple('DC39', [('cf68', Any), ('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC39({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC39) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC40(D14, NamedTuple('DC40', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC40({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC40) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC38(D14, NamedTuple('DC38', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC38({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC38) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC41(D14, NamedTuple('DC41', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC41({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC41) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC43(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D15_DC42)

class D15_DC43(D15, NamedTuple('DC43', [('cf74', Any), ('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43({_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43) and self.cf74 == __o.cf74 and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC42(D15, NamedTuple('DC42', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC42({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC42) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC45(int(0), _dafny.CodePoint('D'), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D16_DC45)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D16_DC44)

class D16_DC45(D16, NamedTuple('DC45', [('cf78', Any), ('cf79', Any), ('cf80', Any), ('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC45({_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC45) and self.cf78 == __o.cf78 and self.cf79 == __o.cf79 and self.cf80 == __o.cf80 and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC44(D16, NamedTuple('DC44', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC44({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC44) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D17_DC46)

class D17_DC46(D17, NamedTuple('DC46', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC46({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC46) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D18_DC47)

class D18_DC47(D18, NamedTuple('DC47', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC47({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC47) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D19_DC48)

class D19_DC48(D19, NamedTuple('DC48', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC48({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC48) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC50(False, int(0), int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D20_DC50)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D20_DC49)

class D20_DC50(D20, NamedTuple('DC50', [('cf86', Any), ('cf87', Any), ('cf88', Any), ('cf89', Any), ('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC50({_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC50) and self.cf86 == __o.cf86 and self.cf87 == __o.cf87 and self.cf88 == __o.cf88 and self.cf89 == __o.cf89 and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC49(D20, NamedTuple('DC49', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC49({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC49) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D21_DC51)

class D21_DC51(D21, NamedTuple('DC51', [('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC51({_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC51) and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D22_DC52)

class D22_DC52(D22, NamedTuple('DC52', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC52({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC52) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC54()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D23_DC54)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D23_DC55)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D23_DC53)

class D23_DC54(D23, NamedTuple('DC54', [])):
    def __dafnystr__(self) -> str:
        return f'D23.DC54'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC54)
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC55(D23, NamedTuple('DC55', [('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC55({_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC55) and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC53(D23, NamedTuple('DC53', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC53({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC53) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D24_DC56)

class D24_DC56(D24, NamedTuple('DC56', [('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC56({_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC56) and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC58(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D25_DC58)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D25_DC57)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D25_DC59)

class D25_DC58(D25, NamedTuple('DC58', [('cf97', Any), ('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC58({_dafny.string_of(self.cf97)}, {_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC58) and self.cf97 == __o.cf97 and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC57(D25, NamedTuple('DC57', [('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC57({_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC57) and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC59(D25, NamedTuple('DC59', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC59({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC59) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D26_DC60)

class D26_DC60(D26, NamedTuple('DC60', [('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC60({_dafny.string_of(self.cf100)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC60) and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC62(False, False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D27_DC62)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D27_DC63)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D27_DC61)

class D27_DC62(D27, NamedTuple('DC62', [('cf102', Any), ('cf103', Any), ('cf104', Any), ('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC62({_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)}, {_dafny.string_of(self.cf104)}, {_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC62) and self.cf102 == __o.cf102 and self.cf103 == __o.cf103 and self.cf104 == __o.cf104 and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC63(D27, NamedTuple('DC63', [('cf106', Any), ('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC63({_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC63) and self.cf106 == __o.cf106 and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC61(D27, NamedTuple('DC61', [('cf101', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC61({_dafny.string_of(self.cf101)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC61) and self.cf101 == __o.cf101
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, globalState):
        pass

    def m2(self, p0, globalState):
        pass


class T1:
    pass
    @property
    def f17(self):
        return self._f17
    @f17.setter
    def f17(self, value):
        self._f17 = value
    def m11(self, p0, p1, globalState):
        pass

    def m12(self, p0, p1, p2, p3, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f1: bool = False
        self.f6: int = int(0)
        self.f9: _dafny.Array = _dafny.Array(None, 0)
        self.f14: _dafny.Array = _dafny.Array(None, 0)
        self._f0: int = int(0)
        self._f2: bool = False
        self._f3: int = int(0)
        self._f4: _dafny.Map = _dafny.Map({})
        self._f5: _dafny.MultiSet = _dafny.MultiSet({})
        self._f7: bool = False
        self._f8: _dafny.Seq = _dafny.Seq({})
        self._f10: bool = False
        self._f11: _dafny.Set = _dafny.Set({})
        self._f12: bool = False
        self._f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14):
        (self)._f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f14 = f14

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
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
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

class C0:
    def  __init__(self):
        self._f22: _dafny.Seq = _dafny.Seq({})
        self._f23: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f22, f23):
        (self)._f22 = f22
        (self)._f23 = f23

    def fm17(self, p0, p1, p2, globalState):
        return (self).f23

    @property
    def f22(self):
        return self._f22
    @property
    def f23(self):
        return self._f23

class C1(T0):
    def  __init__(self):
        self._f24: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f24):
        (self)._f24 = f24

    def fm3(self, p0, p1, p2, p3, globalState):
        source5_ = D1_DC4(_dafny.CodePoint('a'))
        if source5_.is_DC4:
            d_386___mcc_h0_ = source5_.cf4
            d_387_cf4_ = d_386___mcc_h0_
            return D0_DC2(D0_DC2(D0_DC0(False)))
        elif True:
            d_388___mcc_h1_ = source5_.cf3
            d_389_cf3_ = d_388___mcc_h1_
            return D0_DC2(D0_DC2(D0_DC1((self).f24)))

    def fm24(self, p0, p1, globalState):
        def iife60_():
            coll28_ = _dafny.Set()
            compr_28_: int
            for compr_28_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(self).f24: _dafny.MultiSet([True, True])}))])).Elements:
                d_390_v0_: int = compr_28_
                if (d_390_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(self).f24: _dafny.MultiSet([True, True])}))])):
                    coll28_ = coll28_.union(_dafny.Set([default__.safeModuloInt(d_390_v0_, len(_dafny.SeqWithoutIsStrInference([(0) - (-199)])))]))
            return _dafny.Set(coll28_)
        return not(((383) < ((self).f24)) == (not(((self).f24) == (len(iife60_()
        )))))

    def fm25(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(self).f24 for d_391_i0_ in range(default__.abs(920))])) + (_dafny.SeqWithoutIsStrInference([(self).f24]))) + (_dafny.SeqWithoutIsStrInference([(0) - ((self).f24), (self).f24]))

    def m1(self, globalState):
        r0: int = int(0)
        r1: bool = False
        hi2_ = (self).f24
        for d_392_i0_ in range((0) - ((self).f24), hi2_):
            d_393_v0_: _dafny.Array
            nw73_ = _dafny.Array(int(0), 13)
            d_393_v0_ = nw73_
            index77_ = default__.safeIndex(50, (d_393_v0_).length(0))
            (d_393_v0_)[index77_] = ((self).f24) * ((self).f24)
            d_394_v1_: bool
            d_394_v1_ = True
            d_395_v2_: _dafny.Set
            d_395_v2_ = _dafny.Set({d_394_v1_})
            d_396_v3_: _dafny.Set
            d_396_v3_ = _dafny.Set({d_395_v2_, d_395_v2_})
            d_397_v4_: _dafny.Map
            d_397_v4_ = _dafny.Map({d_394_v1_: (self).f24})
            d_398_v5_: _dafny.Seq
            d_398_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kffytsi"))
            d_399_v6_: _dafny.Map
            d_399_v6_ = _dafny.Map({d_397_v4_: d_398_v5_})
            index78_ = default__.safeIndex(50, (d_393_v0_).length(0))
            (d_393_v0_)[index78_] = (0) - (default__.safeModuloInt(len(d_396_v3_), default__.safeModuloInt(len(d_399_v6_), (self).f24)))
            d_400_v7_: str
            d_400_v7_ = _dafny.CodePoint('p')
            d_400_v7_ = d_400_v7_
            d_401_v8_: _dafny.Array
            def lambda13_(d_402_v1_):
                def lambda14_(d_403_i1_):
                    return D1_DC3(_dafny.Map({d_402_v1_: d_402_v1_}))

                return lambda14_

            init7_ = lambda13_(d_394_v1_)
            nw74_ = _dafny.Array(None, 25)
            for i0_7_ in range(nw74_.length(0)):
                nw74_[i0_7_] = init7_(i0_7_)
            d_401_v8_ = nw74_
            d_404_v9_: _dafny.Map
            d_404_v9_ = _dafny.Map({d_394_v1_: d_394_v1_})
            index79_ = default__.safeIndex(863, (d_401_v8_).length(0))
            (d_401_v8_)[index79_] = D1_DC3(d_404_v9_)
            d_405_v10_: D1
            d_405_v10_ = D1_DC3(d_404_v9_)
            index80_ = default__.safeIndex(863, (d_401_v8_).length(0))
            rhs55_ = d_405_v10_
            rhs56_ = not((((d_398_v5_) + (d_398_v5_)).set(default__.safeIndex(-886, len((d_398_v5_) + (d_398_v5_))), d_400_v7_)) <= (d_398_v5_))
            lhs49_ = d_401_v8_
            lhs50_ = default__.safeIndex(863, (d_401_v8_).length(0))
            lhs49_[lhs50_] = rhs55_
            r1 = rhs56_
            r1 = d_394_v1_
        d_406_v11_: bool
        d_406_v11_ = False
        d_407_v12_: _dafny.Set
        d_407_v12_ = _dafny.Set({d_406_v11_})
        r1 = (d_407_v12_).issubset(d_407_v12_)
        d_408_v13_: _dafny.Array
        def lambda15_(d_409_i2_):
            return default__.safeModuloInt(d_409_i2_, (self).f24)

        init8_ = lambda15_
        nw75_ = _dafny.Array(None, 3)
        for i0_8_ in range(nw75_.length(0)):
            nw75_[i0_8_] = init8_(i0_8_)
        d_408_v13_ = nw75_
        d_410_v14_: _dafny.Seq
        d_410_v14_ = _dafny.SeqWithoutIsStrInference([d_408_v13_, d_408_v13_, d_408_v13_, d_408_v13_])
        d_411_v15_: _dafny.Seq
        d_411_v15_ = _dafny.SeqWithoutIsStrInference([D3_DC6((d_410_v14_)[default__.safeIndex((0) - ((self).f24), len(d_410_v14_))])])
        d_412_v16_: C0
        nw76_ = C0()
        nw76_.ctor__(d_411_v15_, d_406_v11_)
        d_412_v16_ = nw76_
        d_413_v17_: _dafny.Seq
        d_413_v17_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b')])
        d_414_v18_: _dafny.Array
        def lambda16_(d_415_i3_):
            return False

        init9_ = lambda16_
        nw77_ = _dafny.Array(None, 8)
        for i0_9_ in range(nw77_.length(0)):
            nw77_[i0_9_] = init9_(i0_9_)
        d_414_v18_ = nw77_
        d_416_v19_: _dafny.Map
        d_416_v19_ = _dafny.Map({len(d_413_v17_): d_414_v18_})
        d_417_v20_: _dafny.Map
        d_417_v20_ = _dafny.Map({((d_416_v19_)[(self).f24] if ((self).f24) in (d_416_v19_) else d_414_v18_): (d_412_v16_).f23})
        d_418_v21_: _dafny.Map
        d_418_v21_ = _dafny.Map({True: d_417_v20_})
        d_418_v21_ = (d_418_v21_).set((d_412_v16_).f23, (d_417_v20_ if d_406_v11_ else d_417_v20_))
        index81_ = default__.safeIndex(596, (d_408_v13_).length(0))
        (d_408_v13_)[index81_] = (self).f24
        d_419_v22_: _dafny.MultiSet
        d_419_v22_ = _dafny.MultiSet([not((d_412_v16_).f23)])
        index82_ = default__.safeIndex(596, (d_408_v13_).length(0))
        (d_408_v13_)[index82_] = (192) + ((d_419_v22_).cardinality)
        d_420_i4_: int
        d_420_i4_ = 0
        with _dafny.label("1"):
            while (d_413_v17_) < (d_413_v17_):
                with _dafny.c_label("1"):
                    if (d_420_i4_) >= (100):
                        raise _dafny.Break("1")
                    d_420_i4_ = (d_420_i4_) + (1)
                    d_421_v23_: D0
                    d_421_v23_ = D0_DC0(not(d_406_v11_))
                    d_422_v24_: _dafny.Seq
                    d_422_v24_ = _dafny.SeqWithoutIsStrInference([d_406_v11_, True])
                    d_423_v25_: _dafny.Map
                    d_423_v25_ = _dafny.Map({d_421_v23_: d_422_v24_})
                    d_423_v25_ = (d_423_v25_).set(d_421_v23_, d_422_v24_)
                    d_424_v26_: _dafny.Map
                    d_424_v26_ = _dafny.Map({(d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))]: True})
                    d_425_v27_: _dafny.Map
                    d_425_v27_ = _dafny.Map({not((d_412_v16_).f23): (self).f24})
                    d_424_v26_ = (d_424_v26_).set(((d_425_v27_)[(d_412_v16_).f23] if ((d_412_v16_).f23) in (d_425_v27_) else (self).f24), d_406_v11_)
                    d_426_v28_: str
                    d_426_v28_ = _dafny.CodePoint('s')
                    index83_ = default__.safeIndex(321, (d_414_v18_).length(0))
                    (d_414_v18_)[index83_] = (d_412_v16_).f23
                    index84_ = default__.safeIndex(321, (d_414_v18_).length(0))
                    rhs57_ = (d_426_v28_ if (d_412_v16_).f23 else d_426_v28_)
                    rhs58_ = False
                    rhs59_ = (0) - ((d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))])
                    lhs51_ = d_414_v18_
                    lhs52_ = default__.safeIndex(321, (d_414_v18_).length(0))
                    d_426_v28_ = rhs57_
                    lhs51_[lhs52_] = rhs58_
                    r0 = rhs59_
                    d_427_v29_: _dafny.Set
                    d_427_v29_ = _dafny.Set({(self).f24, (d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))]})
                    d_428_v30_: _dafny.Array
                    nw78_ = _dafny.Array(None, 10)
                    nw78_[int(0)] = d_406_v11_
                    nw78_[int(1)] = d_406_v11_
                    nw78_[int(2)] = (d_412_v16_).f23
                    nw78_[int(3)] = d_406_v11_
                    nw78_[int(4)] = not((d_412_v16_).f23)
                    nw78_[int(5)] = (d_414_v18_)[default__.safeIndex(321, (d_414_v18_).length(0))]
                    nw78_[int(6)] = d_406_v11_
                    nw78_[int(7)] = (d_412_v16_).f23
                    nw78_[int(8)] = (d_414_v18_)[default__.safeIndex(321, (d_414_v18_).length(0))]
                    nw78_[int(9)] = (d_414_v18_)[default__.safeIndex(321, (d_414_v18_).length(0))]
                    d_428_v30_ = nw78_
                    d_429_v31_: _dafny.Set
                    d_429_v31_ = _dafny.Set({d_428_v30_, d_428_v30_})
                    d_430_v32_: _dafny.Map
                    d_430_v32_ = _dafny.Map({len(d_429_v31_): d_427_v29_})
                    d_431_v33_: D0
                    d_431_v33_ = D0_DC0(d_406_v11_)
                    d_432_v34_: D0
                    d_432_v34_ = D0_DC2(d_431_v33_)
                    d_433_v35_: _dafny.Array
                    nw79_ = _dafny.Array(None, 12)
                    nw79_[int(0)] = d_413_v17_
                    nw79_[int(1)] = d_413_v17_
                    nw79_[int(2)] = d_413_v17_
                    nw79_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocslm"))
                    nw79_[int(4)] = _dafny.SeqWithoutIsStrInference([d_426_v28_ for d_434_i5_ in range(default__.abs(43))])
                    nw79_[int(5)] = d_413_v17_
                    nw79_[int(6)] = d_413_v17_
                    nw79_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmbkie"))
                    nw79_[int(8)] = d_413_v17_
                    nw79_[int(9)] = d_413_v17_
                    nw79_[int(10)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_435_i6_ in range(default__.abs(-616))])
                    nw79_[int(11)] = d_413_v17_
                    d_433_v35_ = nw79_
                    d_436_v36_: _dafny.Map
                    d_436_v36_ = _dafny.Map({d_433_v35_: True})
                    d_437_v38_: _dafny.Map
                    d_437_v38_ = _dafny.Map({d_406_v11_: d_430_v32_})
                    d_438_v41_: _dafny.Seq
                    d_438_v41_ = _dafny.SeqWithoutIsStrInference([len(d_422_v24_)])
                    d_439_v42_: _dafny.MultiSet
                    d_439_v42_ = _dafny.MultiSet([(d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))], (d_438_v41_)[default__.safeIndex((d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))], len(d_438_v41_))], (d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))], len(d_438_v41_)])
                    d_440_v43_: _dafny.Map
                    d_440_v43_ = _dafny.Map({(d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))]: (d_439_v42_).cardinality})
                    d_441_v44_: _dafny.Array
                    nw80_ = _dafny.Array(None, 18)
                    nw80_[int(0)] = _dafny.Map({(0) - ((self).f24): d_427_v29_})
                    nw80_[int(1)] = d_430_v32_
                    nw80_[int(2)] = (default__.fm26(d_426_v28_, d_432_v34_, ((d_436_v36_)[d_433_v35_] if (d_433_v35_) in (d_436_v36_) else d_406_v11_), (d_412_v16_).f23, globalState))
                    nw80_[int(3)] = (_dafny.Map({(self).f24: d_427_v29_})) | (d_430_v32_)
                    def iife61_():
                        coll29_ = _dafny.Map()
                        compr_29_: int
                        for compr_29_ in _dafny.IntegerRange(229, 9):
                            d_442_v37_: int = compr_29_
                            if ((229) <= (d_442_v37_)) and ((d_442_v37_) < (9)):
                                coll29_[(d_442_v37_) * ((d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))])] = _dafny.Set({511, (self).f24, (self).f24})
                        return _dafny.Map(coll29_)
                    nw80_[int(4)] = (iife61_()
                    ) | (d_430_v32_)
                    nw80_[int(5)] = d_430_v32_
                    nw80_[int(6)] = (d_430_v32_) | (d_430_v32_)
                    nw80_[int(7)] = (d_430_v32_) | (d_430_v32_)
                    nw80_[int(8)] = (d_430_v32_) | (d_430_v32_)
                    nw80_[int(9)] = (_dafny.Map({(self).f24: _dafny.Set({(d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))], (d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))]})})) | (default__.fm27((self).f24, (d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))], (self).f24, (d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))], globalState))
                    nw80_[int(10)] = d_430_v32_
                    nw80_[int(11)] = _dafny.Map({(d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))]: d_427_v29_})
                    nw80_[int(12)] = d_430_v32_
                    def iife62_():
                        coll30_ = _dafny.Set()
                        compr_30_: int
                        for compr_30_ in _dafny.IntegerRange(84, 808):
                            d_443_v39_: int = compr_30_
                            if ((84) <= (d_443_v39_)) and ((d_443_v39_) < (808)):
                                coll30_ = coll30_.union(_dafny.Set([(d_443_v39_) * ((d_419_v22_).cardinality)]))
                        return _dafny.Set(coll30_)
                    nw80_[int(13)] = ((((d_437_v38_)[(d_414_v18_)[default__.safeIndex(321, (d_414_v18_).length(0))]] if ((d_414_v18_)[default__.safeIndex(321, (d_414_v18_).length(0))]) in (d_437_v38_) else _dafny.Map({(self).f24: iife62_()
                    }))).set((self).f24, d_427_v29_) if (d_412_v16_).f23 else d_430_v32_)
                    nw80_[int(14)] = d_430_v32_
                    nw80_[int(15)] = d_430_v32_
                    def iife63_():
                        coll31_ = _dafny.Map()
                        compr_31_: int
                        for compr_31_ in (d_440_v43_).keys.Elements:
                            d_444_v40_: int = compr_31_
                            if (d_444_v40_) in (d_440_v43_):
                                coll31_[(d_444_v40_) + (816)] = d_427_v29_
                        return _dafny.Map(coll31_)
                    nw80_[int(16)] = (iife63_()
                    ) | (d_430_v32_)
                    nw80_[int(17)] = (d_430_v32_) | ((d_430_v32_).set(len(d_413_v17_), _dafny.Set({(self).f24})))
                    d_441_v44_ = nw80_
                    index85_ = default__.safeIndex(166, (d_441_v44_).length(0))
                    (d_441_v44_)[index85_] = _dafny.Map({(d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))]: d_427_v29_})
                    d_445_v45_: _dafny.Array
                    def lambda17_(d_446_v28_):
                        def lambda18_(d_447_i7_):
                            return d_446_v28_

                        return lambda18_

                    init10_ = lambda17_(d_426_v28_)
                    nw81_ = _dafny.Array(None, 29)
                    for i0_10_ in range(nw81_.length(0)):
                        nw81_[i0_10_] = init10_(i0_10_)
                    d_445_v45_ = nw81_
                    d_448_v46_: _dafny.Map
                    d_448_v46_ = _dafny.Map({d_445_v45_: d_430_v32_})
                    d_449_v47_: _dafny.Seq
                    d_449_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(0) - (len(d_422_v24_)): d_427_v29_})])
                    index86_ = default__.safeIndex(166, (d_441_v44_).length(0))
                    (d_441_v44_)[index86_] = ((d_448_v46_)[d_445_v45_] if (d_445_v45_) in (d_448_v46_) else (d_430_v32_ if (d_412_v16_).f23 else (d_449_v47_)[default__.safeIndex((d_408_v13_)[default__.safeIndex(596, (d_408_v13_).length(0))], len(d_449_v47_))]))
                    pass
            pass
        r0 = ((0) - ((self).f24)) * ((self).f24)
        r1 = not((d_410_v14_) != (d_410_v14_))
        return r0, r1

    def m2(self, p0, globalState):
        d_450_v0_: _dafny.Seq
        d_450_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
        d_451_v1_: _dafny.Seq
        d_451_v1_ = _dafny.SeqWithoutIsStrInference([d_450_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wkapm")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_452_i0_ in range(default__.abs(925))])])
        d_453_v2_: _dafny.Seq
        d_453_v2_ = _dafny.SeqWithoutIsStrInference([d_451_v1_])
        d_451_v1_ = (d_453_v2_)[default__.safeIndex(((self).f24) + ((self).f24), len(d_453_v2_))]
        d_454_v3_: _dafny.Seq
        d_454_v3_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        hi3_ = default__.fm1(p0, globalState)
        for d_455_i1_ in range(((self).f24 if (d_454_v3_)[default__.safeIndex((self).f24, len(d_454_v3_))] else 323), hi3_):
            (globalState).f1 = p0
            d_456_v4_: D5
            d_456_v4_ = D5_DC11(d_455_i1_)
            d_457_v5_: D5
            d_457_v5_ = D5_DC13(d_456_v4_)
            d_458_v6_: _dafny.Map
            d_458_v6_ = _dafny.Map({d_455_i1_: p0})
            d_459_v7_: _dafny.Map
            d_459_v7_ = _dafny.Map({d_458_v6_: (self).f24})
            rhs60_ = (self).f24
            rhs61_ = len((d_450_v0_) + (d_450_v0_))
            rhs62_ = d_457_v5_
            rhs63_ = (0) - (((d_459_v7_)[_dafny.Map({(self).f24: not(False)})] if (_dafny.Map({(self).f24: not(False)})) in (d_459_v7_) else (self).f24))
            rhs64_ = default__.fm1((p0) or (p0), globalState)
            lhs53_ = globalState
            lhs54_ = globalState
            lhs55_ = globalState
            lhs56_ = globalState
            lhs53_.f6 = rhs60_
            lhs54_.f6 = rhs61_
            d_457_v5_ = rhs62_
            lhs55_.f6 = rhs63_
            lhs56_.f6 = rhs64_
            (globalState).f1 = not(True)
            d_460_v8_: _dafny.Array
            def lambda19_(d_461_i2_):
                return default__.safeDivisionInt(d_461_i2_, -187)

            init11_ = lambda19_
            nw82_ = _dafny.Array(None, 19)
            for i0_11_ in range(nw82_.length(0)):
                nw82_[i0_11_] = init11_(i0_11_)
            d_460_v8_ = nw82_
            d_462_v9_: _dafny.Map
            d_462_v9_ = _dafny.Map({p0: (_dafny.MultiSet([len(_dafny.Map({d_455_i1_: d_460_v8_}))])).cardinality})
            d_463_v10_: _dafny.Map
            d_463_v10_ = _dafny.Map({d_462_v9_: (self).f24})
            (globalState).f6 = ((self).f24) - (((d_463_v10_)[d_462_v9_] if (d_462_v9_) in (d_463_v10_) else (self).f24))
        d_464_v11_: _dafny.Array
        def lambda20_(d_465_i3_):
            return (d_465_i3_) + ((self).f24)

        init12_ = lambda20_
        nw83_ = _dafny.Array(None, 18)
        for i0_12_ in range(nw83_.length(0)):
            nw83_[i0_12_] = init12_(i0_12_)
        d_464_v11_ = nw83_
        d_466_v12_: D3
        d_466_v12_ = D3_DC6(d_464_v11_)
        d_467_v13_: _dafny.Seq
        d_467_v13_ = _dafny.SeqWithoutIsStrInference([D3_DC6(d_464_v11_), d_466_v12_])
        d_468_v14_: C0
        nw84_ = C0()
        nw84_.ctor__(d_467_v13_, p0)
        d_468_v14_ = nw84_
        d_469_v15_: _dafny.Map
        d_469_v15_ = _dafny.Map({d_468_v14_: p0})
        (globalState).f6 = (default__.safeModuloInt(397, len(d_469_v15_))) + ((self).f24)
        d_470_i4_: int
        d_470_i4_ = 0
        with _dafny.label("2"):
            while (d_468_v14_).f23:
                with _dafny.c_label("2"):
                    if (d_470_i4_) >= (100):
                        raise _dafny.Break("2")
                    d_470_i4_ = (d_470_i4_) + (1)
                    d_471_v16_: str
                    d_471_v16_ = _dafny.CodePoint('y')
                    d_472_v17_: _dafny.Map
                    d_472_v17_ = _dafny.Map({d_471_v16_: (self).f24})
                    d_472_v17_ = (d_472_v17_).set(d_471_v16_, ((self).f24) * ((self).f24))
                    (globalState).f1 = (_dafny.Set({(d_468_v14_).f23, (d_468_v14_).f23})).isdisjoint(_dafny.Set({True, p0}))
                    d_471_v16_ = d_471_v16_
                    d_473_v18_: _dafny.Map
                    d_473_v18_ = _dafny.Map({p0: (self).f24})
                    d_474_v19_: _dafny.Map
                    d_474_v19_ = _dafny.Map({(self).f24: d_473_v18_})
                    d_475_v20_: _dafny.Seq
                    d_475_v20_ = _dafny.SeqWithoutIsStrInference([(self).f24, (self).f24, len(d_474_v19_), (self).f24, (self).f24])
                    d_476_v21_: _dafny.Map
                    d_476_v21_ = _dafny.Map({p0: _dafny.Set({p0, (d_468_v14_).f23, (d_468_v14_).f23, (d_468_v14_).f23, (d_468_v14_).f23})})
                    d_477_v22_: _dafny.Map
                    d_477_v22_ = _dafny.Map({(self).f24: (self).f24})
                    d_478_v23_: _dafny.Set
                    d_478_v23_ = _dafny.Set({True})
                    (globalState).f6 = default__.safeModuloInt((d_475_v20_)[default__.safeIndex(len((d_454_v3_).set(default__.safeIndex(len(((d_476_v21_)[(d_468_v14_).fm17(864, len(_dafny.SeqWithoutIsStrInference([len(d_477_v22_)])), (d_468_v14_).f23, globalState)] if ((d_468_v14_).fm17(864, len(_dafny.SeqWithoutIsStrInference([len(d_477_v22_)])), (d_468_v14_).f23, globalState)) in (d_476_v21_) else d_478_v23_)), len(d_454_v3_)), (d_468_v14_).f23)), len(d_475_v20_))], default__.safeDivisionInt((0) - ((self).f24), (self).f24))
                    pass
            pass
        (globalState).f6 = (0) - ((self).f24)
        d_450_v0_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_479_i5_ in range(default__.abs(403))])

    @property
    def f24(self):
        return self._f24

class C2(T1):
    def  __init__(self):
        self._f17: _dafny.Map = _dafny.Map({})
        self.f21: int = int(0)
        self._f20: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f17(self):
        return self._f17
    @f17.setter
    def f17(self, value):
        self._f17 = value
    def ctor__(self, f20, f21, f17):
        (self)._f20 = f20
        (self).f21 = f21
        (self).f17 = f17

    def fm15(self, globalState):
        return (0) - ((0) - ((len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_480_i0_ in range(default__.abs(-329))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_481_i1_ in range(default__.abs(542))])))) - (len((_dafny.Set({D3_DC7(len(_dafny.SeqWithoutIsStrInference([False, False])))})) - (_dafny.Set({D3_DC7(self.f21)}))))))

    def fm16(self, p0, p1, globalState):
        return (((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([559, len(_dafny.Map({self.f21: not(True)})), 111, self.f21, len(_dafny.Set({self.f21}))]))) - (_dafny.MultiSet([len(_dafny.Map({False: (self).f20})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkfiqofmd")))]))).cardinality) <= (self.f21)

    def m11(self, p0, p1, globalState):
        r0: str = _dafny.CodePoint('D')
        d_482_v0_: bool
        d_482_v0_ = True
        (globalState).f1 = d_482_v0_
        d_483_v1_: D0
        d_483_v1_ = D0_DC0(d_482_v0_)
        d_484_v2_: D3
        d_484_v2_ = D3_DC7((self).f20)
        hi4_ = (0) - ((d_484_v2_).cf7)
        for d_485_i0_ in range(len(_dafny.Map({(self).fm16(d_483_v1_, d_482_v0_, globalState): (d_484_v2_).cf7})), hi4_):
            d_486_v3_: _dafny.Array
            nw85_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_486_v3_ = nw85_
            d_487_v4_: _dafny.Array
            nw86_ = _dafny.Array(_dafny.Array(None, 0), 16)
            d_487_v4_ = nw86_
            d_488_v5_: D5
            d_488_v5_ = D5_DC10(d_487_v4_)
            index87_ = default__.safeIndex(810, (d_486_v3_).length(0))
            (d_486_v3_)[index87_] = (d_488_v5_).cf13
            index88_ = default__.safeIndex(810, (d_486_v3_).length(0))
            (d_486_v3_)[index88_] = d_487_v4_
            d_489_v6_: str
            d_489_v6_ = _dafny.CodePoint('v')
            d_490_v7_: _dafny.Map
            d_490_v7_ = _dafny.Map({len(_dafny.Map({d_489_v6_: d_485_i0_})): 811})
            (globalState).f6 = default__.safeModuloInt(len((d_490_v7_) | (d_490_v7_)), 279)
            d_491_v8_: _dafny.Seq
            d_491_v8_ = _dafny.SeqWithoutIsStrInference([p1, self.f21])
            d_492_v9_: _dafny.Array
            nw87_ = _dafny.Array(None, 27)
            nw87_[int(0)] = p1
            nw87_[int(1)] = len(d_490_v7_)
            nw87_[int(2)] = 659
            nw87_[int(3)] = (self).f20
            nw87_[int(4)] = d_485_i0_
            nw87_[int(5)] = len(_dafny.SeqWithoutIsStrInference([d_489_v6_ for d_493_i1_ in range(default__.abs(568))]))
            nw87_[int(6)] = len(default__.fm0(d_482_v0_, d_482_v0_, d_485_i0_, globalState))
            nw87_[int(7)] = p1
            nw87_[int(8)] = (self).f20
            nw87_[int(9)] = (self).f20
            nw87_[int(10)] = (self).f20
            nw87_[int(11)] = 527
            nw87_[int(12)] = p1
            nw87_[int(13)] = len(_dafny.SeqWithoutIsStrInference([d_485_i0_ for d_494_i2_ in range(default__.abs(755))]))
            nw87_[int(14)] = len((d_491_v8_).set(default__.safeIndex((0) - (d_485_i0_), len(d_491_v8_)), (self).f20))
            nw87_[int(15)] = (self).f20
            nw87_[int(16)] = (self).f20
            nw87_[int(17)] = (self).f20
            nw87_[int(18)] = self.f21
            nw87_[int(19)] = self.f21
            nw87_[int(20)] = 876
            nw87_[int(21)] = 693
            nw87_[int(22)] = len((p0).set(default__.safeIndex(d_485_i0_, len(p0)), d_489_v6_))
            nw87_[int(23)] = p1
            nw87_[int(24)] = 279
            nw87_[int(25)] = p1
            nw87_[int(26)] = -252
            d_492_v9_ = nw87_
            d_495_v10_: D3
            d_495_v10_ = D3_DC6(d_492_v9_)
            d_496_v11_: _dafny.Seq
            d_496_v11_ = _dafny.SeqWithoutIsStrInference([d_495_v10_])
            d_497_v12_: C0
            nw88_ = C0()
            nw88_.ctor__((d_496_v11_) + (d_496_v11_), d_482_v0_)
            d_497_v12_ = nw88_
            d_498_v13_: _dafny.Array
            def lambda21_(d_499_i0_):
                def lambda22_(d_500_i3_):
                    return (_dafny.MultiSet([d_499_i0_, self.f21])) | (_dafny.MultiSet([(self).f20]))

                return lambda22_

            init13_ = lambda21_(d_485_i0_)
            nw89_ = _dafny.Array(None, 25)
            for i0_13_ in range(nw89_.length(0)):
                nw89_[i0_13_] = init13_(i0_13_)
            d_498_v13_ = nw89_
            d_501_v14_: _dafny.MultiSet
            d_501_v14_ = _dafny.MultiSet([p1])
            index89_ = default__.safeIndex(64, (d_498_v13_).length(0))
            (d_498_v13_)[index89_] = d_501_v14_
            index90_ = default__.safeIndex(64, (d_498_v13_).length(0))
            (d_498_v13_)[index90_] = d_501_v14_
        d_502_v15_: _dafny.Seq
        d_502_v15_ = _dafny.SeqWithoutIsStrInference([(self).f20])
        d_503_v16_: _dafny.Array
        def lambda23_(d_504_v0_):
            def lambda24_(d_505_i4_):
                return d_504_v0_

            return lambda24_

        init14_ = lambda23_(d_482_v0_)
        nw90_ = _dafny.Array(None, 15)
        for i0_14_ in range(nw90_.length(0)):
            nw90_[i0_14_] = init14_(i0_14_)
        d_503_v16_ = nw90_
        d_506_v17_: _dafny.Seq
        d_506_v17_ = _dafny.SeqWithoutIsStrInference([d_482_v0_])
        d_507_v18_: _dafny.Map
        d_507_v18_ = _dafny.Map({_dafny.MultiSet(d_506_v17_): (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))))})
        d_508_v19_: _dafny.Array
        nw91_ = _dafny.Array(None, 14)
        nw91_[int(0)] = self.f21
        nw91_[int(1)] = (d_502_v15_)[default__.safeIndex(self.f21, len(d_502_v15_))]
        nw91_[int(2)] = p1
        nw91_[int(3)] = p1
        nw91_[int(4)] = self.f21
        nw91_[int(5)] = p1
        nw91_[int(6)] = len(_dafny.Set({d_503_v16_, d_503_v16_}))
        nw91_[int(7)] = p1
        nw91_[int(8)] = p1
        nw91_[int(9)] = self.f21
        nw91_[int(10)] = self.f21
        nw91_[int(11)] = len(d_507_v18_)
        nw91_[int(12)] = p1
        nw91_[int(13)] = (self).f20
        d_508_v19_ = nw91_
        d_509_v20_: D3
        d_509_v20_ = D3_DC6(d_508_v19_)
        source6_ = d_509_v20_
        if source6_.is_DC7:
            d_510___mcc_h0_ = source6_.cf7
            d_511_cf7_ = d_510___mcc_h0_
            index91_ = default__.safeIndex(429, (d_503_v16_).length(0))
            (d_503_v16_)[index91_] = d_482_v0_
            index92_ = default__.safeIndex(429, (d_503_v16_).length(0))
            (d_503_v16_)[index92_] = d_482_v0_
            pat_let_tv8_ = d_508_v19_
            d_512_v21_: C0
            nw92_ = C0()
            def iife64_(_pat_let16_0):
                def iife65_(d_513_dt__update__tmp_h0_):
                    def iife66_(_pat_let17_0):
                        def iife67_(d_514_dt__update_hcf6_h0_):
                            return D3_DC6(d_514_dt__update_hcf6_h0_)
                        return iife67_(_pat_let17_0)
                    return iife66_(pat_let_tv8_)
                return iife65_(_pat_let16_0)
            nw92_.ctor__(_dafny.SeqWithoutIsStrInference([iife64_(d_509_v20_), d_509_v20_, d_509_v20_, d_509_v20_]), not(d_482_v0_))
            d_512_v21_ = nw92_
            d_515_v22_: _dafny.Array
            def lambda25_(d_516_v21_):
                def lambda26_(d_517_i5_):
                    return D0_DC0((d_516_v21_).f23)

                return lambda26_

            init15_ = lambda25_(d_512_v21_)
            nw93_ = _dafny.Array(None, 2)
            for i0_15_ in range(nw93_.length(0)):
                nw93_[i0_15_] = init15_(i0_15_)
            d_515_v22_ = nw93_
            d_518_v23_: _dafny.Map
            d_518_v23_ = _dafny.Map({(d_503_v16_)[default__.safeIndex(429, (d_503_v16_).length(0))]: d_482_v0_})
            d_519_v24_: _dafny.Set
            d_519_v24_ = _dafny.Set({(d_512_v21_).f23})
            d_520_v25_: _dafny.Set
            d_520_v25_ = _dafny.Set({len(d_519_v24_)})
            d_521_v26_: str
            d_521_v26_ = _dafny.CodePoint('y')
            d_522_v27_: D3
            d_522_v27_ = D3_DC8(d_502_v15_, p0, (d_503_v16_)[default__.safeIndex(429, (d_503_v16_).length(0))], d_521_v26_)
            pat_let_tv9_ = d_503_v16_
            pat_let_tv10_ = d_503_v16_
            pat_let_tv11_ = d_482_v0_
            nw94_ = _dafny.Array(None, 24)
            nw94_[int(0)] = d_483_v1_
            nw94_[int(1)] = d_483_v1_
            nw94_[int(2)] = d_483_v1_
            def iife68_(_pat_let18_0):
                def iife69_(d_523_dt__update__tmp_h1_):
                    def iife70_(_pat_let19_0):
                        def iife71_(d_524_dt__update_hcf0_h0_):
                            return D0_DC0(d_524_dt__update_hcf0_h0_)
                        return iife71_(_pat_let19_0)
                    return iife70_((pat_let_tv10_)[default__.safeIndex(429, (pat_let_tv9_).length(0))])
                return iife69_(_pat_let18_0)
            nw94_[int(3)] = iife68_(d_483_v1_)
            nw94_[int(4)] = d_483_v1_
            nw94_[int(5)] = d_483_v1_
            def iife72_(_pat_let20_0):
                def iife73_(d_525_dt__update__tmp_h2_):
                    def iife74_(_pat_let21_0):
                        def iife75_(d_526_dt__update_hcf0_h1_):
                            return D0_DC0(d_526_dt__update_hcf0_h1_)
                        return iife75_(_pat_let21_0)
                    return iife74_(pat_let_tv11_)
                return iife73_(_pat_let20_0)
            nw94_[int(6)] = iife72_(d_483_v1_)
            nw94_[int(7)] = d_483_v1_
            nw94_[int(8)] = d_483_v1_
            nw94_[int(9)] = default__.fm18((self).f20, d_502_v15_, globalState)
            nw94_[int(10)] = default__.fm18(d_511_cf7_, d_502_v15_, globalState)
            nw94_[int(11)] = (d_483_v1_ if (d_512_v21_).fm17(len(_dafny.SeqWithoutIsStrInference([(d_503_v16_)[default__.safeIndex(429, (d_503_v16_).length(0))]])), p1, not((d_512_v21_).f23), globalState) else d_483_v1_)
            nw94_[int(12)] = default__.fm18(len(d_518_v23_), d_502_v15_, globalState)
            nw94_[int(13)] = d_483_v1_
            nw94_[int(14)] = d_483_v1_
            nw94_[int(15)] = D0_DC0(d_482_v0_)
            nw94_[int(16)] = d_483_v1_
            nw94_[int(17)] = d_483_v1_
            nw94_[int(18)] = d_483_v1_
            nw94_[int(19)] = default__.fm18((0) - (self.f21), d_502_v15_, globalState)
            nw94_[int(20)] = d_483_v1_
            nw94_[int(21)] = default__.fm18(len(d_520_v25_), (d_522_v27_).cf8, globalState)
            nw94_[int(22)] = default__.fm18(self.f21, d_502_v15_, globalState)
            nw94_[int(23)] = d_483_v1_
            d_515_v22_ = nw94_
            d_527_v28_: _dafny.MultiSet
            d_527_v28_ = _dafny.MultiSet([(d_512_v21_).f23, d_482_v0_])
            def iife76_():
                coll32_ = _dafny.Map()
                compr_32_: int
                for compr_32_ in (_dafny.Map({self.f21: 132})).keys.Elements:
                    d_528_v29_: int = compr_32_
                    if (d_528_v29_) in (_dafny.Map({self.f21: 132})):
                        coll32_[default__.safeDivisionInt(d_528_v29_, p1)] = _dafny.Map({(d_512_v21_).f23: self.f21})
                return _dafny.Map(coll32_)
            rhs65_ = (_dafny.MultiSet(d_506_v17_)).isdisjoint(((_dafny.MultiSet([d_482_v0_, (d_503_v16_)[default__.safeIndex(429, (d_503_v16_).length(0))], (d_512_v21_).f23, (d_512_v21_).f23])).set((d_512_v21_).f23, default__.abs(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npqmwqi")))))) | (d_527_v28_))
            rhs66_ = 115
            rhs67_ = not (not(d_482_v0_)) or (((d_503_v16_)[default__.safeIndex(429, (d_503_v16_).length(0))] if d_482_v0_ else d_482_v0_))
            rhs68_ = False
            rhs69_ = (0) - ((len(iife76_()
            )) - ((self).f20))
            lhs57_ = globalState
            lhs58_ = globalState
            lhs59_ = globalState
            lhs60_ = globalState
            lhs61_ = globalState
            lhs57_.f1 = rhs65_
            lhs58_.f6 = rhs66_
            lhs59_.f1 = rhs67_
            lhs60_.f1 = rhs68_
            lhs61_.f6 = rhs69_
        elif source6_.is_DC8:
            d_529___mcc_h1_ = source6_.cf8
            d_530___mcc_h2_ = source6_.cf9
            d_531___mcc_h3_ = source6_.cf10
            d_532___mcc_h4_ = source6_.cf11
            d_533_cf11_ = d_532___mcc_h4_
            d_534_cf10_ = d_531___mcc_h3_
            d_535_cf9_ = d_530___mcc_h2_
            d_536_cf8_ = d_529___mcc_h1_
            index93_ = default__.safeIndex(404, (d_508_v19_).length(0))
            (d_508_v19_)[index93_] = len(d_506_v17_)
            index94_ = default__.safeIndex(404, (d_508_v19_).length(0))
            (d_508_v19_)[index94_] = p1
            d_482_v0_ = d_534_cf10_
            d_537_v30_: _dafny.Map
            d_537_v30_ = _dafny.Map({p0: d_503_v16_})
            d_538_v31_: _dafny.Set
            d_538_v31_ = _dafny.Set({len((default__.fm19(globalState)).set(default__.safeIndex((d_508_v19_)[default__.safeIndex(404, (d_508_v19_).length(0))], len(default__.fm19(globalState))), d_534_cf10_)), p1})
            d_539_v32_: _dafny.Map
            d_539_v32_ = _dafny.Map({d_538_v31_: d_503_v16_})
            def iife77_():
                coll33_ = _dafny.Set()
                compr_33_: int
                for compr_33_ in _dafny.IntegerRange(593, 214):
                    d_540_v33_: int = compr_33_
                    if ((593) <= (d_540_v33_)) and ((d_540_v33_) < (214)):
                        coll33_ = coll33_.union(_dafny.Set([(d_540_v33_) * (699)]))
                return _dafny.Set(coll33_)
            def iife78_():
                coll34_ = _dafny.Set()
                compr_34_: int
                for compr_34_ in _dafny.IntegerRange(593, 214):
                    d_541_v33_: int = compr_34_
                    if ((593) <= (d_541_v33_)) and ((d_541_v33_) < (214)):
                        coll34_ = coll34_.union(_dafny.Set([(d_541_v33_) * (699)]))
                return _dafny.Set(coll34_)
            rhs70_ = (d_537_v30_) | (d_537_v30_)
            rhs71_ = ((d_539_v32_)[iife77_()
            ] if (iife78_()
            ) in (d_539_v32_) else d_503_v16_)
            rhs72_ = (d_534_cf10_ if d_482_v0_ else d_534_cf10_)
            rhs73_ = d_534_cf10_
            lhs62_ = globalState
            lhs63_ = globalState
            d_537_v30_ = rhs70_
            d_503_v16_ = rhs71_
            lhs62_.f1 = rhs72_
            lhs63_.f1 = rhs73_
            (globalState).f1 = (D0_DC0(d_482_v0_)).cf0
        elif True:
            d_542___mcc_h5_ = source6_.cf6
            d_543_cf6_ = d_542___mcc_h5_
            d_544_v34_: _dafny.Set
            d_544_v34_ = _dafny.Set({d_482_v0_, d_482_v0_})
            (globalState).f1 = (self).fm16((d_483_v1_ if (self).fm16(default__.fm18(len(d_544_v34_), _dafny.SeqWithoutIsStrInference([p1 for d_545_i6_ in range(default__.abs(411))]), globalState), d_482_v0_, globalState) else d_483_v1_), d_482_v0_, globalState)
            d_546_v35_: _dafny.Map
            d_546_v35_ = _dafny.Map({d_502_v15_: self.f21})
            d_547_v36_: _dafny.Seq
            d_547_v36_ = d_502_v15_
            d_546_v35_ = (d_546_v35_).set(d_547_v36_, (self).f20)
            d_506_v17_ = d_506_v17_
            if d_482_v0_:
                d_548_v37_: _dafny.Seq
                d_548_v37_ = _dafny.SeqWithoutIsStrInference([d_509_v20_, d_509_v20_])
                d_549_v38_: C0
                nw95_ = C0()
                nw95_.ctor__(d_548_v37_, d_482_v0_)
                d_549_v38_ = nw95_
                (self).f21 = (0) - ((0) - (self.f21))
                index95_ = default__.safeIndex(919, (d_503_v16_).length(0))
                (d_503_v16_)[index95_] = (d_549_v38_).f23
                d_550_v39_: _dafny.Seq
                d_550_v39_ = _dafny.SeqWithoutIsStrInference([p0])
                d_551_v40_: _dafny.Map
                d_551_v40_ = _dafny.Map({d_482_v0_: p1})
                index96_ = default__.safeIndex(919, (d_503_v16_).length(0))
                (d_503_v16_)[index96_] = (p0) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ervcnys"))) + ((d_550_v39_)[default__.safeIndex(((d_551_v40_)[(d_549_v38_).f23] if ((d_549_v38_).f23) in (d_551_v40_) else (0) - ((self).f20)), len(d_550_v39_))]))
                d_552_v41_: _dafny.Array
                nw96_ = _dafny.Array(_dafny.Map({}), 19)
                d_552_v41_ = nw96_
                d_553_v42_: _dafny.Map
                d_553_v42_ = _dafny.Map({default__.fm2(globalState): d_482_v0_})
                index97_ = default__.safeIndex(177, (d_552_v41_).length(0))
                (d_552_v41_)[index97_] = _dafny.Map({len(d_553_v42_): (d_549_v38_).f23})
                d_554_v43_: _dafny.Set
                d_554_v43_ = _dafny.Set({(self.f21) - (p1)})
                d_555_v44_: _dafny.Seq
                d_555_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iburkss"))
                d_556_v45_: _dafny.Map
                d_556_v45_ = _dafny.Map({self.f21: True})
                d_557_v46_: _dafny.Seq
                d_557_v46_ = _dafny.SeqWithoutIsStrInference([d_556_v45_])
                d_558_v47_: _dafny.Seq
                d_558_v47_ = _dafny.SeqWithoutIsStrInference([d_506_v17_, d_506_v17_, d_506_v17_])
                index98_ = default__.safeIndex(177, (d_552_v41_).length(0))
                rhs74_ = (d_557_v46_)[default__.safeIndex(p1, len(d_557_v46_))]
                rhs75_ = (d_554_v43_) | (_dafny.Set({self.f21, self.f21}))
                rhs76_ = (_dafny.MultiSet([(d_558_v47_)[default__.safeIndex(len(d_506_v17_), len(d_558_v47_))]])) != (default__.fm20(313, globalState))
                rhs77_ = d_555_v44_
                rhs78_ = (d_549_v38_).f23
                lhs64_ = d_552_v41_
                lhs65_ = default__.safeIndex(177, (d_552_v41_).length(0))
                lhs66_ = globalState
                lhs67_ = globalState
                lhs64_[lhs65_] = rhs74_
                d_554_v43_ = rhs75_
                lhs66_.f1 = rhs76_
                d_555_v44_ = rhs77_
                lhs67_.f1 = rhs78_
                (globalState).f6 = p1
            elif True:
                d_559_v48_: _dafny.MultiSet
                d_559_v48_ = _dafny.MultiSet([d_482_v0_, d_482_v0_, d_482_v0_, d_482_v0_, d_482_v0_])
                (globalState).f1 = (d_559_v48_) != ((((_dafny.MultiSet(d_506_v17_)).set(d_482_v0_, default__.abs(991))).set(True, default__.abs(p1))).intersection(d_559_v48_))
                d_560_v49_: _dafny.Array
                nw97_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_560_v49_ = nw97_
                index99_ = default__.safeIndex(857, (d_560_v49_).length(0))
                (d_560_v49_)[index99_] = d_503_v16_
                index100_ = default__.safeIndex(857, (d_560_v49_).length(0))
                (d_560_v49_)[index100_] = d_503_v16_
                (globalState).f9 = d_508_v19_
                d_506_v17_ = d_506_v17_
                index101_ = default__.safeIndex(93, (d_503_v16_).length(0))
                (d_503_v16_)[index101_] = d_482_v0_
                d_561_v50_: _dafny.Map
                d_561_v50_ = _dafny.Map({d_508_v19_: p0})
                index102_ = default__.safeIndex(93, (d_503_v16_).length(0))
                (d_503_v16_)[index102_] = not((d_543_cf6_) not in (d_561_v50_))
        d_562_v51_: _dafny.Array
        def lambda27_(d_563_i7_):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvrhikco"))

        init16_ = lambda27_
        nw98_ = _dafny.Array(None, 13)
        for i0_16_ in range(nw98_.length(0)):
            nw98_[i0_16_] = init16_(i0_16_)
        d_562_v51_ = nw98_
        d_562_v51_ = d_562_v51_
        (globalState).f6 = p1
        d_564_i8_: int
        d_564_i8_ = 0
        with _dafny.label("3"):
            while d_482_v0_:
                with _dafny.c_label("3"):
                    if (d_564_i8_) >= (100):
                        raise _dafny.Break("3")
                    d_564_i8_ = (d_564_i8_) + (1)
                    d_565_v52_: str
                    d_565_v52_ = _dafny.CodePoint('m')
                    d_566_v53_: D1
                    d_566_v53_ = D1_DC4(d_565_v52_)
                    r0 = (d_566_v53_).cf4
                    (self).f21 = p1
                    (globalState).f1 = False
                    index103_ = default__.safeIndex(331, (d_508_v19_).length(0))
                    (d_508_v19_)[index103_] = (self).f20
                    index104_ = default__.safeIndex(331, (d_508_v19_).length(0))
                    (d_508_v19_)[index104_] = self.f21
                    pass
            pass
        d_567_v54_: str
        d_567_v54_ = _dafny.CodePoint('r')
        r0 = d_567_v54_
        return r0

    def m12(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        if (p2) <= (default__.fm21(globalState)):
            d_568_v0_: bool
            d_568_v0_ = True
            d_569_v1_: _dafny.MultiSet
            d_569_v1_ = _dafny.MultiSet([self.f21, 585, p0, (self).f20])
            d_570_v2_: _dafny.Map
            d_570_v2_ = _dafny.Map({(d_569_v1_).cardinality: d_568_v0_})
            d_571_v3_: _dafny.MultiSet
            d_571_v3_ = _dafny.MultiSet([((d_570_v2_)[(self).f20] if ((self).f20) in (d_570_v2_) else d_568_v0_), not(d_568_v0_)])
            d_572_v4_: _dafny.Map
            d_572_v4_ = _dafny.Map({(p1 if d_568_v0_ else self.f21): (d_571_v3_).cardinality})
            d_572_v4_ = (d_572_v4_).set((self).f20, default__.safeDivisionInt(p1, p1))
            d_573_v5_: _dafny.Array
            def lambda28_(d_574_p1_, d_575_p2_):
                def lambda29_(d_576_i0_):
                    return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f20, 829]) for d_577_i1_ in range(default__.abs(124))]), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_574_p1_, -382]), _dafny.SeqWithoutIsStrInference([self.f21 for d_578_i2_ in range(default__.abs(31))])])])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_575_p2_) for d_579_i5_ in range(default__.abs(751))]) for d_580_i4_ in range(default__.abs(760))]) for d_581_i3_ in range(default__.abs(508))])))

                return lambda29_

            init17_ = lambda28_(p1, p2)
            nw99_ = _dafny.Array(None, 22)
            for i0_17_ in range(nw99_.length(0)):
                nw99_[i0_17_] = init17_(i0_17_)
            d_573_v5_ = nw99_
            d_582_v6_: _dafny.Seq
            d_582_v6_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_583_i6_ in range(default__.abs(-122))]))])
            d_584_v7_: _dafny.Seq
            d_584_v7_ = _dafny.SeqWithoutIsStrInference([d_582_v6_])
            d_585_v8_: _dafny.MultiSet
            d_585_v8_ = _dafny.MultiSet([d_584_v7_, d_584_v7_, d_584_v7_, d_584_v7_, d_584_v7_])
            index105_ = default__.safeIndex(163, (d_573_v5_).length(0))
            (d_573_v5_)[index105_] = d_585_v8_
            d_586_v9_: _dafny.Seq
            d_586_v9_ = _dafny.SeqWithoutIsStrInference([d_585_v8_, (d_585_v8_).intersection(_dafny.MultiSet([default__.fm22(d_568_v0_, d_568_v0_, globalState)])), d_585_v8_, (d_585_v8_).intersection(d_585_v8_)])
            index106_ = default__.safeIndex(163, (d_573_v5_).length(0))
            (d_573_v5_)[index106_] = (d_586_v9_)[default__.safeIndex(((d_571_v3_)[not(default__.fm2(globalState))] if (not(default__.fm2(globalState))) in (d_571_v3_) else -794), len(d_586_v9_))]
            d_587_v10_: str
            d_587_v10_ = _dafny.CodePoint('i')
            d_588_v11_: _dafny.Seq
            d_588_v11_ = _dafny.SeqWithoutIsStrInference([d_568_v0_, d_568_v0_])
            d_589_v12_: D5
            d_589_v12_ = D5_DC12(15, (d_588_v11_)[default__.safeIndex(self.f21, len(d_588_v11_))], d_587_v10_)
            d_590_v13_: D5
            d_590_v13_ = D5_DC11(p0)
            d_591_v14_: _dafny.Seq
            d_591_v14_ = _dafny.SeqWithoutIsStrInference([p1, 714, len(d_572_v4_)])
            d_592_v15_: _dafny.Array
            nw100_ = _dafny.Array(int(0), 20)
            d_592_v15_ = nw100_
            d_593_v16_: _dafny.Map
            d_593_v16_ = _dafny.Map({True: d_592_v15_})
            d_594_v17_: _dafny.Array
            nw101_ = _dafny.Array(None, 29)
            nw101_[int(0)] = (d_569_v1_).cardinality
            nw101_[int(1)] = p0
            nw101_[int(2)] = self.f21
            nw101_[int(3)] = p1
            nw101_[int(4)] = len(_dafny.Map({d_568_v0_: d_587_v10_}))
            nw101_[int(5)] = p1
            nw101_[int(6)] = self.f21
            nw101_[int(7)] = (_dafny.MultiSet([79, p0])).cardinality
            nw101_[int(8)] = 712
            nw101_[int(9)] = self.f21
            nw101_[int(10)] = p1
            nw101_[int(11)] = p1
            nw101_[int(12)] = p1
            nw101_[int(13)] = (self).f20
            nw101_[int(14)] = len(_dafny.SeqWithoutIsStrInference([(d_589_v12_).cf15]))
            nw101_[int(15)] = p1
            nw101_[int(16)] = (self).f20
            nw101_[int(17)] = (self).f20
            nw101_[int(18)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxavydsx")))
            nw101_[int(19)] = (self).f20
            nw101_[int(20)] = p1
            nw101_[int(21)] = 556
            nw101_[int(22)] = (0) - ((d_590_v13_).cf14)
            nw101_[int(23)] = len(d_591_v14_)
            nw101_[int(24)] = p0
            nw101_[int(25)] = self.f21
            nw101_[int(26)] = (self).f20
            nw101_[int(27)] = (0) - (((d_572_v4_)[self.f21] if (self.f21) in (d_572_v4_) else (self).f20))
            nw101_[int(28)] = len(d_593_v16_)
            d_594_v17_ = nw101_
            d_595_v18_: D3
            d_595_v18_ = D3_DC6(d_592_v15_)
            d_596_v19_: _dafny.Seq
            d_596_v19_ = _dafny.SeqWithoutIsStrInference([d_595_v18_])
            d_597_v20_: C0
            nw102_ = C0()
            nw102_.ctor__((d_596_v19_) + (d_596_v19_), d_568_v0_)
            d_597_v20_ = nw102_
            (globalState).f6 = (0) - (((self).f20) + (p1))
            d_598_v21_: _dafny.Map
            d_598_v21_ = _dafny.Map({d_571_v3_: 998})
            d_599_v22_: _dafny.Set
            d_599_v22_ = _dafny.Set({d_594_v17_, (d_595_v18_).cf6, d_592_v15_, d_592_v15_})
            d_600_v23_: _dafny.Map
            d_600_v23_ = _dafny.Map({d_568_v0_: len(d_599_v22_)})
            d_598_v21_ = (d_598_v21_).set(d_571_v3_, len(d_600_v23_))
        elif True:
            d_601_v24_: str
            d_601_v24_ = _dafny.CodePoint('v')
            d_602_v25_: _dafny.Set
            d_602_v25_ = _dafny.Set({d_601_v24_})
            d_602_v25_ = (d_602_v25_).intersection((_dafny.Set({d_601_v24_})).intersection(_dafny.Set({_dafny.CodePoint('j'), d_601_v24_})))
            d_603_v26_: _dafny.Array
            nw103_ = _dafny.Array(False, 28)
            d_603_v26_ = nw103_
            index107_ = default__.safeIndex(792, (d_603_v26_).length(0))
            (d_603_v26_)[index107_] = True
            d_604_v27_: bool
            d_604_v27_ = False
            index108_ = default__.safeIndex(792, (d_603_v26_).length(0))
            (d_603_v26_)[index108_] = d_604_v27_
            d_605_v28_: _dafny.Array
            nw104_ = _dafny.Array(int(0), 17)
            d_605_v28_ = nw104_
            d_606_v29_: _dafny.Map
            d_606_v29_ = _dafny.Map({default__.safeModuloInt(self.f21, (self).f20): (d_605_v28_ if d_604_v27_ else d_605_v28_)})
            d_607_v30_: _dafny.Map
            d_607_v30_ = _dafny.Map({(D5_DC12(p0, (d_603_v26_)[default__.safeIndex(792, (d_603_v26_).length(0))], d_601_v24_)).cf16: d_605_v28_})
            d_606_v29_ = ((d_606_v29_) | (d_606_v29_)).set(p1, ((d_607_v30_)[(d_603_v26_)[default__.safeIndex(792, (d_603_v26_).length(0))]] if ((d_603_v26_)[default__.safeIndex(792, (d_603_v26_).length(0))]) in (d_607_v30_) else d_605_v28_))
            d_608_v31_: _dafny.Seq
            d_608_v31_ = _dafny.SeqWithoutIsStrInference([d_604_v27_, d_604_v27_, d_604_v27_])
            d_609_v32_: _dafny.Seq
            d_609_v32_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_608_v31_))])
            d_610_v33_: _dafny.Seq
            d_610_v33_ = d_609_v32_
            d_611_v34_: _dafny.Map
            d_611_v34_ = _dafny.Map({(d_603_v26_)[default__.safeIndex(792, (d_603_v26_).length(0))]: d_601_v24_})
            d_612_v35_: D3
            d_612_v35_ = D3_DC8(d_609_v32_, p2, (d_603_v26_)[default__.safeIndex(792, (d_603_v26_).length(0))], d_601_v24_)
            d_610_v33_ = default__.fm23(d_611_v34_, (p2) + ((d_612_v35_).cf9), (d_603_v26_)[default__.safeIndex(792, (d_603_v26_).length(0))], (self).f20, globalState)
            r0 = d_604_v27_
        d_613_v36_: bool
        d_613_v36_ = True
        d_614_v37_: _dafny.Seq
        d_614_v37_ = _dafny.SeqWithoutIsStrInference([d_613_v36_])
        hi5_ = len((d_614_v37_).set(default__.safeIndex(p0, len(d_614_v37_)), d_613_v36_))
        for d_615_i7_ in range(default__.safeDivisionInt((self).fm15(globalState), self.f21), hi5_):
            d_616_v38_: _dafny.Array
            nw105_ = _dafny.Array(False, 20)
            d_616_v38_ = nw105_
            d_617_v39_: _dafny.Set
            d_617_v39_ = _dafny.Set({d_616_v38_})
            d_618_v40_: T0
            nw106_ = C1()
            nw106_.ctor__(len(d_617_v39_))
            d_618_v40_ = nw106_
            d_619_v41_: _dafny.Map
            d_619_v41_ = _dafny.Map({d_618_v40_: d_614_v37_})
            d_620_v42_: _dafny.Map
            d_620_v42_ = _dafny.Map({True: _dafny.Map({d_618_v40_: d_614_v37_})})
            d_621_v43_: _dafny.Array
            nw107_ = _dafny.Array(None, 23)
            nw107_[int(0)] = d_619_v41_
            nw107_[int(1)] = d_619_v41_
            nw107_[int(2)] = d_619_v41_
            nw107_[int(3)] = d_619_v41_
            nw107_[int(4)] = d_619_v41_
            nw107_[int(5)] = (d_619_v41_) | (d_619_v41_)
            nw107_[int(6)] = d_619_v41_
            nw107_[int(7)] = (d_619_v41_) | (d_619_v41_)
            nw107_[int(8)] = d_619_v41_
            nw107_[int(9)] = (_dafny.Map({d_618_v40_: d_614_v37_})) | (d_619_v41_)
            nw107_[int(10)] = d_619_v41_
            nw107_[int(11)] = d_619_v41_
            nw107_[int(12)] = (d_619_v41_).set(d_618_v40_, d_614_v37_)
            nw107_[int(13)] = (d_619_v41_).set(d_618_v40_, d_614_v37_)
            nw107_[int(14)] = d_619_v41_
            nw107_[int(15)] = d_619_v41_
            nw107_[int(16)] = _dafny.Map({d_618_v40_: d_614_v37_})
            nw107_[int(17)] = d_619_v41_
            nw107_[int(18)] = (d_619_v41_).set(d_618_v40_, d_614_v37_)
            nw107_[int(19)] = d_619_v41_
            nw107_[int(20)] = d_619_v41_
            nw107_[int(21)] = ((d_620_v42_)[d_613_v36_] if (d_613_v36_) in (d_620_v42_) else d_619_v41_)
            nw107_[int(22)] = (d_619_v41_) | ((d_619_v41_).set(d_618_v40_, d_614_v37_))
            d_621_v43_ = nw107_
            index109_ = default__.safeIndex(277, (d_621_v43_).length(0))
            (d_621_v43_)[index109_] = _dafny.Map({d_618_v40_: d_614_v37_})
            index110_ = default__.safeIndex(277, (d_621_v43_).length(0))
            (d_621_v43_)[index110_] = d_619_v41_
            d_622_v44_: _dafny.Map
            d_622_v44_ = _dafny.Map({((0) - ((0) - (p1))) > (d_615_i7_): -522})
            d_622_v44_ = (d_622_v44_).set((d_613_v36_) and (d_613_v36_), (0) - (p1))
            d_623_v45_: _dafny.Array
            nw108_ = _dafny.Array(_dafny.Array(None, 0), 21)
            d_623_v45_ = nw108_
            d_624_v46_: D5
            d_624_v46_ = D5_DC10(d_623_v45_)
            d_625_v47_: _dafny.Seq
            d_625_v47_ = _dafny.SeqWithoutIsStrInference([d_624_v46_])
            d_626_v48_: D7
            d_626_v48_ = D7_DC15(d_618_v40_)
            d_627_v49_: _dafny.MultiSet
            d_627_v49_ = _dafny.MultiSet([(d_626_v48_).cf20])
            d_628_v50_: _dafny.Array
            nw109_ = _dafny.Array(None, 13)
            nw109_[int(0)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_629_i8_ in range(default__.abs(573))]))
            nw109_[int(1)] = 680
            nw109_[int(2)] = p1
            nw109_[int(3)] = (0) - (((self).f20) * (len((d_625_v47_).set(default__.safeIndex(p0, len(d_625_v47_)), D5_DC10(d_623_v45_)))))
            nw109_[int(4)] = (self).f20
            nw109_[int(5)] = 210
            nw109_[int(6)] = ((d_627_v49_)[d_618_v40_] if (d_618_v40_) in (d_627_v49_) else self.f21)
            nw109_[int(7)] = len(d_614_v37_)
            nw109_[int(8)] = default__.safeDivisionInt(len(_dafny.Set({d_613_v36_})), 367)
            nw109_[int(9)] = (p0) * (704)
            nw109_[int(10)] = p1
            nw109_[int(11)] = ((self).f20) + (p0)
            nw109_[int(12)] = p0
            d_628_v50_ = nw109_
            index111_ = default__.safeIndex(702, (d_628_v50_).length(0))
            (d_628_v50_)[index111_] = -622
            index112_ = default__.safeIndex(702, (d_628_v50_).length(0))
            rhs79_ = (self).f20
            rhs80_ = 210
            lhs68_ = d_628_v50_
            lhs69_ = default__.safeIndex(702, (d_628_v50_).length(0))
            lhs70_ = globalState
            lhs68_[lhs69_] = rhs79_
            lhs70_.f6 = rhs80_
            d_630_v51_: _dafny.Map
            d_630_v51_ = _dafny.Map({d_613_v36_: d_613_v36_})
            d_631_v52_: _dafny.Set
            d_631_v52_ = _dafny.Set({d_613_v36_})
            if ((len(d_630_v51_)) <= (len(d_631_v52_))) == (False):
                d_632_v53_: _dafny.Set
                d_632_v53_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))), ((d_628_v50_)[default__.safeIndex(702, (d_628_v50_).length(0))]) * ((self).f20), default__.safeDivisionInt((d_628_v50_)[default__.safeIndex(702, (d_628_v50_).length(0))], p0)})
                d_632_v53_ = d_632_v53_
                d_633_v54_: _dafny.Array
                nw110_ = _dafny.Array(_dafny.Map({}), 2)
                d_633_v54_ = nw110_
                index113_ = default__.safeIndex(735, (d_633_v54_).length(0))
                (d_633_v54_)[index113_] = _dafny.Map({d_628_v50_: d_615_i7_})
                d_634_v55_: _dafny.Map
                d_634_v55_ = _dafny.Map({d_628_v50_: p1})
                index114_ = default__.safeIndex(735, (d_633_v54_).length(0))
                (d_633_v54_)[index114_] = (_dafny.Map({d_628_v50_: p0})) | ((d_634_v55_ if d_613_v36_ else _dafny.Map({d_628_v50_: len(d_614_v37_)})))
                (globalState).f6 = -256
                d_635_v56_: _dafny.MultiSet
                d_635_v56_ = _dafny.MultiSet([d_613_v36_])
                d_636_v57_: D5
                d_636_v57_ = D5_DC11(p0)
                d_637_v58_: _dafny.Array
                nw111_ = _dafny.Array(None, 4)
                nw111_[int(0)] = D5_DC11(p1)
                nw111_[int(1)] = d_636_v57_
                nw111_[int(2)] = d_636_v57_
                nw111_[int(3)] = d_636_v57_
                d_637_v58_ = nw111_
                d_638_v59_: _dafny.Seq
                d_638_v59_ = _dafny.SeqWithoutIsStrInference([d_637_v58_, d_637_v58_, d_637_v58_])
                d_639_v60_: _dafny.Array
                nw112_ = _dafny.Array(None, 27)
                nw112_[int(0)] = d_637_v58_
                nw112_[int(1)] = d_637_v58_
                nw112_[int(2)] = d_637_v58_
                nw112_[int(3)] = d_637_v58_
                nw112_[int(4)] = d_637_v58_
                nw112_[int(5)] = d_637_v58_
                nw112_[int(6)] = d_637_v58_
                nw112_[int(7)] = d_637_v58_
                nw112_[int(8)] = d_637_v58_
                nw112_[int(9)] = d_637_v58_
                nw112_[int(10)] = d_637_v58_
                nw112_[int(11)] = d_637_v58_
                nw112_[int(12)] = d_637_v58_
                nw112_[int(13)] = d_637_v58_
                nw112_[int(14)] = d_637_v58_
                nw112_[int(15)] = d_637_v58_
                nw112_[int(16)] = d_637_v58_
                nw112_[int(17)] = d_637_v58_
                nw112_[int(18)] = d_637_v58_
                nw112_[int(19)] = d_637_v58_
                nw112_[int(20)] = d_637_v58_
                nw112_[int(21)] = d_637_v58_
                nw112_[int(22)] = d_637_v58_
                nw112_[int(23)] = d_637_v58_
                nw112_[int(24)] = d_637_v58_
                nw112_[int(25)] = (d_638_v59_)[default__.safeIndex(d_615_i7_, len(d_638_v59_))]
                nw112_[int(26)] = d_637_v58_
                d_639_v60_ = nw112_
                d_640_v61_: _dafny.Map
                d_640_v61_ = _dafny.Map({(d_635_v56_).isdisjoint(d_635_v56_): d_639_v60_})
                d_640_v61_ = (d_640_v61_).set((d_614_v37_)[default__.safeIndex(p0, len(d_614_v37_))], d_639_v60_)
                d_641_v62_: str
                d_641_v62_ = _dafny.CodePoint('b')
                d_641_v62_ = d_641_v62_
            elif True:
                d_642_v63_: _dafny.Seq
                d_642_v63_ = _dafny.SeqWithoutIsStrInference([p1, self.f21, len((p2) + (p2))])
                rhs81_ = (d_642_v63_)[default__.safeIndex(len(p3), len(d_642_v63_))]
                rhs82_ = d_613_v36_
                lhs71_ = globalState
                lhs71_.f6 = rhs81_
                d_613_v36_ = rhs82_
                d_643_v64_: _dafny.Set
                d_643_v64_ = _dafny.Set({(d_628_v50_)[default__.safeIndex(702, (d_628_v50_).length(0))], (self).f20, d_615_i7_})
                d_644_v65_: _dafny.Map
                d_644_v65_ = _dafny.Map({D3_DC6(d_628_v50_): (d_643_v64_) | (_dafny.Set({self.f21}))})
                d_645_v66_: D3
                d_645_v66_ = D3_DC6(d_628_v50_)
                pat_let_tv12_ = d_628_v50_
                pat_let_tv13_ = d_628_v50_
                def iife79_(_pat_let22_0):
                    def iife80_(d_646_dt__update__tmp_h1_):
                        def iife81_(_pat_let23_0):
                            def iife82_(d_647_dt__update_hcf6_h1_):
                                return D3_DC6(d_647_dt__update_hcf6_h1_)
                            return iife82_(_pat_let23_0)
                        return iife81_(pat_let_tv12_)
                    return iife80_(_pat_let22_0)
                def iife83_(_pat_let24_0):
                    def iife84_(d_648_dt__update__tmp_h0_):
                        def iife85_(_pat_let25_0):
                            def iife86_(d_649_dt__update_hcf6_h0_):
                                return D3_DC6(d_649_dt__update_hcf6_h0_)
                            return iife86_(_pat_let25_0)
                        return iife85_(pat_let_tv13_)
                    return iife84_(_pat_let24_0)
                d_643_v64_ = ((d_644_v65_)[iife79_(d_645_v66_)] if (iife83_(d_645_v66_)) in (d_644_v65_) else (default__.fm28((d_628_v50_)[default__.safeIndex(702, (d_628_v50_).length(0))], d_613_v36_, d_613_v36_, globalState) if d_613_v36_ else d_643_v64_))
                d_650_v67_: _dafny.MultiSet
                d_650_v67_ = _dafny.MultiSet([d_615_i7_, default__.fm1(d_613_v36_, globalState)])
                d_650_v67_ = _dafny.MultiSet([(0) - (self.f21)])
                (globalState).f1 = not(d_613_v36_)
                d_651_v68_: str
                d_651_v68_ = _dafny.CodePoint('j')
                d_652_v69_: _dafny.MultiSet
                d_652_v69_ = _dafny.MultiSet([d_613_v36_])
                d_651_v68_ = default__.fm29(d_613_v36_, (((d_650_v67_)[(d_628_v50_)[default__.safeIndex(702, (d_628_v50_).length(0))]] if ((d_628_v50_)[default__.safeIndex(702, (d_628_v50_).length(0))]) in (d_650_v67_) else (d_652_v69_).cardinality)) * (p0), len(d_642_v63_), (d_631_v52_).ispropersubset(d_631_v52_), globalState)
        rhs83_ = (default__.fm1(not(d_613_v36_), globalState)) - ((self).f20)
        rhs84_ = len(p2)
        rhs85_ = d_613_v36_
        lhs72_ = globalState
        lhs73_ = globalState
        lhs74_ = globalState
        lhs72_.f6 = rhs83_
        lhs73_.f6 = rhs84_
        lhs74_.f1 = rhs85_
        d_653_v70_: _dafny.Map
        d_653_v70_ = _dafny.Map({d_613_v36_: d_613_v36_})
        d_654_v71_: _dafny.Map
        d_654_v71_ = _dafny.Map({d_613_v36_: len(d_653_v70_)})
        d_655_v72_: _dafny.Seq
        d_655_v72_ = _dafny.SeqWithoutIsStrInference([((d_654_v71_)[d_613_v36_] if (d_613_v36_) in (d_654_v71_) else (self).f20), p0, (self).f20, self.f21])
        d_656_v73_: str
        d_656_v73_ = _dafny.CodePoint('u')
        d_657_v74_: D3
        d_657_v74_ = D3_DC8(d_655_v72_, p2, d_613_v36_, d_656_v73_)
        d_658_v75_: D5
        d_658_v75_ = D5_DC11(p0)
        (globalState).f6 = default__.safeDivisionInt((0) - ((0) - (len((d_657_v74_).cf9))), (d_658_v75_).cf14)
        d_659_i9_: int
        d_659_i9_ = 0
        with _dafny.label("4"):
            while (p0) >= ((self).f20):
                with _dafny.c_label("4"):
                    if (d_659_i9_) >= (100):
                        raise _dafny.Break("4")
                    d_659_i9_ = (d_659_i9_) + (1)
                    d_660_v76_: _dafny.Array
                    nw113_ = _dafny.Array(False, 7)
                    d_660_v76_ = nw113_
                    index115_ = default__.safeIndex(386, (d_660_v76_).length(0))
                    (d_660_v76_)[index115_] = d_613_v36_
                    d_661_v77_: _dafny.MultiSet
                    d_661_v77_ = _dafny.MultiSet([d_613_v36_, False])
                    d_662_v78_: _dafny.MultiSet
                    d_662_v78_ = _dafny.MultiSet([d_655_v72_])
                    index116_ = default__.safeIndex(386, (d_660_v76_).length(0))
                    (d_660_v76_)[index116_] = not((((d_661_v77_).cardinality) > (p1)) and ((d_662_v78_).isdisjoint(_dafny.MultiSet([d_655_v72_]))))
                    d_663_v79_: _dafny.Array
                    def lambda30_(d_664_i10_):
                        return (d_664_i10_) - ((self).f20)

                    init18_ = lambda30_
                    nw114_ = _dafny.Array(None, 17)
                    for i0_18_ in range(nw114_.length(0)):
                        nw114_[i0_18_] = init18_(i0_18_)
                    d_663_v79_ = nw114_
                    index117_ = default__.safeIndex(275, (d_663_v79_).length(0))
                    (d_663_v79_)[index117_] = (0) - (p1)
                    index118_ = default__.safeIndex(275, (d_663_v79_).length(0))
                    (d_663_v79_)[index118_] = default__.safeDivisionInt(p0, (0) - (self.f21))
                    d_665_v80_: D3
                    d_665_v80_ = D3_DC6(d_663_v79_)
                    d_666_v81_: _dafny.Array
                    nw115_ = _dafny.Array(None, 10)
                    nw115_[int(0)] = d_663_v79_
                    nw115_[int(1)] = d_663_v79_
                    nw115_[int(2)] = d_663_v79_
                    nw115_[int(3)] = d_663_v79_
                    nw115_[int(4)] = d_663_v79_
                    nw115_[int(5)] = (d_663_v79_ if not((d_660_v76_)[default__.safeIndex(386, (d_660_v76_).length(0))]) else d_663_v79_)
                    nw115_[int(6)] = (d_665_v80_).cf6
                    nw115_[int(7)] = d_663_v79_
                    nw115_[int(8)] = d_663_v79_
                    nw115_[int(9)] = d_663_v79_
                    d_666_v81_ = nw115_
                    d_667_v82_: _dafny.Map
                    d_667_v82_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, d_613_v36_])): d_663_v79_})
                    index119_ = default__.safeIndex(355, (d_666_v81_).length(0))
                    (d_666_v81_)[index119_] = ((d_667_v82_)[len(d_655_v72_)] if (len(d_655_v72_)) in (d_667_v82_) else d_663_v79_)
                    index120_ = default__.safeIndex(355, (d_666_v81_).length(0))
                    (d_666_v81_)[index120_] = d_663_v79_
                    d_668_v83_: _dafny.MultiSet
                    d_668_v83_ = _dafny.MultiSet([d_661_v77_, d_661_v77_, d_661_v77_])
                    d_669_v84_: _dafny.MultiSet
                    d_669_v84_ = _dafny.MultiSet([d_614_v37_])
                    d_670_v85_: _dafny.MultiSet
                    d_670_v85_ = _dafny.MultiSet([p0])
                    d_671_v86_: _dafny.Seq
                    d_671_v86_ = _dafny.SeqWithoutIsStrInference([d_661_v77_])
                    d_672_v87_: _dafny.Seq
                    d_672_v87_ = _dafny.SeqWithoutIsStrInference([d_668_v83_])
                    d_673_v88_: _dafny.Set
                    d_673_v88_ = _dafny.Set({(d_660_v76_)[default__.safeIndex(386, (d_660_v76_).length(0))], (d_660_v76_)[default__.safeIndex(386, (d_660_v76_).length(0))], (d_660_v76_)[default__.safeIndex(386, (d_660_v76_).length(0))]})
                    d_674_v89_: _dafny.Array
                    nw116_ = _dafny.Array(None, 22)
                    nw116_[int(0)] = d_668_v83_
                    nw116_[int(1)] = default__.fm30(not((d_660_v76_)[default__.safeIndex(386, (d_660_v76_).length(0))]), _dafny.Set({(d_669_v84_).cardinality}), globalState)
                    nw116_[int(2)] = d_668_v83_
                    nw116_[int(3)] = d_668_v83_
                    nw116_[int(4)] = d_668_v83_
                    nw116_[int(5)] = _dafny.MultiSet([d_661_v77_])
                    nw116_[int(6)] = (_dafny.MultiSet([_dafny.MultiSet([(d_660_v76_)[default__.safeIndex(386, (d_660_v76_).length(0))]])])) - (d_668_v83_)
                    nw116_[int(7)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_613_v36_]) for d_675_i11_ in range(default__.abs(562))]))
                    nw116_[int(8)] = (d_668_v83_) - (_dafny.MultiSet([d_661_v77_, d_661_v77_]))
                    nw116_[int(9)] = ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_613_v36_])), d_661_v77_, d_661_v77_, d_661_v77_]))).set(_dafny.MultiSet([(d_660_v76_)[default__.safeIndex(386, (d_660_v76_).length(0))], (d_660_v76_)[default__.safeIndex(386, (d_660_v76_).length(0))], d_613_v36_, (d_660_v76_)[default__.safeIndex(386, (d_660_v76_).length(0))], (d_660_v76_)[default__.safeIndex(386, (d_660_v76_).length(0))]]), default__.abs(((d_670_v85_)[self.f21] if (self.f21) in (d_670_v85_) else (self).f20)))) - (d_668_v83_)
                    nw116_[int(10)] = (_dafny.MultiSet(d_671_v86_)) | (d_668_v83_)
                    nw116_[int(11)] = _dafny.MultiSet(d_671_v86_)
                    nw116_[int(12)] = d_668_v83_
                    nw116_[int(13)] = (d_672_v87_)[default__.safeIndex(len(d_673_v88_), len(d_672_v87_))]
                    nw116_[int(14)] = _dafny.MultiSet([d_661_v77_])
                    nw116_[int(15)] = d_668_v83_
                    nw116_[int(16)] = d_668_v83_
                    nw116_[int(17)] = d_668_v83_
                    nw116_[int(18)] = d_668_v83_
                    nw116_[int(19)] = (d_668_v83_).set(d_661_v77_, default__.abs(len((p2).set(default__.safeIndex((d_663_v79_)[default__.safeIndex(275, (d_663_v79_).length(0))], len(p2)), d_656_v73_))))
                    nw116_[int(20)] = (_dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_613_v36_])), _dafny.MultiSet(d_614_v37_)]) if (d_660_v76_)[default__.safeIndex(386, (d_660_v76_).length(0))] else _dafny.MultiSet([d_661_v77_, d_661_v77_]))
                    nw116_[int(21)] = _dafny.MultiSet((d_671_v86_) + (d_671_v86_))
                    d_674_v89_ = nw116_
                    index121_ = default__.safeIndex(286, (d_674_v89_).length(0))
                    (d_674_v89_)[index121_] = (d_668_v83_).set(d_661_v77_, default__.abs((d_663_v79_)[default__.safeIndex(275, (d_663_v79_).length(0))]))
                    index122_ = default__.safeIndex(286, (d_674_v89_).length(0))
                    (d_674_v89_)[index122_] = d_668_v83_
                    pass
            pass
        d_676_v90_: _dafny.Array
        nw117_ = _dafny.Array(int(0), 4)
        d_676_v90_ = nw117_
        d_677_v91_: D3
        d_677_v91_ = D3_DC6(d_676_v90_)
        d_678_v92_: D3
        d_678_v92_ = D3_DC6((d_677_v91_).cf6)
        d_679_v93_: _dafny.Seq
        d_679_v93_ = _dafny.SeqWithoutIsStrInference([d_677_v91_])
        d_680_v94_: C0
        nw118_ = C0()
        nw118_.ctor__(d_679_v93_, d_613_v36_)
        d_680_v94_ = nw118_
        d_681_v95_: _dafny.Map
        d_681_v95_ = _dafny.Map({d_680_v94_: p0})
        hi6_ = 945
        for d_682_i12_ in range(default__.safeModuloInt(len(d_653_v70_), len(d_681_v95_)), hi6_):
            d_653_v70_ = (d_653_v70_).set(not(default__.fm2(globalState)), (d_680_v94_).f23)
            d_683_v96_: _dafny.Set
            d_683_v96_ = _dafny.Set({(d_680_v94_).f23})
            d_684_v97_: D0
            d_684_v97_ = D0_DC1((0) - (len((_dafny.SeqWithoutIsStrInference([p1])) + (_dafny.SeqWithoutIsStrInference([218, p0, self.f21, len(d_683_v96_)])))))
            pat_let_tv14_ = d_613_v36_
            pat_let_tv15_ = globalState
            def iife87_(_pat_let26_0):
                def iife88_(d_685_dt__update__tmp_h2_):
                    def iife89_(_pat_let27_0):
                        def iife90_(d_686_dt__update_hcf1_h0_):
                            return D0_DC1(d_686_dt__update_hcf1_h0_)
                        return iife90_(_pat_let27_0)
                    return iife89_(default__.fm1(pat_let_tv14_, pat_let_tv15_))
                return iife88_(_pat_let26_0)
            d_684_v97_ = iife87_(D0_DC1((self).f20))
            d_687_v98_: _dafny.Array
            nw119_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
            d_687_v98_ = nw119_
            index123_ = default__.safeIndex(192, (d_687_v98_).length(0))
            (d_687_v98_)[index123_] = p3
            index124_ = default__.safeIndex(192, (d_687_v98_).length(0))
            (d_687_v98_)[index124_] = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_688_i13_ in range(default__.abs(4))])) + (p3)) + (p3)
            if d_613_v36_:
                (globalState).f6 = self.f21
                d_689_v99_: C0
                nw120_ = C0()
                nw120_.ctor__((d_680_v94_).f22, d_613_v36_)
                d_689_v99_ = nw120_
                (globalState).f1 = (d_680_v94_).f23
                d_690_v100_: _dafny.MultiSet
                d_690_v100_ = _dafny.MultiSet([((d_680_v94_).f23) == ((d_689_v99_).f23)])
                d_691_v102_: _dafny.Map
                d_691_v102_ = _dafny.Map({d_654_v71_: (d_680_v94_).f23})
                d_692_v103_: _dafny.Array
                nw121_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_692_v103_ = nw121_
                d_693_v104_: _dafny.Seq
                d_693_v104_ = _dafny.SeqWithoutIsStrInference([d_692_v103_, d_692_v103_, d_692_v103_])
                d_694_v105_: D5
                d_694_v105_ = D5_DC10((d_693_v104_)[default__.safeIndex(p1, len(d_693_v104_))])
                index125_ = default__.safeIndex(192, (d_687_v98_).length(0))
                def iife91_():
                    coll35_ = _dafny.Map()
                    compr_35_: _dafny.Map
                    for compr_35_ in (d_691_v102_).keys.Elements:
                        d_695_v101_: _dafny.Map = compr_35_
                        if (d_695_v101_) in (d_691_v102_):
                            coll35_[d_695_v101_] = (p3) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsri")))
                    return _dafny.Map(coll35_)
                rhs86_ = (0) - (((self).f20) - ((self.f21) + (d_682_i12_)))
                rhs87_ = (0) - (len(iife91_()
                ))
                rhs88_ = (d_690_v100_).set(d_613_v36_, default__.abs((0) - ((self).fm15(globalState))))
                rhs89_ = (d_687_v98_)[default__.safeIndex(192, (d_687_v98_).length(0))]
                rhs90_ = (((d_689_v99_).f23 if (d_680_v94_).f23 else (d_680_v94_).f23) if (D7_DC17((default__.fm18(self.f21, d_655_v72_, globalState)).cf0, (d_680_v94_).fm17(d_682_i12_, p0, d_613_v36_, globalState), d_689_v99_, d_694_v105_, False)).cf24 else (d_689_v99_).f23)
                lhs75_ = globalState
                lhs76_ = globalState
                lhs77_ = d_687_v98_
                lhs78_ = default__.safeIndex(192, (d_687_v98_).length(0))
                lhs79_ = globalState
                lhs75_.f6 = rhs86_
                lhs76_.f6 = rhs87_
                d_690_v100_ = rhs88_
                lhs77_[lhs78_] = rhs89_
                lhs79_.f1 = rhs90_
                d_696_v106_: _dafny.Map
                d_696_v106_ = _dafny.Map({987: p2})
                d_696_v106_ = (d_696_v106_).set((self).f20, (d_687_v98_)[default__.safeIndex(192, (d_687_v98_).length(0))])
            elif True:
                d_697_v107_: _dafny.Map
                d_697_v107_ = (d_654_v71_) | (d_654_v71_)
                d_697_v107_ = d_697_v107_
                d_698_v108_: _dafny.Array
                def lambda31_(d_699_v94_):
                    def lambda32_(d_700_i14_):
                        return (d_699_v94_).f23

                    return lambda32_

                init19_ = lambda31_(d_680_v94_)
                nw122_ = _dafny.Array(None, 4)
                for i0_19_ in range(nw122_.length(0)):
                    nw122_[i0_19_] = init19_(i0_19_)
                d_698_v108_ = nw122_
                d_701_v109_: _dafny.Map
                d_701_v109_ = _dafny.Map({len(d_614_v37_): d_698_v108_})
                d_701_v109_ = ((d_701_v109_) | (_dafny.Map({p0: d_698_v108_}))) | (d_701_v109_)
                d_702_v110_: D0
                d_702_v110_ = D0_DC0(not(not(False)))
                d_701_v109_ = (d_701_v109_).set(len(_dafny.Set({(self).fm16(d_702_v110_, (d_680_v94_).f23, globalState)})), d_698_v108_)
                (globalState).f1 = (default__.safeDivisionInt(len(d_655_v72_), self.f21)) == (default__.fm1((d_680_v94_).f23, globalState))
                r1 = (default__.fm21(globalState)) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_703_i15_ in range(default__.abs(-971))]))
        r0 = not(d_613_v36_)
        r1 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))) + (p2)) + (p2)
        return r0, r1

    def m14(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_704_v0_: _dafny.Array
        nw123_ = _dafny.Array(int(0), 5)
        d_704_v0_ = nw123_
        d_705_v1_: D3
        d_705_v1_ = D3_DC6(d_704_v0_)
        d_706_v2_: _dafny.Set
        d_706_v2_ = _dafny.Set({p1})
        d_707_v3_: bool
        d_708_v4_: _dafny.Map
        d_709_v5_: _dafny.Array
        out9_: bool
        out10_: _dafny.Map
        out11_: _dafny.Array
        out9_, out10_, out11_ = (self).m15(d_705_v1_, p0, (len(d_706_v2_)) * (p2), p2, globalState)
        d_707_v3_ = out9_
        d_708_v4_ = out10_
        d_709_v5_ = out11_
        d_710_v6_: _dafny.Map
        d_710_v6_ = _dafny.Map({True: p2})
        d_711_v7_: _dafny.Map
        d_711_v7_ = d_710_v6_
        source7_ = d_711_v7_
        d_712___mcc_h0_ = source7_
        d_713_cf12_ = d_712___mcc_h0_
        d_714_v8_: _dafny.Map
        d_714_v8_ = _dafny.Map({d_707_v3_: d_707_v3_})
        d_715_v9_: C1
        nw124_ = C1()
        nw124_.ctor__(len(d_714_v8_))
        d_715_v9_ = nw124_
        (globalState).f6 = p2
        index126_ = default__.safeIndex(500, (d_704_v0_).length(0))
        (d_704_v0_)[index126_] = 247
        index127_ = default__.safeIndex(500, (d_704_v0_).length(0))
        (d_704_v0_)[index127_] = (0) - ((self).f20)
        if (len((d_714_v8_).set(d_707_v3_, p0))) > (self.f21):
            d_716_v10_: _dafny.Seq
            d_716_v10_ = _dafny.SeqWithoutIsStrInference([d_705_v1_])
            d_717_v11_: _dafny.Array
            nw125_ = _dafny.Array(None, 20)
            nw125_[int(0)] = d_707_v3_
            nw125_[int(1)] = False
            nw125_[int(2)] = p1
            nw125_[int(3)] = d_707_v3_
            nw125_[int(4)] = p0
            nw125_[int(5)] = True
            nw125_[int(6)] = True
            nw125_[int(7)] = d_707_v3_
            nw125_[int(8)] = p0
            nw125_[int(9)] = p1
            nw125_[int(10)] = False
            nw125_[int(11)] = d_707_v3_
            nw125_[int(12)] = d_707_v3_
            nw125_[int(13)] = not(d_707_v3_)
            nw125_[int(14)] = True
            nw125_[int(15)] = p1
            nw125_[int(16)] = not(False)
            nw125_[int(17)] = p1
            nw125_[int(18)] = p1
            nw125_[int(19)] = d_707_v3_
            d_717_v11_ = nw125_
            d_718_v12_: _dafny.Map
            d_718_v12_ = _dafny.Map({d_717_v11_: d_717_v11_})
            d_719_v13_: C0
            nw126_ = C0()
            nw126_.ctor__(d_716_v10_, (self.f21) > (len(d_718_v12_)))
            d_719_v13_ = nw126_
            d_720_v14_: _dafny.Seq
            d_720_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdistbhg"))
            d_720_v14_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_721_i0_ in range(default__.abs(-230))])
            d_722_v15_: _dafny.Set
            d_722_v15_ = _dafny.Set({d_708_v4_, d_708_v4_})
            d_723_v16_: _dafny.Seq
            d_723_v16_ = _dafny.SeqWithoutIsStrInference([(0) - (((self).f20) + (252)), (d_704_v0_)[default__.safeIndex(500, (d_704_v0_).length(0))], (d_715_v9_).f24, (len(d_722_v15_)) - (p2), (d_715_v9_).f24])
            d_723_v16_ = d_723_v16_
            d_724_v17_: _dafny.Seq
            d_724_v17_ = _dafny.SeqWithoutIsStrInference([d_706_v2_, d_706_v2_])
            d_725_v18_: _dafny.Array
            nw127_ = _dafny.Array(None, 26)
            nw127_[int(0)] = (d_706_v2_) - (d_706_v2_)
            nw127_[int(1)] = (d_706_v2_) | (d_706_v2_)
            nw127_[int(2)] = (d_706_v2_) - (_dafny.Set({(d_719_v13_).f23, (d_719_v13_).f23}))
            nw127_[int(3)] = d_706_v2_
            nw127_[int(4)] = (d_706_v2_).intersection(_dafny.Set({p0}))
            nw127_[int(5)] = d_706_v2_
            nw127_[int(6)] = d_706_v2_
            nw127_[int(7)] = (d_706_v2_).intersection(d_706_v2_)
            nw127_[int(8)] = default__.fm31(self.f21, p0, globalState)
            nw127_[int(9)] = d_706_v2_
            nw127_[int(10)] = d_706_v2_
            nw127_[int(11)] = d_706_v2_
            nw127_[int(12)] = _dafny.Set({False, True})
            nw127_[int(13)] = (_dafny.Set({d_707_v3_}) if d_707_v3_ else _dafny.Set({p0}))
            nw127_[int(14)] = d_706_v2_
            nw127_[int(15)] = d_706_v2_
            nw127_[int(16)] = d_706_v2_
            nw127_[int(17)] = d_706_v2_
            nw127_[int(18)] = d_706_v2_
            nw127_[int(19)] = d_706_v2_
            nw127_[int(20)] = (d_724_v17_)[default__.safeIndex((d_715_v9_).f24, len(d_724_v17_))]
            nw127_[int(21)] = (d_706_v2_) - (_dafny.Set({True, (d_719_v13_).f23}))
            nw127_[int(22)] = d_706_v2_
            nw127_[int(23)] = (d_706_v2_) - (_dafny.Set({d_707_v3_, p1, d_707_v3_}))
            nw127_[int(24)] = (d_706_v2_) - (d_706_v2_)
            nw127_[int(25)] = d_706_v2_
            d_725_v18_ = nw127_
            index128_ = default__.safeIndex(323, (d_725_v18_).length(0))
            (d_725_v18_)[index128_] = (d_706_v2_ if (d_719_v13_).f23 else d_706_v2_)
            index129_ = default__.safeIndex(323, (d_725_v18_).length(0))
            (d_725_v18_)[index129_] = (d_706_v2_).intersection(d_706_v2_)
            d_707_v3_ = (d_719_v13_).f23
        elif True:
            d_726_v19_: _dafny.Array
            nw128_ = _dafny.Array(None, 11)
            nw128_[int(0)] = d_707_v3_
            nw128_[int(1)] = d_707_v3_
            nw128_[int(2)] = p1
            nw128_[int(3)] = d_707_v3_
            nw128_[int(4)] = (self.f21) <= ((d_715_v9_).f24)
            nw128_[int(5)] = (d_706_v2_).issubset(d_706_v2_)
            nw128_[int(6)] = p1
            nw128_[int(7)] = ((self).f20) <= ((d_704_v0_)[default__.safeIndex(500, (d_704_v0_).length(0))])
            nw128_[int(8)] = (d_707_v3_) and (p1)
            nw128_[int(9)] = ((d_704_v0_)[default__.safeIndex(500, (d_704_v0_).length(0))]) == ((d_704_v0_)[default__.safeIndex(500, (d_704_v0_).length(0))])
            nw128_[int(10)] = d_707_v3_
            d_726_v19_ = nw128_
            index130_ = default__.safeIndex(309, (d_726_v19_).length(0))
            (d_726_v19_)[index130_] = p0
            index131_ = default__.safeIndex(309, (d_726_v19_).length(0))
            (d_726_v19_)[index131_] = p0
            index132_ = default__.safeIndex(500, (d_704_v0_).length(0))
            (d_704_v0_)[index132_] = self.f21
            d_727_v20_: _dafny.Seq
            d_727_v20_ = _dafny.SeqWithoutIsStrInference([d_705_v1_, d_705_v1_])
            d_728_v21_: C0
            nw129_ = C0()
            nw129_.ctor__((d_727_v20_ if False else d_727_v20_), d_707_v3_)
            d_728_v21_ = nw129_
            d_729_v22_: _dafny.Array
            nw130_ = _dafny.Array(_dafny.Seq({}), 13)
            d_729_v22_ = nw130_
            d_730_v23_: _dafny.Seq
            d_730_v23_ = _dafny.SeqWithoutIsStrInference([p3, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bt"))])
            index133_ = default__.safeIndex(737, (d_729_v22_).length(0))
            (d_729_v22_)[index133_] = d_730_v23_
            index134_ = default__.safeIndex(737, (d_729_v22_).length(0))
            (d_729_v22_)[index134_] = d_730_v23_
            d_714_v8_ = (d_714_v8_).set((d_728_v21_).f23, p0)
        d_731_i1_: int
        d_731_i1_ = 0
        with _dafny.label("5"):
            while default__.fm2(globalState):
                with _dafny.c_label("5"):
                    if (d_731_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_731_i1_ = (d_731_i1_) + (1)
                    d_732_v24_: _dafny.MultiSet
                    d_732_v24_ = _dafny.MultiSet([True])
                    d_733_v25_: _dafny.Map
                    d_733_v25_ = _dafny.Map({(d_732_v24_).issubset(_dafny.MultiSet([not(p1)])): p3})
                    (self).f21 = len(((d_733_v25_)[(p2) != (p2)] if ((p2) != (p2)) in (d_733_v25_) else (p3) + (p3)))
                    d_734_v26_: _dafny.Seq
                    d_734_v26_ = _dafny.SeqWithoutIsStrInference([d_705_v1_, d_705_v1_])
                    d_735_v27_: _dafny.MultiSet
                    d_735_v27_ = _dafny.MultiSet([p2])
                    d_736_v28_: C0
                    nw131_ = C0()
                    nw131_.ctor__(d_734_v26_, (p2) not in (d_735_v27_))
                    d_736_v28_ = nw131_
                    d_737_v29_: _dafny.Seq
                    d_737_v29_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_738_v30_: _dafny.Seq
                    d_738_v30_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hnxddpbbf")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhsnvto")), p3])
                    d_739_v31_: D3
                    d_739_v31_ = D3_DC7((self).f20)
                    d_740_v32_: _dafny.Set
                    d_740_v32_ = _dafny.Set({(d_739_v31_).cf7})
                    d_741_v33_: D0
                    d_741_v33_ = D0_DC0(False)
                    d_742_v34_: D8
                    d_742_v34_ = D8_DC21(p3, d_706_v2_)
                    d_743_v35_: _dafny.Array
                    nw132_ = _dafny.Array(None, 28)
                    nw132_[int(0)] = not((self.f21) in (default__.fm32((0) - (len(d_737_v29_)), p2, d_707_v3_, globalState)))
                    nw132_[int(1)] = False
                    nw132_[int(2)] = (d_736_v28_).f23
                    nw132_[int(3)] = p1
                    nw132_[int(4)] = (p1) in (d_737_v29_)
                    nw132_[int(5)] = d_707_v3_
                    nw132_[int(6)] = (d_736_v28_).f23
                    nw132_[int(7)] = p0
                    nw132_[int(8)] = p0
                    nw132_[int(9)] = True
                    nw132_[int(10)] = True
                    nw132_[int(11)] = (d_736_v28_).f23
                    nw132_[int(12)] = (p2) >= (len(d_738_v30_))
                    nw132_[int(13)] = p1
                    nw132_[int(14)] = True
                    nw132_[int(15)] = p0
                    nw132_[int(16)] = not(p0)
                    nw132_[int(17)] = True
                    nw132_[int(18)] = p0
                    nw132_[int(19)] = (d_740_v32_).ispropersubset(d_740_v32_)
                    nw132_[int(20)] = d_707_v3_
                    nw132_[int(21)] = (self).fm16(d_741_v33_, d_707_v3_, globalState)
                    nw132_[int(22)] = (d_736_v28_).f23
                    nw132_[int(23)] = (d_706_v2_).issubset((d_742_v34_).cf35)
                    nw132_[int(24)] = (not(not(p0)) if p1 else (self).fm16(D0_DC0(d_707_v3_), p0, globalState))
                    nw132_[int(25)] = p0
                    nw132_[int(26)] = (d_736_v28_).f23
                    nw132_[int(27)] = p0
                    d_743_v35_ = nw132_
                    index135_ = default__.safeIndex(864, (d_743_v35_).length(0))
                    (d_743_v35_)[index135_] = (p1) and (p1)
                    index136_ = default__.safeIndex(864, (d_743_v35_).length(0))
                    (d_743_v35_)[index136_] = d_707_v3_
                    d_744_v36_: C1
                    nw133_ = C1()
                    nw133_.ctor__(((0) - ((self).f20)) + (self.f21))
                    d_744_v36_ = nw133_
                    pass
            pass
        rhs91_ = d_710_v6_
        rhs92_ = (0) - (self.f21)
        d_710_v6_ = rhs91_
        r0 = rhs92_
        if ((p2) + ((self).f20)) > ((self).f20):
            d_745_v37_: _dafny.Array
            nw134_ = _dafny.Array(_dafny.Set({}), 13)
            d_745_v37_ = nw134_
            index137_ = default__.safeIndex(515, (d_745_v37_).length(0))
            (d_745_v37_)[index137_] = d_706_v2_
            index138_ = default__.safeIndex(515, (d_745_v37_).length(0))
            (d_745_v37_)[index138_] = ((d_706_v2_ if p0 else d_706_v2_)).intersection(d_706_v2_)
            d_746_v38_: str
            d_746_v38_ = _dafny.CodePoint('y')
            d_746_v38_ = d_746_v38_
            d_706_v2_ = d_706_v2_
            if p0:
                (globalState).f6 = p2
                (globalState).f6 = p2
                d_747_v39_: D3
                d_747_v39_ = D3_DC7(p2)
                d_748_v40_: _dafny.Seq
                d_748_v40_ = _dafny.SeqWithoutIsStrInference([p2, p2])
                pat_let_tv16_ = d_748_v40_
                d_749_v41_: _dafny.Array
                nw135_ = _dafny.Array(None, 19)
                nw135_[int(0)] = d_747_v39_
                nw135_[int(1)] = d_747_v39_
                nw135_[int(2)] = d_747_v39_
                nw135_[int(3)] = d_747_v39_
                nw135_[int(4)] = d_747_v39_
                nw135_[int(5)] = d_747_v39_
                nw135_[int(6)] = d_747_v39_
                nw135_[int(7)] = d_747_v39_
                nw135_[int(8)] = d_747_v39_
                nw135_[int(9)] = D3_DC7(236)
                nw135_[int(10)] = d_747_v39_
                nw135_[int(11)] = d_747_v39_
                nw135_[int(12)] = d_747_v39_
                nw135_[int(13)] = D3_DC7((0) - (p2))
                nw135_[int(14)] = D3_DC7((self).f20)
                nw135_[int(15)] = d_747_v39_
                nw135_[int(16)] = d_747_v39_
                def iife92_(_pat_let28_0):
                    def iife93_(d_750_dt__update__tmp_h0_):
                        def iife94_(_pat_let29_0):
                            def iife95_(d_751_dt__update_hcf7_h0_):
                                return D3_DC7(d_751_dt__update_hcf7_h0_)
                            return iife95_(_pat_let29_0)
                        return iife94_((0) - (len(pat_let_tv16_)))
                    return iife93_(_pat_let28_0)
                nw135_[int(17)] = iife92_(d_747_v39_)
                nw135_[int(18)] = d_747_v39_
                d_749_v41_ = nw135_
                rhs93_ = (0) - ((p2) - ((p2) - (578)))
                rhs94_ = ((d_749_v41_ if p1 else d_749_v41_) if p1 else d_749_v41_)
                lhs80_ = globalState
                lhs80_.f6 = rhs93_
                d_749_v41_ = rhs94_
                d_752_v42_: _dafny.Seq
                d_752_v42_ = _dafny.SeqWithoutIsStrInference([d_705_v1_])
                d_753_v43_: C0
                nw136_ = C0()
                nw136_.ctor__(d_752_v42_, (d_746_v38_) in (p3))
                d_753_v43_ = nw136_
                d_754_v44_: _dafny.Map
                d_754_v44_ = _dafny.Map({self.f21: p1})
                d_755_v45_: _dafny.Array
                nw137_ = _dafny.Array(None, 3)
                nw137_[int(0)] = d_754_v44_
                nw137_[int(1)] = d_754_v44_
                nw137_[int(2)] = d_754_v44_
                d_755_v45_ = nw137_
                d_756_v46_: D0
                d_756_v46_ = D0_DC0(p1)
                d_757_v47_: _dafny.Map
                d_757_v47_ = _dafny.Map({False: (self).fm16(d_756_v46_, (d_753_v43_).f23, globalState)})
                d_758_v48_: _dafny.Array
                nw138_ = _dafny.Array(None, 19)
                nw138_[int(0)] = (d_748_v40_).set(default__.safeIndex(309, len(d_748_v40_)), (0) - (p2))
                nw138_[int(1)] = d_748_v40_
                nw138_[int(2)] = default__.fm32((self).f20, len(d_757_v47_), p0, globalState)
                nw138_[int(3)] = default__.fm32(self.f21, (self).f20, p1, globalState)
                nw138_[int(4)] = d_748_v40_
                nw138_[int(5)] = d_748_v40_
                nw138_[int(6)] = _dafny.SeqWithoutIsStrInference([(0) - (p2), self.f21, p2])
                nw138_[int(7)] = d_748_v40_
                nw138_[int(8)] = d_748_v40_
                nw138_[int(9)] = (default__.fm32(len(default__.fm32(self.f21, self.f21, p0, globalState)), self.f21, p1, globalState)).set(default__.safeIndex(485, len(default__.fm32(len(default__.fm32(self.f21, self.f21, p0, globalState)), self.f21, p1, globalState))), self.f21)
                nw138_[int(10)] = d_748_v40_
                nw138_[int(11)] = (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f20 for d_759_i2_ in range(default__.abs(307))]))])).set(default__.safeIndex((self).f20, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f20 for d_760_i2_ in range(default__.abs(307))]))]))), (d_747_v39_).cf7)
                nw138_[int(12)] = d_748_v40_
                nw138_[int(13)] = d_748_v40_
                nw138_[int(14)] = d_748_v40_
                nw138_[int(15)] = (d_748_v40_) + (d_748_v40_)
                nw138_[int(16)] = (d_748_v40_) + (_dafny.SeqWithoutIsStrInference([p2, (self).f20]))
                nw138_[int(17)] = _dafny.SeqWithoutIsStrInference([len((D8_DC21(_dafny.SeqWithoutIsStrInference([d_746_v38_ for d_761_i4_ in range(default__.abs(167))]), _dafny.Set({p0}))).cf34) for d_762_i3_ in range(default__.abs(920))])
                nw138_[int(18)] = d_748_v40_
                d_758_v48_ = nw138_
                index139_ = default__.safeIndex(257, (d_758_v48_).length(0))
                (d_758_v48_)[index139_] = (d_748_v40_) + (d_748_v40_)
                d_763_v49_: D9
                d_763_v49_ = D9_DC22(d_755_v45_)
                index140_ = default__.safeIndex(515, (d_745_v37_).length(0))
                index141_ = default__.safeIndex(257, (d_758_v48_).length(0))
                rhs95_ = (d_763_v49_).cf36
                rhs96_ = (d_745_v37_)[default__.safeIndex(515, (d_745_v37_).length(0))]
                rhs97_ = (d_748_v40_).set(default__.safeIndex(self.f21, len(d_748_v40_)), (self).f20)
                rhs98_ = p0
                lhs81_ = d_745_v37_
                lhs82_ = default__.safeIndex(515, (d_745_v37_).length(0))
                lhs83_ = d_758_v48_
                lhs84_ = default__.safeIndex(257, (d_758_v48_).length(0))
                lhs85_ = globalState
                d_755_v45_ = rhs95_
                lhs81_[lhs82_] = rhs96_
                lhs83_[lhs84_] = rhs97_
                lhs85_.f1 = rhs98_
            elif True:
                d_707_v3_ = ((p3) + (p3)) < (p3)
                d_764_v50_: _dafny.Seq
                d_764_v50_ = _dafny.SeqWithoutIsStrInference([d_705_v1_])
                d_765_v51_: C0
                nw139_ = C0()
                nw139_.ctor__(d_764_v50_, p0)
                d_765_v51_ = nw139_
                d_708_v4_ = (d_708_v4_).set(-112, (0) - (-286))
                d_766_v52_: _dafny.Array
                def lambda33_(d_767_p0_):
                    def lambda34_(d_768_i5_):
                        return d_767_p0_

                    return lambda34_

                init20_ = lambda33_(p0)
                nw140_ = _dafny.Array(None, 7)
                for i0_20_ in range(nw140_.length(0)):
                    nw140_[i0_20_] = init20_(i0_20_)
                d_766_v52_ = nw140_
                index142_ = default__.safeIndex(459, (d_766_v52_).length(0))
                (d_766_v52_)[index142_] = p0
                index143_ = default__.safeIndex(459, (d_766_v52_).length(0))
                (d_766_v52_)[index143_] = (d_765_v51_).f23
                d_769_v53_: _dafny.Set
                d_769_v53_ = _dafny.Set({(self).f20})
                d_770_v54_: _dafny.Map
                d_770_v54_ = _dafny.Map({p3: d_769_v53_})
                d_771_v55_: _dafny.Seq
                d_771_v55_ = _dafny.SeqWithoutIsStrInference([(d_765_v51_).f23, (d_766_v52_)[default__.safeIndex(459, (d_766_v52_).length(0))], False, (d_765_v51_).f23, True])
                d_772_v56_: _dafny.Seq
                d_772_v56_ = _dafny.SeqWithoutIsStrInference([len(d_771_v55_), (self).f20])
                d_770_v54_ = (d_770_v54_).set((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tff"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xubovus"))), _dafny.Set({(d_772_v56_)[default__.safeIndex(p2, len(d_772_v56_))], p2}))
            d_773_v57_: _dafny.Seq
            d_773_v57_ = _dafny.SeqWithoutIsStrInference([d_705_v1_])
            d_774_v58_: C0
            nw141_ = C0()
            nw141_.ctor__(d_773_v57_, p0)
            d_774_v58_ = nw141_
            d_775_v59_: _dafny.Array
            nw142_ = _dafny.Array(_dafny.Array(None, 0), 13)
            d_775_v59_ = nw142_
            d_776_v60_: D5
            d_776_v60_ = D5_DC10(d_775_v59_)
            d_777_v61_: D7
            d_777_v61_ = D7_DC17(p0, p0, d_774_v58_, d_776_v60_, d_707_v3_)
            d_707_v3_ = (d_777_v61_).cf28
        elif True:
            if d_707_v3_:
                (globalState).f1 = not(p1)
                d_778_v62_: _dafny.Array
                nw143_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_778_v62_ = nw143_
                d_779_v63_: _dafny.MultiSet
                d_779_v63_ = _dafny.MultiSet([self.f21])
                d_780_v64_: T0
                nw144_ = C1()
                nw144_.ctor__((d_779_v63_).cardinality)
                d_780_v64_ = nw144_
                d_781_v65_: D7
                d_781_v65_ = D7_DC15(d_780_v64_)
                pat_let_tv17_ = d_780_v64_
                d_782_v66_: _dafny.Array
                nw145_ = _dafny.Array(None, 12)
                nw145_[int(0)] = D7_DC15(d_780_v64_)
                nw145_[int(1)] = D7_DC15(d_780_v64_)
                nw145_[int(2)] = d_781_v65_
                nw145_[int(3)] = d_781_v65_
                nw145_[int(4)] = D7_DC15(d_780_v64_)
                nw145_[int(5)] = d_781_v65_
                nw145_[int(6)] = d_781_v65_
                nw145_[int(7)] = d_781_v65_
                def iife96_(_pat_let30_0):
                    def iife97_(d_783_dt__update__tmp_h1_):
                        def iife98_(_pat_let31_0):
                            def iife99_(d_784_dt__update_hcf20_h0_):
                                return D7_DC15(d_784_dt__update_hcf20_h0_)
                            return iife99_(_pat_let31_0)
                        return iife98_(pat_let_tv17_)
                    return iife97_(_pat_let30_0)
                nw145_[int(8)] = iife96_(d_781_v65_)
                nw145_[int(9)] = d_781_v65_
                nw145_[int(10)] = d_781_v65_
                nw145_[int(11)] = d_781_v65_
                d_782_v66_ = nw145_
                d_785_v67_: _dafny.Array
                nw146_ = _dafny.Array(None, 2)
                nw146_[int(0)] = d_782_v66_
                nw146_[int(1)] = d_782_v66_
                d_785_v67_ = nw146_
                index144_ = default__.safeIndex(403, (d_778_v62_).length(0))
                (d_778_v62_)[index144_] = d_785_v67_
                d_786_v68_: _dafny.Array
                nw147_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                d_786_v68_ = nw147_
                d_787_v69_: str
                d_787_v69_ = _dafny.CodePoint('w')
                index145_ = default__.safeIndex(123, (d_786_v68_).length(0))
                (d_786_v68_)[index145_] = d_787_v69_
                d_788_v70_: D10
                d_788_v70_ = D10_DC26(d_785_v67_)
                pat_let_tv18_ = d_785_v67_
                index146_ = default__.safeIndex(403, (d_778_v62_).length(0))
                index147_ = default__.safeIndex(123, (d_786_v68_).length(0))
                def iife100_(_pat_let32_0):
                    def iife101_(d_789_dt__update__tmp_h2_):
                        def iife102_(_pat_let33_0):
                            def iife103_(d_790_dt__update_hcf45_h0_):
                                return D10_DC26(d_790_dt__update_hcf45_h0_)
                            return iife103_(_pat_let33_0)
                        return iife102_(pat_let_tv18_)
                    return iife101_(_pat_let32_0)
                rhs99_ = (iife100_(d_788_v70_)).cf45
                rhs100_ = d_787_v69_
                rhs101_ = p0
                lhs86_ = d_778_v62_
                lhs87_ = default__.safeIndex(403, (d_778_v62_).length(0))
                lhs88_ = d_786_v68_
                lhs89_ = default__.safeIndex(123, (d_786_v68_).length(0))
                lhs90_ = globalState
                lhs86_[lhs87_] = rhs99_
                lhs88_[lhs89_] = rhs100_
                lhs90_.f1 = rhs101_
                (globalState).f1 = p0
                (globalState).f1 = (True) or (True)
                (self).f21 = default__.safeModuloInt((default__.fm1(p1, globalState)) * (default__.fm1(p0, globalState)), default__.safeModuloInt(self.f21, self.f21))
            elif True:
                d_791_v71_: _dafny.Seq
                d_791_v71_ = _dafny.SeqWithoutIsStrInference([d_705_v1_, d_705_v1_])
                d_792_v72_: D11
                d_792_v72_ = D11_DC30(d_791_v71_)
                d_793_v73_: C0
                nw148_ = C0()
                nw148_.ctor__((d_792_v72_).cf53, p1)
                d_793_v73_ = nw148_
                d_794_v74_: _dafny.MultiSet
                d_794_v74_ = _dafny.MultiSet([_dafny.Map({self.f21: (self).f20})])
                d_795_v75_: _dafny.Seq
                d_795_v75_ = _dafny.SeqWithoutIsStrInference([(self).f20, self.f21])
                d_796_v76_: _dafny.MultiSet
                d_796_v76_ = _dafny.MultiSet([d_795_v75_, default__.fm32((self).f20, (self).f20, p0, globalState)])
                d_797_v77_: _dafny.Array
                nw149_ = _dafny.Array(False, 14)
                d_797_v77_ = nw149_
                index148_ = default__.safeIndex(878, (d_797_v77_).length(0))
                (d_797_v77_)[index148_] = (_dafny.CodePoint('t')) not in (p3)
                d_798_v79_: D12
                def iife104_():
                    coll36_ = _dafny.Map()
                    compr_36_: int
                    for compr_36_ in _dafny.IntegerRange(-882, 664):
                        d_799_v78_: int = compr_36_
                        if ((-882) <= (d_799_v78_)) and ((d_799_v78_) < (664)):
                            coll36_[(d_799_v78_) * ((0) - (p2))] = (self).f20
                    return _dafny.Map(coll36_)
                d_798_v79_ = D12_DC34(iife104_()
)
                d_800_v80_: _dafny.Map
                d_800_v80_ = _dafny.Map({(d_793_v73_).f23: (d_793_v73_).f23})
                d_801_v81_: str
                d_801_v81_ = _dafny.CodePoint('a')
                d_802_v82_: _dafny.Map
                d_802_v82_ = _dafny.Map({d_801_v81_: d_796_v76_})
                d_803_v83_: D11
                d_803_v83_ = D11_DC32(p0, self.f21, d_707_v3_)
                d_804_v84_: D7
                d_804_v84_ = D7_DC16(p1, (d_803_v83_).cf60, p0)
                pat_let_tv19_ = d_708_v4_
                index149_ = default__.safeIndex(878, (d_797_v77_).length(0))
                def iife105_(_pat_let34_0):
                    def iife106_(d_805_dt__update__tmp_h3_):
                        def iife107_(_pat_let35_0):
                            def iife108_(d_806_dt__update_hcf63_h0_):
                                return D12_DC34(d_806_dt__update_hcf63_h0_)
                            return iife108_(_pat_let35_0)
                        return iife107_(pat_let_tv19_)
                    return iife106_(_pat_let34_0)
                rhs102_ = (d_794_v74_) - (_dafny.MultiSet([(iife105_(d_798_v79_)).cf63]))
                rhs103_ = ((p0) in (d_800_v80_) if p1 else (d_707_v3_ if p0 else (d_793_v73_).f23))
                rhs104_ = ((d_796_v76_).set(d_795_v75_, default__.abs((0) - (self.f21)))) - (((d_802_v82_)[d_801_v81_] if (d_801_v81_) in (d_802_v82_) else d_796_v76_))
                rhs105_ = (d_804_v84_).cf23
                lhs91_ = globalState
                lhs92_ = d_797_v77_
                lhs93_ = default__.safeIndex(878, (d_797_v77_).length(0))
                d_794_v74_ = rhs102_
                lhs91_.f1 = rhs103_
                d_796_v76_ = rhs104_
                lhs92_[lhs93_] = rhs105_
                d_807_v85_: _dafny.MultiSet
                d_807_v85_ = _dafny.MultiSet([p2, len(default__.fm33((0) - (len(d_706_v2_)), (self).f20, True, globalState)), p2, self.f21, p2])
                d_808_v87_: D0
                d_808_v87_ = D0_DC1((self).f20)
                d_809_v88_: D0
                d_809_v88_ = D0_DC2(d_808_v87_)
                d_810_v89_: D0
                d_810_v89_ = D0_DC2((d_809_v88_).cf2)
                pat_let_tv20_ = d_808_v87_
                pat_let_tv21_ = d_808_v87_
                d_811_v90_: _dafny.Array
                nw150_ = _dafny.Array(None, 19)
                def iife109_(_pat_let36_0):
                    def iife110_(d_812_dt__update__tmp_h4_):
                        def iife111_(_pat_let37_0):
                            def iife112_(d_813_dt__update_hcf2_h0_):
                                return D0_DC2(d_813_dt__update_hcf2_h0_)
                            return iife112_(_pat_let37_0)
                        return iife111_(pat_let_tv20_)
                    return iife110_(_pat_let36_0)
                nw150_[int(0)] = iife109_(d_809_v88_)
                nw150_[int(1)] = D0_DC2(d_808_v87_)
                nw150_[int(2)] = d_809_v88_
                nw150_[int(3)] = d_810_v89_
                nw150_[int(4)] = d_810_v89_
                nw150_[int(5)] = d_809_v88_
                nw150_[int(6)] = d_810_v89_
                nw150_[int(7)] = d_809_v88_
                nw150_[int(8)] = d_809_v88_
                nw150_[int(9)] = d_809_v88_
                nw150_[int(10)] = d_809_v88_
                nw150_[int(11)] = d_809_v88_
                nw150_[int(12)] = d_810_v89_
                nw150_[int(13)] = d_809_v88_
                def iife113_(_pat_let38_0):
                    def iife114_(d_814_dt__update__tmp_h5_):
                        def iife115_(_pat_let39_0):
                            def iife116_(d_815_dt__update_hcf2_h1_):
                                return D0_DC2(d_815_dt__update_hcf2_h1_)
                            return iife116_(_pat_let39_0)
                        return iife115_(pat_let_tv21_)
                    return iife114_(_pat_let38_0)
                nw150_[int(14)] = iife113_(d_809_v88_)
                nw150_[int(15)] = d_809_v88_
                nw150_[int(16)] = D0_DC2(d_808_v87_)
                nw150_[int(17)] = d_810_v89_
                nw150_[int(18)] = d_810_v89_
                d_811_v90_ = nw150_
                def iife117_():
                    coll37_ = _dafny.Set()
                    compr_37_: int
                    for compr_37_ in (d_807_v85_).Elements:
                        d_816_v86_: int = compr_37_
                        if (d_816_v86_) in (d_807_v85_):
                            coll37_ = coll37_.union(_dafny.Set([(d_816_v86_) - (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_817_i6_ in range(default__.abs(-153))]))})))]))
                    return _dafny.Set(coll37_)
                default__.m0(iife117_()
                , d_811_v90_, d_704_v0_, p3, globalState)
                d_818_v91_: _dafny.Map
                d_818_v91_ = _dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iug"))})
                d_819_v92_: D3
                d_819_v92_ = D3_DC7(self.f21)
                d_820_v93_: _dafny.Map
                d_820_v93_ = _dafny.Map({d_797_v77_: True})
                d_821_v94_: D10
                d_821_v94_ = D10_DC27(len(d_800_v80_), p3, p2, d_819_v92_, d_820_v93_)
                d_822_v95_: D3
                d_822_v95_ = D3_DC8(d_795_v75_, (d_821_v94_).cf47, p0, d_801_v81_)
                index150_ = default__.safeIndex(878, (d_797_v77_).length(0))
                (d_797_v77_)[index150_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wefcbtsh"))) < ((p3) + (((d_818_v91_)[(d_793_v73_).f23] if ((d_793_v73_).f23) in (d_818_v91_) else (d_822_v95_).cf9)))
                d_823_v96_: C1
                nw151_ = C1()
                nw151_.ctor__((self).f20)
                d_823_v96_ = nw151_
            pat_let_tv22_ = d_704_v0_
            pat_let_tv23_ = d_704_v0_
            pat_let_tv24_ = d_704_v0_
            d_824_v97_: _dafny.Array
            nw152_ = _dafny.Array(None, 11)
            nw152_[int(0)] = d_705_v1_
            nw152_[int(1)] = d_705_v1_
            def iife118_(_pat_let40_0):
                def iife119_(d_825_dt__update__tmp_h6_):
                    def iife120_(_pat_let41_0):
                        def iife121_(d_826_dt__update_hcf6_h0_):
                            return D3_DC6(d_826_dt__update_hcf6_h0_)
                        return iife121_(_pat_let41_0)
                    return iife120_(pat_let_tv22_)
                return iife119_(_pat_let40_0)
            nw152_[int(2)] = iife118_(d_705_v1_)
            nw152_[int(3)] = D3_DC6(d_704_v0_)
            nw152_[int(4)] = d_705_v1_
            def iife122_(_pat_let42_0):
                def iife123_(d_827_dt__update__tmp_h7_):
                    def iife124_(_pat_let43_0):
                        def iife125_(d_828_dt__update_hcf6_h1_):
                            return D3_DC6(d_828_dt__update_hcf6_h1_)
                        return iife125_(_pat_let43_0)
                    return iife124_(pat_let_tv23_)
                return iife123_(_pat_let42_0)
            nw152_[int(5)] = iife122_(d_705_v1_)
            nw152_[int(6)] = d_705_v1_
            def iife126_(_pat_let44_0):
                def iife127_(d_829_dt__update__tmp_h8_):
                    def iife128_(_pat_let45_0):
                        def iife129_(d_830_dt__update_hcf6_h2_):
                            return D3_DC6(d_830_dt__update_hcf6_h2_)
                        return iife129_(_pat_let45_0)
                    return iife128_(pat_let_tv24_)
                return iife127_(_pat_let44_0)
            nw152_[int(7)] = iife126_(d_705_v1_)
            nw152_[int(8)] = d_705_v1_
            nw152_[int(9)] = d_705_v1_
            nw152_[int(10)] = d_705_v1_
            d_824_v97_ = nw152_
            d_831_v98_: _dafny.Set
            d_831_v98_ = _dafny.Set({d_824_v97_, d_824_v97_, d_824_v97_, d_824_v97_})
            d_832_v99_: D9
            d_832_v99_ = D9_DC23(d_831_v98_, p3, default__.fm2(globalState), d_704_v0_, default__.fm29(d_707_v3_, (self).f20, (self).f20, p0, globalState))
            d_833_v100_: _dafny.Seq
            d_833_v100_ = _dafny.SeqWithoutIsStrInference([not(False), d_707_v3_])
            d_834_v101_: _dafny.MultiSet
            d_834_v101_ = _dafny.MultiSet([(self).f20, -776])
            d_835_v102_: _dafny.MultiSet
            d_835_v102_ = _dafny.MultiSet([p0, default__.fm2(globalState)])
            d_836_v103_: D3
            d_836_v103_ = D3_DC7(((d_834_v101_)[p2] if (p2) in (d_834_v101_) else (d_835_v102_).cardinality))
            d_837_v104_: _dafny.Array
            def lambda35_(d_838_p0_):
                def lambda36_(d_839_i7_):
                    return d_838_p0_

                return lambda36_

            init21_ = lambda35_(p0)
            nw153_ = _dafny.Array(None, 6)
            for i0_21_ in range(nw153_.length(0)):
                nw153_[i0_21_] = init21_(i0_21_)
            d_837_v104_ = nw153_
            d_840_v105_: _dafny.Map
            d_840_v105_ = _dafny.Map({d_837_v104_: p1})
            d_841_v106_: D10
            d_841_v106_ = D10_DC27((len(_dafny.Map({d_707_v3_: not(p0)}))) - ((0) - (self.f21)), (p3) + ((d_832_v99_).cf38), (len(d_833_v100_)) + (self.f21), d_836_v103_, d_840_v105_)
            d_841_v106_ = D10_DC27((964 if p0 else (self).f20), p3, (self.f21) * ((self).f20), d_836_v103_, _dafny.Map({d_837_v104_: p1}))
            r0 = default__.safeModuloInt(342, (self.f21) * ((self).f20))
            d_706_v2_ = d_706_v2_
            index151_ = default__.safeIndex(0, (d_837_v104_).length(0))
            (d_837_v104_)[index151_] = p1
            index152_ = default__.safeIndex(0, (d_837_v104_).length(0))
            (d_837_v104_)[index152_] = False
        d_842_v107_: _dafny.Array
        def lambda37_(d_843_p3_):
            def lambda38_(d_844_i8_):
                return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_845_i9_ in range(default__.abs(370))])) + (d_843_p3_)

            return lambda38_

        init22_ = lambda37_(p3)
        nw154_ = _dafny.Array(None, 15)
        for i0_22_ in range(nw154_.length(0)):
            nw154_[i0_22_] = init22_(i0_22_)
        d_842_v107_ = nw154_
        index153_ = default__.safeIndex(868, (d_842_v107_).length(0))
        (d_842_v107_)[index153_] = p3
        d_846_v108_: _dafny.MultiSet
        d_846_v108_ = _dafny.MultiSet([self.f21])
        index154_ = default__.safeIndex(868, (d_842_v107_).length(0))
        rhs106_ = p3
        rhs107_ = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f21]))).issubset(d_846_v108_)
        rhs108_ = d_707_v3_
        lhs94_ = d_842_v107_
        lhs95_ = default__.safeIndex(868, (d_842_v107_).length(0))
        lhs96_ = globalState
        lhs94_[lhs95_] = rhs106_
        d_707_v3_ = rhs107_
        lhs96_.f1 = rhs108_
        r0 = (self).f20
        return r0

    def m15(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        r2: _dafny.Array = _dafny.Array(None, 0)
        if p1:
            (globalState).f6 = (self).fm15(globalState)
            d_847_v0_: _dafny.Seq
            d_847_v0_ = _dafny.SeqWithoutIsStrInference([default__.fm29(p1, p3, self.f21, p1, globalState)])
            d_848_v1_: D10
            d_848_v1_ = D10_DC28(p3)
            d_849_v2_: _dafny.Map
            d_849_v2_ = _dafny.Map({True: d_848_v1_})
            d_850_v3_: _dafny.Seq
            d_850_v3_ = _dafny.SeqWithoutIsStrInference([d_849_v2_, d_849_v2_])
            d_851_v4_: _dafny.Array
            nw155_ = _dafny.Array(None, 19)
            nw155_[int(0)] = _dafny.Map({True: D10_DC28(len(d_847_v0_))})
            nw155_[int(1)] = d_849_v2_
            nw155_[int(2)] = d_849_v2_
            nw155_[int(3)] = d_849_v2_
            nw155_[int(4)] = d_849_v2_
            nw155_[int(5)] = d_849_v2_
            nw155_[int(6)] = d_849_v2_
            nw155_[int(7)] = d_849_v2_
            nw155_[int(8)] = _dafny.Map({p1: d_848_v1_})
            nw155_[int(9)] = d_849_v2_
            nw155_[int(10)] = (d_850_v3_)[default__.safeIndex(837, len(d_850_v3_))]
            nw155_[int(11)] = d_849_v2_
            nw155_[int(12)] = d_849_v2_
            nw155_[int(13)] = _dafny.Map({p1: d_848_v1_})
            nw155_[int(14)] = d_849_v2_
            nw155_[int(15)] = d_849_v2_
            nw155_[int(16)] = d_849_v2_
            nw155_[int(17)] = d_849_v2_
            nw155_[int(18)] = _dafny.Map({p1: d_848_v1_})
            d_851_v4_ = nw155_
            d_852_v5_: _dafny.Map
            d_852_v5_ = _dafny.Map({len(_dafny.Map({(self).f20: p1})): d_851_v4_})
            d_852_v5_ = (d_852_v5_).set((0) - (p2), d_851_v4_)
            d_853_v6_: C1
            nw156_ = C1()
            nw156_.ctor__((self).f20)
            d_853_v6_ = nw156_
            (globalState).f6 = (self).f20
            d_854_v7_: _dafny.MultiSet
            d_854_v7_ = _dafny.MultiSet([p3])
            d_855_v8_: _dafny.Seq
            d_855_v8_ = _dafny.SeqWithoutIsStrInference([((d_854_v7_)[857] if (857) in (d_854_v7_) else self.f21)])
            (globalState).f6 = (d_855_v8_)[default__.safeIndex(self.f21, len(d_855_v8_))]
        elif True:
            d_856_v9_: _dafny.Seq
            d_856_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkn"))
            (globalState).f6 = len((default__.fm21(globalState)) + (d_856_v9_))
            d_857_v10_: _dafny.Map
            d_857_v10_ = _dafny.Map({p1: self.f21})
            d_857_v10_ = (d_857_v10_).set(default__.fm2(globalState), (p3) - (p2))
            d_858_v11_: D3
            d_858_v11_ = D3_DC7(12)
            d_859_v12_: str
            d_859_v12_ = _dafny.CodePoint('y')
            d_860_v13_: _dafny.Array
            nw157_ = _dafny.Array(None, 22)
            nw157_[int(0)] = d_856_v9_
            nw157_[int(1)] = (((d_856_v9_).set(default__.safeIndex((d_858_v11_).cf7, len(d_856_v9_)), d_859_v12_)).set(default__.safeIndex(p3, len((d_856_v9_).set(default__.safeIndex((d_858_v11_).cf7, len(d_856_v9_)), d_859_v12_))), d_859_v12_) if True else d_856_v9_)
            nw157_[int(2)] = d_856_v9_
            nw157_[int(3)] = d_856_v9_
            nw157_[int(4)] = d_856_v9_
            nw157_[int(5)] = (d_856_v9_) + (d_856_v9_)
            nw157_[int(6)] = d_856_v9_
            nw157_[int(7)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))) + (d_856_v9_)
            nw157_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hnastust"))
            nw157_[int(9)] = d_856_v9_
            nw157_[int(10)] = (d_856_v9_) + (d_856_v9_)
            nw157_[int(11)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsv"))) + (d_856_v9_)
            nw157_[int(12)] = (default__.fm21(globalState)) + (d_856_v9_)
            nw157_[int(13)] = (d_856_v9_ if p1 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtgqkrddi")))
            nw157_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nflhvqi"))
            nw157_[int(15)] = d_856_v9_
            nw157_[int(16)] = d_856_v9_
            nw157_[int(17)] = d_856_v9_
            nw157_[int(18)] = d_856_v9_
            nw157_[int(19)] = (d_856_v9_) + ((_dafny.SeqWithoutIsStrInference([d_859_v12_ for d_861_i0_ in range(default__.abs(958))])).set(default__.safeIndex(p3, len(_dafny.SeqWithoutIsStrInference([d_859_v12_ for d_862_i0_ in range(default__.abs(958))]))), d_859_v12_))
            nw157_[int(20)] = ((d_856_v9_) + (d_856_v9_)).set(default__.safeIndex((self).f20, len((d_856_v9_) + (d_856_v9_))), d_859_v12_)
            nw157_[int(21)] = (d_856_v9_).set(default__.safeIndex(p3, len(d_856_v9_)), _dafny.CodePoint('t'))
            d_860_v13_ = nw157_
            d_860_v13_ = d_860_v13_
            d_863_v14_: _dafny.Seq
            d_863_v14_ = _dafny.SeqWithoutIsStrInference([True])
            d_864_v15_: _dafny.Map
            d_864_v15_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p2, len(d_863_v14_), (self).fm15(globalState), 191, -819]): not(default__.fm2(globalState))})
            d_865_v16_: _dafny.Seq
            d_865_v16_ = _dafny.SeqWithoutIsStrInference([p1, ((d_864_v15_)[_dafny.SeqWithoutIsStrInference([p3])] if (_dafny.SeqWithoutIsStrInference([p3])) in (d_864_v15_) else p1), not(p1)])
            (globalState).f1 = (d_863_v14_)[default__.safeIndex(157, len(d_863_v14_))]
            (self).f21 = (0) - (((0) - (-411)) + (p2))
        d_866_v18_: C1
        nw158_ = C1()
        def iife130_():
            coll38_ = _dafny.Set()
            compr_38_: int
            for compr_38_ in _dafny.IntegerRange(-8, -685):
                d_867_v17_: int = compr_38_
                if ((-8) <= (d_867_v17_)) and ((d_867_v17_) < (-685)):
                    coll38_ = coll38_.union(_dafny.Set([(d_867_v17_) + ((self).f20)]))
            return _dafny.Set(coll38_)
        nw158_.ctor__(len(iife130_()
        ))
        d_866_v18_ = nw158_
        (globalState).f1 = p1
        d_868_v19_: _dafny.Set
        d_868_v19_ = _dafny.Set({p1})
        d_869_v20_: _dafny.Map
        d_869_v20_ = _dafny.Map({(d_866_v18_).f24: (d_868_v19_).intersection(d_868_v19_)})
        d_869_v20_ = _dafny.Map({default__.safeDivisionInt(p3, (self).f20): (d_868_v19_).intersection(d_868_v19_)})
        d_870_v21_: _dafny.Seq
        d_870_v21_ = _dafny.SeqWithoutIsStrInference([p1, p1, False, p1])
        d_871_v22_: _dafny.Array
        nw159_ = _dafny.Array(False, 25)
        d_871_v22_ = nw159_
        d_872_v23_: _dafny.MultiSet
        d_872_v23_ = _dafny.MultiSet([d_871_v22_])
        if (d_870_v21_)[default__.safeIndex((d_872_v23_).cardinality, len(d_870_v21_))]:
            d_873_v24_: _dafny.Map
            d_873_v24_ = _dafny.Map({p1: not (p1) or (p1)})
            d_874_v25_: _dafny.Map
            d_874_v25_ = _dafny.Map({p3: p1})
            d_875_v26_: D7
            d_875_v26_ = D7_DC16(p1, self.f21, p1)
            d_876_v27_: _dafny.Seq
            d_876_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcmsmlo"))
            d_877_v28_: _dafny.Map
            d_877_v28_ = _dafny.Map({d_875_v26_: d_876_v27_})
            d_878_v29_: D13
            d_878_v29_ = D13_DC36(d_877_v28_)
            d_873_v24_ = (d_873_v24_).set(((d_874_v25_)[(0) - (len((d_878_v29_).cf64))] if ((0) - (len((d_878_v29_).cf64))) in (d_874_v25_) else p1), p1)
            index155_ = default__.safeIndex(740, (d_871_v22_).length(0))
            (d_871_v22_)[index155_] = p1
            index156_ = default__.safeIndex(740, (d_871_v22_).length(0))
            (d_871_v22_)[index156_] = p1
            d_879_v30_: _dafny.Seq
            d_879_v30_ = _dafny.SeqWithoutIsStrInference([self.f21, (self).f20, p2])
            d_880_v31_: _dafny.Seq
            d_880_v31_ = _dafny.SeqWithoutIsStrInference([d_873_v24_])
            d_881_v32_: _dafny.Map
            d_881_v32_ = _dafny.Map({(d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))]: (d_866_v18_).f24})
            d_882_v33_: _dafny.Seq
            d_882_v33_ = d_879_v30_
            d_883_v34_: _dafny.Map
            d_883_v34_ = _dafny.Map({(d_880_v31_)[default__.safeIndex(((d_881_v32_)[(d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))]] if ((d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))]) in (d_881_v32_) else self.f21), len(d_880_v31_))]: (d_882_v33_)})
            d_884_v35_: str
            d_884_v35_ = _dafny.CodePoint('x')
            d_885_v36_: _dafny.Array
            nw160_ = _dafny.Array(None, 14)
            nw160_[int(0)] = (d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))]
            nw160_[int(1)] = (d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))]
            nw160_[int(2)] = (d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))]
            nw160_[int(3)] = (d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))]
            nw160_[int(4)] = (d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))]
            nw160_[int(5)] = p1
            nw160_[int(6)] = (d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))]
            nw160_[int(7)] = (d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))]
            nw160_[int(8)] = (d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))]
            nw160_[int(9)] = p1
            nw160_[int(10)] = p1
            nw160_[int(11)] = (d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))]
            nw160_[int(12)] = p1
            nw160_[int(13)] = p1
            d_885_v36_ = nw160_
            d_886_v37_: _dafny.Set
            d_886_v37_ = _dafny.Set({d_885_v36_, d_885_v36_})
            d_887_v38_: _dafny.Map
            d_887_v38_ = _dafny.Map({len(d_886_v37_): (d_866_v18_).f24})
            d_879_v30_ = ((d_883_v34_)[(_dafny.Map({False: not((d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))])})) | (default__.fm34(p2, p1, default__.fm1((d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))], globalState), d_884_v35_, globalState))] if ((_dafny.Map({False: not((d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))])})) | (default__.fm34(p2, p1, default__.fm1((d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))], globalState), d_884_v35_, globalState))) in (d_883_v34_) else _dafny.SeqWithoutIsStrInference([len(d_873_v24_), (0) - (p2), len(d_887_v38_), len(d_868_v19_)]))
            d_888_v39_: _dafny.Array
            nw161_ = _dafny.Array(D7.default()(), 28)
            d_888_v39_ = nw161_
            d_889_v40_: _dafny.Map
            d_889_v40_ = _dafny.Map({default__.fm29((d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))], self.f21, p2, (d_871_v22_)[default__.safeIndex(740, (d_871_v22_).length(0))], globalState): d_885_v36_})
            d_890_v41_: _dafny.Map
            d_890_v41_ = _dafny.Map({d_889_v40_: d_888_v39_})
            d_888_v39_ = ((d_890_v41_)[d_889_v40_] if (d_889_v40_) in (d_890_v41_) else d_888_v39_)
            index157_ = default__.safeIndex(740, (d_871_v22_).length(0))
            (d_871_v22_)[index157_] = p1
        elif True:
            (globalState).f1 = p1
            d_891_v42_: _dafny.Map
            d_891_v42_ = _dafny.Map({p2: p1})
            d_892_v43_: _dafny.Map
            d_892_v43_ = _dafny.Map({d_891_v42_: p1})
            if (p1 if p1 else (d_866_v18_).fm24(p1, d_892_v43_, globalState)):
                d_893_v44_: _dafny.Array
                nw162_ = _dafny.Array(D3.default()(), 13)
                d_893_v44_ = nw162_
                d_894_v45_: _dafny.Set
                d_894_v45_ = _dafny.Set({d_893_v44_})
                d_895_v46_: _dafny.Seq
                d_895_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))
                d_896_v47_: _dafny.MultiSet
                d_896_v47_ = _dafny.MultiSet([p1])
                d_897_v48_: _dafny.Map
                d_897_v48_ = _dafny.Map({p1: d_896_v47_})
                d_898_v49_: str
                d_898_v49_ = _dafny.CodePoint('j')
                d_899_v50_: _dafny.Array
                def lambda39_(d_900_v18_):
                    def lambda40_(d_901_i1_):
                        return default__.safeDivisionInt(d_901_i1_, (d_900_v18_).f24)

                    return lambda40_

                init23_ = lambda39_(d_866_v18_)
                nw163_ = _dafny.Array(None, 23)
                for i0_23_ in range(nw163_.length(0)):
                    nw163_[i0_23_] = init23_(i0_23_)
                d_899_v50_ = nw163_
                rhs109_ = p1
                rhs110_ = (D9_DC23(d_894_v45_, (d_895_v46_).set(default__.safeIndex((((d_897_v48_)[p1] if (p1) in (d_897_v48_) else d_896_v47_)).cardinality, len(d_895_v46_)), d_898_v49_), p1, d_899_v50_, d_898_v49_)).cf40
                rhs111_ = (d_866_v18_).f24
                lhs97_ = globalState
                lhs98_ = globalState
                r0 = rhs109_
                lhs97_.f9 = rhs110_
                lhs98_.f6 = rhs111_
                d_902_v51_: _dafny.MultiSet
                d_902_v51_ = _dafny.MultiSet([(0) - ((d_866_v18_).f24), (d_866_v18_).f24])
                d_903_v52_: _dafny.Map
                d_903_v52_ = _dafny.Map({p1: (0) - ((self).f20)})
                d_904_v53_: _dafny.Seq
                d_904_v53_ = _dafny.SeqWithoutIsStrInference([self.f21])
                d_905_v54_: _dafny.Seq
                d_905_v54_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(d_866_v18_).f24, self.f21])), (d_866_v18_).f24]), d_902_v51_, d_902_v51_])
                d_906_v55_: _dafny.Map
                d_906_v55_ = _dafny.Map({d_898_v49_: p2})
                d_907_v56_: _dafny.Array
                nw164_ = _dafny.Array(None, 23)
                nw164_[int(0)] = (d_902_v51_).set(len(d_903_v52_), default__.abs((D0_DC1(len(_dafny.SeqWithoutIsStrInference([d_898_v49_ for d_908_i2_ in range(default__.abs(738))])))).cf1))
                nw164_[int(1)] = (_dafny.MultiSet([(d_866_v18_).f24, (self).f20])).set((d_866_v18_).f24, default__.abs((d_866_v18_).f24))
                nw164_[int(2)] = d_902_v51_
                nw164_[int(3)] = (d_902_v51_).intersection(_dafny.MultiSet(d_904_v53_))
                nw164_[int(4)] = (d_902_v51_) - (_dafny.MultiSet([69]))
                nw164_[int(5)] = d_902_v51_
                nw164_[int(6)] = d_902_v51_
                nw164_[int(7)] = _dafny.MultiSet([(0) - (self.f21)])
                nw164_[int(8)] = _dafny.MultiSet([(d_866_v18_).f24, len(d_895_v46_)])
                nw164_[int(9)] = ((d_905_v54_)[default__.safeIndex((d_866_v18_).f24, len(d_905_v54_))]) - (d_902_v51_)
                nw164_[int(10)] = _dafny.MultiSet([609, self.f21, p3])
                nw164_[int(11)] = d_902_v51_
                nw164_[int(12)] = d_902_v51_
                nw164_[int(13)] = (d_902_v51_) - (d_902_v51_)
                nw164_[int(14)] = d_902_v51_
                nw164_[int(15)] = d_902_v51_
                nw164_[int(16)] = _dafny.MultiSet([(d_866_v18_).f24, p3])
                nw164_[int(17)] = d_902_v51_
                nw164_[int(18)] = ((d_902_v51_).set(len(d_906_v55_), default__.abs(p2))) | (d_902_v51_)
                nw164_[int(19)] = d_902_v51_
                nw164_[int(20)] = d_902_v51_
                nw164_[int(21)] = d_902_v51_
                nw164_[int(22)] = d_902_v51_
                d_907_v56_ = nw164_
                index158_ = default__.safeIndex(843, (d_907_v56_).length(0))
                (d_907_v56_)[index158_] = _dafny.MultiSet([(self).f20, (d_866_v18_).f24, len(d_868_v19_), (d_866_v18_).f24, len(d_895_v46_)])
                d_909_v57_: _dafny.Seq
                d_909_v57_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f20: p1}), d_891_v42_])
                index159_ = default__.safeIndex(428, (d_871_v22_).length(0))
                (d_871_v22_)[index159_] = (_dafny.SeqWithoutIsStrInference([d_891_v42_ for d_910_i3_ in range(default__.abs(450))])) <= (d_909_v57_)
                d_911_v58_: _dafny.Array
                def lambda41_(d_912_v49_):
                    def lambda42_(d_913_i4_):
                        return D1_DC4(d_912_v49_)

                    return lambda42_

                init24_ = lambda41_(d_898_v49_)
                nw165_ = _dafny.Array(None, 9)
                for i0_24_ in range(nw165_.length(0)):
                    nw165_[i0_24_] = init24_(i0_24_)
                d_911_v58_ = nw165_
                d_914_v59_: _dafny.Array
                def lambda43_(d_915_v21_):
                    def lambda44_(d_916_i5_):
                        return d_915_v21_

                    return lambda44_

                init25_ = lambda43_(d_870_v21_)
                nw166_ = _dafny.Array(None, 12)
                for i0_25_ in range(nw166_.length(0)):
                    nw166_[i0_25_] = init25_(i0_25_)
                d_914_v59_ = nw166_
                d_917_v60_: _dafny.Map
                d_917_v60_ = _dafny.Map({d_914_v59_: d_911_v58_})
                index160_ = default__.safeIndex(843, (d_907_v56_).length(0))
                index161_ = default__.safeIndex(428, (d_871_v22_).length(0))
                rhs112_ = (p3) - ((d_866_v18_).f24)
                rhs113_ = d_902_v51_
                rhs114_ = p1
                rhs115_ = ((d_917_v60_)[d_914_v59_] if (d_914_v59_) in (d_917_v60_) else d_911_v58_)
                lhs99_ = self
                lhs100_ = d_907_v56_
                lhs101_ = default__.safeIndex(843, (d_907_v56_).length(0))
                lhs102_ = d_871_v22_
                lhs103_ = default__.safeIndex(428, (d_871_v22_).length(0))
                lhs99_.f21 = rhs112_
                lhs100_[lhs101_] = rhs113_
                lhs102_[lhs103_] = rhs114_
                d_911_v58_ = rhs115_
                (globalState).f6 = len(default__.fm28(len(d_868_v19_), default__.fm2(globalState), (p2) > ((d_866_v18_).f24), globalState))
                r0 = (d_871_v22_)[default__.safeIndex(428, (d_871_v22_).length(0))]
                (globalState).f1 = p1
            elif True:
                d_918_v61_: _dafny.Seq
                d_918_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydlpbiue"))
                d_919_v62_: str
                d_919_v62_ = _dafny.CodePoint('i')
                d_920_v63_: _dafny.Array
                nw167_ = _dafny.Array(int(0), 4)
                d_920_v63_ = nw167_
                d_921_v64_: _dafny.Seq
                d_921_v64_ = _dafny.SeqWithoutIsStrInference([D3_DC6(d_920_v63_)])
                d_922_v65_: C0
                nw168_ = C0()
                nw168_.ctor__(d_921_v64_, p1)
                d_922_v65_ = nw168_
                d_923_v66_: _dafny.Map
                d_923_v66_ = _dafny.Map({p2: d_871_v22_})
                d_924_v67_: _dafny.Seq
                d_924_v67_ = _dafny.SeqWithoutIsStrInference([d_871_v22_])
                d_925_v68_: _dafny.Seq
                d_925_v68_ = _dafny.SeqWithoutIsStrInference([p2])
                d_926_v69_: _dafny.Array
                nw169_ = _dafny.Array(None, 27)
                nw169_[int(0)] = d_871_v22_
                nw169_[int(1)] = d_871_v22_
                nw169_[int(2)] = d_871_v22_
                nw169_[int(3)] = d_871_v22_
                nw169_[int(4)] = d_871_v22_
                nw169_[int(5)] = d_871_v22_
                nw169_[int(6)] = d_871_v22_
                nw169_[int(7)] = ((d_923_v66_)[(d_866_v18_).f24] if ((d_866_v18_).f24) in (d_923_v66_) else d_871_v22_)
                nw169_[int(8)] = d_871_v22_
                nw169_[int(9)] = d_871_v22_
                nw169_[int(10)] = d_871_v22_
                nw169_[int(11)] = (d_924_v67_)[default__.safeIndex(len(d_925_v68_), len(d_924_v67_))]
                nw169_[int(12)] = d_871_v22_
                nw169_[int(13)] = d_871_v22_
                nw169_[int(14)] = d_871_v22_
                nw169_[int(15)] = d_871_v22_
                nw169_[int(16)] = d_871_v22_
                nw169_[int(17)] = d_871_v22_
                nw169_[int(18)] = d_871_v22_
                nw169_[int(19)] = d_871_v22_
                nw169_[int(20)] = d_871_v22_
                nw169_[int(21)] = d_871_v22_
                nw169_[int(22)] = d_871_v22_
                nw169_[int(23)] = d_871_v22_
                nw169_[int(24)] = d_871_v22_
                nw169_[int(25)] = d_871_v22_
                nw169_[int(26)] = d_871_v22_
                d_926_v69_ = nw169_
                d_927_v70_: D7
                d_927_v70_ = D7_DC17(p1, (len((d_918_v61_).set(default__.safeIndex((d_866_v18_).f24, len(d_918_v61_)), d_919_v62_))) in (_dafny.SeqWithoutIsStrInference([len(d_870_v21_) for d_928_i6_ in range(default__.abs(940))])), d_922_v65_, D5_DC10(d_926_v69_), not(p1))
                d_927_v70_ = d_927_v70_
                (globalState).f6 = p2
                d_929_v71_: _dafny.Array
                def lambda45_(d_930_v68_):
                    def lambda46_(d_931_i7_):
                        return (d_930_v68_) + (d_930_v68_)

                    return lambda46_

                init26_ = lambda45_(d_925_v68_)
                nw170_ = _dafny.Array(None, 10)
                for i0_26_ in range(nw170_.length(0)):
                    nw170_[i0_26_] = init26_(i0_26_)
                d_929_v71_ = nw170_
                d_932_v72_: _dafny.Seq
                d_932_v72_ = _dafny.SeqWithoutIsStrInference([d_929_v71_, d_929_v71_, (D14_DC38(d_929_v71_)).cf67])
                d_929_v71_ = (d_932_v72_)[default__.safeIndex(279, len(d_932_v72_))]
                index162_ = default__.safeIndex(35, (d_871_v22_).length(0))
                (d_871_v22_)[index162_] = (d_922_v65_).f23
                index163_ = default__.safeIndex(424, (d_871_v22_).length(0))
                (d_871_v22_)[index163_] = (d_922_v65_).f23
                index164_ = default__.safeIndex(35, (d_871_v22_).length(0))
                index165_ = default__.safeIndex(424, (d_871_v22_).length(0))
                rhs116_ = (d_922_v65_).f23
                rhs117_ = (d_866_v18_).f24
                rhs118_ = (d_922_v65_).f23
                rhs119_ = (d_922_v65_).f23
                lhs104_ = d_871_v22_
                lhs105_ = default__.safeIndex(35, (d_871_v22_).length(0))
                lhs106_ = globalState
                lhs107_ = d_871_v22_
                lhs108_ = default__.safeIndex(424, (d_871_v22_).length(0))
                lhs109_ = globalState
                lhs104_[lhs105_] = rhs116_
                lhs106_.f6 = rhs117_
                lhs107_[lhs108_] = rhs118_
                lhs109_.f1 = rhs119_
                d_933_v73_: C1
                nw171_ = C1()
                nw171_.ctor__(self.f21)
                d_933_v73_ = nw171_
            d_934_v74_: _dafny.Seq
            d_934_v74_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsek"))
            (globalState).f6 = default__.safeModuloInt(p2, len(d_934_v74_))
            d_935_v75_: str
            d_935_v75_ = _dafny.CodePoint('a')
            if (d_935_v75_) not in (d_934_v74_):
                (self).f21 = (self).f20
                (globalState).f6 = default__.fm1((p1 if p1 else p1), globalState)
                d_936_v76_: _dafny.MultiSet
                d_936_v76_ = _dafny.MultiSet([(d_866_v18_).f24, (d_866_v18_).f24, self.f21])
                d_937_v77_: _dafny.Seq
                d_937_v77_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({p1: p1}))])
                d_938_v78_: _dafny.Seq
                d_938_v78_ = _dafny.SeqWithoutIsStrInference([(d_937_v77_)[default__.safeIndex(self.f21, len(d_937_v77_))], (d_866_v18_).f24, (self).f20])
                d_936_v76_ = _dafny.MultiSet((d_937_v77_) + (d_938_v78_))
                (globalState).f1 = ((d_866_v18_).f24) > (165)
                d_939_v79_: _dafny.Array
                nw172_ = _dafny.Array(int(0), 12)
                d_939_v79_ = nw172_
                index166_ = default__.safeIndex(701, (d_939_v79_).length(0))
                (d_939_v79_)[index166_] = -963
                index167_ = default__.safeIndex(701, (d_939_v79_).length(0))
                (d_939_v79_)[index167_] = (0) - ((self).fm15(globalState))
            elif True:
                d_940_v80_: _dafny.Map
                d_940_v80_ = _dafny.Map({not(not (p1) or (p1)): d_934_v74_})
                d_940_v80_ = (d_940_v80_).set(p1, d_934_v74_)
                d_868_v19_ = (d_868_v19_) | (d_868_v19_)
                (globalState).f6 = default__.fm1(not(p1), globalState)
                (globalState).f6 = ((d_866_v18_).f24) * (p3)
                (globalState).f6 = (((self).f20) - (self.f21) if p1 else -212)
            if p1:
                d_941_v81_: _dafny.Map
                d_941_v81_ = _dafny.Map({default__.fm19(globalState): default__.safeDivisionInt(p2, p2)})
                d_942_v82_: _dafny.Seq
                d_942_v82_ = _dafny.SeqWithoutIsStrInference([p3])
                d_943_v83_: _dafny.Map
                d_943_v83_ = _dafny.Map({d_942_v82_: (_dafny.Map({d_870_v21_: (d_866_v18_).f24})) | (d_941_v81_)})
                d_941_v81_ = ((d_943_v83_)[(d_942_v82_) + (d_942_v82_)] if ((d_942_v82_) + (d_942_v82_)) in (d_943_v83_) else d_941_v81_)
                d_944_v84_: _dafny.MultiSet
                d_944_v84_ = _dafny.MultiSet([(self).f20, (d_866_v18_).f24])
                d_945_v85_: _dafny.Set
                d_945_v85_ = _dafny.Set({d_944_v84_})
                d_946_v86_: _dafny.Map
                d_946_v86_ = _dafny.Map({d_945_v85_: default__.safeDivisionInt(p2, (d_866_v18_).f24)})
                d_946_v86_ = (d_946_v86_).set(d_945_v85_, (self).f20)
                (globalState).f6 = -356
                index168_ = default__.safeIndex(178, (d_871_v22_).length(0))
                (d_871_v22_)[index168_] = p1
                index169_ = default__.safeIndex(178, (d_871_v22_).length(0))
                rhs120_ = ((d_934_v74_) <= (d_934_v74_) if p1 else p1)
                rhs121_ = ((d_866_v18_).f24 if p1 else 421)
                lhs110_ = d_871_v22_
                lhs111_ = default__.safeIndex(178, (d_871_v22_).length(0))
                lhs112_ = self
                lhs110_[lhs111_] = rhs120_
                lhs112_.f21 = rhs121_
                d_947_v87_: _dafny.Map
                d_947_v87_ = _dafny.Map({d_935_v75_: d_934_v74_})
                r0 = (d_866_v18_).fm24((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eisnpxsp"))) != (((d_947_v87_)[d_935_v75_] if (d_935_v75_) in (d_947_v87_) else _dafny.SeqWithoutIsStrInference([d_935_v75_ for d_948_i8_ in range(default__.abs(268))]))), d_892_v43_, globalState)
            elif True:
                (globalState).f6 = ((d_866_v18_).f24) - ((d_866_v18_).f24)
                d_949_v88_: _dafny.Set
                d_949_v88_ = _dafny.Set({d_934_v74_})
                rhs122_ = (d_866_v18_).f24
                rhs123_ = default__.fm2(globalState)
                rhs124_ = (d_949_v88_).ispropersubset(d_949_v88_)
                rhs125_ = d_871_v22_
                lhs113_ = globalState
                lhs114_ = globalState
                lhs115_ = globalState
                lhs113_.f6 = rhs122_
                lhs114_.f1 = rhs123_
                lhs115_.f1 = rhs124_
                d_871_v22_ = rhs125_
                d_950_v89_: D5
                d_950_v89_ = D5_DC12(len(d_934_v74_), p1, d_935_v75_)
                pat_let_tv25_ = d_935_v75_
                pat_let_tv26_ = d_866_v18_
                d_951_v90_: _dafny.MultiSet
                def iife131_(_pat_let46_0):
                    def iife132_(d_952_dt__update__tmp_h0_):
                        def iife133_(_pat_let47_0):
                            def iife134_(d_953_dt__update_hcf17_h0_):
                                def iife135_(_pat_let48_0):
                                    def iife136_(d_954_dt__update_hcf15_h0_):
                                        return D5_DC12(d_954_dt__update_hcf15_h0_, (d_952_dt__update__tmp_h0_).cf16, d_953_dt__update_hcf17_h0_)
                                    return iife136_(_pat_let48_0)
                                return iife135_((pat_let_tv26_).f24)
                            return iife134_(_pat_let47_0)
                        return iife133_(pat_let_tv25_)
                    return iife132_(_pat_let46_0)
                d_951_v90_ = _dafny.MultiSet([D5_DC12((d_866_v18_).f24, p1, d_935_v75_), iife131_(d_950_v89_), d_950_v89_])
                d_955_v91_: _dafny.Seq
                d_955_v91_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_956_v92_: C0
                nw173_ = C0()
                nw173_.ctor__(d_955_v91_, True)
                d_956_v92_ = nw173_
                d_957_v93_: _dafny.Seq
                d_957_v93_ = _dafny.SeqWithoutIsStrInference([d_956_v92_, d_956_v92_])
                d_958_v94_: _dafny.Set
                d_958_v94_ = _dafny.Set({(d_934_v74_)[default__.safeIndex(len(d_957_v93_), len(d_934_v74_))], d_935_v75_})
                d_959_v95_: _dafny.Array
                nw174_ = _dafny.Array(None, 25)
                nw174_[int(0)] = _dafny.Map({676: p1})
                nw174_[int(1)] = d_891_v42_
                nw174_[int(2)] = d_891_v42_
                nw174_[int(3)] = d_891_v42_
                nw174_[int(4)] = _dafny.Map({-939: p1})
                nw174_[int(5)] = d_891_v42_
                nw174_[int(6)] = d_891_v42_
                nw174_[int(7)] = d_891_v42_
                nw174_[int(8)] = d_891_v42_
                nw174_[int(9)] = d_891_v42_
                nw174_[int(10)] = default__.fm0(p1, p1, self.f21, globalState)
                nw174_[int(11)] = d_891_v42_
                nw174_[int(12)] = d_891_v42_
                nw174_[int(13)] = d_891_v42_
                nw174_[int(14)] = _dafny.Map({(0) - (-542): True})
                nw174_[int(15)] = d_891_v42_
                nw174_[int(16)] = d_891_v42_
                nw174_[int(17)] = d_891_v42_
                nw174_[int(18)] = d_891_v42_
                nw174_[int(19)] = d_891_v42_
                nw174_[int(20)] = d_891_v42_
                nw174_[int(21)] = (d_891_v42_).set(len(d_958_v94_), p1)
                nw174_[int(22)] = d_891_v42_
                nw174_[int(23)] = d_891_v42_
                nw174_[int(24)] = d_891_v42_
                d_959_v95_ = nw174_
                d_960_v96_: D9
                d_960_v96_ = D9_DC22(d_959_v95_)
                d_961_v97_: _dafny.Array
                nw175_ = _dafny.Array(None, 9)
                nw175_[int(0)] = len(d_934_v74_)
                nw175_[int(1)] = len(d_934_v74_)
                nw175_[int(2)] = -550
                nw175_[int(3)] = (d_866_v18_).f24
                nw175_[int(4)] = p2
                nw175_[int(5)] = (d_866_v18_).f24
                nw175_[int(6)] = (self).f20
                nw175_[int(7)] = p3
                nw175_[int(8)] = (d_866_v18_).f24
                d_961_v97_ = nw175_
                d_962_v98_: _dafny.Array
                nw176_ = _dafny.Array(None, 10)
                nw176_[int(0)] = p0
                nw176_[int(1)] = p0
                nw176_[int(2)] = p0
                nw176_[int(3)] = p0
                nw176_[int(4)] = p0
                nw176_[int(5)] = p0
                nw176_[int(6)] = D3_DC6(d_961_v97_)
                nw176_[int(7)] = D3_DC6(d_961_v97_)
                nw176_[int(8)] = p0
                nw176_[int(9)] = p0
                d_962_v98_ = nw176_
                d_963_v99_: _dafny.Set
                d_963_v99_ = _dafny.Set({d_962_v98_})
                d_964_v100_: D9
                d_964_v100_ = D9_DC25(D9_DC23(d_963_v99_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkxf")), (d_956_v92_).f23, d_961_v97_, d_935_v75_))
                d_965_v101_: _dafny.Map
                d_965_v101_ = _dafny.Map({p1: d_870_v21_})
                d_966_v103_: _dafny.Array
                nw177_ = _dafny.Array(None, 15)
                nw177_[int(0)] = (self.f21) - ((d_951_v90_).cardinality)
                nw177_[int(1)] = (p3 if p1 else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hty"))))
                nw177_[int(2)] = (0) - (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([D9_DC25(d_960_v96_), d_964_v100_])), p3))
                nw177_[int(3)] = (d_866_v18_).f24
                nw177_[int(4)] = len(((d_965_v101_)[(d_956_v92_).f23] if ((d_956_v92_).f23) in (d_965_v101_) else default__.fm19(globalState)))
                nw177_[int(5)] = (len(d_870_v21_)) - (self.f21)
                nw177_[int(6)] = default__.safeModuloInt(len(d_870_v21_), (d_866_v18_).f24)
                nw177_[int(7)] = (d_866_v18_).f24
                nw177_[int(8)] = (0) - (len(d_870_v21_))
                nw177_[int(9)] = default__.safeModuloInt((self).f20, (self).f20)
                nw177_[int(10)] = p3
                def iife137_():
                    coll39_ = _dafny.Map()
                    compr_39_: int
                    for compr_39_ in (_dafny.Map({self.f21: len(d_870_v21_)})).keys.Elements:
                        d_967_v102_: int = compr_39_
                        if (d_967_v102_) in (_dafny.Map({self.f21: len(d_870_v21_)})):
                            coll39_[(d_967_v102_) * ((d_866_v18_).f24)] = d_934_v74_
                    return _dafny.Map(coll39_)
                nw177_[int(11)] = len(iife137_()
                )
                nw177_[int(12)] = (self).f20
                nw177_[int(13)] = (d_866_v18_).f24
                nw177_[int(14)] = p2
                d_966_v103_ = nw177_
                index170_ = default__.safeIndex(396, (d_966_v103_).length(0))
                (d_966_v103_)[index170_] = self.f21
                d_968_v104_: _dafny.Seq
                d_968_v104_ = _dafny.SeqWithoutIsStrInference([default__.fm1(False, globalState), p3])
                index171_ = default__.safeIndex(396, (d_966_v103_).length(0))
                (d_966_v103_)[index171_] = (0) - (default__.safeDivisionInt((len((D3_DC8(d_968_v104_, d_934_v74_, p1, d_935_v75_)).cf9)) + (391), default__.fm1((d_956_v92_).fm17(len(_dafny.Map({p1: d_935_v75_})), 24, (d_956_v92_).f23, globalState), globalState)))
                d_969_v105_: C0
                nw178_ = C0()
                nw178_.ctor__((d_955_v91_) + ((d_956_v92_).f22), True)
                d_969_v105_ = nw178_
                (d_866_v18_).m2((d_969_v105_).f23, globalState)
        hi7_ = (0) - ((d_866_v18_).f24)
        for d_970_i9_ in range(default__.fm1(p1, globalState), hi7_):
            (globalState).f6 = (d_866_v18_).f24
            (globalState).f6 = (d_866_v18_).f24
            d_971_v106_: _dafny.Set
            d_971_v106_ = _dafny.Set({self.f21})
            d_972_v107_: _dafny.Map
            d_972_v107_ = _dafny.Map({p1: (d_971_v106_) | (d_971_v106_)})
            d_972_v107_ = (d_972_v107_).set(p1, default__.fm28((d_866_v18_).f24, p1, p1, globalState))
            if p1:
                index172_ = default__.safeIndex(598, (d_871_v22_).length(0))
                (d_871_v22_)[index172_] = p1
                d_973_v108_: _dafny.Seq
                d_973_v108_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpy"))
                index173_ = default__.safeIndex(598, (d_871_v22_).length(0))
                (d_871_v22_)[index173_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ja"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_974_i10_ in range(default__.abs(622))]))) != (d_973_v108_)
                (globalState).f1 = (-549) <= (len(d_870_v21_))
                d_975_v109_: _dafny.Seq
                d_975_v109_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_976_v110_: C0
                nw179_ = C0()
                nw179_.ctor__(d_975_v109_, (d_871_v22_)[default__.safeIndex(598, (d_871_v22_).length(0))])
                d_976_v110_ = nw179_
                d_977_v111_: _dafny.MultiSet
                d_977_v111_ = _dafny.MultiSet([d_976_v110_, d_976_v110_, d_976_v110_])
                d_978_v112_: _dafny.Seq
                d_978_v112_ = _dafny.SeqWithoutIsStrInference([d_976_v110_])
                d_979_v113_: _dafny.Seq
                d_979_v113_ = _dafny.SeqWithoutIsStrInference([d_978_v112_, d_978_v112_, d_978_v112_])
                d_980_v114_: _dafny.Array
                nw180_ = _dafny.Array(None, 22)
                nw180_[int(0)] = (d_977_v111_).set(d_976_v110_, default__.abs((d_866_v18_).f24))
                nw180_[int(1)] = d_977_v111_
                nw180_[int(2)] = (d_977_v111_).set(d_976_v110_, default__.abs(p3))
                nw180_[int(3)] = d_977_v111_
                nw180_[int(4)] = d_977_v111_
                nw180_[int(5)] = d_977_v111_
                nw180_[int(6)] = d_977_v111_
                nw180_[int(7)] = d_977_v111_
                nw180_[int(8)] = d_977_v111_
                nw180_[int(9)] = _dafny.MultiSet((d_978_v112_) + (d_978_v112_))
                nw180_[int(10)] = d_977_v111_
                nw180_[int(11)] = _dafny.MultiSet(((d_979_v113_)[default__.safeIndex(p2, len(d_979_v113_))]) + (d_978_v112_))
                nw180_[int(12)] = (d_977_v111_) | (d_977_v111_)
                nw180_[int(13)] = d_977_v111_
                nw180_[int(14)] = (_dafny.MultiSet([d_976_v110_, d_976_v110_])).set(d_976_v110_, default__.abs((d_866_v18_).f24))
                nw180_[int(15)] = (d_977_v111_).intersection(d_977_v111_)
                nw180_[int(16)] = d_977_v111_
                nw180_[int(17)] = d_977_v111_
                nw180_[int(18)] = d_977_v111_
                nw180_[int(19)] = d_977_v111_
                nw180_[int(20)] = d_977_v111_
                nw180_[int(21)] = (_dafny.MultiSet([d_976_v110_, d_976_v110_])).set(d_976_v110_, default__.abs((d_866_v18_).f24))
                d_980_v114_ = nw180_
                index174_ = default__.safeIndex(726, (d_980_v114_).length(0))
                (d_980_v114_)[index174_] = d_977_v111_
                d_981_v115_: str
                d_981_v115_ = _dafny.CodePoint('g')
                d_982_v116_: _dafny.Map
                d_982_v116_ = _dafny.Map({d_981_v115_: (d_871_v22_)[default__.safeIndex(598, (d_871_v22_).length(0))]})
                d_983_v117_: _dafny.Map
                d_983_v117_ = _dafny.Map({p1: ((d_982_v116_)[d_981_v115_] if (d_981_v115_) in (d_982_v116_) else not((d_871_v22_)[default__.safeIndex(598, (d_871_v22_).length(0))]))})
                d_984_v118_: _dafny.Map
                d_984_v118_ = _dafny.Map({len(d_983_v117_): _dafny.MultiSet(d_978_v112_)})
                d_985_v119_: _dafny.Array
                nw181_ = _dafny.Array(False, 7)
                d_985_v119_ = nw181_
                index175_ = default__.safeIndex(726, (d_980_v114_).length(0))
                rhs126_ = (((d_984_v118_)[(d_866_v18_).f24] if ((d_866_v18_).f24) in (d_984_v118_) else d_977_v111_)) - (d_977_v111_)
                rhs127_ = (d_871_v22_)[default__.safeIndex(598, (d_871_v22_).length(0))]
                rhs128_ = not((d_871_v22_)[default__.safeIndex(598, (d_871_v22_).length(0))])
                rhs129_ = d_985_v119_
                lhs116_ = d_980_v114_
                lhs117_ = default__.safeIndex(726, (d_980_v114_).length(0))
                lhs118_ = globalState
                lhs119_ = globalState
                lhs116_[lhs117_] = rhs126_
                lhs118_.f1 = rhs127_
                r0 = rhs128_
                lhs119_.f14 = rhs129_
                d_971_v106_ = default__.fm28(self.f21, (d_871_v22_)[default__.safeIndex(598, (d_871_v22_).length(0))], (d_868_v19_).ispropersubset(d_868_v19_), globalState)
                d_986_v120_: _dafny.Seq
                d_986_v120_ = _dafny.SeqWithoutIsStrInference([d_866_v18_, d_866_v18_])
                d_866_v18_ = (d_986_v120_)[default__.safeIndex(-25, len(d_986_v120_))]
            elif True:
                d_987_v121_: _dafny.Array
                def lambda47_(d_988_i11_):
                    return (d_988_i11_) - (905)

                init27_ = lambda47_
                nw182_ = _dafny.Array(None, 16)
                for i0_27_ in range(nw182_.length(0)):
                    nw182_[i0_27_] = init27_(i0_27_)
                d_987_v121_ = nw182_
                index176_ = default__.safeIndex(821, (d_987_v121_).length(0))
                (d_987_v121_)[index176_] = len((d_868_v19_ if p1 else d_868_v19_))
                d_989_v122_: _dafny.Seq
                d_989_v122_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kn"))
                d_990_v123_: str
                d_990_v123_ = _dafny.CodePoint('a')
                d_991_v124_: _dafny.Map
                d_991_v124_ = _dafny.Map({d_990_v123_: _dafny.Map({p1: p1})})
                d_992_v125_: _dafny.Map
                d_992_v125_ = _dafny.Map({p1: p1})
                d_993_v126_: _dafny.Map
                d_993_v126_ = _dafny.Map({(d_866_v18_).f24: ((d_991_v124_)[d_990_v123_] if (d_990_v123_) in (d_991_v124_) else d_992_v125_)})
                index177_ = default__.safeIndex(821, (d_987_v121_).length(0))
                rhs130_ = (d_866_v18_).f24
                rhs131_ = (self.f21) == ((0) - (len(d_993_v126_)))
                rhs132_ = (d_989_v122_).set(default__.safeIndex(d_970_i9_, len(d_989_v122_)), _dafny.CodePoint('r'))
                rhs133_ = p1
                lhs120_ = d_987_v121_
                lhs121_ = default__.safeIndex(821, (d_987_v121_).length(0))
                lhs122_ = globalState
                lhs123_ = globalState
                lhs120_[lhs121_] = rhs130_
                lhs122_.f1 = rhs131_
                d_989_v122_ = rhs132_
                lhs123_.f1 = rhs133_
                rhs134_ = (self).fm15(globalState)
                rhs135_ = p1
                lhs124_ = self
                lhs124_.f21 = rhs134_
                r0 = rhs135_
                d_994_v127_: _dafny.Array
                nw183_ = _dafny.Array(_dafny.Map({}), 12)
                d_994_v127_ = nw183_
                d_995_v128_: _dafny.MultiSet
                d_995_v128_ = _dafny.MultiSet([d_971_v106_])
                d_996_v129_: _dafny.Seq
                d_996_v129_ = _dafny.SeqWithoutIsStrInference([d_989_v122_])
                d_997_v130_: _dafny.Map
                d_997_v130_ = _dafny.Map({d_995_v128_: (d_996_v129_)[default__.safeIndex((d_866_v18_).f24, len(d_996_v129_))]})
                rhs136_ = (d_994_v127_ if p1 else d_994_v127_)
                rhs137_ = not(not(p1))
                rhs138_ = len(((d_997_v130_)[d_995_v128_] if (d_995_v128_) in (d_997_v130_) else d_989_v122_))
                rhs139_ = _dafny.CodePoint('v')
                lhs125_ = globalState
                lhs126_ = globalState
                r2 = rhs136_
                lhs125_.f1 = rhs137_
                lhs126_.f6 = rhs138_
                d_990_v123_ = rhs139_
                index178_ = default__.safeIndex(96, (d_871_v22_).length(0))
                (d_871_v22_)[index178_] = p1
                index179_ = default__.safeIndex(96, (d_871_v22_).length(0))
                (d_871_v22_)[index179_] = p1
                (globalState).f1 = (d_871_v22_)[default__.safeIndex(96, (d_871_v22_).length(0))]
        r0 = not(not(p1))
        d_998_v131_: _dafny.Seq
        d_998_v131_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwtwcswh"))
        d_999_v132_: _dafny.Map
        d_999_v132_ = _dafny.Map({((self).f20) * (p2): ((0) - (self.f21)) - (len(d_998_v131_))})
        r1 = d_999_v132_
        d_1000_v133_: _dafny.Array
        nw184_ = _dafny.Array(_dafny.Map({}), 25)
        d_1000_v133_ = nw184_
        r2 = d_1000_v133_
        return r0, r1, r2

    @property
    def f20(self):
        return self._f20

class C3(T0, T1):
    def  __init__(self):
        self._f17: _dafny.Map = _dafny.Map({})
        self._f18: bool = False
        self._f19: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f17(self):
        return self._f17
    @f17.setter
    def f17(self, value):
        self._f17 = value
    def ctor__(self, f18, f19, f17):
        (self)._f18 = f18
        (self)._f19 = f19
        (self).f17 = f17

    def fm3(self, p0, p1, p2, p3, globalState):
        return D0_DC2(D0_DC0((self).f18))

    def m1(self, globalState):
        r0: int = int(0)
        r1: bool = False
        (globalState).f1 = not(((0) - ((self).f19)) <= ((self).f19))
        d_1001_v0_: _dafny.Array
        nw185_ = _dafny.Array(_dafny.Seq({}), 8)
        d_1001_v0_ = nw185_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_1001_v0_).length(0)):
            d_1002_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_1002_i0_)) and ((d_1002_i0_) < ((d_1001_v0_).length(0)))):
                (d_1001_v0_)[(d_1002_i0_)] = (_dafny.SeqWithoutIsStrInference([(self).f19 for d_1003_i1_ in range(default__.abs(659))])) + (_dafny.SeqWithoutIsStrInference([(self).f19]))
        d_1004_i2_: int
        d_1004_i2_ = 0
        with _dafny.label("6"):
            while (self).f18:
                with _dafny.c_label("6"):
                    if (d_1004_i2_) >= (100):
                        raise _dafny.Break("6")
                    d_1004_i2_ = (d_1004_i2_) + (1)
                    d_1005_v1_: _dafny.Set
                    d_1005_v1_ = _dafny.Set({(self).f18, False})
                    if not (not((self).f18)) or ((_dafny.Set({(self).f18})).ispropersubset(d_1005_v1_)):
                        (globalState).f1 = (not ((self).f18) or (True) if (self).f18 else (self).f18)
                        d_1006_v2_: _dafny.Seq
                        d_1006_v2_ = _dafny.SeqWithoutIsStrInference([(self).f18])
                        d_1007_v3_: _dafny.Array
                        nw186_ = _dafny.Array(int(0), 29)
                        d_1007_v3_ = nw186_
                        d_1008_v4_: D3
                        d_1008_v4_ = D3_DC6(d_1007_v3_)
                        d_1009_v5_: _dafny.Seq
                        d_1009_v5_ = _dafny.SeqWithoutIsStrInference([d_1008_v4_])
                        d_1010_v6_: _dafny.MultiSet
                        d_1010_v6_ = _dafny.MultiSet([(self).f18, False, (self).f18])
                        d_1011_v7_: C0
                        nw187_ = C0()
                        nw187_.ctor__(((d_1009_v5_).set(default__.safeIndex((d_1010_v6_).cardinality, len(d_1009_v5_)), d_1008_v4_) if (d_1006_v2_)[default__.safeIndex((self).f19, len(d_1006_v2_))] else d_1009_v5_), (self).f18)
                        d_1011_v7_ = nw187_
                        (globalState).f6 = (self).f19
                        d_1012_v8_: _dafny.Array
                        nw188_ = _dafny.Array(_dafny.Array(None, 0), 24)
                        d_1012_v8_ = nw188_
                        d_1013_v9_: D5
                        d_1013_v9_ = D5_DC10(d_1012_v8_)
                        d_1014_v10_: D7
                        d_1014_v10_ = D7_DC17((self).f18, (self).f18, d_1011_v7_, (d_1013_v9_ if (self).f18 else D5_DC10(d_1012_v8_)), (d_1011_v7_).f23)
                        d_1015_v11_: _dafny.Array
                        nw189_ = _dafny.Array(False, 1)
                        d_1015_v11_ = nw189_
                        index180_ = default__.safeIndex(885, (d_1015_v11_).length(0))
                        (d_1015_v11_)[index180_] = (d_1011_v7_).f23
                        d_1016_v12_: _dafny.Seq
                        d_1016_v12_ = _dafny.SeqWithoutIsStrInference([(self).f19])
                        d_1017_v13_: str
                        d_1017_v13_ = _dafny.CodePoint('m')
                        d_1018_v14_: D3
                        d_1018_v14_ = D3_DC8(d_1016_v12_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brgrt")), (d_1011_v7_).f23, d_1017_v13_)
                        d_1019_v15_: _dafny.Seq
                        d_1019_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnu"))
                        d_1020_v16_: _dafny.MultiSet
                        d_1020_v16_ = _dafny.MultiSet([len(d_1019_v15_), (self).f19])
                        index181_ = default__.safeIndex(885, (d_1015_v11_).length(0))
                        rhs140_ = D7_DC17((d_1011_v7_).fm17(len(d_1016_v12_), len((d_1018_v14_).cf8), (d_1011_v7_).f23, globalState), (d_1011_v7_).f23, d_1011_v7_, (d_1013_v9_ if (self).f18 else d_1013_v9_), (d_1011_v7_).f23)
                        rhs141_ = default__.fm2(globalState)
                        rhs142_ = ((d_1020_v16_) - (d_1020_v16_)).isdisjoint(d_1020_v16_)
                        rhs143_ = (self).f18
                        lhs127_ = d_1015_v11_
                        lhs128_ = default__.safeIndex(885, (d_1015_v11_).length(0))
                        lhs129_ = globalState
                        lhs130_ = globalState
                        d_1014_v10_ = rhs140_
                        lhs127_[lhs128_] = rhs141_
                        lhs129_.f1 = rhs142_
                        lhs130_.f1 = rhs143_
                        (globalState).f1 = default__.fm2(globalState)
                    elif True:
                        r0 = (0) - ((self).f19)
                        (globalState).f1 = not((self).f18)
                        d_1021_v17_: _dafny.Map
                        d_1021_v17_ = _dafny.Map({(self).f18: _dafny.MultiSet([(self).f19])})
                        d_1022_v18_: _dafny.MultiSet
                        d_1022_v18_ = _dafny.MultiSet([len(d_1005_v1_)])
                        d_1023_v19_: _dafny.Set
                        d_1023_v19_ = _dafny.Set({((d_1021_v17_)[(self).f18] if ((self).f18) in (d_1021_v17_) else d_1022_v18_)})
                        d_1024_v20_: _dafny.Seq
                        d_1024_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "itmiqpxoa"))
                        d_1025_v21_: D3
                        d_1025_v21_ = D3_DC7((self).f19)
                        d_1026_v22_: _dafny.Array
                        nw190_ = _dafny.Array(False, 24)
                        d_1026_v22_ = nw190_
                        d_1027_v23_: _dafny.Map
                        d_1027_v23_ = _dafny.Map({d_1026_v22_: (self).f18})
                        d_1028_v24_: D10
                        d_1028_v24_ = D10_DC27(len(d_1023_v19_), d_1024_v20_, (self).f19, d_1025_v21_, d_1027_v23_)
                        (globalState).f6 = (d_1028_v24_).cf46
                        d_1029_v25_: str
                        d_1029_v25_ = _dafny.CodePoint('o')
                        d_1029_v25_ = d_1029_v25_
                        d_1030_v26_: C2
                        nw191_ = C2()
                        nw191_.ctor__(((self).f19) + ((self).f19), (self).f19, self.f17)
                        d_1030_v26_ = nw191_
                    d_1031_v27_: _dafny.Seq
                    d_1031_v27_ = _dafny.SeqWithoutIsStrInference([(self).f19, (self).f19])
                    d_1032_v28_: _dafny.Seq
                    d_1032_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cpvoudn"))
                    d_1033_v29_: str
                    d_1033_v29_ = _dafny.CodePoint('u')
                    d_1034_v30_: D3
                    d_1034_v30_ = D3_DC8((d_1031_v27_).set(default__.safeIndex((self).f19, len(d_1031_v27_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlep")))), d_1032_v28_, (self).f18, d_1033_v29_)
                    d_1035_v31_: _dafny.Seq
                    d_1035_v31_ = _dafny.SeqWithoutIsStrInference([(d_1034_v30_).cf10, (self).f18, False, (self).f18])
                    d_1036_v32_: _dafny.Map
                    d_1036_v32_ = _dafny.Map({(self).f18: not((self).f18)})
                    d_1037_v33_: _dafny.Array
                    def lambda48_(d_1038_i3_):
                        return (self).f18

                    init28_ = lambda48_
                    nw192_ = _dafny.Array(None, 5)
                    for i0_28_ in range(nw192_.length(0)):
                        nw192_[i0_28_] = init28_(i0_28_)
                    d_1037_v33_ = nw192_
                    d_1039_v34_: _dafny.Seq
                    d_1039_v34_ = _dafny.SeqWithoutIsStrInference([d_1037_v33_])
                    d_1035_v31_ = (((d_1035_v31_ if ((d_1036_v32_)[(self).f18] if ((self).f18) in (d_1036_v32_) else (self).f18) else d_1035_v31_)) + (d_1035_v31_)).set(default__.safeIndex(len(d_1039_v34_), len(((d_1035_v31_ if ((d_1036_v32_)[(self).f18] if ((self).f18) in (d_1036_v32_) else (self).f18) else d_1035_v31_)) + (d_1035_v31_))), (self).f18)
                    d_1036_v32_ = (d_1036_v32_ if (_dafny.MultiSet([(self).f19])).ispropersubset(_dafny.MultiSet([(self).f19, (self).f19, (0) - ((self).f19)])) else (d_1036_v32_) | (_dafny.Map({(self).f18: (self).f18})))
                    (globalState).f6 = (self).f19
                    pass
            pass
        d_1040_i4_: int
        d_1040_i4_ = 0
        with _dafny.label("7"):
            while not ((self).f18) or ((self).f18):
                with _dafny.c_label("7"):
                    if (d_1040_i4_) >= (100):
                        raise _dafny.Break("7")
                    d_1040_i4_ = (d_1040_i4_) + (1)
                    d_1041_v35_: _dafny.Seq
                    d_1041_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
                    d_1041_v35_ = default__.fm21(globalState)
                    d_1042_v36_: _dafny.Array
                    nw193_ = _dafny.Array(None, 1)
                    nw193_[int(0)] = (self).f18
                    d_1042_v36_ = nw193_
                    d_1043_v37_: _dafny.Array
                    nw194_ = _dafny.Array(None, 2)
                    nw194_[int(0)] = d_1042_v36_
                    nw194_[int(1)] = d_1042_v36_
                    d_1043_v37_ = nw194_
                    d_1044_v38_: D5
                    d_1044_v38_ = D5_DC10(d_1043_v37_)
                    pat_let_tv27_ = d_1043_v37_
                    pat_let_tv28_ = d_1043_v37_
                    def iife139_(_pat_let50_0):
                        def iife140_(d_1045_dt__update__tmp_h0_):
                            def iife141_(_pat_let51_0):
                                def iife142_(d_1046_dt__update_hcf13_h0_):
                                    return D5_DC10(d_1046_dt__update_hcf13_h0_)
                                return iife142_(_pat_let51_0)
                            return iife141_(pat_let_tv27_)
                        return iife140_(_pat_let50_0)
                    def iife138_(_pat_let49_0):
                        def iife143_(d_1047_dt__update__tmp_h1_):
                            def iife144_(_pat_let52_0):
                                def iife145_(d_1048_dt__update_hcf13_h1_):
                                    return D5_DC10(d_1048_dt__update_hcf13_h1_)
                                return iife145_(_pat_let52_0)
                            return iife144_(pat_let_tv28_)
                        return iife143_(_pat_let49_0)
                    source8_ = iife138_((iife139_(d_1044_v38_) if False else d_1044_v38_))
                    if source8_.is_DC11:
                        d_1049___mcc_h0_ = source8_.cf14
                        d_1050_cf14_ = d_1049___mcc_h0_
                        d_1051_v39_: str
                        d_1051_v39_ = _dafny.CodePoint('x')
                        d_1051_v39_ = (d_1041_v35_)[default__.safeIndex((self).f19, len(d_1041_v35_))]
                        r1 = (self).f18
                        (globalState).f1 = not(True)
                        d_1052_v40_: _dafny.Array
                        nw195_ = _dafny.Array(_dafny.Map({}), 12)
                        d_1052_v40_ = nw195_
                        d_1053_v42_: D15
                        def iife146_():
                            coll40_ = _dafny.Set()
                            compr_40_: int
                            for compr_40_ in _dafny.IntegerRange(688, 977):
                                d_1054_v41_: int = compr_40_
                                if ((688) <= (d_1054_v41_)) and ((d_1054_v41_) < (977)):
                                    coll40_ = coll40_.union(_dafny.Set([default__.safeDivisionInt(d_1054_v41_, d_1050_cf14_)]))
                            return _dafny.Set(coll40_)
                        d_1053_v42_ = D15_DC42(iife146_()
)
                        d_1055_v43_: _dafny.Map
                        d_1055_v43_ = _dafny.Map({(d_1053_v42_).cf73: d_1042_v36_})
                        index182_ = default__.safeIndex(508, (d_1052_v40_).length(0))
                        (d_1052_v40_)[index182_] = d_1055_v43_
                        index183_ = default__.safeIndex(508, (d_1052_v40_).length(0))
                        (d_1052_v40_)[index183_] = d_1055_v43_
                    elif source8_.is_DC12:
                        d_1056___mcc_h1_ = source8_.cf15
                        d_1057___mcc_h2_ = source8_.cf16
                        d_1058___mcc_h3_ = source8_.cf17
                        d_1059_cf17_ = d_1058___mcc_h3_
                        d_1060_cf16_ = d_1057___mcc_h2_
                        d_1061_cf15_ = d_1056___mcc_h1_
                        d_1041_v35_ = (default__.fm21(globalState)) + ((d_1041_v35_) + (d_1041_v35_))
                        d_1062_v44_: _dafny.Seq
                        d_1062_v44_ = _dafny.SeqWithoutIsStrInference([(self).f18, d_1060_cf16_])
                        d_1061_cf15_ = ((len(d_1062_v44_)) + (853)) * (d_1061_cf15_)
                        r0 = d_1061_cf15_
                        d_1063_v45_: _dafny.Seq
                        d_1063_v45_ = _dafny.SeqWithoutIsStrInference([125])
                        d_1063_v45_ = d_1063_v45_
                    elif source8_.is_DC10:
                        d_1064___mcc_h4_ = source8_.cf13
                        d_1065_cf13_ = d_1064___mcc_h4_
                        index184_ = default__.safeIndex(986, (d_1042_v36_).length(0))
                        (d_1042_v36_)[index184_] = (self).f18
                        index185_ = default__.safeIndex(986, (d_1042_v36_).length(0))
                        (d_1042_v36_)[index185_] = ((self).f19) < ((self).f19)
                        (globalState).f6 = ((self).f19) + ((self).f19)
                        d_1066_v46_: _dafny.MultiSet
                        d_1066_v46_ = _dafny.MultiSet([(self).f19])
                        d_1067_v47_: C2
                        nw196_ = C2()
                        nw196_.ctor__(default__.safeModuloInt((0) - ((self).f19), (d_1066_v46_).cardinality), (self).f19, (self.f17 if (self).f18 else self.f17))
                        d_1067_v47_ = nw196_
                        d_1068_v49_: _dafny.Set
                        d_1068_v49_ = _dafny.Set({False})
                        d_1069_v50_: _dafny.Map
                        d_1069_v50_ = _dafny.Map({len(d_1068_v49_): (d_1067_v47_).f20})
                        d_1070_v52_: _dafny.Array
                        def lambda49_(d_1071_i5_):
                            return (d_1071_i5_) * (314)

                        init29_ = lambda49_
                        nw197_ = _dafny.Array(None, 21)
                        for i0_29_ in range(nw197_.length(0)):
                            nw197_[i0_29_] = init29_(i0_29_)
                        d_1070_v52_ = nw197_
                        d_1072_v53_: _dafny.Set
                        d_1072_v53_ = _dafny.Set({d_1070_v52_})
                        d_1073_v54_: _dafny.Seq
                        d_1073_v54_ = _dafny.SeqWithoutIsStrInference([d_1072_v53_])
                        d_1074_v55_: _dafny.Seq
                        d_1074_v55_ = _dafny.SeqWithoutIsStrInference([-736, default__.fm1((self).f18, globalState), len((d_1073_v54_)[default__.safeIndex(d_1067_v47_.f21, len(d_1073_v54_))])])
                        d_1075_v56_: _dafny.Map
                        d_1075_v56_ = _dafny.Map({(d_1067_v47_).f20: (d_1042_v36_)[default__.safeIndex(986, (d_1042_v36_).length(0))]})
                        d_1076_v57_: _dafny.Set
                        d_1076_v57_ = _dafny.Set({(d_1067_v47_).f20, (self).f19})
                        d_1077_v58_: _dafny.Map
                        d_1077_v58_ = _dafny.Map({(d_1042_v36_)[default__.safeIndex(986, (d_1042_v36_).length(0))]: len(d_1076_v57_)})
                        d_1078_v59_: _dafny.Array
                        nw198_ = _dafny.Array(None, 10)
                        nw198_[int(0)] = d_1067_v47_.f21
                        nw198_[int(1)] = (-249) + ((self).f19)
                        def iife147_():
                            def iife148_():
                                coll42_ = _dafny.Map()
                                compr_42_: int
                                for compr_42_ in _dafny.IntegerRange(378, 24):
                                    d_1080_v51_: int = compr_42_
                                    if ((378) <= (d_1080_v51_)) and ((d_1080_v51_) < (24)):
                                        coll42_[(d_1080_v51_) + (91)] = not((self).f18)
                                return _dafny.Map(coll42_)
                            coll41_ = _dafny.Map()
                            compr_41_: int
                            for compr_41_ in (d_1069_v50_).keys.Elements:
                                d_1079_v48_: int = compr_41_
                                if (d_1079_v48_) in (d_1069_v50_):
                                    coll41_[default__.safeModuloInt(d_1079_v48_, len(iife148_()
                                    ))] = (self).f18
                            return _dafny.Map(coll41_)
                        nw198_[int(2)] = (len(iife147_()
                        )) * ((d_1067_v47_).f20)
                        nw198_[int(3)] = (self).f19
                        nw198_[int(4)] = len(d_1074_v55_)
                        nw198_[int(5)] = d_1067_v47_.f21
                        nw198_[int(6)] = default__.safeDivisionInt(len(d_1075_v56_), (0) - (d_1067_v47_.f21))
                        nw198_[int(7)] = ((d_1077_v58_)[(d_1042_v36_)[default__.safeIndex(986, (d_1042_v36_).length(0))]] if ((d_1042_v36_)[default__.safeIndex(986, (d_1042_v36_).length(0))]) in (d_1077_v58_) else -546)
                        nw198_[int(8)] = ((d_1067_v47_).f20) * ((d_1067_v47_).f20)
                        nw198_[int(9)] = d_1067_v47_.f21
                        d_1078_v59_ = nw198_
                        index186_ = default__.safeIndex(351, (d_1070_v52_).length(0))
                        (d_1070_v52_)[index186_] = len(default__.fm28((0) - ((d_1067_v47_).f20), (self).f18, (d_1042_v36_)[default__.safeIndex(986, (d_1042_v36_).length(0))], globalState))
                        index187_ = default__.safeIndex(351, (d_1070_v52_).length(0))
                        (d_1070_v52_)[index187_] = (0) - ((self).f19)
                    elif True:
                        d_1081___mcc_h5_ = source8_.cf18
                        d_1082_cf18_ = d_1081___mcc_h5_
                        d_1083_v60_: D13
                        d_1083_v60_ = D13_DC37((self).f18, (False if (self).f18 else (self).f18))
                        d_1083_v60_ = D13_DC37((self).f18, (self).f18)
                        (globalState).f6 = (self).f19
                        d_1084_v61_: _dafny.Map
                        d_1084_v61_ = _dafny.Map({default__.fm2(globalState): (self).f19})
                        d_1085_v62_: C2
                        nw199_ = C2()
                        nw199_.ctor__(len(d_1084_v61_), (self).f19, (self.f17) | (self.f17))
                        d_1085_v62_ = nw199_
                        d_1086_v63_: _dafny.Array
                        def lambda50_(d_1087_i6_):
                            return _dafny.SeqWithoutIsStrInference([(self).f18])

                        init30_ = lambda50_
                        nw200_ = _dafny.Array(None, 7)
                        for i0_30_ in range(nw200_.length(0)):
                            nw200_[i0_30_] = init30_(i0_30_)
                        d_1086_v63_ = nw200_
                        d_1088_v64_: _dafny.Seq
                        d_1088_v64_ = _dafny.SeqWithoutIsStrInference([(self).f18])
                        index188_ = default__.safeIndex(484, (d_1086_v63_).length(0))
                        (d_1086_v63_)[index188_] = d_1088_v64_
                        index189_ = default__.safeIndex(484, (d_1086_v63_).length(0))
                        (d_1086_v63_)[index189_] = (d_1088_v64_) + (d_1088_v64_)
                    d_1089_v65_: str
                    d_1089_v65_ = _dafny.CodePoint('a')
                    d_1090_v66_: D5
                    d_1090_v66_ = D5_DC12((self).f19, (self).f18, d_1089_v65_)
                    d_1091_v67_: D5
                    d_1091_v67_ = D5_DC13(d_1090_v66_)
                    rhs144_ = d_1091_v67_
                    rhs145_ = (self).f18
                    rhs146_ = not(not(not(default__.fm2(globalState))))
                    lhs131_ = globalState
                    d_1091_v67_ = rhs144_
                    r1 = rhs145_
                    lhs131_.f1 = rhs146_
                    d_1092_v68_: _dafny.Array
                    def lambda51_(d_1093_i7_):
                        return (d_1093_i7_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))))

                    init31_ = lambda51_
                    nw201_ = _dafny.Array(None, 18)
                    for i0_31_ in range(nw201_.length(0)):
                        nw201_[i0_31_] = init31_(i0_31_)
                    d_1092_v68_ = nw201_
                    d_1094_v69_: D3
                    d_1094_v69_ = D3_DC6(d_1092_v68_)
                    d_1095_v70_: _dafny.Array
                    nw202_ = _dafny.Array(None, 29)
                    nw202_[int(0)] = D3_DC6(d_1092_v68_)
                    nw202_[int(1)] = d_1094_v69_
                    nw202_[int(2)] = d_1094_v69_
                    nw202_[int(3)] = d_1094_v69_
                    nw202_[int(4)] = d_1094_v69_
                    nw202_[int(5)] = d_1094_v69_
                    nw202_[int(6)] = d_1094_v69_
                    nw202_[int(7)] = d_1094_v69_
                    nw202_[int(8)] = d_1094_v69_
                    nw202_[int(9)] = d_1094_v69_
                    nw202_[int(10)] = d_1094_v69_
                    nw202_[int(11)] = d_1094_v69_
                    nw202_[int(12)] = d_1094_v69_
                    nw202_[int(13)] = d_1094_v69_
                    nw202_[int(14)] = d_1094_v69_
                    nw202_[int(15)] = d_1094_v69_
                    nw202_[int(16)] = d_1094_v69_
                    nw202_[int(17)] = d_1094_v69_
                    nw202_[int(18)] = d_1094_v69_
                    nw202_[int(19)] = d_1094_v69_
                    nw202_[int(20)] = d_1094_v69_
                    nw202_[int(21)] = d_1094_v69_
                    nw202_[int(22)] = d_1094_v69_
                    nw202_[int(23)] = d_1094_v69_
                    nw202_[int(24)] = D3_DC6(d_1092_v68_)
                    nw202_[int(25)] = d_1094_v69_
                    nw202_[int(26)] = d_1094_v69_
                    nw202_[int(27)] = d_1094_v69_
                    nw202_[int(28)] = d_1094_v69_
                    d_1095_v70_ = nw202_
                    d_1096_v71_: _dafny.Map
                    d_1096_v71_ = _dafny.Map({not((self).f18): d_1095_v70_})
                    d_1097_v72_: _dafny.Set
                    d_1097_v72_ = _dafny.Set({((d_1096_v71_)[(self).f18] if ((self).f18) in (d_1096_v71_) else d_1095_v70_), d_1095_v70_})
                    d_1098_v73_: D9
                    d_1098_v73_ = D9_DC23(d_1097_v72_, d_1041_v35_, (self).f18, d_1092_v68_, d_1089_v65_)
                    d_1099_v74_: D9
                    d_1099_v74_ = D9_DC25(d_1098_v73_)
                    d_1100_v75_: D9
                    d_1100_v75_ = D9_DC25(d_1099_v74_)
                    d_1101_v76_: _dafny.MultiSet
                    d_1101_v76_ = _dafny.MultiSet([d_1100_v75_])
                    d_1101_v76_ = _dafny.MultiSet([d_1100_v75_, d_1100_v75_, d_1100_v75_])
                    pass
            pass
        d_1102_v77_: _dafny.Array
        nw203_ = _dafny.Array(_dafny.Array(None, 0), 18)
        d_1102_v77_ = nw203_
        d_1102_v77_ = d_1102_v77_
        d_1103_v78_: D12
        d_1103_v78_ = D12_DC35()
        d_1104_v79_: _dafny.Array
        def lambda52_(d_1105_i8_):
            return D8_DC20((self).f19, (self).f19, (0) - (len(_dafny.Map({(self).f18: (self).f18}))))

        init32_ = lambda52_
        nw204_ = _dafny.Array(None, 9)
        for i0_32_ in range(nw204_.length(0)):
            nw204_[i0_32_] = init32_(i0_32_)
        d_1104_v79_ = nw204_
        d_1106_v80_: D14
        d_1106_v80_ = D14_DC38(d_1001_v0_)
        d_1107_v81_: _dafny.Map
        d_1107_v81_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1106_v80_]): (self).f18})
        d_1108_v82_: D8
        d_1108_v82_ = D8_DC20((self).f19, len(d_1107_v81_), (self).f19)
        index190_ = default__.safeIndex(120, (d_1104_v79_).length(0))
        (d_1104_v79_)[index190_] = d_1108_v82_
        d_1109_v83_: _dafny.Seq
        d_1109_v83_ = _dafny.SeqWithoutIsStrInference([True, True, ((self).f18 if (self).f18 else (self).f18), ((self).f18) or ((self).f18), (self).f18])
        d_1110_v84_: _dafny.Seq
        d_1110_v84_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([(self).f18, True])) + (_dafny.SeqWithoutIsStrInference([(self).f18])), d_1109_v83_, d_1109_v83_, (d_1109_v83_) + (d_1109_v83_), _dafny.SeqWithoutIsStrInference([(self).f18, (d_1109_v83_)[default__.safeIndex((self).f19, len(d_1109_v83_))], (self).f18, default__.fm2(globalState), (self).f18])])
        d_1111_v85_: D10
        d_1111_v85_ = D10_DC28(default__.fm1((self).f18, globalState))
        d_1112_v86_: _dafny.Set
        d_1112_v86_ = _dafny.Set({d_1111_v85_})
        d_1113_v87_: _dafny.Set
        d_1113_v87_ = _dafny.Set({_dafny.Set({D10_DC28((D5_DC11(65)).cf14), d_1111_v85_, d_1111_v85_, d_1111_v85_}), d_1112_v86_})
        index191_ = default__.safeIndex(120, (d_1104_v79_).length(0))
        rhs147_ = D12_DC35()
        rhs148_ = (self).f18
        rhs149_ = d_1108_v82_
        rhs150_ = (d_1110_v84_)[default__.safeIndex((0) - (len(d_1113_v87_)), len(d_1110_v84_))]
        rhs151_ = default__.fm2(globalState)
        lhs132_ = d_1104_v79_
        lhs133_ = default__.safeIndex(120, (d_1104_v79_).length(0))
        lhs134_ = globalState
        d_1103_v78_ = rhs147_
        r1 = rhs148_
        lhs132_[lhs133_] = rhs149_
        d_1109_v83_ = rhs150_
        lhs134_.f1 = rhs151_
        r0 = default__.safeDivisionInt((self).f19, (self).f19)
        r1 = not((self).f18)
        return r0, r1

    def m2(self, p0, globalState):
        d_1114_v0_: _dafny.Seq
        d_1114_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elwpjwmet"))
        rhs152_ = default__.safeDivisionInt(default__.safeModuloInt((self).f19, (self).f19), len(d_1114_v0_))
        rhs153_ = p0
        lhs135_ = globalState
        lhs136_ = globalState
        lhs135_.f6 = rhs152_
        lhs136_.f1 = rhs153_
        (globalState).f1 = ((1 if (self).f18 else (self).f19)) > (333)
        d_1115_v1_: _dafny.MultiSet
        d_1115_v1_ = _dafny.MultiSet([(self).f19])
        d_1116_v2_: str
        d_1116_v2_ = _dafny.CodePoint('c')
        d_1117_v3_: _dafny.Map
        d_1117_v3_ = _dafny.Map({((d_1115_v1_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rccywrduq")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rccywrduq")))) in (d_1115_v1_) else -967): d_1116_v2_})
        d_1118_v4_: _dafny.Map
        d_1118_v4_ = _dafny.Map({((d_1117_v3_)[(self).f19] if ((self).f19) in (d_1117_v3_) else d_1116_v2_): (0) - ((self).f19)})
        d_1118_v4_ = (d_1118_v4_).set((d_1116_v2_ if p0 else d_1116_v2_), (0) - (-416))
        d_1119_i0_: int
        d_1119_i0_ = 0
        with _dafny.label("8"):
            while True:
                with _dafny.c_label("8"):
                    if (d_1119_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_1119_i0_ = (d_1119_i0_) + (1)
                    (globalState).f6 = (self).f19
                    d_1120_v5_: _dafny.Map
                    d_1120_v5_ = _dafny.Map({((self).f18 if (self).f18 else True): (0) - ((self).f19)})
                    d_1120_v5_ = (d_1120_v5_).set((self).f18, (self).f19)
                    (globalState).f6 = 551
                    d_1121_v6_: _dafny.Seq
                    d_1121_v6_ = _dafny.SeqWithoutIsStrInference([576, (self).f19, (0) - ((self).f19), (self).f19])
                    d_1122_v7_: _dafny.Set
                    d_1122_v7_ = _dafny.Set({d_1121_v6_, (d_1121_v6_).set(default__.safeIndex((self).f19, len(d_1121_v6_)), (self).f19), default__.fm32((0) - ((self).f19), (self).f19, (self).f18, globalState)})
                    (globalState).f1 = (d_1122_v7_).issubset(_dafny.Set({d_1121_v6_, _dafny.SeqWithoutIsStrInference([(self).f19 for d_1123_i1_ in range(default__.abs(214))]), d_1121_v6_, d_1121_v6_, d_1121_v6_}))
                    pass
            pass
        if (self).f18:
            d_1124_v8_: _dafny.Set
            d_1124_v8_ = _dafny.Set({p0})
            d_1124_v8_ = d_1124_v8_
            if (((self).f19) < ((self).f19) if not (True) or ((self).f18) else False):
                (self).m13(globalState)
                d_1114_v0_ = (d_1114_v0_) + ((d_1114_v0_) + ((d_1114_v0_).set(default__.safeIndex((self).f19, len(d_1114_v0_)), _dafny.CodePoint('m'))))
                d_1125_v9_: D3
                d_1125_v9_ = D3_DC7((self).f19)
                d_1126_v10_: _dafny.Map
                d_1126_v10_ = _dafny.Map({p0: (self).f19})
                d_1127_v11_: _dafny.Array
                nw205_ = _dafny.Array(None, 11)
                nw205_[int(0)] = (0) - (default__.safeDivisionInt(-104, (self).f19))
                nw205_[int(1)] = (self).f19
                nw205_[int(2)] = (self).f19
                nw205_[int(3)] = (self).f19
                nw205_[int(4)] = (self).f19
                nw205_[int(5)] = len(_dafny.SeqWithoutIsStrInference([p0, p0]))
                nw205_[int(6)] = (d_1125_v9_).cf7
                nw205_[int(7)] = ((self).f19) - ((self).f19)
                nw205_[int(8)] = len(d_1114_v0_)
                nw205_[int(9)] = (self).f19
                nw205_[int(10)] = (len(d_1126_v10_)) + (523)
                d_1127_v11_ = nw205_
                index192_ = default__.safeIndex(222, (d_1127_v11_).length(0))
                (d_1127_v11_)[index192_] = (self).f19
                d_1128_v12_: _dafny.Seq
                d_1128_v12_ = _dafny.SeqWithoutIsStrInference([((self).f19) - ((0) - ((self).f19)), (self).f19])
                index193_ = default__.safeIndex(222, (d_1127_v11_).length(0))
                (d_1127_v11_)[index193_] = (d_1128_v12_)[default__.safeIndex((len(d_1114_v0_) if (self).f18 else len(d_1114_v0_)), len(d_1128_v12_))]
                (globalState).f6 = (self).f19
                d_1129_v14_: _dafny.Map
                d_1129_v14_ = _dafny.Map({(d_1127_v11_)[default__.safeIndex(222, (d_1127_v11_).length(0))]: p0})
                def iife149_():
                    coll43_ = _dafny.Map()
                    compr_43_: int
                    for compr_43_ in (d_1129_v14_).keys.Elements:
                        d_1130_v13_: int = compr_43_
                        if (d_1130_v13_) in (d_1129_v14_):
                            coll43_[default__.safeModuloInt(d_1130_v13_, (self).f19)] = d_1116_v2_
                    return _dafny.Map(coll43_)
                def iife150_():
                    coll44_ = _dafny.Map()
                    compr_44_: int
                    for compr_44_ in (d_1129_v14_).keys.Elements:
                        d_1131_v13_: int = compr_44_
                        if (d_1131_v13_) in (d_1129_v14_):
                            coll44_[default__.safeModuloInt(d_1131_v13_, (self).f19)] = d_1116_v2_
                    return _dafny.Map(coll44_)
                d_1116_v2_ = ((d_1117_v3_)[len(iife149_()
                )] if (len(iife150_()
                )) in (d_1117_v3_) else d_1116_v2_)
            elif True:
                d_1132_v16_: _dafny.Array
                def lambda53_(d_1133_i2_):
                    return D0_DC2(D0_DC1((self).f19))

                init33_ = lambda53_
                nw206_ = _dafny.Array(None, 27)
                for i0_33_ in range(nw206_.length(0)):
                    nw206_[i0_33_] = init33_(i0_33_)
                d_1132_v16_ = nw206_
                d_1134_v17_: _dafny.Array
                def lambda54_(d_1135_i3_):
                    return default__.safeDivisionInt(d_1135_i3_, (self).f19)

                init34_ = lambda54_
                nw207_ = _dafny.Array(None, 22)
                for i0_34_ in range(nw207_.length(0)):
                    nw207_[i0_34_] = init34_(i0_34_)
                d_1134_v17_ = nw207_
                def iife151_():
                    coll45_ = _dafny.Set()
                    compr_45_: int
                    for compr_45_ in (_dafny.SeqWithoutIsStrInference([default__.fm1((self).f18, globalState)])).Elements:
                        d_1136_v15_: int = compr_45_
                        if (d_1136_v15_) in (_dafny.SeqWithoutIsStrInference([default__.fm1((self).f18, globalState)])):
                            coll45_ = coll45_.union(_dafny.Set([(d_1136_v15_) - (701)]))
                    return _dafny.Set(coll45_)
                default__.m0(iife151_()
                , d_1132_v16_, d_1134_v17_, d_1114_v0_, globalState)
                d_1137_v18_: D3
                d_1137_v18_ = D3_DC6(d_1134_v17_)
                d_1138_v19_: _dafny.Seq
                d_1138_v19_ = _dafny.SeqWithoutIsStrInference([d_1137_v18_, D3_DC6(d_1134_v17_), d_1137_v18_, d_1137_v18_, D3_DC6(d_1134_v17_)])
                d_1139_v20_: C0
                nw208_ = C0()
                nw208_.ctor__(((d_1138_v19_) + (d_1138_v19_)).set(default__.safeIndex(52, len((d_1138_v19_) + (d_1138_v19_))), d_1137_v18_), (self).f18)
                d_1139_v20_ = nw208_
                index194_ = default__.safeIndex(392, (d_1134_v17_).length(0))
                (d_1134_v17_)[index194_] = (self).f19
                index195_ = default__.safeIndex(392, (d_1134_v17_).length(0))
                (d_1134_v17_)[index195_] = default__.fm1((253) == ((self).f19), globalState)
                d_1140_v21_: _dafny.Seq
                d_1140_v21_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1141_v22_: D11
                d_1141_v22_ = D11_DC31(d_1140_v21_, (d_1134_v17_)[default__.safeIndex(392, (d_1134_v17_).length(0))], (self).f19, default__.fm1(True, globalState), (self).f19)
                d_1142_v23_: _dafny.Map
                d_1142_v23_ = _dafny.Map({(d_1139_v20_).f23: not((self).f18)})
                d_1143_v24_: _dafny.Array
                nw209_ = _dafny.Array(None, 13)
                nw209_[int(0)] = False
                nw209_[int(1)] = (len(d_1114_v0_)) < ((self).f19)
                nw209_[int(2)] = (d_1139_v20_).f23
                nw209_[int(3)] = (d_1139_v20_).f23
                nw209_[int(4)] = (self).f18
                nw209_[int(5)] = p0
                nw209_[int(6)] = (d_1139_v20_).f23
                nw209_[int(7)] = True
                nw209_[int(8)] = (d_1139_v20_).fm17(len((d_1141_v22_).cf54), (0) - (len(d_1142_v23_)), not(not(p0)), globalState)
                nw209_[int(9)] = (d_1139_v20_).f23
                nw209_[int(10)] = p0
                nw209_[int(11)] = not(False)
                nw209_[int(12)] = p0
                d_1143_v24_ = nw209_
                index196_ = default__.safeIndex(932, (d_1143_v24_).length(0))
                (d_1143_v24_)[index196_] = (self).f18
                index197_ = default__.safeIndex(932, (d_1143_v24_).length(0))
                (d_1143_v24_)[index197_] = (d_1139_v20_).fm17(default__.safeModuloInt((0) - (default__.fm1((self).f18, globalState)), (d_1134_v17_)[default__.safeIndex(392, (d_1134_v17_).length(0))]), (self).f19, p0, globalState)
                d_1144_v25_: _dafny.Array
                nw210_ = _dafny.Array(_dafny.Map({}), 21)
                d_1144_v25_ = nw210_
                d_1144_v25_ = d_1144_v25_
            d_1145_v26_: _dafny.Array
            nw211_ = _dafny.Array(False, 21)
            d_1145_v26_ = nw211_
            index198_ = default__.safeIndex(741, (d_1145_v26_).length(0))
            (d_1145_v26_)[index198_] = not(not(p0))
            index199_ = default__.safeIndex(741, (d_1145_v26_).length(0))
            rhs154_ = (self).f18
            rhs155_ = default__.safeModuloInt((self).f19, ((self).f19) * ((0) - ((self).f19)))
            rhs156_ = d_1114_v0_
            rhs157_ = (d_1114_v0_) + (d_1114_v0_)
            lhs137_ = d_1145_v26_
            lhs138_ = default__.safeIndex(741, (d_1145_v26_).length(0))
            lhs139_ = globalState
            lhs137_[lhs138_] = rhs154_
            lhs139_.f6 = rhs155_
            d_1114_v0_ = rhs156_
            d_1114_v0_ = rhs157_
            (globalState).f6 = (self).f19
            d_1146_v27_: _dafny.Array
            nw212_ = _dafny.Array(_dafny.CodePoint('D'), 20)
            d_1146_v27_ = nw212_
            index200_ = default__.safeIndex(141, (d_1146_v27_).length(0))
            (d_1146_v27_)[index200_] = d_1116_v2_
            index201_ = default__.safeIndex(141, (d_1146_v27_).length(0))
            (d_1146_v27_)[index201_] = (d_1114_v0_)[default__.safeIndex((self).f19, len(d_1114_v0_))]
        elif True:
            (globalState).f1 = p0
            d_1147_v28_: _dafny.Array
            nw213_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_1147_v28_ = nw213_
            d_1148_v29_: _dafny.Map
            d_1148_v29_ = _dafny.Map({p0: d_1147_v28_})
            d_1149_v30_: _dafny.Array
            nw214_ = _dafny.Array(None, 24)
            nw214_[int(0)] = (d_1147_v28_ if p0 else d_1147_v28_)
            nw214_[int(1)] = ((d_1148_v29_)[False] if (False) in (d_1148_v29_) else d_1147_v28_)
            nw214_[int(2)] = d_1147_v28_
            nw214_[int(3)] = d_1147_v28_
            nw214_[int(4)] = d_1147_v28_
            nw214_[int(5)] = d_1147_v28_
            nw214_[int(6)] = d_1147_v28_
            nw214_[int(7)] = d_1147_v28_
            nw214_[int(8)] = d_1147_v28_
            nw214_[int(9)] = d_1147_v28_
            nw214_[int(10)] = d_1147_v28_
            nw214_[int(11)] = d_1147_v28_
            nw214_[int(12)] = d_1147_v28_
            nw214_[int(13)] = d_1147_v28_
            nw214_[int(14)] = d_1147_v28_
            nw214_[int(15)] = d_1147_v28_
            nw214_[int(16)] = d_1147_v28_
            nw214_[int(17)] = d_1147_v28_
            nw214_[int(18)] = d_1147_v28_
            nw214_[int(19)] = d_1147_v28_
            nw214_[int(20)] = d_1147_v28_
            nw214_[int(21)] = d_1147_v28_
            nw214_[int(22)] = d_1147_v28_
            nw214_[int(23)] = d_1147_v28_
            d_1149_v30_ = nw214_
            index202_ = default__.safeIndex(292, (d_1149_v30_).length(0))
            (d_1149_v30_)[index202_] = d_1147_v28_
            index203_ = default__.safeIndex(292, (d_1149_v30_).length(0))
            (d_1149_v30_)[index203_] = d_1147_v28_
            if (p0) and (p0):
                d_1147_v28_ = (d_1149_v30_)[default__.safeIndex(292, (d_1149_v30_).length(0))]
                d_1150_v31_: _dafny.Seq
                d_1150_v31_ = _dafny.SeqWithoutIsStrInference([(self).f19])
                d_1151_v32_: _dafny.Map
                d_1151_v32_ = _dafny.Map({not(True): (self).f18})
                d_1152_v33_: _dafny.Array
                nw215_ = _dafny.Array(None, 7)
                nw215_[int(0)] = (self).f19
                nw215_[int(1)] = len(_dafny.SeqWithoutIsStrInference([len(d_1150_v31_)]))
                nw215_[int(2)] = (self).f19
                nw215_[int(3)] = len(d_1114_v0_)
                nw215_[int(4)] = (self).f19
                nw215_[int(5)] = len(d_1150_v31_)
                nw215_[int(6)] = len(d_1151_v32_)
                d_1152_v33_ = nw215_
                d_1153_v34_: D3
                d_1153_v34_ = D3_DC6(d_1152_v33_)
                d_1154_v35_: _dafny.Seq
                d_1154_v35_ = _dafny.SeqWithoutIsStrInference([d_1153_v34_, d_1153_v34_, d_1153_v34_])
                d_1155_v36_: C0
                nw216_ = C0()
                nw216_.ctor__((d_1154_v35_) + (d_1154_v35_), p0)
                d_1155_v36_ = nw216_
                d_1156_v37_: _dafny.Seq
                d_1156_v37_ = _dafny.SeqWithoutIsStrInference([(d_1155_v36_).f23])
                d_1157_v38_: _dafny.Array
                nw217_ = _dafny.Array(None, 5)
                nw217_[int(0)] = (d_1155_v36_).f23
                nw217_[int(1)] = p0
                nw217_[int(2)] = (d_1156_v37_)[default__.safeIndex((self).f19, len(d_1156_v37_))]
                nw217_[int(3)] = (self).f18
                nw217_[int(4)] = p0
                d_1157_v38_ = nw217_
                d_1158_v39_: _dafny.Map
                d_1158_v39_ = _dafny.Map({d_1117_v3_: d_1157_v38_})
                d_1158_v39_ = d_1158_v39_
                (globalState).f1 = not (False) or ((d_1155_v36_).f23)
                d_1159_v40_: T0
                nw218_ = C1()
                nw218_.ctor__((self).f19)
                d_1159_v40_ = nw218_
                d_1160_v41_: D7
                d_1160_v41_ = D7_DC15(d_1159_v40_)
                pat_let_tv29_ = d_1160_v41_
                d_1161_v42_: _dafny.Array
                nw219_ = _dafny.Array(None, 17)
                nw219_[int(0)] = d_1160_v41_
                nw219_[int(1)] = d_1160_v41_
                def iife152_(_pat_let53_0):
                    def iife153_(d_1162_dt__update__tmp_h0_):
                        def iife154_(_pat_let54_0):
                            def iife155_(d_1163_dt__update_hcf20_h0_):
                                return D7_DC15(d_1163_dt__update_hcf20_h0_)
                            return iife155_(_pat_let54_0)
                        return iife154_((pat_let_tv29_).cf20)
                    return iife153_(_pat_let53_0)
                nw219_[int(2)] = iife152_(d_1160_v41_)
                nw219_[int(3)] = D7_DC15(d_1159_v40_)
                nw219_[int(4)] = D7_DC15(d_1159_v40_)
                nw219_[int(5)] = d_1160_v41_
                nw219_[int(6)] = d_1160_v41_
                nw219_[int(7)] = d_1160_v41_
                nw219_[int(8)] = d_1160_v41_
                nw219_[int(9)] = D7_DC15(d_1159_v40_)
                nw219_[int(10)] = d_1160_v41_
                nw219_[int(11)] = D7_DC15(d_1159_v40_)
                nw219_[int(12)] = d_1160_v41_
                nw219_[int(13)] = d_1160_v41_
                nw219_[int(14)] = d_1160_v41_
                nw219_[int(15)] = d_1160_v41_
                nw219_[int(16)] = d_1160_v41_
                d_1161_v42_ = nw219_
                d_1164_v43_: _dafny.Seq
                d_1164_v43_ = _dafny.SeqWithoutIsStrInference([d_1161_v42_])
                d_1165_v44_: _dafny.Array
                nw220_ = _dafny.Array(None, 25)
                nw220_[int(0)] = d_1161_v42_
                nw220_[int(1)] = d_1161_v42_
                nw220_[int(2)] = d_1161_v42_
                nw220_[int(3)] = d_1161_v42_
                nw220_[int(4)] = d_1161_v42_
                nw220_[int(5)] = d_1161_v42_
                nw220_[int(6)] = d_1161_v42_
                nw220_[int(7)] = d_1161_v42_
                nw220_[int(8)] = d_1161_v42_
                nw220_[int(9)] = d_1161_v42_
                nw220_[int(10)] = d_1161_v42_
                nw220_[int(11)] = (d_1164_v43_)[default__.safeIndex((self).f19, len(d_1164_v43_))]
                nw220_[int(12)] = d_1161_v42_
                nw220_[int(13)] = d_1161_v42_
                nw220_[int(14)] = d_1161_v42_
                nw220_[int(15)] = d_1161_v42_
                nw220_[int(16)] = d_1161_v42_
                nw220_[int(17)] = d_1161_v42_
                nw220_[int(18)] = d_1161_v42_
                nw220_[int(19)] = d_1161_v42_
                nw220_[int(20)] = d_1161_v42_
                nw220_[int(21)] = d_1161_v42_
                nw220_[int(22)] = d_1161_v42_
                nw220_[int(23)] = d_1161_v42_
                nw220_[int(24)] = d_1161_v42_
                d_1165_v44_ = nw220_
                d_1166_v45_: _dafny.Map
                d_1166_v45_ = _dafny.Map({(self).f19: d_1165_v44_})
                d_1167_v46_: D5
                d_1167_v46_ = D5_DC12((self).f19, default__.fm2(globalState), d_1116_v2_)
                d_1168_v47_: _dafny.Map
                d_1168_v47_ = _dafny.Map({D10_DC26(((d_1166_v45_)[len(_dafny.Map({(d_1167_v46_).cf16: (self).f19}))] if (len(_dafny.Map({(d_1167_v46_).cf16: (self).f19}))) in (d_1166_v45_) else d_1165_v44_)): ((self).f19) * ((self).f19)})
                d_1169_v48_: D10
                d_1169_v48_ = D10_DC26(d_1165_v44_)
                d_1168_v47_ = (d_1168_v47_).set(d_1169_v48_, len(_dafny.SeqWithoutIsStrInference([d_1116_v2_ for d_1170_i4_ in range(default__.abs(281))])))
            elif True:
                d_1171_v49_: _dafny.Map
                d_1171_v49_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpswinn"))) != (d_1114_v0_): ((0) - ((self).f19)) >= (118)})
                (globalState).f1 = ((d_1171_v49_)[not((self).f18)] if (not((self).f18)) in (d_1171_v49_) else not(True))
                d_1172_v50_: _dafny.Map
                d_1172_v50_ = _dafny.Map({(d_1115_v1_).cardinality: p0})
                d_1173_v51_: _dafny.Array
                nw221_ = _dafny.Array(None, 29)
                nw221_[int(0)] = p0
                nw221_[int(1)] = (self).f18
                nw221_[int(2)] = True
                nw221_[int(3)] = (self).f18
                nw221_[int(4)] = (self).f18
                nw221_[int(5)] = not((self).f18)
                nw221_[int(6)] = (self).f18
                nw221_[int(7)] = (self).f18
                nw221_[int(8)] = p0
                nw221_[int(9)] = p0
                nw221_[int(10)] = (self).f18
                nw221_[int(11)] = (self).f18
                nw221_[int(12)] = (self).f18
                nw221_[int(13)] = (self).f18
                nw221_[int(14)] = (self).f18
                nw221_[int(15)] = (self).f18
                nw221_[int(16)] = p0
                nw221_[int(17)] = (self).f18
                nw221_[int(18)] = (self).f18
                nw221_[int(19)] = (self).f18
                nw221_[int(20)] = p0
                nw221_[int(21)] = (self).f18
                nw221_[int(22)] = p0
                nw221_[int(23)] = (self).f18
                nw221_[int(24)] = (self).f18
                nw221_[int(25)] = p0
                nw221_[int(26)] = p0
                nw221_[int(27)] = default__.fm2(globalState)
                nw221_[int(28)] = ((d_1172_v50_)[(self).f19] if ((self).f19) in (d_1172_v50_) else p0)
                d_1173_v51_ = nw221_
                d_1174_v52_: _dafny.Array
                nw222_ = _dafny.Array(None, 1)
                nw222_[int(0)] = d_1173_v51_
                d_1174_v52_ = nw222_
                index204_ = default__.safeIndex(905, (d_1174_v52_).length(0))
                (d_1174_v52_)[index204_] = d_1173_v51_
                index205_ = default__.safeIndex(905, (d_1174_v52_).length(0))
                (d_1174_v52_)[index205_] = d_1173_v51_
                (globalState).f1 = p0
                (globalState).f1 = False
                d_1172_v50_ = (d_1172_v50_).set((self).f19, ((self).f18 if (self).f18 else p0))
            (globalState).f1 = (self).f18
            d_1175_v53_: _dafny.Seq
            d_1175_v53_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1176_v54_: D0
            d_1176_v54_ = D0_DC0(not(p0))
            d_1177_v55_: _dafny.Seq
            d_1177_v55_ = _dafny.SeqWithoutIsStrInference([d_1176_v54_])
            d_1178_v56_: _dafny.Map
            d_1178_v56_ = _dafny.Map({(default__.fm35(d_1175_v53_, globalState)) + (d_1177_v55_): (self).f19})
            d_1178_v56_ = (d_1178_v56_).set((d_1177_v55_ if (self).f18 else d_1177_v55_), (self).f19)
        d_1179_v57_: _dafny.MultiSet
        d_1179_v57_ = _dafny.MultiSet([p0, p0])
        d_1180_v58_: _dafny.MultiSet
        d_1180_v58_ = _dafny.MultiSet([d_1116_v2_, d_1116_v2_, d_1116_v2_, d_1116_v2_, d_1116_v2_])
        d_1181_v59_: _dafny.Map
        d_1181_v59_ = _dafny.Map({d_1179_v57_: (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f19])) for d_1182_i5_ in range(default__.abs(115))])) + (_dafny.SeqWithoutIsStrInference([(d_1180_v58_).cardinality, (self).f19]))})
        d_1183_v60_: _dafny.Map
        d_1183_v60_ = _dafny.Map({(self).f18: default__.safeModuloInt(151, (self).f19)})
        d_1184_v61_: D12
        d_1184_v61_ = D12_DC34(_dafny.Map({(self).f19: (self).f19}))
        rhs158_ = d_1181_v59_
        rhs159_ = default__.fm36(d_1116_v2_, d_1184_v61_, globalState)
        d_1181_v59_ = rhs158_
        d_1183_v60_ = rhs159_

    def m11(self, p0, p1, globalState):
        r0: str = _dafny.CodePoint('D')
        (self).m13(globalState)
        d_1185_v0_: _dafny.Set
        d_1185_v0_ = _dafny.Set({p1, default__.fm1((self).f18, globalState)})
        d_1186_v1_: _dafny.Array
        nw223_ = _dafny.Array(D0.default()(), 27)
        d_1186_v1_ = nw223_
        d_1187_v2_: _dafny.Array
        nw224_ = _dafny.Array(None, 2)
        nw224_[int(0)] = p1
        nw224_[int(1)] = p1
        d_1187_v2_ = nw224_
        default__.m0(d_1185_v0_, d_1186_v1_, d_1187_v2_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1188_i0_ in range(default__.abs(747))]), globalState)
        d_1189_v3_: _dafny.Array
        nw225_ = _dafny.Array(_dafny.Seq({}), 13)
        d_1189_v3_ = nw225_
        d_1190_v4_: _dafny.Seq
        d_1190_v4_ = _dafny.SeqWithoutIsStrInference([(self).f19, p1])
        index206_ = default__.safeIndex(519, (d_1189_v3_).length(0))
        (d_1189_v3_)[index206_] = (d_1190_v4_) + (d_1190_v4_)
        index207_ = default__.safeIndex(519, (d_1189_v3_).length(0))
        (d_1189_v3_)[index207_] = _dafny.SeqWithoutIsStrInference([p1])
        (globalState).f6 = ((0) - ((self).f19)) + (p1)
        (globalState).f1 = (self).f18
        (globalState).f6 = ((d_1189_v3_)[default__.safeIndex(519, (d_1189_v3_).length(0))])[default__.safeIndex((self).f19, len((d_1189_v3_)[default__.safeIndex(519, (d_1189_v3_).length(0))]))]
        d_1191_v5_: str
        d_1191_v5_ = _dafny.CodePoint('r')
        r0 = d_1191_v5_
        return r0

    def m12(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1 = (default__.fm21(globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klxjhxnre")))
        (globalState).f1 = (self).f18
        hi8_ = (self).f19
        for d_1192_i0_ in range(p0, hi8_):
            d_1193_v0_: _dafny.Array
            nw226_ = _dafny.Array(False, 2)
            d_1193_v0_ = nw226_
            d_1194_v1_: _dafny.Seq
            d_1194_v1_ = _dafny.SeqWithoutIsStrInference([(self).f18, True, (self).f18, (self).f18, default__.fm2(globalState)])
            index208_ = default__.safeIndex(678, (d_1193_v0_).length(0))
            (d_1193_v0_)[index208_] = (d_1194_v1_)[default__.safeIndex(len(p3), len(d_1194_v1_))]
            d_1195_v2_: _dafny.MultiSet
            d_1195_v2_ = _dafny.MultiSet([(self).f18, (self).f18])
            d_1196_v3_: _dafny.Map
            d_1196_v3_ = _dafny.Map({(d_1195_v2_).intersection(_dafny.MultiSet(d_1194_v1_)): p2})
            d_1197_v4_: _dafny.MultiSet
            d_1197_v4_ = _dafny.MultiSet([p0])
            d_1198_v5_: _dafny.Map
            d_1198_v5_ = _dafny.Map({False: -496})
            d_1199_v6_: _dafny.Map
            d_1199_v6_ = _dafny.Map({d_1192_i0_: (self).f18})
            d_1200_v7_: D7
            d_1200_v7_ = D7_DC16((self).f18, len((d_1198_v5_).set((self).f18, (self).f19)), ((d_1199_v6_)[p0] if (p0) in (d_1199_v6_) else (self).f18))
            d_1201_v8_: D13
            d_1201_v8_ = D13_DC37((self).f18, (d_1200_v7_).cf21)
            d_1202_v10_: _dafny.Seq
            d_1202_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f18]), _dafny.MultiSet(d_1194_v1_)])
            d_1203_v11_: _dafny.Set
            d_1203_v11_ = _dafny.Set({(self).f18})
            index209_ = default__.safeIndex(678, (d_1193_v0_).length(0))
            def iife156_():
                coll46_ = _dafny.Map()
                compr_46_: _dafny.MultiSet
                for compr_46_ in (d_1202_v10_).Elements:
                    d_1204_v9_: _dafny.MultiSet = compr_46_
                    if (d_1204_v9_) in (d_1202_v10_):
                        coll46_[d_1204_v9_] = p3
                return _dafny.Map(coll46_)
            rhs160_ = (d_1197_v4_).isdisjoint(d_1197_v4_)
            rhs161_ = (d_1196_v3_ if (d_1201_v8_).cf66 else iife156_()
            )
            rhs162_ = not((d_1203_v11_).ispropersubset(d_1203_v11_))
            rhs163_ = (p3 if ((self).f18) not in (d_1203_v11_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1205_i1_ in range(default__.abs(590))]))
            lhs140_ = d_1193_v0_
            lhs141_ = default__.safeIndex(678, (d_1193_v0_).length(0))
            lhs140_[lhs141_] = rhs160_
            d_1196_v3_ = rhs161_
            r0 = rhs162_
            r1 = rhs163_
            d_1206_v12_: C2
            nw227_ = C2()
            nw227_.ctor__(p1, (self).f19, self.f17)
            d_1206_v12_ = nw227_
            d_1207_v13_: str
            d_1207_v13_ = _dafny.CodePoint('r')
            d_1208_v14_: _dafny.Map
            d_1208_v14_ = _dafny.Map({len(d_1198_v5_): len(_dafny.Set({d_1207_v13_, d_1207_v13_}))})
            d_1209_v15_: _dafny.Seq
            d_1209_v15_ = _dafny.SeqWithoutIsStrInference([len(p2)])
            d_1210_v16_: _dafny.Array
            nw228_ = _dafny.Array(None, 17)
            nw228_[int(0)] = d_1192_i0_
            nw228_[int(1)] = (d_1206_v12_).f20
            nw228_[int(2)] = (d_1206_v12_).f20
            nw228_[int(3)] = (d_1206_v12_).f20
            nw228_[int(4)] = (len(p3)) - (859)
            nw228_[int(5)] = (d_1206_v12_).f20
            nw228_[int(6)] = (d_1206_v12_).f20
            nw228_[int(7)] = (len(_dafny.SeqWithoutIsStrInference([d_1193_v0_]))) + (len(d_1203_v11_))
            nw228_[int(8)] = p1
            nw228_[int(9)] = (p1) * (default__.fm1((d_1193_v0_)[default__.safeIndex(678, (d_1193_v0_).length(0))], globalState))
            nw228_[int(10)] = (len(_dafny.SeqWithoutIsStrInference([(self).f18]))) * (d_1192_i0_)
            nw228_[int(11)] = (0) - (d_1192_i0_)
            nw228_[int(12)] = len(d_1203_v11_)
            nw228_[int(13)] = p1
            nw228_[int(14)] = (-761) * (len(d_1208_v14_))
            nw228_[int(15)] = ((self).f19) * ((self).f19)
            nw228_[int(16)] = (len(d_1209_v15_) if not((self).f18) else (self).f19)
            d_1210_v16_ = nw228_
            d_1211_v17_: D11
            d_1211_v17_ = D11_DC31(d_1194_v1_, d_1192_i0_, p1, p1, p0)
            index210_ = default__.safeIndex(327, (d_1210_v16_).length(0))
            (d_1210_v16_)[index210_] = (d_1211_v17_).cf56
            index211_ = default__.safeIndex(327, (d_1210_v16_).length(0))
            (d_1210_v16_)[index211_] = d_1192_i0_
            index212_ = default__.safeIndex(327, (d_1210_v16_).length(0))
            (d_1210_v16_)[index212_] = 515
        (globalState).f6 = (0) - (p1)
        (globalState).f6 = p1
        if (self).f18:
            d_1212_v18_: _dafny.Array
            nw229_ = _dafny.Array(int(0), 22)
            d_1212_v18_ = nw229_
            d_1213_v19_: _dafny.Map
            d_1213_v19_ = _dafny.Map({(-157) != (p0): d_1212_v18_})
            (globalState).f9 = ((d_1213_v19_)[not((self).f18)] if (not((self).f18)) in (d_1213_v19_) else d_1212_v18_)
            index213_ = default__.safeIndex(376, (d_1212_v18_).length(0))
            (d_1212_v18_)[index213_] = p0
            index214_ = default__.safeIndex(376, (d_1212_v18_).length(0))
            (d_1212_v18_)[index214_] = (self).f19
            d_1214_v20_: _dafny.MultiSet
            d_1214_v20_ = _dafny.MultiSet([(d_1212_v18_)[default__.safeIndex(376, (d_1212_v18_).length(0))], p1])
            d_1215_v22_: _dafny.Map
            d_1215_v22_ = _dafny.Map({p0: (self).f18})
            d_1216_v23_: _dafny.Set
            def iife157_():
                coll47_ = _dafny.Set()
                compr_47_: int
                for compr_47_ in (d_1214_v20_).Elements:
                    d_1217_v21_: int = compr_47_
                    if (d_1217_v21_) in (d_1214_v20_):
                        coll47_ = coll47_.union(_dafny.Set([(d_1217_v21_) + (784)]))
                return _dafny.Set(coll47_)
            d_1216_v23_ = _dafny.Set({(iife157_()
            ).ispropersubset(_dafny.Set({len(d_1215_v22_)}))})
            d_1216_v23_ = d_1216_v23_
            d_1218_v24_: D3
            d_1218_v24_ = D3_DC6(d_1212_v18_)
            d_1219_v25_: _dafny.Seq
            d_1219_v25_ = _dafny.SeqWithoutIsStrInference([D3_DC6(d_1212_v18_), d_1218_v24_, d_1218_v24_, d_1218_v24_, d_1218_v24_])
            d_1220_v26_: C0
            nw230_ = C0()
            nw230_.ctor__((d_1219_v25_) + (d_1219_v25_), (self).f18)
            d_1220_v26_ = nw230_
            d_1221_v27_: _dafny.Seq
            d_1221_v27_ = _dafny.SeqWithoutIsStrInference([p2])
            if ((d_1221_v27_) + (_dafny.SeqWithoutIsStrInference([p3 for d_1222_i2_ in range(default__.abs(-375))]))) == (d_1221_v27_):
                d_1223_v28_: _dafny.Seq
                d_1223_v28_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1224_v29_: _dafny.Map
                d_1224_v29_ = _dafny.Map({not((self).f18): d_1223_v28_})
                d_1225_v30_: _dafny.Seq
                d_1225_v30_ = _dafny.SeqWithoutIsStrInference([((d_1224_v29_)[(self).f18] if ((self).f18) in (d_1224_v29_) else d_1223_v28_)])
                d_1225_v30_ = d_1225_v30_
                (globalState).f6 = (0) - ((0) - ((self).f19))
                d_1226_v31_: _dafny.Array
                nw231_ = _dafny.Array(_dafny.Map({}), 1)
                d_1226_v31_ = nw231_
                d_1226_v31_ = d_1226_v31_
                d_1227_v32_: _dafny.Map
                d_1227_v32_ = _dafny.Map({(d_1220_v26_).f23: (d_1212_v18_)[default__.safeIndex(376, (d_1212_v18_).length(0))]})
                d_1228_v33_: _dafny.Map
                d_1228_v33_ = _dafny.Map({d_1227_v32_: len(default__.fm28(p1, False, (d_1220_v26_).f23, globalState))})
                d_1228_v33_ = (d_1228_v33_).set(d_1227_v32_, (self).f19)
                d_1229_v34_: C1
                nw232_ = C1()
                nw232_.ctor__((721) + (len(p2)))
                d_1229_v34_ = nw232_
            elif True:
                d_1230_v35_: _dafny.Map
                d_1230_v35_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_1212_v18_)[default__.safeIndex(376, (d_1212_v18_).length(0))], (self).f19]): not((d_1220_v26_).f23)})
                index215_ = default__.safeIndex(376, (d_1212_v18_).length(0))
                def iife158_():
                    coll48_ = _dafny.Set()
                    compr_48_: _dafny.Seq
                    for compr_48_ in (d_1230_v35_).keys.Elements:
                        d_1231_v36_: _dafny.Seq = compr_48_
                        if (d_1231_v36_) in (d_1230_v35_):
                            coll48_ = coll48_.union(_dafny.Set([d_1231_v36_]))
                    return _dafny.Set(coll48_)
                (d_1212_v18_)[index215_] = len(iife158_()
                )
                (globalState).f6 = ((d_1212_v18_)[default__.safeIndex(376, (d_1212_v18_).length(0))]) - (23)
                (globalState).f1 = False
                d_1232_v37_: _dafny.Array
                nw233_ = _dafny.Array(_dafny.Array(None, 0), 23)
                d_1232_v37_ = nw233_
                index216_ = default__.safeIndex(841, (d_1232_v37_).length(0))
                (d_1232_v37_)[index216_] = d_1212_v18_
                d_1233_v38_: _dafny.Array
                nw234_ = _dafny.Array(None, 10)
                nw234_[int(0)] = d_1218_v24_
                nw234_[int(1)] = d_1218_v24_
                nw234_[int(2)] = d_1218_v24_
                nw234_[int(3)] = D3_DC6(d_1212_v18_)
                nw234_[int(4)] = d_1218_v24_
                nw234_[int(5)] = d_1218_v24_
                nw234_[int(6)] = d_1218_v24_
                nw234_[int(7)] = d_1218_v24_
                nw234_[int(8)] = d_1218_v24_
                nw234_[int(9)] = d_1218_v24_
                d_1233_v38_ = nw234_
                d_1234_v39_: _dafny.Set
                d_1234_v39_ = _dafny.Set({d_1233_v38_, d_1233_v38_})
                d_1235_v40_: str
                d_1235_v40_ = _dafny.CodePoint('n')
                d_1236_v41_: D9
                d_1236_v41_ = D9_DC23(d_1234_v39_, p2, not((d_1220_v26_).f23), d_1212_v18_, d_1235_v40_)
                index217_ = default__.safeIndex(841, (d_1232_v37_).length(0))
                (d_1232_v37_)[index217_] = (d_1236_v41_).cf40
                d_1237_v42_: _dafny.Map
                d_1237_v42_ = _dafny.Map({p0: d_1214_v20_})
                d_1238_v43_: _dafny.Seq
                d_1238_v43_ = _dafny.SeqWithoutIsStrInference([self.f17, self.f17, self.f17])
                d_1239_v44_: _dafny.Seq
                d_1239_v44_ = _dafny.SeqWithoutIsStrInference([531])
                d_1240_v45_: _dafny.Seq
                d_1240_v45_ = d_1239_v44_
                d_1241_v46_: C2
                nw235_ = C2()
                nw235_.ctor__(len(d_1237_v42_), (d_1212_v18_)[default__.safeIndex(376, (d_1212_v18_).length(0))], (d_1238_v43_)[default__.safeIndex((d_1239_v44_)[default__.safeIndex(len((d_1240_v45_)), len(d_1239_v44_))], len(d_1238_v43_))])
                d_1241_v46_ = nw235_
        elif True:
            d_1242_v47_: _dafny.Map
            d_1242_v47_ = _dafny.Map({(self).f19: p2})
            d_1243_v48_: T0
            nw236_ = C1()
            nw236_.ctor__(p1)
            d_1243_v48_ = nw236_
            d_1244_v49_: D7
            d_1244_v49_ = D7_DC15(d_1243_v48_)
            d_1245_v50_: _dafny.Map
            d_1245_v50_ = _dafny.Map({193: p0})
            d_1246_v51_: _dafny.Map
            d_1246_v51_ = _dafny.Map({d_1244_v49_: d_1245_v50_})
            d_1247_v52_: str
            d_1247_v52_ = _dafny.CodePoint('m')
            d_1242_v47_ = (d_1242_v47_).set(default__.safeModuloInt(len(((d_1246_v51_)[d_1244_v49_] if (d_1244_v49_) in (d_1246_v51_) else d_1245_v50_)), p1), (p2).set(default__.safeIndex((self).f19, len(p2)), d_1247_v52_))
            if (self).f18:
                r0 = (self).f18
                (globalState).f6 = default__.safeDivisionInt(p0, 20)
                d_1248_v53_: T1
                nw237_ = C2()
                nw237_.ctor__(len(_dafny.Map({(self).f19: (self).f18})), p1, self.f17)
                d_1248_v53_ = nw237_
                d_1249_v54_: _dafny.Map
                d_1249_v54_ = _dafny.Map({d_1248_v53_: (self).f18})
                d_1250_v55_: _dafny.Set
                d_1250_v55_ = _dafny.Set({len(d_1249_v54_), p1})
                d_1251_v56_: _dafny.Map
                d_1251_v56_ = _dafny.Map({(self).f18: d_1250_v55_})
                d_1252_v57_: _dafny.Seq
                d_1252_v57_ = _dafny.SeqWithoutIsStrInference([d_1251_v56_, d_1251_v56_])
                d_1253_v58_: _dafny.MultiSet
                d_1253_v58_ = _dafny.MultiSet([(self).f19])
                d_1254_v59_: _dafny.Map
                d_1254_v59_ = _dafny.Map({(d_1252_v57_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p0, (d_1253_v58_).cardinality])), len(d_1252_v57_))]: (self).f19})
                d_1254_v59_ = (d_1254_v59_).set(_dafny.Map({(self).f18: d_1250_v55_}), (0) - (((self).f19) - (p1)))
                d_1255_v60_: D10
                d_1255_v60_ = D10_DC28((self).f19)
                d_1256_v61_: _dafny.Map
                d_1256_v61_ = _dafny.Map({(self).f18: (self).f18})
                d_1257_v62_: D15
                d_1257_v62_ = D15_DC42(d_1250_v55_)
                d_1255_v60_ = default__.fm37(d_1256_v61_, d_1257_v62_, globalState)
                (globalState).f1 = ((self).f18 if (self).f18 else True)
            elif True:
                d_1258_v63_: _dafny.Array
                nw238_ = _dafny.Array(int(0), 15)
                d_1258_v63_ = nw238_
                d_1259_v64_: _dafny.Set
                d_1259_v64_ = _dafny.Set({d_1247_v52_})
                d_1260_v65_: _dafny.Map
                d_1260_v65_ = _dafny.Map({d_1242_v47_: d_1259_v64_})
                d_1261_v66_: _dafny.Set
                d_1261_v66_ = _dafny.Set({d_1259_v64_, ((d_1260_v65_)[d_1242_v47_] if (d_1242_v47_) in (d_1260_v65_) else d_1259_v64_), _dafny.Set({d_1247_v52_})})
                index218_ = default__.safeIndex(265, (d_1258_v63_).length(0))
                (d_1258_v63_)[index218_] = len(d_1261_v66_)
                d_1262_v67_: _dafny.Array
                nw239_ = _dafny.Array(_dafny.CodePoint('D'), 11)
                d_1262_v67_ = nw239_
                d_1263_v68_: _dafny.Array
                nw240_ = _dafny.Array(False, 3)
                d_1263_v68_ = nw240_
                index219_ = default__.safeIndex(389, (d_1263_v68_).length(0))
                (d_1263_v68_)[index219_] = (self).f18
                index220_ = default__.safeIndex(265, (d_1258_v63_).length(0))
                index221_ = default__.safeIndex(389, (d_1263_v68_).length(0))
                rhs164_ = 444
                rhs165_ = d_1262_v67_
                rhs166_ = 231
                rhs167_ = 49
                rhs168_ = (self).f18
                lhs142_ = d_1258_v63_
                lhs143_ = default__.safeIndex(265, (d_1258_v63_).length(0))
                lhs144_ = globalState
                lhs145_ = globalState
                lhs146_ = d_1263_v68_
                lhs147_ = default__.safeIndex(389, (d_1263_v68_).length(0))
                lhs142_[lhs143_] = rhs164_
                d_1262_v67_ = rhs165_
                lhs144_.f6 = rhs166_
                lhs145_.f6 = rhs167_
                lhs146_[lhs147_] = rhs168_
                d_1264_v69_: _dafny.MultiSet
                d_1264_v69_ = _dafny.MultiSet([(d_1263_v68_)[default__.safeIndex(389, (d_1263_v68_).length(0))], (d_1263_v68_)[default__.safeIndex(389, (d_1263_v68_).length(0))]])
                d_1265_v70_: _dafny.Seq
                d_1265_v70_ = _dafny.SeqWithoutIsStrInference([(d_1263_v68_)[default__.safeIndex(389, (d_1263_v68_).length(0))], (self).f18])
                (globalState).f1 = ((d_1264_v69_) | (d_1264_v69_)).issubset(_dafny.MultiSet(((d_1265_v70_).set(default__.safeIndex((self).f19, len(d_1265_v70_)), (self).f18)) + (d_1265_v70_)))
                d_1245_v50_ = (d_1245_v50_).set(592, p1)
                (globalState).f6 = len(p2)
                d_1266_v71_: D11
                d_1266_v71_ = D11_DC31(d_1265_v70_, (self).f19, p1, 300, (d_1258_v63_)[default__.safeIndex(265, (d_1258_v63_).length(0))])
                d_1267_v72_: _dafny.Seq
                d_1267_v72_ = _dafny.SeqWithoutIsStrInference([(self).f19, (d_1258_v63_)[default__.safeIndex(265, (d_1258_v63_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unna"))), default__.safeModuloInt(p1, (d_1266_v71_).cf58)])
                rhs169_ = _dafny.SeqWithoutIsStrInference([172, (self).f19, 334, len(p3), p0])
                rhs170_ = (0) - ((self).f19)
                lhs148_ = globalState
                d_1267_v72_ = rhs169_
                lhs148_.f6 = rhs170_
            d_1247_v52_ = (d_1247_v52_ if (self).f18 else d_1247_v52_)
            d_1268_v73_: C1
            nw241_ = C1()
            nw241_.ctor__(p0)
            d_1268_v73_ = nw241_
            d_1269_v74_: D0
            d_1269_v74_ = D0_DC1(886)
            pat_let_tv30_ = p1
            def iife159_(_pat_let55_0):
                def iife160_(d_1270_dt__update__tmp_h0_):
                    def iife161_(_pat_let56_0):
                        def iife162_(d_1271_dt__update_hcf1_h0_):
                            return D0_DC1(d_1271_dt__update_hcf1_h0_)
                        return iife162_(_pat_let56_0)
                    return iife161_(pat_let_tv30_)
                return iife160_(_pat_let55_0)
            d_1269_v74_ = iife159_(d_1269_v74_)
        r0 = (self).f18
        r1 = (p3) + (p2)
        return r0, r1

    def m13(self, globalState):
        d_1272_v0_: _dafny.Array
        nw242_ = _dafny.Array(_dafny.Seq({}), 9)
        d_1272_v0_ = nw242_
        d_1273_v1_: _dafny.Array
        nw243_ = _dafny.Array(_dafny.CodePoint('D'), 1)
        d_1273_v1_ = nw243_
        d_1274_v2_: _dafny.Seq
        d_1274_v2_ = _dafny.SeqWithoutIsStrInference([d_1273_v1_])
        index222_ = default__.safeIndex(246, (d_1272_v0_).length(0))
        (d_1272_v0_)[index222_] = (d_1274_v2_) + (_dafny.SeqWithoutIsStrInference([d_1273_v1_, d_1273_v1_, d_1273_v1_]))
        index223_ = default__.safeIndex(246, (d_1272_v0_).length(0))
        (d_1272_v0_)[index223_] = (d_1274_v2_).set(default__.safeIndex(default__.fm1((self).f18, globalState), len(d_1274_v2_)), d_1273_v1_)
        d_1275_v3_: _dafny.MultiSet
        d_1275_v3_ = _dafny.MultiSet([(self).f18])
        (globalState).f6 = default__.safeDivisionInt(default__.safeModuloInt((self).f19, (d_1275_v3_).cardinality), ((d_1275_v3_)[(self).f18] if ((self).f18) in (d_1275_v3_) else (self).f19))
        d_1276_v4_: _dafny.Seq
        d_1276_v4_ = _dafny.SeqWithoutIsStrInference([(self).f18])
        d_1277_v5_: _dafny.Array
        nw244_ = _dafny.Array(None, 10)
        nw244_[int(0)] = (self).f18
        nw244_[int(1)] = (self).f18
        nw244_[int(2)] = (self).f18
        nw244_[int(3)] = (self).f18
        nw244_[int(4)] = ((self).f18 if (self).f18 else (self).f18)
        nw244_[int(5)] = (self).f18
        nw244_[int(6)] = ((self).f19) < (731)
        nw244_[int(7)] = (d_1275_v3_).ispropersubset((_dafny.MultiSet(d_1276_v4_)).set((self).f18, default__.abs((self).f19)))
        nw244_[int(8)] = (self).f18
        nw244_[int(9)] = (self).f18
        d_1277_v5_ = nw244_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1277_v5_).length(0)):
            d_1278_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_1278_i0_)) and ((d_1278_i0_) < ((d_1277_v5_).length(0)))):
                (d_1277_v5_)[(d_1278_i0_)] = (self).f18
        index224_ = default__.safeIndex(148, (d_1277_v5_).length(0))
        (d_1277_v5_)[index224_] = (self).f18
        d_1279_v6_: C1
        nw245_ = C1()
        nw245_.ctor__((self).f19)
        d_1279_v6_ = nw245_
        d_1280_v7_: _dafny.Set
        d_1280_v7_ = _dafny.Set({d_1279_v6_, d_1279_v6_})
        d_1281_v8_: _dafny.Array
        nw246_ = _dafny.Array(int(0), 21)
        d_1281_v8_ = nw246_
        d_1282_v9_: D3
        d_1282_v9_ = D3_DC6(d_1281_v8_)
        d_1283_v10_: _dafny.Seq
        d_1283_v10_ = _dafny.SeqWithoutIsStrInference([d_1282_v9_])
        d_1284_v11_: _dafny.Map
        d_1284_v11_ = _dafny.Map({(self).f18: d_1283_v10_})
        d_1285_v12_: D11
        d_1285_v12_ = D11_DC30(((d_1284_v11_)[(self).f18] if ((self).f18) in (d_1284_v11_) else d_1283_v10_))
        d_1286_v13_: _dafny.Map
        d_1286_v13_ = _dafny.Map({d_1280_v7_: d_1285_v12_})
        d_1287_v14_: _dafny.Array
        nw247_ = _dafny.Array(None, 3)
        nw247_[int(0)] = d_1281_v8_
        nw247_[int(1)] = d_1281_v8_
        nw247_[int(2)] = d_1281_v8_
        d_1287_v14_ = nw247_
        index225_ = default__.safeIndex(318, (d_1287_v14_).length(0))
        (d_1287_v14_)[index225_] = d_1281_v8_
        d_1288_v15_: _dafny.Map
        d_1288_v15_ = _dafny.Map({(d_1279_v6_).f24: not((self).f18)})
        d_1289_v16_: _dafny.Seq
        d_1289_v16_ = _dafny.SeqWithoutIsStrInference([len(d_1288_v15_)])
        d_1290_v17_: _dafny.Map
        d_1290_v17_ = _dafny.Map({not(((d_1279_v6_).f24) not in (d_1289_v16_)): 891})
        index226_ = default__.safeIndex(148, (d_1277_v5_).length(0))
        index227_ = default__.safeIndex(318, (d_1287_v14_).length(0))
        rhs171_ = ((self).f18) == (((self).f18 if False else (self).f18))
        rhs172_ = d_1286_v13_
        rhs173_ = ((d_1290_v17_)[True] if (True) in (d_1290_v17_) else (0) - ((self).f19))
        rhs174_ = d_1281_v8_
        lhs149_ = d_1277_v5_
        lhs150_ = default__.safeIndex(148, (d_1277_v5_).length(0))
        lhs151_ = globalState
        lhs152_ = d_1287_v14_
        lhs153_ = default__.safeIndex(318, (d_1287_v14_).length(0))
        lhs149_[lhs150_] = rhs171_
        d_1286_v13_ = rhs172_
        lhs151_.f6 = rhs173_
        lhs152_[lhs153_] = rhs174_
        (globalState).f1 = (True if (d_1277_v5_)[default__.safeIndex(148, (d_1277_v5_).length(0))] else (d_1277_v5_)[default__.safeIndex(148, (d_1277_v5_).length(0))])
        d_1291_v18_: _dafny.Set
        d_1291_v18_ = _dafny.Set({(self).f18})
        d_1292_i1_: int
        d_1292_i1_ = 0
        with _dafny.label("9"):
            while not (not((_dafny.Set({(d_1277_v5_)[default__.safeIndex(148, (d_1277_v5_).length(0))], (d_1277_v5_)[default__.safeIndex(148, (d_1277_v5_).length(0))], (d_1277_v5_)[default__.safeIndex(148, (d_1277_v5_).length(0))], (d_1277_v5_)[default__.safeIndex(148, (d_1277_v5_).length(0))], (self).f18})).isdisjoint(d_1291_v18_))) or ((self).f18):
                with _dafny.c_label("9"):
                    if (d_1292_i1_) >= (100):
                        raise _dafny.Break("9")
                    d_1292_i1_ = (d_1292_i1_) + (1)
                    d_1293_v21_: _dafny.Set
                    def iife163_():
                        def iife165_():
                            coll51_ = _dafny.Set()
                            compr_51_: int
                            for compr_51_ in (d_1289_v16_).Elements:
                                d_1296_v20_: int = compr_51_
                                if (d_1296_v20_) in (d_1289_v16_):
                                    coll51_ = coll51_.union(_dafny.Set([(d_1296_v20_) - (466)]))
                            return _dafny.Set(coll51_)
                        coll49_ = _dafny.Map()
                        def iife164_():
                            coll50_ = _dafny.Set()
                            compr_50_: int
                            for compr_50_ in (d_1289_v16_).Elements:
                                d_1294_v20_: int = compr_50_
                                if (d_1294_v20_) in (d_1289_v16_):
                                    coll50_ = coll50_.union(_dafny.Set([(d_1294_v20_) - (466)]))
                            return _dafny.Set(coll50_)
                        compr_49_: int
                        for compr_49_ in (iife164_()
                        ).Elements:
                            d_1295_v19_: int = compr_49_
                            if (d_1295_v19_) in (iife165_()
                            ):
                                coll49_[default__.safeDivisionInt(d_1295_v19_, 240)] = (d_1279_v6_).f24
                        return _dafny.Map(coll49_)
                    d_1293_v21_ = _dafny.Set({len(iife163_()
                    )})
                    d_1297_v22_: _dafny.Seq
                    d_1297_v22_ = _dafny.SeqWithoutIsStrInference([d_1293_v21_, d_1293_v21_])
                    if not((d_1293_v21_).ispropersubset((d_1297_v22_)[default__.safeIndex((d_1279_v6_).f24, len(d_1297_v22_))])):
                        index228_ = default__.safeIndex(148, (d_1277_v5_).length(0))
                        (d_1277_v5_)[index228_] = True
                        d_1291_v18_ = _dafny.Set({(d_1291_v18_).issubset(_dafny.Set({not((d_1277_v5_)[default__.safeIndex(148, (d_1277_v5_).length(0))])}))})
                        d_1288_v15_ = (d_1288_v15_).set((d_1279_v6_).f24, False)
                        d_1298_v24_: _dafny.Map
                        d_1298_v24_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18]): 72})
                        def iife166_():
                            coll52_ = _dafny.Map()
                            compr_52_: _dafny.Seq
                            for compr_52_ in (_dafny.Map({d_1276_v4_: (0) - ((d_1279_v6_).f24)})).keys.Elements:
                                d_1299_v23_: _dafny.Seq = compr_52_
                                if (d_1299_v23_) in (_dafny.Map({d_1276_v4_: (0) - ((d_1279_v6_).f24)})):
                                    coll52_[d_1299_v23_] = (d_1279_v6_).f24
                            return _dafny.Map(coll52_)
                        (globalState).f1 = ((iife166_()
                        ).set(d_1276_v4_, (d_1279_v6_).f24)) != (d_1298_v24_)
                        d_1300_v25_: _dafny.Map
                        d_1300_v25_ = _dafny.Map({(self).f18: (d_1277_v5_)[default__.safeIndex(148, (d_1277_v5_).length(0))]})
                        d_1301_v26_: _dafny.MultiSet
                        d_1301_v26_ = _dafny.MultiSet([d_1291_v18_])
                        d_1300_v25_ = (d_1300_v25_).set((self).f18, (d_1301_v26_).ispropersubset(d_1301_v26_))
                    elif True:
                        d_1302_v27_: _dafny.Seq
                        d_1302_v27_ = _dafny.SeqWithoutIsStrInference([d_1277_v5_])
                        d_1303_v28_: _dafny.Array
                        nw248_ = _dafny.Array(None, 13)
                        nw248_[int(0)] = d_1277_v5_
                        nw248_[int(1)] = d_1277_v5_
                        nw248_[int(2)] = d_1277_v5_
                        nw248_[int(3)] = (d_1277_v5_ if (self).f18 else d_1277_v5_)
                        nw248_[int(4)] = d_1277_v5_
                        nw248_[int(5)] = d_1277_v5_
                        nw248_[int(6)] = d_1277_v5_
                        nw248_[int(7)] = (d_1277_v5_ if False else d_1277_v5_)
                        nw248_[int(8)] = d_1277_v5_
                        nw248_[int(9)] = (d_1302_v27_)[default__.safeIndex((self).f19, len(d_1302_v27_))]
                        nw248_[int(10)] = d_1277_v5_
                        nw248_[int(11)] = d_1277_v5_
                        nw248_[int(12)] = d_1277_v5_
                        d_1303_v28_ = nw248_
                        d_1303_v28_ = d_1303_v28_
                        d_1304_v29_: _dafny.Map
                        d_1304_v29_ = _dafny.Map({not(((d_1279_v6_).f24) > ((0) - ((self).f19))): True})
                        d_1304_v29_ = (d_1304_v29_).set((d_1277_v5_)[default__.safeIndex(148, (d_1277_v5_).length(0))], (self).f18)
                        (globalState).f14 = d_1277_v5_
                        index229_ = default__.safeIndex(71, (d_1281_v8_).length(0))
                        (d_1281_v8_)[index229_] = default__.fm1(default__.fm2(globalState), globalState)
                        d_1305_v30_: _dafny.MultiSet
                        d_1305_v30_ = _dafny.MultiSet([((self).f19) * (153), ((d_1279_v6_).f24) * ((d_1279_v6_).f24)])
                        d_1306_v31_: _dafny.Map
                        d_1306_v31_ = _dafny.Map({((d_1275_v3_).intersection(d_1275_v3_)).cardinality: default__.fm1((self).f18, globalState)})
                        d_1307_v32_: str
                        d_1307_v32_ = _dafny.CodePoint('y')
                        d_1308_v33_: _dafny.Map
                        d_1308_v33_ = _dafny.Map({310: _dafny.Map({(d_1279_v6_).f24: (self).f19})})
                        d_1309_v34_: _dafny.Seq
                        d_1309_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rboenvi"))
                        index230_ = default__.safeIndex(71, (d_1281_v8_).length(0))
                        rhs175_ = (0) - (len(_dafny.SeqWithoutIsStrInference([d_1307_v32_, d_1307_v32_, d_1307_v32_])))
                        rhs176_ = d_1305_v30_
                        rhs177_ = ((((d_1308_v33_)[(d_1279_v6_).f24] if ((d_1279_v6_).f24) in (d_1308_v33_) else d_1306_v31_) if False else default__.fm33((d_1279_v6_).f24, (d_1279_v6_).f24, (self).f18, globalState))) | (_dafny.Map({(d_1279_v6_).f24: len(d_1309_v34_)}))
                        rhs178_ = d_1277_v5_
                        lhs154_ = d_1281_v8_
                        lhs155_ = default__.safeIndex(71, (d_1281_v8_).length(0))
                        lhs156_ = globalState
                        lhs154_[lhs155_] = rhs175_
                        d_1305_v30_ = rhs176_
                        d_1306_v31_ = rhs177_
                        lhs156_.f14 = rhs178_
                        d_1304_v29_ = (d_1304_v29_).set((self).f18, (self).f18)
                    d_1310_v36_: _dafny.Map
                    def iife167_():
                        coll53_ = _dafny.Map()
                        compr_53_: int
                        for compr_53_ in _dafny.IntegerRange(305, 589):
                            d_1311_v35_: int = compr_53_
                            if ((305) <= (d_1311_v35_)) and ((d_1311_v35_) < (589)):
                                coll53_[(d_1311_v35_) * ((d_1279_v6_).f24)] = False
                        return _dafny.Map(coll53_)
                    d_1310_v36_ = _dafny.Map({iife167_()
                    : True})
                    index231_ = default__.safeIndex(148, (d_1277_v5_).length(0))
                    (d_1277_v5_)[index231_] = ((d_1277_v5_)[default__.safeIndex(148, (d_1277_v5_).length(0))] if (d_1279_v6_).fm24((d_1277_v5_)[default__.safeIndex(148, (d_1277_v5_).length(0))], d_1310_v36_, globalState) else (self).f18)
                    d_1312_v37_: _dafny.Map
                    d_1312_v37_ = _dafny.Map({(d_1279_v6_).f24: (d_1279_v6_).f24})
                    rhs179_ = (d_1312_v37_) | (d_1312_v37_)
                    rhs180_ = (self).f18
                    rhs181_ = (d_1279_v6_).f24
                    lhs157_ = globalState
                    lhs158_ = globalState
                    d_1312_v37_ = rhs179_
                    lhs157_.f1 = rhs180_
                    lhs158_.f6 = rhs181_
                    source9_ = D13_DC37(True, True)
                    if source9_.is_DC37:
                        d_1313___mcc_h0_ = source9_.cf65
                        d_1314___mcc_h1_ = source9_.cf66
                        d_1315_cf66_ = d_1314___mcc_h1_
                        d_1316_cf65_ = d_1313___mcc_h0_
                        d_1317_v38_: str
                        d_1317_v38_ = _dafny.CodePoint('y')
                        d_1318_v39_: _dafny.Seq
                        d_1318_v39_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f'), d_1317_v38_])
                        d_1319_v40_: _dafny.Map
                        d_1319_v40_ = _dafny.Map({(d_1279_v6_).f24: _dafny.CodePoint('m')})
                        d_1320_v41_: _dafny.Map
                        d_1320_v41_ = _dafny.Map({len(d_1318_v39_): d_1318_v39_})
                        nw249_ = _dafny.Array(None, 11)
                        nw249_[int(0)] = (self).f19
                        nw249_[int(1)] = default__.safeModuloInt((d_1279_v6_).f24, len(d_1318_v39_))
                        nw249_[int(2)] = len(d_1319_v40_)
                        nw249_[int(3)] = (d_1279_v6_).f24
                        nw249_[int(4)] = (d_1279_v6_).f24
                        nw249_[int(5)] = (d_1279_v6_).f24
                        nw249_[int(6)] = len((_dafny.SeqWithoutIsStrInference([(self).f19 for d_1321_i2_ in range(default__.abs(765))])) + (d_1289_v16_))
                        nw249_[int(7)] = (d_1279_v6_).f24
                        nw249_[int(8)] = -42
                        nw249_[int(9)] = (d_1279_v6_).f24
                        nw249_[int(10)] = len((d_1320_v41_) | (default__.fm38((self).f19, (d_1277_v5_)[default__.safeIndex(148, (d_1277_v5_).length(0))], True, (self).f18, globalState)))
                        (globalState).f9 = nw249_
                        (globalState).f1 = not((self).f18)
                        (globalState).f6 = (0) - ((d_1279_v6_).f24)
                        d_1322_v42_: D5
                        d_1322_v42_ = D5_DC12((self).f19, (self).f18, d_1317_v38_)
                        d_1323_v43_: D5
                        d_1323_v43_ = D5_DC13(d_1322_v42_)
                        d_1323_v43_ = d_1323_v43_
                    elif True:
                        d_1324___mcc_h2_ = source9_.cf64
                        d_1325_cf64_ = d_1324___mcc_h2_
                        d_1326_v44_: _dafny.Seq
                        d_1326_v44_ = _dafny.SeqWithoutIsStrInference([d_1283_v10_, (d_1283_v10_).set(default__.safeIndex((self).f19, len(d_1283_v10_)), d_1282_v9_)])
                        d_1327_v45_: C0
                        nw250_ = C0()
                        nw250_.ctor__(((d_1326_v44_)[default__.safeIndex((d_1279_v6_).f24, len(d_1326_v44_))]).set(default__.safeIndex((d_1279_v6_).f24, len((d_1326_v44_)[default__.safeIndex((d_1279_v6_).f24, len(d_1326_v44_))])), d_1282_v9_), False)
                        d_1327_v45_ = nw250_
                        d_1328_v46_: _dafny.Map
                        d_1328_v46_ = _dafny.Map({(d_1327_v45_).f23: (d_1327_v45_).f23})
                        d_1290_v17_ = (d_1290_v17_).set((d_1277_v5_)[default__.safeIndex(148, (d_1277_v5_).length(0))], (0) - (len(d_1328_v46_)))
                        d_1329_v47_: C2
                        nw251_ = C2()
                        nw251_.ctor__((d_1279_v6_).f24, (266 if (d_1327_v45_).f23 else (d_1279_v6_).f24), self.f17)
                        d_1329_v47_ = nw251_
                        d_1328_v46_ = d_1328_v46_
                    pass
            pass

    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19

class C4(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self):
        pass
        pass

    def fm3(self, p0, p1, p2, p3, globalState):
        return D0_DC2(D0_DC1(len(_dafny.SeqWithoutIsStrInference([946, -169, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(True), False, False])): True})), 586, (_dafny.MultiSet([458])).cardinality]))))

    def m1(self, globalState):
        r0: int = int(0)
        r1: bool = False
        d_1330_v0_: _dafny.Seq
        d_1330_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jyvixh"))
        d_1330_v0_ = d_1330_v0_
        d_1331_v1_: _dafny.Array
        def lambda55_(d_1332_i0_):
            return (d_1332_i0_) + (len(_dafny.Map({_dafny.Map({not(True): True}): 783})))

        init35_ = lambda55_
        nw252_ = _dafny.Array(None, 13)
        for i0_35_ in range(nw252_.length(0)):
            nw252_[i0_35_] = init35_(i0_35_)
        d_1331_v1_ = nw252_
        d_1333_v2_: int
        d_1333_v2_ = 814
        index232_ = default__.safeIndex(822, (d_1331_v1_).length(0))
        (d_1331_v1_)[index232_] = (0) - (d_1333_v2_)
        d_1334_v3_: _dafny.Array
        nw253_ = _dafny.Array(False, 2)
        d_1334_v3_ = nw253_
        d_1335_v4_: str
        d_1335_v4_ = _dafny.CodePoint('l')
        d_1336_v5_: _dafny.MultiSet
        d_1336_v5_ = _dafny.MultiSet([d_1335_v4_])
        index233_ = default__.safeIndex(511, (d_1334_v3_).length(0))
        (d_1334_v3_)[index233_] = (d_1335_v4_) not in (d_1336_v5_)
        d_1337_v6_: bool
        d_1337_v6_ = True
        index234_ = default__.safeIndex(822, (d_1331_v1_).length(0))
        index235_ = default__.safeIndex(511, (d_1334_v3_).length(0))
        rhs182_ = default__.safeModuloInt(337, (d_1333_v2_) - (d_1333_v2_))
        rhs183_ = default__.fm1(d_1337_v6_, globalState)
        rhs184_ = d_1337_v6_
        rhs185_ = default__.safeDivisionInt(d_1333_v2_, (d_1333_v2_) + (-621))
        rhs186_ = (d_1330_v0_ if d_1337_v6_ else d_1330_v0_)
        lhs159_ = globalState
        lhs160_ = d_1331_v1_
        lhs161_ = default__.safeIndex(822, (d_1331_v1_).length(0))
        lhs162_ = d_1334_v3_
        lhs163_ = default__.safeIndex(511, (d_1334_v3_).length(0))
        lhs164_ = globalState
        lhs159_.f6 = rhs182_
        lhs160_[lhs161_] = rhs183_
        lhs162_[lhs163_] = rhs184_
        lhs164_.f6 = rhs185_
        d_1330_v0_ = rhs186_
        d_1338_v7_: _dafny.Map
        d_1338_v7_ = _dafny.Map({569: (d_1334_v3_)[default__.safeIndex(511, (d_1334_v3_).length(0))]})
        d_1339_v8_: _dafny.Map
        d_1339_v8_ = _dafny.Map({d_1338_v7_: d_1333_v2_})
        d_1340_v9_: C3
        nw254_ = C3()
        nw254_.ctor__(False, d_1333_v2_, d_1339_v8_)
        d_1340_v9_ = nw254_
        index236_ = default__.safeIndex(822, (d_1331_v1_).length(0))
        (d_1331_v1_)[index236_] = (d_1340_v9_).f19
        d_1341_v10_: _dafny.Map
        d_1341_v10_ = _dafny.Map({len(d_1338_v7_): d_1333_v2_})
        d_1342_v11_: D12
        d_1342_v11_ = D12_DC34(d_1341_v10_)
        pat_let_tv31_ = d_1340_v9_
        pat_let_tv32_ = d_1340_v9_
        pat_let_tv33_ = d_1331_v1_
        pat_let_tv34_ = d_1331_v1_
        index237_ = default__.safeIndex(822, (d_1331_v1_).length(0))
        def lambda56_(source10_):
            if source10_.is_DC35:
                return default__.safeModuloInt((pat_let_tv31_).f19, (pat_let_tv32_).f19)
            elif True:
                d_1343___mcc_h0_ = source10_.cf63
                d_1344_cf63_ = d_1343___mcc_h0_
                return (pat_let_tv34_)[default__.safeIndex(822, (pat_let_tv33_).length(0))]

        (d_1331_v1_)[index237_] = lambda56_(d_1342_v11_)
        d_1337_v6_ = not(False)
        r0 = (d_1331_v1_)[default__.safeIndex(822, (d_1331_v1_).length(0))]
        d_1345_v12_: _dafny.Seq
        d_1345_v12_ = _dafny.SeqWithoutIsStrInference([(d_1334_v3_)[default__.safeIndex(511, (d_1334_v3_).length(0))], default__.fm2(globalState), (d_1334_v3_)[default__.safeIndex(511, (d_1334_v3_).length(0))], d_1337_v6_, (d_1340_v9_).f18])
        r1 = not((False if d_1337_v6_ else not((d_1333_v2_) == (len(d_1345_v12_)))))
        return r0, r1

    def m2(self, p0, globalState):
        d_1346_v0_: _dafny.Array
        nw255_ = _dafny.Array(_dafny.CodePoint('D'), 27)
        d_1346_v0_ = nw255_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1346_v0_).length(0)):
            d_1347_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_1347_i0_)) and ((d_1347_i0_) < ((d_1346_v0_).length(0)))):
                (d_1346_v0_)[(d_1347_i0_)] = (_dafny.CodePoint('u') if False else (_dafny.CodePoint('n') if False else (D5_DC12((_dafny.MultiSet([968])).cardinality, p0, _dafny.CodePoint('r'))).cf17))
        d_1348_v1_: int
        d_1348_v1_ = -497
        d_1349_v2_: _dafny.Map
        d_1349_v2_ = _dafny.Map({_dafny.Set({p0, p0}): d_1348_v1_})
        d_1349_v2_ = d_1349_v2_
        d_1350_v3_: _dafny.Array
        def lambda57_(d_1351_p0_):
            def lambda58_(d_1352_i1_):
                return _dafny.MultiSet([d_1351_p0_, d_1351_p0_, d_1351_p0_, True])

            return lambda58_

        init36_ = lambda57_(p0)
        nw256_ = _dafny.Array(None, 17)
        for i0_36_ in range(nw256_.length(0)):
            nw256_[i0_36_] = init36_(i0_36_)
        d_1350_v3_ = nw256_
        d_1350_v3_ = d_1350_v3_
        if p0:
            (globalState).f6 = default__.safeDivisionInt(d_1348_v1_, d_1348_v1_)
            d_1353_v4_: _dafny.Array
            nw257_ = _dafny.Array(int(0), 25)
            d_1353_v4_ = nw257_
            index238_ = default__.safeIndex(526, (d_1353_v4_).length(0))
            (d_1353_v4_)[index238_] = -199
            index239_ = default__.safeIndex(526, (d_1353_v4_).length(0))
            (d_1353_v4_)[index239_] = d_1348_v1_
            index240_ = default__.safeIndex(526, (d_1353_v4_).length(0))
            (d_1353_v4_)[index240_] = (d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))]
            d_1354_v5_: str
            d_1354_v5_ = _dafny.CodePoint('n')
            d_1355_v6_: _dafny.MultiSet
            d_1355_v6_ = _dafny.MultiSet([d_1354_v5_, d_1354_v5_, d_1354_v5_, _dafny.CodePoint('d')])
            d_1356_v7_: _dafny.Map
            d_1356_v7_ = _dafny.Map({p0: (d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))]})
            d_1357_v9_: _dafny.MultiSet
            d_1357_v9_ = _dafny.MultiSet([d_1348_v1_, d_1348_v1_])
            d_1358_v10_: D8
            d_1358_v10_ = D8_DC20(65, d_1348_v1_, (d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))])
            d_1359_v11_: _dafny.Map
            d_1359_v11_ = _dafny.Map({(d_1358_v10_).cf31: (d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))]})
            d_1360_v12_: _dafny.Seq
            d_1360_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
            d_1361_v13_: D5
            d_1361_v13_ = D5_DC12((d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))], p0, d_1354_v5_)
            d_1362_v14_: _dafny.Set
            d_1362_v14_ = _dafny.Set({d_1361_v13_})
            d_1363_v15_: _dafny.Seq
            d_1363_v15_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1364_v16_: _dafny.Map
            d_1364_v16_ = _dafny.Map({p0: d_1363_v15_})
            d_1365_v17_: _dafny.MultiSet
            d_1365_v17_ = _dafny.MultiSet([p0])
            d_1366_v18_: _dafny.Map
            d_1366_v18_ = _dafny.Map({p0: p0})
            nw258_ = _dafny.Array(None, 26)
            nw258_[int(0)] = ((d_1355_v6_)[d_1354_v5_] if (d_1354_v5_) in (d_1355_v6_) else default__.fm1(p0, globalState))
            nw258_[int(1)] = d_1348_v1_
            nw258_[int(2)] = ((d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))]) * (d_1348_v1_)
            nw258_[int(3)] = len(d_1356_v7_)
            nw258_[int(4)] = (d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))]
            def iife168_():
                coll54_ = _dafny.Map()
                compr_54_: int
                for compr_54_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({p0})) for d_1367_i2_ in range(default__.abs(936))])).Elements:
                    d_1368_v8_: int = compr_54_
                    if (d_1368_v8_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({p0})) for d_1367_i2_ in range(default__.abs(936))])):
                        coll54_[default__.safeDivisionInt(d_1368_v8_, (d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))])] = (0) - (d_1348_v1_)
                return _dafny.Map(coll54_)
            nw258_[int(5)] = len(((iife168_()
            ).set((d_1357_v9_).cardinality, d_1348_v1_)) | (d_1359_v11_))
            nw258_[int(6)] = (d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))]
            nw258_[int(7)] = default__.safeDivisionInt(d_1348_v1_, (d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))])
            nw258_[int(8)] = ((d_1359_v11_)[default__.fm1(p0, globalState)] if (default__.fm1(p0, globalState)) in (d_1359_v11_) else len(d_1360_v12_))
            nw258_[int(9)] = len((d_1356_v7_) | (d_1356_v7_))
            nw258_[int(10)] = (d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))]
            nw258_[int(11)] = len((d_1362_v14_).intersection(d_1362_v14_))
            nw258_[int(12)] = (d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))]
            nw258_[int(13)] = d_1348_v1_
            nw258_[int(14)] = d_1348_v1_
            nw258_[int(15)] = (d_1348_v1_) * ((d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))])
            nw258_[int(16)] = (d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))]
            nw258_[int(17)] = default__.safeModuloInt(d_1348_v1_, (0) - (d_1348_v1_))
            nw258_[int(18)] = ((d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))]) + (761)
            nw258_[int(19)] = default__.fm1(p0, globalState)
            nw258_[int(20)] = (0) - (len(d_1364_v16_))
            nw258_[int(21)] = ((d_1365_v17_).cardinality) * ((d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))])
            nw258_[int(22)] = (d_1353_v4_)[default__.safeIndex(526, (d_1353_v4_).length(0))]
            nw258_[int(23)] = len(d_1359_v11_)
            nw258_[int(24)] = d_1348_v1_
            nw258_[int(25)] = len(d_1366_v18_)
            d_1353_v4_ = nw258_
            (globalState).f6 = (d_1348_v1_ if not(p0) else d_1348_v1_)
        elif True:
            d_1348_v1_ = d_1348_v1_
            d_1369_v20_: _dafny.Map
            d_1369_v20_ = _dafny.Map({d_1348_v1_: p0})
            d_1370_v21_: _dafny.MultiSet
            def iife169_():
                coll55_ = _dafny.Map()
                compr_55_: int
                for compr_55_ in (d_1369_v20_).keys.Elements:
                    d_1371_v19_: int = compr_55_
                    if (d_1371_v19_) in (d_1369_v20_):
                        coll55_[(d_1371_v19_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1372_i3_ in range(default__.abs(485))])))] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpfcroym"))
                return _dafny.Map(coll55_)
            d_1370_v21_ = _dafny.MultiSet([len(iife169_()
            ), d_1348_v1_])
            d_1373_v22_: _dafny.Seq
            d_1373_v22_ = _dafny.SeqWithoutIsStrInference([d_1370_v21_])
            d_1374_v23_: _dafny.Map
            d_1374_v23_ = _dafny.Map({p0: d_1373_v22_})
            d_1375_v24_: _dafny.Map
            d_1375_v24_ = _dafny.Map({p0: p0})
            d_1374_v23_ = (d_1374_v23_).set(((d_1375_v24_)[p0] if (p0) in (d_1375_v24_) else p0), d_1373_v22_)
            d_1376_v25_: _dafny.MultiSet
            d_1376_v25_ = _dafny.MultiSet([p0])
            (globalState).f6 = (((d_1376_v25_)[p0] if (p0) in (d_1376_v25_) else d_1348_v1_)) + ((0) - (d_1348_v1_))
            d_1377_v26_: _dafny.Seq
            d_1377_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            d_1378_v27_: str
            d_1378_v27_ = _dafny.CodePoint('a')
            d_1379_v28_: _dafny.Seq
            d_1379_v28_ = _dafny.SeqWithoutIsStrInference([d_1377_v26_])
            d_1380_v29_: _dafny.Array
            nw259_ = _dafny.Array(None, 20)
            nw259_[int(0)] = d_1377_v26_
            nw259_[int(1)] = d_1377_v26_
            nw259_[int(2)] = (default__.fm21(globalState)).set(default__.safeIndex(d_1348_v1_, len(default__.fm21(globalState))), d_1378_v27_)
            nw259_[int(3)] = d_1377_v26_
            nw259_[int(4)] = _dafny.SeqWithoutIsStrInference([d_1378_v27_ for d_1381_i4_ in range(default__.abs(258))])
            nw259_[int(5)] = d_1377_v26_
            nw259_[int(6)] = d_1377_v26_
            nw259_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            nw259_[int(8)] = d_1377_v26_
            nw259_[int(9)] = d_1377_v26_
            nw259_[int(10)] = d_1377_v26_
            nw259_[int(11)] = d_1377_v26_
            nw259_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtuxv"))
            nw259_[int(13)] = d_1377_v26_
            nw259_[int(14)] = _dafny.SeqWithoutIsStrInference([d_1378_v27_ for d_1382_i5_ in range(default__.abs(406))])
            nw259_[int(15)] = d_1377_v26_
            nw259_[int(16)] = (d_1379_v28_)[default__.safeIndex(d_1348_v1_, len(d_1379_v28_))]
            nw259_[int(17)] = d_1377_v26_
            nw259_[int(18)] = d_1377_v26_
            nw259_[int(19)] = d_1377_v26_
            d_1380_v29_ = nw259_
            d_1383_v30_: _dafny.Array
            nw260_ = _dafny.Array(None, 4)
            nw260_[int(0)] = d_1380_v29_
            nw260_[int(1)] = d_1380_v29_
            nw260_[int(2)] = d_1380_v29_
            nw260_[int(3)] = d_1380_v29_
            d_1383_v30_ = nw260_
            index241_ = default__.safeIndex(282, (d_1383_v30_).length(0))
            (d_1383_v30_)[index241_] = d_1380_v29_
            index242_ = default__.safeIndex(282, (d_1383_v30_).length(0))
            def lambda59_(d_1384_v26_):
                def lambda60_(d_1385_i6_):
                    return d_1384_v26_

                return lambda60_

            init37_ = lambda59_(d_1377_v26_)
            nw261_ = _dafny.Array(None, 9)
            for i0_37_ in range(nw261_.length(0)):
                nw261_[i0_37_] = init37_(i0_37_)
            (d_1383_v30_)[index242_] = nw261_
            (globalState).f1 = p0
        d_1386_i7_: int
        d_1386_i7_ = 0
        with _dafny.label("10"):
            while p0:
                with _dafny.c_label("10"):
                    if (d_1386_i7_) >= (100):
                        raise _dafny.Break("10")
                    d_1386_i7_ = (d_1386_i7_) + (1)
                    d_1387_v31_: _dafny.Set
                    d_1387_v31_ = _dafny.Set({d_1348_v1_, d_1348_v1_, d_1348_v1_})
                    d_1388_v32_: C1
                    nw262_ = C1()
                    nw262_.ctor__(len(d_1387_v31_))
                    d_1388_v32_ = nw262_
                    d_1389_v33_: _dafny.Array
                    nw263_ = _dafny.Array(int(0), 16)
                    d_1389_v33_ = nw263_
                    d_1390_v34_: D3
                    d_1390_v34_ = D3_DC6(d_1389_v33_)
                    d_1391_v35_: _dafny.Seq
                    d_1391_v35_ = _dafny.SeqWithoutIsStrInference([d_1390_v34_, d_1390_v34_, d_1390_v34_, d_1390_v34_, d_1390_v34_])
                    d_1392_v36_: C0
                    nw264_ = C0()
                    nw264_.ctor__(d_1391_v35_, p0)
                    d_1392_v36_ = nw264_
                    if True:
                        d_1393_v37_: str
                        d_1393_v37_ = _dafny.CodePoint('i')
                        d_1393_v37_ = d_1393_v37_
                        d_1394_v38_: _dafny.Seq
                        d_1394_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgkj"))
                        d_1395_v39_: _dafny.MultiSet
                        d_1395_v39_ = _dafny.MultiSet([d_1387_v31_])
                        d_1396_v40_: _dafny.Map
                        d_1396_v40_ = _dafny.Map({len(d_1394_v38_): (d_1395_v39_).cardinality})
                        d_1397_v41_: D12
                        d_1397_v41_ = D12_DC34(d_1396_v40_)
                        d_1398_v43_: _dafny.Map
                        def iife170_():
                            coll56_ = _dafny.Set()
                            compr_56_: int
                            for compr_56_ in _dafny.IntegerRange(470, 885):
                                d_1399_v42_: int = compr_56_
                                if ((470) <= (d_1399_v42_)) and ((d_1399_v42_) < (885)):
                                    coll56_ = coll56_.union(_dafny.Set([(d_1399_v42_) * ((d_1388_v32_).f24)]))
                            return _dafny.Set(coll56_)
                        d_1398_v43_ = _dafny.Map({d_1397_v41_: iife170_()
                        })
                        d_1400_v44_: _dafny.Array
                        nw265_ = _dafny.Array(None, 4)
                        nw265_[int(0)] = d_1398_v43_
                        nw265_[int(1)] = d_1398_v43_
                        nw265_[int(2)] = d_1398_v43_
                        nw265_[int(3)] = d_1398_v43_
                        d_1400_v44_ = nw265_
                        index243_ = default__.safeIndex(220, (d_1400_v44_).length(0))
                        (d_1400_v44_)[index243_] = d_1398_v43_
                        d_1401_v45_: _dafny.MultiSet
                        d_1401_v45_ = _dafny.MultiSet([(d_1388_v32_).f24, d_1348_v1_])
                        d_1402_v46_: _dafny.Seq
                        d_1402_v46_ = _dafny.SeqWithoutIsStrInference([d_1401_v45_])
                        d_1403_v47_: _dafny.Seq
                        d_1403_v47_ = _dafny.SeqWithoutIsStrInference([len(d_1402_v46_)])
                        d_1404_v50_: _dafny.Set
                        d_1404_v50_ = _dafny.Set({d_1397_v41_})
                        index244_ = default__.safeIndex(220, (d_1400_v44_).length(0))
                        def iife171_():
                            def iife173_():
                                coll59_ = _dafny.Map()
                                compr_59_: D12
                                for compr_59_ in (d_1404_v50_).Elements:
                                    d_1405_v49_: D12 = compr_59_
                                    if (d_1405_v49_) in (d_1404_v50_):
                                        coll59_[d_1405_v49_] = d_1393_v37_
                                return _dafny.Map(coll59_)
                            coll57_ = _dafny.Map()
                            def iife172_():
                                coll58_ = _dafny.Map()
                                compr_58_: D12
                                for compr_58_ in (d_1404_v50_).Elements:
                                    d_1405_v49_: D12 = compr_58_
                                    if (d_1405_v49_) in (d_1404_v50_):
                                        coll58_[d_1405_v49_] = d_1393_v37_
                                return _dafny.Map(coll58_)
                            compr_57_: D12
                            for compr_57_ in (iife172_()
                            ).keys.Elements:
                                d_1406_v48_: D12 = compr_57_
                                if (d_1406_v48_) in (iife173_()
                                ):
                                    coll57_[d_1406_v48_] = d_1387_v31_
                            return _dafny.Map(coll57_)
                        rhs187_ = ((iife171_()
                        ) | (_dafny.Map({d_1397_v41_: d_1387_v31_}))) | (d_1398_v43_)
                        rhs188_ = d_1403_v47_
                        lhs165_ = d_1400_v44_
                        lhs166_ = default__.safeIndex(220, (d_1400_v44_).length(0))
                        lhs165_[lhs166_] = rhs187_
                        d_1403_v47_ = rhs188_
                        (globalState).f1 = True
                        d_1348_v1_ = ((d_1388_v32_).f24) - ((d_1388_v32_).f24)
                        (globalState).f1 = default__.fm2(globalState)
                    elif True:
                        d_1407_v51_: str
                        d_1407_v51_ = _dafny.CodePoint('k')
                        d_1408_v52_: _dafny.Seq
                        d_1408_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okij"))
                        d_1409_v53_: _dafny.Map
                        d_1409_v53_ = _dafny.Map({d_1407_v51_: d_1408_v52_})
                        d_1410_v54_: _dafny.Map
                        d_1410_v54_ = _dafny.Map({d_1407_v51_: (d_1388_v32_).f24})
                        index245_ = default__.safeIndex(301, (d_1389_v33_).length(0))
                        (d_1389_v33_)[index245_] = default__.safeDivisionInt(len(d_1409_v53_), len(default__.fm28(len(d_1410_v54_), (d_1392_v36_).f23, True, globalState)))
                        index246_ = default__.safeIndex(301, (d_1389_v33_).length(0))
                        (d_1389_v33_)[index246_] = len(d_1408_v52_)
                        (globalState).f1 = (d_1392_v36_).f23
                        rhs189_ = (d_1392_v36_).f23
                        rhs190_ = d_1408_v52_
                        rhs191_ = p0
                        lhs167_ = globalState
                        lhs168_ = globalState
                        lhs167_.f1 = rhs189_
                        d_1408_v52_ = rhs190_
                        lhs168_.f1 = rhs191_
                        d_1411_v55_: _dafny.Array
                        def lambda61_(d_1412_v51_):
                            def lambda62_(d_1413_i8_):
                                return D1_DC4((D1_DC4(d_1412_v51_)).cf4)

                            return lambda62_

                        init38_ = lambda61_(d_1407_v51_)
                        nw266_ = _dafny.Array(None, 12)
                        for i0_38_ in range(nw266_.length(0)):
                            nw266_[i0_38_] = init38_(i0_38_)
                        d_1411_v55_ = nw266_
                        d_1414_v56_: _dafny.Seq
                        d_1414_v56_ = _dafny.SeqWithoutIsStrInference([(d_1392_v36_).f23])
                        d_1415_v57_: _dafny.MultiSet
                        d_1415_v57_ = _dafny.MultiSet([(d_1414_v56_)[default__.safeIndex((d_1389_v33_)[default__.safeIndex(301, (d_1389_v33_).length(0))], len(d_1414_v56_))], not(p0), p0])
                        rhs192_ = d_1411_v55_
                        rhs193_ = d_1415_v57_
                        rhs194_ = not((default__.fm2(globalState) if False else p0))
                        rhs195_ = False
                        lhs169_ = globalState
                        lhs170_ = globalState
                        d_1411_v55_ = rhs192_
                        d_1415_v57_ = rhs193_
                        lhs169_.f1 = rhs194_
                        lhs170_.f1 = rhs195_
                        d_1408_v52_ = ((d_1408_v52_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "siy")))) + (d_1408_v52_)
                    d_1416_v58_: _dafny.Map
                    d_1416_v58_ = _dafny.Map({p0: d_1388_v32_})
                    d_1417_v59_: _dafny.Array
                    nw267_ = _dafny.Array(None, 26)
                    nw267_[int(0)] = d_1388_v32_
                    nw267_[int(1)] = d_1388_v32_
                    nw267_[int(2)] = d_1388_v32_
                    nw267_[int(3)] = d_1388_v32_
                    nw267_[int(4)] = d_1388_v32_
                    nw267_[int(5)] = d_1388_v32_
                    nw267_[int(6)] = d_1388_v32_
                    nw267_[int(7)] = d_1388_v32_
                    nw267_[int(8)] = d_1388_v32_
                    nw267_[int(9)] = d_1388_v32_
                    nw267_[int(10)] = d_1388_v32_
                    nw267_[int(11)] = d_1388_v32_
                    nw267_[int(12)] = d_1388_v32_
                    nw267_[int(13)] = d_1388_v32_
                    nw267_[int(14)] = d_1388_v32_
                    nw267_[int(15)] = d_1388_v32_
                    nw267_[int(16)] = ((d_1416_v58_)[(d_1392_v36_).f23] if ((d_1392_v36_).f23) in (d_1416_v58_) else d_1388_v32_)
                    nw267_[int(17)] = d_1388_v32_
                    nw267_[int(18)] = d_1388_v32_
                    nw267_[int(19)] = d_1388_v32_
                    nw267_[int(20)] = d_1388_v32_
                    nw267_[int(21)] = d_1388_v32_
                    nw267_[int(22)] = d_1388_v32_
                    nw267_[int(23)] = d_1388_v32_
                    nw267_[int(24)] = d_1388_v32_
                    nw267_[int(25)] = d_1388_v32_
                    d_1417_v59_ = nw267_
                    d_1418_v60_: _dafny.Map
                    d_1418_v60_ = _dafny.Map({d_1417_v59_: d_1348_v1_})
                    (globalState).f6 = ((d_1418_v60_)[d_1417_v59_] if (d_1417_v59_) in (d_1418_v60_) else d_1348_v1_)
                    pass
            pass
        if default__.fm2(globalState):
            d_1419_v61_: _dafny.Map
            d_1419_v61_ = _dafny.Map({d_1348_v1_: default__.fm21(globalState)})
            d_1420_v62_: _dafny.Seq
            d_1420_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxdqb"))
            d_1421_v63_: _dafny.Array
            nw268_ = _dafny.Array(None, 8)
            nw268_[int(0)] = d_1348_v1_
            nw268_[int(1)] = d_1348_v1_
            nw268_[int(2)] = d_1348_v1_
            nw268_[int(3)] = d_1348_v1_
            nw268_[int(4)] = d_1348_v1_
            nw268_[int(5)] = d_1348_v1_
            nw268_[int(6)] = d_1348_v1_
            nw268_[int(7)] = len(((d_1419_v61_)[d_1348_v1_] if (d_1348_v1_) in (d_1419_v61_) else d_1420_v62_))
            d_1421_v63_ = nw268_
            d_1422_v64_: _dafny.Seq
            d_1422_v64_ = _dafny.SeqWithoutIsStrInference([d_1421_v63_])
            d_1422_v64_ = _dafny.SeqWithoutIsStrInference([d_1421_v63_, d_1421_v63_])
            (globalState).f1 = p0
            d_1423_v65_: _dafny.Map
            d_1423_v65_ = _dafny.Map({d_1348_v1_: len(default__.fm21(globalState))})
            d_1424_v66_: _dafny.Seq
            d_1424_v66_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1425_v67_: _dafny.Map
            d_1425_v67_ = _dafny.Map({226: p0})
            d_1426_v68_: _dafny.Map
            d_1426_v68_ = _dafny.Map({d_1425_v67_: d_1348_v1_})
            d_1427_v69_: _dafny.Map
            d_1427_v69_ = _dafny.Map({p0: _dafny.Map({d_1425_v67_: len(d_1424_v66_)})})
            d_1428_v71_: _dafny.MultiSet
            d_1428_v71_ = _dafny.MultiSet([d_1425_v67_, d_1425_v67_])
            d_1429_v72_: C3
            nw269_ = C3()
            def iife174_():
                coll60_ = _dafny.Map()
                compr_60_: _dafny.Map
                for compr_60_ in (d_1428_v71_).Elements:
                    d_1430_v70_: _dafny.Map = compr_60_
                    if (d_1430_v70_) in (d_1428_v71_):
                        coll60_[d_1430_v70_] = d_1348_v1_
                return _dafny.Map(coll60_)
            nw269_.ctor__(p0, (0) - (((d_1423_v65_)[(0) - ((0) - (len(d_1424_v66_)))] if ((0) - ((0) - (len(d_1424_v66_)))) in (d_1423_v65_) else d_1348_v1_)), (d_1426_v68_) | (((d_1427_v69_)[p0] if (p0) in (d_1427_v69_) else iife174_()
            )))
            d_1429_v72_ = nw269_
            (globalState).f1 = (d_1429_v72_).f18
            (globalState).f1 = p0
        elif True:
            (globalState).f6 = d_1348_v1_
            (globalState).f1 = p0
            (globalState).f1 = (default__.fm2(globalState) if p0 else p0)
            d_1431_v73_: str
            d_1431_v73_ = _dafny.CodePoint('a')
            d_1432_v74_: _dafny.Seq
            d_1432_v74_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yehtu"))
            d_1433_v75_: _dafny.Map
            d_1433_v75_ = _dafny.Map({p0: 166})
            d_1431_v73_ = (d_1432_v74_)[default__.safeIndex((len(d_1433_v75_)) - (d_1348_v1_), len(d_1432_v74_))]
            d_1434_v76_: _dafny.MultiSet
            d_1434_v76_ = _dafny.MultiSet([p0, p0])
            (globalState).f6 = (d_1348_v1_) - (((d_1434_v76_)[p0] if (p0) in (d_1434_v76_) else d_1348_v1_))

    def m10(self, globalState):
        d_1435_v0_: bool
        d_1435_v0_ = True
        d_1436_v1_: _dafny.Map
        d_1436_v1_ = _dafny.Map({(d_1435_v0_) == (d_1435_v0_): default__.fm1(d_1435_v0_, globalState)})
        d_1436_v1_ = d_1436_v1_
        d_1437_v2_: _dafny.Seq
        d_1437_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "as"))
        d_1438_v3_: str
        d_1438_v3_ = _dafny.CodePoint('m')
        d_1439_i0_: int
        d_1439_i0_ = 0
        with _dafny.label("11"):
            while not((d_1435_v0_) or (((d_1437_v2_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.Set({False}), _dafny.Set({d_1435_v0_})])) for d_1469_i1_ in range(default__.abs(44))])), len(d_1437_v2_)), d_1438_v3_)) <= (d_1437_v2_))):
                with _dafny.c_label("11"):
                    if (d_1439_i0_) >= (100):
                        raise _dafny.Break("11")
                    d_1439_i0_ = (d_1439_i0_) + (1)
                    d_1440_v4_: int
                    d_1440_v4_ = -666
                    d_1441_v5_: _dafny.Seq
                    d_1441_v5_ = _dafny.SeqWithoutIsStrInference([d_1435_v0_, d_1435_v0_, d_1435_v0_, False])
                    d_1442_v6_: _dafny.Map
                    d_1442_v6_ = _dafny.Map({695: 16})
                    d_1443_v7_: D12
                    d_1443_v7_ = D12_DC34(d_1442_v6_)
                    d_1444_v8_: _dafny.Array
                    nw270_ = _dafny.Array(None, 16)
                    nw270_[int(0)] = (d_1440_v4_ if True else d_1440_v4_)
                    nw270_[int(1)] = d_1440_v4_
                    nw270_[int(2)] = d_1440_v4_
                    nw270_[int(3)] = default__.fm1(d_1435_v0_, globalState)
                    nw270_[int(4)] = d_1440_v4_
                    nw270_[int(5)] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1435_v0_])) for d_1445_i2_ in range(default__.abs(629))])), d_1440_v4_, d_1440_v4_, d_1440_v4_]))
                    nw270_[int(6)] = d_1440_v4_
                    nw270_[int(7)] = (0) - (d_1440_v4_)
                    nw270_[int(8)] = len(d_1441_v5_)
                    nw270_[int(9)] = len(d_1437_v2_)
                    nw270_[int(10)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btflmnhpm")))
                    nw270_[int(11)] = (len(_dafny.Map({d_1440_v4_: d_1443_v7_}))) * (d_1440_v4_)
                    nw270_[int(12)] = default__.fm1(True, globalState)
                    nw270_[int(13)] = (default__.fm1(d_1435_v0_, globalState)) - (d_1440_v4_)
                    nw270_[int(14)] = d_1440_v4_
                    nw270_[int(15)] = d_1440_v4_
                    d_1444_v8_ = nw270_
                    index247_ = default__.safeIndex(246, (d_1444_v8_).length(0))
                    (d_1444_v8_)[index247_] = d_1440_v4_
                    d_1446_v9_: _dafny.MultiSet
                    d_1446_v9_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1438_v3_ for d_1447_i3_ in range(default__.abs(192))]), d_1437_v2_])
                    index248_ = default__.safeIndex(246, (d_1444_v8_).length(0))
                    (d_1444_v8_)[index248_] = (0) - (((d_1446_v9_)[d_1437_v2_] if (d_1437_v2_) in (d_1446_v9_) else d_1440_v4_))
                    d_1448_v10_: _dafny.Set
                    d_1448_v10_ = _dafny.Set({not(not(default__.fm2(globalState)))})
                    d_1448_v10_ = d_1448_v10_
                    if True:
                        d_1449_v11_: _dafny.Array
                        nw271_ = _dafny.Array(False, 24)
                        d_1449_v11_ = nw271_
                        index249_ = default__.safeIndex(818, (d_1449_v11_).length(0))
                        (d_1449_v11_)[index249_] = d_1435_v0_
                        index250_ = default__.safeIndex(818, (d_1449_v11_).length(0))
                        (d_1449_v11_)[index250_] = d_1435_v0_
                        d_1450_v12_: _dafny.Set
                        d_1450_v12_ = _dafny.Set({(d_1444_v8_)[default__.safeIndex(246, (d_1444_v8_).length(0))], (d_1444_v8_)[default__.safeIndex(246, (d_1444_v8_).length(0))]})
                        d_1451_v13_: D15
                        d_1451_v13_ = D15_DC42(d_1450_v12_)
                        index251_ = default__.safeIndex(246, (d_1444_v8_).length(0))
                        rhs196_ = d_1440_v4_
                        rhs197_ = d_1451_v13_
                        lhs171_ = d_1444_v8_
                        lhs172_ = default__.safeIndex(246, (d_1444_v8_).length(0))
                        lhs171_[lhs172_] = rhs196_
                        d_1451_v13_ = rhs197_
                        d_1452_v14_: _dafny.MultiSet
                        d_1452_v14_ = _dafny.MultiSet([default__.fm2(globalState)])
                        d_1453_v15_: _dafny.Array
                        nw272_ = _dafny.Array(_dafny.Array(None, 0), 23)
                        d_1453_v15_ = nw272_
                        d_1454_v16_: _dafny.Map
                        d_1454_v16_ = _dafny.Map({(d_1452_v14_).issubset(d_1452_v14_): d_1453_v15_})
                        d_1454_v16_ = (d_1454_v16_).set(d_1435_v0_, d_1453_v15_)
                        (globalState).f6 = (0) - ((0) - (default__.safeDivisionInt(d_1440_v4_, ((d_1444_v8_)[default__.safeIndex(246, (d_1444_v8_).length(0))] if (d_1449_v11_)[default__.safeIndex(818, (d_1449_v11_).length(0))] else default__.fm1(d_1435_v0_, globalState)))))
                        d_1455_v17_: T0
                        nw273_ = C1()
                        nw273_.ctor__((d_1444_v8_)[default__.safeIndex(246, (d_1444_v8_).length(0))])
                        d_1455_v17_ = nw273_
                        d_1456_v18_: D7
                        d_1456_v18_ = D7_DC15(d_1455_v17_)
                        d_1457_v19_: _dafny.Seq
                        d_1457_v19_ = _dafny.SeqWithoutIsStrInference([d_1455_v17_, d_1455_v17_])
                        d_1458_v20_: _dafny.Array
                        nw274_ = _dafny.Array(None, 16)
                        nw274_[int(0)] = d_1456_v18_
                        nw274_[int(1)] = d_1456_v18_
                        nw274_[int(2)] = d_1456_v18_
                        nw274_[int(3)] = d_1456_v18_
                        nw274_[int(4)] = d_1456_v18_
                        nw274_[int(5)] = d_1456_v18_
                        nw274_[int(6)] = d_1456_v18_
                        nw274_[int(7)] = d_1456_v18_
                        nw274_[int(8)] = d_1456_v18_
                        nw274_[int(9)] = D7_DC15((d_1457_v19_)[default__.safeIndex((d_1444_v8_)[default__.safeIndex(246, (d_1444_v8_).length(0))], len(d_1457_v19_))])
                        nw274_[int(10)] = d_1456_v18_
                        nw274_[int(11)] = d_1456_v18_
                        nw274_[int(12)] = d_1456_v18_
                        nw274_[int(13)] = d_1456_v18_
                        nw274_[int(14)] = d_1456_v18_
                        nw274_[int(15)] = d_1456_v18_
                        d_1458_v20_ = nw274_
                        d_1459_v21_: _dafny.MultiSet
                        d_1459_v21_ = _dafny.MultiSet([d_1458_v20_, d_1458_v20_, d_1458_v20_])
                        d_1460_v22_: C2
                        nw275_ = C2()
                        nw275_.ctor__((d_1444_v8_)[default__.safeIndex(246, (d_1444_v8_).length(0))], d_1440_v4_, _dafny.Map({_dafny.Map({(d_1444_v8_)[default__.safeIndex(246, (d_1444_v8_).length(0))]: d_1435_v0_}): d_1440_v4_}))
                        d_1460_v22_ = nw275_
                        d_1461_v23_: _dafny.Seq
                        d_1461_v23_ = _dafny.SeqWithoutIsStrInference([d_1460_v22_])
                        d_1462_v24_: _dafny.Array
                        nw276_ = _dafny.Array(None, 20)
                        nw276_[int(0)] = d_1459_v21_
                        nw276_[int(1)] = d_1459_v21_
                        nw276_[int(2)] = d_1459_v21_
                        nw276_[int(3)] = _dafny.MultiSet([d_1458_v20_, d_1458_v20_, d_1458_v20_])
                        nw276_[int(4)] = d_1459_v21_
                        nw276_[int(5)] = d_1459_v21_
                        nw276_[int(6)] = d_1459_v21_
                        nw276_[int(7)] = d_1459_v21_
                        nw276_[int(8)] = ((d_1459_v21_).set(d_1458_v20_, default__.abs(d_1440_v4_))) - (d_1459_v21_)
                        nw276_[int(9)] = d_1459_v21_
                        nw276_[int(10)] = d_1459_v21_
                        nw276_[int(11)] = d_1459_v21_
                        nw276_[int(12)] = d_1459_v21_
                        nw276_[int(13)] = (d_1459_v21_).intersection((d_1459_v21_).set(d_1458_v20_, default__.abs((d_1444_v8_)[default__.safeIndex(246, (d_1444_v8_).length(0))])))
                        nw276_[int(14)] = _dafny.MultiSet([d_1458_v20_, d_1458_v20_, d_1458_v20_, d_1458_v20_, d_1458_v20_])
                        nw276_[int(15)] = (d_1459_v21_).intersection(d_1459_v21_)
                        nw276_[int(16)] = _dafny.MultiSet([d_1458_v20_])
                        nw276_[int(17)] = (d_1459_v21_).set(d_1458_v20_, default__.abs(len(d_1461_v23_)))
                        nw276_[int(18)] = d_1459_v21_
                        nw276_[int(19)] = d_1459_v21_
                        d_1462_v24_ = nw276_
                        index252_ = default__.safeIndex(672, (d_1462_v24_).length(0))
                        (d_1462_v24_)[index252_] = d_1459_v21_
                        d_1463_v25_: D16
                        d_1463_v25_ = D16_DC44(d_1452_v14_)
                        pat_let_tv35_ = d_1452_v14_
                        index253_ = default__.safeIndex(672, (d_1462_v24_).length(0))
                        def iife175_(_pat_let57_0):
                            def iife176_(d_1464_dt__update__tmp_h0_):
                                def iife177_(_pat_let58_0):
                                    def iife178_(d_1465_dt__update_hcf77_h0_):
                                        return D16_DC44(d_1465_dt__update_hcf77_h0_)
                                    return iife178_(_pat_let58_0)
                                return iife177_(pat_let_tv35_)
                            return iife176_(_pat_let57_0)
                        rhs198_ = d_1460_v22_.f21
                        rhs199_ = ((d_1459_v21_).intersection(_dafny.MultiSet([d_1458_v20_, d_1458_v20_]))).intersection(d_1459_v21_)
                        rhs200_ = default__.fm2(globalState)
                        rhs201_ = (False) not in (((iife175_(d_1463_v25_)).cf77).intersection(d_1452_v14_))
                        lhs173_ = globalState
                        lhs174_ = d_1462_v24_
                        lhs175_ = default__.safeIndex(672, (d_1462_v24_).length(0))
                        lhs176_ = globalState
                        lhs177_ = globalState
                        lhs173_.f6 = rhs198_
                        lhs174_[lhs175_] = rhs199_
                        lhs176_.f1 = rhs200_
                        lhs177_.f1 = rhs201_
                    elif True:
                        d_1435_v0_ = d_1435_v0_
                        d_1466_v26_: C1
                        nw277_ = C1()
                        nw277_.ctor__((d_1440_v4_) * (d_1440_v4_))
                        d_1466_v26_ = nw277_
                        index254_ = default__.safeIndex(246, (d_1444_v8_).length(0))
                        (d_1444_v8_)[index254_] = (d_1440_v4_) + (d_1440_v4_)
                        d_1467_v27_: _dafny.Map
                        d_1467_v27_ = _dafny.Map({(d_1466_v26_).fm24(not(d_1435_v0_), default__.fm39(d_1435_v0_, globalState), globalState): False})
                        d_1467_v27_ = (d_1467_v27_).set(not(d_1435_v0_), not (True) or (d_1435_v0_))
                        index255_ = default__.safeIndex(246, (d_1444_v8_).length(0))
                        (d_1444_v8_)[index255_] = len(d_1441_v5_)
                    d_1468_v28_: _dafny.Set
                    d_1468_v28_ = _dafny.Set({d_1444_v8_, d_1444_v8_, d_1444_v8_})
                    d_1468_v28_ = d_1468_v28_
                    pass
            pass
        d_1470_v29_: int
        d_1470_v29_ = 538
        d_1471_v30_: _dafny.Array
        nw278_ = _dafny.Array(None, 5)
        d_1471_v30_ = nw278_
        d_1472_v31_: _dafny.Map
        d_1472_v31_ = _dafny.Map({d_1438_v3_: d_1470_v29_})
        d_1473_v32_: _dafny.Map
        d_1473_v32_ = _dafny.Map({d_1471_v30_: d_1472_v31_})
        hi9_ = default__.safeDivisionInt(len(d_1473_v32_), d_1470_v29_)
        for d_1474_i4_ in range(d_1470_v29_, hi9_):
            d_1475_v33_: _dafny.Array
            nw279_ = _dafny.Array(False, 22)
            d_1475_v33_ = nw279_
            (globalState).f14 = d_1475_v33_
            d_1476_v34_: _dafny.Map
            d_1476_v34_ = _dafny.Map({(d_1435_v0_ if d_1435_v0_ else d_1435_v0_): d_1438_v3_})
            d_1477_v35_: _dafny.Map
            d_1477_v35_ = _dafny.Map({d_1435_v0_: default__.fm29(d_1435_v0_, d_1470_v29_, d_1474_i4_, True, globalState)})
            d_1476_v34_ = (d_1477_v35_)
            d_1478_v36_: _dafny.Array
            nw280_ = _dafny.Array(int(0), 19)
            d_1478_v36_ = nw280_
            index256_ = default__.safeIndex(106, (d_1478_v36_).length(0))
            (d_1478_v36_)[index256_] = d_1470_v29_
            index257_ = default__.safeIndex(106, (d_1478_v36_).length(0))
            (d_1478_v36_)[index257_] = d_1470_v29_
            d_1437_v2_ = (default__.fm21(globalState)).set(default__.safeIndex((d_1478_v36_)[default__.safeIndex(106, (d_1478_v36_).length(0))], len(default__.fm21(globalState))), d_1438_v3_)
        (globalState).f6 = d_1470_v29_
        if d_1435_v0_:
            d_1479_v37_: _dafny.Set
            d_1479_v37_ = _dafny.Set({d_1470_v29_})
            d_1480_v38_: _dafny.Map
            d_1480_v38_ = _dafny.Map({d_1479_v37_: d_1437_v2_})
            d_1481_v39_: _dafny.Map
            d_1481_v39_ = _dafny.Map({d_1470_v29_: ((d_1436_v1_)[not(d_1435_v0_)] if (not(d_1435_v0_)) in (d_1436_v1_) else len(((d_1480_v38_)[d_1479_v37_] if (d_1479_v37_) in (d_1480_v38_) else d_1437_v2_)))})
            d_1482_v40_: _dafny.Array
            def lambda63_(d_1483_v0_, d_1484_v2_):
                def lambda64_(d_1485_i5_):
                    return (_dafny.Map({d_1483_v0_: d_1484_v2_})) | (_dafny.Map({d_1483_v0_: d_1484_v2_}))

                return lambda64_

            init39_ = lambda63_(d_1435_v0_, d_1437_v2_)
            nw281_ = _dafny.Array(None, 2)
            for i0_39_ in range(nw281_.length(0)):
                nw281_[i0_39_] = init39_(i0_39_)
            d_1482_v40_ = nw281_
            d_1486_v41_: _dafny.Map
            d_1486_v41_ = _dafny.Map({d_1435_v0_: d_1437_v2_})
            index258_ = default__.safeIndex(264, (d_1482_v40_).length(0))
            (d_1482_v40_)[index258_] = d_1486_v41_
            d_1487_v43_: _dafny.Map
            d_1487_v43_ = _dafny.Map({d_1470_v29_: d_1435_v0_})
            d_1488_v44_: _dafny.Seq
            def iife179_():
                coll61_ = _dafny.Map()
                compr_61_: int
                for compr_61_ in (d_1487_v43_).keys.Elements:
                    d_1489_v42_: int = compr_61_
                    if (d_1489_v42_) in (d_1487_v43_):
                        coll61_[(d_1489_v42_) * (d_1470_v29_)] = d_1470_v29_
                return _dafny.Map(coll61_)
            d_1488_v44_ = _dafny.SeqWithoutIsStrInference([d_1481_v39_, iife179_()
            , d_1481_v39_, d_1481_v39_])
            index259_ = default__.safeIndex(264, (d_1482_v40_).length(0))
            rhs202_ = (d_1488_v44_)[default__.safeIndex((0) - (d_1470_v29_), len(d_1488_v44_))]
            rhs203_ = d_1486_v41_
            rhs204_ = default__.fm29(d_1435_v0_, d_1470_v29_, default__.fm1(d_1435_v0_, globalState), False, globalState)
            lhs178_ = d_1482_v40_
            lhs179_ = default__.safeIndex(264, (d_1482_v40_).length(0))
            d_1481_v39_ = rhs202_
            lhs178_[lhs179_] = rhs203_
            d_1438_v3_ = rhs204_
            d_1435_v0_ = d_1435_v0_
            (globalState).f6 = 788
            if d_1435_v0_:
                (globalState).f6 = ((d_1470_v29_) - (-168)) + (d_1470_v29_)
                (globalState).f1 = True
                d_1490_v45_: C1
                nw282_ = C1()
                nw282_.ctor__(len(d_1437_v2_))
                d_1490_v45_ = nw282_
                d_1491_v46_: _dafny.Map
                d_1491_v46_ = _dafny.Map({d_1435_v0_: d_1435_v0_})
                (globalState).f6 = (0) - ((0) - ((len((d_1491_v46_).set(d_1435_v0_, d_1435_v0_))) - (-717)))
                d_1492_v47_: _dafny.Array
                def lambda65_(d_1493_v29_, d_1494_v0_, d_1495_v2_, d_1496_v45_, d_1497_v37_):
                    def lambda66_(d_1498_i6_):
                        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([D9_DC24(_dafny.Map({d_1493_v29_: _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oorciw"))): _dafny.Set({d_1493_v29_})})}), _dafny.SeqWithoutIsStrInference([d_1494_v0_])), D9_DC24(_dafny.Map({len(d_1495_v2_): _dafny.Map({(d_1496_v45_).f24: d_1497_v37_})}), _dafny.SeqWithoutIsStrInference([d_1494_v0_, d_1494_v0_]))]): 86})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([D9_DC24(_dafny.Map({d_1493_v29_: _dafny.Map({(d_1496_v45_).f24: d_1497_v37_})}), _dafny.SeqWithoutIsStrInference([d_1494_v0_, d_1494_v0_])) for d_1499_i7_ in range(default__.abs(802))]): 726}))

                    return lambda66_

                init40_ = lambda65_(d_1470_v29_, d_1435_v0_, d_1437_v2_, d_1490_v45_, d_1479_v37_)
                nw283_ = _dafny.Array(None, 28)
                for i0_40_ in range(nw283_.length(0)):
                    nw283_[i0_40_] = init40_(i0_40_)
                d_1492_v47_ = nw283_
                d_1500_v49_: _dafny.Map
                def iife180_():
                    coll62_ = _dafny.Map()
                    compr_62_: int
                    for compr_62_ in _dafny.IntegerRange(-144, 515):
                        d_1501_v48_: int = compr_62_
                        if ((-144) <= (d_1501_v48_)) and ((d_1501_v48_) < (515)):
                            coll62_[default__.safeDivisionInt(d_1501_v48_, (d_1490_v45_).f24)] = d_1479_v37_
                    return _dafny.Map(coll62_)
                d_1500_v49_ = iife180_()
                
                d_1502_v50_: _dafny.Map
                d_1502_v50_ = _dafny.Map({(d_1490_v45_).f24: d_1500_v49_})
                d_1503_v51_: _dafny.Seq
                d_1503_v51_ = _dafny.SeqWithoutIsStrInference([d_1435_v0_, d_1435_v0_, d_1435_v0_, d_1435_v0_])
                d_1504_v52_: D9
                d_1504_v52_ = D9_DC24(d_1502_v50_, d_1503_v51_)
                d_1505_v53_: _dafny.Map
                d_1505_v53_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1504_v52_]): (0) - (d_1470_v29_)})
                index260_ = default__.safeIndex(760, (d_1492_v47_).length(0))
                (d_1492_v47_)[index260_] = d_1505_v53_
                d_1506_v54_: _dafny.Seq
                d_1506_v54_ = _dafny.SeqWithoutIsStrInference([d_1436_v1_])
                d_1507_v55_: _dafny.MultiSet
                d_1507_v55_ = _dafny.MultiSet([d_1435_v0_, False])
                d_1508_v56_: _dafny.MultiSet
                d_1508_v56_ = _dafny.MultiSet([713, ((d_1507_v55_)[d_1435_v0_] if (d_1435_v0_) in (d_1507_v55_) else d_1470_v29_)])
                index261_ = default__.safeIndex(760, (d_1492_v47_).length(0))
                rhs205_ = d_1470_v29_
                rhs206_ = default__.fm40(d_1470_v29_, globalState)
                rhs207_ = (((d_1506_v54_).set(default__.safeIndex(len((d_1482_v40_)[default__.safeIndex(264, (d_1482_v40_).length(0))]), len(d_1506_v54_)), d_1436_v1_)) + (d_1506_v54_)) + (d_1506_v54_)
                rhs208_ = -133
                rhs209_ = ((default__.fm1(d_1435_v0_, globalState)) - (((d_1508_v56_)[(d_1490_v45_).f24] if ((d_1490_v45_).f24) in (d_1508_v56_) else (d_1490_v45_).f24))) - (default__.safeModuloInt((d_1490_v45_).f24, d_1470_v29_))
                lhs180_ = d_1492_v47_
                lhs181_ = default__.safeIndex(760, (d_1492_v47_).length(0))
                lhs182_ = globalState
                d_1470_v29_ = rhs205_
                lhs180_[lhs181_] = rhs206_
                d_1506_v54_ = rhs207_
                lhs182_.f6 = rhs208_
                d_1470_v29_ = rhs209_
            elif True:
                d_1509_v57_: _dafny.Map
                d_1509_v57_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))): d_1437_v2_})
                d_1510_v58_: _dafny.Array
                nw284_ = _dafny.Array(None, 23)
                nw284_[int(0)] = d_1437_v2_
                nw284_[int(1)] = default__.fm21(globalState)
                nw284_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfhwoiw"))
                nw284_[int(3)] = d_1437_v2_
                nw284_[int(4)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1511_i8_ in range(default__.abs(79))])) + (d_1437_v2_)
                nw284_[int(5)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))) + (d_1437_v2_)
                nw284_[int(6)] = d_1437_v2_
                nw284_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjvd"))
                nw284_[int(8)] = d_1437_v2_
                nw284_[int(9)] = d_1437_v2_
                nw284_[int(10)] = (d_1437_v2_) + (d_1437_v2_)
                nw284_[int(11)] = d_1437_v2_
                nw284_[int(12)] = (d_1437_v2_) + (d_1437_v2_)
                nw284_[int(13)] = d_1437_v2_
                nw284_[int(14)] = d_1437_v2_
                nw284_[int(15)] = ((d_1509_v57_)[d_1470_v29_] if (d_1470_v29_) in (d_1509_v57_) else d_1437_v2_)
                nw284_[int(16)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gqj"))) + (_dafny.SeqWithoutIsStrInference([d_1438_v3_ for d_1512_i9_ in range(default__.abs(882))]))
                nw284_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gv"))
                nw284_[int(18)] = d_1437_v2_
                nw284_[int(19)] = d_1437_v2_
                nw284_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjbixrfgg"))
                nw284_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oarbcrux"))
                nw284_[int(22)] = (d_1437_v2_) + (d_1437_v2_)
                d_1510_v58_ = nw284_
                d_1513_v59_: _dafny.Seq
                d_1513_v59_ = _dafny.SeqWithoutIsStrInference([d_1437_v2_, d_1437_v2_, d_1437_v2_, d_1437_v2_, d_1437_v2_])
                index262_ = default__.safeIndex(316, (d_1510_v58_).length(0))
                (d_1510_v58_)[index262_] = (d_1513_v59_)[default__.safeIndex(d_1470_v29_, len(d_1513_v59_))]
                index263_ = default__.safeIndex(316, (d_1510_v58_).length(0))
                (d_1510_v58_)[index263_] = d_1437_v2_
                d_1514_v60_: _dafny.MultiSet
                d_1514_v60_ = _dafny.MultiSet([not (d_1435_v0_) or (d_1435_v0_), not(not((d_1435_v0_) or (False))), d_1435_v0_, (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqtqm")))) >= (d_1470_v29_), d_1435_v0_])
                rhs210_ = (d_1514_v60_).cardinality
                rhs211_ = (default__.fm1(d_1435_v0_, globalState)) != (default__.fm1(d_1435_v0_, globalState))
                lhs183_ = globalState
                lhs184_ = globalState
                lhs183_.f6 = rhs210_
                lhs184_.f1 = rhs211_
                (globalState).f6 = (d_1470_v29_) * (default__.fm1(d_1435_v0_, globalState))
                d_1515_v61_: _dafny.Map
                d_1515_v61_ = _dafny.Map({(d_1510_v58_)[default__.safeIndex(316, (d_1510_v58_).length(0))]: d_1470_v29_})
                d_1516_v62_: _dafny.Array
                def lambda67_(d_1517_v29_):
                    def lambda68_(d_1518_i10_):
                        return default__.safeDivisionInt(d_1518_i10_, d_1517_v29_)

                    return lambda68_

                init41_ = lambda67_(d_1470_v29_)
                nw285_ = _dafny.Array(None, 24)
                for i0_41_ in range(nw285_.length(0)):
                    nw285_[i0_41_] = init41_(i0_41_)
                d_1516_v62_ = nw285_
                d_1519_v63_: _dafny.Seq
                d_1519_v63_ = _dafny.SeqWithoutIsStrInference([d_1435_v0_, d_1435_v0_, not(False)])
                d_1520_v64_: D12
                d_1520_v64_ = D12_DC34(_dafny.Map({(_dafny.MultiSet(d_1519_v63_)).cardinality: len(d_1437_v2_)}))
                index264_ = default__.safeIndex(879, (d_1516_v62_).length(0))
                (d_1516_v62_)[index264_] = len(default__.fm36(d_1438_v3_, d_1520_v64_, globalState))
                d_1521_v65_: _dafny.MultiSet
                d_1521_v65_ = _dafny.MultiSet([d_1470_v29_])
                d_1522_v66_: _dafny.Seq
                d_1522_v66_ = _dafny.SeqWithoutIsStrInference([d_1521_v65_])
                d_1523_v67_: _dafny.Seq
                d_1523_v67_ = _dafny.SeqWithoutIsStrInference([d_1470_v29_])
                d_1524_v68_: _dafny.Set
                d_1524_v68_ = _dafny.Set({(d_1510_v58_)[default__.safeIndex(316, (d_1510_v58_).length(0))], (d_1437_v2_).set(default__.safeIndex(d_1470_v29_, len(d_1437_v2_)), d_1438_v3_)})
                index265_ = default__.safeIndex(879, (d_1516_v62_).length(0))
                rhs212_ = (957) - ((0) - (d_1470_v29_))
                rhs213_ = ((d_1515_v61_ if d_1435_v0_ else d_1515_v61_)) | (d_1515_v61_)
                rhs214_ = (d_1470_v29_) * (d_1470_v29_)
                rhs215_ = not(((d_1522_v66_) + (default__.fm41(len(d_1523_v67_), globalState))) != (_dafny.SeqWithoutIsStrInference([d_1521_v65_])))
                rhs216_ = (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxbpkigo"))})).isdisjoint((d_1524_v68_).intersection(d_1524_v68_))
                lhs185_ = globalState
                lhs186_ = d_1516_v62_
                lhs187_ = default__.safeIndex(879, (d_1516_v62_).length(0))
                lhs188_ = globalState
                lhs189_ = globalState
                lhs185_.f6 = rhs212_
                d_1515_v61_ = rhs213_
                lhs186_[lhs187_] = rhs214_
                lhs188_.f1 = rhs215_
                lhs189_.f1 = rhs216_
                d_1525_v69_: _dafny.Array
                nw286_ = _dafny.Array(False, 6)
                d_1525_v69_ = nw286_
                index266_ = default__.safeIndex(576, (d_1525_v69_).length(0))
                (d_1525_v69_)[index266_] = d_1435_v0_
                d_1526_v70_: _dafny.Map
                d_1526_v70_ = _dafny.Map({(d_1516_v62_)[default__.safeIndex(879, (d_1516_v62_).length(0))]: d_1516_v62_})
                index267_ = default__.safeIndex(576, (d_1525_v69_).length(0))
                (d_1525_v69_)[index267_] = (len(d_1526_v70_)) != (default__.safeModuloInt((0) - (d_1470_v29_), (d_1516_v62_)[default__.safeIndex(879, (d_1516_v62_).length(0))]))
            d_1527_v71_: _dafny.Set
            d_1527_v71_ = _dafny.Set({not(d_1435_v0_), d_1435_v0_})
            d_1528_v72_: D8
            d_1528_v72_ = D8_DC21(d_1437_v2_, d_1527_v71_)
            source11_ = d_1528_v72_
            if source11_.is_DC20:
                d_1529___mcc_h0_ = source11_.cf31
                d_1530___mcc_h1_ = source11_.cf32
                d_1531___mcc_h2_ = source11_.cf33
                d_1532_cf33_ = d_1531___mcc_h2_
                d_1533_cf32_ = d_1530___mcc_h1_
                d_1534_cf31_ = d_1529___mcc_h0_
                (globalState).f1 = not (d_1435_v0_) or (d_1435_v0_)
                (globalState).f6 = default__.fm1((_dafny.Set({len(d_1437_v2_), d_1470_v29_})).issubset(d_1479_v37_), globalState)
                d_1535_v73_: _dafny.Map
                d_1535_v73_ = _dafny.Map({d_1435_v0_: d_1435_v0_})
                d_1536_v74_: _dafny.MultiSet
                d_1536_v74_ = _dafny.MultiSet([len(d_1535_v73_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnmgv")))])
                d_1537_v75_: _dafny.Map
                d_1537_v75_ = _dafny.Map({d_1533_cf32_: default__.fm42(d_1533_cf32_, d_1438_v3_, d_1534_cf31_, d_1435_v0_, globalState)})
                d_1538_v76_: _dafny.MultiSet
                d_1538_v76_ = _dafny.MultiSet([d_1536_v74_, d_1536_v74_, ((d_1537_v75_)[d_1533_cf32_] if (d_1533_cf32_) in (d_1537_v75_) else d_1536_v74_), d_1536_v74_, _dafny.MultiSet([d_1532_cf33_, d_1470_v29_])])
                d_1539_v77_: _dafny.Array
                nw287_ = _dafny.Array(False, 15)
                d_1539_v77_ = nw287_
                d_1540_v78_: _dafny.MultiSet
                d_1540_v78_ = _dafny.MultiSet([d_1539_v77_])
                d_1541_v79_: _dafny.Seq
                d_1541_v79_ = _dafny.SeqWithoutIsStrInference([d_1470_v29_])
                d_1542_v80_: _dafny.Seq
                d_1542_v80_ = _dafny.SeqWithoutIsStrInference([d_1536_v74_, _dafny.MultiSet([d_1470_v29_])])
                d_1543_v81_: _dafny.Array
                nw288_ = _dafny.Array(None, 29)
                nw288_[int(0)] = _dafny.MultiSet([d_1536_v74_])
                nw288_[int(1)] = d_1538_v76_
                nw288_[int(2)] = d_1538_v76_
                nw288_[int(3)] = d_1538_v76_
                nw288_[int(4)] = _dafny.MultiSet([_dafny.MultiSet([len(d_1479_v37_)]), d_1536_v74_])
                nw288_[int(5)] = d_1538_v76_
                nw288_[int(6)] = d_1538_v76_
                nw288_[int(7)] = d_1538_v76_
                nw288_[int(8)] = (d_1538_v76_).set((d_1536_v74_).set(((d_1540_v78_)[d_1539_v77_] if (d_1539_v77_) in (d_1540_v78_) else d_1470_v29_), default__.abs(d_1470_v29_)), default__.abs(-251))
                nw288_[int(9)] = d_1538_v76_
                nw288_[int(10)] = default__.fm43(d_1435_v0_, default__.fm2(globalState), d_1435_v0_, d_1532_cf33_, globalState)
                nw288_[int(11)] = d_1538_v76_
                nw288_[int(12)] = d_1538_v76_
                nw288_[int(13)] = d_1538_v76_
                nw288_[int(14)] = (d_1538_v76_ if d_1435_v0_ else d_1538_v76_)
                nw288_[int(15)] = (d_1538_v76_) | (_dafny.MultiSet([d_1536_v74_, d_1536_v74_]))
                nw288_[int(16)] = d_1538_v76_
                nw288_[int(17)] = default__.fm43(d_1435_v0_, d_1435_v0_, not(d_1435_v0_), -376, globalState)
                nw288_[int(18)] = (d_1538_v76_).intersection(d_1538_v76_)
                nw288_[int(19)] = d_1538_v76_
                nw288_[int(20)] = (d_1538_v76_) | (d_1538_v76_)
                nw288_[int(21)] = (d_1538_v76_) | (((d_1538_v76_).set(_dafny.MultiSet(d_1541_v79_), default__.abs(d_1470_v29_))).set(d_1536_v74_, default__.abs(len(_dafny.SeqWithoutIsStrInference([D9_DC24(_dafny.Map({d_1533_cf32_: _dafny.Map({d_1532_cf33_: d_1479_v37_})}), _dafny.SeqWithoutIsStrInference([d_1435_v0_, d_1435_v0_])) for d_1544_i11_ in range(default__.abs(313))])))))
                nw288_[int(22)] = (_dafny.MultiSet(d_1542_v80_)) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_1536_v74_).set(((d_1436_v1_)[d_1435_v0_] if (d_1435_v0_) in (d_1436_v1_) else d_1532_cf33_), default__.abs(d_1533_cf32_)), d_1536_v74_])))
                nw288_[int(23)] = (d_1538_v76_) - (d_1538_v76_)
                nw288_[int(24)] = (d_1538_v76_) - (d_1538_v76_)
                nw288_[int(25)] = d_1538_v76_
                nw288_[int(26)] = (_dafny.MultiSet(d_1542_v80_)) - (d_1538_v76_)
                nw288_[int(27)] = d_1538_v76_
                nw288_[int(28)] = _dafny.MultiSet(d_1542_v80_)
                d_1543_v81_ = nw288_
                index268_ = default__.safeIndex(934, (d_1543_v81_).length(0))
                (d_1543_v81_)[index268_] = _dafny.MultiSet([d_1536_v74_])
                index269_ = default__.safeIndex(934, (d_1543_v81_).length(0))
                (d_1543_v81_)[index269_] = _dafny.MultiSet([(_dafny.MultiSet([(0) - (d_1533_cf32_)])) | (d_1536_v74_)])
                d_1545_v82_: C1
                nw289_ = C1()
                nw289_.ctor__(d_1534_cf31_)
                d_1545_v82_ = nw289_
                d_1546_v83_: _dafny.Map
                d_1546_v83_ = _dafny.Map({d_1532_cf33_: d_1545_v82_})
                d_1546_v83_ = (d_1546_v83_).set(d_1533_cf32_, d_1545_v82_)
            elif source11_.is_DC21:
                d_1547___mcc_h3_ = source11_.cf34
                d_1548___mcc_h4_ = source11_.cf35
                d_1549_cf35_ = d_1548___mcc_h4_
                d_1550_cf34_ = d_1547___mcc_h3_
                d_1551_v84_: _dafny.Array
                def lambda69_(d_1552_v2_):
                    def lambda70_(d_1553_i12_):
                        return d_1552_v2_

                    return lambda70_

                init42_ = lambda69_(d_1437_v2_)
                nw290_ = _dafny.Array(None, 13)
                for i0_42_ in range(nw290_.length(0)):
                    nw290_[i0_42_] = init42_(i0_42_)
                d_1551_v84_ = nw290_
                d_1551_v84_ = d_1551_v84_
                d_1554_v85_: _dafny.Array
                nw291_ = _dafny.Array(int(0), 18)
                d_1554_v85_ = nw291_
                d_1555_v86_: D3
                d_1555_v86_ = D3_DC6(d_1554_v85_)
                pat_let_tv36_ = d_1554_v85_
                d_1556_v87_: _dafny.Seq
                def iife181_(_pat_let59_0):
                    def iife182_(d_1557_dt__update__tmp_h1_):
                        def iife183_(_pat_let60_0):
                            def iife184_(d_1558_dt__update_hcf6_h0_):
                                return D3_DC6(d_1558_dt__update_hcf6_h0_)
                            return iife184_(_pat_let60_0)
                        return iife183_(pat_let_tv36_)
                    return iife182_(_pat_let59_0)
                d_1556_v87_ = _dafny.SeqWithoutIsStrInference([iife181_(d_1555_v86_), d_1555_v86_])
                d_1559_v88_: C0
                nw292_ = C0()
                nw292_.ctor__(d_1556_v87_, default__.fm2(globalState))
                d_1559_v88_ = nw292_
                d_1560_v89_: _dafny.Map
                d_1560_v89_ = _dafny.Map({d_1559_v88_: (d_1559_v88_).f23})
                d_1561_v90_: _dafny.Seq
                d_1561_v90_ = _dafny.SeqWithoutIsStrInference([len(d_1550_cf34_)])
                d_1562_v91_: _dafny.MultiSet
                d_1562_v91_ = _dafny.MultiSet([len(d_1561_v90_)])
                d_1563_v92_: _dafny.Array
                nw293_ = _dafny.Array(None, 7)
                nw293_[int(0)] = (0) - (len(d_1560_v89_))
                nw293_[int(1)] = (d_1562_v91_).cardinality
                nw293_[int(2)] = d_1470_v29_
                nw293_[int(3)] = (0) - (d_1470_v29_)
                nw293_[int(4)] = len(d_1436_v1_)
                nw293_[int(5)] = len(d_1550_cf34_)
                nw293_[int(6)] = d_1470_v29_
                d_1563_v92_ = nw293_
                d_1564_v93_: D3
                d_1564_v93_ = D3_DC6(d_1563_v92_)
                d_1565_v94_: _dafny.Seq
                d_1565_v94_ = _dafny.SeqWithoutIsStrInference([d_1564_v93_, D3_DC6(d_1554_v85_), d_1564_v93_, d_1555_v86_, d_1555_v86_])
                d_1566_v95_: C0
                nw294_ = C0()
                nw294_.ctor__((d_1556_v87_) + (_dafny.SeqWithoutIsStrInference([d_1555_v86_])), default__.fm2(globalState))
                d_1566_v95_ = nw294_
                d_1566_v95_ = d_1566_v95_
                (globalState).f6 = d_1470_v29_
                (globalState).f6 = (default__.safeModuloInt(d_1470_v29_, d_1470_v29_)) + ((d_1470_v29_) * (len(d_1561_v90_)))
            elif True:
                d_1567___mcc_h5_ = source11_.cf30
                d_1568_cf30_ = d_1567___mcc_h5_
                d_1569_v96_: _dafny.Seq
                d_1569_v96_ = _dafny.SeqWithoutIsStrInference([d_1470_v29_, d_1470_v29_, d_1470_v29_])
                d_1569_v96_ = d_1569_v96_
                d_1570_v97_: _dafny.Map
                d_1570_v97_ = _dafny.Map({d_1487_v43_: d_1470_v29_})
                d_1571_v98_: C2
                nw295_ = C2()
                nw295_.ctor__((d_1470_v29_) + (d_1470_v29_), d_1470_v29_, d_1570_v97_)
                d_1571_v98_ = nw295_
                d_1572_v99_: _dafny.Array
                nw296_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
                d_1572_v99_ = nw296_
                index270_ = default__.safeIndex(720, (d_1572_v99_).length(0))
                (d_1572_v99_)[index270_] = d_1437_v2_
                index271_ = default__.safeIndex(720, (d_1572_v99_).length(0))
                (d_1572_v99_)[index271_] = d_1437_v2_
                d_1573_v100_: _dafny.MultiSet
                d_1573_v100_ = _dafny.MultiSet([d_1435_v0_])
                d_1573_v100_ = d_1573_v100_
        elif True:
            d_1574_v101_: _dafny.Array
            def lambda71_(d_1575_v2_, d_1576_v29_, d_1577_v3_):
                def lambda72_(d_1578_i13_):
                    return ((d_1575_v2_).set(default__.safeIndex(d_1576_v29_, len(d_1575_v2_)), d_1577_v3_)) == (d_1575_v2_)

                return lambda72_

            init43_ = lambda71_(d_1437_v2_, d_1470_v29_, d_1438_v3_)
            nw297_ = _dafny.Array(None, 8)
            for i0_43_ in range(nw297_.length(0)):
                nw297_[i0_43_] = init43_(i0_43_)
            d_1574_v101_ = nw297_
            index272_ = default__.safeIndex(317, (d_1574_v101_).length(0))
            (d_1574_v101_)[index272_] = (d_1470_v29_) == (103)
            d_1579_v102_: _dafny.Array
            nw298_ = _dafny.Array(int(0), 23)
            d_1579_v102_ = nw298_
            index273_ = default__.safeIndex(372, (d_1579_v102_).length(0))
            (d_1579_v102_)[index273_] = d_1470_v29_
            d_1580_v103_: _dafny.Map
            d_1580_v103_ = _dafny.Map({d_1470_v29_: d_1470_v29_})
            d_1581_v104_: D12
            d_1581_v104_ = D12_DC34(d_1580_v103_)
            d_1582_v105_: _dafny.Seq
            d_1582_v105_ = _dafny.SeqWithoutIsStrInference([d_1435_v0_, d_1435_v0_, d_1435_v0_, d_1435_v0_])
            d_1583_v106_: _dafny.Map
            d_1583_v106_ = _dafny.Map({d_1435_v0_: d_1435_v0_})
            d_1584_v107_: _dafny.Array
            nw299_ = _dafny.Array(None, 29)
            nw299_[int(0)] = d_1436_v1_
            nw299_[int(1)] = d_1436_v1_
            nw299_[int(2)] = (default__.fm36(d_1438_v3_, d_1581_v104_, globalState)) | (d_1436_v1_)
            nw299_[int(3)] = d_1436_v1_
            nw299_[int(4)] = d_1436_v1_
            nw299_[int(5)] = d_1436_v1_
            nw299_[int(6)] = d_1436_v1_
            nw299_[int(7)] = _dafny.Map({d_1435_v0_: d_1470_v29_})
            nw299_[int(8)] = (d_1436_v1_) | (_dafny.Map({d_1435_v0_: d_1470_v29_}))
            nw299_[int(9)] = (d_1436_v1_) | ((d_1436_v1_).set(d_1435_v0_, len(d_1437_v2_)))
            nw299_[int(10)] = d_1436_v1_
            nw299_[int(11)] = d_1436_v1_
            nw299_[int(12)] = d_1436_v1_
            nw299_[int(13)] = d_1436_v1_
            nw299_[int(14)] = (d_1436_v1_) | (d_1436_v1_)
            nw299_[int(15)] = d_1436_v1_
            nw299_[int(16)] = (d_1436_v1_ if (d_1582_v105_)[default__.safeIndex(699, len(d_1582_v105_))] else (d_1436_v1_).set(d_1435_v0_, d_1470_v29_))
            nw299_[int(17)] = ((d_1436_v1_).set(True, d_1470_v29_)) | (d_1436_v1_)
            nw299_[int(18)] = d_1436_v1_
            nw299_[int(19)] = d_1436_v1_
            nw299_[int(20)] = d_1436_v1_
            nw299_[int(21)] = d_1436_v1_
            nw299_[int(22)] = d_1436_v1_
            nw299_[int(23)] = _dafny.Map({d_1435_v0_: len(d_1583_v106_)})
            nw299_[int(24)] = d_1436_v1_
            nw299_[int(25)] = (d_1436_v1_) | (default__.fm36(d_1438_v3_, d_1581_v104_, globalState))
            nw299_[int(26)] = d_1436_v1_
            nw299_[int(27)] = d_1436_v1_
            nw299_[int(28)] = d_1436_v1_
            d_1584_v107_ = nw299_
            d_1585_v108_: _dafny.Map
            d_1585_v108_ = _dafny.Map({d_1470_v29_: d_1436_v1_})
            index274_ = default__.safeIndex(326, (d_1584_v107_).length(0))
            (d_1584_v107_)[index274_] = ((d_1585_v108_)[d_1470_v29_] if (d_1470_v29_) in (d_1585_v108_) else d_1436_v1_)
            d_1586_v109_: _dafny.Map
            d_1586_v109_ = _dafny.Map({d_1574_v101_: d_1437_v2_})
            index275_ = default__.safeIndex(317, (d_1574_v101_).length(0))
            index276_ = default__.safeIndex(372, (d_1579_v102_).length(0))
            index277_ = default__.safeIndex(326, (d_1584_v107_).length(0))
            rhs217_ = (d_1435_v0_ if d_1435_v0_ else (len(((d_1586_v109_)[d_1574_v101_] if (d_1574_v101_) in (d_1586_v109_) else d_1437_v2_))) >= (default__.fm1(d_1435_v0_, globalState)))
            rhs218_ = d_1470_v29_
            rhs219_ = d_1436_v1_
            rhs220_ = default__.safeDivisionInt(d_1470_v29_, d_1470_v29_)
            rhs221_ = d_1435_v0_
            lhs190_ = d_1574_v101_
            lhs191_ = default__.safeIndex(317, (d_1574_v101_).length(0))
            lhs192_ = d_1579_v102_
            lhs193_ = default__.safeIndex(372, (d_1579_v102_).length(0))
            lhs194_ = d_1584_v107_
            lhs195_ = default__.safeIndex(326, (d_1584_v107_).length(0))
            lhs196_ = globalState
            lhs197_ = globalState
            lhs190_[lhs191_] = rhs217_
            lhs192_[lhs193_] = rhs218_
            lhs194_[lhs195_] = rhs219_
            lhs196_.f6 = rhs220_
            lhs197_.f1 = rhs221_
            d_1587_v110_: _dafny.Array
            nw300_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_1587_v110_ = nw300_
            d_1588_v111_: _dafny.Array
            def lambda73_(d_1589_i14_):
                return _dafny.MultiSet([True])

            init44_ = lambda73_
            nw301_ = _dafny.Array(None, 23)
            for i0_44_ in range(nw301_.length(0)):
                nw301_[i0_44_] = init44_(i0_44_)
            d_1588_v111_ = nw301_
            d_1590_v112_: _dafny.Seq
            d_1590_v112_ = _dafny.SeqWithoutIsStrInference([d_1588_v111_])
            index278_ = default__.safeIndex(703, (d_1587_v110_).length(0))
            (d_1587_v110_)[index278_] = (d_1590_v112_)[default__.safeIndex(d_1470_v29_, len(d_1590_v112_))]
            d_1591_v113_: _dafny.Map
            d_1591_v113_ = _dafny.Map({(d_1579_v102_)[default__.safeIndex(372, (d_1579_v102_).length(0))]: d_1435_v0_})
            index279_ = default__.safeIndex(372, (d_1579_v102_).length(0))
            index280_ = default__.safeIndex(703, (d_1587_v110_).length(0))
            rhs222_ = default__.fm1(False, globalState)
            rhs223_ = d_1435_v0_
            rhs224_ = (90 if d_1435_v0_ else (d_1579_v102_)[default__.safeIndex(372, (d_1579_v102_).length(0))])
            rhs225_ = (d_1435_v0_) and ((432) not in (d_1591_v113_))
            rhs226_ = d_1588_v111_
            lhs198_ = globalState
            lhs199_ = d_1579_v102_
            lhs200_ = default__.safeIndex(372, (d_1579_v102_).length(0))
            lhs201_ = globalState
            lhs202_ = d_1587_v110_
            lhs203_ = default__.safeIndex(703, (d_1587_v110_).length(0))
            lhs198_.f6 = rhs222_
            d_1435_v0_ = rhs223_
            lhs199_[lhs200_] = rhs224_
            lhs201_.f1 = rhs225_
            lhs202_[lhs203_] = rhs226_
            d_1437_v2_ = d_1437_v2_
            d_1592_v114_: _dafny.MultiSet
            d_1592_v114_ = _dafny.MultiSet([(d_1574_v101_)[default__.safeIndex(317, (d_1574_v101_).length(0))]])
            d_1593_v115_: _dafny.Seq
            d_1593_v115_ = _dafny.SeqWithoutIsStrInference([d_1591_v113_])
            d_1594_v116_: _dafny.Map
            d_1594_v116_ = _dafny.Map({(((d_1593_v115_)[default__.safeIndex(d_1470_v29_, len(d_1593_v115_))]).set(d_1470_v29_, (d_1574_v101_)[default__.safeIndex(317, (d_1574_v101_).length(0))])): d_1470_v29_})
            d_1595_v117_: _dafny.Map
            d_1595_v117_ = d_1594_v116_
            d_1596_v118_: C2
            nw302_ = C2()
            nw302_.ctor__(d_1470_v29_, (0) - ((((d_1592_v114_).set((d_1574_v101_)[default__.safeIndex(317, (d_1574_v101_).length(0))], default__.abs(len(d_1591_v113_)))) - (d_1592_v114_)).cardinality), ((d_1595_v117_)) | (d_1594_v116_))
            d_1596_v118_ = nw302_
            d_1597_v120_: C3
            nw303_ = C3()
            def iife185_():
                coll63_ = _dafny.Map()
                compr_63_: _dafny.Map
                for compr_63_ in (d_1594_v116_).keys.Elements:
                    d_1598_v119_: _dafny.Map = compr_63_
                    if (d_1598_v119_) in (d_1594_v116_):
                        coll63_[d_1598_v119_] = len(d_1583_v106_)
                return _dafny.Map(coll63_)
            nw303_.ctor__((d_1574_v101_)[default__.safeIndex(317, (d_1574_v101_).length(0))], (0) - ((0) - ((159) + ((d_1596_v118_).f20))), iife185_()
            )
            d_1597_v120_ = nw303_
        (globalState).f1 = default__.fm2(globalState)


class C5(T0):
    def  __init__(self):
        self._f16: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f16):
        (self)._f16 = f16

    def fm3(self, p0, p1, p2, p3, globalState):
        return D0_DC2(D0_DC0(False))

    def fm11(self, p0, globalState):
        return (_dafny.MultiSet([True])) | (_dafny.MultiSet([True]))

    def m1(self, globalState):
        r0: int = int(0)
        r1: bool = False
        d_1599_v0_: bool
        d_1599_v0_ = False
        d_1600_v1_: _dafny.Seq
        d_1600_v1_ = _dafny.SeqWithoutIsStrInference([d_1599_v0_, d_1599_v0_])
        d_1601_v2_: _dafny.Array
        nw304_ = _dafny.Array(None, 12)
        nw304_[int(0)] = d_1599_v0_
        nw304_[int(1)] = not(d_1599_v0_)
        nw304_[int(2)] = d_1599_v0_
        nw304_[int(3)] = d_1599_v0_
        nw304_[int(4)] = False
        nw304_[int(5)] = d_1599_v0_
        nw304_[int(6)] = d_1599_v0_
        nw304_[int(7)] = d_1599_v0_
        nw304_[int(8)] = d_1599_v0_
        nw304_[int(9)] = d_1599_v0_
        nw304_[int(10)] = d_1599_v0_
        nw304_[int(11)] = d_1599_v0_
        d_1601_v2_ = nw304_
        d_1602_v3_: _dafny.Seq
        d_1602_v3_ = _dafny.SeqWithoutIsStrInference([d_1601_v2_, d_1601_v2_])
        d_1603_v4_: _dafny.Seq
        d_1603_v4_ = _dafny.SeqWithoutIsStrInference([len(d_1602_v3_)])
        d_1604_v5_: _dafny.Array
        nw305_ = _dafny.Array(None, 11)
        nw305_[int(0)] = d_1599_v0_
        nw305_[int(1)] = d_1599_v0_
        nw305_[int(2)] = True
        nw305_[int(3)] = d_1599_v0_
        nw305_[int(4)] = d_1599_v0_
        nw305_[int(5)] = (d_1600_v1_)[default__.safeIndex((self).f16, len(d_1600_v1_))]
        nw305_[int(6)] = ((0) - ((0) - ((d_1603_v4_)[default__.safeIndex((self).f16, len(d_1603_v4_))]))) < ((self).f16)
        nw305_[int(7)] = (not(d_1599_v0_) if d_1599_v0_ else d_1599_v0_)
        nw305_[int(8)] = d_1599_v0_
        nw305_[int(9)] = default__.fm2(globalState)
        nw305_[int(10)] = d_1599_v0_
        d_1604_v5_ = nw305_
        d_1605_v6_: str
        d_1605_v6_ = _dafny.CodePoint('n')
        d_1606_v7_: _dafny.Seq
        d_1606_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
        index281_ = default__.safeIndex(854, (d_1601_v2_).length(0))
        (d_1601_v2_)[index281_] = (d_1605_v6_) in (d_1606_v7_)
        index282_ = default__.safeIndex(854, (d_1601_v2_).length(0))
        (d_1601_v2_)[index282_] = d_1599_v0_
        d_1607_i0_: int
        d_1607_i0_ = 0
        with _dafny.label("12"):
            while (False if d_1599_v0_ else ((d_1603_v4_)[default__.safeIndex((self).f16, len(d_1603_v4_))]) > (len(d_1606_v7_))):
                with _dafny.c_label("12"):
                    if (d_1607_i0_) >= (100):
                        raise _dafny.Break("12")
                    d_1607_i0_ = (d_1607_i0_) + (1)
                    d_1608_v8_: _dafny.Array
                    nw306_ = _dafny.Array(_dafny.Seq({}), 18)
                    d_1608_v8_ = nw306_
                    d_1609_v9_: _dafny.Map
                    d_1609_v9_ = _dafny.Map({default__.safeDivisionInt((self).f16, (self).f16): d_1608_v8_})
                    d_1609_v9_ = (d_1609_v9_).set(((self).f16 if (d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))] else (self).f16), d_1608_v8_)
                    d_1610_v10_: _dafny.Set
                    d_1610_v10_ = _dafny.Set({(self).f16})
                    d_1611_v11_: _dafny.Array
                    def lambda74_(d_1612_v2_):
                        def lambda75_(d_1613_i1_):
                            return D0_DC2(D0_DC2(D0_DC0((d_1612_v2_)[default__.safeIndex(854, (d_1612_v2_).length(0))])))

                        return lambda75_

                    init45_ = lambda74_(d_1601_v2_)
                    nw307_ = _dafny.Array(None, 26)
                    for i0_45_ in range(nw307_.length(0)):
                        nw307_[i0_45_] = init45_(i0_45_)
                    d_1611_v11_ = nw307_
                    d_1614_v12_: _dafny.Array
                    nw308_ = _dafny.Array(int(0), 24)
                    d_1614_v12_ = nw308_
                    default__.m0((d_1610_v10_ if (d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))] else d_1610_v10_), d_1611_v11_, d_1614_v12_, d_1606_v7_, globalState)
                    d_1615_v13_: _dafny.MultiSet
                    d_1615_v13_ = _dafny.MultiSet([d_1599_v0_, False])
                    index283_ = default__.safeIndex(908, (d_1614_v12_).length(0))
                    (d_1614_v12_)[index283_] = default__.safeDivisionInt(len(default__.fm12(d_1615_v13_, (self).f16, globalState)), (self).f16)
                    index284_ = default__.safeIndex(908, (d_1614_v12_).length(0))
                    (d_1614_v12_)[index284_] = (self).f16
                    index285_ = default__.safeIndex(854, (d_1601_v2_).length(0))
                    (d_1601_v2_)[index285_] = not(not((d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]))
                    pass
            pass
        d_1616_v14_: _dafny.Array
        nw309_ = _dafny.Array(None, 14)
        nw309_[int(0)] = d_1600_v1_
        nw309_[int(1)] = d_1600_v1_
        nw309_[int(2)] = default__.fm13(d_1599_v0_, d_1605_v6_, (self).f16, globalState)
        nw309_[int(3)] = d_1600_v1_
        nw309_[int(4)] = d_1600_v1_
        nw309_[int(5)] = d_1600_v1_
        nw309_[int(6)] = d_1600_v1_
        nw309_[int(7)] = d_1600_v1_
        nw309_[int(8)] = d_1600_v1_
        nw309_[int(9)] = d_1600_v1_
        nw309_[int(10)] = d_1600_v1_
        nw309_[int(11)] = d_1600_v1_
        nw309_[int(12)] = d_1600_v1_
        nw309_[int(13)] = (d_1600_v1_) + (d_1600_v1_)
        d_1616_v14_ = nw309_
        index286_ = default__.safeIndex(43, (d_1616_v14_).length(0))
        (d_1616_v14_)[index286_] = d_1600_v1_
        index287_ = default__.safeIndex(43, (d_1616_v14_).length(0))
        (d_1616_v14_)[index287_] = default__.fm13(((self).f16) > ((self).f16), d_1605_v6_, len(d_1606_v7_), globalState)
        d_1617_v15_: _dafny.MultiSet
        d_1617_v15_ = _dafny.MultiSet([False])
        if (d_1617_v15_).ispropersubset((self).fm11(103, globalState)):
            d_1618_v16_: D0
            d_1618_v16_ = D0_DC0((d_1600_v1_)[default__.safeIndex((self).f16, len(d_1600_v1_))])
            d_1619_v17_: _dafny.MultiSet
            d_1619_v17_ = _dafny.MultiSet([d_1618_v16_, d_1618_v16_])
            d_1620_v18_: _dafny.Seq
            d_1620_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1599_v0_, False])])
            d_1621_v19_: _dafny.Map
            d_1621_v19_ = _dafny.Map({((d_1619_v17_)[d_1618_v16_] if (d_1618_v16_) in (d_1619_v17_) else (0) - ((self).f16)): d_1620_v18_})
            d_1621_v19_ = (d_1621_v19_).set((self).f16, d_1620_v18_)
            d_1622_v20_: _dafny.Seq
            d_1622_v20_ = d_1603_v4_
            d_1603_v4_ = ((d_1603_v4_) + (_dafny.SeqWithoutIsStrInference([(self).f16 for d_1623_i2_ in range(default__.abs(-332))]))) + ((d_1622_v20_))
            d_1624_v21_: _dafny.Map
            d_1624_v21_ = _dafny.Map({len(d_1606_v7_): (self).f16})
            (globalState).f1 = (d_1624_v21_) != (d_1624_v21_)
            index288_ = default__.safeIndex(854, (d_1601_v2_).length(0))
            (d_1601_v2_)[index288_] = (d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]
            d_1625_v22_: _dafny.Array
            nw310_ = _dafny.Array(None, 28)
            nw310_[int(0)] = (self).f16
            nw310_[int(1)] = (self).f16
            nw310_[int(2)] = (self).f16
            nw310_[int(3)] = (self).f16
            nw310_[int(4)] = (self).f16
            nw310_[int(5)] = (self).f16
            nw310_[int(6)] = (self).f16
            nw310_[int(7)] = 840
            nw310_[int(8)] = (self).f16
            nw310_[int(9)] = -130
            nw310_[int(10)] = (self).f16
            nw310_[int(11)] = len(_dafny.Map({251: (self).f16}))
            nw310_[int(12)] = (self).f16
            nw310_[int(13)] = (self).f16
            nw310_[int(14)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvp")))
            nw310_[int(15)] = (self).f16
            nw310_[int(16)] = (self).f16
            nw310_[int(17)] = (self).f16
            nw310_[int(18)] = (self).f16
            nw310_[int(19)] = (self).f16
            nw310_[int(20)] = (self).f16
            nw310_[int(21)] = (self).f16
            nw310_[int(22)] = (self).f16
            nw310_[int(23)] = (self).f16
            nw310_[int(24)] = (0) - ((self).f16)
            nw310_[int(25)] = (0) - ((self).f16)
            nw310_[int(26)] = (self).f16
            nw310_[int(27)] = len(d_1606_v7_)
            d_1625_v22_ = nw310_
            d_1626_v23_: D3
            d_1626_v23_ = D3_DC6(d_1625_v22_)
            d_1627_v24_: _dafny.Set
            d_1627_v24_ = _dafny.Set({(D3_DC6(d_1625_v22_)).cf6, (d_1626_v23_).cf6, d_1625_v22_, d_1625_v22_})
            d_1627_v24_ = (_dafny.Set({d_1625_v22_})) | ((d_1627_v24_) - (d_1627_v24_))
        elif True:
            d_1628_v25_: D0
            d_1628_v25_ = D0_DC0((d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))])
            source12_ = d_1628_v25_
            if source12_.is_DC1:
                d_1629___mcc_h0_ = source12_.cf1
                d_1630_cf1_ = d_1629___mcc_h0_
                d_1631_v26_: _dafny.Set
                d_1631_v26_ = _dafny.Set({d_1630_cf1_})
                d_1632_v27_: _dafny.Set
                d_1632_v27_ = _dafny.Set({len(d_1631_v26_), (self).f16})
                d_1633_v28_: _dafny.Array
                def lambda76_(d_1634_i3_):
                    return D0_DC2(D0_DC1((D0_DC1(778)).cf1))

                init46_ = lambda76_
                nw311_ = _dafny.Array(None, 19)
                for i0_46_ in range(nw311_.length(0)):
                    nw311_[i0_46_] = init46_(i0_46_)
                d_1633_v28_ = nw311_
                d_1635_v29_: _dafny.Array
                nw312_ = _dafny.Array(None, 6)
                nw312_[int(0)] = (self).f16
                nw312_[int(1)] = (self).f16
                nw312_[int(2)] = 450
                nw312_[int(3)] = len(d_1632_v27_)
                nw312_[int(4)] = (self).f16
                nw312_[int(5)] = (self).f16
                d_1635_v29_ = nw312_
                default__.m0(d_1631_v26_, d_1633_v28_, d_1635_v29_, (_dafny.SeqWithoutIsStrInference([d_1605_v6_ for d_1636_i4_ in range(default__.abs(214))])) + (d_1606_v7_), globalState)
                index289_ = default__.safeIndex(854, (d_1601_v2_).length(0))
                (d_1601_v2_)[index289_] = d_1599_v0_
                d_1637_v30_: _dafny.MultiSet
                d_1637_v30_ = _dafny.MultiSet([d_1630_cf1_])
                d_1638_v31_: D1
                d_1638_v31_ = D1_DC4(d_1605_v6_)
                d_1639_v32_: _dafny.Array
                nw313_ = _dafny.Array(None, 10)
                nw313_[int(0)] = d_1605_v6_
                nw313_[int(1)] = d_1605_v6_
                nw313_[int(2)] = (d_1606_v7_)[default__.safeIndex((d_1637_v30_).cardinality, len(d_1606_v7_))]
                nw313_[int(3)] = d_1605_v6_
                nw313_[int(4)] = (d_1638_v31_).cf4
                nw313_[int(5)] = _dafny.CodePoint('i')
                nw313_[int(6)] = d_1605_v6_
                nw313_[int(7)] = d_1605_v6_
                nw313_[int(8)] = d_1605_v6_
                nw313_[int(9)] = d_1605_v6_
                d_1639_v32_ = nw313_
                index290_ = default__.safeIndex(655, (d_1639_v32_).length(0))
                (d_1639_v32_)[index290_] = d_1605_v6_
                index291_ = default__.safeIndex(655, (d_1639_v32_).length(0))
                (d_1639_v32_)[index291_] = d_1605_v6_
                d_1640_v33_: _dafny.Map
                d_1640_v33_ = _dafny.Map({d_1630_cf1_: d_1606_v7_})
                d_1641_v34_: _dafny.MultiSet
                d_1641_v34_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mh")), ((d_1640_v33_)[d_1630_cf1_] if (d_1630_cf1_) in (d_1640_v33_) else _dafny.SeqWithoutIsStrInference([d_1605_v6_ for d_1642_i5_ in range(default__.abs(221))]))])
                d_1643_v35_: _dafny.Map
                d_1643_v35_ = _dafny.Map({(d_1641_v34_).cardinality: (self).f16})
                (globalState).f6 = default__.safeModuloInt(len(_dafny.Set({(d_1617_v15_).cardinality})), ((d_1643_v35_)[d_1630_cf1_] if (d_1630_cf1_) in (d_1643_v35_) else d_1630_cf1_))
            elif source12_.is_DC0:
                d_1644___mcc_h1_ = source12_.cf0
                d_1645_cf0_ = d_1644___mcc_h1_
                d_1646_v36_: _dafny.Map
                d_1646_v36_ = _dafny.Map({d_1599_v0_: (_dafny.MultiSet([(self).f16, (self).f16])).cardinality})
                d_1647_v37_: _dafny.Map
                d_1647_v37_ = d_1646_v36_
                d_1648_v38_: _dafny.Map
                d_1648_v38_ = _dafny.Map({((d_1647_v37_)) | ((d_1647_v37_)): d_1599_v0_})
                d_1649_v39_: _dafny.Seq
                d_1649_v39_ = d_1603_v4_
                d_1650_v40_: _dafny.Map
                d_1650_v40_ = _dafny.Map({default__.fm1((d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))], globalState): d_1649_v39_})
                d_1648_v38_ = (d_1648_v38_).set(default__.fm14(d_1650_v40_, (self).f16, d_1645_cf0_, globalState), d_1599_v0_)
                (globalState).f1 = (default__.fm1(d_1645_cf0_, globalState)) == (len(d_1606_v7_))
                d_1651_v42_: _dafny.Map
                d_1651_v42_ = _dafny.Map({776: (self).f16})
                d_1652_v43_: D12
                d_1652_v43_ = D12_DC34(d_1651_v42_)
                d_1653_v44_: C3
                nw314_ = C3()
                def iife186_():
                    coll64_ = _dafny.Map()
                    compr_64_: int
                    for compr_64_ in ((d_1652_v43_).cf63).keys.Elements:
                        d_1654_v41_: int = compr_64_
                        if (d_1654_v41_) in ((d_1652_v43_).cf63):
                            coll64_[default__.safeDivisionInt(d_1654_v41_, (self).f16)] = d_1599_v0_
                    return _dafny.Map(coll64_)
                nw314_.ctor__(d_1599_v0_, default__.safeModuloInt((self).f16, 851), _dafny.Map({iife186_()
                : (self).f16}))
                d_1653_v44_ = nw314_
                d_1655_v45_: D3
                d_1655_v45_ = D3_DC7((self).f16)
                d_1656_v46_: _dafny.Map
                d_1656_v46_ = _dafny.Map({d_1601_v2_: (d_1653_v44_).f18})
                d_1657_v47_: D10
                d_1657_v47_ = D10_DC27((self).f16, d_1606_v7_, -202, d_1655_v45_, d_1656_v46_)
                d_1658_v48_: D10
                d_1658_v48_ = D10_DC29(d_1657_v47_)
                d_1659_v49_: D10
                d_1659_v49_ = D10_DC29(d_1658_v48_)
                d_1660_v50_: D10
                d_1660_v50_ = D10_DC29(d_1659_v49_)
                d_1661_v51_: D10
                d_1661_v51_ = D10_DC29(d_1659_v49_)
                d_1662_v52_: _dafny.MultiSet
                d_1662_v52_ = _dafny.MultiSet([d_1661_v51_, d_1661_v51_, d_1661_v51_])
                d_1662_v52_ = d_1662_v52_
            elif True:
                d_1663___mcc_h2_ = source12_.cf2
                d_1664_cf2_ = d_1663___mcc_h2_
                (globalState).f6 = (self).f16
                r0 = len((d_1616_v14_)[default__.safeIndex(43, (d_1616_v14_).length(0))])
                index292_ = default__.safeIndex(854, (d_1601_v2_).length(0))
                (d_1601_v2_)[index292_] = d_1599_v0_
                (globalState).f6 = (self).f16
            index293_ = default__.safeIndex(854, (d_1601_v2_).length(0))
            rhs227_ = (0) - ((0) - ((self).f16))
            rhs228_ = (((self).f16 if d_1599_v0_ else 799)) < (705)
            rhs229_ = default__.fm1((d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))], globalState)
            rhs230_ = d_1599_v0_
            lhs204_ = globalState
            lhs205_ = globalState
            lhs206_ = globalState
            lhs207_ = d_1601_v2_
            lhs208_ = default__.safeIndex(854, (d_1601_v2_).length(0))
            lhs204_.f6 = rhs227_
            lhs205_.f1 = rhs228_
            lhs206_.f6 = rhs229_
            lhs207_[lhs208_] = rhs230_
            d_1665_v53_: _dafny.MultiSet
            d_1665_v53_ = _dafny.MultiSet([_dafny.CodePoint('k'), d_1605_v6_])
            d_1666_v54_: _dafny.MultiSet
            d_1666_v54_ = _dafny.MultiSet([(self).f16, (d_1665_v53_).cardinality, (self).f16])
            d_1667_v55_: _dafny.Set
            d_1667_v55_ = _dafny.Set({(d_1666_v54_).cardinality})
            d_1668_v56_: _dafny.Map
            d_1668_v56_ = _dafny.Map({(self).f16: d_1667_v55_})
            d_1669_v57_: _dafny.Map
            d_1669_v57_ = _dafny.Map({(d_1603_v4_)[default__.safeIndex((self).f16, len(d_1603_v4_))]: d_1668_v56_})
            d_1670_v58_: D9
            d_1670_v58_ = D9_DC24(d_1669_v57_, d_1600_v1_)
            source13_ = d_1670_v58_
            if source13_.is_DC23:
                d_1671___mcc_h3_ = source13_.cf37
                d_1672___mcc_h4_ = source13_.cf38
                d_1673___mcc_h5_ = source13_.cf39
                d_1674___mcc_h6_ = source13_.cf40
                d_1675___mcc_h7_ = source13_.cf41
                d_1676_cf41_ = d_1675___mcc_h7_
                d_1677_cf40_ = d_1674___mcc_h6_
                d_1678_cf39_ = d_1673___mcc_h5_
                d_1679_cf38_ = d_1672___mcc_h4_
                d_1680_cf37_ = d_1671___mcc_h3_
                d_1681_v59_: D16
                d_1681_v59_ = D16_DC44((_dafny.MultiSet([False, d_1678_cf39_])) - (d_1617_v15_))
                rhs231_ = (self).f16
                rhs232_ = ((d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))] if d_1599_v0_ else d_1599_v0_)
                rhs233_ = d_1681_v59_
                lhs209_ = globalState
                lhs209_.f6 = rhs231_
                d_1678_cf39_ = rhs232_
                d_1681_v59_ = rhs233_
                d_1682_v60_: _dafny.Map
                d_1682_v60_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f16 for d_1683_i6_ in range(default__.abs(579))]): (self).f16})
                d_1684_v61_: _dafny.Array
                nw315_ = _dafny.Array(None, 6)
                nw315_[int(0)] = d_1682_v60_
                nw315_[int(1)] = d_1682_v60_
                nw315_[int(2)] = d_1682_v60_
                nw315_[int(3)] = d_1682_v60_
                nw315_[int(4)] = (d_1682_v60_).set(d_1603_v4_, (self).f16)
                nw315_[int(5)] = (d_1682_v60_).set(_dafny.SeqWithoutIsStrInference([(self).f16]), (self).f16)
                d_1684_v61_ = nw315_
                index294_ = default__.safeIndex(218, (d_1684_v61_).length(0))
                (d_1684_v61_)[index294_] = d_1682_v60_
                d_1685_v62_: _dafny.Set
                d_1685_v62_ = _dafny.Set({d_1677_cf40_})
                d_1686_v63_: _dafny.Set
                d_1686_v63_ = _dafny.Set({d_1599_v0_})
                index295_ = default__.safeIndex(218, (d_1684_v61_).length(0))
                rhs234_ = (self).f16
                rhs235_ = (d_1685_v62_).issubset(d_1685_v62_)
                rhs236_ = (d_1682_v60_) | ((d_1682_v60_) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f16]): 118})))
                rhs237_ = len(d_1603_v4_)
                rhs238_ = len(d_1686_v63_)
                lhs210_ = globalState
                lhs211_ = d_1684_v61_
                lhs212_ = default__.safeIndex(218, (d_1684_v61_).length(0))
                lhs213_ = globalState
                lhs214_ = globalState
                lhs210_.f6 = rhs234_
                d_1678_cf39_ = rhs235_
                lhs211_[lhs212_] = rhs236_
                lhs213_.f6 = rhs237_
                lhs214_.f6 = rhs238_
                d_1687_v64_: _dafny.Map
                d_1687_v64_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1676_cf41_ for d_1688_i7_ in range(default__.abs(48))]): (((d_1616_v14_)[default__.safeIndex(43, (d_1616_v14_).length(0))])[default__.safeIndex((self).f16, len((d_1616_v14_)[default__.safeIndex(43, (d_1616_v14_).length(0))]))]) and (d_1678_cf39_)})
                d_1689_v65_: _dafny.Map
                d_1689_v65_ = _dafny.Map({len(d_1679_cf38_): (d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]})
                index296_ = default__.safeIndex(854, (d_1601_v2_).length(0))
                rhs239_ = (self).f16
                rhs240_ = (d_1687_v64_) | (d_1687_v64_)
                rhs241_ = not(((d_1616_v14_)[default__.safeIndex(43, (d_1616_v14_).length(0))])[default__.safeIndex((len(d_1689_v65_)) + ((self).f16), len((d_1616_v14_)[default__.safeIndex(43, (d_1616_v14_).length(0))]))])
                lhs215_ = d_1601_v2_
                lhs216_ = default__.safeIndex(854, (d_1601_v2_).length(0))
                r0 = rhs239_
                d_1687_v64_ = rhs240_
                lhs215_[lhs216_] = rhs241_
                (globalState).f6 = (self).f16
            elif source13_.is_DC24:
                d_1690___mcc_h8_ = source13_.cf42
                d_1691___mcc_h9_ = source13_.cf43
                d_1692_cf43_ = d_1691___mcc_h9_
                d_1693_cf42_ = d_1690___mcc_h8_
                d_1606_v7_ = d_1606_v7_
                d_1694_v66_: D10
                d_1694_v66_ = D10_DC28((self).f16)
                d_1694_v66_ = D10_DC28((self).f16)
                index297_ = default__.safeIndex(854, (d_1601_v2_).length(0))
                (d_1601_v2_)[index297_] = default__.fm2(globalState)
                d_1695_v67_: _dafny.MultiSet
                d_1695_v67_ = _dafny.MultiSet([d_1600_v1_, d_1692_cf43_, _dafny.SeqWithoutIsStrInference([(d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))], d_1599_v0_]), _dafny.SeqWithoutIsStrInference([default__.fm2(globalState)])])
                d_1696_v68_: _dafny.Map
                d_1696_v68_ = _dafny.Map({(self).f16: (d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]})
                d_1697_v69_: _dafny.Map
                d_1697_v69_ = _dafny.Map({d_1696_v68_: (self).f16})
                d_1698_v70_: C2
                nw316_ = C2()
                nw316_.ctor__(((d_1695_v67_)[_dafny.SeqWithoutIsStrInference([(d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))], (d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))], d_1599_v0_])] if (_dafny.SeqWithoutIsStrInference([(d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))], (d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))], d_1599_v0_])) in (d_1695_v67_) else (self).f16), (self).f16, d_1697_v69_)
                d_1698_v70_ = nw316_
            elif source13_.is_DC22:
                d_1699___mcc_h10_ = source13_.cf36
                d_1700_cf36_ = d_1699___mcc_h10_
                d_1701_v71_: _dafny.Map
                d_1701_v71_ = _dafny.Map({d_1599_v0_: _dafny.SeqWithoutIsStrInference([d_1667_v55_ for d_1702_i8_ in range(default__.abs(553))])})
                d_1703_v72_: _dafny.Seq
                d_1703_v72_ = _dafny.SeqWithoutIsStrInference([d_1667_v55_])
                d_1704_v73_: _dafny.Map
                d_1704_v73_ = _dafny.Map({480: d_1599_v0_})
                d_1705_v74_: _dafny.Map
                d_1705_v74_ = _dafny.Map({d_1704_v73_: (self).f16})
                d_1706_v75_: C2
                nw317_ = C2()
                nw317_.ctor__((self).f16, ((self).f16) - (len(((d_1701_v71_)[(d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]] if ((d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]) in (d_1701_v71_) else d_1703_v72_))), (d_1705_v74_) | (d_1705_v74_))
                d_1706_v75_ = nw317_
                d_1707_v76_: _dafny.Set
                d_1707_v76_ = _dafny.Set({d_1599_v0_, (d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]})
                d_1708_v77_: D11
                d_1708_v77_ = D11_DC32(not ((d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]) or (d_1599_v0_), (0) - ((d_1706_v75_).f20), ((d_1666_v54_).set(len(d_1707_v76_), default__.abs(d_1706_v75_.f21))).ispropersubset(d_1666_v54_))
                d_1708_v77_ = ((d_1708_v77_ if (d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))] else d_1708_v77_) if default__.fm2(globalState) else d_1708_v77_)
                d_1709_v79_: _dafny.Array
                def lambda77_(d_1710_v6_, d_1711_v7_, d_1712_v0_):
                    def lambda78_(d_1713_i9_):
                        def iife187_():
                            coll65_ = _dafny.Map()
                            compr_65_: D1
                            for compr_65_ in (_dafny.Map({D1_DC4(d_1710_v6_): d_1711_v7_})).keys.Elements:
                                d_1714_v78_: D1 = compr_65_
                                if (d_1714_v78_) in (_dafny.Map({D1_DC4(d_1710_v6_): d_1711_v7_})):
                                    coll65_[d_1714_v78_] = d_1712_v0_
                            return _dafny.Map(coll65_)
                        return iife187_()
                        

                    return lambda78_

                init47_ = lambda77_(d_1605_v6_, d_1606_v7_, d_1599_v0_)
                nw318_ = _dafny.Array(None, 27)
                for i0_47_ in range(nw318_.length(0)):
                    nw318_[i0_47_] = init47_(i0_47_)
                d_1709_v79_ = nw318_
                d_1715_v80_: D20
                d_1715_v80_ = D20_DC49(d_1709_v79_)
                d_1716_v81_: _dafny.Array
                nw319_ = _dafny.Array(None, 26)
                nw319_[int(0)] = d_1709_v79_
                nw319_[int(1)] = d_1709_v79_
                nw319_[int(2)] = d_1709_v79_
                nw319_[int(3)] = d_1709_v79_
                nw319_[int(4)] = d_1709_v79_
                nw319_[int(5)] = d_1709_v79_
                nw319_[int(6)] = d_1709_v79_
                nw319_[int(7)] = d_1709_v79_
                nw319_[int(8)] = d_1709_v79_
                nw319_[int(9)] = d_1709_v79_
                nw319_[int(10)] = d_1709_v79_
                nw319_[int(11)] = d_1709_v79_
                nw319_[int(12)] = d_1709_v79_
                nw319_[int(13)] = d_1709_v79_
                nw319_[int(14)] = (d_1715_v80_).cf85
                nw319_[int(15)] = d_1709_v79_
                nw319_[int(16)] = d_1709_v79_
                nw319_[int(17)] = d_1709_v79_
                nw319_[int(18)] = d_1709_v79_
                nw319_[int(19)] = d_1709_v79_
                nw319_[int(20)] = d_1709_v79_
                nw319_[int(21)] = d_1709_v79_
                nw319_[int(22)] = d_1709_v79_
                nw319_[int(23)] = d_1709_v79_
                nw319_[int(24)] = d_1709_v79_
                nw319_[int(25)] = d_1709_v79_
                d_1716_v81_ = nw319_
                index298_ = default__.safeIndex(154, (d_1716_v81_).length(0))
                (d_1716_v81_)[index298_] = (d_1709_v79_ if d_1599_v0_ else d_1709_v79_)
                index299_ = default__.safeIndex(154, (d_1716_v81_).length(0))
                rhs242_ = ((d_1706_v75_).f20) <= (d_1706_v75_.f21)
                rhs243_ = (d_1709_v79_ if not(d_1599_v0_) else d_1709_v79_)
                rhs244_ = d_1667_v55_
                lhs217_ = globalState
                lhs218_ = d_1716_v81_
                lhs219_ = default__.safeIndex(154, (d_1716_v81_).length(0))
                lhs217_.f1 = rhs242_
                lhs218_[lhs219_] = rhs243_
                d_1667_v55_ = rhs244_
                d_1717_v82_: _dafny.Array
                nw320_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
                d_1717_v82_ = nw320_
                d_1717_v82_ = d_1717_v82_
            elif True:
                d_1718___mcc_h11_ = source13_.cf44
                d_1719_cf44_ = d_1718___mcc_h11_
                d_1720_v83_: _dafny.Array
                def lambda79_(d_1721_v4_):
                    def lambda80_(d_1722_i10_):
                        return (d_1722_i10_) + (len((d_1721_v4_).set(default__.safeIndex((self).f16, len(d_1721_v4_)), (self).f16)))

                    return lambda80_

                init48_ = lambda79_(d_1603_v4_)
                nw321_ = _dafny.Array(None, 20)
                for i0_48_ in range(nw321_.length(0)):
                    nw321_[i0_48_] = init48_(i0_48_)
                d_1720_v83_ = nw321_
                index300_ = default__.safeIndex(615, (d_1720_v83_).length(0))
                (d_1720_v83_)[index300_] = (self).f16
                index301_ = default__.safeIndex(615, (d_1720_v83_).length(0))
                (d_1720_v83_)[index301_] = (self).f16
                d_1723_v84_: _dafny.Map
                d_1723_v84_ = _dafny.Map({(d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]: d_1720_v83_})
                (globalState).f6 = ((d_1720_v83_)[default__.safeIndex(615, (d_1720_v83_).length(0))]) - (len(d_1723_v84_))
                d_1724_v85_: _dafny.MultiSet
                d_1724_v85_ = _dafny.MultiSet([d_1617_v15_, d_1617_v15_])
                d_1724_v85_ = (d_1724_v85_ if d_1599_v0_ else d_1724_v85_)
                d_1725_v86_: _dafny.Map
                d_1725_v86_ = _dafny.Map({True: d_1599_v0_})
                index302_ = default__.safeIndex(854, (d_1601_v2_).length(0))
                rhs245_ = False
                rhs246_ = (((d_1725_v86_)[(d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]] if ((d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]) in (d_1725_v86_) else (d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]) if (d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))] else (d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))])
                rhs247_ = (self).f16
                lhs220_ = globalState
                lhs221_ = d_1601_v2_
                lhs222_ = default__.safeIndex(854, (d_1601_v2_).length(0))
                lhs223_ = globalState
                lhs220_.f1 = rhs245_
                lhs221_[lhs222_] = rhs246_
                lhs223_.f6 = rhs247_
            d_1726_v87_: _dafny.Set
            d_1726_v87_ = _dafny.Set({d_1599_v0_})
            d_1727_v88_: D8
            d_1727_v88_ = D8_DC19(d_1726_v87_)
            d_1727_v88_ = d_1727_v88_
            if ((d_1603_v4_)[default__.safeIndex((self).f16, len(d_1603_v4_))]) < ((self).f16):
                r1 = True
                d_1728_v89_: _dafny.Map
                d_1728_v89_ = _dafny.Map({d_1605_v6_: d_1606_v7_})
                d_1729_v90_: D12
                d_1729_v90_ = D12_DC35()
                d_1730_v91_: _dafny.Map
                d_1730_v91_ = _dafny.Map({d_1729_v90_: _dafny.Map({d_1605_v6_: d_1606_v7_})})
                d_1728_v89_ = ((d_1730_v91_)[d_1729_v90_] if (d_1729_v90_) in (d_1730_v91_) else d_1728_v89_)
                (globalState).f6 = (0) - ((self).f16)
                (globalState).f6 = (self).f16
                d_1731_v92_: _dafny.Array
                def lambda81_(d_1732_i11_):
                    return (d_1732_i11_) + (639)

                init49_ = lambda81_
                nw322_ = _dafny.Array(None, 28)
                for i0_49_ in range(nw322_.length(0)):
                    nw322_[i0_49_] = init49_(i0_49_)
                d_1731_v92_ = nw322_
                (globalState).f9 = d_1731_v92_
            elif True:
                (globalState).f6 = len(d_1726_v87_)
                d_1733_v93_: _dafny.Array
                nw323_ = _dafny.Array(_dafny.Map({}), 5)
                d_1733_v93_ = nw323_
                d_1734_v94_: _dafny.Array
                nw324_ = _dafny.Array(int(0), 22)
                d_1734_v94_ = nw324_
                d_1735_v95_: _dafny.Map
                d_1735_v95_ = _dafny.Map({d_1733_v93_: (d_1734_v94_ if (d_1600_v1_)[default__.safeIndex((self).f16, len(d_1600_v1_))] else d_1734_v94_)})
                d_1736_v96_: _dafny.Seq
                d_1736_v96_ = _dafny.SeqWithoutIsStrInference([d_1735_v95_])
                d_1737_v97_: _dafny.Map
                d_1737_v97_ = _dafny.Map({len(d_1603_v4_): (self).f16})
                d_1735_v95_ = (d_1736_v96_)[default__.safeIndex(((self).f16) - (len(d_1737_v97_)), len(d_1736_v96_))]
                d_1666_v54_ = (d_1666_v54_) - ((_dafny.MultiSet((d_1603_v4_).set(default__.safeIndex((self).f16, len(d_1603_v4_)), (self).f16))).set(52, default__.abs((self).f16)))
                index303_ = default__.safeIndex(854, (d_1601_v2_).length(0))
                (d_1601_v2_)[index303_] = d_1599_v0_
                index304_ = default__.safeIndex(906, (d_1734_v94_).length(0))
                (d_1734_v94_)[index304_] = (self).f16
                index305_ = default__.safeIndex(906, (d_1734_v94_).length(0))
                rhs248_ = d_1599_v0_
                rhs249_ = default__.safeDivisionInt((self).f16, (self).f16)
                rhs250_ = (749) < (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ayqx"))))
                rhs251_ = d_1599_v0_
                rhs252_ = ((d_1666_v54_)[default__.safeDivisionInt(150, ((d_1617_v15_)[(d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]] if ((d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]) in (d_1617_v15_) else (self).f16))] if (default__.safeDivisionInt(150, ((d_1617_v15_)[(d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]] if ((d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))]) in (d_1617_v15_) else (self).f16))) in (d_1666_v54_) else (((d_1666_v54_)[582] if (582) in (d_1666_v54_) else (self).f16)) * ((0) - ((self).f16)))
                lhs224_ = globalState
                lhs225_ = d_1734_v94_
                lhs226_ = default__.safeIndex(906, (d_1734_v94_).length(0))
                lhs227_ = globalState
                lhs228_ = globalState
                lhs224_.f1 = rhs248_
                lhs225_[lhs226_] = rhs249_
                lhs227_.f1 = rhs250_
                lhs228_.f1 = rhs251_
                r0 = rhs252_
        d_1599_v0_ = ((self).f16) >= (default__.safeModuloInt(len(_dafny.Map({(self).f16: d_1605_v6_})), (self).f16))
        (globalState).f6 = (self).f16
        r0 = (self).f16
        r1 = ((True) or ((d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))])) in ((default__.fm13((d_1601_v2_)[default__.safeIndex(854, (d_1601_v2_).length(0))], _dafny.CodePoint('f'), (self).f16, globalState)) + (d_1600_v1_))
        return r0, r1

    def m2(self, p0, globalState):
        d_1738_i0_: int
        d_1738_i0_ = 0
        with _dafny.label("13"):
            while p0:
                with _dafny.c_label("13"):
                    if (d_1738_i0_) >= (100):
                        raise _dafny.Break("13")
                    d_1738_i0_ = (d_1738_i0_) + (1)
                    d_1739_v0_: D14
                    d_1739_v0_ = D14_DC40(False)
                    d_1740_v1_: _dafny.Seq
                    d_1740_v1_ = _dafny.SeqWithoutIsStrInference([p0, (d_1739_v0_).cf71])
                    d_1740_v1_ = (_dafny.SeqWithoutIsStrInference([p0])).set(default__.safeIndex(((self).f16 if p0 else (self).f16), len(_dafny.SeqWithoutIsStrInference([p0]))), p0)
                    (globalState).f6 = 386
                    d_1741_v2_: _dafny.Array
                    nw325_ = _dafny.Array(D3.default()(), 12)
                    d_1741_v2_ = nw325_
                    d_1742_v3_: _dafny.Set
                    d_1742_v3_ = _dafny.Set({d_1741_v2_, d_1741_v2_})
                    d_1743_v4_: _dafny.Seq
                    d_1743_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqngik"))
                    d_1744_v5_: _dafny.Seq
                    d_1744_v5_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, True]), d_1740_v1_, d_1740_v1_])
                    d_1745_v6_: T0
                    nw326_ = C1()
                    nw326_.ctor__((self).f16)
                    d_1745_v6_ = nw326_
                    d_1746_v7_: _dafny.Seq
                    d_1746_v7_ = _dafny.SeqWithoutIsStrInference([d_1745_v6_, d_1745_v6_])
                    d_1747_v8_: _dafny.Map
                    d_1747_v8_ = _dafny.Map({p0: d_1746_v7_})
                    d_1748_v9_: _dafny.Map
                    d_1748_v9_ = _dafny.Map({(self).f16: len(d_1740_v1_)})
                    d_1749_v10_: _dafny.Array
                    nw327_ = _dafny.Array(None, 5)
                    nw327_[int(0)] = len(d_1747_v8_)
                    nw327_[int(1)] = ((d_1748_v9_)[(self).f16] if ((self).f16) in (d_1748_v9_) else (self).f16)
                    nw327_[int(2)] = (self).f16
                    nw327_[int(3)] = (0) - ((self).f16)
                    nw327_[int(4)] = (self).f16
                    d_1749_v10_ = nw327_
                    d_1750_v11_: D9
                    d_1750_v11_ = D9_DC23(d_1742_v3_, d_1743_v4_, ((d_1744_v5_)[default__.safeIndex((self).f16, len(d_1744_v5_))]) < (_dafny.SeqWithoutIsStrInference([p0, p0])), d_1749_v10_, default__.fm29(p0, (self).f16, (self).f16, p0, globalState))
                    source14_ = d_1750_v11_
                    if source14_.is_DC23:
                        d_1751___mcc_h0_ = source14_.cf37
                        d_1752___mcc_h1_ = source14_.cf38
                        d_1753___mcc_h2_ = source14_.cf39
                        d_1754___mcc_h3_ = source14_.cf40
                        d_1755___mcc_h4_ = source14_.cf41
                        d_1756_cf41_ = d_1755___mcc_h4_
                        d_1757_cf40_ = d_1754___mcc_h3_
                        d_1758_cf39_ = d_1753___mcc_h2_
                        d_1759_cf38_ = d_1752___mcc_h1_
                        d_1760_cf37_ = d_1751___mcc_h0_
                        d_1761_v12_: _dafny.Set
                        d_1761_v12_ = _dafny.Set({p0, p0})
                        d_1761_v12_ = ((d_1761_v12_) - (d_1761_v12_) if (d_1740_v1_)[default__.safeIndex((self).f16, len(d_1740_v1_))] else d_1761_v12_)
                        d_1740_v1_ = d_1740_v1_
                        d_1762_v13_: _dafny.Array
                        nw328_ = _dafny.Array(_dafny.Array(None, 0), 19)
                        d_1762_v13_ = nw328_
                        index306_ = default__.safeIndex(569, (d_1762_v13_).length(0))
                        (d_1762_v13_)[index306_] = d_1749_v10_
                        index307_ = default__.safeIndex(569, (d_1762_v13_).length(0))
                        (d_1762_v13_)[index307_] = d_1749_v10_
                        d_1758_cf39_ = default__.fm2(globalState)
                    elif source14_.is_DC24:
                        d_1763___mcc_h5_ = source14_.cf42
                        d_1764___mcc_h6_ = source14_.cf43
                        d_1765_cf43_ = d_1764___mcc_h6_
                        d_1766_cf42_ = d_1763___mcc_h5_
                        (globalState).f1 = default__.fm2(globalState)
                        d_1767_v14_: C1
                        nw329_ = C1()
                        nw329_.ctor__((self).f16)
                        d_1767_v14_ = nw329_
                        d_1768_v15_: _dafny.Seq
                        d_1768_v15_ = _dafny.SeqWithoutIsStrInference([(d_1767_v14_).f24, -959, (self).f16, (self).f16, 148])
                        d_1769_v16_: str
                        d_1769_v16_ = _dafny.CodePoint('e')
                        d_1770_v17_: D3
                        d_1770_v17_ = D3_DC8(d_1768_v15_, d_1743_v4_, p0, d_1769_v16_)
                        d_1771_v18_: _dafny.Map
                        d_1771_v18_ = _dafny.Map({p0: d_1770_v17_})
                        d_1772_v19_: _dafny.MultiSet
                        d_1772_v19_ = _dafny.MultiSet([p0])
                        d_1771_v18_ = (d_1771_v18_).set((d_1772_v19_).ispropersubset(d_1772_v19_), d_1770_v17_)
                        d_1748_v9_ = (_dafny.Map({len(d_1765_cf43_): (self).f16}) if (p0 if p0 else p0) else d_1748_v9_)
                    elif source14_.is_DC22:
                        d_1773___mcc_h7_ = source14_.cf36
                        d_1774_cf36_ = d_1773___mcc_h7_
                        (globalState).f1 = p0
                        d_1775_v21_: _dafny.Map
                        def iife188_():
                            coll66_ = _dafny.Map()
                            compr_66_: int
                            for compr_66_ in (d_1748_v9_).keys.Elements:
                                d_1776_v20_: int = compr_66_
                                if (d_1776_v20_) in (d_1748_v9_):
                                    coll66_[default__.safeModuloInt(d_1776_v20_, (self).f16)] = p0
                            return _dafny.Map(coll66_)
                        d_1775_v21_ = _dafny.Map({iife188_()
                        : p0})
                        d_1777_v22_: _dafny.Map
                        d_1777_v22_ = _dafny.Map({(self).f16: default__.fm2(globalState)})
                        d_1775_v21_ = (d_1775_v21_).set(d_1777_v22_, p0)
                        d_1778_v23_: D20
                        d_1778_v23_ = D20_DC50(p0, (self).f16, 719, p0, p0)
                        d_1779_v24_: _dafny.Seq
                        d_1779_v24_ = _dafny.SeqWithoutIsStrInference([d_1778_v23_, d_1778_v23_])
                        (globalState).f6 = len(d_1779_v24_)
                        (globalState).f6 = ((self).f16) + ((self).f16)
                    elif True:
                        d_1780___mcc_h8_ = source14_.cf44
                        d_1781_cf44_ = d_1780___mcc_h8_
                        d_1782_v25_: _dafny.Seq
                        d_1782_v25_ = _dafny.SeqWithoutIsStrInference([d_1749_v10_, d_1749_v10_, d_1749_v10_, (d_1750_v11_).cf40, d_1749_v10_])
                        d_1783_v26_: str
                        d_1783_v26_ = _dafny.CodePoint('a')
                        d_1784_v27_: _dafny.Map
                        d_1784_v27_ = _dafny.Map({len(d_1740_v1_): d_1783_v26_})
                        d_1785_v28_: _dafny.MultiSet
                        d_1785_v28_ = _dafny.MultiSet([len(d_1784_v27_)])
                        d_1786_v29_: _dafny.Map
                        d_1786_v29_ = _dafny.Map({p0: (d_1785_v28_).cardinality})
                        d_1782_v25_ = ((d_1782_v25_ if not (p0) or (p0) else (d_1782_v25_ if p0 else d_1782_v25_))).set(default__.safeIndex((0) - (default__.safeModuloInt((self).f16, ((d_1786_v29_)[p0] if (p0) in (d_1786_v29_) else (self).f16))), len((d_1782_v25_ if not (p0) or (p0) else (d_1782_v25_ if p0 else d_1782_v25_)))), d_1749_v10_)
                        d_1787_v30_: _dafny.Array
                        nw330_ = _dafny.Array(None, 2)
                        nw330_[int(0)] = p0
                        nw330_[int(1)] = p0
                        d_1787_v30_ = nw330_
                        index308_ = default__.safeIndex(529, (d_1787_v30_).length(0))
                        (d_1787_v30_)[index308_] = p0
                        index309_ = default__.safeIndex(529, (d_1787_v30_).length(0))
                        (d_1787_v30_)[index309_] = p0
                        d_1745_v6_ = d_1745_v6_
                        d_1788_v31_: int
                        d_1789_v32_: bool
                        out12_: int
                        out13_: bool
                        out12_, out13_ = (d_1745_v6_).m1(globalState)
                        d_1788_v31_ = out12_
                        d_1789_v32_ = out13_
                    d_1790_v33_: _dafny.Seq
                    d_1790_v33_ = _dafny.SeqWithoutIsStrInference([(self).f16, 971])
                    d_1791_v34_: _dafny.MultiSet
                    d_1791_v34_ = _dafny.MultiSet([(self).f16])
                    d_1792_v36_: _dafny.Map
                    d_1792_v36_ = _dafny.Map({(self).f16: p0})
                    d_1793_v37_: _dafny.Map
                    def iife189_():
                        coll67_ = _dafny.Map()
                        compr_67_: int
                        for compr_67_ in (d_1792_v36_).keys.Elements:
                            d_1794_v35_: int = compr_67_
                            if (d_1794_v35_) in (d_1792_v36_):
                                coll67_[(d_1794_v35_) * (630)] = p0
                        return _dafny.Map(coll67_)
                    d_1793_v37_ = _dafny.Map({iife189_()
                    : (self).f16})
                    d_1795_v38_: C2
                    nw331_ = C2()
                    nw331_.ctor__((d_1790_v33_)[default__.safeIndex(((d_1748_v9_)[727] if (727) in (d_1748_v9_) else (d_1791_v34_).cardinality), len(d_1790_v33_))], default__.safeModuloInt((0) - ((self).f16), (self).f16), d_1793_v37_)
                    d_1795_v38_ = nw331_
                    pass
            pass
        d_1796_v39_: _dafny.Map
        d_1796_v39_ = _dafny.Map({693: _dafny.Set({(self).f16})})
        source15_ = d_1796_v39_
        d_1797___mcc_h9_ = source15_
        d_1798_cf19_ = d_1797___mcc_h9_
        d_1799_v41_: _dafny.Array
        nw332_ = _dafny.Array(D0.default()(), 19)
        d_1799_v41_ = nw332_
        d_1800_v42_: _dafny.Array
        d_1800_v42_ = d_1799_v41_
        d_1801_v43_: _dafny.Array
        nw333_ = _dafny.Array(int(0), 24)
        d_1801_v43_ = nw333_
        d_1802_v44_: _dafny.Seq
        d_1802_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bogbeg"))
        def iife190_():
            coll68_ = _dafny.Set()
            compr_68_: int
            for compr_68_ in _dafny.IntegerRange(972, 980):
                d_1803_v40_: int = compr_68_
                if ((972) <= (d_1803_v40_)) and ((d_1803_v40_) < (980)):
                    coll68_ = coll68_.union(_dafny.Set([(d_1803_v40_) * ((self).f16)]))
            return _dafny.Set(coll68_)
        default__.m0(iife190_()
        , (d_1800_v42_), d_1801_v43_, d_1802_v44_, globalState)
        (globalState).f9 = d_1801_v43_
        d_1804_v45_: _dafny.Map
        d_1804_v45_ = _dafny.Map({(self).f16: p0})
        d_1805_v46_: _dafny.Map
        d_1805_v46_ = _dafny.Map({d_1801_v43_: (self).f16})
        d_1806_v47_: _dafny.Map
        d_1806_v47_ = _dafny.Map({d_1804_v45_: (self).f16})
        d_1807_v48_: T1
        nw334_ = C3()
        nw334_.ctor__(p0, (self).f16, d_1806_v47_)
        d_1807_v48_ = nw334_
        d_1808_v49_: _dafny.Map
        d_1808_v49_ = _dafny.Map({p0: d_1807_v48_})
        d_1809_v50_: C3
        nw335_ = C3()
        nw335_.ctor__(p0, len(d_1802_v44_), _dafny.Map({d_1804_v45_: ((d_1805_v46_)[d_1801_v43_] if (d_1801_v43_) in (d_1805_v46_) else len(d_1808_v49_))}))
        d_1809_v50_ = nw335_
        if p0:
            d_1810_v51_: D20
            d_1810_v51_ = D20_DC50((d_1809_v50_).f18, (d_1809_v50_).f19, (self).f16, not(default__.fm2(globalState)), default__.fm2(globalState))
            d_1811_v52_: _dafny.MultiSet
            d_1811_v52_ = _dafny.MultiSet([(d_1809_v50_).f18])
            d_1812_v53_: _dafny.MultiSet
            d_1812_v53_ = _dafny.MultiSet([((d_1811_v52_)[p0] if (p0) in (d_1811_v52_) else (0) - ((d_1809_v50_).f19))])
            pat_let_tv37_ = d_1809_v50_
            pat_let_tv38_ = d_1812_v53_
            pat_let_tv39_ = d_1809_v50_
            def iife191_(_pat_let61_0):
                def iife192_(d_1813_dt__update__tmp_h0_):
                    def iife193_(_pat_let62_0):
                        def iife194_(d_1814_dt__update_hcf89_h0_):
                            def iife195_(_pat_let63_0):
                                def iife196_(d_1815_dt__update_hcf87_h0_):
                                    def iife197_(_pat_let64_0):
                                        def iife198_(d_1816_dt__update_hcf88_h0_):
                                            return D20_DC50((d_1813_dt__update__tmp_h0_).cf86, d_1815_dt__update_hcf87_h0_, d_1816_dt__update_hcf88_h0_, d_1814_dt__update_hcf89_h0_, (d_1813_dt__update__tmp_h0_).cf90)
                                        return iife198_(_pat_let64_0)
                                    return iife197_((pat_let_tv39_).f19)
                                return iife196_(_pat_let63_0)
                            return iife195_((pat_let_tv38_).cardinality)
                        return iife194_(_pat_let62_0)
                    return iife193_((pat_let_tv37_).f18)
                return iife192_(_pat_let61_0)
            (globalState).f1 = not((iife191_(d_1810_v51_)).cf86)
            d_1802_v44_ = d_1802_v44_
            d_1817_v54_: _dafny.Seq
            d_1817_v54_ = _dafny.SeqWithoutIsStrInference([d_1801_v43_])
            d_1818_v55_: _dafny.Set
            d_1818_v55_ = _dafny.Set({(d_1809_v50_).f19})
            d_1819_v56_: D3
            d_1819_v56_ = D3_DC6((d_1817_v54_)[default__.safeIndex(len(d_1818_v55_), len(d_1817_v54_))])
            d_1820_v57_: _dafny.Seq
            d_1820_v57_ = _dafny.SeqWithoutIsStrInference([d_1819_v56_])
            d_1821_v58_: C0
            nw336_ = C0()
            nw336_.ctor__(d_1820_v57_, (d_1809_v50_).f18)
            d_1821_v58_ = nw336_
            d_1822_v59_: _dafny.Map
            d_1822_v59_ = _dafny.Map({(d_1809_v50_).f19: d_1821_v58_})
            d_1823_v60_: _dafny.Array
            nw337_ = _dafny.Array(False, 12)
            d_1823_v60_ = nw337_
            d_1824_v61_: _dafny.Seq
            d_1824_v61_ = _dafny.SeqWithoutIsStrInference([d_1823_v60_, d_1823_v60_])
            index310_ = default__.safeIndex(584, (d_1801_v43_).length(0))
            (d_1801_v43_)[index310_] = len(((d_1822_v59_).set((d_1809_v50_).f19, d_1821_v58_)) | (_dafny.Map({len(d_1824_v61_): d_1821_v58_})))
            index311_ = default__.safeIndex(584, (d_1801_v43_).length(0))
            (d_1801_v43_)[index311_] = (d_1809_v50_).f19
            (globalState).f6 = (609) * ((d_1809_v50_).f19)
            index312_ = default__.safeIndex(584, (d_1801_v43_).length(0))
            (d_1801_v43_)[index312_] = len(d_1802_v44_)
        elif True:
            (globalState).f6 = default__.fm1(p0, globalState)
            d_1825_v62_: C4
            nw338_ = C4()
            nw338_.ctor__()
            d_1825_v62_ = nw338_
            d_1826_v63_: str
            d_1826_v63_ = _dafny.CodePoint('j')
            d_1826_v63_ = d_1826_v63_
            index313_ = default__.safeIndex(761, (d_1801_v43_).length(0))
            (d_1801_v43_)[index313_] = (self).f16
            index314_ = default__.safeIndex(761, (d_1801_v43_).length(0))
            (d_1801_v43_)[index314_] = (d_1809_v50_).f19
            d_1827_v64_: _dafny.MultiSet
            d_1827_v64_ = _dafny.MultiSet([(d_1809_v50_).f18])
            (globalState).f6 = (((d_1827_v64_)[default__.fm2(globalState)] if (default__.fm2(globalState)) in (d_1827_v64_) else 328) if (d_1809_v50_).f18 else default__.safeDivisionInt((self).f16, (d_1809_v50_).f19))
        d_1828_v65_: _dafny.Map
        d_1828_v65_ = _dafny.Map({(self).f16: p0})
        d_1829_v66_: _dafny.Map
        d_1829_v66_ = _dafny.Map({p0: (0) - ((self).f16)})
        d_1830_v67_: _dafny.Map
        d_1830_v67_ = d_1829_v66_
        def lambda82_(source16_):
            d_1831___mcc_h10_ = source16_
            d_1832_cf12_ = d_1831___mcc_h10_
            def iife199_():
                coll69_ = _dafny.Map()
                compr_69_: int
                for compr_69_ in (_dafny.Set({(self).f16, (self).f16, 715})).Elements:
                    d_1833_v68_: int = compr_69_
                    if (d_1833_v68_) in (_dafny.Set({(self).f16, (self).f16, 715})):
                        coll69_[default__.safeModuloInt(d_1833_v68_, (self).f16)] = False
                return _dafny.Map(coll69_)
            return iife199_()
            

        d_1828_v65_ = lambda82_(d_1830_v67_)
        d_1834_v69_: _dafny.Array
        nw339_ = _dafny.Array(_dafny.Array(None, 0), 29)
        d_1834_v69_ = nw339_
        d_1835_v70_: D10
        d_1835_v70_ = D10_DC26(d_1834_v69_)
        d_1835_v70_ = D10_DC26(d_1834_v69_)
        d_1836_v71_: C4
        nw340_ = C4()
        nw340_.ctor__()
        d_1836_v71_ = nw340_
        d_1837_v72_: _dafny.Array
        def lambda83_(d_1838_p0_):
            def lambda84_(d_1839_i1_):
                return d_1838_p0_

            return lambda84_

        init50_ = lambda83_(p0)
        nw341_ = _dafny.Array(None, 3)
        for i0_50_ in range(nw341_.length(0)):
            nw341_[i0_50_] = init50_(i0_50_)
        d_1837_v72_ = nw341_
        d_1840_v73_: _dafny.Array
        nw342_ = _dafny.Array(None, 7)
        nw342_[int(0)] = d_1837_v72_
        nw342_[int(1)] = d_1837_v72_
        nw342_[int(2)] = d_1837_v72_
        nw342_[int(3)] = d_1837_v72_
        nw342_[int(4)] = d_1837_v72_
        nw342_[int(5)] = d_1837_v72_
        nw342_[int(6)] = d_1837_v72_
        d_1840_v73_ = nw342_
        d_1841_v74_: D5
        d_1841_v74_ = D5_DC10(d_1840_v73_)
        source17_ = d_1841_v74_
        if source17_.is_DC11:
            d_1842___mcc_h11_ = source17_.cf14
            d_1843_cf14_ = d_1842___mcc_h11_
            d_1844_v75_: str
            d_1844_v75_ = _dafny.CodePoint('s')
            d_1845_v76_: _dafny.Map
            d_1845_v76_ = _dafny.Map({p0: d_1844_v75_})
            d_1846_v77_: _dafny.Map
            d_1846_v77_ = _dafny.Map({False: p0})
            d_1847_v78_: D14
            d_1847_v78_ = D14_DC39(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p0, p0])) for d_1848_i2_ in range(default__.abs(682))]), d_1846_v77_, p0)
            d_1849_v79_: _dafny.Array
            nw343_ = _dafny.Array(None, 24)
            nw343_[int(0)] = (_dafny.Map({p0: d_1844_v75_})).set(p0, d_1844_v75_)
            nw343_[int(1)] = (d_1845_v76_) | (d_1845_v76_)
            nw343_[int(2)] = (d_1845_v76_) | (d_1845_v76_)
            nw343_[int(3)] = d_1845_v76_
            nw343_[int(4)] = (_dafny.Map({(d_1847_v78_).cf70: d_1844_v75_})) | (d_1845_v76_)
            nw343_[int(5)] = d_1845_v76_
            nw343_[int(6)] = d_1845_v76_
            nw343_[int(7)] = d_1845_v76_
            nw343_[int(8)] = (d_1845_v76_).set(p0, d_1844_v75_)
            nw343_[int(9)] = _dafny.Map({p0: _dafny.CodePoint('g')})
            nw343_[int(10)] = d_1845_v76_
            nw343_[int(11)] = d_1845_v76_
            nw343_[int(12)] = d_1845_v76_
            nw343_[int(13)] = d_1845_v76_
            nw343_[int(14)] = (_dafny.Map({p0: d_1844_v75_})) | (d_1845_v76_)
            nw343_[int(15)] = d_1845_v76_
            nw343_[int(16)] = d_1845_v76_
            nw343_[int(17)] = (d_1845_v76_) | (_dafny.Map({p0: d_1844_v75_}))
            nw343_[int(18)] = d_1845_v76_
            nw343_[int(19)] = _dafny.Map({p0: d_1844_v75_})
            nw343_[int(20)] = d_1845_v76_
            nw343_[int(21)] = d_1845_v76_
            nw343_[int(22)] = d_1845_v76_
            nw343_[int(23)] = d_1845_v76_
            d_1849_v79_ = nw343_
            d_1850_v80_: _dafny.Seq
            d_1850_v80_ = _dafny.SeqWithoutIsStrInference([(0) - (d_1843_cf14_), d_1843_cf14_])
            d_1851_v81_: _dafny.Seq
            d_1851_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkxwacsi"))
            rhs253_ = d_1849_v79_
            rhs254_ = d_1850_v80_
            rhs255_ = d_1844_v75_
            rhs256_ = d_1851_v81_
            d_1849_v79_ = rhs253_
            d_1850_v80_ = rhs254_
            d_1844_v75_ = rhs255_
            d_1851_v81_ = rhs256_
            if p0:
                d_1852_v82_: D5
                d_1852_v82_ = D5_DC11(d_1843_cf14_)
                d_1853_v83_: _dafny.Array
                nw344_ = _dafny.Array(None, 10)
                nw344_[int(0)] = (d_1852_v82_).cf14
                nw344_[int(1)] = (d_1850_v80_)[default__.safeIndex(d_1843_cf14_, len(d_1850_v80_))]
                nw344_[int(2)] = d_1843_cf14_
                nw344_[int(3)] = len(_dafny.Set({(self).f16}))
                nw344_[int(4)] = default__.fm1(p0, globalState)
                nw344_[int(5)] = -536
                nw344_[int(6)] = d_1843_cf14_
                nw344_[int(7)] = 363
                nw344_[int(8)] = (self).f16
                nw344_[int(9)] = (self).f16
                d_1853_v83_ = nw344_
                d_1854_v84_: _dafny.Set
                d_1854_v84_ = _dafny.Set({p0, p0})
                index315_ = default__.safeIndex(232, (d_1853_v83_).length(0))
                (d_1853_v83_)[index315_] = len(d_1854_v84_)
                index316_ = default__.safeIndex(232, (d_1853_v83_).length(0))
                (d_1853_v83_)[index316_] = (self).f16
                (globalState).f6 = 407
                d_1855_v85_: _dafny.Array
                nw345_ = _dafny.Array(D14.default()(), 27)
                d_1855_v85_ = nw345_
                d_1855_v85_ = d_1855_v85_
                d_1851_v81_ = d_1851_v81_
                (globalState).f1 = not ((d_1843_cf14_) >= ((0) - ((self).f16))) or (default__.fm2(globalState))
            elif True:
                (globalState).f1 = not(p0)
                d_1856_v86_: _dafny.Array
                def lambda85_(d_1857_cf14_):
                    def lambda86_(d_1858_i3_):
                        return (_dafny.MultiSet([(self).f16, 860, (self).f16, (self).f16])) | (_dafny.MultiSet([d_1857_cf14_]))

                    return lambda86_

                init51_ = lambda85_(d_1843_cf14_)
                nw346_ = _dafny.Array(None, 2)
                for i0_51_ in range(nw346_.length(0)):
                    nw346_[i0_51_] = init51_(i0_51_)
                d_1856_v86_ = nw346_
                d_1859_v89_: _dafny.Seq
                def iife200_():
                    coll70_ = _dafny.Set()
                    compr_70_: int
                    for compr_70_ in _dafny.IntegerRange(880, -220):
                        d_1860_v88_: int = compr_70_
                        if ((880) <= (d_1860_v88_)) and ((d_1860_v88_) < (-220)):
                            coll70_ = coll70_.union(_dafny.Set([default__.safeModuloInt(d_1860_v88_, 239)]))
                    return _dafny.Set(coll70_)
                d_1859_v89_ = _dafny.SeqWithoutIsStrInference([iife200_()
                ])
                index317_ = default__.safeIndex(188, (d_1856_v86_).length(0))
                def iife201_():
                    coll71_ = _dafny.Map()
                    compr_71_: _dafny.Set
                    for compr_71_ in (d_1859_v89_).Elements:
                        d_1861_v87_: _dafny.Set = compr_71_
                        if (d_1861_v87_) in (d_1859_v89_):
                            coll71_[d_1861_v87_] = p0
                    return _dafny.Map(coll71_)
                (d_1856_v86_)[index317_] = _dafny.MultiSet([((d_1829_v66_)[p0] if (p0) in (d_1829_v66_) else len(iife201_()
                ))])
                d_1862_v90_: _dafny.Map
                d_1862_v90_ = _dafny.Map({len(d_1851_v81_): (0) - (d_1843_cf14_)})
                d_1863_v91_: _dafny.MultiSet
                d_1863_v91_ = _dafny.MultiSet([(self).f16])
                index318_ = default__.safeIndex(188, (d_1856_v86_).length(0))
                (d_1856_v86_)[index318_] = (_dafny.MultiSet([d_1843_cf14_, ((d_1862_v90_)[(self).f16] if ((self).f16) in (d_1862_v90_) else (self).f16)])) - (d_1863_v91_)
                (globalState).f1 = ((p0 if p0 else True) if not (p0) or (p0) else p0)
                (globalState).f6 = len(((d_1851_v81_) + (_dafny.SeqWithoutIsStrInference([d_1844_v75_ for d_1864_i4_ in range(default__.abs(878))]))) + (d_1851_v81_))
                index319_ = default__.safeIndex(131, (d_1837_v72_).length(0))
                (d_1837_v72_)[index319_] = (d_1844_v75_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emichywvs")))
                index320_ = default__.safeIndex(131, (d_1837_v72_).length(0))
                (d_1837_v72_)[index320_] = (d_1843_cf14_) >= (len(default__.fm21(globalState)))
            d_1865_v92_: _dafny.Map
            d_1865_v92_ = _dafny.Map({_dafny.Map({d_1843_cf14_: p0}): d_1843_cf14_})
            d_1866_v93_: C3
            nw347_ = C3()
            nw347_.ctor__(p0, len(d_1846_v77_), d_1865_v92_)
            d_1866_v93_ = nw347_
            d_1867_v94_: _dafny.Set
            d_1867_v94_ = _dafny.Set({d_1866_v93_})
            d_1868_v95_: _dafny.Map
            d_1868_v95_ = _dafny.Map({(0) - (((d_1829_v66_)[p0] if (p0) in (d_1829_v66_) else len(d_1851_v81_))): d_1867_v94_})
            d_1868_v95_ = d_1868_v95_
            index321_ = default__.safeIndex(135, (d_1837_v72_).length(0))
            (d_1837_v72_)[index321_] = p0
            d_1869_v96_: _dafny.MultiSet
            d_1869_v96_ = _dafny.MultiSet([(d_1866_v93_).f18])
            d_1870_v97_: _dafny.Set
            d_1870_v97_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([len(d_1828_v65_) for d_1871_i5_ in range(default__.abs(368))]), _dafny.SeqWithoutIsStrInference([443, (d_1866_v93_).f19]), default__.fm12(d_1869_v96_, d_1843_cf14_, globalState)})
            d_1872_v98_: _dafny.Set
            d_1872_v98_ = _dafny.Set({(self).f16})
            index322_ = default__.safeIndex(135, (d_1837_v72_).length(0))
            rhs257_ = not(((d_1828_v65_)[831] if (831) in (d_1828_v65_) else not(p0)))
            rhs258_ = (_dafny.Set({_dafny.SeqWithoutIsStrInference([(d_1869_v96_).cardinality for d_1873_i6_ in range(default__.abs(752))]), _dafny.SeqWithoutIsStrInference([(d_1866_v93_).f19 for d_1874_i7_ in range(default__.abs(266))])})).issubset((d_1870_v97_).intersection(d_1870_v97_))
            rhs259_ = (d_1872_v98_).issubset((d_1872_v98_).intersection(d_1872_v98_))
            lhs229_ = globalState
            lhs230_ = globalState
            lhs231_ = d_1837_v72_
            lhs232_ = default__.safeIndex(135, (d_1837_v72_).length(0))
            lhs229_.f1 = rhs257_
            lhs230_.f1 = rhs258_
            lhs231_[lhs232_] = rhs259_
        elif source17_.is_DC12:
            d_1875___mcc_h12_ = source17_.cf15
            d_1876___mcc_h13_ = source17_.cf16
            d_1877___mcc_h14_ = source17_.cf17
            d_1878_cf17_ = d_1877___mcc_h14_
            d_1879_cf16_ = d_1876___mcc_h13_
            d_1880_cf15_ = d_1875___mcc_h12_
            (globalState).f1 = not(not(d_1879_cf16_))
            d_1881_v99_: int
            d_1882_v100_: bool
            out14_: int
            out15_: bool
            out14_, out15_ = (d_1836_v71_).m1(globalState)
            d_1881_v99_ = out14_
            d_1882_v100_ = out15_
            d_1883_v101_: _dafny.Seq
            d_1883_v101_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iohscrqtc"))
            d_1884_v102_: _dafny.Map
            d_1884_v102_ = _dafny.Map({(self).f16: len(d_1883_v101_)})
            if (len(d_1884_v102_)) == (d_1880_cf15_):
                (globalState).f6 = (self).f16
                d_1885_v104_: _dafny.Array
                def lambda87_(d_1886_i8_):
                    def iife202_():
                        coll72_ = _dafny.Set()
                        compr_72_: int
                        for compr_72_ in _dafny.IntegerRange(-897, -991):
                            d_1887_v103_: int = compr_72_
                            if ((-897) <= (d_1887_v103_)) and ((d_1887_v103_) < (-991)):
                                coll72_ = coll72_.union(_dafny.Set([(d_1887_v103_) + ((self).f16)]))
                        return _dafny.Set(coll72_)
                    return D0_DC2(D0_DC1(len(iife202_()
)))

                init52_ = lambda87_
                nw348_ = _dafny.Array(None, 4)
                for i0_52_ in range(nw348_.length(0)):
                    nw348_[i0_52_] = init52_(i0_52_)
                d_1885_v104_ = nw348_
                d_1888_v105_: _dafny.Array
                d_1888_v105_ = (d_1885_v104_ if p0 else d_1885_v104_)
                d_1888_v105_ = d_1888_v105_
                d_1880_cf15_ = (d_1880_cf15_) - (len(d_1829_v66_))
                d_1881_v99_ = d_1881_v99_
                d_1889_v106_: _dafny.MultiSet
                d_1889_v106_ = _dafny.MultiSet([d_1879_cf16_])
                d_1890_v107_: _dafny.Set
                d_1890_v107_ = _dafny.Set({d_1879_cf16_})
                d_1891_v108_: _dafny.Seq
                d_1891_v108_ = _dafny.SeqWithoutIsStrInference([533])
                d_1892_v109_: _dafny.Array
                nw349_ = _dafny.Array(None, 20)
                nw349_[int(0)] = d_1880_cf15_
                nw349_[int(1)] = (default__.fm1(not(p0), globalState)) * ((_dafny.MultiSet([d_1878_cf17_, d_1878_cf17_])).cardinality)
                nw349_[int(2)] = (self).f16
                nw349_[int(3)] = d_1881_v99_
                nw349_[int(4)] = (d_1880_cf15_) - (d_1880_cf15_)
                nw349_[int(5)] = len(_dafny.SeqWithoutIsStrInference([d_1880_cf15_]))
                nw349_[int(6)] = d_1880_cf15_
                nw349_[int(7)] = d_1881_v99_
                nw349_[int(8)] = (self).f16
                nw349_[int(9)] = default__.safeDivisionInt(d_1880_cf15_, 551)
                nw349_[int(10)] = d_1881_v99_
                nw349_[int(11)] = ((d_1889_v106_).cardinality) + (809)
                nw349_[int(12)] = default__.safeModuloInt(d_1880_cf15_, (self).f16)
                nw349_[int(13)] = len(d_1890_v107_)
                nw349_[int(14)] = (self).f16
                nw349_[int(15)] = default__.fm1(d_1882_v100_, globalState)
                nw349_[int(16)] = (d_1880_cf15_ if p0 else len(d_1891_v108_))
                nw349_[int(17)] = (d_1880_cf15_) - (d_1880_cf15_)
                nw349_[int(18)] = d_1881_v99_
                nw349_[int(19)] = d_1881_v99_
                d_1892_v109_ = nw349_
                index323_ = default__.safeIndex(284, (d_1892_v109_).length(0))
                (d_1892_v109_)[index323_] = d_1881_v99_
                d_1893_v110_: _dafny.Array
                nw350_ = _dafny.Array(D15.default()(), 14)
                d_1893_v110_ = nw350_
                d_1894_v111_: _dafny.Map
                d_1894_v111_ = _dafny.Map({not(d_1882_v100_): d_1893_v110_})
                d_1895_v112_: _dafny.Map
                d_1895_v112_ = _dafny.Map({d_1894_v111_: d_1881_v99_})
                index324_ = default__.safeIndex(284, (d_1892_v109_).length(0))
                (d_1892_v109_)[index324_] = (((d_1895_v112_)[d_1894_v111_] if (d_1894_v111_) in (d_1895_v112_) else len(d_1890_v107_)) if d_1879_cf16_ else d_1880_cf15_)
            elif True:
                (globalState).f1 = not(p0)
                (globalState).f1 = not(default__.fm2(globalState))
                (globalState).f14 = d_1837_v72_
                d_1896_v113_: _dafny.Set
                d_1896_v113_ = _dafny.Set({d_1879_cf16_})
                d_1897_v114_: _dafny.Set
                d_1897_v114_ = _dafny.Set({d_1896_v113_})
                d_1898_v115_: _dafny.Map
                d_1898_v115_ = _dafny.Map({_dafny.Map({d_1881_v99_: True}): (_dafny.Set({d_1896_v113_, d_1896_v113_})).issubset(d_1897_v114_)})
                d_1898_v115_ = (d_1898_v115_).set(d_1828_v65_, d_1879_cf16_)
                d_1899_v116_: _dafny.Array
                nw351_ = _dafny.Array(int(0), 13)
                d_1899_v116_ = nw351_
                d_1900_v117_: _dafny.MultiSet
                d_1900_v117_ = _dafny.MultiSet([d_1879_cf16_, p0])
                index325_ = default__.safeIndex(571, (d_1899_v116_).length(0))
                (d_1899_v116_)[index325_] = (d_1900_v117_).cardinality
                index326_ = default__.safeIndex(571, (d_1899_v116_).length(0))
                (d_1899_v116_)[index326_] = default__.safeDivisionInt((len(_dafny.Map({d_1881_v99_: d_1880_cf15_}))) + (154), d_1880_cf15_)
            d_1901_v118_: _dafny.Array
            nw352_ = _dafny.Array(int(0), 23)
            d_1901_v118_ = nw352_
            index327_ = default__.safeIndex(239, (d_1901_v118_).length(0))
            (d_1901_v118_)[index327_] = (788) + (d_1880_cf15_)
            index328_ = default__.safeIndex(239, (d_1901_v118_).length(0))
            (d_1901_v118_)[index328_] = ((d_1884_v102_)[((d_1884_v102_)[(self).f16] if ((self).f16) in (d_1884_v102_) else (0) - (d_1880_cf15_))] if (((d_1884_v102_)[(self).f16] if ((self).f16) in (d_1884_v102_) else (0) - (d_1880_cf15_))) in (d_1884_v102_) else d_1881_v99_)
        elif source17_.is_DC10:
            d_1902___mcc_h15_ = source17_.cf13
            d_1903_cf13_ = d_1902___mcc_h15_
            (globalState).f6 = default__.safeDivisionInt((self).f16, (self).f16)
            d_1904_v119_: _dafny.Array
            nw353_ = _dafny.Array(_dafny.MultiSet({}), 18)
            d_1904_v119_ = nw353_
            d_1905_v120_: _dafny.MultiSet
            d_1905_v120_ = _dafny.MultiSet([p0, p0])
            index329_ = default__.safeIndex(837, (d_1904_v119_).length(0))
            (d_1904_v119_)[index329_] = (d_1905_v120_).intersection(_dafny.MultiSet([not(p0), p0, p0, p0]))
            index330_ = default__.safeIndex(837, (d_1904_v119_).length(0))
            (d_1904_v119_)[index330_] = d_1905_v120_
            d_1906_v121_: D0
            d_1906_v121_ = D0_DC1((self).f16)
            source18_ = d_1906_v121_
            if source18_.is_DC1:
                d_1907___mcc_h17_ = source18_.cf1
                d_1908_cf1_ = d_1907___mcc_h17_
                d_1909_v122_: D0
                d_1909_v122_ = D0_DC1(d_1908_cf1_)
                d_1910_v123_: D0
                d_1910_v123_ = D0_DC2(d_1909_v122_)
                pat_let_tv40_ = d_1909_v122_
                d_1911_v124_: _dafny.Array
                nw354_ = _dafny.Array(None, 13)
                def iife203_(_pat_let65_0):
                    def iife204_(d_1912_dt__update__tmp_h1_):
                        def iife205_(_pat_let66_0):
                            def iife206_(d_1913_dt__update_hcf2_h0_):
                                return D0_DC2(d_1913_dt__update_hcf2_h0_)
                            return iife206_(_pat_let66_0)
                        return iife205_(pat_let_tv40_)
                    return iife204_(_pat_let65_0)
                nw354_[int(0)] = iife203_(d_1910_v123_)
                nw354_[int(1)] = d_1910_v123_
                nw354_[int(2)] = d_1910_v123_
                nw354_[int(3)] = d_1910_v123_
                nw354_[int(4)] = d_1910_v123_
                nw354_[int(5)] = d_1910_v123_
                nw354_[int(6)] = D0_DC2(D0_DC1((self).f16))
                nw354_[int(7)] = D0_DC2(d_1909_v122_)
                nw354_[int(8)] = d_1910_v123_
                nw354_[int(9)] = d_1910_v123_
                nw354_[int(10)] = d_1910_v123_
                nw354_[int(11)] = d_1910_v123_
                nw354_[int(12)] = d_1910_v123_
                d_1911_v124_ = nw354_
                d_1914_v125_: _dafny.Array
                nw355_ = _dafny.Array(None, 4)
                nw355_[int(0)] = 220
                nw355_[int(1)] = (self).f16
                nw355_[int(2)] = d_1908_cf1_
                nw355_[int(3)] = d_1908_cf1_
                d_1914_v125_ = nw355_
                default__.m0(_dafny.Set({(self).f16, d_1908_cf1_, d_1908_cf1_, d_1908_cf1_}), d_1911_v124_, d_1914_v125_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nruo")), globalState)
                d_1915_v126_: _dafny.Seq
                d_1915_v126_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1916_v127_: _dafny.Seq
                d_1916_v127_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mypoqw"))
                d_1917_v128_: _dafny.Map
                d_1917_v128_ = _dafny.Map({(self).f16: default__.safeDivisionInt(len(d_1915_v126_), len(d_1916_v127_))})
                d_1917_v128_ = (d_1917_v128_).set(d_1908_cf1_, (self).f16)
                d_1918_v129_: _dafny.MultiSet
                d_1918_v129_ = _dafny.MultiSet([(len(d_1916_v127_)) - (d_1908_cf1_)])
                d_1919_v130_: _dafny.Seq
                d_1919_v130_ = _dafny.SeqWithoutIsStrInference([d_1918_v129_, d_1918_v129_])
                d_1920_v131_: D11
                d_1920_v131_ = D11_DC32(p0, (self).f16, p0)
                d_1918_v129_ = (d_1919_v130_)[default__.safeIndex((d_1920_v131_).cf60, len(d_1919_v130_))]
                d_1908_cf1_ = (self).f16
            elif source18_.is_DC0:
                d_1921___mcc_h18_ = source18_.cf0
                d_1922_cf0_ = d_1921___mcc_h18_
                d_1923_v132_: _dafny.Map
                d_1923_v132_ = _dafny.Map({(self).f16: ((self).f16 if p0 else (self).f16)})
                (globalState).f6 = len(d_1923_v132_)
                (globalState).f1 = d_1922_cf0_
                d_1924_v133_: _dafny.Seq
                d_1924_v133_ = _dafny.SeqWithoutIsStrInference([not(d_1922_cf0_), d_1922_cf0_, d_1922_cf0_])
                d_1925_v134_: D13
                d_1925_v134_ = D13_DC37(not (not((d_1924_v133_)[default__.safeIndex((self).f16, len(d_1924_v133_))])) or (d_1922_cf0_), (d_1905_v120_).ispropersubset((d_1904_v119_)[default__.safeIndex(837, (d_1904_v119_).length(0))]))
                d_1925_v134_ = (d_1925_v134_ if p0 else d_1925_v134_)
                (globalState).f1 = p0
            elif True:
                d_1926___mcc_h19_ = source18_.cf2
                d_1927_cf2_ = d_1926___mcc_h19_
                d_1928_v135_: _dafny.Seq
                d_1928_v135_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt((self).f16, (self).f16)])
                d_1929_v136_: _dafny.Map
                d_1929_v136_ = _dafny.Map({(self).f16: (d_1928_v135_) + (_dafny.SeqWithoutIsStrInference([(0) - ((self).f16)]))})
                d_1928_v135_ = ((d_1929_v136_)[(self).f16] if ((self).f16) in (d_1929_v136_) else d_1928_v135_)
                (globalState).f6 = (self).f16
                (globalState).f6 = (0) - ((self).f16)
                (globalState).f1 = p0
            if p0:
                (globalState).f1 = False
                d_1930_v137_: _dafny.Seq
                d_1930_v137_ = _dafny.SeqWithoutIsStrInference([(self).f16, (self).f16, (self).f16])
                d_1931_v138_: _dafny.Array
                nw356_ = _dafny.Array(D0.default()(), 23)
                d_1931_v138_ = nw356_
                d_1932_v139_: _dafny.Array
                def lambda88_(d_1933_i9_):
                    return (d_1933_i9_) - ((0) - ((self).f16))

                init53_ = lambda88_
                nw357_ = _dafny.Array(None, 19)
                for i0_53_ in range(nw357_.length(0)):
                    nw357_[i0_53_] = init53_(i0_53_)
                d_1932_v139_ = nw357_
                default__.m0(_dafny.Set({(self).f16, len(d_1930_v137_), (self).f16, 687}), d_1931_v138_, d_1932_v139_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugbdhjeb")), globalState)
                d_1934_v140_: _dafny.Seq
                d_1934_v140_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgycqsugd"))
                d_1935_v141_: _dafny.MultiSet
                d_1935_v141_ = _dafny.MultiSet([len(_dafny.Map({(self).f16: p0})), len(d_1934_v140_)])
                (globalState).f1 = ((self).f16) < ((d_1935_v141_).cardinality)
                d_1936_v142_: C4
                nw358_ = C4()
                nw358_.ctor__()
                d_1936_v142_ = nw358_
                d_1828_v65_ = (d_1828_v65_).set((self).f16, p0)
            elif True:
                d_1937_v143_: str
                d_1937_v143_ = _dafny.CodePoint('x')
                d_1938_v144_: _dafny.Seq
                d_1938_v144_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydawcir"))
                d_1939_v145_: _dafny.Seq
                d_1939_v145_ = _dafny.SeqWithoutIsStrInference([361])
                d_1940_v146_: _dafny.Array
                nw359_ = _dafny.Array(None, 18)
                nw359_[int(0)] = (self).f16
                nw359_[int(1)] = (len(d_1938_v144_)) * ((self).f16)
                nw359_[int(2)] = (self).f16
                nw359_[int(3)] = len(d_1939_v145_)
                nw359_[int(4)] = ((self).f16) + ((self).f16)
                nw359_[int(5)] = (self).f16
                nw359_[int(6)] = (self).f16
                nw359_[int(7)] = (self).f16
                nw359_[int(8)] = (self).f16
                nw359_[int(9)] = (self).f16
                nw359_[int(10)] = (d_1939_v145_)[default__.safeIndex((self).f16, len(d_1939_v145_))]
                nw359_[int(11)] = (self).f16
                nw359_[int(12)] = (0) - ((self).f16)
                nw359_[int(13)] = (self).f16
                nw359_[int(14)] = 619
                nw359_[int(15)] = default__.safeDivisionInt((self).f16, -314)
                nw359_[int(16)] = (0) - (len(d_1829_v66_))
                nw359_[int(17)] = (self).f16
                d_1940_v146_ = nw359_
                d_1941_v147_: _dafny.Map
                d_1941_v147_ = _dafny.Map({(self).f16: (self).f16})
                rhs260_ = d_1937_v143_
                rhs261_ = 534
                rhs262_ = d_1940_v146_
                rhs263_ = ((d_1941_v147_)[(self).f16] if ((self).f16) in (d_1941_v147_) else (len(d_1938_v144_)) + ((self).f16))
                rhs264_ = (d_1939_v145_)[default__.safeIndex((self).f16, len(d_1939_v145_))]
                lhs233_ = globalState
                lhs234_ = globalState
                lhs235_ = globalState
                lhs236_ = globalState
                d_1937_v143_ = rhs260_
                lhs233_.f6 = rhs261_
                lhs234_.f9 = rhs262_
                lhs235_.f6 = rhs263_
                lhs236_.f6 = rhs264_
                d_1942_v148_: _dafny.Array
                nw360_ = _dafny.Array(_dafny.Map({}), 11)
                d_1942_v148_ = nw360_
                def lambda89_(d_1943_v147_):
                    def lambda90_(d_1944_i10_):
                        return (d_1943_v147_) | (d_1943_v147_)

                    return lambda90_

                init54_ = lambda89_(d_1941_v147_)
                nw361_ = _dafny.Array(None, 21)
                for i0_54_ in range(nw361_.length(0)):
                    nw361_[i0_54_] = init54_(i0_54_)
                d_1942_v148_ = nw361_
                index331_ = default__.safeIndex(617, (d_1837_v72_).length(0))
                (d_1837_v72_)[index331_] = p0
                index332_ = default__.safeIndex(617, (d_1837_v72_).length(0))
                (d_1837_v72_)[index332_] = p0
                d_1938_v144_ = default__.fm21(globalState)
                d_1945_v149_: _dafny.MultiSet
                d_1945_v149_ = _dafny.MultiSet([(D16_DC45((self).f16, d_1937_v143_, (self).f16, (d_1837_v72_)[default__.safeIndex(617, (d_1837_v72_).length(0))])).cf78])
                (globalState).f6 = (len(((d_1938_v144_).set(default__.safeIndex(((d_1945_v149_).set(default__.fm1(p0, globalState), default__.abs((self).f16))).cardinality, len(d_1938_v144_)), d_1937_v143_)) + (default__.fm21(globalState)))) - (len(d_1829_v66_))
        elif True:
            d_1946___mcc_h16_ = source17_.cf18
            d_1947_cf18_ = d_1946___mcc_h16_
            (globalState).f1 = p0
            index333_ = default__.safeIndex(270, (d_1840_v73_).length(0))
            (d_1840_v73_)[index333_] = d_1837_v72_
            index334_ = default__.safeIndex(270, (d_1840_v73_).length(0))
            (d_1840_v73_)[index334_] = d_1837_v72_
            (globalState).f1 = p0
            if p0:
                d_1948_v150_: str
                d_1948_v150_ = _dafny.CodePoint('t')
                d_1949_v151_: _dafny.MultiSet
                d_1949_v151_ = _dafny.MultiSet([d_1948_v150_, d_1948_v150_])
                d_1950_v152_: _dafny.Seq
                d_1950_v152_ = _dafny.SeqWithoutIsStrInference([(self).f16, (self).f16])
                d_1951_v153_: _dafny.Seq
                d_1951_v153_ = _dafny.SeqWithoutIsStrInference([(d_1950_v152_)[default__.safeIndex((self).f16, len(d_1950_v152_))]])
                d_1952_v154_: _dafny.Array
                nw362_ = _dafny.Array(None, 6)
                nw362_[int(0)] = (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1953_i11_ in range(default__.abs(565))]))) * ((self).f16)
                nw362_[int(1)] = (0) - ((self).f16)
                nw362_[int(2)] = (self).f16
                nw362_[int(3)] = (self).f16
                nw362_[int(4)] = default__.safeModuloInt(-430, (self).f16)
                nw362_[int(5)] = (((d_1949_v151_).set(d_1948_v150_, default__.abs(len(d_1950_v152_)))).cardinality) + (default__.fm1(p0, globalState))
                d_1952_v154_ = nw362_
                index335_ = default__.safeIndex(771, (d_1952_v154_).length(0))
                (d_1952_v154_)[index335_] = (self).f16
                index336_ = default__.safeIndex(771, (d_1952_v154_).length(0))
                (d_1952_v154_)[index336_] = default__.safeDivisionInt((self).f16, (self).f16)
                d_1954_v155_: _dafny.Seq
                d_1954_v155_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gg"))
                (globalState).f6 = default__.safeModuloInt(len(d_1828_v65_), ((d_1952_v154_)[default__.safeIndex(771, (d_1952_v154_).length(0))]) * (len(d_1954_v155_)))
                (globalState).f6 = (((d_1952_v154_)[default__.safeIndex(771, (d_1952_v154_).length(0))] if False else (d_1952_v154_)[default__.safeIndex(771, (d_1952_v154_).length(0))])) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gunwfodep"))))
                d_1955_v156_: int
                d_1956_v157_: bool
                out16_: int
                out17_: bool
                out16_, out17_ = (d_1836_v71_).m1(globalState)
                d_1955_v156_ = out16_
                d_1956_v157_ = out17_
                index337_ = default__.safeIndex(771, (d_1952_v154_).length(0))
                rhs265_ = d_1956_v157_
                rhs266_ = (self).f16
                lhs237_ = globalState
                lhs238_ = d_1952_v154_
                lhs239_ = default__.safeIndex(771, (d_1952_v154_).length(0))
                lhs237_.f1 = rhs265_
                lhs238_[lhs239_] = rhs266_
            elif True:
                d_1957_v158_: _dafny.Map
                d_1957_v158_ = _dafny.Map({p0: p0})
                d_1958_v159_: _dafny.Map
                d_1958_v159_ = _dafny.Map({not(((d_1957_v158_)[False] if (False) in (d_1957_v158_) else p0)): p0})
                d_1957_v158_ = (d_1958_v159_) | (d_1958_v159_)
                d_1959_v160_: _dafny.Map
                d_1959_v160_ = _dafny.Map({(self).f16: len(_dafny.SeqWithoutIsStrInference([p0, p0, True]))})
                d_1959_v160_ = (d_1959_v160_).set((self).f16, (self).f16)
                d_1960_v161_: str
                d_1960_v161_ = _dafny.CodePoint('b')
                d_1960_v161_ = d_1960_v161_
                d_1961_v162_: _dafny.Map
                d_1961_v162_ = _dafny.Map({d_1828_v65_: (0) - ((self).f16)})
                d_1962_v163_: C3
                nw363_ = C3()
                nw363_.ctor__(False, (self).f16, d_1961_v162_)
                d_1962_v163_ = nw363_
                (globalState).f6 = ((d_1962_v163_).f19) * ((self).f16)

    @property
    def f16(self):
        return self._f16

class C6(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self):
        pass
        pass

    def fm3(self, p0, p1, p2, p3, globalState):
        return D0_DC2(D0_DC0(True))

    def fm9(self, p0, globalState):
        return 995

    def fm10(self, p0, p1, p2, p3, globalState):
        return not(not (((_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbhi"))))])).cardinality) <= (720)) or ((True) == ((D0_DC0(False)).cf0)))

    def m1(self, globalState):
        r0: int = int(0)
        r1: bool = False
        d_1963_v0_: int
        d_1963_v0_ = 277
        d_1964_v1_: bool
        d_1964_v1_ = False
        d_1965_v2_: _dafny.Map
        d_1965_v2_ = _dafny.Map({d_1964_v1_: d_1964_v1_})
        d_1966_v3_: _dafny.Array
        nw364_ = _dafny.Array(None, 13)
        nw364_[int(0)] = (d_1963_v0_) + ((0) - (d_1963_v0_))
        nw364_[int(1)] = d_1963_v0_
        nw364_[int(2)] = (317) - (d_1963_v0_)
        nw364_[int(3)] = d_1963_v0_
        nw364_[int(4)] = d_1963_v0_
        nw364_[int(5)] = d_1963_v0_
        nw364_[int(6)] = d_1963_v0_
        nw364_[int(7)] = default__.fm1(not(d_1964_v1_), globalState)
        nw364_[int(8)] = (0) - (d_1963_v0_)
        nw364_[int(9)] = d_1963_v0_
        nw364_[int(10)] = (0) - (default__.safeModuloInt(d_1963_v0_, d_1963_v0_))
        nw364_[int(11)] = default__.safeDivisionInt(d_1963_v0_, d_1963_v0_)
        nw364_[int(12)] = (len((D1_DC3(d_1965_v2_)).cf3)) * (-379)
        d_1966_v3_ = nw364_
        d_1967_v4_: _dafny.Array
        nw365_ = _dafny.Array(None, 11)
        nw365_[int(0)] = d_1964_v1_
        nw365_[int(1)] = d_1964_v1_
        nw365_[int(2)] = d_1964_v1_
        nw365_[int(3)] = False
        nw365_[int(4)] = d_1964_v1_
        nw365_[int(5)] = d_1964_v1_
        nw365_[int(6)] = d_1964_v1_
        nw365_[int(7)] = d_1964_v1_
        nw365_[int(8)] = d_1964_v1_
        nw365_[int(9)] = d_1964_v1_
        nw365_[int(10)] = not(not(d_1964_v1_))
        d_1967_v4_ = nw365_
        d_1968_v5_: _dafny.Map
        d_1968_v5_ = _dafny.Map({d_1964_v1_: d_1967_v4_})
        index338_ = default__.safeIndex(8, (d_1966_v3_).length(0))
        (d_1966_v3_)[index338_] = len((d_1968_v5_) | (_dafny.Map({d_1964_v1_: d_1967_v4_})))
        index339_ = default__.safeIndex(8, (d_1966_v3_).length(0))
        (d_1966_v3_)[index339_] = (d_1963_v0_) * (default__.safeModuloInt(d_1963_v0_, 932))
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1966_v3_).length(0)):
            d_1969_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_1969_i0_)) and ((d_1969_i0_) < ((d_1966_v3_).length(0)))):
                (d_1966_v3_)[(d_1969_i0_)] = (d_1969_i0_) * (-161)
        d_1970_v6_: D0
        d_1970_v6_ = D0_DC0(d_1964_v1_)
        pat_let_tv41_ = d_1966_v3_
        pat_let_tv42_ = d_1966_v3_
        def lambda91_(source19_):
            if source19_.is_DC1:
                d_1971___mcc_h0_ = source19_.cf1
                d_1972_cf1_ = d_1971___mcc_h0_
                return d_1972_cf1_
            elif source19_.is_DC0:
                d_1973___mcc_h1_ = source19_.cf0
                d_1974_cf0_ = d_1973___mcc_h1_
                return (pat_let_tv42_)[default__.safeIndex(8, (pat_let_tv41_).length(0))]
            elif True:
                d_1975___mcc_h2_ = source19_.cf2
                d_1976_cf2_ = d_1975___mcc_h2_
                return (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krasuv"))))

        d_1963_v0_ = lambda91_(d_1970_v6_)
        d_1977_v7_: _dafny.Map
        d_1977_v7_ = _dafny.Map({(d_1966_v3_)[default__.safeIndex(8, (d_1966_v3_).length(0))]: True})
        d_1978_v8_: _dafny.Map
        d_1978_v8_ = _dafny.Map({d_1977_v7_: (d_1966_v3_)[default__.safeIndex(8, (d_1966_v3_).length(0))]})
        d_1979_v9_: C3
        nw366_ = C3()
        nw366_.ctor__(d_1964_v1_, d_1963_v0_, d_1978_v8_)
        d_1979_v9_ = nw366_
        d_1963_v0_ = d_1963_v0_
        nw367_ = _dafny.Array(False, 13)
        (globalState).f14 = nw367_
        r0 = (d_1966_v3_)[default__.safeIndex(8, (d_1966_v3_).length(0))]
        r1 = d_1964_v1_
        return r0, r1

    def m2(self, p0, globalState):
        d_1980_v2_: _dafny.Array
        def lambda92_(d_1981_p0_):
            def lambda93_(d_1982_i0_):
                def iife207_():
                    def iife209_():
                        coll75_ = _dafny.Set()
                        compr_75_: int
                        for compr_75_ in _dafny.IntegerRange(9, 452):
                            d_1985_v1_: int = compr_75_
                            if ((9) <= (d_1985_v1_)) and ((d_1985_v1_) < (452)):
                                coll75_ = coll75_.union(_dafny.Set([default__.safeModuloInt(d_1985_v1_, 466)]))
                        return _dafny.Set(coll75_)
                    coll73_ = _dafny.Map()
                    def iife208_():
                        coll74_ = _dafny.Set()
                        compr_74_: int
                        for compr_74_ in _dafny.IntegerRange(9, 452):
                            d_1983_v1_: int = compr_74_
                            if ((9) <= (d_1983_v1_)) and ((d_1983_v1_) < (452)):
                                coll74_ = coll74_.union(_dafny.Set([default__.safeModuloInt(d_1983_v1_, 466)]))
                        return _dafny.Set(coll74_)
                    compr_73_: int
                    for compr_73_ in (_dafny.Map({len(iife208_()
                    ): d_1981_p0_})).keys.Elements:
                        d_1984_v0_: int = compr_73_
                        if (d_1984_v0_) in (_dafny.Map({len(iife209_()
                        ): d_1981_p0_})):
                            coll73_[(d_1984_v0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cc"))))] = (_dafny.MultiSet([d_1981_p0_])).cardinality
                    return _dafny.Map(coll73_)
                return (_dafny.MultiSet([729, 484, (0) - (len(iife207_()
                ))])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-753, -323])))

            return lambda93_

        init55_ = lambda92_(p0)
        nw368_ = _dafny.Array(None, 15)
        for i0_55_ in range(nw368_.length(0)):
            nw368_[i0_55_] = init55_(i0_55_)
        d_1980_v2_ = nw368_
        d_1986_v3_: int
        d_1986_v3_ = 495
        d_1987_v4_: _dafny.MultiSet
        d_1987_v4_ = _dafny.MultiSet([d_1986_v3_])
        index340_ = default__.safeIndex(226, (d_1980_v2_).length(0))
        (d_1980_v2_)[index340_] = d_1987_v4_
        d_1988_v5_: str
        d_1988_v5_ = _dafny.CodePoint('r')
        d_1989_v6_: _dafny.Array
        nw369_ = _dafny.Array(int(0), 12)
        d_1989_v6_ = nw369_
        d_1990_v7_: _dafny.Map
        d_1990_v7_ = _dafny.Map({d_1986_v3_: d_1986_v3_})
        d_1991_v8_: _dafny.Map
        d_1991_v8_ = _dafny.Map({len(d_1990_v7_): d_1986_v3_})
        index341_ = default__.safeIndex(715, (d_1989_v6_).length(0))
        (d_1989_v6_)[index341_] = len(d_1991_v8_)
        index342_ = default__.safeIndex(226, (d_1980_v2_).length(0))
        index343_ = default__.safeIndex(715, (d_1989_v6_).length(0))
        rhs267_ = (d_1987_v4_) - (d_1987_v4_)
        rhs268_ = p0
        rhs269_ = d_1988_v5_
        rhs270_ = d_1986_v3_
        lhs240_ = d_1980_v2_
        lhs241_ = default__.safeIndex(226, (d_1980_v2_).length(0))
        lhs242_ = globalState
        lhs243_ = d_1989_v6_
        lhs244_ = default__.safeIndex(715, (d_1989_v6_).length(0))
        lhs240_[lhs241_] = rhs267_
        lhs242_.f1 = rhs268_
        d_1988_v5_ = rhs269_
        lhs243_[lhs244_] = rhs270_
        d_1992_v9_: _dafny.Array
        nw370_ = _dafny.Array(False, 3)
        d_1992_v9_ = nw370_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1992_v9_).length(0)):
            d_1993_i1_: int = guard_loop_5_
            if (True) and (((0) <= (d_1993_i1_)) and ((d_1993_i1_) < ((d_1992_v9_).length(0)))):
                (d_1992_v9_)[(d_1993_i1_)] = (p0) and (p0)
        index344_ = default__.safeIndex(715, (d_1989_v6_).length(0))
        (d_1989_v6_)[index344_] = (d_1989_v6_)[default__.safeIndex(715, (d_1989_v6_).length(0))]
        d_1994_i2_: int
        d_1994_i2_ = 0
        with _dafny.label("14"):
            while p0:
                with _dafny.c_label("14"):
                    if (d_1994_i2_) >= (100):
                        raise _dafny.Break("14")
                    d_1994_i2_ = (d_1994_i2_) + (1)
                    d_1995_v10_: _dafny.Map
                    d_1995_v10_ = _dafny.Map({d_1986_v3_: p0})
                    d_1996_v11_: _dafny.Map
                    d_1996_v11_ = _dafny.Map({d_1995_v10_: (0) - ((d_1989_v6_)[default__.safeIndex(715, (d_1989_v6_).length(0))])})
                    d_1997_v12_: C2
                    nw371_ = C2()
                    nw371_.ctor__((837) + (-800), d_1986_v3_, d_1996_v11_)
                    d_1997_v12_ = nw371_
                    if p0:
                        d_1998_v13_: _dafny.Array
                        d_1998_v13_ = d_1992_v9_
                        d_1999_v14_: _dafny.Seq
                        d_1999_v14_ = _dafny.SeqWithoutIsStrInference([d_1992_v9_, d_1992_v9_, (d_1998_v13_), d_1992_v9_, d_1992_v9_])
                        d_2000_v15_: _dafny.Array
                        nw372_ = _dafny.Array(None, 11)
                        nw372_[int(0)] = (d_1999_v14_)[default__.safeIndex(d_1997_v12_.f21, len(d_1999_v14_))]
                        nw372_[int(1)] = d_1992_v9_
                        nw372_[int(2)] = d_1992_v9_
                        nw372_[int(3)] = d_1992_v9_
                        nw372_[int(4)] = d_1992_v9_
                        nw372_[int(5)] = d_1992_v9_
                        nw372_[int(6)] = (d_1999_v14_)[default__.safeIndex(d_1986_v3_, len(d_1999_v14_))]
                        nw372_[int(7)] = d_1992_v9_
                        nw372_[int(8)] = d_1992_v9_
                        nw372_[int(9)] = d_1992_v9_
                        nw372_[int(10)] = d_1992_v9_
                        d_2000_v15_ = nw372_
                        d_2000_v15_ = d_2000_v15_
                        d_2001_v16_: C5
                        nw373_ = C5()
                        nw373_.ctor__(len(_dafny.Set({d_1989_v6_})))
                        d_2001_v16_ = nw373_
                        (globalState).f6 = d_1997_v12_.f21
                        index345_ = default__.safeIndex(456, (d_1992_v9_).length(0))
                        (d_1992_v9_)[index345_] = (d_1986_v3_) < (114)
                        index346_ = default__.safeIndex(456, (d_1992_v9_).length(0))
                        (d_1992_v9_)[index346_] = p0
                        (globalState).f1 = p0
                    elif True:
                        index347_ = default__.safeIndex(715, (d_1989_v6_).length(0))
                        (d_1989_v6_)[index347_] = default__.safeModuloInt(d_1997_v12_.f21, d_1986_v3_)
                        d_2002_v17_: _dafny.Seq
                        d_2002_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idh"))
                        d_2002_v17_ = d_2002_v17_
                        (globalState).f1 = p0
                        d_2003_v18_: _dafny.Map
                        d_2003_v18_ = _dafny.Map({d_1986_v3_: d_2002_v17_})
                        rhs271_ = default__.fm38(default__.safeModuloInt(default__.fm1(p0, globalState), default__.fm1(p0, globalState)), p0, p0, not(p0), globalState)
                        rhs272_ = 234
                        lhs245_ = d_1997_v12_
                        d_2003_v18_ = rhs271_
                        lhs245_.f21 = rhs272_
                        d_2004_v19_: _dafny.Array
                        nw374_ = _dafny.Array(None, 14)
                        nw374_[int(0)] = (_dafny.SeqWithoutIsStrInference([d_1988_v5_ for d_2005_i3_ in range(default__.abs(68))]) if p0 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ckh")))
                        nw374_[int(1)] = d_2002_v17_
                        nw374_[int(2)] = d_2002_v17_
                        nw374_[int(3)] = d_2002_v17_
                        nw374_[int(4)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyjs")) if p0 else d_2002_v17_)
                        nw374_[int(5)] = d_2002_v17_
                        nw374_[int(6)] = d_2002_v17_
                        nw374_[int(7)] = d_2002_v17_
                        nw374_[int(8)] = d_2002_v17_
                        nw374_[int(9)] = d_2002_v17_
                        nw374_[int(10)] = d_2002_v17_
                        nw374_[int(11)] = d_2002_v17_
                        nw374_[int(12)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2006_i4_ in range(default__.abs(686))])
                        nw374_[int(13)] = d_2002_v17_
                        d_2004_v19_ = nw374_
                        index348_ = default__.safeIndex(510, (d_2004_v19_).length(0))
                        (d_2004_v19_)[index348_] = ((d_2003_v18_)[d_1997_v12_.f21] if (d_1997_v12_.f21) in (d_2003_v18_) else d_2002_v17_)
                        index349_ = default__.safeIndex(510, (d_2004_v19_).length(0))
                        (d_2004_v19_)[index349_] = d_2002_v17_
                    d_1995_v10_ = (d_1995_v10_).set(d_1986_v3_, p0)
                    d_2007_v20_: _dafny.Array
                    nw375_ = _dafny.Array(_dafny.MultiSet({}), 12)
                    d_2007_v20_ = nw375_
                    d_2008_v21_: _dafny.Seq
                    d_2008_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvfjavan"))
                    index350_ = default__.safeIndex(852, (d_2007_v20_).length(0))
                    (d_2007_v20_)[index350_] = (_dafny.MultiSet([d_1989_v6_, d_1989_v6_])).set(d_1989_v6_, default__.abs(len(d_2008_v21_)))
                    d_2009_v22_: _dafny.Seq
                    d_2009_v22_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_2010_v23_: _dafny.MultiSet
                    d_2010_v23_ = _dafny.MultiSet([d_1989_v6_])
                    index351_ = default__.safeIndex(852, (d_2007_v20_).length(0))
                    rhs273_ = (default__.fm44(globalState)).cf65
                    rhs274_ = d_2010_v23_
                    rhs275_ = p0
                    rhs276_ = d_2009_v22_
                    lhs246_ = globalState
                    lhs247_ = d_2007_v20_
                    lhs248_ = default__.safeIndex(852, (d_2007_v20_).length(0))
                    lhs249_ = globalState
                    lhs246_.f1 = rhs273_
                    lhs247_[lhs248_] = rhs274_
                    lhs249_.f1 = rhs275_
                    d_2009_v22_ = rhs276_
                    pass
            pass
        index352_ = default__.safeIndex(715, (d_1989_v6_).length(0))
        (d_1989_v6_)[index352_] = (default__.safeDivisionInt(d_1986_v3_, 880)) + (default__.fm1(p0, globalState))
        d_2011_v24_: _dafny.Seq
        d_2011_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
        d_2011_v24_ = default__.fm21(globalState)

    def m8(self, p0, globalState):
        r0: bool = False
        d_2012_v0_: _dafny.Array
        nw376_ = _dafny.Array(_dafny.Map({}), 2)
        d_2012_v0_ = nw376_
        d_2013_v1_: D20
        d_2013_v1_ = D20_DC49(d_2012_v0_)
        pat_let_tv43_ = d_2012_v0_
        pat_let_tv44_ = d_2012_v0_
        d_2014_v2_: _dafny.Array
        nw377_ = _dafny.Array(None, 8)
        nw377_[int(0)] = D20_DC49(d_2012_v0_)
        nw377_[int(1)] = d_2013_v1_
        def iife210_(_pat_let67_0):
            def iife211_(d_2015_dt__update__tmp_h0_):
                def iife212_(_pat_let68_0):
                    def iife213_(d_2016_dt__update_hcf85_h0_):
                        return D20_DC49(d_2016_dt__update_hcf85_h0_)
                    return iife213_(_pat_let68_0)
                return iife212_(pat_let_tv43_)
            return iife211_(_pat_let67_0)
        nw377_[int(2)] = iife210_(d_2013_v1_)
        nw377_[int(3)] = d_2013_v1_
        def iife214_(_pat_let69_0):
            def iife215_(d_2017_dt__update__tmp_h1_):
                def iife216_(_pat_let70_0):
                    def iife217_(d_2018_dt__update_hcf85_h1_):
                        return D20_DC49(d_2018_dt__update_hcf85_h1_)
                    return iife217_(_pat_let70_0)
                return iife216_(pat_let_tv44_)
            return iife215_(_pat_let69_0)
        nw377_[int(4)] = iife214_(D20_DC49(d_2012_v0_))
        nw377_[int(5)] = d_2013_v1_
        nw377_[int(6)] = d_2013_v1_
        nw377_[int(7)] = d_2013_v1_
        d_2014_v2_ = nw377_
        index353_ = default__.safeIndex(545, (d_2014_v2_).length(0))
        (d_2014_v2_)[index353_] = d_2013_v1_
        index354_ = default__.safeIndex(545, (d_2014_v2_).length(0))
        rhs277_ = d_2013_v1_
        rhs278_ = default__.fm2(globalState)
        lhs250_ = d_2014_v2_
        lhs251_ = default__.safeIndex(545, (d_2014_v2_).length(0))
        lhs252_ = globalState
        lhs250_[lhs251_] = rhs277_
        lhs252_.f1 = rhs278_
        d_2019_v3_: bool
        d_2019_v3_ = False
        (globalState).f1 = d_2019_v3_
        d_2020_v4_: _dafny.Array
        def lambda94_(d_2021_v3_):
            def lambda95_(d_2022_i0_):
                return d_2021_v3_

            return lambda95_

        init56_ = lambda94_(d_2019_v3_)
        nw378_ = _dafny.Array(None, 3)
        for i0_56_ in range(nw378_.length(0)):
            nw378_[i0_56_] = init56_(i0_56_)
        d_2020_v4_ = nw378_
        d_2023_v5_: _dafny.Array
        d_2023_v5_ = d_2020_v4_
        d_2024_v6_: _dafny.Seq
        d_2024_v6_ = _dafny.SeqWithoutIsStrInference([d_2019_v3_])
        d_2025_v7_: _dafny.MultiSet
        d_2025_v7_ = _dafny.MultiSet([(d_2024_v6_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqjejkkcx"))), len(d_2024_v6_))], d_2019_v3_])
        d_2026_v8_: int
        d_2026_v8_ = -321
        rhs279_ = d_2023_v5_
        rhs280_ = (d_2025_v7_).issubset(d_2025_v7_)
        rhs281_ = d_2026_v8_
        lhs253_ = globalState
        d_2023_v5_ = rhs279_
        r0 = rhs280_
        lhs253_.f6 = rhs281_
        if d_2019_v3_:
            r0 = False
            d_2027_v9_: _dafny.Array
            nw379_ = _dafny.Array(_dafny.CodePoint('D'), 22)
            d_2027_v9_ = nw379_
            index355_ = default__.safeIndex(193, (d_2027_v9_).length(0))
            (d_2027_v9_)[index355_] = p0
            index356_ = default__.safeIndex(193, (d_2027_v9_).length(0))
            (d_2027_v9_)[index356_] = _dafny.CodePoint('x')
            d_2028_v10_: _dafny.Array
            nw380_ = _dafny.Array(_dafny.Seq({}), 1)
            d_2028_v10_ = nw380_
            d_2029_v11_: _dafny.Map
            d_2029_v11_ = _dafny.Map({d_2028_v10_: not(d_2019_v3_)})
            d_2029_v11_ = (d_2029_v11_).set(d_2028_v10_, not (False) or (d_2019_v3_))
            d_2030_v12_: D14
            d_2030_v12_ = D14_DC38(d_2028_v10_)
            pat_let_tv45_ = d_2028_v10_
            def iife218_(_pat_let71_0):
                def iife219_(d_2031_dt__update__tmp_h2_):
                    def iife220_(_pat_let72_0):
                        def iife221_(d_2032_dt__update_hcf67_h0_):
                            return D14_DC38(d_2032_dt__update_hcf67_h0_)
                        return iife221_(_pat_let72_0)
                    return iife220_(pat_let_tv45_)
                return iife219_(_pat_let71_0)
            d_2030_v12_ = iife218_(d_2030_v12_)
            d_2033_v13_: _dafny.Seq
            d_2033_v13_ = _dafny.SeqWithoutIsStrInference([d_2026_v8_, d_2026_v8_, d_2026_v8_])
            d_2034_v14_: _dafny.Seq
            d_2034_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwmvti"))
            d_2035_v15_: _dafny.Array
            def lambda96_(d_2036_v8_):
                def lambda97_(d_2037_i1_):
                    return default__.safeModuloInt(d_2037_i1_, d_2036_v8_)

                return lambda97_

            init57_ = lambda96_(d_2026_v8_)
            nw381_ = _dafny.Array(None, 29)
            for i0_57_ in range(nw381_.length(0)):
                nw381_[i0_57_] = init57_(i0_57_)
            d_2035_v15_ = nw381_
            d_2038_v16_: _dafny.Map
            d_2038_v16_ = _dafny.Map({d_2035_v15_: d_2019_v3_})
            d_2033_v13_ = (d_2033_v13_).set(default__.safeIndex((default__.fm1(True, globalState)) + (len(d_2034_v14_)), len(d_2033_v13_)), default__.fm1(((d_2038_v16_)[d_2035_v15_] if (d_2035_v15_) in (d_2038_v16_) else d_2019_v3_), globalState))
        elif True:
            (globalState).f1 = d_2019_v3_
            d_2039_v17_: str
            d_2039_v17_ = _dafny.CodePoint('t')
            d_2040_v18_: D5
            d_2040_v18_ = D5_DC12(d_2026_v8_, (d_2024_v6_)[default__.safeIndex(d_2026_v8_, len(d_2024_v6_))], p0)
            rhs282_ = len(d_2024_v6_)
            rhs283_ = (d_2040_v18_).cf16
            rhs284_ = d_2019_v3_
            rhs285_ = p0
            lhs254_ = globalState
            lhs255_ = globalState
            lhs254_.f6 = rhs282_
            d_2019_v3_ = rhs283_
            lhs255_.f1 = rhs284_
            d_2039_v17_ = rhs285_
            d_2041_v19_: _dafny.Array
            nw382_ = _dafny.Array(int(0), 22)
            d_2041_v19_ = nw382_
            d_2042_v20_: _dafny.Seq
            d_2042_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbvvji"))
            d_2043_v21_: _dafny.Set
            d_2043_v21_ = _dafny.Set({d_2042_v20_})
            d_2044_v22_: _dafny.Seq
            d_2044_v22_ = _dafny.SeqWithoutIsStrInference([d_2026_v8_])
            d_2045_v23_: _dafny.Seq
            d_2045_v23_ = _dafny.SeqWithoutIsStrInference([len(d_2043_v21_), len(d_2044_v22_)])
            d_2046_v24_: _dafny.Map
            d_2046_v24_ = _dafny.Map({(d_2044_v22_): d_2026_v8_})
            index357_ = default__.safeIndex(639, (d_2041_v19_).length(0))
            (d_2041_v19_)[index357_] = len(d_2046_v24_)
            index358_ = default__.safeIndex(639, (d_2041_v19_).length(0))
            (d_2041_v19_)[index358_] = d_2026_v8_
            d_2047_v25_: _dafny.Map
            d_2047_v25_ = _dafny.Map({d_2019_v3_: d_2042_v20_})
            d_2047_v25_ = (_dafny.Map({d_2019_v3_: _dafny.SeqWithoutIsStrInference([p0 for d_2048_i2_ in range(default__.abs(35))])})) | ((_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spti"))})) | (_dafny.Map({d_2019_v3_: d_2042_v20_})))
            d_2049_v26_: D3
            d_2049_v26_ = D3_DC6(d_2041_v19_)
            d_2050_v27_: _dafny.Seq
            d_2050_v27_ = _dafny.SeqWithoutIsStrInference([d_2049_v26_])
            d_2051_v28_: C0
            nw383_ = C0()
            nw383_.ctor__(d_2050_v27_, False)
            d_2051_v28_ = nw383_
        d_2052_v29_: _dafny.Map
        d_2052_v29_ = _dafny.Map({d_2019_v3_: True})
        d_2053_v30_: _dafny.Array
        def lambda98_(d_2054_v8_):
            def lambda99_(d_2055_i3_):
                return (d_2055_i3_) + (d_2054_v8_)

            return lambda99_

        init58_ = lambda98_(d_2026_v8_)
        nw384_ = _dafny.Array(None, 9)
        for i0_58_ in range(nw384_.length(0)):
            nw384_[i0_58_] = init58_(i0_58_)
        d_2053_v30_ = nw384_
        d_2056_v31_: D3
        d_2056_v31_ = D3_DC6(d_2053_v30_)
        d_2057_v32_: C0
        nw385_ = C0()
        nw385_.ctor__(_dafny.SeqWithoutIsStrInference([d_2056_v31_, d_2056_v31_, d_2056_v31_]), False)
        d_2057_v32_ = nw385_
        d_2058_v33_: _dafny.Map
        d_2058_v33_ = _dafny.Map({d_2057_v32_: d_2026_v8_})
        d_2059_v34_: _dafny.Seq
        d_2059_v34_ = _dafny.SeqWithoutIsStrInference([d_2026_v8_, ((d_2058_v33_)[d_2057_v32_] if (d_2057_v32_) in (d_2058_v33_) else d_2026_v8_)])
        d_2052_v29_ = ((d_2052_v29_) | (d_2052_v29_)).set((d_2026_v8_) in (d_2059_v34_), d_2019_v3_)
        d_2060_v35_: _dafny.Seq
        d_2060_v35_ = _dafny.SeqWithoutIsStrInference([default__.fm19(globalState), d_2024_v6_])
        d_2060_v35_ = d_2060_v35_
        r0 = (d_2057_v32_).f23
        return r0

    def m9(self, p0, globalState):
        r0: int = int(0)
        d_2061_v0_: _dafny.Seq
        d_2061_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twnmp"))
        d_2062_v1_: _dafny.Seq
        d_2062_v1_ = _dafny.SeqWithoutIsStrInference([len(d_2061_v0_)])
        d_2063_v2_: str
        d_2063_v2_ = _dafny.CodePoint('q')
        d_2064_v3_: D3
        d_2064_v3_ = D3_DC7(p0)
        d_2065_v4_: _dafny.Array
        nw386_ = _dafny.Array(False, 7)
        d_2065_v4_ = nw386_
        d_2066_v5_: bool
        d_2066_v5_ = True
        d_2067_v6_: _dafny.Map
        d_2067_v6_ = _dafny.Map({d_2065_v4_: d_2066_v5_})
        d_2068_v7_: D10
        d_2068_v7_ = D10_DC27(len(_dafny.Map({d_2062_v1_: d_2063_v2_})), d_2061_v0_, p0, d_2064_v3_, d_2067_v6_)
        d_2069_v8_: _dafny.Map
        d_2069_v8_ = _dafny.Map({len(d_2061_v0_): p0})
        d_2070_v9_: _dafny.MultiSet
        d_2070_v9_ = _dafny.MultiSet([p0])
        d_2071_v10_: _dafny.Map
        d_2071_v10_ = _dafny.Map({d_2066_v5_: p0})
        d_2072_v11_: _dafny.Array
        nw387_ = _dafny.Array(None, 21)
        nw387_[int(0)] = p0
        nw387_[int(1)] = 800
        nw387_[int(2)] = p0
        nw387_[int(3)] = p0
        nw387_[int(4)] = p0
        nw387_[int(5)] = p0
        nw387_[int(6)] = len(_dafny.Map({p0: p0}))
        nw387_[int(7)] = p0
        nw387_[int(8)] = (d_2068_v7_).cf48
        nw387_[int(9)] = len(d_2069_v8_)
        nw387_[int(10)] = p0
        nw387_[int(11)] = p0
        nw387_[int(12)] = p0
        nw387_[int(13)] = p0
        nw387_[int(14)] = p0
        nw387_[int(15)] = p0
        nw387_[int(16)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_2073_i0_ in range(default__.abs(280))]))
        nw387_[int(17)] = p0
        nw387_[int(18)] = (d_2070_v9_).cardinality
        nw387_[int(19)] = p0
        nw387_[int(20)] = len(d_2071_v10_)
        d_2072_v11_ = nw387_
        d_2074_v12_: D3
        d_2074_v12_ = D3_DC6(d_2072_v11_)
        pat_let_tv46_ = d_2072_v11_
        pat_let_tv47_ = d_2072_v11_
        d_2075_v13_: _dafny.Array
        nw388_ = _dafny.Array(None, 24)
        nw388_[int(0)] = d_2074_v12_
        nw388_[int(1)] = d_2074_v12_
        nw388_[int(2)] = D3_DC6(d_2072_v11_)
        nw388_[int(3)] = d_2074_v12_
        nw388_[int(4)] = d_2074_v12_
        nw388_[int(5)] = d_2074_v12_
        nw388_[int(6)] = d_2074_v12_
        def iife222_(_pat_let73_0):
            def iife223_(d_2076_dt__update__tmp_h0_):
                def iife224_(_pat_let74_0):
                    def iife225_(d_2077_dt__update_hcf6_h0_):
                        return D3_DC6(d_2077_dt__update_hcf6_h0_)
                    return iife225_(_pat_let74_0)
                return iife224_(pat_let_tv46_)
            return iife223_(_pat_let73_0)
        nw388_[int(7)] = iife222_(d_2074_v12_)
        def iife226_(_pat_let75_0):
            def iife227_(d_2078_dt__update__tmp_h1_):
                def iife228_(_pat_let76_0):
                    def iife229_(d_2079_dt__update_hcf6_h1_):
                        return D3_DC6(d_2079_dt__update_hcf6_h1_)
                    return iife229_(_pat_let76_0)
                return iife228_(pat_let_tv47_)
            return iife227_(_pat_let75_0)
        nw388_[int(8)] = iife226_(D3_DC6(d_2072_v11_))
        nw388_[int(9)] = d_2074_v12_
        nw388_[int(10)] = d_2074_v12_
        nw388_[int(11)] = d_2074_v12_
        nw388_[int(12)] = d_2074_v12_
        nw388_[int(13)] = d_2074_v12_
        nw388_[int(14)] = D3_DC6(d_2072_v11_)
        nw388_[int(15)] = D3_DC6(d_2072_v11_)
        nw388_[int(16)] = d_2074_v12_
        nw388_[int(17)] = d_2074_v12_
        nw388_[int(18)] = d_2074_v12_
        nw388_[int(19)] = d_2074_v12_
        nw388_[int(20)] = d_2074_v12_
        nw388_[int(21)] = d_2074_v12_
        nw388_[int(22)] = d_2074_v12_
        nw388_[int(23)] = d_2074_v12_
        d_2075_v13_ = nw388_
        d_2080_v14_: _dafny.Set
        d_2080_v14_ = _dafny.Set({d_2075_v13_})
        d_2081_v15_: D9
        d_2081_v15_ = D9_DC23(d_2080_v14_, d_2061_v0_, d_2066_v5_, d_2072_v11_, d_2063_v2_)
        source20_ = d_2081_v15_
        if source20_.is_DC23:
            d_2082___mcc_h0_ = source20_.cf37
            d_2083___mcc_h1_ = source20_.cf38
            d_2084___mcc_h2_ = source20_.cf39
            d_2085___mcc_h3_ = source20_.cf40
            d_2086___mcc_h4_ = source20_.cf41
            d_2087_cf41_ = d_2086___mcc_h4_
            d_2088_cf40_ = d_2085___mcc_h3_
            d_2089_cf39_ = d_2084___mcc_h2_
            d_2090_cf38_ = d_2083___mcc_h1_
            d_2091_cf37_ = d_2082___mcc_h0_
            d_2092_v16_: _dafny.MultiSet
            d_2092_v16_ = _dafny.MultiSet([d_2066_v5_])
            d_2093_v17_: _dafny.Map
            d_2093_v17_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([default__.fm19(globalState)]): d_2092_v16_})
            d_2094_v18_: _dafny.Seq
            d_2094_v18_ = _dafny.SeqWithoutIsStrInference([False, d_2089_cf39_])
            d_2095_v19_: _dafny.Seq
            d_2095_v19_ = _dafny.SeqWithoutIsStrInference([d_2094_v18_])
            d_2096_v20_: _dafny.Array
            nw389_ = _dafny.Array(None, 1)
            nw389_[int(0)] = ((d_2093_v17_)[d_2095_v19_] if (d_2095_v19_) in (d_2093_v17_) else d_2092_v16_)
            d_2096_v20_ = nw389_
            index359_ = default__.safeIndex(465, (d_2096_v20_).length(0))
            (d_2096_v20_)[index359_] = (d_2092_v16_).intersection(d_2092_v16_)
            index360_ = default__.safeIndex(465, (d_2096_v20_).length(0))
            (d_2096_v20_)[index360_] = _dafny.MultiSet(d_2094_v18_)
            d_2089_cf39_ = (_dafny.CodePoint('v')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esckela")))
            index361_ = default__.safeIndex(228, (d_2065_v4_).length(0))
            (d_2065_v4_)[index361_] = d_2066_v5_
            index362_ = default__.safeIndex(228, (d_2065_v4_).length(0))
            (d_2065_v4_)[index362_] = d_2066_v5_
            d_2097_v21_: _dafny.Array
            def lambda100_(d_2098_v5_):
                def lambda101_(d_2099_i1_):
                    return not(d_2098_v5_)

                return lambda101_

            init59_ = lambda100_(d_2066_v5_)
            nw390_ = _dafny.Array(None, 20)
            for i0_59_ in range(nw390_.length(0)):
                nw390_[i0_59_] = init59_(i0_59_)
            d_2097_v21_ = nw390_
            (globalState).f14 = d_2097_v21_
        elif source20_.is_DC24:
            d_2100___mcc_h5_ = source20_.cf42
            d_2101___mcc_h6_ = source20_.cf43
            d_2102_cf43_ = d_2101___mcc_h6_
            d_2103_cf42_ = d_2100___mcc_h5_
            d_2104_v22_: _dafny.MultiSet
            d_2104_v22_ = _dafny.MultiSet([d_2066_v5_, d_2066_v5_, d_2066_v5_, d_2066_v5_, d_2066_v5_])
            index363_ = default__.safeIndex(279, (d_2072_v11_).length(0))
            (d_2072_v11_)[index363_] = (((d_2104_v22_)[d_2066_v5_] if (d_2066_v5_) in (d_2104_v22_) else p0)) * (p0)
            index364_ = default__.safeIndex(279, (d_2072_v11_).length(0))
            (d_2072_v11_)[index364_] = default__.safeModuloInt(p0, (0) - (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bb"))), (0) - (p0))))
            (globalState).f6 = (0) - ((d_2072_v11_)[default__.safeIndex(279, (d_2072_v11_).length(0))])
            d_2066_v5_ = d_2066_v5_
            d_2105_v23_: _dafny.Map
            d_2105_v23_ = _dafny.Map({(d_2072_v11_)[default__.safeIndex(279, (d_2072_v11_).length(0))]: False})
            d_2106_v24_: _dafny.Map
            d_2106_v24_ = _dafny.Map({d_2105_v23_: (d_2072_v11_)[default__.safeIndex(279, (d_2072_v11_).length(0))]})
            d_2107_v25_: C2
            nw391_ = C2()
            nw391_.ctor__((d_2072_v11_)[default__.safeIndex(279, (d_2072_v11_).length(0))], ((d_2069_v8_)[(d_2072_v11_)[default__.safeIndex(279, (d_2072_v11_).length(0))]] if ((d_2072_v11_)[default__.safeIndex(279, (d_2072_v11_).length(0))]) in (d_2069_v8_) else p0), (d_2106_v24_) | (d_2106_v24_))
            d_2107_v25_ = nw391_
        elif source20_.is_DC22:
            d_2108___mcc_h7_ = source20_.cf36
            d_2109_cf36_ = d_2108___mcc_h7_
            d_2110_v26_: _dafny.Array
            def lambda102_(d_2111_v5_, d_2112_p0_):
                def lambda103_(d_2113_i2_):
                    return (_dafny.Set({d_2112_p0_}) if d_2111_v5_ else _dafny.Set({d_2112_p0_, d_2112_p0_, d_2112_p0_, (0) - (d_2112_p0_), len(_dafny.SeqWithoutIsStrInference([True]))}))

                return lambda103_

            init60_ = lambda102_(d_2066_v5_, p0)
            nw392_ = _dafny.Array(None, 25)
            for i0_60_ in range(nw392_.length(0)):
                nw392_[i0_60_] = init60_(i0_60_)
            d_2110_v26_ = nw392_
            d_2114_v28_: _dafny.Set
            def iife230_():
                coll76_ = _dafny.Set()
                compr_76_: int
                for compr_76_ in _dafny.IntegerRange(271, 842):
                    d_2115_v27_: int = compr_76_
                    if ((271) <= (d_2115_v27_)) and ((d_2115_v27_) < (842)):
                        coll76_ = coll76_.union(_dafny.Set([(d_2115_v27_) - (p0)]))
                return _dafny.Set(coll76_)
            d_2114_v28_ = _dafny.Set({p0, len(iife230_()
            ), len(_dafny.Set({d_2066_v5_, d_2066_v5_})), default__.fm1(d_2066_v5_, globalState)})
            index365_ = default__.safeIndex(438, (d_2110_v26_).length(0))
            (d_2110_v26_)[index365_] = d_2114_v28_
            index366_ = default__.safeIndex(438, (d_2110_v26_).length(0))
            (d_2110_v26_)[index366_] = _dafny.Set({p0})
            d_2116_v29_: _dafny.Set
            d_2116_v29_ = _dafny.Set({d_2061_v0_, d_2061_v0_, d_2061_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkytdei")), (d_2061_v0_).set(default__.safeIndex(393, len(d_2061_v0_)), d_2063_v2_)})
            index367_ = default__.safeIndex(557, (d_2072_v11_).length(0))
            (d_2072_v11_)[index367_] = (p0) - (len(d_2116_v29_))
            index368_ = default__.safeIndex(557, (d_2072_v11_).length(0))
            (d_2072_v11_)[index368_] = (p0) * (p0)
            if ((default__.fm21(globalState)) + (d_2061_v0_)) <= (d_2061_v0_):
                (globalState).f1 = d_2066_v5_
                (globalState).f6 = (d_2072_v11_)[default__.safeIndex(557, (d_2072_v11_).length(0))]
                d_2117_v30_: _dafny.Map
                d_2117_v30_ = _dafny.Map({d_2066_v5_: d_2066_v5_})
                d_2118_v31_: D14
                d_2118_v31_ = D14_DC39(d_2062_v1_, d_2117_v30_, d_2066_v5_)
                d_2119_v32_: D15
                d_2119_v32_ = D15_DC43(d_2066_v5_, (d_2118_v31_).cf70, d_2066_v5_)
                d_2120_v33_: _dafny.MultiSet
                d_2120_v33_ = _dafny.MultiSet([d_2066_v5_])
                d_2121_v34_: _dafny.Map
                d_2121_v34_ = _dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(d_2072_v11_)[default__.safeIndex(557, (d_2072_v11_).length(0))]])): d_2066_v5_}): p0})
                d_2122_v35_: _dafny.Map
                d_2122_v35_ = _dafny.Map({p0: d_2066_v5_})
                d_2123_v36_: C3
                nw393_ = C3()
                nw393_.ctor__((d_2066_v5_) or ((d_2119_v32_).cf74), ((d_2120_v33_).cardinality if d_2066_v5_ else (D5_DC12(len(_dafny.SeqWithoutIsStrInference([d_2063_v2_ for d_2124_i3_ in range(default__.abs(855))])), d_2066_v5_, d_2063_v2_)).cf15), ((d_2121_v34_).set(d_2122_v35_, -637)) | ((d_2121_v34_)))
                d_2123_v36_ = nw393_
                d_2125_v37_: _dafny.Set
                d_2125_v37_ = _dafny.Set({_dafny.CodePoint('p')})
                rhs286_ = ((p0) * ((d_2072_v11_)[default__.safeIndex(557, (d_2072_v11_).length(0))])) - (default__.safeModuloInt((d_2123_v36_).f19, (d_2062_v1_)[default__.safeIndex(((d_2120_v33_)[d_2066_v5_] if (d_2066_v5_) in (d_2120_v33_) else len(d_2067_v6_)), len(d_2062_v1_))]))
                rhs287_ = (len((d_2125_v37_) - (d_2125_v37_))) <= (p0)
                rhs288_ = d_2123_v36_
                lhs256_ = globalState
                lhs256_.f6 = rhs286_
                d_2066_v5_ = rhs287_
                d_2123_v36_ = rhs288_
                d_2126_v38_: _dafny.Array
                nw394_ = _dafny.Array(_dafny.Seq({}), 15)
                d_2126_v38_ = nw394_
                d_2127_v39_: _dafny.Seq
                d_2127_v39_ = _dafny.SeqWithoutIsStrInference([(d_2123_v36_).f18])
                d_2128_v40_: D11
                d_2128_v40_ = D11_DC31(d_2127_v39_, (d_2070_v9_).cardinality, (d_2123_v36_).f19, p0, (d_2123_v36_).f19)
                index369_ = default__.safeIndex(955, (d_2126_v38_).length(0))
                (d_2126_v38_)[index369_] = (d_2128_v40_).cf54
                index370_ = default__.safeIndex(955, (d_2126_v38_).length(0))
                rhs289_ = d_2063_v2_
                rhs290_ = (d_2127_v39_) + (d_2127_v39_)
                rhs291_ = (default__.safeModuloInt(258, (d_2123_v36_).f19)) - (p0)
                lhs257_ = d_2126_v38_
                lhs258_ = default__.safeIndex(955, (d_2126_v38_).length(0))
                d_2063_v2_ = rhs289_
                lhs257_[lhs258_] = rhs290_
                r0 = rhs291_
                (globalState).f1 = d_2066_v5_
            elif True:
                index371_ = default__.safeIndex(557, (d_2072_v11_).length(0))
                (d_2072_v11_)[index371_] = (d_2072_v11_)[default__.safeIndex(557, (d_2072_v11_).length(0))]
                (globalState).f6 = ((d_2072_v11_)[default__.safeIndex(557, (d_2072_v11_).length(0))] if d_2066_v5_ else p0)
                d_2129_v41_: _dafny.Array
                nw395_ = _dafny.Array(D16.default()(), 11)
                d_2129_v41_ = nw395_
                d_2130_v42_: _dafny.Map
                d_2130_v42_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyhai")): d_2129_v41_})
                d_2129_v41_ = ((d_2130_v42_)[d_2061_v0_] if (d_2061_v0_) in (d_2130_v42_) else d_2129_v41_)
                d_2061_v0_ = d_2061_v0_
                d_2131_v43_: C0
                nw396_ = C0()
                nw396_.ctor__((_dafny.SeqWithoutIsStrInference([d_2074_v12_])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([d_2074_v12_]))), d_2074_v12_), d_2066_v5_)
                d_2131_v43_ = nw396_
            rhs292_ = default__.safeModuloInt(418, p0)
            rhs293_ = (d_2061_v0_) + (d_2061_v0_)
            rhs294_ = (len((d_2110_v26_)[default__.safeIndex(438, (d_2110_v26_).length(0))])) == (p0)
            rhs295_ = d_2065_v4_
            lhs259_ = globalState
            lhs260_ = globalState
            lhs261_ = globalState
            lhs259_.f6 = rhs292_
            d_2061_v0_ = rhs293_
            lhs260_.f1 = rhs294_
            lhs261_.f14 = rhs295_
        elif True:
            d_2132___mcc_h8_ = source20_.cf44
            d_2133_cf44_ = d_2132___mcc_h8_
            if (D20_DC50(d_2066_v5_, p0, p0, d_2066_v5_, default__.fm2(globalState))).cf89:
                d_2134_v44_: C1
                nw397_ = C1()
                nw397_.ctor__((p0) * (p0))
                d_2134_v44_ = nw397_
                d_2135_v45_: _dafny.Map
                d_2135_v45_ = _dafny.Map({(d_2134_v44_).f24: default__.fm2(globalState)})
                d_2135_v45_ = (d_2135_v45_).set(p0, d_2066_v5_)
                d_2136_v46_: _dafny.Array
                nw398_ = _dafny.Array(_dafny.CodePoint('D'), 2)
                d_2136_v46_ = nw398_
                index372_ = default__.safeIndex(42, (d_2136_v46_).length(0))
                (d_2136_v46_)[index372_] = d_2063_v2_
                index373_ = default__.safeIndex(42, (d_2136_v46_).length(0))
                (d_2136_v46_)[index373_] = d_2063_v2_
                index374_ = default__.safeIndex(980, (d_2065_v4_).length(0))
                (d_2065_v4_)[index374_] = d_2066_v5_
                d_2137_v47_: _dafny.Map
                d_2137_v47_ = _dafny.Map({(d_2070_v9_).cardinality: _dafny.Set({619, -1})})
                d_2138_v48_: _dafny.Set
                d_2138_v48_ = _dafny.Set({p0})
                d_2139_v49_: _dafny.Map
                d_2139_v49_ = _dafny.Map({(d_2134_v44_).f24: d_2138_v48_})
                d_2140_v50_: _dafny.Set
                d_2140_v50_ = _dafny.Set({d_2066_v5_})
                d_2141_v51_: D0
                d_2141_v51_ = D0_DC1(len(_dafny.SeqWithoutIsStrInference([d_2066_v5_])))
                d_2142_v52_: D0
                d_2142_v52_ = D0_DC2(d_2141_v51_)
                d_2143_v53_: _dafny.Array
                nw399_ = _dafny.Array(None, 29)
                nw399_[int(0)] = d_2137_v47_
                nw399_[int(1)] = d_2139_v49_
                nw399_[int(2)] = d_2139_v49_
                nw399_[int(3)] = d_2139_v49_
                nw399_[int(4)] = d_2137_v47_
                nw399_[int(5)] = d_2139_v49_
                nw399_[int(6)] = d_2139_v49_
                nw399_[int(7)] = d_2139_v49_
                nw399_[int(8)] = _dafny.Map({len(d_2135_v45_): _dafny.Set({p0, len(d_2140_v50_)})})
                nw399_[int(9)] = d_2139_v49_
                nw399_[int(10)] = d_2137_v47_
                nw399_[int(11)] = d_2139_v49_
                nw399_[int(12)] = default__.fm26((d_2136_v46_)[default__.safeIndex(42, (d_2136_v46_).length(0))], d_2142_v52_, d_2066_v5_, d_2066_v5_, globalState)
                nw399_[int(13)] = d_2139_v49_
                nw399_[int(14)] = d_2137_v47_
                nw399_[int(15)] = _dafny.Map({(self).fm9(_dafny.SeqWithoutIsStrInference([d_2063_v2_ for d_2144_i4_ in range(default__.abs(601))]), globalState): _dafny.Set({(d_2134_v44_).f24, p0})})
                nw399_[int(16)] = d_2139_v49_
                nw399_[int(17)] = d_2139_v49_
                nw399_[int(18)] = d_2139_v49_
                nw399_[int(19)] = d_2139_v49_
                nw399_[int(20)] = d_2139_v49_
                nw399_[int(21)] = d_2139_v49_
                nw399_[int(22)] = d_2139_v49_
                nw399_[int(23)] = d_2139_v49_
                nw399_[int(24)] = d_2139_v49_
                nw399_[int(25)] = d_2139_v49_
                nw399_[int(26)] = _dafny.Map({len(d_2062_v1_): d_2138_v48_})
                nw399_[int(27)] = d_2139_v49_
                nw399_[int(28)] = d_2139_v49_
                d_2143_v53_ = nw399_
                d_2145_v54_: _dafny.Array
                nw400_ = _dafny.Array(None, 5)
                nw400_[int(0)] = d_2143_v53_
                nw400_[int(1)] = d_2143_v53_
                nw400_[int(2)] = d_2143_v53_
                nw400_[int(3)] = d_2143_v53_
                nw400_[int(4)] = d_2143_v53_
                d_2145_v54_ = nw400_
                index375_ = default__.safeIndex(980, (d_2065_v4_).length(0))
                rhs296_ = not((d_2066_v5_) or (d_2066_v5_))
                rhs297_ = d_2145_v54_
                lhs262_ = d_2065_v4_
                lhs263_ = default__.safeIndex(980, (d_2065_v4_).length(0))
                lhs262_[lhs263_] = rhs296_
                d_2145_v54_ = rhs297_
                d_2146_v55_: C5
                nw401_ = C5()
                nw401_.ctor__(default__.fm1(d_2066_v5_, globalState))
                d_2146_v55_ = nw401_
            elif True:
                (globalState).f6 = p0
                d_2147_v56_: D0
                d_2147_v56_ = D0_DC0(True)
                d_2148_v57_: _dafny.Map
                d_2148_v57_ = _dafny.Map({d_2066_v5_: d_2066_v5_})
                d_2149_v58_: _dafny.Map
                d_2149_v58_ = _dafny.Map({p0: d_2063_v2_})
                d_2150_v59_: D3
                d_2150_v59_ = D3_DC8(d_2062_v1_, _dafny.SeqWithoutIsStrInference([d_2063_v2_ for d_2151_i7_ in range(default__.abs(193))]), d_2066_v5_, _dafny.CodePoint('n'))
                d_2152_v60_: _dafny.Array
                nw402_ = _dafny.Array(None, 21)
                nw402_[int(0)] = d_2061_v0_
                nw402_[int(1)] = (_dafny.SeqWithoutIsStrInference([d_2063_v2_ for d_2153_i5_ in range(default__.abs(204))])) + (_dafny.SeqWithoutIsStrInference([d_2063_v2_ for d_2154_i6_ in range(default__.abs(374))]))
                nw402_[int(2)] = d_2061_v0_
                nw402_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcf"))
                nw402_[int(4)] = d_2061_v0_
                nw402_[int(5)] = d_2061_v0_
                nw402_[int(6)] = d_2061_v0_
                nw402_[int(7)] = d_2061_v0_
                nw402_[int(8)] = d_2061_v0_
                nw402_[int(9)] = d_2061_v0_
                nw402_[int(10)] = (d_2061_v0_ if (self).fm10(d_2147_v56_, d_2148_v57_, ((d_2149_v58_)[p0] if (p0) in (d_2149_v58_) else _dafny.CodePoint('p')), p0, globalState) else d_2061_v0_)
                nw402_[int(11)] = (d_2150_v59_).cf9
                nw402_[int(12)] = d_2061_v0_
                nw402_[int(13)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fmsrvacn"))) + (d_2061_v0_)
                nw402_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "timqowl"))
                nw402_[int(15)] = d_2061_v0_
                nw402_[int(16)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))) + (d_2061_v0_)
                nw402_[int(17)] = d_2061_v0_
                nw402_[int(18)] = d_2061_v0_
                nw402_[int(19)] = d_2061_v0_
                nw402_[int(20)] = d_2061_v0_
                d_2152_v60_ = nw402_
                d_2155_v61_: D8
                d_2155_v61_ = D8_DC21(d_2061_v0_, _dafny.Set({d_2066_v5_, d_2066_v5_}))
                index376_ = default__.safeIndex(216, (d_2152_v60_).length(0))
                (d_2152_v60_)[index376_] = (d_2155_v61_).cf34
                d_2156_v62_: _dafny.Seq
                d_2156_v62_ = _dafny.SeqWithoutIsStrInference([(d_2061_v0_).set(default__.safeIndex(p0, len(d_2061_v0_)), d_2063_v2_), d_2061_v0_, (d_2061_v0_) + (d_2061_v0_), d_2061_v0_])
                index377_ = default__.safeIndex(216, (d_2152_v60_).length(0))
                (d_2152_v60_)[index377_] = (d_2156_v62_)[default__.safeIndex(p0, len(d_2156_v62_))]
                d_2066_v5_ = (p0) == ((p0) * (p0))
                d_2157_v63_: _dafny.Set
                d_2157_v63_ = _dafny.Set({d_2066_v5_})
                (globalState).f6 = len((d_2157_v63_) | ((d_2157_v63_).intersection(default__.fm31((0) - (p0), d_2066_v5_, globalState))))
                index378_ = default__.safeIndex(50, (d_2065_v4_).length(0))
                (d_2065_v4_)[index378_] = d_2066_v5_
                index379_ = default__.safeIndex(50, (d_2065_v4_).length(0))
                (d_2065_v4_)[index379_] = d_2066_v5_
            d_2158_v64_: C4
            nw403_ = C4()
            nw403_.ctor__()
            d_2158_v64_ = nw403_
            d_2159_v65_: C5
            nw404_ = C5()
            nw404_.ctor__(p0)
            d_2159_v65_ = nw404_
            d_2160_v66_: D23
            d_2160_v66_ = D23_DC53(d_2159_v65_)
            d_2161_v67_: _dafny.Map
            d_2161_v67_ = _dafny.Map({(len(_dafny.SeqWithoutIsStrInference([d_2063_v2_ for d_2162_i8_ in range(default__.abs(335))]))) * (default__.fm1(d_2066_v5_, globalState)): (d_2160_v66_).cf93})
            d_2161_v67_ = (d_2161_v67_).set(p0, d_2159_v65_)
            d_2163_v68_: _dafny.MultiSet
            d_2163_v68_ = _dafny.MultiSet([d_2066_v5_, d_2066_v5_, d_2066_v5_, d_2066_v5_, not(d_2066_v5_)])
            rhs298_ = (d_2066_v5_) and (((0) - (p0)) < (p0))
            rhs299_ = d_2065_v4_
            rhs300_ = (0) - ((d_2159_v65_).f16)
            rhs301_ = d_2061_v0_
            rhs302_ = not(not(not(((d_2163_v68_) - ((d_2159_v65_).fm11(18, globalState))).ispropersubset(d_2163_v68_))))
            lhs264_ = globalState
            lhs265_ = globalState
            lhs266_ = globalState
            lhs264_.f1 = rhs298_
            lhs265_.f14 = rhs299_
            lhs266_.f6 = rhs300_
            d_2061_v0_ = rhs301_
            d_2066_v5_ = rhs302_
        d_2072_v11_ = d_2072_v11_
        d_2164_v69_: D14
        d_2164_v69_ = D14_DC40(default__.fm2(globalState))
        pat_let_tv48_ = d_2066_v5_
        def iife231_(_pat_let77_0):
            def iife232_(d_2165_dt__update__tmp_h5_):
                def iife233_(_pat_let78_0):
                    def iife234_(d_2166_dt__update_hcf71_h0_):
                        return D14_DC40(d_2166_dt__update_hcf71_h0_)
                    return iife234_(_pat_let78_0)
                return iife233_(pat_let_tv48_)
            return iife232_(_pat_let77_0)
        d_2164_v69_ = iife231_(d_2164_v69_)
        d_2066_v5_ = d_2066_v5_
        d_2167_v71_: C3
        nw405_ = C3()
        def iife235_():
            coll77_ = _dafny.Map()
            compr_77_: _dafny.Map
            for compr_77_ in (default__.fm45(p0, globalState)).Elements:
                d_2168_v70_: _dafny.Map = compr_77_
                if (d_2168_v70_) in (default__.fm45(p0, globalState)):
                    coll77_[d_2168_v70_] = p0
            return _dafny.Map(coll77_)
        nw405_.ctor__((p0) in (d_2070_v9_), default__.safeModuloInt(p0, p0), iife235_()
        )
        d_2167_v71_ = nw405_
        hi10_ = (d_2167_v71_).f19
        for d_2169_i9_ in range(len((d_2061_v0_) + (d_2061_v0_)), hi10_):
            (globalState).f14 = d_2065_v4_
            d_2066_v5_ = (d_2167_v71_).f18
            d_2066_v5_ = d_2066_v5_
            (globalState).f1 = (((d_2167_v71_).f19) > (default__.fm1(d_2066_v5_, globalState)) if ((d_2167_v71_).f19) < ((d_2167_v71_).f19) else ((d_2167_v71_).f18) and (not((d_2167_v71_).f18)))
        r0 = ((p0) - ((0) - (p0))) - ((d_2167_v71_).f19)
        return r0


class C7(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self):
        pass
        pass

    def fm3(self, p0, p1, p2, p3, globalState):
        if (not(False)) not in (_dafny.Set({False, not(False), not(True), True})):
            return D0_DC2(D0_DC0(True))
        elif True:
            return D0_DC2(D0_DC0((D0_DC0(True)).cf0))

    def fm8(self, p0, p1, p2, globalState):
        if False:
            return not(True)
        elif True:
            return False

    def m1(self, globalState):
        r0: int = int(0)
        r1: bool = False
        d_2170_v0_: bool
        d_2170_v0_ = True
        if not(((True) == (True) if d_2170_v0_ else d_2170_v0_)):
            d_2171_v1_: _dafny.MultiSet
            d_2171_v1_ = _dafny.MultiSet([not(d_2170_v0_), d_2170_v0_, d_2170_v0_])
            (globalState).f6 = default__.fm1((d_2171_v1_).ispropersubset(d_2171_v1_), globalState)
            (self).m7(globalState)
            if d_2170_v0_:
                d_2172_v2_: int
                d_2172_v2_ = 834
                (globalState).f6 = default__.safeModuloInt(d_2172_v2_, d_2172_v2_)
                d_2173_v3_: _dafny.Seq
                d_2173_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjdokiwra"))
                d_2174_v4_: _dafny.Array
                def lambda104_(d_2175_v0_):
                    def lambda105_(d_2176_i0_):
                        return d_2175_v0_

                    return lambda105_

                init61_ = lambda104_(d_2170_v0_)
                nw406_ = _dafny.Array(None, 20)
                for i0_61_ in range(nw406_.length(0)):
                    nw406_[i0_61_] = init61_(i0_61_)
                d_2174_v4_ = nw406_
                d_2177_v5_: _dafny.Seq
                d_2177_v5_ = _dafny.SeqWithoutIsStrInference([d_2170_v0_, d_2170_v0_])
                index380_ = default__.safeIndex(882, (d_2174_v4_).length(0))
                (d_2174_v4_)[index380_] = (d_2177_v5_)[default__.safeIndex(d_2172_v2_, len(d_2177_v5_))]
                d_2178_v6_: _dafny.Seq
                d_2178_v6_ = _dafny.SeqWithoutIsStrInference([d_2172_v2_, d_2172_v2_])
                index381_ = default__.safeIndex(882, (d_2174_v4_).length(0))
                rhs303_ = d_2173_v3_
                rhs304_ = (d_2178_v6_) <= (d_2178_v6_)
                lhs267_ = d_2174_v4_
                lhs268_ = default__.safeIndex(882, (d_2174_v4_).length(0))
                d_2173_v3_ = rhs303_
                lhs267_[lhs268_] = rhs304_
                d_2179_v7_: C4
                nw407_ = C4()
                nw407_.ctor__()
                d_2179_v7_ = nw407_
                d_2180_v8_: C6
                nw408_ = C6()
                nw408_.ctor__()
                d_2180_v8_ = nw408_
                d_2181_v9_: str
                d_2181_v9_ = _dafny.CodePoint('i')
                d_2182_v10_: D5
                d_2182_v10_ = D5_DC11(d_2172_v2_)
                d_2183_v11_: _dafny.Seq
                d_2183_v11_ = _dafny.SeqWithoutIsStrInference([d_2182_v10_])
                d_2184_v12_: _dafny.MultiSet
                d_2184_v12_ = _dafny.MultiSet([d_2172_v2_])
                d_2185_v13_: _dafny.Set
                d_2185_v13_ = _dafny.Set({d_2170_v0_, d_2170_v0_})
                d_2186_v14_: _dafny.Map
                d_2186_v14_ = _dafny.Map({len(d_2173_v3_): len(d_2178_v6_)})
                d_2187_v15_: _dafny.Array
                nw409_ = _dafny.Array(None, 10)
                nw409_[int(0)] = (len(d_2183_v11_)) - (d_2172_v2_)
                nw409_[int(1)] = d_2172_v2_
                nw409_[int(2)] = d_2172_v2_
                nw409_[int(3)] = d_2172_v2_
                nw409_[int(4)] = d_2172_v2_
                nw409_[int(5)] = d_2172_v2_
                nw409_[int(6)] = d_2172_v2_
                nw409_[int(7)] = (((d_2184_v12_)[d_2172_v2_] if (d_2172_v2_) in (d_2184_v12_) else len(d_2178_v6_))) * (d_2172_v2_)
                nw409_[int(8)] = len(d_2185_v13_)
                nw409_[int(9)] = (len(d_2186_v14_)) * (d_2172_v2_)
                d_2187_v15_ = nw409_
                d_2188_v16_: _dafny.Seq
                d_2188_v16_ = _dafny.SeqWithoutIsStrInference([d_2173_v3_])
                d_2189_v17_: _dafny.Map
                d_2189_v17_ = _dafny.Map({len(d_2188_v16_): False})
                index382_ = default__.safeIndex(235, (d_2187_v15_).length(0))
                (d_2187_v15_)[index382_] = len((d_2189_v17_) | (d_2189_v17_))
                d_2190_v18_: _dafny.Map
                d_2190_v18_ = _dafny.Map({d_2170_v0_: d_2170_v0_})
                index383_ = default__.safeIndex(235, (d_2187_v15_).length(0))
                rhs305_ = default__.fm29((d_2174_v4_)[default__.safeIndex(882, (d_2174_v4_).length(0))], d_2172_v2_, d_2172_v2_, ((d_2190_v18_)[d_2170_v0_] if (d_2170_v0_) in (d_2190_v18_) else d_2170_v0_), globalState)
                rhs306_ = 16
                lhs269_ = d_2187_v15_
                lhs270_ = default__.safeIndex(235, (d_2187_v15_).length(0))
                d_2181_v9_ = rhs305_
                lhs269_[lhs270_] = rhs306_
            elif True:
                d_2191_v19_: int
                d_2191_v19_ = 990
                d_2192_v20_: _dafny.Seq
                d_2192_v20_ = _dafny.SeqWithoutIsStrInference([d_2191_v19_])
                d_2193_v21_: str
                d_2193_v21_ = _dafny.CodePoint('d')
                d_2194_v22_: D3
                d_2194_v22_ = D3_DC8(d_2192_v20_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgrydx")), d_2170_v0_, d_2193_v21_)
                d_2195_v23_: _dafny.Seq
                d_2195_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okujifd"))
                pat_let_tv49_ = d_2195_v23_
                pat_let_tv50_ = d_2195_v23_
                pat_let_tv51_ = d_2191_v19_
                pat_let_tv52_ = d_2195_v23_
                def iife236_(_pat_let79_0):
                    def iife237_(d_2196_dt__update__tmp_h0_):
                        def iife238_(_pat_let80_0):
                            def iife239_(d_2197_dt__update_hcf9_h0_):
                                def iife240_(_pat_let81_0):
                                    def iife241_(d_2198_dt__update_hcf11_h0_):
                                        return D3_DC8((d_2196_dt__update__tmp_h0_).cf8, d_2197_dt__update_hcf9_h0_, (d_2196_dt__update__tmp_h0_).cf10, d_2198_dt__update_hcf11_h0_)
                                    return iife241_(_pat_let81_0)
                                return iife240_((pat_let_tv50_)[default__.safeIndex(pat_let_tv51_, len(pat_let_tv52_))])
                            return iife239_(_pat_let80_0)
                        return iife238_(pat_let_tv49_)
                    return iife237_(_pat_let79_0)
                d_2170_v0_ = not(not((iife236_(d_2194_v22_)).cf10))
                d_2199_v24_: D13
                d_2199_v24_ = D13_DC37(not (d_2170_v0_) or (d_2170_v0_), d_2170_v0_)
                d_2199_v24_ = d_2199_v24_
                d_2200_v25_: _dafny.Array
                def lambda106_(d_2201_v23_):
                    def lambda107_(d_2202_i1_):
                        return d_2201_v23_

                    return lambda107_

                init62_ = lambda106_(d_2195_v23_)
                nw410_ = _dafny.Array(None, 5)
                for i0_62_ in range(nw410_.length(0)):
                    nw410_[i0_62_] = init62_(i0_62_)
                d_2200_v25_ = nw410_
                index384_ = default__.safeIndex(24, (d_2200_v25_).length(0))
                (d_2200_v25_)[index384_] = (d_2195_v23_) + (d_2195_v23_)
                d_2203_v26_: _dafny.Array
                nw411_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_2203_v26_ = nw411_
                d_2204_v27_: _dafny.Array
                nw412_ = _dafny.Array(False, 29)
                d_2204_v27_ = nw412_
                index385_ = default__.safeIndex(939, (d_2203_v26_).length(0))
                (d_2203_v26_)[index385_] = d_2204_v27_
                d_2205_v28_: _dafny.Seq
                d_2205_v28_ = _dafny.SeqWithoutIsStrInference([False])
                index386_ = default__.safeIndex(24, (d_2200_v25_).length(0))
                index387_ = default__.safeIndex(939, (d_2203_v26_).length(0))
                rhs307_ = d_2191_v19_
                rhs308_ = (d_2195_v23_) + ((d_2195_v23_) + (d_2195_v23_))
                rhs309_ = ((0) - ((d_2191_v19_) - (d_2191_v19_)) if d_2170_v0_ else len((d_2205_v28_) + (d_2205_v28_)))
                rhs310_ = d_2204_v27_
                lhs271_ = d_2200_v25_
                lhs272_ = default__.safeIndex(24, (d_2200_v25_).length(0))
                lhs273_ = d_2203_v26_
                lhs274_ = default__.safeIndex(939, (d_2203_v26_).length(0))
                r0 = rhs307_
                lhs271_[lhs272_] = rhs308_
                d_2191_v19_ = rhs309_
                lhs273_[lhs274_] = rhs310_
                (globalState).f6 = (110) * ((d_2191_v19_) + (d_2191_v19_))
                d_2206_v29_: D0
                d_2206_v29_ = D0_DC0(d_2170_v0_)
                (globalState).f1 = (d_2206_v29_).cf0
            d_2207_v30_: _dafny.Array
            def lambda108_(d_2208_v0_):
                def lambda109_(d_2209_i2_):
                    return d_2208_v0_

                return lambda109_

            init63_ = lambda108_(d_2170_v0_)
            nw413_ = _dafny.Array(None, 4)
            for i0_63_ in range(nw413_.length(0)):
                nw413_[i0_63_] = init63_(i0_63_)
            d_2207_v30_ = nw413_
            d_2210_v31_: int
            d_2210_v31_ = 605
            d_2211_v32_: _dafny.Seq
            d_2211_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krqgpb"))
            d_2212_v34_: _dafny.Map
            d_2212_v34_ = _dafny.Map({d_2210_v31_: d_2170_v0_})
            d_2213_v35_: _dafny.MultiSet
            d_2213_v35_ = _dafny.MultiSet([d_2212_v34_])
            d_2214_v36_: C2
            nw414_ = C2()
            def iife242_():
                coll78_ = _dafny.Map()
                compr_78_: _dafny.Map
                for compr_78_ in (d_2213_v35_).Elements:
                    d_2215_v33_: _dafny.Map = compr_78_
                    if (d_2215_v33_) in (d_2213_v35_):
                        coll78_[d_2215_v33_] = d_2210_v31_
                return _dafny.Map(coll78_)
            nw414_.ctor__(d_2210_v31_, len(d_2211_v32_), iife242_()
            )
            d_2214_v36_ = nw414_
            d_2216_v37_: _dafny.Map
            d_2216_v37_ = _dafny.Map({d_2214_v36_: d_2170_v0_})
            d_2217_v38_: _dafny.Map
            d_2217_v38_ = _dafny.Map({d_2170_v0_: d_2207_v30_})
            d_2218_v39_: D0
            d_2218_v39_ = D0_DC0(d_2170_v0_)
            d_2219_v40_: _dafny.Array
            d_2219_v40_ = d_2207_v30_
            d_2220_v41_: _dafny.Array
            nw415_ = _dafny.Array(None, 12)
            nw415_[int(0)] = d_2207_v30_
            nw415_[int(1)] = (d_2207_v30_ if ((d_2216_v37_)[d_2214_v36_] if (d_2214_v36_) in (d_2216_v37_) else d_2170_v0_) else d_2207_v30_)
            nw415_[int(2)] = ((d_2217_v38_)[(d_2214_v36_).fm16(d_2218_v39_, d_2170_v0_, globalState)] if ((d_2214_v36_).fm16(d_2218_v39_, d_2170_v0_, globalState)) in (d_2217_v38_) else d_2207_v30_)
            nw415_[int(3)] = d_2207_v30_
            nw415_[int(4)] = (d_2219_v40_)
            nw415_[int(5)] = d_2207_v30_
            nw415_[int(6)] = d_2207_v30_
            nw415_[int(7)] = d_2207_v30_
            nw415_[int(8)] = d_2207_v30_
            nw415_[int(9)] = d_2207_v30_
            nw415_[int(10)] = d_2207_v30_
            nw415_[int(11)] = d_2207_v30_
            d_2220_v41_ = nw415_
            d_2221_v42_: _dafny.Array
            nw416_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
            d_2221_v42_ = nw416_
            d_2222_v43_: _dafny.Seq
            d_2222_v43_ = _dafny.SeqWithoutIsStrInference([d_2211_v32_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddvwnpi")), d_2211_v32_])
            d_2223_v44_: str
            d_2223_v44_ = _dafny.CodePoint('b')
            index388_ = default__.safeIndex(803, (d_2221_v42_).length(0))
            (d_2221_v42_)[index388_] = (d_2222_v43_)[default__.safeIndex(len(default__.fm13(d_2170_v0_, d_2223_v44_, d_2210_v31_, globalState)), len(d_2222_v43_))]
            d_2224_v45_: _dafny.Seq
            d_2224_v45_ = _dafny.SeqWithoutIsStrInference([not (d_2170_v0_) or (default__.fm2(globalState)), d_2170_v0_, False, d_2170_v0_, d_2170_v0_])
            index389_ = default__.safeIndex(803, (d_2221_v42_).length(0))
            rhs311_ = (d_2224_v45_)[default__.safeIndex((d_2214_v36_).f20, len(d_2224_v45_))]
            rhs312_ = d_2220_v41_
            rhs313_ = (d_2211_v32_) + (d_2211_v32_)
            rhs314_ = default__.fm2(globalState)
            lhs275_ = d_2221_v42_
            lhs276_ = default__.safeIndex(803, (d_2221_v42_).length(0))
            lhs277_ = globalState
            r1 = rhs311_
            d_2220_v41_ = rhs312_
            lhs275_[lhs276_] = rhs313_
            lhs277_.f1 = rhs314_
            if ((d_2214_v36_.f21) <= (default__.fm1(d_2170_v0_, globalState))) == ((not(d_2170_v0_) if d_2170_v0_ else d_2170_v0_)):
                d_2218_v39_ = d_2218_v39_
                d_2225_v46_: _dafny.Array
                nw417_ = _dafny.Array(int(0), 21)
                d_2225_v46_ = nw417_
                d_2226_v47_: D3
                d_2226_v47_ = D3_DC6(d_2225_v46_)
                d_2227_v48_: _dafny.Seq
                d_2227_v48_ = _dafny.SeqWithoutIsStrInference([d_2226_v47_, d_2226_v47_])
                d_2228_v49_: D11
                d_2228_v49_ = D11_DC30(d_2227_v48_)
                d_2229_v50_: _dafny.Map
                d_2229_v50_ = _dafny.Map({(d_2214_v36_).f20: (d_2214_v36_).f20})
                d_2230_v51_: _dafny.Map
                d_2230_v51_ = _dafny.Map({d_2228_v49_: d_2229_v50_})
                d_2230_v51_ = (d_2230_v51_).set(d_2228_v49_, d_2229_v50_)
                d_2222_v43_ = d_2222_v43_
                (globalState).f1 = not(d_2170_v0_)
                (globalState).f14 = d_2207_v30_
            elif True:
                d_2231_v52_: D14
                d_2231_v52_ = D14_DC40(d_2170_v0_)
                d_2231_v52_ = D14_DC40(d_2170_v0_)
                d_2232_v53_: _dafny.Array
                def lambda110_(d_2233_v34_, d_2234_v36_):
                    def lambda111_(d_2235_i3_):
                        return _dafny.Map({d_2233_v34_: (d_2234_v36_).f20})

                    return lambda111_

                init64_ = lambda110_(d_2212_v34_, d_2214_v36_)
                nw418_ = _dafny.Array(None, 23)
                for i0_64_ in range(nw418_.length(0)):
                    nw418_[i0_64_] = init64_(i0_64_)
                d_2232_v53_ = nw418_
                d_2232_v53_ = d_2232_v53_
                d_2236_v54_: _dafny.Map
                d_2236_v54_ = _dafny.Map({((d_2214_v36_).f20) + (d_2210_v31_): len((d_2221_v42_)[default__.safeIndex(803, (d_2221_v42_).length(0))])})
                d_2237_v55_: D3
                d_2237_v55_ = D3_DC7(d_2210_v31_)
                rhs315_ = ((d_2236_v54_)[(d_2237_v55_).cf7] if ((d_2237_v55_).cf7) in (d_2236_v54_) else (0) - ((0) - (d_2210_v31_)))
                rhs316_ = not(d_2170_v0_)
                rhs317_ = (0) - ((0) - ((((0) - (d_2214_v36_.f21)) - (d_2214_v36_.f21)) + (59)))
                lhs278_ = globalState
                lhs279_ = d_2214_v36_
                lhs278_.f6 = rhs315_
                r1 = rhs316_
                lhs279_.f21 = rhs317_
                (globalState).f1 = True
                index390_ = default__.safeIndex(669, (d_2207_v30_).length(0))
                (d_2207_v30_)[index390_] = False
                index391_ = default__.safeIndex(669, (d_2207_v30_).length(0))
                rhs318_ = False
                rhs319_ = default__.fm1((d_2218_v39_).cf0, globalState)
                lhs280_ = d_2207_v30_
                lhs281_ = default__.safeIndex(669, (d_2207_v30_).length(0))
                lhs282_ = globalState
                lhs280_[lhs281_] = rhs318_
                lhs282_.f6 = rhs319_
        elif True:
            d_2238_v56_: _dafny.Array
            nw419_ = _dafny.Array(None, 12)
            nw419_[int(0)] = d_2170_v0_
            nw419_[int(1)] = d_2170_v0_
            nw419_[int(2)] = d_2170_v0_
            nw419_[int(3)] = d_2170_v0_
            nw419_[int(4)] = d_2170_v0_
            nw419_[int(5)] = d_2170_v0_
            nw419_[int(6)] = d_2170_v0_
            nw419_[int(7)] = default__.fm2(globalState)
            nw419_[int(8)] = d_2170_v0_
            nw419_[int(9)] = d_2170_v0_
            nw419_[int(10)] = d_2170_v0_
            nw419_[int(11)] = d_2170_v0_
            d_2238_v56_ = nw419_
            d_2239_v57_: _dafny.Seq
            d_2239_v57_ = _dafny.SeqWithoutIsStrInference([d_2238_v56_, d_2238_v56_])
            d_2240_v58_: _dafny.Map
            d_2240_v58_ = _dafny.Map({d_2170_v0_: len(d_2239_v57_)})
            d_2241_v59_: C1
            nw420_ = C1()
            nw420_.ctor__(default__.fm1(d_2170_v0_, globalState))
            d_2241_v59_ = nw420_
            d_2242_v60_: _dafny.Map
            d_2242_v60_ = _dafny.Map({d_2170_v0_: d_2241_v59_})
            d_2240_v58_ = (d_2240_v58_).set(d_2170_v0_, default__.safeModuloInt(len(d_2242_v60_), (0) - ((d_2241_v59_).f24)))
            d_2243_v61_: _dafny.Seq
            d_2243_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqob"))
            d_2243_v61_ = (d_2243_v61_ if d_2170_v0_ else d_2243_v61_)
            d_2244_v62_: _dafny.Array
            nw421_ = _dafny.Array(D16.default()(), 9)
            d_2244_v62_ = nw421_
            d_2245_v63_: str
            d_2245_v63_ = _dafny.CodePoint('s')
            index392_ = default__.safeIndex(436, (d_2244_v62_).length(0))
            (d_2244_v62_)[index392_] = D16_DC45((d_2241_v59_).f24, d_2245_v63_, len(d_2243_v61_), d_2170_v0_)
            d_2246_v64_: D16
            d_2246_v64_ = D16_DC45(len(d_2243_v61_), d_2245_v63_, (d_2241_v59_).f24, False)
            d_2247_v65_: _dafny.Map
            d_2247_v65_ = _dafny.Map({d_2170_v0_: d_2245_v63_})
            pat_let_tv53_ = d_2241_v59_
            pat_let_tv54_ = d_2170_v0_
            pat_let_tv55_ = d_2247_v65_
            pat_let_tv56_ = d_2170_v0_
            pat_let_tv57_ = globalState
            index393_ = default__.safeIndex(436, (d_2244_v62_).length(0))
            def iife244_(_pat_let83_0):
                def iife245_(d_2248_dt__update__tmp_h1_):
                    def iife246_(_pat_let84_0):
                        def iife247_(d_2249_dt__update_hcf80_h0_):
                            return D16_DC45((d_2248_dt__update__tmp_h1_).cf78, (d_2248_dt__update__tmp_h1_).cf79, d_2249_dt__update_hcf80_h0_, (d_2248_dt__update__tmp_h1_).cf81)
                        return iife247_(_pat_let84_0)
                    return iife246_((pat_let_tv53_).f24)
                return iife245_(_pat_let83_0)
            def iife243_(_pat_let82_0):
                def iife248_(d_2250_dt__update__tmp_h2_):
                    def iife249_(_pat_let85_0):
                        def iife250_(d_2251_dt__update_hcf81_h0_):
                            def iife251_(_pat_let86_0):
                                def iife252_(d_2252_dt__update_hcf78_h0_):
                                    return D16_DC45(d_2252_dt__update_hcf78_h0_, (d_2250_dt__update__tmp_h2_).cf79, (d_2250_dt__update__tmp_h2_).cf80, d_2251_dt__update_hcf81_h0_)
                                return iife252_(_pat_let86_0)
                            return iife251_((len(pat_let_tv55_)) - (default__.fm1(pat_let_tv56_, pat_let_tv57_)))
                        return iife250_(_pat_let85_0)
                    return iife249_(pat_let_tv54_)
                return iife248_(_pat_let82_0)
            (d_2244_v62_)[index393_] = iife243_(iife244_(d_2246_v64_))
            (globalState).f1 = not(d_2170_v0_)
            d_2253_v66_: _dafny.Array
            nw422_ = _dafny.Array(_dafny.Map({}), 19)
            d_2253_v66_ = nw422_
            d_2254_v67_: D9
            d_2254_v67_ = D9_DC22(d_2253_v66_)
            d_2255_v68_: _dafny.Map
            d_2255_v68_ = _dafny.Map({d_2254_v67_: d_2238_v56_})
            d_2255_v68_ = d_2255_v68_
        r1 = not(d_2170_v0_)
        d_2256_v69_: _dafny.Array
        nw423_ = _dafny.Array(int(0), 23)
        d_2256_v69_ = nw423_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_2256_v69_).length(0)):
            d_2257_i4_: int = guard_loop_6_
            if (True) and (((0) <= (d_2257_i4_)) and ((d_2257_i4_) < ((d_2256_v69_).length(0)))):
                (d_2256_v69_)[(d_2257_i4_)] = (d_2257_i4_) - ((len((D14_DC39(_dafny.SeqWithoutIsStrInference([824]), (D1_DC3(_dafny.Map({d_2170_v0_: d_2170_v0_}))).cf3, False)).cf68)) - (887))
        d_2258_v70_: _dafny.Array
        nw424_ = _dafny.Array(_dafny.Seq({}), 17)
        d_2258_v70_ = nw424_
        d_2258_v70_ = d_2258_v70_
        r0 = 625
        d_2259_v71_: _dafny.Set
        d_2259_v71_ = _dafny.Set({d_2170_v0_, d_2170_v0_, not(d_2170_v0_)})
        d_2259_v71_ = d_2259_v71_
        d_2260_v72_: int
        d_2260_v72_ = 374
        d_2261_v73_: D7
        d_2261_v73_ = D7_DC16(d_2170_v0_, d_2260_v72_, d_2170_v0_)
        d_2262_v74_: _dafny.Map
        d_2262_v74_ = _dafny.Map({d_2261_v73_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ty"))})
        d_2263_v75_: D13
        d_2263_v75_ = D13_DC36(d_2262_v74_)
        pat_let_tv58_ = d_2260_v72_
        pat_let_tv59_ = d_2260_v72_
        pat_let_tv60_ = d_2260_v72_
        def lambda112_(source21_):
            if source21_.is_DC37:
                d_2264___mcc_h0_ = source21_.cf65
                d_2265___mcc_h1_ = source21_.cf66
                d_2266_cf66_ = d_2265___mcc_h1_
                d_2267_cf65_ = d_2264___mcc_h0_
                return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vuc")))) + (pat_let_tv58_)
            elif True:
                d_2268___mcc_h2_ = source21_.cf64
                d_2269_cf64_ = d_2268___mcc_h2_
                return default__.safeModuloInt(pat_let_tv59_, pat_let_tv60_)

        r0 = lambda112_(d_2263_v75_)
        d_2270_v76_: T0
        nw425_ = C4()
        nw425_.ctor__()
        d_2270_v76_ = nw425_
        d_2271_v77_: _dafny.Map
        d_2271_v77_ = _dafny.Map({d_2260_v72_: d_2270_v76_})
        r1 = not((len(d_2271_v77_)) < (-718))
        return r0, r1

    def m2(self, p0, globalState):
        if p0:
            d_2272_v0_: _dafny.Array
            nw426_ = _dafny.Array(_dafny.Set({}), 7)
            d_2272_v0_ = nw426_
            d_2273_v1_: int
            d_2273_v1_ = 693
            d_2274_v2_: _dafny.Set
            d_2274_v2_ = _dafny.Set({d_2273_v1_})
            index394_ = default__.safeIndex(753, (d_2272_v0_).length(0))
            (d_2272_v0_)[index394_] = (d_2274_v2_) | (d_2274_v2_)
            index395_ = default__.safeIndex(753, (d_2272_v0_).length(0))
            (d_2272_v0_)[index395_] = (d_2274_v2_).intersection(d_2274_v2_)
            d_2275_v3_: _dafny.Array
            def lambda113_(d_2276_p0_):
                def lambda114_(d_2277_i0_):
                    return d_2276_p0_

                return lambda114_

            init65_ = lambda113_(p0)
            nw427_ = _dafny.Array(None, 8)
            for i0_65_ in range(nw427_.length(0)):
                nw427_[i0_65_] = init65_(i0_65_)
            d_2275_v3_ = nw427_
            d_2278_v4_: _dafny.Array
            nw428_ = _dafny.Array(int(0), 7)
            d_2278_v4_ = nw428_
            d_2279_v5_: _dafny.Map
            d_2279_v5_ = _dafny.Map({d_2275_v3_: d_2278_v4_})
            d_2279_v5_ = (d_2279_v5_).set((d_2275_v3_), d_2278_v4_)
            d_2273_v1_ = d_2273_v1_
            (globalState).f1 = p0
            d_2280_v6_: _dafny.Map
            d_2280_v6_ = _dafny.Map({p0: d_2273_v1_})
            (globalState).f1 = (p0) in (d_2280_v6_)
        elif True:
            if p0:
                d_2281_v7_: str
                d_2281_v7_ = _dafny.CodePoint('q')
                d_2281_v7_ = d_2281_v7_
                d_2282_v8_: _dafny.Array
                nw429_ = _dafny.Array(int(0), 11)
                d_2282_v8_ = nw429_
                d_2283_v9_: D3
                d_2283_v9_ = D3_DC6(d_2282_v8_)
                d_2284_v10_: _dafny.Seq
                d_2284_v10_ = _dafny.SeqWithoutIsStrInference([D3_DC6(d_2282_v8_), d_2283_v9_, d_2283_v9_])
                d_2285_v11_: C0
                nw430_ = C0()
                nw430_.ctor__((_dafny.SeqWithoutIsStrInference([d_2283_v9_])) + (d_2284_v10_), p0)
                d_2285_v11_ = nw430_
                d_2286_v12_: _dafny.Array
                def lambda115_(d_2287_p0_):
                    def lambda116_(d_2288_i1_):
                        return d_2287_p0_

                    return lambda116_

                init66_ = lambda115_(p0)
                nw431_ = _dafny.Array(None, 2)
                for i0_66_ in range(nw431_.length(0)):
                    nw431_[i0_66_] = init66_(i0_66_)
                d_2286_v12_ = nw431_
                index396_ = default__.safeIndex(592, (d_2286_v12_).length(0))
                (d_2286_v12_)[index396_] = (p0) == (p0)
                index397_ = default__.safeIndex(592, (d_2286_v12_).length(0))
                (d_2286_v12_)[index397_] = (d_2285_v11_).f23
                d_2289_v13_: int
                d_2289_v13_ = 946
                d_2290_v14_: _dafny.Seq
                d_2290_v14_ = _dafny.SeqWithoutIsStrInference([False])
                index398_ = default__.safeIndex(592, (d_2286_v12_).length(0))
                (d_2286_v12_)[index398_] = (default__.fm13(p0, _dafny.CodePoint('q'), d_2289_v13_, globalState)) <= (d_2290_v14_)
                (globalState).f1 = (95) < ((0) - (d_2289_v13_))
            elif True:
                d_2291_v15_: _dafny.Array
                nw432_ = _dafny.Array(False, 5)
                d_2291_v15_ = nw432_
                (globalState).f14 = (d_2291_v15_ if p0 else d_2291_v15_)
                d_2292_v16_: _dafny.Seq
                d_2292_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "asuvdd"))
                d_2292_v16_ = d_2292_v16_
                d_2293_v17_: str
                d_2293_v17_ = _dafny.CodePoint('s')
                d_2294_v18_: _dafny.Map
                d_2294_v18_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_2293_v17_ for d_2295_i2_ in range(default__.abs(356))])): p0})
                d_2296_v19_: int
                d_2296_v19_ = 335
                d_2293_v17_ = (d_2292_v16_)[default__.safeIndex((len(d_2294_v18_) if p0 else d_2296_v19_), len(d_2292_v16_))]
                d_2297_v20_: _dafny.Array
                nw433_ = _dafny.Array(_dafny.Map({}), 3)
                d_2297_v20_ = nw433_
                d_2298_v21_: _dafny.Map
                d_2298_v21_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmetnnr")): (0) - (d_2296_v19_)})
                index399_ = default__.safeIndex(317, (d_2297_v20_).length(0))
                (d_2297_v20_)[index399_] = d_2298_v21_
                index400_ = default__.safeIndex(317, (d_2297_v20_).length(0))
                (d_2297_v20_)[index400_] = d_2298_v21_
                d_2299_v22_: _dafny.Seq
                d_2299_v22_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, (D14_DC40(p0)).cf71, p0])
                d_2300_v23_: _dafny.MultiSet
                d_2300_v23_ = _dafny.MultiSet([p0])
                d_2301_v24_: C6
                nw434_ = C6()
                nw434_.ctor__()
                d_2301_v24_ = nw434_
                d_2302_v25_: _dafny.Map
                d_2302_v25_ = _dafny.Map({p0: d_2301_v24_})
                d_2303_v26_: _dafny.Array
                nw435_ = _dafny.Array(None, 15)
                nw435_[int(0)] = default__.fm29(p0, d_2296_v19_, d_2296_v19_, (d_2299_v22_)[default__.safeIndex(len((default__.fm12(d_2300_v23_, d_2296_v19_, globalState)).set(default__.safeIndex(d_2296_v19_, len(default__.fm12(d_2300_v23_, d_2296_v19_, globalState))), d_2296_v19_)), len(d_2299_v22_))], globalState)
                nw435_[int(1)] = _dafny.CodePoint('q')
                nw435_[int(2)] = d_2293_v17_
                nw435_[int(3)] = d_2293_v17_
                nw435_[int(4)] = d_2293_v17_
                nw435_[int(5)] = d_2293_v17_
                nw435_[int(6)] = d_2293_v17_
                nw435_[int(7)] = _dafny.CodePoint('v')
                nw435_[int(8)] = d_2293_v17_
                nw435_[int(9)] = default__.fm29(p0, len(d_2302_v25_), d_2296_v19_, p0, globalState)
                nw435_[int(10)] = d_2293_v17_
                nw435_[int(11)] = d_2293_v17_
                nw435_[int(12)] = d_2293_v17_
                nw435_[int(13)] = d_2293_v17_
                nw435_[int(14)] = d_2293_v17_
                d_2303_v26_ = nw435_
                rhs320_ = d_2303_v26_
                rhs321_ = d_2292_v16_
                d_2303_v26_ = rhs320_
                d_2292_v16_ = rhs321_
            d_2304_v27_: _dafny.Array
            nw436_ = _dafny.Array(_dafny.Map({}), 5)
            d_2304_v27_ = nw436_
            source22_ = D9_DC22(d_2304_v27_)
            if source22_.is_DC23:
                d_2305___mcc_h0_ = source22_.cf37
                d_2306___mcc_h1_ = source22_.cf38
                d_2307___mcc_h2_ = source22_.cf39
                d_2308___mcc_h3_ = source22_.cf40
                d_2309___mcc_h4_ = source22_.cf41
                d_2310_cf41_ = d_2309___mcc_h4_
                d_2311_cf40_ = d_2308___mcc_h3_
                d_2312_cf39_ = d_2307___mcc_h2_
                d_2313_cf38_ = d_2306___mcc_h1_
                d_2314_cf37_ = d_2305___mcc_h0_
                d_2315_v28_: _dafny.Array
                nw437_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_2315_v28_ = nw437_
                d_2316_v29_: _dafny.Seq
                d_2316_v29_ = _dafny.SeqWithoutIsStrInference([d_2315_v28_, d_2315_v28_, d_2315_v28_])
                d_2315_v28_ = (d_2316_v29_)[default__.safeIndex(943, len(d_2316_v29_))]
                (globalState).f1 = p0
                d_2317_v30_: _dafny.MultiSet
                d_2317_v30_ = _dafny.MultiSet([d_2313_cf38_])
                index401_ = default__.safeIndex(916, (d_2311_cf40_).length(0))
                (d_2311_cf40_)[index401_] = ((d_2317_v30_).set(d_2313_cf38_, default__.abs(default__.fm1(d_2312_cf39_, globalState)))).cardinality
                d_2318_v31_: int
                d_2318_v31_ = 133
                d_2319_v32_: _dafny.Seq
                d_2319_v32_ = _dafny.SeqWithoutIsStrInference([d_2318_v31_, d_2318_v31_])
                index402_ = default__.safeIndex(916, (d_2311_cf40_).length(0))
                (d_2311_cf40_)[index402_] = (0) - (default__.safeDivisionInt((d_2319_v32_)[default__.safeIndex(d_2318_v31_, len(d_2319_v32_))], d_2318_v31_))
                d_2320_v33_: _dafny.MultiSet
                d_2320_v33_ = _dafny.MultiSet([p0])
                d_2321_v34_: _dafny.Map
                d_2321_v34_ = _dafny.Map({d_2320_v33_: False})
                d_2321_v34_ = (d_2321_v34_).set(d_2320_v33_, p0)
            elif source22_.is_DC24:
                d_2322___mcc_h5_ = source22_.cf42
                d_2323___mcc_h6_ = source22_.cf43
                d_2324_cf43_ = d_2323___mcc_h6_
                d_2325_cf42_ = d_2322___mcc_h5_
                d_2326_v36_: _dafny.Array
                def lambda117_(d_2327_i3_):
                    def iife253_():
                        coll79_ = _dafny.Map()
                        compr_79_: int
                        for compr_79_ in _dafny.IntegerRange(819, -438):
                            d_2328_v35_: int = compr_79_
                            if ((819) <= (d_2328_v35_)) and ((d_2328_v35_) < (-438)):
                                coll79_[(d_2328_v35_) - (-571)] = 890
                        return _dafny.Map(coll79_)
                    return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))): -215})) | (_dafny.Map({50: len(iife253_()
                    )}))

                init67_ = lambda117_
                nw438_ = _dafny.Array(None, 26)
                for i0_67_ in range(nw438_.length(0)):
                    nw438_[i0_67_] = init67_(i0_67_)
                d_2326_v36_ = nw438_
                d_2329_v37_: _dafny.MultiSet
                d_2329_v37_ = _dafny.MultiSet([default__.fm1(not(p0), globalState)])
                d_2330_v38_: int
                d_2330_v38_ = -416
                d_2331_v39_: _dafny.Map
                d_2331_v39_ = _dafny.Map({(((d_2329_v37_)[d_2330_v38_] if (d_2330_v38_) in (d_2329_v37_) else d_2330_v38_)) > ((_dafny.MultiSet([d_2330_v38_])).cardinality): d_2330_v38_})
                rhs322_ = len(d_2331_v39_)
                rhs323_ = p0
                rhs324_ = d_2326_v36_
                rhs325_ = d_2330_v38_
                lhs283_ = globalState
                lhs284_ = globalState
                lhs285_ = globalState
                lhs283_.f6 = rhs322_
                lhs284_.f1 = rhs323_
                d_2326_v36_ = rhs324_
                lhs285_.f6 = rhs325_
                d_2332_v40_: _dafny.Seq
                d_2332_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ced"))
                d_2333_v41_: D8
                d_2333_v41_ = D8_DC21(d_2332_v40_, _dafny.Set({p0}))
                d_2334_v42_: _dafny.Map
                d_2334_v42_ = _dafny.Map({((d_2333_v41_).cf34) + (d_2332_v40_): (0) - (d_2330_v38_)})
                d_2335_v43_: _dafny.Set
                d_2335_v43_ = _dafny.Set({d_2332_v40_})
                d_2334_v42_ = (d_2334_v42_).set(d_2332_v40_, len(d_2335_v43_))
                d_2336_v44_: _dafny.Array
                def lambda118_(d_2337_p0_):
                    def lambda119_(d_2338_i4_):
                        return not(not(d_2337_p0_))

                    return lambda119_

                init68_ = lambda118_(p0)
                nw439_ = _dafny.Array(None, 22)
                for i0_68_ in range(nw439_.length(0)):
                    nw439_[i0_68_] = init68_(i0_68_)
                d_2336_v44_ = nw439_
                index403_ = default__.safeIndex(13, (d_2336_v44_).length(0))
                (d_2336_v44_)[index403_] = p0
                index404_ = default__.safeIndex(13, (d_2336_v44_).length(0))
                (d_2336_v44_)[index404_] = (d_2330_v38_) < (d_2330_v38_)
                index405_ = default__.safeIndex(13, (d_2336_v44_).length(0))
                (d_2336_v44_)[index405_] = ((d_2331_v39_) | (d_2331_v39_)) != (d_2331_v39_)
            elif source22_.is_DC22:
                d_2339___mcc_h7_ = source22_.cf36
                d_2340_cf36_ = d_2339___mcc_h7_
                d_2341_v45_: _dafny.Array
                nw440_ = _dafny.Array(False, 3)
                d_2341_v45_ = nw440_
                index406_ = default__.safeIndex(292, (d_2341_v45_).length(0))
                (d_2341_v45_)[index406_] = default__.fm2(globalState)
                d_2342_v46_: int
                d_2342_v46_ = 865
                index407_ = default__.safeIndex(292, (d_2341_v45_).length(0))
                (d_2341_v45_)[index407_] = (default__.fm1(True, globalState)) < (d_2342_v46_)
                d_2343_v47_: C4
                nw441_ = C4()
                nw441_.ctor__()
                d_2343_v47_ = nw441_
                (globalState).f6 = (0) - ((D23_DC55(d_2342_v46_)).cf94)
                d_2344_v48_: _dafny.Map
                d_2344_v48_ = _dafny.Map({d_2341_v45_: d_2341_v45_})
                d_2345_v49_: _dafny.Seq
                d_2345_v49_ = _dafny.SeqWithoutIsStrInference([((d_2344_v48_)[d_2341_v45_] if (d_2341_v45_) in (d_2344_v48_) else d_2341_v45_)])
                d_2341_v45_ = ((d_2341_v45_ if p0 else d_2341_v45_) if p0 else (d_2345_v49_)[default__.safeIndex(d_2342_v46_, len(d_2345_v49_))])
            elif True:
                d_2346___mcc_h8_ = source22_.cf44
                d_2347_cf44_ = d_2346___mcc_h8_
                d_2348_v50_: str
                d_2348_v50_ = _dafny.CodePoint('h')
                d_2348_v50_ = d_2348_v50_
                d_2349_v51_: C6
                nw442_ = C6()
                nw442_.ctor__()
                d_2349_v51_ = nw442_
                d_2350_v52_: _dafny.Array
                nw443_ = _dafny.Array(None, 22)
                d_2350_v52_ = nw443_
                d_2351_v53_: _dafny.Seq
                d_2351_v53_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2352_v54_: int
                d_2352_v54_ = 981
                d_2353_v55_: D11
                d_2353_v55_ = D11_DC31(d_2351_v53_, d_2352_v54_, len(d_2351_v53_), d_2352_v54_, 982)
                d_2354_v56_: _dafny.Map
                d_2354_v56_ = _dafny.Map({_dafny.Map({(d_2353_v55_).cf56: p0}): d_2352_v54_})
                d_2355_v57_: T1
                nw444_ = C3()
                nw444_.ctor__(p0, 345, d_2354_v56_)
                d_2355_v57_ = nw444_
                index408_ = default__.safeIndex(995, (d_2350_v52_).length(0))
                (d_2350_v52_)[index408_] = d_2355_v57_
                index409_ = default__.safeIndex(995, (d_2350_v52_).length(0))
                (d_2350_v52_)[index409_] = d_2355_v57_
                d_2356_v58_: _dafny.Map
                d_2356_v58_ = _dafny.Map({_dafny.MultiSet([p0]): d_2352_v54_})
                (globalState).f6 = default__.safeModuloInt(default__.safeDivisionInt(d_2352_v54_, len(d_2356_v58_)), d_2352_v54_)
            d_2357_v59_: int
            d_2357_v59_ = -464
            (globalState).f6 = d_2357_v59_
            d_2357_v59_ = (0) - ((0) - (d_2357_v59_))
            d_2358_v60_: _dafny.Set
            d_2358_v60_ = _dafny.Set({d_2357_v59_})
            d_2359_v61_: _dafny.MultiSet
            d_2359_v61_ = _dafny.MultiSet([d_2358_v60_, (d_2358_v60_).intersection(d_2358_v60_)])
            d_2359_v61_ = d_2359_v61_
        d_2360_v62_: int
        d_2360_v62_ = 994
        d_2361_v63_: _dafny.Seq
        d_2361_v63_ = _dafny.SeqWithoutIsStrInference([d_2360_v62_, d_2360_v62_])
        hi11_ = (d_2360_v62_) - (-67)
        for d_2362_i5_ in range(len((d_2361_v63_) + (d_2361_v63_)), hi11_):
            if not(p0):
                d_2363_v64_: str
                d_2363_v64_ = _dafny.CodePoint('n')
                d_2364_v65_: _dafny.MultiSet
                d_2364_v65_ = _dafny.MultiSet([d_2360_v62_])
                (globalState).f1 = ((d_2364_v65_).ispropersubset(default__.fm42(d_2360_v62_, d_2363_v64_, d_2360_v62_, p0, globalState)) if p0 else p0)
                d_2365_v66_: _dafny.Array
                def lambda120_(d_2366_i6_):
                    return D0_DC2(D0_DC2(D0_DC2(D0_DC2(D0_DC2(D0_DC0(True))))))

                init69_ = lambda120_
                nw445_ = _dafny.Array(None, 22)
                for i0_69_ in range(nw445_.length(0)):
                    nw445_[i0_69_] = init69_(i0_69_)
                d_2365_v66_ = nw445_
                d_2367_v67_: _dafny.Array
                def lambda121_(d_2368_i5_):
                    def lambda122_(d_2369_i7_):
                        return (d_2369_i7_) + (d_2368_i5_)

                    return lambda122_

                init70_ = lambda121_(d_2362_i5_)
                nw446_ = _dafny.Array(None, 2)
                for i0_70_ in range(nw446_.length(0)):
                    nw446_[i0_70_] = init70_(i0_70_)
                d_2367_v67_ = nw446_
                default__.m0(_dafny.Set({d_2362_i5_, d_2362_i5_}), d_2365_v66_, d_2367_v67_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unfypjx")), globalState)
                d_2370_v68_: _dafny.Set
                d_2370_v68_ = _dafny.Set({(0) - (d_2360_v62_), d_2362_i5_})
                d_2371_v69_: _dafny.Seq
                d_2371_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxjyeoe"))
                default__.m0(d_2370_v68_, d_2365_v66_, d_2367_v67_, d_2371_v69_, globalState)
                (globalState).f1 = p0
                def iife254_():
                    coll80_ = _dafny.Map()
                    compr_80_: int
                    for compr_80_ in _dafny.IntegerRange(490, 853):
                        d_2372_v70_: int = compr_80_
                        if ((490) <= (d_2372_v70_)) and ((d_2372_v70_) < (853)):
                            coll80_[default__.safeDivisionInt(d_2372_v70_, d_2360_v62_)] = p0
                    return _dafny.Map(coll80_)
                (globalState).f6 = (542) + ((0) - (len(iife254_()
                )))
            elif True:
                d_2373_v71_: str
                d_2373_v71_ = _dafny.CodePoint('j')
                d_2373_v71_ = d_2373_v71_
                d_2374_v72_: _dafny.Set
                d_2374_v72_ = _dafny.Set({(0) - (d_2362_i5_)})
                d_2374_v72_ = d_2374_v72_
                (globalState).f1 = p0
                d_2375_v73_: _dafny.MultiSet
                d_2375_v73_ = _dafny.MultiSet([d_2373_v71_, d_2373_v71_, d_2373_v71_])
                d_2361_v63_ = ((_dafny.SeqWithoutIsStrInference([len(d_2361_v63_) for d_2376_i8_ in range(default__.abs(-148))])) + (d_2361_v63_)).set(default__.safeIndex((d_2375_v73_).cardinality, len((_dafny.SeqWithoutIsStrInference([len(d_2361_v63_) for d_2377_i8_ in range(default__.abs(-148))])) + (d_2361_v63_))), (d_2360_v62_) + (d_2362_i5_))
                d_2360_v62_ = d_2362_i5_
            d_2378_v74_: _dafny.Map
            d_2378_v74_ = _dafny.Map({d_2362_i5_: p0})
            d_2379_v75_: _dafny.Map
            d_2379_v75_ = _dafny.Map({d_2378_v74_: len(_dafny.SeqWithoutIsStrInference([not(True), p0]))})
            d_2380_v76_: C3
            nw447_ = C3()
            nw447_.ctor__(False, -476, d_2379_v75_)
            d_2380_v76_ = nw447_
            d_2381_v77_: _dafny.Seq
            d_2381_v77_ = _dafny.SeqWithoutIsStrInference([d_2380_v76_])
            d_2382_v78_: _dafny.Seq
            d_2382_v78_ = _dafny.SeqWithoutIsStrInference([d_2381_v77_])
            d_2383_v79_: _dafny.Map
            d_2383_v79_ = _dafny.Map({p0: (d_2362_i5_) > (len(d_2382_v78_))})
            d_2383_v79_ = (d_2383_v79_).set(True, False)
            d_2384_v80_: str
            d_2384_v80_ = _dafny.CodePoint('k')
            d_2385_v81_: _dafny.Seq
            d_2385_v81_ = _dafny.SeqWithoutIsStrInference([d_2384_v80_, d_2384_v80_, d_2384_v80_])
            d_2386_v82_: _dafny.Seq
            d_2386_v82_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, (D5_DC12((d_2380_v76_).f19, p0, d_2384_v80_)).cf16, (d_2380_v76_).f18])
            d_2387_v83_: _dafny.Seq
            d_2387_v83_ = _dafny.SeqWithoutIsStrInference([d_2385_v81_, d_2385_v81_])
            d_2388_v84_: _dafny.Array
            nw448_ = _dafny.Array(None, 29)
            nw448_[int(0)] = default__.safeDivisionInt(d_2360_v62_, d_2362_i5_)
            nw448_[int(1)] = d_2360_v62_
            nw448_[int(2)] = (len(default__.fm31(d_2360_v62_, True, globalState)) if p0 else len(_dafny.Map({(d_2380_v76_).f18: d_2360_v62_})))
            nw448_[int(3)] = (0) - (d_2362_i5_)
            nw448_[int(4)] = d_2360_v62_
            nw448_[int(5)] = (0) - (d_2362_i5_)
            nw448_[int(6)] = -829
            nw448_[int(7)] = d_2360_v62_
            nw448_[int(8)] = d_2360_v62_
            nw448_[int(9)] = d_2362_i5_
            nw448_[int(10)] = default__.safeDivisionInt(d_2362_i5_, d_2360_v62_)
            nw448_[int(11)] = (d_2362_i5_) * (d_2360_v62_)
            nw448_[int(12)] = d_2362_i5_
            nw448_[int(13)] = 461
            nw448_[int(14)] = (_dafny.MultiSet(d_2385_v81_)).cardinality
            nw448_[int(15)] = (d_2380_v76_).f19
            nw448_[int(16)] = len(default__.fm28(d_2360_v62_, (d_2380_v76_).f18, True, globalState))
            nw448_[int(17)] = -899
            nw448_[int(18)] = default__.safeModuloInt(d_2362_i5_, d_2362_i5_)
            nw448_[int(19)] = len(d_2386_v82_)
            nw448_[int(20)] = default__.safeDivisionInt((d_2380_v76_).f19, d_2362_i5_)
            nw448_[int(21)] = d_2362_i5_
            nw448_[int(22)] = (d_2380_v76_).f19
            nw448_[int(23)] = d_2362_i5_
            nw448_[int(24)] = len((d_2387_v83_)[default__.safeIndex(d_2362_i5_, len(d_2387_v83_))])
            nw448_[int(25)] = d_2362_i5_
            nw448_[int(26)] = default__.safeDivisionInt((0) - (d_2360_v62_), 673)
            nw448_[int(27)] = d_2360_v62_
            nw448_[int(28)] = d_2362_i5_
            d_2388_v84_ = nw448_
            (globalState).f9 = d_2388_v84_
            (globalState).f1 = p0
        d_2389_v85_: _dafny.Array
        nw449_ = _dafny.Array(None, 17)
        nw449_[int(0)] = True
        nw449_[int(1)] = (self).fm8(d_2360_v62_, d_2360_v62_, (0) - (d_2360_v62_), globalState)
        nw449_[int(2)] = p0
        nw449_[int(3)] = p0
        nw449_[int(4)] = p0
        nw449_[int(5)] = p0
        nw449_[int(6)] = p0
        nw449_[int(7)] = p0
        nw449_[int(8)] = p0
        nw449_[int(9)] = True
        nw449_[int(10)] = p0
        nw449_[int(11)] = default__.fm2(globalState)
        nw449_[int(12)] = p0
        nw449_[int(13)] = False
        nw449_[int(14)] = p0
        nw449_[int(15)] = p0
        nw449_[int(16)] = True
        d_2389_v85_ = nw449_
        d_2390_v86_: _dafny.MultiSet
        d_2390_v86_ = _dafny.MultiSet([d_2389_v85_])
        hi12_ = (d_2390_v86_).cardinality
        for d_2391_i9_ in range(d_2360_v62_, hi12_):
            d_2392_v87_: _dafny.Array
            def lambda123_(d_2393_v62_):
                def lambda124_(d_2394_i10_):
                    return (d_2394_i10_) - (d_2393_v62_)

                return lambda124_

            init71_ = lambda123_(d_2360_v62_)
            nw450_ = _dafny.Array(None, 25)
            for i0_71_ in range(nw450_.length(0)):
                nw450_[i0_71_] = init71_(i0_71_)
            d_2392_v87_ = nw450_
            index410_ = default__.safeIndex(357, (d_2392_v87_).length(0))
            (d_2392_v87_)[index410_] = (d_2360_v62_ if p0 else d_2391_i9_)
            index411_ = default__.safeIndex(357, (d_2392_v87_).length(0))
            (d_2392_v87_)[index411_] = d_2391_i9_
            d_2395_v88_: _dafny.Seq
            d_2395_v88_ = _dafny.SeqWithoutIsStrInference([default__.fm2(globalState)])
            d_2360_v62_ = default__.safeModuloInt(len(d_2395_v88_), d_2391_i9_)
            d_2396_v90_: str
            d_2396_v90_ = _dafny.CodePoint('o')
            d_2397_v91_: _dafny.Set
            d_2397_v91_ = _dafny.Set({default__.fm2(globalState), p0, p0})
            d_2398_v92_: _dafny.Map
            def iife255_():
                coll81_ = _dafny.Map()
                compr_81_: str
                for compr_81_ in (_dafny.Map({d_2396_v90_: p0})).keys.Elements:
                    d_2399_v89_: str = compr_81_
                    if (d_2399_v89_) in (_dafny.Map({d_2396_v90_: p0})):
                        coll81_[d_2399_v89_] = len(_dafny.Set({p0}))
                return _dafny.Map(coll81_)
            d_2398_v92_ = _dafny.Map({iife255_()
            : (d_2397_v91_).ispropersubset(d_2397_v91_)})
            d_2400_v93_: _dafny.Map
            d_2400_v93_ = _dafny.Map({d_2396_v90_: (d_2392_v87_)[default__.safeIndex(357, (d_2392_v87_).length(0))]})
            d_2398_v92_ = (d_2398_v92_).set(d_2400_v93_, (p0 if p0 else p0))
            d_2401_v94_: C5
            nw451_ = C5()
            nw451_.ctor__(d_2360_v62_)
            d_2401_v94_ = nw451_
            d_2402_v95_: _dafny.Map
            d_2402_v95_ = _dafny.Map({d_2401_v94_: d_2391_i9_})
            d_2403_v96_: _dafny.Set
            d_2403_v96_ = _dafny.Set({d_2361_v63_, d_2361_v63_})
            d_2404_v97_: _dafny.Map
            d_2404_v97_ = _dafny.Map({p0: (((d_2402_v95_)[d_2401_v94_] if (d_2401_v94_) in (d_2402_v95_) else len(d_2403_v96_))) > ((d_2401_v94_).f16)})
            d_2404_v97_ = (d_2404_v97_).set(not(p0), p0)
        d_2405_v98_: T0
        nw452_ = C1()
        nw452_.ctor__((0) - (d_2360_v62_))
        d_2405_v98_ = nw452_
        d_2406_v99_: _dafny.Map
        d_2406_v99_ = _dafny.Map({d_2405_v98_: p0})
        d_2406_v99_ = (d_2406_v99_) | (d_2406_v99_)
        d_2407_v100_: _dafny.Array
        nw453_ = _dafny.Array(int(0), 25)
        d_2407_v100_ = nw453_
        index412_ = default__.safeIndex(802, (d_2407_v100_).length(0))
        (d_2407_v100_)[index412_] = d_2360_v62_
        d_2408_v101_: C4
        nw454_ = C4()
        nw454_.ctor__()
        d_2408_v101_ = nw454_
        d_2409_v102_: _dafny.Map
        d_2409_v102_ = _dafny.Map({len(d_2361_v63_): d_2408_v101_})
        index413_ = default__.safeIndex(802, (d_2407_v100_).length(0))
        (d_2407_v100_)[index413_] = len(_dafny.SeqWithoutIsStrInference([((d_2409_v102_)[d_2360_v62_] if (d_2360_v62_) in (d_2409_v102_) else d_2408_v101_)]))
        d_2410_v103_: _dafny.Seq
        d_2410_v103_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
        d_2411_v104_: str
        d_2411_v104_ = _dafny.CodePoint('v')
        d_2412_v105_: _dafny.MultiSet
        d_2412_v105_ = _dafny.MultiSet([True, p0, p0, p0, True])
        d_2413_v106_: _dafny.Map
        d_2413_v106_ = _dafny.Map({d_2412_v105_: d_2389_v85_})
        rhs326_ = (d_2410_v103_ if not((p0) and (p0)) else default__.fm13(default__.fm2(globalState), d_2411_v104_, d_2360_v62_, globalState))
        rhs327_ = default__.safeDivisionInt((d_2361_v63_)[default__.safeIndex(len(d_2413_v106_), len(d_2361_v63_))], (d_2407_v100_)[default__.safeIndex(802, (d_2407_v100_).length(0))])
        d_2410_v103_ = rhs326_
        d_2360_v62_ = rhs327_

    def m7(self, globalState):
        d_2414_v0_: bool
        d_2414_v0_ = True
        d_2415_i0_: int
        d_2415_i0_ = 0
        with _dafny.label("15"):
            while (d_2414_v0_ if d_2414_v0_ else d_2414_v0_):
                with _dafny.c_label("15"):
                    if (d_2415_i0_) >= (100):
                        raise _dafny.Break("15")
                    d_2415_i0_ = (d_2415_i0_) + (1)
                    d_2416_v1_: int
                    d_2416_v1_ = -499
                    (globalState).f1 = (self).fm8(d_2416_v1_, d_2416_v1_, 710, globalState)
                    d_2417_v2_: _dafny.Map
                    d_2417_v2_ = _dafny.Map({d_2416_v1_: d_2416_v1_})
                    d_2418_v3_: _dafny.Map
                    d_2418_v3_ = _dafny.Map({d_2414_v0_: d_2414_v0_})
                    d_2419_v4_: _dafny.Map
                    d_2419_v4_ = _dafny.Map({len(d_2418_v3_): d_2414_v0_})
                    d_2417_v2_ = (d_2417_v2_).set(d_2416_v1_, len(d_2419_v4_))
                    d_2420_v5_: _dafny.Set
                    d_2420_v5_ = _dafny.Set({True})
                    d_2421_v6_: _dafny.MultiSet
                    d_2421_v6_ = _dafny.MultiSet([(d_2420_v5_).issubset(_dafny.Set({d_2414_v0_})), d_2414_v0_])
                    d_2421_v6_ = d_2421_v6_
                    (globalState).f1 = not(((d_2418_v3_)[d_2414_v0_] if (d_2414_v0_) in (d_2418_v3_) else d_2414_v0_))
                    pass
            pass
        d_2422_v7_: _dafny.Set
        d_2422_v7_ = _dafny.Set({554})
        d_2423_v8_: int
        d_2423_v8_ = 867
        d_2424_v9_: _dafny.Seq
        d_2424_v9_ = _dafny.SeqWithoutIsStrInference([d_2423_v8_, d_2423_v8_])
        d_2425_v11_: _dafny.Map
        d_2425_v11_ = _dafny.Map({d_2423_v8_: d_2414_v0_})
        def iife256_():
            coll82_ = _dafny.Set()
            compr_82_: int
            for compr_82_ in (d_2424_v9_).Elements:
                d_2426_v10_: int = compr_82_
                if (d_2426_v10_) in (d_2424_v9_):
                    coll82_ = coll82_.union(_dafny.Set([(d_2426_v10_) * (len(_dafny.SeqWithoutIsStrInference([104])))]))
            return _dafny.Set(coll82_)
        (globalState).f6 = (-388 if (d_2422_v7_).isdisjoint(iife256_()
        ) else len(d_2425_v11_))
        rhs328_ = default__.safeDivisionInt(13, (0) - (d_2423_v8_))
        rhs329_ = (d_2414_v0_) and (d_2414_v0_)
        d_2423_v8_ = rhs328_
        d_2414_v0_ = rhs329_
        (globalState).f6 = -335
        d_2427_v12_: _dafny.MultiSet
        d_2427_v12_ = _dafny.MultiSet([default__.fm2(globalState)])
        d_2427_v12_ = d_2427_v12_
        if False:
            d_2428_v13_: _dafny.MultiSet
            d_2428_v13_ = _dafny.MultiSet([678])
            d_2429_v14_: _dafny.Seq
            d_2429_v14_ = _dafny.SeqWithoutIsStrInference([d_2414_v0_])
            (globalState).f1 = (len(d_2422_v7_)) == (((d_2428_v13_) | (_dafny.MultiSet([d_2423_v8_, d_2423_v8_, d_2423_v8_, len(d_2429_v14_), d_2423_v8_]))).cardinality)
            d_2430_v15_: str
            d_2430_v15_ = _dafny.CodePoint('d')
            d_2431_v16_: _dafny.MultiSet
            d_2431_v16_ = _dafny.MultiSet([d_2430_v15_])
            d_2432_v17_: _dafny.Map
            d_2432_v17_ = _dafny.Map({not(d_2414_v0_): (d_2431_v16_) - (d_2431_v16_)})
            d_2432_v17_ = d_2432_v17_
            (globalState).f6 = d_2423_v8_
            d_2433_v18_: _dafny.Seq
            d_2433_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dla"))
            d_2434_v19_: _dafny.Map
            d_2434_v19_ = _dafny.Map({(((d_2427_v12_)[d_2414_v0_] if (d_2414_v0_) in (d_2427_v12_) else d_2423_v8_)) * (len(d_2433_v18_)): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ug"))})
            d_2435_v20_: _dafny.Seq
            d_2435_v20_ = _dafny.SeqWithoutIsStrInference([d_2433_v18_])
            d_2433_v18_ = ((d_2434_v19_)[(d_2423_v8_) - (d_2423_v8_)] if ((d_2423_v8_) - (d_2423_v8_)) in (d_2434_v19_) else (d_2435_v20_)[default__.safeIndex(d_2423_v8_, len(d_2435_v20_))])
            d_2436_v21_: C5
            nw455_ = C5()
            nw455_.ctor__(d_2423_v8_)
            d_2436_v21_ = nw455_
        elif True:
            rhs330_ = (d_2422_v7_) | (d_2422_v7_)
            rhs331_ = default__.fm1(d_2414_v0_, globalState)
            lhs286_ = globalState
            d_2422_v7_ = rhs330_
            lhs286_.f6 = rhs331_
            d_2437_v22_: _dafny.Seq
            d_2437_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mserh"))
            d_2437_v22_ = ((d_2437_v22_) + (d_2437_v22_)) + (d_2437_v22_)
            d_2438_v23_: str
            d_2438_v23_ = _dafny.CodePoint('k')
            d_2439_v24_: _dafny.MultiSet
            d_2439_v24_ = _dafny.MultiSet([d_2423_v8_, len(_dafny.SeqWithoutIsStrInference([d_2414_v0_]))])
            d_2440_v25_: _dafny.Map
            d_2440_v25_ = _dafny.Map({d_2423_v8_: d_2438_v23_})
            d_2441_v26_: _dafny.Map
            d_2441_v26_ = _dafny.Map({_dafny.Map({d_2438_v23_: d_2423_v8_}): len(d_2440_v25_)})
            d_2442_v27_: _dafny.Map
            d_2442_v27_ = _dafny.Map({d_2414_v0_: d_2438_v23_})
            d_2443_v28_: _dafny.Map
            d_2443_v28_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_2439_v24_).cardinality for d_2444_i1_ in range(default__.abs(125))]): _dafny.CodePoint('o')})
            d_2445_v29_: _dafny.Array
            nw456_ = _dafny.Array(None, 20)
            nw456_[int(0)] = d_2438_v23_
            nw456_[int(1)] = d_2438_v23_
            nw456_[int(2)] = d_2438_v23_
            nw456_[int(3)] = d_2438_v23_
            nw456_[int(4)] = d_2438_v23_
            nw456_[int(5)] = d_2438_v23_
            nw456_[int(6)] = d_2438_v23_
            nw456_[int(7)] = default__.fm29(not(d_2414_v0_), d_2423_v8_, ((d_2439_v24_)[d_2423_v8_] if (d_2423_v8_) in (d_2439_v24_) else len(d_2441_v26_)), d_2414_v0_, globalState)
            nw456_[int(8)] = ((d_2442_v27_)[d_2414_v0_] if (d_2414_v0_) in (d_2442_v27_) else d_2438_v23_)
            nw456_[int(9)] = d_2438_v23_
            nw456_[int(10)] = ((d_2443_v28_)[d_2424_v9_] if (d_2424_v9_) in (d_2443_v28_) else d_2438_v23_)
            nw456_[int(11)] = d_2438_v23_
            nw456_[int(12)] = _dafny.CodePoint('g')
            nw456_[int(13)] = d_2438_v23_
            nw456_[int(14)] = d_2438_v23_
            nw456_[int(15)] = _dafny.CodePoint('p')
            nw456_[int(16)] = d_2438_v23_
            nw456_[int(17)] = d_2438_v23_
            nw456_[int(18)] = _dafny.CodePoint('b')
            nw456_[int(19)] = d_2438_v23_
            d_2445_v29_ = nw456_
            index414_ = default__.safeIndex(974, (d_2445_v29_).length(0))
            (d_2445_v29_)[index414_] = d_2438_v23_
            index415_ = default__.safeIndex(974, (d_2445_v29_).length(0))
            (d_2445_v29_)[index415_] = d_2438_v23_
            d_2446_v30_: D12
            d_2446_v30_ = D12_DC35()
            d_2446_v30_ = d_2446_v30_
            d_2447_v31_: _dafny.Array
            nw457_ = _dafny.Array(False, 24)
            d_2447_v31_ = nw457_
            index416_ = default__.safeIndex(680, (d_2447_v31_).length(0))
            (d_2447_v31_)[index416_] = d_2414_v0_
            index417_ = default__.safeIndex(680, (d_2447_v31_).length(0))
            (d_2447_v31_)[index417_] = d_2414_v0_


class C8:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self):
        pass
        pass

    def fm6(self, p0, p1, globalState):
        return default__.safeModuloInt(-572, len(_dafny.Map({871: 529})))

    def fm7(self, p0, p1, p2, p3, globalState):
        return _dafny.Map({(True) == (False): not((not(False) if False else True))})

    def m5(self, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: bool = False
        d_2448_v0_: bool
        d_2448_v0_ = False
        d_2449_v1_: D0
        d_2449_v1_ = D0_DC0(d_2448_v0_)
        d_2450_v2_: D0
        d_2450_v2_ = D0_DC2(d_2449_v1_)
        d_2451_v3_: D0
        d_2451_v3_ = D0_DC2(d_2449_v1_)
        d_2452_v4_: D0
        d_2452_v4_ = D0_DC2(D0_DC2(d_2449_v1_))
        d_2452_v4_ = d_2452_v4_
        d_2453_v5_: int
        d_2453_v5_ = 844
        d_2454_i0_: int
        d_2454_i0_ = 0
        with _dafny.label("16"):
            while (d_2453_v5_) > ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([not(d_2448_v0_)]))))):
                with _dafny.c_label("16"):
                    if (d_2454_i0_) >= (100):
                        raise _dafny.Break("16")
                    d_2454_i0_ = (d_2454_i0_) + (1)
                    d_2455_v6_: _dafny.Map
                    d_2455_v6_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([-939]): d_2448_v0_})
                    d_2456_v7_: _dafny.Seq
                    d_2456_v7_ = _dafny.SeqWithoutIsStrInference([(d_2453_v5_) >= (len(d_2455_v6_))])
                    d_2457_v8_: _dafny.Seq
                    d_2457_v8_ = _dafny.SeqWithoutIsStrInference([d_2453_v5_])
                    d_2448_v0_ = not(not(not((d_2456_v7_)[default__.safeIndex((d_2457_v8_)[default__.safeIndex(d_2453_v5_, len(d_2457_v8_))], len(d_2456_v7_))])))
                    d_2458_v9_: _dafny.Array
                    nw458_ = _dafny.Array(None, 14)
                    nw458_[int(0)] = d_2448_v0_
                    nw458_[int(1)] = d_2448_v0_
                    nw458_[int(2)] = not(d_2448_v0_)
                    nw458_[int(3)] = d_2448_v0_
                    nw458_[int(4)] = d_2448_v0_
                    nw458_[int(5)] = d_2448_v0_
                    nw458_[int(6)] = d_2448_v0_
                    nw458_[int(7)] = d_2448_v0_
                    nw458_[int(8)] = d_2448_v0_
                    nw458_[int(9)] = False
                    nw458_[int(10)] = d_2448_v0_
                    nw458_[int(11)] = d_2448_v0_
                    nw458_[int(12)] = d_2448_v0_
                    nw458_[int(13)] = d_2448_v0_
                    d_2458_v9_ = nw458_
                    d_2459_v10_: _dafny.Seq
                    d_2459_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umolis"))
                    d_2460_v11_: _dafny.Map
                    d_2460_v11_ = _dafny.Map({d_2458_v9_: (_dafny.MultiSet([d_2453_v5_, len(d_2459_v10_), d_2453_v5_])).cardinality})
                    d_2461_v12_: _dafny.Map
                    d_2461_v12_ = _dafny.Map({not(d_2448_v0_): 868})
                    d_2462_v13_: _dafny.Map
                    d_2462_v13_ = _dafny.Map({((d_2460_v11_)[d_2458_v9_] if (d_2458_v9_) in (d_2460_v11_) else d_2453_v5_): d_2461_v12_})
                    d_2462_v13_ = (d_2462_v13_).set(d_2453_v5_, (_dafny.Map({d_2448_v0_: 845})) | (_dafny.Map({d_2448_v0_: d_2453_v5_})))
                    d_2453_v5_ = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxvm"))), len(d_2456_v7_))
                    d_2463_v14_: _dafny.MultiSet
                    d_2463_v14_ = _dafny.MultiSet([d_2448_v0_, d_2448_v0_])
                    d_2464_v15_: _dafny.MultiSet
                    d_2464_v15_ = _dafny.MultiSet([d_2463_v14_, d_2463_v14_, (d_2463_v14_).set(d_2448_v0_, default__.abs(d_2453_v5_)), d_2463_v14_])
                    d_2465_v16_: C5
                    nw459_ = C5()
                    nw459_.ctor__((d_2464_v15_).cardinality)
                    d_2465_v16_ = nw459_
                    pass
            pass
        d_2466_v17_: _dafny.Seq
        d_2466_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkeuhyxcb"))
        r1 = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2467_i1_ in range(default__.abs(334))])) + (d_2466_v17_)
        d_2468_v18_: _dafny.Array
        def lambda125_(d_2469_v0_):
            def lambda126_(d_2470_i3_):
                return d_2469_v0_

            return lambda126_

        init72_ = lambda125_(d_2448_v0_)
        nw460_ = _dafny.Array(None, 28)
        for i0_72_ in range(nw460_.length(0)):
            nw460_[i0_72_] = init72_(i0_72_)
        d_2468_v18_ = nw460_
        d_2471_v19_: _dafny.MultiSet
        d_2471_v19_ = _dafny.MultiSet([d_2453_v5_])
        hi13_ = ((d_2471_v19_).cardinality) + (d_2453_v5_)
        for d_2472_i2_ in range((d_2453_v5_) - ((_dafny.MultiSet([d_2468_v18_])).cardinality), hi13_):
            if d_2448_v0_:
                (globalState).f6 = (0) - (d_2453_v5_)
                d_2473_v20_: _dafny.Seq
                d_2473_v20_ = _dafny.SeqWithoutIsStrInference([d_2453_v5_])
                d_2474_v21_: str
                d_2474_v21_ = _dafny.CodePoint('c')
                d_2475_v22_: _dafny.Map
                d_2475_v22_ = _dafny.Map({d_2474_v21_: len(default__.fm21(globalState))})
                d_2476_v23_: _dafny.Map
                d_2476_v23_ = _dafny.Map({len(d_2475_v22_): d_2448_v0_})
                d_2477_v24_: C2
                nw461_ = C2()
                nw461_.ctor__((len(d_2473_v20_)) - (len(d_2473_v20_)), default__.safeDivisionInt(280, d_2453_v5_), _dafny.Map({d_2476_v23_: d_2453_v5_}))
                d_2477_v24_ = nw461_
                (globalState).f6 = (self).fm6((d_2448_v0_) == (not(d_2448_v0_)), len(d_2466_v17_), globalState)
                d_2478_v25_: _dafny.Set
                d_2478_v25_ = _dafny.Set({d_2453_v5_})
                d_2474_v21_ = default__.fm29(((d_2473_v20_)[default__.safeIndex((d_2473_v20_)[default__.safeIndex(d_2477_v24_.f21, len(d_2473_v20_))], len(d_2473_v20_))]) <= (d_2453_v5_), len(d_2473_v20_), d_2477_v24_.f21, (d_2478_v25_).issubset(d_2478_v25_), globalState)
                rhs332_ = ((d_2477_v24_).f20 if d_2448_v0_ else d_2453_v5_)
                rhs333_ = d_2448_v0_
                r0 = rhs332_
                d_2448_v0_ = rhs333_
            elif True:
                (globalState).f6 = d_2472_i2_
                d_2448_v0_ = d_2448_v0_
                d_2479_v26_: _dafny.Array
                nw462_ = _dafny.Array(None, 13)
                nw462_[int(0)] = d_2468_v18_
                nw462_[int(1)] = d_2468_v18_
                nw462_[int(2)] = d_2468_v18_
                nw462_[int(3)] = d_2468_v18_
                nw462_[int(4)] = d_2468_v18_
                nw462_[int(5)] = d_2468_v18_
                nw462_[int(6)] = d_2468_v18_
                nw462_[int(7)] = d_2468_v18_
                nw462_[int(8)] = d_2468_v18_
                nw462_[int(9)] = d_2468_v18_
                nw462_[int(10)] = d_2468_v18_
                nw462_[int(11)] = d_2468_v18_
                nw462_[int(12)] = d_2468_v18_
                d_2479_v26_ = nw462_
                index418_ = default__.safeIndex(344, (d_2479_v26_).length(0))
                (d_2479_v26_)[index418_] = d_2468_v18_
                d_2480_v27_: _dafny.Map
                d_2480_v27_ = _dafny.Map({default__.fm2(globalState): d_2468_v18_})
                index419_ = default__.safeIndex(344, (d_2479_v26_).length(0))
                (d_2479_v26_)[index419_] = ((d_2480_v27_)[(len(_dafny.Map({d_2448_v0_: d_2448_v0_}))) <= (d_2453_v5_)] if ((len(_dafny.Map({d_2448_v0_: d_2448_v0_}))) <= (d_2453_v5_)) in (d_2480_v27_) else d_2468_v18_)
                d_2481_v28_: _dafny.Map
                d_2481_v28_ = _dafny.Map({(not(d_2448_v0_) if d_2448_v0_ else d_2448_v0_): d_2448_v0_})
                d_2481_v28_ = (d_2481_v28_).set(d_2448_v0_, d_2448_v0_)
                d_2482_v29_: _dafny.Array
                nw463_ = _dafny.Array(int(0), 17)
                d_2482_v29_ = nw463_
                index420_ = default__.safeIndex(181, (d_2482_v29_).length(0))
                (d_2482_v29_)[index420_] = d_2453_v5_
                index421_ = default__.safeIndex(795, (d_2482_v29_).length(0))
                (d_2482_v29_)[index421_] = (0) - ((self).fm6(not(d_2448_v0_), d_2472_i2_, globalState))
                index422_ = default__.safeIndex(783, (d_2468_v18_).length(0))
                (d_2468_v18_)[index422_] = ((0) - (d_2453_v5_)) >= (d_2453_v5_)
                index423_ = default__.safeIndex(181, (d_2482_v29_).length(0))
                index424_ = default__.safeIndex(795, (d_2482_v29_).length(0))
                index425_ = default__.safeIndex(783, (d_2468_v18_).length(0))
                rhs334_ = d_2472_i2_
                rhs335_ = d_2472_i2_
                rhs336_ = (d_2453_v5_) - (d_2453_v5_)
                rhs337_ = not(d_2448_v0_)
                rhs338_ = d_2472_i2_
                lhs287_ = d_2482_v29_
                lhs288_ = default__.safeIndex(181, (d_2482_v29_).length(0))
                lhs289_ = d_2482_v29_
                lhs290_ = default__.safeIndex(795, (d_2482_v29_).length(0))
                lhs291_ = globalState
                lhs292_ = d_2468_v18_
                lhs293_ = default__.safeIndex(783, (d_2468_v18_).length(0))
                lhs287_[lhs288_] = rhs334_
                lhs289_[lhs290_] = rhs335_
                lhs291_.f6 = rhs336_
                lhs292_[lhs293_] = rhs337_
                d_2453_v5_ = rhs338_
            r0 = d_2453_v5_
            d_2483_v30_: _dafny.Seq
            d_2483_v30_ = _dafny.SeqWithoutIsStrInference([d_2453_v5_])
            d_2484_v31_: _dafny.Set
            d_2484_v31_ = _dafny.Set({d_2472_i2_, len(d_2483_v30_)})
            d_2485_v32_: D15
            d_2485_v32_ = D15_DC42(d_2484_v31_)
            d_2485_v32_ = D15_DC42((d_2484_v31_) - (d_2484_v31_))
            d_2484_v31_ = default__.fm28(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdyvu"))), d_2448_v0_, d_2448_v0_, globalState)
        d_2486_v33_: C6
        nw464_ = C6()
        nw464_.ctor__()
        d_2486_v33_ = nw464_
        d_2487_v34_: _dafny.Array
        nw465_ = _dafny.Array(_dafny.Set({}), 29)
        d_2487_v34_ = nw465_
        d_2488_v35_: _dafny.MultiSet
        d_2488_v35_ = _dafny.MultiSet([True])
        d_2489_v36_: _dafny.Set
        d_2489_v36_ = _dafny.Set({((d_2488_v35_)[d_2448_v0_] if (d_2448_v0_) in (d_2488_v35_) else d_2453_v5_)})
        d_2490_v37_: _dafny.Map
        d_2490_v37_ = _dafny.Map({d_2487_v34_: d_2489_v36_})
        d_2490_v37_ = (d_2490_v37_).set(d_2487_v34_, d_2489_v36_)
        r0 = (d_2453_v5_) * (default__.safeModuloInt(d_2453_v5_, d_2453_v5_))
        d_2491_v38_: _dafny.Map
        d_2491_v38_ = _dafny.Map({(0) - (len(d_2466_v17_)): d_2489_v36_})
        pat_let_tv61_ = d_2466_v17_
        def lambda127_(source23_):
            d_2492___mcc_h0_ = source23_
            d_2493_cf19_ = d_2492___mcc_h0_
            return pat_let_tv61_

        r1 = lambda127_(d_2491_v38_)
        d_2494_v39_: _dafny.Seq
        d_2494_v39_ = _dafny.SeqWithoutIsStrInference([d_2448_v0_])
        r2 = not(not ((d_2494_v39_)[default__.safeIndex(d_2453_v5_, len(d_2494_v39_))]) or ((d_2448_v0_) or (d_2448_v0_)))
        return r0, r1, r2

    def m6(self, p0, p1, globalState):
        d_2495_v0_: int
        d_2495_v0_ = 150
        d_2496_v1_: _dafny.Seq
        d_2496_v1_ = _dafny.SeqWithoutIsStrInference([d_2495_v0_])
        d_2497_v2_: D10
        d_2497_v2_ = D10_DC28(len(d_2496_v1_))
        hi14_ = -270
        for d_2498_i0_ in range(default__.safeDivisionInt(d_2495_v0_, (d_2497_v2_).cf51), hi14_):
            d_2499_v3_: C5
            nw466_ = C5()
            nw466_.ctor__(d_2495_v0_)
            d_2499_v3_ = nw466_
            d_2500_v4_: _dafny.Map
            d_2500_v4_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference([d_2498_i0_])})
            d_2496_v1_ = (((d_2500_v4_)[p0] if (p0) in (d_2500_v4_) else d_2496_v1_)) + (d_2496_v1_)
            d_2501_v5_: C1
            nw467_ = C1()
            nw467_.ctor__(d_2495_v0_)
            d_2501_v5_ = nw467_
            if p0:
                d_2502_v6_: str
                d_2502_v6_ = _dafny.CodePoint('c')
                d_2503_v7_: _dafny.Array
                nw468_ = _dafny.Array(None, 8)
                nw468_[int(0)] = d_2502_v6_
                nw468_[int(1)] = default__.fm29(p0, (d_2499_v3_).f16, (d_2501_v5_).f24, p0, globalState)
                nw468_[int(2)] = d_2502_v6_
                nw468_[int(3)] = d_2502_v6_
                nw468_[int(4)] = _dafny.CodePoint('n')
                nw468_[int(5)] = _dafny.CodePoint('t')
                nw468_[int(6)] = d_2502_v6_
                nw468_[int(7)] = _dafny.CodePoint('b')
                d_2503_v7_ = nw468_
                index426_ = default__.safeIndex(702, (d_2503_v7_).length(0))
                (d_2503_v7_)[index426_] = d_2502_v6_
                index427_ = default__.safeIndex(702, (d_2503_v7_).length(0))
                (d_2503_v7_)[index427_] = _dafny.CodePoint('a')
                d_2504_v8_: _dafny.Array
                def lambda128_(d_2505_v0_):
                    def lambda129_(d_2506_i1_):
                        return default__.safeDivisionInt(d_2506_i1_, d_2505_v0_)

                    return lambda129_

                init73_ = lambda128_(d_2495_v0_)
                nw469_ = _dafny.Array(None, 29)
                for i0_73_ in range(nw469_.length(0)):
                    nw469_[i0_73_] = init73_(i0_73_)
                d_2504_v8_ = nw469_
                d_2507_v9_: D3
                d_2507_v9_ = D3_DC6(d_2504_v8_)
                d_2508_v10_: _dafny.Seq
                d_2508_v10_ = _dafny.SeqWithoutIsStrInference([d_2507_v9_, d_2507_v9_, d_2507_v9_, d_2507_v9_, D3_DC6(d_2504_v8_)])
                d_2509_v11_: _dafny.Seq
                d_2509_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iivhhe"))
                d_2510_v12_: C0
                nw470_ = C0()
                nw470_.ctor__((d_2508_v10_) + (d_2508_v10_), (d_2498_i0_) <= (len(d_2509_v11_)))
                d_2510_v12_ = nw470_
                d_2511_v13_: _dafny.MultiSet
                d_2511_v13_ = _dafny.MultiSet([(self).fm6(p0, len(d_2496_v1_), globalState)])
                d_2512_v14_: _dafny.MultiSet
                d_2512_v14_ = d_2511_v13_
                d_2513_v15_: _dafny.Array
                nw471_ = _dafny.Array(None, 7)
                nw471_[int(0)] = d_2511_v13_
                nw471_[int(1)] = d_2511_v13_
                nw471_[int(2)] = (_dafny.MultiSet(d_2496_v1_)) | (d_2511_v13_)
                nw471_[int(3)] = d_2511_v13_
                nw471_[int(4)] = (d_2512_v14_)
                nw471_[int(5)] = d_2511_v13_
                nw471_[int(6)] = d_2511_v13_
                d_2513_v15_ = nw471_
                d_2513_v15_ = d_2513_v15_
                d_2514_v16_: _dafny.Seq
                d_2514_v16_ = _dafny.SeqWithoutIsStrInference([d_2511_v13_, d_2511_v13_])
                d_2515_v17_: _dafny.Map
                d_2515_v17_ = _dafny.Map({d_2498_i0_: (d_2514_v16_)[default__.safeIndex((d_2501_v5_).f24, len(d_2514_v16_))]})
                (globalState).f6 = (((self).fm6((d_2510_v12_).f23, len(d_2496_v1_), globalState)) + (len(d_2515_v17_))) * (((d_2499_v3_).f16) * (d_2498_i0_))
                d_2516_v18_: C6
                nw472_ = C6()
                nw472_.ctor__()
                d_2516_v18_ = nw472_
            elif True:
                d_2517_v19_: _dafny.Seq
                d_2517_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flk"))
                d_2517_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olvqfpuoo"))
                d_2518_v20_: _dafny.Array
                nw473_ = _dafny.Array(int(0), 26)
                d_2518_v20_ = nw473_
                d_2519_v21_: D3
                d_2519_v21_ = D3_DC6(d_2518_v20_)
                d_2520_v22_: _dafny.Seq
                d_2520_v22_ = _dafny.SeqWithoutIsStrInference([D3_DC6(d_2518_v20_), d_2519_v21_])
                d_2521_v23_: C0
                nw474_ = C0()
                nw474_.ctor__(d_2520_v22_, p0)
                d_2521_v23_ = nw474_
                d_2522_v24_: _dafny.Array
                nw475_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_2522_v24_ = nw475_
                d_2523_v25_: D5
                d_2523_v25_ = D5_DC10(d_2522_v24_)
                d_2524_v26_: D7
                d_2524_v26_ = D7_DC17(p0, (p1 if p0 else p1), d_2521_v23_, d_2523_v25_, (p1) and (p0))
                d_2525_v27_: _dafny.Map
                d_2525_v27_ = _dafny.Map({not((d_2521_v23_).f23): d_2517_v19_})
                d_2524_v26_ = D7_DC17((d_2517_v19_) < (((d_2525_v27_)[True] if (True) in (d_2525_v27_) else d_2517_v19_)), ((d_2521_v23_).f23) and (p1), d_2521_v23_, d_2523_v25_, True)
                d_2495_v0_ = (d_2499_v3_).f16
                d_2526_v28_: _dafny.Set
                d_2526_v28_ = _dafny.Set({(d_2521_v23_).f23, p0})
                d_2527_v30_: _dafny.Map
                d_2527_v30_ = _dafny.Map({True: (d_2499_v3_).f16})
                d_2528_v31_: _dafny.MultiSet
                d_2528_v31_ = _dafny.MultiSet([len(d_2527_v30_)])
                d_2529_v32_: _dafny.Set
                d_2529_v32_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttqx"))), (d_2499_v3_).f16, ((d_2528_v31_)[d_2495_v0_] if (d_2495_v0_) in (d_2528_v31_) else (d_2499_v3_).f16), default__.fm1(p1, globalState)})
                d_2530_v33_: _dafny.Array
                nw476_ = _dafny.Array(None, 15)
                nw476_[int(0)] = (d_2517_v19_) != (d_2517_v19_)
                nw476_[int(1)] = not(p1)
                nw476_[int(2)] = p0
                nw476_[int(3)] = (d_2521_v23_).f23
                nw476_[int(4)] = (d_2521_v23_).fm17((d_2501_v5_).f24, d_2495_v0_, p0, globalState)
                nw476_[int(5)] = True
                nw476_[int(6)] = ((d_2499_v3_).f16) < ((d_2499_v3_).f16)
                nw476_[int(7)] = p0
                nw476_[int(8)] = not((d_2521_v23_).f23)
                nw476_[int(9)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2531_i2_ in range(default__.abs(326))])) < (d_2517_v19_)
                nw476_[int(10)] = (d_2521_v23_).f23
                nw476_[int(11)] = default__.fm2(globalState)
                nw476_[int(12)] = (d_2521_v23_).f23
                nw476_[int(13)] = not((False) not in (d_2526_v28_))
                def iife257_():
                    coll83_ = _dafny.Set()
                    compr_83_: int
                    for compr_83_ in _dafny.IntegerRange(136, 970):
                        d_2532_v29_: int = compr_83_
                        if ((136) <= (d_2532_v29_)) and ((d_2532_v29_) < (970)):
                            coll83_ = coll83_.union(_dafny.Set([default__.safeModuloInt(d_2532_v29_, (_dafny.MultiSet([408])).cardinality)]))
                    return _dafny.Set(coll83_)
                nw476_[int(14)] = not((iife257_()
                ).isdisjoint(d_2529_v32_))
                d_2530_v33_ = nw476_
                d_2533_v34_: _dafny.Seq
                d_2533_v34_ = _dafny.SeqWithoutIsStrInference([p1])
                index428_ = default__.safeIndex(699, (d_2530_v33_).length(0))
                (d_2530_v33_)[index428_] = (d_2533_v34_) == (d_2533_v34_)
                index429_ = default__.safeIndex(699, (d_2530_v33_).length(0))
                (d_2530_v33_)[index429_] = (d_2498_i0_) <= (default__.safeDivisionInt((0) - ((d_2499_v3_).f16), d_2495_v0_))
                (globalState).f6 = (d_2501_v5_).f24
        d_2495_v0_ = -736
        d_2534_v35_: _dafny.Map
        d_2534_v35_ = _dafny.Map({d_2495_v0_: -10})
        d_2535_v36_: C7
        nw477_ = C7()
        nw477_.ctor__()
        d_2535_v36_ = nw477_
        d_2536_v37_: _dafny.MultiSet
        d_2536_v37_ = _dafny.MultiSet([d_2535_v36_, d_2535_v36_, d_2535_v36_, d_2535_v36_, d_2535_v36_])
        d_2537_v38_: _dafny.Map
        d_2537_v38_ = _dafny.Map({p0: d_2495_v0_})
        d_2538_v39_: _dafny.Seq
        d_2538_v39_ = _dafny.SeqWithoutIsStrInference([not(p1), False])
        d_2539_v40_: _dafny.Seq
        d_2539_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
        d_2540_v41_: _dafny.MultiSet
        d_2540_v41_ = _dafny.MultiSet([d_2495_v0_, d_2495_v0_, 21])
        d_2541_v42_: _dafny.Map
        d_2541_v42_ = _dafny.Map({d_2495_v0_: d_2540_v41_})
        d_2542_v43_: _dafny.Set
        d_2542_v43_ = _dafny.Set({d_2495_v0_, d_2495_v0_})
        d_2543_v44_: _dafny.Array
        nw478_ = _dafny.Array(None, 19)
        nw478_[int(0)] = (D0_DC1(len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpdpge")))})))).cf1
        nw478_[int(1)] = d_2495_v0_
        nw478_[int(2)] = ((d_2534_v35_)[(0) - (d_2495_v0_)] if ((0) - (d_2495_v0_)) in (d_2534_v35_) else d_2495_v0_)
        nw478_[int(3)] = d_2495_v0_
        nw478_[int(4)] = d_2495_v0_
        nw478_[int(5)] = d_2495_v0_
        nw478_[int(6)] = ((d_2536_v37_)[d_2535_v36_] if (d_2535_v36_) in (d_2536_v37_) else d_2495_v0_)
        nw478_[int(7)] = ((d_2537_v38_)[False] if (False) in (d_2537_v38_) else 646)
        nw478_[int(8)] = d_2495_v0_
        nw478_[int(9)] = (d_2495_v0_) * (len(d_2538_v39_))
        nw478_[int(10)] = -468
        nw478_[int(11)] = (0) - (d_2495_v0_)
        nw478_[int(12)] = (d_2495_v0_) + (d_2495_v0_)
        nw478_[int(13)] = (0) - (len(_dafny.Set({d_2539_v40_, d_2539_v40_})))
        nw478_[int(14)] = (0) - (((d_2540_v41_)[d_2495_v0_] if (d_2495_v0_) in (d_2540_v41_) else d_2495_v0_))
        nw478_[int(15)] = d_2495_v0_
        nw478_[int(16)] = (len(d_2541_v42_)) * (d_2495_v0_)
        nw478_[int(17)] = len(d_2542_v43_)
        nw478_[int(18)] = d_2495_v0_
        d_2543_v44_ = nw478_
        index430_ = default__.safeIndex(956, (d_2543_v44_).length(0))
        (d_2543_v44_)[index430_] = d_2495_v0_
        d_2544_v45_: str
        d_2544_v45_ = _dafny.CodePoint('m')
        index431_ = default__.safeIndex(956, (d_2543_v44_).length(0))
        (d_2543_v44_)[index431_] = len(((d_2539_v40_) + (d_2539_v40_)).set(default__.safeIndex(d_2495_v0_, len((d_2539_v40_) + (d_2539_v40_))), d_2544_v45_))
        d_2545_v46_: _dafny.Array
        nw479_ = _dafny.Array(None, 5)
        d_2545_v46_ = nw479_
        d_2546_v47_: _dafny.Map
        d_2546_v47_ = _dafny.Map({d_2495_v0_: p0})
        d_2547_v48_: _dafny.Map
        d_2547_v48_ = _dafny.Map({d_2546_v47_: (d_2543_v44_)[default__.safeIndex(956, (d_2543_v44_).length(0))]})
        d_2548_v49_: T1
        nw480_ = C2()
        nw480_.ctor__(d_2495_v0_, len(d_2496_v1_), d_2547_v48_)
        d_2548_v49_ = nw480_
        index432_ = default__.safeIndex(261, (d_2545_v46_).length(0))
        (d_2545_v46_)[index432_] = d_2548_v49_
        index433_ = default__.safeIndex(261, (d_2545_v46_).length(0))
        (d_2545_v46_)[index433_] = d_2548_v49_
        def iife259_():
            coll85_ = _dafny.Set()
            compr_85_: int
            for compr_85_ in _dafny.IntegerRange(426, 668):
                d_2556_v50_: int = compr_85_
                if ((426) <= (d_2556_v50_)) and ((d_2556_v50_) < (668)):
                    coll85_ = coll85_.union(_dafny.Set([default__.safeModuloInt(d_2556_v50_, (d_2543_v44_)[default__.safeIndex(956, (d_2543_v44_).length(0))])]))
            return _dafny.Set(coll85_)
        hi15_ = (d_2543_v44_)[default__.safeIndex(956, (d_2543_v44_).length(0))]
        for d_2549_i3_ in range(default__.safeModuloInt(len(iife259_()
        ), d_2495_v0_), hi15_):
            d_2550_v51_: _dafny.Array
            nw481_ = _dafny.Array(None, 7)
            nw481_[int(0)] = d_2539_v40_
            nw481_[int(1)] = d_2539_v40_
            nw481_[int(2)] = ((d_2539_v40_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "io"))), len(d_2539_v40_)), _dafny.CodePoint('g')) if True else d_2539_v40_)
            nw481_[int(3)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2551_i4_ in range(default__.abs(504))])
            nw481_[int(4)] = d_2539_v40_
            nw481_[int(5)] = (d_2539_v40_) + (default__.fm21(globalState))
            nw481_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            d_2550_v51_ = nw481_
            d_2550_v51_ = d_2550_v51_
            (globalState).f6 = d_2495_v0_
            d_2552_v52_: C7
            nw482_ = C7()
            nw482_.ctor__()
            d_2552_v52_ = nw482_
            d_2553_v53_: _dafny.Map
            d_2553_v53_ = _dafny.Map({(d_2543_v44_)[default__.safeIndex(956, (d_2543_v44_).length(0))]: d_2537_v38_})
            d_2554_v55_: D25
            def iife258_():
                coll84_ = _dafny.Map()
                compr_84_: int
                for compr_84_ in _dafny.IntegerRange(-418, -301):
                    d_2555_v54_: int = compr_84_
                    if ((-418) <= (d_2555_v54_)) and ((d_2555_v54_) < (-301)):
                        coll84_[default__.safeDivisionInt(d_2555_v54_, (d_2543_v44_)[default__.safeIndex(956, (d_2543_v44_).length(0))])] = d_2496_v1_
                return _dafny.Map(coll84_)
            d_2554_v55_ = D25_DC57(iife258_()
)
            d_2553_v53_ = (d_2553_v53_).set((0) - (len(default__.fm21(globalState))), (d_2537_v38_) | (default__.fm14((d_2554_v55_).cf96, d_2495_v0_, True, globalState)))
        d_2557_i5_: int
        d_2557_i5_ = 0
        with _dafny.label("17"):
            while p0:
                with _dafny.c_label("17"):
                    if (d_2557_i5_) >= (100):
                        raise _dafny.Break("17")
                    d_2557_i5_ = (d_2557_i5_) + (1)
                    (globalState).f6 = ((d_2543_v44_)[default__.safeIndex(956, (d_2543_v44_).length(0))]) + ((d_2543_v44_)[default__.safeIndex(956, (d_2543_v44_).length(0))])
                    d_2558_v56_: C2
                    nw483_ = C2()
                    nw483_.ctor__(default__.fm1((d_2538_v39_)[default__.safeIndex((d_2543_v44_)[default__.safeIndex(956, (d_2543_v44_).length(0))], len(d_2538_v39_))], globalState), len((_dafny.Map({d_2495_v0_: len(d_2496_v1_)})) | (_dafny.Map({d_2495_v0_: d_2495_v0_}))), (d_2547_v48_) | (d_2548_v49_.f17))
                    d_2558_v56_ = nw483_
                    d_2559_v58_: _dafny.Map
                    d_2559_v58_ = _dafny.Map({not(p0): p0})
                    def iife260_():
                        coll86_ = _dafny.Set()
                        compr_86_: int
                        for compr_86_ in _dafny.IntegerRange(78, 315):
                            d_2560_v57_: int = compr_86_
                            if ((78) <= (d_2560_v57_)) and ((d_2560_v57_) < (315)):
                                coll86_ = coll86_.union(_dafny.Set([(d_2560_v57_) + (len(d_2539_v40_))]))
                        return _dafny.Set(coll86_)
                    (globalState).f1 = (d_2558_v56_.f21) not in ((iife260_()
                    ) - (_dafny.Set({len(d_2559_v58_)})))
                    d_2561_v59_: _dafny.Seq
                    d_2561_v59_ = _dafny.SeqWithoutIsStrInference([d_2558_v56_.f21])
                    index434_ = default__.safeIndex(956, (d_2543_v44_).length(0))
                    rhs339_ = (d_2543_v44_)[default__.safeIndex(956, (d_2543_v44_).length(0))]
                    rhs340_ = ((d_2558_v56_).fm15(globalState)) - ((0) - (len((d_2561_v59_))))
                    lhs294_ = globalState
                    lhs295_ = d_2543_v44_
                    lhs296_ = default__.safeIndex(956, (d_2543_v44_).length(0))
                    lhs294_.f6 = rhs339_
                    lhs295_[lhs296_] = rhs340_
                    pass
            pass


class C9(T0):
    def  __init__(self):
        self._f15: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    def ctor__(self, f15):
        (self)._f15 = f15

    def fm3(self, p0, p1, p2, p3, globalState):
        source24_ = D0_DC0(not(True))
        if source24_.is_DC1:
            d_2562___mcc_h0_ = source24_.cf1
            d_2563_cf1_ = d_2562___mcc_h0_
            return D0_DC2(D0_DC2(D0_DC1((self).f15)))
        elif source24_.is_DC0:
            d_2564___mcc_h1_ = source24_.cf0
            d_2565_cf0_ = d_2564___mcc_h1_
            if d_2565_cf0_:
                return D0_DC2(D0_DC2(D0_DC2(D0_DC2(D0_DC1((0) - ((self).f15))))))
            elif True:
                return D0_DC2(D0_DC1((self).f15))
        elif True:
            d_2566___mcc_h2_ = source24_.cf2
            d_2567_cf2_ = d_2566___mcc_h2_
            return D0_DC2(D0_DC1((self).f15))

    def fm4(self, p0, globalState):
        return (((D0_DC1((self).f15)).cf1) + (len(_dafny.SeqWithoutIsStrInference([(self).f15 for d_2568_i0_ in range(default__.abs(892))])))) <= ((len(_dafny.Map({(0) - ((self).f15): True}))) - ((0) - ((0) - ((self).f15))))

    def fm5(self, p0, globalState):
        return ((self).f15) + ((self).f15)

    def m1(self, globalState):
        r0: int = int(0)
        r1: bool = False
        d_2569_v0_: bool
        d_2569_v0_ = False
        d_2570_i0_: int
        d_2570_i0_ = 0
        with _dafny.label("18"):
            while d_2569_v0_:
                with _dafny.c_label("18"):
                    if (d_2570_i0_) >= (100):
                        raise _dafny.Break("18")
                    d_2570_i0_ = (d_2570_i0_) + (1)
                    d_2571_v1_: _dafny.Seq
                    d_2571_v1_ = _dafny.SeqWithoutIsStrInference([d_2569_v0_, d_2569_v0_])
                    d_2572_v2_: bool
                    d_2573_v3_: bool
                    d_2574_v4_: int
                    out18_: bool
                    out19_: bool
                    out20_: int
                    out18_, out19_, out20_ = (self).m3((d_2571_v1_) + (d_2571_v1_), globalState)
                    d_2572_v2_ = out18_
                    d_2573_v3_ = out19_
                    d_2574_v4_ = out20_
                    d_2575_v5_: _dafny.Set
                    d_2575_v5_ = _dafny.Set({d_2573_v3_, True, d_2572_v2_, d_2569_v0_})
                    d_2576_v6_: _dafny.Seq
                    d_2576_v6_ = _dafny.SeqWithoutIsStrInference([len(d_2575_v5_), default__.safeModuloInt(781, (0) - (len(d_2571_v1_))), ((self).f15) + (d_2574_v4_), (self).f15])
                    d_2576_v6_ = _dafny.SeqWithoutIsStrInference([((_dafny.MultiSet(d_2571_v1_)).set(d_2572_v2_, default__.abs((0) - (d_2574_v4_)))).cardinality, (0) - ((0) - (d_2574_v4_)), ((self).f15 if d_2572_v2_ else (self).f15)])
                    d_2577_v7_: C5
                    nw484_ = C5()
                    nw484_.ctor__((self).f15)
                    d_2577_v7_ = nw484_
                    d_2578_v8_: _dafny.Map
                    d_2578_v8_ = _dafny.Map({561: (self).f15})
                    d_2578_v8_ = (d_2578_v8_).set((0) - ((self).f15), (self).f15)
                    pass
            pass
        d_2579_v9_: _dafny.Map
        d_2579_v9_ = _dafny.Map({d_2569_v0_: d_2569_v0_})
        d_2580_v10_: _dafny.MultiSet
        d_2580_v10_ = _dafny.MultiSet([_dafny.Map({d_2569_v0_: d_2569_v0_}), d_2579_v9_])
        hi16_ = (self).f15
        for d_2581_i1_ in range((0) - (((d_2580_v10_) | (d_2580_v10_)).cardinality), hi16_):
            d_2582_v11_: _dafny.Seq
            d_2582_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epptbqs"))
            d_2583_v12_: _dafny.Seq
            d_2583_v12_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egwsccse"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scuxx"))), d_2582_v11_])
            d_2584_v13_: D13
            d_2584_v13_ = D13_DC37(d_2569_v0_, d_2569_v0_)
            d_2585_v14_: _dafny.Map
            d_2585_v14_ = _dafny.Map({(0) - (d_2581_i1_): d_2569_v0_})
            rhs341_ = (d_2569_v0_ if not(d_2569_v0_) else False)
            rhs342_ = (d_2569_v0_) or (not(not (d_2569_v0_) or (d_2569_v0_)))
            rhs343_ = d_2583_v12_
            rhs344_ = 815
            rhs345_ = not((len(default__.fm0(not((d_2584_v13_).cf66), True, (self).f15, globalState))) >= (len(d_2585_v14_)))
            lhs297_ = globalState
            d_2569_v0_ = rhs341_
            d_2569_v0_ = rhs342_
            d_2583_v12_ = rhs343_
            lhs297_.f6 = rhs344_
            r1 = rhs345_
            (globalState).f6 = (self).f15
            d_2586_v15_: _dafny.Set
            d_2586_v15_ = _dafny.Set({(self).f15, len(d_2582_v11_), (self).f15, d_2581_i1_, 21})
            d_2587_v16_: D0
            d_2587_v16_ = D0_DC2(D0_DC2((D0_DC2(D0_DC1((self).f15))).cf2))
            d_2588_v17_: D0
            d_2588_v17_ = D0_DC0(d_2569_v0_)
            d_2589_v18_: _dafny.Array
            nw485_ = _dafny.Array(None, 27)
            nw485_[int(0)] = d_2587_v16_
            nw485_[int(1)] = D0_DC2(d_2588_v17_)
            nw485_[int(2)] = d_2587_v16_
            nw485_[int(3)] = d_2587_v16_
            nw485_[int(4)] = D0_DC2(d_2588_v17_)
            nw485_[int(5)] = D0_DC2(d_2588_v17_)
            nw485_[int(6)] = D0_DC2(d_2588_v17_)
            nw485_[int(7)] = D0_DC2(d_2588_v17_)
            nw485_[int(8)] = d_2587_v16_
            nw485_[int(9)] = d_2587_v16_
            nw485_[int(10)] = d_2587_v16_
            nw485_[int(11)] = d_2587_v16_
            nw485_[int(12)] = d_2587_v16_
            nw485_[int(13)] = d_2587_v16_
            nw485_[int(14)] = (self).fm3(d_2569_v0_, (self).f15, d_2569_v0_, d_2569_v0_, globalState)
            nw485_[int(15)] = (self).fm3(not(d_2569_v0_), d_2581_i1_, d_2569_v0_, not(d_2569_v0_), globalState)
            nw485_[int(16)] = d_2587_v16_
            nw485_[int(17)] = d_2587_v16_
            nw485_[int(18)] = d_2587_v16_
            nw485_[int(19)] = D0_DC2(d_2588_v17_)
            nw485_[int(20)] = (self).fm3(False, (self).f15, d_2569_v0_, d_2569_v0_, globalState)
            nw485_[int(21)] = d_2587_v16_
            nw485_[int(22)] = d_2587_v16_
            nw485_[int(23)] = d_2587_v16_
            nw485_[int(24)] = d_2587_v16_
            nw485_[int(25)] = (self).fm3(d_2569_v0_, (self).f15, d_2569_v0_, not(d_2569_v0_), globalState)
            nw485_[int(26)] = d_2587_v16_
            d_2589_v18_ = nw485_
            d_2590_v19_: _dafny.Array
            def lambda130_(d_2591_v14_):
                def lambda131_(d_2592_i2_):
                    return (d_2592_i2_) + (len(d_2591_v14_))

                return lambda131_

            init74_ = lambda130_(d_2585_v14_)
            nw486_ = _dafny.Array(None, 6)
            for i0_74_ in range(nw486_.length(0)):
                nw486_[i0_74_] = init74_(i0_74_)
            d_2590_v19_ = nw486_
            default__.m0(d_2586_v15_, d_2589_v18_, d_2590_v19_, d_2582_v11_, globalState)
            if d_2569_v0_:
                index435_ = default__.safeIndex(476, (d_2590_v19_).length(0))
                (d_2590_v19_)[index435_] = d_2581_i1_
                index436_ = default__.safeIndex(476, (d_2590_v19_).length(0))
                (d_2590_v19_)[index436_] = default__.safeModuloInt(d_2581_i1_, (d_2581_i1_) - ((self).f15))
                r1 = d_2569_v0_
                d_2593_v20_: str
                d_2593_v20_ = _dafny.CodePoint('y')
                d_2594_v21_: bool
                d_2595_v22_: bool
                out21_: bool
                out22_: bool
                out21_, out22_ = (self).m4((d_2593_v20_) in (d_2582_v11_), globalState)
                d_2594_v21_ = out21_
                d_2595_v22_ = out22_
                d_2596_v23_: _dafny.Array
                nw487_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
                d_2596_v23_ = nw487_
                d_2597_v24_: _dafny.Map
                d_2597_v24_ = _dafny.Map({d_2593_v20_: d_2596_v23_})
                d_2597_v24_ = (d_2597_v24_).set(d_2593_v20_, d_2596_v23_)
                index437_ = default__.safeIndex(476, (d_2590_v19_).length(0))
                (d_2590_v19_)[index437_] = (d_2590_v19_)[default__.safeIndex(476, (d_2590_v19_).length(0))]
            elif True:
                r1 = (d_2586_v15_).ispropersubset(_dafny.Set({d_2581_i1_, d_2581_i1_, default__.fm1(d_2569_v0_, globalState)}))
                d_2598_v25_: C5
                nw488_ = C5()
                nw488_.ctor__(((self).f15) * (d_2581_i1_))
                d_2598_v25_ = nw488_
                d_2599_v26_: _dafny.Seq
                d_2599_v26_ = _dafny.SeqWithoutIsStrInference([d_2581_i1_, d_2581_i1_, (d_2598_v25_).f16])
                (globalState).f6 = (d_2581_i1_) - ((d_2599_v26_)[default__.safeIndex(d_2581_i1_, len(d_2599_v26_))])
                d_2600_v27_: str
                d_2600_v27_ = _dafny.CodePoint('p')
                rhs346_ = (D3_DC8(d_2599_v26_, d_2582_v11_, d_2569_v0_, d_2600_v27_)).cf9
                rhs347_ = (d_2599_v26_) + (d_2599_v26_)
                d_2582_v11_ = rhs346_
                d_2599_v26_ = rhs347_
                default__.m0(d_2586_v15_, d_2589_v18_, d_2590_v19_, (d_2582_v11_ if not(d_2569_v0_) else d_2582_v11_), globalState)
        d_2601_v28_: _dafny.Seq
        d_2601_v28_ = _dafny.SeqWithoutIsStrInference([(self).f15])
        d_2602_v29_: _dafny.Seq
        d_2602_v29_ = _dafny.SeqWithoutIsStrInference([(len(d_2601_v28_)) < ((self).f15)])
        if (d_2602_v29_)[default__.safeIndex((self).f15, len(d_2602_v29_))]:
            d_2603_v30_: _dafny.Array
            def lambda132_(d_2604_i3_):
                return default__.safeDivisionInt(d_2604_i3_, (self).f15)

            init75_ = lambda132_
            nw489_ = _dafny.Array(None, 12)
            for i0_75_ in range(nw489_.length(0)):
                nw489_[i0_75_] = init75_(i0_75_)
            d_2603_v30_ = nw489_
            d_2605_v31_: _dafny.Array
            nw490_ = _dafny.Array(False, 24)
            d_2605_v31_ = nw490_
            d_2606_v32_: _dafny.Map
            d_2606_v32_ = _dafny.Map({d_2603_v30_: d_2605_v31_})
            pat_let_tv62_ = d_2603_v30_
            def iife261_(_pat_let87_0):
                def iife262_(d_2607_dt__update__tmp_h0_):
                    def iife263_(_pat_let88_0):
                        def iife264_(d_2608_dt__update_hcf6_h0_):
                            return D3_DC6(d_2608_dt__update_hcf6_h0_)
                        return iife264_(_pat_let88_0)
                    return iife263_(pat_let_tv62_)
                return iife262_(_pat_let87_0)
            d_2606_v32_ = (d_2606_v32_) | (_dafny.Map({(iife261_(D3_DC6(d_2603_v30_))).cf6: d_2605_v31_}))
            d_2579_v9_ = (d_2579_v9_).set(d_2569_v0_, d_2569_v0_)
            (globalState).f1 = d_2569_v0_
            if d_2569_v0_:
                (globalState).f1 = d_2569_v0_
                d_2609_v33_: _dafny.Array
                nw491_ = _dafny.Array(None, 7)
                d_2609_v33_ = nw491_
                d_2610_v34_: C4
                nw492_ = C4()
                nw492_.ctor__()
                d_2610_v34_ = nw492_
                index438_ = default__.safeIndex(625, (d_2609_v33_).length(0))
                (d_2609_v33_)[index438_] = d_2610_v34_
                d_2611_v35_: _dafny.Seq
                d_2611_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wlaqsdmhf"))
                d_2612_v36_: C4
                d_2612_v36_ = d_2610_v34_
                index439_ = default__.safeIndex(625, (d_2609_v33_).length(0))
                rhs348_ = (d_2612_v36_)
                rhs349_ = d_2569_v0_
                rhs350_ = d_2611_v35_
                lhs298_ = d_2609_v33_
                lhs299_ = default__.safeIndex(625, (d_2609_v33_).length(0))
                lhs300_ = globalState
                lhs298_[lhs299_] = rhs348_
                lhs300_.f1 = rhs349_
                d_2611_v35_ = rhs350_
                d_2613_v37_: _dafny.Map
                d_2613_v37_ = _dafny.Map({(self).f15: (self).f15})
                d_2614_v38_: _dafny.Map
                d_2614_v38_ = _dafny.Map({d_2613_v37_: d_2569_v0_})
                d_2615_v40_: _dafny.Map
                d_2615_v40_ = _dafny.Map({_dafny.Map({(self).f15: (self).f15}): len(d_2613_v37_)})
                def iife265_():
                    coll87_ = _dafny.Map()
                    compr_87_: _dafny.Map
                    for compr_87_ in (d_2615_v40_).keys.Elements:
                        d_2616_v39_: _dafny.Map = compr_87_
                        if (d_2616_v39_) in (d_2615_v40_):
                            coll87_[d_2616_v39_] = d_2569_v0_
                    return _dafny.Map(coll87_)
                d_2569_v0_ = ((d_2614_v38_) | (_dafny.Map({d_2613_v37_: d_2569_v0_}))) != ((iife265_()
                ) | (d_2614_v38_))
                d_2569_v0_ = not(True)
                (globalState).f1 = d_2569_v0_
            elif True:
                d_2601_v28_ = _dafny.SeqWithoutIsStrInference([(self).f15 for d_2617_i4_ in range(default__.abs(415))])
                d_2618_v41_: _dafny.Set
                d_2618_v41_ = _dafny.Set({(0) - ((self).f15), (0) - ((self).f15)})
                (globalState).f6 = (len(d_2618_v41_)) * ((self).f15)
                d_2619_v42_: _dafny.Array
                nw493_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
                d_2619_v42_ = nw493_
                d_2620_v43_: _dafny.Seq
                d_2620_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eioyj"))
                index440_ = default__.safeIndex(752, (d_2619_v42_).length(0))
                (d_2619_v42_)[index440_] = d_2620_v43_
                index441_ = default__.safeIndex(752, (d_2619_v42_).length(0))
                (d_2619_v42_)[index441_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sewysnmpc"))
                d_2620_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udky"))
                d_2621_v44_: bool
                d_2622_v45_: bool
                d_2623_v46_: int
                out23_: bool
                out24_: bool
                out25_: int
                out23_, out24_, out25_ = (self).m3(_dafny.SeqWithoutIsStrInference([d_2569_v0_]), globalState)
                d_2621_v44_ = out23_
                d_2622_v45_ = out24_
                d_2623_v46_ = out25_
            r0 = (d_2601_v28_)[default__.safeIndex(default__.safeDivisionInt((self).f15, -890), len(d_2601_v28_))]
        elif True:
            d_2624_v47_: T0
            nw494_ = C4()
            nw494_.ctor__()
            d_2624_v47_ = nw494_
            d_2625_v48_: _dafny.Seq
            d_2625_v48_ = _dafny.SeqWithoutIsStrInference([d_2624_v47_])
            d_2579_v9_ = (d_2579_v9_).set(d_2569_v0_, ((self).f15) != (len((d_2625_v48_).set(default__.safeIndex((d_2601_v28_)[default__.safeIndex((self).f15, len(d_2601_v28_))], len(d_2625_v48_)), d_2624_v47_))))
            d_2626_v49_: _dafny.Seq
            d_2626_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hutlw"))
            d_2627_v50_: _dafny.Seq
            d_2627_v50_ = _dafny.SeqWithoutIsStrInference([d_2602_v29_])
            (globalState).f6 = (((self).f15) + (len(d_2626_v49_))) - (len((d_2627_v50_)[default__.safeIndex((self).f15, len(d_2627_v50_))]))
            d_2628_v51_: _dafny.Array
            nw495_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_2628_v51_ = nw495_
            d_2629_v52_: _dafny.Array
            nw496_ = _dafny.Array(False, 16)
            d_2629_v52_ = nw496_
            index442_ = default__.safeIndex(893, (d_2628_v51_).length(0))
            (d_2628_v51_)[index442_] = d_2629_v52_
            index443_ = default__.safeIndex(893, (d_2628_v51_).length(0))
            (d_2628_v51_)[index443_] = d_2629_v52_
            d_2630_v53_: C4
            nw497_ = C4()
            nw497_.ctor__()
            d_2630_v53_ = nw497_
            d_2631_v54_: D14
            d_2631_v54_ = D14_DC40(d_2569_v0_)
            d_2632_v55_: D14
            d_2632_v55_ = D14_DC41(d_2631_v54_)
            d_2633_v56_: D14
            d_2633_v56_ = D14_DC41(d_2632_v55_)
            d_2634_v57_: D14
            d_2634_v57_ = D14_DC41(d_2632_v55_)
            d_2635_v58_: _dafny.MultiSet
            d_2635_v58_ = _dafny.MultiSet([len(d_2626_v49_)])
            rhs351_ = d_2630_v53_
            rhs352_ = d_2634_v57_
            rhs353_ = ((d_2635_v58_).cardinality) * ((self).f15)
            lhs301_ = globalState
            d_2630_v53_ = rhs351_
            d_2634_v57_ = rhs352_
            lhs301_.f6 = rhs353_
            arr0_ = (d_2628_v51_)[default__.safeIndex(893, (d_2628_v51_).length(0))]
            index444_ = default__.safeIndex(124, ((d_2628_v51_)[default__.safeIndex(893, (d_2628_v51_).length(0))]).length(0))
            arr0_[index444_] = d_2569_v0_
            d_2636_v59_: str
            d_2636_v59_ = _dafny.CodePoint('b')
            d_2637_v60_: _dafny.MultiSet
            d_2637_v60_ = _dafny.MultiSet([d_2636_v59_])
            arr1_ = (d_2628_v51_)[default__.safeIndex(893, (d_2628_v51_).length(0))]
            index445_ = default__.safeIndex(124, ((d_2628_v51_)[default__.safeIndex(893, (d_2628_v51_).length(0))]).length(0))
            arr1_[index445_] = ((default__.fm46(not(d_2569_v0_), True, d_2569_v0_, d_2569_v0_, globalState)) - (d_2637_v60_)).ispropersubset(d_2637_v60_)
        d_2638_v61_: _dafny.Array
        def lambda133_(d_2639_v0_):
            def lambda134_(d_2640_i5_):
                return d_2639_v0_

            return lambda134_

        init76_ = lambda133_(d_2569_v0_)
        nw498_ = _dafny.Array(None, 25)
        for i0_76_ in range(nw498_.length(0)):
            nw498_[i0_76_] = init76_(i0_76_)
        d_2638_v61_ = nw498_
        index446_ = default__.safeIndex(287, (d_2638_v61_).length(0))
        (d_2638_v61_)[index446_] = True
        d_2641_v62_: _dafny.Map
        d_2641_v62_ = _dafny.Map({(self).f15: d_2569_v0_})
        d_2642_v63_: D16
        d_2642_v63_ = D16_DC44(_dafny.MultiSet(d_2602_v29_))
        d_2643_v64_: _dafny.Seq
        d_2643_v64_ = _dafny.SeqWithoutIsStrInference([d_2642_v63_, d_2642_v63_])
        index447_ = default__.safeIndex(287, (d_2638_v61_).length(0))
        rhs354_ = d_2569_v0_
        rhs355_ = (d_2641_v62_) | (d_2641_v62_)
        rhs356_ = (_dafny.SeqWithoutIsStrInference([d_2642_v63_, d_2642_v63_])) != ((d_2643_v64_) + (_dafny.SeqWithoutIsStrInference([D16_DC44(_dafny.MultiSet([d_2569_v0_])) for d_2644_i6_ in range(default__.abs(500))])))
        lhs302_ = d_2638_v61_
        lhs303_ = default__.safeIndex(287, (d_2638_v61_).length(0))
        lhs302_[lhs303_] = rhs354_
        d_2641_v62_ = rhs355_
        r1 = rhs356_
        d_2645_v65_: _dafny.Array
        nw499_ = _dafny.Array(int(0), 10)
        d_2645_v65_ = nw499_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_2645_v65_).length(0)):
            d_2646_i7_: int = guard_loop_7_
            if (True) and (((0) <= (d_2646_i7_)) and ((d_2646_i7_) < ((d_2645_v65_).length(0)))):
                (d_2645_v65_)[(d_2646_i7_)] = (d_2646_i7_) - ((self).f15)
        d_2647_v66_: C7
        nw500_ = C7()
        nw500_.ctor__()
        d_2647_v66_ = nw500_
        d_2648_v67_: _dafny.MultiSet
        d_2648_v67_ = _dafny.MultiSet([d_2647_v66_])
        d_2649_v68_: _dafny.Map
        d_2649_v68_ = _dafny.Map({d_2569_v0_: d_2648_v67_})
        d_2650_v69_: _dafny.Seq
        d_2650_v69_ = _dafny.SeqWithoutIsStrInference([d_2647_v66_, d_2647_v66_])
        d_2651_i8_: int
        d_2651_i8_ = 0
        with _dafny.label("19"):
            while (d_2569_v0_ if (d_2638_v61_)[default__.safeIndex(287, (d_2638_v61_).length(0))] else (_dafny.MultiSet(d_2650_v69_)).ispropersubset(((d_2649_v68_)[d_2569_v0_] if (d_2569_v0_) in (d_2649_v68_) else _dafny.MultiSet([d_2647_v66_, d_2647_v66_])))):
                with _dafny.c_label("19"):
                    if (d_2651_i8_) >= (100):
                        raise _dafny.Break("19")
                    d_2651_i8_ = (d_2651_i8_) + (1)
                    d_2652_v70_: _dafny.MultiSet
                    d_2652_v70_ = _dafny.MultiSet([False, (d_2638_v61_)[default__.safeIndex(287, (d_2638_v61_).length(0))], (d_2638_v61_)[default__.safeIndex(287, (d_2638_v61_).length(0))]])
                    d_2653_v71_: _dafny.Seq
                    d_2653_v71_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_2602_v29_)) | (d_2652_v70_), d_2652_v70_, d_2652_v70_, d_2652_v70_])
                    rhs357_ = (456 if False else (self).f15)
                    rhs358_ = ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_2638_v61_)[default__.safeIndex(287, (d_2638_v61_).length(0))]])])) + (d_2653_v71_)).set(default__.safeIndex(((self).f15) + ((self).f15), len((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_2638_v61_)[default__.safeIndex(287, (d_2638_v61_).length(0))]])])) + (d_2653_v71_))), (d_2652_v70_).intersection(d_2652_v70_))
                    lhs304_ = globalState
                    lhs304_.f6 = rhs357_
                    d_2653_v71_ = rhs358_
                    index448_ = default__.safeIndex(991, (d_2645_v65_).length(0))
                    (d_2645_v65_)[index448_] = (self).f15
                    index449_ = default__.safeIndex(991, (d_2645_v65_).length(0))
                    (d_2645_v65_)[index449_] = (default__.safeDivisionInt((self).f15, (0) - ((self).f15))) - ((self).f15)
                    d_2654_v72_: _dafny.Seq
                    d_2654_v72_ = _dafny.SeqWithoutIsStrInference([d_2638_v61_])
                    r0 = len((d_2654_v72_) + (((_dafny.SeqWithoutIsStrInference([d_2638_v61_, d_2638_v61_, d_2638_v61_])).set(default__.safeIndex((self).f15, len(_dafny.SeqWithoutIsStrInference([d_2638_v61_, d_2638_v61_, d_2638_v61_]))), d_2638_v61_)).set(default__.safeIndex((self).f15, len((_dafny.SeqWithoutIsStrInference([d_2638_v61_, d_2638_v61_, d_2638_v61_])).set(default__.safeIndex((self).f15, len(_dafny.SeqWithoutIsStrInference([d_2638_v61_, d_2638_v61_, d_2638_v61_]))), d_2638_v61_))), d_2638_v61_)))
                    index450_ = default__.safeIndex(287, (d_2638_v61_).length(0))
                    (d_2638_v61_)[index450_] = (d_2569_v0_) not in (_dafny.SeqWithoutIsStrInference([True]))
                    pass
            pass
        r0 = (self).f15
        r1 = (d_2638_v61_)[default__.safeIndex(287, (d_2638_v61_).length(0))]
        return r0, r1

    def m2(self, p0, globalState):
        d_2655_v0_: _dafny.Array
        nw501_ = _dafny.Array(False, 22)
        d_2655_v0_ = nw501_
        d_2656_v1_: _dafny.Set
        d_2656_v1_ = _dafny.Set({d_2655_v0_})
        d_2656_v1_ = d_2656_v1_
        d_2657_v2_: _dafny.Array
        nw502_ = _dafny.Array(D15.default()(), 18)
        d_2657_v2_ = nw502_
        d_2658_v3_: _dafny.Map
        d_2658_v3_ = _dafny.Map({d_2657_v2_: (self).f15})
        d_2658_v3_ = (d_2658_v3_).set((d_2657_v2_ if p0 else d_2657_v2_), (self).f15)
        if p0:
            d_2659_v4_: _dafny.Map
            d_2659_v4_ = _dafny.Map({(718) * ((self).f15): not(p0)})
            d_2660_v5_: _dafny.Seq
            d_2660_v5_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2661_v6_: _dafny.Map
            d_2661_v6_ = _dafny.Map({p0: (self).f15})
            d_2662_v7_: _dafny.Map
            d_2662_v7_ = d_2661_v6_
            d_2663_v8_: _dafny.Set
            d_2663_v8_ = _dafny.Set({d_2662_v7_})
            d_2664_v9_: _dafny.MultiSet
            d_2664_v9_ = _dafny.MultiSet([p0, p0, (d_2660_v5_)[default__.safeIndex(len(d_2663_v8_), len(d_2660_v5_))], p0, p0])
            (globalState).f1 = ((d_2659_v4_)[((self).f15) * ((self).f15)] if (((self).f15) * ((self).f15)) in (d_2659_v4_) else (_dafny.MultiSet([p0, p0])).issubset(d_2664_v9_))
            d_2665_v10_: _dafny.Array
            nw503_ = _dafny.Array(_dafny.Map({}), 1)
            d_2665_v10_ = nw503_
            d_2666_v11_: D20
            d_2666_v11_ = D20_DC49(d_2665_v10_)
            d_2667_v12_: _dafny.Set
            d_2667_v12_ = _dafny.Set({d_2666_v11_})
            (globalState).f6 = (0) - ((((d_2661_v6_)[p0] if (p0) in (d_2661_v6_) else (self).f15)) + (len((d_2667_v12_) | (d_2667_v12_))))
            if (d_2664_v9_).ispropersubset(d_2664_v9_):
                d_2668_v13_: str
                d_2668_v13_ = _dafny.CodePoint('v')
                d_2668_v13_ = d_2668_v13_
                d_2669_v14_: _dafny.Seq
                d_2669_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cevl"))
                d_2669_v14_ = default__.fm21(globalState)
                (globalState).f1 = True
                d_2670_v15_: _dafny.Seq
                d_2670_v15_ = _dafny.SeqWithoutIsStrInference([d_2655_v0_])
                d_2671_v16_: _dafny.Map
                d_2671_v16_ = _dafny.Map({default__.fm1(p0, globalState): _dafny.SeqWithoutIsStrInference([d_2655_v0_])})
                d_2672_v17_: C7
                nw504_ = C7()
                nw504_.ctor__()
                d_2672_v17_ = nw504_
                d_2673_v18_: _dafny.Seq
                d_2673_v18_ = _dafny.SeqWithoutIsStrInference([d_2672_v17_])
                d_2674_v19_: _dafny.MultiSet
                d_2674_v19_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_2668_v13_ for d_2675_i0_ in range(default__.abs(191))])), (self).f15])
                d_2676_v20_: _dafny.MultiSet
                d_2676_v20_ = _dafny.MultiSet([d_2672_v17_, (d_2673_v18_)[default__.safeIndex((d_2674_v19_).cardinality, len(d_2673_v18_))]])
                d_2677_v21_: _dafny.Array
                nw505_ = _dafny.Array(None, 27)
                nw505_[int(0)] = _dafny.SeqWithoutIsStrInference([d_2655_v0_])
                nw505_[int(1)] = d_2670_v15_
                nw505_[int(2)] = d_2670_v15_
                nw505_[int(3)] = d_2670_v15_
                nw505_[int(4)] = d_2670_v15_
                nw505_[int(5)] = d_2670_v15_
                nw505_[int(6)] = d_2670_v15_
                nw505_[int(7)] = _dafny.SeqWithoutIsStrInference([d_2655_v0_, d_2655_v0_])
                nw505_[int(8)] = ((d_2671_v16_)[((d_2676_v20_)[d_2672_v17_] if (d_2672_v17_) in (d_2676_v20_) else (self).f15)] if (((d_2676_v20_)[d_2672_v17_] if (d_2672_v17_) in (d_2676_v20_) else (self).f15)) in (d_2671_v16_) else d_2670_v15_)
                nw505_[int(9)] = (d_2670_v15_) + (d_2670_v15_)
                nw505_[int(10)] = (_dafny.SeqWithoutIsStrInference([d_2655_v0_])) + (d_2670_v15_)
                nw505_[int(11)] = d_2670_v15_
                nw505_[int(12)] = d_2670_v15_
                nw505_[int(13)] = d_2670_v15_
                nw505_[int(14)] = _dafny.SeqWithoutIsStrInference([d_2655_v0_, d_2655_v0_, d_2655_v0_, d_2655_v0_])
                nw505_[int(15)] = (d_2670_v15_) + (d_2670_v15_)
                nw505_[int(16)] = (d_2670_v15_) + ((d_2670_v15_).set(default__.safeIndex((self).f15, len(d_2670_v15_)), d_2655_v0_))
                nw505_[int(17)] = d_2670_v15_
                nw505_[int(18)] = d_2670_v15_
                nw505_[int(19)] = d_2670_v15_
                nw505_[int(20)] = d_2670_v15_
                nw505_[int(21)] = d_2670_v15_
                nw505_[int(22)] = (d_2670_v15_ if p0 else d_2670_v15_)
                nw505_[int(23)] = (d_2670_v15_) + (d_2670_v15_)
                nw505_[int(24)] = ((_dafny.SeqWithoutIsStrInference([d_2655_v0_, d_2655_v0_, d_2655_v0_, d_2655_v0_, d_2655_v0_])) + (d_2670_v15_)).set(default__.safeIndex(380, len((_dafny.SeqWithoutIsStrInference([d_2655_v0_, d_2655_v0_, d_2655_v0_, d_2655_v0_, d_2655_v0_])) + (d_2670_v15_))), d_2655_v0_)
                nw505_[int(25)] = d_2670_v15_
                nw505_[int(26)] = ((d_2670_v15_).set(default__.safeIndex((self).f15, len(d_2670_v15_)), d_2655_v0_)) + (d_2670_v15_)
                d_2677_v21_ = nw505_
                index451_ = default__.safeIndex(388, (d_2677_v21_).length(0))
                (d_2677_v21_)[index451_] = (d_2670_v15_ if p0 else d_2670_v15_)
                d_2678_v22_: _dafny.Array
                nw506_ = _dafny.Array(_dafny.Map({}), 26)
                d_2678_v22_ = nw506_
                d_2679_v23_: _dafny.Map
                d_2679_v23_ = _dafny.Map({len(d_2659_v4_): d_2668_v13_})
                index452_ = default__.safeIndex(446, (d_2678_v22_).length(0))
                def iife266_():
                    coll88_ = _dafny.Map()
                    compr_88_: int
                    for compr_88_ in _dafny.IntegerRange(163, 781):
                        d_2680_v24_: int = compr_88_
                        if ((163) <= (d_2680_v24_)) and ((d_2680_v24_) < (781)):
                            coll88_[default__.safeModuloInt(d_2680_v24_, (d_2674_v19_).cardinality)] = d_2668_v13_
                    return _dafny.Map(coll88_)
                (d_2678_v22_)[index452_] = (d_2679_v23_) | (iife266_()
                )
                d_2681_v25_: _dafny.Array
                nw507_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_2681_v25_ = nw507_
                d_2682_v26_: _dafny.Map
                d_2682_v26_ = _dafny.Map({d_2668_v13_: (self).f15})
                d_2683_v27_: T1
                nw508_ = C2()
                nw508_.ctor__((self).f15, (self).f15, default__.fm47((self).f15, default__.fm1(p0, globalState), (self).f15, len(d_2682_v26_), globalState))
                d_2683_v27_ = nw508_
                d_2684_v28_: _dafny.Seq
                d_2684_v28_ = _dafny.SeqWithoutIsStrInference([d_2683_v27_])
                d_2685_v29_: _dafny.Array
                nw509_ = _dafny.Array(D3.default()(), 17)
                d_2685_v29_ = nw509_
                d_2686_v30_: _dafny.Set
                d_2686_v30_ = _dafny.Set({d_2685_v29_})
                d_2687_v31_: _dafny.Map
                d_2687_v31_ = _dafny.Map({d_2660_v5_: d_2669_v14_})
                d_2688_v32_: _dafny.Map
                d_2688_v32_ = _dafny.Map({p0: p0})
                d_2689_v33_: _dafny.MultiSet
                d_2689_v33_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f15])])
                d_2690_v34_: _dafny.Array
                nw510_ = _dafny.Array(None, 11)
                nw510_[int(0)] = (self).f15
                nw510_[int(1)] = len(d_2688_v32_)
                nw510_[int(2)] = (d_2664_v9_).cardinality
                nw510_[int(3)] = (self).f15
                nw510_[int(4)] = (self).f15
                nw510_[int(5)] = (self).f15
                nw510_[int(6)] = (self).f15
                nw510_[int(7)] = (self).f15
                nw510_[int(8)] = (self).f15
                nw510_[int(9)] = (self).f15
                nw510_[int(10)] = (d_2689_v33_).cardinality
                d_2690_v34_ = nw510_
                d_2691_v35_: D9
                d_2691_v35_ = D9_DC23(d_2686_v30_, ((d_2687_v31_)[d_2660_v5_] if (d_2660_v5_) in (d_2687_v31_) else d_2669_v14_), p0, d_2690_v34_, _dafny.CodePoint('l'))
                d_2692_v36_: _dafny.Array
                nw511_ = _dafny.Array(None, 25)
                nw511_[int(0)] = _dafny.CodePoint('n')
                nw511_[int(1)] = d_2668_v13_
                nw511_[int(2)] = d_2668_v13_
                nw511_[int(3)] = d_2668_v13_
                nw511_[int(4)] = d_2668_v13_
                nw511_[int(5)] = (d_2669_v14_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2668_v13_ for d_2693_i1_ in range(default__.abs(551))])), len(d_2669_v14_))]
                nw511_[int(6)] = d_2668_v13_
                nw511_[int(7)] = d_2668_v13_
                nw511_[int(8)] = d_2668_v13_
                nw511_[int(9)] = _dafny.CodePoint('p')
                nw511_[int(10)] = d_2668_v13_
                nw511_[int(11)] = d_2668_v13_
                nw511_[int(12)] = d_2668_v13_
                nw511_[int(13)] = d_2668_v13_
                nw511_[int(14)] = _dafny.CodePoint('b')
                nw511_[int(15)] = d_2668_v13_
                nw511_[int(16)] = d_2668_v13_
                nw511_[int(17)] = d_2668_v13_
                nw511_[int(18)] = d_2668_v13_
                nw511_[int(19)] = d_2668_v13_
                nw511_[int(20)] = default__.fm29(False, len((d_2684_v28_).set(default__.safeIndex((self).f15, len(d_2684_v28_)), d_2683_v27_)), ((d_2682_v26_)[d_2668_v13_] if (d_2668_v13_) in (d_2682_v26_) else (self).f15), p0, globalState)
                nw511_[int(21)] = d_2668_v13_
                nw511_[int(22)] = d_2668_v13_
                nw511_[int(23)] = d_2668_v13_
                nw511_[int(24)] = (d_2691_v35_).cf41
                d_2692_v36_ = nw511_
                index453_ = default__.safeIndex(195, (d_2681_v25_).length(0))
                (d_2681_v25_)[index453_] = d_2692_v36_
                index454_ = default__.safeIndex(388, (d_2677_v21_).length(0))
                index455_ = default__.safeIndex(446, (d_2678_v22_).length(0))
                index456_ = default__.safeIndex(195, (d_2681_v25_).length(0))
                rhs359_ = d_2670_v15_
                rhs360_ = d_2662_v7_
                rhs361_ = d_2668_v13_
                rhs362_ = (d_2679_v23_) | ((d_2679_v23_).set((self).f15, d_2668_v13_))
                rhs363_ = d_2692_v36_
                lhs305_ = d_2677_v21_
                lhs306_ = default__.safeIndex(388, (d_2677_v21_).length(0))
                lhs307_ = d_2678_v22_
                lhs308_ = default__.safeIndex(446, (d_2678_v22_).length(0))
                lhs309_ = d_2681_v25_
                lhs310_ = default__.safeIndex(195, (d_2681_v25_).length(0))
                lhs305_[lhs306_] = rhs359_
                d_2662_v7_ = rhs360_
                d_2668_v13_ = rhs361_
                lhs307_[lhs308_] = rhs362_
                lhs309_[lhs310_] = rhs363_
                (globalState).f1 = (d_2660_v5_)[default__.safeIndex((self).f15, len(d_2660_v5_))]
            elif True:
                (globalState).f1 = p0
                d_2694_v37_: C5
                nw512_ = C5()
                nw512_.ctor__((0) - ((self).f15))
                d_2694_v37_ = nw512_
                d_2695_v38_: _dafny.Seq
                d_2695_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
                d_2696_v39_: _dafny.Map
                d_2696_v39_ = _dafny.Map({D23_DC53(d_2694_v37_): d_2695_v38_})
                d_2696_v39_ = d_2696_v39_
                (globalState).f14 = d_2655_v0_
                d_2697_v40_: str
                d_2697_v40_ = _dafny.CodePoint('u')
                d_2697_v40_ = d_2697_v40_
                (globalState).f6 = -41
            d_2698_v41_: _dafny.Map
            d_2698_v41_ = _dafny.Map({d_2659_v4_: len(_dafny.Map({(self).f15: (self).f15}))})
            d_2699_v42_: T1
            nw513_ = C3()
            nw513_.ctor__(p0, (self).f15, d_2698_v41_)
            d_2699_v42_ = nw513_
            d_2699_v42_ = ((d_2699_v42_ if p0 else d_2699_v42_) if False else d_2699_v42_)
            d_2700_v43_: _dafny.Array
            def lambda135_(d_2701_i2_):
                return (d_2701_i2_) + ((self).f15)

            init77_ = lambda135_
            nw514_ = _dafny.Array(None, 10)
            for i0_77_ in range(nw514_.length(0)):
                nw514_[i0_77_] = init77_(i0_77_)
            d_2700_v43_ = nw514_
            (globalState).f9 = d_2700_v43_
        elif True:
            (globalState).f6 = (self).f15
            d_2702_v44_: str
            d_2702_v44_ = _dafny.CodePoint('s')
            d_2703_v45_: _dafny.MultiSet
            d_2703_v45_ = _dafny.MultiSet([(self).f15, (self).f15, (self).f15])
            pat_let_tv63_ = d_2703_v45_
            pat_let_tv64_ = d_2703_v45_
            pat_let_tv65_ = d_2702_v44_
            def iife267_(_pat_let89_0):
                def iife268_(d_2704_dt__update__tmp_h0_):
                    def iife269_(_pat_let90_0):
                        def iife270_(d_2705_dt__update_hcf16_h0_):
                            def iife271_(_pat_let91_0):
                                def iife272_(d_2706_dt__update_hcf17_h0_):
                                    return D5_DC12((d_2704_dt__update__tmp_h0_).cf15, d_2705_dt__update_hcf16_h0_, d_2706_dt__update_hcf17_h0_)
                                return iife272_(_pat_let91_0)
                            return iife271_(pat_let_tv65_)
                        return iife270_(_pat_let90_0)
                    return iife269_((pat_let_tv63_).issubset(pat_let_tv64_))
                return iife268_(_pat_let89_0)
            source25_ = iife267_(D5_DC12((self).f15, p0, d_2702_v44_))
            if source25_.is_DC11:
                d_2707___mcc_h0_ = source25_.cf14
                d_2708_cf14_ = d_2707___mcc_h0_
                d_2709_v46_: _dafny.Map
                d_2709_v46_ = _dafny.Map({not (p0) or (True): p0})
                d_2709_v46_ = (d_2709_v46_).set(p0, p0)
                d_2710_v47_: _dafny.Set
                d_2710_v47_ = _dafny.Set({d_2708_cf14_})
                d_2711_v48_: _dafny.Map
                d_2711_v48_ = _dafny.Map({default__.fm1(p0, globalState): d_2710_v47_})
                d_2712_v49_: _dafny.Map
                d_2712_v49_ = d_2711_v48_
                d_2712_v49_ = d_2711_v48_
                d_2708_cf14_ = (((self).f15) * (d_2708_cf14_) if not (p0) or ((self).fm4(d_2708_cf14_, globalState)) else (self).f15)
                d_2713_v50_: _dafny.Array
                nw515_ = _dafny.Array(int(0), 24)
                d_2713_v50_ = nw515_
                d_2714_v51_: _dafny.Seq
                d_2714_v51_ = _dafny.SeqWithoutIsStrInference([p0])
                index457_ = default__.safeIndex(314, (d_2713_v50_).length(0))
                (d_2713_v50_)[index457_] = default__.safeModuloInt(len(d_2714_v51_), d_2708_cf14_)
                d_2715_v52_: _dafny.Set
                d_2715_v52_ = _dafny.Set({default__.fm2(globalState), p0})
                d_2716_v53_: _dafny.Map
                d_2716_v53_ = _dafny.Map({default__.fm2(globalState): d_2715_v52_})
                d_2717_v54_: D0
                d_2717_v54_ = D0_DC0(p0)
                d_2718_v55_: _dafny.Seq
                d_2718_v55_ = _dafny.SeqWithoutIsStrInference([d_2710_v47_, d_2710_v47_])
                d_2719_v56_: _dafny.Map
                d_2719_v56_ = _dafny.Map({(d_2718_v55_)[default__.safeIndex(d_2708_cf14_, len(d_2718_v55_))]: (0) - (len(d_2714_v51_))})
                d_2720_v57_: _dafny.MultiSet
                d_2720_v57_ = _dafny.MultiSet([d_2715_v52_])
                index458_ = default__.safeIndex(314, (d_2713_v50_).length(0))
                rhs364_ = (((d_2719_v56_)[d_2710_v47_] if (d_2710_v47_) in (d_2719_v56_) else d_2708_cf14_)) - (((d_2720_v57_)[d_2715_v52_] if (d_2715_v52_) in (d_2720_v57_) else 318))
                rhs365_ = p0
                rhs366_ = d_2713_v50_
                rhs367_ = d_2716_v53_
                rhs368_ = d_2717_v54_
                lhs311_ = d_2713_v50_
                lhs312_ = default__.safeIndex(314, (d_2713_v50_).length(0))
                lhs313_ = globalState
                lhs314_ = globalState
                lhs311_[lhs312_] = rhs364_
                lhs313_.f1 = rhs365_
                lhs314_.f9 = rhs366_
                d_2716_v53_ = rhs367_
                d_2717_v54_ = rhs368_
            elif source25_.is_DC12:
                d_2721___mcc_h1_ = source25_.cf15
                d_2722___mcc_h2_ = source25_.cf16
                d_2723___mcc_h3_ = source25_.cf17
                d_2724_cf17_ = d_2723___mcc_h3_
                d_2725_cf16_ = d_2722___mcc_h2_
                d_2726_cf15_ = d_2721___mcc_h1_
                d_2727_v58_: _dafny.Seq
                d_2727_v58_ = _dafny.SeqWithoutIsStrInference([d_2724_cf17_, d_2702_v44_, d_2724_cf17_])
                d_2727_v58_ = default__.fm21(globalState)
                d_2728_v59_: _dafny.Array
                nw516_ = _dafny.Array(D13.default()(), 3)
                d_2728_v59_ = nw516_
                d_2729_v60_: _dafny.Map
                d_2729_v60_ = _dafny.Map({p0: d_2728_v59_})
                d_2726_cf15_ = len(d_2729_v60_)
                d_2730_v61_: _dafny.MultiSet
                d_2730_v61_ = _dafny.MultiSet([True, d_2725_cf16_, d_2725_cf16_])
                d_2731_v62_: _dafny.Seq
                d_2731_v62_ = _dafny.SeqWithoutIsStrInference([d_2725_cf16_])
                d_2732_v63_: _dafny.Map
                d_2732_v63_ = _dafny.Map({d_2730_v61_: d_2731_v62_})
                d_2732_v63_ = d_2732_v63_
                d_2733_v64_: _dafny.Seq
                d_2733_v64_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: d_2725_cf16_}))])
                d_2734_v65_: _dafny.Map
                d_2734_v65_ = _dafny.Map({d_2733_v64_: len(_dafny.Map({d_2726_cf15_: d_2725_cf16_}))})
                d_2735_v66_: _dafny.Map
                d_2735_v66_ = _dafny.Map({-974: p0})
                d_2736_v67_: _dafny.Seq
                d_2736_v67_ = _dafny.SeqWithoutIsStrInference([d_2735_v66_])
                d_2737_v68_: _dafny.Map
                d_2737_v68_ = _dafny.Map({(d_2736_v67_)[default__.safeIndex(d_2726_cf15_, len(d_2736_v67_))]: 425})
                d_2738_v69_: T1
                nw517_ = C3()
                nw517_.ctor__(p0, len(d_2734_v65_), d_2737_v68_)
                d_2738_v69_ = nw517_
                d_2739_v70_: _dafny.Map
                d_2739_v70_ = _dafny.Map({(self).f15: d_2738_v69_})
                (globalState).f1 = ((default__.fm48(_dafny.Map({d_2702_v44_: d_2702_v44_}), globalState)).cf7) == (len(d_2739_v70_))
            elif source25_.is_DC10:
                d_2740___mcc_h4_ = source25_.cf13
                d_2741_cf13_ = d_2740___mcc_h4_
                d_2742_v71_: _dafny.Array
                nw518_ = _dafny.Array(D0.default()(), 13)
                d_2742_v71_ = nw518_
                d_2743_v72_: _dafny.Array
                d_2743_v72_ = d_2742_v71_
                d_2744_v73_: _dafny.Map
                d_2744_v73_ = _dafny.Map({default__.fm0(p0, p0, (self).f15, globalState): (self).f15})
                d_2745_v74_: _dafny.Map
                d_2745_v74_ = _dafny.Map({d_2743_v72_: d_2744_v73_})
                d_2746_v75_: _dafny.Map
                d_2746_v75_ = ((d_2745_v74_)[d_2743_v72_] if (d_2743_v72_) in (d_2745_v74_) else d_2744_v73_)
                d_2746_v75_ = d_2746_v75_
                d_2747_v76_: _dafny.Array
                nw519_ = _dafny.Array(int(0), 3)
                d_2747_v76_ = nw519_
                index459_ = default__.safeIndex(833, (d_2747_v76_).length(0))
                (d_2747_v76_)[index459_] = (self).f15
                index460_ = default__.safeIndex(833, (d_2747_v76_).length(0))
                (d_2747_v76_)[index460_] = (0) - ((self).f15)
                (globalState).f1 = p0
                d_2748_v77_: _dafny.Seq
                d_2748_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "do"))
                index461_ = default__.safeIndex(833, (d_2747_v76_).length(0))
                (d_2747_v76_)[index461_] = len((d_2748_v77_ if p0 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecflprnxd"))))
            elif True:
                d_2749___mcc_h5_ = source25_.cf18
                d_2750_cf18_ = d_2749___mcc_h5_
                d_2751_v78_: _dafny.Set
                d_2751_v78_ = _dafny.Set({(self).f15, (self).f15})
                d_2752_v79_: D0
                d_2752_v79_ = D0_DC0(p0)
                d_2753_v80_: D0
                d_2753_v80_ = D0_DC2(d_2752_v79_)
                d_2754_v81_: _dafny.Seq
                d_2754_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vydknp"))
                d_2755_v82_: D16
                d_2755_v82_ = D16_DC45((d_2703_v45_).cardinality, d_2702_v44_, len(_dafny.SeqWithoutIsStrInference([p0])), True)
                pat_let_tv66_ = d_2752_v79_
                pat_let_tv67_ = d_2752_v79_
                pat_let_tv68_ = d_2752_v79_
                pat_let_tv69_ = d_2752_v79_
                pat_let_tv70_ = p0
                d_2756_v83_: _dafny.Array
                nw520_ = _dafny.Array(None, 27)
                nw520_[int(0)] = d_2753_v80_
                nw520_[int(1)] = d_2753_v80_
                nw520_[int(2)] = d_2753_v80_
                nw520_[int(3)] = d_2753_v80_
                def iife274_(_pat_let93_0):
                    def iife275_(d_2757_dt__update__tmp_h1_):
                        def iife276_(_pat_let94_0):
                            def iife277_(d_2758_dt__update_hcf2_h0_):
                                return D0_DC2(d_2758_dt__update_hcf2_h0_)
                            return iife277_(_pat_let94_0)
                        return iife276_(pat_let_tv66_)
                    return iife275_(_pat_let93_0)
                def iife273_(_pat_let92_0):
                    def iife278_(d_2759_dt__update__tmp_h2_):
                        def iife279_(_pat_let95_0):
                            def iife280_(d_2760_dt__update_hcf2_h1_):
                                return D0_DC2(d_2760_dt__update_hcf2_h1_)
                            return iife280_(_pat_let95_0)
                        return iife279_(pat_let_tv67_)
                    return iife278_(_pat_let92_0)
                nw520_[int(4)] = iife273_(iife274_(d_2753_v80_))
                nw520_[int(5)] = d_2753_v80_
                nw520_[int(6)] = d_2753_v80_
                nw520_[int(7)] = d_2753_v80_
                nw520_[int(8)] = d_2753_v80_
                nw520_[int(9)] = d_2753_v80_
                nw520_[int(10)] = d_2753_v80_
                def iife281_(_pat_let96_0):
                    def iife282_(d_2761_dt__update__tmp_h3_):
                        def iife283_(_pat_let97_0):
                            def iife284_(d_2762_dt__update_hcf2_h2_):
                                return D0_DC2(d_2762_dt__update_hcf2_h2_)
                            return iife284_(_pat_let97_0)
                        return iife283_(pat_let_tv68_)
                    return iife282_(_pat_let96_0)
                nw520_[int(11)] = iife281_(d_2753_v80_)
                nw520_[int(12)] = d_2753_v80_
                nw520_[int(13)] = d_2753_v80_
                def iife285_(_pat_let98_0):
                    def iife286_(d_2763_dt__update__tmp_h4_):
                        def iife287_(_pat_let99_0):
                            def iife288_(d_2764_dt__update_hcf2_h3_):
                                return D0_DC2(d_2764_dt__update_hcf2_h3_)
                            return iife288_(_pat_let99_0)
                        return iife287_(pat_let_tv69_)
                    return iife286_(_pat_let98_0)
                nw520_[int(14)] = iife285_(D0_DC2(d_2752_v79_))
                nw520_[int(15)] = d_2753_v80_
                nw520_[int(16)] = d_2753_v80_
                nw520_[int(17)] = D0_DC2(d_2752_v79_)
                nw520_[int(18)] = d_2753_v80_
                nw520_[int(19)] = D0_DC2(d_2752_v79_)
                nw520_[int(20)] = d_2753_v80_
                def iife289_(_pat_let100_0):
                    def iife290_(d_2765_dt__update__tmp_h5_):
                        def iife291_(_pat_let101_0):
                            def iife292_(d_2766_dt__update_hcf2_h4_):
                                return D0_DC2(d_2766_dt__update_hcf2_h4_)
                            return iife292_(_pat_let101_0)
                        return iife291_(D0_DC0(pat_let_tv70_))
                    return iife290_(_pat_let100_0)
                nw520_[int(21)] = iife289_(d_2753_v80_)
                nw520_[int(22)] = d_2753_v80_
                nw520_[int(23)] = d_2753_v80_
                nw520_[int(24)] = d_2753_v80_
                nw520_[int(25)] = (self).fm3(p0, len(d_2754_v81_), p0, (d_2755_v82_).cf81, globalState)
                nw520_[int(26)] = d_2753_v80_
                d_2756_v83_ = nw520_
                d_2767_v84_: _dafny.Seq
                d_2767_v84_ = _dafny.SeqWithoutIsStrInference([d_2756_v83_, d_2756_v83_])
                d_2768_v85_: _dafny.Array
                def lambda136_(d_2769_i3_):
                    return (d_2769_i3_) * ((self).f15)

                init78_ = lambda136_
                nw521_ = _dafny.Array(None, 24)
                for i0_78_ in range(nw521_.length(0)):
                    nw521_[i0_78_] = init78_(i0_78_)
                d_2768_v85_ = nw521_
                default__.m0((d_2751_v78_ if not(False) else d_2751_v78_), (d_2767_v84_)[default__.safeIndex((self).f15, len(d_2767_v84_))], d_2768_v85_, d_2754_v81_, globalState)
                d_2770_v86_: D3
                d_2770_v86_ = D3_DC6(d_2768_v85_)
                d_2771_v87_: _dafny.Array
                nw522_ = _dafny.Array(None, 20)
                nw522_[int(0)] = D3_DC6(d_2768_v85_)
                nw522_[int(1)] = d_2770_v86_
                nw522_[int(2)] = d_2770_v86_
                nw522_[int(3)] = d_2770_v86_
                nw522_[int(4)] = d_2770_v86_
                nw522_[int(5)] = d_2770_v86_
                nw522_[int(6)] = d_2770_v86_
                nw522_[int(7)] = d_2770_v86_
                nw522_[int(8)] = d_2770_v86_
                nw522_[int(9)] = d_2770_v86_
                nw522_[int(10)] = D3_DC6(d_2768_v85_)
                nw522_[int(11)] = d_2770_v86_
                nw522_[int(12)] = d_2770_v86_
                nw522_[int(13)] = d_2770_v86_
                nw522_[int(14)] = D3_DC6(d_2768_v85_)
                nw522_[int(15)] = D3_DC6(d_2768_v85_)
                nw522_[int(16)] = d_2770_v86_
                nw522_[int(17)] = d_2770_v86_
                nw522_[int(18)] = d_2770_v86_
                nw522_[int(19)] = D3_DC6(d_2768_v85_)
                d_2771_v87_ = nw522_
                d_2772_v88_: _dafny.Set
                d_2772_v88_ = _dafny.Set({d_2771_v87_})
                (globalState).f1 = (d_2754_v81_) <= ((D9_DC23(d_2772_v88_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "luiufanv")), p0, d_2768_v85_, d_2702_v44_)).cf38)
                index462_ = default__.safeIndex(660, (d_2768_v85_).length(0))
                (d_2768_v85_)[index462_] = (self).f15
                index463_ = default__.safeIndex(660, (d_2768_v85_).length(0))
                (d_2768_v85_)[index463_] = (self).f15
                d_2773_v89_: C6
                nw523_ = C6()
                nw523_.ctor__()
                d_2773_v89_ = nw523_
            d_2774_v90_: _dafny.Map
            d_2774_v90_ = _dafny.Map({145: (self).f15})
            d_2775_v91_: _dafny.Set
            d_2775_v91_ = _dafny.Set({p0})
            d_2776_v92_: _dafny.Map
            d_2776_v92_ = _dafny.Map({(d_2774_v90_) | (d_2774_v90_): d_2775_v91_})
            d_2777_v93_: _dafny.Seq
            d_2777_v93_ = _dafny.SeqWithoutIsStrInference([((self).f15) * ((self).f15), ((self).f15) + (576), (self).f15, (self).f15, (self).f15])
            d_2778_v94_: _dafny.Seq
            d_2778_v94_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2779_v95_: D11
            d_2779_v95_ = D11_DC31(d_2778_v94_, (self).f15, default__.fm1(p0, globalState), len(d_2774_v90_), (self).f15)
            d_2780_v96_: _dafny.Map
            d_2780_v96_ = _dafny.Map({(D3_DC8(d_2777_v93_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsnffddq")), p0, _dafny.CodePoint('o'))).cf10: d_2702_v44_})
            d_2781_v97_: _dafny.Array
            nw524_ = _dafny.Array(None, 9)
            nw524_[int(0)] = (self).f15
            nw524_[int(1)] = (self).f15
            nw524_[int(2)] = (self).f15
            nw524_[int(3)] = (d_2779_v95_).cf56
            nw524_[int(4)] = (self).f15
            nw524_[int(5)] = (self).f15
            nw524_[int(6)] = (self).f15
            nw524_[int(7)] = (self).f15
            nw524_[int(8)] = len(_dafny.Set({len(d_2780_v96_), (self).f15}))
            d_2781_v97_ = nw524_
            d_2782_v98_: _dafny.Set
            d_2782_v98_ = _dafny.Set({d_2781_v97_})
            d_2783_v99_: _dafny.Map
            d_2783_v99_ = _dafny.Map({d_2782_v98_: d_2777_v93_})
            rhs369_ = _dafny.Map({_dafny.Map({(self).f15: (0) - ((self).f15)}): default__.fm31((self).f15, default__.fm2(globalState), globalState)})
            rhs370_ = ((d_2783_v99_)[d_2782_v98_] if (d_2782_v98_) in (d_2783_v99_) else d_2777_v93_)
            rhs371_ = d_2655_v0_
            lhs315_ = globalState
            d_2776_v92_ = rhs369_
            d_2777_v93_ = rhs370_
            lhs315_.f14 = rhs371_
            index464_ = default__.safeIndex(467, (d_2655_v0_).length(0))
            (d_2655_v0_)[index464_] = not(p0)
            index465_ = default__.safeIndex(467, (d_2655_v0_).length(0))
            (d_2655_v0_)[index465_] = default__.fm2(globalState)
            d_2784_v100_: C7
            nw525_ = C7()
            nw525_.ctor__()
            d_2784_v100_ = nw525_
        (globalState).f6 = 309
        d_2785_i4_: int
        d_2785_i4_ = 0
        with _dafny.label("20"):
            while p0:
                with _dafny.c_label("20"):
                    if (d_2785_i4_) >= (100):
                        raise _dafny.Break("20")
                    d_2785_i4_ = (d_2785_i4_) + (1)
                    d_2786_v101_: _dafny.Array
                    nw526_ = _dafny.Array(_dafny.Map({}), 9)
                    d_2786_v101_ = nw526_
                    index466_ = default__.safeIndex(636, (d_2786_v101_).length(0))
                    (d_2786_v101_)[index466_] = default__.fm49(not(p0), globalState)
                    d_2787_v102_: _dafny.Map
                    d_2787_v102_ = _dafny.Map({p0: (self).f15})
                    d_2788_v103_: _dafny.Map
                    d_2788_v103_ = d_2787_v102_
                    index467_ = default__.safeIndex(636, (d_2786_v101_).length(0))
                    (d_2786_v101_)[index467_] = d_2787_v102_
                    d_2789_v104_: str
                    d_2789_v104_ = _dafny.CodePoint('a')
                    d_2790_v105_: _dafny.MultiSet
                    d_2790_v105_ = _dafny.MultiSet([d_2789_v104_])
                    d_2791_v106_: _dafny.Map
                    d_2791_v106_ = _dafny.Map({(d_2790_v105_).cardinality: (0) - ((self).f15)})
                    d_2792_v107_: _dafny.Seq
                    d_2792_v107_ = _dafny.SeqWithoutIsStrInference([False])
                    d_2793_v108_: _dafny.Map
                    d_2793_v108_ = _dafny.Map({d_2655_v0_: d_2791_v106_})
                    d_2794_v109_: _dafny.Map
                    d_2794_v109_ = _dafny.Map({p0: _dafny.Map({(self).f15: (self).f15})})
                    d_2795_v111_: _dafny.Array
                    nw527_ = _dafny.Array(None, 14)
                    nw527_[int(0)] = d_2791_v106_
                    nw527_[int(1)] = _dafny.Map({863: (self).f15})
                    nw527_[int(2)] = _dafny.Map({len(d_2792_v107_): (self).f15})
                    nw527_[int(3)] = (d_2791_v106_) | (d_2791_v106_)
                    nw527_[int(4)] = (d_2791_v106_).set((self).f15, (self).f15)
                    nw527_[int(5)] = d_2791_v106_
                    nw527_[int(6)] = d_2791_v106_
                    nw527_[int(7)] = d_2791_v106_
                    nw527_[int(8)] = d_2791_v106_
                    nw527_[int(9)] = (d_2791_v106_) | (_dafny.Map({(self).f15: 313}))
                    nw527_[int(10)] = (((d_2793_v108_)[d_2655_v0_] if (d_2655_v0_) in (d_2793_v108_) else _dafny.Map({(self).f15: (self).f15}))) | (((d_2794_v109_)[default__.fm2(globalState)] if (default__.fm2(globalState)) in (d_2794_v109_) else d_2791_v106_))
                    def iife293_():
                        coll89_ = _dafny.Map()
                        compr_89_: int
                        for compr_89_ in _dafny.IntegerRange(16, 622):
                            d_2796_v110_: int = compr_89_
                            if ((16) <= (d_2796_v110_)) and ((d_2796_v110_) < (622)):
                                coll89_[(d_2796_v110_) - ((self).f15)] = (self).f15
                        return _dafny.Map(coll89_)
                    nw527_[int(11)] = iife293_()
                    
                    nw527_[int(12)] = ((d_2791_v106_).set((self).f15, (self).f15)) | (d_2791_v106_)
                    nw527_[int(13)] = d_2791_v106_
                    d_2795_v111_ = nw527_
                    index468_ = default__.safeIndex(234, (d_2795_v111_).length(0))
                    (d_2795_v111_)[index468_] = d_2791_v106_
                    index469_ = default__.safeIndex(234, (d_2795_v111_).length(0))
                    (d_2795_v111_)[index469_] = d_2791_v106_
                    (globalState).f6 = (self).f15
                    (globalState).f1 = default__.fm2(globalState)
                    pass
            pass
        (globalState).f1 = (p0 if p0 else p0)

    def m3(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_2797_v0_: _dafny.Seq
        d_2797_v0_ = _dafny.SeqWithoutIsStrInference([(self).f15])
        d_2798_v1_: _dafny.Map
        d_2798_v1_ = _dafny.Map({(len(p0)) * (len(d_2797_v0_)): default__.safeModuloInt((self).f15, (self).f15)})
        d_2799_v2_: _dafny.Seq
        d_2799_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
        d_2798_v1_ = (d_2798_v1_).set((self).f15, len(d_2799_v2_))
        d_2800_v3_: bool
        d_2800_v3_ = True
        if d_2800_v3_:
            d_2800_v3_ = ((self).f15) > ((self).fm5(d_2800_v3_, globalState))
            d_2797_v0_ = d_2797_v0_
            d_2801_v4_: _dafny.Array
            def lambda137_(d_2802_i0_):
                return default__.safeModuloInt(d_2802_i0_, (self).f15)

            init79_ = lambda137_
            nw528_ = _dafny.Array(None, 7)
            for i0_79_ in range(nw528_.length(0)):
                nw528_[i0_79_] = init79_(i0_79_)
            d_2801_v4_ = nw528_
            d_2803_v5_: D3
            d_2803_v5_ = D3_DC6(d_2801_v4_)
            d_2804_v6_: _dafny.Seq
            d_2804_v6_ = _dafny.SeqWithoutIsStrInference([d_2803_v5_, d_2803_v5_, d_2803_v5_])
            d_2805_v7_: _dafny.Map
            d_2805_v7_ = _dafny.Map({d_2800_v3_: d_2804_v6_})
            d_2806_v8_: _dafny.MultiSet
            d_2806_v8_ = _dafny.MultiSet([d_2801_v4_, d_2801_v4_])
            d_2807_v9_: _dafny.Map
            d_2807_v9_ = _dafny.Map({d_2806_v8_: not(default__.fm2(globalState))})
            d_2808_v10_: C0
            nw529_ = C0()
            nw529_.ctor__(((d_2805_v7_)[d_2800_v3_] if (d_2800_v3_) in (d_2805_v7_) else d_2804_v6_), ((d_2807_v9_)[d_2806_v8_] if (d_2806_v8_) in (d_2807_v9_) else d_2800_v3_))
            d_2808_v10_ = nw529_
            d_2809_v11_: D11
            d_2809_v11_ = D11_DC32(d_2800_v3_, (self).f15, d_2800_v3_)
            d_2810_v12_: _dafny.Map
            d_2810_v12_ = _dafny.Map({False: (self).f15})
            d_2811_v13_: _dafny.Array
            nw530_ = _dafny.Array(None, 27)
            nw530_[int(0)] = (d_2808_v10_).f23
            nw530_[int(1)] = (d_2808_v10_).f23
            nw530_[int(2)] = not(((d_2808_v10_).f23) or (d_2800_v3_))
            nw530_[int(3)] = (d_2800_v3_) or (not(False))
            nw530_[int(4)] = d_2800_v3_
            nw530_[int(5)] = (d_2808_v10_).f23
            nw530_[int(6)] = not((d_2800_v3_ if (d_2808_v10_).f23 else d_2800_v3_))
            nw530_[int(7)] = not((d_2808_v10_).f23)
            nw530_[int(8)] = (d_2800_v3_) or (d_2800_v3_)
            nw530_[int(9)] = True
            nw530_[int(10)] = not((d_2808_v10_).f23)
            nw530_[int(11)] = ((_dafny.MultiSet(d_2797_v0_)).set(len(p0), default__.abs(((d_2798_v1_)[(self).f15] if ((self).f15) in (d_2798_v1_) else (self).f15)))) != (_dafny.MultiSet([(self).f15]))
            nw530_[int(12)] = d_2800_v3_
            nw530_[int(13)] = not (d_2800_v3_) or (d_2800_v3_)
            nw530_[int(14)] = (d_2808_v10_).fm17(-577, 982, (d_2808_v10_).f23, globalState)
            nw530_[int(15)] = ((self).f15) != ((self).f15)
            nw530_[int(16)] = d_2800_v3_
            nw530_[int(17)] = ((d_2809_v11_).cf60) not in (_dafny.Set({(self).f15}))
            nw530_[int(18)] = ((d_2808_v10_).f23 if (d_2808_v10_).f23 else d_2800_v3_)
            nw530_[int(19)] = ((self).f15) > (len(p0))
            nw530_[int(20)] = default__.fm2(globalState)
            nw530_[int(21)] = (self).fm4(len(_dafny.SeqWithoutIsStrInference([(self).f15, len(d_2810_v12_)])), globalState)
            nw530_[int(22)] = ((self).f15) <= (len(d_2799_v2_))
            nw530_[int(23)] = not(((self).f15) < (len(_dafny.Map({len(d_2797_v0_): (d_2808_v10_).f23}))))
            nw530_[int(24)] = (d_2800_v3_) and (d_2800_v3_)
            nw530_[int(25)] = (d_2808_v10_).f23
            nw530_[int(26)] = False
            d_2811_v13_ = nw530_
            index470_ = default__.safeIndex(782, (d_2811_v13_).length(0))
            (d_2811_v13_)[index470_] = not ((d_2808_v10_).f23) or (d_2800_v3_)
            d_2812_v14_: D20
            d_2812_v14_ = D20_DC50(d_2800_v3_, (self).f15, (self).f15, d_2800_v3_, (d_2808_v10_).f23)
            index471_ = default__.safeIndex(782, (d_2811_v13_).length(0))
            (d_2811_v13_)[index471_] = (d_2812_v14_).cf90
            d_2813_v15_: _dafny.Array
            nw531_ = _dafny.Array(None, 10)
            nw531_[int(0)] = d_2799_v2_
            nw531_[int(1)] = (d_2799_v2_) + (default__.fm21(globalState))
            nw531_[int(2)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_2814_i1_ in range(default__.abs(255))])
            nw531_[int(3)] = default__.fm21(globalState)
            nw531_[int(4)] = (d_2799_v2_ if (d_2808_v10_).f23 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "togmll")))
            nw531_[int(5)] = d_2799_v2_
            nw531_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvnvbtp"))
            nw531_[int(7)] = (d_2799_v2_) + (d_2799_v2_)
            nw531_[int(8)] = d_2799_v2_
            nw531_[int(9)] = d_2799_v2_
            d_2813_v15_ = nw531_
            d_2815_v16_: str
            d_2815_v16_ = _dafny.CodePoint('k')
            d_2816_v17_: _dafny.Map
            d_2816_v17_ = _dafny.Map({(self).f15: True})
            d_2817_v18_: _dafny.Set
            d_2817_v18_ = _dafny.Set({d_2800_v3_, False})
            d_2818_v19_: _dafny.Map
            d_2818_v19_ = _dafny.Map({(self).f15: d_2811_v13_})
            d_2819_v20_: _dafny.Set
            d_2819_v20_ = _dafny.Set({((d_2818_v19_)[(self).f15] if ((self).f15) in (d_2818_v19_) else d_2811_v13_), d_2811_v13_, d_2811_v13_, d_2811_v13_})
            d_2820_v21_: _dafny.Map
            d_2820_v21_ = _dafny.Map({(d_2808_v10_).f23: (d_2808_v10_).f23})
            d_2821_v22_: D14
            d_2821_v22_ = D14_DC40((d_2811_v13_)[default__.safeIndex(782, (d_2811_v13_).length(0))])
            d_2822_v23_: _dafny.MultiSet
            d_2822_v23_ = _dafny.MultiSet([(d_2808_v10_).f23, (d_2811_v13_)[default__.safeIndex(782, (d_2811_v13_).length(0))], d_2800_v3_, ((d_2820_v21_)[d_2800_v3_] if (d_2800_v3_) in (d_2820_v21_) else (d_2821_v22_).cf71), (d_2808_v10_).f23])
            rhs372_ = d_2813_v15_
            rhs373_ = (self).f15
            rhs374_ = (default__.safeModuloInt(256, len((d_2799_v2_).set(default__.safeIndex((self).f15, len(d_2799_v2_)), d_2815_v16_))) if (len(d_2816_v17_)) >= (len(d_2817_v18_)) else default__.safeModuloInt((self).f15, len(d_2819_v20_)))
            rhs375_ = ((d_2822_v23_)[(d_2799_v2_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbpuyrmd")))] if ((d_2799_v2_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbpuyrmd")))) in (d_2822_v23_) else (self).f15)
            rhs376_ = d_2815_v16_
            lhs316_ = globalState
            lhs317_ = globalState
            lhs318_ = globalState
            d_2813_v15_ = rhs372_
            lhs316_.f6 = rhs373_
            lhs317_.f6 = rhs374_
            lhs318_.f6 = rhs375_
            d_2815_v16_ = rhs376_
        elif True:
            (globalState).f1 = d_2800_v3_
            r1 = not(d_2800_v3_)
            d_2823_v24_: _dafny.Array
            nw532_ = _dafny.Array(D16.default()(), 29)
            d_2823_v24_ = nw532_
            index472_ = default__.safeIndex(4, (d_2823_v24_).length(0))
            (d_2823_v24_)[index472_] = D16_DC44(_dafny.MultiSet([False, d_2800_v3_, d_2800_v3_]))
            d_2824_v25_: _dafny.MultiSet
            d_2824_v25_ = _dafny.MultiSet([d_2800_v3_, d_2800_v3_])
            d_2825_v26_: D16
            d_2825_v26_ = D16_DC44(d_2824_v25_)
            index473_ = default__.safeIndex(4, (d_2823_v24_).length(0))
            rhs377_ = ((self).f15) * (((0) - (default__.fm1(d_2800_v3_, globalState))) - (len(d_2797_v0_)))
            rhs378_ = (self).f15
            rhs379_ = d_2825_v26_
            rhs380_ = (0) - ((self).f15)
            rhs381_ = (d_2799_v2_ if d_2800_v3_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orpgjwi")))
            lhs319_ = globalState
            lhs320_ = d_2823_v24_
            lhs321_ = default__.safeIndex(4, (d_2823_v24_).length(0))
            r2 = rhs377_
            lhs319_.f6 = rhs378_
            lhs320_[lhs321_] = rhs379_
            r2 = rhs380_
            d_2799_v2_ = rhs381_
            if (((self).f15) - ((self).f15)) < ((self).f15):
                d_2826_v27_: C6
                nw533_ = C6()
                nw533_.ctor__()
                d_2826_v27_ = nw533_
                d_2827_v28_: _dafny.Array
                nw534_ = _dafny.Array(False, 26)
                d_2827_v28_ = nw534_
                d_2828_v29_: _dafny.Seq
                d_2828_v29_ = _dafny.SeqWithoutIsStrInference([d_2827_v28_, d_2827_v28_, d_2827_v28_])
                d_2829_v30_: _dafny.Set
                d_2829_v30_ = _dafny.Set({362, 187, len(d_2799_v2_), (self).f15, (self).f15})
                rhs382_ = len((d_2829_v30_) | (d_2829_v30_))
                rhs383_ = (self).f15
                rhs384_ = (self).f15
                rhs385_ = (d_2828_v29_) + (d_2828_v29_)
                lhs322_ = globalState
                lhs323_ = globalState
                lhs324_ = globalState
                lhs322_.f6 = rhs382_
                lhs323_.f6 = rhs383_
                lhs324_.f6 = rhs384_
                d_2828_v29_ = rhs385_
                d_2830_v31_: _dafny.MultiSet
                d_2830_v31_ = _dafny.MultiSet([(self).f15, 486, len(d_2799_v2_)])
                d_2831_v32_: str
                d_2831_v32_ = _dafny.CodePoint('j')
                d_2832_v33_: _dafny.Array
                nw535_ = _dafny.Array(None, 6)
                nw535_[int(0)] = default__.safeModuloInt((self).f15, (self).f15)
                nw535_[int(1)] = ((d_2830_v31_).cardinality if not(not(d_2800_v3_)) else (self).f15)
                nw535_[int(2)] = ((self).f15) + ((d_2826_v27_).fm9(_dafny.SeqWithoutIsStrInference([d_2831_v32_, d_2831_v32_]), globalState))
                nw535_[int(3)] = (0) - ((d_2797_v0_)[default__.safeIndex((self).f15, len(d_2797_v0_))])
                nw535_[int(4)] = (self).f15
                nw535_[int(5)] = (d_2824_v25_).cardinality
                d_2832_v33_ = nw535_
                index474_ = default__.safeIndex(303, (d_2832_v33_).length(0))
                (d_2832_v33_)[index474_] = len(d_2799_v2_)
                index475_ = default__.safeIndex(303, (d_2832_v33_).length(0))
                (d_2832_v33_)[index475_] = 133
                d_2831_v32_ = d_2831_v32_
                d_2833_v34_: _dafny.Map
                d_2833_v34_ = _dafny.Map({d_2800_v3_: _dafny.CodePoint('u')})
                d_2834_v35_: _dafny.Map
                d_2834_v35_ = _dafny.Map({d_2800_v3_: d_2800_v3_})
                index476_ = default__.safeIndex(303, (d_2832_v33_).length(0))
                rhs386_ = d_2833_v34_
                rhs387_ = ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsimrwma"))))) * ((d_2832_v33_)[default__.safeIndex(303, (d_2832_v33_).length(0))])
                rhs388_ = ((d_2834_v35_)[not(not(d_2800_v3_))] if (not(not(d_2800_v3_))) in (d_2834_v35_) else d_2800_v3_)
                rhs389_ = d_2799_v2_
                lhs325_ = d_2832_v33_
                lhs326_ = default__.safeIndex(303, (d_2832_v33_).length(0))
                lhs327_ = globalState
                d_2833_v34_ = rhs386_
                lhs325_[lhs326_] = rhs387_
                lhs327_.f1 = rhs388_
                d_2799_v2_ = rhs389_
            elif True:
                (globalState).f1 = (d_2800_v3_) not in (d_2824_v25_)
                d_2835_v36_: _dafny.Array
                nw536_ = _dafny.Array(False, 3)
                d_2835_v36_ = nw536_
                index477_ = default__.safeIndex(171, (d_2835_v36_).length(0))
                (d_2835_v36_)[index477_] = d_2800_v3_
                index478_ = default__.safeIndex(171, (d_2835_v36_).length(0))
                (d_2835_v36_)[index478_] = d_2800_v3_
                d_2836_v37_: C4
                nw537_ = C4()
                nw537_.ctor__()
                d_2836_v37_ = nw537_
                d_2837_v38_: _dafny.Array
                nw538_ = _dafny.Array(int(0), 18)
                d_2837_v38_ = nw538_
                index479_ = default__.safeIndex(128, (d_2837_v38_).length(0))
                (d_2837_v38_)[index479_] = (self).f15
                d_2838_v39_: C5
                nw539_ = C5()
                nw539_.ctor__((self).f15)
                d_2838_v39_ = nw539_
                d_2839_v40_: _dafny.Map
                d_2839_v40_ = _dafny.Map({d_2824_v25_: d_2838_v39_})
                index480_ = default__.safeIndex(128, (d_2837_v38_).length(0))
                (d_2837_v38_)[index480_] = len((d_2839_v40_) | (d_2839_v40_))
                (globalState).f6 = (0) - ((690) + (len((_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_2835_v36_)[default__.safeIndex(171, (d_2835_v36_).length(0))]: d_2800_v3_}) for d_2840_i2_ in range(default__.abs(-429))])) + (default__.fm50((d_2838_v39_).f16, -868, globalState)))))
            d_2841_v41_: _dafny.Array
            def lambda138_(d_2842_v3_):
                def lambda139_(d_2843_i3_):
                    return d_2842_v3_

                return lambda139_

            init80_ = lambda138_(d_2800_v3_)
            nw540_ = _dafny.Array(None, 9)
            for i0_80_ in range(nw540_.length(0)):
                nw540_[i0_80_] = init80_(i0_80_)
            d_2841_v41_ = nw540_
            index481_ = default__.safeIndex(203, (d_2841_v41_).length(0))
            (d_2841_v41_)[index481_] = d_2800_v3_
            index482_ = default__.safeIndex(203, (d_2841_v41_).length(0))
            rhs390_ = not(d_2800_v3_)
            rhs391_ = not(d_2800_v3_)
            lhs328_ = d_2841_v41_
            lhs329_ = default__.safeIndex(203, (d_2841_v41_).length(0))
            r1 = rhs390_
            lhs328_[lhs329_] = rhs391_
        hi17_ = (0) - ((self).f15)
        for d_2844_i4_ in range((self).f15, hi17_):
            d_2845_v42_: _dafny.Array
            def lambda140_(d_2846_i5_):
                return (d_2846_i5_) - ((0) - ((self).f15))

            init81_ = lambda140_
            nw541_ = _dafny.Array(None, 18)
            for i0_81_ in range(nw541_.length(0)):
                nw541_[i0_81_] = init81_(i0_81_)
            d_2845_v42_ = nw541_
            d_2847_v43_: str
            d_2847_v43_ = _dafny.CodePoint('j')
            index483_ = default__.safeIndex(710, (d_2845_v42_).length(0))
            (d_2845_v42_)[index483_] = default__.safeDivisionInt((self).f15, (0) - (len((d_2799_v2_).set(default__.safeIndex(d_2844_i4_, len(d_2799_v2_)), d_2847_v43_))))
            d_2848_v44_: _dafny.Map
            d_2848_v44_ = _dafny.Map({d_2800_v3_: d_2844_i4_})
            index484_ = default__.safeIndex(710, (d_2845_v42_).length(0))
            (d_2845_v42_)[index484_] = (len(d_2848_v44_)) - (len(d_2799_v2_))
            rhs392_ = (((d_2799_v2_ if d_2800_v3_ else d_2799_v2_)) + (d_2799_v2_)).set(default__.safeIndex(len((d_2799_v2_) + (d_2799_v2_)), len(((d_2799_v2_ if d_2800_v3_ else d_2799_v2_)) + (d_2799_v2_))), d_2847_v43_)
            rhs393_ = d_2800_v3_
            rhs394_ = d_2799_v2_
            rhs395_ = d_2800_v3_
            rhs396_ = (self).f15
            lhs330_ = globalState
            d_2799_v2_ = rhs392_
            lhs330_.f1 = rhs393_
            d_2799_v2_ = rhs394_
            r1 = rhs395_
            r2 = rhs396_
            d_2849_v45_: _dafny.Map
            d_2849_v45_ = _dafny.Map({d_2800_v3_: d_2800_v3_})
            d_2850_v46_: D1
            d_2850_v46_ = D1_DC3(d_2849_v45_)
            source26_ = d_2850_v46_
            if source26_.is_DC4:
                d_2851___mcc_h0_ = source26_.cf4
                d_2852_cf4_ = d_2851___mcc_h0_
                d_2853_v47_: _dafny.Map
                d_2853_v47_ = _dafny.Map({False: d_2799_v2_})
                d_2799_v2_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhk"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fiwsofh"))) if d_2800_v3_ else (((d_2853_v47_)[d_2800_v3_] if (d_2800_v3_) in (d_2853_v47_) else _dafny.SeqWithoutIsStrInference([d_2852_cf4_ for d_2854_i6_ in range(default__.abs(720))]))).set(default__.safeIndex((self).f15, len(((d_2853_v47_)[d_2800_v3_] if (d_2800_v3_) in (d_2853_v47_) else _dafny.SeqWithoutIsStrInference([d_2852_cf4_ for d_2855_i6_ in range(default__.abs(720))])))), d_2852_cf4_))
                d_2856_v48_: C4
                nw542_ = C4()
                nw542_.ctor__()
                d_2856_v48_ = nw542_
                d_2857_v49_: _dafny.Seq
                d_2857_v49_ = _dafny.SeqWithoutIsStrInference([d_2845_v42_])
                (globalState).f6 = (0) - ((((d_2845_v42_)[default__.safeIndex(710, (d_2845_v42_).length(0))]) * (len(d_2857_v49_))) * (651))
                d_2858_v50_: _dafny.Array
                nw543_ = _dafny.Array(None, 12)
                d_2858_v50_ = nw543_
                nw544_ = _dafny.Array(None, 22)
                d_2858_v50_ = nw544_
            elif True:
                d_2859___mcc_h1_ = source26_.cf3
                d_2860_cf3_ = d_2859___mcc_h1_
                d_2861_v51_: _dafny.Set
                d_2861_v51_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([661, d_2844_i4_]))})
                (globalState).f6 = (len(d_2797_v0_) if d_2800_v3_ else len(d_2861_v51_))
                d_2862_v52_: _dafny.Map
                d_2862_v52_ = _dafny.Map({d_2845_v42_: d_2799_v2_})
                d_2863_v53_: _dafny.Array
                def lambda141_(d_2864_v3_):
                    def lambda142_(d_2865_i7_):
                        return not(d_2864_v3_)

                    return lambda142_

                init82_ = lambda141_(d_2800_v3_)
                nw545_ = _dafny.Array(None, 7)
                for i0_82_ in range(nw545_.length(0)):
                    nw545_[i0_82_] = init82_(i0_82_)
                d_2863_v53_ = nw545_
                index485_ = default__.safeIndex(804, (d_2863_v53_).length(0))
                (d_2863_v53_)[index485_] = d_2800_v3_
                d_2866_v54_: _dafny.Array
                nw546_ = _dafny.Array(None, 5)
                nw546_[int(0)] = d_2863_v53_
                nw546_[int(1)] = d_2863_v53_
                nw546_[int(2)] = d_2863_v53_
                nw546_[int(3)] = d_2863_v53_
                nw546_[int(4)] = d_2863_v53_
                d_2866_v54_ = nw546_
                d_2867_v55_: _dafny.Map
                d_2867_v55_ = _dafny.Map({default__.fm2(globalState): d_2863_v53_})
                index486_ = default__.safeIndex(196, (d_2866_v54_).length(0))
                (d_2866_v54_)[index486_] = ((d_2867_v55_)[d_2800_v3_] if (d_2800_v3_) in (d_2867_v55_) else d_2863_v53_)
                index487_ = default__.safeIndex(804, (d_2863_v53_).length(0))
                index488_ = default__.safeIndex(196, (d_2866_v54_).length(0))
                index489_ = default__.safeIndex(710, (d_2845_v42_).length(0))
                rhs397_ = d_2862_v52_
                rhs398_ = (len(d_2861_v51_)) == (((self).f15) - (d_2844_i4_))
                rhs399_ = d_2863_v53_
                rhs400_ = (len(d_2799_v2_)) * ((0) - (default__.fm1(True, globalState)))
                lhs331_ = d_2863_v53_
                lhs332_ = default__.safeIndex(804, (d_2863_v53_).length(0))
                lhs333_ = d_2866_v54_
                lhs334_ = default__.safeIndex(196, (d_2866_v54_).length(0))
                lhs335_ = d_2845_v42_
                lhs336_ = default__.safeIndex(710, (d_2845_v42_).length(0))
                d_2862_v52_ = rhs397_
                lhs331_[lhs332_] = rhs398_
                lhs333_[lhs334_] = rhs399_
                lhs335_[lhs336_] = rhs400_
                d_2868_v56_: _dafny.Array
                nw547_ = _dafny.Array(None, 19)
                nw547_[int(0)] = d_2845_v42_
                nw547_[int(1)] = d_2845_v42_
                nw547_[int(2)] = d_2845_v42_
                nw547_[int(3)] = d_2845_v42_
                nw547_[int(4)] = d_2845_v42_
                nw547_[int(5)] = d_2845_v42_
                nw547_[int(6)] = d_2845_v42_
                nw547_[int(7)] = d_2845_v42_
                nw547_[int(8)] = d_2845_v42_
                nw547_[int(9)] = d_2845_v42_
                nw547_[int(10)] = d_2845_v42_
                nw547_[int(11)] = d_2845_v42_
                nw547_[int(12)] = d_2845_v42_
                nw547_[int(13)] = d_2845_v42_
                nw547_[int(14)] = d_2845_v42_
                nw547_[int(15)] = d_2845_v42_
                nw547_[int(16)] = d_2845_v42_
                nw547_[int(17)] = d_2845_v42_
                nw547_[int(18)] = d_2845_v42_
                d_2868_v56_ = nw547_
                d_2868_v56_ = d_2868_v56_
                d_2869_v57_: _dafny.MultiSet
                d_2869_v57_ = _dafny.MultiSet([d_2863_v53_, (d_2866_v54_)[default__.safeIndex(196, (d_2866_v54_).length(0))], d_2863_v53_, (d_2866_v54_)[default__.safeIndex(196, (d_2866_v54_).length(0))]])
                rhs401_ = (d_2845_v42_)[default__.safeIndex(710, (d_2845_v42_).length(0))]
                rhs402_ = (d_2869_v57_).issubset(d_2869_v57_)
                r2 = rhs401_
                r1 = rhs402_
            d_2870_v58_: _dafny.MultiSet
            d_2870_v58_ = _dafny.MultiSet([d_2800_v3_, d_2800_v3_, d_2800_v3_, d_2800_v3_, d_2800_v3_])
            d_2871_v59_: _dafny.Map
            d_2871_v59_ = _dafny.Map({(d_2845_v42_)[default__.safeIndex(710, (d_2845_v42_).length(0))]: d_2870_v58_})
            d_2872_v60_: D20
            d_2872_v60_ = D20_DC50(not(d_2800_v3_), (self).f15, (d_2845_v42_)[default__.safeIndex(710, (d_2845_v42_).length(0))], d_2800_v3_, not((d_2847_v43_) not in (d_2799_v2_)))
            d_2873_v61_: _dafny.Map
            d_2873_v61_ = _dafny.Map({d_2800_v3_: d_2871_v59_})
            d_2874_v62_: _dafny.Array
            nw548_ = _dafny.Array(None, 29)
            nw548_[int(0)] = d_2800_v3_
            nw548_[int(1)] = d_2800_v3_
            nw548_[int(2)] = d_2800_v3_
            nw548_[int(3)] = d_2800_v3_
            nw548_[int(4)] = d_2800_v3_
            nw548_[int(5)] = (p0)[default__.safeIndex((self).f15, len(p0))]
            nw548_[int(6)] = d_2800_v3_
            nw548_[int(7)] = True
            nw548_[int(8)] = d_2800_v3_
            nw548_[int(9)] = False
            nw548_[int(10)] = d_2800_v3_
            nw548_[int(11)] = not(d_2800_v3_)
            nw548_[int(12)] = d_2800_v3_
            nw548_[int(13)] = d_2800_v3_
            nw548_[int(14)] = True
            nw548_[int(15)] = d_2800_v3_
            nw548_[int(16)] = not(d_2800_v3_)
            nw548_[int(17)] = True
            nw548_[int(18)] = d_2800_v3_
            nw548_[int(19)] = d_2800_v3_
            nw548_[int(20)] = False
            nw548_[int(21)] = d_2800_v3_
            nw548_[int(22)] = d_2800_v3_
            nw548_[int(23)] = not(d_2800_v3_)
            nw548_[int(24)] = d_2800_v3_
            nw548_[int(25)] = d_2800_v3_
            nw548_[int(26)] = d_2800_v3_
            nw548_[int(27)] = d_2800_v3_
            nw548_[int(28)] = d_2800_v3_
            d_2874_v62_ = nw548_
            d_2875_v63_: _dafny.Array
            nw549_ = _dafny.Array(_dafny.Array(None, 0), 27)
            d_2875_v63_ = nw549_
            d_2876_v64_: _dafny.Seq
            d_2876_v64_ = _dafny.SeqWithoutIsStrInference([d_2875_v63_, d_2875_v63_])
            d_2877_v65_: D5
            d_2877_v65_ = D5_DC10((d_2876_v64_)[default__.safeIndex((0) - ((d_2845_v42_)[default__.safeIndex(710, (d_2845_v42_).length(0))]), len(d_2876_v64_))])
            d_2878_v66_: _dafny.Map
            d_2878_v66_ = _dafny.Map({d_2874_v62_: d_2877_v65_})
            d_2879_v67_: _dafny.Map
            d_2879_v67_ = _dafny.Map({(d_2845_v42_)[default__.safeIndex(710, (d_2845_v42_).length(0))]: d_2800_v3_})
            d_2880_v68_: D0
            d_2880_v68_ = D0_DC1((self).f15)
            rhs403_ = (((d_2873_v61_)[not(d_2800_v3_)] if (not(d_2800_v3_)) in (d_2873_v61_) else default__.fm51(-345, globalState))).set(len((d_2878_v66_) | (d_2878_v66_)), d_2870_v58_)
            rhs404_ = D20_DC50(not(d_2800_v3_), d_2844_i4_, (self).f15, (d_2844_i4_) not in ((d_2879_v67_).set((d_2845_v42_)[default__.safeIndex(710, (d_2845_v42_).length(0))], d_2800_v3_)), not(not((False) == (False))))
            rhs405_ = (d_2848_v44_) | (d_2848_v44_)
            rhs406_ = _dafny.Map({d_2800_v3_: (d_2880_v68_).cf1})
            d_2871_v59_ = rhs403_
            d_2872_v60_ = rhs404_
            d_2848_v44_ = rhs405_
            d_2848_v44_ = rhs406_
        d_2881_v69_: _dafny.MultiSet
        d_2881_v69_ = _dafny.MultiSet([d_2800_v3_, d_2800_v3_, d_2800_v3_, d_2800_v3_])
        d_2882_v70_: D8
        d_2882_v70_ = D8_DC20((d_2881_v69_).cardinality, -620, (self).f15)
        source27_ = default__.fm52(d_2882_v70_, (self).f15, globalState)
        if source27_.is_DC7:
            d_2883___mcc_h2_ = source27_.cf7
            d_2884_cf7_ = d_2883___mcc_h2_
            d_2885_v71_: _dafny.Array
            def lambda143_(d_2886_v3_):
                def lambda144_(d_2887_i8_):
                    return d_2886_v3_

                return lambda144_

            init83_ = lambda143_(d_2800_v3_)
            nw550_ = _dafny.Array(None, 8)
            for i0_83_ in range(nw550_.length(0)):
                nw550_[i0_83_] = init83_(i0_83_)
            d_2885_v71_ = nw550_
            index490_ = default__.safeIndex(37, (d_2885_v71_).length(0))
            (d_2885_v71_)[index490_] = d_2800_v3_
            d_2888_v72_: _dafny.MultiSet
            d_2888_v72_ = _dafny.MultiSet([d_2799_v2_, d_2799_v2_, d_2799_v2_, d_2799_v2_])
            index491_ = default__.safeIndex(37, (d_2885_v71_).length(0))
            (d_2885_v71_)[index491_] = (d_2888_v72_).isdisjoint((d_2888_v72_ if d_2800_v3_ else d_2888_v72_))
            d_2800_v3_ = (d_2885_v71_)[default__.safeIndex(37, (d_2885_v71_).length(0))]
            d_2889_v73_: _dafny.Set
            d_2889_v73_ = _dafny.Set({default__.fm1(True, globalState), d_2884_cf7_})
            d_2890_v75_: _dafny.Seq
            def iife294_():
                coll90_ = _dafny.Set()
                compr_90_: int
                for compr_90_ in _dafny.IntegerRange(616, 728):
                    d_2891_v74_: int = compr_90_
                    if ((616) <= (d_2891_v74_)) and ((d_2891_v74_) < (728)):
                        coll90_ = coll90_.union(_dafny.Set([(d_2891_v74_) - (d_2884_cf7_)]))
                return _dafny.Set(coll90_)
            d_2890_v75_ = _dafny.SeqWithoutIsStrInference([d_2889_v73_, iife294_()
            , (d_2889_v73_) - (d_2889_v73_), d_2889_v73_])
            d_2890_v75_ = ((d_2890_v75_) + (_dafny.SeqWithoutIsStrInference([d_2889_v73_ for d_2892_i9_ in range(default__.abs(755))]))) + (d_2890_v75_)
            r0 = not(d_2800_v3_)
        elif source27_.is_DC8:
            d_2893___mcc_h3_ = source27_.cf8
            d_2894___mcc_h4_ = source27_.cf9
            d_2895___mcc_h5_ = source27_.cf10
            d_2896___mcc_h6_ = source27_.cf11
            d_2897_cf11_ = d_2896___mcc_h6_
            d_2898_cf10_ = d_2895___mcc_h5_
            d_2899_cf9_ = d_2894___mcc_h4_
            d_2900_cf8_ = d_2893___mcc_h3_
            d_2901_v76_: _dafny.Set
            d_2901_v76_ = _dafny.Set({(self).f15, (self).f15})
            d_2902_v77_: _dafny.Map
            d_2902_v77_ = _dafny.Map({((self).f15) == (-976): d_2901_v76_})
            def iife295_():
                coll91_ = _dafny.Set()
                compr_91_: int
                for compr_91_ in _dafny.IntegerRange(42, 375):
                    d_2903_v78_: int = compr_91_
                    if ((42) <= (d_2903_v78_)) and ((d_2903_v78_) < (375)):
                        coll91_ = coll91_.union(_dafny.Set([default__.safeModuloInt(d_2903_v78_, len(p0))]))
                return _dafny.Set(coll91_)
            d_2902_v77_ = (d_2902_v77_).set(default__.fm2(globalState), iife295_()
            )
            r2 = (self).f15
            d_2899_cf9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxctkag"))
            d_2904_v79_: C1
            nw551_ = C1()
            nw551_.ctor__(345)
            d_2904_v79_ = nw551_
        elif True:
            d_2905___mcc_h7_ = source27_.cf6
            d_2906_cf6_ = d_2905___mcc_h7_
            d_2907_v80_: D15
            d_2907_v80_ = D15_DC43(d_2800_v3_, (d_2800_v3_) and (d_2800_v3_), (d_2800_v3_ if d_2800_v3_ else d_2800_v3_))
            d_2908_v81_: _dafny.Array
            nw552_ = _dafny.Array(None, 6)
            nw552_[int(0)] = d_2800_v3_
            nw552_[int(1)] = d_2800_v3_
            nw552_[int(2)] = d_2800_v3_
            nw552_[int(3)] = d_2800_v3_
            nw552_[int(4)] = d_2800_v3_
            nw552_[int(5)] = d_2800_v3_
            d_2908_v81_ = nw552_
            d_2909_v82_: _dafny.Map
            d_2909_v82_ = _dafny.Map({(self).f15: d_2908_v81_})
            rhs407_ = (D15_DC43(d_2800_v3_, d_2800_v3_, (self).fm4((self).f15, globalState))).cf74
            rhs408_ = False
            rhs409_ = default__.fm53(((self).f15) in (d_2909_v82_), (self).f15, globalState)
            r1 = rhs407_
            r1 = rhs408_
            d_2907_v80_ = rhs409_
            d_2910_v83_: _dafny.Map
            d_2910_v83_ = _dafny.Map({d_2800_v3_: d_2800_v3_})
            d_2911_v84_: _dafny.Set
            d_2911_v84_ = _dafny.Set({d_2800_v3_, d_2800_v3_, d_2800_v3_})
            d_2912_v85_: _dafny.Seq
            d_2912_v85_ = _dafny.SeqWithoutIsStrInference([d_2800_v3_, d_2800_v3_, d_2800_v3_, (_dafny.Map({True: d_2800_v3_})) == (d_2910_v83_), not((d_2911_v84_).isdisjoint(default__.fm31((self).f15, d_2800_v3_, globalState)))])
            d_2913_v86_: _dafny.Seq
            d_2913_v86_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([default__.fm2(globalState), not(d_2800_v3_), d_2800_v3_, d_2800_v3_]), p0])
            d_2912_v85_ = (d_2913_v86_)[default__.safeIndex((0) - (-413), len(d_2913_v86_))]
            index492_ = default__.safeIndex(823, (d_2906_cf6_).length(0))
            (d_2906_cf6_)[index492_] = len(d_2799_v2_)
            index493_ = default__.safeIndex(823, (d_2906_cf6_).length(0))
            (d_2906_cf6_)[index493_] = (self).f15
            d_2914_v87_: C6
            nw553_ = C6()
            nw553_.ctor__()
            d_2914_v87_ = nw553_
            d_2914_v87_ = d_2914_v87_
        d_2915_v88_: _dafny.Map
        d_2915_v88_ = _dafny.Map({d_2799_v2_: not(d_2800_v3_)})
        d_2916_v89_: bool
        d_2917_v90_: bool
        out26_: bool
        out27_: bool
        out26_, out27_ = (self).m4(((d_2881_v69_).cardinality) != (len(d_2915_v88_)), globalState)
        d_2916_v89_ = out26_
        d_2917_v90_ = out27_
        d_2918_v91_: _dafny.Set
        d_2918_v91_ = _dafny.Set({(self).f15})
        d_2919_v92_: str
        d_2919_v92_ = _dafny.CodePoint('y')
        d_2920_v93_: D5
        d_2920_v93_ = D5_DC12(len(d_2918_v91_), d_2800_v3_, d_2919_v92_)
        source28_ = d_2920_v93_
        if source28_.is_DC11:
            d_2921___mcc_h8_ = source28_.cf14
            d_2922_cf14_ = d_2921___mcc_h8_
            if d_2916_v89_:
                d_2923_v94_: _dafny.Array
                nw554_ = _dafny.Array(_dafny.Seq({}), 14)
                d_2923_v94_ = nw554_
                d_2924_v95_: _dafny.Map
                d_2924_v95_ = _dafny.Map({False: d_2917_v90_})
                d_2925_v96_: _dafny.Seq
                d_2925_v96_ = _dafny.SeqWithoutIsStrInference([d_2924_v95_])
                index494_ = default__.safeIndex(76, (d_2923_v94_).length(0))
                (d_2923_v94_)[index494_] = (d_2925_v96_) + (d_2925_v96_)
                index495_ = default__.safeIndex(76, (d_2923_v94_).length(0))
                (d_2923_v94_)[index495_] = _dafny.SeqWithoutIsStrInference([d_2924_v95_, d_2924_v95_, (d_2924_v95_) | (d_2924_v95_)])
                d_2926_v97_: _dafny.Array
                nw555_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
                d_2926_v97_ = nw555_
                index496_ = default__.safeIndex(876, (d_2926_v97_).length(0))
                (d_2926_v97_)[index496_] = d_2799_v2_
                index497_ = default__.safeIndex(876, (d_2926_v97_).length(0))
                (d_2926_v97_)[index497_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mltwv"))) + (d_2799_v2_) if True else (d_2799_v2_).set(default__.safeIndex(d_2922_cf14_, len(d_2799_v2_)), d_2919_v92_))
                d_2927_v98_: _dafny.Array
                nw556_ = _dafny.Array(_dafny.Set({}), 22)
                d_2927_v98_ = nw556_
                d_2928_v99_: _dafny.Map
                d_2928_v99_ = _dafny.Map({(d_2927_v98_ if d_2800_v3_ else d_2927_v98_): default__.fm1(d_2916_v89_, globalState)})
                d_2928_v99_ = (d_2928_v99_).set(d_2927_v98_, (self).f15)
                d_2922_cf14_ = (0) - ((self).f15)
                d_2929_v100_: _dafny.Map
                d_2929_v100_ = _dafny.Map({((self).f15) * (d_2922_cf14_): _dafny.CodePoint('b')})
                d_2929_v100_ = ((_dafny.Map({len((d_2926_v97_)[default__.safeIndex(876, (d_2926_v97_).length(0))]): d_2919_v92_})) | (default__.fm54(d_2924_v95_, globalState))) | (d_2929_v100_)
            elif True:
                d_2930_v101_: _dafny.Set
                d_2930_v101_ = _dafny.Set({d_2917_v90_, not((d_2916_v89_ if d_2917_v90_ else False))})
                d_2931_v102_: _dafny.Array
                nw557_ = _dafny.Array(None, 1)
                nw557_[int(0)] = _dafny.Map({_dafny.CodePoint('v'): (self).f15})
                d_2931_v102_ = nw557_
                d_2932_v103_: _dafny.MultiSet
                d_2932_v103_ = _dafny.MultiSet([len(d_2799_v2_), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qyropc"))).set(default__.safeIndex((self).f15, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qyropc")))), d_2919_v92_)), d_2922_cf14_])
                rhs410_ = d_2930_v101_
                rhs411_ = d_2931_v102_
                rhs412_ = ((self).f15) not in (d_2932_v103_)
                rhs413_ = d_2919_v92_
                rhs414_ = d_2919_v92_
                lhs337_ = globalState
                d_2930_v101_ = rhs410_
                d_2931_v102_ = rhs411_
                lhs337_.f1 = rhs412_
                d_2919_v92_ = rhs413_
                d_2919_v92_ = rhs414_
                (globalState).f6 = len(default__.fm13((d_2916_v89_) == (d_2800_v3_), d_2919_v92_, 218, globalState))
                d_2933_v104_: _dafny.Array
                nw558_ = _dafny.Array(None, 7)
                nw558_[int(0)] = d_2881_v69_
                nw558_[int(1)] = _dafny.MultiSet([(D20_DC50(d_2917_v90_, (0) - (d_2922_cf14_), (self).f15, not(d_2917_v90_), d_2800_v3_)).cf90])
                nw558_[int(2)] = _dafny.MultiSet([True])
                nw558_[int(3)] = d_2881_v69_
                nw558_[int(4)] = d_2881_v69_
                nw558_[int(5)] = (d_2881_v69_).intersection(d_2881_v69_)
                nw558_[int(6)] = d_2881_v69_
                d_2933_v104_ = nw558_
                d_2933_v104_ = d_2933_v104_
                (globalState).f6 = (self).f15
                d_2917_v90_ = d_2917_v90_
            d_2934_v105_: _dafny.Set
            d_2934_v105_ = _dafny.Set({d_2916_v89_})
            if (d_2800_v3_) or ((d_2934_v105_).ispropersubset(d_2934_v105_)):
                d_2935_v106_: _dafny.MultiSet
                d_2935_v106_ = _dafny.MultiSet([d_2922_cf14_, (self).f15])
                d_2935_v106_ = d_2935_v106_
                d_2800_v3_ = d_2800_v3_
                d_2936_v107_: C6
                nw559_ = C6()
                nw559_.ctor__()
                d_2936_v107_ = nw559_
                d_2936_v107_ = d_2936_v107_
                d_2800_v3_ = default__.fm2(globalState)
                d_2937_v108_: C8
                nw560_ = C8()
                nw560_.ctor__()
                d_2937_v108_ = nw560_
            elif True:
                (globalState).f1 = ((d_2881_v69_).ispropersubset(d_2881_v69_)) and (d_2800_v3_)
                d_2938_v109_: C6
                nw561_ = C6()
                nw561_.ctor__()
                d_2938_v109_ = nw561_
                d_2939_v110_: _dafny.Map
                d_2939_v110_ = _dafny.Map({default__.safeDivisionInt(d_2922_cf14_, d_2922_cf14_): d_2938_v109_})
                d_2939_v110_ = d_2939_v110_
                d_2940_v111_: _dafny.MultiSet
                d_2940_v111_ = _dafny.MultiSet([default__.safeModuloInt(d_2922_cf14_, -338), default__.safeDivisionInt(d_2922_cf14_, len(p0)), default__.safeDivisionInt((self).f15, (self).f15)])
                d_2940_v111_ = d_2940_v111_
                d_2941_v112_: _dafny.Map
                d_2941_v112_ = _dafny.Map({d_2916_v89_: d_2934_v105_})
                (globalState).f1 = not((((d_2941_v112_)[d_2917_v90_] if (d_2917_v90_) in (d_2941_v112_) else d_2934_v105_)).ispropersubset((d_2934_v105_).intersection(d_2934_v105_)))
                (globalState).f6 = ((0) - ((0) - ((self).f15))) * ((self).f15)
            d_2942_v113_: T0
            nw562_ = C1()
            nw562_.ctor__(len(d_2799_v2_))
            d_2942_v113_ = nw562_
            d_2943_v114_: _dafny.Seq
            d_2943_v114_ = _dafny.SeqWithoutIsStrInference([d_2942_v113_, d_2942_v113_])
            d_2944_v115_: _dafny.Seq
            d_2944_v115_ = _dafny.SeqWithoutIsStrInference([default__.fm32((self).f15, (self).f15, d_2800_v3_, globalState), d_2797_v0_])
            d_2945_v116_: _dafny.Map
            d_2945_v116_ = _dafny.Map({d_2916_v89_: d_2917_v90_})
            d_2946_v117_: _dafny.Map
            d_2946_v117_ = _dafny.Map({default__.fm2(globalState): (self).f15})
            d_2947_v118_: _dafny.Array
            nw563_ = _dafny.Array(None, 24)
            nw563_[int(0)] = (len(d_2943_v114_)) * ((self).f15)
            nw563_[int(1)] = d_2922_cf14_
            nw563_[int(2)] = len(d_2797_v0_)
            nw563_[int(3)] = (self).f15
            nw563_[int(4)] = (len(p0)) - ((self).f15)
            nw563_[int(5)] = d_2922_cf14_
            nw563_[int(6)] = d_2922_cf14_
            nw563_[int(7)] = d_2922_cf14_
            nw563_[int(8)] = (self).f15
            nw563_[int(9)] = default__.safeDivisionInt(-849, (self).f15)
            nw563_[int(10)] = d_2922_cf14_
            nw563_[int(11)] = -277
            nw563_[int(12)] = d_2922_cf14_
            nw563_[int(13)] = len(((d_2944_v115_)[default__.safeIndex((self).f15, len(d_2944_v115_))]) + (_dafny.SeqWithoutIsStrInference([(self).f15 for d_2948_i10_ in range(default__.abs(265))])))
            nw563_[int(14)] = d_2922_cf14_
            nw563_[int(15)] = d_2922_cf14_
            nw563_[int(16)] = len(d_2945_v116_)
            nw563_[int(17)] = ((d_2946_v117_)[False] if (False) in (d_2946_v117_) else (0) - ((self).f15))
            nw563_[int(18)] = len((d_2934_v105_) - (d_2934_v105_))
            nw563_[int(19)] = d_2922_cf14_
            nw563_[int(20)] = d_2922_cf14_
            nw563_[int(21)] = (d_2922_cf14_) - (d_2922_cf14_)
            nw563_[int(22)] = 355
            nw563_[int(23)] = (self).f15
            d_2947_v118_ = nw563_
            index498_ = default__.safeIndex(508, (d_2947_v118_).length(0))
            (d_2947_v118_)[index498_] = (self).f15
            index499_ = default__.safeIndex(508, (d_2947_v118_).length(0))
            (d_2947_v118_)[index499_] = d_2922_cf14_
            d_2881_v69_ = d_2881_v69_
        elif source28_.is_DC12:
            d_2949___mcc_h9_ = source28_.cf15
            d_2950___mcc_h10_ = source28_.cf16
            d_2951___mcc_h11_ = source28_.cf17
            d_2952_cf17_ = d_2951___mcc_h11_
            d_2953_cf16_ = d_2950___mcc_h10_
            d_2954_cf15_ = d_2949___mcc_h9_
            d_2955_v119_: C8
            nw564_ = C8()
            nw564_.ctor__()
            d_2955_v119_ = nw564_
            d_2956_v120_: _dafny.MultiSet
            d_2956_v120_ = _dafny.MultiSet([d_2952_cf17_, d_2952_cf17_])
            d_2957_v121_: D20
            d_2957_v121_ = D20_DC50(d_2953_cf16_, d_2954_cf15_, (self).f15, d_2953_cf16_, not(d_2800_v3_))
            d_2958_v122_: _dafny.Map
            d_2958_v122_ = _dafny.Map({d_2800_v3_: d_2954_cf15_})
            d_2959_v123_: _dafny.Array
            nw565_ = _dafny.Array(None, 24)
            nw565_[int(0)] = True
            nw565_[int(1)] = d_2916_v89_
            nw565_[int(2)] = (True) or (not(not(d_2800_v3_)))
            nw565_[int(3)] = d_2917_v90_
            nw565_[int(4)] = True
            nw565_[int(5)] = (d_2916_v89_) or (d_2917_v90_)
            nw565_[int(6)] = (d_2916_v89_) and (d_2917_v90_)
            nw565_[int(7)] = (_dafny.MultiSet([_dafny.CodePoint('h'), default__.fm29(d_2916_v89_, (d_2881_v69_).cardinality, (self).f15, d_2800_v3_, globalState), d_2952_cf17_, default__.fm29(default__.fm2(globalState), (d_2920_v93_).cf15, (self).f15, d_2917_v90_, globalState), d_2952_cf17_])).ispropersubset(d_2956_v120_)
            nw565_[int(8)] = d_2800_v3_
            nw565_[int(9)] = d_2953_cf16_
            nw565_[int(10)] = (p0)[default__.safeIndex(len(_dafny.Map({d_2953_cf16_: (self).f15})), len(p0))]
            nw565_[int(11)] = d_2953_cf16_
            nw565_[int(12)] = not((default__.fm2(globalState)) == (d_2916_v89_))
            nw565_[int(13)] = (d_2957_v121_).cf86
            nw565_[int(14)] = d_2953_cf16_
            nw565_[int(15)] = d_2953_cf16_
            nw565_[int(16)] = d_2953_cf16_
            nw565_[int(17)] = (((d_2958_v122_)[d_2953_cf16_] if (d_2953_cf16_) in (d_2958_v122_) else d_2954_cf15_)) >= ((self).f15)
            nw565_[int(18)] = d_2917_v90_
            nw565_[int(19)] = d_2916_v89_
            nw565_[int(20)] = (d_2800_v3_) and (d_2953_cf16_)
            nw565_[int(21)] = not (d_2917_v90_) or (d_2953_cf16_)
            nw565_[int(22)] = d_2916_v89_
            nw565_[int(23)] = d_2800_v3_
            d_2959_v123_ = nw565_
            index500_ = default__.safeIndex(321, (d_2959_v123_).length(0))
            (d_2959_v123_)[index500_] = d_2800_v3_
            index501_ = default__.safeIndex(321, (d_2959_v123_).length(0))
            (d_2959_v123_)[index501_] = d_2953_cf16_
            d_2960_v124_: _dafny.Array
            nw566_ = _dafny.Array(int(0), 25)
            d_2960_v124_ = nw566_
            d_2961_v125_: D3
            d_2961_v125_ = D3_DC6(d_2960_v124_)
            source29_ = d_2961_v125_
            if source29_.is_DC7:
                d_2962___mcc_h14_ = source29_.cf7
                d_2963_cf7_ = d_2962___mcc_h14_
                d_2964_v126_: D16
                d_2964_v126_ = D16_DC45(478, d_2919_v92_, d_2954_cf15_, d_2917_v90_)
                d_2965_v127_: _dafny.Map
                d_2965_v127_ = _dafny.Map({d_2954_cf15_: d_2964_v126_})
                d_2966_v128_: _dafny.Set
                d_2966_v128_ = _dafny.Set({default__.fm2(globalState)})
                pat_let_tv71_ = d_2954_cf15_
                def iife296_(_pat_let102_0):
                    def iife297_(d_2967_dt__update__tmp_h0_):
                        def iife298_(_pat_let103_0):
                            def iife299_(d_2968_dt__update_hcf80_h0_):
                                return D16_DC45((d_2967_dt__update__tmp_h0_).cf78, (d_2967_dt__update__tmp_h0_).cf79, d_2968_dt__update_hcf80_h0_, (d_2967_dt__update__tmp_h0_).cf81)
                            return iife299_(_pat_let103_0)
                        return iife298_(pat_let_tv71_)
                    return iife297_(_pat_let102_0)
                d_2965_v127_ = (d_2965_v127_).set(len(d_2966_v128_), iife296_(d_2964_v126_))
                d_2969_v129_: D1
                d_2969_v129_ = D1_DC4(d_2919_v92_)
                d_2969_v129_ = d_2969_v129_
                d_2970_v130_: D3
                d_2970_v130_ = D3_DC7((d_2797_v0_)[default__.safeIndex(-905, len(d_2797_v0_))])
                (globalState).f6 = (d_2970_v130_).cf7
                d_2971_v131_: C5
                nw567_ = C5()
                nw567_.ctor__((self).f15)
                d_2971_v131_ = nw567_
            elif source29_.is_DC8:
                d_2972___mcc_h15_ = source29_.cf8
                d_2973___mcc_h16_ = source29_.cf9
                d_2974___mcc_h17_ = source29_.cf10
                d_2975___mcc_h18_ = source29_.cf11
                d_2976_cf11_ = d_2975___mcc_h18_
                d_2977_cf10_ = d_2974___mcc_h17_
                d_2978_cf9_ = d_2973___mcc_h16_
                d_2979_cf8_ = d_2972___mcc_h15_
                d_2980_v132_: _dafny.Map
                d_2980_v132_ = _dafny.Map({((0) - (default__.fm1(d_2917_v90_, globalState)) if False else d_2954_cf15_): (len(d_2958_v122_)) != (d_2954_cf15_)})
                d_2980_v132_ = (d_2980_v132_).set(d_2954_cf15_, d_2917_v90_)
                d_2981_v133_: C4
                nw568_ = C4()
                nw568_.ctor__()
                d_2981_v133_ = nw568_
                (globalState).f6 = d_2954_cf15_
                (globalState).f14 = d_2959_v123_
            elif True:
                d_2982___mcc_h19_ = source29_.cf6
                d_2983_cf6_ = d_2982___mcc_h19_
                rhs415_ = (d_2916_v89_) == (d_2916_v89_)
                rhs416_ = d_2916_v89_
                r1 = rhs415_
                d_2917_v90_ = rhs416_
                d_2984_v134_: C1
                nw569_ = C1()
                nw569_.ctor__((self).f15)
                d_2984_v134_ = nw569_
                index502_ = default__.safeIndex(321, (d_2959_v123_).length(0))
                (d_2959_v123_)[index502_] = (d_2984_v134_).fm24(d_2953_cf16_, default__.fm39(d_2917_v90_, globalState), globalState)
                d_2985_v135_: _dafny.Array
                nw570_ = _dafny.Array(_dafny.Map({}), 8)
                d_2985_v135_ = nw570_
                d_2985_v135_ = d_2985_v135_
            d_2986_v138_: _dafny.Seq
            d_2986_v138_ = _dafny.SeqWithoutIsStrInference([d_2797_v0_])
            d_2987_v139_: _dafny.Map
            d_2987_v139_ = _dafny.Map({_dafny.Map({d_2954_cf15_: (d_2959_v123_)[default__.safeIndex(321, (d_2959_v123_).length(0))]}): (self).f15})
            d_2988_v140_: T1
            nw571_ = C3()
            nw571_.ctor__((d_2959_v123_)[default__.safeIndex(321, (d_2959_v123_).length(0))], (self).f15, d_2987_v139_)
            d_2988_v140_ = nw571_
            d_2989_v141_: _dafny.Set
            d_2989_v141_ = _dafny.Set({d_2988_v140_, d_2988_v140_})
            d_2990_v142_: _dafny.Map
            d_2990_v142_ = _dafny.Map({False: _dafny.Set({d_2954_cf15_})})
            d_2991_v143_: _dafny.Map
            def iife300_():
                def iife302_():
                    coll94_ = _dafny.Map()
                    compr_94_: int
                    for compr_94_ in ((d_2986_v138_)[default__.safeIndex(len(d_2989_v141_), len(d_2986_v138_))]).Elements:
                        d_2992_v137_: int = compr_94_
                        if (d_2992_v137_) in ((d_2986_v138_)[default__.safeIndex(len(d_2989_v141_), len(d_2986_v138_))]):
                            coll94_[(d_2992_v137_) + (d_2954_cf15_)] = (d_2959_v123_)[default__.safeIndex(321, (d_2959_v123_).length(0))]
                    return _dafny.Map(coll94_)
                coll92_ = _dafny.Map()
                def iife301_():
                    coll93_ = _dafny.Map()
                    compr_93_: int
                    for compr_93_ in ((d_2986_v138_)[default__.safeIndex(len(d_2989_v141_), len(d_2986_v138_))]).Elements:
                        d_2992_v137_: int = compr_93_
                        if (d_2992_v137_) in ((d_2986_v138_)[default__.safeIndex(len(d_2989_v141_), len(d_2986_v138_))]):
                            coll93_[(d_2992_v137_) + (d_2954_cf15_)] = (d_2959_v123_)[default__.safeIndex(321, (d_2959_v123_).length(0))]
                    return _dafny.Map(coll93_)
                compr_92_: int
                for compr_92_ in (iife301_()
                ).keys.Elements:
                    d_2993_v136_: int = compr_92_
                    if (d_2993_v136_) in (iife302_()
                    ):
                        coll92_[(d_2993_v136_) * ((self).f15)] = len(d_2799_v2_)
                return _dafny.Map(coll92_)
            d_2991_v143_ = _dafny.Map({(len(iife300_()
            )) <= (len(d_2918_v91_)): (d_2918_v91_).issubset(((d_2990_v142_)[d_2917_v90_] if (d_2917_v90_) in (d_2990_v142_) else _dafny.Set({len(d_2799_v2_)})))})
            index503_ = default__.safeIndex(321, (d_2959_v123_).length(0))
            (d_2959_v123_)[index503_] = ((d_2991_v143_)[d_2800_v3_] if (d_2800_v3_) in (d_2991_v143_) else d_2800_v3_)
        elif source28_.is_DC10:
            d_2994___mcc_h12_ = source28_.cf13
            d_2995_cf13_ = d_2994___mcc_h12_
            (globalState).f6 = len(((d_2797_v0_) + ((d_2797_v0_).set(default__.safeIndex((self).f15, len(d_2797_v0_)), (self).f15))) + ((d_2797_v0_) + (_dafny.SeqWithoutIsStrInference([(self).f15 for d_2996_i11_ in range(default__.abs(-720))]))))
            d_2918_v91_ = d_2918_v91_
            if d_2917_v90_:
                d_2997_v144_: bool
                d_2998_v145_: bool
                out28_: bool
                out29_: bool
                out28_, out29_ = (self).m4(True, globalState)
                d_2997_v144_ = out28_
                d_2998_v145_ = out29_
                d_2997_v144_ = (not(d_2916_v89_) if not(d_2997_v144_) else d_2997_v144_)
                (globalState).f6 = (415) - ((self).f15)
                d_2999_v146_: _dafny.Map
                d_2999_v146_ = _dafny.Map({(self).f15: d_2917_v90_})
                d_3000_v147_: _dafny.Map
                d_3000_v147_ = _dafny.Map({d_2999_v146_: (self).f15})
                d_3001_v148_: C2
                nw572_ = C2()
                nw572_.ctor__(71, default__.safeModuloInt(len(d_2999_v146_), (self).f15), d_3000_v147_)
                d_3001_v148_ = nw572_
                d_3002_v149_: _dafny.Array
                nw573_ = _dafny.Array(False, 5)
                d_3002_v149_ = nw573_
                d_3003_v150_: _dafny.Map
                d_3003_v150_ = _dafny.Map({d_3002_v149_: d_3000_v147_})
                d_3004_v151_: _dafny.Map
                d_3004_v151_ = ((d_3003_v150_)[d_3002_v149_] if (d_3002_v149_) in (d_3003_v150_) else d_3000_v147_)
                d_3004_v151_ = d_3004_v151_
            elif True:
                d_3005_v152_: _dafny.Array
                def lambda145_(d_3006_v89_):
                    def lambda146_(d_3007_i12_):
                        return not(d_3006_v89_)

                    return lambda146_

                init84_ = lambda145_(d_2916_v89_)
                nw574_ = _dafny.Array(None, 19)
                for i0_84_ in range(nw574_.length(0)):
                    nw574_[i0_84_] = init84_(i0_84_)
                d_3005_v152_ = nw574_
                index504_ = default__.safeIndex(446, (d_3005_v152_).length(0))
                (d_3005_v152_)[index504_] = d_2916_v89_
                d_3008_v153_: _dafny.Array
                nw575_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_3008_v153_ = nw575_
                d_3009_v154_: D10
                d_3009_v154_ = D10_DC26(d_3008_v153_)
                d_3010_v155_: _dafny.Set
                d_3010_v155_ = _dafny.Set({d_3009_v154_, d_3009_v154_, D10_DC26(d_3008_v153_), d_3009_v154_})
                index505_ = default__.safeIndex(446, (d_3005_v152_).length(0))
                (d_3005_v152_)[index505_] = (d_3010_v155_).issubset(d_3010_v155_)
                d_3011_v156_: _dafny.Array
                nw576_ = _dafny.Array(int(0), 20)
                d_3011_v156_ = nw576_
                index506_ = default__.safeIndex(127, (d_3011_v156_).length(0))
                (d_3011_v156_)[index506_] = (self).f15
                index507_ = default__.safeIndex(127, (d_3011_v156_).length(0))
                (d_3011_v156_)[index507_] = (self).f15
                index508_ = default__.safeIndex(127, (d_3011_v156_).length(0))
                (d_3011_v156_)[index508_] = (self).f15
                r2 = default__.safeDivisionInt((d_3011_v156_)[default__.safeIndex(127, (d_3011_v156_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_2919_v92_ for d_3012_i13_ in range(default__.abs(-962))])))
                (globalState).f6 = default__.safeModuloInt((d_3011_v156_)[default__.safeIndex(127, (d_3011_v156_).length(0))], len(d_2799_v2_))
            (globalState).f6 = ((self).f15) - ((self).f15)
        elif True:
            d_3013___mcc_h13_ = source28_.cf18
            d_3014_cf18_ = d_3013___mcc_h13_
            if d_2800_v3_:
                (globalState).f1 = d_2800_v3_
                d_3015_v157_: _dafny.Set
                d_3015_v157_ = _dafny.Set({d_2800_v3_})
                d_3016_v158_: _dafny.Array
                nw577_ = _dafny.Array(int(0), 2)
                d_3016_v158_ = nw577_
                d_3017_v159_: _dafny.Set
                d_3017_v159_ = _dafny.Set({d_3016_v158_, d_3016_v158_, d_3016_v158_, d_3016_v158_, d_3016_v158_})
                rhs417_ = (self).f15
                rhs418_ = (d_3015_v157_) - ((_dafny.Set({d_2800_v3_, not(d_2916_v89_), False})) - (_dafny.Set({d_2800_v3_, d_2916_v89_})))
                rhs419_ = ((d_3017_v159_).intersection(d_3017_v159_)).ispropersubset(d_3017_v159_)
                lhs338_ = globalState
                lhs339_ = globalState
                lhs338_.f6 = rhs417_
                d_3015_v157_ = rhs418_
                lhs339_.f1 = rhs419_
                d_3018_v160_: _dafny.Seq
                d_3018_v160_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dd")), d_2799_v2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvbevapf")), d_2799_v2_])
                d_3019_v161_: _dafny.Seq
                d_3019_v161_ = _dafny.SeqWithoutIsStrInference([(d_3018_v160_)[default__.safeIndex((self).f15, len(d_3018_v160_))], d_2799_v2_, default__.fm21(globalState), d_2799_v2_])
                d_3020_v162_: D8
                d_3020_v162_ = D8_DC21((d_3018_v160_)[default__.safeIndex((self).f15, len(d_3018_v160_))], d_3015_v157_)
                d_3020_v162_ = d_3020_v162_
                index509_ = default__.safeIndex(984, (d_3016_v158_).length(0))
                (d_3016_v158_)[index509_] = (self).f15
                d_3021_v163_: _dafny.Map
                d_3021_v163_ = _dafny.Map({(self).f15: not(d_2917_v90_)})
                index510_ = default__.safeIndex(984, (d_3016_v158_).length(0))
                (d_3016_v158_)[index510_] = ((self).f15) * (len(d_3021_v163_))
                d_3022_v164_: C7
                nw578_ = C7()
                nw578_.ctor__()
                d_3022_v164_ = nw578_
            elif True:
                d_3023_v165_: _dafny.Set
                d_3023_v165_ = _dafny.Set({d_2919_v92_})
                d_3023_v165_ = d_3023_v165_
                d_3024_v166_: D1
                d_3024_v166_ = D1_DC4(d_2919_v92_)
                d_3025_v167_: _dafny.Array
                nw579_ = _dafny.Array(int(0), 22)
                d_3025_v167_ = nw579_
                rhs420_ = d_3025_v167_
                rhs421_ = d_3024_v166_
                rhs422_ = (p0)[default__.safeIndex((d_2797_v0_)[default__.safeIndex(-491, len(d_2797_v0_))], len(p0))]
                rhs423_ = not(((self).f15) > (727))
                lhs340_ = globalState
                lhs341_ = globalState
                lhs340_.f9 = rhs420_
                d_3024_v166_ = rhs421_
                lhs341_.f1 = rhs422_
                d_2917_v90_ = rhs423_
                d_3026_v168_: _dafny.Set
                d_3026_v168_ = _dafny.Set({d_2800_v3_, d_2800_v3_})
                d_2800_v3_ = (d_3026_v168_) != ((d_3026_v168_) | (d_3026_v168_))
                d_3027_v169_: _dafny.Map
                d_3027_v169_ = _dafny.Map({d_2800_v3_: d_2916_v89_})
                d_3028_v170_: _dafny.MultiSet
                d_3028_v170_ = _dafny.MultiSet([default__.fm1(d_2916_v89_, globalState), (len(((D1_DC3(d_3027_v169_)).cf3).set(d_2916_v89_, False))) - (255), (self).f15, (self).f15])
                d_3028_v170_ = d_3028_v170_
                r2 = (len(p0)) - ((0) - ((self).f15))
            d_2919_v92_ = d_2919_v92_
            d_3029_v171_: _dafny.Seq
            d_3029_v171_ = _dafny.SeqWithoutIsStrInference([d_2916_v89_, d_2916_v89_, d_2916_v89_])
            d_3030_v172_: _dafny.MultiSet
            d_3030_v172_ = _dafny.MultiSet([940])
            d_3031_v173_: D11
            d_3031_v173_ = D11_DC31(_dafny.SeqWithoutIsStrInference([d_2916_v89_]), (0) - ((self).f15), (d_3030_v172_).cardinality, (0) - ((self).f15), (self).f15)
            d_3029_v171_ = (d_3031_v173_).cf54
            (globalState).f6 = (self).f15
        r0 = d_2800_v3_
        r1 = ((self).f15) < ((self).f15)
        d_3032_v174_: _dafny.Map
        d_3032_v174_ = _dafny.Map({d_2797_v0_: len(d_2798_v1_)})
        d_3033_v175_: _dafny.Map
        d_3033_v175_ = _dafny.Map({d_2800_v3_: (0) - (len(d_3032_v174_))})
        r2 = (0) - (((self).f15 if not (d_2917_v90_) or (d_2800_v3_) else len(d_3033_v175_)))
        return r0, r1, r2

    def m4(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        hi18_ = (self).f15
        for d_3034_i0_ in range((self).f15, hi18_):
            if (default__.fm2(globalState) if p0 else p0):
                (globalState).f6 = d_3034_i0_
                d_3035_v0_: _dafny.Array
                def lambda147_(d_3036_i1_):
                    return True

                init85_ = lambda147_
                nw580_ = _dafny.Array(None, 19)
                for i0_85_ in range(nw580_.length(0)):
                    nw580_[i0_85_] = init85_(i0_85_)
                d_3035_v0_ = nw580_
                index511_ = default__.safeIndex(557, (d_3035_v0_).length(0))
                (d_3035_v0_)[index511_] = not(p0)
                index512_ = default__.safeIndex(557, (d_3035_v0_).length(0))
                (d_3035_v0_)[index512_] = False
                d_3037_v1_: _dafny.Seq
                d_3037_v1_ = _dafny.SeqWithoutIsStrInference([p0, (d_3035_v0_)[default__.safeIndex(557, (d_3035_v0_).length(0))], (d_3035_v0_)[default__.safeIndex(557, (d_3035_v0_).length(0))], not(p0), (d_3035_v0_)[default__.safeIndex(557, (d_3035_v0_).length(0))]])
                d_3037_v1_ = d_3037_v1_
                (globalState).f1 = (d_3035_v0_)[default__.safeIndex(557, (d_3035_v0_).length(0))]
                r1 = p0
            elif True:
                d_3038_v2_: _dafny.Map
                d_3038_v2_ = _dafny.Map({d_3034_i0_: p0})
                d_3039_v3_: _dafny.Map
                d_3039_v3_ = _dafny.Map({(self).f15: d_3038_v2_})
                d_3040_v5_: _dafny.Map
                def iife303_():
                    coll95_ = _dafny.Map()
                    compr_95_: int
                    for compr_95_ in _dafny.IntegerRange(728, 190):
                        d_3041_v4_: int = compr_95_
                        if ((728) <= (d_3041_v4_)) and ((d_3041_v4_) < (190)):
                            coll95_[default__.safeModuloInt(d_3041_v4_, (self).f15)] = p0
                    return _dafny.Map(coll95_)
                d_3040_v5_ = _dafny.Map({((d_3039_v3_)[(0) - (d_3034_i0_)] if ((0) - (d_3034_i0_)) in (d_3039_v3_) else iife303_()
                ): d_3034_i0_})
                d_3042_v6_: C3
                nw581_ = C3()
                nw581_.ctor__(not (p0) or (p0), 233, d_3040_v5_)
                d_3042_v6_ = nw581_
                d_3043_v7_: _dafny.Seq
                d_3043_v7_ = _dafny.SeqWithoutIsStrInference([p0])
                d_3044_v8_: _dafny.Set
                d_3044_v8_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference([False])) + (d_3043_v7_), _dafny.SeqWithoutIsStrInference([(d_3042_v6_).f18, (d_3042_v6_).f18, not(p0), p0]), d_3043_v7_, _dafny.SeqWithoutIsStrInference([False]), (_dafny.SeqWithoutIsStrInference([p0])) + (d_3043_v7_)})
                d_3044_v8_ = (d_3044_v8_) - (d_3044_v8_)
                d_3045_v9_: _dafny.MultiSet
                d_3045_v9_ = _dafny.MultiSet([p0])
                d_3046_v10_: _dafny.Array
                nw582_ = _dafny.Array(None, 9)
                nw582_[int(0)] = 512
                nw582_[int(1)] = (d_3042_v6_).f19
                nw582_[int(2)] = ((self).f15) + (d_3034_i0_)
                nw582_[int(3)] = 729
                nw582_[int(4)] = default__.safeModuloInt(d_3034_i0_, d_3034_i0_)
                nw582_[int(5)] = 154
                nw582_[int(6)] = d_3034_i0_
                nw582_[int(7)] = len(default__.fm12(d_3045_v9_, 759, globalState))
                nw582_[int(8)] = (d_3042_v6_).f19
                d_3046_v10_ = nw582_
                d_3047_v11_: _dafny.Seq
                d_3047_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tviyhhlfk"))
                index513_ = default__.safeIndex(13, (d_3046_v10_).length(0))
                (d_3046_v10_)[index513_] = len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mpbabn")), d_3047_v11_]))
                index514_ = default__.safeIndex(13, (d_3046_v10_).length(0))
                (d_3046_v10_)[index514_] = default__.safeModuloInt((self).f15, (self).fm5(not(p0), globalState))
                d_3048_v12_: _dafny.Map
                d_3048_v12_ = _dafny.Map({(d_3046_v10_)[default__.safeIndex(13, (d_3046_v10_).length(0))]: (self).f15})
                d_3049_v13_: _dafny.Seq
                d_3049_v13_ = _dafny.SeqWithoutIsStrInference([((d_3048_v12_)[d_3034_i0_] if (d_3034_i0_) in (d_3048_v12_) else (d_3045_v9_).cardinality)])
                d_3050_v14_: _dafny.Map
                d_3050_v14_ = _dafny.Map({default__.fm21(globalState): d_3049_v13_})
                d_3050_v14_ = (d_3050_v14_).set(default__.fm21(globalState), d_3049_v13_)
                (globalState).f6 = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_3051_i2_ in range(default__.abs(981))]))
            d_3052_v15_: _dafny.Seq
            d_3052_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpvilna"))
            d_3052_v15_ = d_3052_v15_
            d_3053_v16_: _dafny.Array
            nw583_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_3053_v16_ = nw583_
            d_3054_v17_: _dafny.Array
            def lambda148_(d_3055_p0_):
                def lambda149_(d_3056_i3_):
                    return _dafny.SeqWithoutIsStrInference([d_3055_p0_, d_3055_p0_, d_3055_p0_])

                return lambda149_

            init86_ = lambda148_(p0)
            nw584_ = _dafny.Array(None, 5)
            for i0_86_ in range(nw584_.length(0)):
                nw584_[i0_86_] = init86_(i0_86_)
            d_3054_v17_ = nw584_
            index515_ = default__.safeIndex(46, (d_3053_v16_).length(0))
            (d_3053_v16_)[index515_] = d_3054_v17_
            index516_ = default__.safeIndex(46, (d_3053_v16_).length(0))
            (d_3053_v16_)[index516_] = d_3054_v17_
            (globalState).f1 = p0
        (globalState).f6 = (self).f15
        d_3057_v18_: _dafny.Seq
        d_3057_v18_ = _dafny.SeqWithoutIsStrInference([(self).f15])
        d_3058_v19_: _dafny.Seq
        d_3058_v19_ = _dafny.SeqWithoutIsStrInference([d_3057_v18_])
        d_3059_v20_: _dafny.Map
        d_3059_v20_ = _dafny.Map({-546: d_3057_v18_})
        d_3060_v21_: _dafny.Array
        nw585_ = _dafny.Array(None, 18)
        nw585_[int(0)] = d_3057_v18_
        nw585_[int(1)] = d_3057_v18_
        nw585_[int(2)] = _dafny.SeqWithoutIsStrInference([(self).f15, (self).f15, (self).f15, 994])
        nw585_[int(3)] = d_3057_v18_
        nw585_[int(4)] = d_3057_v18_
        nw585_[int(5)] = (d_3058_v19_)[default__.safeIndex(-473, len(d_3058_v19_))]
        nw585_[int(6)] = _dafny.SeqWithoutIsStrInference([670 for d_3061_i4_ in range(default__.abs(340))])
        nw585_[int(7)] = d_3057_v18_
        nw585_[int(8)] = d_3057_v18_
        nw585_[int(9)] = d_3057_v18_
        nw585_[int(10)] = d_3057_v18_
        nw585_[int(11)] = d_3057_v18_
        nw585_[int(12)] = ((d_3059_v20_)[(self).f15] if ((self).f15) in (d_3059_v20_) else d_3057_v18_)
        nw585_[int(13)] = d_3057_v18_
        nw585_[int(14)] = _dafny.SeqWithoutIsStrInference([(self).f15, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "doh")))])
        nw585_[int(15)] = d_3057_v18_
        nw585_[int(16)] = d_3057_v18_
        nw585_[int(17)] = d_3057_v18_
        d_3060_v21_ = nw585_
        d_3062_v22_: _dafny.Seq
        d_3062_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "osjud"))
        d_3063_v23_: _dafny.Map
        d_3063_v23_ = _dafny.Map({d_3060_v21_: (d_3062_v22_) != (d_3062_v22_)})
        d_3064_v24_: _dafny.Map
        d_3064_v24_ = _dafny.Map({p0: p0})
        d_3065_v25_: _dafny.MultiSet
        d_3065_v25_ = _dafny.MultiSet([d_3064_v24_, d_3064_v24_])
        d_3063_v23_ = (d_3063_v23_).set(d_3060_v21_, ((self).f15) == ((d_3065_v25_).cardinality))
        d_3066_v26_: _dafny.Map
        d_3066_v26_ = _dafny.Map({(self).f15: 961})
        d_3067_v27_: C1
        nw586_ = C1()
        nw586_.ctor__(774)
        d_3067_v27_ = nw586_
        d_3068_v28_: _dafny.Seq
        d_3068_v28_ = _dafny.SeqWithoutIsStrInference([d_3067_v27_])
        d_3069_v29_: str
        d_3069_v29_ = _dafny.CodePoint('j')
        d_3070_v30_: D16
        d_3070_v30_ = D16_DC45(len(d_3068_v28_), d_3069_v29_, (self).f15, p0)
        d_3066_v26_ = (d_3066_v26_).set((self).f15, default__.fm1((d_3070_v30_).cf81, globalState))
        d_3069_v29_ = d_3069_v29_
        d_3071_v31_: C5
        nw587_ = C5()
        nw587_.ctor__((d_3067_v27_).f24)
        d_3071_v31_ = nw587_
        d_3071_v31_ = d_3071_v31_
        r0 = True
        r1 = not((d_3069_v29_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jayoxhlrs"))))
        return r0, r1

    @property
    def f15(self):
        return self._f15
