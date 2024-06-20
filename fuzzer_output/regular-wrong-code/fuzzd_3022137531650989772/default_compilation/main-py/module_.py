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
        if False:
            return 34
        elif True:
            return default__.safeDivisionInt(419, len(_dafny.SeqWithoutIsStrInference([True, not(False), False, False, True])))

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(-920, 636):
                d_0_v0_: int = compr_0_
                if ((-920) <= (d_0_v0_)) and ((d_0_v0_) < (636)):
                    coll0_[(d_0_v0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ff"))))] = 96
            return _dafny.Map(coll0_)
        return (_dafny.Map({447: len(_dafny.Map({True: False}))})) | ((_dafny.Map({881: 850})) | (iife0_()
        ))

    @staticmethod
    def fm3(p0, globalState):
        source0_ = D3_DC11()
        if source0_.is_DC11:
            if True:
                return _dafny.CodePoint('h')
            elif True:
                return _dafny.CodePoint('v')
        elif True:
            d_1___mcc_h0_ = source0_.cf18
            d_2_cf18_ = d_1___mcc_h0_
            return _dafny.CodePoint('t')

    @staticmethod
    def fm4(p0, p1, globalState):
        return (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_3_i0_ in range(default__.abs(-89))])), -546)) < (default__.safeDivisionInt(-556, 756))

    @staticmethod
    def fm5(globalState):
        return _dafny.Map({default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_4_i0_ in range(default__.abs(733))])), (_dafny.MultiSet([not(False)])).cardinality): False})

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([(D1_DC4(len(_dafny.Map({False: (_dafny.MultiSet([False])).cardinality})), False, _dafny.Map({-523: _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylexru")))])}))).cf9, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True]))]))])).Elements:
                d_5_v0_: int = compr_1_
                if (d_5_v0_) in (_dafny.SeqWithoutIsStrInference([(D1_DC4(len(_dafny.Map({False: (_dafny.MultiSet([False])).cardinality})), False, _dafny.Map({-523: _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylexru")))])}))).cf9, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True]))]))])):
                    coll1_[(d_5_v0_) + (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djl"))), len(_dafny.SeqWithoutIsStrInference([False, not(True), False]))})))] = 529
            return _dafny.Map(coll1_)
        return _dafny.MultiSet([True, (-765) < (len(iife1_()
        )), (True) or (not(True))])

    @staticmethod
    def fm7(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([not((_dafny.MultiSet([True])).ispropersubset(_dafny.MultiSet([False, True])))])

    @staticmethod
    def fm8(p0, p1, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(-409, 468):
                d_6_v0_: int = compr_2_
                if ((-409) <= (d_6_v0_)) and ((d_6_v0_) < (468)):
                    coll2_[(d_6_v0_) * ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uicbdh")))))] = (_dafny.MultiSet([687, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_7_i0_ in range(default__.abs(550))]))])).cardinality
            return _dafny.Map(coll2_)
        return D1_DC3((iife2_()
) | (_dafny.Map({len(_dafny.Set({True})): -962})), (0) - (len((_dafny.SeqWithoutIsStrInference([-196, 427])) + (_dafny.SeqWithoutIsStrInference([223])))), (D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlkawrtbe"))), False, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_8_i1_ in range(default__.abs(124))]), _dafny.Map({675: _dafny.Map({-276: False})})) if False else D0_DC0(205, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xhtbrf")), _dafny.Map({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bq"))]))])).cardinality: _dafny.Map({-318: True})}))))

    @staticmethod
    def fm9(globalState):
        def iife3_():
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(619, 917):
                    d_12_v1_: int = compr_4_
                    if ((619) <= (d_12_v1_)) and ((d_12_v1_) < (917)):
                        def iife5_():
                            coll5_ = _dafny.Map()
                            compr_5_: int
                            for compr_5_ in _dafny.IntegerRange(712, 236):
                                d_13_v2_: int = compr_5_
                                if ((712) <= (d_13_v2_)) and ((d_13_v2_) < (236)):
                                    coll5_[default__.safeDivisionInt(d_13_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qokpwk"))))] = False
                            return _dafny.Map(coll5_)
                        coll4_[(d_12_v1_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fcscl"))])))] = iife5_()

                return _dafny.Map(coll4_)
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(364, 851):
                d_10_v0_: int = compr_3_
                if ((364) <= (d_10_v0_)) and ((d_10_v0_) < (851)):
                    coll3_[(d_10_v0_) - ((0) - (len(_dafny.SeqWithoutIsStrInference([D0_DC0(-898, True, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_11_i1_ in range(default__.abs(569))]), iife4_()
)]))))] = False
            return _dafny.Map(coll3_)
        def iife6_():
            coll6_ = _dafny.Set()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(-634, 9):
                d_14_v3_: int = compr_6_
                if ((-634) <= (d_14_v3_)) and ((d_14_v3_) < (9)):
                    coll6_ = coll6_.union(_dafny.Set([default__.safeModuloInt(d_14_v3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bk"))))]))
            return _dafny.Set(coll6_)
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(620, 477):
                d_15_v4_: int = compr_7_
                if ((620) <= (d_15_v4_)) and ((d_15_v4_) < (477)):
                    coll7_[default__.safeModuloInt(d_15_v4_, 77)] = 486
            return _dafny.Map(coll7_)
        return (_dafny.SeqWithoutIsStrInference([D1_DC3(_dafny.Map({807: len(_dafny.Map({False: True}))}), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdqdlcew"))), 266, 135])).cardinality, D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))), False, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_9_i0_ in range(default__.abs(839))]), _dafny.Map({717: iife3_()
})))])) + (_dafny.SeqWithoutIsStrInference([D1_DC3(_dafny.Map({-962: len((D5_DC17(iife6_()
)).cf29)}), 662, D0_DC0(164, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "newt")), _dafny.Map({len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxly")))): False})): _dafny.Map({447: True})}))), D1_DC3(iife7_()
, -215, D0_DC0(651, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhxxrjsnl")), _dafny.Map({295: _dafny.Map({-593: True})}))), D1_DC3(_dafny.Map({414: (_dafny.MultiSet([False])).cardinality}), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_16_i2_ in range(default__.abs(356))])), D0_DC0(len(_dafny.Map({not(True): 867})), True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouh")), _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))): _dafny.Map({207: False})})))]))

    @staticmethod
    def fm10(p0, globalState):
        return (_dafny.Map({_dafny.MultiSet([True, False, True, False]): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_17_i0_ in range(default__.abs(603))]))})) | ((_dafny.Map({_dafny.MultiSet([False]): -742})) | ((D6_DC20(_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True), True])): 916}))).cf33))

    @staticmethod
    def fm11(p0, p1, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(True)])): True})).keys.Elements:
                d_18_v0_: int = compr_8_
                if (d_18_v0_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(True)])): True})):
                    coll8_[default__.safeDivisionInt(d_18_v0_, len(_dafny.Map({623: False})))] = True
            return _dafny.Map(coll8_)
        return (_dafny.Map({758: iife8_()
        })) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bre"))): _dafny.Map({-798: True})})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(True)])): _dafny.Map({-689: False})})))

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcyqkio"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egwe"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sewnuy"))))

    @staticmethod
    def fm13(globalState):
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(357, 568):
                d_19_v0_: int = compr_9_
                if ((357) <= (d_19_v0_)) and ((d_19_v0_) < (568)):
                    coll9_[default__.safeModuloInt(d_19_v0_, 793)] = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_20_i0_ in range(default__.abs(13))])))
            return _dafny.Map(coll9_)
        def iife10_():
            coll10_ = _dafny.Set()
            compr_10_: str
            for compr_10_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_22_i2_ in range(default__.abs(581))])).Elements:
                d_23_v1_: str = compr_10_
                if (d_23_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_22_i2_ in range(default__.abs(581))])):
                    coll10_ = coll10_.union(_dafny.Set([d_23_v1_]))
            return _dafny.Set(coll10_)
        return D1_DC4(795, not(not((len(iife9_()
)) == (907))), (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_21_i1_ in range(default__.abs(-463))])): _dafny.MultiSet([len(_dafny.Map({True: len(_dafny.Map({True: 343}))})), -382, (0) - (len(iife10_()
))])})) | (_dafny.Map({514: _dafny.MultiSet([726])})))

    @staticmethod
    def fm14(globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([-785]), _dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxmprhotv"))))]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, True, False})), 378, (0) - (-951), (0) - ((_dafny.MultiSet([False])).cardinality)]))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-950])) for d_24_i0_ in range(default__.abs(-648))])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([-132, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_25_i2_ in range(default__.abs(914))]))), -96]) for d_26_i1_ in range(default__.abs(242))])))

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        return _dafny.Map({_dafny.CodePoint('d'): (not(False)) and (False)})

    @staticmethod
    def fm16(p0, p1, p2, globalState):
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(671, 265):
                d_27_v0_: int = compr_11_
                if ((671) <= (d_27_v0_)) and ((d_27_v0_) < (265)):
                    def iife12_():
                        coll12_ = _dafny.Map()
                        compr_12_: int
                        for compr_12_ in _dafny.IntegerRange(858, 123):
                            d_28_v1_: int = compr_12_
                            if ((858) <= (d_28_v1_)) and ((d_28_v1_) < (123)):
                                coll12_[default__.safeDivisionInt(d_28_v1_, 410)] = not(False)
                        return _dafny.Map(coll12_)
                    coll11_[default__.safeModuloInt(d_27_v0_, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_29_i0_ in range(default__.abs(-506))]))))] = iife12_()

            return _dafny.Map(coll11_)
        return D0_DC0(-10, (683) > (len(_dafny.Set({316}))), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")) if True else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpsbupx"))), iife11_()
)

    @staticmethod
    def fm17(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([822 for d_30_i0_ in range(default__.abs(642))])) + ((_dafny.SeqWithoutIsStrInference([219, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agbh"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvgfxoaqc"))), len(_dafny.Set({-211, 157, 330})), 31])) + (_dafny.SeqWithoutIsStrInference([801, 912])))

    @staticmethod
    def fm18(p0, p1, p2, p3, globalState):
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvxhidyxj"))})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_31_i0_ in range(default__.abs(146))])}))

    @staticmethod
    def fm19(p0, globalState):
        return ((_dafny.Map({True: _dafny.Map({False: len(_dafny.Map({False: False}))})})) | (_dafny.Map({True: _dafny.Map({True: 878})}))) | ((_dafny.Map({True: _dafny.Map({(D6_DC21(True, _dafny.CodePoint('g'), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjujka"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxlnkch"))))).cf34: 317})})) | (_dafny.Map({not(False): _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpd")))]))})})))

    @staticmethod
    def fm20(p0, globalState):
        return _dafny.MultiSet([default__.safeModuloInt(931, 846), -596])

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Set = _dafny.Set({})
        d_32_v0_: bool
        d_32_v0_ = True
        d_33_v1_: str
        d_33_v1_ = _dafny.CodePoint('c')
        d_34_v2_: _dafny.Set
        d_34_v2_ = _dafny.Set({d_33_v1_})
        d_32_v0_ = (p1) > (default__.safeDivisionInt(-677, len(d_34_v2_)))
        (globalState).f6 = p1
        r0 = ((-453) + (p1)) + ((0) - ((p1) * ((0) - (p1))))
        d_35_v3_: _dafny.Seq
        d_35_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhpactow"))
        d_36_v4_: D0
        d_36_v4_ = D0_DC0(p1, p3, d_35_v3_, default__.fm11(p1, d_32_v0_, globalState))
        pat_let_tv0_ = d_33_v1_
        pat_let_tv1_ = p3
        pat_let_tv2_ = d_36_v4_
        pat_let_tv3_ = p1
        pat_let_tv4_ = d_33_v1_
        pat_let_tv5_ = p1
        pat_let_tv6_ = globalState
        def lambda0_(source2_):
            d_37___mcc_h9_ = source2_.cf0
            d_38___mcc_h10_ = source2_.cf1
            d_39___mcc_h11_ = source2_.cf2
            d_40___mcc_h12_ = source2_.cf3
            d_41_cf3_ = d_40___mcc_h12_
            d_42_cf2_ = d_39___mcc_h11_
            d_43_cf1_ = d_38___mcc_h10_
            d_44_cf0_ = d_37___mcc_h9_
            return D1_DC2(pat_let_tv0_)

        def iife13_(_pat_let0_0):
            def iife14_(d_45_dt__update__tmp_h0_):
                def iife15_(_pat_let1_0):
                    def iife16_(d_46_dt__update_hcf1_h0_):
                        def iife17_(_pat_let2_0):
                            def iife18_(d_47_dt__update_hcf2_h0_):
                                return D0_DC0((d_45_dt__update__tmp_h0_).cf0, d_46_dt__update_hcf1_h0_, d_47_dt__update_hcf2_h0_, (d_45_dt__update__tmp_h0_).cf3)
                            return iife18_(_pat_let2_0)
                        return iife17_(default__.fm12(pat_let_tv2_, pat_let_tv3_, pat_let_tv4_, pat_let_tv5_, pat_let_tv6_))
                    return iife16_(_pat_let1_0)
                return iife15_(pat_let_tv1_)
            return iife14_(_pat_let0_0)
        source1_ = lambda0_(iife13_(d_36_v4_))
        if source1_.is_DC2:
            d_48___mcc_h0_ = source1_.cf5
            d_49_cf5_ = d_48___mcc_h0_
            d_50_v5_: _dafny.Array
            nw0_ = _dafny.Array(False, 9)
            d_50_v5_ = nw0_
            rhs0_ = p2
            rhs1_ = d_32_v0_
            rhs2_ = (d_35_v3_) + ((d_35_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))))
            rhs3_ = d_50_v5_
            rhs4_ = d_35_v3_
            d_32_v0_ = rhs0_
            d_32_v0_ = rhs1_
            d_35_v3_ = rhs2_
            d_50_v5_ = rhs3_
            d_35_v3_ = rhs4_
            d_51_v6_: _dafny.Map
            d_51_v6_ = _dafny.Map({default__.safeDivisionInt(default__.fm0(p1, _dafny.Map({len(d_35_v3_): len((d_35_v3_).set(default__.safeIndex(p1, len(d_35_v3_)), d_49_cf5_))}), globalState), p1): (p1) - (p1)})
            d_52_v7_: _dafny.Set
            d_52_v7_ = _dafny.Set({d_35_v3_, d_35_v3_})
            d_51_v6_ = (d_51_v6_).set(p1, len(d_52_v7_))
            index0_ = default__.safeIndex(124, (d_50_v5_).length(0))
            (d_50_v5_)[index0_] = True
            index1_ = default__.safeIndex(124, (d_50_v5_).length(0))
            rhs5_ = p1
            rhs6_ = d_32_v0_
            rhs7_ = p1
            lhs0_ = globalState
            lhs1_ = d_50_v5_
            lhs2_ = default__.safeIndex(124, (d_50_v5_).length(0))
            lhs3_ = globalState
            lhs0_.f11 = rhs5_
            lhs1_[lhs2_] = rhs6_
            lhs3_.f16 = rhs7_
            d_53_v8_: _dafny.Array
            def lambda1_(d_54_p1_):
                def lambda2_(d_55_i0_):
                    return (d_55_i0_) + (d_54_p1_)

                return lambda2_

            init0_ = lambda1_(p1)
            nw1_ = _dafny.Array(None, 26)
            for i0_0_ in range(nw1_.length(0)):
                nw1_[i0_0_] = init0_(i0_0_)
            d_53_v8_ = nw1_
            index2_ = default__.safeIndex(544, (d_53_v8_).length(0))
            (d_53_v8_)[index2_] = p1
            index3_ = default__.safeIndex(544, (d_53_v8_).length(0))
            index4_ = default__.safeIndex(124, (d_50_v5_).length(0))
            rhs8_ = d_33_v1_
            rhs9_ = p1
            rhs10_ = d_32_v0_
            lhs4_ = d_53_v8_
            lhs5_ = default__.safeIndex(544, (d_53_v8_).length(0))
            lhs6_ = d_50_v5_
            lhs7_ = default__.safeIndex(124, (d_50_v5_).length(0))
            d_33_v1_ = rhs8_
            lhs4_[lhs5_] = rhs9_
            lhs6_[lhs7_] = rhs10_
        elif source1_.is_DC3:
            d_56___mcc_h1_ = source1_.cf6
            d_57___mcc_h2_ = source1_.cf7
            d_58___mcc_h3_ = source1_.cf8
            d_59_cf8_ = d_58___mcc_h3_
            d_60_cf7_ = d_57___mcc_h2_
            d_61_cf6_ = d_56___mcc_h1_
            d_62_v9_: C0
            nw2_ = C0()
            nw2_.ctor__(d_32_v0_)
            d_62_v9_ = nw2_
            d_63_v10_: _dafny.MultiSet
            d_63_v10_ = _dafny.MultiSet([d_62_v9_])
            d_64_v11_: _dafny.Map
            d_64_v11_ = _dafny.Map({d_62_v9_: ((d_63_v10_)[d_62_v9_] if (d_62_v9_) in (d_63_v10_) else p1)})
            d_65_v12_: _dafny.Set
            d_65_v12_ = _dafny.Set({d_64_v11_, (d_64_v11_) | (d_64_v11_), d_64_v11_, (d_64_v11_) | (d_64_v11_)})
            d_65_v12_ = d_65_v12_
            d_66_v14_: _dafny.Seq
            d_66_v14_ = _dafny.SeqWithoutIsStrInference([d_60_cf7_])
            d_67_v15_: _dafny.Seq
            d_67_v15_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_62_v9_).f17])])
            d_68_v16_: _dafny.Map
            d_68_v16_ = _dafny.Map({d_60_cf7_: p3})
            d_69_v17_: _dafny.Set
            d_69_v17_ = _dafny.Set({191})
            d_70_v18_: _dafny.Map
            d_70_v18_ = _dafny.Map({((d_68_v16_)[898] if (898) in (d_68_v16_) else p2): d_69_v17_})
            d_71_v19_: _dafny.Set
            d_71_v19_ = _dafny.Set({d_60_cf7_, (0) - (len(((d_70_v18_)[p2] if (p2) in (d_70_v18_) else d_69_v17_))), p1, d_60_cf7_})
            d_72_v20_: _dafny.Array
            nw3_ = _dafny.Array(None, 28)
            nw3_[int(0)] = (0) - (p1)
            nw3_[int(1)] = 763
            nw3_[int(2)] = (p1) - (default__.fm0(p1, d_61_cf6_, globalState))
            nw3_[int(3)] = p1
            nw3_[int(4)] = 947
            nw3_[int(5)] = p1
            def iife19_():
                coll13_ = _dafny.Map()
                compr_13_: int
                for compr_13_ in (d_61_cf6_).keys.Elements:
                    d_73_v13_: int = compr_13_
                    if (d_73_v13_) in (d_61_cf6_):
                        coll13_[(d_73_v13_) * (d_60_cf7_)] = p1
                return _dafny.Map(coll13_)
            nw3_[int(6)] = (d_60_cf7_) - (len(iife19_()
            ))
            nw3_[int(7)] = len(((d_35_v3_).set(default__.safeIndex(p1, len(d_35_v3_)), d_33_v1_)) + (d_35_v3_))
            nw3_[int(8)] = p1
            nw3_[int(9)] = ((d_63_v10_)[d_62_v9_] if (d_62_v9_) in (d_63_v10_) else d_60_cf7_)
            nw3_[int(10)] = p1
            nw3_[int(11)] = d_60_cf7_
            nw3_[int(12)] = d_60_cf7_
            nw3_[int(13)] = d_60_cf7_
            nw3_[int(14)] = d_60_cf7_
            nw3_[int(15)] = p1
            nw3_[int(16)] = d_60_cf7_
            nw3_[int(17)] = (d_66_v14_)[default__.safeIndex(p1, len(d_66_v14_))]
            nw3_[int(18)] = (d_66_v14_)[default__.safeIndex(d_60_cf7_, len(d_66_v14_))]
            nw3_[int(19)] = d_60_cf7_
            nw3_[int(20)] = len((d_67_v15_)[default__.safeIndex(p1, len(d_67_v15_))])
            nw3_[int(21)] = default__.safeModuloInt(d_60_cf7_, -884)
            nw3_[int(22)] = len(d_71_v19_)
            nw3_[int(23)] = p1
            nw3_[int(24)] = d_60_cf7_
            nw3_[int(25)] = d_60_cf7_
            nw3_[int(26)] = p1
            nw3_[int(27)] = len(d_35_v3_)
            d_72_v20_ = nw3_
            index5_ = default__.safeIndex(705, (d_72_v20_).length(0))
            (d_72_v20_)[index5_] = (183 if not(p3) else p1)
            index6_ = default__.safeIndex(705, (d_72_v20_).length(0))
            rhs11_ = default__.fm0(p1, d_61_cf6_, globalState)
            rhs12_ = 909
            rhs13_ = True
            rhs14_ = True
            rhs15_ = d_32_v0_
            lhs8_ = d_72_v20_
            lhs9_ = default__.safeIndex(705, (d_72_v20_).length(0))
            lhs10_ = globalState
            lhs8_[lhs9_] = rhs11_
            lhs10_.f16 = rhs12_
            d_32_v0_ = rhs13_
            d_32_v0_ = rhs14_
            d_32_v0_ = rhs15_
            d_74_v21_: _dafny.MultiSet
            d_74_v21_ = _dafny.MultiSet([p1])
            d_75_v22_: _dafny.Set
            d_75_v22_ = _dafny.Set({(d_62_v9_).f17, (d_62_v9_).f17})
            d_76_v23_: _dafny.Map
            d_76_v23_ = _dafny.Map({(d_74_v21_) - (d_74_v21_): d_75_v22_})
            d_76_v23_ = d_76_v23_
            d_77_v24_: _dafny.Seq
            d_77_v24_ = _dafny.SeqWithoutIsStrInference([d_72_v20_, d_72_v20_, d_72_v20_, d_72_v20_])
            rhs16_ = d_66_v14_
            rhs17_ = d_62_v9_
            rhs18_ = d_77_v24_
            d_66_v14_ = rhs16_
            d_62_v9_ = rhs17_
            d_77_v24_ = rhs18_
        elif source1_.is_DC4:
            d_78___mcc_h4_ = source1_.cf9
            d_79___mcc_h5_ = source1_.cf10
            d_80___mcc_h6_ = source1_.cf11
            d_81_cf11_ = d_80___mcc_h6_
            d_82_cf10_ = d_79___mcc_h5_
            d_83_cf9_ = d_78___mcc_h4_
            d_84_v25_: _dafny.Seq
            d_84_v25_ = _dafny.SeqWithoutIsStrInference([d_82_cf10_, p2, p2, p3, p2])
            d_85_v26_: D1
            d_85_v26_ = D1_DC1(d_84_v25_)
            d_85_v26_ = D1_DC1(_dafny.SeqWithoutIsStrInference([p2, p2]))
            source3_ = D1_DC2(d_33_v1_)
            if source3_.is_DC2:
                d_86___mcc_h13_ = source3_.cf5
                d_87_cf5_ = d_86___mcc_h13_
                d_88_v27_: C0
                nw4_ = C0()
                nw4_.ctor__(p3)
                d_88_v27_ = nw4_
                d_89_v28_: _dafny.Seq
                d_89_v28_ = _dafny.SeqWithoutIsStrInference([d_88_v27_])
                d_90_v29_: _dafny.Array
                nw5_ = _dafny.Array(False, 7)
                d_90_v29_ = nw5_
                d_91_v30_: _dafny.Map
                d_91_v30_ = _dafny.Map({d_83_cf9_: len(_dafny.Map({d_90_v29_: d_87_cf5_}))})
                d_92_v31_: _dafny.Map
                d_92_v31_ = _dafny.Map({d_82_cf10_: D1_DC4(d_83_cf9_, default__.fm4(len(d_89_v28_), (d_88_v27_).f17, globalState), _dafny.Map({d_83_cf9_: _dafny.MultiSet([d_83_cf9_, default__.fm0(len(d_35_v3_), d_91_v30_, globalState)])}))})
                d_93_v33_: _dafny.Map
                d_93_v33_ = _dafny.Map({len(d_89_v28_): d_32_v0_})
                d_94_v34_: D1
                def iife20_():
                    coll14_ = _dafny.Map()
                    compr_14_: int
                    for compr_14_ in (d_93_v33_).keys.Elements:
                        d_95_v32_: int = compr_14_
                        if (d_95_v32_) in (d_93_v33_):
                            coll14_[(d_95_v32_) + (p1)] = _dafny.MultiSet([d_83_cf9_, p1])
                    return _dafny.Map(coll14_)
                d_94_v34_ = D1_DC4(d_83_cf9_, (default__.fm13(globalState)).cf10, iife20_()
)
                d_82_cf10_ = ((d_92_v31_) | (d_92_v31_)) != (_dafny.Map({d_32_v0_: d_94_v34_}))
                d_96_v37_: _dafny.Array
                def lambda3_(d_97_cf5_, d_98_v3_, d_99_p1_, d_100_v1_, d_101_cf9_):
                    def lambda4_(d_102_i1_):
                        def iife21_():
                            coll15_ = _dafny.Map()
                            compr_15_: str
                            for compr_15_ in (_dafny.SeqWithoutIsStrInference([(d_98_v3_)[default__.safeIndex(d_99_p1_, len(d_98_v3_))], d_100_v1_])).Elements:
                                d_103_v35_: str = compr_15_
                                if (d_103_v35_) in (_dafny.SeqWithoutIsStrInference([(d_98_v3_)[default__.safeIndex(d_99_p1_, len(d_98_v3_))], d_100_v1_])):
                                    coll15_[d_103_v35_] = d_99_p1_
                            return _dafny.Map(coll15_)
                        def iife22_():
                            coll16_ = _dafny.Map()
                            compr_16_: str
                            for compr_16_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_100_v1_]))).Elements:
                                d_104_v36_: str = compr_16_
                                if (d_104_v36_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_100_v1_]))):
                                    coll16_[d_104_v36_] = (0) - (d_99_p1_)
                            return _dafny.Map(coll16_)
                        return (_dafny.Set({iife21_()
                        , iife22_()
                        })) | (_dafny.Set({_dafny.Map({d_97_cf5_: d_101_cf9_})}))

                    return lambda4_

                init1_ = lambda3_(d_87_cf5_, d_35_v3_, p1, d_33_v1_, d_83_cf9_)
                nw6_ = _dafny.Array(None, 18)
                for i0_1_ in range(nw6_.length(0)):
                    nw6_[i0_1_] = init1_(i0_1_)
                d_96_v37_ = nw6_
                d_105_v39_: _dafny.MultiSet
                d_105_v39_ = _dafny.MultiSet([_dafny.CodePoint('k')])
                d_106_v40_: _dafny.Map
                d_106_v40_ = _dafny.Map({d_87_cf5_: default__.fm0(p1, d_91_v30_, globalState)})
                d_107_v41_: _dafny.MultiSet
                def iife23_():
                    coll17_ = _dafny.Map()
                    compr_17_: str
                    for compr_17_ in (d_105_v39_).Elements:
                        d_108_v38_: str = compr_17_
                        if (d_108_v38_) in (d_105_v39_):
                            coll17_[d_108_v38_] = (_dafny.MultiSet([d_82_cf10_])).cardinality
                    return _dafny.Map(coll17_)
                d_107_v41_ = _dafny.MultiSet([iife23_()
                , d_106_v40_])
                d_109_v43_: _dafny.Set
                d_109_v43_ = _dafny.Set({d_106_v40_})
                index7_ = default__.safeIndex(87, (d_96_v37_).length(0))
                def iife24_():
                    coll18_ = _dafny.Set()
                    compr_18_: _dafny.Map
                    for compr_18_ in (d_107_v41_).Elements:
                        d_110_v42_: _dafny.Map = compr_18_
                        if (d_110_v42_) in (d_107_v41_):
                            coll18_ = coll18_.union(_dafny.Set([d_110_v42_]))
                    return _dafny.Set(coll18_)
                (d_96_v37_)[index7_] = (iife24_()
                ) - (d_109_v43_)
                index8_ = default__.safeIndex(87, (d_96_v37_).length(0))
                (d_96_v37_)[index8_] = (d_109_v43_).intersection(d_109_v43_)
                d_91_v30_ = (d_91_v30_) | (d_91_v30_)
            elif source3_.is_DC3:
                d_111___mcc_h14_ = source3_.cf6
                d_112___mcc_h15_ = source3_.cf7
                d_113___mcc_h16_ = source3_.cf8
                d_114_cf8_ = d_113___mcc_h16_
                d_115_cf7_ = d_112___mcc_h15_
                d_116_cf6_ = d_111___mcc_h14_
                d_117_v44_: _dafny.Set
                d_117_v44_ = _dafny.Set({len(d_116_cf6_)})
                d_118_v45_: _dafny.Array
                def lambda5_(d_119_v1_):
                    def lambda6_(d_120_i3_):
                        return d_119_v1_

                    return lambda6_

                init2_ = lambda5_(d_33_v1_)
                nw7_ = _dafny.Array(None, 16)
                for i0_2_ in range(nw7_.length(0)):
                    nw7_[i0_2_] = init2_(i0_2_)
                d_118_v45_ = nw7_
                d_82_cf10_ = (len(_dafny.SeqWithoutIsStrInference([len(d_35_v3_) for d_121_i2_ in range(default__.abs(607))]))) <= (default__.safeModuloInt(len(d_117_v44_), (_dafny.MultiSet([d_118_v45_])).cardinality))
                d_122_v46_: _dafny.Array
                nw8_ = _dafny.Array(False, 8)
                d_122_v46_ = nw8_
                d_123_v47_: C0
                nw9_ = C0()
                nw9_.ctor__(d_82_cf10_)
                d_123_v47_ = nw9_
                d_124_v48_: _dafny.MultiSet
                d_124_v48_ = _dafny.MultiSet([p2])
                d_125_v49_: _dafny.Map
                d_125_v49_ = _dafny.Map({p1: (d_123_v47_).f17})
                d_126_v50_: _dafny.Map
                d_126_v50_ = _dafny.Map({d_83_cf9_: d_125_v49_})
                d_127_v51_: _dafny.Map
                d_127_v51_ = _dafny.Map({d_123_v47_: default__.fm4(d_83_cf9_, (D0_DC0(((d_124_v48_)[default__.fm4(-968, p3, globalState)] if (default__.fm4(-968, p3, globalState)) in (d_124_v48_) else d_83_cf9_), default__.fm4(d_83_cf9_, not(d_32_v0_), globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "suaxnms")), d_126_v50_)).cf1, globalState)})
                index9_ = default__.safeIndex(780, (d_122_v46_).length(0))
                (d_122_v46_)[index9_] = ((d_127_v51_)[d_123_v47_] if (d_123_v47_) in (d_127_v51_) else d_32_v0_)
                index10_ = default__.safeIndex(780, (d_122_v46_).length(0))
                (d_122_v46_)[index10_] = (d_123_v47_).f17
                index11_ = default__.safeIndex(780, (d_122_v46_).length(0))
                def iife25_():
                    coll19_ = _dafny.Set()
                    compr_19_: int
                    for compr_19_ in _dafny.IntegerRange(509, 976):
                        d_128_v52_: int = compr_19_
                        if ((509) <= (d_128_v52_)) and ((d_128_v52_) < (976)):
                            coll19_ = coll19_.union(_dafny.Set([(d_128_v52_) + (d_83_cf9_)]))
                    return _dafny.Set(coll19_)
                (d_122_v46_)[index11_] = (d_117_v44_).issubset(iife25_()
                )
                d_129_v53_: _dafny.Seq
                d_129_v53_ = _dafny.SeqWithoutIsStrInference([d_122_v46_, d_122_v46_, d_122_v46_, d_122_v46_, d_122_v46_])
                d_129_v53_ = d_129_v53_
            elif source3_.is_DC4:
                d_130___mcc_h17_ = source3_.cf9
                d_131___mcc_h18_ = source3_.cf10
                d_132___mcc_h19_ = source3_.cf11
                d_133_cf11_ = d_132___mcc_h19_
                d_134_cf10_ = d_131___mcc_h18_
                d_135_cf9_ = d_130___mcc_h17_
                d_33_v1_ = d_33_v1_
                d_136_v54_: _dafny.MultiSet
                d_136_v54_ = _dafny.MultiSet([d_32_v0_])
                (globalState).f6 = ((d_136_v54_)[d_32_v0_] if (d_32_v0_) in (d_136_v54_) else (len(_dafny.SeqWithoutIsStrInference([not(p3)]))) + (d_135_cf9_))
                d_137_v55_: _dafny.Array
                nw10_ = _dafny.Array(int(0), 14)
                d_137_v55_ = nw10_
                d_138_v56_: _dafny.Map
                d_138_v56_ = _dafny.Map({d_134_cf10_: d_137_v55_})
                d_139_v57_: _dafny.Seq
                d_139_v57_ = _dafny.SeqWithoutIsStrInference([d_135_cf9_])
                d_140_v58_: _dafny.Array
                nw11_ = _dafny.Array(None, 3)
                nw11_[int(0)] = p3
                nw11_[int(1)] = p3
                nw11_[int(2)] = d_134_cf10_
                d_140_v58_ = nw11_
                d_141_v59_: _dafny.Map
                d_141_v59_ = _dafny.Map({(p1) + ((d_139_v57_)[default__.safeIndex(p1, len(d_139_v57_))]): d_140_v58_})
                d_142_v60_: _dafny.Seq
                d_142_v60_ = _dafny.SeqWithoutIsStrInference([d_138_v56_, (d_138_v56_) | (d_138_v56_)])
                d_143_v61_: _dafny.Map
                d_143_v61_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_33_v1_ for d_144_i4_ in range(default__.abs(823))]): p1})
                d_145_v62_: _dafny.Set
                d_145_v62_ = _dafny.Set({d_135_cf9_})
                d_146_v63_: C0
                nw12_ = C0()
                nw12_.ctor__(default__.fm4(len(_dafny.Set({len(_dafny.Map({d_32_v0_: p3})), d_135_cf9_})), p3, globalState))
                d_146_v63_ = nw12_
                d_147_v64_: _dafny.MultiSet
                d_147_v64_ = _dafny.MultiSet([d_146_v63_, d_146_v63_, d_146_v63_])
                d_148_v65_: _dafny.Map
                d_148_v65_ = _dafny.Map({len(d_145_v62_): (0) - ((d_147_v64_).cardinality)})
                d_149_v66_: _dafny.Map
                d_149_v66_ = _dafny.Map({p2: d_148_v65_})
                rhs19_ = (0) - ((0) - (default__.safeModuloInt(d_135_cf9_, d_135_cf9_)))
                rhs20_ = (d_142_v60_)[default__.safeIndex((d_135_cf9_) - (p1), len(d_142_v60_))]
                rhs21_ = default__.fm0((p1) * (len(d_143_v61_)), ((d_149_v66_)[True] if (True) in (d_149_v66_) else d_148_v65_), globalState)
                rhs22_ = d_141_v59_
                lhs11_ = globalState
                lhs12_ = globalState
                lhs11_.f11 = rhs19_
                d_138_v56_ = rhs20_
                lhs12_.f11 = rhs21_
                d_141_v59_ = rhs22_
                d_150_v67_: _dafny.Seq
                d_150_v67_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_151_i5_ in range(default__.abs(982))]), (_dafny.SeqWithoutIsStrInference([d_33_v1_ for d_152_i6_ in range(default__.abs(-300))])) + (d_35_v3_)])
                d_35_v3_ = ((d_150_v67_)[default__.safeIndex(default__.safeModuloInt(p1, p1), len(d_150_v67_))]).set(default__.safeIndex(d_135_cf9_, len((d_150_v67_)[default__.safeIndex(default__.safeModuloInt(p1, p1), len(d_150_v67_))])), _dafny.CodePoint('j'))
            elif source3_.is_DC1:
                d_153___mcc_h20_ = source3_.cf4
                d_154_cf4_ = d_153___mcc_h20_
                (globalState).f6 = d_83_cf9_
                d_155_v69_: C0
                nw13_ = C0()
                nw13_.ctor__(not(d_82_cf10_))
                d_155_v69_ = nw13_
                d_156_v70_: _dafny.Map
                d_156_v70_ = _dafny.Map({d_155_v69_: d_83_cf9_})
                def iife26_():
                    coll20_ = _dafny.Map()
                    compr_20_: str
                    for compr_20_ in (d_35_v3_).Elements:
                        d_157_v68_: str = compr_20_
                        if (d_157_v68_) in (d_35_v3_):
                            coll20_[d_157_v68_] = 860
                    return _dafny.Map(coll20_)
                (globalState).f6 = (len(iife26_()
                )) - ((len(d_156_v70_)) + (p1))
                d_155_v69_ = d_155_v69_
                d_158_v71_: D1
                d_158_v71_ = D1_DC2(default__.fm3(d_82_cf10_, globalState))
                def iife27_(_pat_let3_0):
                    def iife28_(d_159_dt__update__tmp_h1_):
                        def iife29_(_pat_let4_0):
                            def iife30_(d_160_dt__update_hcf5_h0_):
                                return D1_DC2(d_160_dt__update_hcf5_h0_)
                            return iife30_(_pat_let4_0)
                        return iife29_(_dafny.CodePoint('q'))
                    return iife28_(_pat_let3_0)
                d_158_v71_ = iife27_(d_158_v71_)
            elif True:
                d_161___mcc_h21_ = source3_.cf12
                d_162_cf12_ = d_161___mcc_h21_
                d_163_v72_: D1
                d_163_v72_ = D1_DC4(p1, True, d_81_cf11_)
                d_164_v73_: _dafny.Map
                d_164_v73_ = _dafny.Map({(d_163_v72_).cf10: _dafny.SeqWithoutIsStrInference([d_83_cf9_])})
                d_165_v74_: _dafny.Array
                nw14_ = _dafny.Array(int(0), 4)
                d_165_v74_ = nw14_
                d_166_v75_: C0
                nw15_ = C0()
                nw15_.ctor__(p2)
                d_166_v75_ = nw15_
                d_167_v76_: _dafny.Array
                nw16_ = _dafny.Array(None, 8)
                nw16_[int(0)] = d_166_v75_
                nw16_[int(1)] = d_166_v75_
                nw16_[int(2)] = d_166_v75_
                nw16_[int(3)] = d_166_v75_
                nw16_[int(4)] = d_166_v75_
                nw16_[int(5)] = d_166_v75_
                nw16_[int(6)] = d_166_v75_
                nw16_[int(7)] = d_166_v75_
                d_167_v76_ = nw16_
                d_168_v77_: _dafny.Map
                d_168_v77_ = _dafny.Map({d_165_v74_: d_167_v76_})
                d_169_v78_: _dafny.Set
                d_169_v78_ = _dafny.Set({(_dafny.MultiSet([d_83_cf9_])).cardinality})
                d_170_v79_: _dafny.Map
                d_170_v79_ = _dafny.Map({d_83_cf9_: len(d_169_v78_)})
                rhs23_ = d_83_cf9_
                rhs24_ = len((d_168_v77_) | (d_168_v77_))
                rhs25_ = (d_164_v73_) | (d_164_v73_)
                rhs26_ = default__.safeDivisionInt(default__.safeDivisionInt(p1, d_83_cf9_), len(d_170_v79_))
                rhs27_ = d_33_v1_
                lhs13_ = globalState
                lhs14_ = globalState
                r0 = rhs23_
                lhs13_.f16 = rhs24_
                d_164_v73_ = rhs25_
                lhs14_.f6 = rhs26_
                d_33_v1_ = rhs27_
                d_171_v80_: _dafny.Array
                nw17_ = _dafny.Array(False, 28)
                d_171_v80_ = nw17_
                index12_ = default__.safeIndex(77, (d_165_v74_).length(0))
                (d_165_v74_)[index12_] = ((d_170_v79_)[p1] if (p1) in (d_170_v79_) else (d_36_v4_).cf0)
                d_172_v81_: _dafny.Seq
                d_172_v81_ = _dafny.SeqWithoutIsStrInference([d_171_v80_, d_171_v80_])
                index13_ = default__.safeIndex(77, (d_165_v74_).length(0))
                rhs28_ = d_171_v80_
                rhs29_ = default__.safeDivisionInt(d_83_cf9_, p1)
                rhs30_ = len(d_172_v81_)
                rhs31_ = default__.safeDivisionInt(258, p1)
                lhs15_ = d_165_v74_
                lhs16_ = default__.safeIndex(77, (d_165_v74_).length(0))
                lhs17_ = globalState
                d_171_v80_ = rhs28_
                d_83_cf9_ = rhs29_
                lhs15_[lhs16_] = rhs30_
                lhs17_.f6 = rhs31_
                d_173_v82_: _dafny.MultiSet
                d_173_v82_ = _dafny.MultiSet([p3, (d_82_cf10_ if (d_166_v75_).f17 else False), d_32_v0_])
                d_173_v82_ = d_173_v82_
                index14_ = default__.safeIndex(72, (d_167_v76_).length(0))
                (d_167_v76_)[index14_] = d_166_v75_
                d_174_v83_: _dafny.Array
                nw18_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_174_v83_ = nw18_
                index15_ = default__.safeIndex(278, (d_174_v83_).length(0))
                (d_174_v83_)[index15_] = (d_172_v81_)[default__.safeIndex(-373, len(d_172_v81_))]
                index16_ = default__.safeIndex(562, (d_171_v80_).length(0))
                (d_171_v80_)[index16_] = p3
                index17_ = default__.safeIndex(72, (d_167_v76_).length(0))
                index18_ = default__.safeIndex(278, (d_174_v83_).length(0))
                index19_ = default__.safeIndex(562, (d_171_v80_).length(0))
                rhs32_ = d_166_v75_
                rhs33_ = (default__.safeModuloInt(p1, default__.fm0((d_165_v74_)[default__.safeIndex(77, (d_165_v74_).length(0))], d_170_v79_, globalState))) > ((d_165_v74_)[default__.safeIndex(77, (d_165_v74_).length(0))])
                rhs34_ = d_171_v80_
                rhs35_ = p2
                rhs36_ = p2
                lhs18_ = d_167_v76_
                lhs19_ = default__.safeIndex(72, (d_167_v76_).length(0))
                lhs20_ = d_174_v83_
                lhs21_ = default__.safeIndex(278, (d_174_v83_).length(0))
                lhs22_ = d_171_v80_
                lhs23_ = default__.safeIndex(562, (d_171_v80_).length(0))
                lhs18_[lhs19_] = rhs32_
                d_82_cf10_ = rhs33_
                lhs20_[lhs21_] = rhs34_
                lhs22_[lhs23_] = rhs35_
                d_32_v0_ = rhs36_
            d_175_v84_: _dafny.Array
            nw19_ = _dafny.Array(_dafny.Map({}), 9)
            d_175_v84_ = nw19_
            d_175_v84_ = d_175_v84_
            d_176_v85_: _dafny.Map
            d_176_v85_ = _dafny.Map({p1: d_83_cf9_})
            d_177_v86_: _dafny.MultiSet
            d_177_v86_ = _dafny.MultiSet([(d_84_v25_)[default__.safeIndex(p1, len(d_84_v25_))]])
            pat_let_tv7_ = d_177_v86_
            def iife31_(_pat_let5_0):
                def iife32_(d_178_dt__update__tmp_h2_):
                    def iife33_(_pat_let6_0):
                        def iife34_(d_179_dt__update_hcf7_h0_):
                            return D1_DC3((d_178_dt__update__tmp_h2_).cf6, d_179_dt__update_hcf7_h0_, (d_178_dt__update__tmp_h2_).cf8)
                        return iife34_(_pat_let6_0)
                    return iife33_((pat_let_tv7_).cardinality)
                return iife32_(_pat_let5_0)
            source4_ = iife31_(D1_DC3(d_176_v85_, p1, d_36_v4_))
            if source4_.is_DC2:
                d_180___mcc_h22_ = source4_.cf5
                d_181_cf5_ = d_180___mcc_h22_
                rhs37_ = (d_82_cf10_) or ((not(d_32_v0_)) or (d_32_v0_))
                rhs38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pufjq"))
                d_82_cf10_ = rhs37_
                d_35_v3_ = rhs38_
                d_182_v87_: _dafny.Array
                nw20_ = _dafny.Array(_dafny.Set({}), 5)
                d_182_v87_ = nw20_
                d_183_v88_: _dafny.Array
                nw21_ = _dafny.Array(int(0), 9)
                d_183_v88_ = nw21_
                index20_ = default__.safeIndex(634, (d_182_v87_).length(0))
                (d_182_v87_)[index20_] = _dafny.Set({d_183_v88_, d_183_v88_, d_183_v88_})
                d_184_v89_: _dafny.Set
                d_184_v89_ = _dafny.Set({d_183_v88_, d_183_v88_, d_183_v88_, d_183_v88_})
                d_185_v90_: _dafny.Seq
                d_185_v90_ = _dafny.SeqWithoutIsStrInference([d_184_v89_, d_184_v89_, _dafny.Set({d_183_v88_, d_183_v88_}), d_184_v89_])
                index21_ = default__.safeIndex(634, (d_182_v87_).length(0))
                (d_182_v87_)[index21_] = ((d_184_v89_).intersection((d_185_v90_)[default__.safeIndex(p1, len(d_185_v90_))])).intersection(d_184_v89_)
                d_82_cf10_ = (d_177_v86_).ispropersubset((_dafny.MultiSet(d_84_v25_)).set(p3, default__.abs(d_83_cf9_)))
                d_186_v91_: C0
                nw22_ = C0()
                nw22_.ctor__(p3)
                d_186_v91_ = nw22_
            elif source4_.is_DC3:
                d_187___mcc_h23_ = source4_.cf6
                d_188___mcc_h24_ = source4_.cf7
                d_189___mcc_h25_ = source4_.cf8
                d_190_cf8_ = d_189___mcc_h25_
                d_191_cf7_ = d_188___mcc_h24_
                d_192_cf6_ = d_187___mcc_h23_
                d_33_v1_ = d_33_v1_
                d_193_v92_: _dafny.Array
                nw23_ = _dafny.Array(int(0), 19)
                d_193_v92_ = nw23_
                d_194_v93_: _dafny.Map
                d_194_v93_ = _dafny.Map({p1: (d_35_v3_).set(default__.safeIndex(7, len(d_35_v3_)), _dafny.CodePoint('t'))})
                d_195_v94_: _dafny.Map
                d_195_v94_ = _dafny.Map({d_32_v0_: p1})
                index22_ = default__.safeIndex(57, (d_193_v92_).length(0))
                (d_193_v92_)[index22_] = len(((d_194_v93_)[((d_195_v94_)[d_82_cf10_] if (d_82_cf10_) in (d_195_v94_) else d_191_cf7_)] if (((d_195_v94_)[d_82_cf10_] if (d_82_cf10_) in (d_195_v94_) else d_191_cf7_)) in (d_194_v93_) else _dafny.SeqWithoutIsStrInference([d_33_v1_ for d_196_i7_ in range(default__.abs(587))])))
                d_197_v95_: _dafny.Array
                nw24_ = _dafny.Array(None, 18)
                nw24_[int(0)] = d_82_cf10_
                nw24_[int(1)] = d_82_cf10_
                nw24_[int(2)] = p2
                nw24_[int(3)] = p2
                nw24_[int(4)] = p3
                nw24_[int(5)] = p2
                nw24_[int(6)] = (-500) <= (p1)
                nw24_[int(7)] = d_82_cf10_
                nw24_[int(8)] = not((d_84_v25_)[default__.safeIndex(d_191_cf7_, len(d_84_v25_))])
                nw24_[int(9)] = d_32_v0_
                nw24_[int(10)] = ((default__.fm6(p2, p3, p2, globalState)).set(p2, default__.abs(d_83_cf9_))).isdisjoint(d_177_v86_)
                nw24_[int(11)] = p3
                nw24_[int(12)] = p3
                nw24_[int(13)] = False
                nw24_[int(14)] = (p1) == (d_191_cf7_)
                nw24_[int(15)] = p2
                nw24_[int(16)] = p2
                nw24_[int(17)] = True
                d_197_v95_ = nw24_
                index23_ = default__.safeIndex(156, (d_197_v95_).length(0))
                (d_197_v95_)[index23_] = d_32_v0_
                index24_ = default__.safeIndex(57, (d_193_v92_).length(0))
                index25_ = default__.safeIndex(156, (d_197_v95_).length(0))
                rhs39_ = (len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udu"))) + (d_35_v3_)).set(default__.safeIndex(p1, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udu"))) + (d_35_v3_))), d_33_v1_))) >= (345)
                rhs40_ = p3
                rhs41_ = default__.safeModuloInt(d_191_cf7_, d_83_cf9_)
                rhs42_ = d_82_cf10_
                rhs43_ = (d_191_cf7_) - (p1)
                lhs24_ = d_193_v92_
                lhs25_ = default__.safeIndex(57, (d_193_v92_).length(0))
                lhs26_ = d_197_v95_
                lhs27_ = default__.safeIndex(156, (d_197_v95_).length(0))
                lhs28_ = globalState
                d_32_v0_ = rhs39_
                d_82_cf10_ = rhs40_
                lhs24_[lhs25_] = rhs41_
                lhs26_[lhs27_] = rhs42_
                lhs28_.f16 = rhs43_
                d_198_v96_: D2
                d_198_v96_ = D2_DC6(d_193_v92_)
                d_199_v97_: _dafny.Seq
                d_199_v97_ = _dafny.SeqWithoutIsStrInference([(d_198_v96_).cf13, d_193_v92_, d_193_v92_])
                d_200_v98_: _dafny.Array
                nw25_ = _dafny.Array(None, 19)
                d_200_v98_ = nw25_
                d_201_v99_: C0
                nw26_ = C0()
                nw26_.ctor__(d_82_cf10_)
                d_201_v99_ = nw26_
                index26_ = default__.safeIndex(687, (d_200_v98_).length(0))
                (d_200_v98_)[index26_] = d_201_v99_
                index27_ = default__.safeIndex(156, (d_197_v95_).length(0))
                index28_ = default__.safeIndex(687, (d_200_v98_).length(0))
                rhs44_ = (d_84_v25_)[default__.safeIndex(d_83_cf9_, len(d_84_v25_))]
                rhs45_ = (((d_199_v97_) + (d_199_v97_) if (d_83_cf9_) > (p1) else d_199_v97_)).set(default__.safeIndex((d_193_v92_)[default__.safeIndex(57, (d_193_v92_).length(0))], len(((d_199_v97_) + (d_199_v97_) if (d_83_cf9_) > (p1) else d_199_v97_))), d_193_v92_)
                rhs46_ = d_201_v99_
                rhs47_ = d_197_v95_
                lhs29_ = d_197_v95_
                lhs30_ = default__.safeIndex(156, (d_197_v95_).length(0))
                lhs31_ = d_200_v98_
                lhs32_ = default__.safeIndex(687, (d_200_v98_).length(0))
                lhs29_[lhs30_] = rhs44_
                d_199_v97_ = rhs45_
                lhs31_[lhs32_] = rhs46_
                d_197_v95_ = rhs47_
                index29_ = default__.safeIndex(156, (d_197_v95_).length(0))
                (d_197_v95_)[index29_] = (d_83_cf9_) <= (p1)
            elif source4_.is_DC4:
                d_202___mcc_h26_ = source4_.cf9
                d_203___mcc_h27_ = source4_.cf10
                d_204___mcc_h28_ = source4_.cf11
                d_205_cf11_ = d_204___mcc_h28_
                d_206_cf10_ = d_203___mcc_h27_
                d_207_cf9_ = d_202___mcc_h26_
                d_208_v100_: C0
                nw27_ = C0()
                nw27_.ctor__(d_206_cf10_)
                d_208_v100_ = nw27_
                d_209_v101_: _dafny.Map
                d_209_v101_ = _dafny.Map({d_32_v0_: d_206_cf10_})
                (globalState).f6 = len(d_209_v101_)
                d_210_v102_: _dafny.Array
                def lambda7_(d_211_p3_):
                    def lambda8_(d_212_i8_):
                        return d_211_p3_

                    return lambda8_

                init3_ = lambda7_(p3)
                nw28_ = _dafny.Array(None, 2)
                for i0_3_ in range(nw28_.length(0)):
                    nw28_[i0_3_] = init3_(i0_3_)
                d_210_v102_ = nw28_
                index30_ = default__.safeIndex(678, (d_210_v102_).length(0))
                (d_210_v102_)[index30_] = p2
                index31_ = default__.safeIndex(678, (d_210_v102_).length(0))
                (d_210_v102_)[index31_] = ((d_209_v101_)[not((d_35_v3_) <= (d_35_v3_))] if (not((d_35_v3_) <= (d_35_v3_))) in (d_209_v101_) else d_206_cf10_)
                (globalState).f16 = len(d_35_v3_)
            elif source4_.is_DC1:
                d_213___mcc_h29_ = source4_.cf4
                d_214_cf4_ = d_213___mcc_h29_
                d_215_v103_: C0
                nw29_ = C0()
                nw29_.ctor__(not (p2) or (not(not(p2))))
                d_215_v103_ = nw29_
                d_216_v104_: _dafny.Set
                d_216_v104_ = _dafny.Set({(d_215_v103_).f17})
                d_217_v105_: _dafny.Map
                d_217_v105_ = _dafny.Map({d_83_cf9_: (d_84_v25_)[default__.safeIndex(len(_dafny.Map({(d_215_v103_).f17: (d_215_v103_).f17})), len(d_84_v25_))]})
                d_32_v0_ = default__.fm4(default__.safeDivisionInt(p1, len(d_216_v104_)), ((d_217_v105_)[d_83_cf9_] if (d_83_cf9_) in (d_217_v105_) else p2), globalState)
                d_218_v106_: _dafny.Array
                nw30_ = _dafny.Array(False, 2)
                d_218_v106_ = nw30_
                d_218_v106_ = d_218_v106_
                d_219_v107_: _dafny.Seq
                d_219_v107_ = _dafny.SeqWithoutIsStrInference([d_215_v103_])
                d_220_v108_: _dafny.Seq
                d_220_v108_ = _dafny.SeqWithoutIsStrInference([p1, p1, d_83_cf9_, p1, p1])
                rhs48_ = d_215_v103_
                rhs49_ = d_215_v103_
                rhs50_ = not((d_219_v107_) != (d_219_v107_))
                rhs51_ = (d_219_v107_)[default__.safeIndex((p1) * ((d_220_v108_)[default__.safeIndex(369, len(d_220_v108_))]), len(d_219_v107_))]
                d_215_v103_ = rhs48_
                d_215_v103_ = rhs49_
                d_32_v0_ = rhs50_
                d_215_v103_ = rhs51_
            elif True:
                d_221___mcc_h30_ = source4_.cf12
                d_222_cf12_ = d_221___mcc_h30_
                (globalState).f6 = p1
                d_223_v109_: C0
                nw31_ = C0()
                nw31_.ctor__(p3)
                d_223_v109_ = nw31_
                d_224_v110_: _dafny.Map
                d_224_v110_ = _dafny.Map({d_83_cf9_: d_223_v109_})
                d_224_v110_ = (d_224_v110_).set((0) - (d_83_cf9_), d_223_v109_)
                d_225_v111_: C0
                nw32_ = C0()
                nw32_.ctor__(p2)
                d_225_v111_ = nw32_
                rhs52_ = -453
                rhs53_ = (((0) - (p1)) - (p1)) * (len(default__.fm14(globalState)))
                lhs33_ = globalState
                lhs34_ = globalState
                lhs33_.f11 = rhs52_
                lhs34_.f6 = rhs53_
        elif source1_.is_DC1:
            d_226___mcc_h7_ = source1_.cf4
            d_227_cf4_ = d_226___mcc_h7_
            d_228_v112_: C0
            nw33_ = C0()
            nw33_.ctor__(p2)
            d_228_v112_ = nw33_
            d_228_v112_ = d_228_v112_
            d_229_v113_: _dafny.Map
            d_229_v113_ = _dafny.Map({d_35_v3_: d_228_v112_})
            d_228_v112_ = ((d_229_v113_)[(d_35_v3_ if not(not(p2)) else d_35_v3_)] if ((d_35_v3_ if not(not(p2)) else d_35_v3_)) in (d_229_v113_) else (d_228_v112_ if False else d_228_v112_))
            r0 = p1
            d_230_v114_: _dafny.Map
            d_230_v114_ = _dafny.Map({(d_228_v112_).f17: p1})
            d_231_v115_: _dafny.Seq
            d_231_v115_ = _dafny.SeqWithoutIsStrInference([d_228_v112_])
            d_232_v116_: _dafny.Map
            d_232_v116_ = _dafny.Map({p1: p1})
            d_233_v117_: _dafny.Map
            d_233_v117_ = _dafny.Map({default__.fm0(len(d_231_v115_), d_232_v116_, globalState): p1})
            d_234_v118_: _dafny.Map
            d_234_v118_ = _dafny.Map({p1: d_32_v0_})
            d_235_v119_: D1
            d_235_v119_ = D1_DC4(default__.safeDivisionInt(p1, p1), (d_227_cf4_)[default__.safeIndex(p1, len(d_227_cf4_))], _dafny.Map({p1: _dafny.MultiSet([((d_230_v114_)[(d_227_cf4_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([446, ((d_232_v116_)[p1] if (p1) in (d_232_v116_) else p1)])), len(d_227_cf4_))]] if ((d_227_cf4_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([446, ((d_232_v116_)[p1] if (p1) in (d_232_v116_) else p1)])), len(d_227_cf4_))]) in (d_230_v114_) else len(d_234_v118_))])}))
            source5_ = d_235_v119_
            if source5_.is_DC2:
                d_236___mcc_h31_ = source5_.cf5
                d_237_cf5_ = d_236___mcc_h31_
                d_238_v120_: _dafny.Set
                d_238_v120_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_237_cf5_ for d_239_i9_ in range(default__.abs(124))])})
                d_240_v122_: _dafny.Set
                d_240_v122_ = _dafny.Set({not((d_228_v112_).f17), False, p3, (d_228_v112_).f17, (default__.fm13(globalState)).cf10})
                d_241_v124_: _dafny.Seq
                def iife35_():
                    coll21_ = _dafny.Map()
                    compr_21_: int
                    for compr_21_ in _dafny.IntegerRange(92, 86):
                        d_242_v123_: int = compr_21_
                        if ((92) <= (d_242_v123_)) and ((d_242_v123_) < (86)):
                            coll21_[(d_242_v123_) * (2)] = p2
                    return _dafny.Map(coll21_)
                d_241_v124_ = _dafny.SeqWithoutIsStrInference([d_234_v118_, iife35_()
                ])
                d_243_v125_: _dafny.Array
                nw34_ = _dafny.Array(None, 15)
                nw34_[int(0)] = p2
                nw34_[int(1)] = p3
                nw34_[int(2)] = not(p3)
                nw34_[int(3)] = (d_228_v112_).f17
                nw34_[int(4)] = not(d_32_v0_)
                def iife36_():
                    coll22_ = _dafny.Set()
                    compr_22_: _dafny.Seq
                    for compr_22_ in (d_238_v120_).Elements:
                        d_244_v121_: _dafny.Seq = compr_22_
                        if (d_244_v121_) in (d_238_v120_):
                            coll22_ = coll22_.union(_dafny.Set([d_244_v121_]))
                    return _dafny.Set(coll22_)
                nw34_[int(5)] = (iife36_()
                ).ispropersubset(d_238_v120_)
                nw34_[int(6)] = (d_228_v112_).f17
                nw34_[int(7)] = p3
                nw34_[int(8)] = (len(d_240_v122_)) > (p1)
                nw34_[int(9)] = p2
                nw34_[int(10)] = p2
                nw34_[int(11)] = p3
                nw34_[int(12)] = p3
                nw34_[int(13)] = (d_228_v112_).f17
                nw34_[int(14)] = (_dafny.Map({p1: True})) in (d_241_v124_)
                d_243_v125_ = nw34_
                d_245_v126_: _dafny.MultiSet
                d_245_v126_ = _dafny.MultiSet([p1])
                d_246_v127_: _dafny.Map
                d_246_v127_ = _dafny.Map({p1: d_245_v126_})
                index32_ = default__.safeIndex(328, (d_243_v125_).length(0))
                (d_243_v125_)[index32_] = ((D1_DC4(p1, p2, d_246_v127_)).cf10 if (d_228_v112_).f17 else (d_228_v112_).f17)
                d_247_v128_: _dafny.Set
                d_247_v128_ = _dafny.Set({p1})
                index33_ = default__.safeIndex(328, (d_243_v125_).length(0))
                (d_243_v125_)[index33_] = ((d_247_v128_).ispropersubset(d_247_v128_)) and (p3)
                d_248_v129_: C0
                nw35_ = C0()
                nw35_.ctor__((p1) > (p1))
                d_248_v129_ = nw35_
                d_249_v131_: _dafny.Array
                nw36_ = _dafny.Array(None, 27)
                nw36_[int(0)] = p1
                nw36_[int(1)] = p1
                nw36_[int(2)] = p1
                nw36_[int(3)] = p1
                nw36_[int(4)] = p1
                nw36_[int(5)] = p1
                nw36_[int(6)] = p1
                nw36_[int(7)] = p1
                nw36_[int(8)] = p1
                nw36_[int(9)] = p1
                nw36_[int(10)] = (0) - (p1)
                def iife37_():
                    coll23_ = _dafny.Map()
                    compr_23_: int
                    for compr_23_ in (d_234_v118_).keys.Elements:
                        d_250_v130_: int = compr_23_
                        if (d_250_v130_) in (d_234_v118_):
                            coll23_[(d_250_v130_) + (p1)] = (d_228_v112_).f17
                    return _dafny.Map(coll23_)
                nw36_[int(11)] = len(iife37_()
                )
                nw36_[int(12)] = p1
                nw36_[int(13)] = len(d_35_v3_)
                nw36_[int(14)] = p1
                nw36_[int(15)] = p1
                nw36_[int(16)] = p1
                nw36_[int(17)] = p1
                nw36_[int(18)] = (d_245_v126_).cardinality
                nw36_[int(19)] = (0) - (p1)
                nw36_[int(20)] = len(d_35_v3_)
                nw36_[int(21)] = p1
                nw36_[int(22)] = p1
                nw36_[int(23)] = p1
                nw36_[int(24)] = 930
                nw36_[int(25)] = (0) - (p1)
                nw36_[int(26)] = p1
                d_249_v131_ = nw36_
                d_251_v132_: _dafny.Seq
                d_251_v132_ = _dafny.SeqWithoutIsStrInference([d_249_v131_, d_249_v131_])
                d_252_v133_: _dafny.Array
                nw37_ = _dafny.Array(None, 12)
                nw37_[int(0)] = d_251_v132_
                nw37_[int(1)] = d_251_v132_
                nw37_[int(2)] = _dafny.SeqWithoutIsStrInference([d_249_v131_])
                nw37_[int(3)] = d_251_v132_
                nw37_[int(4)] = d_251_v132_
                nw37_[int(5)] = d_251_v132_
                nw37_[int(6)] = (_dafny.SeqWithoutIsStrInference([d_249_v131_])) + (d_251_v132_)
                nw37_[int(7)] = d_251_v132_
                nw37_[int(8)] = (d_251_v132_) + (d_251_v132_)
                nw37_[int(9)] = (_dafny.SeqWithoutIsStrInference([d_249_v131_]) if p2 else d_251_v132_)
                nw37_[int(10)] = (d_251_v132_) + (d_251_v132_)
                nw37_[int(11)] = d_251_v132_
                d_252_v133_ = nw37_
                index34_ = default__.safeIndex(807, (d_252_v133_).length(0))
                (d_252_v133_)[index34_] = d_251_v132_
                index35_ = default__.safeIndex(807, (d_252_v133_).length(0))
                rhs54_ = d_251_v132_
                rhs55_ = ((len(d_240_v122_)) + (p1)) * (p1)
                rhs56_ = not(not((p2) or ((d_228_v112_).f17)))
                rhs57_ = p1
                rhs58_ = d_249_v131_
                lhs35_ = d_252_v133_
                lhs36_ = default__.safeIndex(807, (d_252_v133_).length(0))
                lhs37_ = globalState
                lhs38_ = globalState
                lhs35_[lhs36_] = rhs54_
                lhs37_.f11 = rhs55_
                d_32_v0_ = rhs56_
                lhs38_.f11 = rhs57_
                d_249_v131_ = rhs58_
                d_32_v0_ = (d_243_v125_)[default__.safeIndex(328, (d_243_v125_).length(0))]
            elif source5_.is_DC3:
                d_253___mcc_h32_ = source5_.cf6
                d_254___mcc_h33_ = source5_.cf7
                d_255___mcc_h34_ = source5_.cf8
                d_256_cf8_ = d_255___mcc_h34_
                d_257_cf7_ = d_254___mcc_h33_
                d_258_cf6_ = d_253___mcc_h32_
                (globalState).f16 = (p1 if d_32_v0_ else d_257_cf7_)
                d_259_v134_: _dafny.Map
                d_259_v134_ = _dafny.Map({d_33_v1_: not(p3)})
                d_260_v135_: _dafny.Set
                d_260_v135_ = _dafny.Set({(d_228_v112_).f17, True})
                d_259_v134_ = default__.fm15(d_32_v0_, d_260_v135_, not(p2), globalState)
                d_261_v136_: D3
                d_261_v136_ = D3_DC10(d_231_v115_)
                d_231_v115_ = (d_261_v136_).cf18
                d_232_v116_ = (d_232_v116_).set(d_257_cf7_, (d_257_cf7_) * (default__.fm0(p1, d_258_cf6_, globalState)))
            elif source5_.is_DC4:
                d_262___mcc_h35_ = source5_.cf9
                d_263___mcc_h36_ = source5_.cf10
                d_264___mcc_h37_ = source5_.cf11
                d_265_cf11_ = d_264___mcc_h37_
                d_266_cf10_ = d_263___mcc_h36_
                d_267_cf9_ = d_262___mcc_h35_
                d_268_v137_: _dafny.Array
                nw38_ = _dafny.Array(_dafny.Seq({}), 4)
                d_268_v137_ = nw38_
                d_269_v138_: _dafny.Set
                d_269_v138_ = _dafny.Set({d_228_v112_, d_228_v112_})
                d_270_v139_: _dafny.Array
                nw39_ = _dafny.Array(None, 4)
                nw39_[int(0)] = -142
                nw39_[int(1)] = -626
                nw39_[int(2)] = p1
                nw39_[int(3)] = len(d_269_v138_)
                d_270_v139_ = nw39_
                d_271_v140_: _dafny.MultiSet
                d_271_v140_ = _dafny.MultiSet([d_270_v139_])
                d_272_v141_: _dafny.Seq
                d_272_v141_ = _dafny.SeqWithoutIsStrInference([default__.fm0((d_271_v140_).cardinality, d_232_v116_, globalState)])
                index36_ = default__.safeIndex(633, (d_268_v137_).length(0))
                (d_268_v137_)[index36_] = d_272_v141_
                d_273_v142_: D1
                d_273_v142_ = D1_DC2(d_33_v1_)
                d_274_v143_: _dafny.Seq
                d_274_v143_ = _dafny.SeqWithoutIsStrInference([d_273_v142_, d_273_v142_, d_273_v142_])
                d_275_v145_: _dafny.MultiSet
                def iife38_():
                    coll24_ = _dafny.Map()
                    compr_24_: int
                    for compr_24_ in _dafny.IntegerRange(998, 785):
                        d_276_v144_: int = compr_24_
                        if ((998) <= (d_276_v144_)) and ((d_276_v144_) < (785)):
                            coll24_[(d_276_v144_) - (p1)] = d_267_cf9_
                    return _dafny.Map(coll24_)
                d_275_v145_ = _dafny.MultiSet([iife38_()
                , d_232_v116_, default__.fm1(d_232_v116_, d_267_cf9_, d_267_cf9_, globalState), d_232_v116_])
                d_277_v146_: _dafny.Seq
                d_277_v146_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_267_cf9_ for d_278_i10_ in range(default__.abs(541))]), d_272_v141_, _dafny.SeqWithoutIsStrInference([p1, len(_dafny.SeqWithoutIsStrInference([454, len(d_274_v143_), p1, (0) - (len(d_35_v3_))])), p1]), d_272_v141_, (_dafny.SeqWithoutIsStrInference([((d_275_v145_)[d_232_v116_] if (d_232_v116_) in (d_275_v145_) else d_267_cf9_)])).set(default__.safeIndex(d_267_cf9_, len(_dafny.SeqWithoutIsStrInference([((d_275_v145_)[d_232_v116_] if (d_232_v116_) in (d_275_v145_) else d_267_cf9_)]))), p1)])
                index37_ = default__.safeIndex(633, (d_268_v137_).length(0))
                (d_268_v137_)[index37_] = (d_277_v146_)[default__.safeIndex(p1, len(d_277_v146_))]
                d_279_v147_: C0
                nw40_ = C0()
                nw40_.ctor__(d_266_cf10_)
                d_279_v147_ = nw40_
                d_280_v148_: D4
                d_280_v148_ = D4_DC12(d_230_v114_)
                d_230_v114_ = (d_230_v114_) | ((d_280_v148_).cf19)
                d_281_v149_: _dafny.Array
                nw41_ = _dafny.Array(False, 28)
                d_281_v149_ = nw41_
                index38_ = default__.safeIndex(609, (d_281_v149_).length(0))
                (d_281_v149_)[index38_] = default__.fm4(default__.fm0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ns"))), d_232_v116_, globalState), p2, globalState)
                index39_ = default__.safeIndex(609, (d_281_v149_).length(0))
                (d_281_v149_)[index39_] = d_266_cf10_
            elif source5_.is_DC1:
                d_282___mcc_h38_ = source5_.cf4
                d_283_cf4_ = d_282___mcc_h38_
                d_232_v116_ = d_232_v116_
                d_284_v150_: _dafny.MultiSet
                d_284_v150_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhlfxth")))])
                d_284_v150_ = d_284_v150_
                d_285_v151_: _dafny.Seq
                d_285_v151_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_286_v152_: _dafny.Set
                d_286_v152_ = _dafny.Set({p1, default__.fm0(len(_dafny.SeqWithoutIsStrInference([p2, p3])), d_232_v116_, globalState), p1, default__.fm0((d_285_v151_)[default__.safeIndex(p1, len(d_285_v151_))], _dafny.Map({p1: p1}), globalState)})
                d_287_v153_: _dafny.Map
                d_287_v153_ = _dafny.Map({p1: d_286_v152_})
                d_288_v154_: _dafny.Set
                d_288_v154_ = _dafny.Set({((d_287_v153_)[p1] if (p1) in (d_287_v153_) else _dafny.Set({p1, p1, p1})), d_286_v152_, d_286_v152_, _dafny.Set({len(_dafny.Set({(d_228_v112_).f17}))})})
                d_32_v0_ = (d_286_v152_) in (d_288_v154_)
                rhs59_ = False
                rhs60_ = p2
                d_32_v0_ = rhs59_
                d_32_v0_ = rhs60_
            elif True:
                d_289___mcc_h39_ = source5_.cf12
                d_290_cf12_ = d_289___mcc_h39_
                d_291_v155_: _dafny.Map
                d_291_v155_ = _dafny.Map({p1: d_234_v118_})
                pat_let_tv8_ = d_35_v3_
                pat_let_tv9_ = p1
                pat_let_tv10_ = d_291_v155_
                d_292_v156_: _dafny.Array
                nw42_ = _dafny.Array(None, 4)
                def iife39_(_pat_let7_0):
                    def iife40_(d_293_dt__update__tmp_h3_):
                        def iife41_(_pat_let8_0):
                            def iife42_(d_294_dt__update_hcf2_h1_):
                                def iife43_(_pat_let9_0):
                                    def iife44_(d_295_dt__update_hcf0_h0_):
                                        def iife45_(_pat_let10_0):
                                            def iife46_(d_296_dt__update_hcf3_h0_):
                                                return D0_DC0(d_295_dt__update_hcf0_h0_, (d_293_dt__update__tmp_h3_).cf1, d_294_dt__update_hcf2_h1_, d_296_dt__update_hcf3_h0_)
                                            return iife46_(_pat_let10_0)
                                        return iife45_(pat_let_tv10_)
                                    return iife44_(_pat_let9_0)
                                return iife43_(pat_let_tv9_)
                            return iife42_(_pat_let8_0)
                        return iife41_(pat_let_tv8_)
                    return iife40_(_pat_let7_0)
                nw42_[int(0)] = iife39_(d_36_v4_)
                nw42_[int(1)] = d_36_v4_
                nw42_[int(2)] = d_36_v4_
                nw42_[int(3)] = d_36_v4_
                d_292_v156_ = nw42_
                index40_ = default__.safeIndex(210, (d_292_v156_).length(0))
                (d_292_v156_)[index40_] = d_36_v4_
                d_297_v158_: _dafny.Map
                d_297_v158_ = _dafny.Map({D3_DC11(): (d_228_v112_).f17})
                pat_let_tv11_ = d_227_cf4_
                pat_let_tv12_ = p2
                pat_let_tv13_ = globalState
                pat_let_tv14_ = p1
                index41_ = default__.safeIndex(210, (d_292_v156_).length(0))
                def iife48_():
                    coll25_ = _dafny.Map()
                    compr_25_: D3
                    for compr_25_ in (d_297_v158_).keys.Elements:
                        d_298_v157_: D3 = compr_25_
                        if (d_298_v157_) in (d_297_v158_):
                            coll25_[d_298_v157_] = p1
                    return _dafny.Map(coll25_)
                def iife47_(_pat_let11_0):
                    def iife49_(d_299_dt__update__tmp_h4_):
                        def iife50_(_pat_let12_0):
                            def iife51_(d_300_dt__update_hcf3_h1_):
                                def iife52_(_pat_let13_0):
                                    def iife53_(d_301_dt__update_hcf0_h1_):
                                        return D0_DC0(d_301_dt__update_hcf0_h1_, (d_299_dt__update__tmp_h4_).cf1, (d_299_dt__update__tmp_h4_).cf2, d_300_dt__update_hcf3_h1_)
                                    return iife53_(_pat_let13_0)
                                return iife52_(pat_let_tv14_)
                            return iife51_(_pat_let12_0)
                        return iife50_(default__.fm11(len(pat_let_tv11_), pat_let_tv12_, pat_let_tv13_))
                    return iife49_(_pat_let11_0)
                (d_292_v156_)[index41_] = iife47_(default__.fm16(d_33_v1_, default__.fm17(597, globalState), iife48_()
                , globalState))
                d_302_v159_: C0
                nw43_ = C0()
                nw43_.ctor__((d_228_v112_).f17)
                d_302_v159_ = nw43_
                d_303_v160_: _dafny.Array
                nw44_ = _dafny.Array(int(0), 4)
                d_303_v160_ = nw44_
                d_304_v161_: _dafny.Seq
                d_304_v161_ = _dafny.SeqWithoutIsStrInference([(d_303_v160_ if p3 else d_303_v160_), d_303_v160_, d_303_v160_, d_303_v160_])
                d_304_v161_ = (_dafny.SeqWithoutIsStrInference([d_303_v160_])) + (d_304_v161_)
                d_32_v0_ = d_32_v0_
        elif True:
            d_305___mcc_h8_ = source1_.cf12
            d_306_cf12_ = d_305___mcc_h8_
            if d_32_v0_:
                d_307_v162_: _dafny.MultiSet
                d_307_v162_ = _dafny.MultiSet([p2, not(not(p3))])
                d_308_v163_: C0
                nw45_ = C0()
                nw45_.ctor__(p2)
                d_308_v163_ = nw45_
                d_309_v164_: _dafny.Map
                d_309_v164_ = _dafny.Map({(d_307_v162_).intersection((d_307_v162_).set(d_32_v0_, default__.abs(p1))): d_308_v163_})
                d_309_v164_ = (d_309_v164_).set(d_307_v162_, d_308_v163_)
                d_32_v0_ = p3
                d_310_v165_: _dafny.Array
                nw46_ = _dafny.Array(None, 19)
                d_310_v165_ = nw46_
                index42_ = default__.safeIndex(738, (d_310_v165_).length(0))
                (d_310_v165_)[index42_] = d_308_v163_
                index43_ = default__.safeIndex(738, (d_310_v165_).length(0))
                (d_310_v165_)[index43_] = d_308_v163_
                d_311_v166_: _dafny.Map
                d_311_v166_ = _dafny.Map({d_32_v0_: True})
                d_311_v166_ = ((d_311_v166_).set((d_308_v163_).f17, (d_308_v163_).f17) if (d_308_v163_).f17 else (d_311_v166_) | (d_311_v166_))
                (globalState).f11 = p1
            elif True:
                rhs61_ = p1
                rhs62_ = not(p2)
                lhs39_ = globalState
                lhs39_.f6 = rhs61_
                d_32_v0_ = rhs62_
                d_312_v167_: C0
                nw47_ = C0()
                nw47_.ctor__(p3)
                d_312_v167_ = nw47_
                d_313_v168_: _dafny.Array
                def lambda9_(d_314_i11_):
                    return True

                init4_ = lambda9_
                nw48_ = _dafny.Array(None, 28)
                for i0_4_ in range(nw48_.length(0)):
                    nw48_[i0_4_] = init4_(i0_4_)
                d_313_v168_ = nw48_
                index44_ = default__.safeIndex(524, (d_313_v168_).length(0))
                (d_313_v168_)[index44_] = ((d_312_v167_).f17 if (d_312_v167_).f17 else d_32_v0_)
                index45_ = default__.safeIndex(524, (d_313_v168_).length(0))
                (d_313_v168_)[index45_] = p2
                d_32_v0_ = default__.fm4(p1, (d_313_v168_)[default__.safeIndex(524, (d_313_v168_).length(0))], globalState)
                d_315_v169_: _dafny.Map
                d_315_v169_ = _dafny.Map({d_32_v0_: p1})
                d_316_v170_: _dafny.Set
                d_316_v170_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([(d_313_v168_)[default__.safeIndex(524, (d_313_v168_).length(0))], (d_313_v168_)[default__.safeIndex(524, (d_313_v168_).length(0))]])), ((d_315_v169_)[p2] if (p2) in (d_315_v169_) else p1), (0) - (p1)})
                index46_ = default__.safeIndex(524, (d_313_v168_).length(0))
                (d_313_v168_)[index46_] = not((_dafny.Set({p1})).isdisjoint(d_316_v170_))
            d_317_v171_: _dafny.MultiSet
            d_317_v171_ = _dafny.MultiSet([d_32_v0_, True])
            if default__.fm4(p1, ((d_317_v171_).cardinality) != (p1), globalState):
                d_318_v172_: _dafny.Array
                nw49_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
                d_318_v172_ = nw49_
                index47_ = default__.safeIndex(131, (d_318_v172_).length(0))
                (d_318_v172_)[index47_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kinwdfjx"))) + (d_35_v3_)
                index48_ = default__.safeIndex(131, (d_318_v172_).length(0))
                (d_318_v172_)[index48_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_319_i12_ in range(default__.abs(468))])
                d_320_v173_: C0
                nw50_ = C0()
                nw50_.ctor__((d_317_v171_).ispropersubset(d_317_v171_))
                d_320_v173_ = nw50_
                d_321_v174_: _dafny.Map
                d_321_v174_ = _dafny.Map({522: p1})
                d_322_v175_: _dafny.Map
                d_322_v175_ = _dafny.Map({p1: True})
                (globalState).f16 = (default__.fm0(p1, d_321_v174_, globalState)) + (len((d_322_v175_ if p3 else default__.fm5(globalState))))
                index49_ = default__.safeIndex(131, (d_318_v172_).length(0))
                (d_318_v172_)[index49_] = (d_318_v172_)[default__.safeIndex(131, (d_318_v172_).length(0))]
                d_323_v176_: _dafny.Array
                nw51_ = _dafny.Array(_dafny.MultiSet({}), 12)
                d_323_v176_ = nw51_
                index50_ = default__.safeIndex(489, (d_323_v176_).length(0))
                (d_323_v176_)[index50_] = (d_317_v171_).set(d_32_v0_, default__.abs(p1))
                d_324_v177_: _dafny.Array
                nw52_ = _dafny.Array(D4.default()(), 20)
                d_324_v177_ = nw52_
                d_325_v178_: _dafny.Array
                def lambda10_(d_326_v1_):
                    def lambda11_(d_327_i13_):
                        return d_326_v1_

                    return lambda11_

                init5_ = lambda10_(d_33_v1_)
                nw53_ = _dafny.Array(None, 7)
                for i0_5_ in range(nw53_.length(0)):
                    nw53_[i0_5_] = init5_(i0_5_)
                d_325_v178_ = nw53_
                d_328_v179_: _dafny.Set
                d_328_v179_ = _dafny.Set({d_325_v178_})
                d_329_v180_: D4
                d_329_v180_ = D4_DC15(d_328_v179_, d_33_v1_)
                index51_ = default__.safeIndex(875, (d_324_v177_).length(0))
                (d_324_v177_)[index51_] = d_329_v180_
                d_330_v181_: _dafny.Array
                nw54_ = _dafny.Array(False, 27)
                d_330_v181_ = nw54_
                index52_ = default__.safeIndex(489, (d_323_v176_).length(0))
                index53_ = default__.safeIndex(875, (d_324_v177_).length(0))
                rhs63_ = (default__.fm6(p2, p3, p2, globalState) if p2 else d_317_v171_)
                rhs64_ = D4_DC15(d_328_v179_, d_33_v1_)
                rhs65_ = (D4_DC13(default__.fm4(p1, p3, globalState), p1, d_330_v181_)).cf20
                lhs40_ = d_323_v176_
                lhs41_ = default__.safeIndex(489, (d_323_v176_).length(0))
                lhs42_ = d_324_v177_
                lhs43_ = default__.safeIndex(875, (d_324_v177_).length(0))
                lhs40_[lhs41_] = rhs63_
                lhs42_[lhs43_] = rhs64_
                d_32_v0_ = rhs65_
            elif True:
                d_331_v182_: _dafny.Seq
                d_331_v182_ = _dafny.SeqWithoutIsStrInference([not(p2)])
                d_32_v0_ = (d_331_v182_)[default__.safeIndex((0) - ((0) - (-856)), len(d_331_v182_))]
                d_332_v183_: _dafny.Map
                d_332_v183_ = _dafny.Map({p1: p1})
                d_333_v184_: C0
                nw55_ = C0()
                nw55_.ctor__(p2)
                d_333_v184_ = nw55_
                d_334_v185_: _dafny.Map
                d_334_v185_ = _dafny.Map({p1: d_333_v184_})
                d_335_v186_: _dafny.Map
                d_335_v186_ = _dafny.Map({default__.fm0(p1, d_332_v183_, globalState): (d_334_v185_) | (d_334_v185_)})
                d_336_v187_: _dafny.Seq
                d_336_v187_ = _dafny.SeqWithoutIsStrInference([d_335_v186_])
                d_335_v186_ = (d_335_v186_) | ((d_336_v187_)[default__.safeIndex(len(_dafny.Map({not(False): True})), len(d_336_v187_))])
                d_337_v188_: _dafny.Map
                d_337_v188_ = _dafny.Map({(d_333_v184_).f17: d_33_v1_})
                d_338_v189_: _dafny.Set
                d_338_v189_ = _dafny.Set({d_35_v3_})
                d_339_v190_: D4
                d_339_v190_ = D4_DC14(len(d_337_v188_), (len(d_332_v183_)) - (p1), d_338_v189_)
                d_339_v190_ = d_339_v190_
                d_340_v191_: _dafny.Array
                nw56_ = _dafny.Array(False, 26)
                d_340_v191_ = nw56_
                d_340_v191_ = d_340_v191_
                d_32_v0_ = not(not (p3) or (d_32_v0_))
            d_341_v192_: _dafny.Array
            def lambda12_(d_342_v171_, d_343_v3_):
                def lambda13_(d_344_i14_):
                    return D4_DC14((d_342_v171_).cardinality, 996, _dafny.Set({d_343_v3_, d_343_v3_}))

                return lambda13_

            init6_ = lambda12_(d_317_v171_, d_35_v3_)
            nw57_ = _dafny.Array(None, 1)
            for i0_6_ in range(nw57_.length(0)):
                nw57_[i0_6_] = init6_(i0_6_)
            d_341_v192_ = nw57_
            d_345_v193_: C0
            nw58_ = C0()
            nw58_.ctor__(not(p3))
            d_345_v193_ = nw58_
            d_346_v194_: D4
            d_346_v194_ = D4_DC14(len(d_35_v3_), len(_dafny.SeqWithoutIsStrInference([d_345_v193_])), default__.fm18(p1, p1, p1, d_33_v1_, globalState))
            d_347_v195_: _dafny.Set
            d_347_v195_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpagmi"))})
            pat_let_tv15_ = d_347_v195_
            index54_ = default__.safeIndex(13, (d_341_v192_).length(0))
            def iife54_(_pat_let14_0):
                def iife55_(d_348_dt__update__tmp_h5_):
                    def iife56_(_pat_let15_0):
                        def iife57_(d_349_dt__update_hcf25_h0_):
                            return D4_DC14((d_348_dt__update__tmp_h5_).cf23, (d_348_dt__update__tmp_h5_).cf24, d_349_dt__update_hcf25_h0_)
                        return iife57_(_pat_let15_0)
                    return iife56_(pat_let_tv15_)
                return iife55_(_pat_let14_0)
            (d_341_v192_)[index54_] = iife54_(d_346_v194_)
            index55_ = default__.safeIndex(13, (d_341_v192_).length(0))
            (d_341_v192_)[index55_] = D4_DC14(p1, p1, _dafny.Set({d_35_v3_}))
            d_350_v196_: _dafny.Seq
            d_350_v196_ = _dafny.SeqWithoutIsStrInference([d_32_v0_])
            d_351_v197_: _dafny.Map
            d_351_v197_ = _dafny.Map({p3: len(d_350_v196_)})
            d_352_v198_: _dafny.Map
            d_352_v198_ = _dafny.Map({d_32_v0_: d_351_v197_})
            d_352_v198_ = default__.fm19((_dafny.MultiSet([d_32_v0_, p3, (d_345_v193_).f17, p2])) - (d_317_v171_), globalState)
        d_353_v199_: C0
        nw59_ = C0()
        nw59_.ctor__(p2)
        d_353_v199_ = nw59_
        d_354_i15_: int
        d_354_i15_ = 0
        with _dafny.label("0"):
            while p3:
                with _dafny.c_label("0"):
                    if (d_354_i15_) >= (100):
                        raise _dafny.Break("0")
                    d_354_i15_ = (d_354_i15_) + (1)
                    d_355_v200_: _dafny.Array
                    nw60_ = _dafny.Array(_dafny.Array(None, 0), 25)
                    d_355_v200_ = nw60_
                    d_356_v201_: _dafny.Map
                    d_356_v201_ = _dafny.Map({d_355_v200_: d_353_v199_})
                    d_356_v201_ = (d_356_v201_).set(d_355_v200_, d_353_v199_)
                    (globalState).f16 = p1
                    if (d_353_v199_).f17:
                        d_32_v0_ = p3
                        d_33_v1_ = d_33_v1_
                        d_357_v202_: _dafny.Array
                        nw61_ = _dafny.Array(_dafny.Set({}), 1)
                        d_357_v202_ = nw61_
                        index56_ = default__.safeIndex(172, (d_357_v202_).length(0))
                        (d_357_v202_)[index56_] = _dafny.Set({p3, default__.fm4(p1, (d_353_v199_).f17, globalState)})
                        d_358_v203_: _dafny.Set
                        d_358_v203_ = _dafny.Set({not((d_353_v199_).f17)})
                        index57_ = default__.safeIndex(172, (d_357_v202_).length(0))
                        (d_357_v202_)[index57_] = d_358_v203_
                        nw62_ = C0()
                        nw62_.ctor__((d_353_v199_).f17)
                        d_353_v199_ = nw62_
                        d_359_v204_: _dafny.Array
                        nw63_ = _dafny.Array(None, 13)
                        d_359_v204_ = nw63_
                        index58_ = default__.safeIndex(107, (d_359_v204_).length(0))
                        (d_359_v204_)[index58_] = d_353_v199_
                        index59_ = default__.safeIndex(107, (d_359_v204_).length(0))
                        rhs66_ = d_32_v0_
                        rhs67_ = d_35_v3_
                        rhs68_ = len(d_35_v3_)
                        rhs69_ = (d_353_v199_ if (d_353_v199_).f17 else d_353_v199_)
                        lhs44_ = globalState
                        lhs45_ = d_359_v204_
                        lhs46_ = default__.safeIndex(107, (d_359_v204_).length(0))
                        d_32_v0_ = rhs66_
                        d_35_v3_ = rhs67_
                        lhs44_.f16 = rhs68_
                        lhs45_[lhs46_] = rhs69_
                    elif True:
                        d_360_v205_: D2
                        d_360_v205_ = D2_DC7(p1, d_33_v1_)
                        d_360_v205_ = d_360_v205_
                        d_361_v206_: _dafny.Array
                        nw64_ = _dafny.Array(int(0), 10)
                        d_361_v206_ = nw64_
                        index60_ = default__.safeIndex(610, (d_355_v200_).length(0))
                        (d_355_v200_)[index60_] = d_361_v206_
                        index61_ = default__.safeIndex(610, (d_355_v200_).length(0))
                        (d_355_v200_)[index61_] = d_361_v206_
                        d_362_v207_: _dafny.Map
                        d_362_v207_ = _dafny.Map({d_353_v199_: d_33_v1_})
                        d_32_v0_ = default__.fm4((0) - (p1), (d_353_v199_) not in ((d_362_v207_).set(d_353_v199_, d_33_v1_)), globalState)
                        d_32_v0_ = ((d_353_v199_).f17 if (p1) == ((default__.fm20(p3, globalState)).cardinality) else p2)
                        d_363_v208_: _dafny.MultiSet
                        d_363_v208_ = _dafny.MultiSet([p1])
                        index62_ = default__.safeIndex(113, (p0).length(0))
                        (p0)[index62_] = d_363_v208_
                        arr0_ = (d_355_v200_)[default__.safeIndex(610, (d_355_v200_).length(0))]
                        index63_ = default__.safeIndex(887, ((d_355_v200_)[default__.safeIndex(610, (d_355_v200_).length(0))]).length(0))
                        arr0_[index63_] = p1
                        d_364_v209_: _dafny.Array
                        nw65_ = _dafny.Array(None, 1)
                        nw65_[int(0)] = (d_353_v199_).f17
                        d_364_v209_ = nw65_
                        d_365_v210_: _dafny.Seq
                        d_365_v210_ = _dafny.SeqWithoutIsStrInference([(d_353_v199_).f17, False, (d_353_v199_).f17])
                        d_366_v211_: _dafny.Set
                        d_366_v211_ = _dafny.Set({len(d_365_v210_), p1})
                        index64_ = default__.safeIndex(649, (d_364_v209_).length(0))
                        (d_364_v209_)[index64_] = (_dafny.Set({p1})).ispropersubset(d_366_v211_)
                        d_367_v212_: _dafny.Seq
                        d_367_v212_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([p1])) - (d_363_v208_), d_363_v208_, d_363_v208_, (d_363_v208_) - (_dafny.MultiSet([p1, p1]))])
                        d_368_v213_: _dafny.Map
                        d_368_v213_ = _dafny.Map({p1: d_353_v199_})
                        d_369_v214_: _dafny.Map
                        d_369_v214_ = _dafny.Map({p1: d_368_v213_})
                        d_370_v215_: _dafny.Set
                        d_370_v215_ = _dafny.Set({False})
                        d_371_v216_: _dafny.Map
                        d_371_v216_ = _dafny.Map({p1: p1})
                        index65_ = default__.safeIndex(113, (p0).length(0))
                        arr1_ = (d_355_v200_)[default__.safeIndex(610, (d_355_v200_).length(0))]
                        index66_ = default__.safeIndex(887, ((d_355_v200_)[default__.safeIndex(610, (d_355_v200_).length(0))]).length(0))
                        index67_ = default__.safeIndex(649, (d_364_v209_).length(0))
                        rhs70_ = not(not((d_32_v0_) or ((d_353_v199_).f17)))
                        rhs71_ = (d_367_v212_)[default__.safeIndex(p1, len(d_367_v212_))]
                        rhs72_ = p1
                        rhs73_ = ((d_363_v208_)[len(d_369_v214_)] if (len(d_369_v214_)) in (d_363_v208_) else default__.fm0(len(d_370_v215_), d_371_v216_, globalState))
                        rhs74_ = False
                        lhs47_ = p0
                        lhs48_ = default__.safeIndex(113, (p0).length(0))
                        lhs49_ = (d_355_v200_)[default__.safeIndex(610, (d_355_v200_).length(0))]
                        lhs50_ = default__.safeIndex(887, ((d_355_v200_)[default__.safeIndex(610, (d_355_v200_).length(0))]).length(0))
                        lhs51_ = globalState
                        lhs52_ = d_364_v209_
                        lhs53_ = default__.safeIndex(649, (d_364_v209_).length(0))
                        d_32_v0_ = rhs70_
                        lhs47_[lhs48_] = rhs71_
                        lhs49_[lhs50_] = rhs72_
                        lhs51_.f6 = rhs73_
                        lhs52_[lhs53_] = rhs74_
                    d_372_v217_: _dafny.Array
                    nw66_ = _dafny.Array(int(0), 12)
                    d_372_v217_ = nw66_
                    d_373_v218_: _dafny.Seq
                    d_373_v218_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_374_v219_: _dafny.Seq
                    d_374_v219_ = _dafny.SeqWithoutIsStrInference([d_373_v218_])
                    d_375_v220_: _dafny.Map
                    d_375_v220_ = _dafny.Map({d_32_v0_: p1})
                    index68_ = default__.safeIndex(549, (d_372_v217_).length(0))
                    (d_372_v217_)[index68_] = len(((d_374_v219_)[default__.safeIndex(((d_375_v220_)[(d_353_v199_).f17] if ((d_353_v199_).f17) in (d_375_v220_) else p1), len(d_374_v219_))]) + (_dafny.SeqWithoutIsStrInference([p1])))
                    index69_ = default__.safeIndex(549, (d_372_v217_).length(0))
                    (d_372_v217_)[index69_] = (p1) + ((p1) + (p1))
                    pass
            pass
        d_376_v221_: _dafny.Map
        d_376_v221_ = _dafny.Map({p1: 415})
        d_377_v222_: _dafny.Seq
        d_377_v222_ = _dafny.SeqWithoutIsStrInference([p3])
        r0 = ((default__.fm0(p1, d_376_v221_, globalState)) + (786)) - (len((d_377_v222_) + (_dafny.SeqWithoutIsStrInference([p3]))))
        d_378_v223_: D1
        d_378_v223_ = D1_DC2(d_33_v1_)
        d_379_v224_: _dafny.Map
        d_379_v224_ = _dafny.Map({d_378_v223_: d_32_v0_})
        d_380_v225_: _dafny.MultiSet
        d_380_v225_ = _dafny.MultiSet([(d_377_v222_)[default__.safeIndex(p1, len(d_377_v222_))]])
        d_381_v226_: _dafny.Map
        d_381_v226_ = _dafny.Map({p3: p1})
        d_382_v227_: _dafny.Array
        nw67_ = _dafny.Array(None, 27)
        nw67_[int(0)] = len(_dafny.SeqWithoutIsStrInference([d_32_v0_, (d_353_v199_).f17, (d_353_v199_).f17, p2, d_32_v0_]))
        nw67_[int(1)] = (p1) + ((0) - (p1))
        nw67_[int(2)] = ((0) - ((0) - (p1)) if p3 else p1)
        nw67_[int(3)] = p1
        nw67_[int(4)] = p1
        nw67_[int(5)] = default__.safeDivisionInt(290, p1)
        nw67_[int(6)] = p1
        nw67_[int(7)] = 752
        nw67_[int(8)] = (p1 if p2 else len(d_379_v224_))
        nw67_[int(9)] = p1
        nw67_[int(10)] = (0) - (len(d_35_v3_))
        nw67_[int(11)] = p1
        nw67_[int(12)] = (64) + (len((_dafny.SeqWithoutIsStrInference([d_33_v1_ for d_383_i16_ in range(default__.abs(552))])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([d_33_v1_ for d_384_i16_ in range(default__.abs(552))]))), d_33_v1_)))
        nw67_[int(13)] = default__.safeDivisionInt((d_380_v225_).cardinality, len(_dafny.Map({((d_381_v226_)[p2] if (p2) in (d_381_v226_) else len(d_376_v221_)): d_353_v199_})))
        nw67_[int(14)] = p1
        nw67_[int(15)] = p1
        nw67_[int(16)] = p1
        nw67_[int(17)] = p1
        nw67_[int(18)] = (p1) + ((0) - (p1))
        nw67_[int(19)] = p1
        nw67_[int(20)] = p1
        nw67_[int(21)] = p1
        nw67_[int(22)] = p1
        nw67_[int(23)] = len(_dafny.SeqWithoutIsStrInference([p1 for d_385_i17_ in range(default__.abs(470))]))
        nw67_[int(24)] = p1
        nw67_[int(25)] = p1
        nw67_[int(26)] = p1
        d_382_v227_ = nw67_
        r1 = d_382_v227_
        d_386_v228_: _dafny.Set
        d_386_v228_ = _dafny.Set({((d_353_v199_).f17 if d_32_v0_ else d_32_v0_), (d_353_v199_).f17, p3})
        r2 = d_386_v228_
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_387_v1_: str
        d_387_v1_ = _dafny.CodePoint('d')
        d_388_v2_: int
        d_388_v2_ = 514
        d_389_v4_: bool
        d_389_v4_ = False
        d_390_v5_: _dafny.Seq
        d_390_v5_ = _dafny.SeqWithoutIsStrInference([d_389_v4_, True])
        d_391_v6_: _dafny.Map
        d_391_v6_ = _dafny.Map({len(d_390_v5_): True})
        d_392_v7_: _dafny.Set
        d_392_v7_ = _dafny.Set({((d_391_v6_)[d_388_v2_] if (d_388_v2_) in (d_391_v6_) else not(d_389_v4_)), d_389_v4_, d_389_v4_, d_389_v4_})
        d_393_v8_: _dafny.Seq
        d_393_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frrp"))
        d_394_v9_: _dafny.Array
        def lambda14_(d_395_i0_):
            return default__.safeDivisionInt(d_395_i0_, 535)

        init7_ = lambda14_
        nw68_ = _dafny.Array(None, 7)
        for i0_7_ in range(nw68_.length(0)):
            nw68_[i0_7_] = init7_(i0_7_)
        d_394_v9_ = nw68_
        d_396_globalState_: GlobalState
        nw69_ = GlobalState()
        def iife58_():
            coll26_ = _dafny.Map()
            compr_26_: str
            for compr_26_ in (_dafny.SeqWithoutIsStrInference([d_387_v1_])).Elements:
                d_397_v0_: str = compr_26_
                if (d_397_v0_) in (_dafny.SeqWithoutIsStrInference([d_387_v1_])):
                    coll26_[d_397_v0_] = 68
            return _dafny.Map(coll26_)
        def iife59_():
            coll27_ = _dafny.Set()
            compr_27_: int
            for compr_27_ in (_dafny.SeqWithoutIsStrInference([d_388_v2_])).Elements:
                d_398_v3_: int = compr_27_
                if (d_398_v3_) in (_dafny.SeqWithoutIsStrInference([d_388_v2_])):
                    coll27_ = coll27_.union(_dafny.Set([(d_398_v3_) * ((0) - (-233))]))
            return _dafny.Set(coll27_)
        nw69_.ctor__(455, 119, (iife58_()
        ).set(d_387_v1_, d_388_v2_), True, 936, 277, 911, _dafny.MultiSet([d_388_v2_, (0) - (len(iife59_()
        )), len(d_392_v7_), len(d_393_v8_)]), True, 828, 622, -572, 262, False, d_394_v9_, 726, 237)
        d_396_globalState_ = nw69_
        d_399_v10_: _dafny.Map
        d_399_v10_ = _dafny.Map({d_388_v2_: d_388_v2_})
        hi0_ = d_388_v2_
        for d_400_i1_ in range(default__.fm0(d_388_v2_, (d_399_v10_).set(-923, d_388_v2_), d_396_globalState_), hi0_):
            (d_396_globalState_).f16 = d_400_i1_
            d_401_v11_: _dafny.MultiSet
            d_401_v11_ = _dafny.MultiSet([not(d_389_v4_), d_389_v4_, d_389_v4_])
            d_402_v12_: _dafny.MultiSet
            d_402_v12_ = _dafny.MultiSet([d_401_v11_])
            d_403_v13_: _dafny.Seq
            d_403_v13_ = _dafny.SeqWithoutIsStrInference([((d_402_v12_)[d_401_v11_] if (d_401_v11_) in (d_402_v12_) else d_400_i1_)])
            d_403_v13_ = d_403_v13_
            d_404_v14_: _dafny.Map
            d_404_v14_ = _dafny.Map({d_389_v4_: (d_391_v6_) | (d_391_v6_)})
            d_404_v14_ = d_404_v14_
            d_405_v16_: D0
            def iife60_():
                coll28_ = _dafny.Map()
                compr_28_: int
                for compr_28_ in _dafny.IntegerRange(396, -156):
                    d_406_v15_: int = compr_28_
                    if ((396) <= (d_406_v15_)) and ((d_406_v15_) < (-156)):
                        coll28_[(d_406_v15_) - (len(d_393_v8_))] = d_391_v6_
                return _dafny.Map(coll28_)
            d_405_v16_ = D0_DC0(d_400_i1_, d_389_v4_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yicqicecd")), iife60_()
)
            source6_ = d_405_v16_
            d_407___mcc_h0_ = source6_.cf0
            d_408___mcc_h1_ = source6_.cf1
            d_409___mcc_h2_ = source6_.cf2
            d_410___mcc_h3_ = source6_.cf3
            d_411_cf3_ = d_410___mcc_h3_
            d_412_cf2_ = d_409___mcc_h2_
            d_413_cf1_ = d_408___mcc_h1_
            d_414_cf0_ = d_407___mcc_h0_
            index70_ = default__.safeIndex(540, (d_394_v9_).length(0))
            (d_394_v9_)[index70_] = (0) - ((0) - ((0) - ((default__.fm0(d_414_cf0_, d_399_v10_, d_396_globalState_)) * (d_400_i1_))))
            index71_ = default__.safeIndex(540, (d_394_v9_).length(0))
            (d_394_v9_)[index71_] = d_388_v2_
            d_399_v10_ = (d_399_v10_).set((0) - ((d_394_v9_)[default__.safeIndex(540, (d_394_v9_).length(0))]), (d_394_v9_)[default__.safeIndex(540, (d_394_v9_).length(0))])
            d_415_v17_: D1
            d_415_v17_ = D1_DC1(d_390_v5_)
            d_416_v18_: _dafny.MultiSet
            d_416_v18_ = _dafny.MultiSet([d_414_cf0_])
            d_417_v19_: _dafny.Map
            d_417_v19_ = _dafny.Map({d_413_cf1_: d_389_v4_})
            d_418_v20_: _dafny.Map
            d_418_v20_ = _dafny.Map({(0) - ((d_394_v9_)[default__.safeIndex(540, (d_394_v9_).length(0))]): d_390_v5_})
            d_419_v21_: _dafny.Array
            nw70_ = _dafny.Array(None, 25)
            nw70_[int(0)] = d_390_v5_
            nw70_[int(1)] = d_390_v5_
            nw70_[int(2)] = ((d_415_v17_).cf4) + (d_390_v5_)
            nw70_[int(3)] = d_390_v5_
            nw70_[int(4)] = d_390_v5_
            nw70_[int(5)] = d_390_v5_
            nw70_[int(6)] = d_390_v5_
            nw70_[int(7)] = (d_390_v5_).set(default__.safeIndex((0) - (d_414_cf0_), len(d_390_v5_)), d_413_cf1_)
            nw70_[int(8)] = (((d_390_v5_).set(default__.safeIndex((d_416_v18_).cardinality, len(d_390_v5_)), d_389_v4_)).set(default__.safeIndex(d_400_i1_, len((d_390_v5_).set(default__.safeIndex((d_416_v18_).cardinality, len(d_390_v5_)), d_389_v4_))), d_413_cf1_)) + (_dafny.SeqWithoutIsStrInference([False, ((d_417_v19_)[d_413_cf1_] if (d_413_cf1_) in (d_417_v19_) else d_389_v4_)]))
            nw70_[int(9)] = d_390_v5_
            nw70_[int(10)] = d_390_v5_
            nw70_[int(11)] = d_390_v5_
            nw70_[int(12)] = d_390_v5_
            nw70_[int(13)] = d_390_v5_
            nw70_[int(14)] = ((d_418_v20_)[-371] if (-371) in (d_418_v20_) else _dafny.SeqWithoutIsStrInference([d_413_cf1_]))
            nw70_[int(15)] = d_390_v5_
            nw70_[int(16)] = (d_390_v5_) + (d_390_v5_)
            nw70_[int(17)] = _dafny.SeqWithoutIsStrInference([False])
            nw70_[int(18)] = (d_390_v5_) + (d_390_v5_)
            nw70_[int(19)] = d_390_v5_
            nw70_[int(20)] = (d_390_v5_) + ((D1_DC1(d_390_v5_)).cf4)
            nw70_[int(21)] = d_390_v5_
            nw70_[int(22)] = d_390_v5_
            nw70_[int(23)] = d_390_v5_
            nw70_[int(24)] = d_390_v5_
            d_419_v21_ = nw70_
            index72_ = default__.safeIndex(128, (d_419_v21_).length(0))
            (d_419_v21_)[index72_] = (_dafny.SeqWithoutIsStrInference([d_389_v4_, False])) + (d_390_v5_)
            d_420_v22_: _dafny.MultiSet
            d_420_v22_ = _dafny.MultiSet([d_393_v8_])
            d_421_v23_: _dafny.Array
            def lambda15_(d_422_i1_, d_423_cf1_, d_424_v8_, d_425_cf3_):
                def lambda16_(d_426_i2_):
                    return D0_DC0(d_422_i1_, d_423_cf1_, d_424_v8_, d_425_cf3_)

                return lambda16_

            init8_ = lambda15_(d_400_i1_, d_413_cf1_, d_393_v8_, d_411_cf3_)
            nw71_ = _dafny.Array(None, 28)
            for i0_8_ in range(nw71_.length(0)):
                nw71_[i0_8_] = init8_(i0_8_)
            d_421_v23_ = nw71_
            d_427_v24_: _dafny.Seq
            d_427_v24_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_389_v4_})])
            d_428_v25_: _dafny.Array
            nw72_ = _dafny.Array(int(0), 17)
            d_428_v25_ = nw72_
            d_429_v26_: _dafny.Set
            d_429_v26_ = _dafny.Set({d_428_v25_})
            index73_ = default__.safeIndex(128, (d_419_v21_).length(0))
            rhs75_ = (d_390_v5_) + (d_390_v5_)
            rhs76_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jaeue"))])
            rhs77_ = d_421_v23_
            rhs78_ = _dafny.SeqWithoutIsStrInference([d_392_v7_ for d_430_i3_ in range(default__.abs(203))])
            rhs79_ = default__.fm0(d_388_v2_, (default__.fm1(d_399_v10_, len(d_399_v10_), len(d_429_v26_), d_396_globalState_)) | (d_399_v10_), d_396_globalState_)
            lhs54_ = d_419_v21_
            lhs55_ = default__.safeIndex(128, (d_419_v21_).length(0))
            lhs56_ = d_396_globalState_
            lhs54_[lhs55_] = rhs75_
            d_420_v22_ = rhs76_
            d_421_v23_ = rhs77_
            d_427_v24_ = rhs78_
            lhs56_.f6 = rhs79_
            d_431_v27_: _dafny.MultiSet
            d_431_v27_ = _dafny.MultiSet([d_428_v25_])
            d_412_cf2_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flvuwul"))).set(default__.safeIndex(d_414_cf0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flvuwul")))), (d_387_v1_ if d_389_v4_ else (d_393_v8_)[default__.safeIndex(d_414_cf0_, len(d_393_v8_))]))).set(default__.safeIndex(default__.safeModuloInt(240, (d_431_v27_).cardinality), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flvuwul"))).set(default__.safeIndex(d_414_cf0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flvuwul")))), (d_387_v1_ if d_389_v4_ else (d_393_v8_)[default__.safeIndex(d_414_cf0_, len(d_393_v8_))])))), d_387_v1_)
        index74_ = default__.safeIndex(12, (d_394_v9_).length(0))
        (d_394_v9_)[index74_] = d_388_v2_
        index75_ = default__.safeIndex(12, (d_394_v9_).length(0))
        (d_394_v9_)[index75_] = default__.safeModuloInt(d_388_v2_, d_388_v2_)
        d_432_v28_: C0
        nw73_ = C0()
        nw73_.ctor__(d_389_v4_)
        d_432_v28_ = nw73_
        d_433_v29_: _dafny.MultiSet
        d_433_v29_ = _dafny.MultiSet([d_388_v2_])
        d_434_v30_: _dafny.Map
        d_434_v30_ = _dafny.Map({(d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]: _dafny.MultiSet([76])})
        d_435_v31_: D1
        d_435_v31_ = D1_DC4(d_388_v2_, (d_432_v28_).f17, d_434_v30_)
        d_436_v32_: D1
        d_436_v32_ = D1_DC5(d_435_v31_)
        pat_let_tv16_ = d_389_v4_
        pat_let_tv17_ = d_432_v28_
        pat_let_tv18_ = d_389_v4_
        pat_let_tv19_ = d_432_v28_
        pat_let_tv20_ = d_389_v4_
        def lambda17_(source7_):
            if source7_.is_DC2:
                d_437___mcc_h4_ = source7_.cf5
                d_438_cf5_ = d_437___mcc_h4_
                return True
            elif source7_.is_DC3:
                d_439___mcc_h5_ = source7_.cf6
                d_440___mcc_h6_ = source7_.cf7
                d_441___mcc_h7_ = source7_.cf8
                d_442_cf8_ = d_441___mcc_h7_
                d_443_cf7_ = d_440___mcc_h6_
                d_444_cf6_ = d_439___mcc_h5_
                return pat_let_tv16_
            elif source7_.is_DC4:
                d_445___mcc_h8_ = source7_.cf9
                d_446___mcc_h9_ = source7_.cf10
                d_447___mcc_h10_ = source7_.cf11
                d_448_cf11_ = d_447___mcc_h10_
                d_449_cf10_ = d_446___mcc_h9_
                d_450_cf9_ = d_445___mcc_h8_
                return (pat_let_tv17_).f17
            elif source7_.is_DC1:
                d_451___mcc_h11_ = source7_.cf4
                d_452_cf4_ = d_451___mcc_h11_
                return pat_let_tv18_
            elif True:
                d_453___mcc_h12_ = source7_.cf12
                d_454_cf12_ = d_453___mcc_h12_
                return not ((pat_let_tv19_).f17) or (pat_let_tv20_)

        rhs80_ = default__.fm3(d_389_v4_, d_396_globalState_)
        rhs81_ = not (d_389_v4_) or ((d_433_v29_).ispropersubset(_dafny.MultiSet([187, len(d_399_v10_)])))
        rhs82_ = lambda17_(d_436_v32_)
        rhs83_ = (d_387_v1_ if d_389_v4_ else d_387_v1_)
        rhs84_ = d_387_v1_
        d_387_v1_ = rhs80_
        d_389_v4_ = rhs81_
        d_389_v4_ = rhs82_
        d_387_v1_ = rhs83_
        d_387_v1_ = rhs84_
        d_455_v33_: _dafny.MultiSet
        d_455_v33_ = _dafny.MultiSet([False])
        d_456_v34_: _dafny.Array
        nw74_ = _dafny.Array(False, 6)
        d_456_v34_ = nw74_
        index76_ = default__.safeIndex(701, (d_456_v34_).length(0))
        (d_456_v34_)[index76_] = d_389_v4_
        index77_ = default__.safeIndex(168, (d_456_v34_).length(0))
        (d_456_v34_)[index77_] = (d_432_v28_).f17
        d_457_v35_: _dafny.Map
        d_457_v35_ = _dafny.Map({(d_432_v28_).f17: (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]})
        d_458_v36_: _dafny.Seq
        d_458_v36_ = _dafny.SeqWithoutIsStrInference([d_399_v10_, d_399_v10_])
        index78_ = default__.safeIndex(701, (d_456_v34_).length(0))
        index79_ = default__.safeIndex(168, (d_456_v34_).length(0))
        rhs85_ = d_455_v33_
        rhs86_ = default__.fm4(((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]) + ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]), ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]) >= ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]), d_396_globalState_)
        rhs87_ = default__.fm4(len((d_393_v8_).set(default__.safeIndex(((d_457_v35_)[default__.fm4((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], (d_432_v28_).f17, d_396_globalState_)] if (default__.fm4((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], (d_432_v28_).f17, d_396_globalState_)) in (d_457_v35_) else default__.fm0(16, (d_458_v36_)[default__.safeIndex((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], len(d_458_v36_))], d_396_globalState_)), len(d_393_v8_)), d_387_v1_)), (d_432_v28_).f17, d_396_globalState_)
        rhs88_ = default__.fm4(d_388_v2_, (d_432_v28_).f17, d_396_globalState_)
        lhs57_ = d_456_v34_
        lhs58_ = default__.safeIndex(701, (d_456_v34_).length(0))
        lhs59_ = d_456_v34_
        lhs60_ = default__.safeIndex(168, (d_456_v34_).length(0))
        d_455_v33_ = rhs85_
        d_389_v4_ = rhs86_
        lhs57_[lhs58_] = rhs87_
        lhs59_[lhs60_] = rhs88_
        d_459_v37_: C0
        nw75_ = C0()
        nw75_.ctor__((d_456_v34_)[default__.safeIndex(701, (d_456_v34_).length(0))])
        d_459_v37_ = nw75_
        d_460_v38_: _dafny.Map
        d_460_v38_ = _dafny.Map({(d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]: d_393_v8_})
        d_460_v38_ = (d_460_v38_).set(len(d_390_v5_), d_393_v8_)
        index80_ = default__.safeIndex(701, (d_456_v34_).length(0))
        (d_456_v34_)[index80_] = ((d_393_v8_ if False else (d_393_v8_).set(default__.safeIndex(d_388_v2_, len(d_393_v8_)), d_387_v1_))) <= (d_393_v8_)
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_394_v9_).length(0)):
            d_461_i4_: int = guard_loop_0_
            if (True) and (((0) <= (d_461_i4_)) and ((d_461_i4_) < ((d_394_v9_).length(0)))):
                (d_394_v9_)[(d_461_i4_)] = (d_461_i4_) * (671)
        d_462_v39_: _dafny.Array
        nw76_ = _dafny.Array(_dafny.Seq({}), 14)
        d_462_v39_ = nw76_
        _ingredientsd_0 = []
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_462_v39_).length(0)):
            d_463_i5_: int = guard_loop_1_
            if (True) and (((0) <= (d_463_i5_)) and ((d_463_i5_) < ((d_462_v39_).length(0)))):
                _ingredientsd_0.append((d_462_v39_, int(d_463_i5_), (_dafny.SeqWithoutIsStrInference([(d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]])) + ((_dafny.SeqWithoutIsStrInference([d_388_v2_])) + (_dafny.SeqWithoutIsStrInference([129])))))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        (d_396_globalState_).f16 = d_388_v2_
        d_464_i6_: int
        d_464_i6_ = 0
        with _dafny.label("1"):
            while (d_459_v37_).f17:
                with _dafny.c_label("1"):
                    if (d_464_i6_) >= (100):
                        raise _dafny.Break("1")
                    d_464_i6_ = (d_464_i6_) + (1)
                    d_465_v40_: _dafny.Map
                    d_465_v40_ = _dafny.Map({D1_DC4(861, (d_456_v34_)[default__.safeIndex(701, (d_456_v34_).length(0))], d_434_v30_): d_394_v9_})
                    d_466_v41_: D1
                    d_466_v41_ = D1_DC4(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksjsyh"))), (d_432_v28_).f17, _dafny.Map({(d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]: d_433_v29_}))
                    d_467_v42_: _dafny.MultiSet
                    d_467_v42_ = _dafny.MultiSet([d_394_v9_, d_394_v9_, ((d_465_v40_)[d_466_v41_] if (d_466_v41_) in (d_465_v40_) else d_394_v9_), d_394_v9_, d_394_v9_])
                    d_467_v42_ = d_467_v42_
                    d_468_v43_: _dafny.Seq
                    d_468_v43_ = _dafny.SeqWithoutIsStrInference([d_456_v34_, d_456_v34_, d_456_v34_])
                    index81_ = default__.safeIndex(701, (d_456_v34_).length(0))
                    (d_456_v34_)[index81_] = not((d_468_v43_) <= (d_468_v43_))
                    d_469_v44_: _dafny.Array
                    nw77_ = _dafny.Array(None, 8)
                    nw77_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqsutwajj"))
                    nw77_[int(1)] = d_393_v8_
                    nw77_[int(2)] = d_393_v8_
                    nw77_[int(3)] = d_393_v8_
                    nw77_[int(4)] = d_393_v8_
                    nw77_[int(5)] = d_393_v8_
                    nw77_[int(6)] = (d_393_v8_).set(default__.safeIndex(len(_dafny.Map({(d_456_v34_)[default__.safeIndex(701, (d_456_v34_).length(0))]: d_388_v2_})), len(d_393_v8_)), d_387_v1_)
                    nw77_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibgqhjni"))
                    d_469_v44_ = nw77_
                    d_470_v45_: _dafny.Map
                    d_470_v45_ = _dafny.Map({(d_432_v28_).f17: d_469_v44_})
                    d_470_v45_ = (d_470_v45_).set(True, d_469_v44_)
                    d_471_v47_: C0
                    nw78_ = C0()
                    def iife61_():
                        coll29_ = _dafny.Map()
                        compr_29_: int
                        for compr_29_ in (default__.fm5(d_396_globalState_)).keys.Elements:
                            d_472_v46_: int = compr_29_
                            if (d_472_v46_) in (default__.fm5(d_396_globalState_)):
                                coll29_[(d_472_v46_) - (d_388_v2_)] = (d_456_v34_)[default__.safeIndex(701, (d_456_v34_).length(0))]
                        return _dafny.Map(coll29_)
                    nw78_.ctor__((d_391_v6_) != (iife61_()
                    ))
                    d_471_v47_ = nw78_
                    pass
            pass
        index82_ = default__.safeIndex(701, (d_456_v34_).length(0))
        (d_456_v34_)[index82_] = (d_459_v37_).f17
        hi1_ = (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]
        for d_473_i7_ in range(85, hi1_):
            d_474_v48_: C0
            nw79_ = C0()
            nw79_.ctor__((d_459_v37_).f17)
            d_474_v48_ = nw79_
            d_475_v49_: _dafny.Seq
            d_475_v49_ = _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([d_393_v8_ for d_476_i8_ in range(default__.abs(874))])))])
            index83_ = default__.safeIndex(569, (d_462_v39_).length(0))
            (d_462_v39_)[index83_] = d_475_v49_
            index84_ = default__.safeIndex(569, (d_462_v39_).length(0))
            (d_462_v39_)[index84_] = d_475_v49_
            d_477_v50_: _dafny.Map
            d_477_v50_ = _dafny.Map({d_388_v2_: d_391_v6_})
            d_478_v51_: D0
            d_478_v51_ = D0_DC0(863, (d_459_v37_).f17, d_393_v8_, d_477_v50_)
            d_479_v52_: D1
            d_479_v52_ = D1_DC3(d_399_v10_, d_388_v2_, d_478_v51_)
            d_480_v53_: _dafny.Set
            d_480_v53_ = _dafny.Set({d_479_v52_, D1_DC3(d_399_v10_, d_473_i7_, d_478_v51_), d_479_v52_})
            index85_ = default__.safeIndex(12, (d_394_v9_).length(0))
            (d_394_v9_)[index85_] = default__.safeModuloInt(d_473_i7_, len((d_480_v53_).intersection(_dafny.Set({d_479_v52_}))))
            d_457_v35_ = (d_457_v35_).set((d_456_v34_)[default__.safeIndex(701, (d_456_v34_).length(0))], (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))])
        if (d_393_v8_) <= (d_393_v8_):
            d_388_v2_ = (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]
            d_481_v54_: D1
            d_481_v54_ = D1_DC2(d_387_v1_)
            d_482_v55_: D1
            d_482_v55_ = D1_DC4(d_388_v2_, (d_456_v34_)[default__.safeIndex(701, (d_456_v34_).length(0))], d_434_v30_)
            d_483_v56_: _dafny.Map
            d_483_v56_ = _dafny.Map({d_432_v28_: d_455_v33_})
            d_484_v58_: _dafny.Set
            def iife62_():
                coll30_ = _dafny.Map()
                compr_30_: int
                for compr_30_ in _dafny.IntegerRange(-794, 807):
                    d_485_v57_: int = compr_30_
                    if ((-794) <= (d_485_v57_)) and ((d_485_v57_) < (807)):
                        coll30_[default__.safeDivisionInt(d_485_v57_, d_388_v2_)] = d_387_v1_
                return _dafny.Map(coll30_)
            d_484_v58_ = _dafny.Set({len(iife62_()
            )})
            d_486_v59_: _dafny.Array
            nw80_ = _dafny.Array(None, 20)
            nw80_[int(0)] = default__.fm6(False, d_389_v4_, (d_432_v28_).f17, d_396_globalState_)
            nw80_[int(1)] = _dafny.MultiSet([(d_482_v55_).cf10, (d_432_v28_).f17, False, (d_432_v28_).f17, True])
            nw80_[int(2)] = d_455_v33_
            nw80_[int(3)] = d_455_v33_
            nw80_[int(4)] = _dafny.MultiSet((default__.fm7((d_481_v54_).cf5, d_396_globalState_)) + (d_390_v5_))
            nw80_[int(5)] = d_455_v33_
            nw80_[int(6)] = default__.fm6(True, (d_456_v34_)[default__.safeIndex(701, (d_456_v34_).length(0))], (d_459_v37_).f17, d_396_globalState_)
            nw80_[int(7)] = d_455_v33_
            nw80_[int(8)] = d_455_v33_
            nw80_[int(9)] = d_455_v33_
            nw80_[int(10)] = d_455_v33_
            nw80_[int(11)] = d_455_v33_
            nw80_[int(12)] = d_455_v33_
            nw80_[int(13)] = (d_455_v33_).set((d_459_v37_).f17, default__.abs(d_388_v2_))
            nw80_[int(14)] = ((d_483_v56_)[d_459_v37_] if (d_459_v37_) in (d_483_v56_) else (d_455_v33_).set((d_432_v28_).f17, default__.abs(len(d_484_v58_))))
            nw80_[int(15)] = _dafny.MultiSet([True, (d_459_v37_).f17])
            nw80_[int(16)] = d_455_v33_
            nw80_[int(17)] = _dafny.MultiSet([(d_456_v34_)[default__.safeIndex(701, (d_456_v34_).length(0))], (d_432_v28_).f17])
            nw80_[int(18)] = d_455_v33_
            nw80_[int(19)] = d_455_v33_
            d_486_v59_ = nw80_
            index86_ = default__.safeIndex(406, (d_486_v59_).length(0))
            (d_486_v59_)[index86_] = d_455_v33_
            index87_ = default__.safeIndex(406, (d_486_v59_).length(0))
            rhs89_ = d_481_v54_
            rhs90_ = (_dafny.MultiSet([(d_459_v37_).f17])) | (d_455_v33_)
            lhs61_ = d_486_v59_
            lhs62_ = default__.safeIndex(406, (d_486_v59_).length(0))
            d_481_v54_ = rhs89_
            lhs61_[lhs62_] = rhs90_
            d_487_v60_: C0
            nw81_ = C0()
            nw81_.ctor__((not((d_459_v37_).f17) if d_389_v4_ else (d_459_v37_).f17))
            d_487_v60_ = nw81_
            d_488_v61_: _dafny.Set
            d_488_v61_ = _dafny.Set({d_390_v5_})
            if ((d_488_v61_).intersection(d_488_v61_)).ispropersubset(d_488_v61_):
                d_489_v62_: _dafny.Array
                def lambda18_(d_490_i9_):
                    return _dafny.CodePoint('o')

                init9_ = lambda18_
                nw82_ = _dafny.Array(None, 25)
                for i0_9_ in range(nw82_.length(0)):
                    nw82_[i0_9_] = init9_(i0_9_)
                d_489_v62_ = nw82_
                rhs91_ = d_489_v62_
                rhs92_ = (d_487_v60_).f17
                d_489_v62_ = rhs91_
                d_389_v4_ = rhs92_
                d_491_v63_: _dafny.Array
                nw83_ = _dafny.Array(None, 15)
                d_491_v63_ = nw83_
                d_492_v64_: _dafny.Set
                d_492_v64_ = _dafny.Set({d_491_v63_})
                index88_ = default__.safeIndex(701, (d_456_v34_).length(0))
                rhs93_ = d_394_v9_
                rhs94_ = (d_492_v64_).ispropersubset(_dafny.Set({d_491_v63_}))
                rhs95_ = default__.fm4((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], (d_459_v37_).f17, d_396_globalState_)
                rhs96_ = d_432_v28_
                lhs63_ = d_456_v34_
                lhs64_ = default__.safeIndex(701, (d_456_v34_).length(0))
                d_394_v9_ = rhs93_
                lhs63_[lhs64_] = rhs94_
                d_389_v4_ = rhs95_
                d_432_v28_ = rhs96_
                d_493_v65_: _dafny.Map
                d_493_v65_ = _dafny.Map({d_388_v2_: d_391_v6_})
                d_494_v66_: D0
                d_494_v66_ = D0_DC0(d_388_v2_, (d_487_v60_).f17, d_393_v8_, d_493_v65_)
                d_495_v67_: D1
                d_495_v67_ = D1_DC3(d_399_v10_, (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], d_494_v66_)
                d_495_v67_ = default__.fm8((d_432_v28_).f17, (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], d_396_globalState_)
                d_496_v68_: _dafny.Map
                d_496_v68_ = _dafny.Map({(d_459_v37_).f17: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyk"))})
                d_496_v68_ = (d_496_v68_).set((d_459_v37_).f17, d_393_v8_)
                d_497_v69_: _dafny.Seq
                d_497_v69_ = _dafny.SeqWithoutIsStrInference([D1_DC3((d_399_v10_).set(d_388_v2_, (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]), (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], d_494_v66_), d_495_v67_, D1_DC3(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_484_v58_, d_484_v58_])): 214}), d_388_v2_, D0_DC0(len(d_392_v7_), d_389_v4_, _dafny.SeqWithoutIsStrInference([d_387_v1_ for d_498_i10_ in range(default__.abs(590))]), d_493_v65_))])
                d_499_v70_: _dafny.MultiSet
                d_499_v70_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_495_v67_, d_495_v67_]), d_497_v69_, d_497_v69_, d_497_v69_, default__.fm9(d_396_globalState_)])
                index89_ = default__.safeIndex(701, (d_456_v34_).length(0))
                (d_456_v34_)[index89_] = (d_499_v70_).ispropersubset(d_499_v70_)
            elif True:
                d_500_v71_: _dafny.Seq
                d_500_v71_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], d_388_v2_), d_388_v2_])
                index90_ = default__.safeIndex(12, (d_394_v9_).length(0))
                (d_394_v9_)[index90_] = (0) - ((d_500_v71_)[default__.safeIndex(((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]) + ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]), len(d_500_v71_))])
                d_501_v72_: C0
                nw84_ = C0()
                nw84_.ctor__((d_392_v7_).isdisjoint(d_392_v7_))
                d_501_v72_ = nw84_
                d_502_v73_: _dafny.Map
                d_502_v73_ = _dafny.Map({-377: d_391_v6_})
                d_503_v74_: D0
                d_503_v74_ = D0_DC0(len(d_393_v8_), (d_459_v37_).f17, d_393_v8_, d_502_v73_)
                d_504_v75_: D1
                d_504_v75_ = D1_DC3(d_399_v10_, ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]) * (d_388_v2_), d_503_v74_)
                pat_let_tv21_ = d_500_v71_
                pat_let_tv22_ = d_500_v71_
                pat_let_tv23_ = d_399_v10_
                pat_let_tv24_ = d_394_v9_
                pat_let_tv25_ = d_394_v9_
                def iife63_(_pat_let16_0):
                    def iife64_(d_505_dt__update__tmp_h0_):
                        def iife66_():
                            coll31_ = _dafny.Map()
                            compr_31_: int
                            for compr_31_ in (pat_let_tv21_).Elements:
                                d_506_v76_: int = compr_31_
                                if (d_506_v76_) in (pat_let_tv22_):
                                    coll31_[(d_506_v76_) * ((pat_let_tv25_)[default__.safeIndex(12, (pat_let_tv24_).length(0))])] = len(pat_let_tv23_)
                            return _dafny.Map(coll31_)
                        def iife65_(_pat_let17_0):
                            def iife67_(d_507_dt__update_hcf6_h0_):
                                return D1_DC3(d_507_dt__update_hcf6_h0_, (d_505_dt__update__tmp_h0_).cf7, (d_505_dt__update__tmp_h0_).cf8)
                            return iife67_(_pat_let17_0)
                        return iife65_(iife66_()
                        )
                    return iife64_(_pat_let16_0)
                d_504_v75_ = iife63_(d_504_v75_)
                d_508_v77_: _dafny.Map
                d_508_v77_ = _dafny.Map({(d_432_v28_).f17: (d_459_v37_).f17})
                d_509_v78_: _dafny.Map
                d_509_v78_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fh")): ((d_508_v77_)[(d_459_v37_).f17] if ((d_459_v37_).f17) in (d_508_v77_) else default__.fm4(d_388_v2_, default__.fm4(d_388_v2_, (d_501_v72_).f17, d_396_globalState_), d_396_globalState_))})
                index91_ = default__.safeIndex(701, (d_456_v34_).length(0))
                (d_456_v34_)[index91_] = ((d_509_v78_)[(d_393_v8_) + (d_393_v8_)] if ((d_393_v8_) + (d_393_v8_)) in (d_509_v78_) else ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]) <= (len(d_390_v5_)))
                (d_396_globalState_).f16 = (0) - ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))])
            index92_ = default__.safeIndex(12, (d_394_v9_).length(0))
            (d_394_v9_)[index92_] = ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))] if True else d_388_v2_)
        elif True:
            d_394_v9_ = d_394_v9_
            if (d_392_v7_).ispropersubset(d_392_v7_):
                (d_396_globalState_).f16 = (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]
                d_389_v4_ = (d_432_v28_).f17
                d_510_v79_: _dafny.Array
                nw85_ = _dafny.Array(None, 21)
                d_510_v79_ = nw85_
                index93_ = default__.safeIndex(643, (d_510_v79_).length(0))
                (d_510_v79_)[index93_] = d_432_v28_
                index94_ = default__.safeIndex(643, (d_510_v79_).length(0))
                (d_510_v79_)[index94_] = d_432_v28_
                d_388_v2_ = default__.safeDivisionInt((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], ((d_457_v35_)[default__.fm4((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], (d_432_v28_).f17, d_396_globalState_)] if (default__.fm4((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], (d_432_v28_).f17, d_396_globalState_)) in (d_457_v35_) else default__.fm0((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], d_399_v10_, d_396_globalState_)))
                (d_396_globalState_).f6 = default__.safeModuloInt((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], (0) - (len(d_390_v5_)))
            elif True:
                index95_ = default__.safeIndex(701, (d_456_v34_).length(0))
                (d_456_v34_)[index95_] = (d_432_v28_).f17
                d_387_v1_ = d_387_v1_
                d_511_v80_: _dafny.Array
                nw86_ = _dafny.Array(_dafny.MultiSet({}), 7)
                d_511_v80_ = nw86_
                index96_ = default__.safeIndex(686, (d_511_v80_).length(0))
                (d_511_v80_)[index96_] = d_433_v29_
                index97_ = default__.safeIndex(686, (d_511_v80_).length(0))
                rhs97_ = (d_393_v8_).set(default__.safeIndex((-462) - (d_388_v2_), len(d_393_v8_)), d_387_v1_)
                rhs98_ = default__.safeModuloInt(len(_dafny.Map({d_394_v9_: d_387_v1_})), default__.safeModuloInt(d_388_v2_, (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]))
                rhs99_ = _dafny.MultiSet([(0) - (d_388_v2_)])
                rhs100_ = d_432_v28_
                rhs101_ = d_394_v9_
                lhs65_ = d_511_v80_
                lhs66_ = default__.safeIndex(686, (d_511_v80_).length(0))
                lhs67_ = d_396_globalState_
                d_393_v8_ = rhs97_
                d_388_v2_ = rhs98_
                lhs65_[lhs66_] = rhs99_
                d_432_v28_ = rhs100_
                lhs67_.f14 = rhs101_
                index98_ = default__.safeIndex(701, (d_456_v34_).length(0))
                (d_456_v34_)[index98_] = (d_459_v37_).f17
                index99_ = default__.safeIndex(12, (d_394_v9_).length(0))
                (d_394_v9_)[index99_] = d_388_v2_
            if (d_390_v5_)[default__.safeIndex(((d_399_v10_)[(d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]] if ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]) in (d_399_v10_) else d_388_v2_), len(d_390_v5_))]:
                d_512_v81_: _dafny.Map
                d_512_v81_ = _dafny.Map({d_393_v8_: d_433_v29_})
                d_513_v82_: _dafny.Array
                nw87_ = _dafny.Array(None, 27)
                nw87_[int(0)] = d_433_v29_
                nw87_[int(1)] = _dafny.MultiSet([d_388_v2_])
                nw87_[int(2)] = _dafny.MultiSet([(d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]])
                nw87_[int(3)] = d_433_v29_
                nw87_[int(4)] = d_433_v29_
                nw87_[int(5)] = d_433_v29_
                nw87_[int(6)] = d_433_v29_
                nw87_[int(7)] = d_433_v29_
                nw87_[int(8)] = d_433_v29_
                nw87_[int(9)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]]))
                nw87_[int(10)] = d_433_v29_
                nw87_[int(11)] = d_433_v29_
                nw87_[int(12)] = d_433_v29_
                nw87_[int(13)] = d_433_v29_
                nw87_[int(14)] = d_433_v29_
                nw87_[int(15)] = ((d_512_v81_)[d_393_v8_] if (d_393_v8_) in (d_512_v81_) else d_433_v29_)
                nw87_[int(16)] = d_433_v29_
                nw87_[int(17)] = d_433_v29_
                nw87_[int(18)] = d_433_v29_
                nw87_[int(19)] = d_433_v29_
                nw87_[int(20)] = d_433_v29_
                nw87_[int(21)] = d_433_v29_
                nw87_[int(22)] = d_433_v29_
                nw87_[int(23)] = ((d_434_v30_)[(d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]] if ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]) in (d_434_v30_) else d_433_v29_)
                nw87_[int(24)] = d_433_v29_
                nw87_[int(25)] = d_433_v29_
                nw87_[int(26)] = d_433_v29_
                d_513_v82_ = nw87_
                d_514_v83_: _dafny.Map
                d_514_v83_ = _dafny.Map({d_388_v2_: _dafny.Map({d_388_v2_: d_389_v4_})})
                d_515_v84_: D0
                d_515_v84_ = D0_DC0((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], (d_432_v28_).f17, _dafny.SeqWithoutIsStrInference([d_387_v1_ for d_516_i11_ in range(default__.abs(105))]), d_514_v83_)
                d_517_v85_: _dafny.Map
                d_517_v85_ = _dafny.Map({D1_DC5(D1_DC3(d_399_v10_, (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], d_515_v84_)): d_459_v37_})
                d_518_v86_: _dafny.Seq
                d_518_v86_ = _dafny.SeqWithoutIsStrInference([(d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], len(d_517_v85_), (0) - ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))])])
                d_519_v87_: _dafny.Map
                d_519_v87_ = _dafny.Map({d_518_v86_: (d_432_v28_).f17})
                d_520_v88_: int
                d_521_v89_: _dafny.Array
                d_522_v90_: _dafny.Set
                out0_: int
                out1_: _dafny.Array
                out2_: _dafny.Set
                out0_, out1_, out2_ = default__.m0(d_513_v82_, (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], (len(d_519_v87_)) > (d_388_v2_), ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]) != (default__.fm0(d_388_v2_, _dafny.Map({-732: -173}), d_396_globalState_)), d_396_globalState_)
                d_520_v88_ = out0_
                d_521_v89_ = out1_
                d_522_v90_ = out2_
                d_390_v5_ = d_390_v5_
                d_523_v91_: int
                d_524_v92_: _dafny.Array
                d_525_v93_: _dafny.Set
                out3_: int
                out4_: _dafny.Array
                out5_: _dafny.Set
                out3_, out4_, out5_ = default__.m0(d_513_v82_, d_520_v88_, not((690) <= ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))])), (d_459_v37_).f17, d_396_globalState_)
                d_523_v91_ = out3_
                d_524_v92_ = out4_
                d_525_v93_ = out5_
                d_526_v94_: _dafny.Set
                d_526_v94_ = _dafny.Set({d_520_v88_})
                d_527_v95_: _dafny.Seq
                d_527_v95_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_387_v1_ for d_528_i13_ in range(default__.abs(611))])])
                d_529_v96_: _dafny.Map
                d_529_v96_ = _dafny.Map({(d_526_v94_).intersection(d_526_v94_): (_dafny.SeqWithoutIsStrInference([(d_515_v84_).cf2 for d_530_i12_ in range(default__.abs(277))])) + (d_527_v95_)})
                (d_396_globalState_).f16 = len(((d_529_v96_)[(d_526_v94_).intersection(d_526_v94_)] if ((d_526_v94_).intersection(d_526_v94_)) in (d_529_v96_) else d_527_v95_))
                pat_let_tv26_ = d_435_v31_
                def iife68_(_pat_let18_0):
                    def iife69_(d_531_dt__update__tmp_h1_):
                        def iife70_(_pat_let19_0):
                            def iife71_(d_532_dt__update_hcf12_h0_):
                                return D1_DC5(d_532_dt__update_hcf12_h0_)
                            return iife71_(_pat_let19_0)
                        return iife70_(pat_let_tv26_)
                    return iife69_(_pat_let18_0)
                d_436_v32_ = iife68_(d_436_v32_)
            elif True:
                d_388_v2_ = 781
                d_533_v97_: _dafny.Array
                nw88_ = _dafny.Array(_dafny.MultiSet({}), 11)
                d_533_v97_ = nw88_
                d_534_v98_: _dafny.Seq
                d_534_v98_ = _dafny.SeqWithoutIsStrInference([d_432_v28_, d_432_v28_])
                d_535_v99_: int
                d_536_v100_: _dafny.Array
                d_537_v101_: _dafny.Set
                out6_: int
                out7_: _dafny.Array
                out8_: _dafny.Set
                out6_, out7_, out8_ = default__.m0(d_533_v97_, (0) - ((len(d_534_v98_) if not((d_432_v28_).f17) else d_388_v2_)), (d_432_v28_).f17, default__.fm4((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], (d_456_v34_)[default__.safeIndex(701, (d_456_v34_).length(0))], d_396_globalState_), d_396_globalState_)
                d_535_v99_ = out6_
                d_536_v100_ = out7_
                d_537_v101_ = out8_
                d_389_v4_ = (d_456_v34_)[default__.safeIndex(701, (d_456_v34_).length(0))]
                d_432_v28_ = (d_459_v37_ if (d_432_v28_).f17 else d_459_v37_)
                d_538_v102_: _dafny.Set
                d_538_v102_ = _dafny.Set({d_432_v28_})
                d_539_v103_: _dafny.Map
                d_539_v103_ = _dafny.Map({d_538_v102_: (d_456_v34_)[default__.safeIndex(701, (d_456_v34_).length(0))]})
                d_540_v104_: _dafny.Map
                d_540_v104_ = _dafny.Map({default__.fm4(len(d_539_v103_), (d_459_v37_).f17, d_396_globalState_): d_456_v34_})
                d_541_v105_: _dafny.Map
                d_541_v105_ = _dafny.Map({d_534_v98_: d_456_v34_})
                d_540_v104_ = (d_540_v104_).set(default__.fm4(d_388_v2_, (d_432_v28_).f17, d_396_globalState_), ((d_541_v105_)[d_534_v98_] if (d_534_v98_) in (d_541_v105_) else d_456_v34_))
            (d_396_globalState_).f11 = d_388_v2_
            if (True) or ((d_393_v8_) <= (d_393_v8_)):
                d_389_v4_ = (d_390_v5_)[default__.safeIndex((0) - ((0) - (len(d_392_v7_))), len(d_390_v5_))]
                index100_ = default__.safeIndex(701, (d_456_v34_).length(0))
                (d_456_v34_)[index100_] = not((d_459_v37_).f17)
                d_542_v106_: _dafny.Map
                d_542_v106_ = _dafny.Map({((d_457_v35_)[(d_432_v28_).f17] if ((d_432_v28_).f17) in (d_457_v35_) else d_388_v2_): d_459_v37_})
                d_543_v107_: _dafny.Seq
                d_543_v107_ = _dafny.SeqWithoutIsStrInference([d_542_v106_])
                index101_ = default__.safeIndex(701, (d_456_v34_).length(0))
                (d_456_v34_)[index101_] = not(((d_542_v106_) | (d_542_v106_)) in ((d_543_v107_) + (d_543_v107_)))
                d_544_v108_: _dafny.Array
                nw89_ = _dafny.Array(_dafny.MultiSet({}), 8)
                d_544_v108_ = nw89_
                d_545_v109_: int
                d_546_v110_: _dafny.Array
                d_547_v111_: _dafny.Set
                out9_: int
                out10_: _dafny.Array
                out11_: _dafny.Set
                out9_, out10_, out11_ = default__.m0(d_544_v108_, len(_dafny.Set({(d_459_v37_).f17})), ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]) < (((d_399_v10_)[d_388_v2_] if (d_388_v2_) in (d_399_v10_) else len(_dafny.SeqWithoutIsStrInference([(d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]])))), (d_459_v37_).f17, d_396_globalState_)
                d_545_v109_ = out9_
                d_546_v110_ = out10_
                d_547_v111_ = out11_
                d_548_v112_: _dafny.Map
                d_548_v112_ = _dafny.Map({d_455_v33_: (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]})
                d_549_v115_: _dafny.Seq
                d_549_v115_ = _dafny.SeqWithoutIsStrInference([d_455_v33_])
                d_550_v116_: _dafny.Array
                nw90_ = _dafny.Array(None, 8)
                nw90_[int(0)] = d_548_v112_
                nw90_[int(1)] = d_548_v112_
                nw90_[int(2)] = (d_548_v112_) | (d_548_v112_)
                def iife72_():
                    coll32_ = _dafny.Map()
                    compr_32_: _dafny.MultiSet
                    for compr_32_ in (d_548_v112_).keys.Elements:
                        d_551_v113_: _dafny.MultiSet = compr_32_
                        if (d_551_v113_) in (d_548_v112_):
                            coll32_[d_551_v113_] = 560
                    return _dafny.Map(coll32_)
                nw90_[int(3)] = iife72_()
                
                nw90_[int(4)] = default__.fm10((d_459_v37_).f17, d_396_globalState_)
                def iife73_():
                    coll33_ = _dafny.Map()
                    compr_33_: _dafny.MultiSet
                    for compr_33_ in (_dafny.MultiSet(d_549_v115_)).Elements:
                        d_552_v114_: _dafny.MultiSet = compr_33_
                        if (d_552_v114_) in (_dafny.MultiSet(d_549_v115_)):
                            coll33_[d_552_v114_] = (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]
                    return _dafny.Map(coll33_)
                nw90_[int(5)] = iife73_()
                
                nw90_[int(6)] = (default__.fm10((d_432_v28_).f17, d_396_globalState_)) | (d_548_v112_)
                nw90_[int(7)] = d_548_v112_
                d_550_v116_ = nw90_
                index102_ = default__.safeIndex(162, (d_550_v116_).length(0))
                (d_550_v116_)[index102_] = d_548_v112_
                index103_ = default__.safeIndex(162, (d_550_v116_).length(0))
                (d_550_v116_)[index103_] = d_548_v112_
            elif True:
                d_553_v117_: _dafny.Seq
                d_553_v117_ = _dafny.SeqWithoutIsStrInference([d_456_v34_])
                d_389_v4_ = ((_dafny.SeqWithoutIsStrInference([d_456_v34_])) + (d_553_v117_)) != (d_553_v117_)
                d_389_v4_ = (d_456_v34_)[default__.safeIndex(701, (d_456_v34_).length(0))]
                index104_ = default__.safeIndex(701, (d_456_v34_).length(0))
                rhs102_ = default__.fm4((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], default__.fm4(d_388_v2_, (d_432_v28_).f17, d_396_globalState_), d_396_globalState_)
                rhs103_ = d_459_v37_
                rhs104_ = default__.safeModuloInt((0) - (d_388_v2_), (0) - (((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]) - (d_388_v2_)))
                lhs68_ = d_456_v34_
                lhs69_ = default__.safeIndex(701, (d_456_v34_).length(0))
                lhs70_ = d_396_globalState_
                lhs68_[lhs69_] = rhs102_
                d_432_v28_ = rhs103_
                lhs70_.f6 = rhs104_
                d_554_v118_: _dafny.Array
                nw91_ = _dafny.Array(_dafny.Set({}), 8)
                d_554_v118_ = nw91_
                d_555_v119_: _dafny.Set
                d_555_v119_ = _dafny.Set({len(d_392_v7_), (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]})
                index105_ = default__.safeIndex(596, (d_554_v118_).length(0))
                (d_554_v118_)[index105_] = d_555_v119_
                index106_ = default__.safeIndex(596, (d_554_v118_).length(0))
                (d_554_v118_)[index106_] = (d_555_v119_) | (d_555_v119_)
                index107_ = default__.safeIndex(701, (d_456_v34_).length(0))
                (d_456_v34_)[index107_] = ((_dafny.MultiSet(d_390_v5_)) - (d_455_v33_)).issubset(d_455_v33_)
        hi2_ = d_388_v2_
        for d_556_i14_ in range((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], hi2_):
            d_389_v4_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "to"))) <= ((d_393_v8_ if default__.fm4((0) - (d_556_i14_), True, d_396_globalState_) else d_393_v8_))
            d_557_v120_: _dafny.Array
            def lambda19_(d_558_v9_):
                def lambda20_(d_559_i15_):
                    return _dafny.MultiSet([(d_558_v9_)[default__.safeIndex(12, (d_558_v9_).length(0))]])

                return lambda20_

            init10_ = lambda19_(d_394_v9_)
            nw92_ = _dafny.Array(None, 16)
            for i0_10_ in range(nw92_.length(0)):
                nw92_[i0_10_] = init10_(i0_10_)
            d_557_v120_ = nw92_
            d_560_v121_: int
            d_561_v122_: _dafny.Array
            d_562_v123_: _dafny.Set
            out12_: int
            out13_: _dafny.Array
            out14_: _dafny.Set
            out12_, out13_, out14_ = default__.m0(d_557_v120_, (d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))], ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]) != ((d_394_v9_)[default__.safeIndex(12, (d_394_v9_).length(0))]), (_dafny.CodePoint('t')) in (d_393_v8_), d_396_globalState_)
            d_560_v121_ = out12_
            d_561_v122_ = out13_
            d_562_v123_ = out14_
            d_563_v124_: int
            d_564_v125_: _dafny.Array
            d_565_v126_: _dafny.Set
            out15_: int
            out16_: _dafny.Array
            out17_: _dafny.Set
            out15_, out16_, out17_ = default__.m0(d_557_v120_, d_560_v121_, (d_392_v7_).isdisjoint(_dafny.Set({(d_456_v34_)[default__.safeIndex(701, (d_456_v34_).length(0))]})), (d_432_v28_).f17, d_396_globalState_)
            d_563_v124_ = out15_
            d_564_v125_ = out16_
            d_565_v126_ = out17_
            d_566_v127_: int
            d_567_v128_: _dafny.Array
            d_568_v129_: _dafny.Set
            out18_: int
            out19_: _dafny.Array
            out20_: _dafny.Set
            out18_, out19_, out20_ = default__.m0(d_557_v120_, default__.safeModuloInt(len(d_393_v8_), d_556_i14_), (d_432_v28_).f17, (d_432_v28_).f17, d_396_globalState_)
            d_566_v127_ = out18_
            d_567_v128_ = out19_
            d_568_v129_ = out20_
        _dafny.print(_dafny.string_of(d_387_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_388_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_389_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_390_v5_) == (_dafny.SeqWithoutIsStrInference([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_391_v6_) == (_dafny.Map({2: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_392_v7_) == (_dafny.Set({True, False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_393_v8_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_394_v9_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_394_v9_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_394_v9_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_394_v9_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_394_v9_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_394_v9_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_394_v9_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_396_globalState_).f2) == (_dafny.Map({_dafny.CodePoint('d'): 514}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_396_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_396_globalState_).f7) == (_dafny.MultiSet([514, -1, 2, 4]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_396_globalState_.f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_.f14)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_.f14)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_.f14)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_.f14)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_.f14)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_.f14)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_.f14)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_396_globalState_.f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_399_v10_) == (_dafny.Map({514: 514, -514: 514}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_432_v28_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_433_v29_) == (_dafny.MultiSet([514]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_434_v30_) == (_dafny.Map({0: _dafny.MultiSet([76])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_435_v31_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_435_v31_).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_435_v31_).cf11) == (_dafny.Map({0: _dafny.MultiSet([76])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_436_v32_).cf12).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_436_v32_).cf12).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_436_v32_).cf12).cf11) == (_dafny.Map({0: _dafny.MultiSet([76])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_455_v33_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_456_v34_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_456_v34_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_457_v35_) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_458_v36_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({514: 514, -514: 514}), _dafny.Map({514: 514, -514: 514})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_459_v37_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_460_v38_) == (_dafny.Map({0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frrp")), 2: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frrp"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[0]) == (_dafny.SeqWithoutIsStrInference([3355, 514, 129]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[1]) == (_dafny.SeqWithoutIsStrInference([3355, 514, 129]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[2]) == (_dafny.SeqWithoutIsStrInference([3355, 514, 129]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[3]) == (_dafny.SeqWithoutIsStrInference([3355, 514, 129]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[4]) == (_dafny.SeqWithoutIsStrInference([3355, 514, 129]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[5]) == (_dafny.SeqWithoutIsStrInference([3355, 514, 129]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[6]) == (_dafny.SeqWithoutIsStrInference([3355, 514, 129]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[7]) == (_dafny.SeqWithoutIsStrInference([3355, 514, 129]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[8]) == (_dafny.SeqWithoutIsStrInference([3355, 514, 129]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[9]) == (_dafny.SeqWithoutIsStrInference([-874]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[10]) == (_dafny.SeqWithoutIsStrInference([3355, 514, 129]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[11]) == (_dafny.SeqWithoutIsStrInference([3355, 514, 129]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[12]) == (_dafny.SeqWithoutIsStrInference([3355, 514, 129]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_462_v39_)[13]) == (_dafny.SeqWithoutIsStrInference([3355, 514, 129]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_464_i6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(int(0), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)}, {self.cf2.VerbatimString(True)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC2(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC2(D1, NamedTuple('DC2', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(int(0), _dafny.CodePoint('D'))
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

class D2_DC7(D2, NamedTuple('DC7', [('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf17 == __o.cf17
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
        return lambda: D3_DC11()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC11(D3, NamedTuple('DC11', [])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11)
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
        return lambda: D4_DC13(False, int(0), _dafny.Array(None, 0))
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

class D4_DC13(D4, NamedTuple('DC13', [('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf23', Any), ('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
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

class D4_DC12(D4, NamedTuple('DC12', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf19 == __o.cf19
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
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D5_DC19)

class D5_DC18(D5, NamedTuple('DC18', [('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC17(D5, NamedTuple('DC17', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC19(D5, NamedTuple('DC19', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC19({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC19) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC21(False, _dafny.CodePoint('D'), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D6_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)

class D6_DC21(D6, NamedTuple('DC21', [('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC21({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC21) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC20(D6, NamedTuple('DC20', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f6: int = int(0)
        self.f11: int = int(0)
        self.f14: _dafny.Array = _dafny.Array(None, 0)
        self.f16: int = int(0)
        self._f0: int = int(0)
        self._f1: int = int(0)
        self._f2: _dafny.Map = _dafny.Map({})
        self._f3: bool = False
        self._f4: int = int(0)
        self._f5: int = int(0)
        self._f7: _dafny.MultiSet = _dafny.MultiSet({})
        self._f8: bool = False
        self._f9: int = int(0)
        self._f10: int = int(0)
        self._f12: int = int(0)
        self._f13: bool = False
        self._f15: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f14 = f14
        (self)._f15 = f15
        (self).f16 = f16

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
    def f5(self):
        return self._f5
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
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    @property
    def f15(self):
        return self._f15

class C0:
    def  __init__(self):
        self._f17: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f17):
        (self)._f17 = f17

    def fm2(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_569_i0_ in range(default__.abs(396))])) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcohubbi"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfaki"))))

    @property
    def f17(self):
        return self._f17
