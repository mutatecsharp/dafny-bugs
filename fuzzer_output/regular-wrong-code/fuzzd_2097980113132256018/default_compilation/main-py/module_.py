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
    def fm4(p0, globalState):
        return (D7_DC17(len(_dafny.Map({691: 995})), not(not(False)), True, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ruhwlvnnb")))))).cf27

    @staticmethod
    def fm6(p0, globalState):
        return not((_dafny.Set({_dafny.Set({not(True), False, True}), _dafny.Set({False}), _dafny.Set({not(True)})})).ispropersubset(_dafny.Set({_dafny.Set({True, True}), _dafny.Set({True})})))

    @staticmethod
    def fm8(p0, globalState):
        return _dafny.CodePoint('y')

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        return D0_DC1(not(True), (_dafny.CodePoint('y') if False else _dafny.CodePoint('j')))

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return (_dafny.Map({D0_DC1(False, _dafny.CodePoint('x')): True})) | (_dafny.Map({D0_DC1(False, _dafny.CodePoint('v')): True}))

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        source0_ = (D23_DC59(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhpqe")), 300, 162) if not(not(False)) else D23_DC59(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_0_i0_ in range(default__.abs(380))]), -580, (0) - ((_dafny.MultiSet([583])).cardinality)))
        if source0_.is_DC59:
            d_1___mcc_h0_ = source0_.cf93
            d_2___mcc_h1_ = source0_.cf94
            d_3___mcc_h2_ = source0_.cf95
            d_4_cf95_ = d_3___mcc_h2_
            d_5_cf94_ = d_2___mcc_h1_
            d_6_cf93_ = d_1___mcc_h0_
            return (_dafny.Map({d_6_cf93_: d_5_cf94_})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")): d_4_cf95_}))
        elif source0_.is_DC60:
            d_7___mcc_h3_ = source0_.cf96
            d_8___mcc_h4_ = source0_.cf97
            d_9_cf97_ = d_8___mcc_h4_
            d_10_cf96_ = d_7___mcc_h3_
            def iife0_():
                coll0_ = _dafny.Map()
                compr_0_: _dafny.Seq
                for compr_0_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvbfh")) for d_11_i1_ in range(default__.abs(248))])).Elements:
                    d_12_v0_: _dafny.Seq = compr_0_
                    if (d_12_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvbfh")) for d_11_i1_ in range(default__.abs(248))])):
                        coll0_[d_12_v0_] = 447
                return _dafny.Map(coll0_)
            return (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bk")): len(_dafny.Map({d_9_cf97_: False}))})) | (iife0_()
            )
        elif source0_.is_DC61:
            d_13___mcc_h5_ = source0_.cf98
            d_14_cf98_ = d_13___mcc_h5_
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: _dafny.Seq
                for compr_1_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aubgkgk"))])).Elements:
                    d_15_v1_: _dafny.Seq = compr_1_
                    if (d_15_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aubgkgk"))])):
                        coll1_[d_15_v1_] = d_14_cf98_
                return _dafny.Map(coll1_)
            return iife1_()
            
        elif True:
            d_16___mcc_h6_ = source0_.cf92
            d_17_cf92_ = d_16___mcc_h6_
            return _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_18_i2_ in range(default__.abs(694))]): -715})

    @staticmethod
    def fm14(globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eb"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mehe")))) < ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_19_i0_ in range(default__.abs(-217))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lypkvvier"))))

    @staticmethod
    def fm15(p0, p1, p2, p3, globalState):
        return (_dafny.Map({len(_dafny.Map({False: not(False)})): True})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgftfxi"))): False}))

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        return not((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: False})) for d_20_i0_ in range(default__.abs(971))])) < (_dafny.SeqWithoutIsStrInference([485, 908, len(_dafny.Set({14})), 375])))

    @staticmethod
    def fm19(globalState):
        return _dafny.SeqWithoutIsStrInference([705])

    @staticmethod
    def fm20(p0, globalState):
        return (_dafny.MultiSet([False])) | (_dafny.MultiSet([False, False, False]))

    @staticmethod
    def fm21(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjryf"))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ri")))

    @staticmethod
    def fm22(p0, p1, p2, globalState):
        return _dafny.MultiSet([(_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, not(True)]))).cardinality: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmue"))))})) | (_dafny.Map({len(_dafny.Set({308})): -232})), _dafny.Map({280: 540})])

    @staticmethod
    def fm23(p0, p1, globalState):
        source1_ = D10_DC23()
        if source1_.is_DC23:
            if False:
                return _dafny.CodePoint('f')
            elif True:
                return _dafny.CodePoint('e')
        elif True:
            d_21___mcc_h0_ = source1_.cf36
            d_22_cf36_ = d_21___mcc_h0_
            return _dafny.CodePoint('v')

    @staticmethod
    def fm24(globalState):
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptxr")))]) if not(True) else _dafny.SeqWithoutIsStrInference([622])) if (True if False else not(not(False))) else (_dafny.SeqWithoutIsStrInference([-587 for d_23_i0_ in range(default__.abs(107))])) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality]))))

    @staticmethod
    def fm25(globalState):
        source2_ = (D11_DC24(_dafny.Map({len(_dafny.Map({349: False})): (_dafny.MultiSet([850])).cardinality})) if False else D11_DC24(_dafny.Map({(_dafny.MultiSet([len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hflpwwb"))})), 749, -493])).cardinality: 856})))
        if source2_.is_DC25:
            d_24___mcc_h0_ = source2_.cf38
            d_25_cf38_ = d_24___mcc_h0_
            if True:
                return _dafny.Map({-330: len(_dafny.Map({not(True): not(True)}))})
            elif True:
                return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pneydj"))): 553})
        elif True:
            d_26___mcc_h1_ = source2_.cf37
            d_27_cf37_ = d_26___mcc_h1_
            return d_27_cf37_

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        return _dafny.Map({not((not(not(True))) or (False)): True})

    @staticmethod
    def fm27(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: str
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l'), _dafny.CodePoint('v')])).Elements:
                d_28_v0_: str = compr_2_
                if (d_28_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l'), _dafny.CodePoint('v')])):
                    coll2_ = coll2_.union(_dafny.Set([d_28_v0_]))
            return _dafny.Set(coll2_)
        return iife2_()
        

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        return _dafny.Map({(_dafny.Set({False})) == (_dafny.Set({True})): (0) - ((-977) * (len(_dafny.Set({168, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljpjdt"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwdyqpvqq"))), -599}))))})

    @staticmethod
    def fm29(p0, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: str
            for compr_3_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_30_i1_ in range(default__.abs(376))]))).Elements:
                d_31_v0_: str = compr_3_
                if (d_31_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_30_i1_ in range(default__.abs(376))]))):
                    coll3_[d_31_v0_] = 640
            return _dafny.Map(coll3_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-25]))]) for d_29_i0_ in range(default__.abs(-939))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([153, (0) - (len(iife3_()
        ))])]))

    @staticmethod
    def fm30(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_32_i0_ in range(default__.abs(990))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))

    @staticmethod
    def fm31(p0, p1, p2, p3, globalState):
        def lambda0_(source3_):
            if source3_.is_DC25:
                d_33___mcc_h0_ = source3_.cf38
                d_34_cf38_ = d_33___mcc_h0_
                return (d_34_cf38_).set(-436, False)
            elif True:
                d_35___mcc_h1_ = source3_.cf37
                d_36_cf37_ = d_35___mcc_h1_
                def iife4_():
                    coll4_ = _dafny.Map()
                    compr_4_: int
                    for compr_4_ in (_dafny.SeqWithoutIsStrInference([-199])).Elements:
                        d_37_v0_: int = compr_4_
                        if (d_37_v0_) in (_dafny.SeqWithoutIsStrInference([-199])):
                            coll4_[(d_37_v0_) + ((_dafny.MultiSet([(D19_DC45(len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.Set({True})), True, False, _dafny.SeqWithoutIsStrInference([True, False]))).cf71])).cardinality)] = False
                    return _dafny.Map(coll4_)
                return iife4_()
                

        return len(lambda0_(D11_DC24(_dafny.Map({943: 383}))))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(634, 314):
                d_39_v0_: int = compr_5_
                if ((634) <= (d_39_v0_)) and ((d_39_v0_) < (314)):
                    coll5_[default__.safeDivisionInt(d_39_v0_, (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crmga"))))))] = False
            return _dafny.Map(coll5_)
        return _dafny.Set({(0) - (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([not(False), True, False])), len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-693 for d_38_i0_ in range(default__.abs(-555))]), _dafny.SeqWithoutIsStrInference([len(iife5_()
        ), len(_dafny.SeqWithoutIsStrInference([True, True]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_40_i2_ in range(default__.abs(130))])) for d_41_i1_ in range(default__.abs(572))])])))), -951, (739) * (313)})

    @staticmethod
    def fm33(p0, p1, globalState):
        def iife6_():
            coll6_ = _dafny.Set()
            compr_6_: int
            for compr_6_ in (_dafny.SeqWithoutIsStrInference([322])).Elements:
                d_42_v0_: int = compr_6_
                if (d_42_v0_) in (_dafny.SeqWithoutIsStrInference([322])):
                    coll6_ = coll6_.union(_dafny.Set([default__.safeDivisionInt(d_42_v0_, 244)]))
            return _dafny.Set(coll6_)
        return (_dafny.Set({True, (D22_DC56((0) - (len(iife6_()
)), False)).cf90, True})) | (_dafny.Set({True}))

    @staticmethod
    def fm34(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([True, False])) + (_dafny.SeqWithoutIsStrInference([not(False), False]))

    @staticmethod
    def fm35(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yg"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "strrna"))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_43_i0_ in range(default__.abs(1))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ifw"))])))

    @staticmethod
    def fm36(p0, p1, globalState):
        return _dafny.Map({(500) + (len(_dafny.SeqWithoutIsStrInference([not(False)]))): _dafny.Set({True, True, not(True)})})

    @staticmethod
    def fm37(p0, globalState):
        if ((0) - ((0) - ((0) - (len(_dafny.Set({len(_dafny.Map({False: len(_dafny.Map({_dafny.Map({not(True): False}): True}))})), 607, (0) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_44_i0_ in range(default__.abs(380))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_45_i1_ in range(default__.abs(995))])])).cardinality)})))))) == (-718):
            return D16_DC36(_dafny.Set({-887}), not(True), True, False)
        elif True:
            return D16_DC36(_dafny.Set({141}), False, False, not(False))

    @staticmethod
    def fm38(p0, p1, globalState):
        return D12_DC28(D12_DC28(D12_DC28(D12_DC28(D12_DC27(-792)))))

    @staticmethod
    def fm39(p0, p1, p2, p3, globalState):
        return D3_DC10((D3_DC10(_dafny.MultiSet([True, False]), D1_DC4(-992))).cf16, D1_DC4(len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_46_i0_ in range(default__.abs(575))])): 12}))))

    @staticmethod
    def fm40(p0, p1, p2, p3, globalState):
        return (_dafny.Map({False: _dafny.CodePoint('i')})) | (_dafny.Map({True: _dafny.CodePoint('l')}))

    @staticmethod
    def fm41(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emqvugbr")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwm"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "byfhiw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thtg"))]))

    @staticmethod
    def fm42(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([D4_DC12(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpkjcwei"))) for d_47_i0_ in range(default__.abs(391))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkajua"))), len(_dafny.SeqWithoutIsStrInference([True]))])), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(True): False}))])), False)])) - (_dafny.MultiSet([D4_DC12((_dafny.MultiSet([False])).cardinality, -41, (D22_DC56(560, False)).cf90)]))) - (_dafny.MultiSet((D27_DC68(_dafny.SeqWithoutIsStrInference([D4_DC12((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pw")))), -595, True), D4_DC12(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmeenbgie"))), 885, True)]))).cf110))

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        return D4_DC13(D4_DC11(_dafny.SeqWithoutIsStrInference([False, True])))

    @staticmethod
    def fm44(p0, p1, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: _dafny.Seq
            for compr_7_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([374])])).Elements:
                d_48_v0_: _dafny.Seq = compr_7_
                if (d_48_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([374])])):
                    coll7_[d_48_v0_] = _dafny.CodePoint('o')
            return _dafny.Map(coll7_)
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: _dafny.Seq
            for compr_8_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('a'): (_dafny.MultiSet([846, -631])).cardinality})), 542]): len(_dafny.SeqWithoutIsStrInference([not(not(False)), False, False]))})).keys.Elements:
                d_49_v1_: _dafny.Seq = compr_8_
                if (d_49_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('a'): (_dafny.MultiSet([846, -631])).cardinality})), 542]): len(_dafny.SeqWithoutIsStrInference([not(not(False)), False, False]))})):
                    coll8_[d_49_v1_] = _dafny.CodePoint('f')
            return _dafny.Map(coll8_)
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: _dafny.Seq
            for compr_9_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-617]) for d_51_i1_ in range(default__.abs(73))]))).Elements:
                d_52_v2_: _dafny.Seq = compr_9_
                if (d_52_v2_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-617]) for d_51_i1_ in range(default__.abs(73))]))):
                    coll9_[d_52_v2_] = _dafny.CodePoint('i')
            return _dafny.Map(coll9_)
        return _dafny.Set({(_dafny.Set({D17_DC37(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))), 55]): _dafny.CodePoint('b')})), D17_DC37(_dafny.Map({_dafny.SeqWithoutIsStrInference([456, (0) - (len(_dafny.SeqWithoutIsStrInference([False, False])))]): _dafny.CodePoint('x')})), D17_DC37(iife7_()
)})) - (_dafny.Set({D17_DC37(iife8_()
), D17_DC37(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcdjuq"))) for d_50_i0_ in range(default__.abs(615))]): _dafny.CodePoint('q')}))})), _dafny.Set({D17_DC37(iife9_()
)})})

    @staticmethod
    def fm45(p0, p1, globalState):
        return _dafny.Map({_dafny.CodePoint('i'): _dafny.CodePoint('u')})

    @staticmethod
    def fm46(globalState):
        return _dafny.Map({(499) < ((0) - ((D1_DC5(False, False, (_dafny.MultiSet([True])).cardinality, (0) - (-179))).cf11)): (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_53_i0_ in range(default__.abs(47))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwrbjb")))})

    @staticmethod
    def fm47(p0, p1, p2, p3, globalState):
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: _dafny.Seq
            for compr_10_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ji")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yns")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_54_i0_ in range(default__.abs(934))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_55_i1_ in range(default__.abs(734))])]))).Elements:
                d_56_v0_: _dafny.Seq = compr_10_
                if (d_56_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ji")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yns")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_54_i0_ in range(default__.abs(934))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_55_i1_ in range(default__.abs(734))])]))):
                    coll10_[d_56_v0_] = False
            return _dafny.Map(coll10_)
        return _dafny.Map({_dafny.Set({len(iife10_()
        )}): 630})

    @staticmethod
    def fm48(p0, globalState):
        if (_dafny.Set({not(True)})).ispropersubset(_dafny.Set({True})):
            return D1_DC5(True, True, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False})) for d_57_i0_ in range(default__.abs(-114))])), 941, -254, 841})), 117)
        elif True:
            return D1_DC5(not(True), True, (_dafny.MultiSet([(_dafny.MultiSet([len(_dafny.Map({len(_dafny.Map({not(False): _dafny.CodePoint('x')})): False}))])).cardinality, -864])).cardinality, len(_dafny.Map({_dafny.CodePoint('n'): False})))

    @staticmethod
    def fm49(p0, globalState):
        return D23_DC59(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_58_i0_ in range(default__.abs(507))]), 8, (len(_dafny.Set({True, True}))) * (487))

    @staticmethod
    def fm50(p0, globalState):
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: _dafny.Seq
            for compr_11_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-751]))]): True})).keys.Elements:
                d_59_v0_: _dafny.Seq = compr_11_
                if (d_59_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-751]))]): True})):
                    coll11_[d_59_v0_] = _dafny.CodePoint('p')
            return _dafny.Map(coll11_)
        return iife11_()
        

    @staticmethod
    def fm51(globalState):
        return D0_DC2((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True]))).ispropersubset(_dafny.MultiSet([False, True])), (len(_dafny.SeqWithoutIsStrInference([False])) if True else len(_dafny.Set({True}))), (0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yj"))) if False else (_dafny.MultiSet([not(True)])).cardinality)))

    @staticmethod
    def m12(p0, p1, p2, p3, globalState):
        r0: D11 = D11.default()()
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.Seq({})
        r3: C2 = None
        d_60_v0_: _dafny.Seq
        d_60_v0_ = _dafny.SeqWithoutIsStrInference([p2, not(not(p2))])
        hi0_ = default__.fm31(p2, p0, not(p3), d_60_v0_, globalState)
        for d_61_i0_ in range(p0, hi0_):
            r1 = (d_61_i0_) + ((d_61_i0_) - (p0))
            d_62_v1_: _dafny.Seq
            d_62_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wowdbx"))
            d_63_v2_: _dafny.Array
            def lambda1_(d_64_p2_):
                def lambda2_(d_65_i1_):
                    return d_64_p2_

                return lambda2_

            init0_ = lambda1_(p2)
            nw0_ = _dafny.Array(None, 16)
            for i0_0_ in range(nw0_.length(0)):
                nw0_[i0_0_] = init0_(i0_0_)
            d_63_v2_ = nw0_
            d_66_v3_: _dafny.Map
            d_66_v3_ = _dafny.Map({p0: d_63_v2_})
            d_67_v4_: C1
            nw1_ = C1()
            nw1_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")), p3)
            d_67_v4_ = nw1_
            d_68_v5_: D15
            d_68_v5_ = D15_DC33(d_63_v2_, d_67_v4_)
            d_69_v6_: _dafny.Seq
            d_69_v6_ = _dafny.SeqWithoutIsStrInference([d_63_v2_])
            d_70_v7_: _dafny.Array
            nw2_ = _dafny.Array(None, 18)
            nw2_[int(0)] = ((d_66_v3_)[d_61_i0_] if (d_61_i0_) in (d_66_v3_) else d_63_v2_)
            nw2_[int(1)] = d_63_v2_
            nw2_[int(2)] = d_63_v2_
            nw2_[int(3)] = d_63_v2_
            nw2_[int(4)] = d_63_v2_
            nw2_[int(5)] = d_63_v2_
            nw2_[int(6)] = d_63_v2_
            nw2_[int(7)] = (d_68_v5_).cf47
            nw2_[int(8)] = d_63_v2_
            nw2_[int(9)] = d_63_v2_
            nw2_[int(10)] = d_63_v2_
            nw2_[int(11)] = (d_63_v2_ if p2 else d_63_v2_)
            nw2_[int(12)] = d_63_v2_
            nw2_[int(13)] = d_63_v2_
            nw2_[int(14)] = (d_69_v6_)[default__.safeIndex(-869, len(d_69_v6_))]
            nw2_[int(15)] = d_63_v2_
            nw2_[int(16)] = d_63_v2_
            nw2_[int(17)] = d_63_v2_
            d_70_v7_ = nw2_
            index0_ = default__.safeIndex(44, (d_70_v7_).length(0))
            (d_70_v7_)[index0_] = d_63_v2_
            d_71_v8_: C3
            nw3_ = C3()
            nw3_.ctor__((d_67_v4_).f13, p2)
            d_71_v8_ = nw3_
            d_72_v9_: D22
            d_72_v9_ = D22_DC55(d_71_v8_)
            index1_ = default__.safeIndex(44, (d_70_v7_).length(0))
            rhs0_ = _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('m') if (d_67_v4_).f13 else _dafny.CodePoint('o')) for d_73_i2_ in range(default__.abs(753))])
            rhs1_ = d_63_v2_
            rhs2_ = p3
            rhs3_ = d_63_v2_
            rhs4_ = d_72_v9_
            lhs0_ = d_70_v7_
            lhs1_ = default__.safeIndex(44, (d_70_v7_).length(0))
            lhs2_ = globalState
            d_62_v1_ = rhs0_
            lhs0_[lhs1_] = rhs1_
            lhs2_.f0 = rhs2_
            d_63_v2_ = rhs3_
            d_72_v9_ = rhs4_
            d_74_v10_: C4
            nw4_ = C4()
            nw4_.ctor__(p1, d_61_i0_, False, p1)
            d_74_v10_ = nw4_
            d_74_v10_ = d_74_v10_
            d_75_v11_: T1
            nw5_ = C6()
            nw5_.ctor__(p1, p2, p1)
            d_75_v11_ = nw5_
            d_76_v12_: C5
            nw6_ = C5()
            nw6_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")), (d_74_v10_).f9, d_74_v10_.f8, p3)
            d_76_v12_ = nw6_
            d_77_v13_: _dafny.Map
            d_77_v13_ = _dafny.Map({d_61_i0_: d_76_v12_})
            d_78_v14_: _dafny.Set
            d_78_v14_ = _dafny.Set({d_77_v13_, d_77_v13_})
            d_79_v15_: _dafny.Map
            d_79_v15_ = _dafny.Map({p2: _dafny.Map({d_76_v12_.f7: d_76_v12_})})
            d_80_v16_: _dafny.Map
            d_80_v16_ = _dafny.Map({d_61_i0_: d_76_v12_.f7})
            d_81_v17_: _dafny.MultiSet
            d_81_v17_ = _dafny.MultiSet([d_80_v16_, d_80_v16_])
            d_82_v18_: D2
            d_82_v18_ = D2_DC8((d_70_v7_)[default__.safeIndex(44, (d_70_v7_).length(0))], (d_67_v4_).f12)
            d_83_v19_: str
            d_83_v19_ = _dafny.CodePoint('i')
            d_84_v20_: C0
            nw7_ = C0()
            nw7_.ctor__()
            d_84_v20_ = nw7_
            d_85_v21_: _dafny.Seq
            d_85_v21_ = _dafny.SeqWithoutIsStrInference([d_84_v20_])
            d_86_v22_: _dafny.MultiSet
            d_86_v22_ = _dafny.MultiSet([d_76_v12_.f7, 870])
            d_87_v23_: _dafny.Set
            d_87_v23_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_83_v19_ for d_88_i3_ in range(default__.abs(341))])), (0) - (p0)})
            d_89_v24_: _dafny.Array
            nw8_ = _dafny.Array(None, 20)
            nw8_[int(0)] = (p0) * (d_61_i0_)
            nw8_[int(1)] = d_61_i0_
            nw8_[int(2)] = d_61_i0_
            nw8_[int(3)] = (24) - (len(_dafny.SeqWithoutIsStrInference([d_75_v11_])))
            nw8_[int(4)] = (d_74_v10_).f9
            nw8_[int(5)] = len((d_78_v14_).intersection(_dafny.Set({((d_79_v15_)[default__.fm21(d_61_i0_, d_81_v17_, globalState)] if (default__.fm21(d_61_i0_, d_81_v17_, globalState)) in (d_79_v15_) else d_77_v13_)})))
            nw8_[int(6)] = p0
            nw8_[int(7)] = d_61_i0_
            nw8_[int(8)] = len(default__.fm30(D1_DC5((d_67_v4_).f13, (d_67_v4_).f13, d_61_i0_, (d_74_v10_).f9), (d_74_v10_).fm5((d_67_v4_).f13, ((d_82_v18_).cf14).set(default__.safeIndex(d_61_i0_, len((d_82_v18_).cf14)), d_83_v19_), globalState), len((d_76_v12_).f6), globalState))
            nw8_[int(9)] = (d_76_v12_.f7) * (d_76_v12_.f7)
            nw8_[int(10)] = (len(d_60_v0_)) * (p0)
            nw8_[int(11)] = len((_dafny.SeqWithoutIsStrInference([d_84_v20_, d_84_v20_, d_84_v20_, d_84_v20_])) + ((d_85_v21_).set(default__.safeIndex((d_86_v22_).cardinality, len(d_85_v21_)), d_84_v20_)))
            nw8_[int(12)] = (d_74_v10_).fm0(globalState)
            nw8_[int(13)] = (d_74_v10_).f9
            nw8_[int(14)] = (d_76_v12_.f7) - (p0)
            nw8_[int(15)] = d_76_v12_.f7
            nw8_[int(16)] = -490
            nw8_[int(17)] = d_61_i0_
            nw8_[int(18)] = default__.safeModuloInt((0) - ((0) - (p0)), len(d_87_v23_))
            nw8_[int(19)] = ((d_74_v10_).fm5(p2, d_62_v1_, globalState)) * (d_61_i0_)
            d_89_v24_ = nw8_
            index2_ = default__.safeIndex(883, (d_89_v24_).length(0))
            (d_89_v24_)[index2_] = d_76_v12_.f7
            index3_ = default__.safeIndex(883, (d_89_v24_).length(0))
            (d_89_v24_)[index3_] = p0
        d_90_v25_: _dafny.Array
        nw9_ = _dafny.Array(int(0), 7)
        d_90_v25_ = nw9_
        d_91_v26_: _dafny.Array
        def lambda3_(d_92_p1_):
            def lambda4_(d_93_i4_):
                return d_92_p1_

            return lambda4_

        init1_ = lambda3_(p1)
        nw10_ = _dafny.Array(None, 2)
        for i0_1_ in range(nw10_.length(0)):
            nw10_[i0_1_] = init1_(i0_1_)
        d_91_v26_ = nw10_
        index4_ = default__.safeIndex(969, (d_91_v26_).length(0))
        (d_91_v26_)[index4_] = p1
        d_94_v27_: _dafny.Seq
        d_94_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkcu"))
        d_95_v28_: C1
        nw11_ = C1()
        nw11_.ctor__(d_94_v27_, p2)
        d_95_v28_ = nw11_
        d_96_v29_: D15
        d_96_v29_ = D15_DC33(d_91_v26_, d_95_v28_)
        d_97_v30_: _dafny.Seq
        d_97_v30_ = _dafny.SeqWithoutIsStrInference([D15_DC33(d_91_v26_, d_95_v28_), d_96_v29_, D15_DC33(d_91_v26_, d_95_v28_)])
        d_98_v31_: _dafny.Map
        d_98_v31_ = _dafny.Map({p0: p3})
        d_99_v32_: _dafny.Set
        d_99_v32_ = _dafny.Set({d_98_v31_})
        d_100_v33_: D8
        d_100_v33_ = D8_DC18(d_98_v31_)
        d_101_v34_: _dafny.Seq
        d_101_v34_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p0: True})])
        d_102_v35_: _dafny.Seq
        d_102_v35_ = _dafny.SeqWithoutIsStrInference([(d_101_v34_)[default__.safeIndex(len((d_95_v28_).f12), len(d_101_v34_))], d_98_v31_])
        index5_ = default__.safeIndex(969, (d_91_v26_).length(0))
        def iife12_():
            coll12_ = _dafny.Set()
            compr_12_: _dafny.Map
            for compr_12_ in (d_102_v35_).Elements:
                d_103_v36_: _dafny.Map = compr_12_
                if (d_103_v36_) in (d_102_v35_):
                    coll12_ = coll12_.union(_dafny.Set([d_103_v36_]))
            return _dafny.Set(coll12_)
        rhs5_ = not((d_97_v30_) <= ((d_97_v30_) + (d_97_v30_)))
        rhs6_ = p0
        rhs7_ = d_90_v25_
        rhs8_ = ((d_99_v32_) | (_dafny.Set({d_98_v31_, _dafny.Map({len((d_100_v33_).cf30): (d_95_v28_).f13})}))).isdisjoint(iife12_()
        )
        lhs3_ = globalState
        lhs4_ = d_91_v26_
        lhs5_ = default__.safeIndex(969, (d_91_v26_).length(0))
        lhs3_.f0 = rhs5_
        r1 = rhs6_
        d_90_v25_ = rhs7_
        lhs4_[lhs5_] = rhs8_
        d_104_v37_: C4
        nw12_ = C4()
        nw12_.ctor__(not(not(False)), p0, (d_95_v28_).f13, (d_91_v26_)[default__.safeIndex(969, (d_91_v26_).length(0))])
        d_104_v37_ = nw12_
        d_105_v38_: _dafny.Map
        d_105_v38_ = _dafny.Map({d_104_v37_: d_104_v37_.f8})
        d_105_v38_ = (d_105_v38_).set(d_104_v37_, ((d_95_v28_).f12) < (d_94_v27_))
        hi1_ = (0) - (p0)
        for d_106_i5_ in range(838, hi1_):
            index6_ = default__.safeIndex(813, (d_90_v25_).length(0))
            (d_90_v25_)[index6_] = (d_104_v37_).f9
            d_107_v39_: D19
            d_107_v39_ = D19_DC45(d_106_i5_, (0) - (d_106_i5_), default__.fm6((d_104_v37_).f9, globalState), p1, d_60_v0_)
            d_108_v40_: C6
            nw13_ = C6()
            nw13_.ctor__(not((d_95_v28_).f13), (d_95_v28_).f13, (d_107_v39_).cf72)
            d_108_v40_ = nw13_
            d_109_v41_: _dafny.Set
            d_109_v41_ = _dafny.Set({d_108_v40_, d_108_v40_, d_108_v40_})
            index7_ = default__.safeIndex(813, (d_90_v25_).length(0))
            rhs9_ = (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lg"))) + ((d_95_v28_).f12))) + ((len(d_109_v41_) if (d_91_v26_)[default__.safeIndex(969, (d_91_v26_).length(0))] else (d_104_v37_).f9))
            rhs10_ = (d_106_i5_) <= (len(d_94_v27_))
            rhs11_ = p0
            rhs12_ = (d_104_v37_).f9
            lhs6_ = d_104_v37_
            lhs7_ = d_90_v25_
            lhs8_ = default__.safeIndex(813, (d_90_v25_).length(0))
            r1 = rhs9_
            lhs6_.f8 = rhs10_
            lhs7_[lhs8_] = rhs11_
            r1 = rhs12_
            index8_ = default__.safeIndex(813, (d_90_v25_).length(0))
            (d_90_v25_)[index8_] = (0) - (p0)
            d_110_v42_: _dafny.Seq
            d_110_v42_ = _dafny.SeqWithoutIsStrInference([p0, (d_104_v37_).f9, d_106_i5_, p0, p0])
            r1 = default__.safeModuloInt(((d_108_v40_).fm1((d_95_v28_).f13, (d_90_v25_)[default__.safeIndex(813, (d_90_v25_).length(0))], True, default__.fm29(d_106_i5_, globalState), globalState)) + (948), (len(d_110_v42_)) * (-952))
            d_111_v43_: str
            d_111_v43_ = _dafny.CodePoint('r')
            d_112_v44_: _dafny.Array
            nw14_ = _dafny.Array(None, 15)
            nw14_[int(0)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_113_i6_ in range(default__.abs(973))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_114_i7_ in range(default__.abs(197))]))
            nw14_[int(1)] = (d_95_v28_).f12
            nw14_[int(2)] = (d_94_v27_).set(default__.safeIndex((d_90_v25_)[default__.safeIndex(813, (d_90_v25_).length(0))], len(d_94_v27_)), d_111_v43_)
            nw14_[int(3)] = (d_95_v28_).f12
            nw14_[int(4)] = (d_94_v27_) + ((d_95_v28_).f12)
            nw14_[int(5)] = d_94_v27_
            nw14_[int(6)] = ((d_95_v28_).f12) + ((d_95_v28_).f12)
            nw14_[int(7)] = d_94_v27_
            nw14_[int(8)] = d_94_v27_
            nw14_[int(9)] = d_94_v27_
            nw14_[int(10)] = (d_95_v28_).f12
            nw14_[int(11)] = d_94_v27_
            nw14_[int(12)] = (d_94_v27_) + ((d_95_v28_).f12)
            nw14_[int(13)] = d_94_v27_
            nw14_[int(14)] = d_94_v27_
            d_112_v44_ = nw14_
            index9_ = default__.safeIndex(298, (d_112_v44_).length(0))
            (d_112_v44_)[index9_] = (d_95_v28_).f12
            index10_ = default__.safeIndex(298, (d_112_v44_).length(0))
            (d_112_v44_)[index10_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugxsphir"))) + (_dafny.SeqWithoutIsStrInference([d_111_v43_ for d_115_i8_ in range(default__.abs(-287))])) if (default__.fm51(globalState)).cf3 else (d_94_v27_) + (d_94_v27_))
        d_116_v45_: _dafny.MultiSet
        d_116_v45_ = _dafny.MultiSet([(d_95_v28_).f13])
        index11_ = default__.safeIndex(697, (d_90_v25_).length(0))
        (d_90_v25_)[index11_] = ((d_116_v45_)[d_104_v37_.f8] if (d_104_v37_.f8) in (d_116_v45_) else p0)
        index12_ = default__.safeIndex(697, (d_90_v25_).length(0))
        rhs13_ = p0
        rhs14_ = ((d_104_v37_).f9) + ((0) - ((d_104_v37_).fm5(True, d_94_v27_, globalState)))
        rhs15_ = (d_91_v26_ if p1 else d_91_v26_)
        lhs9_ = d_90_v25_
        lhs10_ = default__.safeIndex(697, (d_90_v25_).length(0))
        r1 = rhs13_
        lhs9_[lhs10_] = rhs14_
        d_91_v26_ = rhs15_
        d_117_i9_: int
        d_117_i9_ = 0
        with _dafny.label("0"):
            while (d_91_v26_)[default__.safeIndex(969, (d_91_v26_).length(0))]:
                with _dafny.c_label("0"):
                    if (d_117_i9_) >= (100):
                        raise _dafny.Break("0")
                    d_117_i9_ = (d_117_i9_) + (1)
                    d_118_v46_: _dafny.Array
                    def lambda5_(d_119_v0_):
                        def lambda6_(d_120_i10_):
                            return D4_DC11(d_119_v0_)

                        return lambda6_

                    init2_ = lambda5_(d_60_v0_)
                    nw15_ = _dafny.Array(None, 1)
                    for i0_2_ in range(nw15_.length(0)):
                        nw15_[i0_2_] = init2_(i0_2_)
                    d_118_v46_ = nw15_
                    d_121_v47_: D4
                    d_121_v47_ = D4_DC11(d_60_v0_)
                    index13_ = default__.safeIndex(199, (d_118_v46_).length(0))
                    (d_118_v46_)[index13_] = d_121_v47_
                    pat_let_tv0_ = d_104_v37_
                    index14_ = default__.safeIndex(697, (d_90_v25_).length(0))
                    index15_ = default__.safeIndex(199, (d_118_v46_).length(0))
                    def iife13_(_pat_let0_0):
                        def iife14_(d_122_dt__update__tmp_h0_):
                            def iife15_(_pat_let1_0):
                                def iife16_(d_123_dt__update_hcf18_h0_):
                                    return D4_DC11(d_123_dt__update_hcf18_h0_)
                                return iife16_(_pat_let1_0)
                            return iife15_(_dafny.SeqWithoutIsStrInference([pat_let_tv0_.f8]))
                        return iife14_(_pat_let0_0)
                    rhs16_ = (d_90_v25_)[default__.safeIndex(697, (d_90_v25_).length(0))]
                    rhs17_ = iife13_(d_121_v47_)
                    lhs11_ = d_90_v25_
                    lhs12_ = default__.safeIndex(697, (d_90_v25_).length(0))
                    lhs13_ = d_118_v46_
                    lhs14_ = default__.safeIndex(199, (d_118_v46_).length(0))
                    lhs11_[lhs12_] = rhs16_
                    lhs13_[lhs14_] = rhs17_
                    index16_ = default__.safeIndex(969, (d_91_v26_).length(0))
                    (d_91_v26_)[index16_] = p1
                    d_124_v48_: _dafny.Seq
                    d_124_v48_ = _dafny.SeqWithoutIsStrInference([((d_116_v45_)[p3] if (p3) in (d_116_v45_) else (d_104_v37_).f9), 593, (d_104_v37_).f9, p0])
                    r1 = ((len(d_124_v48_)) + ((d_104_v37_).f9)) * ((0) - ((d_104_v37_).f9))
                    d_125_v49_: C1
                    nw16_ = C1()
                    nw16_.ctor__(((d_95_v28_).f12) + (d_94_v27_), (d_95_v28_).f13)
                    d_125_v49_ = nw16_
                    pass
            pass
        d_126_v50_: _dafny.Map
        d_126_v50_ = _dafny.Map({p0: (d_90_v25_)[default__.safeIndex(697, (d_90_v25_).length(0))]})
        d_127_v51_: _dafny.Seq
        d_127_v51_ = _dafny.SeqWithoutIsStrInference([d_126_v50_])
        d_128_v52_: _dafny.MultiSet
        d_128_v52_ = _dafny.MultiSet([d_94_v27_])
        r0 = D11_DC24((d_127_v51_)[default__.safeIndex((d_128_v52_).cardinality, len(d_127_v51_))])
        r1 = (d_90_v25_)[default__.safeIndex(697, (d_90_v25_).length(0))]
        d_129_v53_: T1
        nw17_ = C3()
        nw17_.ctor__(p2, not(not(default__.fm21((d_90_v25_)[default__.safeIndex(697, (d_90_v25_).length(0))], _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_126_v50_, d_126_v50_, d_126_v50_])), globalState))))
        d_129_v53_ = nw17_
        d_130_v54_: _dafny.Seq
        d_130_v54_ = _dafny.SeqWithoutIsStrInference([d_129_v53_])
        r2 = d_130_v54_
        nw18_ = C2()
        nw18_.ctor__((d_91_v26_)[default__.safeIndex(969, (d_91_v26_).length(0))], d_104_v37_.f8, ((0) - ((d_104_v37_).f9)) <= ((0) - (p0)), p1)
        r3 = nw18_
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_131_globalState_: GlobalState
        nw19_ = GlobalState()
        nw19_.ctor__(True, True, True)
        d_131_globalState_ = nw19_
        d_132_v0_: _dafny.Seq
        d_132_v0_ = _dafny.SeqWithoutIsStrInference([True])
        d_133_v1_: bool
        d_133_v1_ = False
        d_134_i0_: int
        d_134_i0_ = 0
        with _dafny.label("1"):
            while ((d_132_v0_) <= (_dafny.SeqWithoutIsStrInference([d_133_v1_, d_133_v1_]))) and (d_133_v1_):
                with _dafny.c_label("1"):
                    if (d_134_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_134_i0_ = (d_134_i0_) + (1)
                    d_135_v2_: C6
                    nw20_ = C6()
                    nw20_.ctor__(d_133_v1_, d_133_v1_, d_133_v1_)
                    d_135_v2_ = nw20_
                    d_136_v3_: _dafny.MultiSet
                    d_136_v3_ = _dafny.MultiSet([d_135_v2_.f5, True])
                    d_137_v4_: int
                    d_137_v4_ = 194
                    d_138_v5_: C4
                    nw21_ = C4()
                    nw21_.ctor__(default__.fm6(d_137_v4_, d_131_globalState_), d_137_v4_, d_133_v1_, d_135_v2_.f5)
                    d_138_v5_ = nw21_
                    d_139_v6_: _dafny.Map
                    d_139_v6_ = _dafny.Map({d_138_v5_: d_138_v5_.f8})
                    d_140_v7_: _dafny.Map
                    d_140_v7_ = _dafny.Map({not(((d_139_v6_)[d_138_v5_] if (d_138_v5_) in (d_139_v6_) else d_138_v5_.f8)): d_133_v1_})
                    d_141_v8_: _dafny.Map
                    d_141_v8_ = _dafny.Map({((d_140_v7_)[not(d_138_v5_.f8)] if (not(d_138_v5_.f8)) in (d_140_v7_) else d_138_v5_.f8): d_136_v3_})
                    d_142_v9_: _dafny.Seq
                    d_142_v9_ = _dafny.SeqWithoutIsStrInference([default__.fm20(d_135_v2_.f5, d_131_globalState_), d_136_v3_, d_136_v3_, d_136_v3_, d_136_v3_])
                    d_143_v10_: _dafny.Array
                    nw22_ = _dafny.Array(None, 14)
                    nw22_[int(0)] = (d_136_v3_).intersection(d_136_v3_)
                    nw22_[int(1)] = d_136_v3_
                    nw22_[int(2)] = ((d_141_v8_)[d_138_v5_.f8] if (d_138_v5_.f8) in (d_141_v8_) else default__.fm20(d_135_v2_.f5, d_131_globalState_))
                    nw22_[int(3)] = d_136_v3_
                    nw22_[int(4)] = (d_136_v3_) - (d_136_v3_)
                    nw22_[int(5)] = default__.fm20(d_135_v2_.f5, d_131_globalState_)
                    nw22_[int(6)] = d_136_v3_
                    nw22_[int(7)] = d_136_v3_
                    nw22_[int(8)] = d_136_v3_
                    nw22_[int(9)] = d_136_v3_
                    nw22_[int(10)] = (d_142_v9_)[default__.safeIndex(-596, len(d_142_v9_))]
                    nw22_[int(11)] = _dafny.MultiSet([d_135_v2_.f5])
                    nw22_[int(12)] = (d_136_v3_) | (d_136_v3_)
                    nw22_[int(13)] = d_136_v3_
                    d_143_v10_ = nw22_
                    index17_ = default__.safeIndex(329, (d_143_v10_).length(0))
                    (d_143_v10_)[index17_] = d_136_v3_
                    index18_ = default__.safeIndex(329, (d_143_v10_).length(0))
                    (d_143_v10_)[index18_] = default__.fm20(d_135_v2_.f5, d_131_globalState_)
                    d_144_v11_: C3
                    nw23_ = C3()
                    nw23_.ctor__(not(d_138_v5_.f8), d_138_v5_.f8)
                    d_144_v11_ = nw23_
                    nw24_ = C3()
                    nw24_.ctor__(d_138_v5_.f8, d_138_v5_.f8)
                    d_144_v11_ = nw24_
                    d_137_v4_ = (786) - ((0) - (d_137_v4_))
                    pass
            pass
        if ((d_133_v1_) == (d_133_v1_) if d_133_v1_ else d_133_v1_):
            d_145_v12_: int
            d_145_v12_ = 958
            d_146_v13_: _dafny.MultiSet
            d_146_v13_ = _dafny.MultiSet([d_145_v12_])
            d_147_v14_: D23
            d_147_v14_ = D23_DC61((d_146_v13_).cardinality)
            pat_let_tv1_ = d_145_v12_
            pat_let_tv2_ = d_133_v1_
            pat_let_tv3_ = d_145_v12_
            def iife17_(_pat_let2_0):
                def iife18_(d_148_dt__update__tmp_h0_):
                    def iife19_(_pat_let3_0):
                        def iife20_(d_149_dt__update_hcf98_h0_):
                            return D23_DC61(d_149_dt__update_hcf98_h0_)
                        return iife20_(_pat_let3_0)
                    return iife19_((pat_let_tv1_ if pat_let_tv2_ else pat_let_tv3_))
                return iife18_(_pat_let2_0)
            source4_ = iife17_(d_147_v14_)
            if source4_.is_DC59:
                d_150___mcc_h0_ = source4_.cf93
                d_151___mcc_h1_ = source4_.cf94
                d_152___mcc_h2_ = source4_.cf95
                d_153_cf95_ = d_152___mcc_h2_
                d_154_cf94_ = d_151___mcc_h1_
                d_155_cf93_ = d_150___mcc_h0_
                d_133_v1_ = (d_145_v12_) != ((0) - ((882) * ((0) - (d_153_cf95_))))
                (d_131_globalState_).f0 = (d_133_v1_ if d_133_v1_ else d_133_v1_)
                d_154_cf94_ = d_154_cf94_
                d_155_cf93_ = (d_155_cf93_) + (d_155_cf93_)
            elif source4_.is_DC60:
                d_156___mcc_h3_ = source4_.cf96
                d_157___mcc_h4_ = source4_.cf97
                d_158_cf97_ = d_157___mcc_h4_
                d_159_cf96_ = d_156___mcc_h3_
                d_160_v15_: _dafny.Seq
                d_160_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
                d_161_v17_: str
                d_161_v17_ = _dafny.CodePoint('u')
                d_162_v18_: _dafny.Array
                nw25_ = _dafny.Array(None, 13)
                nw25_[int(0)] = (0) - (d_145_v12_)
                nw25_[int(1)] = d_145_v12_
                nw25_[int(2)] = 856
                def iife21_():
                    coll13_ = _dafny.Map()
                    compr_13_: int
                    for compr_13_ in (_dafny.SeqWithoutIsStrInference([342])).Elements:
                        d_163_v16_: int = compr_13_
                        if (d_163_v16_) in (_dafny.SeqWithoutIsStrInference([342])):
                            coll13_[(d_163_v16_) * (d_145_v12_)] = d_145_v12_
                    return _dafny.Map(coll13_)
                nw25_[int(3)] = (0) - (len(iife21_()
                ))
                nw25_[int(4)] = d_145_v12_
                nw25_[int(5)] = default__.fm31(d_133_v1_, 924, d_159_cf96_, _dafny.SeqWithoutIsStrInference([d_133_v1_]), d_131_globalState_)
                nw25_[int(6)] = d_145_v12_
                nw25_[int(7)] = d_145_v12_
                nw25_[int(8)] = d_145_v12_
                nw25_[int(9)] = len(_dafny.Map({d_161_v17_: d_145_v12_}))
                nw25_[int(10)] = d_145_v12_
                nw25_[int(11)] = d_145_v12_
                nw25_[int(12)] = d_145_v12_
                d_162_v18_ = nw25_
                d_164_v19_: D17
                d_164_v19_ = D17_DC39(d_162_v18_, d_133_v1_, d_133_v1_, d_145_v12_, d_145_v12_)
                d_165_v20_: D10
                d_165_v20_ = D10_DC22(d_162_v18_)
                d_166_v21_: _dafny.Seq
                d_166_v21_ = _dafny.SeqWithoutIsStrInference([d_162_v18_, d_162_v18_])
                d_167_v22_: _dafny.Array
                nw26_ = _dafny.Array(None, 27)
                nw26_[int(0)] = (d_164_v19_).cf57
                nw26_[int(1)] = d_162_v18_
                nw26_[int(2)] = d_162_v18_
                nw26_[int(3)] = d_162_v18_
                nw26_[int(4)] = d_162_v18_
                nw26_[int(5)] = d_162_v18_
                nw26_[int(6)] = (d_165_v20_).cf36
                nw26_[int(7)] = d_162_v18_
                nw26_[int(8)] = d_162_v18_
                nw26_[int(9)] = d_162_v18_
                nw26_[int(10)] = d_162_v18_
                nw26_[int(11)] = d_162_v18_
                nw26_[int(12)] = d_162_v18_
                nw26_[int(13)] = d_162_v18_
                nw26_[int(14)] = d_162_v18_
                nw26_[int(15)] = d_162_v18_
                nw26_[int(16)] = d_162_v18_
                nw26_[int(17)] = d_162_v18_
                nw26_[int(18)] = d_162_v18_
                nw26_[int(19)] = d_162_v18_
                nw26_[int(20)] = d_162_v18_
                nw26_[int(21)] = d_162_v18_
                nw26_[int(22)] = d_162_v18_
                nw26_[int(23)] = d_162_v18_
                nw26_[int(24)] = (d_166_v21_)[default__.safeIndex(default__.fm31(d_159_cf96_, d_145_v12_, d_158_cf97_, d_132_v0_, d_131_globalState_), len(d_166_v21_))]
                nw26_[int(25)] = d_162_v18_
                nw26_[int(26)] = d_162_v18_
                d_167_v22_ = nw26_
                d_168_v23_: _dafny.Array
                nw27_ = _dafny.Array(None, 17)
                nw27_[int(0)] = d_167_v22_
                nw27_[int(1)] = d_167_v22_
                nw27_[int(2)] = d_167_v22_
                nw27_[int(3)] = d_167_v22_
                nw27_[int(4)] = d_167_v22_
                nw27_[int(5)] = d_167_v22_
                nw27_[int(6)] = d_167_v22_
                nw27_[int(7)] = d_167_v22_
                nw27_[int(8)] = d_167_v22_
                nw27_[int(9)] = d_167_v22_
                nw27_[int(10)] = d_167_v22_
                nw27_[int(11)] = d_167_v22_
                nw27_[int(12)] = (d_167_v22_ if d_159_cf96_ else d_167_v22_)
                nw27_[int(13)] = d_167_v22_
                nw27_[int(14)] = d_167_v22_
                nw27_[int(15)] = d_167_v22_
                nw27_[int(16)] = d_167_v22_
                d_168_v23_ = nw27_
                index19_ = default__.safeIndex(704, (d_168_v23_).length(0))
                (d_168_v23_)[index19_] = d_167_v22_
                d_169_v24_: _dafny.MultiSet
                d_169_v24_ = _dafny.MultiSet([_dafny.Map({d_145_v12_: d_145_v12_})])
                d_170_v25_: _dafny.MultiSet
                d_170_v25_ = _dafny.MultiSet([d_159_cf96_, default__.fm21(d_145_v12_, d_169_v24_, d_131_globalState_)])
                index20_ = default__.safeIndex(704, (d_168_v23_).length(0))
                rhs18_ = d_160_v15_
                rhs19_ = d_167_v22_
                rhs20_ = (_dafny.MultiSet([d_158_cf97_])).ispropersubset((_dafny.MultiSet([d_158_cf97_])).intersection(d_170_v25_))
                lhs15_ = d_168_v23_
                lhs16_ = default__.safeIndex(704, (d_168_v23_).length(0))
                lhs17_ = d_131_globalState_
                d_160_v15_ = rhs18_
                lhs15_[lhs16_] = rhs19_
                lhs17_.f0 = rhs20_
                index21_ = default__.safeIndex(443, (d_162_v18_).length(0))
                (d_162_v18_)[index21_] = d_145_v12_
                d_171_v26_: _dafny.Map
                d_171_v26_ = _dafny.Map({d_145_v12_: default__.safeModuloInt((0) - (d_145_v12_), d_145_v12_)})
                index22_ = default__.safeIndex(443, (d_162_v18_).length(0))
                (d_162_v18_)[index22_] = ((d_171_v26_)[298] if (298) in (d_171_v26_) else d_145_v12_)
                d_172_v27_: T1
                nw28_ = C5()
                nw28_.ctor__(d_160_v15_, len(_dafny.Set({True})), d_158_cf97_, True)
                d_172_v27_ = nw28_
                d_173_v28_: _dafny.Map
                d_173_v28_ = _dafny.Map({d_158_cf97_: d_172_v27_})
                d_174_v29_: _dafny.Seq
                d_174_v29_ = _dafny.SeqWithoutIsStrInference([d_145_v12_, (d_162_v18_)[default__.safeIndex(443, (d_162_v18_).length(0))]])
                d_175_v30_: D23
                d_175_v30_ = D23_DC59((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xa"))).set(default__.safeIndex(len(d_174_v29_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xa")))), (d_160_v15_)[default__.safeIndex((D8_DC19(d_145_v12_, (d_162_v18_)[default__.safeIndex(443, (d_162_v18_).length(0))])).cf32, len(d_160_v15_))]), (d_170_v25_).cardinality, d_145_v12_)
                d_176_v31_: D1
                d_176_v31_ = D1_DC5(d_158_cf97_, default__.fm4(d_133_v1_, d_131_globalState_), len(_dafny.Set({default__.fm49(len(d_173_v28_), d_131_globalState_), d_175_v30_})), ((d_146_v13_)[430] if (430) in (d_146_v13_) else len(d_171_v26_)))
                d_160_v15_ = (((d_160_v15_ if d_133_v1_ else d_160_v15_)) + (default__.fm30(d_176_v31_, len(d_160_v15_), len(_dafny.Map({d_145_v12_: (0) - (d_145_v12_)})), d_131_globalState_))).set(default__.safeIndex(d_145_v12_, len(((d_160_v15_ if d_133_v1_ else d_160_v15_)) + (default__.fm30(d_176_v31_, len(d_160_v15_), len(_dafny.Map({d_145_v12_: (0) - (d_145_v12_)})), d_131_globalState_)))), d_161_v17_)
                d_160_v15_ = ((_dafny.SeqWithoutIsStrInference([d_161_v17_ for d_177_i1_ in range(default__.abs(452))])).set(default__.safeIndex(d_145_v12_, len(_dafny.SeqWithoutIsStrInference([d_161_v17_ for d_178_i1_ in range(default__.abs(452))]))), d_161_v17_)).set(default__.safeIndex((d_162_v18_)[default__.safeIndex(443, (d_162_v18_).length(0))], len((_dafny.SeqWithoutIsStrInference([d_161_v17_ for d_179_i1_ in range(default__.abs(452))])).set(default__.safeIndex(d_145_v12_, len(_dafny.SeqWithoutIsStrInference([d_161_v17_ for d_180_i1_ in range(default__.abs(452))]))), d_161_v17_))), (d_160_v15_)[default__.safeIndex(d_145_v12_, len(d_160_v15_))])
            elif source4_.is_DC61:
                d_181___mcc_h5_ = source4_.cf98
                d_182_cf98_ = d_181___mcc_h5_
                d_183_v32_: _dafny.Array
                nw29_ = _dafny.Array(None, 27)
                nw29_[int(0)] = d_133_v1_
                nw29_[int(1)] = d_133_v1_
                nw29_[int(2)] = False
                nw29_[int(3)] = d_133_v1_
                nw29_[int(4)] = d_133_v1_
                nw29_[int(5)] = d_133_v1_
                nw29_[int(6)] = d_133_v1_
                nw29_[int(7)] = d_133_v1_
                nw29_[int(8)] = not(d_133_v1_)
                nw29_[int(9)] = d_133_v1_
                nw29_[int(10)] = d_133_v1_
                nw29_[int(11)] = d_133_v1_
                nw29_[int(12)] = False
                nw29_[int(13)] = d_133_v1_
                nw29_[int(14)] = d_133_v1_
                nw29_[int(15)] = d_133_v1_
                nw29_[int(16)] = d_133_v1_
                nw29_[int(17)] = d_133_v1_
                nw29_[int(18)] = d_133_v1_
                nw29_[int(19)] = d_133_v1_
                nw29_[int(20)] = True
                nw29_[int(21)] = d_133_v1_
                nw29_[int(22)] = d_133_v1_
                nw29_[int(23)] = d_133_v1_
                nw29_[int(24)] = d_133_v1_
                nw29_[int(25)] = d_133_v1_
                nw29_[int(26)] = d_133_v1_
                d_183_v32_ = nw29_
                d_184_v33_: _dafny.Map
                d_184_v33_ = _dafny.Map({d_183_v32_: d_145_v12_})
                d_184_v33_ = (d_184_v33_) | (d_184_v33_)
                d_145_v12_ = d_145_v12_
                d_185_v34_: _dafny.Seq
                d_185_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpqwyym"))
                d_186_v35_: str
                d_186_v35_ = _dafny.CodePoint('a')
                d_185_v34_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bedn"))).set(default__.safeIndex((d_182_cf98_) - (d_182_cf98_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bedn")))), d_186_v35_)
                d_187_v36_: D2
                d_187_v36_ = D2_DC8(d_183_v32_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvrbrm")))
                d_188_v37_: D1
                d_188_v37_ = D1_DC5(d_133_v1_, d_133_v1_, d_145_v12_, d_182_cf98_)
                d_189_v38_: T0
                nw30_ = C5()
                nw30_.ctor__(d_185_v34_, d_182_cf98_, not(False), d_133_v1_)
                d_189_v38_ = nw30_
                d_190_v39_: _dafny.Map
                d_190_v39_ = _dafny.Map({d_133_v1_: d_189_v38_})
                d_191_v40_: _dafny.Array
                nw31_ = _dafny.Array(None, 18)
                nw31_[int(0)] = d_185_v34_
                nw31_[int(1)] = (d_185_v34_ if d_133_v1_ else d_185_v34_)
                nw31_[int(2)] = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), d_186_v35_]) if d_133_v1_ else d_185_v34_)).set(default__.safeIndex(d_182_cf98_, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), d_186_v35_]) if d_133_v1_ else d_185_v34_))), d_186_v35_)
                nw31_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_186_v35_])) + (d_185_v34_)
                nw31_[int(4)] = (d_185_v34_) + ((d_187_v36_).cf14)
                nw31_[int(5)] = _dafny.SeqWithoutIsStrInference([d_186_v35_ for d_192_i2_ in range(default__.abs(531))])
                nw31_[int(6)] = d_185_v34_
                nw31_[int(7)] = (_dafny.SeqWithoutIsStrInference([d_186_v35_ for d_193_i3_ in range(default__.abs(581))])).set(default__.safeIndex(len((d_132_v0_).set(default__.safeIndex(d_145_v12_, len(d_132_v0_)), not(d_133_v1_))), len(_dafny.SeqWithoutIsStrInference([d_186_v35_ for d_194_i3_ in range(default__.abs(581))]))), d_186_v35_)
                nw31_[int(8)] = d_185_v34_
                nw31_[int(9)] = d_185_v34_
                nw31_[int(10)] = d_185_v34_
                nw31_[int(11)] = default__.fm30(d_188_v37_, 698, (0) - (d_145_v12_), d_131_globalState_)
                nw31_[int(12)] = d_185_v34_
                nw31_[int(13)] = (d_185_v34_ if default__.fm4(True, d_131_globalState_) else d_185_v34_)
                nw31_[int(14)] = (d_185_v34_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), d_186_v35_, (d_185_v34_)[default__.safeIndex(len((d_190_v39_).set((d_189_v38_).f4, d_189_v38_)), len(d_185_v34_))]]))
                nw31_[int(15)] = d_185_v34_
                nw31_[int(16)] = _dafny.SeqWithoutIsStrInference([d_186_v35_])
                nw31_[int(17)] = d_185_v34_
                d_191_v40_ = nw31_
                index23_ = default__.safeIndex(113, (d_191_v40_).length(0))
                (d_191_v40_)[index23_] = d_185_v34_
                d_195_v41_: T1
                nw32_ = C3()
                nw32_.ctor__(d_133_v1_, (d_189_v38_).f4)
                d_195_v41_ = nw32_
                index24_ = default__.safeIndex(113, (d_191_v40_).length(0))
                rhs21_ = ((d_185_v34_) + (_dafny.SeqWithoutIsStrInference([d_186_v35_ for d_196_i4_ in range(default__.abs(-641))]))) + ((_dafny.SeqWithoutIsStrInference([d_186_v35_ for d_197_i5_ in range(default__.abs(223))])) + (d_185_v34_))
                rhs22_ = d_185_v34_
                rhs23_ = d_195_v41_
                rhs24_ = d_182_cf98_
                rhs25_ = (d_185_v34_) + (d_185_v34_)
                lhs18_ = d_191_v40_
                lhs19_ = default__.safeIndex(113, (d_191_v40_).length(0))
                lhs18_[lhs19_] = rhs21_
                d_185_v34_ = rhs22_
                d_195_v41_ = rhs23_
                d_182_cf98_ = rhs24_
                d_185_v34_ = rhs25_
            elif True:
                d_198___mcc_h6_ = source4_.cf92
                d_199_cf92_ = d_198___mcc_h6_
                d_200_v43_: _dafny.Map
                def iife22_():
                    coll14_ = _dafny.Set()
                    compr_14_: _dafny.Seq
                    for compr_14_ in (default__.fm50(d_133_v1_, d_131_globalState_)).keys.Elements:
                        d_201_v42_: _dafny.Seq = compr_14_
                        if (d_201_v42_) in (default__.fm50(d_133_v1_, d_131_globalState_)):
                            coll14_ = coll14_.union(_dafny.Set([d_201_v42_]))
                    return _dafny.Set(coll14_)
                d_200_v43_ = _dafny.Map({False: len(iife22_()
                )})
                d_202_v44_: _dafny.Seq
                d_202_v44_ = _dafny.SeqWithoutIsStrInference([default__.fm28(d_133_v1_, d_133_v1_, False, (0) - (d_145_v12_), d_131_globalState_), (d_200_v43_) | (d_200_v43_), d_200_v43_])
                d_203_v45_: _dafny.Set
                d_203_v45_ = _dafny.Set({len(d_132_v0_)})
                d_204_v46_: _dafny.Map
                d_204_v46_ = _dafny.Map({d_145_v12_: d_133_v1_})
                d_202_v44_ = (_dafny.SeqWithoutIsStrInference([d_200_v43_, d_200_v43_, d_200_v43_, d_200_v43_]) if (_dafny.Set({len(d_204_v46_)})).issubset(d_203_v45_) else d_202_v44_)
                d_205_v47_: _dafny.Array
                def lambda7_(d_206_v12_):
                    def lambda8_(d_207_i6_):
                        return default__.safeModuloInt(d_207_i6_, d_206_v12_)

                    return lambda8_

                init3_ = lambda7_(d_145_v12_)
                nw33_ = _dafny.Array(None, 27)
                for i0_3_ in range(nw33_.length(0)):
                    nw33_[i0_3_] = init3_(i0_3_)
                d_205_v47_ = nw33_
                d_208_v48_: T1
                nw34_ = C5()
                nw34_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_209_i7_ in range(default__.abs(873))]), d_145_v12_, d_133_v1_, not(d_133_v1_))
                d_208_v48_ = nw34_
                d_210_v49_: _dafny.Map
                d_210_v49_ = _dafny.Map({d_208_v48_: d_205_v47_})
                d_211_v50_: D21
                d_211_v50_ = D21_DC52(d_145_v12_, d_205_v47_)
                d_212_v51_: _dafny.Seq
                d_212_v51_ = _dafny.SeqWithoutIsStrInference([d_205_v47_, d_205_v47_, (d_211_v50_).cf85])
                d_213_v52_: _dafny.Array
                nw35_ = _dafny.Array(None, 14)
                nw35_[int(0)] = d_205_v47_
                nw35_[int(1)] = d_205_v47_
                nw35_[int(2)] = d_205_v47_
                nw35_[int(3)] = d_205_v47_
                nw35_[int(4)] = d_205_v47_
                nw35_[int(5)] = d_205_v47_
                nw35_[int(6)] = d_205_v47_
                nw35_[int(7)] = (d_205_v47_ if d_133_v1_ else d_205_v47_)
                nw35_[int(8)] = d_205_v47_
                nw35_[int(9)] = d_205_v47_
                nw35_[int(10)] = d_205_v47_
                nw35_[int(11)] = ((d_210_v49_)[d_208_v48_] if (d_208_v48_) in (d_210_v49_) else d_205_v47_)
                nw35_[int(12)] = d_205_v47_
                nw35_[int(13)] = (d_212_v51_)[default__.safeIndex(d_145_v12_, len(d_212_v51_))]
                d_213_v52_ = nw35_
                d_213_v52_ = d_213_v52_
                d_214_v53_: _dafny.Seq
                d_214_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvdye"))
                d_215_v54_: C0
                nw36_ = C0()
                nw36_.ctor__()
                d_215_v54_ = nw36_
                d_216_v55_: _dafny.Seq
                d_216_v55_ = _dafny.SeqWithoutIsStrInference([d_215_v54_])
                d_217_v56_: _dafny.Array
                nw37_ = _dafny.Array(None, 8)
                nw37_[int(0)] = d_215_v54_
                nw37_[int(1)] = d_215_v54_
                nw37_[int(2)] = d_215_v54_
                nw37_[int(3)] = d_215_v54_
                nw37_[int(4)] = d_215_v54_
                nw37_[int(5)] = d_215_v54_
                nw37_[int(6)] = d_215_v54_
                nw37_[int(7)] = (d_216_v55_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_145_v12_, d_145_v12_])), len(d_216_v55_))]
                d_217_v56_ = nw37_
                d_218_v57_: _dafny.Map
                d_218_v57_ = _dafny.Map({True: d_217_v56_})
                d_219_v58_: D19
                d_219_v58_ = D19_DC44(d_218_v57_)
                index25_ = default__.safeIndex(764, (d_205_v47_).length(0))
                (d_205_v47_)[index25_] = (d_145_v12_) + (len(_dafny.SeqWithoutIsStrInference([d_208_v48_])))
                d_220_v59_: _dafny.Seq
                d_220_v59_ = _dafny.SeqWithoutIsStrInference([d_145_v12_])
                index26_ = default__.safeIndex(764, (d_205_v47_).length(0))
                rhs26_ = d_214_v53_
                rhs27_ = d_133_v1_
                rhs28_ = d_219_v58_
                rhs29_ = (0) - (d_145_v12_)
                rhs30_ = (d_215_v54_).fm10((d_220_v59_)[default__.safeIndex((0) - (d_145_v12_), len(d_220_v59_))], d_133_v1_, d_146_v13_, d_131_globalState_)
                lhs20_ = d_131_globalState_
                lhs21_ = d_205_v47_
                lhs22_ = default__.safeIndex(764, (d_205_v47_).length(0))
                d_214_v53_ = rhs26_
                lhs20_.f0 = rhs27_
                d_219_v58_ = rhs28_
                d_145_v12_ = rhs29_
                lhs21_[lhs22_] = rhs30_
                d_221_v61_: D1
                def iife23_():
                    coll15_ = _dafny.Map()
                    compr_15_: int
                    for compr_15_ in _dafny.IntegerRange(25, -331):
                        d_222_v60_: int = compr_15_
                        if ((25) <= (d_222_v60_)) and ((d_222_v60_) < (-331)):
                            coll15_[default__.safeDivisionInt(d_222_v60_, d_145_v12_)] = 869
                    return _dafny.Map(coll15_)
                d_221_v61_ = D1_DC5(default__.fm4(d_133_v1_, d_131_globalState_), True, len(iife23_()
), (d_215_v54_).fm10(len(_dafny.Map({(d_205_v47_)[default__.safeIndex(764, (d_205_v47_).length(0))]: -742})), d_133_v1_, d_146_v13_, d_131_globalState_))
                d_223_v62_: D19
                d_223_v62_ = D19_DC46((0) - ((d_205_v47_)[default__.safeIndex(764, (d_205_v47_).length(0))]), d_133_v1_, d_133_v1_, d_221_v61_, (0) - ((d_205_v47_)[default__.safeIndex(764, (d_205_v47_).length(0))]))
                index27_ = default__.safeIndex(764, (d_205_v47_).length(0))
                (d_205_v47_)[index27_] = (d_223_v62_).cf78
            d_224_v63_: str
            d_224_v63_ = _dafny.CodePoint('y')
            d_225_v64_: C4
            nw38_ = C4()
            nw38_.ctor__(d_133_v1_, d_145_v12_, d_133_v1_, d_133_v1_)
            d_225_v64_ = nw38_
            d_226_v65_: _dafny.Map
            d_226_v65_ = _dafny.Map({d_225_v64_: not(d_225_v64_.f8)})
            d_227_v66_: _dafny.Array
            nw39_ = _dafny.Array(None, 25)
            nw39_[int(0)] = d_133_v1_
            nw39_[int(1)] = False
            nw39_[int(2)] = d_133_v1_
            nw39_[int(3)] = d_133_v1_
            nw39_[int(4)] = d_133_v1_
            nw39_[int(5)] = d_133_v1_
            nw39_[int(6)] = d_133_v1_
            nw39_[int(7)] = default__.fm18(d_145_v12_, d_224_v63_, d_145_v12_, d_131_globalState_)
            nw39_[int(8)] = d_133_v1_
            nw39_[int(9)] = d_133_v1_
            nw39_[int(10)] = d_133_v1_
            nw39_[int(11)] = d_133_v1_
            nw39_[int(12)] = not(d_133_v1_)
            nw39_[int(13)] = d_133_v1_
            nw39_[int(14)] = d_133_v1_
            nw39_[int(15)] = d_133_v1_
            nw39_[int(16)] = d_133_v1_
            nw39_[int(17)] = ((d_226_v65_)[d_225_v64_] if (d_225_v64_) in (d_226_v65_) else d_133_v1_)
            nw39_[int(18)] = d_133_v1_
            nw39_[int(19)] = d_133_v1_
            nw39_[int(20)] = d_133_v1_
            nw39_[int(21)] = d_225_v64_.f8
            nw39_[int(22)] = d_225_v64_.f8
            nw39_[int(23)] = d_133_v1_
            nw39_[int(24)] = d_133_v1_
            d_227_v66_ = nw39_
            d_228_v67_: D2
            d_228_v67_ = D2_DC8(d_227_v66_, _dafny.SeqWithoutIsStrInference([d_224_v63_ for d_229_i8_ in range(default__.abs(928))]))
            pat_let_tv4_ = d_227_v66_
            def iife24_(_pat_let4_0):
                def iife25_(d_230_dt__update__tmp_h1_):
                    def iife26_(_pat_let5_0):
                        def iife27_(d_231_dt__update_hcf13_h0_):
                            return D2_DC8(d_231_dt__update_hcf13_h0_, (d_230_dt__update__tmp_h1_).cf14)
                        return iife27_(_pat_let5_0)
                    return iife26_(pat_let_tv4_)
                return iife25_(_pat_let4_0)
            source5_ = iife24_(d_228_v67_)
            if source5_.is_DC8:
                d_232___mcc_h7_ = source5_.cf13
                d_233___mcc_h8_ = source5_.cf14
                d_234_cf14_ = d_233___mcc_h8_
                d_235_cf13_ = d_232___mcc_h7_
                d_236_v68_: bool
                d_237_v69_: int
                d_238_v70_: _dafny.Array
                out0_: bool
                out1_: int
                out2_: _dafny.Array
                out0_, out1_, out2_ = (d_225_v64_).m6(d_225_v64_.f8, d_131_globalState_)
                d_236_v68_ = out0_
                d_237_v69_ = out1_
                d_238_v70_ = out2_
                d_145_v12_ = (0) - (((d_225_v64_).f9) * (d_237_v69_))
                d_225_v64_ = d_225_v64_
                d_239_v71_: T0
                nw40_ = C4()
                nw40_.ctor__(d_133_v1_, (d_225_v64_).f9, d_225_v64_.f8, d_225_v64_.f8)
                d_239_v71_ = nw40_
                d_240_v72_: _dafny.Set
                d_240_v72_ = _dafny.Set({d_239_v71_})
                d_241_v74_: _dafny.Map
                d_241_v74_ = _dafny.Map({d_237_v69_: d_133_v1_})
                d_242_v75_: _dafny.Map
                d_242_v75_ = _dafny.Map({(d_225_v64_).f9: len(d_241_v74_)})
                d_243_v76_: _dafny.Map
                d_243_v76_ = _dafny.Map({d_236_v68_: d_145_v12_})
                d_244_v78_: _dafny.Seq
                d_244_v78_ = _dafny.SeqWithoutIsStrInference([d_239_v71_])
                d_245_v79_: _dafny.Seq
                def iife28_():
                    coll16_ = _dafny.Map()
                    compr_16_: int
                    for compr_16_ in (d_242_v75_).keys.Elements:
                        d_246_v73_: int = compr_16_
                        if (d_246_v73_) in (d_242_v75_):
                            coll16_[default__.safeDivisionInt(d_246_v73_, d_145_v12_)] = (d_239_v71_).f4
                    return _dafny.Map(coll16_)
                def iife29_():
                    coll17_ = _dafny.Map()
                    compr_17_: int
                    for compr_17_ in (d_241_v74_).keys.Elements:
                        d_247_v77_: int = compr_17_
                        if (d_247_v77_) in (d_241_v74_):
                            coll17_[default__.safeDivisionInt(d_247_v77_, 600)] = d_225_v64_.f8
                    return _dafny.Map(coll17_)
                d_245_v79_ = _dafny.SeqWithoutIsStrInference([iife28_()
                , _dafny.Map({len(d_243_v76_): d_225_v64_.f8}), d_241_v74_, (iife29_()
                ).set(d_145_v12_, d_133_v1_), _dafny.Map({len(d_244_v78_): d_239_v71_.f3})])
                d_248_v80_: _dafny.Array
                d_249_v81_: int
                out3_: _dafny.Array
                out4_: int
                out3_, out4_ = (d_225_v64_).m0((d_240_v72_).issubset(d_240_v72_), ((d_245_v79_)[default__.safeIndex((d_225_v64_).f9, len(d_245_v79_))]).set((d_225_v64_).f9, d_225_v64_.f8), (d_225_v64_).f9, d_131_globalState_)
                d_248_v80_ = out3_
                d_249_v81_ = out4_
            elif True:
                d_250___mcc_h9_ = source5_.cf12
                d_251_cf12_ = d_250___mcc_h9_
                d_252_v82_: T0
                nw41_ = C2()
                nw41_.ctor__(d_225_v64_.f8, False, True, d_133_v1_)
                d_252_v82_ = nw41_
                d_253_v83_: int
                out5_: int
                out5_ = (d_225_v64_).m1(((d_225_v64_).f9) * ((d_225_v64_).f9), d_252_v82_, d_131_globalState_)
                d_253_v83_ = out5_
                index28_ = default__.safeIndex(754, (d_227_v66_).length(0))
                (d_227_v66_)[index28_] = (d_252_v82_).f4
                d_254_v84_: _dafny.Seq
                d_254_v84_ = _dafny.SeqWithoutIsStrInference([d_253_v83_])
                d_255_v85_: _dafny.MultiSet
                d_255_v85_ = _dafny.MultiSet([d_133_v1_])
                d_256_v86_: _dafny.Map
                d_256_v86_ = _dafny.Map({d_254_v84_: d_255_v85_})
                d_257_v87_: _dafny.Array
                nw42_ = _dafny.Array(None, 1)
                nw42_[int(0)] = d_256_v86_
                d_257_v87_ = nw42_
                index29_ = default__.safeIndex(960, (d_257_v87_).length(0))
                (d_257_v87_)[index29_] = d_256_v86_
                d_258_v88_: _dafny.Map
                d_258_v88_ = _dafny.Map({d_225_v64_.f8: d_225_v64_.f8})
                d_259_v89_: _dafny.Map
                d_259_v89_ = d_256_v86_
                index30_ = default__.safeIndex(754, (d_227_v66_).length(0))
                index31_ = default__.safeIndex(960, (d_257_v87_).length(0))
                rhs31_ = (d_225_v64_).f9
                rhs32_ = ((d_258_v88_)[d_225_v64_.f8] if (d_225_v64_.f8) in (d_258_v88_) else d_225_v64_.f8)
                rhs33_ = (d_259_v89_)
                lhs23_ = d_227_v66_
                lhs24_ = default__.safeIndex(754, (d_227_v66_).length(0))
                lhs25_ = d_257_v87_
                lhs26_ = default__.safeIndex(960, (d_257_v87_).length(0))
                d_145_v12_ = rhs31_
                lhs23_[lhs24_] = rhs32_
                lhs25_[lhs26_] = rhs33_
                d_260_v90_: _dafny.Seq
                d_260_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
                d_261_v91_: _dafny.Map
                d_261_v91_ = _dafny.Map({len(d_260_v90_): (d_227_v66_)[default__.safeIndex(754, (d_227_v66_).length(0))]})
                d_262_v92_: _dafny.Array
                d_263_v93_: int
                out6_: _dafny.Array
                out7_: int
                out6_, out7_ = (d_252_v82_).m0(d_225_v64_.f8, (d_261_v91_) | (d_261_v91_), -15, d_131_globalState_)
                d_262_v92_ = out6_
                d_263_v93_ = out7_
                d_264_v94_: D17
                d_264_v94_ = D17_DC38(d_260_v90_)
                d_260_v90_ = (d_264_v94_).cf56
            d_265_v95_: _dafny.Seq
            d_265_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgg"))
            d_266_v96_: C5
            nw43_ = C5()
            nw43_.ctor__(d_265_v95_, (0) - (d_145_v12_), True, ((d_225_v64_).f9) == (d_145_v12_))
            d_266_v96_ = nw43_
            d_266_v96_ = d_266_v96_
            d_267_v97_: bool
            d_268_v98_: int
            d_269_v99_: _dafny.Array
            out8_: bool
            out9_: int
            out10_: _dafny.Array
            out8_, out9_, out10_ = (d_225_v64_).m6(d_133_v1_, d_131_globalState_)
            d_267_v97_ = out8_
            d_268_v98_ = out9_
            d_269_v99_ = out10_
            d_270_v100_: _dafny.Map
            d_270_v100_ = _dafny.Map({d_225_v64_.f8: d_266_v96_.f7})
            d_271_v101_: D1
            d_271_v101_ = D1_DC5(d_133_v1_, True, len((d_265_v95_).set(default__.safeIndex(((d_270_v100_)[d_133_v1_] if (d_133_v1_) in (d_270_v100_) else d_268_v98_), len(d_265_v95_)), d_224_v63_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgwyb"))))
            d_272_v102_: _dafny.Seq
            d_272_v102_ = _dafny.SeqWithoutIsStrInference([default__.fm31(d_225_v64_.f8, d_266_v96_.f7, d_133_v1_, d_132_v0_, d_131_globalState_)])
            d_273_v103_: D0
            d_273_v103_ = D0_DC0(d_145_v12_)
            d_265_v95_ = default__.fm30(d_271_v101_, (d_272_v102_)[default__.safeIndex(d_266_v96_.f7, len(d_272_v102_))], ((d_273_v103_).cf0 if d_267_v97_ else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcroj")))), d_131_globalState_)
        elif True:
            d_274_v104_: int
            d_274_v104_ = 29
            d_275_v105_: _dafny.MultiSet
            d_275_v105_ = _dafny.MultiSet([(0) - (d_274_v104_)])
            d_274_v104_ = ((d_275_v105_)[d_274_v104_] if (d_274_v104_) in (d_275_v105_) else (d_274_v104_) - (d_274_v104_))
            d_276_v106_: _dafny.Map
            d_276_v106_ = _dafny.Map({d_274_v104_: d_274_v104_})
            d_277_v107_: _dafny.MultiSet
            d_277_v107_ = _dafny.MultiSet([d_276_v106_, d_276_v106_])
            d_278_v108_: _dafny.Array
            nw44_ = _dafny.Array(None, 21)
            nw44_[int(0)] = d_133_v1_
            nw44_[int(1)] = d_133_v1_
            nw44_[int(2)] = d_133_v1_
            nw44_[int(3)] = d_133_v1_
            nw44_[int(4)] = d_133_v1_
            nw44_[int(5)] = d_133_v1_
            nw44_[int(6)] = d_133_v1_
            nw44_[int(7)] = d_133_v1_
            nw44_[int(8)] = True
            nw44_[int(9)] = d_133_v1_
            nw44_[int(10)] = d_133_v1_
            nw44_[int(11)] = not(d_133_v1_)
            nw44_[int(12)] = (d_132_v0_)[default__.safeIndex(d_274_v104_, len(d_132_v0_))]
            nw44_[int(13)] = d_133_v1_
            nw44_[int(14)] = d_133_v1_
            nw44_[int(15)] = default__.fm21(d_274_v104_, d_277_v107_, d_131_globalState_)
            nw44_[int(16)] = False
            nw44_[int(17)] = d_133_v1_
            nw44_[int(18)] = d_133_v1_
            nw44_[int(19)] = not(d_133_v1_)
            nw44_[int(20)] = d_133_v1_
            d_278_v108_ = nw44_
            d_279_v109_: _dafny.Map
            d_279_v109_ = _dafny.Map({d_278_v108_: d_133_v1_})
            d_280_v110_: D19
            d_280_v110_ = D19_DC45(d_274_v104_, d_274_v104_, d_133_v1_, d_133_v1_, d_132_v0_)
            d_281_v111_: C0
            nw45_ = C0()
            nw45_.ctor__()
            d_281_v111_ = nw45_
            d_282_v112_: _dafny.Map
            d_282_v112_ = _dafny.Map({d_280_v110_: d_281_v111_})
            d_279_v109_ = (d_279_v109_).set(d_278_v108_, (d_282_v112_) == (d_282_v112_))
            d_283_v113_: D11
            d_284_v114_: int
            d_285_v115_: _dafny.Seq
            d_286_v116_: C2
            out11_: D11
            out12_: int
            out13_: _dafny.Seq
            out14_: C2
            out11_, out12_, out13_, out14_ = default__.m12((default__.fm31(d_133_v1_, d_274_v104_, d_133_v1_, default__.fm34(d_274_v104_, default__.fm31(d_133_v1_, (d_281_v111_).fm10(d_274_v104_, d_133_v1_, d_275_v105_, d_131_globalState_), d_133_v1_, d_132_v0_, d_131_globalState_), d_131_globalState_), d_131_globalState_)) - (d_274_v104_), d_133_v1_, (d_133_v1_) and (not(not(not(d_133_v1_)))), (d_274_v104_) >= (len(d_279_v109_)), d_131_globalState_)
            d_283_v113_ = out11_
            d_284_v114_ = out12_
            d_285_v115_ = out13_
            d_286_v116_ = out14_
            if not ((d_286_v116_).f10) or ((d_286_v116_).f10):
                d_287_v117_: _dafny.Seq
                d_287_v117_ = _dafny.SeqWithoutIsStrInference([((d_275_v105_)[len(d_132_v0_)] if (len(d_132_v0_)) in (d_275_v105_) else d_284_v114_), d_284_v114_, d_284_v114_, d_284_v114_])
                d_288_v118_: _dafny.Map
                d_288_v118_ = _dafny.Map({d_274_v104_: d_286_v116_})
                d_289_v119_: _dafny.Map
                d_289_v119_ = _dafny.Map({d_288_v118_: d_132_v0_})
                d_290_v120_: _dafny.Map
                d_290_v120_ = _dafny.Map({(d_287_v117_)[default__.safeIndex(d_274_v104_, len(d_287_v117_))]: d_289_v119_})
                d_290_v120_ = (d_290_v120_).set(d_274_v104_, (d_289_v119_) | (d_289_v119_))
                d_291_v121_: _dafny.Seq
                d_291_v121_ = _dafny.SeqWithoutIsStrInference([d_278_v108_])
                d_292_v122_: _dafny.Array
                def lambda9_(d_293_v104_):
                    def lambda10_(d_294_i9_):
                        return (d_294_i9_) + (d_293_v104_)

                    return lambda10_

                init4_ = lambda9_(d_274_v104_)
                nw46_ = _dafny.Array(None, 29)
                for i0_4_ in range(nw46_.length(0)):
                    nw46_[i0_4_] = init4_(i0_4_)
                d_292_v122_ = nw46_
                d_295_v123_: D21
                d_295_v123_ = D21_DC52(d_274_v104_, d_292_v122_)
                d_296_v124_: _dafny.Seq
                d_296_v124_ = _dafny.SeqWithoutIsStrInference([d_295_v123_])
                d_278_v108_ = (d_291_v121_)[default__.safeIndex((len(d_296_v124_)) * (d_274_v104_), len(d_291_v121_))]
                d_133_v1_ = not((d_286_v116_).f11)
                d_297_v125_: _dafny.Array
                nw47_ = _dafny.Array(_dafny.CodePoint('D'), 26)
                d_297_v125_ = nw47_
                d_298_v126_: D21
                d_298_v126_ = D21_DC51(d_297_v125_)
                d_299_v127_: _dafny.MultiSet
                d_299_v127_ = _dafny.MultiSet([d_298_v126_, d_298_v126_, d_298_v126_, d_298_v126_])
                d_300_v128_: T0
                nw48_ = C3()
                nw48_.ctor__((d_286_v116_).f11, (d_286_v116_).f10)
                d_300_v128_ = nw48_
                d_301_v129_: _dafny.Map
                d_301_v129_ = _dafny.Map({(d_286_v116_).f10: d_300_v128_})
                d_302_v130_: str
                d_302_v130_ = _dafny.CodePoint('u')
                d_303_v131_: _dafny.MultiSet
                d_303_v131_ = _dafny.MultiSet([d_302_v130_])
                d_304_v132_: _dafny.Map
                d_304_v132_ = _dafny.Map({d_284_v114_: True})
                d_305_v133_: _dafny.Map
                d_305_v133_ = _dafny.Map({(d_286_v116_).f11: d_133_v1_})
                rhs34_ = d_299_v127_
                rhs35_ = d_133_v1_
                rhs36_ = (((d_286_v116_).fm16(len(d_301_v129_), (d_286_v116_).f10, (d_303_v131_).cardinality, d_304_v132_, d_131_globalState_)) + (d_284_v114_)) + ((0) - ((d_284_v114_ if False else len(d_305_v133_))))
                lhs27_ = d_131_globalState_
                d_299_v127_ = rhs34_
                lhs27_.f0 = rhs35_
                d_274_v104_ = rhs36_
                d_302_v130_ = d_302_v130_
            elif True:
                d_306_v134_: _dafny.Array
                nw49_ = _dafny.Array(int(0), 10)
                d_306_v134_ = nw49_
                d_307_v135_: _dafny.MultiSet
                d_307_v135_ = _dafny.MultiSet([d_306_v134_, d_306_v134_])
                d_308_v136_: _dafny.MultiSet
                d_308_v136_ = _dafny.MultiSet([d_307_v135_, d_307_v135_, _dafny.MultiSet([d_306_v134_, d_306_v134_])])
                d_309_v137_: _dafny.MultiSet
                d_309_v137_ = _dafny.MultiSet([(d_286_v116_).f11, d_133_v1_])
                d_284_v114_ = ((d_308_v136_)[d_307_v135_] if (d_307_v135_) in (d_308_v136_) else (0) - (((d_309_v137_)[False] if (False) in (d_309_v137_) else d_274_v104_)))
                index32_ = default__.safeIndex(685, (d_306_v134_).length(0))
                (d_306_v134_)[index32_] = d_274_v104_
                index33_ = default__.safeIndex(88, (d_278_v108_).length(0))
                (d_278_v108_)[index33_] = (d_284_v114_) > (d_274_v104_)
                index34_ = default__.safeIndex(685, (d_306_v134_).length(0))
                index35_ = default__.safeIndex(88, (d_278_v108_).length(0))
                rhs37_ = d_284_v114_
                rhs38_ = d_133_v1_
                lhs28_ = d_306_v134_
                lhs29_ = default__.safeIndex(685, (d_306_v134_).length(0))
                lhs30_ = d_278_v108_
                lhs31_ = default__.safeIndex(88, (d_278_v108_).length(0))
                lhs28_[lhs29_] = rhs37_
                lhs30_[lhs31_] = rhs38_
                d_310_v138_: _dafny.Seq
                d_310_v138_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
                d_311_v139_: C1
                nw50_ = C1()
                nw50_.ctor__(d_310_v138_, (False) or ((d_278_v108_)[default__.safeIndex(88, (d_278_v108_).length(0))]))
                d_311_v139_ = nw50_
                d_312_v140_: _dafny.Array
                nw51_ = _dafny.Array(None, 9)
                nw51_[int(0)] = d_306_v134_
                nw51_[int(1)] = d_306_v134_
                nw51_[int(2)] = d_306_v134_
                nw51_[int(3)] = d_306_v134_
                nw51_[int(4)] = d_306_v134_
                nw51_[int(5)] = d_306_v134_
                nw51_[int(6)] = d_306_v134_
                nw51_[int(7)] = d_306_v134_
                nw51_[int(8)] = d_306_v134_
                d_312_v140_ = nw51_
                index36_ = default__.safeIndex(705, (d_312_v140_).length(0))
                (d_312_v140_)[index36_] = (d_306_v134_ if (d_286_v116_).f10 else d_306_v134_)
                index37_ = default__.safeIndex(705, (d_312_v140_).length(0))
                (d_312_v140_)[index37_] = d_306_v134_
                d_274_v104_ = d_284_v114_
            d_313_v141_: C6
            nw52_ = C6()
            nw52_.ctor__((d_286_v116_).f10, (d_286_v116_).f11, (d_286_v116_).f10)
            d_313_v141_ = nw52_
            d_314_v142_: _dafny.Map
            d_314_v142_ = _dafny.Map({d_313_v141_: d_133_v1_})
            d_315_v143_: _dafny.Map
            d_315_v143_ = _dafny.Map({d_284_v114_: (d_286_v116_).f10})
            d_314_v142_ = (d_314_v142_).set(d_313_v141_, ((d_315_v143_)[d_284_v114_] if (d_284_v114_) in (d_315_v143_) else (d_286_v116_).f11))
        d_316_v144_: int
        d_316_v144_ = 937
        d_317_v145_: _dafny.Set
        d_317_v145_ = _dafny.Set({d_133_v1_, d_133_v1_})
        d_318_v146_: _dafny.Seq
        d_318_v146_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_317_v145_))])
        d_316_v144_ = default__.safeModuloInt(d_316_v144_, default__.safeDivisionInt(d_316_v144_, (d_318_v146_)[default__.safeIndex(len(d_132_v0_), len(d_318_v146_))]))
        d_319_v147_: _dafny.Array
        nw53_ = _dafny.Array(int(0), 12)
        d_319_v147_ = nw53_
        index38_ = default__.safeIndex(624, (d_319_v147_).length(0))
        (d_319_v147_)[index38_] = d_316_v144_
        d_320_v148_: _dafny.Seq
        d_320_v148_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
        d_321_v149_: D1
        d_321_v149_ = D1_DC5(d_133_v1_, d_133_v1_, d_316_v144_, len(d_320_v148_))
        pat_let_tv5_ = d_316_v144_
        d_322_v150_: D19
        def iife30_(_pat_let6_0):
            def iife31_(d_323_dt__update__tmp_h2_):
                def iife32_(_pat_let7_0):
                    def iife33_(d_324_dt__update_hcf10_h0_):
                        def iife34_(_pat_let8_0):
                            def iife35_(d_325_dt__update_hcf11_h0_):
                                return D1_DC5((d_323_dt__update__tmp_h2_).cf8, (d_323_dt__update__tmp_h2_).cf9, d_324_dt__update_hcf10_h0_, d_325_dt__update_hcf11_h0_)
                            return iife35_(_pat_let8_0)
                        return iife34_(pat_let_tv5_)
                    return iife33_(_pat_let7_0)
                return iife32_(747)
            return iife31_(_pat_let6_0)
        d_322_v150_ = D19_DC46(d_316_v144_, (D22_DC56(d_316_v144_, False)).cf90, d_133_v1_, iife30_(d_321_v149_), (0) - (d_316_v144_))
        pat_let_tv6_ = d_320_v148_
        pat_let_tv7_ = d_316_v144_
        index39_ = default__.safeIndex(624, (d_319_v147_).length(0))
        def lambda11_(source6_):
            if source6_.is_DC45:
                d_326___mcc_h10_ = source6_.cf69
                d_327___mcc_h11_ = source6_.cf70
                d_328___mcc_h12_ = source6_.cf71
                d_329___mcc_h13_ = source6_.cf72
                d_330___mcc_h14_ = source6_.cf73
                d_331_cf73_ = d_330___mcc_h14_
                d_332_cf72_ = d_329___mcc_h13_
                d_333_cf71_ = d_328___mcc_h12_
                d_334_cf70_ = d_327___mcc_h11_
                d_335_cf69_ = d_326___mcc_h10_
                return d_334_cf70_
            elif source6_.is_DC46:
                d_336___mcc_h15_ = source6_.cf74
                d_337___mcc_h16_ = source6_.cf75
                d_338___mcc_h17_ = source6_.cf76
                d_339___mcc_h18_ = source6_.cf77
                d_340___mcc_h19_ = source6_.cf78
                d_341_cf78_ = d_340___mcc_h19_
                d_342_cf77_ = d_339___mcc_h18_
                d_343_cf76_ = d_338___mcc_h17_
                d_344_cf75_ = d_337___mcc_h16_
                d_345_cf74_ = d_336___mcc_h15_
                return (964) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))))
            elif source6_.is_DC44:
                d_346___mcc_h20_ = source6_.cf68
                d_347_cf68_ = d_346___mcc_h20_
                return len(pat_let_tv6_)
            elif True:
                d_348___mcc_h21_ = source6_.cf79
                d_349_cf79_ = d_348___mcc_h21_
                return (0) - (pat_let_tv7_)

        rhs39_ = lambda11_(d_322_v150_)
        rhs40_ = d_316_v144_
        rhs41_ = d_316_v144_
        lhs32_ = d_319_v147_
        lhs33_ = default__.safeIndex(624, (d_319_v147_).length(0))
        lhs32_[lhs33_] = rhs39_
        d_316_v144_ = rhs40_
        d_316_v144_ = rhs41_
        d_350_v151_: _dafny.Array
        nw54_ = _dafny.Array(False, 12)
        d_350_v151_ = nw54_
        d_350_v151_ = d_350_v151_
        d_351_v152_: C4
        nw55_ = C4()
        nw55_.ctor__((len(d_317_v145_)) <= ((d_319_v147_)[default__.safeIndex(624, (d_319_v147_).length(0))]), (d_319_v147_)[default__.safeIndex(624, (d_319_v147_).length(0))], d_133_v1_, d_133_v1_)
        d_351_v152_ = nw55_
        d_351_v152_ = d_351_v152_
        d_352_v153_: T0
        nw56_ = C4()
        nw56_.ctor__(d_133_v1_, d_316_v144_, False, d_133_v1_)
        d_352_v153_ = nw56_
        d_353_v154_: int
        out15_: int
        out15_ = (d_351_v152_).m1(((d_319_v147_)[default__.safeIndex(624, (d_319_v147_).length(0))]) * (-84), d_352_v153_, d_131_globalState_)
        d_353_v154_ = out15_
        hi2_ = (d_353_v154_) - (d_316_v144_)
        for d_354_i10_ in range(d_353_v154_, hi2_):
            index40_ = default__.safeIndex(624, (d_319_v147_).length(0))
            (d_319_v147_)[index40_] = (d_351_v152_).fm0(d_131_globalState_)
            index41_ = default__.safeIndex(624, (d_319_v147_).length(0))
            (d_319_v147_)[index41_] = d_353_v154_
            d_353_v154_ = 763
            d_355_v155_: C6
            nw57_ = C6()
            nw57_.ctor__(d_351_v152_.f8, False, (d_352_v153_).f4)
            d_355_v155_ = nw57_
            d_356_v156_: _dafny.Set
            d_356_v156_ = _dafny.Set({d_355_v155_})
            d_357_v157_: _dafny.Map
            d_357_v157_ = _dafny.Map({not(d_133_v1_): d_356_v156_})
            d_358_v158_: D7
            d_358_v158_ = D7_DC17(d_354_i10_, d_352_v153_.f3, d_351_v152_.f8, len(((d_357_v157_)[True] if (True) in (d_357_v157_) else d_356_v156_)))
            pat_let_tv8_ = d_321_v149_
            def iife36_(_pat_let9_0):
                def iife37_(d_359_dt__update__tmp_h3_):
                    def iife38_(_pat_let10_0):
                        def iife39_(d_360_dt__update_hcf27_h0_):
                            return D7_DC17((d_359_dt__update__tmp_h3_).cf26, d_360_dt__update_hcf27_h0_, (d_359_dt__update__tmp_h3_).cf28, (d_359_dt__update__tmp_h3_).cf29)
                        return iife39_(_pat_let10_0)
                    return iife38_((pat_let_tv8_).cf9)
                return iife37_(_pat_let9_0)
            source7_ = iife36_(d_358_v158_)
            if source7_.is_DC17:
                d_361___mcc_h22_ = source7_.cf26
                d_362___mcc_h23_ = source7_.cf27
                d_363___mcc_h24_ = source7_.cf28
                d_364___mcc_h25_ = source7_.cf29
                d_365_cf29_ = d_364___mcc_h25_
                d_366_cf28_ = d_363___mcc_h24_
                d_367_cf27_ = d_362___mcc_h23_
                d_368_cf26_ = d_361___mcc_h22_
                d_369_v159_: _dafny.Map
                d_369_v159_ = _dafny.Map({d_355_v155_: (d_351_v152_).f9})
                d_369_v159_ = (d_369_v159_).set(d_355_v155_, default__.safeDivisionInt((d_319_v147_)[default__.safeIndex(624, (d_319_v147_).length(0))], d_368_cf26_))
                (d_352_v153_).f3 = ((_dafny.Set({d_352_v153_.f3})).isdisjoint(d_317_v145_) if d_355_v155_.f5 else d_352_v153_.f3)
                d_351_v152_ = d_351_v152_
                d_370_v160_: C1
                nw58_ = C1()
                nw58_.ctor__((d_320_v148_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))), d_351_v152_.f8)
                d_370_v160_ = nw58_
            elif True:
                d_371___mcc_h26_ = source7_.cf25
                d_372_cf25_ = d_371___mcc_h26_
                index42_ = default__.safeIndex(624, (d_319_v147_).length(0))
                (d_319_v147_)[index42_] = len(((d_320_v148_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "po")))) + (d_320_v148_))
                index43_ = default__.safeIndex(624, (d_319_v147_).length(0))
                (d_319_v147_)[index43_] = default__.safeModuloInt((d_351_v152_).f9, d_353_v154_)
                d_373_v161_: _dafny.Map
                d_373_v161_ = _dafny.Map({len(d_320_v148_): d_352_v153_.f3})
                d_374_v162_: _dafny.Map
                d_374_v162_ = _dafny.Map({len(d_373_v161_): d_352_v153_.f3})
                d_375_v163_: _dafny.Array
                d_376_v164_: int
                out16_: _dafny.Array
                out17_: int
                out16_, out17_ = (d_352_v153_).m0(d_352_v153_.f3, ((d_374_v162_).set((d_351_v152_).f9, d_351_v152_.f8)) | (d_373_v161_), d_316_v144_, d_131_globalState_)
                d_375_v163_ = out16_
                d_376_v164_ = out17_
                d_377_v165_: C1
                nw59_ = C1()
                nw59_.ctor__(d_320_v148_, d_133_v1_)
                d_377_v165_ = nw59_
                d_377_v165_ = d_377_v165_
        d_378_v166_: _dafny.Map
        d_378_v166_ = _dafny.Map({d_351_v152_.f8: d_133_v1_})
        (d_352_v153_).f3 = ((d_378_v166_)[d_133_v1_] if (d_133_v1_) in (d_378_v166_) else d_133_v1_)
        d_379_v169_: D17
        def iife40_():
            coll18_ = _dafny.Map()
            compr_18_: _dafny.Seq
            for compr_18_ in (_dafny.SeqWithoutIsStrInference([d_318_v146_])).Elements:
                d_380_v168_: _dafny.Seq = compr_18_
                if (d_380_v168_) in (_dafny.SeqWithoutIsStrInference([d_318_v146_])):
                    coll18_[d_380_v168_] = _dafny.CodePoint('l')
            return _dafny.Map(coll18_)
        d_379_v169_ = D17_DC37(iife40_()
)
        d_381_v170_: _dafny.Set
        d_381_v170_ = _dafny.Set({d_379_v169_})
        d_382_v171_: _dafny.Seq
        d_382_v171_ = _dafny.SeqWithoutIsStrInference([d_381_v170_])
        d_383_v173_: D23
        def iife41_():
            def iife43_():
                coll21_ = _dafny.Map()
                compr_21_: _dafny.Set
                for compr_21_ in (d_382_v171_).Elements:
                    d_384_v167_: _dafny.Set = compr_21_
                    if (d_384_v167_) in (d_382_v171_):
                        coll21_[d_384_v167_] = d_353_v154_
                return _dafny.Map(coll21_)
            coll19_ = _dafny.Set()
            def iife42_():
                coll20_ = _dafny.Map()
                compr_20_: _dafny.Set
                for compr_20_ in (d_382_v171_).Elements:
                    d_384_v167_: _dafny.Set = compr_20_
                    if (d_384_v167_) in (d_382_v171_):
                        coll20_[d_384_v167_] = d_353_v154_
                return _dafny.Map(coll20_)
            compr_19_: _dafny.Set
            for compr_19_ in (iife42_()
            ).keys.Elements:
                d_385_v172_: _dafny.Set = compr_19_
                if (d_385_v172_) in (iife43_()
                ):
                    coll19_ = coll19_.union(_dafny.Set([d_385_v172_]))
            return _dafny.Set(coll19_)
        d_383_v173_ = D23_DC58(iife41_()
)
        pat_let_tv9_ = d_351_v152_
        pat_let_tv10_ = d_320_v148_
        pat_let_tv11_ = d_351_v152_
        pat_let_tv12_ = d_318_v146_
        pat_let_tv13_ = d_316_v144_
        pat_let_tv14_ = d_318_v146_
        pat_let_tv15_ = d_351_v152_
        pat_let_tv16_ = d_320_v148_
        pat_let_tv17_ = d_352_v153_
        pat_let_tv18_ = d_316_v144_
        pat_let_tv19_ = d_316_v144_
        def lambda12_(source8_):
            if source8_.is_DC59:
                d_386___mcc_h27_ = source8_.cf93
                d_387___mcc_h28_ = source8_.cf94
                d_388___mcc_h29_ = source8_.cf95
                d_389_cf95_ = d_388___mcc_h29_
                d_390_cf94_ = d_387___mcc_h28_
                d_391_cf93_ = d_386___mcc_h27_
                return ((pat_let_tv9_).f9) - (len(pat_let_tv10_))
            elif source8_.is_DC60:
                d_392___mcc_h30_ = source8_.cf96
                d_393___mcc_h31_ = source8_.cf97
                d_394_cf97_ = d_393___mcc_h31_
                d_395_cf96_ = d_392___mcc_h30_
                return ((0) - ((pat_let_tv11_).f9)) - ((pat_let_tv12_)[default__.safeIndex(pat_let_tv13_, len(pat_let_tv14_))])
            elif source8_.is_DC61:
                d_396___mcc_h32_ = source8_.cf98
                d_397_cf98_ = d_396___mcc_h32_
                return (D24_DC63((pat_let_tv15_).f9, pat_let_tv16_, (pat_let_tv17_).f4, _dafny.CodePoint('s'))).cf100
            elif True:
                d_398___mcc_h33_ = source8_.cf92
                d_399_cf92_ = d_398___mcc_h33_
                return (pat_let_tv18_) * (pat_let_tv19_)

        d_353_v154_ = lambda12_(d_383_v173_)
        d_400_v174_: _dafny.Array
        nw60_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
        d_400_v174_ = nw60_
        index44_ = default__.safeIndex(230, (d_400_v174_).length(0))
        (d_400_v174_)[index44_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvlmx"))
        index45_ = default__.safeIndex(230, (d_400_v174_).length(0))
        (d_400_v174_)[index45_] = d_320_v148_
        d_353_v154_ = d_353_v154_
        index46_ = default__.safeIndex(497, (d_350_v151_).length(0))
        (d_350_v151_)[index46_] = d_351_v152_.f8
        index47_ = default__.safeIndex(497, (d_350_v151_).length(0))
        (d_350_v151_)[index47_] = (d_320_v148_) <= (d_320_v148_)
        d_401_v175_: _dafny.Array
        nw61_ = _dafny.Array(None, 7)
        nw61_[int(0)] = d_318_v146_
        nw61_[int(1)] = d_318_v146_
        nw61_[int(2)] = d_318_v146_
        nw61_[int(3)] = (d_318_v146_) + (_dafny.SeqWithoutIsStrInference([(d_351_v152_).f9 for d_402_i11_ in range(default__.abs(203))]))
        nw61_[int(4)] = d_318_v146_
        nw61_[int(5)] = default__.fm19(d_131_globalState_)
        nw61_[int(6)] = d_318_v146_
        d_401_v175_ = nw61_
        d_401_v175_ = d_401_v175_
        d_403_v176_: _dafny.Map
        d_403_v176_ = _dafny.Map({(d_352_v153_).f4: (d_351_v152_).f9})
        d_404_v177_: _dafny.MultiSet
        d_404_v177_ = _dafny.MultiSet([(d_403_v176_).set(True, 160)])
        d_405_v178_: C2
        nw62_ = C2()
        nw62_.ctor__((d_353_v154_) != ((d_319_v147_)[default__.safeIndex(624, (d_319_v147_).length(0))]), d_352_v153_.f3, (d_404_v177_).ispropersubset(d_404_v177_), d_352_v153_.f3)
        d_405_v178_ = nw62_
        d_406_v179_: str
        d_406_v179_ = _dafny.CodePoint('a')
        d_407_v180_: _dafny.Seq
        out18_: _dafny.Seq
        out18_ = (d_351_v152_).m2(d_406_v179_, d_352_v153_, d_131_globalState_)
        d_407_v180_ = out18_
        _dafny.print(_dafny.string_of(d_131_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_133_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_134_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_316_v144_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_317_v145_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_318_v146_) == (_dafny.SeqWithoutIsStrInference([-1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_319_v147_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_320_v148_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_321_v149_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_321_v149_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_321_v149_).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_321_v149_).cf11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_322_v150_).cf74))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_322_v150_).cf75))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_322_v150_).cf76))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_322_v150_).cf77).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_322_v150_).cf77).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_322_v150_).cf77).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_322_v150_).cf77).cf11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_322_v150_).cf78))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_350_v151_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_351_v152_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_351_v152_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_352_v153_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_352_v153_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_353_v154_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_378_v166_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_379_v169_).cf55) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([-1]): _dafny.CodePoint('l')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_381_v170_) == (_dafny.Set({D17_DC37(_dafny.Map({_dafny.SeqWithoutIsStrInference([-1]): _dafny.CodePoint('l')}))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_382_v171_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({D17_DC37(_dafny.Map({_dafny.SeqWithoutIsStrInference([-1]): _dafny.CodePoint('l')}))})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_383_v173_).cf92) == (_dafny.Set({_dafny.Set({D17_DC37(_dafny.Map({_dafny.SeqWithoutIsStrInference([-1]): _dafny.CodePoint('l')}))})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_400_v174_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_401_v175_)[0]) == (_dafny.SeqWithoutIsStrInference([-1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_401_v175_)[1]) == (_dafny.SeqWithoutIsStrInference([-1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_401_v175_)[2]) == (_dafny.SeqWithoutIsStrInference([-1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_401_v175_)[3]) == (_dafny.SeqWithoutIsStrInference([-1, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965, 965]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_401_v175_)[4]) == (_dafny.SeqWithoutIsStrInference([-1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_401_v175_)[5]) == (_dafny.SeqWithoutIsStrInference([705]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_401_v175_)[6]) == (_dafny.SeqWithoutIsStrInference([-1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_403_v176_) == (_dafny.Map({True: 965}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_404_v177_) == (_dafny.MultiSet([_dafny.Map({True: 160})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_405_v178_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_405_v178_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_406_v179_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v180_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, _dafny.CodePoint('D'))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
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
        return lambda: D1_DC4(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6)
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
        return lambda: D2_DC8(_dafny.Array(None, 0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf13)}, {self.cf14.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(_dafny.MultiSet({}), D1.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC12(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

class D4_DC12(D4, NamedTuple('DC12', [('cf19', Any), ('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

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
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)

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
        return lambda: D7_DC17(int(0), False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)

class D7_DC17(D7, NamedTuple('DC17', [('cf26', Any), ('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC16(D7, NamedTuple('DC16', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC19(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)

class D8_DC19(D8, NamedTuple('DC19', [('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC18(D8, NamedTuple('DC18', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)

class D9_DC21(D9, NamedTuple('DC21', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC23()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D10_DC22)

class D10_DC23(D10, NamedTuple('DC23', [])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23)
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC22(D10, NamedTuple('DC22', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC22({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC22) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC25(_dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D11_DC24)

class D11_DC25(D11, NamedTuple('DC25', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC24(D11, NamedTuple('DC24', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC24({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC24) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC27(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D12_DC27)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D12_DC26)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)

class D12_DC27(D12, NamedTuple('DC27', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC27({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC27) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC26(D12, NamedTuple('DC26', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC26({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC26) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC28(D12, NamedTuple('DC28', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D13_DC29)

class D13_DC29(D13, NamedTuple('DC29', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC29({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC29) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC31(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D14_DC31)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D14_DC30)

class D14_DC31(D14, NamedTuple('DC31', [('cf44', Any), ('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC31({_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC31) and self.cf44 == __o.cf44 and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC30(D14, NamedTuple('DC30', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC30({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC30) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC33(_dafny.Array(None, 0), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D15_DC33)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D15_DC32)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D15_DC34)

class D15_DC33(D15, NamedTuple('DC33', [('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC33({_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC33) and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC32(D15, NamedTuple('DC32', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC32({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC32) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC34(D15, NamedTuple('DC34', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC34({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC34) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC36(_dafny.Set({}), False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D16_DC36)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D16_DC35)

class D16_DC36(D16, NamedTuple('DC36', [('cf51', Any), ('cf52', Any), ('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC36({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC36) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC35(D16, NamedTuple('DC35', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC35({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC35) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC38(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D17_DC38)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D17_DC39)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D17_DC37)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D17_DC40)

class D17_DC38(D17, NamedTuple('DC38', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC38({self.cf56.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC38) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC39(D17, NamedTuple('DC39', [('cf57', Any), ('cf58', Any), ('cf59', Any), ('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC39({_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC39) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC37(D17, NamedTuple('DC37', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC37({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC37) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC40(D17, NamedTuple('DC40', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC40({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC40) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC42(int(0))
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

class D18_DC42(D18, NamedTuple('DC42', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC42({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC42) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC43(D18, NamedTuple('DC43', [('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC43({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC43) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC41(D18, NamedTuple('DC41', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC41({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC41) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC45(int(0), int(0), False, False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D19_DC45)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D19_DC46)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D19_DC44)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D19_DC47)

class D19_DC45(D19, NamedTuple('DC45', [('cf69', Any), ('cf70', Any), ('cf71', Any), ('cf72', Any), ('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC45({_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC45) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72 and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC46(D19, NamedTuple('DC46', [('cf74', Any), ('cf75', Any), ('cf76', Any), ('cf77', Any), ('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC46({_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC46) and self.cf74 == __o.cf74 and self.cf75 == __o.cf75 and self.cf76 == __o.cf76 and self.cf77 == __o.cf77 and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC44(D19, NamedTuple('DC44', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC44({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC44) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC47(D19, NamedTuple('DC47', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC47({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC47) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC49()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D20_DC49)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D20_DC50)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D20_DC48)

class D20_DC49(D20, NamedTuple('DC49', [])):
    def __dafnystr__(self) -> str:
        return f'D20.DC49'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC49)
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC50(D20, NamedTuple('DC50', [('cf81', Any), ('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC50({_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC50) and self.cf81 == __o.cf81 and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC48(D20, NamedTuple('DC48', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC48({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC48) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC52(int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D21_DC52)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D21_DC53)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D21_DC51)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D21_DC54)

class D21_DC52(D21, NamedTuple('DC52', [('cf84', Any), ('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC52({_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC52) and self.cf84 == __o.cf84 and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC53(D21, NamedTuple('DC53', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC53({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC53) and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC51(D21, NamedTuple('DC51', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC51({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC51) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC54(D21, NamedTuple('DC54', [('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC54({_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC54) and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC56(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D22_DC56)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D22_DC55)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D22_DC57)

class D22_DC56(D22, NamedTuple('DC56', [('cf89', Any), ('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC56({_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC56) and self.cf89 == __o.cf89 and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC55(D22, NamedTuple('DC55', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC55({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC55) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC57(D22, NamedTuple('DC57', [('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC57({_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC57) and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC59(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D23_DC59)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D23_DC60)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D23_DC61)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D23_DC58)

class D23_DC59(D23, NamedTuple('DC59', [('cf93', Any), ('cf94', Any), ('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC59({self.cf93.VerbatimString(True)}, {_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC59) and self.cf93 == __o.cf93 and self.cf94 == __o.cf94 and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC60(D23, NamedTuple('DC60', [('cf96', Any), ('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC60({_dafny.string_of(self.cf96)}, {_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC60) and self.cf96 == __o.cf96 and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC61(D23, NamedTuple('DC61', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC61({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC61) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC58(D23, NamedTuple('DC58', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC58({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC58) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC63(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D24_DC63)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D24_DC64)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D24_DC62)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D24_DC65)

class D24_DC63(D24, NamedTuple('DC63', [('cf100', Any), ('cf101', Any), ('cf102', Any), ('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC63({_dafny.string_of(self.cf100)}, {self.cf101.VerbatimString(True)}, {_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC63) and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102 and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC64(D24, NamedTuple('DC64', [('cf104', Any), ('cf105', Any), ('cf106', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC64({_dafny.string_of(self.cf104)}, {_dafny.string_of(self.cf105)}, {_dafny.string_of(self.cf106)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC64) and self.cf104 == __o.cf104 and self.cf105 == __o.cf105 and self.cf106 == __o.cf106
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC62(D24, NamedTuple('DC62', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC62({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC62) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC65(D24, NamedTuple('DC65', [('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC65({_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC65) and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D25_DC66)

class D25_DC66(D25, NamedTuple('DC66', [('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC66({_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC66) and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D26_DC67)

class D26_DC67(D26, NamedTuple('DC67', [('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC67({_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC67) and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC69(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D27_DC69)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D27_DC68)

class D27_DC69(D27, NamedTuple('DC69', [('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC69({_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC69) and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC68(D27, NamedTuple('DC68', [('cf110', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC68({_dafny.string_of(self.cf110)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC68) and self.cf110 == __o.cf110
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f3(self):
        return self._f3
    @f3.setter
    def f3(self, value):
        self._f3 = value
    def m0(self, p0, p1, p2, globalState):
        pass


class T1:
    pass
    def m1(self, p0, p1, globalState):
        pass

    def m2(self, p0, p1, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self._f1: bool = False
        self._f2: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2

    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm10(self, p0, p1, p2, globalState):
        return len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmyjnxlg")))

    def fm11(self, p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbgjfgxoq"))


class C1(T1):
    def  __init__(self):
        self._f12: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f13: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f12, f13):
        (self)._f12 = f12
        (self)._f13 = f13

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        if (p0) > (p0):
            d_408_v0_: _dafny.Array
            nw63_ = _dafny.Array(False, 8)
            d_408_v0_ = nw63_
            d_408_v0_ = d_408_v0_
            if p1.f3:
                d_409_v1_: _dafny.Map
                d_409_v1_ = _dafny.Map({p1.f3: len((self).f12)})
                d_409_v1_ = (d_409_v1_).set(p1.f3, (p1).fm0(globalState))
                d_410_v3_: _dafny.Seq
                d_410_v3_ = _dafny.SeqWithoutIsStrInference([164])
                d_411_v4_: _dafny.Map
                d_411_v4_ = _dafny.Map({p0: p0})
                d_412_v5_: _dafny.MultiSet
                def iife44_():
                    coll22_ = _dafny.Map()
                    compr_22_: int
                    for compr_22_ in (d_410_v3_).Elements:
                        d_413_v2_: int = compr_22_
                        if (d_413_v2_) in (d_410_v3_):
                            coll22_[(d_413_v2_) * (p0)] = p0
                    return _dafny.Map(coll22_)
                d_412_v5_ = _dafny.MultiSet([iife44_()
                , _dafny.Map({443: p0}), d_411_v4_])
                rhs42_ = not (default__.fm21(647, d_412_v5_, globalState)) or (p1.f3)
                rhs43_ = p0
                rhs44_ = p0
                rhs45_ = (p1).fm0(globalState)
                lhs34_ = globalState
                lhs34_.f0 = rhs42_
                r0 = rhs43_
                r0 = rhs44_
                r0 = rhs45_
                (globalState).f0 = (default__.safeModuloInt(p0, p0)) < (p0)
                d_414_v6_: C0
                nw64_ = C0()
                nw64_.ctor__()
                d_414_v6_ = nw64_
                d_415_v7_: _dafny.Map
                d_415_v7_ = _dafny.Map({d_411_v4_: (p1).f4})
                d_416_v8_: _dafny.MultiSet
                d_416_v8_ = _dafny.MultiSet([p1.f3])
                d_415_v7_ = (d_415_v7_).set(d_411_v4_, (d_416_v8_).ispropersubset(d_416_v8_))
            elif True:
                d_417_v9_: D1
                d_417_v9_ = D1_DC6()
                d_417_v9_ = d_417_v9_
                d_418_v10_: _dafny.Array
                nw65_ = _dafny.Array(_dafny.MultiSet({}), 25)
                d_418_v10_ = nw65_
                index48_ = default__.safeIndex(429, (d_418_v10_).length(0))
                (d_418_v10_)[index48_] = _dafny.MultiSet([(p1).f4])
                d_419_v11_: _dafny.MultiSet
                d_419_v11_ = _dafny.MultiSet([p1.f3])
                index49_ = default__.safeIndex(429, (d_418_v10_).length(0))
                (d_418_v10_)[index49_] = (d_419_v11_) - (_dafny.MultiSet([(p1).f4, (p1).f4, (p1).f4, (p1).f4]))
                d_420_v12_: _dafny.Array
                def lambda13_(d_421_p0_):
                    def lambda14_(d_422_i0_):
                        return _dafny.Set({d_421_p0_})

                    return lambda14_

                init5_ = lambda13_(p0)
                nw66_ = _dafny.Array(None, 13)
                for i0_5_ in range(nw66_.length(0)):
                    nw66_[i0_5_] = init5_(i0_5_)
                d_420_v12_ = nw66_
                d_420_v12_ = d_420_v12_
                d_423_v13_: _dafny.Set
                d_423_v13_ = _dafny.Set({(p1).f4, not(p1.f3)})
                (p1).f3 = (d_423_v13_).isdisjoint(d_423_v13_)
                (p1).f3 = (p1).f4
            r0 = p0
            d_424_v14_: _dafny.Array
            def lambda15_(d_425_i1_):
                return _dafny.CodePoint('f')

            init6_ = lambda15_
            nw67_ = _dafny.Array(None, 22)
            for i0_6_ in range(nw67_.length(0)):
                nw67_[i0_6_] = init6_(i0_6_)
            d_424_v14_ = nw67_
            d_426_v15_: str
            d_426_v15_ = _dafny.CodePoint('f')
            index50_ = default__.safeIndex(917, (d_424_v14_).length(0))
            (d_424_v14_)[index50_] = d_426_v15_
            index51_ = default__.safeIndex(917, (d_424_v14_).length(0))
            (d_424_v14_)[index51_] = d_426_v15_
            d_427_v16_: _dafny.Set
            d_427_v16_ = _dafny.Set({(p1).f4, (p1).f4})
            d_428_v17_: _dafny.Map
            d_428_v17_ = _dafny.Map({p0: len((self).f12)})
            d_429_v18_: _dafny.Seq
            d_429_v18_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
            d_430_v19_: _dafny.Seq
            d_430_v19_ = _dafny.SeqWithoutIsStrInference([p0, len(d_427_v16_), 660, (0) - ((131) - (((d_428_v17_)[p0] if (p0) in (d_428_v17_) else len(d_429_v18_)))), p0])
            rhs46_ = 293
            rhs47_ = ((self).f12) == ((self).f12)
            rhs48_ = _dafny.SeqWithoutIsStrInference([p0, (p1).fm0(globalState), p0, -118])
            lhs35_ = globalState
            r0 = rhs46_
            lhs35_.f0 = rhs47_
            d_429_v18_ = rhs48_
        elif True:
            if p1.f3:
                d_431_v20_: str
                d_431_v20_ = _dafny.CodePoint('u')
                d_432_v21_: _dafny.Set
                d_432_v21_ = _dafny.Set({(p1).f4, (self).f13})
                d_433_v22_: _dafny.Array
                nw68_ = _dafny.Array(False, 9)
                d_433_v22_ = nw68_
                index52_ = default__.safeIndex(930, (d_433_v22_).length(0))
                (d_433_v22_)[index52_] = (p1).f4
                d_434_v23_: _dafny.Map
                d_434_v23_ = _dafny.Map({(self).f13: p0})
                index53_ = default__.safeIndex(930, (d_433_v22_).length(0))
                rhs49_ = ((self).f12)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_434_v23_, (d_434_v23_).set((p1).f4, (0) - (p0))])), len((self).f12))]
                rhs50_ = _dafny.Set({(p1).f4, p1.f3})
                rhs51_ = ((p0) - (p0)) >= (p0)
                rhs52_ = default__.fm21(339, default__.fm22(p0, p0, len(d_434_v23_), globalState), globalState)
                lhs36_ = d_433_v22_
                lhs37_ = default__.safeIndex(930, (d_433_v22_).length(0))
                lhs38_ = globalState
                d_431_v20_ = rhs49_
                d_432_v21_ = rhs50_
                lhs36_[lhs37_] = rhs51_
                lhs38_.f0 = rhs52_
                d_435_v24_: _dafny.Array
                nw69_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_435_v24_ = nw69_
                d_436_v25_: _dafny.Array
                nw70_ = _dafny.Array(_dafny.Seq({}), 6)
                d_436_v25_ = nw70_
                index54_ = default__.safeIndex(232, (d_435_v24_).length(0))
                (d_435_v24_)[index54_] = d_436_v25_
                index55_ = default__.safeIndex(232, (d_435_v24_).length(0))
                nw71_ = _dafny.Array(_dafny.Seq({}), 13)
                (d_435_v24_)[index55_] = nw71_
                d_437_v26_: _dafny.Seq
                d_437_v26_ = _dafny.SeqWithoutIsStrInference([((p1).f4 if p1.f3 else True)])
                d_438_v27_: _dafny.Seq
                d_438_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cw"))
                d_439_v28_: _dafny.Map
                d_439_v28_ = _dafny.Map({p0: d_431_v20_})
                d_440_v29_: _dafny.Map
                d_440_v29_ = _dafny.Map({p1.f3: d_431_v20_})
                rhs53_ = (_dafny.SeqWithoutIsStrInference([(p1).f4, (p1).f4, False, (d_433_v22_)[default__.safeIndex(930, (d_433_v22_).length(0))], ((self).f12) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ter")))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([(p1).f4, (p1).f4, False, (d_433_v22_)[default__.safeIndex(930, (d_433_v22_).length(0))], ((self).f12) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ter")))]))), False)
                rhs54_ = ((_dafny.SeqWithoutIsStrInference([d_431_v20_ for d_441_i2_ in range(default__.abs(-167))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjmrx")))) + ((self).f12)
                rhs55_ = p0
                rhs56_ = ((d_439_v28_)[len(_dafny.SeqWithoutIsStrInference([d_431_v20_ for d_442_i3_ in range(default__.abs(50))]))] if (len(_dafny.SeqWithoutIsStrInference([d_431_v20_ for d_443_i3_ in range(default__.abs(50))]))) in (d_439_v28_) else ((d_440_v29_)[(self).f13] if ((self).f13) in (d_440_v29_) else default__.fm23((p1).f4, 2, globalState)))
                rhs57_ = (((self).f12) + ((self).f12)) == ((d_438_v27_) + ((self).f12))
                lhs39_ = p1
                d_437_v26_ = rhs53_
                d_438_v27_ = rhs54_
                r0 = rhs55_
                d_431_v20_ = rhs56_
                lhs39_.f3 = rhs57_
                d_444_v30_: _dafny.MultiSet
                d_444_v30_ = _dafny.MultiSet([p0, 134, len((_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_438_v27_)).cardinality, 382, len(d_434_v23_)])).set(default__.safeIndex(len(d_438_v27_), len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_438_v27_)).cardinality, 382, len(d_434_v23_)]))), p0))])
                d_445_v31_: _dafny.Set
                d_445_v31_ = _dafny.Set({(d_444_v30_).cardinality, p0, p0, p0})
                d_445_v31_ = d_445_v31_
                d_438_v27_ = (self).f12
            elif True:
                d_446_v32_: _dafny.Map
                d_446_v32_ = _dafny.Map({p0: (p1).f4})
                d_447_v33_: _dafny.Set
                d_447_v33_ = _dafny.Set({p1.f3})
                (p1).f3 = (((d_446_v32_)[p0] if (p0) in (d_446_v32_) else (p1).f4)) in (d_447_v33_)
                d_448_v34_: D1
                d_448_v34_ = D1_DC5(p1.f3, (p1).f4, default__.safeModuloInt(p0, p0), p0)
                rhs58_ = ((0) - ((0) - (p0))) < (p0)
                rhs59_ = d_448_v34_
                rhs60_ = (p1).fm0(globalState)
                lhs40_ = p1
                lhs40_.f3 = rhs58_
                d_448_v34_ = rhs59_
                r0 = rhs60_
                d_449_v35_: _dafny.Map
                d_449_v35_ = _dafny.Map({p0: p0})
                d_450_v37_: _dafny.Array
                nw72_ = _dafny.Array(None, 9)
                nw72_[int(0)] = d_449_v35_
                nw72_[int(1)] = (d_449_v35_).set(-880, p0)
                nw72_[int(2)] = d_449_v35_
                nw72_[int(3)] = _dafny.Map({p0: p0})
                nw72_[int(4)] = (_dafny.Map({p0: 256})) | (d_449_v35_)
                nw72_[int(5)] = d_449_v35_
                def iife45_():
                    coll23_ = _dafny.Map()
                    compr_23_: int
                    for compr_23_ in _dafny.IntegerRange(958, 360):
                        d_451_v36_: int = compr_23_
                        if ((958) <= (d_451_v36_)) and ((d_451_v36_) < (360)):
                            coll23_[(d_451_v36_) + (p0)] = p0
                    return _dafny.Map(coll23_)
                nw72_[int(6)] = iife45_()
                
                nw72_[int(7)] = d_449_v35_
                nw72_[int(8)] = d_449_v35_
                d_450_v37_ = nw72_
                rhs61_ = (p1).f4
                rhs62_ = p0
                rhs63_ = d_450_v37_
                rhs64_ = 290
                lhs41_ = globalState
                lhs41_.f0 = rhs61_
                r0 = rhs62_
                d_450_v37_ = rhs63_
                r0 = rhs64_
                (p1).f3 = (p1).f4
                (p1).f3 = False
            d_452_v38_: C0
            nw73_ = C0()
            nw73_.ctor__()
            d_452_v38_ = nw73_
            d_453_v39_: _dafny.MultiSet
            d_453_v39_ = _dafny.MultiSet([d_452_v38_])
            d_454_v40_: _dafny.Seq
            d_454_v40_ = _dafny.SeqWithoutIsStrInference([d_452_v38_])
            d_455_v41_: D2
            d_455_v41_ = D2_DC7(d_453_v39_)
            d_456_v42_: _dafny.Array
            nw74_ = _dafny.Array(None, 10)
            nw74_[int(0)] = d_453_v39_
            nw74_[int(1)] = _dafny.MultiSet([(d_454_v40_)[default__.safeIndex(p0, len(d_454_v40_))], d_452_v38_, d_452_v38_, d_452_v38_, d_452_v38_])
            nw74_[int(2)] = d_453_v39_
            nw74_[int(3)] = (d_453_v39_) | (_dafny.MultiSet([d_452_v38_]))
            nw74_[int(4)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_452_v38_]))
            nw74_[int(5)] = (d_453_v39_) - ((_dafny.MultiSet([d_452_v38_])).set(d_452_v38_, default__.abs(p0)))
            nw74_[int(6)] = (d_453_v39_).set(d_452_v38_, default__.abs(p0))
            nw74_[int(7)] = _dafny.MultiSet((d_454_v40_ if not(False) else d_454_v40_))
            nw74_[int(8)] = (d_455_v41_).cf12
            nw74_[int(9)] = d_453_v39_
            d_456_v42_ = nw74_
            index56_ = default__.safeIndex(637, (d_456_v42_).length(0))
            (d_456_v42_)[index56_] = d_453_v39_
            index57_ = default__.safeIndex(637, (d_456_v42_).length(0))
            (d_456_v42_)[index57_] = d_453_v39_
            d_457_v43_: _dafny.Set
            d_457_v43_ = _dafny.Set({p0, p0})
            d_458_v44_: _dafny.Map
            d_458_v44_ = _dafny.Map({(not((self).f13) if not((p1).f4) else (p1).f4): len(d_457_v43_)})
            d_458_v44_ = (d_458_v44_).set(False, (p0) - (p0))
            r0 = (0) - ((0) - (p0))
            d_459_v45_: _dafny.Seq
            d_459_v45_ = _dafny.SeqWithoutIsStrInference([(self).f12, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlvktxw")), (self).f12, (self).f12])
            d_460_v46_: str
            d_460_v46_ = _dafny.CodePoint('o')
            d_461_v47_: _dafny.Set
            d_461_v47_ = _dafny.Set({(self).f13})
            d_462_v48_: _dafny.Seq
            d_462_v48_ = _dafny.SeqWithoutIsStrInference([(d_459_v45_)[default__.safeIndex(len(_dafny.Set({p1.f3, p1.f3, p1.f3})), len(d_459_v45_))], (((d_452_v38_).fm11(d_460_v46_, d_461_v47_, globalState)).set(default__.safeIndex((0) - (p0), len((d_452_v38_).fm11(d_460_v46_, d_461_v47_, globalState))), ((self).f12)[default__.safeIndex(p0, len((self).f12))])).set(default__.safeIndex(p0, len(((d_452_v38_).fm11(d_460_v46_, d_461_v47_, globalState)).set(default__.safeIndex((0) - (p0), len((d_452_v38_).fm11(d_460_v46_, d_461_v47_, globalState))), ((self).f12)[default__.safeIndex(p0, len((self).f12))]))), d_460_v46_)])
            d_463_v49_: _dafny.Seq
            d_463_v49_ = _dafny.SeqWithoutIsStrInference([True])
            d_462_v48_ = _dafny.SeqWithoutIsStrInference([((self).f12) + ((self).f12), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))).set(default__.safeIndex(len(_dafny.Set({(p1).f4, (self).f13, (p1).f4, (d_463_v49_)[default__.safeIndex(267, len(d_463_v49_))], True})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")))), d_460_v46_)])
        d_464_v51_: str
        d_464_v51_ = _dafny.CodePoint('v')
        d_465_v52_: _dafny.Set
        d_465_v52_ = _dafny.Set({d_464_v51_, d_464_v51_, d_464_v51_, d_464_v51_, d_464_v51_})
        d_466_v53_: _dafny.Map
        d_466_v53_ = _dafny.Map({_dafny.CodePoint('l'): default__.fm24(globalState)})
        def iife46_():
            coll24_ = _dafny.Map()
            compr_24_: str
            for compr_24_ in (d_465_v52_).Elements:
                d_467_v50_: str = compr_24_
                if (d_467_v50_) in (d_465_v52_):
                    coll24_[d_467_v50_] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len((self).f12), p0, len(_dafny.Map({False: not(p1.f3)})), p0, p0]))
            return _dafny.Map(coll24_)
        if ((iife46_()
        ) | (d_466_v53_)) != (d_466_v53_):
            d_468_v54_: _dafny.Set
            d_468_v54_ = _dafny.Set({p1.f3})
            if (len(d_468_v54_)) == (p0):
                d_469_v55_: _dafny.MultiSet
                d_469_v55_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfsrcr")), (self).f12, (self).f12, (self).f12, (self).f12])
                r0 = default__.safeModuloInt(p0, (0) - ((d_469_v55_).cardinality))
                r0 = p0
                d_470_v56_: _dafny.Seq
                d_470_v56_ = _dafny.SeqWithoutIsStrInference([p0, p0, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_471_i4_ in range(default__.abs(173))]))])
                d_472_v57_: _dafny.Map
                d_472_v57_ = _dafny.Map({-484: d_470_v56_})
                d_472_v57_ = (d_472_v57_).set(p0, (d_470_v56_) + (_dafny.SeqWithoutIsStrInference([p0])))
                d_473_v58_: _dafny.Map
                d_473_v58_ = _dafny.Map({p0: (D0_DC2(False, p0, 607)).cf3})
                d_473_v58_ = (d_473_v58_).set(p0, ((p1).f4) == (p1.f3))
                d_474_v59_: _dafny.Array
                nw75_ = _dafny.Array(_dafny.Seq({}), 13)
                d_474_v59_ = nw75_
                index58_ = default__.safeIndex(935, (d_474_v59_).length(0))
                (d_474_v59_)[index58_] = d_470_v56_
                index59_ = default__.safeIndex(935, (d_474_v59_).length(0))
                (d_474_v59_)[index59_] = d_470_v56_
            elif True:
                d_464_v51_ = d_464_v51_
                d_475_v60_: _dafny.MultiSet
                d_475_v60_ = _dafny.MultiSet([p0, p0])
                d_475_v60_ = d_475_v60_
                d_476_v61_: _dafny.Array
                def lambda16_(d_477_i5_):
                    return True

                init7_ = lambda16_
                nw76_ = _dafny.Array(None, 2)
                for i0_7_ in range(nw76_.length(0)):
                    nw76_[i0_7_] = init7_(i0_7_)
                d_476_v61_ = nw76_
                index60_ = default__.safeIndex(213, (d_476_v61_).length(0))
                (d_476_v61_)[index60_] = (p1).f4
                index61_ = default__.safeIndex(213, (d_476_v61_).length(0))
                (d_476_v61_)[index61_] = (((0) - (p0) if not(True) else p0)) != (p0)
                (p1).f3 = p1.f3
                d_478_v62_: _dafny.Map
                d_478_v62_ = _dafny.Map({p0: p0})
                d_479_v63_: _dafny.MultiSet
                d_479_v63_ = _dafny.MultiSet([d_478_v62_])
                (globalState).f0 = default__.fm21(p0, d_479_v63_, globalState)
            d_480_v64_: _dafny.Map
            d_480_v64_ = _dafny.Map({d_464_v51_: (p1).f4})
            r0 = len(d_480_v64_)
            r0 = p0
            d_481_v65_: _dafny.Array
            nw77_ = _dafny.Array(int(0), 11)
            d_481_v65_ = nw77_
            index62_ = default__.safeIndex(531, (d_481_v65_).length(0))
            (d_481_v65_)[index62_] = p0
            index63_ = default__.safeIndex(531, (d_481_v65_).length(0))
            rhs65_ = p0
            rhs66_ = p0
            lhs42_ = d_481_v65_
            lhs43_ = default__.safeIndex(531, (d_481_v65_).length(0))
            r0 = rhs65_
            lhs42_[lhs43_] = rhs66_
            d_482_v66_: C0
            nw78_ = C0()
            nw78_.ctor__()
            d_482_v66_ = nw78_
        elif True:
            d_483_v67_: _dafny.Array
            nw79_ = _dafny.Array(False, 4)
            d_483_v67_ = nw79_
            d_484_v68_: _dafny.Seq
            d_484_v68_ = _dafny.SeqWithoutIsStrInference([d_483_v67_, d_483_v67_, d_483_v67_])
            r0 = len(d_484_v68_)
            (globalState).f0 = (p1).f4
            d_485_v69_: _dafny.Seq
            d_485_v69_ = _dafny.SeqWithoutIsStrInference([p1.f3])
            d_486_v70_: _dafny.Map
            d_486_v70_ = _dafny.Map({len(d_485_v69_): p1.f3})
            d_487_v71_: _dafny.Map
            d_487_v71_ = _dafny.Map({p0: p0})
            d_488_v72_: _dafny.MultiSet
            d_488_v72_ = _dafny.MultiSet([d_487_v71_, d_487_v71_])
            d_489_v73_: _dafny.Set
            d_489_v73_ = _dafny.Set({p1.f3, False, default__.fm21(len(d_486_v70_), d_488_v72_, globalState)})
            (globalState).f0 = (d_489_v73_).ispropersubset(d_489_v73_)
            r0 = p0
            r0 = p0
        d_490_v74_: _dafny.Array
        nw80_ = _dafny.Array(False, 14)
        d_490_v74_ = nw80_
        d_490_v74_ = d_490_v74_
        d_491_v75_: _dafny.Seq
        d_491_v75_ = _dafny.SeqWithoutIsStrInference([p1.f3])
        d_492_v76_: _dafny.Set
        d_492_v76_ = _dafny.Set({(p1).f4})
        d_493_v77_: _dafny.MultiSet
        d_493_v77_ = _dafny.MultiSet([p1.f3, (p1).f4, (self).f13])
        d_494_v78_: _dafny.Map
        d_494_v78_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, p1.f3])): p0})
        d_495_v79_: _dafny.MultiSet
        d_495_v79_ = _dafny.MultiSet([p0])
        d_496_v80_: _dafny.Map
        d_496_v80_ = _dafny.Map({(self).f13: (D0_DC0(((d_495_v79_)[p0] if (p0) in (d_495_v79_) else p0))).cf0})
        d_497_v81_: _dafny.Map
        d_497_v81_ = _dafny.Map({(p1).f4: len(d_496_v80_)})
        d_498_v82_: _dafny.Map
        d_498_v82_ = _dafny.Map({(self).f13: (self).f13})
        d_499_v83_: _dafny.Seq
        d_499_v83_ = _dafny.SeqWithoutIsStrInference([len(d_498_v82_), p0])
        d_500_v84_: _dafny.Array
        nw81_ = _dafny.Array(None, 24)
        nw81_[int(0)] = p0
        nw81_[int(1)] = (p0 if (d_491_v75_)[default__.safeIndex(p0, len(d_491_v75_))] else len(d_492_v76_))
        nw81_[int(2)] = p0
        nw81_[int(3)] = p0
        nw81_[int(4)] = p0
        nw81_[int(5)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "myku")))) - (p0)
        nw81_[int(6)] = ((d_493_v77_)[(p1).f4] if ((p1).f4) in (d_493_v77_) else p0)
        nw81_[int(7)] = default__.safeDivisionInt(((d_494_v78_)[p0] if (p0) in (d_494_v78_) else p0), p0)
        nw81_[int(8)] = p0
        nw81_[int(9)] = (p0 if (p1).f4 else p0)
        nw81_[int(10)] = p0
        nw81_[int(11)] = (0) - ((p0) * (p0))
        nw81_[int(12)] = 373
        nw81_[int(13)] = 974
        nw81_[int(14)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrv")))
        nw81_[int(15)] = (p0) - (p0)
        nw81_[int(16)] = len((self).f12)
        nw81_[int(17)] = p0
        nw81_[int(18)] = 1
        nw81_[int(19)] = ((d_496_v80_)[p1.f3] if (p1.f3) in (d_496_v80_) else p0)
        nw81_[int(20)] = -114
        nw81_[int(21)] = len(default__.fm25(globalState))
        nw81_[int(22)] = p0
        nw81_[int(23)] = len(d_499_v83_)
        d_500_v84_ = nw81_
        index64_ = default__.safeIndex(188, (d_500_v84_).length(0))
        (d_500_v84_)[index64_] = len(((self).f12) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywt"))))
        d_501_v85_: _dafny.MultiSet
        d_501_v85_ = _dafny.MultiSet([((self).f12) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))])
        index65_ = default__.safeIndex(188, (d_500_v84_).length(0))
        (d_500_v84_)[index65_] = ((d_501_v85_)[(self).f12] if ((self).f12) in (d_501_v85_) else p0)
        if (p1).f4:
            d_502_v86_: _dafny.Map
            d_502_v86_ = _dafny.Map({p1.f3: d_492_v76_})
            d_502_v86_ = ((d_502_v86_) | ((d_502_v86_).set((p1).f4, d_492_v76_))).set((p1).f4, d_492_v76_)
            d_503_v87_: _dafny.Seq
            d_503_v87_ = _dafny.SeqWithoutIsStrInference([d_500_v84_, d_500_v84_])
            rhs67_ = (d_503_v87_)[default__.safeIndex(p0, len(d_503_v87_))]
            rhs68_ = (p1).f4
            lhs44_ = globalState
            d_500_v84_ = rhs67_
            lhs44_.f0 = rhs68_
            d_504_v88_: _dafny.Set
            d_504_v88_ = _dafny.Set({p0})
            index66_ = default__.safeIndex(673, (d_490_v74_).length(0))
            (d_490_v74_)[index66_] = (d_504_v88_).issubset(d_504_v88_)
            index67_ = default__.safeIndex(673, (d_490_v74_).length(0))
            (d_490_v74_)[index67_] = (self).f13
            d_505_v89_: _dafny.Seq
            d_505_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kttjwxwja"))
            d_505_v89_ = (_dafny.SeqWithoutIsStrInference([(D0_DC1(False, d_464_v51_)).cf2 for d_506_i6_ in range(default__.abs(584))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([(D0_DC1(False, d_464_v51_)).cf2 for d_507_i6_ in range(default__.abs(584))]))), default__.fm23(p1.f3, (0) - (p0), globalState))
            r0 = (0) - ((len(d_492_v76_)) - (default__.safeModuloInt((d_500_v84_)[default__.safeIndex(188, (d_500_v84_).length(0))], len(d_505_v89_))))
        elif True:
            d_508_v90_: _dafny.Seq
            d_508_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvigjoi"))
            d_508_v90_ = ((d_508_v90_) + ((d_508_v90_).set(default__.safeIndex((0) - (p0), len(d_508_v90_)), d_464_v51_))) + ((d_508_v90_) + ((self).f12))
            d_509_v91_: C0
            nw82_ = C0()
            nw82_.ctor__()
            d_509_v91_ = nw82_
            d_510_v92_: D1
            d_510_v92_ = D1_DC6()
            source9_ = d_510_v92_
            if source9_.is_DC4:
                d_511___mcc_h0_ = source9_.cf7
                d_512_cf7_ = d_511___mcc_h0_
                r0 = (0) - (default__.safeDivisionInt(d_512_cf7_, (d_500_v84_)[default__.safeIndex(188, (d_500_v84_).length(0))]))
                d_512_cf7_ = d_512_cf7_
                d_513_v93_: _dafny.Array
                nw83_ = _dafny.Array(D1.default()(), 14)
                d_513_v93_ = nw83_
                index68_ = default__.safeIndex(587, (d_513_v93_).length(0))
                (d_513_v93_)[index68_] = d_510_v92_
                index69_ = default__.safeIndex(587, (d_513_v93_).length(0))
                (d_513_v93_)[index69_] = d_510_v92_
                (globalState).f0 = False
            elif source9_.is_DC5:
                d_514___mcc_h1_ = source9_.cf8
                d_515___mcc_h2_ = source9_.cf9
                d_516___mcc_h3_ = source9_.cf10
                d_517___mcc_h4_ = source9_.cf11
                d_518_cf11_ = d_517___mcc_h4_
                d_519_cf10_ = d_516___mcc_h3_
                d_520_cf9_ = d_515___mcc_h2_
                d_521_cf8_ = d_514___mcc_h1_
                (globalState).f0 = d_520_cf9_
                r0 = (p0) * (default__.safeDivisionInt((d_500_v84_)[default__.safeIndex(188, (d_500_v84_).length(0))], (d_500_v84_)[default__.safeIndex(188, (d_500_v84_).length(0))]))
                d_522_v94_: _dafny.Array
                def lambda17_(d_523_v90_):
                    def lambda18_(d_524_i7_):
                        return _dafny.Set({len(d_523_v90_)})

                    return lambda18_

                init8_ = lambda17_(d_508_v90_)
                nw84_ = _dafny.Array(None, 24)
                for i0_8_ in range(nw84_.length(0)):
                    nw84_[i0_8_] = init8_(i0_8_)
                d_522_v94_ = nw84_
                d_525_v95_: _dafny.Set
                d_525_v95_ = _dafny.Set({(d_500_v84_)[default__.safeIndex(188, (d_500_v84_).length(0))]})
                index70_ = default__.safeIndex(132, (d_522_v94_).length(0))
                (d_522_v94_)[index70_] = (d_525_v95_) | (d_525_v95_)
                d_526_v96_: _dafny.MultiSet
                d_526_v96_ = _dafny.MultiSet([d_509_v91_])
                d_527_v97_: D2
                d_527_v97_ = D2_DC7(d_526_v96_)
                d_528_v98_: _dafny.Map
                d_528_v98_ = _dafny.Map({d_527_v97_: (p1).fm0(globalState)})
                index71_ = default__.safeIndex(188, (d_500_v84_).length(0))
                index72_ = default__.safeIndex(132, (d_522_v94_).length(0))
                index73_ = default__.safeIndex(188, (d_500_v84_).length(0))
                rhs69_ = (d_509_v91_).fm10(d_518_cf11_, (p1).f4, d_495_v79_, globalState)
                rhs70_ = d_525_v95_
                rhs71_ = (((d_495_v79_).cardinality if True else ((d_528_v98_)[d_527_v97_] if (d_527_v97_) in (d_528_v98_) else (0) - ((d_495_v79_).cardinality)))) + (d_519_cf10_)
                rhs72_ = p1.f3
                lhs45_ = d_500_v84_
                lhs46_ = default__.safeIndex(188, (d_500_v84_).length(0))
                lhs47_ = d_522_v94_
                lhs48_ = default__.safeIndex(132, (d_522_v94_).length(0))
                lhs49_ = d_500_v84_
                lhs50_ = default__.safeIndex(188, (d_500_v84_).length(0))
                lhs45_[lhs46_] = rhs69_
                lhs47_[lhs48_] = rhs70_
                lhs49_[lhs50_] = rhs71_
                d_520_cf9_ = rhs72_
                d_518_cf11_ = d_518_cf11_
            elif source9_.is_DC6:
                (globalState).f0 = p1.f3
                d_529_v99_: _dafny.Set
                d_529_v99_ = _dafny.Set({(d_499_v83_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qyylnovqs"))), len(d_499_v83_))]})
                d_496_v80_ = (d_496_v80_).set((p1).f4, len((d_529_v99_).intersection(_dafny.Set({len(d_492_v76_)}))))
                index74_ = default__.safeIndex(188, (d_500_v84_).length(0))
                (d_500_v84_)[index74_] = p0
                d_530_v100_: _dafny.Map
                d_530_v100_ = _dafny.Map({d_491_v75_: d_494_v78_})
                d_531_v101_: _dafny.Map
                d_531_v101_ = _dafny.Map({d_530_v100_: p0})
                d_531_v101_ = (d_531_v101_).set(d_530_v100_, (0) - (p0))
            elif True:
                d_532___mcc_h5_ = source9_.cf6
                d_533_cf6_ = d_532___mcc_h5_
                index75_ = default__.safeIndex(716, (d_490_v74_).length(0))
                (d_490_v74_)[index75_] = (self).f13
                index76_ = default__.safeIndex(716, (d_490_v74_).length(0))
                (d_490_v74_)[index76_] = ((p1).f4 if (p1.f3 if True else p1.f3) else not(p1.f3))
                d_534_v102_: C0
                nw85_ = C0()
                nw85_.ctor__()
                d_534_v102_ = nw85_
                d_535_v103_: _dafny.Map
                d_535_v103_ = _dafny.Map({p0: (p1).f4})
                d_497_v81_ = (d_497_v81_).set(((d_535_v103_)[(d_500_v84_)[default__.safeIndex(188, (d_500_v84_).length(0))]] if ((d_500_v84_)[default__.safeIndex(188, (d_500_v84_).length(0))]) in (d_535_v103_) else not((p1).f4)), len(_dafny.SeqWithoutIsStrInference([(p1).f4])))
                index77_ = default__.safeIndex(716, (d_490_v74_).length(0))
                (d_490_v74_)[index77_] = (p1.f3 if not (p1.f3) or (p1.f3) else not(((p1).f4) or (p1.f3)))
            index78_ = default__.safeIndex(188, (d_500_v84_).length(0))
            (d_500_v84_)[index78_] = ((((d_494_v78_)[(p1).fm0(globalState)] if ((p1).fm0(globalState)) in (d_494_v78_) else p0)) + (p0)) + ((d_500_v84_)[default__.safeIndex(188, (d_500_v84_).length(0))])
            d_536_v104_: C0
            nw86_ = C0()
            nw86_.ctor__()
            d_536_v104_ = nw86_
        d_537_v105_: D0
        d_537_v105_ = D0_DC2((p1).f4, len((self).f12), p0)
        pat_let_tv20_ = p1
        pat_let_tv21_ = p1
        pat_let_tv22_ = p0
        pat_let_tv23_ = p0
        pat_let_tv24_ = p1
        pat_let_tv25_ = d_465_v52_
        pat_let_tv26_ = p1
        pat_let_tv27_ = p1
        pat_let_tv28_ = d_500_v84_
        pat_let_tv29_ = d_500_v84_
        pat_let_tv30_ = d_494_v78_
        def lambda19_(source10_):
            if source10_.is_DC1:
                d_538___mcc_h6_ = source10_.cf1
                d_539___mcc_h7_ = source10_.cf2
                d_540_cf2_ = d_539___mcc_h7_
                d_541_cf1_ = d_538___mcc_h6_
                def iife47_():
                    coll25_ = _dafny.Map()
                    compr_25_: int
                    for compr_25_ in _dafny.IntegerRange(377, 23):
                        d_542_v106_: int = compr_25_
                        if ((377) <= (d_542_v106_)) and ((d_542_v106_) < (23)):
                            coll25_[default__.safeDivisionInt(d_542_v106_, 441)] = pat_let_tv20_.f3
                    return _dafny.Map(coll25_)
                def iife48_():
                    coll26_ = _dafny.Map()
                    compr_26_: int
                    for compr_26_ in _dafny.IntegerRange(-707, 92):
                        d_543_v107_: int = compr_26_
                        if ((-707) <= (d_543_v107_)) and ((d_543_v107_) < (92)):
                            coll26_[default__.safeModuloInt(d_543_v107_, pat_let_tv22_)] = (pat_let_tv21_).f4
                    return _dafny.Map(coll26_)
                return (_dafny.MultiSet([iife47_()
                , iife48_()
                ])).issubset(_dafny.MultiSet([_dafny.Map({pat_let_tv23_: pat_let_tv24_.f3}), _dafny.Map({len(pat_let_tv25_): (pat_let_tv26_).f4})]))
            elif source10_.is_DC2:
                d_544___mcc_h8_ = source10_.cf3
                d_545___mcc_h9_ = source10_.cf4
                d_546___mcc_h10_ = source10_.cf5
                d_547_cf5_ = d_546___mcc_h10_
                d_548_cf4_ = d_545___mcc_h9_
                d_549_cf3_ = d_544___mcc_h8_
                return not(((pat_let_tv27_).f4) or (True))
            elif True:
                d_550___mcc_h11_ = source10_.cf0
                d_551_cf0_ = d_550___mcc_h11_
                return (_dafny.Set({(0) - ((pat_let_tv29_)[default__.safeIndex(188, (pat_let_tv28_).length(0))])})).ispropersubset(_dafny.Set({len(pat_let_tv30_)}))

        (globalState).f0 = lambda19_(d_537_v105_)
        r0 = (d_500_v84_)[default__.safeIndex(188, (d_500_v84_).length(0))]
        return r0

    def m2(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        if not((self).f13):
            d_552_v0_: _dafny.Array
            def lambda20_(d_553_p1_):
                def lambda21_(d_554_i0_):
                    return d_553_p1_.f3

                return lambda21_

            init9_ = lambda20_(p1)
            nw87_ = _dafny.Array(None, 22)
            for i0_9_ in range(nw87_.length(0)):
                nw87_[i0_9_] = init9_(i0_9_)
            d_552_v0_ = nw87_
            d_555_v1_: _dafny.MultiSet
            d_555_v1_ = _dafny.MultiSet([(p1).fm0(globalState)])
            d_556_v2_: _dafny.Map
            d_556_v2_ = _dafny.Map({d_555_v1_: p1.f3})
            index79_ = default__.safeIndex(126, (d_552_v0_).length(0))
            (d_552_v0_)[index79_] = ((d_556_v2_)[d_555_v1_] if (d_555_v1_) in (d_556_v2_) else (self).f13)
            d_557_v3_: _dafny.Seq
            d_557_v3_ = _dafny.SeqWithoutIsStrInference([not((p1).f4), True, p1.f3, (self).f13])
            d_558_v4_: _dafny.Seq
            d_558_v4_ = _dafny.SeqWithoutIsStrInference([d_557_v3_, d_557_v3_])
            d_559_v5_: int
            d_559_v5_ = 252
            d_560_v6_: _dafny.Map
            d_560_v6_ = _dafny.Map({(p1).f4: d_559_v5_})
            d_561_v7_: _dafny.Map
            d_561_v7_ = _dafny.Map({False: p1.f3})
            index80_ = default__.safeIndex(126, (d_552_v0_).length(0))
            rhs73_ = (default__.fm26(((d_560_v6_)[p1.f3] if (p1.f3) in (d_560_v6_) else 434), p0, False, globalState)) != (d_561_v7_)
            rhs74_ = (d_559_v5_) != ((p1).fm0(globalState))
            rhs75_ = p1.f3
            rhs76_ = d_558_v4_
            lhs51_ = d_552_v0_
            lhs52_ = default__.safeIndex(126, (d_552_v0_).length(0))
            lhs53_ = p1
            lhs54_ = globalState
            lhs51_[lhs52_] = rhs73_
            lhs53_.f3 = rhs74_
            lhs54_.f0 = rhs75_
            d_558_v4_ = rhs76_
            d_562_v8_: _dafny.Array
            nw88_ = _dafny.Array(int(0), 3)
            d_562_v8_ = nw88_
            d_562_v8_ = (d_562_v8_ if p1.f3 else d_562_v8_)
            d_562_v8_ = d_562_v8_
            d_563_v9_: _dafny.Array
            nw89_ = _dafny.Array(_dafny.Set({}), 26)
            d_563_v9_ = nw89_
            d_564_v10_: _dafny.Set
            d_564_v10_ = _dafny.Set({-462, ((d_555_v1_)[d_559_v5_] if (d_559_v5_) in (d_555_v1_) else d_559_v5_)})
            index81_ = default__.safeIndex(298, (d_563_v9_).length(0))
            (d_563_v9_)[index81_] = d_564_v10_
            d_565_v11_: _dafny.Set
            d_565_v11_ = _dafny.Set({p0})
            index82_ = default__.safeIndex(298, (d_563_v9_).length(0))
            rhs77_ = d_559_v5_
            rhs78_ = d_564_v10_
            rhs79_ = default__.fm27((len(d_560_v6_) if (d_552_v0_)[default__.safeIndex(126, (d_552_v0_).length(0))] else (p1).fm0(globalState)), p1.f3, d_559_v5_, globalState)
            lhs55_ = d_563_v9_
            lhs56_ = default__.safeIndex(298, (d_563_v9_).length(0))
            d_559_v5_ = rhs77_
            lhs55_[lhs56_] = rhs78_
            d_565_v11_ = rhs79_
            d_566_v12_: D0
            d_566_v12_ = D0_DC2((d_552_v0_)[default__.safeIndex(126, (d_552_v0_).length(0))], len(d_560_v6_), d_559_v5_)
            rhs80_ = ((d_566_v12_ if (d_552_v0_)[default__.safeIndex(126, (d_552_v0_).length(0))] else d_566_v12_)).cf5
            rhs81_ = (p1).f4
            lhs57_ = globalState
            d_559_v5_ = rhs80_
            lhs57_.f0 = rhs81_
        elif True:
            d_567_v13_: int
            d_567_v13_ = 500
            d_568_v14_: _dafny.MultiSet
            d_568_v14_ = _dafny.MultiSet([_dafny.Map({d_567_v13_: d_567_v13_})])
            d_569_v15_: _dafny.MultiSet
            d_569_v15_ = _dafny.MultiSet([d_567_v13_, -52])
            d_570_v16_: _dafny.MultiSet
            d_570_v16_ = _dafny.MultiSet([False, not((p1).f4)])
            d_571_v17_: D1
            d_571_v17_ = D1_DC5(p1.f3, not(p1.f3), d_567_v13_, d_567_v13_)
            d_572_v18_: _dafny.Set
            d_572_v18_ = _dafny.Set({True, p1.f3})
            d_573_v19_: _dafny.Array
            nw90_ = _dafny.Array(None, 28)
            nw90_[int(0)] = p1.f3
            nw90_[int(1)] = default__.fm21(d_567_v13_, d_568_v14_, globalState)
            nw90_[int(2)] = p1.f3
            nw90_[int(3)] = True
            nw90_[int(4)] = p1.f3
            nw90_[int(5)] = (self).f13
            nw90_[int(6)] = (p1).f4
            nw90_[int(7)] = (self).f13
            nw90_[int(8)] = p1.f3
            nw90_[int(9)] = p1.f3
            nw90_[int(10)] = (p1).f4
            nw90_[int(11)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhyqkwo"))).set(default__.safeIndex(d_567_v13_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhyqkwo")))), p0)) != ((self).f12)
            nw90_[int(12)] = False
            nw90_[int(13)] = default__.fm21(d_567_v13_, d_568_v14_, globalState)
            nw90_[int(14)] = (d_569_v15_).isdisjoint(d_569_v15_)
            nw90_[int(15)] = (_dafny.MultiSet([(p1).f4, not((p1).f4), (p1).f4])).ispropersubset(d_570_v16_)
            nw90_[int(16)] = p1.f3
            nw90_[int(17)] = (d_571_v17_).cf8
            nw90_[int(18)] = True
            nw90_[int(19)] = (p1).f4
            nw90_[int(20)] = ((p1).f4) == (p1.f3)
            nw90_[int(21)] = p1.f3
            nw90_[int(22)] = not((p1).f4)
            nw90_[int(23)] = (d_572_v18_) != (d_572_v18_)
            nw90_[int(24)] = p1.f3
            nw90_[int(25)] = (self).f13
            nw90_[int(26)] = (self).f13
            nw90_[int(27)] = p1.f3
            d_573_v19_ = nw90_
            d_573_v19_ = d_573_v19_
            d_567_v13_ = default__.safeModuloInt(70, -157)
            d_574_v20_: _dafny.Map
            d_574_v20_ = _dafny.Map({d_567_v13_: (p1).f4})
            d_575_v21_: _dafny.MultiSet
            d_575_v21_ = _dafny.MultiSet([d_574_v20_])
            d_576_v22_: D3
            d_576_v22_ = D3_DC9(d_575_v21_)
            d_567_v13_ = ((d_576_v22_).cf15).cardinality
            d_577_v23_: _dafny.Seq
            d_577_v23_ = _dafny.SeqWithoutIsStrInference([(p1).f4, default__.fm21(d_567_v13_, d_568_v14_, globalState), not((p1).f4)])
            d_578_v24_: _dafny.Map
            d_578_v24_ = _dafny.Map({(d_570_v16_).cardinality: d_577_v23_})
            d_579_v25_: _dafny.Seq
            d_579_v25_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")))])
            d_578_v24_ = (d_578_v24_).set((d_579_v25_)[default__.safeIndex(d_567_v13_, len(d_579_v25_))], d_577_v23_)
            r0 = ((self).f12) + ((self).f12)
        d_580_v26_: int
        d_580_v26_ = 859
        d_581_v27_: _dafny.Map
        d_581_v27_ = _dafny.Map({default__.fm28((p1).f4, (self).f13, p1.f3, d_580_v26_, globalState): True})
        d_582_v28_: _dafny.Set
        d_582_v28_ = _dafny.Set({-226})
        d_583_v29_: _dafny.MultiSet
        d_583_v29_ = _dafny.MultiSet([684, d_580_v26_])
        d_584_v30_: _dafny.MultiSet
        d_584_v30_ = _dafny.MultiSet([p0, default__.fm23(False, 179, globalState), p0, p0, p0])
        d_585_v31_: _dafny.Seq
        d_585_v31_ = _dafny.SeqWithoutIsStrInference([p1.f3])
        d_586_v32_: _dafny.Array
        nw91_ = _dafny.Array(None, 14)
        nw91_[int(0)] = (d_580_v26_) - (len(d_581_v27_))
        nw91_[int(1)] = d_580_v26_
        nw91_[int(2)] = d_580_v26_
        nw91_[int(3)] = 477
        nw91_[int(4)] = d_580_v26_
        nw91_[int(5)] = (0) - ((d_580_v26_) + (len(d_582_v28_)))
        nw91_[int(6)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_587_i1_ in range(default__.abs(733))]))
        nw91_[int(7)] = (len((self).f12)) - (d_580_v26_)
        nw91_[int(8)] = (((d_583_v29_)[d_580_v26_] if (d_580_v26_) in (d_583_v29_) else len((self).f12))) * (d_580_v26_)
        nw91_[int(9)] = ((d_584_v30_)[p0] if (p0) in (d_584_v30_) else d_580_v26_)
        nw91_[int(10)] = (d_583_v29_).cardinality
        nw91_[int(11)] = 972
        nw91_[int(12)] = d_580_v26_
        nw91_[int(13)] = len(d_585_v31_)
        d_586_v32_ = nw91_
        d_588_v33_: D4
        d_588_v33_ = D4_DC11(_dafny.SeqWithoutIsStrInference([(p1).f4]))
        index83_ = default__.safeIndex(829, (d_586_v32_).length(0))
        (d_586_v32_)[index83_] = (0) - (len((d_585_v31_) + ((d_588_v33_).cf18)))
        index84_ = default__.safeIndex(829, (d_586_v32_).length(0))
        (d_586_v32_)[index84_] = (p1).fm0(globalState)
        d_589_v34_: _dafny.MultiSet
        d_589_v34_ = _dafny.MultiSet([d_585_v31_, d_585_v31_])
        hi3_ = default__.safeModuloInt(len((self).f12), ((d_589_v34_)[d_585_v31_] if (d_585_v31_) in (d_589_v34_) else (0) - ((d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))])))
        for d_590_i2_ in range(472, hi3_):
            (globalState).f0 = (p1).f4
            d_591_v35_: _dafny.Map
            d_591_v35_ = _dafny.Map({866: d_590_i2_})
            d_592_v36_: _dafny.MultiSet
            d_592_v36_ = _dafny.MultiSet([_dafny.Map({(0) - ((d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))]): (d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))]}), d_591_v35_, d_591_v35_])
            d_593_v37_: _dafny.MultiSet
            d_593_v37_ = _dafny.MultiSet([p1.f3, default__.fm21(d_590_i2_, d_592_v36_, globalState), (p1).f4])
            d_594_v38_: _dafny.Seq
            d_594_v38_ = _dafny.SeqWithoutIsStrInference([d_593_v37_])
            d_595_v39_: _dafny.Array
            nw92_ = _dafny.Array(None, 26)
            nw92_[int(0)] = d_593_v37_
            nw92_[int(1)] = d_593_v37_
            nw92_[int(2)] = _dafny.MultiSet([p1.f3])
            nw92_[int(3)] = _dafny.MultiSet([(p1).f4, p1.f3, p1.f3, (self).f13, p1.f3])
            nw92_[int(4)] = (d_593_v37_) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p1.f3, p1.f3, p1.f3, p1.f3])))
            nw92_[int(5)] = _dafny.MultiSet([p1.f3, False])
            nw92_[int(6)] = d_593_v37_
            nw92_[int(7)] = (d_594_v38_)[default__.safeIndex((0) - ((d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))]), len(d_594_v38_))]
            nw92_[int(8)] = d_593_v37_
            nw92_[int(9)] = d_593_v37_
            nw92_[int(10)] = (d_593_v37_) - (d_593_v37_)
            nw92_[int(11)] = d_593_v37_
            nw92_[int(12)] = (d_593_v37_) | (d_593_v37_)
            nw92_[int(13)] = d_593_v37_
            nw92_[int(14)] = d_593_v37_
            nw92_[int(15)] = (_dafny.MultiSet([p1.f3, (p1).f4])).intersection(d_593_v37_)
            nw92_[int(16)] = d_593_v37_
            nw92_[int(17)] = d_593_v37_
            nw92_[int(18)] = d_593_v37_
            nw92_[int(19)] = d_593_v37_
            nw92_[int(20)] = _dafny.MultiSet([(p1).f4, (p1).f4, (p1).f4, p1.f3])
            nw92_[int(21)] = d_593_v37_
            nw92_[int(22)] = d_593_v37_
            nw92_[int(23)] = d_593_v37_
            nw92_[int(24)] = _dafny.MultiSet([(p1).f4])
            nw92_[int(25)] = d_593_v37_
            d_595_v39_ = nw92_
            d_596_v40_: _dafny.Seq
            d_596_v40_ = _dafny.SeqWithoutIsStrInference([d_595_v39_])
            d_597_v41_: _dafny.Array
            def lambda22_(d_598_p1_):
                def lambda23_(d_599_i3_):
                    return (d_598_p1_).f4

                return lambda23_

            init10_ = lambda22_(p1)
            nw93_ = _dafny.Array(None, 17)
            for i0_10_ in range(nw93_.length(0)):
                nw93_[i0_10_] = init10_(i0_10_)
            d_597_v41_ = nw93_
            d_600_v42_: _dafny.Set
            d_600_v42_ = _dafny.Set({d_597_v41_, d_597_v41_})
            d_601_v43_: _dafny.Map
            d_601_v43_ = _dafny.Map({(p1).f4: d_600_v42_})
            d_602_v44_: _dafny.Seq
            d_602_v44_ = _dafny.SeqWithoutIsStrInference([d_600_v42_, d_600_v42_, ((d_601_v43_)[False] if (False) in (d_601_v43_) else d_600_v42_)])
            d_595_v39_ = (d_596_v40_)[default__.safeIndex((0) - ((len((d_602_v44_)[default__.safeIndex(d_590_i2_, len(d_602_v44_))])) - (d_580_v26_)), len(d_596_v40_))]
            (globalState).f0 = not(not(((d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))]) == ((d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))])))
            index85_ = default__.safeIndex(829, (d_586_v32_).length(0))
            (d_586_v32_)[index85_] = ((d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))]) + ((p1).fm0(globalState))
        d_603_v45_: _dafny.Seq
        d_603_v45_ = _dafny.SeqWithoutIsStrInference([(d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))]])
        d_604_v46_: _dafny.Seq
        d_604_v46_ = _dafny.SeqWithoutIsStrInference([d_603_v45_, d_603_v45_])
        d_605_v47_: _dafny.Seq
        d_605_v47_ = _dafny.SeqWithoutIsStrInference([d_604_v46_, _dafny.SeqWithoutIsStrInference([d_603_v45_ for d_606_i7_ in range(default__.abs(-127))]), d_604_v46_, d_604_v46_, d_604_v46_])
        d_607_v48_: _dafny.Array
        nw94_ = _dafny.Array(None, 21)
        nw94_[int(0)] = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([p1.f3, p1.f3])).cardinality]) for d_608_i4_ in range(default__.abs(30))])
        nw94_[int(1)] = d_604_v46_
        nw94_[int(2)] = d_604_v46_
        nw94_[int(3)] = d_604_v46_
        nw94_[int(4)] = (_dafny.SeqWithoutIsStrInference([d_603_v45_ for d_609_i5_ in range(default__.abs(856))])) + (_dafny.SeqWithoutIsStrInference([d_603_v45_ for d_610_i6_ in range(default__.abs(875))]))
        nw94_[int(5)] = default__.fm29(d_580_v26_, globalState)
        nw94_[int(6)] = ((d_604_v46_) + ((d_605_v47_)[default__.safeIndex(893, len(d_605_v47_))])).set(default__.safeIndex((d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))], len((d_604_v46_) + ((d_605_v47_)[default__.safeIndex(893, len(d_605_v47_))]))), d_603_v45_)
        nw94_[int(7)] = d_604_v46_
        nw94_[int(8)] = _dafny.SeqWithoutIsStrInference([d_603_v45_ for d_611_i8_ in range(default__.abs(-488))])
        nw94_[int(9)] = d_604_v46_
        nw94_[int(10)] = d_604_v46_
        nw94_[int(11)] = (d_604_v46_ if (p1).f4 else default__.fm29((0) - (d_580_v26_), globalState))
        nw94_[int(12)] = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_580_v26_])])
        nw94_[int(13)] = d_604_v46_
        nw94_[int(14)] = (d_604_v46_ if (p1).f4 else d_604_v46_)
        nw94_[int(15)] = ((d_604_v46_).set(default__.safeIndex((0) - (len(_dafny.Set({p1.f3}))), len(d_604_v46_)), d_603_v45_)) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_580_v26_, (d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))]]), d_603_v45_, _dafny.SeqWithoutIsStrInference([282 for d_612_i9_ in range(default__.abs(997))])])).set(default__.safeIndex(len((self).f12), len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_580_v26_, (d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))]]), d_603_v45_, _dafny.SeqWithoutIsStrInference([282 for d_613_i9_ in range(default__.abs(997))])]))), d_603_v45_))
        nw94_[int(16)] = (d_604_v46_) + ((d_604_v46_).set(default__.safeIndex((0) - (len(((self).f12).set(default__.safeIndex(d_580_v26_, len((self).f12)), default__.fm23(p1.f3, (d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))], globalState)))), len(d_604_v46_)), d_603_v45_))
        nw94_[int(17)] = d_604_v46_
        nw94_[int(18)] = d_604_v46_
        nw94_[int(19)] = d_604_v46_
        nw94_[int(20)] = d_604_v46_
        d_607_v48_ = nw94_
        index86_ = default__.safeIndex(594, (d_607_v48_).length(0))
        (d_607_v48_)[index86_] = d_604_v46_
        index87_ = default__.safeIndex(594, (d_607_v48_).length(0))
        (d_607_v48_)[index87_] = (d_604_v46_) + (d_604_v46_)
        d_614_v49_: _dafny.Map
        d_614_v49_ = _dafny.Map({len(d_603_v45_): d_580_v26_})
        d_615_v50_: _dafny.MultiSet
        d_615_v50_ = _dafny.MultiSet([default__.fm25(globalState), _dafny.Map({(d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))]: len(d_603_v45_)}), d_614_v49_, d_614_v49_, _dafny.Map({d_580_v26_: d_580_v26_})])
        (globalState).f0 = default__.fm21(280, (d_615_v50_) - (d_615_v50_), globalState)
        _ingredientsd_0 = []
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_586_v32_).length(0)):
            d_616_i10_: int = guard_loop_0_
            if (True) and (((0) <= (d_616_i10_)) and ((d_616_i10_) < ((d_586_v32_).length(0)))):
                _ingredientsd_0.append((d_586_v32_, int(d_616_i10_), (d_616_i10_) * ((d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))])))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        d_617_v51_: D1
        d_617_v51_ = D1_DC5(p1.f3, (self).f13, len((self).f12), (d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))])
        r0 = (default__.fm30(d_617_v51_, (0) - ((d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))]), (d_586_v32_)[default__.safeIndex(829, (d_586_v32_).length(0))], globalState)) + ((self).f12)
        return r0

    def m10(self, globalState):
        d_618_v0_: int
        d_618_v0_ = 391
        d_618_v0_ = d_618_v0_
        d_619_v1_: _dafny.Map
        d_619_v1_ = _dafny.Map({(self).f13: d_618_v0_})
        d_620_v2_: _dafny.Set
        d_620_v2_ = _dafny.Set({(self).f13})
        d_621_v3_: _dafny.Map
        out19_: _dafny.Map
        out19_ = (self).m11((self).f13, len(d_619_v1_), (d_620_v2_) - (d_620_v2_), globalState)
        d_621_v3_ = out19_
        d_622_v4_: _dafny.Seq
        d_622_v4_ = _dafny.SeqWithoutIsStrInference([(self).f13])
        d_623_v5_: _dafny.Array
        nw95_ = _dafny.Array(None, 11)
        nw95_[int(0)] = d_618_v0_
        nw95_[int(1)] = len((self).f12)
        nw95_[int(2)] = len((self).f12)
        nw95_[int(3)] = 204
        nw95_[int(4)] = d_618_v0_
        nw95_[int(5)] = d_618_v0_
        nw95_[int(6)] = d_618_v0_
        nw95_[int(7)] = d_618_v0_
        nw95_[int(8)] = d_618_v0_
        nw95_[int(9)] = d_618_v0_
        nw95_[int(10)] = default__.fm31((self).f13, d_618_v0_, (self).f13, d_622_v4_, globalState)
        d_623_v5_ = nw95_
        d_624_v6_: _dafny.Seq
        d_624_v6_ = _dafny.SeqWithoutIsStrInference([d_623_v5_, d_623_v5_, d_623_v5_])
        d_625_v7_: _dafny.Seq
        d_625_v7_ = d_624_v6_
        d_626_v8_: _dafny.Array
        nw96_ = _dafny.Array(None, 14)
        nw96_[int(0)] = (((d_624_v6_).set(default__.safeIndex(590, len(d_624_v6_)), d_623_v5_)).set(default__.safeIndex(d_618_v0_, len((d_624_v6_).set(default__.safeIndex(590, len(d_624_v6_)), d_623_v5_))), d_623_v5_)) + ((_dafny.SeqWithoutIsStrInference([d_623_v5_])).set(default__.safeIndex(d_618_v0_, len(_dafny.SeqWithoutIsStrInference([d_623_v5_]))), d_623_v5_))
        nw96_[int(1)] = (d_625_v7_)
        nw96_[int(2)] = d_624_v6_
        nw96_[int(3)] = d_624_v6_
        nw96_[int(4)] = d_624_v6_
        nw96_[int(5)] = (d_624_v6_) + (d_624_v6_)
        nw96_[int(6)] = d_624_v6_
        nw96_[int(7)] = d_624_v6_
        nw96_[int(8)] = d_624_v6_
        nw96_[int(9)] = d_624_v6_
        nw96_[int(10)] = (d_624_v6_) + (d_624_v6_)
        nw96_[int(11)] = _dafny.SeqWithoutIsStrInference([d_623_v5_])
        nw96_[int(12)] = (d_624_v6_) + ((_dafny.SeqWithoutIsStrInference([d_623_v5_])).set(default__.safeIndex(d_618_v0_, len(_dafny.SeqWithoutIsStrInference([d_623_v5_]))), d_623_v5_))
        nw96_[int(13)] = d_624_v6_
        d_626_v8_ = nw96_
        index88_ = default__.safeIndex(387, (d_626_v8_).length(0))
        (d_626_v8_)[index88_] = (d_624_v6_) + (d_624_v6_)
        index89_ = default__.safeIndex(387, (d_626_v8_).length(0))
        (d_626_v8_)[index89_] = d_624_v6_
        d_618_v0_ = (0) - (default__.safeModuloInt(d_618_v0_, d_618_v0_))
        d_627_v9_: _dafny.Seq
        d_627_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ssfq"))
        d_627_v9_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqxt"))) + (((self).f12 if True else d_627_v9_))
        (globalState).f0 = (self).f13

    def m11(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_628_v0_: _dafny.Array
        def lambda24_(d_629_p0_):
            def lambda25_(d_630_i0_):
                return d_629_p0_

            return lambda25_

        init11_ = lambda24_(p0)
        nw97_ = _dafny.Array(None, 5)
        for i0_11_ in range(nw97_.length(0)):
            nw97_[i0_11_] = init11_(i0_11_)
        d_628_v0_ = nw97_
        d_628_v0_ = d_628_v0_
        d_631_v1_: _dafny.Array
        nw98_ = _dafny.Array(int(0), 28)
        d_631_v1_ = nw98_
        d_631_v1_ = d_631_v1_
        d_632_i1_: int
        d_632_i1_ = 0
        with _dafny.label("2"):
            while p0:
                with _dafny.c_label("2"):
                    if (d_632_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_632_i1_ = (d_632_i1_) + (1)
                    d_633_v2_: str
                    d_633_v2_ = _dafny.CodePoint('t')
                    d_634_v3_: _dafny.MultiSet
                    d_634_v3_ = _dafny.MultiSet([_dafny.CodePoint('t'), d_633_v2_])
                    d_634_v3_ = d_634_v3_
                    d_635_v4_: int
                    d_635_v4_ = -429
                    d_635_v4_ = p1
                    d_636_v5_: _dafny.Seq
                    d_636_v5_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_633_v2_ for d_637_i2_ in range(default__.abs(253))])) + ((self).f12), ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbd"))).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbd")))), d_633_v2_)) + ((self).f12), (self).f12, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkugwmyv"))) + ((self).f12)])
                    d_636_v5_ = (d_636_v5_) + (d_636_v5_)
                    def lambda26_(d_638_v4_):
                        def lambda27_(d_639_i3_):
                            return (d_638_v4_) == (d_638_v4_)

                        return lambda27_

                    init12_ = lambda26_(d_635_v4_)
                    nw99_ = _dafny.Array(None, 11)
                    for i0_12_ in range(nw99_.length(0)):
                        nw99_[i0_12_] = init12_(i0_12_)
                    d_628_v0_ = nw99_
                    pass
            pass
        d_640_v6_: _dafny.Seq
        d_640_v6_ = _dafny.SeqWithoutIsStrInference([p1])
        (globalState).f0 = not((_dafny.SeqWithoutIsStrInference([p1, p1, p1, p1, p1])) <= ((d_640_v6_) + (d_640_v6_)))
        d_641_v7_: _dafny.Map
        d_641_v7_ = _dafny.Map({p1: p1})
        d_642_v8_: _dafny.Map
        d_642_v8_ = _dafny.Map({p0: p1})
        d_643_v9_: _dafny.Map
        d_643_v9_ = _dafny.Map({len(d_642_v8_): d_641_v7_})
        d_641_v7_ = (((d_643_v9_)[p1] if (p1) in (d_643_v9_) else d_641_v7_)).set(p1, p1)
        d_644_v10_: _dafny.Seq
        d_644_v10_ = _dafny.SeqWithoutIsStrInference([(self).f12])
        d_645_v11_: _dafny.Seq
        d_645_v11_ = d_644_v10_
        d_646_v12_: D0
        d_646_v12_ = D0_DC2((self).f13, len((d_645_v11_)), p1)
        pat_let_tv31_ = p0
        def lambda28_(source11_):
            if source11_.is_DC1:
                d_647___mcc_h6_ = source11_.cf1
                d_648___mcc_h7_ = source11_.cf2
                d_649_cf2_ = d_648___mcc_h7_
                d_650_cf1_ = d_647___mcc_h6_
                return d_650_cf1_
            elif source11_.is_DC2:
                d_651___mcc_h8_ = source11_.cf3
                d_652___mcc_h9_ = source11_.cf4
                d_653___mcc_h10_ = source11_.cf5
                d_654_cf5_ = d_653___mcc_h10_
                d_655_cf4_ = d_652___mcc_h9_
                d_656_cf3_ = d_651___mcc_h8_
                return (self).f13
            elif True:
                d_657___mcc_h11_ = source11_.cf0
                d_658_cf0_ = d_657___mcc_h11_
                return pat_let_tv31_

        if lambda28_(d_646_v12_):
            d_659_v13_: D1
            d_659_v13_ = D1_DC3(d_628_v0_)
            source12_ = d_659_v13_
            if source12_.is_DC4:
                d_660___mcc_h0_ = source12_.cf7
                d_661_cf7_ = d_660___mcc_h0_
                d_662_v14_: _dafny.Set
                d_662_v14_ = _dafny.Set({p1})
                d_663_v15_: _dafny.Seq
                d_663_v15_ = _dafny.SeqWithoutIsStrInference([(self).f13])
                d_664_v16_: _dafny.Map
                d_664_v16_ = _dafny.Map({len((self).f12): (self).f12})
                d_665_v17_: _dafny.Seq
                d_665_v17_ = _dafny.SeqWithoutIsStrInference([default__.fm32(True, False, p0, (self).f13, globalState), d_662_v14_, _dafny.Set({d_661_cf7_}), _dafny.Set({d_661_cf7_, (0) - (len(d_663_v15_))}), _dafny.Set({p1, len(d_664_v16_)})])
                d_666_v19_: _dafny.MultiSet
                d_666_v19_ = _dafny.MultiSet([default__.fm25(globalState)])
                def iife49_():
                    coll27_ = _dafny.Set()
                    compr_27_: int
                    for compr_27_ in (_dafny.MultiSet([p1])).Elements:
                        d_667_v18_: int = compr_27_
                        if (d_667_v18_) in (_dafny.MultiSet([p1])):
                            coll27_ = coll27_.union(_dafny.Set([(d_667_v18_) - (762)]))
                    return _dafny.Set(coll27_)
                (globalState).f0 = default__.fm21(len((d_665_v17_) + (_dafny.SeqWithoutIsStrInference([iife49_()
 for d_668_i4_ in range(default__.abs(251))]))), d_666_v19_, globalState)
                (globalState).f0 = (d_663_v15_)[default__.safeIndex(d_661_cf7_, len(d_663_v15_))]
                d_669_v20_: C0
                nw100_ = C0()
                nw100_.ctor__()
                d_669_v20_ = nw100_
                d_670_v21_: _dafny.Map
                d_670_v21_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_671_i5_ in range(default__.abs(-805))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcygaqfhi"))): d_642_v8_})
                d_672_v22_: _dafny.Array
                nw101_ = _dafny.Array(_dafny.MultiSet({}), 20)
                d_672_v22_ = nw101_
                d_673_v23_: _dafny.MultiSet
                d_673_v23_ = _dafny.MultiSet([d_662_v14_])
                index90_ = default__.safeIndex(716, (d_672_v22_).length(0))
                (d_672_v22_)[index90_] = (d_673_v23_) | (d_673_v23_)
                d_674_v25_: _dafny.Set
                d_674_v25_ = _dafny.Set({(self).f12})
                d_675_v28_: _dafny.Map
                def iife50_():
                    coll28_ = _dafny.Set()
                    compr_28_: int
                    for compr_28_ in _dafny.IntegerRange(-500, 373):
                        d_676_v27_: int = compr_28_
                        if ((-500) <= (d_676_v27_)) and ((d_676_v27_) < (373)):
                            coll28_ = coll28_.union(_dafny.Set([(d_676_v27_) + (len(_dafny.Set({(self).f13})))]))
                    return _dafny.Set(coll28_)
                d_675_v28_ = _dafny.Map({d_645_v11_: iife50_()
                })
                index91_ = default__.safeIndex(716, (d_672_v22_).length(0))
                def iife51_():
                    coll29_ = _dafny.Map()
                    compr_29_: _dafny.Seq
                    for compr_29_ in (d_674_v25_).Elements:
                        d_677_v24_: _dafny.Seq = compr_29_
                        if (d_677_v24_) in (d_674_v25_):
                            coll29_[d_677_v24_] = d_642_v8_
                    return _dafny.Map(coll29_)
                def iife52_():
                    coll30_ = _dafny.Map()
                    compr_30_: _dafny.Seq
                    for compr_30_ in (d_644_v10_).Elements:
                        d_678_v26_: _dafny.Seq = compr_30_
                        if (d_678_v26_) in (d_644_v10_):
                            coll30_[d_678_v26_] = d_642_v8_
                    return _dafny.Map(coll30_)
                rhs82_ = (iife51_()
                ) | ((iife52_()
                ).set((self).f12, d_642_v8_))
                rhs83_ = (p1) >= ((0) - (default__.safeModuloInt(d_661_cf7_, -787)))
                rhs84_ = d_661_cf7_
                rhs85_ = _dafny.MultiSet([((d_675_v28_)[d_645_v11_] if (d_645_v11_) in (d_675_v28_) else d_662_v14_)])
                lhs58_ = globalState
                lhs59_ = d_672_v22_
                lhs60_ = default__.safeIndex(716, (d_672_v22_).length(0))
                d_670_v21_ = rhs82_
                lhs58_.f0 = rhs83_
                d_661_cf7_ = rhs84_
                lhs59_[lhs60_] = rhs85_
            elif source12_.is_DC5:
                d_679___mcc_h1_ = source12_.cf8
                d_680___mcc_h2_ = source12_.cf9
                d_681___mcc_h3_ = source12_.cf10
                d_682___mcc_h4_ = source12_.cf11
                d_683_cf11_ = d_682___mcc_h4_
                d_684_cf10_ = d_681___mcc_h3_
                d_685_cf9_ = d_680___mcc_h2_
                d_686_cf8_ = d_679___mcc_h1_
                d_687_v29_: _dafny.Map
                d_687_v29_ = _dafny.Map({(True if d_685_cf9_ else (self).f13): d_686_cf8_})
                d_687_v29_ = (d_687_v29_).set((d_683_cf11_) in (d_641_v7_), False)
                d_688_v30_: _dafny.MultiSet
                d_688_v30_ = _dafny.MultiSet([d_641_v7_, d_641_v7_])
                index92_ = default__.safeIndex(220, (d_628_v0_).length(0))
                (d_628_v0_)[index92_] = default__.fm21(d_683_cf11_, d_688_v30_, globalState)
                index93_ = default__.safeIndex(220, (d_628_v0_).length(0))
                (d_628_v0_)[index93_] = d_686_cf8_
                d_689_v31_: D1
                d_689_v31_ = D1_DC5(d_685_cf9_, (d_628_v0_)[default__.safeIndex(220, (d_628_v0_).length(0))], d_683_cf11_, (0) - ((d_640_v6_)[default__.safeIndex(d_684_cf10_, len(d_640_v6_))]))
                d_690_v32_: _dafny.Map
                d_690_v32_ = _dafny.Map({d_689_v31_: p1})
                d_690_v32_ = d_690_v32_
                d_684_cf10_ = d_684_cf10_
            elif source12_.is_DC6:
                index94_ = default__.safeIndex(313, (d_631_v1_).length(0))
                (d_631_v1_)[index94_] = p1
                index95_ = default__.safeIndex(313, (d_631_v1_).length(0))
                (d_631_v1_)[index95_] = default__.safeModuloInt(p1, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsas"))) + ((self).f12)))
                d_691_v33_: _dafny.Array
                nw102_ = _dafny.Array(_dafny.Seq({}), 18)
                d_691_v33_ = nw102_
                d_691_v33_ = d_691_v33_
                d_692_v34_: _dafny.Map
                d_692_v34_ = _dafny.Map({p0: (self).f12})
                d_693_v36_: D7
                def iife53_():
                    coll31_ = _dafny.Map()
                    compr_31_: int
                    for compr_31_ in _dafny.IntegerRange(483, 839):
                        d_694_v35_: int = compr_31_
                        if ((483) <= (d_694_v35_)) and ((d_694_v35_) < (839)):
                            coll31_[default__.safeDivisionInt(d_694_v35_, len(_dafny.SeqWithoutIsStrInference([p0])))] = p2
                    return _dafny.Map(coll31_)
                d_693_v36_ = D7_DC16(iife53_()
)
                d_695_v37_: _dafny.Map
                d_695_v37_ = _dafny.Map({((d_692_v34_)[not((self).f13)] if (not((self).f13)) in (d_692_v34_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bchjxx"))): (d_693_v36_).cf25})
                d_696_v38_: D1
                d_696_v38_ = D1_DC5(True, p0, len((self).f12), (d_631_v1_)[default__.safeIndex(313, (d_631_v1_).length(0))])
                d_697_v39_: _dafny.MultiSet
                d_697_v39_ = _dafny.MultiSet([(d_631_v1_)[default__.safeIndex(313, (d_631_v1_).length(0))]])
                d_698_v40_: _dafny.Map
                d_698_v40_ = _dafny.Map({p1: p2})
                pat_let_tv32_ = p0
                def iife54_(_pat_let11_0):
                    def iife55_(d_699_dt__update__tmp_h0_):
                        def iife56_(_pat_let12_0):
                            def iife57_(d_700_dt__update_hcf9_h0_):
                                def iife58_(_pat_let13_0):
                                    def iife59_(d_701_dt__update_hcf8_h0_):
                                        return D1_DC5(d_701_dt__update_hcf8_h0_, d_700_dt__update_hcf9_h0_, (d_699_dt__update__tmp_h0_).cf10, (d_699_dt__update__tmp_h0_).cf11)
                                    return iife59_(_pat_let13_0)
                                return iife58_(True)
                            return iife57_(_pat_let12_0)
                        return iife56_(pat_let_tv32_)
                    return iife55_(_pat_let11_0)
                d_695_v37_ = (d_695_v37_).set(default__.fm30(iife54_(d_696_v38_), (d_631_v1_)[default__.safeIndex(313, (d_631_v1_).length(0))], (d_697_v39_).cardinality, globalState), d_698_v40_)
                d_702_v41_: D0
                d_702_v41_ = D0_DC0((d_631_v1_)[default__.safeIndex(313, (d_631_v1_).length(0))])
                (globalState).f0 = (_dafny.MultiSet([(0) - ((d_631_v1_)[default__.safeIndex(313, (d_631_v1_).length(0))]), p1, (d_631_v1_)[default__.safeIndex(313, (d_631_v1_).length(0))], (d_631_v1_)[default__.safeIndex(313, (d_631_v1_).length(0))], (d_702_v41_).cf0])).issubset(((d_697_v39_).set(211, default__.abs(178))).intersection(d_697_v39_))
            elif True:
                d_703___mcc_h5_ = source12_.cf6
                d_704_cf6_ = d_703___mcc_h5_
                d_705_v42_: int
                d_705_v42_ = -940
                d_706_v43_: _dafny.Array
                nw103_ = _dafny.Array(_dafny.CodePoint('D'), 21)
                d_706_v43_ = nw103_
                d_707_v44_: str
                d_707_v44_ = _dafny.CodePoint('f')
                index96_ = default__.safeIndex(517, (d_706_v43_).length(0))
                (d_706_v43_)[index96_] = d_707_v44_
                d_708_v45_: _dafny.MultiSet
                d_708_v45_ = _dafny.MultiSet([D7_DC17(p1, p0, p0, p1)])
                index97_ = default__.safeIndex(517, (d_706_v43_).length(0))
                rhs86_ = len((((self).f12) + ((self).f12)) + ((self).f12))
                rhs87_ = ((self).f12)[default__.safeIndex(d_705_v42_, len((self).f12))]
                rhs88_ = not((d_708_v45_).issubset(d_708_v45_))
                rhs89_ = d_628_v0_
                lhs61_ = d_706_v43_
                lhs62_ = default__.safeIndex(517, (d_706_v43_).length(0))
                lhs63_ = globalState
                d_705_v42_ = rhs86_
                lhs61_[lhs62_] = rhs87_
                lhs63_.f0 = rhs88_
                d_704_cf6_ = rhs89_
                (globalState).f0 = p0
                d_706_v43_ = d_706_v43_
                d_709_v46_: _dafny.MultiSet
                d_709_v46_ = _dafny.MultiSet([d_631_v1_])
                d_710_v47_: _dafny.MultiSet
                d_710_v47_ = _dafny.MultiSet([p0, p0])
                d_711_v48_: _dafny.MultiSet
                d_711_v48_ = _dafny.MultiSet([((d_710_v47_) | (d_710_v47_)).cardinality])
                d_712_v49_: _dafny.Seq
                d_712_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnrvw"))
                rhs90_ = (d_709_v46_) | ((_dafny.MultiSet([d_631_v1_])) | (_dafny.MultiSet([d_631_v1_])))
                rhs91_ = default__.fm24(globalState)
                rhs92_ = p1
                rhs93_ = (d_712_v49_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwbrdrf")))
                rhs94_ = p1
                d_709_v46_ = rhs90_
                d_711_v48_ = rhs91_
                d_705_v42_ = rhs92_
                d_712_v49_ = rhs93_
                d_705_v42_ = rhs94_
            d_713_v50_: int
            d_713_v50_ = -678
            d_714_v51_: _dafny.Seq
            d_714_v51_ = _dafny.SeqWithoutIsStrInference([(self).f13, p0])
            d_713_v50_ = (D7_DC17(p1, p0, False, default__.fm31(p0, d_713_v50_, (self).f13, d_714_v51_, globalState))).cf26
            index98_ = default__.safeIndex(486, (d_628_v0_).length(0))
            (d_628_v0_)[index98_] = (self).f13
            index99_ = default__.safeIndex(486, (d_628_v0_).length(0))
            (d_628_v0_)[index99_] = (self).f13
            index100_ = default__.safeIndex(486, (d_628_v0_).length(0))
            (d_628_v0_)[index100_] = (self).f13
            if (p1) != (p1):
                index101_ = default__.safeIndex(486, (d_628_v0_).length(0))
                (d_628_v0_)[index101_] = (d_628_v0_)[default__.safeIndex(486, (d_628_v0_).length(0))]
                d_715_v52_: _dafny.MultiSet
                d_715_v52_ = _dafny.MultiSet([(self).f13])
                d_716_v53_: _dafny.Seq
                d_716_v53_ = _dafny.SeqWithoutIsStrInference([d_715_v52_])
                (globalState).f0 = ((d_716_v53_)[default__.safeIndex(d_713_v50_, len(d_716_v53_))]).isdisjoint(d_715_v52_)
                index102_ = default__.safeIndex(492, (d_631_v1_).length(0))
                (d_631_v1_)[index102_] = 709
                index103_ = default__.safeIndex(896, (d_631_v1_).length(0))
                (d_631_v1_)[index103_] = d_713_v50_
                index104_ = default__.safeIndex(855, (d_631_v1_).length(0))
                (d_631_v1_)[index104_] = d_713_v50_
                d_717_v54_: _dafny.Set
                d_717_v54_ = _dafny.Set({(d_628_v0_)[default__.safeIndex(486, (d_628_v0_).length(0))]})
                d_718_v55_: D1
                d_718_v55_ = D1_DC5(False, p0, d_713_v50_, p1)
                d_719_v56_: D0
                d_719_v56_ = D0_DC0(p1)
                d_720_v57_: _dafny.Map
                d_720_v57_ = _dafny.Map({(self).f13: d_642_v8_})
                d_721_v58_: _dafny.Seq
                d_721_v58_ = _dafny.SeqWithoutIsStrInference([d_720_v57_, d_720_v57_])
                index105_ = default__.safeIndex(492, (d_631_v1_).length(0))
                index106_ = default__.safeIndex(896, (d_631_v1_).length(0))
                index107_ = default__.safeIndex(855, (d_631_v1_).length(0))
                rhs95_ = default__.safeDivisionInt(p1, default__.safeModuloInt(len(_dafny.Map({(d_628_v0_)[default__.safeIndex(486, (d_628_v0_).length(0))]: 152})), d_713_v50_))
                rhs96_ = (default__.fm30(d_718_v55_, (d_719_v56_).cf0, d_713_v50_, globalState)) == (((self).f12) + (default__.fm30(d_718_v55_, p1, p1, globalState)))
                rhs97_ = d_713_v50_
                rhs98_ = len((((d_721_v58_)[default__.safeIndex(len(d_641_v7_), len(d_721_v58_))]) | (d_720_v57_) if (d_714_v51_) <= (d_714_v51_) else d_720_v57_))
                rhs99_ = (p2) - (p2)
                lhs64_ = d_631_v1_
                lhs65_ = default__.safeIndex(492, (d_631_v1_).length(0))
                lhs66_ = globalState
                lhs67_ = d_631_v1_
                lhs68_ = default__.safeIndex(896, (d_631_v1_).length(0))
                lhs69_ = d_631_v1_
                lhs70_ = default__.safeIndex(855, (d_631_v1_).length(0))
                lhs64_[lhs65_] = rhs95_
                lhs66_.f0 = rhs96_
                lhs67_[lhs68_] = rhs97_
                lhs69_[lhs70_] = rhs98_
                d_717_v54_ = rhs99_
                d_722_v59_: _dafny.MultiSet
                d_722_v59_ = _dafny.MultiSet([d_713_v50_])
                d_723_v60_: _dafny.MultiSet
                d_723_v60_ = _dafny.MultiSet([d_722_v59_])
                d_724_v61_: _dafny.Map
                d_724_v61_ = _dafny.Map({p1: d_723_v60_})
                d_725_v62_: _dafny.Seq
                d_725_v62_ = _dafny.SeqWithoutIsStrInference([d_722_v59_])
                rhs100_ = 108
                rhs101_ = _dafny.SeqWithoutIsStrInference([len((d_642_v8_) | (d_642_v8_))])
                rhs102_ = ((d_724_v61_)[(d_631_v1_)[default__.safeIndex(492, (d_631_v1_).length(0))]] if ((d_631_v1_)[default__.safeIndex(492, (d_631_v1_).length(0))]) in (d_724_v61_) else _dafny.MultiSet(d_725_v62_))
                d_713_v50_ = rhs100_
                d_640_v6_ = rhs101_
                d_723_v60_ = rhs102_
                d_726_v63_: _dafny.Array
                def lambda29_(d_727_v1_):
                    def lambda30_(d_728_i6_):
                        return (d_728_i6_) - ((d_727_v1_)[default__.safeIndex(492, (d_727_v1_).length(0))])

                    return lambda30_

                init13_ = lambda29_(d_631_v1_)
                nw104_ = _dafny.Array(None, 11)
                for i0_13_ in range(nw104_.length(0)):
                    nw104_[i0_13_] = init13_(i0_13_)
                d_726_v63_ = nw104_
                d_726_v63_ = (d_726_v63_ if (self).f13 else d_726_v63_)
            elif True:
                d_729_v64_: _dafny.Map
                d_729_v64_ = _dafny.Map({d_713_v50_: not(p0)})
                d_730_v66_: _dafny.Set
                d_730_v66_ = _dafny.Set({p1})
                d_731_v67_: _dafny.MultiSet
                def iife60_():
                    coll32_ = _dafny.Map()
                    compr_32_: int
                    for compr_32_ in (d_730_v66_).Elements:
                        d_732_v65_: int = compr_32_
                        if (d_732_v65_) in (d_730_v66_):
                            coll32_[default__.safeDivisionInt(d_732_v65_, p1)] = p0
                    return _dafny.Map(coll32_)
                d_731_v67_ = _dafny.MultiSet([d_729_v64_, iife60_()
                ])
                d_733_v68_: D3
                d_733_v68_ = D3_DC9(d_731_v67_)
                d_733_v68_ = d_733_v68_
                d_734_v69_: str
                d_734_v69_ = _dafny.CodePoint('x')
                d_735_v70_: _dafny.Array
                def lambda31_(d_736_v69_):
                    def lambda32_(d_737_i7_):
                        return d_736_v69_

                    return lambda32_

                init14_ = lambda31_(d_734_v69_)
                nw105_ = _dafny.Array(None, 1)
                for i0_14_ in range(nw105_.length(0)):
                    nw105_[i0_14_] = init14_(i0_14_)
                d_735_v70_ = nw105_
                d_738_v71_: _dafny.Set
                d_738_v71_ = _dafny.Set({d_735_v70_})
                d_734_v69_ = default__.fm23((d_738_v71_).isdisjoint(d_738_v71_), len(d_640_v6_), globalState)
                d_739_v72_: C0
                nw106_ = C0()
                nw106_.ctor__()
                d_739_v72_ = nw106_
                index108_ = default__.safeIndex(212, (d_631_v1_).length(0))
                (d_631_v1_)[index108_] = d_713_v50_
                d_740_v73_: _dafny.Map
                d_740_v73_ = _dafny.Map({_dafny.CodePoint('h'): d_734_v69_})
                d_741_v74_: _dafny.Map
                d_741_v74_ = _dafny.Map({d_713_v50_: d_740_v73_})
                d_742_v75_: _dafny.MultiSet
                d_742_v75_ = _dafny.MultiSet([((d_741_v74_)[899] if (899) in (d_741_v74_) else d_740_v73_)])
                d_743_v76_: _dafny.Seq
                d_743_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xslucuycd"))
                index109_ = default__.safeIndex(212, (d_631_v1_).length(0))
                rhs103_ = d_713_v50_
                rhs104_ = default__.safeDivisionInt(default__.safeDivisionInt(p1, d_713_v50_), (p1) * (p1))
                rhs105_ = (d_742_v75_) - ((d_742_v75_) | (d_742_v75_))
                rhs106_ = d_734_v69_
                rhs107_ = d_743_v76_
                lhs71_ = d_631_v1_
                lhs72_ = default__.safeIndex(212, (d_631_v1_).length(0))
                d_713_v50_ = rhs103_
                lhs71_[lhs72_] = rhs104_
                d_742_v75_ = rhs105_
                d_734_v69_ = rhs106_
                d_743_v76_ = rhs107_
                (globalState).f0 = (self).f13
        elif True:
            d_744_v77_: _dafny.Seq
            d_744_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvlcjvnp"))
            d_744_v77_ = (self).f12
            d_745_v78_: str
            d_745_v78_ = _dafny.CodePoint('q')
            d_746_v79_: _dafny.Map
            d_746_v79_ = _dafny.Map({(self).f13: d_745_v78_})
            d_747_v80_: _dafny.Map
            d_747_v80_ = _dafny.Map({not((p1) < (p1)): d_746_v79_})
            d_747_v80_ = _dafny.Map({p0: d_746_v79_})
            d_748_v81_: int
            d_748_v81_ = 337
            d_749_v82_: _dafny.Seq
            d_749_v82_ = _dafny.SeqWithoutIsStrInference([False])
            d_750_v83_: _dafny.MultiSet
            d_750_v83_ = _dafny.MultiSet([d_641_v7_])
            d_751_v84_: _dafny.Seq
            d_751_v84_ = _dafny.SeqWithoutIsStrInference([(self).f13, default__.fm21(len(_dafny.SeqWithoutIsStrInference([d_749_v82_, _dafny.SeqWithoutIsStrInference([p0])])), d_750_v83_, globalState)])
            d_748_v81_ = ((886) * (p1)) + (default__.fm31(p0, d_748_v81_, False, d_751_v84_, globalState))
            d_752_v85_: _dafny.Array
            def lambda33_(d_753_i8_):
                return (self).f12

            init15_ = lambda33_
            nw107_ = _dafny.Array(None, 12)
            for i0_15_ in range(nw107_.length(0)):
                nw107_[i0_15_] = init15_(i0_15_)
            d_752_v85_ = nw107_
            index110_ = default__.safeIndex(603, (d_752_v85_).length(0))
            (d_752_v85_)[index110_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xoffjr"))
            index111_ = default__.safeIndex(603, (d_752_v85_).length(0))
            (d_752_v85_)[index111_] = (self).f12
            index112_ = default__.safeIndex(436, (d_631_v1_).length(0))
            (d_631_v1_)[index112_] = 267
            index113_ = default__.safeIndex(436, (d_631_v1_).length(0))
            (d_631_v1_)[index113_] = p1
        d_754_v86_: _dafny.Map
        d_754_v86_ = _dafny.Map({d_644_v10_: d_631_v1_})
        d_755_v87_: D1
        d_755_v87_ = D1_DC3(d_628_v0_)
        d_756_v88_: _dafny.Map
        d_756_v88_ = _dafny.Map({((d_754_v86_)[_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tyathfrj")) for d_757_i9_ in range(default__.abs(303))])] if (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tyathfrj")) for d_758_i9_ in range(default__.abs(303))])) in (d_754_v86_) else d_631_v1_): d_755_v87_})
        r0 = d_756_v88_
        return r0

    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13

class C2(T0, T1):
    def  __init__(self):
        self._f3: bool = False
        self._f4: bool = False
        self._f10: bool = False
        self._f11: bool = False
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
    def f4(self):
        return self._f4
    def ctor__(self, f10, f11, f3, f4):
        (self)._f10 = f10
        (self)._f11 = f11
        (self).f3 = f3
        (self)._f4 = f4

    def fm0(self, globalState):
        return (((_dafny.MultiSet([_dafny.Map({338: 547})])) - (_dafny.MultiSet([_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))): 597})]))).cardinality) * ((0) - (((0) - (len(_dafny.Map({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([832]))])).cardinality: self.f3})))) * (337)))

    def fm16(self, p0, p1, p2, p3, globalState):
        source13_ = D0_DC0(565)
        if source13_.is_DC1:
            d_759___mcc_h0_ = source13_.cf1
            d_760___mcc_h1_ = source13_.cf2
            d_761_cf2_ = d_760___mcc_h1_
            d_762_cf1_ = d_759___mcc_h0_
            return 482
        elif source13_.is_DC2:
            d_763___mcc_h2_ = source13_.cf3
            d_764___mcc_h3_ = source13_.cf4
            d_765___mcc_h4_ = source13_.cf5
            d_766_cf5_ = d_765___mcc_h4_
            d_767_cf4_ = d_764___mcc_h3_
            d_768_cf3_ = d_763___mcc_h2_
            return (d_767_cf4_) - (-730)
        elif True:
            d_769___mcc_h5_ = source13_.cf0
            d_770_cf0_ = d_769___mcc_h5_
            return len(_dafny.Map({(self).f10: (0) - (d_770_cf0_)}))

    def fm17(self, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdorts"))

    def m0(self, p0, p1, p2, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        d_771_v0_: _dafny.Seq
        d_771_v0_ = _dafny.SeqWithoutIsStrInference([p2])
        hi4_ = (d_771_v0_)[default__.safeIndex(p2, len(d_771_v0_))]
        for d_772_i0_ in range(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gqk"))), hi4_):
            d_773_v1_: _dafny.Array
            def lambda34_(d_774_p2_):
                def lambda35_(d_775_i1_):
                    return (d_774_p2_) <= (33)

                return lambda35_

            init16_ = lambda34_(p2)
            nw108_ = _dafny.Array(None, 15)
            for i0_16_ in range(nw108_.length(0)):
                nw108_[i0_16_] = init16_(i0_16_)
            d_773_v1_ = nw108_
            index114_ = default__.safeIndex(835, (d_773_v1_).length(0))
            (d_773_v1_)[index114_] = True
            index115_ = default__.safeIndex(835, (d_773_v1_).length(0))
            (d_773_v1_)[index115_] = not((self).f4)
            r1 = (p2 if False else p2)
            d_776_v2_: str
            d_776_v2_ = _dafny.CodePoint('v')
            d_777_v3_: _dafny.Map
            d_777_v3_ = _dafny.Map({(self).f10: d_773_v1_})
            d_778_v4_: D1
            d_778_v4_ = D1_DC3(d_773_v1_)
            d_779_v5_: _dafny.Array
            nw109_ = _dafny.Array(None, 28)
            nw109_[int(0)] = d_773_v1_
            nw109_[int(1)] = d_773_v1_
            nw109_[int(2)] = d_773_v1_
            nw109_[int(3)] = d_773_v1_
            nw109_[int(4)] = d_773_v1_
            nw109_[int(5)] = (d_773_v1_ if default__.fm18(d_772_i0_, d_776_v2_, d_772_i0_, globalState) else d_773_v1_)
            nw109_[int(6)] = d_773_v1_
            nw109_[int(7)] = d_773_v1_
            nw109_[int(8)] = d_773_v1_
            nw109_[int(9)] = ((d_777_v3_)[(self).f11] if ((self).f11) in (d_777_v3_) else d_773_v1_)
            nw109_[int(10)] = d_773_v1_
            nw109_[int(11)] = d_773_v1_
            nw109_[int(12)] = d_773_v1_
            nw109_[int(13)] = d_773_v1_
            nw109_[int(14)] = d_773_v1_
            nw109_[int(15)] = (((d_777_v3_)[(self).f4] if ((self).f4) in (d_777_v3_) else (d_778_v4_).cf6) if p0 else d_773_v1_)
            nw109_[int(16)] = d_773_v1_
            nw109_[int(17)] = d_773_v1_
            nw109_[int(18)] = d_773_v1_
            nw109_[int(19)] = d_773_v1_
            nw109_[int(20)] = (d_773_v1_ if (d_773_v1_)[default__.safeIndex(835, (d_773_v1_).length(0))] else d_773_v1_)
            nw109_[int(21)] = d_773_v1_
            nw109_[int(22)] = d_773_v1_
            nw109_[int(23)] = d_773_v1_
            nw109_[int(24)] = d_773_v1_
            nw109_[int(25)] = (d_773_v1_ if True else d_773_v1_)
            nw109_[int(26)] = d_773_v1_
            nw109_[int(27)] = d_773_v1_
            d_779_v5_ = nw109_
            d_780_v6_: _dafny.Seq
            d_780_v6_ = _dafny.SeqWithoutIsStrInference([self.f3, not((self).f4)])
            d_781_v7_: _dafny.Seq
            d_781_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
            d_782_v8_: _dafny.Map
            d_782_v8_ = _dafny.Map({len((d_780_v6_).set(default__.safeIndex(len(d_781_v7_), len(d_780_v6_)), False)): d_773_v1_})
            index116_ = default__.safeIndex(276, (d_779_v5_).length(0))
            (d_779_v5_)[index116_] = ((d_782_v8_)[d_772_i0_] if (d_772_i0_) in (d_782_v8_) else d_773_v1_)
            index117_ = default__.safeIndex(276, (d_779_v5_).length(0))
            (d_779_v5_)[index117_] = d_773_v1_
            r1 = (0) - (((d_772_i0_) + (819)) * (d_772_i0_))
        d_783_v9_: _dafny.Map
        d_783_v9_ = _dafny.Map({self.f3: self.f3})
        r1 = default__.safeDivisionInt(p2, len(d_783_v9_))
        (self).f3 = not(True)
        d_784_v10_: _dafny.Set
        d_784_v10_ = _dafny.Set({(self).f11})
        d_785_v11_: _dafny.Map
        d_785_v11_ = _dafny.Map({p2: len(d_784_v10_)})
        d_786_v12_: _dafny.MultiSet
        d_786_v12_ = _dafny.MultiSet([(self).f4, (self).f10, not((self).f4)])
        d_787_v13_: _dafny.Map
        d_787_v13_ = _dafny.Map({(self).f4: p2})
        d_788_v14_: _dafny.Map
        d_788_v14_ = _dafny.Map({self.f3: d_787_v13_})
        d_789_v15_: _dafny.Array
        nw110_ = _dafny.Array(None, 13)
        nw110_[int(0)] = ((d_785_v11_)[p2] if (p2) in (d_785_v11_) else p2)
        nw110_[int(1)] = p2
        nw110_[int(2)] = 491
        nw110_[int(3)] = (d_771_v0_)[default__.safeIndex(p2, len(d_771_v0_))]
        nw110_[int(4)] = ((d_786_v12_)[(self).f4] if ((self).f4) in (d_786_v12_) else p2)
        nw110_[int(5)] = p2
        nw110_[int(6)] = p2
        nw110_[int(7)] = len(d_788_v14_)
        nw110_[int(8)] = p2
        nw110_[int(9)] = p2
        nw110_[int(10)] = (0) - (p2)
        nw110_[int(11)] = p2
        nw110_[int(12)] = p2
        d_789_v15_ = nw110_
        d_790_v16_: _dafny.Map
        d_790_v16_ = _dafny.Map({d_789_v15_: False})
        (globalState).f0 = (d_789_v15_) in (d_790_v16_)
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_789_v15_).length(0)):
            d_791_i2_: int = guard_loop_1_
            if (True) and (((0) <= (d_791_i2_)) and ((d_791_i2_) < ((d_789_v15_).length(0)))):
                (d_789_v15_)[(d_791_i2_)] = default__.safeModuloInt(d_791_i2_, p2)
        if (self).f4:
            d_792_v17_: str
            d_792_v17_ = _dafny.CodePoint('n')
            d_793_v18_: _dafny.Map
            d_793_v18_ = _dafny.Map({not(default__.fm18(p2, d_792_v17_, p2, globalState)): (self).fm17(globalState)})
            d_793_v18_ = (_dafny.Map({p0: _dafny.SeqWithoutIsStrInference([d_792_v17_ for d_794_i3_ in range(default__.abs(593))])})) | ((d_793_v18_) | (d_793_v18_))
            d_795_v19_: _dafny.MultiSet
            d_795_v19_ = _dafny.MultiSet([d_783_v9_])
            d_771_v0_ = _dafny.SeqWithoutIsStrInference([(p2) + ((d_795_v19_).cardinality)])
            d_796_v20_: _dafny.Seq
            d_796_v20_ = _dafny.SeqWithoutIsStrInference([d_787_v13_])
            r1 = (self).fm16(p2, (p0) not in (_dafny.MultiSet([self.f3, (self).f4, (self).f11, (self).f11, (self).f4])), (0) - (len(d_796_v20_)), (p1) | (p1), globalState)
            (self).f3 = self.f3
            index118_ = default__.safeIndex(532, (d_789_v15_).length(0))
            (d_789_v15_)[index118_] = (270) + (p2)
            index119_ = default__.safeIndex(532, (d_789_v15_).length(0))
            (d_789_v15_)[index119_] = p2
        elif True:
            d_797_v21_: C0
            nw111_ = C0()
            nw111_.ctor__()
            d_797_v21_ = nw111_
            d_798_v22_: _dafny.Array
            nw112_ = _dafny.Array(_dafny.Seq({}), 19)
            d_798_v22_ = nw112_
            d_799_v23_: _dafny.Seq
            d_799_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pgmkqgjiw"))
            index120_ = default__.safeIndex(623, (d_798_v22_).length(0))
            (d_798_v22_)[index120_] = (default__.fm19(globalState)).set(default__.safeIndex(len(d_799_v23_), len(default__.fm19(globalState))), p2)
            d_800_v24_: str
            d_800_v24_ = _dafny.CodePoint('e')
            d_801_v25_: _dafny.Seq
            d_801_v25_ = _dafny.SeqWithoutIsStrInference([default__.fm20(p0, globalState), d_786_v12_])
            index121_ = default__.safeIndex(623, (d_798_v22_).length(0))
            rhs108_ = (d_797_v21_ if (self).f10 else d_797_v21_)
            rhs109_ = False
            rhs110_ = not(default__.fm18(p2, d_800_v24_, (d_797_v21_).fm10(771, (self).f10, _dafny.MultiSet([len(d_801_v25_)]), globalState), globalState))
            rhs111_ = default__.fm19(globalState)
            rhs112_ = self.f3
            lhs73_ = self
            lhs74_ = globalState
            lhs75_ = d_798_v22_
            lhs76_ = default__.safeIndex(623, (d_798_v22_).length(0))
            lhs77_ = globalState
            d_797_v21_ = rhs108_
            lhs73_.f3 = rhs109_
            lhs74_.f0 = rhs110_
            lhs75_[lhs76_] = rhs111_
            lhs77_.f0 = rhs112_
            d_800_v24_ = d_800_v24_
            d_802_v26_: D1
            d_802_v26_ = D1_DC5(default__.fm18(p2, d_800_v24_, p2, globalState), (self).f4, p2, (0) - (p2))
            (globalState).f0 = (d_802_v26_).cf9
            d_799_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgljdqvut"))
            (globalState).f0 = (self).f11
        r0 = (d_789_v15_ if self.f3 else d_789_v15_)
        r1 = p2
        return r0, r1

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        d_803_v0_: str
        d_803_v0_ = _dafny.CodePoint('y')
        d_803_v0_ = _dafny.CodePoint('p')
        d_804_v1_: _dafny.Seq
        d_804_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tenl"))
        d_805_v2_: T1
        nw113_ = C1()
        nw113_.ctor__(d_804_v1_, (self).f4)
        d_805_v2_ = nw113_
        d_805_v2_ = d_805_v2_
        d_806_v3_: _dafny.Array
        nw114_ = _dafny.Array(False, 14)
        d_806_v3_ = nw114_
        index122_ = default__.safeIndex(226, (d_806_v3_).length(0))
        (d_806_v3_)[index122_] = False
        d_807_v4_: _dafny.Array
        nw115_ = _dafny.Array(_dafny.Map({}), 20)
        d_807_v4_ = nw115_
        d_808_v5_: _dafny.Set
        d_808_v5_ = _dafny.Set({(p1).f4})
        index123_ = default__.safeIndex(226, (d_806_v3_).length(0))
        rhs113_ = ((p1).f4 if self.f3 else (p1).f4)
        rhs114_ = ((d_808_v5_) | (d_808_v5_)).issubset(d_808_v5_)
        rhs115_ = d_807_v4_
        lhs78_ = self
        lhs79_ = d_806_v3_
        lhs80_ = default__.safeIndex(226, (d_806_v3_).length(0))
        lhs78_.f3 = rhs113_
        lhs79_[lhs80_] = rhs114_
        d_807_v4_ = rhs115_
        d_809_v6_: _dafny.Seq
        d_809_v6_ = _dafny.SeqWithoutIsStrInference([(d_806_v3_)[default__.safeIndex(226, (d_806_v3_).length(0))], (p1).f4, (p1).f4, (self).f4, (self).f4])
        d_810_v7_: _dafny.Map
        d_810_v7_ = _dafny.Map({(d_806_v3_)[default__.safeIndex(226, (d_806_v3_).length(0))]: len(d_809_v6_)})
        d_811_v8_: _dafny.Seq
        d_811_v8_ = _dafny.SeqWithoutIsStrInference([p0, p0, len(d_810_v7_)])
        r0 = len(d_811_v8_)
        d_812_v9_: _dafny.Set
        d_812_v9_ = _dafny.Set({d_803_v0_})
        d_813_v10_: _dafny.Seq
        d_813_v10_ = _dafny.SeqWithoutIsStrInference([d_812_v9_, d_812_v9_])
        d_814_v11_: C0
        nw116_ = C0()
        nw116_.ctor__()
        d_814_v11_ = nw116_
        d_815_v12_: _dafny.Map
        d_815_v12_ = _dafny.Map({d_814_v11_: (d_806_v3_)[default__.safeIndex(226, (d_806_v3_).length(0))]})
        d_816_v13_: _dafny.Map
        d_816_v13_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_817_i0_ in range(default__.abs(705))])): (_dafny.Map({d_814_v11_: (d_806_v3_)[default__.safeIndex(226, (d_806_v3_).length(0))]})) | (d_815_v12_)})
        rhs116_ = p0
        rhs117_ = d_813_v10_
        rhs118_ = d_803_v0_
        rhs119_ = d_816_v13_
        r0 = rhs116_
        d_813_v10_ = rhs117_
        d_803_v0_ = rhs118_
        d_816_v13_ = rhs119_
        if (d_806_v3_)[default__.safeIndex(226, (d_806_v3_).length(0))]:
            (globalState).f0 = (249) != (p0)
            if p1.f3:
                d_818_v14_: _dafny.Set
                d_818_v14_ = _dafny.Set({default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([p0])), p0), len((d_809_v6_) + (_dafny.SeqWithoutIsStrInference([(self).f11, p1.f3])))})
                d_818_v14_ = d_818_v14_
                d_819_v15_: _dafny.Map
                d_819_v15_ = _dafny.Map({p1: p0})
                d_819_v15_ = (d_819_v15_).set(p1, p0)
                d_820_v16_: _dafny.Array
                nw117_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
                d_820_v16_ = nw117_
                index124_ = default__.safeIndex(147, (d_820_v16_).length(0))
                (d_820_v16_)[index124_] = d_804_v1_
                index125_ = default__.safeIndex(147, (d_820_v16_).length(0))
                (d_820_v16_)[index125_] = (d_804_v1_) + ((d_804_v1_) + (d_804_v1_))
                d_821_v17_: _dafny.Map
                d_821_v17_ = _dafny.Map({p0: p0})
                d_822_v18_: _dafny.Seq
                d_822_v18_ = _dafny.SeqWithoutIsStrInference([d_821_v17_, d_821_v17_])
                d_823_v19_: _dafny.Seq
                d_823_v19_ = _dafny.SeqWithoutIsStrInference([d_821_v17_, (d_822_v18_)[default__.safeIndex(len(d_809_v6_), len(d_822_v18_))]])
                d_824_v20_: _dafny.Map
                d_824_v20_ = _dafny.Map({d_823_v19_: True})
                (p1).f3 = ((d_824_v20_)[(d_823_v19_) + (d_822_v18_)] if ((d_823_v19_) + (d_822_v18_)) in (d_824_v20_) else ((d_806_v3_)[default__.safeIndex(226, (d_806_v3_).length(0))] if p1.f3 else p1.f3))
                rhs120_ = p0
                rhs121_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccdmppcba"))) + (d_804_v1_)) + (d_804_v1_)
                r0 = rhs120_
                d_804_v1_ = rhs121_
            elif True:
                d_825_v21_: _dafny.Map
                d_825_v21_ = _dafny.Map({p0: -386})
                d_826_v22_: _dafny.MultiSet
                d_826_v22_ = _dafny.MultiSet([d_825_v21_, d_825_v21_])
                index126_ = default__.safeIndex(226, (d_806_v3_).length(0))
                (d_806_v3_)[index126_] = not(default__.fm21(p0, d_826_v22_, globalState))
                d_803_v0_ = d_803_v0_
                (globalState).f0 = (p1).f4
                d_803_v0_ = (d_803_v0_ if not(p1.f3) else d_803_v0_)
                d_811_v8_ = ((d_811_v8_) + (_dafny.SeqWithoutIsStrInference([p0 for d_827_i1_ in range(default__.abs(189))]))).set(default__.safeIndex(p0, len((d_811_v8_) + (_dafny.SeqWithoutIsStrInference([p0 for d_828_i1_ in range(default__.abs(189))])))), (p0) - (p0))
            d_829_v23_: _dafny.Map
            d_829_v23_ = _dafny.Map({p0: True})
            d_830_v24_: _dafny.Map
            d_830_v24_ = _dafny.Map({len(d_829_v23_): p0})
            d_831_v25_: _dafny.MultiSet
            d_831_v25_ = _dafny.MultiSet([len(d_830_v24_)])
            d_810_v7_ = (d_810_v7_).set((p1).f4, ((d_831_v25_)[(0) - (p0)] if ((0) - (p0)) in (d_831_v25_) else (self).fm16(p0, (self).f4, p0, (_dafny.Map({(0) - (p0): self.f3})).set(len(d_804_v1_), (self).f4), globalState)))
            index127_ = default__.safeIndex(226, (d_806_v3_).length(0))
            (d_806_v3_)[index127_] = (self).f11
            (globalState).f0 = p1.f3
        elif True:
            d_832_v26_: _dafny.Array
            def lambda36_(d_833_p0_):
                def lambda37_(d_834_i2_):
                    return (d_834_i2_) * (d_833_p0_)

                return lambda37_

            init17_ = lambda36_(p0)
            nw118_ = _dafny.Array(None, 28)
            for i0_17_ in range(nw118_.length(0)):
                nw118_[i0_17_] = init17_(i0_17_)
            d_832_v26_ = nw118_
            d_835_v27_: _dafny.MultiSet
            d_835_v27_ = _dafny.MultiSet([d_832_v26_, d_832_v26_])
            d_809_v6_ = _dafny.SeqWithoutIsStrInference([(d_835_v27_).ispropersubset(d_835_v27_)])
            r0 = p0
            d_832_v26_ = (d_832_v26_ if not((p1).f4) else d_832_v26_)
            d_836_v28_: _dafny.Map
            d_836_v28_ = _dafny.Map({p1.f3: (self).f10})
            d_836_v28_ = _dafny.Map({(p1).f4: not(not ((d_806_v3_)[default__.safeIndex(226, (d_806_v3_).length(0))]) or (not((p1).f4)))})
            (globalState).f0 = not(self.f3)
        d_837_v29_: _dafny.Map
        d_837_v29_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0]): p0})
        d_838_v30_: _dafny.MultiSet
        d_838_v30_ = _dafny.MultiSet([p0, 682, p0, ((d_837_v29_)[d_811_v8_] if (d_811_v8_) in (d_837_v29_) else p0)])
        r0 = (d_814_v11_).fm10(len((d_804_v1_) + (d_804_v1_)), True, (d_838_v30_).intersection(d_838_v30_), globalState)
        return r0

    def m2(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_839_v0_: _dafny.Set
        d_839_v0_ = _dafny.Set({_dafny.Map({True: self.f3})})
        d_840_i0_: int
        d_840_i0_ = 0
        with _dafny.label("3"):
            while (d_839_v0_).isdisjoint(d_839_v0_):
                with _dafny.c_label("3"):
                    if (d_840_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_840_i0_ = (d_840_i0_) + (1)
                    (p1).f3 = p1.f3
                    d_841_v1_: _dafny.Map
                    d_841_v1_ = _dafny.Map({not((p1).f4): -940})
                    d_841_v1_ = d_841_v1_
                    d_842_v3_: _dafny.Array
                    def lambda38_(d_843_p1_):
                        def lambda39_(d_844_i1_):
                            def iife61_():
                                coll33_ = _dafny.Map()
                                compr_33_: int
                                for compr_33_ in (_dafny.Set({413})).Elements:
                                    d_845_v2_: int = compr_33_
                                    if (d_845_v2_) in (_dafny.Set({413})):
                                        coll33_[(d_845_v2_) - (len(_dafny.Map({d_843_p1_.f3: (d_843_p1_).f4})))] = True
                                return _dafny.Map(coll33_)
                            return (_dafny.Map({(_dafny.MultiSet([(self).f4])).cardinality: (self).f10})) | (iife61_()
                            )

                        return lambda39_

                    init18_ = lambda38_(p1)
                    nw119_ = _dafny.Array(None, 27)
                    for i0_18_ in range(nw119_.length(0)):
                        nw119_[i0_18_] = init18_(i0_18_)
                    d_842_v3_ = nw119_
                    d_846_v4_: _dafny.Seq
                    d_846_v4_ = _dafny.SeqWithoutIsStrInference([d_842_v3_])
                    d_847_v5_: int
                    d_847_v5_ = 395
                    d_842_v3_ = (d_846_v4_)[default__.safeIndex((d_847_v5_) - (d_847_v5_), len(d_846_v4_))]
                    d_848_v6_: _dafny.Seq
                    d_848_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
                    d_849_v7_: _dafny.Map
                    d_849_v7_ = _dafny.Map({d_847_v5_: p1.f3})
                    d_850_v8_: _dafny.Seq
                    d_850_v8_ = _dafny.SeqWithoutIsStrInference([d_847_v5_])
                    d_851_v9_: _dafny.Map
                    d_851_v9_ = _dafny.Map({p1.f3: d_850_v8_})
                    d_852_v10_: D7
                    d_852_v10_ = D7_DC17(d_847_v5_, (self).f4, (p1).f4, d_847_v5_)
                    d_853_v11_: _dafny.Map
                    d_853_v11_ = _dafny.Map({d_852_v10_: (self).f11})
                    d_854_v12_: _dafny.Array
                    nw120_ = _dafny.Array(None, 3)
                    nw120_[int(0)] = d_847_v5_
                    nw120_[int(1)] = len(d_853_v11_)
                    nw120_[int(2)] = d_847_v5_
                    d_854_v12_ = nw120_
                    d_855_v13_: _dafny.Seq
                    d_855_v13_ = _dafny.SeqWithoutIsStrInference([d_854_v12_, d_854_v12_])
                    d_856_v15_: _dafny.MultiSet
                    d_856_v15_ = _dafny.MultiSet([d_847_v5_])
                    d_857_v16_: _dafny.Seq
                    def iife62_():
                        coll34_ = _dafny.Map()
                        compr_34_: int
                        for compr_34_ in (d_856_v15_).Elements:
                            d_858_v14_: int = compr_34_
                            if (d_858_v14_) in (d_856_v15_):
                                coll34_[default__.safeDivisionInt(d_858_v14_, d_847_v5_)] = p1.f3
                        return _dafny.Map(coll34_)
                    d_857_v16_ = _dafny.SeqWithoutIsStrInference([default__.fm21(d_847_v5_, default__.fm22(d_847_v5_, (self).fm16(d_847_v5_, p1.f3, d_847_v5_, iife62_()
                    , globalState), d_847_v5_, globalState), globalState), (p1).f4, (self).f11])
                    d_859_v18_: _dafny.Map
                    d_859_v18_ = _dafny.Map({d_847_v5_: _dafny.SeqWithoutIsStrInference([p0 for d_860_i2_ in range(default__.abs(-868))])})
                    d_861_v19_: _dafny.Map
                    d_861_v19_ = _dafny.Map({True: d_854_v12_})
                    d_862_v20_: _dafny.Map
                    d_862_v20_ = _dafny.Map({(p1).f4: _dafny.MultiSet([d_854_v12_, ((d_861_v19_)[(p1).f4] if ((p1).f4) in (d_861_v19_) else d_854_v12_)])})
                    d_863_v21_: _dafny.Map
                    d_863_v21_ = _dafny.Map({d_847_v5_: d_854_v12_})
                    d_864_v22_: _dafny.MultiSet
                    d_864_v22_ = _dafny.MultiSet([((d_863_v21_)[d_847_v5_] if (d_847_v5_) in (d_863_v21_) else d_854_v12_)])
                    d_865_v23_: _dafny.Array
                    nw121_ = _dafny.Array(None, 27)
                    nw121_[int(0)] = not(p1.f3)
                    nw121_[int(1)] = not(self.f3)
                    nw121_[int(2)] = (p1).f4
                    nw121_[int(3)] = (self).f4
                    nw121_[int(4)] = not((p1).f4)
                    nw121_[int(5)] = self.f3
                    nw121_[int(6)] = False
                    nw121_[int(7)] = (p1).f4
                    nw121_[int(8)] = (self).f4
                    nw121_[int(9)] = p1.f3
                    nw121_[int(10)] = not((p1).f4)
                    nw121_[int(11)] = (self).f11
                    nw121_[int(12)] = (d_848_v6_) < (d_848_v6_)
                    nw121_[int(13)] = ((p1).f4 if ((d_849_v7_)[d_847_v5_] if (d_847_v5_) in (d_849_v7_) else (self).f4) else not(p1.f3))
                    nw121_[int(14)] = (d_850_v8_) != (((d_851_v9_)[True] if (True) in (d_851_v9_) else _dafny.SeqWithoutIsStrInference([d_847_v5_])))
                    nw121_[int(15)] = (p1).f4
                    nw121_[int(16)] = not((d_854_v12_) in (d_855_v13_))
                    nw121_[int(17)] = (self).f4
                    nw121_[int(18)] = (p1).f4
                    nw121_[int(19)] = (p1).f4
                    nw121_[int(20)] = False
                    nw121_[int(21)] = (p1.f3) == ((d_857_v16_)[default__.safeIndex(d_847_v5_, len(d_857_v16_))])
                    def iife63_():
                        coll35_ = _dafny.Map()
                        compr_35_: int
                        for compr_35_ in _dafny.IntegerRange(118, 515):
                            d_866_v17_: int = compr_35_
                            if ((118) <= (d_866_v17_)) and ((d_866_v17_) < (515)):
                                coll35_[default__.safeDivisionInt(d_866_v17_, 579)] = d_848_v6_
                        return _dafny.Map(coll35_)
                    nw121_[int(22)] = (iife63_()
                    ) == (d_859_v18_)
                    nw121_[int(23)] = (_dafny.MultiSet([d_847_v5_])).issubset(_dafny.MultiSet([len(d_855_v13_)]))
                    nw121_[int(24)] = (self).f11
                    nw121_[int(25)] = (((d_862_v20_)[(p1).f4] if ((p1).f4) in (d_862_v20_) else d_864_v22_)).isdisjoint(d_864_v22_)
                    nw121_[int(26)] = (not((p1).f4) if (d_857_v16_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p1.f3, (self).f11])), len(d_857_v16_))] else (p1).f4)
                    d_865_v23_ = nw121_
                    index128_ = default__.safeIndex(940, (d_865_v23_).length(0))
                    (d_865_v23_)[index128_] = self.f3
                    d_867_v24_: _dafny.Map
                    d_867_v24_ = _dafny.Map({d_865_v23_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))})
                    d_868_v25_: _dafny.Map
                    d_868_v25_ = _dafny.Map({d_847_v5_: d_847_v5_})
                    d_869_v26_: _dafny.MultiSet
                    d_869_v26_ = _dafny.MultiSet([d_868_v25_, ((d_868_v25_).set(d_847_v5_, d_847_v5_)).set(d_847_v5_, d_847_v5_), d_868_v25_])
                    index129_ = default__.safeIndex(431, (d_865_v23_).length(0))
                    (d_865_v23_)[index129_] = default__.fm21(len(d_867_v24_), d_869_v26_, globalState)
                    d_870_v27_: D0
                    d_870_v27_ = D0_DC2(p1.f3, d_847_v5_, len(d_848_v6_))
                    d_871_v28_: _dafny.Seq
                    d_871_v28_ = _dafny.SeqWithoutIsStrInference([d_865_v23_])
                    index130_ = default__.safeIndex(940, (d_865_v23_).length(0))
                    index131_ = default__.safeIndex(431, (d_865_v23_).length(0))
                    rhs122_ = (d_870_v27_).cf3
                    rhs123_ = not ((self).f4) or ((_dafny.SeqWithoutIsStrInference([d_865_v23_])) != (d_871_v28_))
                    lhs81_ = d_865_v23_
                    lhs82_ = default__.safeIndex(940, (d_865_v23_).length(0))
                    lhs83_ = d_865_v23_
                    lhs84_ = default__.safeIndex(431, (d_865_v23_).length(0))
                    lhs81_[lhs82_] = rhs122_
                    lhs83_[lhs84_] = rhs123_
                    pass
            pass
        d_872_v29_: int
        d_872_v29_ = 443
        d_873_v30_: _dafny.Seq
        d_873_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
        d_874_v31_: _dafny.Seq
        d_874_v31_ = _dafny.SeqWithoutIsStrInference([len(d_873_v30_), (0) - (d_872_v29_)])
        d_875_v32_: _dafny.Map
        d_875_v32_ = _dafny.Map({d_872_v29_: (d_874_v31_)[default__.safeIndex(d_872_v29_, len(d_874_v31_))]})
        d_876_v33_: _dafny.MultiSet
        d_876_v33_ = _dafny.MultiSet([d_875_v32_])
        d_877_v34_: _dafny.Seq
        d_877_v34_ = _dafny.SeqWithoutIsStrInference([True])
        d_878_v36_: _dafny.Map
        d_878_v36_ = _dafny.Map({-61: (self).f4})
        d_879_v37_: _dafny.Array
        d_880_v38_: int
        out20_: _dafny.Array
        out21_: int
        def iife64_():
            coll36_ = _dafny.Map()
            compr_36_: int
            for compr_36_ in (d_878_v36_).keys.Elements:
                d_881_v35_: int = compr_36_
                if (d_881_v35_) in (d_878_v36_):
                    coll36_[(d_881_v35_) + (d_872_v29_)] = (p1).f4
            return _dafny.Map(coll36_)
        out20_, out21_ = (p1).m0((d_872_v29_) >= (default__.fm31(default__.fm21(d_872_v29_, d_876_v33_, globalState), d_872_v29_, (self).f10, d_877_v34_, globalState)), (D8_DC18(iife64_()
)).cf30, (209) + (d_872_v29_), globalState)
        d_879_v37_ = out20_
        d_880_v38_ = out21_
        (p1).f3 = (d_877_v34_)[default__.safeIndex(default__.safeModuloInt(d_880_v38_, (0) - (d_872_v29_)), len(d_877_v34_))]
        d_872_v29_ = (0) - (default__.safeDivisionInt(len(d_873_v30_), (d_880_v38_) * (d_872_v29_)))
        d_882_v39_: C1
        nw122_ = C1()
        nw122_.ctor__(d_873_v30_, (self).f10)
        d_882_v39_ = nw122_
        d_883_v40_: str
        d_883_v40_ = _dafny.CodePoint('y')
        d_883_v40_ = p0
        r0 = d_873_v30_
        return r0

    def m9(self, globalState):
        d_884_v0_: int
        d_884_v0_ = 44
        d_885_v1_: _dafny.Set
        d_885_v1_ = _dafny.Set({d_884_v0_})
        d_886_v2_: _dafny.Map
        d_886_v2_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htgmqesq"))): d_885_v1_})
        d_887_v3_: _dafny.Set
        d_887_v3_ = d_885_v1_
        if (((d_886_v2_)[d_884_v0_] if (d_884_v0_) in (d_886_v2_) else d_885_v1_)).isdisjoint((d_887_v3_)):
            d_888_v4_: _dafny.Seq
            d_888_v4_ = _dafny.SeqWithoutIsStrInference([(self).f4])
            if not((d_888_v4_)[default__.safeIndex(d_884_v0_, len(d_888_v4_))]):
                d_889_v5_: _dafny.Array
                nw123_ = _dafny.Array(int(0), 16)
                d_889_v5_ = nw123_
                d_890_v6_: _dafny.Array
                nw124_ = _dafny.Array(None, 9)
                nw124_[int(0)] = d_889_v5_
                nw124_[int(1)] = d_889_v5_
                nw124_[int(2)] = d_889_v5_
                nw124_[int(3)] = d_889_v5_
                nw124_[int(4)] = d_889_v5_
                nw124_[int(5)] = d_889_v5_
                nw124_[int(6)] = d_889_v5_
                nw124_[int(7)] = d_889_v5_
                nw124_[int(8)] = d_889_v5_
                d_890_v6_ = nw124_
                index132_ = default__.safeIndex(580, (d_890_v6_).length(0))
                (d_890_v6_)[index132_] = d_889_v5_
                index133_ = default__.safeIndex(580, (d_890_v6_).length(0))
                def lambda40_(d_891_v0_):
                    def lambda41_(d_892_i0_):
                        return (d_892_i0_) - ((0) - (d_891_v0_))

                    return lambda41_

                init19_ = lambda40_(d_884_v0_)
                nw125_ = _dafny.Array(None, 23)
                for i0_19_ in range(nw125_.length(0)):
                    nw125_[i0_19_] = init19_(i0_19_)
                (d_890_v6_)[index133_] = nw125_
                d_893_v7_: _dafny.Map
                d_893_v7_ = _dafny.Map({self.f3: d_884_v0_})
                d_894_v8_: _dafny.Set
                d_894_v8_ = _dafny.Set({d_893_v7_})
                (globalState).f0 = (d_894_v8_).isdisjoint(d_894_v8_)
                d_895_v9_: _dafny.Map
                d_895_v9_ = _dafny.Map({d_884_v0_: d_884_v0_})
                (self).f3 = (d_895_v9_) == (d_895_v9_)
                d_896_v10_: _dafny.Array
                nw126_ = _dafny.Array(False, 24)
                d_896_v10_ = nw126_
                d_897_v11_: _dafny.Map
                d_897_v11_ = _dafny.Map({d_896_v10_: _dafny.Map({d_884_v0_: d_884_v0_})})
                d_897_v11_ = (d_897_v11_).set(d_896_v10_, (d_895_v9_) | (d_895_v9_))
                (self).f3 = not(self.f3)
            elif True:
                d_898_v12_: _dafny.Seq
                d_898_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldgivosl"))
                d_899_v13_: _dafny.Map
                d_899_v13_ = _dafny.Map({d_884_v0_: d_884_v0_})
                d_900_v14_: _dafny.Map
                d_900_v14_ = _dafny.Map({d_884_v0_: len(d_899_v13_)})
                d_901_v15_: _dafny.MultiSet
                d_901_v15_ = _dafny.MultiSet([d_900_v14_, d_900_v14_])
                d_902_v16_: _dafny.Map
                d_902_v16_ = _dafny.Map({d_884_v0_: default__.fm21(len(d_898_v12_), d_901_v15_, globalState)})
                d_902_v16_ = (d_902_v16_).set(((d_900_v14_)[d_884_v0_] if (d_884_v0_) in (d_900_v14_) else 360), (self).f4)
                d_903_v17_: _dafny.Array
                nw127_ = _dafny.Array(int(0), 22)
                d_903_v17_ = nw127_
                index134_ = default__.safeIndex(511, (d_903_v17_).length(0))
                (d_903_v17_)[index134_] = default__.safeDivisionInt((0) - (len(d_898_v12_)), d_884_v0_)
                d_904_v18_: _dafny.MultiSet
                d_904_v18_ = _dafny.MultiSet([d_884_v0_, d_884_v0_])
                index135_ = default__.safeIndex(511, (d_903_v17_).length(0))
                (d_903_v17_)[index135_] = ((d_904_v18_) | ((d_904_v18_) - (d_904_v18_))).cardinality
                d_905_v19_: str
                d_905_v19_ = _dafny.CodePoint('u')
                d_906_v20_: D8
                d_906_v20_ = D8_DC20(d_905_v19_, True)
                pat_let_tv33_ = d_905_v19_
                def iife65_(_pat_let14_0):
                    def iife66_(d_907_dt__update__tmp_h0_):
                        def iife67_(_pat_let15_0):
                            def iife68_(d_908_dt__update_hcf33_h0_):
                                return D8_DC20(d_908_dt__update_hcf33_h0_, (d_907_dt__update__tmp_h0_).cf34)
                            return iife68_(_pat_let15_0)
                        return iife67_(pat_let_tv33_)
                    return iife66_(_pat_let14_0)
                d_906_v20_ = iife65_(d_906_v20_)
                d_909_v21_: C1
                nw128_ = C1()
                nw128_.ctor__(_dafny.SeqWithoutIsStrInference([d_905_v19_ for d_910_i1_ in range(default__.abs(802))]), self.f3)
                d_909_v21_ = nw128_
                d_909_v21_ = d_909_v21_
                d_911_v22_: _dafny.Array
                def lambda42_(d_912_v19_):
                    def lambda43_(d_913_i2_):
                        return d_912_v19_

                    return lambda43_

                init20_ = lambda42_(d_905_v19_)
                nw129_ = _dafny.Array(None, 25)
                for i0_20_ in range(nw129_.length(0)):
                    nw129_[i0_20_] = init20_(i0_20_)
                d_911_v22_ = nw129_
                d_914_v23_: _dafny.MultiSet
                d_914_v23_ = _dafny.MultiSet([d_909_v21_])
                index136_ = default__.safeIndex(883, (d_911_v22_).length(0))
                (d_911_v22_)[index136_] = default__.fm23((d_909_v21_).f13, ((d_914_v23_).set(d_909_v21_, default__.abs((self).fm16(d_884_v0_, (d_909_v21_).f13, d_884_v0_, d_902_v16_, globalState)))).cardinality, globalState)
                d_915_v24_: D0
                d_915_v24_ = D0_DC2(self.f3, (0) - (d_884_v0_), len(_dafny.SeqWithoutIsStrInference([(0) - (d_884_v0_)])))
                d_916_v25_: _dafny.Seq
                d_916_v25_ = _dafny.SeqWithoutIsStrInference([(d_898_v12_) + ((d_909_v21_).f12), (d_898_v12_ if (d_915_v24_).cf3 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvkiceyku")))])
                d_917_v26_: C0
                nw130_ = C0()
                nw130_.ctor__()
                d_917_v26_ = nw130_
                d_918_v27_: _dafny.MultiSet
                d_918_v27_ = _dafny.MultiSet([d_917_v26_, d_917_v26_])
                d_919_v28_: _dafny.Map
                d_919_v28_ = _dafny.Map({d_884_v0_: d_888_v4_})
                d_920_v29_: _dafny.Set
                d_920_v29_ = _dafny.Set({(self).f11, (self).f4})
                index137_ = default__.safeIndex(511, (d_903_v17_).length(0))
                index138_ = default__.safeIndex(883, (d_911_v22_).length(0))
                rhs124_ = (0) - ((0) - (default__.fm31((d_918_v27_).ispropersubset(d_918_v27_), len((_dafny.Map({self.f3: (d_903_v17_)[default__.safeIndex(511, (d_903_v17_).length(0))]})).set(self.f3, (d_903_v17_)[default__.safeIndex(511, (d_903_v17_).length(0))])), (self).f4, d_888_v4_, globalState)))
                rhs125_ = d_905_v19_
                rhs126_ = (D10_DC22(d_903_v17_)).cf36
                rhs127_ = (d_917_v26_).fm11(d_905_v19_, default__.fm33((d_903_v17_)[default__.safeIndex(511, (d_903_v17_).length(0))], (((d_919_v28_)[len(_dafny.Map({(d_903_v17_)[default__.safeIndex(511, (d_903_v17_).length(0))]: (d_909_v21_).f13}))] if (len(_dafny.Map({(d_903_v17_)[default__.safeIndex(511, (d_903_v17_).length(0))]: (d_909_v21_).f13}))) in (d_919_v28_) else default__.fm34(len(_dafny.Map({(d_903_v17_)[default__.safeIndex(511, (d_903_v17_).length(0))]: (self).f11})), len(d_920_v29_), globalState))).set(default__.safeIndex(151, len(((d_919_v28_)[len(_dafny.Map({(d_903_v17_)[default__.safeIndex(511, (d_903_v17_).length(0))]: (d_909_v21_).f13}))] if (len(_dafny.Map({(d_903_v17_)[default__.safeIndex(511, (d_903_v17_).length(0))]: (d_909_v21_).f13}))) in (d_919_v28_) else default__.fm34(len(_dafny.Map({(d_903_v17_)[default__.safeIndex(511, (d_903_v17_).length(0))]: (self).f11})), len(d_920_v29_), globalState)))), (d_909_v21_).f13), globalState), globalState)
                rhs128_ = default__.fm35(not(self.f3), globalState)
                lhs85_ = d_903_v17_
                lhs86_ = default__.safeIndex(511, (d_903_v17_).length(0))
                lhs87_ = d_911_v22_
                lhs88_ = default__.safeIndex(883, (d_911_v22_).length(0))
                lhs85_[lhs86_] = rhs124_
                lhs87_[lhs88_] = rhs125_
                d_903_v17_ = rhs126_
                d_898_v12_ = rhs127_
                d_916_v25_ = rhs128_
            d_921_v30_: _dafny.Seq
            d_921_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "px"))
            d_921_v30_ = d_921_v30_
            if ((d_921_v30_) + (d_921_v30_)) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edlbvgnpp"))):
                d_922_v31_: _dafny.Map
                d_922_v31_ = _dafny.Map({(self).f4: True})
                d_923_v32_: C1
                nw131_ = C1()
                nw131_.ctor__(d_921_v30_, ((d_922_v31_)[self.f3] if (self.f3) in (d_922_v31_) else self.f3))
                d_923_v32_ = nw131_
                d_924_v33_: _dafny.Map
                d_924_v33_ = _dafny.Map({(self).f4: d_885_v1_})
                (globalState).f0 = (self.f3 if (self).f4 else (len(d_885_v1_)) in (((d_924_v33_)[(d_923_v32_).f13] if ((d_923_v32_).f13) in (d_924_v33_) else d_885_v1_)))
                d_925_v34_: _dafny.Map
                d_925_v34_ = _dafny.Map({d_884_v0_: d_884_v0_})
                d_926_v36_: _dafny.Seq
                d_926_v36_ = _dafny.SeqWithoutIsStrInference([d_884_v0_, d_884_v0_])
                d_927_v37_: _dafny.Map
                d_927_v37_ = _dafny.Map({346: default__.fm25(globalState)})
                d_928_v39_: _dafny.Array
                nw132_ = _dafny.Array(None, 20)
                nw132_[int(0)] = d_925_v34_
                nw132_[int(1)] = d_925_v34_
                nw132_[int(2)] = _dafny.Map({d_884_v0_: d_884_v0_})
                def iife69_():
                    coll37_ = _dafny.Map()
                    compr_37_: int
                    for compr_37_ in (d_926_v36_).Elements:
                        d_929_v35_: int = compr_37_
                        if (d_929_v35_) in (d_926_v36_):
                            coll37_[default__.safeModuloInt(d_929_v35_, d_884_v0_)] = d_884_v0_
                    return _dafny.Map(coll37_)
                nw132_[int(3)] = iife69_()
                
                nw132_[int(4)] = d_925_v34_
                nw132_[int(5)] = d_925_v34_
                nw132_[int(6)] = (d_925_v34_) | (_dafny.Map({d_884_v0_: d_884_v0_}))
                nw132_[int(7)] = (D11_DC24(d_925_v34_)).cf37
                nw132_[int(8)] = d_925_v34_
                nw132_[int(9)] = d_925_v34_
                nw132_[int(10)] = (d_925_v34_) | (d_925_v34_)
                nw132_[int(11)] = default__.fm25(globalState)
                nw132_[int(12)] = (default__.fm25(globalState)) | (((d_927_v37_)[d_884_v0_] if (d_884_v0_) in (d_927_v37_) else d_925_v34_))
                nw132_[int(13)] = d_925_v34_
                def iife70_():
                    coll38_ = _dafny.Map()
                    compr_38_: int
                    for compr_38_ in _dafny.IntegerRange(924, 0):
                        d_930_v38_: int = compr_38_
                        if ((924) <= (d_930_v38_)) and ((d_930_v38_) < (0)):
                            coll38_[default__.safeDivisionInt(d_930_v38_, len((d_923_v32_).f12))] = d_884_v0_
                    return _dafny.Map(coll38_)
                nw132_[int(14)] = (default__.fm25(globalState)) | (iife70_()
                )
                nw132_[int(15)] = _dafny.Map({d_884_v0_: d_884_v0_})
                nw132_[int(16)] = d_925_v34_
                nw132_[int(17)] = d_925_v34_
                nw132_[int(18)] = d_925_v34_
                nw132_[int(19)] = d_925_v34_
                d_928_v39_ = nw132_
                d_928_v39_ = d_928_v39_
                d_921_v30_ = (d_923_v32_).f12
                d_931_v40_: D10
                d_931_v40_ = D10_DC23()
                d_931_v40_ = d_931_v40_
            elif True:
                d_932_v41_: D0
                d_932_v41_ = D0_DC2((self).f4, d_884_v0_, d_884_v0_)
                d_884_v0_ = (d_932_v41_).cf4
                d_933_v42_: C0
                nw133_ = C0()
                nw133_.ctor__()
                d_933_v42_ = nw133_
                d_934_v43_: str
                d_934_v43_ = _dafny.CodePoint('g')
                d_935_v44_: _dafny.MultiSet
                d_935_v44_ = _dafny.MultiSet([self.f3])
                d_936_v45_: _dafny.Map
                d_936_v45_ = _dafny.Map({(_dafny.Set({_dafny.CodePoint('u')})) | (_dafny.Set({d_934_v43_, d_934_v43_})): (d_935_v44_) | (d_935_v44_)})
                d_936_v45_ = d_936_v45_
                d_937_v46_: _dafny.Array
                def lambda44_(d_938_i3_):
                    return not((self).f11)

                init21_ = lambda44_
                nw134_ = _dafny.Array(None, 18)
                for i0_21_ in range(nw134_.length(0)):
                    nw134_[i0_21_] = init21_(i0_21_)
                d_937_v46_ = nw134_
                d_939_v47_: _dafny.MultiSet
                d_939_v47_ = _dafny.MultiSet([d_937_v46_])
                (globalState).f0 = (d_939_v47_) == (d_939_v47_)
                d_940_v48_: C0
                nw135_ = C0()
                nw135_.ctor__()
                d_940_v48_ = nw135_
            d_941_v49_: D4
            d_941_v49_ = D4_DC12(d_884_v0_, 950, (self).f4)
            pat_let_tv34_ = d_884_v0_
            def iife71_(_pat_let16_0):
                def iife72_(d_942_dt__update__tmp_h1_):
                    def iife73_(_pat_let17_0):
                        def iife74_(d_943_dt__update_hcf20_h0_):
                            def iife75_(_pat_let18_0):
                                def iife76_(d_944_dt__update_hcf19_h0_):
                                    return D4_DC12(d_944_dt__update_hcf19_h0_, d_943_dt__update_hcf20_h0_, (d_942_dt__update__tmp_h1_).cf21)
                                return iife76_(_pat_let18_0)
                            return iife75_(len(_dafny.Map({(self).f10: (self).f10})))
                        return iife74_(_pat_let17_0)
                    return iife73_(pat_let_tv34_)
                return iife72_(_pat_let16_0)
            (globalState).f0 = ((self).f10) or (((self).f10) == (not((iife71_(d_941_v49_)).cf21)))
            d_945_v50_: _dafny.Array
            nw136_ = _dafny.Array(int(0), 17)
            d_945_v50_ = nw136_
            index139_ = default__.safeIndex(246, (d_945_v50_).length(0))
            (d_945_v50_)[index139_] = len(d_885_v1_)
            d_946_v51_: _dafny.MultiSet
            d_946_v51_ = _dafny.MultiSet([len(d_886_v2_)])
            index140_ = default__.safeIndex(246, (d_945_v50_).length(0))
            (d_945_v50_)[index140_] = (default__.safeModuloInt((d_946_v51_).cardinality, 547)) - (d_884_v0_)
        elif True:
            d_947_v52_: C0
            nw137_ = C0()
            nw137_.ctor__()
            d_947_v52_ = nw137_
            d_948_v53_: _dafny.Seq
            d_948_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptxas"))
            d_949_v54_: _dafny.Seq
            d_949_v54_ = _dafny.SeqWithoutIsStrInference([d_948_v53_, d_948_v53_])
            d_950_v55_: _dafny.MultiSet
            d_950_v55_ = _dafny.MultiSet([d_948_v53_, d_948_v53_, d_948_v53_, (d_948_v53_).set(default__.safeIndex(d_884_v0_, len(d_948_v53_)), _dafny.CodePoint('w')), d_948_v53_])
            d_951_v56_: _dafny.Map
            d_951_v56_ = _dafny.Map({(d_948_v53_).set(default__.safeIndex(len(d_949_v54_), len(d_948_v53_)), _dafny.CodePoint('j')): (d_950_v55_).cardinality})
            d_951_v56_ = (d_951_v56_).set(d_948_v53_, d_884_v0_)
            (globalState).f0 = (self).f11
            d_952_v57_: _dafny.Map
            d_952_v57_ = _dafny.Map({d_884_v0_: self.f3})
            d_953_v58_: _dafny.Seq
            d_953_v58_ = _dafny.SeqWithoutIsStrInference([(self).f4, (self).f11, (self).f10, (self).f10])
            d_954_v59_: D1
            d_954_v59_ = D1_DC4(len(d_948_v53_))
            d_955_v60_: _dafny.Set
            d_955_v60_ = _dafny.Set({not((self).f11)})
            d_956_v61_: D12
            d_956_v61_ = D12_DC26(d_955_v60_)
            d_957_v62_: _dafny.Seq
            d_957_v62_ = _dafny.SeqWithoutIsStrInference([d_884_v0_])
            d_958_v64_: _dafny.Map
            d_958_v64_ = _dafny.Map({d_884_v0_: len(d_948_v53_)})
            d_959_v65_: _dafny.MultiSet
            d_959_v65_ = _dafny.MultiSet([d_958_v64_])
            d_960_v66_: _dafny.Array
            nw138_ = _dafny.Array(None, 22)
            nw138_[int(0)] = ((d_952_v57_)[len(d_948_v53_)] if (len(d_948_v53_)) in (d_952_v57_) else (self).f4)
            nw138_[int(1)] = (not(self.f3)) in (d_953_v58_)
            nw138_[int(2)] = (False) in (d_953_v58_)
            nw138_[int(3)] = (self.f3) not in (d_953_v58_)
            nw138_[int(4)] = (self).f4
            nw138_[int(5)] = True
            nw138_[int(6)] = ((d_954_v59_).cf7) != (d_884_v0_)
            nw138_[int(7)] = (False) == ((self).f10)
            nw138_[int(8)] = not((_dafny.Set({False, self.f3, (self).f4, self.f3, (self).f11})).ispropersubset((d_956_v61_).cf39))
            nw138_[int(9)] = (self).f11
            nw138_[int(10)] = (self.f3 if (self).f11 else (self).f10)
            nw138_[int(11)] = (not((self).f10) if (self).f4 else False)
            nw138_[int(12)] = ((0) - ((d_957_v62_)[default__.safeIndex(642, len(d_957_v62_))])) <= (233)
            def iife77_():
                coll39_ = _dafny.Set()
                compr_39_: int
                for compr_39_ in _dafny.IntegerRange(403, -422):
                    d_961_v63_: int = compr_39_
                    if ((403) <= (d_961_v63_)) and ((d_961_v63_) < (-422)):
                        coll39_ = coll39_.union(_dafny.Set([(d_961_v63_) * (d_884_v0_)]))
                return _dafny.Set(coll39_)
            nw138_[int(13)] = not((iife77_()
            ).isdisjoint(d_885_v1_))
            nw138_[int(14)] = (self).f10
            nw138_[int(15)] = (not((self).f4) if (self).f4 else (self).f10)
            nw138_[int(16)] = (self).f10
            nw138_[int(17)] = (self).f11
            nw138_[int(18)] = (d_884_v0_) <= (d_884_v0_)
            nw138_[int(19)] = (_dafny.Set({default__.fm21(-514, d_959_v65_, globalState)})).issubset(_dafny.Set({(self).f11}))
            nw138_[int(20)] = (self).f10
            nw138_[int(21)] = (_dafny.SeqWithoutIsStrInference([(self).f11])) != (_dafny.SeqWithoutIsStrInference([(self).f4, (self).f4]))
            d_960_v66_ = nw138_
            index141_ = default__.safeIndex(264, (d_960_v66_).length(0))
            (d_960_v66_)[index141_] = not((self).f4)
            index142_ = default__.safeIndex(264, (d_960_v66_).length(0))
            (d_960_v66_)[index142_] = (self).f10
            d_962_v67_: _dafny.Array
            nw139_ = _dafny.Array(D2.default()(), 18)
            d_962_v67_ = nw139_
            index143_ = default__.safeIndex(26, (d_962_v67_).length(0))
            (d_962_v67_)[index143_] = D2_DC8(d_960_v66_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iuu")))
            d_963_v68_: str
            d_963_v68_ = _dafny.CodePoint('v')
            d_964_v69_: D2
            d_964_v69_ = D2_DC8(d_960_v66_, (d_947_v52_).fm11(d_963_v68_, d_955_v60_, globalState))
            index144_ = default__.safeIndex(26, (d_962_v67_).length(0))
            (d_962_v67_)[index144_] = d_964_v69_
        hi5_ = (self).fm0(globalState)
        for d_965_i4_ in range(d_884_v0_, hi5_):
            d_966_v70_: _dafny.Array
            def lambda45_(d_967_i5_):
                return (self).f10

            init22_ = lambda45_
            nw140_ = _dafny.Array(None, 25)
            for i0_22_ in range(nw140_.length(0)):
                nw140_[i0_22_] = init22_(i0_22_)
            d_966_v70_ = nw140_
            d_968_v71_: _dafny.Set
            d_968_v71_ = _dafny.Set({d_966_v70_})
            rhs129_ = True
            rhs130_ = (d_968_v71_) | (d_968_v71_)
            lhs89_ = globalState
            lhs89_.f0 = rhs129_
            d_968_v71_ = rhs130_
            d_884_v0_ = (d_965_i4_) * (d_884_v0_)
            index145_ = default__.safeIndex(461, (d_966_v70_).length(0))
            (d_966_v70_)[index145_] = (self).f10
            d_969_v72_: _dafny.Seq
            d_969_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhy"))
            index146_ = default__.safeIndex(461, (d_966_v70_).length(0))
            (d_966_v70_)[index146_] = ((self).f11 if (self).f4 else (len(d_969_v72_)) >= ((0) - (d_965_i4_)))
            index147_ = default__.safeIndex(461, (d_966_v70_).length(0))
            (d_966_v70_)[index147_] = (d_966_v70_)[default__.safeIndex(461, (d_966_v70_).length(0))]
        (globalState).f0 = (self).f4
        d_884_v0_ = default__.safeDivisionInt(d_884_v0_, d_884_v0_)
        d_884_v0_ = (0) - (default__.safeModuloInt((d_884_v0_) - (d_884_v0_), d_884_v0_))
        d_970_v73_: _dafny.Seq
        d_970_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hd"))
        d_971_v74_: _dafny.Array
        nw141_ = _dafny.Array(None, 7)
        nw141_[int(0)] = d_884_v0_
        nw141_[int(1)] = 989
        nw141_[int(2)] = len(d_970_v73_)
        nw141_[int(3)] = d_884_v0_
        nw141_[int(4)] = default__.safeDivisionInt(d_884_v0_, d_884_v0_)
        nw141_[int(5)] = d_884_v0_
        nw141_[int(6)] = d_884_v0_
        d_971_v74_ = nw141_
        index148_ = default__.safeIndex(642, (d_971_v74_).length(0))
        (d_971_v74_)[index148_] = default__.safeModuloInt(d_884_v0_, d_884_v0_)
        index149_ = default__.safeIndex(642, (d_971_v74_).length(0))
        (d_971_v74_)[index149_] = (0) - (d_884_v0_)

    @property
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11

class C3(T1, T0):
    def  __init__(self):
        self._f3: bool = False
        self._f4: bool = False
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
    def f4(self):
        return self._f4
    def ctor__(self, f3, f4):
        (self).f3 = f3
        (self)._f4 = f4

    def fm0(self, globalState):
        return (len(_dafny.Map({_dafny.Map({len(_dafny.Map({(0) - ((_dafny.MultiSet([False, self.f3])).cardinality): 314})): self.f3}): self.f3}))) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gett")))])), 850, len(_dafny.SeqWithoutIsStrInference([False]))])))

    def fm7(self, p0, p1, p2, p3, globalState):
        return (_dafny.Set({(self).f4})).intersection((_dafny.Set({self.f3})) | (_dafny.Set({self.f3})))

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        d_972_v0_: D0
        d_972_v0_ = D0_DC0(p0)
        d_973_v1_: str
        d_973_v1_ = _dafny.CodePoint('b')
        d_974_v2_: _dafny.MultiSet
        d_974_v2_ = _dafny.MultiSet([d_973_v1_])
        d_975_v3_: _dafny.Map
        d_975_v3_ = _dafny.Map({(d_974_v2_).cardinality: p0})
        d_976_v4_: _dafny.MultiSet
        d_976_v4_ = _dafny.MultiSet([(self).f4])
        d_977_v5_: _dafny.Map
        d_977_v5_ = _dafny.Map({(p1).f4: default__.fm8((d_976_v4_).cardinality, globalState)})
        d_978_v6_: _dafny.Seq
        d_978_v6_ = _dafny.SeqWithoutIsStrInference([self.f3, (self).f4, (self).f4])
        d_979_v7_: _dafny.Array
        nw142_ = _dafny.Array(None, 28)
        nw142_[int(0)] = d_973_v1_
        nw142_[int(1)] = d_973_v1_
        nw142_[int(2)] = d_973_v1_
        nw142_[int(3)] = d_973_v1_
        nw142_[int(4)] = d_973_v1_
        nw142_[int(5)] = d_973_v1_
        nw142_[int(6)] = d_973_v1_
        nw142_[int(7)] = d_973_v1_
        nw142_[int(8)] = d_973_v1_
        nw142_[int(9)] = _dafny.CodePoint('u')
        nw142_[int(10)] = d_973_v1_
        nw142_[int(11)] = _dafny.CodePoint('h')
        nw142_[int(12)] = d_973_v1_
        nw142_[int(13)] = default__.fm8(589, globalState)
        nw142_[int(14)] = default__.fm8(p0, globalState)
        nw142_[int(15)] = ((d_977_v5_)[(self).f4] if ((self).f4) in (d_977_v5_) else d_973_v1_)
        nw142_[int(16)] = d_973_v1_
        nw142_[int(17)] = d_973_v1_
        nw142_[int(18)] = d_973_v1_
        nw142_[int(19)] = _dafny.CodePoint('s')
        nw142_[int(20)] = d_973_v1_
        nw142_[int(21)] = _dafny.CodePoint('i')
        nw142_[int(22)] = d_973_v1_
        nw142_[int(23)] = d_973_v1_
        nw142_[int(24)] = d_973_v1_
        nw142_[int(25)] = d_973_v1_
        nw142_[int(26)] = d_973_v1_
        nw142_[int(27)] = (default__.fm9(p0, d_978_v6_, p1.f3, p0, globalState)).cf2
        d_979_v7_ = nw142_
        d_980_v8_: _dafny.Map
        d_980_v8_ = _dafny.Map({((d_975_v3_)[816] if (816) in (d_975_v3_) else p0): d_979_v7_})
        hi6_ = len(d_980_v8_)
        for d_981_i0_ in range(((d_972_v0_).cf0) * ((self).fm0(globalState)), hi6_):
            r0 = ((0) - ((0) - (p0))) * (p0)
            r0 = p0
            (globalState).f0 = False
            d_982_v9_: _dafny.Seq
            d_982_v9_ = _dafny.SeqWithoutIsStrInference([d_973_v1_, d_973_v1_])
            d_983_v10_: _dafny.Map
            d_983_v10_ = _dafny.Map({d_982_v9_: (p1).f4})
            d_984_v11_: _dafny.Array
            nw143_ = _dafny.Array(int(0), 24)
            d_984_v11_ = nw143_
            d_985_v12_: _dafny.Map
            d_985_v12_ = _dafny.Map({d_984_v11_: d_981_i0_})
            d_986_v13_: _dafny.Seq
            d_986_v13_ = _dafny.SeqWithoutIsStrInference([p0, len(d_985_v12_)])
            if ((d_983_v10_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_973_v1_ for d_987_i1_ in range(default__.abs(0))]): (p1).f4}))) or ((len(d_982_v9_)) <= ((d_986_v13_)[default__.safeIndex((0) - (d_981_i0_), len(d_986_v13_))])):
                d_988_v14_: _dafny.Array
                def lambda46_(d_989_p1_):
                    def lambda47_(d_990_i2_):
                        return (d_989_p1_).f4

                    return lambda47_

                init23_ = lambda46_(p1)
                nw144_ = _dafny.Array(None, 19)
                for i0_23_ in range(nw144_.length(0)):
                    nw144_[i0_23_] = init23_(i0_23_)
                d_988_v14_ = nw144_
                d_988_v14_ = d_988_v14_
                d_991_v15_: D0
                d_991_v15_ = D0_DC2((p1.f3 if False else self.f3), (d_976_v4_).cardinality, p0)
                d_991_v15_ = D0_DC2((p1).f4, d_981_i0_, default__.safeDivisionInt(p0, p0))
                d_991_v15_ = d_991_v15_
                d_992_v16_: _dafny.Seq
                d_992_v16_ = _dafny.SeqWithoutIsStrInference([d_972_v0_])
                d_993_v17_: _dafny.Map
                d_993_v17_ = _dafny.Map({d_981_i0_: d_992_v16_})
                d_994_v19_: _dafny.Map
                d_994_v19_ = _dafny.Map({p0: (self).f4})
                def iife78_():
                    coll40_ = _dafny.Map()
                    compr_40_: int
                    for compr_40_ in (d_994_v19_).keys.Elements:
                        d_995_v18_: int = compr_40_
                        if (d_995_v18_) in (d_994_v19_):
                            coll40_[(d_995_v18_) * (p0)] = d_992_v16_
                    return _dafny.Map(coll40_)
                rhs131_ = p0
                rhs132_ = (iife78_()
                 if (p1).f4 else (d_993_v17_) | (d_993_v17_))
                r0 = rhs131_
                d_993_v17_ = rhs132_
                d_996_v20_: _dafny.MultiSet
                d_996_v20_ = _dafny.MultiSet([p0])
                index150_ = default__.safeIndex(186, (d_984_v11_).length(0))
                (d_984_v11_)[index150_] = (0) - (((d_996_v20_)[p0] if (p0) in (d_996_v20_) else p0))
                index151_ = default__.safeIndex(186, (d_984_v11_).length(0))
                (d_984_v11_)[index151_] = (d_991_v15_).cf5
            elif True:
                d_997_v21_: _dafny.Array
                nw145_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_997_v21_ = nw145_
                d_998_v22_: _dafny.Array
                nw146_ = _dafny.Array(False, 24)
                d_998_v22_ = nw146_
                index152_ = default__.safeIndex(891, (d_997_v21_).length(0))
                (d_997_v21_)[index152_] = d_998_v22_
                d_999_v23_: _dafny.Map
                d_999_v23_ = _dafny.Map({p1.f3: d_981_i0_})
                d_1000_v24_: _dafny.MultiSet
                d_1000_v24_ = _dafny.MultiSet([d_999_v23_])
                index153_ = default__.safeIndex(891, (d_997_v21_).length(0))
                rhs133_ = True
                rhs134_ = d_998_v22_
                rhs135_ = (d_1000_v24_).ispropersubset(d_1000_v24_)
                rhs136_ = (p1).f4
                rhs137_ = d_998_v22_
                lhs90_ = p1
                lhs91_ = d_997_v21_
                lhs92_ = default__.safeIndex(891, (d_997_v21_).length(0))
                lhs93_ = p1
                lhs94_ = globalState
                lhs90_.f3 = rhs133_
                lhs91_[lhs92_] = rhs134_
                lhs93_.f3 = rhs135_
                lhs94_.f0 = rhs136_
                d_998_v22_ = rhs137_
                d_1001_v25_: _dafny.Map
                d_1001_v25_ = _dafny.Map({(self).f4: (p1).f4})
                (globalState).f0 = ((len(d_978_v6_)) + (len(d_1001_v25_))) < ((d_981_i0_) + (d_981_i0_))
                d_983_v10_ = (d_983_v10_).set((d_982_v9_) + (d_982_v9_), ((p1).f4) == (p1.f3))
                d_1002_v26_: _dafny.Set
                d_1002_v26_ = _dafny.Set({len(d_986_v13_)})
                (globalState).f0 = ((len(d_982_v9_)) + (len(d_1002_v26_))) != ((p1).fm0(globalState))
                r0 = d_981_i0_
        if (p1).f4:
            d_1003_v27_: _dafny.Array
            nw147_ = _dafny.Array(False, 9)
            d_1003_v27_ = nw147_
            d_1004_v28_: _dafny.Map
            d_1004_v28_ = _dafny.Map({p1.f3: p0})
            d_1005_v29_: _dafny.Set
            d_1005_v29_ = _dafny.Set({d_1004_v28_})
            d_1006_v31_: _dafny.Map
            def iife79_():
                coll41_ = _dafny.Set()
                compr_41_: _dafny.Map
                for compr_41_ in (_dafny.SeqWithoutIsStrInference([d_1004_v28_, d_1004_v28_, d_1004_v28_])).Elements:
                    d_1007_v30_: _dafny.Map = compr_41_
                    if (d_1007_v30_) in (_dafny.SeqWithoutIsStrInference([d_1004_v28_, d_1004_v28_, d_1004_v28_])):
                        coll41_ = coll41_.union(_dafny.Set([d_1007_v30_]))
                return _dafny.Set(coll41_)
            d_1006_v31_ = _dafny.Map({(d_1005_v29_) - (iife79_()
            ): d_1003_v27_})
            d_1003_v27_ = ((d_1006_v31_)[(d_1005_v29_).intersection(d_1005_v29_)] if ((d_1005_v29_).intersection(d_1005_v29_)) in (d_1006_v31_) else d_1003_v27_)
            d_1008_v32_: _dafny.Set
            d_1008_v32_ = _dafny.Set({470, p0})
            rhs138_ = (D0_DC0(p0)).cf0
            rhs139_ = d_973_v1_
            rhs140_ = (d_1008_v32_).ispropersubset(d_1008_v32_)
            rhs141_ = (0) - (p0)
            lhs95_ = p1
            r0 = rhs138_
            d_973_v1_ = rhs139_
            lhs95_.f3 = rhs140_
            r0 = rhs141_
            if (self).f4:
                d_1009_v33_: _dafny.Array
                def lambda48_(d_1010_p0_):
                    def lambda49_(d_1011_i3_):
                        return (d_1011_i3_) - (d_1010_p0_)

                    return lambda49_

                init24_ = lambda48_(p0)
                nw148_ = _dafny.Array(None, 23)
                for i0_24_ in range(nw148_.length(0)):
                    nw148_[i0_24_] = init24_(i0_24_)
                d_1009_v33_ = nw148_
                d_1012_v34_: _dafny.Seq
                d_1012_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yowfhhl"))
                index154_ = default__.safeIndex(766, (d_1009_v33_).length(0))
                (d_1009_v33_)[index154_] = len(d_1012_v34_)
                pat_let_tv35_ = p0
                d_1013_v35_: _dafny.Map
                def iife80_(_pat_let19_0):
                    def iife81_(d_1014_dt__update__tmp_h0_):
                        def iife82_(_pat_let20_0):
                            def iife83_(d_1015_dt__update_hcf4_h0_):
                                return D0_DC2((d_1014_dt__update__tmp_h0_).cf3, d_1015_dt__update_hcf4_h0_, (d_1014_dt__update__tmp_h0_).cf5)
                            return iife83_(_pat_let20_0)
                        return iife82_((D0_DC0(pat_let_tv35_)).cf0)
                    return iife81_(_pat_let19_0)
                d_1013_v35_ = _dafny.Map({(p1).f4: (iife80_(D0_DC2((self).f4, p0, p0))).cf3})
                d_1016_v36_: _dafny.Seq
                d_1016_v36_ = _dafny.SeqWithoutIsStrInference([d_1013_v35_])
                index155_ = default__.safeIndex(766, (d_1009_v33_).length(0))
                rhs142_ = True
                rhs143_ = (0) - (len(((d_1016_v36_) + (d_1016_v36_)) + (((d_1016_v36_).set(default__.safeIndex(890, len(d_1016_v36_)), d_1013_v35_)).set(default__.safeIndex(p0, len((d_1016_v36_).set(default__.safeIndex(890, len(d_1016_v36_)), d_1013_v35_))), d_1013_v35_))))
                lhs96_ = globalState
                lhs97_ = d_1009_v33_
                lhs98_ = default__.safeIndex(766, (d_1009_v33_).length(0))
                lhs96_.f0 = rhs142_
                lhs97_[lhs98_] = rhs143_
                d_1017_v37_: D0
                d_1017_v37_ = D0_DC2((_dafny.CodePoint('u')) not in (d_1012_v34_), -262, ((d_976_v4_).cardinality) + (len(_dafny.SeqWithoutIsStrInference([d_1012_v34_]))))
                d_1018_v38_: _dafny.Map
                d_1018_v38_ = _dafny.Map({90: p1.f3})
                d_1017_v37_ = D0_DC2(((d_1018_v38_)[449] if (449) in (d_1018_v38_) else (p1).f4), ((d_1009_v33_)[default__.safeIndex(766, (d_1009_v33_).length(0))]) + (len(d_1012_v34_)), (self).fm0(globalState))
                d_1019_v39_: _dafny.Set
                d_1019_v39_ = _dafny.Set({(p1).f4})
                d_1020_v40_: _dafny.MultiSet
                d_1020_v40_ = _dafny.MultiSet([d_1019_v39_])
                d_1020_v40_ = _dafny.MultiSet([(d_1019_v39_).intersection(d_1019_v39_), d_1019_v39_, d_1019_v39_, d_1019_v39_])
                d_1021_v41_: _dafny.Map
                d_1021_v41_ = _dafny.Map({((d_976_v4_)[p1.f3] if (p1.f3) in (d_976_v4_) else (d_1009_v33_)[default__.safeIndex(766, (d_1009_v33_).length(0))]): d_973_v1_})
                d_1022_v42_: _dafny.Map
                d_1022_v42_ = _dafny.Map({(374) + ((d_1009_v33_)[default__.safeIndex(766, (d_1009_v33_).length(0))]): d_1021_v41_})
                d_1022_v42_ = d_1022_v42_
                d_1023_v43_: C0
                nw149_ = C0()
                nw149_.ctor__()
                d_1023_v43_ = nw149_
            elif True:
                d_1024_v44_: _dafny.Seq
                d_1024_v44_ = _dafny.SeqWithoutIsStrInference([d_1003_v27_, d_1003_v27_])
                d_1025_v45_: _dafny.Map
                d_1025_v45_ = _dafny.Map({p0: d_1003_v27_})
                d_1026_v46_: _dafny.Array
                nw150_ = _dafny.Array(None, 25)
                nw150_[int(0)] = d_1003_v27_
                nw150_[int(1)] = d_1003_v27_
                nw150_[int(2)] = d_1003_v27_
                nw150_[int(3)] = d_1003_v27_
                nw150_[int(4)] = d_1003_v27_
                nw150_[int(5)] = d_1003_v27_
                nw150_[int(6)] = d_1003_v27_
                nw150_[int(7)] = d_1003_v27_
                nw150_[int(8)] = d_1003_v27_
                nw150_[int(9)] = d_1003_v27_
                nw150_[int(10)] = d_1003_v27_
                nw150_[int(11)] = d_1003_v27_
                nw150_[int(12)] = d_1003_v27_
                nw150_[int(13)] = (d_1024_v44_)[default__.safeIndex(p0, len(d_1024_v44_))]
                nw150_[int(14)] = d_1003_v27_
                nw150_[int(15)] = ((d_1025_v45_)[len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "juonh")): (p1).f4}))] if (len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "juonh")): (p1).f4}))) in (d_1025_v45_) else d_1003_v27_)
                nw150_[int(16)] = d_1003_v27_
                nw150_[int(17)] = d_1003_v27_
                nw150_[int(18)] = d_1003_v27_
                nw150_[int(19)] = d_1003_v27_
                nw150_[int(20)] = d_1003_v27_
                nw150_[int(21)] = d_1003_v27_
                nw150_[int(22)] = d_1003_v27_
                nw150_[int(23)] = d_1003_v27_
                nw150_[int(24)] = d_1003_v27_
                d_1026_v46_ = nw150_
                index156_ = default__.safeIndex(420, (d_1026_v46_).length(0))
                (d_1026_v46_)[index156_] = d_1003_v27_
                d_1027_v47_: _dafny.Seq
                d_1027_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mstiflvo"))
                d_1028_v48_: _dafny.Set
                d_1028_v48_ = _dafny.Set({d_1027_v47_, d_1027_v47_, d_1027_v47_})
                d_1029_v49_: _dafny.Set
                d_1029_v49_ = _dafny.Set({(p1).f4})
                d_1030_v50_: _dafny.Seq
                d_1030_v50_ = _dafny.SeqWithoutIsStrInference([d_1029_v49_])
                d_1031_v51_: _dafny.Seq
                d_1031_v51_ = _dafny.SeqWithoutIsStrInference([p0, p0, len(d_1030_v50_), default__.safeDivisionInt(p0, p0)])
                index157_ = default__.safeIndex(420, (d_1026_v46_).length(0))
                rhs144_ = d_1003_v27_
                rhs145_ = (p0) * (705)
                rhs146_ = d_1028_v48_
                rhs147_ = len(d_1031_v51_)
                lhs99_ = d_1026_v46_
                lhs100_ = default__.safeIndex(420, (d_1026_v46_).length(0))
                lhs99_[lhs100_] = rhs144_
                r0 = rhs145_
                d_1028_v48_ = rhs146_
                r0 = rhs147_
                d_1032_v52_: _dafny.Map
                d_1032_v52_ = _dafny.Map({d_1027_v47_: p0})
                d_1033_v55_: _dafny.Array
                def lambda50_(d_1034_p1_, d_1035_v2_, d_1036_p0_):
                    def lambda51_(d_1037_i4_):
                        def iife84_():
                            coll42_ = _dafny.Map()
                            compr_42_: str
                            for compr_42_ in (d_1035_v2_).Elements:
                                d_1038_v53_: str = compr_42_
                                if (d_1038_v53_) in (d_1035_v2_):
                                    coll42_[d_1038_v53_] = d_1036_p0_
                            return _dafny.Map(coll42_)
                        def iife85_():
                            coll43_ = _dafny.Map()
                            compr_43_: str
                            for compr_43_ in (d_1035_v2_).Elements:
                                d_1039_v54_: str = compr_43_
                                if (d_1039_v54_) in (d_1035_v2_):
                                    coll43_[d_1039_v54_] = d_1036_p0_
                            return _dafny.Map(coll43_)
                        return (iife84_()
                         if d_1034_p1_.f3 else iife85_()
                        )

                    return lambda51_

                init25_ = lambda50_(p1, d_974_v2_, p0)
                nw151_ = _dafny.Array(None, 18)
                for i0_25_ in range(nw151_.length(0)):
                    nw151_[i0_25_] = init25_(i0_25_)
                d_1033_v55_ = nw151_
                d_1040_v56_: _dafny.Map
                d_1040_v56_ = _dafny.Map({d_973_v1_: len(_dafny.SeqWithoutIsStrInference([p0, p0]))})
                index158_ = default__.safeIndex(250, (d_1033_v55_).length(0))
                (d_1033_v55_)[index158_] = d_1040_v56_
                d_1041_v57_: _dafny.Map
                d_1041_v57_ = _dafny.Map({(p1).f4: d_1027_v47_})
                d_1042_v58_: _dafny.MultiSet
                d_1042_v58_ = _dafny.MultiSet([len(d_978_v6_)])
                index159_ = default__.safeIndex(250, (d_1033_v55_).length(0))
                rhs148_ = d_1032_v52_
                rhs149_ = (_dafny.Map({d_973_v1_: p0})).set(d_973_v1_, p0)
                rhs150_ = (0) - (((d_1042_v58_)[p0] if (p0) in (d_1042_v58_) else p0))
                rhs151_ = (len((d_975_v3_ if (p1).f4 else d_975_v3_))) + ((self).fm0(globalState))
                rhs152_ = (d_1041_v57_ if p1.f3 else (d_1041_v57_) | (d_1041_v57_))
                lhs101_ = d_1033_v55_
                lhs102_ = default__.safeIndex(250, (d_1033_v55_).length(0))
                d_1032_v52_ = rhs148_
                lhs101_[lhs102_] = rhs149_
                r0 = rhs150_
                r0 = rhs151_
                d_1041_v57_ = rhs152_
                r0 = 385
                d_1004_v28_ = (d_1004_v28_).set((self).f4, (p0 if (self).f4 else (d_1031_v51_)[default__.safeIndex(p0, len(d_1031_v51_))]))
                (globalState).f0 = self.f3
            d_1043_v59_: D0
            d_1043_v59_ = D0_DC1((self).f4, _dafny.CodePoint('t'))
            d_1044_v60_: _dafny.Map
            d_1044_v60_ = _dafny.Map({p0: d_1043_v59_})
            d_1044_v60_ = d_1044_v60_
            d_1045_v61_: _dafny.Array
            def lambda52_(d_1046_p0_):
                def lambda53_(d_1047_i5_):
                    return default__.safeDivisionInt(d_1047_i5_, d_1046_p0_)

                return lambda53_

            init26_ = lambda52_(p0)
            nw152_ = _dafny.Array(None, 21)
            for i0_26_ in range(nw152_.length(0)):
                nw152_[i0_26_] = init26_(i0_26_)
            d_1045_v61_ = nw152_
            d_1048_v62_: _dafny.Seq
            d_1048_v62_ = _dafny.SeqWithoutIsStrInference([d_1003_v27_, d_1003_v27_, d_1003_v27_])
            index160_ = default__.safeIndex(588, (d_1045_v61_).length(0))
            (d_1045_v61_)[index160_] = len(d_1048_v62_)
            index161_ = default__.safeIndex(588, (d_1045_v61_).length(0))
            (d_1045_v61_)[index161_] = (382) * (len((default__.fm12(p0, self.f3, p0, globalState)).set(d_1043_v59_, (self).f4)))
        elif True:
            r0 = (p1).fm0(globalState)
            d_1049_v64_: _dafny.Array
            def lambda54_(d_1050_p0_, d_1051_v1_):
                def lambda55_(d_1052_i6_):
                    def iife86_():
                        coll44_ = _dafny.Set()
                        compr_44_: int
                        for compr_44_ in _dafny.IntegerRange(550, -61):
                            d_1053_v63_: int = compr_44_
                            if ((550) <= (d_1053_v63_)) and ((d_1053_v63_) < (-61)):
                                coll44_ = coll44_.union(_dafny.Set([(d_1053_v63_) + (len(_dafny.SeqWithoutIsStrInference([d_1051_v1_ for d_1054_i7_ in range(default__.abs(269))])))]))
                        return _dafny.Set(coll44_)
                    return (_dafny.Set({d_1050_p0_, d_1050_p0_})).isdisjoint(iife86_()
                    )

                return lambda55_

            init27_ = lambda54_(p0, d_973_v1_)
            nw153_ = _dafny.Array(None, 20)
            for i0_27_ in range(nw153_.length(0)):
                nw153_[i0_27_] = init27_(i0_27_)
            d_1049_v64_ = nw153_
            d_1055_v65_: _dafny.MultiSet
            d_1055_v65_ = _dafny.MultiSet([d_972_v0_])
            d_1056_v66_: _dafny.Set
            d_1056_v66_ = _dafny.Set({508, p0})
            index162_ = default__.safeIndex(805, (d_1049_v64_).length(0))
            (d_1049_v64_)[index162_] = (d_1056_v66_).ispropersubset(_dafny.Set({(d_1055_v65_).cardinality}))
            d_1057_v67_: _dafny.Map
            d_1057_v67_ = _dafny.Map({p0: (p1).f4})
            index163_ = default__.safeIndex(805, (d_1049_v64_).length(0))
            (d_1049_v64_)[index163_] = not((((d_1057_v67_)[302] if (302) in (d_1057_v67_) else (p1).f4)) and (False))
            d_1058_v68_: _dafny.Seq
            d_1058_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqqftlo"))
            d_1059_v69_: _dafny.Map
            d_1059_v69_ = _dafny.Map({d_1058_v68_: p0})
            d_1059_v69_ = default__.fm13((p0) * (p0), (0) - (len(d_975_v3_)), p0, (self).f4, globalState)
            (p1).f3 = p1.f3
            index164_ = default__.safeIndex(805, (d_1049_v64_).length(0))
            (d_1049_v64_)[index164_] = (d_978_v6_)[default__.safeIndex(((d_976_v4_)[False] if (False) in (d_976_v4_) else p0), len(d_978_v6_))]
        hi7_ = p0
        for d_1060_i8_ in range((p0 if self.f3 else -876), hi7_):
            d_1061_v70_: _dafny.Array
            nw154_ = _dafny.Array(int(0), 1)
            d_1061_v70_ = nw154_
            index165_ = default__.safeIndex(710, (d_1061_v70_).length(0))
            (d_1061_v70_)[index165_] = default__.safeModuloInt(p0, len(d_975_v3_))
            index166_ = default__.safeIndex(710, (d_1061_v70_).length(0))
            (d_1061_v70_)[index166_] = default__.safeModuloInt(d_1060_i8_, (d_1060_i8_) * (p0))
            index167_ = default__.safeIndex(710, (d_1061_v70_).length(0))
            (d_1061_v70_)[index167_] = p0
            d_1062_v71_: C0
            nw155_ = C0()
            nw155_.ctor__()
            d_1062_v71_ = nw155_
            index168_ = default__.safeIndex(710, (d_1061_v70_).length(0))
            (d_1061_v70_)[index168_] = (d_974_v2_).cardinality
        d_1063_v72_: _dafny.Array
        def lambda56_(d_1064_p0_):
            def lambda57_(d_1065_i9_):
                return (d_1065_i9_) + (d_1064_p0_)

            return lambda57_

        init28_ = lambda56_(p0)
        nw156_ = _dafny.Array(None, 14)
        for i0_28_ in range(nw156_.length(0)):
            nw156_[i0_28_] = init28_(i0_28_)
        d_1063_v72_ = nw156_
        d_1066_v73_: _dafny.Map
        d_1066_v73_ = _dafny.Map({(p1).f4: d_1063_v72_})
        d_1067_v74_: _dafny.Array
        nw157_ = _dafny.Array(None, 17)
        nw157_[int(0)] = d_1063_v72_
        nw157_[int(1)] = d_1063_v72_
        nw157_[int(2)] = d_1063_v72_
        nw157_[int(3)] = d_1063_v72_
        nw157_[int(4)] = d_1063_v72_
        nw157_[int(5)] = d_1063_v72_
        nw157_[int(6)] = (d_1063_v72_ if (p1).f4 else d_1063_v72_)
        nw157_[int(7)] = d_1063_v72_
        nw157_[int(8)] = d_1063_v72_
        nw157_[int(9)] = d_1063_v72_
        nw157_[int(10)] = d_1063_v72_
        nw157_[int(11)] = (d_1063_v72_ if self.f3 else d_1063_v72_)
        nw157_[int(12)] = d_1063_v72_
        nw157_[int(13)] = ((d_1066_v73_)[(self).f4] if ((self).f4) in (d_1066_v73_) else d_1063_v72_)
        nw157_[int(14)] = d_1063_v72_
        nw157_[int(15)] = d_1063_v72_
        nw157_[int(16)] = d_1063_v72_
        d_1067_v74_ = nw157_
        d_1068_v75_: _dafny.Map
        d_1068_v75_ = _dafny.Map({d_973_v1_: d_1063_v72_})
        index169_ = default__.safeIndex(145, (d_1067_v74_).length(0))
        (d_1067_v74_)[index169_] = ((d_1068_v75_)[d_973_v1_] if (d_973_v1_) in (d_1068_v75_) else d_1063_v72_)
        d_1069_v76_: _dafny.Seq
        d_1069_v76_ = _dafny.SeqWithoutIsStrInference([d_1063_v72_, d_1063_v72_, d_1063_v72_, d_1063_v72_, d_1063_v72_])
        d_1070_v77_: _dafny.Map
        d_1070_v77_ = _dafny.Map({True: (self).f4})
        index170_ = default__.safeIndex(145, (d_1067_v74_).length(0))
        (d_1067_v74_)[index170_] = (d_1069_v76_)[default__.safeIndex((p0) - (len(d_1070_v77_)), len(d_1069_v76_))]
        d_1071_v78_: _dafny.Seq
        d_1071_v78_ = _dafny.SeqWithoutIsStrInference([p0])
        hi8_ = p0
        for d_1072_i10_ in range((len(d_1071_v78_)) * (p0), hi8_):
            d_1073_v79_: _dafny.Seq
            d_1073_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lnyygcmow"))
            d_1074_v80_: _dafny.Seq
            d_1074_v80_ = _dafny.SeqWithoutIsStrInference([d_1073_v79_])
            d_1074_v80_ = d_1074_v80_
            d_1075_v81_: _dafny.Array
            nw158_ = _dafny.Array(False, 23)
            d_1075_v81_ = nw158_
            d_1076_v82_: _dafny.MultiSet
            d_1076_v82_ = _dafny.MultiSet([d_1075_v81_, d_1075_v81_])
            d_1076_v82_ = d_1076_v82_
            def lambda58_(d_1077_p1_, d_1078_v1_):
                def lambda59_(d_1079_i11_):
                    return not((D0_DC1((d_1077_p1_).f4, d_1078_v1_)).cf1)

                return lambda59_

            init29_ = lambda58_(p1, d_973_v1_)
            nw159_ = _dafny.Array(None, 24)
            for i0_29_ in range(nw159_.length(0)):
                nw159_[i0_29_] = init29_(i0_29_)
            d_1075_v81_ = nw159_
            (p1).f3 = self.f3
        d_1080_v83_: _dafny.Array
        nw160_ = _dafny.Array(False, 25)
        d_1080_v83_ = nw160_
        d_1081_v84_: _dafny.Seq
        d_1081_v84_ = _dafny.SeqWithoutIsStrInference([d_1080_v83_, d_1080_v83_])
        d_1082_v85_: _dafny.Seq
        d_1082_v85_ = _dafny.SeqWithoutIsStrInference([d_1081_v84_])
        d_1083_v86_: _dafny.Map
        d_1083_v86_ = _dafny.Map({p0: (self).f4})
        d_1081_v84_ = ((d_1082_v85_)[default__.safeIndex(len(d_1083_v86_), len(d_1082_v85_))] if (p0) >= (p0) else d_1081_v84_)
        r0 = p0
        return r0

    def m2(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1084_v0_: _dafny.Map
        d_1084_v0_ = _dafny.Map({p1.f3: self.f3})
        d_1085_v1_: D0
        d_1085_v1_ = D0_DC1(((d_1084_v0_)[default__.fm14(globalState)] if (default__.fm14(globalState)) in (d_1084_v0_) else not((self).f4)), p0)
        d_1085_v1_ = d_1085_v1_
        if self.f3:
            d_1086_v2_: _dafny.Seq
            d_1086_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ua"))
            d_1087_v3_: _dafny.Array
            nw161_ = _dafny.Array(None, 4)
            nw161_[int(0)] = self.f3
            nw161_[int(1)] = p1.f3
            nw161_[int(2)] = (self.f3) == ((self).f4)
            nw161_[int(3)] = (d_1086_v2_) < (d_1086_v2_)
            d_1087_v3_ = nw161_
            d_1087_v3_ = d_1087_v3_
            d_1088_v4_: int
            d_1088_v4_ = 902
            d_1089_v5_: _dafny.Seq
            d_1089_v5_ = _dafny.SeqWithoutIsStrInference([self.f3])
            d_1090_v6_: _dafny.Array
            nw162_ = _dafny.Array(None, 10)
            nw162_[int(0)] = len(d_1089_v5_)
            nw162_[int(1)] = d_1088_v4_
            nw162_[int(2)] = d_1088_v4_
            nw162_[int(3)] = 241
            nw162_[int(4)] = d_1088_v4_
            nw162_[int(5)] = d_1088_v4_
            nw162_[int(6)] = d_1088_v4_
            nw162_[int(7)] = 733
            nw162_[int(8)] = (0) - (-98)
            nw162_[int(9)] = d_1088_v4_
            d_1090_v6_ = nw162_
            d_1091_v7_: _dafny.Seq
            d_1091_v7_ = _dafny.SeqWithoutIsStrInference([d_1090_v6_])
            d_1092_v8_: _dafny.Set
            d_1092_v8_ = _dafny.Set({d_1088_v4_, len(d_1091_v7_)})
            (self).f3 = (d_1088_v4_) > (len(d_1092_v8_))
            d_1093_v9_: _dafny.Set
            d_1093_v9_ = _dafny.Set({_dafny.MultiSet([34, (0) - (d_1088_v4_)])})
            d_1093_v9_ = d_1093_v9_
            d_1094_v10_: C0
            nw163_ = C0()
            nw163_.ctor__()
            d_1094_v10_ = nw163_
            d_1095_v11_: C0
            nw164_ = C0()
            nw164_.ctor__()
            d_1095_v11_ = nw164_
        elif True:
            d_1096_v12_: int
            d_1096_v12_ = 123
            d_1096_v12_ = d_1096_v12_
            d_1097_v13_: _dafny.MultiSet
            d_1097_v13_ = _dafny.MultiSet([(p1).f4])
            d_1098_v14_: _dafny.Seq
            d_1098_v14_ = _dafny.SeqWithoutIsStrInference([d_1096_v12_, ((d_1097_v13_).set((p1).f4, default__.abs(d_1096_v12_))).cardinality, d_1096_v12_])
            d_1096_v12_ = (d_1098_v14_)[default__.safeIndex(default__.safeModuloInt(len(d_1098_v14_), (d_1098_v14_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p0 for d_1099_i0_ in range(default__.abs(698))])), len(d_1098_v14_))]), len(d_1098_v14_))]
            d_1100_v15_: _dafny.Array
            nw165_ = _dafny.Array(D0.default()(), 12)
            d_1100_v15_ = nw165_
            d_1100_v15_ = d_1100_v15_
            d_1101_v16_: _dafny.Map
            d_1101_v16_ = _dafny.Map({default__.fm14(globalState): d_1096_v12_})
            d_1101_v16_ = (d_1101_v16_).set((self).f4, d_1096_v12_)
            d_1102_v17_: _dafny.Array
            nw166_ = _dafny.Array(_dafny.Seq({}), 1)
            d_1102_v17_ = nw166_
            d_1103_v18_: _dafny.Seq
            d_1103_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1098_v14_])])
            index171_ = default__.safeIndex(261, (d_1102_v17_).length(0))
            (d_1102_v17_)[index171_] = (d_1103_v18_)[default__.safeIndex(721, len(d_1103_v18_))]
            d_1104_v19_: _dafny.Array
            nw167_ = _dafny.Array(False, 10)
            d_1104_v19_ = nw167_
            d_1105_v20_: D1
            d_1105_v20_ = D1_DC3(d_1104_v19_)
            d_1106_v21_: _dafny.Map
            d_1106_v21_ = _dafny.Map({(d_1105_v20_).cf6: (d_1096_v12_) - (d_1096_v12_)})
            index172_ = default__.safeIndex(261, (d_1102_v17_).length(0))
            rhs153_ = _dafny.SeqWithoutIsStrInference([(d_1098_v14_) + (d_1098_v14_) for d_1107_i1_ in range(default__.abs(453))])
            rhs154_ = len(d_1106_v21_)
            lhs103_ = d_1102_v17_
            lhs104_ = default__.safeIndex(261, (d_1102_v17_).length(0))
            lhs103_[lhs104_] = rhs153_
            d_1096_v12_ = rhs154_
        d_1108_i2_: int
        d_1108_i2_ = 0
        with _dafny.label("4"):
            while self.f3:
                with _dafny.c_label("4"):
                    if (d_1108_i2_) >= (100):
                        raise _dafny.Break("4")
                    d_1108_i2_ = (d_1108_i2_) + (1)
                    d_1109_v22_: _dafny.Array
                    nw168_ = _dafny.Array(int(0), 8)
                    d_1109_v22_ = nw168_
                    d_1110_v23_: int
                    d_1110_v23_ = 63
                    index173_ = default__.safeIndex(522, (d_1109_v22_).length(0))
                    (d_1109_v22_)[index173_] = d_1110_v23_
                    index174_ = default__.safeIndex(522, (d_1109_v22_).length(0))
                    (d_1109_v22_)[index174_] = ((d_1110_v23_) * (764)) * (d_1110_v23_)
                    (p1).f3 = (p1).f4
                    d_1110_v23_ = (p1).fm0(globalState)
                    index175_ = default__.safeIndex(522, (d_1109_v22_).length(0))
                    (d_1109_v22_)[index175_] = (p1).fm0(globalState)
                    pass
            pass
        if not((self).f4):
            d_1111_v24_: int
            d_1111_v24_ = 225
            d_1111_v24_ = d_1111_v24_
            d_1112_v25_: C0
            nw169_ = C0()
            nw169_.ctor__()
            d_1112_v25_ = nw169_
            d_1113_v26_: _dafny.Map
            d_1113_v26_ = _dafny.Map({(p1).f4: p0})
            d_1114_v27_: D0
            d_1114_v27_ = D0_DC2(True, len(d_1113_v26_), (0) - (d_1111_v24_))
            d_1115_v28_: D0
            d_1115_v28_ = D0_DC2((d_1114_v27_).cf3, (0) - (d_1111_v24_), d_1111_v24_)
            rhs155_ = d_1112_v25_
            rhs156_ = d_1115_v28_
            d_1112_v25_ = rhs155_
            d_1114_v27_ = rhs156_
            d_1112_v25_ = (d_1112_v25_ if ((d_1084_v0_)[default__.fm14(globalState)] if (default__.fm14(globalState)) in (d_1084_v0_) else (p1).f4) else d_1112_v25_)
            d_1116_v29_: _dafny.Map
            d_1116_v29_ = _dafny.Map({982: ((p1).f4) == ((self).f4)})
            d_1117_v30_: _dafny.Array
            def lambda60_(d_1118_i3_):
                return (self).f4

            init30_ = lambda60_
            nw170_ = _dafny.Array(None, 22)
            for i0_30_ in range(nw170_.length(0)):
                nw170_[i0_30_] = init30_(i0_30_)
            d_1117_v30_ = nw170_
            d_1119_v31_: _dafny.Map
            d_1119_v31_ = _dafny.Map({d_1117_v30_: d_1084_v0_})
            d_1120_v32_: _dafny.Seq
            d_1120_v32_ = _dafny.SeqWithoutIsStrInference([not((p1).f4)])
            rhs157_ = ((d_1116_v29_)[d_1111_v24_] if (d_1111_v24_) in (d_1116_v29_) else (len(d_1119_v31_)) == (len(d_1120_v32_)))
            rhs158_ = False
            rhs159_ = self.f3
            rhs160_ = False
            rhs161_ = p1.f3
            lhs105_ = self
            lhs106_ = p1
            lhs107_ = globalState
            lhs108_ = p1
            lhs109_ = globalState
            lhs105_.f3 = rhs157_
            lhs106_.f3 = rhs158_
            lhs107_.f0 = rhs159_
            lhs108_.f3 = rhs160_
            lhs109_.f0 = rhs161_
            d_1121_v33_: _dafny.MultiSet
            d_1121_v33_ = _dafny.MultiSet([d_1111_v24_])
            d_1122_v34_: _dafny.Seq
            d_1122_v34_ = _dafny.SeqWithoutIsStrInference([(d_1112_v25_).fm10(d_1111_v24_, (p1).f4, d_1121_v33_, globalState), d_1111_v24_, d_1111_v24_, d_1111_v24_, d_1111_v24_])
            d_1111_v24_ = (d_1122_v34_)[default__.safeIndex((0) - ((d_1111_v24_) - ((d_1122_v34_)[default__.safeIndex(d_1111_v24_, len(d_1122_v34_))])), len(d_1122_v34_))]
        elif True:
            d_1123_v35_: int
            d_1123_v35_ = 877
            d_1124_v36_: _dafny.Seq
            d_1124_v36_ = _dafny.SeqWithoutIsStrInference([d_1123_v35_])
            d_1125_v37_: _dafny.Seq
            d_1125_v37_ = _dafny.SeqWithoutIsStrInference([p1.f3])
            d_1126_v38_: _dafny.Seq
            d_1126_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywhnqq"))
            d_1124_v36_ = ((d_1124_v36_) + (_dafny.SeqWithoutIsStrInference([d_1123_v35_, d_1123_v35_, d_1123_v35_, len(d_1125_v37_), len(d_1126_v38_)]))) + (d_1124_v36_)
            d_1127_v39_: _dafny.Map
            d_1127_v39_ = _dafny.Map({d_1123_v35_: d_1123_v35_})
            d_1127_v39_ = (d_1127_v39_).set(d_1123_v35_, (0) - (d_1123_v35_))
            d_1128_v40_: _dafny.Set
            d_1128_v40_ = _dafny.Set({p1.f3})
            d_1129_v41_: _dafny.Set
            d_1129_v41_ = _dafny.Set({d_1128_v40_})
            d_1129_v41_ = d_1129_v41_
            d_1130_v42_: _dafny.Map
            d_1130_v42_ = _dafny.Map({d_1123_v35_: p1.f3})
            d_1130_v42_ = ((d_1130_v42_) | (default__.fm15((p1).f4, d_1123_v35_, d_1123_v35_, (p1).f4, globalState))) | ((d_1130_v42_) | (d_1130_v42_))
            d_1131_v43_: _dafny.Map
            d_1131_v43_ = _dafny.Map({p1: D0_DC2(p1.f3, len(_dafny.SeqWithoutIsStrInference([p0 for d_1132_i4_ in range(default__.abs(-986))])), d_1123_v35_)})
            d_1133_v44_: _dafny.Seq
            d_1133_v44_ = _dafny.SeqWithoutIsStrInference([d_1131_v43_])
            (self).f3 = (d_1133_v44_) != (d_1133_v44_)
        d_1134_v45_: _dafny.Array
        def lambda61_(d_1135_p1_):
            def lambda62_(d_1136_i5_):
                return ((d_1135_p1_).f4 if d_1135_p1_.f3 else d_1135_p1_.f3)

            return lambda62_

        init31_ = lambda61_(p1)
        nw171_ = _dafny.Array(None, 28)
        for i0_31_ in range(nw171_.length(0)):
            nw171_[i0_31_] = init31_(i0_31_)
        d_1134_v45_ = nw171_
        d_1137_v46_: int
        d_1137_v46_ = 270
        d_1138_v47_: _dafny.Seq
        d_1138_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vg"))
        rhs162_ = ((0) - (d_1137_v46_)) != ((d_1137_v46_ if self.f3 else d_1137_v46_))
        rhs163_ = (d_1134_v45_ if p1.f3 else d_1134_v45_)
        rhs164_ = len(((d_1138_v47_) + (d_1138_v47_)) + (d_1138_v47_))
        lhs110_ = self
        lhs110_.f3 = rhs162_
        d_1134_v45_ = rhs163_
        d_1137_v46_ = rhs164_
        hi9_ = d_1137_v46_
        for d_1139_i6_ in range((d_1137_v46_) * (d_1137_v46_), hi9_):
            d_1140_v48_: _dafny.Map
            d_1140_v48_ = _dafny.Map({self.f3: p0})
            d_1141_v49_: _dafny.Seq
            d_1141_v49_ = _dafny.SeqWithoutIsStrInference([d_1140_v48_])
            d_1142_v50_: _dafny.Array
            nw172_ = _dafny.Array(None, 12)
            nw172_[int(0)] = (d_1137_v46_) * (d_1137_v46_)
            nw172_[int(1)] = len(d_1138_v47_)
            nw172_[int(2)] = len(d_1138_v47_)
            nw172_[int(3)] = d_1139_i6_
            nw172_[int(4)] = len(d_1138_v47_)
            nw172_[int(5)] = (d_1137_v46_) + (d_1139_i6_)
            nw172_[int(6)] = (0) - ((0) - ((0) - (default__.safeModuloInt(len(d_1141_v49_), d_1137_v46_))))
            nw172_[int(7)] = (p1).fm0(globalState)
            nw172_[int(8)] = len(d_1138_v47_)
            nw172_[int(9)] = (d_1139_i6_) + (767)
            nw172_[int(10)] = d_1137_v46_
            nw172_[int(11)] = d_1139_i6_
            d_1142_v50_ = nw172_
            index176_ = default__.safeIndex(370, (d_1142_v50_).length(0))
            (d_1142_v50_)[index176_] = d_1137_v46_
            d_1143_v51_: C0
            nw173_ = C0()
            nw173_.ctor__()
            d_1143_v51_ = nw173_
            d_1144_v52_: _dafny.Set
            d_1144_v52_ = _dafny.Set({d_1143_v51_, d_1143_v51_})
            index177_ = default__.safeIndex(370, (d_1142_v50_).length(0))
            (d_1142_v50_)[index177_] = (d_1139_i6_) * ((d_1139_i6_) * ((0) - (len(d_1144_v52_))))
            d_1145_v53_: _dafny.Set
            d_1145_v53_ = _dafny.Set({p1.f3, p1.f3, self.f3, (p1).f4, default__.fm14(globalState)})
            d_1146_v54_: D0
            d_1146_v54_ = D0_DC0(len((d_1143_v51_).fm11(p0, d_1145_v53_, globalState)))
            source14_ = d_1146_v54_
            if source14_.is_DC1:
                d_1147___mcc_h0_ = source14_.cf1
                d_1148___mcc_h1_ = source14_.cf2
                d_1149_cf2_ = d_1148___mcc_h1_
                d_1150_cf1_ = d_1147___mcc_h0_
                d_1151_v55_: C0
                nw174_ = C0()
                nw174_.ctor__()
                d_1151_v55_ = nw174_
                d_1152_v56_: _dafny.Map
                d_1152_v56_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(d_1142_v50_)[default__.safeIndex(370, (d_1142_v50_).length(0))], (0) - (d_1137_v46_), d_1137_v46_, d_1137_v46_, (0) - ((d_1142_v50_)[default__.safeIndex(370, (d_1142_v50_).length(0))])])): (self).f4})
                def iife87_():
                    coll45_ = _dafny.Map()
                    compr_45_: int
                    for compr_45_ in _dafny.IntegerRange(652, -955):
                        d_1153_v57_: int = compr_45_
                        if ((652) <= (d_1153_v57_)) and ((d_1153_v57_) < (-955)):
                            coll45_[default__.safeModuloInt(d_1153_v57_, d_1139_i6_)] = True
                    return _dafny.Map(coll45_)
                d_1152_v56_ = (iife87_()
                ) | (d_1152_v56_)
                d_1154_v58_: D1
                d_1154_v58_ = D1_DC3(d_1134_v45_)
                d_1154_v58_ = d_1154_v58_
                d_1155_v59_: C0
                nw175_ = C0()
                nw175_.ctor__()
                d_1155_v59_ = nw175_
            elif source14_.is_DC2:
                d_1156___mcc_h2_ = source14_.cf3
                d_1157___mcc_h3_ = source14_.cf4
                d_1158___mcc_h4_ = source14_.cf5
                d_1159_cf5_ = d_1158___mcc_h4_
                d_1160_cf4_ = d_1157___mcc_h3_
                d_1161_cf3_ = d_1156___mcc_h2_
                d_1162_v60_: _dafny.Seq
                d_1162_v60_ = _dafny.SeqWithoutIsStrInference([(p1).f4, self.f3])
                d_1163_v61_: _dafny.Set
                d_1163_v61_ = _dafny.Set({(d_1142_v50_)[default__.safeIndex(370, (d_1142_v50_).length(0))], (d_1142_v50_)[default__.safeIndex(370, (d_1142_v50_).length(0))]})
                (p1).f3 = (len(d_1162_v60_)) in (d_1163_v61_)
                d_1143_v51_ = d_1143_v51_
                d_1084_v0_ = (d_1084_v0_).set(not((p1).f4), ((0) - (d_1137_v46_)) == ((d_1142_v50_)[default__.safeIndex(370, (d_1142_v50_).length(0))]))
                index178_ = default__.safeIndex(370, (d_1142_v50_).length(0))
                (d_1142_v50_)[index178_] = default__.safeDivisionInt(d_1160_cf4_, d_1139_i6_)
            elif True:
                d_1164___mcc_h5_ = source14_.cf0
                d_1165_cf0_ = d_1164___mcc_h5_
                d_1166_v62_: T1
                nw176_ = C1()
                nw176_.ctor__(d_1138_v47_, p1.f3)
                d_1166_v62_ = nw176_
                d_1167_v63_: T1
                d_1167_v63_ = d_1166_v62_
                d_1168_v64_: _dafny.Seq
                d_1168_v64_ = _dafny.SeqWithoutIsStrInference([d_1166_v62_, (d_1167_v63_)])
                d_1169_v65_: _dafny.Map
                d_1169_v65_ = _dafny.Map({(not((p1).f4)) == (True): d_1168_v64_})
                d_1170_v66_: _dafny.Map
                d_1170_v66_ = _dafny.Map({(d_1142_v50_)[default__.safeIndex(370, (d_1142_v50_).length(0))]: d_1169_v65_})
                d_1169_v65_ = (((d_1170_v66_)[d_1165_cf0_] if (d_1165_cf0_) in (d_1170_v66_) else _dafny.Map({self.f3: d_1168_v64_}))) | ((_dafny.Map({False: d_1168_v64_})) | (d_1169_v65_))
                (globalState).f0 = p1.f3
                (globalState).f0 = (p1).f4
                index179_ = default__.safeIndex(370, (d_1142_v50_).length(0))
                (d_1142_v50_)[index179_] = ((d_1165_cf0_) - ((self).fm0(globalState)) if (p1).f4 else d_1139_i6_)
            d_1084_v0_ = (d_1084_v0_).set((p1).f4, self.f3)
            d_1171_v67_: C1
            nw177_ = C1()
            nw177_.ctor__(d_1138_v47_, ((p1).f4) or (p1.f3))
            d_1171_v67_ = nw177_
            d_1171_v67_ = d_1171_v67_
        r0 = ((d_1138_v47_ if self.f3 else ((_dafny.SeqWithoutIsStrInference([p0 for d_1172_i7_ in range(default__.abs(340))])).set(default__.safeIndex(d_1137_v46_, len(_dafny.SeqWithoutIsStrInference([p0 for d_1173_i7_ in range(default__.abs(340))]))), p0)).set(default__.safeIndex(d_1137_v46_, len((_dafny.SeqWithoutIsStrInference([p0 for d_1174_i7_ in range(default__.abs(340))])).set(default__.safeIndex(d_1137_v46_, len(_dafny.SeqWithoutIsStrInference([p0 for d_1175_i7_ in range(default__.abs(340))]))), p0))), p0))) + (d_1138_v47_)
        return r0

    def m0(self, p0, p1, p2, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        d_1176_v0_: _dafny.Array
        def lambda63_(d_1177_p2_):
            def lambda64_(d_1178_i0_):
                return (d_1178_i0_) * (len(_dafny.Map({d_1177_p2_: d_1177_p2_})))

            return lambda64_

        init32_ = lambda63_(p2)
        nw178_ = _dafny.Array(None, 5)
        for i0_32_ in range(nw178_.length(0)):
            nw178_[i0_32_] = init32_(i0_32_)
        d_1176_v0_ = nw178_
        index180_ = default__.safeIndex(847, (d_1176_v0_).length(0))
        (d_1176_v0_)[index180_] = p2
        index181_ = default__.safeIndex(847, (d_1176_v0_).length(0))
        (d_1176_v0_)[index181_] = p2
        d_1179_v1_: _dafny.Array
        def lambda65_(d_1180_i1_):
            return not((self.f3) or (True))

        init33_ = lambda65_
        nw179_ = _dafny.Array(None, 28)
        for i0_33_ in range(nw179_.length(0)):
            nw179_[i0_33_] = init33_(i0_33_)
        d_1179_v1_ = nw179_
        index182_ = default__.safeIndex(458, (d_1179_v1_).length(0))
        (d_1179_v1_)[index182_] = False
        d_1181_v2_: str
        d_1181_v2_ = _dafny.CodePoint('g')
        d_1182_v3_: _dafny.Map
        d_1182_v3_ = _dafny.Map({p1: d_1181_v2_})
        d_1183_v4_: _dafny.Seq
        d_1183_v4_ = _dafny.SeqWithoutIsStrInference([d_1181_v2_, ((d_1182_v3_)[p1] if (p1) in (d_1182_v3_) else _dafny.CodePoint('x')), default__.fm23((self).f4, (d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))], globalState)])
        d_1184_v5_: _dafny.Seq
        d_1184_v5_ = _dafny.SeqWithoutIsStrInference([386])
        d_1185_v6_: _dafny.Seq
        d_1185_v6_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - (p2) for d_1186_i2_ in range(default__.abs(580))]), d_1184_v5_])
        d_1187_v7_: _dafny.MultiSet
        d_1187_v7_ = _dafny.MultiSet([len(d_1183_v4_), (d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))], len(d_1185_v6_), (d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))], 756])
        index183_ = default__.safeIndex(847, (d_1176_v0_).length(0))
        index184_ = default__.safeIndex(458, (d_1179_v1_).length(0))
        rhs165_ = (self.f3) and (self.f3)
        rhs166_ = (d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))]
        rhs167_ = ((d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))]) + ((d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))])
        rhs168_ = ((p2) + ((d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))])) <= ((d_1187_v7_).cardinality)
        rhs169_ = (0) - ((p2) - ((0) - (default__.safeDivisionInt(p2, (d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))]))))
        lhs111_ = self
        lhs112_ = d_1176_v0_
        lhs113_ = default__.safeIndex(847, (d_1176_v0_).length(0))
        lhs114_ = d_1179_v1_
        lhs115_ = default__.safeIndex(458, (d_1179_v1_).length(0))
        lhs111_.f3 = rhs165_
        lhs112_[lhs113_] = rhs166_
        r1 = rhs167_
        lhs114_[lhs115_] = rhs168_
        r1 = rhs169_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1179_v1_).length(0)):
            d_1188_i3_: int = guard_loop_2_
            if (True) and (((0) <= (d_1188_i3_)) and ((d_1188_i3_) < ((d_1179_v1_).length(0)))):
                (d_1179_v1_)[(d_1188_i3_)] = not((self).f4)
        d_1189_i4_: int
        d_1189_i4_ = 0
        with _dafny.label("5"):
            while ((d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))]) >= ((d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))]):
                with _dafny.c_label("5"):
                    if (d_1189_i4_) >= (100):
                        raise _dafny.Break("5")
                    d_1189_i4_ = (d_1189_i4_) + (1)
                    d_1190_v8_: _dafny.Map
                    d_1190_v8_ = _dafny.Map({p2: p2})
                    d_1190_v8_ = d_1190_v8_
                    d_1191_v9_: _dafny.Seq
                    d_1191_v9_ = _dafny.SeqWithoutIsStrInference([(self).f4, True])
                    d_1192_v10_: _dafny.Seq
                    d_1192_v10_ = _dafny.SeqWithoutIsStrInference([d_1183_v4_ for d_1193_i5_ in range(default__.abs(450))])
                    d_1194_v11_: D1
                    d_1194_v11_ = D1_DC5(False, self.f3, default__.fm31(False, (d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))], (self).f4, d_1191_v9_, globalState), len(_dafny.Map({p2: d_1192_v10_})))
                    index185_ = default__.safeIndex(458, (d_1179_v1_).length(0))
                    (d_1179_v1_)[index185_] = (d_1194_v11_).cf8
                    d_1195_v12_: _dafny.Array
                    def lambda66_(d_1196_v1_, d_1197_p0_, d_1198_v0_):
                        def lambda67_(d_1199_i6_):
                            return (_dafny.SeqWithoutIsStrInference([(d_1196_v1_)[default__.safeIndex(458, (d_1196_v1_).length(0))], d_1197_p0_, self.f3])).set(default__.safeIndex((d_1198_v0_)[default__.safeIndex(847, (d_1198_v0_).length(0))], len(_dafny.SeqWithoutIsStrInference([(d_1196_v1_)[default__.safeIndex(458, (d_1196_v1_).length(0))], d_1197_p0_, self.f3]))), (d_1196_v1_)[default__.safeIndex(458, (d_1196_v1_).length(0))])

                        return lambda67_

                    init34_ = lambda66_(d_1179_v1_, p0, d_1176_v0_)
                    nw180_ = _dafny.Array(None, 6)
                    for i0_34_ in range(nw180_.length(0)):
                        nw180_[i0_34_] = init34_(i0_34_)
                    d_1195_v12_ = nw180_
                    rhs170_ = (default__.safeDivisionInt(684, len(_dafny.SeqWithoutIsStrInference([(d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))], p2])))) in (d_1184_v5_)
                    rhs171_ = d_1195_v12_
                    rhs172_ = (p2) < (158)
                    lhs116_ = globalState
                    lhs117_ = globalState
                    lhs116_.f0 = rhs170_
                    d_1195_v12_ = rhs171_
                    lhs117_.f0 = rhs172_
                    d_1200_v13_: _dafny.Array
                    nw181_ = _dafny.Array(_dafny.Map({}), 27)
                    d_1200_v13_ = nw181_
                    d_1201_v14_: _dafny.Map
                    d_1201_v14_ = _dafny.Map({False: (0) - ((d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))])})
                    index186_ = default__.safeIndex(908, (d_1200_v13_).length(0))
                    (d_1200_v13_)[index186_] = (d_1201_v14_) | (d_1201_v14_)
                    index187_ = default__.safeIndex(908, (d_1200_v13_).length(0))
                    (d_1200_v13_)[index187_] = d_1201_v14_
                    pass
            pass
        index188_ = default__.safeIndex(458, (d_1179_v1_).length(0))
        (d_1179_v1_)[index188_] = (d_1179_v1_)[default__.safeIndex(458, (d_1179_v1_).length(0))]
        d_1202_v15_: _dafny.Set
        d_1202_v15_ = _dafny.Set({p2})
        d_1203_v17_: D7
        def iife88_():
            coll46_ = _dafny.Set()
            compr_46_: int
            for compr_46_ in (d_1202_v15_).Elements:
                d_1204_v16_: int = compr_46_
                if (d_1204_v16_) in (d_1202_v15_):
                    coll46_ = coll46_.union(_dafny.Set([(d_1204_v16_) + (len(_dafny.SeqWithoutIsStrInference([True, False])))]))
            return _dafny.Set(coll46_)
        d_1203_v17_ = D7_DC16(default__.fm36(len(iife88_()
), (d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))], globalState))
        d_1205_v18_: D1
        d_1205_v18_ = D1_DC4((d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))])
        d_1206_v19_: _dafny.Map
        d_1206_v19_ = _dafny.Map({d_1203_v17_: d_1205_v18_})
        d_1207_v20_: _dafny.Set
        d_1207_v20_ = _dafny.Set({(D1_DC5((self).f4, (self).f4, (d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))], -772)).cf9})
        d_1208_v21_: _dafny.Map
        d_1208_v21_ = _dafny.Map({len(p1): d_1207_v20_})
        pat_let_tv36_ = d_1208_v21_
        pat_let_tv37_ = d_1208_v21_
        pat_let_tv38_ = p2
        pat_let_tv39_ = d_1207_v20_
        def iife90_(_pat_let22_0):
            def iife91_(d_1209_dt__update__tmp_h0_):
                def iife92_(_pat_let23_0):
                    def iife93_(d_1210_dt__update_hcf25_h0_):
                        return D7_DC16(d_1210_dt__update_hcf25_h0_)
                    return iife93_(_pat_let23_0)
                return iife92_(pat_let_tv36_)
            return iife91_(_pat_let22_0)
        def iife89_(_pat_let21_0):
            def iife94_(d_1211_dt__update__tmp_h1_):
                def iife95_(_pat_let24_0):
                    def iife96_(d_1212_dt__update_hcf25_h1_):
                        return D7_DC16(d_1212_dt__update_hcf25_h1_)
                    return iife96_(_pat_let24_0)
                return iife95_((pat_let_tv37_).set(pat_let_tv38_, pat_let_tv39_))
            return iife94_(_pat_let21_0)
        d_1206_v19_ = (d_1206_v19_).set(iife89_(iife90_(d_1203_v17_)), d_1205_v18_)
        d_1213_v22_: _dafny.Map
        d_1213_v22_ = _dafny.Map({not (p0) or ((d_1179_v1_)[default__.safeIndex(458, (d_1179_v1_).length(0))]): d_1176_v0_})
        r0 = ((d_1213_v22_)[((self).f4 if (d_1179_v1_)[default__.safeIndex(458, (d_1179_v1_).length(0))] else (d_1179_v1_)[default__.safeIndex(458, (d_1179_v1_).length(0))])] if (((self).f4 if (d_1179_v1_)[default__.safeIndex(458, (d_1179_v1_).length(0))] else (d_1179_v1_)[default__.safeIndex(458, (d_1179_v1_).length(0))])) in (d_1213_v22_) else d_1176_v0_)
        r1 = (d_1176_v0_)[default__.safeIndex(847, (d_1176_v0_).length(0))]
        return r0, r1

    def m7(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        if (self).f4:
            r0 = False
            d_1214_v0_: D1
            d_1214_v0_ = D1_DC5((self).f4, (self).f4, p0, -259)
            def iife97_(_pat_let25_0):
                def iife98_(d_1215_dt__update__tmp_h0_):
                    def iife99_(_pat_let26_0):
                        def iife100_(d_1216_dt__update_hcf8_h0_):
                            def iife101_(_pat_let27_0):
                                def iife102_(d_1217_dt__update_hcf9_h0_):
                                    return D1_DC5(d_1216_dt__update_hcf8_h0_, d_1217_dt__update_hcf9_h0_, (d_1215_dt__update__tmp_h0_).cf10, (d_1215_dt__update__tmp_h0_).cf11)
                                return iife102_(_pat_let27_0)
                            return iife101_(False)
                        return iife100_(_pat_let26_0)
                    return iife99_(self.f3)
                return iife98_(_pat_let25_0)
            (self).f3 = (p0) < ((len(default__.fm30(iife97_(d_1214_v0_), p0, p0, globalState))) + (p0))
            d_1218_v1_: _dafny.Seq
            d_1218_v1_ = _dafny.SeqWithoutIsStrInference([(self).f4])
            d_1219_v2_: C1
            nw182_ = C1()
            nw182_.ctor__((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pavx"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wibc"))), (not((self).f4) if (d_1218_v1_)[default__.safeIndex(p0, len(d_1218_v1_))] else (self).f4))
            d_1219_v2_ = nw182_
            d_1220_v3_: _dafny.MultiSet
            d_1220_v3_ = _dafny.MultiSet([p0])
            d_1221_v4_: _dafny.Map
            d_1221_v4_ = _dafny.Map({self.f3: d_1220_v3_})
            d_1222_v5_: _dafny.Map
            d_1222_v5_ = _dafny.Map({(len(d_1221_v4_)) + (p0): p0})
            d_1222_v5_ = (d_1222_v5_).set((self).fm0(globalState), (p0) - (p0))
            d_1223_v6_: C0
            nw183_ = C0()
            nw183_.ctor__()
            d_1223_v6_ = nw183_
        elif True:
            d_1224_v7_: C0
            nw184_ = C0()
            nw184_.ctor__()
            d_1224_v7_ = nw184_
            d_1225_v8_: _dafny.Seq
            d_1226_v9_: _dafny.Array
            out22_: _dafny.Seq
            out23_: _dafny.Array
            out22_, out23_ = (self).m8(716, p0, globalState)
            d_1225_v8_ = out22_
            d_1226_v9_ = out23_
            (self).f3 = (self).f4
            d_1227_v10_: _dafny.Map
            d_1227_v10_ = _dafny.Map({(self).f4: _dafny.MultiSet([True])})
            d_1227_v10_ = (d_1227_v10_).set((self.f3) and (self.f3), default__.fm20(self.f3, globalState))
            d_1228_v11_: _dafny.Seq
            d_1228_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgbmpb"))
            d_1228_v11_ = d_1228_v11_
        d_1229_v12_: int
        d_1229_v12_ = -138
        d_1230_v13_: _dafny.Seq
        d_1230_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvdeve"))
        d_1231_v14_: _dafny.Seq
        d_1231_v14_ = _dafny.SeqWithoutIsStrInference([d_1230_v13_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1232_i0_ in range(default__.abs(99))])])
        d_1233_v15_: _dafny.Seq
        d_1233_v15_ = _dafny.SeqWithoutIsStrInference([(self).f4, (self).f4])
        d_1234_v16_: _dafny.Map
        d_1234_v16_ = _dafny.Map({self.f3: (self).f4})
        d_1235_v17_: _dafny.Map
        d_1235_v17_ = _dafny.Map({(self).f4: d_1229_v12_})
        d_1236_v19_: _dafny.Array
        nw185_ = _dafny.Array(None, 22)
        nw185_[int(0)] = d_1229_v12_
        nw185_[int(1)] = p0
        nw185_[int(2)] = len(d_1233_v15_)
        nw185_[int(3)] = len(d_1233_v15_)
        nw185_[int(4)] = -836
        nw185_[int(5)] = len(d_1234_v16_)
        nw185_[int(6)] = -414
        nw185_[int(7)] = d_1229_v12_
        nw185_[int(8)] = d_1229_v12_
        nw185_[int(9)] = ((d_1235_v17_)[(self).f4] if ((self).f4) in (d_1235_v17_) else p0)
        nw185_[int(10)] = default__.fm31(True, d_1229_v12_, self.f3, d_1233_v15_, globalState)
        nw185_[int(11)] = d_1229_v12_
        nw185_[int(12)] = d_1229_v12_
        nw185_[int(13)] = len(d_1230_v13_)
        nw185_[int(14)] = p0
        def iife103_():
            coll47_ = _dafny.Map()
            compr_47_: int
            for compr_47_ in _dafny.IntegerRange(-59, 797):
                d_1237_v18_: int = compr_47_
                if ((-59) <= (d_1237_v18_)) and ((d_1237_v18_) < (797)):
                    coll47_[(d_1237_v18_) - ((0) - (p0))] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_1229_v12_: 902}))]))
            return _dafny.Map(coll47_)
        nw185_[int(15)] = len(iife103_()
        )
        nw185_[int(16)] = d_1229_v12_
        nw185_[int(17)] = p0
        nw185_[int(18)] = d_1229_v12_
        nw185_[int(19)] = d_1229_v12_
        nw185_[int(20)] = d_1229_v12_
        nw185_[int(21)] = len(d_1230_v13_)
        d_1236_v19_ = nw185_
        def lambda68_(source15_):
            d_1238___mcc_h0_ = source15_
            d_1239_cf24_ = d_1238___mcc_h0_
            return self.f3

        rhs173_ = lambda68_(d_1231_v14_)
        rhs174_ = (_dafny.MultiSet([d_1236_v19_])).cardinality
        lhs118_ = globalState
        lhs118_.f0 = rhs173_
        d_1229_v12_ = rhs174_
        hi10_ = d_1229_v12_
        for d_1240_i1_ in range(d_1229_v12_, hi10_):
            index189_ = default__.safeIndex(46, (d_1236_v19_).length(0))
            (d_1236_v19_)[index189_] = d_1229_v12_
            index190_ = default__.safeIndex(46, (d_1236_v19_).length(0))
            (d_1236_v19_)[index190_] = d_1229_v12_
            d_1241_v20_: D1
            d_1241_v20_ = D1_DC6()
            d_1241_v20_ = d_1241_v20_
            index191_ = default__.safeIndex(46, (d_1236_v19_).length(0))
            (d_1236_v19_)[index191_] = 614
            if False:
                d_1242_v21_: _dafny.Set
                d_1242_v21_ = _dafny.Set({(d_1236_v19_)[default__.safeIndex(46, (d_1236_v19_).length(0))]})
                index192_ = default__.safeIndex(46, (d_1236_v19_).length(0))
                (d_1236_v19_)[index192_] = default__.fm31(self.f3, d_1229_v12_, (d_1242_v21_).ispropersubset(d_1242_v21_), d_1233_v15_, globalState)
                d_1243_v22_: _dafny.Array
                nw186_ = _dafny.Array(_dafny.Seq({}), 3)
                d_1243_v22_ = nw186_
                d_1244_v23_: _dafny.Seq
                d_1244_v23_ = _dafny.SeqWithoutIsStrInference([d_1229_v12_, (0) - ((d_1236_v19_)[default__.safeIndex(46, (d_1236_v19_).length(0))])])
                index193_ = default__.safeIndex(375, (d_1243_v22_).length(0))
                (d_1243_v22_)[index193_] = d_1244_v23_
                index194_ = default__.safeIndex(375, (d_1243_v22_).length(0))
                (d_1243_v22_)[index194_] = ((_dafny.SeqWithoutIsStrInference([(d_1236_v19_)[default__.safeIndex(46, (d_1236_v19_).length(0))]])).set(default__.safeIndex((d_1236_v19_)[default__.safeIndex(46, (d_1236_v19_).length(0))], len(_dafny.SeqWithoutIsStrInference([(d_1236_v19_)[default__.safeIndex(46, (d_1236_v19_).length(0))]]))), len(d_1233_v15_))) + (d_1244_v23_)
                index195_ = default__.safeIndex(46, (d_1236_v19_).length(0))
                (d_1236_v19_)[index195_] = d_1240_i1_
                (globalState).f0 = (self).f4
                d_1245_v24_: _dafny.MultiSet
                d_1245_v24_ = _dafny.MultiSet([p0, default__.safeDivisionInt((_dafny.MultiSet([self.f3])).cardinality, (d_1236_v19_)[default__.safeIndex(46, (d_1236_v19_).length(0))]), p0, d_1240_i1_])
                d_1245_v24_ = (d_1245_v24_).set(default__.safeModuloInt((d_1236_v19_)[default__.safeIndex(46, (d_1236_v19_).length(0))], len(d_1242_v21_)), default__.abs(d_1240_i1_))
            elif True:
                d_1246_v25_: _dafny.MultiSet
                d_1246_v25_ = _dafny.MultiSet([(d_1236_v19_)[default__.safeIndex(46, (d_1236_v19_).length(0))], (d_1236_v19_)[default__.safeIndex(46, (d_1236_v19_).length(0))]])
                d_1247_v26_: _dafny.Map
                d_1247_v26_ = _dafny.Map({-624: d_1246_v25_})
                d_1248_v27_: _dafny.Seq
                d_1248_v27_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1249_v28_: D14
                d_1249_v28_ = D14_DC30(_dafny.MultiSet([d_1229_v12_]))
                (globalState).f0 = ((((d_1247_v26_)[(d_1248_v27_)[default__.safeIndex((d_1236_v19_)[default__.safeIndex(46, (d_1236_v19_).length(0))], len(d_1248_v27_))]] if ((d_1248_v27_)[default__.safeIndex((d_1236_v19_)[default__.safeIndex(46, (d_1236_v19_).length(0))], len(d_1248_v27_))]) in (d_1247_v26_) else (d_1249_v28_).cf43)).intersection(d_1246_v25_)).issubset((default__.fm24(globalState)) | (d_1246_v25_))
                d_1250_v29_: _dafny.Set
                d_1250_v29_ = _dafny.Set({(self).f4, self.f3, (self).f4, self.f3})
                d_1250_v29_ = d_1250_v29_
                d_1251_v30_: str
                d_1251_v30_ = _dafny.CodePoint('n')
                d_1252_v31_: _dafny.Map
                d_1252_v31_ = _dafny.Map({d_1251_v30_: d_1251_v30_})
                d_1252_v31_ = (d_1252_v31_).set(d_1251_v30_, d_1251_v30_)
                d_1253_v32_: _dafny.Map
                d_1253_v32_ = _dafny.Map({(d_1236_v19_)[default__.safeIndex(46, (d_1236_v19_).length(0))]: d_1251_v30_})
                d_1253_v32_ = (d_1253_v32_).set((d_1236_v19_)[default__.safeIndex(46, (d_1236_v19_).length(0))], d_1251_v30_)
                d_1234_v16_ = d_1234_v16_
        d_1254_v33_: str
        d_1254_v33_ = _dafny.CodePoint('b')
        d_1255_v34_: _dafny.MultiSet
        d_1255_v34_ = _dafny.MultiSet([d_1254_v33_, d_1254_v33_])
        (globalState).f0 = (d_1255_v34_) in (_dafny.Set({d_1255_v34_}))
        hi11_ = default__.safeModuloInt(d_1229_v12_, len(d_1230_v13_))
        for d_1256_i2_ in range(d_1229_v12_, hi11_):
            d_1257_v35_: C0
            nw187_ = C0()
            nw187_.ctor__()
            d_1257_v35_ = nw187_
            d_1229_v12_ = d_1229_v12_
            d_1258_v36_: D4
            d_1258_v36_ = D4_DC11((_dafny.SeqWithoutIsStrInference([False, False])) + (d_1233_v15_))
            source16_ = d_1258_v36_
            if source16_.is_DC12:
                d_1259___mcc_h1_ = source16_.cf19
                d_1260___mcc_h2_ = source16_.cf20
                d_1261___mcc_h3_ = source16_.cf21
                d_1262_cf21_ = d_1261___mcc_h3_
                d_1263_cf20_ = d_1260___mcc_h2_
                d_1264_cf19_ = d_1259___mcc_h1_
                d_1265_v37_: _dafny.Map
                d_1265_v37_ = _dafny.Map({375: self.f3})
                d_1266_v38_: _dafny.MultiSet
                d_1266_v38_ = _dafny.MultiSet([d_1230_v13_])
                d_1267_v39_: _dafny.MultiSet
                d_1267_v39_ = _dafny.MultiSet([(self).f4])
                d_1268_v40_: _dafny.Array
                nw188_ = _dafny.Array(None, 26)
                nw188_[int(0)] = ((d_1265_v37_)[d_1229_v12_] if (d_1229_v12_) in (d_1265_v37_) else (self).f4)
                nw188_[int(1)] = d_1262_cf21_
                nw188_[int(2)] = default__.fm14(globalState)
                nw188_[int(3)] = (self).f4
                nw188_[int(4)] = (self).f4
                nw188_[int(5)] = (self).f4
                nw188_[int(6)] = self.f3
                nw188_[int(7)] = self.f3
                nw188_[int(8)] = d_1262_cf21_
                nw188_[int(9)] = not (self.f3) or (d_1262_cf21_)
                nw188_[int(10)] = d_1262_cf21_
                nw188_[int(11)] = ((self).f4 if self.f3 else (self).f4)
                nw188_[int(12)] = ((self).f4 if self.f3 else (self).f4)
                nw188_[int(13)] = default__.fm14(globalState)
                nw188_[int(14)] = default__.fm18(722, d_1254_v33_, (d_1266_v38_).cardinality, globalState)
                nw188_[int(15)] = (d_1263_cf20_) > (((d_1267_v39_)[(self).f4] if ((self).f4) in (d_1267_v39_) else d_1263_cf20_))
                nw188_[int(16)] = True
                nw188_[int(17)] = d_1262_cf21_
                nw188_[int(18)] = (p0) >= ((0) - (d_1229_v12_))
                nw188_[int(19)] = self.f3
                nw188_[int(20)] = self.f3
                nw188_[int(21)] = False
                nw188_[int(22)] = self.f3
                nw188_[int(23)] = self.f3
                nw188_[int(24)] = (self).f4
                nw188_[int(25)] = d_1262_cf21_
                d_1268_v40_ = nw188_
                index196_ = default__.safeIndex(174, (d_1268_v40_).length(0))
                (d_1268_v40_)[index196_] = self.f3
                d_1269_v41_: _dafny.Map
                d_1269_v41_ = _dafny.Map({d_1263_cf20_: d_1264_cf19_})
                index197_ = default__.safeIndex(174, (d_1268_v40_).length(0))
                (d_1268_v40_)[index197_] = default__.fm21(d_1229_v12_, _dafny.MultiSet([d_1269_v41_, d_1269_v41_]), globalState)
                (self).f3 = self.f3
                d_1229_v12_ = d_1264_cf19_
                (globalState).f0 = (self).f4
            elif source16_.is_DC11:
                d_1270___mcc_h4_ = source16_.cf18
                d_1271_cf18_ = d_1270___mcc_h4_
                d_1272_v42_: _dafny.Array
                nw189_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_1272_v42_ = nw189_
                d_1273_v43_: _dafny.Map
                d_1273_v43_ = _dafny.Map({not(self.f3): d_1271_cf18_})
                d_1274_v44_: _dafny.Array
                nw190_ = _dafny.Array(None, 18)
                nw190_[int(0)] = d_1233_v15_
                nw190_[int(1)] = d_1233_v15_
                nw190_[int(2)] = d_1271_cf18_
                nw190_[int(3)] = d_1271_cf18_
                nw190_[int(4)] = _dafny.SeqWithoutIsStrInference([(self).f4])
                nw190_[int(5)] = d_1271_cf18_
                nw190_[int(6)] = (d_1271_cf18_).set(default__.safeIndex(d_1229_v12_, len(d_1271_cf18_)), self.f3)
                nw190_[int(7)] = d_1233_v15_
                nw190_[int(8)] = _dafny.SeqWithoutIsStrInference([self.f3, (self).f4])
                nw190_[int(9)] = d_1271_cf18_
                nw190_[int(10)] = d_1271_cf18_
                nw190_[int(11)] = d_1271_cf18_
                nw190_[int(12)] = d_1271_cf18_
                nw190_[int(13)] = d_1233_v15_
                nw190_[int(14)] = d_1233_v15_
                nw190_[int(15)] = default__.fm34(d_1256_i2_, d_1229_v12_, globalState)
                nw190_[int(16)] = d_1271_cf18_
                nw190_[int(17)] = ((d_1273_v43_)[self.f3] if (self.f3) in (d_1273_v43_) else d_1271_cf18_)
                d_1274_v44_ = nw190_
                index198_ = default__.safeIndex(487, (d_1272_v42_).length(0))
                (d_1272_v42_)[index198_] = d_1274_v44_
                index199_ = default__.safeIndex(487, (d_1272_v42_).length(0))
                rhs175_ = d_1254_v33_
                rhs176_ = d_1274_v44_
                lhs119_ = d_1272_v42_
                lhs120_ = default__.safeIndex(487, (d_1272_v42_).length(0))
                d_1254_v33_ = rhs175_
                lhs119_[lhs120_] = rhs176_
                d_1275_v45_: _dafny.Seq
                d_1275_v45_ = _dafny.SeqWithoutIsStrInference([len(d_1230_v13_)])
                d_1276_v46_: _dafny.Seq
                d_1276_v46_ = _dafny.SeqWithoutIsStrInference([d_1275_v45_, _dafny.SeqWithoutIsStrInference([d_1229_v12_ for d_1277_i3_ in range(default__.abs(39))])])
                d_1278_v47_: _dafny.Set
                d_1278_v47_ = _dafny.Set({(d_1276_v46_)[default__.safeIndex(d_1229_v12_, len(d_1276_v46_))], _dafny.SeqWithoutIsStrInference([p0, d_1229_v12_]), _dafny.SeqWithoutIsStrInference([d_1256_i2_])})
                d_1278_v47_ = d_1278_v47_
                d_1279_v48_: _dafny.MultiSet
                d_1279_v48_ = _dafny.MultiSet([(self).f4])
                d_1279_v48_ = d_1279_v48_
                d_1280_v49_: D7
                d_1280_v49_ = D7_DC17(d_1256_i2_, (self).f4, (self).f4, len(d_1230_v13_))
                d_1281_v50_: _dafny.Set
                d_1281_v50_ = _dafny.Set({(d_1280_v49_).cf28, self.f3})
                d_1282_v51_: D7
                d_1282_v51_ = D7_DC16(_dafny.Map({len(d_1275_v45_): d_1281_v50_}))
                d_1283_v52_: _dafny.Map
                d_1283_v52_ = _dafny.Map({(0) - (d_1229_v12_): d_1281_v50_})
                d_1282_v51_ = D7_DC16(d_1283_v52_)
            elif True:
                d_1284___mcc_h5_ = source16_.cf22
                d_1285_cf22_ = d_1284___mcc_h5_
                index200_ = default__.safeIndex(514, (d_1236_v19_).length(0))
                (d_1236_v19_)[index200_] = d_1229_v12_
                index201_ = default__.safeIndex(514, (d_1236_v19_).length(0))
                (d_1236_v19_)[index201_] = default__.safeModuloInt(d_1229_v12_, d_1229_v12_)
                d_1286_v53_: _dafny.Seq
                d_1286_v53_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1286_v53_ = d_1286_v53_
                (globalState).f0 = self.f3
                index202_ = default__.safeIndex(514, (d_1236_v19_).length(0))
                (d_1236_v19_)[index202_] = (d_1236_v19_)[default__.safeIndex(514, (d_1236_v19_).length(0))]
            r1 = self.f3
        d_1287_v54_: D1
        d_1287_v54_ = D1_DC5(self.f3, (p0) <= (p0), (d_1229_v12_) - (default__.fm31((self).f4, len(d_1235_v17_), (self).f4, d_1233_v15_, globalState)), (d_1229_v12_) - (d_1229_v12_))
        source17_ = d_1287_v54_
        if source17_.is_DC4:
            d_1288___mcc_h6_ = source17_.cf7
            d_1289_cf7_ = d_1288___mcc_h6_
            (self).f3 = (self).f4
            d_1290_v55_: _dafny.Seq
            d_1291_v56_: _dafny.Array
            out24_: _dafny.Seq
            out25_: _dafny.Array
            out24_, out25_ = (self).m8((d_1255_v34_).cardinality, (d_1229_v12_) - (p0), globalState)
            d_1290_v55_ = out24_
            d_1291_v56_ = out25_
            r0 = self.f3
            d_1292_v57_: _dafny.Array
            nw191_ = _dafny.Array(_dafny.Map({}), 5)
            d_1292_v57_ = nw191_
            index203_ = default__.safeIndex(39, (d_1292_v57_).length(0))
            (d_1292_v57_)[index203_] = d_1234_v16_
            index204_ = default__.safeIndex(39, (d_1292_v57_).length(0))
            (d_1292_v57_)[index204_] = d_1234_v16_
        elif source17_.is_DC5:
            d_1293___mcc_h7_ = source17_.cf8
            d_1294___mcc_h8_ = source17_.cf9
            d_1295___mcc_h9_ = source17_.cf10
            d_1296___mcc_h10_ = source17_.cf11
            d_1297_cf11_ = d_1296___mcc_h10_
            d_1298_cf10_ = d_1295___mcc_h9_
            d_1299_cf9_ = d_1294___mcc_h8_
            d_1300_cf8_ = d_1293___mcc_h7_
            (globalState).f0 = self.f3
            d_1299_cf9_ = (d_1233_v15_) == (d_1233_v15_)
            (globalState).f0 = not (d_1300_cf8_) or (not(d_1299_cf9_))
            d_1301_v58_: C2
            nw192_ = C2()
            nw192_.ctor__(not(False), (self).f4, self.f3, d_1300_cf8_)
            d_1301_v58_ = nw192_
        elif source17_.is_DC6:
            (globalState).f0 = (self).f4
            if self.f3:
                (globalState).f0 = (self).f4
                d_1302_v59_: C0
                nw193_ = C0()
                nw193_.ctor__()
                d_1302_v59_ = nw193_
                d_1303_v60_: _dafny.Seq
                d_1303_v60_ = _dafny.SeqWithoutIsStrInference([((d_1235_v17_)[(self).f4] if ((self).f4) in (d_1235_v17_) else len(_dafny.SeqWithoutIsStrInference([self.f3, self.f3])))])
                d_1304_v61_: _dafny.Map
                d_1304_v61_ = _dafny.Map({(0) - ((d_1303_v60_)[default__.safeIndex(d_1229_v12_, len(d_1303_v60_))]): (0) - (p0)})
                d_1304_v61_ = (d_1304_v61_).set((d_1229_v12_) - (p0), p0)
                d_1305_v62_: _dafny.Set
                d_1305_v62_ = _dafny.Set({True})
                d_1306_v63_: D12
                d_1306_v63_ = D12_DC26(d_1305_v62_)
                d_1307_v64_: D12
                d_1307_v64_ = D12_DC28(d_1306_v63_)
                d_1308_v65_: _dafny.Map
                d_1308_v65_ = _dafny.Map({self.f3: d_1307_v64_})
                d_1309_v66_: _dafny.Map
                d_1309_v66_ = _dafny.Map({d_1308_v65_: (d_1302_v59_).fm11(_dafny.CodePoint('m'), d_1305_v62_, globalState)})
                rhs177_ = _dafny.Map({_dafny.Map({self.f3: d_1307_v64_}): d_1230_v13_})
                rhs178_ = d_1233_v15_
                rhs179_ = (self).f4
                lhs121_ = globalState
                d_1309_v66_ = rhs177_
                d_1233_v15_ = rhs178_
                lhs121_.f0 = rhs179_
                d_1310_v67_: _dafny.Array
                nw194_ = _dafny.Array(_dafny.CodePoint('D'), 22)
                d_1310_v67_ = nw194_
                index205_ = default__.safeIndex(19, (d_1310_v67_).length(0))
                (d_1310_v67_)[index205_] = _dafny.CodePoint('a')
                index206_ = default__.safeIndex(19, (d_1310_v67_).length(0))
                (d_1310_v67_)[index206_] = d_1254_v33_
            elif True:
                index207_ = default__.safeIndex(941, (d_1236_v19_).length(0))
                (d_1236_v19_)[index207_] = (0) - (default__.safeModuloInt(p0, p0))
                index208_ = default__.safeIndex(941, (d_1236_v19_).length(0))
                (d_1236_v19_)[index208_] = d_1229_v12_
                index209_ = default__.safeIndex(941, (d_1236_v19_).length(0))
                (d_1236_v19_)[index209_] = p0
                d_1229_v12_ = (d_1236_v19_)[default__.safeIndex(941, (d_1236_v19_).length(0))]
                d_1311_v68_: _dafny.Array
                nw195_ = _dafny.Array(False, 22)
                d_1311_v68_ = nw195_
                d_1312_v69_: D0
                d_1312_v69_ = D0_DC2(self.f3, 656, p0)
                d_1313_v70_: _dafny.Seq
                d_1313_v70_ = _dafny.SeqWithoutIsStrInference([len(d_1230_v13_)])
                index210_ = default__.safeIndex(941, (d_1236_v19_).length(0))
                rhs180_ = (self.f3) == (False)
                rhs181_ = default__.safeModuloInt(((d_1236_v19_)[default__.safeIndex(941, (d_1236_v19_).length(0))]) * ((d_1312_v69_).cf4), d_1229_v12_)
                rhs182_ = (d_1313_v70_)[default__.safeIndex((d_1236_v19_)[default__.safeIndex(941, (d_1236_v19_).length(0))], len(d_1313_v70_))]
                rhs183_ = (default__.fm30(d_1287_v54_, (d_1236_v19_)[default__.safeIndex(941, (d_1236_v19_).length(0))], p0, globalState)) <= ((d_1230_v13_).set(default__.safeIndex((d_1236_v19_)[default__.safeIndex(941, (d_1236_v19_).length(0))], len(d_1230_v13_)), d_1254_v33_))
                rhs184_ = d_1311_v68_
                lhs122_ = globalState
                lhs123_ = d_1236_v19_
                lhs124_ = default__.safeIndex(941, (d_1236_v19_).length(0))
                lhs125_ = globalState
                lhs122_.f0 = rhs180_
                lhs123_[lhs124_] = rhs181_
                d_1229_v12_ = rhs182_
                lhs125_.f0 = rhs183_
                d_1311_v68_ = rhs184_
                d_1314_v71_: C0
                nw196_ = C0()
                nw196_.ctor__()
                d_1314_v71_ = nw196_
            d_1315_v72_: _dafny.Seq
            d_1316_v73_: _dafny.Array
            out26_: _dafny.Seq
            out27_: _dafny.Array
            out26_, out27_ = (self).m8((d_1229_v12_) + ((self).fm0(globalState)), default__.fm31((self).f4, len(d_1230_v13_), self.f3, d_1233_v15_, globalState), globalState)
            d_1315_v72_ = out26_
            d_1316_v73_ = out27_
            d_1317_v74_: _dafny.Set
            d_1317_v74_ = _dafny.Set({d_1229_v12_})
            d_1317_v74_ = d_1317_v74_
        elif True:
            d_1318___mcc_h11_ = source17_.cf6
            d_1319_cf6_ = d_1318___mcc_h11_
            d_1229_v12_ = p0
            if not(not (not((d_1229_v12_) > (d_1229_v12_))) or ((p0) != ((0) - (d_1229_v12_)))):
                r1 = (self).f4
                d_1320_v75_: _dafny.Set
                d_1320_v75_ = _dafny.Set({self.f3})
                d_1321_v76_: _dafny.Set
                d_1321_v76_ = _dafny.Set({(d_1229_v12_) != ((0) - (d_1229_v12_)), (not((self).f4)) not in (d_1320_v75_), (self).f4})
                d_1322_v77_: _dafny.MultiSet
                d_1322_v77_ = _dafny.MultiSet([((d_1234_v16_)[(self).f4] if ((self).f4) in (d_1234_v16_) else (self).f4)])
                d_1321_v76_ = _dafny.Set({(self).f4, self.f3, self.f3, (d_1322_v77_).ispropersubset(_dafny.MultiSet(d_1233_v15_)), (self).f4})
                d_1323_v78_: _dafny.Array
                def lambda69_(d_1324_i4_):
                    return _dafny.Set({709})

                init35_ = lambda69_
                nw197_ = _dafny.Array(None, 29)
                for i0_35_ in range(nw197_.length(0)):
                    nw197_[i0_35_] = init35_(i0_35_)
                d_1323_v78_ = nw197_
                d_1325_v79_: D15
                d_1325_v79_ = D15_DC32(d_1323_v78_)
                d_1326_v80_: _dafny.Seq
                d_1326_v80_ = _dafny.SeqWithoutIsStrInference([d_1323_v78_, d_1323_v78_, d_1323_v78_])
                d_1327_v81_: _dafny.Array
                nw198_ = _dafny.Array(None, 26)
                nw198_[int(0)] = (d_1323_v78_ if (self).f4 else d_1323_v78_)
                nw198_[int(1)] = (d_1325_v79_).cf46
                nw198_[int(2)] = d_1323_v78_
                nw198_[int(3)] = d_1323_v78_
                nw198_[int(4)] = d_1323_v78_
                nw198_[int(5)] = d_1323_v78_
                nw198_[int(6)] = d_1323_v78_
                nw198_[int(7)] = d_1323_v78_
                nw198_[int(8)] = d_1323_v78_
                nw198_[int(9)] = d_1323_v78_
                nw198_[int(10)] = d_1323_v78_
                nw198_[int(11)] = d_1323_v78_
                nw198_[int(12)] = d_1323_v78_
                nw198_[int(13)] = d_1323_v78_
                nw198_[int(14)] = d_1323_v78_
                nw198_[int(15)] = d_1323_v78_
                nw198_[int(16)] = d_1323_v78_
                nw198_[int(17)] = (d_1326_v80_)[default__.safeIndex((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0 for d_1328_i5_ in range(default__.abs(863))]))).cardinality, len(d_1326_v80_))]
                nw198_[int(18)] = d_1323_v78_
                nw198_[int(19)] = d_1323_v78_
                nw198_[int(20)] = d_1323_v78_
                nw198_[int(21)] = d_1323_v78_
                nw198_[int(22)] = d_1323_v78_
                nw198_[int(23)] = d_1323_v78_
                nw198_[int(24)] = d_1323_v78_
                nw198_[int(25)] = (d_1323_v78_ if (self).f4 else d_1323_v78_)
                d_1327_v81_ = nw198_
                d_1329_v82_: _dafny.Map
                d_1329_v82_ = _dafny.Map({self.f3: d_1323_v78_})
                index211_ = default__.safeIndex(716, (d_1327_v81_).length(0))
                (d_1327_v81_)[index211_] = ((d_1329_v82_)[(self).f4] if ((self).f4) in (d_1329_v82_) else d_1323_v78_)
                d_1330_v83_: _dafny.Array
                nw199_ = _dafny.Array(_dafny.Seq({}), 27)
                d_1330_v83_ = nw199_
                d_1331_v84_: _dafny.Set
                d_1331_v84_ = _dafny.Set({p0})
                d_1332_v85_: _dafny.Seq
                d_1332_v85_ = _dafny.SeqWithoutIsStrInference([len(default__.fm34(p0, p0, globalState))])
                index212_ = default__.safeIndex(774, (d_1330_v83_).length(0))
                (d_1330_v83_)[index212_] = (_dafny.SeqWithoutIsStrInference([len(d_1331_v84_), (0) - (d_1229_v12_), p0, 33, len(d_1233_v15_)])) + (d_1332_v85_)
                index213_ = default__.safeIndex(716, (d_1327_v81_).length(0))
                index214_ = default__.safeIndex(774, (d_1330_v83_).length(0))
                rhs185_ = d_1323_v78_
                rhs186_ = _dafny.SeqWithoutIsStrInference([p0 for d_1333_i6_ in range(default__.abs(959))])
                lhs126_ = d_1327_v81_
                lhs127_ = default__.safeIndex(716, (d_1327_v81_).length(0))
                lhs128_ = d_1330_v83_
                lhs129_ = default__.safeIndex(774, (d_1330_v83_).length(0))
                lhs126_[lhs127_] = rhs185_
                lhs128_[lhs129_] = rhs186_
                def lambda70_(d_1334_p0_):
                    def lambda71_(d_1335_i7_):
                        return (d_1335_i7_) * (d_1334_p0_)

                    return lambda71_

                init36_ = lambda70_(p0)
                nw200_ = _dafny.Array(None, 9)
                for i0_36_ in range(nw200_.length(0)):
                    nw200_[i0_36_] = init36_(i0_36_)
                d_1236_v19_ = nw200_
                d_1336_v86_: _dafny.Array
                nw201_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_1336_v86_ = nw201_
                d_1336_v86_ = d_1336_v86_
            elif True:
                d_1319_cf6_ = d_1319_cf6_
                d_1230_v13_ = (d_1230_v13_) + ((d_1230_v13_).set(default__.safeIndex(len((_dafny.Map({True: len(d_1233_v15_)})).set(self.f3, p0)), len(d_1230_v13_)), d_1254_v33_))
                index215_ = default__.safeIndex(715, (d_1236_v19_).length(0))
                (d_1236_v19_)[index215_] = (p0) * (p0)
                d_1337_v87_: D14
                d_1337_v87_ = D14_DC31((0) - (p0), p0)
                index216_ = default__.safeIndex(715, (d_1236_v19_).length(0))
                (d_1236_v19_)[index216_] = (d_1337_v87_).cf44
                r1 = True
                d_1338_v88_: _dafny.Array
                nw202_ = _dafny.Array(None, 17)
                nw202_[int(0)] = d_1235_v17_
                nw202_[int(1)] = (d_1235_v17_) | (d_1235_v17_)
                nw202_[int(2)] = d_1235_v17_
                nw202_[int(3)] = (_dafny.Map({self.f3: p0})) | (_dafny.Map({(self).f4: p0}))
                nw202_[int(4)] = d_1235_v17_
                nw202_[int(5)] = d_1235_v17_
                nw202_[int(6)] = (d_1235_v17_).set((self).f4, (d_1236_v19_)[default__.safeIndex(715, (d_1236_v19_).length(0))])
                nw202_[int(7)] = d_1235_v17_
                nw202_[int(8)] = (d_1235_v17_) | (_dafny.Map({False: d_1229_v12_}))
                nw202_[int(9)] = default__.fm28(False, self.f3, self.f3, len(d_1230_v13_), globalState)
                nw202_[int(10)] = d_1235_v17_
                nw202_[int(11)] = (d_1235_v17_).set((self).f4, 848)
                nw202_[int(12)] = (D16_DC35(d_1235_v17_)).cf50
                nw202_[int(13)] = d_1235_v17_
                nw202_[int(14)] = d_1235_v17_
                nw202_[int(15)] = (_dafny.Map({self.f3: (d_1236_v19_)[default__.safeIndex(715, (d_1236_v19_).length(0))]})).set((self).f4, (d_1236_v19_)[default__.safeIndex(715, (d_1236_v19_).length(0))])
                nw202_[int(16)] = d_1235_v17_
                d_1338_v88_ = nw202_
                index217_ = default__.safeIndex(958, (d_1338_v88_).length(0))
                (d_1338_v88_)[index217_] = d_1235_v17_
                index218_ = default__.safeIndex(958, (d_1338_v88_).length(0))
                (d_1338_v88_)[index218_] = d_1235_v17_
            d_1339_v89_: _dafny.Map
            d_1339_v89_ = _dafny.Map({default__.safeDivisionInt(d_1229_v12_, p0): _dafny.MultiSet([self.f3])})
            index219_ = default__.safeIndex(152, (d_1236_v19_).length(0))
            (d_1236_v19_)[index219_] = d_1229_v12_
            index220_ = default__.safeIndex(490, (d_1236_v19_).length(0))
            (d_1236_v19_)[index220_] = p0
            index221_ = default__.safeIndex(152, (d_1236_v19_).length(0))
            index222_ = default__.safeIndex(490, (d_1236_v19_).length(0))
            rhs187_ = 954
            rhs188_ = _dafny.Map({len(d_1233_v15_): _dafny.MultiSet([True, (self).f4])})
            rhs189_ = (d_1229_v12_) * ((d_1229_v12_) * (p0))
            rhs190_ = ((d_1229_v12_) * (p0)) - ((p0) + (d_1229_v12_))
            rhs191_ = d_1229_v12_
            lhs130_ = d_1236_v19_
            lhs131_ = default__.safeIndex(152, (d_1236_v19_).length(0))
            lhs132_ = d_1236_v19_
            lhs133_ = default__.safeIndex(490, (d_1236_v19_).length(0))
            d_1229_v12_ = rhs187_
            d_1339_v89_ = rhs188_
            lhs130_[lhs131_] = rhs189_
            lhs132_[lhs133_] = rhs190_
            d_1229_v12_ = rhs191_
            d_1340_v90_: C0
            nw203_ = C0()
            nw203_.ctor__()
            d_1340_v90_ = nw203_
        r0 = default__.fm14(globalState)
        d_1341_v91_: _dafny.Set
        d_1341_v91_ = _dafny.Set({(self).f4})
        r1 = not ((d_1341_v91_).isdisjoint(d_1341_v91_)) or ((p0) <= (p0))
        return r0, r1

    def m8(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_1342_i0_: int
        d_1342_i0_ = 0
        with _dafny.label("6"):
            while (default__.safeModuloInt(p1, p0)) <= ((self).fm0(globalState)):
                with _dafny.c_label("6"):
                    if (d_1342_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1342_i0_ = (d_1342_i0_) + (1)
                    d_1343_v0_: _dafny.Seq
                    d_1343_v0_ = _dafny.SeqWithoutIsStrInference([(self).f4, self.f3])
                    d_1344_v1_: D1
                    d_1344_v1_ = D1_DC5((self).f4, (self).f4, p1, p1)
                    d_1345_v2_: _dafny.Array
                    nw204_ = _dafny.Array(None, 7)
                    nw204_[int(0)] = p0
                    nw204_[int(1)] = default__.safeDivisionInt((0) - (p0), p1)
                    nw204_[int(2)] = (p1 if (self).f4 else 631)
                    nw204_[int(3)] = p0
                    nw204_[int(4)] = (default__.fm31(False, p0, (self).f4, d_1343_v0_, globalState)) * (p0)
                    nw204_[int(5)] = len((d_1343_v0_) + (d_1343_v0_))
                    nw204_[int(6)] = len(_dafny.Map({d_1344_v1_: (self).f4}))
                    d_1345_v2_ = nw204_
                    d_1346_v3_: _dafny.Map
                    d_1346_v3_ = _dafny.Map({(self).f4: 83})
                    index223_ = default__.safeIndex(296, (d_1345_v2_).length(0))
                    (d_1345_v2_)[index223_] = (len(d_1346_v3_)) * (84)
                    index224_ = default__.safeIndex(296, (d_1345_v2_).length(0))
                    (d_1345_v2_)[index224_] = p0
                    (self).f3 = not (not(False)) or ((self).f4)
                    d_1347_v4_: C2
                    nw205_ = C2()
                    nw205_.ctor__((self).f4, self.f3, (self).f4, (self).f4)
                    d_1347_v4_ = nw205_
                    d_1348_v5_: _dafny.Set
                    d_1348_v5_ = _dafny.Set({len(d_1346_v3_), p1})
                    d_1349_v6_: _dafny.Set
                    d_1349_v6_ = d_1348_v5_
                    d_1349_v6_ = d_1349_v6_
                    pass
            pass
        d_1350_v7_: _dafny.Map
        d_1350_v7_ = _dafny.Map({260: self.f3})
        d_1351_v8_: _dafny.Seq
        d_1351_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihae"))
        d_1352_v9_: _dafny.Seq
        d_1352_v9_ = _dafny.SeqWithoutIsStrInference([False])
        def iife104_():
            coll48_ = _dafny.Map()
            compr_48_: int
            for compr_48_ in (_dafny.Set({p0, p0, p1, len(d_1351_v8_), len(d_1351_v8_)})).Elements:
                d_1353_v10_: int = compr_48_
                if (d_1353_v10_) in (_dafny.Set({p0, p0, p1, len(d_1351_v8_), len(d_1351_v8_)})):
                    coll48_[(d_1353_v10_) + (p0)] = self.f3
            return _dafny.Map(coll48_)
        d_1350_v7_ = (d_1350_v7_) | ((_dafny.Map({default__.fm31(self.f3, len(d_1351_v8_), self.f3, d_1352_v9_, globalState): True})) | (iife104_()
        ))
        d_1354_v11_: int
        d_1354_v11_ = -243
        d_1354_v11_ = default__.safeDivisionInt(d_1354_v11_, (-33) * (p1))
        d_1355_v12_: _dafny.Array
        def lambda72_(d_1356_i2_):
            return _dafny.Map({True: self.f3})

        init37_ = lambda72_
        nw206_ = _dafny.Array(None, 14)
        for i0_37_ in range(nw206_.length(0)):
            nw206_[i0_37_] = init37_(i0_37_)
        d_1355_v12_ = nw206_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1355_v12_).length(0)):
            d_1357_i1_: int = guard_loop_3_
            if (True) and (((0) <= (d_1357_i1_)) and ((d_1357_i1_) < ((d_1355_v12_).length(0)))):
                (d_1355_v12_)[(d_1357_i1_)] = ((_dafny.Map({(self).f4: self.f3})) | (_dafny.Map({(self).f4: False})) if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p1 for d_1358_i3_ in range(default__.abs(296))]))).issubset(_dafny.MultiSet([d_1354_v11_])) else (_dafny.Map({(self).f4: self.f3})) | (_dafny.Map({(self).f4: self.f3})))
        d_1359_v13_: _dafny.Seq
        d_1359_v13_ = _dafny.SeqWithoutIsStrInference([p0, (self).fm0(globalState), p1, d_1354_v11_, d_1354_v11_])
        d_1360_v14_: _dafny.MultiSet
        d_1360_v14_ = _dafny.MultiSet([default__.safeModuloInt(len(d_1359_v13_), 55)])
        d_1361_v15_: _dafny.Set
        d_1361_v15_ = _dafny.Set({-890})
        d_1360_v14_ = ((d_1360_v14_) - (default__.fm24(globalState))).set(len(d_1361_v15_), default__.abs(498))
        if self.f3:
            d_1354_v11_ = (p1) - (d_1354_v11_)
            d_1362_v16_: _dafny.Array
            def lambda73_(d_1363_i4_):
                return self.f3

            init38_ = lambda73_
            nw207_ = _dafny.Array(None, 1)
            for i0_38_ in range(nw207_.length(0)):
                nw207_[i0_38_] = init38_(i0_38_)
            d_1362_v16_ = nw207_
            index225_ = default__.safeIndex(739, (d_1362_v16_).length(0))
            (d_1362_v16_)[index225_] = default__.fm14(globalState)
            d_1364_v17_: _dafny.Map
            d_1364_v17_ = _dafny.Map({self.f3: self.f3})
            d_1365_v18_: _dafny.Map
            d_1365_v18_ = _dafny.Map({664: len(d_1364_v17_)})
            d_1366_v19_: _dafny.MultiSet
            d_1366_v19_ = _dafny.MultiSet([d_1365_v18_, d_1365_v18_])
            index226_ = default__.safeIndex(739, (d_1362_v16_).length(0))
            (d_1362_v16_)[index226_] = default__.fm21((p0) * (p0), (_dafny.MultiSet([d_1365_v18_])).intersection(d_1366_v19_), globalState)
            if (d_1362_v16_)[default__.safeIndex(739, (d_1362_v16_).length(0))]:
                d_1367_v20_: _dafny.Array
                def lambda74_(d_1368_v9_):
                    def lambda75_(d_1369_i5_):
                        return d_1368_v9_

                    return lambda75_

                init39_ = lambda74_(d_1352_v9_)
                nw208_ = _dafny.Array(None, 23)
                for i0_39_ in range(nw208_.length(0)):
                    nw208_[i0_39_] = init39_(i0_39_)
                d_1367_v20_ = nw208_
                index227_ = default__.safeIndex(516, (d_1367_v20_).length(0))
                (d_1367_v20_)[index227_] = d_1352_v9_
                index228_ = default__.safeIndex(516, (d_1367_v20_).length(0))
                (d_1367_v20_)[index228_] = d_1352_v9_
                d_1370_v21_: _dafny.MultiSet
                d_1370_v21_ = _dafny.MultiSet([(d_1362_v16_)[default__.safeIndex(739, (d_1362_v16_).length(0))]])
                (globalState).f0 = (True if (((d_1370_v21_)[(d_1362_v16_)[default__.safeIndex(739, (d_1362_v16_).length(0))]] if ((d_1362_v16_)[default__.safeIndex(739, (d_1362_v16_).length(0))]) in (d_1370_v21_) else d_1354_v11_)) == (len(d_1352_v9_)) else (self).f4)
                d_1362_v16_ = d_1362_v16_
                (globalState).f0 = ((d_1362_v16_)[default__.safeIndex(739, (d_1362_v16_).length(0))]) or (default__.fm14(globalState))
                d_1354_v11_ = 808
            elif True:
                d_1371_v23_: _dafny.Set
                d_1371_v23_ = _dafny.Set({(d_1362_v16_)[default__.safeIndex(739, (d_1362_v16_).length(0))]})
                d_1372_v24_: str
                d_1372_v24_ = _dafny.CodePoint('n')
                d_1373_v25_: _dafny.Map
                d_1373_v25_ = _dafny.Map({d_1372_v24_: d_1371_v23_})
                d_1374_v26_: _dafny.MultiSet
                d_1374_v26_ = _dafny.MultiSet([d_1359_v13_, d_1359_v13_, d_1359_v13_, d_1359_v13_])
                def iife105_():
                    coll49_ = _dafny.Map()
                    compr_49_: int
                    for compr_49_ in (_dafny.SeqWithoutIsStrInference([206])).Elements:
                        d_1375_v22_: int = compr_49_
                        if (d_1375_v22_) in (_dafny.SeqWithoutIsStrInference([206])):
                            coll49_[default__.safeDivisionInt(d_1375_v22_, (0) - (p0))] = not(self.f3)
                    return _dafny.Map(coll49_)
                (globalState).f0 = ((((d_1373_v25_)[d_1372_v24_] if (d_1372_v24_) in (d_1373_v25_) else d_1371_v23_)).issubset(d_1371_v23_) if (p1) in (iife105_()
                ) else (d_1374_v26_).isdisjoint(d_1374_v26_))
                index229_ = default__.safeIndex(739, (d_1362_v16_).length(0))
                index230_ = default__.safeIndex(739, (d_1362_v16_).length(0))
                rhs192_ = (734) + (p1)
                rhs193_ = (d_1362_v16_)[default__.safeIndex(739, (d_1362_v16_).length(0))]
                rhs194_ = (self).f4
                rhs195_ = d_1372_v24_
                lhs134_ = d_1362_v16_
                lhs135_ = default__.safeIndex(739, (d_1362_v16_).length(0))
                lhs136_ = d_1362_v16_
                lhs137_ = default__.safeIndex(739, (d_1362_v16_).length(0))
                d_1354_v11_ = rhs192_
                lhs134_[lhs135_] = rhs193_
                lhs136_[lhs137_] = rhs194_
                d_1372_v24_ = rhs195_
                d_1376_v27_: _dafny.Array
                def lambda76_(d_1377_v11_):
                    def lambda77_(d_1378_i6_):
                        return (d_1378_i6_) * (d_1377_v11_)

                    return lambda77_

                init40_ = lambda76_(d_1354_v11_)
                nw209_ = _dafny.Array(None, 15)
                for i0_40_ in range(nw209_.length(0)):
                    nw209_[i0_40_] = init40_(i0_40_)
                d_1376_v27_ = nw209_
                index231_ = default__.safeIndex(246, (d_1376_v27_).length(0))
                (d_1376_v27_)[index231_] = (0) - (p0)
                d_1379_v28_: _dafny.Map
                d_1379_v28_ = _dafny.Map({(d_1360_v14_).ispropersubset((d_1360_v14_).set(d_1354_v11_, default__.abs(211))): d_1354_v11_})
                index232_ = default__.safeIndex(246, (d_1376_v27_).length(0))
                (d_1376_v27_)[index232_] = len(d_1379_v28_)
                index233_ = default__.safeIndex(739, (d_1362_v16_).length(0))
                (d_1362_v16_)[index233_] = self.f3
                index234_ = default__.safeIndex(246, (d_1376_v27_).length(0))
                (d_1376_v27_)[index234_] = ((d_1376_v27_)[default__.safeIndex(246, (d_1376_v27_).length(0))]) - ((d_1376_v27_)[default__.safeIndex(246, (d_1376_v27_).length(0))])
            d_1380_v29_: C0
            nw210_ = C0()
            nw210_.ctor__()
            d_1380_v29_ = nw210_
            d_1381_v30_: _dafny.Array
            nw211_ = _dafny.Array(_dafny.Seq({}), 22)
            d_1381_v30_ = nw211_
            d_1382_v31_: _dafny.Map
            d_1382_v31_ = _dafny.Map({(self).f4: d_1381_v30_})
            d_1381_v30_ = (((d_1382_v31_)[self.f3] if (self.f3) in (d_1382_v31_) else d_1381_v30_) if (d_1362_v16_)[default__.safeIndex(739, (d_1362_v16_).length(0))] else ((d_1382_v31_)[True] if (True) in (d_1382_v31_) else d_1381_v30_))
        elif True:
            d_1383_v32_: str
            d_1383_v32_ = _dafny.CodePoint('i')
            d_1383_v32_ = d_1383_v32_
            d_1354_v11_ = p0
            d_1384_v33_: _dafny.Array
            nw212_ = _dafny.Array(False, 4)
            d_1384_v33_ = nw212_
            index235_ = default__.safeIndex(2, (d_1384_v33_).length(0))
            (d_1384_v33_)[index235_] = self.f3
            d_1385_v34_: _dafny.Set
            d_1385_v34_ = _dafny.Set({(self).f4})
            d_1386_v35_: _dafny.Map
            d_1386_v35_ = _dafny.Map({(0) - (p1): d_1385_v34_})
            d_1387_v36_: D7
            d_1387_v36_ = D7_DC16((d_1386_v35_) | (d_1386_v35_))
            d_1388_v37_: _dafny.Set
            d_1388_v37_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhxyt"))})
            index236_ = default__.safeIndex(2, (d_1384_v33_).length(0))
            rhs196_ = ((d_1388_v37_) | (d_1388_v37_)) == (d_1388_v37_)
            rhs197_ = 145
            rhs198_ = d_1387_v36_
            lhs138_ = d_1384_v33_
            lhs139_ = default__.safeIndex(2, (d_1384_v33_).length(0))
            lhs138_[lhs139_] = rhs196_
            d_1354_v11_ = rhs197_
            d_1387_v36_ = rhs198_
            d_1389_v38_: _dafny.Map
            d_1389_v38_ = _dafny.Map({d_1352_v9_: p1})
            d_1390_v39_: D7
            d_1390_v39_ = D7_DC17(default__.fm31(self.f3, len(d_1389_v38_), self.f3, d_1352_v9_, globalState), not(self.f3), (d_1384_v33_)[default__.safeIndex(2, (d_1384_v33_).length(0))], (_dafny.MultiSet([(d_1384_v33_)[default__.safeIndex(2, (d_1384_v33_).length(0))]])).cardinality)
            d_1391_v40_: C2
            nw213_ = C2()
            nw213_.ctor__((d_1384_v33_)[default__.safeIndex(2, (d_1384_v33_).length(0))], False, self.f3, (d_1390_v39_).cf27)
            d_1391_v40_ = nw213_
            d_1392_v41_: _dafny.Map
            d_1392_v41_ = _dafny.Map({p0: d_1391_v40_})
            d_1393_v42_: _dafny.Map
            d_1393_v42_ = _dafny.Map({p1: d_1354_v11_})
            d_1394_v43_: _dafny.Map
            d_1394_v43_ = _dafny.Map({((d_1392_v41_)[p0] if (p0) in (d_1392_v41_) else d_1391_v40_): d_1393_v42_})
            d_1394_v43_ = (d_1394_v43_).set(d_1391_v40_, d_1393_v42_)
            d_1395_v44_: C0
            nw214_ = C0()
            nw214_.ctor__()
            d_1395_v44_ = nw214_
            d_1395_v44_ = d_1395_v44_
        r0 = d_1352_v9_
        d_1396_v45_: _dafny.Array
        nw215_ = _dafny.Array(int(0), 13)
        d_1396_v45_ = nw215_
        r1 = d_1396_v45_
        return r0, r1


class C4(T1, T0):
    def  __init__(self):
        self._f3: bool = False
        self._f4: bool = False
        self.f8: bool = False
        self._f9: int = int(0)
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
    def f4(self):
        return self._f4
    def ctor__(self, f8, f9, f3, f4):
        (self).f8 = f8
        (self)._f9 = f9
        (self).f3 = f3
        (self)._f4 = f4

    def fm0(self, globalState):
        return (self).f9

    def fm5(self, p0, p1, globalState):
        return (D0_DC2((self).f4, (self).f9, (self).f9)).cf4

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        d_1397_v0_: str
        d_1397_v0_ = _dafny.CodePoint('j')
        d_1398_v1_: _dafny.Seq
        d_1398_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), d_1397_v0_, d_1397_v0_, d_1397_v0_])
        if ((p1).f4) == (default__.fm6((self).fm5(False, d_1398_v1_, globalState), globalState)):
            d_1399_v2_: D0
            d_1399_v2_ = D0_DC2(not((p1).f4), (self).f9, p0)
            d_1400_v3_: _dafny.Seq
            d_1400_v3_ = _dafny.SeqWithoutIsStrInference([d_1399_v2_, (d_1399_v2_ if self.f3 else D0_DC2(self.f3, p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irpitaca"))))), d_1399_v2_, d_1399_v2_])
            d_1400_v3_ = (d_1400_v3_) + (d_1400_v3_)
            (globalState).f0 = True
            d_1401_v4_: _dafny.Map
            d_1401_v4_ = _dafny.Map({(self).f9: (self).f9})
            d_1402_v5_: _dafny.MultiSet
            d_1402_v5_ = _dafny.MultiSet([d_1401_v4_])
            d_1403_v6_: C3
            nw216_ = C3()
            nw216_.ctor__(default__.fm21((self).f9, d_1402_v5_, globalState), not(not((p1.f3 if not(p1.f3) else False))))
            d_1403_v6_ = nw216_
            r0 = p0
            d_1404_v7_: T0
            nw217_ = C3()
            nw217_.ctor__(self.f8, (p1).f4)
            d_1404_v7_ = nw217_
            d_1404_v7_ = p1
        elif True:
            d_1405_v8_: _dafny.Seq
            d_1405_v8_ = _dafny.SeqWithoutIsStrInference([(p1).f4, p1.f3, p1.f3])
            d_1406_v9_: _dafny.Map
            d_1406_v9_ = _dafny.Map({self.f8: p1.f3})
            r0 = default__.fm31(False, (self).f9, (d_1405_v8_)[default__.safeIndex(len(d_1406_v9_), len(d_1405_v8_))], d_1405_v8_, globalState)
            d_1407_v10_: _dafny.Array
            def lambda78_(d_1408_v1_, d_1409_v0_, d_1410_p0_):
                def lambda79_(d_1411_i0_):
                    return (_dafny.Map({(d_1408_v1_).set(default__.safeIndex(569, len(d_1408_v1_)), d_1409_v0_): d_1410_p0_})) | (_dafny.Map({d_1408_v1_: (self).f9}))

                return lambda79_

            init41_ = lambda78_(d_1398_v1_, d_1397_v0_, p0)
            nw218_ = _dafny.Array(None, 10)
            for i0_41_ in range(nw218_.length(0)):
                nw218_[i0_41_] = init41_(i0_41_)
            d_1407_v10_ = nw218_
            index237_ = default__.safeIndex(924, (d_1407_v10_).length(0))
            (d_1407_v10_)[index237_] = _dafny.Map({d_1398_v1_: 608})
            d_1412_v11_: _dafny.Set
            d_1412_v11_ = _dafny.Set({self.f8, p1.f3, (p1).f4, self.f8})
            d_1413_v12_: _dafny.Map
            d_1413_v12_ = _dafny.Map({(0) - (p0): d_1412_v11_})
            d_1414_v13_: D7
            d_1414_v13_ = D7_DC16(d_1413_v12_)
            d_1415_v14_: _dafny.Array
            def lambda80_(d_1416_p1_, d_1417_v1_):
                def lambda81_(d_1418_i1_):
                    return (_dafny.Map({d_1416_p1_.f3: d_1417_v1_}) if (d_1416_p1_).f4 else _dafny.Map({(d_1416_p1_).f4: d_1417_v1_}))

                return lambda81_

            init42_ = lambda80_(p1, d_1398_v1_)
            nw219_ = _dafny.Array(None, 14)
            for i0_42_ in range(nw219_.length(0)):
                nw219_[i0_42_] = init42_(i0_42_)
            d_1415_v14_ = nw219_
            d_1419_v15_: _dafny.Map
            d_1419_v15_ = _dafny.Map({self.f8: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oh"))})
            index238_ = default__.safeIndex(99, (d_1415_v14_).length(0))
            (d_1415_v14_)[index238_] = d_1419_v15_
            d_1420_v16_: _dafny.Array
            nw220_ = _dafny.Array(int(0), 27)
            d_1420_v16_ = nw220_
            d_1421_v17_: D4
            d_1421_v17_ = D4_DC12((self).f9, (self).f9, (self).f4)
            index239_ = default__.safeIndex(627, (d_1420_v16_).length(0))
            (d_1420_v16_)[index239_] = (d_1421_v17_).cf20
            d_1422_v18_: _dafny.Map
            d_1422_v18_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnwad"))) + (d_1398_v1_): (p0) - ((self).f9)})
            d_1423_v19_: _dafny.MultiSet
            d_1423_v19_ = _dafny.MultiSet([(self).f4, (p1).f4])
            d_1424_v20_: _dafny.Seq
            d_1424_v20_ = _dafny.SeqWithoutIsStrInference([p0, (0) - ((d_1423_v19_).cardinality), (self).f9, 145, (self).f9])
            d_1425_v21_: _dafny.Map
            d_1425_v21_ = _dafny.Map({(self).f9: (self).f9})
            d_1426_v22_: _dafny.Map
            d_1426_v22_ = _dafny.Map({(p1).f4: d_1425_v21_})
            d_1427_v23_: _dafny.Seq
            d_1427_v23_ = _dafny.SeqWithoutIsStrInference([d_1398_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwnoak"))])
            pat_let_tv40_ = d_1414_v13_
            index240_ = default__.safeIndex(924, (d_1407_v10_).length(0))
            index241_ = default__.safeIndex(99, (d_1415_v14_).length(0))
            index242_ = default__.safeIndex(627, (d_1420_v16_).length(0))
            def iife106_(_pat_let28_0):
                def iife107_(d_1428_dt__update__tmp_h0_):
                    def iife108_(_pat_let29_0):
                        def iife109_(d_1429_dt__update_hcf25_h0_):
                            return D7_DC16(d_1429_dt__update_hcf25_h0_)
                        return iife109_(_pat_let29_0)
                    return iife108_((pat_let_tv40_).cf25)
                return iife107_(_pat_let28_0)
            rhs199_ = d_1422_v18_
            rhs200_ = iife106_(d_1414_v13_)
            rhs201_ = ((p1).f4) not in (d_1423_v19_)
            rhs202_ = (_dafny.Map({default__.fm6(len(d_1412_v11_), globalState): d_1398_v1_})) | (d_1419_v15_)
            rhs203_ = len((((_dafny.SeqWithoutIsStrInference([(self).f9 for d_1430_i2_ in range(default__.abs(13))])) + ((d_1424_v20_) + (_dafny.SeqWithoutIsStrInference([(self).f9 for d_1431_i3_ in range(default__.abs(987))])))).set(default__.safeIndex((self).fm5((self).f4, d_1398_v1_, globalState), len((_dafny.SeqWithoutIsStrInference([(self).f9 for d_1432_i2_ in range(default__.abs(13))])) + ((d_1424_v20_) + (_dafny.SeqWithoutIsStrInference([(self).f9 for d_1433_i3_ in range(default__.abs(987))]))))), (len(d_1426_v22_)) * (len((d_1427_v23_))))).set(default__.safeIndex(((self).f9) * (p0), len(((_dafny.SeqWithoutIsStrInference([(self).f9 for d_1434_i2_ in range(default__.abs(13))])) + ((d_1424_v20_) + (_dafny.SeqWithoutIsStrInference([(self).f9 for d_1435_i3_ in range(default__.abs(987))])))).set(default__.safeIndex((self).fm5((self).f4, d_1398_v1_, globalState), len((_dafny.SeqWithoutIsStrInference([(self).f9 for d_1436_i2_ in range(default__.abs(13))])) + ((d_1424_v20_) + (_dafny.SeqWithoutIsStrInference([(self).f9 for d_1437_i3_ in range(default__.abs(987))]))))), (len(d_1426_v22_)) * (len((d_1427_v23_)))))), p0))
            lhs140_ = d_1407_v10_
            lhs141_ = default__.safeIndex(924, (d_1407_v10_).length(0))
            lhs142_ = p1
            lhs143_ = d_1415_v14_
            lhs144_ = default__.safeIndex(99, (d_1415_v14_).length(0))
            lhs145_ = d_1420_v16_
            lhs146_ = default__.safeIndex(627, (d_1420_v16_).length(0))
            lhs140_[lhs141_] = rhs199_
            d_1414_v13_ = rhs200_
            lhs142_.f3 = rhs201_
            lhs143_[lhs144_] = rhs202_
            lhs145_[lhs146_] = rhs203_
            d_1398_v1_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrmn")) if (d_1397_v0_) not in (d_1398_v1_) else d_1398_v1_)
            (globalState).f0 = (default__.safeModuloInt((self).f9, p0)) > (p0)
            d_1438_v24_: _dafny.Map
            d_1438_v24_ = _dafny.Map({(d_1420_v16_)[default__.safeIndex(627, (d_1420_v16_).length(0))]: p1.f3})
            d_1439_v25_: _dafny.Map
            d_1439_v25_ = _dafny.Map({((d_1438_v24_).set(p0, (p1).f4) if (self).f4 else d_1438_v24_): (self).f9})
            d_1439_v25_ = (d_1439_v25_) | (d_1439_v25_)
        r0 = len(d_1398_v1_)
        hi12_ = ((self).f9) + (91)
        for d_1440_i4_ in range(p0, hi12_):
            if True:
                d_1441_v26_: _dafny.Array
                def lambda82_(d_1442_p1_):
                    def lambda83_(d_1443_i5_):
                        return (d_1442_p1_).f4

                    return lambda83_

                init43_ = lambda82_(p1)
                nw221_ = _dafny.Array(None, 29)
                for i0_43_ in range(nw221_.length(0)):
                    nw221_[i0_43_] = init43_(i0_43_)
                d_1441_v26_ = nw221_
                index243_ = default__.safeIndex(752, (d_1441_v26_).length(0))
                (d_1441_v26_)[index243_] = p1.f3
                d_1444_v27_: _dafny.Array
                def lambda84_(d_1445_v0_):
                    def lambda85_(d_1446_i6_):
                        return d_1445_v0_

                    return lambda85_

                init44_ = lambda84_(d_1397_v0_)
                nw222_ = _dafny.Array(None, 25)
                for i0_44_ in range(nw222_.length(0)):
                    nw222_[i0_44_] = init44_(i0_44_)
                d_1444_v27_ = nw222_
                d_1447_v28_: _dafny.MultiSet
                d_1447_v28_ = _dafny.MultiSet([self.f3, not(p1.f3), (self).f4, self.f8])
                d_1448_v29_: _dafny.Map
                d_1448_v29_ = _dafny.Map({d_1444_v27_: (d_1447_v28_) == (d_1447_v28_)})
                index244_ = default__.safeIndex(752, (d_1441_v26_).length(0))
                (d_1441_v26_)[index244_] = ((d_1448_v29_)[d_1444_v27_] if (d_1444_v27_) in (d_1448_v29_) else not((self).f4))
                d_1447_v28_ = _dafny.MultiSet([(p1).f4, (d_1441_v26_)[default__.safeIndex(752, (d_1441_v26_).length(0))], self.f8])
                d_1449_v30_: _dafny.Seq
                d_1449_v30_ = _dafny.SeqWithoutIsStrInference([self.f3, p1.f3])
                d_1450_v31_: _dafny.Seq
                d_1450_v31_ = _dafny.SeqWithoutIsStrInference([self.f8, (d_1449_v30_)[default__.safeIndex(p0, len(d_1449_v30_))]])
                d_1450_v31_ = d_1450_v31_
                (self).f3 = (p1).f4
                d_1451_v32_: _dafny.Array
                nw223_ = _dafny.Array(int(0), 2)
                d_1451_v32_ = nw223_
                index245_ = default__.safeIndex(841, (d_1451_v32_).length(0))
                (d_1451_v32_)[index245_] = p0
                d_1452_v33_: _dafny.Map
                d_1452_v33_ = _dafny.Map({(self).f4: (d_1398_v1_)[default__.safeIndex(d_1440_i4_, len(d_1398_v1_))]})
                d_1453_v34_: _dafny.Map
                d_1453_v34_ = _dafny.Map({((d_1452_v33_)[(p1).f4] if ((p1).f4) in (d_1452_v33_) else _dafny.CodePoint('a')): (d_1398_v1_) + (d_1398_v1_)})
                index246_ = default__.safeIndex(841, (d_1451_v32_).length(0))
                rhs204_ = (p1).f4
                rhs205_ = 175
                rhs206_ = len(d_1453_v34_)
                rhs207_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
                lhs147_ = globalState
                lhs148_ = d_1451_v32_
                lhs149_ = default__.safeIndex(841, (d_1451_v32_).length(0))
                lhs147_.f0 = rhs204_
                r0 = rhs205_
                lhs148_[lhs149_] = rhs206_
                d_1398_v1_ = rhs207_
            elif True:
                r0 = (p1).fm0(globalState)
                d_1454_v35_: _dafny.Array
                def lambda86_(d_1455_i7_):
                    return _dafny.Set({(self).f9, len(_dafny.Map({(self).f4: self.f3}))})

                init45_ = lambda86_
                nw224_ = _dafny.Array(None, 17)
                for i0_45_ in range(nw224_.length(0)):
                    nw224_[i0_45_] = init45_(i0_45_)
                d_1454_v35_ = nw224_
                d_1456_v36_: D15
                d_1456_v36_ = D15_DC32(d_1454_v35_)
                d_1457_v37_: D15
                d_1457_v37_ = D15_DC34(d_1456_v36_)
                d_1458_v38_: D15
                d_1458_v38_ = D15_DC34(d_1457_v37_)
                d_1458_v38_ = d_1458_v38_
                r0 = p0
                d_1459_v39_: _dafny.Seq
                d_1459_v39_ = _dafny.SeqWithoutIsStrInference([(self).f4, p1.f3, p1.f3, (p1).f4])
                d_1460_v40_: _dafny.MultiSet
                d_1460_v40_ = _dafny.MultiSet([len(_dafny.Map({d_1397_v0_: (p1).f4})), len(d_1459_v39_), (self).f9])
                d_1461_v41_: _dafny.Map
                d_1461_v41_ = _dafny.Map({(self).f9: (self).f9})
                d_1462_v42_: _dafny.Array
                nw225_ = _dafny.Array(int(0), 5)
                d_1462_v42_ = nw225_
                d_1463_v43_: _dafny.Seq
                d_1463_v43_ = _dafny.SeqWithoutIsStrInference([d_1462_v42_])
                d_1464_v44_: _dafny.Seq
                d_1464_v44_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                d_1465_v45_: _dafny.Array
                nw226_ = _dafny.Array(None, 24)
                nw226_[int(0)] = (0) - ((self).f9)
                nw226_[int(1)] = ((d_1460_v40_).cardinality) * (453)
                nw226_[int(2)] = p0
                nw226_[int(3)] = p0
                nw226_[int(4)] = p0
                nw226_[int(5)] = (self).f9
                nw226_[int(6)] = p0
                nw226_[int(7)] = -944
                nw226_[int(8)] = (395) + (len(d_1461_v41_))
                nw226_[int(9)] = len(default__.fm34(d_1440_i4_, (self).f9, globalState))
                nw226_[int(10)] = (self).f9
                nw226_[int(11)] = d_1440_i4_
                nw226_[int(12)] = (len(d_1463_v43_)) + (p0)
                nw226_[int(13)] = (self).f9
                nw226_[int(14)] = (self).f9
                nw226_[int(15)] = default__.safeModuloInt((self).f9, -732)
                nw226_[int(16)] = default__.safeModuloInt(98, (self).f9)
                nw226_[int(17)] = d_1440_i4_
                nw226_[int(18)] = 660
                nw226_[int(19)] = p0
                nw226_[int(20)] = (self).f9
                nw226_[int(21)] = d_1440_i4_
                nw226_[int(22)] = (self).f9
                nw226_[int(23)] = len((d_1459_v39_) + (default__.fm34(len(d_1464_v44_), p0, globalState)))
                d_1465_v45_ = nw226_
                index247_ = default__.safeIndex(7, (d_1462_v42_).length(0))
                (d_1462_v42_)[index247_] = (p0 if p1.f3 else d_1440_i4_)
                index248_ = default__.safeIndex(7, (d_1462_v42_).length(0))
                (d_1462_v42_)[index248_] = p0
                d_1466_v46_: _dafny.MultiSet
                d_1466_v46_ = _dafny.MultiSet([False])
                d_1467_v47_: _dafny.Seq
                d_1467_v47_ = _dafny.SeqWithoutIsStrInference([d_1459_v39_])
                d_1468_v48_: _dafny.Seq
                d_1468_v48_ = _dafny.SeqWithoutIsStrInference([d_1464_v44_, (d_1464_v44_).set(default__.safeIndex((d_1462_v42_)[default__.safeIndex(7, (d_1462_v42_).length(0))], len(d_1464_v44_)), (d_1462_v42_)[default__.safeIndex(7, (d_1462_v42_).length(0))]), d_1464_v44_, _dafny.SeqWithoutIsStrInference([p0])])
                d_1469_v49_: _dafny.Map
                d_1469_v49_ = _dafny.Map({(d_1440_i4_) - (((d_1466_v46_)[self.f8] if (self.f8) in (d_1466_v46_) else default__.fm31(self.f3, (0) - (len(d_1459_v39_)), not((p1).f4), (d_1467_v47_)[default__.safeIndex(p0, len(d_1467_v47_))], globalState))): d_1468_v48_})
                d_1469_v49_ = (d_1469_v49_).set(d_1440_i4_, _dafny.SeqWithoutIsStrInference([d_1464_v44_, d_1464_v44_, d_1464_v44_]))
            if (default__.fm37(d_1440_i4_, globalState)).cf53:
                d_1470_v50_: C2
                nw227_ = C2()
                nw227_.ctor__(p1.f3, p1.f3, self.f3, True)
                d_1470_v50_ = nw227_
                d_1471_v51_: _dafny.Map
                d_1471_v51_ = _dafny.Map({self.f3: (p1).f4})
                d_1472_v52_: D0
                d_1472_v52_ = D0_DC2((d_1470_v50_).f10, p0, (self).f9)
                d_1473_v53_: _dafny.Seq
                d_1473_v53_ = _dafny.SeqWithoutIsStrInference([self.f3, (d_1472_v52_).cf3])
                d_1474_v54_: C1
                nw228_ = C1()
                nw228_.ctor__(d_1398_v1_, ((self).f4 if ((d_1471_v51_)[p1.f3] if (p1.f3) in (d_1471_v51_) else (self).f4) else (d_1473_v53_)[default__.safeIndex((self).f9, len(d_1473_v53_))]))
                d_1474_v54_ = nw228_
                d_1475_v55_: _dafny.Array
                nw229_ = _dafny.Array(None, 10)
                d_1475_v55_ = nw229_
                d_1475_v55_ = d_1475_v55_
                d_1476_v56_: _dafny.Set
                d_1476_v56_ = _dafny.Set({False})
                d_1477_v57_: _dafny.MultiSet
                d_1477_v57_ = _dafny.MultiSet([self.f3, (d_1473_v53_)[default__.safeIndex(len(d_1476_v56_), len(d_1473_v53_))], self.f8])
                d_1478_v58_: D12
                d_1478_v58_ = D12_DC26(d_1476_v56_)
                d_1479_v59_: _dafny.Map
                d_1479_v59_ = _dafny.Map({p1.f3: d_1478_v58_})
                d_1480_v60_: _dafny.Seq
                d_1480_v60_ = _dafny.SeqWithoutIsStrInference([((d_1477_v57_)[(self).f4] if ((self).f4) in (d_1477_v57_) else d_1440_i4_), len(d_1479_v59_)])
                d_1481_v61_: D1
                d_1481_v61_ = D1_DC4(866)
                d_1482_v62_: _dafny.Map
                d_1482_v62_ = _dafny.Map({(d_1481_v61_).cf7: p1})
                d_1483_v63_: _dafny.Seq
                d_1483_v63_ = _dafny.SeqWithoutIsStrInference([((d_1482_v62_)[d_1440_i4_] if (d_1440_i4_) in (d_1482_v62_) else p1), p1])
                d_1484_v64_: _dafny.Map
                d_1484_v64_ = _dafny.Map({p0: len(d_1483_v63_)})
                d_1485_v65_: _dafny.Seq
                d_1485_v65_ = _dafny.SeqWithoutIsStrInference([d_1484_v64_, d_1484_v64_])
                d_1486_v66_: _dafny.Seq
                d_1486_v66_ = _dafny.SeqWithoutIsStrInference([d_1485_v65_, d_1485_v65_])
                d_1487_v67_: _dafny.Array
                nw230_ = _dafny.Array(None, 25)
                nw230_[int(0)] = (364) >= ((self).f9)
                nw230_[int(1)] = (d_1474_v54_).f13
                nw230_[int(2)] = (d_1474_v54_).f13
                nw230_[int(3)] = (p1).f4
                nw230_[int(4)] = (d_1470_v50_).f10
                nw230_[int(5)] = default__.fm18(p0, d_1397_v0_, p0, globalState)
                nw230_[int(6)] = (d_1480_v60_) < (d_1480_v60_)
                nw230_[int(7)] = (self).f4
                nw230_[int(8)] = (d_1397_v0_) not in ((d_1474_v54_).f12)
                nw230_[int(9)] = p1.f3
                nw230_[int(10)] = (p1).f4
                nw230_[int(11)] = default__.fm21((self).f9, _dafny.MultiSet(((d_1486_v66_)[default__.safeIndex(d_1440_i4_, len(d_1486_v66_))]).set(default__.safeIndex(p0, len((d_1486_v66_)[default__.safeIndex(d_1440_i4_, len(d_1486_v66_))])), d_1484_v64_)), globalState)
                nw230_[int(12)] = self.f8
                nw230_[int(13)] = (d_1470_v50_).f10
                nw230_[int(14)] = not((p1).f4)
                nw230_[int(15)] = self.f8
                nw230_[int(16)] = p1.f3
                nw230_[int(17)] = False
                nw230_[int(18)] = (self).f4
                nw230_[int(19)] = (p1).f4
                nw230_[int(20)] = (self.f3) or (p1.f3)
                nw230_[int(21)] = not ((d_1470_v50_).f11) or (self.f8)
                nw230_[int(22)] = default__.fm18(d_1440_i4_, _dafny.CodePoint('d'), len(_dafny.SeqWithoutIsStrInference([(p1).f4, (d_1473_v53_)[default__.safeIndex(len(d_1471_v51_), len(d_1473_v53_))], (d_1470_v50_).f10])), globalState)
                nw230_[int(23)] = not((d_1470_v50_).f11)
                nw230_[int(24)] = not ((p1).f4) or ((d_1474_v54_).f13)
                d_1487_v67_ = nw230_
                d_1487_v67_ = d_1487_v67_
                d_1488_v68_: D2
                d_1488_v68_ = D2_DC8(d_1487_v67_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmomdi")))
                pat_let_tv41_ = d_1487_v67_
                pat_let_tv42_ = d_1487_v67_
                d_1489_v69_: _dafny.Set
                def iife111_(_pat_let31_0):
                    def iife112_(d_1490_dt__update__tmp_h1_):
                        def iife113_(_pat_let32_0):
                            def iife114_(d_1491_dt__update_hcf13_h0_):
                                return D2_DC8(d_1491_dt__update_hcf13_h0_, (d_1490_dt__update__tmp_h1_).cf14)
                            return iife114_(_pat_let32_0)
                        return iife113_(pat_let_tv41_)
                    return iife112_(_pat_let31_0)
                def iife110_(_pat_let30_0):
                    def iife115_(d_1492_dt__update__tmp_h2_):
                        def iife116_(_pat_let33_0):
                            def iife117_(d_1493_dt__update_hcf13_h1_):
                                return D2_DC8(d_1493_dt__update_hcf13_h1_, (d_1492_dt__update__tmp_h2_).cf14)
                            return iife117_(_pat_let33_0)
                        return iife116_(pat_let_tv42_)
                    return iife115_(_pat_let30_0)
                d_1489_v69_ = _dafny.Set({d_1488_v68_, d_1488_v68_, iife110_(iife111_(d_1488_v68_))})
                d_1489_v69_ = d_1489_v69_
            elif True:
                d_1494_v70_: C1
                nw231_ = C1()
                nw231_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipej")), p1.f3)
                d_1494_v70_ = nw231_
                d_1495_v71_: _dafny.Set
                d_1495_v71_ = _dafny.Set({d_1494_v70_, d_1494_v70_, d_1494_v70_})
                (p1).f3 = (d_1495_v71_).issubset(d_1495_v71_)
                d_1496_v72_: _dafny.Set
                d_1496_v72_ = _dafny.Set({False})
                d_1497_v73_: _dafny.Map
                d_1497_v73_ = _dafny.Map({(d_1496_v72_).issubset(_dafny.Set({p1.f3, (p1).f4})): (0) - (default__.safeDivisionInt(len((d_1494_v70_).f12), (self).f9))})
                d_1497_v73_ = (d_1497_v73_).set((d_1494_v70_).f13, len(d_1398_v1_))
                d_1498_v74_: _dafny.Seq
                d_1498_v74_ = _dafny.SeqWithoutIsStrInference([(p1).f4, False, self.f3])
                d_1499_v75_: D4
                d_1499_v75_ = D4_DC11(d_1498_v74_)
                d_1500_v76_: D4
                d_1500_v76_ = D4_DC13(d_1499_v75_)
                d_1501_v77_: D4
                d_1501_v77_ = D4_DC13(d_1499_v75_)
                d_1501_v77_ = d_1501_v77_
                d_1502_v78_: _dafny.Array
                def lambda87_(d_1503_i8_):
                    return default__.safeDivisionInt(d_1503_i8_, 478)

                init46_ = lambda87_
                nw232_ = _dafny.Array(None, 17)
                for i0_46_ in range(nw232_.length(0)):
                    nw232_[i0_46_] = init46_(i0_46_)
                d_1502_v78_ = nw232_
                d_1504_v79_: _dafny.Map
                d_1504_v79_ = _dafny.Map({d_1502_v78_: (-516) * (d_1440_i4_)})
                d_1504_v79_ = (d_1504_v79_).set(d_1502_v78_, p0)
                d_1505_v80_: _dafny.Array
                nw233_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
                d_1505_v80_ = nw233_
                d_1506_v81_: _dafny.Array
                nw234_ = _dafny.Array(False, 13)
                d_1506_v81_ = nw234_
                d_1507_v82_: _dafny.Map
                d_1507_v82_ = _dafny.Map({d_1505_v80_: d_1506_v81_})
                d_1507_v82_ = (d_1507_v82_).set(d_1505_v80_, d_1506_v81_)
            d_1508_v84_: _dafny.Set
            d_1508_v84_ = _dafny.Set({d_1440_i4_})
            d_1509_v85_: _dafny.MultiSet
            def iife118_():
                coll50_ = _dafny.Map()
                compr_50_: int
                for compr_50_ in (d_1508_v84_).Elements:
                    d_1510_v83_: int = compr_50_
                    if (d_1510_v83_) in (d_1508_v84_):
                        coll50_[default__.safeModuloInt(d_1510_v83_, (self).f9)] = (self).f9
                return _dafny.Map(coll50_)
            d_1509_v85_ = _dafny.MultiSet([iife118_()
            ])
            d_1511_v86_: C3
            nw235_ = C3()
            nw235_.ctor__((p1).f4, default__.fm21((0) - (-822), d_1509_v85_, globalState))
            d_1511_v86_ = nw235_
            d_1512_v87_: C1
            nw236_ = C1()
            nw236_.ctor__((d_1398_v1_) + (d_1398_v1_), self.f8)
            d_1512_v87_ = nw236_
        d_1513_v88_: _dafny.Seq
        d_1513_v88_ = _dafny.SeqWithoutIsStrInference([(self).f4, self.f8])
        d_1514_v89_: _dafny.Seq
        d_1514_v89_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_1513_v88_)), (0) - (p0)])
        d_1515_v90_: D4
        d_1515_v90_ = D4_DC12((self).f9, len(d_1514_v89_), self.f3)
        r0 = (d_1515_v90_).cf19
        d_1516_i9_: int
        d_1516_i9_ = 0
        with _dafny.label("7"):
            while (((self).f9) * ((self).f9)) > (len((d_1398_v1_) + (_dafny.SeqWithoutIsStrInference([d_1397_v0_])))):
                with _dafny.c_label("7"):
                    if (d_1516_i9_) >= (100):
                        raise _dafny.Break("7")
                    d_1516_i9_ = (d_1516_i9_) + (1)
                    d_1517_v91_: D7
                    d_1517_v91_ = D7_DC17(p0, p1.f3, False, (self).f9)
                    d_1518_v92_: _dafny.Map
                    d_1518_v92_ = _dafny.Map({p0: 860})
                    d_1519_v93_: _dafny.MultiSet
                    d_1519_v93_ = _dafny.MultiSet([p1.f3, not(p1.f3)])
                    d_1520_v94_: _dafny.Array
                    nw237_ = _dafny.Array(None, 14)
                    nw237_[int(0)] = (self).f9
                    nw237_[int(1)] = p0
                    nw237_[int(2)] = ((self).fm0(globalState)) + ((d_1517_v91_).cf29)
                    nw237_[int(3)] = (p0 if p1.f3 else -207)
                    nw237_[int(4)] = len(d_1518_v92_)
                    nw237_[int(5)] = ((self).f9) + (p0)
                    nw237_[int(6)] = (d_1519_v93_).cardinality
                    nw237_[int(7)] = p0
                    nw237_[int(8)] = ((d_1519_v93_)[p1.f3] if (p1.f3) in (d_1519_v93_) else p0)
                    nw237_[int(9)] = (self).f9
                    nw237_[int(10)] = (self).f9
                    nw237_[int(11)] = default__.safeDivisionInt(p0, p0)
                    nw237_[int(12)] = (0) - (len(_dafny.Map({p1.f3: default__.fm14(globalState)})))
                    nw237_[int(13)] = (self).f9
                    d_1520_v94_ = nw237_
                    index249_ = default__.safeIndex(887, (d_1520_v94_).length(0))
                    (d_1520_v94_)[index249_] = len(default__.fm19(globalState))
                    d_1521_v95_: _dafny.Map
                    d_1521_v95_ = _dafny.Map({(p1).f4: (p1).f4})
                    index250_ = default__.safeIndex(887, (d_1520_v94_).length(0))
                    rhs208_ = True
                    rhs209_ = len((d_1521_v95_) | (_dafny.Map({(p1).f4: (self).f4})))
                    rhs210_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvhykh"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fccwodxf")))
                    rhs211_ = default__.safeDivisionInt(((d_1518_v92_)[len((d_1513_v88_).set(default__.safeIndex(len(d_1513_v88_), len(d_1513_v88_)), default__.fm6(p0, globalState)))] if (len((d_1513_v88_).set(default__.safeIndex(len(d_1513_v88_), len(d_1513_v88_)), default__.fm6(p0, globalState)))) in (d_1518_v92_) else p0), len(d_1398_v1_))
                    rhs212_ = d_1397_v0_
                    lhs150_ = globalState
                    lhs151_ = d_1520_v94_
                    lhs152_ = default__.safeIndex(887, (d_1520_v94_).length(0))
                    lhs150_.f0 = rhs208_
                    r0 = rhs209_
                    d_1398_v1_ = rhs210_
                    lhs151_[lhs152_] = rhs211_
                    d_1397_v0_ = rhs212_
                    d_1522_v96_: _dafny.Set
                    d_1522_v96_ = _dafny.Set({True, self.f8})
                    d_1523_v97_: D12
                    d_1523_v97_ = D12_DC26(d_1522_v96_)
                    d_1524_v98_: D12
                    d_1524_v98_ = D12_DC28(d_1523_v97_)
                    pat_let_tv43_ = d_1523_v97_
                    d_1525_v99_: _dafny.Array
                    nw238_ = _dafny.Array(None, 16)
                    nw238_[int(0)] = (D12_DC28(d_1523_v97_) if self.f8 else d_1524_v98_)
                    nw238_[int(1)] = d_1524_v98_
                    nw238_[int(2)] = d_1524_v98_
                    nw238_[int(3)] = d_1524_v98_
                    nw238_[int(4)] = d_1524_v98_
                    nw238_[int(5)] = d_1524_v98_
                    def iife119_(_pat_let34_0):
                        def iife120_(d_1526_dt__update__tmp_h3_):
                            def iife121_(_pat_let35_0):
                                def iife122_(d_1527_dt__update_hcf41_h0_):
                                    return D12_DC28(d_1527_dt__update_hcf41_h0_)
                                return iife122_(_pat_let35_0)
                            return iife121_(pat_let_tv43_)
                        return iife120_(_pat_let34_0)
                    nw238_[int(6)] = iife119_(d_1524_v98_)
                    nw238_[int(7)] = d_1524_v98_
                    nw238_[int(8)] = d_1524_v98_
                    nw238_[int(9)] = default__.fm38(p0, p1.f3, globalState)
                    nw238_[int(10)] = d_1524_v98_
                    nw238_[int(11)] = d_1524_v98_
                    nw238_[int(12)] = d_1524_v98_
                    nw238_[int(13)] = d_1524_v98_
                    nw238_[int(14)] = default__.fm38((d_1520_v94_)[default__.safeIndex(887, (d_1520_v94_).length(0))], not(p1.f3), globalState)
                    nw238_[int(15)] = default__.fm38((d_1520_v94_)[default__.safeIndex(887, (d_1520_v94_).length(0))], p1.f3, globalState)
                    d_1525_v99_ = nw238_
                    index251_ = default__.safeIndex(751, (d_1525_v99_).length(0))
                    (d_1525_v99_)[index251_] = d_1524_v98_
                    index252_ = default__.safeIndex(751, (d_1525_v99_).length(0))
                    (d_1525_v99_)[index252_] = D12_DC28(default__.fm38(p0, (p1).f4, globalState))
                    d_1398_v1_ = d_1398_v1_
                    d_1521_v95_ = (d_1521_v95_).set((d_1398_v1_) < (d_1398_v1_), True)
                    pass
            pass
        d_1528_v100_: _dafny.Array
        nw239_ = _dafny.Array(_dafny.Set({}), 10)
        d_1528_v100_ = nw239_
        d_1528_v100_ = d_1528_v100_
        r0 = p0
        return r0

    def m2(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1529_v0_: int
        d_1530_v1_: bool
        out28_: int
        out29_: bool
        out28_, out29_ = (self).m5(p1.f3, self.f8, globalState)
        d_1529_v0_ = out28_
        d_1530_v1_ = out29_
        d_1531_v2_: _dafny.Array
        def lambda88_(d_1532_v0_, d_1533_p0_, d_1534_p1_):
            def lambda89_(d_1535_i0_):
                return _dafny.SeqWithoutIsStrInference([d_1532_v0_, len(_dafny.SeqWithoutIsStrInference([d_1533_p0_, d_1533_p0_])), len(_dafny.SeqWithoutIsStrInference([(d_1534_p1_).f4, (d_1534_p1_).f4])), len(_dafny.Set({(self).f9, (_dafny.MultiSet([d_1532_v0_, (self).f9])).cardinality}))])

            return lambda89_

        init47_ = lambda88_(d_1529_v0_, p0, p1)
        nw240_ = _dafny.Array(None, 19)
        for i0_47_ in range(nw240_.length(0)):
            nw240_[i0_47_] = init47_(i0_47_)
        d_1531_v2_ = nw240_
        d_1536_v3_: _dafny.Array
        def lambda90_(d_1537_v0_):
            def lambda91_(d_1538_i1_):
                return _dafny.Set({d_1537_v0_, (self).f9})

            return lambda91_

        init48_ = lambda90_(d_1529_v0_)
        nw241_ = _dafny.Array(None, 24)
        for i0_48_ in range(nw241_.length(0)):
            nw241_[i0_48_] = init48_(i0_48_)
        d_1536_v3_ = nw241_
        d_1539_v4_: D15
        d_1539_v4_ = D15_DC32(d_1536_v3_)
        d_1540_v5_: D15
        d_1540_v5_ = D15_DC34(d_1539_v4_)
        d_1541_v6_: D15
        d_1541_v6_ = D15_DC34(d_1539_v4_)
        d_1542_v7_: D15
        d_1542_v7_ = D15_DC34(d_1541_v6_)
        d_1543_v8_: D15
        d_1543_v8_ = D15_DC34(d_1542_v7_)
        d_1544_v9_: D15
        d_1544_v9_ = D15_DC34(d_1539_v4_)
        d_1545_v10_: D15
        d_1545_v10_ = D15_DC34(d_1543_v8_)
        d_1546_v11_: D15
        d_1546_v11_ = D15_DC34(d_1543_v8_)
        d_1547_v12_: _dafny.Set
        d_1547_v12_ = _dafny.Set({d_1546_v11_, d_1546_v11_})
        d_1548_v13_: _dafny.Seq
        d_1548_v13_ = _dafny.SeqWithoutIsStrInference([len(d_1547_v12_), (self).f9, (0) - (d_1529_v0_), 442, len(_dafny.Map({(self).f9: self.f3}))])
        d_1549_v14_: _dafny.Seq
        d_1549_v14_ = _dafny.SeqWithoutIsStrInference([(p1).f4])
        index253_ = default__.safeIndex(237, (d_1531_v2_).length(0))
        (d_1531_v2_)[index253_] = (d_1548_v13_) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_1549_v14_)).cardinality]))
        index254_ = default__.safeIndex(237, (d_1531_v2_).length(0))
        (d_1531_v2_)[index254_] = d_1548_v13_
        d_1550_v15_: _dafny.Array
        nw242_ = _dafny.Array(False, 1)
        d_1550_v15_ = nw242_
        index255_ = default__.safeIndex(317, (d_1550_v15_).length(0))
        (d_1550_v15_)[index255_] = (self).f4
        d_1551_v16_: _dafny.MultiSet
        d_1551_v16_ = _dafny.MultiSet([not(d_1530_v1_), True])
        index256_ = default__.safeIndex(317, (d_1550_v15_).length(0))
        (d_1550_v15_)[index256_] = (d_1551_v16_).ispropersubset(d_1551_v16_)
        d_1529_v0_ = (self).f9
        if (False if (self).f4 else (True if True else (p1).f4)):
            index257_ = default__.safeIndex(317, (d_1550_v15_).length(0))
            (d_1550_v15_)[index257_] = not(default__.fm14(globalState))
            d_1552_v17_: D12
            d_1552_v17_ = D12_DC27((self).f9)
            d_1553_v18_: D12
            d_1553_v18_ = D12_DC28(d_1552_v17_)
            rhs213_ = (self).f9
            rhs214_ = d_1553_v18_
            d_1529_v0_ = rhs213_
            d_1553_v18_ = rhs214_
            d_1529_v0_ = default__.safeDivisionInt((self).f9, d_1529_v0_)
            d_1554_v19_: _dafny.Seq
            d_1554_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abbgbhwi"))
            d_1555_v20_: _dafny.Map
            d_1555_v20_ = _dafny.Map({d_1529_v0_: (len(d_1554_v19_)) >= ((self).f9)})
            d_1555_v20_ = (d_1555_v20_).set(default__.safeModuloInt(221, (d_1551_v16_).cardinality), self.f3)
            (self).f8 = p1.f3
        elif True:
            d_1556_v21_: _dafny.Map
            d_1556_v21_ = _dafny.Map({(self).f9: self.f3})
            d_1557_v22_: D11
            d_1557_v22_ = D11_DC25(d_1556_v21_)
            d_1557_v22_ = (d_1557_v22_ if self.f3 else d_1557_v22_)
            if ((d_1556_v21_)[(d_1551_v16_).cardinality] if ((d_1551_v16_).cardinality) in (d_1556_v21_) else (d_1550_v15_)[default__.safeIndex(317, (d_1550_v15_).length(0))]):
                d_1558_v23_: _dafny.Array
                nw243_ = _dafny.Array(int(0), 28)
                d_1558_v23_ = nw243_
                rhs215_ = ((d_1548_v13_)[default__.safeIndex(d_1529_v0_, len(d_1548_v13_))]) * (d_1529_v0_)
                rhs216_ = d_1558_v23_
                d_1529_v0_ = rhs215_
                d_1558_v23_ = rhs216_
                d_1559_v24_: D16
                d_1559_v24_ = D16_DC35(_dafny.Map({False: len(d_1556_v21_)}))
                d_1560_v25_: _dafny.Map
                d_1560_v25_ = _dafny.Map({len((d_1559_v24_).cf50): (self).f9})
                d_1561_v26_: _dafny.Seq
                d_1561_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
                d_1562_v27_: _dafny.Seq
                d_1562_v27_ = _dafny.SeqWithoutIsStrInference([d_1561_v26_])
                d_1563_v28_: _dafny.Set
                d_1563_v28_ = _dafny.Set({(self).f9})
                d_1564_v29_: D4
                d_1564_v29_ = D4_DC12(len(d_1560_v25_), len((d_1562_v27_)[default__.safeIndex(len(d_1563_v28_), len(d_1562_v27_))]), (d_1549_v14_)[default__.safeIndex((0) - (d_1529_v0_), len(d_1549_v14_))])
                d_1565_v30_: D7
                d_1565_v30_ = D7_DC17(d_1529_v0_, (self).f4, True, 316)
                d_1566_v31_: D8
                d_1566_v31_ = D8_DC20(p0, (p1).f4)
                d_1567_v32_: _dafny.Map
                d_1567_v32_ = _dafny.Map({p0: d_1566_v31_})
                pat_let_tv44_ = d_1567_v32_
                pat_let_tv45_ = d_1567_v32_
                pat_let_tv46_ = d_1565_v30_
                def iife123_(_pat_let36_0):
                    def iife124_(d_1568_dt__update__tmp_h0_):
                        def iife126_():
                            coll51_ = _dafny.Set()
                            compr_51_: str
                            for compr_51_ in (pat_let_tv44_).keys.Elements:
                                d_1569_v33_: str = compr_51_
                                if (d_1569_v33_) in (pat_let_tv45_):
                                    coll51_ = coll51_.union(_dafny.Set([d_1569_v33_]))
                            return _dafny.Set(coll51_)
                        def iife125_(_pat_let37_0):
                            def iife127_(d_1570_dt__update_hcf19_h0_):
                                def iife128_(_pat_let38_0):
                                    def iife129_(d_1571_dt__update_hcf21_h0_):
                                        return D4_DC12(d_1570_dt__update_hcf19_h0_, (d_1568_dt__update__tmp_h0_).cf20, d_1571_dt__update_hcf21_h0_)
                                    return iife129_(_pat_let38_0)
                                return iife128_((pat_let_tv46_).cf28)
                            return iife127_(_pat_let37_0)
                        return iife125_(len(iife126_()
                        ))
                    return iife124_(_pat_let36_0)
                (p1).f3 = (iife123_(d_1564_v29_)).cf21
                d_1572_v34_: _dafny.MultiSet
                d_1572_v34_ = _dafny.MultiSet([(self).f9])
                d_1573_v35_: C2
                nw244_ = C2()
                nw244_.ctor__(self.f8, ((_dafny.MultiSet([(self).f9])).set(len(d_1561_v26_), default__.abs((self).f9))) == (d_1572_v34_), (d_1549_v14_)[default__.safeIndex((self).f9, len(d_1549_v14_))], d_1530_v1_)
                d_1573_v35_ = nw244_
                d_1530_v1_ = default__.fm14(globalState)
                index258_ = default__.safeIndex(317, (d_1550_v15_).length(0))
                (d_1550_v15_)[index258_] = self.f8
            elif True:
                d_1574_v36_: _dafny.Array
                nw245_ = _dafny.Array(int(0), 22)
                d_1574_v36_ = nw245_
                d_1575_v37_: _dafny.Map
                d_1575_v37_ = _dafny.Map({p0: d_1574_v36_})
                d_1575_v37_ = (d_1575_v37_).set(p0, d_1574_v36_)
                d_1576_v38_: C0
                nw246_ = C0()
                nw246_.ctor__()
                d_1576_v38_ = nw246_
                d_1529_v0_ = default__.safeModuloInt(d_1529_v0_, d_1529_v0_)
                (self).f8 = default__.fm18((d_1529_v0_) * ((self).f9), p0, (self).f9, globalState)
                d_1577_v39_: _dafny.Map
                d_1577_v39_ = _dafny.Map({D12_DC27(d_1529_v0_): (0) - (d_1529_v0_)})
                d_1578_v40_: _dafny.Seq
                d_1578_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcg"))
                d_1577_v39_ = (d_1577_v39_).set(D12_DC27(948), len(d_1578_v40_))
            source18_ = d_1557_v22_
            if source18_.is_DC25:
                d_1579___mcc_h0_ = source18_.cf38
                d_1580_cf38_ = d_1579___mcc_h0_
                d_1581_v41_: _dafny.Map
                d_1581_v41_ = _dafny.Map({-533: default__.fm34((self).f9, (self).f9, globalState)})
                d_1582_v42_: _dafny.Map
                d_1582_v42_ = _dafny.Map({False: p0})
                d_1581_v41_ = (d_1581_v41_).set((len(_dafny.Set({len(d_1582_v42_)}))) + ((self).f9), d_1549_v14_)
                d_1583_v43_: _dafny.Map
                d_1583_v43_ = _dafny.Map({(self).f9: (self).f9})
                d_1584_v44_: _dafny.MultiSet
                d_1584_v44_ = _dafny.MultiSet([d_1583_v43_])
                d_1585_v45_: _dafny.Seq
                d_1585_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jophvjl"))
                d_1586_v46_: _dafny.Map
                d_1586_v46_ = _dafny.Map({not(default__.fm21(402, d_1584_v44_, globalState)): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhflclhlo"))) + (d_1585_v45_)})
                d_1587_v47_: D1
                d_1587_v47_ = D1_DC5((p1).f4, p1.f3, len(d_1548_v13_), (self).f9)
                d_1586_v46_ = (d_1586_v46_).set(((p1).f4) and (p1.f3), default__.fm30(d_1587_v47_, d_1529_v0_, (self).f9, globalState))
                d_1588_v48_: D1
                d_1588_v48_ = D1_DC4((self).f9)
                d_1589_v49_: D3
                d_1589_v49_ = D3_DC10(d_1551_v16_, d_1588_v48_)
                d_1590_v50_: _dafny.Map
                d_1590_v50_ = _dafny.Map({d_1589_v49_: p0})
                d_1590_v50_ = (d_1590_v50_).set(default__.fm39((self).f4, d_1529_v0_, (self).f9, d_1529_v0_, globalState), p0)
                d_1591_v51_: C3
                nw247_ = C3()
                nw247_.ctor__(default__.fm14(globalState), (p1).f4)
                d_1591_v51_ = nw247_
                d_1592_v52_: _dafny.MultiSet
                d_1592_v52_ = _dafny.MultiSet([d_1591_v51_])
                d_1593_v53_: _dafny.MultiSet
                d_1593_v53_ = _dafny.MultiSet([(self).f9, d_1529_v0_])
                rhs217_ = (((d_1592_v52_)[d_1591_v51_] if (d_1591_v51_) in (d_1592_v52_) else len((d_1531_v2_)[default__.safeIndex(237, (d_1531_v2_).length(0))]))) - (default__.safeModuloInt(d_1529_v0_, ((d_1593_v53_)[360] if (360) in (d_1593_v53_) else (self).f9)))
                rhs218_ = _dafny.SeqWithoutIsStrInference([p0 for d_1594_i2_ in range(default__.abs(545))])
                rhs219_ = default__.fm6(((d_1531_v2_)[default__.safeIndex(237, (d_1531_v2_).length(0))])[default__.safeIndex(862, len((d_1531_v2_)[default__.safeIndex(237, (d_1531_v2_).length(0))]))], globalState)
                lhs153_ = p1
                d_1529_v0_ = rhs217_
                r0 = rhs218_
                lhs153_.f3 = rhs219_
            elif True:
                d_1595___mcc_h1_ = source18_.cf37
                d_1596_cf37_ = d_1595___mcc_h1_
                d_1597_v54_: _dafny.Seq
                d_1597_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "grqnosdc"))
                (globalState).f0 = ((d_1597_v54_) + (d_1597_v54_)) < (d_1597_v54_)
                d_1598_v55_: _dafny.Array
                nw248_ = _dafny.Array(_dafny.Map({}), 16)
                d_1598_v55_ = nw248_
                d_1599_v56_: D0
                d_1599_v56_ = D0_DC2(d_1530_v1_, -720, d_1529_v0_)
                d_1598_v55_ = (d_1598_v55_ if (d_1599_v56_).cf3 else d_1598_v55_)
                d_1600_v57_: _dafny.Array
                def lambda92_(d_1601_p1_):
                    def lambda93_(d_1602_i3_):
                        return D7_DC17((self).f9, (d_1601_p1_).f4, self.f3, (self).f9)

                    return lambda93_

                init49_ = lambda92_(p1)
                nw249_ = _dafny.Array(None, 6)
                for i0_49_ in range(nw249_.length(0)):
                    nw249_[i0_49_] = init49_(i0_49_)
                d_1600_v57_ = nw249_
                def lambda94_(d_1603_v13_, d_1604_p1_, d_1605_v1_):
                    def lambda95_(d_1606_i4_):
                        return D7_DC17(len(d_1603_v13_), (d_1604_p1_).f4, d_1605_v1_, (self).f9)

                    return lambda95_

                init50_ = lambda94_(d_1548_v13_, p1, d_1530_v1_)
                nw250_ = _dafny.Array(None, 26)
                for i0_50_ in range(nw250_.length(0)):
                    nw250_[i0_50_] = init50_(i0_50_)
                d_1600_v57_ = nw250_
                (globalState).f0 = (d_1550_v15_)[default__.safeIndex(317, (d_1550_v15_).length(0))]
            d_1607_v58_: _dafny.Set
            d_1607_v58_ = _dafny.Set({(self).f9})
            d_1608_v59_: D7
            d_1608_v59_ = D7_DC17((self).f9, default__.fm18(len(d_1607_v58_), p0, d_1529_v0_, globalState), True, d_1529_v0_)
            d_1609_v60_: _dafny.Seq
            d_1609_v60_ = _dafny.SeqWithoutIsStrInference([d_1608_v59_, d_1608_v59_])
            (self).f8 = ((len(d_1549_v14_)) == (d_1529_v0_) if p1.f3 else ((d_1609_v60_).set(default__.safeIndex(d_1529_v0_, len(d_1609_v60_)), d_1608_v59_)) < (d_1609_v60_))
            d_1529_v0_ = (self).f9
        d_1610_v61_: _dafny.Map
        d_1610_v61_ = _dafny.Map({not(not(p1.f3)): d_1529_v0_})
        d_1611_v62_: D16
        d_1611_v62_ = D16_DC35((d_1610_v61_) | (d_1610_v61_))
        source19_ = d_1611_v62_
        if source19_.is_DC36:
            d_1612___mcc_h2_ = source19_.cf51
            d_1613___mcc_h3_ = source19_.cf52
            d_1614___mcc_h4_ = source19_.cf53
            d_1615___mcc_h5_ = source19_.cf54
            d_1616_cf54_ = d_1615___mcc_h5_
            d_1617_cf53_ = d_1614___mcc_h4_
            d_1618_cf52_ = d_1613___mcc_h3_
            d_1619_cf51_ = d_1612___mcc_h2_
            d_1620_v63_: _dafny.Seq
            d_1620_v63_ = _dafny.SeqWithoutIsStrInference([d_1550_v15_])
            d_1621_v64_: _dafny.MultiSet
            d_1621_v64_ = _dafny.MultiSet([(0) - (d_1529_v0_)])
            d_1622_v65_: _dafny.Map
            d_1622_v65_ = _dafny.Map({(0) - (d_1529_v0_): (p1).f4})
            d_1623_v66_: _dafny.Map
            d_1623_v66_ = _dafny.Map({d_1529_v0_: -818})
            d_1624_v67_: _dafny.Seq
            d_1624_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnfqnbnws"))
            d_1625_v69_: D1
            d_1625_v69_ = D1_DC5(d_1530_v1_, (p1).f4, (self).f9, len(d_1624_v67_))
            d_1626_v70_: _dafny.Array
            nw251_ = _dafny.Array(None, 28)
            nw251_[int(0)] = (self).f9
            nw251_[int(1)] = len(((d_1620_v63_) + (d_1620_v63_)).set(default__.safeIndex((d_1621_v64_).cardinality, len((d_1620_v63_) + (d_1620_v63_))), d_1550_v15_))
            nw251_[int(2)] = 101
            nw251_[int(3)] = (len(d_1622_v65_)) + (len(d_1623_v66_))
            nw251_[int(4)] = d_1529_v0_
            nw251_[int(5)] = (d_1529_v0_) + (d_1529_v0_)
            nw251_[int(6)] = (0) - ((d_1529_v0_) + ((self).f9))
            nw251_[int(7)] = d_1529_v0_
            nw251_[int(8)] = d_1529_v0_
            nw251_[int(9)] = (self).f9
            nw251_[int(10)] = d_1529_v0_
            nw251_[int(11)] = (0) - ((self).f9)
            nw251_[int(12)] = d_1529_v0_
            nw251_[int(13)] = len(_dafny.Map({d_1624_v67_: (self).f4}))
            nw251_[int(14)] = len((d_1624_v67_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "swjrixy"))))
            nw251_[int(15)] = (0) - ((self).f9)
            nw251_[int(16)] = default__.safeModuloInt(((d_1551_v16_)[self.f3] if (self.f3) in (d_1551_v16_) else d_1529_v0_), (self).f9)
            nw251_[int(17)] = (self).f9
            nw251_[int(18)] = (self).f9
            def iife130_():
                coll52_ = _dafny.Set()
                compr_52_: int
                for compr_52_ in _dafny.IntegerRange(283, 923):
                    d_1627_v68_: int = compr_52_
                    if ((283) <= (d_1627_v68_)) and ((d_1627_v68_) < (923)):
                        coll52_ = coll52_.union(_dafny.Set([default__.safeDivisionInt(d_1627_v68_, len(d_1623_v66_))]))
                return _dafny.Set(coll52_)
            nw251_[int(19)] = len(iife130_()
            )
            nw251_[int(20)] = ((self).f9) * ((d_1625_v69_).cf10)
            nw251_[int(21)] = ((self).f9) - ((self).f9)
            nw251_[int(22)] = (self).f9
            nw251_[int(23)] = d_1529_v0_
            nw251_[int(24)] = (self).f9
            nw251_[int(25)] = d_1529_v0_
            nw251_[int(26)] = default__.safeDivisionInt((d_1548_v13_)[default__.safeIndex(d_1529_v0_, len(d_1548_v13_))], (self).fm5((self).f4, d_1624_v67_, globalState))
            nw251_[int(27)] = d_1529_v0_
            d_1626_v70_ = nw251_
            index259_ = default__.safeIndex(815, (d_1626_v70_).length(0))
            (d_1626_v70_)[index259_] = d_1529_v0_
            d_1628_v71_: _dafny.Map
            d_1628_v71_ = _dafny.Map({d_1617_cf53_: (p1).f4})
            d_1629_v72_: _dafny.Set
            d_1629_v72_ = _dafny.Set({d_1628_v71_, d_1628_v71_})
            d_1630_v73_: _dafny.MultiSet
            d_1630_v73_ = _dafny.MultiSet([d_1628_v71_])
            index260_ = default__.safeIndex(815, (d_1626_v70_).length(0))
            def iife131_():
                coll53_ = _dafny.Set()
                compr_53_: _dafny.Map
                for compr_53_ in ((d_1630_v73_).set(d_1628_v71_, default__.abs(d_1529_v0_))).Elements:
                    d_1631_v74_: _dafny.Map = compr_53_
                    if (d_1631_v74_) in ((d_1630_v73_).set(d_1628_v71_, default__.abs(d_1529_v0_))):
                        coll53_ = coll53_.union(_dafny.Set([d_1631_v74_]))
                return _dafny.Set(coll53_)
            (d_1626_v70_)[index260_] = len((d_1629_v72_) | ((d_1629_v72_).intersection(iife131_()
            )))
            if d_1530_v1_:
                index261_ = default__.safeIndex(815, (d_1626_v70_).length(0))
                (d_1626_v70_)[index261_] = (self).f9
                d_1632_v75_: T1
                nw252_ = C1()
                nw252_.ctor__(d_1624_v67_, ((self).f9) > (83))
                d_1632_v75_ = nw252_
                d_1632_v75_ = d_1632_v75_
                d_1529_v0_ = ((self).f9) + (((self).f9) - (265))
                d_1529_v0_ = ((d_1626_v70_)[default__.safeIndex(815, (d_1626_v70_).length(0))]) - ((d_1626_v70_)[default__.safeIndex(815, (d_1626_v70_).length(0))])
                d_1633_v76_: _dafny.Map
                d_1633_v76_ = _dafny.Map({(d_1531_v2_)[default__.safeIndex(237, (d_1531_v2_).length(0))]: p0})
                d_1634_v77_: D17
                d_1634_v77_ = D17_DC37(d_1633_v76_)
                d_1635_v79_: _dafny.Set
                d_1635_v79_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_1529_v0_])})
                d_1636_v80_: D18
                d_1636_v80_ = D18_DC41(d_1635_v79_)
                d_1637_v81_: _dafny.Seq
                d_1637_v81_ = _dafny.SeqWithoutIsStrInference([d_1633_v76_, d_1633_v76_, d_1633_v76_])
                d_1638_v83_: _dafny.MultiSet
                d_1638_v83_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1529_v0_ for d_1639_i6_ in range(default__.abs(964))])])
                d_1640_v85_: _dafny.Map
                d_1640_v85_ = _dafny.Map({(d_1531_v2_)[default__.safeIndex(237, (d_1531_v2_).length(0))]: (p1).f4})
                d_1641_v87_: _dafny.Array
                nw253_ = _dafny.Array(None, 27)
                nw253_[int(0)] = (d_1633_v76_).set(default__.fm19(globalState), p0)
                nw253_[int(1)] = (d_1634_v77_).cf55
                nw253_[int(2)] = _dafny.Map({_dafny.SeqWithoutIsStrInference([-261 for d_1642_i5_ in range(default__.abs(639))]): p0})
                nw253_[int(3)] = d_1633_v76_
                nw253_[int(4)] = d_1633_v76_
                nw253_[int(5)] = d_1633_v76_
                def iife132_():
                    coll54_ = _dafny.Map()
                    compr_54_: _dafny.Seq
                    for compr_54_ in ((d_1636_v80_).cf63).Elements:
                        d_1643_v78_: _dafny.Seq = compr_54_
                        if (d_1643_v78_) in ((d_1636_v80_).cf63):
                            coll54_[d_1643_v78_] = p0
                    return _dafny.Map(coll54_)
                nw253_[int(6)] = iife132_()
                
                nw253_[int(7)] = (d_1637_v81_)[default__.safeIndex((d_1626_v70_)[default__.safeIndex(815, (d_1626_v70_).length(0))], len(d_1637_v81_))]
                nw253_[int(8)] = (d_1633_v76_).set((d_1531_v2_)[default__.safeIndex(237, (d_1531_v2_).length(0))], p0)
                nw253_[int(9)] = (_dafny.Map({d_1548_v13_: p0})) | (d_1633_v76_)
                nw253_[int(10)] = d_1633_v76_
                nw253_[int(11)] = _dafny.Map({(d_1531_v2_)[default__.safeIndex(237, (d_1531_v2_).length(0))]: p0})
                nw253_[int(12)] = _dafny.Map({_dafny.SeqWithoutIsStrInference([len(d_1624_v67_)]): p0})
                nw253_[int(13)] = d_1633_v76_
                nw253_[int(14)] = (d_1633_v76_) | (d_1633_v76_)
                nw253_[int(15)] = _dafny.Map({d_1548_v13_: p0})
                nw253_[int(16)] = d_1633_v76_
                def iife133_():
                    coll55_ = _dafny.Map()
                    compr_55_: _dafny.Seq
                    for compr_55_ in (d_1638_v83_).Elements:
                        d_1644_v82_: _dafny.Seq = compr_55_
                        if (d_1644_v82_) in (d_1638_v83_):
                            coll55_[d_1644_v82_] = p0
                    return _dafny.Map(coll55_)
                nw253_[int(17)] = iife133_()
                
                nw253_[int(18)] = d_1633_v76_
                nw253_[int(19)] = d_1633_v76_
                nw253_[int(20)] = d_1633_v76_
                nw253_[int(21)] = d_1633_v76_
                nw253_[int(22)] = d_1633_v76_
                nw253_[int(23)] = (_dafny.Map({d_1548_v13_: p0})) | (_dafny.Map({(d_1531_v2_)[default__.safeIndex(237, (d_1531_v2_).length(0))]: p0}))
                def iife134_():
                    coll56_ = _dafny.Map()
                    compr_56_: _dafny.Seq
                    for compr_56_ in (d_1640_v85_).keys.Elements:
                        d_1645_v84_: _dafny.Seq = compr_56_
                        if (d_1645_v84_) in (d_1640_v85_):
                            coll56_[d_1645_v84_] = _dafny.CodePoint('p')
                    return _dafny.Map(coll56_)
                def iife135_():
                    coll57_ = _dafny.Map()
                    compr_57_: _dafny.Seq
                    for compr_57_ in (d_1638_v83_).Elements:
                        d_1646_v86_: _dafny.Seq = compr_57_
                        if (d_1646_v86_) in (d_1638_v83_):
                            coll57_[d_1646_v86_] = p0
                    return _dafny.Map(coll57_)
                nw253_[int(24)] = (iife134_()
                ) | (iife135_()
                )
                nw253_[int(25)] = d_1633_v76_
                nw253_[int(26)] = ((d_1633_v76_).set(d_1548_v13_, p0)) | (d_1633_v76_)
                d_1641_v87_ = nw253_
                index262_ = default__.safeIndex(760, (d_1641_v87_).length(0))
                (d_1641_v87_)[index262_] = d_1633_v76_
                index263_ = default__.safeIndex(760, (d_1641_v87_).length(0))
                (d_1641_v87_)[index263_] = (d_1633_v76_) | (d_1633_v76_)
            elif True:
                d_1647_v89_: _dafny.Seq
                def iife136_():
                    coll58_ = _dafny.Map()
                    compr_58_: int
                    for compr_58_ in (_dafny.SeqWithoutIsStrInference([(d_1626_v70_)[default__.safeIndex(815, (d_1626_v70_).length(0))]])).Elements:
                        d_1648_v88_: int = compr_58_
                        if (d_1648_v88_) in (_dafny.SeqWithoutIsStrInference([(d_1626_v70_)[default__.safeIndex(815, (d_1626_v70_).length(0))]])):
                            coll58_[default__.safeDivisionInt(d_1648_v88_, (d_1626_v70_)[default__.safeIndex(815, (d_1626_v70_).length(0))])] = not(True)
                    return _dafny.Map(coll58_)
                d_1647_v89_ = _dafny.SeqWithoutIsStrInference([d_1622_v65_, iife136_()
                , (d_1622_v65_ if (p1).f4 else d_1622_v65_), (d_1622_v65_) | (d_1622_v65_), _dafny.Map({(d_1626_v70_)[default__.safeIndex(815, (d_1626_v70_).length(0))]: d_1616_cf54_})])
                def iife137_():
                    coll59_ = _dafny.Set()
                    compr_59_: int
                    for compr_59_ in _dafny.IntegerRange(626, 615):
                        d_1649_v90_: int = compr_59_
                        if ((626) <= (d_1649_v90_)) and ((d_1649_v90_) < (615)):
                            coll59_ = coll59_.union(_dafny.Set([default__.safeDivisionInt(d_1649_v90_, len(_dafny.SeqWithoutIsStrInference([self.f3, not(self.f3), self.f3])))]))
                    return _dafny.Set(coll59_)
                rhs220_ = d_1647_v89_
                rhs221_ = (iife137_()
                ).issubset(_dafny.Set({len(_dafny.Map({(d_1626_v70_)[default__.safeIndex(815, (d_1626_v70_).length(0))]: d_1616_cf54_}))}))
                rhs222_ = (((p1).f4 if (p1).f4 else p1.f3)) or ((d_1618_cf52_) == (not(True)))
                rhs223_ = (self).f9
                lhs154_ = globalState
                d_1647_v89_ = rhs220_
                lhs154_.f0 = rhs221_
                d_1530_v1_ = rhs222_
                d_1529_v0_ = rhs223_
                d_1650_v92_: D8
                def iife138_():
                    coll60_ = _dafny.Map()
                    compr_60_: int
                    for compr_60_ in (d_1621_v64_).Elements:
                        d_1651_v91_: int = compr_60_
                        if (d_1651_v91_) in (d_1621_v64_):
                            coll60_[(d_1651_v91_) + ((self).f9)] = d_1616_cf54_
                    return _dafny.Map(coll60_)
                d_1650_v92_ = D8_DC18(iife138_()
)
                pat_let_tv47_ = d_1622_v65_
                d_1652_v93_: _dafny.Map
                def iife139_(_pat_let39_0):
                    def iife140_(d_1653_dt__update__tmp_h1_):
                        def iife141_(_pat_let40_0):
                            def iife142_(d_1654_dt__update_hcf30_h0_):
                                return D8_DC18(d_1654_dt__update_hcf30_h0_)
                            return iife142_(_pat_let40_0)
                        return iife141_(pat_let_tv47_)
                    return iife140_(_pat_let39_0)
                d_1652_v93_ = _dafny.Map({d_1529_v0_: iife139_(d_1650_v92_)})
                d_1655_v94_: _dafny.MultiSet
                d_1655_v94_ = _dafny.MultiSet([d_1652_v93_])
                d_1655_v94_ = (d_1655_v94_ if (p1).f4 else d_1655_v94_)
                d_1656_v95_: D4
                d_1656_v95_ = D4_DC12((d_1626_v70_)[default__.safeIndex(815, (d_1626_v70_).length(0))], (self).f9, self.f8)
                (globalState).f0 = not ((d_1550_v15_)[default__.safeIndex(317, (d_1550_v15_).length(0))]) or ((d_1656_v95_).cf21)
                index264_ = default__.safeIndex(815, (d_1626_v70_).length(0))
                (d_1626_v70_)[index264_] = d_1529_v0_
                index265_ = default__.safeIndex(815, (d_1626_v70_).length(0))
                (d_1626_v70_)[index265_] = (0) - ((self).fm0(globalState))
            (p1).f3 = not(d_1530_v1_)
            d_1657_v96_: D0
            d_1657_v96_ = D0_DC1(p1.f3, p0)
            d_1658_v97_: _dafny.Seq
            d_1658_v97_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1659_i7_ in range(default__.abs(-829))]), _dafny.SeqWithoutIsStrInference([p0 for d_1660_i8_ in range(default__.abs(249))])])])
            d_1661_v98_: _dafny.Seq
            d_1661_v98_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykv"))])
            d_1662_v99_: _dafny.Seq
            d_1662_v99_ = d_1661_v98_
            index266_ = default__.safeIndex(317, (d_1550_v15_).length(0))
            rhs224_ = d_1657_v96_
            rhs225_ = (((self).f9) + ((d_1626_v70_)[default__.safeIndex(815, (d_1626_v70_).length(0))])) >= (default__.safeModuloInt(len((d_1658_v97_).set(default__.safeIndex(d_1529_v0_, len(d_1658_v97_)), d_1662_v99_)), d_1529_v0_))
            lhs155_ = d_1550_v15_
            lhs156_ = default__.safeIndex(317, (d_1550_v15_).length(0))
            d_1657_v96_ = rhs224_
            lhs155_[lhs156_] = rhs225_
        elif True:
            d_1663___mcc_h6_ = source19_.cf50
            d_1664_cf50_ = d_1663___mcc_h6_
            if self.f8:
                d_1665_v100_: int
                d_1666_v101_: bool
                out30_: int
                out31_: bool
                out30_, out31_ = (self).m5((p1).f4, p1.f3, globalState)
                d_1665_v100_ = out30_
                d_1666_v101_ = out31_
                d_1667_v102_: _dafny.Set
                d_1667_v102_ = _dafny.Set({len(d_1664_cf50_), len(_dafny.Map({(p1).f4: p0})), len(_dafny.SeqWithoutIsStrInference([p0 for d_1668_i9_ in range(default__.abs(-156))]))})
                d_1669_v103_: _dafny.Array
                def lambda96_(d_1670_v100_):
                    def lambda97_(d_1671_i10_):
                        return (d_1671_i10_) - (d_1670_v100_)

                    return lambda97_

                init51_ = lambda96_(d_1665_v100_)
                nw254_ = _dafny.Array(None, 26)
                for i0_51_ in range(nw254_.length(0)):
                    nw254_[i0_51_] = init51_(i0_51_)
                d_1669_v103_ = nw254_
                d_1672_v104_: _dafny.Seq
                d_1672_v104_ = _dafny.SeqWithoutIsStrInference([d_1669_v103_])
                d_1673_v105_: _dafny.Array
                nw255_ = _dafny.Array(None, 21)
                nw255_[int(0)] = (d_1672_v104_)[default__.safeIndex(856, len(d_1672_v104_))]
                nw255_[int(1)] = d_1669_v103_
                nw255_[int(2)] = d_1669_v103_
                nw255_[int(3)] = d_1669_v103_
                nw255_[int(4)] = d_1669_v103_
                nw255_[int(5)] = d_1669_v103_
                nw255_[int(6)] = d_1669_v103_
                nw255_[int(7)] = d_1669_v103_
                nw255_[int(8)] = d_1669_v103_
                nw255_[int(9)] = d_1669_v103_
                nw255_[int(10)] = d_1669_v103_
                nw255_[int(11)] = d_1669_v103_
                nw255_[int(12)] = d_1669_v103_
                nw255_[int(13)] = d_1669_v103_
                nw255_[int(14)] = d_1669_v103_
                nw255_[int(15)] = d_1669_v103_
                nw255_[int(16)] = d_1669_v103_
                nw255_[int(17)] = d_1669_v103_
                nw255_[int(18)] = d_1669_v103_
                nw255_[int(19)] = d_1669_v103_
                nw255_[int(20)] = d_1669_v103_
                d_1673_v105_ = nw255_
                d_1674_v106_: _dafny.Map
                d_1674_v106_ = _dafny.Map({d_1673_v105_: p1.f3})
                d_1675_v107_: D8
                d_1675_v107_ = D8_DC19(d_1529_v0_, (self).f9)
                d_1676_v108_: _dafny.Map
                d_1676_v108_ = _dafny.Map({(self).f9: (self).f9})
                d_1677_v109_: C3
                nw256_ = C3()
                nw256_.ctor__((d_1550_v15_)[default__.safeIndex(317, (d_1550_v15_).length(0))], True)
                d_1677_v109_ = nw256_
                d_1678_v110_: D0
                d_1678_v110_ = D0_DC2(p1.f3, (self).f9, d_1665_v100_)
                d_1679_v111_: _dafny.Array
                nw257_ = _dafny.Array(None, 13)
                nw257_[int(0)] = default__.fm31((self).f4, 170, not(False), _dafny.SeqWithoutIsStrInference([default__.fm18(len(d_1667_v102_), _dafny.CodePoint('e'), d_1529_v0_, globalState), ((d_1674_v106_)[d_1673_v105_] if (d_1673_v105_) in (d_1674_v106_) else d_1530_v1_)]), globalState)
                nw257_[int(1)] = default__.safeDivisionInt((self).f9, d_1529_v0_)
                nw257_[int(2)] = d_1665_v100_
                nw257_[int(3)] = d_1665_v100_
                nw257_[int(4)] = (0) - ((0) - ((0) - ((d_1529_v0_ if p1.f3 else (d_1675_v107_).cf31))))
                nw257_[int(5)] = (self).f9
                nw257_[int(6)] = (self).f9
                nw257_[int(7)] = (self).f9
                nw257_[int(8)] = 112
                nw257_[int(9)] = len((d_1676_v108_) | (d_1676_v108_))
                nw257_[int(10)] = (-166) * (len(_dafny.SeqWithoutIsStrInference([d_1677_v109_])))
                nw257_[int(11)] = (d_1529_v0_) - (d_1529_v0_)
                nw257_[int(12)] = (d_1678_v110_).cf4
                d_1679_v111_ = nw257_
                d_1679_v111_ = d_1679_v111_
                d_1680_v112_: _dafny.Map
                d_1680_v112_ = _dafny.Map({d_1548_v13_: d_1529_v0_})
                d_1681_v114_: _dafny.Seq
                d_1681_v114_ = _dafny.SeqWithoutIsStrInference([d_1548_v13_])
                def iife143_():
                    def iife145_():
                        coll63_ = _dafny.Set()
                        compr_63_: _dafny.Seq
                        for compr_63_ in (d_1681_v114_).Elements:
                            d_1684_v115_: _dafny.Seq = compr_63_
                            if (d_1684_v115_) in (d_1681_v114_):
                                coll63_ = coll63_.union(_dafny.Set([d_1684_v115_]))
                        return _dafny.Set(coll63_)
                    coll61_ = _dafny.Map()
                    def iife144_():
                        coll62_ = _dafny.Set()
                        compr_62_: _dafny.Seq
                        for compr_62_ in (d_1681_v114_).Elements:
                            d_1682_v115_: _dafny.Seq = compr_62_
                            if (d_1682_v115_) in (d_1681_v114_):
                                coll62_ = coll62_.union(_dafny.Set([d_1682_v115_]))
                        return _dafny.Set(coll62_)
                    compr_61_: _dafny.Seq
                    for compr_61_ in (iife144_()
                    ).Elements:
                        d_1683_v113_: _dafny.Seq = compr_61_
                        if (d_1683_v113_) in (iife145_()
                        ):
                            coll61_[d_1683_v113_] = (self).f9
                    return _dafny.Map(coll61_)
                d_1680_v112_ = (d_1680_v112_) | (iife143_()
                )
                d_1685_v116_: T1
                nw258_ = C3()
                nw258_.ctor__(d_1530_v1_, (d_1529_v0_) != (d_1665_v100_))
                d_1685_v116_ = nw258_
                d_1685_v116_ = d_1685_v116_
                d_1676_v108_ = (d_1676_v108_).set(d_1665_v100_, 420)
            elif True:
                d_1686_v117_: _dafny.Array
                nw259_ = _dafny.Array(int(0), 2)
                d_1686_v117_ = nw259_
                index267_ = default__.safeIndex(658, (d_1686_v117_).length(0))
                (d_1686_v117_)[index267_] = default__.safeModuloInt(d_1529_v0_, -683)
                d_1687_v118_: D18
                d_1687_v118_ = D18_DC43(d_1529_v0_, (self).f9, (p1).f4)
                index268_ = default__.safeIndex(658, (d_1686_v117_).length(0))
                (d_1686_v117_)[index268_] = (0) - ((d_1687_v118_).cf65)
                d_1688_v119_: _dafny.Seq
                d_1688_v119_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhsbcmbd"))
                d_1689_v120_: _dafny.Seq
                d_1689_v120_ = _dafny.SeqWithoutIsStrInference([d_1688_v119_, d_1688_v119_])
                d_1690_v121_: _dafny.Seq
                d_1690_v121_ = _dafny.SeqWithoutIsStrInference([(d_1689_v120_)[default__.safeIndex((d_1686_v117_)[default__.safeIndex(658, (d_1686_v117_).length(0))], len(d_1689_v120_))], (d_1688_v119_) + (d_1688_v119_), d_1688_v119_])
                d_1691_v122_: D7
                d_1691_v122_ = D7_DC17(((d_1664_cf50_)[(d_1549_v14_)[default__.safeIndex((d_1686_v117_)[default__.safeIndex(658, (d_1686_v117_).length(0))], len(d_1549_v14_))]] if ((d_1549_v14_)[default__.safeIndex((d_1686_v117_)[default__.safeIndex(658, (d_1686_v117_).length(0))], len(d_1549_v14_))]) in (d_1664_cf50_) else (d_1686_v117_)[default__.safeIndex(658, (d_1686_v117_).length(0))]), (d_1550_v15_)[default__.safeIndex(317, (d_1550_v15_).length(0))], self.f3, len((d_1531_v2_)[default__.safeIndex(237, (d_1531_v2_).length(0))]))
                r0 = (d_1690_v121_)[default__.safeIndex(default__.safeModuloInt((d_1691_v122_).cf26, d_1529_v0_), len(d_1690_v121_))]
                d_1692_v123_: _dafny.MultiSet
                d_1692_v123_ = _dafny.MultiSet([(self).f9, (d_1686_v117_)[default__.safeIndex(658, (d_1686_v117_).length(0))], d_1529_v0_])
                d_1693_v124_: C2
                nw260_ = C2()
                nw260_.ctor__((D8_DC20(p0, (p1).f4)).cf34, (self).f4, (d_1529_v0_) == ((d_1692_v123_).cardinality), p1.f3)
                d_1693_v124_ = nw260_
                r0 = (d_1688_v119_) + (_dafny.SeqWithoutIsStrInference([p0 for d_1694_i11_ in range(default__.abs(92))]))
                index269_ = default__.safeIndex(658, (d_1686_v117_).length(0))
                (d_1686_v117_)[index269_] = default__.safeModuloInt(915, (d_1686_v117_)[default__.safeIndex(658, (d_1686_v117_).length(0))])
            d_1695_v125_: _dafny.Map
            d_1695_v125_ = _dafny.Map({self.f8: p0})
            d_1696_v126_: D14
            d_1696_v126_ = D14_DC31(d_1529_v0_, len((((d_1695_v125_).set((self).f4, p0)).set(self.f8, p0)) | (default__.fm40(p0, self.f3, d_1529_v0_, d_1529_v0_, globalState))))
            source20_ = d_1696_v126_
            if source20_.is_DC31:
                d_1697___mcc_h7_ = source20_.cf44
                d_1698___mcc_h8_ = source20_.cf45
                d_1699_cf45_ = d_1698___mcc_h8_
                d_1700_cf44_ = d_1697___mcc_h7_
                d_1701_v127_: C3
                nw261_ = C3()
                nw261_.ctor__(((d_1550_v15_)[default__.safeIndex(317, (d_1550_v15_).length(0))]) == ((p1).f4), (p1).f4)
                d_1701_v127_ = nw261_
                d_1529_v0_ = d_1700_cf44_
                d_1702_v128_: _dafny.Seq
                d_1702_v128_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxfvhd"))
                d_1703_v129_: _dafny.Map
                d_1703_v129_ = _dafny.Map({d_1702_v128_: (d_1550_v15_)[default__.safeIndex(317, (d_1550_v15_).length(0))]})
                d_1704_v130_: D1
                d_1704_v130_ = D1_DC5(p1.f3, (p1).f4, d_1529_v0_, len(_dafny.Set({d_1699_cf45_})))
                d_1705_v131_: _dafny.Set
                d_1705_v131_ = _dafny.Set({d_1529_v0_, d_1700_cf44_, d_1699_cf45_})
                d_1706_v132_: _dafny.Map
                d_1706_v132_ = _dafny.Map({d_1699_cf45_: (d_1550_v15_)[default__.safeIndex(317, (d_1550_v15_).length(0))]})
                pat_let_tv48_ = d_1549_v14_
                pat_let_tv49_ = p1
                pat_let_tv50_ = globalState
                pat_let_tv51_ = d_1549_v14_
                pat_let_tv52_ = d_1702_v128_
                pat_let_tv53_ = globalState
                def iife146_(_pat_let41_0):
                    def iife147_(d_1707_dt__update__tmp_h2_):
                        def iife148_(_pat_let42_0):
                            def iife149_(d_1708_dt__update_hcf11_h0_):
                                return D1_DC5((d_1707_dt__update__tmp_h2_).cf8, (d_1707_dt__update__tmp_h2_).cf9, (d_1707_dt__update__tmp_h2_).cf10, d_1708_dt__update_hcf11_h0_)
                            return iife149_(_pat_let42_0)
                        return iife148_((0) - ((self).fm5((pat_let_tv48_)[default__.safeIndex((pat_let_tv49_).fm0(pat_let_tv50_), len(pat_let_tv51_))], pat_let_tv52_, pat_let_tv53_)))
                    return iife147_(_pat_let41_0)
                d_1703_v129_ = (d_1703_v129_).set(default__.fm30(iife146_(d_1704_v130_), len(d_1705_v131_), len(d_1706_v132_), globalState), (d_1550_v15_)[default__.safeIndex(317, (d_1550_v15_).length(0))])
                d_1709_v133_: _dafny.Array
                def lambda98_(d_1710_p0_):
                    def lambda99_(d_1711_i12_):
                        return d_1710_p0_

                    return lambda99_

                init52_ = lambda98_(p0)
                nw262_ = _dafny.Array(None, 12)
                for i0_52_ in range(nw262_.length(0)):
                    nw262_[i0_52_] = init52_(i0_52_)
                d_1709_v133_ = nw262_
                index270_ = default__.safeIndex(902, (d_1709_v133_).length(0))
                (d_1709_v133_)[index270_] = (d_1702_v128_)[default__.safeIndex(len((d_1531_v2_)[default__.safeIndex(237, (d_1531_v2_).length(0))]), len(d_1702_v128_))]
                d_1712_v134_: str
                d_1712_v134_ = _dafny.CodePoint('x')
                index271_ = default__.safeIndex(317, (d_1550_v15_).length(0))
                index272_ = default__.safeIndex(902, (d_1709_v133_).length(0))
                rhs226_ = not(not(p1.f3))
                rhs227_ = (D0_DC1((p1).f4, p0)).cf2
                rhs228_ = default__.fm8((0) - (d_1529_v0_), globalState)
                lhs157_ = d_1550_v15_
                lhs158_ = default__.safeIndex(317, (d_1550_v15_).length(0))
                lhs159_ = d_1709_v133_
                lhs160_ = default__.safeIndex(902, (d_1709_v133_).length(0))
                lhs157_[lhs158_] = rhs226_
                lhs159_[lhs160_] = rhs227_
                d_1712_v134_ = rhs228_
            elif True:
                d_1713___mcc_h9_ = source20_.cf43
                d_1714_cf43_ = d_1713___mcc_h9_
                d_1529_v0_ = d_1529_v0_
                d_1715_v135_: _dafny.Map
                d_1715_v135_ = _dafny.Map({-101: d_1551_v16_})
                d_1551_v16_ = (((d_1715_v135_)[(self).fm0(globalState)] if ((self).fm0(globalState)) in (d_1715_v135_) else d_1551_v16_)).set(p1.f3, default__.abs((self).f9))
                index273_ = default__.safeIndex(317, (d_1550_v15_).length(0))
                (d_1550_v15_)[index273_] = self.f8
                d_1716_v136_: C0
                nw263_ = C0()
                nw263_.ctor__()
                d_1716_v136_ = nw263_
            d_1717_v137_: _dafny.Map
            d_1717_v137_ = _dafny.Map({p1.f3: d_1530_v1_})
            d_1717_v137_ = (d_1717_v137_).set(d_1530_v1_, (d_1550_v15_)[default__.safeIndex(317, (d_1550_v15_).length(0))])
            d_1529_v0_ = d_1529_v0_
        r0 = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1718_i13_ in range(default__.abs(-423))])
        return r0

    def m0(self, p0, p1, p2, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        r1 = p2
        d_1719_v0_: _dafny.Array
        nw264_ = _dafny.Array(None, 8)
        nw264_[int(0)] = default__.fm6(p2, globalState)
        nw264_[int(1)] = not((self).f4)
        nw264_[int(2)] = self.f3
        nw264_[int(3)] = p0
        nw264_[int(4)] = (self.f3) == (self.f3)
        nw264_[int(5)] = p0
        nw264_[int(6)] = (self).f4
        nw264_[int(7)] = (self).f4
        d_1719_v0_ = nw264_
        d_1719_v0_ = d_1719_v0_
        d_1720_v1_: _dafny.Set
        d_1720_v1_ = _dafny.Set({self.f3, p0, default__.fm6(p2, globalState)})
        d_1721_v2_: _dafny.MultiSet
        d_1721_v2_ = _dafny.MultiSet([(d_1720_v1_).intersection(d_1720_v1_), d_1720_v1_])
        index274_ = default__.safeIndex(738, (d_1719_v0_).length(0))
        (d_1719_v0_)[index274_] = p0
        d_1722_v3_: _dafny.Seq
        d_1722_v3_ = _dafny.SeqWithoutIsStrInference([d_1720_v1_])
        d_1723_v4_: D12
        d_1723_v4_ = D12_DC26(d_1720_v1_)
        d_1724_v5_: str
        d_1724_v5_ = _dafny.CodePoint('k')
        d_1725_v6_: _dafny.Seq
        d_1725_v6_ = _dafny.SeqWithoutIsStrInference([d_1724_v5_, d_1724_v5_])
        d_1726_v7_: _dafny.Map
        d_1726_v7_ = _dafny.Map({len(d_1725_v6_): d_1720_v1_})
        d_1727_v8_: D7
        d_1727_v8_ = D7_DC16(d_1726_v7_)
        pat_let_tv54_ = d_1724_v5_
        index275_ = default__.safeIndex(738, (d_1719_v0_).length(0))
        def lambda100_(source21_):
            if source21_.is_DC17:
                d_1728___mcc_h0_ = source21_.cf26
                d_1729___mcc_h1_ = source21_.cf27
                d_1730___mcc_h2_ = source21_.cf28
                d_1731___mcc_h3_ = source21_.cf29
                d_1732_cf29_ = d_1731___mcc_h3_
                d_1733_cf28_ = d_1730___mcc_h2_
                d_1734_cf27_ = d_1729___mcc_h1_
                d_1735_cf26_ = d_1728___mcc_h0_
                return (_dafny.Map({False: (0) - (len(_dafny.Map({pat_let_tv54_: d_1733_cf28_})))})) in (_dafny.Map({_dafny.Map({self.f8: 521}): self.f8}))
            elif True:
                d_1736___mcc_h4_ = source21_.cf25
                d_1737_cf25_ = d_1736___mcc_h4_
                return (self).f4

        rhs229_ = ((d_1721_v2_) | (d_1721_v2_)) - (_dafny.MultiSet((d_1722_v3_) + (_dafny.SeqWithoutIsStrInference([(d_1723_v4_).cf39]))))
        rhs230_ = lambda100_(d_1727_v8_)
        lhs161_ = d_1719_v0_
        lhs162_ = default__.safeIndex(738, (d_1719_v0_).length(0))
        d_1721_v2_ = rhs229_
        lhs161_[lhs162_] = rhs230_
        index276_ = default__.safeIndex(738, (d_1719_v0_).length(0))
        (d_1719_v0_)[index276_] = self.f3
        d_1738_v9_: D7
        d_1738_v9_ = D7_DC17((self).f9, self.f3, self.f8, (self).f9)
        d_1739_v10_: _dafny.MultiSet
        d_1739_v10_ = _dafny.MultiSet([d_1738_v9_, d_1738_v9_, d_1738_v9_, D7_DC17(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bydy"))), False, self.f8, p2), d_1738_v9_])
        d_1740_v11_: C2
        nw265_ = C2()
        nw265_.ctor__(p0, not (self.f8) or ((d_1719_v0_)[default__.safeIndex(738, (d_1719_v0_).length(0))]), (not(default__.fm14(globalState)) if self.f3 else (d_1719_v0_)[default__.safeIndex(738, (d_1719_v0_).length(0))]), (d_1738_v9_) in ((d_1739_v10_).set(d_1738_v9_, default__.abs((self).f9))))
        d_1740_v11_ = nw265_
        if not(self.f3):
            d_1741_v12_: D10
            d_1741_v12_ = D10_DC23()
            d_1741_v12_ = d_1741_v12_
            d_1742_v13_: C1
            nw266_ = C1()
            nw266_.ctor__(d_1725_v6_, self.f3)
            d_1742_v13_ = nw266_
            d_1743_v14_: _dafny.Map
            d_1743_v14_ = _dafny.Map({d_1720_v1_: (p2) * ((self).f9)})
            d_1743_v14_ = (d_1743_v14_).set(_dafny.Set({default__.fm18((self).f9, d_1724_v5_, p2, globalState), (d_1740_v11_).f11, p0, (d_1740_v11_).f10, self.f8}), p2)
            d_1744_v15_: _dafny.Map
            d_1744_v15_ = _dafny.Map({(d_1742_v13_).f13: -268})
            d_1745_v16_: _dafny.Map
            d_1745_v16_ = _dafny.Map({d_1719_v0_: ((d_1744_v15_)[p0] if (p0) in (d_1744_v15_) else (self).f9)})
            d_1746_v17_: _dafny.Seq
            d_1746_v17_ = _dafny.SeqWithoutIsStrInference([d_1725_v6_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwvhmvke"))])
            d_1747_v18_: _dafny.Seq
            d_1747_v18_ = _dafny.SeqWithoutIsStrInference([len((d_1746_v17_)[default__.safeIndex((self).f9, len(d_1746_v17_))]), p2])
            d_1748_v19_: _dafny.Seq
            d_1748_v19_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p2]), d_1747_v18_, d_1747_v18_, d_1747_v18_])
            d_1749_v20_: _dafny.MultiSet
            d_1749_v20_ = _dafny.MultiSet([p2, len((d_1748_v19_)[default__.safeIndex((0) - ((0) - ((self).f9)), len(d_1748_v19_))])])
            d_1745_v16_ = (_dafny.Map({d_1719_v0_: (self).f9})) | ((d_1745_v16_ if (d_1719_v0_)[default__.safeIndex(738, (d_1719_v0_).length(0))] else _dafny.Map({d_1719_v0_: (d_1749_v20_).cardinality})))
            r1 = len(d_1725_v6_)
        elif True:
            d_1750_v21_: _dafny.Set
            d_1750_v21_ = _dafny.Set({(self).f9, len(d_1725_v6_), 702, p2, (self).f9})
            d_1751_v23_: C3
            nw267_ = C3()
            def iife150_():
                coll64_ = _dafny.Set()
                compr_64_: int
                for compr_64_ in _dafny.IntegerRange(157, 522):
                    d_1752_v22_: int = compr_64_
                    if ((157) <= (d_1752_v22_)) and ((d_1752_v22_) < (522)):
                        coll64_ = coll64_.union(_dafny.Set([(d_1752_v22_) - (p2)]))
                return _dafny.Set(coll64_)
            nw267_.ctor__((d_1719_v0_)[default__.safeIndex(738, (d_1719_v0_).length(0))], (iife150_()
            ).issubset(d_1750_v21_))
            d_1751_v23_ = nw267_
            d_1751_v23_ = d_1751_v23_
            d_1753_v24_: _dafny.Array
            def lambda101_(d_1754_i0_):
                return (d_1754_i0_) * ((self).f9)

            init53_ = lambda101_
            nw268_ = _dafny.Array(None, 29)
            for i0_53_ in range(nw268_.length(0)):
                nw268_[i0_53_] = init53_(i0_53_)
            d_1753_v24_ = nw268_
            d_1755_v25_: _dafny.Seq
            d_1755_v25_ = _dafny.SeqWithoutIsStrInference([d_1753_v24_, d_1753_v24_])
            d_1756_v26_: _dafny.Seq
            d_1756_v26_ = d_1755_v25_
            source22_ = d_1756_v26_
            d_1757___mcc_h5_ = source22_
            d_1758_cf23_ = d_1757___mcc_h5_
            (d_1740_v11_).m9(globalState)
            d_1759_v27_: _dafny.Map
            d_1759_v27_ = _dafny.Map({False: (d_1740_v11_).f11})
            d_1760_v28_: _dafny.MultiSet
            d_1760_v28_ = _dafny.MultiSet([d_1759_v27_])
            index277_ = default__.safeIndex(251, (d_1753_v24_).length(0))
            (d_1753_v24_)[index277_] = p2
            index278_ = default__.safeIndex(251, (d_1753_v24_).length(0))
            rhs231_ = d_1760_v28_
            rhs232_ = ((self).f9) - ((self).f9)
            lhs163_ = d_1753_v24_
            lhs164_ = default__.safeIndex(251, (d_1753_v24_).length(0))
            d_1760_v28_ = rhs231_
            lhs163_[lhs164_] = rhs232_
            index279_ = default__.safeIndex(251, (d_1753_v24_).length(0))
            (d_1753_v24_)[index279_] = (d_1740_v11_).fm16(p2, True, p2, p1, globalState)
            d_1761_v29_: _dafny.Seq
            d_1761_v29_ = _dafny.SeqWithoutIsStrInference([(d_1740_v11_).f10])
            d_1762_v30_: _dafny.MultiSet
            d_1762_v30_ = _dafny.MultiSet([(d_1761_v29_)[default__.safeIndex((self).f9, len(d_1761_v29_))], self.f8])
            d_1763_v31_: _dafny.Map
            d_1763_v31_ = _dafny.Map({d_1761_v29_: d_1762_v30_})
            d_1763_v31_ = ((_dafny.Map({d_1761_v29_: _dafny.MultiSet(d_1761_v29_)})) | (d_1763_v31_)) | (d_1763_v31_)
            (globalState).f0 = (d_1719_v0_)[default__.safeIndex(738, (d_1719_v0_).length(0))]
            d_1764_v32_: _dafny.Set
            d_1764_v32_ = _dafny.Set({d_1719_v0_})
            if (_dafny.Set({d_1719_v0_})).ispropersubset(d_1764_v32_):
                index280_ = default__.safeIndex(738, (d_1719_v0_).length(0))
                rhs233_ = (d_1740_v11_).f10
                rhs234_ = 220
                rhs235_ = d_1724_v5_
                lhs165_ = d_1719_v0_
                lhs166_ = default__.safeIndex(738, (d_1719_v0_).length(0))
                lhs165_[lhs166_] = rhs233_
                r1 = rhs234_
                d_1724_v5_ = rhs235_
                d_1765_v34_: _dafny.Map
                d_1765_v34_ = _dafny.Map({261: p2})
                d_1766_v35_: _dafny.MultiSet
                def iife151_():
                    coll65_ = _dafny.Map()
                    compr_65_: int
                    for compr_65_ in (d_1765_v34_).keys.Elements:
                        d_1767_v33_: int = compr_65_
                        if (d_1767_v33_) in (d_1765_v34_):
                            coll65_[default__.safeModuloInt(d_1767_v33_, len(d_1720_v1_))] = p2
                    return _dafny.Map(coll65_)
                d_1766_v35_ = _dafny.MultiSet([iife151_()
                , d_1765_v34_, d_1765_v34_])
                d_1768_v36_: _dafny.MultiSet
                d_1768_v36_ = _dafny.MultiSet([not((d_1719_v0_)[default__.safeIndex(738, (d_1719_v0_).length(0))]), (d_1740_v11_).f11])
                pat_let_tv55_ = p2
                def iife152_(_pat_let43_0):
                    def iife153_(d_1769_dt__update__tmp_h0_):
                        def iife154_(_pat_let44_0):
                            def iife155_(d_1770_dt__update_hcf61_h0_):
                                def iife156_(_pat_let45_0):
                                    def iife157_(d_1771_dt__update_hcf58_h0_):
                                        return D17_DC39((d_1769_dt__update__tmp_h0_).cf57, d_1771_dt__update_hcf58_h0_, (d_1769_dt__update__tmp_h0_).cf59, (d_1769_dt__update__tmp_h0_).cf60, d_1770_dt__update_hcf61_h0_)
                                    return iife157_(_pat_let45_0)
                                return iife156_((self).f4)
                            return iife155_(_pat_let44_0)
                        return iife154_(pat_let_tv55_)
                    return iife153_(_pat_let43_0)
                r1 = ((iife152_(D17_DC39(d_1753_v24_, p0, (d_1740_v11_).f10, len((d_1751_v23_).fm7((d_1719_v0_)[default__.safeIndex(738, (d_1719_v0_).length(0))], d_1726_v7_, default__.fm21((self).f9, d_1766_v35_, globalState), d_1768_v36_, globalState)), p2))).cf60) - (default__.safeModuloInt((self).f9, (self).f9))
                index281_ = default__.safeIndex(504, (d_1753_v24_).length(0))
                (d_1753_v24_)[index281_] = (self).f9
                index282_ = default__.safeIndex(504, (d_1753_v24_).length(0))
                index283_ = default__.safeIndex(738, (d_1719_v0_).length(0))
                index284_ = default__.safeIndex(738, (d_1719_v0_).length(0))
                rhs236_ = default__.safeDivisionInt(p2, (p2) + (p2))
                rhs237_ = (self).f4
                rhs238_ = ((0) - (default__.safeDivisionInt(p2, p2))) + (p2)
                rhs239_ = (d_1740_v11_).f10
                lhs167_ = d_1753_v24_
                lhs168_ = default__.safeIndex(504, (d_1753_v24_).length(0))
                lhs169_ = d_1719_v0_
                lhs170_ = default__.safeIndex(738, (d_1719_v0_).length(0))
                lhs171_ = d_1719_v0_
                lhs172_ = default__.safeIndex(738, (d_1719_v0_).length(0))
                lhs167_[lhs168_] = rhs236_
                lhs169_[lhs170_] = rhs237_
                r1 = rhs238_
                lhs171_[lhs172_] = rhs239_
                d_1772_v37_: C3
                nw269_ = C3()
                nw269_.ctor__((p2) >= (p2), (d_1740_v11_).f10)
                d_1772_v37_ = nw269_
                d_1773_v38_: _dafny.Array
                nw270_ = _dafny.Array(int(0), 21)
                d_1773_v38_ = nw270_
                d_1774_v39_: _dafny.Seq
                d_1774_v39_ = _dafny.SeqWithoutIsStrInference([(d_1740_v11_).f10, True])
                d_1775_v40_: D10
                d_1775_v40_ = D10_DC22(d_1773_v38_)
                d_1776_v41_: _dafny.Array
                nw271_ = _dafny.Array(None, 21)
                nw271_[int(0)] = d_1773_v38_
                nw271_[int(1)] = d_1773_v38_
                nw271_[int(2)] = d_1773_v38_
                nw271_[int(3)] = d_1773_v38_
                nw271_[int(4)] = d_1773_v38_
                nw271_[int(5)] = (d_1773_v38_ if (d_1774_v39_)[default__.safeIndex(len(d_1765_v34_), len(d_1774_v39_))] else d_1773_v38_)
                nw271_[int(6)] = d_1773_v38_
                nw271_[int(7)] = (d_1755_v25_)[default__.safeIndex((d_1753_v24_)[default__.safeIndex(504, (d_1753_v24_).length(0))], len(d_1755_v25_))]
                nw271_[int(8)] = d_1773_v38_
                nw271_[int(9)] = d_1773_v38_
                nw271_[int(10)] = d_1773_v38_
                nw271_[int(11)] = d_1773_v38_
                nw271_[int(12)] = d_1773_v38_
                nw271_[int(13)] = d_1773_v38_
                nw271_[int(14)] = d_1773_v38_
                nw271_[int(15)] = d_1773_v38_
                nw271_[int(16)] = d_1773_v38_
                nw271_[int(17)] = d_1773_v38_
                nw271_[int(18)] = (d_1775_v40_).cf36
                nw271_[int(19)] = d_1773_v38_
                nw271_[int(20)] = d_1773_v38_
                d_1776_v41_ = nw271_
                d_1777_v42_: _dafny.Seq
                d_1777_v42_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_1753_v24_)[default__.safeIndex(504, (d_1753_v24_).length(0))]: d_1724_v5_})), p2, len(d_1774_v39_)])
                index285_ = default__.safeIndex(187, (d_1776_v41_).length(0))
                (d_1776_v41_)[index285_] = (d_1755_v25_)[default__.safeIndex(len(d_1777_v42_), len(d_1755_v25_))]
                index286_ = default__.safeIndex(187, (d_1776_v41_).length(0))
                (d_1776_v41_)[index286_] = d_1773_v38_
            elif True:
                d_1778_v43_: _dafny.Map
                d_1778_v43_ = _dafny.Map({not(False): d_1724_v5_})
                d_1779_v44_: _dafny.Map
                d_1779_v44_ = _dafny.Map({p0: (self).fm5((d_1740_v11_).f10, d_1725_v6_, globalState)})
                d_1780_v45_: _dafny.MultiSet
                d_1780_v45_ = _dafny.MultiSet([len(d_1779_v44_)])
                d_1781_v46_: _dafny.MultiSet
                d_1781_v46_ = _dafny.MultiSet([len(d_1778_v43_), (d_1780_v45_).cardinality])
                index287_ = default__.safeIndex(738, (d_1719_v0_).length(0))
                rhs240_ = self.f8
                rhs241_ = (d_1780_v45_).issubset((d_1781_v46_) - (d_1781_v46_))
                rhs242_ = d_1725_v6_
                rhs243_ = ((295) >= (p2) if not (default__.fm18((self).f9, d_1724_v5_, p2, globalState)) or ((d_1740_v11_).f11) else p0)
                lhs173_ = d_1719_v0_
                lhs174_ = default__.safeIndex(738, (d_1719_v0_).length(0))
                lhs175_ = globalState
                lhs176_ = self
                lhs173_[lhs174_] = rhs240_
                lhs175_.f0 = rhs241_
                d_1725_v6_ = rhs242_
                lhs176_.f3 = rhs243_
                d_1782_v47_: _dafny.Map
                d_1782_v47_ = _dafny.Map({(d_1740_v11_).f10: (self).f4})
                d_1783_v48_: C2
                nw272_ = C2()
                nw272_.ctor__((d_1740_v11_).f10, (d_1725_v6_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffkltwrcq"))), (d_1740_v11_).f11, (default__.fm18(len(d_1782_v47_), d_1724_v5_, len((d_1740_v11_).fm17(globalState)), globalState)) or ((d_1719_v0_)[default__.safeIndex(738, (d_1719_v0_).length(0))]))
                d_1783_v48_ = nw272_
                d_1784_v49_: C3
                nw273_ = C3()
                nw273_.ctor__((d_1740_v11_).f11, (d_1724_v5_) not in ((d_1740_v11_).fm17(globalState)))
                d_1784_v49_ = nw273_
                d_1785_v50_: _dafny.Map
                d_1785_v50_ = _dafny.Map({p2: p0})
                d_1785_v50_ = (d_1785_v50_).set(len((d_1725_v6_) + (d_1725_v6_)), (d_1740_v11_).f11)
                d_1786_v51_: D0
                d_1786_v51_ = D0_DC1((d_1740_v11_).f11, d_1724_v5_)
                d_1787_v52_: _dafny.Map
                d_1787_v52_ = _dafny.Map({871: d_1724_v5_})
                d_1788_v53_: _dafny.Seq
                d_1788_v53_ = _dafny.SeqWithoutIsStrInference([self.f3, True, (d_1740_v11_).f10, (d_1740_v11_).f11])
                d_1789_v54_: _dafny.Array
                nw274_ = _dafny.Array(None, 25)
                nw274_[int(0)] = d_1724_v5_
                nw274_[int(1)] = d_1724_v5_
                nw274_[int(2)] = d_1724_v5_
                nw274_[int(3)] = d_1724_v5_
                nw274_[int(4)] = d_1724_v5_
                nw274_[int(5)] = d_1724_v5_
                nw274_[int(6)] = (d_1786_v51_).cf2
                nw274_[int(7)] = d_1724_v5_
                nw274_[int(8)] = _dafny.CodePoint('o')
                nw274_[int(9)] = d_1724_v5_
                nw274_[int(10)] = d_1724_v5_
                nw274_[int(11)] = default__.fm8(60, globalState)
                nw274_[int(12)] = d_1724_v5_
                nw274_[int(13)] = d_1724_v5_
                nw274_[int(14)] = d_1724_v5_
                nw274_[int(15)] = d_1724_v5_
                nw274_[int(16)] = ((d_1787_v52_)[len(d_1788_v53_)] if (len(d_1788_v53_)) in (d_1787_v52_) else _dafny.CodePoint('w'))
                nw274_[int(17)] = d_1724_v5_
                nw274_[int(18)] = d_1724_v5_
                nw274_[int(19)] = d_1724_v5_
                nw274_[int(20)] = default__.fm8(len(_dafny.Set({p2})), globalState)
                nw274_[int(21)] = d_1724_v5_
                nw274_[int(22)] = (d_1725_v6_)[default__.safeIndex(875, len(d_1725_v6_))]
                nw274_[int(23)] = _dafny.CodePoint('q')
                nw274_[int(24)] = default__.fm23((d_1740_v11_).f11, len(d_1725_v6_), globalState)
                d_1789_v54_ = nw274_
                index288_ = default__.safeIndex(42, (d_1789_v54_).length(0))
                (d_1789_v54_)[index288_] = d_1724_v5_
                index289_ = default__.safeIndex(42, (d_1789_v54_).length(0))
                (d_1789_v54_)[index289_] = _dafny.CodePoint('v')
            d_1790_v55_: _dafny.MultiSet
            d_1790_v55_ = _dafny.MultiSet([d_1725_v6_])
            r1 = (((d_1790_v55_)[d_1725_v6_] if (d_1725_v6_) in (d_1790_v55_) else (self).f9)) * (default__.safeDivisionInt(p2, p2))
        d_1791_v56_: _dafny.Array
        def lambda102_(d_1792_p2_):
            def lambda103_(d_1793_i1_):
                return (d_1793_i1_) - (d_1792_p2_)

            return lambda103_

        init54_ = lambda102_(p2)
        nw275_ = _dafny.Array(None, 12)
        for i0_54_ in range(nw275_.length(0)):
            nw275_[i0_54_] = init54_(i0_54_)
        d_1791_v56_ = nw275_
        r0 = d_1791_v56_
        d_1794_v57_: _dafny.Array
        nw276_ = _dafny.Array(None, 24)
        nw276_[int(0)] = d_1740_v11_
        nw276_[int(1)] = d_1740_v11_
        nw276_[int(2)] = d_1740_v11_
        nw276_[int(3)] = d_1740_v11_
        nw276_[int(4)] = d_1740_v11_
        nw276_[int(5)] = d_1740_v11_
        nw276_[int(6)] = d_1740_v11_
        nw276_[int(7)] = d_1740_v11_
        nw276_[int(8)] = d_1740_v11_
        nw276_[int(9)] = d_1740_v11_
        nw276_[int(10)] = d_1740_v11_
        nw276_[int(11)] = d_1740_v11_
        nw276_[int(12)] = d_1740_v11_
        nw276_[int(13)] = d_1740_v11_
        nw276_[int(14)] = d_1740_v11_
        nw276_[int(15)] = d_1740_v11_
        nw276_[int(16)] = d_1740_v11_
        nw276_[int(17)] = d_1740_v11_
        nw276_[int(18)] = d_1740_v11_
        nw276_[int(19)] = d_1740_v11_
        nw276_[int(20)] = d_1740_v11_
        nw276_[int(21)] = d_1740_v11_
        nw276_[int(22)] = d_1740_v11_
        nw276_[int(23)] = d_1740_v11_
        d_1794_v57_ = nw276_
        d_1795_v58_: _dafny.Seq
        d_1795_v58_ = _dafny.SeqWithoutIsStrInference([d_1794_v57_, d_1794_v57_])
        d_1796_v59_: _dafny.Set
        d_1796_v59_ = _dafny.Set({p2, (self).f9})
        d_1797_v60_: _dafny.Seq
        d_1797_v60_ = _dafny.SeqWithoutIsStrInference([(d_1795_v58_)[default__.safeIndex(len(d_1796_v59_), len(d_1795_v58_))], d_1794_v57_, d_1794_v57_])
        d_1798_v61_: _dafny.Map
        d_1798_v61_ = _dafny.Map({p2: (self).f9})
        r1 = (len(d_1797_v60_)) + (len((d_1798_v61_) | (d_1798_v61_)))
        return r0, r1

    def m5(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        d_1799_v0_: str
        d_1799_v0_ = _dafny.CodePoint('w')
        d_1800_v1_: D8
        d_1800_v1_ = D8_DC20(d_1799_v0_, p1)
        d_1801_v2_: _dafny.Seq
        d_1801_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxqsrb"))
        d_1802_v4_: C0
        nw277_ = C0()
        nw277_.ctor__()
        d_1802_v4_ = nw277_
        d_1803_v5_: _dafny.Array
        nw278_ = _dafny.Array(None, 19)
        nw278_[int(0)] = d_1802_v4_
        nw278_[int(1)] = d_1802_v4_
        nw278_[int(2)] = d_1802_v4_
        nw278_[int(3)] = d_1802_v4_
        nw278_[int(4)] = d_1802_v4_
        nw278_[int(5)] = d_1802_v4_
        nw278_[int(6)] = d_1802_v4_
        nw278_[int(7)] = d_1802_v4_
        nw278_[int(8)] = d_1802_v4_
        nw278_[int(9)] = d_1802_v4_
        nw278_[int(10)] = d_1802_v4_
        nw278_[int(11)] = d_1802_v4_
        nw278_[int(12)] = d_1802_v4_
        nw278_[int(13)] = d_1802_v4_
        nw278_[int(14)] = d_1802_v4_
        nw278_[int(15)] = d_1802_v4_
        nw278_[int(16)] = d_1802_v4_
        nw278_[int(17)] = d_1802_v4_
        nw278_[int(18)] = d_1802_v4_
        d_1803_v5_ = nw278_
        d_1804_v6_: _dafny.Map
        d_1804_v6_ = _dafny.Map({p0: d_1803_v5_})
        d_1805_v7_: D19
        d_1805_v7_ = D19_DC44(d_1804_v6_)
        d_1806_v8_: _dafny.Array
        nw279_ = _dafny.Array(None, 10)
        nw279_[int(0)] = (d_1800_v1_).cf34
        nw279_[int(1)] = self.f8
        nw279_[int(2)] = p0
        nw279_[int(3)] = (D7_DC17(len(d_1801_v2_), self.f8, (self).f4, (self).f9)).cf28
        nw279_[int(4)] = self.f8
        nw279_[int(5)] = self.f8
        def iife158_():
            coll66_ = _dafny.Map()
            compr_66_: int
            for compr_66_ in _dafny.IntegerRange(128, 194):
                d_1807_v3_: int = compr_66_
                if ((128) <= (d_1807_v3_)) and ((d_1807_v3_) < (194)):
                    coll66_[default__.safeDivisionInt(d_1807_v3_, (self).f9)] = (self).f9
            return _dafny.Map(coll66_)
        nw279_[int(6)] = (len(iife158_()
        )) > ((self).f9)
        nw279_[int(7)] = ((self).f4) not in ((d_1805_v7_).cf68)
        nw279_[int(8)] = (self.f3 if (D0_DC1(True, d_1799_v0_)).cf1 else self.f8)
        nw279_[int(9)] = (self).f4
        d_1806_v8_ = nw279_
        index290_ = default__.safeIndex(217, (d_1806_v8_).length(0))
        (d_1806_v8_)[index290_] = default__.fm14(globalState)
        index291_ = default__.safeIndex(217, (d_1806_v8_).length(0))
        (d_1806_v8_)[index291_] = self.f3
        if self.f3:
            d_1808_v9_: _dafny.Set
            d_1808_v9_ = _dafny.Set({-848})
            d_1809_v10_: _dafny.Map
            d_1809_v10_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvjswysav")): len(d_1808_v9_)})
            r0 = ((d_1809_v10_)[d_1801_v2_] if (d_1801_v2_) in (d_1809_v10_) else (0) - ((self).f9))
            d_1810_v11_: C3
            nw280_ = C3()
            nw280_.ctor__((d_1806_v8_)[default__.safeIndex(217, (d_1806_v8_).length(0))], (d_1806_v8_)[default__.safeIndex(217, (d_1806_v8_).length(0))])
            d_1810_v11_ = nw280_
            (globalState).f0 = p1
            d_1811_v12_: _dafny.Seq
            d_1811_v12_ = _dafny.SeqWithoutIsStrInference([(self).f9])
            d_1812_v14_: D8
            def iife159_():
                coll67_ = _dafny.Map()
                compr_67_: int
                for compr_67_ in (d_1811_v12_).Elements:
                    d_1813_v13_: int = compr_67_
                    if (d_1813_v13_) in (d_1811_v12_):
                        coll67_[(d_1813_v13_) + ((self).f9)] = False
                return _dafny.Map(coll67_)
            d_1812_v14_ = D8_DC18(iife159_()
)
            index292_ = default__.safeIndex(217, (d_1806_v8_).length(0))
            rhs244_ = not ((d_1811_v12_) < (d_1811_v12_)) or (False)
            rhs245_ = (0) - ((((self).f9) * (len((d_1812_v14_).cf30))) * ((self).f9))
            lhs177_ = d_1806_v8_
            lhs178_ = default__.safeIndex(217, (d_1806_v8_).length(0))
            lhs177_[lhs178_] = rhs244_
            r0 = rhs245_
            d_1814_v15_: _dafny.Array
            nw281_ = _dafny.Array(_dafny.MultiSet({}), 13)
            d_1814_v15_ = nw281_
            d_1815_v16_: _dafny.MultiSet
            d_1815_v16_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(self).f9 for d_1816_i0_ in range(default__.abs(230))]))])
            index293_ = default__.safeIndex(828, (d_1814_v15_).length(0))
            (d_1814_v15_)[index293_] = d_1815_v16_
            d_1817_v17_: _dafny.Array
            def lambda104_(d_1818_i1_):
                return default__.safeDivisionInt(d_1818_i1_, (self).f9)

            init55_ = lambda104_
            nw282_ = _dafny.Array(None, 9)
            for i0_55_ in range(nw282_.length(0)):
                nw282_[i0_55_] = init55_(i0_55_)
            d_1817_v17_ = nw282_
            d_1819_v18_: _dafny.Seq
            d_1819_v18_ = _dafny.SeqWithoutIsStrInference([not(False)])
            d_1820_v19_: _dafny.Map
            d_1820_v19_ = _dafny.Map({len(d_1819_v18_): p0})
            d_1821_v20_: D11
            d_1821_v20_ = D11_DC25(d_1820_v19_)
            d_1822_v21_: _dafny.Seq
            d_1822_v21_ = _dafny.SeqWithoutIsStrInference([d_1821_v20_])
            d_1823_v22_: _dafny.Map
            d_1823_v22_ = _dafny.Map({d_1822_v21_: (self).f4})
            index294_ = default__.safeIndex(301, (d_1817_v17_).length(0))
            (d_1817_v17_)[index294_] = (917 if ((d_1823_v22_)[d_1822_v21_] if (d_1822_v21_) in (d_1823_v22_) else p0) else (self).f9)
            index295_ = default__.safeIndex(828, (d_1814_v15_).length(0))
            index296_ = default__.safeIndex(301, (d_1817_v17_).length(0))
            rhs246_ = d_1815_v16_
            rhs247_ = (self).f9
            rhs248_ = (self).f9
            lhs179_ = d_1814_v15_
            lhs180_ = default__.safeIndex(828, (d_1814_v15_).length(0))
            lhs181_ = d_1817_v17_
            lhs182_ = default__.safeIndex(301, (d_1817_v17_).length(0))
            lhs179_[lhs180_] = rhs246_
            lhs181_[lhs182_] = rhs247_
            r0 = rhs248_
        elif True:
            r0 = (self).f9
            r0 = ((self).f9) + ((0) - (default__.safeModuloInt(-505, (self).f9)))
            d_1824_v23_: C3
            nw283_ = C3()
            nw283_.ctor__(p0, (d_1806_v8_)[default__.safeIndex(217, (d_1806_v8_).length(0))])
            d_1824_v23_ = nw283_
            d_1825_v24_: _dafny.Map
            d_1825_v24_ = _dafny.Map({(self).f9: self.f8})
            d_1826_v25_: _dafny.Map
            d_1826_v25_ = _dafny.Map({len(d_1825_v24_): not(not((d_1806_v8_)[default__.safeIndex(217, (d_1806_v8_).length(0))]))})
            d_1827_v26_: D8
            d_1827_v26_ = D8_DC18(d_1825_v24_)
            d_1827_v26_ = d_1827_v26_
            d_1828_v27_: _dafny.Seq
            d_1828_v27_ = _dafny.SeqWithoutIsStrInference([p0, p1])
            d_1829_v28_: _dafny.Seq
            d_1829_v28_ = _dafny.SeqWithoutIsStrInference([d_1828_v27_, d_1828_v27_, d_1828_v27_, (default__.fm34(654, (self).f9, globalState)) + (d_1828_v27_)])
            d_1829_v28_ = _dafny.SeqWithoutIsStrInference([d_1828_v27_, (default__.fm34((self).f9, (self).f9, globalState)) + (_dafny.SeqWithoutIsStrInference([not(False), (d_1806_v8_)[default__.safeIndex(217, (d_1806_v8_).length(0))], p1]))])
        d_1830_v29_: _dafny.Array
        def lambda105_(d_1831_i3_):
            return (d_1831_i3_) * ((self).f9)

        init56_ = lambda105_
        nw284_ = _dafny.Array(None, 2)
        for i0_56_ in range(nw284_.length(0)):
            nw284_[i0_56_] = init56_(i0_56_)
        d_1830_v29_ = nw284_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1830_v29_).length(0)):
            d_1832_i2_: int = guard_loop_4_
            if (True) and (((0) <= (d_1832_i2_)) and ((d_1832_i2_) < ((d_1830_v29_).length(0)))):
                (d_1830_v29_)[(d_1832_i2_)] = (d_1832_i2_) - (((self).f9 if p0 else 4))
        d_1801_v2_ = d_1801_v2_
        if self.f3:
            d_1833_v30_: D1
            d_1833_v30_ = D1_DC3(d_1806_v8_)
            pat_let_tv56_ = d_1806_v8_
            def iife160_(_pat_let46_0):
                def iife161_(d_1834_dt__update__tmp_h0_):
                    def iife162_(_pat_let47_0):
                        def iife163_(d_1835_dt__update_hcf6_h0_):
                            return D1_DC3(d_1835_dt__update_hcf6_h0_)
                        return iife163_(_pat_let47_0)
                    return iife162_(pat_let_tv56_)
                return iife161_(_pat_let46_0)
            source23_ = iife160_(d_1833_v30_)
            if source23_.is_DC4:
                d_1836___mcc_h0_ = source23_.cf7
                d_1837_cf7_ = d_1836___mcc_h0_
                nw285_ = _dafny.Array(int(0), 15)
                d_1830_v29_ = nw285_
                d_1838_v31_: _dafny.Map
                d_1838_v31_ = _dafny.Map({self.f8: (d_1806_v8_)[default__.safeIndex(217, (d_1806_v8_).length(0))]})
                d_1838_v31_ = (d_1838_v31_).set(p0, p0)
                d_1839_v32_: bool
                d_1840_v33_: int
                d_1841_v34_: _dafny.Array
                out32_: bool
                out33_: int
                out34_: _dafny.Array
                out32_, out33_, out34_ = (self).m6(p1, globalState)
                d_1839_v32_ = out32_
                d_1840_v33_ = out33_
                d_1841_v34_ = out34_
                d_1842_v35_: _dafny.Map
                d_1842_v35_ = _dafny.Map({d_1840_v33_: d_1837_cf7_})
                d_1840_v33_ = (0) - (((self).f9 if (d_1806_v8_)[default__.safeIndex(217, (d_1806_v8_).length(0))] else len(d_1842_v35_)))
            elif source23_.is_DC5:
                d_1843___mcc_h1_ = source23_.cf8
                d_1844___mcc_h2_ = source23_.cf9
                d_1845___mcc_h3_ = source23_.cf10
                d_1846___mcc_h4_ = source23_.cf11
                d_1847_cf11_ = d_1846___mcc_h4_
                d_1848_cf10_ = d_1845___mcc_h3_
                d_1849_cf9_ = d_1844___mcc_h2_
                d_1850_cf8_ = d_1843___mcc_h1_
                d_1849_cf9_ = default__.fm14(globalState)
                d_1849_cf9_ = False
                d_1851_v36_: _dafny.Map
                d_1851_v36_ = _dafny.Map({True: d_1848_cf10_})
                index297_ = default__.safeIndex(774, (d_1830_v29_).length(0))
                (d_1830_v29_)[index297_] = (len((d_1851_v36_).set(p1, (self).f9))) + (d_1847_cf11_)
                index298_ = default__.safeIndex(774, (d_1830_v29_).length(0))
                (d_1830_v29_)[index298_] = (self).f9
                d_1852_v37_: _dafny.Array
                nw286_ = _dafny.Array(_dafny.MultiSet({}), 8)
                d_1852_v37_ = nw286_
                d_1853_v38_: _dafny.Set
                d_1853_v38_ = _dafny.Set({d_1848_cf10_, d_1848_cf10_})
                d_1854_v39_: D16
                d_1854_v39_ = D16_DC36(d_1853_v38_, d_1850_cf8_, (self).f4, d_1850_cf8_)
                d_1855_v40_: _dafny.MultiSet
                d_1855_v40_ = _dafny.MultiSet([d_1854_v39_, d_1854_v39_])
                index299_ = default__.safeIndex(86, (d_1852_v37_).length(0))
                (d_1852_v37_)[index299_] = d_1855_v40_
                index300_ = default__.safeIndex(86, (d_1852_v37_).length(0))
                (d_1852_v37_)[index300_] = d_1855_v40_
            elif source23_.is_DC6:
                (globalState).f0 = ((_dafny.SeqWithoutIsStrInference([d_1799_v0_ for d_1856_i4_ in range(default__.abs(263))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ng")))) < ((d_1801_v2_) + (d_1801_v2_))
                d_1857_v41_: bool
                d_1858_v42_: int
                d_1859_v43_: _dafny.Array
                out35_: bool
                out36_: int
                out37_: _dafny.Array
                out35_, out36_, out37_ = (self).m6(p1, globalState)
                d_1857_v41_ = out35_
                d_1858_v42_ = out36_
                d_1859_v43_ = out37_
                r0 = (self).f9
                d_1860_v44_: C1
                nw287_ = C1()
                nw287_.ctor__(d_1801_v2_, d_1857_v41_)
                d_1860_v44_ = nw287_
            elif True:
                d_1861___mcc_h5_ = source23_.cf6
                d_1862_cf6_ = d_1861___mcc_h5_
                d_1863_v45_: _dafny.MultiSet
                d_1863_v45_ = _dafny.MultiSet([(d_1806_v8_)[default__.safeIndex(217, (d_1806_v8_).length(0))], default__.fm6((self).f9, globalState)])
                d_1864_v46_: _dafny.Map
                d_1864_v46_ = _dafny.Map({self.f3: (d_1806_v8_)[default__.safeIndex(217, (d_1806_v8_).length(0))]})
                d_1865_v47_: _dafny.Seq
                d_1865_v47_ = _dafny.SeqWithoutIsStrInference([self.f3])
                rhs249_ = ((d_1864_v46_)[self.f3] if (self.f3) in (d_1864_v46_) else (d_1865_v47_)[default__.safeIndex(818, len(d_1865_v47_))])
                rhs250_ = ((_dafny.MultiSet([not((self).f4)])).intersection(d_1863_v45_)) - (d_1863_v45_)
                lhs183_ = self
                lhs183_.f8 = rhs249_
                d_1863_v45_ = rhs250_
                r0 = (self).f9
                d_1866_v48_: _dafny.Map
                d_1866_v48_ = _dafny.Map({False: (self).f9})
                d_1867_v49_: _dafny.Seq
                d_1867_v49_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                d_1868_v50_: _dafny.Map
                d_1868_v50_ = _dafny.Map({(self).f9: (d_1802_v4_).fm10(((d_1866_v48_)[self.f3] if (self.f3) in (d_1866_v48_) else (self).f9), p0, _dafny.MultiSet(d_1867_v49_), globalState)})
                d_1869_v51_: _dafny.Set
                d_1869_v51_ = _dafny.Set({self.f8, p1, (self).f4})
                r0 = ((d_1868_v50_)[(self).f9] if ((self).f9) in (d_1868_v50_) else len(d_1869_v51_))
                d_1870_v52_: _dafny.Map
                d_1870_v52_ = _dafny.Map({p0: d_1806_v8_})
                d_1870_v52_ = (d_1870_v52_).set((d_1806_v8_)[default__.safeIndex(217, (d_1806_v8_).length(0))], d_1862_cf6_)
            d_1871_v53_: _dafny.Map
            d_1871_v53_ = _dafny.Map({(self).f9: self.f8})
            d_1872_v54_: C2
            nw288_ = C2()
            nw288_.ctor__(not (((d_1871_v53_)[(self).f9] if ((self).f9) in (d_1871_v53_) else (self).f4)) or (p1), p1, p1, True)
            d_1872_v54_ = nw288_
            d_1873_v55_: _dafny.Seq
            d_1873_v55_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1874_v56_: _dafny.Map
            d_1874_v56_ = _dafny.Map({d_1873_v55_: ((d_1872_v54_).fm17(globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))})
            d_1874_v56_ = (d_1874_v56_).set(d_1873_v55_, (d_1801_v2_).set(default__.safeIndex((self).f9, len(d_1801_v2_)), d_1799_v0_))
            index301_ = default__.safeIndex(369, (d_1830_v29_).length(0))
            (d_1830_v29_)[index301_] = 85
            index302_ = default__.safeIndex(369, (d_1830_v29_).length(0))
            (d_1830_v29_)[index302_] = (self).f9
            if (self).f4:
                d_1875_v57_: D19
                d_1875_v57_ = D19_DC45(default__.safeModuloInt(47, (d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]), (self).f9, (d_1872_v54_).f10, (d_1872_v54_).f11, d_1873_v55_)
                d_1875_v57_ = d_1875_v57_
                d_1876_v58_: _dafny.Map
                d_1876_v58_ = _dafny.Map({(d_1801_v2_) + (d_1801_v2_): not((self).f4)})
                d_1877_v59_: D1
                d_1877_v59_ = D1_DC5((self).f4, (d_1872_v54_).f11, (self).f9, (self).f9)
                index303_ = default__.safeIndex(217, (d_1806_v8_).length(0))
                (d_1806_v8_)[index303_] = ((d_1876_v58_)[default__.fm30(d_1877_v59_, len(d_1871_v53_), (self).f9, globalState)] if (default__.fm30(d_1877_v59_, len(d_1871_v53_), (self).f9, globalState)) in (d_1876_v58_) else False)
                d_1878_v60_: D2
                d_1878_v60_ = D2_DC8(d_1806_v8_, d_1801_v2_)
                d_1879_v61_: _dafny.Map
                d_1879_v61_ = _dafny.Map({d_1878_v60_: (self).f9})
                index304_ = default__.safeIndex(217, (d_1806_v8_).length(0))
                (d_1806_v8_)[index304_] = (default__.safeDivisionInt((d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))], len(_dafny.Map({d_1879_v61_: p0})))) >= ((self).fm5((d_1872_v54_).f10, d_1801_v2_, globalState))
                index305_ = default__.safeIndex(369, (d_1830_v29_).length(0))
                (d_1830_v29_)[index305_] = (d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]
                d_1880_v62_: _dafny.Seq
                d_1880_v62_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                d_1881_v63_: _dafny.Map
                d_1881_v63_ = _dafny.Map({d_1880_v62_: d_1799_v0_})
                d_1882_v64_: D17
                d_1882_v64_ = D17_DC37(d_1881_v63_)
                d_1883_v65_: _dafny.Map
                d_1883_v65_ = _dafny.Map({d_1882_v64_: (d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]})
                d_1883_v65_ = (d_1883_v65_).set(D17_DC37(d_1881_v63_), (d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))])
            elif True:
                d_1801_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "empl"))
                d_1884_v66_: _dafny.Seq
                d_1884_v66_ = _dafny.SeqWithoutIsStrInference([(d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]])
                d_1885_v67_: _dafny.Map
                d_1885_v67_ = _dafny.Map({(d_1872_v54_).f10: (d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]})
                r0 = default__.fm31(not(((d_1871_v53_)[(self).f9] if ((self).f9) in (d_1871_v53_) else (d_1806_v8_)[default__.safeIndex(217, (d_1806_v8_).length(0))])), (d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))], ((d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]) not in (d_1884_v66_), (d_1873_v55_).set(default__.safeIndex(((d_1885_v67_)[((d_1871_v53_)[(0) - (len(d_1801_v2_))] if ((0) - (len(d_1801_v2_))) in (d_1871_v53_) else (self).f4)] if (((d_1871_v53_)[(0) - (len(d_1801_v2_))] if ((0) - (len(d_1801_v2_))) in (d_1871_v53_) else (self).f4)) in (d_1885_v67_) else len(d_1873_v55_)), len(d_1873_v55_)), (self).f4), globalState)
                d_1886_v68_: _dafny.Map
                d_1886_v68_ = _dafny.Map({d_1806_v8_: (d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]})
                d_1887_v69_: D18
                d_1887_v69_ = D18_DC42((self).f9)
                d_1886_v68_ = (d_1886_v68_).set(d_1806_v8_, len(((_dafny.SeqWithoutIsStrInference([(self).f9])).set(default__.safeIndex((self).f9, len(_dafny.SeqWithoutIsStrInference([(self).f9]))), (self).f9)) + (_dafny.SeqWithoutIsStrInference([(d_1887_v69_).cf64]))))
                d_1888_v70_: _dafny.Array
                nw289_ = _dafny.Array(None, 13)
                nw289_[int(0)] = (d_1884_v66_) + (d_1884_v66_)
                nw289_[int(1)] = d_1884_v66_
                nw289_[int(2)] = d_1884_v66_
                nw289_[int(3)] = d_1884_v66_
                nw289_[int(4)] = _dafny.SeqWithoutIsStrInference([(d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))] for d_1889_i5_ in range(default__.abs(-229))])
                nw289_[int(5)] = (_dafny.SeqWithoutIsStrInference([(self).f9, (self).f9, (self).f9])) + (d_1884_v66_)
                nw289_[int(6)] = (default__.fm19(globalState)).set(default__.safeIndex(len(d_1884_v66_), len(default__.fm19(globalState))), 158)
                nw289_[int(7)] = d_1884_v66_
                nw289_[int(8)] = d_1884_v66_
                nw289_[int(9)] = _dafny.SeqWithoutIsStrInference([len((D20_DC48(_dafny.SeqWithoutIsStrInference([(d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]]))).cf80) for d_1890_i6_ in range(default__.abs(-327))])
                nw289_[int(10)] = d_1884_v66_
                nw289_[int(11)] = (d_1884_v66_ if p1 else d_1884_v66_)
                nw289_[int(12)] = d_1884_v66_
                d_1888_v70_ = nw289_
                d_1891_v72_: _dafny.Array
                def lambda106_(d_1892_v55_, d_1893_v2_, d_1894_v54_, d_1895_v0_, d_1896_v29_):
                    def lambda107_(d_1897_i7_):
                        def iife164_():
                            coll68_ = _dafny.Map()
                            compr_68_: int
                            for compr_68_ in _dafny.IntegerRange(137, -50):
                                d_1898_v71_: int = compr_68_
                                if ((137) <= (d_1898_v71_)) and ((d_1898_v71_) < (-50)):
                                    coll68_[default__.safeDivisionInt(d_1898_v71_, len(d_1893_v2_))] = _dafny.Map({(d_1894_v54_).f10: d_1895_v0_})
                            return _dafny.Map(coll68_)
                        return (iife164_()
                        ) | (_dafny.Map({len(_dafny.Set({(d_1892_v55_)[default__.safeIndex((d_1896_v29_)[default__.safeIndex(369, (d_1896_v29_).length(0))], len(d_1892_v55_))]})): _dafny.Map({(d_1894_v54_).f11: _dafny.CodePoint('p')})}))

                    return lambda107_

                init57_ = lambda106_(d_1873_v55_, d_1801_v2_, d_1872_v54_, d_1799_v0_, d_1830_v29_)
                nw290_ = _dafny.Array(None, 24)
                for i0_57_ in range(nw290_.length(0)):
                    nw290_[i0_57_] = init57_(i0_57_)
                d_1891_v72_ = nw290_
                d_1899_v73_: _dafny.Map
                d_1899_v73_ = _dafny.Map({not((d_1806_v8_)[default__.safeIndex(217, (d_1806_v8_).length(0))]): d_1799_v0_})
                index306_ = default__.safeIndex(696, (d_1891_v72_).length(0))
                (d_1891_v72_)[index306_] = _dafny.Map({(d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]: d_1899_v73_})
                d_1900_v74_: _dafny.Array
                def lambda108_(d_1901_v29_):
                    def lambda109_(d_1902_i8_):
                        return D12_DC27((d_1901_v29_)[default__.safeIndex(369, (d_1901_v29_).length(0))])

                    return lambda109_

                init58_ = lambda108_(d_1830_v29_)
                nw291_ = _dafny.Array(None, 20)
                for i0_58_ in range(nw291_.length(0)):
                    nw291_[i0_58_] = init58_(i0_58_)
                d_1900_v74_ = nw291_
                d_1903_v75_: _dafny.Set
                d_1903_v75_ = _dafny.Set({len(d_1873_v55_)})
                d_1904_v76_: _dafny.Map
                d_1904_v76_ = _dafny.Map({(d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]: d_1903_v75_})
                d_1905_v77_: _dafny.Map
                d_1905_v77_ = _dafny.Map({(self).f9: d_1899_v73_})
                d_1906_v78_: _dafny.Seq
                d_1906_v78_ = _dafny.SeqWithoutIsStrInference([d_1900_v74_, (d_1900_v74_ if p0 else d_1900_v74_), d_1900_v74_, d_1900_v74_, d_1900_v74_])
                index307_ = default__.safeIndex(696, (d_1891_v72_).length(0))
                rhs251_ = ((d_1903_v75_).intersection(((d_1904_v76_)[(self).f9] if ((self).f9) in (d_1904_v76_) else _dafny.Set({(self).f9, (d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]})))).ispropersubset(d_1903_v75_)
                rhs252_ = d_1888_v70_
                rhs253_ = d_1905_v77_
                rhs254_ = (d_1906_v78_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([len(d_1801_v2_), (d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]])), len(d_1906_v78_))]
                lhs184_ = d_1891_v72_
                lhs185_ = default__.safeIndex(696, (d_1891_v72_).length(0))
                r1 = rhs251_
                d_1888_v70_ = rhs252_
                lhs184_[lhs185_] = rhs253_
                d_1900_v74_ = rhs254_
                d_1907_v79_: _dafny.Seq
                d_1907_v79_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsnalpc"))])
                d_1908_v80_: _dafny.Map
                d_1908_v80_ = _dafny.Map({(d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]: d_1907_v79_})
                d_1909_v81_: _dafny.Seq
                d_1909_v81_ = ((d_1908_v80_)[(d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]] if ((d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]) in (d_1908_v80_) else d_1907_v79_)
                d_1910_v82_: _dafny.Array
                nw292_ = _dafny.Array(None, 28)
                nw292_[int(0)] = d_1909_v81_
                nw292_[int(1)] = default__.fm35((d_1872_v54_).f10, globalState)
                nw292_[int(2)] = d_1909_v81_
                nw292_[int(3)] = d_1909_v81_
                nw292_[int(4)] = d_1909_v81_
                nw292_[int(5)] = d_1907_v79_
                nw292_[int(6)] = d_1909_v81_
                nw292_[int(7)] = d_1909_v81_
                nw292_[int(8)] = d_1909_v81_
                nw292_[int(9)] = d_1909_v81_
                nw292_[int(10)] = d_1909_v81_
                nw292_[int(11)] = d_1907_v79_
                nw292_[int(12)] = d_1909_v81_
                nw292_[int(13)] = d_1909_v81_
                nw292_[int(14)] = d_1909_v81_
                nw292_[int(15)] = d_1909_v81_
                nw292_[int(16)] = d_1909_v81_
                nw292_[int(17)] = d_1909_v81_
                nw292_[int(18)] = d_1909_v81_
                nw292_[int(19)] = d_1909_v81_
                nw292_[int(20)] = d_1909_v81_
                nw292_[int(21)] = d_1907_v79_
                nw292_[int(22)] = d_1909_v81_
                nw292_[int(23)] = d_1909_v81_
                nw292_[int(24)] = d_1907_v79_
                nw292_[int(25)] = d_1907_v79_
                nw292_[int(26)] = default__.fm41((self).f9, globalState)
                nw292_[int(27)] = d_1909_v81_
                d_1910_v82_ = nw292_
                index308_ = default__.safeIndex(369, (d_1830_v29_).length(0))
                rhs255_ = d_1910_v82_
                rhs256_ = default__.safeModuloInt((self).f9, ((d_1830_v29_)[default__.safeIndex(369, (d_1830_v29_).length(0))]) * (len(d_1873_v55_)))
                rhs257_ = (self).f9
                rhs258_ = self.f3
                lhs186_ = d_1830_v29_
                lhs187_ = default__.safeIndex(369, (d_1830_v29_).length(0))
                d_1910_v82_ = rhs255_
                lhs186_[lhs187_] = rhs256_
                r0 = rhs257_
                r1 = rhs258_
        elif True:
            d_1911_v83_: _dafny.Seq
            d_1911_v83_ = _dafny.SeqWithoutIsStrInference([(self).f9, (self).f9])
            index309_ = default__.safeIndex(605, (d_1830_v29_).length(0))
            (d_1830_v29_)[index309_] = len(d_1911_v83_)
            index310_ = default__.safeIndex(605, (d_1830_v29_).length(0))
            (d_1830_v29_)[index310_] = 258
            d_1912_v84_: C3
            nw293_ = C3()
            nw293_.ctor__(self.f3, p0)
            d_1912_v84_ = nw293_
            d_1913_v85_: T0
            nw294_ = C2()
            nw294_.ctor__((self).f4, True, self.f3, (self).f4)
            d_1913_v85_ = nw294_
            d_1914_v86_: _dafny.Map
            d_1914_v86_ = _dafny.Map({d_1799_v0_: d_1913_v85_})
            d_1915_v87_: _dafny.MultiSet
            d_1915_v87_ = _dafny.MultiSet([(d_1830_v29_)[default__.safeIndex(605, (d_1830_v29_).length(0))], (d_1830_v29_)[default__.safeIndex(605, (d_1830_v29_).length(0))], len(d_1914_v86_), (self).f9])
            d_1916_v88_: _dafny.Seq
            d_1916_v88_ = _dafny.SeqWithoutIsStrInference([True, not(not((d_1806_v8_)[default__.safeIndex(217, (d_1806_v8_).length(0))])), (_dafny.MultiSet([len(d_1911_v83_)])).issubset((d_1915_v87_).set((self).f9, default__.abs(292)))])
            d_1916_v88_ = (d_1916_v88_).set(default__.safeIndex(len(d_1801_v2_), len(d_1916_v88_)), d_1913_v85_.f3)
            (globalState).f0 = d_1913_v85_.f3
            (self).f8 = (d_1913_v85_.f3) == (p0)
        index311_ = default__.safeIndex(115, (d_1830_v29_).length(0))
        (d_1830_v29_)[index311_] = (self).f9
        d_1917_v89_: _dafny.MultiSet
        d_1917_v89_ = _dafny.MultiSet([self.f3])
        index312_ = default__.safeIndex(115, (d_1830_v29_).length(0))
        (d_1830_v29_)[index312_] = (d_1917_v89_).cardinality
        d_1918_v90_: _dafny.Map
        d_1918_v90_ = _dafny.Map({(self).f9: d_1801_v2_})
        r0 = ((d_1830_v29_)[default__.safeIndex(115, (d_1830_v29_).length(0))] if ((d_1830_v29_)[default__.safeIndex(115, (d_1830_v29_).length(0))]) in (d_1918_v90_) else (self).f9)
        r1 = p1
        return r0, r1

    def m6(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        (self).f8 = p0
        d_1919_v0_: _dafny.Array
        nw295_ = _dafny.Array(_dafny.Set({}), 8)
        d_1919_v0_ = nw295_
        d_1920_v1_: T0
        nw296_ = C3()
        nw296_.ctor__(p0, not(self.f3))
        d_1920_v1_ = nw296_
        d_1921_v2_: _dafny.Seq
        d_1921_v2_ = _dafny.SeqWithoutIsStrInference([d_1920_v1_, d_1920_v1_])
        d_1922_v3_: _dafny.Map
        d_1922_v3_ = _dafny.Map({(self).f9: d_1921_v2_})
        d_1923_v4_: _dafny.Seq
        d_1923_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrkmvyivs"))
        d_1924_v5_: _dafny.Set
        d_1924_v5_ = _dafny.Set({d_1921_v2_, (((d_1922_v3_)[(self).f9] if ((self).f9) in (d_1922_v3_) else d_1921_v2_)).set(default__.safeIndex(len(d_1923_v4_), len(((d_1922_v3_)[(self).f9] if ((self).f9) in (d_1922_v3_) else d_1921_v2_))), d_1920_v1_)})
        index313_ = default__.safeIndex(336, (d_1919_v0_).length(0))
        (d_1919_v0_)[index313_] = d_1924_v5_
        index314_ = default__.safeIndex(336, (d_1919_v0_).length(0))
        rhs259_ = self.f3
        rhs260_ = ((d_1924_v5_) - (d_1924_v5_)).intersection(d_1924_v5_)
        lhs188_ = d_1919_v0_
        lhs189_ = default__.safeIndex(336, (d_1919_v0_).length(0))
        r0 = rhs259_
        lhs188_[lhs189_] = rhs260_
        d_1925_v6_: _dafny.Array
        nw297_ = _dafny.Array(_dafny.Array(None, 0), 6)
        d_1925_v6_ = nw297_
        d_1926_v7_: _dafny.Array
        def lambda110_(d_1927_i0_):
            return True

        init59_ = lambda110_
        nw298_ = _dafny.Array(None, 19)
        for i0_59_ in range(nw298_.length(0)):
            nw298_[i0_59_] = init59_(i0_59_)
        d_1926_v7_ = nw298_
        d_1928_v8_: _dafny.Set
        d_1928_v8_ = _dafny.Set({(self).f9, len(d_1923_v4_)})
        d_1929_v9_: D16
        d_1929_v9_ = D16_DC36(d_1928_v8_, (self).f4, (self).f4, False)
        nw299_ = _dafny.Array(None, 23)
        nw299_[int(0)] = d_1926_v7_
        nw299_[int(1)] = d_1926_v7_
        nw299_[int(2)] = d_1926_v7_
        nw299_[int(3)] = d_1926_v7_
        nw299_[int(4)] = d_1926_v7_
        nw299_[int(5)] = d_1926_v7_
        nw299_[int(6)] = d_1926_v7_
        nw299_[int(7)] = d_1926_v7_
        nw299_[int(8)] = d_1926_v7_
        nw299_[int(9)] = d_1926_v7_
        nw299_[int(10)] = d_1926_v7_
        nw299_[int(11)] = d_1926_v7_
        nw299_[int(12)] = d_1926_v7_
        nw299_[int(13)] = d_1926_v7_
        nw299_[int(14)] = d_1926_v7_
        nw299_[int(15)] = (d_1926_v7_ if self.f8 else d_1926_v7_)
        nw299_[int(16)] = d_1926_v7_
        nw299_[int(17)] = d_1926_v7_
        nw299_[int(18)] = d_1926_v7_
        nw299_[int(19)] = d_1926_v7_
        nw299_[int(20)] = (d_1926_v7_ if p0 else d_1926_v7_)
        nw299_[int(21)] = (d_1926_v7_ if (d_1929_v9_).cf52 else d_1926_v7_)
        nw299_[int(22)] = d_1926_v7_
        d_1925_v6_ = nw299_
        d_1928_v8_ = d_1928_v8_
        d_1930_v10_: _dafny.Map
        d_1930_v10_ = _dafny.Map({d_1920_v1_: d_1920_v1_})
        d_1931_v11_: _dafny.Seq
        d_1931_v11_ = _dafny.SeqWithoutIsStrInference([d_1930_v10_, d_1930_v10_])
        d_1932_v12_: _dafny.Array
        nw300_ = _dafny.Array(None, 3)
        nw300_[int(0)] = d_1931_v11_
        nw300_[int(1)] = d_1931_v11_
        nw300_[int(2)] = d_1931_v11_
        d_1932_v12_ = nw300_
        index315_ = default__.safeIndex(588, (d_1932_v12_).length(0))
        (d_1932_v12_)[index315_] = d_1931_v11_
        index316_ = default__.safeIndex(588, (d_1932_v12_).length(0))
        (d_1932_v12_)[index316_] = (_dafny.SeqWithoutIsStrInference([(d_1930_v10_).set(d_1920_v1_, d_1920_v1_)])) + (d_1931_v11_)
        d_1933_v13_: _dafny.Array
        def lambda111_(d_1934_i2_):
            return _dafny.SeqWithoutIsStrInference([(self).f9, (self).f9, (self).f9])

        init60_ = lambda111_
        nw301_ = _dafny.Array(None, 21)
        for i0_60_ in range(nw301_.length(0)):
            nw301_[i0_60_] = init60_(i0_60_)
        d_1933_v13_ = nw301_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1933_v13_).length(0)):
            d_1935_i1_: int = guard_loop_5_
            if (True) and (((0) <= (d_1935_i1_)) and ((d_1935_i1_) < ((d_1933_v13_).length(0)))):
                (d_1933_v13_)[(d_1935_i1_)] = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt((self).f9, (self).f9) for d_1936_i3_ in range(default__.abs(297))])
        r0 = d_1920_v1_.f3
        d_1937_v14_: _dafny.MultiSet
        d_1937_v14_ = _dafny.MultiSet([default__.fm14(globalState), not(((self).f4) or (self.f3)), (d_1920_v1_.f3) and (self.f3), not ((d_1920_v1_).f4) or (p0)])
        r1 = ((d_1937_v14_)[default__.fm6((_dafny.MultiSet([(self).f9])).cardinality, globalState)] if (default__.fm6((_dafny.MultiSet([(self).f9])).cardinality, globalState)) in (d_1937_v14_) else (self).f9)
        d_1938_v15_: _dafny.Array
        nw302_ = _dafny.Array(int(0), 26)
        d_1938_v15_ = nw302_
        r2 = d_1938_v15_
        return r0, r1, r2

    @property
    def f9(self):
        return self._f9

class C5(T0, T1):
    def  __init__(self):
        self._f3: bool = False
        self._f4: bool = False
        self.f7: int = int(0)
        self._f6: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f3(self):
        return self._f3
    @f3.setter
    def f3(self, value):
        self._f3 = value
    @property
    def f4(self):
        return self._f4
    def ctor__(self, f6, f7, f3, f4):
        (self)._f6 = f6
        (self).f7 = f7
        (self).f3 = f3
        (self)._f4 = f4

    def fm0(self, globalState):
        if self.f3:
            return ((_dafny.MultiSet([(self).f4])).cardinality) * (self.f7)
        elif True:
            return 192

    def fm3(self, globalState):
        return (self.f7) * ((0) - (self.f7))

    def m0(self, p0, p1, p2, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        (self).f7 = self.f7
        d_1939_v0_: _dafny.Seq
        d_1939_v0_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([self.f7]), _dafny.MultiSet([p2])])
        d_1940_v1_: _dafny.Map
        d_1940_v1_ = _dafny.Map({d_1939_v0_: p0})
        (self).f3 = ((d_1940_v1_)[d_1939_v0_] if (d_1939_v0_) in (d_1940_v1_) else default__.fm4(self.f3, globalState))
        d_1941_v2_: D1
        d_1941_v2_ = D1_DC5((self).f4, self.f3, p2, self.f7)
        d_1942_v3_: C4
        nw303_ = C4()
        nw303_.ctor__(False, 367, self.f3, (self.f3 if self.f3 else (d_1941_v2_).cf8))
        d_1942_v3_ = nw303_
        d_1943_v4_: _dafny.Map
        d_1943_v4_ = _dafny.Map({p0: self.f3})
        if (len((d_1943_v4_) | (d_1943_v4_))) <= (((d_1942_v3_).f9 if d_1942_v3_.f8 else p2)):
            (d_1942_v3_).f8 = p0
            if d_1942_v3_.f8:
                d_1944_v5_: _dafny.Set
                d_1944_v5_ = _dafny.Set({p0})
                rhs261_ = (d_1944_v5_).ispropersubset(d_1944_v5_)
                rhs262_ = 735
                rhs263_ = (self).f4
                lhs190_ = globalState
                lhs191_ = globalState
                lhs190_.f0 = rhs261_
                r1 = rhs262_
                lhs191_.f0 = rhs263_
                (d_1942_v3_).f8 = d_1942_v3_.f8
                d_1945_v6_: _dafny.Seq
                d_1945_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yuup"))
                d_1945_v6_ = ((d_1945_v6_) + (default__.fm30(d_1941_v2_, (0) - (p2), p2, globalState))) + ((self).f6)
                d_1946_v7_: _dafny.MultiSet
                d_1946_v7_ = _dafny.MultiSet([(self).f6, (self).f6, d_1945_v6_])
                r1 = ((d_1946_v7_)[(self).f6] if ((self).f6) in (d_1946_v7_) else self.f7)
                d_1947_v8_: _dafny.MultiSet
                d_1947_v8_ = _dafny.MultiSet([p2, self.f7])
                d_1945_v6_ = ((d_1945_v6_).set(default__.safeIndex((d_1947_v8_).cardinality, len(d_1945_v6_)), _dafny.CodePoint('r'))) + ((d_1945_v6_) + ((self).f6))
            elif True:
                d_1948_v9_: _dafny.Array
                nw304_ = _dafny.Array(_dafny.Seq({}), 12)
                d_1948_v9_ = nw304_
                d_1948_v9_ = d_1948_v9_
                d_1949_v10_: _dafny.Seq
                d_1949_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbfhlbhiw"))
                d_1950_v11_: _dafny.Array
                nw305_ = _dafny.Array(None, 8)
                d_1950_v11_ = nw305_
                d_1951_v12_: _dafny.Array
                def lambda112_(d_1952_v3_):
                    def lambda113_(d_1953_i0_):
                        return _dafny.Set({(d_1952_v3_).f9})

                    return lambda113_

                init61_ = lambda112_(d_1942_v3_)
                nw306_ = _dafny.Array(None, 16)
                for i0_61_ in range(nw306_.length(0)):
                    nw306_[i0_61_] = init61_(i0_61_)
                d_1951_v12_ = nw306_
                d_1954_v13_: D15
                d_1954_v13_ = D15_DC32(d_1951_v12_)
                index317_ = default__.safeIndex(368, (d_1950_v11_).length(0))
                (d_1950_v11_)[index317_] = d_1954_v13_
                d_1955_v14_: _dafny.Array
                nw307_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
                d_1955_v14_ = nw307_
                d_1956_v15_: _dafny.Array
                nw308_ = _dafny.Array(None, 17)
                nw308_[int(0)] = d_1955_v14_
                nw308_[int(1)] = d_1955_v14_
                nw308_[int(2)] = d_1955_v14_
                nw308_[int(3)] = d_1955_v14_
                nw308_[int(4)] = d_1955_v14_
                nw308_[int(5)] = d_1955_v14_
                nw308_[int(6)] = d_1955_v14_
                nw308_[int(7)] = d_1955_v14_
                nw308_[int(8)] = d_1955_v14_
                nw308_[int(9)] = d_1955_v14_
                nw308_[int(10)] = d_1955_v14_
                nw308_[int(11)] = d_1955_v14_
                nw308_[int(12)] = d_1955_v14_
                nw308_[int(13)] = d_1955_v14_
                nw308_[int(14)] = d_1955_v14_
                nw308_[int(15)] = d_1955_v14_
                nw308_[int(16)] = d_1955_v14_
                d_1956_v15_ = nw308_
                index318_ = default__.safeIndex(122, (d_1956_v15_).length(0))
                (d_1956_v15_)[index318_] = d_1955_v14_
                d_1957_v16_: _dafny.Array
                def lambda114_(d_1958_i1_):
                    return (d_1958_i1_) - (len(_dafny.SeqWithoutIsStrInference([601])))

                init62_ = lambda114_
                nw309_ = _dafny.Array(None, 9)
                for i0_62_ in range(nw309_.length(0)):
                    nw309_[i0_62_] = init62_(i0_62_)
                d_1957_v16_ = nw309_
                pat_let_tv57_ = d_1951_v12_
                index319_ = default__.safeIndex(368, (d_1950_v11_).length(0))
                index320_ = default__.safeIndex(122, (d_1956_v15_).length(0))
                def iife165_(_pat_let48_0):
                    def iife166_(d_1959_dt__update__tmp_h0_):
                        def iife167_(_pat_let49_0):
                            def iife168_(d_1960_dt__update_hcf46_h0_):
                                return D15_DC32(d_1960_dt__update_hcf46_h0_)
                            return iife168_(_pat_let49_0)
                        return iife167_(pat_let_tv57_)
                    return iife166_(_pat_let48_0)
                rhs264_ = d_1957_v16_
                rhs265_ = ((d_1949_v10_ if self.f3 else (self).f6)) + (d_1949_v10_)
                rhs266_ = iife165_(d_1954_v13_)
                rhs267_ = d_1955_v14_
                lhs192_ = d_1950_v11_
                lhs193_ = default__.safeIndex(368, (d_1950_v11_).length(0))
                lhs194_ = d_1956_v15_
                lhs195_ = default__.safeIndex(122, (d_1956_v15_).length(0))
                r0 = rhs264_
                d_1949_v10_ = rhs265_
                lhs192_[lhs193_] = rhs266_
                lhs194_[lhs195_] = rhs267_
                d_1961_v17_: _dafny.Seq
                d_1961_v17_ = _dafny.SeqWithoutIsStrInference([(self).f4])
                d_1962_v18_: _dafny.Map
                d_1962_v18_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(d_1961_v17_)])): (d_1942_v3_).f9})
                d_1963_v19_: _dafny.MultiSet
                d_1963_v19_ = _dafny.MultiSet([d_1962_v18_, d_1962_v18_])
                d_1964_v20_: C4
                nw310_ = C4()
                nw310_.ctor__(not(default__.fm21((self).fm3(globalState), d_1963_v19_, globalState)), (0) - (self.f7), not(d_1942_v3_.f8), (self).f4)
                d_1964_v20_ = nw310_
                d_1965_v21_: _dafny.Map
                d_1965_v21_ = _dafny.Map({p0: (d_1942_v3_).f9})
                r1 = default__.safeDivisionInt((d_1942_v3_).f9, default__.safeModuloInt((d_1964_v20_).f9, ((d_1965_v21_)[d_1942_v3_.f8] if (d_1942_v3_.f8) in (d_1965_v21_) else (d_1942_v3_).f9)))
                d_1966_v22_: C3
                nw311_ = C3()
                nw311_.ctor__(True, not(d_1942_v3_.f8))
                d_1966_v22_ = nw311_
                d_1967_v23_: _dafny.Map
                d_1967_v23_ = _dafny.Map({d_1966_v22_: d_1942_v3_.f8})
                d_1968_v24_: _dafny.Map
                d_1968_v24_ = _dafny.Map({_dafny.Set({(self).f4}): d_1957_v16_})
                d_1969_v25_: _dafny.Map
                d_1969_v25_ = _dafny.Map({((d_1968_v24_)[_dafny.Set({self.f3, self.f3, d_1942_v3_.f8, not(True)})] if (_dafny.Set({self.f3, self.f3, d_1942_v3_.f8, not(True)})) in (d_1968_v24_) else d_1957_v16_): d_1964_v20_.f8})
                d_1970_v26_: _dafny.Map
                d_1970_v26_ = _dafny.Map({len(default__.fm33((d_1964_v20_).f9, d_1961_v17_, globalState)): (self).f6})
                d_1971_v27_: _dafny.Array
                nw312_ = _dafny.Array(None, 20)
                nw312_[int(0)] = not(p0)
                nw312_[int(1)] = not(self.f3)
                nw312_[int(2)] = d_1942_v3_.f8
                nw312_[int(3)] = self.f3
                nw312_[int(4)] = d_1942_v3_.f8
                nw312_[int(5)] = d_1942_v3_.f8
                nw312_[int(6)] = d_1942_v3_.f8
                nw312_[int(7)] = (509) != (232)
                nw312_[int(8)] = ((d_1967_v23_)[d_1966_v22_] if (d_1966_v22_) in (d_1967_v23_) else not(d_1942_v3_.f8))
                nw312_[int(9)] = ((d_1942_v3_).f9) <= ((d_1942_v3_).f9)
                nw312_[int(10)] = self.f3
                nw312_[int(11)] = False
                nw312_[int(12)] = ((d_1942_v3_).f9) < ((d_1942_v3_).f9)
                nw312_[int(13)] = not(d_1942_v3_.f8)
                nw312_[int(14)] = d_1942_v3_.f8
                nw312_[int(15)] = d_1964_v20_.f8
                nw312_[int(16)] = p0
                nw312_[int(17)] = (len(d_1969_v25_)) > (len(d_1970_v26_))
                nw312_[int(18)] = p0
                nw312_[int(19)] = d_1942_v3_.f8
                d_1971_v27_ = nw312_
                index321_ = default__.safeIndex(155, (d_1971_v27_).length(0))
                (d_1971_v27_)[index321_] = d_1942_v3_.f8
                index322_ = default__.safeIndex(155, (d_1971_v27_).length(0))
                def iife169_():
                    coll69_ = _dafny.Map()
                    compr_69_: int
                    for compr_69_ in (p1).keys.Elements:
                        d_1972_v28_: int = compr_69_
                        if (d_1972_v28_) in (p1):
                            coll69_[(d_1972_v28_) * ((d_1942_v3_).f9)] = self.f7
                    return _dafny.Map(coll69_)
                (d_1971_v27_)[index322_] = not(not(default__.fm21((d_1942_v3_).f9, (d_1963_v19_) | ((d_1963_v19_).set(iife169_()
                , default__.abs((d_1942_v3_).f9))), globalState)))
            d_1973_v29_: _dafny.Array
            nw313_ = _dafny.Array(int(0), 22)
            d_1973_v29_ = nw313_
            index323_ = default__.safeIndex(322, (d_1973_v29_).length(0))
            (d_1973_v29_)[index323_] = p2
            d_1974_v30_: _dafny.Seq
            d_1974_v30_ = _dafny.SeqWithoutIsStrInference([p2])
            index324_ = default__.safeIndex(322, (d_1973_v29_).length(0))
            (d_1973_v29_)[index324_] = ((D19_DC46((d_1942_v3_).f9, self.f3, (self).f4, D1_DC5(self.f3, d_1942_v3_.f8, len((self).f6), (d_1942_v3_).f9), len(d_1974_v30_))).cf78) + (((d_1942_v3_).f9) * ((d_1942_v3_).f9))
            d_1975_v31_: _dafny.Seq
            d_1975_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqudafc"))
            d_1975_v31_ = (self).f6
            d_1976_v32_: _dafny.MultiSet
            d_1976_v32_ = _dafny.MultiSet([default__.safeDivisionInt(self.f7, (d_1973_v29_)[default__.safeIndex(322, (d_1973_v29_).length(0))]), (p2) + (p2)])
            d_1976_v32_ = _dafny.MultiSet([self.f7, len(d_1974_v30_), self.f7, (d_1942_v3_).f9, (0) - (len((d_1974_v30_) + (d_1974_v30_)))])
        elif True:
            d_1977_v33_: _dafny.MultiSet
            d_1977_v33_ = _dafny.MultiSet([(d_1942_v3_).f9])
            r1 = default__.safeDivisionInt((d_1942_v3_).f9, ((d_1977_v33_)[(d_1942_v3_).f9] if ((d_1942_v3_).f9) in (d_1977_v33_) else self.f7))
            d_1978_v34_: _dafny.Set
            d_1978_v34_ = _dafny.Set({p0})
            d_1979_v35_: _dafny.Map
            d_1979_v35_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gfiwjkr")): d_1978_v34_})
            d_1980_v36_: _dafny.Map
            d_1980_v36_ = _dafny.Map({(self).f4: self.f7})
            d_1981_v37_: _dafny.Seq
            d_1981_v37_ = _dafny.SeqWithoutIsStrInference([d_1980_v36_, d_1980_v36_, _dafny.Map({False: (d_1942_v3_).f9})])
            d_1982_v38_: str
            d_1982_v38_ = _dafny.CodePoint('g')
            d_1983_v39_: _dafny.Map
            d_1983_v39_ = _dafny.Map({p2: d_1982_v38_})
            d_1979_v35_ = (d_1979_v35_).set(default__.fm30(d_1941_v2_, self.f7, len((d_1981_v37_)[default__.safeIndex(len(d_1983_v39_), len(d_1981_v37_))]), globalState), d_1978_v34_)
            (self).f7 = default__.safeDivisionInt((self.f7) * (len((self).f6)), self.f7)
            d_1984_v40_: _dafny.Array
            def lambda115_(d_1985_v33_):
                def lambda116_(d_1986_i2_):
                    return d_1985_v33_

                return lambda116_

            init63_ = lambda115_(d_1977_v33_)
            nw314_ = _dafny.Array(None, 15)
            for i0_63_ in range(nw314_.length(0)):
                nw314_[i0_63_] = init63_(i0_63_)
            d_1984_v40_ = nw314_
            d_1984_v40_ = d_1984_v40_
            d_1987_v41_: _dafny.Array
            def lambda117_(d_1988_i4_):
                return (self).f4

            init64_ = lambda117_
            nw315_ = _dafny.Array(None, 4)
            for i0_64_ in range(nw315_.length(0)):
                nw315_[i0_64_] = init64_(i0_64_)
            d_1987_v41_ = nw315_
            d_1989_v42_: D2
            d_1989_v42_ = D2_DC8(d_1987_v41_, _dafny.SeqWithoutIsStrInference([d_1982_v38_ for d_1990_i5_ in range(default__.abs(-310))]))
            d_1991_v43_: _dafny.Seq
            d_1991_v43_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_1982_v38_ for d_1992_i3_ in range(default__.abs(-888))])) + ((self).f6), (d_1989_v42_).cf14, (self).f6, ((self).f6) + ((self).f6)])
            d_1993_v44_: _dafny.Seq
            d_1993_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnqupg"))
            rhs268_ = _dafny.SeqWithoutIsStrInference([d_1993_v44_ for d_1994_i6_ in range(default__.abs(971))])
            rhs269_ = ((p0 if self.f3 else (self).f4)) and ((_dafny.MultiSet([(d_1942_v3_).f9])).ispropersubset(d_1977_v33_))
            rhs270_ = self.f7
            rhs271_ = ((d_1989_v42_).cf14) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqw"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fetrpyq"))))
            lhs196_ = globalState
            d_1991_v43_ = rhs268_
            lhs196_.f0 = rhs269_
            r1 = rhs270_
            d_1993_v44_ = rhs271_
        source24_ = d_1941_v2_
        if source24_.is_DC4:
            d_1995___mcc_h0_ = source24_.cf7
            d_1996_cf7_ = d_1995___mcc_h0_
            d_1997_v45_: _dafny.Set
            d_1997_v45_ = _dafny.Set({d_1942_v3_.f8, self.f3})
            d_1998_v46_: _dafny.Seq
            d_1998_v46_ = _dafny.SeqWithoutIsStrInference([d_1997_v45_, d_1997_v45_, d_1997_v45_])
            d_1999_v47_: str
            d_1999_v47_ = _dafny.CodePoint('p')
            d_2000_v48_: _dafny.MultiSet
            d_2000_v48_ = _dafny.MultiSet([d_1996_cf7_])
            d_2001_v49_: _dafny.Seq
            d_2001_v49_ = _dafny.SeqWithoutIsStrInference([default__.fm18(len((d_1998_v46_)[default__.safeIndex((d_1942_v3_).f9, len(d_1998_v46_))]), d_1999_v47_, (d_1942_v3_).f9, globalState), p0, d_1942_v3_.f8, (d_2000_v48_).issubset(d_2000_v48_), False])
            rhs272_ = (d_2001_v49_)[default__.safeIndex((d_1942_v3_).f9, len(d_2001_v49_))]
            rhs273_ = d_1942_v3_.f8
            lhs197_ = globalState
            lhs198_ = globalState
            lhs197_.f0 = rhs272_
            lhs198_.f0 = rhs273_
            d_2002_v50_: _dafny.Map
            d_2002_v50_ = _dafny.Map({d_1942_v3_: p0})
            d_2003_v51_: _dafny.Array
            nw316_ = _dafny.Array(None, 28)
            nw316_[int(0)] = d_1942_v3_.f8
            nw316_[int(1)] = not((self).f4)
            nw316_[int(2)] = (self).f4
            nw316_[int(3)] = True
            nw316_[int(4)] = d_1942_v3_.f8
            nw316_[int(5)] = False
            nw316_[int(6)] = self.f3
            nw316_[int(7)] = p0
            nw316_[int(8)] = self.f3
            nw316_[int(9)] = p0
            nw316_[int(10)] = self.f3
            nw316_[int(11)] = p0
            nw316_[int(12)] = d_1942_v3_.f8
            nw316_[int(13)] = ((d_2002_v50_)[d_1942_v3_] if (d_1942_v3_) in (d_2002_v50_) else d_1942_v3_.f8)
            nw316_[int(14)] = p0
            nw316_[int(15)] = d_1942_v3_.f8
            nw316_[int(16)] = self.f3
            nw316_[int(17)] = d_1942_v3_.f8
            nw316_[int(18)] = not(True)
            nw316_[int(19)] = self.f3
            nw316_[int(20)] = not(d_1942_v3_.f8)
            nw316_[int(21)] = d_1942_v3_.f8
            nw316_[int(22)] = d_1942_v3_.f8
            nw316_[int(23)] = d_1942_v3_.f8
            nw316_[int(24)] = d_1942_v3_.f8
            nw316_[int(25)] = p0
            nw316_[int(26)] = self.f3
            nw316_[int(27)] = (self).f4
            d_2003_v51_ = nw316_
            d_2004_v52_: D4
            d_2004_v52_ = D4_DC11(d_2001_v49_)
            d_2005_v53_: D19
            d_2005_v53_ = D19_DC45((d_1942_v3_).f9, d_1996_cf7_, d_1942_v3_.f8, False, (d_2004_v52_).cf18)
            d_2006_v54_: D19
            d_2006_v54_ = D19_DC47(d_2005_v53_)
            d_2007_v55_: _dafny.Map
            d_2007_v55_ = _dafny.Map({d_1996_cf7_: d_2006_v54_})
            d_2008_v56_: D18
            d_2008_v56_ = D18_DC43(d_1996_cf7_, p2, d_1942_v3_.f8)
            d_2009_v57_: _dafny.Seq
            d_2009_v57_ = _dafny.SeqWithoutIsStrInference([195])
            d_2010_v58_: D8
            d_2010_v58_ = D8_DC19(len(d_2009_v57_), p2)
            d_2011_v59_: _dafny.MultiSet
            d_2011_v59_ = _dafny.MultiSet([d_1942_v3_.f8])
            d_2012_v60_: _dafny.Map
            d_2012_v60_ = _dafny.Map({(d_1942_v3_).f9: d_2001_v49_})
            d_2013_v61_: _dafny.Map
            d_2013_v61_ = _dafny.Map({self.f7: d_2012_v60_})
            d_2014_v62_: _dafny.Map
            d_2014_v62_ = _dafny.Map({d_2000_v48_: default__.fm31(d_1942_v3_.f8, d_1996_cf7_, self.f3, d_2001_v49_, globalState)})
            d_2015_v63_: _dafny.Array
            nw317_ = _dafny.Array(None, 26)
            nw317_[int(0)] = (0) - (d_1996_cf7_)
            nw317_[int(1)] = p2
            nw317_[int(2)] = default__.safeDivisionInt((d_1942_v3_).f9, d_1996_cf7_)
            nw317_[int(3)] = len(((D2_DC8(d_2003_v51_, (self).f6)).cf14) + ((self).f6))
            nw317_[int(4)] = len((self).f6)
            nw317_[int(5)] = len((d_2007_v55_).set(len(_dafny.Map({d_2008_v56_: (self).f4})), d_2006_v54_))
            nw317_[int(6)] = default__.safeDivisionInt((d_1942_v3_).f9, (default__.fm42(d_2010_v58_, d_1942_v3_.f8, 557, globalState)).cardinality)
            nw317_[int(7)] = (d_1942_v3_).f9
            nw317_[int(8)] = ((default__.fm20((self).f4, globalState)) - (d_2011_v59_)).cardinality
            nw317_[int(9)] = (0) - (self.f7)
            nw317_[int(10)] = self.f7
            nw317_[int(11)] = 843
            nw317_[int(12)] = (0) - ((d_1942_v3_).f9)
            nw317_[int(13)] = d_1996_cf7_
            nw317_[int(14)] = (d_1942_v3_).f9
            nw317_[int(15)] = (self.f7) - (self.f7)
            nw317_[int(16)] = (d_1942_v3_).f9
            nw317_[int(17)] = default__.safeModuloInt(len(((d_2013_v61_)[p2] if (p2) in (d_2013_v61_) else d_2012_v60_)), (0) - ((d_1942_v3_).f9))
            nw317_[int(18)] = len(default__.fm34(len((self).f6), (d_1942_v3_).f9, globalState))
            nw317_[int(19)] = d_1996_cf7_
            nw317_[int(20)] = (d_1942_v3_).f9
            nw317_[int(21)] = (0) - (self.f7)
            nw317_[int(22)] = (default__.fm31((self).f4, (0) - ((_dafny.MultiSet([p2, (d_1942_v3_).f9])).cardinality), d_1942_v3_.f8, _dafny.SeqWithoutIsStrInference([True, (self).f4, self.f3, d_1942_v3_.f8]), globalState)) * ((d_1942_v3_).f9)
            nw317_[int(23)] = self.f7
            nw317_[int(24)] = self.f7
            nw317_[int(25)] = ((0) - (self.f7)) - (len(d_2014_v62_))
            d_2015_v63_ = nw317_
            d_2016_v64_: _dafny.Map
            d_2016_v64_ = _dafny.Map({d_1999_v47_: (self).f4})
            index325_ = default__.safeIndex(319, (d_2015_v63_).length(0))
            (d_2015_v63_)[index325_] = len(d_2016_v64_)
            index326_ = default__.safeIndex(319, (d_2015_v63_).length(0))
            (d_2015_v63_)[index326_] = ((len((self).f6) if self.f3 else d_1996_cf7_)) * ((d_1942_v3_).fm5(d_1942_v3_.f8, (self).f6, globalState))
            pat_let_tv58_ = d_1942_v3_
            pat_let_tv59_ = d_1942_v3_
            def iife170_(_pat_let50_0):
                def iife171_(d_2017_dt__update__tmp_h1_):
                    def iife172_(_pat_let51_0):
                        def iife173_(d_2018_dt__update_hcf30_h0_):
                            return D8_DC18(d_2018_dt__update_hcf30_h0_)
                        return iife173_(_pat_let51_0)
                    return iife172_(_dafny.Map({(pat_let_tv58_).f9: pat_let_tv59_.f8}))
                return iife171_(_pat_let50_0)
            source25_ = iife170_(D8_DC18(p1))
            if source25_.is_DC19:
                d_2019___mcc_h6_ = source25_.cf31
                d_2020___mcc_h7_ = source25_.cf32
                d_2021_cf32_ = d_2020___mcc_h7_
                d_2022_cf31_ = d_2019___mcc_h6_
                d_2023_v65_: _dafny.Set
                d_2023_v65_ = _dafny.Set({(self).f6})
                d_2024_v66_: _dafny.MultiSet
                d_2024_v66_ = _dafny.MultiSet([(self).f6, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwncwwteh"))])
                d_2025_v68_: _dafny.Map
                def iife174_():
                    coll70_ = _dafny.Set()
                    compr_70_: _dafny.Seq
                    for compr_70_ in (d_2024_v66_).Elements:
                        d_2026_v67_: _dafny.Seq = compr_70_
                        if (d_2026_v67_) in (d_2024_v66_):
                            coll70_ = coll70_.union(_dafny.Set([d_2026_v67_]))
                    return _dafny.Set(coll70_)
                d_2025_v68_ = _dafny.Map({d_2021_cf32_: iife174_()
                })
                d_2027_v70_: _dafny.Seq
                def iife175_():
                    coll71_ = _dafny.Set()
                    compr_71_: _dafny.Seq
                    for compr_71_ in (d_2023_v65_).Elements:
                        d_2028_v69_: _dafny.Seq = compr_71_
                        if (d_2028_v69_) in (d_2023_v65_):
                            coll71_ = coll71_.union(_dafny.Set([d_2028_v69_]))
                    return _dafny.Set(coll71_)
                d_2027_v70_ = _dafny.SeqWithoutIsStrInference([d_2023_v65_, _dafny.Set({(self).f6}), ((d_2025_v68_)[d_2022_cf31_] if (d_2022_cf31_) in (d_2025_v68_) else iife175_()
                ), _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uurou"))})])
                d_2029_v71_: _dafny.Map
                d_2029_v71_ = _dafny.Map({d_1942_v3_.f8: d_1997_v45_})
                d_2030_v72_: _dafny.Map
                d_2030_v72_ = _dafny.Map({(d_1942_v3_).f9: (d_1942_v3_).f9})
                d_2031_v73_: _dafny.Map
                d_2031_v73_ = _dafny.Map({len(((d_2029_v71_)[self.f3] if (self.f3) in (d_2029_v71_) else d_1997_v45_)): d_2030_v72_})
                rhs274_ = ((d_2023_v65_) - ((d_2027_v70_)[default__.safeIndex(len(d_2031_v73_), len(d_2027_v70_))])).isdisjoint(d_2023_v65_)
                rhs275_ = d_1941_v2_
                rhs276_ = d_1996_cf7_
                lhs199_ = globalState
                lhs199_.f0 = rhs274_
                d_1941_v2_ = rhs275_
                r1 = rhs276_
                d_2001_v49_ = default__.fm34((d_1942_v3_).f9, (d_2022_cf31_ if d_1942_v3_.f8 else d_1996_cf7_), globalState)
                (d_1942_v3_).f8 = not(p0)
                (globalState).f0 = d_1942_v3_.f8
            elif source25_.is_DC20:
                d_2032___mcc_h8_ = source25_.cf33
                d_2033___mcc_h9_ = source25_.cf34
                d_2034_cf34_ = d_2033___mcc_h9_
                d_2035_cf33_ = d_2032___mcc_h8_
                (globalState).f0 = d_1942_v3_.f8
                d_2003_v51_ = (d_2003_v51_ if p0 else (d_2003_v51_ if not(False) else d_2003_v51_))
                d_2036_v74_: _dafny.Map
                d_2036_v74_ = _dafny.Map({self.f7: (d_1942_v3_).f9})
                d_2037_v75_: _dafny.Set
                d_2037_v75_ = _dafny.Set({(d_1942_v3_).f9, (459 if (self).f4 else p2), ((0) - ((d_1942_v3_).f9)) * (self.f7), len(d_2036_v74_)})
                d_2037_v75_ = d_2037_v75_
                d_2038_v76_: int
                d_2039_v77_: bool
                out38_: int
                out39_: bool
                out38_, out39_ = (d_1942_v3_).m5(d_1942_v3_.f8, d_1942_v3_.f8, globalState)
                d_2038_v76_ = out38_
                d_2039_v77_ = out39_
            elif True:
                d_2040___mcc_h10_ = source25_.cf30
                d_2041_cf30_ = d_2040___mcc_h10_
                rhs277_ = (self).f4
                rhs278_ = (((d_1942_v3_).f9 if d_1942_v3_.f8 else 294)) == (p2)
                lhs200_ = globalState
                lhs201_ = self
                lhs200_.f0 = rhs277_
                lhs201_.f3 = rhs278_
                (globalState).f0 = ((default__.fm30(d_1941_v2_, len((self).f6), (d_2015_v63_)[default__.safeIndex(319, (d_2015_v63_).length(0))], globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnoyfw")))) < ((self).f6)
                d_1943_v4_ = (d_1943_v4_).set(((self).f6) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_2042_i7_ in range(default__.abs(882))])), not(not(not(self.f3))))
                d_2043_v78_: _dafny.Map
                d_2043_v78_ = _dafny.Map({d_2003_v51_: 169})
                d_2043_v78_ = (d_2043_v78_).set(d_2003_v51_, (d_1942_v3_).f9)
            d_2044_v79_: C0
            nw318_ = C0()
            nw318_.ctor__()
            d_2044_v79_ = nw318_
        elif source24_.is_DC5:
            d_2045___mcc_h1_ = source24_.cf8
            d_2046___mcc_h2_ = source24_.cf9
            d_2047___mcc_h3_ = source24_.cf10
            d_2048___mcc_h4_ = source24_.cf11
            d_2049_cf11_ = d_2048___mcc_h4_
            d_2050_cf10_ = d_2047___mcc_h3_
            d_2051_cf9_ = d_2046___mcc_h2_
            d_2052_cf8_ = d_2045___mcc_h1_
            if False:
                d_2053_v80_: _dafny.Array
                def lambda118_(d_2054_i8_):
                    return _dafny.Map({_dafny.CodePoint('e'): len(_dafny.SeqWithoutIsStrInference([-863, len((self).f6)]))})

                init65_ = lambda118_
                nw319_ = _dafny.Array(None, 23)
                for i0_65_ in range(nw319_.length(0)):
                    nw319_[i0_65_] = init65_(i0_65_)
                d_2053_v80_ = nw319_
                d_2055_v81_: str
                d_2055_v81_ = _dafny.CodePoint('b')
                d_2056_v82_: _dafny.Map
                d_2056_v82_ = _dafny.Map({d_2055_v81_: (d_1942_v3_).f9})
                index327_ = default__.safeIndex(170, (d_2053_v80_).length(0))
                (d_2053_v80_)[index327_] = d_2056_v82_
                index328_ = default__.safeIndex(170, (d_2053_v80_).length(0))
                (d_2053_v80_)[index328_] = d_2056_v82_
                d_2057_v83_: _dafny.MultiSet
                d_2057_v83_ = _dafny.MultiSet([(self).f4, d_2052_cf8_])
                (globalState).f0 = (d_2057_v83_).issubset(d_2057_v83_)
                d_2050_cf10_ = ((d_1942_v3_).f9) * (d_2049_cf11_)
                d_2058_v84_: _dafny.Array
                nw320_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_2058_v84_ = nw320_
                d_2059_v85_: _dafny.Array
                def lambda119_(d_2060_v3_):
                    def lambda120_(d_2061_i9_):
                        return d_2060_v3_.f8

                    return lambda120_

                init66_ = lambda119_(d_1942_v3_)
                nw321_ = _dafny.Array(None, 7)
                for i0_66_ in range(nw321_.length(0)):
                    nw321_[i0_66_] = init66_(i0_66_)
                d_2059_v85_ = nw321_
                index329_ = default__.safeIndex(213, (d_2058_v84_).length(0))
                (d_2058_v84_)[index329_] = d_2059_v85_
                index330_ = default__.safeIndex(213, (d_2058_v84_).length(0))
                (d_2058_v84_)[index330_] = d_2059_v85_
                d_2062_v86_: C3
                nw322_ = C3()
                nw322_.ctor__((self.f3) or (d_2051_cf9_), (self).f4)
                d_2062_v86_ = nw322_
            elif True:
                d_2063_v87_: _dafny.Array
                nw323_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_2063_v87_ = nw323_
                d_2064_v88_: _dafny.Array
                nw324_ = _dafny.Array(_dafny.MultiSet({}), 24)
                d_2064_v88_ = nw324_
                index331_ = default__.safeIndex(339, (d_2063_v87_).length(0))
                (d_2063_v87_)[index331_] = d_2064_v88_
                d_2065_v89_: _dafny.Array
                def lambda121_(d_2066_cf8_):
                    def lambda122_(d_2067_i10_):
                        return d_2066_cf8_

                    return lambda122_

                init67_ = lambda121_(d_2052_cf8_)
                nw325_ = _dafny.Array(None, 7)
                for i0_67_ in range(nw325_.length(0)):
                    nw325_[i0_67_] = init67_(i0_67_)
                d_2065_v89_ = nw325_
                d_2068_v90_: _dafny.MultiSet
                d_2068_v90_ = _dafny.MultiSet([not(d_1942_v3_.f8), True])
                index332_ = default__.safeIndex(121, (d_2065_v89_).length(0))
                (d_2065_v89_)[index332_] = (d_2068_v90_).issubset(d_2068_v90_)
                d_2069_v91_: _dafny.Seq
                d_2069_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfcgkb"))
                index333_ = default__.safeIndex(339, (d_2063_v87_).length(0))
                index334_ = default__.safeIndex(121, (d_2065_v89_).length(0))
                rhs279_ = (d_2064_v88_ if d_1942_v3_.f8 else d_2064_v88_)
                rhs280_ = default__.safeDivisionInt((d_2050_cf10_) * (p2), 276)
                rhs281_ = not((self).f4)
                rhs282_ = (self).f6
                rhs283_ = d_2065_v89_
                lhs202_ = d_2063_v87_
                lhs203_ = default__.safeIndex(339, (d_2063_v87_).length(0))
                lhs204_ = d_2065_v89_
                lhs205_ = default__.safeIndex(121, (d_2065_v89_).length(0))
                lhs202_[lhs203_] = rhs279_
                d_2049_cf11_ = rhs280_
                lhs204_[lhs205_] = rhs281_
                d_2069_v91_ = rhs282_
                d_2065_v89_ = rhs283_
                d_2070_v92_: _dafny.MultiSet
                d_2070_v92_ = _dafny.MultiSet([(d_1942_v3_).f9, d_2050_cf10_])
                (self).f3 = (d_2049_cf11_) >= (((d_2070_v92_)[(d_1942_v3_).f9] if ((d_1942_v3_).f9) in (d_2070_v92_) else d_2049_cf11_))
                d_2071_v93_: _dafny.Seq
                d_2071_v93_ = _dafny.SeqWithoutIsStrInference([(d_1943_v4_).set(self.f3, d_2052_cf8_), d_1943_v4_])
                r1 = len(d_2071_v93_)
                d_2072_v94_: _dafny.Map
                d_2072_v94_ = _dafny.Map({d_2050_cf10_: (0) - (self.f7)})
                d_2072_v94_ = (d_2072_v94_).set(p2, d_2049_cf11_)
                d_2073_v95_: bool
                d_2074_v96_: bool
                d_2075_v97_: int
                d_2076_v98_: int
                out40_: bool
                out41_: bool
                out42_: int
                out43_: int
                out40_, out41_, out42_, out43_ = (self).m3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkowqfgre")), (self.f7) >= ((d_1942_v3_).f9), (d_2051_cf9_) and (self.f3), (d_1942_v3_).f9, globalState)
                d_2073_v95_ = out40_
                d_2074_v96_ = out41_
                d_2075_v97_ = out42_
                d_2076_v98_ = out43_
            d_2077_v99_: _dafny.Array
            nw326_ = _dafny.Array(_dafny.MultiSet({}), 12)
            d_2077_v99_ = nw326_
            index335_ = default__.safeIndex(605, (d_2077_v99_).length(0))
            (d_2077_v99_)[index335_] = _dafny.MultiSet([d_2051_cf9_])
            d_2078_v100_: _dafny.MultiSet
            d_2078_v100_ = _dafny.MultiSet([d_2052_cf8_, d_2052_cf8_])
            index336_ = default__.safeIndex(605, (d_2077_v99_).length(0))
            (d_2077_v99_)[index336_] = d_2078_v100_
            d_2079_v101_: bool
            d_2080_v102_: int
            d_2081_v103_: _dafny.Array
            out44_: bool
            out45_: int
            out46_: _dafny.Array
            out44_, out45_, out46_ = (d_1942_v3_).m6(d_1942_v3_.f8, globalState)
            d_2079_v101_ = out44_
            d_2080_v102_ = out45_
            d_2081_v103_ = out46_
            r1 = p2
        elif source24_.is_DC6:
            d_2082_v104_: _dafny.Seq
            d_2082_v104_ = _dafny.SeqWithoutIsStrInference([False, not(self.f3), d_1942_v3_.f8, (self).f4])
            d_2083_v105_: D11
            d_2083_v105_ = D11_DC24(_dafny.Map({len(d_2082_v104_): 180}))
            d_2084_v106_: _dafny.Map
            d_2084_v106_ = _dafny.Map({self.f7: (d_1942_v3_).f9})
            d_2083_v105_ = (d_2083_v105_ if d_1942_v3_.f8 else D11_DC24(d_2084_v106_))
            d_2085_v107_: _dafny.Set
            d_2085_v107_ = _dafny.Set({(self).f4})
            (d_1942_v3_).f8 = (d_2085_v107_).ispropersubset(_dafny.Set({(d_2082_v104_)[default__.safeIndex((d_1942_v3_).f9, len(d_2082_v104_))]}))
            if ((0) - (p2)) in (default__.fm15((self).f4, self.f7, p2, default__.fm4(False, globalState), globalState)):
                d_2086_v108_: _dafny.Map
                d_2086_v108_ = _dafny.Map({(d_1942_v3_).f9: default__.fm30(d_1941_v2_, (d_1942_v3_).f9, (d_1942_v3_).f9, globalState)})
                d_2087_v109_: _dafny.Seq
                d_2087_v109_ = _dafny.SeqWithoutIsStrInference([(d_1942_v3_).f9])
                d_2086_v108_ = (d_2086_v108_).set((0) - (((d_1942_v3_).f9) + ((d_2087_v109_)[default__.safeIndex(p2, len(d_2087_v109_))])), ((self).f6) + ((self).f6))
                d_2088_v110_: _dafny.Set
                d_2088_v110_ = _dafny.Set({self.f7, self.f7, (d_1942_v3_).f9})
                d_2089_v111_: _dafny.Array
                nw327_ = _dafny.Array(None, 1)
                nw327_[int(0)] = len(d_2088_v110_)
                d_2089_v111_ = nw327_
                index337_ = default__.safeIndex(352, (d_2089_v111_).length(0))
                (d_2089_v111_)[index337_] = len((self).f6)
                index338_ = default__.safeIndex(363, (d_2089_v111_).length(0))
                (d_2089_v111_)[index338_] = (d_1942_v3_).f9
                index339_ = default__.safeIndex(352, (d_2089_v111_).length(0))
                index340_ = default__.safeIndex(363, (d_2089_v111_).length(0))
                rhs284_ = (0) - (p2)
                rhs285_ = True
                rhs286_ = p2
                lhs206_ = d_2089_v111_
                lhs207_ = default__.safeIndex(352, (d_2089_v111_).length(0))
                lhs208_ = d_1942_v3_
                lhs209_ = d_2089_v111_
                lhs210_ = default__.safeIndex(363, (d_2089_v111_).length(0))
                lhs206_[lhs207_] = rhs284_
                lhs208_.f8 = rhs285_
                lhs209_[lhs210_] = rhs286_
                pat_let_tv60_ = d_1942_v3_
                pat_let_tv61_ = d_2084_v106_
                pat_let_tv62_ = globalState
                pat_let_tv63_ = d_1942_v3_
                d_2090_v112_: D19
                def iife176_(_pat_let52_0):
                    def iife177_(d_2091_dt__update__tmp_h2_):
                        def iife178_(_pat_let53_0):
                            def iife179_(d_2092_dt__update_hcf8_h0_):
                                def iife180_(_pat_let54_0):
                                    def iife181_(d_2093_dt__update_hcf9_h0_):
                                        def iife182_(_pat_let55_0):
                                            def iife183_(d_2094_dt__update_hcf10_h0_):
                                                return D1_DC5(d_2092_dt__update_hcf8_h0_, d_2093_dt__update_hcf9_h0_, d_2094_dt__update_hcf10_h0_, (d_2091_dt__update__tmp_h2_).cf11)
                                            return iife183_(_pat_let55_0)
                                        return iife182_((pat_let_tv63_).f9)
                                    return iife181_(_pat_let54_0)
                                return iife180_(not(default__.fm21((pat_let_tv60_).f9, _dafny.MultiSet([pat_let_tv61_]), pat_let_tv62_)))
                            return iife179_(_pat_let53_0)
                        return iife178_((self).f4)
                    return iife177_(_pat_let52_0)
                d_2090_v112_ = D19_DC46(((d_1942_v3_).f9) * (626), p0, d_1942_v3_.f8, iife176_(d_1941_v2_), (d_2089_v111_)[default__.safeIndex(352, (d_2089_v111_).length(0))])
                pat_let_tv64_ = d_1941_v2_
                pat_let_tv65_ = d_1942_v3_
                def iife184_(_pat_let56_0):
                    def iife185_(d_2095_dt__update__tmp_h3_):
                        def iife186_(_pat_let57_0):
                            def iife187_(d_2096_dt__update_hcf75_h0_):
                                def iife188_(_pat_let58_0):
                                    def iife189_(d_2097_dt__update_hcf77_h0_):
                                        def iife190_(_pat_let59_0):
                                            def iife191_(d_2098_dt__update_hcf74_h0_):
                                                return D19_DC46(d_2098_dt__update_hcf74_h0_, d_2096_dt__update_hcf75_h0_, (d_2095_dt__update__tmp_h3_).cf76, d_2097_dt__update_hcf77_h0_, (d_2095_dt__update__tmp_h3_).cf78)
                                            return iife191_(_pat_let59_0)
                                        return iife190_((pat_let_tv65_).f9)
                                    return iife189_(_pat_let58_0)
                                return iife188_(pat_let_tv64_)
                            return iife187_(_pat_let57_0)
                        return iife186_((self).f4)
                    return iife185_(_pat_let56_0)
                d_2090_v112_ = iife184_(D19_DC46(p2, d_1942_v3_.f8, p0, d_1941_v2_, (0) - ((d_2089_v111_)[default__.safeIndex(352, (d_2089_v111_).length(0))])))
                index341_ = default__.safeIndex(352, (d_2089_v111_).length(0))
                (d_2089_v111_)[index341_] = 932
                (d_1942_v3_).f8 = (self).f4
            elif True:
                d_2099_v113_: _dafny.Array
                nw328_ = _dafny.Array(False, 26)
                d_2099_v113_ = nw328_
                index342_ = default__.safeIndex(269, (d_2099_v113_).length(0))
                (d_2099_v113_)[index342_] = (self.f3 if not((self).f4) else d_1942_v3_.f8)
                index343_ = default__.safeIndex(269, (d_2099_v113_).length(0))
                (d_2099_v113_)[index343_] = d_1942_v3_.f8
                d_2100_v114_: _dafny.MultiSet
                d_2100_v114_ = _dafny.MultiSet([self.f7, p2])
                d_2101_v115_: C2
                nw329_ = C2()
                nw329_.ctor__((d_2100_v114_).issubset(d_2100_v114_), not(default__.fm6((d_1942_v3_).f9, globalState)), (d_2099_v113_)[default__.safeIndex(269, (d_2099_v113_).length(0))], self.f3)
                d_2101_v115_ = nw329_
                d_2102_v116_: _dafny.Map
                d_2102_v116_ = _dafny.Map({d_1942_v3_.f8: len(d_2084_v106_)})
                (globalState).f0 = (((d_2102_v116_)[not(p0)] if (not(p0)) in (d_2102_v116_) else self.f7)) <= (p2)
                d_2103_v117_: _dafny.Seq
                d_2103_v117_ = _dafny.SeqWithoutIsStrInference([p2, (0) - (self.f7), p2])
                d_2104_v118_: D8
                d_2104_v118_ = D8_DC19(len((self).f6), (d_2101_v115_).fm16((d_1942_v3_).f9, (d_2082_v104_)[default__.safeIndex((d_2103_v117_)[default__.safeIndex(self.f7, len(d_2103_v117_))], len(d_2082_v104_))], self.f7, p1, globalState))
                d_2105_v119_: _dafny.Map
                d_2105_v119_ = _dafny.Map({d_2104_v118_: (d_1942_v3_).f9})
                d_2105_v119_ = (d_2105_v119_).set(D8_DC19(self.f7, self.f7), len((self).f6))
                (d_1942_v3_).f8 = (d_2085_v107_).isdisjoint(d_2085_v107_)
            r1 = self.f7
        elif True:
            d_2106___mcc_h5_ = source24_.cf6
            d_2107_cf6_ = d_2106___mcc_h5_
            d_2108_v120_: _dafny.Map
            d_2108_v120_ = _dafny.Map({(self).f6: self.f7})
            d_2109_v121_: str
            d_2109_v121_ = _dafny.CodePoint('y')
            d_2108_v120_ = (d_2108_v120_).set((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2110_i11_ in range(default__.abs(-557))])).set(default__.safeIndex((d_1942_v3_).f9, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2111_i11_ in range(default__.abs(-557))]))), d_2109_v121_), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))) - (-532))
            (self).f7 = self.f7
            d_2112_v122_: _dafny.Set
            d_2112_v122_ = _dafny.Set({d_1942_v3_.f8})
            (self).f7 = len(d_2112_v122_)
            d_2113_v123_: _dafny.Seq
            d_2113_v123_ = _dafny.SeqWithoutIsStrInference([236])
            d_2114_v124_: _dafny.Map
            d_2114_v124_ = _dafny.Map({d_2113_v123_: d_2109_v121_})
            d_2115_v125_: D17
            d_2115_v125_ = D17_DC37((d_2114_v124_) | (d_2114_v124_))
            source26_ = d_2115_v125_
            if source26_.is_DC38:
                d_2116___mcc_h11_ = source26_.cf56
                d_2117_cf56_ = d_2116___mcc_h11_
                d_2118_v126_: _dafny.MultiSet
                d_2118_v126_ = _dafny.MultiSet([(self).f4])
                d_2119_v127_: _dafny.Seq
                d_2119_v127_ = _dafny.SeqWithoutIsStrInference([d_2107_cf6_])
                d_2120_v128_: D17
                d_2120_v128_ = D17_DC38(d_2117_cf56_)
                d_2121_v129_: _dafny.Array
                nw330_ = _dafny.Array(None, 18)
                nw330_[int(0)] = (d_1942_v3_).f9
                nw330_[int(1)] = self.f7
                nw330_[int(2)] = (d_1942_v3_).f9
                nw330_[int(3)] = (d_2113_v123_)[default__.safeIndex(((d_2118_v126_)[(self).f4] if ((self).f4) in (d_2118_v126_) else p2), len(d_2113_v123_))]
                nw330_[int(4)] = self.f7
                nw330_[int(5)] = (d_1942_v3_).f9
                nw330_[int(6)] = len((d_2119_v127_ if d_1942_v3_.f8 else (d_2119_v127_).set(default__.safeIndex((d_1942_v3_).f9, len(d_2119_v127_)), d_2107_cf6_)))
                nw330_[int(7)] = (d_1942_v3_).f9
                nw330_[int(8)] = ((d_1942_v3_).f9) - ((d_1942_v3_).f9)
                nw330_[int(9)] = (0) - (((d_1942_v3_).f9) * ((d_1942_v3_).f9))
                nw330_[int(10)] = (p2 if d_1942_v3_.f8 else len((d_2120_v128_).cf56))
                nw330_[int(11)] = (d_1942_v3_).f9
                nw330_[int(12)] = (d_1942_v3_).f9
                nw330_[int(13)] = (d_1942_v3_).f9
                nw330_[int(14)] = (d_1942_v3_).f9
                nw330_[int(15)] = (len(d_2113_v123_) if not(d_1942_v3_.f8) else (d_1942_v3_).f9)
                nw330_[int(16)] = (d_2113_v123_)[default__.safeIndex(len(d_2113_v123_), len(d_2113_v123_))]
                nw330_[int(17)] = self.f7
                d_2121_v129_ = nw330_
                index344_ = default__.safeIndex(790, (d_2121_v129_).length(0))
                (d_2121_v129_)[index344_] = (p2) + (p2)
                d_2122_v130_: _dafny.Map
                d_2122_v130_ = _dafny.Map({_dafny.Set({(self).f4}): len((self).f6)})
                index345_ = default__.safeIndex(790, (d_2121_v129_).length(0))
                (d_2121_v129_)[index345_] = len((d_2122_v130_) | (d_2122_v130_))
                index346_ = default__.safeIndex(790, (d_2121_v129_).length(0))
                (d_2121_v129_)[index346_] = default__.safeModuloInt((self.f7) + (p2), (d_1942_v3_).f9)
                (d_1942_v3_).f8 = p0
                d_2123_v131_: D1
                d_2123_v131_ = D1_DC4((d_1942_v3_).f9)
                d_2124_v132_: D3
                def iife192_(_pat_let60_0):
                    def iife193_(d_2125_dt__update__tmp_h4_):
                        def iife194_(_pat_let61_0):
                            def iife195_(d_2126_dt__update_hcf7_h0_):
                                return D1_DC4(d_2126_dt__update_hcf7_h0_)
                            return iife195_(_pat_let61_0)
                        return iife194_(self.f7)
                    return iife193_(_pat_let60_0)
                d_2124_v132_ = D3_DC10((d_2118_v126_) - ((d_2118_v126_).set(not((self).f4), default__.abs((d_1942_v3_).f9))), iife192_(d_2123_v131_))
                d_2127_v133_: _dafny.Seq
                d_2127_v133_ = _dafny.SeqWithoutIsStrInference([(self).f4, False, d_1942_v3_.f8])
                d_2128_v134_: _dafny.MultiSet
                d_2128_v134_ = _dafny.MultiSet([len(d_2127_v133_), (d_1942_v3_).f9])
                d_2124_v132_ = default__.fm39(d_1942_v3_.f8, (d_1942_v3_).f9, ((_dafny.MultiSet([self.f7, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjili"))), (d_1942_v3_).f9])).intersection(d_2128_v134_)).cardinality, self.f7, globalState)
            elif source26_.is_DC39:
                d_2129___mcc_h12_ = source26_.cf57
                d_2130___mcc_h13_ = source26_.cf58
                d_2131___mcc_h14_ = source26_.cf59
                d_2132___mcc_h15_ = source26_.cf60
                d_2133___mcc_h16_ = source26_.cf61
                d_2134_cf61_ = d_2133___mcc_h16_
                d_2135_cf60_ = d_2132___mcc_h15_
                d_2136_cf59_ = d_2131___mcc_h14_
                d_2137_cf58_ = d_2130___mcc_h13_
                d_2138_cf57_ = d_2129___mcc_h12_
                (self).f7 = (d_1942_v3_).f9
                d_2134_cf61_ = len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: (d_1942_v3_).f9}) for d_2139_i12_ in range(default__.abs(510))]))
                index347_ = default__.safeIndex(817, (d_2107_cf6_).length(0))
                (d_2107_cf6_)[index347_] = not(d_2137_cf58_)
                index348_ = default__.safeIndex(817, (d_2107_cf6_).length(0))
                (d_2107_cf6_)[index348_] = default__.fm18(p2, d_2109_v121_, default__.safeDivisionInt(d_2134_cf61_, self.f7), globalState)
                d_2140_v135_: _dafny.Array
                nw331_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_2140_v135_ = nw331_
                d_2141_v136_: _dafny.Map
                d_2141_v136_ = _dafny.Map({d_2140_v135_: self.f7})
                d_2141_v136_ = (d_2141_v136_).set(d_2140_v135_, (len(p1)) - (len((self).f6)))
            elif source26_.is_DC37:
                d_2142___mcc_h17_ = source26_.cf55
                d_2143_cf55_ = d_2142___mcc_h17_
                d_2144_v137_: _dafny.Seq
                d_2144_v137_ = _dafny.SeqWithoutIsStrInference([d_1942_v3_.f8, d_1942_v3_.f8, p0, d_1942_v3_.f8, False])
                d_2145_v138_: _dafny.Map
                d_2145_v138_ = _dafny.Map({d_1942_v3_.f8: d_2109_v121_})
                d_2146_v139_: _dafny.Map
                d_2146_v139_ = _dafny.Map({(d_2144_v137_)[default__.safeIndex((0) - (len((d_2145_v138_).set(d_1942_v3_.f8, d_2109_v121_))), len(d_2144_v137_))]: p2})
                d_2146_v139_ = d_2146_v139_
                (self).f7 = default__.safeDivisionInt((len(_dafny.SeqWithoutIsStrInference([d_1942_v3_.f8, d_1942_v3_.f8]))) + (self.f7), (d_1942_v3_).f9)
                d_2147_v140_: _dafny.Array
                def lambda123_(d_2148_i13_):
                    return (d_2148_i13_) + (self.f7)

                init68_ = lambda123_
                nw332_ = _dafny.Array(None, 6)
                for i0_68_ in range(nw332_.length(0)):
                    nw332_[i0_68_] = init68_(i0_68_)
                d_2147_v140_ = nw332_
                d_2149_v141_: _dafny.Seq
                d_2149_v141_ = _dafny.SeqWithoutIsStrInference([d_2147_v140_])
                d_2150_v142_: _dafny.Array
                nw333_ = _dafny.Array(None, 7)
                nw333_[int(0)] = (d_2149_v141_) + (d_2149_v141_)
                nw333_[int(1)] = d_2149_v141_
                nw333_[int(2)] = d_2149_v141_
                nw333_[int(3)] = (d_2149_v141_) + ((d_2149_v141_).set(default__.safeIndex(p2, len(d_2149_v141_)), d_2147_v140_))
                nw333_[int(4)] = d_2149_v141_
                nw333_[int(5)] = (d_2149_v141_) + (_dafny.SeqWithoutIsStrInference([d_2147_v140_, d_2147_v140_]))
                nw333_[int(6)] = _dafny.SeqWithoutIsStrInference([d_2147_v140_, d_2147_v140_])
                d_2150_v142_ = nw333_
                index349_ = default__.safeIndex(616, (d_2150_v142_).length(0))
                (d_2150_v142_)[index349_] = d_2149_v141_
                index350_ = default__.safeIndex(616, (d_2150_v142_).length(0))
                (d_2150_v142_)[index350_] = (d_2149_v141_) + ((d_2149_v141_) + (d_2149_v141_))
                d_2151_v143_: _dafny.Array
                nw334_ = _dafny.Array(_dafny.Map({}), 20)
                d_2151_v143_ = nw334_
                d_2152_v144_: _dafny.Map
                d_2152_v144_ = _dafny.Map({(d_1942_v3_).f9: p2})
                index351_ = default__.safeIndex(107, (d_2151_v143_).length(0))
                (d_2151_v143_)[index351_] = d_2152_v144_
                index352_ = default__.safeIndex(107, (d_2151_v143_).length(0))
                rhs287_ = self.f7
                rhs288_ = (_dafny.Map({len((self).f6): len((self).f6)})) | (d_2152_v144_)
                lhs211_ = self
                lhs212_ = d_2151_v143_
                lhs213_ = default__.safeIndex(107, (d_2151_v143_).length(0))
                lhs211_.f7 = rhs287_
                lhs212_[lhs213_] = rhs288_
            elif True:
                d_2153___mcc_h18_ = source26_.cf62
                d_2154_cf62_ = d_2153___mcc_h18_
                d_2155_v145_: C3
                nw335_ = C3()
                nw335_.ctor__(True, (self).f4)
                d_2155_v145_ = nw335_
                index353_ = default__.safeIndex(844, (d_2107_cf6_).length(0))
                (d_2107_cf6_)[index353_] = p0
                index354_ = default__.safeIndex(844, (d_2107_cf6_).length(0))
                (d_2107_cf6_)[index354_] = (self).f4
                (self).f7 = 282
                r1 = (((self).fm0(globalState)) * ((d_1942_v3_).f9)) + (default__.safeModuloInt(len((self).f6), (d_1942_v3_).f9))
        (globalState).f0 = ((default__.fm24(globalState)).cardinality) > ((d_1942_v3_).f9)
        d_2156_v146_: _dafny.Map
        d_2156_v146_ = _dafny.Map({(self).f6: self.f7})
        d_2157_v147_: _dafny.Seq
        d_2157_v147_ = _dafny.SeqWithoutIsStrInference([False, d_1942_v3_.f8, d_1942_v3_.f8, p0])
        d_2158_v148_: _dafny.Set
        d_2158_v148_ = _dafny.Set({p2})
        d_2159_v149_: _dafny.Map
        d_2159_v149_ = _dafny.Map({False: len(d_2158_v148_)})
        d_2160_v150_: _dafny.Array
        nw336_ = _dafny.Array(None, 3)
        nw336_[int(0)] = d_1942_v3_.f8
        nw336_[int(1)] = not(p0)
        nw336_[int(2)] = p0
        d_2160_v150_ = nw336_
        d_2161_v151_: D1
        d_2161_v151_ = D1_DC3(d_2160_v150_)
        d_2162_v152_: _dafny.Set
        d_2162_v152_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2163_i14_ in range(default__.abs(-315))])})
        d_2164_v153_: _dafny.Set
        d_2164_v153_ = _dafny.Set({self.f3})
        nw337_ = _dafny.Array(None, 27)
        nw337_[int(0)] = self.f7
        nw337_[int(1)] = (0) - ((d_1942_v3_).f9)
        nw337_[int(2)] = ((d_2156_v146_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ochh"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ochh"))) in (d_2156_v146_) else len(d_2157_v147_))
        nw337_[int(3)] = (d_1942_v3_).f9
        nw337_[int(4)] = self.f7
        nw337_[int(5)] = (d_1942_v3_).f9
        nw337_[int(6)] = (d_1942_v3_).f9
        nw337_[int(7)] = len(p1)
        nw337_[int(8)] = ((d_2159_v149_)[p0] if (p0) in (d_2159_v149_) else p2)
        nw337_[int(9)] = (d_1942_v3_).f9
        nw337_[int(10)] = len((_dafny.Map({(d_1942_v3_).f9: d_2160_v150_})) | (_dafny.Map({p2: (d_2161_v151_).cf6})))
        nw337_[int(11)] = default__.safeDivisionInt(p2, self.f7)
        nw337_[int(12)] = len((self).f6)
        nw337_[int(13)] = len(d_2162_v152_)
        nw337_[int(14)] = p2
        nw337_[int(15)] = self.f7
        nw337_[int(16)] = (len(d_1943_v4_)) + (self.f7)
        nw337_[int(17)] = self.f7
        nw337_[int(18)] = self.f7
        nw337_[int(19)] = (p2 if p0 else self.f7)
        nw337_[int(20)] = (d_1942_v3_).f9
        nw337_[int(21)] = (p2) - (p2)
        nw337_[int(22)] = p2
        nw337_[int(23)] = 714
        nw337_[int(24)] = default__.safeModuloInt((d_1942_v3_).f9, ((d_2159_v149_)[p0] if (p0) in (d_2159_v149_) else self.f7))
        nw337_[int(25)] = p2
        nw337_[int(26)] = len(_dafny.Set({(d_1942_v3_).f9, p2, len(d_2164_v153_)}))
        r0 = nw337_
        d_2165_v154_: _dafny.MultiSet
        d_2165_v154_ = _dafny.MultiSet([not(not (d_1942_v3_.f8) or ((self).f4))])
        d_2166_v155_: str
        d_2166_v155_ = _dafny.CodePoint('m')
        d_2167_v156_: C2
        nw338_ = C2()
        nw338_.ctor__(self.f3, p0, (self).f4, self.f3)
        d_2167_v156_ = nw338_
        d_2168_v157_: _dafny.MultiSet
        d_2168_v157_ = _dafny.MultiSet([d_2167_v156_])
        d_2169_v158_: _dafny.Seq
        d_2169_v158_ = _dafny.SeqWithoutIsStrInference([(d_2168_v157_).cardinality])
        r1 = ((d_2165_v154_)[default__.fm18((0) - (self.f7), d_2166_v155_, p2, globalState)] if (default__.fm18((0) - (self.f7), d_2166_v155_, p2, globalState)) in (d_2165_v154_) else len(d_2169_v158_))
        return r0, r1

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        d_2170_i0_: int
        d_2170_i0_ = 0
        with _dafny.label("8"):
            while not (p1.f3) or (not (not(False)) or (default__.fm4(p1.f3, globalState))):
                with _dafny.c_label("8"):
                    if (d_2170_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_2170_i0_ = (d_2170_i0_) + (1)
                    d_2171_v0_: C1
                    nw339_ = C1()
                    nw339_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2172_i1_ in range(default__.abs(497))]), p1.f3)
                    d_2171_v0_ = nw339_
                    d_2173_v1_: _dafny.Map
                    d_2173_v1_ = _dafny.Map({(d_2171_v0_).f13: (p1).f4})
                    d_2174_v2_: _dafny.Seq
                    d_2174_v2_ = _dafny.SeqWithoutIsStrInference([p0, (p0) + (self.f7), (self).fm0(globalState), p0, self.f7])
                    d_2175_v3_: _dafny.Seq
                    d_2175_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fscgsd"))
                    rhs289_ = d_2173_v1_
                    rhs290_ = ((_dafny.SeqWithoutIsStrInference([p0, self.f7])) + (d_2174_v2_)).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference([p0, self.f7])) + (d_2174_v2_))), default__.safeDivisionInt(p0, p0))
                    rhs291_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psukcoa"))) + (d_2175_v3_)
                    d_2173_v1_ = rhs289_
                    d_2174_v2_ = rhs290_
                    d_2175_v3_ = rhs291_
                    (p1).f3 = True
                    d_2176_v4_: _dafny.Map
                    d_2176_v4_ = _dafny.Map({494: (p1).f4})
                    d_2177_v5_: _dafny.MultiSet
                    d_2177_v5_ = _dafny.MultiSet([p1.f3, (self).f4, p1.f3, (p1).f4, (p1).f4])
                    d_2178_v6_: _dafny.Map
                    d_2178_v6_ = _dafny.Map({d_2177_v5_: len(d_2175_v3_)})
                    d_2179_v7_: C4
                    nw340_ = C4()
                    nw340_.ctor__(((d_2176_v4_)[p0] if (p0) in (d_2176_v4_) else p1.f3), ((d_2178_v6_)[d_2177_v5_] if (d_2177_v5_) in (d_2178_v6_) else self.f7), p1.f3, not((d_2171_v0_).f13))
                    d_2179_v7_ = nw340_
                    pass
            pass
        (self).f3 = False
        d_2180_v8_: _dafny.Array
        def lambda124_(d_2181_p0_):
            def lambda125_(d_2182_i2_):
                return D14_DC30(_dafny.MultiSet([d_2181_p0_, d_2181_p0_, 349, len(_dafny.Map({_dafny.CodePoint('v'): False}))]))

            return lambda125_

        init69_ = lambda124_(p0)
        nw341_ = _dafny.Array(None, 7)
        for i0_69_ in range(nw341_.length(0)):
            nw341_[i0_69_] = init69_(i0_69_)
        d_2180_v8_ = nw341_
        d_2180_v8_ = d_2180_v8_
        d_2183_v9_: _dafny.Array
        nw342_ = _dafny.Array(_dafny.CodePoint('D'), 18)
        d_2183_v9_ = nw342_
        d_2184_v10_: _dafny.Array
        nw343_ = _dafny.Array(int(0), 19)
        d_2184_v10_ = nw343_
        index355_ = default__.safeIndex(234, (d_2184_v10_).length(0))
        (d_2184_v10_)[index355_] = 665
        d_2185_v11_: _dafny.Map
        d_2185_v11_ = _dafny.Map({(p1).f4: p0})
        d_2186_v12_: D16
        d_2186_v12_ = D16_DC35((d_2185_v11_ if (p1).f4 else d_2185_v11_))
        d_2187_v13_: _dafny.Array
        nw344_ = _dafny.Array(None, 2)
        nw344_[int(0)] = (self).f6
        nw344_[int(1)] = (self).f6
        d_2187_v13_ = nw344_
        index356_ = default__.safeIndex(984, (d_2187_v13_).length(0))
        (d_2187_v13_)[index356_] = (self).f6
        d_2188_v14_: D21
        d_2188_v14_ = D21_DC51(d_2183_v9_)
        d_2189_v15_: _dafny.MultiSet
        d_2189_v15_ = _dafny.MultiSet([self.f7, self.f7])
        d_2190_v16_: _dafny.Map
        d_2190_v16_ = _dafny.Map({p0: p1.f3})
        d_2191_v17_: _dafny.Seq
        d_2191_v17_ = _dafny.SeqWithoutIsStrInference([p0])
        d_2192_v18_: _dafny.Array
        nw345_ = _dafny.Array(None, 28)
        nw345_[int(0)] = not(self.f3)
        nw345_[int(1)] = True
        nw345_[int(2)] = (self).f4
        nw345_[int(3)] = (self).f4
        nw345_[int(4)] = (self).f4
        nw345_[int(5)] = p1.f3
        nw345_[int(6)] = self.f3
        nw345_[int(7)] = self.f3
        nw345_[int(8)] = p1.f3
        nw345_[int(9)] = (self).f4
        nw345_[int(10)] = p1.f3
        nw345_[int(11)] = (self).f4
        nw345_[int(12)] = self.f3
        nw345_[int(13)] = p1.f3
        nw345_[int(14)] = not(self.f3)
        nw345_[int(15)] = self.f3
        nw345_[int(16)] = not(p1.f3)
        nw345_[int(17)] = self.f3
        nw345_[int(18)] = (p1).f4
        nw345_[int(19)] = (self).f4
        nw345_[int(20)] = self.f3
        nw345_[int(21)] = False
        nw345_[int(22)] = self.f3
        nw345_[int(23)] = not(not(p1.f3))
        nw345_[int(24)] = ((d_2190_v16_)[len(d_2191_v17_)] if (len(d_2191_v17_)) in (d_2190_v16_) else True)
        nw345_[int(25)] = False
        nw345_[int(26)] = (self).f4
        nw345_[int(27)] = (self).f4
        d_2192_v18_ = nw345_
        d_2193_v19_: _dafny.Map
        d_2193_v19_ = _dafny.Map({default__.fm14(globalState): d_2192_v18_})
        d_2194_v20_: D19
        d_2194_v20_ = D19_DC45((0) - (self.f7), len(d_2193_v19_), (self).f4, p1.f3, _dafny.SeqWithoutIsStrInference([p1.f3]))
        index357_ = default__.safeIndex(234, (d_2184_v10_).length(0))
        index358_ = default__.safeIndex(984, (d_2187_v13_).length(0))
        rhs292_ = (d_2188_v14_).cf83
        rhs293_ = (-408) <= (((d_2189_v15_)[(0) - (self.f7)] if ((0) - (self.f7)) in (d_2189_v15_) else self.f7))
        rhs294_ = (d_2194_v20_).cf70
        rhs295_ = d_2186_v12_
        rhs296_ = (self).f6
        lhs214_ = p1
        lhs215_ = d_2184_v10_
        lhs216_ = default__.safeIndex(234, (d_2184_v10_).length(0))
        lhs217_ = d_2187_v13_
        lhs218_ = default__.safeIndex(984, (d_2187_v13_).length(0))
        d_2183_v9_ = rhs292_
        lhs214_.f3 = rhs293_
        lhs215_[lhs216_] = rhs294_
        d_2186_v12_ = rhs295_
        lhs217_[lhs218_] = rhs296_
        source27_ = d_2186_v12_
        if source27_.is_DC36:
            d_2195___mcc_h0_ = source27_.cf51
            d_2196___mcc_h1_ = source27_.cf52
            d_2197___mcc_h2_ = source27_.cf53
            d_2198___mcc_h3_ = source27_.cf54
            d_2199_cf54_ = d_2198___mcc_h3_
            d_2200_cf53_ = d_2197___mcc_h2_
            d_2201_cf52_ = d_2196___mcc_h1_
            d_2202_cf51_ = d_2195___mcc_h0_
            d_2203_v21_: C2
            nw346_ = C2()
            nw346_.ctor__(False, d_2200_cf53_, ((d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))]) <= (self.f7), p1.f3)
            d_2203_v21_ = nw346_
            (self).f7 = p0
            d_2204_v22_: C0
            nw347_ = C0()
            nw347_.ctor__()
            d_2204_v22_ = nw347_
            if (d_2189_v15_) == ((d_2189_v15_).intersection(d_2189_v15_)):
                def iife196_():
                    def iife198_():
                        coll74_ = _dafny.Set()
                        compr_74_: int
                        for compr_74_ in (default__.fm19(globalState)).Elements:
                            d_2207_v23_: int = compr_74_
                            if (d_2207_v23_) in (default__.fm19(globalState)):
                                coll74_ = coll74_.union(_dafny.Set([(d_2207_v23_) * (-208)]))
                        return _dafny.Set(coll74_)
                    coll72_ = _dafny.Set()
                    def iife197_():
                        coll73_ = _dafny.Set()
                        compr_73_: int
                        for compr_73_ in (default__.fm19(globalState)).Elements:
                            d_2205_v23_: int = compr_73_
                            if (d_2205_v23_) in (default__.fm19(globalState)):
                                coll73_ = coll73_.union(_dafny.Set([(d_2205_v23_) * (-208)]))
                        return _dafny.Set(coll73_)
                    compr_72_: int
                    for compr_72_ in (iife197_()
                    ).Elements:
                        d_2206_v24_: int = compr_72_
                        if (d_2206_v24_) in (iife198_()
                        ):
                            def iife199_():
                                coll75_ = _dafny.Map()
                                compr_75_: int
                                for compr_75_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxtq"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cb")))])), (0) - (-986)])).Elements:
                                    d_2208_v25_: int = compr_75_
                                    if (d_2208_v25_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxtq"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cb")))])), (0) - (-986)])):
                                        coll75_[default__.safeModuloInt(d_2208_v25_, 178)] = False
                                return _dafny.Map(coll75_)
                            coll72_ = coll72_.union(_dafny.Set([default__.safeModuloInt(d_2206_v24_, (678 if not(False) else len(_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(not(False))])), 824, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyhfjwks"))), len(_dafny.Set({not(True)}))]), _dafny.SeqWithoutIsStrInference([719, len(iife199_()
)])}))))]))
                    return _dafny.Set(coll72_)
                rhs297_ = (d_2203_v21_).f11
                rhs298_ = iife196_()

                rhs299_ = ((d_2203_v21_).f10 if not((self).f4) else (d_2203_v21_).f10)
                rhs300_ = (((d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))]) + ((d_2189_v15_).cardinality)) >= ((0) - ((d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))]))
                lhs219_ = p1
                lhs220_ = p1
                lhs219_.f3 = rhs297_
                d_2202_cf51_ = rhs298_
                d_2200_cf53_ = rhs299_
                lhs220_.f3 = rhs300_
                d_2209_v26_: D1
                d_2209_v26_ = D1_DC5(p1.f3, d_2200_cf53_, (self).fm3(globalState), p0)
                d_2210_v27_: _dafny.Map
                d_2210_v27_ = _dafny.Map({(d_2209_v26_).cf11: (D1_DC3(d_2192_v18_)).cf6})
                d_2210_v27_ = (d_2210_v27_).set((d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))], d_2192_v18_)
                (self).f7 = (p0) * ((d_2191_v17_)[default__.safeIndex(len(d_2202_cf51_), len(d_2191_v17_))])
                (d_2203_v21_).m9(globalState)
                index359_ = default__.safeIndex(3, (d_2192_v18_).length(0))
                (d_2192_v18_)[index359_] = p1.f3
                d_2211_v29_: _dafny.Seq
                d_2211_v29_ = _dafny.SeqWithoutIsStrInference([d_2192_v18_])
                d_2212_v30_: _dafny.Map
                d_2212_v30_ = _dafny.Map({len(d_2211_v29_): p0})
                d_2213_v31_: _dafny.Map
                def iife200_():
                    coll76_ = _dafny.Map()
                    compr_76_: int
                    for compr_76_ in _dafny.IntegerRange(952, -100):
                        d_2214_v28_: int = compr_76_
                        if ((952) <= (d_2214_v28_)) and ((d_2214_v28_) < (-100)):
                            coll76_[(d_2214_v28_) * (len((d_2187_v13_)[default__.safeIndex(984, (d_2187_v13_).length(0))]))] = 105
                    return _dafny.Map(coll76_)
                d_2213_v31_ = _dafny.Map({(iife200_()
                ) | (d_2212_v30_): d_2199_cf54_})
                index360_ = default__.safeIndex(234, (d_2184_v10_).length(0))
                index361_ = default__.safeIndex(3, (d_2192_v18_).length(0))
                rhs301_ = p0
                rhs302_ = d_2187_v13_
                rhs303_ = ((self.f7) > (self.f7)) and (d_2201_cf52_)
                rhs304_ = (983) * ((836 if (d_2203_v21_).f10 else p0))
                rhs305_ = d_2213_v31_
                lhs221_ = d_2184_v10_
                lhs222_ = default__.safeIndex(234, (d_2184_v10_).length(0))
                lhs223_ = d_2192_v18_
                lhs224_ = default__.safeIndex(3, (d_2192_v18_).length(0))
                lhs221_[lhs222_] = rhs301_
                d_2187_v13_ = rhs302_
                lhs223_[lhs224_] = rhs303_
                r0 = rhs304_
                d_2213_v31_ = rhs305_
            elif True:
                d_2215_v32_: D1
                d_2215_v32_ = D1_DC5(p1.f3, p1.f3, (d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))], (d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))])
                d_2216_v33_: _dafny.Seq
                d_2216_v33_ = _dafny.SeqWithoutIsStrInference([d_2184_v10_, d_2184_v10_, d_2184_v10_])
                d_2217_v34_: D19
                d_2217_v34_ = D19_DC46(p0, d_2201_cf52_, p1.f3, d_2215_v32_, len(d_2216_v33_))
                d_2218_v35_: _dafny.Map
                d_2218_v35_ = _dafny.Map({d_2200_cf53_: (d_2203_v21_).f11})
                d_2219_v36_: _dafny.Set
                d_2219_v36_ = _dafny.Set({False})
                d_2217_v34_ = D19_DC46((735) - ((d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))]), (p1).f4, (d_2219_v36_).ispropersubset(_dafny.Set({((d_2218_v35_)[d_2200_cf53_] if (d_2200_cf53_) in (d_2218_v35_) else ((d_2190_v16_)[(d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))]] if ((d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))]) in (d_2190_v16_) else False))})), D1_DC5(p1.f3, (p1).f4, self.f7, (d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))]), (d_2204_v22_).fm10((d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))], d_2199_cf54_, d_2189_v15_, globalState))
                d_2220_v37_: _dafny.Array
                nw348_ = _dafny.Array(None, 19)
                d_2220_v37_ = nw348_
                d_2221_v38_: _dafny.Map
                d_2221_v38_ = _dafny.Map({p1.f3: d_2220_v37_})
                d_2222_v39_: D19
                d_2222_v39_ = D19_DC44(d_2221_v38_)
                d_2223_v40_: _dafny.Map
                d_2223_v40_ = _dafny.Map({d_2222_v39_: default__.safeDivisionInt(self.f7, len(d_2185_v11_))})
                d_2223_v40_ = (d_2223_v40_).set(d_2222_v39_, (d_2189_v15_).cardinality)
                (self).f7 = default__.safeModuloInt(p0, p0)
                d_2218_v35_ = (d_2218_v35_).set(not(p1.f3), (d_2203_v21_).f11)
                d_2224_v41_: _dafny.Seq
                d_2224_v41_ = _dafny.SeqWithoutIsStrInference([d_2200_cf53_])
                d_2225_v42_: D12
                d_2225_v42_ = D12_DC26(d_2219_v36_)
                d_2226_v43_: D12
                d_2226_v43_ = D12_DC28(d_2225_v42_)
                d_2227_v44_: D0
                d_2227_v44_ = D0_DC0((d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))])
                d_2228_v45_: _dafny.Set
                d_2228_v45_ = _dafny.Set({d_2226_v43_})
                pat_let_tv66_ = d_2225_v42_
                def iife201_(_pat_let62_0):
                    def iife202_(d_2229_dt__update__tmp_h0_):
                        def iife203_(_pat_let63_0):
                            def iife204_(d_2230_dt__update_hcf41_h0_):
                                return D12_DC28(d_2230_dt__update_hcf41_h0_)
                            return iife204_(_pat_let63_0)
                        return iife203_(pat_let_tv66_)
                    return iife202_(_pat_let62_0)
                rhs306_ = len((_dafny.SeqWithoutIsStrInference([p1.f3])) + (d_2224_v41_))
                rhs307_ = d_2184_v10_
                rhs308_ = (_dafny.Set({d_2226_v43_, default__.fm38((d_2227_v44_).cf0, self.f3, globalState), d_2226_v43_, iife201_(d_2226_v43_), d_2226_v43_})).issubset(d_2228_v45_)
                lhs225_ = self
                lhs225_.f7 = rhs306_
                d_2184_v10_ = rhs307_
                d_2201_cf52_ = rhs308_
        elif True:
            d_2231___mcc_h4_ = source27_.cf50
            d_2232_cf50_ = d_2231___mcc_h4_
            d_2233_v46_: str
            d_2233_v46_ = _dafny.CodePoint('f')
            d_2234_v47_: _dafny.Map
            d_2234_v47_ = _dafny.Map({True: ((d_2190_v16_)[(self).fm3(globalState)] if ((self).fm3(globalState)) in (d_2190_v16_) else (p1).f4)})
            d_2235_v48_: D8
            d_2235_v48_ = D8_DC20(d_2233_v46_, ((d_2234_v47_)[False] if (False) in (d_2234_v47_) else p1.f3))
            index362_ = default__.safeIndex(608, (d_2192_v18_).length(0))
            (d_2192_v18_)[index362_] = (p1).f4
            index363_ = default__.safeIndex(608, (d_2192_v18_).length(0))
            rhs309_ = d_2235_v48_
            rhs310_ = (self).f4
            lhs226_ = d_2192_v18_
            lhs227_ = default__.safeIndex(608, (d_2192_v18_).length(0))
            d_2235_v48_ = rhs309_
            lhs226_[lhs227_] = rhs310_
            (p1).f3 = not(self.f3)
            d_2236_v49_: _dafny.Seq
            d_2236_v49_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_2233_v46_ for d_2237_i3_ in range(default__.abs(355))])) + ((d_2187_v13_)[default__.safeIndex(984, (d_2187_v13_).length(0))]), (self).f6, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hxq"))])
            d_2236_v49_ = d_2236_v49_
            d_2238_v50_: _dafny.Map
            d_2238_v50_ = _dafny.Map({d_2189_v15_: not(self.f3)})
            d_2239_v51_: D17
            d_2239_v51_ = D17_DC39(d_2184_v10_, (p1).f4, ((p1).f4 if (d_2192_v18_)[default__.safeIndex(608, (d_2192_v18_).length(0))] else (self).f4), (len((d_2238_v50_).set(d_2189_v15_, self.f3))) + (p0), self.f7)
            source28_ = d_2239_v51_
            if source28_.is_DC38:
                d_2240___mcc_h5_ = source28_.cf56
                d_2241_cf56_ = d_2240___mcc_h5_
                index364_ = default__.safeIndex(234, (d_2184_v10_).length(0))
                (d_2184_v10_)[index364_] = p0
                index365_ = default__.safeIndex(608, (d_2192_v18_).length(0))
                (d_2192_v18_)[index365_] = not(default__.fm14(globalState))
                d_2242_v52_: _dafny.MultiSet
                d_2242_v52_ = _dafny.MultiSet([(p1).f4, (p1).f4])
                d_2243_v53_: _dafny.Map
                d_2243_v53_ = _dafny.Map({(self).f6: d_2242_v52_})
                d_2244_v54_: _dafny.Set
                d_2244_v54_ = _dafny.Set({False, False, (((d_2243_v53_)[d_2241_cf56_] if (d_2241_cf56_) in (d_2243_v53_) else default__.fm20(p1.f3, globalState))).ispropersubset(d_2242_v52_), ((d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))]) == (self.f7), (self).f4})
                d_2245_v55_: _dafny.Seq
                d_2245_v55_ = _dafny.SeqWithoutIsStrInference([p1.f3])
                pat_let_tv67_ = d_2244_v54_
                def iife205_(_pat_let64_0):
                    def iife206_(d_2246_dt__update__tmp_h1_):
                        def iife207_(_pat_let65_0):
                            def iife208_(d_2247_dt__update_hcf39_h0_):
                                return D12_DC26(d_2247_dt__update_hcf39_h0_)
                            return iife208_(_pat_let65_0)
                        return iife207_(pat_let_tv67_)
                    return iife206_(_pat_let64_0)
                d_2244_v54_ = ((iife205_(D12_DC26(_dafny.Set({(self).f4})))).cf39) - ((d_2244_v54_).intersection(default__.fm33(((d_2189_v15_)[self.f7] if (self.f7) in (d_2189_v15_) else len((d_2187_v13_)[default__.safeIndex(984, (d_2187_v13_).length(0))])), d_2245_v55_, globalState)))
                (self).f3 = p1.f3
            elif source28_.is_DC39:
                d_2248___mcc_h6_ = source28_.cf57
                d_2249___mcc_h7_ = source28_.cf58
                d_2250___mcc_h8_ = source28_.cf59
                d_2251___mcc_h9_ = source28_.cf60
                d_2252___mcc_h10_ = source28_.cf61
                d_2253_cf61_ = d_2252___mcc_h10_
                d_2254_cf60_ = d_2251___mcc_h9_
                d_2255_cf59_ = d_2250___mcc_h8_
                d_2256_cf58_ = d_2249___mcc_h7_
                d_2257_cf57_ = d_2248___mcc_h6_
                d_2258_v56_: _dafny.Array
                nw349_ = _dafny.Array(_dafny.MultiSet({}), 10)
                d_2258_v56_ = nw349_
                d_2259_v57_: _dafny.Seq
                d_2259_v57_ = _dafny.SeqWithoutIsStrInference([d_2256_cf58_, d_2255_cf59_])
                index366_ = default__.safeIndex(714, (d_2258_v56_).length(0))
                (d_2258_v56_)[index366_] = default__.fm20(default__.fm6(default__.fm31(p1.f3, p0, d_2256_cf58_, d_2259_v57_, globalState), globalState), globalState)
                d_2260_v58_: _dafny.MultiSet
                d_2260_v58_ = _dafny.MultiSet([not((d_2192_v18_)[default__.safeIndex(608, (d_2192_v18_).length(0))])])
                index367_ = default__.safeIndex(714, (d_2258_v56_).length(0))
                (d_2258_v56_)[index367_] = d_2260_v58_
                (globalState).f0 = (p1).f4
                d_2261_v59_: C1
                nw350_ = C1()
                nw350_.ctor__((self).f6, p1.f3)
                d_2261_v59_ = nw350_
                d_2262_v60_: _dafny.Seq
                d_2262_v60_ = _dafny.SeqWithoutIsStrInference([d_2261_v59_])
                d_2263_v61_: _dafny.Array
                def lambda126_(d_2264_i4_):
                    return True

                init70_ = lambda126_
                nw351_ = _dafny.Array(None, 7)
                for i0_70_ in range(nw351_.length(0)):
                    nw351_[i0_70_] = init70_(i0_70_)
                d_2263_v61_ = nw351_
                d_2265_v62_: _dafny.Array
                nw352_ = _dafny.Array(None, 12)
                nw352_[int(0)] = d_2261_v59_
                nw352_[int(1)] = d_2261_v59_
                nw352_[int(2)] = d_2261_v59_
                nw352_[int(3)] = d_2261_v59_
                nw352_[int(4)] = d_2261_v59_
                nw352_[int(5)] = (d_2262_v60_)[default__.safeIndex(999, len(d_2262_v60_))]
                nw352_[int(6)] = d_2261_v59_
                nw352_[int(7)] = (D15_DC33(d_2263_v61_, d_2261_v59_)).cf48
                nw352_[int(8)] = d_2261_v59_
                nw352_[int(9)] = d_2261_v59_
                nw352_[int(10)] = d_2261_v59_
                nw352_[int(11)] = d_2261_v59_
                d_2265_v62_ = nw352_
                d_2266_v63_: _dafny.Map
                d_2266_v63_ = _dafny.Map({(d_2261_v59_).f12: d_2265_v62_})
                d_2267_v64_: _dafny.Map
                d_2267_v64_ = _dafny.Map({(d_2258_v56_)[default__.safeIndex(714, (d_2258_v56_).length(0))]: d_2265_v62_})
                d_2268_v65_: _dafny.Array
                nw353_ = _dafny.Array(None, 12)
                nw353_[int(0)] = d_2265_v62_
                nw353_[int(1)] = ((d_2266_v63_)[_dafny.SeqWithoutIsStrInference([d_2233_v46_ for d_2269_i5_ in range(default__.abs(807))])] if (_dafny.SeqWithoutIsStrInference([d_2233_v46_ for d_2270_i5_ in range(default__.abs(807))])) in (d_2266_v63_) else d_2265_v62_)
                nw353_[int(2)] = d_2265_v62_
                nw353_[int(3)] = d_2265_v62_
                nw353_[int(4)] = d_2265_v62_
                nw353_[int(5)] = d_2265_v62_
                nw353_[int(6)] = (d_2265_v62_ if (self).f4 else d_2265_v62_)
                nw353_[int(7)] = d_2265_v62_
                nw353_[int(8)] = d_2265_v62_
                nw353_[int(9)] = d_2265_v62_
                nw353_[int(10)] = d_2265_v62_
                nw353_[int(11)] = ((d_2267_v64_)[d_2260_v58_] if (d_2260_v58_) in (d_2267_v64_) else d_2265_v62_)
                d_2268_v65_ = nw353_
                d_2271_v66_: _dafny.Seq
                d_2271_v66_ = _dafny.SeqWithoutIsStrInference([d_2265_v62_, d_2265_v62_, d_2265_v62_, d_2265_v62_, d_2265_v62_])
                index368_ = default__.safeIndex(626, (d_2268_v65_).length(0))
                (d_2268_v65_)[index368_] = (d_2271_v66_)[default__.safeIndex(d_2254_cf60_, len(d_2271_v66_))]
                d_2272_v67_: _dafny.Set
                d_2272_v67_ = _dafny.Set({d_2254_cf60_})
                d_2273_v68_: _dafny.Map
                d_2273_v68_ = _dafny.Map({p0: d_2233_v46_})
                d_2274_v70_: _dafny.Seq
                def iife209_():
                    coll77_ = _dafny.Map()
                    compr_77_: int
                    for compr_77_ in (d_2191_v17_).Elements:
                        d_2275_v69_: int = compr_77_
                        if (d_2275_v69_) in (d_2191_v17_):
                            coll77_[default__.safeDivisionInt(d_2275_v69_, d_2253_cf61_)] = d_2233_v46_
                    return _dafny.Map(coll77_)
                d_2274_v70_ = _dafny.SeqWithoutIsStrInference([d_2273_v68_, d_2273_v68_, iife209_()
                ])
                d_2276_v71_: _dafny.Map
                d_2276_v71_ = _dafny.Map({d_2272_v67_: ((d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))]) + ((0) - (len((d_2274_v70_).set(default__.safeIndex(len((self).f6), len(d_2274_v70_)), _dafny.Map({(d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))]: d_2233_v46_})))))})
                d_2277_v72_: D1
                d_2277_v72_ = D1_DC4((0) - (d_2253_cf61_))
                d_2278_v73_: D1
                d_2278_v73_ = D1_DC5(d_2255_cf59_, (p1).f4, (d_2277_v72_).cf7, -848)
                index369_ = default__.safeIndex(626, (d_2268_v65_).length(0))
                index370_ = default__.safeIndex(608, (d_2192_v18_).length(0))
                rhs311_ = d_2184_v10_
                rhs312_ = d_2265_v62_
                rhs313_ = ((d_2276_v71_) | (d_2276_v71_)) | (d_2276_v71_)
                rhs314_ = (d_2278_v73_).cf9
                lhs228_ = d_2268_v65_
                lhs229_ = default__.safeIndex(626, (d_2268_v65_).length(0))
                lhs230_ = d_2192_v18_
                lhs231_ = default__.safeIndex(608, (d_2192_v18_).length(0))
                d_2257_cf57_ = rhs311_
                lhs228_[lhs229_] = rhs312_
                d_2276_v71_ = rhs313_
                lhs230_[lhs231_] = rhs314_
                index371_ = default__.safeIndex(608, (d_2192_v18_).length(0))
                (d_2192_v18_)[index371_] = p1.f3
            elif source28_.is_DC37:
                d_2279___mcc_h11_ = source28_.cf55
                d_2280_cf55_ = d_2279___mcc_h11_
                d_2281_v74_: _dafny.Map
                d_2281_v74_ = _dafny.Map({d_2184_v10_: True})
                index372_ = default__.safeIndex(608, (d_2192_v18_).length(0))
                rhs315_ = not(not((d_2184_v10_) in (d_2281_v74_)))
                rhs316_ = not((p0) == (default__.safeModuloInt(p0, p0)))
                lhs232_ = d_2192_v18_
                lhs233_ = default__.safeIndex(608, (d_2192_v18_).length(0))
                lhs234_ = p1
                lhs232_[lhs233_] = rhs315_
                lhs234_.f3 = rhs316_
                d_2282_v75_: C1
                nw354_ = C1()
                nw354_.ctor__(((self).f6) + ((d_2187_v13_)[default__.safeIndex(984, (d_2187_v13_).length(0))]), p1.f3)
                d_2282_v75_ = nw354_
                index373_ = default__.safeIndex(608, (d_2192_v18_).length(0))
                (d_2192_v18_)[index373_] = self.f3
                r0 = (len(_dafny.SeqWithoutIsStrInference([p0 for d_2283_i6_ in range(default__.abs(779))]))) + ((d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))])
            elif True:
                d_2284___mcc_h12_ = source28_.cf62
                d_2285_cf62_ = d_2284___mcc_h12_
                index374_ = default__.safeIndex(234, (d_2184_v10_).length(0))
                (d_2184_v10_)[index374_] = (-692 if (d_2235_v48_).cf34 else 973)
                d_2286_v76_: _dafny.Set
                d_2286_v76_ = _dafny.Set({(p1).f4})
                d_2287_v77_: _dafny.Map
                d_2287_v77_ = _dafny.Map({d_2286_v76_: self.f7})
                d_2287_v77_ = (d_2287_v77_).set(d_2286_v76_, self.f7)
                index375_ = default__.safeIndex(106, (d_2183_v9_).length(0))
                (d_2183_v9_)[index375_] = d_2233_v46_
                index376_ = default__.safeIndex(106, (d_2183_v9_).length(0))
                (d_2183_v9_)[index376_] = _dafny.CodePoint('e')
                r0 = (d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))]
        d_2288_v78_: _dafny.Seq
        d_2288_v78_ = _dafny.SeqWithoutIsStrInference([(self).f4])
        hi13_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjx")))
        for d_2289_i7_ in range(len(d_2288_v78_), hi13_):
            (self).f7 = p0
            (self).f7 = (p0 if (p1).f4 else (p0) - ((d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))]))
            d_2290_v79_: _dafny.Array
            nw355_ = _dafny.Array(_dafny.Seq({}), 12)
            d_2290_v79_ = nw355_
            index377_ = default__.safeIndex(598, (d_2290_v79_).length(0))
            (d_2290_v79_)[index377_] = d_2288_v78_
            index378_ = default__.safeIndex(598, (d_2290_v79_).length(0))
            (d_2290_v79_)[index378_] = d_2288_v78_
            (self).f7 = -240
        r0 = (d_2184_v10_)[default__.safeIndex(234, (d_2184_v10_).length(0))]
        return r0

    def m2(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_2291_v0_: _dafny.Array
        def lambda127_(d_2292_p1_):
            def lambda128_(d_2293_i0_):
                return not(d_2292_p1_.f3)

            return lambda128_

        init71_ = lambda127_(p1)
        nw356_ = _dafny.Array(None, 24)
        for i0_71_ in range(nw356_.length(0)):
            nw356_[i0_71_] = init71_(i0_71_)
        d_2291_v0_ = nw356_
        d_2294_v1_: D1
        d_2294_v1_ = D1_DC3(d_2291_v0_)
        source29_ = d_2294_v1_
        if source29_.is_DC4:
            d_2295___mcc_h0_ = source29_.cf7
            d_2296_cf7_ = d_2295___mcc_h0_
            d_2297_v2_: str
            d_2297_v2_ = _dafny.CodePoint('k')
            d_2298_v3_: _dafny.Set
            d_2298_v3_ = _dafny.Set({self.f7, self.f7})
            d_2299_v4_: _dafny.Set
            d_2299_v4_ = _dafny.Set({len(d_2298_v3_), d_2296_cf7_})
            d_2300_v5_: _dafny.Set
            d_2300_v5_ = _dafny.Set({d_2299_v4_, d_2299_v4_, d_2299_v4_})
            d_2301_v6_: _dafny.Map
            d_2301_v6_ = _dafny.Map({(p1).f4: False})
            rhs317_ = (d_2300_v5_).isdisjoint(d_2300_v5_)
            rhs318_ = len(((d_2301_v6_) | (d_2301_v6_)) | (d_2301_v6_))
            rhs319_ = d_2296_cf7_
            rhs320_ = p0
            rhs321_ = self.f3
            lhs235_ = p1
            lhs236_ = self
            lhs237_ = self
            lhs238_ = self
            lhs235_.f3 = rhs317_
            lhs236_.f7 = rhs318_
            lhs237_.f7 = rhs319_
            d_2297_v2_ = rhs320_
            lhs238_.f3 = rhs321_
            d_2302_v7_: _dafny.Seq
            d_2302_v7_ = _dafny.SeqWithoutIsStrInference([self.f3, p1.f3])
            d_2303_v8_: _dafny.Seq
            d_2303_v8_ = _dafny.SeqWithoutIsStrInference([self.f7, self.f7, self.f7])
            d_2304_v9_: _dafny.Array
            nw357_ = _dafny.Array(int(0), 27)
            d_2304_v9_ = nw357_
            d_2305_v10_: D17
            d_2305_v10_ = D17_DC39(d_2304_v9_, self.f3, (p1).f4, d_2296_cf7_, d_2296_cf7_)
            d_2306_v11_: _dafny.MultiSet
            d_2306_v11_ = _dafny.MultiSet([self.f7, d_2296_cf7_, (d_2305_v10_).cf61, self.f7, (_dafny.MultiSet([p0, p0])).cardinality])
            rhs322_ = (p1).f4
            rhs323_ = (d_2301_v6_) | ((d_2301_v6_) | (d_2301_v6_))
            rhs324_ = default__.fm32(self.f3, True, (d_2302_v7_)[default__.safeIndex(d_2296_cf7_, len(d_2302_v7_))], not((_dafny.MultiSet(d_2303_v8_)).isdisjoint(d_2306_v11_)), globalState)
            lhs239_ = p1
            lhs239_.f3 = rhs322_
            d_2301_v6_ = rhs323_
            d_2298_v3_ = rhs324_
            d_2307_v12_: _dafny.Array
            nw358_ = _dafny.Array(_dafny.Set({}), 5)
            d_2307_v12_ = nw358_
            d_2308_v13_: D15
            d_2308_v13_ = D15_DC32(d_2307_v12_)
            d_2309_v14_: D15
            d_2309_v14_ = D15_DC34(D15_DC34(d_2308_v13_))
            d_2310_v15_: D15
            d_2310_v15_ = D15_DC34(d_2309_v14_)
            source30_ = d_2310_v15_
            if source30_.is_DC33:
                d_2311___mcc_h6_ = source30_.cf47
                d_2312___mcc_h7_ = source30_.cf48
                d_2313_cf48_ = d_2312___mcc_h7_
                d_2314_cf47_ = d_2311___mcc_h6_
                d_2315_v16_: D20
                d_2315_v16_ = D20_DC48(d_2303_v8_)
                d_2316_v17_: _dafny.Map
                d_2316_v17_ = _dafny.Map({(d_2315_v16_).cf80: self.f7})
                rhs325_ = len((d_2302_v7_ if (not((d_2313_cf48_).f13) if (p1).f4 else True) else d_2302_v7_))
                rhs326_ = d_2297_v2_
                rhs327_ = (self).f4
                rhs328_ = ((d_2316_v17_) | (d_2316_v17_) if p1.f3 else d_2316_v17_)
                lhs240_ = self
                lhs241_ = p1
                lhs240_.f7 = rhs325_
                d_2297_v2_ = rhs326_
                lhs241_.f3 = rhs327_
                d_2316_v17_ = rhs328_
                def lambda129_(d_2317_i1_):
                    return (d_2317_i1_) - (self.f7)

                init72_ = lambda129_
                nw359_ = _dafny.Array(None, 15)
                for i0_72_ in range(nw359_.length(0)):
                    nw359_[i0_72_] = init72_(i0_72_)
                d_2304_v9_ = nw359_
                d_2318_v18_: D17
                d_2318_v18_ = D17_DC38(((self).f6) + ((self).f6))
                rhs329_ = d_2318_v18_
                rhs330_ = d_2301_v6_
                rhs331_ = (d_2296_cf7_ if not((d_2298_v3_).isdisjoint(d_2298_v3_)) else d_2296_cf7_)
                d_2318_v18_ = rhs329_
                d_2301_v6_ = rhs330_
                d_2296_cf7_ = rhs331_
                (self).f7 = self.f7
            elif source30_.is_DC32:
                d_2319___mcc_h8_ = source30_.cf46
                d_2320_cf46_ = d_2319___mcc_h8_
                r0 = _dafny.SeqWithoutIsStrInference([(D0_DC1((p1).f4, d_2297_v2_)).cf2 for d_2321_i2_ in range(default__.abs(186))])
                (self).f3 = (p1).f4
                (p1).f3 = (p1).f4
                d_2322_v19_: _dafny.MultiSet
                d_2322_v19_ = _dafny.MultiSet([p1.f3])
                d_2323_v20_: _dafny.Seq
                d_2323_v20_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_2303_v8_), d_2306_v11_])
                d_2324_v21_: _dafny.Seq
                d_2324_v21_ = _dafny.SeqWithoutIsStrInference([(d_2303_v8_).set(default__.safeIndex(len((self).f6), len(d_2303_v8_)), self.f7)])
                d_2325_v22_: _dafny.Array
                nw360_ = _dafny.Array(None, 26)
                nw360_[int(0)] = d_2306_v11_
                nw360_[int(1)] = d_2306_v11_
                nw360_[int(2)] = ((d_2306_v11_).set(-532, default__.abs(d_2296_cf7_))) | (_dafny.MultiSet([(0) - ((self).fm3(globalState))]))
                nw360_[int(3)] = _dafny.MultiSet(d_2303_v8_)
                nw360_[int(4)] = _dafny.MultiSet(d_2303_v8_)
                nw360_[int(5)] = (d_2306_v11_).intersection(d_2306_v11_)
                nw360_[int(6)] = ((d_2306_v11_).set(self.f7, default__.abs(((d_2322_v19_)[False] if (False) in (d_2322_v19_) else self.f7)))).intersection((d_2323_v20_)[default__.safeIndex(d_2296_cf7_, len(d_2323_v20_))])
                nw360_[int(7)] = (_dafny.MultiSet([400, self.f7])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2296_cf7_, d_2296_cf7_, (d_2322_v19_).cardinality])))
                nw360_[int(8)] = _dafny.MultiSet([d_2296_cf7_, self.f7, self.f7])
                nw360_[int(9)] = (d_2323_v20_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2297_v2_ for d_2326_i3_ in range(default__.abs(607))])), len(d_2323_v20_))]
                nw360_[int(10)] = d_2306_v11_
                nw360_[int(11)] = (d_2306_v11_).intersection(d_2306_v11_)
                nw360_[int(12)] = d_2306_v11_
                nw360_[int(13)] = d_2306_v11_
                nw360_[int(14)] = d_2306_v11_
                nw360_[int(15)] = d_2306_v11_
                nw360_[int(16)] = d_2306_v11_
                nw360_[int(17)] = _dafny.MultiSet(d_2303_v8_)
                nw360_[int(18)] = d_2306_v11_
                nw360_[int(19)] = _dafny.MultiSet([d_2296_cf7_, self.f7, d_2296_cf7_])
                nw360_[int(20)] = d_2306_v11_
                nw360_[int(21)] = d_2306_v11_
                nw360_[int(22)] = _dafny.MultiSet((d_2324_v21_)[default__.safeIndex(len((self).f6), len(d_2324_v21_))])
                nw360_[int(23)] = _dafny.MultiSet([self.f7])
                nw360_[int(24)] = d_2306_v11_
                nw360_[int(25)] = d_2306_v11_
                d_2325_v22_ = nw360_
                index379_ = default__.safeIndex(381, (d_2325_v22_).length(0))
                (d_2325_v22_)[index379_] = d_2306_v11_
                index380_ = default__.safeIndex(381, (d_2325_v22_).length(0))
                rhs332_ = self.f7
                rhs333_ = (d_2306_v11_ if (d_2296_cf7_) != (d_2296_cf7_) else d_2306_v11_)
                lhs242_ = d_2325_v22_
                lhs243_ = default__.safeIndex(381, (d_2325_v22_).length(0))
                d_2296_cf7_ = rhs332_
                lhs242_[lhs243_] = rhs333_
            elif True:
                d_2327___mcc_h9_ = source30_.cf49
                d_2328_cf49_ = d_2327___mcc_h9_
                d_2329_v23_: C1
                nw361_ = C1()
                nw361_.ctor__((self).f6, p1.f3)
                d_2329_v23_ = nw361_
                d_2330_v24_: D1
                d_2330_v24_ = D1_DC5((p1).f4, True, self.f7, self.f7)
                d_2331_v25_: D19
                d_2331_v25_ = D19_DC46(self.f7, p1.f3, self.f3, d_2330_v24_, d_2296_cf7_)
                r0 = (_dafny.SeqWithoutIsStrInference([d_2297_v2_ for d_2332_i4_ in range(default__.abs(169))]) if not ((self).f4) or ((d_2331_v25_).cf75) else (d_2329_v23_).f12)
                d_2333_v26_: _dafny.Array
                nw362_ = _dafny.Array(_dafny.CodePoint('D'), 19)
                d_2333_v26_ = nw362_
                d_2334_v27_: _dafny.Array
                nw363_ = _dafny.Array(None, 13)
                nw363_[int(0)] = d_2333_v26_
                nw363_[int(1)] = d_2333_v26_
                nw363_[int(2)] = d_2333_v26_
                nw363_[int(3)] = d_2333_v26_
                nw363_[int(4)] = d_2333_v26_
                nw363_[int(5)] = d_2333_v26_
                nw363_[int(6)] = d_2333_v26_
                nw363_[int(7)] = d_2333_v26_
                nw363_[int(8)] = d_2333_v26_
                nw363_[int(9)] = d_2333_v26_
                nw363_[int(10)] = d_2333_v26_
                nw363_[int(11)] = d_2333_v26_
                nw363_[int(12)] = d_2333_v26_
                d_2334_v27_ = nw363_
                d_2335_v28_: _dafny.Array
                nw364_ = _dafny.Array(None, 25)
                nw364_[int(0)] = d_2291_v0_
                nw364_[int(1)] = d_2291_v0_
                nw364_[int(2)] = d_2291_v0_
                nw364_[int(3)] = d_2291_v0_
                nw364_[int(4)] = d_2291_v0_
                nw364_[int(5)] = d_2291_v0_
                nw364_[int(6)] = d_2291_v0_
                nw364_[int(7)] = (d_2291_v0_ if (p1).f4 else d_2291_v0_)
                nw364_[int(8)] = d_2291_v0_
                nw364_[int(9)] = d_2291_v0_
                nw364_[int(10)] = d_2291_v0_
                nw364_[int(11)] = (d_2291_v0_ if (p1).f4 else d_2291_v0_)
                nw364_[int(12)] = d_2291_v0_
                nw364_[int(13)] = d_2291_v0_
                nw364_[int(14)] = d_2291_v0_
                nw364_[int(15)] = (d_2291_v0_ if False else d_2291_v0_)
                nw364_[int(16)] = d_2291_v0_
                nw364_[int(17)] = (d_2291_v0_ if p1.f3 else d_2291_v0_)
                nw364_[int(18)] = d_2291_v0_
                nw364_[int(19)] = d_2291_v0_
                nw364_[int(20)] = d_2291_v0_
                nw364_[int(21)] = d_2291_v0_
                nw364_[int(22)] = d_2291_v0_
                nw364_[int(23)] = d_2291_v0_
                nw364_[int(24)] = d_2291_v0_
                d_2335_v28_ = nw364_
                index381_ = default__.safeIndex(703, (d_2335_v28_).length(0))
                (d_2335_v28_)[index381_] = d_2291_v0_
                index382_ = default__.safeIndex(703, (d_2335_v28_).length(0))
                rhs334_ = d_2303_v8_
                rhs335_ = d_2334_v27_
                rhs336_ = d_2291_v0_
                rhs337_ = self.f7
                lhs244_ = d_2335_v28_
                lhs245_ = default__.safeIndex(703, (d_2335_v28_).length(0))
                lhs246_ = self
                d_2303_v8_ = rhs334_
                d_2334_v27_ = rhs335_
                lhs244_[lhs245_] = rhs336_
                lhs246_.f7 = rhs337_
                (globalState).f0 = (default__.fm34(d_2296_cf7_, self.f7, globalState)) <= ((d_2302_v7_) + (_dafny.SeqWithoutIsStrInference([p1.f3, p1.f3])))
            d_2336_v29_: _dafny.Array
            nw365_ = _dafny.Array(None, 1)
            d_2336_v29_ = nw365_
            d_2337_v30_: C3
            nw366_ = C3()
            nw366_.ctor__(p1.f3, self.f3)
            d_2337_v30_ = nw366_
            index383_ = default__.safeIndex(224, (d_2336_v29_).length(0))
            (d_2336_v29_)[index383_] = d_2337_v30_
            d_2338_v31_: _dafny.Map
            d_2338_v31_ = _dafny.Map({len((self).f6): 128})
            d_2339_v32_: _dafny.MultiSet
            d_2339_v32_ = _dafny.MultiSet([d_2338_v31_])
            d_2340_v33_: D22
            d_2340_v33_ = D22_DC55(d_2337_v30_)
            index384_ = default__.safeIndex(224, (d_2336_v29_).length(0))
            (d_2336_v29_)[index384_] = (d_2337_v30_ if (d_2302_v7_) == (_dafny.SeqWithoutIsStrInference([((d_2301_v6_)[default__.fm21(d_2296_cf7_, d_2339_v32_, globalState)] if (default__.fm21(d_2296_cf7_, d_2339_v32_, globalState)) in (d_2301_v6_) else (self).f4)])) else (d_2337_v30_ if True else (d_2340_v33_).cf88))
        elif source29_.is_DC5:
            d_2341___mcc_h1_ = source29_.cf8
            d_2342___mcc_h2_ = source29_.cf9
            d_2343___mcc_h3_ = source29_.cf10
            d_2344___mcc_h4_ = source29_.cf11
            d_2345_cf11_ = d_2344___mcc_h4_
            d_2346_cf10_ = d_2343___mcc_h3_
            d_2347_cf9_ = d_2342___mcc_h2_
            d_2348_cf8_ = d_2341___mcc_h1_
            d_2349_v34_: _dafny.Array
            nw367_ = _dafny.Array(None, 18)
            d_2349_v34_ = nw367_
            (globalState).f0 = not((len(_dafny.SeqWithoutIsStrInference([d_2349_v34_, d_2349_v34_, d_2349_v34_, d_2349_v34_]))) != (self.f7))
            d_2350_v35_: C1
            nw368_ = C1()
            nw368_.ctor__((self).f6, (p1).f4)
            d_2350_v35_ = nw368_
            d_2345_cf11_ = d_2345_cf11_
            d_2351_v36_: _dafny.Seq
            d_2351_v36_ = _dafny.SeqWithoutIsStrInference([d_2347_cf9_, not(not(p1.f3))])
            d_2352_v37_: _dafny.Seq
            d_2352_v37_ = _dafny.SeqWithoutIsStrInference([False, (d_2351_v36_)[default__.safeIndex((self).fm3(globalState), len(d_2351_v36_))], (d_2345_cf11_) >= (d_2346_cf10_)])
            d_2351_v36_ = default__.fm34(default__.fm31(p1.f3, d_2345_cf11_, not(self.f3), _dafny.SeqWithoutIsStrInference([not((p1).f4), not((p1).f4), (d_2350_v35_).f13]), globalState), self.f7, globalState)
        elif source29_.is_DC6:
            if self.f3:
                d_2353_v38_: C0
                nw369_ = C0()
                nw369_.ctor__()
                d_2353_v38_ = nw369_
                (self).f7 = self.f7
                d_2354_v39_: C2
                nw370_ = C2()
                nw370_.ctor__(True, (self).f4, p1.f3, p1.f3)
                d_2354_v39_ = nw370_
                d_2355_v40_: _dafny.Seq
                d_2355_v40_ = _dafny.SeqWithoutIsStrInference([d_2354_v39_])
                d_2356_v41_: _dafny.Seq
                d_2356_v41_ = _dafny.SeqWithoutIsStrInference([self.f7, self.f7, len(d_2355_v40_), self.f7])
                d_2357_v42_: _dafny.Map
                d_2357_v42_ = _dafny.Map({d_2356_v41_: not((d_2354_v39_).f10)})
                d_2358_v43_: _dafny.Seq
                d_2358_v43_ = _dafny.SeqWithoutIsStrInference([(self).f4, p1.f3, p1.f3])
                d_2359_v44_: _dafny.Set
                d_2359_v44_ = _dafny.Set({p1.f3, False})
                d_2360_v45_: _dafny.Map
                d_2360_v45_ = _dafny.Map({self.f3: p1.f3})
                d_2361_v46_: _dafny.Array
                nw371_ = _dafny.Array(None, 14)
                nw371_[int(0)] = self.f7
                nw371_[int(1)] = (self.f7 if default__.fm4(False, globalState) else self.f7)
                nw371_[int(2)] = self.f7
                nw371_[int(3)] = self.f7
                nw371_[int(4)] = self.f7
                nw371_[int(5)] = default__.fm31(((d_2357_v42_)[d_2356_v41_] if (d_2356_v41_) in (d_2357_v42_) else (d_2358_v43_)[default__.safeIndex(self.f7, len(d_2358_v43_))]), self.f7, (self).f4, d_2358_v43_, globalState)
                nw371_[int(6)] = len(d_2358_v43_)
                nw371_[int(7)] = len(d_2359_v44_)
                nw371_[int(8)] = self.f7
                nw371_[int(9)] = (self.f7) * (self.f7)
                nw371_[int(10)] = default__.safeModuloInt(self.f7, len(d_2360_v45_))
                nw371_[int(11)] = -452
                nw371_[int(12)] = self.f7
                nw371_[int(13)] = len(d_2356_v41_)
                d_2361_v46_ = nw371_
                d_2362_v47_: _dafny.Set
                d_2362_v47_ = _dafny.Set({self.f7})
                d_2363_v48_: _dafny.Seq
                d_2363_v48_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6, (self).f6, (self).f6, _dafny.SeqWithoutIsStrInference([p0 for d_2364_i5_ in range(default__.abs(719))])])
                rhs338_ = (d_2354_v39_).fm0(globalState)
                rhs339_ = (self.f7 if (p1).f4 else len(d_2362_v47_))
                rhs340_ = (len((d_2363_v48_)[default__.safeIndex(self.f7, len(d_2363_v48_))])) * (self.f7)
                rhs341_ = d_2361_v46_
                lhs247_ = self
                lhs248_ = self
                lhs249_ = self
                lhs247_.f7 = rhs338_
                lhs248_.f7 = rhs339_
                lhs249_.f7 = rhs340_
                d_2361_v46_ = rhs341_
                d_2365_v49_: _dafny.Map
                d_2365_v49_ = _dafny.Map({(self).f4: self.f7})
                d_2366_v50_: D0
                d_2366_v50_ = D0_DC2((self).f4, (d_2356_v41_)[default__.safeIndex(len(d_2365_v49_), len(d_2356_v41_))], self.f7)
                d_2367_v51_: _dafny.MultiSet
                d_2367_v51_ = _dafny.MultiSet([(d_2366_v50_).cf5, 341, len(d_2358_v43_)])
                d_2367_v51_ = _dafny.MultiSet([self.f7])
                d_2368_v52_: _dafny.Array
                nw372_ = _dafny.Array(_dafny.Map({}), 10)
                d_2368_v52_ = nw372_
                d_2369_v53_: _dafny.Map
                d_2369_v53_ = _dafny.Map({self.f7: p1.f3})
                d_2370_v54_: _dafny.Map
                d_2370_v54_ = _dafny.Map({d_2369_v53_: p1.f3})
                index385_ = default__.safeIndex(960, (d_2368_v52_).length(0))
                (d_2368_v52_)[index385_] = (_dafny.Map({default__.fm15(self.f3, self.f7, self.f7, p1.f3, globalState): (self).f4})) | (d_2370_v54_)
                d_2371_v56_: _dafny.Map
                d_2371_v56_ = _dafny.Map({self.f7: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_2372_i6_ in range(default__.abs(692))]))})
                index386_ = default__.safeIndex(960, (d_2368_v52_).length(0))
                def iife210_():
                    coll78_ = _dafny.Map()
                    compr_78_: _dafny.Map
                    for compr_78_ in ((_dafny.Map({_dafny.Map({self.f7: (p1).f4}): self.f7})).set(d_2369_v53_, len(d_2371_v56_))).keys.Elements:
                        d_2373_v55_: _dafny.Map = compr_78_
                        if (d_2373_v55_) in ((_dafny.Map({_dafny.Map({self.f7: (p1).f4}): self.f7})).set(d_2369_v53_, len(d_2371_v56_))):
                            coll78_[d_2373_v55_] = not(False)
                    return _dafny.Map(coll78_)
                (d_2368_v52_)[index386_] = iife210_()
                
            elif True:
                d_2374_v57_: _dafny.Seq
                d_2374_v57_ = _dafny.SeqWithoutIsStrInference([(p1).f4])
                (self).f3 = not((d_2374_v57_) <= (d_2374_v57_))
                rhs342_ = p1.f3
                rhs343_ = self.f7
                lhs250_ = globalState
                lhs251_ = self
                lhs250_.f0 = rhs342_
                lhs251_.f7 = rhs343_
                d_2375_v58_: _dafny.Map
                d_2375_v58_ = _dafny.Map({self.f7: (p1).f4})
                d_2376_v59_: _dafny.Map
                d_2376_v59_ = _dafny.Map({((d_2375_v58_)[(p1).fm0(globalState)] if ((p1).fm0(globalState)) in (d_2375_v58_) else p1.f3): 365})
                d_2377_v60_: D14
                d_2377_v60_ = D14_DC31(self.f7, self.f7)
                pat_let_tv68_ = globalState
                d_2378_v61_: _dafny.Map
                def iife211_(_pat_let66_0):
                    def iife212_(d_2379_dt__update__tmp_h0_):
                        def iife213_(_pat_let67_0):
                            def iife214_(d_2380_dt__update_hcf44_h0_):
                                return D14_DC31(d_2380_dt__update_hcf44_h0_, (d_2379_dt__update__tmp_h0_).cf45)
                            return iife214_(_pat_let67_0)
                        return iife213_((self).fm3(pat_let_tv68_))
                    return iife212_(_pat_let66_0)
                d_2378_v61_ = _dafny.Map({self.f7: ((d_2376_v59_)[(p1).f4] if ((p1).f4) in (d_2376_v59_) else (iife211_(d_2377_v60_)).cf44)})
                d_2381_v62_: _dafny.MultiSet
                d_2381_v62_ = _dafny.MultiSet([d_2378_v61_])
                d_2381_v62_ = d_2381_v62_
                (globalState).f0 = ((0) - (default__.safeDivisionInt(175, self.f7))) == (216)
                d_2382_v63_: _dafny.MultiSet
                d_2382_v63_ = _dafny.MultiSet([p1.f3])
                d_2383_v64_: str
                d_2383_v64_ = _dafny.CodePoint('r')
                rhs344_ = (d_2382_v63_).intersection(d_2382_v63_)
                rhs345_ = d_2383_v64_
                rhs346_ = self.f7
                rhs347_ = (p1).f4
                lhs252_ = self
                lhs253_ = globalState
                d_2382_v63_ = rhs344_
                d_2383_v64_ = rhs345_
                lhs252_.f7 = rhs346_
                lhs253_.f0 = rhs347_
            (self).f7 = self.f7
            d_2384_v65_: _dafny.MultiSet
            d_2384_v65_ = _dafny.MultiSet([self.f7])
            d_2384_v65_ = d_2384_v65_
            d_2385_v66_: T1
            nw373_ = C4()
            nw373_.ctor__(p1.f3, self.f7, (p1).f4, (p1).f4)
            d_2385_v66_ = nw373_
            d_2386_v67_: _dafny.Map
            d_2386_v67_ = _dafny.Map({self.f7: d_2385_v66_})
            d_2387_v68_: _dafny.Array
            nw374_ = _dafny.Array(None, 10)
            d_2387_v68_ = nw374_
            d_2388_v69_: _dafny.MultiSet
            d_2388_v69_ = _dafny.MultiSet([d_2387_v68_, d_2387_v68_, d_2387_v68_, d_2387_v68_, d_2387_v68_])
            d_2389_v70_: C2
            nw375_ = C2()
            nw375_.ctor__(self.f3, ((0) - (self.f7)) >= (len(_dafny.Set({self.f7, len(d_2386_v67_), self.f7}))), (d_2388_v69_).issubset(d_2388_v69_), not(default__.fm4(False, globalState)))
            d_2389_v70_ = nw375_
        elif True:
            d_2390___mcc_h5_ = source29_.cf6
            d_2391_cf6_ = d_2390___mcc_h5_
            index387_ = default__.safeIndex(752, (d_2391_cf6_).length(0))
            (d_2391_cf6_)[index387_] = (self).f4
            index388_ = default__.safeIndex(752, (d_2391_cf6_).length(0))
            (d_2391_cf6_)[index388_] = p1.f3
            d_2392_v71_: _dafny.Array
            nw376_ = _dafny.Array(int(0), 2)
            d_2392_v71_ = nw376_
            d_2393_v72_: _dafny.Set
            d_2393_v72_ = _dafny.Set({d_2392_v71_})
            (self).f7 = len((_dafny.Set({d_2392_v71_})).intersection(d_2393_v72_))
            d_2394_v73_: _dafny.Array
            nw377_ = _dafny.Array(None, 13)
            nw377_[int(0)] = ((self).f6) + ((self).f6)
            nw377_[int(1)] = ((self).f6).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urawnxuq"))), len((self).f6)), default__.fm23(self.f3, self.f7, globalState))
            nw377_[int(2)] = ((self).f6) + ((self).f6)
            nw377_[int(3)] = ((self).f6 if p1.f3 else (self).f6)
            nw377_[int(4)] = ((self).f6) + ((self).f6)
            nw377_[int(5)] = ((D17_DC38((self).f6)).cf56).set(default__.safeIndex(219, len((D17_DC38((self).f6)).cf56)), p0)
            nw377_[int(6)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "doqvtbv"))) + ((self).f6)
            nw377_[int(7)] = _dafny.SeqWithoutIsStrInference([p0 for d_2395_i7_ in range(default__.abs(603))])
            nw377_[int(8)] = (self).f6
            nw377_[int(9)] = _dafny.SeqWithoutIsStrInference([p0 for d_2396_i8_ in range(default__.abs(168))])
            nw377_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "porrudmcw"))
            nw377_[int(11)] = (self).f6
            nw377_[int(12)] = (self).f6
            d_2394_v73_ = nw377_
            index389_ = default__.safeIndex(747, (d_2394_v73_).length(0))
            (d_2394_v73_)[index389_] = (self).f6
            index390_ = default__.safeIndex(747, (d_2394_v73_).length(0))
            (d_2394_v73_)[index390_] = ((self).f6) + (_dafny.SeqWithoutIsStrInference([p0 for d_2397_i9_ in range(default__.abs(-63))]))
            d_2398_v74_: _dafny.Map
            d_2398_v74_ = _dafny.Map({p1.f3: 707})
            d_2399_v75_: _dafny.Seq
            d_2399_v75_ = _dafny.SeqWithoutIsStrInference([(p1).f4, not(True), not(self.f3)])
            d_2398_v74_ = (d_2398_v74_).set((p1).f4, len(d_2399_v75_))
        d_2400_v76_: _dafny.Array
        nw378_ = _dafny.Array(int(0), 25)
        d_2400_v76_ = nw378_
        d_2401_v77_: _dafny.Seq
        d_2401_v77_ = _dafny.SeqWithoutIsStrInference([d_2400_v76_])
        d_2402_v78_: _dafny.Map
        d_2402_v78_ = _dafny.Map({(self).f4: 810})
        if ((0) - (default__.safeDivisionInt(len(d_2401_v77_), self.f7))) < (default__.safeModuloInt(len(d_2402_v78_), self.f7)):
            d_2403_v79_: _dafny.Set
            d_2403_v79_ = _dafny.Set({self.f7})
            d_2404_v80_: D16
            d_2404_v80_ = D16_DC36(d_2403_v79_, p1.f3, False, not(default__.fm14(globalState)))
            d_2405_v81_: _dafny.Seq
            d_2405_v81_ = _dafny.SeqWithoutIsStrInference([True, (p1).f4])
            pat_let_tv69_ = d_2405_v81_
            pat_let_tv70_ = d_2405_v81_
            def iife215_(_pat_let68_0):
                def iife216_(d_2406_dt__update__tmp_h1_):
                    def iife217_(_pat_let69_0):
                        def iife218_(d_2407_dt__update_hcf53_h0_):
                            def iife219_(_pat_let70_0):
                                def iife220_(d_2408_dt__update_hcf51_h0_):
                                    return D16_DC36(d_2408_dt__update_hcf51_h0_, (d_2406_dt__update__tmp_h1_).cf52, d_2407_dt__update_hcf53_h0_, (d_2406_dt__update__tmp_h1_).cf54)
                                return iife220_(_pat_let70_0)
                            return iife219_(_dafny.Set({706}))
                        return iife218_(_pat_let69_0)
                    return iife217_((pat_let_tv69_)[default__.safeIndex(self.f7, len(pat_let_tv70_))])
                return iife216_(_pat_let68_0)
            source31_ = iife215_(d_2404_v80_)
            if source31_.is_DC36:
                d_2409___mcc_h10_ = source31_.cf51
                d_2410___mcc_h11_ = source31_.cf52
                d_2411___mcc_h12_ = source31_.cf53
                d_2412___mcc_h13_ = source31_.cf54
                d_2413_cf54_ = d_2412___mcc_h13_
                d_2414_cf53_ = d_2411___mcc_h12_
                d_2415_cf52_ = d_2410___mcc_h11_
                d_2416_cf51_ = d_2409___mcc_h10_
                d_2417_v82_: _dafny.Array
                nw379_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
                d_2417_v82_ = nw379_
                index391_ = default__.safeIndex(106, (d_2417_v82_).length(0))
                (d_2417_v82_)[index391_] = (self).f6
                index392_ = default__.safeIndex(106, (d_2417_v82_).length(0))
                (d_2417_v82_)[index392_] = (self).f6
                d_2400_v76_ = d_2400_v76_
                d_2400_v76_ = d_2400_v76_
                (self).m4(globalState)
            elif True:
                d_2418___mcc_h14_ = source31_.cf50
                d_2419_cf50_ = d_2418___mcc_h14_
                d_2420_v83_: _dafny.Array
                def lambda130_(d_2421_i10_):
                    return D12_DC28(D12_DC27(self.f7))

                init73_ = lambda130_
                nw380_ = _dafny.Array(None, 14)
                for i0_73_ in range(nw380_.length(0)):
                    nw380_[i0_73_] = init73_(i0_73_)
                d_2420_v83_ = nw380_
                d_2422_v84_: D12
                d_2422_v84_ = D12_DC27(self.f7)
                d_2423_v85_: D12
                d_2423_v85_ = D12_DC28(d_2422_v84_)
                d_2424_v86_: D12
                d_2424_v86_ = D12_DC28(d_2422_v84_)
                index393_ = default__.safeIndex(811, (d_2420_v83_).length(0))
                (d_2420_v83_)[index393_] = d_2424_v86_
                pat_let_tv71_ = d_2423_v85_
                index394_ = default__.safeIndex(811, (d_2420_v83_).length(0))
                def iife221_(_pat_let71_0):
                    def iife222_(d_2425_dt__update__tmp_h2_):
                        def iife223_(_pat_let72_0):
                            def iife224_(d_2426_dt__update_hcf41_h0_):
                                return D12_DC28(d_2426_dt__update_hcf41_h0_)
                            return iife224_(_pat_let72_0)
                        return iife223_(pat_let_tv71_)
                    return iife222_(_pat_let71_0)
                (d_2420_v83_)[index394_] = iife221_(d_2424_v86_)
                (globalState).f0 = (not(not((p1.f3) == (p1.f3))) if self.f3 else p1.f3)
                d_2427_v87_: D11
                d_2427_v87_ = D11_DC24(_dafny.Map({len(d_2405_v81_): default__.fm31(p1.f3, 175, self.f3, _dafny.SeqWithoutIsStrInference([(self).f4, (self).f4]), globalState)}))
                d_2428_v88_: _dafny.Map
                d_2428_v88_ = _dafny.Map({(0) - (self.f7): (d_2427_v87_).cf37})
                d_2429_v89_: _dafny.Array
                nw381_ = _dafny.Array(None, 11)
                d_2429_v89_ = nw381_
                d_2430_v90_: _dafny.Seq
                d_2430_v90_ = _dafny.SeqWithoutIsStrInference([d_2429_v89_])
                d_2431_v91_: _dafny.Map
                d_2431_v91_ = _dafny.Map({self.f7: self.f7})
                d_2432_v92_: _dafny.Map
                d_2432_v92_ = _dafny.Map({self.f7: p0})
                (self).f3 = default__.fm18(len(_dafny.Map({len(((d_2428_v88_)[len((d_2430_v90_).set(default__.safeIndex((0) - (default__.fm31((p1).f4, self.f7, True, d_2405_v81_, globalState)), len(d_2430_v90_)), d_2429_v89_))] if (len((d_2430_v90_).set(default__.safeIndex((0) - (default__.fm31((p1).f4, self.f7, True, d_2405_v81_, globalState)), len(d_2430_v90_)), d_2429_v89_))) in (d_2428_v88_) else d_2431_v91_)): len(d_2432_v92_)})), p0, self.f7, globalState)
                d_2433_v93_: _dafny.Map
                d_2433_v93_ = _dafny.Map({self.f7: False})
                (p1).f3 = (self.f7) >= ((0) - (len(d_2433_v93_)))
            index395_ = default__.safeIndex(336, (d_2400_v76_).length(0))
            (d_2400_v76_)[index395_] = self.f7
            d_2434_v95_: _dafny.Array
            def lambda131_(d_2435_i11_):
                def iife225_():
                    coll79_ = _dafny.Set()
                    compr_79_: int
                    for compr_79_ in _dafny.IntegerRange(351, 239):
                        d_2436_v94_: int = compr_79_
                        if ((351) <= (d_2436_v94_)) and ((d_2436_v94_) < (239)):
                            coll79_ = coll79_.union(_dafny.Set([default__.safeDivisionInt(d_2436_v94_, self.f7)]))
                    return _dafny.Set(coll79_)
                return iife225_()
                

            init74_ = lambda131_
            nw382_ = _dafny.Array(None, 15)
            for i0_74_ in range(nw382_.length(0)):
                nw382_[i0_74_] = init74_(i0_74_)
            d_2434_v95_ = nw382_
            d_2437_v96_: _dafny.Seq
            d_2437_v96_ = _dafny.SeqWithoutIsStrInference([d_2434_v95_])
            d_2438_v97_: D15
            d_2438_v97_ = D15_DC32(d_2434_v95_)
            d_2439_v98_: _dafny.MultiSet
            d_2439_v98_ = _dafny.MultiSet([D15_DC32((d_2437_v96_)[default__.safeIndex(self.f7, len(d_2437_v96_))]), d_2438_v97_, d_2438_v97_, d_2438_v97_])
            d_2440_v99_: _dafny.Map
            d_2440_v99_ = _dafny.Map({d_2439_v98_: (self.f7) < (self.f7)})
            index396_ = default__.safeIndex(336, (d_2400_v76_).length(0))
            rhs348_ = default__.safeModuloInt((len(d_2402_v78_)) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cpdfirkd"))))), len(default__.fm19(globalState)))
            rhs349_ = ((_dafny.Map({d_2439_v98_: (p1).f4})).set(d_2439_v98_, (p1).f4)) | (d_2440_v99_)
            lhs254_ = d_2400_v76_
            lhs255_ = default__.safeIndex(336, (d_2400_v76_).length(0))
            lhs254_[lhs255_] = rhs348_
            d_2440_v99_ = rhs349_
            d_2441_v100_: _dafny.Map
            d_2441_v100_ = _dafny.Map({(self).f6: (p1).f4})
            d_2441_v100_ = (d_2441_v100_).set(((self).f6) + ((self).f6), not(True))
            d_2291_v0_ = (d_2294_v1_).cf6
            d_2442_v101_: _dafny.MultiSet
            d_2442_v101_ = _dafny.MultiSet([self.f3])
            d_2443_v102_: _dafny.Map
            d_2443_v102_ = _dafny.Map({self.f7: (d_2442_v101_).cardinality})
            d_2444_v103_: _dafny.Map
            d_2444_v103_ = _dafny.Map({_dafny.Map({not((p1).f4): default__.fm18((d_2400_v76_)[default__.safeIndex(336, (d_2400_v76_).length(0))], p0, 5, globalState)}): d_2443_v102_})
            d_2444_v103_ = d_2444_v103_
        elif True:
            (self).f7 = ((d_2402_v78_)[p1.f3] if (p1.f3) in (d_2402_v78_) else (self.f7) * (self.f7))
            d_2445_v104_: D1
            d_2445_v104_ = D1_DC5((p1).f4, p1.f3, self.f7, self.f7)
            d_2446_v105_: _dafny.Seq
            d_2446_v105_ = _dafny.SeqWithoutIsStrInference([(self).f6])
            rhs350_ = (default__.fm30(d_2445_v104_, self.f7, self.f7, globalState)) != ((d_2446_v105_)[default__.safeIndex(self.f7, len(d_2446_v105_))])
            rhs351_ = d_2291_v0_
            rhs352_ = True
            rhs353_ = self.f7
            lhs256_ = p1
            lhs257_ = globalState
            lhs258_ = self
            lhs256_.f3 = rhs350_
            d_2291_v0_ = rhs351_
            lhs257_.f0 = rhs352_
            lhs258_.f7 = rhs353_
            if ((self.f7) == (len(_dafny.SeqWithoutIsStrInference([p1]))) if p1.f3 else (p1).f4):
                d_2447_v106_: _dafny.Array
                nw383_ = _dafny.Array(D17.default()(), 22)
                d_2447_v106_ = nw383_
                d_2447_v106_ = d_2447_v106_
                (self).m4(globalState)
                d_2448_v107_: _dafny.Map
                d_2448_v107_ = _dafny.Map({d_2291_v0_: (self.f7) - (self.f7)})
                d_2448_v107_ = (d_2448_v107_).set(d_2291_v0_, self.f7)
                d_2449_v108_: _dafny.Seq
                d_2449_v108_ = _dafny.SeqWithoutIsStrInference([p1.f3])
                d_2450_v109_: D4
                d_2450_v109_ = D4_DC13(D4_DC11(d_2449_v108_))
                d_2451_v110_: _dafny.MultiSet
                d_2451_v110_ = _dafny.MultiSet([d_2450_v109_, d_2450_v109_, d_2450_v109_, default__.fm43(self.f3, self.f7, (p1).f4, globalState), d_2450_v109_])
                d_2452_v111_: C2
                nw384_ = C2()
                nw384_.ctor__((d_2450_v109_) in (d_2451_v110_), ((self).f6) != ((default__.fm30(d_2445_v104_, 86, self.f7, globalState)).set(default__.safeIndex(self.f7, len(default__.fm30(d_2445_v104_, 86, self.f7, globalState))), _dafny.CodePoint('i'))), self.f3, (p1).f4)
                d_2452_v111_ = nw384_
                d_2453_v112_: _dafny.Map
                d_2453_v112_ = _dafny.Map({d_2400_v76_: (d_2452_v111_).f11})
                d_2454_v113_: _dafny.Set
                d_2454_v113_ = _dafny.Set({len(d_2453_v112_), self.f7})
                d_2454_v113_ = d_2454_v113_
            elif True:
                d_2455_v114_: C1
                nw385_ = C1()
                nw385_.ctor__((self).f6, self.f3)
                d_2455_v114_ = nw385_
                d_2456_v115_: _dafny.Map
                d_2456_v115_ = _dafny.Map({d_2455_v114_: (self).f6})
                d_2457_v116_: _dafny.Map
                d_2457_v116_ = _dafny.Map({True: (p1).f4})
                d_2458_v117_: _dafny.Map
                d_2458_v117_ = _dafny.Map({True: (d_2455_v114_).f12})
                d_2459_v118_: D17
                d_2459_v118_ = D17_DC38((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymgcfot"))).set(default__.safeIndex(self.f7, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymgcfot")))), p0))
                d_2460_v119_: _dafny.Array
                nw386_ = _dafny.Array(None, 29)
                nw386_[int(0)] = (self).f6
                nw386_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tufjnmcu"))
                nw386_[int(2)] = (self).f6
                nw386_[int(3)] = (self).f6
                nw386_[int(4)] = _dafny.SeqWithoutIsStrInference([p0 for d_2461_i12_ in range(default__.abs(616))])
                nw386_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "drlwyki"))
                nw386_[int(6)] = (((d_2456_v115_)[d_2455_v114_] if (d_2455_v114_) in (d_2456_v115_) else ((d_2455_v114_).f12).set(default__.safeIndex(len(d_2457_v116_), len((d_2455_v114_).f12)), _dafny.CodePoint('p')))) + ((self).f6)
                nw386_[int(7)] = _dafny.SeqWithoutIsStrInference([p0 for d_2462_i13_ in range(default__.abs(-178))])
                nw386_[int(8)] = ((self).f6) + ((d_2455_v114_).f12)
                nw386_[int(9)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iv"))) + ((d_2455_v114_).f12)
                nw386_[int(10)] = (self).f6
                nw386_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kagcscus"))
                nw386_[int(12)] = (self).f6
                nw386_[int(13)] = (d_2455_v114_).f12
                nw386_[int(14)] = (self).f6
                nw386_[int(15)] = ((d_2455_v114_).f12 if (self).f4 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldrf")))
                nw386_[int(16)] = ((d_2455_v114_).f12 if self.f3 else (self).f6)
                nw386_[int(17)] = (d_2455_v114_).f12
                nw386_[int(18)] = (d_2455_v114_).f12
                nw386_[int(19)] = (d_2455_v114_).f12
                nw386_[int(20)] = (d_2455_v114_).f12
                nw386_[int(21)] = (self).f6
                nw386_[int(22)] = (self).f6
                nw386_[int(23)] = ((d_2458_v117_)[(p1).f4] if ((p1).f4) in (d_2458_v117_) else (self).f6)
                nw386_[int(24)] = (self).f6
                nw386_[int(25)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfmg"))) + ((d_2459_v118_).cf56)
                nw386_[int(26)] = ((self).f6) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukbecno")))
                nw386_[int(27)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxn"))) + ((self).f6)
                nw386_[int(28)] = (self).f6
                d_2460_v119_ = nw386_
                index397_ = default__.safeIndex(279, (d_2460_v119_).length(0))
                (d_2460_v119_)[index397_] = (d_2455_v114_).f12
                index398_ = default__.safeIndex(279, (d_2460_v119_).length(0))
                (d_2460_v119_)[index398_] = ((_dafny.SeqWithoutIsStrInference([p0 for d_2463_i14_ in range(default__.abs(31))])) + ((d_2455_v114_).f12)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pnqsnqy"))) + ((self).f6))
                (p1).f3 = (p1).f4
                (p1).f3 = not(self.f3)
                d_2464_v120_: _dafny.Array
                nw387_ = _dafny.Array(None, 11)
                nw387_[int(0)] = d_2400_v76_
                nw387_[int(1)] = d_2400_v76_
                nw387_[int(2)] = d_2400_v76_
                nw387_[int(3)] = d_2400_v76_
                nw387_[int(4)] = (d_2401_v77_)[default__.safeIndex(self.f7, len(d_2401_v77_))]
                nw387_[int(5)] = d_2400_v76_
                nw387_[int(6)] = d_2400_v76_
                nw387_[int(7)] = d_2400_v76_
                nw387_[int(8)] = d_2400_v76_
                nw387_[int(9)] = d_2400_v76_
                nw387_[int(10)] = d_2400_v76_
                d_2464_v120_ = nw387_
                d_2465_v121_: D20
                d_2465_v121_ = D20_DC50(self.f7, (p1).fm0(globalState))
                d_2466_v122_: _dafny.MultiSet
                d_2466_v122_ = _dafny.MultiSet([p1.f3, (self).f4, not((p1).f4), p1.f3, p1.f3])
                rhs354_ = default__.safeDivisionInt(default__.fm31((self).f4, 172, (p1).f4, (default__.fm34(self.f7, 294, globalState)).set(default__.safeIndex(len((self).f6), len(default__.fm34(self.f7, 294, globalState))), p1.f3), globalState), ((d_2402_v78_)[(p1).f4] if ((p1).f4) in (d_2402_v78_) else self.f7))
                rhs355_ = (p1).f4
                rhs356_ = d_2464_v120_
                rhs357_ = ((d_2465_v121_).cf82) + ((self.f7) * (self.f7))
                rhs358_ = (((p1).f4) in (d_2466_v122_) if p1.f3 else (not(p1.f3) if not((d_2455_v114_).f13) else self.f3))
                lhs259_ = self
                lhs260_ = p1
                lhs261_ = self
                lhs262_ = self
                lhs259_.f7 = rhs354_
                lhs260_.f3 = rhs355_
                d_2464_v120_ = rhs356_
                lhs261_.f7 = rhs357_
                lhs262_.f3 = rhs358_
                (p1).f3 = (d_2455_v114_).f13
            (self).f7 = 91
            (self).f7 = self.f7
        (self).f7 = self.f7
        index399_ = default__.safeIndex(46, (d_2400_v76_).length(0))
        (d_2400_v76_)[index399_] = len((self).f6)
        index400_ = default__.safeIndex(46, (d_2400_v76_).length(0))
        (d_2400_v76_)[index400_] = default__.safeModuloInt(self.f7, 750)
        d_2467_v123_: _dafny.MultiSet
        d_2467_v123_ = _dafny.MultiSet([(d_2400_v76_)[default__.safeIndex(46, (d_2400_v76_).length(0))]])
        d_2468_v124_: C1
        nw388_ = C1()
        nw388_.ctor__((self).f6, (d_2467_v123_).issubset(d_2467_v123_))
        d_2468_v124_ = nw388_
        d_2469_v125_: _dafny.Seq
        d_2469_v125_ = _dafny.SeqWithoutIsStrInference([(d_2400_v76_)[default__.safeIndex(46, (d_2400_v76_).length(0))]])
        d_2470_v128_: _dafny.Array
        def lambda132_(d_2471_p0_):
            def lambda133_(d_2472_i15_):
                def iife226_():
                    def iife228_():
                        coll82_ = _dafny.Set()
                        compr_82_: str
                        for compr_82_ in (_dafny.Map({d_2471_p0_: self.f3})).keys.Elements:
                            d_2475_v126_: str = compr_82_
                            if (d_2475_v126_) in (_dafny.Map({d_2471_p0_: self.f3})):
                                coll82_ = coll82_.union(_dafny.Set([d_2475_v126_]))
                        return _dafny.Set(coll82_)
                    coll80_ = _dafny.Set()
                    def iife227_():
                        coll81_ = _dafny.Set()
                        compr_81_: str
                        for compr_81_ in (_dafny.Map({d_2471_p0_: self.f3})).keys.Elements:
                            d_2473_v126_: str = compr_81_
                            if (d_2473_v126_) in (_dafny.Map({d_2471_p0_: self.f3})):
                                coll81_ = coll81_.union(_dafny.Set([d_2473_v126_]))
                        return _dafny.Set(coll81_)
                    compr_80_: str
                    for compr_80_ in (iife227_()
                    ).Elements:
                        d_2474_v127_: str = compr_80_
                        if (d_2474_v127_) in (iife228_()
                        ):
                            coll80_ = coll80_.union(_dafny.Set([d_2474_v127_]))
                    return _dafny.Set(coll80_)
                return iife226_()
                

            return lambda133_

        init75_ = lambda132_(p0)
        nw389_ = _dafny.Array(None, 20)
        for i0_75_ in range(nw389_.length(0)):
            nw389_[i0_75_] = init75_(i0_75_)
        d_2470_v128_ = nw389_
        d_2476_v129_: _dafny.Seq
        d_2476_v129_ = _dafny.SeqWithoutIsStrInference([d_2470_v128_, d_2470_v128_])
        d_2477_v130_: _dafny.Map
        d_2477_v130_ = _dafny.Map({d_2469_v125_: (d_2476_v129_)[default__.safeIndex((d_2400_v76_)[default__.safeIndex(46, (d_2400_v76_).length(0))], len(d_2476_v129_))]})
        d_2477_v130_ = (d_2477_v130_).set(_dafny.SeqWithoutIsStrInference([(p1).fm0(globalState)]), d_2470_v128_)
        r0 = (d_2468_v124_).f12
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r3: int = int(0)
        d_2478_v0_: str
        d_2478_v0_ = _dafny.CodePoint('y')
        d_2479_v1_: _dafny.Seq
        d_2479_v1_ = _dafny.SeqWithoutIsStrInference([d_2478_v0_, d_2478_v0_, d_2478_v0_, d_2478_v0_])
        d_2480_v2_: _dafny.Map
        d_2480_v2_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference([d_2478_v0_, d_2478_v0_])})
        d_2479_v1_ = (d_2479_v1_) + (((d_2480_v2_)[self.f3] if (self.f3) in (d_2480_v2_) else p0))
        if (self).f4:
            d_2481_v3_: C3
            nw390_ = C3()
            nw390_.ctor__(p1, (self).f4)
            d_2481_v3_ = nw390_
            d_2482_v4_: _dafny.Array
            def lambda134_(d_2483_i0_):
                return (not((self).f4) if False else (self).f4)

            init76_ = lambda134_
            nw391_ = _dafny.Array(None, 17)
            for i0_76_ in range(nw391_.length(0)):
                nw391_[i0_76_] = init76_(i0_76_)
            d_2482_v4_ = nw391_
            d_2484_v5_: C0
            nw392_ = C0()
            nw392_.ctor__()
            d_2484_v5_ = nw392_
            d_2485_v6_: _dafny.Map
            d_2485_v6_ = _dafny.Map({138: p2})
            d_2486_v7_: _dafny.Map
            d_2486_v7_ = _dafny.Map({d_2484_v5_: len(d_2485_v6_)})
            d_2487_v8_: _dafny.Map
            d_2487_v8_ = _dafny.Map({self.f7: p3})
            d_2488_v9_: D11
            d_2488_v9_ = D11_DC24(d_2487_v8_)
            pat_let_tv72_ = d_2487_v8_
            d_2489_v10_: _dafny.MultiSet
            def iife229_(_pat_let73_0):
                def iife230_(d_2490_dt__update__tmp_h0_):
                    def iife231_(_pat_let74_0):
                        def iife232_(d_2491_dt__update_hcf37_h0_):
                            return D11_DC24(d_2491_dt__update_hcf37_h0_)
                        return iife232_(_pat_let74_0)
                    return iife231_(pat_let_tv72_)
                return iife230_(_pat_let73_0)
            d_2489_v10_ = _dafny.MultiSet([_dafny.Map({self.f7: ((d_2486_v7_)[d_2484_v5_] if (d_2484_v5_) in (d_2486_v7_) else len(_dafny.SeqWithoutIsStrInference([335])))}), (iife229_(d_2488_v9_)).cf37, d_2487_v8_])
            d_2492_v11_: _dafny.Seq
            d_2492_v11_ = _dafny.SeqWithoutIsStrInference([self.f7])
            d_2493_v12_: _dafny.Map
            d_2493_v12_ = _dafny.Map({d_2492_v11_: d_2478_v0_})
            d_2494_v13_: D17
            d_2494_v13_ = D17_DC37(d_2493_v12_)
            d_2495_v14_: _dafny.Seq
            d_2495_v14_ = _dafny.SeqWithoutIsStrInference([d_2494_v13_])
            d_2496_v16_: _dafny.Seq
            def iife233_():
                coll83_ = _dafny.Set()
                compr_83_: D17
                for compr_83_ in (d_2495_v14_).Elements:
                    d_2497_v15_: D17 = compr_83_
                    if (d_2497_v15_) in (d_2495_v14_):
                        coll83_ = coll83_.union(_dafny.Set([d_2497_v15_]))
                return _dafny.Set(coll83_)
            d_2496_v16_ = _dafny.SeqWithoutIsStrInference([iife233_()
            ])
            d_2498_v17_: _dafny.Seq
            d_2498_v17_ = _dafny.SeqWithoutIsStrInference([(d_2496_v16_)[default__.safeIndex(self.f7, len(d_2496_v16_))]])
            d_2499_v19_: _dafny.Map
            d_2499_v19_ = _dafny.Map({self.f3: d_2478_v0_})
            d_2500_v20_: _dafny.Map
            d_2500_v20_ = _dafny.Map({p3: d_2499_v19_})
            d_2501_v21_: _dafny.Set
            d_2501_v21_ = _dafny.Set({d_2494_v13_})
            d_2502_v23_: _dafny.Map
            def iife234_():
                coll84_ = _dafny.Set()
                compr_84_: _dafny.Set
                for compr_84_ in (_dafny.Map({d_2501_v21_: d_2478_v0_})).keys.Elements:
                    d_2503_v22_: _dafny.Set = compr_84_
                    if (d_2503_v22_) in (_dafny.Map({d_2501_v21_: d_2478_v0_})):
                        coll84_ = coll84_.union(_dafny.Set([d_2503_v22_]))
                return _dafny.Set(coll84_)
            d_2502_v23_ = _dafny.Map({d_2500_v20_: iife234_()
            })
            d_2504_v24_: D23
            d_2504_v24_ = D23_DC58(default__.fm44(self.f3, self.f3, globalState))
            d_2505_v25_: _dafny.Set
            d_2505_v25_ = _dafny.Set({(self).f6})
            nw393_ = _dafny.Array(None, 9)
            nw393_[int(0)] = self.f3
            nw393_[int(1)] = default__.fm21(self.f7, d_2489_v10_, globalState)
            nw393_[int(2)] = (self.f3 if p2 else self.f3)
            nw393_[int(3)] = (p0) <= ((self).f6)
            nw393_[int(4)] = (self).f4
            def iife235_():
                coll85_ = _dafny.Set()
                compr_85_: _dafny.Set
                for compr_85_ in (d_2498_v17_).Elements:
                    d_2506_v18_: _dafny.Set = compr_85_
                    if (d_2506_v18_) in (d_2498_v17_):
                        coll85_ = coll85_.union(_dafny.Set([d_2506_v18_]))
                return _dafny.Set(coll85_)
            nw393_[int(5)] = (((d_2502_v23_)[d_2500_v20_] if (d_2500_v20_) in (d_2502_v23_) else (d_2504_v24_).cf92)).issubset(iife235_()
            )
            nw393_[int(6)] = p2
            nw393_[int(7)] = (d_2505_v25_) != (d_2505_v25_)
            nw393_[int(8)] = (self).f4
            d_2482_v4_ = nw393_
            d_2507_v26_: _dafny.Set
            d_2507_v26_ = _dafny.Set({p1, p1})
            d_2508_v27_: _dafny.Map
            d_2508_v27_ = _dafny.Map({d_2507_v26_: self.f3})
            r3 = len((d_2508_v27_) | (_dafny.Map({_dafny.Set({True}): (self).f4})))
            d_2509_v28_: D0
            d_2509_v28_ = D0_DC2(p2, 999, (861 if True else p3))
            source32_ = d_2509_v28_
            if source32_.is_DC1:
                d_2510___mcc_h0_ = source32_.cf1
                d_2511___mcc_h1_ = source32_.cf2
                d_2512_cf2_ = d_2511___mcc_h1_
                d_2513_cf1_ = d_2510___mcc_h0_
                index401_ = default__.safeIndex(91, (d_2482_v4_).length(0))
                (d_2482_v4_)[index401_] = (self).f4
                d_2514_v29_: _dafny.Array
                def lambda135_(d_2515_p2_, d_2516_p0_, d_2517_p3_):
                    def lambda136_(d_2518_i1_):
                        return _dafny.Map({d_2515_p2_: _dafny.MultiSet([len(d_2516_p0_), d_2517_p3_])})

                    return lambda136_

                init77_ = lambda135_(p2, p0, p3)
                nw394_ = _dafny.Array(None, 5)
                for i0_77_ in range(nw394_.length(0)):
                    nw394_[i0_77_] = init77_(i0_77_)
                d_2514_v29_ = nw394_
                d_2519_v30_: _dafny.MultiSet
                d_2519_v30_ = _dafny.MultiSet([p3])
                d_2520_v31_: _dafny.Map
                d_2520_v31_ = _dafny.Map({(self).f4: d_2519_v30_})
                index402_ = default__.safeIndex(465, (d_2514_v29_).length(0))
                (d_2514_v29_)[index402_] = d_2520_v31_
                d_2521_v32_: _dafny.Array
                nw395_ = _dafny.Array(int(0), 8)
                d_2521_v32_ = nw395_
                d_2522_v33_: D17
                d_2522_v33_ = D17_DC39(d_2521_v32_, d_2513_cf1_, True, self.f7, p3)
                index403_ = default__.safeIndex(91, (d_2482_v4_).length(0))
                index404_ = default__.safeIndex(465, (d_2514_v29_).length(0))
                rhs359_ = self.f7
                rhs360_ = (self).f6
                rhs361_ = p1
                rhs362_ = (_dafny.Map({(self).f4: (_dafny.MultiSet([(d_2522_v33_).cf60])).set(p3, default__.abs(p3))})) | (d_2520_v31_)
                rhs363_ = d_2478_v0_
                lhs263_ = self
                lhs264_ = d_2482_v4_
                lhs265_ = default__.safeIndex(91, (d_2482_v4_).length(0))
                lhs266_ = d_2514_v29_
                lhs267_ = default__.safeIndex(465, (d_2514_v29_).length(0))
                lhs263_.f7 = rhs359_
                d_2479_v1_ = rhs360_
                lhs264_[lhs265_] = rhs361_
                lhs266_[lhs267_] = rhs362_
                d_2512_cf2_ = rhs363_
                (globalState).f0 = default__.fm18(p3, _dafny.CodePoint('v'), default__.fm31(self.f3, p3, False, _dafny.SeqWithoutIsStrInference([self.f3]), globalState), globalState)
                d_2523_v34_: _dafny.Seq
                d_2523_v34_ = _dafny.SeqWithoutIsStrInference([(d_2519_v30_).set(p3, default__.abs(p3))])
                d_2524_v35_: _dafny.Set
                d_2524_v35_ = _dafny.Set({self.f7})
                (self).f7 = (d_2484_v5_).fm10((self.f7) * (p3), (self).f4, (d_2523_v34_)[default__.safeIndex(len(d_2524_v35_), len(d_2523_v34_))], globalState)
                d_2525_v36_: C2
                nw396_ = C2()
                nw396_.ctor__(d_2513_cf1_, self.f3, not(p2), p1)
                d_2525_v36_ = nw396_
            elif source32_.is_DC2:
                d_2526___mcc_h2_ = source32_.cf3
                d_2527___mcc_h3_ = source32_.cf4
                d_2528___mcc_h4_ = source32_.cf5
                d_2529_cf5_ = d_2528___mcc_h4_
                d_2530_cf4_ = d_2527___mcc_h3_
                d_2531_cf3_ = d_2526___mcc_h2_
                rhs364_ = (len(d_2492_v11_)) + ((self.f7) * (self.f7))
                rhs365_ = self.f3
                rhs366_ = d_2482_v4_
                lhs268_ = self
                lhs269_ = self
                lhs268_.f7 = rhs364_
                lhs269_.f3 = rhs365_
                d_2482_v4_ = rhs366_
                d_2479_v1_ = _dafny.SeqWithoutIsStrInference([(D0_DC1((self).f4, _dafny.CodePoint('m'))).cf2 for d_2532_i2_ in range(default__.abs(828))])
                d_2533_v37_: _dafny.Map
                d_2533_v37_ = _dafny.Map({(d_2479_v1_)[default__.safeIndex(p3, len(d_2479_v1_))]: d_2478_v0_})
                d_2534_v39_: _dafny.Map
                d_2534_v39_ = _dafny.Map({(self).f4: default__.fm45(p3, self.f3, globalState)})
                d_2535_v40_: _dafny.Seq
                d_2535_v40_ = _dafny.SeqWithoutIsStrInference([default__.fm21(316, d_2489_v10_, globalState), p2, d_2531_cf3_, self.f3])
                d_2536_v41_: _dafny.Seq
                d_2536_v41_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({p3: len(d_2535_v40_)})) != (d_2487_v8_)])
                d_2537_v42_: _dafny.MultiSet
                d_2537_v42_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_2531_cf3_, not(d_2531_cf3_)]))])
                def iife236_():
                    coll86_ = _dafny.Map()
                    compr_86_: str
                    for compr_86_ in ((self).f6).Elements:
                        d_2538_v38_: str = compr_86_
                        if (d_2538_v38_) in ((self).f6):
                            coll86_[d_2538_v38_] = d_2478_v0_
                    return _dafny.Map(coll86_)
                rhs367_ = ((iife236_()
                ) | (d_2533_v37_)) | (((d_2534_v39_)[p1] if (p1) in (d_2534_v39_) else d_2533_v37_))
                rhs368_ = len(d_2536_v41_)
                rhs369_ = ((0) - ((d_2537_v42_).cardinality)) != (-14)
                lhs270_ = globalState
                d_2533_v37_ = rhs367_
                r2 = rhs368_
                lhs270_.f0 = rhs369_
                d_2539_v43_: _dafny.MultiSet
                d_2539_v43_ = _dafny.MultiSet([d_2531_cf3_])
                d_2540_v44_: _dafny.Map
                d_2540_v44_ = _dafny.Map({((d_2539_v43_).set((self).f4, default__.abs(565))).cardinality: d_2492_v11_})
                d_2529_cf5_ = (0) - (default__.safeDivisionInt((self.f7) * (len(d_2540_v44_)), (self).fm3(globalState)))
            elif True:
                d_2541___mcc_h5_ = source32_.cf0
                d_2542_cf0_ = d_2541___mcc_h5_
                d_2485_v6_ = (d_2485_v6_).set(self.f7, ((self).f6) != (d_2479_v1_))
                r1 = p1
                d_2542_cf0_ = self.f7
                d_2543_v45_: _dafny.Seq
                d_2543_v45_ = _dafny.SeqWithoutIsStrInference([False])
                d_2544_v46_: _dafny.MultiSet
                d_2544_v46_ = _dafny.MultiSet([p1, not((self).f4)])
                r1 = ((d_2544_v46_).set(p1, default__.abs(15))).ispropersubset(_dafny.MultiSet(d_2543_v45_))
            d_2545_v47_: _dafny.Array
            def lambda137_(d_2546_i3_):
                return default__.safeModuloInt(d_2546_i3_, self.f7)

            init78_ = lambda137_
            nw397_ = _dafny.Array(None, 20)
            for i0_78_ in range(nw397_.length(0)):
                nw397_[i0_78_] = init78_(i0_78_)
            d_2545_v47_ = nw397_
            index405_ = default__.safeIndex(412, (d_2545_v47_).length(0))
            (d_2545_v47_)[index405_] = p3
            d_2547_v48_: _dafny.Seq
            d_2547_v48_ = _dafny.SeqWithoutIsStrInference([not(self.f3), p1])
            d_2548_v49_: _dafny.Seq
            d_2548_v49_ = _dafny.SeqWithoutIsStrInference([((d_2547_v48_).set(default__.safeIndex(self.f7, len(d_2547_v48_)), (self).f4)) + (d_2547_v48_)])
            index406_ = default__.safeIndex(412, (d_2545_v47_).length(0))
            (d_2545_v47_)[index406_] = len((d_2548_v49_)[default__.safeIndex(p3, len(d_2548_v49_))])
        elif True:
            r2 = p3
            if (self.f7) != ((default__.fm20((self).f4, globalState)).cardinality):
                d_2549_v50_: _dafny.Seq
                d_2549_v50_ = _dafny.SeqWithoutIsStrInference([self.f3])
                d_2550_v51_: _dafny.Seq
                d_2550_v51_ = _dafny.SeqWithoutIsStrInference([d_2549_v50_])
                d_2551_v52_: _dafny.Map
                d_2551_v52_ = _dafny.Map({len(d_2550_v51_): (self).f4})
                d_2552_v53_: _dafny.Seq
                d_2552_v53_ = _dafny.SeqWithoutIsStrInference([d_2551_v52_, _dafny.Map({531: p1}), d_2551_v52_])
                d_2552_v53_ = d_2552_v53_
                (self).f3 = self.f3
                d_2553_v54_: _dafny.Array
                nw398_ = _dafny.Array(False, 21)
                d_2553_v54_ = nw398_
                index407_ = default__.safeIndex(911, (d_2553_v54_).length(0))
                (d_2553_v54_)[index407_] = not((p1) or ((self).f4))
                d_2554_v55_: _dafny.Array
                nw399_ = _dafny.Array(None, 23)
                d_2554_v55_ = nw399_
                d_2555_v56_: _dafny.MultiSet
                d_2555_v56_ = _dafny.MultiSet([d_2554_v55_, d_2554_v55_, d_2554_v55_])
                index408_ = default__.safeIndex(911, (d_2553_v54_).length(0))
                (d_2553_v54_)[index408_] = (d_2555_v56_).isdisjoint(_dafny.MultiSet([d_2554_v55_, d_2554_v55_]))
                d_2556_v57_: _dafny.Map
                d_2556_v57_ = _dafny.Map({p1: (self).fm3(globalState)})
                d_2557_v58_: _dafny.Set
                d_2557_v58_ = _dafny.Set({(d_2553_v54_)[default__.safeIndex(911, (d_2553_v54_).length(0))]})
                d_2558_v59_: _dafny.Map
                d_2558_v59_ = _dafny.Map({d_2557_v58_: d_2551_v52_})
                d_2556_v57_ = (d_2556_v57_).set(p2, (len(d_2558_v59_) if p2 else self.f7))
                d_2559_v60_: _dafny.Map
                d_2559_v60_ = _dafny.Map({p3: self.f7})
                d_2559_v60_ = ((d_2559_v60_ if not((d_2553_v54_)[default__.safeIndex(911, (d_2553_v54_).length(0))]) else d_2559_v60_)).set(p3, p3)
            elif True:
                d_2560_v61_: C0
                nw400_ = C0()
                nw400_.ctor__()
                d_2560_v61_ = nw400_
                d_2561_v62_: _dafny.MultiSet
                d_2561_v62_ = _dafny.MultiSet([d_2560_v61_])
                d_2562_v63_: T0
                nw401_ = C3()
                nw401_.ctor__((d_2561_v62_).issubset(d_2561_v62_), self.f3)
                d_2562_v63_ = nw401_
                d_2562_v63_ = d_2562_v63_
                d_2563_v64_: _dafny.Array
                nw402_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_2563_v64_ = nw402_
                d_2564_v65_: _dafny.Array
                nw403_ = _dafny.Array(_dafny.CodePoint('D'), 4)
                d_2564_v65_ = nw403_
                d_2565_v66_: _dafny.Map
                d_2565_v66_ = _dafny.Map({p1: d_2564_v65_})
                d_2566_v67_: D21
                d_2566_v67_ = D21_DC51(d_2564_v65_)
                d_2567_v68_: _dafny.Array
                nw404_ = _dafny.Array(None, 14)
                nw404_[int(0)] = d_2564_v65_
                nw404_[int(1)] = ((d_2565_v66_)[p1] if (p1) in (d_2565_v66_) else d_2564_v65_)
                nw404_[int(2)] = d_2564_v65_
                nw404_[int(3)] = d_2564_v65_
                nw404_[int(4)] = d_2564_v65_
                nw404_[int(5)] = d_2564_v65_
                nw404_[int(6)] = d_2564_v65_
                nw404_[int(7)] = d_2564_v65_
                nw404_[int(8)] = d_2564_v65_
                nw404_[int(9)] = d_2564_v65_
                nw404_[int(10)] = d_2564_v65_
                nw404_[int(11)] = d_2564_v65_
                nw404_[int(12)] = (d_2566_v67_).cf83
                nw404_[int(13)] = d_2564_v65_
                d_2567_v68_ = nw404_
                index409_ = default__.safeIndex(642, (d_2563_v64_).length(0))
                (d_2563_v64_)[index409_] = d_2567_v68_
                d_2568_v69_: D21
                d_2568_v69_ = D21_DC51(d_2564_v65_)
                d_2569_v70_: D21
                d_2569_v70_ = D21_DC54(d_2568_v69_)
                d_2570_v71_: _dafny.Seq
                d_2570_v71_ = _dafny.SeqWithoutIsStrInference([d_2562_v63_.f3])
                d_2571_v72_: _dafny.Seq
                d_2571_v72_ = _dafny.SeqWithoutIsStrInference([d_2568_v69_])
                pat_let_tv73_ = d_2568_v69_
                pat_let_tv74_ = d_2568_v69_
                pat_let_tv75_ = d_2568_v69_
                pat_let_tv76_ = d_2568_v69_
                d_2572_v73_: _dafny.Array
                nw405_ = _dafny.Array(None, 20)
                nw405_[int(0)] = d_2569_v70_
                def iife237_(_pat_let75_0):
                    def iife238_(d_2573_dt__update__tmp_h1_):
                        def iife239_(_pat_let76_0):
                            def iife240_(d_2574_dt__update_hcf87_h0_):
                                return D21_DC54(d_2574_dt__update_hcf87_h0_)
                            return iife240_(_pat_let76_0)
                        return iife239_(pat_let_tv73_)
                    return iife238_(_pat_let75_0)
                nw405_[int(1)] = iife237_(d_2569_v70_)
                nw405_[int(2)] = d_2569_v70_
                nw405_[int(3)] = (d_2569_v70_ if (self).f4 else d_2569_v70_)
                nw405_[int(4)] = d_2569_v70_
                nw405_[int(5)] = d_2569_v70_
                nw405_[int(6)] = d_2569_v70_
                def iife241_(_pat_let77_0):
                    def iife242_(d_2575_dt__update__tmp_h2_):
                        def iife243_(_pat_let78_0):
                            def iife244_(d_2576_dt__update_hcf87_h1_):
                                return D21_DC54(d_2576_dt__update_hcf87_h1_)
                            return iife244_(_pat_let78_0)
                        return iife243_(pat_let_tv74_)
                    return iife242_(_pat_let77_0)
                nw405_[int(7)] = (iife241_(D21_DC54(d_2568_v69_)) if d_2562_v63_.f3 else D21_DC54(d_2568_v69_))
                nw405_[int(8)] = (d_2569_v70_ if (d_2570_v71_)[default__.safeIndex(len(_dafny.Map({d_2562_v63_.f3: self.f3})), len(d_2570_v71_))] else d_2569_v70_)
                def iife245_(_pat_let79_0):
                    def iife246_(d_2577_dt__update__tmp_h3_):
                        def iife247_(_pat_let80_0):
                            def iife248_(d_2578_dt__update_hcf87_h2_):
                                return D21_DC54(d_2578_dt__update_hcf87_h2_)
                            return iife248_(_pat_let80_0)
                        return iife247_(pat_let_tv75_)
                    return iife246_(_pat_let79_0)
                nw405_[int(9)] = iife245_(D21_DC54(d_2568_v69_))
                nw405_[int(10)] = (D21_DC54(d_2568_v69_) if p2 else d_2569_v70_)
                nw405_[int(11)] = d_2569_v70_
                def iife249_(_pat_let81_0):
                    def iife250_(d_2579_dt__update__tmp_h4_):
                        def iife251_(_pat_let82_0):
                            def iife252_(d_2580_dt__update_hcf87_h3_):
                                return D21_DC54(d_2580_dt__update_hcf87_h3_)
                            return iife252_(_pat_let82_0)
                        return iife251_(pat_let_tv76_)
                    return iife250_(_pat_let81_0)
                nw405_[int(12)] = iife249_(d_2569_v70_)
                nw405_[int(13)] = d_2569_v70_
                nw405_[int(14)] = d_2569_v70_
                nw405_[int(15)] = D21_DC54(d_2568_v69_)
                nw405_[int(16)] = d_2569_v70_
                nw405_[int(17)] = (D21_DC54(d_2568_v69_) if self.f3 else D21_DC54(d_2568_v69_))
                nw405_[int(18)] = D21_DC54((d_2571_v72_)[default__.safeIndex(self.f7, len(d_2571_v72_))])
                nw405_[int(19)] = d_2569_v70_
                d_2572_v73_ = nw405_
                index410_ = default__.safeIndex(343, (d_2572_v73_).length(0))
                (d_2572_v73_)[index410_] = d_2569_v70_
                d_2581_v74_: _dafny.Array
                nw406_ = _dafny.Array(int(0), 11)
                d_2581_v74_ = nw406_
                d_2582_v75_: _dafny.Seq
                d_2582_v75_ = _dafny.SeqWithoutIsStrInference([d_2581_v74_])
                index411_ = default__.safeIndex(642, (d_2563_v64_).length(0))
                index412_ = default__.safeIndex(343, (d_2572_v73_).length(0))
                rhs370_ = False
                rhs371_ = (0) - (default__.safeDivisionInt(p3, self.f7))
                rhs372_ = d_2567_v68_
                rhs373_ = len(d_2582_v75_)
                rhs374_ = d_2569_v70_
                lhs271_ = d_2563_v64_
                lhs272_ = default__.safeIndex(642, (d_2563_v64_).length(0))
                lhs273_ = self
                lhs274_ = d_2572_v73_
                lhs275_ = default__.safeIndex(343, (d_2572_v73_).length(0))
                r1 = rhs370_
                r2 = rhs371_
                lhs271_[lhs272_] = rhs372_
                lhs273_.f7 = rhs373_
                lhs274_[lhs275_] = rhs374_
                d_2583_v76_: _dafny.Map
                d_2583_v76_ = _dafny.Map({(self).f4: (d_2562_v63_).f4})
                d_2583_v76_ = d_2583_v76_
                d_2584_v77_: _dafny.Set
                d_2584_v77_ = _dafny.Set({not((self).f4)})
                (d_2562_v63_).f3 = (d_2584_v77_).ispropersubset(d_2584_v77_)
                d_2478_v0_ = d_2478_v0_
            d_2585_v78_: _dafny.Map
            d_2585_v78_ = _dafny.Map({len(d_2479_v1_): p2})
            d_2586_v79_: D11
            d_2586_v79_ = D11_DC25(d_2585_v78_)
            pat_let_tv77_ = d_2585_v78_
            pat_let_tv78_ = d_2585_v78_
            def iife254_(_pat_let84_0):
                def iife255_(d_2587_dt__update__tmp_h5_):
                    def iife256_(_pat_let85_0):
                        def iife257_(d_2588_dt__update_hcf38_h0_):
                            return D11_DC25(d_2588_dt__update_hcf38_h0_)
                        return iife257_(_pat_let85_0)
                    return iife256_(pat_let_tv77_)
                return iife255_(_pat_let84_0)
            def iife253_(_pat_let83_0):
                def iife258_(d_2589_dt__update__tmp_h6_):
                    def iife259_(_pat_let86_0):
                        def iife260_(d_2590_dt__update_hcf38_h1_):
                            return D11_DC25(d_2590_dt__update_hcf38_h1_)
                        return iife260_(_pat_let86_0)
                    return iife259_(pat_let_tv78_)
                return iife258_(_pat_let83_0)
            d_2586_v79_ = iife253_(iife254_(d_2586_v79_))
            d_2591_v80_: D21
            d_2591_v80_ = D21_DC53(self.f3)
            d_2592_v81_: _dafny.Seq
            d_2592_v81_ = _dafny.SeqWithoutIsStrInference([(d_2591_v80_).cf86, p1, self.f3, p2])
            d_2593_v82_: _dafny.Set
            d_2593_v82_ = _dafny.Set({(self).f4})
            d_2594_v83_: _dafny.MultiSet
            d_2594_v83_ = _dafny.MultiSet([_dafny.Set({self.f3, (self).f4})])
            d_2595_v84_: _dafny.Seq
            d_2595_v84_ = _dafny.SeqWithoutIsStrInference([d_2594_v83_, d_2594_v83_])
            rhs375_ = default__.fm31(p1, len((d_2479_v1_) + (p0)), False, d_2592_v81_, globalState)
            rhs376_ = d_2478_v0_
            rhs377_ = (d_2593_v82_) in ((d_2594_v83_ if (self).f4 else (d_2595_v84_)[default__.safeIndex(p3, len(d_2595_v84_))]))
            rhs378_ = p2
            rhs379_ = default__.safeModuloInt(p3, (self).fm0(globalState))
            lhs276_ = self
            lhs277_ = globalState
            lhs276_.f7 = rhs375_
            d_2478_v0_ = rhs376_
            lhs277_.f0 = rhs377_
            r1 = rhs378_
            r3 = rhs379_
            (globalState).f0 = (d_2478_v0_) in (p0)
        (self).f7 = len((p0) + ((d_2479_v1_) + (p0)))
        (self).m4(globalState)
        d_2596_v85_: D8
        d_2596_v85_ = D8_DC20(d_2478_v0_, (self).f4)
        d_2597_v86_: C4
        nw407_ = C4()
        nw407_.ctor__((p2) and (True), -394, (d_2596_v85_).cf34, default__.fm6(54, globalState))
        d_2597_v86_ = nw407_
        (self).f3 = d_2597_v86_.f8
        r0 = self.f3
        r1 = d_2597_v86_.f8
        r2 = p3
        d_2598_v87_: _dafny.Array
        nw408_ = _dafny.Array(False, 10)
        d_2598_v87_ = nw408_
        d_2599_v88_: _dafny.MultiSet
        d_2599_v88_ = _dafny.MultiSet([p1, d_2597_v86_.f8])
        d_2600_v89_: _dafny.Map
        d_2600_v89_ = _dafny.Map({d_2598_v87_: d_2599_v88_})
        r3 = (0) - ((((d_2600_v89_)[(d_2598_v87_ if p2 else d_2598_v87_)] if ((d_2598_v87_ if p2 else d_2598_v87_)) in (d_2600_v89_) else (d_2599_v88_) - (d_2599_v88_))).cardinality)
        return r0, r1, r2, r3

    def m4(self, globalState):
        d_2601_v0_: C0
        nw409_ = C0()
        nw409_.ctor__()
        d_2601_v0_ = nw409_
        d_2602_v1_: D2
        d_2602_v1_ = D2_DC7((_dafny.MultiSet([d_2601_v0_, d_2601_v0_])).set(d_2601_v0_, default__.abs(self.f7)))
        d_2603_v2_: _dafny.Map
        d_2603_v2_ = _dafny.Map({d_2602_v1_: (self.f7) >= (len(_dafny.SeqWithoutIsStrInference([self.f3])))})
        d_2603_v2_ = (d_2603_v2_).set(d_2602_v1_, (self).f4)
        hi14_ = self.f7
        for d_2604_i0_ in range(self.f7, hi14_):
            d_2605_v3_: _dafny.Array
            nw410_ = _dafny.Array(int(0), 2)
            d_2605_v3_ = nw410_
            index413_ = default__.safeIndex(675, (d_2605_v3_).length(0))
            (d_2605_v3_)[index413_] = d_2604_i0_
            d_2606_v4_: C4
            nw411_ = C4()
            nw411_.ctor__(self.f3, len(_dafny.Set({d_2604_i0_})), not((self).f4), self.f3)
            d_2606_v4_ = nw411_
            d_2607_v5_: _dafny.Map
            d_2607_v5_ = _dafny.Map({d_2606_v4_: (d_2606_v4_).f9})
            d_2608_v6_: _dafny.Set
            d_2608_v6_ = _dafny.Set({d_2607_v5_})
            d_2609_v7_: D24
            d_2609_v7_ = D24_DC62(d_2608_v6_)
            index414_ = default__.safeIndex(675, (d_2605_v3_).length(0))
            rhs380_ = ((d_2608_v6_) - (d_2608_v6_)) != ((d_2609_v7_).cf99)
            rhs381_ = len(_dafny.SeqWithoutIsStrInference([((0) - (len(_dafny.Set({self.f7, 153})))) * (-334) for d_2610_i1_ in range(default__.abs(-804))]))
            lhs278_ = self
            lhs279_ = d_2605_v3_
            lhs280_ = default__.safeIndex(675, (d_2605_v3_).length(0))
            lhs278_.f3 = rhs380_
            lhs279_[lhs280_] = rhs381_
            d_2611_v8_: str
            d_2611_v8_ = _dafny.CodePoint('x')
            d_2612_v9_: _dafny.MultiSet
            d_2612_v9_ = _dafny.MultiSet([(d_2606_v4_).f9])
            d_2613_v10_: _dafny.Seq
            d_2613_v10_ = _dafny.SeqWithoutIsStrInference([(d_2612_v9_).cardinality])
            d_2614_v11_: _dafny.Seq
            d_2614_v11_ = _dafny.SeqWithoutIsStrInference([d_2606_v4_.f8, d_2606_v4_.f8])
            index415_ = default__.safeIndex(675, (d_2605_v3_).length(0))
            rhs382_ = ((d_2604_i0_ if not(self.f3) else 793)) + ((self.f7) - (len(d_2613_v10_)))
            rhs383_ = (self).f4
            rhs384_ = d_2611_v8_
            rhs385_ = default__.safeDivisionInt((len(d_2614_v11_)) + (d_2604_i0_), d_2604_i0_)
            lhs281_ = d_2605_v3_
            lhs282_ = default__.safeIndex(675, (d_2605_v3_).length(0))
            lhs283_ = self
            lhs284_ = self
            lhs281_[lhs282_] = rhs382_
            lhs283_.f3 = rhs383_
            d_2611_v8_ = rhs384_
            lhs284_.f7 = rhs385_
            d_2615_v12_: _dafny.Array
            nw412_ = _dafny.Array(D16.default()(), 4)
            d_2615_v12_ = nw412_
            d_2616_v13_: _dafny.Map
            d_2616_v13_ = _dafny.Map({self.f3: d_2604_i0_})
            d_2617_v14_: D16
            d_2617_v14_ = D16_DC35(d_2616_v13_)
            index416_ = default__.safeIndex(924, (d_2615_v12_).length(0))
            (d_2615_v12_)[index416_] = d_2617_v14_
            index417_ = default__.safeIndex(924, (d_2615_v12_).length(0))
            (d_2615_v12_)[index417_] = d_2617_v14_
            (self).f7 = (D14_DC31(len((default__.fm46(globalState)).set(d_2606_v4_.f8, (self).f6)), len(d_2614_v11_))).cf45
        d_2618_v15_: _dafny.Map
        d_2618_v15_ = _dafny.Map({self.f7: self.f3})
        d_2619_v16_: _dafny.Array
        nw413_ = _dafny.Array(None, 1)
        nw413_[int(0)] = len(d_2618_v15_)
        d_2619_v16_ = nw413_
        d_2620_v17_: _dafny.MultiSet
        d_2620_v17_ = _dafny.MultiSet([not(self.f3)])
        d_2621_v18_: _dafny.Map
        d_2621_v18_ = _dafny.Map({self.f7: self.f7})
        d_2622_v19_: _dafny.MultiSet
        d_2622_v19_ = _dafny.MultiSet([d_2621_v18_])
        d_2623_v20_: _dafny.Map
        d_2623_v20_ = _dafny.Map({(d_2619_v16_ if self.f3 else d_2619_v16_): len(_dafny.Set({d_2620_v17_, _dafny.MultiSet([(D0_DC2(self.f3, 491, len(d_2618_v15_))).cf3, default__.fm21(self.f7, d_2622_v19_, globalState), self.f3]), d_2620_v17_, d_2620_v17_, (d_2620_v17_).set(self.f3, default__.abs(self.f7))}))})
        d_2623_v20_ = (d_2623_v20_).set(d_2619_v16_, (self.f7) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2624_i2_ in range(default__.abs(976))]))))
        (self).f7 = self.f7
        d_2625_v21_: _dafny.Array
        nw414_ = _dafny.Array(None, 6)
        nw414_[int(0)] = d_2618_v15_
        nw414_[int(1)] = d_2618_v15_
        nw414_[int(2)] = d_2618_v15_
        nw414_[int(3)] = d_2618_v15_
        nw414_[int(4)] = d_2618_v15_
        nw414_[int(5)] = d_2618_v15_
        d_2625_v21_ = nw414_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_2625_v21_).length(0)):
            d_2626_i3_: int = guard_loop_6_
            if (True) and (((0) <= (d_2626_i3_)) and ((d_2626_i3_) < ((d_2625_v21_).length(0)))):
                (d_2625_v21_)[(d_2626_i3_)] = ((d_2618_v15_) | (d_2618_v15_)) | (d_2618_v15_)
        d_2627_v22_: C4
        nw415_ = C4()
        nw415_.ctor__(default__.fm21(self.f7, d_2622_v19_, globalState), 759, (self.f3 if not(self.f3) else (self).f4), default__.fm21(self.f7, (d_2622_v19_).set(d_2621_v18_, default__.abs(self.f7)), globalState))
        d_2627_v22_ = nw415_
        d_2627_v22_ = d_2627_v22_

    @property
    def f6(self):
        return self._f6

class C6(T0, T1):
    def  __init__(self):
        self._f3: bool = False
        self._f4: bool = False
        self.f5: bool = False
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
    def f4(self):
        return self._f4
    def ctor__(self, f5, f3, f4):
        (self).f5 = f5
        (self).f3 = f3
        (self)._f4 = f4

    def fm0(self, globalState):
        return len(_dafny.SeqWithoutIsStrInference([(0) - (-455)]))

    def fm1(self, p0, p1, p2, p3, globalState):
        return (0) - (default__.safeDivisionInt(default__.safeModuloInt(828, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))), 522))

    def fm2(self, p0, p1, p2, globalState):
        return len(_dafny.SeqWithoutIsStrInference([900, ((0) - (len(_dafny.SeqWithoutIsStrInference([694]))) if not(self.f3) else 643)]))

    def m0(self, p0, p1, p2, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        d_2628_v0_: _dafny.Seq
        d_2628_v0_ = _dafny.SeqWithoutIsStrInference([p2, -135])
        r1 = len(_dafny.Map({-472: (d_2628_v0_) + (d_2628_v0_)}))
        d_2629_v1_: _dafny.Seq
        d_2629_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxmev"))
        d_2630_v2_: _dafny.Seq
        d_2630_v2_ = _dafny.SeqWithoutIsStrInference([d_2629_v1_])
        d_2631_v3_: C5
        nw416_ = C5()
        nw416_.ctor__((d_2630_v2_)[default__.safeIndex(p2, len(d_2630_v2_))], (0) - (len((d_2629_v1_) + (d_2629_v1_))), (self).f4, (self).f4)
        d_2631_v3_ = nw416_
        hi15_ = p2
        for d_2632_i0_ in range(p2, hi15_):
            r1 = (d_2628_v0_)[default__.safeIndex(d_2631_v3_.f7, len(d_2628_v0_))]
            d_2633_v4_: _dafny.Array
            nw417_ = _dafny.Array(None, 1)
            nw417_[int(0)] = (self.f3) and (p0)
            d_2633_v4_ = nw417_
            index418_ = default__.safeIndex(931, (d_2633_v4_).length(0))
            (d_2633_v4_)[index418_] = self.f5
            d_2634_v5_: str
            d_2634_v5_ = _dafny.CodePoint('f')
            d_2635_v6_: _dafny.Map
            d_2635_v6_ = _dafny.Map({self.f3: self.f5})
            index419_ = default__.safeIndex(931, (d_2633_v4_).length(0))
            rhs386_ = (d_2631_v3_).f6
            rhs387_ = default__.fm18(996, d_2634_v5_, len(d_2635_v6_), globalState)
            lhs285_ = d_2633_v4_
            lhs286_ = default__.safeIndex(931, (d_2633_v4_).length(0))
            d_2629_v1_ = rhs386_
            lhs285_[lhs286_] = rhs387_
            if (self).f4:
                d_2636_v7_: _dafny.Array
                nw418_ = _dafny.Array(int(0), 24)
                d_2636_v7_ = nw418_
                index420_ = default__.safeIndex(196, (d_2636_v7_).length(0))
                (d_2636_v7_)[index420_] = d_2632_i0_
                index421_ = default__.safeIndex(196, (d_2636_v7_).length(0))
                (d_2636_v7_)[index421_] = ((p2) - (p2)) + (d_2631_v3_.f7)
                d_2637_v8_: _dafny.Array
                nw419_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_2637_v8_ = nw419_
                index422_ = default__.safeIndex(250, (d_2637_v8_).length(0))
                (d_2637_v8_)[index422_] = d_2636_v7_
                index423_ = default__.safeIndex(250, (d_2637_v8_).length(0))
                (d_2637_v8_)[index423_] = d_2636_v7_
                index424_ = default__.safeIndex(196, (d_2636_v7_).length(0))
                (d_2636_v7_)[index424_] = d_2631_v3_.f7
                d_2638_v9_: _dafny.Seq
                d_2638_v9_ = _dafny.SeqWithoutIsStrInference([d_2628_v0_, d_2628_v0_])
                d_2639_v10_: C4
                nw420_ = C4()
                nw420_.ctor__((d_2632_i0_) < ((self).fm1(not(self.f5), d_2631_v3_.f7, p0, d_2638_v9_, globalState)), (0) - (d_2632_i0_), True, (not((self).f4) if self.f5 else False))
                d_2639_v10_ = nw420_
                d_2639_v10_ = d_2639_v10_
                d_2640_v11_: _dafny.Map
                d_2640_v11_ = _dafny.Map({(d_2631_v3_.f7) * ((d_2636_v7_)[default__.safeIndex(196, (d_2636_v7_).length(0))]): d_2631_v3_.f7})
                d_2640_v11_ = (d_2640_v11_).set(default__.safeModuloInt((d_2636_v7_)[default__.safeIndex(196, (d_2636_v7_).length(0))], d_2631_v3_.f7), d_2631_v3_.f7)
            elif True:
                d_2641_v12_: C2
                nw421_ = C2()
                nw421_.ctor__(self.f3, (d_2633_v4_)[default__.safeIndex(931, (d_2633_v4_).length(0))], (self).f4, self.f5)
                d_2641_v12_ = nw421_
                d_2642_v13_: _dafny.Map
                d_2642_v13_ = _dafny.Map({d_2641_v12_: d_2634_v5_})
                d_2643_v14_: _dafny.Map
                d_2643_v14_ = _dafny.Map({((d_2642_v13_)[d_2641_v12_] if (d_2641_v12_) in (d_2642_v13_) else d_2634_v5_): (d_2641_v12_).f10})
                (self).f5 = (_dafny.Map({d_2634_v5_: False})) != (d_2643_v14_)
                d_2644_v15_: _dafny.Array
                def lambda138_(d_2645_v3_):
                    def lambda139_(d_2646_i1_):
                        return default__.safeModuloInt(d_2646_i1_, len((d_2645_v3_).f6))

                    return lambda139_

                init79_ = lambda138_(d_2631_v3_)
                nw422_ = _dafny.Array(None, 12)
                for i0_79_ in range(nw422_.length(0)):
                    nw422_[i0_79_] = init79_(i0_79_)
                d_2644_v15_ = nw422_
                r0 = d_2644_v15_
                d_2647_v16_: C0
                nw423_ = C0()
                nw423_.ctor__()
                d_2647_v16_ = nw423_
                d_2648_v17_: _dafny.Map
                d_2648_v17_ = _dafny.Map({d_2647_v16_: (d_2641_v12_).f10})
                d_2649_v18_: T0
                nw424_ = C4()
                nw424_.ctor__(((d_2633_v4_)[default__.safeIndex(931, (d_2633_v4_).length(0))]) == (p0), len(d_2648_v17_), (d_2641_v12_).f10, (d_2641_v12_).f10)
                d_2649_v18_ = nw424_
                d_2649_v18_ = d_2649_v18_
                d_2650_v19_: _dafny.MultiSet
                d_2650_v19_ = _dafny.MultiSet([d_2632_i0_, len(d_2628_v0_)])
                d_2651_v20_: C3
                nw425_ = C3()
                nw425_.ctor__(True, (d_2633_v4_)[default__.safeIndex(931, (d_2633_v4_).length(0))])
                d_2651_v20_ = nw425_
                d_2652_v21_: _dafny.Map
                d_2652_v21_ = _dafny.Map({d_2651_v20_: _dafny.MultiSet([d_2632_i0_, d_2631_v3_.f7, 858])})
                d_2653_v22_: C3
                nw426_ = C3()
                nw426_.ctor__((d_2633_v4_)[default__.safeIndex(931, (d_2633_v4_).length(0))], (((d_2652_v21_)[d_2651_v20_] if (d_2651_v20_) in (d_2652_v21_) else _dafny.MultiSet(d_2628_v0_))).issubset(d_2650_v19_))
                d_2653_v22_ = nw426_
                d_2654_v23_: _dafny.Map
                d_2654_v23_ = _dafny.Map({(d_2641_v12_).f10: (d_2628_v0_)[default__.safeIndex(p2, len(d_2628_v0_))]})
                d_2654_v23_ = (d_2654_v23_).set((p0) or ((d_2649_v18_).f4), 868)
            rhs388_ = (d_2631_v3_).f6
            rhs389_ = self.f3
            lhs287_ = globalState
            d_2629_v1_ = rhs388_
            lhs287_.f0 = rhs389_
        d_2655_v24_: _dafny.Set
        d_2655_v24_ = _dafny.Set({d_2631_v3_.f7, d_2631_v3_.f7, d_2631_v3_.f7, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnsnspud")))})
        d_2656_v25_: _dafny.Map
        d_2656_v25_ = _dafny.Map({(228) * (len(d_2655_v24_)): (0) - (p2)})
        d_2656_v25_ = d_2656_v25_
        r1 = default__.safeModuloInt(default__.safeDivisionInt(d_2631_v3_.f7, p2), (p2) - (d_2631_v3_.f7))
        d_2657_v26_: _dafny.Array
        nw427_ = _dafny.Array(None, 1)
        nw427_[int(0)] = p2
        d_2657_v26_ = nw427_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_2657_v26_).length(0)):
            d_2658_i2_: int = guard_loop_7_
            if (True) and (((0) <= (d_2658_i2_)) and ((d_2658_i2_) < ((d_2657_v26_).length(0)))):
                (d_2657_v26_)[(d_2658_i2_)] = (d_2658_i2_) * (d_2631_v3_.f7)
        def lambda140_(d_2659_v0_):
            def lambda141_(d_2660_i3_):
                return (d_2660_i3_) * (len(d_2659_v0_))

            return lambda141_

        init80_ = lambda140_(d_2628_v0_)
        nw428_ = _dafny.Array(None, 11)
        for i0_80_ in range(nw428_.length(0)):
            nw428_[i0_80_] = init80_(i0_80_)
        r0 = nw428_
        r1 = (0) - (d_2631_v3_.f7)
        return r0, r1

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        d_2661_v0_: _dafny.Set
        d_2661_v0_ = _dafny.Set({p0, p0})
        d_2662_v1_: D16
        d_2662_v1_ = D16_DC36(d_2661_v0_, (p1).f4, p1.f3, (p1).f4)
        pat_let_tv79_ = p0
        pat_let_tv80_ = p1
        d_2663_v2_: _dafny.Set
        def iife261_(_pat_let87_0):
            def iife262_(d_2664_dt__update__tmp_h0_):
                def iife263_(_pat_let88_0):
                    def iife264_(d_2665_dt__update_hcf51_h0_):
                        def iife265_(_pat_let89_0):
                            def iife266_(d_2666_dt__update_hcf53_h0_):
                                def iife267_(_pat_let90_0):
                                    def iife268_(d_2667_dt__update_hcf52_h0_):
                                        return D16_DC36(d_2665_dt__update_hcf51_h0_, d_2667_dt__update_hcf52_h0_, d_2666_dt__update_hcf53_h0_, (d_2664_dt__update__tmp_h0_).cf54)
                                    return iife268_(_pat_let90_0)
                                return iife267_((pat_let_tv80_).f4)
                            return iife266_(_pat_let89_0)
                        return iife265_(self.f3)
                    return iife264_(_pat_let88_0)
                return iife263_(_dafny.Set({pat_let_tv79_}))
            return iife262_(_pat_let87_0)
        d_2663_v2_ = _dafny.Set({iife261_(d_2662_v1_)})
        d_2668_v3_: _dafny.Map
        d_2668_v3_ = _dafny.Map({d_2663_v2_: False})
        (p1).f3 = ((d_2668_v3_)[d_2663_v2_] if (d_2663_v2_) in (d_2668_v3_) else p1.f3)
        d_2669_v4_: _dafny.Map
        d_2669_v4_ = _dafny.Map({default__.fm32(p1.f3, p1.f3, False, p1.f3, globalState): (p0 if p1.f3 else 746)})
        d_2670_v5_: _dafny.Map
        d_2670_v5_ = _dafny.Map({default__.fm4(self.f5, globalState): p1})
        d_2671_v6_: _dafny.Seq
        d_2671_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtokh"))
        d_2672_v7_: _dafny.MultiSet
        d_2672_v7_ = _dafny.MultiSet([p0, len(d_2670_v5_), len(d_2671_v6_)])
        d_2673_v8_: _dafny.Seq
        d_2673_v8_ = _dafny.SeqWithoutIsStrInference([len(d_2671_v6_), p0])
        d_2674_v9_: D20
        d_2674_v9_ = D20_DC48(d_2673_v8_)
        d_2675_v10_: str
        d_2675_v10_ = _dafny.CodePoint('x')
        d_2676_v11_: _dafny.Seq
        d_2676_v11_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2677_i0_ in range(default__.abs(-105))]), (d_2671_v6_).set(default__.safeIndex(p0, len(d_2671_v6_)), d_2675_v10_)])
        d_2669_v4_ = (_dafny.Map({d_2661_v0_: p0})) | (default__.fm47((d_2672_v7_).cardinality, self.f3, d_2674_v9_, (d_2676_v11_)[default__.safeIndex(p0, len(d_2676_v11_))], globalState))
        d_2671_v6_ = (d_2671_v6_).set(default__.safeIndex((0) - (p0), len(d_2671_v6_)), d_2675_v10_)
        d_2678_i1_: int
        d_2678_i1_ = 0
        with _dafny.label("9"):
            while p1.f3:
                with _dafny.c_label("9"):
                    if (d_2678_i1_) >= (100):
                        raise _dafny.Break("9")
                    d_2678_i1_ = (d_2678_i1_) + (1)
                    d_2679_v12_: _dafny.Seq
                    d_2679_v12_ = _dafny.SeqWithoutIsStrInference([self.f5])
                    d_2679_v12_ = (d_2679_v12_) + (d_2679_v12_)
                    if (p1).f4:
                        d_2680_v13_: _dafny.MultiSet
                        d_2680_v13_ = _dafny.MultiSet([d_2675_v10_, d_2675_v10_, d_2675_v10_, d_2675_v10_, d_2675_v10_])
                        d_2681_v14_: C4
                        nw429_ = C4()
                        nw429_.ctor__(not ((p1).f4) or ((p1).f4), ((d_2680_v13_)[d_2675_v10_] if (d_2675_v10_) in (d_2680_v13_) else p0), (p1).f4, self.f3)
                        d_2681_v14_ = nw429_
                        d_2682_v15_: C3
                        nw430_ = C3()
                        nw430_.ctor__(p1.f3, p1.f3)
                        d_2682_v15_ = nw430_
                        d_2683_v16_: _dafny.MultiSet
                        d_2683_v16_ = _dafny.MultiSet([True])
                        d_2684_v17_: _dafny.Map
                        d_2684_v17_ = _dafny.Map({d_2683_v16_: not(default__.fm4((p1).f4, globalState))})
                        d_2684_v17_ = (d_2684_v17_) | (d_2684_v17_)
                        d_2685_v18_: _dafny.Seq
                        d_2685_v18_ = _dafny.SeqWithoutIsStrInference([d_2679_v12_])
                        d_2686_v19_: _dafny.Array
                        nw431_ = _dafny.Array(None, 15)
                        nw431_[int(0)] = d_2679_v12_
                        nw431_[int(1)] = d_2679_v12_
                        nw431_[int(2)] = _dafny.SeqWithoutIsStrInference([(self).f4, True, (p1).f4, p1.f3, True])
                        nw431_[int(3)] = d_2679_v12_
                        nw431_[int(4)] = d_2679_v12_
                        nw431_[int(5)] = (d_2679_v12_) + (d_2679_v12_)
                        nw431_[int(6)] = d_2679_v12_
                        nw431_[int(7)] = d_2679_v12_
                        nw431_[int(8)] = d_2679_v12_
                        nw431_[int(9)] = (d_2679_v12_) + (d_2679_v12_)
                        nw431_[int(10)] = d_2679_v12_
                        nw431_[int(11)] = ((d_2679_v12_) + ((d_2685_v18_)[default__.safeIndex(p0, len(d_2685_v18_))])).set(default__.safeIndex(len(d_2679_v12_), len((d_2679_v12_) + ((d_2685_v18_)[default__.safeIndex(p0, len(d_2685_v18_))]))), p1.f3)
                        nw431_[int(12)] = d_2679_v12_
                        nw431_[int(13)] = d_2679_v12_
                        nw431_[int(14)] = _dafny.SeqWithoutIsStrInference([False])
                        d_2686_v19_ = nw431_
                        index425_ = default__.safeIndex(613, (d_2686_v19_).length(0))
                        (d_2686_v19_)[index425_] = (d_2679_v12_) + (d_2679_v12_)
                        index426_ = default__.safeIndex(613, (d_2686_v19_).length(0))
                        (d_2686_v19_)[index426_] = ((default__.fm34(p0, (d_2681_v14_).f9, globalState)) + (d_2679_v12_)) + (d_2679_v12_)
                        r0 = default__.safeModuloInt(p0, (p0) * (len(d_2671_v6_)))
                    elif True:
                        d_2687_v20_: _dafny.Map
                        d_2687_v20_ = _dafny.Map({len(d_2679_v12_): (p1).f4})
                        d_2688_v21_: _dafny.Array
                        d_2689_v22_: int
                        out47_: _dafny.Array
                        out48_: int
                        out47_, out48_ = (p1).m0((p1).f4, d_2687_v20_, default__.safeDivisionInt(p0, p0), globalState)
                        d_2688_v21_ = out47_
                        d_2689_v22_ = out48_
                        d_2690_v23_: _dafny.Map
                        d_2690_v23_ = _dafny.Map({d_2689_v22_: -272})
                        d_2691_v24_: _dafny.MultiSet
                        d_2691_v24_ = _dafny.MultiSet([self.f3])
                        (globalState).f0 = (D0_DC2(p1.f3, -562, len((d_2690_v23_).set(len(d_2673_v8_), (d_2691_v24_).cardinality)))).cf3
                        index427_ = default__.safeIndex(155, (d_2688_v21_).length(0))
                        (d_2688_v21_)[index427_] = d_2689_v22_
                        d_2692_v25_: _dafny.Seq
                        d_2692_v25_ = _dafny.SeqWithoutIsStrInference([d_2673_v8_])
                        d_2693_v26_: _dafny.Map
                        d_2693_v26_ = _dafny.Map({(0) - (p0): d_2692_v25_})
                        d_2694_v27_: _dafny.Seq
                        d_2694_v27_ = ((d_2693_v26_)[p0] if (p0) in (d_2693_v26_) else d_2692_v25_)
                        index428_ = default__.safeIndex(155, (d_2688_v21_).length(0))
                        rhs390_ = (self).fm1(default__.fm4(p1.f3, globalState), d_2689_v22_, ((d_2687_v20_)[d_2689_v22_] if (d_2689_v22_) in (d_2687_v20_) else p1.f3), (d_2694_v27_), globalState)
                        rhs391_ = (0) - ((((0) - (p0) if (p1).f4 else len(d_2673_v8_))) - (d_2689_v22_))
                        lhs288_ = d_2688_v21_
                        lhs289_ = default__.safeIndex(155, (d_2688_v21_).length(0))
                        d_2689_v22_ = rhs390_
                        lhs288_[lhs289_] = rhs391_
                        d_2695_v28_: D17
                        d_2695_v28_ = D17_DC39(d_2688_v21_, p1.f3, p1.f3, p0, (d_2688_v21_)[default__.safeIndex(155, (d_2688_v21_).length(0))])
                        index429_ = default__.safeIndex(155, (d_2688_v21_).length(0))
                        (d_2688_v21_)[index429_] = (d_2695_v28_).cf60
                        index430_ = default__.safeIndex(155, (d_2688_v21_).length(0))
                        (d_2688_v21_)[index430_] = p0
                    d_2673_v8_ = d_2673_v8_
                    d_2696_v29_: _dafny.MultiSet
                    d_2696_v29_ = _dafny.MultiSet([default__.fm6(p0, globalState), self.f3, p1.f3, self.f3])
                    (self).f5 = (d_2696_v29_).isdisjoint(_dafny.MultiSet([self.f3]))
                    pass
            pass
        d_2697_v30_: C2
        nw432_ = C2()
        nw432_.ctor__((self).f4, self.f5, p1.f3, not (self.f3) or (p1.f3))
        d_2697_v30_ = nw432_
        d_2698_v31_: C3
        nw433_ = C3()
        nw433_.ctor__(self.f5, (self).f4)
        d_2698_v31_ = nw433_
        d_2699_v32_: D22
        d_2699_v32_ = D22_DC55(d_2698_v31_)
        pat_let_tv81_ = d_2698_v31_
        pat_let_tv82_ = d_2698_v31_
        def iife270_(_pat_let92_0):
            def iife271_(d_2700_dt__update__tmp_h1_):
                def iife272_(_pat_let93_0):
                    def iife273_(d_2701_dt__update_hcf88_h0_):
                        return D22_DC55(d_2701_dt__update_hcf88_h0_)
                    return iife273_(_pat_let93_0)
                return iife272_(pat_let_tv81_)
            return iife271_(_pat_let92_0)
        def iife269_(_pat_let91_0):
            def iife274_(d_2702_dt__update__tmp_h2_):
                def iife275_(_pat_let94_0):
                    def iife276_(d_2703_dt__update_hcf88_h1_):
                        return D22_DC55(d_2703_dt__update_hcf88_h1_)
                    return iife276_(_pat_let94_0)
                return iife275_(pat_let_tv82_)
            return iife274_(_pat_let91_0)
        source33_ = iife269_(iife270_(d_2699_v32_))
        if source33_.is_DC56:
            d_2704___mcc_h0_ = source33_.cf89
            d_2705___mcc_h1_ = source33_.cf90
            d_2706_cf90_ = d_2705___mcc_h1_
            d_2707_cf89_ = d_2704___mcc_h0_
            d_2708_v33_: _dafny.Array
            nw434_ = _dafny.Array(int(0), 1)
            d_2708_v33_ = nw434_
            d_2709_v34_: _dafny.Seq
            d_2709_v34_ = _dafny.SeqWithoutIsStrInference([d_2708_v33_])
            d_2710_v35_: _dafny.Map
            d_2710_v35_ = _dafny.Map({d_2675_v10_: (d_2709_v34_)[default__.safeIndex(d_2707_cf89_, len(d_2709_v34_))]})
            d_2711_v36_: _dafny.Array
            nw435_ = _dafny.Array(None, 20)
            nw435_[int(0)] = d_2708_v33_
            nw435_[int(1)] = d_2708_v33_
            nw435_[int(2)] = d_2708_v33_
            nw435_[int(3)] = d_2708_v33_
            nw435_[int(4)] = d_2708_v33_
            nw435_[int(5)] = d_2708_v33_
            nw435_[int(6)] = d_2708_v33_
            nw435_[int(7)] = d_2708_v33_
            nw435_[int(8)] = d_2708_v33_
            nw435_[int(9)] = d_2708_v33_
            nw435_[int(10)] = d_2708_v33_
            nw435_[int(11)] = d_2708_v33_
            nw435_[int(12)] = d_2708_v33_
            nw435_[int(13)] = d_2708_v33_
            nw435_[int(14)] = d_2708_v33_
            nw435_[int(15)] = d_2708_v33_
            nw435_[int(16)] = ((d_2710_v35_)[d_2675_v10_] if (d_2675_v10_) in (d_2710_v35_) else d_2708_v33_)
            nw435_[int(17)] = d_2708_v33_
            nw435_[int(18)] = d_2708_v33_
            nw435_[int(19)] = d_2708_v33_
            d_2711_v36_ = nw435_
            index431_ = default__.safeIndex(265, (d_2711_v36_).length(0))
            (d_2711_v36_)[index431_] = d_2708_v33_
            d_2712_v37_: C0
            nw436_ = C0()
            nw436_.ctor__()
            d_2712_v37_ = nw436_
            d_2713_v38_: _dafny.Array
            nw437_ = _dafny.Array(None, 13)
            nw437_[int(0)] = d_2712_v37_
            nw437_[int(1)] = d_2712_v37_
            nw437_[int(2)] = d_2712_v37_
            nw437_[int(3)] = d_2712_v37_
            nw437_[int(4)] = d_2712_v37_
            nw437_[int(5)] = d_2712_v37_
            nw437_[int(6)] = d_2712_v37_
            nw437_[int(7)] = d_2712_v37_
            nw437_[int(8)] = d_2712_v37_
            nw437_[int(9)] = d_2712_v37_
            nw437_[int(10)] = d_2712_v37_
            nw437_[int(11)] = d_2712_v37_
            nw437_[int(12)] = d_2712_v37_
            d_2713_v38_ = nw437_
            d_2714_v39_: _dafny.Map
            d_2714_v39_ = _dafny.Map({(d_2697_v30_).f11: d_2713_v38_})
            d_2715_v40_: D19
            d_2715_v40_ = D19_DC44((d_2714_v39_) | ((_dafny.Map({p1.f3: d_2713_v38_})).set((p1).f4, d_2713_v38_)))
            d_2716_v41_: _dafny.Array
            nw438_ = _dafny.Array(False, 5)
            d_2716_v41_ = nw438_
            index432_ = default__.safeIndex(579, (d_2716_v41_).length(0))
            (d_2716_v41_)[index432_] = (d_2697_v30_).f11
            d_2717_v42_: D0
            d_2717_v42_ = D0_DC0(335)
            d_2718_v43_: _dafny.Set
            d_2718_v43_ = _dafny.Set({(d_2697_v30_).f11, p1.f3})
            index433_ = default__.safeIndex(265, (d_2711_v36_).length(0))
            index434_ = default__.safeIndex(579, (d_2716_v41_).length(0))
            rhs392_ = d_2708_v33_
            rhs393_ = _dafny.MultiSet([default__.safeDivisionInt(d_2707_cf89_, d_2707_cf89_)])
            rhs394_ = (len(d_2718_v43_) if default__.fm18(p0, d_2675_v10_, (d_2717_v42_).cf0, globalState) else len(_dafny.Set({default__.fm4((p1).f4, globalState), (p1).f4})))
            rhs395_ = d_2715_v40_
            rhs396_ = (self).f4
            lhs290_ = d_2711_v36_
            lhs291_ = default__.safeIndex(265, (d_2711_v36_).length(0))
            lhs292_ = d_2716_v41_
            lhs293_ = default__.safeIndex(579, (d_2716_v41_).length(0))
            lhs290_[lhs291_] = rhs392_
            d_2672_v7_ = rhs393_
            r0 = rhs394_
            d_2715_v40_ = rhs395_
            lhs292_[lhs293_] = rhs396_
            d_2719_v44_: _dafny.Array
            def lambda142_(d_2720_v10_):
                def lambda143_(d_2721_i2_):
                    return d_2720_v10_

                return lambda143_

            init81_ = lambda142_(d_2675_v10_)
            nw439_ = _dafny.Array(None, 24)
            for i0_81_ in range(nw439_.length(0)):
                nw439_[i0_81_] = init81_(i0_81_)
            d_2719_v44_ = nw439_
            index435_ = default__.safeIndex(813, (d_2719_v44_).length(0))
            (d_2719_v44_)[index435_] = d_2675_v10_
            index436_ = default__.safeIndex(813, (d_2719_v44_).length(0))
            (d_2719_v44_)[index436_] = d_2675_v10_
            index437_ = default__.safeIndex(813, (d_2719_v44_).length(0))
            (d_2719_v44_)[index437_] = d_2675_v10_
            arr0_ = (d_2711_v36_)[default__.safeIndex(265, (d_2711_v36_).length(0))]
            index438_ = default__.safeIndex(209, ((d_2711_v36_)[default__.safeIndex(265, (d_2711_v36_).length(0))]).length(0))
            arr0_[index438_] = len(d_2673_v8_)
            arr1_ = (d_2711_v36_)[default__.safeIndex(265, (d_2711_v36_).length(0))]
            index439_ = default__.safeIndex(209, ((d_2711_v36_)[default__.safeIndex(265, (d_2711_v36_).length(0))]).length(0))
            arr1_[index439_] = d_2707_cf89_
        elif source33_.is_DC55:
            d_2722___mcc_h2_ = source33_.cf88
            d_2723_cf88_ = d_2722___mcc_h2_
            d_2724_v45_: _dafny.Array
            def lambda144_(d_2725_v8_, d_2726_p0_, d_2727_v30_, d_2728_p1_):
                def lambda145_(d_2729_i3_):
                    return (D18_DC43((d_2725_v8_)[default__.safeIndex(d_2726_p0_, len(d_2725_v8_))], (D19_DC45(d_2726_p0_, 167, (d_2727_v30_).f10, self.f5, _dafny.SeqWithoutIsStrInference([self.f5]))).cf70, (d_2728_p1_).f4)).cf67

                return lambda145_

            init82_ = lambda144_(d_2673_v8_, p0, d_2697_v30_, p1)
            nw440_ = _dafny.Array(None, 18)
            for i0_82_ in range(nw440_.length(0)):
                nw440_[i0_82_] = init82_(i0_82_)
            d_2724_v45_ = nw440_
            d_2724_v45_ = d_2724_v45_
            d_2730_v46_: _dafny.Seq
            d_2730_v46_ = _dafny.SeqWithoutIsStrInference([(d_2697_v30_).f10, (self).f4, (p1).f4, not(self.f5), (p1).f4])
            d_2731_v47_: D18
            d_2731_v47_ = D18_DC43(default__.fm31(self.f3, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqlotnfij"))), self.f5, d_2730_v46_, globalState), (p0) - (p0), self.f3)
            d_2732_v48_: _dafny.Array
            nw441_ = _dafny.Array(int(0), 7)
            d_2732_v48_ = nw441_
            d_2733_v49_: D17
            d_2733_v49_ = D17_DC39(d_2732_v48_, self.f5, (d_2697_v30_).f11, p0, p0)
            pat_let_tv83_ = d_2697_v30_
            def iife277_(_pat_let95_0):
                def iife278_(d_2734_dt__update__tmp_h3_):
                    def iife279_(_pat_let96_0):
                        def iife280_(d_2735_dt__update_hcf58_h0_):
                            def iife281_(_pat_let97_0):
                                def iife282_(d_2736_dt__update_hcf59_h0_):
                                    return D17_DC39((d_2734_dt__update__tmp_h3_).cf57, d_2735_dt__update_hcf58_h0_, d_2736_dt__update_hcf59_h0_, (d_2734_dt__update__tmp_h3_).cf60, (d_2734_dt__update__tmp_h3_).cf61)
                                return iife282_(_pat_let97_0)
                            return iife281_((self).f4)
                        return iife280_(_pat_let96_0)
                    return iife279_((pat_let_tv83_).f10)
                return iife278_(_pat_let95_0)
            rhs397_ = not((iife277_(d_2733_v49_)).cf58)
            rhs398_ = d_2731_v47_
            lhs294_ = p1
            lhs294_.f3 = rhs397_
            d_2731_v47_ = rhs398_
            index440_ = default__.safeIndex(537, (d_2724_v45_).length(0))
            (d_2724_v45_)[index440_] = p1.f3
            d_2737_v50_: _dafny.MultiSet
            d_2737_v50_ = _dafny.MultiSet([d_2671_v6_])
            d_2738_v51_: _dafny.Map
            d_2738_v51_ = _dafny.Map({(d_2697_v30_).f11: 665})
            d_2739_v52_: D1
            d_2739_v52_ = D1_DC5(p1.f3, (p1).f4, ((d_2737_v50_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fiktmwl")), default__.abs(len(d_2738_v51_)))).cardinality, p0)
            d_2740_v53_: _dafny.Array
            nw442_ = _dafny.Array(None, 8)
            nw442_[int(0)] = d_2739_v52_
            nw442_[int(1)] = D1_DC5(not(True), (d_2697_v30_).f11, -544, p0)
            nw442_[int(2)] = D1_DC5(p1.f3, p1.f3, p0, p0)
            nw442_[int(3)] = d_2739_v52_
            nw442_[int(4)] = default__.fm48((d_2697_v30_).f10, globalState)
            nw442_[int(5)] = d_2739_v52_
            nw442_[int(6)] = d_2739_v52_
            nw442_[int(7)] = d_2739_v52_
            d_2740_v53_ = nw442_
            index441_ = default__.safeIndex(537, (d_2724_v45_).length(0))
            rhs399_ = (p1).f4
            rhs400_ = (d_2740_v53_ if ((d_2697_v30_).f11) or (True) else d_2740_v53_)
            rhs401_ = (p1).f4
            lhs295_ = d_2724_v45_
            lhs296_ = default__.safeIndex(537, (d_2724_v45_).length(0))
            lhs297_ = p1
            lhs295_[lhs296_] = rhs399_
            d_2740_v53_ = rhs400_
            lhs297_.f3 = rhs401_
            r0 = ((d_2738_v51_)[(d_2730_v46_)[default__.safeIndex(p0, len(d_2730_v46_))]] if ((d_2730_v46_)[default__.safeIndex(p0, len(d_2730_v46_))]) in (d_2738_v51_) else (p0) + (p0))
        elif True:
            d_2741___mcc_h3_ = source33_.cf91
            d_2742_cf91_ = d_2741___mcc_h3_
            (p1).f3 = (p1).f4
            r0 = (0) - ((837) + (p0))
            d_2675_v10_ = d_2675_v10_
            d_2743_v54_: _dafny.Seq
            d_2743_v54_ = _dafny.SeqWithoutIsStrInference([(d_2697_v30_).f10, (d_2697_v30_).f11])
            d_2744_v55_: _dafny.MultiSet
            d_2744_v55_ = _dafny.MultiSet([d_2743_v54_])
            d_2745_v56_: D19
            d_2745_v56_ = D19_DC45(len(d_2671_v6_), (d_2673_v8_)[default__.safeIndex((d_2744_v55_).cardinality, len(d_2673_v8_))], True, False, d_2743_v54_)
            (globalState).f0 = not(((d_2745_v56_ if self.f3 else d_2745_v56_)).cf72)
        r0 = p0
        return r0

    def m2(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_2746_v0_: int
        d_2746_v0_ = 304
        d_2747_v1_: _dafny.Map
        d_2747_v1_ = _dafny.Map({d_2746_v0_: (p1).f4})
        d_2748_v2_: _dafny.MultiSet
        d_2748_v2_ = _dafny.MultiSet([p1.f3])
        d_2749_v3_: _dafny.Seq
        d_2749_v3_ = _dafny.SeqWithoutIsStrInference([d_2746_v0_])
        d_2750_v4_: _dafny.Seq
        d_2750_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oq"))
        d_2751_v5_: _dafny.Seq
        d_2751_v5_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(D19_DC46(d_2746_v0_, (p1).f4, (p1).f4, D1_DC5((p1).f4, (self).f4, 119, d_2746_v0_), 795)).cf74])])
        d_2752_v6_: _dafny.Array
        d_2753_v7_: int
        out49_: _dafny.Array
        out50_: int
        out49_, out50_ = (p1).m0(self.f5, d_2747_v1_, ((d_2748_v2_).set(p1.f3, default__.abs((0) - ((d_2749_v3_)[default__.safeIndex((self).fm1(True, len(d_2750_v4_), self.f5, d_2751_v5_, globalState), len(d_2749_v3_))])))).cardinality, globalState)
        d_2752_v6_ = out49_
        d_2753_v7_ = out50_
        d_2754_v8_: _dafny.Map
        d_2754_v8_ = _dafny.Map({((d_2748_v2_).cardinality) <= (d_2753_v7_): self.f5})
        (self).f3 = ((d_2754_v8_)[((d_2754_v8_)[(self).f4] if ((self).f4) in (d_2754_v8_) else self.f5)] if (((d_2754_v8_)[(self).f4] if ((self).f4) in (d_2754_v8_) else self.f5)) in (d_2754_v8_) else self.f5)
        d_2755_v9_: _dafny.MultiSet
        d_2755_v9_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p0 for d_2756_i0_ in range(default__.abs(973))]), d_2750_v4_, default__.fm30(D1_DC5(not(self.f3), self.f5, d_2746_v0_, (self).fm1((p1).f4, len(_dafny.SeqWithoutIsStrInference([self.f5])), p1.f3, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2746_v0_]) for d_2757_i1_ in range(default__.abs(383))]), globalState)), d_2753_v7_, d_2753_v7_, globalState), d_2750_v4_])
        d_2758_v10_: _dafny.Map
        d_2758_v10_ = _dafny.Map({d_2746_v0_: d_2749_v3_})
        d_2759_v11_: D17
        d_2759_v11_ = D17_DC39(d_2752_v6_, (self).f4, self.f3, d_2753_v7_, d_2753_v7_)
        d_2760_v12_: _dafny.Map
        d_2760_v12_ = _dafny.Map({p1.f3: d_2759_v11_})
        nw443_ = _dafny.Array(None, 13)
        nw443_[int(0)] = d_2753_v7_
        nw443_[int(1)] = (d_2753_v7_ if not((self).f4) else -46)
        nw443_[int(2)] = ((d_2755_v9_)[_dafny.SeqWithoutIsStrInference([p0 for d_2761_i2_ in range(default__.abs(-891))])] if (_dafny.SeqWithoutIsStrInference([p0 for d_2762_i2_ in range(default__.abs(-891))])) in (d_2755_v9_) else d_2753_v7_)
        nw443_[int(3)] = len(((d_2758_v10_)[d_2746_v0_] if (d_2746_v0_) in (d_2758_v10_) else _dafny.SeqWithoutIsStrInference([d_2746_v0_ for d_2763_i3_ in range(default__.abs(630))])))
        nw443_[int(4)] = len((d_2760_v12_).set((self).f4, d_2759_v11_))
        nw443_[int(5)] = d_2746_v0_
        nw443_[int(6)] = d_2746_v0_
        nw443_[int(7)] = d_2746_v0_
        nw443_[int(8)] = default__.safeDivisionInt(d_2746_v0_, (0) - (d_2746_v0_))
        nw443_[int(9)] = len(d_2750_v4_)
        nw443_[int(10)] = d_2746_v0_
        nw443_[int(11)] = d_2753_v7_
        nw443_[int(12)] = d_2753_v7_
        d_2752_v6_ = nw443_
        hi16_ = (d_2753_v7_) - (d_2746_v0_)
        for d_2764_i4_ in range(d_2746_v0_, hi16_):
            d_2746_v0_ = d_2764_i4_
            d_2753_v7_ = d_2746_v0_
            d_2765_v13_: _dafny.Seq
            d_2765_v13_ = _dafny.SeqWithoutIsStrInference([(self).f4])
            d_2748_v2_ = default__.fm20((_dafny.SeqWithoutIsStrInference([(self).f4])) == (d_2765_v13_), globalState)
            d_2765_v13_ = ((d_2765_v13_) + (_dafny.SeqWithoutIsStrInference([self.f5, p1.f3]))) + ((d_2765_v13_ if p1.f3 else d_2765_v13_))
        d_2766_i5_: int
        d_2766_i5_ = 0
        with _dafny.label("10"):
            while self.f5:
                with _dafny.c_label("10"):
                    if (d_2766_i5_) >= (100):
                        raise _dafny.Break("10")
                    d_2766_i5_ = (d_2766_i5_) + (1)
                    (p1).f3 = (p1).f4
                    index442_ = default__.safeIndex(799, (d_2752_v6_).length(0))
                    (d_2752_v6_)[index442_] = (23) + (d_2753_v7_)
                    index443_ = default__.safeIndex(799, (d_2752_v6_).length(0))
                    def iife283_():
                        coll87_ = _dafny.Set()
                        compr_87_: int
                        for compr_87_ in (d_2747_v1_).keys.Elements:
                            d_2767_v14_: int = compr_87_
                            if (d_2767_v14_) in (d_2747_v1_):
                                coll87_ = coll87_.union(_dafny.Set([(d_2767_v14_) * ((len(_dafny.Map({not(True): (0) - ((0) - (len(_dafny.Map({True: _dafny.CodePoint('w')}))))}))) * ((D1_DC5(not(True), not(True), 803, (_dafny.MultiSet([True])).cardinality)).cf11))]))
                        return _dafny.Set(coll87_)
                    (d_2752_v6_)[index443_] = len(iife283_()
                    )
                    d_2746_v0_ = d_2753_v7_
                    (self).f5 = False
                    pass
            pass
        d_2768_v15_: _dafny.Array
        d_2769_v16_: int
        out51_: _dafny.Array
        out52_: int
        out51_, out52_ = (p1).m0(p1.f3, d_2747_v1_, d_2746_v0_, globalState)
        d_2768_v15_ = out51_
        d_2769_v16_ = out52_
        r0 = (d_2750_v4_) + (d_2750_v4_)
        return r0

