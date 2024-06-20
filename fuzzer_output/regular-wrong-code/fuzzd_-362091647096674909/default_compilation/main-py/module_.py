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
        return default__.safeModuloInt(-739, (121 if False else (_dafny.MultiSet([False])).cardinality))

    @staticmethod
    def fm1(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cww"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iekbdyjkt")) if True else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_0_i0_ in range(default__.abs(-653))])))

    @staticmethod
    def fm2(globalState):
        return not(((len(_dafny.Set({True}))) >= (-923)) or (True))

    @staticmethod
    def fm3(p0, globalState):
        return _dafny.MultiSet([(False) not in (_dafny.Map({False: True})), not(False), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvsmrmtk")))) == (-139)])

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return D0_DC1((0) - ((391) + (820)))

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({True: D0_DC1(len(_dafny.Map({553: len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rejvchf"))}))})))})) | (_dafny.Map({not(False): D0_DC1((_dafny.MultiSet([True, False])).cardinality)}))) | (_dafny.Map({not(False): D0_DC1(-482)}))

    @staticmethod
    def fm14(p0, globalState):
        return D0_DC1(((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kd"))))) - (74))

    @staticmethod
    def fm18(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(-377, 741):
                d_1_v0_: int = compr_0_
                if ((-377) <= (d_1_v0_)) and ((d_1_v0_) < (741)):
                    coll0_[(d_1_v0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "levb"))))] = (0) - (len(_dafny.Set({True, False, False})))
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(393, -101):
                d_2_v1_: int = compr_1_
                if ((393) <= (d_2_v1_)) and ((d_2_v1_) < (-101)):
                    coll1_[(d_2_v1_) + (len(_dafny.Set({True, True})))] = -736
            return _dafny.Map(coll1_)
        return _dafny.MultiSet([(0) - (len((iife0_()
        ) | (iife1_()
        ))), (-318) * (841)])

    @staticmethod
    def fm19(globalState):
        return (_dafny.Map({False: False})) | (_dafny.Map({True: not(not(True))}))

    @staticmethod
    def fm20(p0, p1, globalState):
        return D0_DC0(_dafny.MultiSet([False]))

    @staticmethod
    def fm21(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_3_i0_ in range(default__.abs(230))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtjpii"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnq"))]))

    @staticmethod
    def fm27(p0, p1, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: _dafny.Seq
            for compr_2_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([D5_DC13((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epbtsmm"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "upj"))), len(_dafny.SeqWithoutIsStrInference([-734 for d_4_i0_ in range(default__.abs(435))]))])).cardinality, False, False)]), _dafny.SeqWithoutIsStrInference([D5_DC13(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ju"))), False, True)]), _dafny.SeqWithoutIsStrInference([D5_DC13(-703, True, False), D5_DC13((_dafny.MultiSet([_dafny.CodePoint('m')])).cardinality, True, False)]), _dafny.SeqWithoutIsStrInference([D5_DC13(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_5_i1_ in range(default__.abs(194))])), True, True)])})).Elements:
                d_6_v0_: _dafny.Seq = compr_2_
                if (d_6_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([D5_DC13((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epbtsmm"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "upj"))), len(_dafny.SeqWithoutIsStrInference([-734 for d_4_i0_ in range(default__.abs(435))]))])).cardinality, False, False)]), _dafny.SeqWithoutIsStrInference([D5_DC13(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ju"))), False, True)]), _dafny.SeqWithoutIsStrInference([D5_DC13(-703, True, False), D5_DC13((_dafny.MultiSet([_dafny.CodePoint('m')])).cardinality, True, False)]), _dafny.SeqWithoutIsStrInference([D5_DC13(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_5_i1_ in range(default__.abs(194))])), True, True)])})):
                    coll2_ = coll2_.union(_dafny.Set([d_6_v0_]))
            return _dafny.Set(coll2_)
        return (iife2_()
        ).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([D5_DC13((0) - (len(_dafny.Map({False: _dafny.CodePoint('n')}))), True, False)]), _dafny.SeqWithoutIsStrInference([D5_DC13(-505, True, not(False))]), _dafny.SeqWithoutIsStrInference([D5_DC13(953, True, False)])}))

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        return D0_DC2(not((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "teghdwfgg")))}))]))).issubset(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhgmuljeq")))]))), default__.safeDivisionInt(len(_dafny.Set({False})), len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([False, True]))}))), (True if True else not(True)))

    @staticmethod
    def fm29(p0, p1, p2, p3, globalState):
        source0_ = D0_DC0(_dafny.MultiSet([False, False, False]))
        if source0_.is_DC1:
            d_7___mcc_h0_ = source0_.cf1
            d_8_cf1_ = d_7___mcc_h0_
            return _dafny.CodePoint('u')
        elif source0_.is_DC2:
            d_9___mcc_h1_ = source0_.cf2
            d_10___mcc_h2_ = source0_.cf3
            d_11___mcc_h3_ = source0_.cf4
            d_12_cf4_ = d_11___mcc_h3_
            d_13_cf3_ = d_10___mcc_h2_
            d_14_cf2_ = d_9___mcc_h1_
            return _dafny.CodePoint('h')
        elif source0_.is_DC0:
            d_15___mcc_h4_ = source0_.cf0
            d_16_cf0_ = d_15___mcc_h4_
            return _dafny.CodePoint('k')
        elif True:
            d_17___mcc_h5_ = source0_.cf5
            d_18_cf5_ = d_17___mcc_h5_
            return _dafny.CodePoint('w')

    @staticmethod
    def fm30(p0, p1, globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: _dafny.Map
            for compr_3_ in (_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umygtlqs"))), 503])) for d_20_i1_ in range(default__.abs(20))]))])): False})})).Elements:
                d_21_v0_: _dafny.Map = compr_3_
                if (d_21_v0_) in (_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umygtlqs"))), 503])) for d_20_i1_ in range(default__.abs(20))]))])): False})})):
                    coll3_ = coll3_.union(_dafny.Set([d_21_v0_]))
            return _dafny.Set(coll3_)
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: _dafny.MultiSet
            for compr_4_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(0) - (-43)]), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False, False]))])])).Elements:
                d_22_v1_: _dafny.MultiSet = compr_4_
                if (d_22_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(0) - (-43)]), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False, False]))])])):
                    coll4_ = coll4_.union(_dafny.Set([d_22_v1_]))
            return _dafny.Set(coll4_)
        return (_dafny.Set({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-681 for d_19_i0_ in range(default__.abs(966))]))]), _dafny.MultiSet([-454]), _dafny.MultiSet([-972]), _dafny.MultiSet([756, 399, len(iife3_()
        )]), _dafny.MultiSet([817, len(_dafny.Map({not(not(True)): True}))])})).intersection((_dafny.Set({_dafny.MultiSet([529, 926])})) | (iife4_()
        ))

    @staticmethod
    def fm31(p0, p1, p2, p3, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(554, 634):
                d_23_v0_: int = compr_5_
                if ((554) <= (d_23_v0_)) and ((d_23_v0_) < (634)):
                    coll5_[(d_23_v0_) + (707)] = not(False)
            return _dafny.Map(coll5_)
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(593, -150):
                d_25_v1_: int = compr_6_
                if ((593) <= (d_25_v1_)) and ((d_25_v1_) < (-150)):
                    coll6_[(d_25_v1_) * (680)] = 551
            return _dafny.Map(coll6_)
        return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(iife5_()
        )])])), len((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(True), True])): _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('q'), _dafny.CodePoint('v')]))) for d_24_i0_ in range(default__.abs(793))])})))])), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iytbot")))])).intersection(_dafny.MultiSet([len(iife6_()
        )])), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(True): not(False)})) for d_26_i1_ in range(default__.abs(-732))]))])

    @staticmethod
    def fm34(p0, p1, globalState):
        return _dafny.CodePoint('g')

    @staticmethod
    def fm35(p0, globalState):
        return (_dafny.Set({745})) - ((_dafny.Set({69})) | (_dafny.Set({184, 786, 16})))

    @staticmethod
    def fm36(globalState):
        return _dafny.SeqWithoutIsStrInference([(False) == (True), not((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqt")))) <= (695)), True, True])

    @staticmethod
    def fm37(p0, p1, globalState):
        source1_ = D21_DC50(_dafny.MultiSet([D9_DC20(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ere")), True, -948), D9_DC20(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmnxpj")), True, 189)]))
        if source1_.is_DC51:
            d_27___mcc_h0_ = source1_.cf81
            d_28_cf81_ = d_27___mcc_h0_
            def iife7_():
                coll7_ = _dafny.Map()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(-654, -939):
                    d_29_v0_: int = compr_7_
                    if ((-654) <= (d_29_v0_)) and ((d_29_v0_) < (-939)):
                        coll7_[(d_29_v0_) * (-920)] = _dafny.SeqWithoutIsStrInference([d_28_cf81_])
                return _dafny.Map(coll7_)
            return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlotprgxp"))): _dafny.SeqWithoutIsStrInference([True])})) | (iife7_()
            )
        elif source1_.is_DC50:
            d_30___mcc_h1_ = source1_.cf80
            d_31_cf80_ = d_30___mcc_h1_
            return _dafny.Map({-855: _dafny.SeqWithoutIsStrInference([not(False)])})
        elif True:
            d_32___mcc_h2_ = source1_.cf82
            d_33_cf82_ = d_32___mcc_h2_
            def iife8_():
                coll8_ = _dafny.Map()
                compr_8_: int
                for compr_8_ in _dafny.IntegerRange(201, 952):
                    d_34_v1_: int = compr_8_
                    if ((201) <= (d_34_v1_)) and ((d_34_v1_) < (952)):
                        coll8_[(d_34_v1_) + (637)] = _dafny.SeqWithoutIsStrInference([True, False, True])
                return _dafny.Map(coll8_)
            return iife8_()
            

    @staticmethod
    def fm38(p0, p1, p2, globalState):
        return D2_DC7(not(False), not((_dafny.Set({664})).issubset(_dafny.Set({len(_dafny.Set({_dafny.Set({False, False}), _dafny.Set({False})}))}))))

    @staticmethod
    def fm39(p0, globalState):
        source2_ = D5_DC14(not(False))
        if source2_.is_DC13:
            d_35___mcc_h0_ = source2_.cf20
            d_36___mcc_h1_ = source2_.cf21
            d_37___mcc_h2_ = source2_.cf22
            d_38_cf22_ = d_37___mcc_h2_
            d_39_cf21_ = d_36___mcc_h1_
            d_40_cf20_ = d_35___mcc_h0_
            def iife9_():
                coll9_ = _dafny.Map()
                compr_9_: int
                for compr_9_ in _dafny.IntegerRange(153, -772):
                    d_41_v0_: int = compr_9_
                    if ((153) <= (d_41_v0_)) and ((d_41_v0_) < (-772)):
                        coll9_[(d_41_v0_) + (d_40_cf20_)] = _dafny.SeqWithoutIsStrInference([d_39_cf21_])
                return _dafny.Map(coll9_)
            return D6_DC15(iife9_()
)
        elif source2_.is_DC14:
            d_42___mcc_h3_ = source2_.cf23
            d_43_cf23_ = d_42___mcc_h3_
            return D6_DC15(_dafny.Map({490: _dafny.SeqWithoutIsStrInference([d_43_cf23_])}))
        elif True:
            d_44___mcc_h4_ = source2_.cf19
            d_45_cf19_ = d_44___mcc_h4_
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(555, 490):
                    d_46_v1_: int = compr_10_
                    if ((555) <= (d_46_v1_)) and ((d_46_v1_) < (490)):
                        coll10_[(d_46_v1_) + (355)] = _dafny.SeqWithoutIsStrInference([True, False, True, False])
                return _dafny.Map(coll10_)
            return D6_DC15(iife10_()
)

    @staticmethod
    def fm40(p0, globalState):
        return D5_DC13(((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsqstwekh")))))) + (-868), False, True)

    @staticmethod
    def fm41(globalState):
        return (_dafny.Set({False})) | ((_dafny.Set({True, not(True), False, not(False), False})).intersection(_dafny.Set({not(True), True, True})))

    @staticmethod
    def fm42(p0, p1, globalState):
        return ((_dafny.MultiSet([898, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_47_i0_ in range(default__.abs(-776))]))]) if False else _dafny.MultiSet([827]))).intersection(_dafny.MultiSet([-34]))

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        return D0_DC0(_dafny.MultiSet([False, False, False, False]))

    @staticmethod
    def fm44(p0, globalState):
        if False:
            return D4_DC10(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_48_i0_ in range(default__.abs(582))]))
        elif True:
            return D4_DC10(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_49_i1_ in range(default__.abs(370))]))

    @staticmethod
    def fm45(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r')])) for d_50_i0_ in range(default__.abs(982))])) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqu"))) for d_51_i1_ in range(default__.abs(89))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehybm"))), len(_dafny.SeqWithoutIsStrInference([588 for d_52_i2_ in range(default__.abs(485))]))])))

    @staticmethod
    def fm46(p0, p1, globalState):
        return (_dafny.Map({-665: 731}) if True else _dafny.Map({413: 65}))

    @staticmethod
    def fm47(p0, globalState):
        return _dafny.Map({default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.Map({542: False})])), 396): _dafny.SeqWithoutIsStrInference([-939, (0) - (len(_dafny.SeqWithoutIsStrInference([False, False])))])})

    @staticmethod
    def fm48(globalState):
        return (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uarmdrl"))})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpsc"))}))

    @staticmethod
    def fm49(globalState):
        return _dafny.Map({(True if not(not(not(True))) else True): 657})

    @staticmethod
    def fm50(p0, p1, globalState):
        return (_dafny.MultiSet([D13_DC27(_dafny.Map({True: _dafny.CodePoint('u')}))])) | (_dafny.MultiSet([D13_DC27(_dafny.Map({(D13_DC29(False, False)).cf49: _dafny.CodePoint('i')})), D13_DC27(_dafny.Map({True: _dafny.CodePoint('y')})), D13_DC27(_dafny.Map({not(not(not(True))): _dafny.CodePoint('c')}))]))

    @staticmethod
    def fm51(p0, p1, p2, p3, globalState):
        def iife11_():
            def iife15_():
                def iife17_():
                    coll17_ = _dafny.Map()
                    compr_17_: int
                    for compr_17_ in _dafny.IntegerRange(297, 936):
                        d_53_v1_: int = compr_17_
                        if ((297) <= (d_53_v1_)) and ((d_53_v1_) < (936)):
                            coll17_[default__.safeModuloInt(d_53_v1_, -866)] = (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([292])), len(_dafny.Map({False: False}))])).cardinality
                    return _dafny.Map(coll17_)
                coll15_ = _dafny.Map()
                def iife16_():
                    coll16_ = _dafny.Map()
                    compr_16_: int
                    for compr_16_ in _dafny.IntegerRange(297, 936):
                        d_53_v1_: int = compr_16_
                        if ((297) <= (d_53_v1_)) and ((d_53_v1_) < (936)):
                            coll16_[default__.safeModuloInt(d_53_v1_, -866)] = (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([292])), len(_dafny.Map({False: False}))])).cardinality
                    return _dafny.Map(coll16_)
                compr_15_: _dafny.Map
                for compr_15_ in (_dafny.SeqWithoutIsStrInference([iife16_()
                , _dafny.Map({(_dafny.MultiSet([False])).cardinality: (D11_DC23(not(True), -234, _dafny.Map({True: True}), D5_DC14(False))).cf37})])).Elements:
                    d_54_v0_: _dafny.Map = compr_15_
                    if (d_54_v0_) in (_dafny.SeqWithoutIsStrInference([iife17_()
                    , _dafny.Map({(_dafny.MultiSet([False])).cardinality: (D11_DC23(not(True), -234, _dafny.Map({True: True}), D5_DC14(False))).cf37})])):
                        coll15_[d_54_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccyc"))
                return _dafny.Map(coll15_)
            coll11_ = _dafny.Set()
            def iife12_():
                def iife14_():
                    coll14_ = _dafny.Map()
                    compr_14_: int
                    for compr_14_ in _dafny.IntegerRange(297, 936):
                        d_53_v1_: int = compr_14_
                        if ((297) <= (d_53_v1_)) and ((d_53_v1_) < (936)):
                            coll14_[default__.safeModuloInt(d_53_v1_, -866)] = (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([292])), len(_dafny.Map({False: False}))])).cardinality
                    return _dafny.Map(coll14_)
                coll12_ = _dafny.Map()
                def iife13_():
                    coll13_ = _dafny.Map()
                    compr_13_: int
                    for compr_13_ in _dafny.IntegerRange(297, 936):
                        d_53_v1_: int = compr_13_
                        if ((297) <= (d_53_v1_)) and ((d_53_v1_) < (936)):
                            coll13_[default__.safeModuloInt(d_53_v1_, -866)] = (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([292])), len(_dafny.Map({False: False}))])).cardinality
                    return _dafny.Map(coll13_)
                compr_12_: _dafny.Map
                for compr_12_ in (_dafny.SeqWithoutIsStrInference([iife13_()
                , _dafny.Map({(_dafny.MultiSet([False])).cardinality: (D11_DC23(not(True), -234, _dafny.Map({True: True}), D5_DC14(False))).cf37})])).Elements:
                    d_54_v0_: _dafny.Map = compr_12_
                    if (d_54_v0_) in (_dafny.SeqWithoutIsStrInference([iife14_()
                    , _dafny.Map({(_dafny.MultiSet([False])).cardinality: (D11_DC23(not(True), -234, _dafny.Map({True: True}), D5_DC14(False))).cf37})])):
                        coll12_[d_54_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccyc"))
                return _dafny.Map(coll12_)
            compr_11_: _dafny.Map
            for compr_11_ in ((iife12_()
            ) | (_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): 973}): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_55_i0_ in range(default__.abs(-461))])}))).keys.Elements:
                d_56_v2_: _dafny.Map = compr_11_
                if (d_56_v2_) in ((iife15_()
                ) | (_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): 973}): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_55_i0_ in range(default__.abs(-461))])}))):
                    coll11_ = coll11_.union(_dafny.Set([d_56_v2_]))
            return _dafny.Set(coll11_)
        return iife11_()
        

    @staticmethod
    def fm52(p0, p1, p2, p3, globalState):
        return (_dafny.Map({False: _dafny.CodePoint('b')})) | (_dafny.Map({not(True): _dafny.CodePoint('f')}))

    @staticmethod
    def fm53(globalState):
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttv")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctxpl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljm"))})) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_57_i0_ in range(default__.abs(505))])}))

    @staticmethod
    def fm54(p0, globalState):
        def iife18_():
            def iife22_():
                def iife24_():
                    coll24_ = _dafny.Map()
                    compr_24_: int
                    for compr_24_ in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_58_i0_ in range(default__.abs(513))])), len(_dafny.Map({True: 47}))})).Elements:
                        d_59_v1_: int = compr_24_
                        if (d_59_v1_) in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_58_i0_ in range(default__.abs(513))])), len(_dafny.Map({True: 47}))})):
                            coll24_[(d_59_v1_) * (666)] = not(True)
                    return _dafny.Map(coll24_)
                coll22_ = _dafny.Map()
                def iife23_():
                    coll23_ = _dafny.Map()
                    compr_23_: int
                    for compr_23_ in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_58_i0_ in range(default__.abs(513))])), len(_dafny.Map({True: 47}))})).Elements:
                        d_59_v1_: int = compr_23_
                        if (d_59_v1_) in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_58_i0_ in range(default__.abs(513))])), len(_dafny.Map({True: 47}))})):
                            coll23_[(d_59_v1_) * (666)] = not(True)
                    return _dafny.Map(coll23_)
                compr_22_: _dafny.Set
                for compr_22_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))}), _dafny.Set({len(iife23_()
                )}), _dafny.Set({296}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([False, not(True)]))})])).Elements:
                    d_60_v0_: _dafny.Set = compr_22_
                    if (d_60_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))}), _dafny.Set({len(iife24_()
                    )}), _dafny.Set({296}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([False, not(True)]))})])):
                        coll22_[d_60_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgfnoiq"))
                return _dafny.Map(coll22_)
            coll18_ = _dafny.Set()
            def iife19_():
                def iife21_():
                    coll21_ = _dafny.Map()
                    compr_21_: int
                    for compr_21_ in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_58_i0_ in range(default__.abs(513))])), len(_dafny.Map({True: 47}))})).Elements:
                        d_59_v1_: int = compr_21_
                        if (d_59_v1_) in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_58_i0_ in range(default__.abs(513))])), len(_dafny.Map({True: 47}))})):
                            coll21_[(d_59_v1_) * (666)] = not(True)
                    return _dafny.Map(coll21_)
                coll19_ = _dafny.Map()
                def iife20_():
                    coll20_ = _dafny.Map()
                    compr_20_: int
                    for compr_20_ in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_58_i0_ in range(default__.abs(513))])), len(_dafny.Map({True: 47}))})).Elements:
                        d_59_v1_: int = compr_20_
                        if (d_59_v1_) in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_58_i0_ in range(default__.abs(513))])), len(_dafny.Map({True: 47}))})):
                            coll20_[(d_59_v1_) * (666)] = not(True)
                    return _dafny.Map(coll20_)
                compr_19_: _dafny.Set
                for compr_19_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))}), _dafny.Set({len(iife20_()
                )}), _dafny.Set({296}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([False, not(True)]))})])).Elements:
                    d_60_v0_: _dafny.Set = compr_19_
                    if (d_60_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))}), _dafny.Set({len(iife21_()
                    )}), _dafny.Set({296}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([False, not(True)]))})])):
                        coll19_[d_60_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgfnoiq"))
                return _dafny.Map(coll19_)
            compr_18_: _dafny.Set
            for compr_18_ in (iife19_()
            ).keys.Elements:
                d_61_v2_: _dafny.Set = compr_18_
                if (d_61_v2_) in (iife22_()
                ):
                    coll18_ = coll18_.union(_dafny.Set([d_61_v2_]))
            return _dafny.Set(coll18_)
        return iife18_()
        

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.Map({True: _dafny.CodePoint('f')}) for d_62_i0_ in range(default__.abs(8))])

    @staticmethod
    def fm56(p0, p1, globalState):
        def iife25_():
            coll25_ = _dafny.Map()
            compr_25_: int
            for compr_25_ in (_dafny.Map({784: (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([-19]), _dafny.SeqWithoutIsStrInference([428]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({True, True, False})), len(_dafny.Set({False})), 437])])).cardinality})).keys.Elements:
                d_63_v0_: int = compr_25_
                if (d_63_v0_) in (_dafny.Map({784: (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([-19]), _dafny.SeqWithoutIsStrInference([428]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({True, True, False})), len(_dafny.Set({False})), 437])])).cardinality})):
                    coll25_[default__.safeDivisionInt(d_63_v0_, 581)] = True
            return _dafny.Map(coll25_)
        def iife26_():
            def iife28_():
                coll28_ = _dafny.Set()
                compr_28_: int
                for compr_28_ in (_dafny.Map({88: True})).keys.Elements:
                    d_66_v2_: int = compr_28_
                    if (d_66_v2_) in (_dafny.Map({88: True})):
                        coll28_ = coll28_.union(_dafny.Set([(d_66_v2_) * ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ss")))])))))]))
                return _dafny.Set(coll28_)
            coll26_ = _dafny.Map()
            def iife27_():
                coll27_ = _dafny.Set()
                compr_27_: int
                for compr_27_ in (_dafny.Map({88: True})).keys.Elements:
                    d_64_v2_: int = compr_27_
                    if (d_64_v2_) in (_dafny.Map({88: True})):
                        coll27_ = coll27_.union(_dafny.Set([(d_64_v2_) * ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ss")))])))))]))
                return _dafny.Set(coll27_)
            compr_26_: int
            for compr_26_ in (_dafny.Set({(_dafny.MultiSet([True])).cardinality, len(iife27_()
            ), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqn"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dltidiom"))), 399})).Elements:
                d_65_v1_: int = compr_26_
                if (d_65_v1_) in (_dafny.Set({(_dafny.MultiSet([True])).cardinality, len(iife28_()
                ), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqn"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dltidiom"))), 399})):
                    coll26_[(d_65_v1_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_67_i0_ in range(default__.abs(185))])))] = _dafny.MultiSet([148, len(_dafny.Map({False: False})), 10, 813, -879])
            return _dafny.Map(coll26_)
        return (iife25_()
        ) | ((_dafny.Map({(_dafny.MultiSet([False, True])).cardinality: True})) | (_dafny.Map({len(iife26_()
        ): True})))

    @staticmethod
    def fm57(globalState):
        return D19_DC47((_dafny.Set({not(False), False})).issubset(_dafny.Set({False})))

    @staticmethod
    def fm58(p0, p1, p2, globalState):
        return (_dafny.Map({D5_DC14(True): False})) | ((_dafny.Map({D5_DC14(True): False})) | (_dafny.Map({D5_DC14(False): False})))

    @staticmethod
    def fm59(p0, globalState):
        def iife29_():
            coll29_ = _dafny.Map()
            compr_29_: _dafny.Set
            for compr_29_ in (_dafny.Map({_dafny.Set({D5_DC14(False)}): True})).keys.Elements:
                d_68_v0_: _dafny.Set = compr_29_
                if (d_68_v0_) in (_dafny.Map({_dafny.Set({D5_DC14(False)}): True})):
                    coll29_[d_68_v0_] = _dafny.Map({-697: -69})
            return _dafny.Map(coll29_)
        def iife30_():
            coll30_ = _dafny.Map()
            compr_30_: int
            for compr_30_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({109: not(False)}))])).Elements:
                d_69_v1_: int = compr_30_
                if (d_69_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({109: not(False)}))])):
                    coll30_[(d_69_v1_) * (-893)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qubdos")))
            return _dafny.Map(coll30_)
        return (_dafny.Map({_dafny.Set({D5_DC14(True), D5_DC14(True), D5_DC14(False)}): _dafny.Map({655: -349})})) | ((iife29_()
        ) | ((D31_DC68(_dafny.Map({_dafny.Set({D5_DC14(True)}): iife30_()
}))).cf112))

    @staticmethod
    def fm60(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([486 for d_70_i0_ in range(default__.abs(219))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-574 for d_71_i2_ in range(default__.abs(171))]) for d_72_i1_ in range(default__.abs(449))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([337 for d_73_i3_ in range(default__.abs(78))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-185, (_dafny.MultiSet([False])).cardinality, 734, len(_dafny.SeqWithoutIsStrInference([781 for d_74_i5_ in range(default__.abs(913))])), -302]) for d_75_i4_ in range(default__.abs(974))])))

    @staticmethod
    def fm61(p0, p1, p2, p3, globalState):
        def iife31_():
            def iife33_():
                coll33_ = _dafny.Set()
                compr_33_: int
                for compr_33_ in (_dafny.Set({-448, 124})).Elements:
                    d_79_v0_: int = compr_33_
                    if (d_79_v0_) in (_dafny.Set({-448, 124})):
                        coll33_ = coll33_.union(_dafny.Set([(d_79_v0_) + (212)]))
                return _dafny.Set(coll33_)
            coll31_ = _dafny.Set()
            def iife32_():
                coll32_ = _dafny.Set()
                compr_32_: int
                for compr_32_ in (_dafny.Set({-448, 124})).Elements:
                    d_76_v0_: int = compr_32_
                    if (d_76_v0_) in (_dafny.Set({-448, 124})):
                        coll32_ = coll32_.union(_dafny.Set([(d_76_v0_) + (212)]))
                return _dafny.Set(coll32_)
            compr_31_: _dafny.Seq
            for compr_31_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, False])), len(_dafny.SeqWithoutIsStrInference([len(iife32_()
) for d_77_i0_ in range(default__.abs(-252))]))])])).Elements:
                d_78_v1_: _dafny.Seq = compr_31_
                if (d_78_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, False])), len(_dafny.SeqWithoutIsStrInference([len(iife33_()
) for d_77_i0_ in range(default__.abs(-252))]))])])):
                    coll31_ = coll31_.union(_dafny.Set([d_78_v1_]))
            return _dafny.Set(coll31_)
        def iife34_():
            coll34_ = _dafny.Set()
            compr_34_: _dafny.Seq
            for compr_34_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([837]): D5_DC14(False)})).keys.Elements:
                d_82_v2_: _dafny.Seq = compr_34_
                if (d_82_v2_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([837]): D5_DC14(False)})):
                    coll34_ = coll34_.union(_dafny.Set([d_82_v2_]))
            return _dafny.Set(coll34_)
        return ((iife31_()
        ) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "faysnx"))), -59, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_80_i1_ in range(default__.abs(-141))])), len(_dafny.Map({False: True})), 494})), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-772, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_81_i2_ in range(default__.abs(-478))]))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qs"))), 995]))]))])}))) - ((_dafny.Set({_dafny.SeqWithoutIsStrInference([88]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, True, False}))]), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False, True, False, False])).cardinality]), _dafny.SeqWithoutIsStrInference([-711, len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eavsgwtj")))]), _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({_dafny.Map({_dafny.CodePoint('j'): len(_dafny.SeqWithoutIsStrInference([True]))}): -41}))), 753])})) | (iife34_()
        ))

    @staticmethod
    def fm62(p0, globalState):
        source3_ = D4_DC11((0) - (len(_dafny.SeqWithoutIsStrInference([not(False)]))))
        if source3_.is_DC11:
            d_83___mcc_h0_ = source3_.cf18
            d_84_cf18_ = d_83___mcc_h0_
            def iife35_():
                def iife39_():
                    def iife41_():
                        coll41_ = _dafny.Map()
                        compr_41_: _dafny.Seq
                        for compr_41_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True})).keys.Elements:
                            d_85_v1_: _dafny.Seq = compr_41_
                            if (d_85_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True})):
                                coll41_[d_85_v1_] = d_84_cf18_
                        return _dafny.Map(coll41_)
                    coll39_ = _dafny.Set()
                    def iife40_():
                        coll40_ = _dafny.Map()
                        compr_40_: _dafny.Seq
                        for compr_40_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True})).keys.Elements:
                            d_85_v1_: _dafny.Seq = compr_40_
                            if (d_85_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True})):
                                coll40_[d_85_v1_] = d_84_cf18_
                        return _dafny.Map(coll40_)
                    compr_39_: _dafny.Seq
                    for compr_39_ in (iife40_()
                    ).keys.Elements:
                        d_88_v2_: _dafny.Seq = compr_39_
                        if (d_88_v2_) in (iife41_()
                        ):
                            coll39_ = coll39_.union(_dafny.Set([d_88_v2_]))
                    return _dafny.Set(coll39_)
                coll35_ = _dafny.Map()
                def iife36_():
                    def iife38_():
                        coll38_ = _dafny.Map()
                        compr_38_: _dafny.Seq
                        for compr_38_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True})).keys.Elements:
                            d_85_v1_: _dafny.Seq = compr_38_
                            if (d_85_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True})):
                                coll38_[d_85_v1_] = d_84_cf18_
                        return _dafny.Map(coll38_)
                    coll36_ = _dafny.Set()
                    def iife37_():
                        coll37_ = _dafny.Map()
                        compr_37_: _dafny.Seq
                        for compr_37_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True})).keys.Elements:
                            d_85_v1_: _dafny.Seq = compr_37_
                            if (d_85_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True})):
                                coll37_[d_85_v1_] = d_84_cf18_
                        return _dafny.Map(coll37_)
                    compr_36_: _dafny.Seq
                    for compr_36_ in (iife37_()
                    ).keys.Elements:
                        d_86_v2_: _dafny.Seq = compr_36_
                        if (d_86_v2_) in (iife38_()
                        ):
                            coll36_ = coll36_.union(_dafny.Set([d_86_v2_]))
                    return _dafny.Set(coll36_)
                compr_35_: _dafny.Seq
                for compr_35_ in (iife36_()
                ).Elements:
                    d_87_v0_: _dafny.Seq = compr_35_
                    if (d_87_v0_) in (iife39_()
                    ):
                        coll35_[d_87_v0_] = True
                return _dafny.Map(coll35_)
            return iife35_()
            
        elif True:
            d_89___mcc_h1_ = source3_.cf17
            d_90_cf17_ = d_89___mcc_h1_
            return (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): not(True)})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True}))

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: int = int(0)
        d_91_v0_: _dafny.Array
        nw0_ = _dafny.Array(_dafny.Array(None, 0), 16)
        d_91_v0_ = nw0_
        d_92_v1_: _dafny.Map
        d_92_v1_ = _dafny.Map({d_91_v0_: (p0) != (p3)})
        if ((d_92_v1_)[d_91_v0_] if (d_91_v0_) in (d_92_v1_) else not (p2) or (not(p2))):
            d_93_v2_: _dafny.Set
            d_93_v2_ = _dafny.Set({False, p2})
            d_94_v3_: _dafny.Map
            d_94_v3_ = _dafny.Map({d_93_v2_: p0})
            d_95_v4_: _dafny.Seq
            d_95_v4_ = _dafny.SeqWithoutIsStrInference([p3])
            d_96_v5_: _dafny.Seq
            d_96_v5_ = _dafny.SeqWithoutIsStrInference([p2, p2])
            d_97_v6_: _dafny.Map
            d_97_v6_ = _dafny.Map({p0: p3})
            d_98_v7_: _dafny.Map
            d_98_v7_ = _dafny.Map({len(d_95_v4_): len((d_96_v5_).set(default__.safeIndex((0) - (len(d_97_v6_)), len(d_96_v5_)), False))})
            d_94_v3_ = (d_94_v3_).set(_dafny.Set({p2}), ((d_97_v6_)[-908] if (-908) in (d_97_v6_) else p0))
            d_99_v8_: _dafny.Array
            def lambda0_(d_100_p2_):
                def lambda1_(d_101_i0_):
                    return d_100_p2_

                return lambda1_

            init0_ = lambda0_(p2)
            nw1_ = _dafny.Array(None, 17)
            for i0_0_ in range(nw1_.length(0)):
                nw1_[i0_0_] = init0_(i0_0_)
            d_99_v8_ = nw1_
            d_102_v9_: _dafny.Seq
            d_102_v9_ = _dafny.SeqWithoutIsStrInference([d_99_v8_])
            d_103_v10_: D27
            d_103_v10_ = D27_DC63((p1).cardinality, p0, p0, p2)
            d_104_v11_: _dafny.Seq
            d_104_v11_ = _dafny.SeqWithoutIsStrInference([d_95_v4_, d_95_v4_, _dafny.SeqWithoutIsStrInference([p0 for d_105_i2_ in range(default__.abs(985))]), d_95_v4_])
            d_106_v12_: _dafny.Map
            d_106_v12_ = _dafny.Map({(d_103_v10_).cf107: d_104_v11_})
            d_107_v13_: D1
            d_107_v13_ = D1_DC5(p3, p0, ((d_106_v12_)[True] if (True) in (d_106_v12_) else d_104_v11_), d_99_v8_, p0)
            d_108_v14_: _dafny.Set
            d_108_v14_ = _dafny.Set({len(_dafny.Map({p0: p3}))})
            d_109_v15_: _dafny.Map
            d_109_v15_ = _dafny.Map({p2: d_99_v8_})
            pat_let_tv0_ = d_104_v11_
            pat_let_tv1_ = p3
            d_110_v16_: _dafny.Array
            nw2_ = _dafny.Array(None, 17)
            nw2_[int(0)] = D1_DC5(p3, p3, _dafny.SeqWithoutIsStrInference([d_95_v4_ for d_111_i1_ in range(default__.abs(844))]), d_99_v8_, 685)
            nw2_[int(1)] = d_107_v13_
            nw2_[int(2)] = d_107_v13_
            nw2_[int(3)] = d_107_v13_
            def iife42_(_pat_let0_0):
                def iife43_(d_112_dt__update__tmp_h0_):
                    def iife44_(_pat_let1_0):
                        def iife45_(d_113_dt__update_hcf9_h0_):
                            return D1_DC5((d_112_dt__update__tmp_h0_).cf7, (d_112_dt__update__tmp_h0_).cf8, d_113_dt__update_hcf9_h0_, (d_112_dt__update__tmp_h0_).cf10, (d_112_dt__update__tmp_h0_).cf11)
                        return iife45_(_pat_let1_0)
                    return iife44_(pat_let_tv0_)
                return iife43_(_pat_let0_0)
            nw2_[int(4)] = iife42_(d_107_v13_)
            nw2_[int(5)] = d_107_v13_
            nw2_[int(6)] = (d_107_v13_ if p2 else d_107_v13_)
            nw2_[int(7)] = (d_107_v13_ if p2 else D1_DC5(p0, p3, d_104_v11_, d_99_v8_, (0) - (p3)))
            nw2_[int(8)] = D1_DC5(p3, 806, d_104_v11_, d_99_v8_, p0)
            nw2_[int(9)] = D1_DC5(p0, p3, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(p1).cardinality, p3, p0, p3, len(d_108_v14_)])]), ((d_109_v15_)[p2] if (p2) in (d_109_v15_) else d_99_v8_), p3)
            nw2_[int(10)] = D1_DC5(p3, p3, d_104_v11_, d_99_v8_, p0)
            nw2_[int(11)] = d_107_v13_
            nw2_[int(12)] = d_107_v13_
            nw2_[int(13)] = d_107_v13_
            nw2_[int(14)] = d_107_v13_
            nw2_[int(15)] = d_107_v13_
            def iife46_(_pat_let2_0):
                def iife47_(d_114_dt__update__tmp_h1_):
                    def iife48_(_pat_let3_0):
                        def iife49_(d_115_dt__update_hcf11_h0_):
                            return D1_DC5((d_114_dt__update__tmp_h1_).cf7, (d_114_dt__update__tmp_h1_).cf8, (d_114_dt__update__tmp_h1_).cf9, (d_114_dt__update__tmp_h1_).cf10, d_115_dt__update_hcf11_h0_)
                        return iife49_(_pat_let3_0)
                    return iife48_(pat_let_tv1_)
                return iife47_(_pat_let2_0)
            nw2_[int(16)] = iife46_(d_107_v13_)
            d_110_v16_ = nw2_
            index0_ = default__.safeIndex(593, (d_110_v16_).length(0))
            (d_110_v16_)[index0_] = d_107_v13_
            d_116_v17_: _dafny.Seq
            d_116_v17_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_99_v8_]), d_102_v9_, d_102_v9_, d_102_v9_, d_102_v9_])
            d_117_v18_: _dafny.Map
            d_117_v18_ = _dafny.Map({p0: (d_116_v17_)[default__.safeIndex(p3, len(d_116_v17_))]})
            d_118_v19_: _dafny.Seq
            d_118_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlsyfds"))
            index1_ = default__.safeIndex(593, (d_110_v16_).length(0))
            rhs0_ = (d_102_v9_) + (((d_117_v18_)[len(d_118_v19_)] if (len(d_118_v19_)) in (d_117_v18_) else d_102_v9_))
            rhs1_ = d_107_v13_
            lhs0_ = d_110_v16_
            lhs1_ = default__.safeIndex(593, (d_110_v16_).length(0))
            d_102_v9_ = rhs0_
            lhs0_[lhs1_] = rhs1_
            d_119_v20_: D0
            d_119_v20_ = D0_DC0(p1)
            d_120_v21_: D0
            d_120_v21_ = D0_DC0((p1) | ((d_119_v20_).cf0))
            source4_ = d_120_v21_
            if source4_.is_DC1:
                d_121___mcc_h0_ = source4_.cf1
                d_122_cf1_ = d_121___mcc_h0_
                d_123_v22_: D9
                d_123_v22_ = D9_DC20(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_124_i3_ in range(default__.abs(86))]), p2, p3)
                d_125_v23_: _dafny.MultiSet
                d_125_v23_ = _dafny.MultiSet([d_123_v22_])
                d_126_v24_: D21
                d_126_v24_ = D21_DC50(d_125_v23_)
                d_127_v25_: _dafny.Map
                d_127_v25_ = _dafny.Map({(d_96_v5_)[default__.safeIndex(p0, len(d_96_v5_))]: p2})
                d_128_v26_: D5
                d_128_v26_ = D5_DC14(p2)
                d_129_v27_: D11
                d_129_v27_ = D11_DC23(p2, d_122_cf1_, _dafny.Map({p2: p2}), d_128_v26_)
                d_130_v28_: _dafny.Array
                nw3_ = _dafny.Array(None, 15)
                nw3_[int(0)] = _dafny.Map({p2: p2})
                nw3_[int(1)] = _dafny.Map({p2: p2})
                nw3_[int(2)] = d_127_v25_
                nw3_[int(3)] = d_127_v25_
                nw3_[int(4)] = d_127_v25_
                nw3_[int(5)] = d_127_v25_
                nw3_[int(6)] = d_127_v25_
                nw3_[int(7)] = d_127_v25_
                nw3_[int(8)] = d_127_v25_
                nw3_[int(9)] = (d_129_v27_).cf38
                nw3_[int(10)] = d_127_v25_
                nw3_[int(11)] = (d_129_v27_).cf38
                nw3_[int(12)] = _dafny.Map({(d_96_v5_)[default__.safeIndex(p3, len(d_96_v5_))]: p2})
                nw3_[int(13)] = d_127_v25_
                nw3_[int(14)] = d_127_v25_
                d_130_v28_ = nw3_
                d_131_v29_: _dafny.MultiSet
                d_131_v29_ = _dafny.MultiSet([d_118_v19_])
                d_132_v30_: _dafny.MultiSet
                d_132_v30_ = d_131_v29_
                r0 = (d_122_cf1_) - ((0) - ((D23_DC55(d_126_v24_, d_130_v28_, d_132_v30_, p2, len(_dafny.Map({p2: p2})))).cf89))
                pat_let_tv2_ = d_93_v2_
                pat_let_tv3_ = d_93_v2_
                def iife50_(_pat_let4_0):
                    def iife51_(d_133_dt__update__tmp_h2_):
                        def iife52_(_pat_let5_0):
                            def iife53_(d_134_dt__update_hcf23_h0_):
                                return D5_DC14(d_134_dt__update_hcf23_h0_)
                            return iife53_(_pat_let5_0)
                        return iife52_((pat_let_tv2_).issubset(pat_let_tv3_))
                    return iife51_(_pat_let4_0)
                rhs2_ = p2
                rhs3_ = iife50_(d_128_v26_)
                rhs4_ = d_118_v19_
                rhs5_ = default__.safeDivisionInt((p3) - (len(d_96_v5_)), p0)
                lhs2_ = globalState
                lhs2_.f13 = rhs2_
                d_128_v26_ = rhs3_
                d_118_v19_ = rhs4_
                r0 = rhs5_
                index2_ = default__.safeIndex(212, (d_99_v8_).length(0))
                (d_99_v8_)[index2_] = (not(p2)) or (p2)
                d_135_v31_: _dafny.Array
                nw4_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
                d_135_v31_ = nw4_
                index3_ = default__.safeIndex(212, (d_99_v8_).length(0))
                rhs6_ = (((_dafny.SeqWithoutIsStrInference([(d_118_v19_)[default__.safeIndex(p0, len(d_118_v19_))] for d_136_i4_ in range(default__.abs(-374))])) + (d_118_v19_)) + ((d_118_v19_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnqxtfmaj"))))).set(default__.safeIndex(default__.safeDivisionInt(p3, d_122_cf1_), len(((_dafny.SeqWithoutIsStrInference([(d_118_v19_)[default__.safeIndex(p0, len(d_118_v19_))] for d_137_i4_ in range(default__.abs(-374))])) + (d_118_v19_)) + ((d_118_v19_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnqxtfmaj")))))), _dafny.CodePoint('q'))
                rhs7_ = (d_122_cf1_) + (p0)
                rhs8_ = (p2 if p2 else p2)
                rhs9_ = p2
                rhs10_ = d_135_v31_
                lhs3_ = globalState
                lhs4_ = d_99_v8_
                lhs5_ = default__.safeIndex(212, (d_99_v8_).length(0))
                d_118_v19_ = rhs6_
                lhs3_.f7 = rhs7_
                r1 = rhs8_
                lhs4_[lhs5_] = rhs9_
                d_135_v31_ = rhs10_
                d_138_v32_: str
                d_138_v32_ = _dafny.CodePoint('q')
                d_139_v33_: T1
                nw5_ = C6()
                nw5_.ctor__(d_138_v32_, (d_96_v5_)[default__.safeIndex(p0, len(d_96_v5_))])
                d_139_v33_ = nw5_
                d_140_v34_: _dafny.Array
                nw6_ = _dafny.Array(None, 4)
                nw6_[int(0)] = d_139_v33_.f23
                nw6_[int(1)] = d_139_v33_.f23
                nw6_[int(2)] = d_138_v32_
                nw6_[int(3)] = d_139_v33_.f23
                d_140_v34_ = nw6_
                nw7_ = C10()
                nw7_.ctor__(p2, default__.safeDivisionInt(872, p0), d_138_v32_, p2, d_140_v34_, default__.fm2(globalState))
                d_139_v33_ = nw7_
            elif source4_.is_DC2:
                d_141___mcc_h1_ = source4_.cf2
                d_142___mcc_h2_ = source4_.cf3
                d_143___mcc_h3_ = source4_.cf4
                d_144_cf4_ = d_143___mcc_h3_
                d_145_cf3_ = d_142___mcc_h2_
                d_146_cf2_ = d_141___mcc_h1_
                d_147_v35_: _dafny.Array
                nw8_ = _dafny.Array(int(0), 27)
                d_147_v35_ = nw8_
                d_148_v36_: _dafny.Seq
                d_148_v36_ = _dafny.SeqWithoutIsStrInference([d_147_v35_, d_147_v35_])
                d_149_v37_: D18
                d_149_v37_ = D18_DC42((p3) - (p3), (d_147_v35_) in (d_148_v36_))
                d_149_v37_ = d_149_v37_
                d_150_v38_: _dafny.Map
                d_150_v38_ = _dafny.Map({(d_118_v19_) < (d_118_v19_): p0})
                d_150_v38_ = (d_150_v38_).set((d_146_cf2_) == (False), p0)
                d_151_v39_: _dafny.Array
                nw9_ = _dafny.Array(_dafny.Map({}), 20)
                d_151_v39_ = nw9_
                d_152_v40_: _dafny.Map
                d_152_v40_ = _dafny.Map({d_144_cf4_: d_146_cf2_})
                index4_ = default__.safeIndex(375, (d_151_v39_).length(0))
                (d_151_v39_)[index4_] = d_152_v40_
                index5_ = default__.safeIndex(375, (d_151_v39_).length(0))
                (d_151_v39_)[index5_] = d_152_v40_
                d_153_v41_: D19
                d_153_v41_ = D19_DC47(True)
                d_154_v42_: D19
                d_154_v42_ = D19_DC48(d_153_v41_)
                d_155_v43_: D19
                d_155_v43_ = D19_DC48(d_154_v42_)
                pat_let_tv4_ = d_154_v42_
                def iife54_(_pat_let6_0):
                    def iife55_(d_156_dt__update__tmp_h3_):
                        def iife56_(_pat_let7_0):
                            def iife57_(d_157_dt__update_hcf78_h0_):
                                return D19_DC48(d_157_dt__update_hcf78_h0_)
                            return iife57_(_pat_let7_0)
                        return iife56_(pat_let_tv4_)
                    return iife55_(_pat_let6_0)
                d_155_v43_ = iife54_(d_155_v43_)
            elif source4_.is_DC0:
                d_158___mcc_h4_ = source4_.cf0
                d_159_cf0_ = d_158___mcc_h4_
                d_160_v44_: str
                d_160_v44_ = _dafny.CodePoint('f')
                d_161_v45_: _dafny.MultiSet
                d_161_v45_ = _dafny.MultiSet([d_160_v44_])
                d_161_v45_ = d_161_v45_
                d_162_v46_: _dafny.Array
                nw10_ = _dafny.Array(_dafny.CodePoint('D'), 26)
                d_162_v46_ = nw10_
                index6_ = default__.safeIndex(199, (d_162_v46_).length(0))
                (d_162_v46_)[index6_] = d_160_v44_
                index7_ = default__.safeIndex(199, (d_162_v46_).length(0))
                (d_162_v46_)[index7_] = d_160_v44_
                d_163_v47_: _dafny.MultiSet
                d_163_v47_ = _dafny.MultiSet([d_118_v19_, d_118_v19_])
                d_164_v48_: _dafny.Map
                d_164_v48_ = _dafny.Map({p3: d_163_v47_})
                d_163_v47_ = ((d_164_v48_)[(p3) + (p0)] if ((p3) + (p0)) in (d_164_v48_) else _dafny.MultiSet([d_118_v19_]))
                d_165_v49_: _dafny.Seq
                d_165_v49_ = _dafny.SeqWithoutIsStrInference([d_118_v19_, d_118_v19_])
                d_118_v19_ = ((d_118_v19_) + (d_118_v19_)) + ((d_165_v49_)[default__.safeIndex(p3, len(d_165_v49_))])
            elif True:
                d_166___mcc_h5_ = source4_.cf5
                d_167_cf5_ = d_166___mcc_h5_
                d_168_v50_: str
                d_168_v50_ = _dafny.CodePoint('v')
                d_169_v51_: T0
                nw11_ = C0()
                nw11_.ctor__(default__.fm48(globalState), p2, d_168_v50_, p2)
                d_169_v51_ = nw11_
                d_170_v52_: _dafny.Array
                nw12_ = _dafny.Array(None, 16)
                nw12_[int(0)] = d_169_v51_
                nw12_[int(1)] = d_169_v51_
                nw12_[int(2)] = d_169_v51_
                nw12_[int(3)] = d_169_v51_
                nw12_[int(4)] = d_169_v51_
                nw12_[int(5)] = d_169_v51_
                nw12_[int(6)] = d_169_v51_
                nw12_[int(7)] = d_169_v51_
                nw12_[int(8)] = d_169_v51_
                nw12_[int(9)] = d_169_v51_
                nw12_[int(10)] = d_169_v51_
                nw12_[int(11)] = d_169_v51_
                nw12_[int(12)] = (d_169_v51_)
                nw12_[int(13)] = d_169_v51_
                nw12_[int(14)] = d_169_v51_
                nw12_[int(15)] = d_169_v51_
                d_170_v52_ = nw12_
                index8_ = default__.safeIndex(56, (d_170_v52_).length(0))
                (d_170_v52_)[index8_] = d_169_v51_
                index9_ = default__.safeIndex(56, (d_170_v52_).length(0))
                (d_170_v52_)[index9_] = d_169_v51_
                d_171_v53_: _dafny.Array
                nw13_ = _dafny.Array(_dafny.Set({}), 17)
                d_171_v53_ = nw13_
                d_172_v55_: _dafny.Set
                d_172_v55_ = _dafny.Set({d_169_v51_.f23})
                index10_ = default__.safeIndex(588, (d_171_v53_).length(0))
                def iife58_():
                    def iife60_():
                        coll44_ = _dafny.Map()
                        compr_44_: str
                        for compr_44_ in (d_172_v55_).Elements:
                            d_173_v54_: str = compr_44_
                            if (d_173_v54_) in (d_172_v55_):
                                coll44_[d_173_v54_] = not((d_169_v51_).f24)
                        return _dafny.Map(coll44_)
                    coll42_ = _dafny.Set()
                    def iife59_():
                        coll43_ = _dafny.Map()
                        compr_43_: str
                        for compr_43_ in (d_172_v55_).Elements:
                            d_173_v54_: str = compr_43_
                            if (d_173_v54_) in (d_172_v55_):
                                coll43_[d_173_v54_] = not((d_169_v51_).f24)
                        return _dafny.Map(coll43_)
                    compr_42_: str
                    for compr_42_ in (iife59_()
                    ).keys.Elements:
                        d_174_v56_: str = compr_42_
                        if (d_174_v56_) in (iife60_()
                        ):
                            coll42_ = coll42_.union(_dafny.Set([d_174_v56_]))
                    return _dafny.Set(coll42_)
                (d_171_v53_)[index10_] = (iife58_()
                ).intersection(d_172_v55_)
                index11_ = default__.safeIndex(588, (d_171_v53_).length(0))
                (d_171_v53_)[index11_] = d_172_v55_
                d_175_v57_: _dafny.Array
                def lambda2_(d_176_v51_):
                    def lambda3_(d_177_i5_):
                        return d_176_v51_.f23

                    return lambda3_

                init1_ = lambda2_(d_169_v51_)
                nw14_ = _dafny.Array(None, 5)
                for i0_1_ in range(nw14_.length(0)):
                    nw14_[i0_1_] = init1_(i0_1_)
                d_175_v57_ = nw14_
                d_178_v58_: C7
                nw15_ = C7()
                nw15_.ctor__((d_169_v51_).f24, True, d_175_v57_, d_169_v51_.f23, False, not((_dafny.MultiSet([p0, 448, p0])).ispropersubset(_dafny.MultiSet(d_95_v4_))))
                d_178_v58_ = nw15_
                d_178_v58_ = d_178_v58_
                (globalState).f13 = (d_178_v58_).f31
            d_179_v59_: _dafny.Map
            d_179_v59_ = d_97_v6_
            d_180_v60_: _dafny.Array
            def lambda4_(d_181_i6_):
                return _dafny.CodePoint('q')

            init2_ = lambda4_
            nw16_ = _dafny.Array(None, 29)
            for i0_2_ in range(nw16_.length(0)):
                nw16_[i0_2_] = init2_(i0_2_)
            d_180_v60_ = nw16_
            d_182_v61_: _dafny.Map
            d_182_v61_ = _dafny.Map({p2: d_180_v60_})
            d_183_v62_: str
            d_183_v62_ = _dafny.CodePoint('s')
            d_184_v63_: C4
            nw17_ = C4()
            nw17_.ctor__(len(d_118_v19_), d_179_v59_, ((d_182_v61_)[False] if (False) in (d_182_v61_) else d_180_v60_), d_183_v62_, p2)
            d_184_v63_ = nw17_
            (globalState).f13 = (p2 if (d_184_v63_).fm22(p0, p2, globalState) else p2)
        elif True:
            r2 = default__.safeDivisionInt(-599, len(default__.fm62(p0, globalState)))
            r0 = p0
            d_185_v64_: _dafny.Array
            nw18_ = _dafny.Array(False, 23)
            d_185_v64_ = nw18_
            index12_ = default__.safeIndex(7, (d_185_v64_).length(0))
            (d_185_v64_)[index12_] = False
            index13_ = default__.safeIndex(7, (d_185_v64_).length(0))
            (d_185_v64_)[index13_] = p2
            d_186_v65_: _dafny.Seq
            d_186_v65_ = _dafny.SeqWithoutIsStrInference([not((d_185_v64_)[default__.safeIndex(7, (d_185_v64_).length(0))])])
            d_187_v66_: _dafny.Map
            d_187_v66_ = _dafny.Map({p0: d_186_v65_})
            d_188_v67_: _dafny.Set
            d_188_v67_ = _dafny.Set({425})
            d_189_v68_: _dafny.Seq
            d_189_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
            d_190_v69_: _dafny.Map
            d_190_v69_ = _dafny.Map({p2: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfiap")))})
            d_191_v70_: _dafny.Seq
            d_191_v70_ = _dafny.SeqWithoutIsStrInference([p3, p3, p0, 141, p3])
            d_192_v71_: _dafny.Set
            d_192_v71_ = _dafny.Set({(d_185_v64_)[default__.safeIndex(7, (d_185_v64_).length(0))], (d_185_v64_)[default__.safeIndex(7, (d_185_v64_).length(0))]})
            d_193_v72_: _dafny.MultiSet
            d_193_v72_ = _dafny.MultiSet([d_192_v71_])
            d_194_v74_: _dafny.Array
            nw19_ = _dafny.Array(None, 29)
            nw19_[int(0)] = (p3) * (len(d_187_v66_))
            nw19_[int(1)] = len(d_188_v67_)
            nw19_[int(2)] = p3
            nw19_[int(3)] = -354
            nw19_[int(4)] = -37
            nw19_[int(5)] = default__.safeModuloInt(p3, p3)
            nw19_[int(6)] = p3
            nw19_[int(7)] = len(d_189_v68_)
            nw19_[int(8)] = 641
            nw19_[int(9)] = (len(d_190_v69_)) - (((d_190_v69_)[not((d_185_v64_)[default__.safeIndex(7, (d_185_v64_).length(0))])] if (not((d_185_v64_)[default__.safeIndex(7, (d_185_v64_).length(0))])) in (d_190_v69_) else p3))
            nw19_[int(10)] = p3
            nw19_[int(11)] = len(d_191_v70_)
            nw19_[int(12)] = default__.safeDivisionInt(p0, p0)
            nw19_[int(13)] = (p3 if True else (0) - (p3))
            nw19_[int(14)] = p3
            nw19_[int(15)] = ((d_193_v72_)[d_192_v71_] if (d_192_v71_) in (d_193_v72_) else 184)
            nw19_[int(16)] = p3
            nw19_[int(17)] = p0
            nw19_[int(18)] = p0
            nw19_[int(19)] = p3
            nw19_[int(20)] = p0
            nw19_[int(21)] = p3
            nw19_[int(22)] = p0
            nw19_[int(23)] = p3
            nw19_[int(24)] = p3
            nw19_[int(25)] = -424
            nw19_[int(26)] = p3
            def iife61_():
                coll45_ = _dafny.Set()
                compr_45_: int
                for compr_45_ in _dafny.IntegerRange(649, 62):
                    d_195_v73_: int = compr_45_
                    if ((649) <= (d_195_v73_)) and ((d_195_v73_) < (62)):
                        coll45_ = coll45_.union(_dafny.Set([default__.safeDivisionInt(d_195_v73_, 234)]))
                return _dafny.Set(coll45_)
            nw19_[int(27)] = len((iife61_()
            ).intersection(d_188_v67_))
            nw19_[int(28)] = p0
            d_194_v74_ = nw19_
            index14_ = default__.safeIndex(319, (d_194_v74_).length(0))
            (d_194_v74_)[index14_] = p0
            index15_ = default__.safeIndex(319, (d_194_v74_).length(0))
            rhs11_ = p0
            rhs12_ = (p0) * (p3)
            lhs6_ = d_194_v74_
            lhs7_ = default__.safeIndex(319, (d_194_v74_).length(0))
            lhs8_ = globalState
            lhs6_[lhs7_] = rhs11_
            lhs8_.f7 = rhs12_
            d_196_v75_: _dafny.Array
            nw20_ = _dafny.Array(_dafny.Set({}), 28)
            d_196_v75_ = nw20_
            index16_ = default__.safeIndex(809, (d_196_v75_).length(0))
            (d_196_v75_)[index16_] = d_192_v71_
            d_197_v76_: _dafny.Seq
            d_197_v76_ = _dafny.SeqWithoutIsStrInference([d_192_v71_])
            d_198_v77_: str
            d_198_v77_ = _dafny.CodePoint('h')
            d_199_v78_: _dafny.Map
            d_199_v78_ = _dafny.Map({(d_185_v64_)[default__.safeIndex(7, (d_185_v64_).length(0))]: p2})
            index17_ = default__.safeIndex(809, (d_196_v75_).length(0))
            rhs13_ = default__.fm2(globalState)
            rhs14_ = ((d_192_v71_) | (d_192_v71_)).intersection(((d_197_v76_)[default__.safeIndex(default__.fm0(p2, len(_dafny.Set({d_198_v77_})), globalState), len(d_197_v76_))]) - (_dafny.Set({((d_199_v78_)[(d_185_v64_)[default__.safeIndex(7, (d_185_v64_).length(0))]] if ((d_185_v64_)[default__.safeIndex(7, (d_185_v64_).length(0))]) in (d_199_v78_) else (d_185_v64_)[default__.safeIndex(7, (d_185_v64_).length(0))]), p2})))
            lhs9_ = d_196_v75_
            lhs10_ = default__.safeIndex(809, (d_196_v75_).length(0))
            r1 = rhs13_
            lhs9_[lhs10_] = rhs14_
        d_200_v79_: str
        d_200_v79_ = _dafny.CodePoint('a')
        d_201_v80_: C6
        nw21_ = C6()
        nw21_.ctor__(d_200_v79_, True)
        d_201_v80_ = nw21_
        d_202_v81_: D23
        d_202_v81_ = D23_DC54(d_201_v80_)
        d_201_v80_ = (d_202_v81_).cf84
        d_203_v82_: _dafny.Seq
        d_203_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yj"))
        d_204_v83_: D9
        d_204_v83_ = D9_DC20(d_203_v82_, p2, p0)
        d_205_v84_: _dafny.MultiSet
        d_205_v84_ = _dafny.MultiSet([d_204_v83_, d_204_v83_, d_204_v83_])
        d_206_v85_: D21
        d_206_v85_ = D21_DC50(d_205_v84_)
        d_207_v86_: _dafny.Seq
        d_207_v86_ = _dafny.SeqWithoutIsStrInference([d_206_v85_, d_206_v85_, D21_DC50(d_205_v84_)])
        d_207_v86_ = d_207_v86_
        d_208_v87_: _dafny.Array
        def lambda5_(d_209_p2_):
            def lambda6_(d_210_i7_):
                return d_209_p2_

            return lambda6_

        init3_ = lambda5_(p2)
        nw22_ = _dafny.Array(None, 29)
        for i0_3_ in range(nw22_.length(0)):
            nw22_[i0_3_] = init3_(i0_3_)
        d_208_v87_ = nw22_
        d_208_v87_ = d_208_v87_
        index18_ = default__.safeIndex(776, (d_208_v87_).length(0))
        (d_208_v87_)[index18_] = p2
        d_211_v88_: _dafny.Map
        d_211_v88_ = _dafny.Map({(0) - (len(d_203_v82_)): p2})
        d_212_v89_: _dafny.Array
        nw23_ = _dafny.Array(None, 7)
        nw23_[int(0)] = len(d_211_v88_)
        nw23_[int(1)] = p0
        nw23_[int(2)] = p0
        nw23_[int(3)] = p3
        nw23_[int(4)] = len(d_203_v82_)
        nw23_[int(5)] = p3
        nw23_[int(6)] = p3
        d_212_v89_ = nw23_
        d_213_v90_: _dafny.Map
        d_213_v90_ = _dafny.Map({d_212_v89_: p1})
        d_214_v91_: _dafny.Seq
        d_214_v91_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_215_v92_: _dafny.Seq
        d_215_v92_ = _dafny.SeqWithoutIsStrInference([d_214_v91_, d_214_v91_])
        index19_ = default__.safeIndex(776, (d_208_v87_).length(0))
        (d_208_v87_)[index19_] = (((d_213_v90_)[d_212_v89_] if (d_212_v89_) in (d_213_v90_) else p1)).isdisjoint((p1).set(p2, default__.abs(len(d_215_v92_))))
        d_216_v93_: _dafny.Map
        d_216_v93_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nrtq")): d_212_v89_})
        r2 = len((d_216_v93_) | (d_216_v93_))
        r0 = p0
        r1 = (d_208_v87_)[default__.safeIndex(776, (d_208_v87_).length(0))]
        r2 = p0
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_217_v0_: bool
        d_217_v0_ = True
        d_218_v1_: _dafny.Array
        nw24_ = _dafny.Array(None, 1)
        nw24_[int(0)] = d_217_v0_
        d_218_v1_ = nw24_
        d_219_v2_: int
        d_219_v2_ = 230
        d_220_v3_: _dafny.Seq
        d_220_v3_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_218_v1_: (0) - (d_219_v2_)})])
        d_221_v4_: _dafny.Map
        d_221_v4_ = _dafny.Map({298: d_217_v0_})
        d_222_v5_: _dafny.Map
        d_222_v5_ = _dafny.Map({d_219_v2_: d_219_v2_})
        d_223_v6_: _dafny.Map
        d_223_v6_ = _dafny.Map({True: _dafny.CodePoint('t')})
        d_224_v7_: _dafny.Seq
        d_224_v7_ = _dafny.SeqWithoutIsStrInference([d_217_v0_])
        d_225_v8_: _dafny.Map
        d_225_v8_ = _dafny.Map({d_217_v0_: d_217_v0_})
        d_226_v9_: _dafny.Map
        d_226_v9_ = _dafny.Map({d_224_v7_: True})
        d_227_v10_: _dafny.Seq
        d_227_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjoqku"))
        d_228_v11_: _dafny.Array
        nw25_ = _dafny.Array(None, 26)
        nw25_[int(0)] = d_219_v2_
        nw25_[int(1)] = len(d_221_v4_)
        nw25_[int(2)] = d_219_v2_
        nw25_[int(3)] = ((d_222_v5_)[d_219_v2_] if (d_219_v2_) in (d_222_v5_) else d_219_v2_)
        nw25_[int(4)] = d_219_v2_
        nw25_[int(5)] = len(d_223_v6_)
        nw25_[int(6)] = d_219_v2_
        nw25_[int(7)] = d_219_v2_
        nw25_[int(8)] = -846
        nw25_[int(9)] = d_219_v2_
        nw25_[int(10)] = 379
        nw25_[int(11)] = d_219_v2_
        nw25_[int(12)] = d_219_v2_
        nw25_[int(13)] = d_219_v2_
        nw25_[int(14)] = d_219_v2_
        nw25_[int(15)] = (0) - (d_219_v2_)
        nw25_[int(16)] = len((d_224_v7_).set(default__.safeIndex(len(d_224_v7_), len(d_224_v7_)), d_217_v0_))
        nw25_[int(17)] = d_219_v2_
        nw25_[int(18)] = (0) - (d_219_v2_)
        nw25_[int(19)] = d_219_v2_
        nw25_[int(20)] = (0) - (len(d_225_v8_))
        nw25_[int(21)] = len(d_226_v9_)
        nw25_[int(22)] = len(d_227_v10_)
        nw25_[int(23)] = d_219_v2_
        nw25_[int(24)] = len(d_221_v4_)
        nw25_[int(25)] = d_219_v2_
        d_228_v11_ = nw25_
        d_229_v12_: _dafny.Set
        d_229_v12_ = _dafny.Set({d_218_v1_})
        d_230_v13_: str
        d_230_v13_ = _dafny.CodePoint('a')
        d_231_v14_: _dafny.Map
        d_231_v14_ = _dafny.Map({(0) - (d_219_v2_): d_225_v8_})
        d_232_globalState_: GlobalState
        nw26_ = GlobalState()
        nw26_.ctor__(d_218_v1_, 986, (d_220_v3_)[default__.safeIndex(d_219_v2_, len(d_220_v3_))], True, d_228_v11_, False, d_218_v1_, 986, d_227_v10_, True, d_229_v12_, 737, (d_227_v10_).set(default__.safeIndex(d_219_v2_, len(d_227_v10_)), d_230_v13_), False, _dafny.CodePoint('a'), (d_225_v8_) | (((d_231_v14_)[d_219_v2_] if (d_219_v2_) in (d_231_v14_) else d_225_v8_)), 108, (d_227_v10_) + (d_227_v10_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")), False, -137, d_228_v11_, d_228_v11_)
        d_232_globalState_ = nw26_
        hi0_ = (0) - (((0) - (d_219_v2_) if True else default__.fm0(d_217_v0_, d_219_v2_, d_232_globalState_)))
        for d_233_i0_ in range(425, hi0_):
            if True:
                (d_232_globalState_).f7 = d_219_v2_
                (d_232_globalState_).f7 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sulow")))
                d_227_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_234_i1_ in range(default__.abs(723))])
                (d_232_globalState_).f7 = d_219_v2_
                d_235_v15_: _dafny.MultiSet
                d_235_v15_ = _dafny.MultiSet([d_217_v0_])
                d_236_v16_: int
                d_237_v17_: bool
                d_238_v18_: int
                out0_: int
                out1_: bool
                out2_: int
                out0_, out1_, out2_ = default__.m0((0) - (d_233_i0_), (d_235_v15_).intersection(d_235_v15_), d_217_v0_, len(d_224_v7_), d_232_globalState_)
                d_236_v16_ = out0_
                d_237_v17_ = out1_
                d_238_v18_ = out2_
            elif True:
                index20_ = default__.safeIndex(240, (d_218_v1_).length(0))
                (d_218_v1_)[index20_] = d_217_v0_
                d_239_v19_: _dafny.Seq
                d_239_v19_ = _dafny.SeqWithoutIsStrInference([313])
                d_240_v20_: _dafny.Seq
                d_240_v20_ = _dafny.SeqWithoutIsStrInference([d_239_v19_, d_239_v19_, d_239_v19_])
                index21_ = default__.safeIndex(240, (d_218_v1_).length(0))
                (d_218_v1_)[index21_] = ((d_240_v20_)[default__.safeIndex((0) - (d_219_v2_), len(d_240_v20_))]) == (d_239_v19_)
                index22_ = default__.safeIndex(240, (d_218_v1_).length(0))
                (d_218_v1_)[index22_] = ((len(default__.fm1(default__.fm2(d_232_globalState_), d_232_globalState_))) + (d_219_v2_)) <= (591)
                d_241_v21_: _dafny.MultiSet
                d_241_v21_ = _dafny.MultiSet([d_217_v0_, not(d_217_v0_), False, d_217_v0_, d_217_v0_])
                d_242_v22_: int
                d_243_v23_: bool
                d_244_v24_: int
                out3_: int
                out4_: bool
                out5_: int
                out3_, out4_, out5_ = default__.m0(d_219_v2_, d_241_v21_, (d_218_v1_)[default__.safeIndex(240, (d_218_v1_).length(0))], d_233_i0_, d_232_globalState_)
                d_242_v22_ = out3_
                d_243_v23_ = out4_
                d_244_v24_ = out5_
                d_245_v25_: D0
                d_245_v25_ = D0_DC0(default__.fm3(d_217_v0_, d_232_globalState_))
                d_246_v26_: _dafny.Set
                d_246_v26_ = _dafny.Set({d_244_v24_})
                d_247_v27_: int
                d_248_v28_: bool
                d_249_v29_: int
                out6_: int
                out7_: bool
                out8_: int
                out6_, out7_, out8_ = default__.m0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrbqh"))), ((d_245_v25_).cf0) - (d_241_v21_), d_217_v0_, (d_244_v24_) * (len(d_246_v26_)), d_232_globalState_)
                d_247_v27_ = out6_
                d_248_v28_ = out7_
                d_249_v29_ = out8_
                (d_232_globalState_).f7 = d_219_v2_
            d_250_v30_: _dafny.Array
            def lambda7_(d_251_i2_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "feg"))

            init4_ = lambda7_
            nw27_ = _dafny.Array(None, 27)
            for i0_4_ in range(nw27_.length(0)):
                nw27_[i0_4_] = init4_(i0_4_)
            d_250_v30_ = nw27_
            index23_ = default__.safeIndex(561, (d_250_v30_).length(0))
            (d_250_v30_)[index23_] = (d_227_v10_) + (d_227_v10_)
            d_252_v31_: _dafny.MultiSet
            d_252_v31_ = _dafny.MultiSet([d_217_v0_, d_217_v0_, d_217_v0_])
            index24_ = default__.safeIndex(561, (d_250_v30_).length(0))
            (d_250_v30_)[index24_] = default__.fm1((d_219_v2_) <= (((d_252_v31_)[not(d_217_v0_)] if (not(d_217_v0_)) in (d_252_v31_) else d_233_i0_)), d_232_globalState_)
            d_253_v32_: _dafny.Array
            nw28_ = _dafny.Array(D0.default()(), 5)
            d_253_v32_ = nw28_
            index25_ = default__.safeIndex(95, (d_253_v32_).length(0))
            (d_253_v32_)[index25_] = D0_DC2(d_217_v0_, d_233_i0_, d_217_v0_)
            d_254_v34_: _dafny.Set
            d_254_v34_ = _dafny.Set({d_219_v2_})
            d_255_v35_: D0
            def iife62_():
                coll46_ = _dafny.Map()
                compr_46_: int
                for compr_46_ in (d_254_v34_).Elements:
                    d_256_v33_: int = compr_46_
                    if (d_256_v33_) in (d_254_v34_):
                        coll46_[default__.safeDivisionInt(d_256_v33_, d_233_i0_)] = d_230_v13_
                return _dafny.Map(coll46_)
            d_255_v35_ = D0_DC2(not(d_217_v0_), len(iife62_()
), d_217_v0_)
            pat_let_tv5_ = d_217_v0_
            index26_ = default__.safeIndex(95, (d_253_v32_).length(0))
            def iife63_(_pat_let8_0):
                def iife64_(d_257_dt__update__tmp_h0_):
                    def iife65_(_pat_let9_0):
                        def iife66_(d_258_dt__update_hcf2_h0_):
                            return D0_DC2(d_258_dt__update_hcf2_h0_, (d_257_dt__update__tmp_h0_).cf3, (d_257_dt__update__tmp_h0_).cf4)
                        return iife66_(_pat_let9_0)
                    return iife65_(not(pat_let_tv5_))
                return iife64_(_pat_let8_0)
            (d_253_v32_)[index26_] = iife63_(d_255_v35_)
            rhs15_ = d_233_i0_
            rhs16_ = default__.fm0(False, d_233_i0_, d_232_globalState_)
            rhs17_ = ((_dafny.MultiSet([d_217_v0_, d_217_v0_])) | (_dafny.MultiSet([d_217_v0_, d_217_v0_]))).intersection(_dafny.MultiSet([d_217_v0_, not(d_217_v0_), True]))
            rhs18_ = d_217_v0_
            lhs11_ = d_232_globalState_
            lhs12_ = d_232_globalState_
            lhs11_.f7 = rhs15_
            d_219_v2_ = rhs16_
            d_252_v31_ = rhs17_
            lhs12_.f13 = rhs18_
        (d_232_globalState_).f13 = not ((d_219_v2_) >= (d_219_v2_)) or (not((d_227_v10_) <= (_dafny.SeqWithoutIsStrInference([d_230_v13_ for d_259_i3_ in range(default__.abs(215))]))))
        d_260_v36_: _dafny.Set
        d_260_v36_ = _dafny.Set({d_217_v0_})
        d_261_v37_: _dafny.Set
        d_261_v37_ = _dafny.Set({default__.fm0(d_217_v0_, d_219_v2_, d_232_globalState_), d_219_v2_})
        d_262_v38_: _dafny.Set
        d_262_v38_ = _dafny.Set({462, len(d_261_v37_), d_219_v2_, d_219_v2_, d_219_v2_})
        d_219_v2_ = (d_219_v2_) - (default__.safeModuloInt(len((d_222_v5_).set(len(d_260_v36_), len(d_261_v37_))), default__.fm0(d_217_v0_, len(_dafny.SeqWithoutIsStrInference([d_217_v0_])), d_232_globalState_)))
        d_263_v39_: _dafny.MultiSet
        d_263_v39_ = _dafny.MultiSet([(0) - (default__.safeModuloInt(d_219_v2_, (0) - (d_219_v2_)))])
        d_263_v39_ = (d_263_v39_).intersection(d_263_v39_)
        index27_ = default__.safeIndex(95, (d_228_v11_).length(0))
        (d_228_v11_)[index27_] = d_219_v2_
        index28_ = default__.safeIndex(95, (d_228_v11_).length(0))
        (d_228_v11_)[index28_] = d_219_v2_
        d_264_v40_: _dafny.Array
        nw29_ = _dafny.Array(_dafny.MultiSet({}), 22)
        d_264_v40_ = nw29_
        d_265_v41_: _dafny.MultiSet
        d_265_v41_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpbndwsy")), default__.fm1(d_217_v0_, d_232_globalState_), d_227_v10_, d_227_v10_])
        index29_ = default__.safeIndex(357, (d_264_v40_).length(0))
        (d_264_v40_)[index29_] = d_265_v41_
        d_266_v42_: D0
        d_266_v42_ = D0_DC1(d_219_v2_)
        pat_let_tv6_ = d_265_v41_
        pat_let_tv7_ = d_227_v10_
        pat_let_tv8_ = d_224_v7_
        pat_let_tv9_ = d_265_v41_
        pat_let_tv10_ = d_265_v41_
        index30_ = default__.safeIndex(357, (d_264_v40_).length(0))
        def lambda8_(source5_):
            if source5_.is_DC1:
                d_267___mcc_h0_ = source5_.cf1
                d_268_cf1_ = d_267___mcc_h0_
                return (pat_let_tv6_).set(pat_let_tv7_, default__.abs(len(pat_let_tv8_)))
            elif source5_.is_DC2:
                d_269___mcc_h1_ = source5_.cf2
                d_270___mcc_h2_ = source5_.cf3
                d_271___mcc_h3_ = source5_.cf4
                d_272_cf4_ = d_271___mcc_h3_
                d_273_cf3_ = d_270___mcc_h2_
                d_274_cf2_ = d_269___mcc_h1_
                return pat_let_tv9_
            elif source5_.is_DC0:
                d_275___mcc_h4_ = source5_.cf0
                d_276_cf0_ = d_275___mcc_h4_
                return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nt"))])
            elif True:
                d_277___mcc_h5_ = source5_.cf5
                d_278_cf5_ = d_277___mcc_h5_
                return pat_let_tv10_

        rhs19_ = lambda8_(d_266_v42_)
        rhs20_ = ((0) - ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))])) == (default__.safeModuloInt((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], d_219_v2_))
        lhs13_ = d_264_v40_
        lhs14_ = default__.safeIndex(357, (d_264_v40_).length(0))
        lhs15_ = d_232_globalState_
        lhs13_[lhs14_] = rhs19_
        lhs15_.f13 = rhs20_
        d_279_v43_: _dafny.Map
        d_279_v43_ = _dafny.Map({d_218_v1_: (d_227_v10_)[default__.safeIndex(d_219_v2_, len(d_227_v10_))]})
        d_280_v44_: _dafny.Array
        def lambda9_(d_281_v13_):
            def lambda10_(d_282_i4_):
                return d_281_v13_

            return lambda10_

        init5_ = lambda9_(d_230_v13_)
        nw30_ = _dafny.Array(None, 5)
        for i0_5_ in range(nw30_.length(0)):
            nw30_[i0_5_] = init5_(i0_5_)
        d_280_v44_ = nw30_
        d_283_v45_: C3
        nw31_ = C3()
        nw31_.ctor__(d_279_v43_, d_280_v44_, default__.fm29((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], d_217_v0_, d_217_v0_, (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], d_232_globalState_), ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]) != ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]))
        d_283_v45_ = nw31_
        d_284_v46_: D5
        d_284_v46_ = D5_DC14(d_217_v0_)
        d_285_v47_: _dafny.Set
        def iife67_(_pat_let10_0):
            def iife68_(d_286_dt__update__tmp_h1_):
                def iife69_(_pat_let11_0):
                    def iife70_(d_287_dt__update_hcf23_h0_):
                        return D5_DC14(d_287_dt__update_hcf23_h0_)
                    return iife70_(_pat_let11_0)
                return iife69_(True)
            return iife68_(_pat_let10_0)
        d_285_v47_ = _dafny.Set({d_284_v46_, d_284_v46_, D5_DC14(d_217_v0_), iife67_(D5_DC14(d_217_v0_)), d_284_v46_})
        d_288_v49_: _dafny.Map
        def iife71_():
            coll47_ = _dafny.Map()
            compr_47_: int
            for compr_47_ in _dafny.IntegerRange(568, 81):
                d_289_v48_: int = compr_47_
                if ((568) <= (d_289_v48_)) and ((d_289_v48_) < (81)):
                    coll47_[default__.safeModuloInt(d_289_v48_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vlhwkla"))))] = ((d_263_v39_)[395] if (395) in (d_263_v39_) else (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))])
            return _dafny.Map(coll47_)
        d_288_v49_ = _dafny.Map({d_285_v47_: iife71_()
        })
        d_290_v52_: _dafny.Seq
        def iife72_():
            def iife74_():
                coll50_ = _dafny.Set()
                compr_50_: D5
                for compr_50_ in (default__.fm58(d_230_v13_, len(d_227_v10_), d_217_v0_, d_232_globalState_)).keys.Elements:
                    d_293_v51_: D5 = compr_50_
                    if (d_293_v51_) in (default__.fm58(d_230_v13_, len(d_227_v10_), d_217_v0_, d_232_globalState_)):
                        coll50_ = coll50_.union(_dafny.Set([d_293_v51_]))
                return _dafny.Set(coll50_)
            coll48_ = _dafny.Map()
            def iife73_():
                coll49_ = _dafny.Set()
                compr_49_: D5
                for compr_49_ in (default__.fm58(d_230_v13_, len(d_227_v10_), d_217_v0_, d_232_globalState_)).keys.Elements:
                    d_291_v51_: D5 = compr_49_
                    if (d_291_v51_) in (default__.fm58(d_230_v13_, len(d_227_v10_), d_217_v0_, d_232_globalState_)):
                        coll49_ = coll49_.union(_dafny.Set([d_291_v51_]))
                return _dafny.Set(coll49_)
            compr_48_: _dafny.Set
            for compr_48_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({d_284_v46_}), d_285_v47_, iife73_()
            ])).Elements:
                d_292_v50_: _dafny.Set = compr_48_
                if (d_292_v50_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({d_284_v46_}), d_285_v47_, iife74_()
                ])):
                    coll48_[d_292_v50_] = d_222_v5_
            return _dafny.Map(coll48_)
        d_290_v52_ = _dafny.SeqWithoutIsStrInference([d_288_v49_, iife72_()
        , (d_288_v49_) | (d_288_v49_), default__.fm59(d_227_v10_, d_232_globalState_), d_288_v49_])
        d_288_v49_ = (d_290_v52_)[default__.safeIndex(-874, len(d_290_v52_))]
        d_294_v53_: _dafny.Array
        nw32_ = _dafny.Array(_dafny.Seq({}), 2)
        d_294_v53_ = nw32_
        index31_ = default__.safeIndex(669, (d_294_v53_).length(0))
        (d_294_v53_)[index31_] = _dafny.SeqWithoutIsStrInference([d_218_v1_])
        d_295_v54_: _dafny.Seq
        d_295_v54_ = _dafny.SeqWithoutIsStrInference([d_218_v1_, d_218_v1_, d_218_v1_, d_218_v1_, d_218_v1_])
        d_296_v55_: _dafny.Seq
        d_296_v55_ = _dafny.SeqWithoutIsStrInference([d_295_v54_, _dafny.SeqWithoutIsStrInference([d_218_v1_])])
        index32_ = default__.safeIndex(669, (d_294_v53_).length(0))
        (d_294_v53_)[index32_] = (d_296_v55_)[default__.safeIndex((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], len(d_296_v55_))]
        hi1_ = d_219_v2_
        for d_297_i5_ in range(default__.safeDivisionInt(d_219_v2_, d_219_v2_), hi1_):
            d_298_v56_: _dafny.Array
            def lambda11_(d_299_v0_, d_300_v2_):
                def lambda12_(d_301_i6_):
                    return _dafny.Map({d_299_v0_: d_300_v2_})

                return lambda12_

            init6_ = lambda11_(d_217_v0_, d_219_v2_)
            nw33_ = _dafny.Array(None, 7)
            for i0_6_ in range(nw33_.length(0)):
                nw33_[i0_6_] = init6_(i0_6_)
            d_298_v56_ = nw33_
            d_302_v57_: C1
            nw34_ = C1()
            nw34_.ctor__(True, False)
            d_302_v57_ = nw34_
            d_303_v58_: _dafny.Map
            d_303_v58_ = _dafny.Map({d_298_v56_: d_302_v57_})
            d_304_v59_: _dafny.Seq
            d_304_v59_ = _dafny.SeqWithoutIsStrInference([((d_303_v58_)[d_298_v56_] if (d_298_v56_) in (d_303_v58_) else d_302_v57_), (d_302_v57_ if (d_302_v57_).f39 else d_302_v57_)])
            rhs21_ = d_228_v11_
            rhs22_ = (d_304_v59_) + (d_304_v59_)
            rhs23_ = d_218_v1_
            lhs16_ = d_232_globalState_
            lhs16_.f22 = rhs21_
            d_304_v59_ = rhs22_
            d_218_v1_ = rhs23_
            (d_232_globalState_).f13 = (d_224_v7_)[default__.safeIndex(default__.fm0((d_302_v57_).f39, len(d_222_v5_), d_232_globalState_), len(d_224_v7_))]
            if (62) == ((d_297_i5_) - ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))])):
                d_305_v60_: C10
                nw35_ = C10()
                nw35_.ctor__((True) == ((d_302_v57_).f39), (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], _dafny.CodePoint('j'), (d_219_v2_) <= (d_219_v2_), d_280_v44_, d_217_v0_)
                d_305_v60_ = nw35_
                d_306_v61_: _dafny.Seq
                d_306_v61_ = _dafny.SeqWithoutIsStrInference([(d_305_v60_.f29 if (d_302_v57_).f39 else d_305_v60_.f29)])
                rhs24_ = d_305_v60_
                rhs25_ = (((_dafny.SeqWithoutIsStrInference([d_305_v60_.f29 for d_307_i7_ in range(default__.abs(238))]) if d_217_v0_ else d_306_v61_)) + (d_306_v61_)).set(default__.safeIndex(len(d_227_v10_), len(((_dafny.SeqWithoutIsStrInference([d_305_v60_.f29 for d_308_i7_ in range(default__.abs(238))]) if d_217_v0_ else d_306_v61_)) + (d_306_v61_))), d_297_i5_)
                rhs26_ = (d_306_v61_) + (d_306_v61_)
                rhs27_ = d_219_v2_
                rhs28_ = d_297_i5_
                lhs17_ = d_232_globalState_
                lhs18_ = d_232_globalState_
                d_305_v60_ = rhs24_
                d_306_v61_ = rhs25_
                d_306_v61_ = rhs26_
                lhs17_.f7 = rhs27_
                lhs18_.f7 = rhs28_
                d_309_v62_: _dafny.MultiSet
                d_309_v62_ = _dafny.MultiSet([d_217_v0_, (d_302_v57_).f39, (d_302_v57_).f39])
                d_310_v63_: C2
                nw36_ = C2()
                nw36_.ctor__(d_309_v62_, d_280_v44_, d_230_v13_, (d_305_v60_).f28)
                d_310_v63_ = nw36_
                (d_232_globalState_).f13 = d_217_v0_
                d_311_v64_: C0
                nw37_ = C0()
                nw37_.ctor__(_dafny.Map({(d_305_v60_).f28: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))}), ((d_302_v57_).f39) and (d_217_v0_), _dafny.CodePoint('l'), (d_305_v60_).f28)
                d_311_v64_ = nw37_
                index33_ = default__.safeIndex(95, (d_228_v11_).length(0))
                rhs29_ = d_311_v64_
                rhs30_ = d_309_v62_
                rhs31_ = default__.safeModuloInt(d_219_v2_, d_305_v60_.f29)
                rhs32_ = d_283_v45_
                rhs33_ = ((default__.fm3(not((d_302_v57_).f39), d_232_globalState_)).cardinality if d_217_v0_ else (d_305_v60_).fm9(d_306_v61_, d_305_v60_.f29, (d_302_v57_).f39, d_232_globalState_))
                lhs19_ = d_232_globalState_
                lhs20_ = d_228_v11_
                lhs21_ = default__.safeIndex(95, (d_228_v11_).length(0))
                d_311_v64_ = rhs29_
                d_309_v62_ = rhs30_
                lhs19_.f7 = rhs31_
                d_283_v45_ = rhs32_
                lhs20_[lhs21_] = rhs33_
                d_217_v0_ = (d_311_v64_).f38
            elif True:
                d_312_v65_: _dafny.Map
                d_312_v65_ = _dafny.Map({(d_302_v57_).f39: default__.fm1((d_302_v57_).f39, d_232_globalState_)})
                d_313_v66_: T2
                nw38_ = C2()
                nw38_.ctor__(_dafny.MultiSet(d_224_v7_), d_280_v44_, d_230_v13_, (d_302_v57_).f39)
                d_313_v66_ = nw38_
                d_314_v67_: _dafny.Map
                d_314_v67_ = _dafny.Map({((d_312_v65_)[(d_302_v57_).f39] if ((d_302_v57_).f39) in (d_312_v65_) else d_227_v10_): (d_313_v66_ if d_217_v0_ else d_313_v66_)})
                d_314_v67_ = (d_314_v67_).set(d_227_v10_, d_313_v66_)
                d_315_v68_: _dafny.Seq
                d_315_v68_ = _dafny.SeqWithoutIsStrInference([d_297_i5_, (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]])
                d_316_v69_: _dafny.Map
                d_316_v69_ = _dafny.Map({(d_313_v66_).f24: _dafny.Map({len(d_222_v5_): (d_313_v66_).f24})})
                d_317_v71_: _dafny.Seq
                def iife75_():
                    coll51_ = _dafny.Map()
                    compr_51_: int
                    for compr_51_ in _dafny.IntegerRange(-632, 635):
                        d_319_v70_: int = compr_51_
                        if ((-632) <= (d_319_v70_)) and ((d_319_v70_) < (635)):
                            coll51_[(d_319_v70_) + (len(d_221_v4_))] = (D6_DC16(d_225_v8_, (d_302_v57_).f39, False)).cf27
                    return _dafny.Map(coll51_)
                d_317_v71_ = _dafny.SeqWithoutIsStrInference([((_dafny.SeqWithoutIsStrInference([d_297_i5_ for d_318_i8_ in range(default__.abs(933))])).set(default__.safeIndex(len(((d_316_v69_)[(d_302_v57_).f39] if ((d_302_v57_).f39) in (d_316_v69_) else iife75_()
                )), len(_dafny.SeqWithoutIsStrInference([d_297_i5_ for d_320_i8_ in range(default__.abs(933))]))), len(d_227_v10_))) + (d_315_v68_), d_315_v68_, _dafny.SeqWithoutIsStrInference([(d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))] for d_321_i9_ in range(default__.abs(930))]), d_315_v68_, d_315_v68_])
                d_315_v68_ = (d_317_v71_)[default__.safeIndex(((d_302_v57_).fm33(d_232_globalState_)) - (d_297_i5_), len(d_317_v71_))]
                d_322_v72_: bool
                d_323_v73_: str
                d_324_v74_: bool
                out9_: bool
                out10_: str
                out11_: bool
                out9_, out10_, out11_ = (d_283_v45_).m11(d_217_v0_, d_232_globalState_)
                d_322_v72_ = out9_
                d_323_v73_ = out10_
                d_324_v74_ = out11_
                d_325_v75_: T1
                nw39_ = C9()
                nw39_.ctor__(d_230_v13_, d_324_v74_)
                d_325_v75_ = nw39_
                d_325_v75_ = d_325_v75_
                d_317_v71_ = (default__.fm60((d_302_v57_).f39, d_232_globalState_) if d_322_v72_ else d_317_v71_)
            d_230_v13_ = d_230_v13_
        d_326_v76_: D21
        d_326_v76_ = D21_DC51((d_217_v0_) or (d_217_v0_))
        source6_ = d_326_v76_
        if source6_.is_DC51:
            d_327___mcc_h6_ = source6_.cf81
            d_328_cf81_ = d_327___mcc_h6_
            (d_232_globalState_).f13 = not(d_217_v0_)
            if ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]) >= (-409):
                d_329_v77_: C1
                nw40_ = C1()
                nw40_.ctor__(d_217_v0_, d_328_cf81_)
                d_329_v77_ = nw40_
                d_330_v78_: _dafny.MultiSet
                d_330_v78_ = _dafny.MultiSet([(d_329_v77_).fm32(d_328_cf81_, d_219_v2_, d_232_globalState_), (d_224_v7_)[default__.safeIndex((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], len(d_224_v7_))]])
                d_331_v79_: int
                d_332_v80_: bool
                d_333_v81_: int
                out12_: int
                out13_: bool
                out14_: int
                out12_, out13_, out14_ = default__.m0((len(_dafny.Map({d_329_v77_: _dafny.CodePoint('p')}))) + ((d_263_v39_).cardinality), d_330_v78_, d_217_v0_, d_219_v2_, d_232_globalState_)
                d_331_v79_ = out12_
                d_332_v80_ = out13_
                d_333_v81_ = out14_
                d_334_v82_: _dafny.Array
                nw41_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_334_v82_ = nw41_
                index34_ = default__.safeIndex(259, (d_334_v82_).length(0))
                (d_334_v82_)[index34_] = d_228_v11_
                index35_ = default__.safeIndex(259, (d_334_v82_).length(0))
                index36_ = default__.safeIndex(95, (d_228_v11_).length(0))
                rhs34_ = (d_221_v4_) | (d_221_v4_)
                rhs35_ = d_332_v80_
                rhs36_ = d_228_v11_
                rhs37_ = default__.safeDivisionInt(d_219_v2_, d_219_v2_)
                rhs38_ = d_332_v80_
                lhs22_ = d_334_v82_
                lhs23_ = default__.safeIndex(259, (d_334_v82_).length(0))
                lhs24_ = d_228_v11_
                lhs25_ = default__.safeIndex(95, (d_228_v11_).length(0))
                d_221_v4_ = rhs34_
                d_332_v80_ = rhs35_
                lhs22_[lhs23_] = rhs36_
                lhs24_[lhs25_] = rhs37_
                d_217_v0_ = rhs38_
                d_335_v83_: C2
                nw42_ = C2()
                nw42_.ctor__(d_330_v78_, d_280_v44_, (d_230_v13_ if (d_329_v77_).f39 else d_230_v13_), d_332_v80_)
                d_335_v83_ = nw42_
                d_328_cf81_ = d_332_v80_
                d_336_v84_: C4
                nw43_ = C4()
                nw43_.ctor__((d_335_v83_).fm24((d_335_v83_).fm24(d_333_v81_, d_328_cf81_, d_232_globalState_), (d_329_v77_).f39, d_232_globalState_), d_222_v5_, d_280_v44_, d_230_v13_, ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]) >= (d_331_v79_))
                d_336_v84_ = nw43_
                nw44_ = C4()
                nw44_.ctor__(d_331_v79_, d_336_v84_.f34, d_280_v44_, d_230_v13_, d_217_v0_)
                d_336_v84_ = nw44_
            elif True:
                d_337_v85_: bool
                d_338_v86_: str
                d_339_v87_: bool
                out15_: bool
                out16_: str
                out17_: bool
                out15_, out16_, out17_ = (d_283_v45_).m11(False, d_232_globalState_)
                d_337_v85_ = out15_
                d_338_v86_ = out16_
                d_339_v87_ = out17_
                d_340_v88_: C5
                nw45_ = C5()
                nw45_.ctor__(d_218_v1_)
                d_340_v88_ = nw45_
                d_341_v89_: _dafny.Array
                nw46_ = _dafny.Array(None, 12)
                nw46_[int(0)] = d_340_v88_
                nw46_[int(1)] = d_340_v88_
                nw46_[int(2)] = d_340_v88_
                nw46_[int(3)] = d_340_v88_
                nw46_[int(4)] = d_340_v88_
                nw46_[int(5)] = d_340_v88_
                nw46_[int(6)] = d_340_v88_
                nw46_[int(7)] = d_340_v88_
                nw46_[int(8)] = d_340_v88_
                nw46_[int(9)] = d_340_v88_
                nw46_[int(10)] = d_340_v88_
                nw46_[int(11)] = d_340_v88_
                d_341_v89_ = nw46_
                d_342_v90_: D16
                d_342_v90_ = D16_DC37(d_340_v88_)
                index37_ = default__.safeIndex(403, (d_341_v89_).length(0))
                (d_341_v89_)[index37_] = (d_342_v90_).cf61
                index38_ = default__.safeIndex(403, (d_341_v89_).length(0))
                nw47_ = C5()
                nw47_.ctor__(d_218_v1_)
                (d_341_v89_)[index38_] = nw47_
                d_227_v10_ = d_227_v10_
                d_343_v91_: _dafny.MultiSet
                d_343_v91_ = _dafny.MultiSet([(-238) == (d_219_v2_), d_328_cf81_])
                d_344_v94_: C4
                nw48_ = C4()
                nw48_.ctor__(d_219_v2_, d_222_v5_, d_280_v44_, d_338_v86_, d_328_cf81_)
                d_344_v94_ = nw48_
                d_345_v95_: D24
                d_345_v95_ = D24_DC56(d_344_v94_)
                d_346_v96_: _dafny.Map
                def iife76_():
                    coll52_ = _dafny.Set()
                    compr_52_: int
                    for compr_52_ in _dafny.IntegerRange(-204, 551):
                        d_347_v93_: int = compr_52_
                        if ((-204) <= (d_347_v93_)) and ((d_347_v93_) < (551)):
                            coll52_ = coll52_.union(_dafny.Set([(d_347_v93_) * ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))])]))
                    return _dafny.Set(coll52_)
                d_346_v96_ = _dafny.Map({iife76_()
                : (d_345_v95_).cf90})
                def iife77_():
                    coll53_ = _dafny.Set()
                    compr_53_: int
                    for compr_53_ in _dafny.IntegerRange(609, -929):
                        d_348_v92_: int = compr_53_
                        if ((609) <= (d_348_v92_)) and ((d_348_v92_) < (-929)):
                            coll53_ = coll53_.union(_dafny.Set([(d_348_v92_) - ((0) - (d_219_v2_))]))
                    return _dafny.Set(coll53_)
                rhs39_ = default__.safeModuloInt(d_219_v2_, ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]) - ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]))
                rhs40_ = default__.fm3(d_339_v87_, d_232_globalState_)
                rhs41_ = (iife77_()
                ) in (d_346_v96_)
                lhs26_ = d_232_globalState_
                lhs26_.f7 = rhs39_
                d_343_v91_ = rhs40_
                d_328_cf81_ = rhs41_
                d_349_v97_: _dafny.Map
                d_349_v97_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")): d_228_v11_})
                d_350_v98_: _dafny.Array
                nw49_ = _dafny.Array(None, 4)
                nw49_[int(0)] = d_228_v11_
                nw49_[int(1)] = (d_228_v11_ if d_339_v87_ else d_228_v11_)
                nw49_[int(2)] = ((d_349_v97_)[(d_227_v10_).set(default__.safeIndex(default__.fm0(d_339_v87_, (d_344_v94_).f33, d_232_globalState_), len(d_227_v10_)), d_230_v13_)] if ((d_227_v10_).set(default__.safeIndex(default__.fm0(d_339_v87_, (d_344_v94_).f33, d_232_globalState_), len(d_227_v10_)), d_230_v13_)) in (d_349_v97_) else d_228_v11_)
                nw49_[int(3)] = d_228_v11_
                d_350_v98_ = nw49_
                d_351_v99_: D25
                d_351_v99_ = D25_DC58(d_350_v98_)
                d_350_v98_ = (d_351_v99_).cf96
            d_352_v100_: _dafny.Set
            d_352_v100_ = _dafny.Set({d_280_v44_, d_280_v44_})
            d_353_v101_: _dafny.Map
            d_353_v101_ = _dafny.Map({d_217_v0_: d_352_v100_})
            if (len(((d_353_v101_)[d_328_cf81_] if (d_328_cf81_) in (d_353_v101_) else d_352_v100_))) >= (default__.safeDivisionInt((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))])):
                index39_ = default__.safeIndex(777, (d_218_v1_).length(0))
                (d_218_v1_)[index39_] = d_217_v0_
                index40_ = default__.safeIndex(777, (d_218_v1_).length(0))
                index41_ = default__.safeIndex(95, (d_228_v11_).length(0))
                rhs42_ = d_230_v13_
                rhs43_ = d_217_v0_
                rhs44_ = (_dafny.SeqWithoutIsStrInference([d_230_v13_ for d_354_i10_ in range(default__.abs(793))])) != (d_227_v10_)
                rhs45_ = not(d_328_cf81_)
                rhs46_ = (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]
                lhs27_ = d_232_globalState_
                lhs28_ = d_218_v1_
                lhs29_ = default__.safeIndex(777, (d_218_v1_).length(0))
                lhs30_ = d_232_globalState_
                lhs31_ = d_228_v11_
                lhs32_ = default__.safeIndex(95, (d_228_v11_).length(0))
                d_230_v13_ = rhs42_
                lhs27_.f13 = rhs43_
                lhs28_[lhs29_] = rhs44_
                lhs30_.f13 = rhs45_
                lhs31_[lhs32_] = rhs46_
                d_355_v102_: _dafny.MultiSet
                d_355_v102_ = d_263_v39_
                d_356_v103_: _dafny.Seq
                d_356_v103_ = _dafny.SeqWithoutIsStrInference([(d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]])
                nw50_ = _dafny.Array(None, 27)
                nw50_[int(0)] = d_219_v2_
                nw50_[int(1)] = d_219_v2_
                nw50_[int(2)] = default__.safeModuloInt(d_219_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fghnvwo"))))
                nw50_[int(3)] = (d_266_v42_).cf1
                nw50_[int(4)] = ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]) * ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))])
                nw50_[int(5)] = 946
                nw50_[int(6)] = default__.safeDivisionInt(d_219_v2_, (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))])
                nw50_[int(7)] = (0) - (len((d_224_v7_).set(default__.safeIndex(480, len(d_224_v7_)), (d_218_v1_)[default__.safeIndex(777, (d_218_v1_).length(0))])))
                nw50_[int(8)] = (d_219_v2_) + (d_219_v2_)
                nw50_[int(9)] = (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]
                nw50_[int(10)] = d_219_v2_
                nw50_[int(11)] = (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]
                nw50_[int(12)] = default__.fm0((d_218_v1_)[default__.safeIndex(777, (d_218_v1_).length(0))], d_219_v2_, d_232_globalState_)
                nw50_[int(13)] = len(default__.fm21(d_260_v36_, (d_263_v39_), d_232_globalState_))
                nw50_[int(14)] = default__.safeDivisionInt(d_219_v2_, ((d_222_v5_)[(d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]] if ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]) in (d_222_v5_) else d_219_v2_))
                nw50_[int(15)] = (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]
                nw50_[int(16)] = (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]
                nw50_[int(17)] = 86
                nw50_[int(18)] = (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]
                nw50_[int(19)] = d_219_v2_
                nw50_[int(20)] = (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]
                nw50_[int(21)] = d_219_v2_
                nw50_[int(22)] = d_219_v2_
                nw50_[int(23)] = (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]
                nw50_[int(24)] = default__.fm0(d_217_v0_, len(d_221_v4_), d_232_globalState_)
                nw50_[int(25)] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([d_230_v13_ for d_357_i11_ in range(default__.abs(284))])), (d_356_v103_)[default__.safeIndex((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], len(d_356_v103_))])
                nw50_[int(26)] = d_219_v2_
                d_228_v11_ = nw50_
                (d_232_globalState_).f7 = default__.safeModuloInt(((d_356_v103_)[default__.safeIndex(d_219_v2_, len(d_356_v103_))]) + ((0) - ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))])), -101)
                index42_ = default__.safeIndex(95, (d_228_v11_).length(0))
                (d_228_v11_)[index42_] = default__.safeDivisionInt((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], (0) - (((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]) - (d_219_v2_)))
                (d_232_globalState_).f13 = d_328_cf81_
            elif True:
                d_358_v104_: _dafny.Seq
                d_358_v104_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], d_219_v2_)])
                d_359_v105_: _dafny.Set
                d_359_v105_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_219_v2_ for d_360_i12_ in range(default__.abs(865))])})
                d_361_v107_: _dafny.Map
                d_361_v107_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([672 for d_362_i13_ in range(default__.abs(715))]): d_328_cf81_})
                d_363_v109_: _dafny.Seq
                d_363_v109_ = _dafny.SeqWithoutIsStrInference([d_359_v105_])
                d_364_v110_: _dafny.Map
                d_364_v110_ = _dafny.Map({d_358_v104_: d_230_v13_})
                d_365_v112_: D13
                d_365_v112_ = D13_DC28(d_328_cf81_, (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], d_219_v2_)
                d_366_v114_: _dafny.Map
                d_366_v114_ = _dafny.Map({d_217_v0_: d_227_v10_})
                d_367_v115_: D2
                d_367_v115_ = D2_DC6(d_366_v114_)
                d_368_v116_: _dafny.Array
                nw51_ = _dafny.Array(None, 22)
                def iife78_():
                    coll54_ = _dafny.Set()
                    compr_54_: _dafny.Seq
                    for compr_54_ in (d_359_v105_).Elements:
                        d_369_v106_: _dafny.Seq = compr_54_
                        if (d_369_v106_) in (d_359_v105_):
                            coll54_ = coll54_.union(_dafny.Set([d_369_v106_]))
                    return _dafny.Set(coll54_)
                nw51_[int(0)] = iife78_()
                
                nw51_[int(1)] = (d_359_v105_).intersection(d_359_v105_)
                nw51_[int(2)] = d_359_v105_
                nw51_[int(3)] = _dafny.Set({d_358_v104_})
                nw51_[int(4)] = (d_359_v105_ if d_328_cf81_ else d_359_v105_)
                nw51_[int(5)] = d_359_v105_
                def iife79_():
                    coll55_ = _dafny.Set()
                    compr_55_: _dafny.Seq
                    for compr_55_ in (d_361_v107_).keys.Elements:
                        d_370_v108_: _dafny.Seq = compr_55_
                        if (d_370_v108_) in (d_361_v107_):
                            coll55_ = coll55_.union(_dafny.Set([d_370_v108_]))
                    return _dafny.Set(coll55_)
                nw51_[int(6)] = (iife79_()
                ).intersection(d_359_v105_)
                nw51_[int(7)] = (d_363_v109_)[default__.safeIndex(d_219_v2_, len(d_363_v109_))]
                nw51_[int(8)] = (d_363_v109_)[default__.safeIndex(d_219_v2_, len(d_363_v109_))]
                nw51_[int(9)] = d_359_v105_
                nw51_[int(10)] = ((d_363_v109_)[default__.safeIndex((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], len(d_363_v109_))]).intersection(d_359_v105_)
                def iife80_():
                    coll56_ = _dafny.Set()
                    compr_56_: _dafny.Seq
                    for compr_56_ in (d_364_v110_).keys.Elements:
                        d_371_v111_: _dafny.Seq = compr_56_
                        if (d_371_v111_) in (d_364_v110_):
                            coll56_ = coll56_.union(_dafny.Set([d_371_v111_]))
                    return _dafny.Set(coll56_)
                nw51_[int(11)] = iife80_()
                
                nw51_[int(12)] = d_359_v105_
                nw51_[int(13)] = (d_359_v105_) - (d_359_v105_)
                nw51_[int(14)] = d_359_v105_
                def iife81_():
                    coll57_ = _dafny.Set()
                    compr_57_: _dafny.Seq
                    for compr_57_ in (_dafny.SeqWithoutIsStrInference([(d_283_v45_).fm7(d_232_globalState_)])).Elements:
                        d_372_v113_: _dafny.Seq = compr_57_
                        if (d_372_v113_) in (_dafny.SeqWithoutIsStrInference([(d_283_v45_).fm7(d_232_globalState_)])):
                            coll57_ = coll57_.union(_dafny.Set([d_372_v113_]))
                    return _dafny.Set(coll57_)
                nw51_[int(15)] = (iife81_()
                 if (d_365_v112_).cf45 else d_359_v105_)
                nw51_[int(16)] = _dafny.Set({d_358_v104_, default__.fm45(d_217_v0_, d_367_v115_, ((d_222_v5_)[591] if (591) in (d_222_v5_) else d_219_v2_), d_232_globalState_)})
                nw51_[int(17)] = _dafny.Set({d_358_v104_})
                nw51_[int(18)] = (d_363_v109_)[default__.safeIndex(d_219_v2_, len(d_363_v109_))]
                nw51_[int(19)] = d_359_v105_
                nw51_[int(20)] = d_359_v105_
                nw51_[int(21)] = d_359_v105_
                d_368_v116_ = nw51_
                index43_ = default__.safeIndex(785, (d_368_v116_).length(0))
                (d_368_v116_)[index43_] = _dafny.Set({d_358_v104_})
                d_373_v117_: T1
                nw52_ = C10()
                nw52_.ctor__(d_217_v0_, (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], d_230_v13_, d_217_v0_, d_280_v44_, d_328_cf81_)
                d_373_v117_ = nw52_
                d_374_v118_: _dafny.Map
                d_374_v118_ = _dafny.Map({d_328_cf81_: d_373_v117_})
                d_375_v119_: _dafny.Array
                nw53_ = _dafny.Array(None, 18)
                d_375_v119_ = nw53_
                d_376_v120_: D18
                d_376_v120_ = D18_DC44(D18_DC41(d_375_v119_))
                d_377_v121_: D11
                d_377_v121_ = D11_DC23((d_373_v117_).f24, d_219_v2_, d_225_v8_, d_284_v46_)
                d_378_v122_: D0
                d_378_v122_ = D0_DC2(d_328_cf81_, d_219_v2_, d_328_cf81_)
                d_379_v123_: D27
                d_379_v123_ = D27_DC61(d_374_v118_)
                d_380_v124_: _dafny.Seq
                d_380_v124_ = _dafny.SeqWithoutIsStrInference([((d_379_v123_).cf99 if d_328_cf81_ else d_374_v118_), d_374_v118_])
                index44_ = default__.safeIndex(785, (d_368_v116_).length(0))
                rhs47_ = True
                rhs48_ = (d_358_v104_) + ((d_358_v104_) + ((_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_376_v120_])).cardinality])).set(default__.safeIndex(d_219_v2_, len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_376_v120_])).cardinality]))), d_219_v2_)))
                rhs49_ = d_224_v7_
                rhs50_ = default__.fm61((d_377_v121_).cf37, default__.fm1(d_217_v0_, d_232_globalState_), (d_378_v122_ if not(False) else d_378_v122_), d_219_v2_, d_232_globalState_)
                rhs51_ = (d_380_v124_)[default__.safeIndex((d_219_v2_) + ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]), len(d_380_v124_))]
                lhs33_ = d_232_globalState_
                lhs34_ = d_368_v116_
                lhs35_ = default__.safeIndex(785, (d_368_v116_).length(0))
                lhs33_.f13 = rhs47_
                d_358_v104_ = rhs48_
                d_224_v7_ = rhs49_
                lhs34_[lhs35_] = rhs50_
                d_374_v118_ = rhs51_
                (d_232_globalState_).f13 = (d_373_v117_).f24
                index45_ = default__.safeIndex(95, (d_228_v11_).length(0))
                (d_228_v11_)[index45_] = d_219_v2_
                (d_232_globalState_).f13 = d_217_v0_
                d_381_v125_: T2
                nw54_ = C4()
                nw54_.ctor__(d_219_v2_, d_222_v5_, d_280_v44_, d_373_v117_.f23, d_328_cf81_)
                d_381_v125_ = nw54_
                d_382_v126_: _dafny.Map
                d_382_v126_ = _dafny.Map({d_381_v125_: d_328_cf81_})
                d_383_v127_: bool
                d_384_v128_: int
                d_385_v129_: bool
                out18_: bool
                out19_: int
                out20_: bool
                out18_, out19_, out20_ = (d_283_v45_).m3(((d_382_v126_)[d_381_v125_] if (d_381_v125_) in (d_382_v126_) else (d_381_v125_).f24), d_230_v13_, d_232_globalState_)
                d_383_v127_ = out18_
                d_384_v128_ = out19_
                d_385_v129_ = out20_
            d_217_v0_ = d_217_v0_
        elif source6_.is_DC50:
            d_386___mcc_h7_ = source6_.cf80
            d_387_cf80_ = d_386___mcc_h7_
            d_388_v130_: C10
            nw55_ = C10()
            nw55_.ctor__((41) >= ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]), d_219_v2_, d_230_v13_, False, d_280_v44_, d_217_v0_)
            d_388_v130_ = nw55_
            d_389_v131_: _dafny.Seq
            d_389_v131_ = _dafny.SeqWithoutIsStrInference([685, 402, d_219_v2_])
            d_389_v131_ = d_389_v131_
            index46_ = default__.safeIndex(95, (d_228_v11_).length(0))
            (d_228_v11_)[index46_] = d_388_v130_.f29
            d_227_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ueed"))
        elif True:
            d_390___mcc_h8_ = source6_.cf82
            d_391_cf82_ = d_390___mcc_h8_
            d_392_v132_: _dafny.Array
            nw56_ = _dafny.Array(_dafny.Array(None, 0), 29)
            d_392_v132_ = nw56_
            index47_ = default__.safeIndex(144, (d_392_v132_).length(0))
            (d_392_v132_)[index47_] = d_228_v11_
            d_393_v133_: _dafny.Map
            d_393_v133_ = _dafny.Map({(d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]: d_228_v11_})
            index48_ = default__.safeIndex(144, (d_392_v132_).length(0))
            index49_ = default__.safeIndex(95, (d_228_v11_).length(0))
            rhs52_ = d_228_v11_
            rhs53_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xy")))
            rhs54_ = d_219_v2_
            rhs55_ = d_393_v133_
            lhs36_ = d_392_v132_
            lhs37_ = default__.safeIndex(144, (d_392_v132_).length(0))
            lhs38_ = d_228_v11_
            lhs39_ = default__.safeIndex(95, (d_228_v11_).length(0))
            lhs36_[lhs37_] = rhs52_
            d_219_v2_ = rhs53_
            lhs38_[lhs39_] = rhs54_
            d_393_v133_ = rhs55_
            d_394_v134_: C7
            nw57_ = C7()
            nw57_.ctor__(not((d_217_v0_ if d_217_v0_ else d_217_v0_)), False, d_280_v44_, d_230_v13_, (d_224_v7_) < (d_224_v7_), (d_283_v45_).fm4(d_219_v2_, d_224_v7_, d_232_globalState_))
            d_394_v134_ = nw57_
            d_227_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
            d_227_v10_ = ((d_227_v10_).set(default__.safeIndex((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], len(d_227_v10_)), d_230_v13_)) + (d_227_v10_)
        rhs56_ = (0) - (d_219_v2_)
        rhs57_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbtyeu"))) + (d_227_v10_)) != ((default__.fm1(d_217_v0_, d_232_globalState_)).set(default__.safeIndex(d_219_v2_, len(default__.fm1(d_217_v0_, d_232_globalState_))), d_230_v13_))
        lhs40_ = d_232_globalState_
        lhs40_.f7 = rhs56_
        d_217_v0_ = rhs57_
        index50_ = default__.safeIndex(261, (d_218_v1_).length(0))
        (d_218_v1_)[index50_] = (d_217_v0_ if d_217_v0_ else d_217_v0_)
        d_395_v135_: D18
        d_395_v135_ = D18_DC42((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], d_217_v0_)
        d_396_v136_: _dafny.Set
        d_396_v136_ = _dafny.Set({d_395_v135_})
        d_397_v137_: _dafny.Map
        d_397_v137_ = _dafny.Map({d_395_v135_: len(d_260_v36_)})
        index51_ = default__.safeIndex(261, (d_218_v1_).length(0))
        def iife82_():
            coll58_ = _dafny.Set()
            compr_58_: D18
            for compr_58_ in (d_397_v137_).keys.Elements:
                d_398_v138_: D18 = compr_58_
                if (d_398_v138_) in (d_397_v137_):
                    coll58_ = coll58_.union(_dafny.Set([d_398_v138_]))
            return _dafny.Set(coll58_)
        (d_218_v1_)[index51_] = (d_396_v136_) != (iife82_()
        )
        _ingredientsd_0 = []
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_218_v1_).length(0)):
            d_399_i14_: int = guard_loop_0_
            if (True) and (((0) <= (d_399_i14_)) and ((d_399_i14_) < ((d_218_v1_).length(0)))):
                _ingredientsd_0.append((d_218_v1_, int(d_399_i14_), (d_218_v1_)[default__.safeIndex(261, (d_218_v1_).length(0))]))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        d_400_v139_: D9
        d_400_v139_ = D9_DC20(d_227_v10_, d_217_v0_, (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))])
        d_401_v140_: D4
        d_401_v140_ = D4_DC10(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnm"))).set(default__.safeIndex((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnm")))), d_230_v13_)) + ((d_400_v139_).cf31))
        source7_ = d_401_v140_
        if source7_.is_DC11:
            d_402___mcc_h9_ = source7_.cf18
            d_403_cf18_ = d_402___mcc_h9_
            d_404_v141_: _dafny.Seq
            d_404_v141_ = _dafny.SeqWithoutIsStrInference([(d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], d_403_cf18_, (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]])
            d_405_v142_: _dafny.Map
            d_405_v142_ = d_222_v5_
            d_406_v143_: C4
            nw58_ = C4()
            nw58_.ctor__(len(d_404_v141_), d_405_v142_, d_280_v44_, d_230_v13_, (d_283_v45_).fm4(((d_222_v5_)[d_219_v2_] if (d_219_v2_) in (d_222_v5_) else (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]), d_224_v7_, d_232_globalState_))
            d_406_v143_ = nw58_
            d_407_v144_: D24
            d_407_v144_ = D24_DC56(d_406_v143_)
            index52_ = default__.safeIndex(261, (d_218_v1_).length(0))
            index53_ = default__.safeIndex(261, (d_218_v1_).length(0))
            rhs58_ = (d_218_v1_)[default__.safeIndex(261, (d_218_v1_).length(0))]
            rhs59_ = d_407_v144_
            rhs60_ = d_217_v0_
            rhs61_ = (d_218_v1_)[default__.safeIndex(261, (d_218_v1_).length(0))]
            lhs41_ = d_218_v1_
            lhs42_ = default__.safeIndex(261, (d_218_v1_).length(0))
            lhs43_ = d_218_v1_
            lhs44_ = default__.safeIndex(261, (d_218_v1_).length(0))
            lhs45_ = d_232_globalState_
            lhs41_[lhs42_] = rhs58_
            d_407_v144_ = rhs59_
            lhs43_[lhs44_] = rhs60_
            lhs45_.f13 = rhs61_
            (d_232_globalState_).f7 = (0) - ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))])
            d_408_v145_: bool
            d_409_v146_: str
            d_410_v147_: bool
            out21_: bool
            out22_: str
            out23_: bool
            out21_, out22_, out23_ = (d_283_v45_).m11((d_217_v0_ if (d_218_v1_)[default__.safeIndex(261, (d_218_v1_).length(0))] else d_217_v0_), d_232_globalState_)
            d_408_v145_ = out21_
            d_409_v146_ = out22_
            d_410_v147_ = out23_
            index54_ = default__.safeIndex(261, (d_218_v1_).length(0))
            (d_218_v1_)[index54_] = not(((d_224_v7_) + (_dafny.SeqWithoutIsStrInference([(d_224_v7_)[default__.safeIndex(d_403_cf18_, len(d_224_v7_))]]))) == ((d_224_v7_) + (d_224_v7_)))
        elif True:
            d_411___mcc_h10_ = source7_.cf17
            d_412_cf17_ = d_411___mcc_h10_
            d_413_v148_: _dafny.Array
            def lambda13_(d_414_v0_, d_415_v13_):
                def lambda14_(d_416_i15_):
                    return _dafny.Map({d_414_v0_: d_415_v13_})

                return lambda14_

            init7_ = lambda13_(d_217_v0_, d_230_v13_)
            nw59_ = _dafny.Array(None, 5)
            for i0_7_ in range(nw59_.length(0)):
                nw59_[i0_7_] = init7_(i0_7_)
            d_413_v148_ = nw59_
            d_417_v149_: D14
            d_417_v149_ = D14_DC30(d_413_v148_)
            d_418_v150_: D14
            d_418_v150_ = D14_DC32(d_417_v149_)
            d_419_v151_: D14
            d_419_v151_ = D14_DC32((d_418_v150_ if d_217_v0_ else D14_DC31(d_217_v0_)))
            d_420_v152_: _dafny.Map
            d_420_v152_ = _dafny.Map({(d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]: d_219_v2_})
            d_421_v153_: C4
            nw60_ = C4()
            nw60_.ctor__(d_219_v2_, d_420_v152_, d_280_v44_, d_230_v13_, True)
            d_421_v153_ = nw60_
            d_422_v155_: _dafny.Map
            d_422_v155_ = _dafny.Map({d_219_v2_: d_395_v135_})
            d_423_v156_: _dafny.Seq
            d_423_v156_ = _dafny.SeqWithoutIsStrInference([default__.fm0((d_218_v1_)[default__.safeIndex(261, (d_218_v1_).length(0))], (d_421_v153_).f33, d_232_globalState_)])
            def iife83_():
                coll59_ = _dafny.Map()
                compr_59_: int
                for compr_59_ in ((d_422_v155_).set((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], d_395_v135_)).keys.Elements:
                    d_424_v154_: int = compr_59_
                    if (d_424_v154_) in ((d_422_v155_).set((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], d_395_v135_)):
                        coll59_[default__.safeModuloInt(d_424_v154_, (d_421_v153_).f33)] = d_230_v13_
                return _dafny.Map(coll59_)
            rhs62_ = (d_283_v45_).fm4((d_219_v2_) + ((_dafny.MultiSet([d_421_v153_, d_421_v153_])).cardinality), d_224_v7_, d_232_globalState_)
            rhs63_ = d_419_v151_
            rhs64_ = (d_218_v1_)[default__.safeIndex(261, (d_218_v1_).length(0))]
            rhs65_ = len(iife83_()
            )
            rhs66_ = ((d_423_v156_ if not((d_218_v1_)[default__.safeIndex(261, (d_218_v1_).length(0))]) else d_423_v156_)) < ((d_423_v156_).set(default__.safeIndex((d_421_v153_).f33, len(d_423_v156_)), (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]))
            lhs46_ = d_232_globalState_
            lhs47_ = d_232_globalState_
            lhs48_ = d_232_globalState_
            lhs49_ = d_232_globalState_
            lhs46_.f13 = rhs62_
            d_419_v151_ = rhs63_
            lhs47_.f13 = rhs64_
            lhs48_.f7 = rhs65_
            lhs49_.f13 = rhs66_
            if (_dafny.CodePoint('a')) not in ((d_412_cf17_) + (d_227_v10_)):
                index55_ = default__.safeIndex(261, (d_218_v1_).length(0))
                (d_218_v1_)[index55_] = (d_218_v1_)[default__.safeIndex(261, (d_218_v1_).length(0))]
                (d_232_globalState_).f7 = default__.safeModuloInt((d_219_v2_) * (len((d_224_v7_).set(default__.safeIndex(len(d_412_cf17_), len(d_224_v7_)), d_217_v0_))), (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))])
                d_425_v157_: C8
                nw61_ = C8()
                nw61_.ctor__(d_217_v0_)
                d_425_v157_ = nw61_
                d_426_v158_: C8
                d_426_v158_ = d_425_v157_
                d_427_v159_: _dafny.Map
                d_427_v159_ = _dafny.Map({(d_421_v153_).f33: (d_426_v158_)})
                d_427_v159_ = d_427_v159_
                d_219_v2_ = d_219_v2_
                d_428_v160_: C3
                d_428_v160_ = d_283_v45_
                d_283_v45_ = (d_428_v160_)
            elif True:
                index56_ = default__.safeIndex(95, (d_228_v11_).length(0))
                (d_228_v11_)[index56_] = ((0) - ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))])) * (d_219_v2_)
                (d_232_globalState_).f7 = (d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]
                d_429_v161_: D14
                d_429_v161_ = D14_DC30(d_413_v148_)
                d_429_v161_ = d_429_v161_
                d_430_v162_: bool
                d_431_v163_: int
                d_432_v164_: bool
                out24_: bool
                out25_: int
                out26_: bool
                out24_, out25_, out26_ = (d_421_v153_).m3((d_219_v2_) > ((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]), d_230_v13_, d_232_globalState_)
                d_430_v162_ = out24_
                d_431_v163_ = out25_
                d_432_v164_ = out26_
                (d_232_globalState_).f13 = (d_224_v7_)[default__.safeIndex(d_219_v2_, len(d_224_v7_))]
            d_433_v165_: _dafny.MultiSet
            d_433_v165_ = _dafny.MultiSet([d_217_v0_])
            d_434_v166_: C2
            nw62_ = C2()
            nw62_.ctor__(d_433_v165_, (d_280_v44_ if (d_218_v1_)[default__.safeIndex(261, (d_218_v1_).length(0))] else d_280_v44_), d_230_v13_, d_217_v0_)
            d_434_v166_ = nw62_
            (d_232_globalState_).f7 = len(_dafny.Map({(_dafny.SeqWithoutIsStrInference([(d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], d_219_v2_])) + (_dafny.SeqWithoutIsStrInference([(d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))]])): not(d_217_v0_)}))
        d_435_v167_: _dafny.Map
        d_435_v167_ = d_222_v5_
        source8_ = d_435_v167_
        d_436___mcc_h11_ = source8_
        d_437_cf16_ = d_436___mcc_h11_
        d_438_v168_: int
        d_439_v169_: bool
        d_440_v170_: int
        out27_: int
        out28_: bool
        out29_: int
        out27_, out28_, out29_ = default__.m0((d_228_v11_)[default__.safeIndex(95, (d_228_v11_).length(0))], _dafny.MultiSet([(d_218_v1_)[default__.safeIndex(261, (d_218_v1_).length(0))]]), ((d_218_v1_)[default__.safeIndex(261, (d_218_v1_).length(0))]) == (True), d_219_v2_, d_232_globalState_)
        d_438_v168_ = out27_
        d_439_v169_ = out28_
        d_440_v170_ = out29_
        d_441_v171_: _dafny.MultiSet
        d_441_v171_ = _dafny.MultiSet([d_228_v11_, d_228_v11_])
        d_441_v171_ = d_441_v171_
        (d_232_globalState_).f7 = len(_dafny.Map({default__.fm2(d_232_globalState_): d_217_v0_}))
        d_442_v172_: _dafny.MultiSet
        d_442_v172_ = _dafny.MultiSet([(d_218_v1_)[default__.safeIndex(261, (d_218_v1_).length(0))], d_439_v169_, d_217_v0_])
        d_443_v173_: C2
        nw63_ = C2()
        nw63_.ctor__(d_442_v172_, d_280_v44_, (d_230_v13_ if (d_283_v45_).fm4(d_440_v170_, d_224_v7_, d_232_globalState_) else _dafny.CodePoint('a')), (False) == (not(d_217_v0_)))
        d_443_v173_ = nw63_
        _dafny.print(_dafny.string_of(d_217_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v1_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_219_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_220_v3_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_v4_) == (_dafny.Map({298: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_222_v5_) == (_dafny.Map({230: 230}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_223_v6_) == (_dafny.Map({True: _dafny.CodePoint('t')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_224_v7_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_225_v8_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_v9_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_227_v10_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v11_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_229_v12_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_230_v13_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_231_v14_) == (_dafny.Map({-230: _dafny.Map({True: True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f0)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_232_globalState_).f2)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f4)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f6)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_232_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_232_globalState_).f8).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_232_globalState_.f10)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_232_globalState_).f12).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_232_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f15) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_232_globalState_).f17).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_232_globalState_).f18).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_globalState_).f21)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_globalState_.f22)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v36_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_261_v37_) == (_dafny.Set({0, 230}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v38_) == (_dafny.Set({462, 2, 230}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v39_) == (_dafny.MultiSet([0]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_264_v40_)[5]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpbndwsy")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwwgiekbdyjkt")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjoqku"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_265_v41_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpbndwsy")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwwgiekbdyjkt")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjoqku")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjoqku"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_266_v42_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_279_v43_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v44_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v44_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v44_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v44_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v44_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_283_v45_).f35)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v46_).cf23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v47_) == (_dafny.Set({D5_DC14(True)}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_288_v49_) == (_dafny.Map({_dafny.Set({D5_DC14(True)}): _dafny.Map({})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v52_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.Set({D5_DC14(True)}): _dafny.Map({})}), _dafny.Map({_dafny.Set({D5_DC14(True)}): _dafny.Map({230: 230}), _dafny.Set({D5_DC14(True), D5_DC14(False)}): _dafny.Map({230: 230})}), _dafny.Map({_dafny.Set({D5_DC14(True)}): _dafny.Map({})}), _dafny.Map({_dafny.Set({D5_DC14(True), D5_DC14(False)}): _dafny.Map({655: -349}), _dafny.Set({D5_DC14(False)}): _dafny.Map({-697: -69}), _dafny.Set({D5_DC14(True)}): _dafny.Map({-893: 6})}), _dafny.Map({_dafny.Set({D5_DC14(True)}): _dafny.Map({})})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_294_v53_)[1])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_295_v54_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_296_v55_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_326_v76_).cf81))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_395_v135_).cf68))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_395_v135_).cf69))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_396_v136_) == (_dafny.Set({D18_DC42(0, True)}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_397_v137_) == (_dafny.Map({D18_DC42(0, True): 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_400_v139_).cf31).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_400_v139_).cf32))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_400_v139_).cf33))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_401_v140_).cf17).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_435_v167_)) == (_dafny.Map({230: 230}))))
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
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC5(int(0), int(0), _dafny.Seq({}), _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC5(D1, NamedTuple('DC5', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

class D2_DC7(D2, NamedTuple('DC7', [('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC9(D3, NamedTuple('DC9', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC11(D4, NamedTuple('DC11', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({self.cf17.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC13(D5, NamedTuple('DC13', [('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(_dafny.Map({}), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)

class D6_DC16(D6, NamedTuple('DC16', [('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)

class D7_DC17(D7, NamedTuple('DC17', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)

class D8_DC18(D8, NamedTuple('DC18', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC20(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D9_DC19)

class D9_DC20(D9, NamedTuple('DC20', [('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({self.cf31.VerbatimString(True)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC19(D9, NamedTuple('DC19', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC19({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC19) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D10_DC21)

class D10_DC21(D10, NamedTuple('DC21', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC21({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC21) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC23(False, int(0), _dafny.Map({}), D5.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D11_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D11_DC22)

class D11_DC23(D11, NamedTuple('DC23', [('cf36', Any), ('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC23({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC23) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC22(D11, NamedTuple('DC22', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC22({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC22) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC25(False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D12_DC25)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D12_DC24)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D12_DC26)

class D12_DC25(D12, NamedTuple('DC25', [('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC25({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC25) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC24(D12, NamedTuple('DC24', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC24({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC24) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC26(D12, NamedTuple('DC26', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC26({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC26) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC28(False, int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D13_DC28)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D13_DC29)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D13_DC27)

class D13_DC28(D13, NamedTuple('DC28', [('cf45', Any), ('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC28({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC28) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC29(D13, NamedTuple('DC29', [('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC29({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC29) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC27(D13, NamedTuple('DC27', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC27({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC27) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC31(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D14_DC31)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D14_DC30)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D14_DC32)

class D14_DC31(D14, NamedTuple('DC31', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC31({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC31) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC30(D14, NamedTuple('DC30', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC30({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC30) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC32(D14, NamedTuple('DC32', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC32({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC32) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC34(False)
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

class D15_DC34(D15, NamedTuple('DC34', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC34({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC34) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC35(D15, NamedTuple('DC35', [('cf56', Any), ('cf57', Any), ('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC35({_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC35) and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC36(D15, NamedTuple('DC36', [('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC36({_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC36) and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC33(D15, NamedTuple('DC33', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC33({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC33) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC38(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Array(None, 0), _dafny.Set({}))
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

class D16_DC38(D16, NamedTuple('DC38', [('cf62', Any), ('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC38({self.cf62.VerbatimString(True)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC38) and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC37(D16, NamedTuple('DC37', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC37({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC37) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC39(D16, NamedTuple('DC39', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC39({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC39) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D17_DC40)

class D17_DC40(D17, NamedTuple('DC40', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC40({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC40) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC42(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D18_DC42)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D18_DC43)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D18_DC41)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D18_DC44)

class D18_DC42(D18, NamedTuple('DC42', [('cf68', Any), ('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC42({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC42) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC43(D18, NamedTuple('DC43', [('cf70', Any), ('cf71', Any), ('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC43({_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)}, {self.cf72.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC43) and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC41(D18, NamedTuple('DC41', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC41({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC41) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC44(D18, NamedTuple('DC44', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC44({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC44) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC46(_dafny.Set({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D19_DC46)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D19_DC47)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D19_DC45)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D19_DC48)

class D19_DC46(D19, NamedTuple('DC46', [('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC46({_dafny.string_of(self.cf75)}, {self.cf76.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC46) and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC47(D19, NamedTuple('DC47', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC47({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC47) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC45(D19, NamedTuple('DC45', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC45({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC45) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC48(D19, NamedTuple('DC48', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC48({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC48) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D20_DC49)

class D20_DC49(D20, NamedTuple('DC49', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC49({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC49) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC51(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D21_DC51)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D21_DC50)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D21_DC52)

class D21_DC51(D21, NamedTuple('DC51', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC51({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC51) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC50(D21, NamedTuple('DC50', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC50({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC50) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC52(D21, NamedTuple('DC52', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC52({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC52) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D22_DC53)

class D22_DC53(D22, NamedTuple('DC53', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC53({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC53) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC55(D21.default()(), _dafny.Array(None, 0), _dafny.MultiSet({}), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D23_DC55)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D23_DC54)

class D23_DC55(D23, NamedTuple('DC55', [('cf85', Any), ('cf86', Any), ('cf87', Any), ('cf88', Any), ('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC55({_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC55) and self.cf85 == __o.cf85 and self.cf86 == __o.cf86 and self.cf87 == __o.cf87 and self.cf88 == __o.cf88 and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC54(D23, NamedTuple('DC54', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC54({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC54) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC57(None, False, _dafny.CodePoint('D'), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D24_DC57)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D24_DC56)

class D24_DC57(D24, NamedTuple('DC57', [('cf91', Any), ('cf92', Any), ('cf93', Any), ('cf94', Any), ('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC57({_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC57) and self.cf91 == __o.cf91 and self.cf92 == __o.cf92 and self.cf93 == __o.cf93 and self.cf94 == __o.cf94 and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC56(D24, NamedTuple('DC56', [('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC56({_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC56) and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC59(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D25_DC59)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D25_DC58)

class D25_DC59(D25, NamedTuple('DC59', [('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC59({_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC59) and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC58(D25, NamedTuple('DC58', [('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC58({_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC58) and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D26_DC60)

class D26_DC60(D26, NamedTuple('DC60', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC60({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC60) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC62(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.CodePoint('D'), False, _dafny.Seq({}))
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
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D27_DC64)

class D27_DC62(D27, NamedTuple('DC62', [('cf100', Any), ('cf101', Any), ('cf102', Any), ('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC62({self.cf100.VerbatimString(True)}, {_dafny.string_of(self.cf101)}, {_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC62) and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102 and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC63(D27, NamedTuple('DC63', [('cf104', Any), ('cf105', Any), ('cf106', Any), ('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC63({_dafny.string_of(self.cf104)}, {_dafny.string_of(self.cf105)}, {_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC63) and self.cf104 == __o.cf104 and self.cf105 == __o.cf105 and self.cf106 == __o.cf106 and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC61(D27, NamedTuple('DC61', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC61({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC61) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC64(D27, NamedTuple('DC64', [('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC64({_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC64) and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D28_DC65)

class D28_DC65(D28, NamedTuple('DC65', [('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC65({_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC65) and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D29_DC66)

class D29_DC66(D29, NamedTuple('DC66', [('cf110', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC66({_dafny.string_of(self.cf110)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC66) and self.cf110 == __o.cf110
    def __hash__(self) -> int:
        return super().__hash__()


class D30:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D30_DC67)

class D30_DC67(D30, NamedTuple('DC67', [('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC67({_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC67) and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()


class D31:
    @classmethod
    def default(cls, ):
        return lambda: D31_DC69(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D31_DC69)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D31_DC70)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D31_DC68)

class D31_DC69(D31, NamedTuple('DC69', [('cf113', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC69({_dafny.string_of(self.cf113)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC69) and self.cf113 == __o.cf113
    def __hash__(self) -> int:
        return super().__hash__()

class D31_DC70(D31, NamedTuple('DC70', [('cf114', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC70({_dafny.string_of(self.cf114)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC70) and self.cf114 == __o.cf114
    def __hash__(self) -> int:
        return super().__hash__()

class D31_DC68(D31, NamedTuple('DC68', [('cf112', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC68({_dafny.string_of(self.cf112)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC68) and self.cf112 == __o.cf112
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f23(self):
        return self._f23
    @f23.setter
    def f23(self, value):
        self._f23 = value

class T1:
    pass
    def m1(self, p0, p1, globalState):
        pass

    def m2(self, globalState):
        pass


class T2:
    pass
    def m3(self, p0, p1, globalState):
        pass


class T3:
    pass
    def m4(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f7: int = int(0)
        self.f10: _dafny.Set = _dafny.Set({})
        self.f13: bool = False
        self.f22: _dafny.Array = _dafny.Array(None, 0)
        self._f0: _dafny.Array = _dafny.Array(None, 0)
        self._f1: int = int(0)
        self._f2: _dafny.Map = _dafny.Map({})
        self._f3: bool = False
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f5: bool = False
        self._f6: _dafny.Array = _dafny.Array(None, 0)
        self._f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f9: bool = False
        self._f11: int = int(0)
        self._f12: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f14: str = _dafny.CodePoint('D')
        self._f15: _dafny.Map = _dafny.Map({})
        self._f16: int = int(0)
        self._f17: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f18: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f19: bool = False
        self._f20: int = int(0)
        self._f21: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21
        (self).f22 = f22

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
    def f6(self):
        return self._f6
    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f11(self):
        return self._f11
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
    def f16(self):
        return self._f16
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

class C0(T0):
    def  __init__(self):
        self._f23: str = _dafny.CodePoint('D')
        self._f24: bool = False
        self.f37: _dafny.Map = _dafny.Map({})
        self._f38: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f23(self):
        return self._f23
    @f23.setter
    def f23(self, value):
        self._f23 = value
    @property
    def f24(self):
        return self._f24
    def ctor__(self, f37, f38, f23, f24):
        (self).f37 = f37
        (self)._f38 = f38
        (self).f23 = f23
        (self)._f24 = f24

    def fm4(self, p0, p1, globalState):
        return not(not(((0) - (len((_dafny.Map({(self).f24: (self).f24})) | (_dafny.Map({True: (self).f24}))))) >= ((0) - ((0) - ((len(_dafny.Map({(self).f38: (D0_DC2((self).f24, len(_dafny.Map({624: True})), (self).f24)).cf4}))) - ((_dafny.MultiSet([(self).f24])).cardinality))))))

    def fm26(self, p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([(self).f24, (self).f24])), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({(self).f38}))]))), (403) * ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdnuosq")))))])

    @property
    def f38(self):
        return self._f38

class C1(T3):
    def  __init__(self):
        self._f27: bool = False
        self._f39: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f39, f27):
        (self)._f39 = f39
        (self)._f27 = f27

    def fm9(self, p0, p1, p2, globalState):
        return default__.safeDivisionInt((0) - (((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snfxoe")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_444_i0_ in range(default__.abs(957))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_445_i1_ in range(default__.abs(-692))])])).cardinality) * ((_dafny.MultiSet([(self).f39, (self).f27, (self).f27])).cardinality)), 934)

    def fm32(self, p0, p1, globalState):
        return not(((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "afadxteqa")): 915})): (self).f39}))])) + (_dafny.SeqWithoutIsStrInference([155, len(_dafny.SeqWithoutIsStrInference([(self).f39]))]))) <= ((_dafny.SeqWithoutIsStrInference([590, len(_dafny.SeqWithoutIsStrInference([(self).f39]))])) + (_dafny.SeqWithoutIsStrInference([671, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yahpsabt")))]))))

    def fm33(self, globalState):
        def iife84_():
            coll60_ = _dafny.Set()
            compr_60_: D5
            for compr_60_ in (_dafny.SeqWithoutIsStrInference([D5_DC14((self).f39)])).Elements:
                d_446_v0_: D5 = compr_60_
                if (d_446_v0_) in (_dafny.SeqWithoutIsStrInference([D5_DC14((self).f39)])):
                    coll60_ = coll60_.union(_dafny.Set([d_446_v0_]))
            return _dafny.Set(coll60_)
        return default__.safeModuloInt(506, len((iife84_()
        ) | (_dafny.Set({D5_DC14((self).f39), D5_DC14((self).f27)}))))

    def m4(self, p0, globalState):
        r0: bool = False
        d_447_i0_: int
        d_447_i0_ = 0
        with _dafny.label("0"):
            while True:
                with _dafny.c_label("0"):
                    if (d_447_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_447_i0_ = (d_447_i0_) + (1)
                    d_448_v0_: str
                    d_448_v0_ = _dafny.CodePoint('h')
                    d_448_v0_ = d_448_v0_
                    if (self).f39:
                        d_449_v1_: int
                        d_449_v1_ = 852
                        d_450_v2_: _dafny.Seq
                        d_450_v2_ = _dafny.SeqWithoutIsStrInference([d_449_v1_, d_449_v1_, d_449_v1_])
                        d_451_v3_: _dafny.MultiSet
                        d_451_v3_ = _dafny.MultiSet([d_449_v1_, d_449_v1_, len(d_450_v2_)])
                        d_452_v4_: _dafny.Map
                        d_452_v4_ = _dafny.Map({(self).f39: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dr"))).set(default__.safeIndex(d_449_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dr")))), _dafny.CodePoint('h'))})
                        d_453_v5_: D2
                        d_453_v5_ = D2_DC6(d_452_v4_)
                        d_454_v6_: _dafny.Map
                        d_454_v6_ = _dafny.Map({d_453_v5_: p0})
                        d_455_v7_: _dafny.Map
                        d_455_v7_ = _dafny.Map({d_449_v1_: True})
                        d_456_v8_: _dafny.Seq
                        d_456_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nubq"))
                        d_457_v9_: _dafny.Map
                        d_457_v9_ = _dafny.Map({(self).f39: p0})
                        d_458_v10_: _dafny.Array
                        nw64_ = _dafny.Array(None, 18)
                        nw64_[int(0)] = d_449_v1_
                        nw64_[int(1)] = d_449_v1_
                        nw64_[int(2)] = default__.safeModuloInt(len(d_450_v2_), (self).fm33(globalState))
                        nw64_[int(3)] = (d_451_v3_).cardinality
                        nw64_[int(4)] = len(d_454_v6_)
                        nw64_[int(5)] = (0) - ((_dafny.MultiSet([((d_455_v7_)[483] if (483) in (d_455_v7_) else p0)])).cardinality)
                        nw64_[int(6)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jovube")))
                        nw64_[int(7)] = (self).fm9(d_450_v2_, d_449_v1_, (self).f27, globalState)
                        nw64_[int(8)] = (self).fm33(globalState)
                        nw64_[int(9)] = d_449_v1_
                        nw64_[int(10)] = d_449_v1_
                        nw64_[int(11)] = d_449_v1_
                        nw64_[int(12)] = d_449_v1_
                        nw64_[int(13)] = -64
                        nw64_[int(14)] = len((d_456_v8_) + (d_456_v8_))
                        nw64_[int(15)] = (d_449_v1_ if (self).f39 else d_449_v1_)
                        nw64_[int(16)] = (d_451_v3_).cardinality
                        nw64_[int(17)] = len(d_457_v9_)
                        d_458_v10_ = nw64_
                        (globalState).f22 = d_458_v10_
                        d_448_v0_ = default__.fm34((d_449_v1_) - (d_449_v1_), p0, globalState)
                        (globalState).f7 = d_449_v1_
                        d_459_v11_: _dafny.Set
                        d_459_v11_ = _dafny.Set({d_449_v1_})
                        d_460_v12_: _dafny.Seq
                        d_460_v12_ = _dafny.SeqWithoutIsStrInference([d_459_v11_])
                        d_461_v13_: _dafny.MultiSet
                        d_461_v13_ = _dafny.MultiSet([_dafny.Set({d_449_v1_}), (d_460_v12_)[default__.safeIndex(d_449_v1_, len(d_460_v12_))], _dafny.Set({d_449_v1_}), d_459_v11_, default__.fm35((self).f27, globalState)])
                        d_450_v2_ = (((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrwmpm"))), d_449_v1_])) + (d_450_v2_)) + ((_dafny.SeqWithoutIsStrInference([(self).fm9(d_450_v2_, d_449_v1_, (self).f27, globalState), d_449_v1_, len(d_455_v7_)])) + (_dafny.SeqWithoutIsStrInference([(d_461_v13_).cardinality, d_449_v1_])))).set(default__.safeIndex((d_449_v1_) - ((self).fm9(d_450_v2_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulr"))).set(default__.safeIndex(d_449_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulr")))), d_448_v0_)), p0, globalState)), len(((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrwmpm"))), d_449_v1_])) + (d_450_v2_)) + ((_dafny.SeqWithoutIsStrInference([(self).fm9(d_450_v2_, d_449_v1_, (self).f27, globalState), d_449_v1_, len(d_455_v7_)])) + (_dafny.SeqWithoutIsStrInference([(d_461_v13_).cardinality, d_449_v1_]))))), len(d_450_v2_))
                        d_449_v1_ = (0) - (d_449_v1_)
                    elif True:
                        d_462_v14_: int
                        d_462_v14_ = 699
                        d_463_v15_: _dafny.Seq
                        d_463_v15_ = _dafny.SeqWithoutIsStrInference([d_462_v14_])
                        d_464_v16_: _dafny.Seq
                        d_464_v16_ = _dafny.SeqWithoutIsStrInference([d_463_v15_, d_463_v15_])
                        d_465_v17_: _dafny.Array
                        def lambda15_(d_466_i1_):
                            return (self).f39

                        init8_ = lambda15_
                        nw65_ = _dafny.Array(None, 19)
                        for i0_8_ in range(nw65_.length(0)):
                            nw65_[i0_8_] = init8_(i0_8_)
                        d_465_v17_ = nw65_
                        d_467_v18_: D1
                        d_467_v18_ = D1_DC5(d_462_v14_, d_462_v14_, d_464_v16_, d_465_v17_, d_462_v14_)
                        rhs67_ = (d_467_v18_).cf8
                        rhs68_ = not ((self).f27) or (p0)
                        rhs69_ = (self).f39
                        lhs50_ = globalState
                        lhs50_.f7 = rhs67_
                        r0 = rhs68_
                        r0 = rhs69_
                        (globalState).f13 = p0
                        d_468_v19_: _dafny.Seq
                        d_468_v19_ = _dafny.SeqWithoutIsStrInference([p0])
                        d_462_v14_ = len(d_468_v19_)
                        (globalState).f13 = (self).f39
                        d_469_v20_: _dafny.Array
                        nw66_ = _dafny.Array(int(0), 21)
                        d_469_v20_ = nw66_
                        d_470_v21_: _dafny.Array
                        nw67_ = _dafny.Array(None, 9)
                        nw67_[int(0)] = d_469_v20_
                        nw67_[int(1)] = d_469_v20_
                        nw67_[int(2)] = d_469_v20_
                        nw67_[int(3)] = d_469_v20_
                        nw67_[int(4)] = d_469_v20_
                        nw67_[int(5)] = d_469_v20_
                        nw67_[int(6)] = d_469_v20_
                        nw67_[int(7)] = d_469_v20_
                        nw67_[int(8)] = d_469_v20_
                        d_470_v21_ = nw67_
                        index57_ = default__.safeIndex(856, (d_470_v21_).length(0))
                        (d_470_v21_)[index57_] = d_469_v20_
                        index58_ = default__.safeIndex(856, (d_470_v21_).length(0))
                        (d_470_v21_)[index58_] = d_469_v20_
                    d_471_v22_: int
                    d_471_v22_ = 683
                    (globalState).f7 = (0) - (d_471_v22_)
                    d_472_v23_: _dafny.Map
                    d_472_v23_ = _dafny.Map({d_471_v22_: default__.fm36(globalState)})
                    d_473_v24_: _dafny.Array
                    def lambda16_(d_474_p0_):
                        def lambda17_(d_475_i2_):
                            return _dafny.Map({d_474_p0_: d_474_p0_})

                        return lambda17_

                    init9_ = lambda16_(p0)
                    nw68_ = _dafny.Array(None, 27)
                    for i0_9_ in range(nw68_.length(0)):
                        nw68_[i0_9_] = init9_(i0_9_)
                    d_473_v24_ = nw68_
                    d_476_v25_: _dafny.Set
                    d_476_v25_ = _dafny.Set({d_471_v22_})
                    d_477_v26_: _dafny.Seq
                    d_477_v26_ = _dafny.SeqWithoutIsStrInference([False])
                    d_478_v27_: _dafny.Map
                    d_478_v27_ = _dafny.Map({d_473_v24_: (d_472_v23_) | (((D6_DC15(d_472_v23_)).cf24).set(len(d_476_v25_), d_477_v26_))})
                    d_472_v23_ = ((d_478_v27_)[d_473_v24_] if (d_473_v24_) in (d_478_v27_) else default__.fm37(d_471_v22_, p0, globalState))
                    pass
            pass
        d_479_v28_: _dafny.MultiSet
        d_479_v28_ = _dafny.MultiSet([(self).f27])
        d_480_v29_: int
        d_480_v29_ = -13
        d_481_v30_: _dafny.Set
        d_481_v30_ = _dafny.Set({-220, d_480_v29_})
        d_482_v31_: _dafny.Map
        d_482_v31_ = _dafny.Map({d_479_v28_: (d_481_v30_).ispropersubset(_dafny.Set({len(d_481_v30_), ((d_479_v28_)[(self).f27] if ((self).f27) in (d_479_v28_) else d_480_v29_)}))})
        d_482_v31_ = d_482_v31_
        d_483_v32_: _dafny.Seq
        d_483_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkn"))
        d_484_v33_: C0
        nw69_ = C0()
        nw69_.ctor__((_dafny.Map({(self).f39: d_483_v32_})).set(p0, d_483_v32_), p0, _dafny.CodePoint('m'), p0)
        d_484_v33_ = nw69_
        (globalState).f13 = p0
        d_485_v34_: D5
        d_485_v34_ = D5_DC14((self).f27)
        d_485_v34_ = D5_DC14((self).f39)
        d_486_v35_: str
        d_486_v35_ = _dafny.CodePoint('u')
        d_487_v36_: _dafny.Map
        d_487_v36_ = _dafny.Map({d_486_v35_: (d_485_v34_).cf23})
        d_487_v36_ = ((d_487_v36_) | (d_487_v36_)) | (d_487_v36_)
        r0 = (self).f27
        return r0

    @property
    def f39(self):
        return self._f39

class C2(T2, T0):
    def  __init__(self):
        self._f23: str = _dafny.CodePoint('D')
        self._f24: bool = False
        self._f25: _dafny.Array = _dafny.Array(None, 0)
        self._f36: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f23(self):
        return self._f23
    @f23.setter
    def f23(self, value):
        self._f23 = value
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f36, f25, f23, f24):
        (self)._f36 = f36
        (self)._f25 = f25
        (self).f23 = f23
        (self)._f24 = f24

    def fm7(self, globalState):
        return _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([(self).f24])), len(_dafny.SeqWithoutIsStrInference([(self).f24]))), len((_dafny.Set({not((self).f24), (self).f24})) | (_dafny.Set({(self).f24}))), (998) + (870)])

    def fm8(self, p0, p1, globalState):
        return (_dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptgbrq"))})) | (_dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxchetkux"))}))

    def fm4(self, p0, p1, globalState):
        return (self).f24

    def fm24(self, p0, p1, globalState):
        def iife85_():
            coll61_ = _dafny.Set()
            compr_61_: int
            for compr_61_ in _dafny.IntegerRange(584, -980):
                d_488_v0_: int = compr_61_
                if ((584) <= (d_488_v0_)) and ((d_488_v0_) < (-980)):
                    coll61_ = coll61_.union(_dafny.Set([default__.safeModuloInt(d_488_v0_, 982)]))
            return _dafny.Set(coll61_)
        return (0) - (len(((_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))))})) | (iife85_()
        )) | ((D5_DC12(_dafny.Set({len(_dafny.Map({(self).f24: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jayht")))}))}))).cf19)))

    def fm25(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True, False])) + (_dafny.SeqWithoutIsStrInference([(self).f24]))) + ((_dafny.SeqWithoutIsStrInference([(self).f24])) + (_dafny.SeqWithoutIsStrInference([False])))

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        d_489_v0_: int
        d_489_v0_ = 268
        d_490_v1_: _dafny.Set
        d_490_v1_ = _dafny.Set({d_489_v0_, d_489_v0_})
        d_491_v2_: D5
        d_491_v2_ = D5_DC12(d_490_v1_)
        source9_ = d_491_v2_
        if source9_.is_DC13:
            d_492___mcc_h0_ = source9_.cf20
            d_493___mcc_h1_ = source9_.cf21
            d_494___mcc_h2_ = source9_.cf22
            d_495_cf22_ = d_494___mcc_h2_
            d_496_cf21_ = d_493___mcc_h1_
            d_497_cf20_ = d_492___mcc_h0_
            index59_ = default__.safeIndex(332, ((self).f25).length(0))
            ((self).f25)[index59_] = p1
            index60_ = default__.safeIndex(332, ((self).f25).length(0))
            ((self).f25)[index60_] = _dafny.CodePoint('d')
            d_498_v3_: D0
            d_498_v3_ = D0_DC1(d_489_v0_)
            d_499_v4_: _dafny.Map
            d_499_v4_ = _dafny.Map({(d_498_v3_) not in (_dafny.SeqWithoutIsStrInference([d_498_v3_])): d_489_v0_})
            d_500_v5_: _dafny.Seq
            d_500_v5_ = _dafny.SeqWithoutIsStrInference([p0, (self).f24])
            d_499_v4_ = (d_499_v4_).set((d_496_cf21_ if (d_500_v5_)[default__.safeIndex(default__.fm0((self).f24, d_489_v0_, globalState), len(d_500_v5_))] else d_495_cf22_), (0) - (len(d_500_v5_)))
            d_501_v6_: _dafny.Map
            d_501_v6_ = _dafny.Map({(d_497_cf20_) * (d_489_v0_): p0})
            d_501_v6_ = _dafny.Map({d_489_v0_: (d_495_cf22_) == ((self).f24)})
            d_502_v7_: _dafny.Array
            nw70_ = _dafny.Array(False, 22)
            d_502_v7_ = nw70_
            index61_ = default__.safeIndex(736, (d_502_v7_).length(0))
            (d_502_v7_)[index61_] = False
            index62_ = default__.safeIndex(736, (d_502_v7_).length(0))
            (d_502_v7_)[index62_] = p0
        elif source9_.is_DC14:
            d_503___mcc_h3_ = source9_.cf23
            d_504_cf23_ = d_503___mcc_h3_
            (globalState).f7 = (0) - ((d_489_v0_ if d_504_cf23_ else d_489_v0_))
            d_505_v8_: _dafny.Seq
            d_505_v8_ = _dafny.SeqWithoutIsStrInference([(self).f24])
            d_506_v9_: _dafny.Array
            nw71_ = _dafny.Array(None, 12)
            nw71_[int(0)] = p0
            nw71_[int(1)] = p0
            nw71_[int(2)] = d_504_cf23_
            nw71_[int(3)] = p0
            nw71_[int(4)] = d_504_cf23_
            nw71_[int(5)] = p0
            nw71_[int(6)] = (d_489_v0_) != ((0) - (d_489_v0_))
            nw71_[int(7)] = (self).f24
            nw71_[int(8)] = default__.fm2(globalState)
            nw71_[int(9)] = ((self).fm4(d_489_v0_, (d_505_v8_).set(default__.safeIndex(-415, len(d_505_v8_)), p0), globalState) if True else p0)
            nw71_[int(10)] = p0
            nw71_[int(11)] = ((self).f24 if d_504_cf23_ else p0)
            d_506_v9_ = nw71_
            index63_ = default__.safeIndex(32, (d_506_v9_).length(0))
            (d_506_v9_)[index63_] = True
            index64_ = default__.safeIndex(32, (d_506_v9_).length(0))
            (d_506_v9_)[index64_] = (self).f24
            d_507_v10_: _dafny.Seq
            d_507_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yf"))
            d_508_v11_: _dafny.Seq
            d_508_v11_ = _dafny.SeqWithoutIsStrInference([d_489_v0_])
            d_509_v12_: _dafny.Seq
            d_509_v12_ = _dafny.SeqWithoutIsStrInference([d_508_v11_])
            d_510_v13_: D1
            d_510_v13_ = D1_DC5(len(d_507_v10_), d_489_v0_, d_509_v12_, d_506_v9_, d_489_v0_)
            pat_let_tv11_ = d_489_v0_
            pat_let_tv12_ = d_506_v9_
            pat_let_tv13_ = d_509_v12_
            pat_let_tv14_ = d_489_v0_
            pat_let_tv15_ = d_508_v11_
            pat_let_tv16_ = d_489_v0_
            pat_let_tv17_ = d_508_v11_
            def iife87_(_pat_let13_0):
                def iife88_(d_511_dt__update__tmp_h0_):
                    def iife89_(_pat_let14_0):
                        def iife90_(d_512_dt__update_hcf7_h0_):
                            def iife91_(_pat_let15_0):
                                def iife92_(d_513_dt__update_hcf10_h0_):
                                    def iife93_(_pat_let16_0):
                                        def iife94_(d_514_dt__update_hcf9_h0_):
                                            return D1_DC5(d_512_dt__update_hcf7_h0_, (d_511_dt__update__tmp_h0_).cf8, d_514_dt__update_hcf9_h0_, d_513_dt__update_hcf10_h0_, (d_511_dt__update__tmp_h0_).cf11)
                                        return iife94_(_pat_let16_0)
                                    return iife93_(pat_let_tv13_)
                                return iife92_(_pat_let15_0)
                            return iife91_(pat_let_tv12_)
                        return iife90_(_pat_let14_0)
                    return iife89_(pat_let_tv11_)
                return iife88_(_pat_let13_0)
            def iife86_(_pat_let12_0):
                def iife95_(d_515_dt__update__tmp_h1_):
                    def iife96_(_pat_let17_0):
                        def iife97_(d_516_dt__update_hcf7_h1_):
                            def iife98_(_pat_let18_0):
                                def iife99_(d_518_dt__update_hcf8_h0_):
                                    return D1_DC5(d_516_dt__update_hcf7_h1_, d_518_dt__update_hcf8_h0_, (d_515_dt__update__tmp_h1_).cf9, (d_515_dt__update__tmp_h1_).cf10, (d_515_dt__update__tmp_h1_).cf11)
                                return iife99_(_pat_let18_0)
                            return iife98_((pat_let_tv15_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([pat_let_tv16_ for d_517_i0_ in range(default__.abs(-263))])), len(pat_let_tv17_))])
                        return iife97_(_pat_let17_0)
                    return iife96_(pat_let_tv14_)
                return iife95_(_pat_let12_0)
            d_506_v9_ = (iife86_(iife87_(d_510_v13_))).cf10
            d_507_v10_ = (d_507_v10_).set(default__.safeIndex(d_489_v0_, len(d_507_v10_)), self.f23)
        elif True:
            d_519___mcc_h4_ = source9_.cf19
            d_520_cf19_ = d_519___mcc_h4_
            (globalState).f13 = (self).f24
            d_521_v14_: _dafny.Seq
            d_521_v14_ = _dafny.SeqWithoutIsStrInference([(self).f24, not(p0)])
            d_522_v15_: _dafny.Array
            nw72_ = _dafny.Array(None, 21)
            nw72_[int(0)] = p0
            nw72_[int(1)] = p0
            nw72_[int(2)] = default__.fm2(globalState)
            nw72_[int(3)] = (self).f24
            nw72_[int(4)] = p0
            nw72_[int(5)] = p0
            nw72_[int(6)] = p0
            nw72_[int(7)] = (d_521_v14_)[default__.safeIndex(d_489_v0_, len(d_521_v14_))]
            nw72_[int(8)] = (self).f24
            nw72_[int(9)] = p0
            nw72_[int(10)] = not((self).f24)
            nw72_[int(11)] = False
            nw72_[int(12)] = (self).f24
            nw72_[int(13)] = False
            nw72_[int(14)] = (self).f24
            nw72_[int(15)] = p0
            nw72_[int(16)] = (self).f24
            nw72_[int(17)] = (self).f24
            nw72_[int(18)] = not(p0)
            nw72_[int(19)] = (self).f24
            nw72_[int(20)] = not(False)
            d_522_v15_ = nw72_
            d_523_v16_: _dafny.Array
            nw73_ = _dafny.Array(None, 1)
            nw73_[int(0)] = d_522_v15_
            d_523_v16_ = nw73_
            index65_ = default__.safeIndex(261, (d_523_v16_).length(0))
            (d_523_v16_)[index65_] = d_522_v15_
            index66_ = default__.safeIndex(261, (d_523_v16_).length(0))
            (d_523_v16_)[index66_] = d_522_v15_
            d_524_v17_: C0
            nw74_ = C0()
            nw74_.ctor__(_dafny.Map({False: default__.fm1(p0, globalState)}), False, p1, False)
            d_524_v17_ = nw74_
            d_525_v18_: _dafny.Seq
            d_525_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
            arr0_ = (d_523_v16_)[default__.safeIndex(261, (d_523_v16_).length(0))]
            index67_ = default__.safeIndex(146, ((d_523_v16_)[default__.safeIndex(261, (d_523_v16_).length(0))]).length(0))
            arr0_[index67_] = (p0) and (p0)
            d_526_v19_: _dafny.Set
            d_526_v19_ = _dafny.Set({(d_521_v14_).set(default__.safeIndex(d_489_v0_, len(d_521_v14_)), (self).f24)})
            arr1_ = (d_523_v16_)[default__.safeIndex(261, (d_523_v16_).length(0))]
            index68_ = default__.safeIndex(146, ((d_523_v16_)[default__.safeIndex(261, (d_523_v16_).length(0))]).length(0))
            rhs70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eeheeq"))
            rhs71_ = (_dafny.Set({d_521_v14_})) != (d_526_v19_)
            lhs51_ = (d_523_v16_)[default__.safeIndex(261, (d_523_v16_).length(0))]
            lhs52_ = default__.safeIndex(146, ((d_523_v16_)[default__.safeIndex(261, (d_523_v16_).length(0))]).length(0))
            d_525_v18_ = rhs70_
            lhs51_[lhs52_] = rhs71_
        d_527_v20_: _dafny.Seq
        d_527_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcs"))
        d_528_v21_: _dafny.Map
        d_528_v21_ = _dafny.Map({d_489_v0_: False})
        d_529_v22_: _dafny.Seq
        d_529_v22_ = _dafny.SeqWithoutIsStrInference([len(d_527_v20_), len(d_528_v21_)])
        d_530_v23_: _dafny.Seq
        d_530_v23_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(D4_DC11(d_489_v0_)).cf18 for d_531_i2_ in range(default__.abs(-728))]), d_529_v22_])
        d_532_v24_: _dafny.Seq
        d_532_v24_ = _dafny.SeqWithoutIsStrInference([(self).f24])
        d_533_v25_: _dafny.Array
        nw75_ = _dafny.Array(None, 29)
        nw75_[int(0)] = p0
        nw75_[int(1)] = not(False)
        nw75_[int(2)] = (self).f24
        nw75_[int(3)] = p0
        nw75_[int(4)] = True
        nw75_[int(5)] = (self).f24
        nw75_[int(6)] = (self).f24
        nw75_[int(7)] = default__.fm2(globalState)
        nw75_[int(8)] = (self).f24
        nw75_[int(9)] = (self).f24
        nw75_[int(10)] = p0
        nw75_[int(11)] = p0
        nw75_[int(12)] = not((self).f24)
        nw75_[int(13)] = (self).f24
        nw75_[int(14)] = p0
        nw75_[int(15)] = p0
        nw75_[int(16)] = True
        nw75_[int(17)] = p0
        nw75_[int(18)] = p0
        nw75_[int(19)] = (self).fm4(-538, d_532_v24_, globalState)
        nw75_[int(20)] = p0
        nw75_[int(21)] = p0
        nw75_[int(22)] = p0
        nw75_[int(23)] = p0
        nw75_[int(24)] = p0
        nw75_[int(25)] = (self).f24
        nw75_[int(26)] = (self).f24
        nw75_[int(27)] = (self).f24
        nw75_[int(28)] = True
        d_533_v25_ = nw75_
        d_534_v26_: D1
        d_534_v26_ = D1_DC5(d_489_v0_, d_489_v0_, d_530_v23_, d_533_v25_, (0) - (len(d_490_v1_)))
        d_535_v27_: _dafny.MultiSet
        d_535_v27_ = _dafny.MultiSet([(d_534_v26_).cf10, d_533_v25_, d_533_v25_, d_533_v25_])
        d_536_i1_: int
        d_536_i1_ = 0
        with _dafny.label("1"):
            while not(((d_535_v27_) | (d_535_v27_)).issubset(d_535_v27_)):
                with _dafny.c_label("1"):
                    if (d_536_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_536_i1_ = (d_536_i1_) + (1)
                    d_537_v28_: _dafny.Seq
                    d_537_v28_ = _dafny.SeqWithoutIsStrInference([(d_527_v20_).set(default__.safeIndex(d_489_v0_, len(d_527_v20_)), self.f23), d_527_v20_])
                    if ((d_537_v28_)[default__.safeIndex(d_489_v0_, len(d_537_v28_))]) == (d_527_v20_):
                        (globalState).f13 = (self).f24
                        d_529_v22_ = d_529_v22_
                        d_538_v29_: _dafny.Array
                        def lambda18_(d_539_v0_):
                            def lambda19_(d_540_i3_):
                                return (d_540_i3_) * (d_539_v0_)

                            return lambda19_

                        init10_ = lambda18_(d_489_v0_)
                        nw76_ = _dafny.Array(None, 2)
                        for i0_10_ in range(nw76_.length(0)):
                            nw76_[i0_10_] = init10_(i0_10_)
                        d_538_v29_ = nw76_
                        index69_ = default__.safeIndex(9, (d_538_v29_).length(0))
                        (d_538_v29_)[index69_] = d_489_v0_
                        index70_ = default__.safeIndex(9, (d_538_v29_).length(0))
                        (d_538_v29_)[index70_] = 553
                        d_541_v30_: _dafny.Map
                        d_541_v30_ = _dafny.Map({(self).f24: (d_538_v29_)[default__.safeIndex(9, (d_538_v29_).length(0))]})
                        d_542_v31_: _dafny.Map
                        d_542_v31_ = _dafny.Map({d_541_v30_: p0})
                        d_542_v31_ = (d_542_v31_).set(d_541_v30_, (d_489_v0_) != (len(d_529_v22_)))
                        r0 = ((d_489_v0_) + (d_489_v0_)) > (((d_538_v29_)[default__.safeIndex(9, (d_538_v29_).length(0))]) - ((self).fm24(975, (self).f24, globalState)))
                    elif True:
                        (globalState).f13 = (d_489_v0_) < (d_489_v0_)
                        r1 = default__.fm0((d_527_v20_) < (d_527_v20_), d_489_v0_, globalState)
                        d_543_v32_: D0
                        d_543_v32_ = D0_DC1(d_489_v0_)
                        d_543_v32_ = D0_DC1((d_489_v0_) * (d_489_v0_))
                        d_544_v33_: _dafny.Array
                        nw77_ = _dafny.Array(int(0), 18)
                        d_544_v33_ = nw77_
                        index71_ = default__.safeIndex(415, (d_544_v33_).length(0))
                        (d_544_v33_)[index71_] = (d_489_v0_) + (604)
                        d_545_v34_: _dafny.Map
                        d_545_v34_ = _dafny.Map({d_544_v33_: (d_529_v22_) + (d_529_v22_)})
                        index72_ = default__.safeIndex(939, (d_544_v33_).length(0))
                        (d_544_v33_)[index72_] = len(default__.fm1(p0, globalState))
                        index73_ = default__.safeIndex(415, (d_544_v33_).length(0))
                        index74_ = default__.safeIndex(939, (d_544_v33_).length(0))
                        rhs72_ = (d_489_v0_) + ((d_489_v0_) - (d_489_v0_))
                        rhs73_ = (self).f24
                        rhs74_ = d_545_v34_
                        rhs75_ = len(d_527_v20_)
                        rhs76_ = d_489_v0_
                        lhs53_ = d_544_v33_
                        lhs54_ = default__.safeIndex(415, (d_544_v33_).length(0))
                        lhs55_ = globalState
                        lhs56_ = d_544_v33_
                        lhs57_ = default__.safeIndex(939, (d_544_v33_).length(0))
                        lhs53_[lhs54_] = rhs72_
                        r0 = rhs73_
                        d_545_v34_ = rhs74_
                        lhs55_.f7 = rhs75_
                        lhs56_[lhs57_] = rhs76_
                        d_546_v35_: _dafny.Map
                        d_546_v35_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_547_i4_ in range(default__.abs(389))]): 472})
                        d_546_v35_ = (d_546_v35_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjl")), (d_544_v33_)[default__.safeIndex(415, (d_544_v33_).length(0))])
                    if not(not(p0)):
                        d_548_v36_: _dafny.Map
                        d_548_v36_ = _dafny.Map({d_533_v25_: d_527_v20_})
                        d_549_v37_: D4
                        d_549_v37_ = D4_DC10(((d_548_v36_)[d_533_v25_] if (d_533_v25_) in (d_548_v36_) else d_527_v20_))
                        rhs77_ = (d_489_v0_) - (d_489_v0_)
                        rhs78_ = d_549_v37_
                        rhs79_ = ((0) - (d_489_v0_)) >= (len((d_527_v20_) + (d_527_v20_)))
                        rhs80_ = (0) - (d_489_v0_)
                        rhs81_ = p0
                        lhs58_ = globalState
                        lhs59_ = globalState
                        lhs60_ = globalState
                        lhs61_ = globalState
                        lhs58_.f7 = rhs77_
                        d_549_v37_ = rhs78_
                        lhs59_.f13 = rhs79_
                        lhs60_.f7 = rhs80_
                        lhs61_.f13 = rhs81_
                        (globalState).f7 = (d_489_v0_) - (d_489_v0_)
                        (self).f23 = self.f23
                        r0 = (D2_DC7(not((self).f24), p0)).cf14
                        d_527_v20_ = d_527_v20_
                    elif True:
                        (globalState).f13 = (self).f24
                        d_550_v38_: _dafny.Set
                        d_550_v38_ = _dafny.Set({not((self).f24)})
                        d_551_v39_: _dafny.Map
                        d_551_v39_ = _dafny.Map({p1: p0})
                        d_552_v40_: _dafny.Map
                        d_552_v40_ = _dafny.Map({(self).f24: len((_dafny.Map({_dafny.CodePoint('h'): (self).f24})) | (d_551_v39_))})
                        d_553_v41_: D0
                        d_553_v41_ = D0_DC1(d_489_v0_)
                        d_554_v42_: D0
                        d_554_v42_ = D0_DC3(d_553_v41_)
                        d_555_v43_: D0
                        d_555_v43_ = D0_DC3(d_554_v42_)
                        d_556_v44_: _dafny.Array
                        nw78_ = _dafny.Array(D4.default()(), 9)
                        d_556_v44_ = nw78_
                        d_557_v45_: D4
                        d_557_v45_ = D4_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgil")))
                        index75_ = default__.safeIndex(306, (d_556_v44_).length(0))
                        (d_556_v44_)[index75_] = d_557_v45_
                        d_558_v46_: _dafny.Map
                        d_558_v46_ = _dafny.Map({(self).f24: (self).f24})
                        index76_ = default__.safeIndex(306, (d_556_v44_).length(0))
                        rhs82_ = d_550_v38_
                        rhs83_ = _dafny.Map({(d_489_v0_) < (len(d_558_v46_)): (d_489_v0_) - (d_489_v0_)})
                        rhs84_ = d_555_v43_
                        rhs85_ = d_557_v45_
                        lhs62_ = d_556_v44_
                        lhs63_ = default__.safeIndex(306, (d_556_v44_).length(0))
                        d_550_v38_ = rhs82_
                        d_552_v40_ = rhs83_
                        d_555_v43_ = rhs84_
                        lhs62_[lhs63_] = rhs85_
                        d_489_v0_ = d_489_v0_
                        (globalState).f7 = d_489_v0_
                        d_559_v47_: _dafny.Array
                        nw79_ = _dafny.Array(_dafny.Map({}), 16)
                        d_559_v47_ = nw79_
                        d_560_v48_: _dafny.Array
                        def lambda20_(d_561_i5_):
                            return (d_561_i5_) + (291)

                        init11_ = lambda20_
                        nw80_ = _dafny.Array(None, 7)
                        for i0_11_ in range(nw80_.length(0)):
                            nw80_[i0_11_] = init11_(i0_11_)
                        d_560_v48_ = nw80_
                        d_562_v49_: _dafny.Map
                        d_562_v49_ = _dafny.Map({d_560_v48_: d_489_v0_})
                        index77_ = default__.safeIndex(359, (d_559_v47_).length(0))
                        (d_559_v47_)[index77_] = d_562_v49_
                        index78_ = default__.safeIndex(359, (d_559_v47_).length(0))
                        rhs86_ = p0
                        rhs87_ = ((((d_562_v49_).set(d_560_v48_, len(d_558_v46_))).set(d_560_v48_, d_489_v0_) if default__.fm2(globalState) else d_562_v49_)) | (d_562_v49_)
                        lhs64_ = d_559_v47_
                        lhs65_ = default__.safeIndex(359, (d_559_v47_).length(0))
                        r2 = rhs86_
                        lhs64_[lhs65_] = rhs87_
                    d_563_v50_: _dafny.Map
                    d_563_v50_ = _dafny.Map({p0: (self).f24})
                    d_564_v51_: D0
                    d_564_v51_ = D0_DC2(p0, d_489_v0_, ((d_563_v50_)[p0] if (p0) in (d_563_v50_) else not(p0)))
                    index79_ = default__.safeIndex(349, (d_533_v25_).length(0))
                    (d_533_v25_)[index79_] = (d_564_v51_).cf4
                    index80_ = default__.safeIndex(349, (d_533_v25_).length(0))
                    (d_533_v25_)[index80_] = (self).f24
                    r2 = p0
                    pass
            pass
        d_565_v52_: _dafny.Array
        def lambda21_(d_566_v0_):
            def lambda22_(d_567_i6_):
                return _dafny.MultiSet([d_566_v0_])

            return lambda22_

        init12_ = lambda21_(d_489_v0_)
        nw81_ = _dafny.Array(None, 3)
        for i0_12_ in range(nw81_.length(0)):
            nw81_[i0_12_] = init12_(i0_12_)
        d_565_v52_ = nw81_
        d_568_v53_: _dafny.Map
        d_568_v53_ = _dafny.Map({(self).f24: (self).f24})
        d_569_v54_: _dafny.MultiSet
        d_569_v54_ = _dafny.MultiSet([d_489_v0_, len(d_568_v53_)])
        index81_ = default__.safeIndex(944, (d_565_v52_).length(0))
        (d_565_v52_)[index81_] = (_dafny.MultiSet([d_489_v0_])).intersection(d_569_v54_)
        index82_ = default__.safeIndex(944, (d_565_v52_).length(0))
        (d_565_v52_)[index82_] = (_dafny.MultiSet(d_529_v22_)).intersection(d_569_v54_)
        hi2_ = ((0) - (d_489_v0_) if (self).f24 else (0) - (d_489_v0_))
        for d_570_i7_ in range(d_489_v0_, hi2_):
            (globalState).f13 = (d_490_v1_).isdisjoint(d_490_v1_)
            r0 = (self).f24
            d_571_v55_: _dafny.Array
            nw82_ = _dafny.Array(_dafny.Map({}), 16)
            d_571_v55_ = nw82_
            d_572_v56_: D5
            d_572_v56_ = D5_DC13(d_489_v0_, (self).f24, p0)
            d_573_v57_: _dafny.Set
            d_573_v57_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_572_v56_, d_572_v56_])})
            d_574_v58_: _dafny.Map
            d_574_v58_ = _dafny.Map({(self).f24: d_573_v57_})
            rhs88_ = (default__.fm27((self).f24, p0, globalState)) == (((d_574_v58_)[p0] if (p0) in (d_574_v58_) else d_573_v57_))
            rhs89_ = self.f23
            rhs90_ = (d_571_v55_ if (self).f24 else d_571_v55_)
            lhs66_ = globalState
            lhs67_ = self
            lhs66_.f13 = rhs88_
            lhs67_.f23 = rhs89_
            d_571_v55_ = rhs90_
            d_575_v59_: C0
            nw83_ = C0()
            nw83_.ctor__((_dafny.Map({p0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owqe"))})).set((self).f24, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdpkd"))), (self).f24, self.f23, True)
            d_575_v59_ = nw83_
        d_576_v60_: _dafny.Array
        nw84_ = _dafny.Array(int(0), 14)
        d_576_v60_ = nw84_
        index83_ = default__.safeIndex(370, (d_576_v60_).length(0))
        (d_576_v60_)[index83_] = d_489_v0_
        index84_ = default__.safeIndex(370, (d_576_v60_).length(0))
        (d_576_v60_)[index84_] = default__.safeDivisionInt((d_569_v54_).cardinality, d_489_v0_)
        r1 = len(d_527_v20_)
        r0 = (_dafny.MultiSet([False])).ispropersubset((self).f36)
        r1 = d_489_v0_
        r2 = (self).f24
        return r0, r1, r2

    def m12(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: int = int(0)
        d_577_v0_: _dafny.Seq
        d_577_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulpqjjbk"))
        r1 = (d_577_v0_) + (_dafny.SeqWithoutIsStrInference([self.f23 for d_578_i0_ in range(default__.abs(783))]))
        d_579_v1_: _dafny.Seq
        d_579_v1_ = _dafny.SeqWithoutIsStrInference([(self).f24])
        rhs91_ = (p1) + (p1)
        rhs92_ = len(d_579_v1_)
        lhs68_ = globalState
        lhs69_ = globalState
        lhs68_.f7 = rhs91_
        lhs69_.f7 = rhs92_
        d_580_v2_: D1
        d_580_v2_ = D1_DC4((d_577_v0_)[default__.safeIndex(p1, len(d_577_v0_))])
        source10_ = d_580_v2_
        if source10_.is_DC5:
            d_581___mcc_h0_ = source10_.cf7
            d_582___mcc_h1_ = source10_.cf8
            d_583___mcc_h2_ = source10_.cf9
            d_584___mcc_h3_ = source10_.cf10
            d_585___mcc_h4_ = source10_.cf11
            d_586_cf11_ = d_585___mcc_h4_
            d_587_cf10_ = d_584___mcc_h3_
            d_588_cf9_ = d_583___mcc_h2_
            d_589_cf8_ = d_582___mcc_h1_
            d_590_cf7_ = d_581___mcc_h0_
            d_591_v3_: _dafny.Seq
            d_591_v3_ = _dafny.SeqWithoutIsStrInference([d_586_cf11_])
            d_592_v4_: _dafny.Set
            d_592_v4_ = _dafny.Set({d_590_cf7_, (0) - (p1), (d_591_v3_)[default__.safeIndex(d_586_cf11_, len(d_591_v3_))]})
            rhs93_ = default__.fm0((self).f24, d_590_cf7_, globalState)
            rhs94_ = default__.safeDivisionInt((0) - (len(d_592_v4_)), len(d_579_v1_))
            rhs95_ = (p1) < ((0) - (d_589_cf8_))
            lhs70_ = globalState
            lhs71_ = globalState
            d_586_cf11_ = rhs93_
            lhs70_.f7 = rhs94_
            lhs71_.f13 = rhs95_
            d_593_v5_: D0
            d_593_v5_ = D0_DC2((self).f24, d_590_cf7_, (self).f24)
            d_594_v6_: _dafny.Map
            d_594_v6_ = _dafny.Map({self.f23: d_590_cf7_})
            pat_let_tv18_ = d_594_v6_
            def iife100_(_pat_let19_0):
                def iife101_(d_595_dt__update__tmp_h0_):
                    def iife102_(_pat_let20_0):
                        def iife103_(d_596_dt__update_hcf3_h0_):
                            def iife104_(_pat_let21_0):
                                def iife105_(d_597_dt__update_hcf4_h0_):
                                    return D0_DC2((d_595_dt__update__tmp_h0_).cf2, d_596_dt__update_hcf3_h0_, d_597_dt__update_hcf4_h0_)
                                return iife105_(_pat_let21_0)
                            return iife104_(True)
                        return iife103_(_pat_let20_0)
                    return iife102_(len(pat_let_tv18_))
                return iife101_(_pat_let19_0)
            d_593_v5_ = iife100_(d_593_v5_)
            d_598_v7_: _dafny.Map
            d_598_v7_ = _dafny.Map({not(False): d_577_v0_})
            d_599_v8_: D2
            d_599_v8_ = D2_DC6(d_598_v7_)
            d_599_v8_ = (d_599_v8_ if (self).f24 else D2_DC6(d_598_v7_))
            d_600_v9_: _dafny.Map
            d_600_v9_ = _dafny.Map({(self).f24: -372})
            d_601_v10_: _dafny.Set
            d_601_v10_ = _dafny.Set({_dafny.Map({(self).f24: d_589_cf8_}), d_600_v9_, d_600_v9_, d_600_v9_})
            source11_ = default__.fm28(d_577_v0_, d_591_v3_, (d_579_v1_)[default__.safeIndex(len((d_577_v0_).set(default__.safeIndex(len(d_577_v0_), len(d_577_v0_)), self.f23)), len(d_579_v1_))], d_601_v10_, globalState)
            if source11_.is_DC1:
                d_602___mcc_h6_ = source11_.cf1
                d_603_cf1_ = d_602___mcc_h6_
                d_604_v11_: C0
                nw85_ = C0()
                nw85_.ctor__(d_598_v7_, (D0_DC2(True, d_590_cf7_, (self).f24)).cf4, self.f23, (self).f24)
                d_604_v11_ = nw85_
                d_605_v12_: _dafny.Seq
                d_605_v12_ = _dafny.SeqWithoutIsStrInference([d_587_cf10_])
                d_606_v13_: _dafny.Map
                d_606_v13_ = _dafny.Map({default__.fm0((d_604_v11_).f38, d_589_cf8_, globalState): (d_605_v12_)[default__.safeIndex(len(d_592_v4_), len(d_605_v12_))]})
                d_606_v13_ = (d_606_v13_).set((d_591_v3_)[default__.safeIndex(d_586_cf11_, len(d_591_v3_))], d_587_cf10_)
                d_607_v14_: _dafny.Map
                d_607_v14_ = _dafny.Map({len(d_577_v0_): d_586_cf11_})
                d_607_v14_ = (d_607_v14_).set((843) - (p1), 474)
                d_608_v15_: _dafny.Map
                d_608_v15_ = _dafny.Map({len(d_600_v9_): d_579_v1_})
                d_608_v15_ = (d_608_v15_).set((0) - (d_589_cf8_), (d_579_v1_) + (d_579_v1_))
            elif source11_.is_DC2:
                d_609___mcc_h7_ = source11_.cf2
                d_610___mcc_h8_ = source11_.cf3
                d_611___mcc_h9_ = source11_.cf4
                d_612_cf4_ = d_611___mcc_h9_
                d_613_cf3_ = d_610___mcc_h8_
                d_614_cf2_ = d_609___mcc_h7_
                (globalState).f7 = default__.safeModuloInt((d_590_cf7_) + (d_586_cf11_), d_590_cf7_)
                (globalState).f7 = d_613_cf3_
                d_615_v16_: _dafny.Map
                d_615_v16_ = _dafny.Map({p1: d_586_cf11_})
                pat_let_tv19_ = p1
                pat_let_tv20_ = d_614_cf2_
                def iife106_(_pat_let22_0):
                    def iife107_(d_616_dt__update__tmp_h1_):
                        def iife108_(_pat_let23_0):
                            def iife109_(d_617_dt__update_hcf3_h1_):
                                def iife110_(_pat_let24_0):
                                    def iife111_(d_618_dt__update_hcf4_h1_):
                                        return D0_DC2((d_616_dt__update__tmp_h1_).cf2, d_617_dt__update_hcf3_h1_, d_618_dt__update_hcf4_h1_)
                                    return iife111_(_pat_let24_0)
                                return iife110_(pat_let_tv20_)
                            return iife109_(_pat_let23_0)
                        return iife108_(pat_let_tv19_)
                    return iife107_(_pat_let22_0)
                def iife112_():
                    coll62_ = _dafny.Map()
                    compr_62_: int
                    for compr_62_ in (d_591_v3_).Elements:
                        d_619_v17_: int = compr_62_
                        if (d_619_v17_) in (d_591_v3_):
                            coll62_[default__.safeModuloInt(d_619_v17_, d_613_cf3_)] = d_586_cf11_
                    return _dafny.Map(coll62_)
                rhs96_ = (self.f23 if (iife106_(d_593_v5_)).cf4 else default__.fm29(d_613_cf3_, d_612_cf4_, d_614_cf2_, d_586_cf11_, globalState))
                rhs97_ = (iife112_()
                ) | ((d_615_v16_) | (_dafny.Map({d_589_cf8_: d_586_cf11_})))
                lhs72_ = self
                lhs72_.f23 = rhs96_
                d_615_v16_ = rhs97_
                rhs98_ = (d_577_v0_) + (((d_577_v0_) + (d_577_v0_)).set(default__.safeIndex(d_589_cf8_, len((d_577_v0_) + (d_577_v0_))), _dafny.CodePoint('v')))
                rhs99_ = self.f23
                lhs73_ = self
                r1 = rhs98_
                lhs73_.f23 = rhs99_
            elif source11_.is_DC0:
                d_620___mcc_h10_ = source11_.cf0
                d_621_cf0_ = d_620___mcc_h10_
                d_577_v0_ = (d_577_v0_).set(default__.safeIndex(d_589_cf8_, len(d_577_v0_)), _dafny.CodePoint('a'))
                d_622_v18_: _dafny.Array
                def lambda23_(d_623_v2_):
                    def lambda24_(d_624_i1_):
                        return (d_624_i1_) * (len(_dafny.Set({(d_623_v2_).cf6, self.f23})))

                    return lambda24_

                init13_ = lambda23_(d_580_v2_)
                nw86_ = _dafny.Array(None, 11)
                for i0_13_ in range(nw86_.length(0)):
                    nw86_[i0_13_] = init13_(i0_13_)
                d_622_v18_ = nw86_
                rhs100_ = ((0) - (d_589_cf8_)) * ((0) - (default__.safeModuloInt(d_586_cf11_, d_586_cf11_)))
                rhs101_ = (False) in ((d_579_v1_) + (d_579_v1_))
                rhs102_ = (d_590_cf7_) - ((0) - (d_590_cf7_))
                rhs103_ = d_622_v18_
                lhs74_ = globalState
                lhs75_ = globalState
                lhs76_ = globalState
                lhs77_ = globalState
                lhs74_.f7 = rhs100_
                lhs75_.f13 = rhs101_
                lhs76_.f7 = rhs102_
                lhs77_.f22 = rhs103_
                d_625_v19_: D2
                d_625_v19_ = D2_DC7((self).f24, (self).f24)
                d_626_v20_: C0
                nw87_ = C0()
                nw87_.ctor__(d_598_v7_, (d_625_v19_).cf14, self.f23, (self).f24)
                d_626_v20_ = nw87_
                d_627_v21_: _dafny.MultiSet
                d_627_v21_ = _dafny.MultiSet([len(d_577_v0_)])
                d_628_v22_: _dafny.Set
                d_628_v22_ = _dafny.Set({(d_627_v21_) - (d_627_v21_)})
                d_629_v23_: _dafny.Seq
                d_629_v23_ = _dafny.SeqWithoutIsStrInference([d_577_v0_, d_577_v0_])
                rhs104_ = default__.fm30((d_599_v8_ if (self).f24 else d_599_v8_), d_629_v23_, globalState)
                rhs105_ = (0) - (d_590_cf7_)
                d_628_v22_ = rhs104_
                r0 = rhs105_
            elif True:
                d_630___mcc_h11_ = source11_.cf5
                d_631_cf5_ = d_630___mcc_h11_
                d_586_cf11_ = d_590_cf7_
                d_632_v24_: _dafny.MultiSet
                d_632_v24_ = _dafny.MultiSet([len(d_579_v1_), d_590_cf7_, p1])
                d_633_v25_: _dafny.Seq
                d_633_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_590_cf7_]), d_632_v24_, d_632_v24_])
                d_634_v26_: D5
                d_634_v26_ = D5_DC12(_dafny.Set({p1}))
                d_635_v27_: _dafny.Array
                nw88_ = _dafny.Array(None, 15)
                nw88_[int(0)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogv"))) + (d_577_v0_))
                nw88_[int(1)] = d_589_cf8_
                nw88_[int(2)] = len(d_592_v4_)
                nw88_[int(3)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcybk")))
                nw88_[int(4)] = (len((d_634_v26_).cf19)) + (((d_632_v24_)[d_590_cf7_] if (d_590_cf7_) in (d_632_v24_) else len(d_577_v0_)))
                nw88_[int(5)] = 276
                nw88_[int(6)] = d_589_cf8_
                nw88_[int(7)] = (len(d_591_v3_)) + (((d_600_v9_)[(d_579_v1_)[default__.safeIndex(len(d_591_v3_), len(d_579_v1_))]] if ((d_579_v1_)[default__.safeIndex(len(d_591_v3_), len(d_579_v1_))]) in (d_600_v9_) else p1))
                nw88_[int(8)] = p1
                nw88_[int(9)] = d_590_cf7_
                nw88_[int(10)] = p1
                nw88_[int(11)] = (0) - ((p1) + (p1))
                nw88_[int(12)] = default__.fm0(False, d_589_cf8_, globalState)
                nw88_[int(13)] = d_586_cf11_
                nw88_[int(14)] = len((D4_DC10(d_577_v0_)).cf17)
                d_635_v27_ = nw88_
                index85_ = default__.safeIndex(36, (d_635_v27_).length(0))
                (d_635_v27_)[index85_] = p1
                d_636_v28_: D5
                d_636_v28_ = D5_DC13(d_586_cf11_, True, (self).f24)
                index86_ = default__.safeIndex(578, (d_587_cf10_).length(0))
                def iife113_(_pat_let25_0):
                    def iife114_(d_637_dt__update__tmp_h2_):
                        def iife115_(_pat_let26_0):
                            def iife116_(d_638_dt__update_hcf21_h0_):
                                return D5_DC13((d_637_dt__update__tmp_h2_).cf20, d_638_dt__update_hcf21_h0_, (d_637_dt__update__tmp_h2_).cf22)
                            return iife116_(_pat_let26_0)
                        return iife115_((self).f24)
                    return iife114_(_pat_let25_0)
                (d_587_cf10_)[index86_] = (iife113_(d_636_v28_)).cf21
                d_639_v29_: _dafny.Map
                d_639_v29_ = _dafny.Map({(self).f24: d_587_cf10_})
                d_640_v30_: _dafny.Set
                d_640_v30_ = _dafny.Set({(self).f24})
                index87_ = default__.safeIndex(36, (d_635_v27_).length(0))
                index88_ = default__.safeIndex(578, (d_587_cf10_).length(0))
                rhs106_ = default__.fm31((self).f24, not((self).f24), (self).f24, d_593_v5_, globalState)
                rhs107_ = (default__.safeDivisionInt(d_590_cf7_, d_590_cf7_)) * ((len(d_639_v29_)) - (d_586_cf11_))
                rhs108_ = (d_640_v30_).isdisjoint((d_640_v30_).intersection(_dafny.Set({(d_579_v1_)[default__.safeIndex(d_589_cf8_, len(d_579_v1_))]})))
                lhs78_ = d_635_v27_
                lhs79_ = default__.safeIndex(36, (d_635_v27_).length(0))
                lhs80_ = d_587_cf10_
                lhs81_ = default__.safeIndex(578, (d_587_cf10_).length(0))
                d_633_v25_ = rhs106_
                lhs78_[lhs79_] = rhs107_
                lhs80_[lhs81_] = rhs108_
                d_641_v31_: _dafny.Array
                nw89_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
                d_641_v31_ = nw89_
                index89_ = default__.safeIndex(445, (d_641_v31_).length(0))
                (d_641_v31_)[index89_] = d_577_v0_
                index90_ = default__.safeIndex(445, (d_641_v31_).length(0))
                (d_641_v31_)[index90_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxxltvc"))
                d_642_v32_: _dafny.Map
                d_642_v32_ = _dafny.Map({d_632_v24_: d_586_cf11_})
                d_643_v33_: _dafny.Map
                d_643_v33_ = _dafny.Map({d_586_cf11_: d_586_cf11_})
                d_642_v32_ = (d_642_v32_).set(d_632_v24_, default__.fm0(not((d_587_cf10_)[default__.safeIndex(578, (d_587_cf10_).length(0))]), len(d_643_v33_), globalState))
        elif True:
            d_644___mcc_h5_ = source10_.cf6
            d_645_cf6_ = d_644___mcc_h5_
            source12_ = D2_DC7((self).f24, (self).f24)
            if source12_.is_DC7:
                d_646___mcc_h12_ = source12_.cf13
                d_647___mcc_h13_ = source12_.cf14
                d_648_cf14_ = d_647___mcc_h13_
                d_649_cf13_ = d_646___mcc_h12_
                d_650_v34_: D4
                d_650_v34_ = D4_DC11(p1)
                pat_let_tv21_ = p1
                pat_let_tv22_ = p1
                def iife117_(_pat_let27_0):
                    def iife118_(d_651_dt__update__tmp_h3_):
                        def iife119_(_pat_let28_0):
                            def iife120_(d_652_dt__update_hcf18_h0_):
                                return D4_DC11(d_652_dt__update_hcf18_h0_)
                            return iife120_(_pat_let28_0)
                        return iife119_((pat_let_tv21_) + (pat_let_tv22_))
                    return iife118_(_pat_let27_0)
                d_650_v34_ = iife117_(D4_DC11(p1))
                d_653_v35_: _dafny.Map
                d_653_v35_ = _dafny.Map({(self).f24: p1})
                d_653_v35_ = (d_653_v35_).set(not (not(True)) or (d_648_cf14_), p1)
                d_654_v36_: _dafny.Seq
                d_654_v36_ = _dafny.SeqWithoutIsStrInference([p1])
                d_655_v37_: T3
                nw90_ = C1()
                nw90_.ctor__(True, d_649_cf13_)
                d_655_v37_ = nw90_
                d_656_v38_: _dafny.Array
                nw91_ = _dafny.Array(None, 13)
                nw91_[int(0)] = d_655_v37_
                nw91_[int(1)] = d_655_v37_
                nw91_[int(2)] = d_655_v37_
                nw91_[int(3)] = d_655_v37_
                nw91_[int(4)] = d_655_v37_
                nw91_[int(5)] = d_655_v37_
                nw91_[int(6)] = d_655_v37_
                nw91_[int(7)] = d_655_v37_
                nw91_[int(8)] = d_655_v37_
                nw91_[int(9)] = d_655_v37_
                nw91_[int(10)] = d_655_v37_
                nw91_[int(11)] = d_655_v37_
                nw91_[int(12)] = d_655_v37_
                d_656_v38_ = nw91_
                d_657_v39_: _dafny.Map
                d_657_v39_ = _dafny.Map({d_656_v38_: self.f23})
                rhs109_ = (_dafny.SeqWithoutIsStrInference([-18]) if d_648_cf14_ else (d_654_v36_) + (d_654_v36_))
                rhs110_ = d_657_v39_
                d_654_v36_ = rhs109_
                d_657_v39_ = rhs110_
                d_658_v40_: _dafny.Array
                nw92_ = _dafny.Array(None, 18)
                nw92_[int(0)] = d_579_v1_
                nw92_[int(1)] = (d_579_v1_ if d_649_cf13_ else _dafny.SeqWithoutIsStrInference([(d_655_v37_).f27]))
                nw92_[int(2)] = d_579_v1_
                nw92_[int(3)] = d_579_v1_
                nw92_[int(4)] = d_579_v1_
                nw92_[int(5)] = d_579_v1_
                nw92_[int(6)] = d_579_v1_
                nw92_[int(7)] = d_579_v1_
                nw92_[int(8)] = d_579_v1_
                nw92_[int(9)] = (_dafny.SeqWithoutIsStrInference([(d_655_v37_).f27, (self).f24])) + (_dafny.SeqWithoutIsStrInference([d_648_cf14_]))
                nw92_[int(10)] = (_dafny.SeqWithoutIsStrInference([(self).f24])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([(self).f24]))), (d_655_v37_).f27)
                nw92_[int(11)] = (_dafny.SeqWithoutIsStrInference([default__.fm2(globalState)])) + (d_579_v1_)
                nw92_[int(12)] = (_dafny.SeqWithoutIsStrInference([True, d_649_cf13_])) + (d_579_v1_)
                nw92_[int(13)] = d_579_v1_
                nw92_[int(14)] = (d_579_v1_) + (d_579_v1_)
                nw92_[int(15)] = default__.fm36(globalState)
                nw92_[int(16)] = d_579_v1_
                nw92_[int(17)] = (d_579_v1_) + (_dafny.SeqWithoutIsStrInference([(d_655_v37_).f27]))
                d_658_v40_ = nw92_
                d_659_v41_: _dafny.Map
                d_659_v41_ = _dafny.Map({440: d_579_v1_})
                d_660_v42_: _dafny.Map
                d_660_v42_ = _dafny.Map({(self).f24: (d_655_v37_).f27})
                d_661_v43_: D6
                d_661_v43_ = D6_DC16(d_660_v42_, (d_655_v37_).f27, (d_655_v37_).f27)
                d_662_v44_: _dafny.Set
                d_662_v44_ = _dafny.Set({(self).f25})
                d_663_v45_: _dafny.Array
                nw93_ = _dafny.Array(None, 27)
                nw93_[int(0)] = d_648_cf14_
                nw93_[int(1)] = (d_645_cf6_) in (d_577_v0_)
                nw93_[int(2)] = d_648_cf14_
                nw93_[int(3)] = (self).f24
                nw93_[int(4)] = d_648_cf14_
                nw93_[int(5)] = ((self).f36).ispropersubset(_dafny.MultiSet([False, (self).f24]))
                nw93_[int(6)] = (d_659_v41_) != (d_659_v41_)
                nw93_[int(7)] = (d_648_cf14_) == (d_649_cf13_)
                nw93_[int(8)] = (d_661_v43_).cf27
                nw93_[int(9)] = False
                nw93_[int(10)] = not(d_648_cf14_)
                nw93_[int(11)] = (D2_DC7((d_655_v37_).f27, (d_655_v37_).f27)).cf13
                nw93_[int(12)] = d_649_cf13_
                nw93_[int(13)] = not((d_662_v44_).ispropersubset(d_662_v44_))
                nw93_[int(14)] = (self).f24
                nw93_[int(15)] = (d_655_v37_).f27
                nw93_[int(16)] = (self).f24
                nw93_[int(17)] = (d_655_v37_).f27
                nw93_[int(18)] = (self).f24
                nw93_[int(19)] = True
                nw93_[int(20)] = (d_655_v37_).f27
                nw93_[int(21)] = not(False)
                nw93_[int(22)] = (self).f24
                nw93_[int(23)] = (self).f24
                nw93_[int(24)] = not((p1) > (p1))
                nw93_[int(25)] = True
                nw93_[int(26)] = (self).f24
                d_663_v45_ = nw93_
                d_664_v46_: _dafny.Map
                d_664_v46_ = _dafny.Map({(0) - (p1): d_658_v40_})
                rhs111_ = ((d_664_v46_)[default__.safeDivisionInt(675, p1)] if (default__.safeDivisionInt(675, p1)) in (d_664_v46_) else d_658_v40_)
                rhs112_ = True
                rhs113_ = d_663_v45_
                lhs82_ = globalState
                d_658_v40_ = rhs111_
                lhs82_.f13 = rhs112_
                d_663_v45_ = rhs113_
            elif source12_.is_DC6:
                d_665___mcc_h14_ = source12_.cf12
                d_666_cf12_ = d_665___mcc_h14_
                d_667_v47_: _dafny.Map
                d_667_v47_ = _dafny.Map({d_577_v0_: len(d_577_v0_)})
                r0 = ((d_667_v47_)[d_577_v0_] if (d_577_v0_) in (d_667_v47_) else p1)
                (globalState).f13 = not((self).f24)
                d_668_v48_: _dafny.Array
                nw94_ = _dafny.Array(int(0), 19)
                d_668_v48_ = nw94_
                index91_ = default__.safeIndex(369, (d_668_v48_).length(0))
                (d_668_v48_)[index91_] = p1
                index92_ = default__.safeIndex(369, (d_668_v48_).length(0))
                (d_668_v48_)[index92_] = default__.safeDivisionInt(p1, p1)
                index93_ = default__.safeIndex(369, (d_668_v48_).length(0))
                (d_668_v48_)[index93_] = (d_668_v48_)[default__.safeIndex(369, (d_668_v48_).length(0))]
            elif True:
                d_669___mcc_h15_ = source12_.cf15
                d_670_cf15_ = d_669___mcc_h15_
                d_671_v49_: _dafny.MultiSet
                d_671_v49_ = _dafny.MultiSet([535])
                d_672_v50_: _dafny.Map
                d_672_v50_ = _dafny.Map({(self).f24: True})
                d_673_v51_: _dafny.Map
                d_673_v51_ = _dafny.Map({d_672_v50_: -793})
                d_674_v52_: _dafny.Array
                nw95_ = _dafny.Array(None, 16)
                nw95_[int(0)] = p1
                nw95_[int(1)] = (0) - (p1)
                nw95_[int(2)] = p1
                nw95_[int(3)] = p1
                nw95_[int(4)] = (((self).f36)[(self).f24] if ((self).f24) in ((self).f36) else -489)
                nw95_[int(5)] = (d_671_v49_).cardinality
                nw95_[int(6)] = p1
                nw95_[int(7)] = p1
                nw95_[int(8)] = default__.fm0(True, len(d_577_v0_), globalState)
                nw95_[int(9)] = 529
                nw95_[int(10)] = ((d_673_v51_)[_dafny.Map({(self).f24: True})] if (_dafny.Map({(self).f24: True})) in (d_673_v51_) else 444)
                nw95_[int(11)] = p1
                nw95_[int(12)] = (0) - ((_dafny.MultiSet([(d_579_v1_)[default__.safeIndex(p1, len(d_579_v1_))]])).cardinality)
                nw95_[int(13)] = p1
                nw95_[int(14)] = p1
                nw95_[int(15)] = p1
                d_674_v52_ = nw95_
                d_675_v53_: _dafny.Map
                d_675_v53_ = _dafny.Map({d_674_v52_: (self).f24})
                d_675_v53_ = d_675_v53_
                d_676_v54_: D6
                d_676_v54_ = D6_DC16(_dafny.Map({False: (self).f24}), (self).f24, (self).f24)
                d_677_v55_: _dafny.Map
                d_677_v55_ = _dafny.Map({True: p1})
                d_678_v56_: _dafny.Set
                d_678_v56_ = _dafny.Set({p1, p1, p1, len(_dafny.Map({False: d_674_v52_})), p1})
                pat_let_tv23_ = d_672_v50_
                d_679_v57_: _dafny.Array
                nw96_ = _dafny.Array(None, 20)
                nw96_[int(0)] = (self).f24
                nw96_[int(1)] = (self).f24
                nw96_[int(2)] = (self).f24
                nw96_[int(3)] = (self).f24
                nw96_[int(4)] = (self).f24
                nw96_[int(5)] = (self).f24
                nw96_[int(6)] = (self).f24
                nw96_[int(7)] = (self).f24
                nw96_[int(8)] = (self).f24
                nw96_[int(9)] = ((0) - (p1)) not in (d_671_v49_)
                nw96_[int(10)] = (self).f24
                nw96_[int(11)] = (self).f24
                nw96_[int(12)] = (self).f24
                def iife121_(_pat_let29_0):
                    def iife122_(d_680_dt__update__tmp_h4_):
                        def iife123_(_pat_let30_0):
                            def iife124_(d_681_dt__update_hcf25_h0_):
                                return D6_DC16(d_681_dt__update_hcf25_h0_, (d_680_dt__update__tmp_h4_).cf26, (d_680_dt__update__tmp_h4_).cf27)
                            return iife124_(_pat_let30_0)
                        return iife123_(pat_let_tv23_)
                    return iife122_(_pat_let29_0)
                nw96_[int(13)] = not ((iife121_(d_676_v54_)).cf27) or ((self).f24)
                nw96_[int(14)] = (p1) > ((0) - (((d_677_v55_)[(self).f24] if ((self).f24) in (d_677_v55_) else p1)))
                nw96_[int(15)] = (self).f24
                nw96_[int(16)] = (p1) != (len(d_678_v56_))
                nw96_[int(17)] = not((self).f24)
                nw96_[int(18)] = ((self).f36).ispropersubset((self).f36)
                nw96_[int(19)] = (self).f24
                d_679_v57_ = nw96_
                index94_ = default__.safeIndex(281, (d_679_v57_).length(0))
                (d_679_v57_)[index94_] = (self).f24
                index95_ = default__.safeIndex(281, (d_679_v57_).length(0))
                (d_679_v57_)[index95_] = (False) == (False)
                d_682_v58_: _dafny.Seq
                d_682_v58_ = _dafny.SeqWithoutIsStrInference([d_674_v52_, d_674_v52_, d_674_v52_])
                d_683_v59_: _dafny.Map
                d_683_v59_ = _dafny.Map({p1: (self).f24})
                def iife125_():
                    coll63_ = _dafny.Set()
                    compr_63_: int
                    for compr_63_ in (d_683_v59_).keys.Elements:
                        d_684_v60_: int = compr_63_
                        if (d_684_v60_) in (d_683_v59_):
                            coll63_ = coll63_.union(_dafny.Set([(d_684_v60_) * (len(_dafny.Map({-807: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mo")))])})))]))
                    return _dafny.Set(coll63_)
                d_682_v58_ = (d_682_v58_ if (iife125_()
                ).issubset(d_678_v56_) else (d_682_v58_) + (d_682_v58_))
                (globalState).f13 = (d_679_v57_)[default__.safeIndex(281, (d_679_v57_).length(0))]
            d_685_v61_: C1
            nw97_ = C1()
            nw97_.ctor__((self).f24, (self).f24)
            d_685_v61_ = nw97_
            d_686_v62_: _dafny.Array
            def lambda25_(d_687_v0_):
                def lambda26_(d_688_i2_):
                    return d_687_v0_

                return lambda26_

            init14_ = lambda25_(d_577_v0_)
            nw98_ = _dafny.Array(None, 25)
            for i0_14_ in range(nw98_.length(0)):
                nw98_[i0_14_] = init14_(i0_14_)
            d_686_v62_ = nw98_
            index96_ = default__.safeIndex(249, (d_686_v62_).length(0))
            (d_686_v62_)[index96_] = d_577_v0_
            d_689_v63_: D4
            d_689_v63_ = D4_DC10(d_577_v0_)
            index97_ = default__.safeIndex(249, (d_686_v62_).length(0))
            (d_686_v62_)[index97_] = ((d_689_v63_ if (self).f24 else d_689_v63_)).cf17
            if default__.fm2(globalState):
                d_690_v64_: _dafny.Array
                def lambda27_(d_691_p1_):
                    def lambda28_(d_692_i3_):
                        return default__.safeDivisionInt(d_692_i3_, d_691_p1_)

                    return lambda28_

                init15_ = lambda27_(p1)
                nw99_ = _dafny.Array(None, 23)
                for i0_15_ in range(nw99_.length(0)):
                    nw99_[i0_15_] = init15_(i0_15_)
                d_690_v64_ = nw99_
                d_693_v65_: _dafny.Map
                d_693_v65_ = _dafny.Map({p1: d_690_v64_})
                d_693_v65_ = (d_693_v65_).set(p1, d_690_v64_)
                (globalState).f13 = ((d_685_v61_).f39) not in ((self).f36)
                d_694_v66_: _dafny.Array
                nw100_ = _dafny.Array(False, 24)
                d_694_v66_ = nw100_
                d_695_v67_: _dafny.Set
                d_695_v67_ = _dafny.Set({p1})
                index98_ = default__.safeIndex(1, (d_694_v66_).length(0))
                (d_694_v66_)[index98_] = not((d_685_v61_).fm32((d_685_v61_).f39, len(d_695_v67_), globalState))
                d_696_v68_: _dafny.Seq
                d_696_v68_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sn")), (d_686_v62_)[default__.safeIndex(249, (d_686_v62_).length(0))]])
                index99_ = default__.safeIndex(1, (d_694_v66_).length(0))
                rhs114_ = (d_685_v61_).f39
                rhs115_ = default__.safeModuloInt(p1, p1)
                rhs116_ = (d_696_v68_)[default__.safeIndex(len(d_577_v0_), len(d_696_v68_))]
                rhs117_ = (p1) <= (p1)
                rhs118_ = p1
                lhs83_ = globalState
                lhs84_ = d_694_v66_
                lhs85_ = default__.safeIndex(1, (d_694_v66_).length(0))
                lhs86_ = globalState
                lhs83_.f13 = rhs114_
                r2 = rhs115_
                r1 = rhs116_
                lhs84_[lhs85_] = rhs117_
                lhs86_.f7 = rhs118_
                d_697_v69_: _dafny.Array
                def lambda29_(d_698_p1_, d_699_v61_):
                    def lambda30_(d_700_i4_):
                        return (_dafny.Map({(self).f24: len(_dafny.Map({-632: d_698_p1_}))})) | (_dafny.Map({(d_699_v61_).f39: d_698_p1_}))

                    return lambda30_

                init16_ = lambda29_(p1, d_685_v61_)
                nw101_ = _dafny.Array(None, 2)
                for i0_16_ in range(nw101_.length(0)):
                    nw101_[i0_16_] = init16_(i0_16_)
                d_697_v69_ = nw101_
                nw102_ = _dafny.Array(_dafny.Map({}), 7)
                d_697_v69_ = nw102_
                d_701_v70_: _dafny.Set
                d_701_v70_ = _dafny.Set({(self).f24})
                index100_ = default__.safeIndex(1, (d_694_v66_).length(0))
                (d_694_v66_)[index100_] = ((self).f24 if (d_694_v66_)[default__.safeIndex(1, (d_694_v66_).length(0))] else (d_701_v70_).issubset(d_701_v70_))
            elif True:
                (globalState).f13 = ((d_579_v1_)[default__.safeIndex(p1, len(d_579_v1_))] if ((d_686_v62_)[default__.safeIndex(249, (d_686_v62_).length(0))]) == ((d_686_v62_)[default__.safeIndex(249, (d_686_v62_).length(0))]) else not((self).f24))
                (globalState).f13 = (d_685_v61_).f39
                index101_ = default__.safeIndex(249, (d_686_v62_).length(0))
                (d_686_v62_)[index101_] = (d_686_v62_)[default__.safeIndex(249, (d_686_v62_).length(0))]
                (globalState).f13 = (d_685_v61_).f39
                d_702_v71_: D2
                d_702_v71_ = D2_DC7(True, (d_685_v61_).f39)
                d_703_v72_: _dafny.Array
                nw103_ = _dafny.Array(None, 23)
                nw103_[int(0)] = d_702_v71_
                nw103_[int(1)] = d_702_v71_
                nw103_[int(2)] = d_702_v71_
                nw103_[int(3)] = D2_DC7((self).fm4(p1, d_579_v1_, globalState), (d_685_v61_).f39)
                nw103_[int(4)] = d_702_v71_
                nw103_[int(5)] = d_702_v71_
                nw103_[int(6)] = d_702_v71_
                nw103_[int(7)] = d_702_v71_
                nw103_[int(8)] = d_702_v71_
                nw103_[int(9)] = D2_DC7((d_685_v61_).f39, (d_685_v61_).f39)
                nw103_[int(10)] = d_702_v71_
                nw103_[int(11)] = d_702_v71_
                nw103_[int(12)] = d_702_v71_
                nw103_[int(13)] = d_702_v71_
                def iife126_(_pat_let31_0):
                    def iife127_(d_704_dt__update__tmp_h5_):
                        def iife128_(_pat_let32_0):
                            def iife129_(d_705_dt__update_hcf13_h0_):
                                return D2_DC7(d_705_dt__update_hcf13_h0_, (d_704_dt__update__tmp_h5_).cf14)
                            return iife129_(_pat_let32_0)
                        return iife128_((self).f24)
                    return iife127_(_pat_let31_0)
                nw103_[int(14)] = iife126_(d_702_v71_)
                nw103_[int(15)] = d_702_v71_
                nw103_[int(16)] = d_702_v71_
                nw103_[int(17)] = default__.fm38(True, (d_685_v61_).f39, len(d_579_v1_), globalState)
                nw103_[int(18)] = d_702_v71_
                nw103_[int(19)] = d_702_v71_
                nw103_[int(20)] = d_702_v71_
                nw103_[int(21)] = d_702_v71_
                nw103_[int(22)] = d_702_v71_
                d_703_v72_ = nw103_
                d_706_v73_: _dafny.Seq
                d_706_v73_ = _dafny.SeqWithoutIsStrInference([d_703_v72_, d_703_v72_, d_703_v72_, d_703_v72_, d_703_v72_])
                d_707_v74_: _dafny.Map
                d_707_v74_ = _dafny.Map({p1: (d_685_v61_).f39})
                d_708_v75_: _dafny.Seq
                d_708_v75_ = _dafny.SeqWithoutIsStrInference([p1])
                d_709_v76_: _dafny.Seq
                d_709_v76_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_708_v75_)[default__.safeIndex(p1, len(d_708_v75_))])])
                rhs119_ = d_706_v73_
                rhs120_ = (d_685_v61_).f39
                rhs121_ = (self).f24
                rhs122_ = (0) - ((p1) * (default__.safeModuloInt(len(d_707_v74_), (d_709_v76_)[default__.safeIndex(p1, len(d_709_v76_))])))
                rhs123_ = True
                lhs87_ = globalState
                lhs88_ = globalState
                lhs89_ = globalState
                d_706_v73_ = rhs119_
                lhs87_.f13 = rhs120_
                lhs88_.f13 = rhs121_
                r2 = rhs122_
                lhs89_.f13 = rhs123_
        d_710_v77_: _dafny.Array
        def lambda31_(d_711_i6_):
            return self.f23

        init17_ = lambda31_
        nw104_ = _dafny.Array(None, 9)
        for i0_17_ in range(nw104_.length(0)):
            nw104_[i0_17_] = init17_(i0_17_)
        d_710_v77_ = nw104_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_710_v77_).length(0)):
            d_712_i5_: int = guard_loop_1_
            if (True) and (((0) <= (d_712_i5_)) and ((d_712_i5_) < ((d_710_v77_).length(0)))):
                (d_710_v77_)[(d_712_i5_)] = self.f23
        d_713_v78_: T3
        nw105_ = C1()
        nw105_.ctor__((self).f24, (self).f24)
        d_713_v78_ = nw105_
        d_714_v79_: _dafny.Array
        def lambda32_(d_715_i7_):
            return (self).f24

        init18_ = lambda32_
        nw106_ = _dafny.Array(None, 13)
        for i0_18_ in range(nw106_.length(0)):
            nw106_[i0_18_] = init18_(i0_18_)
        d_714_v79_ = nw106_
        d_716_v80_: _dafny.Map
        d_716_v80_ = _dafny.Map({d_713_v78_: (d_714_v79_ if (d_713_v78_).f27 else d_714_v79_)})
        d_716_v80_ = ((d_716_v80_) | (d_716_v80_)) | ((d_716_v80_).set(d_713_v78_, d_714_v79_))
        if (self.f23) in (d_577_v0_):
            d_717_v81_: D2
            d_717_v81_ = D2_DC7((d_713_v78_).f27, (d_713_v78_).f27)
            (globalState).f13 = (d_717_v81_).cf14
            d_718_v82_: _dafny.Map
            d_718_v82_ = _dafny.Map({(d_713_v78_).f27: d_577_v0_})
            d_719_v83_: C0
            nw107_ = C0()
            nw107_.ctor__(d_718_v82_, not((self).f24), (d_577_v0_)[default__.safeIndex(-439, len(d_577_v0_))], (p1) >= (len(d_577_v0_)))
            d_719_v83_ = nw107_
            d_720_v84_: _dafny.Map
            d_720_v84_ = _dafny.Map({p1: d_579_v1_})
            d_721_v85_: D6
            d_721_v85_ = D6_DC15(d_720_v84_)
            pat_let_tv24_ = d_579_v1_
            d_722_v87_: _dafny.Array
            nw108_ = _dafny.Array(None, 8)
            nw108_[int(0)] = d_721_v85_
            nw108_[int(1)] = default__.fm39(p1, globalState)
            nw108_[int(2)] = (d_721_v85_ if (d_719_v83_).f38 else d_721_v85_)
            def iife130_(_pat_let33_0):
                def iife131_(d_723_dt__update__tmp_h6_):
                    def iife133_():
                        coll64_ = _dafny.Map()
                        compr_64_: int
                        for compr_64_ in _dafny.IntegerRange(36, -788):
                            d_724_v86_: int = compr_64_
                            if ((36) <= (d_724_v86_)) and ((d_724_v86_) < (-788)):
                                coll64_[(d_724_v86_) * (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exgcvpsfo")))])))] = pat_let_tv24_
                        return _dafny.Map(coll64_)
                    def iife132_(_pat_let34_0):
                        def iife134_(d_725_dt__update_hcf24_h0_):
                            return D6_DC15(d_725_dt__update_hcf24_h0_)
                        return iife134_(_pat_let34_0)
                    return iife132_(iife133_()
                    )
                return iife131_(_pat_let33_0)
            nw108_[int(3)] = iife130_(default__.fm39(278, globalState))
            nw108_[int(4)] = d_721_v85_
            nw108_[int(5)] = d_721_v85_
            nw108_[int(6)] = D6_DC15(d_720_v84_)
            nw108_[int(7)] = d_721_v85_
            d_722_v87_ = nw108_
            index102_ = default__.safeIndex(810, (d_722_v87_).length(0))
            (d_722_v87_)[index102_] = d_721_v85_
            index103_ = default__.safeIndex(810, (d_722_v87_).length(0))
            (d_722_v87_)[index103_] = d_721_v85_
            d_726_v88_: D2
            d_726_v88_ = D2_DC6((d_718_v82_) | (d_718_v82_))
            d_726_v88_ = d_726_v88_
            d_727_v89_: _dafny.Seq
            d_727_v89_ = _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([556 for d_728_i8_ in range(default__.abs(-678))])))])
            (globalState).f13 = ((_dafny.SeqWithoutIsStrInference([54, p1, p1])) + (_dafny.SeqWithoutIsStrInference([p1]))) < ((d_727_v89_) + (d_727_v89_))
        elif True:
            d_729_v90_: _dafny.Map
            d_729_v90_ = _dafny.Map({(d_713_v78_).f27: (d_713_v78_).f27})
            d_730_v91_: _dafny.Seq
            d_730_v91_ = _dafny.SeqWithoutIsStrInference([d_729_v90_, d_729_v90_, d_729_v90_, d_729_v90_])
            d_731_v92_: _dafny.Map
            d_731_v92_ = _dafny.Map({d_580_v2_: d_730_v91_})
            r0 = (0) - (len(((d_731_v92_)[d_580_v2_] if (d_580_v2_) in (d_731_v92_) else (d_730_v91_) + (_dafny.SeqWithoutIsStrInference([d_729_v90_])))))
            d_732_v93_: D0
            d_732_v93_ = D0_DC2((self).fm4(p1, d_579_v1_, globalState), p1, (d_713_v78_).f27)
            d_733_v94_: _dafny.Map
            d_733_v94_ = _dafny.Map({d_732_v93_: (self).f24})
            if ((d_733_v94_)[d_732_v93_] if (d_732_v93_) in (d_733_v94_) else (d_713_v78_).f27):
                (self).f23 = (default__.fm29(p1, (d_713_v78_).f27, (d_713_v78_).f27, p1, globalState) if (self).f24 else self.f23)
                (globalState).f7 = (len(d_577_v0_)) + ((0) - (p1))
                d_734_v95_: _dafny.Map
                d_734_v95_ = _dafny.Map({(default__.fm40((d_713_v78_).f27, globalState)).cf20: not((d_713_v78_).f27)})
                d_735_v96_: _dafny.Seq
                d_735_v96_ = _dafny.SeqWithoutIsStrInference([p1])
                d_734_v95_ = (d_734_v95_).set(((d_713_v78_).fm9(d_735_v96_, p1, (d_713_v78_).f27, globalState)) + (p1), (d_713_v78_).f27)
                (globalState).f7 = default__.safeDivisionInt((0) - ((p1) * (len(d_735_v96_))), p1)
                d_736_v97_: _dafny.Set
                d_736_v97_ = _dafny.Set({(d_713_v78_).f27, (d_713_v78_).f27})
                d_734_v95_ = (d_734_v95_).set(len(d_736_v97_), (d_735_v96_) == (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbav"))), p1, len(d_736_v97_)])))
            elif True:
                d_737_v98_: _dafny.Seq
                d_737_v98_ = _dafny.SeqWithoutIsStrInference([d_577_v0_, d_577_v0_, ((d_577_v0_) + (d_577_v0_)).set(default__.safeIndex(p1, len((d_577_v0_) + (d_577_v0_))), self.f23), (d_577_v0_) + (d_577_v0_)])
                r1 = (d_737_v98_)[default__.safeIndex(p1, len(d_737_v98_))]
                (globalState).f13 = (d_713_v78_).f27
                (globalState).f7 = p1
                (globalState).f7 = p1
                d_738_v99_: _dafny.Set
                d_738_v99_ = _dafny.Set({_dafny.CodePoint('r'), self.f23, self.f23})
                def iife135_(_pat_let35_0):
                    def iife136_(d_739_dt__update__tmp_h7_):
                        def iife137_(_pat_let36_0):
                            def iife138_(d_740_dt__update_hcf6_h0_):
                                return D1_DC4(d_740_dt__update_hcf6_h0_)
                            return iife138_(_pat_let36_0)
                        return iife137_(self.f23)
                    return iife136_(_pat_let35_0)
                d_738_v99_ = _dafny.Set({(iife135_(d_580_v2_)).cf6, self.f23})
            if (d_713_v78_).f27:
                rhs124_ = self.f23
                rhs125_ = (d_713_v78_).f27
                rhs126_ = (self).f24
                lhs90_ = self
                lhs91_ = globalState
                lhs92_ = globalState
                lhs90_.f23 = rhs124_
                lhs91_.f13 = rhs125_
                lhs92_.f13 = rhs126_
                (globalState).f13 = (d_713_v78_).f27
                d_741_v100_: D2
                d_741_v100_ = D2_DC7((d_713_v78_).f27, (d_713_v78_).f27)
                d_742_v101_: D5
                d_742_v101_ = D5_DC13(377, (d_579_v1_) != (d_579_v1_), not ((d_741_v100_).cf13) or (default__.fm2(globalState)))
                d_742_v101_ = d_742_v101_
                nw109_ = _dafny.Array(False, 16)
                d_714_v79_ = nw109_
                nw110_ = _dafny.Array(False, 18)
                d_714_v79_ = nw110_
            elif True:
                d_743_v102_: _dafny.Map
                d_743_v102_ = _dafny.Map({True: d_577_v0_})
                d_744_v103_: _dafny.MultiSet
                d_744_v103_ = _dafny.MultiSet([975])
                d_745_v104_: C0
                nw111_ = C0()
                nw111_.ctor__(d_743_v102_, (self).fm4((d_744_v103_).cardinality, d_579_v1_, globalState), self.f23, (d_713_v78_).f27)
                d_745_v104_ = nw111_
                d_746_v105_: _dafny.Array
                nw112_ = _dafny.Array(None, 26)
                d_746_v105_ = nw112_
                d_747_v106_: C1
                nw113_ = C1()
                nw113_.ctor__((d_713_v78_).f27, (d_713_v78_).f27)
                d_747_v106_ = nw113_
                index104_ = default__.safeIndex(536, (d_746_v105_).length(0))
                (d_746_v105_)[index104_] = d_747_v106_
                d_748_v107_: C1
                d_748_v107_ = d_747_v106_
                index105_ = default__.safeIndex(536, (d_746_v105_).length(0))
                (d_746_v105_)[index105_] = (d_748_v107_)
                d_749_v108_: bool
                out30_: bool
                out30_ = (d_747_v106_).m4((self).f24, globalState)
                d_749_v108_ = out30_
                d_749_v108_ = not((self).f24)
                (globalState).f13 = ((self).f24) or (((self).f24) or ((d_713_v78_).f27))
            r0 = (len(d_577_v0_)) * ((self).fm24(p1, (d_713_v78_).f27, globalState))
            d_750_v109_: _dafny.Seq
            d_750_v109_ = _dafny.SeqWithoutIsStrInference([p1, p1, default__.fm0((d_713_v78_).f27, 395, globalState)])
            d_751_v110_: _dafny.Map
            d_751_v110_ = _dafny.Map({self.f23: (self).f24})
            (globalState).f7 = (571) - ((0) - ((d_750_v109_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_751_v110_])), len(d_750_v109_))]))
        r0 = (((d_713_v78_).fm9(_dafny.SeqWithoutIsStrInference([p1, p1]), p1, (self).f24, globalState) if (d_713_v78_).f27 else p1) if (self).f24 else (0) - ((p1) - (p1)))
        r1 = d_577_v0_
        r2 = default__.safeModuloInt(406, (p1 if (self).f24 else p1))
        return r0, r1, r2

    @property
    def f36(self):
        return self._f36

class C3(T2, T0):
    def  __init__(self):
        self._f23: str = _dafny.CodePoint('D')
        self._f24: bool = False
        self._f25: _dafny.Array = _dafny.Array(None, 0)
        self._f35: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f23(self):
        return self._f23
    @f23.setter
    def f23(self, value):
        self._f23 = value
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f35, f25, f23, f24):
        (self)._f35 = f35
        (self)._f25 = f25
        (self).f23 = f23
        (self)._f24 = f24

    def fm7(self, globalState):
        if ((self).f24) or (False):
            return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "er"))) for d_752_i0_ in range(default__.abs(926))])
        elif True:
            return _dafny.SeqWithoutIsStrInference([278])

    def fm8(self, p0, p1, globalState):
        return (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))})) | (_dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))}))

    def fm4(self, p0, p1, globalState):
        return ((538) - (len(_dafny.SeqWithoutIsStrInference([self.f23 for d_753_i0_ in range(default__.abs(947))])))) <= (len((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkrmexiw"))}) if (self).f24 else _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aaucpfakq")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrrodm")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgsjiaqub"))}))))

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        d_754_v0_: int
        d_754_v0_ = 331
        (globalState).f7 = (d_754_v0_) - (d_754_v0_)
        d_755_v1_: _dafny.Seq
        d_755_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbplvd"))
        d_756_v2_: D4
        d_756_v2_ = D4_DC10(d_755_v1_)
        d_755_v1_ = (d_755_v1_) + ((d_756_v2_).cf17)
        d_757_v3_: _dafny.MultiSet
        d_757_v3_ = _dafny.MultiSet([d_754_v0_, d_754_v0_, d_754_v0_, d_754_v0_, d_754_v0_])
        d_758_i0_: int
        d_758_i0_ = 0
        with _dafny.label("2"):
            while not((d_757_v3_).isdisjoint((d_757_v3_ if (self).f24 else d_757_v3_))):
                with _dafny.c_label("2"):
                    if (d_758_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_758_i0_ = (d_758_i0_) + (1)
                    d_759_v4_: bool
                    d_760_v5_: str
                    d_761_v6_: bool
                    out31_: bool
                    out32_: str
                    out33_: bool
                    out31_, out32_, out33_ = (self).m11(p0, globalState)
                    d_759_v4_ = out31_
                    d_760_v5_ = out32_
                    d_761_v6_ = out33_
                    d_762_v7_: _dafny.Map
                    d_762_v7_ = _dafny.Map({p0: d_755_v1_})
                    d_763_v8_: C0
                    nw114_ = C0()
                    nw114_.ctor__(d_762_v7_, d_761_v6_, p1, True)
                    d_763_v8_ = nw114_
                    r1 = d_754_v0_
                    (globalState).f13 = (self).f24
                    pass
            pass
        d_764_v9_: _dafny.Map
        d_764_v9_ = _dafny.Map({d_755_v1_: d_754_v0_})
        (globalState).f7 = len((_dafny.Map({d_755_v1_: d_754_v0_})) | (d_764_v9_))
        r0 = p0
        (globalState).f13 = (self).f24
        r0 = p0
        r1 = d_754_v0_
        d_765_v10_: _dafny.MultiSet
        d_765_v10_ = _dafny.MultiSet([(self).f24, p0])
        d_766_v11_: _dafny.Map
        d_766_v11_ = _dafny.Map({d_755_v1_: d_765_v10_})
        r2 = (p0) not in (((d_766_v11_)[d_755_v1_] if (d_755_v1_) in (d_766_v11_) else d_765_v10_))
        return r0, r1, r2

    def m11(self, p0, globalState):
        r0: bool = False
        r1: str = _dafny.CodePoint('D')
        r2: bool = False
        d_767_v0_: D5
        d_767_v0_ = D5_DC14(p0)
        r0 = (d_767_v0_).cf23
        d_768_i0_: int
        d_768_i0_ = 0
        with _dafny.label("3"):
            while (self).f24:
                with _dafny.c_label("3"):
                    if (d_768_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_768_i0_ = (d_768_i0_) + (1)
                    d_769_v1_: int
                    d_769_v1_ = 49
                    (globalState).f7 = d_769_v1_
                    source13_ = d_767_v0_
                    if source13_.is_DC13:
                        d_770___mcc_h0_ = source13_.cf20
                        d_771___mcc_h1_ = source13_.cf21
                        d_772___mcc_h2_ = source13_.cf22
                        d_773_cf22_ = d_772___mcc_h2_
                        d_774_cf21_ = d_771___mcc_h1_
                        d_775_cf20_ = d_770___mcc_h0_
                        d_776_v2_: _dafny.Seq
                        d_776_v2_ = _dafny.SeqWithoutIsStrInference([d_769_v1_, (0) - (d_769_v1_)])
                        d_777_v3_: _dafny.Seq
                        d_777_v3_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_769_v1_, -258, d_769_v1_, d_769_v1_, d_775_cf20_])])
                        (globalState).f7 = len((d_776_v2_) + ((d_777_v3_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([self.f23 for d_778_i1_ in range(default__.abs(690))])), len(d_777_v3_))]))
                        (self).f23 = self.f23
                        d_779_v4_: C1
                        nw115_ = C1()
                        nw115_.ctor__(d_773_cf22_, p0)
                        d_779_v4_ = nw115_
                        (globalState).f7 = d_769_v1_
                    elif source13_.is_DC14:
                        d_780___mcc_h3_ = source13_.cf23
                        d_781_cf23_ = d_780___mcc_h3_
                        d_782_v5_: _dafny.Array
                        nw116_ = _dafny.Array(int(0), 12)
                        d_782_v5_ = nw116_
                        (globalState).f22 = d_782_v5_
                        d_783_v6_: _dafny.Set
                        d_783_v6_ = _dafny.Set({p0})
                        d_783_v6_ = default__.fm41(globalState)
                        d_784_v7_: _dafny.Seq
                        d_784_v7_ = _dafny.SeqWithoutIsStrInference([(self).f24, p0, False, not(d_781_cf23_)])
                        d_785_v8_: _dafny.Map
                        d_785_v8_ = _dafny.Map({d_784_v7_: d_769_v1_})
                        d_786_v9_: _dafny.Map
                        d_786_v9_ = _dafny.Map({p0: default__.fm36(globalState)})
                        d_787_v10_: _dafny.Array
                        def lambda33_(d_788_i2_):
                            return False

                        init19_ = lambda33_
                        nw117_ = _dafny.Array(None, 4)
                        for i0_19_ in range(nw117_.length(0)):
                            nw117_[i0_19_] = init19_(i0_19_)
                        d_787_v10_ = nw117_
                        d_789_v11_: _dafny.Map
                        d_789_v11_ = _dafny.Map({d_787_v10_: d_769_v1_})
                        d_785_v8_ = (d_785_v8_).set((d_784_v7_) + (((d_786_v9_)[p0] if (p0) in (d_786_v9_) else d_784_v7_)), ((d_789_v11_)[d_787_v10_] if (d_787_v10_) in (d_789_v11_) else d_769_v1_))
                        d_781_cf23_ = p0
                    elif True:
                        d_790___mcc_h4_ = source13_.cf19
                        d_791_cf19_ = d_790___mcc_h4_
                        d_792_v12_: _dafny.Set
                        d_792_v12_ = _dafny.Set({self.f23, self.f23, self.f23})
                        d_793_v13_: _dafny.Seq
                        d_793_v13_ = _dafny.SeqWithoutIsStrInference([d_769_v1_])
                        d_794_v14_: _dafny.Seq
                        d_794_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cemw"))
                        d_795_v15_: _dafny.Map
                        d_795_v15_ = _dafny.Map({d_769_v1_: (0) - ((d_793_v13_)[default__.safeIndex(default__.fm0(p0, len(d_794_v14_), globalState), len(d_793_v13_))])})
                        d_796_v16_: _dafny.Map
                        d_796_v16_ = _dafny.Map({(d_792_v12_).ispropersubset(d_792_v12_): len(d_795_v15_)})
                        d_796_v16_ = (d_796_v16_).set(p0, d_769_v1_)
                        d_797_v18_: _dafny.Map
                        def iife139_():
                            coll65_ = _dafny.Map()
                            compr_65_: int
                            for compr_65_ in (d_793_v13_).Elements:
                                d_798_v17_: int = compr_65_
                                if (d_798_v17_) in (d_793_v13_):
                                    coll65_[default__.safeDivisionInt(d_798_v17_, d_769_v1_)] = d_769_v1_
                            return _dafny.Map(coll65_)
                        d_797_v18_ = (d_795_v15_) | (iife139_()
                        )
                        d_797_v18_ = d_795_v15_
                        d_799_v19_: _dafny.Seq
                        d_799_v19_ = _dafny.SeqWithoutIsStrInference([p0])
                        d_800_v20_: D0
                        d_800_v20_ = D0_DC0(_dafny.MultiSet(d_799_v19_))
                        d_801_v21_: _dafny.Map
                        d_801_v21_ = _dafny.Map({p0: (self).f24})
                        d_802_v22_: _dafny.Array
                        nw118_ = _dafny.Array(None, 26)
                        nw118_[int(0)] = p0
                        nw118_[int(1)] = (d_769_v1_) < (default__.fm0(not((self).f24), d_769_v1_, globalState))
                        nw118_[int(2)] = (self).f24
                        nw118_[int(3)] = (d_799_v19_) < (d_799_v19_)
                        nw118_[int(4)] = (self).f24
                        nw118_[int(5)] = p0
                        nw118_[int(6)] = (self).f24
                        nw118_[int(7)] = not(False)
                        nw118_[int(8)] = (p0) or (p0)
                        nw118_[int(9)] = p0
                        nw118_[int(10)] = (self).f24
                        nw118_[int(11)] = p0
                        nw118_[int(12)] = ((d_800_v20_).cf0).isdisjoint(_dafny.MultiSet([p0, (self).f24]))
                        nw118_[int(13)] = (len(d_801_v21_)) <= (len(d_801_v21_))
                        nw118_[int(14)] = p0
                        nw118_[int(15)] = default__.fm2(globalState)
                        nw118_[int(16)] = p0
                        nw118_[int(17)] = p0
                        nw118_[int(18)] = (d_799_v19_)[default__.safeIndex(d_769_v1_, len(d_799_v19_))]
                        nw118_[int(19)] = p0
                        nw118_[int(20)] = (self).f24
                        nw118_[int(21)] = (self).f24
                        nw118_[int(22)] = (len(d_794_v14_)) == (d_769_v1_)
                        nw118_[int(23)] = p0
                        nw118_[int(24)] = (self).f24
                        nw118_[int(25)] = (self).f24
                        d_802_v22_ = nw118_
                        d_803_v23_: _dafny.Map
                        d_803_v23_ = _dafny.Map({-580: (self).f24})
                        index106_ = default__.safeIndex(140, (d_802_v22_).length(0))
                        (d_802_v22_)[index106_] = ((d_803_v23_)[len(d_793_v13_)] if (len(d_793_v13_)) in (d_803_v23_) else not((self).f24))
                        index107_ = default__.safeIndex(140, (d_802_v22_).length(0))
                        (d_802_v22_)[index107_] = (self).f24
                        d_804_v24_: _dafny.Array
                        nw119_ = _dafny.Array(int(0), 23)
                        d_804_v24_ = nw119_
                        index108_ = default__.safeIndex(703, (d_804_v24_).length(0))
                        (d_804_v24_)[index108_] = (0) - (default__.safeModuloInt(d_769_v1_, d_769_v1_))
                        index109_ = default__.safeIndex(703, (d_804_v24_).length(0))
                        (d_804_v24_)[index109_] = d_769_v1_
                    d_805_v25_: _dafny.Array
                    def lambda34_(d_806_v1_):
                        def lambda35_(d_807_i3_):
                            return D4_DC11(d_806_v1_)

                        return lambda35_

                    init20_ = lambda34_(d_769_v1_)
                    nw120_ = _dafny.Array(None, 1)
                    for i0_20_ in range(nw120_.length(0)):
                        nw120_[i0_20_] = init20_(i0_20_)
                    d_805_v25_ = nw120_
                    d_808_v26_: D4
                    d_808_v26_ = D4_DC11(911)
                    index110_ = default__.safeIndex(434, (d_805_v25_).length(0))
                    (d_805_v25_)[index110_] = d_808_v26_
                    index111_ = default__.safeIndex(434, (d_805_v25_).length(0))
                    (d_805_v25_)[index111_] = d_808_v26_
                    d_809_v27_: _dafny.Array
                    nw121_ = _dafny.Array(int(0), 19)
                    d_809_v27_ = nw121_
                    index112_ = default__.safeIndex(426, (d_809_v27_).length(0))
                    (d_809_v27_)[index112_] = (0) - (d_769_v1_)
                    index113_ = default__.safeIndex(426, (d_809_v27_).length(0))
                    (d_809_v27_)[index113_] = (d_769_v1_) * (d_769_v1_)
                    pass
            pass
        (globalState).f13 = default__.fm2(globalState)
        d_810_v28_: _dafny.Seq
        d_810_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plwsf"))
        d_811_v29_: int
        d_811_v29_ = 669
        d_812_v30_: _dafny.Array
        nw122_ = _dafny.Array(None, 24)
        nw122_[int(0)] = d_810_v28_
        nw122_[int(1)] = d_810_v28_
        nw122_[int(2)] = d_810_v28_
        nw122_[int(3)] = d_810_v28_
        nw122_[int(4)] = d_810_v28_
        nw122_[int(5)] = d_810_v28_
        nw122_[int(6)] = d_810_v28_
        nw122_[int(7)] = d_810_v28_
        nw122_[int(8)] = d_810_v28_
        nw122_[int(9)] = _dafny.SeqWithoutIsStrInference([self.f23 for d_813_i4_ in range(default__.abs(774))])
        nw122_[int(10)] = d_810_v28_
        nw122_[int(11)] = d_810_v28_
        nw122_[int(12)] = d_810_v28_
        nw122_[int(13)] = (d_810_v28_) + ((_dafny.SeqWithoutIsStrInference([self.f23 for d_814_i5_ in range(default__.abs(-90))])).set(default__.safeIndex(d_811_v29_, len(_dafny.SeqWithoutIsStrInference([self.f23 for d_815_i5_ in range(default__.abs(-90))]))), self.f23))
        nw122_[int(14)] = d_810_v28_
        nw122_[int(15)] = (d_810_v28_) + (d_810_v28_)
        nw122_[int(16)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_816_i6_ in range(default__.abs(439))])
        nw122_[int(17)] = _dafny.SeqWithoutIsStrInference([self.f23 for d_817_i7_ in range(default__.abs(769))])
        nw122_[int(18)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lti")) if (self).f24 else d_810_v28_)
        nw122_[int(19)] = d_810_v28_
        nw122_[int(20)] = d_810_v28_
        nw122_[int(21)] = d_810_v28_
        nw122_[int(22)] = d_810_v28_
        nw122_[int(23)] = d_810_v28_
        d_812_v30_ = nw122_
        d_818_v31_: _dafny.Seq
        d_818_v31_ = _dafny.SeqWithoutIsStrInference([p0])
        d_819_v32_: _dafny.Map
        d_819_v32_ = _dafny.Map({len(d_818_v31_): p0})
        d_820_v33_: _dafny.Set
        d_820_v33_ = _dafny.Set({((d_819_v32_)[default__.fm0((self).f24, d_811_v29_, globalState)] if (default__.fm0((self).f24, d_811_v29_, globalState)) in (d_819_v32_) else (self).f24), True, (self).f24})
        pat_let_tv25_ = d_811_v29_
        pat_let_tv26_ = d_811_v29_
        def lambda36_(source14_):
            if source14_.is_DC13:
                d_821___mcc_h5_ = source14_.cf20
                d_822___mcc_h6_ = source14_.cf21
                d_823___mcc_h7_ = source14_.cf22
                d_824_cf22_ = d_823___mcc_h7_
                d_825_cf21_ = d_822___mcc_h6_
                d_826_cf20_ = d_821___mcc_h5_
                return d_826_cf20_
            elif source14_.is_DC14:
                d_827___mcc_h8_ = source14_.cf23
                d_828_cf23_ = d_827___mcc_h8_
                return pat_let_tv25_
            elif True:
                d_829___mcc_h9_ = source14_.cf19
                d_830_cf19_ = d_829___mcc_h9_
                return pat_let_tv26_

        rhs127_ = d_812_v30_
        rhs128_ = not(p0)
        rhs129_ = lambda36_(D5_DC13((0) - (len(d_820_v33_)), p0, False))
        rhs130_ = p0
        rhs131_ = (self).f24
        lhs93_ = globalState
        lhs94_ = globalState
        lhs95_ = globalState
        d_812_v30_ = rhs127_
        lhs93_.f13 = rhs128_
        lhs94_.f7 = rhs129_
        lhs95_.f13 = rhs130_
        r0 = rhs131_
        d_831_v34_: _dafny.Array
        def lambda37_(d_832_p0_):
            def lambda38_(d_833_i9_):
                return d_832_p0_

            return lambda38_

        init21_ = lambda37_(p0)
        nw123_ = _dafny.Array(None, 28)
        for i0_21_ in range(nw123_.length(0)):
            nw123_[i0_21_] = init21_(i0_21_)
        d_831_v34_ = nw123_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_831_v34_).length(0)):
            d_834_i8_: int = guard_loop_2_
            if (True) and (((0) <= (d_834_i8_)) and ((d_834_i8_) < ((d_831_v34_).length(0)))):
                (d_831_v34_)[(d_834_i8_)] = (_dafny.MultiSet([(self).f24, p0])).issubset(_dafny.MultiSet([False, (self).f24]))
        (globalState).f13 = (self).f24
        r0 = p0
        r1 = default__.fm34((d_811_v29_ if True else d_811_v29_), not(False), globalState)
        r2 = (self).f24
        return r0, r1, r2

    @property
    def f35(self):
        return self._f35

class C4(T2, T0):
    def  __init__(self):
        self._f23: str = _dafny.CodePoint('D')
        self._f24: bool = False
        self._f25: _dafny.Array = _dafny.Array(None, 0)
        self.f34: _dafny.Map = _dafny.Map({})
        self._f33: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f23(self):
        return self._f23
    @f23.setter
    def f23(self, value):
        self._f23 = value
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f33, f34, f25, f23, f24):
        (self)._f33 = f33
        (self).f34 = f34
        (self)._f25 = f25
        (self).f23 = f23
        (self)._f24 = f24

    def fm7(self, globalState):
        def iife140_():
            coll66_ = _dafny.Map()
            compr_66_: _dafny.Map
            for compr_66_ in (_dafny.MultiSet([_dafny.Map({(self).f24: False}), _dafny.Map({(self).f24: (self).f24})])).Elements:
                d_835_v0_: _dafny.Map = compr_66_
                if (d_835_v0_) in (_dafny.MultiSet([_dafny.Map({(self).f24: False}), _dafny.Map({(self).f24: (self).f24})])):
                    coll66_[d_835_v0_] = (self).f33
            return _dafny.Map(coll66_)
        return _dafny.SeqWithoutIsStrInference([207, (len(_dafny.Map({(self).f24: -927}))) * (len(iife140_()
        ))])

    def fm8(self, p0, p1, globalState):
        return ((D2_DC6(_dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))})) if (self).f24 else D2_DC6(_dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))})))).cf12

    def fm4(self, p0, p1, globalState):
        return (self).f24

    def fm22(self, p0, p1, globalState):
        return (self).f24

    def fm23(self, p0, p1, p2, globalState):
        return (self).f24

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        if p0:
            d_836_v0_: _dafny.Seq
            d_836_v0_ = _dafny.SeqWithoutIsStrInference([not(p0), p0])
            d_837_v1_: _dafny.MultiSet
            d_837_v1_ = _dafny.MultiSet([not((self).f24), p0, (d_836_v0_)[default__.safeIndex((self).f33, len(d_836_v0_))]])
            d_838_v2_: C2
            nw124_ = C2()
            nw124_.ctor__((d_837_v1_).intersection(_dafny.MultiSet([(self).f24, (self).f24])), (self).f25, _dafny.CodePoint('c'), default__.fm2(globalState))
            d_838_v2_ = nw124_
            d_839_v3_: _dafny.Set
            d_839_v3_ = _dafny.Set({315})
            (globalState).f13 = (d_839_v3_).issubset(d_839_v3_)
            d_840_v4_: C1
            nw125_ = C1()
            nw125_.ctor__(True, (self).f24)
            d_840_v4_ = nw125_
            d_841_v5_: _dafny.Seq
            d_841_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
            d_842_v6_: _dafny.Map
            d_842_v6_ = _dafny.Map({(d_840_v4_).f39: d_841_v5_})
            source15_ = D2_DC6(d_842_v6_)
            if source15_.is_DC7:
                d_843___mcc_h0_ = source15_.cf13
                d_844___mcc_h1_ = source15_.cf14
                d_845_cf14_ = d_844___mcc_h1_
                d_846_cf13_ = d_843___mcc_h0_
                d_847_v7_: _dafny.Seq
                d_847_v7_ = _dafny.SeqWithoutIsStrInference([(self).f33])
                d_848_v8_: _dafny.Seq
                d_848_v8_ = _dafny.SeqWithoutIsStrInference([d_847_v7_])
                d_849_v9_: _dafny.Map
                d_849_v9_ = _dafny.Map({43: len(d_848_v8_)})
                d_850_v11_: _dafny.Map
                d_850_v11_ = _dafny.Map({(d_840_v4_).f39: p0})
                d_851_v12_: _dafny.Set
                d_851_v12_ = _dafny.Set({D6_DC16(d_850_v11_, (d_840_v4_).f39, (d_840_v4_).f39)})
                def iife141_():
                    coll67_ = _dafny.Map()
                    compr_67_: int
                    for compr_67_ in _dafny.IntegerRange(452, 718):
                        d_852_v10_: int = compr_67_
                        if ((452) <= (d_852_v10_)) and ((d_852_v10_) < (718)):
                            coll67_[default__.safeDivisionInt(d_852_v10_, (self).f33)] = (len(_dafny.Set({d_846_cf13_, (d_840_v4_).f39}))) + (len(d_839_v3_))
                    return _dafny.Map(coll67_)
                rhs132_ = d_845_cf14_
                rhs133_ = _dafny.SeqWithoutIsStrInference([d_845_cf14_, (d_840_v4_).f39])
                rhs134_ = iife141_()

                rhs135_ = ((self).f33) - ((len(_dafny.Map({False: d_851_v12_}))) + (len(_dafny.SeqWithoutIsStrInference([(self).f33 for d_853_i0_ in range(default__.abs(662))]))))
                lhs96_ = globalState
                d_846_cf13_ = rhs132_
                d_836_v0_ = rhs133_
                d_849_v9_ = rhs134_
                lhs96_.f7 = rhs135_
                d_854_v13_: D5
                d_854_v13_ = D5_DC13((self).f33, (d_840_v4_).f39, (d_840_v4_).f39)
                d_855_v14_: _dafny.MultiSet
                d_855_v14_ = _dafny.MultiSet([d_854_v13_, d_854_v13_])
                d_856_v15_: _dafny.MultiSet
                d_856_v15_ = _dafny.MultiSet([d_841_v5_])
                d_857_v16_: _dafny.Array
                def lambda39_(d_858_v0_):
                    def lambda40_(d_859_i1_):
                        return d_858_v0_

                    return lambda40_

                init22_ = lambda39_(d_836_v0_)
                nw126_ = _dafny.Array(None, 5)
                for i0_22_ in range(nw126_.length(0)):
                    nw126_[i0_22_] = init22_(i0_22_)
                d_857_v16_ = nw126_
                index114_ = default__.safeIndex(257, (d_857_v16_).length(0))
                (d_857_v16_)[index114_] = (d_838_v2_).fm25(d_846_cf13_, globalState)
                d_860_v17_: _dafny.MultiSet
                d_860_v17_ = _dafny.MultiSet([(self).f33])
                d_861_v18_: _dafny.Set
                d_861_v18_ = _dafny.Set({d_860_v17_, d_860_v17_, default__.fm42((self).f24, len(d_841_v5_), globalState)})
                index115_ = default__.safeIndex(257, (d_857_v16_).length(0))
                def iife142_():
                    coll68_ = _dafny.Set()
                    compr_68_: _dafny.MultiSet
                    for compr_68_ in (_dafny.Map({d_860_v17_: (self).f33})).keys.Elements:
                        d_862_v19_: _dafny.MultiSet = compr_68_
                        if (d_862_v19_) in (_dafny.Map({d_860_v17_: (self).f33})):
                            coll68_ = coll68_.union(_dafny.Set([d_862_v19_]))
                    return _dafny.Set(coll68_)
                rhs136_ = (d_855_v14_) | (d_855_v14_)
                rhs137_ = d_856_v15_
                rhs138_ = (d_840_v4_).fm32(d_846_cf13_, (self).f33, globalState)
                rhs139_ = not((d_861_v18_).issubset(iife142_()
                ))
                rhs140_ = d_836_v0_
                lhs97_ = globalState
                lhs98_ = globalState
                lhs99_ = d_857_v16_
                lhs100_ = default__.safeIndex(257, (d_857_v16_).length(0))
                d_855_v14_ = rhs136_
                d_856_v15_ = rhs137_
                lhs97_.f13 = rhs138_
                lhs98_.f13 = rhs139_
                lhs99_[lhs100_] = rhs140_
                d_863_v20_: _dafny.Array
                nw127_ = _dafny.Array(False, 20)
                d_863_v20_ = nw127_
                index116_ = default__.safeIndex(360, (d_863_v20_).length(0))
                (d_863_v20_)[index116_] = (d_840_v4_).f39
                index117_ = default__.safeIndex(360, (d_863_v20_).length(0))
                rhs141_ = default__.safeDivisionInt((self).f33, (0) - (((0) - ((self).f33)) - ((self).f33)))
                rhs142_ = (d_840_v4_).f39
                lhs101_ = globalState
                lhs102_ = d_863_v20_
                lhs103_ = default__.safeIndex(360, (d_863_v20_).length(0))
                lhs101_.f7 = rhs141_
                lhs102_[lhs103_] = rhs142_
                r0 = True
            elif source15_.is_DC6:
                d_864___mcc_h2_ = source15_.cf12
                d_865_cf12_ = d_864___mcc_h2_
                r2 = p0
                d_866_v21_: _dafny.Array
                nw128_ = _dafny.Array(_dafny.Map({}), 29)
                d_866_v21_ = nw128_
                d_867_v23_: _dafny.Map
                def iife143_():
                    coll69_ = _dafny.Set()
                    compr_69_: int
                    for compr_69_ in _dafny.IntegerRange(643, 65):
                        d_868_v22_: int = compr_69_
                        if ((643) <= (d_868_v22_)) and ((d_868_v22_) < (65)):
                            coll69_ = coll69_.union(_dafny.Set([(d_868_v22_) - ((self).f33)]))
                    return _dafny.Set(coll69_)
                d_867_v23_ = _dafny.Map({True: len(iife143_()
                )})
                d_869_v24_: _dafny.Map
                d_869_v24_ = _dafny.Map({len(d_867_v23_): d_837_v1_})
                index118_ = default__.safeIndex(652, (d_866_v21_).length(0))
                (d_866_v21_)[index118_] = d_869_v24_
                index119_ = default__.safeIndex(652, (d_866_v21_).length(0))
                (d_866_v21_)[index119_] = d_869_v24_
                d_870_v25_: _dafny.Map
                d_870_v25_ = _dafny.Map({148: p0})
                (globalState).f13 = (d_836_v0_)[default__.safeIndex(default__.safeDivisionInt((self).f33, len(d_870_v25_)), len(d_836_v0_))]
                r2 = not(default__.fm2(globalState))
            elif True:
                d_871___mcc_h3_ = source15_.cf15
                d_872_cf15_ = d_871___mcc_h3_
                r2 = (d_840_v4_).f39
                d_873_v27_: D4
                def iife144_():
                    coll70_ = _dafny.Map()
                    compr_70_: str
                    for compr_70_ in (_dafny.SeqWithoutIsStrInference([self.f23, p1])).Elements:
                        d_874_v26_: str = compr_70_
                        if (d_874_v26_) in (_dafny.SeqWithoutIsStrInference([self.f23, p1])):
                            coll70_[d_874_v26_] = (d_840_v4_).f39
                    return _dafny.Map(coll70_)
                d_873_v27_ = D4_DC11(len(iife144_()
))
                d_875_v28_: D0
                d_875_v28_ = D0_DC2((d_840_v4_).f39, (d_873_v27_).cf18, (d_840_v4_).f39)
                pat_let_tv27_ = p0
                pat_let_tv28_ = d_840_v4_
                def iife146_(_pat_let38_0):
                    def iife147_(d_876_dt__update__tmp_h0_):
                        def iife148_(_pat_let39_0):
                            def iife149_(d_877_dt__update_hcf4_h0_):
                                def iife150_(_pat_let40_0):
                                    def iife151_(d_878_dt__update_hcf3_h0_):
                                        return D0_DC2((d_876_dt__update__tmp_h0_).cf2, d_878_dt__update_hcf3_h0_, d_877_dt__update_hcf4_h0_)
                                    return iife151_(_pat_let40_0)
                                return iife150_(795)
                            return iife149_(_pat_let39_0)
                        return iife148_(pat_let_tv27_)
                    return iife147_(_pat_let38_0)
                def iife145_(_pat_let37_0):
                    def iife152_(d_879_dt__update__tmp_h1_):
                        def iife153_(_pat_let41_0):
                            def iife154_(d_880_dt__update_hcf4_h1_):
                                def iife155_(_pat_let42_0):
                                    def iife156_(d_881_dt__update_hcf2_h0_):
                                        return D0_DC2(d_881_dt__update_hcf2_h0_, (d_879_dt__update__tmp_h1_).cf3, d_880_dt__update_hcf4_h1_)
                                    return iife156_(_pat_let42_0)
                                return iife155_((pat_let_tv28_).f39)
                            return iife154_(_pat_let41_0)
                        return iife153_((self).f24)
                    return iife152_(_pat_let37_0)
                (globalState).f7 = len(_dafny.Map({(d_841_v5_) == (default__.fm1(True, globalState)): iife145_(iife146_(d_875_v28_))}))
                r0 = (d_840_v4_).f39
                d_882_v29_: C2
                nw129_ = C2()
                nw129_.ctor__((d_837_v1_).set((d_840_v4_).f39, default__.abs((self).f33)), (self).f25, _dafny.CodePoint('f'), (d_841_v5_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))))
                d_882_v29_ = nw129_
            r0 = (d_840_v4_).f39
        elif True:
            d_883_v30_: _dafny.Array
            nw130_ = _dafny.Array(False, 16)
            d_883_v30_ = nw130_
            d_884_v31_: _dafny.Array
            nw131_ = _dafny.Array(None, 9)
            nw131_[int(0)] = d_883_v30_
            nw131_[int(1)] = d_883_v30_
            nw131_[int(2)] = d_883_v30_
            nw131_[int(3)] = d_883_v30_
            nw131_[int(4)] = d_883_v30_
            nw131_[int(5)] = d_883_v30_
            nw131_[int(6)] = d_883_v30_
            nw131_[int(7)] = d_883_v30_
            nw131_[int(8)] = d_883_v30_
            d_884_v31_ = nw131_
            d_884_v31_ = d_884_v31_
            d_885_v32_: _dafny.Array
            nw132_ = _dafny.Array(int(0), 10)
            d_885_v32_ = nw132_
            d_886_v33_: _dafny.Seq
            d_886_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhp"))
            d_887_v34_: _dafny.MultiSet
            d_887_v34_ = _dafny.MultiSet([(self).f33, len(d_886_v33_)])
            d_888_v35_: _dafny.MultiSet
            d_888_v35_ = _dafny.MultiSet([not(p0), (self).f24, (self).f24, (self).f24])
            rhs143_ = d_885_v32_
            rhs144_ = (self).f33
            rhs145_ = d_883_v30_
            rhs146_ = default__.safeModuloInt((((d_887_v34_)[(self).f33] if ((self).f33) in (d_887_v34_) else -963)) * ((d_888_v35_).cardinality), (self).f33)
            lhs104_ = globalState
            lhs104_.f22 = rhs143_
            r1 = rhs144_
            d_883_v30_ = rhs145_
            r1 = rhs146_
            d_889_v36_: _dafny.Map
            d_889_v36_ = _dafny.Map({p1: (self).f33})
            d_889_v36_ = (d_889_v36_).set(p1, -924)
            index120_ = default__.safeIndex(941, (d_885_v32_).length(0))
            (d_885_v32_)[index120_] = (self).f33
            index121_ = default__.safeIndex(941, (d_885_v32_).length(0))
            (d_885_v32_)[index121_] = len((d_886_v33_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fb"))))
            d_890_v37_: _dafny.Map
            d_890_v37_ = _dafny.Map({(p0 if not(False) else (self).f24): (self).f33})
            rhs147_ = p1
            rhs148_ = d_890_v37_
            rhs149_ = ((self).f33) + ((self).f33)
            rhs150_ = p0
            rhs151_ = (self).f24
            lhs105_ = self
            lhs106_ = globalState
            lhs107_ = globalState
            lhs105_.f23 = rhs147_
            d_890_v37_ = rhs148_
            lhs106_.f7 = rhs149_
            r2 = rhs150_
            lhs107_.f13 = rhs151_
        if (p0) or (p0):
            d_891_v38_: _dafny.Map
            d_891_v38_ = _dafny.Map({(self).f33: (self).f24})
            d_892_v39_: _dafny.Seq
            d_892_v39_ = _dafny.SeqWithoutIsStrInference([d_891_v38_, d_891_v38_])
            r0 = ((d_892_v39_) + (d_892_v39_)) == (d_892_v39_)
            d_893_v40_: _dafny.Seq
            d_893_v40_ = _dafny.SeqWithoutIsStrInference([(self).f33])
            d_894_v41_: _dafny.MultiSet
            d_894_v41_ = _dafny.MultiSet([d_893_v40_])
            d_895_v42_: _dafny.Array
            nw133_ = _dafny.Array(None, 2)
            nw133_[int(0)] = (self).f33
            nw133_[int(1)] = (self).f33
            d_895_v42_ = nw133_
            d_896_v43_: _dafny.Map
            d_896_v43_ = _dafny.Map({d_894_v41_: d_895_v42_})
            d_896_v43_ = (d_896_v43_).set(d_894_v41_, (d_895_v42_ if (self).f24 else d_895_v42_))
            d_897_v44_: _dafny.Seq
            d_897_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jp"))
            d_898_v45_: _dafny.MultiSet
            d_898_v45_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([self.f23 for d_899_i2_ in range(default__.abs(13))])])
            d_900_v46_: _dafny.Map
            d_900_v46_ = _dafny.Map({d_897_v44_: d_898_v45_})
            d_900_v46_ = (d_900_v46_).set(d_897_v44_, d_898_v45_)
            d_901_v47_: _dafny.Seq
            d_901_v47_ = _dafny.SeqWithoutIsStrInference([False, p0, not((self).f24), (self).f24])
            d_902_v48_: _dafny.Map
            d_902_v48_ = _dafny.Map({(self).f24: (self).f24})
            d_903_v49_: _dafny.Array
            nw134_ = _dafny.Array(None, 28)
            nw134_[int(0)] = (self).f24
            nw134_[int(1)] = p0
            nw134_[int(2)] = p0
            nw134_[int(3)] = p0
            nw134_[int(4)] = p0
            nw134_[int(5)] = (self).f24
            nw134_[int(6)] = p0
            nw134_[int(7)] = p0
            nw134_[int(8)] = p0
            nw134_[int(9)] = (self).f24
            nw134_[int(10)] = p0
            nw134_[int(11)] = not((d_901_v47_)[default__.safeIndex((self).f33, len(d_901_v47_))])
            nw134_[int(12)] = p0
            nw134_[int(13)] = (self).f24
            nw134_[int(14)] = (self).f24
            nw134_[int(15)] = False
            nw134_[int(16)] = (self).f24
            nw134_[int(17)] = p0
            nw134_[int(18)] = True
            nw134_[int(19)] = ((d_902_v48_)[True] if (True) in (d_902_v48_) else True)
            nw134_[int(20)] = (self).f24
            nw134_[int(21)] = (d_901_v47_)[default__.safeIndex((0) - ((self).f33), len(d_901_v47_))]
            nw134_[int(22)] = (self).fm4((self).f33, d_901_v47_, globalState)
            nw134_[int(23)] = p0
            nw134_[int(24)] = False
            nw134_[int(25)] = (self).f24
            nw134_[int(26)] = (self).f24
            nw134_[int(27)] = True
            d_903_v49_ = nw134_
            d_904_v50_: _dafny.Map
            d_904_v50_ = _dafny.Map({d_903_v49_: self.f23})
            d_905_v51_: _dafny.Set
            d_905_v51_ = _dafny.Set({(d_897_v44_)[default__.safeIndex((self).f33, len(d_897_v44_))], self.f23})
            d_906_v52_: _dafny.MultiSet
            d_906_v52_ = _dafny.MultiSet([(self).f24, p0])
            d_907_v53_: D0
            d_907_v53_ = D0_DC0(d_906_v52_)
            d_908_v54_: D0
            d_908_v54_ = D0_DC3(D0_DC3(D0_DC3(d_907_v53_)))
            d_909_v55_: _dafny.Array
            nw135_ = _dafny.Array(None, 8)
            nw135_[int(0)] = D0_DC3(d_907_v53_)
            nw135_[int(1)] = d_908_v54_
            nw135_[int(2)] = d_908_v54_
            nw135_[int(3)] = d_908_v54_
            nw135_[int(4)] = d_908_v54_
            nw135_[int(5)] = D0_DC3(d_907_v53_)
            nw135_[int(6)] = d_908_v54_
            nw135_[int(7)] = d_908_v54_
            d_909_v55_ = nw135_
            d_910_v56_: _dafny.Map
            d_910_v56_ = _dafny.Map({d_909_v55_: _dafny.Set({self.f23, p1})})
            d_911_v57_: C3
            nw136_ = C3()
            nw136_.ctor__(d_904_v50_, (self).f25, self.f23, not((((d_910_v56_)[d_909_v55_] if (d_909_v55_) in (d_910_v56_) else d_905_v51_)).ispropersubset(d_905_v51_)))
            d_911_v57_ = nw136_
            nw137_ = C3()
            nw137_.ctor__((d_911_v57_).f35, (self).f25, self.f23, (len(_dafny.Set({d_895_v42_, d_895_v42_}))) == ((0) - ((self).f33)))
            d_911_v57_ = nw137_
            d_912_v58_: _dafny.Set
            d_912_v58_ = _dafny.Set({(self).f33, len((d_911_v57_).fm7(globalState))})
            d_913_v59_: D5
            d_913_v59_ = D5_DC12(d_912_v58_)
            d_914_v60_: _dafny.Map
            d_914_v60_ = _dafny.Map({(self).f33: self.f23})
            r0 = not(((d_912_v58_) - (_dafny.Set({(0) - (len(d_902_v48_)), (self).f33, (self).f33, (self).f33, len(d_914_v60_)}))).issubset((d_913_v59_).cf19))
        elif True:
            (globalState).f7 = (self).f33
            d_915_v61_: _dafny.MultiSet
            d_915_v61_ = _dafny.MultiSet([not((self).f24)])
            d_916_v62_: C2
            nw138_ = C2()
            nw138_.ctor__((_dafny.MultiSet([(self).f24, p0, (self).f24])).intersection(d_915_v61_), (self).f25, (self.f23 if p0 else self.f23), (d_915_v61_).isdisjoint(d_915_v61_))
            d_916_v62_ = nw138_
            d_917_v63_: _dafny.Array
            nw139_ = _dafny.Array(False, 22)
            d_917_v63_ = nw139_
            index122_ = default__.safeIndex(144, (d_917_v63_).length(0))
            (d_917_v63_)[index122_] = p0
            index123_ = default__.safeIndex(144, (d_917_v63_).length(0))
            (d_917_v63_)[index123_] = ((self).f33) == ((0) - ((self).f33))
            d_918_v64_: _dafny.Seq
            d_918_v64_ = _dafny.SeqWithoutIsStrInference([p0])
            (globalState).f13 = (d_918_v64_)[default__.safeIndex(default__.safeDivisionInt((self).f33, (self).f33), len(d_918_v64_))]
            d_919_v65_: int
            out34_: int
            out34_ = (self).m10((d_917_v63_)[default__.safeIndex(144, (d_917_v63_).length(0))], globalState)
            d_919_v65_ = out34_
        (globalState).f13 = default__.fm2(globalState)
        d_920_v66_: _dafny.Seq
        d_920_v66_ = _dafny.SeqWithoutIsStrInference([(self).f24, (self).f24, p0, not(p0), p0])
        d_921_v67_: _dafny.Map
        d_921_v67_ = _dafny.Map({p0: d_920_v66_})
        (globalState).f13 = not(((self).f24) in (d_921_v67_))
        hi3_ = default__.fm0((self).f24, (self).f33, globalState)
        for d_922_i3_ in range((self).f33, hi3_):
            if (d_922_i3_) > (d_922_i3_):
                (globalState).f7 = d_922_i3_
                d_923_v68_: _dafny.Set
                d_923_v68_ = _dafny.Set({(self).f24, (self).f24})
                d_924_v69_: _dafny.Seq
                d_924_v69_ = _dafny.SeqWithoutIsStrInference([(d_923_v68_) - (d_923_v68_)])
                d_925_v70_: _dafny.Map
                d_925_v70_ = _dafny.Map({default__.fm2(globalState): (_dafny.SeqWithoutIsStrInference([d_923_v68_, d_923_v68_])) + (d_924_v69_)})
                d_926_v71_: _dafny.Set
                d_926_v71_ = _dafny.Set({999, d_922_i3_})
                d_927_v72_: _dafny.MultiSet
                d_927_v72_ = _dafny.MultiSet([d_926_v71_])
                rhs152_ = not(not((self).f24))
                rhs153_ = ((d_925_v70_)[False] if (False) in (d_925_v70_) else d_924_v69_)
                rhs154_ = (0) - (d_922_i3_)
                rhs155_ = (self).f24
                rhs156_ = ((d_927_v72_).cardinality) * ((self).f33)
                lhs108_ = globalState
                lhs109_ = globalState
                lhs110_ = globalState
                lhs111_ = globalState
                lhs108_.f13 = rhs152_
                d_924_v69_ = rhs153_
                lhs109_.f7 = rhs154_
                lhs110_.f13 = rhs155_
                lhs111_.f7 = rhs156_
                (self).f34 = self.f34
                d_928_v73_: _dafny.Array
                nw140_ = _dafny.Array(False, 4)
                d_928_v73_ = nw140_
                d_929_v74_: _dafny.Map
                d_929_v74_ = _dafny.Map({True: (self).f24})
                index124_ = default__.safeIndex(53, (d_928_v73_).length(0))
                (d_928_v73_)[index124_] = (self).fm23((self).f33, d_929_v74_, d_920_v66_, globalState)
                index125_ = default__.safeIndex(53, (d_928_v73_).length(0))
                (d_928_v73_)[index125_] = (self).fm23(d_922_i3_, d_929_v74_, (d_920_v66_) + (d_920_v66_), globalState)
                d_930_v75_: _dafny.Seq
                d_930_v75_ = _dafny.SeqWithoutIsStrInference([(self).f33, (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f24])))])
                d_931_v76_: _dafny.MultiSet
                d_931_v76_ = _dafny.MultiSet([(self).f24, (self).f24, False])
                d_932_v77_: _dafny.Map
                d_932_v77_ = _dafny.Map({d_930_v75_: (((d_931_v76_)[False] if (False) in (d_931_v76_) else (self).f33)) - (d_922_i3_)})
                d_932_v77_ = (d_932_v77_).set(d_930_v75_, d_922_i3_)
            elif True:
                d_933_v78_: _dafny.Seq
                d_933_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jctwmpljq"))
                d_933_v78_ = (d_933_v78_) + ((d_933_v78_).set(default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference([p0, p0, (self).f24, True, p0]))), len(d_933_v78_)), self.f23))
                d_934_v79_: _dafny.Array
                nw141_ = _dafny.Array(False, 9)
                d_934_v79_ = nw141_
                d_934_v79_ = (d_934_v79_ if p0 else d_934_v79_)
                d_935_v80_: _dafny.Array
                def lambda41_(d_936_p0_):
                    def lambda42_(d_937_i4_):
                        return D0_DC0(_dafny.MultiSet([not(d_936_p0_)]))

                    return lambda42_

                init23_ = lambda41_(p0)
                nw142_ = _dafny.Array(None, 27)
                for i0_23_ in range(nw142_.length(0)):
                    nw142_[i0_23_] = init23_(i0_23_)
                d_935_v80_ = nw142_
                d_938_v81_: _dafny.MultiSet
                d_938_v81_ = _dafny.MultiSet([(self).f24])
                d_939_v82_: D0
                d_939_v82_ = D0_DC0(d_938_v81_)
                index126_ = default__.safeIndex(918, (d_935_v80_).length(0))
                (d_935_v80_)[index126_] = d_939_v82_
                index127_ = default__.safeIndex(918, (d_935_v80_).length(0))
                (d_935_v80_)[index127_] = default__.fm43(p0, (self).f24, p0, globalState)
                d_940_v83_: _dafny.Seq
                d_940_v83_ = _dafny.SeqWithoutIsStrInference([d_933_v78_, d_933_v78_])
                rhs157_ = (self).f24
                rhs158_ = (d_922_i3_) > (d_922_i3_)
                rhs159_ = d_940_v83_
                lhs112_ = globalState
                lhs112_.f13 = rhs157_
                r0 = rhs158_
                d_940_v83_ = rhs159_
                d_941_v84_: _dafny.Map
                d_941_v84_ = _dafny.Map({p0: 43})
                d_942_v85_: _dafny.Map
                d_942_v85_ = _dafny.Map({self.f23: d_941_v84_})
                d_943_v86_: _dafny.Map
                d_943_v86_ = _dafny.Map({(self).f24: (not(p0)) in (((d_942_v85_)[self.f23] if (self.f23) in (d_942_v85_) else d_941_v84_))})
                r1 = len(d_943_v86_)
            d_944_v87_: D5
            d_944_v87_ = D5_DC14(p0)
            d_945_v88_: _dafny.Seq
            d_945_v88_ = _dafny.SeqWithoutIsStrInference([d_944_v87_, d_944_v87_])
            d_946_v89_: _dafny.Map
            d_946_v89_ = _dafny.Map({(self).f24: d_945_v88_})
            d_946_v89_ = (d_946_v89_).set((self).f24, d_945_v88_)
            d_947_v90_: C1
            nw143_ = C1()
            nw143_.ctor__((self).f24, True)
            d_947_v90_ = nw143_
            d_948_v91_: C1
            d_948_v91_ = d_947_v90_
            d_949_v92_: _dafny.Map
            d_949_v92_ = _dafny.Map({(d_922_i3_ if (self).f24 else (self).f33): d_948_v91_})
            d_949_v92_ = (d_949_v92_).set(default__.safeDivisionInt((self).f33, (self).f33), d_948_v91_)
            d_950_v93_: _dafny.Seq
            d_950_v93_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvxmib"))
            (globalState).f7 = default__.safeModuloInt(d_922_i3_, len(d_950_v93_))
        d_951_v94_: _dafny.Array
        def lambda43_(d_952_i5_):
            return False

        init24_ = lambda43_
        nw144_ = _dafny.Array(None, 10)
        for i0_24_ in range(nw144_.length(0)):
            nw144_[i0_24_] = init24_(i0_24_)
        d_951_v94_ = nw144_
        d_953_v95_: _dafny.Map
        d_953_v95_ = _dafny.Map({d_951_v94_: self.f23})
        d_954_v96_: C3
        nw145_ = C3()
        nw145_.ctor__(d_953_v95_, (self).f25, p1, p0)
        d_954_v96_ = nw145_
        r0 = p0
        r1 = (self).f33
        d_955_v98_: _dafny.MultiSet
        d_955_v98_ = _dafny.MultiSet([(self).f33])
        d_956_v99_: _dafny.Set
        d_956_v99_ = _dafny.Set({d_951_v94_})
        d_957_v100_: _dafny.Map
        def iife157_():
            coll71_ = _dafny.Map()
            compr_71_: int
            for compr_71_ in (_dafny.Map({((d_955_v98_)[(self).f33] if ((self).f33) in (d_955_v98_) else (self).f33): 240})).keys.Elements:
                d_958_v97_: int = compr_71_
                if (d_958_v97_) in (_dafny.Map({((d_955_v98_)[(self).f33] if ((self).f33) in (d_955_v98_) else (self).f33): 240})):
                    coll71_[(d_958_v97_) * ((self).f33)] = (self).f33
            return _dafny.Map(coll71_)
        d_957_v100_ = _dafny.Map({iife157_()
        : d_956_v99_})
        d_959_v101_: _dafny.Seq
        d_959_v101_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktaasd"))
        d_960_v102_: _dafny.Map
        d_960_v102_ = _dafny.Map({len(d_959_v101_): (self).f33})
        r2 = (d_956_v99_).issubset(((d_957_v100_)[d_960_v102_] if (d_960_v102_) in (d_957_v100_) else d_956_v99_))
        return r0, r1, r2

    def m10(self, p0, globalState):
        r0: int = int(0)
        d_961_v0_: _dafny.Map
        d_961_v0_ = _dafny.Map({(self).f33: ((self).f33) >= (237)})
        d_961_v0_ = (d_961_v0_).set((self).f33, ((self).f24 if p0 else p0))
        d_962_v1_: _dafny.Map
        d_962_v1_ = _dafny.Map({False: (self).f24})
        d_962_v1_ = (d_962_v1_).set(not((self).f24), (self).f24)
        (globalState).f13 = ((p0 if False else p0)) and (p0)
        hi4_ = (self).f33
        for d_963_i0_ in range((self).f33, hi4_):
            d_962_v1_ = d_962_v1_
            d_964_v2_: _dafny.Array
            nw146_ = _dafny.Array(False, 2)
            d_964_v2_ = nw146_
            index128_ = default__.safeIndex(772, (d_964_v2_).length(0))
            (d_964_v2_)[index128_] = True
            index129_ = default__.safeIndex(772, (d_964_v2_).length(0))
            (d_964_v2_)[index129_] = (self).f24
            (globalState).f13 = not(p0)
            d_965_v3_: _dafny.Seq
            d_965_v3_ = _dafny.SeqWithoutIsStrInference([(self).f24])
            d_966_v4_: _dafny.MultiSet
            d_966_v4_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([(d_964_v2_)[default__.safeIndex(772, (d_964_v2_).length(0))]])) + (d_965_v3_)])
            d_966_v4_ = d_966_v4_
        d_967_i1_: int
        d_967_i1_ = 0
        with _dafny.label("4"):
            while ((self).f33) == ((self).f33):
                with _dafny.c_label("4"):
                    if (d_967_i1_) >= (100):
                        raise _dafny.Break("4")
                    d_967_i1_ = (d_967_i1_) + (1)
                    (globalState).f13 = default__.fm2(globalState)
                    (globalState).f7 = (self).f33
                    d_968_v5_: _dafny.Map
                    d_968_v5_ = _dafny.Map({(_dafny.MultiSet([False])).cardinality: (self).f33})
                    d_969_v6_: _dafny.Seq
                    d_969_v6_ = _dafny.SeqWithoutIsStrInference([(self).f33])
                    d_970_v7_: _dafny.Seq
                    d_970_v7_ = d_969_v6_
                    d_971_v8_: D0
                    d_971_v8_ = D0_DC2(not(not((self).f24)), ((d_968_v5_)[(self).f33] if ((self).f33) in (d_968_v5_) else (0) - (len((d_970_v7_)))), p0)
                    d_972_v9_: _dafny.Array
                    def lambda44_(d_973_i2_):
                        return (d_973_i2_) * ((self).f33)

                    init25_ = lambda44_
                    nw147_ = _dafny.Array(None, 24)
                    for i0_25_ in range(nw147_.length(0)):
                        nw147_[i0_25_] = init25_(i0_25_)
                    d_972_v9_ = nw147_
                    d_974_v10_: _dafny.Map
                    d_974_v10_ = _dafny.Map({d_971_v8_: d_972_v9_})
                    d_975_v11_: _dafny.Seq
                    d_975_v11_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_974_v10_ = (d_974_v10_).set(D0_DC2(not(p0), len(d_975_v11_), (self).f24), d_972_v9_)
                    d_976_v12_: _dafny.Array
                    def lambda45_(d_977_v6_):
                        def lambda46_(d_978_i3_):
                            return d_977_v6_

                        return lambda46_

                    init26_ = lambda45_(d_969_v6_)
                    nw148_ = _dafny.Array(None, 25)
                    for i0_26_ in range(nw148_.length(0)):
                        nw148_[i0_26_] = init26_(i0_26_)
                    d_976_v12_ = nw148_
                    index130_ = default__.safeIndex(225, (d_976_v12_).length(0))
                    (d_976_v12_)[index130_] = (self).fm7(globalState)
                    index131_ = default__.safeIndex(677, (d_972_v9_).length(0))
                    (d_972_v9_)[index131_] = ((self).f33) + ((self).f33)
                    d_979_v13_: _dafny.Seq
                    d_979_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xfjecqwws"))
                    index132_ = default__.safeIndex(225, (d_976_v12_).length(0))
                    index133_ = default__.safeIndex(677, (d_972_v9_).length(0))
                    rhs160_ = d_969_v6_
                    rhs161_ = default__.safeModuloInt((498) - ((0) - (len(_dafny.Set({(self).f33, len(d_979_v13_)})))), (-59) + ((self).f33))
                    rhs162_ = default__.safeModuloInt((self).f33, (self).f33)
                    lhs113_ = d_976_v12_
                    lhs114_ = default__.safeIndex(225, (d_976_v12_).length(0))
                    lhs115_ = globalState
                    lhs116_ = d_972_v9_
                    lhs117_ = default__.safeIndex(677, (d_972_v9_).length(0))
                    lhs113_[lhs114_] = rhs160_
                    lhs115_.f7 = rhs161_
                    lhs116_[lhs117_] = rhs162_
                    pass
            pass
        r0 = (self).f33
        d_980_v14_: _dafny.MultiSet
        d_980_v14_ = _dafny.MultiSet([p0])
        r0 = (0) - ((0) - (default__.safeDivisionInt(((d_980_v14_)[(self).f24] if ((self).f24) in (d_980_v14_) else (self).f33), (self).f33)))
        return r0

    @property
    def f33(self):
        return self._f33

class C5:
    def  __init__(self):
        self._f32: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f32):
        (self)._f32 = f32

    def fm17(self, p0, p1, p2, globalState):
        return ((744) * (-870)) + (((_dafny.MultiSet([_dafny.CodePoint('j')])).intersection(_dafny.MultiSet([_dafny.CodePoint('v')]))).cardinality)

    def m8(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_981_v0_: bool
        d_981_v0_ = True
        d_982_v1_: _dafny.Seq
        d_982_v1_ = _dafny.SeqWithoutIsStrInference([d_981_v0_, d_981_v0_, d_981_v0_])
        d_983_v2_: int
        d_983_v2_ = 987
        d_984_v3_: _dafny.MultiSet
        d_984_v3_ = _dafny.MultiSet([len(d_982_v1_), d_983_v2_])
        d_985_v4_: _dafny.Map
        d_985_v4_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_981_v0_, d_981_v0_, d_981_v0_])): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_986_i0_ in range(default__.abs(-624))])})
        (globalState).f13 = ((default__.fm18(globalState)) - (d_984_v3_)).ispropersubset((d_984_v3_).set(len(((d_985_v4_)[d_983_v2_] if (d_983_v2_) in (d_985_v4_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnfuhyqq")))), default__.abs(d_983_v2_)))
        d_987_v5_: _dafny.Set
        d_987_v5_ = _dafny.Set({164})
        def iife158_():
            coll72_ = _dafny.Set()
            compr_72_: int
            for compr_72_ in _dafny.IntegerRange(413, 860):
                d_988_v6_: int = compr_72_
                if ((413) <= (d_988_v6_)) and ((d_988_v6_) < (860)):
                    coll72_ = coll72_.union(_dafny.Set([(d_988_v6_) - (d_983_v2_)]))
            return _dafny.Set(coll72_)
        d_987_v5_ = (d_987_v5_) | (iife158_()
        )
        hi5_ = d_983_v2_
        for d_989_i1_ in range(d_983_v2_, hi5_):
            if d_981_v0_:
                d_990_v7_: _dafny.Seq
                d_990_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwdbajgpg"))
                d_991_v8_: _dafny.Map
                d_991_v8_ = _dafny.Map({d_981_v0_: ((d_984_v3_)[len(d_987_v5_)] if (len(d_987_v5_)) in (d_984_v3_) else d_983_v2_)})
                d_992_v9_: str
                d_992_v9_ = _dafny.CodePoint('k')
                d_993_v10_: _dafny.Array
                nw149_ = _dafny.Array(None, 6)
                nw149_[int(0)] = d_983_v2_
                nw149_[int(1)] = default__.fm0(d_981_v0_, d_983_v2_, globalState)
                nw149_[int(2)] = (self).fm17(len((d_990_v7_).set(default__.safeIndex(len(d_991_v8_), len(d_990_v7_)), d_992_v9_)), _dafny.CodePoint('g'), default__.fm2(globalState), globalState)
                nw149_[int(3)] = d_989_i1_
                nw149_[int(4)] = (d_989_i1_) - (d_989_i1_)
                nw149_[int(5)] = default__.fm0(d_981_v0_, d_989_i1_, globalState)
                d_993_v10_ = nw149_
                index134_ = default__.safeIndex(716, (d_993_v10_).length(0))
                (d_993_v10_)[index134_] = 454
                index135_ = default__.safeIndex(716, (d_993_v10_).length(0))
                (d_993_v10_)[index135_] = default__.safeDivisionInt((self).fm17(d_989_i1_, d_992_v9_, d_981_v0_, globalState), d_989_i1_)
                (globalState).f7 = d_989_i1_
                d_994_v11_: _dafny.Map
                d_994_v11_ = _dafny.Map({d_981_v0_: d_981_v0_})
                d_995_v12_: _dafny.Map
                d_995_v12_ = _dafny.Map({_dafny.Map({d_981_v0_: (d_982_v1_).set(default__.safeIndex(d_989_i1_, len(d_982_v1_)), d_981_v0_)}): (d_994_v11_) | (default__.fm19(globalState))})
                d_996_v13_: _dafny.Map
                d_996_v13_ = _dafny.Map({d_981_v0_: d_982_v1_})
                d_995_v12_ = (d_995_v12_).set((d_996_v13_ if d_981_v0_ else d_996_v13_), d_994_v11_)
                index136_ = default__.safeIndex(295, ((self).f32).length(0))
                ((self).f32)[index136_] = d_981_v0_
                d_997_v14_: _dafny.Seq
                d_997_v14_ = _dafny.SeqWithoutIsStrInference([(d_993_v10_)[default__.safeIndex(716, (d_993_v10_).length(0))], -537])
                index137_ = default__.safeIndex(295, ((self).f32).length(0))
                ((self).f32)[index137_] = (_dafny.MultiSet([d_997_v14_, d_997_v14_])).ispropersubset(_dafny.MultiSet([d_997_v14_]))
                (globalState).f7 = (d_993_v10_)[default__.safeIndex(716, (d_993_v10_).length(0))]
            elif True:
                d_998_v15_: _dafny.Array
                out35_: _dafny.Array
                out35_ = (self).m9((-814) <= (d_983_v2_), d_981_v0_, globalState)
                d_998_v15_ = out35_
                (globalState).f13 = d_981_v0_
                (globalState).f7 = d_983_v2_
                d_999_v16_: _dafny.Seq
                d_999_v16_ = _dafny.SeqWithoutIsStrInference([d_983_v2_, d_983_v2_])
                (globalState).f7 = default__.fm0(False, default__.safeDivisionInt(d_989_i1_, len(d_999_v16_)), globalState)
                d_1000_v17_: _dafny.Array
                nw150_ = _dafny.Array(None, 29)
                nw150_[int(0)] = _dafny.SeqWithoutIsStrInference([d_989_i1_ for d_1001_i2_ in range(default__.abs(-368))])
                nw150_[int(1)] = _dafny.SeqWithoutIsStrInference([d_989_i1_, d_989_i1_, d_983_v2_])
                nw150_[int(2)] = _dafny.SeqWithoutIsStrInference([d_983_v2_ for d_1002_i3_ in range(default__.abs(182))])
                nw150_[int(3)] = d_999_v16_
                nw150_[int(4)] = d_999_v16_
                nw150_[int(5)] = d_999_v16_
                nw150_[int(6)] = d_999_v16_
                nw150_[int(7)] = d_999_v16_
                nw150_[int(8)] = _dafny.SeqWithoutIsStrInference([d_983_v2_, d_989_i1_])
                nw150_[int(9)] = _dafny.SeqWithoutIsStrInference([(0) - (d_983_v2_) for d_1003_i4_ in range(default__.abs(-645))])
                nw150_[int(10)] = _dafny.SeqWithoutIsStrInference([d_989_i1_ for d_1004_i5_ in range(default__.abs(14))])
                nw150_[int(11)] = d_999_v16_
                nw150_[int(12)] = d_999_v16_
                nw150_[int(13)] = d_999_v16_
                nw150_[int(14)] = d_999_v16_
                nw150_[int(15)] = d_999_v16_
                nw150_[int(16)] = d_999_v16_
                nw150_[int(17)] = d_999_v16_
                nw150_[int(18)] = d_999_v16_
                nw150_[int(19)] = _dafny.SeqWithoutIsStrInference([d_989_i1_, d_983_v2_])
                nw150_[int(20)] = d_999_v16_
                nw150_[int(21)] = d_999_v16_
                nw150_[int(22)] = d_999_v16_
                nw150_[int(23)] = d_999_v16_
                nw150_[int(24)] = _dafny.SeqWithoutIsStrInference([d_989_i1_])
                nw150_[int(25)] = d_999_v16_
                nw150_[int(26)] = _dafny.SeqWithoutIsStrInference([len(d_982_v1_), d_983_v2_])
                nw150_[int(27)] = d_999_v16_
                nw150_[int(28)] = d_999_v16_
                d_1000_v17_ = nw150_
                d_1005_v18_: _dafny.Map
                d_1005_v18_ = _dafny.Map({d_1000_v17_: d_981_v0_})
                d_1006_v19_: _dafny.MultiSet
                d_1006_v19_ = _dafny.MultiSet([((d_1005_v18_)[d_1000_v17_] if (d_1000_v17_) in (d_1005_v18_) else d_981_v0_)])
                d_1007_v20_: D0
                d_1007_v20_ = D0_DC0(d_1006_v19_)
                d_1008_v21_: _dafny.Array
                nw151_ = _dafny.Array(None, 19)
                nw151_[int(0)] = d_1007_v20_
                nw151_[int(1)] = d_1007_v20_
                nw151_[int(2)] = d_1007_v20_
                nw151_[int(3)] = (d_1007_v20_ if not(d_981_v0_) else d_1007_v20_)
                nw151_[int(4)] = d_1007_v20_
                nw151_[int(5)] = D0_DC0(_dafny.MultiSet([not(True), d_981_v0_]))
                nw151_[int(6)] = d_1007_v20_
                nw151_[int(7)] = d_1007_v20_
                nw151_[int(8)] = d_1007_v20_
                nw151_[int(9)] = d_1007_v20_
                nw151_[int(10)] = d_1007_v20_
                nw151_[int(11)] = d_1007_v20_
                nw151_[int(12)] = D0_DC0(d_1006_v19_)
                nw151_[int(13)] = default__.fm20((d_1006_v19_).cardinality, default__.fm20(d_983_v2_, d_1007_v20_, globalState), globalState)
                nw151_[int(14)] = d_1007_v20_
                nw151_[int(15)] = d_1007_v20_
                nw151_[int(16)] = d_1007_v20_
                nw151_[int(17)] = d_1007_v20_
                nw151_[int(18)] = d_1007_v20_
                d_1008_v21_ = nw151_
                d_1009_v22_: _dafny.Seq
                d_1009_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgjysbxnh"))
                d_1010_v23_: _dafny.Array
                def lambda47_(d_1011_v2_):
                    def lambda48_(d_1012_i6_):
                        return (d_1012_i6_) * (d_1011_v2_)

                    return lambda48_

                init27_ = lambda47_(d_983_v2_)
                nw152_ = _dafny.Array(None, 4)
                for i0_27_ in range(nw152_.length(0)):
                    nw152_[i0_27_] = init27_(i0_27_)
                d_1010_v23_ = nw152_
                rhs163_ = d_1008_v21_
                rhs164_ = d_1010_v23_
                rhs165_ = (0) - (default__.safeModuloInt(default__.safeModuloInt(d_983_v2_, d_989_i1_), d_989_i1_))
                rhs166_ = d_981_v0_
                rhs167_ = d_1009_v22_
                lhs118_ = globalState
                lhs119_ = globalState
                lhs120_ = globalState
                d_1008_v21_ = rhs163_
                lhs118_.f22 = rhs164_
                lhs119_.f7 = rhs165_
                lhs120_.f13 = rhs166_
                d_1009_v22_ = rhs167_
            d_983_v2_ = d_989_i1_
            (globalState).f13 = d_981_v0_
            d_1013_v24_: str
            d_1013_v24_ = _dafny.CodePoint('i')
            d_1013_v24_ = d_1013_v24_
        d_1014_v25_: _dafny.Map
        d_1014_v25_ = _dafny.Map({(len(_dafny.Map({d_983_v2_: d_981_v0_}))) <= (d_983_v2_): d_981_v0_})
        d_1014_v25_ = (d_1014_v25_).set(d_981_v0_, d_981_v0_)
        d_1015_v26_: _dafny.Array
        nw153_ = _dafny.Array(int(0), 28)
        d_1015_v26_ = nw153_
        d_1016_v27_: _dafny.Map
        d_1016_v27_ = _dafny.Map({d_983_v2_: d_983_v2_})
        d_1017_v28_: _dafny.Map
        d_1017_v28_ = _dafny.Map({d_1016_v27_: d_981_v0_})
        index138_ = default__.safeIndex(505, (d_1015_v26_).length(0))
        (d_1015_v26_)[index138_] = (len(d_1017_v28_)) + (d_983_v2_)
        d_1018_v29_: _dafny.Seq
        d_1018_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejirnglkd"))
        d_1019_v30_: D2
        d_1019_v30_ = D2_DC7(d_981_v0_, d_981_v0_)
        d_1020_v31_: D2
        d_1020_v31_ = D2_DC7((d_1018_v29_) < (d_1018_v29_), (d_1019_v30_).cf13)
        index139_ = default__.safeIndex(505, (d_1015_v26_).length(0))
        rhs168_ = d_983_v2_
        rhs169_ = d_1019_v30_
        rhs170_ = (d_983_v2_) + (len(((d_982_v1_ if d_981_v0_ else _dafny.SeqWithoutIsStrInference([d_981_v0_]))).set(default__.safeIndex((0) - (d_983_v2_), len((d_982_v1_ if d_981_v0_ else _dafny.SeqWithoutIsStrInference([d_981_v0_])))), d_981_v0_)))
        rhs171_ = d_981_v0_
        rhs172_ = default__.safeDivisionInt(d_983_v2_, (d_983_v2_) * (304))
        lhs121_ = d_1015_v26_
        lhs122_ = default__.safeIndex(505, (d_1015_v26_).length(0))
        lhs123_ = globalState
        lhs124_ = globalState
        lhs125_ = globalState
        lhs121_[lhs122_] = rhs168_
        d_1020_v31_ = rhs169_
        lhs123_.f7 = rhs170_
        lhs124_.f13 = rhs171_
        lhs125_.f7 = rhs172_
        (globalState).f13 = d_981_v0_
        r0 = (self).f32
        return r0

    def m9(self, p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_1021_v0_: _dafny.Seq
        d_1021_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ri"))
        d_1022_v1_: _dafny.Seq
        d_1022_v1_ = _dafny.SeqWithoutIsStrInference([d_1021_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwjxmn"))])
        d_1023_v2_: _dafny.Set
        d_1023_v2_ = _dafny.Set({p1, p0})
        d_1024_v3_: int
        d_1024_v3_ = 858
        d_1025_v4_: _dafny.MultiSet
        d_1025_v4_ = _dafny.MultiSet([d_1024_v3_, d_1024_v3_])
        d_1022_v1_ = (_dafny.SeqWithoutIsStrInference([d_1021_v0_, d_1021_v0_, d_1021_v0_])) + (default__.fm21(d_1023_v2_, d_1025_v4_, globalState))
        if p0:
            d_1026_v5_: str
            d_1026_v5_ = _dafny.CodePoint('v')
            d_1027_v6_: _dafny.Map
            d_1027_v6_ = _dafny.Map({p1: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "av"))).set(default__.safeIndex(d_1024_v3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "av")))), d_1026_v5_)})
            d_1028_v7_: C0
            nw154_ = C0()
            nw154_.ctor__((d_1027_v6_ if p0 else d_1027_v6_), (not(p0)) and (p1), d_1026_v5_, p0)
            d_1028_v7_ = nw154_
            (globalState).f7 = d_1024_v3_
            d_1029_v8_: _dafny.Seq
            d_1029_v8_ = _dafny.SeqWithoutIsStrInference([(d_1028_v7_).f38, not((d_1028_v7_).f38)])
            d_1030_v9_: _dafny.Array
            nw155_ = _dafny.Array(None, 13)
            nw155_[int(0)] = d_1029_v8_
            nw155_[int(1)] = d_1029_v8_
            nw155_[int(2)] = _dafny.SeqWithoutIsStrInference([p1, p1, (d_1028_v7_).f38])
            nw155_[int(3)] = (d_1029_v8_) + (d_1029_v8_)
            nw155_[int(4)] = d_1029_v8_
            nw155_[int(5)] = (_dafny.SeqWithoutIsStrInference([(d_1028_v7_).f38, p0])) + (d_1029_v8_)
            nw155_[int(6)] = (d_1029_v8_) + (d_1029_v8_)
            nw155_[int(7)] = d_1029_v8_
            nw155_[int(8)] = _dafny.SeqWithoutIsStrInference([not((d_1028_v7_).f38), p1, not(p1)])
            nw155_[int(9)] = _dafny.SeqWithoutIsStrInference([p0])
            nw155_[int(10)] = (d_1029_v8_) + (d_1029_v8_)
            nw155_[int(11)] = d_1029_v8_
            nw155_[int(12)] = (_dafny.SeqWithoutIsStrInference([p1])) + (d_1029_v8_)
            d_1030_v9_ = nw155_
            index140_ = default__.safeIndex(686, (d_1030_v9_).length(0))
            (d_1030_v9_)[index140_] = d_1029_v8_
            d_1031_v10_: _dafny.MultiSet
            d_1031_v10_ = _dafny.MultiSet([p0])
            d_1032_v11_: _dafny.Set
            d_1032_v11_ = _dafny.Set({829, d_1024_v3_, d_1024_v3_, d_1024_v3_})
            index141_ = default__.safeIndex(686, (d_1030_v9_).length(0))
            rhs173_ = d_1026_v5_
            rhs174_ = ((_dafny.SeqWithoutIsStrInference([p0])) + (d_1029_v8_)) + ((d_1029_v8_).set(default__.safeIndex(((d_1031_v10_)[True] if (True) in (d_1031_v10_) else d_1024_v3_), len(d_1029_v8_)), p1))
            rhs175_ = False
            rhs176_ = (d_1032_v11_).isdisjoint(d_1032_v11_)
            lhs126_ = d_1030_v9_
            lhs127_ = default__.safeIndex(686, (d_1030_v9_).length(0))
            lhs128_ = globalState
            lhs129_ = globalState
            d_1026_v5_ = rhs173_
            lhs126_[lhs127_] = rhs174_
            lhs128_.f13 = rhs175_
            lhs129_.f13 = rhs176_
            (globalState).f7 = d_1024_v3_
            if (d_1028_v7_).f38:
                d_1033_v12_: _dafny.Seq
                d_1033_v12_ = _dafny.SeqWithoutIsStrInference([(0) - (default__.safeDivisionInt(d_1024_v3_, d_1024_v3_))])
                def iife159_():
                    coll73_ = _dafny.Set()
                    compr_73_: int
                    for compr_73_ in _dafny.IntegerRange(885, 965):
                        d_1034_v13_: int = compr_73_
                        if ((885) <= (d_1034_v13_)) and ((d_1034_v13_) < (965)):
                            coll73_ = coll73_.union(_dafny.Set([default__.safeModuloInt(d_1034_v13_, d_1024_v3_)]))
                    return _dafny.Set(coll73_)
                d_1033_v12_ = (d_1028_v7_).fm26(len(d_1021_v0_), ((d_1030_v9_)[default__.safeIndex(686, (d_1030_v9_).length(0))] if p0 else (d_1030_v9_)[default__.safeIndex(686, (d_1030_v9_).length(0))]), iife159_()
                , (d_1028_v7_).f38, globalState)
                (globalState).f13 = p1
                (globalState).f13 = (d_1028_v7_).f38
                (globalState).f7 = d_1024_v3_
                d_1035_v14_: _dafny.Array
                nw156_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                d_1035_v14_ = nw156_
                index142_ = default__.safeIndex(600, (d_1035_v14_).length(0))
                (d_1035_v14_)[index142_] = d_1026_v5_
                index143_ = default__.safeIndex(600, (d_1035_v14_).length(0))
                rhs177_ = (d_1022_v1_)[default__.safeIndex(d_1024_v3_, len(d_1022_v1_))]
                rhs178_ = _dafny.CodePoint('l')
                lhs130_ = d_1035_v14_
                lhs131_ = default__.safeIndex(600, (d_1035_v14_).length(0))
                d_1021_v0_ = rhs177_
                lhs130_[lhs131_] = rhs178_
            elif True:
                d_1036_v15_: _dafny.Array
                nw157_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                d_1036_v15_ = nw157_
                d_1037_v16_: C2
                nw158_ = C2()
                nw158_.ctor__(d_1031_v10_, d_1036_v15_, d_1026_v5_, (d_1028_v7_).f38)
                d_1037_v16_ = nw158_
                d_1038_v17_: _dafny.Seq
                d_1038_v17_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt((self).fm17(d_1024_v3_, d_1026_v5_, (d_1028_v7_).f38, globalState), d_1024_v3_)])
                d_1038_v17_ = ((d_1037_v16_).fm7(globalState)) + (d_1038_v17_)
                d_1039_v18_: _dafny.MultiSet
                d_1039_v18_ = _dafny.MultiSet([d_1032_v11_, (d_1032_v11_) - (d_1032_v11_), d_1032_v11_, d_1032_v11_, (d_1032_v11_) | (d_1032_v11_)])
                d_1039_v18_ = ((d_1039_v18_) - (d_1039_v18_)).intersection((d_1039_v18_) | (_dafny.MultiSet([d_1032_v11_, d_1032_v11_])))
                index144_ = default__.safeIndex(648, ((self).f32).length(0))
                ((self).f32)[index144_] = (d_1028_v7_).f38
                index145_ = default__.safeIndex(648, ((self).f32).length(0))
                ((self).f32)[index145_] = (d_1028_v7_).f38
                (globalState).f7 = d_1024_v3_
        elif True:
            d_1040_v19_: C1
            nw159_ = C1()
            nw159_.ctor__(p1, p1)
            d_1040_v19_ = nw159_
            index146_ = default__.safeIndex(534, ((self).f32).length(0))
            ((self).f32)[index146_] = p1
            d_1041_v20_: _dafny.Seq
            d_1041_v20_ = _dafny.SeqWithoutIsStrInference([(d_1040_v19_).f39])
            index147_ = default__.safeIndex(534, ((self).f32).length(0))
            ((self).f32)[index147_] = not(((d_1024_v3_) < (len(d_1041_v20_)) if False else True))
            index148_ = default__.safeIndex(534, ((self).f32).length(0))
            ((self).f32)[index148_] = ((self).f32)[default__.safeIndex(534, ((self).f32).length(0))]
            d_1042_v21_: _dafny.Array
            def lambda49_(d_1043_v20_):
                def lambda50_(d_1044_i0_):
                    return _dafny.Map({False: d_1043_v20_})

                return lambda50_

            init28_ = lambda49_(d_1041_v20_)
            nw160_ = _dafny.Array(None, 1)
            for i0_28_ in range(nw160_.length(0)):
                nw160_[i0_28_] = init28_(i0_28_)
            d_1042_v21_ = nw160_
            d_1045_v22_: _dafny.Map
            d_1045_v22_ = _dafny.Map({(d_1040_v19_).f39: default__.fm36(globalState)})
            d_1046_v23_: D9
            d_1046_v23_ = D9_DC19(d_1045_v22_)
            index149_ = default__.safeIndex(497, (d_1042_v21_).length(0))
            (d_1042_v21_)[index149_] = (d_1046_v23_).cf30
            index150_ = default__.safeIndex(497, (d_1042_v21_).length(0))
            index151_ = default__.safeIndex(534, ((self).f32).length(0))
            rhs179_ = d_1045_v22_
            rhs180_ = p0
            rhs181_ = (d_1040_v19_).f39
            rhs182_ = p0
            lhs132_ = d_1042_v21_
            lhs133_ = default__.safeIndex(497, (d_1042_v21_).length(0))
            lhs134_ = globalState
            lhs135_ = globalState
            lhs136_ = (self).f32
            lhs137_ = default__.safeIndex(534, ((self).f32).length(0))
            lhs132_[lhs133_] = rhs179_
            lhs134_.f13 = rhs180_
            lhs135_.f13 = rhs181_
            lhs136_[lhs137_] = rhs182_
            if (d_1041_v20_) <= (d_1041_v20_):
                d_1047_v24_: C1
                nw161_ = C1()
                nw161_.ctor__((((self).f32)[default__.safeIndex(534, ((self).f32).length(0))] if p1 else p1), p1)
                d_1047_v24_ = nw161_
                (globalState).f13 = (((self).f32)[default__.safeIndex(534, ((self).f32).length(0))]) == ((d_1047_v24_).f39)
                (globalState).f7 = d_1024_v3_
                (globalState).f7 = (d_1024_v3_) + ((0) - (d_1024_v3_))
                d_1048_v25_: _dafny.Array
                def lambda51_(d_1049_v19_, d_1050_p1_, d_1051_p0_, d_1052_v24_):
                    def lambda52_(d_1053_i1_):
                        return (_dafny.MultiSet([D6_DC16(_dafny.Map({((self).f32)[default__.safeIndex(534, ((self).f32).length(0))]: ((self).f32)[default__.safeIndex(534, ((self).f32).length(0))]}), (d_1049_v19_).f39, d_1050_p1_), D6_DC16(_dafny.Map({not(d_1051_p0_): d_1051_p0_}), (d_1049_v19_).f39, (d_1052_v24_).f39)])).intersection(_dafny.MultiSet([D6_DC16(_dafny.Map({((self).f32)[default__.safeIndex(534, ((self).f32).length(0))]: (d_1049_v19_).f39}), (d_1052_v24_).f39, not((d_1049_v19_).f39))]))

                    return lambda52_

                init29_ = lambda51_(d_1040_v19_, p1, p0, d_1047_v24_)
                nw162_ = _dafny.Array(None, 25)
                for i0_29_ in range(nw162_.length(0)):
                    nw162_[i0_29_] = init29_(i0_29_)
                d_1048_v25_ = nw162_
                d_1054_v26_: _dafny.Map
                d_1054_v26_ = _dafny.Map({p0: p1})
                d_1055_v27_: D6
                d_1055_v27_ = D6_DC16(d_1054_v26_, True, p1)
                d_1056_v28_: _dafny.MultiSet
                d_1056_v28_ = _dafny.MultiSet([d_1055_v27_, d_1055_v27_])
                index152_ = default__.safeIndex(840, (d_1048_v25_).length(0))
                (d_1048_v25_)[index152_] = (d_1056_v28_) | ((d_1056_v28_).set(d_1055_v27_, default__.abs(d_1024_v3_)))
                index153_ = default__.safeIndex(840, (d_1048_v25_).length(0))
                (d_1048_v25_)[index153_] = d_1056_v28_
            elif True:
                d_1057_v29_: _dafny.Map
                d_1057_v29_ = _dafny.Map({(d_1040_v19_).f39: d_1021_v0_})
                d_1058_v30_: D2
                d_1058_v30_ = D2_DC6(d_1057_v29_)
                d_1058_v30_ = d_1058_v30_
                d_1059_v31_: str
                d_1059_v31_ = _dafny.CodePoint('c')
                d_1059_v31_ = _dafny.CodePoint('d')
                d_1060_v32_: _dafny.MultiSet
                d_1060_v32_ = _dafny.MultiSet([not((d_1040_v19_).f39)])
                d_1061_v33_: D1
                d_1061_v33_ = D1_DC4(d_1059_v31_)
                d_1062_v34_: _dafny.Array
                nw163_ = _dafny.Array(None, 28)
                nw163_[int(0)] = d_1059_v31_
                nw163_[int(1)] = d_1059_v31_
                nw163_[int(2)] = _dafny.CodePoint('j')
                nw163_[int(3)] = d_1059_v31_
                nw163_[int(4)] = d_1059_v31_
                nw163_[int(5)] = d_1059_v31_
                nw163_[int(6)] = _dafny.CodePoint('r')
                nw163_[int(7)] = d_1059_v31_
                nw163_[int(8)] = d_1059_v31_
                nw163_[int(9)] = d_1059_v31_
                nw163_[int(10)] = d_1059_v31_
                nw163_[int(11)] = d_1059_v31_
                nw163_[int(12)] = d_1059_v31_
                nw163_[int(13)] = d_1059_v31_
                nw163_[int(14)] = d_1059_v31_
                nw163_[int(15)] = (d_1061_v33_).cf6
                nw163_[int(16)] = d_1059_v31_
                nw163_[int(17)] = d_1059_v31_
                nw163_[int(18)] = _dafny.CodePoint('o')
                nw163_[int(19)] = d_1059_v31_
                nw163_[int(20)] = d_1059_v31_
                nw163_[int(21)] = d_1059_v31_
                nw163_[int(22)] = d_1059_v31_
                nw163_[int(23)] = (d_1061_v33_).cf6
                nw163_[int(24)] = d_1059_v31_
                nw163_[int(25)] = _dafny.CodePoint('s')
                nw163_[int(26)] = _dafny.CodePoint('a')
                nw163_[int(27)] = d_1059_v31_
                d_1062_v34_ = nw163_
                d_1063_v35_: C2
                nw164_ = C2()
                nw164_.ctor__(d_1060_v32_, d_1062_v34_, d_1059_v31_, (d_1040_v19_).f39)
                d_1063_v35_ = nw164_
                d_1064_v36_: D4
                d_1064_v36_ = D4_DC10(_dafny.SeqWithoutIsStrInference([d_1059_v31_ for d_1065_i2_ in range(default__.abs(198))]))
                d_1066_v37_: _dafny.Set
                d_1066_v37_ = _dafny.Set({-793})
                def iife160_(_pat_let43_0):
                    def iife161_(d_1067_dt__update__tmp_h0_):
                        def iife162_(_pat_let44_0):
                            def iife163_(d_1068_dt__update_hcf17_h0_):
                                return D4_DC10(d_1068_dt__update_hcf17_h0_)
                            return iife163_(_pat_let44_0)
                        return iife162_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrk")))
                    return iife161_(_pat_let43_0)
                d_1064_v36_ = iife160_(default__.fm44(len(d_1066_v37_), globalState))
                d_1069_v38_: _dafny.Map
                d_1069_v38_ = _dafny.Map({(d_1040_v19_).f39: 899})
                d_1070_v39_: _dafny.Seq
                d_1070_v39_ = _dafny.SeqWithoutIsStrInference([len(d_1069_v38_)])
                d_1071_v40_: _dafny.Array
                def lambda53_(d_1072_v3_):
                    def lambda54_(d_1073_i3_):
                        return default__.safeDivisionInt(d_1073_i3_, d_1072_v3_)

                    return lambda54_

                init30_ = lambda53_(d_1024_v3_)
                nw165_ = _dafny.Array(None, 15)
                for i0_30_ in range(nw165_.length(0)):
                    nw165_[i0_30_] = init30_(i0_30_)
                d_1071_v40_ = nw165_
                d_1074_v41_: _dafny.Map
                d_1074_v41_ = _dafny.Map({d_1071_v40_: ((d_1021_v0_).set(default__.safeIndex(d_1024_v3_, len(d_1021_v0_)), d_1059_v31_)).set(default__.safeIndex(d_1024_v3_, len((d_1021_v0_).set(default__.safeIndex(d_1024_v3_, len(d_1021_v0_)), d_1059_v31_))), d_1059_v31_)})
                d_1075_v42_: _dafny.Map
                d_1075_v42_ = _dafny.Map({(d_1040_v19_).fm9(d_1070_v39_, d_1024_v3_, (d_1040_v19_).f39, globalState): d_1074_v41_})
                d_1075_v42_ = (d_1075_v42_).set((d_1040_v19_).fm9(d_1070_v39_, len(_dafny.SeqWithoutIsStrInference([d_1024_v3_ for d_1076_i4_ in range(default__.abs(-662))])), (d_1040_v19_).f39, globalState), d_1074_v41_)
        (globalState).f7 = default__.safeModuloInt((0) - (d_1024_v3_), d_1024_v3_)
        d_1077_v43_: _dafny.Map
        d_1077_v43_ = _dafny.Map({d_1024_v3_: (0) - (d_1024_v3_)})
        hi6_ = ((d_1077_v43_)[d_1024_v3_] if (d_1024_v3_) in (d_1077_v43_) else d_1024_v3_)
        for d_1078_i5_ in range((56) + (431), hi6_):
            d_1079_v44_: _dafny.Array
            nw166_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_1079_v44_ = nw166_
            d_1079_v44_ = d_1079_v44_
            (globalState).f13 = not(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1080_i6_ in range(default__.abs(575))])) + (d_1021_v0_)) == (d_1021_v0_))
            (globalState).f13 = p0
            d_1025_v4_ = (d_1025_v4_).intersection(d_1025_v4_)
        d_1081_v45_: _dafny.Array
        nw167_ = _dafny.Array(_dafny.Set({}), 9)
        d_1081_v45_ = nw167_
        d_1082_v46_: _dafny.Array
        def lambda55_(d_1083_i7_):
            return (d_1083_i7_) - (len(_dafny.SeqWithoutIsStrInference([False])))

        init31_ = lambda55_
        nw168_ = _dafny.Array(None, 5)
        for i0_31_ in range(nw168_.length(0)):
            nw168_[i0_31_] = init31_(i0_31_)
        d_1082_v46_ = nw168_
        d_1084_v47_: _dafny.Set
        d_1084_v47_ = _dafny.Set({d_1082_v46_})
        index154_ = default__.safeIndex(817, (d_1081_v45_).length(0))
        (d_1081_v45_)[index154_] = d_1084_v47_
        index155_ = default__.safeIndex(817, (d_1081_v45_).length(0))
        (d_1081_v45_)[index155_] = (d_1084_v47_ if not(p1) else d_1084_v47_)
        d_1085_v48_: _dafny.Array
        def lambda56_(d_1086_i8_):
            return _dafny.CodePoint('j')

        init32_ = lambda56_
        nw169_ = _dafny.Array(None, 23)
        for i0_32_ in range(nw169_.length(0)):
            nw169_[i0_32_] = init32_(i0_32_)
        d_1085_v48_ = nw169_
        d_1087_v49_: str
        d_1087_v49_ = _dafny.CodePoint('g')
        d_1088_v50_: C4
        nw170_ = C4()
        nw170_.ctor__(-768, d_1077_v43_, d_1085_v48_, d_1087_v49_, p0)
        d_1088_v50_ = nw170_
        d_1089_v51_: _dafny.Array
        nw171_ = _dafny.Array(_dafny.Array(None, 0), 6)
        d_1089_v51_ = nw171_
        d_1090_v52_: _dafny.Map
        d_1090_v52_ = _dafny.Map({p1: d_1089_v51_})
        r0 = ((d_1090_v52_)[(len((d_1022_v1_).set(default__.safeIndex(-248, len(d_1022_v1_)), d_1021_v0_))) >= (d_1024_v3_)] if ((len((d_1022_v1_).set(default__.safeIndex(-248, len(d_1022_v1_)), d_1021_v0_))) >= (d_1024_v3_)) in (d_1090_v52_) else d_1089_v51_)
        return r0

    @property
    def f32(self):
        return self._f32

class C6(T1, T0):
    def  __init__(self):
        self._f23: str = _dafny.CodePoint('D')
        self._f24: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f23(self):
        return self._f23
    @f23.setter
    def f23(self, value):
        self._f23 = value
    @property
    def f24(self):
        return self._f24
    def ctor__(self, f23, f24):
        (self).f23 = f23
        (self)._f24 = f24

    def fm5(self, globalState):
        return D0_DC3(D0_DC1(len(_dafny.Map({len(_dafny.Set({self.f23})): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwt"))}))))

    def fm6(self, globalState):
        return len((_dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1091_i0_ in range(default__.abs(217))])})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gt"))})))

    def fm4(self, p0, p1, globalState):
        return (_dafny.MultiSet([_dafny.MultiSet([(self).f24]), (D0_DC0(_dafny.MultiSet([(self).f24, (self).f24, (self).f24, (self).f24, (self).f24]))).cf0])).isdisjoint((_dafny.MultiSet([_dafny.MultiSet([(self).f24])])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, (self).f24])]))))

    def fm16(self, globalState):
        return ((self).f24) not in (_dafny.Map({(self).f24: 298}))

    def m1(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: D0 = D0.default()()
        r2: bool = False
        d_1092_v0_: int
        d_1092_v0_ = 914
        hi7_ = d_1092_v0_
        for d_1093_i0_ in range(d_1092_v0_, hi7_):
            d_1094_v1_: _dafny.Map
            d_1094_v1_ = _dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hblxvj"))})
            d_1095_v2_: C0
            nw172_ = C0()
            nw172_.ctor__(d_1094_v1_, (self).f24, self.f23, (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xu")))) == (d_1093_i0_))
            d_1095_v2_ = nw172_
            d_1096_v3_: _dafny.Seq
            d_1096_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
            (globalState).f7 = len(d_1096_v3_)
            rhs183_ = default__.safeDivisionInt(d_1093_i0_, d_1092_v0_)
            rhs184_ = d_1092_v0_
            lhs138_ = globalState
            lhs138_.f7 = rhs183_
            d_1092_v0_ = rhs184_
            d_1097_v4_: _dafny.Seq
            d_1097_v4_ = _dafny.SeqWithoutIsStrInference([(self).f24, (self).f24])
            if ((not(p1) if (self).f24 else True) if (d_1095_v2_).f38 else (d_1097_v4_)[default__.safeIndex(d_1092_v0_, len(d_1097_v4_))]):
                (globalState).f7 = default__.safeModuloInt((0) - (d_1092_v0_), (d_1092_v0_) - (len(d_1097_v4_)))
                d_1098_v5_: _dafny.Array
                nw173_ = _dafny.Array(False, 23)
                d_1098_v5_ = nw173_
                d_1099_v6_: _dafny.Map
                d_1099_v6_ = _dafny.Map({d_1098_v5_: d_1092_v0_})
                d_1100_v7_: _dafny.Seq
                d_1100_v7_ = _dafny.SeqWithoutIsStrInference([d_1099_v6_, d_1099_v6_])
                d_1101_v8_: _dafny.Map
                d_1101_v8_ = _dafny.Map({(self).f24: len((d_1100_v7_)[default__.safeIndex((0) - (d_1093_i0_), len(d_1100_v7_))])})
                d_1102_v9_: _dafny.Map
                d_1102_v9_ = _dafny.Map({d_1093_i0_: (d_1101_v8_) | (d_1101_v8_)})
                d_1102_v9_ = (d_1102_v9_).set((0) - (d_1093_i0_), _dafny.Map({(self).f24: len(_dafny.SeqWithoutIsStrInference([d_1092_v0_]))}))
                r0 = (default__.fm1((d_1095_v2_).f38, globalState)) + (d_1096_v3_)
                (self).f23 = self.f23
                d_1092_v0_ = d_1092_v0_
            elif True:
                d_1103_v10_: _dafny.Set
                d_1103_v10_ = _dafny.Set({d_1092_v0_, d_1092_v0_})
                d_1104_v11_: _dafny.Seq
                d_1104_v11_ = _dafny.SeqWithoutIsStrInference([d_1103_v10_, d_1103_v10_])
                (globalState).f7 = len(_dafny.SeqWithoutIsStrInference([(len(d_1104_v11_)) + (d_1093_i0_), d_1092_v0_, -974]))
                d_1105_v12_: _dafny.Array
                def lambda57_(d_1106_v0_):
                    def lambda58_(d_1107_i1_):
                        return (d_1107_i1_) - (d_1106_v0_)

                    return lambda58_

                init33_ = lambda57_(d_1092_v0_)
                nw174_ = _dafny.Array(None, 12)
                for i0_33_ in range(nw174_.length(0)):
                    nw174_[i0_33_] = init33_(i0_33_)
                d_1105_v12_ = nw174_
                (globalState).f7 = default__.safeDivisionInt((0) - (default__.safeModuloInt(d_1092_v0_, d_1093_i0_)), default__.safeModuloInt((_dafny.MultiSet([d_1105_v12_])).cardinality, d_1092_v0_))
                (globalState).f7 = d_1093_i0_
                (globalState).f13 = (self).f24
                d_1108_v13_: _dafny.Map
                d_1108_v13_ = _dafny.Map({(self).f24: p1})
                rhs185_ = d_1093_i0_
                rhs186_ = default__.fm19(globalState)
                d_1092_v0_ = rhs185_
                d_1108_v13_ = rhs186_
        d_1109_v14_: _dafny.Array
        def lambda59_(d_1110_v0_):
            def lambda60_(d_1111_i2_):
                return default__.safeModuloInt(d_1111_i2_, d_1110_v0_)

            return lambda60_

        init34_ = lambda59_(d_1092_v0_)
        nw175_ = _dafny.Array(None, 23)
        for i0_34_ in range(nw175_.length(0)):
            nw175_[i0_34_] = init34_(i0_34_)
        d_1109_v14_ = nw175_
        d_1112_v15_: _dafny.Seq
        d_1112_v15_ = _dafny.SeqWithoutIsStrInference([d_1109_v14_])
        d_1113_v16_: _dafny.Seq
        d_1113_v16_ = _dafny.SeqWithoutIsStrInference([(d_1112_v15_)[default__.safeIndex(d_1092_v0_, len(d_1112_v15_))], d_1109_v14_])
        d_1114_v17_: _dafny.Seq
        d_1114_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bniwatogj"))
        d_1115_v18_: _dafny.Map
        d_1115_v18_ = _dafny.Map({p1: d_1114_v17_})
        d_1116_v19_: D2
        d_1116_v19_ = D2_DC6(d_1115_v18_)
        d_1092_v0_ = (len((d_1113_v16_ if p1 else d_1113_v16_))) - (len(default__.fm45((self).f24, d_1116_v19_, d_1092_v0_, globalState)))
        r2 = (self).f24
        d_1117_v20_: C1
        nw176_ = C1()
        nw176_.ctor__(p1, p1)
        d_1117_v20_ = nw176_
        d_1118_v21_: C1
        d_1118_v21_ = d_1117_v20_
        d_1119_v22_: _dafny.Seq
        d_1119_v22_ = _dafny.SeqWithoutIsStrInference([d_1118_v21_, d_1118_v21_, d_1118_v21_, d_1118_v21_, d_1118_v21_])
        d_1120_v23_: _dafny.Seq
        d_1120_v23_ = d_1119_v22_
        d_1121_v24_: _dafny.MultiSet
        d_1121_v24_ = _dafny.MultiSet([d_1119_v22_, _dafny.SeqWithoutIsStrInference([d_1118_v21_, d_1118_v21_, d_1118_v21_]), d_1119_v22_, (d_1120_v23_), d_1119_v22_])
        d_1121_v24_ = ((d_1121_v24_).set(d_1119_v22_, default__.abs(d_1092_v0_))) - (d_1121_v24_)
        d_1122_i3_: int
        d_1122_i3_ = 0
        with _dafny.label("5"):
            while p1:
                with _dafny.c_label("5"):
                    if (d_1122_i3_) >= (100):
                        raise _dafny.Break("5")
                    d_1122_i3_ = (d_1122_i3_) + (1)
                    (globalState).f7 = d_1092_v0_
                    (globalState).f22 = d_1109_v14_
                    d_1123_v25_: D5
                    d_1123_v25_ = D5_DC13((0) - (d_1092_v0_), (d_1117_v20_).f39, (self).f24)
                    d_1124_v26_: _dafny.Map
                    d_1124_v26_ = _dafny.Map({402: (d_1123_v25_).cf21})
                    d_1124_v26_ = (d_1124_v26_).set(431, p1)
                    d_1125_v27_: _dafny.Seq
                    d_1125_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - (d_1092_v0_) for d_1126_i4_ in range(default__.abs(153))])])
                    d_1127_v28_: _dafny.Array
                    nw177_ = _dafny.Array(None, 10)
                    nw177_[int(0)] = p1
                    nw177_[int(1)] = p1
                    nw177_[int(2)] = (self).f24
                    nw177_[int(3)] = p1
                    nw177_[int(4)] = (self).f24
                    nw177_[int(5)] = (self).f24
                    nw177_[int(6)] = (self).f24
                    nw177_[int(7)] = p1
                    nw177_[int(8)] = (self).f24
                    nw177_[int(9)] = (d_1117_v20_).f39
                    d_1127_v28_ = nw177_
                    d_1128_v29_: D1
                    d_1128_v29_ = D1_DC5((d_1092_v0_) + (d_1092_v0_), len(d_1114_v17_), d_1125_v27_, (d_1127_v28_ if True else d_1127_v28_), d_1092_v0_)
                    source16_ = d_1128_v29_
                    if source16_.is_DC5:
                        d_1129___mcc_h0_ = source16_.cf7
                        d_1130___mcc_h1_ = source16_.cf8
                        d_1131___mcc_h2_ = source16_.cf9
                        d_1132___mcc_h3_ = source16_.cf10
                        d_1133___mcc_h4_ = source16_.cf11
                        d_1134_cf11_ = d_1133___mcc_h4_
                        d_1135_cf10_ = d_1132___mcc_h3_
                        d_1136_cf9_ = d_1131___mcc_h2_
                        d_1137_cf8_ = d_1130___mcc_h1_
                        d_1138_cf7_ = d_1129___mcc_h0_
                        index156_ = default__.safeIndex(455, (d_1135_cf10_).length(0))
                        (d_1135_cf10_)[index156_] = (d_1117_v20_).f39
                        d_1139_v30_: _dafny.Map
                        d_1139_v30_ = _dafny.Map({(d_1117_v20_).f39: (d_1117_v20_).f39})
                        d_1140_v31_: _dafny.MultiSet
                        d_1140_v31_ = _dafny.MultiSet([(d_1117_v20_).f39])
                        d_1141_v32_: _dafny.Seq
                        d_1141_v32_ = _dafny.SeqWithoutIsStrInference([(d_1117_v20_).f39, p1])
                        index157_ = default__.safeIndex(455, (d_1135_cf10_).length(0))
                        (d_1135_cf10_)[index157_] = ((d_1139_v30_)[(d_1117_v20_).f39] if ((d_1117_v20_).f39) in (d_1139_v30_) else (_dafny.MultiSet(d_1141_v32_)).ispropersubset(d_1140_v31_))
                        d_1142_v33_: D11
                        d_1142_v33_ = D11_DC22(d_1109_v14_)
                        d_1143_v34_: _dafny.Map
                        d_1143_v34_ = _dafny.Map({(True if p1 else (self).f24): (d_1142_v33_).cf35})
                        d_1144_v35_: D5
                        d_1144_v35_ = D5_DC14((d_1117_v20_).f39)
                        (globalState).f22 = ((d_1143_v34_)[((d_1117_v20_).fm32((self).f24, d_1137_cf8_, globalState)) == ((d_1144_v35_).cf23)] if (((d_1117_v20_).fm32((self).f24, d_1137_cf8_, globalState)) == ((d_1144_v35_).cf23)) in (d_1143_v34_) else d_1109_v14_)
                        index158_ = default__.safeIndex(455, (d_1135_cf10_).length(0))
                        (d_1135_cf10_)[index158_] = not((self).f24)
                        d_1145_v36_: _dafny.Map
                        d_1145_v36_ = _dafny.Map({p0: d_1114_v17_})
                        d_1137_cf8_ = len(d_1145_v36_)
                    elif True:
                        d_1146___mcc_h5_ = source16_.cf6
                        d_1147_cf6_ = d_1146___mcc_h5_
                        d_1148_v37_: bool
                        out36_: bool
                        out36_ = (d_1117_v20_).m4(False, globalState)
                        d_1148_v37_ = out36_
                        d_1092_v0_ = d_1092_v0_
                        d_1149_v38_: _dafny.Seq
                        d_1149_v38_ = _dafny.SeqWithoutIsStrInference([d_1092_v0_, d_1092_v0_, d_1092_v0_, len(d_1114_v17_)])
                        d_1150_v39_: _dafny.Array
                        def lambda61_(d_1151_i5_):
                            return _dafny.CodePoint('s')

                        init35_ = lambda61_
                        nw178_ = _dafny.Array(None, 7)
                        for i0_35_ in range(nw178_.length(0)):
                            nw178_[i0_35_] = init35_(i0_35_)
                        d_1150_v39_ = nw178_
                        d_1152_v40_: D4
                        d_1152_v40_ = D4_DC10(d_1114_v17_)
                        d_1153_v41_: C4
                        nw179_ = C4()
                        nw179_.ctor__(d_1092_v0_, default__.fm46(d_1092_v0_, d_1149_v38_, globalState), d_1150_v39_, default__.fm34(len((d_1152_v40_).cf17), (d_1117_v20_).f39, globalState), p1)
                        d_1153_v41_ = nw179_
                        rhs187_ = 239
                        rhs188_ = (0) - ((d_1153_v41_).f33)
                        lhs139_ = globalState
                        lhs139_.f7 = rhs187_
                        d_1092_v0_ = rhs188_
                    pass
            pass
        d_1154_v42_: _dafny.Seq
        d_1154_v42_ = _dafny.SeqWithoutIsStrInference([(self).f24, False, (p1 if not((self).f24) else (self).f24)])
        if (d_1154_v42_)[default__.safeIndex(350, len(d_1154_v42_))]:
            d_1155_v43_: _dafny.Set
            d_1155_v43_ = _dafny.Set({(d_1117_v20_).f39})
            (globalState).f7 = default__.safeDivisionInt(len((_dafny.SeqWithoutIsStrInference([p0 for d_1156_i6_ in range(default__.abs(-258))])) + (d_1114_v17_)), (0) - (len(d_1155_v43_)))
            d_1120_v23_ = d_1120_v23_
            d_1112_v15_ = d_1113_v16_
            d_1092_v0_ = d_1092_v0_
            d_1157_v44_: _dafny.Array
            nw180_ = _dafny.Array(False, 20)
            d_1157_v44_ = nw180_
            index159_ = default__.safeIndex(245, (d_1157_v44_).length(0))
            (d_1157_v44_)[index159_] = (d_1092_v0_) > (805)
            d_1158_v45_: _dafny.Seq
            d_1158_v45_ = _dafny.SeqWithoutIsStrInference([d_1092_v0_])
            d_1159_v46_: _dafny.Map
            d_1159_v46_ = _dafny.Map({len(d_1154_v42_): (d_1158_v45_)})
            d_1160_v47_: _dafny.Map
            d_1160_v47_ = _dafny.Map({d_1159_v46_: ((self).f24 if True else (self).f24)})
            d_1161_v48_: _dafny.Set
            d_1161_v48_ = _dafny.Set({d_1092_v0_, d_1092_v0_, 50})
            index160_ = default__.safeIndex(245, (d_1157_v44_).length(0))
            rhs189_ = ((d_1160_v47_)[default__.fm47((default__.fm18(globalState)).cardinality, globalState)] if (default__.fm47((default__.fm18(globalState)).cardinality, globalState)) in (d_1160_v47_) else (_dafny.Set({default__.fm0(p1, d_1092_v0_, globalState)})).ispropersubset(d_1161_v48_))
            rhs190_ = _dafny.CodePoint('i')
            rhs191_ = d_1092_v0_
            lhs140_ = d_1157_v44_
            lhs141_ = default__.safeIndex(245, (d_1157_v44_).length(0))
            lhs142_ = self
            lhs143_ = globalState
            lhs140_[lhs141_] = rhs189_
            lhs142_.f23 = rhs190_
            lhs143_.f7 = rhs191_
        elif True:
            d_1092_v0_ = len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1092_v0_]) for d_1162_i7_ in range(default__.abs(744))]))
            rhs192_ = (d_1117_v20_).f39
            rhs193_ = (d_1092_v0_) * ((d_1092_v0_ if (self).f24 else -78))
            rhs194_ = d_1154_v42_
            lhs144_ = globalState
            r2 = rhs192_
            lhs144_.f7 = rhs193_
            d_1154_v42_ = rhs194_
            d_1163_v49_: _dafny.Array
            nw181_ = _dafny.Array(D6.default()(), 18)
            d_1163_v49_ = nw181_
            d_1164_v51_: D6
            def iife164_():
                coll74_ = _dafny.Map()
                compr_74_: int
                for compr_74_ in _dafny.IntegerRange(327, 336):
                    d_1165_v50_: int = compr_74_
                    if ((327) <= (d_1165_v50_)) and ((d_1165_v50_) < (336)):
                        coll74_[default__.safeDivisionInt(d_1165_v50_, d_1092_v0_)] = d_1154_v42_
                return _dafny.Map(coll74_)
            d_1164_v51_ = D6_DC15(iife164_()
)
            index161_ = default__.safeIndex(738, (d_1163_v49_).length(0))
            (d_1163_v49_)[index161_] = d_1164_v51_
            index162_ = default__.safeIndex(738, (d_1163_v49_).length(0))
            (d_1163_v49_)[index162_] = default__.fm39(d_1092_v0_, globalState)
            d_1166_v52_: _dafny.MultiSet
            d_1166_v52_ = _dafny.MultiSet([d_1092_v0_])
            d_1167_v53_: _dafny.Map
            d_1167_v53_ = _dafny.Map({d_1114_v17_: d_1166_v52_})
            d_1168_v54_: D5
            d_1168_v54_ = D5_DC14((d_1117_v20_).f39)
            d_1169_v55_: _dafny.Map
            d_1169_v55_ = _dafny.Map({410: not(p1)})
            d_1170_v56_: _dafny.Map
            d_1170_v56_ = _dafny.Map({not(p1): (d_1117_v20_).f39})
            d_1171_v57_: _dafny.Map
            d_1171_v57_ = _dafny.Map({False: ((d_1170_v56_)[(d_1117_v20_).f39] if ((d_1117_v20_).f39) in (d_1170_v56_) else (d_1117_v20_).f39)})
            d_1172_v58_: _dafny.Set
            d_1172_v58_ = _dafny.Set({d_1166_v52_})
            d_1173_v59_: _dafny.Map
            d_1173_v59_ = _dafny.Map({(d_1117_v20_).f39: d_1172_v58_})
            d_1174_v60_: _dafny.Array
            nw182_ = _dafny.Array(None, 27)
            nw182_[int(0)] = p1
            nw182_[int(1)] = (p0) in (d_1114_v17_)
            nw182_[int(2)] = ((d_1117_v20_).f39 if (self).f24 else (self).f24)
            nw182_[int(3)] = (self).f24
            nw182_[int(4)] = p1
            nw182_[int(5)] = (d_1166_v52_).isdisjoint(d_1166_v52_)
            nw182_[int(6)] = ((_dafny.SeqWithoutIsStrInference([p0 for d_1175_i8_ in range(default__.abs(770))])).set(default__.safeIndex((0) - (d_1092_v0_), len(_dafny.SeqWithoutIsStrInference([p0 for d_1176_i8_ in range(default__.abs(770))]))), p0)) in (d_1167_v53_)
            nw182_[int(7)] = p1
            nw182_[int(8)] = (d_1092_v0_) != (d_1092_v0_)
            nw182_[int(9)] = (d_1168_v54_).cf23
            nw182_[int(10)] = (False) == ((self).f24)
            nw182_[int(11)] = (d_1117_v20_).f39
            nw182_[int(12)] = (self).f24
            nw182_[int(13)] = (self).f24
            nw182_[int(14)] = (self).f24
            nw182_[int(15)] = p1
            nw182_[int(16)] = (d_1117_v20_).f39
            nw182_[int(17)] = (self).f24
            nw182_[int(18)] = not((self).f24)
            nw182_[int(19)] = p1
            nw182_[int(20)] = p1
            nw182_[int(21)] = ((d_1117_v20_).f39 if (self).f24 else ((d_1169_v55_)[d_1092_v0_] if (d_1092_v0_) in (d_1169_v55_) else (d_1117_v20_).f39))
            nw182_[int(22)] = ((d_1171_v57_)[p1] if (p1) in (d_1171_v57_) else not(False))
            nw182_[int(23)] = p1
            nw182_[int(24)] = (d_1117_v20_).f39
            nw182_[int(25)] = (len(((d_1173_v59_)[(d_1117_v20_).f39] if ((d_1117_v20_).f39) in (d_1173_v59_) else d_1172_v58_))) < (d_1092_v0_)
            nw182_[int(26)] = (d_1117_v20_).f39
            d_1174_v60_ = nw182_
            index163_ = default__.safeIndex(696, (d_1174_v60_).length(0))
            (d_1174_v60_)[index163_] = (d_1117_v20_).f39
            index164_ = default__.safeIndex(696, (d_1174_v60_).length(0))
            (d_1174_v60_)[index164_] = p1
            d_1177_v61_: _dafny.Array
            nw183_ = _dafny.Array(_dafny.Array(None, 0), 27)
            d_1177_v61_ = nw183_
            index165_ = default__.safeIndex(593, (d_1177_v61_).length(0))
            (d_1177_v61_)[index165_] = d_1174_v60_
            d_1178_v62_: _dafny.Map
            d_1178_v62_ = _dafny.Map({not((d_1174_v60_)[default__.safeIndex(696, (d_1174_v60_).length(0))]): d_1174_v60_})
            index166_ = default__.safeIndex(593, (d_1177_v61_).length(0))
            (d_1177_v61_)[index166_] = ((d_1178_v62_)[p1] if (p1) in (d_1178_v62_) else d_1174_v60_)
        r0 = (d_1114_v17_) + (d_1114_v17_)
        d_1179_v63_: D0
        d_1179_v63_ = D0_DC1(len(((d_1115_v18_)[p1] if (p1) in (d_1115_v18_) else d_1114_v17_)))
        r1 = (d_1179_v63_ if (self).f24 else d_1179_v63_)
        r2 = not((self).f24)
        return r0, r1, r2

    def m2(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        source17_ = D2_DC6(_dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hin"))}))
        if source17_.is_DC7:
            d_1180___mcc_h0_ = source17_.cf13
            d_1181___mcc_h1_ = source17_.cf14
            d_1182_cf14_ = d_1181___mcc_h1_
            d_1183_cf13_ = d_1180___mcc_h0_
            d_1184_v0_: _dafny.Array
            nw184_ = _dafny.Array(_dafny.CodePoint('D'), 22)
            d_1184_v0_ = nw184_
            d_1184_v0_ = d_1184_v0_
            d_1185_v1_: _dafny.Array
            nw185_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
            d_1185_v1_ = nw185_
            d_1185_v1_ = d_1185_v1_
            d_1186_v2_: int
            d_1186_v2_ = 486
            r2 = (d_1186_v2_) <= ((0) - (d_1186_v2_))
            (globalState).f7 = (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1187_i0_ in range(default__.abs(711))])).set(default__.safeIndex(-356, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1188_i0_ in range(default__.abs(711))]))), self.f23))) * ((0) - (d_1186_v2_))
        elif source17_.is_DC6:
            d_1189___mcc_h2_ = source17_.cf12
            d_1190_cf12_ = d_1189___mcc_h2_
            d_1191_v3_: _dafny.Seq
            d_1191_v3_ = _dafny.SeqWithoutIsStrInference([False])
            d_1192_v4_: int
            d_1192_v4_ = 158
            d_1193_v5_: _dafny.Seq
            d_1193_v5_ = _dafny.SeqWithoutIsStrInference([len(d_1191_v3_), d_1192_v4_, 430])
            d_1194_v6_: _dafny.Map
            d_1194_v6_ = _dafny.Map({(self).f24: d_1193_v5_})
            d_1194_v6_ = (d_1194_v6_) | (d_1194_v6_)
            d_1195_v7_: _dafny.Array
            nw186_ = _dafny.Array(_dafny.Map({}), 20)
            d_1195_v7_ = nw186_
            d_1196_v8_: _dafny.Map
            d_1196_v8_ = _dafny.Map({(self).f24: d_1192_v4_})
            index167_ = default__.safeIndex(985, (d_1195_v7_).length(0))
            (d_1195_v7_)[index167_] = (d_1196_v8_) | (d_1196_v8_)
            index168_ = default__.safeIndex(985, (d_1195_v7_).length(0))
            (d_1195_v7_)[index168_] = d_1196_v8_
            d_1197_v9_: _dafny.Seq
            d_1197_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urkqv"))
            d_1198_v10_: _dafny.Map
            d_1198_v10_ = _dafny.Map({d_1192_v4_: d_1192_v4_})
            d_1199_v12_: C0
            nw187_ = C0()
            def iife165_():
                coll75_ = _dafny.Set()
                compr_75_: int
                for compr_75_ in (d_1198_v10_).keys.Elements:
                    d_1200_v11_: int = compr_75_
                    if (d_1200_v11_) in (d_1198_v10_):
                        coll75_ = coll75_.union(_dafny.Set([(d_1200_v11_) + (476)]))
                return _dafny.Set(coll75_)
            nw187_.ctor__((_dafny.Map({(self).f24: d_1197_v9_})) | (d_1190_cf12_), (len(iife165_()
            )) < (d_1192_v4_), self.f23, True)
            d_1199_v12_ = nw187_
            d_1191_v3_ = d_1191_v3_
        elif True:
            d_1201___mcc_h3_ = source17_.cf15
            d_1202_cf15_ = d_1201___mcc_h3_
            d_1203_v13_: C1
            nw188_ = C1()
            nw188_.ctor__((self).f24, (self).f24)
            d_1203_v13_ = nw188_
            d_1204_v14_: C1
            d_1204_v14_ = d_1203_v13_
            d_1205_v15_: _dafny.Seq
            d_1205_v15_ = _dafny.SeqWithoutIsStrInference([d_1203_v13_, d_1203_v13_, d_1203_v13_])
            d_1206_v16_: int
            d_1206_v16_ = -984
            source18_ = (d_1205_v15_)[default__.safeIndex(d_1206_v16_, len(d_1205_v15_))]
            d_1207___mcc_h4_ = source18_
            d_1208_cf28_ = d_1207___mcc_h4_
            d_1209_v17_: _dafny.Array
            nw189_ = _dafny.Array(_dafny.Array(None, 0), 22)
            d_1209_v17_ = nw189_
            d_1210_v18_: _dafny.Array
            def lambda62_(d_1211_v16_):
                def lambda63_(d_1212_i1_):
                    return (d_1212_i1_) - (len(_dafny.SeqWithoutIsStrInference([d_1211_v16_])))

                return lambda63_

            init36_ = lambda62_(d_1206_v16_)
            nw190_ = _dafny.Array(None, 8)
            for i0_36_ in range(nw190_.length(0)):
                nw190_[i0_36_] = init36_(i0_36_)
            d_1210_v18_ = nw190_
            index169_ = default__.safeIndex(359, (d_1209_v17_).length(0))
            (d_1209_v17_)[index169_] = d_1210_v18_
            index170_ = default__.safeIndex(359, (d_1209_v17_).length(0))
            (d_1209_v17_)[index170_] = d_1210_v18_
            d_1213_v19_: _dafny.Seq
            d_1213_v19_ = _dafny.SeqWithoutIsStrInference([not((d_1208_cf28_).f39), (d_1203_v13_).f39])
            d_1214_v20_: _dafny.Seq
            d_1214_v20_ = _dafny.SeqWithoutIsStrInference([d_1213_v19_])
            d_1215_v21_: int
            d_1216_v22_: bool
            d_1217_v23_: int
            out37_: int
            out38_: bool
            out39_: int
            out37_, out38_, out39_ = default__.m0(d_1206_v16_, _dafny.MultiSet([(self).f24]), ((self).f24) == ((d_1208_cf28_).f39), len(d_1214_v20_), globalState)
            d_1215_v21_ = out37_
            d_1216_v22_ = out38_
            d_1217_v23_ = out39_
            d_1218_v24_: _dafny.Array
            nw191_ = _dafny.Array(False, 5)
            d_1218_v24_ = nw191_
            index171_ = default__.safeIndex(527, (d_1218_v24_).length(0))
            (d_1218_v24_)[index171_] = d_1216_v22_
            index172_ = default__.safeIndex(24, (d_1218_v24_).length(0))
            (d_1218_v24_)[index172_] = d_1216_v22_
            d_1219_v25_: _dafny.Seq
            d_1219_v25_ = _dafny.SeqWithoutIsStrInference([d_1215_v21_])
            index173_ = default__.safeIndex(527, (d_1218_v24_).length(0))
            index174_ = default__.safeIndex(24, (d_1218_v24_).length(0))
            rhs195_ = (self).f24
            rhs196_ = (d_1214_v20_) + (_dafny.SeqWithoutIsStrInference([d_1213_v19_]))
            rhs197_ = ((d_1219_v25_) != (d_1219_v25_) if d_1216_v22_ else True)
            lhs145_ = d_1218_v24_
            lhs146_ = default__.safeIndex(527, (d_1218_v24_).length(0))
            lhs147_ = d_1218_v24_
            lhs148_ = default__.safeIndex(24, (d_1218_v24_).length(0))
            lhs145_[lhs146_] = rhs195_
            d_1214_v20_ = rhs196_
            lhs147_[lhs148_] = rhs197_
            d_1220_v26_: _dafny.Map
            d_1220_v26_ = _dafny.Map({d_1216_v22_: d_1216_v22_})
            d_1221_v27_: _dafny.Map
            d_1221_v27_ = _dafny.Map({d_1213_v19_: (d_1208_cf28_).f39})
            d_1222_v28_: _dafny.Map
            d_1222_v28_ = _dafny.Map({((d_1220_v26_)[(d_1203_v13_).f39] if ((d_1203_v13_).f39) in (d_1220_v26_) else ((d_1221_v27_)[d_1213_v19_] if (d_1213_v19_) in (d_1221_v27_) else not((d_1203_v13_).f39))): d_1215_v21_})
            d_1223_v29_: _dafny.Seq
            d_1223_v29_ = _dafny.SeqWithoutIsStrInference([d_1222_v28_, _dafny.Map({default__.fm2(globalState): (d_1203_v13_).fm33(globalState)})])
            d_1222_v28_ = (d_1223_v29_)[default__.safeIndex(d_1215_v21_, len(d_1223_v29_))]
            d_1224_v30_: _dafny.Seq
            d_1224_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycfwea"))
            d_1225_v31_: _dafny.Map
            d_1225_v31_ = _dafny.Map({not((self).f24): d_1224_v30_})
            d_1226_v32_: D2
            d_1226_v32_ = D2_DC6(d_1225_v31_)
            source19_ = d_1226_v32_
            if source19_.is_DC7:
                d_1227___mcc_h5_ = source19_.cf13
                d_1228___mcc_h6_ = source19_.cf14
                d_1229_cf14_ = d_1228___mcc_h6_
                d_1230_cf13_ = d_1227___mcc_h5_
                d_1231_v33_: _dafny.Seq
                d_1231_v33_ = _dafny.SeqWithoutIsStrInference([d_1230_cf13_])
                d_1232_v34_: _dafny.Map
                d_1232_v34_ = _dafny.Map({default__.fm2(globalState): d_1231_v33_})
                d_1232_v34_ = (d_1232_v34_).set(d_1230_cf13_, d_1231_v33_)
                d_1233_v35_: _dafny.Map
                d_1233_v35_ = _dafny.Map({(d_1203_v13_).f39: _dafny.Map({len(d_1224_v30_): (d_1203_v13_).f39})})
                d_1234_v36_: _dafny.Map
                d_1234_v36_ = _dafny.Map({d_1206_v16_: default__.fm2(globalState)})
                d_1233_v35_ = (d_1233_v35_).set((self).f24, d_1234_v36_)
                d_1235_v37_: _dafny.Map
                d_1235_v37_ = _dafny.Map({len(d_1232_v34_): d_1206_v16_})
                d_1236_v38_: _dafny.Seq
                d_1236_v38_ = _dafny.SeqWithoutIsStrInference([d_1206_v16_])
                d_1237_v39_: _dafny.Set
                d_1237_v39_ = _dafny.Set({(self).f24, (D5_DC13(d_1206_v16_, (d_1203_v13_).f39, (d_1231_v33_)[default__.safeIndex(d_1206_v16_, len(d_1231_v33_))])).cf21})
                d_1238_v40_: _dafny.Seq
                d_1238_v40_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_1203_v13_).f39, (d_1203_v13_).f39}), d_1237_v39_])
                d_1239_v41_: _dafny.Array
                nw192_ = _dafny.Array(None, 23)
                nw192_[int(0)] = d_1206_v16_
                nw192_[int(1)] = d_1206_v16_
                nw192_[int(2)] = default__.safeModuloInt((0) - (d_1206_v16_), d_1206_v16_)
                nw192_[int(3)] = d_1206_v16_
                nw192_[int(4)] = d_1206_v16_
                nw192_[int(5)] = d_1206_v16_
                nw192_[int(6)] = d_1206_v16_
                nw192_[int(7)] = d_1206_v16_
                nw192_[int(8)] = default__.safeModuloInt(d_1206_v16_, ((d_1235_v37_)[len(d_1236_v38_)] if (len(d_1236_v38_)) in (d_1235_v37_) else d_1206_v16_))
                nw192_[int(9)] = d_1206_v16_
                nw192_[int(10)] = (self).fm6(globalState)
                nw192_[int(11)] = (d_1206_v16_) * (len((d_1238_v40_)[default__.safeIndex(d_1206_v16_, len(d_1238_v40_))]))
                nw192_[int(12)] = d_1206_v16_
                nw192_[int(13)] = (0) - (d_1206_v16_)
                nw192_[int(14)] = d_1206_v16_
                nw192_[int(15)] = d_1206_v16_
                nw192_[int(16)] = (d_1236_v38_)[default__.safeIndex(d_1206_v16_, len(d_1236_v38_))]
                nw192_[int(17)] = len(d_1234_v36_)
                nw192_[int(18)] = (0) - ((0) - (d_1206_v16_))
                nw192_[int(19)] = d_1206_v16_
                nw192_[int(20)] = d_1206_v16_
                nw192_[int(21)] = (0) - (default__.safeModuloInt(len(_dafny.Map({d_1230_cf13_: d_1206_v16_})), d_1206_v16_))
                nw192_[int(22)] = (d_1206_v16_) + (len(d_1224_v30_))
                d_1239_v41_ = nw192_
                index175_ = default__.safeIndex(684, (d_1239_v41_).length(0))
                (d_1239_v41_)[index175_] = len(d_1235_v37_)
                d_1240_v42_: _dafny.MultiSet
                d_1240_v42_ = _dafny.MultiSet([self.f23])
                d_1241_v43_: _dafny.Map
                d_1241_v43_ = _dafny.Map({d_1230_cf13_: False})
                d_1242_v44_: _dafny.MultiSet
                d_1242_v44_ = _dafny.MultiSet([d_1241_v43_, d_1241_v43_, d_1241_v43_])
                d_1243_v45_: _dafny.MultiSet
                d_1243_v45_ = _dafny.MultiSet([d_1242_v44_, d_1242_v44_, (_dafny.MultiSet([d_1241_v43_])).set(d_1241_v43_, default__.abs(d_1206_v16_)), d_1242_v44_, d_1242_v44_])
                index176_ = default__.safeIndex(684, (d_1239_v41_).length(0))
                rhs198_ = d_1206_v16_
                rhs199_ = ((d_1240_v42_)[self.f23] if (self.f23) in (d_1240_v42_) else ((d_1243_v45_)[d_1242_v44_] if (d_1242_v44_) in (d_1243_v45_) else d_1206_v16_))
                rhs200_ = ((821) * (default__.fm0((d_1203_v13_).f39, d_1206_v16_, globalState))) * (d_1206_v16_)
                lhs149_ = d_1239_v41_
                lhs150_ = default__.safeIndex(684, (d_1239_v41_).length(0))
                lhs151_ = globalState
                lhs152_ = globalState
                lhs149_[lhs150_] = rhs198_
                lhs151_.f7 = rhs199_
                lhs152_.f7 = rhs200_
                (globalState).f7 = default__.fm0((d_1203_v13_).f39, d_1206_v16_, globalState)
            elif source19_.is_DC6:
                d_1244___mcc_h7_ = source19_.cf12
                d_1245_cf12_ = d_1244___mcc_h7_
                d_1246_v46_: _dafny.MultiSet
                d_1246_v46_ = _dafny.MultiSet([(d_1203_v13_).f39, (self).f24])
                d_1246_v46_ = d_1246_v46_
                d_1247_v48_: _dafny.Map
                def iife166_():
                    coll76_ = _dafny.Map()
                    compr_76_: int
                    for compr_76_ in _dafny.IntegerRange(704, 468):
                        d_1248_v47_: int = compr_76_
                        if ((704) <= (d_1248_v47_)) and ((d_1248_v47_) < (468)):
                            coll76_[(d_1248_v47_) * (d_1206_v16_)] = d_1206_v16_
                    return _dafny.Map(coll76_)
                d_1247_v48_ = _dafny.Map({iife166_()
                : 39})
                d_1249_v49_: _dafny.MultiSet
                d_1249_v49_ = _dafny.MultiSet([d_1206_v16_, d_1206_v16_, d_1206_v16_])
                d_1250_v50_: _dafny.Set
                d_1250_v50_ = _dafny.Set({(d_1203_v13_).f39})
                nw193_ = _dafny.Array(None, 19)
                nw193_[int(0)] = d_1206_v16_
                nw193_[int(1)] = 31
                nw193_[int(2)] = (d_1206_v16_) + (d_1206_v16_)
                nw193_[int(3)] = len(d_1247_v48_)
                nw193_[int(4)] = d_1206_v16_
                nw193_[int(5)] = ((d_1249_v49_)[d_1206_v16_] if (d_1206_v16_) in (d_1249_v49_) else d_1206_v16_)
                nw193_[int(6)] = d_1206_v16_
                nw193_[int(7)] = d_1206_v16_
                nw193_[int(8)] = (d_1206_v16_) - (d_1206_v16_)
                nw193_[int(9)] = d_1206_v16_
                nw193_[int(10)] = default__.safeDivisionInt(d_1206_v16_, d_1206_v16_)
                nw193_[int(11)] = (d_1206_v16_) + (default__.fm0((d_1203_v13_).f39, d_1206_v16_, globalState))
                nw193_[int(12)] = d_1206_v16_
                nw193_[int(13)] = d_1206_v16_
                nw193_[int(14)] = (d_1206_v16_) + ((0) - (d_1206_v16_))
                nw193_[int(15)] = d_1206_v16_
                nw193_[int(16)] = d_1206_v16_
                nw193_[int(17)] = default__.fm0((self).f24, len(d_1250_v50_), globalState)
                nw193_[int(18)] = d_1206_v16_
                (globalState).f22 = nw193_
                (globalState).f7 = default__.safeModuloInt(d_1206_v16_, d_1206_v16_)
                d_1251_v51_: _dafny.Map
                d_1251_v51_ = _dafny.Map({105: d_1206_v16_})
                d_1252_v52_: _dafny.Array
                nw194_ = _dafny.Array(_dafny.CodePoint('D'), 1)
                d_1252_v52_ = nw194_
                d_1253_v53_: D12
                d_1253_v53_ = D12_DC24(d_1252_v52_)
                d_1254_v54_: D0
                d_1254_v54_ = D0_DC1((d_1246_v46_).cardinality)
                d_1255_v55_: C4
                nw195_ = C4()
                nw195_.ctor__((d_1206_v16_) - (d_1206_v16_), d_1251_v51_, (d_1253_v53_).cf40, default__.fm34((d_1254_v54_).cf1, False, globalState), (d_1203_v13_).f39)
                d_1255_v55_ = nw195_
            elif True:
                d_1256___mcc_h8_ = source19_.cf15
                d_1257_cf15_ = d_1256___mcc_h8_
                d_1258_v56_: _dafny.Map
                d_1258_v56_ = _dafny.Map({(d_1203_v13_).f39: (d_1203_v13_).f39})
                d_1258_v56_ = (d_1258_v56_).set((d_1203_v13_).f39, (d_1203_v13_).f39)
                d_1259_v59_: _dafny.Map
                def iife167_():
                    def iife168_():
                        coll78_ = _dafny.Map()
                        compr_78_: int
                        for compr_78_ in _dafny.IntegerRange(-155, 293):
                            d_1261_v58_: int = compr_78_
                            if ((-155) <= (d_1261_v58_)) and ((d_1261_v58_) < (293)):
                                coll78_[default__.safeDivisionInt(d_1261_v58_, d_1206_v16_)] = d_1258_v56_
                        return _dafny.Map(coll78_)
                    coll77_ = _dafny.Map()
                    compr_77_: int
                    for compr_77_ in _dafny.IntegerRange(183, 311):
                        d_1260_v57_: int = compr_77_
                        if ((183) <= (d_1260_v57_)) and ((d_1260_v57_) < (311)):
                            coll77_[(d_1260_v57_) * (len(iife168_()
                            ))] = d_1206_v16_
                    return _dafny.Map(coll77_)
                d_1259_v59_ = _dafny.Map({len(iife167_()
                ): (default__.fm1((self).f24, globalState)).set(default__.safeIndex(d_1206_v16_, len(default__.fm1((self).f24, globalState))), self.f23)})
                d_1224_v30_ = ((d_1259_v59_)[d_1206_v16_] if (d_1206_v16_) in (d_1259_v59_) else d_1224_v30_)
                d_1224_v30_ = d_1224_v30_
                d_1262_v60_: _dafny.Array
                nw196_ = _dafny.Array(False, 23)
                d_1262_v60_ = nw196_
                d_1263_v61_: C5
                nw197_ = C5()
                nw197_.ctor__(d_1262_v60_)
                d_1263_v61_ = nw197_
            (globalState).f13 = (self).f24
            d_1264_v62_: _dafny.MultiSet
            d_1264_v62_ = _dafny.MultiSet([d_1206_v16_, d_1206_v16_])
            d_1265_v63_: _dafny.Map
            d_1265_v63_ = _dafny.Map({(d_1206_v16_) < ((0) - ((d_1264_v62_).cardinality)): d_1204_v14_})
            d_1265_v63_ = (d_1265_v63_).set(((d_1203_v13_).f39) == ((self).f24), d_1204_v14_)
        d_1266_v64_: _dafny.Map
        d_1266_v64_ = _dafny.Map({(self).f24: (self).f24})
        d_1267_v65_: _dafny.Array
        nw198_ = _dafny.Array(None, 24)
        nw198_[int(0)] = (d_1266_v64_) | (d_1266_v64_)
        nw198_[int(1)] = (d_1266_v64_) | (d_1266_v64_)
        nw198_[int(2)] = (default__.fm19(globalState)) | (d_1266_v64_)
        nw198_[int(3)] = _dafny.Map({not((self).f24): (self).f24})
        nw198_[int(4)] = _dafny.Map({(self).fm16(globalState): (self).f24})
        nw198_[int(5)] = d_1266_v64_
        nw198_[int(6)] = d_1266_v64_
        nw198_[int(7)] = _dafny.Map({(self).f24: (self).f24})
        nw198_[int(8)] = d_1266_v64_
        nw198_[int(9)] = d_1266_v64_
        nw198_[int(10)] = d_1266_v64_
        nw198_[int(11)] = (d_1266_v64_) | (d_1266_v64_)
        nw198_[int(12)] = d_1266_v64_
        nw198_[int(13)] = _dafny.Map({(self).f24: (self).f24})
        nw198_[int(14)] = _dafny.Map({(self).f24: False})
        nw198_[int(15)] = (d_1266_v64_) | (d_1266_v64_)
        nw198_[int(16)] = (_dafny.Map({(self).f24: not((self).f24)})).set((self).f24, (self).f24)
        nw198_[int(17)] = d_1266_v64_
        nw198_[int(18)] = d_1266_v64_
        nw198_[int(19)] = _dafny.Map({(self).f24: (self).f24})
        nw198_[int(20)] = _dafny.Map({(self).f24: (self).f24})
        nw198_[int(21)] = (d_1266_v64_) | (_dafny.Map({(self).f24: (self).f24}))
        nw198_[int(22)] = d_1266_v64_
        nw198_[int(23)] = (d_1266_v64_) | (d_1266_v64_)
        d_1267_v65_ = nw198_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1267_v65_).length(0)):
            d_1268_i2_: int = guard_loop_3_
            if (True) and (((0) <= (d_1268_i2_)) and ((d_1268_i2_) < ((d_1267_v65_).length(0)))):
                (d_1267_v65_)[(d_1268_i2_)] = (d_1266_v64_) | (_dafny.Map({(self).f24: (self).f24}))
        d_1269_v66_: int
        d_1269_v66_ = 590
        d_1270_v67_: _dafny.Map
        d_1270_v67_ = _dafny.Map({d_1269_v66_: d_1269_v66_})
        d_1271_v68_: _dafny.Map
        d_1271_v68_ = _dafny.Map({d_1269_v66_: d_1270_v67_})
        d_1272_v69_: _dafny.Seq
        d_1272_v69_ = _dafny.SeqWithoutIsStrInference([len(((d_1271_v68_)[d_1269_v66_] if (d_1269_v66_) in (d_1271_v68_) else d_1270_v67_)), d_1269_v66_, d_1269_v66_])
        d_1272_v69_ = (d_1272_v69_) + (d_1272_v69_)
        r1 = d_1269_v66_
        d_1273_v70_: _dafny.Seq
        d_1273_v70_ = _dafny.SeqWithoutIsStrInference([not(True)])
        d_1274_v71_: _dafny.Array
        nw199_ = _dafny.Array(_dafny.CodePoint('D'), 24)
        d_1274_v71_ = nw199_
        d_1275_v72_: _dafny.MultiSet
        d_1275_v72_ = _dafny.MultiSet([self.f23])
        d_1276_v73_: _dafny.Set
        d_1276_v73_ = _dafny.Set({((d_1275_v72_)[self.f23] if (self.f23) in (d_1275_v72_) else (self).fm6(globalState)), -668})
        d_1277_v74_: C2
        nw200_ = C2()
        nw200_.ctor__(_dafny.MultiSet(d_1273_v70_), (d_1274_v71_ if (self).f24 else d_1274_v71_), self.f23, (len(d_1276_v73_)) < (d_1269_v66_))
        d_1277_v74_ = nw200_
        r0 = True
        r0 = (self).f24
        r1 = d_1269_v66_
        r2 = (self).f24
        return r0, r1, r2


class C7(T2, T3, T0):
    def  __init__(self):
        self._f23: str = _dafny.CodePoint('D')
        self._f24: bool = False
        self._f25: _dafny.Array = _dafny.Array(None, 0)
        self._f27: bool = False
        self._f30: bool = False
        self._f31: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f23(self):
        return self._f23
    @f23.setter
    def f23(self, value):
        self._f23 = value
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f30, f31, f25, f23, f24, f27):
        (self)._f30 = f30
        (self)._f31 = f31
        (self)._f25 = f25
        (self).f23 = f23
        (self)._f24 = f24
        (self)._f27 = f27

    def fm7(self, globalState):
        return ((_dafny.SeqWithoutIsStrInference([564])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([self.f23 for d_1278_i0_ in range(default__.abs(892))]))]))) + ((_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f24]))).cardinality, (0) - (len(_dafny.SeqWithoutIsStrInference([not((self).f30)])))])) + (_dafny.SeqWithoutIsStrInference([802 for d_1279_i1_ in range(default__.abs(933))])))

    def fm8(self, p0, p1, globalState):
        return ((D2_DC6(_dafny.Map({(self).f30: _dafny.SeqWithoutIsStrInference([self.f23 for d_1280_i0_ in range(default__.abs(4))])}))).cf12) | ((_dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))})) | (_dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sufkeaph"))})))

    def fm4(self, p0, p1, globalState):
        return (252) > (len(_dafny.SeqWithoutIsStrInference([not(not((self).f30))])))

    def fm9(self, p0, p1, p2, globalState):
        return default__.safeModuloInt((0) - ((0) - ((-200 if (self).f30 else -98))), default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({891}))])), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ych"))): len((_dafny.Map({-903: -298})))}))))

    def fm15(self, p0, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([(self).f27]) if (self).f31 else _dafny.SeqWithoutIsStrInference([True, (self).f24, (self).f24]))])

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        d_1281_v0_: _dafny.Array
        def lambda64_(d_1282_i0_):
            return (self).f30

        init37_ = lambda64_
        nw201_ = _dafny.Array(None, 11)
        for i0_37_ in range(nw201_.length(0)):
            nw201_[i0_37_] = init37_(i0_37_)
        d_1281_v0_ = nw201_
        d_1283_v1_: int
        d_1283_v1_ = 235
        d_1284_v2_: _dafny.Map
        d_1284_v2_ = _dafny.Map({p0: p0})
        d_1285_v3_: _dafny.Array
        nw202_ = _dafny.Array(None, 10)
        nw202_[int(0)] = d_1283_v1_
        nw202_[int(1)] = 441
        nw202_[int(2)] = -229
        nw202_[int(3)] = d_1283_v1_
        nw202_[int(4)] = d_1283_v1_
        nw202_[int(5)] = d_1283_v1_
        nw202_[int(6)] = d_1283_v1_
        nw202_[int(7)] = 459
        nw202_[int(8)] = 855
        nw202_[int(9)] = len(d_1284_v2_)
        d_1285_v3_ = nw202_
        d_1286_v4_: _dafny.Seq
        d_1286_v4_ = _dafny.SeqWithoutIsStrInference([d_1285_v3_])
        index177_ = default__.safeIndex(567, (d_1281_v0_).length(0))
        (d_1281_v0_)[index177_] = (d_1286_v4_) == (d_1286_v4_)
        d_1287_v5_: _dafny.Map
        d_1287_v5_ = _dafny.Map({p0: p1})
        d_1288_v6_: _dafny.Map
        d_1288_v6_ = _dafny.Map({len(d_1287_v5_): d_1283_v1_})
        d_1289_v7_: _dafny.Map
        d_1289_v7_ = _dafny.Map({p0: d_1288_v6_})
        d_1290_v8_: _dafny.Map
        d_1290_v8_ = ((d_1289_v7_)[(self).f31] if ((self).f31) in (d_1289_v7_) else d_1288_v6_)
        index178_ = default__.safeIndex(567, (d_1281_v0_).length(0))
        def lambda65_(source20_):
            d_1291___mcc_h0_ = source20_
            d_1292_cf16_ = d_1291___mcc_h0_
            return not ((self).f24) or ((self).f30)

        (d_1281_v0_)[index178_] = lambda65_(d_1290_v8_)
        d_1293_v9_: _dafny.Array
        nw203_ = _dafny.Array(_dafny.Array(None, 0), 5)
        d_1293_v9_ = nw203_
        index179_ = default__.safeIndex(97, (d_1293_v9_).length(0))
        (d_1293_v9_)[index179_] = d_1281_v0_
        d_1294_v10_: _dafny.Map
        d_1294_v10_ = _dafny.Map({d_1283_v1_: False})
        d_1295_v11_: _dafny.Map
        d_1295_v11_ = _dafny.Map({d_1294_v10_: d_1281_v0_})
        index180_ = default__.safeIndex(97, (d_1293_v9_).length(0))
        (d_1293_v9_)[index180_] = ((d_1295_v11_)[d_1294_v10_] if (d_1294_v10_) in (d_1295_v11_) else d_1281_v0_)
        d_1296_v12_: _dafny.Array
        nw204_ = _dafny.Array(_dafny.Array(None, 0), 12)
        d_1296_v12_ = nw204_
        d_1297_v13_: _dafny.Array
        nw205_ = _dafny.Array(None, 14)
        nw205_[int(0)] = d_1285_v3_
        nw205_[int(1)] = d_1285_v3_
        nw205_[int(2)] = d_1285_v3_
        nw205_[int(3)] = d_1285_v3_
        nw205_[int(4)] = d_1285_v3_
        nw205_[int(5)] = d_1285_v3_
        nw205_[int(6)] = d_1285_v3_
        nw205_[int(7)] = d_1285_v3_
        nw205_[int(8)] = d_1285_v3_
        nw205_[int(9)] = d_1285_v3_
        nw205_[int(10)] = d_1285_v3_
        nw205_[int(11)] = d_1285_v3_
        nw205_[int(12)] = d_1285_v3_
        nw205_[int(13)] = d_1285_v3_
        d_1297_v13_ = nw205_
        index181_ = default__.safeIndex(476, (d_1296_v12_).length(0))
        (d_1296_v12_)[index181_] = d_1297_v13_
        d_1298_v14_: T1
        nw206_ = C6()
        nw206_.ctor__(self.f23, True)
        d_1298_v14_ = nw206_
        index182_ = default__.safeIndex(476, (d_1296_v12_).length(0))
        index183_ = default__.safeIndex(567, (d_1281_v0_).length(0))
        rhs201_ = ((d_1297_v13_ if (self).f30 else d_1297_v13_) if (self).f30 else d_1297_v13_)
        rhs202_ = (d_1283_v1_) < (72)
        rhs203_ = d_1298_v14_
        lhs153_ = d_1296_v12_
        lhs154_ = default__.safeIndex(476, (d_1296_v12_).length(0))
        lhs155_ = d_1281_v0_
        lhs156_ = default__.safeIndex(567, (d_1281_v0_).length(0))
        lhs153_[lhs154_] = rhs201_
        lhs155_[lhs156_] = rhs202_
        d_1298_v14_ = rhs203_
        d_1299_v15_: _dafny.Set
        d_1299_v15_ = _dafny.Set({d_1283_v1_})
        d_1300_v16_: D5
        d_1300_v16_ = D5_DC12((d_1299_v15_) | (d_1299_v15_))
        source21_ = d_1300_v16_
        if source21_.is_DC13:
            d_1301___mcc_h1_ = source21_.cf20
            d_1302___mcc_h2_ = source21_.cf21
            d_1303___mcc_h3_ = source21_.cf22
            d_1304_cf22_ = d_1303___mcc_h3_
            d_1305_cf21_ = d_1302___mcc_h2_
            d_1306_cf20_ = d_1301___mcc_h1_
            d_1307_v17_: _dafny.MultiSet
            d_1307_v17_ = _dafny.MultiSet([p0])
            r1 = default__.safeModuloInt(d_1283_v1_, (len((_dafny.SeqWithoutIsStrInference([48 for d_1308_i1_ in range(default__.abs(650))])).set(default__.safeIndex(d_1306_cf20_, len(_dafny.SeqWithoutIsStrInference([48 for d_1309_i1_ in range(default__.abs(650))]))), len(_dafny.Set({d_1306_cf20_, (d_1307_v17_).cardinality, d_1306_cf20_}))))) * ((d_1298_v14_).fm6(globalState)))
            if not((self).f24):
                d_1285_v3_ = d_1285_v3_
                d_1310_v18_: C1
                nw207_ = C1()
                nw207_.ctor__(not((d_1305_cf21_ if (d_1298_v14_).f24 else d_1305_cf21_)), p0)
                d_1310_v18_ = nw207_
                d_1311_v19_: D12
                d_1311_v19_ = D12_DC25(((self).f24) and ((d_1298_v14_).f24), d_1298_v14_.f23)
                d_1311_v19_ = d_1311_v19_
                d_1312_v20_: _dafny.Seq
                d_1312_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwgqydm"))
                d_1313_v21_: _dafny.Set
                d_1313_v21_ = _dafny.Set({(d_1298_v14_).f24})
                d_1314_v22_: _dafny.Map
                d_1314_v22_ = _dafny.Map({d_1312_v20_: d_1313_v21_})
                d_1314_v22_ = (d_1314_v22_).set(d_1312_v20_, d_1313_v21_)
                d_1304_cf22_ = (self).f31
            elif True:
                (globalState).f22 = d_1285_v3_
                d_1315_v23_: _dafny.Seq
                d_1315_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlmrjcjpx"))
                d_1316_v24_: D5
                d_1316_v24_ = D5_DC14(default__.fm2(globalState))
                d_1317_v25_: _dafny.Map
                d_1317_v25_ = _dafny.Map({d_1316_v24_: default__.fm0(d_1305_cf21_, d_1306_cf20_, globalState)})
                d_1318_v26_: _dafny.Map
                d_1318_v26_ = _dafny.Map({(d_1298_v14_.f23) in (d_1315_v23_): d_1317_v25_})
                d_1319_v28_: _dafny.Set
                d_1319_v28_ = _dafny.Set({d_1316_v24_})
                def iife169_():
                    coll79_ = _dafny.Map()
                    compr_79_: D5
                    for compr_79_ in (d_1319_v28_).Elements:
                        d_1320_v27_: D5 = compr_79_
                        if (d_1320_v27_) in (d_1319_v28_):
                            coll79_[d_1320_v27_] = d_1306_cf20_
                    return _dafny.Map(coll79_)
                d_1318_v26_ = (d_1318_v26_).set(d_1305_cf21_, iife169_()
                )
                d_1306_cf20_ = d_1306_cf20_
                d_1321_v29_: _dafny.Seq
                d_1321_v29_ = _dafny.SeqWithoutIsStrInference([d_1283_v1_, d_1306_cf20_, d_1306_cf20_, d_1283_v1_])
                index184_ = default__.safeIndex(355, (d_1285_v3_).length(0))
                (d_1285_v3_)[index184_] = len(d_1321_v29_)
                d_1322_v30_: _dafny.Map
                d_1322_v30_ = _dafny.Map({d_1283_v1_: d_1285_v3_})
                d_1323_v31_: _dafny.Map
                d_1323_v31_ = _dafny.Map({d_1284_v2_: d_1305_cf21_})
                index185_ = default__.safeIndex(355, (d_1285_v3_).length(0))
                rhs204_ = (d_1283_v1_) == ((0) - (d_1306_cf20_))
                rhs205_ = (d_1306_cf20_) + (default__.fm0(d_1305_cf21_, d_1283_v1_, globalState))
                rhs206_ = ((d_1322_v30_)[len(d_1323_v31_)] if (len(d_1323_v31_)) in (d_1322_v30_) else (d_1285_v3_ if False else d_1285_v3_))
                rhs207_ = (0) - (d_1283_v1_)
                lhs157_ = globalState
                lhs158_ = globalState
                lhs159_ = globalState
                lhs160_ = d_1285_v3_
                lhs161_ = default__.safeIndex(355, (d_1285_v3_).length(0))
                lhs157_.f13 = rhs204_
                lhs158_.f7 = rhs205_
                lhs159_.f22 = rhs206_
                lhs160_[lhs161_] = rhs207_
                d_1305_cf21_ = (d_1306_cf20_) != ((d_1285_v3_)[default__.safeIndex(355, (d_1285_v3_).length(0))])
            d_1305_cf21_ = (self).f30
            index186_ = default__.safeIndex(567, (d_1281_v0_).length(0))
            (d_1281_v0_)[index186_] = (d_1281_v0_)[default__.safeIndex(567, (d_1281_v0_).length(0))]
        elif source21_.is_DC14:
            d_1324___mcc_h4_ = source21_.cf23
            d_1325_cf23_ = d_1324___mcc_h4_
            index187_ = default__.safeIndex(191, (d_1285_v3_).length(0))
            (d_1285_v3_)[index187_] = default__.safeModuloInt(d_1283_v1_, default__.fm0((self).f24, 262, globalState))
            d_1326_v32_: _dafny.MultiSet
            d_1326_v32_ = _dafny.MultiSet([(self).f24])
            index188_ = default__.safeIndex(191, (d_1285_v3_).length(0))
            (d_1285_v3_)[index188_] = ((d_1326_v32_)[False] if (False) in (d_1326_v32_) else d_1283_v1_)
            d_1327_v33_: _dafny.Seq
            d_1327_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqbli"))
            d_1327_v33_ = (d_1327_v33_) + (d_1327_v33_)
            d_1328_v34_: _dafny.Array
            nw208_ = _dafny.Array(_dafny.Map({}), 26)
            d_1328_v34_ = nw208_
            d_1329_v35_: _dafny.Map
            d_1329_v35_ = _dafny.Map({(d_1285_v3_)[default__.safeIndex(191, (d_1285_v3_).length(0))]: (self).f25})
            index189_ = default__.safeIndex(383, (d_1328_v34_).length(0))
            (d_1328_v34_)[index189_] = d_1329_v35_
            index190_ = default__.safeIndex(383, (d_1328_v34_).length(0))
            (d_1328_v34_)[index190_] = d_1329_v35_
            d_1327_v33_ = (default__.fm1((d_1281_v0_)[default__.safeIndex(567, (d_1281_v0_).length(0))], globalState)).set(default__.safeIndex(default__.safeModuloInt((d_1285_v3_)[default__.safeIndex(191, (d_1285_v3_).length(0))], d_1283_v1_), len(default__.fm1((d_1281_v0_)[default__.safeIndex(567, (d_1281_v0_).length(0))], globalState))), d_1298_v14_.f23)
        elif True:
            d_1330___mcc_h5_ = source21_.cf19
            d_1331_cf19_ = d_1330___mcc_h5_
            d_1332_v36_: _dafny.Seq
            d_1332_v36_ = _dafny.SeqWithoutIsStrInference([(d_1293_v9_)[default__.safeIndex(97, (d_1293_v9_).length(0))], d_1281_v0_, (d_1293_v9_)[default__.safeIndex(97, (d_1293_v9_).length(0))], (d_1293_v9_)[default__.safeIndex(97, (d_1293_v9_).length(0))]])
            d_1332_v36_ = _dafny.SeqWithoutIsStrInference([d_1281_v0_, ((d_1332_v36_)[default__.safeIndex(d_1283_v1_, len(d_1332_v36_))] if (d_1298_v14_).f24 else d_1281_v0_), (d_1293_v9_)[default__.safeIndex(97, (d_1293_v9_).length(0))]])
            index191_ = default__.safeIndex(567, (d_1281_v0_).length(0))
            (d_1281_v0_)[index191_] = (d_1298_v14_).f24
            d_1333_v37_: _dafny.Seq
            d_1333_v37_ = _dafny.SeqWithoutIsStrInference([d_1283_v1_])
            d_1333_v37_ = d_1333_v37_
            d_1334_v38_: _dafny.Seq
            d_1334_v38_ = _dafny.SeqWithoutIsStrInference([(self).f24, not((self).f31), (d_1298_v14_).f24, (self).f24])
            d_1335_v39_: D4
            d_1335_v39_ = D4_DC11(d_1283_v1_)
            rhs208_ = (self).fm4((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dupgk")))) * (d_1283_v1_), d_1334_v38_, globalState)
            rhs209_ = (d_1335_v39_).cf18
            rhs210_ = (d_1334_v38_)[default__.safeIndex(d_1283_v1_, len(d_1334_v38_))]
            rhs211_ = p0
            lhs162_ = globalState
            lhs163_ = globalState
            lhs162_.f13 = rhs208_
            d_1283_v1_ = rhs209_
            lhs163_.f13 = rhs210_
            r2 = rhs211_
        d_1288_v6_ = (d_1288_v6_).set(d_1283_v1_, d_1283_v1_)
        if (((d_1294_v10_)[-678] if (-678) in (d_1294_v10_) else (d_1298_v14_).f24) if (d_1281_v0_)[default__.safeIndex(567, (d_1281_v0_).length(0))] else p0):
            (globalState).f7 = d_1283_v1_
            (self).f23 = (d_1298_v14_.f23 if (self).f31 else d_1298_v14_.f23)
            d_1336_v40_: _dafny.Map
            d_1336_v40_ = _dafny.Map({(d_1298_v14_).f24: d_1283_v1_})
            d_1337_v41_: _dafny.Seq
            d_1337_v41_ = _dafny.SeqWithoutIsStrInference([len(d_1336_v40_)])
            r1 = len((d_1337_v41_) + (d_1337_v41_))
            d_1338_v42_: _dafny.MultiSet
            d_1338_v42_ = _dafny.MultiSet([default__.fm2(globalState)])
            r1 = (((d_1338_v42_) | ((d_1338_v42_).set(p0, default__.abs(d_1283_v1_)))).cardinality) + ((0) - (d_1283_v1_))
            index192_ = default__.safeIndex(159, (d_1285_v3_).length(0))
            (d_1285_v3_)[index192_] = 342
            index193_ = default__.safeIndex(159, (d_1285_v3_).length(0))
            (d_1285_v3_)[index193_] = default__.safeModuloInt(default__.safeDivisionInt(d_1283_v1_, d_1283_v1_), d_1283_v1_)
        elif True:
            d_1339_v43_: D5
            d_1339_v43_ = D5_DC14((d_1298_v14_).f24)
            d_1339_v43_ = d_1339_v43_
            d_1340_v44_: _dafny.Seq
            d_1340_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qy"))
            d_1341_v45_: _dafny.Map
            d_1341_v45_ = _dafny.Map({(self).f30: d_1340_v44_})
            d_1342_v46_: _dafny.Seq
            d_1342_v46_ = _dafny.SeqWithoutIsStrInference([(d_1281_v0_)[default__.safeIndex(567, (d_1281_v0_).length(0))], (self).f30])
            d_1343_v47_: _dafny.Set
            d_1343_v47_ = _dafny.Set({d_1340_v44_, ((d_1341_v45_)[(d_1298_v14_).fm4(d_1283_v1_, d_1342_v46_, globalState)] if ((d_1298_v14_).fm4(d_1283_v1_, d_1342_v46_, globalState)) in (d_1341_v45_) else d_1340_v44_)})
            d_1344_v48_: _dafny.Map
            d_1344_v48_ = _dafny.Map({d_1340_v44_: (self).f24})
            def iife170_():
                coll80_ = _dafny.Set()
                compr_80_: _dafny.Seq
                for compr_80_ in (d_1344_v48_).keys.Elements:
                    d_1345_v49_: _dafny.Seq = compr_80_
                    if (d_1345_v49_) in (d_1344_v48_):
                        coll80_ = coll80_.union(_dafny.Set([d_1345_v49_]))
                return _dafny.Set(coll80_)
            d_1343_v47_ = (iife170_()
            ) - (d_1343_v47_)
            d_1340_v44_ = d_1340_v44_
            d_1346_v50_: _dafny.Set
            d_1346_v50_ = _dafny.Set({(d_1342_v46_)[default__.safeIndex(d_1283_v1_, len(d_1342_v46_))], (self).f30})
            d_1347_v51_: _dafny.Map
            d_1347_v51_ = _dafny.Map({(self).f31: d_1346_v50_})
            d_1347_v51_ = (d_1347_v51_).set(False, d_1346_v50_)
            (globalState).f7 = (d_1283_v1_) + ((0) - (d_1283_v1_))
        d_1348_v52_: _dafny.MultiSet
        d_1348_v52_ = _dafny.MultiSet([False])
        d_1349_v53_: _dafny.Seq
        d_1349_v53_ = _dafny.SeqWithoutIsStrInference([d_1348_v52_])
        d_1350_v54_: _dafny.Seq
        d_1350_v54_ = _dafny.SeqWithoutIsStrInference([d_1283_v1_, d_1283_v1_])
        def iife171_():
            coll81_ = _dafny.Set()
            compr_81_: _dafny.MultiSet
            for compr_81_ in (_dafny.MultiSet(d_1349_v53_)).Elements:
                d_1351_v55_: _dafny.MultiSet = compr_81_
                if (d_1351_v55_) in (_dafny.MultiSet(d_1349_v53_)):
                    coll81_ = coll81_.union(_dafny.Set([d_1351_v55_]))
            return _dafny.Set(coll81_)
        r0 = (_dafny.Set({(d_1349_v53_)[default__.safeIndex(len(d_1350_v54_), len(d_1349_v53_))], _dafny.MultiSet([(d_1298_v14_).f24])})).isdisjoint(iife171_()
        )
        d_1352_v57_: _dafny.Set
        def iife172_():
            coll82_ = _dafny.Set()
            compr_82_: int
            for compr_82_ in _dafny.IntegerRange(-668, -506):
                d_1353_v56_: int = compr_82_
                if ((-668) <= (d_1353_v56_)) and ((d_1353_v56_) < (-506)):
                    coll82_ = coll82_.union(_dafny.Set([(d_1353_v56_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1354_i2_ in range(default__.abs(948))])))]))
            return _dafny.Set(coll82_)
        d_1352_v57_ = _dafny.Set({D5_DC12(iife172_()
), d_1300_v16_})
        d_1355_v58_: _dafny.Set
        d_1355_v58_ = _dafny.Set({d_1352_v57_})
        r1 = len(d_1355_v58_)
        r2 = p0
        return r0, r1, r2

    def m4(self, p0, globalState):
        r0: bool = False
        (globalState).f13 = (self).f24
        d_1356_v0_: _dafny.MultiSet
        d_1356_v0_ = _dafny.MultiSet([(self).f30, (self).f24])
        d_1357_v1_: D0
        d_1357_v1_ = D0_DC0(d_1356_v0_)
        d_1358_v2_: C2
        nw209_ = C2()
        nw209_.ctor__((d_1356_v0_ if (self).f31 else (d_1357_v1_).cf0), (self).f25, self.f23, (self).f24)
        d_1358_v2_ = nw209_
        d_1359_i0_: int
        d_1359_i0_ = 0
        with _dafny.label("6"):
            while p0:
                with _dafny.c_label("6"):
                    if (d_1359_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1359_i0_ = (d_1359_i0_) + (1)
                    d_1360_v3_: _dafny.Array
                    nw210_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
                    d_1360_v3_ = nw210_
                    d_1361_v4_: _dafny.Seq
                    d_1361_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtxxn"))
                    index194_ = default__.safeIndex(47, (d_1360_v3_).length(0))
                    (d_1360_v3_)[index194_] = d_1361_v4_
                    d_1362_v5_: int
                    d_1362_v5_ = 924
                    d_1363_v6_: _dafny.Seq
                    d_1363_v6_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                    index195_ = default__.safeIndex(47, (d_1360_v3_).length(0))
                    (d_1360_v3_)[index195_] = (d_1361_v4_).set(default__.safeIndex((default__.fm0((self).f27, d_1362_v5_, globalState)) * (len(d_1363_v6_)), len(d_1361_v4_)), self.f23)
                    d_1364_v7_: _dafny.Seq
                    d_1364_v7_ = _dafny.SeqWithoutIsStrInference([default__.fm0(True, d_1362_v5_, globalState)])
                    d_1365_v8_: _dafny.Seq
                    d_1365_v8_ = _dafny.SeqWithoutIsStrInference([(d_1364_v7_)[default__.safeIndex(len(d_1364_v7_), len(d_1364_v7_))]])
                    d_1365_v8_ = (d_1365_v8_ if p0 else d_1365_v8_)
                    d_1366_v9_: D5
                    d_1366_v9_ = D5_DC13(d_1362_v5_, True, (self).f30)
                    rhs212_ = d_1362_v5_
                    rhs213_ = D5_DC13(d_1362_v5_, (self).f24, ((self).f31) not in (d_1363_v6_))
                    lhs164_ = globalState
                    lhs164_.f7 = rhs212_
                    d_1366_v9_ = rhs213_
                    d_1367_v10_: _dafny.Array
                    nw211_ = _dafny.Array(False, 24)
                    d_1367_v10_ = nw211_
                    d_1368_v11_: _dafny.Map
                    d_1368_v11_ = _dafny.Map({d_1367_v10_: self.f23})
                    d_1369_v12_: C3
                    nw212_ = C3()
                    nw212_.ctor__((d_1368_v11_) | (_dafny.Map({d_1367_v10_: self.f23})), (self).f25, self.f23, True)
                    d_1369_v12_ = nw212_
                    pass
            pass
        if (self).f31:
            (globalState).f13 = not((self).f31)
            d_1370_v13_: _dafny.Array
            nw213_ = _dafny.Array(int(0), 9)
            d_1370_v13_ = nw213_
            index196_ = default__.safeIndex(360, (d_1370_v13_).length(0))
            (d_1370_v13_)[index196_] = -212
            d_1371_v14_: int
            d_1371_v14_ = 17
            index197_ = default__.safeIndex(360, (d_1370_v13_).length(0))
            (d_1370_v13_)[index197_] = d_1371_v14_
            d_1372_v15_: C1
            nw214_ = C1()
            nw214_.ctor__((self).f27, (self).f30)
            d_1372_v15_ = nw214_
            d_1373_v16_: C1
            d_1373_v16_ = d_1372_v15_
            d_1374_v17_: _dafny.Seq
            d_1374_v17_ = _dafny.SeqWithoutIsStrInference([d_1373_v16_])
            d_1375_v18_: _dafny.Seq
            d_1375_v18_ = (_dafny.SeqWithoutIsStrInference([d_1373_v16_]) if (self).f27 else d_1374_v17_)
            source22_ = d_1375_v18_
            d_1376___mcc_h0_ = source22_
            d_1377_cf34_ = d_1376___mcc_h0_
            (globalState).f13 = (self).f31
            d_1378_v19_: _dafny.Seq
            d_1378_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqrnaakjo"))
            d_1379_v20_: _dafny.Set
            d_1379_v20_ = _dafny.Set({883, d_1371_v14_})
            d_1380_v21_: _dafny.Seq
            d_1380_v21_ = _dafny.SeqWithoutIsStrInference([d_1371_v14_, len(d_1378_v19_), len(d_1379_v20_), (0) - ((0) - (d_1371_v14_))])
            d_1381_v22_: _dafny.Seq
            d_1381_v22_ = _dafny.SeqWithoutIsStrInference([d_1372_v15_])
            d_1382_v23_: _dafny.MultiSet
            d_1382_v23_ = _dafny.MultiSet([d_1372_v15_, d_1372_v15_, d_1372_v15_])
            d_1383_v24_: _dafny.Seq
            d_1383_v24_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1372_v15_, d_1372_v15_]), _dafny.MultiSet(d_1381_v22_), d_1382_v23_, d_1382_v23_])
            d_1384_v25_: _dafny.Map
            d_1384_v25_ = _dafny.Map({d_1371_v14_: (d_1370_v13_)[default__.safeIndex(360, (d_1370_v13_).length(0))]})
            rhs214_ = (self).fm9(d_1380_v21_, (d_1370_v13_)[default__.safeIndex(360, (d_1370_v13_).length(0))], (d_1372_v15_).f39, globalState)
            rhs215_ = (len(d_1383_v24_)) - (d_1371_v14_)
            rhs216_ = ((self).f31) or (((0) - ((d_1370_v13_)[default__.safeIndex(360, (d_1370_v13_).length(0))])) not in (d_1384_v25_))
            lhs165_ = globalState
            d_1371_v14_ = rhs214_
            d_1371_v14_ = rhs215_
            lhs165_.f13 = rhs216_
            d_1385_v26_: _dafny.Map
            d_1385_v26_ = _dafny.Map({self.f23: (d_1370_v13_)[default__.safeIndex(360, (d_1370_v13_).length(0))]})
            (globalState).f7 = (len(d_1385_v26_)) * (default__.safeDivisionInt(d_1371_v14_, d_1371_v14_))
            (globalState).f7 = (d_1370_v13_)[default__.safeIndex(360, (d_1370_v13_).length(0))]
            (globalState).f7 = d_1371_v14_
            d_1386_v27_: _dafny.Map
            d_1386_v27_ = _dafny.Map({((0) - (d_1371_v14_)) * (d_1371_v14_): p0})
            d_1386_v27_ = (d_1386_v27_).set(d_1371_v14_, (d_1372_v15_).f39)
        elif True:
            d_1387_v28_: int
            d_1387_v28_ = 782
            d_1388_v29_: _dafny.Seq
            d_1388_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "byo"))
            if (d_1387_v28_) > (len(d_1388_v29_)):
                (globalState).f13 = not((d_1388_v29_) <= (d_1388_v29_))
                d_1389_v30_: _dafny.Seq
                d_1389_v30_ = _dafny.SeqWithoutIsStrInference([default__.fm2(globalState), (p0 if (self).f30 else p0), (self).f24])
                d_1390_v31_: _dafny.Map
                d_1390_v31_ = _dafny.Map({d_1387_v28_: d_1387_v28_})
                rhs217_ = (0) - (((386 if (self).f30 else d_1387_v28_)) * (default__.fm0(not((self).f27), ((d_1356_v0_)[True] if (True) in (d_1356_v0_) else d_1387_v28_), globalState)))
                rhs218_ = d_1389_v30_
                rhs219_ = (0) - (((d_1390_v31_)[d_1387_v28_] if (d_1387_v28_) in (d_1390_v31_) else (0) - (d_1387_v28_)))
                lhs166_ = globalState
                lhs167_ = globalState
                lhs166_.f7 = rhs217_
                d_1389_v30_ = rhs218_
                lhs167_.f7 = rhs219_
                d_1391_v32_: _dafny.Set
                d_1391_v32_ = _dafny.Set({(self).f24})
                d_1392_v33_: _dafny.Map
                d_1392_v33_ = _dafny.Map({(self).f24: len(d_1391_v32_)})
                d_1393_v34_: _dafny.Set
                d_1393_v34_ = _dafny.Set({d_1392_v33_, d_1392_v33_})
                d_1394_v35_: D0
                d_1394_v35_ = D0_DC3(default__.fm28(d_1388_v29_, _dafny.SeqWithoutIsStrInference([d_1387_v28_]), default__.fm2(globalState), d_1393_v34_, globalState))
                d_1395_v36_: _dafny.Map
                d_1395_v36_ = _dafny.Map({(d_1391_v32_).isdisjoint(d_1391_v32_): d_1394_v35_})
                d_1395_v36_ = (d_1395_v36_).set((self).f31, d_1394_v35_)
                d_1396_v37_: _dafny.Array
                def lambda66_(d_1397_v29_):
                    def lambda67_(d_1398_i1_):
                        return (d_1398_i1_) - ((0) - (len(d_1397_v29_)))

                    return lambda67_

                init38_ = lambda66_(d_1388_v29_)
                nw215_ = _dafny.Array(None, 20)
                for i0_38_ in range(nw215_.length(0)):
                    nw215_[i0_38_] = init38_(i0_38_)
                d_1396_v37_ = nw215_
                d_1399_v38_: _dafny.Set
                d_1399_v38_ = _dafny.Set({d_1387_v28_, d_1387_v28_, d_1387_v28_, d_1387_v28_})
                d_1400_v39_: _dafny.MultiSet
                d_1400_v39_ = _dafny.MultiSet([len(d_1389_v30_)])
                index198_ = default__.safeIndex(26, (d_1396_v37_).length(0))
                def iife173_():
                    coll83_ = _dafny.Set()
                    compr_83_: int
                    for compr_83_ in (d_1400_v39_).Elements:
                        d_1401_v40_: int = compr_83_
                        if (d_1401_v40_) in (d_1400_v39_):
                            coll83_ = coll83_.union(_dafny.Set([default__.safeModuloInt(d_1401_v40_, 321)]))
                    return _dafny.Set(coll83_)
                (d_1396_v37_)[index198_] = len((d_1399_v38_) | (iife173_()
                ))
                index199_ = default__.safeIndex(26, (d_1396_v37_).length(0))
                (d_1396_v37_)[index199_] = d_1387_v28_
                d_1402_v41_: C1
                nw216_ = C1()
                nw216_.ctor__(True, (self).f30)
                d_1402_v41_ = nw216_
                d_1403_v42_: C1
                d_1403_v42_ = d_1402_v41_
                d_1403_v42_ = d_1403_v42_
            elif True:
                d_1404_v43_: C2
                nw217_ = C2()
                nw217_.ctor__(d_1356_v0_, (self).f25, self.f23, False)
                d_1404_v43_ = nw217_
                d_1405_v44_: _dafny.MultiSet
                d_1405_v44_ = _dafny.MultiSet([d_1387_v28_])
                d_1406_v45_: _dafny.Map
                d_1406_v45_ = _dafny.Map({(d_1388_v29_)[default__.safeIndex(d_1387_v28_, len(d_1388_v29_))]: (d_1405_v44_).ispropersubset(d_1405_v44_)})
                d_1407_v46_: _dafny.Array
                def lambda68_(d_1408_i2_):
                    return (self).f27

                init39_ = lambda68_
                nw218_ = _dafny.Array(None, 18)
                for i0_39_ in range(nw218_.length(0)):
                    nw218_[i0_39_] = init39_(i0_39_)
                d_1407_v46_ = nw218_
                d_1409_v47_: _dafny.Map
                d_1409_v47_ = _dafny.Map({(d_1407_v46_ if (self).f24 else d_1407_v46_): d_1406_v45_})
                d_1406_v45_ = ((d_1409_v47_)[d_1407_v46_] if (d_1407_v46_) in (d_1409_v47_) else ((d_1406_v45_).set(self.f23, p0)) | (d_1406_v45_))
                (globalState).f13 = (self).f30
                d_1387_v28_ = d_1387_v28_
                d_1410_v48_: _dafny.Seq
                d_1410_v48_ = _dafny.SeqWithoutIsStrInference([(self).f31, (self).f24, (self).f31, (self).f30, (self).f31])
                d_1411_v49_: _dafny.Map
                d_1411_v49_ = _dafny.Map({d_1387_v28_: d_1410_v48_})
                d_1412_v50_: D6
                d_1412_v50_ = D6_DC15(d_1411_v49_)
                d_1413_v51_: _dafny.Seq
                d_1413_v51_ = _dafny.SeqWithoutIsStrInference([d_1412_v50_])
                d_1414_v52_: _dafny.Map
                d_1414_v52_ = _dafny.Map({(d_1413_v51_) <= (d_1413_v51_): d_1387_v28_})
                (globalState).f7 = len(d_1414_v52_)
            d_1415_v53_: _dafny.Array
            nw219_ = _dafny.Array(False, 25)
            d_1415_v53_ = nw219_
            d_1416_v54_: _dafny.Map
            d_1416_v54_ = _dafny.Map({d_1415_v53_: default__.fm34(123, (self).f24, globalState)})
            d_1417_v55_: C3
            nw220_ = C3()
            nw220_.ctor__(d_1416_v54_, (self).f25, self.f23, (self).f31)
            d_1417_v55_ = nw220_
            d_1418_v56_: _dafny.MultiSet
            d_1418_v56_ = _dafny.MultiSet([self.f23])
            d_1419_v57_: _dafny.Map
            d_1419_v57_ = _dafny.Map({(d_1418_v56_).cardinality: d_1417_v55_})
            d_1420_v58_: _dafny.Seq
            d_1420_v58_ = _dafny.SeqWithoutIsStrInference([d_1417_v55_, d_1417_v55_])
            d_1421_v59_: _dafny.Array
            nw221_ = _dafny.Array(None, 27)
            nw221_[int(0)] = d_1417_v55_
            nw221_[int(1)] = d_1417_v55_
            nw221_[int(2)] = d_1417_v55_
            nw221_[int(3)] = d_1417_v55_
            nw221_[int(4)] = ((d_1419_v57_)[d_1387_v28_] if (d_1387_v28_) in (d_1419_v57_) else d_1417_v55_)
            nw221_[int(5)] = d_1417_v55_
            nw221_[int(6)] = d_1417_v55_
            nw221_[int(7)] = d_1417_v55_
            nw221_[int(8)] = d_1417_v55_
            nw221_[int(9)] = d_1417_v55_
            nw221_[int(10)] = d_1417_v55_
            nw221_[int(11)] = d_1417_v55_
            nw221_[int(12)] = d_1417_v55_
            nw221_[int(13)] = d_1417_v55_
            nw221_[int(14)] = (d_1420_v58_)[default__.safeIndex(d_1387_v28_, len(d_1420_v58_))]
            nw221_[int(15)] = d_1417_v55_
            nw221_[int(16)] = d_1417_v55_
            nw221_[int(17)] = d_1417_v55_
            nw221_[int(18)] = d_1417_v55_
            nw221_[int(19)] = d_1417_v55_
            nw221_[int(20)] = d_1417_v55_
            nw221_[int(21)] = d_1417_v55_
            nw221_[int(22)] = d_1417_v55_
            nw221_[int(23)] = d_1417_v55_
            nw221_[int(24)] = (d_1420_v58_)[default__.safeIndex(d_1387_v28_, len(d_1420_v58_))]
            nw221_[int(25)] = d_1417_v55_
            nw221_[int(26)] = d_1417_v55_
            d_1421_v59_ = nw221_
            index200_ = default__.safeIndex(936, (d_1421_v59_).length(0))
            (d_1421_v59_)[index200_] = d_1417_v55_
            index201_ = default__.safeIndex(936, (d_1421_v59_).length(0))
            (d_1421_v59_)[index201_] = d_1417_v55_
            d_1422_v60_: D2
            d_1422_v60_ = D2_DC7(p0, p0)
            d_1423_v61_: _dafny.Set
            d_1423_v61_ = _dafny.Set({d_1422_v60_, d_1422_v60_, d_1422_v60_, d_1422_v60_})
            pat_let_tv29_ = p0
            def iife174_(_pat_let45_0):
                def iife175_(d_1424_dt__update__tmp_h0_):
                    def iife176_(_pat_let46_0):
                        def iife177_(d_1425_dt__update_hcf14_h0_):
                            return D2_DC7((d_1424_dt__update__tmp_h0_).cf13, d_1425_dt__update_hcf14_h0_)
                        return iife177_(_pat_let46_0)
                    return iife176_(pat_let_tv29_)
                return iife175_(_pat_let45_0)
            d_1423_v61_ = _dafny.Set({iife174_(d_1422_v60_)})
            d_1426_v63_: _dafny.Set
            d_1426_v63_ = _dafny.Set({d_1388_v29_})
            def iife178_():
                coll84_ = _dafny.Map()
                compr_84_: _dafny.Seq
                for compr_84_ in (d_1426_v63_).Elements:
                    d_1427_v62_: _dafny.Seq = compr_84_
                    if (d_1427_v62_) in (d_1426_v63_):
                        coll84_[d_1427_v62_] = (D9_DC20(d_1388_v29_, p0, d_1387_v28_)).cf31
                return _dafny.Map(coll84_)
            if ((D4_DC10(default__.fm1(p0, globalState))).cf17) not in ((iife178_()
            ).set(d_1388_v29_, d_1388_v29_)):
                d_1428_v64_: C3
                nw222_ = C3()
                nw222_.ctor__(_dafny.Map({d_1415_v53_: self.f23}), ((self).f25 if (self).f30 else (self).f25), _dafny.CodePoint('e'), (self).f31)
                d_1428_v64_ = nw222_
                d_1429_v65_: _dafny.Map
                d_1429_v65_ = _dafny.Map({d_1415_v53_: default__.fm0((self).f30, (d_1358_v2_).fm24((0) - (d_1387_v28_), True, globalState), globalState)})
                rhs220_ = (self.f23 if p0 else self.f23)
                rhs221_ = (d_1387_v28_ if (d_1387_v28_) < (d_1387_v28_) else len(d_1429_v65_))
                lhs168_ = self
                lhs169_ = globalState
                lhs168_.f23 = rhs220_
                lhs169_.f7 = rhs221_
                d_1387_v28_ = len(d_1388_v29_)
                d_1430_v66_: D12
                d_1430_v66_ = D12_DC25((self).f31, self.f23)
                d_1430_v66_ = d_1430_v66_
                d_1431_v67_: bool
                d_1432_v68_: int
                d_1433_v69_: bool
                out40_: bool
                out41_: int
                out42_: bool
                out40_, out41_, out42_ = (d_1417_v55_).m3((self).f27, self.f23, globalState)
                d_1431_v67_ = out40_
                d_1432_v68_ = out41_
                d_1433_v69_ = out42_
            elif True:
                (globalState).f7 = d_1387_v28_
                d_1434_v70_: _dafny.Map
                d_1434_v70_ = _dafny.Map({(self).f24: _dafny.CodePoint('q')})
                d_1435_v71_: D13
                d_1435_v71_ = D13_DC27(d_1434_v70_)
                d_1436_v72_: D5
                d_1436_v72_ = D5_DC14((self).f27)
                d_1437_v73_: D11
                d_1437_v73_ = D11_DC23(p0, (0) - (d_1387_v28_), _dafny.Map({(self).f27: p0}), d_1436_v72_)
                d_1438_v74_: _dafny.Seq
                d_1438_v74_ = _dafny.SeqWithoutIsStrInference([(d_1437_v73_).cf36])
                d_1439_v75_: _dafny.Array
                nw223_ = _dafny.Array(None, 8)
                nw223_[int(0)] = (d_1434_v70_).set((self).f24, self.f23)
                nw223_[int(1)] = (d_1435_v71_).cf44
                nw223_[int(2)] = (d_1434_v70_).set((self).f27, self.f23)
                nw223_[int(3)] = _dafny.Map({(d_1438_v74_)[default__.safeIndex(d_1387_v28_, len(d_1438_v74_))]: self.f23})
                nw223_[int(4)] = d_1434_v70_
                nw223_[int(5)] = d_1434_v70_
                nw223_[int(6)] = d_1434_v70_
                nw223_[int(7)] = d_1434_v70_
                d_1439_v75_ = nw223_
                d_1440_v76_: _dafny.Array
                nw224_ = _dafny.Array(int(0), 15)
                d_1440_v76_ = nw224_
                d_1441_v77_: _dafny.Seq
                d_1441_v77_ = _dafny.SeqWithoutIsStrInference([d_1439_v75_, d_1439_v75_, d_1439_v75_, d_1439_v75_, d_1439_v75_])
                d_1442_v78_: _dafny.Map
                d_1442_v78_ = _dafny.Map({d_1440_v76_: (d_1441_v77_)[default__.safeIndex(d_1387_v28_, len(d_1441_v77_))]})
                d_1443_v79_: D14
                d_1443_v79_ = D14_DC30(d_1439_v75_)
                d_1444_v80_: _dafny.Map
                d_1444_v80_ = _dafny.Map({(self).f27: d_1439_v75_})
                d_1445_v81_: _dafny.Array
                nw225_ = _dafny.Array(None, 25)
                nw225_[int(0)] = d_1439_v75_
                nw225_[int(1)] = d_1439_v75_
                nw225_[int(2)] = d_1439_v75_
                nw225_[int(3)] = ((d_1442_v78_)[d_1440_v76_] if (d_1440_v76_) in (d_1442_v78_) else d_1439_v75_)
                nw225_[int(4)] = d_1439_v75_
                nw225_[int(5)] = d_1439_v75_
                nw225_[int(6)] = d_1439_v75_
                nw225_[int(7)] = d_1439_v75_
                nw225_[int(8)] = d_1439_v75_
                nw225_[int(9)] = d_1439_v75_
                nw225_[int(10)] = d_1439_v75_
                nw225_[int(11)] = d_1439_v75_
                nw225_[int(12)] = d_1439_v75_
                nw225_[int(13)] = d_1439_v75_
                nw225_[int(14)] = d_1439_v75_
                nw225_[int(15)] = d_1439_v75_
                nw225_[int(16)] = (d_1443_v79_).cf51
                nw225_[int(17)] = d_1439_v75_
                nw225_[int(18)] = d_1439_v75_
                nw225_[int(19)] = ((d_1444_v80_)[(self).f24] if ((self).f24) in (d_1444_v80_) else d_1439_v75_)
                nw225_[int(20)] = d_1439_v75_
                nw225_[int(21)] = d_1439_v75_
                nw225_[int(22)] = d_1439_v75_
                nw225_[int(23)] = d_1439_v75_
                nw225_[int(24)] = d_1439_v75_
                d_1445_v81_ = nw225_
                index202_ = default__.safeIndex(656, (d_1445_v81_).length(0))
                (d_1445_v81_)[index202_] = d_1439_v75_
                index203_ = default__.safeIndex(656, (d_1445_v81_).length(0))
                (d_1445_v81_)[index203_] = d_1439_v75_
                (self).f23 = self.f23
                d_1446_v82_: _dafny.Seq
                d_1446_v82_ = _dafny.SeqWithoutIsStrInference([d_1387_v28_])
                d_1447_v83_: _dafny.MultiSet
                d_1447_v83_ = _dafny.MultiSet([d_1387_v28_, d_1387_v28_, d_1387_v28_])
                (globalState).f13 = (d_1447_v83_).ispropersubset(_dafny.MultiSet(d_1446_v82_))
                index204_ = default__.safeIndex(919, (d_1415_v53_).length(0))
                (d_1415_v53_)[index204_] = ((self).f24) and ((self).f27)
                index205_ = default__.safeIndex(919, (d_1415_v53_).length(0))
                (d_1415_v53_)[index205_] = p0
            d_1448_v84_: C5
            nw226_ = C5()
            nw226_.ctor__(d_1415_v53_)
            d_1448_v84_ = nw226_
        d_1449_v85_: _dafny.Array
        def lambda69_(d_1450_p0_):
            def lambda70_(d_1451_i3_):
                return d_1450_p0_

            return lambda70_

        init40_ = lambda69_(p0)
        nw227_ = _dafny.Array(None, 2)
        for i0_40_ in range(nw227_.length(0)):
            nw227_[i0_40_] = init40_(i0_40_)
        d_1449_v85_ = nw227_
        d_1452_v86_: _dafny.Map
        d_1452_v86_ = _dafny.Map({d_1449_v85_: 550})
        d_1453_v87_: int
        d_1453_v87_ = 629
        d_1454_v88_: _dafny.MultiSet
        d_1454_v88_ = _dafny.MultiSet([(0) - (d_1453_v87_), d_1453_v87_, d_1453_v87_, d_1453_v87_])
        d_1455_v89_: _dafny.Map
        d_1455_v89_ = _dafny.Map({not((self).f31): (d_1358_v2_).fm4(d_1453_v87_, _dafny.SeqWithoutIsStrInference([(self).f24, (self).f31, (self).f30]), globalState)})
        rhs222_ = d_1452_v86_
        rhs223_ = ((d_1454_v88_) - (default__.fm18(globalState))) - (d_1454_v88_)
        rhs224_ = (d_1453_v87_) + (len(d_1455_v89_))
        lhs170_ = globalState
        d_1452_v86_ = rhs222_
        d_1454_v88_ = rhs223_
        lhs170_.f7 = rhs224_
        d_1456_v90_: _dafny.Array
        nw228_ = _dafny.Array(int(0), 8)
        d_1456_v90_ = nw228_
        index206_ = default__.safeIndex(317, (d_1456_v90_).length(0))
        (d_1456_v90_)[index206_] = d_1453_v87_
        d_1457_v91_: _dafny.Seq
        d_1457_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oaniuldr"))
        d_1458_v92_: _dafny.Map
        d_1458_v92_ = _dafny.Map({(self).f30: (d_1457_v91_).set(default__.safeIndex(d_1453_v87_, len(d_1457_v91_)), self.f23)})
        index207_ = default__.safeIndex(317, (d_1456_v90_).length(0))
        rhs225_ = (d_1458_v92_) == ((d_1458_v92_) | (d_1458_v92_))
        rhs226_ = d_1453_v87_
        lhs171_ = globalState
        lhs172_ = d_1456_v90_
        lhs173_ = default__.safeIndex(317, (d_1456_v90_).length(0))
        lhs171_.f13 = rhs225_
        lhs172_[lhs173_] = rhs226_
        r0 = ((self).f27) and ((self).f24)
        return r0

    def m7(self, p0, p1, globalState):
        d_1459_v0_: int
        d_1460_v1_: bool
        d_1461_v2_: int
        out43_: int
        out44_: bool
        out45_: int
        out43_, out44_, out45_ = default__.m0((p0) + (p0), default__.fm3((self).f30, globalState), True, (p0) + (p0), globalState)
        d_1459_v0_ = out43_
        d_1460_v1_ = out44_
        d_1461_v2_ = out45_
        (globalState).f7 = (p0) - (p0)
        d_1462_v3_: _dafny.Map
        d_1462_v3_ = _dafny.Map({(self).f24: p0})
        d_1463_v4_: _dafny.Seq
        d_1463_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dfrewfj"))
        d_1464_v5_: _dafny.Set
        d_1464_v5_ = _dafny.Set({d_1461_v2_})
        d_1465_v6_: D5
        d_1465_v6_ = D5_DC12(d_1464_v5_)
        pat_let_tv30_ = p1
        pat_let_tv31_ = d_1459_v0_
        def lambda71_(source23_):
            if source23_.is_DC13:
                d_1466___mcc_h0_ = source23_.cf20
                d_1467___mcc_h1_ = source23_.cf21
                d_1468___mcc_h2_ = source23_.cf22
                d_1469_cf22_ = d_1468___mcc_h2_
                d_1470_cf21_ = d_1467___mcc_h1_
                d_1471_cf20_ = d_1466___mcc_h0_
                return not (pat_let_tv30_) or (d_1470_cf21_)
            elif source23_.is_DC14:
                d_1472___mcc_h3_ = source23_.cf23
                d_1473_cf23_ = d_1472___mcc_h3_
                return d_1473_cf23_
            elif True:
                d_1474___mcc_h4_ = source23_.cf19
                d_1475_cf19_ = d_1474___mcc_h4_
                return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wpj")))) < (pat_let_tv31_)

        rhs227_ = d_1462_v3_
        rhs228_ = p0
        rhs229_ = (self).f24
        rhs230_ = not(lambda71_(d_1465_v6_))
        rhs231_ = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1476_i0_ in range(default__.abs(744))])).set(default__.safeIndex(d_1461_v2_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1477_i0_ in range(default__.abs(744))]))), self.f23) if d_1460_v1_ else (d_1463_v4_) + (d_1463_v4_))
        d_1462_v3_ = rhs227_
        d_1461_v2_ = rhs228_
        d_1460_v1_ = rhs229_
        d_1460_v1_ = rhs230_
        d_1463_v4_ = rhs231_
        hi8_ = (p0) + (d_1461_v2_)
        for d_1478_i1_ in range(p0, hi8_):
            d_1479_v7_: _dafny.Map
            d_1479_v7_ = _dafny.Map({(d_1459_v0_ if p1 else p0): not((self).f27)})
            d_1480_v8_: _dafny.Seq
            d_1480_v8_ = _dafny.SeqWithoutIsStrInference([p1, (self).f27])
            d_1481_v9_: _dafny.Seq
            d_1481_v9_ = _dafny.SeqWithoutIsStrInference([False, (self).fm4(p0, _dafny.SeqWithoutIsStrInference([(self).f30]), globalState), (d_1480_v8_)[default__.safeIndex(d_1461_v2_, len(d_1480_v8_))], (self).f30, (self).f27])
            d_1479_v7_ = (d_1479_v7_).set(default__.safeDivisionInt(d_1461_v2_, d_1461_v2_), (len(d_1481_v9_)) >= (d_1459_v0_))
            d_1460_v1_ = p1
            (globalState).f13 = p1
            (globalState).f13 = (self).fm4(d_1459_v0_, d_1480_v8_, globalState)
        d_1482_v10_: _dafny.Map
        d_1482_v10_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([p0, d_1459_v0_])})
        d_1483_v11_: D4
        d_1483_v11_ = D4_DC10(d_1463_v4_)
        d_1484_v12_: _dafny.Map
        d_1484_v12_ = _dafny.Map({d_1482_v10_: d_1483_v11_})
        d_1485_v14_: T0
        nw229_ = C0()
        nw229_.ctor__(_dafny.Map({(self).f31: d_1463_v4_}), (self).f31, self.f23, (self).f31)
        d_1485_v14_ = nw229_
        d_1486_v15_: _dafny.Map
        d_1486_v15_ = _dafny.Map({d_1460_v1_: d_1485_v14_})
        d_1487_v16_: _dafny.Map
        d_1487_v16_ = _dafny.Map({d_1461_v2_: d_1486_v15_})
        d_1488_v17_: _dafny.Seq
        d_1488_v17_ = _dafny.SeqWithoutIsStrInference([d_1459_v0_, len(d_1487_v16_), p0, d_1459_v0_])
        def iife179_():
            coll85_ = _dafny.Map()
            compr_85_: int
            for compr_85_ in _dafny.IntegerRange(-590, 391):
                d_1489_v13_: int = compr_85_
                if ((-590) <= (d_1489_v13_)) and ((d_1489_v13_) < (391)):
                    coll85_[(d_1489_v13_) + (202)] = _dafny.SeqWithoutIsStrInference([p0])
            return _dafny.Map(coll85_)
        d_1484_v12_ = (d_1484_v12_).set((iife179_()
         if not(not((self).f31)) else _dafny.Map({len(d_1463_v4_): d_1488_v17_})), D4_DC10(d_1463_v4_))
        hi9_ = default__.safeDivisionInt((D0_DC2(not(d_1460_v1_), d_1459_v0_, d_1460_v1_)).cf3, d_1461_v2_)
        for d_1490_i2_ in range(743, hi9_):
            if not((d_1464_v5_).ispropersubset(d_1464_v5_)):
                rhs232_ = (((d_1485_v14_).f24 if False else (d_1485_v14_).f24) if (self).f27 else (self).f31)
                rhs233_ = d_1463_v4_
                lhs174_ = globalState
                lhs174_.f13 = rhs232_
                d_1463_v4_ = rhs233_
                d_1491_v18_: _dafny.Seq
                d_1491_v18_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                d_1492_v19_: D14
                d_1492_v19_ = D14_DC31((d_1485_v14_).f24)
                d_1491_v18_ = (d_1491_v18_) + (_dafny.SeqWithoutIsStrInference([(d_1492_v19_).cf52]))
                d_1493_v20_: _dafny.Seq
                d_1493_v20_ = _dafny.SeqWithoutIsStrInference([default__.fm36(globalState)])
                d_1494_v21_: _dafny.MultiSet
                d_1494_v21_ = _dafny.MultiSet([len(d_1493_v20_), d_1461_v2_, d_1490_i2_])
                d_1495_v22_: _dafny.Map
                d_1495_v22_ = _dafny.Map({(d_1494_v21_).isdisjoint(d_1494_v21_): d_1464_v5_})
                d_1461_v2_ = len(d_1495_v22_)
                d_1496_v23_: _dafny.Map
                d_1496_v23_ = _dafny.Map({(d_1490_i2_) >= (d_1461_v2_): not ((self).f27) or (not(d_1460_v1_))})
                d_1496_v23_ = (d_1496_v23_).set(True, ((0) - (d_1459_v0_)) != (p0))
                d_1491_v18_ = d_1491_v18_
            elif True:
                d_1497_v24_: _dafny.MultiSet
                d_1497_v24_ = _dafny.MultiSet([(d_1485_v14_).f24])
                d_1498_v25_: _dafny.Map
                d_1498_v25_ = _dafny.Map({d_1490_i2_: d_1461_v2_})
                d_1499_v26_: _dafny.Array
                nw230_ = _dafny.Array(None, 17)
                nw230_[int(0)] = len(d_1463_v4_)
                nw230_[int(1)] = p0
                nw230_[int(2)] = ((0) - (d_1490_i2_)) - (862)
                nw230_[int(3)] = d_1461_v2_
                nw230_[int(4)] = (788) * ((d_1497_v24_).cardinality)
                nw230_[int(5)] = ((d_1497_v24_)[(self).f24] if ((self).f24) in (d_1497_v24_) else p0)
                nw230_[int(6)] = -929
                nw230_[int(7)] = p0
                nw230_[int(8)] = d_1490_i2_
                nw230_[int(9)] = -59
                nw230_[int(10)] = (len(_dafny.SeqWithoutIsStrInference([self.f23 for d_1500_i3_ in range(default__.abs(851))]))) + (d_1459_v0_)
                nw230_[int(11)] = d_1461_v2_
                nw230_[int(12)] = (p0) - ((0) - (d_1490_i2_))
                nw230_[int(13)] = p0
                nw230_[int(14)] = (0) - ((0) - (d_1461_v2_))
                nw230_[int(15)] = p0
                nw230_[int(16)] = ((d_1498_v25_)[d_1461_v2_] if (d_1461_v2_) in (d_1498_v25_) else d_1459_v0_)
                d_1499_v26_ = nw230_
                index208_ = default__.safeIndex(816, (d_1499_v26_).length(0))
                (d_1499_v26_)[index208_] = d_1459_v0_
                index209_ = default__.safeIndex(816, (d_1499_v26_).length(0))
                (d_1499_v26_)[index209_] = ((d_1497_v24_)[not((d_1485_v14_).f24)] if (not((d_1485_v14_).f24)) in (d_1497_v24_) else (0) - ((d_1488_v17_)[default__.safeIndex(355, len(d_1488_v17_))]))
                d_1501_v27_: _dafny.Seq
                d_1501_v27_ = _dafny.SeqWithoutIsStrInference([(d_1463_v4_) + (d_1463_v4_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ifjf")), d_1463_v4_])
                d_1501_v27_ = ((d_1501_v27_) + (d_1501_v27_)) + ((d_1501_v27_) + (d_1501_v27_))
                d_1502_v28_: D11
                d_1502_v28_ = D11_DC22(d_1499_v26_)
                (globalState).f22 = (d_1502_v28_).cf35
                d_1460_v1_ = (self).f24
                d_1503_v29_: _dafny.Map
                d_1503_v29_ = _dafny.Map({False: d_1463_v4_})
                d_1504_v30_: _dafny.Seq
                d_1504_v30_ = _dafny.SeqWithoutIsStrInference([d_1460_v1_, (self).f27, (self).f27])
                d_1505_v31_: C0
                nw231_ = C0()
                nw231_.ctor__(((d_1503_v29_).set((self).f30, _dafny.SeqWithoutIsStrInference([d_1485_v14_.f23 for d_1506_i4_ in range(default__.abs(728))]))) | (_dafny.Map({p1: d_1463_v4_})), (d_1459_v0_) >= (len(d_1504_v30_)), d_1485_v14_.f23, (self).f24)
                d_1505_v31_ = nw231_
            d_1460_v1_ = (d_1485_v14_).f24
            d_1507_v32_: D12
            d_1507_v32_ = D12_DC25(not(p1), d_1485_v14_.f23)
            d_1508_v33_: D12
            d_1508_v33_ = D12_DC26(d_1507_v32_)
            d_1508_v33_ = d_1508_v33_
            (globalState).f13 = (self).f31

    @property
    def f30(self):
        return self._f30
    @property
    def f31(self):
        return self._f31

class C8(T3):
    def  __init__(self):
        self._f27: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f27):
        (self)._f27 = f27

    def fm9(self, p0, p1, p2, globalState):
        return -598

    def m4(self, p0, globalState):
        r0: bool = False
        if (self).f27:
            d_1509_v0_: _dafny.MultiSet
            d_1509_v0_ = _dafny.MultiSet([False])
            d_1510_v1_: _dafny.Set
            d_1510_v1_ = _dafny.Set({D0_DC0(d_1509_v0_)})
            (globalState).f13 = (d_1510_v1_).isdisjoint(d_1510_v1_)
            d_1511_v2_: _dafny.Seq
            d_1511_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktgrmbaw"))
            d_1512_v3_: _dafny.Set
            d_1512_v3_ = _dafny.Set({len(d_1511_v2_)})
            d_1513_v4_: _dafny.Seq
            d_1513_v4_ = _dafny.SeqWithoutIsStrInference([d_1512_v3_])
            d_1514_v5_: int
            d_1514_v5_ = 57
            (globalState).f13 = (d_1512_v3_) != (((d_1513_v4_)[default__.safeIndex(d_1514_v5_, len(d_1513_v4_))]) | (_dafny.Set({639, d_1514_v5_})))
            d_1515_v6_: str
            d_1515_v6_ = _dafny.CodePoint('y')
            d_1515_v6_ = d_1515_v6_
            d_1516_v7_: _dafny.Seq
            d_1516_v7_ = _dafny.SeqWithoutIsStrInference([d_1514_v5_, d_1514_v5_, d_1514_v5_])
            d_1517_v8_: _dafny.Seq
            d_1517_v8_ = _dafny.SeqWithoutIsStrInference([d_1516_v7_])
            d_1518_v9_: _dafny.Array
            nw232_ = _dafny.Array(int(0), 8)
            d_1518_v9_ = nw232_
            d_1519_v10_: _dafny.Map
            d_1519_v10_ = _dafny.Map({d_1514_v5_: d_1518_v9_})
            d_1520_v11_: _dafny.Set
            d_1520_v11_ = _dafny.Set({((d_1519_v10_)[d_1514_v5_] if (d_1514_v5_) in (d_1519_v10_) else d_1518_v9_), d_1518_v9_, d_1518_v9_, d_1518_v9_})
            d_1521_v12_: _dafny.Seq
            d_1521_v12_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])
            d_1522_v13_: _dafny.Seq
            d_1522_v13_ = _dafny.SeqWithoutIsStrInference([(d_1521_v12_).set(default__.safeIndex(len(d_1511_v2_), len(d_1521_v12_)), (self).f27)])
            rhs234_ = (_dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_1514_v5_, len(d_1520_v11_), d_1514_v5_])).set(default__.safeIndex(d_1514_v5_, len(_dafny.SeqWithoutIsStrInference([d_1514_v5_, len(d_1520_v11_), d_1514_v5_]))), len(d_1511_v2_)), d_1516_v7_])) + (d_1517_v8_)
            rhs235_ = d_1514_v5_
            rhs236_ = ((d_1522_v13_ if p0 else d_1522_v13_)) < (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, not(not((self).f27))])]))
            rhs237_ = d_1516_v7_
            lhs175_ = globalState
            lhs176_ = globalState
            d_1517_v8_ = rhs234_
            lhs175_.f7 = rhs235_
            lhs176_.f13 = rhs236_
            d_1516_v7_ = rhs237_
            d_1523_v14_: _dafny.Array
            nw233_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
            d_1523_v14_ = nw233_
            d_1524_v15_: D1
            d_1524_v15_ = D1_DC4(d_1515_v6_)
            rhs238_ = (d_1524_v15_).cf6
            rhs239_ = d_1523_v14_
            rhs240_ = (d_1514_v5_) + ((d_1514_v5_) + (default__.fm0(p0, d_1514_v5_, globalState)))
            lhs177_ = globalState
            d_1515_v6_ = rhs238_
            d_1523_v14_ = rhs239_
            lhs177_.f7 = rhs240_
        elif True:
            d_1525_v16_: int
            d_1525_v16_ = 253
            d_1526_v17_: _dafny.Map
            d_1526_v17_ = _dafny.Map({default__.fm0((self).f27, 570, globalState): not((D0_DC2((self).f27, d_1525_v16_, (self).f27)).cf4)})
            if ((d_1526_v17_)[166] if (166) in (d_1526_v17_) else p0):
                d_1527_v18_: _dafny.Array
                def lambda72_(d_1528_v16_):
                    def lambda73_(d_1529_i0_):
                        return default__.safeDivisionInt(d_1529_i0_, d_1528_v16_)

                    return lambda73_

                init41_ = lambda72_(d_1525_v16_)
                nw234_ = _dafny.Array(None, 24)
                for i0_41_ in range(nw234_.length(0)):
                    nw234_[i0_41_] = init41_(i0_41_)
                d_1527_v18_ = nw234_
                d_1530_v19_: _dafny.Set
                d_1530_v19_ = _dafny.Set({409, d_1525_v16_, d_1525_v16_})
                rhs241_ = d_1527_v18_
                rhs242_ = d_1525_v16_
                rhs243_ = (d_1530_v19_).isdisjoint(d_1530_v19_)
                lhs178_ = globalState
                lhs179_ = globalState
                lhs180_ = globalState
                lhs178_.f22 = rhs241_
                lhs179_.f7 = rhs242_
                lhs180_.f13 = rhs243_
                d_1531_v20_: _dafny.Array
                def lambda74_(d_1532_p0_):
                    def lambda75_(d_1533_i1_):
                        return d_1532_p0_

                    return lambda75_

                init42_ = lambda74_(p0)
                nw235_ = _dafny.Array(None, 29)
                for i0_42_ in range(nw235_.length(0)):
                    nw235_[i0_42_] = init42_(i0_42_)
                d_1531_v20_ = nw235_
                d_1534_v21_: _dafny.Array
                def lambda76_(d_1535_v16_):
                    def lambda77_(d_1536_i2_):
                        return D0_DC1(d_1535_v16_)

                    return lambda77_

                init43_ = lambda76_(d_1525_v16_)
                nw236_ = _dafny.Array(None, 7)
                for i0_43_ in range(nw236_.length(0)):
                    nw236_[i0_43_] = init43_(i0_43_)
                d_1534_v21_ = nw236_
                d_1537_v22_: _dafny.Set
                d_1537_v22_ = _dafny.Set({d_1534_v21_})
                rhs244_ = (0) - (len(((d_1537_v22_) | (d_1537_v22_)) | (d_1537_v22_)))
                rhs245_ = d_1531_v20_
                rhs246_ = (d_1525_v16_) + (d_1525_v16_)
                lhs181_ = globalState
                d_1525_v16_ = rhs244_
                d_1531_v20_ = rhs245_
                lhs181_.f7 = rhs246_
                d_1538_v23_: str
                d_1538_v23_ = _dafny.CodePoint('o')
                d_1539_v24_: _dafny.Array
                nw237_ = _dafny.Array(None, 7)
                nw237_[int(0)] = d_1538_v23_
                nw237_[int(1)] = d_1538_v23_
                nw237_[int(2)] = d_1538_v23_
                nw237_[int(3)] = d_1538_v23_
                nw237_[int(4)] = d_1538_v23_
                nw237_[int(5)] = d_1538_v23_
                nw237_[int(6)] = d_1538_v23_
                d_1539_v24_ = nw237_
                index210_ = default__.safeIndex(410, (d_1539_v24_).length(0))
                (d_1539_v24_)[index210_] = d_1538_v23_
                index211_ = default__.safeIndex(410, (d_1539_v24_).length(0))
                (d_1539_v24_)[index211_] = d_1538_v23_
                (globalState).f7 = default__.fm0((self).f27, d_1525_v16_, globalState)
                d_1540_v25_: D0
                d_1540_v25_ = D0_DC2(p0, d_1525_v16_, p0)
                d_1541_v26_: _dafny.Set
                d_1541_v26_ = _dafny.Set({(p0 if (self).f27 else not(p0)), (d_1540_v25_).cf2})
                d_1541_v26_ = (d_1541_v26_) - ((d_1541_v26_) | (d_1541_v26_))
            elif True:
                d_1542_v27_: _dafny.Seq
                d_1542_v27_ = _dafny.SeqWithoutIsStrInference([d_1525_v16_])
                d_1543_v28_: _dafny.Seq
                d_1543_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwpgw"))
                d_1544_v29_: _dafny.Array
                nw238_ = _dafny.Array(None, 10)
                nw238_[int(0)] = d_1525_v16_
                nw238_[int(1)] = d_1525_v16_
                nw238_[int(2)] = (137 if (self).f27 else d_1525_v16_)
                nw238_[int(3)] = d_1525_v16_
                nw238_[int(4)] = d_1525_v16_
                nw238_[int(5)] = d_1525_v16_
                nw238_[int(6)] = default__.safeModuloInt((0) - (len(d_1542_v27_)), d_1525_v16_)
                nw238_[int(7)] = d_1525_v16_
                nw238_[int(8)] = default__.fm0(p0, d_1525_v16_, globalState)
                nw238_[int(9)] = len((d_1543_v28_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1545_i3_ in range(default__.abs(2))])))
                d_1544_v29_ = nw238_
                (globalState).f22 = d_1544_v29_
                d_1546_v30_: _dafny.Seq
                d_1546_v30_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                d_1547_v31_: D0
                d_1547_v31_ = D0_DC1(d_1525_v16_)
                d_1548_v32_: _dafny.Map
                d_1548_v32_ = _dafny.Map({(p0) not in (d_1546_v30_): d_1547_v31_})
                d_1549_v33_: _dafny.Seq
                d_1549_v33_ = _dafny.SeqWithoutIsStrInference([d_1542_v27_, _dafny.SeqWithoutIsStrInference([d_1525_v16_])])
                d_1550_v34_: str
                d_1550_v34_ = _dafny.CodePoint('c')
                d_1551_v35_: _dafny.Array
                nw239_ = _dafny.Array(None, 19)
                nw239_[int(0)] = d_1550_v34_
                nw239_[int(1)] = d_1550_v34_
                nw239_[int(2)] = d_1550_v34_
                nw239_[int(3)] = d_1550_v34_
                nw239_[int(4)] = d_1550_v34_
                nw239_[int(5)] = d_1550_v34_
                nw239_[int(6)] = d_1550_v34_
                nw239_[int(7)] = d_1550_v34_
                nw239_[int(8)] = d_1550_v34_
                nw239_[int(9)] = d_1550_v34_
                nw239_[int(10)] = d_1550_v34_
                nw239_[int(11)] = d_1550_v34_
                nw239_[int(12)] = d_1550_v34_
                nw239_[int(13)] = d_1550_v34_
                nw239_[int(14)] = d_1550_v34_
                nw239_[int(15)] = d_1550_v34_
                nw239_[int(16)] = d_1550_v34_
                nw239_[int(17)] = d_1550_v34_
                nw239_[int(18)] = d_1550_v34_
                d_1551_v35_ = nw239_
                d_1552_v36_: D12
                d_1552_v36_ = D12_DC24(d_1551_v35_)
                d_1553_v37_: T3
                nw240_ = C7()
                nw240_.ctor__((self).f27, not(True), (d_1552_v36_).cf40, d_1550_v34_, (self).f27, p0)
                d_1553_v37_ = nw240_
                d_1554_v38_: _dafny.Map
                d_1554_v38_ = _dafny.Map({d_1553_v37_: (self).f27})
                rhs247_ = default__.fm13(default__.fm14((self).f27, globalState), not((d_1549_v33_) < ((d_1549_v33_).set(default__.safeIndex(975, len(d_1549_v33_)), _dafny.SeqWithoutIsStrInference([d_1525_v16_])))), d_1525_v16_, ((d_1554_v38_)[d_1553_v37_] if (d_1553_v37_) in (d_1554_v38_) else (d_1553_v37_).f27), globalState)
                rhs248_ = default__.safeDivisionInt((d_1542_v27_)[default__.safeIndex(d_1525_v16_, len(d_1542_v27_))], d_1525_v16_)
                lhs182_ = globalState
                d_1548_v32_ = rhs247_
                lhs182_.f7 = rhs248_
                d_1543_v28_ = (d_1543_v28_) + (d_1543_v28_)
                nw241_ = _dafny.Array(int(0), 21)
                d_1544_v29_ = nw241_
                d_1555_v39_: _dafny.Array
                nw242_ = _dafny.Array(_dafny.Map({}), 25)
                d_1555_v39_ = nw242_
                d_1556_v40_: _dafny.Map
                d_1556_v40_ = _dafny.Map({len(d_1546_v30_): D14_DC30(d_1555_v39_)})
                d_1557_v41_: _dafny.Seq
                d_1557_v41_ = _dafny.SeqWithoutIsStrInference([d_1556_v40_, d_1556_v40_, d_1556_v40_])
                (globalState).f7 = (self).fm9(d_1542_v27_, len(d_1557_v41_), (_dafny.CodePoint('f')) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))), globalState)
            d_1558_v42_: _dafny.Set
            d_1558_v42_ = _dafny.Set({-926})
            d_1559_v43_: _dafny.Map
            d_1559_v43_ = _dafny.Map({d_1558_v42_: d_1558_v42_})
            d_1560_v44_: _dafny.Map
            d_1560_v44_ = _dafny.Map({((d_1559_v43_)[d_1558_v42_] if (d_1558_v42_) in (d_1559_v43_) else d_1558_v42_): d_1525_v16_})
            d_1561_v45_: _dafny.MultiSet
            d_1561_v45_ = _dafny.MultiSet([d_1525_v16_])
            d_1525_v16_ = default__.safeModuloInt(len((d_1560_v44_).set(d_1558_v42_, 101)), ((d_1561_v45_).cardinality if p0 else d_1525_v16_))
            (globalState).f7 = (363) - (default__.safeDivisionInt(d_1525_v16_, d_1525_v16_))
            d_1562_v46_: _dafny.Map
            d_1562_v46_ = _dafny.Map({(self).f27: (self).f27})
            source24_ = D11_DC23((self).f27, (d_1525_v16_) * (d_1525_v16_), d_1562_v46_, D5_DC14(p0))
            if source24_.is_DC23:
                d_1563___mcc_h0_ = source24_.cf36
                d_1564___mcc_h1_ = source24_.cf37
                d_1565___mcc_h2_ = source24_.cf38
                d_1566___mcc_h3_ = source24_.cf39
                d_1567_cf39_ = d_1566___mcc_h3_
                d_1568_cf38_ = d_1565___mcc_h2_
                d_1569_cf37_ = d_1564___mcc_h1_
                d_1570_cf36_ = d_1563___mcc_h0_
                d_1569_cf37_ = d_1525_v16_
                d_1571_v47_: _dafny.Array
                def lambda78_(d_1572_cf37_):
                    def lambda79_(d_1573_i4_):
                        return (d_1573_i4_) + (d_1572_cf37_)

                    return lambda79_

                init44_ = lambda78_(d_1569_cf37_)
                nw243_ = _dafny.Array(None, 23)
                for i0_44_ in range(nw243_.length(0)):
                    nw243_[i0_44_] = init44_(i0_44_)
                d_1571_v47_ = nw243_
                d_1574_v48_: _dafny.MultiSet
                d_1574_v48_ = _dafny.MultiSet([d_1571_v47_])
                d_1575_v49_: str
                d_1575_v49_ = _dafny.CodePoint('e')
                d_1576_v50_: _dafny.Seq
                d_1576_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
                d_1577_v51_: _dafny.Seq
                d_1577_v51_ = _dafny.SeqWithoutIsStrInference([True, (self).f27])
                d_1578_v52_: _dafny.Map
                d_1578_v52_ = _dafny.Map({d_1569_cf37_: d_1577_v51_})
                d_1579_v53_: _dafny.Map
                d_1579_v53_ = _dafny.Map({(d_1575_v49_) in (d_1576_v50_): len(d_1578_v52_)})
                rhs249_ = d_1574_v48_
                rhs250_ = d_1579_v53_
                d_1574_v48_ = rhs249_
                d_1579_v53_ = rhs250_
                d_1580_v54_: _dafny.Array
                def lambda80_(d_1581_p0_):
                    def lambda81_(d_1582_i5_):
                        return d_1581_p0_

                    return lambda81_

                init45_ = lambda80_(p0)
                nw244_ = _dafny.Array(None, 15)
                for i0_45_ in range(nw244_.length(0)):
                    nw244_[i0_45_] = init45_(i0_45_)
                d_1580_v54_ = nw244_
                d_1583_v55_: _dafny.Set
                d_1583_v55_ = _dafny.Set({d_1580_v54_})
                (globalState).f10 = d_1583_v55_
                index212_ = default__.safeIndex(986, (d_1571_v47_).length(0))
                (d_1571_v47_)[index212_] = (59) * (default__.fm0(p0, (0) - (d_1569_cf37_), globalState))
                index213_ = default__.safeIndex(833, (d_1571_v47_).length(0))
                (d_1571_v47_)[index213_] = d_1525_v16_
                d_1584_v56_: _dafny.Seq
                d_1584_v56_ = _dafny.SeqWithoutIsStrInference([d_1525_v16_, d_1525_v16_, d_1525_v16_])
                d_1585_v57_: _dafny.Map
                d_1585_v57_ = _dafny.Map({d_1569_cf37_: d_1569_cf37_})
                d_1586_v58_: _dafny.Map
                d_1586_v58_ = _dafny.Map({default__.safeModuloInt(len(d_1584_v56_), len(d_1577_v51_)): d_1585_v57_})
                d_1587_v59_: D9
                d_1587_v59_ = D9_DC20(d_1576_v50_, (self).f27, d_1569_cf37_)
                d_1588_v60_: _dafny.Map
                d_1588_v60_ = _dafny.Map({d_1587_v59_: default__.fm0(d_1570_cf36_, d_1569_cf37_, globalState)})
                d_1589_v61_: _dafny.Map
                d_1589_v61_ = _dafny.Map({d_1588_v60_: d_1576_v50_})
                d_1590_v63_: _dafny.MultiSet
                d_1590_v63_ = _dafny.MultiSet([D9_DC20(d_1576_v50_, p0, d_1525_v16_)])
                d_1591_v65_: D15
                def iife180_():
                    def iife182_():
                        coll88_ = _dafny.Set()
                        compr_88_: D9
                        for compr_88_ in (d_1590_v63_).Elements:
                            d_1594_v64_: D9 = compr_88_
                            if (d_1594_v64_) in (d_1590_v63_):
                                coll88_ = coll88_.union(_dafny.Set([d_1594_v64_]))
                        return _dafny.Set(coll88_)
                    coll86_ = _dafny.Map()
                    def iife181_():
                        coll87_ = _dafny.Set()
                        compr_87_: D9
                        for compr_87_ in (d_1590_v63_).Elements:
                            d_1592_v64_: D9 = compr_87_
                            if (d_1592_v64_) in (d_1590_v63_):
                                coll87_ = coll87_.union(_dafny.Set([d_1592_v64_]))
                        return _dafny.Set(coll87_)
                    compr_86_: D9
                    for compr_86_ in (iife181_()
                    ).Elements:
                        d_1593_v62_: D9 = compr_86_
                        if (d_1593_v62_) in (iife182_()
                        ):
                            coll86_[d_1593_v62_] = d_1525_v16_
                    return _dafny.Map(coll86_)
                d_1591_v65_ = D15_DC33(iife180_()
)
                index214_ = default__.safeIndex(986, (d_1571_v47_).length(0))
                index215_ = default__.safeIndex(833, (d_1571_v47_).length(0))
                rhs251_ = len(d_1586_v58_)
                rhs252_ = (d_1569_cf37_) - ((0) - (default__.fm0((self).f27, d_1525_v16_, globalState)))
                rhs253_ = ((d_1589_v61_)[(d_1588_v60_) | ((d_1591_v65_).cf54)] if ((d_1588_v60_) | ((d_1591_v65_).cf54)) in (d_1589_v61_) else d_1576_v50_)
                rhs254_ = ((default__.fm36(globalState)).set(default__.safeIndex(d_1525_v16_, len(default__.fm36(globalState))), (self).f27) if not(d_1570_cf36_) else (d_1577_v51_) + (d_1577_v51_))
                lhs183_ = d_1571_v47_
                lhs184_ = default__.safeIndex(986, (d_1571_v47_).length(0))
                lhs185_ = d_1571_v47_
                lhs186_ = default__.safeIndex(833, (d_1571_v47_).length(0))
                lhs183_[lhs184_] = rhs251_
                lhs185_[lhs186_] = rhs252_
                d_1576_v50_ = rhs253_
                d_1577_v51_ = rhs254_
            elif True:
                d_1595___mcc_h4_ = source24_.cf35
                d_1596_cf35_ = d_1595___mcc_h4_
                d_1597_v66_: _dafny.Array
                nw245_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_1597_v66_ = nw245_
                d_1598_v67_: _dafny.Array
                nw246_ = _dafny.Array(False, 16)
                d_1598_v67_ = nw246_
                index216_ = default__.safeIndex(749, (d_1597_v66_).length(0))
                (d_1597_v66_)[index216_] = d_1598_v67_
                d_1599_v68_: _dafny.Array
                nw247_ = _dafny.Array(None, 3)
                d_1599_v68_ = nw247_
                d_1600_v69_: _dafny.Map
                d_1600_v69_ = _dafny.Map({d_1599_v68_: d_1598_v67_})
                index217_ = default__.safeIndex(749, (d_1597_v66_).length(0))
                (d_1597_v66_)[index217_] = ((d_1600_v69_)[d_1599_v68_] if (d_1599_v68_) in (d_1600_v69_) else d_1598_v67_)
                d_1601_v70_: str
                d_1601_v70_ = _dafny.CodePoint('q')
                d_1602_v71_: _dafny.Map
                d_1602_v71_ = _dafny.Map({(d_1525_v16_) - (d_1525_v16_): d_1601_v70_})
                d_1602_v71_ = (d_1602_v71_).set(d_1525_v16_, d_1601_v70_)
                d_1603_v72_: _dafny.Array
                nw248_ = _dafny.Array(D12.default()(), 25)
                d_1603_v72_ = nw248_
                d_1604_v73_: _dafny.Map
                d_1604_v73_ = _dafny.Map({d_1525_v16_: d_1603_v72_})
                d_1605_v74_: _dafny.Seq
                d_1605_v74_ = _dafny.SeqWithoutIsStrInference([d_1525_v16_, d_1525_v16_, len(_dafny.SeqWithoutIsStrInference([d_1601_v70_ for d_1606_i6_ in range(default__.abs(949))]))])
                d_1604_v73_ = (d_1604_v73_).set(len(d_1605_v74_), d_1603_v72_)
                d_1607_v75_: _dafny.Map
                d_1607_v75_ = _dafny.Map({d_1525_v16_: d_1561_v45_})
                d_1607_v75_ = _dafny.Map({d_1525_v16_: d_1561_v45_})
            (globalState).f13 = not (default__.fm2(globalState)) or (not(p0))
        d_1608_v76_: int
        d_1608_v76_ = 549
        hi10_ = d_1608_v76_
        for d_1609_i7_ in range(d_1608_v76_, hi10_):
            (globalState).f13 = (p0) == (False)
            d_1610_v77_: _dafny.Array
            nw249_ = _dafny.Array(_dafny.CodePoint('D'), 14)
            d_1610_v77_ = nw249_
            d_1611_v78_: str
            d_1611_v78_ = _dafny.CodePoint('d')
            d_1612_v79_: _dafny.Seq
            d_1612_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjjnna"))
            d_1613_v80_: D9
            d_1613_v80_ = D9_DC20(d_1612_v79_, p0, len(_dafny.Set({d_1609_i7_})))
            d_1614_v81_: _dafny.Seq
            d_1614_v81_ = _dafny.SeqWithoutIsStrInference([d_1609_i7_, (d_1613_v80_).cf33])
            d_1615_v82_: _dafny.Seq
            d_1615_v82_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(0) - (len(d_1612_v79_)): (self).f27})), 720, len(d_1614_v81_)])
            d_1616_v83_: C7
            nw250_ = C7()
            nw250_.ctor__(True, not(p0), d_1610_v77_, d_1611_v78_, not(not((_dafny.SeqWithoutIsStrInference([d_1609_i7_ for d_1617_i8_ in range(default__.abs(432))])) != (d_1614_v81_))), not (default__.fm2(globalState)) or (p0))
            d_1616_v83_ = nw250_
            d_1618_v84_: _dafny.Map
            d_1618_v84_ = _dafny.Map({p0: d_1614_v81_})
            d_1619_v85_: _dafny.MultiSet
            d_1619_v85_ = _dafny.MultiSet([True])
            d_1620_v86_: int
            d_1621_v87_: bool
            d_1622_v88_: int
            out46_: int
            out47_: bool
            out48_: int
            out46_, out47_, out48_ = default__.m0(len((d_1618_v84_ if (d_1616_v83_).f31 else d_1618_v84_)), d_1619_v85_, (d_1616_v83_).f30, d_1609_i7_, globalState)
            d_1620_v86_ = out46_
            d_1621_v87_ = out47_
            d_1622_v88_ = out48_
            d_1612_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
        d_1623_v89_: _dafny.Seq
        d_1623_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgc"))
        d_1623_v89_ = d_1623_v89_
        d_1608_v76_ = len(((d_1623_v89_) + (d_1623_v89_)) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1624_i9_ in range(default__.abs(272))])))
        r0 = default__.fm2(globalState)
        d_1608_v76_ = default__.safeDivisionInt(d_1608_v76_, (d_1608_v76_) + (d_1608_v76_))
        r0 = (D5_DC13((0) - (len(d_1623_v89_)), p0, p0)).cf22
        return r0


class C9(T1, T0):
    def  __init__(self):
        self._f23: str = _dafny.CodePoint('D')
        self._f24: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f23(self):
        return self._f23
    @f23.setter
    def f23(self, value):
        self._f23 = value
    @property
    def f24(self):
        return self._f24
    def ctor__(self, f23, f24):
        (self).f23 = f23
        (self)._f24 = f24

    def fm5(self, globalState):
        return D0_DC3(D0_DC3(D0_DC2((self).f24, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f24, not((self).f24)])): len(_dafny.Map({D0_DC3(D0_DC0(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f24])))): (self).f24}))})), (self).f24)))

    def fm6(self, globalState):
        return len(((_dafny.Set({D0_DC1(219), D0_DC1(367)})) - (_dafny.Set({D0_DC1(len(_dafny.SeqWithoutIsStrInference([self.f23, _dafny.CodePoint('u')]))), D0_DC1(715), D0_DC1((_dafny.MultiSet([len(_dafny.Map({_dafny.Map({(self).f24: 380}): (self).f24}))])).cardinality)}))).intersection((_dafny.Set({D0_DC1(890)})) | (_dafny.Set({D0_DC1(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obaunck")))), D0_DC1(-720), D0_DC1(len(_dafny.SeqWithoutIsStrInference([(self).f24]))), D0_DC1(len(_dafny.SeqWithoutIsStrInference([83, 946])))}))))

    def fm4(self, p0, p1, globalState):
        def iife183_():
            coll89_ = _dafny.Map()
            compr_89_: int
            for compr_89_ in _dafny.IntegerRange(456, 762):
                d_1625_v0_: int = compr_89_
                if ((456) <= (d_1625_v0_)) and ((d_1625_v0_) < (762)):
                    coll89_[(d_1625_v0_) + (36)] = False
            return _dafny.Map(coll89_)
        return ((0) - (len((iife183_()
        ) | (_dafny.Map({613: (self).f24}))))) > (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([(self).f24])), len(_dafny.SeqWithoutIsStrInference([D0_DC3(D0_DC1(987)) for d_1626_i0_ in range(default__.abs(760))]))))

    def fm11(self, p0, p1, globalState):
        return (self).f24

    def m1(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: D0 = D0.default()()
        r2: bool = False
        d_1627_v0_: int
        d_1627_v0_ = 135
        d_1628_v1_: D0
        d_1628_v1_ = D0_DC2(p1, d_1627_v0_, False)
        if (d_1628_v1_).cf2:
            d_1629_v2_: _dafny.Seq
            d_1629_v2_ = _dafny.SeqWithoutIsStrInference([357])
            d_1630_v3_: _dafny.Map
            d_1630_v3_ = _dafny.Map({default__.fm1(p1, globalState): len(d_1629_v2_)})
            d_1631_v4_: _dafny.Seq
            d_1631_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rptgus"))
            d_1630_v3_ = (d_1630_v3_).set(d_1631_v4_, d_1627_v0_)
            d_1632_v5_: _dafny.Array
            nw251_ = _dafny.Array(False, 5)
            d_1632_v5_ = nw251_
            d_1633_v6_: _dafny.Array
            nw252_ = _dafny.Array(int(0), 16)
            d_1633_v6_ = nw252_
            index218_ = default__.safeIndex(607, (d_1633_v6_).length(0))
            (d_1633_v6_)[index218_] = d_1627_v0_
            d_1634_v7_: _dafny.Array
            nw253_ = _dafny.Array(_dafny.Map({}), 7)
            d_1634_v7_ = nw253_
            index219_ = default__.safeIndex(607, (d_1633_v6_).length(0))
            rhs255_ = (d_1632_v5_ if p1 else d_1632_v5_)
            rhs256_ = d_1627_v0_
            rhs257_ = d_1634_v7_
            lhs187_ = d_1633_v6_
            lhs188_ = default__.safeIndex(607, (d_1633_v6_).length(0))
            d_1632_v5_ = rhs255_
            lhs187_[lhs188_] = rhs256_
            d_1634_v7_ = rhs257_
            d_1627_v0_ = (default__.fm12(self.f23, p1, d_1627_v0_, globalState)).cf1
            d_1631_v4_ = d_1631_v4_
            index220_ = default__.safeIndex(698, (d_1632_v5_).length(0))
            (d_1632_v5_)[index220_] = (self).f24
            d_1635_v8_: _dafny.Set
            d_1635_v8_ = _dafny.Set({(d_1633_v6_)[default__.safeIndex(607, (d_1633_v6_).length(0))]})
            index221_ = default__.safeIndex(698, (d_1632_v5_).length(0))
            (d_1632_v5_)[index221_] = (d_1635_v8_).isdisjoint(d_1635_v8_)
        elif True:
            d_1636_v9_: _dafny.MultiSet
            d_1636_v9_ = _dafny.MultiSet([p1])
            d_1637_v10_: _dafny.Map
            d_1637_v10_ = _dafny.Map({(p1) or ((self).f24): ((self).f24) in (d_1636_v9_)})
            d_1638_v11_: _dafny.Array
            nw254_ = _dafny.Array(_dafny.CodePoint('D'), 21)
            d_1638_v11_ = nw254_
            d_1639_v12_: T3
            nw255_ = C7()
            nw255_.ctor__(p1, p1, d_1638_v11_, _dafny.CodePoint('j'), p1, True)
            d_1639_v12_ = nw255_
            d_1640_v13_: _dafny.Map
            d_1640_v13_ = _dafny.Map({d_1639_v12_: self.f23})
            d_1641_v14_: _dafny.Map
            d_1641_v14_ = _dafny.Map({d_1627_v0_: d_1640_v13_})
            d_1637_v10_ = (d_1637_v10_).set((_dafny.Map({d_1627_v0_: d_1640_v13_})) == (d_1641_v14_), (self).f24)
            (globalState).f7 = d_1627_v0_
            if p1:
                d_1642_v15_: _dafny.Seq
                d_1642_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hapquyumu"))
                d_1643_v16_: _dafny.Map
                d_1643_v16_ = _dafny.Map({(self).f24: d_1642_v15_})
                d_1644_v17_: T0
                nw256_ = C0()
                nw256_.ctor__((d_1643_v16_) | (_dafny.Map({p1: d_1642_v15_})), p1, self.f23, p1)
                d_1644_v17_ = nw256_
                d_1644_v17_ = d_1644_v17_
                pat_let_tv32_ = d_1639_v12_
                def iife184_(_pat_let47_0):
                    def iife185_(d_1645_dt__update__tmp_h0_):
                        def iife186_(_pat_let48_0):
                            def iife187_(d_1646_dt__update_hcf2_h0_):
                                return D0_DC2(d_1646_dt__update_hcf2_h0_, (d_1645_dt__update__tmp_h0_).cf3, (d_1645_dt__update__tmp_h0_).cf4)
                            return iife187_(_pat_let48_0)
                        return iife186_((pat_let_tv32_).f27)
                    return iife185_(_pat_let47_0)
                r2 = (iife184_(d_1628_v1_)).cf4
                d_1647_v18_: _dafny.Array
                def lambda82_(d_1648_v0_):
                    def lambda83_(d_1649_i0_):
                        return (d_1649_i0_) + (d_1648_v0_)

                    return lambda83_

                init46_ = lambda82_(d_1627_v0_)
                nw257_ = _dafny.Array(None, 19)
                for i0_46_ in range(nw257_.length(0)):
                    nw257_[i0_46_] = init46_(i0_46_)
                d_1647_v18_ = nw257_
                index222_ = default__.safeIndex(135, (d_1647_v18_).length(0))
                (d_1647_v18_)[index222_] = d_1627_v0_
                d_1650_v19_: _dafny.Seq
                d_1650_v19_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_1627_v0_) + (d_1627_v0_))])
                index223_ = default__.safeIndex(135, (d_1647_v18_).length(0))
                rhs258_ = len(d_1650_v19_)
                rhs259_ = d_1627_v0_
                lhs189_ = d_1647_v18_
                lhs190_ = default__.safeIndex(135, (d_1647_v18_).length(0))
                lhs189_[lhs190_] = rhs258_
                d_1627_v0_ = rhs259_
                d_1651_v20_: C8
                nw258_ = C8()
                nw258_.ctor__((self).f24)
                d_1651_v20_ = nw258_
                d_1627_v0_ = (0) - (default__.safeModuloInt((d_1647_v18_)[default__.safeIndex(135, (d_1647_v18_).length(0))], d_1627_v0_))
            elif True:
                d_1652_v21_: C6
                nw259_ = C6()
                nw259_.ctor__(p0, not((d_1639_v12_).f27))
                d_1652_v21_ = nw259_
                d_1653_v22_: _dafny.Seq
                d_1653_v22_ = _dafny.SeqWithoutIsStrInference([d_1638_v11_])
                d_1654_v23_: D14
                d_1654_v23_ = D14_DC31((d_1639_v12_).f27)
                d_1655_v24_: C7
                nw260_ = C7()
                nw260_.ctor__(False, (d_1639_v12_).f27, (d_1653_v22_)[default__.safeIndex((0) - (d_1627_v0_), len(d_1653_v22_))], _dafny.CodePoint('w'), (self).f24, (d_1654_v23_).cf52)
                d_1655_v24_ = nw260_
                d_1656_v25_: _dafny.Array
                nw261_ = _dafny.Array(False, 2)
                d_1656_v25_ = nw261_
                d_1657_v26_: C5
                nw262_ = C5()
                nw262_.ctor__(d_1656_v25_)
                d_1657_v26_ = nw262_
                d_1658_v27_: _dafny.Array
                def lambda84_(d_1659_i1_):
                    return D15_DC34(False)

                init47_ = lambda84_
                nw263_ = _dafny.Array(None, 6)
                for i0_47_ in range(nw263_.length(0)):
                    nw263_[i0_47_] = init47_(i0_47_)
                d_1658_v27_ = nw263_
                d_1660_v28_: _dafny.Seq
                d_1660_v28_ = _dafny.SeqWithoutIsStrInference([d_1658_v27_])
                d_1661_v29_: D16
                d_1661_v29_ = D16_DC37(d_1657_v26_)
                d_1662_v30_: D16
                d_1662_v30_ = D16_DC37((d_1661_v29_).cf61)
                rhs260_ = len(d_1660_v28_)
                rhs261_ = d_1627_v0_
                rhs262_ = (d_1662_v30_).cf61
                rhs263_ = (self).fm6(globalState)
                rhs264_ = True
                lhs191_ = globalState
                lhs192_ = globalState
                lhs193_ = globalState
                d_1627_v0_ = rhs260_
                lhs191_.f7 = rhs261_
                d_1657_v26_ = rhs262_
                lhs192_.f7 = rhs263_
                lhs193_.f13 = rhs264_
                d_1663_v31_: _dafny.Seq
                d_1663_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtuyeu"))
                r0 = d_1663_v31_
                r2 = ((d_1639_v12_).f27) or (False)
            d_1664_v32_: _dafny.Array
            nw264_ = _dafny.Array(int(0), 5)
            d_1664_v32_ = nw264_
            d_1665_v33_: _dafny.Seq
            d_1665_v33_ = _dafny.SeqWithoutIsStrInference([d_1664_v32_])
            d_1666_v34_: _dafny.Array
            nw265_ = _dafny.Array(None, 17)
            nw265_[int(0)] = d_1664_v32_
            nw265_[int(1)] = d_1664_v32_
            nw265_[int(2)] = d_1664_v32_
            nw265_[int(3)] = d_1664_v32_
            nw265_[int(4)] = d_1664_v32_
            nw265_[int(5)] = d_1664_v32_
            nw265_[int(6)] = d_1664_v32_
            nw265_[int(7)] = d_1664_v32_
            nw265_[int(8)] = d_1664_v32_
            nw265_[int(9)] = d_1664_v32_
            nw265_[int(10)] = d_1664_v32_
            nw265_[int(11)] = d_1664_v32_
            nw265_[int(12)] = d_1664_v32_
            nw265_[int(13)] = d_1664_v32_
            nw265_[int(14)] = d_1664_v32_
            nw265_[int(15)] = d_1664_v32_
            nw265_[int(16)] = (d_1665_v33_)[default__.safeIndex(d_1627_v0_, len(d_1665_v33_))]
            d_1666_v34_ = nw265_
            d_1667_v35_: _dafny.Map
            d_1667_v35_ = _dafny.Map({p0: d_1666_v34_})
            d_1667_v35_ = (d_1667_v35_).set(_dafny.CodePoint('y'), d_1666_v34_)
            d_1668_v36_: _dafny.MultiSet
            d_1668_v36_ = _dafny.MultiSet([p0])
            d_1669_v37_: _dafny.MultiSet
            d_1669_v37_ = _dafny.MultiSet([d_1668_v36_])
            d_1670_v38_: C0
            nw266_ = C0()
            nw266_.ctor__((default__.fm48(globalState)) | (_dafny.Map({(d_1639_v12_).f27: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufkfma"))})), (self).f24, self.f23, (d_1669_v37_).ispropersubset(d_1669_v37_))
            d_1670_v38_ = nw266_
        d_1671_v39_: D5
        d_1671_v39_ = D5_DC13(d_1627_v0_, not((self).f24), (self).f24)
        source25_ = d_1671_v39_
        if source25_.is_DC13:
            d_1672___mcc_h0_ = source25_.cf20
            d_1673___mcc_h1_ = source25_.cf21
            d_1674___mcc_h2_ = source25_.cf22
            d_1675_cf22_ = d_1674___mcc_h2_
            d_1676_cf21_ = d_1673___mcc_h1_
            d_1677_cf20_ = d_1672___mcc_h0_
            d_1678_v40_: _dafny.Seq
            d_1678_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntvyfg"))
            (globalState).f13 = (d_1678_v40_) == ((d_1678_v40_) + (d_1678_v40_))
            d_1679_v41_: _dafny.Array
            nw267_ = _dafny.Array(int(0), 2)
            d_1679_v41_ = nw267_
            index224_ = default__.safeIndex(977, (d_1679_v41_).length(0))
            (d_1679_v41_)[index224_] = (self).fm6(globalState)
            d_1680_v42_: _dafny.Map
            d_1680_v42_ = _dafny.Map({not(not(p1)): d_1627_v0_})
            d_1681_v43_: _dafny.Seq
            d_1681_v43_ = _dafny.SeqWithoutIsStrInference([d_1680_v42_])
            d_1682_v44_: _dafny.Set
            d_1682_v44_ = _dafny.Set({p0, self.f23, self.f23, self.f23, self.f23})
            index225_ = default__.safeIndex(977, (d_1679_v41_).length(0))
            (d_1679_v41_)[index225_] = len((d_1681_v43_) + ((d_1681_v43_).set(default__.safeIndex(len(d_1682_v44_), len(d_1681_v43_)), d_1680_v42_)))
            d_1683_v45_: _dafny.MultiSet
            d_1683_v45_ = _dafny.MultiSet([p1])
            d_1684_v46_: _dafny.Seq
            d_1684_v46_ = _dafny.SeqWithoutIsStrInference([d_1683_v45_])
            d_1685_v47_: _dafny.Array
            nw268_ = _dafny.Array(_dafny.CodePoint('D'), 7)
            d_1685_v47_ = nw268_
            d_1686_v48_: C2
            nw269_ = C2()
            nw269_.ctor__((d_1684_v46_)[default__.safeIndex((d_1679_v41_)[default__.safeIndex(977, (d_1679_v41_).length(0))], len(d_1684_v46_))], d_1685_v47_, _dafny.CodePoint('s'), p1)
            d_1686_v48_ = nw269_
            d_1686_v48_ = d_1686_v48_
            d_1687_v49_: _dafny.Map
            d_1687_v49_ = _dafny.Map({(0) - ((d_1679_v41_)[default__.safeIndex(977, (d_1679_v41_).length(0))]): d_1676_cf21_})
            d_1687_v49_ = (d_1687_v49_).set((d_1679_v41_)[default__.safeIndex(977, (d_1679_v41_).length(0))], d_1675_cf22_)
        elif source25_.is_DC14:
            d_1688___mcc_h3_ = source25_.cf23
            d_1689_cf23_ = d_1688___mcc_h3_
            (globalState).f13 = d_1689_cf23_
            d_1628_v1_ = d_1628_v1_
            if (((self).fm11(d_1627_v0_, (self).f24, globalState) if p1 else p1)) and ((self).f24):
                d_1690_v50_: _dafny.Seq
                d_1690_v50_ = _dafny.SeqWithoutIsStrInference([p1, False, (self).f24, d_1689_cf23_, d_1689_cf23_])
                d_1691_v51_: _dafny.Map
                d_1691_v51_ = _dafny.Map({d_1627_v0_: self.f23})
                d_1692_v52_: _dafny.Map
                d_1692_v52_ = _dafny.Map({(d_1690_v50_)[default__.safeIndex(len(d_1691_v51_), len(d_1690_v50_))]: d_1689_cf23_})
                r2 = ((self).f24 if False else (self).fm4(len(d_1692_v52_), default__.fm36(globalState), globalState))
                d_1693_v53_: _dafny.Seq
                d_1693_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hj"))
                r0 = d_1693_v53_
                (globalState).f13 = d_1689_cf23_
                (globalState).f7 = d_1627_v0_
                d_1694_v54_: _dafny.MultiSet
                d_1694_v54_ = _dafny.MultiSet([not(False)])
                d_1695_v55_: int
                d_1696_v56_: bool
                d_1697_v57_: int
                out49_: int
                out50_: bool
                out51_: int
                out49_, out50_, out51_ = default__.m0(d_1627_v0_, (_dafny.MultiSet([p1, (self).f24])).intersection(d_1694_v54_), p1, default__.safeDivisionInt(d_1627_v0_, 878), globalState)
                d_1695_v55_ = out49_
                d_1696_v56_ = out50_
                d_1697_v57_ = out51_
            elif True:
                d_1698_v58_: _dafny.Array
                def lambda85_(d_1699_i2_):
                    return False

                init48_ = lambda85_
                nw270_ = _dafny.Array(None, 4)
                for i0_48_ in range(nw270_.length(0)):
                    nw270_[i0_48_] = init48_(i0_48_)
                d_1698_v58_ = nw270_
                d_1700_v59_: _dafny.Map
                d_1700_v59_ = _dafny.Map({d_1698_v58_: p0})
                d_1701_v60_: _dafny.Map
                d_1701_v60_ = _dafny.Map({default__.safeDivisionInt(d_1627_v0_, d_1627_v0_): d_1700_v59_})
                d_1701_v60_ = (d_1701_v60_).set(d_1627_v0_, d_1700_v59_)
                d_1702_v61_: _dafny.Array
                def lambda86_(d_1703_v0_):
                    def lambda87_(d_1704_i3_):
                        return default__.safeModuloInt(d_1704_i3_, d_1703_v0_)

                    return lambda87_

                init49_ = lambda86_(d_1627_v0_)
                nw271_ = _dafny.Array(None, 8)
                for i0_49_ in range(nw271_.length(0)):
                    nw271_[i0_49_] = init49_(i0_49_)
                d_1702_v61_ = nw271_
                d_1705_v62_: _dafny.Seq
                d_1705_v62_ = _dafny.SeqWithoutIsStrInference([p1])
                index226_ = default__.safeIndex(247, (d_1702_v61_).length(0))
                (d_1702_v61_)[index226_] = (d_1627_v0_) * (default__.fm0(d_1689_cf23_, len(d_1705_v62_), globalState))
                index227_ = default__.safeIndex(247, (d_1702_v61_).length(0))
                (d_1702_v61_)[index227_] = d_1627_v0_
                (globalState).f7 = d_1627_v0_
                d_1706_v63_: _dafny.Seq
                d_1706_v63_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_1707_v64_: _dafny.Array
                nw272_ = _dafny.Array(None, 27)
                nw272_[int(0)] = p0
                nw272_[int(1)] = self.f23
                nw272_[int(2)] = (d_1706_v63_)[default__.safeIndex(len(d_1706_v63_), len(d_1706_v63_))]
                nw272_[int(3)] = self.f23
                nw272_[int(4)] = p0
                nw272_[int(5)] = p0
                nw272_[int(6)] = (d_1706_v63_)[default__.safeIndex(d_1627_v0_, len(d_1706_v63_))]
                nw272_[int(7)] = self.f23
                nw272_[int(8)] = self.f23
                nw272_[int(9)] = self.f23
                nw272_[int(10)] = p0
                nw272_[int(11)] = p0
                nw272_[int(12)] = self.f23
                nw272_[int(13)] = _dafny.CodePoint('u')
                nw272_[int(14)] = p0
                nw272_[int(15)] = self.f23
                nw272_[int(16)] = self.f23
                nw272_[int(17)] = p0
                nw272_[int(18)] = self.f23
                nw272_[int(19)] = p0
                nw272_[int(20)] = p0
                nw272_[int(21)] = (d_1706_v63_)[default__.safeIndex((d_1702_v61_)[default__.safeIndex(247, (d_1702_v61_).length(0))], len(d_1706_v63_))]
                nw272_[int(22)] = p0
                nw272_[int(23)] = p0
                nw272_[int(24)] = default__.fm29((d_1702_v61_)[default__.safeIndex(247, (d_1702_v61_).length(0))], p1, d_1689_cf23_, d_1627_v0_, globalState)
                nw272_[int(25)] = p0
                nw272_[int(26)] = p0
                d_1707_v64_ = nw272_
                d_1708_v65_: D12
                d_1708_v65_ = D12_DC24(d_1707_v64_)
                d_1709_v66_: C7
                nw273_ = C7()
                nw273_.ctor__(((self).f24 if True else True), (self).f24, (d_1708_v65_).cf40, self.f23, not(((d_1702_v61_)[default__.safeIndex(247, (d_1702_v61_).length(0))]) > (d_1627_v0_)), p1)
                d_1709_v66_ = nw273_
                d_1710_v67_: _dafny.Array
                nw274_ = _dafny.Array(D13.default()(), 6)
                d_1710_v67_ = nw274_
                d_1711_v68_: D13
                d_1711_v68_ = D13_DC29((d_1709_v66_).f30, (d_1709_v66_).f31)
                index228_ = default__.safeIndex(852, (d_1710_v67_).length(0))
                (d_1710_v67_)[index228_] = d_1711_v68_
                index229_ = default__.safeIndex(852, (d_1710_v67_).length(0))
                (d_1710_v67_)[index229_] = d_1711_v68_
            d_1689_cf23_ = (self).f24
        elif True:
            d_1712___mcc_h4_ = source25_.cf19
            d_1713_cf19_ = d_1712___mcc_h4_
            d_1714_v69_: _dafny.Array
            def lambda88_(d_1715_v0_):
                def lambda89_(d_1716_i4_):
                    return (d_1716_i4_) - (d_1715_v0_)

                return lambda89_

            init50_ = lambda88_(d_1627_v0_)
            nw275_ = _dafny.Array(None, 16)
            for i0_50_ in range(nw275_.length(0)):
                nw275_[i0_50_] = init50_(i0_50_)
            d_1714_v69_ = nw275_
            index230_ = default__.safeIndex(24, (d_1714_v69_).length(0))
            (d_1714_v69_)[index230_] = d_1627_v0_
            index231_ = default__.safeIndex(24, (d_1714_v69_).length(0))
            (d_1714_v69_)[index231_] = d_1627_v0_
            d_1717_v70_: _dafny.Seq
            d_1717_v70_ = _dafny.SeqWithoutIsStrInference([d_1627_v0_])
            rhs265_ = (d_1717_v70_) + (d_1717_v70_)
            rhs266_ = not(p1)
            lhs194_ = globalState
            d_1717_v70_ = rhs265_
            lhs194_.f13 = rhs266_
            d_1718_v71_: _dafny.Array
            def lambda90_(d_1719_p1_):
                def lambda91_(d_1720_i5_):
                    return (D2_DC7(d_1719_p1_, (self).f24)).cf14

                return lambda91_

            init51_ = lambda90_(p1)
            nw276_ = _dafny.Array(None, 5)
            for i0_51_ in range(nw276_.length(0)):
                nw276_[i0_51_] = init51_(i0_51_)
            d_1718_v71_ = nw276_
            d_1721_v72_: _dafny.Seq
            d_1721_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbxbjje"))
            d_1722_v73_: _dafny.Map
            d_1722_v73_ = _dafny.Map({(self).f24: d_1721_v72_})
            d_1723_v74_: _dafny.Seq
            d_1723_v74_ = _dafny.SeqWithoutIsStrInference([((d_1722_v73_)[False] if (False) in (d_1722_v73_) else d_1721_v72_)])
            d_1724_v75_: _dafny.Map
            d_1724_v75_ = _dafny.Map({not(p1): p1})
            d_1725_v76_: _dafny.Map
            d_1725_v76_ = _dafny.Map({default__.safeModuloInt(len(d_1723_v74_), len(d_1724_v75_)): d_1718_v71_})
            d_1718_v71_ = ((d_1725_v76_)[d_1627_v0_] if (d_1627_v0_) in (d_1725_v76_) else d_1718_v71_)
            d_1726_v77_: D0
            d_1726_v77_ = D0_DC1(-87)
            r1 = d_1726_v77_
        if p1:
            d_1727_v78_: _dafny.Array
            def lambda92_(d_1728_v0_):
                def lambda93_(d_1729_i6_):
                    return (d_1729_i6_) - (d_1728_v0_)

                return lambda93_

            init52_ = lambda92_(d_1627_v0_)
            nw277_ = _dafny.Array(None, 26)
            for i0_52_ in range(nw277_.length(0)):
                nw277_[i0_52_] = init52_(i0_52_)
            d_1727_v78_ = nw277_
            (globalState).f22 = d_1727_v78_
            d_1730_v79_: _dafny.Map
            d_1730_v79_ = _dafny.Map({p0: (self).f24})
            (globalState).f7 = default__.fm0(p1, len(d_1730_v79_), globalState)
            (globalState).f13 = (self).f24
            d_1731_v80_: T1
            nw278_ = C6()
            nw278_.ctor__(self.f23, (self).f24)
            d_1731_v80_ = nw278_
            d_1731_v80_ = d_1731_v80_
            (globalState).f13 = ((d_1731_v80_).f24 if False else (self).f24)
        elif True:
            rhs267_ = 86
            rhs268_ = ((self).fm11(d_1627_v0_, p1, globalState) if p1 else not (p1) or (p1))
            rhs269_ = (self).fm11(-862, (self).f24, globalState)
            rhs270_ = (self).fm11(d_1627_v0_, p1, globalState)
            rhs271_ = 9
            lhs195_ = globalState
            lhs196_ = globalState
            lhs197_ = globalState
            lhs198_ = globalState
            lhs195_.f7 = rhs267_
            r2 = rhs268_
            lhs196_.f13 = rhs269_
            lhs197_.f13 = rhs270_
            lhs198_.f7 = rhs271_
            d_1732_v81_: _dafny.Seq
            d_1732_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
            d_1733_v82_: _dafny.Array
            def lambda94_(d_1734_i7_):
                return (self).f24

            init53_ = lambda94_
            nw279_ = _dafny.Array(None, 14)
            for i0_53_ in range(nw279_.length(0)):
                nw279_[i0_53_] = init53_(i0_53_)
            d_1733_v82_ = nw279_
            d_1735_v83_: _dafny.Set
            d_1735_v83_ = _dafny.Set({(self).f24, (self).f24})
            d_1736_v84_: D16
            d_1736_v84_ = D16_DC38((d_1732_v81_) + (d_1732_v81_), d_1733_v82_, d_1735_v83_)
            source26_ = d_1736_v84_
            if source26_.is_DC38:
                d_1737___mcc_h5_ = source26_.cf62
                d_1738___mcc_h6_ = source26_.cf63
                d_1739___mcc_h7_ = source26_.cf64
                d_1740_cf64_ = d_1739___mcc_h7_
                d_1741_cf63_ = d_1738___mcc_h6_
                d_1742_cf62_ = d_1737___mcc_h5_
                d_1743_v85_: _dafny.Map
                d_1743_v85_ = _dafny.Map({524: d_1627_v0_})
                d_1744_v86_: _dafny.Map
                d_1744_v86_ = _dafny.Map({d_1627_v0_: len(d_1743_v85_)})
                d_1745_v87_: _dafny.Map
                d_1745_v87_ = _dafny.Map({(d_1627_v0_) - (len(d_1744_v86_)): d_1627_v0_})
                d_1743_v85_ = d_1743_v85_
                d_1746_v88_: C5
                nw280_ = C5()
                nw280_.ctor__(d_1741_cf63_)
                d_1746_v88_ = nw280_
                d_1747_v89_: _dafny.Array
                nw281_ = _dafny.Array(int(0), 21)
                d_1747_v89_ = nw281_
                index232_ = default__.safeIndex(466, (d_1747_v89_).length(0))
                (d_1747_v89_)[index232_] = (0) - (d_1627_v0_)
                d_1748_v90_: _dafny.Map
                d_1748_v90_ = _dafny.Map({(self).f24: p1})
                d_1749_v91_: D5
                d_1749_v91_ = D5_DC14(p1)
                d_1750_v92_: D11
                d_1750_v92_ = D11_DC23(p1, len(_dafny.SeqWithoutIsStrInference([p1, not(p1)])), (d_1748_v90_).set((self).f24, not((self).f24)), d_1749_v91_)
                index233_ = default__.safeIndex(466, (d_1747_v89_).length(0))
                rhs272_ = d_1627_v0_
                rhs273_ = (d_1750_v92_).cf37
                rhs274_ = (self).f24
                lhs199_ = d_1747_v89_
                lhs200_ = default__.safeIndex(466, (d_1747_v89_).length(0))
                lhs201_ = globalState
                lhs202_ = globalState
                lhs199_[lhs200_] = rhs272_
                lhs201_.f7 = rhs273_
                lhs202_.f13 = rhs274_
                d_1748_v90_ = (d_1748_v90_).set((self).f24, (self).f24)
            elif source26_.is_DC37:
                d_1751___mcc_h8_ = source26_.cf61
                d_1752_cf61_ = d_1751___mcc_h8_
                d_1753_v93_: _dafny.Array
                nw282_ = _dafny.Array(int(0), 6)
                d_1753_v93_ = nw282_
                index234_ = default__.safeIndex(374, (d_1753_v93_).length(0))
                (d_1753_v93_)[index234_] = d_1627_v0_
                index235_ = default__.safeIndex(374, (d_1753_v93_).length(0))
                (d_1753_v93_)[index235_] = (d_1627_v0_) * (d_1627_v0_)
                (self).f23 = p0
                d_1627_v0_ = default__.fm0(not((d_1627_v0_) >= (default__.fm0(True, 364, globalState))), (d_1753_v93_)[default__.safeIndex(374, (d_1753_v93_).length(0))], globalState)
                d_1754_v94_: _dafny.Array
                nw283_ = _dafny.Array(_dafny.MultiSet({}), 25)
                d_1754_v94_ = nw283_
                d_1755_v95_: _dafny.MultiSet
                d_1755_v95_ = _dafny.MultiSet([d_1627_v0_])
                index236_ = default__.safeIndex(144, (d_1754_v94_).length(0))
                (d_1754_v94_)[index236_] = (d_1755_v95_).intersection(d_1755_v95_)
                index237_ = default__.safeIndex(144, (d_1754_v94_).length(0))
                (d_1754_v94_)[index237_] = d_1755_v95_
            elif True:
                d_1756___mcc_h9_ = source26_.cf65
                d_1757_cf65_ = d_1756___mcc_h9_
                (globalState).f13 = not (p1) or (False)
                d_1758_v96_: _dafny.Seq
                d_1758_v96_ = _dafny.SeqWithoutIsStrInference([d_1627_v0_])
                d_1759_v97_: _dafny.Array
                nw284_ = _dafny.Array(None, 25)
                nw284_[int(0)] = (0) - (len(d_1732_v81_))
                nw284_[int(1)] = (d_1758_v96_)[default__.safeIndex(657, len(d_1758_v96_))]
                nw284_[int(2)] = d_1627_v0_
                nw284_[int(3)] = 521
                nw284_[int(4)] = (0) - (d_1627_v0_)
                nw284_[int(5)] = (_dafny.MultiSet(d_1758_v96_)).cardinality
                nw284_[int(6)] = d_1627_v0_
                nw284_[int(7)] = (0) - (d_1627_v0_)
                nw284_[int(8)] = (d_1628_v1_).cf3
                nw284_[int(9)] = d_1627_v0_
                nw284_[int(10)] = d_1627_v0_
                nw284_[int(11)] = d_1627_v0_
                nw284_[int(12)] = d_1627_v0_
                nw284_[int(13)] = d_1627_v0_
                nw284_[int(14)] = default__.fm0(not(p1), d_1627_v0_, globalState)
                nw284_[int(15)] = d_1627_v0_
                nw284_[int(16)] = d_1627_v0_
                nw284_[int(17)] = d_1627_v0_
                nw284_[int(18)] = d_1627_v0_
                nw284_[int(19)] = d_1627_v0_
                nw284_[int(20)] = d_1627_v0_
                nw284_[int(21)] = (d_1627_v0_) * (d_1627_v0_)
                nw284_[int(22)] = d_1627_v0_
                nw284_[int(23)] = d_1627_v0_
                nw284_[int(24)] = (d_1627_v0_) + (len(d_1732_v81_))
                d_1759_v97_ = nw284_
                d_1760_v98_: D5
                d_1760_v98_ = D5_DC14(p1)
                d_1761_v99_: D11
                d_1761_v99_ = D11_DC23(p1, d_1627_v0_, _dafny.Map({p1: p1}), d_1760_v98_)
                d_1762_v100_: _dafny.Seq
                d_1762_v100_ = _dafny.SeqWithoutIsStrInference([d_1761_v99_, d_1761_v99_, d_1761_v99_, d_1761_v99_, d_1761_v99_])
                index238_ = default__.safeIndex(119, (d_1759_v97_).length(0))
                (d_1759_v97_)[index238_] = (((_dafny.MultiSet(d_1762_v100_)).set(d_1761_v99_, default__.abs(d_1627_v0_))).intersection(_dafny.MultiSet([d_1761_v99_]))).cardinality
                index239_ = default__.safeIndex(119, (d_1759_v97_).length(0))
                (d_1759_v97_)[index239_] = default__.safeDivisionInt(d_1627_v0_, ((0) - (d_1627_v0_)) - (d_1627_v0_))
                d_1627_v0_ = 375
                index240_ = default__.safeIndex(119, (d_1759_v97_).length(0))
                (d_1759_v97_)[index240_] = d_1627_v0_
            d_1763_v101_: _dafny.Array
            nw285_ = _dafny.Array(int(0), 17)
            d_1763_v101_ = nw285_
            index241_ = default__.safeIndex(378, (d_1763_v101_).length(0))
            (d_1763_v101_)[index241_] = d_1627_v0_
            d_1764_v102_: _dafny.Map
            d_1764_v102_ = _dafny.Map({(self).f24: d_1627_v0_})
            d_1765_v103_: _dafny.Seq
            d_1765_v103_ = _dafny.SeqWithoutIsStrInference([p1])
            index242_ = default__.safeIndex(378, (d_1763_v101_).length(0))
            rhs275_ = ((d_1764_v102_)[(self).f24] if ((self).f24) in (d_1764_v102_) else len(d_1765_v103_))
            rhs276_ = d_1627_v0_
            lhs203_ = globalState
            lhs204_ = d_1763_v101_
            lhs205_ = default__.safeIndex(378, (d_1763_v101_).length(0))
            lhs203_.f7 = rhs275_
            lhs204_[lhs205_] = rhs276_
            d_1766_v104_: C8
            nw286_ = C8()
            nw286_.ctor__((self).f24)
            d_1766_v104_ = nw286_
            d_1767_v105_: _dafny.Seq
            d_1767_v105_ = _dafny.SeqWithoutIsStrInference([d_1627_v0_, (d_1763_v101_)[default__.safeIndex(378, (d_1763_v101_).length(0))]])
            d_1768_v106_: _dafny.Seq
            d_1768_v106_ = _dafny.SeqWithoutIsStrInference([d_1627_v0_, (d_1767_v105_)[default__.safeIndex((d_1763_v101_)[default__.safeIndex(378, (d_1763_v101_).length(0))], len(d_1767_v105_))]])
            d_1769_v107_: _dafny.MultiSet
            d_1769_v107_ = _dafny.MultiSet([d_1627_v0_, len(d_1768_v106_), d_1627_v0_, d_1627_v0_])
            r2 = ((_dafny.MultiSet([(d_1763_v101_)[default__.safeIndex(378, (d_1763_v101_).length(0))]])).isdisjoint(d_1769_v107_) if not((self).f24) else (self).f24)
        d_1770_v108_: _dafny.Map
        d_1770_v108_ = _dafny.Map({d_1671_v39_: (self).f24})
        d_1770_v108_ = (d_1770_v108_).set(d_1671_v39_, (self).f24)
        if (self).f24:
            d_1771_v109_: _dafny.Array
            nw287_ = _dafny.Array(False, 21)
            d_1771_v109_ = nw287_
            d_1772_v110_: _dafny.Map
            d_1772_v110_ = _dafny.Map({d_1771_v109_: _dafny.CodePoint('y')})
            d_1773_v111_: _dafny.Seq
            d_1773_v111_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hu"))
            d_1774_v112_: _dafny.Array
            nw288_ = _dafny.Array(None, 11)
            nw288_[int(0)] = self.f23
            nw288_[int(1)] = _dafny.CodePoint('p')
            nw288_[int(2)] = p0
            nw288_[int(3)] = p0
            nw288_[int(4)] = self.f23
            nw288_[int(5)] = p0
            nw288_[int(6)] = self.f23
            nw288_[int(7)] = _dafny.CodePoint('h')
            nw288_[int(8)] = p0
            nw288_[int(9)] = p0
            nw288_[int(10)] = (d_1773_v111_)[default__.safeIndex(d_1627_v0_, len(d_1773_v111_))]
            d_1774_v112_ = nw288_
            d_1775_v113_: C3
            nw289_ = C3()
            nw289_.ctor__(d_1772_v110_, d_1774_v112_, p0, (self).f24)
            d_1775_v113_ = nw289_
            d_1776_v114_: _dafny.Seq
            d_1776_v114_ = _dafny.SeqWithoutIsStrInference([d_1627_v0_, d_1627_v0_, 129, d_1627_v0_])
            d_1777_v115_: _dafny.Map
            d_1777_v115_ = _dafny.Map({d_1775_v113_: default__.safeDivisionInt((d_1776_v114_)[default__.safeIndex(d_1627_v0_, len(d_1776_v114_))], d_1627_v0_)})
            d_1777_v115_ = d_1777_v115_
            d_1778_v116_: T3
            nw290_ = C7()
            nw290_.ctor__(False, (self).f24, d_1774_v112_, p0, p1, not(True))
            d_1778_v116_ = nw290_
            d_1779_v117_: _dafny.Array
            nw291_ = _dafny.Array(None, 7)
            nw291_[int(0)] = d_1778_v116_
            nw291_[int(1)] = d_1778_v116_
            nw291_[int(2)] = d_1778_v116_
            nw291_[int(3)] = d_1778_v116_
            nw291_[int(4)] = d_1778_v116_
            nw291_[int(5)] = d_1778_v116_
            nw291_[int(6)] = d_1778_v116_
            d_1779_v117_ = nw291_
            index243_ = default__.safeIndex(538, (d_1779_v117_).length(0))
            (d_1779_v117_)[index243_] = d_1778_v116_
            index244_ = default__.safeIndex(538, (d_1779_v117_).length(0))
            (d_1779_v117_)[index244_] = (d_1778_v116_ if (d_1778_v116_).f27 else d_1778_v116_)
            d_1780_v118_: _dafny.Map
            d_1780_v118_ = _dafny.Map({d_1773_v111_: (d_1771_v109_ if not((d_1778_v116_).f27) else d_1771_v109_)})
            d_1781_v119_: _dafny.Seq
            d_1781_v119_ = _dafny.SeqWithoutIsStrInference([(self).f24, (self).f24])
            d_1782_v120_: D4
            d_1782_v120_ = D4_DC10(d_1773_v111_)
            d_1783_v121_: _dafny.Map
            d_1783_v121_ = _dafny.Map({(self).f24: d_1773_v111_})
            d_1784_v122_: _dafny.Seq
            d_1784_v122_ = _dafny.SeqWithoutIsStrInference([((d_1783_v121_)[p1] if (p1) in (d_1783_v121_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpa"))), d_1773_v111_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfygf"))])
            d_1785_v123_: _dafny.MultiSet
            d_1785_v123_ = _dafny.MultiSet([(d_1778_v116_).f27])
            d_1786_v124_: C2
            nw292_ = C2()
            nw292_.ctor__(d_1785_v123_, d_1774_v112_, self.f23, (self).f24)
            d_1786_v124_ = nw292_
            d_1787_v125_: _dafny.Seq
            d_1787_v125_ = _dafny.SeqWithoutIsStrInference([d_1786_v124_, d_1786_v124_])
            rhs277_ = ((self).fm4((self).fm6(globalState), d_1781_v119_, globalState)) or (p1)
            rhs278_ = d_1780_v118_
            rhs279_ = (d_1627_v0_) - (len(_dafny.SeqWithoutIsStrInference([d_1782_v120_])))
            rhs280_ = (self).f24
            rhs281_ = (d_1784_v122_)[default__.safeIndex(default__.safeModuloInt(d_1627_v0_, len(d_1787_v125_)), len(d_1784_v122_))]
            lhs206_ = globalState
            lhs207_ = globalState
            r2 = rhs277_
            d_1780_v118_ = rhs278_
            lhs206_.f7 = rhs279_
            lhs207_.f13 = rhs280_
            d_1773_v111_ = rhs281_
            d_1788_v126_: C1
            nw293_ = C1()
            nw293_.ctor__(p1, (_dafny.SeqWithoutIsStrInference([default__.fm2(globalState), (d_1778_v116_).f27])) == (d_1781_v119_))
            d_1788_v126_ = nw293_
            r2 = (self).f24
        elif True:
            d_1789_v127_: _dafny.Map
            d_1789_v127_ = _dafny.Map({(d_1627_v0_) * (d_1627_v0_): d_1627_v0_})
            d_1789_v127_ = d_1789_v127_
            (globalState).f7 = (0) - (d_1627_v0_)
            d_1628_v1_ = d_1628_v1_
            d_1790_v128_: _dafny.Array
            def lambda95_(d_1791_v0_):
                def lambda96_(d_1792_i8_):
                    return (d_1792_i8_) * (d_1791_v0_)

                return lambda96_

            init54_ = lambda95_(d_1627_v0_)
            nw294_ = _dafny.Array(None, 12)
            for i0_54_ in range(nw294_.length(0)):
                nw294_[i0_54_] = init54_(i0_54_)
            d_1790_v128_ = nw294_
            d_1793_v129_: _dafny.Set
            d_1793_v129_ = _dafny.Set({d_1627_v0_})
            index245_ = default__.safeIndex(172, (d_1790_v128_).length(0))
            (d_1790_v128_)[index245_] = len(d_1793_v129_)
            index246_ = default__.safeIndex(172, (d_1790_v128_).length(0))
            (d_1790_v128_)[index246_] = ((d_1627_v0_ if (self).f24 else d_1627_v0_)) + (d_1627_v0_)
            d_1794_v130_: _dafny.MultiSet
            d_1794_v130_ = _dafny.MultiSet([d_1627_v0_])
            if ((d_1794_v130_) - (_dafny.MultiSet([(self).fm6(globalState), d_1627_v0_]))).issubset(d_1794_v130_):
                d_1789_v127_ = (d_1789_v127_).set((0) - (((d_1790_v128_)[default__.safeIndex(172, (d_1790_v128_).length(0))]) + ((d_1790_v128_)[default__.safeIndex(172, (d_1790_v128_).length(0))])), (0) - (d_1627_v0_))
                d_1795_v131_: _dafny.Array
                nw295_ = _dafny.Array(False, 9)
                d_1795_v131_ = nw295_
                d_1796_v132_: C5
                nw296_ = C5()
                nw296_.ctor__(d_1795_v131_)
                d_1796_v132_ = nw296_
                r2 = (self).f24
                (globalState).f13 = (D5_DC13(len(_dafny.SeqWithoutIsStrInference([(0) - (d_1627_v0_)])), not((self).f24), p1)).cf22
                d_1797_v133_: _dafny.MultiSet
                d_1797_v133_ = _dafny.MultiSet([(self).f24])
                d_1798_v134_: D0
                d_1798_v134_ = D0_DC0(d_1797_v133_)
                d_1798_v134_ = d_1798_v134_
            elif True:
                (globalState).f7 = d_1627_v0_
                d_1799_v135_: _dafny.Map
                d_1799_v135_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f24, (self).f24, p1]): p1})
                d_1800_v136_: _dafny.Seq
                d_1800_v136_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1801_v137_: _dafny.Seq
                d_1801_v137_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "blrabrw"))
                d_1802_v138_: _dafny.Map
                d_1802_v138_ = _dafny.Map({d_1627_v0_: d_1801_v137_})
                d_1803_v139_: _dafny.Map
                d_1803_v139_ = _dafny.Map({(self).fm4(len(d_1799_v135_), d_1800_v136_, globalState): ((d_1802_v138_)[d_1627_v0_] if (d_1627_v0_) in (d_1802_v138_) else _dafny.SeqWithoutIsStrInference([self.f23 for d_1804_i9_ in range(default__.abs(-330))]))})
                d_1805_v140_: C1
                nw297_ = C1()
                nw297_.ctor__((self).f24, False)
                d_1805_v140_ = nw297_
                d_1806_v141_: _dafny.Map
                d_1806_v141_ = _dafny.Map({d_1805_v140_: d_1627_v0_})
                d_1807_v142_: _dafny.Map
                d_1807_v142_ = _dafny.Map({d_1806_v141_: (self).f24})
                d_1808_v143_: _dafny.Map
                d_1808_v143_ = _dafny.Map({not((self).f24): (0) - (len(d_1807_v142_))})
                d_1809_v144_: _dafny.Map
                d_1809_v144_ = _dafny.Map({d_1808_v143_: p1})
                d_1810_v145_: C0
                nw298_ = C0()
                nw298_.ctor__(d_1803_v139_, (self).f24, default__.fm29((d_1790_v128_)[default__.safeIndex(172, (d_1790_v128_).length(0))], p1, p1, (d_1790_v128_)[default__.safeIndex(172, (d_1790_v128_).length(0))], globalState), (d_1809_v144_) == (_dafny.Map({_dafny.Map({(self).f24: len(d_1800_v136_)}): (d_1805_v140_).f39})))
                d_1810_v145_ = nw298_
                (self).f23 = (d_1801_v137_)[default__.safeIndex(d_1627_v0_, len(d_1801_v137_))]
                d_1811_v146_: _dafny.Seq
                d_1811_v146_ = _dafny.SeqWithoutIsStrInference([(d_1790_v128_)[default__.safeIndex(172, (d_1790_v128_).length(0))], ((d_1790_v128_)[default__.safeIndex(172, (d_1790_v128_).length(0))]) - ((d_1790_v128_)[default__.safeIndex(172, (d_1790_v128_).length(0))]), d_1627_v0_, (d_1790_v128_)[default__.safeIndex(172, (d_1790_v128_).length(0))]])
                (globalState).f7 = (d_1811_v146_)[default__.safeIndex((d_1790_v128_)[default__.safeIndex(172, (d_1790_v128_).length(0))], len(d_1811_v146_))]
                d_1812_v147_: _dafny.MultiSet
                d_1812_v147_ = _dafny.MultiSet([(d_1810_v145_).f38])
                (globalState).f7 = ((d_1790_v128_)[default__.safeIndex(172, (d_1790_v128_).length(0))] if (not(not((d_1805_v140_).f39))) not in (d_1812_v147_) else default__.safeModuloInt((d_1790_v128_)[default__.safeIndex(172, (d_1790_v128_).length(0))], (d_1790_v128_)[default__.safeIndex(172, (d_1790_v128_).length(0))]))
        d_1813_v148_: _dafny.Array
        nw299_ = _dafny.Array(int(0), 24)
        d_1813_v148_ = nw299_
        d_1814_v149_: _dafny.Seq
        d_1814_v149_ = _dafny.SeqWithoutIsStrInference([d_1813_v148_, d_1813_v148_])
        hi11_ = (len(d_1814_v149_) if (self).f24 else d_1627_v0_)
        for d_1815_i10_ in range(d_1627_v0_, hi11_):
            d_1816_v150_: _dafny.Seq
            d_1816_v150_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "enqsgx"))
            (globalState).f13 = ((default__.fm1((self).f24, globalState)) + (d_1816_v150_)) < (d_1816_v150_)
            d_1817_v151_: _dafny.Array
            nw300_ = _dafny.Array(False, 16)
            d_1817_v151_ = nw300_
            index247_ = default__.safeIndex(802, (d_1817_v151_).length(0))
            (d_1817_v151_)[index247_] = (self).f24
            d_1818_v152_: _dafny.Map
            d_1818_v152_ = _dafny.Map({p1: d_1816_v150_})
            d_1819_v153_: _dafny.Seq
            d_1819_v153_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1820_v154_: T0
            nw301_ = C0()
            nw301_.ctor__(d_1818_v152_, (d_1819_v153_)[default__.safeIndex(d_1815_i10_, len(d_1819_v153_))], self.f23, not((self).f24))
            d_1820_v154_ = nw301_
            d_1821_v155_: _dafny.MultiSet
            d_1821_v155_ = _dafny.MultiSet([d_1815_i10_])
            index248_ = default__.safeIndex(802, (d_1817_v151_).length(0))
            rhs282_ = False
            rhs283_ = (0) - (d_1815_i10_)
            rhs284_ = d_1820_v154_
            rhs285_ = (d_1821_v155_).intersection(_dafny.MultiSet([d_1815_i10_, d_1815_i10_, d_1627_v0_]))
            rhs286_ = ((d_1819_v153_) + (_dafny.SeqWithoutIsStrInference([True]))) + ((d_1819_v153_) + (d_1819_v153_))
            lhs208_ = d_1817_v151_
            lhs209_ = default__.safeIndex(802, (d_1817_v151_).length(0))
            lhs208_[lhs209_] = rhs282_
            d_1627_v0_ = rhs283_
            d_1820_v154_ = rhs284_
            d_1821_v155_ = rhs285_
            d_1819_v153_ = rhs286_
            d_1822_v156_: C8
            nw302_ = C8()
            nw302_.ctor__(p1)
            d_1822_v156_ = nw302_
            d_1823_v157_: _dafny.MultiSet
            d_1823_v157_ = _dafny.MultiSet([d_1822_v156_])
            d_1824_v158_: _dafny.Set
            d_1824_v158_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lufwkuo"))), d_1815_i10_, ((d_1823_v157_).set(d_1822_v156_, default__.abs(d_1815_i10_))).cardinality})
            d_1825_v159_: _dafny.Set
            d_1825_v159_ = _dafny.Set({len(d_1824_v158_), d_1815_i10_, len(_dafny.SeqWithoutIsStrInference([d_1820_v154_.f23 for d_1826_i11_ in range(default__.abs(114))]))})
            (globalState).f7 = (default__.safeDivisionInt(d_1627_v0_, len(d_1825_v159_)) if p1 else d_1815_i10_)
            d_1827_v160_: _dafny.Map
            d_1827_v160_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")): d_1820_v154_.f23})
            (d_1820_v154_).f23 = ((d_1827_v160_)[(d_1816_v150_) + (d_1816_v150_)] if ((d_1816_v150_) + (d_1816_v150_)) in (d_1827_v160_) else d_1820_v154_.f23)
        r0 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oyguroyys"))
        d_1828_v161_: D0
        d_1828_v161_ = D0_DC1(d_1627_v0_)
        r1 = d_1828_v161_
        d_1829_v162_: C1
        nw303_ = C1()
        nw303_.ctor__(p1, p1)
        d_1829_v162_ = nw303_
        d_1830_v163_: _dafny.Set
        d_1830_v163_ = _dafny.Set({d_1829_v162_, d_1829_v162_})
        r2 = (d_1830_v163_).ispropersubset(d_1830_v163_)
        return r0, r1, r2

    def m2(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        source27_ = D13_DC27(_dafny.Map({(self).f24: self.f23}))
        if source27_.is_DC28:
            d_1831___mcc_h0_ = source27_.cf45
            d_1832___mcc_h1_ = source27_.cf46
            d_1833___mcc_h2_ = source27_.cf47
            d_1834___mcc_h3_ = source27_.cf48
            d_1835_cf48_ = d_1834___mcc_h3_
            d_1836_cf47_ = d_1833___mcc_h2_
            d_1837_cf46_ = d_1832___mcc_h1_
            d_1838_cf45_ = d_1831___mcc_h0_
            d_1839_v0_: _dafny.Array
            def lambda97_(d_1840_cf45_):
                def lambda98_(d_1841_i0_):
                    return d_1840_cf45_

                return lambda98_

            init55_ = lambda97_(d_1838_cf45_)
            nw304_ = _dafny.Array(None, 27)
            for i0_55_ in range(nw304_.length(0)):
                nw304_[i0_55_] = init55_(i0_55_)
            d_1839_v0_ = nw304_
            d_1842_v1_: _dafny.Seq
            d_1842_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ig"))
            d_1843_v2_: _dafny.Map
            d_1843_v2_ = _dafny.Map({d_1839_v0_: (d_1842_v1_)[default__.safeIndex(d_1836_cf47_, len(d_1842_v1_))]})
            d_1844_v3_: _dafny.Array
            nw305_ = _dafny.Array(_dafny.CodePoint('D'), 23)
            d_1844_v3_ = nw305_
            d_1845_v4_: C3
            nw306_ = C3()
            nw306_.ctor__(d_1843_v2_, d_1844_v3_, self.f23, d_1838_cf45_)
            d_1845_v4_ = nw306_
            d_1846_v5_: D2
            d_1846_v5_ = D2_DC6(_dafny.Map({True: d_1842_v1_}))
            d_1847_v6_: _dafny.Seq
            d_1847_v6_ = default__.fm45((self).f24, d_1846_v5_, len(d_1842_v1_), globalState)
            d_1848_v7_: _dafny.Map
            d_1848_v7_ = _dafny.Map({((self).f24) or (True): (d_1847_v6_)})
            d_1849_v8_: _dafny.Seq
            d_1849_v8_ = _dafny.SeqWithoutIsStrInference([-557, d_1836_cf47_])
            d_1848_v7_ = (d_1848_v7_).set((self).f24, d_1849_v8_)
            d_1850_v9_: D12
            d_1850_v9_ = D12_DC25((self).f24, (D1_DC4(self.f23)).cf6)
            if ((self).f24 if d_1838_cf45_ else (d_1850_v9_).cf41):
                d_1850_v9_ = d_1850_v9_
                d_1851_v10_: _dafny.Map
                d_1851_v10_ = _dafny.Map({d_1838_cf45_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sworc"))})
                d_1852_v11_: T0
                nw307_ = C0()
                nw307_.ctor__(d_1851_v10_, d_1838_cf45_, self.f23, d_1838_cf45_)
                d_1852_v11_ = nw307_
                d_1853_v12_: T0
                d_1853_v12_ = d_1852_v11_
                d_1854_v13_: _dafny.Array
                nw308_ = _dafny.Array(None, 14)
                nw308_[int(0)] = (d_1853_v12_)
                nw308_[int(1)] = d_1852_v11_
                nw308_[int(2)] = d_1852_v11_
                nw308_[int(3)] = d_1852_v11_
                nw308_[int(4)] = d_1852_v11_
                nw308_[int(5)] = (d_1852_v11_ if (d_1852_v11_).f24 else d_1852_v11_)
                nw308_[int(6)] = d_1852_v11_
                nw308_[int(7)] = d_1852_v11_
                nw308_[int(8)] = d_1852_v11_
                nw308_[int(9)] = d_1852_v11_
                nw308_[int(10)] = (d_1852_v11_ if (d_1852_v11_).f24 else d_1852_v11_)
                nw308_[int(11)] = d_1852_v11_
                nw308_[int(12)] = (d_1852_v11_ if (d_1852_v11_).f24 else d_1852_v11_)
                nw308_[int(13)] = d_1852_v11_
                d_1854_v13_ = nw308_
                d_1855_v14_: _dafny.MultiSet
                d_1855_v14_ = _dafny.MultiSet([d_1835_cf48_, (0) - (len(d_1842_v1_))])
                d_1856_v15_: _dafny.Map
                d_1856_v15_ = _dafny.Map({d_1838_cf45_: (D5_DC13((0) - (d_1836_cf47_), not((self).f24), (d_1852_v11_).f24)).cf20})
                d_1857_v16_: _dafny.Map
                d_1857_v16_ = _dafny.Map({len(d_1856_v15_): (self).f24})
                d_1858_v17_: _dafny.Array
                nw309_ = _dafny.Array(None, 19)
                nw309_[int(0)] = d_1837_cf46_
                nw309_[int(1)] = default__.safeModuloInt(len(d_1842_v1_), d_1835_cf48_)
                nw309_[int(2)] = d_1835_cf48_
                nw309_[int(3)] = (self).fm6(globalState)
                nw309_[int(4)] = d_1835_cf48_
                nw309_[int(5)] = d_1836_cf47_
                nw309_[int(6)] = default__.safeModuloInt(d_1835_cf48_, len(d_1842_v1_))
                nw309_[int(7)] = 310
                nw309_[int(8)] = 592
                nw309_[int(9)] = default__.safeDivisionInt(((d_1855_v14_)[d_1837_cf46_] if (d_1837_cf46_) in (d_1855_v14_) else len(d_1857_v16_)), d_1836_cf47_)
                nw309_[int(10)] = default__.safeDivisionInt(d_1837_cf46_, d_1836_cf47_)
                nw309_[int(11)] = len(_dafny.Set({d_1835_cf48_, len(d_1842_v1_), (0) - ((0) - (d_1835_cf48_)), d_1837_cf46_, d_1836_cf47_}))
                nw309_[int(12)] = d_1837_cf46_
                nw309_[int(13)] = d_1836_cf47_
                nw309_[int(14)] = d_1835_cf48_
                nw309_[int(15)] = 132
                nw309_[int(16)] = d_1836_cf47_
                nw309_[int(17)] = d_1837_cf46_
                nw309_[int(18)] = default__.safeModuloInt(d_1835_cf48_, d_1836_cf47_)
                d_1858_v17_ = nw309_
                index249_ = default__.safeIndex(827, (d_1858_v17_).length(0))
                (d_1858_v17_)[index249_] = -225
                d_1859_v18_: _dafny.Array
                nw310_ = _dafny.Array(_dafny.Map({}), 22)
                d_1859_v18_ = nw310_
                index250_ = default__.safeIndex(260, (d_1859_v18_).length(0))
                (d_1859_v18_)[index250_] = default__.fm49(globalState)
                d_1860_v19_: D18
                d_1860_v19_ = D18_DC41(d_1854_v13_)
                d_1861_v20_: _dafny.Seq
                d_1861_v20_ = _dafny.SeqWithoutIsStrInference([(d_1852_v11_).f24])
                index251_ = default__.safeIndex(827, (d_1858_v17_).length(0))
                index252_ = default__.safeIndex(260, (d_1859_v18_).length(0))
                rhs287_ = (0) - (d_1836_cf47_)
                rhs288_ = (d_1860_v19_).cf67
                rhs289_ = default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adnjjpgt"))) + (d_1842_v1_)), d_1835_cf48_)
                rhs290_ = ((self).fm4(d_1836_cf47_, d_1861_v20_, globalState)) and ((self).f24)
                rhs291_ = d_1856_v15_
                lhs210_ = globalState
                lhs211_ = d_1858_v17_
                lhs212_ = default__.safeIndex(827, (d_1858_v17_).length(0))
                lhs213_ = globalState
                lhs214_ = d_1859_v18_
                lhs215_ = default__.safeIndex(260, (d_1859_v18_).length(0))
                lhs210_.f7 = rhs287_
                d_1854_v13_ = rhs288_
                lhs211_[lhs212_] = rhs289_
                lhs213_.f13 = rhs290_
                lhs214_[lhs215_] = rhs291_
                index253_ = default__.safeIndex(827, (d_1858_v17_).length(0))
                (d_1858_v17_)[index253_] = (0) - (d_1837_cf46_)
                d_1862_v21_: C1
                nw311_ = C1()
                nw311_.ctor__((self).f24, (self).f24)
                d_1862_v21_ = nw311_
                d_1863_v22_: C1
                d_1863_v22_ = d_1862_v21_
                d_1863_v22_ = d_1863_v22_
                index254_ = default__.safeIndex(264, (d_1839_v0_).length(0))
                (d_1839_v0_)[index254_] = ((self).f24) and ((self).f24)
                index255_ = default__.safeIndex(264, (d_1839_v0_).length(0))
                (d_1839_v0_)[index255_] = (self).f24
            elif True:
                d_1864_v23_: _dafny.Map
                d_1864_v23_ = _dafny.Map({(self).f24: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjgyr")))})
                d_1865_v24_: _dafny.Set
                d_1865_v24_ = _dafny.Set({(self).f24})
                d_1866_v25_: _dafny.Map
                d_1866_v25_ = _dafny.Map({len(d_1865_v24_): 151})
                d_1864_v23_ = (d_1864_v23_).set((self).f24, ((d_1866_v25_)[d_1836_cf47_] if (d_1836_cf47_) in (d_1866_v25_) else d_1836_cf47_))
                d_1867_v26_: _dafny.Set
                d_1867_v26_ = _dafny.Set({(0) - (d_1835_cf48_)})
                d_1868_v27_: _dafny.Set
                d_1868_v27_ = _dafny.Set({d_1836_cf47_, len(d_1867_v26_)})
                d_1869_v28_: _dafny.Seq
                d_1869_v28_ = _dafny.SeqWithoutIsStrInference([d_1838_cf45_])
                d_1870_v29_: _dafny.Set
                d_1870_v29_ = _dafny.Set({d_1868_v27_})
                d_1871_v30_: _dafny.Map
                d_1871_v30_ = _dafny.Map({d_1869_v28_: len(d_1870_v29_)})
                d_1872_v31_: _dafny.Map
                d_1872_v31_ = _dafny.Map({(d_1842_v1_).set(default__.safeIndex(len(d_1867_v26_), len(d_1842_v1_)), self.f23): d_1871_v30_})
                def iife188_():
                    coll90_ = _dafny.Map()
                    compr_90_: _dafny.Seq
                    for compr_90_ in (_dafny.SeqWithoutIsStrInference([d_1869_v28_, d_1869_v28_, d_1869_v28_])).Elements:
                        d_1873_v32_: _dafny.Seq = compr_90_
                        if (d_1873_v32_) in (_dafny.SeqWithoutIsStrInference([d_1869_v28_, d_1869_v28_, d_1869_v28_])):
                            coll90_[d_1873_v32_] = 66
                    return _dafny.Map(coll90_)
                d_1872_v31_ = (d_1872_v31_).set((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uso"))) + (d_1842_v1_), (iife188_()
                ) | (d_1871_v30_))
                d_1874_v33_: _dafny.MultiSet
                d_1874_v33_ = _dafny.MultiSet([d_1838_cf45_, False])
                d_1875_v34_: _dafny.Seq
                d_1875_v34_ = _dafny.SeqWithoutIsStrInference([d_1844_v3_])
                d_1876_v35_: C2
                nw312_ = C2()
                nw312_.ctor__(d_1874_v33_, (d_1875_v34_)[default__.safeIndex(d_1835_cf48_, len(d_1875_v34_))], self.f23, d_1838_cf45_)
                d_1876_v35_ = nw312_
                d_1838_cf45_ = (self).f24
                (globalState).f13 = (self).f24
            d_1877_v36_: _dafny.Map
            d_1877_v36_ = _dafny.Map({(self).f24: d_1839_v0_})
            (globalState).f7 = len(d_1877_v36_)
        elif source27_.is_DC29:
            d_1878___mcc_h4_ = source27_.cf49
            d_1879___mcc_h5_ = source27_.cf50
            d_1880_cf50_ = d_1879___mcc_h5_
            d_1881_cf49_ = d_1878___mcc_h4_
            d_1882_v37_: int
            d_1882_v37_ = -698
            d_1883_v38_: _dafny.Array
            nw313_ = _dafny.Array(None, 7)
            nw313_[int(0)] = d_1882_v37_
            nw313_[int(1)] = ((0) - (d_1882_v37_)) * (944)
            nw313_[int(2)] = d_1882_v37_
            nw313_[int(3)] = d_1882_v37_
            nw313_[int(4)] = d_1882_v37_
            nw313_[int(5)] = d_1882_v37_
            nw313_[int(6)] = d_1882_v37_
            d_1883_v38_ = nw313_
            (globalState).f22 = d_1883_v38_
            d_1884_v39_: _dafny.Array
            def lambda99_(d_1885_i1_):
                return False

            init56_ = lambda99_
            nw314_ = _dafny.Array(None, 16)
            for i0_56_ in range(nw314_.length(0)):
                nw314_[i0_56_] = init56_(i0_56_)
            d_1884_v39_ = nw314_
            index256_ = default__.safeIndex(747, (d_1884_v39_).length(0))
            (d_1884_v39_)[index256_] = (self).f24
            index257_ = default__.safeIndex(747, (d_1884_v39_).length(0))
            (d_1884_v39_)[index257_] = d_1881_cf49_
            d_1886_v40_: _dafny.Array
            def lambda100_(d_1887_cf50_):
                def lambda101_(d_1888_i2_):
                    return _dafny.Map({d_1887_cf50_: 736})

                return lambda101_

            init57_ = lambda100_(d_1880_cf50_)
            nw315_ = _dafny.Array(None, 7)
            for i0_57_ in range(nw315_.length(0)):
                nw315_[i0_57_] = init57_(i0_57_)
            d_1886_v40_ = nw315_
            d_1886_v40_ = d_1886_v40_
            (globalState).f13 = (self).f24
        elif True:
            d_1889___mcc_h6_ = source27_.cf44
            d_1890_cf44_ = d_1889___mcc_h6_
            d_1891_v41_: _dafny.Seq
            d_1891_v41_ = _dafny.SeqWithoutIsStrInference([not((self).f24)])
            d_1892_v42_: _dafny.Set
            d_1892_v42_ = _dafny.Set({(self).f24})
            d_1893_v43_: _dafny.Map
            d_1893_v43_ = _dafny.Map({d_1891_v41_: d_1892_v42_})
            d_1894_v44_: int
            d_1894_v44_ = 35
            d_1895_v45_: _dafny.Seq
            d_1895_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tghfl"))
            rhs292_ = d_1893_v43_
            rhs293_ = d_1894_v44_
            rhs294_ = (self.f23) in (d_1895_v45_)
            rhs295_ = not((self).f24)
            lhs216_ = globalState
            lhs217_ = globalState
            d_1893_v43_ = rhs292_
            r1 = rhs293_
            lhs216_.f13 = rhs294_
            lhs217_.f13 = rhs295_
            d_1896_v46_: T1
            nw316_ = C6()
            nw316_.ctor__(self.f23, (self).f24)
            d_1896_v46_ = nw316_
            d_1897_v47_: _dafny.Seq
            d_1897_v47_ = _dafny.SeqWithoutIsStrInference([d_1896_v46_, d_1896_v46_])
            d_1897_v47_ = ((d_1897_v47_) + (_dafny.SeqWithoutIsStrInference([d_1896_v46_]))) + (d_1897_v47_)
            if ((d_1896_v46_).f24) or ((self).f24):
                (globalState).f7 = d_1894_v44_
                d_1898_v48_: _dafny.Array
                nw317_ = _dafny.Array(int(0), 12)
                d_1898_v48_ = nw317_
                index258_ = default__.safeIndex(346, (d_1898_v48_).length(0))
                (d_1898_v48_)[index258_] = (0) - (len(d_1891_v41_))
                d_1899_v49_: _dafny.Map
                d_1899_v49_ = _dafny.Map({d_1894_v44_: (self).f24})
                index259_ = default__.safeIndex(346, (d_1898_v48_).length(0))
                (d_1898_v48_)[index259_] = len((d_1899_v49_) | ((d_1899_v49_ if (self).f24 else d_1899_v49_)))
                d_1900_v50_: _dafny.MultiSet
                d_1900_v50_ = _dafny.MultiSet([_dafny.Map({(d_1896_v46_).f24: (d_1896_v46_).f24})])
                d_1901_v51_: _dafny.MultiSet
                d_1901_v51_ = _dafny.MultiSet([(self).f24])
                d_1902_v52_: _dafny.Set
                d_1902_v52_ = _dafny.Set({d_1894_v44_})
                d_1903_v53_: int
                d_1904_v54_: bool
                d_1905_v55_: int
                out52_: int
                out53_: bool
                out54_: int
                out52_, out53_, out54_ = default__.m0((d_1900_v50_).cardinality, (d_1901_v51_).set(True, default__.abs(len(d_1902_v52_))), (d_1896_v46_).f24, d_1894_v44_, globalState)
                d_1903_v53_ = out52_
                d_1904_v54_ = out53_
                d_1905_v55_ = out54_
                d_1906_v56_: _dafny.Array
                nw318_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
                d_1906_v56_ = nw318_
                nw319_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
                d_1906_v56_ = nw319_
                d_1891_v41_ = d_1891_v41_
            elif True:
                d_1894_v44_ = default__.safeDivisionInt(726, d_1894_v44_)
                d_1907_v57_: _dafny.Map
                d_1907_v57_ = _dafny.Map({(d_1896_v46_).f24: d_1891_v41_})
                d_1907_v57_ = (d_1907_v57_).set((d_1896_v46_).f24, d_1891_v41_)
                (globalState).f13 = (d_1896_v46_).f24
                d_1908_v58_: _dafny.Map
                d_1908_v58_ = _dafny.Map({d_1894_v44_: (d_1896_v46_).f24})
                d_1909_v59_: _dafny.Seq
                d_1909_v59_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(d_1894_v44_, d_1894_v44_), len(d_1908_v58_)])
                rhs296_ = (d_1894_v44_) <= ((0) - (d_1894_v44_))
                rhs297_ = (d_1894_v44_) * ((len(d_1891_v41_)) - (d_1894_v44_))
                rhs298_ = (_dafny.SeqWithoutIsStrInference([d_1894_v44_, d_1894_v44_])) + (d_1909_v59_)
                rhs299_ = (self).f24
                lhs218_ = globalState
                lhs219_ = globalState
                r2 = rhs296_
                lhs218_.f7 = rhs297_
                d_1909_v59_ = rhs298_
                lhs219_.f13 = rhs299_
                d_1910_v60_: D14
                d_1910_v60_ = D14_DC31((d_1896_v46_).f24)
                r2 = (d_1910_v60_).cf52
            d_1911_v61_: _dafny.Array
            nw320_ = _dafny.Array(False, 25)
            d_1911_v61_ = nw320_
            d_1912_v62_: _dafny.Map
            d_1912_v62_ = _dafny.Map({d_1911_v61_: d_1896_v46_.f23})
            d_1913_v63_: _dafny.Array
            def lambda102_(d_1914_i3_):
                return _dafny.CodePoint('v')

            init58_ = lambda102_
            nw321_ = _dafny.Array(None, 17)
            for i0_58_ in range(nw321_.length(0)):
                nw321_[i0_58_] = init58_(i0_58_)
            d_1913_v63_ = nw321_
            d_1915_v65_: C3
            nw322_ = C3()
            def iife189_():
                coll91_ = _dafny.Set()
                compr_91_: int
                for compr_91_ in _dafny.IntegerRange(164, 617):
                    d_1916_v64_: int = compr_91_
                    if ((164) <= (d_1916_v64_)) and ((d_1916_v64_) < (617)):
                        coll91_ = coll91_.union(_dafny.Set([(d_1916_v64_) - (d_1894_v44_)]))
                return _dafny.Set(coll91_)
            nw322_.ctor__(d_1912_v62_, d_1913_v63_, (d_1895_v45_)[default__.safeIndex(d_1894_v44_, len(d_1895_v45_))], (_dafny.Set({d_1894_v44_})) == (iife189_()
            ))
            d_1915_v65_ = nw322_
        d_1917_v66_: int
        d_1917_v66_ = -734
        d_1918_v67_: _dafny.Seq
        d_1918_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfjo"))
        d_1919_v68_: _dafny.Set
        d_1919_v68_ = _dafny.Set({(self).fm6(globalState), d_1917_v66_, d_1917_v66_})
        d_1920_v69_: _dafny.Map
        d_1920_v69_ = _dafny.Map({d_1917_v66_: self.f23})
        d_1921_v70_: _dafny.Array
        nw323_ = _dafny.Array(int(0), 3)
        d_1921_v70_ = nw323_
        d_1922_v71_: _dafny.Seq
        d_1922_v71_ = _dafny.SeqWithoutIsStrInference([d_1921_v70_])
        d_1923_v72_: _dafny.Seq
        d_1923_v72_ = _dafny.SeqWithoutIsStrInference([d_1920_v69_, d_1920_v69_, d_1920_v69_, _dafny.Map({len(d_1922_v71_): self.f23})])
        d_1924_v73_: _dafny.MultiSet
        d_1924_v73_ = _dafny.MultiSet([(0) - (d_1917_v66_), d_1917_v66_])
        d_1925_v74_: _dafny.Array
        nw324_ = _dafny.Array(None, 21)
        nw324_[int(0)] = d_1917_v66_
        nw324_[int(1)] = d_1917_v66_
        nw324_[int(2)] = (d_1917_v66_) * (len(d_1918_v67_))
        nw324_[int(3)] = d_1917_v66_
        nw324_[int(4)] = (d_1917_v66_) - (434)
        nw324_[int(5)] = d_1917_v66_
        nw324_[int(6)] = d_1917_v66_
        nw324_[int(7)] = len(d_1918_v67_)
        nw324_[int(8)] = len(d_1919_v68_)
        nw324_[int(9)] = d_1917_v66_
        nw324_[int(10)] = (d_1917_v66_) - (d_1917_v66_)
        nw324_[int(11)] = d_1917_v66_
        nw324_[int(12)] = (d_1917_v66_) - (len(d_1923_v72_))
        nw324_[int(13)] = d_1917_v66_
        nw324_[int(14)] = d_1917_v66_
        nw324_[int(15)] = d_1917_v66_
        nw324_[int(16)] = (d_1917_v66_) + (d_1917_v66_)
        nw324_[int(17)] = d_1917_v66_
        nw324_[int(18)] = default__.safeModuloInt((d_1924_v73_).cardinality, d_1917_v66_)
        nw324_[int(19)] = d_1917_v66_
        nw324_[int(20)] = 139
        d_1925_v74_ = nw324_
        d_1926_v75_: _dafny.Map
        d_1926_v75_ = _dafny.Map({d_1917_v66_: 904})
        d_1927_v76_: _dafny.Map
        d_1927_v76_ = d_1926_v75_
        d_1928_v77_: _dafny.Array
        def lambda103_(d_1929_i4_):
            return self.f23

        init59_ = lambda103_
        nw325_ = _dafny.Array(None, 23)
        for i0_59_ in range(nw325_.length(0)):
            nw325_[i0_59_] = init59_(i0_59_)
        d_1928_v77_ = nw325_
        d_1930_v78_: T2
        nw326_ = C4()
        nw326_.ctor__(872, d_1927_v76_, d_1928_v77_, self.f23, not((self).f24))
        d_1930_v78_ = nw326_
        d_1931_v79_: _dafny.Seq
        d_1931_v79_ = _dafny.SeqWithoutIsStrInference([d_1930_v78_, d_1930_v78_, d_1930_v78_, d_1930_v78_, d_1930_v78_])
        d_1932_v80_: _dafny.Seq
        d_1932_v80_ = _dafny.SeqWithoutIsStrInference([(d_1931_v79_)[default__.safeIndex(d_1917_v66_, len(d_1931_v79_))], d_1930_v78_])
        d_1933_v81_: _dafny.Map
        d_1933_v81_ = _dafny.Map({(self).f24: d_1932_v80_})
        index260_ = default__.safeIndex(906, (d_1925_v74_).length(0))
        (d_1925_v74_)[index260_] = len(((d_1933_v81_)[(d_1930_v78_).f24] if ((d_1930_v78_).f24) in (d_1933_v81_) else d_1932_v80_))
        d_1934_v82_: _dafny.Map
        d_1934_v82_ = _dafny.Map({(d_1930_v78_).f24: len(_dafny.SeqWithoutIsStrInference([(self).f24, (self).f24]))})
        index261_ = default__.safeIndex(906, (d_1925_v74_).length(0))
        (d_1925_v74_)[index261_] = default__.safeDivisionInt(d_1917_v66_, ((d_1934_v82_)[(d_1930_v78_).f24] if ((d_1930_v78_).f24) in (d_1934_v82_) else -751))
        r1 = (d_1925_v74_)[default__.safeIndex(906, (d_1925_v74_).length(0))]
        d_1935_v83_: _dafny.Map
        d_1935_v83_ = _dafny.Map({(d_1930_v78_).f25: ((d_1925_v74_)[default__.safeIndex(906, (d_1925_v74_).length(0))]) - ((d_1925_v74_)[default__.safeIndex(906, (d_1925_v74_).length(0))])})
        d_1935_v83_ = (d_1935_v83_).set(d_1928_v77_, d_1917_v66_)
        rhs300_ = d_1930_v78_.f23
        rhs301_ = default__.fm34(default__.safeModuloInt(len(d_1919_v68_), d_1917_v66_), (d_1930_v78_).f24, globalState)
        lhs220_ = d_1930_v78_
        lhs221_ = self
        lhs220_.f23 = rhs300_
        lhs221_.f23 = rhs301_
        d_1936_v84_: _dafny.Array
        def lambda104_(d_1937_v68_):
            def lambda105_(d_1938_i5_):
                return d_1937_v68_

            return lambda105_

        init60_ = lambda104_(d_1919_v68_)
        nw327_ = _dafny.Array(None, 18)
        for i0_60_ in range(nw327_.length(0)):
            nw327_[i0_60_] = init60_(i0_60_)
        d_1936_v84_ = nw327_
        index262_ = default__.safeIndex(676, (d_1936_v84_).length(0))
        (d_1936_v84_)[index262_] = (D5_DC12(d_1919_v68_)).cf19
        index263_ = default__.safeIndex(676, (d_1936_v84_).length(0))
        (d_1936_v84_)[index263_] = d_1919_v68_
        r0 = ((d_1918_v67_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebhil")))) <= (d_1918_v67_)
        r1 = d_1917_v66_
        r2 = (D15_DC34(True)).cf55
        return r0, r1, r2

    def m5(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        d_1939_v0_: _dafny.Seq
        d_1939_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnuau"))
        hi12_ = default__.safeDivisionInt(p3, (0) - (len(d_1939_v0_)))
        for d_1940_i0_ in range(p3, hi12_):
            if False:
                d_1939_v0_ = d_1939_v0_
                (globalState).f13 = p2
                (globalState).f7 = (p3 if p2 else (d_1940_i0_ if (self).f24 else d_1940_i0_))
                d_1941_v1_: _dafny.Array
                def lambda106_(d_1942_i1_):
                    return (d_1942_i1_) + (930)

                init61_ = lambda106_
                nw328_ = _dafny.Array(None, 16)
                for i0_61_ in range(nw328_.length(0)):
                    nw328_[i0_61_] = init61_(i0_61_)
                d_1941_v1_ = nw328_
                index264_ = default__.safeIndex(505, (d_1941_v1_).length(0))
                (d_1941_v1_)[index264_] = d_1940_i0_
                index265_ = default__.safeIndex(505, (d_1941_v1_).length(0))
                (d_1941_v1_)[index265_] = default__.safeDivisionInt(p3, p3)
                d_1943_v2_: _dafny.Seq
                d_1943_v2_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1944_v3_: C1
                nw329_ = C1()
                nw329_.ctor__((self).fm4(475, d_1943_v2_, globalState), p2)
                d_1944_v3_ = nw329_
            elif True:
                rhs302_ = (self).f24
                rhs303_ = not (p2) or ((p1 if p2 else True))
                lhs222_ = globalState
                lhs223_ = globalState
                lhs222_.f13 = rhs302_
                lhs223_.f13 = rhs303_
                d_1945_v4_: _dafny.Set
                d_1945_v4_ = _dafny.Set({(0) - (d_1940_i0_), p3})
                d_1945_v4_ = d_1945_v4_
                d_1946_v5_: _dafny.MultiSet
                d_1946_v5_ = _dafny.MultiSet([d_1940_i0_, default__.fm0((self).f24, p3, globalState)])
                d_1947_v6_: _dafny.Seq
                d_1947_v6_ = _dafny.SeqWithoutIsStrInference([p2, (self).f24])
                d_1948_v7_: _dafny.Map
                d_1948_v7_ = _dafny.Map({d_1939_v0_: (self).f24})
                d_1949_v8_: _dafny.Map
                d_1949_v8_ = _dafny.Map({(d_1946_v5_).isdisjoint((d_1946_v5_).set(len(d_1947_v6_), default__.abs(p3))): len(d_1948_v7_)})
                d_1949_v8_ = d_1949_v8_
                (globalState).f13 = False
                d_1950_v9_: _dafny.Array
                def lambda107_(d_1951_v6_):
                    def lambda108_(d_1952_i2_):
                        return (d_1952_i2_) * (len(d_1951_v6_))

                    return lambda108_

                init62_ = lambda107_(d_1947_v6_)
                nw330_ = _dafny.Array(None, 10)
                for i0_62_ in range(nw330_.length(0)):
                    nw330_[i0_62_] = init62_(i0_62_)
                d_1950_v9_ = nw330_
                index266_ = default__.safeIndex(644, (d_1950_v9_).length(0))
                (d_1950_v9_)[index266_] = p3
                index267_ = default__.safeIndex(644, (d_1950_v9_).length(0))
                (d_1950_v9_)[index267_] = default__.safeModuloInt(d_1940_i0_, p3)
            d_1953_v10_: _dafny.Map
            d_1953_v10_ = _dafny.Map({(self).f24: len(_dafny.SeqWithoutIsStrInference([635 for d_1954_i3_ in range(default__.abs(445))]))})
            d_1955_v11_: _dafny.MultiSet
            d_1955_v11_ = _dafny.MultiSet([d_1940_i0_, (0) - (((d_1953_v10_)[p2] if (p2) in (d_1953_v10_) else d_1940_i0_)), d_1940_i0_])
            d_1956_v12_: _dafny.MultiSet
            d_1956_v12_ = _dafny.MultiSet([d_1939_v0_])
            (globalState).f13 = not ((d_1955_v11_).issubset(d_1955_v11_)) or ((_dafny.MultiSet([d_1939_v0_])).issubset(d_1956_v12_))
            d_1957_v13_: _dafny.Array
            nw331_ = _dafny.Array(int(0), 28)
            d_1957_v13_ = nw331_
            (globalState).f22 = d_1957_v13_
            d_1958_v14_: _dafny.Array
            nw332_ = _dafny.Array(_dafny.Array(None, 0), 25)
            d_1958_v14_ = nw332_
            index268_ = default__.safeIndex(126, (d_1958_v14_).length(0))
            (d_1958_v14_)[index268_] = d_1957_v13_
            index269_ = default__.safeIndex(126, (d_1958_v14_).length(0))
            (d_1958_v14_)[index269_] = (d_1957_v13_ if p1 else d_1957_v13_)
        d_1959_v15_: _dafny.Array
        def lambda109_(d_1960_p1_):
            def lambda110_(d_1961_i4_):
                return d_1960_p1_

            return lambda110_

        init63_ = lambda109_(p1)
        nw333_ = _dafny.Array(None, 18)
        for i0_63_ in range(nw333_.length(0)):
            nw333_[i0_63_] = init63_(i0_63_)
        d_1959_v15_ = nw333_
        index270_ = default__.safeIndex(159, (d_1959_v15_).length(0))
        (d_1959_v15_)[index270_] = p2
        index271_ = default__.safeIndex(159, (d_1959_v15_).length(0))
        (d_1959_v15_)[index271_] = (self).f24
        d_1962_v16_: _dafny.Array
        nw334_ = _dafny.Array(None, 5)
        nw334_[int(0)] = self.f23
        nw334_[int(1)] = (d_1939_v0_)[default__.safeIndex(p3, len(d_1939_v0_))]
        nw334_[int(2)] = self.f23
        nw334_[int(3)] = default__.fm34(p3, (self).f24, globalState)
        nw334_[int(4)] = self.f23
        d_1962_v16_ = nw334_
        source28_ = D12_DC24(d_1962_v16_)
        if source28_.is_DC25:
            d_1963___mcc_h0_ = source28_.cf41
            d_1964___mcc_h1_ = source28_.cf42
            d_1965_cf42_ = d_1964___mcc_h1_
            d_1966_cf41_ = d_1963___mcc_h0_
            (globalState).f13 = p1
            r0 = d_1966_cf41_
            d_1967_v17_: _dafny.Set
            d_1967_v17_ = _dafny.Set({p1})
            if (_dafny.Set({d_1966_cf41_})) == (d_1967_v17_):
                d_1968_v18_: _dafny.Array
                nw335_ = _dafny.Array(D18.default()(), 5)
                d_1968_v18_ = nw335_
                d_1969_v19_: T0
                nw336_ = C0()
                nw336_.ctor__(_dafny.Map({p2: d_1939_v0_}), d_1966_cf41_, self.f23, d_1966_cf41_)
                d_1969_v19_ = nw336_
                d_1970_v20_: _dafny.Map
                d_1970_v20_ = _dafny.Map({True: d_1969_v19_})
                d_1971_v21_: _dafny.Map
                d_1971_v21_ = _dafny.Map({p3: d_1969_v19_})
                d_1972_v22_: _dafny.Map
                d_1972_v22_ = _dafny.Map({d_1969_v19_.f23: d_1969_v19_})
                d_1973_v23_: _dafny.Array
                nw337_ = _dafny.Array(None, 13)
                nw337_[int(0)] = d_1969_v19_
                nw337_[int(1)] = ((d_1970_v20_)[(self).f24] if ((self).f24) in (d_1970_v20_) else d_1969_v19_)
                nw337_[int(2)] = ((d_1971_v21_)[p3] if (p3) in (d_1971_v21_) else d_1969_v19_)
                nw337_[int(3)] = d_1969_v19_
                nw337_[int(4)] = d_1969_v19_
                nw337_[int(5)] = d_1969_v19_
                nw337_[int(6)] = d_1969_v19_
                nw337_[int(7)] = d_1969_v19_
                nw337_[int(8)] = ((d_1972_v22_)[d_1969_v19_.f23] if (d_1969_v19_.f23) in (d_1972_v22_) else d_1969_v19_)
                nw337_[int(9)] = d_1969_v19_
                nw337_[int(10)] = d_1969_v19_
                nw337_[int(11)] = d_1969_v19_
                nw337_[int(12)] = d_1969_v19_
                d_1973_v23_ = nw337_
                pat_let_tv33_ = d_1973_v23_
                index272_ = default__.safeIndex(313, (d_1968_v18_).length(0))
                def iife190_(_pat_let49_0):
                    def iife191_(d_1974_dt__update__tmp_h0_):
                        def iife192_(_pat_let50_0):
                            def iife193_(d_1975_dt__update_hcf67_h0_):
                                return D18_DC41(d_1975_dt__update_hcf67_h0_)
                            return iife193_(_pat_let50_0)
                        return iife192_(pat_let_tv33_)
                    return iife191_(_pat_let49_0)
                (d_1968_v18_)[index272_] = iife190_(D18_DC41(d_1973_v23_))
                d_1976_v24_: D18
                d_1976_v24_ = D18_DC41(d_1973_v23_)
                pat_let_tv34_ = d_1973_v23_
                index273_ = default__.safeIndex(313, (d_1968_v18_).length(0))
                def iife194_(_pat_let51_0):
                    def iife195_(d_1977_dt__update__tmp_h1_):
                        def iife196_(_pat_let52_0):
                            def iife197_(d_1978_dt__update_hcf67_h1_):
                                return D18_DC41(d_1978_dt__update_hcf67_h1_)
                            return iife197_(_pat_let52_0)
                        return iife196_(pat_let_tv34_)
                    return iife195_(_pat_let51_0)
                (d_1968_v18_)[index273_] = iife194_(d_1976_v24_)
                d_1979_v25_: _dafny.Map
                d_1979_v25_ = _dafny.Map({p3: d_1966_cf41_})
                d_1980_v26_: _dafny.Map
                d_1980_v26_ = _dafny.Map({len(d_1979_v25_): p1})
                d_1981_v27_: C8
                nw338_ = C8()
                nw338_.ctor__((d_1979_v25_) == (d_1979_v25_))
                d_1981_v27_ = nw338_
                d_1939_v0_ = (d_1939_v0_ if not(d_1966_cf41_) else d_1939_v0_)
                (globalState).f13 = (self).f24
                d_1982_v28_: _dafny.MultiSet
                d_1982_v28_ = _dafny.MultiSet([(d_1969_v19_).f24, False, (p3) < (p3), (d_1969_v19_).f24])
                def iife198_():
                    coll92_ = _dafny.Set()
                    compr_92_: int
                    for compr_92_ in _dafny.IntegerRange(767, 828):
                        d_1983_v29_: int = compr_92_
                        if ((767) <= (d_1983_v29_)) and ((d_1983_v29_) < (828)):
                            coll92_ = coll92_.union(_dafny.Set([default__.safeDivisionInt(d_1983_v29_, p3)]))
                    return _dafny.Set(coll92_)
                def iife199_():
                    coll93_ = _dafny.Set()
                    compr_93_: int
                    for compr_93_ in _dafny.IntegerRange(529, 499):
                        d_1984_v30_: int = compr_93_
                        if ((529) <= (d_1984_v30_)) and ((d_1984_v30_) < (499)):
                            coll93_ = coll93_.union(_dafny.Set([(d_1984_v30_) - (p3)]))
                    return _dafny.Set(coll93_)
                d_1982_v28_ = _dafny.MultiSet([(iife198_()
                ).isdisjoint(iife199_()
                )])
            elif True:
                d_1985_v31_: _dafny.Map
                d_1985_v31_ = _dafny.Map({p3: p3})
                d_1986_v32_: C4
                nw339_ = C4()
                nw339_.ctor__(p3, d_1985_v31_, d_1962_v16_, self.f23, not(p1))
                d_1986_v32_ = nw339_
                d_1987_v33_: _dafny.Array
                nw340_ = _dafny.Array(_dafny.MultiSet({}), 10)
                d_1987_v33_ = nw340_
                d_1988_v34_: C7
                nw341_ = C7()
                nw341_.ctor__((self).f24, (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))], d_1962_v16_, d_1965_cf42_, p1, p1)
                d_1988_v34_ = nw341_
                d_1989_v35_: _dafny.MultiSet
                d_1989_v35_ = _dafny.MultiSet([d_1988_v34_])
                index274_ = default__.safeIndex(944, (d_1987_v33_).length(0))
                (d_1987_v33_)[index274_] = d_1989_v35_
                d_1990_v36_: _dafny.Seq
                d_1990_v36_ = _dafny.SeqWithoutIsStrInference([p3])
                d_1991_v37_: _dafny.MultiSet
                d_1991_v37_ = _dafny.MultiSet([905, len(_dafny.SeqWithoutIsStrInference([d_1990_v36_])), len(d_1939_v0_)])
                d_1992_v38_: D15
                d_1992_v38_ = D15_DC34(p1)
                d_1993_v39_: _dafny.Seq
                d_1993_v39_ = _dafny.SeqWithoutIsStrInference([(d_1988_v34_).f31, p2])
                index275_ = default__.safeIndex(944, (d_1987_v33_).length(0))
                rhs304_ = d_1967_v17_
                rhs305_ = ((_dafny.MultiSet([d_1988_v34_])).intersection(d_1989_v35_)) - (d_1989_v35_)
                rhs306_ = _dafny.MultiSet([(d_1986_v32_).f33, ((0) - (p3)) + ((d_1986_v32_).f33), ((_dafny.MultiSet([d_1992_v38_])).cardinality) - (len(d_1993_v39_))])
                lhs224_ = d_1987_v33_
                lhs225_ = default__.safeIndex(944, (d_1987_v33_).length(0))
                d_1967_v17_ = rhs304_
                lhs224_[lhs225_] = rhs305_
                d_1991_v37_ = rhs306_
                d_1994_v40_: _dafny.Map
                d_1994_v40_ = _dafny.Map({d_1966_cf41_: (d_1988_v34_).f31})
                (globalState).f13 = ((d_1994_v40_)[d_1966_cf41_] if (d_1966_cf41_) in (d_1994_v40_) else d_1966_cf41_)
                d_1995_v41_: _dafny.Set
                d_1995_v41_ = _dafny.Set({(d_1986_v32_).f33})
                (globalState).f13 = ((_dafny.Set({(d_1986_v32_).f33})) | (d_1995_v41_)).issubset(d_1995_v41_)
                d_1996_v42_: _dafny.Array
                def lambda111_(d_1997_p2_):
                    def lambda112_(d_1998_i5_):
                        return _dafny.Map({d_1997_p2_: self.f23})

                    return lambda112_

                init64_ = lambda111_(p2)
                nw342_ = _dafny.Array(None, 20)
                for i0_64_ in range(nw342_.length(0)):
                    nw342_[i0_64_] = init64_(i0_64_)
                d_1996_v42_ = nw342_
                d_1999_v43_: D14
                d_1999_v43_ = D14_DC30(d_1996_v42_)
                d_2000_v44_: _dafny.Map
                d_2000_v44_ = _dafny.Map({(d_1986_v32_).f33: d_1999_v43_})
                d_2000_v44_ = _dafny.Map({(p3) + (p3): D14_DC30(d_1996_v42_)})
            d_2001_v45_: D12
            d_2001_v45_ = D12_DC25((p3) >= (p3), d_1965_cf42_)
            d_2001_v45_ = d_2001_v45_
        elif source28_.is_DC24:
            d_2002___mcc_h2_ = source28_.cf40
            d_2003_cf40_ = d_2002___mcc_h2_
            d_2004_v46_: _dafny.Map
            d_2004_v46_ = _dafny.Map({p3: default__.fm2(globalState)})
            d_2005_v47_: _dafny.Set
            d_2005_v47_ = _dafny.Set({((d_2004_v46_)[p3] if (p3) in (d_2004_v46_) else (self).f24), (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]})
            d_2006_v48_: _dafny.Map
            d_2006_v48_ = _dafny.Map({p3: (d_2005_v47_) | (d_2005_v47_)})
            d_2006_v48_ = (d_2006_v48_).set(p3, _dafny.Set({p1, False}))
            if p1:
                d_2007_v49_: D12
                d_2007_v49_ = D12_DC25((self).f24, self.f23)
                index276_ = default__.safeIndex(12, (d_1962_v16_).length(0))
                (d_1962_v16_)[index276_] = (d_2007_v49_).cf42
                index277_ = default__.safeIndex(12, (d_1962_v16_).length(0))
                rhs307_ = self.f23
                rhs308_ = not ((self).f24) or (not((self).f24))
                rhs309_ = p2
                lhs226_ = d_1962_v16_
                lhs227_ = default__.safeIndex(12, (d_1962_v16_).length(0))
                lhs228_ = globalState
                lhs229_ = globalState
                lhs226_[lhs227_] = rhs307_
                lhs228_.f13 = rhs308_
                lhs229_.f13 = rhs309_
                (globalState).f7 = default__.safeModuloInt((p3) + ((0) - (p3)), (p3 if (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))] else p3))
                (globalState).f13 = p2
                d_2008_v50_: _dafny.Array
                nw343_ = _dafny.Array(int(0), 18)
                d_2008_v50_ = nw343_
                index278_ = default__.safeIndex(357, (d_2008_v50_).length(0))
                (d_2008_v50_)[index278_] = p3
                index279_ = default__.safeIndex(357, (d_2008_v50_).length(0))
                (d_2008_v50_)[index279_] = (0) - (p3)
                (globalState).f7 = p3
            elif True:
                d_2009_v51_: _dafny.Map
                d_2009_v51_ = _dafny.Map({p2: p3})
                (globalState).f7 = default__.fm0((d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))], ((d_2009_v51_)[p2] if (p2) in (d_2009_v51_) else len(d_2004_v46_)), globalState)
                d_2010_v52_: _dafny.Array
                nw344_ = _dafny.Array(_dafny.Map({}), 13)
                d_2010_v52_ = nw344_
                d_2011_v53_: D15
                d_2011_v53_ = D15_DC34((self).f24)
                d_2012_v54_: _dafny.Map
                d_2012_v54_ = _dafny.Map({(self).f24: d_2011_v53_})
                index280_ = default__.safeIndex(312, (d_2010_v52_).length(0))
                (d_2010_v52_)[index280_] = d_2012_v54_
                index281_ = default__.safeIndex(312, (d_2010_v52_).length(0))
                (d_2010_v52_)[index281_] = d_2012_v54_
                (globalState).f7 = 263
                d_1939_v0_ = (d_1939_v0_) + ((d_1939_v0_).set(default__.safeIndex(p3, len(d_1939_v0_)), _dafny.CodePoint('o')))
                d_2013_v55_: _dafny.Seq
                d_2013_v55_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({not((d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]), p2, p1, (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]})])
                (globalState).f13 = not(((p3) + (p3)) < (len((d_2013_v55_) + (_dafny.SeqWithoutIsStrInference([default__.fm41(globalState)])))))
            (self).f23 = self.f23
            (self).f23 = self.f23
        elif True:
            d_2014___mcc_h3_ = source28_.cf43
            d_2015_cf43_ = d_2014___mcc_h3_
            d_1939_v0_ = d_1939_v0_
            d_2016_v56_: _dafny.Seq
            d_2016_v56_ = _dafny.SeqWithoutIsStrInference([(d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]])
            d_2017_v57_: _dafny.MultiSet
            d_2017_v57_ = _dafny.MultiSet([(self).fm11(-498, p1, globalState), (self).f24])
            d_2018_v58_: _dafny.Seq
            d_2018_v58_ = _dafny.SeqWithoutIsStrInference([p3])
            d_2019_v59_: _dafny.Set
            d_2019_v59_ = _dafny.Set({d_2018_v58_})
            d_2020_v60_: int
            d_2021_v61_: bool
            d_2022_v62_: int
            out55_: int
            out56_: bool
            out57_: int
            out55_, out56_, out57_ = default__.m0(len(d_2016_v56_), d_2017_v57_, not(not((d_2019_v59_).ispropersubset(d_2019_v59_))), (p3) * ((0) - (p3)), globalState)
            d_2020_v60_ = out55_
            d_2021_v61_ = out56_
            d_2022_v62_ = out57_
            d_2023_v63_: _dafny.Set
            d_2023_v63_ = _dafny.Set({p1, p2, (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]})
            d_2024_v64_: C8
            nw345_ = C8()
            nw345_.ctor__((d_2023_v63_) != (d_2023_v63_))
            d_2024_v64_ = nw345_
            d_2024_v64_ = d_2024_v64_
            d_2025_v65_: _dafny.Set
            d_2025_v65_ = _dafny.Set({(d_2020_v60_) * ((d_2018_v58_)[default__.safeIndex(len(d_2023_v63_), len(d_2018_v58_))]), d_2022_v62_, p3, d_2020_v60_, d_2020_v60_})
            d_2025_v65_ = _dafny.Set({p3})
        hi13_ = p3
        for d_2026_i6_ in range((0) - (p3), hi13_):
            if (p1 if (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))] else p1):
                d_2027_v66_: _dafny.Seq
                d_2027_v66_ = _dafny.SeqWithoutIsStrInference([p1, (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))], p1])
                (globalState).f7 = (len(d_2027_v66_)) * (p3)
                d_2028_v67_: _dafny.Map
                d_2028_v67_ = _dafny.Map({d_1959_v15_: p3})
                d_2029_v68_: _dafny.MultiSet
                d_2029_v68_ = _dafny.MultiSet([p2])
                d_2028_v67_ = (d_2028_v67_).set(d_1959_v15_, default__.safeModuloInt(d_2026_i6_, (d_2029_v68_).cardinality))
                index282_ = default__.safeIndex(159, (d_1959_v15_).length(0))
                (d_1959_v15_)[index282_] = (d_2029_v68_) == ((d_2029_v68_) | (d_2029_v68_))
                (globalState).f13 = not (p1) or ((d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))])
                d_2030_v69_: _dafny.Map
                d_2030_v69_ = _dafny.Map({d_1959_v15_: self.f23})
                d_2031_v70_: C3
                nw346_ = C3()
                nw346_.ctor__(d_2030_v69_, d_1962_v16_, self.f23, not ((d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]) or (p2))
                d_2031_v70_ = nw346_
            elif True:
                d_2032_v71_: _dafny.Array
                nw347_ = _dafny.Array(int(0), 21)
                d_2032_v71_ = nw347_
                index283_ = default__.safeIndex(257, (d_2032_v71_).length(0))
                (d_2032_v71_)[index283_] = p3
                index284_ = default__.safeIndex(257, (d_2032_v71_).length(0))
                (d_2032_v71_)[index284_] = len((d_1939_v0_) + (d_1939_v0_))
                d_2033_v72_: _dafny.Map
                d_2033_v72_ = _dafny.Map({d_1959_v15_: self.f23})
                d_2034_v73_: C3
                nw348_ = C3()
                nw348_.ctor__(d_2033_v72_, d_1962_v16_, self.f23, p1)
                d_2034_v73_ = nw348_
                index285_ = default__.safeIndex(257, (d_2032_v71_).length(0))
                (d_2032_v71_)[index285_] = (d_2032_v71_)[default__.safeIndex(257, (d_2032_v71_).length(0))]
                d_2035_v74_: _dafny.MultiSet
                d_2035_v74_ = _dafny.MultiSet([(d_2032_v71_)[default__.safeIndex(257, (d_2032_v71_).length(0))]])
                d_2036_v75_: C1
                nw349_ = C1()
                nw349_.ctor__(True, (_dafny.MultiSet([-913])).ispropersubset(d_2035_v74_))
                d_2036_v75_ = nw349_
                d_2037_v76_: D5
                d_2037_v76_ = D5_DC12(_dafny.Set({p3}))
                d_2038_v77_: _dafny.MultiSet
                d_2038_v77_ = _dafny.MultiSet([d_2037_v76_])
                d_2038_v77_ = d_2038_v77_
            index286_ = default__.safeIndex(247, (d_1962_v16_).length(0))
            (d_1962_v16_)[index286_] = self.f23
            index287_ = default__.safeIndex(247, (d_1962_v16_).length(0))
            rhs310_ = default__.safeDivisionInt(d_2026_i6_, len(((_dafny.SeqWithoutIsStrInference([(d_1939_v0_)[default__.safeIndex(p3, len(d_1939_v0_))] for d_2039_i7_ in range(default__.abs(341))])).set(default__.safeIndex(p3, len(_dafny.SeqWithoutIsStrInference([(d_1939_v0_)[default__.safeIndex(p3, len(d_1939_v0_))] for d_2040_i7_ in range(default__.abs(341))]))), self.f23)) + (d_1939_v0_)))
            rhs311_ = _dafny.CodePoint('n')
            rhs312_ = d_1959_v15_
            rhs313_ = self.f23
            lhs230_ = globalState
            lhs231_ = d_1962_v16_
            lhs232_ = default__.safeIndex(247, (d_1962_v16_).length(0))
            lhs233_ = self
            lhs230_.f7 = rhs310_
            lhs231_[lhs232_] = rhs311_
            d_1959_v15_ = rhs312_
            lhs233_.f23 = rhs313_
            r0 = not(True)
            d_2041_v78_: _dafny.Seq
            d_2041_v78_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2042_v79_: _dafny.Set
            d_2042_v79_ = _dafny.Set({(self).f24, not((self).f24), (self).fm4(-465, d_2041_v78_, globalState)})
            (globalState).f7 = len(d_2042_v79_)
        if p2:
            d_2043_v80_: _dafny.MultiSet
            d_2043_v80_ = _dafny.MultiSet([not(p1)])
            d_2044_v81_: int
            d_2045_v82_: bool
            d_2046_v83_: int
            out58_: int
            out59_: bool
            out60_: int
            out58_, out59_, out60_ = default__.m0((self).fm6(globalState), (default__.fm3(p2, globalState)) | (d_2043_v80_), p2, p3, globalState)
            d_2044_v81_ = out58_
            d_2045_v82_ = out59_
            d_2046_v83_ = out60_
            d_2047_v84_: D13
            d_2047_v84_ = D13_DC27(_dafny.Map({True: self.f23}))
            pat_let_tv35_ = d_2045_v82_
            d_2048_v85_: _dafny.MultiSet
            def iife200_(_pat_let53_0):
                def iife201_(d_2049_dt__update__tmp_h2_):
                    def iife202_(_pat_let54_0):
                        def iife203_(d_2050_dt__update_hcf44_h0_):
                            return D13_DC27(d_2050_dt__update_hcf44_h0_)
                        return iife203_(_pat_let54_0)
                    return iife202_(_dafny.Map({pat_let_tv35_: self.f23}))
                return iife201_(_pat_let53_0)
            d_2048_v85_ = _dafny.MultiSet([d_2047_v84_, (iife200_(d_2047_v84_) if (self).f24 else d_2047_v84_), d_2047_v84_])
            d_2048_v85_ = default__.fm50((d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))], d_2044_v81_, globalState)
            d_2051_v86_: _dafny.Seq
            d_2051_v86_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(d_1939_v0_) for d_2052_i8_ in range(default__.abs(599))]))])
            d_2053_v87_: _dafny.Map
            d_2053_v87_ = _dafny.Map({len((d_2051_v86_).set(default__.safeIndex(d_2044_v81_, len(d_2051_v86_)), d_2046_v83_)): self.f23})
            (self).f23 = (self.f23 if d_2045_v82_ else ((d_2053_v87_)[-794] if (-794) in (d_2053_v87_) else _dafny.CodePoint('b')))
            d_2044_v81_ = p3
            d_2054_v88_: _dafny.Set
            d_2054_v88_ = _dafny.Set({d_2045_v82_})
            (globalState).f7 = (0) - (len((d_2054_v88_) | (d_2054_v88_)))
        elif True:
            (globalState).f7 = 650
            index288_ = default__.safeIndex(159, (d_1959_v15_).length(0))
            (d_1959_v15_)[index288_] = (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]
            d_2055_v89_: D2
            d_2055_v89_ = D2_DC7(not(False), p2)
            d_2056_v90_: _dafny.Seq
            d_2056_v90_ = _dafny.SeqWithoutIsStrInference([d_2055_v89_])
            d_2056_v90_ = (d_2056_v90_) + ((d_2056_v90_) + (_dafny.SeqWithoutIsStrInference([d_2055_v89_])))
            d_2057_v91_: _dafny.MultiSet
            d_2057_v91_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aynidalob")), d_1939_v0_, d_1939_v0_, d_1939_v0_, d_1939_v0_])
            d_2058_v92_: _dafny.Set
            d_2058_v92_ = _dafny.Set({271, ((d_2057_v91_)[d_1939_v0_] if (d_1939_v0_) in (d_2057_v91_) else p3)})
            d_2058_v92_ = d_2058_v92_
            d_2059_v93_: _dafny.Set
            d_2059_v93_ = _dafny.Set({(self).f24})
            d_2060_v94_: _dafny.Map
            d_2060_v94_ = _dafny.Map({p3: len(d_2059_v93_)})
            d_2060_v94_ = d_2060_v94_
        if ((self).f24 if True else (self).f24):
            d_2061_v95_: _dafny.Map
            d_2061_v95_ = _dafny.Map({default__.fm0((d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))], len(d_1939_v0_), globalState): p3})
            d_2062_v96_: _dafny.Set
            d_2062_v96_ = _dafny.Set({d_2061_v95_})
            d_2063_v97_: _dafny.Map
            d_2063_v97_ = _dafny.Map({(d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]: (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]})
            d_2064_v98_: _dafny.Seq
            d_2064_v98_ = _dafny.SeqWithoutIsStrInference([(self).f24, (self).f24])
            d_2065_v99_: C8
            nw350_ = C8()
            nw350_.ctor__(p2)
            d_2065_v99_ = nw350_
            d_2066_v100_: _dafny.Map
            d_2066_v100_ = _dafny.Map({(d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]: d_2065_v99_})
            d_2062_v96_ = default__.fm51(_dafny.Map({d_2063_v97_: d_2064_v98_}), p3, p3, ((d_2064_v98_)[default__.safeIndex(len(d_2066_v100_), len(d_2064_v98_))] if p1 else p1), globalState)
            d_2067_v101_: _dafny.MultiSet
            d_2067_v101_ = _dafny.MultiSet([p2, (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))], (d_2064_v98_)[default__.safeIndex(p3, len(d_2064_v98_))]])
            if (default__.fm0(not((self).f24), p3, globalState)) not in ((_dafny.MultiSet([p3])).set((d_2067_v101_).cardinality, default__.abs(p3))):
                d_2068_v102_: int
                d_2069_v103_: bool
                d_2070_v104_: int
                out61_: int
                out62_: bool
                out63_: int
                out61_, out62_, out63_ = default__.m0(p3, (d_2067_v101_) | (d_2067_v101_), p1, p3, globalState)
                d_2068_v102_ = out61_
                d_2069_v103_ = out62_
                d_2070_v104_ = out63_
                (globalState).f7 = (0) - (len(d_1939_v0_))
                d_2071_v105_: _dafny.Array
                nw351_ = _dafny.Array(None, 23)
                d_2071_v105_ = nw351_
                d_2072_v106_: C6
                nw352_ = C6()
                nw352_.ctor__(self.f23, not(default__.fm2(globalState)))
                d_2072_v106_ = nw352_
                index289_ = default__.safeIndex(635, (d_2071_v105_).length(0))
                (d_2071_v105_)[index289_] = d_2072_v106_
                index290_ = default__.safeIndex(635, (d_2071_v105_).length(0))
                (d_2071_v105_)[index290_] = d_2072_v106_
                d_2073_v107_: D13
                d_2073_v107_ = D13_DC28((d_1939_v0_) < (d_1939_v0_), (0) - ((len(((d_1939_v0_).set(default__.safeIndex(len(d_2064_v98_), len(d_1939_v0_)), self.f23)).set(default__.safeIndex((0) - (d_2068_v102_), len((d_1939_v0_).set(default__.safeIndex(len(d_2064_v98_), len(d_1939_v0_)), self.f23))), self.f23))) * (p3)), (d_2070_v104_) + (920), (d_2070_v104_) * (592))
                rhs314_ = d_2073_v107_
                rhs315_ = not (d_2069_v103_) or ((self).f24)
                lhs234_ = globalState
                d_2073_v107_ = rhs314_
                lhs234_.f13 = rhs315_
                d_2070_v104_ = (D15_DC36(default__.fm2(globalState), len(d_2063_v97_))).cf60
            elif True:
                d_2074_v108_: _dafny.Map
                d_2074_v108_ = _dafny.Map({(self).f24: _dafny.CodePoint('t')})
                d_2075_v109_: C5
                nw353_ = C5()
                nw353_.ctor__(d_1959_v15_)
                d_2075_v109_ = nw353_
                d_2076_v110_: _dafny.Map
                d_2076_v110_ = _dafny.Map({d_2075_v109_: (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]})
                d_2077_v111_: _dafny.Map
                d_2077_v111_ = _dafny.Map({len(d_2076_v110_): p1})
                d_2074_v108_ = default__.fm52(len(d_2077_v111_), (_dafny.SeqWithoutIsStrInference([self.f23])) != (d_1939_v0_), p3, p3, globalState)
                d_2078_v112_: T1
                nw354_ = C6()
                nw354_.ctor__(self.f23, (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))])
                d_2078_v112_ = nw354_
                d_2079_v113_: D5
                d_2079_v113_ = D5_DC13((0) - (default__.safeModuloInt(p3, ((d_2067_v101_)[(self).f24] if ((self).f24) in (d_2067_v101_) else len(default__.fm1((d_2078_v112_).f24, globalState))))), (d_2078_v112_).f24, (d_2078_v112_).f24)
                d_2080_v114_: _dafny.Map
                d_2080_v114_ = _dafny.Map({(d_2078_v112_).f24: 522})
                d_2081_v115_: _dafny.Map
                d_2081_v115_ = _dafny.Map({d_1959_v15_: (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]})
                d_2082_v116_: _dafny.Map
                d_2082_v116_ = _dafny.Map({_dafny.CodePoint('t'): len(d_2081_v115_)})
                d_2083_v117_: C6
                nw355_ = C6()
                nw355_.ctor__(d_2078_v112_.f23, p2)
                d_2083_v117_ = nw355_
                d_2084_v118_: _dafny.Map
                d_2084_v118_ = _dafny.Map({d_2083_v117_: p3})
                d_2085_v119_: _dafny.Array
                nw356_ = _dafny.Array(None, 29)
                nw356_[int(0)] = p3
                nw356_[int(1)] = 565
                nw356_[int(2)] = p3
                nw356_[int(3)] = p3
                nw356_[int(4)] = p3
                nw356_[int(5)] = len((_dafny.Map({p2: self.f23})) | (d_2074_v108_))
                nw356_[int(6)] = default__.safeDivisionInt(216, len(d_2080_v114_))
                nw356_[int(7)] = p3
                nw356_[int(8)] = p3
                nw356_[int(9)] = (p3) * (len(d_1939_v0_))
                nw356_[int(10)] = p3
                nw356_[int(11)] = len((_dafny.Map({self.f23: (self).fm6(globalState)})) | (d_2082_v116_))
                nw356_[int(12)] = p3
                nw356_[int(13)] = p3
                nw356_[int(14)] = ((d_2067_v101_)[(d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]] if ((d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]) in (d_2067_v101_) else 218)
                nw356_[int(15)] = (p3) - (p3)
                nw356_[int(16)] = default__.safeDivisionInt(len(d_2084_v118_), p3)
                nw356_[int(17)] = p3
                nw356_[int(18)] = p3
                nw356_[int(19)] = 463
                nw356_[int(20)] = default__.safeDivisionInt(p3, p3)
                nw356_[int(21)] = len(d_1939_v0_)
                nw356_[int(22)] = default__.safeDivisionInt(287, p3)
                nw356_[int(23)] = -90
                nw356_[int(24)] = len(d_2064_v98_)
                nw356_[int(25)] = p3
                nw356_[int(26)] = p3
                nw356_[int(27)] = 420
                nw356_[int(28)] = p3
                d_2085_v119_ = nw356_
                index291_ = default__.safeIndex(187, (d_2085_v119_).length(0))
                (d_2085_v119_)[index291_] = len(_dafny.SeqWithoutIsStrInference([p3, p3, p3]))
                d_2086_v120_: _dafny.MultiSet
                d_2086_v120_ = _dafny.MultiSet([p3])
                d_2087_v121_: D11
                d_2087_v121_ = D11_DC23((self).f24, (d_2086_v120_).cardinality, d_2063_v97_, D5_DC14((d_2078_v112_).f24))
                pat_let_tv36_ = p3
                index292_ = default__.safeIndex(187, (d_2085_v119_).length(0))
                def iife204_(_pat_let55_0):
                    def iife205_(d_2088_dt__update__tmp_h3_):
                        def iife206_(_pat_let56_0):
                            def iife207_(d_2089_dt__update_hcf20_h0_):
                                return D5_DC13(d_2089_dt__update_hcf20_h0_, (d_2088_dt__update__tmp_h3_).cf21, (d_2088_dt__update__tmp_h3_).cf22)
                            return iife207_(_pat_let56_0)
                        return iife206_(pat_let_tv36_)
                    return iife205_(_pat_let55_0)
                rhs316_ = d_2078_v112_
                rhs317_ = not(((self).f24 if (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))] else p1))
                rhs318_ = iife204_(d_2079_v113_)
                rhs319_ = default__.fm0(p2, (d_2067_v101_).cardinality, globalState)
                rhs320_ = (390 if (d_2087_v121_).cf36 else ((d_2086_v120_)[p3] if (p3) in (d_2086_v120_) else len(d_2080_v114_)))
                lhs235_ = globalState
                lhs236_ = d_2085_v119_
                lhs237_ = default__.safeIndex(187, (d_2085_v119_).length(0))
                lhs238_ = globalState
                d_2078_v112_ = rhs316_
                lhs235_.f13 = rhs317_
                d_2079_v113_ = rhs318_
                lhs236_[lhs237_] = rhs319_
                lhs238_.f7 = rhs320_
                (globalState).f13 = p1
                index293_ = default__.safeIndex(187, (d_2085_v119_).length(0))
                (d_2085_v119_)[index293_] = (d_2078_v112_).fm6(globalState)
                index294_ = default__.safeIndex(159, (d_1959_v15_).length(0))
                (d_1959_v15_)[index294_] = not (not (True) or (p1)) or (not ((d_2078_v112_).f24) or ((d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]))
            if (self).f24:
                d_2090_v122_: C6
                nw357_ = C6()
                nw357_.ctor__(self.f23, p1)
                d_2090_v122_ = nw357_
                d_2091_v123_: D5
                d_2091_v123_ = D5_DC14(p1)
                d_2092_v124_: D11
                d_2092_v124_ = D11_DC23((self).f24, len(d_1939_v0_), d_2063_v97_, d_2091_v123_)
                d_1959_v15_ = (d_1959_v15_ if (d_2092_v124_).cf36 else d_1959_v15_)
                d_2093_v125_: _dafny.Array
                def lambda113_(d_2094_v0_):
                    def lambda114_(d_2095_i9_):
                        return d_2094_v0_

                    return lambda114_

                init65_ = lambda113_(d_1939_v0_)
                nw358_ = _dafny.Array(None, 18)
                for i0_65_ in range(nw358_.length(0)):
                    nw358_[i0_65_] = init65_(i0_65_)
                d_2093_v125_ = nw358_
                d_2093_v125_ = d_2093_v125_
                d_2061_v95_ = (d_2061_v95_).set(p3, ((d_2067_v101_)[False] if (False) in (d_2067_v101_) else p3))
                d_2096_v126_: _dafny.Set
                d_2096_v126_ = _dafny.Set({((d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))]) == ((self).f24)})
                d_2096_v126_ = (d_2096_v126_) | (d_2096_v126_)
            elif True:
                r0 = p1
                (globalState).f13 = p1
                (globalState).f13 = (self.f23) not in (d_1939_v0_)
                r0 = p2
                d_2097_v127_: _dafny.Set
                d_2097_v127_ = _dafny.Set({(d_1939_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykpfixig")))})
                d_2097_v127_ = default__.fm53(globalState)
            d_2063_v97_ = (d_2063_v97_).set(p2, (d_1959_v15_)[default__.safeIndex(159, (d_1959_v15_).length(0))])
            (globalState).f7 = p3
        elif True:
            d_2098_v128_: _dafny.Map
            d_2098_v128_ = _dafny.Map({p3: p3})
            d_2099_v129_: T2
            nw359_ = C4()
            nw359_.ctor__(p3, d_2098_v128_, d_1962_v16_, self.f23, (self).f24)
            d_2099_v129_ = nw359_
            d_2100_v130_: _dafny.Map
            d_2100_v130_ = _dafny.Map({p3: d_2099_v129_})
            d_2101_v131_: _dafny.Seq
            d_2101_v131_ = _dafny.SeqWithoutIsStrInference([p2])
            d_2100_v130_ = (d_2100_v130_).set((len(d_2101_v131_)) + (p3), d_2099_v129_)
            d_2102_v132_: _dafny.MultiSet
            d_2102_v132_ = _dafny.MultiSet([default__.fm2(globalState)])
            index295_ = default__.safeIndex(159, (d_1959_v15_).length(0))
            (d_1959_v15_)[index295_] = (not (p1) or (True)) or ((d_2102_v132_).issubset(d_2102_v132_))
            index296_ = default__.safeIndex(159, (d_1959_v15_).length(0))
            (d_1959_v15_)[index296_] = not((self).f24)
            (globalState).f7 = 57
            (globalState).f13 = p1
        pat_let_tv37_ = p1
        pat_let_tv38_ = p1
        pat_let_tv39_ = d_1959_v15_
        pat_let_tv40_ = d_1959_v15_
        pat_let_tv41_ = p3
        pat_let_tv42_ = p3
        pat_let_tv43_ = p2
        pat_let_tv44_ = d_1939_v0_
        def lambda115_(source29_):
            if source29_.is_DC13:
                d_2103___mcc_h4_ = source29_.cf20
                d_2104___mcc_h5_ = source29_.cf21
                d_2105___mcc_h6_ = source29_.cf22
                d_2106_cf22_ = d_2105___mcc_h6_
                d_2107_cf21_ = d_2104___mcc_h5_
                d_2108_cf20_ = d_2103___mcc_h4_
                return (pat_let_tv37_) == (pat_let_tv38_)
            elif source29_.is_DC14:
                d_2109___mcc_h7_ = source29_.cf23
                d_2110_cf23_ = d_2109___mcc_h7_
                return (pat_let_tv40_)[default__.safeIndex(159, (pat_let_tv39_).length(0))]
            elif True:
                d_2111___mcc_h8_ = source29_.cf19
                d_2112_cf19_ = d_2111___mcc_h8_
                return (_dafny.MultiSet([pat_let_tv41_, pat_let_tv42_])).issubset(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([pat_let_tv43_])), len(pat_let_tv44_), 622]))

        r0 = lambda115_(D5_DC13(len(d_1939_v0_), True, p1))
        return r0

    def m6(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        if p3:
            d_2113_v0_: _dafny.MultiSet
            d_2113_v0_ = _dafny.MultiSet([True])
            d_2114_v1_: _dafny.Map
            d_2114_v1_ = _dafny.Map({d_2113_v0_: self.f23})
            d_2114_v1_ = d_2114_v1_
            d_2115_v2_: int
            d_2115_v2_ = 641
            d_2116_v3_: _dafny.Seq
            d_2116_v3_ = _dafny.SeqWithoutIsStrInference([True])
            d_2117_v4_: _dafny.Map
            d_2117_v4_ = _dafny.Map({d_2115_v2_: d_2115_v2_})
            d_2118_v5_: _dafny.Map
            d_2118_v5_ = _dafny.Map({p0: d_2115_v2_})
            d_2119_v6_: _dafny.Array
            nw360_ = _dafny.Array(None, 10)
            nw360_[int(0)] = d_2117_v4_
            nw360_[int(1)] = d_2117_v4_
            nw360_[int(2)] = _dafny.Map({len(p2): d_2115_v2_})
            nw360_[int(3)] = d_2117_v4_
            nw360_[int(4)] = d_2117_v4_
            nw360_[int(5)] = d_2117_v4_
            nw360_[int(6)] = d_2117_v4_
            nw360_[int(7)] = _dafny.Map({(d_2113_v0_).cardinality: (0) - (len(d_2118_v5_))})
            nw360_[int(8)] = d_2117_v4_
            nw360_[int(9)] = d_2117_v4_
            d_2119_v6_ = nw360_
            d_2120_v7_: _dafny.Map
            d_2120_v7_ = _dafny.Map({d_2119_v6_: (self).fm6(globalState)})
            d_2121_v8_: D9
            d_2121_v8_ = D9_DC20(p2, p0, d_2115_v2_)
            d_2122_v9_: _dafny.Seq
            d_2122_v9_ = _dafny.SeqWithoutIsStrInference([-691])
            d_2123_v10_: _dafny.MultiSet
            d_2123_v10_ = _dafny.MultiSet([len(d_2122_v9_)])
            d_2124_v11_: _dafny.Array
            nw361_ = _dafny.Array(None, 22)
            nw361_[int(0)] = d_2115_v2_
            nw361_[int(1)] = d_2115_v2_
            nw361_[int(2)] = d_2115_v2_
            nw361_[int(3)] = d_2115_v2_
            nw361_[int(4)] = d_2115_v2_
            nw361_[int(5)] = (d_2115_v2_) + (d_2115_v2_)
            nw361_[int(6)] = d_2115_v2_
            nw361_[int(7)] = d_2115_v2_
            nw361_[int(8)] = ((d_2113_v0_)[(self).f24] if ((self).f24) in (d_2113_v0_) else d_2115_v2_)
            nw361_[int(9)] = 648
            nw361_[int(10)] = d_2115_v2_
            nw361_[int(11)] = d_2115_v2_
            nw361_[int(12)] = (d_2115_v2_ if (d_2116_v3_)[default__.safeIndex(d_2115_v2_, len(d_2116_v3_))] else ((d_2120_v7_)[d_2119_v6_] if (d_2119_v6_) in (d_2120_v7_) else d_2115_v2_))
            nw361_[int(13)] = (d_2115_v2_) + (((_dafny.MultiSet([d_2121_v8_])).set(D9_DC20(p2, (self).f24, (d_2123_v10_).cardinality), default__.abs(d_2115_v2_))).cardinality)
            nw361_[int(14)] = default__.safeDivisionInt(d_2115_v2_, len(p2))
            nw361_[int(15)] = (0) - (d_2115_v2_)
            nw361_[int(16)] = d_2115_v2_
            nw361_[int(17)] = (d_2115_v2_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xhna"))))
            nw361_[int(18)] = d_2115_v2_
            nw361_[int(19)] = d_2115_v2_
            nw361_[int(20)] = d_2115_v2_
            nw361_[int(21)] = (d_2115_v2_ if (self).f24 else d_2115_v2_)
            d_2124_v11_ = nw361_
            index297_ = default__.safeIndex(583, (d_2124_v11_).length(0))
            (d_2124_v11_)[index297_] = d_2115_v2_
            index298_ = default__.safeIndex(583, (d_2124_v11_).length(0))
            (d_2124_v11_)[index298_] = default__.safeDivisionInt((0) - ((p1).cf3), d_2115_v2_)
            d_2125_v12_: _dafny.Set
            d_2125_v12_ = _dafny.Set({d_2113_v0_, d_2113_v0_, d_2113_v0_, (_dafny.MultiSet([(self).f24])).intersection(_dafny.MultiSet(d_2116_v3_))})
            d_2125_v12_ = _dafny.Set({default__.fm3(p3, globalState), (d_2113_v0_).set((self).f24, default__.abs((d_2124_v11_)[default__.safeIndex(583, (d_2124_v11_).length(0))]))})
            d_2126_v13_: _dafny.Seq
            d_2126_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
            d_2127_v14_: _dafny.Map
            d_2127_v14_ = _dafny.Map({d_2115_v2_: self.f23})
            d_2126_v13_ = (default__.fm1((p3 if (self).f24 else (self).f24), globalState)).set(default__.safeIndex((d_2124_v11_)[default__.safeIndex(583, (d_2124_v11_).length(0))], len(default__.fm1((p3 if (self).f24 else (self).f24), globalState))), ((d_2127_v14_)[702] if (702) in (d_2127_v14_) else self.f23))
            d_2128_v15_: _dafny.Array
            nw362_ = _dafny.Array(False, 9)
            d_2128_v15_ = nw362_
            index299_ = default__.safeIndex(808, (d_2128_v15_).length(0))
            (d_2128_v15_)[index299_] = not ((self).f24) or (p3)
            d_2129_v16_: _dafny.Map
            d_2129_v16_ = _dafny.Map({(self).f24: False})
            d_2130_v17_: _dafny.Map
            d_2130_v17_ = _dafny.Map({True: d_2126_v13_})
            d_2131_v18_: T0
            nw363_ = C0()
            nw363_.ctor__(d_2130_v17_, True, self.f23, (self).f24)
            d_2131_v18_ = nw363_
            d_2132_v19_: _dafny.Map
            d_2132_v19_ = _dafny.Map({d_2131_v18_: (self).f24})
            d_2133_v20_: T3
            nw364_ = C1()
            nw364_.ctor__(p3, ((d_2129_v16_)[((d_2132_v19_)[d_2131_v18_] if (d_2131_v18_) in (d_2132_v19_) else True)] if (((d_2132_v19_)[d_2131_v18_] if (d_2131_v18_) in (d_2132_v19_) else True)) in (d_2129_v16_) else True))
            d_2133_v20_ = nw364_
            d_2134_v21_: _dafny.Seq
            d_2134_v21_ = _dafny.SeqWithoutIsStrInference([d_2133_v20_])
            d_2135_v22_: _dafny.Map
            d_2135_v22_ = _dafny.Map({d_2126_v13_: d_2134_v21_})
            index300_ = default__.safeIndex(808, (d_2128_v15_).length(0))
            rhs321_ = len((d_2135_v22_).set(d_2126_v13_, d_2134_v21_))
            rhs322_ = ((0) - (-745)) - (len(d_2126_v13_))
            rhs323_ = (d_2131_v18_).f24
            rhs324_ = (d_2133_v20_).f27
            lhs239_ = globalState
            lhs240_ = globalState
            lhs241_ = d_2128_v15_
            lhs242_ = default__.safeIndex(808, (d_2128_v15_).length(0))
            lhs239_.f7 = rhs321_
            lhs240_.f7 = rhs322_
            r0 = rhs323_
            lhs241_[lhs242_] = rhs324_
        elif True:
            d_2136_v23_: int
            d_2136_v23_ = 861
            d_2137_v24_: _dafny.Map
            d_2137_v24_ = _dafny.Map({(self).f24: p2})
            d_2138_v25_: _dafny.Map
            d_2138_v25_ = _dafny.Map({d_2136_v23_: d_2137_v24_})
            d_2139_v26_: T0
            nw365_ = C0()
            nw365_.ctor__((_dafny.Map({(D2_DC7(not((self).f24), (self).f24)).cf14: p2})) | (((d_2138_v25_)[d_2136_v23_] if (d_2136_v23_) in (d_2138_v25_) else d_2137_v24_)), p3, self.f23, p0)
            d_2139_v26_ = nw365_
            d_2139_v26_ = d_2139_v26_
            d_2140_v27_: _dafny.Set
            d_2140_v27_ = _dafny.Set({(self).f24})
            d_2141_v28_: _dafny.Set
            d_2141_v28_ = _dafny.Set({(0) - (d_2136_v23_), d_2136_v23_})
            d_2142_v29_: _dafny.Set
            d_2142_v29_ = _dafny.Set({d_2141_v28_, d_2141_v28_})
            (globalState).f7 = default__.fm0((d_2142_v29_).issubset(default__.fm54(d_2140_v27_, globalState)), (d_2136_v23_) - (d_2136_v23_), globalState)
            d_2143_v30_: _dafny.Array
            nw366_ = _dafny.Array(None, 23)
            nw366_[int(0)] = p0
            nw366_[int(1)] = (self).fm11(d_2136_v23_, (d_2139_v26_).f24, globalState)
            nw366_[int(2)] = p3
            nw366_[int(3)] = (d_2139_v26_).f24
            nw366_[int(4)] = p3
            nw366_[int(5)] = p0
            nw366_[int(6)] = p3
            nw366_[int(7)] = (self).f24
            nw366_[int(8)] = (self).f24
            nw366_[int(9)] = p3
            nw366_[int(10)] = (d_2139_v26_).f24
            nw366_[int(11)] = (self).f24
            nw366_[int(12)] = p0
            nw366_[int(13)] = p3
            nw366_[int(14)] = (self).f24
            nw366_[int(15)] = p0
            nw366_[int(16)] = p3
            nw366_[int(17)] = False
            nw366_[int(18)] = (self).f24
            nw366_[int(19)] = (self).f24
            nw366_[int(20)] = (d_2139_v26_).f24
            nw366_[int(21)] = p3
            nw366_[int(22)] = p0
            d_2143_v30_ = nw366_
            d_2144_v31_: _dafny.Seq
            d_2144_v31_ = _dafny.SeqWithoutIsStrInference([d_2143_v30_, d_2143_v30_])
            d_2145_v32_: _dafny.Array
            nw367_ = _dafny.Array(None, 1)
            nw367_[int(0)] = (d_2144_v31_) + (d_2144_v31_)
            d_2145_v32_ = nw367_
            index301_ = default__.safeIndex(211, (d_2145_v32_).length(0))
            (d_2145_v32_)[index301_] = d_2144_v31_
            index302_ = default__.safeIndex(211, (d_2145_v32_).length(0))
            (d_2145_v32_)[index302_] = _dafny.SeqWithoutIsStrInference([d_2143_v30_])
            (globalState).f7 = (d_2136_v23_) * (d_2136_v23_)
            index303_ = default__.safeIndex(39, (d_2143_v30_).length(0))
            (d_2143_v30_)[index303_] = (self).f24
            d_2146_v33_: _dafny.Seq
            d_2146_v33_ = _dafny.SeqWithoutIsStrInference([p3, (self).f24, p0, p3, (d_2139_v26_).f24])
            d_2147_v34_: _dafny.Seq
            d_2147_v34_ = _dafny.SeqWithoutIsStrInference([d_2146_v33_])
            d_2148_v35_: _dafny.Map
            d_2148_v35_ = _dafny.Map({d_2136_v23_: d_2136_v23_})
            d_2149_v36_: _dafny.Seq
            d_2149_v36_ = _dafny.SeqWithoutIsStrInference([d_2136_v23_, (self).fm6(globalState)])
            d_2150_v37_: C1
            nw368_ = C1()
            nw368_.ctor__(p3, p0)
            d_2150_v37_ = nw368_
            d_2151_v38_: _dafny.Map
            d_2151_v38_ = _dafny.Map({d_2150_v37_: d_2136_v23_})
            d_2152_v39_: _dafny.Map
            d_2152_v39_ = _dafny.Map({d_2136_v23_: d_2151_v38_})
            d_2153_v40_: _dafny.Map
            d_2153_v40_ = _dafny.Map({d_2136_v23_: False})
            d_2154_v41_: D19
            d_2154_v41_ = D19_DC45(_dafny.Map({d_2150_v37_: d_2136_v23_}))
            index304_ = default__.safeIndex(39, (d_2143_v30_).length(0))
            rhs325_ = ((0) - (len((((_dafny.SeqWithoutIsStrInference([(self).f24])) + (d_2146_v33_)).set(default__.safeIndex(912, len((_dafny.SeqWithoutIsStrInference([(self).f24])) + (d_2146_v33_))), (self).f24)).set(default__.safeIndex(len(d_2147_v34_), len(((_dafny.SeqWithoutIsStrInference([(self).f24])) + (d_2146_v33_)).set(default__.safeIndex(912, len((_dafny.SeqWithoutIsStrInference([(self).f24])) + (d_2146_v33_))), (self).f24))), p0)))) in (d_2148_v35_)
            rhs326_ = (len(d_2149_v36_) if not (p0) or ((self).f24) else d_2136_v23_)
            rhs327_ = (-491) > (len((d_2152_v39_).set(len(d_2153_v40_), (d_2154_v41_).cf74)))
            rhs328_ = ((d_2150_v37_).f39 if (self).f24 else (self).fm4(d_2136_v23_, d_2146_v33_, globalState))
            lhs243_ = globalState
            lhs244_ = d_2143_v30_
            lhs245_ = default__.safeIndex(39, (d_2143_v30_).length(0))
            r0 = rhs325_
            d_2136_v23_ = rhs326_
            lhs243_.f13 = rhs327_
            lhs244_[lhs245_] = rhs328_
        d_2155_v42_: int
        d_2155_v42_ = 458
        d_2156_v43_: _dafny.Map
        d_2156_v43_ = _dafny.Map({d_2155_v42_: p0})
        if not(((d_2156_v43_)[d_2155_v42_] if (d_2155_v42_) in (d_2156_v43_) else p3)):
            d_2157_v44_: _dafny.Seq
            d_2157_v44_ = _dafny.SeqWithoutIsStrInference([(d_2155_v42_) * (d_2155_v42_)])
            rhs329_ = self.f23
            rhs330_ = ((d_2157_v44_) + (d_2157_v44_)) + (d_2157_v44_)
            lhs246_ = self
            lhs246_.f23 = rhs329_
            d_2157_v44_ = rhs330_
            d_2158_v45_: C8
            nw369_ = C8()
            nw369_.ctor__(p3)
            d_2158_v45_ = nw369_
            d_2159_v46_: _dafny.Seq
            d_2159_v46_ = _dafny.SeqWithoutIsStrInference([p0, (self).f24])
            d_2160_v47_: _dafny.Set
            d_2160_v47_ = _dafny.Set({(self).f24, not((d_2159_v46_)[default__.safeIndex(d_2155_v42_, len(d_2159_v46_))]), (self).f24, p3, (self).f24})
            d_2161_v48_: _dafny.Map
            d_2161_v48_ = _dafny.Map({d_2160_v47_: _dafny.MultiSet(d_2157_v44_)})
            d_2162_v49_: D15
            d_2162_v49_ = D15_DC35(d_2155_v42_, d_2155_v42_, len(d_2160_v47_))
            d_2163_v50_: _dafny.Seq
            d_2163_v50_ = _dafny.SeqWithoutIsStrInference([d_2162_v49_, d_2162_v49_])
            d_2161_v48_ = (d_2161_v48_) | ((d_2161_v48_ if p3 else _dafny.Map({d_2160_v47_: default__.fm42(p3, len(d_2163_v50_), globalState)})))
            d_2164_v51_: _dafny.Map
            d_2164_v51_ = _dafny.Map({p0: p2})
            d_2165_v52_: _dafny.Seq
            d_2165_v52_ = _dafny.SeqWithoutIsStrInference([d_2157_v44_])
            d_2166_v53_: _dafny.Map
            d_2166_v53_ = _dafny.Map({p3: 399})
            d_2167_v54_: _dafny.Set
            d_2167_v54_ = _dafny.Set({d_2155_v42_, len(d_2165_v52_), len(d_2166_v53_), d_2155_v42_, len(p2)})
            d_2168_v55_: C0
            nw370_ = C0()
            nw370_.ctor__(((_dafny.Map({(self).f24: p2})).set(default__.fm2(globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))) | (d_2164_v51_), (d_2167_v54_).issubset(_dafny.Set({(self).fm6(globalState)})), self.f23, True)
            d_2168_v55_ = nw370_
            d_2169_v56_: D15
            d_2169_v56_ = D15_DC34((self).f24)
            source30_ = d_2169_v56_
            if source30_.is_DC34:
                d_2170___mcc_h0_ = source30_.cf55
                d_2171_cf55_ = d_2170___mcc_h0_
                (globalState).f7 = d_2155_v42_
                (globalState).f7 = d_2155_v42_
                d_2172_v57_: _dafny.Array
                nw371_ = _dafny.Array(_dafny.Seq({}), 3)
                d_2172_v57_ = nw371_
                d_2173_v58_: _dafny.Map
                d_2173_v58_ = _dafny.Map({d_2155_v42_: -824})
                d_2174_v59_: _dafny.Map
                d_2174_v59_ = _dafny.Map({True: self.f23})
                d_2175_v60_: _dafny.Seq
                d_2175_v60_ = _dafny.SeqWithoutIsStrInference([d_2174_v59_])
                index305_ = default__.safeIndex(727, (d_2172_v57_).length(0))
                (d_2172_v57_)[index305_] = (default__.fm55(d_2171_cf55_, d_2173_v58_, _dafny.SeqWithoutIsStrInference([d_2155_v42_ for d_2176_i0_ in range(default__.abs(404))]), globalState)) + (d_2175_v60_)
                d_2177_v61_: _dafny.Array
                nw372_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
                d_2177_v61_ = nw372_
                index306_ = default__.safeIndex(97, (d_2177_v61_).length(0))
                (d_2177_v61_)[index306_] = p2
                d_2178_v62_: _dafny.Array
                def lambda116_(d_2179_v42_):
                    def lambda117_(d_2180_i1_):
                        return (d_2180_i1_) * (d_2179_v42_)

                    return lambda117_

                init66_ = lambda116_(d_2155_v42_)
                nw373_ = _dafny.Array(None, 6)
                for i0_66_ in range(nw373_.length(0)):
                    nw373_[i0_66_] = init66_(i0_66_)
                d_2178_v62_ = nw373_
                index307_ = default__.safeIndex(665, (d_2178_v62_).length(0))
                (d_2178_v62_)[index307_] = ((self).fm6(globalState)) + (len(d_2157_v44_))
                d_2181_v63_: _dafny.MultiSet
                d_2181_v63_ = _dafny.MultiSet([p2, _dafny.SeqWithoutIsStrInference([self.f23 for d_2182_i2_ in range(default__.abs(887))]), p2, p2, p2])
                d_2183_v64_: _dafny.MultiSet
                d_2183_v64_ = d_2181_v63_
                d_2184_v65_: _dafny.Map
                d_2184_v65_ = _dafny.Map({self.f23: _dafny.MultiSet([p2])})
                index308_ = default__.safeIndex(727, (d_2172_v57_).length(0))
                index309_ = default__.safeIndex(97, (d_2177_v61_).length(0))
                index310_ = default__.safeIndex(665, (d_2178_v62_).length(0))
                rhs331_ = d_2175_v60_
                rhs332_ = not (p3) or (((d_2183_v64_)).ispropersubset(((d_2184_v65_)[_dafny.CodePoint('j')] if (_dafny.CodePoint('j')) in (d_2184_v65_) else d_2181_v63_)))
                rhs333_ = ((p2) + (p2)) + ((p2) + (p2))
                rhs334_ = d_2155_v42_
                lhs247_ = d_2172_v57_
                lhs248_ = default__.safeIndex(727, (d_2172_v57_).length(0))
                lhs249_ = d_2177_v61_
                lhs250_ = default__.safeIndex(97, (d_2177_v61_).length(0))
                lhs251_ = d_2178_v62_
                lhs252_ = default__.safeIndex(665, (d_2178_v62_).length(0))
                lhs247_[lhs248_] = rhs331_
                d_2171_cf55_ = rhs332_
                lhs249_[lhs250_] = rhs333_
                lhs251_[lhs252_] = rhs334_
                d_2185_v66_: _dafny.Map
                d_2185_v66_ = _dafny.Map({(d_2168_v55_).fm26((d_2178_v62_)[default__.safeIndex(665, (d_2178_v62_).length(0))], d_2159_v46_, d_2167_v54_, (d_2168_v55_).f38, globalState): (d_2178_v62_)[default__.safeIndex(665, (d_2178_v62_).length(0))]})
                d_2186_v67_: _dafny.MultiSet
                d_2186_v67_ = _dafny.MultiSet([len(d_2185_v66_), d_2155_v42_])
                (globalState).f13 = (d_2186_v67_).issubset(d_2186_v67_)
            elif source30_.is_DC35:
                d_2187___mcc_h1_ = source30_.cf56
                d_2188___mcc_h2_ = source30_.cf57
                d_2189___mcc_h3_ = source30_.cf58
                d_2190_cf58_ = d_2189___mcc_h3_
                d_2191_cf57_ = d_2188___mcc_h2_
                d_2192_cf56_ = d_2187___mcc_h1_
                d_2193_v68_: _dafny.Seq
                d_2193_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uiky"))
                d_2194_v69_: D19
                d_2194_v69_ = D19_DC46(d_2160_v47_, p2)
                d_2195_v70_: _dafny.Seq
                d_2195_v70_ = _dafny.SeqWithoutIsStrInference([d_2193_v68_, (default__.fm1((self).f24, globalState)) + (p2), p2, (d_2194_v69_).cf76])
                d_2193_v68_ = (d_2195_v70_)[default__.safeIndex(d_2191_cf57_, len(d_2195_v70_))]
                d_2196_v71_: _dafny.MultiSet
                d_2196_v71_ = _dafny.MultiSet([(d_2168_v55_).f38, (self).f24])
                d_2197_v72_: _dafny.Map
                d_2197_v72_ = _dafny.Map({d_2191_cf57_: d_2196_v71_})
                d_2156_v43_ = (d_2156_v43_).set(default__.safeDivisionInt(d_2155_v42_, (((d_2197_v72_)[d_2190_cf58_] if (d_2190_cf58_) in (d_2197_v72_) else _dafny.MultiSet([(d_2168_v55_).f38, p3]))).cardinality), not ((d_2168_v55_).f38) or (p3))
                d_2198_v73_: _dafny.Map
                d_2198_v73_ = _dafny.Map({d_2155_v42_: _dafny.SeqWithoutIsStrInference([d_2191_cf57_ for d_2199_i3_ in range(default__.abs(972))])})
                d_2198_v73_ = (d_2198_v73_).set((len(d_2193_v68_)) - (d_2155_v42_), d_2157_v44_)
                d_2200_v74_: _dafny.Map
                d_2200_v74_ = _dafny.Map({d_2157_v44_: self.f23})
                d_2201_v75_: _dafny.Map
                d_2201_v75_ = _dafny.Map({d_2193_v68_: d_2200_v74_})
                d_2201_v75_ = (d_2201_v75_).set(p2, d_2200_v74_)
            elif source30_.is_DC36:
                d_2202___mcc_h4_ = source30_.cf59
                d_2203___mcc_h5_ = source30_.cf60
                d_2204_cf60_ = d_2203___mcc_h5_
                d_2205_cf59_ = d_2202___mcc_h4_
                d_2206_v76_: _dafny.Array
                def lambda118_(d_2207_i4_):
                    return self.f23

                init67_ = lambda118_
                nw374_ = _dafny.Array(None, 28)
                for i0_67_ in range(nw374_.length(0)):
                    nw374_[i0_67_] = init67_(i0_67_)
                d_2206_v76_ = nw374_
                index311_ = default__.safeIndex(228, (d_2206_v76_).length(0))
                (d_2206_v76_)[index311_] = self.f23
                d_2208_v77_: _dafny.Array
                nw375_ = _dafny.Array(False, 8)
                d_2208_v77_ = nw375_
                d_2209_v78_: _dafny.MultiSet
                d_2209_v78_ = _dafny.MultiSet([d_2208_v77_, d_2208_v77_])
                index312_ = default__.safeIndex(228, (d_2206_v76_).length(0))
                rhs335_ = len((default__.fm1(d_2205_cf59_, globalState)) + ((p2) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qyif")))))
                rhs336_ = self.f23
                rhs337_ = ((d_2209_v78_)[d_2208_v77_] if (d_2208_v77_) in (d_2209_v78_) else d_2204_cf60_)
                rhs338_ = p3
                lhs253_ = globalState
                lhs254_ = d_2206_v76_
                lhs255_ = default__.safeIndex(228, (d_2206_v76_).length(0))
                lhs256_ = globalState
                lhs257_ = globalState
                lhs253_.f7 = rhs335_
                lhs254_[lhs255_] = rhs336_
                lhs256_.f7 = rhs337_
                lhs257_.f13 = rhs338_
                index313_ = default__.safeIndex(228, (d_2206_v76_).length(0))
                (d_2206_v76_)[index313_] = self.f23
                d_2204_cf60_ = d_2155_v42_
                d_2210_v79_: _dafny.Array
                nw376_ = _dafny.Array(int(0), 5)
                d_2210_v79_ = nw376_
                d_2211_v80_: _dafny.Set
                d_2211_v80_ = _dafny.Set({d_2210_v79_, d_2210_v79_})
                (globalState).f7 = (0) - ((0) - ((0) - (default__.safeModuloInt(len(d_2211_v80_), d_2155_v42_))))
            elif True:
                d_2212___mcc_h6_ = source30_.cf54
                d_2213_cf54_ = d_2212___mcc_h6_
                d_2214_v81_: _dafny.Array
                nw377_ = _dafny.Array(int(0), 20)
                d_2214_v81_ = nw377_
                index314_ = default__.safeIndex(181, (d_2214_v81_).length(0))
                (d_2214_v81_)[index314_] = -378
                d_2215_v82_: _dafny.MultiSet
                d_2215_v82_ = _dafny.MultiSet([d_2166_v53_])
                index315_ = default__.safeIndex(181, (d_2214_v81_).length(0))
                (d_2214_v81_)[index315_] = ((d_2215_v82_).intersection((d_2215_v82_) - ((d_2215_v82_).set(d_2166_v53_, default__.abs(d_2155_v42_))))).cardinality
                d_2216_v83_: T0
                nw378_ = C0()
                nw378_.ctor__(d_2164_v51_, (d_2168_v55_).f38, self.f23, (d_2168_v55_).f38)
                d_2216_v83_ = nw378_
                d_2217_v84_: T0
                d_2217_v84_ = d_2216_v83_
                d_2217_v84_ = d_2217_v84_
                d_2218_v85_: _dafny.Array
                nw379_ = _dafny.Array(None, 13)
                d_2218_v85_ = nw379_
                d_2219_v86_: _dafny.Map
                d_2219_v86_ = _dafny.Map({d_2218_v85_: p2})
                d_2219_v86_ = (d_2219_v86_).set(d_2218_v85_, p2)
                (globalState).f13 = p0
        elif True:
            (globalState).f7 = (d_2155_v42_) * ((d_2155_v42_) + (d_2155_v42_))
            d_2220_v87_: _dafny.Map
            d_2220_v87_ = _dafny.Map({p0: self.f23})
            (globalState).f13 = not(not((d_2220_v87_) != (d_2220_v87_)))
            d_2221_v88_: _dafny.Map
            d_2221_v88_ = _dafny.Map({p0: p2})
            d_2222_v89_: C0
            nw380_ = C0()
            nw380_.ctor__(d_2221_v88_, p0, self.f23, p0)
            d_2222_v89_ = nw380_
            d_2155_v42_ = ((self).fm6(globalState)) - (d_2155_v42_)
            d_2223_v90_: _dafny.Map
            d_2223_v90_ = _dafny.Map({(0) - ((len(_dafny.Map({len(p2): (self).f24}))) - (d_2155_v42_)): default__.fm56(p3, d_2155_v42_, globalState)})
            def iife208_():
                coll94_ = _dafny.Map()
                compr_94_: int
                for compr_94_ in (_dafny.MultiSet([d_2155_v42_])).Elements:
                    d_2224_v91_: int = compr_94_
                    if (d_2224_v91_) in (_dafny.MultiSet([d_2155_v42_])):
                        coll94_[default__.safeModuloInt(d_2224_v91_, d_2155_v42_)] = d_2156_v43_
                return _dafny.Map(coll94_)
            d_2223_v90_ = (iife208_()
            ) | (d_2223_v90_)
        d_2225_v92_: _dafny.Seq
        d_2225_v92_ = _dafny.SeqWithoutIsStrInference([len(p2)])
        d_2226_v93_: _dafny.Set
        d_2226_v93_ = _dafny.Set({(self).fm6(globalState), len((_dafny.SeqWithoutIsStrInference([d_2155_v42_])) + (d_2225_v92_))})
        d_2226_v93_ = d_2226_v93_
        d_2227_v94_: _dafny.Set
        d_2227_v94_ = _dafny.Set({True})
        d_2228_v95_: _dafny.MultiSet
        d_2228_v95_ = _dafny.MultiSet([d_2155_v42_, (0) - (d_2155_v42_)])
        d_2229_v96_: _dafny.Map
        d_2229_v96_ = _dafny.Map({len(_dafny.Set({d_2227_v94_, _dafny.Set({p0, p3, True})})): (d_2228_v95_).cardinality})
        d_2229_v96_ = (d_2229_v96_) | (d_2229_v96_)
        d_2230_v97_: D9
        d_2230_v97_ = D9_DC20(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "of")), (self).f24, d_2155_v42_)
        d_2231_v98_: _dafny.MultiSet
        d_2231_v98_ = _dafny.MultiSet([d_2230_v97_])
        d_2232_v99_: D21
        d_2232_v99_ = D21_DC50(d_2231_v98_)
        d_2231_v98_ = (d_2232_v99_).cf80
        d_2233_i5_: int
        d_2233_i5_ = 0
        with _dafny.label("7"):
            while (self).f24:
                with _dafny.c_label("7"):
                    if (d_2233_i5_) >= (100):
                        raise _dafny.Break("7")
                    d_2233_i5_ = (d_2233_i5_) + (1)
                    d_2156_v43_ = (d_2156_v43_).set((0) - ((d_2155_v42_) + (d_2155_v42_)), (d_2155_v42_) == (d_2155_v42_))
                    d_2234_v100_: _dafny.MultiSet
                    d_2234_v100_ = _dafny.MultiSet([d_2226_v93_])
                    if (d_2234_v100_).ispropersubset((d_2234_v100_ if (self).f24 else d_2234_v100_)):
                        d_2235_v101_: _dafny.Array
                        nw381_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
                        d_2235_v101_ = nw381_
                        d_2236_v102_: _dafny.Array
                        nw382_ = _dafny.Array(None, 7)
                        nw382_[int(0)] = p0
                        nw382_[int(1)] = default__.fm2(globalState)
                        nw382_[int(2)] = p0
                        nw382_[int(3)] = (self).f24
                        nw382_[int(4)] = True
                        nw382_[int(5)] = p0
                        nw382_[int(6)] = (self).f24
                        d_2236_v102_ = nw382_
                        d_2237_v103_: _dafny.Array
                        nw383_ = _dafny.Array(None, 3)
                        nw383_[int(0)] = d_2236_v102_
                        nw383_[int(1)] = d_2236_v102_
                        nw383_[int(2)] = d_2236_v102_
                        d_2237_v103_ = nw383_
                        index316_ = default__.safeIndex(987, (d_2237_v103_).length(0))
                        (d_2237_v103_)[index316_] = d_2236_v102_
                        index317_ = default__.safeIndex(987, (d_2237_v103_).length(0))
                        rhs339_ = d_2235_v101_
                        rhs340_ = (self).f24
                        rhs341_ = (p2) <= ((_dafny.SeqWithoutIsStrInference([(p2)[default__.safeIndex(d_2155_v42_, len(p2))] for d_2238_i6_ in range(default__.abs(729))])) + (p2))
                        rhs342_ = d_2155_v42_
                        rhs343_ = d_2236_v102_
                        lhs258_ = globalState
                        lhs259_ = globalState
                        lhs260_ = d_2237_v103_
                        lhs261_ = default__.safeIndex(987, (d_2237_v103_).length(0))
                        d_2235_v101_ = rhs339_
                        lhs258_.f13 = rhs340_
                        r0 = rhs341_
                        lhs259_.f7 = rhs342_
                        lhs260_[lhs261_] = rhs343_
                        d_2239_v104_: _dafny.Array
                        def lambda119_(d_2240_i7_):
                            return self.f23

                        init68_ = lambda119_
                        nw384_ = _dafny.Array(None, 2)
                        for i0_68_ in range(nw384_.length(0)):
                            nw384_[i0_68_] = init68_(i0_68_)
                        d_2239_v104_ = nw384_
                        d_2239_v104_ = d_2239_v104_
                        d_2241_v105_: _dafny.MultiSet
                        d_2241_v105_ = _dafny.MultiSet([p0])
                        d_2242_v106_: C2
                        nw385_ = C2()
                        nw385_.ctor__(d_2241_v105_, d_2239_v104_, self.f23, p3)
                        d_2242_v106_ = nw385_
                        d_2243_v107_: D5
                        d_2243_v107_ = D5_DC13(d_2155_v42_, p3, p3)
                        rhs344_ = (d_2243_v107_).cf20
                        rhs345_ = d_2155_v42_
                        rhs346_ = (self.f23) not in (p2)
                        rhs347_ = (0) - (d_2155_v42_)
                        lhs262_ = globalState
                        lhs263_ = globalState
                        lhs264_ = globalState
                        lhs262_.f7 = rhs344_
                        d_2155_v42_ = rhs345_
                        lhs263_.f13 = rhs346_
                        lhs264_.f7 = rhs347_
                        r0 = not((self).f24)
                    elif True:
                        d_2244_v108_: _dafny.Array
                        def lambda120_(d_2245_v95_):
                            def lambda121_(d_2246_i8_):
                                return d_2245_v95_

                            return lambda121_

                        init69_ = lambda120_(d_2228_v95_)
                        nw386_ = _dafny.Array(None, 28)
                        for i0_69_ in range(nw386_.length(0)):
                            nw386_[i0_69_] = init69_(i0_69_)
                        d_2244_v108_ = nw386_
                        d_2244_v108_ = d_2244_v108_
                        (globalState).f13 = (d_2155_v42_) != (d_2155_v42_)
                        d_2247_v109_: _dafny.Seq
                        d_2247_v109_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjiuamnmc"))
                        d_2247_v109_ = d_2247_v109_
                        d_2248_v110_: D0
                        d_2248_v110_ = D0_DC1(len(d_2229_v96_))
                        d_2249_v111_: _dafny.Seq
                        d_2249_v111_ = _dafny.SeqWithoutIsStrInference([D0_DC1(38), d_2248_v110_, d_2248_v110_])
                        d_2250_v112_: _dafny.Map
                        d_2250_v112_ = _dafny.Map({339: _dafny.Set({d_2155_v42_})})
                        d_2251_v113_: _dafny.Seq
                        d_2251_v113_ = _dafny.SeqWithoutIsStrInference([True, p3])
                        d_2252_v114_: _dafny.Array
                        nw387_ = _dafny.Array(None, 28)
                        nw387_[int(0)] = p3
                        nw387_[int(1)] = p3
                        nw387_[int(2)] = p3
                        nw387_[int(3)] = (self).f24
                        nw387_[int(4)] = False
                        nw387_[int(5)] = p0
                        nw387_[int(6)] = p0
                        nw387_[int(7)] = p3
                        nw387_[int(8)] = (self).fm4(d_2155_v42_, d_2251_v113_, globalState)
                        nw387_[int(9)] = default__.fm2(globalState)
                        nw387_[int(10)] = p0
                        nw387_[int(11)] = False
                        nw387_[int(12)] = p0
                        nw387_[int(13)] = p3
                        nw387_[int(14)] = p0
                        nw387_[int(15)] = p3
                        nw387_[int(16)] = False
                        nw387_[int(17)] = p0
                        nw387_[int(18)] = p0
                        nw387_[int(19)] = not((self).f24)
                        nw387_[int(20)] = p0
                        nw387_[int(21)] = p0
                        nw387_[int(22)] = p3
                        nw387_[int(23)] = p0
                        nw387_[int(24)] = True
                        nw387_[int(25)] = p3
                        nw387_[int(26)] = p3
                        nw387_[int(27)] = p0
                        d_2252_v114_ = nw387_
                        d_2253_v115_: _dafny.Array
                        nw388_ = _dafny.Array(_dafny.CodePoint('D'), 15)
                        d_2253_v115_ = nw388_
                        d_2254_v116_: C3
                        nw389_ = C3()
                        nw389_.ctor__(_dafny.Map({d_2252_v114_: self.f23}), d_2253_v115_, _dafny.CodePoint('v'), False)
                        d_2254_v116_ = nw389_
                        d_2255_v117_: _dafny.Seq
                        d_2255_v117_ = _dafny.SeqWithoutIsStrInference([d_2254_v116_, d_2254_v116_])
                        d_2256_v118_: bool
                        out64_: bool
                        out64_ = (self).m5(d_2249_v111_, (d_2226_v93_).isdisjoint(((d_2250_v112_)[d_2155_v42_] if (d_2155_v42_) in (d_2250_v112_) else d_2226_v93_)), not(False), len(d_2255_v117_), globalState)
                        d_2256_v118_ = out64_
                        (globalState).f7 = (0) - (len(d_2247_v109_))
                    d_2257_v119_: _dafny.Map
                    d_2257_v119_ = _dafny.Map({p0: p0})
                    d_2258_v120_: _dafny.Seq
                    d_2258_v120_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_2259_v121_: _dafny.Map
                    d_2259_v121_ = _dafny.Map({d_2258_v120_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ri"))})
                    d_2260_v122_: _dafny.Map
                    d_2260_v122_ = _dafny.Map({(D2_DC7(p3, p0)).cf13: (d_2259_v121_ if (self).fm11(len(d_2257_v119_), (self).f24, globalState) else d_2259_v121_)})
                    d_2260_v122_ = (d_2260_v122_).set((self.f23) not in (p2), d_2259_v121_)
                    d_2261_v123_: _dafny.Array
                    nw390_ = _dafny.Array(D4.default()(), 16)
                    d_2261_v123_ = nw390_
                    d_2262_v124_: D4
                    d_2262_v124_ = D4_DC11(d_2155_v42_)
                    index318_ = default__.safeIndex(134, (d_2261_v123_).length(0))
                    (d_2261_v123_)[index318_] = d_2262_v124_
                    index319_ = default__.safeIndex(134, (d_2261_v123_).length(0))
                    (d_2261_v123_)[index319_] = D4_DC11(d_2155_v42_)
                    pass
            pass
        d_2263_v125_: _dafny.Seq
        d_2263_v125_ = _dafny.SeqWithoutIsStrInference([(d_2155_v42_) in (d_2226_v93_)])
        r0 = (d_2263_v125_)[default__.safeIndex((0) - ((d_2155_v42_ if (self).f24 else d_2155_v42_)), len(d_2263_v125_))]
        return r0


class C10(T1, T2, T3, T0):
    def  __init__(self):
        self._f23: str = _dafny.CodePoint('D')
        self._f24: bool = False
        self._f25: _dafny.Array = _dafny.Array(None, 0)
        self._f27: bool = False
        self.f29: int = int(0)
        self._f28: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    @property
    def f23(self):
        return self._f23
    @f23.setter
    def f23(self, value):
        self._f23 = value
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f28, f29, f23, f24, f25, f27):
        (self)._f28 = f28
        (self).f29 = f29
        (self).f23 = f23
        (self)._f24 = f24
        (self)._f25 = f25
        (self)._f27 = f27

    def fm5(self, globalState):
        return D0_DC3((D0_DC0(_dafny.MultiSet([(self).f28])) if (self).f27 else D0_DC3(D0_DC2(True, len(_dafny.SeqWithoutIsStrInference([997])), (self).f28))))

    def fm6(self, globalState):
        return (0) - ((self.f29) + (self.f29))

    def fm4(self, p0, p1, globalState):
        return not((self).f28)

    def fm7(self, globalState):
        if (self.f29) >= (len(_dafny.SeqWithoutIsStrInference([(self).f27]))):
            return (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f29]))).cardinality])) + (_dafny.SeqWithoutIsStrInference([self.f29 for d_2264_i0_ in range(default__.abs(874))]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([self.f29])

    def fm8(self, p0, p1, globalState):
        return ((_dafny.Map({(self).f27: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "th"))})) | (_dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jedxluog"))}))) | (_dafny.Map({(self).f28: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "afbr"))}))

    def fm9(self, p0, p1, p2, globalState):
        return len(((_dafny.Map({self.f29: self.f29})) | (_dafny.Map({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfg")))})): len(_dafny.Map({self.f29: self.f23}))}))) | ((_dafny.Map({self.f29: self.f29})) | (_dafny.Map({self.f29: self.f29}))))

    def fm10(self, p0, p1, globalState):
        return self.f23

    def m1(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: D0 = D0.default()()
        r2: bool = False
        d_2265_v0_: _dafny.Array
        def lambda122_(d_2266_p0_):
            def lambda123_(d_2267_i0_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qccj"))).set(default__.safeIndex(self.f29, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qccj")))), d_2266_p0_)

            return lambda123_

        init70_ = lambda122_(p0)
        nw391_ = _dafny.Array(None, 13)
        for i0_70_ in range(nw391_.length(0)):
            nw391_[i0_70_] = init70_(i0_70_)
        d_2265_v0_ = nw391_
        d_2268_v1_: _dafny.Seq
        d_2268_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbcmfn"))
        index320_ = default__.safeIndex(454, (d_2265_v0_).length(0))
        (d_2265_v0_)[index320_] = d_2268_v1_
        index321_ = default__.safeIndex(454, (d_2265_v0_).length(0))
        (d_2265_v0_)[index321_] = (d_2268_v1_).set(default__.safeIndex(default__.safeDivisionInt(self.f29, (0) - (self.f29)), len(d_2268_v1_)), p0)
        d_2269_v2_: C1
        nw392_ = C1()
        nw392_.ctor__((not((self).f24)) and (p1), (self).f28)
        d_2269_v2_ = nw392_
        d_2270_v3_: _dafny.Seq
        d_2270_v3_ = _dafny.SeqWithoutIsStrInference([(default__.fm1((self).f28, globalState)) + (_dafny.SeqWithoutIsStrInference([self.f23 for d_2271_i1_ in range(default__.abs(679))])), d_2268_v1_, d_2268_v1_, (d_2268_v1_) + (d_2268_v1_), (d_2265_v0_)[default__.safeIndex(454, (d_2265_v0_).length(0))]])
        d_2270_v3_ = ((default__.fm21(_dafny.Set({(self).f27}), _dafny.MultiSet((self).fm7(globalState)), globalState)) + (d_2270_v3_)) + (d_2270_v3_)
        d_2272_v4_: _dafny.Array
        nw393_ = _dafny.Array(_dafny.Array(None, 0), 2)
        d_2272_v4_ = nw393_
        d_2273_v5_: C7
        nw394_ = C7()
        nw394_.ctor__((d_2269_v2_).f39, (d_2269_v2_).f39, (self).f25, self.f23, not((self).f27), (self).f24)
        d_2273_v5_ = nw394_
        d_2274_v6_: _dafny.Array
        nw395_ = _dafny.Array(None, 14)
        nw395_[int(0)] = d_2273_v5_
        nw395_[int(1)] = d_2273_v5_
        nw395_[int(2)] = d_2273_v5_
        nw395_[int(3)] = d_2273_v5_
        nw395_[int(4)] = d_2273_v5_
        nw395_[int(5)] = d_2273_v5_
        nw395_[int(6)] = d_2273_v5_
        nw395_[int(7)] = d_2273_v5_
        nw395_[int(8)] = d_2273_v5_
        nw395_[int(9)] = d_2273_v5_
        nw395_[int(10)] = d_2273_v5_
        nw395_[int(11)] = d_2273_v5_
        nw395_[int(12)] = d_2273_v5_
        nw395_[int(13)] = d_2273_v5_
        d_2274_v6_ = nw395_
        d_2275_v7_: _dafny.Map
        d_2275_v7_ = _dafny.Map({self.f23: d_2274_v6_})
        index322_ = default__.safeIndex(731, (d_2272_v4_).length(0))
        (d_2272_v4_)[index322_] = ((d_2275_v7_)[self.f23] if (self.f23) in (d_2275_v7_) else d_2274_v6_)
        index323_ = default__.safeIndex(731, (d_2272_v4_).length(0))
        nw396_ = _dafny.Array(None, 4)
        nw396_[int(0)] = d_2273_v5_
        nw396_[int(1)] = d_2273_v5_
        nw396_[int(2)] = (d_2273_v5_ if (d_2273_v5_).f30 else d_2273_v5_)
        nw396_[int(3)] = d_2273_v5_
        (d_2272_v4_)[index323_] = nw396_
        _ingredientsd_1 = []
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_2265_v0_).length(0)):
            d_2276_i2_: int = guard_loop_4_
            if (True) and (((0) <= (d_2276_i2_)) and ((d_2276_i2_) < ((d_2265_v0_).length(0)))):
                _ingredientsd_1.append((d_2265_v0_, int(d_2276_i2_), (((d_2265_v0_)[default__.safeIndex(454, (d_2265_v0_).length(0))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmtvniulv")))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_2277_i3_ in range(default__.abs(310))]))))
        for _tupd_1 in _ingredientsd_1:
            _tupd_1[0][_tupd_1[1]] = _tupd_1[2]
        d_2278_v8_: D9
        d_2278_v8_ = D9_DC20((d_2265_v0_)[default__.safeIndex(454, (d_2265_v0_).length(0))], default__.fm2(globalState), len(d_2268_v1_))
        d_2279_v9_: _dafny.MultiSet
        d_2279_v9_ = _dafny.MultiSet([d_2278_v8_])
        d_2280_v10_: D21
        d_2280_v10_ = D21_DC50((d_2279_v9_) - (d_2279_v9_))
        source31_ = d_2280_v10_
        if source31_.is_DC51:
            d_2281___mcc_h0_ = source31_.cf81
            d_2282_cf81_ = d_2281___mcc_h0_
            d_2268_v1_ = (_dafny.SeqWithoutIsStrInference([p0 for d_2283_i4_ in range(default__.abs(43))]) if not (d_2282_cf81_) or ((d_2273_v5_).f30) else (default__.fm1((self).f24, globalState)) + (d_2268_v1_))
            d_2284_v11_: bool
            out65_: bool
            out65_ = (d_2273_v5_).m4((d_2273_v5_).f30, globalState)
            d_2284_v11_ = out65_
            d_2285_v12_: _dafny.Seq
            d_2285_v12_ = _dafny.SeqWithoutIsStrInference([(d_2273_v5_).f30, not((d_2273_v5_).f30), (d_2278_v8_).cf32])
            d_2284_v11_ = (len(d_2285_v12_)) < (self.f29)
            d_2286_v13_: _dafny.Seq
            d_2286_v13_ = _dafny.SeqWithoutIsStrInference([self.f29])
            d_2287_v14_: _dafny.Map
            d_2287_v14_ = _dafny.Map({843: (d_2286_v13_)[default__.safeIndex(self.f29, len(d_2286_v13_))]})
            d_2285_v12_ = (d_2285_v12_).set(default__.safeIndex(((d_2287_v14_)[157] if (157) in (d_2287_v14_) else self.f29), len(d_2285_v12_)), d_2284_v11_)
        elif source31_.is_DC50:
            d_2288___mcc_h1_ = source31_.cf80
            d_2289_cf80_ = d_2288___mcc_h1_
            d_2290_v15_: C9
            nw397_ = C9()
            nw397_.ctor__(self.f23, (self).f28)
            d_2290_v15_ = nw397_
            d_2291_v16_: _dafny.Seq
            d_2291_v16_ = _dafny.SeqWithoutIsStrInference([(d_2273_v5_).f30, False])
            d_2292_v17_: _dafny.Array
            nw398_ = _dafny.Array(None, 26)
            nw398_[int(0)] = (self).f27
            nw398_[int(1)] = (d_2273_v5_).f30
            nw398_[int(2)] = (self).fm4(self.f29, d_2291_v16_, globalState)
            nw398_[int(3)] = True
            nw398_[int(4)] = (self).f28
            nw398_[int(5)] = p1
            nw398_[int(6)] = (d_2273_v5_).f30
            nw398_[int(7)] = True
            nw398_[int(8)] = (self).f28
            nw398_[int(9)] = (self).f24
            nw398_[int(10)] = False
            nw398_[int(11)] = (self).f28
            nw398_[int(12)] = (d_2269_v2_).f39
            nw398_[int(13)] = (d_2269_v2_).f39
            nw398_[int(14)] = not(not((d_2269_v2_).f39))
            nw398_[int(15)] = (d_2273_v5_).f30
            nw398_[int(16)] = p1
            nw398_[int(17)] = not((d_2269_v2_).f39)
            nw398_[int(18)] = True
            nw398_[int(19)] = (d_2273_v5_).f31
            nw398_[int(20)] = (self).f27
            nw398_[int(21)] = p1
            nw398_[int(22)] = (d_2273_v5_).f31
            nw398_[int(23)] = (d_2273_v5_).fm4(self.f29, (_dafny.SeqWithoutIsStrInference([False, (self).f24, (d_2269_v2_).f39])).set(default__.safeIndex(self.f29, len(_dafny.SeqWithoutIsStrInference([False, (self).f24, (d_2269_v2_).f39]))), (d_2269_v2_).f39), globalState)
            nw398_[int(24)] = (d_2269_v2_).f39
            nw398_[int(25)] = (d_2273_v5_).f31
            d_2292_v17_ = nw398_
            d_2293_v18_: C5
            nw399_ = C5()
            nw399_.ctor__(d_2292_v17_)
            d_2293_v18_ = nw399_
            if (self.f29) >= (self.f29):
                (globalState).f7 = self.f29
                (globalState).f7 = self.f29
                d_2294_v19_: _dafny.MultiSet
                d_2294_v19_ = _dafny.MultiSet([len((d_2265_v0_)[default__.safeIndex(454, (d_2265_v0_).length(0))]), self.f29, self.f29, (0) - ((0) - (self.f29)), self.f29])
                d_2295_v20_: D5
                d_2295_v20_ = D5_DC13(((d_2294_v19_)[self.f29] if (self.f29) in (d_2294_v19_) else self.f29), True, False)
                d_2295_v20_ = d_2295_v20_
                d_2296_v21_: _dafny.Array
                nw400_ = _dafny.Array(int(0), 16)
                d_2296_v21_ = nw400_
                (globalState).f22 = d_2296_v21_
                d_2297_v22_: _dafny.Set
                d_2297_v22_ = _dafny.Set({True, (d_2273_v5_).f31})
                d_2298_v23_: _dafny.Map
                d_2298_v23_ = _dafny.Map({(d_2297_v22_).intersection(d_2297_v22_): d_2270_v3_})
                d_2298_v23_ = (d_2298_v23_).set((d_2297_v22_) - (d_2297_v22_), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))]))
            elif True:
                (globalState).f7 = (D5_DC13(self.f29, (d_2269_v2_).f39, (self).f24)).cf20
                (self).f29 = self.f29
                d_2299_v24_: _dafny.Array
                def lambda124_(d_2300_i5_):
                    return (d_2300_i5_) * (len(_dafny.SeqWithoutIsStrInference([self.f29, self.f29])))

                init71_ = lambda124_
                nw401_ = _dafny.Array(None, 2)
                for i0_71_ in range(nw401_.length(0)):
                    nw401_[i0_71_] = init71_(i0_71_)
                d_2299_v24_ = nw401_
                index324_ = default__.safeIndex(748, (d_2299_v24_).length(0))
                (d_2299_v24_)[index324_] = self.f29
                d_2301_v25_: _dafny.Seq
                d_2301_v25_ = _dafny.SeqWithoutIsStrInference([len((d_2291_v16_) + (d_2291_v16_)), self.f29])
                index325_ = default__.safeIndex(748, (d_2299_v24_).length(0))
                rhs348_ = (0) - (self.f29)
                rhs349_ = self.f23
                rhs350_ = (((_dafny.SeqWithoutIsStrInference([self.f29])).set(default__.safeIndex(self.f29, len(_dafny.SeqWithoutIsStrInference([self.f29]))), self.f29)) + (d_2301_v25_)) + ((d_2301_v25_) + (d_2301_v25_))
                rhs351_ = (0) - (self.f29)
                lhs265_ = d_2299_v24_
                lhs266_ = default__.safeIndex(748, (d_2299_v24_).length(0))
                lhs267_ = self
                lhs268_ = globalState
                lhs265_[lhs266_] = rhs348_
                lhs267_.f23 = rhs349_
                d_2301_v25_ = rhs350_
                lhs268_.f7 = rhs351_
                d_2302_v26_: D21
                d_2302_v26_ = D21_DC51(True)
                d_2303_v27_: _dafny.Map
                d_2303_v27_ = _dafny.Map({(d_2302_v26_).cf81: (d_2273_v5_).f31})
                d_2303_v27_ = (d_2303_v27_).set((d_2273_v5_).f30, (d_2291_v16_)[default__.safeIndex(len(d_2268_v1_), len(d_2291_v16_))])
                (globalState).f13 = ((d_2268_v1_) + ((d_2265_v0_)[default__.safeIndex(454, (d_2265_v0_).length(0))])) < (d_2268_v1_)
            d_2304_v28_: C1
            nw402_ = C1()
            nw402_.ctor__((d_2269_v2_).f39, (d_2269_v2_).f39)
            d_2304_v28_ = nw402_
        elif True:
            d_2305___mcc_h2_ = source31_.cf82
            d_2306_cf82_ = d_2305___mcc_h2_
            (self).f29 = self.f29
            r2 = (d_2273_v5_).f31
            d_2307_v29_: _dafny.Array
            nw403_ = _dafny.Array(False, 26)
            d_2307_v29_ = nw403_
            d_2308_v30_: C5
            nw404_ = C5()
            nw404_.ctor__(d_2307_v29_)
            d_2308_v30_ = nw404_
            d_2309_v31_: D16
            d_2309_v31_ = D16_DC37(d_2308_v30_)
            d_2310_v32_: D16
            d_2310_v32_ = D16_DC39(d_2309_v31_)
            pat_let_tv45_ = d_2309_v31_
            d_2311_v33_: _dafny.Array
            nw405_ = _dafny.Array(None, 17)
            nw405_[int(0)] = d_2310_v32_
            nw405_[int(1)] = d_2310_v32_
            nw405_[int(2)] = d_2310_v32_
            nw405_[int(3)] = d_2310_v32_
            nw405_[int(4)] = d_2310_v32_
            nw405_[int(5)] = d_2310_v32_
            nw405_[int(6)] = D16_DC39(d_2309_v31_)
            nw405_[int(7)] = (d_2310_v32_ if (d_2273_v5_).f31 else d_2310_v32_)
            def iife209_(_pat_let57_0):
                def iife210_(d_2312_dt__update__tmp_h0_):
                    def iife211_(_pat_let58_0):
                        def iife212_(d_2313_dt__update_hcf65_h0_):
                            return D16_DC39(d_2313_dt__update_hcf65_h0_)
                        return iife212_(_pat_let58_0)
                    return iife211_(pat_let_tv45_)
                return iife210_(_pat_let57_0)
            nw405_[int(8)] = iife209_(d_2310_v32_)
            nw405_[int(9)] = D16_DC39(d_2309_v31_)
            nw405_[int(10)] = d_2310_v32_
            nw405_[int(11)] = d_2310_v32_
            nw405_[int(12)] = d_2310_v32_
            nw405_[int(13)] = D16_DC39(d_2309_v31_)
            nw405_[int(14)] = d_2310_v32_
            nw405_[int(15)] = d_2310_v32_
            nw405_[int(16)] = d_2310_v32_
            d_2311_v33_ = nw405_
            index326_ = default__.safeIndex(851, (d_2311_v33_).length(0))
            (d_2311_v33_)[index326_] = D16_DC39(d_2309_v31_)
            index327_ = default__.safeIndex(851, (d_2311_v33_).length(0))
            (d_2311_v33_)[index327_] = d_2310_v32_
            d_2314_v34_: D21
            d_2314_v34_ = D21_DC51((d_2273_v5_).f30)
            d_2314_v34_ = d_2314_v34_
        r0 = _dafny.SeqWithoutIsStrInference([p0 for d_2315_i6_ in range(default__.abs(-100))])
        d_2316_v35_: D0
        d_2316_v35_ = D0_DC1(default__.safeDivisionInt(len((d_2265_v0_)[default__.safeIndex(454, (d_2265_v0_).length(0))]), self.f29))
        r1 = d_2316_v35_
        d_2317_v36_: _dafny.MultiSet
        d_2317_v36_ = _dafny.MultiSet([self.f29])
        r2 = (((d_2317_v36_ if (self).f24 else d_2317_v36_)).cardinality) > (self.f29)
        return r0, r1, r2

    def m2(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        (globalState).f13 = (self).f24
        (globalState).f7 = default__.safeModuloInt(self.f29, self.f29)
        d_2318_v0_: _dafny.Array
        nw406_ = _dafny.Array(D0.default()(), 22)
        d_2318_v0_ = nw406_
        d_2319_v1_: _dafny.Map
        d_2319_v1_ = _dafny.Map({self.f23: -956})
        d_2320_v2_: D0
        d_2320_v2_ = D0_DC1(len(d_2319_v1_))
        d_2321_v3_: D0
        d_2321_v3_ = D0_DC3(d_2320_v2_)
        index328_ = default__.safeIndex(855, (d_2318_v0_).length(0))
        (d_2318_v0_)[index328_] = d_2321_v3_
        index329_ = default__.safeIndex(855, (d_2318_v0_).length(0))
        (d_2318_v0_)[index329_] = d_2321_v3_
        d_2322_v4_: _dafny.Array
        nw407_ = _dafny.Array(int(0), 28)
        d_2322_v4_ = nw407_
        d_2323_v5_: _dafny.MultiSet
        d_2323_v5_ = _dafny.MultiSet([d_2322_v4_])
        (globalState).f7 = (self.f29) - ((d_2323_v5_).cardinality)
        d_2324_v6_: _dafny.Seq
        d_2324_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
        d_2325_v7_: _dafny.MultiSet
        d_2325_v7_ = _dafny.MultiSet([d_2324_v6_, d_2324_v6_])
        d_2326_v8_: _dafny.MultiSet
        d_2326_v8_ = (d_2325_v7_) - (_dafny.MultiSet([d_2324_v6_, d_2324_v6_]))
        source32_ = d_2326_v8_
        d_2327___mcc_h0_ = source32_
        d_2328_cf79_ = d_2327___mcc_h0_
        d_2329_v9_: _dafny.Map
        d_2329_v9_ = _dafny.Map({(self).f24: 865})
        d_2330_v10_: _dafny.MultiSet
        d_2330_v10_ = _dafny.MultiSet([len(d_2329_v9_), self.f29])
        d_2331_v11_: _dafny.Map
        d_2331_v11_ = _dafny.Map({False: (d_2330_v10_).ispropersubset(d_2330_v10_)})
        d_2332_v12_: _dafny.Array
        nw408_ = _dafny.Array(False, 23)
        d_2332_v12_ = nw408_
        d_2333_v13_: C5
        nw409_ = C5()
        nw409_.ctor__(d_2332_v12_)
        d_2333_v13_ = nw409_
        d_2334_v14_: _dafny.Set
        d_2334_v14_ = _dafny.Set({d_2333_v13_, d_2333_v13_, d_2333_v13_, d_2333_v13_, d_2333_v13_})
        d_2331_v11_ = (d_2331_v11_).set(default__.fm2(globalState), (d_2334_v14_).isdisjoint(_dafny.Set({d_2333_v13_})))
        d_2335_v15_: _dafny.Set
        d_2335_v15_ = _dafny.Set({d_2332_v12_, (d_2333_v13_).f32, d_2332_v12_})
        (globalState).f13 = (d_2335_v15_).ispropersubset(d_2335_v15_)
        (globalState).f13 = (self).f28
        d_2336_v16_: _dafny.MultiSet
        d_2336_v16_ = _dafny.MultiSet([(self).f27])
        d_2337_v17_: _dafny.Map
        d_2337_v17_ = _dafny.Map({(self).f27: (self).f25})
        d_2338_v18_: C2
        nw410_ = C2()
        nw410_.ctor__(d_2336_v16_, ((d_2337_v17_)[(self).f27] if ((self).f27) in (d_2337_v17_) else (self).f25), self.f23, (self).f24)
        d_2338_v18_ = nw410_
        hi14_ = self.f29
        for d_2339_i0_ in range((0) - (self.f29), hi14_):
            index330_ = default__.safeIndex(609, (d_2322_v4_).length(0))
            (d_2322_v4_)[index330_] = d_2339_i0_
            d_2340_v19_: _dafny.Seq
            d_2340_v19_ = _dafny.SeqWithoutIsStrInference([self.f29])
            index331_ = default__.safeIndex(609, (d_2322_v4_).length(0))
            (d_2322_v4_)[index331_] = (default__.safeDivisionInt(self.f29, (d_2340_v19_)[default__.safeIndex(self.f29, len(d_2340_v19_))])) * ((d_2339_i0_) - (d_2339_i0_))
            d_2341_v20_: D12
            d_2341_v20_ = D12_DC25((self).f24, (d_2324_v6_)[default__.safeIndex(self.f29, len(d_2324_v6_))])
            d_2342_v21_: D12
            d_2342_v21_ = D12_DC26(d_2341_v20_)
            d_2343_v22_: D12
            d_2343_v22_ = D12_DC26(d_2341_v20_)
            source33_ = d_2343_v22_
            if source33_.is_DC25:
                d_2344___mcc_h1_ = source33_.cf41
                d_2345___mcc_h2_ = source33_.cf42
                d_2346_cf42_ = d_2345___mcc_h2_
                d_2347_cf41_ = d_2344___mcc_h1_
                d_2348_v23_: _dafny.MultiSet
                d_2348_v23_ = _dafny.MultiSet([(self).f28, (self).f24, (self).f27])
                d_2349_v24_: int
                d_2350_v25_: bool
                d_2351_v26_: int
                out66_: int
                out67_: bool
                out68_: int
                out66_, out67_, out68_ = default__.m0(len(d_2324_v6_), d_2348_v23_, d_2347_cf41_, ((d_2322_v4_)[default__.safeIndex(609, (d_2322_v4_).length(0))]) + ((d_2322_v4_)[default__.safeIndex(609, (d_2322_v4_).length(0))]), globalState)
                d_2349_v24_ = out66_
                d_2350_v25_ = out67_
                d_2351_v26_ = out68_
                r2 = (self).f28
                d_2352_v27_: D0
                d_2352_v27_ = D0_DC2((self).f24, (d_2322_v4_)[default__.safeIndex(609, (d_2322_v4_).length(0))], (self).f24)
                d_2347_cf41_ = (d_2352_v27_).cf4
                index332_ = default__.safeIndex(654, ((self).f25).length(0))
                ((self).f25)[index332_] = d_2346_cf42_
                index333_ = default__.safeIndex(654, ((self).f25).length(0))
                ((self).f25)[index333_] = self.f23
            elif source33_.is_DC24:
                d_2353___mcc_h3_ = source33_.cf40
                d_2354_cf40_ = d_2353___mcc_h3_
                d_2355_v28_: _dafny.Array
                nw411_ = _dafny.Array(_dafny.Map({}), 24)
                d_2355_v28_ = nw411_
                d_2356_v29_: _dafny.Array
                nw412_ = _dafny.Array(False, 4)
                d_2356_v29_ = nw412_
                d_2357_v30_: D12
                d_2357_v30_ = D12_DC25((self).f24, self.f23)
                d_2358_v31_: C3
                nw413_ = C3()
                nw413_.ctor__(_dafny.Map({d_2356_v29_: (d_2357_v30_).cf42}), d_2354_cf40_, self.f23, (self).f24)
                d_2358_v31_ = nw413_
                d_2359_v32_: _dafny.Map
                d_2359_v32_ = _dafny.Map({(d_2322_v4_)[default__.safeIndex(609, (d_2322_v4_).length(0))]: d_2358_v31_})
                index334_ = default__.safeIndex(535, (d_2355_v28_).length(0))
                (d_2355_v28_)[index334_] = d_2359_v32_
                d_2360_v34_: _dafny.Set
                d_2360_v34_ = _dafny.Set({(self).fm9(d_2340_v19_, d_2339_i0_, (self).f24, globalState), d_2339_i0_})
                d_2361_v35_: _dafny.Set
                def iife213_():
                    coll95_ = _dafny.Set()
                    compr_95_: int
                    for compr_95_ in _dafny.IntegerRange(-657, 218):
                        d_2362_v33_: int = compr_95_
                        if ((-657) <= (d_2362_v33_)) and ((d_2362_v33_) < (218)):
                            coll95_ = coll95_.union(_dafny.Set([(d_2362_v33_) - (self.f29)]))
                    return _dafny.Set(coll95_)
                d_2361_v35_ = _dafny.Set({not ((self).f24) or (default__.fm2(globalState)), (d_2360_v34_).issubset(iife213_()
                )})
                index335_ = default__.safeIndex(609, (d_2322_v4_).length(0))
                index336_ = default__.safeIndex(535, (d_2355_v28_).length(0))
                rhs352_ = len(d_2361_v35_)
                rhs353_ = d_2359_v32_
                lhs269_ = d_2322_v4_
                lhs270_ = default__.safeIndex(609, (d_2322_v4_).length(0))
                lhs271_ = d_2355_v28_
                lhs272_ = default__.safeIndex(535, (d_2355_v28_).length(0))
                lhs269_[lhs270_] = rhs352_
                lhs271_[lhs272_] = rhs353_
                (globalState).f7 = len(d_2324_v6_)
                d_2363_v36_: D15
                d_2363_v36_ = D15_DC34(default__.fm2(globalState))
                d_2364_v37_: _dafny.Map
                d_2364_v37_ = _dafny.Map({self.f29: (self).f24})
                d_2365_v38_: _dafny.Map
                d_2365_v38_ = _dafny.Map({d_2363_v36_: len(d_2364_v37_)})
                d_2366_v39_: C6
                nw414_ = C6()
                nw414_.ctor__(self.f23, (self).f24)
                d_2366_v39_ = nw414_
                d_2367_v40_: _dafny.Map
                d_2367_v40_ = _dafny.Map({False: d_2366_v39_})
                d_2368_v41_: _dafny.MultiSet
                d_2368_v41_ = _dafny.MultiSet([d_2339_i0_])
                d_2369_v42_: _dafny.Map
                d_2369_v42_ = _dafny.Map({(d_2365_v38_).set(d_2363_v36_, (_dafny.MultiSet([d_2367_v40_, _dafny.Map({(self).f28: d_2366_v39_}), d_2367_v40_, d_2367_v40_, d_2367_v40_])).cardinality): d_2368_v41_})
                d_2370_v43_: _dafny.Set
                d_2370_v43_ = _dafny.Set({d_2324_v6_, (d_2324_v6_).set(default__.safeIndex(self.f29, len(d_2324_v6_)), _dafny.CodePoint('u')), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aeorvcm"))) + (d_2324_v6_), d_2324_v6_})
                d_2371_v44_: _dafny.Map
                d_2371_v44_ = _dafny.Map({(self).f28: d_2356_v29_})
                d_2372_v45_: _dafny.Array
                nw415_ = _dafny.Array(int(0), 10)
                d_2372_v45_ = nw415_
                d_2373_v46_: _dafny.Seq
                d_2373_v46_ = _dafny.SeqWithoutIsStrInference([d_2324_v6_])
                def iife214_():
                    coll96_ = _dafny.Set()
                    compr_96_: _dafny.Seq
                    for compr_96_ in (d_2373_v46_).Elements:
                        d_2374_v47_: _dafny.Seq = compr_96_
                        if (d_2374_v47_) in (d_2373_v46_):
                            coll96_ = coll96_.union(_dafny.Set([d_2374_v47_]))
                    return _dafny.Set(coll96_)
                rhs354_ = d_2369_v42_
                rhs355_ = (True) in (d_2371_v44_)
                rhs356_ = d_2372_v45_
                rhs357_ = iife214_()

                lhs273_ = globalState
                lhs274_ = globalState
                d_2369_v42_ = rhs354_
                lhs273_.f13 = rhs355_
                lhs274_.f22 = rhs356_
                d_2370_v43_ = rhs357_
                index337_ = default__.safeIndex(868, (d_2356_v29_).length(0))
                (d_2356_v29_)[index337_] = (default__.fm0((self).f28, self.f29, globalState)) >= (self.f29)
                d_2375_v48_: _dafny.Seq
                d_2375_v48_ = _dafny.SeqWithoutIsStrInference([not((-418) != (len(d_2324_v6_))), (self).f28, True])
                index338_ = default__.safeIndex(868, (d_2356_v29_).length(0))
                (d_2356_v29_)[index338_] = not((d_2375_v48_)[default__.safeIndex(-95, len(d_2375_v48_))])
            elif True:
                d_2376___mcc_h4_ = source33_.cf43
                d_2377_cf43_ = d_2376___mcc_h4_
                d_2378_v49_: _dafny.Array
                nw416_ = _dafny.Array(False, 23)
                d_2378_v49_ = nw416_
                d_2378_v49_ = d_2378_v49_
                r2 = (self).f24
                d_2379_v50_: _dafny.MultiSet
                d_2379_v50_ = _dafny.MultiSet([False, (self).f24])
                d_2380_v51_: int
                d_2381_v52_: bool
                d_2382_v53_: int
                out69_: int
                out70_: bool
                out71_: int
                out69_, out70_, out71_ = default__.m0(default__.safeDivisionInt((0) - (d_2339_i0_), d_2339_i0_), d_2379_v50_, ((self).f28 if (self).f28 else (self).f28), self.f29, globalState)
                d_2380_v51_ = out69_
                d_2381_v52_ = out70_
                d_2382_v53_ = out71_
                d_2383_v54_: T3
                nw417_ = C8()
                nw417_.ctor__(d_2381_v52_)
                d_2383_v54_ = nw417_
                d_2384_v55_: T3
                d_2384_v55_ = d_2383_v54_
                d_2385_v56_: _dafny.Map
                d_2385_v56_ = _dafny.Map({len(_dafny.Map({298: (d_2322_v4_)[default__.safeIndex(609, (d_2322_v4_).length(0))]})): d_2383_v54_})
                d_2386_v57_: _dafny.Map
                d_2386_v57_ = _dafny.Map({(self).f27: d_2324_v6_})
                d_2387_v58_: D2
                d_2387_v58_ = D2_DC6(d_2386_v57_)
                d_2388_v59_: _dafny.MultiSet
                d_2388_v59_ = _dafny.MultiSet([d_2380_v51_, len(default__.fm45((self).f27, d_2387_v58_, (d_2322_v4_)[default__.safeIndex(609, (d_2322_v4_).length(0))], globalState))])
                d_2389_v60_: _dafny.Array
                nw418_ = _dafny.Array(None, 19)
                nw418_[int(0)] = d_2383_v54_
                nw418_[int(1)] = d_2383_v54_
                nw418_[int(2)] = d_2383_v54_
                nw418_[int(3)] = d_2383_v54_
                nw418_[int(4)] = d_2383_v54_
                nw418_[int(5)] = (d_2384_v55_)
                nw418_[int(6)] = d_2383_v54_
                nw418_[int(7)] = (d_2384_v55_)
                nw418_[int(8)] = d_2383_v54_
                nw418_[int(9)] = d_2383_v54_
                nw418_[int(10)] = d_2383_v54_
                nw418_[int(11)] = d_2383_v54_
                nw418_[int(12)] = d_2383_v54_
                nw418_[int(13)] = ((d_2385_v56_)[((d_2388_v59_)[d_2382_v53_] if (d_2382_v53_) in (d_2388_v59_) else (d_2322_v4_)[default__.safeIndex(609, (d_2322_v4_).length(0))])] if (((d_2388_v59_)[d_2382_v53_] if (d_2382_v53_) in (d_2388_v59_) else (d_2322_v4_)[default__.safeIndex(609, (d_2322_v4_).length(0))])) in (d_2385_v56_) else d_2383_v54_)
                nw418_[int(14)] = (d_2384_v55_)
                nw418_[int(15)] = d_2383_v54_
                nw418_[int(16)] = d_2383_v54_
                nw418_[int(17)] = d_2383_v54_
                nw418_[int(18)] = d_2383_v54_
                d_2389_v60_ = nw418_
                index339_ = default__.safeIndex(707, (d_2389_v60_).length(0))
                (d_2389_v60_)[index339_] = d_2383_v54_
                d_2390_v61_: _dafny.Map
                d_2390_v61_ = _dafny.Map({(d_2383_v54_).f27: False})
                d_2391_v62_: _dafny.Map
                d_2391_v62_ = _dafny.Map({d_2390_v61_: d_2324_v6_})
                d_2392_v63_: _dafny.Seq
                d_2392_v63_ = _dafny.SeqWithoutIsStrInference([d_2324_v6_, ((d_2391_v62_)[d_2390_v61_] if (d_2390_v61_) in (d_2391_v62_) else d_2324_v6_), _dafny.SeqWithoutIsStrInference([self.f23 for d_2393_i1_ in range(default__.abs(311))])])
                d_2394_v64_: _dafny.MultiSet
                d_2394_v64_ = _dafny.MultiSet([d_2340_v19_])
                d_2395_v65_: _dafny.Seq
                d_2395_v65_ = _dafny.SeqWithoutIsStrInference([d_2381_v52_, (self).f24, (self).f27, False])
                d_2396_v67_: _dafny.Map
                d_2396_v67_ = _dafny.Map({self.f23: self.f23})
                d_2397_v68_: _dafny.Map
                d_2397_v68_ = _dafny.Map({(d_2322_v4_)[default__.safeIndex(609, (d_2322_v4_).length(0))]: len(d_2396_v67_)})
                d_2398_v69_: _dafny.Map
                d_2398_v69_ = _dafny.Map({len(d_2397_v68_): (self).f27})
                d_2399_v70_: _dafny.Map
                d_2399_v70_ = _dafny.Map({d_2395_v65_: d_2397_v68_})
                d_2400_v71_: _dafny.Map
                def iife215_():
                    coll97_ = _dafny.Map()
                    compr_97_: int
                    for compr_97_ in (d_2398_v69_).keys.Elements:
                        d_2401_v66_: int = compr_97_
                        if (d_2401_v66_) in (d_2398_v69_):
                            coll97_[default__.safeModuloInt(d_2401_v66_, self.f29)] = len(d_2395_v65_)
                    return _dafny.Map(coll97_)
                d_2400_v71_ = _dafny.Map({(_dafny.Map({d_2395_v65_: iife215_()
                })) | (d_2399_v70_): d_2383_v54_})
                index340_ = default__.safeIndex(707, (d_2389_v60_).length(0))
                rhs358_ = (len(d_2392_v63_)) * (((d_2394_v64_)[d_2340_v19_] if (d_2340_v19_) in (d_2394_v64_) else d_2380_v51_))
                rhs359_ = ((d_2400_v71_)[d_2399_v70_] if (d_2399_v70_) in (d_2400_v71_) else d_2383_v54_)
                lhs275_ = globalState
                lhs276_ = d_2389_v60_
                lhs277_ = default__.safeIndex(707, (d_2389_v60_).length(0))
                lhs275_.f7 = rhs358_
                lhs276_[lhs277_] = rhs359_
            (self).f29 = self.f29
            d_2402_v72_: C1
            nw419_ = C1()
            nw419_.ctor__((self).f28, ((self).f28 if (self).f24 else (self).f28))
            d_2402_v72_ = nw419_
        d_2403_v73_: _dafny.Seq
        d_2403_v73_ = _dafny.SeqWithoutIsStrInference([718, self.f29])
        r0 = ((0) - (len(d_2403_v73_))) < (self.f29)
        r1 = self.f29
        d_2404_v74_: _dafny.Set
        d_2404_v74_ = _dafny.Set({(self).f24})
        d_2405_v75_: _dafny.Seq
        d_2405_v75_ = _dafny.SeqWithoutIsStrInference([(self).f24])
        r2 = not ((self).f28) or ((d_2404_v74_).isdisjoint(_dafny.Set({(d_2405_v75_)[default__.safeIndex(self.f29, len(d_2405_v75_))]})))
        return r0, r1, r2

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        hi15_ = self.f29
        for d_2406_i0_ in range(self.f29, hi15_):
            (globalState).f13 = (self).f28
            d_2407_v0_: _dafny.MultiSet
            d_2407_v0_ = _dafny.MultiSet([(self).f24])
            d_2408_v1_: _dafny.Seq
            d_2408_v1_ = _dafny.SeqWithoutIsStrInference([d_2406_i0_, self.f29, -260])
            d_2409_v2_: int
            d_2410_v3_: bool
            d_2411_v4_: int
            out72_: int
            out73_: bool
            out74_: int
            out72_, out73_, out74_ = default__.m0(self.f29, (d_2407_v0_) - ((d_2407_v0_).set(p0, default__.abs(d_2406_i0_))), (self).f28, default__.fm0((self).f27, (d_2408_v1_)[default__.safeIndex(self.f29, len(d_2408_v1_))], globalState), globalState)
            d_2409_v2_ = out72_
            d_2410_v3_ = out73_
            d_2411_v4_ = out74_
            d_2412_v5_: C6
            nw420_ = C6()
            nw420_.ctor__((self).fm10((self).f27, d_2409_v2_, globalState), default__.fm2(globalState))
            d_2412_v5_ = nw420_
            d_2413_v6_: D23
            d_2413_v6_ = D23_DC54(d_2412_v5_)
            d_2412_v5_ = (d_2413_v6_).cf84
            d_2414_v7_: _dafny.Set
            d_2414_v7_ = _dafny.Set({d_2409_v2_})
            d_2415_v8_: _dafny.Map
            d_2415_v8_ = _dafny.Map({p1: (d_2414_v7_ if not(p0) else d_2414_v7_)})
            d_2416_v10_: _dafny.Map
            d_2416_v10_ = _dafny.Map({d_2406_i0_: d_2409_v2_})
            d_2417_v11_: C4
            nw421_ = C4()
            def iife216_():
                coll98_ = _dafny.Set()
                compr_98_: int
                for compr_98_ in _dafny.IntegerRange(272, 866):
                    d_2418_v9_: int = compr_98_
                    if ((272) <= (d_2418_v9_)) and ((d_2418_v9_) < (866)):
                        coll98_ = coll98_.union(_dafny.Set([(d_2418_v9_) - (d_2411_v4_)]))
                return _dafny.Set(coll98_)
            nw421_.ctor__(len(iife216_()
            ), d_2416_v10_, (self).f25, self.f23, d_2410_v3_)
            d_2417_v11_ = nw421_
            d_2419_v12_: _dafny.Seq
            d_2419_v12_ = _dafny.SeqWithoutIsStrInference([d_2417_v11_])
            d_2420_v13_: _dafny.Seq
            d_2420_v13_ = _dafny.SeqWithoutIsStrInference([p0])
            rhs360_ = (_dafny.Map({self.f23: _dafny.Set({687})})) | ((d_2415_v8_) | (d_2415_v8_))
            rhs361_ = (_dafny.MultiSet(d_2408_v1_)).cardinality
            rhs362_ = default__.safeModuloInt(len(d_2419_v12_), len((d_2420_v13_) + (d_2420_v13_)))
            rhs363_ = d_2410_v3_
            lhs278_ = self
            lhs279_ = globalState
            d_2415_v8_ = rhs360_
            lhs278_.f29 = rhs361_
            r1 = rhs362_
            lhs279_.f13 = rhs363_
        hi16_ = self.f29
        for d_2421_i1_ in range(186, hi16_):
            d_2422_v14_: _dafny.Set
            d_2422_v14_ = _dafny.Set({not((self).f27), p0, p0})
            d_2423_v15_: _dafny.Seq
            d_2423_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eeajmfk"))
            d_2424_v16_: D19
            d_2424_v16_ = D19_DC46((d_2422_v14_) | (d_2422_v14_), d_2423_v15_)
            d_2424_v16_ = d_2424_v16_
            (self).f29 = self.f29
            (globalState).f7 = (d_2421_i1_) + (self.f29)
            (self).f23 = self.f23
        d_2425_v17_: _dafny.Array
        nw422_ = _dafny.Array(_dafny.Set({}), 25)
        d_2425_v17_ = nw422_
        d_2426_v18_: _dafny.Array
        def lambda125_(d_2427_i2_):
            return (self).f28

        init72_ = lambda125_
        nw423_ = _dafny.Array(None, 9)
        for i0_72_ in range(nw423_.length(0)):
            nw423_[i0_72_] = init72_(i0_72_)
        d_2426_v18_ = nw423_
        d_2428_v19_: _dafny.Set
        d_2428_v19_ = _dafny.Set({d_2426_v18_, d_2426_v18_, d_2426_v18_, d_2426_v18_, d_2426_v18_})
        index341_ = default__.safeIndex(958, (d_2425_v17_).length(0))
        (d_2425_v17_)[index341_] = d_2428_v19_
        index342_ = default__.safeIndex(958, (d_2425_v17_).length(0))
        (d_2425_v17_)[index342_] = ((d_2428_v19_ if (self).f27 else _dafny.Set({d_2426_v18_, d_2426_v18_}))).intersection(d_2428_v19_)
        r2 = (self).f24
        d_2429_v20_: _dafny.Array
        nw424_ = _dafny.Array(_dafny.Map({}), 8)
        d_2429_v20_ = nw424_
        d_2430_v21_: _dafny.Array
        nw425_ = _dafny.Array(int(0), 26)
        d_2430_v21_ = nw425_
        d_2431_v22_: _dafny.Map
        d_2431_v22_ = _dafny.Map({(self).f27: d_2430_v21_})
        index343_ = default__.safeIndex(655, (d_2429_v20_).length(0))
        (d_2429_v20_)[index343_] = (d_2431_v22_) | (_dafny.Map({(self).f27: d_2430_v21_}))
        index344_ = default__.safeIndex(655, (d_2429_v20_).length(0))
        (d_2429_v20_)[index344_] = ((_dafny.Map({not(p0): d_2430_v21_})) | (d_2431_v22_)) | (_dafny.Map({(self).f27: d_2430_v21_}))
        hi17_ = self.f29
        for d_2432_i3_ in range(self.f29, hi17_):
            d_2433_v23_: _dafny.Set
            d_2433_v23_ = _dafny.Set({(self).f24})
            d_2434_v24_: C5
            nw426_ = C5()
            nw426_.ctor__(d_2426_v18_)
            d_2434_v24_ = nw426_
            d_2435_v25_: _dafny.Set
            d_2435_v25_ = _dafny.Set({d_2434_v24_, d_2434_v24_})
            rhs364_ = d_2430_v21_
            rhs365_ = True
            rhs366_ = (0) - (len(((d_2433_v23_ if (self).f24 else d_2433_v23_)) | (d_2433_v23_)))
            rhs367_ = ((default__.fm2(globalState) if False else p0) if (self).f27 else (d_2435_v25_).ispropersubset(d_2435_v25_))
            lhs280_ = globalState
            lhs281_ = self
            lhs280_.f22 = rhs364_
            r0 = rhs365_
            lhs281_.f29 = rhs366_
            r2 = rhs367_
            d_2436_v26_: _dafny.Seq
            d_2436_v26_ = _dafny.SeqWithoutIsStrInference([self.f29])
            d_2437_v27_: T1
            nw427_ = C6()
            nw427_.ctor__(p1, p0)
            d_2437_v27_ = nw427_
            d_2438_v28_: _dafny.Array
            nw428_ = _dafny.Array(None, 17)
            nw428_[int(0)] = d_2437_v27_
            nw428_[int(1)] = d_2437_v27_
            nw428_[int(2)] = d_2437_v27_
            nw428_[int(3)] = d_2437_v27_
            nw428_[int(4)] = d_2437_v27_
            nw428_[int(5)] = d_2437_v27_
            nw428_[int(6)] = d_2437_v27_
            nw428_[int(7)] = d_2437_v27_
            nw428_[int(8)] = (d_2437_v27_ if p0 else d_2437_v27_)
            nw428_[int(9)] = d_2437_v27_
            nw428_[int(10)] = d_2437_v27_
            nw428_[int(11)] = d_2437_v27_
            nw428_[int(12)] = d_2437_v27_
            nw428_[int(13)] = d_2437_v27_
            nw428_[int(14)] = d_2437_v27_
            nw428_[int(15)] = d_2437_v27_
            nw428_[int(16)] = d_2437_v27_
            d_2438_v28_ = nw428_
            index345_ = default__.safeIndex(378, (d_2438_v28_).length(0))
            (d_2438_v28_)[index345_] = d_2437_v27_
            d_2439_v29_: _dafny.Seq
            d_2439_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfbyqdoi"))
            d_2440_v30_: _dafny.Seq
            d_2440_v30_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2441_v32_: D5
            def iife217_():
                coll99_ = _dafny.Set()
                compr_99_: int
                for compr_99_ in (_dafny.SeqWithoutIsStrInference([len(d_2439_v29_), len(d_2440_v30_)])).Elements:
                    d_2442_v31_: int = compr_99_
                    if (d_2442_v31_) in (_dafny.SeqWithoutIsStrInference([len(d_2439_v29_), len(d_2440_v30_)])):
                        coll99_ = coll99_.union(_dafny.Set([(d_2442_v31_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ru"))))]))
                return _dafny.Set(coll99_)
            d_2441_v32_ = D5_DC12(iife217_()
)
            d_2443_v33_: _dafny.Map
            d_2443_v33_ = _dafny.Map({p0: d_2439_v29_})
            d_2444_v34_: D2
            d_2444_v34_ = D2_DC6(d_2443_v33_)
            index346_ = default__.safeIndex(378, (d_2438_v28_).length(0))
            rhs368_ = (_dafny.SeqWithoutIsStrInference([d_2432_i3_])) + (default__.fm45((self).f28, d_2444_v34_, 633, globalState))
            rhs369_ = d_2437_v27_.f23
            rhs370_ = d_2437_v27_
            rhs371_ = (d_2439_v29_) + (d_2439_v29_)
            rhs372_ = (d_2441_v32_ if default__.fm2(globalState) else d_2441_v32_)
            lhs282_ = self
            lhs283_ = d_2438_v28_
            lhs284_ = default__.safeIndex(378, (d_2438_v28_).length(0))
            d_2436_v26_ = rhs368_
            lhs282_.f23 = rhs369_
            lhs283_[lhs284_] = rhs370_
            d_2439_v29_ = rhs371_
            d_2441_v32_ = rhs372_
            if not (False) or ((self).f27):
                (self).f23 = d_2437_v27_.f23
                (self).f29 = (0) - (d_2432_i3_)
                d_2445_v35_: D1
                d_2445_v35_ = D1_DC4(d_2437_v27_.f23)
                (self).f23 = (d_2445_v35_).cf6
                d_2446_v36_: _dafny.Array
                def lambda126_(d_2447_v29_):
                    def lambda127_(d_2448_i4_):
                        return d_2447_v29_

                    return lambda127_

                init73_ = lambda126_(d_2439_v29_)
                nw429_ = _dafny.Array(None, 12)
                for i0_73_ in range(nw429_.length(0)):
                    nw429_[i0_73_] = init73_(i0_73_)
                d_2446_v36_ = nw429_
                d_2449_v37_: _dafny.Map
                d_2449_v37_ = _dafny.Map({d_2439_v29_: d_2439_v29_})
                index347_ = default__.safeIndex(455, (d_2446_v36_).length(0))
                (d_2446_v36_)[index347_] = ((d_2449_v37_)[d_2439_v29_] if (d_2439_v29_) in (d_2449_v37_) else d_2439_v29_)
                index348_ = default__.safeIndex(455, (d_2446_v36_).length(0))
                (d_2446_v36_)[index348_] = (d_2439_v29_) + ((d_2439_v29_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lydhu"))))
                (globalState).f7 = self.f29
            elif True:
                d_2450_v38_: _dafny.MultiSet
                d_2450_v38_ = _dafny.MultiSet([(self).f24, (d_2437_v27_).f24, True])
                d_2451_v39_: _dafny.Map
                d_2451_v39_ = _dafny.Map({D0_DC0(d_2450_v38_): (d_2439_v29_) + (d_2439_v29_)})
                d_2452_v40_: D0
                d_2452_v40_ = D0_DC0(d_2450_v38_)
                d_2451_v39_ = (d_2451_v39_).set(d_2452_v40_, (d_2439_v29_) + (_dafny.SeqWithoutIsStrInference([d_2437_v27_.f23 for d_2453_i5_ in range(default__.abs(263))])))
                r1 = (self.f29) * (968)
                d_2454_v41_: _dafny.Array
                nw430_ = _dafny.Array(None, 7)
                nw430_[int(0)] = d_2437_v27_.f23
                nw430_[int(1)] = self.f23
                nw430_[int(2)] = d_2437_v27_.f23
                nw430_[int(3)] = default__.fm29(d_2432_i3_, (self).f28, (self).f27, (0) - (d_2432_i3_), globalState)
                nw430_[int(4)] = _dafny.CodePoint('w')
                nw430_[int(5)] = d_2437_v27_.f23
                nw430_[int(6)] = self.f23
                d_2454_v41_ = nw430_
                d_2454_v41_ = d_2454_v41_
                (globalState).f13 = (default__.safeDivisionInt((0) - ((d_2436_v26_)[default__.safeIndex(self.f29, len(d_2436_v26_))]), 471)) != (552)
                d_2455_v42_: _dafny.Map
                d_2455_v42_ = _dafny.Map({d_2432_i3_: d_2439_v29_})
                d_2455_v42_ = (d_2455_v42_).set(default__.safeModuloInt(self.f29, len(d_2433_v23_)), d_2439_v29_)
            index349_ = default__.safeIndex(766, (d_2426_v18_).length(0))
            (d_2426_v18_)[index349_] = ((self).f28) and (p0)
            d_2456_v43_: _dafny.Map
            d_2456_v43_ = _dafny.Map({(self).f28: (self).f27})
            d_2457_v44_: _dafny.Seq
            d_2457_v44_ = _dafny.SeqWithoutIsStrInference([d_2456_v43_, d_2456_v43_, d_2456_v43_, d_2456_v43_])
            d_2458_v45_: _dafny.Map
            d_2458_v45_ = _dafny.Map({d_2430_v21_: (d_2457_v44_)[default__.safeIndex(len(d_2433_v23_), len(d_2457_v44_))]})
            index350_ = default__.safeIndex(619, ((d_2434_v24_).f32).length(0))
            ((d_2434_v24_).f32)[index350_] = (self).f28
            d_2459_v46_: _dafny.Map
            d_2459_v46_ = _dafny.Map({(d_2434_v24_).f32: p1})
            d_2460_v47_: C3
            nw431_ = C3()
            nw431_.ctor__(d_2459_v46_, (self).f25, self.f23, (self).f27)
            d_2460_v47_ = nw431_
            d_2461_v48_: _dafny.Set
            d_2461_v48_ = _dafny.Set({d_2460_v47_, d_2460_v47_, d_2460_v47_, d_2460_v47_, d_2460_v47_})
            index351_ = default__.safeIndex(766, (d_2426_v18_).length(0))
            index352_ = default__.safeIndex(619, ((d_2434_v24_).f32).length(0))
            rhs373_ = not((916) != ((self.f29) + (d_2432_i3_)))
            rhs374_ = d_2458_v45_
            rhs375_ = (d_2437_v27_).fm4(len(d_2461_v48_), default__.fm36(globalState), globalState)
            lhs285_ = d_2426_v18_
            lhs286_ = default__.safeIndex(766, (d_2426_v18_).length(0))
            lhs287_ = (d_2434_v24_).f32
            lhs288_ = default__.safeIndex(619, ((d_2434_v24_).f32).length(0))
            lhs285_[lhs286_] = rhs373_
            d_2458_v45_ = rhs374_
            lhs287_[lhs288_] = rhs375_
        r0 = (self).f28
        r1 = self.f29
        d_2462_v49_: D6
        d_2462_v49_ = D6_DC16(_dafny.Map({False: p0}), (self).f27, p0)
        r2 = (d_2462_v49_).cf26
        return r0, r1, r2

    def m4(self, p0, globalState):
        r0: bool = False
        d_2463_v0_: _dafny.Seq
        d_2463_v0_ = _dafny.SeqWithoutIsStrInference([self.f29, self.f29])
        d_2464_v1_: _dafny.Map
        d_2464_v1_ = _dafny.Map({d_2463_v0_: self.f23})
        d_2465_v2_: _dafny.MultiSet
        d_2465_v2_ = _dafny.MultiSet([((d_2464_v1_)[d_2463_v0_] if (d_2463_v0_) in (d_2464_v1_) else self.f23), self.f23])
        d_2465_v2_ = d_2465_v2_
        d_2466_v3_: _dafny.Map
        d_2466_v3_ = _dafny.Map({(self).f24: p0})
        d_2467_v4_: _dafny.Map
        d_2467_v4_ = _dafny.Map({((d_2466_v3_)[(self).f24] if ((self).f24) in (d_2466_v3_) else (self).f28): (self.f29) - (self.f29)})
        d_2468_v5_: _dafny.MultiSet
        d_2468_v5_ = _dafny.MultiSet([(self).f28, (self).f28, not((self).f27), (self).f28])
        d_2469_v6_: _dafny.Map
        d_2469_v6_ = _dafny.Map({(d_2468_v5_).cardinality: (self).f24})
        d_2467_v4_ = (d_2467_v4_).set((len(d_2469_v6_)) >= ((default__.fm3((self).f24, globalState)).cardinality), self.f29)
        (globalState).f7 = self.f29
        (self).f29 = (self).fm6(globalState)
        d_2470_v7_: _dafny.Array
        nw432_ = _dafny.Array(None, 17)
        d_2470_v7_ = nw432_
        d_2471_v8_: _dafny.Seq
        d_2471_v8_ = _dafny.SeqWithoutIsStrInference([d_2470_v7_])
        d_2472_v9_: _dafny.Array
        nw433_ = _dafny.Array(None, 13)
        nw433_[int(0)] = d_2470_v7_
        nw433_[int(1)] = (d_2471_v8_)[default__.safeIndex(self.f29, len(d_2471_v8_))]
        nw433_[int(2)] = d_2470_v7_
        nw433_[int(3)] = d_2470_v7_
        nw433_[int(4)] = d_2470_v7_
        nw433_[int(5)] = d_2470_v7_
        nw433_[int(6)] = d_2470_v7_
        nw433_[int(7)] = d_2470_v7_
        nw433_[int(8)] = d_2470_v7_
        nw433_[int(9)] = d_2470_v7_
        nw433_[int(10)] = d_2470_v7_
        nw433_[int(11)] = d_2470_v7_
        nw433_[int(12)] = d_2470_v7_
        d_2472_v9_ = nw433_
        index353_ = default__.safeIndex(380, (d_2472_v9_).length(0))
        (d_2472_v9_)[index353_] = d_2470_v7_
        index354_ = default__.safeIndex(380, (d_2472_v9_).length(0))
        rhs376_ = (self).f24
        rhs377_ = d_2470_v7_
        lhs289_ = d_2472_v9_
        lhs290_ = default__.safeIndex(380, (d_2472_v9_).length(0))
        r0 = rhs376_
        lhs289_[lhs290_] = rhs377_
        d_2473_v10_: _dafny.MultiSet
        d_2473_v10_ = _dafny.MultiSet([self.f29, (d_2468_v5_).cardinality])
        d_2474_v11_: _dafny.Map
        d_2474_v11_ = _dafny.Map({self.f29: ((d_2473_v10_) - (d_2473_v10_)).cardinality})
        d_2475_v12_: D19
        d_2475_v12_ = D19_DC47((self).f24)
        d_2476_v13_: _dafny.Array
        nw434_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
        d_2476_v13_ = nw434_
        d_2477_v14_: _dafny.Seq
        d_2477_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "myrcehyl"))
        index355_ = default__.safeIndex(244, (d_2476_v13_).length(0))
        (d_2476_v13_)[index355_] = d_2477_v14_
        d_2478_v15_: _dafny.Seq
        d_2478_v15_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f28])
        d_2479_v16_: _dafny.Set
        d_2479_v16_ = _dafny.Set({(d_2478_v15_)[default__.safeIndex(self.f29, len(d_2478_v15_))]})
        d_2480_v17_: D0
        d_2480_v17_ = D0_DC2((self).f28, len(d_2466_v3_), not(default__.fm2(globalState)))
        index356_ = default__.safeIndex(244, (d_2476_v13_).length(0))
        rhs378_ = ((d_2474_v11_).set((0) - (self.f29), self.f29)) | (_dafny.Map({self.f29: self.f29}))
        rhs379_ = default__.fm57(globalState)
        rhs380_ = (d_2480_v17_).cf3
        rhs381_ = d_2477_v14_
        rhs382_ = (d_2479_v16_).intersection(d_2479_v16_)
        lhs291_ = globalState
        lhs292_ = d_2476_v13_
        lhs293_ = default__.safeIndex(244, (d_2476_v13_).length(0))
        d_2474_v11_ = rhs378_
        d_2475_v12_ = rhs379_
        lhs291_.f7 = rhs380_
        lhs292_[lhs293_] = rhs381_
        d_2479_v16_ = rhs382_
        r0 = (self).f24
        return r0

    @property
    def f28(self):
        return self._f28
