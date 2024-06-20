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
        return ((_dafny.MultiSet([False])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False, False, False, False])))) not in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])) for d_0_i0_ in range(default__.abs(916))]))

    @staticmethod
    def fm1(p0, p1, globalState):
        return _dafny.CodePoint('m')

    @staticmethod
    def fm2(globalState):
        def iife0_():
            def iife2_():
                coll2_ = _dafny.Set()
                compr_2_: D3
                for compr_2_ in (_dafny.Map({D3_DC8(_dafny.SeqWithoutIsStrInference([False])): True})).keys.Elements:
                    d_4_v1_: D3 = compr_2_
                    if (d_4_v1_) in (_dafny.Map({D3_DC8(_dafny.SeqWithoutIsStrInference([False])): True})):
                        coll2_ = coll2_.union(_dafny.Set([d_4_v1_]))
                return _dafny.Set(coll2_)
            coll0_ = _dafny.Map()
            def iife1_():
                coll1_ = _dafny.Set()
                compr_1_: D3
                for compr_1_ in (_dafny.Map({D3_DC8(_dafny.SeqWithoutIsStrInference([False])): True})).keys.Elements:
                    d_2_v1_: D3 = compr_1_
                    if (d_2_v1_) in (_dafny.Map({D3_DC8(_dafny.SeqWithoutIsStrInference([False])): True})):
                        coll1_ = coll1_.union(_dafny.Set([d_2_v1_]))
                return _dafny.Set(coll1_)
            compr_0_: D3
            for compr_0_ in (iife1_()
            ).Elements:
                d_3_v0_: D3 = compr_0_
                if (d_3_v0_) in (iife2_()
                ):
                    coll0_[d_3_v0_] = True
            return _dafny.Map(coll0_)
        source0_ = (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({len(_dafny.Set({-390}))}))) for d_1_i0_ in range(default__.abs(709))]), _dafny.SeqWithoutIsStrInference([len(iife0_()
        ), 120]), _dafny.SeqWithoutIsStrInference([796]), _dafny.SeqWithoutIsStrInference([126]), _dafny.SeqWithoutIsStrInference([285, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngfivc")))])]) if not(not(not(False))) else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-535, -862]), _dafny.SeqWithoutIsStrInference([569])])))
        d_5___mcc_h0_ = source0_
        d_6_cf88_ = d_5___mcc_h0_
        return D0_DC1()

    @staticmethod
    def fm3(p0, p1, globalState):
        return (0) - ((_dafny.MultiSet([(-560) * (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-401]))])))])).cardinality)

    @staticmethod
    def fm4(p0, p1, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(366, 228):
                d_9_v0_: int = compr_3_
                if ((366) <= (d_9_v0_)) and ((d_9_v0_) < (228)):
                    coll3_[(d_9_v0_) - (((D16_DC33(False, len(_dafny.Map({False: 734})), _dafny.MultiSet([False]), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))), _dafny.Set({True}))).cf54).cardinality)] = _dafny.CodePoint('y')
            return _dafny.Map(coll3_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m')])), -29, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({True: False})]))), len(_dafny.Map({-691: (_dafny.MultiSet([True, not(True)])).cardinality})), -9]) for d_7_i0_ in range(default__.abs(710))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False, False, True])]))]) for d_8_i1_ in range(default__.abs(314))]) if True else _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.Map({len(_dafny.Map({967: 448})): len(iife3_()
        )}))])])))

    @staticmethod
    def fm5(p0, p1, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: D0
            for compr_4_ in (_dafny.Set({D0_DC1()})).Elements:
                d_10_v0_: D0 = compr_4_
                if (d_10_v0_) in (_dafny.Set({D0_DC1()})):
                    coll4_[d_10_v0_] = 558
            return _dafny.Map(coll4_)
        return ((_dafny.Map({D0_DC1(): 603})) | (_dafny.Map({D0_DC1(): 157}))) | ((iife4_()
        ) | (_dafny.Map({D0_DC1(): (_dafny.MultiSet([_dafny.CodePoint('c'), _dafny.CodePoint('e'), _dafny.CodePoint('u')])).cardinality})))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        return D1_DC4(len(_dafny.Map({_dafny.SeqWithoutIsStrInference([-606, len(_dafny.SeqWithoutIsStrInference([False]))]): len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({167})), 841, 909]))})), (0) - ((len(_dafny.Map({464: (D21_DC43(True, False, _dafny.SeqWithoutIsStrInference([not(False)]))).cf70}))) * (159)))

    @staticmethod
    def fm16(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eh"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgorcc")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxshrry")))

    @staticmethod
    def fm18(p0, p1, globalState):
        return _dafny.Map({default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('n')])), 28): (_dafny.Set({True})).intersection(_dafny.Set({False}))})

    @staticmethod
    def fm19(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False, False, False, False]))

    @staticmethod
    def fm20(p0, p1, globalState):
        source1_ = D6_DC15(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False])]))
        if source1_.is_DC16:
            d_11___mcc_h0_ = source1_.cf20
            d_12___mcc_h1_ = source1_.cf21
            d_13___mcc_h2_ = source1_.cf22
            d_14_cf22_ = d_13___mcc_h2_
            d_15_cf21_ = d_12___mcc_h1_
            d_16_cf20_ = d_11___mcc_h0_
            return D3_DC8(_dafny.SeqWithoutIsStrInference([False, d_16_cf20_]))
        elif True:
            d_17___mcc_h3_ = source1_.cf19
            d_18_cf19_ = d_17___mcc_h3_
            if False:
                return D3_DC8(_dafny.SeqWithoutIsStrInference([True]))
            elif True:
                return D3_DC8(_dafny.SeqWithoutIsStrInference([not(False), True]))

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        source2_ = D19_DC38(False, not(True))
        if source2_.is_DC38:
            d_19___mcc_h0_ = source2_.cf63
            d_20___mcc_h1_ = source2_.cf64
            d_21_cf64_ = d_20___mcc_h1_
            d_22_cf63_ = d_19___mcc_h0_
            return _dafny.MultiSet([d_21_cf64_, d_21_cf64_, d_22_cf63_, d_21_cf64_])
        elif source2_.is_DC37:
            d_23___mcc_h2_ = source2_.cf62
            d_24_cf62_ = d_23___mcc_h2_
            return (D16_DC33(True, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: 250}))])), _dafny.MultiSet([True]), 527, _dafny.Set({not(True)}))).cf54
        elif True:
            d_25___mcc_h3_ = source2_.cf65
            d_26_cf65_ = d_25___mcc_h3_
            return (_dafny.MultiSet([False, True])) | (_dafny.MultiSet([True, True]))

    @staticmethod
    def fm22(globalState):
        return (_dafny.Map({_dafny.MultiSet([(_dafny.MultiSet([True])).cardinality]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))})) | ((_dafny.Map({_dafny.MultiSet([855, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xt")))]): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_27_i0_ in range(default__.abs(290))])})) | (_dafny.Map({_dafny.MultiSet([(0) - (len(_dafny.Map({False: False})))]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmrrdi"))})))

    @staticmethod
    def fm23(p0, p1, globalState):
        return _dafny.Set({False})

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([not(not(True))]), _dafny.SeqWithoutIsStrInference([True])])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False])]))) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True])]))

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([D0_DC1(), D0_DC1()])) + ((_dafny.SeqWithoutIsStrInference([D0_DC1(), D0_DC1(), D0_DC1(), D0_DC1()])) + (_dafny.SeqWithoutIsStrInference([D0_DC1()])))

    @staticmethod
    def fm27(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False])])) | ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, not(True), not(True), False]), _dafny.SeqWithoutIsStrInference([False, True, True])])) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True])])))

    @staticmethod
    def fm28(globalState):
        if (958) == (397):
            return _dafny.Map({False: -268})
        elif True:
            return _dafny.Map({not(False): 388})

    @staticmethod
    def fm29(globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: _dafny.Seq
            for compr_5_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([True, False]), _dafny.SeqWithoutIsStrInference([True, True, False, not(True), not(True)]), _dafny.SeqWithoutIsStrInference([True, True, False])})).Elements:
                d_28_v0_: _dafny.Seq = compr_5_
                if (d_28_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([True, False]), _dafny.SeqWithoutIsStrInference([True, True, False, not(True), not(True)]), _dafny.SeqWithoutIsStrInference([True, True, False])})):
                    coll5_[d_28_v0_] = not(True)
            return _dafny.Map(coll5_)
        return (_dafny.Map({452: 863})) | (_dafny.Map({len(iife5_()
        ): -354}))

    @staticmethod
    def fm30(p0, p1, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(122, -15):
                d_29_v0_: int = compr_6_
                if ((122) <= (d_29_v0_)) and ((d_29_v0_) < (-15)):
                    coll6_[(d_29_v0_) * (-977)] = True
            return _dafny.Map(coll6_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([812]), _dafny.SeqWithoutIsStrInference([-968, len(_dafny.SeqWithoutIsStrInference([False, True]))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: False})), len(_dafny.SeqWithoutIsStrInference([True, True])), len(iife6_()
)]) for d_30_i0_ in range(default__.abs(573))]))

    @staticmethod
    def fm31(p0, p1, p2, p3, globalState):
        return (_dafny.Set({-505, len(_dafny.Set({-506}))}))

    @staticmethod
    def fm32(p0, p1, globalState):
        return (_dafny.Set({len(_dafny.Set({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.SeqWithoutIsStrInference([False]))])).cardinality, len(_dafny.Map({not(False): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vyojj"))}))}))})).intersection(_dafny.Set({864}))

    @staticmethod
    def fm33(p0, p1, p2, globalState):
        return D6_DC15((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, True])])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, True])])))

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: _dafny.MultiSet
            for compr_7_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(_dafny.MultiSet([True])).cardinality, 236])])).Elements:
                d_31_v0_: _dafny.MultiSet = compr_7_
                if (d_31_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(_dafny.MultiSet([True])).cardinality, 236])])):
                    coll7_[d_31_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
            return _dafny.Map(coll7_)
        return ((iife7_()
        ) | (_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({299, 935, -557, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chaced"))), 934})), 375])): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnbkik"))}))) | ((_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality for d_32_i0_ in range(default__.abs(343))])): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdpr"))}) if False else _dafny.Map({_dafny.MultiSet([len(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_33_i1_ in range(default__.abs(25))]): -972}))]): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_34_i2_ in range(default__.abs(331))])})))

    @staticmethod
    def fm35(p0, globalState):
        return ((_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tghobrq")))})).intersection(_dafny.Set({-815}))).intersection((_dafny.Set({(0) - ((0) - (len(_dafny.Set({len(_dafny.Set({False, True})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))})))), 186})) | (_dafny.Set({858, -526, (D11_DC24(not(True), True, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfoksstqw")))))).cf38, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_35_i0_ in range(default__.abs(653))]))})))

    @staticmethod
    def fm36(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([35])) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(not(False)), True])).cardinality]))) + (_dafny.SeqWithoutIsStrInference([903 for d_36_i0_ in range(default__.abs(918))]))

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcdnnimh"))

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        if (472) < (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ky")))}))):
            return D0_DC0((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ys")))]))))
        elif not(True):
            return D0_DC0(621)
        elif True:
            return D0_DC0(820)

    @staticmethod
    def fm39(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjnq")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_37_i0_ in range(default__.abs(462))])])), len(_dafny.Map({_dafny.MultiSet([False]): _dafny.CodePoint('d')})))])

    @staticmethod
    def fm40(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([True]))) + ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm41(p0, p1, p2, globalState):
        def iife8_():
            coll8_ = _dafny.Set()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(763, 69):
                d_38_v0_: int = compr_8_
                if ((763) <= (d_38_v0_)) and ((d_38_v0_) < (69)):
                    coll8_ = coll8_.union(_dafny.Set([(d_38_v0_) - (256)]))
            return _dafny.Set(coll8_)
        def iife9_():
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: int
                for compr_11_ in _dafny.IntegerRange(288, 859):
                    d_39_v1_: int = compr_11_
                    if ((288) <= (d_39_v1_)) and ((d_39_v1_) < (859)):
                        coll11_[(d_39_v1_) + (596)] = len(_dafny.Set({False, False, True}))
                return _dafny.Map(coll11_)
            coll9_ = _dafny.Set()
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(288, 859):
                    d_39_v1_: int = compr_10_
                    if ((288) <= (d_39_v1_)) and ((d_39_v1_) < (859)):
                        coll10_[(d_39_v1_) + (596)] = len(_dafny.Set({False, False, True}))
                return _dafny.Map(coll10_)
            compr_9_: int
            for compr_9_ in (iife10_()
            ).keys.Elements:
                d_40_v2_: int = compr_9_
                if (d_40_v2_) in (iife11_()
                ):
                    coll9_ = coll9_.union(_dafny.Set([(d_40_v2_) + (-500)]))
            return _dafny.Set(coll9_)
        return ((_dafny.Set({len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y'), _dafny.CodePoint('a')]))])]))})), len(_dafny.SeqWithoutIsStrInference([False, True])), -193})) - (iife8_()
        )).intersection((iife9_()
        ) | (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([True, True, False]))})))

    @staticmethod
    def fm42(p0, p1, globalState):
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: int
            for compr_12_ in _dafny.IntegerRange(581, 197):
                d_41_v0_: int = compr_12_
                if ((581) <= (d_41_v0_)) and ((d_41_v0_) < (197)):
                    coll12_[default__.safeDivisionInt(d_41_v0_, 501)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(not(not(not(not(False))))), False]))).cardinality
            return _dafny.Map(coll12_)
        return (_dafny.Map({459: (0) - (-980)})) | (iife12_()
        )

    @staticmethod
    def fm43(globalState):
        return D6_DC16((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exmbhmdo"))), not (True) or (True), (_dafny.SeqWithoutIsStrInference([-346])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onku")))])))

    @staticmethod
    def fm44(globalState):
        return _dafny.Map({False: 445})

    @staticmethod
    def fm45(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True, False]), _dafny.MultiSet([True, False, False, True, True]), _dafny.MultiSet([False])])) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False]), _dafny.MultiSet([True])]) if not(True) else _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([not(False)]) for d_42_i0_ in range(default__.abs(824))])))

    @staticmethod
    def fm46(p0, globalState):
        if True:
            return _dafny.Set({False})
        elif True:
            return (D16_DC33(False, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False}))])), _dafny.MultiSet([False, True]), 873, _dafny.Set({not(True), True}))).cf56

    @staticmethod
    def fm47(p0, globalState):
        return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([True])) + ((_dafny.SeqWithoutIsStrInference([False, True]) if True else _dafny.SeqWithoutIsStrInference([not(True)]))))

    @staticmethod
    def fm49(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([191]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({737, -741})), 344])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({606: (D6_DC16(True, False, _dafny.SeqWithoutIsStrInference([len(_dafny.Map({576: len(_dafny.Map({True: False}))}))]))).cf20}))]))]) for d_43_i0_ in range(default__.abs(-224))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, True]))])]))

    @staticmethod
    def fm50(p0, p1, p2, p3, globalState):
        source3_ = D1_DC3(not(not(True)))
        if source3_.is_DC4:
            d_44___mcc_h0_ = source3_.cf3
            d_45___mcc_h1_ = source3_.cf4
            d_46_cf4_ = d_45___mcc_h1_
            d_47_cf3_ = d_44___mcc_h0_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyg"))
        elif True:
            d_48___mcc_h2_ = source3_.cf2
            d_49_cf2_ = d_48___mcc_h2_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dstoi"))

    @staticmethod
    def fm51(p0, p1, globalState):
        return (_dafny.Map({891: False})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([372])): False}))

    @staticmethod
    def fm52(p0, globalState):
        return D3_DC8((_dafny.SeqWithoutIsStrInference([(D16_DC33(False, -609, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True, False, False])), -454, _dafny.Set({False}))).cf52]) if False else _dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm53(p0, p1, globalState):
        return (_dafny.Set({True, False})) - ((_dafny.Set({False, False})).intersection(_dafny.Set({True})))

    @staticmethod
    def fm54(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: False})])) + ((_dafny.SeqWithoutIsStrInference([_dafny.Map({False: True})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: not(False)}) for d_50_i0_ in range(default__.abs(155))])))

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        source4_ = D21_DC43(False, True, _dafny.SeqWithoutIsStrInference([not(not(True))]))
        if source4_.is_DC43:
            d_51___mcc_h0_ = source4_.cf69
            d_52___mcc_h1_ = source4_.cf70
            d_53___mcc_h2_ = source4_.cf71
            d_54_cf71_ = d_53___mcc_h2_
            d_55_cf70_ = d_52___mcc_h1_
            d_56_cf69_ = d_51___mcc_h0_
            return D1_DC3(False)
        elif source4_.is_DC44:
            d_57___mcc_h3_ = source4_.cf72
            d_58_cf72_ = d_57___mcc_h3_
            return D1_DC3(False)
        elif True:
            d_59___mcc_h4_ = source4_.cf68
            d_60_cf68_ = d_59___mcc_h4_
            return D1_DC3(True)

    @staticmethod
    def fm56(p0, p1, p2, globalState):
        if (False) and (not(True)):
            return _dafny.SeqWithoutIsStrInference([False])
        elif True:
            return (_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([not(False)]))

    @staticmethod
    def fm57(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: 681})])) for d_61_i0_ in range(default__.abs(189))])) + (_dafny.SeqWithoutIsStrInference([-469]))) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.Map({_dafny.CodePoint('b'): False}), _dafny.Map({_dafny.CodePoint('x'): False})})) for d_62_i1_ in range(default__.abs(-661))]))

    @staticmethod
    def fm58(globalState):
        return ((_dafny.Map({False: D12_DC27(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])]))})) | (_dafny.Map({True: D12_DC27(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False])]))}))) | (_dafny.Map({True: D12_DC27(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([True])]))}))

    @staticmethod
    def fm59(globalState):
        source5_ = D21_DC43(True, False, _dafny.SeqWithoutIsStrInference([False, False]))
        if source5_.is_DC43:
            d_63___mcc_h0_ = source5_.cf69
            d_64___mcc_h1_ = source5_.cf70
            d_65___mcc_h2_ = source5_.cf71
            d_66_cf71_ = d_65___mcc_h2_
            d_67_cf70_ = d_64___mcc_h1_
            d_68_cf69_ = d_63___mcc_h0_
            def iife13_():
                coll13_ = _dafny.Map()
                compr_13_: D10
                for compr_13_ in (_dafny.SeqWithoutIsStrInference([D10_DC21(_dafny.Map({d_68_cf69_: 804}))])).Elements:
                    d_69_v0_: D10 = compr_13_
                    if (d_69_v0_) in (_dafny.SeqWithoutIsStrInference([D10_DC21(_dafny.Map({d_68_cf69_: 804}))])):
                        coll13_[d_69_v0_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sg")))
                return _dafny.Map(coll13_)
            return iife13_()
            
        elif source5_.is_DC44:
            d_70___mcc_h3_ = source5_.cf72
            d_71_cf72_ = d_70___mcc_h3_
            return _dafny.Map({D10_DC21(_dafny.Map({d_71_cf72_: 18})): (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")))))})
        elif True:
            d_72___mcc_h4_ = source5_.cf68
            d_73_cf68_ = d_72___mcc_h4_
            return (_dafny.Map({D10_DC21(_dafny.Map({True: len(_dafny.Map({False: 2}))})): (0) - (len(_dafny.SeqWithoutIsStrInference([776])))})) | (_dafny.Map({D10_DC21(_dafny.Map({False: 335})): len(_dafny.Set({377, (0) - ((0) - ((_dafny.MultiSet([762, -921])).cardinality))}))}))

    @staticmethod
    def fm60(p0, p1, p2, globalState):
        def iife14_():
            coll14_ = _dafny.Map()
            compr_14_: int
            for compr_14_ in (_dafny.MultiSet([896])).Elements:
                d_74_v0_: int = compr_14_
                if (d_74_v0_) in (_dafny.MultiSet([896])):
                    coll14_[(d_74_v0_) - (75)] = True
            return _dafny.Map(coll14_)
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference([False, False])) + (_dafny.SeqWithoutIsStrInference([False])): (_dafny.Map({len(_dafny.Map({909: 368})): not(not(not(False)))})) | (iife14_()
        )})

    @staticmethod
    def fm61(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, not(True), False, True, False])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(not(False)), True, True, False, not(False)])]))

    @staticmethod
    def fm62(globalState):
        return D4_DC12((_dafny.CodePoint('p') if not(True) else _dafny.CodePoint('a')))

    @staticmethod
    def fm63(globalState):
        return D4_DC13(len((_dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hxnhw")))})) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aog")))}))), (False) == (False), 850, _dafny.SeqWithoutIsStrInference([not(False)]))

    @staticmethod
    def fm64(p0, p1, p2, globalState):
        def iife15_():
            coll15_ = _dafny.Set()
            compr_15_: str
            for compr_15_ in (_dafny.MultiSet([_dafny.CodePoint('f')])).Elements:
                d_75_v0_: str = compr_15_
                if (d_75_v0_) in (_dafny.MultiSet([_dafny.CodePoint('f')])):
                    coll15_ = coll15_.union(_dafny.Set([d_75_v0_]))
            return _dafny.Set(coll15_)
        return ((_dafny.Set({_dafny.CodePoint('e')})) | (_dafny.Set({_dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('x')}))) | ((iife15_()
        ) - (_dafny.Set({_dafny.CodePoint('a')})))

    @staticmethod
    def fm65(p0, p1, p2, globalState):
        return ((_dafny.Map({True: _dafny.CodePoint('f')})) | (_dafny.Map({False: _dafny.CodePoint('r')}))) | ((_dafny.Map({True: _dafny.CodePoint('q')})) | (_dafny.Map({not(False): _dafny.CodePoint('j')})))

    @staticmethod
    def fm66(p0, globalState):
        return D16_DC33(not (True) or (False), default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gk"))), 925), ((D16_DC33(False, len(_dafny.SeqWithoutIsStrInference([False])), _dafny.MultiSet([not(False)]), 249, _dafny.Set({False}))).cf54).intersection(_dafny.MultiSet([True])), -152, (_dafny.Set({False})) - (_dafny.Set({False})))

    @staticmethod
    def fm67(p0, globalState):
        if True:
            return _dafny.Map({_dafny.Map({876: 861}): not(not(True))})
        elif True:
            return (_dafny.Map({_dafny.Map({len(_dafny.Set({720, 548})): (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rss")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))])).cardinality}): True})) | (_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udygyb"))): -885}): not(True)}))

    @staticmethod
    def fm68(globalState):
        return D12_DC27((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]) for d_76_i0_ in range(default__.abs(223))]) if False else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])])))

    @staticmethod
    def fm69(globalState):
        return _dafny.Map({((_dafny.MultiSet([(0) - (-3)])) | (_dafny.MultiSet([-982]))).cardinality: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igvytirrh"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_77_i0_ in range(default__.abs(614))]))})

    @staticmethod
    def fm70(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('d')]) for d_78_i0_ in range(default__.abs(7))])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('j'), _dafny.CodePoint('q')]) for d_79_i1_ in range(default__.abs(866))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('k')]), _dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('c'), _dafny.CodePoint('f'), _dafny.CodePoint('b')]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')])), _dafny.MultiSet([_dafny.CodePoint('b'), _dafny.CodePoint('c')]), _dafny.MultiSet([_dafny.CodePoint('h')])]))

    @staticmethod
    def fm71(globalState):
        return (_dafny.Map({True: True})) | (_dafny.Map({not(False): False}))

    @staticmethod
    def fm72(p0, p1, globalState):
        source6_ = _dafny.Set({350, len(_dafny.SeqWithoutIsStrInference([True, not(True), True])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_80_i0_ in range(default__.abs(846))]))})
        d_81___mcc_h0_ = source6_
        d_82_cf23_ = d_81___mcc_h0_
        return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([180]) for d_83_i1_ in range(default__.abs(510))]))

    @staticmethod
    def fm73(globalState):
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: int
            for compr_16_ in (_dafny.Map({544: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkivckp")))})).keys.Elements:
                d_84_v0_: int = compr_16_
                if (d_84_v0_) in (_dafny.Map({544: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkivckp")))})):
                    coll16_[default__.safeModuloInt(d_84_v0_, 688)] = False
            return _dafny.Map(coll16_)
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(iife16_()
        )]), _dafny.SeqWithoutIsStrInference([len((_dafny.Map({_dafny.Map({False: True}): not(False)})))])])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([732])])))) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({602: -700})) for d_85_i0_ in range(default__.abs(519))])]))

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        if True:
            d_86_v0_: _dafny.Array
            nw0_ = _dafny.Array(_dafny.Seq({}), 5)
            d_86_v0_ = nw0_
            d_87_v1_: _dafny.Map
            d_87_v1_ = _dafny.Map({p1: d_86_v0_})
            d_86_v0_ = ((d_87_v1_)[p3] if (p3) in (d_87_v1_) else d_86_v0_)
            d_88_v3_: int
            d_88_v3_ = 50
            d_89_v4_: _dafny.Seq
            d_89_v4_ = _dafny.SeqWithoutIsStrInference([d_88_v3_, d_88_v3_, d_88_v3_])
            d_90_v5_: _dafny.Set
            d_90_v5_ = _dafny.Set({d_88_v3_})
            d_91_v6_: _dafny.MultiSet
            d_91_v6_ = _dafny.MultiSet([p3])
            d_92_v7_: _dafny.Set
            d_92_v7_ = _dafny.Set({p3, p1, p3, p1, p3})
            d_93_v8_: _dafny.Map
            d_93_v8_ = _dafny.Map({p3: p3})
            d_94_v9_: _dafny.Seq
            d_94_v9_ = _dafny.SeqWithoutIsStrInference([d_93_v8_, _dafny.Map({p1: p3}), _dafny.Map({p1: p1}), d_93_v8_])
            d_95_v10_: _dafny.Array
            nw1_ = _dafny.Array(None, 19)
            nw1_[int(0)] = p3
            nw1_[int(1)] = p3
            nw1_[int(2)] = p1
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in (d_89_v4_).Elements:
                    d_96_v2_: int = compr_17_
                    if (d_96_v2_) in (d_89_v4_):
                        coll17_[default__.safeDivisionInt(d_96_v2_, 88)] = (0) - (d_88_v3_)
                return _dafny.Map(coll17_)
            nw1_[int(3)] = (len(iife17_()
            )) != (803)
            nw1_[int(4)] = p1
            nw1_[int(5)] = p1
            nw1_[int(6)] = (d_90_v5_).ispropersubset(d_90_v5_)
            nw1_[int(7)] = p1
            nw1_[int(8)] = p1
            nw1_[int(9)] = (d_88_v3_) <= (d_88_v3_)
            nw1_[int(10)] = (d_91_v6_).isdisjoint(d_91_v6_)
            nw1_[int(11)] = (d_92_v7_).ispropersubset(default__.fm53(d_88_v3_, d_94_v9_, globalState))
            nw1_[int(12)] = p1
            nw1_[int(13)] = p1
            nw1_[int(14)] = p1
            nw1_[int(15)] = p1
            nw1_[int(16)] = not (p1) or (p3)
            nw1_[int(17)] = p3
            nw1_[int(18)] = p3
            d_95_v10_ = nw1_
            d_97_v11_: _dafny.Seq
            d_97_v11_ = _dafny.SeqWithoutIsStrInference([True])
            index0_ = default__.safeIndex(167, (d_95_v10_).length(0))
            (d_95_v10_)[index0_] = (d_97_v11_)[default__.safeIndex(d_88_v3_, len(d_97_v11_))]
            d_98_v12_: _dafny.Seq
            d_98_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fku"))
            index1_ = default__.safeIndex(167, (d_95_v10_).length(0))
            (d_95_v10_)[index1_] = (d_90_v5_).ispropersubset(_dafny.Set({len(d_98_v12_), d_88_v3_}))
            d_99_v13_: _dafny.Array
            nw2_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_99_v13_ = nw2_
            d_100_v14_: C4
            nw3_ = C4()
            nw3_.ctor__(d_99_v13_, (0) - ((0) - (d_88_v3_)), (d_95_v10_)[default__.safeIndex(167, (d_95_v10_).length(0))])
            d_100_v14_ = nw3_
            d_100_v14_ = d_100_v14_
            d_101_v15_: _dafny.Array
            nw4_ = _dafny.Array(int(0), 5)
            d_101_v15_ = nw4_
            index2_ = default__.safeIndex(173, (d_101_v15_).length(0))
            (d_101_v15_)[index2_] = ((d_91_v6_).cardinality) - (d_88_v3_)
            index3_ = default__.safeIndex(173, (d_101_v15_).length(0))
            (d_101_v15_)[index3_] = ((((p2)[len(d_92_v7_)] if (len(d_92_v7_)) in (p2) else len(d_98_v12_))) + (d_88_v3_)) + (d_88_v3_)
            d_102_v16_: C2
            nw5_ = C2()
            nw5_.ctor__(d_98_v12_, (d_97_v11_)[default__.safeIndex((0) - ((d_101_v15_)[default__.safeIndex(173, (d_101_v15_).length(0))]), len(d_97_v11_))])
            d_102_v16_ = nw5_
            d_103_v17_: D27
            d_103_v17_ = D27_DC56(d_102_v16_)
            d_102_v16_ = (d_103_v17_).cf95
        elif True:
            d_104_v18_: str
            d_104_v18_ = _dafny.CodePoint('r')
            d_105_v19_: _dafny.Array
            nw6_ = _dafny.Array(None, 8)
            nw6_[int(0)] = p1
            nw6_[int(1)] = default__.fm0(d_104_v18_, globalState)
            nw6_[int(2)] = p3
            nw6_[int(3)] = p3
            nw6_[int(4)] = p1
            nw6_[int(5)] = False
            nw6_[int(6)] = p3
            nw6_[int(7)] = p1
            d_105_v19_ = nw6_
            d_106_v20_: _dafny.Map
            d_106_v20_ = _dafny.Map({p3: d_105_v19_})
            d_107_v21_: _dafny.Array
            nw7_ = _dafny.Array(None, 22)
            nw7_[int(0)] = d_105_v19_
            nw7_[int(1)] = d_105_v19_
            nw7_[int(2)] = d_105_v19_
            nw7_[int(3)] = d_105_v19_
            nw7_[int(4)] = d_105_v19_
            nw7_[int(5)] = d_105_v19_
            nw7_[int(6)] = d_105_v19_
            nw7_[int(7)] = d_105_v19_
            nw7_[int(8)] = d_105_v19_
            nw7_[int(9)] = d_105_v19_
            nw7_[int(10)] = d_105_v19_
            nw7_[int(11)] = d_105_v19_
            nw7_[int(12)] = d_105_v19_
            nw7_[int(13)] = d_105_v19_
            nw7_[int(14)] = d_105_v19_
            nw7_[int(15)] = d_105_v19_
            nw7_[int(16)] = d_105_v19_
            nw7_[int(17)] = d_105_v19_
            nw7_[int(18)] = d_105_v19_
            nw7_[int(19)] = d_105_v19_
            nw7_[int(20)] = ((d_106_v20_)[p3] if (p3) in (d_106_v20_) else d_105_v19_)
            nw7_[int(21)] = d_105_v19_
            d_107_v21_ = nw7_
            d_108_v22_: int
            d_108_v22_ = 852
            d_109_v23_: T1
            nw8_ = C4()
            nw8_.ctor__(d_107_v21_, d_108_v22_, p1)
            d_109_v23_ = nw8_
            d_110_v24_: _dafny.Set
            d_110_v24_ = _dafny.Set({d_109_v23_})
            d_111_v25_: D12
            d_111_v25_ = D12_DC28(len(d_110_v24_), d_108_v22_, (d_109_v23_).f4, (d_109_v23_).f4, default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbi"))), (d_109_v23_).f6))
            d_111_v25_ = d_111_v25_
            d_112_v26_: _dafny.Seq
            d_112_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "filnlnc"))
            d_112_v26_ = d_112_v26_
            d_113_v27_: _dafny.Seq
            d_113_v27_ = _dafny.SeqWithoutIsStrInference([(d_109_v23_).f6, d_108_v22_])
            d_114_v28_: _dafny.MultiSet
            d_114_v28_ = _dafny.MultiSet([p3])
            d_115_v29_: _dafny.Array
            nw9_ = _dafny.Array(None, 22)
            nw9_[int(0)] = d_113_v27_
            nw9_[int(1)] = d_113_v27_
            nw9_[int(2)] = d_113_v27_
            nw9_[int(3)] = _dafny.SeqWithoutIsStrInference([len(d_112_v26_)])
            nw9_[int(4)] = (d_113_v27_) + (d_113_v27_)
            nw9_[int(5)] = (d_113_v27_).set(default__.safeIndex(len(d_112_v26_), len(d_113_v27_)), (d_109_v23_).f6)
            nw9_[int(6)] = (d_113_v27_) + (d_113_v27_)
            nw9_[int(7)] = (d_113_v27_) + (d_113_v27_)
            nw9_[int(8)] = d_113_v27_
            nw9_[int(9)] = d_113_v27_
            nw9_[int(10)] = _dafny.SeqWithoutIsStrInference([((d_114_v28_)[p1] if (p1) in (d_114_v28_) else d_108_v22_)])
            nw9_[int(11)] = (d_113_v27_) + (_dafny.SeqWithoutIsStrInference([(d_113_v27_)[default__.safeIndex((d_109_v23_).f6, len(d_113_v27_))], len(d_112_v26_), d_108_v22_, ((d_114_v28_)[False] if (False) in (d_114_v28_) else (p2).cardinality), d_108_v22_]))
            nw9_[int(12)] = _dafny.SeqWithoutIsStrInference([d_108_v22_, len(d_112_v26_)])
            nw9_[int(13)] = (_dafny.SeqWithoutIsStrInference([(d_109_v23_).f6])) + (d_113_v27_)
            nw9_[int(14)] = d_113_v27_
            nw9_[int(15)] = d_113_v27_
            nw9_[int(16)] = d_113_v27_
            nw9_[int(17)] = d_113_v27_
            nw9_[int(18)] = d_113_v27_
            nw9_[int(19)] = d_113_v27_
            nw9_[int(20)] = d_113_v27_
            nw9_[int(21)] = default__.fm57((d_109_v23_).f6, (d_109_v23_).f6, globalState)
            d_115_v29_ = nw9_
            d_116_v30_: _dafny.Map
            d_116_v30_ = _dafny.Map({p1: d_108_v22_})
            d_117_v31_: D10
            d_117_v31_ = D10_DC21((d_116_v30_).set(p3, (d_109_v23_).f6))
            d_118_v32_: _dafny.Seq
            d_118_v32_ = _dafny.SeqWithoutIsStrInference([d_117_v31_])
            d_119_v33_: _dafny.Seq
            d_119_v33_ = _dafny.SeqWithoutIsStrInference([default__.fm0(d_104_v18_, globalState), p3])
            rhs0_ = not((d_119_v33_)[default__.safeIndex(d_108_v22_, len(d_119_v33_))])
            rhs1_ = d_115_v29_
            rhs2_ = (d_118_v32_ if False else (d_118_v32_).set(default__.safeIndex((d_109_v23_).f6, len(d_118_v32_)), d_117_v31_))
            rhs3_ = (d_109_v23_).fm7(globalState)
            r0 = rhs0_
            d_115_v29_ = rhs1_
            d_118_v32_ = rhs2_
            r1 = rhs3_
            d_120_v34_: D16
            d_120_v34_ = D16_DC33((d_109_v23_).f4, (d_109_v23_).fm9(p3, p3, (d_109_v23_).f6, globalState), d_114_v28_, len(d_112_v26_), _dafny.Set({(d_109_v23_).f4, p1}))
            r0 = not(not (((d_120_v34_).cf54).issubset(d_114_v28_)) or ((d_120_v34_).cf52))
            d_121_v35_: _dafny.Map
            d_121_v35_ = _dafny.Map({d_108_v22_: p3})
            d_122_v36_: _dafny.Map
            d_122_v36_ = _dafny.Map({d_104_v18_: d_119_v33_})
            d_123_v37_: C1
            nw10_ = C1()
            nw10_.ctor__((0) - (default__.safeDivisionInt(185, ((d_114_v28_)[((d_121_v35_)[884] if (884) in (d_121_v35_) else p1)] if (((d_121_v35_)[884] if (884) in (d_121_v35_) else p1)) in (d_114_v28_) else d_108_v22_))), d_108_v22_, (d_107_v21_ if p1 else d_109_v23_.f5), len(d_122_v36_), (p2).ispropersubset(p2))
            d_123_v37_ = nw10_
            rhs4_ = d_123_v37_
            rhs5_ = (d_113_v27_) + (_dafny.SeqWithoutIsStrInference([(0) - ((d_109_v23_).f6) for d_124_i0_ in range(default__.abs(476))]))
            rhs6_ = d_105_v19_
            d_123_v37_ = rhs4_
            d_113_v27_ = rhs5_
            d_105_v19_ = rhs6_
        d_125_v38_: int
        d_125_v38_ = 197
        hi0_ = default__.fm3(d_125_v38_, True, globalState)
        for d_126_i1_ in range(d_125_v38_, hi0_):
            d_125_v38_ = default__.safeModuloInt(d_125_v38_, d_125_v38_)
            d_127_v39_: _dafny.Seq
            d_127_v39_ = _dafny.SeqWithoutIsStrInference([d_125_v38_])
            d_128_v40_: _dafny.Seq
            d_128_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etegojdr"))
            r0 = (p3 if False else (d_127_v39_) != ((d_127_v39_).set(default__.safeIndex(len(d_128_v40_), len(d_127_v39_)), d_126_i1_)))
            d_129_v41_: C2
            nw11_ = C2()
            nw11_.ctor__((d_128_v40_) + (d_128_v40_), p3)
            d_129_v41_ = nw11_
            d_129_v41_ = (d_129_v41_ if p1 else d_129_v41_)
            d_130_v42_: _dafny.Seq
            d_130_v42_ = _dafny.SeqWithoutIsStrInference([not(p1)])
            d_131_v43_: _dafny.Set
            d_131_v43_ = _dafny.Set({d_130_v42_, d_130_v42_, d_130_v42_, d_130_v42_})
            d_132_v44_: _dafny.Seq
            d_132_v44_ = _dafny.SeqWithoutIsStrInference([d_130_v42_, d_130_v42_])
            d_133_v46_: _dafny.Array
            nw12_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_133_v46_ = nw12_
            d_134_v47_: C9
            nw13_ = C9()
            def iife18_():
                coll18_ = _dafny.Set()
                compr_18_: _dafny.Seq
                for compr_18_ in (d_132_v44_).Elements:
                    d_135_v45_: _dafny.Seq = compr_18_
                    if (d_135_v45_) in (d_132_v44_):
                        coll18_ = coll18_.union(_dafny.Set([d_135_v45_]))
                return _dafny.Set(coll18_)
            nw13_.ctor__(d_126_i1_, (d_131_v43_) - (iife18_()
            ), d_133_v46_, d_126_i1_, p3)
            d_134_v47_ = nw13_
        d_136_v48_: _dafny.Array
        def lambda0_(d_137_i2_):
            return (d_137_i2_) - (len(_dafny.Map({_dafny.CodePoint('j'): 33})))

        init0_ = lambda0_
        nw14_ = _dafny.Array(None, 4)
        for i0_0_ in range(nw14_.length(0)):
            nw14_[i0_0_] = init0_(i0_0_)
        d_136_v48_ = nw14_
        d_138_v49_: _dafny.Seq
        d_138_v49_ = _dafny.SeqWithoutIsStrInference([d_136_v48_, d_136_v48_, d_136_v48_, d_136_v48_])
        d_138_v49_ = (d_138_v49_) + (d_138_v49_)
        d_125_v38_ = default__.safeDivisionInt(default__.fm3(d_125_v38_, p1, globalState), d_125_v38_)
        if p1:
            index4_ = default__.safeIndex(171, (d_136_v48_).length(0))
            (d_136_v48_)[index4_] = 946
            index5_ = default__.safeIndex(171, (d_136_v48_).length(0))
            (d_136_v48_)[index5_] = -463
            d_139_v50_: _dafny.Array
            def lambda1_(d_140_i3_):
                return _dafny.CodePoint('l')

            init1_ = lambda1_
            nw15_ = _dafny.Array(None, 4)
            for i0_1_ in range(nw15_.length(0)):
                nw15_[i0_1_] = init1_(i0_1_)
            d_139_v50_ = nw15_
            d_141_v51_: str
            d_141_v51_ = _dafny.CodePoint('j')
            index6_ = default__.safeIndex(838, (d_139_v50_).length(0))
            (d_139_v50_)[index6_] = d_141_v51_
            index7_ = default__.safeIndex(838, (d_139_v50_).length(0))
            rhs7_ = d_141_v51_
            rhs8_ = not (p1) or (default__.fm0(d_141_v51_, globalState))
            rhs9_ = (d_136_v48_)[default__.safeIndex(171, (d_136_v48_).length(0))]
            rhs10_ = not (p3) or (not(not(False)))
            lhs0_ = d_139_v50_
            lhs1_ = default__.safeIndex(838, (d_139_v50_).length(0))
            lhs0_[lhs1_] = rhs7_
            r0 = rhs8_
            r1 = rhs9_
            r0 = rhs10_
            index8_ = default__.safeIndex(838, (d_139_v50_).length(0))
            (d_139_v50_)[index8_] = _dafny.CodePoint('b')
            d_142_v52_: _dafny.Seq
            d_142_v52_ = _dafny.SeqWithoutIsStrInference([p1])
            d_143_v53_: _dafny.Seq
            d_143_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjpsw"))
            d_144_v54_: _dafny.Set
            d_144_v54_ = _dafny.Set({d_142_v52_, (_dafny.SeqWithoutIsStrInference([False, p3])).set(default__.safeIndex(len(d_143_v53_), len(_dafny.SeqWithoutIsStrInference([False, p3]))), p3)})
            d_145_v55_: D11
            d_145_v55_ = D11_DC24(p1, p3, (d_136_v48_)[default__.safeIndex(171, (d_136_v48_).length(0))])
            d_146_v56_: D1
            d_146_v56_ = D1_DC3(p1)
            d_147_v57_: _dafny.Array
            nw16_ = _dafny.Array(None, 27)
            nw16_[int(0)] = p3
            nw16_[int(1)] = p3
            nw16_[int(2)] = p3
            nw16_[int(3)] = p1
            nw16_[int(4)] = True
            nw16_[int(5)] = (d_145_v55_).cf36
            nw16_[int(6)] = p1
            nw16_[int(7)] = p3
            nw16_[int(8)] = p3
            nw16_[int(9)] = (d_146_v56_).cf2
            nw16_[int(10)] = not(p3)
            nw16_[int(11)] = p3
            nw16_[int(12)] = p1
            nw16_[int(13)] = p3
            nw16_[int(14)] = p3
            nw16_[int(15)] = p3
            nw16_[int(16)] = p3
            nw16_[int(17)] = p1
            nw16_[int(18)] = p1
            nw16_[int(19)] = not(p1)
            nw16_[int(20)] = p3
            nw16_[int(21)] = False
            nw16_[int(22)] = p1
            nw16_[int(23)] = p3
            nw16_[int(24)] = not(p1)
            nw16_[int(25)] = True
            nw16_[int(26)] = p1
            d_147_v57_ = nw16_
            d_148_v58_: _dafny.Array
            nw17_ = _dafny.Array(None, 8)
            nw17_[int(0)] = d_147_v57_
            nw17_[int(1)] = d_147_v57_
            nw17_[int(2)] = d_147_v57_
            nw17_[int(3)] = d_147_v57_
            nw17_[int(4)] = d_147_v57_
            nw17_[int(5)] = d_147_v57_
            nw17_[int(6)] = d_147_v57_
            nw17_[int(7)] = d_147_v57_
            d_148_v58_ = nw17_
            d_149_v59_: _dafny.Seq
            d_149_v59_ = _dafny.SeqWithoutIsStrInference([d_148_v58_])
            d_150_v60_: _dafny.Set
            d_150_v60_ = _dafny.Set({d_125_v38_})
            d_151_v61_: D12
            d_151_v61_ = D12_DC28(len(d_150_v60_), d_125_v38_, p3, p3, 516)
            pat_let_tv0_ = d_125_v38_
            d_152_v62_: C9
            nw18_ = C9()
            def iife19_(_pat_let0_0):
                def iife20_(d_154_dt__update__tmp_h0_):
                    def iife21_(_pat_let1_0):
                        def iife22_(d_155_dt__update_hcf44_h0_):
                            def iife23_(_pat_let2_0):
                                def iife24_(d_156_dt__update_hcf45_h0_):
                                    return D12_DC28((d_154_dt__update__tmp_h0_).cf43, d_155_dt__update_hcf44_h0_, d_156_dt__update_hcf45_h0_, (d_154_dt__update__tmp_h0_).cf46, (d_154_dt__update__tmp_h0_).cf47)
                                return iife24_(_pat_let2_0)
                            return iife23_(True)
                        return iife22_(_pat_let1_0)
                    return iife21_(pat_let_tv0_)
                return iife20_(_pat_let0_0)
            nw18_.ctor__((len(_dafny.SeqWithoutIsStrInference([(d_139_v50_)[default__.safeIndex(838, (d_139_v50_).length(0))] for d_153_i4_ in range(default__.abs(544))]))) + ((d_136_v48_)[default__.safeIndex(171, (d_136_v48_).length(0))]), d_144_v54_, (d_149_v59_)[default__.safeIndex(d_125_v38_, len(d_149_v59_))], (d_136_v48_)[default__.safeIndex(171, (d_136_v48_).length(0))], (iife19_(d_151_v61_)).cf46)
            d_152_v62_ = nw18_
            d_157_v63_: _dafny.Seq
            d_157_v63_ = _dafny.SeqWithoutIsStrInference([d_152_v62_.f7])
            d_158_v64_: _dafny.Seq
            d_158_v64_ = _dafny.SeqWithoutIsStrInference([(d_136_v48_)[default__.safeIndex(171, (d_136_v48_).length(0))], d_152_v62_.f7, (0) - (d_152_v62_.f7), d_152_v62_.f7, (d_157_v63_)[default__.safeIndex(339, len(d_157_v63_))]])
            d_159_v65_: D9
            d_159_v65_ = D9_DC20(p3, (d_139_v50_)[default__.safeIndex(838, (d_139_v50_).length(0))], d_152_v62_.f7)
            d_160_v66_: _dafny.Array
            nw19_ = _dafny.Array(None, 13)
            nw19_[int(0)] = d_125_v38_
            nw19_[int(1)] = default__.safeModuloInt(len(_dafny.Map({975: False})), d_152_v62_.f7)
            nw19_[int(2)] = d_152_v62_.f7
            nw19_[int(3)] = d_125_v38_
            nw19_[int(4)] = d_152_v62_.f7
            nw19_[int(5)] = (0) - (d_125_v38_)
            nw19_[int(6)] = len((d_157_v63_) + (_dafny.SeqWithoutIsStrInference([(d_136_v48_)[default__.safeIndex(171, (d_136_v48_).length(0))] for d_161_i5_ in range(default__.abs(729))])))
            nw19_[int(7)] = (d_136_v48_)[default__.safeIndex(171, (d_136_v48_).length(0))]
            nw19_[int(8)] = default__.fm3((d_159_v65_).cf28, p3, globalState)
            nw19_[int(9)] = d_152_v62_.f7
            nw19_[int(10)] = (d_136_v48_)[default__.safeIndex(171, (d_136_v48_).length(0))]
            nw19_[int(11)] = (d_125_v38_ if p1 else 715)
            nw19_[int(12)] = d_152_v62_.f7
            d_160_v66_ = nw19_
            index9_ = default__.safeIndex(171, (d_136_v48_).length(0))
            rhs11_ = d_152_v62_
            rhs12_ = default__.fm3(d_125_v38_, p3, globalState)
            rhs13_ = d_160_v66_
            rhs14_ = p1
            lhs2_ = d_136_v48_
            lhs3_ = default__.safeIndex(171, (d_136_v48_).length(0))
            d_152_v62_ = rhs11_
            lhs2_[lhs3_] = rhs12_
            d_160_v66_ = rhs13_
            r0 = rhs14_
            d_162_v67_: C6
            nw20_ = C6()
            nw20_.ctor__(p3, p3, d_148_v58_, -356)
            d_162_v67_ = nw20_
            d_163_v68_: _dafny.Set
            d_163_v68_ = _dafny.Set({d_162_v67_})
            d_164_v69_: _dafny.Map
            d_164_v69_ = _dafny.Map({d_163_v68_: p1})
            index10_ = default__.safeIndex(171, (d_136_v48_).length(0))
            index11_ = default__.safeIndex(171, (d_136_v48_).length(0))
            rhs15_ = (d_125_v38_) * (default__.fm3(d_125_v38_, False, globalState))
            rhs16_ = (d_125_v38_) - ((d_136_v48_)[default__.safeIndex(171, (d_136_v48_).length(0))])
            rhs17_ = ((d_136_v48_)[default__.safeIndex(171, (d_136_v48_).length(0))]) + (default__.safeModuloInt(d_152_v62_.f7, d_152_v62_.f7))
            rhs18_ = ((len(d_164_v69_) if not(p1) else (0) - ((d_136_v48_)[default__.safeIndex(171, (d_136_v48_).length(0))]))) + ((d_125_v38_) + (d_152_v62_.f7))
            lhs4_ = d_136_v48_
            lhs5_ = default__.safeIndex(171, (d_136_v48_).length(0))
            lhs6_ = d_136_v48_
            lhs7_ = default__.safeIndex(171, (d_136_v48_).length(0))
            lhs4_[lhs5_] = rhs15_
            d_125_v38_ = rhs16_
            lhs6_[lhs7_] = rhs17_
            r1 = rhs18_
        elif True:
            d_165_v70_: str
            d_165_v70_ = _dafny.CodePoint('g')
            d_166_v71_: _dafny.Map
            d_166_v71_ = _dafny.Map({d_165_v70_: d_125_v38_})
            d_167_v72_: _dafny.Seq
            d_167_v72_ = _dafny.SeqWithoutIsStrInference([d_125_v38_, d_125_v38_, ((d_166_v71_)[d_165_v70_] if (d_165_v70_) in (d_166_v71_) else d_125_v38_), len(_dafny.SeqWithoutIsStrInference([d_125_v38_ for d_168_i6_ in range(default__.abs(540))]))])
            d_169_v73_: D1
            d_169_v73_ = D1_DC4(d_125_v38_, d_125_v38_)
            d_170_v74_: _dafny.Map
            d_170_v74_ = _dafny.Map({p3: d_125_v38_})
            d_171_v75_: _dafny.Set
            d_171_v75_ = _dafny.Set({default__.safeDivisionInt(len(d_167_v72_), d_125_v38_), d_125_v38_, (d_169_v73_).cf4, ((d_170_v74_)[p1] if (p1) in (d_170_v74_) else d_125_v38_)})
            d_171_v75_ = ((d_171_v75_ if not(p1) else d_171_v75_)) - ((d_171_v75_).intersection(d_171_v75_))
            d_172_v76_: _dafny.Map
            d_172_v76_ = _dafny.Map({p1: p1})
            d_173_v77_: _dafny.MultiSet
            d_173_v77_ = _dafny.MultiSet([p3, p3])
            d_174_v78_: D11
            d_174_v78_ = D11_DC24(default__.fm0(d_165_v70_, globalState), ((d_172_v76_)[False] if (False) in (d_172_v76_) else p3), (d_173_v77_).cardinality)
            d_175_v79_: _dafny.MultiSet
            d_175_v79_ = _dafny.MultiSet([d_174_v78_])
            d_176_v80_: _dafny.Array
            nw21_ = _dafny.Array(None, 26)
            nw21_[int(0)] = p3
            nw21_[int(1)] = p3
            nw21_[int(2)] = p1
            nw21_[int(3)] = p1
            nw21_[int(4)] = p1
            nw21_[int(5)] = p1
            nw21_[int(6)] = p3
            nw21_[int(7)] = default__.fm0(d_165_v70_, globalState)
            nw21_[int(8)] = not(False)
            nw21_[int(9)] = p3
            nw21_[int(10)] = p3
            nw21_[int(11)] = p3
            nw21_[int(12)] = p3
            nw21_[int(13)] = not(p1)
            nw21_[int(14)] = p1
            nw21_[int(15)] = not(p3)
            nw21_[int(16)] = True
            nw21_[int(17)] = p3
            nw21_[int(18)] = False
            nw21_[int(19)] = p1
            nw21_[int(20)] = p3
            nw21_[int(21)] = p3
            nw21_[int(22)] = p1
            nw21_[int(23)] = p1
            nw21_[int(24)] = p3
            nw21_[int(25)] = p1
            d_176_v80_ = nw21_
            d_177_v81_: _dafny.Array
            nw22_ = _dafny.Array(None, 26)
            nw22_[int(0)] = d_176_v80_
            nw22_[int(1)] = d_176_v80_
            nw22_[int(2)] = d_176_v80_
            nw22_[int(3)] = d_176_v80_
            nw22_[int(4)] = d_176_v80_
            nw22_[int(5)] = d_176_v80_
            nw22_[int(6)] = d_176_v80_
            nw22_[int(7)] = d_176_v80_
            nw22_[int(8)] = d_176_v80_
            nw22_[int(9)] = d_176_v80_
            nw22_[int(10)] = d_176_v80_
            nw22_[int(11)] = d_176_v80_
            nw22_[int(12)] = d_176_v80_
            nw22_[int(13)] = d_176_v80_
            nw22_[int(14)] = d_176_v80_
            nw22_[int(15)] = d_176_v80_
            nw22_[int(16)] = d_176_v80_
            nw22_[int(17)] = d_176_v80_
            nw22_[int(18)] = d_176_v80_
            nw22_[int(19)] = d_176_v80_
            nw22_[int(20)] = d_176_v80_
            nw22_[int(21)] = d_176_v80_
            nw22_[int(22)] = d_176_v80_
            nw22_[int(23)] = d_176_v80_
            nw22_[int(24)] = d_176_v80_
            nw22_[int(25)] = d_176_v80_
            d_177_v81_ = nw22_
            d_178_v82_: C5
            nw23_ = C5()
            nw23_.ctor__((p3) == (p3), (d_175_v79_).cardinality, d_177_v81_, d_125_v38_, p3)
            d_178_v82_ = nw23_
            d_178_v82_ = d_178_v82_
            d_179_v83_: _dafny.Array
            nw24_ = _dafny.Array(_dafny.Seq({}), 21)
            d_179_v83_ = nw24_
            d_180_v84_: _dafny.Map
            d_180_v84_ = _dafny.Map({d_179_v83_: (d_178_v82_).f13})
            d_180_v84_ = (d_180_v84_).set(d_179_v83_, p3)
            d_181_v85_: _dafny.Seq
            d_181_v85_ = _dafny.SeqWithoutIsStrInference([p1])
            d_182_v86_: _dafny.Set
            d_182_v86_ = _dafny.Set({d_181_v85_})
            d_183_v87_: C9
            nw25_ = C9()
            nw25_.ctor__((0) - (len(_dafny.SeqWithoutIsStrInference([d_165_v70_ for d_184_i7_ in range(default__.abs(-613))]))), d_182_v86_, d_177_v81_, d_178_v82_.f14, p1)
            d_183_v87_ = nw25_
            d_185_v88_: _dafny.Map
            d_185_v88_ = _dafny.Map({d_183_v87_: p3})
            d_185_v88_ = (d_185_v88_).set((d_183_v87_ if (d_178_v82_).f13 else d_183_v87_), p1)
            d_186_v89_: _dafny.Seq
            d_186_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xu"))
            d_187_v90_: T0
            nw26_ = C2()
            nw26_.ctor__(d_186_v89_, p1)
            d_187_v90_ = nw26_
            d_188_v91_: T0
            d_188_v91_ = d_187_v90_
            d_188_v91_ = d_188_v91_
        if p3:
            d_189_v92_: str
            d_189_v92_ = _dafny.CodePoint('k')
            d_190_v93_: _dafny.Seq
            d_190_v93_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvok"))
            d_191_v94_: D9
            d_191_v94_ = D9_DC19(d_190_v93_)
            d_192_v95_: T0
            nw27_ = C2()
            nw27_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_193_i8_ in range(default__.abs(-339))]), (d_189_v92_) in ((d_191_v94_).cf25))
            d_192_v95_ = nw27_
            rhs19_ = d_192_v95_
            rhs20_ = (0) - (d_125_v38_)
            d_192_v95_ = rhs19_
            d_125_v38_ = rhs20_
            d_194_v96_: _dafny.MultiSet
            d_194_v96_ = _dafny.MultiSet([(d_192_v95_).f4, p3])
            d_195_v97_: _dafny.Map
            d_195_v97_ = _dafny.Map({p3: p1})
            d_196_v98_: _dafny.Seq
            d_196_v98_ = _dafny.SeqWithoutIsStrInference([d_195_v97_, d_195_v97_])
            source7_ = D16_DC33((d_192_v95_).f4, d_125_v38_, (d_194_v96_) - (d_194_v96_), d_125_v38_, default__.fm53(d_125_v38_, d_196_v98_, globalState))
            if source7_.is_DC33:
                d_197___mcc_h0_ = source7_.cf52
                d_198___mcc_h1_ = source7_.cf53
                d_199___mcc_h2_ = source7_.cf54
                d_200___mcc_h3_ = source7_.cf55
                d_201___mcc_h4_ = source7_.cf56
                d_202_cf56_ = d_201___mcc_h4_
                d_203_cf55_ = d_200___mcc_h3_
                d_204_cf54_ = d_199___mcc_h2_
                d_205_cf53_ = d_198___mcc_h1_
                d_206_cf52_ = d_197___mcc_h0_
                d_207_v99_: _dafny.Array
                nw28_ = _dafny.Array(None, 20)
                d_207_v99_ = nw28_
                d_208_v100_: _dafny.Array
                nw29_ = _dafny.Array(_dafny.Array(None, 0), 14)
                d_208_v100_ = nw29_
                d_209_v101_: _dafny.Seq
                d_209_v101_ = _dafny.SeqWithoutIsStrInference([d_205_cf53_])
                d_210_v102_: C4
                nw30_ = C4()
                nw30_.ctor__(d_208_v100_, (d_209_v101_)[default__.safeIndex(len(d_209_v101_), len(d_209_v101_))], p3)
                d_210_v102_ = nw30_
                index12_ = default__.safeIndex(60, (d_207_v99_).length(0))
                (d_207_v99_)[index12_] = d_210_v102_
                d_211_v103_: _dafny.Seq
                d_211_v103_ = _dafny.SeqWithoutIsStrInference([not(p3)])
                d_212_v106_: _dafny.MultiSet
                def iife25_():
                    coll19_ = _dafny.Map()
                    compr_19_: int
                    for compr_19_ in _dafny.IntegerRange(438, 391):
                        d_213_v104_: int = compr_19_
                        if ((438) <= (d_213_v104_)) and ((d_213_v104_) < (391)):
                            coll19_[default__.safeDivisionInt(d_213_v104_, len(d_195_v97_))] = p1
                    return _dafny.Map(coll19_)
                def iife26_():
                    coll20_ = _dafny.Map()
                    compr_20_: int
                    for compr_20_ in _dafny.IntegerRange(355, 267):
                        d_214_v105_: int = compr_20_
                        if ((355) <= (d_214_v105_)) and ((d_214_v105_) < (267)):
                            coll20_[(d_214_v105_) * (d_125_v38_)] = p1
                    return _dafny.Map(coll20_)
                d_212_v106_ = _dafny.MultiSet([iife25_()
                , iife26_()
                ])
                index13_ = default__.safeIndex(60, (d_207_v99_).length(0))
                rhs21_ = d_205_cf53_
                rhs22_ = (d_211_v103_)[default__.safeIndex((d_212_v106_).cardinality, len(d_211_v103_))]
                rhs23_ = (d_189_v92_ if (d_210_v102_).fm8(d_205_cf53_, d_206_cf52_, _dafny.SeqWithoutIsStrInference([d_189_v92_ for d_215_i9_ in range(default__.abs(263))]), d_209_v101_, globalState) else _dafny.CodePoint('p'))
                rhs24_ = d_210_v102_
                lhs8_ = d_207_v99_
                lhs9_ = default__.safeIndex(60, (d_207_v99_).length(0))
                r1 = rhs21_
                r0 = rhs22_
                d_189_v92_ = rhs23_
                lhs8_[lhs9_] = rhs24_
                d_216_v107_: _dafny.Array
                nw31_ = _dafny.Array(None, 3)
                d_216_v107_ = nw31_
                d_217_v108_: _dafny.Set
                d_217_v108_ = _dafny.Set({d_211_v103_, d_211_v103_, _dafny.SeqWithoutIsStrInference([(d_192_v95_).f4]), d_211_v103_, d_211_v103_})
                d_218_v109_: C9
                nw32_ = C9()
                nw32_.ctor__(d_205_cf53_, d_217_v108_, d_208_v100_, d_203_cf55_, d_206_cf52_)
                d_218_v109_ = nw32_
                index14_ = default__.safeIndex(913, (d_216_v107_).length(0))
                (d_216_v107_)[index14_] = d_218_v109_
                index15_ = default__.safeIndex(913, (d_216_v107_).length(0))
                (d_216_v107_)[index15_] = d_218_v109_
                d_206_cf52_ = not(False)
                d_206_cf52_ = ((0) - ((d_210_v102_).fm7(globalState))) < (d_205_cf53_)
            elif True:
                d_219___mcc_h5_ = source7_.cf51
                d_220_cf51_ = d_219___mcc_h5_
                d_189_v92_ = (d_189_v92_ if ((d_192_v95_).f4 if p1 else (d_192_v95_).f4) else default__.fm1(d_125_v38_, (d_192_v95_).f4, globalState))
                d_221_v110_: _dafny.Array
                nw33_ = _dafny.Array(False, 20)
                d_221_v110_ = nw33_
                d_222_v111_: _dafny.Seq
                d_222_v111_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_221_v110_, d_221_v110_, d_221_v110_, d_221_v110_, d_221_v110_}))])
                r0 = ((d_222_v111_)[default__.safeIndex(d_125_v38_, len(d_222_v111_))]) > ((0) - ((default__.fm3(d_125_v38_, (d_192_v95_).f4, globalState) if (d_192_v95_).f4 else d_125_v38_)))
                d_223_v112_: _dafny.Array
                nw34_ = _dafny.Array(_dafny.Set({}), 19)
                d_223_v112_ = nw34_
                d_224_v113_: _dafny.Array
                nw35_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_224_v113_ = nw35_
                d_225_v114_: _dafny.Array
                d_225_v114_ = d_224_v113_
                d_226_v115_: _dafny.Set
                d_226_v115_ = _dafny.Set({d_225_v114_, d_225_v114_})
                d_227_v116_: _dafny.Set
                d_227_v116_ = _dafny.Set({_dafny.Set({d_225_v114_}), d_226_v115_})
                index16_ = default__.safeIndex(506, (d_223_v112_).length(0))
                (d_223_v112_)[index16_] = d_227_v116_
                index17_ = default__.safeIndex(506, (d_223_v112_).length(0))
                (d_223_v112_)[index17_] = _dafny.Set({d_226_v115_})
                r1 = (0) - (d_125_v38_)
            d_228_v117_: _dafny.Map
            d_228_v117_ = _dafny.Map({d_125_v38_: d_136_v48_})
            d_228_v117_ = (d_228_v117_).set(d_125_v38_, d_136_v48_)
            d_125_v38_ = d_125_v38_
            r0 = p1
        elif True:
            r0 = (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_229_i10_ in range(default__.abs(863))]))) >= (d_125_v38_)
            d_230_v118_: _dafny.Array
            nw36_ = _dafny.Array(_dafny.Map({}), 1)
            d_230_v118_ = nw36_
            d_231_v119_: D22
            d_231_v119_ = D22_DC45(d_230_v118_)
            d_231_v119_ = D22_DC45(d_230_v118_)
            d_232_v120_: _dafny.Map
            d_232_v120_ = _dafny.Map({d_125_v38_: p3})
            d_233_v122_: _dafny.Array
            nw37_ = _dafny.Array(_dafny.Array(None, 0), 22)
            d_233_v122_ = nw37_
            d_234_v123_: _dafny.Seq
            d_234_v123_ = _dafny.SeqWithoutIsStrInference([d_125_v38_, d_125_v38_])
            d_235_v124_: C1
            nw38_ = C1()
            def iife27_():
                coll21_ = _dafny.Set()
                compr_21_: int
                for compr_21_ in (d_232_v120_).keys.Elements:
                    d_236_v121_: int = compr_21_
                    if (d_236_v121_) in (d_232_v120_):
                        coll21_ = coll21_.union(_dafny.Set([(d_236_v121_) + (len(_dafny.Set({True, True})))]))
                return _dafny.Set(coll21_)
            nw38_.ctor__(len(iife27_()
            ), d_125_v38_, d_233_v122_, len(d_234_v123_), p1)
            d_235_v124_ = nw38_
            index18_ = default__.safeIndex(240, (d_136_v48_).length(0))
            (d_136_v48_)[index18_] = d_125_v38_
            index19_ = default__.safeIndex(240, (d_136_v48_).length(0))
            (d_136_v48_)[index19_] = d_125_v38_
            d_125_v38_ = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_237_i11_ in range(default__.abs(623))]))
        d_238_v125_: _dafny.Seq
        d_238_v125_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjhatt"))
        d_239_v126_: _dafny.Array
        nw39_ = _dafny.Array(False, 22)
        d_239_v126_ = nw39_
        d_240_v127_: _dafny.Seq
        d_240_v127_ = _dafny.SeqWithoutIsStrInference([d_239_v126_, d_239_v126_])
        d_241_v128_: D20
        d_241_v128_ = D20_DC41(d_125_v38_)
        d_242_v129_: D10
        d_242_v129_ = D10_DC22(d_238_v125_, p1, p3, (d_240_v127_)[default__.safeIndex(d_125_v38_, len(d_240_v127_))], (d_241_v128_).cf67)
        d_243_v130_: _dafny.Map
        d_243_v130_ = _dafny.Map({d_125_v38_: p3})
        d_244_v131_: _dafny.Map
        d_244_v131_ = _dafny.Map({p1: p1})
        d_245_v132_: _dafny.Array
        nw40_ = _dafny.Array(_dafny.Array(None, 0), 24)
        d_245_v132_ = nw40_
        d_246_v133_: C6
        nw41_ = C6()
        nw41_.ctor__(((d_243_v130_)[d_125_v38_] if (d_125_v38_) in (d_243_v130_) else p1), not(((d_244_v131_)[p1] if (p1) in (d_244_v131_) else p1)), d_245_v132_, default__.fm3(d_125_v38_, False, globalState))
        d_246_v133_ = nw41_
        d_247_v134_: _dafny.Map
        d_247_v134_ = _dafny.Map({d_246_v133_: d_125_v38_})
        r0 = (len((d_242_v129_).cf30)) >= (len(d_247_v134_))
        d_248_v135_: str
        d_248_v135_ = _dafny.CodePoint('t')
        r1 = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewmfdfes"))).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_249_i12_ in range(default__.abs(714))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewmfdfes")))), d_248_v135_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nv"))))
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_250_v0_: bool
        d_250_v0_ = False
        d_251_v1_: _dafny.MultiSet
        d_251_v1_ = _dafny.MultiSet([d_250_v0_, d_250_v0_, d_250_v0_])
        d_252_globalState_: GlobalState
        nw42_ = GlobalState()
        nw42_.ctor__(d_251_v1_, -993, False, 862)
        d_252_globalState_ = nw42_
        d_253_v2_: int
        d_253_v2_ = 598
        d_254_v3_: _dafny.Array
        nw43_ = _dafny.Array(None, 14)
        nw43_[int(0)] = d_253_v2_
        nw43_[int(1)] = d_253_v2_
        nw43_[int(2)] = d_253_v2_
        nw43_[int(3)] = (d_253_v2_ if d_250_v0_ else d_253_v2_)
        nw43_[int(4)] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_255_i0_ in range(default__.abs(532))])), 334)
        nw43_[int(5)] = -930
        nw43_[int(6)] = d_253_v2_
        nw43_[int(7)] = d_253_v2_
        nw43_[int(8)] = d_253_v2_
        nw43_[int(9)] = ((D0_DC2(d_253_v2_)).cf1) - ((0) - (d_253_v2_))
        nw43_[int(10)] = d_253_v2_
        nw43_[int(11)] = (d_253_v2_) - (d_253_v2_)
        nw43_[int(12)] = ((0) - ((0) - (d_253_v2_))) * (743)
        nw43_[int(13)] = d_253_v2_
        d_254_v3_ = nw43_
        d_254_v3_ = d_254_v3_
        d_256_v4_: str
        d_256_v4_ = _dafny.CodePoint('j')
        d_256_v4_ = d_256_v4_
        d_257_v5_: _dafny.Map
        d_257_v5_ = _dafny.Map({(d_253_v2_) * (d_253_v2_): d_250_v0_})
        d_257_v5_ = (d_257_v5_).set((d_253_v2_) - (d_253_v2_), d_250_v0_)
        d_258_v6_: _dafny.Set
        d_258_v6_ = _dafny.Set({d_253_v2_, d_253_v2_})
        d_259_v7_: _dafny.Set
        d_259_v7_ = _dafny.Set({d_258_v6_})
        d_260_v8_: _dafny.Seq
        d_260_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smw"))
        d_261_v9_: _dafny.Map
        d_261_v9_ = _dafny.Map({d_259_v7_: (len(d_260_v8_)) != (d_253_v2_)})
        d_250_v0_ = ((d_261_v9_)[d_259_v7_] if (d_259_v7_) in (d_261_v9_) else d_250_v0_)
        d_262_v10_: _dafny.Array
        nw44_ = _dafny.Array(_dafny.Set({}), 2)
        d_262_v10_ = nw44_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_262_v10_).length(0)):
            d_263_i1_: int = guard_loop_0_
            if (True) and (((0) <= (d_263_i1_)) and ((d_263_i1_) < ((d_262_v10_).length(0)))):
                (d_262_v10_)[(d_263_i1_)] = ((_dafny.Set({d_250_v0_})) - (_dafny.Set({d_250_v0_, False}))) - ((_dafny.Set({d_250_v0_, True, d_250_v0_, False})).intersection(_dafny.Set({d_250_v0_, d_250_v0_, (D1_DC3(d_250_v0_)).cf2, d_250_v0_, d_250_v0_})))
        d_250_v0_ = not (d_250_v0_) or (d_250_v0_)
        d_264_v11_: _dafny.Array
        def lambda2_(d_265_v0_):
            def lambda3_(d_266_i2_):
                return d_265_v0_

            return lambda3_

        init2_ = lambda2_(d_250_v0_)
        nw45_ = _dafny.Array(None, 6)
        for i0_2_ in range(nw45_.length(0)):
            nw45_[i0_2_] = init2_(i0_2_)
        d_264_v11_ = nw45_
        d_267_v12_: _dafny.Seq
        d_267_v12_ = _dafny.SeqWithoutIsStrInference([d_264_v11_, d_264_v11_])
        d_268_v13_: _dafny.Map
        d_268_v13_ = _dafny.Map({not(d_250_v0_): d_256_v4_})
        d_269_v14_: _dafny.Seq
        d_269_v14_ = _dafny.SeqWithoutIsStrInference([d_253_v2_])
        d_270_v15_: _dafny.Array
        nw46_ = _dafny.Array(None, 21)
        nw46_[int(0)] = not(False)
        nw46_[int(1)] = default__.fm0(default__.fm1((0) - (d_253_v2_), d_250_v0_, d_252_globalState_), d_252_globalState_)
        nw46_[int(2)] = d_250_v0_
        nw46_[int(3)] = d_250_v0_
        nw46_[int(4)] = not(default__.fm0(_dafny.CodePoint('e'), d_252_globalState_))
        nw46_[int(5)] = d_250_v0_
        nw46_[int(6)] = d_250_v0_
        nw46_[int(7)] = not(d_250_v0_)
        nw46_[int(8)] = d_250_v0_
        nw46_[int(9)] = not(d_250_v0_)
        nw46_[int(10)] = d_250_v0_
        nw46_[int(11)] = (d_264_v11_) in (d_267_v12_)
        nw46_[int(12)] = d_250_v0_
        nw46_[int(13)] = d_250_v0_
        nw46_[int(14)] = d_250_v0_
        nw46_[int(15)] = (d_253_v2_) != (((d_251_v1_)[d_250_v0_] if (d_250_v0_) in (d_251_v1_) else len(d_260_v8_)))
        nw46_[int(16)] = (d_250_v0_ if d_250_v0_ else not(d_250_v0_))
        nw46_[int(17)] = d_250_v0_
        nw46_[int(18)] = d_250_v0_
        nw46_[int(19)] = not(default__.fm0(((d_268_v13_)[d_250_v0_] if (d_250_v0_) in (d_268_v13_) else d_256_v4_), d_252_globalState_))
        nw46_[int(20)] = (len(d_260_v8_)) < (len(d_269_v14_))
        d_270_v15_ = nw46_
        index20_ = default__.safeIndex(99, (d_270_v15_).length(0))
        (d_270_v15_)[index20_] = False
        d_271_v16_: D0
        d_271_v16_ = D0_DC1()
        d_272_v17_: _dafny.Seq
        d_272_v17_ = _dafny.SeqWithoutIsStrInference([d_271_v16_, default__.fm2(d_252_globalState_)])
        index21_ = default__.safeIndex(99, (d_270_v15_).length(0))
        rhs25_ = not((_dafny.SeqWithoutIsStrInference([D0_DC1()])) == ((d_272_v17_) + (d_272_v17_)))
        rhs26_ = default__.fm0(d_256_v4_, d_252_globalState_)
        lhs10_ = d_270_v15_
        lhs11_ = default__.safeIndex(99, (d_270_v15_).length(0))
        d_250_v0_ = rhs25_
        lhs10_[lhs11_] = rhs26_
        d_273_v18_: _dafny.MultiSet
        d_273_v18_ = _dafny.MultiSet([d_253_v2_])
        d_274_v19_: _dafny.Map
        d_274_v19_ = _dafny.Map({d_273_v18_: d_260_v8_})
        d_275_v20_: bool
        d_276_v21_: int
        out0_: bool
        out1_: int
        out0_, out1_ = default__.m0(d_274_v19_, (not(d_250_v0_) if (d_270_v15_)[default__.safeIndex(99, (d_270_v15_).length(0))] else (d_270_v15_)[default__.safeIndex(99, (d_270_v15_).length(0))]), (d_273_v18_).set(d_253_v2_, default__.abs(d_253_v2_)), (d_270_v15_)[default__.safeIndex(99, (d_270_v15_).length(0))], d_252_globalState_)
        d_275_v20_ = out0_
        d_276_v21_ = out1_
        d_277_v22_: _dafny.Array
        def lambda4_(d_278_v14_):
            def lambda5_(d_279_i3_):
                return (d_278_v14_) + (d_278_v14_)

            return lambda5_

        init3_ = lambda4_(d_269_v14_)
        nw47_ = _dafny.Array(None, 24)
        for i0_3_ in range(nw47_.length(0)):
            nw47_[i0_3_] = init3_(i0_3_)
        d_277_v22_ = nw47_
        index22_ = default__.safeIndex(766, (d_277_v22_).length(0))
        (d_277_v22_)[index22_] = d_269_v14_
        d_280_v23_: D0
        d_280_v23_ = D0_DC2(d_253_v2_)
        pat_let_tv1_ = d_271_v16_
        pat_let_tv2_ = d_271_v16_
        pat_let_tv3_ = d_271_v16_
        index23_ = default__.safeIndex(99, (d_270_v15_).length(0))
        index24_ = default__.safeIndex(766, (d_277_v22_).length(0))
        def lambda6_(source8_):
            if source8_.is_DC1:
                return pat_let_tv1_
            elif source8_.is_DC2:
                d_281___mcc_h0_ = source8_.cf1
                d_282_cf1_ = d_281___mcc_h0_
                return pat_let_tv2_
            elif True:
                d_283___mcc_h1_ = source8_.cf0
                d_284_cf0_ = d_283___mcc_h1_
                return pat_let_tv3_

        rhs27_ = lambda6_(d_280_v23_)
        rhs28_ = True
        rhs29_ = default__.fm0(d_256_v4_, d_252_globalState_)
        rhs30_ = d_269_v14_
        rhs31_ = d_253_v2_
        lhs12_ = d_270_v15_
        lhs13_ = default__.safeIndex(99, (d_270_v15_).length(0))
        lhs14_ = d_277_v22_
        lhs15_ = default__.safeIndex(766, (d_277_v22_).length(0))
        d_271_v16_ = rhs27_
        lhs12_[lhs13_] = rhs28_
        d_275_v20_ = rhs29_
        lhs14_[lhs15_] = rhs30_
        d_253_v2_ = rhs31_
        d_280_v23_ = d_280_v23_
        if (d_270_v15_)[default__.safeIndex(99, (d_270_v15_).length(0))]:
            d_285_v24_: _dafny.Map
            d_285_v24_ = _dafny.Map({d_253_v2_: d_254_v3_})
            d_254_v3_ = ((d_285_v24_)[d_253_v2_] if (d_253_v2_) in (d_285_v24_) else d_254_v3_)
            index25_ = default__.safeIndex(124, (d_254_v3_).length(0))
            (d_254_v3_)[index25_] = (d_276_v21_) * (default__.fm3(d_253_v2_, (d_270_v15_)[default__.safeIndex(99, (d_270_v15_).length(0))], d_252_globalState_))
            d_286_v25_: _dafny.Seq
            d_286_v25_ = _dafny.SeqWithoutIsStrInference([d_273_v18_, (d_273_v18_) | (_dafny.MultiSet([d_253_v2_, 673]))])
            index26_ = default__.safeIndex(124, (d_254_v3_).length(0))
            rhs32_ = not(d_250_v0_)
            rhs33_ = default__.safeDivisionInt(d_276_v21_, d_253_v2_)
            rhs34_ = (d_286_v25_) + (default__.fm4(False, True, d_252_globalState_))
            rhs35_ = d_280_v23_
            lhs16_ = d_254_v3_
            lhs17_ = default__.safeIndex(124, (d_254_v3_).length(0))
            d_275_v20_ = rhs32_
            lhs16_[lhs17_] = rhs33_
            d_286_v25_ = rhs34_
            d_280_v23_ = rhs35_
            index27_ = default__.safeIndex(124, (d_254_v3_).length(0))
            (d_254_v3_)[index27_] = 902
            index28_ = default__.safeIndex(99, (d_270_v15_).length(0))
            (d_270_v15_)[index28_] = d_275_v20_
            d_287_v26_: D2
            d_287_v26_ = D2_DC5(d_270_v15_)
            d_264_v11_ = (d_287_v26_).cf5
        elif True:
            index29_ = default__.safeIndex(547, (d_254_v3_).length(0))
            (d_254_v3_)[index29_] = len((d_260_v8_ if d_275_v20_ else d_260_v8_))
            d_288_v27_: _dafny.Array
            nw48_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
            d_288_v27_ = nw48_
            index30_ = default__.safeIndex(397, (d_288_v27_).length(0))
            (d_288_v27_)[index30_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwenxedn"))
            d_289_v28_: _dafny.Seq
            d_289_v28_ = _dafny.SeqWithoutIsStrInference([d_260_v8_, d_260_v8_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qystgqh")), _dafny.SeqWithoutIsStrInference([d_256_v4_ for d_290_i4_ in range(default__.abs(635))]), (d_260_v8_) + (d_260_v8_)])
            index31_ = default__.safeIndex(547, (d_254_v3_).length(0))
            index32_ = default__.safeIndex(397, (d_288_v27_).length(0))
            rhs36_ = len(d_289_v28_)
            rhs37_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_291_i5_ in range(default__.abs(912))])
            lhs18_ = d_254_v3_
            lhs19_ = default__.safeIndex(547, (d_254_v3_).length(0))
            lhs20_ = d_288_v27_
            lhs21_ = default__.safeIndex(397, (d_288_v27_).length(0))
            lhs18_[lhs19_] = rhs36_
            lhs20_[lhs21_] = rhs37_
            d_292_v30_: _dafny.Seq
            d_292_v30_ = _dafny.SeqWithoutIsStrInference([d_273_v18_])
            d_293_v31_: _dafny.Map
            d_293_v31_ = _dafny.Map({True: d_274_v19_})
            d_294_v32_: bool
            d_295_v33_: int
            out2_: bool
            out3_: int
            def iife28_():
                coll22_ = _dafny.Map()
                compr_22_: _dafny.MultiSet
                for compr_22_ in (d_292_v30_).Elements:
                    d_296_v29_: _dafny.MultiSet = compr_22_
                    if (d_296_v29_) in (d_292_v30_):
                        coll22_[d_296_v29_] = d_260_v8_
                return _dafny.Map(coll22_)
            out2_, out3_ = default__.m0((iife28_()
            ) | (((d_293_v31_)[d_275_v20_] if (d_275_v20_) in (d_293_v31_) else d_274_v19_)), d_250_v0_, d_273_v18_, not(d_250_v0_), d_252_globalState_)
            d_294_v32_ = out2_
            d_295_v33_ = out3_
            d_250_v0_ = (d_270_v15_)[default__.safeIndex(99, (d_270_v15_).length(0))]
            d_297_v34_: _dafny.Set
            d_297_v34_ = _dafny.Set({d_260_v8_})
            d_298_v35_: _dafny.Map
            d_298_v35_ = _dafny.Map({d_250_v0_: d_294_v32_})
            index33_ = default__.safeIndex(547, (d_254_v3_).length(0))
            index34_ = default__.safeIndex(547, (d_254_v3_).length(0))
            index35_ = default__.safeIndex(547, (d_254_v3_).length(0))
            rhs38_ = default__.fm3((len(d_297_v34_)) + (d_295_v33_), not((d_270_v15_)[default__.safeIndex(99, (d_270_v15_).length(0))]), d_252_globalState_)
            rhs39_ = d_273_v18_
            rhs40_ = d_276_v21_
            rhs41_ = (0) - (default__.safeModuloInt(d_253_v2_, (0) - ((d_295_v33_ if d_275_v20_ else len(d_298_v35_)))))
            rhs42_ = d_295_v33_
            lhs22_ = d_254_v3_
            lhs23_ = default__.safeIndex(547, (d_254_v3_).length(0))
            lhs24_ = d_254_v3_
            lhs25_ = default__.safeIndex(547, (d_254_v3_).length(0))
            lhs26_ = d_254_v3_
            lhs27_ = default__.safeIndex(547, (d_254_v3_).length(0))
            d_276_v21_ = rhs38_
            d_273_v18_ = rhs39_
            lhs22_[lhs23_] = rhs40_
            lhs24_[lhs25_] = rhs41_
            lhs26_[lhs27_] = rhs42_
            d_250_v0_ = not (d_294_v32_) or (d_294_v32_)
        d_299_v36_: _dafny.Map
        d_299_v36_ = _dafny.Map({d_271_v16_: d_276_v21_})
        d_300_v40_: _dafny.Set
        d_300_v40_ = _dafny.Set({default__.fm0(d_256_v4_, d_252_globalState_)})
        d_301_v41_: _dafny.Array
        nw49_ = _dafny.Array(None, 12)
        nw49_[int(0)] = d_299_v36_
        nw49_[int(1)] = (d_299_v36_) | (d_299_v36_)
        nw49_[int(2)] = d_299_v36_
        nw49_[int(3)] = d_299_v36_
        def iife29_():
            coll23_ = _dafny.Map()
            compr_23_: D0
            for compr_23_ in (d_299_v36_).keys.Elements:
                d_302_v37_: D0 = compr_23_
                if (d_302_v37_) in (d_299_v36_):
                    coll23_[d_302_v37_] = (d_269_v14_)[default__.safeIndex(d_253_v2_, len(d_269_v14_))]
            return _dafny.Map(coll23_)
        nw49_[int(4)] = iife29_()
        
        nw49_[int(5)] = d_299_v36_
        nw49_[int(6)] = d_299_v36_
        nw49_[int(7)] = d_299_v36_
        def iife30_():
            coll24_ = _dafny.Map()
            compr_24_: D0
            for compr_24_ in (_dafny.SeqWithoutIsStrInference([d_271_v16_, d_271_v16_])).Elements:
                d_303_v38_: D0 = compr_24_
                if (d_303_v38_) in (_dafny.SeqWithoutIsStrInference([d_271_v16_, d_271_v16_])):
                    def iife31_():
                        coll25_ = _dafny.Map()
                        compr_25_: D2
                        for compr_25_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D2_DC7(d_253_v2_)]))).Elements:
                            d_304_v39_: D2 = compr_25_
                            if (d_304_v39_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D2_DC7(d_253_v2_)]))):
                                coll25_[d_304_v39_] = d_276_v21_
                        return _dafny.Map(coll25_)
                    coll24_[d_303_v38_] = len(iife31_()
                    )
            return _dafny.Map(coll24_)
        nw49_[int(8)] = (iife30_()
        ).set(D0_DC1(), default__.fm3(len(d_300_v40_), d_250_v0_, d_252_globalState_))
        nw49_[int(9)] = d_299_v36_
        nw49_[int(10)] = default__.fm5(d_276_v21_, (d_270_v15_)[default__.safeIndex(99, (d_270_v15_).length(0))], d_252_globalState_)
        nw49_[int(11)] = d_299_v36_
        d_301_v41_ = nw49_
        d_301_v41_ = (d_301_v41_ if not(d_250_v0_) else d_301_v41_)
        d_305_v42_: _dafny.MultiSet
        d_305_v42_ = _dafny.MultiSet([d_256_v4_])
        d_306_v43_: bool
        d_307_v44_: int
        out4_: bool
        out5_: int
        out4_, out5_ = default__.m0(d_274_v19_, False, d_273_v18_, (_dafny.CodePoint('n')) in (d_305_v42_), d_252_globalState_)
        d_306_v43_ = out4_
        d_307_v44_ = out5_
        index36_ = default__.safeIndex(964, (d_254_v3_).length(0))
        (d_254_v3_)[index36_] = d_307_v44_
        d_308_v45_: _dafny.Map
        d_308_v45_ = _dafny.Map({d_254_v3_: _dafny.SeqWithoutIsStrInference([d_276_v21_ for d_309_i6_ in range(default__.abs(-750))])})
        index37_ = default__.safeIndex(964, (d_254_v3_).length(0))
        (d_254_v3_)[index37_] = (0) - ((len((d_308_v45_ if d_275_v20_ else d_308_v45_))) * (d_253_v2_))
        hi1_ = d_276_v21_
        for d_310_i7_ in range(d_307_v44_, hi1_):
            index38_ = default__.safeIndex(964, (d_254_v3_).length(0))
            (d_254_v3_)[index38_] = (d_253_v2_) + ((d_254_v3_)[default__.safeIndex(964, (d_254_v3_).length(0))])
            index39_ = default__.safeIndex(964, (d_254_v3_).length(0))
            (d_254_v3_)[index39_] = d_253_v2_
            d_311_v46_: D19
            d_311_v46_ = D19_DC38(d_275_v20_, d_306_v43_)
            d_312_v47_: _dafny.Seq
            d_312_v47_ = _dafny.SeqWithoutIsStrInference([False, False])
            d_313_v48_: _dafny.Set
            d_313_v48_ = _dafny.Set({d_312_v47_, d_312_v47_, d_312_v47_})
            d_314_v49_: _dafny.Map
            d_314_v49_ = _dafny.Map({d_250_v0_: d_270_v15_})
            d_315_v50_: D10
            d_315_v50_ = D10_DC22(d_260_v8_, d_250_v0_, d_275_v20_, d_270_v15_, 867)
            d_316_v51_: _dafny.Array
            nw50_ = _dafny.Array(None, 18)
            nw50_[int(0)] = d_264_v11_
            nw50_[int(1)] = d_264_v11_
            nw50_[int(2)] = d_270_v15_
            nw50_[int(3)] = d_270_v15_
            nw50_[int(4)] = d_270_v15_
            nw50_[int(5)] = d_264_v11_
            nw50_[int(6)] = d_264_v11_
            nw50_[int(7)] = d_270_v15_
            nw50_[int(8)] = d_264_v11_
            nw50_[int(9)] = d_264_v11_
            nw50_[int(10)] = d_270_v15_
            nw50_[int(11)] = d_264_v11_
            nw50_[int(12)] = ((d_314_v49_)[d_250_v0_] if (d_250_v0_) in (d_314_v49_) else d_264_v11_)
            nw50_[int(13)] = d_270_v15_
            nw50_[int(14)] = d_264_v11_
            nw50_[int(15)] = d_270_v15_
            nw50_[int(16)] = d_264_v11_
            nw50_[int(17)] = (d_315_v50_).cf33
            d_316_v51_ = nw50_
            d_317_v52_: C9
            nw51_ = C9()
            nw51_.ctor__(default__.fm3(d_310_i7_, (d_311_v46_).cf64, d_252_globalState_), d_313_v48_, d_316_v51_, d_253_v2_, default__.fm0(d_256_v4_, d_252_globalState_))
            d_317_v52_ = nw51_
            d_254_v3_ = d_254_v3_
        if not(not(d_275_v20_)):
            d_318_v53_: _dafny.Array
            nw52_ = _dafny.Array(None, 10)
            nw52_[int(0)] = d_264_v11_
            nw52_[int(1)] = d_264_v11_
            nw52_[int(2)] = d_270_v15_
            nw52_[int(3)] = d_270_v15_
            nw52_[int(4)] = d_264_v11_
            nw52_[int(5)] = d_264_v11_
            nw52_[int(6)] = d_264_v11_
            nw52_[int(7)] = d_264_v11_
            nw52_[int(8)] = d_270_v15_
            nw52_[int(9)] = d_270_v15_
            d_318_v53_ = nw52_
            d_319_v54_: C5
            nw53_ = C5()
            nw53_.ctor__(d_275_v20_, d_253_v2_, d_318_v53_, len(d_260_v8_), d_306_v43_)
            d_319_v54_ = nw53_
            index40_ = default__.safeIndex(964, (d_254_v3_).length(0))
            (d_254_v3_)[index40_] = len((D26_DC53(_dafny.SeqWithoutIsStrInference([d_319_v54_, d_319_v54_, d_319_v54_]))).cf89)
            d_320_v55_: _dafny.Set
            d_320_v55_ = _dafny.Set({d_256_v4_, d_256_v4_})
            d_321_v56_: D6
            d_321_v56_ = D6_DC16((d_319_v54_).f13, ((d_257_v5_)[len(d_320_v55_)] if (len(d_320_v55_)) in (d_257_v5_) else d_306_v43_), (d_277_v22_)[default__.safeIndex(766, (d_277_v22_).length(0))])
            d_322_v57_: _dafny.Seq
            d_322_v57_ = _dafny.SeqWithoutIsStrInference([True, (d_319_v54_).fm8(d_276_v21_, d_250_v0_, d_260_v8_, (d_321_v56_).cf22, d_252_globalState_)])
            d_256_v4_ = (d_260_v8_)[default__.safeIndex(default__.safeModuloInt(len(d_322_v57_), d_253_v2_), len(d_260_v8_))]
            d_323_v59_: _dafny.Seq
            d_323_v59_ = _dafny.SeqWithoutIsStrInference([d_318_v53_])
            d_324_v60_: T1
            nw54_ = C9()
            def iife32_():
                coll26_ = _dafny.Set()
                compr_26_: _dafny.Seq
                for compr_26_ in (_dafny.SeqWithoutIsStrInference([d_322_v57_ for d_325_i8_ in range(default__.abs(-88))])).Elements:
                    d_326_v58_: _dafny.Seq = compr_26_
                    if (d_326_v58_) in (_dafny.SeqWithoutIsStrInference([d_322_v57_ for d_325_i8_ in range(default__.abs(-88))])):
                        coll26_ = coll26_.union(_dafny.Set([d_326_v58_]))
                return _dafny.Set(coll26_)
            nw54_.ctor__(d_253_v2_, iife32_()
            , (d_323_v59_)[default__.safeIndex(d_253_v2_, len(d_323_v59_))], d_307_v44_, (d_319_v54_).f13)
            d_324_v60_ = nw54_
            d_327_v61_: _dafny.Set
            d_327_v61_ = _dafny.Set({d_324_v60_})
            d_327_v61_ = d_327_v61_
            index41_ = default__.safeIndex(964, (d_254_v3_).length(0))
            (d_254_v3_)[index41_] = (d_324_v60_).fm9((d_319_v54_).f13, (d_319_v54_).fm8((d_254_v3_)[default__.safeIndex(964, (d_254_v3_).length(0))], d_306_v43_, d_260_v8_, _dafny.SeqWithoutIsStrInference([d_319_v54_.f14, (d_254_v3_)[default__.safeIndex(964, (d_254_v3_).length(0))]]), d_252_globalState_), d_276_v21_, d_252_globalState_)
            d_306_v43_ = (d_319_v54_).f13
        elif True:
            d_328_v63_: _dafny.MultiSet
            d_328_v63_ = _dafny.MultiSet([d_273_v18_])
            d_329_v64_: bool
            d_330_v65_: int
            out6_: bool
            out7_: int
            def iife33_():
                coll27_ = _dafny.Map()
                compr_27_: _dafny.MultiSet
                for compr_27_ in (d_328_v63_).Elements:
                    d_331_v62_: _dafny.MultiSet = compr_27_
                    if (d_331_v62_) in (d_328_v63_):
                        coll27_[d_331_v62_] = d_260_v8_
                return _dafny.Map(coll27_)
            out6_, out7_ = default__.m0(iife33_()
            , d_275_v20_, _dafny.MultiSet(d_269_v14_), (d_275_v20_) or (d_250_v0_), d_252_globalState_)
            d_329_v64_ = out6_
            d_330_v65_ = out7_
            index42_ = default__.safeIndex(964, (d_254_v3_).length(0))
            (d_254_v3_)[index42_] = ((d_273_v18_).cardinality) * (len(d_258_v6_))
            d_332_v66_: D16
            d_332_v66_ = D16_DC33(d_306_v43_, d_253_v2_, d_251_v1_, d_307_v44_, d_300_v40_)
            if ((980) != ((d_254_v3_)[default__.safeIndex(964, (d_254_v3_).length(0))])) and ((d_332_v66_).cf52):
                d_333_v67_: bool
                d_334_v68_: int
                out8_: bool
                out9_: int
                out8_, out9_ = default__.m0(d_274_v19_, d_306_v43_, d_273_v18_, d_275_v20_, d_252_globalState_)
                d_333_v67_ = out8_
                d_334_v68_ = out9_
                d_333_v67_ = d_329_v64_
                d_335_v69_: _dafny.Map
                d_335_v69_ = _dafny.Map({320: d_273_v18_})
                d_336_v70_: _dafny.Seq
                d_336_v70_ = _dafny.SeqWithoutIsStrInference([d_250_v0_, (d_335_v69_) != (d_335_v69_)])
                index43_ = default__.safeIndex(99, (d_270_v15_).length(0))
                (d_270_v15_)[index43_] = (d_336_v70_)[default__.safeIndex(d_253_v2_, len(d_336_v70_))]
                index44_ = default__.safeIndex(99, (d_270_v15_).length(0))
                (d_270_v15_)[index44_] = d_275_v20_
                d_250_v0_ = (d_250_v0_ if (d_250_v0_ if d_329_v64_ else d_329_v64_) else (d_270_v15_)[default__.safeIndex(99, (d_270_v15_).length(0))])
            elif True:
                d_337_v71_: _dafny.Seq
                d_337_v71_ = _dafny.SeqWithoutIsStrInference([d_329_v64_, d_306_v43_])
                d_250_v0_ = (d_337_v71_)[default__.safeIndex((d_254_v3_)[default__.safeIndex(964, (d_254_v3_).length(0))], len(d_337_v71_))]
                d_260_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "becsmvx"))
                d_338_v72_: bool
                d_339_v73_: int
                out10_: bool
                out11_: int
                out10_, out11_ = default__.m0(d_274_v19_, (d_253_v2_) <= (d_253_v2_), d_273_v18_, d_306_v43_, d_252_globalState_)
                d_338_v72_ = out10_
                d_339_v73_ = out11_
                d_340_v74_: D21
                d_340_v74_ = D21_DC44(False)
                pat_let_tv4_ = d_306_v43_
                def iife34_(_pat_let3_0):
                    def iife35_(d_341_dt__update__tmp_h0_):
                        def iife36_(_pat_let4_0):
                            def iife37_(d_342_dt__update_hcf72_h0_):
                                return D21_DC44(d_342_dt__update_hcf72_h0_)
                            return iife37_(_pat_let4_0)
                        return iife36_(pat_let_tv4_)
                    return iife35_(_pat_let3_0)
                d_340_v74_ = iife34_(d_340_v74_)
                d_338_v72_ = (default__.fm3((d_254_v3_)[default__.safeIndex(964, (d_254_v3_).length(0))], d_338_v72_, d_252_globalState_)) not in ((d_273_v18_) | (_dafny.MultiSet([-977, 547, ((d_251_v1_)[d_275_v20_] if (d_275_v20_) in (d_251_v1_) else d_307_v44_)])))
            d_343_v75_: _dafny.MultiSet
            d_343_v75_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_253_v2_, d_330_v65_]), d_269_v14_])
            d_344_v76_: _dafny.MultiSet
            d_344_v76_ = d_343_v75_
            d_345_v77_: _dafny.Seq
            d_345_v77_ = _dafny.SeqWithoutIsStrInference([d_275_v20_])
            d_346_v78_: _dafny.Seq
            d_346_v78_ = _dafny.SeqWithoutIsStrInference([d_345_v77_])
            d_347_v79_: D12
            d_347_v79_ = D12_DC27(d_346_v78_)
            d_348_v80_: _dafny.Map
            d_348_v80_ = _dafny.Map({d_260_v8_: d_347_v79_})
            d_349_v81_: _dafny.Seq
            d_349_v81_ = _dafny.SeqWithoutIsStrInference([d_348_v80_, d_348_v80_])
            d_350_v82_: _dafny.Array
            nw55_ = _dafny.Array(None, 25)
            nw55_[int(0)] = d_344_v76_
            nw55_[int(1)] = d_344_v76_
            nw55_[int(2)] = d_343_v75_
            nw55_[int(3)] = d_344_v76_
            nw55_[int(4)] = d_344_v76_
            nw55_[int(5)] = d_343_v75_
            nw55_[int(6)] = default__.fm72((d_349_v81_)[default__.safeIndex(d_307_v44_, len(d_349_v81_))], d_273_v18_, d_252_globalState_)
            nw55_[int(7)] = d_344_v76_
            nw55_[int(8)] = d_344_v76_
            nw55_[int(9)] = default__.fm73(d_252_globalState_)
            nw55_[int(10)] = d_343_v75_
            nw55_[int(11)] = d_344_v76_
            nw55_[int(12)] = d_344_v76_
            nw55_[int(13)] = d_344_v76_
            nw55_[int(14)] = d_344_v76_
            nw55_[int(15)] = d_344_v76_
            nw55_[int(16)] = d_344_v76_
            nw55_[int(17)] = d_344_v76_
            nw55_[int(18)] = d_344_v76_
            nw55_[int(19)] = d_344_v76_
            nw55_[int(20)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_269_v14_ for d_351_i9_ in range(default__.abs(790))]))
            nw55_[int(21)] = d_344_v76_
            nw55_[int(22)] = d_344_v76_
            nw55_[int(23)] = d_344_v76_
            nw55_[int(24)] = (d_343_v75_).set((d_277_v22_)[default__.safeIndex(766, (d_277_v22_).length(0))], default__.abs((d_254_v3_)[default__.safeIndex(964, (d_254_v3_).length(0))]))
            d_350_v82_ = nw55_
            index45_ = default__.safeIndex(428, (d_350_v82_).length(0))
            (d_350_v82_)[index45_] = d_344_v76_
            index46_ = default__.safeIndex(428, (d_350_v82_).length(0))
            (d_350_v82_)[index46_] = d_344_v76_
            rhs43_ = d_256_v4_
            rhs44_ = not((d_270_v15_)[default__.safeIndex(99, (d_270_v15_).length(0))])
            rhs45_ = default__.safeDivisionInt(len((d_260_v8_) + (d_260_v8_)), d_276_v21_)
            d_256_v4_ = rhs43_
            d_275_v20_ = rhs44_
            d_253_v2_ = rhs45_
        _dafny.print(_dafny.string_of(d_250_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v1_) == (_dafny.MultiSet([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_globalState_).f0) == (_dafny.MultiSet([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_253_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v3_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_256_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_257_v5_) == (_dafny.Map({357604: False, 0: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_258_v6_) == (_dafny.Set({598}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_259_v7_) == (_dafny.Set({_dafny.Set({598})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_260_v8_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_261_v9_) == (_dafny.Map({_dafny.Set({_dafny.Set({598})}): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v10_)[0]) == (_dafny.Set({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v10_)[1]) == (_dafny.Set({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v11_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_267_v12_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_268_v13_) == (_dafny.Map({False: _dafny.CodePoint('j')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v14_) == (_dafny.SeqWithoutIsStrInference([598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v15_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v17_) == (_dafny.SeqWithoutIsStrInference([D0_DC1(), D0_DC1()]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_273_v18_) == (_dafny.MultiSet([598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v19_) == (_dafny.Map({_dafny.MultiSet([598]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smw"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_275_v20_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_276_v21_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[0]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[1]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[2]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[3]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[4]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[5]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[6]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[7]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[8]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[9]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[10]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[11]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[12]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[13]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[14]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[15]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[16]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[17]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[18]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[19]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[20]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[21]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[22]) == (_dafny.SeqWithoutIsStrInference([598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_v22_)[23]) == (_dafny.SeqWithoutIsStrInference([598, 598]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v23_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_299_v36_) == (_dafny.Map({D0_DC1(): 10}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_300_v40_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v41_)[0]) == (_dafny.Map({D0_DC1(): 10}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v41_)[1]) == (_dafny.Map({D0_DC1(): 10}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v41_)[2]) == (_dafny.Map({D0_DC1(): 10}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v41_)[3]) == (_dafny.Map({D0_DC1(): 10}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v41_)[4]) == (_dafny.Map({D0_DC1(): 598}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v41_)[5]) == (_dafny.Map({D0_DC1(): 10}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v41_)[6]) == (_dafny.Map({D0_DC1(): 10}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v41_)[7]) == (_dafny.Map({D0_DC1(): 10}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v41_)[8]) == (_dafny.Map({D0_DC1(): -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v41_)[9]) == (_dafny.Map({D0_DC1(): 10}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v41_)[10]) == (_dafny.Map({D0_DC1(): 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v41_)[11]) == (_dafny.Map({D0_DC1(): 10}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_v42_) == (_dafny.MultiSet([_dafny.CodePoint('j')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_306_v43_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_307_v44_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_308_v45_)))
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

class D0_DC2(D0, NamedTuple('DC2', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf1 == __o.cf1
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
        return lambda: D1_DC4(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6()
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

class D2_DC6(D2, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: D3_DC9(False, int(0), False)
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

class D3_DC9(D3, NamedTuple('DC9', [('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13(int(0), False, int(0), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC13(D4, NamedTuple('DC13', [('cf14', Any), ('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC14(D5, NamedTuple('DC14', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(False, False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)

class D6_DC16(D6, NamedTuple('DC16', [('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)

class D7_DC17(D7, NamedTuple('DC17', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)

class D8_DC18(D8, NamedTuple('DC18', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC20(False, _dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D9_DC19)

class D9_DC20(D9, NamedTuple('DC20', [('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC19(D9, NamedTuple('DC19', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC19({self.cf25.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC19) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, False, _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D10_DC22)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D10_DC21)

class D10_DC22(D10, NamedTuple('DC22', [('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC22({self.cf30.VerbatimString(True)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC22) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC21(D10, NamedTuple('DC21', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC21({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC21) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC24(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D11_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D11_DC26)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D11_DC23)

class D11_DC24(D11, NamedTuple('DC24', [('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC24({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC24) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC25(D11, NamedTuple('DC25', [('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC26(D11, NamedTuple('DC26', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC23(D11, NamedTuple('DC23', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC23({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC23) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC28(int(0), int(0), False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D12_DC27)

class D12_DC28(D12, NamedTuple('DC28', [('cf43', Any), ('cf44', Any), ('cf45', Any), ('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC27(D12, NamedTuple('DC27', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC27({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC27) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D13_DC29)

class D13_DC29(D13, NamedTuple('DC29', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC29({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC29) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D14_DC30)

class D14_DC30(D14, NamedTuple('DC30', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC30({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC30) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D15_DC31)

class D15_DC31(D15, NamedTuple('DC31', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC31({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC31) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC33(False, int(0), _dafny.MultiSet({}), int(0), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D16_DC33)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D16_DC32)

class D16_DC33(D16, NamedTuple('DC33', [('cf52', Any), ('cf53', Any), ('cf54', Any), ('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC33({_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC33) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC32(D16, NamedTuple('DC32', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC32({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC32) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D17_DC34)

class D17_DC34(D17, NamedTuple('DC34', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC34({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC34) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC36(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D18_DC36)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D18_DC35)

class D18_DC36(D18, NamedTuple('DC36', [('cf59', Any), ('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC36({_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC36) and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC35(D18, NamedTuple('DC35', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC35({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC35) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC38(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D19_DC38)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D19_DC37)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D19_DC39)

class D19_DC38(D19, NamedTuple('DC38', [('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC38({_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC38) and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC37(D19, NamedTuple('DC37', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC37({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC37) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC39(D19, NamedTuple('DC39', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC39({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC39) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC41(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D20_DC41)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D20_DC40)

class D20_DC41(D20, NamedTuple('DC41', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC41({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC41) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC40(D20, NamedTuple('DC40', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC40({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC40) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC43(False, False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D21_DC43)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D21_DC44)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D21_DC42)

class D21_DC43(D21, NamedTuple('DC43', [('cf69', Any), ('cf70', Any), ('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC43({_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC43) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC44(D21, NamedTuple('DC44', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC44({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC44) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC42(D21, NamedTuple('DC42', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC42({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC42) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC46(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D22_DC46)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D22_DC47)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D22_DC45)

class D22_DC46(D22, NamedTuple('DC46', [('cf74', Any), ('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC46({_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC46) and self.cf74 == __o.cf74 and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC47(D22, NamedTuple('DC47', [('cf77', Any), ('cf78', Any), ('cf79', Any), ('cf80', Any), ('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC47({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC47) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78 and self.cf79 == __o.cf79 and self.cf80 == __o.cf80 and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC45(D22, NamedTuple('DC45', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC45({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC45) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC49(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D23_DC49)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D23_DC48)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D23_DC50)

class D23_DC49(D23, NamedTuple('DC49', [('cf83', Any), ('cf84', Any), ('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC49({_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC49) and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC48(D23, NamedTuple('DC48', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC48({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC48) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC50(D23, NamedTuple('DC50', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC50({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC50) and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D24_DC51)

class D24_DC51(D24, NamedTuple('DC51', [('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC51({_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC51) and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D25_DC52)

class D25_DC52(D25, NamedTuple('DC52', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC52({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC52) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC54(int(0), int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D26_DC54)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D26_DC53)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D26_DC55)

class D26_DC54(D26, NamedTuple('DC54', [('cf90', Any), ('cf91', Any), ('cf92', Any), ('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC54({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC54) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91 and self.cf92 == __o.cf92 and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC53(D26, NamedTuple('DC53', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC53({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC53) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC55(D26, NamedTuple('DC55', [('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC55({_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC55) and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC57(_dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D27_DC57)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D27_DC56)

class D27_DC57(D27, NamedTuple('DC57', [('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC57({_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC57) and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC56(D27, NamedTuple('DC56', [('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC56({_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC56) and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D28_DC58)

class D28_DC58(D28, NamedTuple('DC58', [('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC58({_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC58) and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, globalState):
        pass

    def m2(self, p0, globalState):
        pass


class T1:
    pass
    @property
    def f5(self):
        return self._f5
    @f5.setter
    def f5(self, value):
        self._f5 = value

class GlobalState:
    def  __init__(self):
        self._f0: _dafny.MultiSet = _dafny.MultiSet({})
        self._f1: int = int(0)
        self._f2: bool = False
        self._f3: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3

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

class C0:
    def  __init__(self):
        self._f12: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f12):
        (self)._f12 = f12

    def fm17(self, p0, globalState):
        return True

    @property
    def f12(self):
        return self._f12

class C1(T1, T0):
    def  __init__(self):
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self._f6: int = int(0)
        self._f4: bool = False
        self.f16: int = int(0)
        self._f15: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f5(self):
        return self._f5
    @f5.setter
    def f5(self, value):
        self._f5 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f4(self):
        return self._f4
    def ctor__(self, f15, f16, f5, f6, f4):
        (self)._f15 = f15
        (self).f16 = f16
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f4 = f4

    def fm8(self, p0, p1, p2, p3, globalState):
        return (self).f4

    def fm9(self, p0, p1, p2, globalState):
        return default__.safeDivisionInt((self).f15, (0) - ((self).f6))

    def fm6(self, p0, p1, p2, p3, globalState):
        if ((self).f4) or (True):
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f6 for d_352_i0_ in range(default__.abs(43))]))).intersection(_dafny.MultiSet([(self).f6, 469]))
        elif True:
            return (_dafny.MultiSet([self.f16])) | (_dafny.MultiSet([(self).f6]))

    def fm7(self, globalState):
        source9_ = D1_DC3(not((self).f4))
        if source9_.is_DC4:
            d_353___mcc_h0_ = source9_.cf3
            d_354___mcc_h1_ = source9_.cf4
            d_355_cf4_ = d_354___mcc_h1_
            d_356_cf3_ = d_353___mcc_h0_
            return len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iyylxfgh")))
        elif True:
            d_357___mcc_h2_ = source9_.cf2
            d_358_cf2_ = d_357___mcc_h2_
            return len((_dafny.SeqWithoutIsStrInference([d_358_cf2_, not(d_358_cf2_)]) if True else _dafny.SeqWithoutIsStrInference([d_358_cf2_, d_358_cf2_])))

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_359_v0_: _dafny.Set
        d_359_v0_ = _dafny.Set({(self).f15})
        d_360_v1_: _dafny.MultiSet
        d_360_v1_ = _dafny.MultiSet([self.f16, len(d_359_v0_), self.f16])
        d_361_v2_: D3
        d_361_v2_ = D3_DC9((self).f4, (self).f6, not((self).f4))
        def iife38_(_pat_let5_0):
            def iife39_(d_362_dt__update__tmp_h0_):
                def iife40_(_pat_let6_0):
                    def iife41_(d_363_dt__update_hcf8_h0_):
                        return D3_DC9(d_363_dt__update_hcf8_h0_, (d_362_dt__update__tmp_h0_).cf9, (d_362_dt__update__tmp_h0_).cf10)
                    return iife41_(_pat_let6_0)
                return iife40_(False)
            return iife39_(_pat_let5_0)
        source10_ = (d_361_v2_ if (d_360_v1_).ispropersubset(_dafny.MultiSet([(self).f6])) else iife38_(d_361_v2_))
        if source10_.is_DC9:
            d_364___mcc_h0_ = source10_.cf8
            d_365___mcc_h1_ = source10_.cf9
            d_366___mcc_h2_ = source10_.cf10
            d_367_cf10_ = d_366___mcc_h2_
            d_368_cf9_ = d_365___mcc_h1_
            d_369_cf8_ = d_364___mcc_h0_
            d_370_v3_: str
            d_370_v3_ = _dafny.CodePoint('k')
            if (-828) <= ((len(default__.fm37(d_369_cf8_, not(d_367_cf10_), p0, default__.fm0(d_370_v3_, globalState), globalState))) - (d_368_cf9_)):
                d_371_v4_: _dafny.Array
                nw56_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_371_v4_ = nw56_
                d_372_v5_: _dafny.Array
                def lambda7_(d_373_p0_):
                    def lambda8_(d_374_i0_):
                        return (d_374_i0_) * (d_373_p0_)

                    return lambda8_

                init4_ = lambda7_(p0)
                nw57_ = _dafny.Array(None, 22)
                for i0_4_ in range(nw57_.length(0)):
                    nw57_[i0_4_] = init4_(i0_4_)
                d_372_v5_ = nw57_
                index47_ = default__.safeIndex(500, (d_371_v4_).length(0))
                (d_371_v4_)[index47_] = d_372_v5_
                index48_ = default__.safeIndex(500, (d_371_v4_).length(0))
                (d_371_v4_)[index48_] = d_372_v5_
                d_375_v6_: _dafny.Seq
                d_375_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqnumrpc"))
                d_375_v6_ = (_dafny.SeqWithoutIsStrInference([(d_375_v6_)[default__.safeIndex(len((D6_DC16((self).f4, (self).f4, _dafny.SeqWithoutIsStrInference([len(_dafny.Set({732, len(_dafny.Map({d_368_cf9_: p0}))})) for d_376_i2_ in range(default__.abs(566))]))).cf22), len(d_375_v6_))] for d_377_i1_ in range(default__.abs(435))])).set(default__.safeIndex(834, len(_dafny.SeqWithoutIsStrInference([(d_375_v6_)[default__.safeIndex(len((D6_DC16((self).f4, (self).f4, _dafny.SeqWithoutIsStrInference([len(_dafny.Set({732, len(_dafny.Map({d_368_cf9_: p0}))})) for d_378_i2_ in range(default__.abs(566))]))).cf22), len(d_375_v6_))] for d_379_i1_ in range(default__.abs(435))]))), d_370_v3_)
                d_380_v7_: _dafny.Seq
                d_380_v7_ = _dafny.SeqWithoutIsStrInference([d_368_cf9_])
                d_381_v8_: _dafny.Map
                d_381_v8_ = _dafny.Map({(self).fm8((self).fm7(globalState), (self).f4, d_375_v6_, d_380_v7_, globalState): len(d_380_v7_)})
                r0 = d_381_v8_
                d_369_cf8_ = d_367_cf10_
                d_382_v9_: _dafny.Seq
                d_382_v9_ = _dafny.SeqWithoutIsStrInference([(self).f4])
                d_382_v9_ = (_dafny.SeqWithoutIsStrInference([d_369_cf8_])) + (d_382_v9_)
            elif True:
                d_383_v11_: _dafny.Seq
                d_383_v11_ = _dafny.SeqWithoutIsStrInference([(self).f4, d_369_cf8_])
                d_384_v12_: _dafny.Map
                def iife42_():
                    coll28_ = _dafny.Map()
                    compr_28_: int
                    for compr_28_ in (_dafny.MultiSet([(self).f15])).Elements:
                        d_385_v10_: int = compr_28_
                        if (d_385_v10_) in (_dafny.MultiSet([(self).f15])):
                            coll28_[default__.safeModuloInt(d_385_v10_, p0)] = (self).f4
                    return _dafny.Map(coll28_)
                d_384_v12_ = _dafny.Map({iife42_()
                : default__.fm3(d_368_cf9_, (d_383_v11_)[default__.safeIndex(d_368_cf9_, len(d_383_v11_))], globalState)})
                d_384_v12_ = (d_384_v12_) | (d_384_v12_)
                d_386_v13_: C0
                nw58_ = C0()
                nw58_.ctor__(d_368_cf9_)
                d_386_v13_ = nw58_
                (self).f16 = default__.safeDivisionInt((self).f15, (d_386_v13_).f12)
                d_387_v14_: _dafny.Seq
                d_387_v14_ = _dafny.SeqWithoutIsStrInference([(0) - (p0), 886])
                (self).f16 = ((d_387_v14_)[default__.safeIndex(d_368_cf9_, len(d_387_v14_))]) - (760)
                d_388_v15_: _dafny.MultiSet
                d_388_v15_ = _dafny.MultiSet([not(not(d_369_cf8_))])
                d_389_v16_: _dafny.Seq
                d_389_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tplcf"))
                (self).f16 = (d_368_cf9_) - (len((default__.fm37(d_367_cf10_, (self).f4, ((d_388_v15_)[(self).f4] if ((self).f4) in (d_388_v15_) else 150), (self).f4, globalState)) + ((d_389_v16_).set(default__.safeIndex(self.f16, len(d_389_v16_)), d_370_v3_))))
            d_369_cf8_ = d_369_cf8_
            (self).f5 = self.f5
            if d_367_cf10_:
                d_369_cf8_ = not(d_369_cf8_)
                (self).f16 = default__.safeModuloInt(p0, (self).f6)
                d_368_cf9_ = (d_368_cf9_) - (905)
                d_390_v17_: C0
                nw59_ = C0()
                nw59_.ctor__((d_368_cf9_) - (d_368_cf9_))
                d_390_v17_ = nw59_
                nw60_ = C0()
                nw60_.ctor__(921)
                d_390_v17_ = nw60_
                d_391_v18_: _dafny.Seq
                d_391_v18_ = _dafny.SeqWithoutIsStrInference([(self).f4, True, not(d_367_cf10_)])
                d_392_v19_: _dafny.Seq
                d_392_v19_ = _dafny.SeqWithoutIsStrInference([((self).f6) * (len((D12_DC27(_dafny.SeqWithoutIsStrInference([d_391_v18_]))).cf42))])
                d_392_v19_ = d_392_v19_
            elif True:
                (self).f16 = (0) - (d_368_cf9_)
                d_393_v20_: _dafny.Array
                def lambda9_(d_394_cf9_):
                    def lambda10_(d_395_i3_):
                        return _dafny.SeqWithoutIsStrInference([d_394_cf9_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lyl")))])

                    return lambda10_

                init5_ = lambda9_(d_368_cf9_)
                nw61_ = _dafny.Array(None, 8)
                for i0_5_ in range(nw61_.length(0)):
                    nw61_[i0_5_] = init5_(i0_5_)
                d_393_v20_ = nw61_
                d_396_v21_: D6
                d_396_v21_ = D6_DC16(d_367_cf10_, d_367_cf10_, _dafny.SeqWithoutIsStrInference([p0]))
                index49_ = default__.safeIndex(686, (d_393_v20_).length(0))
                (d_393_v20_)[index49_] = (d_396_v21_).cf22
                d_397_v22_: _dafny.Seq
                d_397_v22_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                index50_ = default__.safeIndex(686, (d_393_v20_).length(0))
                (d_393_v20_)[index50_] = d_397_v22_
                d_398_v23_: _dafny.Map
                d_398_v23_ = _dafny.Map({(self).f4: self.f16})
                d_399_v24_: _dafny.Seq
                d_399_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlcjkdars"))
                d_367_cf10_ = (d_367_cf10_ if (self).fm8(((d_398_v23_)[d_369_cf8_] if (d_369_cf8_) in (d_398_v23_) else len(d_399_v24_)), True, d_399_v24_, d_397_v22_, globalState) else not((self).f4))
                d_400_v25_: _dafny.Array
                nw62_ = _dafny.Array(int(0), 18)
                d_400_v25_ = nw62_
                d_400_v25_ = d_400_v25_
                d_370_v3_ = (d_370_v3_ if (self).f4 else d_370_v3_)
        elif source10_.is_DC10:
            d_401___mcc_h3_ = source10_.cf11
            d_402_cf11_ = d_401___mcc_h3_
            if (self).f4:
                d_403_v26_: C0
                nw63_ = C0()
                nw63_.ctor__((self).fm7(globalState))
                d_403_v26_ = nw63_
                d_404_v27_: _dafny.Seq
                d_404_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfgsjgori"))
                d_405_v28_: D9
                d_405_v28_ = D9_DC19(d_404_v27_)
                pat_let_tv5_ = d_404_v27_
                d_406_v29_: _dafny.MultiSet
                def iife43_(_pat_let7_0):
                    def iife44_(d_407_dt__update__tmp_h1_):
                        def iife45_(_pat_let8_0):
                            def iife46_(d_408_dt__update_hcf25_h0_):
                                return D9_DC19(d_408_dt__update_hcf25_h0_)
                            return iife46_(_pat_let8_0)
                        return iife45_(pat_let_tv5_)
                    return iife44_(_pat_let7_0)
                d_406_v29_ = _dafny.MultiSet([iife43_(d_405_v28_), d_405_v28_])
                d_409_v30_: _dafny.Set
                d_409_v30_ = _dafny.Set({True, (self).f4})
                d_410_v31_: _dafny.Seq
                d_410_v31_ = _dafny.SeqWithoutIsStrInference([len(d_409_v30_)])
                d_411_v32_: _dafny.Map
                d_411_v32_ = _dafny.Map({d_406_v29_: d_410_v31_})
                d_412_v34_: _dafny.Set
                d_412_v34_ = _dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_405_v28_])), _dafny.MultiSet([d_405_v28_]), d_406_v29_, d_406_v29_})
                d_413_v35_: _dafny.Seq
                d_413_v35_ = _dafny.SeqWithoutIsStrInference([(self).f4])
                d_414_v36_: _dafny.Array
                nw64_ = _dafny.Array(False, 11)
                d_414_v36_ = nw64_
                d_415_v37_: _dafny.Map
                d_415_v37_ = _dafny.Map({d_414_v36_: d_402_cf11_})
                d_416_v38_: _dafny.Array
                nw65_ = _dafny.Array(None, 14)
                nw65_[int(0)] = self.f16
                def iife47_():
                    coll29_ = _dafny.Map()
                    compr_29_: _dafny.MultiSet
                    for compr_29_ in (d_412_v34_).Elements:
                        d_417_v33_: _dafny.MultiSet = compr_29_
                        if (d_417_v33_) in (d_412_v34_):
                            coll29_[d_417_v33_] = _dafny.SeqWithoutIsStrInference([(d_403_v26_).f12, -615])
                    return _dafny.Map(coll29_)
                nw65_[int(1)] = len((d_411_v32_) | (iife47_()
                ))
                nw65_[int(2)] = p0
                nw65_[int(3)] = d_402_cf11_
                nw65_[int(4)] = (self).f6
                nw65_[int(5)] = default__.safeDivisionInt(len(d_413_v35_), p0)
                nw65_[int(6)] = (self).f6
                nw65_[int(7)] = ((self).f6 if (self).fm8((self).f6, (d_413_v35_)[default__.safeIndex(default__.fm3(130, False, globalState), len(d_413_v35_))], d_404_v27_, d_410_v31_, globalState) else len(d_410_v31_))
                nw65_[int(8)] = self.f16
                nw65_[int(9)] = p0
                nw65_[int(10)] = self.f16
                nw65_[int(11)] = (0) - (((d_415_v37_)[d_414_v36_] if (d_414_v36_) in (d_415_v37_) else len(d_413_v35_)))
                nw65_[int(12)] = (d_403_v26_).f12
                nw65_[int(13)] = default__.safeDivisionInt(d_402_cf11_, p0)
                d_416_v38_ = nw65_
                index51_ = default__.safeIndex(159, (d_416_v38_).length(0))
                (d_416_v38_)[index51_] = default__.safeDivisionInt(-342, default__.fm3((self).f6, (self).f4, globalState))
                index52_ = default__.safeIndex(159, (d_416_v38_).length(0))
                (d_416_v38_)[index52_] = default__.fm3(((self).f6) - ((d_403_v26_).f12), (self).f4, globalState)
                d_418_v39_: _dafny.Seq
                d_418_v39_ = _dafny.SeqWithoutIsStrInference([d_409_v30_])
                d_419_v40_: _dafny.Map
                d_419_v40_ = _dafny.Map({(d_413_v35_) + (d_413_v35_): d_418_v39_})
                d_420_v41_: _dafny.Map
                d_420_v41_ = _dafny.Map({(0) - ((self).f15): d_409_v30_})
                d_419_v40_ = (d_419_v40_).set(d_413_v35_, _dafny.SeqWithoutIsStrInference([((d_420_v41_)[372] if (372) in (d_420_v41_) else d_409_v30_)]))
                d_421_v42_: bool
                d_421_v42_ = True
                d_422_v43_: _dafny.MultiSet
                d_422_v43_ = _dafny.MultiSet([d_421_v42_, d_421_v42_, (d_409_v30_).ispropersubset(_dafny.Set({(self).f4, (self).f4, d_421_v42_, (self).fm8((d_416_v38_)[default__.safeIndex(159, (d_416_v38_).length(0))], (self).f4, d_404_v27_, d_410_v31_, globalState), d_421_v42_}))])
                def iife48_():
                    coll30_ = _dafny.Set()
                    compr_30_: int
                    for compr_30_ in _dafny.IntegerRange(681, 893):
                        d_423_v44_: int = compr_30_
                        if ((681) <= (d_423_v44_)) and ((d_423_v44_) < (893)):
                            coll30_ = coll30_.union(_dafny.Set([(d_423_v44_) * ((d_403_v26_).f12)]))
                    return _dafny.Set(coll30_)
                rhs46_ = (d_403_v26_).f12
                rhs47_ = False
                rhs48_ = iife48_()

                rhs49_ = (_dafny.MultiSet([(self).f4, False])).set(not ((self).f4) or (False), default__.abs((len(d_404_v27_)) + ((d_403_v26_).f12)))
                lhs28_ = self
                lhs28_.f16 = rhs46_
                d_421_v42_ = rhs47_
                d_359_v0_ = rhs48_
                d_422_v43_ = rhs49_
                d_424_v45_: _dafny.Set
                d_424_v45_ = d_359_v0_
                d_425_v46_: _dafny.Set
                d_425_v46_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_426_i4_ in range(default__.abs(710))])})
                d_427_v47_: _dafny.Map
                d_427_v47_ = _dafny.Map({d_424_v45_: (d_425_v46_).isdisjoint(d_425_v46_)})
                d_427_v47_ = (d_427_v47_).set(d_424_v45_, (self).f4)
            elif True:
                d_428_v48_: _dafny.Seq
                d_428_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrngovhln"))
                d_428_v48_ = (d_428_v48_) + (d_428_v48_)
                d_429_v49_: _dafny.MultiSet
                d_429_v49_ = _dafny.MultiSet([d_428_v48_, d_428_v48_, d_428_v48_])
                d_402_cf11_ = default__.safeModuloInt(default__.fm3((self).f6, (self).f4, globalState), ((d_429_v49_)[d_428_v48_] if (d_428_v48_) in (d_429_v49_) else (self).f15))
                d_430_v50_: _dafny.Array
                nw66_ = _dafny.Array(int(0), 9)
                d_430_v50_ = nw66_
                index53_ = default__.safeIndex(442, (d_430_v50_).length(0))
                (d_430_v50_)[index53_] = (0) - (len(d_359_v0_))
                d_431_v51_: _dafny.Set
                d_431_v51_ = _dafny.Set({(self).f4})
                index54_ = default__.safeIndex(442, (d_430_v50_).length(0))
                (d_430_v50_)[index54_] = default__.safeModuloInt((self).fm7(globalState), (0) - (len(_dafny.Map({p0: d_431_v51_}))))
                d_402_cf11_ = (d_430_v50_)[default__.safeIndex(442, (d_430_v50_).length(0))]
                d_432_v52_: _dafny.Array
                def lambda11_(d_433_i5_):
                    return (_dafny.MultiSet([not((self).f4)])).intersection(_dafny.MultiSet([(self).f4, (self).f4]))

                init6_ = lambda11_
                nw67_ = _dafny.Array(None, 9)
                for i0_6_ in range(nw67_.length(0)):
                    nw67_[i0_6_] = init6_(i0_6_)
                d_432_v52_ = nw67_
                d_434_v53_: _dafny.MultiSet
                d_434_v53_ = _dafny.MultiSet([(self).f4])
                d_435_v54_: _dafny.Map
                d_435_v54_ = _dafny.Map({(self).f4: d_434_v53_})
                index55_ = default__.safeIndex(258, (d_432_v52_).length(0))
                (d_432_v52_)[index55_] = (_dafny.MultiSet([(self).f4]) if not((self).f4) else ((d_435_v54_)[(self).f4] if ((self).f4) in (d_435_v54_) else d_434_v53_))
                index56_ = default__.safeIndex(258, (d_432_v52_).length(0))
                (d_432_v52_)[index56_] = (d_434_v53_) - ((_dafny.MultiSet([(self).f4])) - (d_434_v53_))
            d_436_v55_: _dafny.Seq
            d_436_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dohmu"))
            d_437_v56_: _dafny.Map
            d_437_v56_ = _dafny.Map({d_436_v55_: (self).f6})
            d_438_v57_: _dafny.Seq
            d_438_v57_ = _dafny.SeqWithoutIsStrInference([d_437_v56_])
            d_439_v58_: _dafny.Array
            nw68_ = _dafny.Array(False, 26)
            d_439_v58_ = nw68_
            d_440_v59_: int
            out12_: int
            out12_ = (self).m11((self).f4, not((self).f4), ((d_438_v57_)[default__.safeIndex((self).f6, len(d_438_v57_))]).set(d_436_v55_, d_402_cf11_), d_439_v58_, globalState)
            d_440_v59_ = out12_
            d_441_v60_: _dafny.Array
            def lambda12_(d_442_v55_):
                def lambda13_(d_443_i6_):
                    return (d_442_v55_) + (d_442_v55_)

                return lambda13_

            init7_ = lambda12_(d_436_v55_)
            nw69_ = _dafny.Array(None, 25)
            for i0_7_ in range(nw69_.length(0)):
                nw69_[i0_7_] = init7_(i0_7_)
            d_441_v60_ = nw69_
            index57_ = default__.safeIndex(910, (d_441_v60_).length(0))
            (d_441_v60_)[index57_] = d_436_v55_
            d_444_v61_: _dafny.Map
            d_444_v61_ = _dafny.Map({(self).f15: d_402_cf11_})
            d_445_v62_: _dafny.Map
            d_445_v62_ = _dafny.Map({-87: d_444_v61_})
            index58_ = default__.safeIndex(910, (d_441_v60_).length(0))
            (d_441_v60_)[index58_] = default__.fm37((self).f4, (self).f4, p0, ((self).f15) not in (((d_445_v62_)[p0] if (p0) in (d_445_v62_) else (d_444_v61_).set((self).fm9((self).f4, True, d_402_cf11_, globalState), (self).f15))), globalState)
            d_446_v63_: _dafny.Array
            def lambda14_(d_447_i7_):
                return (_dafny.MultiSet([D4_DC12(_dafny.CodePoint('f'))])).intersection(_dafny.MultiSet([D4_DC12(_dafny.CodePoint('l'))]))

            init8_ = lambda14_
            nw70_ = _dafny.Array(None, 21)
            for i0_8_ in range(nw70_.length(0)):
                nw70_[i0_8_] = init8_(i0_8_)
            d_446_v63_ = nw70_
            d_446_v63_ = d_446_v63_
        elif source10_.is_DC8:
            d_448___mcc_h4_ = source10_.cf7
            d_449_cf7_ = d_448___mcc_h4_
            d_450_v64_: _dafny.Array
            nw71_ = _dafny.Array(False, 28)
            d_450_v64_ = nw71_
            index59_ = default__.safeIndex(707, (d_450_v64_).length(0))
            (d_450_v64_)[index59_] = (self.f16) > (548)
            index60_ = default__.safeIndex(707, (d_450_v64_).length(0))
            (d_450_v64_)[index60_] = (self).f4
            d_451_v65_: _dafny.Seq
            d_451_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dv"))
            d_451_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qssuchryd"))
            d_452_v66_: C0
            nw72_ = C0()
            nw72_.ctor__(p0)
            d_452_v66_ = nw72_
            d_453_v67_: _dafny.Array
            nw73_ = _dafny.Array(int(0), 14)
            d_453_v67_ = nw73_
            d_454_v68_: _dafny.Set
            d_454_v68_ = _dafny.Set({d_453_v67_, d_453_v67_})
            d_455_v69_: _dafny.Set
            d_455_v69_ = d_454_v68_
            d_456_v70_: _dafny.Map
            d_456_v70_ = _dafny.Map({(self).f4: (d_455_v69_)})
            d_457_v71_: _dafny.Seq
            d_457_v71_ = _dafny.SeqWithoutIsStrInference([d_454_v68_])
            d_456_v70_ = (d_456_v70_).set(not((self).f4), (d_454_v68_ if (self).f4 else (d_457_v71_)[default__.safeIndex((self).f6, len(d_457_v71_))]))
        elif True:
            d_458___mcc_h5_ = source10_.cf12
            d_459_cf12_ = d_458___mcc_h5_
            source11_ = D11_DC24(not((self).f4), (self).f4, default__.safeModuloInt(self.f16, -625))
            if source11_.is_DC24:
                d_460___mcc_h6_ = source11_.cf36
                d_461___mcc_h7_ = source11_.cf37
                d_462___mcc_h8_ = source11_.cf38
                d_463_cf38_ = d_462___mcc_h8_
                d_464_cf37_ = d_461___mcc_h7_
                d_465_cf36_ = d_460___mcc_h6_
                (self).f16 = d_463_cf38_
                d_465_cf36_ = d_465_cf36_
                d_466_v72_: _dafny.Array
                nw74_ = _dafny.Array(False, 11)
                d_466_v72_ = nw74_
                d_467_v73_: D2
                d_467_v73_ = D2_DC5(d_466_v72_)
                d_467_v73_ = d_467_v73_
                d_468_v74_: _dafny.Array
                nw75_ = _dafny.Array(None, 2)
                nw75_[int(0)] = 445
                nw75_[int(1)] = (self).f15
                d_468_v74_ = nw75_
                d_469_v75_: _dafny.Seq
                d_469_v75_ = _dafny.SeqWithoutIsStrInference([d_468_v74_])
                d_469_v75_ = ((_dafny.SeqWithoutIsStrInference([d_468_v74_])) + ((d_469_v75_).set(default__.safeIndex((self).f6, len(d_469_v75_)), d_468_v74_))) + ((d_469_v75_) + (d_469_v75_))
            elif source11_.is_DC25:
                d_470___mcc_h9_ = source11_.cf39
                d_471___mcc_h10_ = source11_.cf40
                d_472_cf40_ = d_471___mcc_h10_
                d_473_cf39_ = d_470___mcc_h9_
                d_474_v76_: str
                d_474_v76_ = _dafny.CodePoint('b')
                d_475_v77_: _dafny.Map
                d_475_v77_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odwhwaejb")): (_dafny.MultiSet([d_473_cf39_])).cardinality})
                d_476_v79_: _dafny.Seq
                d_476_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqtrskur"))
                d_477_v80_: _dafny.Set
                d_477_v80_ = _dafny.Set({d_476_v79_})
                d_478_v81_: _dafny.Array
                def lambda15_(d_479_i9_):
                    return not((self).f4)

                init9_ = lambda15_
                nw76_ = _dafny.Array(None, 20)
                for i0_9_ in range(nw76_.length(0)):
                    nw76_[i0_9_] = init9_(i0_9_)
                d_478_v81_ = nw76_
                d_480_v82_: int
                out13_: int
                def iife49_():
                    coll31_ = _dafny.Map()
                    compr_31_: _dafny.Seq
                    for compr_31_ in (d_477_v80_).Elements:
                        d_482_v78_: _dafny.Seq = compr_31_
                        if (d_482_v78_) in (d_477_v80_):
                            coll31_[d_482_v78_] = len(_dafny.Set({(self).f4, d_473_cf39_, d_473_cf39_}))
                    return _dafny.Map(coll31_)
                out13_ = (self).m11((default__.fm37(d_473_cf39_, (self).f4, (self).f6, False, globalState)) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_481_i8_ in range(default__.abs(906))])), default__.fm0(d_474_v76_, globalState), (d_475_v77_) | (iife49_()
                ), d_478_v81_, globalState)
                d_480_v82_ = out13_
                d_476_v79_ = _dafny.SeqWithoutIsStrInference([d_474_v76_ for d_483_i10_ in range(default__.abs(812))])
                d_484_v83_: _dafny.Seq
                d_484_v83_ = _dafny.SeqWithoutIsStrInference([self.f16])
                d_485_v84_: _dafny.Map
                d_485_v84_ = _dafny.Map({_dafny.MultiSet(d_484_v83_): d_476_v79_})
                d_486_v85_: bool
                d_487_v86_: int
                out14_: bool
                out15_: int
                out14_, out15_ = default__.m0(d_485_v84_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcxrsang"))) <= (d_476_v79_), (d_360_v1_).intersection(d_360_v1_), (d_474_v76_) not in ((d_476_v79_).set(default__.safeIndex(p0, len(d_476_v79_)), d_474_v76_)), globalState)
                d_486_v85_ = out14_
                d_487_v86_ = out15_
                d_488_v87_: _dafny.Seq
                d_488_v87_ = _dafny.SeqWithoutIsStrInference([d_473_cf39_])
                d_489_v88_: _dafny.Map
                d_489_v88_ = _dafny.Map({d_487_v86_: d_474_v76_})
                d_490_v89_: _dafny.Map
                d_490_v89_ = _dafny.Map({d_487_v86_: d_472_cf40_})
                d_491_v90_: _dafny.Map
                d_491_v90_ = _dafny.Map({len(d_490_v89_): d_472_cf40_})
                d_492_v91_: _dafny.Seq
                d_492_v91_ = _dafny.SeqWithoutIsStrInference([default__.fm40((self).f4, d_486_v85_, d_487_v86_, globalState), d_488_v87_, d_488_v87_])
                d_486_v85_ = not(((d_360_v1_).intersection(_dafny.MultiSet(default__.fm39(d_476_v79_, D12_DC27(d_492_v91_), d_473_cf39_, globalState)))).ispropersubset((d_360_v1_) - ((self).fm6(default__.fm38((self).f15, p0, d_486_v85_, len(d_488_v87_), globalState), D0_DC0((self).f15), d_489_v88_, len(d_491_v90_), globalState))))
            elif source11_.is_DC26:
                d_493___mcc_h11_ = source11_.cf41
                d_494_cf41_ = d_493___mcc_h11_
                (self).f16 = default__.safeDivisionInt((self).f6, (self).f15)
                d_495_v92_: _dafny.Array
                nw77_ = _dafny.Array(False, 18)
                d_495_v92_ = nw77_
                index61_ = default__.safeIndex(725, (d_495_v92_).length(0))
                (d_495_v92_)[index61_] = (self.f16) <= (p0)
                d_496_v93_: _dafny.Seq
                d_496_v93_ = _dafny.SeqWithoutIsStrInference([106, (self).f15])
                index62_ = default__.safeIndex(725, (d_495_v92_).length(0))
                (d_495_v92_)[index62_] = (d_496_v93_) == (d_496_v93_)
                index63_ = default__.safeIndex(725, (d_495_v92_).length(0))
                (d_495_v92_)[index63_] = ((d_495_v92_)[default__.safeIndex(725, (d_495_v92_).length(0))]) == ((d_495_v92_)[default__.safeIndex(725, (d_495_v92_).length(0))])
                d_497_v94_: D0
                d_497_v94_ = D0_DC0((self).f15)
                d_498_v95_: _dafny.Seq
                d_498_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                d_499_v96_: _dafny.Map
                d_499_v96_ = _dafny.Map({p0: d_498_v95_})
                d_500_v97_: _dafny.Seq
                d_500_v97_ = _dafny.SeqWithoutIsStrInference([d_498_v95_, ((d_499_v96_)[-706] if (-706) in (d_499_v96_) else d_498_v95_)])
                d_501_v98_: str
                d_501_v98_ = _dafny.CodePoint('y')
                pat_let_tv6_ = p0
                pat_let_tv7_ = p0
                index64_ = default__.safeIndex(725, (d_495_v92_).length(0))
                def iife50_(_pat_let9_0):
                    def iife51_(d_502_dt__update__tmp_h2_):
                        def iife52_(_pat_let10_0):
                            def iife53_(d_503_dt__update_hcf0_h0_):
                                return D0_DC0(d_503_dt__update_hcf0_h0_)
                            return iife53_(_pat_let10_0)
                        return iife52_((pat_let_tv6_) + (pat_let_tv7_))
                    return iife51_(_pat_let9_0)
                rhs50_ = iife50_(d_497_v94_)
                rhs51_ = (d_500_v97_) + (d_500_v97_)
                rhs52_ = (self).f4
                rhs53_ = d_501_v98_
                lhs29_ = d_495_v92_
                lhs30_ = default__.safeIndex(725, (d_495_v92_).length(0))
                d_497_v94_ = rhs50_
                d_500_v97_ = rhs51_
                lhs29_[lhs30_] = rhs52_
                d_501_v98_ = rhs53_
            elif True:
                d_504___mcc_h12_ = source11_.cf35
                d_505_cf35_ = d_504___mcc_h12_
                d_506_v99_: _dafny.Array
                nw78_ = _dafny.Array(_dafny.Map({}), 12)
                d_506_v99_ = nw78_
                d_507_v100_: _dafny.Seq
                d_507_v100_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfnegvsi"))
                d_508_v101_: _dafny.Array
                def lambda16_(d_509_i11_):
                    return (d_509_i11_) - (self.f16)

                init10_ = lambda16_
                nw79_ = _dafny.Array(None, 19)
                for i0_10_ in range(nw79_.length(0)):
                    nw79_[i0_10_] = init10_(i0_10_)
                d_508_v101_ = nw79_
                d_510_v102_: _dafny.Map
                d_510_v102_ = _dafny.Map({d_507_v100_: d_508_v101_})
                index65_ = default__.safeIndex(305, (d_506_v99_).length(0))
                (d_506_v99_)[index65_] = (d_510_v102_).set(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_511_i12_ in range(default__.abs(252))]), d_508_v101_)
                index66_ = default__.safeIndex(305, (d_506_v99_).length(0))
                (d_506_v99_)[index66_] = d_510_v102_
                d_512_v103_: _dafny.Seq
                d_512_v103_ = _dafny.SeqWithoutIsStrInference([self.f16, 357])
                d_507_v100_ = default__.fm37((self).f4, (p0) != (len(d_512_v103_)), (0) - ((self).f6), (self).f4, globalState)
                d_513_v104_: _dafny.Array
                nw80_ = _dafny.Array(False, 28)
                d_513_v104_ = nw80_
                index67_ = default__.safeIndex(205, (d_513_v104_).length(0))
                (d_513_v104_)[index67_] = (self).f4
                index68_ = default__.safeIndex(205, (d_513_v104_).length(0))
                (d_513_v104_)[index68_] = ((self).f4) and ((self).f4)
                d_514_v105_: _dafny.Map
                d_514_v105_ = _dafny.Map({default__.fm3((self).f15, (self).f4, globalState): (d_513_v104_)[default__.safeIndex(205, (d_513_v104_).length(0))]})
                d_515_v106_: _dafny.Map
                d_515_v106_ = _dafny.Map({d_513_v104_: (self).f15})
                d_516_v107_: str
                d_516_v107_ = _dafny.CodePoint('f')
                d_517_v108_: _dafny.Map
                d_517_v108_ = _dafny.Map({d_516_v107_: (self).f4})
                d_518_v109_: D10
                d_518_v109_ = D10_DC22(d_507_v100_, (self).f4, ((d_517_v108_)[d_516_v107_] if (d_516_v107_) in (d_517_v108_) else not((self).f4)), d_513_v104_, p0)
                d_514_v105_ = (d_514_v105_).set(((d_515_v106_)[d_513_v104_] if (d_513_v104_) in (d_515_v106_) else ((d_360_v1_)[self.f16] if (self.f16) in (d_360_v1_) else (self).f15)), (d_518_v109_).cf32)
            d_519_v110_: _dafny.Array
            nw81_ = _dafny.Array(int(0), 9)
            d_519_v110_ = nw81_
            index69_ = default__.safeIndex(1, (d_519_v110_).length(0))
            (d_519_v110_)[index69_] = default__.safeModuloInt(739, (self).f15)
            index70_ = default__.safeIndex(1, (d_519_v110_).length(0))
            (d_519_v110_)[index70_] = (self).f6
            (self).f16 = ((self).f15) + ((0) - (p0))
            d_520_v111_: C0
            nw82_ = C0()
            nw82_.ctor__((self).f15)
            d_520_v111_ = nw82_
        d_521_v112_: _dafny.Seq
        d_521_v112_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmhajedk"))
        d_522_v113_: _dafny.Set
        d_522_v113_ = _dafny.Set({d_521_v112_})
        d_523_i13_: int
        d_523_i13_ = 0
        with _dafny.label("0"):
            while not((d_522_v113_).ispropersubset((d_522_v113_).intersection(d_522_v113_))):
                with _dafny.c_label("0"):
                    if (d_523_i13_) >= (100):
                        raise _dafny.Break("0")
                    d_523_i13_ = (d_523_i13_) + (1)
                    (self).f16 = len((d_521_v112_) + (d_521_v112_))
                    d_524_v114_: C0
                    nw83_ = C0()
                    nw83_.ctor__(self.f16)
                    d_524_v114_ = nw83_
                    d_525_v115_: str
                    d_525_v115_ = _dafny.CodePoint('n')
                    d_526_v116_: _dafny.Map
                    d_526_v116_ = _dafny.Map({(self).f4: d_525_v115_})
                    d_527_v117_: _dafny.Seq
                    d_527_v117_ = _dafny.SeqWithoutIsStrInference([((self).f15) + (len(d_526_v116_)), (d_524_v114_).f12, (self).f15])
                    d_528_v118_: _dafny.Array
                    nw84_ = _dafny.Array(int(0), 22)
                    d_528_v118_ = nw84_
                    index71_ = default__.safeIndex(23, (d_528_v118_).length(0))
                    (d_528_v118_)[index71_] = (self).f15
                    d_529_v119_: _dafny.Set
                    d_529_v119_ = _dafny.Set({(self).f4, (self).f4})
                    d_530_v120_: _dafny.Map
                    d_530_v120_ = _dafny.Map({self.f16: d_529_v119_})
                    d_531_v121_: _dafny.Map
                    d_531_v121_ = _dafny.Map({(self).f15: (self).f4})
                    index72_ = default__.safeIndex(23, (d_528_v118_).length(0))
                    rhs54_ = (d_527_v117_) + (_dafny.SeqWithoutIsStrInference([len(d_530_v120_), 798]))
                    rhs55_ = (d_524_v114_).f12
                    rhs56_ = ((self).f6) * (len(d_531_v121_))
                    lhs31_ = self
                    lhs32_ = d_528_v118_
                    lhs33_ = default__.safeIndex(23, (d_528_v118_).length(0))
                    d_527_v117_ = rhs54_
                    lhs31_.f16 = rhs55_
                    lhs32_[lhs33_] = rhs56_
                    d_532_v122_: _dafny.Array
                    nw85_ = _dafny.Array(False, 10)
                    d_532_v122_ = nw85_
                    index73_ = default__.safeIndex(6, (d_532_v122_).length(0))
                    (d_532_v122_)[index73_] = (True if (self).f4 else (self).f4)
                    index74_ = default__.safeIndex(23, (d_528_v118_).length(0))
                    index75_ = default__.safeIndex(6, (d_532_v122_).length(0))
                    rhs57_ = (0) - (((self).f6) * ((d_528_v118_)[default__.safeIndex(23, (d_528_v118_).length(0))]))
                    rhs58_ = (self).f4
                    lhs34_ = d_528_v118_
                    lhs35_ = default__.safeIndex(23, (d_528_v118_).length(0))
                    lhs36_ = d_532_v122_
                    lhs37_ = default__.safeIndex(6, (d_532_v122_).length(0))
                    lhs34_[lhs35_] = rhs57_
                    lhs36_[lhs37_] = rhs58_
                    pass
            pass
        d_533_v123_: str
        d_533_v123_ = _dafny.CodePoint('a')
        d_534_v124_: _dafny.Map
        d_534_v124_ = _dafny.Map({(self).f15: (self).f4})
        d_535_v125_: _dafny.MultiSet
        d_535_v125_ = _dafny.MultiSet([(self).f4])
        d_536_v126_: D12
        d_536_v126_ = D12_DC28(len(default__.fm40((self).f4, ((d_534_v124_)[((d_535_v125_)[(self).f4] if ((self).f4) in (d_535_v125_) else (self).f6)] if (((d_535_v125_)[(self).f4] if ((self).f4) in (d_535_v125_) else (self).f6)) in (d_534_v124_) else (self).f4), self.f16, globalState)), (0) - ((self).f6), (self).f4, (self).f4, (self).f15)
        pat_let_tv8_ = d_533_v123_
        pat_let_tv9_ = d_533_v123_
        def lambda17_(source12_):
            if source12_.is_DC28:
                d_537___mcc_h13_ = source12_.cf43
                d_538___mcc_h14_ = source12_.cf44
                d_539___mcc_h15_ = source12_.cf45
                d_540___mcc_h16_ = source12_.cf46
                d_541___mcc_h17_ = source12_.cf47
                d_542_cf47_ = d_541___mcc_h17_
                d_543_cf46_ = d_540___mcc_h16_
                d_544_cf45_ = d_539___mcc_h15_
                d_545_cf44_ = d_538___mcc_h14_
                d_546_cf43_ = d_537___mcc_h13_
                return pat_let_tv8_
            elif True:
                d_547___mcc_h18_ = source12_.cf42
                d_548_cf42_ = d_547___mcc_h18_
                return pat_let_tv9_

        d_533_v123_ = lambda17_(d_536_v126_)
        d_549_v127_: _dafny.Seq
        d_549_v127_ = _dafny.SeqWithoutIsStrInference([(self).f15])
        d_550_v128_: _dafny.Set
        d_550_v128_ = _dafny.Set({D6_DC16((self).f4, (self).f4, d_549_v127_), default__.fm43(globalState)})
        d_551_v129_: _dafny.Array
        nw86_ = _dafny.Array(None, 20)
        nw86_[int(0)] = (self).f4
        nw86_[int(1)] = not(((d_534_v124_)[(self).f6] if ((self).f6) in (d_534_v124_) else (self).f4))
        nw86_[int(2)] = not(False)
        nw86_[int(3)] = (p0) < ((0) - ((self).f15))
        nw86_[int(4)] = (self).f4
        nw86_[int(5)] = (self).f4
        nw86_[int(6)] = (self).f4
        nw86_[int(7)] = not((self).fm8((self).f15, (self).f4, d_521_v112_, d_549_v127_, globalState))
        nw86_[int(8)] = (self).f4
        nw86_[int(9)] = (p0) <= ((self).f15)
        nw86_[int(10)] = False
        nw86_[int(11)] = (self).f4
        nw86_[int(12)] = (default__.fm41(len(d_521_v112_), d_533_v123_, (self).f4, globalState)).issubset(default__.fm41(440, d_533_v123_, (self).f4, globalState))
        nw86_[int(13)] = default__.fm0(d_533_v123_, globalState)
        nw86_[int(14)] = (self).f4
        nw86_[int(15)] = False
        nw86_[int(16)] = not ((self).f4) or ((self).f4)
        nw86_[int(17)] = True
        nw86_[int(18)] = (p0) in (default__.fm42(d_533_v123_, d_550_v128_, globalState))
        nw86_[int(19)] = (self).f4
        d_551_v129_ = nw86_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_551_v129_).length(0)):
            d_552_i14_: int = guard_loop_1_
            if (True) and (((0) <= (d_552_i14_)) and ((d_552_i14_) < ((d_551_v129_).length(0)))):
                (d_551_v129_)[(d_552_i14_)] = (self).f4
        d_553_i15_: int
        d_553_i15_ = 0
        with _dafny.label("1"):
            while (self).f4:
                with _dafny.c_label("1"):
                    if (d_553_i15_) >= (100):
                        raise _dafny.Break("1")
                    d_553_i15_ = (d_553_i15_) + (1)
                    (self).f16 = (0) - ((self.f16) * (p0))
                    d_554_v130_: _dafny.Seq
                    d_554_v130_ = _dafny.SeqWithoutIsStrInference([False, (self).f4, (self).f4])
                    d_555_v131_: D4
                    d_555_v131_ = D4_DC13(self.f16, False, p0, (D4_DC13(362, True, self.f16, d_554_v130_)).cf17)
                    d_556_v132_: _dafny.Map
                    d_556_v132_ = _dafny.Map({d_555_v131_: (self).f4})
                    d_556_v132_ = (d_556_v132_).set(d_555_v131_, (self).f4)
                    d_557_v133_: C0
                    nw87_ = C0()
                    nw87_.ctor__(753)
                    d_557_v133_ = nw87_
                    d_557_v133_ = d_557_v133_
                    pass
            pass
        d_558_v134_: bool
        d_558_v134_ = False
        index76_ = default__.safeIndex(395, (d_551_v129_).length(0))
        (d_551_v129_)[index76_] = True
        d_559_v135_: _dafny.Array
        nw88_ = _dafny.Array(int(0), 3)
        d_559_v135_ = nw88_
        d_560_v136_: _dafny.Set
        d_560_v136_ = _dafny.Set({d_559_v135_})
        d_561_v137_: _dafny.Array
        def lambda18_(d_562_i16_):
            return D0_DC1()

        init11_ = lambda18_
        nw89_ = _dafny.Array(None, 20)
        for i0_11_ in range(nw89_.length(0)):
            nw89_[i0_11_] = init11_(i0_11_)
        d_561_v137_ = nw89_
        d_563_v138_: _dafny.Map
        d_563_v138_ = _dafny.Map({d_560_v136_: d_561_v137_})
        d_564_v139_: _dafny.Seq
        d_564_v139_ = _dafny.SeqWithoutIsStrInference([d_563_v138_])
        index77_ = default__.safeIndex(693, (d_559_v135_).length(0))
        (d_559_v135_)[index77_] = len((d_564_v139_)[default__.safeIndex((self).f15, len(d_564_v139_))])
        d_565_v140_: D11
        d_565_v140_ = D11_DC25(d_558_v134_, p0)
        index78_ = default__.safeIndex(395, (d_551_v129_).length(0))
        index79_ = default__.safeIndex(693, (d_559_v135_).length(0))
        rhs59_ = (d_565_v140_).cf40
        rhs60_ = (d_558_v134_) == (d_558_v134_)
        rhs61_ = d_558_v134_
        rhs62_ = p0
        rhs63_ = (self).f4
        lhs38_ = self
        lhs39_ = d_551_v129_
        lhs40_ = default__.safeIndex(395, (d_551_v129_).length(0))
        lhs41_ = d_559_v135_
        lhs42_ = default__.safeIndex(693, (d_559_v135_).length(0))
        lhs38_.f16 = rhs59_
        d_558_v134_ = rhs60_
        lhs39_[lhs40_] = rhs61_
        lhs41_[lhs42_] = rhs62_
        d_558_v134_ = rhs63_
        d_566_v141_: _dafny.Map
        d_566_v141_ = _dafny.Map({default__.fm0(d_533_v123_, globalState): p0})
        r0 = d_566_v141_
        return r0

    def m2(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_567_v0_: _dafny.Array
        nw90_ = _dafny.Array(False, 28)
        d_567_v0_ = nw90_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_567_v0_).length(0)):
            d_568_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_568_i0_)) and ((d_568_i0_) < ((d_567_v0_).length(0)))):
                (d_567_v0_)[(d_568_i0_)] = not ((self).f4) or ((self).f4)
        d_569_i1_: int
        d_569_i1_ = 0
        with _dafny.label("2"):
            while True:
                with _dafny.c_label("2"):
                    if (d_569_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_569_i1_ = (d_569_i1_) + (1)
                    d_570_v1_: _dafny.Map
                    d_570_v1_ = _dafny.Map({(self).f6: True})
                    d_570_v1_ = (d_570_v1_) | (d_570_v1_)
                    d_571_v2_: _dafny.MultiSet
                    d_571_v2_ = _dafny.MultiSet([p0])
                    d_572_v3_: _dafny.Map
                    d_572_v3_ = _dafny.Map({(_dafny.MultiSet([(self).f4])).isdisjoint(d_571_v2_): (self).f6})
                    d_572_v3_ = (default__.fm44(globalState)) | ((d_572_v3_) | (d_572_v3_))
                    d_573_v4_: _dafny.Seq
                    d_573_v4_ = _dafny.SeqWithoutIsStrInference([not((self).f4)])
                    (self).f16 = len(_dafny.Map({(self).f15: len(d_573_v4_)}))
                    d_574_v5_: C0
                    nw91_ = C0()
                    nw91_.ctor__((self).f6)
                    d_574_v5_ = nw91_
                    pass
            pass
        d_575_v6_: _dafny.Seq
        d_575_v6_ = _dafny.SeqWithoutIsStrInference([p0])
        d_576_v7_: D0
        d_576_v7_ = D0_DC2(len(d_575_v6_))
        source13_ = d_576_v7_
        if source13_.is_DC1:
            d_577_v8_: bool
            d_577_v8_ = False
            d_578_v9_: str
            d_578_v9_ = _dafny.CodePoint('n')
            d_577_v8_ = default__.fm0(d_578_v9_, globalState)
            d_579_v10_: _dafny.Set
            d_579_v10_ = _dafny.Set({(self).f4, d_577_v8_, d_577_v8_, False})
            if ((d_579_v10_).intersection(d_579_v10_)) == (d_579_v10_):
                d_580_v11_: _dafny.Seq
                d_580_v11_ = _dafny.SeqWithoutIsStrInference([d_567_v0_])
                d_581_v12_: _dafny.Array
                nw92_ = _dafny.Array(int(0), 1)
                d_581_v12_ = nw92_
                d_582_v13_: _dafny.Map
                d_582_v13_ = _dafny.Map({p0: not((self).f4)})
                d_583_v14_: _dafny.MultiSet
                d_583_v14_ = _dafny.MultiSet([d_578_v9_, d_578_v9_, d_578_v9_])
                d_584_v15_: _dafny.Seq
                d_584_v15_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                d_585_v16_: _dafny.Array
                nw93_ = _dafny.Array(None, 26)
                nw93_[int(0)] = (self).f6
                nw93_[int(1)] = self.f16
                nw93_[int(2)] = default__.fm3((self).f6, p0, globalState)
                nw93_[int(3)] = self.f16
                nw93_[int(4)] = self.f16
                nw93_[int(5)] = len(_dafny.Map({d_581_v12_: (self).fm7(globalState)}))
                nw93_[int(6)] = (self).f6
                nw93_[int(7)] = (self).f15
                nw93_[int(8)] = (self).f6
                nw93_[int(9)] = len(d_582_v13_)
                nw93_[int(10)] = (self).f6
                nw93_[int(11)] = (self).f15
                nw93_[int(12)] = len(d_575_v6_)
                nw93_[int(13)] = (0) - ((self).f15)
                nw93_[int(14)] = self.f16
                nw93_[int(15)] = (self).f15
                nw93_[int(16)] = (self).f6
                nw93_[int(17)] = self.f16
                nw93_[int(18)] = (self).f15
                nw93_[int(19)] = (self).f15
                nw93_[int(20)] = (d_583_v14_).cardinality
                nw93_[int(21)] = len(d_584_v15_)
                nw93_[int(22)] = self.f16
                nw93_[int(23)] = self.f16
                nw93_[int(24)] = -935
                nw93_[int(25)] = 642
                d_585_v16_ = nw93_
                d_586_v17_: _dafny.MultiSet
                d_586_v17_ = _dafny.MultiSet([d_585_v16_, d_585_v16_])
                d_587_v18_: _dafny.Map
                d_587_v18_ = _dafny.Map({not((_dafny.SeqWithoutIsStrInference([d_567_v0_])) < (d_580_v11_)): (((d_586_v17_)[d_581_v12_] if (d_581_v12_) in (d_586_v17_) else (d_584_v15_)[default__.safeIndex(self.f16, len(d_584_v15_))])) + (907)})
                d_588_v19_: _dafny.Map
                d_588_v19_ = _dafny.Map({155: self.f16})
                d_587_v18_ = (d_587_v18_).set((self).f4, ((d_588_v19_)[self.f16] if (self.f16) in (d_588_v19_) else self.f16))
                d_589_v20_: D3
                d_589_v20_ = D3_DC8(default__.fm40(p0, d_577_v8_, len(default__.fm37(d_577_v8_, (self).f4, (self).f6, (self).f4, globalState)), globalState))
                d_590_v21_: D3
                d_590_v21_ = D3_DC11(d_589_v20_)
                d_590_v21_ = d_590_v21_
                index80_ = default__.safeIndex(299, (d_567_v0_).length(0))
                (d_567_v0_)[index80_] = p0
                d_591_v22_: _dafny.Seq
                d_591_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shvveca"))
                index81_ = default__.safeIndex(299, (d_567_v0_).length(0))
                rhs64_ = (0) - ((0) - (len((d_591_v22_) + (d_591_v22_))))
                rhs65_ = default__.fm3((self).f15, d_577_v8_, globalState)
                rhs66_ = d_577_v8_
                rhs67_ = d_577_v8_
                rhs68_ = d_577_v8_
                lhs43_ = self
                lhs44_ = self
                lhs45_ = d_567_v0_
                lhs46_ = default__.safeIndex(299, (d_567_v0_).length(0))
                lhs43_.f16 = rhs64_
                lhs44_.f16 = rhs65_
                d_577_v8_ = rhs66_
                lhs45_[lhs46_] = rhs67_
                d_577_v8_ = rhs68_
                d_592_v23_: _dafny.Array
                nw94_ = _dafny.Array(_dafny.MultiSet({}), 5)
                d_592_v23_ = nw94_
                d_593_v24_: _dafny.Map
                d_593_v24_ = _dafny.Map({d_578_v9_: (self).f4})
                d_594_v25_: _dafny.Seq
                d_594_v25_ = _dafny.SeqWithoutIsStrInference([d_592_v23_, d_592_v23_])
                d_595_v26_: _dafny.Array
                nw95_ = _dafny.Array(None, 20)
                nw95_[int(0)] = d_592_v23_
                nw95_[int(1)] = d_592_v23_
                nw95_[int(2)] = (d_592_v23_ if ((d_593_v24_)[d_578_v9_] if (d_578_v9_) in (d_593_v24_) else (d_567_v0_)[default__.safeIndex(299, (d_567_v0_).length(0))]) else d_592_v23_)
                nw95_[int(3)] = d_592_v23_
                nw95_[int(4)] = d_592_v23_
                nw95_[int(5)] = d_592_v23_
                nw95_[int(6)] = d_592_v23_
                nw95_[int(7)] = d_592_v23_
                nw95_[int(8)] = d_592_v23_
                nw95_[int(9)] = (d_594_v25_)[default__.safeIndex(self.f16, len(d_594_v25_))]
                nw95_[int(10)] = d_592_v23_
                nw95_[int(11)] = d_592_v23_
                nw95_[int(12)] = d_592_v23_
                nw95_[int(13)] = d_592_v23_
                nw95_[int(14)] = d_592_v23_
                nw95_[int(15)] = d_592_v23_
                nw95_[int(16)] = d_592_v23_
                nw95_[int(17)] = d_592_v23_
                nw95_[int(18)] = d_592_v23_
                nw95_[int(19)] = d_592_v23_
                d_595_v26_ = nw95_
                index82_ = default__.safeIndex(525, (d_595_v26_).length(0))
                (d_595_v26_)[index82_] = d_592_v23_
                index83_ = default__.safeIndex(525, (d_595_v26_).length(0))
                (d_595_v26_)[index83_] = d_592_v23_
                d_591_v22_ = ((d_591_v22_) + (d_591_v22_)) + (d_591_v22_)
            elif True:
                d_596_v27_: _dafny.Map
                d_596_v27_ = _dafny.Map({(self).f6: (self).f4})
                d_597_v28_: _dafny.Map
                d_597_v28_ = _dafny.Map({p0: d_596_v27_})
                d_597_v28_ = (d_597_v28_).set(p0, _dafny.Map({871: d_577_v8_}))
                d_598_v29_: _dafny.Map
                d_598_v29_ = _dafny.Map({self.f16: (self).f15})
                (self).f16 = ((d_598_v29_)[(self).f15] if ((self).f15) in (d_598_v29_) else self.f16)
                (self).f16 = (self).f6
                d_599_v30_: _dafny.MultiSet
                d_599_v30_ = _dafny.MultiSet([(self).f15, -639])
                d_600_v31_: _dafny.Array
                nw96_ = _dafny.Array(int(0), 28)
                d_600_v31_ = nw96_
                d_601_v32_: _dafny.Seq
                d_601_v32_ = _dafny.SeqWithoutIsStrInference([d_600_v31_, d_600_v31_, d_600_v31_])
                rhs69_ = (self.f16) + ((self).f6)
                rhs70_ = d_599_v30_
                rhs71_ = (self).f4
                rhs72_ = d_601_v32_
                lhs47_ = self
                lhs47_.f16 = rhs69_
                d_599_v30_ = rhs70_
                d_577_v8_ = rhs71_
                d_601_v32_ = rhs72_
                (self).f16 = self.f16
            (self).f16 = (self.f16) - (len(_dafny.SeqWithoutIsStrInference([d_578_v9_ for d_602_i2_ in range(default__.abs(591))])))
            d_567_v0_ = d_567_v0_
        elif source13_.is_DC2:
            d_603___mcc_h0_ = source13_.cf1
            d_604_cf1_ = d_603___mcc_h0_
            d_605_v33_: bool
            d_605_v33_ = False
            d_606_v34_: _dafny.Map
            d_606_v34_ = _dafny.Map({(self).f4: d_567_v0_})
            d_607_v35_: _dafny.Seq
            d_607_v35_ = _dafny.SeqWithoutIsStrInference([d_567_v0_, ((d_606_v34_)[(self).f4] if ((self).f4) in (d_606_v34_) else d_567_v0_)])
            d_608_v36_: _dafny.Seq
            d_608_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gx"))
            d_609_v37_: _dafny.Map
            d_609_v37_ = _dafny.Map({(d_607_v35_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htbf"))), len(d_607_v35_))]: default__.fm3(len(d_608_v36_), p0, globalState)})
            d_610_v38_: _dafny.Map
            d_610_v38_ = _dafny.Map({d_604_cf1_: (self).f15})
            d_611_v39_: _dafny.Set
            d_611_v39_ = _dafny.Set({d_610_v38_, d_610_v38_})
            d_612_v40_: _dafny.Map
            d_612_v40_ = _dafny.Map({d_605_v33_: p0})
            d_613_v41_: _dafny.Map
            d_613_v41_ = _dafny.Map({p0: d_604_cf1_})
            d_614_v42_: str
            d_614_v42_ = _dafny.CodePoint('r')
            d_615_v43_: _dafny.Set
            d_615_v43_ = _dafny.Set({d_614_v42_, d_614_v42_})
            d_616_v44_: _dafny.Array
            nw97_ = _dafny.Array(None, 24)
            nw97_[int(0)] = self.f16
            nw97_[int(1)] = len(d_612_v40_)
            nw97_[int(2)] = self.f16
            nw97_[int(3)] = 377
            nw97_[int(4)] = d_604_cf1_
            nw97_[int(5)] = len(_dafny.SeqWithoutIsStrInference([(self).f6, d_604_cf1_, (self).f6, len(d_606_v34_), len(d_608_v36_)]))
            nw97_[int(6)] = (self).f15
            nw97_[int(7)] = 495
            nw97_[int(8)] = self.f16
            nw97_[int(9)] = (_dafny.MultiSet([self.f16, self.f16])).cardinality
            nw97_[int(10)] = ((d_613_v41_)[(self).f4] if ((self).f4) in (d_613_v41_) else self.f16)
            nw97_[int(11)] = (self).f6
            nw97_[int(12)] = d_604_cf1_
            nw97_[int(13)] = (self).f6
            nw97_[int(14)] = self.f16
            nw97_[int(15)] = len(d_615_v43_)
            nw97_[int(16)] = self.f16
            nw97_[int(17)] = self.f16
            nw97_[int(18)] = d_604_cf1_
            nw97_[int(19)] = d_604_cf1_
            nw97_[int(20)] = len(d_608_v36_)
            nw97_[int(21)] = len(d_613_v41_)
            nw97_[int(22)] = len(d_608_v36_)
            nw97_[int(23)] = 180
            d_616_v44_ = nw97_
            d_617_v45_: _dafny.Seq
            d_617_v45_ = _dafny.SeqWithoutIsStrInference([d_616_v44_, d_616_v44_, d_616_v44_, d_616_v44_, d_616_v44_])
            d_618_v46_: _dafny.Map
            d_618_v46_ = _dafny.Map({d_608_v36_: len(_dafny.SeqWithoutIsStrInference([d_610_v38_ for d_619_i3_ in range(default__.abs(62))]))})
            d_620_v47_: D1
            d_620_v47_ = D1_DC3(p0)
            d_621_v48_: _dafny.Map
            d_621_v48_ = _dafny.Map({d_620_v47_: d_604_cf1_})
            d_622_v49_: _dafny.Array
            nw98_ = _dafny.Array(None, 26)
            nw98_[int(0)] = default__.safeModuloInt(self.f16, len(d_608_v36_))
            nw98_[int(1)] = (self).f6
            nw98_[int(2)] = (0) - ((self).f15)
            nw98_[int(3)] = default__.safeModuloInt(len(d_575_v6_), default__.fm3(len(d_611_v39_), d_605_v33_, globalState))
            nw98_[int(4)] = default__.safeModuloInt((0) - (d_604_cf1_), self.f16)
            nw98_[int(5)] = 813
            nw98_[int(6)] = len(d_617_v45_)
            nw98_[int(7)] = (self).f6
            nw98_[int(8)] = self.f16
            nw98_[int(9)] = 206
            nw98_[int(10)] = ((self).f15) * ((self).f6)
            nw98_[int(11)] = -867
            nw98_[int(12)] = self.f16
            nw98_[int(13)] = default__.safeDivisionInt((self).f15, (0) - ((0) - (d_604_cf1_)))
            nw98_[int(14)] = default__.safeDivisionInt(d_604_cf1_, 915)
            nw98_[int(15)] = d_604_cf1_
            nw98_[int(16)] = (self).fm9(p0, p0, ((d_618_v46_)[_dafny.SeqWithoutIsStrInference([d_614_v42_ for d_623_i4_ in range(default__.abs(886))])] if (_dafny.SeqWithoutIsStrInference([d_614_v42_ for d_624_i4_ in range(default__.abs(886))])) in (d_618_v46_) else self.f16), globalState)
            nw98_[int(17)] = default__.safeDivisionInt(len(d_575_v6_), self.f16)
            nw98_[int(18)] = d_604_cf1_
            nw98_[int(19)] = (self).f15
            nw98_[int(20)] = len(d_621_v48_)
            nw98_[int(21)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eadjey")))
            nw98_[int(22)] = d_604_cf1_
            nw98_[int(23)] = default__.safeModuloInt(len(d_575_v6_), len(d_608_v36_))
            nw98_[int(24)] = (self).f15
            nw98_[int(25)] = ((d_613_v41_)[d_605_v33_] if (d_605_v33_) in (d_613_v41_) else -222)
            d_622_v49_ = nw98_
            index84_ = default__.safeIndex(987, (d_616_v44_).length(0))
            (d_616_v44_)[index84_] = d_604_cf1_
            d_625_v50_: _dafny.Seq
            d_625_v50_ = _dafny.SeqWithoutIsStrInference([676, d_604_cf1_])
            index85_ = default__.safeIndex(987, (d_616_v44_).length(0))
            rhs73_ = not(((self).f4) == (p0))
            rhs74_ = (self).f4
            rhs75_ = d_609_v37_
            rhs76_ = (((d_613_v41_)[d_605_v33_] if (d_605_v33_) in (d_613_v41_) else 94)) + (default__.safeDivisionInt(d_604_cf1_, len(d_625_v50_)))
            rhs77_ = d_575_v6_
            lhs48_ = d_616_v44_
            lhs49_ = default__.safeIndex(987, (d_616_v44_).length(0))
            d_605_v33_ = rhs73_
            d_605_v33_ = rhs74_
            d_609_v37_ = rhs75_
            lhs48_[lhs49_] = rhs76_
            d_575_v6_ = rhs77_
            rhs78_ = d_605_v33_
            rhs79_ = d_608_v36_
            d_605_v33_ = rhs78_
            d_608_v36_ = rhs79_
            d_626_v51_: _dafny.Map
            d_626_v51_ = _dafny.Map({d_607_v35_: (self).f4})
            d_627_v52_: int
            out16_: int
            out16_ = (self).m11(((d_626_v51_)[d_607_v35_] if (d_607_v35_) in (d_626_v51_) else d_605_v33_), (d_625_v50_) < (_dafny.SeqWithoutIsStrInference([(self).f6 for d_628_i5_ in range(default__.abs(-844))])), (d_618_v46_) | (d_618_v46_), d_567_v0_, globalState)
            d_627_v52_ = out16_
            if p0:
                d_605_v33_ = (d_575_v6_)[default__.safeIndex(len(d_608_v36_), len(d_575_v6_))]
                index86_ = default__.safeIndex(513, (d_567_v0_).length(0))
                (d_567_v0_)[index86_] = d_605_v33_
                index87_ = default__.safeIndex(513, (d_567_v0_).length(0))
                (d_567_v0_)[index87_] = d_605_v33_
                d_605_v33_ = (self).f4
                d_612_v40_ = (d_612_v40_).set(p0, default__.fm0(d_614_v42_, globalState))
                d_627_v52_ = (628) * ((0) - (((0) - (len(d_610_v38_))) * ((d_616_v44_)[default__.safeIndex(987, (d_616_v44_).length(0))])))
            elif True:
                index88_ = default__.safeIndex(987, (d_616_v44_).length(0))
                (d_616_v44_)[index88_] = default__.safeModuloInt(default__.safeModuloInt(d_604_cf1_, (self).f15), (d_625_v50_)[default__.safeIndex(self.f16, len(d_625_v50_))])
                index89_ = default__.safeIndex(413, (d_567_v0_).length(0))
                (d_567_v0_)[index89_] = (d_575_v6_)[default__.safeIndex((self).f6, len(d_575_v6_))]
                index90_ = default__.safeIndex(413, (d_567_v0_).length(0))
                (d_567_v0_)[index90_] = (not(d_605_v33_) if default__.fm0(d_614_v42_, globalState) else not(True))
                d_629_v53_: _dafny.MultiSet
                d_629_v53_ = _dafny.MultiSet([d_605_v33_, d_605_v33_])
                d_630_v54_: D11
                d_630_v54_ = D11_DC25((d_567_v0_)[default__.safeIndex(413, (d_567_v0_).length(0))], self.f16)
                index91_ = default__.safeIndex(413, (d_567_v0_).length(0))
                index92_ = default__.safeIndex(987, (d_616_v44_).length(0))
                rhs80_ = ((_dafny.MultiSet(d_575_v6_)).ispropersubset(d_629_v53_)) in ((_dafny.SeqWithoutIsStrInference([(self).f4, p0])) + (_dafny.SeqWithoutIsStrInference([True, (self).f4])))
                rhs81_ = ((self).f15) * ((d_630_v54_).cf40)
                rhs82_ = d_614_v42_
                lhs50_ = d_567_v0_
                lhs51_ = default__.safeIndex(413, (d_567_v0_).length(0))
                lhs52_ = d_616_v44_
                lhs53_ = default__.safeIndex(987, (d_616_v44_).length(0))
                lhs50_[lhs51_] = rhs80_
                lhs52_[lhs53_] = rhs81_
                r0 = rhs82_
                index93_ = default__.safeIndex(413, (d_567_v0_).length(0))
                rhs83_ = (0) - (len((d_610_v38_) | (_dafny.Map({(d_616_v44_)[default__.safeIndex(987, (d_616_v44_).length(0))]: d_604_cf1_}))))
                rhs84_ = (d_616_v44_)[default__.safeIndex(987, (d_616_v44_).length(0))]
                rhs85_ = d_622_v49_
                rhs86_ = not(not(d_605_v33_))
                lhs54_ = d_567_v0_
                lhs55_ = default__.safeIndex(413, (d_567_v0_).length(0))
                d_627_v52_ = rhs83_
                d_627_v52_ = rhs84_
                d_622_v49_ = rhs85_
                lhs54_[lhs55_] = rhs86_
                d_605_v33_ = (d_567_v0_)[default__.safeIndex(413, (d_567_v0_).length(0))]
        elif True:
            d_631___mcc_h1_ = source13_.cf0
            d_632_cf0_ = d_631___mcc_h1_
            d_633_v55_: C0
            nw99_ = C0()
            nw99_.ctor__((self.f16) + (d_632_cf0_))
            d_633_v55_ = nw99_
            d_634_v56_: bool
            d_634_v56_ = True
            d_635_v57_: _dafny.Array
            nw100_ = _dafny.Array(int(0), 4)
            d_635_v57_ = nw100_
            d_636_v58_: str
            d_636_v58_ = _dafny.CodePoint('c')
            d_637_v59_: _dafny.Map
            d_637_v59_ = _dafny.Map({d_633_v55_: p0})
            d_638_v60_: _dafny.Seq
            d_638_v60_ = _dafny.SeqWithoutIsStrInference([d_636_v58_])
            d_639_v61_: _dafny.MultiSet
            d_639_v61_ = _dafny.MultiSet([d_567_v0_, d_567_v0_])
            d_640_v62_: _dafny.Seq
            d_640_v62_ = _dafny.SeqWithoutIsStrInference([(d_639_v61_).cardinality])
            d_641_v63_: _dafny.Seq
            d_641_v63_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_635_v57_: d_636_v58_})), (self).f6, default__.fm3((self).f15, (self).fm8((0) - (d_632_cf0_), not(((d_637_v59_)[d_633_v55_] if (d_633_v55_) in (d_637_v59_) else p0)), d_638_v60_, d_640_v62_, globalState), globalState)])
            d_642_v64_: _dafny.Array
            nw101_ = _dafny.Array(None, 4)
            nw101_[int(0)] = (0) - ((d_633_v55_).f12)
            nw101_[int(1)] = (self).f15
            nw101_[int(2)] = (d_641_v63_)[default__.safeIndex(d_632_cf0_, len(d_641_v63_))]
            nw101_[int(3)] = d_632_cf0_
            d_642_v64_ = nw101_
            index94_ = default__.safeIndex(78, (d_635_v57_).length(0))
            (d_635_v57_)[index94_] = (d_633_v55_).f12
            d_643_v65_: _dafny.MultiSet
            d_643_v65_ = _dafny.MultiSet([p0, d_634_v56_])
            index95_ = default__.safeIndex(78, (d_635_v57_).length(0))
            rhs87_ = ((d_643_v65_).issubset(d_643_v65_) if d_634_v56_ else (self).f4)
            rhs88_ = (self.f16 if p0 else self.f16)
            rhs89_ = d_632_cf0_
            lhs56_ = d_635_v57_
            lhs57_ = default__.safeIndex(78, (d_635_v57_).length(0))
            d_634_v56_ = rhs87_
            d_632_cf0_ = rhs88_
            lhs56_[lhs57_] = rhs89_
            d_644_v66_: _dafny.Map
            d_644_v66_ = _dafny.Map({d_632_cf0_: (d_635_v57_)[default__.safeIndex(78, (d_635_v57_).length(0))]})
            d_645_v67_: _dafny.MultiSet
            d_645_v67_ = _dafny.MultiSet([d_632_cf0_, -806, len(d_575_v6_), ((d_644_v66_)[self.f16] if (self.f16) in (d_644_v66_) else (d_633_v55_).f12)])
            d_646_v68_: _dafny.MultiSet
            d_646_v68_ = d_645_v67_
            d_647_v69_: _dafny.Map
            d_647_v69_ = _dafny.Map({d_646_v68_: ((_dafny.Map({(d_635_v57_)[default__.safeIndex(78, (d_635_v57_).length(0))]: (self).f15}))) | (d_644_v66_)})
            d_647_v69_ = d_647_v69_
            index96_ = default__.safeIndex(78, (d_635_v57_).length(0))
            (d_635_v57_)[index96_] = ((d_635_v57_)[default__.safeIndex(78, (d_635_v57_).length(0))]) + (len(d_640_v62_))
        index97_ = default__.safeIndex(243, (d_567_v0_).length(0))
        (d_567_v0_)[index97_] = (self).f4
        d_648_v70_: _dafny.Seq
        d_648_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uehfebb"))
        d_649_v71_: str
        d_649_v71_ = _dafny.CodePoint('v')
        index98_ = default__.safeIndex(243, (d_567_v0_).length(0))
        (d_567_v0_)[index98_] = ((d_649_v71_) not in (d_648_v70_) if (d_648_v70_) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_650_i6_ in range(default__.abs(272))])) else (self).f4)
        d_651_v73_: _dafny.MultiSet
        d_651_v73_ = _dafny.MultiSet([(self).f6])
        def iife59_():
            coll33_ = _dafny.Map()
            compr_33_: int
            for compr_33_ in (d_651_v73_).Elements:
                d_665_v72_: int = compr_33_
                if (d_665_v72_) in (d_651_v73_):
                    coll33_[(d_665_v72_) + ((0) - ((d_651_v73_).cardinality))] = ((d_651_v73_)[(self).f6] if ((self).f6) in (d_651_v73_) else (self).f6)
            return _dafny.Map(coll33_)
        hi2_ = (self.f16) - (self.f16)
        for d_652_i7_ in range((0) - (len(iife59_()
        )), hi2_):
            d_653_v75_: _dafny.Array
            def lambda19_(d_654_i8_):
                def iife54_():
                    coll32_ = _dafny.Set()
                    compr_32_: int
                    for compr_32_ in _dafny.IntegerRange(739, 930):
                        d_655_v74_: int = compr_32_
                        if ((739) <= (d_655_v74_)) and ((d_655_v74_) < (930)):
                            coll32_ = coll32_.union(_dafny.Set([default__.safeDivisionInt(d_655_v74_, self.f16)]))
                    return _dafny.Set(coll32_)
                return (d_654_i8_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.Set({904}), iife54_()
                , _dafny.Set({(self).f6, (self).f15})])))

            init12_ = lambda19_
            nw102_ = _dafny.Array(None, 6)
            for i0_12_ in range(nw102_.length(0)):
                nw102_[i0_12_] = init12_(i0_12_)
            d_653_v75_ = nw102_
            index99_ = default__.safeIndex(117, (d_653_v75_).length(0))
            (d_653_v75_)[index99_] = (self).fm9(p0, (self).f4, d_652_i7_, globalState)
            index100_ = default__.safeIndex(117, (d_653_v75_).length(0))
            (d_653_v75_)[index100_] = (self).f15
            d_656_v76_: C0
            nw103_ = C0()
            nw103_.ctor__((d_653_v75_)[default__.safeIndex(117, (d_653_v75_).length(0))])
            d_656_v76_ = nw103_
            d_657_v77_: _dafny.MultiSet
            d_657_v77_ = _dafny.MultiSet([p0, (d_567_v0_)[default__.safeIndex(243, (d_567_v0_).length(0))]])
            d_658_v78_: _dafny.Seq
            d_658_v78_ = _dafny.SeqWithoutIsStrInference([d_657_v77_, (_dafny.MultiSet([True, (self).f4, (self).f4, (d_567_v0_)[default__.safeIndex(243, (d_567_v0_).length(0))], (self).f4])).intersection(d_657_v77_)])
            d_658_v78_ = ((d_658_v78_) + (default__.fm45(p0, (self).f15, (d_567_v0_)[default__.safeIndex(243, (d_567_v0_).length(0))], (d_656_v76_).f12, globalState))) + (_dafny.SeqWithoutIsStrInference([d_657_v77_]))
            d_659_v79_: D1
            d_659_v79_ = D1_DC3((self).fm8(779, (d_567_v0_)[default__.safeIndex(243, (d_567_v0_).length(0))], d_648_v70_, _dafny.SeqWithoutIsStrInference([(d_653_v75_)[default__.safeIndex(117, (d_653_v75_).length(0))] for d_660_i9_ in range(default__.abs(520))]), globalState))
            d_661_v80_: _dafny.Array
            nw104_ = _dafny.Array(_dafny.Array(None, 0), 4)
            d_661_v80_ = nw104_
            pat_let_tv10_ = p0
            d_662_v81_: _dafny.Map
            def iife55_(_pat_let11_0):
                def iife56_(d_663_dt__update__tmp_h0_):
                    def iife57_(_pat_let12_0):
                        def iife58_(d_664_dt__update_hcf2_h0_):
                            return D1_DC3(d_664_dt__update_hcf2_h0_)
                        return iife58_(_pat_let12_0)
                    return iife57_(pat_let_tv10_)
                return iife56_(_pat_let11_0)
            d_662_v81_ = _dafny.Map({(iife55_(d_659_v79_)).cf2: d_661_v80_})
            d_662_v81_ = (d_662_v81_).set(not(p0), d_661_v80_)
        index101_ = default__.safeIndex(243, (d_567_v0_).length(0))
        (d_567_v0_)[index101_] = (self).f4
        r0 = d_649_v71_
        d_666_v82_: _dafny.Array
        def lambda20_(d_667_i10_):
            return D1_DC3((self).f4)

        init13_ = lambda20_
        nw105_ = _dafny.Array(None, 3)
        for i0_13_ in range(nw105_.length(0)):
            nw105_[i0_13_] = init13_(i0_13_)
        d_666_v82_ = nw105_
        r1 = d_666_v82_
        return r0, r1

    def m11(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_668_v0_: _dafny.Seq
        d_668_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
        d_668_v0_ = d_668_v0_
        d_669_v1_: bool
        d_669_v1_ = False
        arr0_ = self.f5
        index102_ = default__.safeIndex(595, (self.f5).length(0))
        arr0_[index102_] = p3
        d_670_v2_: _dafny.Seq
        d_670_v2_ = _dafny.SeqWithoutIsStrInference([d_669_v1_])
        d_671_v3_: D1
        d_671_v3_ = D1_DC4((self).f15, self.f16)
        pat_let_tv11_ = p1
        pat_let_tv12_ = p0
        arr1_ = self.f5
        index103_ = default__.safeIndex(595, (self.f5).length(0))
        def lambda21_(source14_):
            if source14_.is_DC4:
                d_672___mcc_h0_ = source14_.cf3
                d_673___mcc_h1_ = source14_.cf4
                d_674_cf4_ = d_673___mcc_h1_
                d_675_cf3_ = d_672___mcc_h0_
                return pat_let_tv11_
            elif True:
                d_676___mcc_h2_ = source14_.cf2
                d_677_cf2_ = d_676___mcc_h2_
                return pat_let_tv12_

        def iife60_(_pat_let13_0):
            def iife61_(d_678_dt__update__tmp_h0_):
                def iife62_(_pat_let14_0):
                    def iife63_(d_679_dt__update_hcf3_h0_):
                        return D1_DC4(d_679_dt__update_hcf3_h0_, (d_678_dt__update__tmp_h0_).cf4)
                    return iife63_(_pat_let14_0)
                return iife62_(-948)
            return iife61_(_pat_let13_0)
        rhs90_ = (d_670_v2_)[default__.safeIndex(((self).f15) - ((self).f6), len(d_670_v2_))]
        rhs91_ = p3
        rhs92_ = p0
        rhs93_ = lambda21_(iife60_(d_671_v3_))
        rhs94_ = False
        lhs58_ = self.f5
        lhs59_ = default__.safeIndex(595, (self.f5).length(0))
        d_669_v1_ = rhs90_
        lhs58_[lhs59_] = rhs91_
        d_669_v1_ = rhs92_
        d_669_v1_ = rhs93_
        d_669_v1_ = rhs94_
        d_680_v4_: str
        d_680_v4_ = _dafny.CodePoint('d')
        d_681_v5_: _dafny.Map
        d_681_v5_ = _dafny.Map({d_680_v4_: p0})
        d_682_i0_: int
        d_682_i0_ = 0
        with _dafny.label("3"):
            while ((len(d_681_v5_)) > ((self).f6) if (d_668_v0_) != (d_668_v0_) else p1):
                with _dafny.c_label("3"):
                    if (d_682_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_682_i0_ = (d_682_i0_) + (1)
                    d_683_v6_: _dafny.Array
                    nw106_ = _dafny.Array(int(0), 14)
                    d_683_v6_ = nw106_
                    d_683_v6_ = d_683_v6_
                    d_684_v7_: _dafny.Map
                    d_684_v7_ = _dafny.Map({(self).f6: (0) - ((self).f6)})
                    d_685_v8_: C0
                    nw107_ = C0()
                    nw107_.ctor__((len(d_684_v7_) if (self).f4 else (self).f6))
                    d_685_v8_ = nw107_
                    d_669_v1_ = (default__.fm40((self).f4, p1, (d_685_v8_).f12, globalState)) == (_dafny.SeqWithoutIsStrInference([d_669_v1_]))
                    d_686_v9_: _dafny.Set
                    d_686_v9_ = _dafny.Set({(d_669_v1_) or ((self).f4), p0})
                    d_686_v9_ = d_686_v9_
                    pass
            pass
        if False:
            d_687_v10_: _dafny.Set
            d_687_v10_ = _dafny.Set({p3})
            d_688_v11_: _dafny.Array
            nw108_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_688_v11_ = nw108_
            d_689_v12_: _dafny.Array
            nw109_ = _dafny.Array(None, 12)
            nw109_[int(0)] = d_688_v11_
            nw109_[int(1)] = d_688_v11_
            nw109_[int(2)] = d_688_v11_
            nw109_[int(3)] = d_688_v11_
            nw109_[int(4)] = d_688_v11_
            nw109_[int(5)] = d_688_v11_
            nw109_[int(6)] = d_688_v11_
            nw109_[int(7)] = d_688_v11_
            nw109_[int(8)] = d_688_v11_
            nw109_[int(9)] = d_688_v11_
            nw109_[int(10)] = d_688_v11_
            nw109_[int(11)] = d_688_v11_
            d_689_v12_ = nw109_
            index104_ = default__.safeIndex(401, (d_689_v12_).length(0))
            (d_689_v12_)[index104_] = d_688_v11_
            d_690_v13_: _dafny.Map
            d_690_v13_ = _dafny.Map({default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([self.f16])), (self).fm9((self).f4, p1, (self).f15, globalState)): (self).f15})
            d_691_v14_: _dafny.Seq
            d_691_v14_ = _dafny.SeqWithoutIsStrInference([d_688_v11_, d_688_v11_, d_688_v11_, d_688_v11_])
            d_692_v15_: _dafny.MultiSet
            d_692_v15_ = _dafny.MultiSet([p0])
            index105_ = default__.safeIndex(401, (d_689_v12_).length(0))
            rhs95_ = d_687_v10_
            rhs96_ = (d_691_v14_)[default__.safeIndex(((d_692_v15_)[not(False)] if (not(False)) in (d_692_v15_) else (self).f15), len(d_691_v14_))]
            rhs97_ = (d_690_v13_) | (d_690_v13_)
            lhs60_ = d_689_v12_
            lhs61_ = default__.safeIndex(401, (d_689_v12_).length(0))
            d_687_v10_ = rhs95_
            lhs60_[lhs61_] = rhs96_
            d_690_v13_ = rhs97_
            (self).f16 = (self).fm7(globalState)
            if (self).f4:
                d_693_v16_: _dafny.Array
                nw110_ = _dafny.Array(int(0), 6)
                d_693_v16_ = nw110_
                d_694_v17_: _dafny.Map
                d_694_v17_ = _dafny.Map({len(d_668_v0_): d_693_v16_})
                d_694_v17_ = d_694_v17_
                d_695_v18_: _dafny.MultiSet
                d_695_v18_ = _dafny.MultiSet([(self).f15])
                d_695_v18_ = ((d_695_v18_).intersection(_dafny.MultiSet([(self).f15]))).set(default__.safeDivisionInt((0) - (len(d_668_v0_)), (self).f15), default__.abs((self).f15))
                arr2_ = (self.f5)[default__.safeIndex(595, (self.f5).length(0))]
                index106_ = default__.safeIndex(253, ((self.f5)[default__.safeIndex(595, (self.f5).length(0))]).length(0))
                arr2_[index106_] = (self).f4
                d_696_v19_: _dafny.Map
                d_696_v19_ = _dafny.Map({p0: d_668_v0_})
                d_697_v20_: _dafny.Seq
                d_697_v20_ = _dafny.SeqWithoutIsStrInference([(0) - (self.f16)])
                d_698_v21_: _dafny.Map
                d_698_v21_ = _dafny.Map({True: (0) - (len(d_697_v20_))})
                d_699_v22_: D12
                d_699_v22_ = D12_DC28((self).f6, self.f16, d_669_v1_, not(d_669_v1_), ((d_698_v21_)[(self).f4] if ((self).f4) in (d_698_v21_) else 146))
                arr3_ = (self.f5)[default__.safeIndex(595, (self.f5).length(0))]
                index107_ = default__.safeIndex(253, ((self.f5)[default__.safeIndex(595, (self.f5).length(0))]).length(0))
                rhs98_ = len(d_696_v19_)
                rhs99_ = (d_699_v22_).cf44
                rhs100_ = p0
                lhs62_ = self
                lhs63_ = (self.f5)[default__.safeIndex(595, (self.f5).length(0))]
                lhs64_ = default__.safeIndex(253, ((self.f5)[default__.safeIndex(595, (self.f5).length(0))]).length(0))
                lhs62_.f16 = rhs98_
                r0 = rhs99_
                lhs63_[lhs64_] = rhs100_
                d_700_v23_: _dafny.Seq
                d_700_v23_ = _dafny.SeqWithoutIsStrInference([d_668_v0_, d_668_v0_])
                d_669_v1_ = not(((d_700_v23_) + (d_700_v23_)) <= (d_700_v23_))
                arr4_ = (self.f5)[default__.safeIndex(595, (self.f5).length(0))]
                index108_ = default__.safeIndex(253, ((self.f5)[default__.safeIndex(595, (self.f5).length(0))]).length(0))
                arr4_[index108_] = not((self).f4)
            elif True:
                r0 = ((self).f15) * (self.f16)
                d_701_v24_: D0
                d_701_v24_ = D0_DC2((0) - ((self).f6))
                d_701_v24_ = d_701_v24_
                (self).f16 = self.f16
                d_702_v25_: _dafny.Map
                d_702_v25_ = _dafny.Map({d_670_v2_: p0})
                d_703_v26_: _dafny.Set
                d_703_v26_ = _dafny.Set({(self).f15})
                d_702_v25_ = (d_702_v25_).set(d_670_v2_, (d_703_v26_).issubset(d_703_v26_))
                r0 = len(d_668_v0_)
            d_704_v27_: D0
            d_704_v27_ = D0_DC0((self).f6)
            d_705_v28_: C0
            nw111_ = C0()
            nw111_.ctor__(len(d_670_v2_))
            d_705_v28_ = nw111_
            d_706_v29_: _dafny.Map
            d_706_v29_ = _dafny.Map({d_705_v28_: d_669_v1_})
            d_707_v30_: _dafny.Map
            d_707_v30_ = _dafny.Map({(self).f15: d_706_v29_})
            d_708_v31_: _dafny.Seq
            d_708_v31_ = _dafny.SeqWithoutIsStrInference([-159, (0) - (self.f16)])
            rhs101_ = d_704_v27_
            rhs102_ = ((_dafny.Map({d_705_v28_: (self).f4})).set(d_705_v28_, not((self).f4))) | (((d_707_v30_)[(self).f15] if ((self).f15) in (d_707_v30_) else d_706_v29_))
            rhs103_ = (self).f6
            rhs104_ = (self).fm8(((self).f15) * ((d_705_v28_).f12), ((d_708_v31_)[default__.safeIndex((self).f6, len(d_708_v31_))]) <= ((d_705_v28_).f12), d_668_v0_, d_708_v31_, globalState)
            rhs105_ = (p1) and (p0)
            d_704_v27_ = rhs101_
            d_706_v29_ = rhs102_
            r0 = rhs103_
            d_669_v1_ = rhs104_
            d_669_v1_ = rhs105_
            d_709_v32_: _dafny.Map
            d_709_v32_ = _dafny.Map({_dafny.MultiSet([(d_705_v28_).f12]): ((d_692_v15_)[not(d_669_v1_)] if (not(d_669_v1_)) in (d_692_v15_) else (self).fm7(globalState))})
            d_710_v33_: _dafny.Map
            d_710_v33_ = _dafny.Map({self.f16: d_705_v28_})
            d_711_v34_: _dafny.Map
            d_711_v34_ = _dafny.Map({(d_710_v33_) != (d_710_v33_): d_709_v32_})
            d_709_v32_ = ((d_711_v34_)[p0] if (p0) in (d_711_v34_) else d_709_v32_)
        elif True:
            (self).f16 = self.f16
            d_712_v35_: C0
            nw112_ = C0()
            nw112_.ctor__(520)
            d_712_v35_ = nw112_
            d_713_v36_: _dafny.MultiSet
            d_713_v36_ = _dafny.MultiSet([p1])
            d_714_v37_: D3
            d_714_v37_ = D3_DC10((d_713_v36_).cardinality)
            d_715_v38_: _dafny.Map
            d_715_v38_ = _dafny.Map({self.f16: ((d_714_v37_).cf11) == ((d_712_v35_).f12)})
            d_716_v39_: _dafny.Seq
            d_716_v39_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_717_v40_: _dafny.Seq
            d_717_v40_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_712_v35_).f12]), d_716_v39_, d_716_v39_, _dafny.SeqWithoutIsStrInference([(0) - ((d_712_v35_).f12) for d_718_i1_ in range(default__.abs(-982))]), d_716_v39_])
            d_715_v38_ = (d_715_v38_).set((self).f6, (d_717_v40_) != (d_717_v40_))
            d_719_v41_: _dafny.Array
            def lambda22_(d_720_i3_):
                return (d_720_i3_) - ((self).f6)

            init14_ = lambda22_
            nw113_ = _dafny.Array(None, 17)
            for i0_14_ in range(nw113_.length(0)):
                nw113_[i0_14_] = init14_(i0_14_)
            d_719_v41_ = nw113_
            d_721_v42_: _dafny.Map
            d_721_v42_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wirva"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_722_i2_ in range(default__.abs(-993))])): _dafny.Map({p1: d_719_v41_})})
            d_723_v43_: _dafny.Map
            d_723_v43_ = _dafny.Map({p0: d_719_v41_})
            d_721_v42_ = (d_721_v42_).set(d_668_v0_, (d_723_v43_) | (d_723_v43_))
            d_669_v1_ = (self).f4
        d_724_v44_: _dafny.Map
        d_724_v44_ = _dafny.Map({p0: self.f16})
        r0 = (self.f16) + (default__.safeDivisionInt(len(d_724_v44_), -690))
        d_725_v45_: _dafny.Map
        d_725_v45_ = _dafny.Map({d_680_v4_: (self).f15})
        if (d_725_v45_) == ((d_725_v45_) | ((d_725_v45_).set(d_680_v4_, 220))):
            (self).f16 = default__.safeModuloInt(self.f16, (self).f15)
            d_726_v46_: _dafny.Seq
            d_726_v46_ = _dafny.SeqWithoutIsStrInference([(len(d_668_v0_)) * (668), (self.f16) * ((self).f15)])
            d_727_v47_: D12
            d_727_v47_ = D12_DC27(_dafny.SeqWithoutIsStrInference([d_670_v2_]))
            d_726_v46_ = default__.fm39(d_668_v0_, d_727_v47_, p1, globalState)
            d_728_v48_: _dafny.MultiSet
            d_728_v48_ = _dafny.MultiSet([not(False), p0])
            d_729_v49_: _dafny.MultiSet
            d_729_v49_ = _dafny.MultiSet([self.f16, self.f16, (d_728_v48_).cardinality])
            d_730_v50_: D9
            d_730_v50_ = D9_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imf")))
            d_731_v51_: _dafny.Map
            d_731_v51_ = _dafny.Map({d_729_v49_: (d_730_v50_).cf25})
            d_732_v52_: _dafny.Set
            d_732_v52_ = _dafny.Set({self.f16})
            d_733_v53_: bool
            d_734_v54_: int
            out17_: bool
            out18_: int
            out17_, out18_ = default__.m0(d_731_v51_, (d_732_v52_).issubset(_dafny.Set({(self).f6})), d_729_v49_, d_669_v1_, globalState)
            d_733_v53_ = out17_
            d_734_v54_ = out18_
            d_735_v55_: _dafny.Map
            d_735_v55_ = _dafny.Map({(self).f6: (d_726_v46_)[default__.safeIndex(len(d_668_v0_), len(d_726_v46_))]})
            d_736_v56_: _dafny.Array
            nw114_ = _dafny.Array(None, 6)
            nw114_[int(0)] = (len(d_668_v0_)) + ((0) - ((self).f6))
            nw114_[int(1)] = 398
            nw114_[int(2)] = 62
            nw114_[int(3)] = default__.safeModuloInt(self.f16, (0) - ((d_728_v48_).cardinality))
            nw114_[int(4)] = ((d_735_v55_)[(self).f6] if ((self).f6) in (d_735_v55_) else self.f16)
            nw114_[int(5)] = (d_734_v54_) * (self.f16)
            d_736_v56_ = nw114_
            rhs106_ = d_736_v56_
            rhs107_ = p1
            rhs108_ = (self.f16) + ((len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csd"))) for d_737_i4_ in range(default__.abs(311))]))) + (len(default__.fm46(d_680_v4_, globalState))))
            rhs109_ = ((d_728_v48_)[d_669_v1_] if (d_669_v1_) in (d_728_v48_) else (self).f6)
            lhs65_ = self
            lhs66_ = self
            d_736_v56_ = rhs106_
            d_733_v53_ = rhs107_
            lhs65_.f16 = rhs108_
            lhs66_.f16 = rhs109_
            if (self).f4:
                d_738_v57_: bool
                d_739_v58_: int
                out19_: bool
                out20_: int
                out19_, out20_ = default__.m0((d_731_v51_).set(d_729_v49_, d_668_v0_), p0, d_729_v49_, p1, globalState)
                d_738_v57_ = out19_
                d_739_v58_ = out20_
                d_740_v60_: _dafny.Set
                d_740_v60_ = _dafny.Set({d_729_v49_})
                d_741_v61_: bool
                d_742_v62_: int
                out21_: bool
                out22_: int
                def iife64_():
                    coll34_ = _dafny.Map()
                    compr_34_: _dafny.MultiSet
                    for compr_34_ in (d_740_v60_).Elements:
                        d_743_v59_: _dafny.MultiSet = compr_34_
                        if (d_743_v59_) in (d_740_v60_):
                            coll34_[d_743_v59_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgs"))
                    return _dafny.Map(coll34_)
                out21_, out22_ = default__.m0(iife64_()
                , (len(d_724_v44_)) >= (580), (d_729_v49_) | (d_729_v49_), d_669_v1_, globalState)
                d_741_v61_ = out21_
                d_742_v62_ = out22_
                d_669_v1_ = p1
                d_741_v61_ = p1
                r0 = d_734_v54_
            elif True:
                d_744_v63_: _dafny.Map
                d_744_v63_ = _dafny.Map({(self).f15: p1})
                d_724_v44_ = (d_724_v44_).set((self).f4, ((self).f15) * ((self).fm9(False, p1, len(d_744_v63_), globalState)))
                d_745_v64_: _dafny.Map
                d_745_v64_ = _dafny.Map({d_733_v53_: p0})
                d_746_v65_: _dafny.Seq
                d_746_v65_ = _dafny.SeqWithoutIsStrInference([d_729_v49_])
                d_747_v66_: _dafny.Map
                d_747_v66_ = _dafny.Map({(d_745_v64_) | (d_745_v64_): ((d_746_v65_)[default__.safeIndex((0) - (self.f16), len(d_746_v65_))]).cardinality})
                d_747_v66_ = d_747_v66_
                d_748_v67_: C0
                nw115_ = C0()
                nw115_.ctor__((self.f16) + (843))
                d_748_v67_ = nw115_
                d_728_v48_ = ((d_728_v48_) | (d_728_v48_)) | ((d_728_v48_) | (default__.fm47((self).f6, globalState)))
                d_669_v1_ = d_733_v53_
        elif True:
            d_669_v1_ = (p0) or (p1)
            d_680_v4_ = d_680_v4_
            d_749_v68_: _dafny.Array
            def lambda23_(d_750_i5_):
                return (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f6, -594, self.f16])).cardinality, (0) - ((self).f15)]) if (self).f4 else _dafny.SeqWithoutIsStrInference([self.f16]))

            init15_ = lambda23_
            nw116_ = _dafny.Array(None, 15)
            for i0_15_ in range(nw116_.length(0)):
                nw116_[i0_15_] = init15_(i0_15_)
            d_749_v68_ = nw116_
            d_751_v69_: _dafny.Map
            d_751_v69_ = _dafny.Map({p1: d_669_v1_})
            d_752_v70_: _dafny.Seq
            d_752_v70_ = _dafny.SeqWithoutIsStrInference([len(d_751_v69_)])
            d_753_v71_: _dafny.Seq
            d_753_v71_ = _dafny.SeqWithoutIsStrInference([(self).f15, len(d_752_v70_), (self).f6])
            index109_ = default__.safeIndex(489, (d_749_v68_).length(0))
            (d_749_v68_)[index109_] = d_753_v71_
            d_754_v72_: _dafny.Map
            d_754_v72_ = _dafny.Map({default__.fm0(d_680_v4_, globalState): d_680_v4_})
            index110_ = default__.safeIndex(489, (d_749_v68_).length(0))
            (d_749_v68_)[index110_] = (((_dafny.SeqWithoutIsStrInference([(self).f15])) + (d_753_v71_)).set(default__.safeIndex(876, len((_dafny.SeqWithoutIsStrInference([(self).f15])) + (d_753_v71_))), len(d_754_v72_))) + (_dafny.SeqWithoutIsStrInference([self.f16, (self).f15]))
            d_755_v73_: _dafny.Map
            d_755_v73_ = _dafny.Map({(d_668_v0_) + (d_668_v0_): (self).f6})
            d_755_v73_ = (d_755_v73_).set(d_668_v0_, (self.f16) + (424))
            d_669_v1_ = p0
        r0 = self.f16
        return r0

    @property
    def f15(self):
        return self._f15

class C2(T0):
    def  __init__(self):
        self._f4: bool = False
        self._f17: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f4(self):
        return self._f4
    def ctor__(self, f17, f4):
        (self)._f17 = f17
        (self)._f4 = f4

    def fm6(self, p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([486]))

    def fm7(self, globalState):
        return 831

    def fm48(self, p0, p1, p2, globalState):
        return (D1_DC3((self).f4)).cf2

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        hi3_ = p0
        for d_756_i0_ in range((self).fm7(globalState), hi3_):
            d_757_v0_: _dafny.Array
            nw117_ = _dafny.Array(False, 18)
            d_757_v0_ = nw117_
            def lambda24_(d_758_i1_):
                return (self).f4

            init16_ = lambda24_
            nw118_ = _dafny.Array(None, 2)
            for i0_16_ in range(nw118_.length(0)):
                nw118_[i0_16_] = init16_(i0_16_)
            d_757_v0_ = nw118_
            d_759_v1_: int
            d_759_v1_ = 805
            d_759_v1_ = -78
            d_760_v2_: _dafny.Set
            d_760_v2_ = _dafny.Set({(self).f4})
            d_761_v3_: int
            d_762_v4_: int
            out23_: int
            out24_: int
            out23_, out24_ = (self).m12(d_759_v1_, (d_760_v2_) - (d_760_v2_), _dafny.Map({d_759_v1_: len(d_760_v2_)}), (self).f4, globalState)
            d_761_v3_ = out23_
            d_762_v4_ = out24_
            if (self).f4:
                d_763_v5_: bool
                d_763_v5_ = True
                d_763_v5_ = (d_760_v2_).ispropersubset(_dafny.Set({(self).f4, (self).f4, (self).f4}))
                d_764_v6_: _dafny.Set
                d_764_v6_ = _dafny.Set({d_756_i0_, d_761_v3_})
                d_765_v7_: _dafny.MultiSet
                d_765_v7_ = _dafny.MultiSet([False, d_763_v5_, d_763_v5_, d_763_v5_])
                d_766_v8_: str
                d_766_v8_ = _dafny.CodePoint('p')
                d_767_v9_: _dafny.Set
                d_767_v9_ = _dafny.Set({d_764_v6_, (d_764_v6_).intersection(_dafny.Set({d_756_i0_, (d_765_v7_).cardinality, len(((self).f17).set(default__.safeIndex(d_762_v4_, len((self).f17)), d_766_v8_)), ((d_765_v7_).set(d_763_v5_, default__.abs((0) - (p0)))).cardinality}))})
                d_768_v10_: D11
                d_768_v10_ = D11_DC26(d_763_v5_)
                d_769_v11_: _dafny.Map
                d_769_v11_ = _dafny.Map({d_768_v10_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjgsgbnw"))) + ((self).f17)})
                d_770_v12_: _dafny.Array
                nw119_ = _dafny.Array(D2.default()(), 16)
                d_770_v12_ = nw119_
                index111_ = default__.safeIndex(838, (d_770_v12_).length(0))
                (d_770_v12_)[index111_] = D2_DC6()
                d_771_v14_: _dafny.MultiSet
                d_771_v14_ = _dafny.MultiSet([d_768_v10_])
                d_772_v15_: D2
                d_772_v15_ = D2_DC6()
                index112_ = default__.safeIndex(838, (d_770_v12_).length(0))
                def iife65_():
                    coll35_ = _dafny.Map()
                    compr_35_: D11
                    for compr_35_ in (d_771_v14_).Elements:
                        d_773_v13_: D11 = compr_35_
                        if (d_773_v13_) in (d_771_v14_):
                            coll35_[d_773_v13_] = (self).f17
                    return _dafny.Map(coll35_)
                rhs110_ = (d_767_v9_ if False else (d_767_v9_) - (d_767_v9_))
                rhs111_ = ((d_769_v11_) | (d_769_v11_)) | ((iife65_()
                ) | (d_769_v11_))
                rhs112_ = d_772_v15_
                rhs113_ = (not((self).f4)) and ((d_759_v1_) >= (d_759_v1_))
                rhs114_ = (self).f4
                lhs67_ = d_770_v12_
                lhs68_ = default__.safeIndex(838, (d_770_v12_).length(0))
                d_767_v9_ = rhs110_
                d_769_v11_ = rhs111_
                lhs67_[lhs68_] = rhs112_
                d_763_v5_ = rhs113_
                d_763_v5_ = rhs114_
                d_774_v16_: C0
                nw120_ = C0()
                nw120_.ctor__(-980)
                d_774_v16_ = nw120_
                d_763_v5_ = not(((self).f4 if d_763_v5_ else d_763_v5_))
                d_775_v17_: D2
                d_775_v17_ = D2_DC5(d_757_v0_)
                d_776_v18_: _dafny.Map
                d_776_v18_ = _dafny.Map({(d_763_v5_) or ((self).f4): (d_775_v17_).cf5})
                d_776_v18_ = (d_776_v18_).set(not (d_763_v5_) or (d_763_v5_), d_757_v0_)
            elif True:
                d_777_v19_: _dafny.MultiSet
                d_777_v19_ = _dafny.MultiSet([d_756_i0_])
                d_777_v19_ = (d_777_v19_).intersection(d_777_v19_)
                d_778_v20_: _dafny.Array
                def lambda25_(d_779_v4_):
                    def lambda26_(d_780_i2_):
                        return (d_780_i2_) - (d_779_v4_)

                    return lambda26_

                init17_ = lambda25_(d_762_v4_)
                nw121_ = _dafny.Array(None, 21)
                for i0_17_ in range(nw121_.length(0)):
                    nw121_[i0_17_] = init17_(i0_17_)
                d_778_v20_ = nw121_
                index113_ = default__.safeIndex(217, (d_778_v20_).length(0))
                (d_778_v20_)[index113_] = d_756_i0_
                d_781_v21_: _dafny.Seq
                d_781_v21_ = _dafny.SeqWithoutIsStrInference([d_761_v3_])
                d_782_v22_: _dafny.Seq
                d_782_v22_ = _dafny.SeqWithoutIsStrInference([d_781_v21_])
                d_783_v23_: bool
                d_783_v23_ = False
                d_784_v24_: _dafny.Seq
                d_784_v24_ = _dafny.SeqWithoutIsStrInference([(self).f4, d_783_v23_])
                d_785_v25_: _dafny.Seq
                d_785_v25_ = _dafny.SeqWithoutIsStrInference([(self).f17, (self).f17])
                index114_ = default__.safeIndex(217, (d_778_v20_).length(0))
                rhs115_ = (self).fm7(globalState)
                rhs116_ = ((default__.fm49((self).f17, d_784_v24_, d_761_v3_, globalState)).set(default__.safeIndex(len((d_785_v25_)[default__.safeIndex(d_761_v3_, len(d_785_v25_))]), len(default__.fm49((self).f17, d_784_v24_, d_761_v3_, globalState))), _dafny.SeqWithoutIsStrInference([d_756_i0_ for d_786_i3_ in range(default__.abs(345))]))) + (_dafny.SeqWithoutIsStrInference([d_781_v21_]))
                rhs117_ = True
                rhs118_ = (0) - ((d_761_v3_) + ((0) - (d_759_v1_)))
                rhs119_ = (self).f4
                lhs69_ = d_778_v20_
                lhs70_ = default__.safeIndex(217, (d_778_v20_).length(0))
                lhs69_[lhs70_] = rhs115_
                d_782_v22_ = rhs116_
                d_783_v23_ = rhs117_
                d_762_v4_ = rhs118_
                d_783_v23_ = rhs119_
                d_787_v26_: _dafny.MultiSet
                d_787_v26_ = _dafny.MultiSet([d_783_v23_, d_783_v23_, False])
                index115_ = default__.safeIndex(217, (d_778_v20_).length(0))
                (d_778_v20_)[index115_] = (0) - (((d_787_v26_)[d_783_v23_] if (d_783_v23_) in (d_787_v26_) else d_762_v4_))
                d_760_v2_ = _dafny.Set({not((self).f4)})
                d_788_v27_: _dafny.Array
                nw122_ = _dafny.Array(_dafny.CodePoint('D'), 1)
                d_788_v27_ = nw122_
                index116_ = default__.safeIndex(331, (d_788_v27_).length(0))
                (d_788_v27_)[index116_] = _dafny.CodePoint('l')
                d_789_v28_: str
                d_789_v28_ = _dafny.CodePoint('c')
                index117_ = default__.safeIndex(331, (d_788_v27_).length(0))
                (d_788_v27_)[index117_] = d_789_v28_
        d_790_v29_: _dafny.Array
        nw123_ = _dafny.Array(False, 25)
        d_790_v29_ = nw123_
        d_791_v30_: D10
        d_791_v30_ = D10_DC22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdyfjcy")), ((0) - (p0)) == (122), not((self).f4), d_790_v29_, (0) - (p0))
        source15_ = d_791_v30_
        if source15_.is_DC22:
            d_792___mcc_h0_ = source15_.cf30
            d_793___mcc_h1_ = source15_.cf31
            d_794___mcc_h2_ = source15_.cf32
            d_795___mcc_h3_ = source15_.cf33
            d_796___mcc_h4_ = source15_.cf34
            d_797_cf34_ = d_796___mcc_h4_
            d_798_cf33_ = d_795___mcc_h3_
            d_799_cf32_ = d_794___mcc_h2_
            d_800_cf31_ = d_793___mcc_h1_
            d_801_cf30_ = d_792___mcc_h0_
            d_797_cf34_ = (default__.safeDivisionInt(d_797_cf34_, p0) if d_800_cf31_ else 569)
            d_802_v31_: _dafny.Map
            d_802_v31_ = _dafny.Map({(d_799_cf32_ if (self).f4 else (self).f4): d_797_cf34_})
            d_803_v32_: _dafny.Seq
            d_803_v32_ = _dafny.SeqWithoutIsStrInference([p0])
            d_804_v33_: _dafny.Set
            d_804_v33_ = _dafny.Set({d_799_cf32_, True})
            d_805_v34_: _dafny.Array
            nw124_ = _dafny.Array(None, 9)
            nw124_[int(0)] = p0
            nw124_[int(1)] = d_797_cf34_
            nw124_[int(2)] = len(d_803_v32_)
            nw124_[int(3)] = 156
            nw124_[int(4)] = p0
            nw124_[int(5)] = p0
            nw124_[int(6)] = len(d_801_cf30_)
            nw124_[int(7)] = (self).fm7(globalState)
            nw124_[int(8)] = len(d_804_v33_)
            d_805_v34_ = nw124_
            d_806_v35_: _dafny.Map
            d_806_v35_ = _dafny.Map({d_805_v34_: p0})
            d_802_v31_ = (d_802_v31_).set((_dafny.Map({d_805_v34_: (self).fm7(globalState)})) == ((d_806_v35_).set(d_805_v34_, -360)), p0)
            d_807_v36_: D2
            d_807_v36_ = D2_DC7((d_797_cf34_) + (p0))
            rhs120_ = p0
            rhs121_ = d_799_cf32_
            rhs122_ = d_807_v36_
            d_797_cf34_ = rhs120_
            d_799_cf32_ = rhs121_
            d_807_v36_ = rhs122_
            d_808_v37_: C0
            nw125_ = C0()
            nw125_.ctor__(p0)
            d_808_v37_ = nw125_
        elif True:
            d_809___mcc_h5_ = source15_.cf29
            d_810_cf29_ = d_809___mcc_h5_
            d_811_v38_: D2
            d_811_v38_ = D2_DC6()
            source16_ = d_811_v38_
            if source16_.is_DC6:
                d_812_v39_: str
                d_812_v39_ = _dafny.CodePoint('w')
                d_813_v40_: D9
                d_813_v40_ = D9_DC20((self).f4, d_812_v39_, len((self).f17))
                d_814_v41_: _dafny.Seq
                d_814_v41_ = _dafny.SeqWithoutIsStrInference([d_813_v40_])
                d_815_v42_: _dafny.Seq
                d_815_v42_ = d_814_v41_
                d_814_v41_ = (d_814_v41_)
                index118_ = default__.safeIndex(751, (d_790_v29_).length(0))
                (d_790_v29_)[index118_] = ((self).f4 if (self).f4 else (self).f4)
                d_816_v43_: _dafny.Array
                nw126_ = _dafny.Array(int(0), 16)
                d_816_v43_ = nw126_
                index119_ = default__.safeIndex(437, (d_816_v43_).length(0))
                (d_816_v43_)[index119_] = default__.safeDivisionInt(p0, p0)
                d_817_v44_: _dafny.Array
                nw127_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_817_v44_ = nw127_
                d_818_v45_: bool
                d_818_v45_ = True
                index120_ = default__.safeIndex(751, (d_790_v29_).length(0))
                index121_ = default__.safeIndex(437, (d_816_v43_).length(0))
                rhs123_ = (self).f4
                rhs124_ = p0
                rhs125_ = d_817_v44_
                rhs126_ = not(d_818_v45_)
                lhs71_ = d_790_v29_
                lhs72_ = default__.safeIndex(751, (d_790_v29_).length(0))
                lhs73_ = d_816_v43_
                lhs74_ = default__.safeIndex(437, (d_816_v43_).length(0))
                lhs71_[lhs72_] = rhs123_
                lhs73_[lhs74_] = rhs124_
                d_817_v44_ = rhs125_
                d_818_v45_ = rhs126_
                d_819_v46_: _dafny.Array
                nw128_ = _dafny.Array(False, 14)
                d_819_v46_ = nw128_
                d_819_v46_ = d_819_v46_
                d_820_v47_: _dafny.Seq
                d_820_v47_ = _dafny.SeqWithoutIsStrInference([(d_790_v29_)[default__.safeIndex(751, (d_790_v29_).length(0))]])
                d_821_v48_: _dafny.Seq
                d_821_v48_ = _dafny.SeqWithoutIsStrInference([len(d_820_v47_), p0])
                d_822_v49_: _dafny.MultiSet
                d_822_v49_ = _dafny.MultiSet([False, (d_790_v29_)[default__.safeIndex(751, (d_790_v29_).length(0))]])
                d_823_v50_: _dafny.Map
                d_823_v50_ = _dafny.Map({(d_821_v48_)[default__.safeIndex(len((self).f17), len(d_821_v48_))]: ((0) - (p0) if (d_790_v29_)[default__.safeIndex(751, (d_790_v29_).length(0))] else (d_822_v49_).cardinality)})
                d_824_v51_: _dafny.Set
                d_824_v51_ = _dafny.Set({(d_816_v43_)[default__.safeIndex(437, (d_816_v43_).length(0))]})
                d_825_v52_: _dafny.MultiSet
                d_825_v52_ = _dafny.MultiSet([len(d_824_v51_), p0])
                d_823_v50_ = (d_823_v50_).set((d_816_v43_)[default__.safeIndex(437, (d_816_v43_).length(0))], (460 if d_818_v45_ else (d_825_v52_).cardinality))
            elif source16_.is_DC7:
                d_826___mcc_h6_ = source16_.cf6
                d_827_cf6_ = d_826___mcc_h6_
                d_828_v53_: _dafny.Seq
                d_828_v53_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f4, (self).f4, not(True)})])
                d_828_v53_ = d_828_v53_
                d_829_v54_: _dafny.Seq
                d_829_v54_ = _dafny.SeqWithoutIsStrInference([(self).f4])
                d_830_v55_: _dafny.Seq
                d_830_v55_ = _dafny.SeqWithoutIsStrInference([((d_829_v54_)[default__.safeIndex((0) - (d_827_cf6_), len(d_829_v54_))] if (self).f4 else False)])
                d_830_v55_ = d_829_v54_
                d_831_v56_: _dafny.MultiSet
                d_831_v56_ = _dafny.MultiSet([p0, p0])
                d_832_v57_: _dafny.Map
                d_832_v57_ = _dafny.Map({((self).f17 if (self).fm48((self).f4, d_827_cf6_, p0, globalState) else default__.fm50(892, p0, (d_831_v56_).cardinality, (self).f4, globalState)): p0})
                d_833_v58_: _dafny.Seq
                d_833_v58_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_834_v59_: _dafny.Map
                d_834_v59_ = _dafny.Map({len(d_833_v58_): (self).f17})
                d_835_v60_: _dafny.Seq
                d_835_v60_ = _dafny.SeqWithoutIsStrInference([len(d_834_v59_)])
                d_836_v61_: _dafny.Seq
                d_836_v61_ = _dafny.SeqWithoutIsStrInference([325, d_827_cf6_, p0, (d_835_v60_)[default__.safeIndex(d_827_cf6_, len(d_835_v60_))]])
                d_832_v57_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_837_i4_ in range(default__.abs(228))]): ((d_836_v61_)[default__.safeIndex(p0, len(d_836_v61_))]) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_838_i5_ in range(default__.abs(761))])))})
                d_839_v62_: bool
                d_839_v62_ = False
                d_840_v63_: _dafny.Array
                def lambda27_(d_841_cf6_):
                    def lambda28_(d_842_i6_):
                        return default__.safeDivisionInt(d_842_i6_, d_841_cf6_)

                    return lambda28_

                init18_ = lambda27_(d_827_cf6_)
                nw129_ = _dafny.Array(None, 22)
                for i0_18_ in range(nw129_.length(0)):
                    nw129_[i0_18_] = init18_(i0_18_)
                d_840_v63_ = nw129_
                d_843_v64_: _dafny.Map
                d_843_v64_ = _dafny.Map({p0: p0})
                index122_ = default__.safeIndex(0, (d_840_v63_).length(0))
                (d_840_v63_)[index122_] = (len(_dafny.SeqWithoutIsStrInference([d_827_cf6_]))) * (((d_843_v64_)[p0] if (p0) in (d_843_v64_) else d_827_cf6_))
                index123_ = default__.safeIndex(0, (d_840_v63_).length(0))
                rhs127_ = d_839_v62_
                rhs128_ = ((d_843_v64_)[d_827_cf6_] if (d_827_cf6_) in (d_843_v64_) else p0)
                rhs129_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hasb"))) < ((self).f17)
                lhs75_ = d_840_v63_
                lhs76_ = default__.safeIndex(0, (d_840_v63_).length(0))
                d_839_v62_ = rhs127_
                lhs75_[lhs76_] = rhs128_
                d_839_v62_ = rhs129_
            elif True:
                d_844___mcc_h7_ = source16_.cf5
                d_845_cf5_ = d_844___mcc_h7_
                d_846_v65_: _dafny.Array
                nw130_ = _dafny.Array(_dafny.Set({}), 10)
                d_846_v65_ = nw130_
                d_847_v66_: _dafny.Map
                d_847_v66_ = _dafny.Map({(self).f4: (self).f4})
                d_848_v67_: _dafny.Set
                d_848_v67_ = _dafny.Set({d_847_v66_})
                index124_ = default__.safeIndex(210, (d_846_v65_).length(0))
                (d_846_v65_)[index124_] = (d_848_v67_ if (self).f4 else d_848_v67_)
                d_849_v68_: _dafny.MultiSet
                d_849_v68_ = _dafny.MultiSet([(self).f4])
                index125_ = default__.safeIndex(210, (d_846_v65_).length(0))
                def iife66_():
                    coll36_ = _dafny.Set()
                    compr_36_: _dafny.Map
                    for compr_36_ in ((_dafny.Map({d_847_v66_: _dafny.MultiSet([False])})).set(d_847_v66_, d_849_v68_)).keys.Elements:
                        d_850_v69_: _dafny.Map = compr_36_
                        if (d_850_v69_) in ((_dafny.Map({d_847_v66_: _dafny.MultiSet([False])})).set(d_847_v66_, d_849_v68_)):
                            coll36_ = coll36_.union(_dafny.Set([d_850_v69_]))
                    return _dafny.Set(coll36_)
                (d_846_v65_)[index125_] = iife66_()
                
                d_851_v70_: _dafny.Map
                d_851_v70_ = _dafny.Map({p0: p0})
                d_852_v71_: _dafny.Set
                d_852_v71_ = _dafny.Set({(self).f4, (self).f4})
                d_853_v72_: _dafny.Array
                nw131_ = _dafny.Array(None, 16)
                nw131_[int(0)] = len((d_851_v70_).set(p0, p0))
                nw131_[int(1)] = p0
                nw131_[int(2)] = 245
                nw131_[int(3)] = p0
                nw131_[int(4)] = len(d_852_v71_)
                nw131_[int(5)] = p0
                nw131_[int(6)] = p0
                nw131_[int(7)] = p0
                nw131_[int(8)] = p0
                nw131_[int(9)] = p0
                nw131_[int(10)] = len(d_810_cf29_)
                nw131_[int(11)] = p0
                nw131_[int(12)] = p0
                nw131_[int(13)] = p0
                nw131_[int(14)] = len((self).f17)
                nw131_[int(15)] = len((d_851_v70_) | (d_851_v70_))
                d_853_v72_ = nw131_
                index126_ = default__.safeIndex(918, (d_853_v72_).length(0))
                (d_853_v72_)[index126_] = p0
                index127_ = default__.safeIndex(918, (d_853_v72_).length(0))
                (d_853_v72_)[index127_] = 844
                d_854_v73_: C0
                nw132_ = C0()
                nw132_.ctor__(((d_853_v72_)[default__.safeIndex(918, (d_853_v72_).length(0))]) * ((d_853_v72_)[default__.safeIndex(918, (d_853_v72_).length(0))]))
                d_854_v73_ = nw132_
                d_855_v74_: _dafny.Map
                d_855_v74_ = _dafny.Map({74: d_852_v71_})
                d_856_v75_: _dafny.Seq
                d_856_v75_ = _dafny.SeqWithoutIsStrInference([(d_854_v73_).f12, len(d_855_v74_)])
                d_857_v76_: _dafny.Seq
                d_857_v76_ = _dafny.SeqWithoutIsStrInference([((0) - (len(d_856_v75_))) - (p0), (d_854_v73_).f12])
                d_857_v76_ = d_857_v76_
            d_858_v77_: _dafny.Seq
            d_858_v77_ = _dafny.SeqWithoutIsStrInference([(self).f4])
            d_859_v78_: _dafny.Seq
            d_859_v78_ = _dafny.SeqWithoutIsStrInference([d_858_v77_, d_858_v77_, _dafny.SeqWithoutIsStrInference([(self).f4, (self).f4])])
            d_858_v77_ = (D4_DC13(p0, True, (0) - (len((self).f17)), (d_859_v78_)[default__.safeIndex(p0, len(d_859_v78_))])).cf17
            d_860_v79_: _dafny.Array
            nw133_ = _dafny.Array(int(0), 3)
            d_860_v79_ = nw133_
            d_861_v80_: _dafny.Map
            d_861_v80_ = _dafny.Map({d_860_v79_: (self).f4})
            d_861_v80_ = (d_861_v80_).set(d_860_v79_, (p0) >= (p0))
            if ((p0) <= (p0)) or ((p0) >= (p0)):
                d_862_v81_: _dafny.Array
                nw134_ = _dafny.Array(None, 4)
                nw134_[int(0)] = d_790_v29_
                nw134_[int(1)] = d_790_v29_
                nw134_[int(2)] = d_790_v29_
                nw134_[int(3)] = d_790_v29_
                d_862_v81_ = nw134_
                d_863_v82_: C1
                nw135_ = C1()
                nw135_.ctor__(p0, p0, d_862_v81_, (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bw")))) * (206), (self).f4)
                d_863_v82_ = nw135_
                d_864_v83_: _dafny.Seq
                d_864_v83_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxhimf")))])
                d_865_v84_: _dafny.Map
                d_865_v84_ = _dafny.Map({d_864_v83_: not ((self).f4) or ((self).f4)})
                d_866_v86_: _dafny.Seq
                d_866_v86_ = _dafny.SeqWithoutIsStrInference([d_864_v83_, d_864_v83_])
                def iife67_():
                    coll37_ = _dafny.Map()
                    compr_37_: _dafny.Seq
                    for compr_37_ in (d_866_v86_).Elements:
                        d_867_v85_: _dafny.Seq = compr_37_
                        if (d_867_v85_) in (d_866_v86_):
                            coll37_[d_867_v85_] = (self).f4
                    return _dafny.Map(coll37_)
                d_865_v84_ = ((iife67_()
                ) | (d_865_v84_)) | ((d_865_v84_) | (d_865_v84_))
                d_868_v87_: C0
                nw136_ = C0()
                nw136_.ctor__(p0)
                d_868_v87_ = nw136_
                d_869_v88_: bool
                d_869_v88_ = False
                d_870_v89_: _dafny.Array
                nw137_ = _dafny.Array(None, 4)
                nw137_[int(0)] = ((self).f17) + ((self).f17)
                nw137_[int(1)] = (self).f17
                nw137_[int(2)] = ((self).f17) + ((self).f17)
                nw137_[int(3)] = (self).f17
                d_870_v89_ = nw137_
                index128_ = default__.safeIndex(575, (d_870_v89_).length(0))
                (d_870_v89_)[index128_] = (self).f17
                d_871_v90_: str
                d_871_v90_ = _dafny.CodePoint('e')
                d_872_v91_: _dafny.Map
                d_872_v91_ = _dafny.Map({d_871_v90_: d_790_v29_})
                d_873_v92_: _dafny.MultiSet
                d_873_v92_ = _dafny.MultiSet([d_790_v29_, ((d_872_v91_)[_dafny.CodePoint('o')] if (_dafny.CodePoint('o')) in (d_872_v91_) else d_790_v29_), d_790_v29_, d_790_v29_, (d_790_v29_ if d_869_v88_ else d_790_v29_)])
                d_874_v93_: _dafny.Map
                d_874_v93_ = _dafny.Map({(d_863_v82_).f15: len((self).f17)})
                index129_ = default__.safeIndex(575, (d_870_v89_).length(0))
                rhs130_ = ((d_873_v92_)[d_790_v29_] if (d_790_v29_) in (d_873_v92_) else (((d_874_v93_)[(0) - (d_863_v82_.f16)] if ((0) - (d_863_v82_.f16)) in (d_874_v93_) else (d_864_v83_)[default__.safeIndex(p0, len(d_864_v83_))])) - ((d_863_v82_).f15))
                rhs131_ = ((d_868_v87_).f12) != ((d_868_v87_).f12)
                rhs132_ = d_869_v88_
                rhs133_ = d_869_v88_
                rhs134_ = _dafny.SeqWithoutIsStrInference([d_871_v90_ for d_875_i7_ in range(default__.abs(-804))])
                lhs77_ = d_863_v82_
                lhs78_ = d_870_v89_
                lhs79_ = default__.safeIndex(575, (d_870_v89_).length(0))
                lhs77_.f16 = rhs130_
                d_869_v88_ = rhs131_
                d_869_v88_ = rhs132_
                d_869_v88_ = rhs133_
                lhs78_[lhs79_] = rhs134_
                d_871_v90_ = _dafny.CodePoint('x')
            elif True:
                d_876_v94_: int
                d_876_v94_ = 297
                d_876_v94_ = p0
                d_877_v95_: bool
                d_877_v95_ = False
                d_878_v96_: D1
                d_878_v96_ = D1_DC4(p0, p0)
                d_879_v97_: _dafny.Array
                nw138_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_879_v97_ = nw138_
                d_880_v98_: C1
                nw139_ = C1()
                nw139_.ctor__(p0, p0, d_879_v97_, d_876_v94_, (self).f4)
                d_880_v98_ = nw139_
                d_881_v99_: _dafny.Seq
                d_881_v99_ = _dafny.SeqWithoutIsStrInference([d_880_v98_])
                d_882_v100_: _dafny.Set
                d_882_v100_ = _dafny.Set({(self).f17, (self).f17})
                d_883_v101_: _dafny.Map
                d_883_v101_ = _dafny.Map({d_882_v100_: d_881_v99_})
                d_884_v102_: _dafny.Map
                d_884_v102_ = _dafny.Map({(self).f4: d_877_v95_})
                d_885_v103_: _dafny.Set
                d_885_v103_ = _dafny.Set({(self).f4, ((d_884_v102_)[d_877_v95_] if (d_877_v95_) in (d_884_v102_) else True)})
                d_886_v104_: str
                d_886_v104_ = _dafny.CodePoint('e')
                d_887_v105_: _dafny.Map
                d_887_v105_ = _dafny.Map({d_886_v104_: d_885_v103_})
                rhs135_ = p0
                rhs136_ = (d_881_v99_) != (((d_883_v101_)[d_882_v100_] if (d_882_v100_) in (d_883_v101_) else d_881_v99_))
                rhs137_ = d_878_v96_
                rhs138_ = ((d_885_v103_) | (((d_887_v105_)[d_886_v104_] if (d_886_v104_) in (d_887_v105_) else _dafny.Set({not(True)})))).issubset(_dafny.Set({d_877_v95_, (self).f4}))
                rhs139_ = p0
                d_876_v94_ = rhs135_
                d_877_v95_ = rhs136_
                d_878_v96_ = rhs137_
                d_877_v95_ = rhs138_
                d_876_v94_ = rhs139_
                d_877_v95_ = False
                d_888_v106_: _dafny.Seq
                d_888_v106_ = _dafny.SeqWithoutIsStrInference([d_860_v79_])
                d_889_v107_: D2
                d_889_v107_ = D2_DC5(d_790_v29_)
                d_877_v95_ = ((0) - (default__.safeDivisionInt(len(d_888_v106_), d_876_v94_))) <= (default__.safeDivisionInt((0) - (p0), (0) - (len(_dafny.SeqWithoutIsStrInference([d_889_v107_])))))
                d_890_v108_: _dafny.Map
                d_890_v108_ = _dafny.Map({-14: (d_880_v98_).f15})
                d_891_v109_: int
                d_892_v110_: int
                out25_: int
                out26_: int
                out25_, out26_ = (self).m12(p0, d_885_v103_, (d_890_v108_) | (d_890_v108_), (self).f4, globalState)
                d_891_v109_ = out25_
                d_892_v110_ = out26_
        d_893_v111_: C0
        nw140_ = C0()
        nw140_.ctor__(p0)
        d_893_v111_ = nw140_
        d_894_i8_: int
        d_894_i8_ = 0
        with _dafny.label("4"):
            while ((self).f4) and ((self).f4):
                with _dafny.c_label("4"):
                    if (d_894_i8_) >= (100):
                        raise _dafny.Break("4")
                    d_894_i8_ = (d_894_i8_) + (1)
                    index130_ = default__.safeIndex(978, (d_790_v29_).length(0))
                    (d_790_v29_)[index130_] = (self).f4
                    index131_ = default__.safeIndex(978, (d_790_v29_).length(0))
                    (d_790_v29_)[index131_] = (self).f4
                    d_895_v112_: _dafny.Map
                    d_895_v112_ = _dafny.Map({p0: not((self).f4)})
                    d_896_v113_: _dafny.Map
                    d_896_v113_ = _dafny.Map({(self).f4: (d_790_v29_)[default__.safeIndex(978, (d_790_v29_).length(0))]})
                    d_897_v114_: _dafny.Set
                    d_897_v114_ = _dafny.Set({(self).f4})
                    d_895_v112_ = default__.fm51(_dafny.Map({((d_896_v113_)[default__.fm0(_dafny.CodePoint('c'), globalState)] if (default__.fm0(_dafny.CodePoint('c'), globalState)) in (d_896_v113_) else False): d_897_v114_}), True, globalState)
                    d_898_v115_: _dafny.Seq
                    d_898_v115_ = _dafny.SeqWithoutIsStrInference([((self).f17) <= ((self).f17), ((d_896_v113_)[(self).f4] if ((self).f4) in (d_896_v113_) else (d_790_v29_)[default__.safeIndex(978, (d_790_v29_).length(0))]), (d_790_v29_)[default__.safeIndex(978, (d_790_v29_).length(0))], (self).f4])
                    d_899_v116_: D11
                    d_899_v116_ = D11_DC26((d_790_v29_)[default__.safeIndex(978, (d_790_v29_).length(0))])
                    d_900_v117_: _dafny.MultiSet
                    d_900_v117_ = _dafny.MultiSet([(self).f4])
                    d_901_v118_: D16
                    d_901_v118_ = D16_DC33((d_899_v116_).cf41, (d_893_v111_).f12, d_900_v117_, p0, d_897_v114_)
                    d_902_v119_: D4
                    d_902_v119_ = D4_DC13(len((_dafny.SeqWithoutIsStrInference([d_901_v118_])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([d_901_v118_]))), d_901_v118_)), (d_790_v29_)[default__.safeIndex(978, (d_790_v29_).length(0))], p0, d_898_v115_)
                    pat_let_tv13_ = p0
                    pat_let_tv14_ = d_898_v115_
                    def iife68_(_pat_let15_0):
                        def iife69_(d_903_dt__update__tmp_h1_):
                            def iife70_(_pat_let16_0):
                                def iife71_(d_904_dt__update_hcf14_h0_):
                                    def iife72_(_pat_let17_0):
                                        def iife73_(d_905_dt__update_hcf17_h0_):
                                            return D4_DC13(d_904_dt__update_hcf14_h0_, (d_903_dt__update__tmp_h1_).cf15, (d_903_dt__update__tmp_h1_).cf16, d_905_dt__update_hcf17_h0_)
                                        return iife73_(_pat_let17_0)
                                    return iife72_(pat_let_tv14_)
                                return iife71_(_pat_let16_0)
                            return iife70_(pat_let_tv13_)
                        return iife69_(_pat_let15_0)
                    d_898_v115_ = ((iife68_(d_902_v119_)).cf17) + (d_898_v115_)
                    index132_ = default__.safeIndex(978, (d_790_v29_).length(0))
                    index133_ = default__.safeIndex(978, (d_790_v29_).length(0))
                    rhs140_ = ((d_897_v114_ if (self).f4 else d_897_v114_)).issubset(d_897_v114_)
                    rhs141_ = not((self).f4)
                    lhs80_ = d_790_v29_
                    lhs81_ = default__.safeIndex(978, (d_790_v29_).length(0))
                    lhs82_ = d_790_v29_
                    lhs83_ = default__.safeIndex(978, (d_790_v29_).length(0))
                    lhs80_[lhs81_] = rhs140_
                    lhs82_[lhs83_] = rhs141_
                    pass
            pass
        d_906_v120_: int
        d_906_v120_ = 906
        d_906_v120_ = default__.safeModuloInt(d_906_v120_, (d_893_v111_).f12)
        d_906_v120_ = (len((self).f17)) + (d_906_v120_)
        d_907_v121_: _dafny.Map
        d_907_v121_ = _dafny.Map({((d_893_v111_).fm17((self).f4, globalState) if (self).f4 else (self).f4): ((d_893_v111_).f12) - (518)})
        r0 = d_907_v121_
        return r0

    def m2(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_908_v0_: _dafny.Seq
        d_908_v0_ = _dafny.SeqWithoutIsStrInference([(self).f4, (self).f4])
        d_909_v1_: D3
        d_909_v1_ = D3_DC8(d_908_v0_)
        d_910_v2_: _dafny.Array
        nw141_ = _dafny.Array(None, 26)
        nw141_[int(0)] = d_909_v1_
        nw141_[int(1)] = D3_DC8(d_908_v0_)
        nw141_[int(2)] = d_909_v1_
        nw141_[int(3)] = d_909_v1_
        nw141_[int(4)] = D3_DC8(d_908_v0_)
        nw141_[int(5)] = D3_DC8(d_908_v0_)
        nw141_[int(6)] = d_909_v1_
        nw141_[int(7)] = d_909_v1_
        nw141_[int(8)] = d_909_v1_
        nw141_[int(9)] = d_909_v1_
        nw141_[int(10)] = D3_DC8(d_908_v0_)
        nw141_[int(11)] = d_909_v1_
        nw141_[int(12)] = default__.fm52(-402, globalState)
        nw141_[int(13)] = d_909_v1_
        nw141_[int(14)] = d_909_v1_
        nw141_[int(15)] = d_909_v1_
        nw141_[int(16)] = d_909_v1_
        nw141_[int(17)] = d_909_v1_
        nw141_[int(18)] = d_909_v1_
        nw141_[int(19)] = d_909_v1_
        nw141_[int(20)] = d_909_v1_
        nw141_[int(21)] = d_909_v1_
        nw141_[int(22)] = D3_DC8(d_908_v0_)
        nw141_[int(23)] = D3_DC8(d_908_v0_)
        nw141_[int(24)] = d_909_v1_
        nw141_[int(25)] = d_909_v1_
        d_910_v2_ = nw141_
        index134_ = default__.safeIndex(878, (d_910_v2_).length(0))
        (d_910_v2_)[index134_] = d_909_v1_
        pat_let_tv15_ = d_908_v0_
        index135_ = default__.safeIndex(878, (d_910_v2_).length(0))
        def iife74_(_pat_let18_0):
            def iife75_(d_911_dt__update__tmp_h0_):
                def iife76_(_pat_let19_0):
                    def iife77_(d_912_dt__update_hcf7_h0_):
                        return D3_DC8(d_912_dt__update_hcf7_h0_)
                    return iife77_(_pat_let19_0)
                return iife76_(pat_let_tv15_)
            return iife75_(_pat_let18_0)
        (d_910_v2_)[index135_] = iife74_(d_909_v1_)
        d_913_v3_: int
        d_913_v3_ = 187
        d_914_v4_: _dafny.Array
        nw142_ = _dafny.Array(None, 9)
        nw142_[int(0)] = default__.safeModuloInt(d_913_v3_, len((_dafny.SeqWithoutIsStrInference([d_913_v3_, d_913_v3_])).set(default__.safeIndex(408, len(_dafny.SeqWithoutIsStrInference([d_913_v3_, d_913_v3_]))), d_913_v3_)))
        nw142_[int(1)] = d_913_v3_
        nw142_[int(2)] = d_913_v3_
        nw142_[int(3)] = d_913_v3_
        nw142_[int(4)] = d_913_v3_
        nw142_[int(5)] = (d_913_v3_) + (577)
        nw142_[int(6)] = default__.fm3(d_913_v3_, (self).f4, globalState)
        nw142_[int(7)] = d_913_v3_
        nw142_[int(8)] = (0) - (len((self).f17))
        d_914_v4_ = nw142_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_914_v4_).length(0)):
            d_915_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_915_i0_)) and ((d_915_i0_) < ((d_914_v4_).length(0)))):
                (d_914_v4_)[(d_915_i0_)] = default__.safeModuloInt(d_915_i0_, (0) - ((d_913_v3_) - (d_913_v3_)))
        d_916_v5_: _dafny.Array
        nw143_ = _dafny.Array(None, 1)
        nw143_[int(0)] = (self).f4
        d_916_v5_ = nw143_
        d_917_v6_: _dafny.Map
        d_917_v6_ = _dafny.Map({p0: d_916_v5_})
        d_918_v7_: _dafny.Map
        d_918_v7_ = _dafny.Map({(0) - ((len(d_917_v6_)) * (d_913_v3_)): (0) - (d_913_v3_)})
        d_918_v7_ = (d_918_v7_).set(d_913_v3_, d_913_v3_)
        d_919_v8_: _dafny.MultiSet
        d_919_v8_ = _dafny.MultiSet([not((self).f4), p0])
        d_920_v11_: str
        d_920_v11_ = _dafny.CodePoint('n')
        d_921_v12_: _dafny.Array
        nw144_ = _dafny.Array(None, 13)
        nw144_[int(0)] = d_916_v5_
        nw144_[int(1)] = d_916_v5_
        nw144_[int(2)] = d_916_v5_
        nw144_[int(3)] = d_916_v5_
        nw144_[int(4)] = d_916_v5_
        nw144_[int(5)] = d_916_v5_
        nw144_[int(6)] = d_916_v5_
        nw144_[int(7)] = d_916_v5_
        nw144_[int(8)] = d_916_v5_
        nw144_[int(9)] = d_916_v5_
        nw144_[int(10)] = d_916_v5_
        nw144_[int(11)] = d_916_v5_
        nw144_[int(12)] = d_916_v5_
        d_921_v12_ = nw144_
        d_922_v14_: C1
        nw145_ = C1()
        def iife78_():
            def iife79_():
                coll39_ = _dafny.Map()
                compr_39_: int
                for compr_39_ in _dafny.IntegerRange(72, 693):
                    d_924_v10_: int = compr_39_
                    if ((72) <= (d_924_v10_)) and ((d_924_v10_) < (693)):
                        coll39_[(d_924_v10_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cii"))))] = (self).f4
                return _dafny.Map(coll39_)
            coll38_ = _dafny.Map()
            compr_38_: int
            for compr_38_ in _dafny.IntegerRange(321, 377):
                d_923_v9_: int = compr_38_
                if ((321) <= (d_923_v9_)) and ((d_923_v9_) < (377)):
                    coll38_[default__.safeDivisionInt(d_923_v9_, len(iife79_()
                    ))] = len((self).f17)
            return _dafny.Map(coll38_)
        def iife80_():
            coll40_ = _dafny.Map()
            compr_40_: int
            for compr_40_ in _dafny.IntegerRange(412, -676):
                d_925_v13_: int = compr_40_
                if ((412) <= (d_925_v13_)) and ((d_925_v13_) < (-676)):
                    coll40_[(d_925_v13_) * (d_913_v3_)] = p0
            return _dafny.Map(coll40_)
        nw145_.ctor__(len(default__.fm50(((d_919_v8_)[(self).f4] if ((self).f4) in (d_919_v8_) else d_913_v3_), len(iife78_()
        ), d_913_v3_, default__.fm0(d_920_v11_, globalState), globalState)), d_913_v3_, d_921_v12_, default__.fm3(len(iife80_()
        ), p0, globalState), p0)
        d_922_v14_ = nw145_
        d_926_v15_: _dafny.Set
        d_926_v15_ = _dafny.Set({p0, p0, p0})
        d_926_v15_ = default__.fm53(default__.safeModuloInt(d_913_v3_, d_913_v3_), (default__.fm54(d_922_v14_.f16, globalState)) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: p0}) for d_927_i1_ in range(default__.abs(644))])), globalState)
        d_928_i2_: int
        d_928_i2_ = 0
        with _dafny.label("5"):
            while not(((d_913_v3_ if p0 else (d_922_v14_).f15)) < ((d_922_v14_.f16) * ((d_922_v14_).f15))):
                with _dafny.c_label("5"):
                    if (d_928_i2_) >= (100):
                        raise _dafny.Break("5")
                    d_928_i2_ = (d_928_i2_) + (1)
                    index136_ = default__.safeIndex(549, (d_916_v5_).length(0))
                    (d_916_v5_)[index136_] = ((d_922_v14_).f15) < (21)
                    index137_ = default__.safeIndex(549, (d_916_v5_).length(0))
                    (d_916_v5_)[index137_] = default__.fm0(d_920_v11_, globalState)
                    (d_922_v14_).f16 = default__.fm3((self).fm7(globalState), p0, globalState)
                    index138_ = default__.safeIndex(549, (d_916_v5_).length(0))
                    (d_916_v5_)[index138_] = (self).f4
                    index139_ = default__.safeIndex(549, (d_916_v5_).length(0))
                    (d_916_v5_)[index139_] = (len((self).f17)) <= (d_913_v3_)
                    pass
            pass
        r0 = d_920_v11_
        d_929_v16_: D1
        d_929_v16_ = D1_DC3(p0)
        nw146_ = _dafny.Array(None, 12)
        nw146_[int(0)] = d_929_v16_
        nw146_[int(1)] = d_929_v16_
        nw146_[int(2)] = d_929_v16_
        nw146_[int(3)] = d_929_v16_
        nw146_[int(4)] = d_929_v16_
        nw146_[int(5)] = d_929_v16_
        nw146_[int(6)] = d_929_v16_
        nw146_[int(7)] = d_929_v16_
        nw146_[int(8)] = d_929_v16_
        def iife81_(_pat_let20_0):
            def iife82_(d_930_dt__update__tmp_h1_):
                def iife83_(_pat_let21_0):
                    def iife84_(d_931_dt__update_hcf2_h0_):
                        return D1_DC3(d_931_dt__update_hcf2_h0_)
                    return iife84_(_pat_let21_0)
                return iife83_((self).f4)
            return iife82_(_pat_let20_0)
        nw146_[int(9)] = iife81_(d_929_v16_)
        nw146_[int(10)] = default__.fm55(not(p0), (self).f4, 547, globalState)
        nw146_[int(11)] = D1_DC3((self).f4)
        r1 = nw146_
        return r0, r1

    def m12(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_932_v0_: _dafny.Array
        nw147_ = _dafny.Array(False, 16)
        d_932_v0_ = nw147_
        hi4_ = len(_dafny.SeqWithoutIsStrInference([d_932_v0_, d_932_v0_, d_932_v0_, d_932_v0_]))
        for d_933_i0_ in range(len((p1).intersection(_dafny.Set({(self).f4, (self).f4}))), hi4_):
            d_934_v1_: C0
            nw148_ = C0()
            nw148_.ctor__(271)
            d_934_v1_ = nw148_
            d_935_v2_: _dafny.Map
            d_935_v2_ = _dafny.Map({(self).f4: p0})
            d_936_v3_: _dafny.Seq
            d_936_v3_ = _dafny.SeqWithoutIsStrInference([d_933_i0_])
            d_937_v4_: D9
            d_937_v4_ = D9_DC19(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_938_i1_ in range(default__.abs(-645))]))
            d_939_v5_: _dafny.Array
            nw149_ = _dafny.Array(None, 21)
            nw149_[int(0)] = p0
            nw149_[int(1)] = d_933_i0_
            nw149_[int(2)] = d_933_i0_
            nw149_[int(3)] = 896
            nw149_[int(4)] = (d_934_v1_).f12
            nw149_[int(5)] = (_dafny.MultiSet(d_936_v3_)).cardinality
            nw149_[int(6)] = (d_934_v1_).f12
            nw149_[int(7)] = d_933_i0_
            nw149_[int(8)] = p0
            nw149_[int(9)] = d_933_i0_
            nw149_[int(10)] = d_933_i0_
            nw149_[int(11)] = p0
            nw149_[int(12)] = len((self).f17)
            nw149_[int(13)] = len((self).f17)
            nw149_[int(14)] = (d_934_v1_).f12
            nw149_[int(15)] = len((d_937_v4_).cf25)
            nw149_[int(16)] = -648
            nw149_[int(17)] = d_933_i0_
            nw149_[int(18)] = p0
            nw149_[int(19)] = d_933_i0_
            nw149_[int(20)] = p0
            d_939_v5_ = nw149_
            d_940_v6_: _dafny.Seq
            d_940_v6_ = _dafny.SeqWithoutIsStrInference([d_939_v5_])
            d_941_v7_: _dafny.Map
            d_941_v7_ = _dafny.Map({d_935_v2_: d_940_v6_})
            d_942_v8_: _dafny.Seq
            d_942_v8_ = _dafny.SeqWithoutIsStrInference([d_940_v6_])
            d_941_v7_ = (d_941_v7_).set(d_935_v2_, (d_942_v8_)[default__.safeIndex(p0, len(d_942_v8_))])
            d_943_v9_: _dafny.Array
            nw150_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_943_v9_ = nw150_
            index140_ = default__.safeIndex(170, (d_943_v9_).length(0))
            (d_943_v9_)[index140_] = d_932_v0_
            index141_ = default__.safeIndex(985, (d_932_v0_).length(0))
            (d_932_v0_)[index141_] = False
            d_944_v10_: str
            d_944_v10_ = _dafny.CodePoint('c')
            d_945_v11_: T1
            nw151_ = C1()
            nw151_.ctor__(d_933_i0_, p0, d_943_v9_, (self).fm7(globalState), p3)
            d_945_v11_ = nw151_
            d_946_v12_: _dafny.Map
            d_946_v12_ = _dafny.Map({d_945_v11_: (d_945_v11_).f4})
            index142_ = default__.safeIndex(964, (d_932_v0_).length(0))
            (d_932_v0_)[index142_] = ((d_946_v12_)[d_945_v11_] if (d_945_v11_) in (d_946_v12_) else (d_945_v11_).f4)
            d_947_v13_: _dafny.Seq
            d_947_v13_ = _dafny.SeqWithoutIsStrInference([(d_945_v11_).f4, (d_945_v11_).f4])
            index143_ = default__.safeIndex(170, (d_943_v9_).length(0))
            index144_ = default__.safeIndex(985, (d_932_v0_).length(0))
            index145_ = default__.safeIndex(964, (d_932_v0_).length(0))
            rhs142_ = d_932_v0_
            rhs143_ = (self).fm7(globalState)
            rhs144_ = not((not(False) if (d_945_v11_).f4 else (d_945_v11_).f4))
            rhs145_ = _dafny.CodePoint('k')
            rhs146_ = (not (p3) or ((self).f4) if (d_947_v13_)[default__.safeIndex(p0, len(d_947_v13_))] else p3)
            lhs84_ = d_943_v9_
            lhs85_ = default__.safeIndex(170, (d_943_v9_).length(0))
            lhs86_ = d_932_v0_
            lhs87_ = default__.safeIndex(985, (d_932_v0_).length(0))
            lhs88_ = d_932_v0_
            lhs89_ = default__.safeIndex(964, (d_932_v0_).length(0))
            lhs84_[lhs85_] = rhs142_
            r0 = rhs143_
            lhs86_[lhs87_] = rhs144_
            d_944_v10_ = rhs145_
            lhs88_[lhs89_] = rhs146_
            d_948_v14_: _dafny.Map
            d_948_v14_ = _dafny.Map({p0: (d_945_v11_).f4})
            d_949_v15_: _dafny.Seq
            d_949_v15_ = _dafny.SeqWithoutIsStrInference([d_948_v14_, (_dafny.Map({len(d_948_v14_): p3})) | (_dafny.Map({(d_934_v1_).f12: (d_945_v11_).f4})), d_948_v14_])
            d_949_v15_ = d_949_v15_
        d_950_v16_: str
        d_950_v16_ = _dafny.CodePoint('l')
        d_951_i2_: int
        d_951_i2_ = 0
        with _dafny.label("6"):
            while default__.fm0(d_950_v16_, globalState):
                with _dafny.c_label("6"):
                    if (d_951_i2_) >= (100):
                        raise _dafny.Break("6")
                    d_951_i2_ = (d_951_i2_) + (1)
                    d_952_v17_: D10
                    d_952_v17_ = D10_DC21(_dafny.Map({(self).f4: p0}))
                    d_953_v18_: _dafny.Seq
                    d_953_v18_ = _dafny.SeqWithoutIsStrInference([d_952_v17_, d_952_v17_])
                    d_954_v19_: _dafny.MultiSet
                    d_954_v19_ = _dafny.MultiSet([(390) - (p0), p0, ((0) - (len(d_953_v18_))) - (len(p2))])
                    r0 = ((d_954_v19_)[(0) - ((p0) * ((0) - (p0)))] if ((0) - ((p0) * ((0) - (p0)))) in (d_954_v19_) else p0)
                    d_955_v20_: _dafny.Seq
                    d_955_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulsxdir"))
                    d_955_v20_ = (d_955_v20_) + ((d_955_v20_ if not((self).f4) else (self).f17))
                    d_956_v21_: _dafny.MultiSet
                    d_956_v21_ = _dafny.MultiSet([p3])
                    d_957_v22_: _dafny.Array
                    nw152_ = _dafny.Array(_dafny.Array(None, 0), 8)
                    d_957_v22_ = nw152_
                    d_958_v23_: C1
                    nw153_ = C1()
                    nw153_.ctor__(p0, (0) - ((len(_dafny.SeqWithoutIsStrInference([d_950_v16_, d_950_v16_, d_950_v16_]))) + (((d_956_v21_)[True] if (True) in (d_956_v21_) else p0))), (d_957_v22_ if p3 else d_957_v22_), p0, (p0) != (493))
                    d_958_v23_ = nw153_
                    r0 = ((d_958_v23_).fm9((self).f4, (self).f4, (d_958_v23_).f15, globalState)) * ((d_958_v23_).f15)
                    pass
            pass
        r0 = (0) - (((p0) - (p0)) - (default__.safeModuloInt(p0, p0)))
        d_959_v24_: _dafny.Array
        nw154_ = _dafny.Array(int(0), 14)
        d_959_v24_ = nw154_
        d_960_v25_: _dafny.Map
        d_960_v25_ = _dafny.Map({((0) - ((self).fm7(globalState))) + (p0): d_959_v24_})
        d_959_v24_ = ((d_960_v25_)[718] if (718) in (d_960_v25_) else d_959_v24_)
        d_961_v26_: D2
        d_961_v26_ = D2_DC7((p0 if (self).f4 else p0))
        source17_ = d_961_v26_
        if source17_.is_DC6:
            d_962_v27_: C0
            nw155_ = C0()
            nw155_.ctor__(default__.safeDivisionInt((0) - (-298), p0))
            d_962_v27_ = nw155_
            d_963_v28_: bool
            d_963_v28_ = True
            d_963_v28_ = d_963_v28_
            r0 = p0
            r1 = (d_962_v27_).f12
        elif source17_.is_DC7:
            d_964___mcc_h0_ = source17_.cf6
            d_965_cf6_ = d_964___mcc_h0_
            r1 = p0
            d_966_v29_: bool
            d_966_v29_ = False
            d_966_v29_ = (p0) != (len((self).f17))
            d_967_v30_: _dafny.MultiSet
            d_967_v30_ = _dafny.MultiSet([(0) - (p0)])
            d_968_v31_: _dafny.Map
            d_968_v31_ = _dafny.Map({d_967_v30_: _dafny.SeqWithoutIsStrInference([d_950_v16_ for d_969_i3_ in range(default__.abs(-304))])})
            d_970_v32_: _dafny.Seq
            d_970_v32_ = _dafny.SeqWithoutIsStrInference([p3])
            d_971_v33_: _dafny.Seq
            d_971_v33_ = _dafny.SeqWithoutIsStrInference([default__.fm56((self).f4, d_965_cf6_, p3, globalState), d_970_v32_, d_970_v32_])
            d_972_v34_: _dafny.MultiSet
            d_972_v34_ = _dafny.MultiSet([d_970_v32_])
            d_973_v35_: bool
            d_974_v36_: int
            out27_: bool
            out28_: int
            out27_, out28_ = default__.m0(d_968_v31_, (_dafny.MultiSet(d_971_v33_)).issubset(d_972_v34_), d_967_v30_, (len(p2)) == (p0), globalState)
            d_973_v35_ = out27_
            d_974_v36_ = out28_
            d_975_v37_: _dafny.Array
            nw156_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_975_v37_ = nw156_
            d_976_v38_: C1
            nw157_ = C1()
            nw157_.ctor__(p0, 143, d_975_v37_, d_974_v36_, not(False))
            d_976_v38_ = nw157_
            d_977_v39_: _dafny.MultiSet
            d_977_v39_ = _dafny.MultiSet([(self).f4, p3])
            d_978_v40_: D6
            d_978_v40_ = D6_DC16(False, p3, (_dafny.SeqWithoutIsStrInference([d_965_cf6_])).set(default__.safeIndex(len(_dafny.Map({d_973_v35_: d_976_v38_})), len(_dafny.SeqWithoutIsStrInference([d_965_cf6_]))), (d_977_v39_).cardinality))
            d_979_v41_: _dafny.Map
            d_979_v41_ = _dafny.Map({default__.fm3(d_965_cf6_, (self).f4, globalState): d_978_v40_})
            d_980_v42_: _dafny.Set
            d_980_v42_ = _dafny.Set({d_932_v0_})
            d_979_v41_ = (d_979_v41_).set(len((d_980_v42_) | (d_980_v42_)), d_978_v40_)
        elif True:
            d_981___mcc_h1_ = source17_.cf5
            d_982_cf5_ = d_981___mcc_h1_
            d_983_v43_: _dafny.MultiSet
            d_983_v43_ = _dafny.MultiSet([True])
            d_984_v44_: _dafny.Seq
            d_984_v44_ = _dafny.SeqWithoutIsStrInference([p0, ((d_983_v43_)[p3] if (p3) in (d_983_v43_) else (0) - (p0)), p0])
            rhs147_ = d_932_v0_
            rhs148_ = p0
            rhs149_ = (d_984_v44_)[default__.safeIndex(default__.safeModuloInt(p0, p0), len(d_984_v44_))]
            d_982_cf5_ = rhs147_
            r0 = rhs148_
            r1 = rhs149_
            d_985_v45_: bool
            d_985_v45_ = True
            d_985_v45_ = not((self).f4)
            d_986_v46_: _dafny.Map
            d_986_v46_ = _dafny.Map({d_959_v24_: not((self).f4)})
            d_987_v47_: _dafny.Map
            d_987_v47_ = _dafny.Map({d_985_v45_: d_986_v46_})
            d_987_v47_ = (d_987_v47_).set(p3, (d_986_v46_ if d_985_v45_ else d_986_v46_))
            d_988_v48_: _dafny.Map
            d_988_v48_ = _dafny.Map({(((self).f17) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))).set(default__.safeIndex(298, len(((self).f17) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))))), d_950_v16_): (d_983_v43_).issubset(_dafny.MultiSet([p3]))})
            d_988_v48_ = (d_988_v48_).set((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpo")) if False else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlx"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlx")))), d_950_v16_)), (self).f4)
        if ((self).f4) and (p3):
            index146_ = default__.safeIndex(936, (d_932_v0_).length(0))
            (d_932_v0_)[index146_] = p3
            index147_ = default__.safeIndex(936, (d_932_v0_).length(0))
            (d_932_v0_)[index147_] = (self).f4
            d_989_v49_: _dafny.Seq
            d_989_v49_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_990_v50_: _dafny.Seq
            d_990_v50_ = _dafny.SeqWithoutIsStrInference([((p2)[(d_989_v49_)[default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference([(d_932_v0_)[default__.safeIndex(936, (d_932_v0_).length(0))], p3]))), len(d_989_v49_))]] if ((d_989_v49_)[default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference([(d_932_v0_)[default__.safeIndex(936, (d_932_v0_).length(0))], p3]))), len(d_989_v49_))]) in (p2) else p0), p0])
            d_990_v50_ = (d_989_v49_).set(default__.safeIndex((self).fm7(globalState), len(d_989_v49_)), p0)
            d_991_v51_: _dafny.Array
            nw158_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
            d_991_v51_ = nw158_
            index148_ = default__.safeIndex(893, (d_991_v51_).length(0))
            (d_991_v51_)[index148_] = (self).f17
            d_992_v52_: _dafny.Seq
            d_992_v52_ = _dafny.SeqWithoutIsStrInference([not((d_932_v0_)[default__.safeIndex(936, (d_932_v0_).length(0))])])
            index149_ = default__.safeIndex(893, (d_991_v51_).length(0))
            (d_991_v51_)[index149_] = (((self).f17 if (d_992_v52_)[default__.safeIndex(p0, len(d_992_v52_))] else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dehlggigm")))) + ((self).f17)
            d_993_v53_: _dafny.Array
            nw159_ = _dafny.Array(None, 15)
            nw159_[int(0)] = d_950_v16_
            nw159_[int(1)] = d_950_v16_
            nw159_[int(2)] = d_950_v16_
            nw159_[int(3)] = d_950_v16_
            nw159_[int(4)] = _dafny.CodePoint('e')
            nw159_[int(5)] = _dafny.CodePoint('s')
            nw159_[int(6)] = d_950_v16_
            nw159_[int(7)] = d_950_v16_
            nw159_[int(8)] = ((d_991_v51_)[default__.safeIndex(893, (d_991_v51_).length(0))])[default__.safeIndex(p0, len((d_991_v51_)[default__.safeIndex(893, (d_991_v51_).length(0))]))]
            nw159_[int(9)] = _dafny.CodePoint('g')
            nw159_[int(10)] = default__.fm1(p0, p3, globalState)
            nw159_[int(11)] = d_950_v16_
            nw159_[int(12)] = d_950_v16_
            nw159_[int(13)] = default__.fm1(p0, (d_932_v0_)[default__.safeIndex(936, (d_932_v0_).length(0))], globalState)
            nw159_[int(14)] = d_950_v16_
            d_993_v53_ = nw159_
            d_993_v53_ = d_993_v53_
            d_994_v54_: _dafny.Array
            def lambda29_(d_995_p3_):
                def lambda30_(d_996_i4_):
                    return d_995_p3_

                return lambda30_

            init19_ = lambda29_(p3)
            nw160_ = _dafny.Array(None, 2)
            for i0_19_ in range(nw160_.length(0)):
                nw160_[i0_19_] = init19_(i0_19_)
            d_994_v54_ = nw160_
            d_997_v55_: _dafny.Map
            d_997_v55_ = _dafny.Map({default__.safeDivisionInt(p0, 135): d_994_v54_})
            d_997_v55_ = _dafny.Map({p0: d_994_v54_})
        elif True:
            d_998_v56_: bool
            d_998_v56_ = False
            d_999_v57_: _dafny.Seq
            d_999_v57_ = _dafny.SeqWithoutIsStrInference([(self).f4])
            d_1000_v58_: _dafny.Map
            d_1000_v58_ = _dafny.Map({p0: False})
            d_1001_v59_: _dafny.Seq
            d_1001_v59_ = _dafny.SeqWithoutIsStrInference([len(d_1000_v58_)])
            d_1002_v60_: _dafny.Map
            d_1002_v60_ = _dafny.Map({d_998_v56_: len(d_1001_v59_)})
            d_1003_v61_: _dafny.MultiSet
            d_1003_v61_ = _dafny.MultiSet([(d_999_v57_)[default__.safeIndex(((d_1002_v60_)[p3] if (p3) in (d_1002_v60_) else p0), len(d_999_v57_))], p3])
            d_1004_v62_: _dafny.Array
            nw161_ = _dafny.Array(None, 2)
            nw161_[int(0)] = d_1003_v61_
            nw161_[int(1)] = d_1003_v61_
            d_1004_v62_ = nw161_
            d_1005_v63_: _dafny.Array
            nw162_ = _dafny.Array(None, 7)
            nw162_[int(0)] = d_1004_v62_
            nw162_[int(1)] = d_1004_v62_
            nw162_[int(2)] = d_1004_v62_
            nw162_[int(3)] = d_1004_v62_
            nw162_[int(4)] = d_1004_v62_
            nw162_[int(5)] = (d_1004_v62_ if (self).f4 else d_1004_v62_)
            nw162_[int(6)] = d_1004_v62_
            d_1005_v63_ = nw162_
            index150_ = default__.safeIndex(906, (d_1005_v63_).length(0))
            (d_1005_v63_)[index150_] = d_1004_v62_
            index151_ = default__.safeIndex(906, (d_1005_v63_).length(0))
            rhs150_ = (0) - (p0)
            rhs151_ = d_932_v0_
            rhs152_ = not((self).f4)
            rhs153_ = d_1004_v62_
            lhs90_ = d_1005_v63_
            lhs91_ = default__.safeIndex(906, (d_1005_v63_).length(0))
            r1 = rhs150_
            d_932_v0_ = rhs151_
            d_998_v56_ = rhs152_
            lhs90_[lhs91_] = rhs153_
            d_1006_v64_: _dafny.Array
            nw163_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_1006_v64_ = nw163_
            d_1007_v65_: C1
            nw164_ = C1()
            nw164_.ctor__((d_1001_v59_)[default__.safeIndex(p0, len(d_1001_v59_))], default__.safeModuloInt(p0, p0), d_1006_v64_, default__.safeDivisionInt((_dafny.MultiSet([len(d_999_v57_), 9])).cardinality, p0), p3)
            d_1007_v65_ = nw164_
            d_998_v56_ = d_998_v56_
            d_998_v56_ = ((d_1001_v59_).set(default__.safeIndex((d_1007_v65_).f15, len(d_1001_v59_)), p0)) != (((d_1001_v59_) + (default__.fm57((d_1007_v65_).f15, (0) - (d_1007_v65_.f16), globalState))).set(default__.safeIndex(d_1007_v65_.f16, len((d_1001_v59_) + (default__.fm57((d_1007_v65_).f15, (0) - (d_1007_v65_.f16), globalState)))), d_1007_v65_.f16))
            d_998_v56_ = (d_1007_v65_.f16) <= (d_1007_v65_.f16)
        r0 = -259
        r1 = (0) - (p0)
        return r0, r1

    @property
    def f17(self):
        return self._f17

class C3(T1, T0):
    def  __init__(self):
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self._f6: int = int(0)
        self._f4: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f5(self):
        return self._f5
    @f5.setter
    def f5(self, value):
        self._f5 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f4(self):
        return self._f4
    def ctor__(self, f5, f6, f4):
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f4 = f4

    def fm8(self, p0, p1, p2, p3, globalState):
        return not (((self).f6) != ((0) - ((self).f6))) or ((self).f4)

    def fm9(self, p0, p1, p2, globalState):
        return 476

    def fm6(self, p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([(self).f6, (self).f6])) - (_dafny.MultiSet([(self).f6, (self).f6]))

    def fm7(self, globalState):
        return (self).f6

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_1008_v0_: _dafny.Array
        def lambda31_(d_1009_p0_):
            def lambda32_(d_1010_i1_):
                return default__.safeDivisionInt(d_1010_i1_, d_1009_p0_)

            return lambda32_

        init20_ = lambda31_(p0)
        nw165_ = _dafny.Array(None, 22)
        for i0_20_ in range(nw165_.length(0)):
            nw165_[i0_20_] = init20_(i0_20_)
        d_1008_v0_ = nw165_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1008_v0_).length(0)):
            d_1011_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_1011_i0_)) and ((d_1011_i0_) < ((d_1008_v0_).length(0)))):
                (d_1008_v0_)[(d_1011_i0_)] = (d_1011_i0_) - (p0)
        d_1012_v1_: _dafny.Seq
        d_1012_v1_ = _dafny.SeqWithoutIsStrInference([False])
        d_1013_v2_: _dafny.MultiSet
        d_1013_v2_ = _dafny.MultiSet([d_1012_v1_])
        d_1014_v3_: D6
        d_1014_v3_ = D6_DC15(d_1013_v2_)
        d_1014_v3_ = default__.fm33((self).f4, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqdyv")))), (False if not((self).f4) else not((self).f4)), globalState)
        d_1015_v4_: _dafny.Seq
        d_1015_v4_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1016_v5_: _dafny.Set
        d_1016_v5_ = _dafny.Set({p0, len(d_1015_v4_)})
        hi5_ = (p0) - (len(d_1016_v5_))
        for d_1017_i2_ in range(p0, hi5_):
            d_1018_v6_: _dafny.Seq
            d_1018_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sj"))
            d_1019_v7_: _dafny.Map
            d_1019_v7_ = _dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0 for d_1020_i3_ in range(default__.abs(871))])): d_1018_v6_})
            d_1021_v8_: _dafny.MultiSet
            d_1021_v8_ = _dafny.MultiSet([len(d_1018_v6_)])
            d_1022_v9_: bool
            d_1023_v10_: int
            out29_: bool
            out30_: int
            out29_, out30_ = default__.m0((default__.fm34(d_1017_i2_, p0, not((self).f4), (self).f4, globalState)) | (d_1019_v7_), (self).f4, d_1021_v8_, (self).f4, globalState)
            d_1022_v9_ = out29_
            d_1023_v10_ = out30_
            d_1023_v10_ = ((self).f6) * ((p0 if not((self).f4) else ((d_1021_v8_)[d_1017_i2_] if (d_1017_i2_) in (d_1021_v8_) else p0)))
            d_1024_v11_: C0
            nw166_ = C0()
            nw166_.ctor__(887)
            d_1024_v11_ = nw166_
            d_1025_v12_: _dafny.Set
            d_1025_v12_ = default__.fm35(d_1017_i2_, globalState)
            source18_ = d_1025_v12_
            d_1026___mcc_h0_ = source18_
            d_1027_cf23_ = d_1026___mcc_h0_
            d_1028_v13_: _dafny.Map
            d_1028_v13_ = _dafny.Map({(d_1024_v11_).f12: _dafny.SeqWithoutIsStrInference([p0 for d_1029_i5_ in range(default__.abs(832))])})
            d_1030_v14_: _dafny.Array
            nw167_ = _dafny.Array(None, 28)
            nw167_[int(0)] = _dafny.SeqWithoutIsStrInference([(d_1024_v11_).f12 for d_1031_i4_ in range(default__.abs(870))])
            nw167_[int(1)] = (d_1015_v4_) + (d_1015_v4_)
            nw167_[int(2)] = (d_1015_v4_) + (d_1015_v4_)
            nw167_[int(3)] = (d_1015_v4_) + ((d_1015_v4_).set(default__.safeIndex(523, len(d_1015_v4_)), d_1023_v10_))
            nw167_[int(4)] = (((d_1028_v13_)[(d_1024_v11_).f12] if ((d_1024_v11_).f12) in (d_1028_v13_) else d_1015_v4_)) + (d_1015_v4_)
            nw167_[int(5)] = d_1015_v4_
            nw167_[int(6)] = (d_1015_v4_) + (_dafny.SeqWithoutIsStrInference([d_1017_i2_, (self).fm7(globalState)]))
            nw167_[int(7)] = d_1015_v4_
            nw167_[int(8)] = d_1015_v4_
            nw167_[int(9)] = (d_1015_v4_ if (self).f4 else _dafny.SeqWithoutIsStrInference([d_1017_i2_]))
            nw167_[int(10)] = d_1015_v4_
            nw167_[int(11)] = d_1015_v4_
            nw167_[int(12)] = d_1015_v4_
            nw167_[int(13)] = d_1015_v4_
            nw167_[int(14)] = d_1015_v4_
            nw167_[int(15)] = _dafny.SeqWithoutIsStrInference([d_1023_v10_ for d_1032_i6_ in range(default__.abs(386))])
            nw167_[int(16)] = _dafny.SeqWithoutIsStrInference([(0) - ((d_1024_v11_).f12)])
            nw167_[int(17)] = d_1015_v4_
            nw167_[int(18)] = default__.fm36(True, globalState)
            nw167_[int(19)] = _dafny.SeqWithoutIsStrInference([(d_1021_v8_).cardinality, (0) - (d_1023_v10_)])
            nw167_[int(20)] = d_1015_v4_
            nw167_[int(21)] = _dafny.SeqWithoutIsStrInference([d_1017_i2_, (d_1024_v11_).f12, len(d_1012_v1_), d_1023_v10_])
            nw167_[int(22)] = d_1015_v4_
            nw167_[int(23)] = (d_1015_v4_).set(default__.safeIndex((0) - (d_1017_i2_), len(d_1015_v4_)), len(d_1018_v6_))
            nw167_[int(24)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f4, False])) for d_1033_i7_ in range(default__.abs(633))])
            nw167_[int(25)] = (_dafny.SeqWithoutIsStrInference([p0 for d_1034_i8_ in range(default__.abs(528))])) + (_dafny.SeqWithoutIsStrInference([-466]))
            nw167_[int(26)] = default__.fm36((self).fm8((_dafny.MultiSet([d_1022_v9_, (self).f4, d_1022_v9_])).cardinality, False, d_1018_v6_, d_1015_v4_, globalState), globalState)
            nw167_[int(27)] = d_1015_v4_
            d_1030_v14_ = nw167_
            index152_ = default__.safeIndex(161, (d_1030_v14_).length(0))
            (d_1030_v14_)[index152_] = d_1015_v4_
            index153_ = default__.safeIndex(161, (d_1030_v14_).length(0))
            (d_1030_v14_)[index153_] = d_1015_v4_
            d_1023_v10_ = (d_1024_v11_).f12
            d_1035_v15_: _dafny.Seq
            d_1035_v15_ = _dafny.SeqWithoutIsStrInference([d_1018_v6_])
            d_1036_v16_: _dafny.Seq
            d_1036_v16_ = _dafny.SeqWithoutIsStrInference([d_1018_v6_, (d_1035_v15_)[default__.safeIndex((self).f6, len(d_1035_v15_))], d_1018_v6_])
            d_1037_v17_: D3
            d_1037_v17_ = D3_DC10(len((d_1036_v16_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1023_v10_])), len(d_1036_v16_))]))
            d_1037_v17_ = d_1037_v17_
            d_1022_v9_ = d_1022_v9_
        d_1038_v18_: _dafny.Array
        nw168_ = _dafny.Array(_dafny.Seq({}), 23)
        d_1038_v18_ = nw168_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1038_v18_).length(0)):
            d_1039_i9_: int = guard_loop_5_
            if (True) and (((0) <= (d_1039_i9_)) and ((d_1039_i9_) < ((d_1038_v18_).length(0)))):
                (d_1038_v18_)[(d_1039_i9_)] = d_1012_v1_
        d_1040_v19_: bool
        d_1040_v19_ = True
        d_1040_v19_ = (self).f4
        source19_ = D0_DC2(-988)
        if source19_.is_DC1:
            d_1041_v20_: _dafny.Map
            d_1041_v20_ = _dafny.Map({default__.fm3(p0, d_1040_v19_, globalState): p0})
            d_1042_v21_: _dafny.Seq
            d_1042_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtke"))
            d_1041_v20_ = (d_1041_v20_).set(len(d_1042_v21_), default__.safeDivisionInt(len(d_1016_v5_), (0) - (p0)))
            d_1040_v19_ = (self).f4
            d_1043_v22_: int
            d_1043_v22_ = -943
            d_1044_v23_: _dafny.Set
            d_1044_v23_ = _dafny.Set({d_1008_v0_, d_1008_v0_})
            d_1043_v22_ = len(d_1044_v23_)
            d_1040_v19_ = (self).f4
        elif source19_.is_DC2:
            d_1045___mcc_h1_ = source19_.cf1
            d_1046_cf1_ = d_1045___mcc_h1_
            d_1047_v24_: _dafny.Map
            d_1047_v24_ = _dafny.Map({d_1016_v5_: p0})
            d_1046_cf1_ = default__.safeDivisionInt(((d_1047_v24_)[d_1016_v5_] if (d_1016_v5_) in (d_1047_v24_) else d_1046_cf1_), (self).f6)
            d_1048_v25_: _dafny.Map
            d_1048_v25_ = _dafny.Map({(self).f4: (self).f6})
            d_1049_v26_: D10
            d_1049_v26_ = D10_DC21((d_1048_v25_ if not((self).f4) else d_1048_v25_))
            source20_ = d_1049_v26_
            if source20_.is_DC22:
                d_1050___mcc_h3_ = source20_.cf30
                d_1051___mcc_h4_ = source20_.cf31
                d_1052___mcc_h5_ = source20_.cf32
                d_1053___mcc_h6_ = source20_.cf33
                d_1054___mcc_h7_ = source20_.cf34
                d_1055_cf34_ = d_1054___mcc_h7_
                d_1056_cf33_ = d_1053___mcc_h6_
                d_1057_cf32_ = d_1052___mcc_h5_
                d_1058_cf31_ = d_1051___mcc_h4_
                d_1059_cf30_ = d_1050___mcc_h3_
                d_1060_v27_: str
                d_1060_v27_ = _dafny.CodePoint('t')
                d_1061_v28_: bool
                out31_: bool
                out31_ = (self).m9(default__.fm0(d_1060_v27_, globalState), globalState)
                d_1061_v28_ = out31_
                d_1046_cf1_ = ((d_1046_cf1_) * (d_1046_cf1_)) + (-560)
                d_1062_v29_: _dafny.Array
                nw169_ = _dafny.Array(_dafny.Map({}), 23)
                d_1062_v29_ = nw169_
                d_1063_v30_: _dafny.Map
                d_1063_v30_ = _dafny.Map({d_1062_v29_: d_1040_v19_})
                d_1064_v31_: D11
                d_1064_v31_ = D11_DC23(d_1062_v29_)
                d_1063_v30_ = (d_1063_v30_).set((d_1064_v31_).cf35, (self).f4)
                d_1065_v32_: _dafny.MultiSet
                d_1065_v32_ = _dafny.MultiSet([p0])
                d_1040_v19_ = ((_dafny.SeqWithoutIsStrInference([(self).f6])) <= (d_1015_v4_) if d_1040_v19_ else (d_1065_v32_).issubset(d_1065_v32_))
            elif True:
                d_1066___mcc_h8_ = source20_.cf29
                d_1067_cf29_ = d_1066___mcc_h8_
                d_1013_v2_ = d_1013_v2_
                index154_ = default__.safeIndex(483, (d_1008_v0_).length(0))
                (d_1008_v0_)[index154_] = (self).f6
                d_1068_v33_: _dafny.Set
                d_1068_v33_ = _dafny.Set({d_1040_v19_})
                d_1069_v34_: _dafny.Map
                d_1069_v34_ = _dafny.Map({p0: d_1068_v33_})
                d_1070_v35_: str
                d_1070_v35_ = _dafny.CodePoint('k')
                d_1071_v36_: C0
                nw170_ = C0()
                nw170_.ctor__((self).f6)
                d_1071_v36_ = nw170_
                d_1072_v37_: _dafny.MultiSet
                d_1072_v37_ = _dafny.MultiSet([p0, (0) - ((p0) * ((self).f6)), (d_1046_cf1_) * (386), p0, (self).fm9((self).f4, (self).f4, (d_1071_v36_).f12, globalState)])
                d_1073_v38_: _dafny.Seq
                d_1073_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcnfyfu"))
                d_1074_v39_: _dafny.MultiSet
                d_1074_v39_ = _dafny.MultiSet([d_1072_v37_])
                index155_ = default__.safeIndex(483, (d_1008_v0_).length(0))
                rhs154_ = len(_dafny.Map({((d_1069_v34_)[(self).f6] if ((self).f6) in (d_1069_v34_) else d_1068_v33_): _dafny.Map({d_1070_v35_: d_1071_v36_})}))
                rhs155_ = ((d_1072_v37_)[len(d_1073_v38_)] if (len(d_1073_v38_)) in (d_1072_v37_) else d_1046_cf1_)
                rhs156_ = ((d_1074_v39_)[(d_1072_v37_).intersection(d_1072_v37_)] if ((d_1072_v37_).intersection(d_1072_v37_)) in (d_1074_v39_) else p0)
                rhs157_ = d_1046_cf1_
                lhs92_ = d_1008_v0_
                lhs93_ = default__.safeIndex(483, (d_1008_v0_).length(0))
                d_1046_cf1_ = rhs154_
                d_1046_cf1_ = rhs155_
                d_1046_cf1_ = rhs156_
                lhs92_[lhs93_] = rhs157_
                d_1075_v40_: _dafny.Map
                d_1075_v40_ = _dafny.Map({default__.safeDivisionInt(-194, 341): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mou"))})
                d_1075_v40_ = (d_1075_v40_).set((d_1008_v0_)[default__.safeIndex(483, (d_1008_v0_).length(0))], _dafny.SeqWithoutIsStrInference([d_1070_v35_ for d_1076_i10_ in range(default__.abs(70))]))
                d_1077_v41_: _dafny.Map
                d_1077_v41_ = _dafny.Map({True: d_1070_v35_})
                d_1078_v42_: _dafny.Map
                d_1078_v42_ = _dafny.Map({d_1040_v19_: (self).f4})
                d_1079_v43_: _dafny.Map
                d_1079_v43_ = _dafny.Map({d_1046_cf1_: len(d_1078_v42_)})
                d_1080_v44_: _dafny.Map
                d_1080_v44_ = _dafny.Map({(self).f6: False})
                d_1077_v41_ = (d_1077_v41_).set(((d_1008_v0_)[default__.safeIndex(483, (d_1008_v0_).length(0))]) <= (((d_1079_v43_)[d_1046_cf1_] if (d_1046_cf1_) in (d_1079_v43_) else len(d_1080_v44_))), d_1070_v35_)
            d_1040_v19_ = d_1040_v19_
            source21_ = d_1014_v3_
            if source21_.is_DC16:
                d_1081___mcc_h9_ = source21_.cf20
                d_1082___mcc_h10_ = source21_.cf21
                d_1083___mcc_h11_ = source21_.cf22
                d_1084_cf22_ = d_1083___mcc_h11_
                d_1085_cf21_ = d_1082___mcc_h10_
                d_1086_cf20_ = d_1081___mcc_h9_
                d_1087_v45_: T1
                nw171_ = C1()
                nw171_.ctor__((0) - (d_1046_cf1_), (self).f6, self.f5, 490, False)
                d_1087_v45_ = nw171_
                d_1088_v46_: _dafny.Map
                d_1088_v46_ = _dafny.Map({d_1087_v45_: (d_1087_v45_).f4})
                d_1089_v47_: _dafny.Map
                d_1089_v47_ = _dafny.Map({default__.fm1(p0, d_1085_cf21_, globalState): d_1088_v46_})
                d_1089_v47_ = d_1089_v47_
                d_1090_v48_: str
                d_1090_v48_ = _dafny.CodePoint('r')
                d_1091_v49_: _dafny.Map
                d_1091_v49_ = _dafny.Map({d_1090_v48_: d_1090_v48_})
                d_1092_v50_: _dafny.MultiSet
                d_1092_v50_ = _dafny.MultiSet([d_1091_v49_])
                d_1093_v51_: _dafny.Set
                d_1093_v51_ = _dafny.Set({d_1040_v19_, False})
                d_1094_v52_: _dafny.MultiSet
                d_1094_v52_ = _dafny.MultiSet([(self).f6])
                d_1086_cf20_ = ((d_1094_v52_) | (d_1094_v52_)).issubset(_dafny.MultiSet([(d_1092_v50_).cardinality, len(d_1093_v51_), p0, d_1046_cf1_]))
                d_1095_v53_: _dafny.MultiSet
                d_1095_v53_ = _dafny.MultiSet([d_1040_v19_])
                d_1012_v1_ = (D4_DC13((d_1087_v45_).f6, d_1040_v19_, (d_1095_v53_).cardinality, d_1012_v1_)).cf17
                d_1096_v54_: _dafny.Seq
                d_1096_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfi"))
                d_1097_v55_: _dafny.Map
                d_1097_v55_ = _dafny.Map({d_1094_v52_: d_1096_v54_})
                d_1098_v56_: bool
                d_1099_v57_: int
                out32_: bool
                out33_: int
                out32_, out33_ = default__.m0(d_1097_v55_, (d_1084_cf22_) <= (d_1015_v4_), d_1094_v52_, (d_1087_v45_).f4, globalState)
                d_1098_v56_ = out32_
                d_1099_v57_ = out33_
            elif True:
                d_1100___mcc_h12_ = source21_.cf19
                d_1101_cf19_ = d_1100___mcc_h12_
                d_1046_cf1_ = (0) - (default__.safeModuloInt((self).fm9(d_1040_v19_, True, (self).f6, globalState), ((d_1048_v25_)[(self).f4] if ((self).f4) in (d_1048_v25_) else p0)))
                d_1102_v58_: _dafny.Array
                nw172_ = _dafny.Array(None, 2)
                nw172_[int(0)] = d_1040_v19_
                nw172_[int(1)] = (self).f4
                d_1102_v58_ = nw172_
                d_1103_v59_: _dafny.Map
                d_1103_v59_ = _dafny.Map({d_1102_v58_: d_1046_cf1_})
                d_1104_v60_: _dafny.Array
                d_1104_v60_ = self.f5
                d_1105_v61_: C1
                nw173_ = C1()
                nw173_.ctor__((self).f6, (self).f6, (d_1104_v60_), (self).f6, d_1040_v19_)
                d_1105_v61_ = nw173_
                d_1106_v62_: _dafny.Map
                d_1106_v62_ = _dafny.Map({d_1102_v58_: d_1105_v61_})
                d_1107_v63_: _dafny.Map
                d_1107_v63_ = _dafny.Map({d_1046_cf1_: len(d_1106_v62_)})
                d_1108_v64_: T1
                nw174_ = C1()
                nw174_.ctor__(len(d_1107_v63_), (self).f6, self.f5, (0) - (len(d_1015_v4_)), (self).f4)
                d_1108_v64_ = nw174_
                d_1109_v65_: _dafny.Map
                d_1109_v65_ = _dafny.Map({d_1108_v64_: d_1102_v58_})
                d_1103_v59_ = (d_1103_v59_).set(((d_1109_v65_)[d_1108_v64_] if (d_1108_v64_) in (d_1109_v65_) else d_1102_v58_), d_1105_v61_.f16)
                d_1110_v66_: _dafny.MultiSet
                d_1110_v66_ = _dafny.MultiSet([False, not(d_1040_v19_)])
                d_1111_v67_: D12
                d_1111_v67_ = D12_DC28(p0, (d_1110_v66_).cardinality, False, (d_1040_v19_) or (d_1040_v19_), (self).fm9(d_1040_v19_, d_1040_v19_, d_1105_v61_.f16, globalState))
                d_1111_v67_ = d_1111_v67_
                d_1112_v68_: _dafny.Map
                d_1112_v68_ = _dafny.Map({d_1105_v61_.f16: _dafny.CodePoint('m')})
                (d_1105_v61_).f16 = default__.safeModuloInt(p0, len(d_1112_v68_))
        elif True:
            d_1113___mcc_h2_ = source19_.cf0
            d_1114_cf0_ = d_1113___mcc_h2_
            d_1115_v69_: _dafny.Seq
            d_1115_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlh"))
            d_1116_v70_: _dafny.Map
            d_1116_v70_ = _dafny.Map({_dafny.MultiSet(default__.fm39(d_1115_v69_, D12_DC27(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f4])])), default__.fm0(_dafny.CodePoint('r'), globalState), globalState)): d_1115_v69_})
            d_1117_v71_: bool
            d_1118_v72_: int
            out34_: bool
            out35_: int
            out34_, out35_ = default__.m0((d_1116_v70_) | (d_1116_v70_), (self).f4, _dafny.MultiSet((d_1015_v4_) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_1115_v69_: _dafny.CodePoint('i')})) for d_1119_i11_ in range(default__.abs(68))]))), not(d_1040_v19_), globalState)
            d_1117_v71_ = out34_
            d_1118_v72_ = out35_
            if False:
                d_1117_v71_ = d_1040_v19_
                d_1008_v0_ = d_1008_v0_
                d_1120_v74_: C1
                nw175_ = C1()
                def iife85_():
                    coll41_ = _dafny.Set()
                    compr_41_: int
                    for compr_41_ in _dafny.IntegerRange(779, -489):
                        d_1121_v73_: int = compr_41_
                        if ((779) <= (d_1121_v73_)) and ((d_1121_v73_) < (-489)):
                            coll41_ = coll41_.union(_dafny.Set([(d_1121_v73_) - (p0)]))
                    return _dafny.Set(coll41_)
                nw175_.ctor__((0) - (d_1118_v72_), (d_1015_v4_)[default__.safeIndex((self).f6, len(d_1015_v4_))], self.f5, len((d_1016_v5_) | (iife85_()
                )), (d_1118_v72_) > (d_1118_v72_))
                d_1120_v74_ = nw175_
                d_1040_v19_ = (self).f4
                d_1122_v75_: _dafny.Array
                nw176_ = _dafny.Array(False, 7)
                d_1122_v75_ = nw176_
                d_1122_v75_ = d_1122_v75_
            elif True:
                d_1114_cf0_ = d_1118_v72_
                d_1123_v76_: _dafny.Map
                d_1123_v76_ = _dafny.Map({len(d_1012_v1_): (self).f6})
                d_1124_v77_: _dafny.MultiSet
                d_1124_v77_ = _dafny.MultiSet([226])
                d_1125_v78_: _dafny.Seq
                d_1125_v78_ = _dafny.SeqWithoutIsStrInference([d_1124_v77_])
                d_1117_v71_ = (d_1012_v1_)[default__.safeIndex(default__.safeDivisionInt(len(d_1123_v76_), ((d_1125_v78_)[default__.safeIndex(9, len(d_1125_v78_))]).cardinality), len(d_1012_v1_))]
                d_1126_v79_: _dafny.Map
                d_1126_v79_ = _dafny.Map({(d_1114_cf0_) < (d_1114_cf0_): d_1117_v71_})
                d_1040_v19_ = ((d_1126_v79_)[(self).f4] if ((self).f4) in (d_1126_v79_) else d_1040_v19_)
                d_1127_v80_: str
                d_1127_v80_ = _dafny.CodePoint('v')
                d_1127_v80_ = d_1127_v80_
                d_1128_v81_: _dafny.Array
                nw177_ = _dafny.Array(False, 24)
                d_1128_v81_ = nw177_
                d_1128_v81_ = (D2_DC5(d_1128_v81_)).cf5
            d_1129_v82_: str
            d_1129_v82_ = _dafny.CodePoint('l')
            d_1129_v82_ = d_1129_v82_
            d_1130_v85_: _dafny.Set
            d_1130_v85_ = _dafny.Set({(self).f4, d_1117_v71_, default__.fm0(d_1129_v82_, globalState)})
            d_1131_v86_: _dafny.Array
            nw178_ = _dafny.Array(None, 6)
            nw178_[int(0)] = d_1016_v5_
            def iife86_():
                coll42_ = _dafny.Set()
                compr_42_: int
                for compr_42_ in (d_1015_v4_).Elements:
                    d_1132_v83_: int = compr_42_
                    if (d_1132_v83_) in (d_1015_v4_):
                        coll42_ = coll42_.union(_dafny.Set([default__.safeDivisionInt(d_1132_v83_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1133_i12_ in range(default__.abs(-932))])))]))
                return _dafny.Set(coll42_)
            nw178_[int(1)] = iife86_()
            
            nw178_[int(2)] = d_1016_v5_
            def iife87_():
                coll43_ = _dafny.Set()
                compr_43_: int
                for compr_43_ in (d_1015_v4_).Elements:
                    d_1134_v84_: int = compr_43_
                    if (d_1134_v84_) in (d_1015_v4_):
                        coll43_ = coll43_.union(_dafny.Set([default__.safeDivisionInt(d_1134_v84_, (0) - (-223))]))
                return _dafny.Set(coll43_)
            nw178_[int(3)] = (d_1016_v5_ if d_1117_v71_ else iife87_()
            )
            nw178_[int(4)] = d_1016_v5_
            nw178_[int(5)] = default__.fm35(len(d_1130_v85_), globalState)
            d_1131_v86_ = nw178_
            index156_ = default__.safeIndex(869, (d_1131_v86_).length(0))
            (d_1131_v86_)[index156_] = d_1016_v5_
            d_1135_v87_: _dafny.MultiSet
            d_1135_v87_ = _dafny.MultiSet([p0])
            index157_ = default__.safeIndex(869, (d_1131_v86_).length(0))
            (d_1131_v86_)[index157_] = (d_1016_v5_).intersection(_dafny.Set({((d_1135_v87_)[(self).f6] if ((self).f6) in (d_1135_v87_) else (self).f6), p0, (self).f6, d_1114_cf0_, p0}))
        d_1136_v88_: _dafny.Map
        d_1136_v88_ = _dafny.Map({d_1040_v19_: (self).f6})
        r0 = d_1136_v88_
        return r0

    def m2(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_1137_v0_: _dafny.MultiSet
        d_1137_v0_ = _dafny.MultiSet([(self).f6])
        d_1138_v1_: _dafny.Array
        nw179_ = _dafny.Array(None, 3)
        nw179_[int(0)] = (0) - (((d_1137_v0_)[(self).f6] if ((self).f6) in (d_1137_v0_) else (self).f6))
        nw179_[int(1)] = (self).f6
        nw179_[int(2)] = (0) - (((self).f6) * ((self).fm7(globalState)))
        d_1138_v1_ = nw179_
        index158_ = default__.safeIndex(699, (d_1138_v1_).length(0))
        (d_1138_v1_)[index158_] = (self).f6
        index159_ = default__.safeIndex(699, (d_1138_v1_).length(0))
        (d_1138_v1_)[index159_] = (self).f6
        d_1139_v2_: _dafny.Map
        d_1139_v2_ = _dafny.Map({(d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))]: d_1138_v1_})
        d_1140_v3_: _dafny.Map
        d_1140_v3_ = _dafny.Map({442: ((d_1139_v2_)[(self).f6] if ((self).f6) in (d_1139_v2_) else d_1138_v1_)})
        d_1141_v4_: _dafny.Seq
        d_1141_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivk"))
        d_1142_v5_: _dafny.Seq
        d_1142_v5_ = _dafny.SeqWithoutIsStrInference([d_1138_v1_, d_1138_v1_, d_1138_v1_, d_1138_v1_])
        d_1143_v6_: _dafny.Array
        nw180_ = _dafny.Array(None, 24)
        nw180_[int(0)] = d_1138_v1_
        nw180_[int(1)] = d_1138_v1_
        nw180_[int(2)] = d_1138_v1_
        nw180_[int(3)] = d_1138_v1_
        nw180_[int(4)] = d_1138_v1_
        nw180_[int(5)] = d_1138_v1_
        nw180_[int(6)] = d_1138_v1_
        nw180_[int(7)] = d_1138_v1_
        nw180_[int(8)] = d_1138_v1_
        nw180_[int(9)] = d_1138_v1_
        nw180_[int(10)] = d_1138_v1_
        nw180_[int(11)] = d_1138_v1_
        nw180_[int(12)] = (d_1138_v1_ if not((self).f4) else d_1138_v1_)
        nw180_[int(13)] = d_1138_v1_
        nw180_[int(14)] = d_1138_v1_
        nw180_[int(15)] = (d_1138_v1_ if (self).f4 else d_1138_v1_)
        nw180_[int(16)] = ((d_1140_v3_)[len(d_1141_v4_)] if (len(d_1141_v4_)) in (d_1140_v3_) else d_1138_v1_)
        nw180_[int(17)] = d_1138_v1_
        nw180_[int(18)] = d_1138_v1_
        nw180_[int(19)] = d_1138_v1_
        nw180_[int(20)] = d_1138_v1_
        nw180_[int(21)] = d_1138_v1_
        nw180_[int(22)] = (d_1142_v5_)[default__.safeIndex((d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))], len(d_1142_v5_))]
        nw180_[int(23)] = d_1138_v1_
        d_1143_v6_ = nw180_
        index160_ = default__.safeIndex(463, (d_1143_v6_).length(0))
        (d_1143_v6_)[index160_] = d_1138_v1_
        d_1144_v7_: D2
        d_1144_v7_ = D2_DC7(len(default__.fm37(p0, (self).f4, (self).f6, p0, globalState)))
        d_1145_v8_: _dafny.MultiSet
        d_1145_v8_ = _dafny.MultiSet([d_1144_v7_, d_1144_v7_])
        d_1146_v9_: _dafny.Set
        d_1146_v9_ = _dafny.Set({p0, p0})
        d_1147_v10_: _dafny.Map
        d_1147_v10_ = _dafny.Map({((d_1145_v8_).cardinality) == ((d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))]): (len(d_1146_v9_)) != ((d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))])})
        d_1148_v11_: bool
        d_1148_v11_ = True
        d_1149_v12_: D16
        d_1149_v12_ = D16_DC32(d_1138_v1_)
        d_1150_v13_: _dafny.Map
        d_1150_v13_ = _dafny.Map({d_1138_v1_: (d_1149_v12_).cf51})
        index161_ = default__.safeIndex(463, (d_1143_v6_).length(0))
        index162_ = default__.safeIndex(699, (d_1138_v1_).length(0))
        rhs158_ = ((d_1150_v13_)[d_1138_v1_] if (d_1138_v1_) in (d_1150_v13_) else d_1138_v1_)
        rhs159_ = (self).f6
        rhs160_ = (d_1147_v10_) | ((d_1147_v10_) | (d_1147_v10_))
        rhs161_ = not ((self).f4) or ((self).f4)
        lhs94_ = d_1143_v6_
        lhs95_ = default__.safeIndex(463, (d_1143_v6_).length(0))
        lhs96_ = d_1138_v1_
        lhs97_ = default__.safeIndex(699, (d_1138_v1_).length(0))
        lhs94_[lhs95_] = rhs158_
        lhs96_[lhs97_] = rhs159_
        d_1147_v10_ = rhs160_
        d_1148_v11_ = rhs161_
        d_1151_v14_: _dafny.Array
        nw181_ = _dafny.Array(_dafny.Array(None, 0), 18)
        d_1151_v14_ = nw181_
        d_1152_v15_: _dafny.Seq
        d_1152_v15_ = _dafny.SeqWithoutIsStrInference([(self).f4])
        d_1153_v16_: _dafny.Map
        d_1153_v16_ = _dafny.Map({(self).f4: (d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))]})
        d_1154_v17_: _dafny.Seq
        d_1154_v17_ = _dafny.SeqWithoutIsStrInference([d_1153_v16_, _dafny.Map({p0: (0) - ((d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))])}), d_1153_v16_])
        index163_ = default__.safeIndex(699, (d_1138_v1_).length(0))
        index164_ = default__.safeIndex(699, (d_1138_v1_).length(0))
        rhs162_ = (d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))]
        rhs163_ = len(((d_1152_v15_) + (d_1152_v15_) if (self).f4 else (d_1152_v15_) + (d_1152_v15_)))
        rhs164_ = (d_1153_v16_) not in ((d_1154_v17_) + (d_1154_v17_))
        rhs165_ = d_1148_v11_
        rhs166_ = d_1151_v14_
        lhs98_ = d_1138_v1_
        lhs99_ = default__.safeIndex(699, (d_1138_v1_).length(0))
        lhs100_ = d_1138_v1_
        lhs101_ = default__.safeIndex(699, (d_1138_v1_).length(0))
        lhs98_[lhs99_] = rhs162_
        lhs100_[lhs101_] = rhs163_
        d_1148_v11_ = rhs164_
        d_1148_v11_ = rhs165_
        d_1151_v14_ = rhs166_
        d_1155_v18_: T0
        nw182_ = C2()
        nw182_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1156_i0_ in range(default__.abs(587))]), p0)
        d_1155_v18_ = nw182_
        d_1157_v19_: _dafny.Map
        d_1157_v19_ = _dafny.Map({(d_1155_v18_).f4: d_1155_v18_})
        d_1158_v20_: _dafny.Map
        d_1158_v20_ = _dafny.Map({(self).f6: d_1155_v18_})
        d_1159_v21_: _dafny.Seq
        d_1159_v21_ = _dafny.SeqWithoutIsStrInference([d_1155_v18_, ((d_1157_v19_)[p0] if (p0) in (d_1157_v19_) else ((d_1158_v20_)[-228] if (-228) in (d_1158_v20_) else d_1155_v18_)), d_1155_v18_, d_1155_v18_, d_1155_v18_])
        index165_ = default__.safeIndex(699, (d_1138_v1_).length(0))
        (d_1138_v1_)[index165_] = (len(d_1159_v21_)) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1160_i1_ in range(default__.abs(-701))])))
        d_1161_v22_: _dafny.Array
        def lambda33_(d_1162_i2_):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fag")))

        init21_ = lambda33_
        nw183_ = _dafny.Array(None, 14)
        for i0_21_ in range(nw183_.length(0)):
            nw183_[i0_21_] = init21_(i0_21_)
        d_1161_v22_ = nw183_
        d_1163_v23_: _dafny.Seq
        d_1163_v23_ = _dafny.SeqWithoutIsStrInference([d_1141_v4_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1164_i3_ in range(default__.abs(819))])])
        d_1165_v24_: str
        d_1165_v24_ = _dafny.CodePoint('f')
        index166_ = default__.safeIndex(290, (d_1161_v22_).length(0))
        (d_1161_v22_)[index166_] = (((d_1163_v23_)[default__.safeIndex((self).f6, len(d_1163_v23_))]).set(default__.safeIndex((self).f6, len((d_1163_v23_)[default__.safeIndex((self).f6, len(d_1163_v23_))])), d_1165_v24_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nusipv")))
        index167_ = default__.safeIndex(290, (d_1161_v22_).length(0))
        (d_1161_v22_)[index167_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1166_i4_ in range(default__.abs(0))])
        if default__.fm0(d_1165_v24_, globalState):
            d_1167_v25_: _dafny.Array
            nw184_ = _dafny.Array(_dafny.CodePoint('D'), 21)
            d_1167_v25_ = nw184_
            index168_ = default__.safeIndex(647, (d_1151_v14_).length(0))
            (d_1151_v14_)[index168_] = d_1167_v25_
            index169_ = default__.safeIndex(647, (d_1151_v14_).length(0))
            (d_1151_v14_)[index169_] = d_1167_v25_
            d_1148_v11_ = not(p0)
            d_1168_v26_: D11
            d_1168_v26_ = D11_DC26((d_1155_v18_).f4)
            d_1169_v27_: _dafny.Map
            d_1169_v27_ = _dafny.Map({d_1168_v26_: (d_1155_v18_).f4})
            d_1170_v28_: _dafny.Seq
            d_1170_v28_ = _dafny.SeqWithoutIsStrInference([d_1169_v27_])
            d_1171_v29_: D18
            d_1171_v29_ = D18_DC35(d_1170_v28_)
            d_1172_v30_: _dafny.Seq
            d_1172_v30_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1165_v24_ for d_1173_i5_ in range(default__.abs(451))])), (d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))]])
            d_1174_v31_: _dafny.Seq
            d_1174_v31_ = _dafny.SeqWithoutIsStrInference([(d_1172_v30_)[default__.safeIndex((d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))], len(d_1172_v30_))]])
            d_1175_v32_: _dafny.Map
            d_1175_v32_ = _dafny.Map({len(d_1172_v30_): d_1147_v10_})
            d_1176_v33_: _dafny.Seq
            d_1176_v33_ = _dafny.SeqWithoutIsStrInference([d_1175_v32_, d_1175_v32_])
            d_1177_v34_: C1
            nw185_ = C1()
            nw185_.ctor__((0) - (len((d_1171_v29_).cf58)), len((d_1176_v33_)[default__.safeIndex(len(_dafny.Set({d_1148_v11_})), len(d_1176_v33_))]), self.f5, ((d_1174_v31_)[default__.safeIndex((d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))], len(d_1174_v31_))]) - ((d_1155_v18_).fm7(globalState)), True)
            d_1177_v34_ = nw185_
            d_1178_v35_: _dafny.Set
            d_1178_v35_ = _dafny.Set({_dafny.MultiSet(d_1152_v15_)})
            d_1179_v36_: _dafny.Array
            nw186_ = _dafny.Array(None, 27)
            nw186_[int(0)] = default__.fm0(d_1165_v24_, globalState)
            nw186_[int(1)] = not(((self).f6) <= ((d_1177_v34_).f15))
            nw186_[int(2)] = (d_1178_v35_).issubset(d_1178_v35_)
            nw186_[int(3)] = (self).f4
            nw186_[int(4)] = False
            nw186_[int(5)] = d_1148_v11_
            nw186_[int(6)] = (self).f4
            nw186_[int(7)] = (self).f4
            nw186_[int(8)] = True
            nw186_[int(9)] = (d_1155_v18_).f4
            nw186_[int(10)] = (self).f4
            nw186_[int(11)] = p0
            nw186_[int(12)] = d_1148_v11_
            nw186_[int(13)] = (d_1146_v9_).issubset(d_1146_v9_)
            nw186_[int(14)] = (d_1155_v18_).f4
            nw186_[int(15)] = (self).f4
            nw186_[int(16)] = p0
            nw186_[int(17)] = ((d_1155_v18_).f4) or (p0)
            nw186_[int(18)] = p0
            nw186_[int(19)] = (d_1155_v18_).f4
            nw186_[int(20)] = (self).f4
            nw186_[int(21)] = p0
            nw186_[int(22)] = p0
            nw186_[int(23)] = d_1148_v11_
            nw186_[int(24)] = True
            nw186_[int(25)] = (d_1155_v18_).f4
            nw186_[int(26)] = p0
            d_1179_v36_ = nw186_
            d_1180_v37_: _dafny.Set
            d_1180_v37_ = _dafny.Set({(d_1161_v22_)[default__.safeIndex(290, (d_1161_v22_).length(0))]})
            index170_ = default__.safeIndex(815, (d_1179_v36_).length(0))
            (d_1179_v36_)[index170_] = (d_1180_v37_).ispropersubset(d_1180_v37_)
            index171_ = default__.safeIndex(699, (d_1138_v1_).length(0))
            index172_ = default__.safeIndex(815, (d_1179_v36_).length(0))
            rhs167_ = default__.safeModuloInt((self).f6, ((d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))]) * (d_1177_v34_.f16))
            rhs168_ = p0
            rhs169_ = (d_1155_v18_).f4
            lhs102_ = d_1138_v1_
            lhs103_ = default__.safeIndex(699, (d_1138_v1_).length(0))
            lhs104_ = d_1179_v36_
            lhs105_ = default__.safeIndex(815, (d_1179_v36_).length(0))
            lhs102_[lhs103_] = rhs167_
            d_1148_v11_ = rhs168_
            lhs104_[lhs105_] = rhs169_
            d_1148_v11_ = not((D12_DC28(default__.fm3(len(((d_1152_v15_).set(default__.safeIndex((d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))], len(d_1152_v15_)), not(p0))).set(default__.safeIndex(d_1177_v34_.f16, len((d_1152_v15_).set(default__.safeIndex((d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))], len(d_1152_v15_)), not(p0)))), (d_1179_v36_)[default__.safeIndex(815, (d_1179_v36_).length(0))])), (self).f4, globalState), d_1177_v34_.f16, (d_1155_v18_).f4, (d_1179_v36_)[default__.safeIndex(815, (d_1179_v36_).length(0))], (self).f6)).cf46)
        elif True:
            index173_ = default__.safeIndex(699, (d_1138_v1_).length(0))
            (d_1138_v1_)[index173_] = (self).f6
            index174_ = default__.safeIndex(699, (d_1138_v1_).length(0))
            (d_1138_v1_)[index174_] = (self).f6
            d_1148_v11_ = True
            d_1181_v38_: C2
            nw187_ = C2()
            nw187_.ctor__((d_1161_v22_)[default__.safeIndex(290, (d_1161_v22_).length(0))], d_1148_v11_)
            d_1181_v38_ = nw187_
            d_1182_v39_: _dafny.MultiSet
            d_1182_v39_ = _dafny.MultiSet([d_1181_v38_, d_1181_v38_, d_1181_v38_])
            d_1183_v40_: _dafny.Array
            nw188_ = _dafny.Array(None, 10)
            nw188_[int(0)] = (d_1181_v38_) not in ((d_1182_v39_).set(d_1181_v38_, default__.abs((self).f6)))
            nw188_[int(1)] = True
            nw188_[int(2)] = default__.fm0(d_1165_v24_, globalState)
            nw188_[int(3)] = (d_1148_v11_ if (d_1155_v18_).f4 else True)
            nw188_[int(4)] = (d_1146_v9_) != (d_1146_v9_)
            nw188_[int(5)] = (d_1137_v0_) == (d_1137_v0_)
            nw188_[int(6)] = ((self).f6) > ((self).f6)
            nw188_[int(7)] = d_1148_v11_
            nw188_[int(8)] = (self).f4
            nw188_[int(9)] = (d_1155_v18_).f4
            d_1183_v40_ = nw188_
            index175_ = default__.safeIndex(580, (d_1183_v40_).length(0))
            (d_1183_v40_)[index175_] = d_1148_v11_
            index176_ = default__.safeIndex(580, (d_1183_v40_).length(0))
            (d_1183_v40_)[index176_] = (d_1155_v18_).f4
            d_1184_v41_: C1
            nw189_ = C1()
            nw189_.ctor__((d_1138_v1_)[default__.safeIndex(699, (d_1138_v1_).length(0))], 806, self.f5, ((self).f6) * ((self).f6), d_1148_v11_)
            d_1184_v41_ = nw189_
        r0 = d_1165_v24_
        d_1185_v42_: _dafny.Array
        nw190_ = _dafny.Array(D1.default()(), 23)
        d_1185_v42_ = nw190_
        r1 = d_1185_v42_
        return r0, r1

    def m9(self, p0, globalState):
        r0: bool = False
        r0 = not(p0)
        d_1186_v0_: _dafny.Array
        def lambda34_(d_1187_i0_):
            return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1188_i1_ in range(default__.abs(773))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqvrar")))

        init22_ = lambda34_
        nw191_ = _dafny.Array(None, 22)
        for i0_22_ in range(nw191_.length(0)):
            nw191_[i0_22_] = init22_(i0_22_)
        d_1186_v0_ = nw191_
        index177_ = default__.safeIndex(721, (d_1186_v0_).length(0))
        (d_1186_v0_)[index177_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
        d_1189_v1_: _dafny.Seq
        d_1189_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
        index178_ = default__.safeIndex(721, (d_1186_v0_).length(0))
        (d_1186_v0_)[index178_] = (default__.fm37(not((self).f4), (self).f4, len(d_1189_v1_), (self).f4, globalState)) + (d_1189_v1_)
        d_1190_v2_: str
        d_1190_v2_ = _dafny.CodePoint('v')
        index179_ = default__.safeIndex(721, (d_1186_v0_).length(0))
        (d_1186_v0_)[index179_] = (d_1189_v1_).set(default__.safeIndex((0) - ((self).f6), len(d_1189_v1_)), d_1190_v2_)
        d_1191_v3_: _dafny.Seq
        d_1191_v3_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6, (self).f6])
        d_1192_v4_: C1
        nw192_ = C1()
        nw192_.ctor__((0) - ((self).f6), len((d_1191_v3_) + (d_1191_v3_)), self.f5, (self).f6, not((self).f4))
        d_1192_v4_ = nw192_
        if not((self).f4):
            d_1193_v5_: _dafny.Array
            def lambda35_(d_1194_i2_):
                return (self).f4

            init23_ = lambda35_
            nw193_ = _dafny.Array(None, 5)
            for i0_23_ in range(nw193_.length(0)):
                nw193_[i0_23_] = init23_(i0_23_)
            d_1193_v5_ = nw193_
            d_1195_v6_: _dafny.Map
            d_1195_v6_ = _dafny.Map({True: d_1193_v5_})
            d_1196_v7_: _dafny.Seq
            d_1196_v7_ = _dafny.SeqWithoutIsStrInference([d_1193_v5_, d_1193_v5_, d_1193_v5_])
            d_1193_v5_ = ((d_1195_v6_)[(self).f4] if ((self).f4) in (d_1195_v6_) else (d_1196_v7_)[default__.safeIndex((d_1192_v4_).f15, len(d_1196_v7_))])
            d_1197_v8_: C2
            nw194_ = C2()
            nw194_.ctor__(d_1189_v1_, p0)
            d_1197_v8_ = nw194_
            (d_1192_v4_).f16 = (d_1192_v4_).f15
            d_1198_v9_: _dafny.Map
            d_1198_v9_ = _dafny.Map({(d_1192_v4_).f15: (self).f4})
            d_1198_v9_ = (_dafny.Map({(d_1192_v4_).f15: (self).f4})) | (_dafny.Map({(d_1192_v4_).f15: p0}))
            d_1199_v10_: D10
            d_1199_v10_ = D10_DC22(_dafny.SeqWithoutIsStrInference([d_1190_v2_ for d_1200_i3_ in range(default__.abs(791))]), False, (self).f4, d_1193_v5_, 184)
            d_1201_v11_: D1
            d_1201_v11_ = D1_DC4((d_1199_v10_).cf34, (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lroyicek")))) - (455))
            d_1201_v11_ = d_1201_v11_
        elif True:
            r0 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtossvsa"))) < (((d_1186_v0_)[default__.safeIndex(721, (d_1186_v0_).length(0))]) + ((d_1186_v0_)[default__.safeIndex(721, (d_1186_v0_).length(0))]))
            r0 = ((d_1192_v4_.f16 if (self).f4 else (self).f6)) >= (len((d_1186_v0_)[default__.safeIndex(721, (d_1186_v0_).length(0))]))
            d_1202_v12_: _dafny.Array
            nw195_ = _dafny.Array(None, 2)
            nw195_[int(0)] = d_1190_v2_
            nw195_[int(1)] = d_1190_v2_
            d_1202_v12_ = nw195_
            d_1203_v13_: _dafny.Map
            d_1203_v13_ = _dafny.Map({p0: d_1202_v12_})
            d_1204_v14_: _dafny.Seq
            d_1204_v14_ = _dafny.SeqWithoutIsStrInference([d_1202_v12_])
            d_1205_v15_: _dafny.Array
            nw196_ = _dafny.Array(None, 8)
            nw196_[int(0)] = d_1202_v12_
            nw196_[int(1)] = d_1202_v12_
            nw196_[int(2)] = d_1202_v12_
            nw196_[int(3)] = d_1202_v12_
            nw196_[int(4)] = d_1202_v12_
            nw196_[int(5)] = ((d_1203_v13_)[p0] if (p0) in (d_1203_v13_) else (d_1204_v14_)[default__.safeIndex((0) - ((d_1192_v4_).f15), len(d_1204_v14_))])
            nw196_[int(6)] = d_1202_v12_
            nw196_[int(7)] = d_1202_v12_
            d_1205_v15_ = nw196_
            index180_ = default__.safeIndex(472, (d_1205_v15_).length(0))
            (d_1205_v15_)[index180_] = d_1202_v12_
            d_1206_v16_: _dafny.MultiSet
            d_1206_v16_ = _dafny.MultiSet([d_1192_v4_.f16])
            index181_ = default__.safeIndex(721, (d_1186_v0_).length(0))
            index182_ = default__.safeIndex(472, (d_1205_v15_).length(0))
            rhs170_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsbsmmac"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ps")))) + ((d_1186_v0_)[default__.safeIndex(721, (d_1186_v0_).length(0))])
            rhs171_ = ((d_1206_v16_)[d_1192_v4_.f16] if (d_1192_v4_.f16) in (d_1206_v16_) else (d_1192_v4_).f15)
            rhs172_ = d_1202_v12_
            lhs106_ = d_1186_v0_
            lhs107_ = default__.safeIndex(721, (d_1186_v0_).length(0))
            lhs108_ = d_1192_v4_
            lhs109_ = d_1205_v15_
            lhs110_ = default__.safeIndex(472, (d_1205_v15_).length(0))
            lhs106_[lhs107_] = rhs170_
            lhs108_.f16 = rhs171_
            lhs109_[lhs110_] = rhs172_
            d_1207_v18_: _dafny.Array
            def lambda36_(d_1208_v3_):
                def lambda37_(d_1209_i4_):
                    def iife88_():
                        coll44_ = _dafny.Set()
                        compr_44_: int
                        for compr_44_ in (d_1208_v3_).Elements:
                            d_1210_v17_: int = compr_44_
                            if (d_1210_v17_) in (d_1208_v3_):
                                coll44_ = coll44_.union(_dafny.Set([(d_1210_v17_) + (-142)]))
                        return _dafny.Set(coll44_)
                    return iife88_()
                    

                return lambda37_

            init24_ = lambda36_(d_1191_v3_)
            nw197_ = _dafny.Array(None, 1)
            for i0_24_ in range(nw197_.length(0)):
                nw197_[i0_24_] = init24_(i0_24_)
            d_1207_v18_ = nw197_
            d_1211_v19_: _dafny.Array
            nw198_ = _dafny.Array(None, 4)
            nw198_[int(0)] = d_1207_v18_
            nw198_[int(1)] = (d_1207_v18_ if p0 else d_1207_v18_)
            nw198_[int(2)] = (d_1207_v18_ if p0 else d_1207_v18_)
            nw198_[int(3)] = d_1207_v18_
            d_1211_v19_ = nw198_
            index183_ = default__.safeIndex(373, (d_1211_v19_).length(0))
            (d_1211_v19_)[index183_] = d_1207_v18_
            index184_ = default__.safeIndex(373, (d_1211_v19_).length(0))
            (d_1211_v19_)[index184_] = (d_1207_v18_ if (self).f4 else d_1207_v18_)
            if True:
                d_1212_v20_: _dafny.Array
                def lambda38_(d_1213_v16_):
                    def lambda39_(d_1214_i5_):
                        return _dafny.Map({(self).f4: (d_1213_v16_).cardinality})

                    return lambda39_

                init25_ = lambda38_(d_1206_v16_)
                nw199_ = _dafny.Array(None, 25)
                for i0_25_ in range(nw199_.length(0)):
                    nw199_[i0_25_] = init25_(i0_25_)
                d_1212_v20_ = nw199_
                d_1215_v21_: _dafny.Map
                d_1215_v21_ = _dafny.Map({(self).f4: d_1192_v4_.f16})
                index185_ = default__.safeIndex(655, (d_1212_v20_).length(0))
                (d_1212_v20_)[index185_] = d_1215_v21_
                d_1216_v22_: _dafny.Map
                d_1216_v22_ = _dafny.Map({(self).f4: (self).f4})
                d_1217_v23_: _dafny.Map
                d_1217_v23_ = _dafny.Map({(self).f4: ((d_1216_v22_)[(self).f4] if ((self).f4) in (d_1216_v22_) else p0)})
                d_1218_v24_: _dafny.Seq
                d_1218_v24_ = _dafny.SeqWithoutIsStrInference([d_1216_v22_, d_1217_v23_, d_1216_v22_])
                index186_ = default__.safeIndex(472, (d_1205_v15_).length(0))
                index187_ = default__.safeIndex(655, (d_1212_v20_).length(0))
                rhs173_ = (d_1218_v24_) == (d_1218_v24_)
                rhs174_ = d_1202_v12_
                rhs175_ = d_1215_v21_
                rhs176_ = d_1206_v16_
                lhs111_ = d_1205_v15_
                lhs112_ = default__.safeIndex(472, (d_1205_v15_).length(0))
                lhs113_ = d_1212_v20_
                lhs114_ = default__.safeIndex(655, (d_1212_v20_).length(0))
                r0 = rhs173_
                lhs111_[lhs112_] = rhs174_
                lhs113_[lhs114_] = rhs175_
                d_1206_v16_ = rhs176_
                d_1219_v25_: _dafny.Array
                nw200_ = _dafny.Array(_dafny.Map({}), 22)
                d_1219_v25_ = nw200_
                d_1220_v26_: _dafny.MultiSet
                d_1220_v26_ = _dafny.MultiSet([p0])
                d_1221_v27_: _dafny.Map
                d_1221_v27_ = _dafny.Map({(d_1220_v26_).cardinality: 548})
                index188_ = default__.safeIndex(11, (d_1219_v25_).length(0))
                (d_1219_v25_)[index188_] = d_1221_v27_
                d_1222_v28_: _dafny.Set
                d_1222_v28_ = _dafny.Set({p0})
                index189_ = default__.safeIndex(11, (d_1219_v25_).length(0))
                (d_1219_v25_)[index189_] = _dafny.Map({len(d_1222_v28_): 580})
                r0 = (self).f4
                (d_1192_v4_).f16 = d_1192_v4_.f16
                (d_1192_v4_).f16 = (d_1192_v4_.f16) + (((d_1192_v4_).f15) * (d_1192_v4_.f16))
            elif True:
                d_1223_v29_: _dafny.Map
                d_1223_v29_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1190_v2_ for d_1224_i6_ in range(default__.abs(4))]): 307})
                d_1225_v30_: _dafny.Map
                d_1225_v30_ = _dafny.Map({(self).f4: d_1223_v29_})
                d_1226_v31_: _dafny.Array
                def lambda40_(d_1227_i7_):
                    return True

                init26_ = lambda40_
                nw201_ = _dafny.Array(None, 16)
                for i0_26_ in range(nw201_.length(0)):
                    nw201_[i0_26_] = init26_(i0_26_)
                d_1226_v31_ = nw201_
                d_1228_v32_: int
                out36_: int
                out36_ = (d_1192_v4_).m11(p0, p0, ((d_1225_v30_)[p0] if (p0) in (d_1225_v30_) else _dafny.Map({(d_1186_v0_)[default__.safeIndex(721, (d_1186_v0_).length(0))]: d_1192_v4_.f16})), d_1226_v31_, globalState)
                d_1228_v32_ = out36_
                d_1229_v33_: _dafny.Map
                d_1229_v33_ = _dafny.Map({(d_1192_v4_).f15: p0})
                d_1230_v34_: _dafny.Seq
                d_1230_v34_ = _dafny.SeqWithoutIsStrInference([((d_1229_v33_)[(d_1192_v4_).f15] if ((d_1192_v4_).f15) in (d_1229_v33_) else False)])
                d_1231_v35_: _dafny.Seq
                d_1231_v35_ = _dafny.SeqWithoutIsStrInference([d_1230_v34_, d_1230_v34_])
                d_1232_v36_: D12
                d_1232_v36_ = D12_DC27(d_1231_v35_)
                pat_let_tv16_ = d_1231_v35_
                pat_let_tv17_ = d_1189_v1_
                pat_let_tv18_ = d_1231_v35_
                pat_let_tv19_ = d_1231_v35_
                pat_let_tv20_ = d_1192_v4_
                pat_let_tv21_ = d_1231_v35_
                pat_let_tv22_ = d_1231_v35_
                def iife90_(_pat_let23_0):
                    def iife91_(d_1233_dt__update__tmp_h0_):
                        def iife92_(_pat_let24_0):
                            def iife93_(d_1234_dt__update_hcf42_h0_):
                                return D12_DC27(d_1234_dt__update_hcf42_h0_)
                            return iife93_(_pat_let24_0)
                        return iife92_((pat_let_tv16_).set(default__.safeIndex(len(pat_let_tv17_), len(pat_let_tv18_)), (pat_let_tv19_)[default__.safeIndex(pat_let_tv20_.f16, len(pat_let_tv21_))]))
                    return iife91_(_pat_let23_0)
                def iife89_(_pat_let22_0):
                    def iife94_(d_1235_dt__update__tmp_h1_):
                        def iife95_(_pat_let25_0):
                            def iife96_(d_1236_dt__update_hcf42_h1_):
                                return D12_DC27(d_1236_dt__update_hcf42_h1_)
                            return iife96_(_pat_let25_0)
                        return iife95_(pat_let_tv22_)
                    return iife94_(_pat_let22_0)
                d_1232_v36_ = iife89_(iife90_(d_1232_v36_))
                arr5_ = self.f5
                index190_ = default__.safeIndex(533, (self.f5).length(0))
                arr5_[index190_] = d_1226_v31_
                arr6_ = self.f5
                index191_ = default__.safeIndex(533, (self.f5).length(0))
                arr6_[index191_] = d_1226_v31_
                d_1237_v37_: _dafny.Map
                d_1237_v37_ = _dafny.Map({(self).f4: d_1226_v31_})
                d_1238_v38_: _dafny.Map
                d_1238_v38_ = _dafny.Map({d_1190_v2_: ((d_1237_v37_)[(self).f4] if ((self).f4) in (d_1237_v37_) else (self.f5)[default__.safeIndex(533, (self.f5).length(0))])})
                d_1239_v39_: D19
                d_1239_v39_ = D19_DC37(d_1238_v38_)
                rhs177_ = (self).f4
                rhs178_ = (d_1239_v39_).cf62
                r0 = rhs177_
                d_1238_v38_ = rhs178_
                d_1240_v40_: str
                d_1241_v41_: _dafny.Array
                out37_: str
                out38_: _dafny.Array
                out37_, out38_ = (d_1192_v4_).m2(p0, globalState)
                d_1240_v40_ = out37_
                d_1241_v41_ = out38_
        d_1242_v42_: _dafny.Map
        d_1242_v42_ = _dafny.Map({p0: p0})
        d_1243_v43_: _dafny.Map
        d_1243_v43_ = _dafny.Map({d_1242_v42_: d_1190_v2_})
        d_1244_v44_: _dafny.Map
        d_1244_v44_ = _dafny.Map({(self).f6: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcrpadd"))})
        d_1245_v45_: bool
        d_1246_v46_: int
        d_1247_v47_: int
        d_1248_v48_: bool
        out39_: bool
        out40_: int
        out41_: int
        out42_: bool
        out39_, out40_, out41_, out42_ = (self).m10(_dafny.MultiSet([len(d_1243_v43_), (self).f6, d_1192_v4_.f16]), (d_1189_v1_) < (((d_1244_v44_)[d_1192_v4_.f16] if (d_1192_v4_.f16) in (d_1244_v44_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))), globalState)
        d_1245_v45_ = out39_
        d_1246_v46_ = out40_
        d_1247_v47_ = out41_
        d_1248_v48_ = out42_
        r0 = p0
        return r0

    def m10(self, p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        r3: bool = False
        d_1249_v0_: _dafny.Array
        nw202_ = _dafny.Array(int(0), 28)
        d_1249_v0_ = nw202_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_1249_v0_).length(0)):
            d_1250_i0_: int = guard_loop_6_
            if (True) and (((0) <= (d_1250_i0_)) and ((d_1250_i0_) < ((d_1249_v0_).length(0)))):
                (d_1249_v0_)[(d_1250_i0_)] = default__.safeModuloInt(d_1250_i0_, (self).f6)
        hi6_ = -365
        for d_1251_i1_ in range((self).f6, hi6_):
            d_1252_v1_: D16
            d_1252_v1_ = D16_DC32(d_1249_v0_)
            source22_ = d_1252_v1_
            if source22_.is_DC33:
                d_1253___mcc_h0_ = source22_.cf52
                d_1254___mcc_h1_ = source22_.cf53
                d_1255___mcc_h2_ = source22_.cf54
                d_1256___mcc_h3_ = source22_.cf55
                d_1257___mcc_h4_ = source22_.cf56
                d_1258_cf56_ = d_1257___mcc_h4_
                d_1259_cf55_ = d_1256___mcc_h3_
                d_1260_cf54_ = d_1255___mcc_h2_
                d_1261_cf53_ = d_1254___mcc_h1_
                d_1262_cf52_ = d_1253___mcc_h0_
                d_1262_cf52_ = not(True)
                d_1263_v2_: _dafny.Seq
                d_1263_v2_ = _dafny.SeqWithoutIsStrInference([d_1262_cf52_, False])
                d_1264_v3_: D4
                d_1264_v3_ = D4_DC13((self).f6, d_1262_cf52_, d_1251_i1_, d_1263_v2_)
                r0 = ((d_1264_v3_ if p1 else d_1264_v3_)).cf15
                d_1265_v4_: _dafny.Array
                nw203_ = _dafny.Array(False, 14)
                d_1265_v4_ = nw203_
                d_1265_v4_ = d_1265_v4_
                r2 = (0) - (d_1251_i1_)
            elif True:
                d_1266___mcc_h5_ = source22_.cf51
                d_1267_cf51_ = d_1266___mcc_h5_
                d_1268_v5_: _dafny.Array
                nw204_ = _dafny.Array(False, 16)
                d_1268_v5_ = nw204_
                index192_ = default__.safeIndex(287, (d_1268_v5_).length(0))
                (d_1268_v5_)[index192_] = (self).f4
                d_1269_v6_: D0
                d_1269_v6_ = D0_DC0(len(_dafny.Map({p1: d_1251_i1_})))
                d_1270_v7_: _dafny.Set
                d_1270_v7_ = _dafny.Set({d_1269_v6_})
                d_1271_v8_: _dafny.MultiSet
                d_1271_v8_ = _dafny.MultiSet([(self).f4])
                d_1272_v10_: _dafny.Set
                def iife97_():
                    coll45_ = _dafny.Set()
                    compr_45_: D0
                    for compr_45_ in (_dafny.SeqWithoutIsStrInference([default__.fm38((((d_1271_v8_).set((self).f4, default__.abs(d_1251_i1_))).set((self).f4, default__.abs(d_1251_i1_))).cardinality, d_1251_i1_, p1, (self).f6, globalState)])).Elements:
                        d_1273_v9_: D0 = compr_45_
                        if (d_1273_v9_) in (_dafny.SeqWithoutIsStrInference([default__.fm38((((d_1271_v8_).set((self).f4, default__.abs(d_1251_i1_))).set((self).f4, default__.abs(d_1251_i1_))).cardinality, d_1251_i1_, p1, (self).f6, globalState)])):
                            coll45_ = coll45_.union(_dafny.Set([d_1273_v9_]))
                    return _dafny.Set(coll45_)
                d_1272_v10_ = _dafny.Set({d_1270_v7_, d_1270_v7_, iife97_()
                })
                d_1274_v11_: _dafny.MultiSet
                d_1274_v11_ = _dafny.MultiSet([len((_dafny.Set({d_1270_v7_})) - (d_1272_v10_))])
                index193_ = default__.safeIndex(287, (d_1268_v5_).length(0))
                rhs179_ = not ((self).f4) or ((self).f4)
                rhs180_ = d_1274_v11_
                rhs181_ = len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skjggxj")))]))
                lhs115_ = d_1268_v5_
                lhs116_ = default__.safeIndex(287, (d_1268_v5_).length(0))
                lhs115_[lhs116_] = rhs179_
                d_1274_v11_ = rhs180_
                r1 = rhs181_
                index194_ = default__.safeIndex(287, (d_1268_v5_).length(0))
                (d_1268_v5_)[index194_] = ((d_1268_v5_)[default__.safeIndex(287, (d_1268_v5_).length(0))]) and (((d_1268_v5_)[default__.safeIndex(287, (d_1268_v5_).length(0))]) == (False))
                d_1275_v12_: _dafny.Seq
                d_1275_v12_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f6), (self).f6])
                r2 = len(d_1275_v12_)
                index195_ = default__.safeIndex(287, (d_1268_v5_).length(0))
                (d_1268_v5_)[index195_] = p1
            d_1276_v13_: _dafny.Array
            nw205_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
            d_1276_v13_ = nw205_
            d_1277_v14_: D11
            d_1277_v14_ = D11_DC26((self).f4)
            rhs182_ = d_1276_v13_
            rhs183_ = d_1277_v14_
            d_1276_v13_ = rhs182_
            d_1277_v14_ = rhs183_
            d_1278_v15_: _dafny.Map
            d_1278_v15_ = _dafny.Map({(self).f6: p1})
            d_1279_v16_: C1
            nw206_ = C1()
            nw206_.ctor__((len(d_1278_v15_)) * (d_1251_i1_), d_1251_i1_, self.f5, d_1251_i1_, not ((self).f4) or (not(False)))
            d_1279_v16_ = nw206_
            d_1280_v17_: str
            d_1280_v17_ = _dafny.CodePoint('b')
            r3 = default__.fm0(d_1280_v17_, globalState)
        if (self).f4:
            d_1281_v18_: D3
            d_1281_v18_ = D3_DC10(((self).f6) * ((self).f6))
            source23_ = d_1281_v18_
            if source23_.is_DC9:
                d_1282___mcc_h6_ = source23_.cf8
                d_1283___mcc_h7_ = source23_.cf9
                d_1284___mcc_h8_ = source23_.cf10
                d_1285_cf10_ = d_1284___mcc_h8_
                d_1286_cf9_ = d_1283___mcc_h7_
                d_1287_cf8_ = d_1282___mcc_h6_
                r2 = (self).f6
                d_1288_v19_: _dafny.Seq
                d_1288_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uig"))
                d_1289_v20_: _dafny.Map
                d_1289_v20_ = _dafny.Map({p0: d_1288_v19_})
                d_1290_v21_: D0
                d_1290_v21_ = D0_DC0(767)
                d_1291_v22_: str
                d_1291_v22_ = _dafny.CodePoint('t')
                d_1292_v23_: _dafny.Set
                d_1292_v23_ = _dafny.Set({d_1288_v19_})
                d_1293_v24_: bool
                d_1294_v25_: int
                out43_: bool
                out44_: int
                out43_, out44_ = default__.m0(d_1289_v20_, d_1287_cf8_, (self).fm6(D0_DC0((p0).cardinality), d_1290_v21_, _dafny.Map({d_1286_cf9_: d_1291_v22_}), len(d_1292_v23_), globalState), d_1285_cf10_, globalState)
                d_1293_v24_ = out43_
                d_1294_v25_ = out44_
                d_1295_v26_: _dafny.Array
                nw207_ = _dafny.Array(_dafny.Map({}), 8)
                d_1295_v26_ = nw207_
                d_1296_v27_: _dafny.Map
                d_1296_v27_ = _dafny.Map({d_1281_v18_: default__.fm58(globalState)})
                d_1297_v28_: _dafny.Seq
                d_1297_v28_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                d_1298_v29_: _dafny.Seq
                d_1298_v29_ = _dafny.SeqWithoutIsStrInference([(self).fm8(d_1294_v25_, False, d_1288_v19_, d_1297_v28_, globalState)])
                d_1299_v30_: _dafny.Seq
                d_1299_v30_ = _dafny.SeqWithoutIsStrInference([d_1298_v29_])
                d_1300_v31_: D12
                d_1300_v31_ = D12_DC27((d_1299_v30_).set(default__.safeIndex(len(d_1298_v29_), len(d_1299_v30_)), d_1298_v29_))
                d_1301_v32_: _dafny.Map
                d_1301_v32_ = _dafny.Map({(self).f4: d_1300_v31_})
                index196_ = default__.safeIndex(776, (d_1295_v26_).length(0))
                (d_1295_v26_)[index196_] = ((d_1296_v27_)[d_1281_v18_] if (d_1281_v18_) in (d_1296_v27_) else d_1301_v32_)
                d_1302_v33_: D20
                d_1302_v33_ = D20_DC40(d_1301_v32_)
                index197_ = default__.safeIndex(776, (d_1295_v26_).length(0))
                (d_1295_v26_)[index197_] = (d_1301_v32_) | ((d_1301_v32_) | ((D20_DC40((d_1302_v33_).cf66)).cf66))
                d_1303_v34_: _dafny.Map
                d_1303_v34_ = _dafny.Map({d_1294_v25_: d_1285_cf10_})
                d_1304_v35_: C1
                nw208_ = C1()
                nw208_.ctor__((len(d_1303_v34_)) + (d_1286_cf9_), default__.safeModuloInt(len(d_1288_v19_), d_1294_v25_), self.f5, (self).f6, d_1293_v24_)
                d_1304_v35_ = nw208_
            elif source23_.is_DC10:
                d_1305___mcc_h9_ = source23_.cf11
                d_1306_cf11_ = d_1305___mcc_h9_
                r1 = 491
                d_1307_v36_: _dafny.Seq
                d_1307_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvlmlhsge"))
                r2 = len((d_1307_v36_) + (d_1307_v36_))
                d_1308_v37_: C1
                nw209_ = C1()
                nw209_.ctor__((self).f6, default__.safeModuloInt(753, (self).f6), self.f5, default__.safeModuloInt((self).f6, (self).f6), (self).f4)
                d_1308_v37_ = nw209_
                d_1249_v0_ = d_1249_v0_
            elif source23_.is_DC8:
                d_1309___mcc_h10_ = source23_.cf7
                d_1310_cf7_ = d_1309___mcc_h10_
                d_1311_v38_: str
                d_1311_v38_ = _dafny.CodePoint('e')
                d_1312_v39_: _dafny.Map
                d_1312_v39_ = _dafny.Map({d_1311_v38_: False})
                r0 = not (p1) or (((d_1312_v39_)[d_1311_v38_] if (d_1311_v38_) in (d_1312_v39_) else not((self).f4)))
                d_1313_v40_: _dafny.Seq
                d_1313_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "honftv"))
                r2 = ((self).f6) - ((self).fm9((self).f4, (self).f4, len(d_1313_v40_), globalState))
                d_1314_v41_: _dafny.Array
                def lambda41_(d_1315_i2_):
                    return True

                init27_ = lambda41_
                nw210_ = _dafny.Array(None, 26)
                for i0_27_ in range(nw210_.length(0)):
                    nw210_[i0_27_] = init27_(i0_27_)
                d_1314_v41_ = nw210_
                d_1316_v42_: _dafny.Map
                d_1316_v42_ = _dafny.Map({_dafny.CodePoint('p'): d_1314_v41_})
                d_1317_v43_: D19
                d_1317_v43_ = D19_DC37(d_1316_v42_)
                d_1318_v44_: _dafny.Map
                d_1318_v44_ = _dafny.Map({((self).f4) or (True): d_1317_v43_})
                d_1318_v44_ = (d_1318_v44_).set((self).f4, d_1317_v43_)
                d_1319_v45_: C2
                nw211_ = C2()
                nw211_.ctor__(d_1313_v40_, (self).f4)
                d_1319_v45_ = nw211_
                d_1320_v46_: _dafny.Seq
                d_1320_v46_ = _dafny.SeqWithoutIsStrInference([d_1319_v45_, d_1319_v45_, d_1319_v45_])
                d_1321_v47_: _dafny.Seq
                d_1321_v47_ = _dafny.SeqWithoutIsStrInference([d_1320_v46_, _dafny.SeqWithoutIsStrInference([d_1319_v45_])])
                r3 = ((d_1321_v47_).set(default__.safeIndex((self).f6, len(d_1321_v47_)), d_1320_v46_)) != (d_1321_v47_)
            elif True:
                d_1322___mcc_h11_ = source23_.cf12
                d_1323_cf12_ = d_1322___mcc_h11_
                d_1324_v48_: _dafny.Seq
                d_1324_v48_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])
                d_1324_v48_ = _dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference([p1])) + (_dafny.SeqWithoutIsStrInference([(self).f4, (self).f4]))) for d_1325_i3_ in range(default__.abs(904))])
                d_1326_v49_: _dafny.Seq
                d_1326_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwcahcko"))
                d_1327_v50_: str
                d_1327_v50_ = _dafny.CodePoint('h')
                d_1328_v51_: D9
                d_1328_v51_ = D9_DC19(d_1326_v49_)
                d_1326_v49_ = (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylnncpc")) if (self).f4 else d_1326_v49_)).set(default__.safeIndex((self).f6, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylnncpc")) if (self).f4 else d_1326_v49_))), d_1327_v50_)) + ((d_1328_v51_).cf25)
                d_1329_v52_: D12
                d_1329_v52_ = D12_DC28((self).f6, (0) - ((self).f6), (self).f4, p1, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iunuhhphg")) if (self).f4 else _dafny.SeqWithoutIsStrInference([d_1327_v50_ for d_1330_i4_ in range(default__.abs(651))]))))
                rhs184_ = d_1329_v52_
                rhs185_ = (True) or ((self).f4)
                rhs186_ = d_1326_v49_
                d_1329_v52_ = rhs184_
                r0 = rhs185_
                d_1326_v49_ = rhs186_
                r3 = p1
            d_1331_v53_: _dafny.Seq
            d_1331_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukokag"))
            r1 = len(default__.fm37(p1, not(p1), len(d_1331_v53_), (self).fm8((self).f6, p1, d_1331_v53_, _dafny.SeqWithoutIsStrInference([len(d_1331_v53_), (self).f6]), globalState), globalState))
            r2 = (0) - ((self).f6)
            d_1332_v54_: _dafny.Set
            d_1332_v54_ = _dafny.Set({(self).f4, p1})
            d_1333_v55_: _dafny.Seq
            d_1333_v55_ = _dafny.SeqWithoutIsStrInference([(d_1332_v54_).issubset(d_1332_v54_)])
            r0 = (d_1333_v55_)[default__.safeIndex(((self).f6) - (779), len(d_1333_v55_))]
            d_1334_v56_: str
            d_1334_v56_ = _dafny.CodePoint('i')
            d_1334_v56_ = d_1334_v56_
        elif True:
            r1 = (self).f6
            d_1335_v57_: _dafny.Seq
            d_1335_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcdy"))
            d_1336_v58_: _dafny.Map
            d_1336_v58_ = _dafny.Map({(self).f4: d_1335_v57_})
            d_1337_v59_: _dafny.Map
            d_1337_v59_ = _dafny.Map({d_1335_v57_: d_1336_v58_})
            d_1338_v60_: _dafny.Map
            d_1338_v60_ = _dafny.Map({(self).f6: d_1336_v58_})
            d_1337_v59_ = (d_1337_v59_).set(d_1335_v57_, (((d_1338_v60_)[len(d_1335_v57_)] if (len(d_1335_v57_)) in (d_1338_v60_) else d_1336_v58_)) | (d_1336_v58_))
            d_1339_v61_: _dafny.Seq
            d_1339_v61_ = _dafny.SeqWithoutIsStrInference([(self).f6])
            r3 = (self).fm8((self).f6, (p1) and ((self).f4), d_1335_v57_, d_1339_v61_, globalState)
            index198_ = default__.safeIndex(459, (d_1249_v0_).length(0))
            (d_1249_v0_)[index198_] = (self).f6
            index199_ = default__.safeIndex(459, (d_1249_v0_).length(0))
            (d_1249_v0_)[index199_] = (self).f6
            d_1340_v62_: _dafny.Array
            nw212_ = _dafny.Array(_dafny.Set({}), 14)
            d_1340_v62_ = nw212_
            d_1341_v63_: _dafny.Set
            d_1341_v63_ = _dafny.Set({(self).f6, len(d_1335_v57_)})
            index200_ = default__.safeIndex(216, (d_1340_v62_).length(0))
            (d_1340_v62_)[index200_] = d_1341_v63_
            index201_ = default__.safeIndex(216, (d_1340_v62_).length(0))
            (d_1340_v62_)[index201_] = d_1341_v63_
        d_1342_v64_: _dafny.Seq
        d_1342_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iktm"))
        hi7_ = (self).f6
        for d_1343_i5_ in range(((self).f6) - (len(d_1342_v64_)), hi7_):
            d_1344_v65_: _dafny.Seq
            d_1344_v65_ = _dafny.SeqWithoutIsStrInference([(self).f6])
            r2 = (d_1344_v65_)[default__.safeIndex((self).f6, len(d_1344_v65_))]
            d_1345_v66_: _dafny.MultiSet
            d_1345_v66_ = _dafny.MultiSet([(self).f4, not((self).f4)])
            pat_let_tv23_ = d_1342_v64_
            def iife99_(_pat_let27_0):
                def iife100_(d_1346_dt__update__tmp_h0_):
                    def iife101_(_pat_let28_0):
                        def iife102_(d_1347_dt__update_hcf43_h0_):
                            return D12_DC28(d_1347_dt__update_hcf43_h0_, (d_1346_dt__update__tmp_h0_).cf44, (d_1346_dt__update__tmp_h0_).cf45, (d_1346_dt__update__tmp_h0_).cf46, (d_1346_dt__update__tmp_h0_).cf47)
                        return iife102_(_pat_let28_0)
                    return iife101_((self).f6)
                return iife100_(_pat_let27_0)
            def iife98_(_pat_let26_0):
                def iife103_(d_1348_dt__update__tmp_h1_):
                    def iife104_(_pat_let29_0):
                        def iife105_(d_1349_dt__update_hcf47_h0_):
                            return D12_DC28((d_1348_dt__update__tmp_h1_).cf43, (d_1348_dt__update__tmp_h1_).cf44, (d_1348_dt__update__tmp_h1_).cf45, (d_1348_dt__update__tmp_h1_).cf46, d_1349_dt__update_hcf47_h0_)
                        return iife105_(_pat_let29_0)
                    return iife104_(len(pat_let_tv23_))
                return iife103_(_pat_let26_0)
            source24_ = iife98_(iife99_(D12_DC28((self).f6, ((p0).set(d_1343_i5_, default__.abs((d_1345_v66_).cardinality))).cardinality, p1, True, d_1343_i5_)))
            if source24_.is_DC28:
                d_1350___mcc_h12_ = source24_.cf43
                d_1351___mcc_h13_ = source24_.cf44
                d_1352___mcc_h14_ = source24_.cf45
                d_1353___mcc_h15_ = source24_.cf46
                d_1354___mcc_h16_ = source24_.cf47
                d_1355_cf47_ = d_1354___mcc_h16_
                d_1356_cf46_ = d_1353___mcc_h15_
                d_1357_cf45_ = d_1352___mcc_h14_
                d_1358_cf44_ = d_1351___mcc_h13_
                d_1359_cf43_ = d_1350___mcc_h12_
                d_1360_v67_: D0
                d_1360_v67_ = D0_DC0((879) - (d_1355_cf47_))
                d_1360_v67_ = d_1360_v67_
                d_1342_v64_ = d_1342_v64_
                d_1361_v68_: _dafny.Set
                d_1361_v68_ = _dafny.Set({(self).f6, d_1355_cf47_})
                d_1356_cf46_ = (d_1343_i5_) == ((0) - ((len(d_1361_v68_)) * ((d_1344_v65_)[default__.safeIndex((0) - (len(d_1361_v68_)), len(d_1344_v65_))])))
                d_1362_v69_: _dafny.Map
                d_1362_v69_ = _dafny.Map({d_1355_cf47_: p1})
                d_1363_v71_: _dafny.MultiSet
                def iife106_():
                    coll46_ = _dafny.Map()
                    compr_46_: int
                    for compr_46_ in _dafny.IntegerRange(416, -448):
                        d_1364_v70_: int = compr_46_
                        if ((416) <= (d_1364_v70_)) and ((d_1364_v70_) < (-448)):
                            coll46_[default__.safeModuloInt(d_1364_v70_, (self).f6)] = d_1356_cf46_
                    return _dafny.Map(coll46_)
                d_1363_v71_ = _dafny.MultiSet([(d_1362_v69_ if d_1357_cf45_ else d_1362_v69_), iife106_()
                , d_1362_v69_])
                d_1363_v71_ = d_1363_v71_
            elif True:
                d_1365___mcc_h17_ = source24_.cf42
                d_1366_cf42_ = d_1365___mcc_h17_
                r3 = not ((len(d_1342_v64_)) <= ((self).f6)) or ((self).f4)
                r1 = d_1343_i5_
                d_1367_v72_: _dafny.Array
                nw213_ = _dafny.Array(None, 3)
                nw213_[int(0)] = (not(p1) if not(not(p1)) else (self).f4)
                nw213_[int(1)] = (self).f4
                nw213_[int(2)] = (self).f4
                d_1367_v72_ = nw213_
                index202_ = default__.safeIndex(154, (d_1367_v72_).length(0))
                (d_1367_v72_)[index202_] = p1
                index203_ = default__.safeIndex(154, (d_1367_v72_).length(0))
                rhs187_ = p1
                rhs188_ = p1
                lhs117_ = d_1367_v72_
                lhs118_ = default__.safeIndex(154, (d_1367_v72_).length(0))
                r0 = rhs187_
                lhs117_[lhs118_] = rhs188_
                index204_ = default__.safeIndex(154, (d_1367_v72_).length(0))
                (d_1367_v72_)[index204_] = (d_1367_v72_)[default__.safeIndex(154, (d_1367_v72_).length(0))]
            d_1368_v73_: _dafny.Map
            d_1368_v73_ = _dafny.Map({(self).f4: len(d_1342_v64_)})
            d_1369_v74_: D10
            d_1369_v74_ = D10_DC21(d_1368_v73_)
            d_1370_v75_: _dafny.Map
            d_1370_v75_ = _dafny.Map({d_1369_v74_: len(d_1342_v64_)})
            d_1370_v75_ = default__.fm59(globalState)
            r1 = (0) - ((self).f6)
        d_1371_v76_: _dafny.Seq
        d_1371_v76_ = _dafny.SeqWithoutIsStrInference([(self).f6])
        d_1371_v76_ = (_dafny.SeqWithoutIsStrInference([-887, (self).f6, (self).f6, (self).f6, (self).f6]) if p1 else _dafny.SeqWithoutIsStrInference([(self).f6 for d_1372_i6_ in range(default__.abs(899))]))
        d_1373_v77_: _dafny.Map
        d_1373_v77_ = _dafny.Map({(self).f4: ((self).f6) - ((self).f6)})
        d_1373_v77_ = (default__.fm44(globalState)) | ((d_1373_v77_) | (d_1373_v77_))
        r0 = True
        r1 = (self).f6
        r2 = (self).f6
        r3 = not((self).f4)
        return r0, r1, r2, r3


class C4(T1, T0):
    def  __init__(self):
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self._f6: int = int(0)
        self._f4: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f5(self):
        return self._f5
    @f5.setter
    def f5(self, value):
        self._f5 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f4(self):
        return self._f4
    def ctor__(self, f5, f6, f4):
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f4 = f4

    def fm8(self, p0, p1, p2, p3, globalState):
        return not((self).f4)

    def fm9(self, p0, p1, p2, globalState):
        return (self).f6

    def fm6(self, p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([default__.safeModuloInt(117, 842), ((self).f6 if (self).f4 else (self).f6), (self).f6, default__.safeDivisionInt((self).f6, (self).f6), (self).f6])

    def fm7(self, globalState):
        def lambda42_(source25_):
            if source25_.is_DC20:
                d_1374___mcc_h0_ = source25_.cf26
                d_1375___mcc_h1_ = source25_.cf27
                d_1376___mcc_h2_ = source25_.cf28
                d_1377_cf28_ = d_1376___mcc_h2_
                d_1378_cf27_ = d_1375___mcc_h1_
                d_1379_cf26_ = d_1374___mcc_h0_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uaimtxsem"))) + (_dafny.SeqWithoutIsStrInference([d_1378_cf27_ for d_1380_i0_ in range(default__.abs(744))]))
            elif True:
                d_1381___mcc_h3_ = source25_.cf25
                d_1382_cf25_ = d_1381___mcc_h3_
                return d_1382_cf25_

        return len(lambda42_(D9_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxfde")))))

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_1383_v0_: _dafny.Seq
        d_1383_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tv"))
        d_1384_v1_: _dafny.MultiSet
        d_1384_v1_ = _dafny.MultiSet([len(_dafny.Map({(self).f4: (self).f4})), len(_dafny.SeqWithoutIsStrInference([d_1383_v0_]))])
        d_1385_v2_: bool
        d_1386_v3_: int
        out45_: bool
        out46_: int
        out45_, out46_ = default__.m0(_dafny.Map({d_1384_v1_: d_1383_v0_}), not(((self).f4 if (self).f4 else (self).f4)), d_1384_v1_, (self).f4, globalState)
        d_1385_v2_ = out45_
        d_1386_v3_ = out46_
        d_1387_v4_: _dafny.Array
        def lambda43_(d_1388_i0_):
            return _dafny.CodePoint('x')

        init28_ = lambda43_
        nw214_ = _dafny.Array(None, 9)
        for i0_28_ in range(nw214_.length(0)):
            nw214_[i0_28_] = init28_(i0_28_)
        d_1387_v4_ = nw214_
        d_1389_v5_: str
        d_1389_v5_ = _dafny.CodePoint('i')
        index205_ = default__.safeIndex(290, (d_1387_v4_).length(0))
        (d_1387_v4_)[index205_] = d_1389_v5_
        d_1390_v6_: _dafny.Seq
        d_1390_v6_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(d_1386_v3_, (self).f6)])
        d_1391_v7_: _dafny.Map
        d_1391_v7_ = _dafny.Map({d_1385_v2_: (0) - (len(default__.fm30(d_1390_v6_, True, globalState)))})
        d_1392_v8_: _dafny.Map
        d_1392_v8_ = _dafny.Map({d_1385_v2_: (self).f4})
        d_1393_v9_: _dafny.Map
        d_1393_v9_ = _dafny.Map({d_1392_v8_: len(_dafny.Set({(self).f4, d_1385_v2_}))})
        d_1394_v11_: _dafny.Array
        nw215_ = _dafny.Array(None, 22)
        nw215_[int(0)] = -811
        nw215_[int(1)] = p0
        nw215_[int(2)] = len(_dafny.SeqWithoutIsStrInference([d_1389_v5_ for d_1395_i1_ in range(default__.abs(-476))]))
        nw215_[int(3)] = (self).f6
        nw215_[int(4)] = (self).f6
        nw215_[int(5)] = (self).f6
        nw215_[int(6)] = len(d_1391_v7_)
        nw215_[int(7)] = (self).f6
        nw215_[int(8)] = (self).f6
        nw215_[int(9)] = d_1386_v3_
        nw215_[int(10)] = len(_dafny.SeqWithoutIsStrInference([p0, p0, (0) - (d_1386_v3_), (self).f6]))
        nw215_[int(11)] = (self).f6
        nw215_[int(12)] = p0
        nw215_[int(13)] = (self).f6
        nw215_[int(14)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejh")))
        nw215_[int(15)] = default__.fm3(d_1386_v3_, (self).f4, globalState)
        nw215_[int(16)] = d_1386_v3_
        nw215_[int(17)] = d_1386_v3_
        def iife107_():
            coll47_ = _dafny.Map()
            compr_47_: int
            for compr_47_ in _dafny.IntegerRange(374, 306):
                d_1396_v10_: int = compr_47_
                if ((374) <= (d_1396_v10_)) and ((d_1396_v10_) < (306)):
                    coll47_[(d_1396_v10_) - ((self).f6)] = len(_dafny.Map({(_dafny.MultiSet(d_1390_v6_)).cardinality: (self).f4}))
            return _dafny.Map(coll47_)
        nw215_[int(18)] = ((d_1393_v9_)[_dafny.Map({d_1385_v2_: (self).f4})] if (_dafny.Map({d_1385_v2_: (self).f4})) in (d_1393_v9_) else len(iife107_()
        ))
        nw215_[int(19)] = p0
        nw215_[int(20)] = d_1386_v3_
        nw215_[int(21)] = (self).f6
        d_1394_v11_ = nw215_
        d_1397_v12_: _dafny.Seq
        d_1397_v12_ = _dafny.SeqWithoutIsStrInference([d_1394_v11_])
        d_1398_v13_: _dafny.MultiSet
        d_1398_v13_ = _dafny.MultiSet([d_1394_v11_, (d_1397_v12_)[default__.safeIndex(d_1386_v3_, len(d_1397_v12_))]])
        d_1399_v14_: _dafny.MultiSet
        d_1399_v14_ = _dafny.MultiSet([d_1398_v13_, (d_1398_v13_) - ((d_1398_v13_).set(d_1394_v11_, default__.abs(p0)))])
        index206_ = default__.safeIndex(290, (d_1387_v4_).length(0))
        rhs189_ = d_1389_v5_
        rhs190_ = (self).f6
        rhs191_ = ((d_1399_v14_)[(d_1398_v13_) | (d_1398_v13_)] if ((d_1398_v13_) | (d_1398_v13_)) in (d_1399_v14_) else (d_1386_v3_ if d_1385_v2_ else p0))
        rhs192_ = d_1390_v6_
        lhs119_ = d_1387_v4_
        lhs120_ = default__.safeIndex(290, (d_1387_v4_).length(0))
        lhs119_[lhs120_] = rhs189_
        d_1386_v3_ = rhs190_
        d_1386_v3_ = rhs191_
        d_1390_v6_ = rhs192_
        d_1400_v15_: C0
        nw216_ = C0()
        nw216_.ctor__(default__.safeDivisionInt(d_1386_v3_, len(d_1390_v6_)))
        d_1400_v15_ = nw216_
        d_1401_v16_: _dafny.Map
        d_1401_v16_ = _dafny.Map({d_1389_v5_: d_1385_v2_})
        d_1401_v16_ = (d_1401_v16_).set((d_1387_v4_)[default__.safeIndex(290, (d_1387_v4_).length(0))], (self).f4)
        hi8_ = (0) - (((d_1400_v15_).f12) - (281))
        for d_1402_i2_ in range(-659, hi8_):
            d_1403_v17_: C0
            nw217_ = C0()
            nw217_.ctor__(d_1386_v3_)
            d_1403_v17_ = nw217_
            d_1404_v18_: _dafny.Map
            d_1404_v18_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jl"))): default__.fm31(d_1402_i2_, d_1386_v3_, len(d_1383_v0_), (self).f4, globalState)})
            d_1405_v19_: _dafny.Map
            d_1405_v19_ = _dafny.Map({d_1386_v3_: (self).f4})
            d_1406_v20_: _dafny.Set
            d_1406_v20_ = _dafny.Set({(d_1384_v1_).cardinality, len(d_1391_v7_), (D0_DC0((self).f6)).cf0, len(d_1405_v19_), len(d_1390_v6_)})
            d_1407_v23_: _dafny.Array
            nw218_ = _dafny.Array(None, 8)
            nw218_[int(0)] = ((d_1404_v18_)[(d_1400_v15_).f12] if ((d_1400_v15_).f12) in (d_1404_v18_) else d_1406_v20_)
            def iife108_():
                coll48_ = _dafny.Set()
                compr_48_: int
                for compr_48_ in _dafny.IntegerRange(380, 684):
                    d_1408_v21_: int = compr_48_
                    if ((380) <= (d_1408_v21_)) and ((d_1408_v21_) < (684)):
                        coll48_ = coll48_.union(_dafny.Set([(d_1408_v21_) - ((d_1400_v15_).f12)]))
                return _dafny.Set(coll48_)
            nw218_[int(1)] = iife108_()
            
            nw218_[int(2)] = _dafny.Set({(self).f6, (d_1400_v15_).f12, (d_1403_v17_).f12, (d_1384_v1_).cardinality})
            nw218_[int(3)] = (default__.fm32(558, d_1385_v2_, globalState))
            nw218_[int(4)] = d_1406_v20_
            nw218_[int(5)] = _dafny.Set({d_1402_i2_})
            nw218_[int(6)] = d_1406_v20_
            def iife109_():
                coll49_ = _dafny.Set()
                compr_49_: int
                for compr_49_ in _dafny.IntegerRange(788, 145):
                    d_1409_v22_: int = compr_49_
                    if ((788) <= (d_1409_v22_)) and ((d_1409_v22_) < (145)):
                        coll49_ = coll49_.union(_dafny.Set([(d_1409_v22_) - ((d_1403_v17_).f12)]))
                return _dafny.Set(coll49_)
            nw218_[int(7)] = iife109_()
            
            d_1407_v23_ = nw218_
            d_1410_v24_: _dafny.Set
            d_1410_v24_ = _dafny.Set({(d_1403_v17_).fm17((self).f4, globalState), not(not((self).f4)), d_1385_v2_, (self).f4, not((self).f4)})
            d_1411_v25_: _dafny.Map
            d_1411_v25_ = _dafny.Map({d_1410_v24_: d_1385_v2_})
            d_1412_v26_: _dafny.Map
            d_1412_v26_ = _dafny.Map({(d_1407_v23_ if d_1385_v2_ else d_1407_v23_): len(d_1411_v25_)})
            d_1412_v26_ = (d_1412_v26_).set(d_1407_v23_, default__.safeModuloInt(p0, (d_1400_v15_).f12))
            index207_ = default__.safeIndex(911, (d_1394_v11_).length(0))
            (d_1394_v11_)[index207_] = len(d_1391_v7_)
            d_1413_v27_: _dafny.Seq
            d_1413_v27_ = _dafny.SeqWithoutIsStrInference([(self).f4, d_1385_v2_, d_1385_v2_, d_1385_v2_])
            index208_ = default__.safeIndex(911, (d_1394_v11_).length(0))
            (d_1394_v11_)[index208_] = len(((d_1413_v27_).set(default__.safeIndex(len(d_1413_v27_), len(d_1413_v27_)), (self).f4)) + (d_1413_v27_))
            d_1414_v28_: _dafny.MultiSet
            d_1414_v28_ = _dafny.MultiSet([(d_1400_v15_).fm17((self).f4, globalState), (d_1400_v15_).fm17(True, globalState), d_1385_v2_])
            d_1414_v28_ = d_1414_v28_
        index209_ = default__.safeIndex(290, (d_1387_v4_).length(0))
        (d_1387_v4_)[index209_] = (d_1387_v4_)[default__.safeIndex(290, (d_1387_v4_).length(0))]
        r0 = d_1391_v7_
        return r0

    def m2(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_1415_v0_: _dafny.Map
        d_1415_v0_ = _dafny.Map({((0) - ((self).f6)) * ((self).f6): (self).fm7(globalState)})
        d_1416_v1_: _dafny.Array
        nw219_ = _dafny.Array(int(0), 4)
        d_1416_v1_ = nw219_
        d_1417_v2_: _dafny.Map
        d_1417_v2_ = _dafny.Map({d_1416_v1_: _dafny.Map({(self).f6: (0) - ((self).f6)})})
        d_1415_v0_ = ((d_1417_v2_)[d_1416_v1_] if (d_1416_v1_) in (d_1417_v2_) else _dafny.Map({(self).f6: (self).f6}))
        d_1418_v3_: _dafny.Seq
        d_1418_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjatmipct"))
        d_1419_v4_: _dafny.Seq
        d_1419_v4_ = _dafny.SeqWithoutIsStrInference([len(d_1418_v3_), (self).f6])
        d_1419_v4_ = _dafny.SeqWithoutIsStrInference([(self).f6 for d_1420_i0_ in range(default__.abs(355))])
        hi9_ = len(d_1418_v3_)
        for d_1421_i1_ in range((self).f6, hi9_):
            d_1422_v5_: _dafny.Seq
            d_1422_v5_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1423_v6_: D4
            d_1423_v6_ = D4_DC13(default__.safeModuloInt(len(d_1415_v0_), d_1421_i1_), (self).f4, -262, d_1422_v5_)
            source26_ = d_1423_v6_
            if source26_.is_DC13:
                d_1424___mcc_h0_ = source26_.cf14
                d_1425___mcc_h1_ = source26_.cf15
                d_1426___mcc_h2_ = source26_.cf16
                d_1427___mcc_h3_ = source26_.cf17
                d_1428_cf17_ = d_1427___mcc_h3_
                d_1429_cf16_ = d_1426___mcc_h2_
                d_1430_cf15_ = d_1425___mcc_h1_
                d_1431_cf14_ = d_1424___mcc_h0_
                d_1432_v7_: str
                d_1432_v7_ = _dafny.CodePoint('y')
                d_1429_cf16_ = (d_1431_cf14_ if (d_1432_v7_) not in (d_1418_v3_) else d_1431_cf14_)
                d_1433_v8_: C0
                nw220_ = C0()
                nw220_.ctor__((self).f6)
                d_1433_v8_ = nw220_
                d_1434_v9_: _dafny.Map
                d_1434_v9_ = _dafny.Map({d_1433_v8_: d_1421_i1_})
                d_1435_v10_: _dafny.Map
                d_1435_v10_ = _dafny.Map({d_1430_cf15_: d_1418_v3_})
                d_1436_v11_: _dafny.Array
                nw221_ = _dafny.Array(None, 22)
                nw221_[int(0)] = (self).f4
                nw221_[int(1)] = (self).f4
                nw221_[int(2)] = p0
                nw221_[int(3)] = p0
                nw221_[int(4)] = (self).f4
                nw221_[int(5)] = (self).f4
                nw221_[int(6)] = d_1430_cf15_
                nw221_[int(7)] = p0
                nw221_[int(8)] = p0
                nw221_[int(9)] = (self).f4
                nw221_[int(10)] = not(not(d_1430_cf15_))
                nw221_[int(11)] = (self).f4
                nw221_[int(12)] = (self).f4
                nw221_[int(13)] = d_1430_cf15_
                nw221_[int(14)] = not(p0)
                nw221_[int(15)] = p0
                nw221_[int(16)] = (self).f4
                nw221_[int(17)] = True
                nw221_[int(18)] = d_1430_cf15_
                nw221_[int(19)] = (self).fm8(d_1429_cf16_, d_1430_cf15_, ((d_1435_v10_)[p0] if (p0) in (d_1435_v10_) else d_1418_v3_), d_1419_v4_, globalState)
                nw221_[int(20)] = d_1430_cf15_
                nw221_[int(21)] = (self).f4
                d_1436_v11_ = nw221_
                d_1415_v0_ = (d_1415_v0_).set((len(d_1434_v9_)) - (450), (D10_DC22(d_1418_v3_, p0, False, d_1436_v11_, (d_1433_v8_).f12)).cf34)
                index210_ = default__.safeIndex(991, (d_1416_v1_).length(0))
                (d_1416_v1_)[index210_] = 784
                index211_ = default__.safeIndex(991, (d_1416_v1_).length(0))
                (d_1416_v1_)[index211_] = (d_1421_i1_) * (925)
                d_1437_v12_: D0
                d_1437_v12_ = D0_DC2(d_1429_cf16_)
                d_1438_v13_: _dafny.MultiSet
                d_1438_v13_ = _dafny.MultiSet([d_1432_v7_])
                d_1439_v14_: T1
                nw222_ = C3()
                nw222_.ctor__(self.f5, (d_1438_v13_).cardinality, default__.fm0(_dafny.CodePoint('f'), globalState))
                d_1439_v14_ = nw222_
                d_1440_v15_: _dafny.Seq
                d_1440_v15_ = _dafny.SeqWithoutIsStrInference([d_1439_v14_, d_1439_v14_, d_1439_v14_, d_1439_v14_, d_1439_v14_])
                pat_let_tv24_ = d_1422_v5_
                pat_let_tv25_ = d_1429_cf16_
                pat_let_tv26_ = d_1429_cf16_
                pat_let_tv27_ = d_1416_v1_
                pat_let_tv28_ = d_1416_v1_
                d_1441_v16_: _dafny.Array
                nw223_ = _dafny.Array(None, 21)
                nw223_[int(0)] = d_1437_v12_
                nw223_[int(1)] = d_1437_v12_
                def iife110_(_pat_let30_0):
                    def iife111_(d_1442_dt__update__tmp_h0_):
                        def iife112_(_pat_let31_0):
                            def iife113_(d_1443_dt__update_hcf1_h0_):
                                return D0_DC2(d_1443_dt__update_hcf1_h0_)
                            return iife113_(_pat_let31_0)
                        return iife112_(len(pat_let_tv24_))
                    return iife111_(_pat_let30_0)
                nw223_[int(2)] = iife110_(D0_DC2(d_1429_cf16_))
                nw223_[int(3)] = d_1437_v12_
                nw223_[int(4)] = d_1437_v12_
                nw223_[int(5)] = d_1437_v12_
                nw223_[int(6)] = d_1437_v12_
                def iife114_(_pat_let32_0):
                    def iife115_(d_1444_dt__update__tmp_h1_):
                        def iife116_(_pat_let33_0):
                            def iife117_(d_1445_dt__update_hcf1_h1_):
                                return D0_DC2(d_1445_dt__update_hcf1_h1_)
                            return iife117_(_pat_let33_0)
                        return iife116_(pat_let_tv25_)
                    return iife115_(_pat_let32_0)
                nw223_[int(7)] = iife114_(D0_DC2(291))
                nw223_[int(8)] = D0_DC2(431)
                def iife118_(_pat_let34_0):
                    def iife119_(d_1446_dt__update__tmp_h2_):
                        def iife120_(_pat_let35_0):
                            def iife121_(d_1447_dt__update_hcf1_h2_):
                                return D0_DC2(d_1447_dt__update_hcf1_h2_)
                            return iife121_(_pat_let35_0)
                        return iife120_(pat_let_tv26_)
                    return iife119_(_pat_let34_0)
                nw223_[int(9)] = iife118_(d_1437_v12_)
                nw223_[int(10)] = d_1437_v12_
                nw223_[int(11)] = d_1437_v12_
                def iife122_(_pat_let36_0):
                    def iife123_(d_1448_dt__update__tmp_h3_):
                        def iife124_(_pat_let37_0):
                            def iife125_(d_1449_dt__update_hcf1_h3_):
                                return D0_DC2(d_1449_dt__update_hcf1_h3_)
                            return iife125_(_pat_let37_0)
                        return iife124_((pat_let_tv28_)[default__.safeIndex(991, (pat_let_tv27_).length(0))])
                    return iife123_(_pat_let36_0)
                nw223_[int(12)] = iife122_(d_1437_v12_)
                nw223_[int(13)] = d_1437_v12_
                nw223_[int(14)] = d_1437_v12_
                nw223_[int(15)] = D0_DC2(len(d_1440_v15_))
                nw223_[int(16)] = d_1437_v12_
                nw223_[int(17)] = D0_DC2(d_1421_i1_)
                nw223_[int(18)] = d_1437_v12_
                nw223_[int(19)] = d_1437_v12_
                nw223_[int(20)] = D0_DC2((d_1433_v8_).f12)
                d_1441_v16_ = nw223_
                def lambda44_(d_1450_v12_):
                    def lambda45_(d_1451_i2_):
                        return d_1450_v12_

                    return lambda45_

                init29_ = lambda44_(d_1437_v12_)
                nw224_ = _dafny.Array(None, 22)
                for i0_29_ in range(nw224_.length(0)):
                    nw224_[i0_29_] = init29_(i0_29_)
                d_1441_v16_ = nw224_
            elif True:
                d_1452___mcc_h4_ = source26_.cf13
                d_1453_cf13_ = d_1452___mcc_h4_
                d_1454_v17_: _dafny.Map
                d_1454_v17_ = _dafny.Map({(self).f6: d_1416_v1_})
                d_1454_v17_ = (d_1454_v17_).set(d_1421_i1_, d_1416_v1_)
                d_1455_v18_: int
                d_1455_v18_ = 277
                d_1456_v20_: _dafny.Map
                d_1456_v20_ = _dafny.Map({(0) - (d_1421_i1_): d_1453_cf13_})
                d_1457_v21_: _dafny.Set
                d_1457_v21_ = _dafny.Set({d_1456_v20_})
                def iife126_():
                    coll50_ = _dafny.Map()
                    compr_50_: int
                    for compr_50_ in _dafny.IntegerRange(447, -282):
                        d_1458_v19_: int = compr_50_
                        if ((447) <= (d_1458_v19_)) and ((d_1458_v19_) < (-282)):
                            coll50_[default__.safeDivisionInt(d_1458_v19_, 592)] = d_1421_i1_
                    return _dafny.Map(coll50_)
                d_1455_v18_ = default__.safeDivisionInt(len(iife126_()
                ), default__.safeDivisionInt(len(d_1457_v21_), (self).f6))
                d_1459_v22_: _dafny.Array
                nw225_ = _dafny.Array(_dafny.Seq({}), 29)
                d_1459_v22_ = nw225_
                index212_ = default__.safeIndex(528, (d_1459_v22_).length(0))
                (d_1459_v22_)[index212_] = d_1419_v4_
                index213_ = default__.safeIndex(528, (d_1459_v22_).length(0))
                (d_1459_v22_)[index213_] = _dafny.SeqWithoutIsStrInference([d_1455_v18_])
                d_1460_v23_: _dafny.Array
                nw226_ = _dafny.Array(_dafny.CodePoint('D'), 21)
                d_1460_v23_ = nw226_
                index214_ = default__.safeIndex(885, (d_1460_v23_).length(0))
                (d_1460_v23_)[index214_] = d_1453_cf13_
                index215_ = default__.safeIndex(885, (d_1460_v23_).length(0))
                (d_1460_v23_)[index215_] = d_1453_cf13_
            d_1461_v24_: bool
            d_1461_v24_ = True
            d_1462_v25_: _dafny.MultiSet
            d_1462_v25_ = _dafny.MultiSet([d_1421_i1_, (self).f6])
            d_1463_v26_: _dafny.Set
            d_1463_v26_ = _dafny.Set({p0, (self).f4})
            d_1464_v27_: _dafny.Seq
            d_1464_v27_ = _dafny.SeqWithoutIsStrInference([d_1419_v4_, _dafny.SeqWithoutIsStrInference([d_1421_i1_ for d_1465_i3_ in range(default__.abs(980))])])
            d_1461_v24_ = (self).fm8(((d_1462_v25_)[len(d_1463_v26_)] if (len(d_1463_v26_)) in (d_1462_v25_) else 622), (d_1464_v27_) == (_dafny.SeqWithoutIsStrInference([d_1419_v4_])), d_1418_v3_, (d_1419_v4_) + (d_1419_v4_), globalState)
            d_1466_v28_: D11
            d_1466_v28_ = D11_DC24((self).f4, (self).f4, len(d_1418_v3_))
            d_1467_v29_: D20
            d_1467_v29_ = D20_DC41((d_1421_i1_ if d_1461_v24_ else (d_1466_v28_).cf38))
            d_1467_v29_ = D20_DC41(130)
            d_1468_v30_: int
            d_1468_v30_ = 174
            d_1469_v31_: _dafny.Set
            d_1469_v31_ = _dafny.Set({d_1421_i1_})
            d_1470_v32_: _dafny.Set
            d_1470_v32_ = d_1469_v31_
            d_1471_v33_: _dafny.Set
            d_1471_v33_ = _dafny.Set({d_1470_v32_})
            d_1468_v30_ = ((0) - (((0) - (d_1468_v30_)) - ((self).f6))) + ((d_1468_v30_ if (self).f4 else len(d_1471_v33_)))
        d_1472_v34_: bool
        d_1472_v34_ = False
        index216_ = default__.safeIndex(485, (d_1416_v1_).length(0))
        (d_1416_v1_)[index216_] = ((self).f6) * ((self).f6)
        d_1473_v35_: int
        d_1473_v35_ = 978
        d_1474_v36_: _dafny.Map
        d_1474_v36_ = _dafny.Map({(d_1419_v4_)[default__.safeIndex(d_1473_v35_, len(d_1419_v4_))]: d_1472_v34_})
        d_1475_v37_: _dafny.Map
        d_1475_v37_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f4]): d_1474_v36_})
        d_1476_v38_: _dafny.Seq
        d_1476_v38_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f6])])
        d_1477_v39_: _dafny.MultiSet
        d_1477_v39_ = _dafny.MultiSet([p0])
        index217_ = default__.safeIndex(485, (d_1416_v1_).length(0))
        rhs193_ = (d_1475_v37_) != (default__.fm60(len(d_1476_v38_), (self).f6, d_1473_v35_, globalState))
        rhs194_ = d_1473_v35_
        rhs195_ = default__.safeDivisionInt((d_1477_v39_).cardinality, d_1473_v35_)
        rhs196_ = not(not(d_1472_v34_))
        rhs197_ = (self).f4
        lhs121_ = d_1416_v1_
        lhs122_ = default__.safeIndex(485, (d_1416_v1_).length(0))
        d_1472_v34_ = rhs193_
        lhs121_[lhs122_] = rhs194_
        d_1473_v35_ = rhs195_
        d_1472_v34_ = rhs196_
        d_1472_v34_ = rhs197_
        d_1478_v40_: _dafny.Array
        def lambda46_(d_1479_p0_):
            def lambda47_(d_1480_i4_):
                return D12_DC27(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([d_1479_p0_])]))

            return lambda47_

        init30_ = lambda46_(p0)
        nw227_ = _dafny.Array(None, 28)
        for i0_30_ in range(nw227_.length(0)):
            nw227_[i0_30_] = init30_(i0_30_)
        d_1478_v40_ = nw227_
        d_1481_v41_: _dafny.Set
        d_1481_v41_ = _dafny.Set({(self).f4, (self).f4})
        d_1482_v42_: D12
        d_1482_v42_ = D12_DC27(default__.fm61(d_1481_v41_, globalState))
        index218_ = default__.safeIndex(785, (d_1478_v40_).length(0))
        (d_1478_v40_)[index218_] = d_1482_v42_
        index219_ = default__.safeIndex(785, (d_1478_v40_).length(0))
        rhs198_ = d_1473_v35_
        rhs199_ = d_1482_v42_
        lhs123_ = d_1478_v40_
        lhs124_ = default__.safeIndex(785, (d_1478_v40_).length(0))
        d_1473_v35_ = rhs198_
        lhs123_[lhs124_] = rhs199_
        d_1483_v43_: C1
        nw228_ = C1()
        nw228_.ctor__(d_1473_v35_, default__.safeModuloInt((d_1416_v1_)[default__.safeIndex(485, (d_1416_v1_).length(0))], d_1473_v35_), self.f5, (self).f6, (self).fm8((self).f6, (self).f4, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1484_i5_ in range(default__.abs(566))]), _dafny.SeqWithoutIsStrInference([d_1473_v35_ for d_1485_i6_ in range(default__.abs(-532))]), globalState))
        d_1483_v43_ = nw228_
        d_1486_v44_: str
        d_1486_v44_ = _dafny.CodePoint('f')
        d_1487_v45_: D9
        d_1487_v45_ = D9_DC20((self).f4, d_1486_v44_, d_1483_v43_.f16)
        d_1488_v46_: _dafny.Seq
        d_1488_v46_ = _dafny.SeqWithoutIsStrInference([d_1487_v45_])
        index220_ = default__.safeIndex(485, (d_1416_v1_).length(0))
        def lambda48_(source27_):
            d_1489___mcc_h5_ = source27_
            d_1490_cf57_ = d_1489___mcc_h5_
            return 911

        rhs200_ = (d_1483_v43_ if (((d_1415_v0_)[(d_1416_v1_)[default__.safeIndex(485, (d_1416_v1_).length(0))]] if ((d_1416_v1_)[default__.safeIndex(485, (d_1416_v1_).length(0))]) in (d_1415_v0_) else (d_1416_v1_)[default__.safeIndex(485, (d_1416_v1_).length(0))])) <= ((d_1416_v1_)[default__.safeIndex(485, (d_1416_v1_).length(0))]) else d_1483_v43_)
        rhs201_ = p0
        rhs202_ = p0
        rhs203_ = (0) - (lambda48_(d_1488_v46_))
        lhs125_ = d_1416_v1_
        lhs126_ = default__.safeIndex(485, (d_1416_v1_).length(0))
        d_1483_v43_ = rhs200_
        d_1472_v34_ = rhs201_
        d_1472_v34_ = rhs202_
        lhs125_[lhs126_] = rhs203_
        r0 = d_1486_v44_
        d_1491_v47_: _dafny.Array
        nw229_ = _dafny.Array(D1.default()(), 2)
        d_1491_v47_ = nw229_
        r1 = d_1491_v47_
        return r0, r1


class C5(T1, T0):
    def  __init__(self):
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self._f6: int = int(0)
        self._f4: bool = False
        self.f14: int = int(0)
        self._f13: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f5(self):
        return self._f5
    @f5.setter
    def f5(self, value):
        self._f5 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f4(self):
        return self._f4
    def ctor__(self, f13, f14, f5, f6, f4):
        (self)._f13 = f13
        (self).f14 = f14
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f4 = f4

    def fm8(self, p0, p1, p2, p3, globalState):
        if (self).f13:
            return (_dafny.Map({(self).f4: (self).f4})) == (_dafny.Map({(self).f13: (self).f13}))
        elif True:
            return ((0) - (self.f14)) not in (_dafny.MultiSet([self.f14]))

    def fm9(self, p0, p1, p2, globalState):
        return 30

    def fm6(self, p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([(self).f6, 309, (self).f6])) - (_dafny.MultiSet([(0) - (self.f14)]))

    def fm7(self, globalState):
        source28_ = D3_DC9((self).f13, (self).f6, (self).f4)
        if source28_.is_DC9:
            d_1492___mcc_h0_ = source28_.cf8
            d_1493___mcc_h1_ = source28_.cf9
            d_1494___mcc_h2_ = source28_.cf10
            d_1495_cf10_ = d_1494___mcc_h2_
            d_1496_cf9_ = d_1493___mcc_h1_
            d_1497_cf8_ = d_1492___mcc_h0_
            return (self).f6
        elif source28_.is_DC10:
            d_1498___mcc_h3_ = source28_.cf11
            d_1499_cf11_ = d_1498___mcc_h3_
            return (self.f14) + (self.f14)
        elif source28_.is_DC8:
            d_1500___mcc_h4_ = source28_.cf7
            d_1501_cf7_ = d_1500___mcc_h4_
            return (64) - (len((D9_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fafdlayl")))).cf25))
        elif True:
            d_1502___mcc_h5_ = source28_.cf12
            d_1503_cf12_ = d_1502___mcc_h5_
            return (self).f6

    def fm26(self, p0, p1, p2, globalState):
        return (0) - (self.f14)

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_1504_v0_: _dafny.MultiSet
        d_1504_v0_ = _dafny.MultiSet([(self).f13, (self).f4, (self).f4, (self).f13, (self).f13])
        d_1505_v1_: _dafny.Map
        d_1505_v1_ = _dafny.Map({D6_DC15(default__.fm27((self).f13, 35, d_1504_v0_, (self).f13, globalState)): ((d_1504_v0_).intersection(d_1504_v0_)).cardinality})
        d_1505_v1_ = d_1505_v1_
        (self).f14 = (self).f6
        d_1506_v2_: _dafny.Set
        d_1506_v2_ = _dafny.Set({(self).f4})
        d_1507_v3_: _dafny.Seq
        d_1507_v3_ = _dafny.SeqWithoutIsStrInference([d_1506_v2_])
        d_1508_v4_: _dafny.Map
        d_1508_v4_ = _dafny.Map({(self).f13: len(d_1507_v3_)})
        d_1509_v5_: str
        d_1509_v5_ = _dafny.CodePoint('a')
        d_1510_v6_: D9
        d_1510_v6_ = D9_DC20((self).f13, d_1509_v5_, (self).f6)
        d_1511_v7_: C0
        nw230_ = C0()
        nw230_.ctor__(((d_1508_v4_)[(d_1510_v6_).cf26] if ((d_1510_v6_).cf26) in (d_1508_v4_) else p0))
        d_1511_v7_ = nw230_
        d_1512_v8_: _dafny.Seq
        d_1512_v8_ = _dafny.SeqWithoutIsStrInference([(self).f13])
        d_1513_v9_: _dafny.Seq
        d_1513_v9_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f4, (self).f4, (self).f13, (self).f4, (self).f4])])
        d_1514_v11_: _dafny.Seq
        d_1514_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trc"))
        d_1515_v12_: _dafny.Array
        def lambda49_(d_1516_i0_):
            return default__.safeModuloInt(d_1516_i0_, (self).f6)

        init31_ = lambda49_
        nw231_ = _dafny.Array(None, 1)
        for i0_31_ in range(nw231_.length(0)):
            nw231_[i0_31_] = init31_(i0_31_)
        d_1515_v12_ = nw231_
        d_1517_v13_: _dafny.Map
        d_1517_v13_ = _dafny.Map({(self).f4: d_1515_v12_})
        d_1518_v14_: _dafny.Array
        nw232_ = _dafny.Array(None, 20)
        nw232_[int(0)] = self.f14
        nw232_[int(1)] = (0) - (self.f14)
        nw232_[int(2)] = len(_dafny.SeqWithoutIsStrInference([59, (d_1511_v7_).f12, (d_1511_v7_).f12]))
        nw232_[int(3)] = (self.f14) - (self.f14)
        nw232_[int(4)] = (d_1511_v7_).f12
        nw232_[int(5)] = (self).f6
        nw232_[int(6)] = p0
        nw232_[int(7)] = (self).f6
        nw232_[int(8)] = len(_dafny.Map({(_dafny.MultiSet(d_1512_v8_)).cardinality: (d_1511_v7_).f12}))
        nw232_[int(9)] = (self).f6
        nw232_[int(10)] = self.f14
        nw232_[int(11)] = (d_1511_v7_).f12
        nw232_[int(12)] = len(d_1513_v9_)
        def iife127_():
            coll51_ = _dafny.Set()
            compr_51_: int
            for compr_51_ in _dafny.IntegerRange(253, -291):
                d_1519_v10_: int = compr_51_
                if ((253) <= (d_1519_v10_)) and ((d_1519_v10_) < (-291)):
                    coll51_ = coll51_.union(_dafny.Set([(d_1519_v10_) * ((d_1511_v7_).f12)]))
            return _dafny.Set(coll51_)
        nw232_[int(13)] = len(iife127_()
        )
        nw232_[int(14)] = ((self).f6) - (len(d_1514_v11_))
        nw232_[int(15)] = len(d_1517_v13_)
        nw232_[int(16)] = (self).f6
        nw232_[int(17)] = default__.safeModuloInt(438, self.f14)
        nw232_[int(18)] = (self).f6
        nw232_[int(19)] = (d_1511_v7_).f12
        d_1518_v14_ = nw232_
        index221_ = default__.safeIndex(396, (d_1515_v12_).length(0))
        (d_1515_v12_)[index221_] = self.f14
        d_1520_v15_: D10
        d_1520_v15_ = D10_DC21(default__.fm28(globalState))
        d_1521_v16_: _dafny.Seq
        d_1521_v16_ = _dafny.SeqWithoutIsStrInference([d_1507_v3_, d_1507_v3_])
        index222_ = default__.safeIndex(396, (d_1515_v12_).length(0))
        rhs204_ = ((d_1511_v7_).f12 if (self).f13 else (d_1511_v7_).f12)
        rhs205_ = (d_1520_v15_).cf29
        rhs206_ = (_dafny.SeqWithoutIsStrInference([d_1506_v2_]) if (self.f14) != (521) else (d_1521_v16_)[default__.safeIndex((d_1511_v7_).f12, len(d_1521_v16_))])
        lhs127_ = d_1515_v12_
        lhs128_ = default__.safeIndex(396, (d_1515_v12_).length(0))
        lhs127_[lhs128_] = rhs204_
        d_1508_v4_ = rhs205_
        d_1507_v3_ = rhs206_
        d_1522_v17_: _dafny.Map
        d_1522_v17_ = _dafny.Map({d_1509_v5_: d_1509_v5_})
        d_1522_v17_ = (d_1522_v17_).set(_dafny.CodePoint('w'), d_1509_v5_)
        d_1523_v18_: _dafny.MultiSet
        d_1523_v18_ = _dafny.MultiSet([len(d_1514_v11_), len(d_1506_v2_)])
        d_1524_v19_: _dafny.Map
        d_1524_v19_ = _dafny.Map({d_1523_v18_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydutpac"))})
        d_1525_v20_: bool
        d_1526_v21_: int
        out47_: bool
        out48_: int
        out47_, out48_ = default__.m0((d_1524_v19_) | (d_1524_v19_), False, d_1523_v18_, (self).f4, globalState)
        d_1525_v20_ = out47_
        d_1526_v21_ = out48_
        r0 = d_1508_v4_
        return r0

    def m2(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_1527_v1_: _dafny.Seq
        d_1527_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rptj"))
        d_1528_v2_: _dafny.MultiSet
        d_1528_v2_ = _dafny.MultiSet([len(d_1527_v1_)])
        d_1529_v3_: _dafny.Seq
        d_1529_v3_ = _dafny.SeqWithoutIsStrInference([self.f14, (self).f6, self.f14, self.f14, (self).f6])
        d_1530_v4_: _dafny.Seq
        d_1530_v4_ = _dafny.SeqWithoutIsStrInference([d_1528_v2_, _dafny.MultiSet(d_1529_v3_)])
        d_1531_v5_: bool
        d_1532_v6_: int
        out49_: bool
        out50_: int
        def iife128_():
            coll52_ = _dafny.Map()
            compr_52_: _dafny.MultiSet
            for compr_52_ in (d_1530_v4_).Elements:
                d_1533_v0_: _dafny.MultiSet = compr_52_
                if (d_1533_v0_) in (d_1530_v4_):
                    coll52_[d_1533_v0_] = d_1527_v1_
            return _dafny.Map(coll52_)
        out49_, out50_ = default__.m0(iife128_()
        , (self).f4, d_1528_v2_, ((self).f6) != ((self).f6), globalState)
        d_1531_v5_ = out49_
        d_1532_v6_ = out50_
        d_1527_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1534_i0_ in range(default__.abs(592))])
        d_1535_v7_: C0
        nw233_ = C0()
        nw233_.ctor__(len(_dafny.SeqWithoutIsStrInference([(self).f4, not(d_1531_v5_), not((self).f13)])))
        d_1535_v7_ = nw233_
        d_1536_v8_: str
        d_1536_v8_ = _dafny.CodePoint('p')
        d_1537_v9_: _dafny.Map
        d_1537_v9_ = _dafny.Map({self.f14: default__.fm0(d_1536_v8_, globalState)})
        d_1538_v10_: _dafny.MultiSet
        d_1538_v10_ = _dafny.MultiSet([d_1536_v8_])
        d_1537_v9_ = (d_1537_v9_).set((0) - ((d_1538_v10_).cardinality), (d_1532_v6_) == ((0) - (self.f14)))
        hi10_ = 496
        for d_1539_i1_ in range(((self).f6) - ((self).f6), hi10_):
            if (False) and (not(p0)):
                d_1540_v11_: _dafny.Array
                def lambda50_(d_1541_i2_):
                    return (self).f4

                init32_ = lambda50_
                nw234_ = _dafny.Array(None, 10)
                for i0_32_ in range(nw234_.length(0)):
                    nw234_[i0_32_] = init32_(i0_32_)
                d_1540_v11_ = nw234_
                d_1542_v12_: _dafny.Map
                d_1542_v12_ = _dafny.Map({d_1540_v11_: p0})
                d_1543_v13_: C0
                nw235_ = C0()
                nw235_.ctor__((self).fm26((d_1535_v7_).f12, (self).fm8((self).f6, True, d_1527_v1_, d_1529_v3_, globalState), len(d_1542_v12_), globalState))
                d_1543_v13_ = nw235_
                d_1544_v14_: _dafny.Map
                d_1544_v14_ = _dafny.Map({(self).f13: (self).f6})
                d_1545_v15_: _dafny.Map
                d_1545_v15_ = _dafny.Map({len(d_1544_v14_): len(_dafny.SeqWithoutIsStrInference([(d_1543_v13_).f12 for d_1546_i3_ in range(default__.abs(218))]))})
                def iife129_():
                    coll53_ = _dafny.Map()
                    compr_53_: int
                    for compr_53_ in _dafny.IntegerRange(444, 515):
                        d_1547_v16_: int = compr_53_
                        if ((444) <= (d_1547_v16_)) and ((d_1547_v16_) < (515)):
                            coll53_[(d_1547_v16_) * ((d_1535_v7_).f12)] = (_dafny.MultiSet([610, self.f14])).cardinality
                    return _dafny.Map(coll53_)
                d_1545_v15_ = (d_1545_v15_).set(len(iife129_()
                ), len(_dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference([d_1529_v3_ for d_1548_i4_ in range(default__.abs(82))])).set(default__.safeIndex(len(_dafny.Map({(self).f6: 541})), len(_dafny.SeqWithoutIsStrInference([d_1529_v3_ for d_1549_i4_ in range(default__.abs(82))]))), d_1529_v3_))])))
                d_1550_v17_: _dafny.Array
                def lambda51_(d_1551_v7_):
                    def lambda52_(d_1552_i5_):
                        return default__.safeDivisionInt(d_1552_i5_, (d_1551_v7_).f12)

                    return lambda52_

                init33_ = lambda51_(d_1535_v7_)
                nw236_ = _dafny.Array(None, 6)
                for i0_33_ in range(nw236_.length(0)):
                    nw236_[i0_33_] = init33_(i0_33_)
                d_1550_v17_ = nw236_
                index223_ = default__.safeIndex(99, (d_1550_v17_).length(0))
                (d_1550_v17_)[index223_] = (40) * (self.f14)
                d_1553_v18_: _dafny.MultiSet
                d_1553_v18_ = _dafny.MultiSet([(self).f4, p0, (self).f4])
                index224_ = default__.safeIndex(99, (d_1550_v17_).length(0))
                (d_1550_v17_)[index224_] = default__.safeDivisionInt((self.f14) - ((d_1553_v18_).cardinality), (self).f6)
                d_1531_v5_ = (self).f4
                d_1532_v6_ = (d_1535_v7_).f12
            elif True:
                d_1554_v19_: _dafny.Map
                d_1554_v19_ = _dafny.Map({(self.f14) - ((d_1535_v7_).f12): (-598) - ((d_1535_v7_).f12)})
                d_1554_v19_ = default__.fm29(globalState)
                d_1532_v6_ = (d_1535_v7_).f12
                r0 = d_1536_v8_
                d_1536_v8_ = d_1536_v8_
                rhs207_ = self.f14
                rhs208_ = ((d_1535_v7_).f12) * (len(d_1527_v1_))
                rhs209_ = (self).f6
                rhs210_ = d_1536_v8_
                rhs211_ = (d_1535_v7_).f12
                lhs129_ = self
                lhs130_ = self
                lhs131_ = self
                d_1532_v6_ = rhs207_
                lhs129_.f14 = rhs208_
                lhs130_.f14 = rhs209_
                r0 = rhs210_
                lhs131_.f14 = rhs211_
            if d_1531_v5_:
                (self).f14 = d_1539_i1_
                (self).f14 = (d_1539_i1_) * (d_1532_v6_)
                d_1532_v6_ = (d_1535_v7_).f12
                d_1531_v5_ = not((self).f13)
                d_1555_v20_: _dafny.Array
                nw237_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
                d_1555_v20_ = nw237_
                index225_ = default__.safeIndex(239, (d_1555_v20_).length(0))
                (d_1555_v20_)[index225_] = d_1527_v1_
                d_1556_v21_: _dafny.Seq
                d_1556_v21_ = _dafny.SeqWithoutIsStrInference([d_1527_v1_])
                index226_ = default__.safeIndex(239, (d_1555_v20_).length(0))
                (d_1555_v20_)[index226_] = ((d_1556_v21_)[default__.safeIndex((d_1535_v7_).f12, len(d_1556_v21_))]) + (((d_1556_v21_)[default__.safeIndex(len(d_1529_v3_), len(d_1556_v21_))]) + (d_1527_v1_))
            elif True:
                d_1557_v22_: T0
                nw238_ = C2()
                nw238_.ctor__(d_1527_v1_, p0)
                d_1557_v22_ = nw238_
                d_1558_v23_: _dafny.Map
                d_1558_v23_ = _dafny.Map({d_1531_v5_: d_1557_v22_})
                d_1558_v23_ = (d_1558_v23_).set(not ((self).f13) or ((self).f13), d_1557_v22_)
                d_1531_v5_ = not(p0)
                d_1559_v24_: D4
                d_1559_v24_ = D4_DC12(d_1536_v8_)
                d_1560_v25_: _dafny.Array
                def lambda53_(d_1561_i6_):
                    return _dafny.CodePoint('d')

                init34_ = lambda53_
                nw239_ = _dafny.Array(None, 9)
                for i0_34_ in range(nw239_.length(0)):
                    nw239_[i0_34_] = init34_(i0_34_)
                d_1560_v25_ = nw239_
                index227_ = default__.safeIndex(656, (d_1560_v25_).length(0))
                (d_1560_v25_)[index227_] = d_1536_v8_
                d_1562_v26_: _dafny.Seq
                d_1562_v26_ = _dafny.SeqWithoutIsStrInference([not(d_1531_v5_)])
                d_1563_v27_: _dafny.Array
                nw240_ = _dafny.Array(None, 11)
                nw240_[int(0)] = not((self).f4)
                nw240_[int(1)] = (self).f4
                nw240_[int(2)] = (self).f4
                nw240_[int(3)] = (self).f13
                nw240_[int(4)] = False
                nw240_[int(5)] = (self).f13
                nw240_[int(6)] = False
                nw240_[int(7)] = not (p0) or (True)
                nw240_[int(8)] = not((self).f4)
                nw240_[int(9)] = (d_1562_v26_) == (d_1562_v26_)
                nw240_[int(10)] = default__.fm0(d_1536_v8_, globalState)
                d_1563_v27_ = nw240_
                index228_ = default__.safeIndex(225, (d_1563_v27_).length(0))
                (d_1563_v27_)[index228_] = (self).fm8(self.f14, p0, d_1527_v1_, _dafny.SeqWithoutIsStrInference([self.f14]), globalState)
                d_1564_v28_: D10
                d_1564_v28_ = D10_DC22(d_1527_v1_, p0, (self).f13, d_1563_v27_, d_1539_i1_)
                index229_ = default__.safeIndex(656, (d_1560_v25_).length(0))
                index230_ = default__.safeIndex(225, (d_1563_v27_).length(0))
                rhs212_ = default__.fm62(globalState)
                rhs213_ = (d_1536_v8_ if (d_1564_v28_).cf31 else _dafny.CodePoint('d'))
                rhs214_ = (p0) == ((d_1539_i1_) == ((d_1535_v7_).f12))
                rhs215_ = d_1563_v27_
                lhs132_ = d_1560_v25_
                lhs133_ = default__.safeIndex(656, (d_1560_v25_).length(0))
                lhs134_ = d_1563_v27_
                lhs135_ = default__.safeIndex(225, (d_1563_v27_).length(0))
                d_1559_v24_ = rhs212_
                lhs132_[lhs133_] = rhs213_
                lhs134_[lhs135_] = rhs214_
                d_1563_v27_ = rhs215_
                index231_ = default__.safeIndex(225, (d_1563_v27_).length(0))
                rhs216_ = p0
                rhs217_ = (d_1532_v6_) >= ((909 if d_1531_v5_ else len(d_1527_v1_)))
                lhs136_ = d_1563_v27_
                lhs137_ = default__.safeIndex(225, (d_1563_v27_).length(0))
                lhs136_[lhs137_] = rhs216_
                d_1531_v5_ = rhs217_
                d_1531_v5_ = p0
            d_1565_v29_: _dafny.Array
            nw241_ = _dafny.Array(None, 4)
            nw241_[int(0)] = (self).f13
            nw241_[int(1)] = d_1531_v5_
            nw241_[int(2)] = (49) < (d_1532_v6_)
            nw241_[int(3)] = (self).f4
            d_1565_v29_ = nw241_
            d_1565_v29_ = d_1565_v29_
            d_1566_v30_: _dafny.Set
            d_1566_v30_ = _dafny.Set({p0, False})
            (self).f14 = ((D16_DC33(d_1531_v5_, len(d_1527_v1_), _dafny.MultiSet([p0, True]), d_1532_v6_, d_1566_v30_)).cf54).cardinality
        d_1567_v31_: _dafny.Seq
        d_1567_v31_ = _dafny.SeqWithoutIsStrInference([p0, d_1531_v5_, d_1531_v5_])
        d_1568_v32_: _dafny.Seq
        d_1568_v32_ = _dafny.SeqWithoutIsStrInference([d_1567_v31_])
        d_1569_v33_: D12
        d_1569_v33_ = D12_DC27(d_1568_v32_)
        source29_ = d_1569_v33_
        if source29_.is_DC28:
            d_1570___mcc_h0_ = source29_.cf43
            d_1571___mcc_h1_ = source29_.cf44
            d_1572___mcc_h2_ = source29_.cf45
            d_1573___mcc_h3_ = source29_.cf46
            d_1574___mcc_h4_ = source29_.cf47
            d_1575_cf47_ = d_1574___mcc_h4_
            d_1576_cf46_ = d_1573___mcc_h3_
            d_1577_cf45_ = d_1572___mcc_h2_
            d_1578_cf44_ = d_1571___mcc_h1_
            d_1579_cf43_ = d_1570___mcc_h0_
            d_1580_v34_: _dafny.Array
            def lambda54_(d_1581_i7_):
                return _dafny.Map({(self).f13: False})

            init35_ = lambda54_
            nw242_ = _dafny.Array(None, 29)
            for i0_35_ in range(nw242_.length(0)):
                nw242_[i0_35_] = init35_(i0_35_)
            d_1580_v34_ = nw242_
            d_1580_v34_ = d_1580_v34_
            if default__.fm0(d_1536_v8_, globalState):
                d_1582_v35_: _dafny.Set
                d_1582_v35_ = _dafny.Set({(self).f13, d_1576_cf46_})
                d_1583_v36_: _dafny.Map
                d_1583_v36_ = _dafny.Map({(d_1582_v35_).intersection(d_1582_v35_): (self).f4})
                d_1583_v36_ = (d_1583_v36_).set((d_1582_v35_) | (_dafny.Set({(self).f4})), not ((self).f4) or (d_1531_v5_))
                d_1531_v5_ = (self).f13
                d_1584_v37_: C3
                nw243_ = C3()
                nw243_.ctor__(self.f5, (0) - (self.f14), p0)
                d_1584_v37_ = nw243_
                d_1585_v38_: _dafny.Map
                d_1585_v38_ = _dafny.Map({d_1584_v37_: (self).f13})
                d_1585_v38_ = (d_1585_v38_).set(d_1584_v37_, d_1576_cf46_)
                d_1536_v8_ = d_1536_v8_
                d_1527_v1_ = (d_1527_v1_) + (d_1527_v1_)
            elif True:
                d_1578_cf44_ = 483
                d_1586_v39_: _dafny.Array
                def lambda55_(d_1587_i8_):
                    return (d_1587_i8_) * (146)

                init36_ = lambda55_
                nw244_ = _dafny.Array(None, 7)
                for i0_36_ in range(nw244_.length(0)):
                    nw244_[i0_36_] = init36_(i0_36_)
                d_1586_v39_ = nw244_
                d_1588_v40_: D16
                d_1588_v40_ = D16_DC32(d_1586_v39_)
                d_1586_v39_ = (d_1588_v40_).cf51
                d_1531_v5_ = d_1577_cf45_
                index232_ = default__.safeIndex(769, (d_1586_v39_).length(0))
                (d_1586_v39_)[index232_] = (d_1535_v7_).f12
                index233_ = default__.safeIndex(769, (d_1586_v39_).length(0))
                (d_1586_v39_)[index233_] = 342
                d_1575_cf47_ = (d_1535_v7_).f12
            d_1589_v41_: _dafny.Set
            d_1589_v41_ = _dafny.Set({(self).fm7(globalState)})
            d_1590_v42_: _dafny.Map
            d_1590_v42_ = _dafny.Map({d_1577_cf45_: d_1577_cf45_})
            d_1591_v43_: _dafny.Map
            d_1591_v43_ = _dafny.Map({len(d_1529_v3_): len(d_1590_v42_)})
            def iife130_():
                coll54_ = _dafny.Set()
                compr_54_: int
                for compr_54_ in (d_1528_v2_).Elements:
                    d_1592_v44_: int = compr_54_
                    if (d_1592_v44_) in (d_1528_v2_):
                        coll54_ = coll54_.union(_dafny.Set([(d_1592_v44_) + (len(_dafny.Set({True, True, False, False})))]))
                return _dafny.Set(coll54_)
            rhs218_ = ((_dafny.Set({((d_1591_v43_)[(self).f6] if ((self).f6) in (d_1591_v43_) else (d_1535_v7_).f12), (d_1535_v7_).f12})) - (iife130_()
            )).issubset(d_1589_v41_)
            rhs219_ = ((d_1578_cf44_) * ((self).f6)) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "reowmtl"))))
            d_1531_v5_ = rhs218_
            d_1579_cf43_ = rhs219_
            d_1593_v45_: C2
            nw245_ = C2()
            nw245_.ctor__(d_1527_v1_, d_1577_cf45_)
            d_1593_v45_ = nw245_
        elif True:
            d_1594___mcc_h5_ = source29_.cf42
            d_1595_cf42_ = d_1594___mcc_h5_
            d_1596_v46_: D9
            d_1596_v46_ = D9_DC19(d_1527_v1_)
            d_1597_v47_: _dafny.Array
            nw246_ = _dafny.Array(None, 3)
            nw246_[int(0)] = D9_DC19(default__.fm37((self).f13, False, len(d_1527_v1_), p0, globalState))
            nw246_[int(1)] = d_1596_v46_
            nw246_[int(2)] = d_1596_v46_
            d_1597_v47_ = nw246_
            d_1598_v48_: _dafny.Set
            d_1598_v48_ = _dafny.Set({d_1531_v5_})
            pat_let_tv29_ = d_1532_v6_
            pat_let_tv30_ = d_1598_v48_
            pat_let_tv31_ = globalState
            pat_let_tv32_ = d_1527_v1_
            pat_let_tv33_ = d_1527_v1_
            pat_let_tv34_ = d_1527_v1_
            nw247_ = _dafny.Array(None, 20)
            nw247_[int(0)] = D9_DC19(d_1527_v1_)
            nw247_[int(1)] = d_1596_v46_
            nw247_[int(2)] = D9_DC19(d_1527_v1_)
            def iife131_(_pat_let38_0):
                def iife132_(d_1599_dt__update__tmp_h0_):
                    def iife133_(_pat_let39_0):
                        def iife134_(d_1600_dt__update_hcf25_h0_):
                            return D9_DC19(d_1600_dt__update_hcf25_h0_)
                        return iife134_(_pat_let39_0)
                    return iife133_(default__.fm50(self.f14, pat_let_tv29_, len(pat_let_tv30_), not(not(not(False))), pat_let_tv31_))
                return iife132_(_pat_let38_0)
            nw247_[int(3)] = iife131_(d_1596_v46_)
            nw247_[int(4)] = d_1596_v46_
            nw247_[int(5)] = D9_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ruvt")))
            nw247_[int(6)] = d_1596_v46_
            nw247_[int(7)] = D9_DC19(d_1527_v1_)
            nw247_[int(8)] = d_1596_v46_
            def iife136_(_pat_let41_0):
                def iife137_(d_1601_dt__update__tmp_h1_):
                    def iife138_(_pat_let42_0):
                        def iife139_(d_1602_dt__update_hcf25_h1_):
                            return D9_DC19(d_1602_dt__update_hcf25_h1_)
                        return iife139_(_pat_let42_0)
                    return iife138_(pat_let_tv32_)
                return iife137_(_pat_let41_0)
            def iife135_(_pat_let40_0):
                def iife140_(d_1603_dt__update__tmp_h2_):
                    def iife141_(_pat_let43_0):
                        def iife142_(d_1604_dt__update_hcf25_h2_):
                            return D9_DC19(d_1604_dt__update_hcf25_h2_)
                        return iife142_(_pat_let43_0)
                    return iife141_(pat_let_tv33_)
                return iife140_(_pat_let40_0)
            nw247_[int(9)] = (iife135_(iife136_(d_1596_v46_)) if p0 else d_1596_v46_)
            nw247_[int(10)] = d_1596_v46_
            nw247_[int(11)] = d_1596_v46_
            nw247_[int(12)] = d_1596_v46_
            nw247_[int(13)] = (D9_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khiwbj"))) if (self).f13 else d_1596_v46_)
            nw247_[int(14)] = D9_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qmqtcd")))
            nw247_[int(15)] = d_1596_v46_
            nw247_[int(16)] = d_1596_v46_
            nw247_[int(17)] = D9_DC19(d_1527_v1_)
            def iife143_(_pat_let44_0):
                def iife144_(d_1605_dt__update__tmp_h3_):
                    def iife145_(_pat_let45_0):
                        def iife146_(d_1606_dt__update_hcf25_h3_):
                            return D9_DC19(d_1606_dt__update_hcf25_h3_)
                        return iife146_(_pat_let45_0)
                    return iife145_(pat_let_tv34_)
                return iife144_(_pat_let44_0)
            nw247_[int(18)] = iife143_(d_1596_v46_)
            nw247_[int(19)] = d_1596_v46_
            d_1597_v47_ = nw247_
            d_1607_v49_: _dafny.Array
            nw248_ = _dafny.Array(None, 9)
            nw248_[int(0)] = not((self).f13)
            nw248_[int(1)] = d_1531_v5_
            nw248_[int(2)] = (default__.fm53(d_1532_v6_, (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_1531_v5_: (self).f13}) for d_1608_i9_ in range(default__.abs(22))])).set(default__.safeIndex((d_1535_v7_).f12, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({d_1531_v5_: (self).f13}) for d_1609_i9_ in range(default__.abs(22))]))), (_dafny.Map({(self).f13: d_1531_v5_})).set(p0, (self).f13)), globalState)).issubset(d_1598_v48_)
            nw248_[int(3)] = ((self).f4 if p0 else default__.fm0(d_1536_v8_, globalState))
            nw248_[int(4)] = (d_1567_v31_) <= (d_1567_v31_)
            nw248_[int(5)] = (self).f13
            nw248_[int(6)] = (self).f4
            nw248_[int(7)] = (self).f13
            nw248_[int(8)] = False
            d_1607_v49_ = nw248_
            index234_ = default__.safeIndex(379, (d_1607_v49_).length(0))
            (d_1607_v49_)[index234_] = True
            d_1610_v50_: _dafny.Array
            nw249_ = _dafny.Array(D18.default()(), 1)
            d_1610_v50_ = nw249_
            d_1611_v51_: D11
            d_1611_v51_ = D11_DC26(d_1531_v5_)
            d_1612_v52_: _dafny.Map
            d_1612_v52_ = _dafny.Map({d_1611_v51_: True})
            d_1613_v53_: _dafny.Seq
            d_1613_v53_ = _dafny.SeqWithoutIsStrInference([d_1612_v52_])
            d_1614_v54_: D18
            d_1614_v54_ = D18_DC35(d_1613_v53_)
            index235_ = default__.safeIndex(848, (d_1610_v50_).length(0))
            (d_1610_v50_)[index235_] = d_1614_v54_
            index236_ = default__.safeIndex(379, (d_1607_v49_).length(0))
            index237_ = default__.safeIndex(848, (d_1610_v50_).length(0))
            rhs220_ = ((default__.fm37(p0, (self).f13, d_1532_v6_, (self).f13, globalState)) + (d_1527_v1_)) <= (d_1527_v1_)
            rhs221_ = d_1614_v54_
            lhs138_ = d_1607_v49_
            lhs139_ = default__.safeIndex(379, (d_1607_v49_).length(0))
            lhs140_ = d_1610_v50_
            lhs141_ = default__.safeIndex(848, (d_1610_v50_).length(0))
            lhs138_[lhs139_] = rhs220_
            lhs140_[lhs141_] = rhs221_
            d_1615_v55_: D12
            d_1615_v55_ = D12_DC28((self).f6, (d_1535_v7_).f12, p0, True, d_1532_v6_)
            rhs222_ = default__.safeModuloInt((self).f6, (d_1532_v6_ if not(d_1531_v5_) else d_1532_v6_))
            rhs223_ = ((d_1598_v48_).intersection(d_1598_v48_)).issubset(_dafny.Set({(self).f13, True, (d_1615_v55_).cf45, (self).f4, False}))
            lhs142_ = self
            lhs142_.f14 = rhs222_
            d_1531_v5_ = rhs223_
            d_1616_v56_: C1
            nw250_ = C1()
            nw250_.ctor__(self.f14, 425, (self.f5 if (self).f13 else self.f5), self.f14, not(p0))
            d_1616_v56_ = nw250_
        pat_let_tv35_ = d_1536_v8_
        pat_let_tv36_ = d_1536_v8_
        pat_let_tv37_ = d_1536_v8_
        pat_let_tv38_ = d_1527_v1_
        def lambda56_(source30_):
            if source30_.is_DC24:
                d_1617___mcc_h6_ = source30_.cf36
                d_1618___mcc_h7_ = source30_.cf37
                d_1619___mcc_h8_ = source30_.cf38
                d_1620_cf38_ = d_1619___mcc_h8_
                d_1621_cf37_ = d_1618___mcc_h7_
                d_1622_cf36_ = d_1617___mcc_h6_
                return _dafny.CodePoint('w')
            elif source30_.is_DC25:
                d_1623___mcc_h9_ = source30_.cf39
                d_1624___mcc_h10_ = source30_.cf40
                d_1625_cf40_ = d_1624___mcc_h10_
                d_1626_cf39_ = d_1623___mcc_h9_
                return pat_let_tv35_
            elif source30_.is_DC26:
                d_1627___mcc_h11_ = source30_.cf41
                d_1628_cf41_ = d_1627___mcc_h11_
                return pat_let_tv36_
            elif True:
                d_1629___mcc_h12_ = source30_.cf35
                d_1630_cf35_ = d_1629___mcc_h12_
                return pat_let_tv37_

        def iife147_(_pat_let46_0):
            def iife148_(d_1631_dt__update__tmp_h4_):
                def iife149_(_pat_let47_0):
                    def iife150_(d_1632_dt__update_hcf38_h0_):
                        return D11_DC24((d_1631_dt__update__tmp_h4_).cf36, (d_1631_dt__update__tmp_h4_).cf37, d_1632_dt__update_hcf38_h0_)
                    return iife150_(_pat_let47_0)
                return iife149_(len(pat_let_tv38_))
            return iife148_(_pat_let46_0)
        r0 = lambda56_(iife147_(D11_DC24(p0, not((self).f4), (d_1535_v7_).f12)))
        d_1633_v57_: _dafny.Array
        nw251_ = _dafny.Array(D1.default()(), 19)
        d_1633_v57_ = nw251_
        r1 = d_1633_v57_
        return r0, r1

    @property
    def f13(self):
        return self._f13

class C6(T0, T1):
    def  __init__(self):
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self._f6: int = int(0)
        self._f4: bool = False
        self._f11: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f5(self):
        return self._f5
    @f5.setter
    def f5(self, value):
        self._f5 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f4(self):
        return self._f4
    def ctor__(self, f11, f4, f5, f6):
        (self)._f11 = f11
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6

    def fm6(self, p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f6]))).intersection(_dafny.MultiSet([(self).f6]))

    def fm7(self, globalState):
        return (((self).f6) * ((_dafny.MultiSet([len(_dafny.Map({(self).f11: (self).f6})), (self).f6])).cardinality)) * (default__.safeDivisionInt((self).f6, len(_dafny.Map({not(False): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkuc"))}))))

    def fm8(self, p0, p1, p2, p3, globalState):
        return not(((self).f11) and (((self).f6) < ((_dafny.MultiSet([(self).f4, (self).f4])).cardinality)))

    def fm9(self, p0, p1, p2, globalState):
        return (self).f6

    def fm15(self, globalState):
        source31_ = D0_DC1()
        if source31_.is_DC1:
            return (0) - ((self).f6)
        elif source31_.is_DC2:
            d_1634___mcc_h0_ = source31_.cf1
            d_1635_cf1_ = d_1634___mcc_h0_
            return default__.safeDivisionInt(d_1635_cf1_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f11]))])))
        elif True:
            d_1636___mcc_h1_ = source31_.cf0
            d_1637_cf0_ = d_1636___mcc_h1_
            return len((_dafny.SeqWithoutIsStrInference([(self).f4, (self).f11])) + ((D3_DC8(_dafny.SeqWithoutIsStrInference([False]))).cf7))

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_1638_v0_: _dafny.Seq
        d_1638_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gteox"))
        d_1639_v1_: _dafny.Map
        d_1639_v1_ = _dafny.Map({not((self).f11): d_1638_v0_})
        d_1640_v2_: D2
        d_1640_v2_ = D2_DC6()
        d_1641_v3_: D1
        d_1641_v3_ = D1_DC3((self).f11)
        d_1642_v4_: _dafny.Array
        nw252_ = _dafny.Array(None, 17)
        nw252_[int(0)] = _dafny.Map({not((self).f11): d_1638_v0_})
        nw252_[int(1)] = _dafny.Map({(self).f11: d_1638_v0_})
        nw252_[int(2)] = d_1639_v1_
        nw252_[int(3)] = _dafny.Map({(self).f11: (default__.fm16((self).f4, (self).f4, (self).f11, d_1640_v2_, globalState)).set(default__.safeIndex((self).f6, len(default__.fm16((self).f4, (self).f4, (self).f11, d_1640_v2_, globalState))), _dafny.CodePoint('l'))})
        nw252_[int(4)] = d_1639_v1_
        nw252_[int(5)] = d_1639_v1_
        nw252_[int(6)] = d_1639_v1_
        nw252_[int(7)] = d_1639_v1_
        nw252_[int(8)] = (d_1639_v1_ if (d_1641_v3_).cf2 else _dafny.Map({(self).f11: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwkwygjor"))}))
        nw252_[int(9)] = (d_1639_v1_) | (d_1639_v1_)
        nw252_[int(10)] = (d_1639_v1_) | (d_1639_v1_)
        nw252_[int(11)] = d_1639_v1_
        nw252_[int(12)] = d_1639_v1_
        nw252_[int(13)] = d_1639_v1_
        nw252_[int(14)] = d_1639_v1_
        nw252_[int(15)] = d_1639_v1_
        nw252_[int(16)] = d_1639_v1_
        d_1642_v4_ = nw252_
        index238_ = default__.safeIndex(480, (d_1642_v4_).length(0))
        (d_1642_v4_)[index238_] = (d_1639_v1_) | (d_1639_v1_)
        index239_ = default__.safeIndex(480, (d_1642_v4_).length(0))
        (d_1642_v4_)[index239_] = d_1639_v1_
        hi11_ = (self).fm7(globalState)
        for d_1643_i0_ in range(p0, hi11_):
            d_1644_v5_: _dafny.Array
            def lambda57_(d_1645_i0_):
                def lambda58_(d_1646_i1_):
                    return (d_1646_i1_) - (d_1645_i0_)

                return lambda58_

            init37_ = lambda57_(d_1643_i0_)
            nw253_ = _dafny.Array(None, 23)
            for i0_37_ in range(nw253_.length(0)):
                nw253_[i0_37_] = init37_(i0_37_)
            d_1644_v5_ = nw253_
            d_1647_v6_: _dafny.Map
            d_1647_v6_ = _dafny.Map({(self).f6: (self).f4})
            d_1648_v7_: _dafny.Map
            d_1648_v7_ = _dafny.Map({d_1644_v5_: d_1647_v6_})
            d_1649_v8_: _dafny.Map
            d_1649_v8_ = _dafny.Map({(self).f4: d_1643_i0_})
            d_1648_v7_ = (d_1648_v7_).set(d_1644_v5_, _dafny.Map({((d_1649_v8_)[(self).f11] if ((self).f11) in (d_1649_v8_) else (self).f6): (self).f4}))
            d_1650_v9_: _dafny.MultiSet
            d_1650_v9_ = _dafny.MultiSet([d_1643_i0_])
            if (d_1650_v9_).isdisjoint(d_1650_v9_):
                d_1651_v10_: _dafny.Array
                nw254_ = _dafny.Array(_dafny.Seq({}), 25)
                d_1651_v10_ = nw254_
                d_1652_v11_: _dafny.Seq
                d_1652_v11_ = _dafny.SeqWithoutIsStrInference([d_1644_v5_])
                index240_ = default__.safeIndex(952, (d_1651_v10_).length(0))
                (d_1651_v10_)[index240_] = d_1652_v11_
                index241_ = default__.safeIndex(952, (d_1651_v10_).length(0))
                (d_1651_v10_)[index241_] = d_1652_v11_
                d_1653_v12_: int
                d_1653_v12_ = 85
                d_1654_v13_: bool
                d_1654_v13_ = False
                rhs224_ = (p0) - ((0) - ((self).f6))
                rhs225_ = not((self).f4)
                d_1653_v12_ = rhs224_
                d_1654_v13_ = rhs225_
                d_1655_v14_: _dafny.Array
                def lambda59_(d_1656_v12_):
                    def lambda60_(d_1657_i2_):
                        return D0_DC0(d_1656_v12_)

                    return lambda60_

                init38_ = lambda59_(d_1653_v12_)
                nw255_ = _dafny.Array(None, 16)
                for i0_38_ in range(nw255_.length(0)):
                    nw255_[i0_38_] = init38_(i0_38_)
                d_1655_v14_ = nw255_
                d_1655_v14_ = d_1655_v14_
                d_1658_v15_: _dafny.Set
                d_1658_v15_ = _dafny.Set({-882})
                d_1653_v12_ = (len(d_1658_v15_)) + (p0)
                d_1659_v16_: _dafny.Map
                d_1659_v16_ = _dafny.Map({(self).f11: d_1644_v5_})
                d_1660_v17_: _dafny.Map
                d_1660_v17_ = _dafny.Map({d_1659_v16_: (((d_1649_v8_)[False] if (False) in (d_1649_v8_) else (self).fm9(not(d_1654_v13_), (self).f11, p0, globalState)) if (self).f4 else d_1643_i0_)})
                d_1660_v17_ = (d_1660_v17_).set((d_1659_v16_).set((self).f4, d_1644_v5_), d_1643_i0_)
            elif True:
                d_1661_v18_: C0
                nw256_ = C0()
                nw256_.ctor__(d_1643_i0_)
                d_1661_v18_ = nw256_
                d_1662_v19_: int
                d_1662_v19_ = 759
                d_1663_v20_: _dafny.Map
                d_1663_v20_ = _dafny.Map({p0: 410})
                d_1664_v21_: _dafny.Map
                d_1664_v21_ = _dafny.Map({p0: d_1663_v20_})
                d_1662_v19_ = len(d_1664_v21_)
                d_1662_v19_ = ((d_1661_v18_).f12) * ((0) - (d_1662_v19_))
                d_1665_v22_: bool
                d_1665_v22_ = True
                index242_ = default__.safeIndex(426, (d_1644_v5_).length(0))
                (d_1644_v5_)[index242_] = (default__.fm3(d_1662_v19_, d_1665_v22_, globalState)) + (d_1662_v19_)
                index243_ = default__.safeIndex(914, (d_1644_v5_).length(0))
                (d_1644_v5_)[index243_] = ((d_1649_v8_)[(self).f4] if ((self).f4) in (d_1649_v8_) else d_1643_i0_)
                d_1666_v23_: D2
                d_1666_v23_ = D2_DC7((d_1661_v18_).f12)
                index244_ = default__.safeIndex(426, (d_1644_v5_).length(0))
                index245_ = default__.safeIndex(914, (d_1644_v5_).length(0))
                rhs226_ = d_1665_v22_
                rhs227_ = (0) - (((d_1661_v18_).f12) - ((len(default__.fm16((self).f4, d_1665_v22_, (self).f11, d_1640_v2_, globalState))) + (p0)))
                rhs228_ = (p0) * ((d_1666_v23_).cf6)
                rhs229_ = (self).f6
                rhs230_ = (d_1641_v3_).cf2
                lhs143_ = d_1644_v5_
                lhs144_ = default__.safeIndex(426, (d_1644_v5_).length(0))
                lhs145_ = d_1644_v5_
                lhs146_ = default__.safeIndex(914, (d_1644_v5_).length(0))
                d_1665_v22_ = rhs226_
                lhs143_[lhs144_] = rhs227_
                lhs145_[lhs146_] = rhs228_
                d_1662_v19_ = rhs229_
                d_1665_v22_ = rhs230_
                d_1662_v19_ = (d_1650_v9_).cardinality
            d_1667_v24_: _dafny.Set
            d_1667_v24_ = _dafny.Set({(self).f4})
            d_1668_v25_: _dafny.Map
            d_1668_v25_ = _dafny.Map({d_1667_v24_: (self).f11})
            d_1669_v26_: _dafny.Set
            d_1669_v26_ = _dafny.Set({len(d_1668_v25_), p0, p0, p0})
            d_1669_v26_ = d_1669_v26_
            d_1670_v27_: int
            d_1670_v27_ = -271
            d_1670_v27_ = d_1670_v27_
        d_1671_i3_: int
        d_1671_i3_ = 0
        with _dafny.label("7"):
            while (self).f11:
                with _dafny.c_label("7"):
                    if (d_1671_i3_) >= (100):
                        raise _dafny.Break("7")
                    d_1671_i3_ = (d_1671_i3_) + (1)
                    d_1672_v28_: _dafny.Array
                    nw257_ = _dafny.Array(int(0), 12)
                    d_1672_v28_ = nw257_
                    d_1672_v28_ = d_1672_v28_
                    d_1673_v29_: bool
                    d_1673_v29_ = False
                    d_1674_v30_: _dafny.Seq
                    d_1674_v30_ = _dafny.SeqWithoutIsStrInference([583])
                    d_1673_v29_ = (d_1674_v30_) == (d_1674_v30_)
                    d_1673_v29_ = (self).f11
                    d_1675_v31_: int
                    d_1675_v31_ = -981
                    rhs231_ = (self).f11
                    rhs232_ = 759
                    d_1673_v29_ = rhs231_
                    d_1675_v31_ = rhs232_
                    pass
            pass
        d_1676_v32_: _dafny.MultiSet
        d_1676_v32_ = _dafny.MultiSet([(self).f11])
        d_1677_v33_: _dafny.Map
        d_1677_v33_ = _dafny.Map({(0) - ((self).f6): (self).f4})
        if (((self).f6) + ((self).f6)) == ((p0) * (((d_1676_v32_)[(self).f4] if ((self).f4) in (d_1676_v32_) else len(d_1677_v33_)))):
            d_1678_v34_: _dafny.Seq
            d_1678_v34_ = _dafny.SeqWithoutIsStrInference([327, p0])
            d_1679_v35_: C0
            nw258_ = C0()
            nw258_.ctor__(len(d_1678_v34_))
            d_1679_v35_ = nw258_
            d_1680_v36_: _dafny.Seq
            d_1680_v36_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfhdckno"))])
            d_1681_v37_: C0
            nw259_ = C0()
            nw259_.ctor__(len(d_1680_v36_))
            d_1681_v37_ = nw259_
            d_1682_v38_: _dafny.Array
            def lambda61_(d_1683_v0_, d_1684_v34_):
                def lambda62_(d_1685_i4_):
                    return ((d_1683_v0_) + (d_1683_v0_)).set(default__.safeIndex((d_1684_v34_)[default__.safeIndex(239, len(d_1684_v34_))], len((d_1683_v0_) + (d_1683_v0_))), _dafny.CodePoint('j'))

                return lambda62_

            init39_ = lambda61_(d_1638_v0_, d_1678_v34_)
            nw260_ = _dafny.Array(None, 28)
            for i0_39_ in range(nw260_.length(0)):
                nw260_[i0_39_] = init39_(i0_39_)
            d_1682_v38_ = nw260_
            index246_ = default__.safeIndex(734, (d_1682_v38_).length(0))
            (d_1682_v38_)[index246_] = d_1638_v0_
            d_1686_v39_: str
            d_1686_v39_ = _dafny.CodePoint('s')
            index247_ = default__.safeIndex(734, (d_1682_v38_).length(0))
            (d_1682_v38_)[index247_] = (d_1638_v0_) + ((d_1638_v0_) + ((d_1638_v0_).set(default__.safeIndex((self).f6, len(d_1638_v0_)), d_1686_v39_)))
            d_1687_v40_: _dafny.Set
            d_1687_v40_ = _dafny.Set({p0})
            d_1688_v41_: _dafny.Array
            nw261_ = _dafny.Array(int(0), 13)
            d_1688_v41_ = nw261_
            d_1689_v42_: _dafny.Map
            d_1689_v42_ = _dafny.Map({len((d_1687_v40_).intersection(d_1687_v40_)): d_1688_v41_})
            d_1689_v42_ = (d_1689_v42_).set(-824, d_1688_v41_)
            d_1690_v43_: _dafny.Map
            d_1690_v43_ = _dafny.Map({d_1638_v0_: (d_1681_v37_).f12})
            d_1691_v44_: _dafny.Array
            nw262_ = _dafny.Array(None, 6)
            nw262_[int(0)] = (self).f4
            nw262_[int(1)] = (self).f11
            nw262_[int(2)] = (self).f11
            nw262_[int(3)] = (self).f4
            nw262_[int(4)] = (self).f4
            nw262_[int(5)] = (self).f11
            d_1691_v44_ = nw262_
            d_1692_v45_: _dafny.Map
            d_1692_v45_ = _dafny.Map({d_1691_v44_: (self).f11})
            d_1690_v43_ = (d_1690_v43_).set(_dafny.SeqWithoutIsStrInference([d_1686_v39_ for d_1693_i5_ in range(default__.abs(863))]), len(default__.fm18(d_1638_v0_, len(d_1692_v45_), globalState)))
        elif True:
            d_1694_v46_: int
            d_1694_v46_ = 602
            d_1695_v47_: str
            d_1695_v47_ = _dafny.CodePoint('j')
            rhs233_ = p0
            rhs234_ = p0
            rhs235_ = d_1695_v47_
            d_1694_v46_ = rhs233_
            d_1694_v46_ = rhs234_
            d_1695_v47_ = rhs235_
            if (self).f4:
                d_1694_v46_ = p0
                d_1694_v46_ = (len(default__.fm16((self).f11, (self).f4, (D3_DC9((self).f4, d_1694_v46_, (self).f11)).cf10, d_1640_v2_, globalState)) if (d_1694_v46_) > (p0) else d_1694_v46_)
                d_1696_v48_: _dafny.Set
                d_1696_v48_ = _dafny.Set({274, p0, len(d_1638_v0_)})
                d_1696_v48_ = (d_1696_v48_) | (_dafny.Set({p0}))
                d_1697_v49_: _dafny.Array
                def lambda63_(d_1698_p0_):
                    def lambda64_(d_1699_i6_):
                        return (d_1699_i6_) * (d_1698_p0_)

                    return lambda64_

                init40_ = lambda63_(p0)
                nw263_ = _dafny.Array(None, 11)
                for i0_40_ in range(nw263_.length(0)):
                    nw263_[i0_40_] = init40_(i0_40_)
                d_1697_v49_ = nw263_
                index248_ = default__.safeIndex(996, (d_1697_v49_).length(0))
                (d_1697_v49_)[index248_] = ((self).f6 if (self).f4 else p0)
                index249_ = default__.safeIndex(996, (d_1697_v49_).length(0))
                (d_1697_v49_)[index249_] = 348
                d_1700_v50_: C0
                nw264_ = C0()
                nw264_.ctor__((-268 if (self).f11 else p0))
                d_1700_v50_ = nw264_
            elif True:
                d_1701_v51_: D4
                d_1701_v51_ = D4_DC12(d_1695_v47_)
                d_1695_v47_ = (d_1701_v51_).cf13
                d_1702_v52_: _dafny.Seq
                d_1702_v52_ = _dafny.SeqWithoutIsStrInference([(len(default__.fm16((self).f11, default__.fm0(d_1695_v47_, globalState), (self).f11, d_1640_v2_, globalState))) * ((self).fm15(globalState))])
                d_1703_v53_: _dafny.Seq
                d_1703_v53_ = _dafny.SeqWithoutIsStrInference([(self).f4])
                d_1704_v54_: _dafny.Array
                nw265_ = _dafny.Array(int(0), 29)
                d_1704_v54_ = nw265_
                d_1705_v55_: _dafny.Array
                nw266_ = _dafny.Array(None, 19)
                nw266_[int(0)] = d_1703_v53_
                nw266_[int(1)] = _dafny.SeqWithoutIsStrInference([(self).f4])
                nw266_[int(2)] = default__.fm19((self).f11, default__.fm3((0) - (len(_dafny.Map({d_1695_v47_: d_1704_v54_}))), (self).f11, globalState), globalState)
                nw266_[int(3)] = d_1703_v53_
                nw266_[int(4)] = d_1703_v53_
                nw266_[int(5)] = d_1703_v53_
                nw266_[int(6)] = d_1703_v53_
                nw266_[int(7)] = d_1703_v53_
                nw266_[int(8)] = (d_1703_v53_ if (self).f4 else d_1703_v53_)
                nw266_[int(9)] = d_1703_v53_
                nw266_[int(10)] = d_1703_v53_
                nw266_[int(11)] = ((d_1703_v53_).set(default__.safeIndex(d_1694_v46_, len(d_1703_v53_)), True)).set(default__.safeIndex((self).f6, len((d_1703_v53_).set(default__.safeIndex(d_1694_v46_, len(d_1703_v53_)), True))), False)
                nw266_[int(12)] = (d_1703_v53_) + (d_1703_v53_)
                nw266_[int(13)] = d_1703_v53_
                nw266_[int(14)] = d_1703_v53_
                nw266_[int(15)] = d_1703_v53_
                nw266_[int(16)] = d_1703_v53_
                nw266_[int(17)] = d_1703_v53_
                nw266_[int(18)] = d_1703_v53_
                d_1705_v55_ = nw266_
                index250_ = default__.safeIndex(737, (d_1705_v55_).length(0))
                (d_1705_v55_)[index250_] = ((D4_DC13((self).f6, (self).f11, d_1694_v46_, d_1703_v53_)).cf17) + (d_1703_v53_)
                index251_ = default__.safeIndex(737, (d_1705_v55_).length(0))
                rhs236_ = d_1702_v52_
                rhs237_ = p0
                rhs238_ = -794
                rhs239_ = d_1703_v53_
                lhs147_ = d_1705_v55_
                lhs148_ = default__.safeIndex(737, (d_1705_v55_).length(0))
                d_1702_v52_ = rhs236_
                d_1694_v46_ = rhs237_
                d_1694_v46_ = rhs238_
                lhs147_[lhs148_] = rhs239_
                d_1706_v56_: D3
                d_1706_v56_ = D3_DC10(p0)
                d_1707_v57_: D3
                d_1707_v57_ = D3_DC11(d_1706_v56_)
                d_1708_v58_: _dafny.Array
                nw267_ = _dafny.Array(None, 6)
                nw267_[int(0)] = (self).f11
                nw267_[int(1)] = (self).f4
                nw267_[int(2)] = (self).f4
                nw267_[int(3)] = (self).f4
                nw267_[int(4)] = (self).f4
                nw267_[int(5)] = (self).f4
                d_1708_v58_ = nw267_
                d_1709_v59_: _dafny.MultiSet
                d_1709_v59_ = _dafny.MultiSet([d_1708_v58_])
                d_1710_v60_: _dafny.Map
                d_1710_v60_ = _dafny.Map({d_1707_v57_: (d_1709_v59_) - (d_1709_v59_)})
                d_1710_v60_ = (d_1710_v60_).set(d_1707_v57_, d_1709_v59_)
                d_1711_v61_: _dafny.MultiSet
                d_1711_v61_ = _dafny.MultiSet([(-133 if (self).f4 else p0), d_1694_v46_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1712_i7_ in range(default__.abs(-336))])), p0])
                d_1711_v61_ = (_dafny.MultiSet(d_1702_v52_))
                d_1713_v62_: _dafny.Map
                d_1713_v62_ = _dafny.Map({p0: p0})
                d_1714_v63_: _dafny.Map
                d_1714_v63_ = _dafny.Map({(d_1711_v61_).cardinality: d_1638_v0_})
                d_1715_v64_: _dafny.Seq
                d_1715_v64_ = _dafny.SeqWithoutIsStrInference([d_1714_v63_, d_1714_v63_])
                index252_ = default__.safeIndex(300, (d_1704_v54_).length(0))
                (d_1704_v54_)[index252_] = (default__.fm3(((d_1713_v62_)[d_1694_v46_] if (d_1694_v46_) in (d_1713_v62_) else d_1694_v46_), (self).f4, globalState) if (self).f4 else len((d_1715_v64_)[default__.safeIndex(len(d_1638_v0_), len(d_1715_v64_))]))
                index253_ = default__.safeIndex(300, (d_1704_v54_).length(0))
                (d_1704_v54_)[index253_] = (p0 if (self).f11 else len((D3_DC8((default__.fm20((self).f6, (self).f4, globalState)).cf7)).cf7))
            d_1716_v65_: _dafny.Array
            def lambda65_(d_1717_v46_, d_1718_v0_, d_1719_p0_):
                def lambda66_(d_1720_i8_):
                    return (_dafny.MultiSet([d_1717_v46_, len(d_1718_v0_)])).ispropersubset(_dafny.MultiSet([d_1719_p0_, d_1717_v46_, d_1717_v46_]))

                return lambda66_

            init41_ = lambda65_(d_1694_v46_, d_1638_v0_, p0)
            nw268_ = _dafny.Array(None, 8)
            for i0_41_ in range(nw268_.length(0)):
                nw268_[i0_41_] = init41_(i0_41_)
            d_1716_v65_ = nw268_
            d_1716_v65_ = d_1716_v65_
            d_1677_v33_ = (d_1677_v33_).set(451, (self).f4)
            d_1721_v66_: bool
            d_1721_v66_ = True
            d_1722_v67_: _dafny.Map
            d_1722_v67_ = _dafny.Map({(self).f4: (self).f6})
            d_1721_v66_ = (len(d_1722_v67_)) in (_dafny.MultiSet([(self).f6]))
        d_1723_v68_: bool
        d_1723_v68_ = True
        d_1723_v68_ = (self).f4
        d_1724_v69_: int
        d_1724_v69_ = 678
        d_1724_v69_ = (0) - (((self).f6 if not((d_1723_v68_) == (d_1723_v68_)) else p0))
        d_1725_v70_: _dafny.Map
        d_1725_v70_ = _dafny.Map({d_1723_v68_: p0})
        r0 = d_1725_v70_
        return r0

    def m2(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_1726_v0_: _dafny.Seq
        d_1726_v0_ = _dafny.SeqWithoutIsStrInference([(self).f11, p0, (self).f4])
        d_1727_v1_: _dafny.Seq
        d_1727_v1_ = _dafny.SeqWithoutIsStrInference([(d_1726_v0_)[default__.safeIndex((self).f6, len(d_1726_v0_))], p0, (self).f4])
        d_1728_v2_: D3
        d_1728_v2_ = D3_DC8((d_1726_v0_).set(default__.safeIndex((self).f6, len(d_1726_v0_)), (self).f11))
        source32_ = d_1728_v2_
        if source32_.is_DC9:
            d_1729___mcc_h0_ = source32_.cf8
            d_1730___mcc_h1_ = source32_.cf9
            d_1731___mcc_h2_ = source32_.cf10
            d_1732_cf10_ = d_1731___mcc_h2_
            d_1733_cf9_ = d_1730___mcc_h1_
            d_1734_cf8_ = d_1729___mcc_h0_
            d_1735_v3_: _dafny.Array
            nw269_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_1735_v3_ = nw269_
            d_1736_v4_: _dafny.Array
            def lambda67_(d_1737_i0_):
                return not((self).f11)

            init42_ = lambda67_
            nw270_ = _dafny.Array(None, 18)
            for i0_42_ in range(nw270_.length(0)):
                nw270_[i0_42_] = init42_(i0_42_)
            d_1736_v4_ = nw270_
            index254_ = default__.safeIndex(874, (d_1736_v4_).length(0))
            (d_1736_v4_)[index254_] = (self).f11
            d_1738_v5_: _dafny.Seq
            d_1738_v5_ = _dafny.SeqWithoutIsStrInference([d_1735_v3_, d_1735_v3_, d_1735_v3_, d_1735_v3_, d_1735_v3_])
            d_1739_v6_: _dafny.MultiSet
            d_1739_v6_ = _dafny.MultiSet([d_1736_v4_, d_1736_v4_])
            d_1740_v7_: _dafny.Seq
            d_1740_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmjgkq"))
            index255_ = default__.safeIndex(874, (d_1736_v4_).length(0))
            rhs240_ = (self).fm9(d_1732_cf10_, (self).f11, d_1733_cf9_, globalState)
            rhs241_ = (d_1738_v5_)[default__.safeIndex(((d_1739_v6_)[d_1736_v4_] if (d_1736_v4_) in (d_1739_v6_) else len(d_1740_v7_)), len(d_1738_v5_))]
            rhs242_ = (0) - ((len(d_1740_v7_)) + (default__.safeModuloInt((self).fm9(True, (self).f11, d_1733_cf9_, globalState), (self).f6)))
            rhs243_ = (len(d_1740_v7_) if p0 else d_1733_cf9_)
            rhs244_ = d_1732_cf10_
            lhs149_ = d_1736_v4_
            lhs150_ = default__.safeIndex(874, (d_1736_v4_).length(0))
            d_1733_cf9_ = rhs240_
            d_1735_v3_ = rhs241_
            d_1733_cf9_ = rhs242_
            d_1733_cf9_ = rhs243_
            lhs149_[lhs150_] = rhs244_
            if (self).f4:
                d_1732_cf10_ = False
                arr7_ = self.f5
                index256_ = default__.safeIndex(333, (self.f5).length(0))
                arr7_[index256_] = d_1736_v4_
                arr8_ = self.f5
                index257_ = default__.safeIndex(333, (self.f5).length(0))
                arr8_[index257_] = (D2_DC5(d_1736_v4_)).cf5
                d_1741_v8_: D3
                d_1741_v8_ = D3_DC10(len(d_1740_v7_))
                d_1742_v9_: _dafny.Array
                nw271_ = _dafny.Array(None, 25)
                nw271_[int(0)] = (d_1741_v8_ if d_1732_cf10_ else d_1741_v8_)
                nw271_[int(1)] = d_1741_v8_
                nw271_[int(2)] = d_1741_v8_
                nw271_[int(3)] = d_1741_v8_
                nw271_[int(4)] = d_1741_v8_
                nw271_[int(5)] = d_1741_v8_
                nw271_[int(6)] = d_1741_v8_
                nw271_[int(7)] = d_1741_v8_
                nw271_[int(8)] = d_1741_v8_
                nw271_[int(9)] = d_1741_v8_
                nw271_[int(10)] = d_1741_v8_
                nw271_[int(11)] = D3_DC10(d_1733_cf9_)
                nw271_[int(12)] = d_1741_v8_
                nw271_[int(13)] = d_1741_v8_
                nw271_[int(14)] = D3_DC10((self).f6)
                nw271_[int(15)] = d_1741_v8_
                nw271_[int(16)] = D3_DC10((self).f6)
                nw271_[int(17)] = d_1741_v8_
                nw271_[int(18)] = (d_1741_v8_ if d_1732_cf10_ else d_1741_v8_)
                nw271_[int(19)] = d_1741_v8_
                nw271_[int(20)] = d_1741_v8_
                nw271_[int(21)] = d_1741_v8_
                nw271_[int(22)] = d_1741_v8_
                nw271_[int(23)] = d_1741_v8_
                nw271_[int(24)] = d_1741_v8_
                d_1742_v9_ = nw271_
                index258_ = default__.safeIndex(515, (d_1742_v9_).length(0))
                (d_1742_v9_)[index258_] = d_1741_v8_
                index259_ = default__.safeIndex(515, (d_1742_v9_).length(0))
                (d_1742_v9_)[index259_] = d_1741_v8_
                d_1733_cf9_ = default__.safeModuloInt(d_1733_cf9_, (self).f6)
                d_1743_v10_: _dafny.Map
                d_1743_v10_ = _dafny.Map({True: (d_1736_v4_)[default__.safeIndex(874, (d_1736_v4_).length(0))]})
                d_1744_v11_: _dafny.Seq
                d_1744_v11_ = _dafny.SeqWithoutIsStrInference([(d_1743_v10_).set(d_1732_cf10_, d_1734_cf8_)])
                d_1743_v10_ = ((d_1744_v11_)[default__.safeIndex((self).f6, len(d_1744_v11_))]) | (d_1743_v10_)
            elif True:
                d_1745_v12_: _dafny.Map
                d_1745_v12_ = _dafny.Map({not((d_1727_v1_)[default__.safeIndex(d_1733_cf9_, len(d_1727_v1_))]): (self).f6})
                d_1746_v13_: C0
                nw272_ = C0()
                nw272_.ctor__(((d_1745_v12_)[d_1732_cf10_] if (d_1732_cf10_) in (d_1745_v12_) else (self).f6))
                d_1746_v13_ = nw272_
                d_1747_v14_: D2
                d_1747_v14_ = D2_DC7(default__.safeModuloInt((self).f6, 433))
                d_1748_v15_: _dafny.Set
                d_1748_v15_ = _dafny.Set({(d_1746_v13_).f12})
                d_1749_v16_: _dafny.MultiSet
                d_1749_v16_ = _dafny.MultiSet([d_1748_v15_, d_1748_v15_, d_1748_v15_])
                d_1750_v17_: _dafny.Map
                d_1750_v17_ = _dafny.Map({d_1746_v13_: d_1736_v4_})
                d_1751_v18_: _dafny.Map
                d_1751_v18_ = _dafny.Map({376: len(d_1750_v17_)})
                d_1747_v14_ = (D2_DC7(len(d_1751_v18_)) if (_dafny.Set({(self).f6})) not in ((d_1749_v16_).set(d_1748_v15_, default__.abs((d_1746_v13_).f12))) else d_1747_v14_)
                d_1752_v19_: str
                d_1752_v19_ = _dafny.CodePoint('s')
                r0 = d_1752_v19_
                d_1732_cf10_ = (self).f4
                d_1733_cf9_ = d_1733_cf9_
            d_1753_v20_: C0
            nw273_ = C0()
            nw273_.ctor__(default__.safeDivisionInt(len(d_1726_v0_), (self).f6))
            d_1753_v20_ = nw273_
            index260_ = default__.safeIndex(874, (d_1736_v4_).length(0))
            (d_1736_v4_)[index260_] = (d_1753_v20_).fm17(p0, globalState)
        elif source32_.is_DC10:
            d_1754___mcc_h3_ = source32_.cf11
            d_1755_cf11_ = d_1754___mcc_h3_
            d_1756_v21_: _dafny.Map
            d_1756_v21_ = _dafny.Map({default__.safeModuloInt(137, d_1755_cf11_): ((self).f11 if (self).f11 else p0)})
            d_1756_v21_ = (d_1756_v21_).set(d_1755_cf11_, (self).f11)
            d_1757_v22_: bool
            d_1757_v22_ = False
            d_1758_v23_: _dafny.Array
            nw274_ = _dafny.Array(int(0), 26)
            d_1758_v23_ = nw274_
            d_1759_v24_: _dafny.Seq
            d_1759_v24_ = _dafny.SeqWithoutIsStrInference([d_1755_cf11_])
            rhs245_ = p0
            rhs246_ = d_1755_cf11_
            rhs247_ = (d_1759_v24_) != (d_1759_v24_)
            rhs248_ = d_1758_v23_
            d_1757_v22_ = rhs245_
            d_1755_cf11_ = rhs246_
            d_1757_v22_ = rhs247_
            d_1758_v23_ = rhs248_
            d_1760_v25_: _dafny.Seq
            d_1760_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gan"))
            d_1761_v26_: str
            d_1761_v26_ = _dafny.CodePoint('j')
            d_1760_v25_ = (d_1760_v25_ if (default__.fm0(d_1761_v26_, globalState) if (self).f4 else not((self).f4)) else (d_1760_v25_) + (d_1760_v25_))
            d_1762_v27_: _dafny.Array
            nw275_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_1762_v27_ = nw275_
            d_1762_v27_ = d_1762_v27_
        elif source32_.is_DC8:
            d_1763___mcc_h4_ = source32_.cf7
            d_1764_cf7_ = d_1763___mcc_h4_
            d_1765_v28_: _dafny.Array
            nw276_ = _dafny.Array(int(0), 7)
            d_1765_v28_ = nw276_
            d_1766_v29_: D1
            d_1766_v29_ = D1_DC4((0) - ((self).f6), (self).f6)
            d_1767_v30_: _dafny.Seq
            d_1767_v30_ = _dafny.SeqWithoutIsStrInference([(d_1766_v29_).cf3])
            index261_ = default__.safeIndex(538, (d_1765_v28_).length(0))
            (d_1765_v28_)[index261_] = (d_1767_v30_)[default__.safeIndex((self).f6, len(d_1767_v30_))]
            d_1768_v31_: _dafny.Map
            d_1768_v31_ = _dafny.Map({(self).f4: 476})
            d_1769_v32_: _dafny.MultiSet
            d_1769_v32_ = _dafny.MultiSet([d_1768_v31_, d_1768_v31_])
            d_1770_v33_: str
            d_1770_v33_ = _dafny.CodePoint('v')
            d_1771_v34_: _dafny.Seq
            d_1771_v34_ = _dafny.SeqWithoutIsStrInference([d_1770_v33_])
            d_1772_v35_: _dafny.MultiSet
            d_1772_v35_ = _dafny.MultiSet([p0, (self).fm8(45, (self).f11, d_1771_v34_, d_1767_v30_, globalState), p0, (self).f11, (self).f4])
            index262_ = default__.safeIndex(538, (d_1765_v28_).length(0))
            (d_1765_v28_)[index262_] = ((default__.fm21((0) - ((0) - ((self).f6)), d_1769_v32_, 193, (self).f6, globalState)) - (d_1772_v35_)).cardinality
            d_1773_v36_: C0
            nw277_ = C0()
            nw277_.ctor__((0) - (((d_1765_v28_)[default__.safeIndex(538, (d_1765_v28_).length(0))]) - ((d_1765_v28_)[default__.safeIndex(538, (d_1765_v28_).length(0))])))
            d_1773_v36_ = nw277_
            d_1774_v37_: _dafny.Array
            nw278_ = _dafny.Array(False, 27)
            d_1774_v37_ = nw278_
            index263_ = default__.safeIndex(189, (d_1774_v37_).length(0))
            (d_1774_v37_)[index263_] = ((self).f4) or ((self).f11)
            index264_ = default__.safeIndex(189, (d_1774_v37_).length(0))
            (d_1774_v37_)[index264_] = not((d_1726_v0_)[default__.safeIndex((175) + ((d_1765_v28_)[default__.safeIndex(538, (d_1765_v28_).length(0))]), len(d_1726_v0_))])
            d_1775_v38_: D4
            d_1775_v38_ = D4_DC13((d_1765_v28_)[default__.safeIndex(538, (d_1765_v28_).length(0))], (self).f11, len(d_1771_v34_), d_1727_v1_)
            d_1776_v39_: _dafny.Set
            d_1776_v39_ = _dafny.Set({(d_1775_v38_).cf15, p0, (self).f4})
            index265_ = default__.safeIndex(189, (d_1774_v37_).length(0))
            index266_ = default__.safeIndex(538, (d_1765_v28_).length(0))
            index267_ = default__.safeIndex(538, (d_1765_v28_).length(0))
            rhs249_ = (d_1727_v1_)[default__.safeIndex(len((d_1776_v39_) | (d_1776_v39_)), len(d_1727_v1_))]
            rhs250_ = ((self).f6 if ((self).f4 if p0 else p0) else (d_1773_v36_).f12)
            rhs251_ = (d_1767_v30_)[default__.safeIndex((self).f6, len(d_1767_v30_))]
            lhs151_ = d_1774_v37_
            lhs152_ = default__.safeIndex(189, (d_1774_v37_).length(0))
            lhs153_ = d_1765_v28_
            lhs154_ = default__.safeIndex(538, (d_1765_v28_).length(0))
            lhs155_ = d_1765_v28_
            lhs156_ = default__.safeIndex(538, (d_1765_v28_).length(0))
            lhs151_[lhs152_] = rhs249_
            lhs153_[lhs154_] = rhs250_
            lhs155_[lhs156_] = rhs251_
        elif True:
            d_1777___mcc_h5_ = source32_.cf12
            d_1778_cf12_ = d_1777___mcc_h5_
            d_1779_v40_: _dafny.Array
            def lambda68_(d_1780_i1_):
                return default__.safeModuloInt(d_1780_i1_, (self).f6)

            init43_ = lambda68_
            nw279_ = _dafny.Array(None, 19)
            for i0_43_ in range(nw279_.length(0)):
                nw279_[i0_43_] = init43_(i0_43_)
            d_1779_v40_ = nw279_
            index268_ = default__.safeIndex(30, (d_1779_v40_).length(0))
            (d_1779_v40_)[index268_] = 257
            index269_ = default__.safeIndex(30, (d_1779_v40_).length(0))
            (d_1779_v40_)[index269_] = (self).f6
            index270_ = default__.safeIndex(30, (d_1779_v40_).length(0))
            (d_1779_v40_)[index270_] = (self).f6
            d_1781_v41_: bool
            d_1781_v41_ = False
            d_1782_v42_: _dafny.Array
            nw280_ = _dafny.Array(False, 26)
            d_1782_v42_ = nw280_
            d_1783_v43_: _dafny.MultiSet
            d_1783_v43_ = _dafny.MultiSet([p0, not(d_1781_v41_), (self).f4])
            d_1784_v44_: _dafny.Map
            d_1784_v44_ = _dafny.Map({(self).f11: d_1783_v43_})
            d_1785_v45_: _dafny.MultiSet
            d_1785_v45_ = _dafny.MultiSet([d_1783_v43_, _dafny.MultiSet(d_1727_v1_), ((d_1784_v44_)[d_1781_v41_] if (d_1781_v41_) in (d_1784_v44_) else d_1783_v43_), d_1783_v43_])
            index271_ = default__.safeIndex(420, (d_1782_v42_).length(0))
            (d_1782_v42_)[index271_] = (d_1785_v45_).issubset(d_1785_v45_)
            index272_ = default__.safeIndex(30, (d_1779_v40_).length(0))
            index273_ = default__.safeIndex(420, (d_1782_v42_).length(0))
            index274_ = default__.safeIndex(30, (d_1779_v40_).length(0))
            rhs252_ = True
            rhs253_ = (d_1779_v40_)[default__.safeIndex(30, (d_1779_v40_).length(0))]
            rhs254_ = (self).f11
            rhs255_ = default__.safeModuloInt((d_1779_v40_)[default__.safeIndex(30, (d_1779_v40_).length(0))], ((self).f6 if (self).f11 else (d_1779_v40_)[default__.safeIndex(30, (d_1779_v40_).length(0))]))
            lhs157_ = d_1779_v40_
            lhs158_ = default__.safeIndex(30, (d_1779_v40_).length(0))
            lhs159_ = d_1782_v42_
            lhs160_ = default__.safeIndex(420, (d_1782_v42_).length(0))
            lhs161_ = d_1779_v40_
            lhs162_ = default__.safeIndex(30, (d_1779_v40_).length(0))
            d_1781_v41_ = rhs252_
            lhs157_[lhs158_] = rhs253_
            lhs159_[lhs160_] = rhs254_
            lhs161_[lhs162_] = rhs255_
            if True:
                d_1786_v46_: D0
                d_1786_v46_ = D0_DC2(len(_dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f6, (self).f6})])))
                d_1787_v47_: _dafny.MultiSet
                d_1787_v47_ = _dafny.MultiSet([(d_1779_v40_)[default__.safeIndex(30, (d_1779_v40_).length(0))], (d_1786_v46_).cf1])
                d_1788_v48_: D3
                d_1788_v48_ = D3_DC9((self).f11, (d_1787_v47_).cardinality, (self).f11)
                d_1789_v49_: _dafny.Seq
                d_1789_v49_ = _dafny.SeqWithoutIsStrInference([(d_1788_v48_).cf9, (d_1779_v40_)[default__.safeIndex(30, (d_1779_v40_).length(0))]])
                d_1790_v50_: bool
                d_1791_v51_: int
                out51_: bool
                out52_: int
                out51_, out52_ = default__.m0(default__.fm22(globalState), (p0) == ((self).f4), (_dafny.MultiSet(d_1789_v49_)).intersection(d_1787_v47_), d_1781_v41_, globalState)
                d_1790_v50_ = out51_
                d_1791_v51_ = out52_
                d_1792_v52_: _dafny.Map
                d_1792_v52_ = _dafny.Map({((d_1779_v40_)[default__.safeIndex(30, (d_1779_v40_).length(0))]) + (d_1791_v51_): 200})
                d_1792_v52_ = (d_1792_v52_).set(default__.fm3(d_1791_v51_, d_1781_v41_, globalState), (d_1789_v49_)[default__.safeIndex(-562, len(d_1789_v49_))])
                d_1793_v53_: _dafny.Seq
                d_1793_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owcvyfr"))
                d_1794_v54_: _dafny.Map
                d_1794_v54_ = _dafny.Map({(len(d_1793_v53_)) * (d_1791_v51_): d_1726_v0_})
                d_1794_v54_ = (d_1794_v54_).set(d_1791_v51_, _dafny.SeqWithoutIsStrInference([(d_1782_v42_)[default__.safeIndex(420, (d_1782_v42_).length(0))], (self).f11, (self).f11, p0]))
                d_1795_v55_: str
                d_1795_v55_ = _dafny.CodePoint('d')
                index275_ = default__.safeIndex(420, (d_1782_v42_).length(0))
                (d_1782_v42_)[index275_] = default__.fm0(d_1795_v55_, globalState)
                d_1796_v56_: _dafny.Set
                d_1796_v56_ = _dafny.Set({(d_1782_v42_)[default__.safeIndex(420, (d_1782_v42_).length(0))]})
                d_1797_v57_: _dafny.Seq
                d_1797_v57_ = _dafny.SeqWithoutIsStrInference([d_1779_v40_, d_1779_v40_])
                rhs256_ = default__.fm23(d_1791_v51_, (self).f6, globalState)
                rhs257_ = d_1797_v57_
                rhs258_ = d_1781_v41_
                d_1796_v56_ = rhs256_
                d_1797_v57_ = rhs257_
                d_1790_v50_ = rhs258_
            elif True:
                d_1781_v41_ = not((self).f11)
                d_1798_v58_: _dafny.Map
                d_1798_v58_ = _dafny.Map({(d_1783_v43_).cardinality: True})
                d_1799_v59_: _dafny.MultiSet
                d_1799_v59_ = _dafny.MultiSet([d_1798_v58_])
                d_1800_v60_: _dafny.Seq
                d_1800_v60_ = _dafny.SeqWithoutIsStrInference([d_1798_v58_, d_1798_v58_])
                d_1801_v61_: _dafny.Seq
                d_1801_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgpjpb"))
                d_1802_v62_: _dafny.Map
                d_1802_v62_ = _dafny.Map({(self).f6: (d_1799_v59_) | (_dafny.MultiSet([d_1798_v58_, d_1798_v58_, d_1798_v58_, (d_1800_v60_)[default__.safeIndex(len(d_1801_v61_), len(d_1800_v60_))], d_1798_v58_]))})
                d_1802_v62_ = (d_1802_v62_).set((d_1779_v40_)[default__.safeIndex(30, (d_1779_v40_).length(0))], d_1799_v59_)
                index276_ = default__.safeIndex(30, (d_1779_v40_).length(0))
                (d_1779_v40_)[index276_] = ((self).f6) - ((self).f6)
                index277_ = default__.safeIndex(420, (d_1782_v42_).length(0))
                (d_1782_v42_)[index277_] = (803) >= (((self).fm9(p0, p0, (self).fm15(globalState), globalState)) * (((d_1783_v43_)[p0] if (p0) in (d_1783_v43_) else (self).f6)))
                d_1803_v63_: str
                d_1803_v63_ = _dafny.CodePoint('e')
                r0 = d_1803_v63_
        d_1804_i2_: int
        d_1804_i2_ = 0
        with _dafny.label("8"):
            while (self).f4:
                with _dafny.c_label("8"):
                    if (d_1804_i2_) >= (100):
                        raise _dafny.Break("8")
                    d_1804_i2_ = (d_1804_i2_) + (1)
                    d_1805_v64_: _dafny.Array
                    def lambda69_(d_1806_p0_):
                        def lambda70_(d_1807_i3_):
                            return d_1806_p0_

                        return lambda70_

                    init44_ = lambda69_(p0)
                    nw281_ = _dafny.Array(None, 21)
                    for i0_44_ in range(nw281_.length(0)):
                        nw281_[i0_44_] = init44_(i0_44_)
                    d_1805_v64_ = nw281_
                    d_1805_v64_ = d_1805_v64_
                    d_1808_v65_: bool
                    d_1808_v65_ = True
                    d_1808_v65_ = (self).f4
                    d_1809_v66_: _dafny.MultiSet
                    d_1809_v66_ = _dafny.MultiSet([126, (self).f6])
                    d_1810_v67_: D0
                    d_1810_v67_ = D0_DC2((d_1809_v66_).cardinality)
                    d_1811_v68_: _dafny.Map
                    d_1811_v68_ = _dafny.Map({(self).f11: (d_1810_v67_).cf1})
                    d_1811_v68_ = (d_1811_v68_).set((self).f4, ((self).f6) + ((self).f6))
                    index278_ = default__.safeIndex(910, (d_1805_v64_).length(0))
                    (d_1805_v64_)[index278_] = not(p0)
                    index279_ = default__.safeIndex(910, (d_1805_v64_).length(0))
                    (d_1805_v64_)[index279_] = p0
                    pass
            pass
        d_1812_v69_: _dafny.Map
        d_1812_v69_ = _dafny.Map({(self).f6: (self).f6})
        d_1812_v69_ = (d_1812_v69_) | ((d_1812_v69_) | (d_1812_v69_))
        d_1726_v0_ = (d_1726_v0_) + (d_1727_v1_)
        hi12_ = default__.safeModuloInt((self).f6, (self).f6)
        for d_1813_i4_ in range((self).f6, hi12_):
            d_1814_v70_: int
            d_1814_v70_ = 577
            d_1814_v70_ = (self).f6
            d_1815_v71_: D2
            d_1815_v71_ = D2_DC7((self).f6)
            source33_ = d_1815_v71_
            if source33_.is_DC6:
                d_1816_v72_: bool
                d_1817_v73_: bool
                d_1818_v74_: D1
                out53_: bool
                out54_: bool
                out55_: D1
                out53_, out54_, out55_ = (self).m8(default__.safeModuloInt((self).f6, d_1814_v70_), d_1814_v70_, globalState)
                d_1816_v72_ = out53_
                d_1817_v73_ = out54_
                d_1818_v74_ = out55_
                d_1819_v75_: str
                d_1819_v75_ = _dafny.CodePoint('j')
                d_1820_v76_: _dafny.Seq
                d_1820_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftxvovkv"))
                d_1821_v77_: D0
                d_1821_v77_ = D0_DC0(len(d_1820_v76_))
                d_1822_v78_: _dafny.Map
                d_1822_v78_ = _dafny.Map({d_1821_v77_: d_1816_v72_})
                d_1823_v79_: _dafny.Map
                d_1823_v79_ = _dafny.Map({default__.fm0(d_1819_v75_, globalState): d_1822_v78_})
                d_1824_v80_: _dafny.Array
                nw282_ = _dafny.Array(None, 5)
                nw282_[int(0)] = d_1814_v70_
                nw282_[int(1)] = (0) - (len(((d_1823_v79_)[not(p0)] if (not(p0)) in (d_1823_v79_) else _dafny.Map({D0_DC0(d_1814_v70_): (self).f4}))))
                nw282_[int(2)] = (self).f6
                nw282_[int(3)] = -352
                nw282_[int(4)] = (self).f6
                d_1824_v80_ = nw282_
                d_1825_v81_: _dafny.Set
                d_1825_v81_ = _dafny.Set({d_1824_v80_})
                d_1816_v72_ = (d_1825_v81_).isdisjoint(d_1825_v81_)
                d_1812_v69_ = (d_1812_v69_).set((self).f6, (self).f6)
                d_1826_v82_: D0
                d_1826_v82_ = D0_DC2(d_1814_v70_)
                d_1826_v82_ = d_1826_v82_
            elif source33_.is_DC7:
                d_1827___mcc_h6_ = source33_.cf6
                d_1828_cf6_ = d_1827___mcc_h6_
                d_1829_v83_: _dafny.Array
                def lambda71_(d_1830_p0_):
                    def lambda72_(d_1831_i5_):
                        return d_1830_p0_

                    return lambda72_

                init45_ = lambda71_(p0)
                nw283_ = _dafny.Array(None, 27)
                for i0_45_ in range(nw283_.length(0)):
                    nw283_[i0_45_] = init45_(i0_45_)
                d_1829_v83_ = nw283_
                d_1832_v84_: str
                d_1832_v84_ = _dafny.CodePoint('w')
                d_1833_v85_: _dafny.Map
                d_1833_v85_ = _dafny.Map({d_1832_v84_: d_1813_i4_})
                d_1834_v87_: _dafny.Seq
                d_1834_v87_ = _dafny.SeqWithoutIsStrInference([d_1832_v84_, d_1832_v84_])
                d_1835_v88_: _dafny.Seq
                def iife151_():
                    coll55_ = _dafny.Map()
                    compr_55_: str
                    for compr_55_ in (d_1834_v87_).Elements:
                        d_1836_v86_: str = compr_55_
                        if (d_1836_v86_) in (d_1834_v87_):
                            coll55_[d_1836_v86_] = d_1814_v70_
                    return _dafny.Map(coll55_)
                d_1835_v88_ = _dafny.SeqWithoutIsStrInference([iife151_()
                ])
                d_1837_v89_: _dafny.Map
                d_1837_v89_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1829_v83_]): (_dafny.SeqWithoutIsStrInference([d_1833_v85_, d_1833_v85_])) + ((d_1835_v88_).set(default__.safeIndex(d_1828_cf6_, len(d_1835_v88_)), _dafny.Map({default__.fm1(default__.fm3(d_1828_cf6_, (self).f11, globalState), (self).f11, globalState): (0) - ((self).f6)})))})
                d_1838_v90_: _dafny.Seq
                d_1838_v90_ = _dafny.SeqWithoutIsStrInference([d_1829_v83_, d_1829_v83_, d_1829_v83_])
                d_1839_v91_: _dafny.Seq
                d_1839_v91_ = _dafny.SeqWithoutIsStrInference([d_1829_v83_, d_1829_v83_, (d_1838_v90_)[default__.safeIndex(-553, len(d_1838_v90_))], (D2_DC5(d_1829_v83_)).cf5, d_1829_v83_])
                d_1837_v89_ = (d_1837_v89_).set(d_1838_v90_, _dafny.SeqWithoutIsStrInference([d_1833_v85_ for d_1840_i6_ in range(default__.abs(707))]))
                d_1841_v92_: _dafny.Map
                d_1841_v92_ = _dafny.Map({d_1834_v87_: (self).f6})
                d_1841_v92_ = (d_1841_v92_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpjli")), d_1828_cf6_)
                d_1842_v93_: _dafny.Map
                d_1842_v93_ = _dafny.Map({p0: p0})
                r0 = default__.fm1(len((d_1842_v93_) | (_dafny.Map({False: False}))), (self).f4, globalState)
                d_1843_v94_: _dafny.Array
                nw284_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_1843_v94_ = nw284_
                d_1844_v95_: _dafny.Array
                def lambda73_(d_1845_i7_):
                    return (d_1845_i7_) - (964)

                init46_ = lambda73_
                nw285_ = _dafny.Array(None, 1)
                for i0_46_ in range(nw285_.length(0)):
                    nw285_[i0_46_] = init46_(i0_46_)
                d_1844_v95_ = nw285_
                index280_ = default__.safeIndex(906, (d_1843_v94_).length(0))
                (d_1843_v94_)[index280_] = d_1844_v95_
                index281_ = default__.safeIndex(906, (d_1843_v94_).length(0))
                (d_1843_v94_)[index281_] = d_1844_v95_
            elif True:
                d_1846___mcc_h7_ = source33_.cf5
                d_1847_cf5_ = d_1846___mcc_h7_
                d_1848_v96_: C0
                nw286_ = C0()
                nw286_.ctor__((self).f6)
                d_1848_v96_ = nw286_
                d_1849_v97_: _dafny.MultiSet
                d_1849_v97_ = _dafny.MultiSet([(d_1848_v96_).f12, (d_1848_v96_).f12, (self).f6, d_1814_v70_])
                d_1850_v98_: _dafny.Map
                d_1850_v98_ = _dafny.Map({d_1849_v97_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "synxuae"))})
                d_1851_v99_: bool
                d_1852_v100_: int
                out56_: bool
                out57_: int
                out56_, out57_ = default__.m0(d_1850_v98_, (self).f11, _dafny.MultiSet([d_1814_v70_]), (self).f11, globalState)
                d_1851_v99_ = out56_
                d_1852_v100_ = out57_
                d_1853_v101_: _dafny.Seq
                d_1853_v101_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adljl"))
                d_1854_v102_: _dafny.Set
                d_1854_v102_ = _dafny.Set({(d_1853_v101_ if (d_1726_v0_)[default__.safeIndex(d_1814_v70_, len(d_1726_v0_))] else d_1853_v101_)})
                d_1854_v102_ = d_1854_v102_
                d_1855_v103_: _dafny.Map
                d_1855_v103_ = _dafny.Map({not(not(p0)): (0) - (d_1814_v70_)})
                d_1856_v104_: _dafny.Seq
                d_1856_v104_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1851_v99_: (d_1848_v96_).f12}), d_1855_v103_, d_1855_v103_, d_1855_v103_])
                d_1857_v105_: C0
                nw287_ = C0()
                nw287_.ctor__(len(d_1856_v104_))
                d_1857_v105_ = nw287_
            d_1858_v106_: _dafny.Seq
            d_1858_v106_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tasx"))
            d_1859_v107_: _dafny.Array
            nw288_ = _dafny.Array(int(0), 17)
            d_1859_v107_ = nw288_
            index282_ = default__.safeIndex(289, (d_1859_v107_).length(0))
            (d_1859_v107_)[index282_] = 415
            index283_ = default__.safeIndex(289, (d_1859_v107_).length(0))
            rhs259_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xohxe"))
            rhs260_ = len(d_1858_v106_)
            lhs163_ = d_1859_v107_
            lhs164_ = default__.safeIndex(289, (d_1859_v107_).length(0))
            d_1858_v106_ = rhs259_
            lhs163_[lhs164_] = rhs260_
            index284_ = default__.safeIndex(289, (d_1859_v107_).length(0))
            (d_1859_v107_)[index284_] = d_1814_v70_
        d_1860_v108_: str
        d_1860_v108_ = _dafny.CodePoint('h')
        if default__.fm0(d_1860_v108_, globalState):
            d_1861_v109_: int
            d_1861_v109_ = 639
            d_1862_v110_: _dafny.MultiSet
            d_1862_v110_ = _dafny.MultiSet([(self).f4])
            d_1861_v109_ = ((self).f6) * (len(_dafny.SeqWithoutIsStrInference([d_1861_v109_, (d_1862_v110_).cardinality, (self).f6])))
            d_1863_v111_: bool
            d_1863_v111_ = False
            d_1864_v112_: _dafny.Seq
            d_1864_v112_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
            d_1865_v113_: _dafny.Map
            d_1865_v113_ = _dafny.Map({len(d_1864_v112_): d_1863_v111_})
            d_1863_v111_ = ((d_1865_v113_)[d_1861_v109_] if (d_1861_v109_) in (d_1865_v113_) else (self).f11)
            d_1866_v114_: _dafny.Map
            d_1866_v114_ = _dafny.Map({d_1726_v0_: d_1863_v111_})
            d_1866_v114_ = (d_1866_v114_).set(_dafny.SeqWithoutIsStrInference([p0]), p0)
            d_1867_v115_: _dafny.MultiSet
            d_1867_v115_ = _dafny.MultiSet([default__.fm19(p0, d_1861_v109_, globalState)])
            d_1868_v116_: D1
            d_1868_v116_ = D1_DC3(d_1863_v111_)
            d_1869_v117_: _dafny.Seq
            d_1869_v117_ = _dafny.SeqWithoutIsStrInference([d_1726_v0_, d_1726_v0_])
            d_1870_v118_: _dafny.Array
            nw289_ = _dafny.Array(None, 22)
            nw289_[int(0)] = d_1867_v115_
            nw289_[int(1)] = _dafny.MultiSet([d_1727_v1_, d_1726_v0_, _dafny.SeqWithoutIsStrInference([(self).f11]), d_1726_v0_])
            nw289_[int(2)] = d_1867_v115_
            nw289_[int(3)] = d_1867_v115_
            nw289_[int(4)] = _dafny.MultiSet([d_1727_v1_])
            nw289_[int(5)] = d_1867_v115_
            nw289_[int(6)] = d_1867_v115_
            nw289_[int(7)] = _dafny.MultiSet([d_1727_v1_, _dafny.SeqWithoutIsStrInference([(self).f11, True]), d_1726_v0_, d_1726_v0_])
            nw289_[int(8)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1726_v0_ for d_1871_i8_ in range(default__.abs(146))])) if (d_1868_v116_).cf2 else default__.fm24(d_1861_v109_, d_1860_v108_, d_1861_v109_, globalState))
            nw289_[int(9)] = d_1867_v115_
            nw289_[int(10)] = (d_1867_v115_).intersection(_dafny.MultiSet(d_1869_v117_))
            nw289_[int(11)] = d_1867_v115_
            nw289_[int(12)] = (d_1867_v115_).intersection(d_1867_v115_)
            nw289_[int(13)] = d_1867_v115_
            nw289_[int(14)] = (d_1867_v115_).intersection(d_1867_v115_)
            nw289_[int(15)] = (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f4])])).intersection(d_1867_v115_)
            nw289_[int(16)] = d_1867_v115_
            nw289_[int(17)] = d_1867_v115_
            nw289_[int(18)] = d_1867_v115_
            nw289_[int(19)] = d_1867_v115_
            nw289_[int(20)] = d_1867_v115_
            nw289_[int(21)] = d_1867_v115_
            d_1870_v118_ = nw289_
            d_1872_v119_: _dafny.Map
            d_1872_v119_ = _dafny.Map({d_1863_v111_: _dafny.SeqWithoutIsStrInference([d_1727_v1_])})
            d_1873_v120_: D6
            d_1873_v120_ = D6_DC15(d_1867_v115_)
            index285_ = default__.safeIndex(270, (d_1870_v118_).length(0))
            (d_1870_v118_)[index285_] = (_dafny.MultiSet(((d_1872_v119_)[p0] if (p0) in (d_1872_v119_) else d_1869_v117_))) | ((d_1873_v120_).cf19)
            d_1874_v121_: _dafny.Seq
            d_1874_v121_ = _dafny.SeqWithoutIsStrInference([d_1861_v109_, (0) - (d_1861_v109_)])
            d_1875_v122_: _dafny.MultiSet
            d_1875_v122_ = _dafny.MultiSet([d_1874_v121_])
            d_1876_v123_: _dafny.Array
            nw290_ = _dafny.Array(_dafny.Seq({}), 1)
            d_1876_v123_ = nw290_
            index286_ = default__.safeIndex(270, (d_1870_v118_).length(0))
            rhs261_ = (d_1867_v115_).set(_dafny.SeqWithoutIsStrInference([False]), default__.abs((self).f6))
            rhs262_ = d_1875_v122_
            rhs263_ = d_1876_v123_
            rhs264_ = (self).f11
            lhs165_ = d_1870_v118_
            lhs166_ = default__.safeIndex(270, (d_1870_v118_).length(0))
            lhs165_[lhs166_] = rhs261_
            d_1875_v122_ = rhs262_
            d_1876_v123_ = rhs263_
            d_1863_v111_ = rhs264_
            d_1861_v109_ = default__.safeDivisionInt(313, d_1861_v109_)
        elif True:
            d_1877_v124_: D3
            d_1877_v124_ = D3_DC10((self).f6)
            d_1877_v124_ = D3_DC10((self).f6)
            d_1878_v125_: int
            d_1878_v125_ = 216
            d_1878_v125_ = (((d_1812_v69_)[-103] if (-103) in (d_1812_v69_) else d_1878_v125_) if p0 else len(((d_1726_v0_).set(default__.safeIndex(d_1878_v125_, len(d_1726_v0_)), default__.fm0(d_1860_v108_, globalState))).set(default__.safeIndex(948, len((d_1726_v0_).set(default__.safeIndex(d_1878_v125_, len(d_1726_v0_)), default__.fm0(d_1860_v108_, globalState)))), True)))
            source34_ = D3_DC8(d_1726_v0_)
            if source34_.is_DC9:
                d_1879___mcc_h8_ = source34_.cf8
                d_1880___mcc_h9_ = source34_.cf9
                d_1881___mcc_h10_ = source34_.cf10
                d_1882_cf10_ = d_1881___mcc_h10_
                d_1883_cf9_ = d_1880___mcc_h9_
                d_1884_cf8_ = d_1879___mcc_h8_
                d_1885_v126_: _dafny.MultiSet
                d_1885_v126_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_1860_v108_ for d_1886_i9_ in range(default__.abs(676))])), len(_dafny.Set({(self).f11, not((self).f4), d_1884_cf8_})), len(default__.fm25(d_1883_cf9_, (self).f6, (self).f4, (self).f11, globalState)), d_1878_v125_])
                d_1887_v127_: _dafny.Seq
                d_1887_v127_ = _dafny.SeqWithoutIsStrInference([d_1885_v126_])
                d_1888_v128_: _dafny.Seq
                d_1888_v128_ = _dafny.SeqWithoutIsStrInference([d_1883_cf9_, (self).f6, d_1878_v125_])
                d_1885_v126_ = (d_1887_v127_)[default__.safeIndex(len((_dafny.Map({d_1878_v125_: d_1888_v128_})).set((self).f6, _dafny.SeqWithoutIsStrInference([d_1883_cf9_ for d_1889_i10_ in range(default__.abs(807))]))), len(d_1887_v127_))]
                d_1890_v129_: _dafny.Set
                d_1890_v129_ = _dafny.Set({d_1878_v125_})
                d_1883_cf9_ = (self).fm9((self).f11, d_1884_cf8_, len((d_1890_v129_)), globalState)
                d_1891_v130_: _dafny.Array
                def lambda74_(d_1892_i11_):
                    return (d_1892_i11_) * ((self).f6)

                init47_ = lambda74_
                nw291_ = _dafny.Array(None, 8)
                for i0_47_ in range(nw291_.length(0)):
                    nw291_[i0_47_] = init47_(i0_47_)
                d_1891_v130_ = nw291_
                d_1893_v131_: _dafny.Array
                def lambda75_(d_1894_i12_):
                    return D0_DC0((self).f6)

                init48_ = lambda75_
                nw292_ = _dafny.Array(None, 10)
                for i0_48_ in range(nw292_.length(0)):
                    nw292_[i0_48_] = init48_(i0_48_)
                d_1893_v131_ = nw292_
                d_1895_v132_: D0
                d_1895_v132_ = D0_DC0(d_1883_cf9_)
                index287_ = default__.safeIndex(258, (d_1893_v131_).length(0))
                (d_1893_v131_)[index287_] = d_1895_v132_
                d_1896_v133_: _dafny.Array
                nw293_ = _dafny.Array(False, 27)
                d_1896_v133_ = nw293_
                index288_ = default__.safeIndex(258, (d_1893_v131_).length(0))
                rhs265_ = d_1877_v124_
                rhs266_ = d_1891_v130_
                rhs267_ = (0) - ((0) - (((0) - (default__.safeDivisionInt(533, d_1883_cf9_))) + ((d_1883_cf9_ if d_1882_cf10_ else -614))))
                rhs268_ = d_1895_v132_
                rhs269_ = d_1896_v133_
                lhs167_ = d_1893_v131_
                lhs168_ = default__.safeIndex(258, (d_1893_v131_).length(0))
                d_1877_v124_ = rhs265_
                d_1891_v130_ = rhs266_
                d_1883_cf9_ = rhs267_
                lhs167_[lhs168_] = rhs268_
                d_1896_v133_ = rhs269_
                d_1878_v125_ = (0) - ((0) - (d_1883_cf9_))
            elif source34_.is_DC10:
                d_1897___mcc_h11_ = source34_.cf11
                d_1898_cf11_ = d_1897___mcc_h11_
                d_1899_v134_: _dafny.Array
                def lambda76_(d_1900_i13_):
                    return (d_1900_i13_) + ((self).f6)

                init49_ = lambda76_
                nw294_ = _dafny.Array(None, 19)
                for i0_49_ in range(nw294_.length(0)):
                    nw294_[i0_49_] = init49_(i0_49_)
                d_1899_v134_ = nw294_
                d_1901_v135_: _dafny.Map
                d_1901_v135_ = _dafny.Map({d_1899_v134_: p0})
                d_1901_v135_ = (((d_1901_v135_).set(d_1899_v134_, (self).f4)).set(d_1899_v134_, (self).f4))
                d_1902_v136_: bool
                d_1902_v136_ = True
                rhs270_ = (self).f6
                rhs271_ = ((self).f6) + (d_1898_cf11_)
                rhs272_ = ((D3_DC8(d_1726_v0_)).cf7) != (_dafny.SeqWithoutIsStrInference([p0, p0, (self).f11, d_1902_v136_]))
                d_1898_cf11_ = rhs270_
                d_1898_cf11_ = rhs271_
                d_1902_v136_ = rhs272_
                d_1903_v137_: _dafny.Seq
                d_1903_v137_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjydewlp"))
                d_1904_v138_: T1
                nw295_ = C3()
                nw295_.ctor__(self.f5, len(d_1903_v137_), False)
                d_1904_v138_ = nw295_
                d_1905_v139_: _dafny.Map
                d_1905_v139_ = _dafny.Map({d_1902_v136_: d_1904_v138_})
                d_1906_v140_: _dafny.Map
                d_1906_v140_ = _dafny.Map({(self).f6: d_1904_v138_})
                d_1905_v139_ = (d_1905_v139_).set(((self).f11 if p0 else (d_1904_v138_).f4), ((d_1906_v140_)[(d_1904_v138_).f6] if ((d_1904_v138_).f6) in (d_1906_v140_) else d_1904_v138_))
                d_1878_v125_ = d_1878_v125_
            elif source34_.is_DC8:
                d_1907___mcc_h12_ = source34_.cf7
                d_1908_cf7_ = d_1907___mcc_h12_
                d_1909_v141_: _dafny.Array
                nw296_ = _dafny.Array(False, 11)
                d_1909_v141_ = nw296_
                d_1910_v142_: _dafny.Seq
                d_1910_v142_ = _dafny.SeqWithoutIsStrInference([d_1909_v141_])
                d_1911_v143_: _dafny.Seq
                d_1911_v143_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxv"))
                d_1912_v144_: D4
                d_1912_v144_ = D4_DC13((0) - (d_1878_v125_), (d_1909_v141_) not in (d_1910_v142_), default__.safeModuloInt((0) - (len(d_1911_v143_)), (self).f6), d_1726_v0_)
                d_1913_v145_: bool
                d_1913_v145_ = False
                d_1914_v146_: _dafny.Seq
                d_1914_v146_ = _dafny.SeqWithoutIsStrInference([(self).f6, d_1878_v125_])
                rhs273_ = ((self).f6) + ((self).f6)
                rhs274_ = default__.fm63(globalState)
                rhs275_ = (d_1914_v146_)[default__.safeIndex(d_1878_v125_, len(d_1914_v146_))]
                rhs276_ = (self).f4
                d_1878_v125_ = rhs273_
                d_1912_v144_ = rhs274_
                d_1878_v125_ = rhs275_
                d_1913_v145_ = rhs276_
                d_1915_v147_: _dafny.MultiSet
                d_1915_v147_ = _dafny.MultiSet([(self).f11])
                d_1913_v145_ = (not((self).f11)) == ((d_1915_v147_).ispropersubset(d_1915_v147_))
                d_1916_v148_: _dafny.Map
                d_1916_v148_ = _dafny.Map({d_1860_v108_: (d_1911_v143_) + (d_1911_v143_)})
                d_1916_v148_ = d_1916_v148_
                d_1917_v150_: _dafny.Array
                def lambda77_(d_1918_v125_):
                    def lambda78_(d_1919_i14_):
                        def iife152_():
                            coll56_ = _dafny.Set()
                            compr_56_: int
                            for compr_56_ in _dafny.IntegerRange(528, 26):
                                d_1920_v149_: int = compr_56_
                                if ((528) <= (d_1920_v149_)) and ((d_1920_v149_) < (26)):
                                    coll56_ = coll56_.union(_dafny.Set([default__.safeDivisionInt(d_1920_v149_, d_1918_v125_)]))
                            return _dafny.Set(coll56_)
                        return _dafny.MultiSet([iife152_()
                        , _dafny.Set({(self).f6})])

                    return lambda78_

                init50_ = lambda77_(d_1878_v125_)
                nw297_ = _dafny.Array(None, 11)
                for i0_50_ in range(nw297_.length(0)):
                    nw297_[i0_50_] = init50_(i0_50_)
                d_1917_v150_ = nw297_
                d_1921_v151_: _dafny.Map
                d_1921_v151_ = _dafny.Map({-561: (D3_DC9(d_1913_v145_, (self).fm15(globalState), (self).f11)).cf8})
                d_1922_v152_: _dafny.Set
                d_1922_v152_ = _dafny.Set({len(d_1921_v151_)})
                d_1923_v153_: _dafny.MultiSet
                d_1923_v153_ = _dafny.MultiSet([d_1922_v152_, d_1922_v152_])
                index289_ = default__.safeIndex(763, (d_1917_v150_).length(0))
                (d_1917_v150_)[index289_] = d_1923_v153_
                d_1924_v154_: _dafny.Seq
                d_1924_v154_ = _dafny.SeqWithoutIsStrInference([d_1922_v152_, d_1922_v152_, d_1922_v152_])
                index290_ = default__.safeIndex(763, (d_1917_v150_).length(0))
                (d_1917_v150_)[index290_] = (d_1923_v153_).set(((d_1924_v154_)[default__.safeIndex(len(d_1911_v143_), len(d_1924_v154_))]) - (d_1922_v152_), default__.abs((0) - (default__.safeModuloInt(len(d_1914_v146_), d_1878_v125_))))
            elif True:
                d_1925___mcc_h13_ = source34_.cf12
                d_1926_cf12_ = d_1925___mcc_h13_
                d_1878_v125_ = (self).fm15(globalState)
                d_1927_v155_: _dafny.Map
                d_1927_v155_ = _dafny.Map({(self).f11: 725})
                d_1928_v156_: _dafny.MultiSet
                d_1928_v156_ = _dafny.MultiSet([d_1927_v155_, default__.fm44(globalState), d_1927_v155_, d_1927_v155_, d_1927_v155_])
                d_1928_v156_ = (d_1928_v156_) | ((d_1928_v156_) | ((_dafny.MultiSet([d_1927_v155_])).set(d_1927_v155_, default__.abs(default__.fm3(d_1878_v125_, p0, globalState)))))
                d_1860_v108_ = d_1860_v108_
                d_1927_v155_ = d_1927_v155_
            d_1878_v125_ = ((0) - ((self).f6) if p0 else ((self).f6) + ((self).f6))
            d_1929_v157_: bool
            d_1929_v157_ = True
            d_1929_v157_ = p0
        d_1930_v158_: _dafny.Seq
        d_1930_v158_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehhyianjb"))
        d_1931_v159_: C2
        nw298_ = C2()
        nw298_.ctor__(d_1930_v158_, p0)
        d_1931_v159_ = nw298_
        d_1932_v160_: _dafny.Map
        d_1932_v160_ = _dafny.Map({692: d_1931_v159_})
        r0 = default__.fm1(default__.fm3((self).f6, p0, globalState), ((self).fm7(globalState)) <= (len(d_1932_v160_)), globalState)
        d_1933_v161_: D1
        d_1933_v161_ = D1_DC3((self).f4)
        d_1934_v162_: _dafny.Map
        d_1934_v162_ = _dafny.Map({(self).f11: p0})
        d_1935_v163_: _dafny.Array
        nw299_ = _dafny.Array(None, 29)
        nw299_[int(0)] = d_1933_v161_
        nw299_[int(1)] = d_1933_v161_
        nw299_[int(2)] = d_1933_v161_
        nw299_[int(3)] = default__.fm55(((d_1934_v162_)[p0] if (p0) in (d_1934_v162_) else (self).f11), (self).f4, (self).f6, globalState)
        nw299_[int(4)] = d_1933_v161_
        nw299_[int(5)] = d_1933_v161_
        nw299_[int(6)] = d_1933_v161_
        nw299_[int(7)] = d_1933_v161_
        nw299_[int(8)] = d_1933_v161_
        nw299_[int(9)] = d_1933_v161_
        nw299_[int(10)] = d_1933_v161_
        nw299_[int(11)] = d_1933_v161_
        nw299_[int(12)] = d_1933_v161_
        nw299_[int(13)] = D1_DC3(default__.fm0(d_1860_v108_, globalState))
        nw299_[int(14)] = d_1933_v161_
        nw299_[int(15)] = d_1933_v161_
        nw299_[int(16)] = D1_DC3((self).f11)
        nw299_[int(17)] = d_1933_v161_
        nw299_[int(18)] = d_1933_v161_
        nw299_[int(19)] = d_1933_v161_
        nw299_[int(20)] = d_1933_v161_
        nw299_[int(21)] = d_1933_v161_
        nw299_[int(22)] = d_1933_v161_
        nw299_[int(23)] = d_1933_v161_
        nw299_[int(24)] = d_1933_v161_
        nw299_[int(25)] = d_1933_v161_
        nw299_[int(26)] = d_1933_v161_
        nw299_[int(27)] = d_1933_v161_
        nw299_[int(28)] = d_1933_v161_
        d_1935_v163_ = nw299_
        d_1936_v164_: _dafny.Map
        d_1936_v164_ = _dafny.Map({p0: d_1935_v163_})
        d_1937_v165_: _dafny.Map
        d_1937_v165_ = _dafny.Map({p0: (self).f6})
        d_1938_v166_: _dafny.Map
        d_1938_v166_ = _dafny.Map({d_1937_v165_: (self).f4})
        r1 = ((d_1936_v164_)[((d_1938_v166_)[_dafny.Map({p0: (self).f6})] if (_dafny.Map({p0: (self).f6})) in (d_1938_v166_) else (self).f4)] if (((d_1938_v166_)[_dafny.Map({p0: (self).f6})] if (_dafny.Map({p0: (self).f6})) in (d_1938_v166_) else (self).f4)) in (d_1936_v164_) else d_1935_v163_)
        return r0, r1

    def m7(self, globalState):
        d_1939_v0_: bool
        d_1939_v0_ = True
        d_1939_v0_ = (self).f11
        d_1940_v1_: int
        d_1940_v1_ = 843
        d_1941_v2_: _dafny.Seq
        d_1941_v2_ = _dafny.SeqWithoutIsStrInference([(self).f6])
        d_1942_v3_: _dafny.Seq
        d_1942_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gnnf"))
        d_1943_v4_: _dafny.Seq
        d_1943_v4_ = _dafny.SeqWithoutIsStrInference([True])
        d_1944_v5_: _dafny.Seq
        d_1944_v5_ = _dafny.SeqWithoutIsStrInference([d_1943_v4_])
        d_1945_v6_: D12
        d_1945_v6_ = D12_DC27(d_1944_v5_)
        d_1946_v7_: _dafny.Seq
        d_1946_v7_ = _dafny.SeqWithoutIsStrInference([d_1941_v2_, (default__.fm39(d_1942_v3_, d_1945_v6_, (self).f11, globalState)) + (_dafny.SeqWithoutIsStrInference([d_1940_v1_, (self).f6]))])
        d_1940_v1_ = len((d_1946_v7_)[default__.safeIndex(((self).f6 if (self).f4 else (self).f6), len(d_1946_v7_))])
        d_1947_v8_: _dafny.Array
        nw300_ = _dafny.Array(False, 14)
        d_1947_v8_ = nw300_
        d_1947_v8_ = d_1947_v8_
        d_1948_v9_: C4
        nw301_ = C4()
        nw301_.ctor__(self.f5, d_1940_v1_, (self).f4)
        d_1948_v9_ = nw301_
        rhs277_ = default__.safeDivisionInt(d_1940_v1_, (self).f6)
        rhs278_ = (len((d_1943_v4_ if d_1939_v0_ else (d_1943_v4_).set(default__.safeIndex((self).f6, len(d_1943_v4_)), (self).f11)))) != (d_1940_v1_)
        d_1940_v1_ = rhs277_
        d_1939_v0_ = rhs278_
        d_1949_v10_: _dafny.Map
        d_1949_v10_ = _dafny.Map({(self).f6: d_1940_v1_})
        d_1950_v11_: C1
        nw302_ = C1()
        nw302_.ctor__(d_1940_v1_, (((d_1949_v10_)[d_1940_v1_] if (d_1940_v1_) in (d_1949_v10_) else d_1940_v1_)) * (len(d_1942_v3_)), self.f5, d_1940_v1_, not(d_1939_v0_))
        d_1950_v11_ = nw302_

    def m8(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        r2: D1 = D1.default()()
        d_1951_v0_: str
        d_1951_v0_ = _dafny.CodePoint('a')
        d_1952_v1_: _dafny.Array
        nw303_ = _dafny.Array(False, 15)
        d_1952_v1_ = nw303_
        d_1953_v2_: _dafny.Map
        d_1953_v2_ = _dafny.Map({d_1951_v0_: d_1952_v1_})
        d_1954_v3_: D19
        d_1954_v3_ = D19_DC39(D19_DC37(d_1953_v2_))
        d_1955_v4_: D19
        d_1955_v4_ = D19_DC37((d_1953_v2_).set(d_1951_v0_, d_1952_v1_))
        d_1954_v3_ = D19_DC39(d_1955_v4_)
        if (self).f11:
            d_1956_v5_: int
            d_1956_v5_ = 594
            d_1956_v5_ = (p1) + (default__.fm3((self).f6, (self).f11, globalState))
            d_1957_v6_: _dafny.Set
            d_1957_v6_ = _dafny.Set({not((self).f4)})
            d_1957_v6_ = d_1957_v6_
            r1 = ((self).f11) == (((self).f11 if (self).f4 else (self).f11))
            d_1956_v5_ = (d_1956_v5_) + ((self).f6)
            r0 = (self).f4
        elif True:
            arr9_ = self.f5
            index291_ = default__.safeIndex(951, (self.f5).length(0))
            arr9_[index291_] = d_1952_v1_
            arr10_ = self.f5
            index292_ = default__.safeIndex(951, (self.f5).length(0))
            arr10_[index292_] = d_1952_v1_
            d_1958_v7_: _dafny.Seq
            d_1958_v7_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f4, True, (self).f4])
            d_1959_v8_: _dafny.Seq
            d_1959_v8_ = _dafny.SeqWithoutIsStrInference([(d_1958_v7_) + (d_1958_v7_), (d_1958_v7_) + (d_1958_v7_), (_dafny.SeqWithoutIsStrInference([(self).f11])) + (d_1958_v7_)])
            d_1959_v8_ = _dafny.SeqWithoutIsStrInference([d_1958_v7_ for d_1960_i0_ in range(default__.abs(108))])
            d_1961_v9_: _dafny.MultiSet
            d_1961_v9_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brp"))])
            d_1962_v10_: _dafny.MultiSet
            d_1962_v10_ = _dafny.MultiSet([255, ((d_1961_v9_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cufybjlks"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cufybjlks"))) in (d_1961_v9_) else p0)])
            d_1963_v11_: D0
            d_1963_v11_ = D0_DC0((self).f6)
            d_1964_v12_: _dafny.Map
            d_1964_v12_ = _dafny.Map({p0: d_1951_v0_})
            d_1965_v13_: _dafny.Map
            d_1965_v13_ = _dafny.Map({(self).f4: (self).f4})
            d_1966_v14_: bool
            d_1967_v15_: int
            out58_: bool
            out59_: int
            out58_, out59_ = default__.m0(_dafny.Map({d_1962_v10_: _dafny.SeqWithoutIsStrInference([d_1951_v0_ for d_1968_i1_ in range(default__.abs(577))])}), (self).f4, ((self).fm6(d_1963_v11_, d_1963_v11_, d_1964_v12_, p1, globalState)) | (d_1962_v10_), not (False) or (((d_1965_v13_)[(self).f11] if ((self).f11) in (d_1965_v13_) else (self).f4)), globalState)
            d_1966_v14_ = out58_
            d_1967_v15_ = out59_
            d_1969_v16_: _dafny.Seq
            d_1969_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwdadof"))
            d_1970_v17_: _dafny.Map
            d_1970_v17_ = _dafny.Map({d_1967_v15_: d_1969_v16_})
            d_1971_v18_: D2
            d_1971_v18_ = D2_DC6()
            d_1970_v17_ = (d_1970_v17_).set(default__.safeDivisionInt(p1, 442), ((default__.fm16(d_1966_v14_, (self).f11, False, d_1971_v18_, globalState)).set(default__.safeIndex(p1, len(default__.fm16(d_1966_v14_, (self).f11, False, d_1971_v18_, globalState))), d_1951_v0_)) + (d_1969_v16_))
            d_1972_v19_: _dafny.Array
            nw304_ = _dafny.Array(None, 2)
            nw304_[int(0)] = (self.f5)[default__.safeIndex(951, (self.f5).length(0))]
            nw304_[int(1)] = d_1952_v1_
            d_1972_v19_ = nw304_
            d_1973_v20_: _dafny.Array
            d_1973_v20_ = d_1972_v19_
            d_1974_v21_: C5
            nw305_ = C5()
            nw305_.ctor__((self).f4, (0) - (p0), (d_1973_v20_), (len(d_1969_v16_)) * (p0), d_1966_v14_)
            d_1974_v21_ = nw305_
        d_1975_v22_: int
        d_1975_v22_ = 861
        d_1975_v22_ = (d_1975_v22_) * ((p1) - (p1))
        if (self).f11:
            d_1976_v23_: _dafny.Seq
            d_1976_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "senka"))
            d_1977_v24_: _dafny.Map
            d_1977_v24_ = _dafny.Map({self.f5: (d_1976_v23_) == (d_1976_v23_)})
            d_1977_v24_ = (d_1977_v24_).set(self.f5, (self).f4)
            d_1978_v25_: _dafny.Set
            d_1978_v25_ = _dafny.Set({(self).f4, not((self).f11), (self).f11})
            d_1978_v25_ = (d_1978_v25_).intersection(d_1978_v25_)
            r0 = (self).f4
            (self).m7(globalState)
            d_1951_v0_ = d_1951_v0_
        elif True:
            d_1979_v26_: D1
            d_1979_v26_ = D1_DC3(not(True))
            r2 = d_1979_v26_
            d_1980_v27_: _dafny.Map
            d_1980_v27_ = _dafny.Map({not((self).f4): (self).f4})
            d_1981_v28_: _dafny.Map
            d_1981_v28_ = _dafny.Map({p1: ((d_1980_v27_)[(self).f4] if ((self).f4) in (d_1980_v27_) else not((self).f11))})
            d_1981_v28_ = (d_1981_v28_).set(d_1975_v22_, (self).f11)
            r0 = (self).f11
            index293_ = default__.safeIndex(250, (d_1952_v1_).length(0))
            (d_1952_v1_)[index293_] = (self).f4
            index294_ = default__.safeIndex(250, (d_1952_v1_).length(0))
            (d_1952_v1_)[index294_] = (self).f11
            d_1982_v29_: _dafny.Seq
            d_1982_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pgj"))
            d_1983_v30_: C2
            nw306_ = C2()
            nw306_.ctor__(d_1982_v29_, (d_1952_v1_)[default__.safeIndex(250, (d_1952_v1_).length(0))])
            d_1983_v30_ = nw306_
        d_1984_v31_: _dafny.Seq
        d_1984_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
        d_1985_v32_: T0
        nw307_ = C2()
        nw307_.ctor__(d_1984_v31_, (self).f11)
        d_1985_v32_ = nw307_
        d_1985_v32_ = d_1985_v32_
        d_1986_v33_: _dafny.Map
        d_1986_v33_ = _dafny.Map({((self).f11 if (self).f11 else (self).f4): (self).f11})
        d_1986_v33_ = d_1986_v33_
        r0 = (self).f4
        r1 = ((self).f6) >= (((self).f6) + (d_1975_v22_))
        d_1987_v34_: D1
        d_1987_v34_ = D1_DC3((self).f11)
        r2 = d_1987_v34_
        return r0, r1, r2

    @property
    def f11(self):
        return self._f11

class C7(T1, T0):
    def  __init__(self):
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self._f6: int = int(0)
        self._f4: bool = False
        self._f10: _dafny.Array = _dafny.Array(None, 0)
        self._f9: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f5(self):
        return self._f5
    @f5.setter
    def f5(self, value):
        self._f5 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f4(self):
        return self._f4
    def ctor__(self, f9, f10, f5, f6, f4):
        (self)._f9 = f9
        (self)._f10 = f10
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f4 = f4

    def fm8(self, p0, p1, p2, p3, globalState):
        if ((self).f4 if (self).f4 else (self).f4):
            return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1988_i0_ in range(default__.abs(799))])) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h')]))
        elif True:
            return ((self).f6) == ((0) - ((self).f6))

    def fm9(self, p0, p1, p2, globalState):
        return (self).f6

    def fm6(self, p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([(self).f6]) if False else _dafny.MultiSet([(self).f6]))) - ((_dafny.MultiSet([(self).f6])) - (_dafny.MultiSet([472])))

    def fm7(self, globalState):
        return ((0) - (default__.safeDivisionInt((self).f6, (self).f6))) - (677)

    def fm13(self, p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([(self).f4])) + (_dafny.SeqWithoutIsStrInference([not((self).f4), (self).f4, (self).f4]))

    def fm14(self, p0, p1, p2, p3, globalState):
        return ((self).f6) == (664)

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_1989_v0_: C1
        nw308_ = C1()
        nw308_.ctor__((self).f6, p0, self.f5, p0, (self).f4)
        d_1989_v0_ = nw308_
        d_1990_v1_: _dafny.Seq
        d_1990_v1_ = _dafny.SeqWithoutIsStrInference([(self).f4, (self).f4, True, (self).f4, not((self).f4)])
        source35_ = D4_DC13(194, not((self).f4), ((self).f6) - (d_1989_v0_.f16), d_1990_v1_)
        if source35_.is_DC13:
            d_1991___mcc_h0_ = source35_.cf14
            d_1992___mcc_h1_ = source35_.cf15
            d_1993___mcc_h2_ = source35_.cf16
            d_1994___mcc_h3_ = source35_.cf17
            d_1995_cf17_ = d_1994___mcc_h3_
            d_1996_cf16_ = d_1993___mcc_h2_
            d_1997_cf15_ = d_1992___mcc_h1_
            d_1998_cf14_ = d_1991___mcc_h0_
            d_1999_v2_: _dafny.Set
            d_1999_v2_ = _dafny.Set({(0) - (len(default__.fm64(p0, (self).f6, d_1997_cf15_, globalState)))})
            d_2000_v3_: D9
            d_2000_v3_ = D9_DC20(d_1997_cf15_, _dafny.CodePoint('n'), d_1989_v0_.f16)
            d_2001_v4_: C6
            nw309_ = C6()
            nw309_.ctor__((len(d_1999_v2_)) < (272), (self).f4, self.f5, (d_2000_v3_).cf28)
            d_2001_v4_ = nw309_
            d_1997_cf15_ = True
            d_1997_cf15_ = not(not((d_2001_v4_).f11))
            (d_1989_v0_).f16 = ((d_1989_v0_).f15) - (-913)
        elif True:
            d_2002___mcc_h4_ = source35_.cf13
            d_2003_cf13_ = d_2002___mcc_h4_
            d_2004_v5_: _dafny.Array
            nw310_ = _dafny.Array(int(0), 2)
            d_2004_v5_ = nw310_
            index295_ = default__.safeIndex(382, (d_2004_v5_).length(0))
            (d_2004_v5_)[index295_] = len(d_1990_v1_)
            index296_ = default__.safeIndex(382, (d_2004_v5_).length(0))
            (d_2004_v5_)[index296_] = (d_1989_v0_).f15
            d_2005_v6_: _dafny.Array
            def lambda79_(d_2006_v0_, d_2007_p0_):
                def lambda80_(d_2008_i0_):
                    return (_dafny.MultiSet([d_2006_v0_.f16])) | (_dafny.MultiSet([d_2007_p0_]))

                return lambda80_

            init51_ = lambda79_(d_1989_v0_, p0)
            nw311_ = _dafny.Array(None, 23)
            for i0_51_ in range(nw311_.length(0)):
                nw311_[i0_51_] = init51_(i0_51_)
            d_2005_v6_ = nw311_
            d_2009_v7_: _dafny.MultiSet
            d_2009_v7_ = _dafny.MultiSet([p0])
            index297_ = default__.safeIndex(126, (d_2005_v6_).length(0))
            (d_2005_v6_)[index297_] = (d_2009_v7_ if (self).f4 else d_2009_v7_)
            d_2010_v8_: bool
            d_2010_v8_ = True
            d_2011_v9_: _dafny.Map
            d_2011_v9_ = _dafny.Map({(d_1989_v0_).f15: True})
            d_2012_v10_: C3
            nw312_ = C3()
            nw312_.ctor__(self.f5, (d_1989_v0_).f15, ((self).f4 if (self).fm14(d_2011_v9_, True, d_1989_v0_.f16, False, globalState) else d_2010_v8_))
            d_2012_v10_ = nw312_
            d_2013_v11_: _dafny.Map
            d_2013_v11_ = _dafny.Map({d_2010_v8_: d_2010_v8_})
            d_2014_v12_: _dafny.MultiSet
            d_2014_v12_ = _dafny.MultiSet([d_2010_v8_, ((d_2013_v11_)[(self).f4] if ((self).f4) in (d_2013_v11_) else d_2010_v8_)])
            d_2015_v13_: _dafny.Seq
            d_2015_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdatbjvc"))
            d_2016_v14_: C2
            nw313_ = C2()
            nw313_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyw")), (self).f4)
            d_2016_v14_ = nw313_
            d_2017_v15_: _dafny.Map
            d_2017_v15_ = _dafny.Map({d_1989_v0_.f16: d_2016_v14_})
            d_2018_v16_: D16
            d_2018_v16_ = D16_DC33(False, (self).f6, d_2014_v12_, len(d_2015_v13_), default__.fm23(len(d_2017_v15_), d_1989_v0_.f16, globalState))
            d_2019_v17_: D9
            d_2019_v17_ = D9_DC20((d_2018_v16_).cf52, d_2003_cf13_, d_1989_v0_.f16)
            pat_let_tv39_ = d_2013_v11_
            d_2020_v18_: _dafny.Array
            nw314_ = _dafny.Array(None, 6)
            nw314_[int(0)] = d_2003_cf13_
            nw314_[int(1)] = d_2003_cf13_
            def iife153_(_pat_let48_0):
                def iife154_(d_2021_dt__update__tmp_h0_):
                    def iife155_(_pat_let49_0):
                        def iife156_(d_2022_dt__update_hcf26_h0_):
                            def iife157_(_pat_let50_0):
                                def iife158_(d_2023_dt__update_hcf28_h0_):
                                    return D9_DC20(d_2022_dt__update_hcf26_h0_, (d_2021_dt__update__tmp_h0_).cf27, d_2023_dt__update_hcf28_h0_)
                                return iife158_(_pat_let50_0)
                            return iife157_(len(pat_let_tv39_))
                        return iife156_(_pat_let49_0)
                    return iife155_((self).f4)
                return iife154_(_pat_let48_0)
            nw314_[int(2)] = (iife153_(d_2019_v17_)).cf27
            nw314_[int(3)] = (d_2003_cf13_ if (self).f4 else d_2003_cf13_)
            nw314_[int(4)] = d_2003_cf13_
            nw314_[int(5)] = d_2003_cf13_
            d_2020_v18_ = nw314_
            index298_ = default__.safeIndex(126, (d_2005_v6_).length(0))
            rhs279_ = d_2009_v7_
            rhs280_ = (((self).f4) or (True)) and (((d_1989_v0_).f15) != (d_1989_v0_.f16))
            rhs281_ = d_2012_v10_
            rhs282_ = d_2020_v18_
            lhs169_ = d_2005_v6_
            lhs170_ = default__.safeIndex(126, (d_2005_v6_).length(0))
            lhs169_[lhs170_] = rhs279_
            d_2010_v8_ = rhs280_
            d_2012_v10_ = rhs281_
            d_2020_v18_ = rhs282_
            (d_1989_v0_).f16 = d_1989_v0_.f16
            index299_ = default__.safeIndex(382, (d_2004_v5_).length(0))
            rhs283_ = (default__.safeDivisionInt(989, (d_1989_v0_).f15)) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfu"))))
            rhs284_ = 196
            rhs285_ = (0) - (d_1989_v0_.f16)
            lhs171_ = d_1989_v0_
            lhs172_ = d_1989_v0_
            lhs173_ = d_2004_v5_
            lhs174_ = default__.safeIndex(382, (d_2004_v5_).length(0))
            lhs171_.f16 = rhs283_
            lhs172_.f16 = rhs284_
            lhs173_[lhs174_] = rhs285_
        d_2024_v19_: _dafny.Map
        d_2024_v19_ = _dafny.Map({p0: d_1989_v0_.f16})
        d_2025_v20_: _dafny.Map
        d_2025_v20_ = _dafny.Map({True: d_2024_v19_})
        d_2026_v21_: _dafny.Seq
        d_2026_v21_ = _dafny.SeqWithoutIsStrInference([d_2024_v19_, ((d_2025_v20_)[default__.fm0(_dafny.CodePoint('a'), globalState)] if (default__.fm0(_dafny.CodePoint('a'), globalState)) in (d_2025_v20_) else d_2024_v19_)])
        d_2027_v22_: _dafny.MultiSet
        d_2027_v22_ = _dafny.MultiSet([len((d_2026_v21_)[default__.safeIndex((0) - (p0), len(d_2026_v21_))]), d_1989_v0_.f16, p0])
        d_2028_v23_: _dafny.Map
        d_2028_v23_ = _dafny.Map({d_2027_v22_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rcxfakrum"))})
        d_2029_v24_: bool
        d_2030_v25_: int
        out60_: bool
        out61_: int
        out60_, out61_ = default__.m0(d_2028_v23_, (self).f4, d_2027_v22_, (self).f4, globalState)
        d_2029_v24_ = out60_
        d_2030_v25_ = out61_
        index300_ = default__.safeIndex(270, ((self).f10).length(0))
        ((self).f10)[index300_] = d_2029_v24_
        d_2031_v26_: D11
        d_2031_v26_ = D11_DC26(not((self).f4))
        d_2032_v27_: _dafny.Set
        d_2032_v27_ = _dafny.Set({d_2031_v26_, d_2031_v26_, d_2031_v26_, d_2031_v26_})
        index301_ = default__.safeIndex(270, ((self).f10).length(0))
        ((self).f10)[index301_] = (d_2030_v25_) != (len((d_2032_v27_).intersection(d_2032_v27_)))
        d_2033_v28_: _dafny.Seq
        d_2033_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvmi"))
        d_2034_v29_: _dafny.Map
        d_2034_v29_ = _dafny.Map({False: ((self).f10)[default__.safeIndex(270, ((self).f10).length(0))]})
        d_2035_v30_: _dafny.Seq
        d_2035_v30_ = _dafny.SeqWithoutIsStrInference([(self).f6, len(d_1990_v1_)])
        d_2036_v31_: _dafny.Seq
        d_2036_v31_ = _dafny.SeqWithoutIsStrInference([d_1989_v0_.f16, len(d_2035_v30_)])
        d_2037_v32_: _dafny.Array
        nw315_ = _dafny.Array(None, 14)
        nw315_[int(0)] = (self).fm7(globalState)
        nw315_[int(1)] = (d_1989_v0_).f15
        nw315_[int(2)] = default__.safeModuloInt(len(d_2033_v28_), (d_1989_v0_).f15)
        nw315_[int(3)] = len(d_2034_v29_)
        nw315_[int(4)] = (self).f6
        nw315_[int(5)] = default__.safeModuloInt(d_2030_v25_, len(d_1990_v1_))
        nw315_[int(6)] = ((self).f6) * (p0)
        nw315_[int(7)] = default__.safeDivisionInt(34, 583)
        nw315_[int(8)] = default__.safeDivisionInt(len(d_2036_v31_), 967)
        nw315_[int(9)] = d_1989_v0_.f16
        nw315_[int(10)] = (d_1989_v0_).f15
        nw315_[int(11)] = d_1989_v0_.f16
        nw315_[int(12)] = d_2030_v25_
        nw315_[int(13)] = (d_1989_v0_).f15
        d_2037_v32_ = nw315_
        d_2038_v33_: _dafny.Map
        d_2038_v33_ = _dafny.Map({30: ((self).f10)[default__.safeIndex(270, ((self).f10).length(0))]})
        d_2039_v34_: _dafny.Array
        nw316_ = _dafny.Array(False, 25)
        d_2039_v34_ = nw316_
        d_2040_v35_: D10
        d_2040_v35_ = D10_DC22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsfhsrurn")), d_2029_v24_, (self).f4, d_2039_v34_, (self).f6)
        d_2041_v36_: _dafny.MultiSet
        d_2041_v36_ = _dafny.MultiSet([(self).f4, d_2029_v24_])
        d_2042_v37_: C4
        nw317_ = C4()
        nw317_.ctor__(self.f5, d_2030_v25_, d_2029_v24_)
        d_2042_v37_ = nw317_
        d_2043_v38_: _dafny.Map
        d_2043_v38_ = _dafny.Map({d_2042_v37_: d_2033_v28_})
        d_2044_v39_: _dafny.Set
        d_2044_v39_ = _dafny.Set({(self).f4, d_2029_v24_})
        d_2045_v40_: D16
        d_2045_v40_ = D16_DC33(((d_2038_v33_)[d_1989_v0_.f16] if (d_1989_v0_.f16) in (d_2038_v33_) else ((self).f10)[default__.safeIndex(270, ((self).f10).length(0))]), (d_2042_v37_).fm7(globalState), _dafny.MultiSet([((self).f10)[default__.safeIndex(270, ((self).f10).length(0))]]), (d_2027_v22_).cardinality, d_2044_v39_)
        d_2046_v41_: D16
        d_2046_v41_ = D16_DC33((self).fm14(d_2038_v33_, ((self).f10)[default__.safeIndex(270, ((self).f10).length(0))], (d_1989_v0_).f15, (d_2040_v35_).cf32, globalState), (0) - ((((d_2024_v19_)[d_1989_v0_.f16] if (d_1989_v0_.f16) in (d_2024_v19_) else (d_2027_v22_).cardinality)) + (431)), ((d_2041_v36_).set(d_2029_v24_, default__.abs(len(d_2043_v38_)))).intersection(d_2041_v36_), (0) - ((d_1989_v0_.f16) * (d_1989_v0_.f16)), _dafny.Set({d_2029_v24_, d_2029_v24_, True, not(d_2029_v24_), not((d_2045_v40_).cf52)}))
        index302_ = default__.safeIndex(270, ((self).f10).length(0))
        rhs286_ = 154
        rhs287_ = 842
        rhs288_ = d_2037_v32_
        rhs289_ = d_2045_v40_
        rhs290_ = ((self).f10)[default__.safeIndex(270, ((self).f10).length(0))]
        lhs175_ = d_1989_v0_
        lhs176_ = (self).f10
        lhs177_ = default__.safeIndex(270, ((self).f10).length(0))
        lhs175_.f16 = rhs286_
        d_2030_v25_ = rhs287_
        d_2037_v32_ = rhs288_
        d_2045_v40_ = rhs289_
        lhs176_[lhs177_] = rhs290_
        index303_ = default__.safeIndex(270, ((self).f10).length(0))
        rhs291_ = p0
        rhs292_ = (False if (d_2030_v25_) == (d_1989_v0_.f16) else False)
        lhs178_ = d_1989_v0_
        lhs179_ = (self).f10
        lhs180_ = default__.safeIndex(270, ((self).f10).length(0))
        lhs178_.f16 = rhs291_
        lhs179_[lhs180_] = rhs292_
        d_2047_v42_: _dafny.Map
        d_2047_v42_ = _dafny.Map({(self).f4: p0})
        r0 = d_2047_v42_
        return r0

    def m2(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_2048_v0_: int
        d_2048_v0_ = -830
        d_2049_v1_: _dafny.Seq
        d_2049_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))
        d_2050_v2_: C3
        nw318_ = C3()
        nw318_.ctor__(self.f5, len(d_2049_v1_), p0)
        d_2050_v2_ = nw318_
        d_2051_v3_: _dafny.Map
        d_2051_v3_ = _dafny.Map({d_2050_v2_: False})
        d_2048_v0_ = (default__.safeDivisionInt(len(_dafny.Map({d_2051_v3_: (self).f6})), d_2048_v0_)) * ((self).f6)
        d_2052_i0_: int
        d_2052_i0_ = 0
        with _dafny.label("9"):
            while p0:
                with _dafny.c_label("9"):
                    if (d_2052_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_2052_i0_ = (d_2052_i0_) + (1)
                    d_2053_v4_: _dafny.Array
                    nw319_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
                    d_2053_v4_ = nw319_
                    index304_ = default__.safeIndex(251, (d_2053_v4_).length(0))
                    (d_2053_v4_)[index304_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2054_i1_ in range(default__.abs(578))])
                    d_2055_v5_: bool
                    d_2055_v5_ = False
                    d_2056_v6_: _dafny.Array
                    def lambda81_(d_2057_i2_):
                        return _dafny.CodePoint('d')

                    init52_ = lambda81_
                    nw320_ = _dafny.Array(None, 13)
                    for i0_52_ in range(nw320_.length(0)):
                        nw320_[i0_52_] = init52_(i0_52_)
                    d_2056_v6_ = nw320_
                    d_2058_v7_: str
                    d_2058_v7_ = _dafny.CodePoint('h')
                    index305_ = default__.safeIndex(294, (d_2056_v6_).length(0))
                    (d_2056_v6_)[index305_] = d_2058_v7_
                    d_2059_v8_: _dafny.Set
                    d_2059_v8_ = _dafny.Set({474, (self).f6, (self).f6, (0) - ((self).fm7(globalState))})
                    d_2060_v9_: _dafny.Seq
                    d_2060_v9_ = _dafny.SeqWithoutIsStrInference([not((self).f4)])
                    d_2061_v10_: _dafny.Seq
                    d_2061_v10_ = _dafny.SeqWithoutIsStrInference([d_2049_v1_])
                    d_2062_v11_: _dafny.Map
                    d_2062_v11_ = _dafny.Map({(self).f10: d_2061_v10_})
                    index306_ = default__.safeIndex(251, (d_2053_v4_).length(0))
                    index307_ = default__.safeIndex(294, (d_2056_v6_).length(0))
                    rhs293_ = d_2049_v1_
                    rhs294_ = (True) in (d_2060_v9_)
                    rhs295_ = len(((d_2062_v11_)[(self).f10] if ((self).f10) in (d_2062_v11_) else (d_2061_v10_) + (d_2061_v10_)))
                    rhs296_ = d_2058_v7_
                    rhs297_ = d_2059_v8_
                    lhs181_ = d_2053_v4_
                    lhs182_ = default__.safeIndex(251, (d_2053_v4_).length(0))
                    lhs183_ = d_2056_v6_
                    lhs184_ = default__.safeIndex(294, (d_2056_v6_).length(0))
                    lhs181_[lhs182_] = rhs293_
                    d_2055_v5_ = rhs294_
                    d_2048_v0_ = rhs295_
                    lhs183_[lhs184_] = rhs296_
                    d_2059_v8_ = rhs297_
                    d_2048_v0_ = default__.safeDivisionInt(404, len(d_2049_v1_))
                    d_2063_v12_: _dafny.Map
                    d_2063_v12_ = _dafny.Map({p0: d_2058_v7_})
                    d_2064_v13_: _dafny.Seq
                    d_2064_v13_ = _dafny.SeqWithoutIsStrInference([(default__.fm65(d_2048_v0_, 625, 209, globalState)).set(True, (d_2056_v6_)[default__.safeIndex(294, (d_2056_v6_).length(0))]), d_2063_v12_])
                    d_2048_v0_ = (0) - (len(((d_2063_v12_).set(d_2055_v5_, (d_2056_v6_)[default__.safeIndex(294, (d_2056_v6_).length(0))])) | ((d_2063_v12_) | ((d_2064_v13_)[default__.safeIndex(len(d_2049_v1_), len(d_2064_v13_))]))))
                    d_2055_v5_ = p0
                    pass
            pass
        d_2065_v14_: T0
        nw321_ = C6()
        nw321_.ctor__((self).f4, p0, self.f5, d_2048_v0_)
        d_2065_v14_ = nw321_
        d_2066_v15_: _dafny.Map
        d_2066_v15_ = _dafny.Map({(self).f6: d_2065_v14_})
        d_2067_v16_: _dafny.Map
        d_2067_v16_ = _dafny.Map({(self).f6: (self).f6})
        d_2068_v17_: D3
        d_2068_v17_ = D3_DC9(False, (len(d_2066_v15_)) + (len(d_2067_v16_)), (d_2065_v14_).f4)
        source36_ = d_2068_v17_
        if source36_.is_DC9:
            d_2069___mcc_h0_ = source36_.cf8
            d_2070___mcc_h1_ = source36_.cf9
            d_2071___mcc_h2_ = source36_.cf10
            d_2072_cf10_ = d_2071___mcc_h2_
            d_2073_cf9_ = d_2070___mcc_h1_
            d_2074_cf8_ = d_2069___mcc_h0_
            d_2075_v18_: _dafny.Set
            d_2075_v18_ = _dafny.Set({(d_2065_v14_).f4, p0, (self).f4})
            d_2076_v19_: _dafny.Array
            def lambda82_(d_2077_cf9_):
                def lambda83_(d_2078_i3_):
                    return default__.safeModuloInt(d_2078_i3_, d_2077_cf9_)

                return lambda83_

            init53_ = lambda82_(d_2073_cf9_)
            nw322_ = _dafny.Array(None, 7)
            for i0_53_ in range(nw322_.length(0)):
                nw322_[i0_53_] = init53_(i0_53_)
            d_2076_v19_ = nw322_
            d_2079_v20_: _dafny.Map
            d_2079_v20_ = _dafny.Map({d_2076_v19_: d_2075_v18_})
            d_2080_v21_: str
            d_2080_v21_ = _dafny.CodePoint('x')
            d_2081_v22_: _dafny.Map
            d_2081_v22_ = _dafny.Map({True: (d_2065_v14_).f4})
            d_2082_v23_: _dafny.Array
            nw323_ = _dafny.Array(None, 25)
            nw323_[int(0)] = d_2075_v18_
            nw323_[int(1)] = _dafny.Set({True})
            nw323_[int(2)] = d_2075_v18_
            nw323_[int(3)] = d_2075_v18_
            nw323_[int(4)] = (_dafny.Set({d_2072_cf10_, p0, (d_2065_v14_).f4})) - (d_2075_v18_)
            nw323_[int(5)] = (d_2075_v18_ if d_2074_cf8_ else d_2075_v18_)
            nw323_[int(6)] = d_2075_v18_
            nw323_[int(7)] = d_2075_v18_
            nw323_[int(8)] = _dafny.Set({d_2074_cf8_})
            nw323_[int(9)] = d_2075_v18_
            nw323_[int(10)] = d_2075_v18_
            nw323_[int(11)] = d_2075_v18_
            nw323_[int(12)] = _dafny.Set({p0, (d_2065_v14_).f4, (d_2065_v14_).f4, d_2072_cf10_, (d_2065_v14_).f4})
            nw323_[int(13)] = d_2075_v18_
            nw323_[int(14)] = d_2075_v18_
            nw323_[int(15)] = d_2075_v18_
            nw323_[int(16)] = d_2075_v18_
            nw323_[int(17)] = (d_2075_v18_) - (d_2075_v18_)
            nw323_[int(18)] = d_2075_v18_
            nw323_[int(19)] = d_2075_v18_
            nw323_[int(20)] = (((d_2079_v20_)[d_2076_v19_] if (d_2076_v19_) in (d_2079_v20_) else d_2075_v18_)) | (default__.fm46(d_2080_v21_, globalState))
            nw323_[int(21)] = d_2075_v18_
            nw323_[int(22)] = (d_2075_v18_).intersection(d_2075_v18_)
            nw323_[int(23)] = d_2075_v18_
            nw323_[int(24)] = (d_2075_v18_) | (default__.fm53((self).f6, _dafny.SeqWithoutIsStrInference([d_2081_v22_, d_2081_v22_]), globalState))
            d_2082_v23_ = nw323_
            d_2083_v24_: _dafny.MultiSet
            d_2083_v24_ = _dafny.MultiSet([(d_2065_v14_).f4])
            index308_ = default__.safeIndex(856, (d_2082_v23_).length(0))
            (d_2082_v23_)[index308_] = default__.fm53(d_2073_cf9_, default__.fm54((d_2083_v24_).cardinality, globalState), globalState)
            index309_ = default__.safeIndex(856, (d_2082_v23_).length(0))
            (d_2082_v23_)[index309_] = d_2075_v18_
            d_2084_v25_: _dafny.Array
            d_2084_v25_ = self.f5
            d_2085_v26_: C5
            nw324_ = C5()
            nw324_.ctor__(d_2072_cf10_, (self).f6, (d_2084_v25_), d_2073_cf9_, (d_2065_v14_).f4)
            d_2085_v26_ = nw324_
            if (not((d_2085_v26_).f13)) or ((d_2048_v0_) < (d_2048_v0_)):
                d_2086_v27_: _dafny.Array
                nw325_ = _dafny.Array(_dafny.Set({}), 25)
                d_2086_v27_ = nw325_
                d_2087_v28_: _dafny.Array
                nw326_ = _dafny.Array(False, 6)
                d_2087_v28_ = nw326_
                index310_ = default__.safeIndex(176, (d_2076_v19_).length(0))
                (d_2076_v19_)[index310_] = default__.fm3(d_2048_v0_, d_2074_cf8_, globalState)
                index311_ = default__.safeIndex(176, (d_2076_v19_).length(0))
                rhs298_ = d_2086_v27_
                rhs299_ = d_2087_v28_
                rhs300_ = (0) - (default__.safeDivisionInt((0) - (d_2048_v0_), 61))
                lhs185_ = d_2076_v19_
                lhs186_ = default__.safeIndex(176, (d_2076_v19_).length(0))
                d_2086_v27_ = rhs298_
                d_2087_v28_ = rhs299_
                lhs185_[lhs186_] = rhs300_
                d_2088_v29_: _dafny.MultiSet
                d_2088_v29_ = _dafny.MultiSet([(self).f10])
                d_2088_v29_ = _dafny.MultiSet([(self).f10])
                d_2072_cf10_ = ((d_2048_v0_) + ((self).f6)) < (((0) - ((0) - ((d_2076_v19_)[default__.safeIndex(176, (d_2076_v19_).length(0))]))) * (d_2085_v26_.f14))
                d_2089_v30_: _dafny.Seq
                d_2089_v30_ = _dafny.SeqWithoutIsStrInference([(0) - (d_2085_v26_.f14)])
                d_2048_v0_ = (d_2089_v30_)[default__.safeIndex(d_2085_v26_.f14, len(d_2089_v30_))]
                d_2090_v31_: C4
                nw327_ = C4()
                nw327_.ctor__(self.f5, d_2085_v26_.f14, d_2074_cf8_)
                d_2090_v31_ = nw327_
            elif True:
                d_2091_v32_: _dafny.Map
                d_2091_v32_ = _dafny.Map({p0: d_2048_v0_})
                d_2067_v16_ = (d_2067_v16_).set(((self).f6) + (len(d_2091_v32_)), d_2073_cf9_)
                d_2048_v0_ = default__.safeModuloInt((0) - (d_2085_v26_.f14), d_2085_v26_.f14)
                d_2048_v0_ = len(d_2049_v1_)
                d_2092_v33_: D2
                d_2092_v33_ = D2_DC5((self).f10)
                d_2092_v33_ = d_2092_v33_
                d_2093_v34_: _dafny.Map
                d_2093_v34_ = _dafny.Map({d_2085_v26_.f14: d_2072_cf10_})
                d_2093_v34_ = (d_2093_v34_).set(((self).f6) - ((self).f6), not((d_2065_v14_).f4))
            d_2072_cf10_ = not((d_2065_v14_).f4)
        elif source36_.is_DC10:
            d_2094___mcc_h3_ = source36_.cf11
            d_2095_cf11_ = d_2094___mcc_h3_
            d_2096_v35_: _dafny.Map
            d_2096_v35_ = _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([(self).f4]))): (self).f10})
            d_2096_v35_ = (d_2096_v35_).set((self).f6, (self).f10)
            d_2097_v36_: D12
            d_2097_v36_ = D12_DC28(d_2048_v0_, d_2095_cf11_, False, (self).f4, d_2095_cf11_)
            if (d_2097_v36_).cf46:
                d_2098_v37_: _dafny.Map
                d_2098_v37_ = _dafny.Map({(d_2065_v14_).f4: d_2049_v1_})
                d_2099_v39_: _dafny.Array
                nw328_ = _dafny.Array(None, 7)
                nw328_[int(0)] = 154
                nw328_[int(1)] = default__.safeModuloInt(len(d_2049_v1_), len((_dafny.Map({(self).f4: d_2048_v0_})).set(p0, d_2048_v0_)))
                nw328_[int(2)] = (self).fm7(globalState)
                def iife159_():
                    coll57_ = _dafny.Set()
                    compr_57_: int
                    for compr_57_ in _dafny.IntegerRange(587, 174):
                        d_2100_v38_: int = compr_57_
                        if ((587) <= (d_2100_v38_)) and ((d_2100_v38_) < (174)):
                            coll57_ = coll57_.union(_dafny.Set([(d_2100_v38_) - ((self).f6)]))
                    return _dafny.Set(coll57_)
                nw328_[int(3)] = default__.safeModuloInt(len(iife159_()
                ), d_2095_cf11_)
                nw328_[int(4)] = (self).f6
                nw328_[int(5)] = (_dafny.MultiSet([(d_2065_v14_).f4, (self).f4, p0])).cardinality
                nw328_[int(6)] = (self).f6
                d_2099_v39_ = nw328_
                d_2101_v40_: _dafny.Seq
                d_2101_v40_ = _dafny.SeqWithoutIsStrInference([p0])
                index312_ = default__.safeIndex(891, (d_2099_v39_).length(0))
                (d_2099_v39_)[index312_] = len((d_2101_v40_) + (d_2101_v40_))
                d_2102_v41_: _dafny.Array
                nw329_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_2102_v41_ = nw329_
                index313_ = default__.safeIndex(166, (d_2102_v41_).length(0))
                (d_2102_v41_)[index313_] = d_2099_v39_
                d_2103_v42_: bool
                d_2103_v42_ = False
                d_2104_v43_: str
                d_2104_v43_ = _dafny.CodePoint('h')
                d_2105_v44_: _dafny.Map
                d_2105_v44_ = _dafny.Map({(d_2065_v14_).f4: d_2048_v0_})
                index314_ = default__.safeIndex(891, (d_2099_v39_).length(0))
                index315_ = default__.safeIndex(166, (d_2102_v41_).length(0))
                rhs301_ = (_dafny.Map({(self).f4: d_2049_v1_})) | (((self).f9) | (d_2098_v37_))
                rhs302_ = d_2048_v0_
                rhs303_ = d_2099_v39_
                rhs304_ = default__.fm0(d_2104_v43_, globalState)
                rhs305_ = ((0) - ((0) - (len((d_2105_v44_) | (d_2105_v44_))))) == (-513)
                lhs187_ = d_2099_v39_
                lhs188_ = default__.safeIndex(891, (d_2099_v39_).length(0))
                lhs189_ = d_2102_v41_
                lhs190_ = default__.safeIndex(166, (d_2102_v41_).length(0))
                d_2098_v37_ = rhs301_
                lhs187_[lhs188_] = rhs302_
                lhs189_[lhs190_] = rhs303_
                d_2103_v42_ = rhs304_
                d_2103_v42_ = rhs305_
                d_2106_v45_: _dafny.Set
                d_2106_v45_ = _dafny.Set({not((d_2065_v14_).f4), not((d_2065_v14_).f4)})
                d_2107_v46_: D16
                d_2107_v46_ = D16_DC33(False, d_2095_cf11_, _dafny.MultiSet([False]), (self).f6, d_2106_v45_)
                d_2103_v42_ = (_dafny.MultiSet([(d_2065_v14_).f4])).isdisjoint((d_2107_v46_).cf54)
                d_2108_v47_: _dafny.Array
                nw330_ = _dafny.Array(_dafny.Map({}), 23)
                d_2108_v47_ = nw330_
                d_2109_v48_: _dafny.MultiSet
                d_2109_v48_ = _dafny.MultiSet([(d_2099_v39_)[default__.safeIndex(891, (d_2099_v39_).length(0))], len(d_2105_v44_), d_2095_cf11_])
                d_2110_v49_: _dafny.Array
                nw331_ = _dafny.Array(None, 12)
                nw331_[int(0)] = d_2104_v43_
                nw331_[int(1)] = default__.fm1((d_2099_v39_)[default__.safeIndex(891, (d_2099_v39_).length(0))], (d_2065_v14_).f4, globalState)
                nw331_[int(2)] = _dafny.CodePoint('c')
                nw331_[int(3)] = d_2104_v43_
                nw331_[int(4)] = d_2104_v43_
                nw331_[int(5)] = d_2104_v43_
                nw331_[int(6)] = d_2104_v43_
                nw331_[int(7)] = _dafny.CodePoint('y')
                nw331_[int(8)] = d_2104_v43_
                nw331_[int(9)] = d_2104_v43_
                nw331_[int(10)] = _dafny.CodePoint('l')
                nw331_[int(11)] = d_2104_v43_
                d_2110_v49_ = nw331_
                d_2111_v50_: _dafny.Map
                d_2111_v50_ = _dafny.Map({((d_2109_v48_)[d_2095_cf11_] if (d_2095_cf11_) in (d_2109_v48_) else 879): d_2110_v49_})
                index316_ = default__.safeIndex(836, (d_2108_v47_).length(0))
                (d_2108_v47_)[index316_] = d_2111_v50_
                d_2112_v51_: D11
                d_2112_v51_ = D11_DC25(True, d_2095_cf11_)
                d_2113_v52_: _dafny.Set
                d_2113_v52_ = _dafny.Set({d_2048_v0_})
                d_2114_v53_: _dafny.Seq
                d_2114_v53_ = _dafny.SeqWithoutIsStrInference([(d_2112_v51_).cf40, len(_dafny.SeqWithoutIsStrInference([d_2113_v52_]))])
                d_2115_v54_: _dafny.Map
                d_2115_v54_ = _dafny.Map({len(d_2114_v53_): (d_2065_v14_).f4})
                d_2116_v55_: _dafny.Set
                d_2116_v55_ = _dafny.Set({d_2115_v54_})
                d_2117_v56_: _dafny.MultiSet
                d_2117_v56_ = _dafny.MultiSet([(d_2065_v14_).f4, (d_2116_v55_) != (d_2116_v55_), d_2103_v42_])
                index317_ = default__.safeIndex(836, (d_2108_v47_).length(0))
                rhs306_ = (p0) or ((d_2065_v14_).f4)
                rhs307_ = d_2103_v42_
                rhs308_ = (d_2117_v56_).cardinality
                rhs309_ = ((d_2111_v50_) | (d_2111_v50_) if p0 else (d_2111_v50_) | (d_2111_v50_))
                lhs191_ = d_2108_v47_
                lhs192_ = default__.safeIndex(836, (d_2108_v47_).length(0))
                d_2103_v42_ = rhs306_
                d_2103_v42_ = rhs307_
                d_2048_v0_ = rhs308_
                lhs191_[lhs192_] = rhs309_
                d_2118_v57_: C1
                nw332_ = C1()
                nw332_.ctor__((d_2099_v39_)[default__.safeIndex(891, (d_2099_v39_).length(0))], (d_2099_v39_)[default__.safeIndex(891, (d_2099_v39_).length(0))], self.f5, ((d_2105_v44_)[p0] if (p0) in (d_2105_v44_) else (self).f6), (d_2065_v14_).f4)
                d_2118_v57_ = nw332_
                d_2119_v58_: _dafny.Seq
                d_2119_v58_ = _dafny.SeqWithoutIsStrInference([(d_2118_v57_ if (self).f4 else d_2118_v57_), d_2118_v57_, d_2118_v57_, d_2118_v57_, d_2118_v57_])
                d_2120_v59_: _dafny.Seq
                d_2120_v59_ = _dafny.SeqWithoutIsStrInference([(d_2119_v58_) + ((d_2119_v58_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))), len(d_2119_v58_)), d_2118_v57_))])
                d_2119_v58_ = (d_2120_v59_)[default__.safeIndex(628, len(d_2120_v59_))]
                d_2121_v60_: D9
                d_2121_v60_ = D9_DC20((d_2065_v14_).f4, d_2104_v43_, default__.safeDivisionInt((self).f6, (self).f6))
                d_2122_v61_: _dafny.Map
                d_2122_v61_ = _dafny.Map({d_2104_v43_: (self).f10})
                d_2123_v62_: D19
                d_2123_v62_ = D19_DC37(d_2122_v61_)
                index318_ = default__.safeIndex(891, (d_2099_v39_).length(0))
                def iife160_():
                    coll58_ = _dafny.Set()
                    compr_58_: int
                    for compr_58_ in _dafny.IntegerRange(375, 806):
                        d_2124_v63_: int = compr_58_
                        if ((375) <= (d_2124_v63_)) and ((d_2124_v63_) < (806)):
                            coll58_ = coll58_.union(_dafny.Set([(d_2124_v63_) * ((d_2118_v57_).f15)]))
                    return _dafny.Set(coll58_)
                rhs310_ = (len(iife160_()
                )) * (d_2048_v0_)
                rhs311_ = d_2121_v60_
                rhs312_ = (d_2065_v14_).f4
                rhs313_ = d_2123_v62_
                lhs193_ = d_2099_v39_
                lhs194_ = default__.safeIndex(891, (d_2099_v39_).length(0))
                lhs193_[lhs194_] = rhs310_
                d_2121_v60_ = rhs311_
                d_2103_v42_ = rhs312_
                d_2123_v62_ = rhs313_
            elif True:
                d_2095_cf11_ = d_2048_v0_
                d_2125_v64_: bool
                d_2125_v64_ = False
                index319_ = default__.safeIndex(635, ((self).f10).length(0))
                ((self).f10)[index319_] = not((d_2065_v14_).f4)
                d_2126_v65_: _dafny.Array
                nw333_ = _dafny.Array(int(0), 2)
                d_2126_v65_ = nw333_
                d_2127_v66_: _dafny.Map
                d_2127_v66_ = _dafny.Map({d_2048_v0_: d_2126_v65_})
                d_2128_v67_: _dafny.Set
                d_2128_v67_ = _dafny.Set({False})
                d_2129_v68_: _dafny.Seq
                d_2129_v68_ = _dafny.SeqWithoutIsStrInference([d_2128_v67_, d_2128_v67_, d_2128_v67_, d_2128_v67_])
                d_2130_v69_: _dafny.Map
                d_2130_v69_ = d_2067_v16_
                index320_ = default__.safeIndex(635, ((self).f10).length(0))
                rhs314_ = (d_2048_v0_) >= (len(d_2127_v66_))
                rhs315_ = not((_dafny.MultiSet([d_2095_cf11_])) != ((_dafny.MultiSet([(self).f6, 762])).set((self).f6, default__.abs(len(d_2129_v68_)))))
                rhs316_ = True
                rhs317_ = not(not ((d_2067_v16_) != ((d_2130_v69_))) or ((d_2065_v14_).f4))
                lhs195_ = (self).f10
                lhs196_ = default__.safeIndex(635, ((self).f10).length(0))
                d_2125_v64_ = rhs314_
                d_2125_v64_ = rhs315_
                lhs195_[lhs196_] = rhs316_
                d_2125_v64_ = rhs317_
                d_2048_v0_ = d_2048_v0_
                index321_ = default__.safeIndex(635, ((self).f10).length(0))
                ((self).f10)[index321_] = not (p0) or (((d_2065_v14_).f4) or (p0))
                d_2131_v70_: _dafny.Map
                d_2131_v70_ = _dafny.Map({p0: ((self).f10)[default__.safeIndex(635, ((self).f10).length(0))]})
                d_2132_v71_: _dafny.Seq
                d_2132_v71_ = _dafny.SeqWithoutIsStrInference([d_2095_cf11_])
                d_2133_v72_: _dafny.Map
                d_2133_v72_ = _dafny.Map({D6_DC16(True, ((d_2131_v70_)[(d_2065_v14_).f4] if ((d_2065_v14_).f4) in (d_2131_v70_) else (d_2065_v14_).f4), d_2132_v71_): not(d_2125_v64_)})
                d_2133_v72_ = d_2133_v72_
            d_2134_v73_: _dafny.Seq
            d_2134_v73_ = _dafny.SeqWithoutIsStrInference([True])
            d_2135_v74_: _dafny.Set
            d_2135_v74_ = _dafny.Set({d_2048_v0_, len(d_2134_v73_), d_2048_v0_})
            if ((d_2135_v74_).ispropersubset(d_2135_v74_) if not((d_2065_v14_).f4) else p0):
                d_2136_v75_: bool
                d_2136_v75_ = False
                d_2137_v76_: _dafny.MultiSet
                d_2137_v76_ = _dafny.MultiSet([len((D10_DC22(default__.fm50(d_2048_v0_, d_2048_v0_, d_2095_cf11_, (d_2065_v14_).f4, globalState), p0, (self).f4, (self).f10, (self).f6)).cf30), -163, d_2095_cf11_])
                d_2136_v75_ = (((d_2137_v76_).set(d_2095_cf11_, default__.abs((self).f6))).cardinality) == (d_2048_v0_)
                d_2095_cf11_ = (self).f6
                d_2095_cf11_ = (((self).f6) + ((self).f6)) - (len(d_2049_v1_))
                d_2095_cf11_ = ((d_2050_v2_).fm7(globalState) if (default__.fm66(d_2048_v0_, globalState)).cf52 else len(d_2134_v73_))
                index322_ = default__.safeIndex(31, ((self).f10).length(0))
                ((self).f10)[index322_] = p0
                index323_ = default__.safeIndex(31, ((self).f10).length(0))
                ((self).f10)[index323_] = (d_2048_v0_) == (d_2048_v0_)
            elif True:
                d_2138_v77_: bool
                d_2138_v77_ = True
                d_2139_v78_: D11
                d_2139_v78_ = D11_DC25(d_2138_v77_, (self).f6)
                d_2138_v77_ = (d_2139_v78_).cf39
                d_2140_v79_: _dafny.Map
                d_2140_v79_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2095_cf11_, (self).f6]): (d_2065_v14_).f4})
                d_2141_v80_: _dafny.Seq
                d_2141_v80_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f6])).cardinality, d_2095_cf11_])
                d_2142_v81_: C1
                nw334_ = C1()
                nw334_.ctor__(len((d_2140_v79_).set(d_2141_v80_, (d_2065_v14_).f4)), len(_dafny.SeqWithoutIsStrInference([d_2095_cf11_ for d_2143_i4_ in range(default__.abs(400))])), self.f5, d_2095_cf11_, (d_2065_v14_).f4)
                d_2142_v81_ = nw334_
                d_2144_v82_: _dafny.Map
                d_2144_v82_ = _dafny.Map({d_2138_v77_: (0) - ((self).f6)})
                (d_2142_v81_).f16 = (d_2142_v81_.f16) + (default__.safeModuloInt(len(d_2144_v82_), (d_2065_v14_).fm7(globalState)))
                d_2145_v83_: _dafny.MultiSet
                d_2145_v83_ = _dafny.MultiSet([d_2142_v81_.f16, len(_dafny.Map({d_2142_v81_.f16: d_2138_v77_})), (d_2142_v81_).f15])
                d_2146_v84_: _dafny.Seq
                d_2146_v84_ = _dafny.SeqWithoutIsStrInference([d_2134_v73_])
                d_2147_v85_: _dafny.MultiSet
                d_2147_v85_ = _dafny.MultiSet([((0) - ((0) - (d_2142_v81_.f16))) * (d_2095_cf11_), default__.safeDivisionInt((self).f6, d_2048_v0_), ((d_2145_v83_)[220] if (220) in (d_2145_v83_) else ((d_2144_v82_)[(self).f4] if ((self).f4) in (d_2144_v82_) else (D4_DC13(len(d_2067_v16_), d_2138_v77_, d_2048_v0_, (d_2146_v84_)[default__.safeIndex(d_2142_v81_.f16, len(d_2146_v84_))])).cf16)), (self).f6])
                d_2148_v86_: _dafny.Map
                d_2148_v86_ = _dafny.Map({d_2142_v81_.f16: (self).f4})
                d_2147_v85_ = ((_dafny.MultiSet(d_2141_v80_)).set(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfiwvdm"))), default__.abs(len(d_2148_v86_)))).set(d_2048_v0_, default__.abs((0) - ((0) - ((self).f6))))
                d_2149_v87_: _dafny.Array
                nw335_ = _dafny.Array(_dafny.Seq({}), 24)
                d_2149_v87_ = nw335_
                d_2150_v88_: D10
                d_2150_v88_ = D10_DC21(d_2144_v82_)
                d_2151_v89_: _dafny.Map
                d_2151_v89_ = _dafny.Map({d_2150_v88_: d_2048_v0_})
                d_2152_v90_: _dafny.Seq
                d_2152_v90_ = _dafny.SeqWithoutIsStrInference([d_2151_v89_])
                index324_ = default__.safeIndex(101, (d_2149_v87_).length(0))
                (d_2149_v87_)[index324_] = d_2152_v90_
                index325_ = default__.safeIndex(172, ((self).f10).length(0))
                ((self).f10)[index325_] = (d_2065_v14_).f4
                d_2153_v91_: _dafny.Set
                d_2153_v91_ = _dafny.Set({d_2049_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugmeom")), d_2049_v1_})
                d_2154_v92_: _dafny.Map
                d_2154_v92_ = _dafny.Map({(d_2049_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))): ((_dafny.MultiSet([(self).f6, d_2142_v81_.f16, d_2095_cf11_, d_2095_cf11_, (0) - (len(d_2153_v91_))])).cardinality) * (len(d_2141_v80_))})
                d_2155_v93_: str
                d_2155_v93_ = _dafny.CodePoint('w')
                index326_ = default__.safeIndex(101, (d_2149_v87_).length(0))
                index327_ = default__.safeIndex(172, ((self).f10).length(0))
                rhs318_ = p0
                rhs319_ = d_2152_v90_
                rhs320_ = (d_2065_v14_).f4
                rhs321_ = d_2155_v93_
                rhs322_ = (d_2154_v92_).set(d_2049_v1_, d_2142_v81_.f16)
                lhs197_ = d_2149_v87_
                lhs198_ = default__.safeIndex(101, (d_2149_v87_).length(0))
                lhs199_ = (self).f10
                lhs200_ = default__.safeIndex(172, ((self).f10).length(0))
                d_2138_v77_ = rhs318_
                lhs197_[lhs198_] = rhs319_
                lhs199_[lhs200_] = rhs320_
                r0 = rhs321_
                d_2154_v92_ = rhs322_
            arr11_ = self.f5
            index328_ = default__.safeIndex(896, (self.f5).length(0))
            arr11_[index328_] = (self).f10
            d_2156_v94_: _dafny.Seq
            d_2156_v94_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10, ((self).f10 if (d_2065_v14_).f4 else (self).f10), (self).f10, ((d_2096_v35_)[d_2048_v0_] if (d_2048_v0_) in (d_2096_v35_) else (self).f10)])
            arr12_ = self.f5
            index329_ = default__.safeIndex(896, (self.f5).length(0))
            arr12_[index329_] = (d_2156_v94_)[default__.safeIndex(default__.safeModuloInt(d_2048_v0_, 122), len(d_2156_v94_))]
        elif source36_.is_DC8:
            d_2157___mcc_h4_ = source36_.cf7
            d_2158_cf7_ = d_2157___mcc_h4_
            d_2159_v95_: _dafny.Map
            d_2159_v95_ = default__.fm29(globalState)
            d_2160_v96_: _dafny.Seq
            d_2160_v96_ = _dafny.SeqWithoutIsStrInference([d_2158_cf7_])
            d_2161_v97_: D20
            d_2161_v97_ = D20_DC40(_dafny.Map({(self).f4: D12_DC27(d_2160_v96_)}))
            d_2162_v98_: _dafny.Map
            d_2162_v98_ = _dafny.Map({len(d_2158_cf7_): d_2161_v97_})
            index330_ = default__.safeIndex(211, ((self).f10).length(0))
            ((self).f10)[index330_] = (self).f4
            d_2163_v99_: _dafny.MultiSet
            d_2163_v99_ = _dafny.MultiSet([d_2158_cf7_, _dafny.SeqWithoutIsStrInference([p0, (self).f4, p0])])
            d_2164_v100_: D6
            d_2164_v100_ = D6_DC15((d_2163_v99_).set((d_2158_cf7_).set(default__.safeIndex(-524, len(d_2158_cf7_)), (d_2065_v14_).f4), default__.abs(d_2048_v0_)))
            index331_ = default__.safeIndex(211, ((self).f10).length(0))
            rhs323_ = (self).f6
            rhs324_ = d_2159_v95_
            rhs325_ = _dafny.Map({(self).f6: d_2161_v97_})
            rhs326_ = p0
            rhs327_ = d_2164_v100_
            lhs201_ = (self).f10
            lhs202_ = default__.safeIndex(211, ((self).f10).length(0))
            d_2048_v0_ = rhs323_
            d_2159_v95_ = rhs324_
            d_2162_v98_ = rhs325_
            lhs201_[lhs202_] = rhs326_
            d_2164_v100_ = rhs327_
            d_2165_v101_: _dafny.Array
            def lambda84_(d_2166_v1_):
                def lambda85_(d_2167_i5_):
                    return ((d_2166_v1_) + (d_2166_v1_)).set(default__.safeIndex((self).f6, len((d_2166_v1_) + (d_2166_v1_))), _dafny.CodePoint('s'))

                return lambda85_

            init54_ = lambda84_(d_2049_v1_)
            nw336_ = _dafny.Array(None, 4)
            for i0_54_ in range(nw336_.length(0)):
                nw336_[i0_54_] = init54_(i0_54_)
            d_2165_v101_ = nw336_
            index332_ = default__.safeIndex(642, (d_2165_v101_).length(0))
            (d_2165_v101_)[index332_] = d_2049_v1_
            d_2168_v102_: _dafny.Seq
            d_2168_v102_ = _dafny.SeqWithoutIsStrInference([(self).f6])
            d_2169_v103_: str
            d_2169_v103_ = _dafny.CodePoint('o')
            index333_ = default__.safeIndex(642, (d_2165_v101_).length(0))
            (d_2165_v101_)[index333_] = (default__.fm37((self).f4, ((d_2065_v14_).f4 if (d_2065_v14_).f4 else (d_2050_v2_).fm8(d_2048_v0_, True, d_2049_v1_, d_2168_v102_, globalState)), (self).f6, (d_2065_v14_).f4, globalState)).set(default__.safeIndex(d_2048_v0_, len(default__.fm37((self).f4, ((d_2065_v14_).f4 if (d_2065_v14_).f4 else (d_2050_v2_).fm8(d_2048_v0_, True, d_2049_v1_, d_2168_v102_, globalState)), (self).f6, (d_2065_v14_).f4, globalState))), d_2169_v103_)
            if p0:
                index334_ = default__.safeIndex(211, ((self).f10).length(0))
                ((self).f10)[index334_] = (not ((d_2158_cf7_)[default__.safeIndex(4, len(d_2158_cf7_))]) or ((d_2065_v14_).f4)) and ((self).f4)
                d_2170_v104_: _dafny.Array
                def lambda86_(d_2171_i6_):
                    return (D1_DC3(((self).f10)[default__.safeIndex(211, ((self).f10).length(0))])).cf2

                init55_ = lambda86_
                nw337_ = _dafny.Array(None, 6)
                for i0_55_ in range(nw337_.length(0)):
                    nw337_[i0_55_] = init55_(i0_55_)
                d_2170_v104_ = nw337_
                d_2170_v104_ = (d_2170_v104_ if (d_2168_v102_) <= (d_2168_v102_) else d_2170_v104_)
                d_2048_v0_ = default__.safeDivisionInt(default__.safeDivisionInt((self).f6, len(d_2158_cf7_)), 846)
                index335_ = default__.safeIndex(211, ((self).f10).length(0))
                ((self).f10)[index335_] = not(((self).f10)[default__.safeIndex(211, ((self).f10).length(0))])
                d_2048_v0_ = d_2048_v0_
            elif True:
                index336_ = default__.safeIndex(211, ((self).f10).length(0))
                index337_ = default__.safeIndex(642, (d_2165_v101_).length(0))
                index338_ = default__.safeIndex(211, ((self).f10).length(0))
                rhs328_ = not(not ((d_2065_v14_).f4) or (True))
                rhs329_ = ((d_2165_v101_)[default__.safeIndex(642, (d_2165_v101_).length(0))]) + (d_2049_v1_)
                rhs330_ = ((d_2048_v0_) <= (d_2048_v0_) if (d_2065_v14_).f4 else ((self).f6) >= (d_2048_v0_))
                rhs331_ = d_2048_v0_
                lhs203_ = (self).f10
                lhs204_ = default__.safeIndex(211, ((self).f10).length(0))
                lhs205_ = d_2165_v101_
                lhs206_ = default__.safeIndex(642, (d_2165_v101_).length(0))
                lhs207_ = (self).f10
                lhs208_ = default__.safeIndex(211, ((self).f10).length(0))
                lhs203_[lhs204_] = rhs328_
                lhs205_[lhs206_] = rhs329_
                lhs207_[lhs208_] = rhs330_
                d_2048_v0_ = rhs331_
                index339_ = default__.safeIndex(211, ((self).f10).length(0))
                ((self).f10)[index339_] = ((D3_DC9(default__.fm0(d_2169_v103_, globalState), (self).f6, ((self).f10)[default__.safeIndex(211, ((self).f10).length(0))])).cf8) in (_dafny.Map({True: (d_2065_v14_).f4}))
                d_2172_v105_: D12
                d_2172_v105_ = D12_DC27((d_2160_v96_).set(default__.safeIndex((self).f6, len(d_2160_v96_)), d_2158_cf7_))
                index340_ = default__.safeIndex(211, ((self).f10).length(0))
                index341_ = default__.safeIndex(211, ((self).f10).length(0))
                rhs332_ = (default__.fm39((d_2165_v101_)[default__.safeIndex(642, (d_2165_v101_).length(0))], d_2172_v105_, (d_2065_v14_).f4, globalState)) + ((_dafny.SeqWithoutIsStrInference([d_2048_v0_])) + (d_2168_v102_))
                rhs333_ = (d_2168_v102_)[default__.safeIndex((self).f6, len(d_2168_v102_))]
                rhs334_ = (d_2065_v14_).f4
                rhs335_ = not(False)
                lhs209_ = (self).f10
                lhs210_ = default__.safeIndex(211, ((self).f10).length(0))
                lhs211_ = (self).f10
                lhs212_ = default__.safeIndex(211, ((self).f10).length(0))
                d_2168_v102_ = rhs332_
                d_2048_v0_ = rhs333_
                lhs209_[lhs210_] = rhs334_
                lhs211_[lhs212_] = rhs335_
                d_2173_v106_: T1
                nw338_ = C6()
                nw338_.ctor__(((self).f10)[default__.safeIndex(211, ((self).f10).length(0))], (self).f4, self.f5, (self).f6)
                d_2173_v106_ = nw338_
                d_2174_v107_: _dafny.Map
                d_2174_v107_ = _dafny.Map({d_2173_v106_: d_2049_v1_})
                d_2175_v108_: _dafny.Map
                d_2175_v108_ = _dafny.Map({(d_2065_v14_).f4: (self).f6})
                d_2176_v109_: _dafny.MultiSet
                d_2176_v109_ = _dafny.MultiSet([(d_2173_v106_).f6, ((d_2175_v108_)[not((d_2050_v2_).fm8(len(d_2158_cf7_), (d_2173_v106_).f4, d_2049_v1_, d_2168_v102_, globalState))] if (not((d_2050_v2_).fm8(len(d_2158_cf7_), (d_2173_v106_).f4, d_2049_v1_, d_2168_v102_, globalState))) in (d_2175_v108_) else (self).f6)])
                d_2177_v110_: _dafny.Map
                d_2177_v110_ = _dafny.Map({(d_2173_v106_).f6: d_2049_v1_})
                d_2178_v111_: _dafny.Map
                d_2178_v111_ = _dafny.Map({((self).f10)[default__.safeIndex(211, ((self).f10).length(0))]: (d_2065_v14_).f4})
                d_2179_v112_: _dafny.MultiSet
                d_2179_v112_ = _dafny.MultiSet([((self).f10)[default__.safeIndex(211, ((self).f10).length(0))]])
                d_2180_v113_: _dafny.Array
                nw339_ = _dafny.Array(None, 26)
                nw339_[int(0)] = (d_2050_v2_).fm7(globalState)
                nw339_[int(1)] = d_2048_v0_
                nw339_[int(2)] = (self).fm9(True, False, (self).f6, globalState)
                nw339_[int(3)] = len((d_2067_v16_) | (d_2067_v16_))
                nw339_[int(4)] = d_2048_v0_
                nw339_[int(5)] = d_2048_v0_
                nw339_[int(6)] = (len(d_2174_v107_)) * (len(d_2158_cf7_))
                nw339_[int(7)] = d_2048_v0_
                nw339_[int(8)] = ((self).f6) + (d_2048_v0_)
                nw339_[int(9)] = (d_2173_v106_).f6
                nw339_[int(10)] = len((d_2168_v102_) + (d_2168_v102_))
                nw339_[int(11)] = ((d_2176_v109_) - (d_2176_v109_)).cardinality
                nw339_[int(12)] = (d_2173_v106_).f6
                nw339_[int(13)] = (self).f6
                nw339_[int(14)] = (len(d_2177_v110_) if (d_2173_v106_).f4 else (self).f6)
                nw339_[int(15)] = d_2048_v0_
                nw339_[int(16)] = len(_dafny.SeqWithoutIsStrInference([False]))
                nw339_[int(17)] = (self).f6
                nw339_[int(18)] = (d_2173_v106_).f6
                nw339_[int(19)] = (self).f6
                nw339_[int(20)] = (len(_dafny.Set({((self).f10)[default__.safeIndex(211, ((self).f10).length(0))]}))) * ((self).f6)
                nw339_[int(21)] = d_2048_v0_
                nw339_[int(22)] = len(d_2049_v1_)
                nw339_[int(23)] = len(((d_2165_v101_)[default__.safeIndex(642, (d_2165_v101_).length(0))]) + ((d_2165_v101_)[default__.safeIndex(642, (d_2165_v101_).length(0))]))
                nw339_[int(24)] = len(d_2178_v111_)
                nw339_[int(25)] = (d_2179_v112_).cardinality
                d_2180_v113_ = nw339_
                d_2181_v114_: _dafny.Map
                d_2181_v114_ = _dafny.Map({d_2067_v16_: (d_2065_v14_).f4})
                index342_ = default__.safeIndex(983, (d_2180_v113_).length(0))
                (d_2180_v113_)[index342_] = len((_dafny.SeqWithoutIsStrInference([len(d_2181_v114_)]) if (d_2065_v14_).f4 else d_2168_v102_))
                index343_ = default__.safeIndex(983, (d_2180_v113_).length(0))
                (d_2180_v113_)[index343_] = default__.safeDivisionInt(d_2048_v0_, (self).f6)
                d_2182_v115_: _dafny.Array
                def lambda87_(d_2183_v106_):
                    def lambda88_(d_2184_i7_):
                        return (d_2183_v106_).f4

                    return lambda88_

                init56_ = lambda87_(d_2173_v106_)
                nw340_ = _dafny.Array(None, 1)
                for i0_56_ in range(nw340_.length(0)):
                    nw340_[i0_56_] = init56_(i0_56_)
                d_2182_v115_ = nw340_
                d_2185_v116_: _dafny.Map
                d_2185_v116_ = _dafny.Map({d_2169_v103_: d_2182_v115_})
                d_2186_v117_: D19
                d_2186_v117_ = D19_DC37(d_2185_v116_)
                d_2187_v118_: D19
                d_2187_v118_ = D19_DC39(d_2186_v117_)
                d_2188_v119_: _dafny.Map
                d_2188_v119_ = _dafny.Map({d_2187_v118_: (0) - ((self).f6)})
                d_2189_v120_: _dafny.Map
                d_2189_v120_ = _dafny.Map({d_2169_v103_: len(d_2188_v119_)})
                d_2189_v120_ = (d_2189_v120_).set(d_2169_v103_, 235)
            d_2190_v121_: T1
            nw341_ = C1()
            nw341_.ctor__(669, (self).f6, self.f5, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kv"))), (d_2065_v14_).f4)
            d_2190_v121_ = nw341_
            d_2191_v122_: _dafny.Seq
            d_2191_v122_ = _dafny.SeqWithoutIsStrInference([default__.fm57(d_2048_v0_, d_2048_v0_, globalState), d_2168_v102_, _dafny.SeqWithoutIsStrInference([246 for d_2192_i8_ in range(default__.abs(875))]), d_2168_v102_])
            nw342_ = C5()
            nw342_.ctor__((d_2190_v121_).f4, default__.fm3(default__.fm3((self).f6, True, globalState), ((self).f10)[default__.safeIndex(211, ((self).f10).length(0))], globalState), d_2190_v121_.f5, len(d_2191_v122_), p0)
            d_2190_v121_ = nw342_
        elif True:
            d_2193___mcc_h5_ = source36_.cf12
            d_2194_cf12_ = d_2193___mcc_h5_
            d_2195_v123_: _dafny.Map
            d_2195_v123_ = _dafny.Map({d_2049_v1_: _dafny.CodePoint('a')})
            d_2196_v124_: _dafny.MultiSet
            d_2196_v124_ = _dafny.MultiSet([(d_2065_v14_).f4, True, (d_2065_v14_).f4])
            d_2197_v125_: _dafny.Map
            d_2197_v125_ = _dafny.Map({D2_DC6(): ((d_2067_v16_)[len(d_2195_v123_)] if (len(d_2195_v123_)) in (d_2067_v16_) else ((d_2196_v124_)[(d_2065_v14_).f4] if ((d_2065_v14_).f4) in (d_2196_v124_) else (d_2050_v2_).fm9(True, False, 632, globalState)))})
            d_2198_v126_: D2
            d_2198_v126_ = D2_DC6()
            d_2197_v125_ = (d_2197_v125_).set(d_2198_v126_, default__.safeModuloInt((self).f6, (self).f6))
            d_2199_v127_: _dafny.Array
            def lambda89_(d_2200_v0_):
                def lambda90_(d_2201_i9_):
                    return default__.safeDivisionInt(d_2201_i9_, d_2200_v0_)

                return lambda90_

            init57_ = lambda89_(d_2048_v0_)
            nw343_ = _dafny.Array(None, 25)
            for i0_57_ in range(nw343_.length(0)):
                nw343_[i0_57_] = init57_(i0_57_)
            d_2199_v127_ = nw343_
            d_2202_v128_: D16
            d_2202_v128_ = D16_DC32(d_2199_v127_)
            source37_ = (d_2202_v128_ if p0 else d_2202_v128_)
            if source37_.is_DC33:
                d_2203___mcc_h6_ = source37_.cf52
                d_2204___mcc_h7_ = source37_.cf53
                d_2205___mcc_h8_ = source37_.cf54
                d_2206___mcc_h9_ = source37_.cf55
                d_2207___mcc_h10_ = source37_.cf56
                d_2208_cf56_ = d_2207___mcc_h10_
                d_2209_cf55_ = d_2206___mcc_h9_
                d_2210_cf54_ = d_2205___mcc_h8_
                d_2211_cf53_ = d_2204___mcc_h7_
                d_2212_cf52_ = d_2203___mcc_h6_
                d_2213_v129_: _dafny.Map
                d_2213_v129_ = _dafny.Map({(d_2065_v14_).f4: d_2209_cf55_})
                d_2048_v0_ = default__.safeModuloInt((self).f6, (len(d_2213_v129_)) + (len(d_2049_v1_)))
                d_2214_v130_: str
                d_2214_v130_ = _dafny.CodePoint('f')
                r0 = d_2214_v130_
                d_2215_v131_: D4
                d_2215_v131_ = D4_DC12(_dafny.CodePoint('p'))
                d_2215_v131_ = d_2215_v131_
                d_2216_v132_: _dafny.Array
                def lambda91_(d_2217_v0_, d_2218_cf55_, d_2219_v130_):
                    def lambda92_(d_2220_i10_):
                        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f6, d_2217_v0_, (self).f6]), _dafny.SeqWithoutIsStrInference([d_2218_cf55_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dv"))), len(_dafny.Set({d_2219_v130_, d_2219_v130_}))]), _dafny.SeqWithoutIsStrInference([d_2217_v0_])})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([741])}))

                    return lambda92_

                init58_ = lambda91_(d_2048_v0_, d_2209_cf55_, d_2214_v130_)
                nw344_ = _dafny.Array(None, 6)
                for i0_58_ in range(nw344_.length(0)):
                    nw344_[i0_58_] = init58_(i0_58_)
                d_2216_v132_ = nw344_
                d_2221_v133_: _dafny.Seq
                d_2221_v133_ = _dafny.SeqWithoutIsStrInference([(0) - (d_2209_cf55_)])
                d_2222_v134_: _dafny.MultiSet
                d_2222_v134_ = _dafny.MultiSet([d_2065_v14_])
                d_2223_v135_: _dafny.Set
                d_2223_v135_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tm")))]), _dafny.SeqWithoutIsStrInference([d_2048_v0_ for d_2224_i11_ in range(default__.abs(283))]), d_2221_v133_, d_2221_v133_, _dafny.SeqWithoutIsStrInference([982, (d_2222_v134_).cardinality])})
                index344_ = default__.safeIndex(265, (d_2216_v132_).length(0))
                (d_2216_v132_)[index344_] = d_2223_v135_
                index345_ = default__.safeIndex(963, ((self).f10).length(0))
                ((self).f10)[index345_] = d_2212_cf52_
                d_2225_v136_: D6
                d_2225_v136_ = D6_DC16((d_2065_v14_).f4, (self).f4, d_2221_v133_)
                d_2226_v137_: _dafny.Set
                d_2226_v137_ = _dafny.Set({d_2225_v136_})
                d_2227_v138_: _dafny.Map
                d_2227_v138_ = _dafny.Map({default__.fm42(d_2214_v130_, d_2226_v137_, globalState): d_2214_v130_})
                index346_ = default__.safeIndex(895, (d_2199_v127_).length(0))
                (d_2199_v127_)[index346_] = d_2048_v0_
                d_2228_v139_: _dafny.Seq
                d_2228_v139_ = _dafny.SeqWithoutIsStrInference([d_2212_cf52_, d_2212_cf52_])
                d_2229_v140_: _dafny.Seq
                d_2229_v140_ = _dafny.SeqWithoutIsStrInference([d_2228_v139_, d_2228_v139_, d_2228_v139_])
                d_2230_v141_: D12
                d_2230_v141_ = D12_DC27(d_2229_v140_)
                d_2231_v142_: _dafny.Map
                d_2231_v142_ = _dafny.Map({(d_2065_v14_).f4: (self).f4})
                d_2232_v143_: _dafny.Map
                d_2232_v143_ = _dafny.Map({(self).f6: ((d_2231_v142_)[p0] if (p0) in (d_2231_v142_) else (d_2065_v14_).f4)})
                d_2233_v144_: D9
                d_2233_v144_ = D9_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mskr")))
                index347_ = default__.safeIndex(265, (d_2216_v132_).length(0))
                index348_ = default__.safeIndex(963, ((self).f10).length(0))
                index349_ = default__.safeIndex(895, (d_2199_v127_).length(0))
                rhs336_ = _dafny.Set({(d_2221_v133_) + (default__.fm39(d_2049_v1_, d_2230_v141_, ((d_2232_v143_)[len(d_2221_v133_)] if (len(d_2221_v133_)) in (d_2232_v143_) else p0), globalState))})
                rhs337_ = (self).f4
                rhs338_ = (d_2227_v138_) | (d_2227_v138_)
                rhs339_ = len((d_2233_v144_).cf25)
                lhs213_ = d_2216_v132_
                lhs214_ = default__.safeIndex(265, (d_2216_v132_).length(0))
                lhs215_ = (self).f10
                lhs216_ = default__.safeIndex(963, ((self).f10).length(0))
                lhs217_ = d_2199_v127_
                lhs218_ = default__.safeIndex(895, (d_2199_v127_).length(0))
                lhs213_[lhs214_] = rhs336_
                lhs215_[lhs216_] = rhs337_
                d_2227_v138_ = rhs338_
                lhs217_[lhs218_] = rhs339_
            elif True:
                d_2234___mcc_h11_ = source37_.cf51
                d_2235_cf51_ = d_2234___mcc_h11_
                d_2236_v145_: D3
                d_2236_v145_ = D3_DC10((self).f6)
                d_2048_v0_ = default__.safeModuloInt((d_2236_v145_).cf11, 765)
                d_2237_v146_: _dafny.Seq
                d_2237_v146_ = _dafny.SeqWithoutIsStrInference([False, (d_2050_v2_).fm8(d_2048_v0_, (self).f4, d_2049_v1_, _dafny.SeqWithoutIsStrInference([(self).f6 for d_2238_i12_ in range(default__.abs(133))]), globalState)])
                index350_ = default__.safeIndex(257, ((self).f10).length(0))
                ((self).f10)[index350_] = (d_2237_v146_)[default__.safeIndex((self).f6, len(d_2237_v146_))]
                d_2239_v147_: _dafny.Array
                def lambda93_(d_2240_i13_):
                    return D0_DC1()

                init59_ = lambda93_
                nw345_ = _dafny.Array(None, 12)
                for i0_59_ in range(nw345_.length(0)):
                    nw345_[i0_59_] = init59_(i0_59_)
                d_2239_v147_ = nw345_
                d_2241_v148_: bool
                d_2241_v148_ = True
                d_2242_v149_: str
                d_2242_v149_ = _dafny.CodePoint('p')
                d_2243_v150_: _dafny.Array
                nw346_ = _dafny.Array(None, 10)
                nw346_[int(0)] = d_2049_v1_
                nw346_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkjwa"))
                nw346_[int(2)] = d_2049_v1_
                nw346_[int(3)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc"))) + (d_2049_v1_)).set(default__.safeIndex((0) - (len(_dafny.Map({(d_2065_v14_).f4: d_2242_v149_}))), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc"))) + (d_2049_v1_))), _dafny.CodePoint('c'))
                nw346_[int(4)] = d_2049_v1_
                nw346_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtrm"))
                nw346_[int(6)] = d_2049_v1_
                nw346_[int(7)] = (d_2049_v1_ if (d_2065_v14_).f4 else d_2049_v1_)
                nw346_[int(8)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2244_i14_ in range(default__.abs(471))])
                nw346_[int(9)] = d_2049_v1_
                d_2243_v150_ = nw346_
                index351_ = default__.safeIndex(38, (d_2243_v150_).length(0))
                (d_2243_v150_)[index351_] = (d_2049_v1_ if (self).f4 else d_2049_v1_)
                index352_ = default__.safeIndex(257, ((self).f10).length(0))
                index353_ = default__.safeIndex(38, (d_2243_v150_).length(0))
                rhs340_ = (d_2065_v14_).f4
                rhs341_ = (D21_DC42(d_2239_v147_)).cf68
                rhs342_ = (d_2065_v14_).f4
                rhs343_ = d_2048_v0_
                rhs344_ = (d_2049_v1_) + (d_2049_v1_)
                lhs219_ = (self).f10
                lhs220_ = default__.safeIndex(257, ((self).f10).length(0))
                lhs221_ = d_2243_v150_
                lhs222_ = default__.safeIndex(38, (d_2243_v150_).length(0))
                lhs219_[lhs220_] = rhs340_
                d_2239_v147_ = rhs341_
                d_2241_v148_ = rhs342_
                d_2048_v0_ = rhs343_
                lhs221_[lhs222_] = rhs344_
                d_2245_v151_: _dafny.Map
                d_2245_v151_ = _dafny.Map({False: ((self).f10)[default__.safeIndex(257, ((self).f10).length(0))]})
                d_2246_v152_: _dafny.Set
                d_2246_v152_ = _dafny.Set({(self).f6})
                d_2247_v153_: _dafny.Map
                d_2247_v153_ = _dafny.Map({d_2245_v151_: len(d_2246_v152_)})
                rhs345_ = default__.fm0(_dafny.CodePoint('o'), globalState)
                rhs346_ = not((d_2247_v153_) == (d_2247_v153_))
                d_2241_v148_ = rhs345_
                d_2241_v148_ = rhs346_
                d_2048_v0_ = d_2048_v0_
            d_2248_v154_: _dafny.Map
            d_2248_v154_ = _dafny.Map({d_2049_v1_: len(d_2049_v1_)})
            d_2249_v155_: _dafny.Map
            d_2249_v155_ = _dafny.Map({d_2248_v154_: d_2048_v0_})
            d_2249_v155_ = (d_2249_v155_).set(_dafny.Map({d_2049_v1_: (0) - ((self).f6)}), (self).f6)
            d_2250_v156_: _dafny.Seq
            d_2250_v156_ = _dafny.SeqWithoutIsStrInference([d_2048_v0_])
            d_2251_v157_: _dafny.Set
            d_2251_v157_ = _dafny.Set({p0})
            d_2048_v0_ = (d_2250_v156_)[default__.safeIndex(len(d_2251_v157_), len(d_2250_v156_))]
        d_2252_v158_: bool
        d_2252_v158_ = True
        d_2252_v158_ = (self).f4
        d_2253_v159_: str
        d_2253_v159_ = _dafny.CodePoint('k')
        r0 = d_2253_v159_
        d_2048_v0_ = d_2048_v0_
        r0 = d_2253_v159_
        d_2254_v160_: _dafny.Array
        nw347_ = _dafny.Array(D1.default()(), 19)
        d_2254_v160_ = nw347_
        r1 = d_2254_v160_
        return r0, r1

    def m6(self, globalState):
        r0: D1 = D1.default()()
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_2255_i0_: int
        d_2255_i0_ = 0
        with _dafny.label("10"):
            while (self).f4:
                with _dafny.c_label("10"):
                    if (d_2255_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_2255_i0_ = (d_2255_i0_) + (1)
                    d_2256_v0_: _dafny.Array
                    nw348_ = _dafny.Array(_dafny.Map({}), 13)
                    d_2256_v0_ = nw348_
                    d_2257_v1_: _dafny.Map
                    d_2257_v1_ = _dafny.Map({(self).f10: (self).f4})
                    index354_ = default__.safeIndex(868, (d_2256_v0_).length(0))
                    (d_2256_v0_)[index354_] = (d_2257_v1_) | (d_2257_v1_)
                    index355_ = default__.safeIndex(868, (d_2256_v0_).length(0))
                    (d_2256_v0_)[index355_] = (d_2257_v1_) | (((_dafny.Map({(self).f10: (self).f4})).set((self).f10, (self).f4)) | (d_2257_v1_))
                    d_2258_v2_: _dafny.Array
                    nw349_ = _dafny.Array(int(0), 10)
                    d_2258_v2_ = nw349_
                    index356_ = default__.safeIndex(304, (d_2258_v2_).length(0))
                    (d_2258_v2_)[index356_] = (self).f6
                    d_2259_v3_: str
                    d_2259_v3_ = _dafny.CodePoint('w')
                    index357_ = default__.safeIndex(597, ((self).f10).length(0))
                    ((self).f10)[index357_] = default__.fm0(d_2259_v3_, globalState)
                    d_2260_v4_: bool
                    d_2260_v4_ = True
                    index358_ = default__.safeIndex(304, (d_2258_v2_).length(0))
                    index359_ = default__.safeIndex(597, ((self).f10).length(0))
                    rhs347_ = (self).f6
                    rhs348_ = (self).f4
                    rhs349_ = d_2260_v4_
                    lhs223_ = d_2258_v2_
                    lhs224_ = default__.safeIndex(304, (d_2258_v2_).length(0))
                    lhs225_ = (self).f10
                    lhs226_ = default__.safeIndex(597, ((self).f10).length(0))
                    lhs223_[lhs224_] = rhs347_
                    lhs225_[lhs226_] = rhs348_
                    d_2260_v4_ = rhs349_
                    d_2261_v5_: _dafny.Array
                    def lambda94_(d_2262_v2_):
                        def lambda95_(d_2263_i1_):
                            return (_dafny.Map({(self).f6: (d_2262_v2_)[default__.safeIndex(304, (d_2262_v2_).length(0))]})) | ((_dafny.Map({(self).f6: (self).f6})))

                        return lambda95_

                    init60_ = lambda94_(d_2258_v2_)
                    nw350_ = _dafny.Array(None, 13)
                    for i0_60_ in range(nw350_.length(0)):
                        nw350_[i0_60_] = init60_(i0_60_)
                    d_2261_v5_ = nw350_
                    d_2264_v6_: _dafny.Seq
                    d_2264_v6_ = _dafny.SeqWithoutIsStrInference([(D1_DC4((d_2258_v2_)[default__.safeIndex(304, (d_2258_v2_).length(0))], (d_2258_v2_)[default__.safeIndex(304, (d_2258_v2_).length(0))])).cf3])
                    d_2265_v7_: _dafny.Map
                    d_2265_v7_ = _dafny.Map({len(d_2264_v6_): (self).f6})
                    index360_ = default__.safeIndex(958, (d_2261_v5_).length(0))
                    (d_2261_v5_)[index360_] = d_2265_v7_
                    index361_ = default__.safeIndex(958, (d_2261_v5_).length(0))
                    (d_2261_v5_)[index361_] = d_2265_v7_
                    d_2266_v8_: C6
                    nw351_ = C6()
                    nw351_.ctor__(default__.fm0(d_2259_v3_, globalState), False, self.f5, default__.safeModuloInt((self).f6, (self).f6))
                    d_2266_v8_ = nw351_
                    pass
            pass
        hi13_ = (self).f6
        for d_2267_i2_ in range((self).f6, hi13_):
            index362_ = default__.safeIndex(324, ((self).f10).length(0))
            ((self).f10)[index362_] = (False) or (True)
            index363_ = default__.safeIndex(324, ((self).f10).length(0))
            ((self).f10)[index363_] = ((self).f6) >= (d_2267_i2_)
            index364_ = default__.safeIndex(324, ((self).f10).length(0))
            ((self).f10)[index364_] = ((self).f10)[default__.safeIndex(324, ((self).f10).length(0))]
            d_2268_v9_: _dafny.Array
            def lambda96_(d_2269_i3_):
                return _dafny.Map({True: ((self).f10)[default__.safeIndex(324, ((self).f10).length(0))]})

            init61_ = lambda96_
            nw352_ = _dafny.Array(None, 24)
            for i0_61_ in range(nw352_.length(0)):
                nw352_[i0_61_] = init61_(i0_61_)
            d_2268_v9_ = nw352_
            d_2270_v10_: D11
            d_2270_v10_ = D11_DC23(d_2268_v9_)
            d_2271_v11_: _dafny.Map
            d_2271_v11_ = _dafny.Map({(self).f6: d_2270_v10_})
            d_2272_v12_: int
            d_2272_v12_ = 472
            rhs350_ = d_2271_v11_
            rhs351_ = d_2267_i2_
            d_2271_v11_ = rhs350_
            d_2272_v12_ = rhs351_
            index365_ = default__.safeIndex(324, ((self).f10).length(0))
            index366_ = default__.safeIndex(324, ((self).f10).length(0))
            rhs352_ = not(((self).f10)[default__.safeIndex(324, ((self).f10).length(0))])
            rhs353_ = (0) - ((0) - ((d_2267_i2_) * (697)))
            rhs354_ = ((0) - (d_2272_v12_)) - ((self).fm7(globalState))
            rhs355_ = False
            lhs227_ = (self).f10
            lhs228_ = default__.safeIndex(324, ((self).f10).length(0))
            lhs229_ = (self).f10
            lhs230_ = default__.safeIndex(324, ((self).f10).length(0))
            lhs227_[lhs228_] = rhs352_
            d_2272_v12_ = rhs353_
            d_2272_v12_ = rhs354_
            lhs229_[lhs230_] = rhs355_
        d_2273_v13_: bool
        d_2273_v13_ = False
        d_2273_v13_ = d_2273_v13_
        index367_ = default__.safeIndex(304, ((self).f10).length(0))
        ((self).f10)[index367_] = (d_2273_v13_) or ((self).f4)
        index368_ = default__.safeIndex(304, ((self).f10).length(0))
        ((self).f10)[index368_] = d_2273_v13_
        d_2274_v14_: _dafny.Seq
        d_2274_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptmrel"))
        d_2275_v15_: _dafny.Seq
        d_2275_v15_ = _dafny.SeqWithoutIsStrInference([len((((self).f9)[(self).f4] if ((self).f4) in ((self).f9) else d_2274_v14_)), (self).f6, (self).f6, (self).f6])
        d_2276_v16_: _dafny.Seq
        d_2276_v16_ = _dafny.SeqWithoutIsStrInference([((self).f10)[default__.safeIndex(304, ((self).f10).length(0))]])
        d_2277_v17_: _dafny.MultiSet
        d_2277_v17_ = _dafny.MultiSet([d_2276_v16_])
        d_2278_v18_: _dafny.Map
        d_2278_v18_ = _dafny.Map({True: D6_DC15(d_2277_v17_)})
        if ((d_2275_v15_)[default__.safeIndex((self).f6, len(d_2275_v15_))]) > (len(d_2278_v18_)):
            d_2279_v19_: _dafny.Set
            d_2279_v19_ = _dafny.Set({(self).f6})
            d_2280_v20_: _dafny.Map
            d_2280_v20_ = _dafny.Map({((self).f10)[default__.safeIndex(304, ((self).f10).length(0))]: len(d_2279_v19_)})
            d_2281_v21_: D10
            d_2281_v21_ = D10_DC21(d_2280_v20_)
            source38_ = d_2281_v21_
            if source38_.is_DC22:
                d_2282___mcc_h0_ = source38_.cf30
                d_2283___mcc_h1_ = source38_.cf31
                d_2284___mcc_h2_ = source38_.cf32
                d_2285___mcc_h3_ = source38_.cf33
                d_2286___mcc_h4_ = source38_.cf34
                d_2287_cf34_ = d_2286___mcc_h4_
                d_2288_cf33_ = d_2285___mcc_h3_
                d_2289_cf32_ = d_2284___mcc_h2_
                d_2290_cf31_ = d_2283___mcc_h1_
                d_2291_cf30_ = d_2282___mcc_h0_
                d_2291_cf30_ = (d_2291_cf30_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlwyh"))) + (d_2291_cf30_))
                d_2292_v22_: _dafny.Array
                def lambda97_(d_2293_i4_):
                    return (d_2293_i4_) - (199)

                init62_ = lambda97_
                nw353_ = _dafny.Array(None, 28)
                for i0_62_ in range(nw353_.length(0)):
                    nw353_[i0_62_] = init62_(i0_62_)
                d_2292_v22_ = nw353_
                d_2294_v23_: _dafny.Map
                d_2294_v23_ = _dafny.Map({len(d_2275_v15_): (self).f6})
                index369_ = default__.safeIndex(968, (d_2292_v22_).length(0))
                (d_2292_v22_)[index369_] = ((d_2294_v23_)[len(d_2274_v14_)] if (len(d_2274_v14_)) in (d_2294_v23_) else d_2287_cf34_)
                index370_ = default__.safeIndex(968, (d_2292_v22_).length(0))
                (d_2292_v22_)[index370_] = (self).f6
                d_2295_v24_: D11
                d_2295_v24_ = D11_DC25(d_2290_cf31_, (self).f6)
                pat_let_tv40_ = d_2273_v13_
                d_2296_v25_: C5
                nw354_ = C5()
                def iife161_(_pat_let51_0):
                    def iife162_(d_2297_dt__update__tmp_h0_):
                        def iife163_(_pat_let52_0):
                            def iife164_(d_2298_dt__update_hcf39_h0_):
                                return D11_DC25(d_2298_dt__update_hcf39_h0_, (d_2297_dt__update__tmp_h0_).cf40)
                            return iife164_(_pat_let52_0)
                        return iife163_(pat_let_tv40_)
                    return iife162_(_pat_let51_0)
                nw354_.ctor__((iife161_(d_2295_v24_)).cf39, (0) - ((self).f6), self.f5, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "romg"))), ((self).f10)[default__.safeIndex(304, ((self).f10).length(0))])
                d_2296_v25_ = nw354_
                d_2299_v26_: _dafny.Array
                def lambda98_(d_2300_cf32_):
                    def lambda99_(d_2301_i5_):
                        return _dafny.MultiSet([d_2300_cf32_])

                    return lambda99_

                init63_ = lambda98_(d_2289_cf32_)
                nw355_ = _dafny.Array(None, 18)
                for i0_63_ in range(nw355_.length(0)):
                    nw355_[i0_63_] = init63_(i0_63_)
                d_2299_v26_ = nw355_
                d_2299_v26_ = d_2299_v26_
            elif True:
                d_2302___mcc_h5_ = source38_.cf29
                d_2303_cf29_ = d_2302___mcc_h5_
                d_2304_v27_: int
                d_2304_v27_ = 371
                d_2304_v27_ = d_2304_v27_
                d_2305_v28_: _dafny.MultiSet
                d_2305_v28_ = _dafny.MultiSet([((self).f10)[default__.safeIndex(304, ((self).f10).length(0))], (self).f4, (self).f4, d_2273_v13_])
                d_2306_v29_: _dafny.MultiSet
                d_2306_v29_ = _dafny.MultiSet([d_2304_v27_, (d_2305_v28_).cardinality])
                d_2307_v30_: _dafny.Map
                d_2307_v30_ = _dafny.Map({d_2306_v29_: d_2274_v14_})
                d_2308_v31_: bool
                d_2309_v32_: int
                out62_: bool
                out63_: int
                out62_, out63_ = default__.m0(d_2307_v30_, ((self).f10)[default__.safeIndex(304, ((self).f10).length(0))], d_2306_v29_, (253) > (len((d_2275_v15_).set(default__.safeIndex(d_2304_v27_, len(d_2275_v15_)), d_2304_v27_))), globalState)
                d_2308_v31_ = out62_
                d_2309_v32_ = out63_
                d_2304_v27_ = (self).f6
                index371_ = default__.safeIndex(304, ((self).f10).length(0))
                ((self).f10)[index371_] = (self).f4
            if d_2273_v13_:
                d_2310_v33_: _dafny.MultiSet
                d_2310_v33_ = _dafny.MultiSet([(self).f6])
                d_2311_v34_: _dafny.Map
                d_2311_v34_ = _dafny.Map({d_2310_v33_: d_2274_v14_})
                d_2312_v35_: bool
                d_2313_v36_: int
                out64_: bool
                out65_: int
                out64_, out65_ = default__.m0(d_2311_v34_, True, _dafny.MultiSet(d_2275_v15_), d_2273_v13_, globalState)
                d_2312_v35_ = out64_
                d_2313_v36_ = out65_
                d_2314_v37_: _dafny.Array
                def lambda100_(d_2315_v14_):
                    def lambda101_(d_2316_i6_):
                        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2317_i7_ in range(default__.abs(929))])})) | (_dafny.Set({d_2315_v14_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mj"))}))

                    return lambda101_

                init64_ = lambda100_(d_2274_v14_)
                nw356_ = _dafny.Array(None, 6)
                for i0_64_ in range(nw356_.length(0)):
                    nw356_[i0_64_] = init64_(i0_64_)
                d_2314_v37_ = nw356_
                d_2318_v38_: _dafny.Set
                d_2318_v38_ = _dafny.Set({d_2274_v14_})
                index372_ = default__.safeIndex(140, (d_2314_v37_).length(0))
                (d_2314_v37_)[index372_] = d_2318_v38_
                index373_ = default__.safeIndex(140, (d_2314_v37_).length(0))
                (d_2314_v37_)[index373_] = d_2318_v38_
                d_2313_v36_ = d_2313_v36_
                d_2319_v39_: C4
                nw357_ = C4()
                nw357_.ctor__(self.f5, (d_2275_v15_)[default__.safeIndex((self).f6, len(d_2275_v15_))], (self).f4)
                d_2319_v39_ = nw357_
                d_2320_v40_: D4
                d_2320_v40_ = D4_DC13(d_2313_v36_, d_2312_v35_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdwcp"))), d_2276_v16_)
                d_2321_v41_: str
                d_2321_v41_ = _dafny.CodePoint('i')
                rhs356_ = ((self).f6) - ((d_2320_v40_).cf14)
                rhs357_ = d_2273_v13_
                rhs358_ = len(((d_2274_v14_).set(default__.safeIndex(d_2313_v36_, len(d_2274_v14_)), d_2321_v41_)) + (((d_2274_v14_).set(default__.safeIndex((self).f6, len(d_2274_v14_)), d_2321_v41_)) + (d_2274_v14_)))
                rhs359_ = d_2313_v36_
                rhs360_ = (d_2313_v36_) + (439)
                d_2313_v36_ = rhs356_
                d_2312_v35_ = rhs357_
                d_2313_v36_ = rhs358_
                d_2313_v36_ = rhs359_
                d_2313_v36_ = rhs360_
            elif True:
                d_2322_v42_: _dafny.Array
                def lambda102_(d_2323_i8_):
                    return (d_2323_i8_) - ((self).f6)

                init65_ = lambda102_
                nw358_ = _dafny.Array(None, 26)
                for i0_65_ in range(nw358_.length(0)):
                    nw358_[i0_65_] = init65_(i0_65_)
                d_2322_v42_ = nw358_
                index374_ = default__.safeIndex(291, (d_2322_v42_).length(0))
                (d_2322_v42_)[index374_] = ((self).f6) - ((self).f6)
                index375_ = default__.safeIndex(291, (d_2322_v42_).length(0))
                (d_2322_v42_)[index375_] = (self).f6
                index376_ = default__.safeIndex(291, (d_2322_v42_).length(0))
                (d_2322_v42_)[index376_] = ((d_2322_v42_)[default__.safeIndex(291, (d_2322_v42_).length(0))]) * ((d_2322_v42_)[default__.safeIndex(291, (d_2322_v42_).length(0))])
                d_2324_v44_: _dafny.MultiSet
                d_2324_v44_ = _dafny.MultiSet(d_2275_v15_)
                d_2325_v45_: _dafny.Map
                d_2325_v45_ = _dafny.Map({d_2324_v44_: ((self).f10)[default__.safeIndex(304, ((self).f10).length(0))]})
                d_2326_v46_: _dafny.Map
                def iife165_():
                    coll59_ = _dafny.Map()
                    compr_59_: _dafny.MultiSet
                    for compr_59_ in (d_2325_v45_).keys.Elements:
                        d_2327_v43_: _dafny.MultiSet = compr_59_
                        if (d_2327_v43_) in (d_2325_v45_):
                            coll59_[d_2327_v43_] = len(d_2274_v14_)
                    return _dafny.Map(coll59_)
                d_2326_v46_ = _dafny.Map({len(iife165_()
                ): (self).f4})
                d_2326_v46_ = d_2326_v46_
                d_2273_v13_ = d_2273_v13_
                index377_ = default__.safeIndex(291, (d_2322_v42_).length(0))
                (d_2322_v42_)[index377_] = (self).f6
            d_2275_v15_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt((self).f6, len(d_2280_v20_)), (self).f6])
            d_2328_v47_: int
            d_2328_v47_ = 355
            d_2328_v47_ = ((self).f6 if d_2273_v13_ else d_2328_v47_)
            index378_ = default__.safeIndex(304, ((self).f10).length(0))
            ((self).f10)[index378_] = ((self).f10)[default__.safeIndex(304, ((self).f10).length(0))]
        elif True:
            d_2329_v48_: _dafny.Array
            nw359_ = _dafny.Array(D21.default()(), 1)
            d_2329_v48_ = nw359_
            d_2330_v49_: _dafny.Array
            def lambda103_(d_2331_i9_):
                return D0_DC1()

            init66_ = lambda103_
            nw360_ = _dafny.Array(None, 1)
            for i0_66_ in range(nw360_.length(0)):
                nw360_[i0_66_] = init66_(i0_66_)
            d_2330_v49_ = nw360_
            d_2332_v50_: D21
            d_2332_v50_ = D21_DC42(d_2330_v49_)
            pat_let_tv41_ = d_2330_v49_
            index379_ = default__.safeIndex(777, (d_2329_v48_).length(0))
            def iife166_(_pat_let53_0):
                def iife167_(d_2333_dt__update__tmp_h1_):
                    def iife168_(_pat_let54_0):
                        def iife169_(d_2334_dt__update_hcf68_h0_):
                            return D21_DC42(d_2334_dt__update_hcf68_h0_)
                        return iife169_(_pat_let54_0)
                    return iife168_(pat_let_tv41_)
                return iife167_(_pat_let53_0)
            (d_2329_v48_)[index379_] = iife166_(d_2332_v50_)
            index380_ = default__.safeIndex(777, (d_2329_v48_).length(0))
            (d_2329_v48_)[index380_] = d_2332_v50_
            d_2335_v51_: str
            d_2335_v51_ = _dafny.CodePoint('a')
            d_2335_v51_ = d_2335_v51_
            d_2336_v52_: _dafny.Set
            d_2336_v52_ = _dafny.Set({True})
            index381_ = default__.safeIndex(304, ((self).f10).length(0))
            ((self).f10)[index381_] = (d_2336_v52_).issubset(d_2336_v52_)
            d_2337_v53_: _dafny.Seq
            d_2337_v53_ = _dafny.SeqWithoutIsStrInference([d_2275_v15_, d_2275_v15_, d_2275_v15_])
            d_2338_v54_: _dafny.MultiSet
            d_2338_v54_ = _dafny.MultiSet([(self).f6, (self).f6, len(d_2337_v53_)])
            d_2339_v55_: _dafny.Map
            d_2339_v55_ = _dafny.Map({(d_2338_v54_).cardinality: ((self).f10)[default__.safeIndex(304, ((self).f10).length(0))]})
            index382_ = default__.safeIndex(304, ((self).f10).length(0))
            ((self).f10)[index382_] = (d_2273_v13_) == ((self).fm14(d_2339_v55_, not((self).f4), (self).f6, ((self).f10)[default__.safeIndex(304, ((self).f10).length(0))], globalState))
            d_2340_v56_: _dafny.Array
            def lambda104_(d_2341_i10_):
                return default__.safeDivisionInt(d_2341_i10_, (self).f6)

            init67_ = lambda104_
            nw361_ = _dafny.Array(None, 13)
            for i0_67_ in range(nw361_.length(0)):
                nw361_[i0_67_] = init67_(i0_67_)
            d_2340_v56_ = nw361_
            index383_ = default__.safeIndex(193, (d_2340_v56_).length(0))
            (d_2340_v56_)[index383_] = ((self).f6) + (default__.fm3((self).f6, True, globalState))
            index384_ = default__.safeIndex(193, (d_2340_v56_).length(0))
            (d_2340_v56_)[index384_] = (self).f6
        d_2342_v57_: str
        d_2342_v57_ = _dafny.CodePoint('n')
        if default__.fm0(d_2342_v57_, globalState):
            index385_ = default__.safeIndex(304, ((self).f10).length(0))
            ((self).f10)[index385_] = d_2273_v13_
            d_2343_v58_: C6
            nw362_ = C6()
            nw362_.ctor__((self).f4, d_2273_v13_, self.f5, ((self).f6 if False else (self).f6))
            d_2343_v58_ = nw362_
            d_2344_v59_: int
            d_2344_v59_ = -792
            index386_ = default__.safeIndex(304, ((self).f10).length(0))
            rhs361_ = d_2343_v58_
            rhs362_ = ((self).fm9((d_2343_v58_).f11, ((self).f10)[default__.safeIndex(304, ((self).f10).length(0))], len(d_2274_v14_), globalState)) == (d_2344_v59_)
            rhs363_ = d_2344_v59_
            lhs231_ = (self).f10
            lhs232_ = default__.safeIndex(304, ((self).f10).length(0))
            d_2343_v58_ = rhs361_
            lhs231_[lhs232_] = rhs362_
            d_2344_v59_ = rhs363_
            d_2344_v59_ = (self).f6
            d_2344_v59_ = d_2344_v59_
            d_2345_v60_: _dafny.Array
            def lambda105_(d_2346_v59_):
                def lambda106_(d_2347_i11_):
                    return (d_2347_i11_) * (d_2346_v59_)

                return lambda106_

            init68_ = lambda105_(d_2344_v59_)
            nw363_ = _dafny.Array(None, 15)
            for i0_68_ in range(nw363_.length(0)):
                nw363_[i0_68_] = init68_(i0_68_)
            d_2345_v60_ = nw363_
            index387_ = default__.safeIndex(231, (d_2345_v60_).length(0))
            (d_2345_v60_)[index387_] = default__.safeDivisionInt(d_2344_v59_, (self).f6)
            d_2348_v61_: _dafny.MultiSet
            d_2348_v61_ = _dafny.MultiSet([(0) - (d_2344_v59_)])
            d_2349_v62_: _dafny.Map
            d_2349_v62_ = _dafny.Map({d_2348_v61_: d_2273_v13_})
            d_2350_v63_: _dafny.MultiSet
            d_2350_v63_ = d_2348_v61_
            d_2351_v64_: _dafny.Map
            d_2351_v64_ = _dafny.Map({d_2344_v59_: (d_2350_v63_)})
            d_2352_v65_: _dafny.Map
            d_2352_v65_ = _dafny.Map({((self).f10)[default__.safeIndex(304, ((self).f10).length(0))]: ((d_2349_v62_)[((d_2351_v64_)[d_2344_v59_] if (d_2344_v59_) in (d_2351_v64_) else d_2348_v61_)] if (((d_2351_v64_)[d_2344_v59_] if (d_2344_v59_) in (d_2351_v64_) else d_2348_v61_)) in (d_2349_v62_) else (d_2343_v58_).f11)})
            index388_ = default__.safeIndex(231, (d_2345_v60_).length(0))
            rhs364_ = (0) - ((d_2344_v59_) - ((len(d_2275_v15_)) * (d_2344_v59_)))
            rhs365_ = ((d_2352_v65_)[(d_2343_v58_).f11] if ((d_2343_v58_).f11) in (d_2352_v65_) else (d_2343_v58_).f11)
            rhs366_ = (self).fm7(globalState)
            rhs367_ = d_2343_v58_
            lhs233_ = d_2345_v60_
            lhs234_ = default__.safeIndex(231, (d_2345_v60_).length(0))
            lhs233_[lhs234_] = rhs364_
            d_2273_v13_ = rhs365_
            d_2344_v59_ = rhs366_
            d_2343_v58_ = rhs367_
        elif True:
            d_2353_v66_: int
            d_2353_v66_ = 313
            d_2353_v66_ = default__.safeDivisionInt(d_2353_v66_, ((self).f6) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpm")))))
            d_2354_v67_: _dafny.Array
            nw364_ = _dafny.Array(False, 27)
            d_2354_v67_ = nw364_
            d_2354_v67_ = d_2354_v67_
            d_2355_v68_: D9
            d_2355_v68_ = D9_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")))
            d_2356_v69_: _dafny.Map
            d_2356_v69_ = _dafny.Map({d_2355_v68_: (self).f4})
            d_2357_v70_: _dafny.Map
            d_2357_v70_ = _dafny.Map({(d_2356_v69_) | (d_2356_v69_): d_2342_v57_})
            d_2357_v70_ = (d_2357_v70_).set(d_2356_v69_, d_2342_v57_)
            d_2353_v66_ = (0) - ((d_2353_v66_) + (d_2353_v66_))
            index389_ = default__.safeIndex(304, ((self).f10).length(0))
            ((self).f10)[index389_] = (((self).f6) == ((self).f6)) and (not((d_2275_v15_) != (d_2275_v15_)))
        d_2358_v71_: _dafny.Set
        d_2358_v71_ = _dafny.Set({(self).f6})
        d_2359_v72_: _dafny.Map
        d_2359_v72_ = _dafny.Map({(0) - ((self).f6): not(not(d_2273_v13_))})
        d_2360_v73_: _dafny.Array
        nw365_ = _dafny.Array(None, 28)
        nw365_[int(0)] = (self).f6
        nw365_[int(1)] = (self).f6
        nw365_[int(2)] = (self).f6
        nw365_[int(3)] = len(d_2358_v71_)
        nw365_[int(4)] = (self).f6
        nw365_[int(5)] = (self).f6
        nw365_[int(6)] = (self).f6
        nw365_[int(7)] = 617
        nw365_[int(8)] = (self).f6
        nw365_[int(9)] = 844
        nw365_[int(10)] = (self).f6
        nw365_[int(11)] = (self).f6
        nw365_[int(12)] = -957
        nw365_[int(13)] = (self).f6
        nw365_[int(14)] = 860
        nw365_[int(15)] = len(d_2274_v14_)
        nw365_[int(16)] = (0) - ((self).f6)
        nw365_[int(17)] = (self).f6
        nw365_[int(18)] = (self).f6
        nw365_[int(19)] = (self).f6
        nw365_[int(20)] = len(d_2274_v14_)
        nw365_[int(21)] = (self).f6
        nw365_[int(22)] = (self).f6
        nw365_[int(23)] = (self).f6
        nw365_[int(24)] = (self).f6
        nw365_[int(25)] = (self).f6
        nw365_[int(26)] = len(d_2359_v72_)
        nw365_[int(27)] = 682
        d_2360_v73_ = nw365_
        d_2361_v74_: _dafny.Map
        d_2361_v74_ = _dafny.Map({d_2360_v73_: (self).f6})
        pat_let_tv42_ = d_2361_v74_
        def iife170_(_pat_let55_0):
            def iife171_(d_2363_dt__update__tmp_h2_):
                def iife172_(_pat_let56_0):
                    def iife173_(d_2364_dt__update_hcf2_h0_):
                        return D1_DC3(d_2364_dt__update_hcf2_h0_)
                    return iife173_(_pat_let56_0)
                return iife172_(((self).f6) >= (len(pat_let_tv42_)))
            return iife171_(_pat_let55_0)
        r0 = iife170_(D1_DC3((self).fm8((self).f6, d_2273_v13_, _dafny.SeqWithoutIsStrInference([d_2342_v57_ for d_2362_i12_ in range(default__.abs(-493))]), d_2275_v15_, globalState)))
        r1 = d_2274_v14_
        return r0, r1

    @property
    def f10(self):
        return self._f10
    @property
    def f9(self):
        return self._f9

class C8:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self):
        pass
        pass

    def fm11(self, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prxjdttbg"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "or")))

    def fm12(self, p0, p1, p2, p3, globalState):
        return (len(_dafny.Map({True: _dafny.Map({854: len(_dafny.SeqWithoutIsStrInference([True, False]))})}))) not in ((_dafny.SeqWithoutIsStrInference([-632 for d_2365_i0_ in range(default__.abs(157))])) + (_dafny.SeqWithoutIsStrInference([-337, len(_dafny.Map({False: False}))])))

    def m4(self, p0, p1, p2, globalState):
        d_2366_v0_: D0
        d_2366_v0_ = D0_DC0(986)
        d_2367_i0_: int
        d_2367_i0_ = 0
        with _dafny.label("11"):
            while not((self).fm12(p2, d_2366_v0_, p1, p2, globalState)):
                with _dafny.c_label("11"):
                    if (d_2367_i0_) >= (100):
                        raise _dafny.Break("11")
                    d_2367_i0_ = (d_2367_i0_) + (1)
                    d_2368_v1_: str
                    d_2368_v1_ = _dafny.CodePoint('h')
                    d_2368_v1_ = d_2368_v1_
                    d_2369_v2_: _dafny.Seq
                    d_2369_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhlv"))
                    d_2369_v2_ = (d_2369_v2_) + (d_2369_v2_)
                    d_2370_v3_: bool
                    d_2370_v3_ = True
                    d_2371_v4_: _dafny.Array
                    nw366_ = _dafny.Array(int(0), 25)
                    d_2371_v4_ = nw366_
                    rhs368_ = p2
                    rhs369_ = d_2371_v4_
                    d_2370_v3_ = rhs368_
                    d_2371_v4_ = rhs369_
                    d_2372_v5_: int
                    d_2372_v5_ = 337
                    rhs370_ = ((0) - (p0)) >= (len(d_2369_v2_))
                    rhs371_ = len(d_2369_v2_)
                    d_2370_v3_ = rhs370_
                    d_2372_v5_ = rhs371_
                    pass
            pass
        d_2373_v6_: int
        d_2373_v6_ = 631
        d_2373_v6_ = (p1) * (p0)
        d_2374_v7_: _dafny.Array
        nw367_ = _dafny.Array(int(0), 4)
        d_2374_v7_ = nw367_
        index390_ = default__.safeIndex(803, (d_2374_v7_).length(0))
        (d_2374_v7_)[index390_] = (194 if True else p1)
        d_2375_v8_: D11
        d_2375_v8_ = D11_DC26(p2)
        d_2376_v9_: _dafny.Map
        d_2376_v9_ = _dafny.Map({p2: p2})
        d_2377_v10_: _dafny.Seq
        d_2377_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "drtvj"))
        d_2378_v11_: D12
        d_2378_v11_ = D12_DC28(p0, len(d_2377_v10_), p2, not(p2), p0)
        d_2379_v12_: str
        d_2379_v12_ = _dafny.CodePoint('a')
        d_2380_v13_: _dafny.Seq
        d_2380_v13_ = _dafny.SeqWithoutIsStrInference([p2])
        d_2381_v14_: D21
        d_2381_v14_ = D21_DC43(p2, False, d_2380_v13_)
        d_2382_v15_: _dafny.Seq
        d_2382_v15_ = _dafny.SeqWithoutIsStrInference([p2, (d_2381_v14_).cf70, p2])
        d_2383_v16_: _dafny.Array
        nw368_ = _dafny.Array(None, 21)
        nw368_[int(0)] = p2
        nw368_[int(1)] = p2
        nw368_[int(2)] = True
        nw368_[int(3)] = p2
        nw368_[int(4)] = p2
        nw368_[int(5)] = (d_2375_v8_).cf41
        nw368_[int(6)] = p2
        nw368_[int(7)] = p2
        nw368_[int(8)] = p2
        nw368_[int(9)] = p2
        nw368_[int(10)] = p2
        nw368_[int(11)] = p2
        nw368_[int(12)] = p2
        nw368_[int(13)] = p2
        nw368_[int(14)] = p2
        nw368_[int(15)] = p2
        nw368_[int(16)] = ((d_2376_v9_)[p2] if (p2) in (d_2376_v9_) else True)
        nw368_[int(17)] = p2
        nw368_[int(18)] = (D9_DC20((d_2378_v11_).cf46, d_2379_v12_, d_2373_v6_)).cf26
        nw368_[int(19)] = not((d_2380_v13_)[default__.safeIndex(p0, len(d_2380_v13_))])
        nw368_[int(20)] = p2
        d_2383_v16_ = nw368_
        d_2384_v17_: _dafny.Array
        nw369_ = _dafny.Array(None, 28)
        nw369_[int(0)] = d_2383_v16_
        nw369_[int(1)] = d_2383_v16_
        nw369_[int(2)] = d_2383_v16_
        nw369_[int(3)] = d_2383_v16_
        nw369_[int(4)] = d_2383_v16_
        nw369_[int(5)] = d_2383_v16_
        nw369_[int(6)] = d_2383_v16_
        nw369_[int(7)] = d_2383_v16_
        nw369_[int(8)] = d_2383_v16_
        nw369_[int(9)] = d_2383_v16_
        nw369_[int(10)] = d_2383_v16_
        nw369_[int(11)] = d_2383_v16_
        nw369_[int(12)] = d_2383_v16_
        nw369_[int(13)] = d_2383_v16_
        nw369_[int(14)] = d_2383_v16_
        nw369_[int(15)] = d_2383_v16_
        nw369_[int(16)] = d_2383_v16_
        nw369_[int(17)] = d_2383_v16_
        nw369_[int(18)] = d_2383_v16_
        nw369_[int(19)] = d_2383_v16_
        nw369_[int(20)] = d_2383_v16_
        nw369_[int(21)] = d_2383_v16_
        nw369_[int(22)] = d_2383_v16_
        nw369_[int(23)] = d_2383_v16_
        nw369_[int(24)] = d_2383_v16_
        nw369_[int(25)] = d_2383_v16_
        nw369_[int(26)] = d_2383_v16_
        nw369_[int(27)] = d_2383_v16_
        d_2384_v17_ = nw369_
        d_2385_v18_: _dafny.MultiSet
        d_2385_v18_ = _dafny.MultiSet([-502])
        d_2386_v19_: T1
        nw370_ = C3()
        nw370_.ctor__(d_2384_v17_, (d_2373_v6_) * (p0), (d_2385_v18_).ispropersubset(d_2385_v18_))
        d_2386_v19_ = nw370_
        arr13_ = d_2386_v19_.f5
        index391_ = default__.safeIndex(632, (d_2386_v19_.f5).length(0))
        arr13_[index391_] = d_2383_v16_
        d_2387_v20_: _dafny.Map
        d_2387_v20_ = _dafny.Map({p0: d_2373_v6_})
        d_2388_v21_: _dafny.MultiSet
        d_2388_v21_ = _dafny.MultiSet([(d_2382_v15_).set(default__.safeIndex(len(d_2387_v20_), len(d_2382_v15_)), (d_2386_v19_).f4)])
        d_2389_v22_: D6
        d_2389_v22_ = D6_DC15(d_2388_v21_)
        pat_let_tv43_ = d_2379_v12_
        pat_let_tv44_ = d_2377_v10_
        pat_let_tv45_ = d_2373_v6_
        pat_let_tv46_ = d_2377_v10_
        index392_ = default__.safeIndex(803, (d_2374_v7_).length(0))
        arr14_ = d_2386_v19_.f5
        index393_ = default__.safeIndex(632, (d_2386_v19_.f5).length(0))
        def lambda107_(source39_):
            if source39_.is_DC16:
                d_2390___mcc_h0_ = source39_.cf20
                d_2391___mcc_h1_ = source39_.cf21
                d_2392___mcc_h2_ = source39_.cf22
                d_2393_cf22_ = d_2392___mcc_h2_
                d_2394_cf21_ = d_2391___mcc_h1_
                d_2395_cf20_ = d_2390___mcc_h0_
                return pat_let_tv43_
            elif True:
                d_2396___mcc_h3_ = source39_.cf19
                d_2397_cf19_ = d_2396___mcc_h3_
                return (pat_let_tv44_)[default__.safeIndex(pat_let_tv45_, len(pat_let_tv46_))]

        rhs372_ = p0
        rhs373_ = d_2386_v19_
        rhs374_ = lambda107_(d_2389_v22_)
        rhs375_ = d_2383_v16_
        lhs235_ = d_2374_v7_
        lhs236_ = default__.safeIndex(803, (d_2374_v7_).length(0))
        lhs237_ = d_2386_v19_.f5
        lhs238_ = default__.safeIndex(632, (d_2386_v19_.f5).length(0))
        lhs235_[lhs236_] = rhs372_
        d_2386_v19_ = rhs373_
        d_2379_v12_ = rhs374_
        lhs237_[lhs238_] = rhs375_
        d_2398_i1_: int
        d_2398_i1_ = 0
        with _dafny.label("12"):
            while (d_2386_v19_).f4:
                with _dafny.c_label("12"):
                    if (d_2398_i1_) >= (100):
                        raise _dafny.Break("12")
                    d_2398_i1_ = (d_2398_i1_) + (1)
                    d_2374_v7_ = d_2374_v7_
                    d_2399_v23_: _dafny.Array
                    def lambda108_(d_2400_v12_):
                        def lambda109_(d_2401_i2_):
                            return _dafny.SeqWithoutIsStrInference([d_2400_v12_ for d_2402_i3_ in range(default__.abs(205))])

                        return lambda109_

                    init69_ = lambda108_(d_2379_v12_)
                    nw371_ = _dafny.Array(None, 23)
                    for i0_69_ in range(nw371_.length(0)):
                        nw371_[i0_69_] = init69_(i0_69_)
                    d_2399_v23_ = nw371_
                    d_2403_v24_: D1
                    d_2403_v24_ = D1_DC4((d_2374_v7_)[default__.safeIndex(803, (d_2374_v7_).length(0))], d_2373_v6_)
                    d_2404_v25_: _dafny.Seq
                    d_2404_v25_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([default__.fm47((d_2386_v19_).f6, globalState)]))])
                    d_2405_v26_: C1
                    nw372_ = C1()
                    nw372_.ctor__((d_2374_v7_)[default__.safeIndex(803, (d_2374_v7_).length(0))], -616, d_2386_v19_.f5, len(d_2380_v13_), p2)
                    d_2405_v26_ = nw372_
                    index394_ = default__.safeIndex(626, (d_2399_v23_).length(0))
                    (d_2399_v23_)[index394_] = default__.fm50((d_2403_v24_).cf3, len((d_2404_v25_).set(default__.safeIndex(p1, len(d_2404_v25_)), len(_dafny.SeqWithoutIsStrInference([p1, (d_2386_v19_).f6])))), len(_dafny.Map({default__.fm0(d_2379_v12_, globalState): d_2405_v26_})), (d_2386_v19_).f4, globalState)
                    index395_ = default__.safeIndex(626, (d_2399_v23_).length(0))
                    (d_2399_v23_)[index395_] = (((d_2377_v10_) + (d_2377_v10_) if (d_2386_v19_).f4 else d_2377_v10_)).set(default__.safeIndex((len(d_2377_v10_)) + (len(d_2377_v10_)), len(((d_2377_v10_) + (d_2377_v10_) if (d_2386_v19_).f4 else d_2377_v10_))), d_2379_v12_)
                    d_2406_v27_: _dafny.Map
                    d_2406_v27_ = _dafny.Map({_dafny.CodePoint('a'): p0})
                    d_2406_v27_ = (d_2406_v27_).set(d_2379_v12_, p1)
                    d_2384_v17_ = d_2386_v19_.f5
                    pass
            pass
        d_2407_v28_: _dafny.MultiSet
        d_2407_v28_ = d_2385_v18_
        d_2408_v29_: _dafny.Seq
        d_2408_v29_ = _dafny.SeqWithoutIsStrInference([d_2385_v18_, d_2407_v28_, d_2407_v28_])
        d_2409_v30_: _dafny.Set
        d_2409_v30_ = _dafny.Set({d_2408_v29_})
        d_2409_v30_ = d_2409_v30_
        d_2410_v31_: bool
        d_2410_v31_ = False
        d_2410_v31_ = (d_2386_v19_).f4

    def m5(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        d_2411_v0_: bool
        d_2411_v0_ = False
        if d_2411_v0_:
            r0 = not(d_2411_v0_)
            d_2412_v1_: _dafny.Set
            d_2412_v1_ = _dafny.Set({d_2411_v0_, d_2411_v0_})
            r1 = (d_2412_v1_).issubset(d_2412_v1_)
            d_2413_v2_: _dafny.Map
            d_2413_v2_ = _dafny.Map({p0: (p0) - (len(_dafny.SeqWithoutIsStrInference([259])))})
            d_2413_v2_ = (d_2413_v2_).set(p0, p0)
            r1 = (p0) <= (p0)
            d_2414_v3_: _dafny.Array
            nw373_ = _dafny.Array(_dafny.Array(None, 0), 4)
            d_2414_v3_ = nw373_
            d_2415_v4_: C3
            nw374_ = C3()
            nw374_.ctor__(d_2414_v3_, p0, d_2411_v0_)
            d_2415_v4_ = nw374_
        elif True:
            d_2416_v5_: _dafny.Seq
            d_2416_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
            d_2411_v0_ = ((p0) * ((0) - (len(d_2416_v5_)))) == (793)
            if not(d_2411_v0_):
                d_2417_v6_: D2
                d_2417_v6_ = D2_DC6()
                d_2417_v6_ = d_2417_v6_
                d_2418_v7_: _dafny.Seq
                d_2418_v7_ = _dafny.SeqWithoutIsStrInference([270, (0) - (p0)])
                d_2419_v8_: _dafny.Seq
                d_2419_v8_ = _dafny.SeqWithoutIsStrInference([d_2411_v0_, d_2411_v0_])
                d_2420_v9_: _dafny.MultiSet
                d_2420_v9_ = _dafny.MultiSet([(_dafny.MultiSet(d_2419_v8_)).cardinality])
                d_2421_v10_: _dafny.Array
                nw375_ = _dafny.Array(None, 23)
                nw375_[int(0)] = p0
                nw375_[int(1)] = len(d_2416_v5_)
                nw375_[int(2)] = p0
                nw375_[int(3)] = len(d_2418_v7_)
                nw375_[int(4)] = -91
                nw375_[int(5)] = p0
                nw375_[int(6)] = p0
                nw375_[int(7)] = p0
                nw375_[int(8)] = (d_2420_v9_).cardinality
                nw375_[int(9)] = p0
                nw375_[int(10)] = p0
                nw375_[int(11)] = p0
                nw375_[int(12)] = -870
                nw375_[int(13)] = p0
                nw375_[int(14)] = p0
                nw375_[int(15)] = p0
                nw375_[int(16)] = p0
                nw375_[int(17)] = p0
                nw375_[int(18)] = 951
                nw375_[int(19)] = 424
                nw375_[int(20)] = p0
                nw375_[int(21)] = p0
                nw375_[int(22)] = p0
                d_2421_v10_ = nw375_
                d_2422_v11_: _dafny.Set
                d_2422_v11_ = _dafny.Set({d_2421_v10_, d_2421_v10_, d_2421_v10_})
                d_2423_v12_: _dafny.Set
                d_2423_v12_ = d_2422_v11_
                d_2424_v13_: _dafny.Set
                d_2424_v13_ = _dafny.Set({d_2422_v11_, d_2423_v12_})
                d_2411_v0_ = (d_2424_v13_).ispropersubset(d_2424_v13_)
                d_2425_v14_: _dafny.Map
                d_2425_v14_ = _dafny.Map({(p0) * (p0): d_2411_v0_})
                d_2425_v14_ = (d_2425_v14_).set((p0) - ((d_2418_v7_)[default__.safeIndex(-179, len(d_2418_v7_))]), d_2411_v0_)
                d_2426_v15_: str
                d_2426_v15_ = _dafny.CodePoint('d')
                d_2427_v16_: _dafny.Map
                d_2427_v16_ = _dafny.Map({d_2411_v0_: d_2426_v15_})
                d_2428_v17_: _dafny.Map
                d_2428_v17_ = _dafny.Map({d_2411_v0_: len(d_2427_v16_)})
                d_2429_v18_: D10
                d_2429_v18_ = D10_DC21(d_2428_v17_)
                d_2430_v19_: _dafny.Set
                d_2430_v19_ = _dafny.Set({p0, 85})
                d_2431_v22_: _dafny.Array
                nw376_ = _dafny.Array(None, 19)
                nw376_[int(0)] = (d_2430_v19_).intersection(d_2430_v19_)
                nw376_[int(1)] = (d_2430_v19_) - (d_2430_v19_)
                nw376_[int(2)] = default__.fm31(len(d_2425_v14_), p0, p0, False, globalState)
                nw376_[int(3)] = d_2430_v19_
                nw376_[int(4)] = _dafny.Set({len(d_2416_v5_)})
                nw376_[int(5)] = d_2430_v19_
                nw376_[int(6)] = (_dafny.Set({738, p0})) - (d_2430_v19_)
                nw376_[int(7)] = d_2430_v19_
                nw376_[int(8)] = d_2430_v19_
                nw376_[int(9)] = d_2430_v19_
                nw376_[int(10)] = d_2430_v19_
                def iife174_():
                    coll60_ = _dafny.Set()
                    compr_60_: int
                    for compr_60_ in _dafny.IntegerRange(141, 982):
                        d_2432_v20_: int = compr_60_
                        if ((141) <= (d_2432_v20_)) and ((d_2432_v20_) < (982)):
                            coll60_ = coll60_.union(_dafny.Set([(d_2432_v20_) - (p0)]))
                    return _dafny.Set(coll60_)
                nw376_[int(11)] = iife174_()
                
                nw376_[int(12)] = _dafny.Set({230})
                nw376_[int(13)] = d_2430_v19_
                nw376_[int(14)] = d_2430_v19_
                nw376_[int(15)] = (_dafny.Set({p0})).intersection(d_2430_v19_)
                def iife175_():
                    coll61_ = _dafny.Set()
                    compr_61_: int
                    for compr_61_ in _dafny.IntegerRange(-129, 732):
                        d_2433_v21_: int = compr_61_
                        if ((-129) <= (d_2433_v21_)) and ((d_2433_v21_) < (732)):
                            coll61_ = coll61_.union(_dafny.Set([default__.safeDivisionInt(d_2433_v21_, p0)]))
                    return _dafny.Set(coll61_)
                nw376_[int(16)] = iife175_()
                
                nw376_[int(17)] = d_2430_v19_
                nw376_[int(18)] = d_2430_v19_
                d_2431_v22_ = nw376_
                index396_ = default__.safeIndex(947, (d_2431_v22_).length(0))
                (d_2431_v22_)[index396_] = (d_2430_v19_) - (d_2430_v19_)
                d_2434_v23_: _dafny.Seq
                d_2434_v23_ = _dafny.SeqWithoutIsStrInference([d_2416_v5_])
                d_2435_v24_: _dafny.MultiSet
                d_2435_v24_ = _dafny.MultiSet([d_2411_v0_, False, not(d_2411_v0_)])
                d_2436_v25_: _dafny.Set
                d_2436_v25_ = _dafny.Set({d_2411_v0_})
                d_2437_v26_: D16
                d_2437_v26_ = D16_DC33(d_2411_v0_, len(d_2416_v5_), d_2435_v24_, p0, d_2436_v25_)
                d_2438_v27_: _dafny.Set
                d_2438_v27_ = _dafny.Set({p0, p0, p0, p0, p0})
                index397_ = default__.safeIndex(947, (d_2431_v22_).length(0))
                rhs376_ = not (False) or ((d_2426_v15_) not in ((d_2434_v23_)[default__.safeIndex(p0, len(d_2434_v23_))]))
                rhs377_ = D10_DC21(d_2428_v17_)
                rhs378_ = (d_2437_v26_).cf52
                rhs379_ = (d_2438_v27_)
                lhs239_ = d_2431_v22_
                lhs240_ = default__.safeIndex(947, (d_2431_v22_).length(0))
                d_2411_v0_ = rhs376_
                d_2429_v18_ = rhs377_
                r1 = rhs378_
                lhs239_[lhs240_] = rhs379_
                d_2439_v28_: int
                d_2439_v28_ = 113
                d_2439_v28_ = (0) - (p0)
            elif True:
                d_2440_v29_: int
                d_2440_v29_ = 868
                d_2440_v29_ = (0) - (p0)
                r0 = d_2411_v0_
                d_2441_v30_: _dafny.MultiSet
                d_2441_v30_ = _dafny.MultiSet([d_2440_v29_])
                d_2442_v31_: _dafny.Seq
                d_2442_v31_ = _dafny.SeqWithoutIsStrInference([d_2441_v30_])
                d_2440_v29_ = len(d_2442_v31_)
                d_2443_v32_: _dafny.Array
                nw377_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_2443_v32_ = nw377_
                d_2444_v33_: _dafny.Map
                d_2444_v33_ = _dafny.Map({p0: p0})
                d_2445_v34_: _dafny.Map
                d_2445_v34_ = _dafny.Map({d_2444_v33_: d_2411_v0_})
                d_2446_v35_: C4
                nw378_ = C4()
                nw378_.ctor__(d_2443_v32_, p0, (default__.fm67(d_2411_v0_, globalState)) == (d_2445_v34_))
                d_2446_v35_ = nw378_
                d_2447_v36_: _dafny.Array
                def lambda110_(d_2448_v0_):
                    def lambda111_(d_2449_i0_):
                        return (_dafny.Set({d_2448_v0_, d_2448_v0_})).issubset(_dafny.Set({d_2448_v0_, d_2448_v0_}))

                    return lambda111_

                init70_ = lambda110_(d_2411_v0_)
                nw379_ = _dafny.Array(None, 28)
                for i0_70_ in range(nw379_.length(0)):
                    nw379_[i0_70_] = init70_(i0_70_)
                d_2447_v36_ = nw379_
                d_2450_v37_: D21
                d_2450_v37_ = D21_DC44(d_2411_v0_)
                pat_let_tv47_ = d_2411_v0_
                index398_ = default__.safeIndex(521, (d_2447_v36_).length(0))
                def iife176_(_pat_let57_0):
                    def iife177_(d_2451_dt__update__tmp_h1_):
                        def iife178_(_pat_let58_0):
                            def iife179_(d_2452_dt__update_hcf72_h0_):
                                return D21_DC44(d_2452_dt__update_hcf72_h0_)
                            return iife179_(_pat_let58_0)
                        return iife178_(pat_let_tv47_)
                    return iife177_(_pat_let57_0)
                (d_2447_v36_)[index398_] = (iife176_(d_2450_v37_)).cf72
                index399_ = default__.safeIndex(521, (d_2447_v36_).length(0))
                (d_2447_v36_)[index399_] = True
            d_2416_v5_ = d_2416_v5_
            d_2453_v38_: _dafny.Array
            def lambda112_(d_2454_p0_):
                def lambda113_(d_2455_i1_):
                    return (d_2455_i1_) + (d_2454_p0_)

                return lambda113_

            init71_ = lambda112_(p0)
            nw380_ = _dafny.Array(None, 4)
            for i0_71_ in range(nw380_.length(0)):
                nw380_[i0_71_] = init71_(i0_71_)
            d_2453_v38_ = nw380_
            index400_ = default__.safeIndex(377, (d_2453_v38_).length(0))
            (d_2453_v38_)[index400_] = 647
            d_2456_v39_: _dafny.MultiSet
            d_2456_v39_ = _dafny.MultiSet([not(d_2411_v0_)])
            index401_ = default__.safeIndex(377, (d_2453_v38_).length(0))
            (d_2453_v38_)[index401_] = ((d_2456_v39_).cardinality) * (p0)
            d_2457_v40_: _dafny.Seq
            d_2457_v40_ = _dafny.SeqWithoutIsStrInference([d_2411_v0_, d_2411_v0_, d_2411_v0_, d_2411_v0_])
            d_2458_v41_: D4
            d_2458_v41_ = D4_DC13(137, d_2411_v0_, len(d_2416_v5_), d_2457_v40_)
            d_2459_v42_: _dafny.MultiSet
            d_2459_v42_ = _dafny.MultiSet([(d_2453_v38_)[default__.safeIndex(377, (d_2453_v38_).length(0))]])
            index402_ = default__.safeIndex(377, (d_2453_v38_).length(0))
            (d_2453_v38_)[index402_] = len(((((d_2458_v41_).cf17).set(default__.safeIndex((0) - (((d_2459_v42_)[p0] if (p0) in (d_2459_v42_) else p0)), len((d_2458_v41_).cf17)), d_2411_v0_)) + (d_2457_v40_)) + ((d_2457_v40_) + (d_2457_v40_)))
        d_2460_v43_: _dafny.Set
        d_2460_v43_ = _dafny.Set({False})
        d_2411_v0_ = ((d_2460_v43_) - (d_2460_v43_)).isdisjoint((d_2460_v43_) - (d_2460_v43_))
        if d_2411_v0_:
            d_2461_v44_: str
            d_2461_v44_ = _dafny.CodePoint('p')
            d_2462_v45_: D4
            d_2462_v45_ = D4_DC12(d_2461_v44_)
            pat_let_tv48_ = d_2461_v44_
            pat_let_tv49_ = d_2411_v0_
            pat_let_tv50_ = d_2461_v44_
            def iife180_(_pat_let59_0):
                def iife181_(d_2463_dt__update__tmp_h2_):
                    def iife182_(_pat_let60_0):
                        def iife183_(d_2464_dt__update_hcf13_h0_):
                            return D4_DC12(d_2464_dt__update_hcf13_h0_)
                        return iife183_(_pat_let60_0)
                    return iife182_((pat_let_tv48_ if pat_let_tv49_ else pat_let_tv50_))
                return iife181_(_pat_let59_0)
            source40_ = iife180_(d_2462_v45_)
            if source40_.is_DC13:
                d_2465___mcc_h0_ = source40_.cf14
                d_2466___mcc_h1_ = source40_.cf15
                d_2467___mcc_h2_ = source40_.cf16
                d_2468___mcc_h3_ = source40_.cf17
                d_2469_cf17_ = d_2468___mcc_h3_
                d_2470_cf16_ = d_2467___mcc_h2_
                d_2471_cf15_ = d_2466___mcc_h1_
                d_2472_cf14_ = d_2465___mcc_h0_
                d_2470_cf16_ = default__.fm3(d_2472_cf14_, d_2471_cf15_, globalState)
                d_2473_v46_: C2
                nw381_ = C2()
                nw381_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgj")), d_2411_v0_)
                d_2473_v46_ = nw381_
                d_2474_v47_: _dafny.Seq
                d_2474_v47_ = _dafny.SeqWithoutIsStrInference([(d_2470_cf16_) + (p0), -328])
                d_2474_v47_ = (d_2474_v47_) + (d_2474_v47_)
                r0 = d_2471_cf15_
            elif True:
                d_2475___mcc_h4_ = source40_.cf13
                d_2476_cf13_ = d_2475___mcc_h4_
                d_2477_v48_: _dafny.Array
                nw382_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_2477_v48_ = nw382_
                d_2478_v49_: C1
                nw383_ = C1()
                nw383_.ctor__(p0, p0, d_2477_v48_, p0, d_2411_v0_)
                d_2478_v49_ = nw383_
                d_2478_v49_ = d_2478_v49_
                (d_2478_v49_).f16 = (d_2478_v49_).fm7(globalState)
                d_2479_v50_: _dafny.Map
                d_2479_v50_ = _dafny.Map({d_2411_v0_: d_2411_v0_})
                d_2480_v51_: _dafny.MultiSet
                d_2480_v51_ = _dafny.MultiSet([d_2411_v0_])
                d_2479_v50_ = (d_2479_v50_).set(True, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2411_v0_]))) == (d_2480_v51_))
                d_2481_v52_: _dafny.MultiSet
                d_2481_v52_ = _dafny.MultiSet([d_2461_v44_, d_2461_v44_, d_2476_cf13_, d_2461_v44_])
                d_2481_v52_ = (d_2481_v52_ if d_2411_v0_ else _dafny.MultiSet([d_2476_cf13_]))
            d_2482_v53_: _dafny.Array
            nw384_ = _dafny.Array(_dafny.CodePoint('D'), 5)
            d_2482_v53_ = nw384_
            index403_ = default__.safeIndex(201, (d_2482_v53_).length(0))
            (d_2482_v53_)[index403_] = d_2461_v44_
            index404_ = default__.safeIndex(201, (d_2482_v53_).length(0))
            (d_2482_v53_)[index404_] = d_2461_v44_
            d_2483_v54_: int
            d_2483_v54_ = -581
            d_2483_v54_ = d_2483_v54_
            d_2484_v55_: _dafny.Seq
            d_2484_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbuouf"))
            d_2485_v56_: _dafny.MultiSet
            d_2485_v56_ = _dafny.MultiSet([d_2484_v55_, d_2484_v55_])
            d_2486_v57_: D11
            d_2486_v57_ = D11_DC25(((0) - (p0)) == (p0), default__.safeModuloInt((0) - ((0) - ((d_2485_v56_).cardinality)), (0) - (p0)))
            d_2486_v57_ = d_2486_v57_
            d_2487_v58_: _dafny.Array
            nw385_ = _dafny.Array(None, 4)
            nw385_[int(0)] = (0) - (d_2483_v54_)
            nw385_[int(1)] = p0
            nw385_[int(2)] = d_2483_v54_
            nw385_[int(3)] = d_2483_v54_
            d_2487_v58_ = nw385_
            d_2488_v59_: D16
            d_2488_v59_ = D16_DC32(d_2487_v58_)
            source41_ = d_2488_v59_
            if source41_.is_DC33:
                d_2489___mcc_h5_ = source41_.cf52
                d_2490___mcc_h6_ = source41_.cf53
                d_2491___mcc_h7_ = source41_.cf54
                d_2492___mcc_h8_ = source41_.cf55
                d_2493___mcc_h9_ = source41_.cf56
                d_2494_cf56_ = d_2493___mcc_h9_
                d_2495_cf55_ = d_2492___mcc_h8_
                d_2496_cf54_ = d_2491___mcc_h7_
                d_2497_cf53_ = d_2490___mcc_h6_
                d_2498_cf52_ = d_2489___mcc_h5_
                r0 = not(((d_2411_v0_ if d_2411_v0_ else True) if d_2498_cf52_ else d_2411_v0_))
                d_2499_v60_: D11
                d_2499_v60_ = D11_DC26(d_2411_v0_)
                d_2500_v61_: _dafny.Map
                d_2500_v61_ = _dafny.Map({d_2411_v0_: d_2411_v0_})
                d_2501_v62_: _dafny.Map
                d_2501_v62_ = _dafny.Map({d_2499_v60_: ((d_2500_v61_)[d_2411_v0_] if (d_2411_v0_) in (d_2500_v61_) else d_2498_cf52_)})
                d_2502_v63_: _dafny.Seq
                d_2502_v63_ = _dafny.SeqWithoutIsStrInference([d_2501_v62_])
                d_2503_v64_: D18
                d_2503_v64_ = D18_DC35(d_2502_v63_)
                d_2504_v65_: _dafny.Array
                nw386_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_2504_v65_ = nw386_
                d_2505_v66_: _dafny.Map
                d_2505_v66_ = _dafny.Map({d_2461_v44_: d_2411_v0_})
                d_2506_v67_: _dafny.Array
                nw387_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_2506_v67_ = nw387_
                d_2507_v68_: T1
                nw388_ = C6()
                nw388_.ctor__(d_2411_v0_, ((d_2505_v66_)[_dafny.CodePoint('g')] if (_dafny.CodePoint('g')) in (d_2505_v66_) else d_2411_v0_), d_2506_v67_, d_2495_cf55_)
                d_2507_v68_ = nw388_
                d_2508_v69_: _dafny.Array
                nw389_ = _dafny.Array(None, 27)
                nw389_[int(0)] = d_2507_v68_
                nw389_[int(1)] = d_2507_v68_
                nw389_[int(2)] = d_2507_v68_
                nw389_[int(3)] = d_2507_v68_
                nw389_[int(4)] = d_2507_v68_
                nw389_[int(5)] = d_2507_v68_
                nw389_[int(6)] = d_2507_v68_
                nw389_[int(7)] = d_2507_v68_
                nw389_[int(8)] = d_2507_v68_
                nw389_[int(9)] = d_2507_v68_
                nw389_[int(10)] = d_2507_v68_
                nw389_[int(11)] = d_2507_v68_
                nw389_[int(12)] = d_2507_v68_
                nw389_[int(13)] = d_2507_v68_
                nw389_[int(14)] = d_2507_v68_
                nw389_[int(15)] = d_2507_v68_
                nw389_[int(16)] = d_2507_v68_
                nw389_[int(17)] = d_2507_v68_
                nw389_[int(18)] = d_2507_v68_
                nw389_[int(19)] = d_2507_v68_
                nw389_[int(20)] = d_2507_v68_
                nw389_[int(21)] = d_2507_v68_
                nw389_[int(22)] = d_2507_v68_
                nw389_[int(23)] = d_2507_v68_
                nw389_[int(24)] = d_2507_v68_
                nw389_[int(25)] = d_2507_v68_
                nw389_[int(26)] = d_2507_v68_
                d_2508_v69_ = nw389_
                d_2509_v70_: _dafny.MultiSet
                d_2509_v70_ = _dafny.MultiSet([d_2508_v69_, d_2508_v69_])
                d_2510_v71_: _dafny.Map
                d_2510_v71_ = _dafny.Map({(d_2509_v70_ if d_2498_cf52_ else d_2509_v70_): d_2461_v44_})
                index405_ = default__.safeIndex(201, (d_2482_v53_).length(0))
                rhs380_ = default__.fm1(len(d_2484_v55_), d_2411_v0_, globalState)
                rhs381_ = d_2498_cf52_
                rhs382_ = d_2503_v64_
                rhs383_ = d_2504_v65_
                rhs384_ = ((d_2510_v71_)[_dafny.MultiSet([d_2508_v69_])] if (_dafny.MultiSet([d_2508_v69_])) in (d_2510_v71_) else d_2461_v44_)
                lhs241_ = d_2482_v53_
                lhs242_ = default__.safeIndex(201, (d_2482_v53_).length(0))
                lhs241_[lhs242_] = rhs380_
                d_2498_cf52_ = rhs381_
                d_2503_v64_ = rhs382_
                d_2504_v65_ = rhs383_
                d_2461_v44_ = rhs384_
                d_2511_v72_: _dafny.Seq
                d_2511_v72_ = _dafny.SeqWithoutIsStrInference([d_2494_cf56_])
                d_2512_v73_: _dafny.Map
                d_2512_v73_ = _dafny.Map({(0) - (d_2495_cf55_): -139})
                d_2513_v74_: _dafny.Seq
                d_2513_v74_ = _dafny.SeqWithoutIsStrInference([d_2498_cf52_, (d_2507_v68_).f4])
                d_2514_v75_: D9
                d_2514_v75_ = D9_DC19(d_2484_v55_)
                pat_let_tv51_ = d_2484_v55_
                d_2515_v76_: _dafny.Array
                nw390_ = _dafny.Array(None, 29)
                nw390_[int(0)] = d_2498_cf52_
                nw390_[int(1)] = d_2411_v0_
                nw390_[int(2)] = d_2498_cf52_
                nw390_[int(3)] = True
                nw390_[int(4)] = not(d_2498_cf52_)
                nw390_[int(5)] = (d_2507_v68_).f4
                nw390_[int(6)] = (d_2498_cf52_) and (False)
                nw390_[int(7)] = (len(_dafny.SeqWithoutIsStrInference([(d_2482_v53_)[default__.safeIndex(201, (d_2482_v53_).length(0))] for d_2516_i2_ in range(default__.abs(614))]))) < (d_2483_v54_)
                nw390_[int(8)] = d_2498_cf52_
                nw390_[int(9)] = d_2498_cf52_
                nw390_[int(10)] = d_2411_v0_
                nw390_[int(11)] = (d_2483_v54_) > ((0) - (p0))
                nw390_[int(12)] = (not(d_2411_v0_)) in ((d_2511_v72_)[default__.safeIndex(((d_2512_v73_)[(d_2507_v68_).f6] if ((d_2507_v68_).f6) in (d_2512_v73_) else d_2483_v54_), len(d_2511_v72_))])
                nw390_[int(13)] = d_2411_v0_
                nw390_[int(14)] = default__.fm0(d_2461_v44_, globalState)
                nw390_[int(15)] = (d_2513_v74_)[default__.safeIndex(p0, len(d_2513_v74_))]
                nw390_[int(16)] = d_2411_v0_
                nw390_[int(17)] = False
                nw390_[int(18)] = d_2498_cf52_
                nw390_[int(19)] = (d_2507_v68_).f4
                nw390_[int(20)] = d_2411_v0_
                nw390_[int(21)] = d_2498_cf52_
                nw390_[int(22)] = d_2498_cf52_
                nw390_[int(23)] = d_2411_v0_
                nw390_[int(24)] = (d_2507_v68_).f4
                def iife184_(_pat_let61_0):
                    def iife185_(d_2518_dt__update__tmp_h3_):
                        def iife186_(_pat_let62_0):
                            def iife187_(d_2519_dt__update_hcf25_h0_):
                                return D9_DC19(d_2519_dt__update_hcf25_h0_)
                            return iife187_(_pat_let62_0)
                        return iife186_(pat_let_tv51_)
                    return iife185_(_pat_let61_0)
                nw390_[int(25)] = (_dafny.SeqWithoutIsStrInference([(d_2482_v53_)[default__.safeIndex(201, (d_2482_v53_).length(0))] for d_2517_i3_ in range(default__.abs(895))])) <= ((iife184_(d_2514_v75_)).cf25)
                nw390_[int(26)] = False
                nw390_[int(27)] = (d_2507_v68_).f4
                nw390_[int(28)] = not ((d_2507_v68_).f4) or (d_2498_cf52_)
                d_2515_v76_ = nw390_
                index406_ = default__.safeIndex(172, (d_2515_v76_).length(0))
                (d_2515_v76_)[index406_] = d_2498_cf52_
                d_2520_v77_: _dafny.Map
                d_2520_v77_ = _dafny.Map({d_2512_v73_: (d_2507_v68_).f4})
                d_2521_v78_: _dafny.Map
                d_2521_v78_ = _dafny.Map({(d_2507_v68_).f6: d_2487_v58_})
                d_2522_v79_: _dafny.Map
                d_2522_v79_ = _dafny.Map({default__.fm3(len(d_2520_v77_), d_2411_v0_, globalState): ((d_2521_v78_)[d_2497_cf53_] if (d_2497_cf53_) in (d_2521_v78_) else d_2487_v58_)})
                d_2523_v80_: D20
                d_2523_v80_ = D20_DC41(d_2483_v54_)
                d_2524_v81_: _dafny.MultiSet
                d_2524_v81_ = _dafny.MultiSet([(d_2507_v68_).f6])
                d_2525_v82_: _dafny.Seq
                d_2525_v82_ = _dafny.SeqWithoutIsStrInference([(d_2507_v68_).f6, d_2495_cf55_])
                d_2526_v83_: _dafny.MultiSet
                d_2526_v83_ = _dafny.MultiSet([(d_2524_v81_).cardinality, len(d_2525_v82_), d_2483_v54_, d_2497_cf53_, 943])
                d_2527_v84_: C1
                nw391_ = C1()
                nw391_.ctor__(925, p0, d_2506_v67_, d_2497_cf53_, not(d_2411_v0_))
                d_2527_v84_ = nw391_
                d_2528_v85_: _dafny.Map
                d_2528_v85_ = _dafny.Map({625: _dafny.SeqWithoutIsStrInference([((d_2524_v81_).set(d_2495_cf55_, default__.abs(len(_dafny.Map({d_2527_v84_: (d_2496_cf54_).cardinality}))))).cardinality])})
                index407_ = default__.safeIndex(172, (d_2515_v76_).length(0))
                rhs385_ = (d_2507_v68_).fm8((d_2523_v80_).cf67, d_2498_cf52_, _dafny.SeqWithoutIsStrInference([(d_2482_v53_)[default__.safeIndex(201, (d_2482_v53_).length(0))] for d_2529_i4_ in range(default__.abs(69))]), ((d_2528_v85_)[852] if (852) in (d_2528_v85_) else d_2525_v82_), globalState)
                rhs386_ = (d_2507_v68_).f4
                rhs387_ = not((d_2507_v68_).f4)
                rhs388_ = (d_2522_v79_) | (d_2522_v79_)
                lhs243_ = d_2515_v76_
                lhs244_ = default__.safeIndex(172, (d_2515_v76_).length(0))
                lhs243_[lhs244_] = rhs385_
                d_2498_cf52_ = rhs386_
                r0 = rhs387_
                d_2522_v79_ = rhs388_
                d_2530_v86_: _dafny.Array
                nw392_ = _dafny.Array(_dafny.Map({}), 29)
                d_2530_v86_ = nw392_
                d_2531_v87_: _dafny.Set
                d_2531_v87_ = _dafny.Set({d_2515_v76_, d_2515_v76_})
                d_2532_v88_: D22
                d_2532_v88_ = D22_DC45(d_2530_v86_)
                d_2533_v89_: _dafny.Map
                d_2533_v89_ = _dafny.Map({d_2531_v87_: (d_2532_v88_).cf73})
                d_2530_v86_ = ((d_2533_v89_)[d_2531_v87_] if (d_2531_v87_) in (d_2533_v89_) else d_2530_v86_)
            elif True:
                d_2534___mcc_h10_ = source41_.cf51
                d_2535_cf51_ = d_2534___mcc_h10_
                d_2536_v90_: _dafny.Array
                nw393_ = _dafny.Array(False, 22)
                d_2536_v90_ = nw393_
                d_2537_v91_: _dafny.Seq
                d_2537_v91_ = _dafny.SeqWithoutIsStrInference([d_2483_v54_, p0, p0, 13])
                d_2538_v92_: _dafny.Map
                d_2538_v92_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2536_v90_, d_2536_v90_, d_2536_v90_, d_2536_v90_, d_2536_v90_]): (_dafny.MultiSet(d_2537_v91_)).cardinality})
                d_2539_v93_: _dafny.Map
                d_2539_v93_ = _dafny.Map({((d_2538_v92_)[_dafny.SeqWithoutIsStrInference([d_2536_v90_, d_2536_v90_, d_2536_v90_, d_2536_v90_])] if (_dafny.SeqWithoutIsStrInference([d_2536_v90_, d_2536_v90_, d_2536_v90_, d_2536_v90_])) in (d_2538_v92_) else d_2483_v54_): default__.fm40(False, True, p0, globalState)})
                d_2539_v93_ = (d_2539_v93_).set(192, default__.fm19(d_2411_v0_, p0, globalState))
                index408_ = default__.safeIndex(700, (d_2536_v90_).length(0))
                (d_2536_v90_)[index408_] = (self).fm11(globalState)
                index409_ = default__.safeIndex(700, (d_2536_v90_).length(0))
                (d_2536_v90_)[index409_] = d_2411_v0_
                d_2540_v94_: _dafny.Array
                def lambda114_(d_2541_v44_):
                    def lambda115_(d_2542_i5_):
                        return _dafny.SeqWithoutIsStrInference([d_2541_v44_ for d_2543_i6_ in range(default__.abs(370))])

                    return lambda115_

                init72_ = lambda114_(d_2461_v44_)
                nw394_ = _dafny.Array(None, 5)
                for i0_72_ in range(nw394_.length(0)):
                    nw394_[i0_72_] = init72_(i0_72_)
                d_2540_v94_ = nw394_
                d_2544_v95_: _dafny.Array
                def lambda116_(d_2545_p0_, d_2546_v54_, d_2547_v90_):
                    def lambda117_(d_2548_i7_):
                        return _dafny.Map({len(_dafny.Map({d_2545_p0_: (0) - (d_2546_v54_)})): (d_2547_v90_)[default__.safeIndex(700, (d_2547_v90_).length(0))]})

                    return lambda117_

                init73_ = lambda116_(p0, d_2483_v54_, d_2536_v90_)
                nw395_ = _dafny.Array(None, 24)
                for i0_73_ in range(nw395_.length(0)):
                    nw395_[i0_73_] = init73_(i0_73_)
                d_2544_v95_ = nw395_
                rhs389_ = d_2540_v94_
                rhs390_ = d_2544_v95_
                d_2540_v94_ = rhs389_
                d_2544_v95_ = rhs390_
                d_2549_v96_: _dafny.Seq
                d_2549_v96_ = _dafny.SeqWithoutIsStrInference([d_2484_v55_, _dafny.SeqWithoutIsStrInference([d_2461_v44_ for d_2550_i8_ in range(default__.abs(525))]), d_2484_v55_, d_2484_v55_])
                d_2551_v97_: D12
                d_2551_v97_ = D12_DC27(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2411_v0_]) for d_2552_i9_ in range(default__.abs(195))]))
                d_2553_v98_: _dafny.Set
                d_2553_v98_ = _dafny.Set({d_2483_v54_, p0})
                d_2554_v99_: _dafny.Map
                d_2554_v99_ = _dafny.Map({not((d_2536_v90_)[default__.safeIndex(700, (d_2536_v90_).length(0))]): d_2553_v98_})
                index410_ = default__.safeIndex(700, (d_2536_v90_).length(0))
                rhs391_ = (p0) >= (p0)
                rhs392_ = default__.fm39(((d_2484_v55_).set(default__.safeIndex(d_2483_v54_, len(d_2484_v55_)), (d_2482_v53_)[default__.safeIndex(201, (d_2482_v53_).length(0))])) + ((d_2549_v96_)[default__.safeIndex(p0, len(d_2549_v96_))]), d_2551_v97_, d_2411_v0_, globalState)
                rhs393_ = (self).fm11(globalState)
                rhs394_ = (len(((d_2554_v99_)[False] if (False) in (d_2554_v99_) else d_2553_v98_))) + (d_2483_v54_)
                rhs395_ = d_2483_v54_
                lhs245_ = d_2536_v90_
                lhs246_ = default__.safeIndex(700, (d_2536_v90_).length(0))
                r1 = rhs391_
                d_2537_v91_ = rhs392_
                lhs245_[lhs246_] = rhs393_
                d_2483_v54_ = rhs394_
                d_2483_v54_ = rhs395_
        elif True:
            d_2555_v100_: _dafny.MultiSet
            d_2555_v100_ = _dafny.MultiSet([d_2411_v0_])
            r0 = ((d_2555_v100_) - (d_2555_v100_)).isdisjoint(_dafny.MultiSet([d_2411_v0_, d_2411_v0_, d_2411_v0_]))
            d_2556_v101_: _dafny.Array
            def lambda118_(d_2557_v0_):
                def lambda119_(d_2558_i10_):
                    return (d_2557_v0_) == (d_2557_v0_)

                return lambda119_

            init74_ = lambda118_(d_2411_v0_)
            nw396_ = _dafny.Array(None, 24)
            for i0_74_ in range(nw396_.length(0)):
                nw396_[i0_74_] = init74_(i0_74_)
            d_2556_v101_ = nw396_
            index411_ = default__.safeIndex(935, (d_2556_v101_).length(0))
            (d_2556_v101_)[index411_] = d_2411_v0_
            d_2559_v102_: _dafny.Map
            d_2559_v102_ = _dafny.Map({d_2411_v0_: p0})
            index412_ = default__.safeIndex(935, (d_2556_v101_).length(0))
            (d_2556_v101_)[index412_] = (((d_2559_v102_)[d_2411_v0_] if (d_2411_v0_) in (d_2559_v102_) else p0)) == (default__.fm3(len(d_2460_v43_), d_2411_v0_, globalState))
            d_2560_v103_: _dafny.Array
            def lambda120_(d_2561_p0_):
                def lambda121_(d_2562_i11_):
                    return (d_2562_i11_) - (d_2561_p0_)

                return lambda121_

            init75_ = lambda120_(p0)
            nw397_ = _dafny.Array(None, 11)
            for i0_75_ in range(nw397_.length(0)):
                nw397_[i0_75_] = init75_(i0_75_)
            d_2560_v103_ = nw397_
            index413_ = default__.safeIndex(247, (d_2560_v103_).length(0))
            (d_2560_v103_)[index413_] = (0) - (p0)
            index414_ = default__.safeIndex(247, (d_2560_v103_).length(0))
            (d_2560_v103_)[index414_] = p0
            d_2563_v104_: _dafny.Array
            nw398_ = _dafny.Array(_dafny.Map({}), 20)
            d_2563_v104_ = nw398_
            d_2564_v105_: D22
            d_2564_v105_ = D22_DC45(d_2563_v104_)
            pat_let_tv52_ = d_2564_v105_
            d_2565_v106_: _dafny.Array
            nw399_ = _dafny.Array(None, 4)
            nw399_[int(0)] = d_2564_v105_
            nw399_[int(1)] = d_2564_v105_
            def iife188_(_pat_let63_0):
                def iife189_(d_2566_dt__update__tmp_h4_):
                    def iife190_(_pat_let64_0):
                        def iife191_(d_2567_dt__update_hcf73_h0_):
                            return D22_DC45(d_2567_dt__update_hcf73_h0_)
                        return iife191_(_pat_let64_0)
                    return iife190_((pat_let_tv52_).cf73)
                return iife189_(_pat_let63_0)
            nw399_[int(2)] = iife188_(d_2564_v105_)
            nw399_[int(3)] = d_2564_v105_
            d_2565_v106_ = nw399_
            index415_ = default__.safeIndex(442, (d_2565_v106_).length(0))
            (d_2565_v106_)[index415_] = d_2564_v105_
            index416_ = default__.safeIndex(442, (d_2565_v106_).length(0))
            (d_2565_v106_)[index416_] = d_2564_v105_
            d_2568_v107_: C0
            nw400_ = C0()
            nw400_.ctor__((d_2560_v103_)[default__.safeIndex(247, (d_2560_v103_).length(0))])
            d_2568_v107_ = nw400_
            d_2569_v108_: _dafny.Seq
            d_2569_v108_ = _dafny.SeqWithoutIsStrInference([d_2568_v107_, d_2568_v107_])
            d_2570_v109_: _dafny.Map
            d_2570_v109_ = _dafny.Map({len(_dafny.Map({d_2411_v0_: p0})): d_2568_v107_})
            d_2571_v110_: D23
            d_2571_v110_ = D23_DC48(d_2568_v107_)
            d_2572_v111_: _dafny.Array
            nw401_ = _dafny.Array(None, 24)
            nw401_[int(0)] = d_2568_v107_
            nw401_[int(1)] = d_2568_v107_
            nw401_[int(2)] = d_2568_v107_
            nw401_[int(3)] = d_2568_v107_
            nw401_[int(4)] = d_2568_v107_
            nw401_[int(5)] = d_2568_v107_
            nw401_[int(6)] = d_2568_v107_
            nw401_[int(7)] = d_2568_v107_
            nw401_[int(8)] = d_2568_v107_
            nw401_[int(9)] = d_2568_v107_
            nw401_[int(10)] = d_2568_v107_
            nw401_[int(11)] = d_2568_v107_
            nw401_[int(12)] = (d_2569_v108_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_2568_v107_).f12, p0])), len(d_2569_v108_))]
            nw401_[int(13)] = d_2568_v107_
            nw401_[int(14)] = d_2568_v107_
            nw401_[int(15)] = ((d_2570_v109_)[p0] if (p0) in (d_2570_v109_) else d_2568_v107_)
            nw401_[int(16)] = d_2568_v107_
            nw401_[int(17)] = (d_2571_v110_).cf82
            nw401_[int(18)] = d_2568_v107_
            nw401_[int(19)] = d_2568_v107_
            nw401_[int(20)] = d_2568_v107_
            nw401_[int(21)] = d_2568_v107_
            nw401_[int(22)] = d_2568_v107_
            nw401_[int(23)] = (d_2571_v110_).cf82
            d_2572_v111_ = nw401_
            d_2572_v111_ = d_2572_v111_
        d_2573_v112_: _dafny.Array
        def lambda122_(d_2574_p0_):
            def lambda123_(d_2575_i13_):
                return (d_2575_i13_) * (d_2574_p0_)

            return lambda123_

        init76_ = lambda122_(p0)
        nw402_ = _dafny.Array(None, 29)
        for i0_76_ in range(nw402_.length(0)):
            nw402_[i0_76_] = init76_(i0_76_)
        d_2573_v112_ = nw402_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_2573_v112_).length(0)):
            d_2576_i12_: int = guard_loop_7_
            if (True) and (((0) <= (d_2576_i12_)) and ((d_2576_i12_) < ((d_2573_v112_).length(0)))):
                (d_2573_v112_)[(d_2576_i12_)] = default__.safeModuloInt(d_2576_i12_, (p0) + (p0))
        d_2577_v113_: _dafny.Array
        nw403_ = _dafny.Array(None, 19)
        d_2577_v113_ = nw403_
        d_2578_v114_: _dafny.Array
        def lambda124_(d_2579_v0_):
            def lambda125_(d_2580_i14_):
                return d_2579_v0_

            return lambda125_

        init77_ = lambda124_(d_2411_v0_)
        nw404_ = _dafny.Array(None, 26)
        for i0_77_ in range(nw404_.length(0)):
            nw404_[i0_77_] = init77_(i0_77_)
        d_2578_v114_ = nw404_
        d_2581_v115_: _dafny.Array
        nw405_ = _dafny.Array(None, 23)
        nw405_[int(0)] = d_2578_v114_
        nw405_[int(1)] = d_2578_v114_
        nw405_[int(2)] = d_2578_v114_
        nw405_[int(3)] = d_2578_v114_
        nw405_[int(4)] = d_2578_v114_
        nw405_[int(5)] = d_2578_v114_
        nw405_[int(6)] = d_2578_v114_
        nw405_[int(7)] = d_2578_v114_
        nw405_[int(8)] = d_2578_v114_
        nw405_[int(9)] = d_2578_v114_
        nw405_[int(10)] = d_2578_v114_
        nw405_[int(11)] = d_2578_v114_
        nw405_[int(12)] = d_2578_v114_
        nw405_[int(13)] = d_2578_v114_
        nw405_[int(14)] = d_2578_v114_
        nw405_[int(15)] = d_2578_v114_
        nw405_[int(16)] = d_2578_v114_
        nw405_[int(17)] = d_2578_v114_
        nw405_[int(18)] = d_2578_v114_
        nw405_[int(19)] = d_2578_v114_
        nw405_[int(20)] = d_2578_v114_
        nw405_[int(21)] = d_2578_v114_
        nw405_[int(22)] = d_2578_v114_
        d_2581_v115_ = nw405_
        d_2582_v116_: _dafny.Seq
        d_2582_v116_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "clpbitvi"))
        d_2583_v117_: _dafny.Seq
        d_2583_v117_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvtigi")), d_2582_v116_])
        d_2584_v118_: C3
        nw406_ = C3()
        nw406_.ctor__(d_2581_v115_, len(d_2583_v117_), d_2411_v0_)
        d_2584_v118_ = nw406_
        index417_ = default__.safeIndex(124, (d_2577_v113_).length(0))
        (d_2577_v113_)[index417_] = d_2584_v118_
        d_2585_v119_: _dafny.Map
        d_2585_v119_ = _dafny.Map({d_2411_v0_: p0})
        d_2586_v120_: C0
        nw407_ = C0()
        nw407_.ctor__((0) - (len(d_2585_v119_)))
        d_2586_v120_ = nw407_
        index418_ = default__.safeIndex(124, (d_2577_v113_).length(0))
        rhs396_ = d_2411_v0_
        rhs397_ = d_2584_v118_
        rhs398_ = d_2586_v120_
        lhs247_ = d_2577_v113_
        lhs248_ = default__.safeIndex(124, (d_2577_v113_).length(0))
        r1 = rhs396_
        lhs247_[lhs248_] = rhs397_
        d_2586_v120_ = rhs398_
        d_2587_v121_: int
        d_2587_v121_ = -191
        d_2587_v121_ = ((0) - ((d_2586_v120_).f12)) - ((d_2586_v120_).f12)
        d_2588_v122_: D22
        d_2588_v122_ = D22_DC47(len(d_2582_v116_), not(not(d_2411_v0_)), d_2411_v0_, (d_2586_v120_).f12, (0) - (p0))
        r0 = (d_2588_v122_).cf79
        r1 = d_2411_v0_
        return r0, r1


class C9(T1, T0):
    def  __init__(self):
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self._f6: int = int(0)
        self._f4: bool = False
        self.f7: int = int(0)
        self._f8: _dafny.Set = _dafny.Set({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f5(self):
        return self._f5
    @f5.setter
    def f5(self, value):
        self._f5 = value
    @property
    def f6(self):
        return self._f6
    @property
    def f4(self):
        return self._f4
    def ctor__(self, f7, f8, f5, f6, f4):
        (self).f7 = f7
        (self)._f8 = f8
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f4 = f4

    def fm8(self, p0, p1, p2, p3, globalState):
        return (self).f4

    def fm9(self, p0, p1, p2, globalState):
        source42_ = (D1_DC4(len(_dafny.SeqWithoutIsStrInference([self.f7])), 912) if (self).f4 else D1_DC4((self).f6, 392))
        if source42_.is_DC4:
            d_2589___mcc_h0_ = source42_.cf3
            d_2590___mcc_h1_ = source42_.cf4
            d_2591_cf4_ = d_2590___mcc_h1_
            d_2592_cf3_ = d_2589___mcc_h0_
            return (0) - (d_2592_cf3_)
        elif True:
            d_2593___mcc_h2_ = source42_.cf2
            d_2594_cf2_ = d_2593___mcc_h2_
            return (self).f6

    def fm6(self, p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f4]))).cardinality, (self).f6, self.f7, -280])) | ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fku")))])).intersection(_dafny.MultiSet([len(_dafny.Map({self.f7: (self).f4}))])))

    def fm7(self, globalState):
        return self.f7

    def m1(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_2595_v0_: _dafny.Array
        def lambda126_(d_2596_i1_):
            return (self).f4

        init78_ = lambda126_
        nw408_ = _dafny.Array(None, 18)
        for i0_78_ in range(nw408_.length(0)):
            nw408_[i0_78_] = init78_(i0_78_)
        d_2595_v0_ = nw408_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_2595_v0_).length(0)):
            d_2597_i0_: int = guard_loop_8_
            if (True) and (((0) <= (d_2597_i0_)) and ((d_2597_i0_) < ((d_2595_v0_).length(0)))):
                (d_2595_v0_)[(d_2597_i0_)] = ((self).f6) >= (p0)
        hi14_ = (0) - (default__.safeDivisionInt(self.f7, self.f7))
        for d_2598_i2_ in range(default__.safeModuloInt(self.f7, self.f7), hi14_):
            d_2599_v1_: D1
            d_2599_v1_ = D1_DC4((self).f6, p0)
            d_2600_v2_: _dafny.Map
            d_2600_v2_ = _dafny.Map({not((self).f4): (d_2599_v1_ if False else d_2599_v1_)})
            d_2601_v3_: _dafny.Seq
            d_2601_v3_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')])
            d_2602_v4_: _dafny.Map
            d_2602_v4_ = _dafny.Map({(self).f6: d_2601_v3_})
            d_2603_v5_: str
            d_2603_v5_ = _dafny.CodePoint('h')
            d_2604_v6_: _dafny.Map
            d_2604_v6_ = _dafny.Map({(self).f4: len((d_2601_v3_).set(default__.safeIndex(d_2598_i2_, len(d_2601_v3_)), d_2603_v5_))})
            d_2600_v2_ = (d_2600_v2_).set((self).f4, default__.fm10((self).f4, False, d_2602_v4_, len(d_2604_v6_), globalState))
            d_2605_v7_: bool
            d_2605_v7_ = True
            d_2605_v7_ = (self).f4
            d_2606_v8_: C0
            nw409_ = C0()
            nw409_.ctor__((d_2598_i2_) + ((0) - ((self).f6)))
            d_2606_v8_ = nw409_
            if d_2605_v7_:
                d_2607_v9_: _dafny.MultiSet
                d_2607_v9_ = _dafny.MultiSet([(d_2606_v8_).f12, (_dafny.MultiSet([len(d_2604_v6_), 886])).cardinality])
                d_2608_v10_: _dafny.Map
                d_2608_v10_ = _dafny.Map({(d_2607_v9_).intersection(d_2607_v9_): 451})
                d_2609_v11_: D0
                d_2609_v11_ = D0_DC0(d_2598_i2_)
                d_2610_v12_: _dafny.Map
                d_2610_v12_ = _dafny.Map({(self).f6: d_2603_v5_})
                d_2608_v10_ = (d_2608_v10_).set((_dafny.MultiSet([p0, (0) - ((d_2606_v8_).f12), self.f7])) | ((self).fm6(d_2609_v11_, d_2609_v11_, d_2610_v12_, 251, globalState)), self.f7)
                rhs399_ = (self).f4
                rhs400_ = len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2598_i2_])]))
                lhs249_ = self
                d_2605_v7_ = rhs399_
                lhs249_.f7 = rhs400_
                d_2605_v7_ = d_2605_v7_
                d_2611_v13_: _dafny.Array
                def lambda127_(d_2612_v3_):
                    def lambda128_(d_2613_i3_):
                        return d_2612_v3_

                    return lambda128_

                init79_ = lambda127_(d_2601_v3_)
                nw410_ = _dafny.Array(None, 14)
                for i0_79_ in range(nw410_.length(0)):
                    nw410_[i0_79_] = init79_(i0_79_)
                d_2611_v13_ = nw410_
                index419_ = default__.safeIndex(656, (d_2611_v13_).length(0))
                (d_2611_v13_)[index419_] = default__.fm16(d_2605_v7_, (self).f4, not(d_2605_v7_), D2_DC6(), globalState)
                d_2614_v14_: _dafny.MultiSet
                d_2614_v14_ = _dafny.MultiSet([d_2609_v11_, d_2609_v11_])
                d_2615_v15_: _dafny.Array
                nw411_ = _dafny.Array(None, 2)
                nw411_[int(0)] = d_2614_v14_
                nw411_[int(1)] = d_2614_v14_
                d_2615_v15_ = nw411_
                index420_ = default__.safeIndex(913, (d_2615_v15_).length(0))
                (d_2615_v15_)[index420_] = _dafny.MultiSet([d_2609_v11_])
                d_2616_v16_: _dafny.Seq
                def iife192_(_pat_let65_0):
                    def iife193_(d_2617_dt__update__tmp_h0_):
                        def iife194_(_pat_let66_0):
                            def iife195_(d_2618_dt__update_hcf0_h0_):
                                return D0_DC0(d_2618_dt__update_hcf0_h0_)
                            return iife195_(_pat_let66_0)
                        return iife194_((0) - (d_2598_i2_))
                    return iife193_(_pat_let65_0)
                d_2616_v16_ = _dafny.SeqWithoutIsStrInference([iife192_(d_2609_v11_)])
                index421_ = default__.safeIndex(656, (d_2611_v13_).length(0))
                index422_ = default__.safeIndex(913, (d_2615_v15_).length(0))
                rhs401_ = d_2601_v3_
                rhs402_ = (d_2614_v14_) - (_dafny.MultiSet((d_2616_v16_) + (d_2616_v16_)))
                lhs250_ = d_2611_v13_
                lhs251_ = default__.safeIndex(656, (d_2611_v13_).length(0))
                lhs252_ = d_2615_v15_
                lhs253_ = default__.safeIndex(913, (d_2615_v15_).length(0))
                lhs250_[lhs251_] = rhs401_
                lhs252_[lhs253_] = rhs402_
                d_2619_v17_: D12
                d_2619_v17_ = D12_DC28(p0, self.f7, (self).f4, True, p0)
                d_2620_v18_: _dafny.Seq
                d_2620_v18_ = _dafny.SeqWithoutIsStrInference([(d_2619_v17_).cf46, (self).f4, (d_2605_v7_) == (d_2605_v7_)])
                d_2621_v19_: _dafny.Seq
                d_2621_v19_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                d_2620_v18_ = ((d_2620_v18_).set(default__.safeIndex(d_2598_i2_, len(d_2620_v18_)), (d_2620_v18_)[default__.safeIndex(len(d_2621_v19_), len(d_2620_v18_))])).set(default__.safeIndex(default__.safeModuloInt((self).f6, self.f7), len((d_2620_v18_).set(default__.safeIndex(d_2598_i2_, len(d_2620_v18_)), (d_2620_v18_)[default__.safeIndex(len(d_2621_v19_), len(d_2620_v18_))]))), not(True))
            elif True:
                d_2622_v20_: _dafny.Map
                d_2622_v20_ = _dafny.Map({d_2601_v3_: default__.fm1(self.f7, d_2605_v7_, globalState)})
                d_2622_v20_ = (d_2622_v20_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrkp")), d_2603_v5_)
                d_2623_v21_: C6
                nw412_ = C6()
                nw412_.ctor__((self).f4, (self).f4, (self.f5 if d_2605_v7_ else self.f5), (d_2606_v8_).f12)
                d_2623_v21_ = nw412_
                d_2624_v22_: _dafny.Set
                d_2624_v22_ = _dafny.Set({not(d_2605_v7_)})
                d_2605_v7_ = (d_2624_v22_) != (default__.fm46(_dafny.CodePoint('t'), globalState))
                d_2625_v23_: D3
                d_2625_v23_ = D3_DC9(True, p0, (self).f4)
                d_2626_v24_: _dafny.Map
                d_2626_v24_ = _dafny.Map({(0) - ((self).f6): (d_2625_v23_).cf9})
                d_2627_v25_: _dafny.Seq
                d_2627_v25_ = _dafny.SeqWithoutIsStrInference([d_2605_v7_, (self).f4])
                d_2628_v26_: _dafny.Seq
                d_2628_v26_ = _dafny.SeqWithoutIsStrInference([len(d_2627_v25_), (self).f6])
                d_2629_v27_: D6
                d_2629_v27_ = D6_DC16(False, (d_2623_v21_).f11, d_2628_v26_)
                d_2630_v28_: _dafny.Set
                d_2630_v28_ = _dafny.Set({d_2629_v27_, d_2629_v27_, d_2629_v27_})
                rhs403_ = (default__.fm42(d_2603_v5_, d_2630_v28_, globalState)) | (d_2626_v24_)
                rhs404_ = (d_2598_i2_) < ((d_2606_v8_).f12)
                rhs405_ = ((d_2606_v8_).f12) != ((self).fm7(globalState))
                rhs406_ = (p0) - (345)
                rhs407_ = p0
                lhs254_ = self
                lhs255_ = self
                d_2626_v24_ = rhs403_
                d_2605_v7_ = rhs404_
                d_2605_v7_ = rhs405_
                lhs254_.f7 = rhs406_
                lhs255_.f7 = rhs407_
                d_2605_v7_ = ((d_2606_v8_).f12) != (((self).f6) - (self.f7))
        d_2631_i4_: int
        d_2631_i4_ = 0
        with _dafny.label("13"):
            while (self).f4:
                with _dafny.c_label("13"):
                    if (d_2631_i4_) >= (100):
                        raise _dafny.Break("13")
                    d_2631_i4_ = (d_2631_i4_) + (1)
                    (self).f7 = (0) - ((self).f6)
                    d_2632_v29_: _dafny.Seq
                    d_2632_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnnryc"))
                    d_2633_v30_: T0
                    nw413_ = C2()
                    nw413_.ctor__(d_2632_v29_, (self).f4)
                    d_2633_v30_ = nw413_
                    d_2634_v31_: _dafny.Seq
                    d_2634_v31_ = _dafny.SeqWithoutIsStrInference([d_2633_v30_, d_2633_v30_, d_2633_v30_])
                    d_2635_v32_: T0
                    d_2635_v32_ = d_2633_v30_
                    d_2636_v33_: _dafny.Array
                    nw414_ = _dafny.Array(None, 18)
                    nw414_[int(0)] = d_2633_v30_
                    nw414_[int(1)] = (d_2634_v31_)[default__.safeIndex(p0, len(d_2634_v31_))]
                    nw414_[int(2)] = (d_2634_v31_)[default__.safeIndex(len(d_2632_v29_), len(d_2634_v31_))]
                    nw414_[int(3)] = (d_2633_v30_ if (self).f4 else d_2633_v30_)
                    nw414_[int(4)] = (d_2635_v32_)
                    nw414_[int(5)] = d_2633_v30_
                    nw414_[int(6)] = (d_2633_v30_ if (self).f4 else d_2633_v30_)
                    nw414_[int(7)] = d_2633_v30_
                    nw414_[int(8)] = d_2633_v30_
                    nw414_[int(9)] = d_2633_v30_
                    nw414_[int(10)] = d_2633_v30_
                    nw414_[int(11)] = d_2633_v30_
                    nw414_[int(12)] = d_2633_v30_
                    nw414_[int(13)] = d_2633_v30_
                    nw414_[int(14)] = d_2633_v30_
                    nw414_[int(15)] = d_2633_v30_
                    nw414_[int(16)] = d_2633_v30_
                    nw414_[int(17)] = d_2633_v30_
                    d_2636_v33_ = nw414_
                    index423_ = default__.safeIndex(704, (d_2636_v33_).length(0))
                    (d_2636_v33_)[index423_] = d_2633_v30_
                    index424_ = default__.safeIndex(704, (d_2636_v33_).length(0))
                    (d_2636_v33_)[index424_] = d_2633_v30_
                    source43_ = default__.fm68(globalState)
                    if source43_.is_DC28:
                        d_2637___mcc_h0_ = source43_.cf43
                        d_2638___mcc_h1_ = source43_.cf44
                        d_2639___mcc_h2_ = source43_.cf45
                        d_2640___mcc_h3_ = source43_.cf46
                        d_2641___mcc_h4_ = source43_.cf47
                        d_2642_cf47_ = d_2641___mcc_h4_
                        d_2643_cf46_ = d_2640___mcc_h3_
                        d_2644_cf45_ = d_2639___mcc_h2_
                        d_2645_cf44_ = d_2638___mcc_h1_
                        d_2646_cf43_ = d_2637___mcc_h0_
                        d_2646_cf43_ = len((d_2632_v29_) + (d_2632_v29_))
                        d_2647_v34_: _dafny.Map
                        d_2647_v34_ = _dafny.Map({(d_2633_v30_).f4: not(d_2643_cf46_)})
                        d_2647_v34_ = d_2647_v34_
                        d_2648_v35_: _dafny.Array
                        def lambda129_(d_2649_cf46_, d_2650_v30_, d_2651_cf45_):
                            def lambda130_(d_2652_i5_):
                                return _dafny.Set({len(_dafny.Set({d_2649_cf46_, (d_2650_v30_).f4, not((d_2650_v30_).f4), (self).f4, d_2651_cf45_}))})

                            return lambda130_

                        init80_ = lambda129_(d_2643_cf46_, d_2633_v30_, d_2644_cf45_)
                        nw415_ = _dafny.Array(None, 28)
                        for i0_80_ in range(nw415_.length(0)):
                            nw415_[i0_80_] = init80_(i0_80_)
                        d_2648_v35_ = nw415_
                        index425_ = default__.safeIndex(441, (d_2648_v35_).length(0))
                        (d_2648_v35_)[index425_] = _dafny.Set({d_2645_cf44_})
                        index426_ = default__.safeIndex(441, (d_2648_v35_).length(0))
                        (d_2648_v35_)[index426_] = (_dafny.Set({p0, p0, len(_dafny.SeqWithoutIsStrInference([d_2643_cf46_])), d_2646_cf43_, (self).fm7(globalState)})).intersection(_dafny.Set({(self).f6}))
                        rhs408_ = d_2644_cf45_
                        rhs409_ = ((d_2648_v35_)[default__.safeIndex(441, (d_2648_v35_).length(0))]).ispropersubset(((d_2648_v35_)[default__.safeIndex(441, (d_2648_v35_).length(0))]).intersection((d_2648_v35_)[default__.safeIndex(441, (d_2648_v35_).length(0))]))
                        d_2644_cf45_ = rhs408_
                        d_2643_cf46_ = rhs409_
                    elif True:
                        d_2653___mcc_h5_ = source43_.cf42
                        d_2654_cf42_ = d_2653___mcc_h5_
                        d_2655_v36_: _dafny.Seq
                        d_2655_v36_ = _dafny.SeqWithoutIsStrInference([(self).f4, (d_2633_v30_).f4, (d_2633_v30_).f4])
                        d_2655_v36_ = (d_2655_v36_).set(default__.safeIndex(self.f7, len(d_2655_v36_)), (self).f4)
                        d_2656_v37_: bool
                        d_2656_v37_ = False
                        d_2656_v37_ = d_2656_v37_
                        index427_ = default__.safeIndex(68, (d_2595_v0_).length(0))
                        (d_2595_v0_)[index427_] = d_2656_v37_
                        d_2657_v38_: _dafny.MultiSet
                        d_2657_v38_ = _dafny.MultiSet([(self).f4, (d_2633_v30_).f4, True])
                        index428_ = default__.safeIndex(68, (d_2595_v0_).length(0))
                        (d_2595_v0_)[index428_] = (_dafny.MultiSet([d_2656_v37_])).issubset(d_2657_v38_)
                        (self).f7 = p0
                    d_2658_v39_: _dafny.Map
                    d_2658_v39_ = _dafny.Map({d_2633_v30_: (d_2633_v30_).f4})
                    d_2659_v40_: D12
                    d_2659_v40_ = D12_DC28(self.f7, len(d_2632_v29_), (d_2633_v30_).f4, True, len(d_2632_v29_))
                    d_2660_v41_: _dafny.Set
                    d_2660_v41_ = _dafny.Set({((d_2658_v39_)[(d_2636_v33_)[default__.safeIndex(704, (d_2636_v33_).length(0))]] if ((d_2636_v33_)[default__.safeIndex(704, (d_2636_v33_).length(0))]) in (d_2658_v39_) else (d_2633_v30_).f4), False, (d_2659_v40_).cf45, (self).f4, (self).f4})
                    d_2661_v42_: _dafny.Map
                    d_2661_v42_ = _dafny.Map({(d_2633_v30_).f4: (d_2633_v30_).f4})
                    d_2662_v43_: _dafny.Seq
                    d_2662_v43_ = _dafny.SeqWithoutIsStrInference([d_2661_v42_, _dafny.Map({(self).f4: (d_2633_v30_).f4})])
                    d_2663_v44_: _dafny.Seq
                    d_2663_v44_ = _dafny.SeqWithoutIsStrInference([d_2660_v41_, default__.fm53((self).f6, d_2662_v43_, globalState)])
                    (self).f7 = default__.safeDivisionInt(len((d_2663_v44_)[default__.safeIndex(p0, len(d_2663_v44_))]), (308) + ((self).f6))
                    pass
            pass
        d_2664_v45_: _dafny.Array
        d_2664_v45_ = self.f5
        d_2665_v46_: bool
        d_2665_v46_ = False
        d_2666_v47_: _dafny.Set
        d_2666_v47_ = _dafny.Set({self.f7, self.f7})
        d_2667_v48_: _dafny.Seq
        d_2667_v48_ = _dafny.SeqWithoutIsStrInference([False, (d_2666_v47_).ispropersubset(d_2666_v47_), d_2665_v46_, (d_2665_v46_) == (d_2665_v46_)])
        d_2668_v49_: _dafny.MultiSet
        d_2668_v49_ = _dafny.MultiSet([p0, (self).f6])
        rhs410_ = d_2664_v45_
        rhs411_ = len(d_2667_v48_)
        rhs412_ = (d_2668_v49_) == ((d_2668_v49_) - (d_2668_v49_))
        lhs256_ = self
        d_2664_v45_ = rhs410_
        lhs256_.f7 = rhs411_
        d_2665_v46_ = rhs412_
        d_2669_v50_: C0
        nw416_ = C0()
        nw416_.ctor__(p0)
        d_2669_v50_ = nw416_
        d_2670_v51_: str
        d_2670_v51_ = _dafny.CodePoint('d')
        d_2671_v52_: _dafny.Seq
        d_2671_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nom"))
        d_2672_v53_: D3
        d_2672_v53_ = D3_DC9((self).f4, p0, True)
        index429_ = default__.safeIndex(544, (d_2595_v0_).length(0))
        (d_2595_v0_)[index429_] = (d_2672_v53_).cf10
        d_2673_v54_: _dafny.MultiSet
        d_2673_v54_ = _dafny.MultiSet([d_2668_v49_])
        d_2674_v55_: _dafny.Map
        d_2674_v55_ = _dafny.Map({d_2670_v51_: d_2673_v54_})
        index430_ = default__.safeIndex(544, (d_2595_v0_).length(0))
        rhs413_ = (self.f7 if d_2665_v46_ else -737)
        rhs414_ = d_2670_v51_
        rhs415_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxypm"))
        rhs416_ = (((d_2674_v55_)[d_2670_v51_] if (d_2670_v51_) in (d_2674_v55_) else d_2673_v54_)).ispropersubset(d_2673_v54_)
        rhs417_ = d_2665_v46_
        lhs257_ = self
        lhs258_ = d_2595_v0_
        lhs259_ = default__.safeIndex(544, (d_2595_v0_).length(0))
        lhs257_.f7 = rhs413_
        d_2670_v51_ = rhs414_
        d_2671_v52_ = rhs415_
        d_2665_v46_ = rhs416_
        lhs258_[lhs259_] = rhs417_
        d_2675_v56_: _dafny.Map
        d_2675_v56_ = _dafny.Map({(d_2595_v0_)[default__.safeIndex(544, (d_2595_v0_).length(0))]: (d_2669_v50_).f12})
        r0 = (d_2675_v56_) | (_dafny.Map({not((d_2595_v0_)[default__.safeIndex(544, (d_2595_v0_).length(0))]): (d_2669_v50_).f12}))
        return r0

    def m2(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_2676_v0_: C4
        nw417_ = C4()
        nw417_.ctor__(self.f5, default__.safeDivisionInt((self).fm7(globalState), self.f7), not(p0))
        d_2676_v0_ = nw417_
        d_2677_i0_: int
        d_2677_i0_ = 0
        with _dafny.label("14"):
            while p0:
                with _dafny.c_label("14"):
                    if (d_2677_i0_) >= (100):
                        raise _dafny.Break("14")
                    d_2677_i0_ = (d_2677_i0_) + (1)
                    d_2678_v1_: C8
                    nw418_ = C8()
                    nw418_.ctor__()
                    d_2678_v1_ = nw418_
                    d_2679_v2_: _dafny.Seq
                    d_2679_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bj"))
                    d_2680_v3_: _dafny.Map
                    d_2680_v3_ = _dafny.Map({self.f7: d_2679_v2_})
                    d_2680_v3_ = default__.fm69(globalState)
                    d_2681_v4_: _dafny.Array
                    def lambda131_(d_2682_i1_):
                        return default__.safeModuloInt(d_2682_i1_, len(_dafny.SeqWithoutIsStrInference([(self).f4, False])))

                    init81_ = lambda131_
                    nw419_ = _dafny.Array(None, 22)
                    for i0_81_ in range(nw419_.length(0)):
                        nw419_[i0_81_] = init81_(i0_81_)
                    d_2681_v4_ = nw419_
                    d_2681_v4_ = d_2681_v4_
                    (self).f7 = (self).f6
                    pass
            pass
        d_2683_v5_: _dafny.Array
        nw420_ = _dafny.Array(int(0), 3)
        d_2683_v5_ = nw420_
        d_2683_v5_ = d_2683_v5_
        if p0:
            index431_ = default__.safeIndex(176, (d_2683_v5_).length(0))
            (d_2683_v5_)[index431_] = (self).f6
            index432_ = default__.safeIndex(176, (d_2683_v5_).length(0))
            (d_2683_v5_)[index432_] = default__.safeModuloInt((self).f6, (self).f6)
            d_2684_v6_: bool
            d_2684_v6_ = True
            d_2685_v7_: _dafny.Set
            d_2685_v7_ = _dafny.Set({p0})
            d_2686_v8_: str
            d_2686_v8_ = _dafny.CodePoint('w')
            d_2684_v6_ = ((default__.fm46(d_2686_v8_, globalState)).intersection(d_2685_v7_)).ispropersubset((d_2685_v7_).intersection(d_2685_v7_))
            d_2687_v9_: C3
            nw421_ = C3()
            nw421_.ctor__(self.f5, (d_2683_v5_)[default__.safeIndex(176, (d_2683_v5_).length(0))], d_2684_v6_)
            d_2687_v9_ = nw421_
            d_2688_v10_: _dafny.Array
            nw422_ = _dafny.Array(_dafny.Seq({}), 24)
            d_2688_v10_ = nw422_
            index433_ = default__.safeIndex(799, (d_2688_v10_).length(0))
            (d_2688_v10_)[index433_] = _dafny.SeqWithoutIsStrInference([(d_2683_v5_)[default__.safeIndex(176, (d_2683_v5_).length(0))] for d_2689_i2_ in range(default__.abs(767))])
            d_2690_v11_: _dafny.Seq
            d_2690_v11_ = _dafny.SeqWithoutIsStrInference([self.f7])
            d_2691_v12_: _dafny.Seq
            d_2691_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xondmryd"))
            d_2692_v13_: _dafny.Seq
            d_2692_v13_ = _dafny.SeqWithoutIsStrInference([p0, (self).f4, (self).f4])
            d_2693_v14_: D12
            d_2693_v14_ = D12_DC27(_dafny.SeqWithoutIsStrInference([(d_2692_v13_).set(default__.safeIndex((d_2683_v5_)[default__.safeIndex(176, (d_2683_v5_).length(0))], len(d_2692_v13_)), d_2684_v6_)]))
            d_2694_v15_: _dafny.Seq
            d_2694_v15_ = _dafny.SeqWithoutIsStrInference([(d_2690_v11_) + (_dafny.SeqWithoutIsStrInference([(self).f6])), (d_2690_v11_) + (default__.fm39(d_2691_v12_, d_2693_v14_, (self).f4, globalState))])
            d_2695_v16_: _dafny.MultiSet
            d_2695_v16_ = _dafny.MultiSet([(self).f6, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cprbau")))])
            index434_ = default__.safeIndex(799, (d_2688_v10_).length(0))
            (d_2688_v10_)[index434_] = (d_2694_v15_)[default__.safeIndex((d_2695_v16_).cardinality, len(d_2694_v15_))]
            d_2696_v17_: _dafny.MultiSet
            d_2696_v17_ = _dafny.MultiSet([d_2691_v12_])
            d_2697_v18_: _dafny.Map
            d_2697_v18_ = _dafny.Map({(self).f4: (d_2683_v5_)[default__.safeIndex(176, (d_2683_v5_).length(0))]})
            d_2696_v17_ = (d_2696_v17_) - ((d_2696_v17_).set(d_2691_v12_, default__.abs(len(d_2697_v18_))))
        elif True:
            d_2698_v19_: _dafny.Seq
            d_2698_v19_ = _dafny.SeqWithoutIsStrInference([(self).f6, default__.safeModuloInt(self.f7, self.f7), self.f7])
            d_2698_v19_ = (d_2698_v19_) + ((d_2698_v19_) + (_dafny.SeqWithoutIsStrInference([(self).f6 for d_2699_i3_ in range(default__.abs(484))])))
            d_2700_v20_: str
            d_2700_v20_ = _dafny.CodePoint('f')
            d_2701_v21_: C5
            nw423_ = C5()
            nw423_.ctor__(not(not(p0)), self.f7, self.f5, (self).f6, default__.fm0(d_2700_v20_, globalState))
            d_2701_v21_ = nw423_
            d_2702_v22_: bool
            d_2702_v22_ = True
            d_2703_v23_: _dafny.Set
            d_2703_v23_ = _dafny.Set({p0})
            index435_ = default__.safeIndex(617, (d_2683_v5_).length(0))
            (d_2683_v5_)[index435_] = default__.safeDivisionInt(d_2701_v21_.f14, (0) - (len(d_2703_v23_)))
            d_2704_v24_: _dafny.Set
            d_2704_v24_ = _dafny.Set({(self).f6, 18})
            d_2705_v25_: _dafny.MultiSet
            d_2705_v25_ = _dafny.MultiSet([(d_2701_v21_).f13, (d_2701_v21_).f13])
            d_2706_v26_: _dafny.Map
            d_2706_v26_ = _dafny.Map({(d_2701_v21_).f13: _dafny.MultiSet([p0])})
            index436_ = default__.safeIndex(617, (d_2683_v5_).length(0))
            rhs418_ = (d_2704_v24_).issubset((d_2704_v24_) - (d_2704_v24_))
            rhs419_ = default__.safeModuloInt((self).f6, ((self).f6) * (self.f7))
            rhs420_ = (_dafny.MultiSet([False, p0])).issubset(d_2705_v25_)
            rhs421_ = (((self).f6) * (self.f7)) * (((self).f6) - ((d_2701_v21_).fm7(globalState)))
            rhs422_ = (not ((d_2701_v21_).f13) or (p0)) not in (d_2706_v26_)
            lhs260_ = d_2701_v21_
            lhs261_ = d_2683_v5_
            lhs262_ = default__.safeIndex(617, (d_2683_v5_).length(0))
            d_2702_v22_ = rhs418_
            lhs260_.f14 = rhs419_
            d_2702_v22_ = rhs420_
            lhs261_[lhs262_] = rhs421_
            d_2702_v22_ = rhs422_
            d_2702_v22_ = p0
            index437_ = default__.safeIndex(617, (d_2683_v5_).length(0))
            (d_2683_v5_)[index437_] = 45
        d_2707_v27_: _dafny.Seq
        d_2707_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "packlnnxk"))
        d_2707_v27_ = d_2707_v27_
        d_2708_v28_: _dafny.Seq
        d_2708_v28_ = _dafny.SeqWithoutIsStrInference([(self).f4])
        d_2709_v29_: str
        d_2709_v29_ = _dafny.CodePoint('y')
        d_2710_v30_: _dafny.MultiSet
        d_2710_v30_ = _dafny.MultiSet([d_2709_v29_])
        d_2711_v31_: _dafny.Seq
        d_2711_v31_ = _dafny.SeqWithoutIsStrInference([d_2710_v30_, d_2710_v30_])
        if (default__.fm70((self).f4, 736, d_2708_v28_, p0, globalState)) <= ((_dafny.SeqWithoutIsStrInference([d_2710_v30_])) + (d_2711_v31_)):
            (self).f7 = self.f7
            d_2712_v32_: _dafny.Map
            d_2712_v32_ = _dafny.Map({self.f7: (self).f4})
            rhs423_ = self.f7
            rhs424_ = (self.f7 if not (((d_2712_v32_)[len(d_2707_v27_)] if (len(d_2707_v27_)) in (d_2712_v32_) else (self).f4)) or (p0) else (self).f6)
            rhs425_ = self.f7
            rhs426_ = d_2708_v28_
            lhs263_ = self
            lhs264_ = self
            lhs265_ = self
            lhs263_.f7 = rhs423_
            lhs264_.f7 = rhs424_
            lhs265_.f7 = rhs425_
            d_2708_v28_ = rhs426_
            d_2713_v33_: bool
            d_2713_v33_ = True
            d_2713_v33_ = not(d_2713_v33_)
            d_2714_v34_: _dafny.Seq
            d_2714_v34_ = _dafny.SeqWithoutIsStrInference([self.f7])
            index438_ = default__.safeIndex(227, (d_2683_v5_).length(0))
            (d_2683_v5_)[index438_] = (len(d_2714_v34_)) * ((self).f6)
            index439_ = default__.safeIndex(13, (d_2683_v5_).length(0))
            (d_2683_v5_)[index439_] = (len(d_2708_v28_) if (self).fm8(self.f7, d_2713_v33_, d_2707_v27_, (d_2714_v34_).set(default__.safeIndex(self.f7, len(d_2714_v34_)), self.f7), globalState) else default__.fm3((self).f6, False, globalState))
            d_2715_v35_: _dafny.Map
            d_2715_v35_ = _dafny.Map({p0: False})
            d_2716_v36_: _dafny.MultiSet
            d_2716_v36_ = _dafny.MultiSet([((d_2715_v35_)[d_2713_v33_] if (d_2713_v33_) in (d_2715_v35_) else (d_2708_v28_)[default__.safeIndex(self.f7, len(d_2708_v28_))]), d_2713_v33_, not(d_2713_v33_)])
            index440_ = default__.safeIndex(227, (d_2683_v5_).length(0))
            index441_ = default__.safeIndex(13, (d_2683_v5_).length(0))
            rhs427_ = (0) - ((0) - (default__.safeDivisionInt((self).f6, ((d_2716_v36_)[(d_2676_v0_).fm8(self.f7, False, d_2707_v27_, d_2714_v34_, globalState)] if ((d_2676_v0_).fm8(self.f7, False, d_2707_v27_, d_2714_v34_, globalState)) in (d_2716_v36_) else len(_dafny.SeqWithoutIsStrInference([p0, p0]))))))
            rhs428_ = self.f7
            rhs429_ = default__.fm3((self).f6, d_2713_v33_, globalState)
            rhs430_ = (d_2708_v28_)[default__.safeIndex((self).f6, len(d_2708_v28_))]
            rhs431_ = p0
            lhs266_ = d_2683_v5_
            lhs267_ = default__.safeIndex(227, (d_2683_v5_).length(0))
            lhs268_ = d_2683_v5_
            lhs269_ = default__.safeIndex(13, (d_2683_v5_).length(0))
            lhs270_ = self
            lhs266_[lhs267_] = rhs427_
            lhs268_[lhs269_] = rhs428_
            lhs270_.f7 = rhs429_
            d_2713_v33_ = rhs430_
            d_2713_v33_ = rhs431_
            d_2717_v37_: _dafny.Seq
            d_2717_v37_ = _dafny.SeqWithoutIsStrInference([d_2707_v27_])
            d_2713_v33_ = (d_2676_v0_).fm8((self).f6, (self).f4, (d_2717_v37_)[default__.safeIndex((self).f6, len(d_2717_v37_))], d_2714_v34_, globalState)
        elif True:
            d_2718_v38_: _dafny.Array
            nw424_ = _dafny.Array(False, 1)
            d_2718_v38_ = nw424_
            d_2719_v39_: D2
            d_2719_v39_ = D2_DC5(d_2718_v38_)
            d_2719_v39_ = d_2719_v39_
            d_2718_v38_ = d_2718_v38_
            d_2720_v40_: _dafny.Array
            nw425_ = _dafny.Array(_dafny.Array(None, 0), 22)
            d_2720_v40_ = nw425_
            d_2721_v41_: _dafny.Seq
            d_2721_v41_ = _dafny.SeqWithoutIsStrInference([d_2683_v5_, d_2683_v5_])
            d_2722_v42_: _dafny.Map
            d_2722_v42_ = _dafny.Map({(d_2721_v41_)[default__.safeIndex(self.f7, len(d_2721_v41_))]: d_2720_v40_})
            d_2720_v40_ = ((d_2722_v42_)[(d_2721_v41_)[default__.safeIndex((self).f6, len(d_2721_v41_))]] if ((d_2721_v41_)[default__.safeIndex((self).f6, len(d_2721_v41_))]) in (d_2722_v42_) else d_2720_v40_)
            d_2723_v43_: _dafny.Map
            d_2723_v43_ = _dafny.Map({(0) - ((self).f6): self.f7})
            d_2724_v44_: _dafny.Seq
            d_2724_v44_ = _dafny.SeqWithoutIsStrInference([(self).f6])
            (self).f7 = default__.safeDivisionInt((480) + (((d_2723_v43_)[(d_2724_v44_)[default__.safeIndex(901, len(d_2724_v44_))]] if ((d_2724_v44_)[default__.safeIndex(901, len(d_2724_v44_))]) in (d_2723_v43_) else len(d_2708_v28_))), 444)
            d_2725_v45_: _dafny.Map
            d_2725_v45_ = _dafny.Map({d_2683_v5_: d_2718_v38_})
            (self).f7 = len(d_2725_v45_)
        r0 = d_2709_v29_
        nw426_ = _dafny.Array(D1.default()(), 9)
        r1 = nw426_
        return r0, r1

    def m3(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_2726_v0_: _dafny.Seq
        d_2726_v0_ = _dafny.SeqWithoutIsStrInference([True])
        d_2727_v1_: _dafny.Seq
        d_2727_v1_ = _dafny.SeqWithoutIsStrInference([d_2726_v0_])
        d_2728_v2_: D12
        d_2728_v2_ = D12_DC27(d_2727_v1_)
        d_2729_v3_: _dafny.Seq
        d_2729_v3_ = _dafny.SeqWithoutIsStrInference([len(default__.fm71(globalState)), 953])
        d_2730_v4_: D6
        d_2730_v4_ = D6_DC16((self).f4, not ((self).f4) or ((self).f4), (d_2729_v3_) + (d_2729_v3_))
        d_2731_v5_: _dafny.Seq
        d_2731_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wweol"))
        d_2732_v6_: _dafny.Map
        d_2732_v6_ = _dafny.Map({False: (self).f4})
        d_2733_v7_: _dafny.Set
        d_2733_v7_ = _dafny.Set({False})
        d_2734_v8_: C5
        nw427_ = C5()
        nw427_.ctor__((self).f4, len(d_2733_v7_), self.f5, (self).f6, not(default__.fm0(_dafny.CodePoint('s'), globalState)))
        d_2734_v8_ = nw427_
        d_2735_v9_: _dafny.Seq
        d_2735_v9_ = _dafny.SeqWithoutIsStrInference([d_2734_v8_])
        d_2736_v10_: _dafny.Set
        d_2736_v10_ = _dafny.Set({d_2734_v8_, (d_2735_v9_)[default__.safeIndex(self.f7, len(d_2735_v9_))]})
        d_2737_v11_: _dafny.MultiSet
        d_2737_v11_ = _dafny.MultiSet([not((d_2726_v0_)[default__.safeIndex(d_2734_v8_.f14, len(d_2726_v0_))]), (self).f4, (self).f4])
        d_2738_v12_: _dafny.Array
        nw428_ = _dafny.Array(None, 15)
        nw428_[int(0)] = (self).f6
        nw428_[int(1)] = len((d_2732_v6_) | (d_2732_v6_))
        nw428_[int(2)] = len(_dafny.Set({(self).f4, (self).f4, (self).f4, (self).f4}))
        nw428_[int(3)] = p0
        nw428_[int(4)] = default__.safeModuloInt(p0, len(d_2736_v10_))
        nw428_[int(5)] = (self).f6
        nw428_[int(6)] = p0
        nw428_[int(7)] = (d_2734_v8_).fm9((d_2734_v8_).f13, (d_2734_v8_).f13, p1, globalState)
        nw428_[int(8)] = p1
        nw428_[int(9)] = (d_2734_v8_).fm7(globalState)
        nw428_[int(10)] = (0) - (((d_2737_v11_)[False] if (False) in (d_2737_v11_) else p0))
        nw428_[int(11)] = (p0) - (self.f7)
        nw428_[int(12)] = (len(d_2729_v3_)) + (p0)
        nw428_[int(13)] = len(d_2733_v7_)
        nw428_[int(14)] = p0
        d_2738_v12_ = nw428_
        index442_ = default__.safeIndex(662, (d_2738_v12_).length(0))
        (d_2738_v12_)[index442_] = (D4_DC13(len(d_2731_v5_), (self).f4, self.f7, d_2726_v0_)).cf14
        d_2739_v13_: str
        d_2739_v13_ = _dafny.CodePoint('r')
        d_2740_v14_: D9
        d_2740_v14_ = D9_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dks")))
        index443_ = default__.safeIndex(662, (d_2738_v12_).length(0))
        rhs432_ = d_2728_v2_
        rhs433_ = D6_DC16((d_2734_v8_).f13, default__.fm0(default__.fm1(len(d_2731_v5_), (self).f4, globalState), globalState), (d_2729_v3_).set(default__.safeIndex(p0, len(d_2729_v3_)), p1))
        rhs434_ = ((d_2740_v14_ if (d_2726_v0_)[default__.safeIndex(-8, len(d_2726_v0_))] else d_2740_v14_)).cf25
        rhs435_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvpib")))
        rhs436_ = d_2739_v13_
        lhs271_ = d_2738_v12_
        lhs272_ = default__.safeIndex(662, (d_2738_v12_).length(0))
        d_2728_v2_ = rhs432_
        d_2730_v4_ = rhs433_
        d_2731_v5_ = rhs434_
        lhs271_[lhs272_] = rhs435_
        d_2739_v13_ = rhs436_
        if False:
            d_2741_v15_: bool
            d_2741_v15_ = False
            d_2742_v16_: _dafny.MultiSet
            d_2742_v16_ = _dafny.MultiSet([len(d_2726_v0_)])
            d_2741_v15_ = (d_2742_v16_).issubset((d_2742_v16_) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f7]))))
            index444_ = default__.safeIndex(662, (d_2738_v12_).length(0))
            (d_2738_v12_)[index444_] = (0) - (p0)
            index445_ = default__.safeIndex(662, (d_2738_v12_).length(0))
            (d_2738_v12_)[index445_] = self.f7
            index446_ = default__.safeIndex(662, (d_2738_v12_).length(0))
            (d_2738_v12_)[index446_] = d_2734_v8_.f14
            d_2743_v17_: _dafny.Array
            nw429_ = _dafny.Array(_dafny.Set({}), 17)
            d_2743_v17_ = nw429_
            d_2744_v18_: _dafny.Map
            d_2744_v18_ = _dafny.Map({True: d_2734_v8_.f14})
            d_2745_v19_: _dafny.Set
            d_2745_v19_ = _dafny.Set({d_2744_v18_})
            index447_ = default__.safeIndex(990, (d_2743_v17_).length(0))
            (d_2743_v17_)[index447_] = d_2745_v19_
            index448_ = default__.safeIndex(990, (d_2743_v17_).length(0))
            rhs437_ = (_dafny.Set({_dafny.Map({d_2741_v15_: d_2734_v8_.f14}), d_2744_v18_})) | (_dafny.Set({_dafny.Map({(self).f4: p0}), d_2744_v18_, d_2744_v18_, d_2744_v18_}))
            rhs438_ = (((self).f6) - ((0) - (self.f7))) - (p0)
            rhs439_ = (d_2734_v8_).f13
            lhs273_ = d_2743_v17_
            lhs274_ = default__.safeIndex(990, (d_2743_v17_).length(0))
            lhs275_ = d_2734_v8_
            lhs273_[lhs274_] = rhs437_
            lhs275_.f14 = rhs438_
            d_2741_v15_ = rhs439_
        elif True:
            d_2746_v20_: bool
            d_2746_v20_ = True
            d_2747_v21_: _dafny.Array
            nw430_ = _dafny.Array(False, 21)
            d_2747_v21_ = nw430_
            d_2748_v22_: _dafny.Map
            d_2748_v22_ = _dafny.Map({d_2747_v21_: True})
            d_2746_v20_ = ((d_2748_v22_)[d_2747_v21_] if (d_2747_v21_) in (d_2748_v22_) else (d_2734_v8_).f13)
            d_2746_v20_ = d_2746_v20_
            d_2749_v23_: C5
            nw431_ = C5()
            nw431_.ctor__(((d_2734_v8_).f13) or (d_2746_v20_), self.f7, self.f5, (self).f6, ((d_2734_v8_).f13) and ((d_2734_v8_).f13))
            d_2749_v23_ = nw431_
            d_2746_v20_ = not((d_2749_v23_).f13)
            d_2750_v24_: _dafny.MultiSet
            d_2750_v24_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([685]), d_2729_v3_, d_2729_v3_, d_2729_v3_, default__.fm39(d_2731_v5_, d_2728_v2_, (d_2734_v8_).f13, globalState)])
            d_2751_v25_: _dafny.MultiSet
            d_2751_v25_ = _dafny.MultiSet([d_2729_v3_])
            d_2750_v24_ = (d_2751_v25_)
        index449_ = default__.safeIndex(662, (d_2738_v12_).length(0))
        rhs440_ = (d_2738_v12_)[default__.safeIndex(662, (d_2738_v12_).length(0))]
        rhs441_ = d_2738_v12_
        lhs276_ = d_2738_v12_
        lhs277_ = default__.safeIndex(662, (d_2738_v12_).length(0))
        lhs276_[lhs277_] = rhs440_
        d_2738_v12_ = rhs441_
        d_2752_v26_: _dafny.Map
        d_2752_v26_ = _dafny.Map({p0: (self).f4})
        d_2753_v27_: _dafny.MultiSet
        d_2753_v27_ = _dafny.MultiSet([d_2738_v12_])
        d_2754_v28_: _dafny.Array
        def lambda132_(d_2755_v8_):
            def lambda133_(d_2756_i0_):
                return (d_2755_v8_).f13

            return lambda133_

        init82_ = lambda132_(d_2734_v8_)
        nw432_ = _dafny.Array(None, 12)
        for i0_82_ in range(nw432_.length(0)):
            nw432_[i0_82_] = init82_(i0_82_)
        d_2754_v28_ = nw432_
        d_2757_v29_: _dafny.Map
        d_2757_v29_ = _dafny.Map({len((d_2731_v5_).set(default__.safeIndex((d_2753_v27_).cardinality, len(d_2731_v5_)), p2)): len(_dafny.SeqWithoutIsStrInference([d_2754_v28_]))})
        d_2758_v30_: _dafny.Map
        d_2758_v30_ = _dafny.Map({False: len(d_2757_v29_)})
        d_2759_v31_: _dafny.Array
        nw433_ = _dafny.Array(None, 16)
        nw433_[int(0)] = (d_2734_v8_).f13
        nw433_[int(1)] = (d_2734_v8_).f13
        nw433_[int(2)] = (d_2734_v8_).fm8(p1, (d_2734_v8_).f13, d_2731_v5_, d_2729_v3_, globalState)
        nw433_[int(3)] = not((self).f4)
        nw433_[int(4)] = (self).f4
        nw433_[int(5)] = ((d_2734_v8_).f13) == ((self).f4)
        nw433_[int(6)] = not(default__.fm0(p2, globalState))
        nw433_[int(7)] = (self).f4
        nw433_[int(8)] = (d_2734_v8_).f13
        nw433_[int(9)] = ((self).f4) or ((D4_DC13((d_2738_v12_)[default__.safeIndex(662, (d_2738_v12_).length(0))], (d_2734_v8_).f13, p0, (d_2726_v0_).set(default__.safeIndex(len(d_2726_v0_), len(d_2726_v0_)), (d_2734_v8_).f13))).cf15)
        nw433_[int(10)] = (d_2734_v8_).f13
        nw433_[int(11)] = (d_2737_v11_) != (_dafny.MultiSet([((d_2752_v26_)[len(d_2758_v30_)] if (len(d_2758_v30_)) in (d_2752_v26_) else (d_2734_v8_).f13), False]))
        nw433_[int(12)] = not((_dafny.SeqWithoutIsStrInference([d_2739_v13_ for d_2760_i1_ in range(default__.abs(-420))])) < (d_2731_v5_))
        nw433_[int(13)] = (d_2734_v8_).f13
        nw433_[int(14)] = (self).f4
        nw433_[int(15)] = (self).f4
        d_2759_v31_ = nw433_
        d_2754_v28_ = d_2754_v28_
        d_2761_v32_: bool
        d_2761_v32_ = True
        d_2761_v32_ = False
        if (self).f4:
            (d_2734_v8_).f14 = default__.safeDivisionInt(d_2734_v8_.f14, (0) - ((d_2738_v12_)[default__.safeIndex(662, (d_2738_v12_).length(0))]))
            d_2761_v32_ = ((d_2734_v8_).f13) or ((_dafny.MultiSet([len(d_2726_v0_), self.f7])).isdisjoint(_dafny.MultiSet([p0])))
            index450_ = default__.safeIndex(467, (d_2754_v28_).length(0))
            (d_2754_v28_)[index450_] = (d_2734_v8_).f13
            index451_ = default__.safeIndex(467, (d_2754_v28_).length(0))
            (d_2754_v28_)[index451_] = (d_2734_v8_).f13
            d_2762_v33_: C0
            nw434_ = C0()
            nw434_.ctor__(d_2734_v8_.f14)
            d_2762_v33_ = nw434_
            index452_ = default__.safeIndex(662, (d_2738_v12_).length(0))
            (d_2738_v12_)[index452_] = (d_2762_v33_).f12
        elif True:
            (d_2734_v8_).f14 = (d_2738_v12_)[default__.safeIndex(662, (d_2738_v12_).length(0))]
            d_2732_v6_ = (d_2732_v6_).set((self).f4, (d_2734_v8_).f13)
            (d_2734_v8_).f14 = (d_2734_v8_.f14) * (248)
            index453_ = default__.safeIndex(662, (d_2738_v12_).length(0))
            (d_2738_v12_)[index453_] = p0
            (d_2734_v8_).f14 = (self).f6
        r0 = ((_dafny.Map({self.f7: (self).f6})).set(len(_dafny.Map({d_2734_v8_.f14: len(d_2731_v5_)})), p0)) | (d_2757_v29_)
        return r0

    @property
    def f8(self):
        return self._f8
