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
    def fm0(p0, p1, p2, p3, globalState):
        return ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbkunt")))) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: _dafny.Set({(0) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ich"))), 910])).cardinality), 208})}))])))) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgadqenod"))))

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        return _dafny.Map({(_dafny.Set({not(not(False))})).ispropersubset(_dafny.Set({True, True, not(not(False))})): 646})

    @staticmethod
    def fm5(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([False, False])) + (_dafny.SeqWithoutIsStrInference([not(False)]))

    @staticmethod
    def fm6(globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(-866, 9):
                d_0_v0_: int = compr_0_
                if ((-866) <= (d_0_v0_)) and ((d_0_v0_) < (9)):
                    coll0_ = coll0_.union(_dafny.Set([(d_0_v0_) - ((_dafny.MultiSet([D2_DC7(_dafny.SeqWithoutIsStrInference([False]))])).cardinality)]))
            return _dafny.Set(coll0_)
        return ((iife0_()
        ) | (_dafny.Set({len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")): 696}))}))) - ((_dafny.Set({267})).intersection(_dafny.Set({305})))

    @staticmethod
    def fm7(globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({not(True): True}))) | ((_dafny.Map({True: False})) | (_dafny.Map({False: True})))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([(D9_DC30(_dafny.CodePoint('q'), False, -248, _dafny.Set({588, 628, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "slllsh"))), len(_dafny.Map({False: True})), -740}), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbe"))))).cf54])})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([-855])}))

    @staticmethod
    def fm13(p0, p1, globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([620]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsrbg"))), 617, (_dafny.MultiSet([True])).cardinality])]))).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([False, not(False)]))) for d_1_i0_ in range(default__.abs(481))]), _dafny.SeqWithoutIsStrInference([-578])]))).intersection(_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Set({False})): True}))]), _dafny.SeqWithoutIsStrInference([-82, 371])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-616])]))))

    @staticmethod
    def fm14(globalState):
        return (_dafny.Map({-188: 202})) | ((_dafny.Map({330: 681})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([524 for d_2_i0_ in range(default__.abs(619))])): len(_dafny.SeqWithoutIsStrInference([False]))})))

    @staticmethod
    def fm19(p0, p1, globalState):
        return (_dafny.Set({_dafny.CodePoint('a'), _dafny.CodePoint('d'), _dafny.CodePoint('a')})) - (_dafny.Set({_dafny.CodePoint('g'), _dafny.CodePoint('s'), _dafny.CodePoint('m'), _dafny.CodePoint('d'), _dafny.CodePoint('h')}))

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        source0_ = (D7_DC23(False, 842, False) if False else D7_DC23(True, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])), -164])), not(True)))
        if source0_.is_DC22:
            d_3___mcc_h0_ = source0_.cf37
            d_4___mcc_h1_ = source0_.cf38
            d_5___mcc_h2_ = source0_.cf39
            d_6___mcc_h3_ = source0_.cf40
            d_7___mcc_h4_ = source0_.cf41
            d_8_cf41_ = d_7___mcc_h4_
            d_9_cf40_ = d_6___mcc_h3_
            d_10_cf39_ = d_5___mcc_h2_
            d_11_cf38_ = d_4___mcc_h1_
            d_12_cf37_ = d_3___mcc_h0_
            return D0_DC0(True)
        elif source0_.is_DC23:
            d_13___mcc_h5_ = source0_.cf42
            d_14___mcc_h6_ = source0_.cf43
            d_15___mcc_h7_ = source0_.cf44
            d_16_cf44_ = d_15___mcc_h7_
            d_17_cf43_ = d_14___mcc_h6_
            d_18_cf42_ = d_13___mcc_h5_
            return D0_DC0(d_18_cf42_)
        elif True:
            d_19___mcc_h8_ = source0_.cf36
            d_20_cf36_ = d_19___mcc_h8_
            return D0_DC0(True)

    @staticmethod
    def fm21(p0, p1, p2, globalState):
        source1_ = D4_DC14(True, (D7_DC23(True, 781, True)).cf44)
        if source1_.is_DC13:
            d_21___mcc_h0_ = source1_.cf22
            d_22___mcc_h1_ = source1_.cf23
            d_23___mcc_h2_ = source1_.cf24
            d_24___mcc_h3_ = source1_.cf25
            d_25_cf25_ = d_24___mcc_h3_
            d_26_cf24_ = d_23___mcc_h2_
            d_27_cf23_ = d_22___mcc_h1_
            d_28_cf22_ = d_21___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([d_27_cf23_ for d_29_i0_ in range(default__.abs(339))])
        elif source1_.is_DC14:
            d_30___mcc_h4_ = source1_.cf26
            d_31___mcc_h5_ = source1_.cf27
            d_32_cf27_ = d_31___mcc_h5_
            d_33_cf26_ = d_30___mcc_h4_
            def iife1_():
                coll1_ = _dafny.Set()
                compr_1_: int
                for compr_1_ in _dafny.IntegerRange(752, 428):
                    d_34_v0_: int = compr_1_
                    if ((752) <= (d_34_v0_)) and ((d_34_v0_) < (428)):
                        coll1_ = coll1_.union(_dafny.Set([(d_34_v0_) - (708)]))
                return _dafny.Set(coll1_)
            return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_33_cf26_])), len(iife1_()
)])).cardinality})), 0]))])) for d_35_i1_ in range(default__.abs(-807))])
        elif source1_.is_DC12:
            d_36___mcc_h6_ = source1_.cf21
            d_37_cf21_ = d_36___mcc_h6_
            return _dafny.SeqWithoutIsStrInference([106, len(_dafny.Set({-163, len(_dafny.Map({937: 667})), 795, len(_dafny.Set({True, not(True), not(True)}))}))])
        elif True:
            d_38___mcc_h7_ = source1_.cf28
            d_39_cf28_ = d_38___mcc_h7_
            return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdqjqr"))), (_dafny.MultiSet([141])).cardinality, len(_dafny.SeqWithoutIsStrInference([not(True), False])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvuwwxen"))), (0) - ((0) - ((_dafny.MultiSet([False, not(True), True])).cardinality))])

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        if not(not (True) or (True)):
            def iife2_():
                coll2_ = _dafny.Set()
                compr_2_: int
                for compr_2_ in _dafny.IntegerRange(768, -704):
                    d_40_v0_: int = compr_2_
                    if ((768) <= (d_40_v0_)) and ((d_40_v0_) < (-704)):
                        coll2_ = coll2_.union(_dafny.Set([default__.safeDivisionInt(d_40_v0_, 236)]))
                return _dafny.Set(coll2_)
            return (_dafny.Set({(0) - (len(_dafny.Set({len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nuqceuqr"))), -359]))})), 870})))})) | (iife2_()
            )
        elif True:
            def iife3_():
                coll3_ = _dafny.Set()
                compr_3_: int
                for compr_3_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))]))).Elements:
                    d_41_v1_: int = compr_3_
                    if (d_41_v1_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))]))):
                        coll3_ = coll3_.union(_dafny.Set([default__.safeModuloInt(d_41_v1_, len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([688, 355, len(_dafny.SeqWithoutIsStrInference([not(True), True]))])), len(_dafny.SeqWithoutIsStrInference([384 for d_42_i0_ in range(default__.abs(412))]))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dp"))), -318, (0) - (len(_dafny.Map({True: True}))), 813])))])))]))
                return _dafny.Set(coll3_)
            return _dafny.Set({569, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, not(False)]))).cardinality, len(iife3_()
            ), -904})

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        return True

    @staticmethod
    def fm24(p0, p1, p2, p3, globalState):
        return (_dafny.Map({_dafny.CodePoint('s'): (0) - (-778)})) | (_dafny.Map({_dafny.CodePoint('y'): (0) - (len(_dafny.Map({not(not(True)): (D18_DC56(True, True, len(_dafny.Set({754, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kk")))})))).cf92})))}))

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('g')

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        source2_ = D12_DC38((D9_DC30(_dafny.CodePoint('y'), True, (0) - (len(_dafny.Map({True: -170}))), _dafny.Set({(0) - ((_dafny.MultiSet([(0) - (-48), 262])).cardinality), len(_dafny.Set({True}))}), -96)).cf52)
        if source2_.is_DC38:
            d_43___mcc_h0_ = source2_.cf64
            d_44_cf64_ = d_43___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('x')])])
        elif source2_.is_DC39:
            d_45___mcc_h1_ = source2_.cf65
            d_46_cf65_ = d_45___mcc_h1_
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_47_i1_ in range(default__.abs(277))]) for d_48_i0_ in range(default__.abs(881))])
        elif source2_.is_DC37:
            d_49___mcc_h2_ = source2_.cf63
            d_50_cf63_ = d_49___mcc_h2_
            def iife4_():
                coll4_ = _dafny.Set()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(750, 208):
                    d_52_v0_: int = compr_4_
                    if ((750) <= (d_52_v0_)) and ((d_52_v0_) < (208)):
                        def iife5_():
                            coll5_ = _dafny.Set()
                            compr_5_: int
                            for compr_5_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: False}))])): True})).keys.Elements:
                                d_53_v1_: int = compr_5_
                                if (d_53_v1_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: False}))])): True})):
                                    coll5_ = coll5_.union(_dafny.Set([default__.safeModuloInt(d_53_v1_, 881)]))
                            return _dafny.Set(coll5_)
                        coll4_ = coll4_.union(_dafny.Set([default__.safeModuloInt(d_52_v0_, len(iife5_()
))]))
                return _dafny.Set(coll4_)
            return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_51_i2_ in range(default__.abs(114))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(D9_DC30(_dafny.CodePoint('j'), True, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality, iife4_()
, -338)).cf50]) for d_54_i3_ in range(default__.abs(87))]))
        elif True:
            d_55___mcc_h3_ = source2_.cf66
            d_56_cf66_ = d_55___mcc_h3_
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_57_i4_ in range(default__.abs(874))])])

    @staticmethod
    def fm27(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkgblbh"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_58_i0_ in range(default__.abs(122))]))

    @staticmethod
    def fm28(p0, globalState):
        return D1_DC4((len(_dafny.Map({_dafny.CodePoint('a'): False}))) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyccnriii")))), (-573) - (618))

    @staticmethod
    def fm29(globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([True, not(True)]))])

    @staticmethod
    def fm30(p0, globalState):
        return (_dafny.MultiSet([-690])) - (_dafny.MultiSet([-639, (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False])])).cardinality, -736, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmgn")))])) for d_59_i0_ in range(default__.abs(-596))])), 465]))

    @staticmethod
    def fm31(p0, p1, globalState):
        return _dafny.Map({(0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymdokrv"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_60_i0_ in range(default__.abs(957))])))): _dafny.CodePoint('c')})

    @staticmethod
    def fm32(p0, p1, globalState):
        def iife6_():
            coll6_ = _dafny.Set()
            compr_6_: str
            for compr_6_ in ((_dafny.Map({_dafny.CodePoint('b'): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True]))})) | (_dafny.Map({_dafny.CodePoint('p'): _dafny.MultiSet([False])}))).keys.Elements:
                d_61_v0_: str = compr_6_
                if (d_61_v0_) in ((_dafny.Map({_dafny.CodePoint('b'): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True]))})) | (_dafny.Map({_dafny.CodePoint('p'): _dafny.MultiSet([False])}))):
                    coll6_ = coll6_.union(_dafny.Set([d_61_v0_]))
            return _dafny.Set(coll6_)
        return iife6_()
        

    @staticmethod
    def fm33(p0, globalState):
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([not(not(False)), not(not(True))])) + (_dafny.SeqWithoutIsStrInference([True, True]))) + ((_dafny.SeqWithoutIsStrInference([not(True)])) + (_dafny.SeqWithoutIsStrInference([True, False, True, False]))))

    @staticmethod
    def fm34(globalState):
        def iife7_():
            def iife9_():
                coll9_ = _dafny.Map()
                compr_9_: _dafny.Seq
                for compr_9_ in (_dafny.SeqWithoutIsStrInference([(D22_DC70(True, not(False), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgiu")), (_dafny.MultiSet([331, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ia")))])).cardinality, len(_dafny.Map({-47: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_62_i1_ in range(default__.abs(860))]))})))).cf114 for d_63_i0_ in range(default__.abs(915))])).Elements:
                    d_64_v0_: _dafny.Seq = compr_9_
                    if (d_64_v0_) in (_dafny.SeqWithoutIsStrInference([(D22_DC70(True, not(False), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgiu")), (_dafny.MultiSet([331, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ia")))])).cardinality, len(_dafny.Map({-47: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_62_i1_ in range(default__.abs(860))]))})))).cf114 for d_63_i0_ in range(default__.abs(915))])):
                        coll9_[d_64_v0_] = len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, not(True)])]))
                return _dafny.Map(coll9_)
            coll7_ = _dafny.Set()
            def iife8_():
                coll8_ = _dafny.Map()
                compr_8_: _dafny.Seq
                for compr_8_ in (_dafny.SeqWithoutIsStrInference([(D22_DC70(True, not(False), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgiu")), (_dafny.MultiSet([331, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ia")))])).cardinality, len(_dafny.Map({-47: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_62_i1_ in range(default__.abs(860))]))})))).cf114 for d_63_i0_ in range(default__.abs(915))])).Elements:
                    d_64_v0_: _dafny.Seq = compr_8_
                    if (d_64_v0_) in (_dafny.SeqWithoutIsStrInference([(D22_DC70(True, not(False), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgiu")), (_dafny.MultiSet([331, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ia")))])).cardinality, len(_dafny.Map({-47: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_62_i1_ in range(default__.abs(860))]))})))).cf114 for d_63_i0_ in range(default__.abs(915))])):
                        coll8_[d_64_v0_] = len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, not(True)])]))
                return _dafny.Map(coll8_)
            compr_7_: _dafny.Seq
            for compr_7_ in ((iife8_()
            ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elx")): 919}))).keys.Elements:
                d_65_v1_: _dafny.Seq = compr_7_
                if (d_65_v1_) in ((iife9_()
                ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elx")): 919}))):
                    coll7_ = coll7_.union(_dafny.Set([d_65_v1_]))
            return _dafny.Set(coll7_)
        return iife7_()
        

    @staticmethod
    def fm35(p0, p1, globalState):
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(527, 315):
                d_66_v0_: int = compr_10_
                if ((527) <= (d_66_v0_)) and ((d_66_v0_) < (315)):
                    coll10_[(d_66_v0_) + (909)] = D5_DC18(not(False))
            return _dafny.Map(coll10_)
        return (_dafny.Map({len(_dafny.Map({(D0_DC1(False)).cf1: False})): D5_DC18(False)})) | ((iife10_()
        ) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "niflmco"))): D5_DC18(True)})))

    @staticmethod
    def fm36(p0, p1, p2, p3, globalState):
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(954, 68):
                d_67_v0_: int = compr_11_
                if ((954) <= (d_67_v0_)) and ((d_67_v0_) < (68)):
                    coll11_[(d_67_v0_) + (len(_dafny.Set({False})))] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: _dafny.CodePoint('f')}))]))
            return _dafny.Map(coll11_)
        return ((_dafny.SeqWithoutIsStrInference([_dafny.Map({True: 510})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): len(_dafny.Map({len(_dafny.Map({True: not(False)})): 940}))})]))) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xcail")))}), _dafny.Map({False: len(iife11_()
        )})]))

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        return D0_DC1(not((len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): True}))) < (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvidhsf"))))))

    @staticmethod
    def fm38(globalState):
        source3_ = D5_DC17(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False, False])]), False, _dafny.MultiSet([372, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_68_i0_ in range(default__.abs(142))])): len(_dafny.Set({_dafny.CodePoint('w'), _dafny.CodePoint('n'), _dafny.CodePoint('j'), _dafny.CodePoint('r')}))})), len(_dafny.Map({False: _dafny.MultiSet([839])})), len(_dafny.Map({327: False}))]))
        if source3_.is_DC17:
            d_69___mcc_h0_ = source3_.cf30
            d_70___mcc_h1_ = source3_.cf31
            d_71___mcc_h2_ = source3_.cf32
            d_72_cf32_ = d_71___mcc_h2_
            d_73_cf31_ = d_70___mcc_h1_
            d_74_cf30_ = d_69___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([d_73_cf31_])) + (_dafny.SeqWithoutIsStrInference([d_73_cf31_]))
        elif source3_.is_DC18:
            d_75___mcc_h3_ = source3_.cf33
            d_76_cf33_ = d_75___mcc_h3_
            return _dafny.SeqWithoutIsStrInference([d_76_cf33_, True])
        elif source3_.is_DC16:
            d_77___mcc_h4_ = source3_.cf29
            d_78_cf29_ = d_77___mcc_h4_
            return _dafny.SeqWithoutIsStrInference([True, True])
        elif True:
            d_79___mcc_h5_ = source3_.cf34
            d_80_cf34_ = d_79___mcc_h5_
            return (_dafny.SeqWithoutIsStrInference([False, True, not(False), True, True])) + (_dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm39(p0, p1, p2, p3, globalState):
        def iife12_():
            coll12_ = _dafny.Set()
            compr_12_: int
            for compr_12_ in _dafny.IntegerRange(626, 349):
                d_81_v0_: int = compr_12_
                if ((626) <= (d_81_v0_)) and ((d_81_v0_) < (349)):
                    coll12_ = coll12_.union(_dafny.Set([default__.safeModuloInt(d_81_v0_, 168)]))
            return _dafny.Set(coll12_)
        return D9_DC30(_dafny.CodePoint('w'), (False if True else True), (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfdadptf"))): True}))) - (844), iife12_()
, len(_dafny.Map({len(_dafny.Set({(D7_DC23(False, 232, not(False))).cf43, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ixq"))), len(_dafny.Map({-233: True}))})): True})))

    @staticmethod
    def fm40(globalState):
        return _dafny.SeqWithoutIsStrInference([D2_DC7(_dafny.SeqWithoutIsStrInference([False, True]))])

    @staticmethod
    def fm41(p0, p1, globalState):
        return ((_dafny.Set({not(False), not(True), not(True)})).intersection(_dafny.Set({not(True)}))) | ((_dafny.Set({not(not(not(False)))})) | (_dafny.Set({False})))

    @staticmethod
    def fm42(p0, p1, p2, p3, globalState):
        source4_ = D11_DC33(_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))}))
        if source4_.is_DC34:
            d_82___mcc_h0_ = source4_.cf58
            d_83___mcc_h1_ = source4_.cf59
            d_84___mcc_h2_ = source4_.cf60
            d_85___mcc_h3_ = source4_.cf61
            d_86_cf61_ = d_85___mcc_h3_
            d_87_cf60_ = d_84___mcc_h2_
            d_88_cf59_ = d_83___mcc_h1_
            d_89_cf58_ = d_82___mcc_h0_
            def iife13_():
                coll13_ = _dafny.Map()
                compr_13_: int
                for compr_13_ in ((_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "di"))))]))).Elements:
                    d_90_v0_: int = compr_13_
                    if (d_90_v0_) in ((_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "di"))))]))):
                        coll13_[(d_90_v0_) - (-650)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pc"))
                return _dafny.Map(coll13_)
            return D8_DC26(True, iife13_()
)
        elif source4_.is_DC35:
            return D8_DC26(False, _dafny.Map({66: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_91_i0_ in range(default__.abs(108))])}))
        elif source4_.is_DC33:
            d_92___mcc_h4_ = source4_.cf57
            d_93_cf57_ = d_92___mcc_h4_
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(528, 791):
                    d_94_v1_: int = compr_14_
                    if ((528) <= (d_94_v1_)) and ((d_94_v1_) < (791)):
                        coll14_[(d_94_v1_) - (586)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsqevyqb"))
                return _dafny.Map(coll14_)
            return D8_DC26(not(True), iife14_()
)
        elif True:
            d_95___mcc_h5_ = source4_.cf62
            d_96_cf62_ = d_95___mcc_h5_
            if not(True):
                return D8_DC26(True, _dafny.Map({-119: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_97_i1_ in range(default__.abs(169))])}))
            elif True:
                return D8_DC26(False, _dafny.Map({-953: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))}))

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(D0_DC0(True)).cf0, not(True), True, True])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False])), _dafny.MultiSet([False]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])), _dafny.MultiSet([not(False)])]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False])) for d_98_i0_ in range(default__.abs(-251))])))

    @staticmethod
    def fm44(p0, p1, p2, globalState):
        source5_ = D13_DC41(_dafny.Set({_dafny.SeqWithoutIsStrInference([689])}))
        if source5_.is_DC42:
            d_99___mcc_h0_ = source5_.cf68
            d_100___mcc_h1_ = source5_.cf69
            d_101_cf69_ = d_100___mcc_h1_
            d_102_cf68_ = d_99___mcc_h0_
            return D2_DC8((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_102_cf68_ for d_103_i0_ in range(default__.abs(609))])]))), -147, -214)
        elif source5_.is_DC41:
            d_104___mcc_h2_ = source5_.cf67
            d_105_cf67_ = d_104___mcc_h2_
            return D2_DC8(len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True})), 117])), 293, -852)
        elif True:
            d_106___mcc_h3_ = source5_.cf70
            d_107_cf70_ = d_106___mcc_h3_
            return D2_DC8(len(_dafny.Set({188})), len(_dafny.Map({True: True})), len(_dafny.SeqWithoutIsStrInference([False, False, True])))

    @staticmethod
    def fm45(p0, p1, p2, globalState):
        source6_ = D9_DC30(_dafny.CodePoint('q'), False, len(_dafny.Set({len(_dafny.Map({(0) - ((0) - (len(_dafny.Map({False: True})))): 0})), (_dafny.MultiSet([False, True, False])).cardinality})), _dafny.Set({335, 218, -979}), len(_dafny.Map({-175: len(_dafny.SeqWithoutIsStrInference([not(False)]))})))
        if source6_.is_DC30:
            d_108___mcc_h0_ = source6_.cf50
            d_109___mcc_h1_ = source6_.cf51
            d_110___mcc_h2_ = source6_.cf52
            d_111___mcc_h3_ = source6_.cf53
            d_112___mcc_h4_ = source6_.cf54
            d_113_cf54_ = d_112___mcc_h4_
            d_114_cf53_ = d_111___mcc_h3_
            d_115_cf52_ = d_110___mcc_h2_
            d_116_cf51_ = d_109___mcc_h1_
            d_117_cf50_ = d_108___mcc_h0_
            return D3_DC10(_dafny.Map({d_116_cf51_: d_116_cf51_}))
        elif source6_.is_DC29:
            d_118___mcc_h5_ = source6_.cf49
            d_119_cf49_ = d_118___mcc_h5_
            return D3_DC10(_dafny.Map({True: False}))
        elif True:
            d_120___mcc_h6_ = source6_.cf55
            d_121_cf55_ = d_120___mcc_h6_
            return D3_DC10(_dafny.Map({True: False}))

    @staticmethod
    def fm46(p0, p1, globalState):
        def iife15_():
            coll15_ = _dafny.Map()
            compr_15_: str
            for compr_15_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('g'), _dafny.CodePoint('b'), _dafny.CodePoint('e'), _dafny.CodePoint('a')])).Elements:
                d_122_v0_: str = compr_15_
                if (d_122_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('g'), _dafny.CodePoint('b'), _dafny.CodePoint('e'), _dafny.CodePoint('a')])):
                    coll15_[d_122_v0_] = True
            return _dafny.Map(coll15_)
        return (_dafny.Map({_dafny.CodePoint('p'): not(not(True))})) | (iife15_()
        )

    @staticmethod
    def fm47(globalState):
        def iife16_():
            coll16_ = _dafny.Set()
            compr_16_: _dafny.MultiSet
            for compr_16_ in (_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False, False, True])), _dafny.MultiSet([True]), _dafny.MultiSet([False, not(False)]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True), True, True, True]))})).Elements:
                d_123_v0_: _dafny.MultiSet = compr_16_
                if (d_123_v0_) in (_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False, False, True])), _dafny.MultiSet([True]), _dafny.MultiSet([False, not(False)]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True), True, True, True]))})):
                    coll16_ = coll16_.union(_dafny.Set([d_123_v0_]))
            return _dafny.Set(coll16_)
        return ((_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False])), _dafny.MultiSet([True]), _dafny.MultiSet([False])})) - (_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True, False, False])), _dafny.MultiSet([not(False)])}))) | (iife16_()
        )

    @staticmethod
    def fm48(p0, p1, p2, globalState):
        return _dafny.Map({_dafny.Map({(0) - (len(_dafny.Set({not(True)}))): len(_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True}))}): (871) + (758)})

    @staticmethod
    def fm49(p0, globalState):
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: _dafny.Seq
            for compr_17_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([386]): _dafny.Map({not(True): True})})).keys.Elements:
                d_124_v0_: _dafny.Seq = compr_17_
                if (d_124_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([386]): _dafny.Map({not(True): True})})):
                    coll17_[d_124_v0_] = False
            return _dafny.Map(coll17_)
        def iife18_():
            coll18_ = _dafny.Set()
            compr_18_: int
            for compr_18_ in (_dafny.MultiSet([954])).Elements:
                d_125_v1_: int = compr_18_
                if (d_125_v1_) in (_dafny.MultiSet([954])):
                    coll18_ = coll18_.union(_dafny.Set([default__.safeDivisionInt(d_125_v1_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({-253: len(_dafny.Map({False: 880}))})): D1_DC4(len(_dafny.Set({-15})), (_dafny.MultiSet([974, len(_dafny.SeqWithoutIsStrInference([not(False), False])), len(_dafny.Set({True})), 106, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpse")))])).cardinality)})), 162, len(_dafny.SeqWithoutIsStrInference([False, False, not(False)])), 929])))]))
            return _dafny.Set(coll18_)
        def iife19_():
            coll19_ = _dafny.Map()
            compr_19_: int
            for compr_19_ in _dafny.IntegerRange(140, 644):
                d_126_v2_: int = compr_19_
                if ((140) <= (d_126_v2_)) and ((d_126_v2_) < (644)):
                    coll19_[(d_126_v2_) + (293)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phlpopw")))
            return _dafny.Map(coll19_)
        def iife20_():
            def iife22_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in _dafny.IntegerRange(394, -187):
                    d_129_v3_: int = compr_22_
                    if ((394) <= (d_129_v3_)) and ((d_129_v3_) < (-187)):
                        coll22_[default__.safeDivisionInt(d_129_v3_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_130_i2_ in range(default__.abs(919))])))] = (_dafny.MultiSet([-559, 990])).cardinality
                return _dafny.Map(coll22_)
            coll20_ = _dafny.Set()
            def iife21_():
                coll21_ = _dafny.Map()
                compr_21_: int
                for compr_21_ in _dafny.IntegerRange(394, -187):
                    d_129_v3_: int = compr_21_
                    if ((394) <= (d_129_v3_)) and ((d_129_v3_) < (-187)):
                        coll21_[default__.safeDivisionInt(d_129_v3_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_130_i2_ in range(default__.abs(919))])))] = (_dafny.MultiSet([-559, 990])).cardinality
                return _dafny.Map(coll21_)
            compr_20_: _dafny.Seq
            for compr_20_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D9_DC31(D9_DC31(D9_DC29(iife21_()
))), D9_DC31(D9_DC29(_dafny.Map({(_dafny.MultiSet([False])).cardinality: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_131_i3_ in range(default__.abs(395))]))})))])])).Elements:
                d_132_v4_: _dafny.Seq = compr_20_
                if (d_132_v4_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D9_DC31(D9_DC31(D9_DC29(iife22_()
))), D9_DC31(D9_DC29(_dafny.Map({(_dafny.MultiSet([False])).cardinality: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_131_i3_ in range(default__.abs(395))]))})))])])):
                    coll20_ = coll20_.union(_dafny.Set([d_132_v4_]))
            return _dafny.Set(coll20_)
        def iife23_():
            coll23_ = _dafny.Map()
            compr_23_: int
            for compr_23_ in (_dafny.Map({(_dafny.MultiSet([False, False])).cardinality: False})).keys.Elements:
                d_133_v5_: int = compr_23_
                if (d_133_v5_) in (_dafny.Map({(_dafny.MultiSet([False, False])).cardinality: False})):
                    coll23_[default__.safeDivisionInt(d_133_v5_, len(_dafny.Map({_dafny.CodePoint('j'): len(_dafny.Map({126: len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjfiraa")))}))}))})))] = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "to"))))
            return _dafny.Map(coll23_)
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([D9_DC31(D9_DC30(_dafny.CodePoint('h'), False, len(iife17_()
), iife18_()
, 277))]), _dafny.SeqWithoutIsStrInference([D9_DC31(D9_DC31(D9_DC29(iife19_()
)))]), _dafny.SeqWithoutIsStrInference([D9_DC31(D9_DC31(D9_DC30(_dafny.CodePoint('i'), True, len(_dafny.Map({True: len(_dafny.Map({False: True}))})), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([True, True, True, True]))}), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_127_i1_ in range(default__.abs(-825))]))))) for d_128_i0_ in range(default__.abs(-447))])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([D9_DC31((D9_DC31(D9_DC31(D9_DC30(_dafny.CodePoint('w'), False, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False}))])), _dafny.Set({271, 444}), 168)))).cf55), D9_DC31(D9_DC29(_dafny.Map({733: -491})))])}))).intersection((iife20_()
        ) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([D9_DC31(D9_DC31(D9_DC31(D9_DC31(D9_DC31(D9_DC29(_dafny.Map({-463: 601}))))))), D9_DC31(D9_DC31(D9_DC31(D9_DC29(iife23_()
))))])})))

    @staticmethod
    def fm50(p0, p1, p2, p3, globalState):
        def iife24_():
            coll24_ = _dafny.Map()
            compr_24_: int
            for compr_24_ in _dafny.IntegerRange(614, 484):
                d_134_v0_: int = compr_24_
                if ((614) <= (d_134_v0_)) and ((d_134_v0_) < (484)):
                    coll24_[(d_134_v0_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_135_i0_ in range(default__.abs(-399))])))] = True
            return _dafny.Map(coll24_)
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([len(iife24_()
        )])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, False]))])}))

    @staticmethod
    def fm51(p0, p1, p2, p3, globalState):
        return D13_DC43(D13_DC43(D13_DC42(_dafny.CodePoint('k'), True)))

    @staticmethod
    def fm52(p0, globalState):
        return ((_dafny.Map({len(_dafny.Set({False})): _dafny.SeqWithoutIsStrInference([len(_dafny.Set({(_dafny.MultiSet([True])).cardinality})), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlkqqbehd"))) for d_136_i0_ in range(default__.abs(-309))])), 589])})) | (_dafny.Map({578: _dafny.SeqWithoutIsStrInference([(0) - ((0) - (-859))])}))) | (_dafny.Map({-541: _dafny.SeqWithoutIsStrInference([83])}))

    @staticmethod
    def fm53(p0, p1, globalState):
        def iife25_():
            coll25_ = _dafny.Map()
            compr_25_: str
            for compr_25_ in ((D22_DC67(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hac")))).cf108).Elements:
                d_137_v0_: str = compr_25_
                if (d_137_v0_) in ((D22_DC67(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hac")))).cf108):
                    coll25_[d_137_v0_] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_138_i1_ in range(default__.abs(911))]))) for d_139_i0_ in range(default__.abs(426))]))).cardinality
            return _dafny.Map(coll25_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.Set({-116, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({not(False)}))]), _dafny.SeqWithoutIsStrInference([50])])), 945]))}), _dafny.Set({len(iife25_()
        )}), _dafny.Set({106, (0) - (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False])), 427])).cardinality])))})])) + ((_dafny.SeqWithoutIsStrInference([_dafny.Set({(_dafny.MultiSet([_dafny.CodePoint('q')])).cardinality, 326})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({887})])))

    @staticmethod
    def fm54(p0, p1, globalState):
        if not((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pgi")))) >= ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_140_i0_ in range(default__.abs(371))]))))):
            return _dafny.Set({(D8_DC24(_dafny.Set({(D1_DC5(True, False)).cf7}))).cf45})
        elif True:
            def iife26_():
                coll26_ = _dafny.Set()
                compr_26_: _dafny.Set
                for compr_26_ in (_dafny.Set({_dafny.Set({True, False})})).Elements:
                    d_141_v0_: _dafny.Set = compr_26_
                    if (d_141_v0_) in (_dafny.Set({_dafny.Set({True, False})})):
                        coll26_ = coll26_.union(_dafny.Set([d_141_v0_]))
                return _dafny.Set(coll26_)
            return (iife26_()
            ).intersection(_dafny.Set({_dafny.Set({True, False})}))

    @staticmethod
    def fm55(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvjfiqb")))]) for d_142_i0_ in range(default__.abs(352))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-532, -470])]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality for d_143_i1_ in range(default__.abs(328))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_144_i2_ in range(default__.abs(615))])), len(_dafny.SeqWithoutIsStrInference([False]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xoxncefj"))) for d_145_i3_ in range(default__.abs(-994))]), _dafny.SeqWithoutIsStrInference([809])])))

    @staticmethod
    def m11(p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        d_146_v0_: _dafny.Seq
        d_146_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plwnbed"))
        d_147_v1_: _dafny.Array
        nw0_ = _dafny.Array(False, 1)
        d_147_v1_ = nw0_
        d_148_v2_: D4
        d_148_v2_ = D4_DC13(d_146_v0_, p1, False, d_147_v1_)
        def iife27_(_pat_let0_0):
            def iife28_(d_149_dt__update__tmp_h0_):
                def iife29_(_pat_let1_0):
                    def iife30_(d_150_dt__update_hcf22_h0_):
                        return D4_DC13(d_150_dt__update_hcf22_h0_, (d_149_dt__update__tmp_h0_).cf23, (d_149_dt__update__tmp_h0_).cf24, (d_149_dt__update__tmp_h0_).cf25)
                    return iife30_(_pat_let1_0)
                return iife29_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkgqgh")))
            return iife28_(_pat_let0_0)
        source7_ = iife27_((d_148_v2_ if False else d_148_v2_))
        if source7_.is_DC13:
            d_151___mcc_h0_ = source7_.cf22
            d_152___mcc_h1_ = source7_.cf23
            d_153___mcc_h2_ = source7_.cf24
            d_154___mcc_h3_ = source7_.cf25
            d_155_cf25_ = d_154___mcc_h3_
            d_156_cf24_ = d_153___mcc_h2_
            d_157_cf23_ = d_152___mcc_h1_
            d_158_cf22_ = d_151___mcc_h0_
            d_156_cf24_ = (p3 if True else p2)
            (globalState).f9 = p2
            d_159_v3_: _dafny.MultiSet
            d_159_v3_ = _dafny.MultiSet([p3])
            d_160_v4_: T0
            nw1_ = C6()
            nw1_.ctor__(d_156_cf24_, (d_159_v3_) | ((d_159_v3_).set(not(p3), default__.abs(d_157_cf23_))))
            d_160_v4_ = nw1_
            d_160_v4_ = d_160_v4_
            d_161_v5_: _dafny.Array
            nw2_ = _dafny.Array(_dafny.CodePoint('D'), 22)
            d_161_v5_ = nw2_
            d_162_v6_: _dafny.Map
            d_162_v6_ = _dafny.Map({-374: d_161_v5_})
            d_163_v7_: D18
            d_163_v7_ = D18_DC56(p3, d_160_v4_.f14, d_157_cf23_)
            d_164_v8_: _dafny.Seq
            d_164_v8_ = _dafny.SeqWithoutIsStrInference([p2])
            d_165_v9_: C0
            nw3_ = C0()
            nw3_.ctor__((d_164_v8_)[default__.safeIndex(d_157_cf23_, len(d_164_v8_))], p2)
            d_165_v9_ = nw3_
            d_166_v10_: _dafny.Seq
            d_166_v10_ = _dafny.SeqWithoutIsStrInference([d_165_v9_])
            d_167_v11_: _dafny.Map
            d_167_v11_ = _dafny.Map({(d_163_v7_).cf93: len(d_166_v10_)})
            d_168_v12_: _dafny.Seq
            d_168_v12_ = _dafny.SeqWithoutIsStrInference([d_167_v11_, _dafny.Map({p3: d_157_cf23_}), (d_167_v11_).set((d_165_v9_).f24, p1)])
            d_169_v13_: C8
            nw4_ = C8()
            nw4_.ctor__(d_162_v6_, d_168_v12_, p2, d_159_v3_)
            d_169_v13_ = nw4_
            d_170_v14_: _dafny.Array
            nw5_ = _dafny.Array(None, 5)
            nw5_[int(0)] = d_169_v13_
            nw5_[int(1)] = d_169_v13_
            nw5_[int(2)] = d_169_v13_
            nw5_[int(3)] = d_169_v13_
            nw5_[int(4)] = d_169_v13_
            d_170_v14_ = nw5_
            d_170_v14_ = (d_170_v14_ if d_156_cf24_ else d_170_v14_)
        elif source7_.is_DC14:
            d_171___mcc_h4_ = source7_.cf26
            d_172___mcc_h5_ = source7_.cf27
            d_173_cf27_ = d_172___mcc_h5_
            d_174_cf26_ = d_171___mcc_h4_
            d_175_v15_: _dafny.Seq
            d_175_v15_ = _dafny.SeqWithoutIsStrInference([p2])
            d_176_v16_: C6
            nw6_ = C6()
            nw6_.ctor__(True, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_174_cf26_])))
            d_176_v16_ = nw6_
            d_177_v17_: _dafny.Map
            d_177_v17_ = _dafny.Map({p2: d_176_v16_})
            rhs0_ = default__.safeDivisionInt(len((d_175_v15_ if p3 else d_175_v15_)), (len(_dafny.Map({690: d_174_cf26_}))) - (len((d_177_v17_).set((d_175_v15_)[default__.safeIndex(p1, len(d_175_v15_))], d_176_v16_))))
            rhs1_ = (d_176_v16_).fm2((d_175_v15_) <= (d_175_v15_), p2, globalState)
            lhs0_ = globalState
            lhs1_ = globalState
            lhs0_.f3 = rhs0_
            lhs1_.f9 = rhs1_
            d_178_v18_: _dafny.MultiSet
            d_178_v18_ = _dafny.MultiSet([d_174_cf26_, d_173_cf27_])
            d_179_v19_: C4
            nw7_ = C4()
            nw7_.ctor__(d_173_cf27_, d_174_cf26_, d_178_v18_)
            d_179_v19_ = nw7_
            d_180_v20_: _dafny.Map
            d_180_v20_ = _dafny.Map({d_179_v19_: p2})
            d_181_v21_: C5
            nw8_ = C5()
            nw8_.ctor__((p1) + (len(d_180_v20_)), p2, _dafny.MultiSet([d_174_cf26_, True]))
            d_181_v21_ = nw8_
            d_181_v21_ = d_181_v21_
            d_182_v22_: _dafny.Array
            nw9_ = _dafny.Array(int(0), 24)
            d_182_v22_ = nw9_
            index0_ = default__.safeIndex(372, (d_182_v22_).length(0))
            (d_182_v22_)[index0_] = ((d_181_v21_).f18) * ((0) - (p1))
            index1_ = default__.safeIndex(564, (d_147_v1_).length(0))
            (d_147_v1_)[index1_] = (p1) != (p1)
            d_183_v23_: _dafny.Array
            nw10_ = _dafny.Array(None, 9)
            d_183_v23_ = nw10_
            d_184_v24_: _dafny.Set
            d_184_v24_ = _dafny.Set({d_183_v23_, d_183_v23_, d_183_v23_})
            d_185_v25_: _dafny.Map
            d_185_v25_ = _dafny.Map({(d_176_v16_).fm2(p2, (d_179_v19_).f19, globalState): default__.safeDivisionInt((0) - (p1), len(d_184_v24_))})
            index2_ = default__.safeIndex(372, (d_182_v22_).length(0))
            index3_ = default__.safeIndex(564, (d_147_v1_).length(0))
            rhs2_ = (d_146_v0_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_186_i0_ in range(default__.abs(918))]))
            rhs3_ = d_174_cf26_
            rhs4_ = len(d_185_v25_)
            rhs5_ = p2
            lhs2_ = globalState
            lhs3_ = d_182_v22_
            lhs4_ = default__.safeIndex(372, (d_182_v22_).length(0))
            lhs5_ = d_147_v1_
            lhs6_ = default__.safeIndex(564, (d_147_v1_).length(0))
            d_146_v0_ = rhs2_
            lhs2_.f9 = rhs3_
            lhs3_[lhs4_] = rhs4_
            lhs5_[lhs6_] = rhs5_
            index4_ = default__.safeIndex(564, (d_147_v1_).length(0))
            (d_147_v1_)[index4_] = (d_175_v15_)[default__.safeIndex((d_182_v22_)[default__.safeIndex(372, (d_182_v22_).length(0))], len(d_175_v15_))]
        elif source7_.is_DC12:
            d_187___mcc_h6_ = source7_.cf21
            d_188_cf21_ = d_187___mcc_h6_
            index5_ = default__.safeIndex(719, (d_147_v1_).length(0))
            (d_147_v1_)[index5_] = (p1) < (p1)
            d_189_v26_: _dafny.Set
            d_189_v26_ = _dafny.Set({p3, p2, False})
            d_190_v27_: _dafny.Set
            d_190_v27_ = _dafny.Set({d_189_v26_})
            index6_ = default__.safeIndex(719, (d_147_v1_).length(0))
            rhs6_ = (0) - (default__.safeModuloInt(p1, p1))
            rhs7_ = ((d_190_v27_) | (d_190_v27_)).issubset(default__.fm54(850, d_146_v0_, globalState))
            lhs7_ = globalState
            lhs8_ = d_147_v1_
            lhs9_ = default__.safeIndex(719, (d_147_v1_).length(0))
            lhs7_.f3 = rhs6_
            lhs8_[lhs9_] = rhs7_
            if (d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))]:
                (globalState).f3 = p1
                d_191_v28_: _dafny.Seq
                d_191_v28_ = _dafny.SeqWithoutIsStrInference([p1])
                d_192_v29_: _dafny.MultiSet
                d_192_v29_ = _dafny.MultiSet([p3])
                d_193_v30_: C6
                nw11_ = C6()
                nw11_.ctor__((len(d_191_v28_)) > (p1), d_192_v29_)
                d_193_v30_ = nw11_
                (globalState).f3 = (p1) + (p1)
                d_194_v31_: T0
                nw12_ = C3()
                nw12_.ctor__(p1, d_146_v0_, p3, d_192_v29_)
                d_194_v31_ = nw12_
                d_195_v32_: _dafny.Seq
                d_195_v32_ = _dafny.SeqWithoutIsStrInference([d_194_v31_])
                d_196_v33_: _dafny.Array
                nw13_ = _dafny.Array(int(0), 17)
                d_196_v33_ = nw13_
                d_197_v34_: _dafny.Set
                d_197_v34_ = _dafny.Set({d_196_v33_, d_196_v33_, d_196_v33_})
                d_198_v35_: _dafny.Set
                d_198_v35_ = _dafny.Set({len(d_197_v34_)})
                (globalState).f9 = (len((_dafny.SeqWithoutIsStrInference([(d_195_v32_)[default__.safeIndex(p1, len(d_195_v32_))], d_194_v31_, d_194_v31_, d_194_v31_, d_194_v31_])) + (d_195_v32_))) not in (d_198_v35_)
                index7_ = default__.safeIndex(719, (d_147_v1_).length(0))
                (d_147_v1_)[index7_] = not(True)
            elif True:
                index8_ = default__.safeIndex(719, (d_147_v1_).length(0))
                (d_147_v1_)[index8_] = (p1) >= (p1)
                d_199_v36_: _dafny.Map
                d_199_v36_ = _dafny.Map({p1: (d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))]})
                d_200_v37_: D1
                d_200_v37_ = D1_DC5(p3, True)
                d_201_v38_: _dafny.Map
                d_201_v38_ = _dafny.Map({((d_199_v36_)[p1] if (p1) in (d_199_v36_) else True): d_200_v37_})
                d_201_v38_ = d_201_v38_
                d_202_v39_: _dafny.Set
                d_202_v39_ = _dafny.Set({p1, p1})
                d_203_v40_: _dafny.Array
                nw14_ = _dafny.Array(None, 28)
                d_203_v40_ = nw14_
                d_204_v41_: _dafny.Seq
                d_204_v41_ = _dafny.SeqWithoutIsStrInference([d_202_v39_, default__.fm6(globalState)])
                d_205_v42_: _dafny.Seq
                d_205_v42_ = _dafny.SeqWithoutIsStrInference([p2, p3])
                d_206_v43_: _dafny.Map
                d_206_v43_ = _dafny.Map({(d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))]: (d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))]})
                d_207_v44_: str
                d_207_v44_ = _dafny.CodePoint('c')
                d_208_v45_: _dafny.Map
                d_208_v45_ = _dafny.Map({d_207_v44_: p1})
                d_209_v46_: _dafny.Seq
                d_209_v46_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1, p1, p1])
                d_210_v47_: _dafny.Set
                d_210_v47_ = _dafny.Set({d_209_v46_, d_209_v46_, d_209_v46_, d_209_v46_})
                index9_ = default__.safeIndex(719, (d_147_v1_).length(0))
                rhs8_ = ((d_204_v41_)[default__.safeIndex(len(d_205_v42_), len(d_204_v41_))]) - (d_202_v39_)
                rhs9_ = ((d_206_v43_)[not((d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))])] if (not((d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))])) in (d_206_v43_) else default__.fm23(len(d_199_v36_), ((d_199_v36_)[p1] if (p1) in (d_199_v36_) else not(p3)), d_208_v45_, d_210_v47_, globalState))
                rhs10_ = d_203_v40_
                lhs10_ = d_147_v1_
                lhs11_ = default__.safeIndex(719, (d_147_v1_).length(0))
                d_202_v39_ = rhs8_
                lhs10_[lhs11_] = rhs9_
                d_203_v40_ = rhs10_
                d_211_v48_: _dafny.Array
                nw15_ = _dafny.Array(None, 8)
                nw15_[int(0)] = d_188_cf21_
                nw15_[int(1)] = d_188_cf21_
                nw15_[int(2)] = d_188_cf21_
                nw15_[int(3)] = d_188_cf21_
                nw15_[int(4)] = d_188_cf21_
                nw15_[int(5)] = d_188_cf21_
                nw15_[int(6)] = d_188_cf21_
                nw15_[int(7)] = d_188_cf21_
                d_211_v48_ = nw15_
                index10_ = default__.safeIndex(246, (d_211_v48_).length(0))
                (d_211_v48_)[index10_] = d_188_cf21_
                index11_ = default__.safeIndex(246, (d_211_v48_).length(0))
                nw16_ = _dafny.Array(False, 29)
                (d_211_v48_)[index11_] = nw16_
                d_212_v49_: _dafny.MultiSet
                d_212_v49_ = _dafny.MultiSet([p3, p2])
                d_213_v50_: C5
                nw17_ = C5()
                nw17_.ctor__(p1, p3, (d_212_v49_) - ((d_212_v49_).set(p3, default__.abs((d_212_v49_).cardinality))))
                d_213_v50_ = nw17_
            d_214_v51_: _dafny.Map
            d_214_v51_ = _dafny.Map({p2: p3})
            d_215_v52_: _dafny.MultiSet
            d_215_v52_ = _dafny.MultiSet([(d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))], (d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))], p2])
            d_216_v53_: _dafny.MultiSet
            d_216_v53_ = _dafny.MultiSet([-561, p1])
            d_217_v54_: _dafny.Seq
            d_217_v54_ = _dafny.SeqWithoutIsStrInference([p1])
            d_218_v55_: _dafny.Seq
            d_218_v55_ = _dafny.SeqWithoutIsStrInference([d_217_v54_])
            d_219_v56_: _dafny.Map
            d_219_v56_ = _dafny.Map({False: 622})
            d_220_v57_: _dafny.Array
            nw18_ = _dafny.Array(None, 23)
            nw18_[int(0)] = default__.safeModuloInt(p1, p1)
            nw18_[int(1)] = default__.fm0(p2, (0) - (p1), p1, (_dafny.MultiSet([len(d_146_v0_), -622, p1, p1, (0) - (p1)])).set(p1, default__.abs((0) - (len(d_214_v51_)))), globalState)
            nw18_[int(2)] = p1
            nw18_[int(3)] = (p1) + ((0) - ((d_215_v52_).cardinality))
            nw18_[int(4)] = p1
            nw18_[int(5)] = 978
            nw18_[int(6)] = default__.fm0(False, p1, p1, d_216_v53_, globalState)
            nw18_[int(7)] = len(d_218_v55_)
            nw18_[int(8)] = p1
            nw18_[int(9)] = (p1 if p2 else p1)
            nw18_[int(10)] = ((d_217_v54_)[default__.safeIndex(p1, len(d_217_v54_))] if (d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))] else p1)
            nw18_[int(11)] = default__.safeDivisionInt(p1, p1)
            nw18_[int(12)] = p1
            nw18_[int(13)] = 929
            nw18_[int(14)] = (0) - (len(d_217_v54_))
            nw18_[int(15)] = p1
            nw18_[int(16)] = (0) - (len(d_217_v54_))
            nw18_[int(17)] = p1
            nw18_[int(18)] = ((d_219_v56_)[(d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))]] if ((d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))]) in (d_219_v56_) else 171)
            nw18_[int(19)] = p1
            nw18_[int(20)] = p1
            nw18_[int(21)] = p1
            nw18_[int(22)] = p1
            d_220_v57_ = nw18_
            d_220_v57_ = d_220_v57_
            if ((0) - (default__.safeDivisionInt(p1, p1))) <= (p1):
                (globalState).f9 = not(((d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))] if True else p2))
                d_221_v58_: _dafny.Map
                d_221_v58_ = _dafny.Map({d_188_cf21_: (d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))]})
                d_222_v59_: str
                d_222_v59_ = _dafny.CodePoint('p')
                d_223_v60_: _dafny.Map
                d_223_v60_ = _dafny.Map({d_222_v59_: len(_dafny.SeqWithoutIsStrInference([p1, len(_dafny.Set({p1}))]))})
                d_224_v61_: _dafny.Set
                d_224_v61_ = _dafny.Set({d_217_v54_})
                def iife31_():
                    coll27_ = _dafny.Set()
                    compr_27_: _dafny.Seq
                    for compr_27_ in (default__.fm55(default__.fm23(p1, p2, d_223_v60_, d_224_v61_, globalState), globalState)).Elements:
                        d_225_v62_: _dafny.Seq = compr_27_
                        if (d_225_v62_) in (default__.fm55(default__.fm23(p1, p2, d_223_v60_, d_224_v61_, globalState), globalState)):
                            coll27_ = coll27_.union(_dafny.Set([d_225_v62_]))
                    return _dafny.Set(coll27_)
                d_221_v58_ = _dafny.Map({d_188_cf21_: default__.fm23(p1, False, default__.fm24(-773, p3, True, p1, globalState), iife31_()
                , globalState)})
                d_226_v64_: _dafny.Map
                d_226_v64_ = _dafny.Map({d_217_v54_: p2})
                d_227_v66_: _dafny.Array
                nw19_ = _dafny.Array(None, 15)
                d_227_v66_ = nw19_
                d_228_v67_: C3
                nw20_ = C3()
                def iife32_():
                    coll28_ = _dafny.Map()
                    compr_28_: str
                    for compr_28_ in (d_146_v0_).Elements:
                        d_229_v63_: str = compr_28_
                        if (d_229_v63_) in (d_146_v0_):
                            coll28_[d_229_v63_] = p1
                    return _dafny.Map(coll28_)
                def iife33_():
                    coll29_ = _dafny.Set()
                    compr_29_: _dafny.Seq
                    for compr_29_ in (d_226_v64_).keys.Elements:
                        d_230_v65_: _dafny.Seq = compr_29_
                        if (d_230_v65_) in (d_226_v64_):
                            coll29_ = coll29_.union(_dafny.Set([d_230_v65_]))
                    return _dafny.Set(coll29_)
                nw20_.ctor__(default__.fm0((d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))], p1, p1, d_216_v53_, globalState), ((d_146_v0_).set(default__.safeIndex(len(_dafny.Map({default__.fm0(default__.fm23(p1, False, iife32_()
                , iife33_()
                , globalState), p1, p1, d_216_v53_, globalState): d_227_v66_})), len(d_146_v0_)), d_222_v59_)) + ((d_146_v0_).set(default__.safeIndex(p1, len(d_146_v0_)), d_222_v59_)), (d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))], d_215_v52_)
                d_228_v67_ = nw20_
                nw21_ = C3()
                nw21_.ctor__((d_228_v67_).f20, (d_228_v67_).f21, p3, d_215_v52_)
                d_228_v67_ = nw21_
                d_231_v68_: _dafny.Set
                d_231_v68_ = _dafny.Set({(d_228_v67_).f20, (d_228_v67_).f20})
                d_232_v69_: int
                d_233_v70_: int
                d_234_v71_: bool
                out0_: int
                out1_: int
                out2_: bool
                out0_, out1_, out2_ = (d_228_v67_).m1(d_231_v68_, (d_228_v67_).f20, globalState)
                d_232_v69_ = out0_
                d_233_v70_ = out1_
                d_234_v71_ = out2_
                d_235_v72_: D5
                d_235_v72_ = D5_DC18(False)
                d_236_v73_: _dafny.Map
                d_236_v73_ = _dafny.Map({p3: d_235_v72_})
                d_236_v73_ = (d_236_v73_).set(p3, D5_DC18((d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))]))
            elif True:
                d_237_v75_: str
                d_237_v75_ = _dafny.CodePoint('f')
                d_238_v77_: _dafny.Seq
                d_238_v77_ = _dafny.SeqWithoutIsStrInference([p2])
                d_239_v78_: _dafny.Map
                def iife34_():
                    coll30_ = _dafny.Map()
                    compr_30_: _dafny.Seq
                    for compr_30_ in (_dafny.MultiSet([d_238_v77_])).Elements:
                        d_240_v76_: _dafny.Seq = compr_30_
                        if (d_240_v76_) in (_dafny.MultiSet([d_238_v77_])):
                            coll30_[d_240_v76_] = _dafny.SeqWithoutIsStrInference([p1])
                    return _dafny.Map(coll30_)
                d_239_v78_ = _dafny.Map({len(d_146_v0_): len(iife34_()
                )})
                d_241_v79_: _dafny.Map
                d_241_v79_ = _dafny.Map({d_237_v75_: (d_239_v78_).set(len(d_238_v77_), 165)})
                d_242_v80_: _dafny.Set
                d_242_v80_ = _dafny.Set({p1, p1})
                d_243_v81_: D9
                d_243_v81_ = D9_DC30(default__.fm25(p1, p1, p3, default__.fm1(p2, p1, True, d_242_v80_, globalState), globalState), (d_147_v1_)[default__.safeIndex(719, (d_147_v1_).length(0))], p1, d_242_v80_, 961)
                d_244_v82_: _dafny.Map
                d_244_v82_ = _dafny.Map({(d_243_v81_).cf50: (0) - (len(_dafny.SeqWithoutIsStrInference([d_237_v75_ for d_245_i1_ in range(default__.abs(528))])))})
                d_246_v84_: _dafny.Map
                d_246_v84_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([101]): p1})
                index12_ = default__.safeIndex(719, (d_147_v1_).length(0))
                def iife35_():
                    coll31_ = _dafny.Map()
                    compr_31_: str
                    for compr_31_ in (d_241_v79_).keys.Elements:
                        d_247_v74_: str = compr_31_
                        if (d_247_v74_) in (d_241_v79_):
                            coll31_[d_247_v74_] = 781
                    return _dafny.Map(coll31_)
                def iife36_():
                    def iife38_():
                        coll34_ = _dafny.Map()
                        compr_34_: _dafny.Seq
                        for compr_34_ in (d_246_v84_).keys.Elements:
                            d_248_v83_: _dafny.Seq = compr_34_
                            if (d_248_v83_) in (d_246_v84_):
                                coll34_[d_248_v83_] = 540
                        return _dafny.Map(coll34_)
                    coll32_ = _dafny.Set()
                    def iife37_():
                        coll33_ = _dafny.Map()
                        compr_33_: _dafny.Seq
                        for compr_33_ in (d_246_v84_).keys.Elements:
                            d_248_v83_: _dafny.Seq = compr_33_
                            if (d_248_v83_) in (d_246_v84_):
                                coll33_[d_248_v83_] = 540
                        return _dafny.Map(coll33_)
                    compr_32_: _dafny.Seq
                    for compr_32_ in (iife37_()
                    ).keys.Elements:
                        d_249_v85_: _dafny.Seq = compr_32_
                        if (d_249_v85_) in (iife38_()
                        ):
                            coll32_ = coll32_.union(_dafny.Set([d_249_v85_]))
                    return _dafny.Set(coll32_)
                (d_147_v1_)[index12_] = default__.fm23(p1, False, (iife35_()
                ) | (d_244_v82_), iife36_()
                , globalState)
                d_250_v86_: _dafny.Array
                nw22_ = _dafny.Array(None, 1)
                nw22_[int(0)] = d_146_v0_
                d_250_v86_ = nw22_
                index13_ = default__.safeIndex(664, (d_250_v86_).length(0))
                (d_250_v86_)[index13_] = (d_146_v0_) + (_dafny.SeqWithoutIsStrInference([d_237_v75_ for d_251_i2_ in range(default__.abs(869))]))
                index14_ = default__.safeIndex(664, (d_250_v86_).length(0))
                (d_250_v86_)[index14_] = d_146_v0_
                rhs11_ = 667
                rhs12_ = d_188_cf21_
                rhs13_ = (p1) - (len(((d_250_v86_)[default__.safeIndex(664, (d_250_v86_).length(0))]) + (((d_250_v86_)[default__.safeIndex(664, (d_250_v86_).length(0))]).set(default__.safeIndex(p1, len((d_250_v86_)[default__.safeIndex(664, (d_250_v86_).length(0))])), ((d_250_v86_)[default__.safeIndex(664, (d_250_v86_).length(0))])[default__.safeIndex(len(d_219_v56_), len((d_250_v86_)[default__.safeIndex(664, (d_250_v86_).length(0))]))]))))
                lhs12_ = globalState
                lhs12_.f3 = rhs11_
                d_188_cf21_ = rhs12_
                r0 = rhs13_
                d_242_v80_ = ((d_242_v80_).intersection(d_242_v80_)) | ((_dafny.Set({p1})) - (d_242_v80_))
                d_252_v87_: D1
                d_252_v87_ = D1_DC3(p1)
                index15_ = default__.safeIndex(906, (d_220_v57_).length(0))
                (d_220_v57_)[index15_] = (d_252_v87_).cf3
                index16_ = default__.safeIndex(906, (d_220_v57_).length(0))
                (d_220_v57_)[index16_] = (default__.safeDivisionInt(p1, 474)) - (p1)
        elif True:
            d_253___mcc_h7_ = source7_.cf28
            d_254_cf28_ = d_253___mcc_h7_
            d_255_v88_: str
            d_255_v88_ = _dafny.CodePoint('i')
            d_256_v89_: _dafny.Map
            d_256_v89_ = _dafny.Map({d_255_v88_: p1})
            d_257_v90_: _dafny.MultiSet
            d_257_v90_ = _dafny.MultiSet([p3, p3])
            d_258_v91_: _dafny.Seq
            d_258_v91_ = _dafny.SeqWithoutIsStrInference([(d_257_v90_).cardinality, 555])
            d_259_v92_: _dafny.Set
            d_259_v92_ = _dafny.Set({d_258_v91_, _dafny.SeqWithoutIsStrInference([p1])})
            d_260_v93_: _dafny.Seq
            d_260_v93_ = _dafny.SeqWithoutIsStrInference([default__.fm23(len(_dafny.Set({(0) - (-800), p1})), p2, d_256_v89_, d_259_v92_, globalState)])
            d_261_v94_: _dafny.Array
            nw23_ = _dafny.Array(_dafny.CodePoint('D'), 1)
            d_261_v94_ = nw23_
            d_262_v95_: _dafny.Map
            d_262_v95_ = _dafny.Map({(_dafny.MultiSet(d_260_v93_)).cardinality: d_261_v94_})
            d_263_v96_: _dafny.Map
            d_263_v96_ = _dafny.Map({p3: (0) - (len(d_146_v0_))})
            d_264_v97_: _dafny.Seq
            d_264_v97_ = _dafny.SeqWithoutIsStrInference([d_263_v96_])
            d_265_v98_: T0
            nw24_ = C8()
            nw24_.ctor__(d_262_v95_, d_264_v97_, (d_260_v93_)[default__.safeIndex(len(d_260_v93_), len(d_260_v93_))], (d_257_v90_).set(p3, default__.abs((0) - (p1))))
            d_265_v98_ = nw24_
            d_266_v99_: D2
            d_266_v99_ = D2_DC9(d_265_v98_.f14, 19, (d_146_v0_)[default__.safeIndex(p1, len(d_146_v0_))])
            nw25_ = C3()
            nw25_.ctor__((d_266_v99_).cf14, default__.fm27(p1, d_265_v98_.f14, globalState), d_265_v98_.f14, ((d_265_v98_).f15).intersection(_dafny.MultiSet(d_260_v93_)))
            d_265_v98_ = nw25_
            (d_265_v98_).f14 = False
            d_267_v100_: _dafny.Array
            nw26_ = _dafny.Array(_dafny.Map({}), 20)
            d_267_v100_ = nw26_
            d_268_v101_: _dafny.Map
            d_268_v101_ = _dafny.Map({d_265_v98_.f14: p3})
            index17_ = default__.safeIndex(988, (d_267_v100_).length(0))
            (d_267_v100_)[index17_] = d_268_v101_
            d_269_v102_: _dafny.MultiSet
            d_269_v102_ = _dafny.MultiSet([len(d_146_v0_)])
            d_270_v103_: _dafny.Array
            nw27_ = _dafny.Array(None, 3)
            nw27_[int(0)] = -46
            nw27_[int(1)] = default__.fm0(d_265_v98_.f14, p1, p1, d_269_v102_, globalState)
            nw27_[int(2)] = (0) - (p1)
            d_270_v103_ = nw27_
            index18_ = default__.safeIndex(718, (d_270_v103_).length(0))
            (d_270_v103_)[index18_] = p1
            index19_ = default__.safeIndex(672, (d_147_v1_).length(0))
            (d_147_v1_)[index19_] = d_265_v98_.f14
            index20_ = default__.safeIndex(614, (d_261_v94_).length(0))
            (d_261_v94_)[index20_] = d_255_v88_
            d_271_v104_: _dafny.Seq
            d_271_v104_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({True: d_265_v98_.f14}), d_268_v101_, _dafny.Map({True: p2})])
            d_272_v105_: C8
            nw28_ = C8()
            nw28_.ctor__(_dafny.Map({p1: d_261_v94_}), d_264_v97_, d_265_v98_.f14, (_dafny.MultiSet([p2])).set(d_265_v98_.f14, default__.abs(p1)))
            d_272_v105_ = nw28_
            d_273_v106_: _dafny.MultiSet
            d_273_v106_ = _dafny.MultiSet([d_272_v105_])
            d_274_v107_: _dafny.Map
            d_274_v107_ = _dafny.Map({p1: d_265_v98_.f14})
            index21_ = default__.safeIndex(988, (d_267_v100_).length(0))
            index22_ = default__.safeIndex(718, (d_270_v103_).length(0))
            index23_ = default__.safeIndex(672, (d_147_v1_).length(0))
            index24_ = default__.safeIndex(614, (d_261_v94_).length(0))
            rhs14_ = (d_271_v104_)[default__.safeIndex(p1, len(d_271_v104_))]
            rhs15_ = (0) - (((0) - (((_dafny.MultiSet([(0) - ((d_273_v106_).cardinality)])).cardinality) * (p1))) + (p1))
            rhs16_ = p1
            rhs17_ = ((d_274_v107_)[len(default__.fm27(467, d_265_v98_.f14, globalState))] if (len(default__.fm27(467, d_265_v98_.f14, globalState))) in (d_274_v107_) else p2)
            rhs18_ = d_255_v88_
            lhs13_ = d_267_v100_
            lhs14_ = default__.safeIndex(988, (d_267_v100_).length(0))
            lhs15_ = d_270_v103_
            lhs16_ = default__.safeIndex(718, (d_270_v103_).length(0))
            lhs17_ = d_147_v1_
            lhs18_ = default__.safeIndex(672, (d_147_v1_).length(0))
            lhs19_ = d_261_v94_
            lhs20_ = default__.safeIndex(614, (d_261_v94_).length(0))
            lhs13_[lhs14_] = rhs14_
            r0 = rhs15_
            lhs15_[lhs16_] = rhs16_
            lhs17_[lhs18_] = rhs17_
            lhs19_[lhs20_] = rhs18_
            if not((d_147_v1_)[default__.safeIndex(672, (d_147_v1_).length(0))]):
                d_275_v108_: C2
                nw29_ = C2()
                nw29_.ctor__(d_274_v107_)
                d_275_v108_ = nw29_
                index25_ = default__.safeIndex(718, (d_270_v103_).length(0))
                (d_270_v103_)[index25_] = (d_258_v91_)[default__.safeIndex((d_270_v103_)[default__.safeIndex(718, (d_270_v103_).length(0))], len(d_258_v91_))]
                d_276_v109_: _dafny.Array
                nw30_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_276_v109_ = nw30_
                d_277_v110_: _dafny.Map
                d_277_v110_ = _dafny.Map({(d_270_v103_)[default__.safeIndex(718, (d_270_v103_).length(0))]: d_270_v103_})
                d_278_v111_: _dafny.Map
                d_278_v111_ = _dafny.Map({p2: d_270_v103_})
                d_279_v112_: _dafny.Array
                nw31_ = _dafny.Array(None, 25)
                nw31_[int(0)] = d_270_v103_
                nw31_[int(1)] = d_270_v103_
                nw31_[int(2)] = d_270_v103_
                nw31_[int(3)] = d_270_v103_
                nw31_[int(4)] = d_270_v103_
                nw31_[int(5)] = d_270_v103_
                nw31_[int(6)] = d_270_v103_
                nw31_[int(7)] = d_270_v103_
                nw31_[int(8)] = d_270_v103_
                nw31_[int(9)] = d_270_v103_
                nw31_[int(10)] = d_270_v103_
                nw31_[int(11)] = d_270_v103_
                nw31_[int(12)] = d_270_v103_
                nw31_[int(13)] = d_270_v103_
                nw31_[int(14)] = d_270_v103_
                nw31_[int(15)] = d_270_v103_
                nw31_[int(16)] = ((d_277_v110_)[p1] if (p1) in (d_277_v110_) else d_270_v103_)
                nw31_[int(17)] = d_270_v103_
                nw31_[int(18)] = d_270_v103_
                nw31_[int(19)] = ((d_278_v111_)[d_265_v98_.f14] if (d_265_v98_.f14) in (d_278_v111_) else d_270_v103_)
                nw31_[int(20)] = d_270_v103_
                nw31_[int(21)] = d_270_v103_
                nw31_[int(22)] = d_270_v103_
                nw31_[int(23)] = d_270_v103_
                nw31_[int(24)] = d_270_v103_
                d_279_v112_ = nw31_
                index26_ = default__.safeIndex(822, (d_276_v109_).length(0))
                (d_276_v109_)[index26_] = d_279_v112_
                index27_ = default__.safeIndex(822, (d_276_v109_).length(0))
                rhs19_ = d_279_v112_
                rhs20_ = (d_270_v103_)[default__.safeIndex(718, (d_270_v103_).length(0))]
                rhs21_ = default__.safeDivisionInt(p1, (d_270_v103_)[default__.safeIndex(718, (d_270_v103_).length(0))])
                lhs21_ = d_276_v109_
                lhs22_ = default__.safeIndex(822, (d_276_v109_).length(0))
                lhs23_ = globalState
                lhs24_ = globalState
                lhs21_[lhs22_] = rhs19_
                lhs23_.f3 = rhs20_
                lhs24_.f3 = rhs21_
                d_280_v113_: C1
                nw32_ = C1()
                nw32_.ctor__()
                d_280_v113_ = nw32_
                d_281_v114_: D0
                d_281_v114_ = D0_DC1(p3)
                d_281_v114_ = d_281_v114_
            elif True:
                d_282_v115_: _dafny.Array
                nw33_ = _dafny.Array(False, 12)
                d_282_v115_ = nw33_
                d_283_v116_: _dafny.Map
                d_283_v116_ = _dafny.Map({p1: d_282_v115_})
                (globalState).f2 = ((d_283_v116_)[((d_270_v103_)[default__.safeIndex(718, (d_270_v103_).length(0))] if d_265_v98_.f14 else -369)] if (((d_270_v103_)[default__.safeIndex(718, (d_270_v103_).length(0))] if d_265_v98_.f14 else -369)) in (d_283_v116_) else d_282_v115_)
                d_284_v117_: _dafny.Map
                d_284_v117_ = _dafny.Map({_dafny.Map({(d_270_v103_)[default__.safeIndex(718, (d_270_v103_).length(0))]: d_265_v98_.f14}): d_265_v98_.f14})
                d_285_v118_: _dafny.Map
                d_285_v118_ = _dafny.Map({len(d_284_v117_): d_146_v0_})
                d_286_v119_: D8
                d_286_v119_ = D8_DC26(default__.fm23(p1, d_265_v98_.f14, d_256_v89_, d_259_v92_, globalState), d_285_v118_)
                d_287_v120_: D1
                d_287_v120_ = D1_DC3(808)
                d_288_v121_: D1
                d_288_v121_ = D1_DC6(d_287_v120_)
                d_289_v122_: D1
                d_289_v122_ = D1_DC6(d_287_v120_)
                d_290_v123_: _dafny.Map
                d_290_v123_ = _dafny.Map({d_286_v119_: d_288_v121_})
                d_291_v124_: D1
                d_291_v124_ = D1_DC6(((d_290_v123_)[d_286_v119_] if (d_286_v119_) in (d_290_v123_) else d_287_v120_))
                d_292_v125_: _dafny.Array
                nw34_ = _dafny.Array(None, 1)
                nw34_[int(0)] = d_291_v124_
                d_292_v125_ = nw34_
                index28_ = default__.safeIndex(26, (d_292_v125_).length(0))
                (d_292_v125_)[index28_] = d_291_v124_
                index29_ = default__.safeIndex(26, (d_292_v125_).length(0))
                (d_292_v125_)[index29_] = d_291_v124_
                index30_ = default__.safeIndex(718, (d_270_v103_).length(0))
                rhs22_ = default__.safeModuloInt(667, default__.fm0(d_265_v98_.f14, 173, p1, d_269_v102_, globalState))
                rhs23_ = p3
                lhs25_ = d_270_v103_
                lhs26_ = default__.safeIndex(718, (d_270_v103_).length(0))
                lhs27_ = d_265_v98_
                lhs25_[lhs26_] = rhs22_
                lhs27_.f14 = rhs23_
                (globalState).f2 = d_282_v115_
                d_293_v126_: _dafny.Set
                d_293_v126_ = _dafny.Set({len(d_263_v96_)})
                rhs24_ = d_265_v98_
                rhs25_ = (0) - (default__.safeDivisionInt((len(d_274_v107_)) + ((d_270_v103_)[default__.safeIndex(718, (d_270_v103_).length(0))]), p1))
                rhs26_ = (d_261_v94_)[default__.safeIndex(614, (d_261_v94_).length(0))]
                rhs27_ = (d_293_v126_) | ((default__.fm6(globalState)) | (d_293_v126_))
                lhs28_ = globalState
                d_265_v98_ = rhs24_
                lhs28_.f3 = rhs25_
                d_255_v88_ = rhs26_
                d_293_v126_ = rhs27_
        d_294_v127_: C7
        nw35_ = C7()
        nw35_.ctor__()
        d_294_v127_ = nw35_
        d_294_v127_ = d_294_v127_
        r0 = (0) - (p1)
        d_146_v0_ = (d_148_v2_).cf22
        d_295_v128_: _dafny.Map
        d_295_v128_ = _dafny.Map({p3: p1})
        d_296_v129_: C4
        nw36_ = C4()
        nw36_.ctor__(p3, p3, _dafny.MultiSet([p2]))
        d_296_v129_ = nw36_
        d_297_v130_: D20
        d_297_v130_ = D20_DC60(d_296_v129_)
        d_298_v131_: _dafny.MultiSet
        d_298_v131_ = _dafny.MultiSet([p1, len(default__.fm7(globalState)), p1, len(d_295_v128_), len(_dafny.Map({(d_297_v130_).cf99: p3}))])
        d_299_v132_: _dafny.Map
        d_299_v132_ = _dafny.Map({(0) - (default__.safeDivisionInt(p1, (0) - (p1))): (d_294_v127_).fm10(d_295_v128_, d_298_v131_, (d_296_v129_).f19, d_146_v0_, globalState)})
        if ((d_299_v132_)[default__.safeDivisionInt(p1, p1)] if (default__.safeDivisionInt(p1, p1)) in (d_299_v132_) else p3):
            (globalState).f3 = (p1) + (p1)
            d_300_v133_: _dafny.MultiSet
            d_300_v133_ = _dafny.MultiSet([p2, not(p2), (d_296_v129_).f19])
            d_301_v134_: C5
            nw37_ = C5()
            nw37_.ctor__(709, True, d_300_v133_)
            d_301_v134_ = nw37_
            d_301_v134_ = (d_301_v134_ if (d_296_v129_).f19 else d_301_v134_)
            d_302_v135_: C0
            nw38_ = C0()
            nw38_.ctor__((d_296_v129_).f19, p2)
            d_302_v135_ = nw38_
            (globalState).f9 = not ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbhurb"))) < ((d_148_v2_).cf22)) or ((d_302_v135_).f23)
            d_303_v136_: D1
            d_303_v136_ = D1_DC5((d_296_v129_).f19, (d_302_v135_).f23)
            d_304_v137_: D1
            d_304_v137_ = D1_DC6(D1_DC6(d_303_v136_))
            pat_let_tv0_ = d_303_v136_
            d_305_v138_: _dafny.Map
            def iife39_(_pat_let2_0):
                def iife40_(d_306_dt__update__tmp_h1_):
                    def iife41_(_pat_let3_0):
                        def iife42_(d_307_dt__update_hcf8_h0_):
                            return D1_DC6(d_307_dt__update_hcf8_h0_)
                        return iife42_(_pat_let3_0)
                    return iife41_(D1_DC6(pat_let_tv0_))
                return iife40_(_pat_let2_0)
            d_305_v138_ = _dafny.Map({iife39_(d_304_v137_): p3})
            d_305_v138_ = (d_305_v138_).set(d_304_v137_, p3)
        elif True:
            d_308_v139_: _dafny.Array
            def lambda0_(d_309_p1_):
                def lambda1_(d_310_i3_):
                    return default__.safeDivisionInt(d_310_i3_, d_309_p1_)

                return lambda1_

            init0_ = lambda0_(p1)
            nw39_ = _dafny.Array(None, 4)
            for i0_0_ in range(nw39_.length(0)):
                nw39_[i0_0_] = init0_(i0_0_)
            d_308_v139_ = nw39_
            index31_ = default__.safeIndex(520, (d_308_v139_).length(0))
            (d_308_v139_)[index31_] = p1
            d_311_v140_: _dafny.Seq
            d_311_v140_ = _dafny.SeqWithoutIsStrInference([(421) + (p1)])
            index32_ = default__.safeIndex(520, (d_308_v139_).length(0))
            (d_308_v139_)[index32_] = len(d_311_v140_)
            d_312_v141_: _dafny.Set
            d_312_v141_ = _dafny.Set({(d_294_v127_).fm10(d_295_v128_, d_298_v131_, (d_296_v129_).f19, d_146_v0_, globalState), p2})
            (globalState).f3 = (default__.safeModuloInt(p1, p1)) - (default__.safeModuloInt((d_308_v139_)[default__.safeIndex(520, (d_308_v139_).length(0))], (0) - (len(d_312_v141_))))
            (globalState).f9 = (((d_296_v129_).f19 if (d_296_v129_).f19 else True)) or ((d_296_v129_).f19)
            d_313_v142_: D0
            d_313_v142_ = D0_DC0(p2)
            d_314_v143_: _dafny.MultiSet
            d_314_v143_ = _dafny.MultiSet([(d_313_v142_).cf0, (d_296_v129_).f19])
            d_315_v144_: T0
            nw40_ = C3()
            nw40_.ctor__((d_308_v139_)[default__.safeIndex(520, (d_308_v139_).length(0))], d_146_v0_, (d_296_v129_).f19, d_314_v143_)
            d_315_v144_ = nw40_
            d_316_v145_: C6
            nw41_ = C6()
            nw41_.ctor__(p3, (d_315_v144_).f15)
            d_316_v145_ = nw41_
            d_317_v146_: D21
            d_317_v146_ = D21_DC63(d_316_v145_)
            d_318_v147_: _dafny.Map
            d_318_v147_ = _dafny.Map({d_315_v144_: (d_317_v146_).cf101})
            d_318_v147_ = (d_318_v147_).set(d_315_v144_, d_316_v145_)
            d_319_v148_: str
            d_319_v148_ = _dafny.CodePoint('o')
            d_320_v149_: _dafny.Set
            d_320_v149_ = _dafny.Set({(d_315_v144_).f15})
            rhs28_ = ((d_315_v144_).f15) in (d_320_v149_)
            rhs29_ = d_319_v148_
            lhs29_ = globalState
            lhs29_.f9 = rhs28_
            d_319_v148_ = rhs29_
        d_321_v150_: D5
        d_321_v150_ = D5_DC18((d_294_v127_).fm10(d_295_v128_, d_298_v131_, p2, d_146_v0_, globalState))
        pat_let_tv1_ = p3
        pat_let_tv2_ = p3
        pat_let_tv3_ = p1
        pat_let_tv4_ = p1
        pat_let_tv5_ = p1
        pat_let_tv6_ = p2
        pat_let_tv7_ = d_296_v129_
        pat_let_tv8_ = d_296_v129_
        pat_let_tv9_ = d_296_v129_
        pat_let_tv10_ = p2
        def lambda2_(source8_):
            if source8_.is_DC17:
                d_322___mcc_h8_ = source8_.cf30
                d_323___mcc_h9_ = source8_.cf31
                d_324___mcc_h10_ = source8_.cf32
                d_325_cf32_ = d_324___mcc_h10_
                d_326_cf31_ = d_323___mcc_h9_
                d_327_cf30_ = d_322___mcc_h8_
                return pat_let_tv1_
            elif source8_.is_DC18:
                d_328___mcc_h11_ = source8_.cf33
                d_329_cf33_ = d_328___mcc_h11_
                return pat_let_tv2_
            elif source8_.is_DC16:
                d_330___mcc_h12_ = source8_.cf29
                d_331_cf29_ = d_330___mcc_h12_
                return (459) <= ((0) - (len(_dafny.Map({pat_let_tv3_: (D1_DC4(pat_let_tv4_, pat_let_tv5_)).cf4}))))
            elif True:
                d_332___mcc_h13_ = source8_.cf34
                d_333_cf34_ = d_332___mcc_h13_
                return (_dafny.SeqWithoutIsStrInference([pat_let_tv6_])) != (_dafny.SeqWithoutIsStrInference([(pat_let_tv7_).f19, (pat_let_tv8_).f19, (pat_let_tv9_).f19, pat_let_tv10_]))

        (globalState).f9 = lambda2_(d_321_v150_)
        d_334_v151_: _dafny.Seq
        d_334_v151_ = _dafny.SeqWithoutIsStrInference([(d_296_v129_).f19])
        r0 = (default__.safeModuloInt(len(d_334_v151_), p1)) * (p1)
        d_335_v152_: _dafny.Array
        nw42_ = _dafny.Array(None, 13)
        d_335_v152_ = nw42_
        d_336_v153_: _dafny.Map
        d_336_v153_ = _dafny.Map({p1: d_335_v152_})
        r1 = d_336_v153_
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_337_v0_: int
        d_337_v0_ = 206
        d_338_v1_: _dafny.MultiSet
        d_338_v1_ = _dafny.MultiSet([d_337_v0_])
        d_339_v2_: _dafny.Array
        nw43_ = _dafny.Array(False, 3)
        d_339_v2_ = nw43_
        d_340_v3_: _dafny.Seq
        d_340_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwiunjo"))
        d_341_v4_: bool
        d_341_v4_ = True
        d_342_v5_: _dafny.Set
        d_342_v5_ = _dafny.Set({d_341_v4_, d_341_v4_})
        d_343_v6_: _dafny.Array
        nw44_ = _dafny.Array(_dafny.Set({}), 12)
        d_343_v6_ = nw44_
        d_344_v7_: _dafny.Array
        def lambda3_(d_345_v3_):
            def lambda4_(d_346_i0_):
                return d_345_v3_

            return lambda4_

        init1_ = lambda3_(d_340_v3_)
        nw45_ = _dafny.Array(None, 11)
        for i0_1_ in range(nw45_.length(0)):
            nw45_[i0_1_] = init1_(i0_1_)
        d_344_v7_ = nw45_
        d_347_globalState_: GlobalState
        nw46_ = GlobalState()
        nw46_.ctor__(946, (d_338_v1_) - (d_338_v1_), d_339_v2_, 664, True, d_340_v3_, d_342_v5_, False, 762, False, -427, d_343_v6_, d_344_v7_, _dafny.Set({d_341_v4_, d_341_v4_, False}))
        d_347_globalState_ = nw46_
        d_348_v8_: _dafny.Seq
        d_348_v8_ = _dafny.SeqWithoutIsStrInference([d_341_v4_])
        d_348_v8_ = (d_348_v8_) + (d_348_v8_)
        d_349_v9_: _dafny.Map
        d_349_v9_ = _dafny.Map({(d_337_v0_) != (d_337_v0_): d_342_v5_})
        d_349_v9_ = (d_349_v9_).set(d_341_v4_, d_342_v5_)
        d_350_v10_: _dafny.Seq
        d_350_v10_ = _dafny.SeqWithoutIsStrInference([d_337_v0_, d_337_v0_])
        hi0_ = d_337_v0_
        for d_351_i1_ in range(default__.fm0(d_341_v4_, d_337_v0_, d_337_v0_, (_dafny.MultiSet(d_350_v10_)).set(d_337_v0_, default__.abs(d_337_v0_)), d_347_globalState_), hi0_):
            d_352_v11_: _dafny.Map
            d_352_v11_ = _dafny.Map({d_337_v0_: d_344_v7_})
            d_344_v7_ = ((d_352_v11_)[(d_337_v0_) * (-478)] if ((d_337_v0_) * (-478)) in (d_352_v11_) else d_344_v7_)
            d_353_v12_: _dafny.Map
            d_353_v12_ = _dafny.Map({d_341_v4_: default__.safeModuloInt(d_351_i1_, d_351_i1_)})
            d_354_v13_: _dafny.Array
            nw47_ = _dafny.Array(int(0), 7)
            d_354_v13_ = nw47_
            d_355_v14_: _dafny.MultiSet
            d_355_v14_ = _dafny.MultiSet([d_354_v13_])
            d_356_v15_: _dafny.Map
            d_356_v15_ = _dafny.Map({d_337_v0_: d_354_v13_})
            d_357_v16_: _dafny.Set
            d_357_v16_ = _dafny.Set({len(d_350_v10_), d_337_v0_, d_337_v0_})
            rhs30_ = default__.safeModuloInt((d_351_i1_) - (d_351_i1_), default__.safeDivisionInt(d_337_v0_, d_337_v0_))
            rhs31_ = len((d_340_v3_) + ((d_340_v3_) + (d_340_v3_)))
            rhs32_ = (0) - ((0) - ((((_dafny.MultiSet([d_354_v13_]) if d_341_v4_ else d_355_v14_)).set(((d_356_v15_)[d_351_i1_] if (d_351_i1_) in (d_356_v15_) else d_354_v13_), default__.abs((0) - (d_337_v0_)))).cardinality))
            rhs33_ = (d_353_v12_) | (default__.fm1(False, len(_dafny.Map({d_351_i1_: d_341_v4_})), d_341_v4_, d_357_v16_, d_347_globalState_))
            lhs30_ = d_347_globalState_
            lhs31_ = d_347_globalState_
            d_337_v0_ = rhs30_
            lhs30_.f3 = rhs31_
            lhs31_.f3 = rhs32_
            d_353_v12_ = rhs33_
            d_358_v17_: _dafny.Map
            d_358_v17_ = _dafny.Map({d_337_v0_: d_341_v4_})
            if ((d_358_v17_)[d_337_v0_] if (d_337_v0_) in (d_358_v17_) else True):
                index33_ = default__.safeIndex(0, (d_354_v13_).length(0))
                (d_354_v13_)[index33_] = d_337_v0_
                index34_ = default__.safeIndex(0, (d_354_v13_).length(0))
                (d_354_v13_)[index34_] = 971
                rhs34_ = d_340_v3_
                rhs35_ = d_341_v4_
                lhs32_ = d_347_globalState_
                d_340_v3_ = rhs34_
                lhs32_.f9 = rhs35_
                (d_347_globalState_).f3 = (0) - (-824)
                d_341_v4_ = d_341_v4_
                d_359_v18_: C1
                nw48_ = C1()
                nw48_.ctor__()
                d_359_v18_ = nw48_
            elif True:
                d_360_v19_: str
                d_360_v19_ = _dafny.CodePoint('r')
                d_361_v20_: _dafny.Map
                d_361_v20_ = _dafny.Map({d_360_v19_: d_337_v0_})
                d_361_v20_ = (d_361_v20_).set(d_360_v19_, d_337_v0_)
                d_362_v21_: D5
                d_362_v21_ = D5_DC16(d_354_v13_)
                pat_let_tv11_ = d_354_v13_
                pat_let_tv12_ = d_354_v13_
                pat_let_tv13_ = d_354_v13_
                d_363_v22_: _dafny.Array
                nw49_ = _dafny.Array(None, 15)
                nw49_[int(0)] = d_362_v21_
                nw49_[int(1)] = d_362_v21_
                nw49_[int(2)] = d_362_v21_
                nw49_[int(3)] = d_362_v21_
                nw49_[int(4)] = d_362_v21_
                def iife43_(_pat_let4_0):
                    def iife44_(d_364_dt__update__tmp_h0_):
                        def iife45_(_pat_let5_0):
                            def iife46_(d_365_dt__update_hcf29_h0_):
                                return D5_DC16(d_365_dt__update_hcf29_h0_)
                            return iife46_(_pat_let5_0)
                        return iife45_(pat_let_tv11_)
                    return iife44_(_pat_let4_0)
                nw49_[int(5)] = iife43_(d_362_v21_)
                def iife48_(_pat_let7_0):
                    def iife49_(d_366_dt__update__tmp_h1_):
                        def iife50_(_pat_let8_0):
                            def iife51_(d_367_dt__update_hcf29_h1_):
                                return D5_DC16(d_367_dt__update_hcf29_h1_)
                            return iife51_(_pat_let8_0)
                        return iife50_(pat_let_tv12_)
                    return iife49_(_pat_let7_0)
                def iife47_(_pat_let6_0):
                    def iife52_(d_368_dt__update__tmp_h2_):
                        def iife53_(_pat_let9_0):
                            def iife54_(d_369_dt__update_hcf29_h2_):
                                return D5_DC16(d_369_dt__update_hcf29_h2_)
                            return iife54_(_pat_let9_0)
                        return iife53_(pat_let_tv13_)
                    return iife52_(_pat_let6_0)
                nw49_[int(6)] = iife47_(iife48_(d_362_v21_))
                nw49_[int(7)] = d_362_v21_
                nw49_[int(8)] = d_362_v21_
                nw49_[int(9)] = d_362_v21_
                nw49_[int(10)] = D5_DC16(d_354_v13_)
                nw49_[int(11)] = d_362_v21_
                nw49_[int(12)] = D5_DC16(d_354_v13_)
                nw49_[int(13)] = d_362_v21_
                nw49_[int(14)] = d_362_v21_
                d_363_v22_ = nw49_
                index35_ = default__.safeIndex(543, (d_363_v22_).length(0))
                (d_363_v22_)[index35_] = d_362_v21_
                d_370_v23_: _dafny.MultiSet
                d_370_v23_ = _dafny.MultiSet([(d_341_v4_ if d_341_v4_ else False), (d_337_v0_) not in (_dafny.SeqWithoutIsStrInference([d_351_i1_])), d_341_v4_, (d_351_i1_) >= (d_337_v0_), d_341_v4_])
                index36_ = default__.safeIndex(804, (d_339_v2_).length(0))
                (d_339_v2_)[index36_] = d_341_v4_
                index37_ = default__.safeIndex(543, (d_363_v22_).length(0))
                index38_ = default__.safeIndex(804, (d_339_v2_).length(0))
                rhs36_ = (0) - (d_337_v0_)
                rhs37_ = d_362_v21_
                rhs38_ = ((d_370_v23_ if d_341_v4_ else d_370_v23_)).intersection(d_370_v23_)
                rhs39_ = d_341_v4_
                lhs33_ = d_347_globalState_
                lhs34_ = d_363_v22_
                lhs35_ = default__.safeIndex(543, (d_363_v22_).length(0))
                lhs36_ = d_339_v2_
                lhs37_ = default__.safeIndex(804, (d_339_v2_).length(0))
                lhs33_.f3 = rhs36_
                lhs34_[lhs35_] = rhs37_
                d_370_v23_ = rhs38_
                lhs36_[lhs37_] = rhs39_
                (d_347_globalState_).f3 = (default__.safeModuloInt(d_337_v0_, d_351_i1_) if d_341_v4_ else d_351_i1_)
                d_371_v24_: _dafny.Map
                d_371_v24_ = _dafny.Map({(0) - (d_351_i1_): d_337_v0_})
                d_371_v24_ = (d_371_v24_).set((d_337_v0_) + (d_351_i1_), d_337_v0_)
                d_341_v4_ = d_341_v4_
            d_372_v25_: T0
            nw50_ = C5()
            nw50_.ctor__(len(d_348_v8_), d_341_v4_, _dafny.MultiSet([False, d_341_v4_]))
            d_372_v25_ = nw50_
            d_373_v26_: _dafny.Seq
            d_373_v26_ = _dafny.SeqWithoutIsStrInference([d_372_v25_, d_372_v25_])
            d_374_v27_: _dafny.Array
            nw51_ = _dafny.Array(None, 24)
            nw51_[int(0)] = d_372_v25_
            nw51_[int(1)] = d_372_v25_
            nw51_[int(2)] = d_372_v25_
            nw51_[int(3)] = d_372_v25_
            nw51_[int(4)] = (d_373_v26_)[default__.safeIndex(d_351_i1_, len(d_373_v26_))]
            nw51_[int(5)] = d_372_v25_
            nw51_[int(6)] = d_372_v25_
            nw51_[int(7)] = d_372_v25_
            nw51_[int(8)] = d_372_v25_
            nw51_[int(9)] = d_372_v25_
            nw51_[int(10)] = d_372_v25_
            nw51_[int(11)] = d_372_v25_
            nw51_[int(12)] = d_372_v25_
            nw51_[int(13)] = d_372_v25_
            nw51_[int(14)] = d_372_v25_
            nw51_[int(15)] = d_372_v25_
            nw51_[int(16)] = d_372_v25_
            nw51_[int(17)] = d_372_v25_
            nw51_[int(18)] = d_372_v25_
            nw51_[int(19)] = d_372_v25_
            nw51_[int(20)] = d_372_v25_
            nw51_[int(21)] = d_372_v25_
            nw51_[int(22)] = d_372_v25_
            nw51_[int(23)] = d_372_v25_
            d_374_v27_ = nw51_
            d_375_v28_: _dafny.MultiSet
            d_375_v28_ = _dafny.MultiSet([d_374_v27_, d_374_v27_, d_374_v27_])
            index39_ = default__.safeIndex(802, (d_339_v2_).length(0))
            (d_339_v2_)[index39_] = not((d_375_v28_).issubset(d_375_v28_))
            index40_ = default__.safeIndex(802, (d_339_v2_).length(0))
            (d_339_v2_)[index40_] = d_341_v4_
        d_376_v29_: D13
        d_376_v29_ = D13_DC43(default__.fm51(d_341_v4_, d_341_v4_, d_341_v4_, d_337_v0_, d_347_globalState_))
        d_377_v30_: _dafny.Map
        d_377_v30_ = _dafny.Map({d_337_v0_: d_337_v0_})
        d_378_v31_: str
        d_378_v31_ = _dafny.CodePoint('m')
        d_379_v32_: int
        d_380_v33_: _dafny.Map
        out3_: int
        out4_: _dafny.Map
        out3_, out4_ = default__.m11(d_376_v29_, len((d_377_v30_).set(len((d_340_v3_).set(default__.safeIndex(d_337_v0_, len(d_340_v3_)), d_378_v31_)), d_337_v0_)), d_341_v4_, d_341_v4_, d_347_globalState_)
        d_379_v32_ = out3_
        d_380_v33_ = out4_
        if (d_341_v4_ if not(not((True) or (d_341_v4_))) else (d_379_v32_) <= (d_379_v32_)):
            d_381_v34_: _dafny.Array
            nw52_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_381_v34_ = nw52_
            d_382_v35_: D16
            d_382_v35_ = D16_DC52(d_337_v0_, True, d_381_v34_, d_341_v4_, d_337_v0_)
            d_383_v36_: D2
            d_383_v36_ = D2_DC9(d_341_v4_, d_337_v0_, _dafny.CodePoint('l'))
            d_384_v37_: int
            d_385_v38_: _dafny.Map
            out5_: int
            out6_: _dafny.Map
            out5_, out6_ = default__.m11(d_376_v29_, (d_379_v32_) * ((d_382_v35_).cf88), ((d_383_v36_).cf15) in (d_340_v3_), d_341_v4_, d_347_globalState_)
            d_384_v37_ = out5_
            d_385_v38_ = out6_
            d_386_v39_: _dafny.Array
            def lambda5_(d_387_v0_):
                def lambda6_(d_388_i2_):
                    return default__.safeModuloInt(d_388_i2_, (D1_DC4(d_387_v0_, 0)).cf5)

                return lambda6_

            init2_ = lambda5_(d_337_v0_)
            nw53_ = _dafny.Array(None, 15)
            for i0_2_ in range(nw53_.length(0)):
                nw53_[i0_2_] = init2_(i0_2_)
            d_386_v39_ = nw53_
            index41_ = default__.safeIndex(190, (d_386_v39_).length(0))
            (d_386_v39_)[index41_] = d_379_v32_
            index42_ = default__.safeIndex(190, (d_386_v39_).length(0))
            (d_386_v39_)[index42_] = (0) - (d_379_v32_)
            if d_341_v4_:
                d_389_v40_: _dafny.Map
                d_389_v40_ = _dafny.Map({d_341_v4_: d_348_v8_})
                d_390_v41_: _dafny.Map
                d_390_v41_ = _dafny.Map({d_378_v31_: (d_338_v1_).cardinality})
                d_391_v42_: _dafny.Set
                d_391_v42_ = _dafny.Set({d_350_v10_})
                d_392_v43_: _dafny.Map
                d_392_v43_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pshth")): _dafny.MultiSet(d_348_v8_)})
                d_393_v44_: _dafny.MultiSet
                d_393_v44_ = _dafny.MultiSet([d_341_v4_])
                d_394_v45_: _dafny.Map
                d_394_v45_ = _dafny.Map({d_341_v4_: (d_386_v39_)[default__.safeIndex(190, (d_386_v39_).length(0))]})
                d_395_v46_: _dafny.Map
                d_395_v46_ = _dafny.Map({((d_394_v45_)[d_341_v4_] if (d_341_v4_) in (d_394_v45_) else len((d_350_v10_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_396_i4_ in range(default__.abs(-625))])), len(d_350_v10_)), d_384_v37_))): d_341_v4_})
                d_397_v47_: D1
                d_397_v47_ = D1_DC5(d_341_v4_, False)
                pat_let_tv14_ = d_341_v4_
                d_398_v48_: D11
                def iife55_(_pat_let10_0):
                    def iife56_(d_399_dt__update__tmp_h3_):
                        def iife57_(_pat_let11_0):
                            def iife58_(d_400_dt__update_hcf7_h0_):
                                return D1_DC5((d_399_dt__update__tmp_h3_).cf6, d_400_dt__update_hcf7_h0_)
                            return iife58_(_pat_let11_0)
                        return iife57_(pat_let_tv14_)
                    return iife56_(_pat_let10_0)
                d_398_v48_ = D11_DC34(d_393_v44_, d_338_v1_, d_339_v2_, iife55_(d_397_v47_))
                pat_let_tv15_ = d_393_v44_
                pat_let_tv16_ = d_339_v2_
                d_401_v49_: _dafny.Array
                nw54_ = _dafny.Array(None, 26)
                nw54_[int(0)] = _dafny.MultiSet(((d_389_v40_)[default__.fm23(len(d_348_v8_), d_341_v4_, d_390_v41_, d_391_v42_, d_347_globalState_)] if (default__.fm23(len(d_348_v8_), d_341_v4_, d_390_v41_, d_391_v42_, d_347_globalState_)) in (d_389_v40_) else (d_348_v8_).set(default__.safeIndex((d_386_v39_)[default__.safeIndex(190, (d_386_v39_).length(0))], len(d_348_v8_)), default__.fm23(len(d_348_v8_), not(d_341_v4_), default__.fm24(d_384_v37_, d_341_v4_, d_341_v4_, (d_386_v39_)[default__.safeIndex(190, (d_386_v39_).length(0))], d_347_globalState_), d_391_v42_, d_347_globalState_))))
                nw54_[int(1)] = (((d_392_v43_)[_dafny.SeqWithoutIsStrInference([d_378_v31_ for d_402_i3_ in range(default__.abs(-948))])] if (_dafny.SeqWithoutIsStrInference([d_378_v31_ for d_403_i3_ in range(default__.abs(-948))])) in (d_392_v43_) else d_393_v44_)).intersection(_dafny.MultiSet([d_341_v4_]))
                nw54_[int(2)] = d_393_v44_
                nw54_[int(3)] = d_393_v44_
                nw54_[int(4)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(d_341_v4_), d_341_v4_]))
                nw54_[int(5)] = (default__.fm33(not(False), d_347_globalState_)) | (d_393_v44_)
                nw54_[int(6)] = (d_393_v44_).intersection(d_393_v44_)
                nw54_[int(7)] = _dafny.MultiSet(d_348_v8_)
                nw54_[int(8)] = d_393_v44_
                nw54_[int(9)] = _dafny.MultiSet([False])
                nw54_[int(10)] = d_393_v44_
                nw54_[int(11)] = (d_393_v44_).set(d_341_v4_, default__.abs(d_379_v32_))
                nw54_[int(12)] = _dafny.MultiSet([d_341_v4_, not(d_341_v4_), ((d_395_v46_)[d_384_v37_] if (d_384_v37_) in (d_395_v46_) else False)])
                nw54_[int(13)] = d_393_v44_
                nw54_[int(14)] = _dafny.MultiSet([True])
                nw54_[int(15)] = ((d_393_v44_).set(d_341_v4_, default__.abs(d_384_v37_))).set(d_341_v4_, default__.abs(d_337_v0_))
                nw54_[int(16)] = (d_393_v44_) | (d_393_v44_)
                nw54_[int(17)] = d_393_v44_
                nw54_[int(18)] = d_393_v44_
                nw54_[int(19)] = d_393_v44_
                nw54_[int(20)] = d_393_v44_
                nw54_[int(21)] = _dafny.MultiSet([d_341_v4_, not(d_341_v4_), True, d_341_v4_])
                nw54_[int(22)] = d_393_v44_
                def iife59_(_pat_let12_0):
                    def iife60_(d_404_dt__update__tmp_h4_):
                        def iife61_(_pat_let13_0):
                            def iife62_(d_405_dt__update_hcf58_h0_):
                                def iife63_(_pat_let14_0):
                                    def iife64_(d_406_dt__update_hcf60_h0_):
                                        return D11_DC34(d_405_dt__update_hcf58_h0_, (d_404_dt__update__tmp_h4_).cf59, d_406_dt__update_hcf60_h0_, (d_404_dt__update__tmp_h4_).cf61)
                                    return iife64_(_pat_let14_0)
                                return iife63_(pat_let_tv16_)
                            return iife62_(_pat_let13_0)
                        return iife61_(pat_let_tv15_)
                    return iife60_(_pat_let12_0)
                nw54_[int(23)] = (iife59_(d_398_v48_)).cf58
                nw54_[int(24)] = (d_393_v44_) - (d_393_v44_)
                nw54_[int(25)] = d_393_v44_
                d_401_v49_ = nw54_
                d_407_v50_: _dafny.Array
                nw55_ = _dafny.Array(None, 9)
                d_407_v50_ = nw55_
                d_408_v51_: _dafny.Map
                d_408_v51_ = _dafny.Map({d_407_v50_: d_384_v37_})
                d_409_v52_: _dafny.Array
                def lambda7_(d_410_v31_):
                    def lambda8_(d_411_i5_):
                        return d_410_v31_

                    return lambda8_

                init3_ = lambda7_(d_378_v31_)
                nw56_ = _dafny.Array(None, 1)
                for i0_3_ in range(nw56_.length(0)):
                    nw56_[i0_3_] = init3_(i0_3_)
                d_409_v52_ = nw56_
                d_412_v53_: D16
                d_412_v53_ = D16_DC51(d_409_v52_)
                rhs40_ = default__.fm0((d_379_v32_) >= (len(d_408_v51_)), (0) - (d_379_v32_), d_379_v32_, d_338_v1_, d_347_globalState_)
                rhs41_ = d_384_v37_
                rhs42_ = d_378_v31_
                rhs43_ = d_401_v49_
                rhs44_ = (d_348_v8_)[default__.safeIndex(((_dafny.MultiSet([d_409_v52_])).set((d_412_v53_).cf83, default__.abs(35))).cardinality, len(d_348_v8_))]
                lhs38_ = d_347_globalState_
                lhs39_ = d_347_globalState_
                lhs38_.f3 = rhs40_
                d_384_v37_ = rhs41_
                d_378_v31_ = rhs42_
                d_401_v49_ = rhs43_
                lhs39_.f9 = rhs44_
                d_413_v54_: _dafny.Array
                nw57_ = _dafny.Array(None, 7)
                d_413_v54_ = nw57_
                d_414_v55_: D14
                d_414_v55_ = D14_DC44(d_344_v7_)
                index43_ = default__.safeIndex(715, (d_413_v54_).length(0))
                (d_413_v54_)[index43_] = d_414_v55_
                pat_let_tv17_ = d_344_v7_
                index44_ = default__.safeIndex(715, (d_413_v54_).length(0))
                def iife65_(_pat_let15_0):
                    def iife66_(d_415_dt__update__tmp_h5_):
                        def iife67_(_pat_let16_0):
                            def iife68_(d_416_dt__update_hcf71_h0_):
                                return D14_DC44(d_416_dt__update_hcf71_h0_)
                            return iife68_(_pat_let16_0)
                        return iife67_(pat_let_tv17_)
                    return iife66_(_pat_let15_0)
                (d_413_v54_)[index44_] = iife65_(d_414_v55_)
                d_417_v56_: _dafny.Map
                d_417_v56_ = _dafny.Map({(d_386_v39_)[default__.safeIndex(190, (d_386_v39_).length(0))]: (d_393_v44_).set(d_341_v4_, default__.abs(d_379_v32_))})
                d_418_v57_: C5
                nw58_ = C5()
                nw58_.ctor__(d_384_v37_, d_341_v4_, (((d_417_v56_)[len(_dafny.SeqWithoutIsStrInference([d_378_v31_ for d_419_i6_ in range(default__.abs(387))]))] if (len(_dafny.SeqWithoutIsStrInference([d_378_v31_ for d_420_i6_ in range(default__.abs(387))]))) in (d_417_v56_) else d_393_v44_)) | (_dafny.MultiSet(d_348_v8_)))
                d_418_v57_ = nw58_
                d_378_v31_ = d_378_v31_
                d_421_v58_: _dafny.Array
                def lambda9_(d_422_v4_, d_423_v3_, d_424_v8_):
                    def lambda10_(d_425_i7_):
                        return ((_dafny.SeqWithoutIsStrInference([d_422_v4_, d_422_v4_])).set(default__.safeIndex(len(d_423_v3_), len(_dafny.SeqWithoutIsStrInference([d_422_v4_, d_422_v4_]))), d_422_v4_)) + (d_424_v8_)

                    return lambda10_

                init4_ = lambda9_(d_341_v4_, d_340_v3_, d_348_v8_)
                nw59_ = _dafny.Array(None, 28)
                for i0_4_ in range(nw59_.length(0)):
                    nw59_[i0_4_] = init4_(i0_4_)
                d_421_v58_ = nw59_
                index45_ = default__.safeIndex(201, (d_421_v58_).length(0))
                (d_421_v58_)[index45_] = d_348_v8_
                index46_ = default__.safeIndex(201, (d_421_v58_).length(0))
                (d_421_v58_)[index46_] = d_348_v8_
            elif True:
                (d_347_globalState_).f9 = d_341_v4_
                d_386_v39_ = d_386_v39_
                (d_347_globalState_).f3 = default__.fm0(not(d_341_v4_), d_384_v37_, d_337_v0_, (d_338_v1_).intersection(d_338_v1_), d_347_globalState_)
                d_426_v59_: _dafny.Array
                nw60_ = _dafny.Array(_dafny.Seq({}), 19)
                d_426_v59_ = nw60_
                index47_ = default__.safeIndex(307, (d_426_v59_).length(0))
                (d_426_v59_)[index47_] = d_348_v8_
                d_427_v60_: _dafny.Map
                d_427_v60_ = _dafny.Map({(d_348_v8_) + (_dafny.SeqWithoutIsStrInference([False])): ((0) - (len(default__.fm53(d_341_v4_, d_341_v4_, d_347_globalState_))) if d_341_v4_ else (d_386_v39_)[default__.safeIndex(190, (d_386_v39_).length(0))])})
                d_428_v61_: C4
                nw61_ = C4()
                nw61_.ctor__(d_341_v4_, d_341_v4_, (_dafny.MultiSet([d_341_v4_, d_341_v4_])).set(d_341_v4_, default__.abs((d_338_v1_).cardinality)))
                d_428_v61_ = nw61_
                d_429_v62_: _dafny.Map
                d_429_v62_ = _dafny.Map({d_428_v61_: d_379_v32_})
                index48_ = default__.safeIndex(307, (d_426_v59_).length(0))
                rhs45_ = default__.fm38(d_347_globalState_)
                rhs46_ = ((_dafny.Map({d_348_v8_: len(d_429_v62_)})).set(d_348_v8_, 267)) | (_dafny.Map({(d_348_v8_).set(default__.safeIndex(d_337_v0_, len(d_348_v8_)), (d_428_v61_).f19): d_379_v32_}))
                lhs40_ = d_426_v59_
                lhs41_ = default__.safeIndex(307, (d_426_v59_).length(0))
                lhs40_[lhs41_] = rhs45_
                d_427_v60_ = rhs46_
                index49_ = default__.safeIndex(190, (d_386_v39_).length(0))
                (d_386_v39_)[index49_] = -956
            d_430_v64_: _dafny.Map
            d_430_v64_ = _dafny.Map({(d_350_v10_)[default__.safeIndex((d_386_v39_)[default__.safeIndex(190, (d_386_v39_).length(0))], len(d_350_v10_))]: _dafny.Set({d_341_v4_})})
            d_431_v65_: D0
            def iife69_():
                coll35_ = _dafny.Set()
                compr_35_: int
                for compr_35_ in ((d_350_v10_).set(default__.safeIndex(len(d_340_v3_), len(d_350_v10_)), (d_338_v1_).cardinality)).Elements:
                    d_432_v63_: int = compr_35_
                    if (d_432_v63_) in ((d_350_v10_).set(default__.safeIndex(len(d_340_v3_), len(d_350_v10_)), (d_338_v1_).cardinality)):
                        coll35_ = coll35_.union(_dafny.Set([default__.safeDivisionInt(d_432_v63_, 240)]))
                return _dafny.Set(coll35_)
            d_431_v65_ = D0_DC2(default__.fm37(iife69_()
, d_430_v64_, d_341_v4_, d_341_v4_, d_347_globalState_))
            source9_ = d_431_v65_
            if source9_.is_DC1:
                d_433___mcc_h0_ = source9_.cf1
                d_434_cf1_ = d_433___mcc_h0_
                (d_347_globalState_).f3 = d_384_v37_
                d_379_v32_ = (d_337_v0_) + (default__.fm0(not(False), d_384_v37_, len(d_342_v5_), d_338_v1_, d_347_globalState_))
                d_435_v66_: int
                d_436_v67_: _dafny.Map
                out7_: int
                out8_: _dafny.Map
                out7_, out8_ = default__.m11(d_376_v29_, d_384_v37_, (633) < ((0) - (d_379_v32_)), False, d_347_globalState_)
                d_435_v66_ = out7_
                d_436_v67_ = out8_
                d_435_v66_ = d_379_v32_
            elif source9_.is_DC0:
                d_437___mcc_h1_ = source9_.cf0
                d_438_cf0_ = d_437___mcc_h1_
                d_341_v4_ = d_438_cf0_
                d_384_v37_ = d_379_v32_
                d_439_v68_: _dafny.MultiSet
                d_439_v68_ = _dafny.MultiSet([d_341_v4_, d_438_cf0_])
                d_440_v69_: _dafny.Set
                d_440_v69_ = _dafny.Set({d_379_v32_, (0) - (d_337_v0_), ((d_439_v68_)[False] if (False) in (d_439_v68_) else d_384_v37_)})
                (d_347_globalState_).f9 = (d_440_v69_).ispropersubset(_dafny.Set({981, d_337_v0_}))
                d_441_v70_: _dafny.Map
                d_441_v70_ = _dafny.Map({d_379_v32_: _dafny.SeqWithoutIsStrInference([d_341_v4_, d_438_cf0_])})
                (d_347_globalState_).f9 = ((((d_441_v70_)[(d_386_v39_)[default__.safeIndex(190, (d_386_v39_).length(0))]] if ((d_386_v39_)[default__.safeIndex(190, (d_386_v39_).length(0))]) in (d_441_v70_) else d_348_v8_)) + (d_348_v8_)) < (d_348_v8_)
            elif True:
                d_442___mcc_h2_ = source9_.cf2
                d_443_cf2_ = d_442___mcc_h2_
                d_444_v71_: D13
                d_444_v71_ = D13_DC42(d_378_v31_, not(True))
                d_445_v72_: int
                d_446_v73_: _dafny.Map
                out9_: int
                out10_: _dafny.Map
                out9_, out10_ = default__.m11(D13_DC43(d_444_v71_), d_384_v37_, d_341_v4_, d_341_v4_, d_347_globalState_)
                d_445_v72_ = out9_
                d_446_v73_ = out10_
                d_341_v4_ = (d_338_v1_) == (d_338_v1_)
                d_447_v74_: D5
                d_447_v74_ = D5_DC16(d_386_v39_)
                d_386_v39_ = (d_447_v74_).cf29
                (d_347_globalState_).f9 = d_341_v4_
            d_337_v0_ = d_337_v0_
        elif True:
            (d_347_globalState_).f9 = ((not(d_341_v4_) if d_341_v4_ else d_341_v4_)) in ((d_348_v8_) + (d_348_v8_))
            d_448_v75_: _dafny.Map
            d_448_v75_ = _dafny.Map({d_378_v31_: d_337_v0_})
            d_449_v77_: int
            d_450_v78_: _dafny.Map
            out11_: int
            out12_: _dafny.Map
            def iife70_():
                coll36_ = _dafny.Set()
                compr_36_: _dafny.Seq
                for compr_36_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_337_v0_]): d_337_v0_})).keys.Elements:
                    d_451_v76_: _dafny.Seq = compr_36_
                    if (d_451_v76_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_337_v0_]): d_337_v0_})):
                        coll36_ = coll36_.union(_dafny.Set([d_451_v76_]))
                return _dafny.Set(coll36_)
            out11_, out12_ = default__.m11(d_376_v29_, d_337_v0_, d_341_v4_, default__.fm23(d_337_v0_, d_341_v4_, d_448_v75_, iife70_()
            , d_347_globalState_), d_347_globalState_)
            d_449_v77_ = out11_
            d_450_v78_ = out12_
            d_452_v79_: C4
            nw62_ = C4()
            nw62_.ctor__(d_341_v4_, d_341_v4_, default__.fm33(d_341_v4_, d_347_globalState_))
            d_452_v79_ = nw62_
            d_453_v80_: _dafny.MultiSet
            d_453_v80_ = _dafny.MultiSet([d_452_v79_])
            d_454_v81_: _dafny.MultiSet
            d_454_v81_ = _dafny.MultiSet([d_341_v4_, d_341_v4_])
            d_455_v82_: _dafny.Map
            d_455_v82_ = _dafny.Map({d_453_v80_: ((d_454_v81_)[(d_452_v79_).f19] if ((d_452_v79_).f19) in (d_454_v81_) else d_379_v32_)})
            d_455_v82_ = (d_455_v82_).set(d_453_v80_, default__.safeModuloInt(d_379_v32_, d_449_v77_))
            if (True if d_341_v4_ else ((d_452_v79_).f19 if (d_452_v79_).fm2((d_452_v79_).f19, d_341_v4_, d_347_globalState_) else d_341_v4_)):
                d_456_v83_: _dafny.Array
                def lambda11_(d_457_v1_):
                    def lambda12_(d_458_i8_):
                        return (d_458_i8_) * ((d_457_v1_).cardinality)

                    return lambda12_

                init5_ = lambda11_(d_338_v1_)
                nw63_ = _dafny.Array(None, 12)
                for i0_5_ in range(nw63_.length(0)):
                    nw63_[i0_5_] = init5_(i0_5_)
                d_456_v83_ = nw63_
                index50_ = default__.safeIndex(845, (d_456_v83_).length(0))
                (d_456_v83_)[index50_] = d_449_v77_
                d_459_v84_: D4
                d_459_v84_ = D4_DC12(d_339_v2_)
                d_460_v85_: D4
                d_460_v85_ = D4_DC15(d_459_v84_)
                d_461_v86_: C6
                nw64_ = C6()
                nw64_.ctor__(d_341_v4_, d_454_v81_)
                d_461_v86_ = nw64_
                d_462_v87_: _dafny.Seq
                d_462_v87_ = _dafny.SeqWithoutIsStrInference([d_461_v86_])
                d_463_v88_: _dafny.Map
                d_463_v88_ = _dafny.Map({d_462_v87_: default__.fm6(d_347_globalState_)})
                pat_let_tv18_ = d_459_v84_
                index51_ = default__.safeIndex(845, (d_456_v83_).length(0))
                def iife71_(_pat_let17_0):
                    def iife72_(d_464_dt__update__tmp_h6_):
                        def iife73_(_pat_let18_0):
                            def iife74_(d_465_dt__update_hcf28_h0_):
                                return D4_DC15(d_465_dt__update_hcf28_h0_)
                            return iife74_(_pat_let18_0)
                        return iife73_(pat_let_tv18_)
                    return iife72_(_pat_let17_0)
                rhs47_ = (default__.safeDivisionInt(d_337_v0_, 632)) * (len(d_463_v88_))
                rhs48_ = d_379_v32_
                rhs49_ = d_379_v32_
                rhs50_ = iife71_(d_460_v85_)
                lhs42_ = d_456_v83_
                lhs43_ = default__.safeIndex(845, (d_456_v83_).length(0))
                lhs44_ = d_347_globalState_
                lhs45_ = d_347_globalState_
                lhs42_[lhs43_] = rhs47_
                lhs44_.f3 = rhs48_
                lhs45_.f3 = rhs49_
                d_460_v85_ = rhs50_
                index52_ = default__.safeIndex(845, (d_456_v83_).length(0))
                (d_456_v83_)[index52_] = (d_456_v83_)[default__.safeIndex(845, (d_456_v83_).length(0))]
                d_466_v89_: int
                d_467_v90_: int
                d_468_v91_: _dafny.MultiSet
                d_469_v92_: bool
                out13_: int
                out14_: int
                out15_: _dafny.MultiSet
                out16_: bool
                out13_, out14_, out15_, out16_ = (d_461_v86_).m5(d_347_globalState_)
                d_466_v89_ = out13_
                d_467_v90_ = out14_
                d_468_v91_ = out15_
                d_469_v92_ = out16_
                d_470_v93_: int
                d_471_v94_: bool
                d_472_v95_: _dafny.Map
                d_473_v96_: bool
                out17_: int
                out18_: bool
                out19_: _dafny.Map
                out20_: bool
                out17_, out18_, out19_, out20_ = (d_452_v79_).m0(d_347_globalState_)
                d_470_v93_ = out17_
                d_471_v94_ = out18_
                d_472_v95_ = out19_
                d_473_v96_ = out20_
                d_474_v97_: C7
                nw65_ = C7()
                nw65_.ctor__()
                d_474_v97_ = nw65_
            elif True:
                d_475_v98_: _dafny.Array
                def lambda13_(d_476_v31_):
                    def lambda14_(d_477_i9_):
                        return d_476_v31_

                    return lambda14_

                init6_ = lambda13_(d_378_v31_)
                nw66_ = _dafny.Array(None, 6)
                for i0_6_ in range(nw66_.length(0)):
                    nw66_[i0_6_] = init6_(i0_6_)
                d_475_v98_ = nw66_
                d_478_v99_: _dafny.Map
                d_478_v99_ = _dafny.Map({_dafny.CodePoint('f'): d_475_v98_})
                d_479_v100_: _dafny.Map
                d_479_v100_ = _dafny.Map({(d_452_v79_).f19: d_379_v32_})
                d_480_v101_: _dafny.Array
                nw67_ = _dafny.Array(None, 15)
                nw67_[int(0)] = d_475_v98_
                nw67_[int(1)] = d_475_v98_
                nw67_[int(2)] = d_475_v98_
                nw67_[int(3)] = d_475_v98_
                nw67_[int(4)] = ((d_478_v99_)[default__.fm25(d_337_v0_, len(d_350_v10_), (d_452_v79_).f19, d_479_v100_, d_347_globalState_)] if (default__.fm25(d_337_v0_, len(d_350_v10_), (d_452_v79_).f19, d_479_v100_, d_347_globalState_)) in (d_478_v99_) else d_475_v98_)
                nw67_[int(5)] = d_475_v98_
                nw67_[int(6)] = d_475_v98_
                nw67_[int(7)] = d_475_v98_
                nw67_[int(8)] = d_475_v98_
                nw67_[int(9)] = d_475_v98_
                nw67_[int(10)] = d_475_v98_
                nw67_[int(11)] = (d_475_v98_ if (d_452_v79_).f19 else d_475_v98_)
                nw67_[int(12)] = d_475_v98_
                nw67_[int(13)] = d_475_v98_
                nw67_[int(14)] = d_475_v98_
                d_480_v101_ = nw67_
                index53_ = default__.safeIndex(670, (d_480_v101_).length(0))
                (d_480_v101_)[index53_] = d_475_v98_
                d_481_v102_: D16
                d_481_v102_ = D16_DC51(d_475_v98_)
                index54_ = default__.safeIndex(670, (d_480_v101_).length(0))
                (d_480_v101_)[index54_] = (d_481_v102_).cf83
                d_482_v103_: _dafny.Array
                nw68_ = _dafny.Array(int(0), 7)
                d_482_v103_ = nw68_
                index55_ = default__.safeIndex(207, (d_482_v103_).length(0))
                (d_482_v103_)[index55_] = (d_449_v77_ if (d_452_v79_).f19 else d_337_v0_)
                d_483_v104_: _dafny.Map
                d_483_v104_ = _dafny.Map({d_482_v103_: ((d_348_v8_)[default__.safeIndex(len(d_377_v30_), len(d_348_v8_))] if d_341_v4_ else d_341_v4_)})
                d_484_v105_: D5
                d_484_v105_ = D5_DC18((d_452_v79_).f19)
                d_485_v106_: _dafny.Map
                d_485_v106_ = _dafny.Map({d_337_v0_: D11_DC35()})
                d_486_v107_: _dafny.Set
                d_486_v107_ = _dafny.Set({d_340_v3_})
                index56_ = default__.safeIndex(207, (d_482_v103_).length(0))
                rhs51_ = (d_452_v79_).f19
                rhs52_ = (d_484_v105_).cf33
                rhs53_ = ((d_452_v79_).f19 if (d_485_v106_) != (d_485_v106_) else (d_486_v107_).isdisjoint(d_486_v107_))
                rhs54_ = default__.safeModuloInt(d_337_v0_, d_337_v0_)
                rhs55_ = d_483_v104_
                lhs46_ = d_347_globalState_
                lhs47_ = d_347_globalState_
                lhs48_ = d_482_v103_
                lhs49_ = default__.safeIndex(207, (d_482_v103_).length(0))
                lhs46_.f9 = rhs51_
                lhs47_.f9 = rhs52_
                d_341_v4_ = rhs53_
                lhs48_[lhs49_] = rhs54_
                d_483_v104_ = rhs55_
                d_487_v108_: C3
                nw69_ = C3()
                nw69_.ctor__((d_482_v103_)[default__.safeIndex(207, (d_482_v103_).length(0))], d_340_v3_, d_341_v4_, d_454_v81_)
                d_487_v108_ = nw69_
                d_488_v109_: D18
                d_488_v109_ = D18_DC55(d_487_v108_)
                d_489_v110_: _dafny.MultiSet
                d_489_v110_ = _dafny.MultiSet([(d_488_v109_).cf91, d_487_v108_])
                d_490_v111_: _dafny.Seq
                d_490_v111_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({(d_489_v110_).cardinality: (d_487_v108_).f20})).set(len((d_348_v8_).set(default__.safeIndex((d_482_v103_)[default__.safeIndex(207, (d_482_v103_).length(0))], len(d_348_v8_)), (d_452_v79_).f19)), (0) - ((d_487_v108_).f20))])
                d_490_v111_ = d_490_v111_
                index57_ = default__.safeIndex(919, (d_339_v2_).length(0))
                (d_339_v2_)[index57_] = True
                index58_ = default__.safeIndex(919, (d_339_v2_).length(0))
                (d_339_v2_)[index58_] = (d_452_v79_).f19
                d_491_v112_: _dafny.Map
                d_491_v112_ = _dafny.Map({d_341_v4_: d_378_v31_})
                d_492_v113_: _dafny.MultiSet
                d_492_v113_ = _dafny.MultiSet([(d_491_v112_).set((d_452_v79_).f19, d_378_v31_), d_491_v112_])
                (d_347_globalState_).f9 = (d_492_v113_).ispropersubset(d_492_v113_)
            d_493_v114_: _dafny.Map
            d_493_v114_ = _dafny.Map({d_339_v2_: d_378_v31_})
            d_494_v115_: D15
            d_494_v115_ = D15_DC49((d_493_v114_) | (d_493_v114_), -88)
            d_494_v115_ = D15_DC49(d_493_v114_, d_379_v32_)
        d_337_v0_ = d_379_v32_
        d_341_v4_ = d_341_v4_
        d_495_v116_: _dafny.MultiSet
        d_495_v116_ = _dafny.MultiSet([d_341_v4_, d_341_v4_, d_341_v4_, d_341_v4_, d_341_v4_])
        d_496_v117_: C6
        nw70_ = C6()
        nw70_.ctor__(d_341_v4_, d_495_v116_)
        d_496_v117_ = nw70_
        if (d_350_v10_) != (d_350_v10_):
            d_337_v0_ = default__.fm0(d_341_v4_, ((0) - (d_379_v32_)) + (359), d_379_v32_, d_338_v1_, d_347_globalState_)
            d_497_v118_: D8
            d_497_v118_ = D8_DC24(_dafny.Set({d_341_v4_}))
            d_498_v119_: _dafny.Map
            d_498_v119_ = _dafny.Map({d_379_v32_: d_342_v5_})
            d_499_v120_: _dafny.Map
            d_499_v120_ = _dafny.Map({d_337_v0_: d_341_v4_})
            d_500_v121_: _dafny.Seq
            d_500_v121_ = _dafny.SeqWithoutIsStrInference([d_342_v5_, _dafny.Set({d_341_v4_, d_341_v4_, True}), _dafny.Set({d_341_v4_})])
            d_501_v122_: _dafny.Array
            nw71_ = _dafny.Array(None, 16)
            nw71_[int(0)] = d_342_v5_
            nw71_[int(1)] = d_342_v5_
            nw71_[int(2)] = d_342_v5_
            nw71_[int(3)] = (d_497_v118_).cf45
            nw71_[int(4)] = d_342_v5_
            nw71_[int(5)] = ((d_498_v119_)[d_337_v0_] if (d_337_v0_) in (d_498_v119_) else d_342_v5_)
            nw71_[int(6)] = (d_342_v5_) | (default__.fm41(d_379_v32_, len(d_499_v120_), d_347_globalState_))
            nw71_[int(7)] = ((d_500_v121_)[default__.safeIndex(d_379_v32_, len(d_500_v121_))]).intersection(_dafny.Set({d_341_v4_}))
            nw71_[int(8)] = (d_342_v5_ if d_341_v4_ else d_342_v5_)
            nw71_[int(9)] = default__.fm41(len(d_350_v10_), d_337_v0_, d_347_globalState_)
            nw71_[int(10)] = _dafny.Set({d_341_v4_, d_341_v4_})
            nw71_[int(11)] = d_342_v5_
            nw71_[int(12)] = (d_342_v5_) - (_dafny.Set({d_341_v4_, d_341_v4_, d_341_v4_, d_341_v4_, d_341_v4_}))
            nw71_[int(13)] = (d_342_v5_) | (d_342_v5_)
            nw71_[int(14)] = d_342_v5_
            nw71_[int(15)] = d_342_v5_
            d_501_v122_ = nw71_
            index59_ = default__.safeIndex(78, (d_501_v122_).length(0))
            (d_501_v122_)[index59_] = (d_342_v5_).intersection(d_342_v5_)
            d_502_v123_: _dafny.Array
            nw72_ = _dafny.Array(None, 10)
            nw72_[int(0)] = (0) - (d_337_v0_)
            nw72_[int(1)] = (d_350_v10_)[default__.safeIndex(len(d_340_v3_), len(d_350_v10_))]
            nw72_[int(2)] = len(d_348_v8_)
            nw72_[int(3)] = default__.safeDivisionInt(d_379_v32_, d_337_v0_)
            nw72_[int(4)] = d_337_v0_
            nw72_[int(5)] = (0) - ((d_337_v0_) + (d_337_v0_))
            nw72_[int(6)] = d_379_v32_
            nw72_[int(7)] = len((_dafny.Map({d_341_v4_: d_378_v31_})).set(d_341_v4_, d_378_v31_))
            nw72_[int(8)] = (d_379_v32_) - (d_379_v32_)
            nw72_[int(9)] = d_337_v0_
            d_502_v123_ = nw72_
            index60_ = default__.safeIndex(170, (d_502_v123_).length(0))
            (d_502_v123_)[index60_] = d_337_v0_
            d_503_v124_: D12
            d_503_v124_ = D12_DC39(d_379_v32_)
            index61_ = default__.safeIndex(78, (d_501_v122_).length(0))
            index62_ = default__.safeIndex(170, (d_502_v123_).length(0))
            rhs56_ = default__.fm41(len(d_340_v3_), (d_503_v124_).cf65, d_347_globalState_)
            rhs57_ = (d_342_v5_).intersection(d_342_v5_)
            rhs58_ = (0) - ((d_337_v0_) + (len(d_340_v3_)))
            rhs59_ = default__.fm0(d_341_v4_, d_337_v0_, -571, d_338_v1_, d_347_globalState_)
            lhs50_ = d_501_v122_
            lhs51_ = default__.safeIndex(78, (d_501_v122_).length(0))
            lhs52_ = d_347_globalState_
            lhs53_ = d_502_v123_
            lhs54_ = default__.safeIndex(170, (d_502_v123_).length(0))
            lhs50_[lhs51_] = rhs56_
            lhs52_.f6 = rhs57_
            lhs53_[lhs54_] = rhs58_
            d_379_v32_ = rhs59_
            d_341_v4_ = not(d_341_v4_)
            d_341_v4_ = True
            index63_ = default__.safeIndex(170, (d_502_v123_).length(0))
            (d_502_v123_)[index63_] = d_337_v0_
        elif True:
            if d_341_v4_:
                d_495_v116_ = d_495_v116_
                d_504_v125_: _dafny.Set
                d_504_v125_ = _dafny.Set({78, d_337_v0_})
                d_505_v126_: _dafny.Map
                d_505_v126_ = _dafny.Map({d_378_v31_: d_504_v125_})
                d_506_v127_: int
                d_507_v128_: int
                d_508_v129_: bool
                out21_: int
                out22_: int
                out23_: bool
                out21_, out22_, out23_ = (d_496_v117_).m1((((d_505_v126_)[d_378_v31_] if (d_378_v31_) in (d_505_v126_) else d_504_v125_)) | (_dafny.Set({-222})), 30, d_347_globalState_)
                d_506_v127_ = out21_
                d_507_v128_ = out22_
                d_508_v129_ = out23_
                d_509_v130_: C3
                nw73_ = C3()
                nw73_.ctor__(d_507_v128_, d_340_v3_, True, d_495_v116_)
                d_509_v130_ = nw73_
                d_510_v131_: _dafny.Map
                d_510_v131_ = _dafny.Map({default__.safeDivisionInt(d_507_v128_, d_507_v128_): d_509_v130_})
                d_510_v131_ = (d_510_v131_).set(d_379_v32_, d_509_v130_)
                d_511_v132_: _dafny.Map
                d_511_v132_ = _dafny.Map({d_507_v128_: d_339_v2_})
                d_512_v133_: C2
                nw74_ = C2()
                nw74_.ctor__((_dafny.Map({len(d_511_v132_): d_341_v4_})).set(d_379_v32_, d_341_v4_))
                d_512_v133_ = nw74_
                d_513_v134_: _dafny.Map
                d_513_v134_ = _dafny.Map({d_508_v129_: d_337_v0_})
                d_513_v134_ = (d_513_v134_).set(d_508_v129_, d_337_v0_)
            elif True:
                d_514_v135_: int
                d_515_v136_: int
                d_516_v137_: _dafny.MultiSet
                d_517_v138_: bool
                out24_: int
                out25_: int
                out26_: _dafny.MultiSet
                out27_: bool
                out24_, out25_, out26_, out27_ = (d_496_v117_).m5(d_347_globalState_)
                d_514_v135_ = out24_
                d_515_v136_ = out25_
                d_516_v137_ = out26_
                d_517_v138_ = out27_
                d_518_v139_: C7
                nw75_ = C7()
                nw75_.ctor__()
                d_518_v139_ = nw75_
                nw76_ = C7()
                nw76_.ctor__()
                d_518_v139_ = nw76_
                d_519_v140_: _dafny.Map
                d_519_v140_ = _dafny.Map({len(_dafny.Map({d_341_v4_: d_341_v4_})): d_517_v138_})
                d_520_v141_: _dafny.Array
                nw77_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_520_v141_ = nw77_
                d_521_v142_: D16
                d_521_v142_ = D16_DC52(d_337_v0_, d_517_v138_, (D16_DC52(d_515_v136_, d_517_v138_, d_520_v141_, True, d_337_v0_)).cf86, d_341_v4_, d_515_v136_)
                d_519_v140_ = (d_519_v140_).set(((d_521_v142_).cf88) * (d_379_v32_), d_341_v4_)
                d_522_v143_: _dafny.Array
                def lambda15_(d_523_v136_):
                    def lambda16_(d_524_i10_):
                        return default__.safeDivisionInt(d_524_i10_, d_523_v136_)

                    return lambda16_

                init7_ = lambda15_(d_515_v136_)
                nw78_ = _dafny.Array(None, 11)
                for i0_7_ in range(nw78_.length(0)):
                    nw78_[i0_7_] = init7_(i0_7_)
                d_522_v143_ = nw78_
                index64_ = default__.safeIndex(825, (d_522_v143_).length(0))
                (d_522_v143_)[index64_] = (d_350_v10_)[default__.safeIndex(d_337_v0_, len(d_350_v10_))]
                d_525_v144_: _dafny.Map
                d_525_v144_ = _dafny.Map({d_517_v138_: d_379_v32_})
                d_526_v145_: _dafny.Set
                d_526_v145_ = _dafny.Set({d_379_v32_})
                d_527_v147_: T0
                nw79_ = C5()
                nw79_.ctor__(d_515_v136_, d_517_v138_, d_495_v116_)
                d_527_v147_ = nw79_
                d_528_v148_: D19
                d_528_v148_ = D19_DC59(d_527_v147_, d_525_v144_)
                d_529_v149_: _dafny.Array
                nw80_ = _dafny.Array(None, 29)
                nw80_[int(0)] = (d_525_v144_).set(d_517_v138_, d_514_v135_)
                nw80_[int(1)] = d_525_v144_
                nw80_[int(2)] = (d_525_v144_) | (d_525_v144_)
                nw80_[int(3)] = default__.fm1(d_517_v138_, d_379_v32_, d_341_v4_, d_526_v145_, d_347_globalState_)
                nw80_[int(4)] = d_525_v144_
                nw80_[int(5)] = _dafny.Map({d_341_v4_: d_379_v32_})
                nw80_[int(6)] = d_525_v144_
                def iife75_():
                    coll37_ = _dafny.Set()
                    compr_37_: int
                    for compr_37_ in _dafny.IntegerRange(994, 775):
                        d_530_v146_: int = compr_37_
                        if ((994) <= (d_530_v146_)) and ((d_530_v146_) < (775)):
                            coll37_ = coll37_.union(_dafny.Set([default__.safeDivisionInt(d_530_v146_, d_379_v32_)]))
                    return _dafny.Set(coll37_)
                nw80_[int(7)] = (default__.fm1(d_341_v4_, d_379_v32_, not(d_517_v138_), iife75_()
                , d_347_globalState_) if not(d_341_v4_) else d_525_v144_)
                nw80_[int(8)] = d_525_v144_
                nw80_[int(9)] = d_525_v144_
                nw80_[int(10)] = default__.fm1(d_341_v4_, 705, d_341_v4_, d_526_v145_, d_347_globalState_)
                nw80_[int(11)] = (d_528_v148_).cf98
                nw80_[int(12)] = d_525_v144_
                nw80_[int(13)] = (_dafny.Map({d_517_v138_: d_514_v135_})) | (d_525_v144_)
                nw80_[int(14)] = d_525_v144_
                nw80_[int(15)] = d_525_v144_
                nw80_[int(16)] = (d_525_v144_) | (_dafny.Map({d_517_v138_: d_379_v32_}))
                nw80_[int(17)] = (default__.fm1(not(True), d_515_v136_, d_517_v138_, d_526_v145_, d_347_globalState_)) | (d_525_v144_)
                nw80_[int(18)] = (d_525_v144_) | ((d_525_v144_).set(d_341_v4_, d_379_v32_))
                nw80_[int(19)] = (D19_DC59(d_527_v147_, d_525_v144_)).cf98
                nw80_[int(20)] = d_525_v144_
                nw80_[int(21)] = (d_525_v144_) | (d_525_v144_)
                nw80_[int(22)] = d_525_v144_
                nw80_[int(23)] = d_525_v144_
                nw80_[int(24)] = d_525_v144_
                nw80_[int(25)] = d_525_v144_
                nw80_[int(26)] = _dafny.Map({False: d_337_v0_})
                nw80_[int(27)] = (d_525_v144_) | (default__.fm1(d_517_v138_, d_514_v135_, d_517_v138_, d_526_v145_, d_347_globalState_))
                nw80_[int(28)] = _dafny.Map({d_527_v147_.f14: d_337_v0_})
                d_529_v149_ = nw80_
                d_531_v150_: _dafny.Seq
                d_531_v150_ = _dafny.SeqWithoutIsStrInference([d_525_v144_])
                index65_ = default__.safeIndex(514, (d_529_v149_).length(0))
                (d_529_v149_)[index65_] = ((d_531_v150_)[default__.safeIndex(d_514_v135_, len(d_531_v150_))]) | (_dafny.Map({d_527_v147_.f14: d_515_v136_}))
                index66_ = default__.safeIndex(825, (d_522_v143_).length(0))
                index67_ = default__.safeIndex(514, (d_529_v149_).length(0))
                rhs60_ = d_378_v31_
                rhs61_ = d_379_v32_
                rhs62_ = ((d_525_v144_) | (d_525_v144_)) | (((d_525_v144_).set(d_341_v4_, d_514_v135_)) | (d_525_v144_))
                rhs63_ = d_337_v0_
                rhs64_ = d_378_v31_
                lhs55_ = d_522_v143_
                lhs56_ = default__.safeIndex(825, (d_522_v143_).length(0))
                lhs57_ = d_529_v149_
                lhs58_ = default__.safeIndex(514, (d_529_v149_).length(0))
                d_378_v31_ = rhs60_
                lhs55_[lhs56_] = rhs61_
                lhs57_[lhs58_] = rhs62_
                d_515_v136_ = rhs63_
                d_378_v31_ = rhs64_
                d_337_v0_ = (d_515_v136_) - (default__.safeDivisionInt(d_514_v135_, len(d_340_v3_)))
            d_532_v151_: _dafny.Array
            nw81_ = _dafny.Array(int(0), 25)
            d_532_v151_ = nw81_
            d_532_v151_ = d_532_v151_
            index68_ = default__.safeIndex(349, (d_532_v151_).length(0))
            (d_532_v151_)[index68_] = 815
            index69_ = default__.safeIndex(349, (d_532_v151_).length(0))
            (d_532_v151_)[index69_] = (d_379_v32_ if True else d_379_v32_)
            d_340_v3_ = d_340_v3_
            index70_ = default__.safeIndex(349, (d_532_v151_).length(0))
            (d_532_v151_)[index70_] = (d_532_v151_)[default__.safeIndex(349, (d_532_v151_).length(0))]
        d_533_i11_: int
        d_533_i11_ = 0
        with _dafny.label("0"):
            while d_341_v4_:
                with _dafny.c_label("0"):
                    if (d_533_i11_) >= (100):
                        raise _dafny.Break("0")
                    d_533_i11_ = (d_533_i11_) + (1)
                    d_534_v152_: _dafny.Array
                    def lambda17_(d_535_v4_):
                        def lambda18_(d_536_i12_):
                            return (d_536_i12_) * ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D3_DC10(_dafny.Map({d_535_v4_: d_535_v4_}))]))).cardinality)

                        return lambda18_

                    init8_ = lambda17_(d_341_v4_)
                    nw82_ = _dafny.Array(None, 25)
                    for i0_8_ in range(nw82_.length(0)):
                        nw82_[i0_8_] = init8_(i0_8_)
                    d_534_v152_ = nw82_
                    index71_ = default__.safeIndex(726, (d_534_v152_).length(0))
                    (d_534_v152_)[index71_] = d_379_v32_
                    index72_ = default__.safeIndex(726, (d_534_v152_).length(0))
                    (d_534_v152_)[index72_] = default__.safeModuloInt((d_338_v1_).cardinality, d_379_v32_)
                    d_537_v153_: _dafny.Array
                    nw83_ = _dafny.Array(None, 15)
                    d_537_v153_ = nw83_
                    d_538_v154_: C3
                    nw84_ = C3()
                    nw84_.ctor__(d_337_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihvjxhpxy")), d_341_v4_, d_495_v116_)
                    d_538_v154_ = nw84_
                    index73_ = default__.safeIndex(945, (d_537_v153_).length(0))
                    (d_537_v153_)[index73_] = d_538_v154_
                    index74_ = default__.safeIndex(945, (d_537_v153_).length(0))
                    nw85_ = C3()
                    nw85_.ctor__(len(_dafny.SeqWithoutIsStrInference([True, d_341_v4_, d_341_v4_, d_341_v4_])), ((d_538_v154_).f21) + ((d_538_v154_).f21), d_341_v4_, d_495_v116_)
                    (d_537_v153_)[index74_] = nw85_
                    d_539_v155_: D2
                    d_539_v155_ = D2_DC7(d_348_v8_)
                    d_540_v156_: _dafny.Set
                    d_540_v156_ = _dafny.Set({d_539_v155_, d_539_v155_, d_539_v155_})
                    (d_347_globalState_).f9 = (D2_DC7(d_348_v8_)) not in (d_540_v156_)
                    index75_ = default__.safeIndex(918, (d_339_v2_).length(0))
                    (d_339_v2_)[index75_] = d_341_v4_
                    d_541_v157_: _dafny.Map
                    d_541_v157_ = _dafny.Map({d_341_v4_: ((d_538_v154_).f21)[default__.safeIndex(-660, len((d_538_v154_).f21))]})
                    d_542_v158_: _dafny.Set
                    d_542_v158_ = _dafny.Set({(0) - (len(d_541_v157_))})
                    index76_ = default__.safeIndex(918, (d_339_v2_).length(0))
                    (d_339_v2_)[index76_] = (default__.fm6(d_347_globalState_)).issubset(d_542_v158_)
                    pass
            pass
        d_543_v159_: _dafny.Array
        def lambda19_(d_544_v8_):
            def lambda20_(d_545_i14_):
                return (d_545_i14_) - (len(d_544_v8_))

            return lambda20_

        init9_ = lambda19_(d_348_v8_)
        nw86_ = _dafny.Array(None, 13)
        for i0_9_ in range(nw86_.length(0)):
            nw86_[i0_9_] = init9_(i0_9_)
        d_543_v159_ = nw86_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_543_v159_).length(0)):
            d_546_i13_: int = guard_loop_0_
            if (True) and (((0) <= (d_546_i13_)) and ((d_546_i13_) < ((d_543_v159_).length(0)))):
                (d_543_v159_)[(d_546_i13_)] = (d_546_i13_) - (525)
        index77_ = default__.safeIndex(302, (d_339_v2_).length(0))
        (d_339_v2_)[index77_] = d_341_v4_
        index78_ = default__.safeIndex(302, (d_339_v2_).length(0))
        (d_339_v2_)[index78_] = d_341_v4_
        d_547_v160_: int
        d_548_v161_: bool
        d_549_v162_: _dafny.Map
        d_550_v163_: bool
        out28_: int
        out29_: bool
        out30_: _dafny.Map
        out31_: bool
        out28_, out29_, out30_, out31_ = (d_496_v117_).m0(d_347_globalState_)
        d_547_v160_ = out28_
        d_548_v161_ = out29_
        d_549_v162_ = out30_
        d_550_v163_ = out31_
        d_551_v164_: _dafny.Map
        d_551_v164_ = _dafny.Map({d_547_v160_: d_341_v4_})
        d_551_v164_ = (d_551_v164_).set((0) - (d_337_v0_), (d_339_v2_)[default__.safeIndex(302, (d_339_v2_).length(0))])
        d_552_i15_: int
        d_552_i15_ = 0
        with _dafny.label("1"):
            while (len(d_350_v10_)) <= (default__.safeDivisionInt(d_337_v0_, d_337_v0_)):
                with _dafny.c_label("1"):
                    if (d_552_i15_) >= (100):
                        raise _dafny.Break("1")
                    d_552_i15_ = (d_552_i15_) + (1)
                    d_553_v165_: _dafny.MultiSet
                    d_553_v165_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([188 for d_554_i16_ in range(default__.abs(76))]), _dafny.SeqWithoutIsStrInference([d_379_v32_ for d_555_i17_ in range(default__.abs(731))]), d_350_v10_])
                    d_553_v165_ = d_553_v165_
                    d_556_v166_: D1
                    d_556_v166_ = D1_DC5(d_341_v4_, d_548_v161_)
                    index79_ = default__.safeIndex(302, (d_339_v2_).length(0))
                    rhs65_ = ((d_556_v166_).cf6 if (d_339_v2_)[default__.safeIndex(302, (d_339_v2_).length(0))] else d_548_v161_)
                    rhs66_ = d_550_v163_
                    rhs67_ = (D4_DC13(d_340_v3_, d_547_v160_, d_548_v161_, d_339_v2_)).cf22
                    lhs59_ = d_339_v2_
                    lhs60_ = default__.safeIndex(302, (d_339_v2_).length(0))
                    d_550_v163_ = rhs65_
                    lhs59_[lhs60_] = rhs66_
                    d_340_v3_ = rhs67_
                    d_557_v167_: int
                    d_558_v168_: _dafny.Map
                    out32_: int
                    out33_: _dafny.Map
                    out32_, out33_ = default__.m11(d_376_v29_, len(d_340_v3_), True, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([307, d_379_v32_]))).ispropersubset(d_338_v1_), d_347_globalState_)
                    d_557_v167_ = out32_
                    d_558_v168_ = out33_
                    index80_ = default__.safeIndex(302, (d_339_v2_).length(0))
                    (d_339_v2_)[index80_] = False
                    pass
            pass
        if (d_339_v2_)[default__.safeIndex(302, (d_339_v2_).length(0))]:
            d_559_v169_: _dafny.Set
            d_559_v169_ = _dafny.Set({600})
            d_560_v170_: int
            d_561_v171_: int
            d_562_v172_: bool
            out34_: int
            out35_: int
            out36_: bool
            out34_, out35_, out36_ = (d_496_v117_).m1(d_559_v169_, len(d_551_v164_), d_347_globalState_)
            d_560_v170_ = out34_
            d_561_v171_ = out35_
            d_562_v172_ = out36_
            index81_ = default__.safeIndex(302, (d_339_v2_).length(0))
            (d_339_v2_)[index81_] = d_550_v163_
            d_563_v173_: C1
            nw87_ = C1()
            nw87_.ctor__()
            d_563_v173_ = nw87_
            d_564_v174_: T0
            nw88_ = C5()
            nw88_.ctor__(d_547_v160_, not (d_548_v161_) or (d_548_v161_), _dafny.MultiSet([d_341_v4_]))
            d_564_v174_ = nw88_
            index82_ = default__.safeIndex(302, (d_339_v2_).length(0))
            rhs68_ = d_562_v172_
            rhs69_ = d_564_v174_
            lhs61_ = d_339_v2_
            lhs62_ = default__.safeIndex(302, (d_339_v2_).length(0))
            lhs61_[lhs62_] = rhs68_
            d_564_v174_ = rhs69_
            d_550_v163_ = d_548_v161_
        elif True:
            d_565_v175_: int
            d_566_v176_: _dafny.Map
            out37_: int
            out38_: _dafny.Map
            out37_, out38_ = default__.m11(d_376_v29_, default__.fm0((d_339_v2_)[default__.safeIndex(302, (d_339_v2_).length(0))], d_337_v0_, d_337_v0_, default__.fm30(d_548_v161_, d_347_globalState_), d_347_globalState_), (_dafny.MultiSet(d_348_v8_)).issubset(d_495_v116_), d_550_v163_, d_347_globalState_)
            d_565_v175_ = out37_
            d_566_v176_ = out38_
            d_567_v177_: _dafny.Map
            d_567_v177_ = _dafny.Map({len(d_551_v164_): d_340_v3_})
            d_568_v178_: D8
            d_568_v178_ = D8_DC26(d_548_v161_, d_567_v177_)
            d_569_v179_: _dafny.Set
            d_569_v179_ = _dafny.Set({d_568_v178_, d_568_v178_, d_568_v178_, d_568_v178_})
            d_570_v180_: _dafny.Map
            d_570_v180_ = _dafny.Map({d_550_v163_: not(d_548_v161_)})
            index83_ = default__.safeIndex(302, (d_339_v2_).length(0))
            rhs70_ = (_dafny.SeqWithoutIsStrInference([d_550_v163_])) + (_dafny.SeqWithoutIsStrInference([not(d_550_v163_), ((d_570_v180_)[(d_339_v2_)[default__.safeIndex(302, (d_339_v2_).length(0))]] if ((d_339_v2_)[default__.safeIndex(302, (d_339_v2_).length(0))]) in (d_570_v180_) else d_548_v161_)]))
            rhs71_ = d_550_v163_
            rhs72_ = d_569_v179_
            rhs73_ = d_338_v1_
            rhs74_ = (d_339_v2_)[default__.safeIndex(302, (d_339_v2_).length(0))]
            lhs63_ = d_339_v2_
            lhs64_ = default__.safeIndex(302, (d_339_v2_).length(0))
            d_348_v8_ = rhs70_
            d_550_v163_ = rhs71_
            d_569_v179_ = rhs72_
            d_338_v1_ = rhs73_
            lhs63_[lhs64_] = rhs74_
            d_571_v181_: _dafny.Set
            d_571_v181_ = _dafny.Set({d_337_v0_})
            d_572_v182_: int
            d_573_v183_: int
            d_574_v184_: bool
            out39_: int
            out40_: int
            out41_: bool
            out39_, out40_, out41_ = (d_496_v117_).m1((_dafny.Set({(0) - (d_547_v160_), 314})) | (d_571_v181_), d_379_v32_, d_347_globalState_)
            d_572_v182_ = out39_
            d_573_v183_ = out40_
            d_574_v184_ = out41_
            (d_347_globalState_).f3 = d_547_v160_
            d_575_v185_: C0
            nw89_ = C0()
            nw89_.ctor__(not(d_550_v163_), not(False))
            d_575_v185_ = nw89_
        _dafny.print(_dafny.string_of(d_337_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_338_v1_) == (_dafny.MultiSet([206]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_339_v2_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_339_v2_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_339_v2_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_340_v3_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_341_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_342_v5_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_344_v7_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_344_v7_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_344_v7_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_344_v7_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_344_v7_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_344_v7_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_344_v7_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_344_v7_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_344_v7_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_344_v7_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_344_v7_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_347_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_347_globalState_).f1) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_347_globalState_.f2)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_347_globalState_.f2)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_347_globalState_.f2)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_347_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_347_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_347_globalState_).f5).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_347_globalState_.f6) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_347_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_347_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_347_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_347_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_347_globalState_).f12)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_347_globalState_).f12)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_347_globalState_).f12)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_347_globalState_).f12)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_347_globalState_).f12)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_347_globalState_).f12)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_347_globalState_).f12)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_347_globalState_).f12)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_347_globalState_).f12)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_347_globalState_).f12)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_347_globalState_).f12)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_347_globalState_.f13) == (_dafny.Set({True, False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_348_v8_) == (_dafny.SeqWithoutIsStrInference([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v9_) == (_dafny.Map({False: _dafny.Set({True}), True: _dafny.Set({True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_350_v10_) == (_dafny.SeqWithoutIsStrInference([206, 206]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_376_v29_).cf70).cf70).cf70).cf68))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_376_v29_).cf70).cf70).cf70).cf69))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_377_v30_) == (_dafny.Map({0: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_378_v31_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_379_v32_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_380_v33_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_495_v116_) == (_dafny.MultiSet([True, True, True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_533_i11_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v159_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v159_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v159_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v159_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v159_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v159_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v159_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v159_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v159_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v159_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v159_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v159_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v159_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_547_v160_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_548_v161_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_549_v162_) == (_dafny.Map({_dafny.Map({-153: True}): -804}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_550_v163_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_551_v164_) == (_dafny.Map({-153539075: True, -2: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_552_i15_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False)
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
        return lambda: D1_DC4(int(0), int(0))
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
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)

class D1_DC4(D1, NamedTuple('DC4', [('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC11(int(0), int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC11(D3, NamedTuple('DC11', [('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)

class D4_DC13(D4, NamedTuple('DC13', [('cf22', Any), ('cf23', Any), ('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({self.cf22.VerbatimString(True)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC17(_dafny.Seq({}), False, _dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D5_DC18)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D5_DC19)

class D5_DC17(D5, NamedTuple('DC17', [('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC18(D5, NamedTuple('DC18', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC19(D5, NamedTuple('DC19', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC19({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC19) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)

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
        return lambda: D7_DC22(_dafny.Array(None, 0), _dafny.Seq({}), _dafny.Seq({}), _dafny.Array(None, 0), D5.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D7_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D7_DC23)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC22(D7, NamedTuple('DC22', [('cf37', Any), ('cf38', Any), ('cf39', Any), ('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC23(D7, NamedTuple('DC23', [('cf42', Any), ('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC23({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC23) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: D8_DC25()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D8_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D8_DC27)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D8_DC28)

class D8_DC25(D8, NamedTuple('DC25', [])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25)
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC26(D8, NamedTuple('DC26', [('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC26({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC26) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC27(D8, NamedTuple('DC27', [])):
    def __dafnystr__(self) -> str:
        return f'D8.DC27'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC27)
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC28(D8, NamedTuple('DC28', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC28({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC28) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC30(_dafny.CodePoint('D'), False, int(0), _dafny.Set({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D9_DC30)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D9_DC29)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D9_DC31)

class D9_DC30(D9, NamedTuple('DC30', [('cf50', Any), ('cf51', Any), ('cf52', Any), ('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC30({_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC30) and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC29(D9, NamedTuple('DC29', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC29({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC29) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC31(D9, NamedTuple('DC31', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC31({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC31) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D10_DC32)

class D10_DC32(D10, NamedTuple('DC32', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC32({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC32) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC34(_dafny.MultiSet({}), _dafny.MultiSet({}), _dafny.Array(None, 0), D1.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D11_DC34)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D11_DC35)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D11_DC36)

class D11_DC34(D11, NamedTuple('DC34', [('cf58', Any), ('cf59', Any), ('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC34({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC34) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC35(D11, NamedTuple('DC35', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC35'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC35)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC33(D11, NamedTuple('DC33', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC36(D11, NamedTuple('DC36', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC36({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC36) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC38(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D12_DC38)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D12_DC39)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D12_DC37)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D12_DC40)

class D12_DC38(D12, NamedTuple('DC38', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC38({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC38) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC39(D12, NamedTuple('DC39', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC39({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC39) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC37(D12, NamedTuple('DC37', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC37({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC37) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC40(D12, NamedTuple('DC40', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC40({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC40) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC42(_dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D13_DC42)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D13_DC41)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D13_DC43)

class D13_DC42(D13, NamedTuple('DC42', [('cf68', Any), ('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC42({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC42) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC41(D13, NamedTuple('DC41', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC41({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC41) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC43(D13, NamedTuple('DC43', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC43({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC43) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC45(_dafny.Map({}), None, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D14_DC45)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D14_DC44)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D14_DC46)

class D14_DC45(D14, NamedTuple('DC45', [('cf72', Any), ('cf73', Any), ('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC45({_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC45) and self.cf72 == __o.cf72 and self.cf73 == __o.cf73 and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC44(D14, NamedTuple('DC44', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC44({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC44) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC46(D14, NamedTuple('DC46', [('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC46({_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC46) and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC48(int(0), None, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D15_DC48)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D15_DC49)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D15_DC47)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D15_DC50)

class D15_DC48(D15, NamedTuple('DC48', [('cf77', Any), ('cf78', Any), ('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC48({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC48) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78 and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC49(D15, NamedTuple('DC49', [('cf80', Any), ('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC49({_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC49) and self.cf80 == __o.cf80 and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC47(D15, NamedTuple('DC47', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC47({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC47) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC50(D15, NamedTuple('DC50', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC50({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC50) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC52(int(0), False, _dafny.Array(None, 0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D16_DC52)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D16_DC51)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D16_DC53)

class D16_DC52(D16, NamedTuple('DC52', [('cf84', Any), ('cf85', Any), ('cf86', Any), ('cf87', Any), ('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC52({_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC52) and self.cf84 == __o.cf84 and self.cf85 == __o.cf85 and self.cf86 == __o.cf86 and self.cf87 == __o.cf87 and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC51(D16, NamedTuple('DC51', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC51({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC51) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC53(D16, NamedTuple('DC53', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC53({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC53) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D17_DC54)

class D17_DC54(D17, NamedTuple('DC54', [('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC54({_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC54) and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC56(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D18_DC56)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D18_DC57)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D18_DC55)

class D18_DC56(D18, NamedTuple('DC56', [('cf92', Any), ('cf93', Any), ('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC56({_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC56) and self.cf92 == __o.cf92 and self.cf93 == __o.cf93 and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC57(D18, NamedTuple('DC57', [('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC57({_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC57) and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC55(D18, NamedTuple('DC55', [('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC55({_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC55) and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC59(None, _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D19_DC59)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D19_DC58)

class D19_DC59(D19, NamedTuple('DC59', [('cf97', Any), ('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC59({_dafny.string_of(self.cf97)}, {_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC59) and self.cf97 == __o.cf97 and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC58(D19, NamedTuple('DC58', [('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC58({_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC58) and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC61()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D20_DC61)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D20_DC60)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D20_DC62)

class D20_DC61(D20, NamedTuple('DC61', [])):
    def __dafnystr__(self) -> str:
        return f'D20.DC61'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC61)
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC60(D20, NamedTuple('DC60', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC60({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC60) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC62(D20, NamedTuple('DC62', [('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC62({_dafny.string_of(self.cf100)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC62) and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC64(False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D21_DC64)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D21_DC65)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D21_DC63)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D21_DC66)

class D21_DC64(D21, NamedTuple('DC64', [('cf102', Any), ('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC64({_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC64) and self.cf102 == __o.cf102 and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC65(D21, NamedTuple('DC65', [('cf104', Any), ('cf105', Any), ('cf106', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC65({_dafny.string_of(self.cf104)}, {_dafny.string_of(self.cf105)}, {_dafny.string_of(self.cf106)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC65) and self.cf104 == __o.cf104 and self.cf105 == __o.cf105 and self.cf106 == __o.cf106
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC63(D21, NamedTuple('DC63', [('cf101', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC63({_dafny.string_of(self.cf101)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC63) and self.cf101 == __o.cf101
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC66(D21, NamedTuple('DC66', [('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC66({_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC66) and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC68(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D22_DC68)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D22_DC69)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D22_DC70)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D22_DC67)

class D22_DC68(D22, NamedTuple('DC68', [('cf109', Any), ('cf110', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC68({_dafny.string_of(self.cf109)}, {_dafny.string_of(self.cf110)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC68) and self.cf109 == __o.cf109 and self.cf110 == __o.cf110
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC69(D22, NamedTuple('DC69', [('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC69({_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC69) and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC70(D22, NamedTuple('DC70', [('cf112', Any), ('cf113', Any), ('cf114', Any), ('cf115', Any), ('cf116', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC70({_dafny.string_of(self.cf112)}, {_dafny.string_of(self.cf113)}, {self.cf114.VerbatimString(True)}, {_dafny.string_of(self.cf115)}, {_dafny.string_of(self.cf116)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC70) and self.cf112 == __o.cf112 and self.cf113 == __o.cf113 and self.cf114 == __o.cf114 and self.cf115 == __o.cf115 and self.cf116 == __o.cf116
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC67(D22, NamedTuple('DC67', [('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC67({self.cf108.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC67) and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    def m0(self, globalState):
        pass

    def m1(self, p0, p1, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: _dafny.Array = _dafny.Array(None, 0)
        self.f3: int = int(0)
        self.f6: _dafny.Set = _dafny.Set({})
        self.f9: bool = False
        self.f11: _dafny.Array = _dafny.Array(None, 0)
        self.f13: _dafny.Set = _dafny.Set({})
        self._f0: int = int(0)
        self._f1: _dafny.MultiSet = _dafny.MultiSet({})
        self._f4: bool = False
        self._f5: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f7: bool = False
        self._f8: int = int(0)
        self._f10: int = int(0)
        self._f12: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self).f9 = f9
        (self)._f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self).f13 = f13

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
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
    def f12(self):
        return self._f12

class C0:
    def  __init__(self):
        self._f23: bool = False
        self._f24: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f23, f24):
        (self)._f23 = f23
        (self)._f24 = f24

    @property
    def f23(self):
        return self._f23
    @property
    def f24(self):
        return self._f24

class C1:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def m10(self, p0, p1, p2, globalState):
        r0: D1 = D1.default()()
        r1: int = int(0)
        r2: bool = False
        r3: int = int(0)
        if p2:
            d_576_v0_: _dafny.Array
            nw90_ = _dafny.Array(int(0), 9)
            d_576_v0_ = nw90_
            d_577_v1_: D5
            d_577_v1_ = D5_DC16(p1)
            d_576_v0_ = (d_577_v1_).cf29
            d_578_v2_: _dafny.Seq
            d_578_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhcjdj"))
            d_579_v3_: _dafny.Seq
            d_579_v3_ = _dafny.SeqWithoutIsStrInference([d_578_v2_])
            d_580_v4_: int
            d_580_v4_ = -738
            d_581_v5_: C0
            nw91_ = C0()
            nw91_.ctor__(((d_579_v3_)[default__.safeIndex(d_580_v4_, len(d_579_v3_))]) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rylyonx"))), True)
            d_581_v5_ = nw91_
            d_582_v6_: C0
            nw92_ = C0()
            nw92_.ctor__(not((d_581_v5_).f23), (True) or ((d_581_v5_).f23))
            d_582_v6_ = nw92_
            d_583_v7_: _dafny.Map
            d_583_v7_ = _dafny.Map({p2: d_580_v4_})
            r3 = ((0) - (d_580_v4_)) - (default__.safeDivisionInt((0) - (((d_583_v7_)[p2] if (p2) in (d_583_v7_) else d_580_v4_)), d_580_v4_))
            d_584_v8_: _dafny.Seq
            d_584_v8_ = _dafny.SeqWithoutIsStrInference([d_580_v4_])
            d_584_v8_ = d_584_v8_
        elif True:
            d_585_v9_: _dafny.Array
            nw93_ = _dafny.Array(False, 7)
            d_585_v9_ = nw93_
            (globalState).f2 = d_585_v9_
            d_586_v10_: int
            d_586_v10_ = 489
            d_587_v11_: _dafny.Seq
            d_587_v11_ = _dafny.SeqWithoutIsStrInference([d_586_v10_])
            d_588_v12_: _dafny.Map
            d_588_v12_ = _dafny.Map({p2: p2})
            d_589_v13_: _dafny.MultiSet
            d_589_v13_ = _dafny.MultiSet([p2, p2])
            d_590_v14_: _dafny.Array
            nw94_ = _dafny.Array(None, 27)
            nw94_[int(0)] = len(d_587_v11_)
            nw94_[int(1)] = d_586_v10_
            nw94_[int(2)] = (d_586_v10_) - (620)
            nw94_[int(3)] = d_586_v10_
            nw94_[int(4)] = len(d_588_v12_)
            nw94_[int(5)] = default__.fm0(p2, d_586_v10_, d_586_v10_, _dafny.MultiSet([d_586_v10_, d_586_v10_, d_586_v10_]), globalState)
            nw94_[int(6)] = d_586_v10_
            nw94_[int(7)] = d_586_v10_
            nw94_[int(8)] = d_586_v10_
            nw94_[int(9)] = d_586_v10_
            nw94_[int(10)] = d_586_v10_
            nw94_[int(11)] = d_586_v10_
            nw94_[int(12)] = d_586_v10_
            nw94_[int(13)] = d_586_v10_
            nw94_[int(14)] = default__.fm0(p2, d_586_v10_, d_586_v10_, _dafny.MultiSet([550, d_586_v10_]), globalState)
            nw94_[int(15)] = d_586_v10_
            nw94_[int(16)] = default__.safeModuloInt(d_586_v10_, d_586_v10_)
            nw94_[int(17)] = ((d_589_v13_).intersection(d_589_v13_)).cardinality
            nw94_[int(18)] = d_586_v10_
            nw94_[int(19)] = d_586_v10_
            nw94_[int(20)] = d_586_v10_
            nw94_[int(21)] = len(d_587_v11_)
            nw94_[int(22)] = d_586_v10_
            nw94_[int(23)] = d_586_v10_
            nw94_[int(24)] = 589
            nw94_[int(25)] = d_586_v10_
            nw94_[int(26)] = d_586_v10_
            d_590_v14_ = nw94_
            d_590_v14_ = p1
            d_591_v15_: _dafny.Seq
            d_591_v15_ = _dafny.SeqWithoutIsStrInference([p2, p2])
            index84_ = default__.safeIndex(69, (p1).length(0))
            (p1)[index84_] = len(d_591_v15_)
            index85_ = default__.safeIndex(69, (p1).length(0))
            (p1)[index85_] = d_586_v10_
            d_592_v16_: _dafny.Set
            d_592_v16_ = _dafny.Set({p2})
            d_593_v17_: _dafny.MultiSet
            d_593_v17_ = _dafny.MultiSet([len(d_592_v16_)])
            d_594_v18_: D2
            d_594_v18_ = D2_DC7(d_591_v15_)
            d_595_v19_: _dafny.Seq
            d_595_v19_ = _dafny.SeqWithoutIsStrInference([d_594_v18_, d_594_v18_, d_594_v18_, d_594_v18_, d_594_v18_])
            d_587_v11_ = default__.fm21(D1_DC3(default__.fm0(p2, d_586_v10_, (p1)[default__.safeIndex(69, (p1).length(0))], (d_593_v17_).set(d_586_v10_, default__.abs(d_586_v10_)), globalState)), (d_591_v15_)[default__.safeIndex((p1)[default__.safeIndex(69, (p1).length(0))], len(d_591_v15_))], d_595_v19_, globalState)
            d_596_v20_: _dafny.Set
            d_596_v20_ = _dafny.Set({d_587_v11_, d_587_v11_, d_587_v11_, d_587_v11_})
            d_597_v21_: _dafny.Map
            d_597_v21_ = _dafny.Map({not(not(p2)): len(d_596_v20_)})
            d_598_v22_: _dafny.Seq
            d_598_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ye"))
            d_599_v23_: _dafny.Set
            d_599_v23_ = _dafny.Set({d_586_v10_})
            d_600_v24_: _dafny.Seq
            d_600_v24_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({((d_593_v17_)[651] if (651) in (d_593_v17_) else len(d_599_v23_)): p2})])
            d_601_v25_: _dafny.Array
            nw95_ = _dafny.Array(None, 15)
            nw95_[int(0)] = d_597_v21_
            nw95_[int(1)] = (d_597_v21_ if p2 else d_597_v21_)
            nw95_[int(2)] = d_597_v21_
            nw95_[int(3)] = (d_597_v21_) | (d_597_v21_)
            nw95_[int(4)] = (default__.fm1(((d_588_v12_)[p2] if (p2) in (d_588_v12_) else p2), (p1)[default__.safeIndex(69, (p1).length(0))], True, default__.fm22(len(d_598_v22_), d_600_v24_, p2, p2, globalState), globalState)).set(p2, d_586_v10_)
            nw95_[int(5)] = (d_597_v21_).set(False, (p1)[default__.safeIndex(69, (p1).length(0))])
            nw95_[int(6)] = (d_597_v21_) | (d_597_v21_)
            nw95_[int(7)] = d_597_v21_
            nw95_[int(8)] = d_597_v21_
            nw95_[int(9)] = d_597_v21_
            nw95_[int(10)] = d_597_v21_
            nw95_[int(11)] = d_597_v21_
            nw95_[int(12)] = d_597_v21_
            nw95_[int(13)] = d_597_v21_
            nw95_[int(14)] = (d_597_v21_) | (d_597_v21_)
            d_601_v25_ = nw95_
            index86_ = default__.safeIndex(374, (d_601_v25_).length(0))
            (d_601_v25_)[index86_] = (d_597_v21_) | ((d_597_v21_).set(p2, (p1)[default__.safeIndex(69, (p1).length(0))]))
            d_602_v26_: D1
            d_602_v26_ = D1_DC5(True, p2)
            d_603_v27_: _dafny.Set
            d_603_v27_ = _dafny.Set({D1_DC5(not(p2), p2), d_602_v26_, d_602_v26_, d_602_v26_, d_602_v26_})
            index87_ = default__.safeIndex(374, (d_601_v25_).length(0))
            rhs75_ = len((d_603_v27_) | (d_603_v27_))
            rhs76_ = p2
            rhs77_ = (_dafny.Map({p2: (p1)[default__.safeIndex(69, (p1).length(0))]})) | (d_597_v21_)
            rhs78_ = (d_587_v11_) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('h'): _dafny.CodePoint('o')})) for d_604_i0_ in range(default__.abs(624))]))
            lhs65_ = globalState
            lhs66_ = globalState
            lhs67_ = d_601_v25_
            lhs68_ = default__.safeIndex(374, (d_601_v25_).length(0))
            lhs65_.f3 = rhs75_
            lhs66_.f9 = rhs76_
            lhs67_[lhs68_] = rhs77_
            d_587_v11_ = rhs78_
        d_605_v28_: int
        d_605_v28_ = 267
        d_606_v29_: _dafny.Seq
        d_606_v29_ = _dafny.SeqWithoutIsStrInference([d_605_v28_, d_605_v28_])
        d_607_v30_: str
        d_607_v30_ = _dafny.CodePoint('j')
        d_608_v31_: _dafny.Map
        d_608_v31_ = _dafny.Map({d_607_v30_: d_605_v28_})
        d_609_v32_: _dafny.Set
        d_609_v32_ = _dafny.Set({d_606_v29_})
        if default__.fm23(default__.safeDivisionInt(d_605_v28_, d_605_v28_), default__.fm23(d_605_v28_, p2, default__.fm24(d_605_v28_, p2, p2, d_605_v28_, globalState), _dafny.Set({d_606_v29_, d_606_v29_}), globalState), (d_608_v31_) | ((d_608_v31_).set(d_607_v30_, d_605_v28_)), d_609_v32_, globalState):
            d_610_v33_: _dafny.Seq
            d_610_v33_ = _dafny.SeqWithoutIsStrInference([p2, default__.fm23(-497, False, d_608_v31_, d_609_v32_, globalState)])
            d_611_v34_: D2
            d_611_v34_ = D2_DC7(d_610_v33_)
            d_610_v33_ = (d_611_v34_).cf9
            d_612_v35_: _dafny.Seq
            d_612_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjf"))
            d_613_v36_: _dafny.Map
            d_613_v36_ = _dafny.Map({d_605_v28_: len(d_612_v35_)})
            d_614_v37_: D1
            d_614_v37_ = D1_DC4(len(d_613_v36_), 112)
            source10_ = d_614_v37_
            if source10_.is_DC4:
                d_615___mcc_h0_ = source10_.cf4
                d_616___mcc_h1_ = source10_.cf5
                d_617_cf5_ = d_616___mcc_h1_
                d_618_cf4_ = d_615___mcc_h0_
                d_619_v38_: D1
                d_619_v38_ = D1_DC3(d_617_cf5_)
                d_620_v39_: _dafny.Array
                nw96_ = _dafny.Array(None, 29)
                nw96_[int(0)] = p2
                nw96_[int(1)] = not(not(p2))
                nw96_[int(2)] = p2
                nw96_[int(3)] = False
                nw96_[int(4)] = default__.fm23(d_605_v28_, True, d_608_v31_, _dafny.Set({default__.fm21(d_619_v38_, p2, _dafny.SeqWithoutIsStrInference([d_611_v34_ for d_621_i1_ in range(default__.abs(-918))]), globalState)}), globalState)
                nw96_[int(5)] = p2
                nw96_[int(6)] = p2
                nw96_[int(7)] = p2
                nw96_[int(8)] = p2
                nw96_[int(9)] = p2
                nw96_[int(10)] = False
                nw96_[int(11)] = p2
                nw96_[int(12)] = p2
                nw96_[int(13)] = p2
                nw96_[int(14)] = not((d_610_v33_)[default__.safeIndex(d_617_cf5_, len(d_610_v33_))])
                nw96_[int(15)] = False
                nw96_[int(16)] = p2
                nw96_[int(17)] = True
                nw96_[int(18)] = p2
                nw96_[int(19)] = True
                nw96_[int(20)] = p2
                nw96_[int(21)] = p2
                nw96_[int(22)] = p2
                nw96_[int(23)] = (d_610_v33_)[default__.safeIndex(d_618_cf4_, len(d_610_v33_))]
                nw96_[int(24)] = not(False)
                nw96_[int(25)] = p2
                nw96_[int(26)] = p2
                nw96_[int(27)] = p2
                nw96_[int(28)] = True
                d_620_v39_ = nw96_
                d_622_v40_: _dafny.Map
                d_622_v40_ = _dafny.Map({(d_610_v33_)[default__.safeIndex(d_618_cf4_, len(d_610_v33_))]: d_620_v39_})
                d_623_v41_: C0
                nw97_ = C0()
                nw97_.ctor__(p2, not(p2))
                d_623_v41_ = nw97_
                d_624_v42_: _dafny.Seq
                d_624_v42_ = _dafny.SeqWithoutIsStrInference([d_623_v41_, d_623_v41_, d_623_v41_])
                index88_ = default__.safeIndex(170, (d_620_v39_).length(0))
                (d_620_v39_)[index88_] = (d_624_v42_) == (d_624_v42_)
                index89_ = default__.safeIndex(170, (d_620_v39_).length(0))
                rhs79_ = d_612_v35_
                rhs80_ = d_622_v40_
                rhs81_ = ((d_623_v41_).f23 if (d_618_cf4_) < (d_618_cf4_) else p2)
                lhs69_ = d_620_v39_
                lhs70_ = default__.safeIndex(170, (d_620_v39_).length(0))
                d_612_v35_ = rhs79_
                d_622_v40_ = rhs80_
                lhs69_[lhs70_] = rhs81_
                d_625_v43_: _dafny.Array
                nw98_ = _dafny.Array(D2.default()(), 28)
                d_625_v43_ = nw98_
                index90_ = default__.safeIndex(395, (d_625_v43_).length(0))
                (d_625_v43_)[index90_] = d_611_v34_
                d_626_v44_: _dafny.Map
                d_626_v44_ = _dafny.Map({(d_620_v39_)[default__.safeIndex(170, (d_620_v39_).length(0))]: 284})
                index91_ = default__.safeIndex(395, (d_625_v43_).length(0))
                rhs82_ = default__.fm25(d_618_cf4_, (0) - ((d_618_cf4_) + (d_618_cf4_)), (d_623_v41_).f23, d_626_v44_, globalState)
                rhs83_ = d_623_v41_
                rhs84_ = d_607_v30_
                rhs85_ = 948
                rhs86_ = d_611_v34_
                lhs71_ = d_625_v43_
                lhs72_ = default__.safeIndex(395, (d_625_v43_).length(0))
                d_607_v30_ = rhs82_
                d_623_v41_ = rhs83_
                d_607_v30_ = rhs84_
                d_617_cf5_ = rhs85_
                lhs71_[lhs72_] = rhs86_
                d_627_v45_: _dafny.Map
                d_627_v45_ = _dafny.Map({p1: (d_623_v41_).f23})
                d_627_v45_ = (d_627_v45_).set(p1, (d_623_v41_).f23)
                d_628_v46_: _dafny.MultiSet
                d_628_v46_ = _dafny.MultiSet([(0) - (d_605_v28_)])
                d_629_v47_: _dafny.Map
                d_629_v47_ = _dafny.Map({p1: _dafny.Set({(d_620_v39_)[default__.safeIndex(170, (d_620_v39_).length(0))]})})
                rhs87_ = d_628_v46_
                rhs88_ = d_620_v39_
                rhs89_ = d_612_v35_
                rhs90_ = d_629_v47_
                rhs91_ = 79
                lhs73_ = globalState
                d_628_v46_ = rhs87_
                lhs73_.f2 = rhs88_
                d_612_v35_ = rhs89_
                d_629_v47_ = rhs90_
                r3 = rhs91_
            elif source10_.is_DC5:
                d_630___mcc_h2_ = source10_.cf6
                d_631___mcc_h3_ = source10_.cf7
                d_632_cf7_ = d_631___mcc_h3_
                d_633_cf6_ = d_630___mcc_h2_
                d_634_v48_: D4
                d_634_v48_ = D4_DC14(d_632_cf7_, False)
                d_635_v49_: _dafny.Map
                d_635_v49_ = _dafny.Map({d_634_v48_: (d_632_cf7_ if d_633_cf6_ else True)})
                d_635_v49_ = (d_635_v49_).set(d_634_v48_, (d_610_v33_)[default__.safeIndex(d_605_v28_, len(d_610_v33_))])
                d_636_v50_: _dafny.Array
                nw99_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_636_v50_ = nw99_
                d_637_v51_: _dafny.Array
                nw100_ = _dafny.Array(False, 8)
                d_637_v51_ = nw100_
                index92_ = default__.safeIndex(641, (d_636_v50_).length(0))
                (d_636_v50_)[index92_] = d_637_v51_
                index93_ = default__.safeIndex(641, (d_636_v50_).length(0))
                (d_636_v50_)[index93_] = d_637_v51_
                arr0_ = (d_636_v50_)[default__.safeIndex(641, (d_636_v50_).length(0))]
                index94_ = default__.safeIndex(855, ((d_636_v50_)[default__.safeIndex(641, (d_636_v50_).length(0))]).length(0))
                arr0_[index94_] = p2
                arr1_ = (d_636_v50_)[default__.safeIndex(641, (d_636_v50_).length(0))]
                index95_ = default__.safeIndex(855, ((d_636_v50_)[default__.safeIndex(641, (d_636_v50_).length(0))]).length(0))
                arr1_[index95_] = p2
                d_638_v52_: _dafny.Seq
                d_638_v52_ = _dafny.SeqWithoutIsStrInference([d_612_v35_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbocxqw")), d_612_v35_])
                d_639_v53_: _dafny.Map
                d_639_v53_ = _dafny.Map({d_606_v29_: (d_638_v52_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_605_v28_, d_605_v28_, d_605_v28_])), len(d_638_v52_))]})
                (globalState).f9 = (d_606_v29_) not in (d_639_v53_)
            elif source10_.is_DC3:
                d_640___mcc_h4_ = source10_.cf3
                d_641_cf3_ = d_640___mcc_h4_
                rhs92_ = d_605_v28_
                rhs93_ = (default__.safeModuloInt(d_641_cf3_, 324)) * (len(d_612_v35_))
                r3 = rhs92_
                r3 = rhs93_
                d_642_v54_: _dafny.Array
                def lambda21_(d_643_p2_):
                    def lambda22_(d_644_i2_):
                        return d_643_p2_

                    return lambda22_

                init10_ = lambda21_(p2)
                nw101_ = _dafny.Array(None, 24)
                for i0_10_ in range(nw101_.length(0)):
                    nw101_[i0_10_] = init10_(i0_10_)
                d_642_v54_ = nw101_
                d_645_v55_: D4
                d_645_v55_ = D4_DC13(d_612_v35_, 930, p2, d_642_v54_)
                (globalState).f2 = (d_645_v55_).cf25
                index96_ = default__.safeIndex(457, (d_642_v54_).length(0))
                (d_642_v54_)[index96_] = p2
                index97_ = default__.safeIndex(457, (d_642_v54_).length(0))
                (d_642_v54_)[index97_] = p2
                d_605_v28_ = d_605_v28_
            elif True:
                d_646___mcc_h5_ = source10_.cf8
                d_647_cf8_ = d_646___mcc_h5_
                d_648_v56_: _dafny.Map
                d_648_v56_ = _dafny.Map({p2: True})
                d_649_v57_: _dafny.Array
                nw102_ = _dafny.Array(None, 27)
                nw102_[int(0)] = p2
                nw102_[int(1)] = p2
                nw102_[int(2)] = p2
                nw102_[int(3)] = p2
                nw102_[int(4)] = p2
                nw102_[int(5)] = p2
                nw102_[int(6)] = True
                nw102_[int(7)] = p2
                nw102_[int(8)] = p2
                nw102_[int(9)] = p2
                nw102_[int(10)] = ((d_648_v56_)[p2] if (p2) in (d_648_v56_) else p2)
                nw102_[int(11)] = p2
                nw102_[int(12)] = p2
                nw102_[int(13)] = p2
                nw102_[int(14)] = p2
                nw102_[int(15)] = p2
                nw102_[int(16)] = p2
                nw102_[int(17)] = (d_610_v33_)[default__.safeIndex(d_605_v28_, len(d_610_v33_))]
                nw102_[int(18)] = p2
                nw102_[int(19)] = p2
                nw102_[int(20)] = not(False)
                nw102_[int(21)] = p2
                nw102_[int(22)] = (d_610_v33_)[default__.safeIndex(d_605_v28_, len(d_610_v33_))]
                nw102_[int(23)] = True
                nw102_[int(24)] = p2
                nw102_[int(25)] = p2
                nw102_[int(26)] = False
                d_649_v57_ = nw102_
                d_650_v58_: D4
                d_650_v58_ = D4_DC13(d_612_v35_, len((d_612_v35_).set(default__.safeIndex(747, len(d_612_v35_)), d_607_v30_)), not(False), d_649_v57_)
                d_651_v59_: D4
                d_651_v59_ = D4_DC15(d_650_v58_)
                d_652_v60_: _dafny.Seq
                d_652_v60_ = _dafny.SeqWithoutIsStrInference([d_651_v59_, d_651_v59_])
                r1 = (d_605_v28_) - ((_dafny.MultiSet((d_652_v60_) + (d_652_v60_))).cardinality)
                r2 = (434) != (22)
                d_653_v61_: _dafny.MultiSet
                d_653_v61_ = _dafny.MultiSet([d_605_v28_])
                d_654_v62_: D4
                d_654_v62_ = D4_DC13(d_612_v35_, (d_653_v61_).cardinality, True, d_649_v57_)
                (globalState).f3 = (d_654_v62_).cf23
                d_655_v63_: _dafny.Map
                d_655_v63_ = _dafny.Map({d_605_v28_: p2})
                d_656_v64_: _dafny.Seq
                d_656_v64_ = _dafny.SeqWithoutIsStrInference([p1])
                d_657_v65_: _dafny.Map
                d_657_v65_ = _dafny.Map({True: d_605_v28_})
                d_658_v66_: _dafny.Map
                d_658_v66_ = _dafny.Map({d_655_v63_: (len(d_656_v64_)) != (((d_657_v65_)[p2] if (p2) in (d_657_v65_) else len(d_657_v65_)))})
                d_658_v66_ = d_658_v66_
            d_659_v67_: D1
            d_659_v67_ = D1_DC5(p2, p2)
            d_659_v67_ = d_659_v67_
            (globalState).f3 = d_605_v28_
            d_660_v68_: C0
            nw103_ = C0()
            nw103_.ctor__(p2, p2)
            d_660_v68_ = nw103_
        elif True:
            d_661_v69_: _dafny.Array
            nw104_ = _dafny.Array(_dafny.CodePoint('D'), 18)
            d_661_v69_ = nw104_
            d_662_v70_: _dafny.MultiSet
            d_662_v70_ = _dafny.MultiSet([d_661_v69_, d_661_v69_, d_661_v69_, d_661_v69_, d_661_v69_])
            d_662_v70_ = _dafny.MultiSet([d_661_v69_])
            d_663_v71_: _dafny.MultiSet
            d_663_v71_ = _dafny.MultiSet([d_605_v28_])
            d_664_v72_: _dafny.Map
            d_664_v72_ = _dafny.Map({p2: (_dafny.SeqWithoutIsStrInference([d_607_v30_ for d_665_i3_ in range(default__.abs(707))])).set(default__.safeIndex(default__.fm0(p2, d_605_v28_, d_605_v28_, d_663_v71_, globalState), len(_dafny.SeqWithoutIsStrInference([d_607_v30_ for d_666_i3_ in range(default__.abs(707))]))), d_607_v30_)})
            rhs94_ = ((d_664_v72_) | (d_664_v72_)) | (d_664_v72_)
            rhs95_ = _dafny.SeqWithoutIsStrInference([(D1_DC4(-838, d_605_v28_)).cf4 for d_667_i4_ in range(default__.abs(59))])
            d_664_v72_ = rhs94_
            d_606_v29_ = rhs95_
            d_668_v73_: _dafny.Array
            def lambda23_(d_669_v30_):
                def lambda24_(d_670_i5_):
                    return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h')])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_669_v30_]), _dafny.SeqWithoutIsStrInference([d_669_v30_, d_669_v30_])]))

                return lambda24_

            init11_ = lambda23_(d_607_v30_)
            nw105_ = _dafny.Array(None, 7)
            for i0_11_ in range(nw105_.length(0)):
                nw105_[i0_11_] = init11_(i0_11_)
            d_668_v73_ = nw105_
            d_671_v74_: _dafny.Seq
            d_671_v74_ = _dafny.SeqWithoutIsStrInference([d_607_v30_, d_607_v30_])
            d_672_v75_: _dafny.Seq
            d_672_v75_ = _dafny.SeqWithoutIsStrInference([d_671_v74_, _dafny.SeqWithoutIsStrInference([d_607_v30_, d_607_v30_]), d_671_v74_])
            index98_ = default__.safeIndex(920, (d_668_v73_).length(0))
            (d_668_v73_)[index98_] = d_672_v75_
            index99_ = default__.safeIndex(920, (d_668_v73_).length(0))
            (d_668_v73_)[index99_] = (d_672_v75_) + ((d_672_v75_) + (default__.fm26(d_605_v28_, p2, len(d_671_v74_), globalState)))
            d_673_v76_: _dafny.Array
            def lambda25_(d_674_i6_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tulofkoow"))

            init12_ = lambda25_
            nw106_ = _dafny.Array(None, 6)
            for i0_12_ in range(nw106_.length(0)):
                nw106_[i0_12_] = init12_(i0_12_)
            d_673_v76_ = nw106_
            index100_ = default__.safeIndex(573, (d_673_v76_).length(0))
            (d_673_v76_)[index100_] = (_dafny.SeqWithoutIsStrInference([d_607_v30_ for d_675_i7_ in range(default__.abs(600))])) + ((d_671_v74_).set(default__.safeIndex(d_605_v28_, len(d_671_v74_)), d_607_v30_))
            index101_ = default__.safeIndex(573, (d_673_v76_).length(0))
            (d_673_v76_)[index101_] = (default__.fm27((0) - (d_605_v28_), p2, globalState)) + (d_671_v74_)
            d_676_v77_: _dafny.Seq
            d_676_v77_ = _dafny.SeqWithoutIsStrInference([True, p2, p2, p2, p2])
            d_677_v78_: _dafny.Map
            d_677_v78_ = _dafny.Map({len(d_676_v77_): d_605_v28_})
            d_678_v79_: _dafny.Map
            d_678_v79_ = _dafny.Map({d_677_v78_: d_605_v28_})
            (globalState).f3 = len((d_678_v79_ if p2 else d_678_v79_))
        index102_ = default__.safeIndex(15, (p1).length(0))
        (p1)[index102_] = len(d_606_v29_)
        index103_ = default__.safeIndex(15, (p1).length(0))
        (p1)[index103_] = d_605_v28_
        r3 = (d_605_v28_) * ((p1)[default__.safeIndex(15, (p1).length(0))])
        hi1_ = d_605_v28_
        for d_679_i8_ in range((p1)[default__.safeIndex(15, (p1).length(0))], hi1_):
            d_680_v80_: C0
            nw107_ = C0()
            nw107_.ctor__((p2 if p2 else False), (p2) and (False))
            d_680_v80_ = nw107_
            if (d_605_v28_) <= (d_679_i8_):
                r2 = (d_680_v80_).f24
                d_681_v81_: _dafny.Seq
                d_681_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpi"))
                d_681_v81_ = d_681_v81_
                d_681_v81_ = d_681_v81_
                d_682_v82_: _dafny.Map
                d_682_v82_ = _dafny.Map({d_680_v80_: default__.fm27((0) - (d_679_i8_), p2, globalState)})
                d_683_v83_: _dafny.Array
                nw108_ = _dafny.Array(False, 14)
                d_683_v83_ = nw108_
                d_684_v84_: D4
                d_684_v84_ = D4_DC13(_dafny.SeqWithoutIsStrInference([d_607_v30_ for d_685_i9_ in range(default__.abs(103))]), 12, (d_680_v80_).f23, d_683_v83_)
                d_686_v85_: _dafny.Seq
                d_686_v85_ = _dafny.SeqWithoutIsStrInference([d_681_v81_, d_681_v81_])
                d_687_v86_: _dafny.Array
                nw109_ = _dafny.Array(None, 25)
                nw109_[int(0)] = d_681_v81_
                nw109_[int(1)] = ((d_682_v82_)[d_680_v80_] if (d_680_v80_) in (d_682_v82_) else d_681_v81_)
                nw109_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmsifhi"))) + (d_681_v81_)
                nw109_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sta"))
                nw109_[int(4)] = d_681_v81_
                nw109_[int(5)] = (d_681_v81_) + (default__.fm27(len(d_681_v81_), (d_680_v80_).f24, globalState))
                nw109_[int(6)] = d_681_v81_
                nw109_[int(7)] = default__.fm27(d_605_v28_, (d_680_v80_).f23, globalState)
                nw109_[int(8)] = d_681_v81_
                nw109_[int(9)] = d_681_v81_
                nw109_[int(10)] = d_681_v81_
                nw109_[int(11)] = (d_684_v84_).cf22
                nw109_[int(12)] = d_681_v81_
                nw109_[int(13)] = d_681_v81_
                nw109_[int(14)] = (d_681_v81_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_688_i10_ in range(default__.abs(462))]))
                nw109_[int(15)] = (d_686_v85_)[default__.safeIndex(d_605_v28_, len(d_686_v85_))]
                nw109_[int(16)] = ((d_681_v81_).set(default__.safeIndex(d_679_i8_, len(d_681_v81_)), d_607_v30_)).set(default__.safeIndex((p1)[default__.safeIndex(15, (p1).length(0))], len((d_681_v81_).set(default__.safeIndex(d_679_i8_, len(d_681_v81_)), d_607_v30_))), d_607_v30_)
                nw109_[int(17)] = d_681_v81_
                nw109_[int(18)] = d_681_v81_
                nw109_[int(19)] = (d_684_v84_).cf22
                nw109_[int(20)] = (d_681_v81_).set(default__.safeIndex(len(d_681_v81_), len(d_681_v81_)), d_607_v30_)
                nw109_[int(21)] = (d_681_v81_) + (d_681_v81_)
                nw109_[int(22)] = d_681_v81_
                nw109_[int(23)] = (d_681_v81_ if p2 else d_681_v81_)
                nw109_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhvna"))
                d_687_v86_ = nw109_
                d_687_v86_ = d_687_v86_
                d_689_v87_: _dafny.Array
                def lambda26_(d_690_v80_):
                    def lambda27_(d_691_i11_):
                        return _dafny.SeqWithoutIsStrInference([(d_690_v80_).f24, True, not(not((d_690_v80_).f23))])

                    return lambda27_

                init13_ = lambda26_(d_680_v80_)
                nw110_ = _dafny.Array(None, 21)
                for i0_13_ in range(nw110_.length(0)):
                    nw110_[i0_13_] = init13_(i0_13_)
                d_689_v87_ = nw110_
                d_692_v88_: _dafny.Seq
                d_692_v88_ = _dafny.SeqWithoutIsStrInference([(d_680_v80_).f24, False])
                index104_ = default__.safeIndex(157, (d_689_v87_).length(0))
                (d_689_v87_)[index104_] = d_692_v88_
                index105_ = default__.safeIndex(157, (d_689_v87_).length(0))
                (d_689_v87_)[index105_] = d_692_v88_
            elif True:
                d_693_v89_: _dafny.Seq
                d_693_v89_ = _dafny.SeqWithoutIsStrInference([d_607_v30_])
                (globalState).f3 = default__.fm0(default__.fm23(d_679_i8_, (d_680_v80_).f23, d_608_v31_, _dafny.Set({d_606_v29_}), globalState), (len(d_693_v89_)) + (d_605_v28_), len(_dafny.Set({(p1)[default__.safeIndex(15, (p1).length(0))]})), _dafny.MultiSet([d_679_i8_]), globalState)
                (globalState).f9 = False
                (globalState).f9 = (d_605_v28_) > (d_605_v28_)
                d_694_v90_: C0
                nw111_ = C0()
                nw111_.ctor__((d_680_v80_).f23, (d_680_v80_).f24)
                d_694_v90_ = nw111_
                d_695_v91_: _dafny.Array
                nw112_ = _dafny.Array(False, 11)
                d_695_v91_ = nw112_
                d_696_v92_: _dafny.MultiSet
                d_696_v92_ = _dafny.MultiSet([d_695_v91_, d_695_v91_])
                (globalState).f9 = not((((d_696_v92_).set(d_695_v91_, default__.abs(911))) - (d_696_v92_)) == (d_696_v92_))
            r3 = (p1)[default__.safeIndex(15, (p1).length(0))]
            d_697_v93_: D5
            d_697_v93_ = D5_DC18((d_680_v80_).f24)
            d_698_v94_: _dafny.Map
            d_698_v94_ = _dafny.Map({d_697_v93_: d_605_v28_})
            (globalState).f9 = (d_605_v28_) > ((d_605_v28_) * (len(d_698_v94_)))
        index106_ = default__.safeIndex(15, (p1).length(0))
        (p1)[index106_] = (p1)[default__.safeIndex(15, (p1).length(0))]
        r0 = default__.fm28(p2, globalState)
        r1 = (p1)[default__.safeIndex(15, (p1).length(0))]
        d_699_v95_: _dafny.Seq
        d_699_v95_ = _dafny.SeqWithoutIsStrInference([not(p2)])
        d_700_v96_: _dafny.Seq
        d_700_v96_ = _dafny.SeqWithoutIsStrInference([d_699_v95_, (d_699_v95_).set(default__.safeIndex(len(_dafny.Map({(p1)[default__.safeIndex(15, (p1).length(0))]: (p1)[default__.safeIndex(15, (p1).length(0))]})), len(d_699_v95_)), p2)])
        d_701_v97_: _dafny.MultiSet
        d_701_v97_ = _dafny.MultiSet([(0) - (d_605_v28_)])
        pat_let_tv19_ = p2
        d_702_v98_: _dafny.Seq
        def iife76_(_pat_let19_0):
            def iife77_(d_703_dt__update__tmp_h0_):
                def iife78_(_pat_let20_0):
                    def iife79_(d_704_dt__update_hcf31_h0_):
                        return D5_DC17((d_703_dt__update__tmp_h0_).cf30, d_704_dt__update_hcf31_h0_, (d_703_dt__update__tmp_h0_).cf32)
                    return iife79_(_pat_let20_0)
                return iife78_(pat_let_tv19_)
            return iife77_(_pat_let19_0)
        d_702_v98_ = _dafny.SeqWithoutIsStrInference([iife76_(D5_DC17(d_700_v96_, p2, d_701_v97_))])
        r2 = default__.fm23((p1)[default__.safeIndex(15, (p1).length(0))], p2, (_dafny.Map({d_607_v30_: len(d_702_v98_)})) | (d_608_v31_), d_609_v32_, globalState)
        r3 = d_605_v28_
        return r0, r1, r2, r3


class C2:
    def  __init__(self):
        self.f22: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f22):
        (self).f22 = f22

    def m8(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: D3 = D3.default()()
        d_705_v0_: bool
        d_705_v0_ = True
        d_706_i0_: int
        d_706_i0_ = 0
        with _dafny.label("2"):
            while d_705_v0_:
                with _dafny.c_label("2"):
                    if (d_706_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_706_i0_ = (d_706_i0_) + (1)
                    (globalState).f9 = d_705_v0_
                    d_707_v1_: str
                    d_707_v1_ = _dafny.CodePoint('y')
                    r0 = d_707_v1_
                    d_708_v2_: D0
                    d_708_v2_ = D0_DC1(d_705_v0_)
                    source11_ = d_708_v2_
                    if source11_.is_DC1:
                        d_709___mcc_h0_ = source11_.cf1
                        d_710_cf1_ = d_709___mcc_h0_
                        d_711_v3_: int
                        d_711_v3_ = 8
                        (globalState).f3 = d_711_v3_
                        d_712_v4_: _dafny.Array
                        def lambda28_(d_713_v3_, d_714_cf1_, d_715_v0_):
                            def lambda29_(d_716_i1_):
                                return (_dafny.SeqWithoutIsStrInference([D1_DC4(d_713_v3_, d_713_v3_) for d_717_i2_ in range(default__.abs(981))])) + (_dafny.SeqWithoutIsStrInference([D1_DC4(len(self.f22), len(_dafny.Map({d_713_v3_: len(_dafny.Map({d_714_cf1_: d_715_v0_}))})))]))

                            return lambda29_

                        init14_ = lambda28_(d_711_v3_, d_710_cf1_, d_705_v0_)
                        nw113_ = _dafny.Array(None, 24)
                        for i0_14_ in range(nw113_.length(0)):
                            nw113_[i0_14_] = init14_(i0_14_)
                        d_712_v4_ = nw113_
                        d_718_v5_: _dafny.Map
                        d_718_v5_ = _dafny.Map({d_710_cf1_: d_710_cf1_})
                        d_719_v6_: _dafny.Map
                        d_719_v6_ = _dafny.Map({d_707_v1_: len(d_718_v5_)})
                        d_720_v7_: _dafny.Seq
                        d_720_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pr"))
                        d_721_v8_: _dafny.Seq
                        d_721_v8_ = _dafny.SeqWithoutIsStrInference([((d_719_v6_)[d_707_v1_] if (d_707_v1_) in (d_719_v6_) else d_711_v3_), d_711_v3_, len(d_720_v7_)])
                        d_722_v9_: _dafny.Array
                        nw114_ = _dafny.Array(False, 17)
                        d_722_v9_ = nw114_
                        d_723_v10_: _dafny.Map
                        d_723_v10_ = _dafny.Map({d_722_v9_: d_707_v1_})
                        d_724_v11_: _dafny.Seq
                        d_724_v11_ = _dafny.SeqWithoutIsStrInference([d_705_v0_, d_710_cf1_])
                        d_725_v12_: _dafny.Seq
                        d_725_v12_ = _dafny.SeqWithoutIsStrInference([d_724_v11_, d_724_v11_])
                        d_726_v13_: _dafny.Array
                        nw115_ = _dafny.Array(None, 8)
                        nw115_[int(0)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cudqbqnau")))
                        nw115_[int(1)] = -98
                        nw115_[int(2)] = (d_711_v3_) + (d_711_v3_)
                        nw115_[int(3)] = d_711_v3_
                        nw115_[int(4)] = len((_dafny.SeqWithoutIsStrInference([d_711_v3_]) if False else (d_721_v8_).set(default__.safeIndex(d_711_v3_, len(d_721_v8_)), len(d_723_v10_))))
                        nw115_[int(5)] = (d_711_v3_) - (len(d_721_v8_))
                        nw115_[int(6)] = default__.safeDivisionInt(default__.fm0(d_710_cf1_, d_711_v3_, 441, _dafny.MultiSet(d_721_v8_), globalState), (0) - (len((d_725_v12_)[default__.safeIndex(d_711_v3_, len(d_725_v12_))])))
                        nw115_[int(7)] = (len(d_721_v8_)) + (558)
                        d_726_v13_ = nw115_
                        d_727_v14_: _dafny.Array
                        nw116_ = _dafny.Array(None, 13)
                        nw116_[int(0)] = d_708_v2_
                        nw116_[int(1)] = d_708_v2_
                        nw116_[int(2)] = d_708_v2_
                        nw116_[int(3)] = D0_DC1(((self.f22)[d_711_v3_] if (d_711_v3_) in (self.f22) else False))
                        nw116_[int(4)] = d_708_v2_
                        nw116_[int(5)] = d_708_v2_
                        nw116_[int(6)] = d_708_v2_
                        nw116_[int(7)] = d_708_v2_
                        nw116_[int(8)] = d_708_v2_
                        nw116_[int(9)] = d_708_v2_
                        nw116_[int(10)] = d_708_v2_
                        nw116_[int(11)] = d_708_v2_
                        nw116_[int(12)] = d_708_v2_
                        d_727_v14_ = nw116_
                        index107_ = default__.safeIndex(203, (d_727_v14_).length(0))
                        (d_727_v14_)[index107_] = d_708_v2_
                        d_728_v15_: _dafny.Seq
                        d_728_v15_ = _dafny.SeqWithoutIsStrInference([d_712_v4_])
                        d_729_v16_: _dafny.Map
                        d_729_v16_ = _dafny.Map({d_711_v3_: p0})
                        d_730_v17_: _dafny.MultiSet
                        d_730_v17_ = _dafny.MultiSet([len(self.f22)])
                        pat_let_tv20_ = d_710_cf1_
                        index108_ = default__.safeIndex(203, (d_727_v14_).length(0))
                        def iife80_(_pat_let21_0):
                            def iife81_(d_731_dt__update__tmp_h0_):
                                def iife82_(_pat_let22_0):
                                    def iife83_(d_732_dt__update_hcf1_h0_):
                                        return D0_DC1(d_732_dt__update_hcf1_h0_)
                                    return iife83_(_pat_let22_0)
                                return iife82_(pat_let_tv20_)
                            return iife81_(_pat_let21_0)
                        rhs96_ = (d_728_v15_)[default__.safeIndex((d_711_v3_) * (len(d_729_v16_)), len(d_728_v15_))]
                        rhs97_ = default__.fm0(False, (d_711_v3_) + (358), d_711_v3_, d_730_v17_, globalState)
                        rhs98_ = d_726_v13_
                        rhs99_ = (D4_DC12(d_722_v9_)).cf21
                        rhs100_ = iife80_(d_708_v2_)
                        lhs74_ = d_727_v14_
                        lhs75_ = default__.safeIndex(203, (d_727_v14_).length(0))
                        d_712_v4_ = rhs96_
                        d_711_v3_ = rhs97_
                        d_726_v13_ = rhs98_
                        d_722_v9_ = rhs99_
                        lhs74_[lhs75_] = rhs100_
                        d_733_v18_: C1
                        nw117_ = C1()
                        nw117_.ctor__()
                        d_733_v18_ = nw117_
                        index109_ = default__.safeIndex(418, (p0).length(0))
                        (p0)[index109_] = d_711_v3_
                        index110_ = default__.safeIndex(418, (p0).length(0))
                        (p0)[index110_] = ((d_711_v3_) * (d_711_v3_)) * ((d_711_v3_) - (d_711_v3_))
                    elif source11_.is_DC0:
                        d_734___mcc_h1_ = source11_.cf0
                        d_735_cf0_ = d_734___mcc_h1_
                        d_736_v19_: D0
                        d_736_v19_ = D0_DC0(d_705_v0_)
                        d_736_v19_ = d_736_v19_
                        d_737_v20_: _dafny.Map
                        d_737_v20_ = _dafny.Map({(d_735_cf0_) or (d_705_v0_): _dafny.SeqWithoutIsStrInference([d_707_v1_ for d_738_i3_ in range(default__.abs(529))])})
                        d_739_v21_: _dafny.Seq
                        d_739_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lksdhhcbj"))
                        d_737_v20_ = (d_737_v20_).set(d_735_cf0_, d_739_v21_)
                        d_740_v23_: int
                        out42_: int
                        def iife84_():
                            coll38_ = _dafny.Set()
                            compr_38_: int
                            for compr_38_ in _dafny.IntegerRange(502, 693):
                                d_741_v22_: int = compr_38_
                                if ((502) <= (d_741_v22_)) and ((d_741_v22_) < (693)):
                                    coll38_ = coll38_.union(_dafny.Set([default__.safeDivisionInt(d_741_v22_, len(d_739_v21_))]))
                            return _dafny.Set(coll38_)
                        out42_ = (self).m9(len(iife84_()
                        ), globalState)
                        d_740_v23_ = out42_
                        d_742_v24_: _dafny.Array
                        nw118_ = _dafny.Array(None, 3)
                        nw118_[int(0)] = d_735_cf0_
                        nw118_[int(1)] = d_735_cf0_
                        nw118_[int(2)] = d_735_cf0_
                        d_742_v24_ = nw118_
                        index111_ = default__.safeIndex(461, (d_742_v24_).length(0))
                        (d_742_v24_)[index111_] = d_705_v0_
                        index112_ = default__.safeIndex(461, (d_742_v24_).length(0))
                        (d_742_v24_)[index112_] = d_705_v0_
                    elif True:
                        d_743___mcc_h2_ = source11_.cf2
                        d_744_cf2_ = d_743___mcc_h2_
                        d_745_v25_: int
                        d_745_v25_ = 921
                        r1 = d_745_v25_
                        index113_ = default__.safeIndex(63, (p0).length(0))
                        (p0)[index113_] = 887
                        d_746_v26_: _dafny.Seq
                        d_746_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfyhouxd"))
                        index114_ = default__.safeIndex(63, (p0).length(0))
                        (p0)[index114_] = default__.safeDivisionInt(d_745_v25_, (-526) + (len(d_746_v26_)))
                        d_747_v27_: _dafny.Array
                        nw119_ = _dafny.Array(int(0), 14)
                        d_747_v27_ = nw119_
                        d_747_v27_ = d_747_v27_
                        d_748_v28_: _dafny.Array
                        nw120_ = _dafny.Array(D4.default()(), 14)
                        d_748_v28_ = nw120_
                        d_749_v29_: _dafny.Map
                        d_749_v29_ = _dafny.Map({d_745_v25_: d_748_v28_})
                        d_750_v30_: _dafny.Array
                        nw121_ = _dafny.Array(None, 10)
                        nw121_[int(0)] = d_748_v28_
                        nw121_[int(1)] = d_748_v28_
                        nw121_[int(2)] = d_748_v28_
                        nw121_[int(3)] = d_748_v28_
                        nw121_[int(4)] = d_748_v28_
                        nw121_[int(5)] = (d_748_v28_ if True else d_748_v28_)
                        nw121_[int(6)] = d_748_v28_
                        nw121_[int(7)] = d_748_v28_
                        nw121_[int(8)] = ((d_749_v29_)[(p0)[default__.safeIndex(63, (p0).length(0))]] if ((p0)[default__.safeIndex(63, (p0).length(0))]) in (d_749_v29_) else d_748_v28_)
                        nw121_[int(9)] = d_748_v28_
                        d_750_v30_ = nw121_
                        index115_ = default__.safeIndex(97, (d_750_v30_).length(0))
                        (d_750_v30_)[index115_] = d_748_v28_
                        index116_ = default__.safeIndex(97, (d_750_v30_).length(0))
                        (d_750_v30_)[index116_] = d_748_v28_
                    (globalState).f9 = (d_705_v0_) and (d_705_v0_)
                    pass
            pass
        d_751_v31_: int
        d_751_v31_ = 86
        d_752_v32_: _dafny.Map
        d_752_v32_ = _dafny.Map({(d_751_v31_) != (-586): True})
        d_753_v33_: str
        d_753_v33_ = _dafny.CodePoint('f')
        d_754_v34_: _dafny.Map
        d_754_v34_ = _dafny.Map({d_753_v33_: d_751_v31_})
        d_755_v35_: _dafny.Set
        d_755_v35_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_751_v31_, d_751_v31_])})
        d_752_v32_ = (d_752_v32_).set(d_705_v0_, default__.fm23(d_751_v31_, True, d_754_v34_, d_755_v35_, globalState))
        r0 = d_753_v33_
        d_756_v36_: C1
        nw122_ = C1()
        nw122_.ctor__()
        d_756_v36_ = nw122_
        d_757_i4_: int
        d_757_i4_ = 0
        with _dafny.label("3"):
            while d_705_v0_:
                with _dafny.c_label("3"):
                    if (d_757_i4_) >= (100):
                        raise _dafny.Break("3")
                    d_757_i4_ = (d_757_i4_) + (1)
                    d_758_v37_: _dafny.Set
                    d_758_v37_ = _dafny.Set({365, d_751_v31_, d_751_v31_})
                    d_759_v38_: _dafny.Seq
                    d_759_v38_ = _dafny.SeqWithoutIsStrInference([d_758_v37_, d_758_v37_, d_758_v37_])
                    d_760_v39_: _dafny.Map
                    d_760_v39_ = _dafny.Map({d_751_v31_: (d_759_v38_)[default__.safeIndex(d_751_v31_, len(d_759_v38_))]})
                    d_760_v39_ = (d_760_v39_).set(d_751_v31_, d_758_v37_)
                    d_761_v40_: _dafny.Seq
                    d_761_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncnudk"))
                    d_762_v41_: _dafny.MultiSet
                    d_762_v41_ = _dafny.MultiSet([len(d_761_v40_), d_751_v31_, d_751_v31_, d_751_v31_, d_751_v31_])
                    index117_ = default__.safeIndex(215, (p0).length(0))
                    (p0)[index117_] = (d_762_v41_).cardinality
                    index118_ = default__.safeIndex(215, (p0).length(0))
                    (p0)[index118_] = (len(_dafny.SeqWithoutIsStrInference([d_751_v31_ for d_763_i5_ in range(default__.abs(802))]))) + (default__.fm0(d_705_v0_, d_751_v31_, d_751_v31_, d_762_v41_, globalState))
                    d_764_v42_: C0
                    nw123_ = C0()
                    nw123_.ctor__(d_705_v0_, d_705_v0_)
                    d_764_v42_ = nw123_
                    d_765_v43_: _dafny.Map
                    d_765_v43_ = _dafny.Map({(p0)[default__.safeIndex(215, (p0).length(0))]: d_751_v31_})
                    r1 = (len(_dafny.SeqWithoutIsStrInference([len(d_765_v43_)]))) + ((p0)[default__.safeIndex(215, (p0).length(0))])
                    pass
            pass
        d_766_v44_: _dafny.MultiSet
        d_766_v44_ = _dafny.MultiSet([not(d_705_v0_)])
        d_767_v45_: _dafny.Seq
        d_767_v45_ = _dafny.SeqWithoutIsStrInference([d_751_v31_, d_751_v31_, d_751_v31_, d_751_v31_, 494])
        (globalState).f3 = default__.safeModuloInt((d_766_v44_).cardinality, len((d_767_v45_ if d_705_v0_ else d_767_v45_)))
        r0 = d_753_v33_
        r1 = d_751_v31_
        d_768_v46_: _dafny.Seq
        d_768_v46_ = _dafny.SeqWithoutIsStrInference([d_754_v34_, d_754_v34_, d_754_v34_, _dafny.Map({_dafny.CodePoint('v'): 701}), d_754_v34_])
        d_769_v47_: _dafny.Seq
        d_769_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymak"))
        d_770_v48_: D3
        d_770_v48_ = D3_DC10(d_752_v32_)
        r2 = (d_770_v48_ if default__.fm23(d_751_v31_, d_705_v0_, (d_768_v46_)[default__.safeIndex(len(d_769_v47_), len(d_768_v46_))], d_755_v35_, globalState) else d_770_v48_)
        return r0, r1, r2

    def m9(self, p0, globalState):
        r0: int = int(0)
        d_771_v0_: bool
        d_771_v0_ = True
        d_772_v1_: _dafny.MultiSet
        d_772_v1_ = _dafny.MultiSet([p0, p0, 917, p0])
        d_773_v2_: _dafny.Seq
        d_773_v2_ = _dafny.SeqWithoutIsStrInference([d_772_v1_])
        d_774_v3_: _dafny.Seq
        d_774_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
        d_775_v4_: D5
        d_775_v4_ = D5_DC19(D5_DC17(default__.fm29(globalState), d_771_v0_, (d_773_v2_)[default__.safeIndex((0) - (len(d_774_v3_)), len(d_773_v2_))]))
        d_776_v5_: _dafny.Seq
        d_776_v5_ = _dafny.SeqWithoutIsStrInference([d_775_v4_])
        d_777_v6_: _dafny.Seq
        d_777_v6_ = _dafny.SeqWithoutIsStrInference([d_771_v0_])
        d_778_v7_: D2
        d_778_v7_ = D2_DC7(d_777_v6_)
        d_779_v8_: _dafny.Array
        nw124_ = _dafny.Array(None, 17)
        nw124_[int(0)] = p0
        nw124_[int(1)] = p0
        nw124_[int(2)] = (0) - (p0)
        nw124_[int(3)] = p0
        nw124_[int(4)] = len((_dafny.SeqWithoutIsStrInference([d_775_v4_])) + (d_776_v5_))
        nw124_[int(5)] = p0
        nw124_[int(6)] = p0
        nw124_[int(7)] = (0) - (p0)
        nw124_[int(8)] = p0
        nw124_[int(9)] = p0
        nw124_[int(10)] = (p0) + (p0)
        nw124_[int(11)] = (len((d_778_v7_).cf9)) + (p0)
        nw124_[int(12)] = p0
        nw124_[int(13)] = p0
        nw124_[int(14)] = p0
        nw124_[int(15)] = p0
        nw124_[int(16)] = p0
        d_779_v8_ = nw124_
        index119_ = default__.safeIndex(459, (d_779_v8_).length(0))
        (d_779_v8_)[index119_] = default__.safeModuloInt(p0, p0)
        index120_ = default__.safeIndex(459, (d_779_v8_).length(0))
        (d_779_v8_)[index120_] = -989
        d_780_v9_: _dafny.MultiSet
        d_780_v9_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpyalxhrf")), d_774_v3_])
        (globalState).f9 = (d_777_v6_)[default__.safeIndex(((d_780_v9_) | (d_780_v9_)).cardinality, len(d_777_v6_))]
        d_781_v10_: C1
        nw125_ = C1()
        nw125_.ctor__()
        d_781_v10_ = nw125_
        d_782_v11_: _dafny.Array
        def lambda30_(d_783_i0_):
            return (D5_DC18(True)).cf33

        init15_ = lambda30_
        nw126_ = _dafny.Array(None, 23)
        for i0_15_ in range(nw126_.length(0)):
            nw126_[i0_15_] = init15_(i0_15_)
        d_782_v11_ = nw126_
        index121_ = default__.safeIndex(190, (d_782_v11_).length(0))
        (d_782_v11_)[index121_] = d_771_v0_
        index122_ = default__.safeIndex(512, (d_782_v11_).length(0))
        (d_782_v11_)[index122_] = not(d_771_v0_)
        index123_ = default__.safeIndex(190, (d_782_v11_).length(0))
        index124_ = default__.safeIndex(459, (d_779_v8_).length(0))
        index125_ = default__.safeIndex(512, (d_782_v11_).length(0))
        rhs101_ = d_771_v0_
        rhs102_ = not((default__.safeDivisionInt(p0, (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))])) <= ((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]))
        rhs103_ = default__.safeModuloInt((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))], default__.fm0(d_771_v0_, 177, (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))], d_772_v1_, globalState))
        rhs104_ = d_771_v0_
        lhs76_ = d_782_v11_
        lhs77_ = default__.safeIndex(190, (d_782_v11_).length(0))
        lhs78_ = d_779_v8_
        lhs79_ = default__.safeIndex(459, (d_779_v8_).length(0))
        lhs80_ = d_782_v11_
        lhs81_ = default__.safeIndex(512, (d_782_v11_).length(0))
        lhs76_[lhs77_] = rhs101_
        d_771_v0_ = rhs102_
        lhs78_[lhs79_] = rhs103_
        lhs80_[lhs81_] = rhs104_
        index126_ = default__.safeIndex(190, (d_782_v11_).length(0))
        index127_ = default__.safeIndex(459, (d_779_v8_).length(0))
        rhs105_ = (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]
        rhs106_ = 174
        lhs82_ = d_782_v11_
        lhs83_ = default__.safeIndex(190, (d_782_v11_).length(0))
        lhs84_ = d_779_v8_
        lhs85_ = default__.safeIndex(459, (d_779_v8_).length(0))
        lhs82_[lhs83_] = rhs105_
        lhs84_[lhs85_] = rhs106_
        d_784_v12_: str
        d_784_v12_ = _dafny.CodePoint('q')
        d_785_v13_: _dafny.Map
        d_785_v13_ = _dafny.Map({not(d_771_v0_): (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]})
        d_786_v14_: _dafny.Seq
        d_786_v14_ = _dafny.SeqWithoutIsStrInference([p0, len(d_785_v13_), p0, (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]])
        d_787_v15_: _dafny.Set
        d_787_v15_ = _dafny.Set({d_786_v14_})
        d_788_v16_: _dafny.Set
        d_788_v16_ = _dafny.Set({d_771_v0_, default__.fm23((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))], False, _dafny.Map({d_784_v12_: (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]}), d_787_v15_, globalState)})
        if (d_788_v16_).ispropersubset((d_788_v16_).intersection(d_788_v16_)):
            if (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]:
                d_789_v17_: _dafny.Map
                d_789_v17_ = _dafny.Map({d_771_v0_: p0})
                d_790_v18_: _dafny.Map
                d_790_v18_ = _dafny.Map({d_774_v3_: d_789_v17_})
                d_791_v19_: D3
                d_791_v19_ = D3_DC11(p0, len(d_785_v13_), d_771_v0_, (d_786_v14_)[default__.safeIndex(p0, len(d_786_v14_))])
                d_790_v18_ = (d_790_v18_).set((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "byt")) if d_771_v0_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))), default__.fm1((d_791_v19_).cf19, (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))], not(d_771_v0_), default__.fm22(p0, _dafny.SeqWithoutIsStrInference([_dafny.Map({p0: d_771_v0_}), _dafny.Map({p0: (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]})]), d_771_v0_, not((d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]), globalState), globalState))
                d_792_v20_: _dafny.MultiSet
                d_792_v20_ = _dafny.MultiSet([(d_789_v17_).set((d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))], p0), _dafny.Map({d_771_v0_: p0})])
                index128_ = default__.safeIndex(459, (d_779_v8_).length(0))
                (d_779_v8_)[index128_] = (d_792_v20_).cardinality
                (globalState).f3 = p0
                (globalState).f9 = not(d_771_v0_)
                (globalState).f3 = default__.fm0(True, 312, p0, d_772_v1_, globalState)
            elif True:
                index129_ = default__.safeIndex(190, (d_782_v11_).length(0))
                (d_782_v11_)[index129_] = (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]
                r0 = (d_786_v14_)[default__.safeIndex(726, len(d_786_v14_))]
                d_793_v21_: D1
                d_793_v21_ = D1_DC5((p0) not in (d_786_v14_), False)
                d_794_v22_: _dafny.Seq
                d_794_v22_ = _dafny.SeqWithoutIsStrInference([d_777_v6_])
                d_795_v23_: D5
                d_795_v23_ = D5_DC17(d_794_v22_, (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))], d_772_v1_)
                index130_ = default__.safeIndex(459, (d_779_v8_).length(0))
                index131_ = default__.safeIndex(190, (d_782_v11_).length(0))
                rhs107_ = (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]
                rhs108_ = d_793_v21_
                rhs109_ = (0) - ((default__.safeDivisionInt((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))], p0)) - ((0) - (((d_772_v1_).cardinality) * ((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]))))
                rhs110_ = True
                rhs111_ = not (((d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]) or (not((d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]))) or ((d_795_v23_).cf31)
                lhs86_ = globalState
                lhs87_ = d_779_v8_
                lhs88_ = default__.safeIndex(459, (d_779_v8_).length(0))
                lhs89_ = globalState
                lhs90_ = d_782_v11_
                lhs91_ = default__.safeIndex(190, (d_782_v11_).length(0))
                lhs86_.f9 = rhs107_
                d_793_v21_ = rhs108_
                lhs87_[lhs88_] = rhs109_
                lhs89_.f9 = rhs110_
                lhs90_[lhs91_] = rhs111_
                (globalState).f3 = 31
                index132_ = default__.safeIndex(190, (d_782_v11_).length(0))
                (d_782_v11_)[index132_] = (882) == ((p0) * ((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]))
            index133_ = default__.safeIndex(459, (d_779_v8_).length(0))
            (d_779_v8_)[index133_] = (p0) * (((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]) * (p0))
            if (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]:
                d_796_v24_: _dafny.Map
                d_796_v24_ = _dafny.Map({(d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]: p0})
                d_796_v24_ = ((_dafny.Map({(d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]: (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]})).set((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))], (0) - (p0))) | (d_796_v24_)
                (globalState).f9 = (p0) >= (p0)
                d_797_v25_: D4
                d_797_v25_ = D4_DC14(d_771_v0_, (446) <= (p0))
                d_797_v25_ = d_797_v25_
                d_771_v0_ = d_771_v0_
                index134_ = default__.safeIndex(459, (d_779_v8_).length(0))
                (d_779_v8_)[index134_] = (0) - (((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))] if (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))] else default__.safeDivisionInt((0) - (p0), len(d_774_v3_))))
            elif True:
                (globalState).f9 = not(not(d_771_v0_))
                d_798_v26_: _dafny.Array
                def lambda31_(d_799_v12_):
                    def lambda32_(d_800_i1_):
                        return d_799_v12_

                    return lambda32_

                init16_ = lambda31_(d_784_v12_)
                nw127_ = _dafny.Array(None, 16)
                for i0_16_ in range(nw127_.length(0)):
                    nw127_[i0_16_] = init16_(i0_16_)
                d_798_v26_ = nw127_
                index135_ = default__.safeIndex(846, (d_798_v26_).length(0))
                (d_798_v26_)[index135_] = d_784_v12_
                index136_ = default__.safeIndex(846, (d_798_v26_).length(0))
                (d_798_v26_)[index136_] = d_784_v12_
                d_801_v27_: _dafny.Map
                d_801_v27_ = _dafny.Map({(d_798_v26_)[default__.safeIndex(846, (d_798_v26_).length(0))]: 118})
                d_802_v28_: D1
                d_802_v28_ = D1_DC5(d_771_v0_, default__.fm23((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))], (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))], d_801_v27_, d_787_v15_, globalState))
                d_803_v29_: D1
                d_803_v29_ = D1_DC6(d_802_v28_)
                d_804_v30_: D1
                d_804_v30_ = D1_DC6(d_802_v28_)
                d_805_v31_: _dafny.Map
                d_805_v31_ = _dafny.Map({p0: d_804_v30_})
                d_805_v31_ = (d_805_v31_).set(p0, d_804_v30_)
                (globalState).f3 = len(_dafny.Map({default__.fm25(default__.fm0(d_771_v0_, (d_772_v1_).cardinality, (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))], d_772_v1_, globalState), (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))], d_771_v0_, _dafny.Map({default__.fm23((d_772_v1_).cardinality, d_771_v0_, d_801_v27_, d_787_v15_, globalState): (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]}), globalState): d_777_v6_}))
                (globalState).f3 = default__.safeModuloInt((0) - (-204), default__.safeDivisionInt((0) - ((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]), (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]))
            d_771_v0_ = True
            d_806_v32_: D2
            d_806_v32_ = D2_DC9((d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))], p0, d_784_v12_)
            d_807_v33_: D4
            d_807_v33_ = D4_DC13(d_774_v3_, (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))], (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))], d_782_v11_)
            index137_ = default__.safeIndex(459, (d_779_v8_).length(0))
            index138_ = default__.safeIndex(190, (d_782_v11_).length(0))
            rhs112_ = (d_806_v32_).cf14
            rhs113_ = (d_807_v33_).cf22
            rhs114_ = False
            rhs115_ = (default__.safeModuloInt(p0, p0)) + (default__.fm0((d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))], p0, p0, d_772_v1_, globalState))
            lhs92_ = d_779_v8_
            lhs93_ = default__.safeIndex(459, (d_779_v8_).length(0))
            lhs94_ = d_782_v11_
            lhs95_ = default__.safeIndex(190, (d_782_v11_).length(0))
            lhs96_ = globalState
            lhs92_[lhs93_] = rhs112_
            d_774_v3_ = rhs113_
            lhs94_[lhs95_] = rhs114_
            lhs96_.f3 = rhs115_
        elif True:
            (globalState).f3 = 835
            (globalState).f3 = (0) - ((0) - (default__.safeModuloInt(default__.safeDivisionInt((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))], (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]), (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))])))
            d_808_v34_: D4
            d_808_v34_ = D4_DC13(d_774_v3_, 452, d_771_v0_, d_782_v11_)
            d_809_v35_: _dafny.Seq
            d_809_v35_ = _dafny.SeqWithoutIsStrInference([d_808_v34_, d_808_v34_, d_808_v34_])
            d_771_v0_ = (d_809_v35_) <= (d_809_v35_)
            d_810_v36_: D2
            d_810_v36_ = D2_DC9(d_771_v0_, (0) - (p0), d_784_v12_)
            d_811_v37_: _dafny.Set
            d_811_v37_ = _dafny.Set({_dafny.MultiSet([not(True)]), _dafny.MultiSet([(d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]])})
            d_812_v38_: _dafny.Map
            d_812_v38_ = _dafny.Map({((d_810_v36_).cf15 if False else d_784_v12_): ((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]) - (len(d_811_v37_))})
            d_813_v39_: _dafny.MultiSet
            d_813_v39_ = _dafny.MultiSet([d_779_v8_])
            d_812_v38_ = (d_812_v38_).set(_dafny.CodePoint('w'), (d_813_v39_).cardinality)
            if not(((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]) == ((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))])):
                (globalState).f13 = d_788_v16_
                d_814_v40_: _dafny.Map
                d_814_v40_ = _dafny.Map({(d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]: d_784_v12_})
                d_815_v43_: _dafny.Map
                d_815_v43_ = d_814_v40_
                d_816_v44_: _dafny.Array
                nw128_ = _dafny.Array(None, 28)
                nw128_[int(0)] = d_814_v40_
                nw128_[int(1)] = (d_814_v40_).set(p0, _dafny.CodePoint('o'))
                nw128_[int(2)] = d_814_v40_
                def iife85_():
                    coll39_ = _dafny.Map()
                    compr_39_: int
                    for compr_39_ in (default__.fm30((d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))], globalState)).Elements:
                        d_817_v41_: int = compr_39_
                        if (d_817_v41_) in (default__.fm30((d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))], globalState)):
                            coll39_[default__.safeDivisionInt(d_817_v41_, len(d_777_v6_))] = d_784_v12_
                    return _dafny.Map(coll39_)
                nw128_[int(3)] = iife85_()
                
                nw128_[int(4)] = d_814_v40_
                def iife86_():
                    coll40_ = _dafny.Map()
                    compr_40_: int
                    for compr_40_ in _dafny.IntegerRange(-39, -511):
                        d_818_v42_: int = compr_40_
                        if ((-39) <= (d_818_v42_)) and ((d_818_v42_) < (-511)):
                            coll40_[default__.safeModuloInt(d_818_v42_, (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))])] = _dafny.CodePoint('v')
                    return _dafny.Map(coll40_)
                nw128_[int(5)] = iife86_()
                
                nw128_[int(6)] = d_814_v40_
                nw128_[int(7)] = default__.fm31(d_771_v0_, d_771_v0_, globalState)
                nw128_[int(8)] = d_814_v40_
                nw128_[int(9)] = d_814_v40_
                nw128_[int(10)] = d_814_v40_
                nw128_[int(11)] = d_814_v40_
                nw128_[int(12)] = d_814_v40_
                nw128_[int(13)] = d_814_v40_
                nw128_[int(14)] = d_814_v40_
                nw128_[int(15)] = _dafny.Map({p0: d_784_v12_})
                nw128_[int(16)] = d_814_v40_
                nw128_[int(17)] = d_814_v40_
                nw128_[int(18)] = (d_815_v43_)
                nw128_[int(19)] = d_814_v40_
                nw128_[int(20)] = d_814_v40_
                nw128_[int(21)] = default__.fm31(True, d_771_v0_, globalState)
                nw128_[int(22)] = d_814_v40_
                nw128_[int(23)] = d_814_v40_
                nw128_[int(24)] = d_814_v40_
                nw128_[int(25)] = _dafny.Map({(d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]: d_784_v12_})
                nw128_[int(26)] = d_814_v40_
                nw128_[int(27)] = d_814_v40_
                d_816_v44_ = nw128_
                d_819_v45_: D0
                d_819_v45_ = D0_DC1(d_771_v0_)
                d_820_v46_: D1
                d_821_v47_: int
                d_822_v48_: bool
                d_823_v49_: int
                out43_: D1
                out44_: int
                out45_: bool
                out46_: int
                out43_, out44_, out45_, out46_ = (d_781_v10_).m10(d_816_v44_, d_779_v8_, ((d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]) and ((d_819_v45_).cf1), globalState)
                d_820_v46_ = out43_
                d_821_v47_ = out44_
                d_822_v48_ = out45_
                d_823_v49_ = out46_
                d_824_v50_: _dafny.MultiSet
                d_824_v50_ = _dafny.MultiSet([d_784_v12_, d_784_v12_, d_784_v12_, d_784_v12_, d_784_v12_])
                d_825_v51_: D4
                d_825_v51_ = D4_DC14((len(default__.fm29(globalState))) != ((d_824_v50_).cardinality), (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))])
                d_826_v52_: _dafny.MultiSet
                d_826_v52_ = _dafny.MultiSet([d_822_v48_, (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))], d_771_v0_, (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))], d_822_v48_])
                pat_let_tv21_ = d_826_v52_
                pat_let_tv22_ = d_826_v52_
                def iife87_(_pat_let23_0):
                    def iife88_(d_827_dt__update__tmp_h0_):
                        def iife89_(_pat_let24_0):
                            def iife90_(d_828_dt__update_hcf26_h0_):
                                return D4_DC14(d_828_dt__update_hcf26_h0_, (d_827_dt__update__tmp_h0_).cf27)
                            return iife90_(_pat_let24_0)
                        return iife89_((pat_let_tv21_).ispropersubset(pat_let_tv22_))
                    return iife88_(_pat_let23_0)
                d_825_v51_ = iife87_(D4_DC14(((self.f22)[d_823_v49_] if (d_823_v49_) in (self.f22) else d_771_v0_), d_771_v0_))
                (globalState).f3 = d_821_v47_
                d_829_v53_: C0
                nw129_ = C0()
                nw129_.ctor__(d_771_v0_, d_771_v0_)
                d_829_v53_ = nw129_
            elif True:
                index139_ = default__.safeIndex(459, (d_779_v8_).length(0))
                (d_779_v8_)[index139_] = 288
                (globalState).f9 = (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]
                (globalState).f9 = d_771_v0_
                d_830_v54_: _dafny.Map
                d_830_v54_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([d_784_v12_ for d_831_i2_ in range(default__.abs(-811))])})
                d_774_v3_ = (d_774_v3_) + (((d_830_v54_)[(0) - ((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))])] if ((0) - ((d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))])) in (d_830_v54_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tecdvlbj"))))
                d_832_v55_: _dafny.Map
                d_832_v55_ = _dafny.Map({327: d_777_v6_})
                d_833_v56_: _dafny.MultiSet
                d_833_v56_ = _dafny.MultiSet([d_771_v0_, d_771_v0_, d_771_v0_])
                rhs116_ = len(_dafny.Set({(len(d_832_v55_)) <= (970), (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))], d_771_v0_, (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]}))
                rhs117_ = ((d_780_v9_).cardinality if ((d_785_v13_)[(d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]] if ((d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]) in (d_785_v13_) else False) else ((d_833_v56_)[d_771_v0_] if (d_771_v0_) in (d_833_v56_) else default__.fm0((d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))], p0, p0, (d_772_v1_).set(554, default__.abs(len(d_786_v14_))), globalState)))
                r0 = rhs116_
                r0 = rhs117_
        d_834_v58_: _dafny.Set
        d_834_v58_ = _dafny.Set({d_784_v12_})
        def iife91_():
            coll41_ = _dafny.Map()
            compr_41_: str
            for compr_41_ in (d_834_v58_).Elements:
                d_835_v57_: str = compr_41_
                if (d_835_v57_) in (d_834_v58_):
                    coll41_[d_835_v57_] = (d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))]
            return _dafny.Map(coll41_)
        r0 = default__.fm0(not(d_771_v0_), default__.fm0((d_782_v11_)[default__.safeIndex(190, (d_782_v11_).length(0))], (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))], (d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))], _dafny.MultiSet([(d_779_v8_)[default__.safeIndex(459, (d_779_v8_).length(0))]]), globalState), len(iife91_()
        ), (d_772_v1_ if False else d_772_v1_), globalState)
        return r0


class C3(T0):
    def  __init__(self):
        self._f14: bool = False
        self._f15: _dafny.MultiSet = _dafny.MultiSet({})
        self._f20: int = int(0)
        self._f21: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    @property
    def f15(self):
        return self._f15
    def ctor__(self, f20, f21, f14, f15):
        (self)._f20 = f20
        (self)._f21 = f21
        (self).f14 = f14
        (self)._f15 = f15

    def fm2(self, p0, p1, globalState):
        if self.f14:
            return self.f14
        elif True:
            return (-49) != (len((self).f21))

    def fm3(self, p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dowcmd"))) + ((((self).f21).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_836_i0_ in range(default__.abs(952))])), len((self).f21)), _dafny.CodePoint('y'))) + ((self).f21))).set(default__.safeIndex((self).f20, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dowcmd"))) + ((((self).f21).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_837_i0_ in range(default__.abs(952))])), len((self).f21)), _dafny.CodePoint('y'))) + ((self).f21)))), _dafny.CodePoint('a'))

    def m0(self, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        r3: bool = False
        d_838_v0_: _dafny.MultiSet
        d_838_v0_ = _dafny.MultiSet([_dafny.Map({self.f14: self.f14})])
        d_839_v1_: _dafny.Map
        d_839_v1_ = _dafny.Map({self.f14: self.f14})
        d_840_v2_: _dafny.Map
        d_840_v2_ = _dafny.Map({((d_839_v1_)[not(self.f14)] if (not(self.f14)) in (d_839_v1_) else self.f14): self.f14})
        d_841_v3_: _dafny.MultiSet
        d_841_v3_ = _dafny.MultiSet([(self).f20])
        d_838_v0_ = (d_838_v0_) - ((_dafny.MultiSet([(d_840_v2_).set(False, self.f14), d_840_v2_])).set(d_839_v1_, default__.abs(default__.fm0(self.f14, (self).f20, (self).f20, d_841_v3_, globalState))))
        d_842_v4_: D2
        d_842_v4_ = D2_DC7(_dafny.SeqWithoutIsStrInference([self.f14]))
        d_843_v5_: _dafny.MultiSet
        d_843_v5_ = _dafny.MultiSet([d_842_v4_])
        d_843_v5_ = d_843_v5_
        (globalState).f3 = default__.safeDivisionInt(713, ((self).f20) * (115))
        d_844_v6_: _dafny.Seq
        d_844_v6_ = _dafny.SeqWithoutIsStrInference([self.f14, self.f14])
        hi2_ = len(d_844_v6_)
        for d_845_i0_ in range((self).f20, hi2_):
            d_846_v7_: _dafny.Array
            nw130_ = _dafny.Array(int(0), 1)
            d_846_v7_ = nw130_
            index140_ = default__.safeIndex(290, (d_846_v7_).length(0))
            (d_846_v7_)[index140_] = d_845_i0_
            d_847_v8_: _dafny.Array
            nw131_ = _dafny.Array(False, 11)
            d_847_v8_ = nw131_
            d_848_v9_: D1
            d_848_v9_ = D1_DC4(d_845_i0_, len((self).fm3((self).f20, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fo"))), self.f14, (self).f20, globalState)))
            index141_ = default__.safeIndex(290, (d_846_v7_).length(0))
            rhs118_ = d_847_v8_
            rhs119_ = not(self.f14)
            rhs120_ = (((self).f15)[(self).fm2(self.f14, self.f14, globalState)] if ((self).fm2(self.f14, self.f14, globalState)) in ((self).f15) else len(_dafny.SeqWithoutIsStrInference([(self).f21 for d_849_i1_ in range(default__.abs(875))])))
            rhs121_ = (d_848_v9_).cf4
            lhs97_ = globalState
            lhs98_ = globalState
            lhs99_ = d_846_v7_
            lhs100_ = default__.safeIndex(290, (d_846_v7_).length(0))
            lhs101_ = globalState
            lhs97_.f2 = rhs118_
            lhs98_.f9 = rhs119_
            lhs99_[lhs100_] = rhs120_
            lhs101_.f3 = rhs121_
            d_850_v10_: _dafny.Array
            nw132_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
            d_850_v10_ = nw132_
            index142_ = default__.safeIndex(632, (d_850_v10_).length(0))
            (d_850_v10_)[index142_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_851_i2_ in range(default__.abs(-993))])
            d_852_v11_: _dafny.Array
            def lambda33_(d_853_v9_, d_854_v7_):
                def lambda34_(d_855_i3_):
                    def iife92_(_pat_let25_0):
                        def iife93_(d_856_dt__update__tmp_h0_):
                            def iife94_(_pat_let26_0):
                                def iife95_(d_857_dt__update_hcf4_h0_):
                                    return D1_DC4(d_857_dt__update_hcf4_h0_, (d_856_dt__update__tmp_h0_).cf5)
                                return iife95_(_pat_let26_0)
                            return iife94_((0) - ((d_854_v7_)[default__.safeIndex(290, (d_854_v7_).length(0))]))
                        return iife93_(_pat_let25_0)
                    return iife92_(d_853_v9_)

                return lambda34_

            init17_ = lambda33_(d_848_v9_, d_846_v7_)
            nw133_ = _dafny.Array(None, 8)
            for i0_17_ in range(nw133_.length(0)):
                nw133_[i0_17_] = init17_(i0_17_)
            d_852_v11_ = nw133_
            index143_ = default__.safeIndex(817, (d_852_v11_).length(0))
            (d_852_v11_)[index143_] = d_848_v9_
            index144_ = default__.safeIndex(290, (d_846_v7_).length(0))
            index145_ = default__.safeIndex(632, (d_850_v10_).length(0))
            index146_ = default__.safeIndex(817, (d_852_v11_).length(0))
            rhs122_ = d_845_i0_
            rhs123_ = (self).f21
            rhs124_ = ((False if True else self.f14)) or (self.f14)
            rhs125_ = d_845_i0_
            rhs126_ = d_848_v9_
            lhs102_ = d_846_v7_
            lhs103_ = default__.safeIndex(290, (d_846_v7_).length(0))
            lhs104_ = d_850_v10_
            lhs105_ = default__.safeIndex(632, (d_850_v10_).length(0))
            lhs106_ = globalState
            lhs107_ = d_852_v11_
            lhs108_ = default__.safeIndex(817, (d_852_v11_).length(0))
            lhs102_[lhs103_] = rhs122_
            lhs104_[lhs105_] = rhs123_
            r1 = rhs124_
            lhs106_.f3 = rhs125_
            lhs107_[lhs108_] = rhs126_
            index147_ = default__.safeIndex(290, (d_846_v7_).length(0))
            (d_846_v7_)[index147_] = 578
            if True:
                d_858_v12_: _dafny.Map
                d_858_v12_ = _dafny.Map({(self).f20: self.f14})
                d_858_v12_ = (_dafny.Map({d_845_i0_: True})) | ((d_858_v12_) | (d_858_v12_))
                d_859_v13_: D0
                d_859_v13_ = D0_DC0(not(self.f14))
                d_860_v14_: _dafny.Array
                nw134_ = _dafny.Array(None, 13)
                nw134_[int(0)] = d_859_v13_
                nw134_[int(1)] = d_859_v13_
                nw134_[int(2)] = default__.fm20(-575, (d_846_v7_)[default__.safeIndex(290, (d_846_v7_).length(0))], (d_846_v7_)[default__.safeIndex(290, (d_846_v7_).length(0))], globalState)
                nw134_[int(3)] = (d_859_v13_ if self.f14 else d_859_v13_)
                nw134_[int(4)] = default__.fm20(600, (d_846_v7_)[default__.safeIndex(290, (d_846_v7_).length(0))], (d_846_v7_)[default__.safeIndex(290, (d_846_v7_).length(0))], globalState)
                nw134_[int(5)] = d_859_v13_
                nw134_[int(6)] = (d_859_v13_ if self.f14 else d_859_v13_)
                nw134_[int(7)] = d_859_v13_
                nw134_[int(8)] = d_859_v13_
                nw134_[int(9)] = (d_859_v13_ if self.f14 else d_859_v13_)
                nw134_[int(10)] = d_859_v13_
                nw134_[int(11)] = d_859_v13_
                nw134_[int(12)] = d_859_v13_
                d_860_v14_ = nw134_
                index148_ = default__.safeIndex(768, (d_860_v14_).length(0))
                (d_860_v14_)[index148_] = d_859_v13_
                index149_ = default__.safeIndex(768, (d_860_v14_).length(0))
                (d_860_v14_)[index149_] = d_859_v13_
                d_861_v15_: str
                d_861_v15_ = _dafny.CodePoint('e')
                d_862_v16_: _dafny.MultiSet
                d_862_v16_ = _dafny.MultiSet([d_861_v15_, _dafny.CodePoint('p')])
                d_863_v17_: _dafny.Seq
                d_863_v17_ = _dafny.SeqWithoutIsStrInference([(self).f20])
                rhs127_ = (self).f20
                rhs128_ = d_862_v16_
                rhs129_ = d_845_i0_
                rhs130_ = (self).f20
                rhs131_ = (d_863_v17_)[default__.safeIndex((self).f20, len(d_863_v17_))]
                lhs109_ = globalState
                lhs110_ = globalState
                lhs111_ = globalState
                lhs112_ = globalState
                lhs109_.f3 = rhs127_
                d_862_v16_ = rhs128_
                lhs110_.f3 = rhs129_
                lhs111_.f3 = rhs130_
                lhs112_.f3 = rhs131_
                d_864_v18_: _dafny.Map
                d_864_v18_ = _dafny.Map({self.f14: _dafny.MultiSet([self.f14, self.f14, self.f14, self.f14, False])})
                index150_ = default__.safeIndex(322, (d_847_v8_).length(0))
                (d_847_v8_)[index150_] = ((((d_864_v18_)[not(self.f14)] if (not(self.f14)) in (d_864_v18_) else (self).f15)).cardinality) != (d_845_i0_)
                index151_ = default__.safeIndex(322, (d_847_v8_).length(0))
                (d_847_v8_)[index151_] = True
                (globalState).f3 = ((self).f20) - ((d_846_v7_)[default__.safeIndex(290, (d_846_v7_).length(0))])
            elif True:
                d_865_v19_: C2
                nw135_ = C2()
                nw135_.ctor__(_dafny.Map({(self).f20: self.f14}))
                d_865_v19_ = nw135_
                d_866_v20_: _dafny.Map
                d_866_v20_ = _dafny.Map({(self.f14 if self.f14 else self.f14): (self).f20})
                d_866_v20_ = (d_866_v20_).set(self.f14, ((d_846_v7_)[default__.safeIndex(290, (d_846_v7_).length(0))]) * (d_845_i0_))
                index152_ = default__.safeIndex(290, (d_846_v7_).length(0))
                (d_846_v7_)[index152_] = default__.safeDivisionInt(default__.safeModuloInt(657, (d_846_v7_)[default__.safeIndex(290, (d_846_v7_).length(0))]), 631)
                d_867_v21_: _dafny.Map
                d_867_v21_ = _dafny.Map({(d_846_v7_)[default__.safeIndex(290, (d_846_v7_).length(0))]: d_845_i0_})
                d_868_v22_: _dafny.Map
                d_868_v22_ = _dafny.Map({d_846_v7_: d_867_v21_})
                d_869_v23_: _dafny.Seq
                d_869_v23_ = _dafny.SeqWithoutIsStrInference([(d_846_v7_)[default__.safeIndex(290, (d_846_v7_).length(0))]])
                d_870_v24_: _dafny.Array
                nw136_ = _dafny.Array(None, 25)
                nw136_[int(0)] = d_868_v22_
                nw136_[int(1)] = (_dafny.Map({d_846_v7_: d_867_v21_})) | (d_868_v22_)
                nw136_[int(2)] = d_868_v22_
                nw136_[int(3)] = _dafny.Map({d_846_v7_: d_867_v21_})
                nw136_[int(4)] = d_868_v22_
                nw136_[int(5)] = _dafny.Map({d_846_v7_: d_867_v21_})
                nw136_[int(6)] = d_868_v22_
                nw136_[int(7)] = (d_868_v22_) | (d_868_v22_)
                nw136_[int(8)] = d_868_v22_
                nw136_[int(9)] = d_868_v22_
                nw136_[int(10)] = d_868_v22_
                nw136_[int(11)] = d_868_v22_
                nw136_[int(12)] = d_868_v22_
                nw136_[int(13)] = d_868_v22_
                nw136_[int(14)] = d_868_v22_
                nw136_[int(15)] = d_868_v22_
                nw136_[int(16)] = d_868_v22_
                nw136_[int(17)] = (d_868_v22_) | (d_868_v22_)
                nw136_[int(18)] = (d_868_v22_).set(d_846_v7_, d_867_v21_)
                nw136_[int(19)] = (_dafny.Map({d_846_v7_: d_867_v21_})) | (d_868_v22_)
                nw136_[int(20)] = (d_868_v22_) | (d_868_v22_)
                nw136_[int(21)] = (d_868_v22_) | (d_868_v22_)
                nw136_[int(22)] = _dafny.Map({d_846_v7_: d_867_v21_})
                nw136_[int(23)] = (d_868_v22_) | (d_868_v22_)
                nw136_[int(24)] = _dafny.Map({d_846_v7_: _dafny.Map({135: default__.fm0(self.f14, d_845_i0_, (0) - ((d_869_v23_)[default__.safeIndex((d_846_v7_)[default__.safeIndex(290, (d_846_v7_).length(0))], len(d_869_v23_))]), _dafny.MultiSet(d_869_v23_), globalState)})})
                d_870_v24_ = nw136_
                index153_ = default__.safeIndex(349, (d_870_v24_).length(0))
                (d_870_v24_)[index153_] = (d_868_v22_) | (d_868_v22_)
                index154_ = default__.safeIndex(349, (d_870_v24_).length(0))
                (d_870_v24_)[index154_] = (d_868_v22_).set(d_846_v7_, d_867_v21_)
                (self).f14 = (self.f14 if (d_844_v6_)[default__.safeIndex(d_845_i0_, len(d_844_v6_))] else self.f14)
        if not(self.f14):
            d_871_v25_: _dafny.Seq
            d_871_v25_ = _dafny.SeqWithoutIsStrInference([301, ((self).f20) * ((0) - ((self).f20)), (self).f20])
            d_871_v25_ = d_871_v25_
            d_872_v26_: D3
            d_872_v26_ = D3_DC11((self).f20, ((self).f20) * ((self).f20), self.f14, (326) + (len(default__.fm32((self).f21, d_841_v3_, globalState))))
            d_872_v26_ = d_872_v26_
            d_873_v27_: _dafny.Array
            nw137_ = _dafny.Array(None, 8)
            nw137_[int(0)] = 898
            nw137_[int(1)] = (self).f20
            nw137_[int(2)] = (self).f20
            nw137_[int(3)] = (self).f20
            nw137_[int(4)] = (self).f20
            nw137_[int(5)] = (self).f20
            nw137_[int(6)] = len(d_844_v6_)
            nw137_[int(7)] = (self).f20
            d_873_v27_ = nw137_
            d_874_v28_: _dafny.Seq
            d_874_v28_ = _dafny.SeqWithoutIsStrInference([d_873_v27_])
            d_874_v28_ = (d_874_v28_) + (d_874_v28_)
            r3 = (((self).f20) > ((0) - ((self).f20))) and (self.f14)
            (globalState).f9 = self.f14
        elif True:
            d_875_v29_: _dafny.Array
            nw138_ = _dafny.Array(_dafny.Array(None, 0), 29)
            d_875_v29_ = nw138_
            d_876_v30_: _dafny.Map
            d_876_v30_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnwodjwai"))) + ((self).f21): d_875_v29_})
            d_876_v30_ = (d_876_v30_).set((self).f21, d_875_v29_)
            d_877_v31_: C0
            nw139_ = C0()
            nw139_.ctor__(self.f14, self.f14)
            d_877_v31_ = nw139_
            d_878_v32_: _dafny.Array
            nw140_ = _dafny.Array(False, 24)
            d_878_v32_ = nw140_
            index155_ = default__.safeIndex(829, (d_878_v32_).length(0))
            (d_878_v32_)[index155_] = self.f14
            index156_ = default__.safeIndex(829, (d_878_v32_).length(0))
            (d_878_v32_)[index156_] = (d_877_v31_).f23
            d_879_v33_: _dafny.Array
            nw141_ = _dafny.Array(_dafny.CodePoint('D'), 7)
            d_879_v33_ = nw141_
            d_880_v34_: str
            d_880_v34_ = _dafny.CodePoint('p')
            d_881_v35_: _dafny.Map
            d_881_v35_ = _dafny.Map({d_879_v33_: ((self).f21).set(default__.safeIndex((self).f20, len((self).f21)), d_880_v34_)})
            d_881_v35_ = d_881_v35_
            (globalState).f3 = 255
        (globalState).f3 = (self).f20
        r0 = ((d_841_v3_)[(self).f20] if ((self).f20) in (d_841_v3_) else (self).f20)
        r1 = self.f14
        d_882_v36_: _dafny.Map
        d_882_v36_ = _dafny.Map({(self).f20: self.f14})
        d_883_v37_: _dafny.Map
        d_883_v37_ = _dafny.Map({d_882_v36_: (self).f20})
        r2 = (D7_DC21(d_883_v37_)).cf36
        r3 = True
        return r0, r1, r2, r3

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        (self).f14 = self.f14
        d_884_v0_: _dafny.Array
        nw142_ = _dafny.Array(False, 14)
        d_884_v0_ = nw142_
        index157_ = default__.safeIndex(887, (d_884_v0_).length(0))
        (d_884_v0_)[index157_] = self.f14
        index158_ = default__.safeIndex(887, (d_884_v0_).length(0))
        (d_884_v0_)[index158_] = self.f14
        d_885_v1_: C0
        nw143_ = C0()
        nw143_.ctor__(self.f14, (self.f14) == (self.f14))
        d_885_v1_ = nw143_
        d_886_v2_: _dafny.Array
        nw144_ = _dafny.Array(int(0), 20)
        d_886_v2_ = nw144_
        index159_ = default__.safeIndex(232, (d_886_v2_).length(0))
        (d_886_v2_)[index159_] = len(default__.fm27(len(p0), (d_885_v1_).f24, globalState))
        index160_ = default__.safeIndex(232, (d_886_v2_).length(0))
        (d_886_v2_)[index160_] = default__.safeDivisionInt(p1, default__.fm0((d_884_v0_)[default__.safeIndex(887, (d_884_v0_).length(0))], (self).f20, p1, default__.fm30(False, globalState), globalState))
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_886_v2_).length(0)):
            d_887_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_887_i0_)) and ((d_887_i0_) < ((d_886_v2_).length(0)))):
                (d_886_v2_)[(d_887_i0_)] = (d_887_i0_) + ((D3_DC11(len(_dafny.Map({(self).f20: (0) - ((self).f20)})), len(_dafny.Map({p0: _dafny.Map({p1: D2_DC9(not((d_885_v1_).f24), p1, _dafny.CodePoint('w'))})})), (d_885_v1_).f24, (self).f20)).cf20)
        if False:
            r2 = ((p1) >= (len((self).f21)) if (d_884_v0_)[default__.safeIndex(887, (d_884_v0_).length(0))] else self.f14)
            d_888_v3_: _dafny.Array
            nw145_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
            d_888_v3_ = nw145_
            index161_ = default__.safeIndex(644, (d_888_v3_).length(0))
            (d_888_v3_)[index161_] = (self).f21
            index162_ = default__.safeIndex(644, (d_888_v3_).length(0))
            (d_888_v3_)[index162_] = (self).f21
            d_889_v4_: _dafny.Map
            d_889_v4_ = _dafny.Map({(self).f21: (d_885_v1_).f23})
            d_890_v5_: _dafny.Seq
            d_890_v5_ = _dafny.SeqWithoutIsStrInference([(d_885_v1_).f23])
            (globalState).f6 = _dafny.Set({True, (((d_889_v4_)[(d_888_v3_)[default__.safeIndex(644, (d_888_v3_).length(0))]] if ((d_888_v3_)[default__.safeIndex(644, (d_888_v3_).length(0))]) in (d_889_v4_) else (d_890_v5_)[default__.safeIndex(p1, len(d_890_v5_))]) if (d_885_v1_).f23 else (d_885_v1_).f24), not(not(False))})
            (self).f14 = not((self).fm2(self.f14, not(((d_885_v1_).f23) == ((d_885_v1_).f24)), globalState))
            d_891_v6_: str
            d_891_v6_ = _dafny.CodePoint('f')
            index163_ = default__.safeIndex(887, (d_884_v0_).length(0))
            rhs132_ = ((d_885_v1_).f24) and ((d_885_v1_).f24)
            rhs133_ = d_891_v6_
            rhs134_ = (d_885_v1_).f24
            lhs113_ = self
            lhs114_ = d_884_v0_
            lhs115_ = default__.safeIndex(887, (d_884_v0_).length(0))
            lhs113_.f14 = rhs132_
            d_891_v6_ = rhs133_
            lhs114_[lhs115_] = rhs134_
        elif True:
            d_892_v7_: _dafny.Array
            nw146_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_892_v7_ = nw146_
            index164_ = default__.safeIndex(283, (d_892_v7_).length(0))
            (d_892_v7_)[index164_] = d_886_v2_
            index165_ = default__.safeIndex(283, (d_892_v7_).length(0))
            nw147_ = _dafny.Array(int(0), 8)
            (d_892_v7_)[index165_] = nw147_
            (globalState).f9 = (d_885_v1_).f24
            d_893_v8_: _dafny.Map
            d_893_v8_ = _dafny.Map({(d_884_v0_)[default__.safeIndex(887, (d_884_v0_).length(0))]: not((d_884_v0_)[default__.safeIndex(887, (d_884_v0_).length(0))])})
            d_894_v9_: D3
            d_894_v9_ = D3_DC10((d_893_v8_).set((d_884_v0_)[default__.safeIndex(887, (d_884_v0_).length(0))], (d_885_v1_).f23))
            d_895_v10_: _dafny.Map
            d_895_v10_ = _dafny.Map({d_894_v9_: ((d_886_v2_)[default__.safeIndex(232, (d_886_v2_).length(0))]) < ((0) - (len(_dafny.SeqWithoutIsStrInference([self.f14, (d_885_v1_).f24]))))})
            d_895_v10_ = (d_895_v10_).set(d_894_v9_, (d_885_v1_).f23)
            d_896_v11_: _dafny.Set
            d_896_v11_ = _dafny.Set({p1})
            d_896_v11_ = (d_896_v11_) - ((d_896_v11_ if (d_884_v0_)[default__.safeIndex(887, (d_884_v0_).length(0))] else d_896_v11_))
            index166_ = default__.safeIndex(887, (d_884_v0_).length(0))
            (d_884_v0_)[index166_] = (d_885_v1_).f23
        d_897_v12_: _dafny.Seq
        d_897_v12_ = _dafny.SeqWithoutIsStrInference([self.f14, False])
        d_898_v13_: str
        d_898_v13_ = _dafny.CodePoint('h')
        d_899_v14_: _dafny.Map
        d_899_v14_ = _dafny.Map({len(d_897_v12_): d_898_v13_})
        r0 = ((((self).f15) - ((self).f15)).cardinality) - (len(d_899_v14_))
        d_900_v15_: _dafny.MultiSet
        d_900_v15_ = _dafny.MultiSet([(d_886_v2_)[default__.safeIndex(232, (d_886_v2_).length(0))], p1])
        r1 = ((len(_dafny.Set({p1}))) * ((d_900_v15_).cardinality)) - (default__.safeDivisionInt((self).f20, 546))
        r2 = not(not (not(False)) or ((p1) > (len(default__.fm27((0) - (len(d_897_v12_)), (d_884_v0_)[default__.safeIndex(887, (d_884_v0_).length(0))], globalState)))))
        return r0, r1, r2

    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21

class C4(T0):
    def  __init__(self):
        self._f14: bool = False
        self._f15: _dafny.MultiSet = _dafny.MultiSet({})
        self._f19: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    @property
    def f15(self):
        return self._f15
    def ctor__(self, f19, f14, f15):
        (self)._f19 = f19
        (self).f14 = f14
        (self)._f15 = f15

    def fm2(self, p0, p1, globalState):
        return (self).f19

    def fm3(self, p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_901_i0_ in range(default__.abs(-13))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlscb")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fu"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_902_i1_ in range(default__.abs(866))])))

    def fm17(self, p0, p1, p2, p3, globalState):
        return (((self).f15).intersection((self).f15)).cardinality

    def fm18(self, p0, p1, p2, p3, globalState):
        return ((D3_DC10(_dafny.Map({self.f14: (self).f19}))).cf16) | (_dafny.Map({(self).f19: self.f14}))

    def m0(self, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        r3: bool = False
        if not((self).f19):
            d_903_v0_: _dafny.Array
            def lambda35_(d_904_i0_):
                return False

            init18_ = lambda35_
            nw148_ = _dafny.Array(None, 10)
            for i0_18_ in range(nw148_.length(0)):
                nw148_[i0_18_] = init18_(i0_18_)
            d_903_v0_ = nw148_
            index167_ = default__.safeIndex(891, (d_903_v0_).length(0))
            (d_903_v0_)[index167_] = not ((self).f19) or ((self).f19)
            d_905_v1_: int
            d_905_v1_ = -689
            d_906_v2_: _dafny.Seq
            d_906_v2_ = _dafny.SeqWithoutIsStrInference([self.f14])
            d_907_v3_: str
            d_907_v3_ = _dafny.CodePoint('v')
            d_908_v4_: _dafny.Set
            d_908_v4_ = _dafny.Set({d_907_v3_})
            d_909_v5_: _dafny.Map
            d_909_v5_ = _dafny.Map({d_907_v3_: (self).f19})
            index168_ = default__.safeIndex(891, (d_903_v0_).length(0))
            def iife96_():
                coll42_ = _dafny.Set()
                compr_42_: str
                for compr_42_ in (d_909_v5_).keys.Elements:
                    d_910_v6_: str = compr_42_
                    if (d_910_v6_) in (d_909_v5_):
                        coll42_ = coll42_.union(_dafny.Set([d_910_v6_]))
                return _dafny.Set(coll42_)
            (d_903_v0_)[index168_] = ((default__.fm19(d_905_v1_, (d_906_v2_)[default__.safeIndex((0) - (d_905_v1_), len(d_906_v2_))], globalState)) - (d_908_v4_)).isdisjoint(iife96_()
            )
            (globalState).f9 = self.f14
            d_911_v7_: _dafny.Map
            d_911_v7_ = _dafny.Map({(d_903_v0_)[default__.safeIndex(891, (d_903_v0_).length(0))]: self.f14})
            d_912_v8_: _dafny.Seq
            d_912_v8_ = _dafny.SeqWithoutIsStrInference([d_911_v7_])
            (globalState).f9 = (_dafny.Map({self.f14: self.f14})) in (d_912_v8_)
            d_913_v9_: _dafny.Seq
            d_913_v9_ = _dafny.SeqWithoutIsStrInference([(self).f15])
            d_914_v10_: _dafny.Array
            def lambda36_(d_915_v1_):
                def lambda37_(d_916_i1_):
                    return (d_916_i1_) * (d_915_v1_)

                return lambda37_

            init19_ = lambda36_(d_905_v1_)
            nw149_ = _dafny.Array(None, 13)
            for i0_19_ in range(nw149_.length(0)):
                nw149_[i0_19_] = init19_(i0_19_)
            d_914_v10_ = nw149_
            d_917_v11_: _dafny.Set
            d_917_v11_ = _dafny.Set({d_914_v10_, d_914_v10_, d_914_v10_})
            d_918_v12_: _dafny.MultiSet
            d_918_v12_ = _dafny.MultiSet([d_905_v1_])
            (self).m7(d_913_v9_, (_dafny.Set({d_914_v10_})).isdisjoint(d_917_v11_), (772 if (self).f19 else d_905_v1_), (d_905_v1_) >= (default__.fm0(self.f14, -955, d_905_v1_, d_918_v12_, globalState)), globalState)
            d_919_v13_: _dafny.Map
            d_919_v13_ = _dafny.Map({d_905_v1_: d_907_v3_})
            d_920_v14_: D3
            d_920_v14_ = D3_DC11(d_905_v1_, d_905_v1_, (self).f19, len(d_919_v13_))
            (globalState).f3 = default__.fm0((d_905_v1_) >= ((((self).f15).set((self).f19, default__.abs(d_905_v1_))).cardinality), d_905_v1_, (d_920_v14_).cf20, d_918_v12_, globalState)
        elif True:
            d_921_v15_: D1
            d_921_v15_ = D1_DC5(self.f14, self.f14)
            d_922_v16_: _dafny.Array
            nw150_ = _dafny.Array(None, 18)
            nw150_[int(0)] = d_921_v15_
            nw150_[int(1)] = d_921_v15_
            nw150_[int(2)] = d_921_v15_
            nw150_[int(3)] = d_921_v15_
            nw150_[int(4)] = d_921_v15_
            nw150_[int(5)] = d_921_v15_
            nw150_[int(6)] = d_921_v15_
            nw150_[int(7)] = d_921_v15_
            nw150_[int(8)] = D1_DC5(False, True)
            nw150_[int(9)] = d_921_v15_
            nw150_[int(10)] = d_921_v15_
            nw150_[int(11)] = d_921_v15_
            nw150_[int(12)] = d_921_v15_
            nw150_[int(13)] = d_921_v15_
            nw150_[int(14)] = d_921_v15_
            nw150_[int(15)] = d_921_v15_
            nw150_[int(16)] = D1_DC5(self.f14, self.f14)
            nw150_[int(17)] = d_921_v15_
            d_922_v16_ = nw150_
            d_923_v17_: _dafny.Seq
            d_923_v17_ = _dafny.SeqWithoutIsStrInference([d_922_v16_, d_922_v16_, d_922_v16_])
            d_924_v18_: int
            d_924_v18_ = -618
            d_925_v19_: _dafny.Array
            nw151_ = _dafny.Array(None, 12)
            nw151_[int(0)] = d_923_v17_
            nw151_[int(1)] = (d_923_v17_) + (d_923_v17_)
            nw151_[int(2)] = ((d_923_v17_) + (d_923_v17_)).set(default__.safeIndex(d_924_v18_, len((d_923_v17_) + (d_923_v17_))), d_922_v16_)
            nw151_[int(3)] = d_923_v17_
            nw151_[int(4)] = d_923_v17_
            nw151_[int(5)] = d_923_v17_
            nw151_[int(6)] = d_923_v17_
            nw151_[int(7)] = _dafny.SeqWithoutIsStrInference([d_922_v16_, d_922_v16_])
            nw151_[int(8)] = d_923_v17_
            nw151_[int(9)] = (d_923_v17_) + (d_923_v17_)
            nw151_[int(10)] = d_923_v17_
            nw151_[int(11)] = d_923_v17_
            d_925_v19_ = nw151_
            index169_ = default__.safeIndex(930, (d_925_v19_).length(0))
            (d_925_v19_)[index169_] = d_923_v17_
            index170_ = default__.safeIndex(930, (d_925_v19_).length(0))
            (d_925_v19_)[index170_] = ((d_923_v17_) + (d_923_v17_)) + ((d_923_v17_) + (d_923_v17_))
            (globalState).f9 = (d_924_v18_) <= ((0) - (d_924_v18_))
            d_926_v20_: _dafny.Array
            nw152_ = _dafny.Array(False, 26)
            d_926_v20_ = nw152_
            d_927_v21_: _dafny.Array
            nw153_ = _dafny.Array(None, 8)
            nw153_[int(0)] = d_926_v20_
            nw153_[int(1)] = d_926_v20_
            nw153_[int(2)] = d_926_v20_
            nw153_[int(3)] = d_926_v20_
            nw153_[int(4)] = d_926_v20_
            nw153_[int(5)] = d_926_v20_
            nw153_[int(6)] = (d_926_v20_ if True else d_926_v20_)
            nw153_[int(7)] = d_926_v20_
            d_927_v21_ = nw153_
            index171_ = default__.safeIndex(371, (d_927_v21_).length(0))
            (d_927_v21_)[index171_] = d_926_v20_
            index172_ = default__.safeIndex(371, (d_927_v21_).length(0))
            nw154_ = _dafny.Array(False, 7)
            (d_927_v21_)[index172_] = nw154_
            d_928_v22_: _dafny.Seq
            d_928_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkgn"))
            d_929_v23_: T0
            nw155_ = C3()
            nw155_.ctor__(d_924_v18_, d_928_v22_, not((self).f19), default__.fm33((self).f19, globalState))
            d_929_v23_ = nw155_
            d_930_v24_: _dafny.Map
            d_930_v24_ = _dafny.Map({(d_924_v18_) + (len(_dafny.SeqWithoutIsStrInference([d_924_v18_]))): d_929_v23_})
            d_931_v25_: _dafny.Map
            d_931_v25_ = _dafny.Map({d_924_v18_: -353})
            d_932_v26_: _dafny.MultiSet
            d_932_v26_ = _dafny.MultiSet([len(d_931_v25_)])
            d_933_v27_: D1
            d_933_v27_ = D1_DC4(d_924_v18_, default__.fm0(d_929_v23_.f14, (0) - (len(d_928_v22_)), 548, d_932_v26_, globalState))
            d_930_v24_ = (d_930_v24_).set((d_933_v27_).cf5, d_929_v23_)
            d_924_v18_ = default__.safeModuloInt(d_924_v18_, (d_924_v18_) * (d_924_v18_))
        d_934_v28_: int
        d_934_v28_ = -847
        d_935_v29_: _dafny.Seq
        d_935_v29_ = _dafny.SeqWithoutIsStrInference([d_934_v28_])
        d_936_v30_: _dafny.MultiSet
        d_936_v30_ = _dafny.MultiSet([(d_935_v29_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_937_i2_ in range(default__.abs(-609))])), len(d_935_v29_))]])
        d_938_v31_: _dafny.Seq
        d_938_v31_ = _dafny.SeqWithoutIsStrInference([d_934_v28_, d_934_v28_, (d_936_v30_).cardinality])
        d_935_v29_ = d_938_v31_
        d_939_v32_: _dafny.Seq
        d_939_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsfh"))
        (self).f14 = not ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eduiltxa"))) <= (d_939_v32_)) or (self.f14)
        d_940_v33_: _dafny.Map
        d_940_v33_ = _dafny.Map({d_934_v28_: _dafny.Map({self.f14: len(_dafny.SeqWithoutIsStrInference([self.f14, False, self.f14]))})})
        d_941_v34_: _dafny.Map
        d_941_v34_ = _dafny.Map({d_940_v33_: (self).f19})
        d_942_v35_: _dafny.Seq
        d_942_v35_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_943_i4_ in range(default__.abs(-201))])])
        d_944_v36_: _dafny.Array
        nw156_ = _dafny.Array(None, 22)
        nw156_[int(0)] = ((d_941_v34_)[d_940_v33_] if (d_940_v33_) in (d_941_v34_) else self.f14)
        nw156_[int(1)] = (377) <= (d_934_v28_)
        nw156_[int(2)] = (self).f19
        nw156_[int(3)] = (d_934_v28_) <= ((0) - (d_934_v28_))
        nw156_[int(4)] = (self).f19
        nw156_[int(5)] = (self).fm2((self).f19, self.f14, globalState)
        nw156_[int(6)] = not((self).f19)
        nw156_[int(7)] = (self).f19
        nw156_[int(8)] = self.f14
        nw156_[int(9)] = (d_939_v32_) != ((d_942_v35_)[default__.safeIndex(d_934_v28_, len(d_942_v35_))])
        nw156_[int(10)] = (self).fm2((self).f19, (self).f19, globalState)
        nw156_[int(11)] = not((self).f19)
        nw156_[int(12)] = ((self).f19) == (self.f14)
        nw156_[int(13)] = not (self.f14) or (self.f14)
        nw156_[int(14)] = False
        nw156_[int(15)] = (True) or (False)
        nw156_[int(16)] = (self.f14 if self.f14 else (self).f19)
        nw156_[int(17)] = self.f14
        nw156_[int(18)] = False
        nw156_[int(19)] = self.f14
        nw156_[int(20)] = (_dafny.MultiSet([d_939_v32_])).isdisjoint(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpfkgsqvw"))]))
        nw156_[int(21)] = not(self.f14)
        d_944_v36_ = nw156_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_944_v36_).length(0)):
            d_945_i3_: int = guard_loop_2_
            if (True) and (((0) <= (d_945_i3_)) and ((d_945_i3_) < ((d_944_v36_).length(0)))):
                (d_944_v36_)[(d_945_i3_)] = (_dafny.Set({self.f14, not(True), self.f14, self.f14})).issubset((_dafny.Set({False, self.f14})) - (_dafny.Set({True, self.f14})))
        d_946_v37_: C1
        nw157_ = C1()
        nw157_.ctor__()
        d_946_v37_ = nw157_
        d_947_v38_: _dafny.Map
        d_947_v38_ = _dafny.Map({(self).f19: (self).f19})
        d_948_v39_: _dafny.Set
        d_948_v39_ = _dafny.Set({d_934_v28_, 608, len(d_947_v38_), d_934_v28_})
        r1 = not((d_948_v39_).issubset(d_948_v39_))
        r0 = len(d_939_v32_)
        d_949_v40_: str
        d_949_v40_ = _dafny.CodePoint('s')
        d_950_v41_: _dafny.Map
        d_950_v41_ = _dafny.Map({d_949_v40_: d_934_v28_})
        d_951_v42_: _dafny.Set
        d_951_v42_ = _dafny.Set({d_935_v29_, d_938_v31_})
        r1 = default__.fm23((d_934_v28_) * (d_934_v28_), (self).f19, (d_950_v41_) | (_dafny.Map({d_949_v40_: d_934_v28_})), d_951_v42_, globalState)
        d_952_v43_: _dafny.Map
        d_952_v43_ = _dafny.Map({d_934_v28_: self.f14})
        r2 = _dafny.Map({(d_952_v43_) | (_dafny.Map({d_934_v28_: (self).fm2(self.f14, self.f14, globalState)})): d_934_v28_})
        r3 = (d_934_v28_) == (d_934_v28_)
        return r0, r1, r2, r3

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        r1 = p1
        if not (not(not ((self).f19) or (self.f14))) or ((self).f19):
            d_953_v0_: _dafny.Array
            def lambda38_(d_954_i0_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mpehub"))

            init20_ = lambda38_
            nw158_ = _dafny.Array(None, 3)
            for i0_20_ in range(nw158_.length(0)):
                nw158_[i0_20_] = init20_(i0_20_)
            d_953_v0_ = nw158_
            d_955_v2_: _dafny.Map
            d_955_v2_ = _dafny.Map({p1: p1})
            def iife97_():
                coll43_ = _dafny.Map()
                compr_43_: int
                for compr_43_ in (d_955_v2_).keys.Elements:
                    d_956_v1_: int = compr_43_
                    if (d_956_v1_) in (d_955_v2_):
                        coll43_[(d_956_v1_) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D2_DC9(self.f14, p1, _dafny.CodePoint('q'))).cf13]))).cardinality)] = self.f14
                return _dafny.Map(coll43_)
            rhs135_ = len(iife97_()
            )
            rhs136_ = d_953_v0_
            r1 = rhs135_
            d_953_v0_ = rhs136_
            r1 = p1
            r2 = True
            d_957_v3_: _dafny.Seq
            d_957_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iycqg"))
            d_958_v4_: T0
            nw159_ = C3()
            nw159_.ctor__(p1, d_957_v3_, self.f14, (self).f15)
            d_958_v4_ = nw159_
            d_959_v5_: _dafny.Array
            def lambda39_(d_960_p1_):
                def lambda40_(d_961_i1_):
                    return (d_961_i1_) * (len(_dafny.SeqWithoutIsStrInference([d_960_p1_])))

                return lambda40_

            init21_ = lambda39_(p1)
            nw160_ = _dafny.Array(None, 27)
            for i0_21_ in range(nw160_.length(0)):
                nw160_[i0_21_] = init21_(i0_21_)
            d_959_v5_ = nw160_
            d_962_v6_: _dafny.Map
            d_962_v6_ = _dafny.Map({d_959_v5_: d_958_v4_})
            if (d_958_v4_) not in (_dafny.Set({d_958_v4_, ((d_962_v6_)[d_959_v5_] if (d_959_v5_) in (d_962_v6_) else d_958_v4_), d_958_v4_, d_958_v4_})):
                d_963_v7_: str
                d_963_v7_ = _dafny.CodePoint('v')
                d_964_v8_: C3
                nw161_ = C3()
                nw161_.ctor__(p1, (d_957_v3_).set(default__.safeIndex(p1, len(d_957_v3_)), d_963_v7_), self.f14, (d_958_v4_).f15)
                d_964_v8_ = nw161_
                index173_ = default__.safeIndex(899, (d_953_v0_).length(0))
                (d_953_v0_)[index173_] = (d_964_v8_).f21
                index174_ = default__.safeIndex(899, (d_953_v0_).length(0))
                (d_953_v0_)[index174_] = ((d_964_v8_).f21 if not((self).f19) else (d_964_v8_).f21)
                d_965_v9_: _dafny.Set
                d_965_v9_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([len((d_964_v8_).f21) for d_966_i2_ in range(default__.abs(749))])})
                d_967_v10_: _dafny.Seq
                d_967_v10_ = _dafny.SeqWithoutIsStrInference([(self).f19])
                d_968_v11_: _dafny.Seq
                d_968_v11_ = _dafny.SeqWithoutIsStrInference([(d_967_v10_)[default__.safeIndex(p1, len(d_967_v10_))]])
                d_969_v12_: _dafny.Set
                d_969_v12_ = _dafny.Set({(d_964_v8_).f21})
                r2 = ((d_968_v11_) == ((d_967_v10_).set(default__.safeIndex((d_964_v8_).f20, len(d_967_v10_)), (self).f19)) if not(default__.fm23((d_964_v8_).f20, d_958_v4_.f14, _dafny.Map({d_963_v7_: p1}), d_965_v9_, globalState)) else (d_969_v12_).ispropersubset(default__.fm34(globalState)))
                d_970_v13_: _dafny.Array
                nw162_ = _dafny.Array(False, 8)
                d_970_v13_ = nw162_
                (globalState).f2 = d_970_v13_
                d_971_v14_: _dafny.Array
                def lambda41_(d_972_v4_):
                    def lambda42_(d_973_i3_):
                        return (_dafny.SeqWithoutIsStrInference([d_972_v4_.f14])).set(default__.safeIndex(len(_dafny.Set({self.f14})), len(_dafny.SeqWithoutIsStrInference([d_972_v4_.f14]))), True)

                    return lambda42_

                init22_ = lambda41_(d_958_v4_)
                nw163_ = _dafny.Array(None, 23)
                for i0_22_ in range(nw163_.length(0)):
                    nw163_[i0_22_] = init22_(i0_22_)
                d_971_v14_ = nw163_
                index175_ = default__.safeIndex(603, (d_971_v14_).length(0))
                (d_971_v14_)[index175_] = _dafny.SeqWithoutIsStrInference([True, d_958_v4_.f14, d_958_v4_.f14])
                index176_ = default__.safeIndex(7, (d_959_v5_).length(0))
                (d_959_v5_)[index176_] = ((d_964_v8_).f20) * (p1)
                index177_ = default__.safeIndex(603, (d_971_v14_).length(0))
                index178_ = default__.safeIndex(7, (d_959_v5_).length(0))
                rhs137_ = d_967_v10_
                rhs138_ = d_963_v7_
                rhs139_ = default__.safeDivisionInt((0) - ((len(d_957_v3_)) + (p1)), (((d_958_v4_).f15) - ((d_958_v4_).f15)).cardinality)
                rhs140_ = d_958_v4_.f14
                lhs116_ = d_971_v14_
                lhs117_ = default__.safeIndex(603, (d_971_v14_).length(0))
                lhs118_ = d_959_v5_
                lhs119_ = default__.safeIndex(7, (d_959_v5_).length(0))
                lhs120_ = self
                lhs116_[lhs117_] = rhs137_
                d_963_v7_ = rhs138_
                lhs118_[lhs119_] = rhs139_
                lhs120_.f14 = rhs140_
            elif True:
                r0 = (p1) * ((0) - (p1))
                r0 = (0) - ((621) - ((p1 if d_958_v4_.f14 else (0) - (p1))))
                d_974_v15_: _dafny.MultiSet
                d_974_v15_ = _dafny.MultiSet([(self).f19, (p1) != (p1), False, False, True])
                d_974_v15_ = (((_dafny.MultiSet([d_958_v4_.f14])).set(d_958_v4_.f14, default__.abs(p1))) - (_dafny.MultiSet([d_958_v4_.f14, self.f14, (self).f19, (self).f19]))) - (((d_958_v4_).f15) | (d_974_v15_))
                index179_ = default__.safeIndex(591, (d_959_v5_).length(0))
                (d_959_v5_)[index179_] = p1
                index180_ = default__.safeIndex(591, (d_959_v5_).length(0))
                (d_959_v5_)[index180_] = p1
                r2 = d_958_v4_.f14
            d_975_v16_: _dafny.Array
            nw164_ = _dafny.Array(False, 21)
            d_975_v16_ = nw164_
            d_976_v17_: D5
            d_976_v17_ = D5_DC18(d_958_v4_.f14)
            index181_ = default__.safeIndex(334, (d_975_v16_).length(0))
            (d_975_v16_)[index181_] = (d_976_v17_).cf33
            index182_ = default__.safeIndex(334, (d_975_v16_).length(0))
            (d_975_v16_)[index182_] = not((not((self).f19)) and (self.f14))
        elif True:
            (globalState).f9 = (p1) > (819)
            d_977_v18_: str
            d_977_v18_ = _dafny.CodePoint('c')
            d_978_v19_: _dafny.Map
            d_978_v19_ = _dafny.Map({d_977_v18_: p1})
            d_979_v20_: _dafny.Seq
            d_979_v20_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_980_v21_: _dafny.MultiSet
            d_980_v21_ = _dafny.MultiSet([-711])
            d_981_v22_: _dafny.Map
            d_981_v22_ = _dafny.Map({p1: p1})
            if not ((default__.fm0(True, p1, default__.fm0((self).f19, 877, p1, d_980_v21_, globalState), d_980_v21_, globalState)) not in (d_981_v22_)) or (default__.fm23(p1, self.f14, d_978_v19_, _dafny.Set({d_979_v20_}), globalState)):
                (globalState).f3 = p1
                d_982_v23_: C1
                nw165_ = C1()
                nw165_.ctor__()
                d_982_v23_ = nw165_
                d_982_v23_ = d_982_v23_
                d_983_v24_: _dafny.Seq
                d_983_v24_ = _dafny.SeqWithoutIsStrInference([self.f14, (self).f19, self.f14, self.f14])
                d_984_v25_: _dafny.MultiSet
                d_984_v25_ = _dafny.MultiSet([d_983_v24_])
                d_985_v26_: _dafny.Seq
                d_985_v26_ = _dafny.SeqWithoutIsStrInference([d_980_v21_])
                r1 = (self).fm17(p1, self.f14, (_dafny.MultiSet([d_983_v24_])) - (d_984_v25_), len(d_985_v26_), globalState)
                (self).m6(globalState)
                d_986_v27_: _dafny.Array
                nw166_ = _dafny.Array(D1.default()(), 6)
                d_986_v27_ = nw166_
                d_987_v28_: _dafny.Seq
                d_987_v28_ = _dafny.SeqWithoutIsStrInference([d_986_v27_, d_986_v27_, d_986_v27_, d_986_v27_])
                r1 = (0) - (len(d_987_v28_))
            elif True:
                d_988_v29_: _dafny.Seq
                d_988_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erx"))
                d_989_v30_: _dafny.Seq
                d_989_v30_ = _dafny.SeqWithoutIsStrInference([(self).f19, self.f14])
                d_990_v31_: C3
                nw167_ = C3()
                nw167_.ctor__(p1, d_988_v29_, not(self.f14), _dafny.MultiSet(d_989_v30_))
                d_990_v31_ = nw167_
                rhs141_ = (-958) != (p1)
                rhs142_ = _dafny.SeqWithoutIsStrInference([p1])
                rhs143_ = default__.fm0((d_990_v31_).fm2(self.f14, (self).f19, globalState), p1, len(default__.fm35(p1, p1, globalState)), d_980_v21_, globalState)
                rhs144_ = (self).f19
                lhs121_ = globalState
                lhs122_ = globalState
                r2 = rhs141_
                d_979_v20_ = rhs142_
                lhs121_.f3 = rhs143_
                lhs122_.f9 = rhs144_
                r0 = p1
                (globalState).f3 = (d_990_v31_).f20
                d_991_v32_: C3
                nw168_ = C3()
                nw168_.ctor__((d_990_v31_).f20, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oek")), self.f14, (self).f15)
                d_991_v32_ = nw168_
            d_977_v18_ = d_977_v18_
            d_992_v33_: D0
            d_992_v33_ = D0_DC0((self).f19)
            d_993_v34_: D0
            d_993_v34_ = D0_DC2(d_992_v33_)
            d_994_v35_: D0
            d_994_v35_ = D0_DC2(d_993_v34_)
            source12_ = d_994_v35_
            if source12_.is_DC1:
                d_995___mcc_h0_ = source12_.cf1
                d_996_cf1_ = d_995___mcc_h0_
                d_997_v36_: _dafny.Seq
                d_997_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukldksby"))
                d_998_v37_: _dafny.Map
                d_998_v37_ = _dafny.Map({(self).f19: p1})
                d_977_v18_ = default__.fm25(p1, default__.safeDivisionInt(p1, p1), (d_977_v18_) in (d_997_v36_), d_998_v37_, globalState)
                (globalState).f3 = p1
                (globalState).f9 = d_996_cf1_
                d_999_v38_: _dafny.Array
                def lambda43_(d_1000_cf1_):
                    def lambda44_(d_1001_i4_):
                        return D3_DC10(_dafny.Map({(D1_DC5(self.f14, (self).f19)).cf6: d_1000_cf1_}))

                    return lambda44_

                init23_ = lambda43_(d_996_cf1_)
                nw169_ = _dafny.Array(None, 8)
                for i0_23_ in range(nw169_.length(0)):
                    nw169_[i0_23_] = init23_(i0_23_)
                d_999_v38_ = nw169_
                d_999_v38_ = d_999_v38_
            elif source12_.is_DC0:
                d_1002___mcc_h1_ = source12_.cf0
                d_1003_cf0_ = d_1002___mcc_h1_
                (globalState).f3 = default__.safeDivisionInt(p1, ((d_981_v22_)[p1] if (p1) in (d_981_v22_) else p1))
                d_1004_v39_: _dafny.Map
                d_1004_v39_ = _dafny.Map({d_1003_cf0_: p1})
                d_1005_v40_: _dafny.Seq
                d_1005_v40_ = _dafny.SeqWithoutIsStrInference([d_1004_v39_])
                d_1006_v41_: _dafny.Seq
                d_1006_v41_ = _dafny.SeqWithoutIsStrInference([default__.fm36(d_980_v21_, p1, d_1003_cf0_, (self).f19, globalState)])
                d_1005_v40_ = ((d_1006_v41_)[default__.safeIndex((0) - (p1), len(d_1006_v41_))]) + (d_1005_v40_)
                d_1007_v42_: _dafny.Array
                nw170_ = _dafny.Array(False, 12)
                d_1007_v42_ = nw170_
                index183_ = default__.safeIndex(853, (d_1007_v42_).length(0))
                (d_1007_v42_)[index183_] = self.f14
                d_1008_v43_: _dafny.Seq
                d_1008_v43_ = _dafny.SeqWithoutIsStrInference([d_1003_cf0_])
                index184_ = default__.safeIndex(853, (d_1007_v42_).length(0))
                (d_1007_v42_)[index184_] = (not ((self).f19) or ((self).f19)) and ((d_1008_v43_)[default__.safeIndex(p1, len(d_1008_v43_))])
                (self).m6(globalState)
            elif True:
                d_1009___mcc_h2_ = source12_.cf2
                d_1010_cf2_ = d_1009___mcc_h2_
                d_1011_v44_: _dafny.Array
                nw171_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_1011_v44_ = nw171_
                d_1012_v45_: _dafny.Array
                def lambda45_(d_1013_i5_):
                    return default__.safeDivisionInt(d_1013_i5_, 588)

                init24_ = lambda45_
                nw172_ = _dafny.Array(None, 25)
                for i0_24_ in range(nw172_.length(0)):
                    nw172_[i0_24_] = init24_(i0_24_)
                d_1012_v45_ = nw172_
                index185_ = default__.safeIndex(265, (d_1011_v44_).length(0))
                (d_1011_v44_)[index185_] = d_1012_v45_
                index186_ = default__.safeIndex(265, (d_1011_v44_).length(0))
                (d_1011_v44_)[index186_] = d_1012_v45_
                d_1014_v46_: D7
                d_1014_v46_ = D7_DC23(self.f14, len((d_979_v20_).set(default__.safeIndex((0) - (p1), len(d_979_v20_)), p1)), self.f14)
                d_1015_v47_: _dafny.Map
                d_1015_v47_ = _dafny.Map({d_1014_v46_: (self).fm2((self).f19, (self).f19, globalState)})
                d_1015_v47_ = (d_1015_v47_).set(d_1014_v46_, (self).f19)
                d_1016_v48_: _dafny.Seq
                d_1016_v48_ = _dafny.SeqWithoutIsStrInference([self.f14])
                d_1017_v49_: _dafny.Array
                def lambda46_(d_1018_i6_):
                    return (self).f19

                init25_ = lambda46_
                nw173_ = _dafny.Array(None, 1)
                for i0_25_ in range(nw173_.length(0)):
                    nw173_[i0_25_] = init25_(i0_25_)
                d_1017_v49_ = nw173_
                d_1019_v50_: _dafny.Set
                d_1019_v50_ = _dafny.Set({(self).f19})
                d_1020_v51_: D8
                d_1020_v51_ = D8_DC24(d_1019_v50_)
                d_1021_v52_: C1
                nw174_ = C1()
                nw174_.ctor__()
                d_1021_v52_ = nw174_
                d_1022_v53_: _dafny.Map
                d_1022_v53_ = _dafny.Map({p1: d_1021_v52_})
                d_1023_v54_: _dafny.Map
                d_1023_v54_ = _dafny.Map({d_1022_v53_: _dafny.Set({(self).fm2(True, (self).f19, globalState), not(False)})})
                d_1024_v55_: _dafny.MultiSet
                d_1024_v55_ = _dafny.MultiSet([d_1016_v48_, d_1016_v48_])
                d_1025_v56_: _dafny.Array
                nw175_ = _dafny.Array(None, 25)
                nw175_[int(0)] = (self).f19
                nw175_[int(1)] = (self).f19
                nw175_[int(2)] = self.f14
                nw175_[int(3)] = not(((self).f19 if (self).f19 else (self).f19))
                nw175_[int(4)] = self.f14
                nw175_[int(5)] = self.f14
                nw175_[int(6)] = self.f14
                nw175_[int(7)] = not((self).f19)
                nw175_[int(8)] = True
                nw175_[int(9)] = (p1) == (len(d_1016_v48_))
                nw175_[int(10)] = (d_1017_v49_) not in (_dafny.Map({d_1017_v49_: d_977_v18_}))
                nw175_[int(11)] = not((self).f19)
                nw175_[int(12)] = not(self.f14)
                nw175_[int(13)] = (((d_981_v22_)[p1] if (p1) in (d_981_v22_) else p1)) <= (len(p0))
                nw175_[int(14)] = ((self).f19 if (self).f19 else (self).f19)
                nw175_[int(15)] = (self).f19
                nw175_[int(16)] = (self).f19
                nw175_[int(17)] = self.f14
                nw175_[int(18)] = ((d_1020_v51_).cf45).isdisjoint(((d_1023_v54_)[d_1022_v53_] if (d_1022_v53_) in (d_1023_v54_) else d_1019_v50_))
                nw175_[int(19)] = True
                nw175_[int(20)] = not(((self).f19 if (self).f19 else (self).fm2((self).f19, self.f14, globalState)))
                nw175_[int(21)] = not ((self).f19) or (self.f14)
                nw175_[int(22)] = (p1) == ((self).fm17(p1, (self).f19, d_1024_v55_, 29, globalState))
                nw175_[int(23)] = self.f14
                nw175_[int(24)] = (self).f19
                d_1025_v56_ = nw175_
                index187_ = default__.safeIndex(870, (d_1017_v49_).length(0))
                (d_1017_v49_)[index187_] = (self).f19
                d_1026_v57_: _dafny.Array
                nw176_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_1026_v57_ = nw176_
                index188_ = default__.safeIndex(25, (d_1026_v57_).length(0))
                (d_1026_v57_)[index188_] = d_1017_v49_
                d_1027_v58_: _dafny.Map
                d_1027_v58_ = _dafny.Map({self.f14: d_1017_v49_})
                index189_ = default__.safeIndex(870, (d_1017_v49_).length(0))
                index190_ = default__.safeIndex(25, (d_1026_v57_).length(0))
                rhs145_ = not(not(not(self.f14)))
                rhs146_ = ((d_1027_v58_)[(self).f19] if ((self).f19) in (d_1027_v58_) else d_1025_v56_)
                lhs123_ = d_1017_v49_
                lhs124_ = default__.safeIndex(870, (d_1017_v49_).length(0))
                lhs125_ = d_1026_v57_
                lhs126_ = default__.safeIndex(25, (d_1026_v57_).length(0))
                lhs123_[lhs124_] = rhs145_
                lhs125_[lhs126_] = rhs146_
                d_1028_v59_: _dafny.Map
                d_1028_v59_ = _dafny.Map({-263: False})
                arr2_ = (d_1011_v44_)[default__.safeIndex(265, (d_1011_v44_).length(0))]
                index191_ = default__.safeIndex(932, ((d_1011_v44_)[default__.safeIndex(265, (d_1011_v44_).length(0))]).length(0))
                arr2_[index191_] = ((d_980_v21_)[p1] if (p1) in (d_980_v21_) else len(d_1028_v59_))
                d_1029_v60_: _dafny.Map
                d_1029_v60_ = _dafny.Map({(d_1017_v49_)[default__.safeIndex(870, (d_1017_v49_).length(0))]: len(_dafny.SeqWithoutIsStrInference([p1 for d_1030_i7_ in range(default__.abs(258))]))})
                arr3_ = (d_1011_v44_)[default__.safeIndex(265, (d_1011_v44_).length(0))]
                index192_ = default__.safeIndex(932, ((d_1011_v44_)[default__.safeIndex(265, (d_1011_v44_).length(0))]).length(0))
                rhs147_ = ((d_1029_v60_)[(self).f19] if ((self).f19) in (d_1029_v60_) else p1)
                rhs148_ = (p1) - (p1)
                lhs127_ = globalState
                lhs128_ = (d_1011_v44_)[default__.safeIndex(265, (d_1011_v44_).length(0))]
                lhs129_ = default__.safeIndex(932, ((d_1011_v44_)[default__.safeIndex(265, (d_1011_v44_).length(0))]).length(0))
                lhs127_.f3 = rhs147_
                lhs128_[lhs129_] = rhs148_
            d_1031_v61_: _dafny.Array
            def lambda47_(d_1032_p1_):
                def lambda48_(d_1033_i8_):
                    return (d_1033_i8_) + (d_1032_p1_)

                return lambda48_

            init26_ = lambda47_(p1)
            nw177_ = _dafny.Array(None, 6)
            for i0_26_ in range(nw177_.length(0)):
                nw177_[i0_26_] = init26_(i0_26_)
            d_1031_v61_ = nw177_
            index193_ = default__.safeIndex(327, (d_1031_v61_).length(0))
            (d_1031_v61_)[index193_] = p1
            d_1034_v62_: _dafny.Seq
            d_1034_v62_ = _dafny.SeqWithoutIsStrInference([(self).fm2(self.f14, False, globalState)])
            index194_ = default__.safeIndex(327, (d_1031_v61_).length(0))
            rhs149_ = p1
            rhs150_ = p1
            rhs151_ = (p1) * (len((d_1034_v62_) + (d_1034_v62_)))
            rhs152_ = (self).fm2(self.f14, (self).f19, globalState)
            rhs153_ = (len((d_1034_v62_) + (d_1034_v62_)) if (self).fm2(self.f14, self.f14, globalState) else p1)
            lhs130_ = d_1031_v61_
            lhs131_ = default__.safeIndex(327, (d_1031_v61_).length(0))
            lhs132_ = globalState
            lhs133_ = globalState
            r0 = rhs149_
            lhs130_[lhs131_] = rhs150_
            lhs132_.f3 = rhs151_
            r2 = rhs152_
            lhs133_.f3 = rhs153_
        d_1035_v63_: _dafny.Seq
        d_1035_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iep"))
        d_1035_v63_ = d_1035_v63_
        if (self).f19:
            (globalState).f9 = (self).f19
            d_1036_v64_: D4
            d_1036_v64_ = D4_DC14(False, (self).f19)
            d_1037_v65_: _dafny.Array
            nw178_ = _dafny.Array(int(0), 5)
            d_1037_v65_ = nw178_
            d_1038_v66_: _dafny.Map
            d_1038_v66_ = _dafny.Map({d_1037_v65_: not((self).f19)})
            d_1039_v67_: D5
            d_1039_v67_ = D5_DC16(d_1037_v65_)
            d_1040_v68_: _dafny.Set
            d_1040_v68_ = _dafny.Set({False})
            d_1041_v69_: _dafny.Array
            nw179_ = _dafny.Array(None, 22)
            nw179_[int(0)] = not((d_1036_v64_).cf26)
            nw179_[int(1)] = (self).f19
            nw179_[int(2)] = False
            nw179_[int(3)] = (self).f19
            nw179_[int(4)] = (self.f14 if True else (self).f19)
            nw179_[int(5)] = (self).fm2(self.f14, (self).f19, globalState)
            nw179_[int(6)] = self.f14
            nw179_[int(7)] = self.f14
            nw179_[int(8)] = (self).f19
            nw179_[int(9)] = self.f14
            nw179_[int(10)] = (self).fm2(self.f14, self.f14, globalState)
            nw179_[int(11)] = ((self).f19 if (self).f19 else (self).f19)
            nw179_[int(12)] = ((self).f19) == (((d_1038_v66_)[(d_1039_v67_).cf29] if ((d_1039_v67_).cf29) in (d_1038_v66_) else self.f14))
            nw179_[int(13)] = not(self.f14)
            nw179_[int(14)] = (self).f19
            nw179_[int(15)] = (self).f19
            nw179_[int(16)] = (_dafny.Set({not((self).f19), self.f14, (self).f19, self.f14})).issubset(d_1040_v68_)
            nw179_[int(17)] = (self).f19
            nw179_[int(18)] = (self).fm2(self.f14, (self).f19, globalState)
            nw179_[int(19)] = self.f14
            nw179_[int(20)] = (p1) == (p1)
            nw179_[int(21)] = not((self).f19)
            d_1041_v69_ = nw179_
            index195_ = default__.safeIndex(949, (d_1041_v69_).length(0))
            (d_1041_v69_)[index195_] = self.f14
            index196_ = default__.safeIndex(949, (d_1041_v69_).length(0))
            (d_1041_v69_)[index196_] = self.f14
            d_1042_v70_: _dafny.Map
            d_1042_v70_ = _dafny.Map({self.f14: d_1041_v69_})
            d_1042_v70_ = (d_1042_v70_).set((self).f19, d_1041_v69_)
            d_1043_v71_: _dafny.Seq
            d_1043_v71_ = _dafny.SeqWithoutIsStrInference([(self).f15])
            (self).m7(d_1043_v71_, (d_1041_v69_)[default__.safeIndex(949, (d_1041_v69_).length(0))], p1, (d_1041_v69_)[default__.safeIndex(949, (d_1041_v69_).length(0))], globalState)
            r2 = (p1) == (p1)
        elif True:
            (globalState).f3 = default__.safeDivisionInt(-458, 715)
            d_1044_v72_: _dafny.Set
            d_1044_v72_ = _dafny.Set({True})
            if (self.f14) not in (d_1044_v72_):
                d_1045_v73_: _dafny.Array
                nw180_ = _dafny.Array(_dafny.Set({}), 8)
                d_1045_v73_ = nw180_
                d_1045_v73_ = d_1045_v73_
                d_1046_v74_: _dafny.Seq
                d_1046_v74_ = _dafny.SeqWithoutIsStrInference([self.f14, (self).f19])
                (globalState).f9 = (d_1046_v74_)[default__.safeIndex(p1, len(d_1046_v74_))]
                d_1047_v75_: D0
                d_1047_v75_ = D0_DC1(not(self.f14))
                d_1048_v76_: _dafny.Map
                d_1048_v76_ = _dafny.Map({500: _dafny.Set({(self).f19, self.f14})})
                d_1049_v77_: _dafny.Array
                nw181_ = _dafny.Array(None, 25)
                nw181_[int(0)] = d_1047_v75_
                nw181_[int(1)] = d_1047_v75_
                nw181_[int(2)] = d_1047_v75_
                nw181_[int(3)] = d_1047_v75_
                def iife98_(_pat_let27_0):
                    def iife99_(d_1050_dt__update__tmp_h0_):
                        def iife100_(_pat_let28_0):
                            def iife101_(d_1051_dt__update_hcf1_h0_):
                                return D0_DC1(d_1051_dt__update_hcf1_h0_)
                            return iife101_(_pat_let28_0)
                        return iife100_((self).f19)
                    return iife99_(_pat_let27_0)
                nw181_[int(4)] = iife98_(d_1047_v75_)
                nw181_[int(5)] = d_1047_v75_
                nw181_[int(6)] = d_1047_v75_
                nw181_[int(7)] = d_1047_v75_
                nw181_[int(8)] = d_1047_v75_
                nw181_[int(9)] = default__.fm37(p0, d_1048_v76_, (self).f19, (D0_DC1(not(self.f14))).cf1, globalState)
                nw181_[int(10)] = d_1047_v75_
                nw181_[int(11)] = d_1047_v75_
                def iife102_(_pat_let29_0):
                    def iife103_(d_1052_dt__update__tmp_h1_):
                        def iife104_(_pat_let30_0):
                            def iife105_(d_1053_dt__update_hcf1_h1_):
                                return D0_DC1(d_1053_dt__update_hcf1_h1_)
                            return iife105_(_pat_let30_0)
                        return iife104_(self.f14)
                    return iife103_(_pat_let29_0)
                nw181_[int(12)] = iife102_(default__.fm37(p0, d_1048_v76_, self.f14, (self).f19, globalState))
                nw181_[int(13)] = d_1047_v75_
                nw181_[int(14)] = d_1047_v75_
                nw181_[int(15)] = D0_DC1(self.f14)
                nw181_[int(16)] = d_1047_v75_
                nw181_[int(17)] = d_1047_v75_
                nw181_[int(18)] = d_1047_v75_
                nw181_[int(19)] = d_1047_v75_
                nw181_[int(20)] = d_1047_v75_
                nw181_[int(21)] = d_1047_v75_
                nw181_[int(22)] = d_1047_v75_
                nw181_[int(23)] = d_1047_v75_
                def iife106_(_pat_let31_0):
                    def iife107_(d_1054_dt__update__tmp_h2_):
                        def iife108_(_pat_let32_0):
                            def iife109_(d_1055_dt__update_hcf1_h2_):
                                return D0_DC1(d_1055_dt__update_hcf1_h2_)
                            return iife109_(_pat_let32_0)
                        return iife108_(self.f14)
                    return iife107_(_pat_let31_0)
                nw181_[int(24)] = iife106_(d_1047_v75_)
                d_1049_v77_ = nw181_
                d_1049_v77_ = d_1049_v77_
                (globalState).f9 = False
                d_1035_v63_ = d_1035_v63_
            elif True:
                (globalState).f9 = (self.f14 if self.f14 else (self).f19)
                d_1056_v78_: _dafny.Array
                def lambda49_(d_1057_i9_):
                    return ((self).f15).isdisjoint((self).f15)

                init27_ = lambda49_
                nw182_ = _dafny.Array(None, 6)
                for i0_27_ in range(nw182_.length(0)):
                    nw182_[i0_27_] = init27_(i0_27_)
                d_1056_v78_ = nw182_
                (globalState).f2 = d_1056_v78_
                r2 = self.f14
                d_1058_v79_: _dafny.Map
                d_1058_v79_ = _dafny.Map({(self).f19: self.f14})
                d_1059_v80_: _dafny.Set
                d_1059_v80_ = _dafny.Set({d_1058_v79_, d_1058_v79_, d_1058_v79_})
                d_1060_v81_: _dafny.Map
                d_1060_v81_ = _dafny.Map({d_1059_v80_: (0) - ((p1) - (48))})
                d_1060_v81_ = (d_1060_v81_).set(d_1059_v80_, p1)
                d_1061_v82_: _dafny.Array
                nw183_ = _dafny.Array(int(0), 25)
                d_1061_v82_ = nw183_
                index197_ = default__.safeIndex(635, (d_1061_v82_).length(0))
                (d_1061_v82_)[index197_] = default__.safeModuloInt(p1, p1)
                index198_ = default__.safeIndex(635, (d_1061_v82_).length(0))
                (d_1061_v82_)[index198_] = -600
            d_1062_v83_: str
            d_1062_v83_ = _dafny.CodePoint('c')
            d_1063_v84_: _dafny.Map
            d_1063_v84_ = _dafny.Map({d_1062_v83_: p1})
            d_1064_v85_: _dafny.Seq
            d_1064_v85_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1065_v86_: _dafny.Map
            d_1065_v86_ = _dafny.Map({d_1064_v85_: (self).f19})
            d_1066_v88_: D8
            def iife110_():
                coll44_ = _dafny.Set()
                compr_44_: _dafny.Seq
                for compr_44_ in (d_1065_v86_).keys.Elements:
                    d_1067_v87_: _dafny.Seq = compr_44_
                    if (d_1067_v87_) in (d_1065_v86_):
                        coll44_ = coll44_.union(_dafny.Set([d_1067_v87_]))
                return _dafny.Set(coll44_)
            d_1066_v88_ = D8_DC24(_dafny.Set({self.f14, default__.fm23(p1, self.f14, d_1063_v84_, iife110_()
, globalState)}))
            d_1066_v88_ = d_1066_v88_
            d_1068_v89_: _dafny.Array
            def lambda50_(d_1069_i10_):
                return default__.safeDivisionInt(d_1069_i10_, 380)

            init28_ = lambda50_
            nw184_ = _dafny.Array(None, 1)
            for i0_28_ in range(nw184_.length(0)):
                nw184_[i0_28_] = init28_(i0_28_)
            d_1068_v89_ = nw184_
            d_1070_v90_: _dafny.Map
            d_1070_v90_ = _dafny.Map({p1: d_1035_v63_})
            d_1071_v91_: D8
            d_1071_v91_ = D8_DC28(D8_DC26(not(self.f14), d_1070_v90_))
            d_1072_v92_: _dafny.MultiSet
            d_1072_v92_ = _dafny.MultiSet([d_1071_v91_])
            d_1073_v93_: _dafny.Map
            d_1073_v93_ = _dafny.Map({d_1068_v89_: p1})
            d_1074_v94_: _dafny.Map
            d_1074_v94_ = _dafny.Map({(self).f19: d_1073_v93_})
            d_1075_v95_: _dafny.Seq
            d_1075_v95_ = _dafny.SeqWithoutIsStrInference([self.f14, not(self.f14)])
            d_1076_v96_: _dafny.Seq
            d_1076_v96_ = _dafny.SeqWithoutIsStrInference([d_1068_v89_, d_1068_v89_])
            d_1077_v97_: D3
            d_1077_v97_ = D3_DC11((d_1072_v92_).cardinality, len(((d_1074_v94_)[(d_1075_v95_)[default__.safeIndex(p1, len(d_1075_v95_))]] if ((d_1075_v95_)[default__.safeIndex(p1, len(d_1075_v95_))]) in (d_1074_v94_) else _dafny.Map({(d_1076_v96_)[default__.safeIndex(p1, len(d_1076_v96_))]: p1}))), (self).f19, p1)
            rhs154_ = d_1068_v89_
            rhs155_ = (self).f19
            rhs156_ = d_1077_v97_
            rhs157_ = not((self).f19)
            lhs134_ = globalState
            lhs135_ = self
            d_1068_v89_ = rhs154_
            lhs134_.f9 = rhs155_
            d_1077_v97_ = rhs156_
            lhs135_.f14 = rhs157_
            (globalState).f9 = (d_1075_v95_)[default__.safeIndex(525, len(d_1075_v95_))]
        if ((self).f19) and (not(self.f14)):
            d_1078_v98_: _dafny.Seq
            d_1078_v98_ = _dafny.SeqWithoutIsStrInference([(self).f19])
            r1 = (((_dafny.MultiSet([self.f14])).intersection((self).f15)) | ((_dafny.MultiSet(d_1078_v98_)) | ((self).f15))).cardinality
            if not((self).fm2((p1) > (p1), (self).f19, globalState)):
                (globalState).f9 = (self).f19
                d_1035_v63_ = d_1035_v63_
                r1 = (p1) + (p1)
                d_1079_v99_: _dafny.Array
                def lambda51_(d_1080_p1_):
                    def lambda52_(d_1081_i11_):
                        return (d_1081_i11_) + ((0) - (d_1080_p1_))

                    return lambda52_

                init29_ = lambda51_(p1)
                nw185_ = _dafny.Array(None, 10)
                for i0_29_ in range(nw185_.length(0)):
                    nw185_[i0_29_] = init29_(i0_29_)
                d_1079_v99_ = nw185_
                index199_ = default__.safeIndex(808, (d_1079_v99_).length(0))
                (d_1079_v99_)[index199_] = (p1) + (p1)
                d_1082_v100_: _dafny.Seq
                d_1082_v100_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f19, self.f14])), p1])
                d_1083_v101_: _dafny.MultiSet
                d_1083_v101_ = _dafny.MultiSet([p1])
                d_1084_v102_: _dafny.Seq
                d_1084_v102_ = _dafny.SeqWithoutIsStrInference([(p1) * (default__.fm0(self.f14, len(d_1082_v100_), p1, d_1083_v101_, globalState))])
                d_1085_v103_: _dafny.Array
                nw186_ = _dafny.Array(None, 9)
                nw186_[int(0)] = not((self).f19)
                nw186_[int(1)] = (self).f19
                nw186_[int(2)] = self.f14
                nw186_[int(3)] = self.f14
                nw186_[int(4)] = (self).f19
                nw186_[int(5)] = (self).f19
                nw186_[int(6)] = not(self.f14)
                nw186_[int(7)] = self.f14
                nw186_[int(8)] = (self).f19
                d_1085_v103_ = nw186_
                d_1086_v104_: str
                d_1086_v104_ = _dafny.CodePoint('n')
                d_1087_v105_: _dafny.Map
                d_1087_v105_ = _dafny.Map({d_1086_v104_: len(d_1035_v63_)})
                d_1088_v106_: _dafny.Set
                d_1088_v106_ = _dafny.Set({d_1082_v100_})
                d_1089_v107_: D7
                d_1089_v107_ = D7_DC23(default__.fm23(len(_dafny.SeqWithoutIsStrInference([d_1085_v103_])), (self).f19, d_1087_v105_, d_1088_v106_, globalState), p1, self.f14)
                index200_ = default__.safeIndex(808, (d_1079_v99_).length(0))
                rhs158_ = (d_1084_v102_)[default__.safeIndex((d_1089_v107_).cf43, len(d_1084_v102_))]
                rhs159_ = p1
                lhs136_ = d_1079_v99_
                lhs137_ = default__.safeIndex(808, (d_1079_v99_).length(0))
                lhs136_[lhs137_] = rhs158_
                r0 = rhs159_
                d_1090_v108_: _dafny.Map
                d_1090_v108_ = _dafny.Map({p1: False})
                d_1091_v109_: _dafny.Map
                d_1091_v109_ = _dafny.Map({d_1090_v108_: (d_1079_v99_)[default__.safeIndex(808, (d_1079_v99_).length(0))]})
                d_1092_v110_: D7
                d_1092_v110_ = D7_DC21(d_1091_v109_)
                d_1093_v111_: _dafny.Set
                d_1093_v111_ = _dafny.Set({self.f14, self.f14, (self).f19, self.f14, self.f14})
                index201_ = default__.safeIndex(808, (d_1079_v99_).length(0))
                rhs160_ = len(((_dafny.Set({self.f14, False, self.f14, self.f14})).intersection(_dafny.Set({(self).f19}))).intersection(d_1093_v111_))
                rhs161_ = (d_1079_v99_)[default__.safeIndex(808, (d_1079_v99_).length(0))]
                rhs162_ = (d_1079_v99_)[default__.safeIndex(808, (d_1079_v99_).length(0))]
                rhs163_ = d_1092_v110_
                lhs138_ = globalState
                lhs139_ = d_1079_v99_
                lhs140_ = default__.safeIndex(808, (d_1079_v99_).length(0))
                lhs138_.f3 = rhs160_
                r0 = rhs161_
                lhs139_[lhs140_] = rhs162_
                d_1092_v110_ = rhs163_
            elif True:
                d_1094_v112_: _dafny.MultiSet
                d_1094_v112_ = _dafny.MultiSet([D8_DC25()])
                d_1095_v113_: D2
                d_1095_v113_ = D2_DC7(d_1078_v98_)
                rhs164_ = len((d_1095_v113_).cf9)
                rhs165_ = default__.safeModuloInt((p1) + (p1), p1)
                rhs166_ = (d_1094_v112_).intersection(d_1094_v112_)
                rhs167_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                lhs141_ = globalState
                lhs141_.f3 = rhs164_
                r1 = rhs165_
                d_1094_v112_ = rhs166_
                d_1035_v63_ = rhs167_
                (globalState).f3 = p1
                d_1096_v114_: _dafny.Seq
                d_1096_v114_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1097_v115_: _dafny.MultiSet
                d_1097_v115_ = _dafny.MultiSet([p1])
                d_1098_v116_: _dafny.Map
                d_1098_v116_ = _dafny.Map({(self).f19: (_dafny.MultiSet(d_1096_v114_)).intersection(d_1097_v115_)})
                d_1098_v116_ = (d_1098_v116_).set(self.f14, (d_1097_v115_).intersection(d_1097_v115_))
                d_1099_v117_: D3
                d_1099_v117_ = D3_DC11(p1, p1, (self).f19, p1)
                d_1100_v118_: C3
                nw187_ = C3()
                nw187_.ctor__(p1, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1101_i12_ in range(default__.abs(669))]), (self).f19, (self).f15)
                d_1100_v118_ = nw187_
                d_1102_v119_: _dafny.MultiSet
                d_1102_v119_ = _dafny.MultiSet([d_1100_v118_, d_1100_v118_])
                d_1103_v120_: str
                d_1103_v120_ = _dafny.CodePoint('q')
                d_1104_v121_: _dafny.Set
                d_1104_v121_ = _dafny.Set({d_1103_v120_, d_1103_v120_})
                d_1105_v122_: _dafny.MultiSet
                d_1105_v122_ = _dafny.MultiSet([d_1104_v121_])
                d_1106_v123_: _dafny.Map
                d_1106_v123_ = _dafny.Map({self.f14: (d_1105_v122_).cardinality})
                d_1107_v124_: D8
                d_1107_v124_ = D8_DC24(_dafny.Set({(self).f19, False}))
                d_1108_v125_: _dafny.Map
                d_1108_v125_ = _dafny.Map({d_1107_v124_: self.f14})
                d_1109_v126_: _dafny.Map
                d_1109_v126_ = _dafny.Map({self.f14: (self).f19})
                pat_let_tv23_ = d_1108_v125_
                pat_let_tv24_ = p1
                pat_let_tv25_ = d_1106_v123_
                pat_let_tv26_ = d_1100_v118_
                pat_let_tv27_ = d_1097_v115_
                pat_let_tv28_ = globalState
                pat_let_tv29_ = d_1100_v118_
                pat_let_tv30_ = d_1097_v115_
                pat_let_tv31_ = globalState
                pat_let_tv32_ = d_1100_v118_
                d_1110_v127_: _dafny.Array
                nw188_ = _dafny.Array(None, 20)
                nw188_[int(0)] = d_1099_v117_
                nw188_[int(1)] = d_1099_v117_
                def iife111_(_pat_let33_0):
                    def iife112_(d_1111_dt__update__tmp_h3_):
                        def iife113_(_pat_let34_0):
                            def iife114_(d_1112_dt__update_hcf19_h0_):
                                return D3_DC11((d_1111_dt__update__tmp_h3_).cf17, (d_1111_dt__update__tmp_h3_).cf18, d_1112_dt__update_hcf19_h0_, (d_1111_dt__update__tmp_h3_).cf20)
                            return iife114_(_pat_let34_0)
                        return iife113_(self.f14)
                    return iife112_(_pat_let33_0)
                nw188_[int(2)] = (D3_DC11(p1, p1, (self).f19, 967) if (self).f19 else iife111_(d_1099_v117_))
                nw188_[int(3)] = D3_DC11(p1, p1, self.f14, 485)
                nw188_[int(4)] = d_1099_v117_
                nw188_[int(5)] = d_1099_v117_
                nw188_[int(6)] = d_1099_v117_
                nw188_[int(7)] = D3_DC11(p1, p1, True, (0) - ((0) - ((d_1102_v119_).cardinality)))
                nw188_[int(8)] = d_1099_v117_
                nw188_[int(9)] = d_1099_v117_
                nw188_[int(10)] = d_1099_v117_
                nw188_[int(11)] = d_1099_v117_
                nw188_[int(12)] = d_1099_v117_
                nw188_[int(13)] = d_1099_v117_
                def iife115_(_pat_let35_0):
                    def iife116_(d_1113_dt__update__tmp_h4_):
                        def iife117_(_pat_let36_0):
                            def iife118_(d_1114_dt__update_hcf19_h1_):
                                def iife119_(_pat_let37_0):
                                    def iife120_(d_1115_dt__update_hcf18_h0_):
                                        def iife121_(_pat_let38_0):
                                            def iife122_(d_1116_dt__update_hcf20_h0_):
                                                return D3_DC11((d_1113_dt__update__tmp_h4_).cf17, d_1115_dt__update_hcf18_h0_, d_1114_dt__update_hcf19_h1_, d_1116_dt__update_hcf20_h0_)
                                            return iife122_(_pat_let38_0)
                                        return iife121_(default__.fm0((self).f19, default__.fm0(self.f14, len(_dafny.Set({pat_let_tv24_, len(pat_let_tv25_)})), (pat_let_tv26_).f20, pat_let_tv27_, pat_let_tv28_), (pat_let_tv29_).f20, pat_let_tv30_, pat_let_tv31_))
                                    return iife120_(_pat_let37_0)
                                return iife119_(len(pat_let_tv23_))
                            return iife118_(_pat_let36_0)
                        return iife117_((self).f19)
                    return iife116_(_pat_let35_0)
                nw188_[int(14)] = iife115_(D3_DC11(p1, default__.fm0(True, default__.fm0((self).f19, 535, p1, d_1097_v115_, globalState), (d_1100_v118_).f20, _dafny.MultiSet(d_1096_v114_), globalState), self.f14, p1))
                nw188_[int(15)] = d_1099_v117_
                nw188_[int(16)] = d_1099_v117_
                nw188_[int(17)] = D3_DC11(len(d_1106_v123_), len((d_1109_v126_).set((self).f19, (self).f19)), self.f14, (d_1100_v118_).f20)
                def iife123_(_pat_let39_0):
                    def iife124_(d_1117_dt__update__tmp_h5_):
                        def iife125_(_pat_let40_0):
                            def iife126_(d_1118_dt__update_hcf20_h1_):
                                return D3_DC11((d_1117_dt__update__tmp_h5_).cf17, (d_1117_dt__update__tmp_h5_).cf18, (d_1117_dt__update__tmp_h5_).cf19, d_1118_dt__update_hcf20_h1_)
                            return iife126_(_pat_let40_0)
                        return iife125_((pat_let_tv32_).f20)
                    return iife124_(_pat_let39_0)
                nw188_[int(18)] = (d_1099_v117_ if (self).f19 else iife123_(d_1099_v117_))
                nw188_[int(19)] = d_1099_v117_
                d_1110_v127_ = nw188_
                index202_ = default__.safeIndex(129, (d_1110_v127_).length(0))
                (d_1110_v127_)[index202_] = d_1099_v117_
                d_1119_v128_: _dafny.Array
                def lambda53_(d_1120_i13_):
                    return (self).f15

                init30_ = lambda53_
                nw189_ = _dafny.Array(None, 1)
                for i0_30_ in range(nw189_.length(0)):
                    nw189_[i0_30_] = init30_(i0_30_)
                d_1119_v128_ = nw189_
                index203_ = default__.safeIndex(891, (d_1119_v128_).length(0))
                (d_1119_v128_)[index203_] = (self).f15
                d_1121_v129_: _dafny.Map
                d_1121_v129_ = _dafny.Map({default__.fm25(p1, len(d_1078_v98_), self.f14, d_1106_v123_, globalState): p1})
                d_1122_v130_: _dafny.Set
                d_1122_v130_ = _dafny.Set({d_1096_v114_})
                d_1123_v133_: _dafny.Map
                d_1123_v133_ = _dafny.Map({d_1078_v98_: False})
                pat_let_tv33_ = d_1103_v120_
                pat_let_tv34_ = d_1121_v129_
                pat_let_tv35_ = p1
                index204_ = default__.safeIndex(129, (d_1110_v127_).length(0))
                index205_ = default__.safeIndex(891, (d_1119_v128_).length(0))
                def iife127_(_pat_let41_0):
                    def iife128_(d_1124_dt__update__tmp_h6_):
                        def iife129_(_pat_let42_0):
                            def iife130_(d_1125_dt__update_hcf19_h2_):
                                def iife131_(_pat_let43_0):
                                    def iife132_(d_1126_dt__update_hcf18_h1_):
                                        return D3_DC11((d_1124_dt__update__tmp_h6_).cf17, d_1126_dt__update_hcf18_h1_, d_1125_dt__update_hcf19_h2_, (d_1124_dt__update__tmp_h6_).cf20)
                                    return iife132_(_pat_let43_0)
                                return iife131_(pat_let_tv35_)
                            return iife130_(_pat_let42_0)
                        return iife129_((pat_let_tv33_) not in (pat_let_tv34_))
                    return iife128_(_pat_let41_0)
                def iife133_():
                    coll45_ = _dafny.Map()
                    compr_45_: str
                    for compr_45_ in ((d_1100_v118_).f21).Elements:
                        d_1127_v131_: str = compr_45_
                        if (d_1127_v131_) in ((d_1100_v118_).f21):
                            coll45_[d_1127_v131_] = p1
                    return _dafny.Map(coll45_)
                def iife134_():
                    coll46_ = _dafny.Set()
                    compr_46_: _dafny.Seq
                    for compr_46_ in (_dafny.MultiSet([d_1096_v114_, _dafny.SeqWithoutIsStrInference([-546])])).Elements:
                        d_1128_v132_: _dafny.Seq = compr_46_
                        if (d_1128_v132_) in (_dafny.MultiSet([d_1096_v114_, _dafny.SeqWithoutIsStrInference([-546])])):
                            coll46_ = coll46_.union(_dafny.Set([d_1128_v132_]))
                    return _dafny.Set(coll46_)
                rhs168_ = iife127_(d_1099_v117_)
                rhs169_ = (self).f15
                rhs170_ = (((_dafny.Map({_dafny.SeqWithoutIsStrInference([default__.fm23(len((d_1100_v118_).f21), default__.fm23((d_1100_v118_).f20, (self).f19, d_1121_v129_, d_1122_v130_, globalState), iife133_()
                , iife134_()
                , globalState), self.f14]): (self).f19})).set(_dafny.SeqWithoutIsStrInference([(self).f19]), (self).f19)).set(d_1078_v98_, (self).f19)) == (d_1123_v133_)
                lhs142_ = d_1110_v127_
                lhs143_ = default__.safeIndex(129, (d_1110_v127_).length(0))
                lhs144_ = d_1119_v128_
                lhs145_ = default__.safeIndex(891, (d_1119_v128_).length(0))
                lhs146_ = globalState
                lhs142_[lhs143_] = rhs168_
                lhs144_[lhs145_] = rhs169_
                lhs146_.f9 = rhs170_
                d_1129_v134_: _dafny.Seq
                d_1129_v134_ = _dafny.SeqWithoutIsStrInference([(d_1109_v126_) | (d_1109_v126_), d_1109_v126_, _dafny.Map({True: False}), d_1109_v126_])
                d_1129_v134_ = (d_1129_v134_) + (d_1129_v134_)
            d_1130_v135_: D7
            d_1130_v135_ = D7_DC23((self).f19, p1, self.f14)
            d_1131_v136_: _dafny.Seq
            d_1131_v136_ = _dafny.SeqWithoutIsStrInference([d_1078_v98_, default__.fm38(globalState), d_1078_v98_])
            (globalState).f3 = ((d_1130_v135_).cf43) * (len(d_1131_v136_))
            (globalState).f3 = p1
            d_1132_v137_: _dafny.Array
            nw190_ = _dafny.Array(int(0), 1)
            d_1132_v137_ = nw190_
            index206_ = default__.safeIndex(773, (d_1132_v137_).length(0))
            (d_1132_v137_)[index206_] = (0) - (len(d_1078_v98_))
            d_1133_v138_: _dafny.Set
            d_1133_v138_ = _dafny.Set({(self).f19, self.f14})
            index207_ = default__.safeIndex(773, (d_1132_v137_).length(0))
            rhs171_ = (0) - (p1)
            rhs172_ = (d_1133_v138_).ispropersubset(d_1133_v138_)
            rhs173_ = default__.safeDivisionInt(p1, p1)
            lhs147_ = globalState
            lhs148_ = self
            lhs149_ = d_1132_v137_
            lhs150_ = default__.safeIndex(773, (d_1132_v137_).length(0))
            lhs147_.f3 = rhs171_
            lhs148_.f14 = rhs172_
            lhs149_[lhs150_] = rhs173_
        elif True:
            d_1134_v139_: _dafny.MultiSet
            d_1134_v139_ = _dafny.MultiSet([((self).f19) or ((self).f19)])
            d_1135_v140_: _dafny.Array
            nw191_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_1135_v140_ = nw191_
            d_1136_v141_: _dafny.Seq
            d_1136_v141_ = _dafny.SeqWithoutIsStrInference([(self).f19])
            d_1137_v142_: D0
            d_1137_v142_ = D0_DC1(self.f14)
            d_1138_v143_: _dafny.Array
            nw192_ = _dafny.Array(None, 12)
            nw192_[int(0)] = self.f14
            nw192_[int(1)] = self.f14
            nw192_[int(2)] = self.f14
            nw192_[int(3)] = self.f14
            nw192_[int(4)] = not(self.f14)
            nw192_[int(5)] = (d_1136_v141_)[default__.safeIndex(p1, len(d_1136_v141_))]
            nw192_[int(6)] = (self).f19
            nw192_[int(7)] = self.f14
            nw192_[int(8)] = (d_1137_v142_).cf1
            nw192_[int(9)] = (self).f19
            nw192_[int(10)] = (self).f19
            nw192_[int(11)] = self.f14
            d_1138_v143_ = nw192_
            index208_ = default__.safeIndex(985, (d_1135_v140_).length(0))
            (d_1135_v140_)[index208_] = d_1138_v143_
            index209_ = default__.safeIndex(985, (d_1135_v140_).length(0))
            rhs174_ = (d_1134_v139_) | ((self).f15)
            rhs175_ = (d_1138_v143_ if False else d_1138_v143_)
            lhs151_ = d_1135_v140_
            lhs152_ = default__.safeIndex(985, (d_1135_v140_).length(0))
            d_1134_v139_ = rhs174_
            lhs151_[lhs152_] = rhs175_
            index210_ = default__.safeIndex(828, (d_1138_v143_).length(0))
            (d_1138_v143_)[index210_] = (self).f19
            d_1139_v144_: _dafny.Set
            d_1139_v144_ = _dafny.Set({(self).f19, (self).f19, False})
            index211_ = default__.safeIndex(828, (d_1138_v143_).length(0))
            (d_1138_v143_)[index211_] = not (((0) - ((0) - (len(d_1139_v144_)))) >= (70)) or (not(self.f14))
            d_1140_v145_: _dafny.Map
            d_1140_v145_ = _dafny.Map({709: p1})
            d_1141_v147_: _dafny.Map
            d_1141_v147_ = _dafny.Map({len(d_1035_v63_): (self).f19})
            d_1142_v148_: _dafny.Seq
            def iife135_():
                coll47_ = _dafny.Map()
                compr_47_: int
                for compr_47_ in (p0).Elements:
                    d_1143_v146_: int = compr_47_
                    if (d_1143_v146_) in (p0):
                        coll47_[(d_1143_v146_) + (p1)] = p1
                return _dafny.Map(coll47_)
            d_1142_v148_ = _dafny.SeqWithoutIsStrInference([(iife135_()
            ).set(889, len(d_1141_v147_)), d_1140_v145_])
            def iife136_():
                coll48_ = _dafny.Map()
                compr_48_: int
                for compr_48_ in _dafny.IntegerRange(581, 439):
                    d_1144_v149_: int = compr_48_
                    if ((581) <= (d_1144_v149_)) and ((d_1144_v149_) < (439)):
                        coll48_[default__.safeDivisionInt(d_1144_v149_, len(d_1139_v144_))] = (self).f19
                return _dafny.Map(coll48_)
            rhs176_ = default__.safeModuloInt((0) - (len(d_1035_v63_)), p1)
            rhs177_ = (d_1140_v145_) | ((d_1142_v148_)[default__.safeIndex(len(iife136_()
            ), len(d_1142_v148_))])
            rhs178_ = d_1035_v63_
            lhs153_ = globalState
            lhs153_.f3 = rhs176_
            d_1140_v145_ = rhs177_
            d_1035_v63_ = rhs178_
            d_1145_v150_: _dafny.MultiSet
            d_1145_v150_ = _dafny.MultiSet([p1, p1])
            d_1146_v151_: D1
            d_1146_v151_ = D1_DC4(default__.fm0((d_1138_v143_)[default__.safeIndex(828, (d_1138_v143_).length(0))], p1, p1, d_1145_v150_, globalState), p1)
            d_1147_v153_: D3
            def iife137_():
                coll49_ = _dafny.Map()
                compr_49_: int
                for compr_49_ in (d_1145_v150_).Elements:
                    d_1148_v152_: int = compr_49_
                    if (d_1148_v152_) in (d_1145_v150_):
                        coll49_[(d_1148_v152_) * (p1)] = d_1140_v145_
                return _dafny.Map(coll49_)
            d_1147_v153_ = D3_DC11((d_1146_v151_).cf5, len(iife137_()
), self.f14, p1)
            source13_ = d_1147_v153_
            if source13_.is_DC11:
                d_1149___mcc_h3_ = source13_.cf17
                d_1150___mcc_h4_ = source13_.cf18
                d_1151___mcc_h5_ = source13_.cf19
                d_1152___mcc_h6_ = source13_.cf20
                d_1153_cf20_ = d_1152___mcc_h6_
                d_1154_cf19_ = d_1151___mcc_h5_
                d_1155_cf18_ = d_1150___mcc_h4_
                d_1156_cf17_ = d_1149___mcc_h3_
                d_1157_v154_: _dafny.Map
                d_1157_v154_ = _dafny.Map({(d_1035_v63_) + (d_1035_v63_): d_1140_v145_})
                d_1158_v155_: str
                d_1158_v155_ = _dafny.CodePoint('d')
                d_1159_v156_: D9
                d_1159_v156_ = D9_DC29(d_1140_v145_)
                d_1157_v154_ = (d_1157_v154_).set((d_1035_v63_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1160_i14_ in range(default__.abs(345))])), len(d_1035_v63_)), d_1158_v155_), (d_1159_v156_).cf49)
                r2 = ((d_1141_v147_)[(d_1153_cf20_) * (d_1156_cf17_)] if ((d_1153_cf20_) * (d_1156_cf17_)) in (d_1141_v147_) else d_1154_cf19_)
                d_1161_v157_: _dafny.Set
                d_1161_v157_ = _dafny.Set({d_1158_v155_, d_1158_v155_})
                d_1162_v158_: _dafny.Seq
                d_1162_v158_ = _dafny.SeqWithoutIsStrInference([d_1161_v157_, (d_1161_v157_) - (d_1161_v157_)])
                d_1162_v158_ = d_1162_v158_
                (self).f14 = not (d_1154_cf19_) or ((_dafny.SeqWithoutIsStrInference([d_1158_v155_ for d_1163_i15_ in range(default__.abs(137))])) != (d_1035_v63_))
            elif True:
                d_1164___mcc_h7_ = source13_.cf16
                d_1165_cf16_ = d_1164___mcc_h7_
                (self).m6(globalState)
                d_1166_v159_: _dafny.Array
                nw193_ = _dafny.Array(None, 1)
                nw193_[int(0)] = (p1) - (p1)
                d_1166_v159_ = nw193_
                index212_ = default__.safeIndex(918, (d_1166_v159_).length(0))
                (d_1166_v159_)[index212_] = 966
                index213_ = default__.safeIndex(918, (d_1166_v159_).length(0))
                (d_1166_v159_)[index213_] = ((p1) + (default__.fm0(self.f14, p1, p1, default__.fm30((self).f19, globalState), globalState))) - (p1)
                r2 = True
                index214_ = default__.safeIndex(828, (d_1138_v143_).length(0))
                rhs179_ = ((d_1166_v159_)[default__.safeIndex(918, (d_1166_v159_).length(0))]) >= ((d_1146_v151_).cf4)
                rhs180_ = d_1138_v143_
                rhs181_ = -942
                rhs182_ = (p1) * ((d_1166_v159_)[default__.safeIndex(918, (d_1166_v159_).length(0))])
                lhs154_ = d_1138_v143_
                lhs155_ = default__.safeIndex(828, (d_1138_v143_).length(0))
                lhs156_ = globalState
                lhs157_ = globalState
                lhs154_[lhs155_] = rhs179_
                lhs156_.f2 = rhs180_
                r0 = rhs181_
                lhs157_.f3 = rhs182_
            d_1167_v160_: _dafny.Map
            d_1167_v160_ = _dafny.Map({(d_1135_v140_)[default__.safeIndex(985, (d_1135_v140_).length(0))]: p1})
            d_1168_v161_: D4
            d_1168_v161_ = D4_DC13(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1169_i16_ in range(default__.abs(671))]), 786, (self).f19, (d_1135_v140_)[default__.safeIndex(985, (d_1135_v140_).length(0))])
            d_1170_v162_: _dafny.MultiSet
            d_1170_v162_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([self.f14, (self).f19])])
            rhs183_ = ((default__.fm30((d_1138_v143_)[default__.safeIndex(828, (d_1138_v143_).length(0))], globalState)).cardinality) + (p1)
            rhs184_ = (0) - ((default__.safeDivisionInt(p1, p1) if ((0) - (len(d_1167_v160_))) != (p1) else p1))
            rhs185_ = ((p1) * (p1)) - ((0) - (p1))
            rhs186_ = (p1) + (((0) - ((0) - (len((d_1168_v161_).cf22)))) * ((0) - (p1)))
            rhs187_ = default__.safeDivisionInt(default__.safeDivisionInt((self).fm17(p1, self.f14, d_1170_v162_, p1, globalState), p1), p1)
            lhs158_ = globalState
            r1 = rhs183_
            lhs158_.f3 = rhs184_
            r1 = rhs185_
            r0 = rhs186_
            r1 = rhs187_
        d_1171_v163_: _dafny.Array
        nw194_ = _dafny.Array(False, 15)
        d_1171_v163_ = nw194_
        d_1172_v164_: _dafny.Seq
        d_1172_v164_ = _dafny.SeqWithoutIsStrInference([d_1171_v163_])
        d_1173_v165_: D4
        d_1173_v165_ = D4_DC13(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uoqpauedn")), p1, self.f14, (d_1172_v164_)[default__.safeIndex(p1, len(d_1172_v164_))])
        def iife138_(_pat_let44_0):
            def iife139_(d_1174_dt__update__tmp_h7_):
                def iife140_(_pat_let45_0):
                    def iife141_(d_1176_dt__update_hcf22_h0_):
                        return D4_DC13(d_1176_dt__update_hcf22_h0_, (d_1174_dt__update__tmp_h7_).cf23, (d_1174_dt__update__tmp_h7_).cf24, (d_1174_dt__update__tmp_h7_).cf25)
                    return iife141_(_pat_let45_0)
                return iife140_(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1175_i17_ in range(default__.abs(431))]))
            return iife139_(_pat_let44_0)
        source14_ = iife138_(d_1173_v165_)
        if source14_.is_DC13:
            d_1177___mcc_h8_ = source14_.cf22
            d_1178___mcc_h9_ = source14_.cf23
            d_1179___mcc_h10_ = source14_.cf24
            d_1180___mcc_h11_ = source14_.cf25
            d_1181_cf25_ = d_1180___mcc_h11_
            d_1182_cf24_ = d_1179___mcc_h10_
            d_1183_cf23_ = d_1178___mcc_h9_
            d_1184_cf22_ = d_1177___mcc_h8_
            d_1185_v166_: _dafny.Map
            d_1185_v166_ = _dafny.Map({d_1183_cf23_: d_1182_cf24_})
            d_1186_v167_: _dafny.Map
            d_1186_v167_ = _dafny.Map({len(d_1035_v63_): d_1185_v166_})
            d_1186_v167_ = (d_1186_v167_).set(default__.safeDivisionInt(d_1183_cf23_, d_1183_cf23_), d_1185_v166_)
            index215_ = default__.safeIndex(241, (d_1181_cf25_).length(0))
            (d_1181_cf25_)[index215_] = d_1182_cf24_
            index216_ = default__.safeIndex(241, (d_1181_cf25_).length(0))
            (d_1181_cf25_)[index216_] = (self).fm2((self).f19, self.f14, globalState)
            if (d_1181_cf25_)[default__.safeIndex(241, (d_1181_cf25_).length(0))]:
                d_1187_v168_: _dafny.Array
                def lambda54_(d_1188_p1_):
                    def lambda55_(d_1189_i18_):
                        return (d_1189_i18_) + (d_1188_p1_)

                    return lambda55_

                init31_ = lambda54_(p1)
                nw195_ = _dafny.Array(None, 17)
                for i0_31_ in range(nw195_.length(0)):
                    nw195_[i0_31_] = init31_(i0_31_)
                d_1187_v168_ = nw195_
                d_1190_v169_: _dafny.Map
                d_1190_v169_ = _dafny.Map({p1: d_1187_v168_})
                (globalState).f3 = (0) - (len(((d_1190_v169_) | (d_1190_v169_)) | (d_1190_v169_)))
                d_1191_v170_: C3
                nw196_ = C3()
                nw196_.ctor__(p1, (d_1035_v63_).set(default__.safeIndex(d_1183_cf23_, len(d_1035_v63_)), _dafny.CodePoint('t')), False, (self).f15)
                d_1191_v170_ = nw196_
                d_1191_v170_ = d_1191_v170_
                d_1192_v171_: C3
                nw197_ = C3()
                nw197_.ctor__((p1) + ((d_1191_v170_).f20), (d_1035_v63_) + (d_1184_cf22_), False, (self).f15)
                d_1192_v171_ = nw197_
                (globalState).f3 = ((self).f15).cardinality
                d_1193_v172_: str
                d_1193_v172_ = _dafny.CodePoint('g')
                d_1194_v173_: D9
                d_1194_v173_ = D9_DC30(d_1193_v172_, d_1182_cf24_, p1, p0, (d_1192_v171_).f20)
                d_1195_v174_: _dafny.Seq
                d_1195_v174_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1196_v175_: D2
                d_1196_v175_ = D2_DC9((d_1181_cf25_)[default__.safeIndex(241, (d_1181_cf25_).length(0))], (d_1192_v171_).f20, d_1193_v172_)
                d_1197_v176_: _dafny.Seq
                d_1197_v176_ = _dafny.SeqWithoutIsStrInference([self.f14])
                d_1198_v177_: _dafny.Set
                d_1198_v177_ = _dafny.Set({d_1035_v63_})
                pat_let_tv36_ = d_1193_v172_
                d_1199_v179_: _dafny.Array
                nw198_ = _dafny.Array(None, 21)
                nw198_[int(0)] = D9_DC30(d_1193_v172_, (self).f19, d_1183_cf23_, p0, p1)
                nw198_[int(1)] = d_1194_v173_
                nw198_[int(2)] = d_1194_v173_
                nw198_[int(3)] = d_1194_v173_
                def iife142_(_pat_let46_0):
                    def iife143_(d_1200_dt__update__tmp_h8_):
                        def iife144_(_pat_let47_0):
                            def iife145_(d_1201_dt__update_hcf50_h0_):
                                return D9_DC30(d_1201_dt__update_hcf50_h0_, (d_1200_dt__update__tmp_h8_).cf51, (d_1200_dt__update__tmp_h8_).cf52, (d_1200_dt__update__tmp_h8_).cf53, (d_1200_dt__update__tmp_h8_).cf54)
                            return iife145_(_pat_let47_0)
                        return iife144_(pat_let_tv36_)
                    return iife143_(_pat_let46_0)
                nw198_[int(4)] = iife142_(d_1194_v173_)
                nw198_[int(5)] = d_1194_v173_
                nw198_[int(6)] = d_1194_v173_
                nw198_[int(7)] = D9_DC30(d_1193_v172_, d_1182_cf24_, (d_1191_v170_).f20, p0, 746)
                nw198_[int(8)] = D9_DC30(d_1193_v172_, (self).f19, (d_1192_v171_).f20, p0, p1)
                nw198_[int(9)] = D9_DC30(d_1193_v172_, False, (0) - (len(_dafny.SeqWithoutIsStrInference([d_1195_v174_, d_1195_v174_]))), _dafny.Set({(d_1192_v171_).f20, 5}), (d_1192_v171_).f20)
                nw198_[int(10)] = D9_DC30((d_1196_v175_).cf15, (d_1181_cf25_)[default__.safeIndex(241, (d_1181_cf25_).length(0))], (d_1191_v170_).f20, p0, (d_1191_v170_).f20)
                nw198_[int(11)] = d_1194_v173_
                nw198_[int(12)] = default__.fm39(len((d_1197_v176_).set(default__.safeIndex(p1, len(d_1197_v176_)), (d_1181_cf25_)[default__.safeIndex(241, (d_1181_cf25_).length(0))])), (d_1191_v170_).f20, p1, len(d_1198_v177_), globalState)
                nw198_[int(13)] = d_1194_v173_
                nw198_[int(14)] = D9_DC30(d_1193_v172_, (self).f19, (d_1191_v170_).f20, p0, p1)
                nw198_[int(15)] = d_1194_v173_
                def iife146_():
                    coll50_ = _dafny.Map()
                    compr_50_: str
                    for compr_50_ in (_dafny.SeqWithoutIsStrInference([d_1193_v172_ for d_1202_i19_ in range(default__.abs(-580))])).Elements:
                        d_1203_v178_: str = compr_50_
                        if (d_1203_v178_) in (_dafny.SeqWithoutIsStrInference([d_1193_v172_ for d_1202_i19_ in range(default__.abs(-580))])):
                            coll50_[d_1203_v178_] = False
                    return _dafny.Map(coll50_)
                nw198_[int(16)] = D9_DC30(d_1193_v172_, (d_1181_cf25_)[default__.safeIndex(241, (d_1181_cf25_).length(0))], len(iife146_()
), _dafny.Set({len((d_1191_v170_).f21)}), (d_1192_v171_).f20)
                nw198_[int(17)] = d_1194_v173_
                nw198_[int(18)] = d_1194_v173_
                nw198_[int(19)] = d_1194_v173_
                nw198_[int(20)] = d_1194_v173_
                d_1199_v179_ = nw198_
                d_1204_v180_: _dafny.Map
                d_1204_v180_ = _dafny.Map({d_1199_v179_: (d_1181_cf25_)[default__.safeIndex(241, (d_1181_cf25_).length(0))]})
                rhs188_ = ((d_1204_v180_)[d_1199_v179_] if (d_1199_v179_) in (d_1204_v180_) else ((D1_DC4((d_1191_v170_).f20, (d_1191_v170_).f20)).cf4) not in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_1182_cf24_: d_1182_cf24_})) for d_1205_i20_ in range(default__.abs(588))])))
                rhs189_ = (d_1192_v171_).fm2(False, d_1182_cf24_, globalState)
                r2 = rhs188_
                d_1182_cf24_ = rhs189_
            elif True:
                (self).f14 = ((p1) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qd"))))) != (d_1183_cf23_)
                d_1206_v181_: D4
                d_1206_v181_ = D4_DC14(self.f14, (self).f19)
                d_1207_v182_: D4
                d_1207_v182_ = D4_DC15(d_1206_v181_)
                d_1208_v183_: _dafny.Map
                d_1208_v183_ = _dafny.Map({d_1207_v182_: d_1183_cf23_})
                d_1209_v184_: _dafny.Set
                d_1209_v184_ = _dafny.Set({self.f14, self.f14})
                d_1210_v185_: _dafny.Map
                d_1210_v185_ = _dafny.Map({(d_1181_cf25_)[default__.safeIndex(241, (d_1181_cf25_).length(0))]: d_1209_v184_})
                d_1211_v186_: _dafny.Map
                d_1211_v186_ = _dafny.Map({p1: (d_1208_v183_).set(D4_DC15(d_1206_v181_), len(d_1210_v185_))})
                d_1211_v186_ = (d_1211_v186_).set(d_1183_cf23_, _dafny.Map({d_1207_v182_: p1}))
                d_1212_v187_: _dafny.Seq
                d_1212_v187_ = _dafny.SeqWithoutIsStrInference([(d_1181_cf25_)[default__.safeIndex(241, (d_1181_cf25_).length(0))]])
                d_1213_v188_: _dafny.MultiSet
                d_1213_v188_ = _dafny.MultiSet([(d_1035_v63_) + ((self).fm3(len(d_1212_v187_), d_1183_cf23_, not((self).f19), d_1183_cf23_, globalState))])
                d_1213_v188_ = _dafny.MultiSet([d_1184_cf22_])
                d_1214_v189_: _dafny.Seq
                d_1214_v189_ = _dafny.SeqWithoutIsStrInference([len(d_1035_v63_), p1])
                d_1215_v190_: _dafny.Map
                d_1215_v190_ = _dafny.Map({d_1181_cf25_: len(d_1214_v189_)})
                d_1216_v191_: str
                d_1216_v191_ = _dafny.CodePoint('y')
                d_1217_v192_: _dafny.Map
                d_1217_v192_ = _dafny.Map({d_1216_v191_: d_1183_cf23_})
                d_1218_v193_: _dafny.Set
                d_1218_v193_ = _dafny.Set({d_1214_v189_})
                d_1219_v194_: _dafny.Seq
                d_1219_v194_ = _dafny.SeqWithoutIsStrInference([d_1218_v193_])
                index217_ = default__.safeIndex(241, (d_1181_cf25_).length(0))
                (d_1181_cf25_)[index217_] = not (default__.fm23(p1, ((d_1185_v166_)[len(d_1215_v190_)] if (len(d_1215_v190_)) in (d_1185_v166_) else self.f14), d_1217_v192_, (d_1219_v194_)[default__.safeIndex(940, len(d_1219_v194_))], globalState)) or ((d_1181_cf25_)[default__.safeIndex(241, (d_1181_cf25_).length(0))])
                d_1220_v195_: C2
                nw199_ = C2()
                nw199_.ctor__((d_1185_v166_) | (d_1185_v166_))
                d_1220_v195_ = nw199_
            d_1221_v196_: _dafny.Map
            d_1221_v196_ = _dafny.Map({d_1182_cf24_: (self).f15})
            d_1222_v197_: C3
            nw200_ = C3()
            nw200_.ctor__(p1, (d_1184_cf22_) + (d_1184_cf22_), d_1182_cf24_, ((d_1221_v196_)[(d_1181_cf25_)[default__.safeIndex(241, (d_1181_cf25_).length(0))]] if ((d_1181_cf25_)[default__.safeIndex(241, (d_1181_cf25_).length(0))]) in (d_1221_v196_) else (self).f15))
            d_1222_v197_ = nw200_
        elif source14_.is_DC14:
            d_1223___mcc_h12_ = source14_.cf26
            d_1224___mcc_h13_ = source14_.cf27
            d_1225_cf27_ = d_1224___mcc_h13_
            d_1226_cf26_ = d_1223___mcc_h12_
            d_1227_v198_: _dafny.MultiSet
            d_1227_v198_ = _dafny.MultiSet([default__.fm0(self.f14, p1, len(d_1035_v63_), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1228_i21_ in range(default__.abs(-852))]))]), globalState), len(d_1035_v63_), 869])
            d_1229_v199_: _dafny.Seq
            d_1229_v199_ = _dafny.SeqWithoutIsStrInference([d_1227_v198_, _dafny.MultiSet([p1])])
            r1 = default__.fm0(d_1226_cf26_, len(d_1035_v63_), p1, ((d_1229_v199_)[default__.safeIndex(p1, len(d_1229_v199_))]).intersection(d_1227_v198_), globalState)
            d_1230_v200_: _dafny.Seq
            d_1230_v200_ = _dafny.SeqWithoutIsStrInference([(self).f15])
            d_1231_v201_: _dafny.Seq
            d_1231_v201_ = _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([p1, p1]))), p1])
            (self).m7((d_1230_v200_) + (_dafny.SeqWithoutIsStrInference([(self).f15 for d_1232_i22_ in range(default__.abs(770))])), (p1) >= (len(d_1231_v201_)), -452, d_1225_cf27_, globalState)
            (self).f14 = (d_1035_v63_) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vlgqwvqgo"))) + (d_1035_v63_))
            (globalState).f3 = p1
        elif source14_.is_DC12:
            d_1233___mcc_h14_ = source14_.cf21
            d_1234_cf21_ = d_1233___mcc_h14_
            d_1235_v202_: _dafny.Set
            d_1235_v202_ = _dafny.Set({self.f14})
            d_1236_v203_: _dafny.Seq
            d_1236_v203_ = _dafny.SeqWithoutIsStrInference([len(d_1235_v202_), p1])
            d_1237_v204_: D1
            d_1237_v204_ = D1_DC3(len(d_1035_v63_))
            d_1238_v205_: _dafny.Seq
            d_1238_v205_ = d_1236_v203_
            d_1239_v206_: _dafny.Map
            d_1239_v206_ = _dafny.Map({p1: (self).f19})
            d_1240_v207_: C2
            nw201_ = C2()
            nw201_.ctor__(d_1239_v206_)
            d_1240_v207_ = nw201_
            d_1241_v208_: _dafny.Map
            d_1241_v208_ = _dafny.Map({d_1240_v207_: d_1236_v203_})
            d_1242_v209_: _dafny.Seq
            d_1242_v209_ = _dafny.SeqWithoutIsStrInference([(self).f19, self.f14])
            d_1243_v210_: _dafny.Array
            nw202_ = _dafny.Array(None, 28)
            nw202_[int(0)] = _dafny.SeqWithoutIsStrInference([681 for d_1244_i23_ in range(default__.abs(955))])
            nw202_[int(1)] = d_1236_v203_
            nw202_[int(2)] = d_1236_v203_
            nw202_[int(3)] = _dafny.SeqWithoutIsStrInference([p1])
            nw202_[int(4)] = default__.fm21(d_1237_v204_, self.f14, default__.fm40(globalState), globalState)
            nw202_[int(5)] = d_1236_v203_
            nw202_[int(6)] = d_1236_v203_
            nw202_[int(7)] = d_1236_v203_
            nw202_[int(8)] = d_1236_v203_
            nw202_[int(9)] = d_1236_v203_
            nw202_[int(10)] = d_1236_v203_
            nw202_[int(11)] = d_1236_v203_
            nw202_[int(12)] = d_1236_v203_
            nw202_[int(13)] = d_1236_v203_
            nw202_[int(14)] = d_1236_v203_
            nw202_[int(15)] = _dafny.SeqWithoutIsStrInference([189])
            nw202_[int(16)] = (d_1238_v205_)
            nw202_[int(17)] = ((d_1241_v208_)[d_1240_v207_] if (d_1240_v207_) in (d_1241_v208_) else d_1236_v203_)
            nw202_[int(18)] = (d_1236_v203_ if (self).f19 else d_1236_v203_)
            nw202_[int(19)] = d_1236_v203_
            nw202_[int(20)] = d_1236_v203_
            nw202_[int(21)] = (d_1236_v203_) + (d_1236_v203_)
            nw202_[int(22)] = _dafny.SeqWithoutIsStrInference([p1, p1])
            nw202_[int(23)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not((self).f19), (self).f19, self.f14, (self).f19, (self).f19])) for d_1245_i24_ in range(default__.abs(-99))])
            nw202_[int(24)] = (d_1236_v203_) + ((d_1236_v203_).set(default__.safeIndex(len(d_1242_v209_), len(d_1236_v203_)), 790))
            nw202_[int(25)] = d_1236_v203_
            nw202_[int(26)] = d_1236_v203_
            nw202_[int(27)] = d_1236_v203_
            d_1243_v210_ = nw202_
            d_1243_v210_ = d_1243_v210_
            d_1246_v211_: _dafny.Map
            d_1246_v211_ = _dafny.Map({p1: d_1235_v202_})
            d_1247_v212_: _dafny.Array
            nw203_ = _dafny.Array(None, 9)
            nw203_[int(0)] = d_1235_v202_
            nw203_[int(1)] = d_1235_v202_
            nw203_[int(2)] = default__.fm41(len(d_1236_v203_), p1, globalState)
            nw203_[int(3)] = ((d_1246_v211_)[996] if (996) in (d_1246_v211_) else d_1235_v202_)
            nw203_[int(4)] = (d_1235_v202_).intersection(d_1235_v202_)
            nw203_[int(5)] = d_1235_v202_
            nw203_[int(6)] = (d_1235_v202_) | (_dafny.Set({False, False}))
            nw203_[int(7)] = d_1235_v202_
            nw203_[int(8)] = d_1235_v202_
            d_1247_v212_ = nw203_
            d_1248_v213_: _dafny.MultiSet
            d_1248_v213_ = _dafny.MultiSet([(0) - (len(d_1035_v63_)), p1])
            d_1249_v215_: _dafny.Set
            def iife147_():
                coll51_ = _dafny.Map()
                compr_51_: int
                for compr_51_ in _dafny.IntegerRange(210, 972):
                    d_1250_v214_: int = compr_51_
                    if ((210) <= (d_1250_v214_)) and ((d_1250_v214_) < (972)):
                        coll51_[(d_1250_v214_) * (p1)] = True
                return _dafny.Map(coll51_)
            d_1249_v215_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({366: len(iife147_()
)})) for d_1251_i25_ in range(default__.abs(518))])})
            index218_ = default__.safeIndex(632, (d_1247_v212_).length(0))
            (d_1247_v212_)[index218_] = _dafny.Set({not(True), default__.fm23(default__.fm0(False, 279, p1, _dafny.MultiSet([((d_1248_v213_)[p1] if (p1) in (d_1248_v213_) else (0) - (p1)), len(d_1242_v209_), p1]), globalState), self.f14, default__.fm24(p1, (self).f19, False, p1, globalState), d_1249_v215_, globalState), (self).f19, (self).f19, (self).f19})
            index219_ = default__.safeIndex(632, (d_1247_v212_).length(0))
            rhs190_ = p1
            rhs191_ = not(not(((d_1035_v63_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhxiqa")))) <= (d_1035_v63_)))
            rhs192_ = (_dafny.Set({not(self.f14), self.f14, self.f14, (self).f19, (self).f19}) if self.f14 else (d_1235_v202_) - (d_1235_v202_))
            lhs159_ = globalState
            lhs160_ = d_1247_v212_
            lhs161_ = default__.safeIndex(632, (d_1247_v212_).length(0))
            lhs159_.f3 = rhs190_
            r2 = rhs191_
            lhs160_[lhs161_] = rhs192_
            d_1252_v216_: _dafny.Array
            nw204_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_1252_v216_ = nw204_
            d_1253_v217_: C3
            nw205_ = C3()
            nw205_.ctor__(6, d_1035_v63_, (self).f19, (self).f15)
            d_1253_v217_ = nw205_
            d_1254_v218_: _dafny.Map
            d_1254_v218_ = _dafny.Map({d_1252_v216_: d_1253_v217_})
            d_1254_v218_ = (d_1254_v218_).set((d_1252_v216_ if (self).f19 else d_1252_v216_), d_1253_v217_)
            d_1255_v219_: _dafny.Map
            d_1255_v219_ = _dafny.Map({_dafny.Set({(d_1253_v217_).f20}): (self).f19})
            d_1256_v220_: _dafny.Array
            def lambda56_(d_1257_p1_, d_1258_v217_):
                def lambda57_(d_1259_i26_):
                    return (d_1259_i26_) - ((D1_DC4((0) - (d_1257_p1_), (0) - ((0) - ((d_1258_v217_).f20)))).cf4)

                return lambda57_

            init32_ = lambda56_(p1, d_1253_v217_)
            nw206_ = _dafny.Array(None, 14)
            for i0_32_ in range(nw206_.length(0)):
                nw206_[i0_32_] = init32_(i0_32_)
            d_1256_v220_ = nw206_
            d_1260_v221_: _dafny.Map
            d_1260_v221_ = _dafny.Map({d_1255_v219_: d_1256_v220_})
            d_1260_v221_ = (d_1260_v221_).set(d_1255_v219_, d_1256_v220_)
        elif True:
            d_1261___mcc_h15_ = source14_.cf28
            d_1262_cf28_ = d_1261___mcc_h15_
            if (p1) > (595):
                index220_ = default__.safeIndex(853, (d_1171_v163_).length(0))
                (d_1171_v163_)[index220_] = not (self.f14) or ((self).f19)
                index221_ = default__.safeIndex(853, (d_1171_v163_).length(0))
                (d_1171_v163_)[index221_] = (self).f19
                d_1263_v222_: str
                d_1263_v222_ = _dafny.CodePoint('r')
                d_1264_v223_: _dafny.Map
                d_1264_v223_ = _dafny.Map({d_1263_v222_: (d_1171_v163_)[default__.safeIndex(853, (d_1171_v163_).length(0))]})
                d_1264_v223_ = (d_1264_v223_).set(d_1263_v222_, (self).f19)
                d_1265_v224_: _dafny.Map
                d_1265_v224_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([not((self).f19)]): len(d_1172_v164_)})
                d_1266_v226_: _dafny.Seq
                d_1266_v226_ = _dafny.SeqWithoutIsStrInference([(d_1171_v163_)[default__.safeIndex(853, (d_1171_v163_).length(0))]])
                d_1267_v227_: _dafny.Map
                d_1267_v227_ = _dafny.Map({p1: d_1035_v63_})
                d_1268_v228_: D8
                d_1268_v228_ = D8_DC26((d_1171_v163_)[default__.safeIndex(853, (d_1171_v163_).length(0))], d_1267_v227_)
                d_1269_v229_: _dafny.MultiSet
                d_1269_v229_ = _dafny.MultiSet([default__.fm42(False, not((d_1171_v163_)[default__.safeIndex(853, (d_1171_v163_).length(0))]), 265, (d_1171_v163_)[default__.safeIndex(853, (d_1171_v163_).length(0))], globalState), d_1268_v228_])
                d_1270_v230_: D3
                d_1270_v230_ = D3_DC11(p1, p1, self.f14, -278)
                d_1271_v231_: _dafny.MultiSet
                d_1271_v231_ = _dafny.MultiSet([(d_1270_v230_).cf18, len(d_1035_v63_)])
                d_1272_v232_: _dafny.Set
                d_1272_v232_ = _dafny.Set({d_1271_v231_, d_1271_v231_})
                index222_ = default__.safeIndex(853, (d_1171_v163_).length(0))
                def iife148_():
                    coll52_ = _dafny.Map()
                    compr_52_: _dafny.Seq
                    for compr_52_ in (_dafny.SeqWithoutIsStrInference([d_1266_v226_])).Elements:
                        d_1273_v225_: _dafny.Seq = compr_52_
                        if (d_1273_v225_) in (_dafny.SeqWithoutIsStrInference([d_1266_v226_])):
                            coll52_[d_1273_v225_] = p1
                    return _dafny.Map(coll52_)
                rhs193_ = iife148_()

                rhs194_ = not((d_1269_v229_).ispropersubset(d_1269_v229_))
                rhs195_ = (d_1171_v163_)[default__.safeIndex(853, (d_1171_v163_).length(0))]
                rhs196_ = (d_1272_v232_).isdisjoint(d_1272_v232_)
                lhs162_ = globalState
                lhs163_ = self
                lhs164_ = d_1171_v163_
                lhs165_ = default__.safeIndex(853, (d_1171_v163_).length(0))
                d_1265_v224_ = rhs193_
                lhs162_.f9 = rhs194_
                lhs163_.f14 = rhs195_
                lhs164_[lhs165_] = rhs196_
                d_1274_v233_: _dafny.Array
                nw207_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_1274_v233_ = nw207_
                d_1275_v234_: _dafny.Array
                def lambda58_(d_1276_p1_):
                    def lambda59_(d_1277_i27_):
                        return default__.safeDivisionInt(d_1277_i27_, d_1276_p1_)

                    return lambda59_

                init33_ = lambda58_(p1)
                nw208_ = _dafny.Array(None, 28)
                for i0_33_ in range(nw208_.length(0)):
                    nw208_[i0_33_] = init33_(i0_33_)
                d_1275_v234_ = nw208_
                index223_ = default__.safeIndex(844, (d_1274_v233_).length(0))
                (d_1274_v233_)[index223_] = d_1275_v234_
                d_1278_v235_: _dafny.Map
                d_1278_v235_ = _dafny.Map({p1: d_1263_v222_})
                d_1279_v236_: _dafny.Map
                d_1279_v236_ = _dafny.Map({self.f14: d_1263_v222_})
                index224_ = default__.safeIndex(844, (d_1274_v233_).length(0))
                rhs197_ = p1
                rhs198_ = p1
                rhs199_ = d_1275_v234_
                rhs200_ = ((self).fm3((0) - (p1), p1, not(((d_1264_v223_)[d_1263_v222_] if (d_1263_v222_) in (d_1264_v223_) else self.f14)), default__.safeModuloInt(len(d_1278_v235_), p1), globalState)).set(default__.safeIndex(p1, len((self).fm3((0) - (p1), p1, not(((d_1264_v223_)[d_1263_v222_] if (d_1263_v222_) in (d_1264_v223_) else self.f14)), default__.safeModuloInt(len(d_1278_v235_), p1), globalState))), ((d_1279_v236_)[not(self.f14)] if (not(self.f14)) in (d_1279_v236_) else d_1263_v222_))
                lhs166_ = globalState
                lhs167_ = globalState
                lhs168_ = d_1274_v233_
                lhs169_ = default__.safeIndex(844, (d_1274_v233_).length(0))
                lhs166_.f3 = rhs197_
                lhs167_.f3 = rhs198_
                lhs168_[lhs169_] = rhs199_
                d_1035_v63_ = rhs200_
                d_1280_v237_: _dafny.Set
                d_1280_v237_ = _dafny.Set({len(_dafny.Set({(self).f19, (self).f19, not((d_1171_v163_)[default__.safeIndex(853, (d_1171_v163_).length(0))])})), (d_1270_v230_).cf18, p1})
                d_1280_v237_ = d_1280_v237_
            elif True:
                d_1281_v238_: _dafny.Array
                def lambda60_(d_1282_p1_):
                    def lambda61_(d_1283_i28_):
                        return _dafny.SeqWithoutIsStrInference([d_1282_p1_])

                    return lambda61_

                init34_ = lambda60_(p1)
                nw209_ = _dafny.Array(None, 14)
                for i0_34_ in range(nw209_.length(0)):
                    nw209_[i0_34_] = init34_(i0_34_)
                d_1281_v238_ = nw209_
                d_1284_v239_: _dafny.Map
                d_1284_v239_ = _dafny.Map({_dafny.CodePoint('s'): p1})
                d_1285_v240_: _dafny.Map
                d_1285_v240_ = _dafny.Map({p1: d_1284_v239_})
                d_1286_v241_: _dafny.Map
                d_1286_v241_ = _dafny.Map({(self).f19: (self).fm2((self).f19, self.f14, globalState)})
                d_1287_v242_: _dafny.Seq
                d_1287_v242_ = _dafny.SeqWithoutIsStrInference([p1, len(((d_1285_v240_)[len(d_1286_v241_)] if (len(d_1286_v241_)) in (d_1285_v240_) else d_1284_v239_)), p1, p1, p1])
                index225_ = default__.safeIndex(696, (d_1281_v238_).length(0))
                (d_1281_v238_)[index225_] = (d_1287_v242_ if self.f14 else d_1287_v242_)
                index226_ = default__.safeIndex(696, (d_1281_v238_).length(0))
                (d_1281_v238_)[index226_] = d_1287_v242_
                d_1288_v243_: _dafny.Seq
                d_1288_v243_ = _dafny.SeqWithoutIsStrInference([d_1035_v63_])
                r1 = (default__.safeDivisionInt(p1, len(d_1035_v63_)) if (self).f19 else len(d_1288_v243_))
                d_1289_v244_: D4
                d_1289_v244_ = D4_DC12(d_1171_v163_)
                d_1290_v245_: _dafny.Array
                nw210_ = _dafny.Array(None, 4)
                nw210_[int(0)] = d_1289_v244_
                nw210_[int(1)] = d_1289_v244_
                nw210_[int(2)] = d_1289_v244_
                nw210_[int(3)] = d_1289_v244_
                d_1290_v245_ = nw210_
                d_1291_v246_: _dafny.Map
                d_1291_v246_ = _dafny.Map({d_1290_v245_: True})
                (globalState).f9 = (d_1290_v245_) not in (d_1291_v246_)
                d_1292_v247_: _dafny.Map
                d_1292_v247_ = _dafny.Map({len(d_1035_v63_): self.f14})
                (self).f14 = ((len(d_1292_v247_)) == (p1)) == ((self).f19)
                d_1293_v248_: _dafny.Seq
                d_1293_v248_ = _dafny.SeqWithoutIsStrInference([(self).f19])
                (self).f14 = not ((d_1293_v248_)[default__.safeIndex(364, len(d_1293_v248_))]) or ((self).f19)
            if not(True):
                d_1294_v249_: _dafny.Seq
                d_1294_v249_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_1295_v250_: _dafny.MultiSet
                d_1295_v250_ = _dafny.MultiSet([p1])
                d_1296_v251_: C3
                nw211_ = C3()
                nw211_.ctor__(default__.fm0((self).f19, len(d_1294_v249_), p1, d_1295_v250_, globalState), d_1035_v63_, (p1) >= (p1), (self).f15)
                d_1296_v251_ = nw211_
                d_1297_v252_: _dafny.Seq
                d_1297_v252_ = _dafny.SeqWithoutIsStrInference([d_1295_v250_, d_1295_v250_, d_1295_v250_, d_1295_v250_])
                (self).m7(default__.fm43((d_1296_v251_).f20, p1, len(_dafny.SeqWithoutIsStrInference([(0) - (default__.fm0(not((self).f19), (d_1296_v251_).f20, p1, d_1295_v250_, globalState))])), globalState), self.f14, ((d_1295_v250_).intersection((d_1297_v252_)[default__.safeIndex(p1, len(d_1297_v252_))])).cardinality, self.f14, globalState)
                d_1298_v253_: D2
                d_1298_v253_ = D2_DC8(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1299_i29_ in range(default__.abs(-565))])), default__.safeDivisionInt((d_1296_v251_).f20, p1), p1)
                d_1300_v254_: _dafny.Array
                nw212_ = _dafny.Array(int(0), 20)
                d_1300_v254_ = nw212_
                index227_ = default__.safeIndex(298, (d_1300_v254_).length(0))
                (d_1300_v254_)[index227_] = (0) - (default__.safeModuloInt((d_1296_v251_).f20, p1))
                index228_ = default__.safeIndex(605, (d_1300_v254_).length(0))
                (d_1300_v254_)[index228_] = (len(d_1294_v249_)) - (p1)
                index229_ = default__.safeIndex(902, (d_1171_v163_).length(0))
                (d_1171_v163_)[index229_] = True
                d_1301_v255_: _dafny.Set
                d_1301_v255_ = _dafny.Set({d_1300_v254_})
                index230_ = default__.safeIndex(298, (d_1300_v254_).length(0))
                index231_ = default__.safeIndex(605, (d_1300_v254_).length(0))
                index232_ = default__.safeIndex(902, (d_1171_v163_).length(0))
                rhs201_ = (self).f19
                rhs202_ = default__.fm44((self).f19, (d_1296_v251_).f20, (d_1296_v251_).f20, globalState)
                rhs203_ = (p1 if self.f14 else p1)
                rhs204_ = (d_1296_v251_).f20
                rhs205_ = ((d_1301_v255_) | (d_1301_v255_)) == (_dafny.Set({d_1300_v254_}))
                lhs170_ = d_1300_v254_
                lhs171_ = default__.safeIndex(298, (d_1300_v254_).length(0))
                lhs172_ = d_1300_v254_
                lhs173_ = default__.safeIndex(605, (d_1300_v254_).length(0))
                lhs174_ = d_1171_v163_
                lhs175_ = default__.safeIndex(902, (d_1171_v163_).length(0))
                r2 = rhs201_
                d_1298_v253_ = rhs202_
                lhs170_[lhs171_] = rhs203_
                lhs172_[lhs173_] = rhs204_
                lhs174_[lhs175_] = rhs205_
                d_1302_v256_: _dafny.Map
                d_1302_v256_ = _dafny.Map({True: 560})
                index233_ = default__.safeIndex(298, (d_1300_v254_).length(0))
                (d_1300_v254_)[index233_] = ((d_1296_v251_).f20) * (((d_1302_v256_)[self.f14] if (self.f14) in (d_1302_v256_) else (d_1296_v251_).f20))
                (globalState).f3 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owt")))
            elif True:
                d_1303_v257_: _dafny.Array
                nw213_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_1303_v257_ = nw213_
                d_1303_v257_ = d_1303_v257_
                d_1304_v258_: D8
                d_1304_v258_ = D8_DC26((self).f19, _dafny.Map({p1: d_1035_v63_}))
                d_1305_v259_: _dafny.Array
                nw214_ = _dafny.Array(int(0), 10)
                d_1305_v259_ = nw214_
                index234_ = default__.safeIndex(559, (d_1305_v259_).length(0))
                (d_1305_v259_)[index234_] = (p1 if (self).f19 else p1)
                d_1306_v260_: D4
                d_1306_v260_ = D4_DC14(self.f14, (self).fm2(self.f14, not(self.f14), globalState))
                d_1307_v261_: _dafny.Map
                d_1307_v261_ = _dafny.Map({(self).fm3(p1, p1, (d_1306_v260_).cf26, p1, globalState): d_1305_v259_})
                d_1308_v262_: _dafny.Seq
                d_1308_v262_ = _dafny.SeqWithoutIsStrInference([d_1035_v63_, d_1035_v63_, d_1035_v63_])
                d_1309_v263_: _dafny.Map
                d_1309_v263_ = _dafny.Map({p1: (d_1308_v262_)[default__.safeIndex(p1, len(d_1308_v262_))]})
                pat_let_tv37_ = d_1309_v263_
                index235_ = default__.safeIndex(559, (d_1305_v259_).length(0))
                def iife149_(_pat_let48_0):
                    def iife150_(d_1310_dt__update__tmp_h9_):
                        def iife151_(_pat_let49_0):
                            def iife152_(d_1311_dt__update_hcf47_h0_):
                                return D8_DC26((d_1310_dt__update__tmp_h9_).cf46, d_1311_dt__update_hcf47_h0_)
                            return iife152_(_pat_let49_0)
                        return iife151_(pat_let_tv37_)
                    return iife150_(_pat_let48_0)
                rhs206_ = p1
                rhs207_ = len(d_1307_v261_)
                rhs208_ = iife149_(d_1304_v258_)
                rhs209_ = default__.safeModuloInt((p1) * (((self).f15).cardinality), p1)
                lhs176_ = globalState
                lhs177_ = d_1305_v259_
                lhs178_ = default__.safeIndex(559, (d_1305_v259_).length(0))
                r0 = rhs206_
                lhs176_.f3 = rhs207_
                d_1304_v258_ = rhs208_
                lhs177_[lhs178_] = rhs209_
                d_1312_v264_: D2
                d_1312_v264_ = D2_DC9((self).f19, (d_1305_v259_)[default__.safeIndex(559, (d_1305_v259_).length(0))], _dafny.CodePoint('n'))
                r0 = (d_1312_v264_).cf14
                r2 = not(self.f14)
                d_1313_v265_: _dafny.Map
                d_1313_v265_ = _dafny.Map({(d_1305_v259_)[default__.safeIndex(559, (d_1305_v259_).length(0))]: len(p0)})
                d_1314_v266_: _dafny.Map
                d_1314_v266_ = _dafny.Map({d_1313_v265_: (d_1305_v259_)[default__.safeIndex(559, (d_1305_v259_).length(0))]})
                d_1314_v266_ = (d_1314_v266_).set((d_1313_v265_) | (d_1313_v265_), (d_1305_v259_)[default__.safeIndex(559, (d_1305_v259_).length(0))])
            (self).f14 = not(self.f14)
            if (self).f19:
                d_1315_v267_: _dafny.Map
                d_1315_v267_ = _dafny.Map({p1: self.f14})
                d_1316_v268_: C2
                nw215_ = C2()
                nw215_.ctor__((d_1315_v267_).set(p1, (self).f19))
                d_1316_v268_ = nw215_
                d_1317_v269_: _dafny.Array
                nw216_ = _dafny.Array(_dafny.CodePoint('D'), 12)
                d_1317_v269_ = nw216_
                d_1318_v270_: str
                d_1318_v270_ = _dafny.CodePoint('a')
                index236_ = default__.safeIndex(236, (d_1317_v269_).length(0))
                (d_1317_v269_)[index236_] = d_1318_v270_
                d_1319_v271_: _dafny.Set
                d_1319_v271_ = _dafny.Set({(self).f15, (self).f15})
                index237_ = default__.safeIndex(236, (d_1317_v269_).length(0))
                rhs210_ = (241) < ((0) - ((0) - (p1)))
                rhs211_ = (d_1319_v271_).issubset(d_1319_v271_)
                rhs212_ = (d_1035_v63_)[default__.safeIndex(p1, len(d_1035_v63_))]
                lhs179_ = globalState
                lhs180_ = globalState
                lhs181_ = d_1317_v269_
                lhs182_ = default__.safeIndex(236, (d_1317_v269_).length(0))
                lhs179_.f9 = rhs210_
                lhs180_.f9 = rhs211_
                lhs181_[lhs182_] = rhs212_
                d_1320_v272_: _dafny.MultiSet
                d_1320_v272_ = _dafny.MultiSet([p1, p1])
                d_1321_v273_: _dafny.Seq
                d_1321_v273_ = _dafny.SeqWithoutIsStrInference([(d_1320_v272_).cardinality])
                d_1322_v274_: _dafny.Set
                d_1322_v274_ = _dafny.Set({d_1321_v273_, _dafny.SeqWithoutIsStrInference([p1]), d_1321_v273_})
                d_1323_v275_: _dafny.Map
                d_1323_v275_ = _dafny.Map({False: (self).f19})
                pat_let_tv38_ = d_1323_v275_
                d_1324_v276_: _dafny.Map
                def iife153_(_pat_let50_0):
                    def iife154_(d_1325_dt__update__tmp_h10_):
                        def iife155_(_pat_let51_0):
                            def iife156_(d_1326_dt__update_hcf16_h0_):
                                return D3_DC10(d_1326_dt__update_hcf16_h0_)
                            return iife156_(_pat_let51_0)
                        return iife155_(pat_let_tv38_)
                    return iife154_(_pat_let50_0)
                d_1324_v276_ = _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwmccp")))): (len(_dafny.Set({p1, p1})) if self.f14 else len((iife153_(default__.fm45(p1, default__.fm23(p1, True, _dafny.Map({d_1318_v270_: p1}), d_1322_v274_, globalState), p1, globalState))).cf16))})
                d_1324_v276_ = (d_1324_v276_).set(p1, len(d_1035_v63_))
                (globalState).f3 = p1
                (globalState).f9 = (d_1320_v272_).isdisjoint((d_1320_v272_).intersection(d_1320_v272_))
            elif True:
                (globalState).f3 = p1
                d_1327_v277_: _dafny.Map
                d_1327_v277_ = _dafny.Map({False: (p1) * (p1)})
                d_1327_v277_ = d_1327_v277_
                r1 = p1
                d_1035_v63_ = (d_1035_v63_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1328_i30_ in range(default__.abs(-916))]))
                d_1329_v278_: _dafny.MultiSet
                d_1329_v278_ = _dafny.MultiSet([p1, p1])
                (globalState).f3 = default__.fm0(not ((self).f19) or (self.f14), p1, (p1 if (self).f19 else p1), d_1329_v278_, globalState)
        r0 = p1
        r1 = default__.safeDivisionInt(p1, p1)
        r2 = self.f14
        return r0, r1, r2

    def m6(self, globalState):
        if self.f14:
            d_1330_v0_: _dafny.Array
            def lambda62_(d_1331_i0_):
                return default__.safeDivisionInt(d_1331_i0_, len(_dafny.SeqWithoutIsStrInference([self.f14])))

            init35_ = lambda62_
            nw217_ = _dafny.Array(None, 22)
            for i0_35_ in range(nw217_.length(0)):
                nw217_[i0_35_] = init35_(i0_35_)
            d_1330_v0_ = nw217_
            d_1332_v1_: int
            d_1332_v1_ = 835
            index238_ = default__.safeIndex(612, (d_1330_v0_).length(0))
            (d_1330_v0_)[index238_] = (d_1332_v1_) * (698)
            d_1333_v2_: _dafny.Seq
            d_1333_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dchmesu"))
            d_1334_v3_: _dafny.Seq
            d_1334_v3_ = _dafny.SeqWithoutIsStrInference([d_1332_v1_])
            d_1335_v4_: _dafny.Seq
            d_1335_v4_ = _dafny.SeqWithoutIsStrInference([self.f14])
            d_1336_v5_: _dafny.Set
            d_1336_v5_ = _dafny.Set({d_1332_v1_})
            index239_ = default__.safeIndex(612, (d_1330_v0_).length(0))
            rhs213_ = (0) - ((0) - ((d_1334_v3_)[default__.safeIndex(d_1332_v1_, len(d_1334_v3_))]))
            rhs214_ = (d_1333_v2_) + (d_1333_v2_)
            rhs215_ = (d_1336_v5_).ispropersubset(_dafny.Set({d_1332_v1_, (0) - ((0) - ((0) - (d_1332_v1_))), len(d_1335_v4_)}))
            rhs216_ = 711
            rhs217_ = default__.safeModuloInt(d_1332_v1_, d_1332_v1_)
            lhs183_ = d_1330_v0_
            lhs184_ = default__.safeIndex(612, (d_1330_v0_).length(0))
            lhs185_ = globalState
            lhs186_ = globalState
            lhs187_ = globalState
            lhs183_[lhs184_] = rhs213_
            d_1333_v2_ = rhs214_
            lhs185_.f9 = rhs215_
            lhs186_.f3 = rhs216_
            lhs187_.f3 = rhs217_
            d_1337_v6_: _dafny.Array
            def lambda63_(d_1338_i1_):
                return ((self).f19) in (_dafny.Map({self.f14: self.f14}))

            init36_ = lambda63_
            nw218_ = _dafny.Array(None, 29)
            for i0_36_ in range(nw218_.length(0)):
                nw218_[i0_36_] = init36_(i0_36_)
            d_1337_v6_ = nw218_
            (globalState).f2 = d_1337_v6_
            index240_ = default__.safeIndex(909, (d_1337_v6_).length(0))
            (d_1337_v6_)[index240_] = self.f14
            index241_ = default__.safeIndex(909, (d_1337_v6_).length(0))
            (d_1337_v6_)[index241_] = self.f14
            d_1339_v7_: _dafny.Map
            d_1339_v7_ = _dafny.Map({(d_1330_v0_)[default__.safeIndex(612, (d_1330_v0_).length(0))]: False})
            d_1336_v5_ = (default__.fm22(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epm"))), _dafny.SeqWithoutIsStrInference([d_1339_v7_]), (self).f19, (self).f19, globalState)).intersection((d_1336_v5_).intersection(d_1336_v5_))
            d_1340_v8_: str
            d_1340_v8_ = _dafny.CodePoint('h')
            rhs218_ = (0) - (len(((d_1333_v2_).set(default__.safeIndex(d_1332_v1_, len(d_1333_v2_)), d_1340_v8_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ob")))))
            rhs219_ = ((d_1330_v0_)[default__.safeIndex(612, (d_1330_v0_).length(0))] if (d_1337_v6_)[default__.safeIndex(909, (d_1337_v6_).length(0))] else len(_dafny.SeqWithoutIsStrInference([(d_1337_v6_)[default__.safeIndex(909, (d_1337_v6_).length(0))], (self).f19, self.f14])))
            rhs220_ = d_1330_v0_
            rhs221_ = d_1333_v2_
            lhs188_ = globalState
            d_1332_v1_ = rhs218_
            lhs188_.f3 = rhs219_
            d_1330_v0_ = rhs220_
            d_1333_v2_ = rhs221_
        elif True:
            d_1341_v9_: int
            d_1341_v9_ = 819
            (globalState).f3 = d_1341_v9_
            d_1342_v10_: _dafny.Array
            def lambda64_(d_1343_i2_):
                return (self).f19

            init37_ = lambda64_
            nw219_ = _dafny.Array(None, 2)
            for i0_37_ in range(nw219_.length(0)):
                nw219_[i0_37_] = init37_(i0_37_)
            d_1342_v10_ = nw219_
            d_1344_v11_: _dafny.Seq
            d_1344_v11_ = _dafny.SeqWithoutIsStrInference([d_1342_v10_, d_1342_v10_, d_1342_v10_])
            d_1345_v12_: _dafny.MultiSet
            d_1345_v12_ = _dafny.MultiSet([(0) - (d_1341_v9_)])
            (globalState).f2 = (d_1344_v11_)[default__.safeIndex(default__.fm0(self.f14, (0) - (d_1341_v9_), d_1341_v9_, d_1345_v12_, globalState), len(d_1344_v11_))]
            d_1346_v13_: _dafny.Array
            nw220_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_1346_v13_ = nw220_
            nw221_ = _dafny.Array(_dafny.Array(None, 0), 8)
            d_1346_v13_ = nw221_
            index242_ = default__.safeIndex(630, (d_1342_v10_).length(0))
            (d_1342_v10_)[index242_] = self.f14
            index243_ = default__.safeIndex(630, (d_1342_v10_).length(0))
            (d_1342_v10_)[index243_] = ((self).f15).ispropersubset((self).f15)
            (self).f14 = (d_1342_v10_)[default__.safeIndex(630, (d_1342_v10_).length(0))]
        d_1347_v14_: int
        d_1347_v14_ = 202
        hi3_ = d_1347_v14_
        for d_1348_i3_ in range(d_1347_v14_, hi3_):
            d_1349_v15_: _dafny.Seq
            d_1349_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmbu"))
            d_1350_v16_: _dafny.MultiSet
            d_1350_v16_ = _dafny.MultiSet([(self).f19])
            d_1351_v17_: D8
            d_1351_v17_ = D8_DC25()
            rhs222_ = d_1349_v15_
            rhs223_ = d_1350_v16_
            rhs224_ = self.f14
            rhs225_ = d_1351_v17_
            lhs189_ = globalState
            d_1349_v15_ = rhs222_
            d_1350_v16_ = rhs223_
            lhs189_.f9 = rhs224_
            d_1351_v17_ = rhs225_
            d_1352_v18_: _dafny.Array
            def lambda65_(d_1353_i3_, d_1354_v14_):
                def lambda66_(d_1355_i4_):
                    return not((d_1353_i3_) != (d_1354_v14_))

                return lambda66_

            init38_ = lambda65_(d_1348_i3_, d_1347_v14_)
            nw222_ = _dafny.Array(None, 25)
            for i0_38_ in range(nw222_.length(0)):
                nw222_[i0_38_] = init38_(i0_38_)
            d_1352_v18_ = nw222_
            index244_ = default__.safeIndex(997, (d_1352_v18_).length(0))
            (d_1352_v18_)[index244_] = (self).f19
            index245_ = default__.safeIndex(997, (d_1352_v18_).length(0))
            (d_1352_v18_)[index245_] = self.f14
            d_1356_v19_: _dafny.Seq
            d_1356_v19_ = _dafny.SeqWithoutIsStrInference([(d_1352_v18_)[default__.safeIndex(997, (d_1352_v18_).length(0))]])
            d_1357_v20_: _dafny.MultiSet
            d_1357_v20_ = _dafny.MultiSet([default__.fm38(globalState)])
            d_1347_v14_ = (self).fm17(d_1348_i3_, not((d_1352_v18_)[default__.safeIndex(997, (d_1352_v18_).length(0))]), (_dafny.MultiSet([d_1356_v19_])).intersection(d_1357_v20_), -899, globalState)
            (globalState).f9 = (d_1356_v19_)[default__.safeIndex(default__.safeDivisionInt(d_1348_i3_, d_1348_i3_), len(d_1356_v19_))]
        if (self).f19:
            d_1358_v21_: _dafny.Array
            def lambda67_(d_1359_i5_):
                return _dafny.Set({self.f14, self.f14})

            init39_ = lambda67_
            nw223_ = _dafny.Array(None, 24)
            for i0_39_ in range(nw223_.length(0)):
                nw223_[i0_39_] = init39_(i0_39_)
            d_1358_v21_ = nw223_
            d_1360_v22_: _dafny.Set
            d_1360_v22_ = _dafny.Set({self.f14})
            index246_ = default__.safeIndex(579, (d_1358_v21_).length(0))
            (d_1358_v21_)[index246_] = (d_1360_v22_) - (d_1360_v22_)
            index247_ = default__.safeIndex(579, (d_1358_v21_).length(0))
            (d_1358_v21_)[index247_] = d_1360_v22_
            if self.f14:
                d_1361_v23_: _dafny.Array
                def lambda68_(d_1362_v14_):
                    def lambda69_(d_1363_i6_):
                        return _dafny.SeqWithoutIsStrInference([(0) - (d_1362_v14_), d_1362_v14_, d_1362_v14_])

                    return lambda69_

                init40_ = lambda68_(d_1347_v14_)
                nw224_ = _dafny.Array(None, 2)
                for i0_40_ in range(nw224_.length(0)):
                    nw224_[i0_40_] = init40_(i0_40_)
                d_1361_v23_ = nw224_
                d_1364_v24_: _dafny.Array
                nw225_ = _dafny.Array(D0.default()(), 24)
                d_1364_v24_ = nw225_
                d_1365_v25_: D0
                d_1365_v25_ = D0_DC1((self).fm2(self.f14, True, globalState))
                d_1366_v26_: D0
                d_1366_v26_ = D0_DC2(d_1365_v25_)
                index248_ = default__.safeIndex(618, (d_1364_v24_).length(0))
                (d_1364_v24_)[index248_] = d_1366_v26_
                d_1367_v27_: _dafny.Array
                def lambda70_(d_1368_i7_):
                    return D0_DC1(self.f14)

                init41_ = lambda70_
                nw226_ = _dafny.Array(None, 25)
                for i0_41_ in range(nw226_.length(0)):
                    nw226_[i0_41_] = init41_(i0_41_)
                d_1367_v27_ = nw226_
                d_1369_v28_: _dafny.Map
                d_1369_v28_ = _dafny.Map({d_1347_v14_: _dafny.Set({self.f14, (self).f19})})
                index249_ = default__.safeIndex(462, (d_1367_v27_).length(0))
                (d_1367_v27_)[index249_] = default__.fm37(_dafny.Set({d_1347_v14_}), d_1369_v28_, self.f14, (self).f19, globalState)
                d_1370_v29_: _dafny.Array
                nw227_ = _dafny.Array(False, 7)
                d_1370_v29_ = nw227_
                index250_ = default__.safeIndex(222, (d_1370_v29_).length(0))
                (d_1370_v29_)[index250_] = not((self).f19)
                d_1371_v30_: _dafny.Set
                d_1371_v30_ = _dafny.Set({d_1347_v14_})
                d_1372_v31_: _dafny.Seq
                d_1372_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pw"))
                d_1373_v32_: _dafny.Map
                d_1373_v32_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1374_i8_ in range(default__.abs(343))])): d_1372_v31_})
                d_1375_v33_: D8
                d_1375_v33_ = D8_DC26(self.f14, d_1373_v32_)
                pat_let_tv39_ = d_1375_v33_
                index251_ = default__.safeIndex(618, (d_1364_v24_).length(0))
                index252_ = default__.safeIndex(462, (d_1367_v27_).length(0))
                index253_ = default__.safeIndex(222, (d_1370_v29_).length(0))
                def iife158_(_pat_let53_0):
                    def iife159_(d_1376_dt__update__tmp_h0_):
                        def iife160_(_pat_let54_0):
                            def iife161_(d_1377_dt__update_hcf1_h0_):
                                return D0_DC1(d_1377_dt__update_hcf1_h0_)
                            return iife161_(_pat_let54_0)
                        return iife160_((pat_let_tv39_).cf46)
                    return iife159_(_pat_let53_0)
                def iife157_(_pat_let52_0):
                    def iife162_(d_1378_dt__update__tmp_h1_):
                        def iife163_(_pat_let55_0):
                            def iife164_(d_1379_dt__update_hcf1_h1_):
                                return D0_DC1(d_1379_dt__update_hcf1_h1_)
                            return iife164_(_pat_let55_0)
                        return iife163_(not((self).f19))
                    return iife162_(_pat_let52_0)
                rhs226_ = d_1361_v23_
                rhs227_ = d_1366_v26_
                rhs228_ = (d_1371_v30_).isdisjoint(d_1371_v30_)
                rhs229_ = iife157_(iife158_(D0_DC1((self).fm2(not((self).f19), self.f14, globalState))))
                rhs230_ = self.f14
                lhs190_ = d_1364_v24_
                lhs191_ = default__.safeIndex(618, (d_1364_v24_).length(0))
                lhs192_ = globalState
                lhs193_ = d_1367_v27_
                lhs194_ = default__.safeIndex(462, (d_1367_v27_).length(0))
                lhs195_ = d_1370_v29_
                lhs196_ = default__.safeIndex(222, (d_1370_v29_).length(0))
                d_1361_v23_ = rhs226_
                lhs190_[lhs191_] = rhs227_
                lhs192_.f9 = rhs228_
                lhs193_[lhs194_] = rhs229_
                lhs195_[lhs196_] = rhs230_
                (globalState).f2 = d_1370_v29_
                (self).f14 = not((d_1370_v29_)[default__.safeIndex(222, (d_1370_v29_).length(0))])
                d_1380_v34_: C1
                nw228_ = C1()
                nw228_.ctor__()
                d_1380_v34_ = nw228_
                d_1381_v35_: _dafny.MultiSet
                d_1381_v35_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwcabk")))])
                d_1382_v36_: _dafny.Seq
                d_1382_v36_ = _dafny.SeqWithoutIsStrInference([d_1381_v35_, d_1381_v35_])
                d_1382_v36_ = d_1382_v36_
            elif True:
                d_1383_v37_: _dafny.Seq
                d_1383_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nenovha"))
                d_1384_v38_: _dafny.Seq
                d_1384_v38_ = _dafny.SeqWithoutIsStrInference([self.f14, (self).f19, self.f14, self.f14, self.f14])
                d_1385_v39_: str
                d_1385_v39_ = _dafny.CodePoint('y')
                (globalState).f3 = len(((d_1383_v37_) + (d_1383_v37_)) + (((d_1383_v37_ if (d_1384_v38_)[default__.safeIndex(d_1347_v14_, len(d_1384_v38_))] else d_1383_v37_)).set(default__.safeIndex((0) - (d_1347_v14_), len((d_1383_v37_ if (d_1384_v38_)[default__.safeIndex(d_1347_v14_, len(d_1384_v38_))] else d_1383_v37_))), d_1385_v39_)))
                d_1386_v40_: _dafny.Seq
                d_1386_v40_ = _dafny.SeqWithoutIsStrInference([571])
                d_1387_v41_: _dafny.Seq
                d_1387_v41_ = d_1386_v40_
                d_1388_v42_: _dafny.Map
                d_1388_v42_ = _dafny.Map({d_1387_v41_: self.f14})
                d_1389_v43_: _dafny.Seq
                d_1389_v43_ = _dafny.SeqWithoutIsStrInference([d_1383_v37_, d_1383_v37_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvqxuld"))])
                d_1390_v44_: _dafny.Array
                nw229_ = _dafny.Array(None, 10)
                nw229_[int(0)] = (d_1383_v37_) + (d_1383_v37_)
                nw229_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebteocn"))
                nw229_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lronodbd"))
                nw229_[int(3)] = d_1383_v37_
                nw229_[int(4)] = d_1383_v37_
                nw229_[int(5)] = default__.fm27(d_1347_v14_, True, globalState)
                nw229_[int(6)] = (d_1383_v37_ if ((d_1388_v42_)[d_1387_v41_] if (d_1387_v41_) in (d_1388_v42_) else (self).f19) else d_1383_v37_)
                nw229_[int(7)] = (d_1389_v43_)[default__.safeIndex(d_1347_v14_, len(d_1389_v43_))]
                nw229_[int(8)] = (d_1383_v37_ if (self).f19 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evmfvq")))
                nw229_[int(9)] = d_1383_v37_
                d_1390_v44_ = nw229_
                index254_ = default__.safeIndex(523, (d_1390_v44_).length(0))
                (d_1390_v44_)[index254_] = d_1383_v37_
                index255_ = default__.safeIndex(523, (d_1390_v44_).length(0))
                (d_1390_v44_)[index255_] = d_1383_v37_
                (globalState).f3 = d_1347_v14_
                d_1391_v45_: _dafny.Seq
                d_1391_v45_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                (self).m7((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([not(self.f14)])])) + (d_1391_v45_), (d_1386_v40_) == (d_1386_v40_), d_1347_v14_, self.f14, globalState)
                d_1392_v46_: _dafny.Array
                nw230_ = _dafny.Array(False, 24)
                d_1392_v46_ = nw230_
                d_1393_v47_: _dafny.MultiSet
                d_1393_v47_ = _dafny.MultiSet([d_1385_v39_, d_1385_v39_, d_1385_v39_, d_1385_v39_, d_1385_v39_])
                index256_ = default__.safeIndex(750, (d_1392_v46_).length(0))
                (d_1392_v46_)[index256_] = (d_1347_v14_) == ((d_1393_v47_).cardinality)
                d_1394_v48_: _dafny.MultiSet
                d_1394_v48_ = _dafny.MultiSet([len(d_1386_v40_), 180])
                d_1395_v49_: _dafny.Array
                def lambda71_(d_1396_v14_):
                    def lambda72_(d_1397_i9_):
                        return _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1396_v14_: False}) for d_1398_i10_ in range(default__.abs(271))])

                    return lambda72_

                init42_ = lambda71_(d_1347_v14_)
                nw231_ = _dafny.Array(None, 28)
                for i0_42_ in range(nw231_.length(0)):
                    nw231_[i0_42_] = init42_(i0_42_)
                d_1395_v49_ = nw231_
                d_1399_v50_: _dafny.Map
                d_1399_v50_ = _dafny.Map({d_1347_v14_: self.f14})
                d_1400_v51_: _dafny.Seq
                d_1400_v51_ = _dafny.SeqWithoutIsStrInference([d_1399_v50_])
                index257_ = default__.safeIndex(598, (d_1395_v49_).length(0))
                (d_1395_v49_)[index257_] = d_1400_v51_
                d_1401_v52_: _dafny.Set
                d_1401_v52_ = _dafny.Set({(d_1394_v48_).cardinality, d_1347_v14_})
                index258_ = default__.safeIndex(750, (d_1392_v46_).length(0))
                index259_ = default__.safeIndex(598, (d_1395_v49_).length(0))
                def iife165_():
                    coll53_ = _dafny.Set()
                    compr_53_: int
                    for compr_53_ in (d_1401_v52_).Elements:
                        d_1402_v53_: int = compr_53_
                        if (d_1402_v53_) in (d_1401_v52_):
                            coll53_ = coll53_.union(_dafny.Set([(d_1402_v53_) - (216)]))
                    return _dafny.Set(coll53_)
                rhs231_ = self.f14
                rhs232_ = (len(iife165_()
                )) < ((d_1347_v14_) - (d_1347_v14_))
                rhs233_ = d_1394_v48_
                rhs234_ = _dafny.SeqWithoutIsStrInference([d_1399_v50_, (d_1399_v50_) | (d_1399_v50_), d_1399_v50_, ((d_1399_v50_).set(298, (self).f19)) | ((d_1399_v50_).set(d_1347_v14_, (self).f19))])
                lhs197_ = d_1392_v46_
                lhs198_ = default__.safeIndex(750, (d_1392_v46_).length(0))
                lhs199_ = globalState
                lhs200_ = d_1395_v49_
                lhs201_ = default__.safeIndex(598, (d_1395_v49_).length(0))
                lhs197_[lhs198_] = rhs231_
                lhs199_.f9 = rhs232_
                d_1394_v48_ = rhs233_
                lhs200_[lhs201_] = rhs234_
            d_1403_v54_: _dafny.Map
            d_1403_v54_ = _dafny.Map({self.f14: d_1347_v14_})
            d_1404_v58_: _dafny.Seq
            d_1404_v58_ = _dafny.SeqWithoutIsStrInference([d_1403_v54_])
            d_1405_v59_: _dafny.Seq
            d_1405_v59_ = _dafny.SeqWithoutIsStrInference([d_1347_v14_, d_1347_v14_])
            d_1406_v60_: _dafny.Set
            d_1406_v60_ = _dafny.Set({d_1347_v14_, (d_1405_v59_)[default__.safeIndex(((self).f15).cardinality, len(d_1405_v59_))], d_1347_v14_})
            d_1407_v61_: _dafny.Seq
            d_1407_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chvabiuu"))
            d_1408_v62_: _dafny.Array
            nw232_ = _dafny.Array(None, 23)
            nw232_[int(0)] = (d_1403_v54_) | (_dafny.Map({self.f14: d_1347_v14_}))
            nw232_[int(1)] = d_1403_v54_
            nw232_[int(2)] = d_1403_v54_
            def iife166_():
                coll54_ = _dafny.Set()
                compr_54_: int
                for compr_54_ in _dafny.IntegerRange(-621, 879):
                    d_1409_v55_: int = compr_54_
                    if ((-621) <= (d_1409_v55_)) and ((d_1409_v55_) < (879)):
                        def iife167_():
                            def iife169_():
                                coll57_ = _dafny.Map()
                                compr_57_: int
                                for compr_57_ in (_dafny.Map({(0) - (d_1347_v14_): self.f14})).keys.Elements:
                                    d_1410_v57_: int = compr_57_
                                    if (d_1410_v57_) in (_dafny.Map({(0) - (d_1347_v14_): self.f14})):
                                        coll57_[default__.safeDivisionInt(d_1410_v57_, d_1347_v14_)] = (self).f19
                                return _dafny.Map(coll57_)
                            coll55_ = _dafny.Map()
                            def iife168_():
                                coll56_ = _dafny.Map()
                                compr_56_: int
                                for compr_56_ in (_dafny.Map({(0) - (d_1347_v14_): self.f14})).keys.Elements:
                                    d_1410_v57_: int = compr_56_
                                    if (d_1410_v57_) in (_dafny.Map({(0) - (d_1347_v14_): self.f14})):
                                        coll56_[default__.safeDivisionInt(d_1410_v57_, d_1347_v14_)] = (self).f19
                                return _dafny.Map(coll56_)
                            compr_55_: int
                            for compr_55_ in (iife168_()
                            ).keys.Elements:
                                d_1411_v56_: int = compr_55_
                                if (d_1411_v56_) in (iife169_()
                                ):
                                    coll55_[default__.safeDivisionInt(d_1411_v56_, 403)] = d_1347_v14_
                            return _dafny.Map(coll55_)
                        coll54_ = coll54_.union(_dafny.Set([(d_1409_v55_) + (len(iife167_()
))]))
                return _dafny.Set(coll54_)
            nw232_[int(3)] = (d_1403_v54_) | (default__.fm1(False, d_1347_v14_, self.f14, iife166_()
            , globalState))
            nw232_[int(4)] = ((d_1404_v58_)[default__.safeIndex(d_1347_v14_, len(d_1404_v58_))]) | (d_1403_v54_)
            nw232_[int(5)] = d_1403_v54_
            nw232_[int(6)] = (d_1403_v54_ if (self).f19 else d_1403_v54_)
            nw232_[int(7)] = d_1403_v54_
            nw232_[int(8)] = _dafny.Map({self.f14: d_1347_v14_})
            nw232_[int(9)] = d_1403_v54_
            nw232_[int(10)] = d_1403_v54_
            nw232_[int(11)] = (d_1403_v54_).set(False, d_1347_v14_)
            nw232_[int(12)] = (_dafny.Map({(self).f19: d_1347_v14_})) | (default__.fm1((self).f19, d_1347_v14_, self.f14, d_1406_v60_, globalState))
            nw232_[int(13)] = d_1403_v54_
            nw232_[int(14)] = (d_1403_v54_) | (d_1403_v54_)
            nw232_[int(15)] = d_1403_v54_
            nw232_[int(16)] = d_1403_v54_
            nw232_[int(17)] = d_1403_v54_
            nw232_[int(18)] = d_1403_v54_
            nw232_[int(19)] = d_1403_v54_
            nw232_[int(20)] = _dafny.Map({(self).f19: d_1347_v14_})
            nw232_[int(21)] = (d_1403_v54_) | (d_1403_v54_)
            nw232_[int(22)] = (d_1403_v54_) | (_dafny.Map({not((self).f19): len(d_1407_v61_)}))
            d_1408_v62_ = nw232_
            index260_ = default__.safeIndex(860, (d_1408_v62_).length(0))
            (d_1408_v62_)[index260_] = d_1403_v54_
            index261_ = default__.safeIndex(860, (d_1408_v62_).length(0))
            (d_1408_v62_)[index261_] = d_1403_v54_
            d_1347_v14_ = (0) - ((d_1347_v14_ if (self).f19 else (0) - (d_1347_v14_)))
            d_1412_v63_: C1
            nw233_ = C1()
            nw233_.ctor__()
            d_1412_v63_ = nw233_
        elif True:
            (globalState).f9 = self.f14
            if (self).f19:
                d_1413_v64_: _dafny.Seq
                d_1413_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frcf"))
                d_1413_v64_ = (d_1413_v64_) + (d_1413_v64_)
                d_1414_v65_: _dafny.Map
                d_1414_v65_ = _dafny.Map({self.f14: ((self).f15).intersection((self).f15)})
                d_1414_v65_ = (d_1414_v65_).set(not ((self).f19) or (not((self).f19)), (self).f15)
                d_1415_v66_: _dafny.Array
                nw234_ = _dafny.Array(False, 3)
                d_1415_v66_ = nw234_
                d_1416_v67_: D4
                d_1416_v67_ = D4_DC13(d_1413_v64_, d_1347_v14_, (self).f19, d_1415_v66_)
                d_1417_v68_: _dafny.MultiSet
                d_1417_v68_ = _dafny.MultiSet([d_1347_v14_, d_1347_v14_, d_1347_v14_])
                d_1413_v64_ = (self).fm3(default__.fm0(self.f14, len((d_1416_v67_).cf22), d_1347_v14_, d_1417_v68_, globalState), len(_dafny.Set({d_1347_v14_})), self.f14, d_1347_v14_, globalState)
                d_1418_v69_: str
                d_1418_v69_ = _dafny.CodePoint('m')
                d_1419_v70_: _dafny.Map
                d_1419_v70_ = _dafny.Map({d_1418_v69_: self.f14})
                index262_ = default__.safeIndex(156, (d_1415_v66_).length(0))
                (d_1415_v66_)[index262_] = False
                d_1420_v71_: _dafny.Seq
                d_1420_v71_ = _dafny.SeqWithoutIsStrInference([(self).f19])
                d_1421_v72_: _dafny.Map
                d_1421_v72_ = _dafny.Map({len(d_1413_v64_): True})
                d_1422_v73_: _dafny.Map
                d_1422_v73_ = _dafny.Map({default__.fm22(d_1347_v14_, _dafny.SeqWithoutIsStrInference([d_1421_v72_]), self.f14, (self).f19, globalState): d_1415_v66_})
                d_1423_v74_: _dafny.Set
                d_1423_v74_ = _dafny.Set({d_1417_v68_})
                d_1424_v75_: _dafny.MultiSet
                d_1424_v75_ = _dafny.MultiSet([d_1420_v71_, d_1420_v71_, d_1420_v71_, d_1420_v71_])
                index263_ = default__.safeIndex(156, (d_1415_v66_).length(0))
                rhs235_ = (self).f19
                rhs236_ = ((d_1419_v70_) | (d_1419_v70_) if (d_1420_v71_)[default__.safeIndex(len(d_1422_v73_), len(d_1420_v71_))] else (default__.fm46(d_1347_v14_, (self).f19, globalState)) | (d_1419_v70_))
                rhs237_ = True
                rhs238_ = (self).fm17(d_1347_v14_, (d_1423_v74_).isdisjoint(d_1423_v74_), d_1424_v75_, d_1347_v14_, globalState)
                lhs202_ = globalState
                lhs203_ = d_1415_v66_
                lhs204_ = default__.safeIndex(156, (d_1415_v66_).length(0))
                lhs202_.f9 = rhs235_
                d_1419_v70_ = rhs236_
                lhs203_[lhs204_] = rhs237_
                d_1347_v14_ = rhs238_
                (globalState).f3 = default__.safeDivisionInt(d_1347_v14_, d_1347_v14_)
            elif True:
                d_1425_v76_: _dafny.Array
                nw235_ = _dafny.Array(False, 9)
                d_1425_v76_ = nw235_
                d_1426_v77_: _dafny.Set
                d_1426_v77_ = _dafny.Set({d_1425_v76_, d_1425_v76_})
                (globalState).f3 = (d_1347_v14_) * (len(d_1426_v77_))
                d_1427_v78_: _dafny.Seq
                d_1427_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wkk"))
                d_1428_v79_: _dafny.Set
                d_1428_v79_ = _dafny.Set({(d_1427_v78_) < (d_1427_v78_)})
                (globalState).f6 = d_1428_v79_
                (globalState).f3 = (len(d_1427_v78_)) * (d_1347_v14_)
                d_1429_v80_: _dafny.Map
                d_1429_v80_ = _dafny.Map({_dafny.CodePoint('i'): d_1347_v14_})
                d_1430_v81_: _dafny.Seq
                d_1430_v81_ = _dafny.SeqWithoutIsStrInference([d_1347_v14_, len(d_1427_v78_)])
                d_1431_v82_: _dafny.Set
                d_1431_v82_ = _dafny.Set({d_1430_v81_})
                d_1432_v83_: _dafny.Map
                d_1432_v83_ = _dafny.Map({(default__.fm23(d_1347_v14_, self.f14, d_1429_v80_, d_1431_v82_, globalState)) or (self.f14): d_1347_v14_})
                d_1432_v83_ = (d_1432_v83_).set(self.f14, d_1347_v14_)
                (globalState).f9 = (self).f19
            d_1433_v84_: _dafny.Array
            def lambda73_(d_1434_i11_):
                return (_dafny.SeqWithoutIsStrInference([(self).f19])) + (_dafny.SeqWithoutIsStrInference([not((self).f19)]))

            init43_ = lambda73_
            nw236_ = _dafny.Array(None, 9)
            for i0_43_ in range(nw236_.length(0)):
                nw236_[i0_43_] = init43_(i0_43_)
            d_1433_v84_ = nw236_
            d_1435_v85_: _dafny.Seq
            d_1435_v85_ = _dafny.SeqWithoutIsStrInference([self.f14, self.f14, (self).f19])
            index264_ = default__.safeIndex(360, (d_1433_v84_).length(0))
            (d_1433_v84_)[index264_] = d_1435_v85_
            d_1436_v86_: _dafny.Map
            d_1436_v86_ = _dafny.Map({not(not(self.f14)): d_1347_v14_})
            d_1437_v87_: _dafny.Seq
            d_1437_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pw"))
            d_1438_v88_: _dafny.Seq
            d_1438_v88_ = _dafny.SeqWithoutIsStrInference([d_1435_v85_])
            d_1439_v89_: _dafny.Seq
            d_1439_v89_ = _dafny.SeqWithoutIsStrInference([d_1347_v14_, (self).fm17(((d_1436_v86_)[self.f14] if (self.f14) in (d_1436_v86_) else d_1347_v14_), self.f14, _dafny.MultiSet([d_1435_v85_, (d_1435_v85_).set(default__.safeIndex(len(d_1437_v87_), len(d_1435_v85_)), (self).f19), d_1435_v85_, d_1435_v85_, d_1435_v85_]), (0) - (len((d_1438_v88_).set(default__.safeIndex((0) - (d_1347_v14_), len(d_1438_v88_)), _dafny.SeqWithoutIsStrInference([(self).f19])))), globalState)])
            index265_ = default__.safeIndex(360, (d_1433_v84_).length(0))
            rhs239_ = (d_1435_v85_).set(default__.safeIndex(d_1347_v14_, len(d_1435_v85_)), (self).f19)
            rhs240_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(d_1347_v14_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldnplt")))) for d_1440_i12_ in range(default__.abs(400))])
            lhs205_ = d_1433_v84_
            lhs206_ = default__.safeIndex(360, (d_1433_v84_).length(0))
            lhs205_[lhs206_] = rhs239_
            d_1439_v89_ = rhs240_
            (globalState).f9 = not(not((self.f14) == (not(self.f14))))
            (globalState).f9 = self.f14
        d_1441_v90_: str
        d_1441_v90_ = _dafny.CodePoint('o')
        d_1441_v90_ = d_1441_v90_
        d_1442_v91_: _dafny.Array
        nw237_ = _dafny.Array(False, 7)
        d_1442_v91_ = nw237_
        d_1443_v92_: _dafny.Seq
        d_1443_v92_ = _dafny.SeqWithoutIsStrInference([(self).f19, True, (self).f19, self.f14, (self).f19])
        d_1444_v93_: D1
        d_1444_v93_ = D1_DC3(338)
        d_1445_v94_: _dafny.Map
        d_1445_v94_ = _dafny.Map({d_1347_v14_: d_1444_v93_})
        index266_ = default__.safeIndex(821, (d_1442_v91_).length(0))
        (d_1442_v91_)[index266_] = (d_1443_v92_)[default__.safeIndex(len(d_1445_v94_), len(d_1443_v92_))]
        index267_ = default__.safeIndex(821, (d_1442_v91_).length(0))
        (d_1442_v91_)[index267_] = (self).f19
        d_1446_v95_: _dafny.Map
        d_1446_v95_ = _dafny.Map({self.f14: 101})
        d_1447_v96_: _dafny.Map
        d_1447_v96_ = _dafny.Map({d_1347_v14_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "voqhvj"))).set(default__.safeIndex(d_1347_v14_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "voqhvj")))), d_1441_v90_)})
        d_1448_v97_: _dafny.MultiSet
        d_1448_v97_ = _dafny.MultiSet([len(d_1446_v95_), d_1347_v14_, len(d_1447_v96_)])
        d_1441_v90_ = default__.fm25(d_1347_v14_, ((0) - (d_1347_v14_)) - ((d_1448_v97_).cardinality), not ((d_1442_v91_)[default__.safeIndex(821, (d_1442_v91_).length(0))]) or ((self).f19), d_1446_v95_, globalState)

    def m7(self, p0, p1, p2, p3, globalState):
        d_1449_v0_: C1
        nw238_ = C1()
        nw238_.ctor__()
        d_1449_v0_ = nw238_
        d_1450_v1_: _dafny.Array
        nw239_ = _dafny.Array(False, 20)
        d_1450_v1_ = nw239_
        index268_ = default__.safeIndex(886, (d_1450_v1_).length(0))
        (d_1450_v1_)[index268_] = p1
        d_1451_v2_: _dafny.Set
        d_1451_v2_ = _dafny.Set({(self).f15})
        d_1452_v3_: _dafny.Array
        nw240_ = _dafny.Array(None, 7)
        nw240_[int(0)] = d_1451_v2_
        nw240_[int(1)] = (D11_DC33(d_1451_v2_)).cf57
        nw240_[int(2)] = d_1451_v2_
        nw240_[int(3)] = d_1451_v2_
        nw240_[int(4)] = (d_1451_v2_) - (d_1451_v2_)
        nw240_[int(5)] = d_1451_v2_
        nw240_[int(6)] = (d_1451_v2_) - (_dafny.Set({(self).f15, (self).f15}))
        d_1452_v3_ = nw240_
        index269_ = default__.safeIndex(43, (d_1452_v3_).length(0))
        (d_1452_v3_)[index269_] = (default__.fm47(globalState)) - (d_1451_v2_)
        d_1453_v4_: _dafny.Set
        d_1453_v4_ = _dafny.Set({p2})
        d_1454_v5_: _dafny.Seq
        d_1454_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovck"))
        d_1455_v6_: _dafny.Seq
        d_1455_v6_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_1453_v4_)), (0) - (-295), p2, len(_dafny.Map({len(d_1454_v5_): d_1449_v0_}))])
        index270_ = default__.safeIndex(886, (d_1450_v1_).length(0))
        index271_ = default__.safeIndex(43, (d_1452_v3_).length(0))
        rhs241_ = (not (p3) or (p3)) and ((d_1455_v6_) <= (d_1455_v6_))
        rhs242_ = d_1451_v2_
        rhs243_ = True
        lhs207_ = d_1450_v1_
        lhs208_ = default__.safeIndex(886, (d_1450_v1_).length(0))
        lhs209_ = d_1452_v3_
        lhs210_ = default__.safeIndex(43, (d_1452_v3_).length(0))
        lhs211_ = globalState
        lhs207_[lhs208_] = rhs241_
        lhs209_[lhs210_] = rhs242_
        lhs211_.f9 = rhs243_
        rhs244_ = p2
        rhs245_ = 538
        lhs212_ = globalState
        lhs213_ = globalState
        lhs212_.f3 = rhs244_
        lhs213_.f3 = rhs245_
        d_1456_v7_: str
        d_1456_v7_ = _dafny.CodePoint('e')
        d_1457_v8_: _dafny.Map
        d_1457_v8_ = _dafny.Map({p2: (0) - (len(d_1454_v5_))})
        d_1458_v9_: D4
        d_1458_v9_ = D4_DC13(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sikdur")), p2, (self).f19, d_1450_v1_)
        d_1459_v11_: _dafny.Map
        d_1459_v11_ = _dafny.Map({p2: p3})
        d_1460_v12_: _dafny.Seq
        def iife170_():
            coll58_ = _dafny.Map()
            compr_58_: int
            for compr_58_ in _dafny.IntegerRange(-314, -878):
                d_1461_v10_: int = compr_58_
                if ((-314) <= (d_1461_v10_)) and ((d_1461_v10_) < (-878)):
                    coll58_[(d_1461_v10_) + (p2)] = self.f14
            return _dafny.Map(coll58_)
        d_1460_v12_ = _dafny.SeqWithoutIsStrInference([iife170_()
        , d_1459_v11_, d_1459_v11_, d_1459_v11_])
        d_1462_v13_: D9
        d_1462_v13_ = D9_DC30(d_1456_v7_, p1, len(d_1457_v8_), default__.fm22((d_1458_v9_).cf23, d_1460_v12_, (self).f19, (d_1450_v1_)[default__.safeIndex(886, (d_1450_v1_).length(0))], globalState), p2)
        d_1463_v14_: _dafny.Map
        d_1463_v14_ = _dafny.Map({d_1462_v13_: p2})
        hi4_ = (len(_dafny.SeqWithoutIsStrInference([d_1456_v7_ for d_1464_i1_ in range(default__.abs(934))]))) + (p2)
        for d_1465_i0_ in range(len((d_1463_v14_).set(d_1462_v13_, p2)), hi4_):
            d_1466_v15_: _dafny.MultiSet
            d_1466_v15_ = _dafny.MultiSet([p2])
            d_1467_v16_: D2
            d_1467_v16_ = D2_DC9((self).f19, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1468_i2_ in range(default__.abs(41))])), d_1456_v7_)
            d_1469_v17_: _dafny.Array
            nw241_ = _dafny.Array(None, 6)
            nw241_[int(0)] = d_1467_v16_
            nw241_[int(1)] = d_1467_v16_
            nw241_[int(2)] = d_1467_v16_
            nw241_[int(3)] = d_1467_v16_
            nw241_[int(4)] = D2_DC9(p3, p2, d_1456_v7_)
            nw241_[int(5)] = D2_DC9(p1, d_1465_i0_, d_1456_v7_)
            d_1469_v17_ = nw241_
            d_1470_v18_: _dafny.Seq
            d_1470_v18_ = _dafny.SeqWithoutIsStrInference([d_1469_v17_])
            d_1471_v19_: _dafny.Array
            nw242_ = _dafny.Array(int(0), 22)
            d_1471_v19_ = nw242_
            d_1472_v20_: D5
            d_1472_v20_ = D5_DC16(d_1471_v19_)
            d_1473_v21_: D5
            d_1473_v21_ = D5_DC19(d_1472_v20_)
            d_1474_v22_: _dafny.Map
            d_1474_v22_ = _dafny.Map({(self).fm2(self.f14, p3, globalState): d_1472_v20_})
            d_1475_v23_: D5
            d_1475_v23_ = D5_DC19(((d_1474_v22_)[(self).f19] if ((self).f19) in (d_1474_v22_) else d_1473_v21_))
            d_1476_v24_: D7
            d_1476_v24_ = D7_DC22(d_1450_v1_, d_1470_v18_, _dafny.SeqWithoutIsStrInference([d_1450_v1_, d_1450_v1_]), d_1450_v1_, d_1475_v23_)
            d_1477_v25_: _dafny.Seq
            d_1477_v25_ = _dafny.SeqWithoutIsStrInference([d_1476_v24_])
            d_1478_v26_: _dafny.Array
            nw243_ = _dafny.Array(None, 11)
            nw243_[int(0)] = p2
            nw243_[int(1)] = (p2) + (p2)
            nw243_[int(2)] = p2
            nw243_[int(3)] = d_1465_i0_
            nw243_[int(4)] = ((((self).f15)[(d_1450_v1_)[default__.safeIndex(886, (d_1450_v1_).length(0))]] if ((d_1450_v1_)[default__.safeIndex(886, (d_1450_v1_).length(0))]) in ((self).f15) else default__.fm0((d_1450_v1_)[default__.safeIndex(886, (d_1450_v1_).length(0))], 24, d_1465_i0_, d_1466_v15_, globalState))) * (p2)
            nw243_[int(5)] = default__.safeDivisionInt((0) - (p2), len(_dafny.Map({d_1465_i0_: p2})))
            nw243_[int(6)] = len(d_1454_v5_)
            nw243_[int(7)] = default__.fm0(self.f14, p2, d_1465_i0_, d_1466_v15_, globalState)
            nw243_[int(8)] = p2
            nw243_[int(9)] = (d_1465_i0_) - (len(d_1454_v5_))
            nw243_[int(10)] = len(d_1477_v25_)
            d_1478_v26_ = nw243_
            d_1479_v27_: _dafny.Seq
            d_1479_v27_ = _dafny.SeqWithoutIsStrInference([self.f14, self.f14, self.f14, p3, self.f14])
            d_1480_v28_: _dafny.MultiSet
            d_1480_v28_ = _dafny.MultiSet([d_1479_v27_])
            index272_ = default__.safeIndex(23, (d_1471_v19_).length(0))
            (d_1471_v19_)[index272_] = (self).fm17(p2, not((self).f19), d_1480_v28_, (0) - (default__.fm0((d_1450_v1_)[default__.safeIndex(886, (d_1450_v1_).length(0))], d_1465_i0_, d_1465_i0_, d_1466_v15_, globalState)), globalState)
            d_1481_v29_: D12
            d_1481_v29_ = D12_DC37(d_1480_v28_)
            index273_ = default__.safeIndex(23, (d_1471_v19_).length(0))
            (d_1471_v19_)[index273_] = default__.safeModuloInt((p2) + (528), (self).fm17(d_1465_i0_, (d_1479_v27_)[default__.safeIndex(p2, len(d_1479_v27_))], (d_1481_v29_).cf63, d_1465_i0_, globalState))
            d_1482_v30_: _dafny.Map
            d_1482_v30_ = _dafny.Map({self.f14: (d_1471_v19_)[default__.safeIndex(23, (d_1471_v19_).length(0))]})
            (globalState).f9 = (128) > ((d_1465_i0_) + (((d_1482_v30_)[(self).f19] if ((self).f19) in (d_1482_v30_) else p2)))
            rhs246_ = d_1456_v7_
            rhs247_ = d_1456_v7_
            rhs248_ = d_1471_v19_
            rhs249_ = ((d_1482_v30_)[p1] if (p1) in (d_1482_v30_) else (p2) - (p2))
            lhs214_ = globalState
            d_1456_v7_ = rhs246_
            d_1456_v7_ = rhs247_
            d_1471_v19_ = rhs248_
            lhs214_.f3 = rhs249_
            (globalState).f9 = (d_1479_v27_) == (d_1479_v27_)
        d_1483_v31_: C2
        nw244_ = C2()
        nw244_.ctor__(d_1459_v11_)
        d_1483_v31_ = nw244_
        d_1483_v31_ = d_1483_v31_
        d_1484_v32_: _dafny.Seq
        d_1484_v32_ = _dafny.SeqWithoutIsStrInference([p1])
        d_1485_v33_: _dafny.Map
        d_1485_v33_ = _dafny.Map({(d_1450_v1_)[default__.safeIndex(886, (d_1450_v1_).length(0))]: (d_1484_v32_) + (d_1484_v32_)})
        rhs250_ = len(((d_1485_v33_)[(self.f14) and (p1)] if ((self.f14) and (p1)) in (d_1485_v33_) else d_1484_v32_))
        rhs251_ = not ((d_1450_v1_)[default__.safeIndex(886, (d_1450_v1_).length(0))]) or ((self).f19)
        rhs252_ = (self).f19
        lhs215_ = globalState
        lhs216_ = globalState
        lhs217_ = globalState
        lhs215_.f3 = rhs250_
        lhs216_.f9 = rhs251_
        lhs217_.f9 = rhs252_

    @property
    def f19(self):
        return self._f19

class C5(T0):
    def  __init__(self):
        self._f14: bool = False
        self._f15: _dafny.MultiSet = _dafny.MultiSet({})
        self._f18: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    @property
    def f15(self):
        return self._f15
    def ctor__(self, f18, f14, f15):
        (self)._f18 = f18
        (self).f14 = f14
        (self)._f15 = f15

    def fm2(self, p0, p1, globalState):
        return self.f14

    def fm3(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")))

    def fm15(self, p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kfx")))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1486_i0_ in range(default__.abs(654))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1487_i1_ in range(default__.abs(25))])))

    def fm16(self, p0, p1, p2, globalState):
        return not(self.f14)

    def m0(self, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        r3: bool = False
        d_1488_v0_: _dafny.Seq
        d_1488_v0_ = _dafny.SeqWithoutIsStrInference([self.f14])
        d_1489_v1_: D2
        d_1489_v1_ = D2_DC7(d_1488_v0_)
        d_1490_v2_: _dafny.Map
        d_1490_v2_ = _dafny.Map({(self).f18: (d_1489_v1_).cf9})
        d_1491_v3_: _dafny.Seq
        d_1491_v3_ = _dafny.SeqWithoutIsStrInference([d_1488_v0_, d_1488_v0_, d_1488_v0_, d_1488_v0_, d_1488_v0_])
        d_1490_v2_ = (d_1490_v2_).set((self).f18, (d_1491_v3_)[default__.safeIndex(840, len(d_1491_v3_))])
        d_1492_v4_: _dafny.Seq
        d_1492_v4_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({True, self.f14, self.f14, True, self.f14})])
        d_1493_v5_: C4
        nw245_ = C4()
        nw245_.ctor__(self.f14, (len(d_1491_v3_)) > (len(d_1492_v4_)), (self).f15)
        d_1493_v5_ = nw245_
        r0 = len((d_1491_v3_)[default__.safeIndex((self).f18, len(d_1491_v3_))])
        r0 = (self).f18
        d_1494_v6_: _dafny.Array
        nw246_ = _dafny.Array(False, 1)
        d_1494_v6_ = nw246_
        index274_ = default__.safeIndex(819, (d_1494_v6_).length(0))
        (d_1494_v6_)[index274_] = self.f14
        index275_ = default__.safeIndex(819, (d_1494_v6_).length(0))
        (d_1494_v6_)[index275_] = (self.f14 if self.f14 else (d_1493_v5_).f19)
        d_1495_v7_: _dafny.Map
        d_1495_v7_ = _dafny.Map({(self).f18: (self).f18})
        d_1496_v8_: _dafny.Map
        d_1496_v8_ = _dafny.Map({(d_1493_v5_).f19: len(d_1495_v7_)})
        d_1497_v9_: _dafny.Seq
        d_1497_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lm"))
        d_1498_v10_: _dafny.Seq
        d_1498_v10_ = _dafny.SeqWithoutIsStrInference([(self).f18])
        d_1499_v11_: _dafny.Array
        nw247_ = _dafny.Array(None, 26)
        nw247_[int(0)] = (self).f18
        nw247_[int(1)] = (self).f18
        nw247_[int(2)] = (self).f18
        nw247_[int(3)] = (self).f18
        nw247_[int(4)] = (self).f18
        nw247_[int(5)] = (self).f18
        nw247_[int(6)] = (self).f18
        nw247_[int(7)] = (self).f18
        nw247_[int(8)] = (self).f18
        nw247_[int(9)] = len(d_1496_v8_)
        nw247_[int(10)] = len(d_1497_v9_)
        nw247_[int(11)] = -604
        nw247_[int(12)] = len(d_1498_v10_)
        nw247_[int(13)] = 817
        nw247_[int(14)] = len(_dafny.SeqWithoutIsStrInference([(d_1494_v6_)[default__.safeIndex(819, (d_1494_v6_).length(0))], (d_1493_v5_).f19]))
        nw247_[int(15)] = (self).f18
        nw247_[int(16)] = (self).f18
        nw247_[int(17)] = (self).f18
        nw247_[int(18)] = (self).f18
        nw247_[int(19)] = 975
        nw247_[int(20)] = (self).f18
        nw247_[int(21)] = -187
        nw247_[int(22)] = (self).f18
        nw247_[int(23)] = (self).f18
        nw247_[int(24)] = (d_1498_v10_)[default__.safeIndex((self).f18, len(d_1498_v10_))]
        nw247_[int(25)] = (self).f18
        d_1499_v11_ = nw247_
        d_1500_v12_: D5
        d_1500_v12_ = D5_DC16(d_1499_v11_)
        d_1501_v13_: _dafny.Map
        d_1501_v13_ = _dafny.Map({d_1497_v9_: d_1499_v11_})
        d_1502_v14_: _dafny.Map
        d_1502_v14_ = _dafny.Map({(d_1493_v5_).f19: d_1499_v11_})
        pat_let_tv40_ = d_1499_v11_
        d_1503_v15_: _dafny.Array
        nw248_ = _dafny.Array(None, 26)
        def iife171_(_pat_let56_0):
            def iife172_(d_1504_dt__update__tmp_h0_):
                def iife173_(_pat_let57_0):
                    def iife174_(d_1505_dt__update_hcf29_h0_):
                        return D5_DC16(d_1505_dt__update_hcf29_h0_)
                    return iife174_(_pat_let57_0)
                return iife173_(pat_let_tv40_)
            return iife172_(_pat_let56_0)
        nw248_[int(0)] = (iife171_(d_1500_v12_)).cf29
        nw248_[int(1)] = d_1499_v11_
        nw248_[int(2)] = d_1499_v11_
        nw248_[int(3)] = d_1499_v11_
        nw248_[int(4)] = d_1499_v11_
        nw248_[int(5)] = (d_1499_v11_ if (d_1493_v5_).f19 else d_1499_v11_)
        nw248_[int(6)] = d_1499_v11_
        nw248_[int(7)] = d_1499_v11_
        nw248_[int(8)] = d_1499_v11_
        nw248_[int(9)] = d_1499_v11_
        nw248_[int(10)] = d_1499_v11_
        nw248_[int(11)] = ((d_1501_v13_)[_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1506_i0_ in range(default__.abs(979))])] if (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1507_i0_ in range(default__.abs(979))])) in (d_1501_v13_) else d_1499_v11_)
        nw248_[int(12)] = d_1499_v11_
        nw248_[int(13)] = d_1499_v11_
        nw248_[int(14)] = d_1499_v11_
        nw248_[int(15)] = d_1499_v11_
        nw248_[int(16)] = d_1499_v11_
        nw248_[int(17)] = d_1499_v11_
        nw248_[int(18)] = d_1499_v11_
        nw248_[int(19)] = d_1499_v11_
        nw248_[int(20)] = d_1499_v11_
        nw248_[int(21)] = d_1499_v11_
        nw248_[int(22)] = d_1499_v11_
        nw248_[int(23)] = d_1499_v11_
        nw248_[int(24)] = ((d_1502_v14_)[(d_1494_v6_)[default__.safeIndex(819, (d_1494_v6_).length(0))]] if ((d_1494_v6_)[default__.safeIndex(819, (d_1494_v6_).length(0))]) in (d_1502_v14_) else d_1499_v11_)
        nw248_[int(25)] = d_1499_v11_
        d_1503_v15_ = nw248_
        nw249_ = _dafny.Array(_dafny.Array(None, 0), 9)
        d_1503_v15_ = nw249_
        r0 = (self).f18
        r1 = (self.f14) and (self.f14)
        d_1508_v16_: _dafny.Seq
        d_1508_v16_ = _dafny.SeqWithoutIsStrInference([d_1494_v6_])
        d_1509_v17_: D4
        d_1509_v17_ = D4_DC13(d_1497_v9_, (self).f18, (d_1494_v6_)[default__.safeIndex(819, (d_1494_v6_).length(0))], (d_1508_v16_)[default__.safeIndex((self).f18, len(d_1508_v16_))])
        d_1510_v18_: _dafny.Map
        d_1510_v18_ = _dafny.Map({((_dafny.Map({(self).f18: (d_1494_v6_)[default__.safeIndex(819, (d_1494_v6_).length(0))]})).set((self).f18, (d_1493_v5_).f19)).set((d_1509_v17_).cf23, self.f14): (self).f18})
        r2 = d_1510_v18_
        r3 = (d_1494_v6_)[default__.safeIndex(819, (d_1494_v6_).length(0))]
        return r0, r1, r2, r3

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        d_1511_v0_: _dafny.Array
        def lambda74_(d_1512_i1_):
            return (d_1512_i1_) * ((_dafny.MultiSet([_dafny.CodePoint('b')])).cardinality)

        init44_ = lambda74_
        nw250_ = _dafny.Array(None, 27)
        for i0_44_ in range(nw250_.length(0)):
            nw250_[i0_44_] = init44_(i0_44_)
        d_1511_v0_ = nw250_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1511_v0_).length(0)):
            d_1513_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_1513_i0_)) and ((d_1513_i0_) < ((d_1511_v0_).length(0)))):
                (d_1511_v0_)[(d_1513_i0_)] = default__.safeModuloInt(d_1513_i0_, p1)
        d_1514_v1_: _dafny.Map
        d_1514_v1_ = _dafny.Map({self.f14: self.f14})
        d_1515_v2_: D3
        d_1515_v2_ = D3_DC10(_dafny.Map({self.f14: self.f14}))
        def lambda75_(source15_):
            if source15_.is_DC11:
                d_1516___mcc_h0_ = source15_.cf17
                d_1517___mcc_h1_ = source15_.cf18
                d_1518___mcc_h2_ = source15_.cf19
                d_1519___mcc_h3_ = source15_.cf20
                d_1520_cf20_ = d_1519___mcc_h3_
                d_1521_cf19_ = d_1518___mcc_h2_
                d_1522_cf18_ = d_1517___mcc_h1_
                d_1523_cf17_ = d_1516___mcc_h0_
                return _dafny.Map({self.f14: d_1521_cf19_})
            elif True:
                d_1524___mcc_h4_ = source15_.cf16
                d_1525_cf16_ = d_1524___mcc_h4_
                return (d_1525_cf16_) | (d_1525_cf16_)

        d_1514_v1_ = lambda75_(d_1515_v2_)
        d_1526_v3_: _dafny.Seq
        d_1526_v3_ = _dafny.SeqWithoutIsStrInference([(self).f18])
        d_1527_v4_: _dafny.Seq
        d_1527_v4_ = _dafny.SeqWithoutIsStrInference([self.f14])
        d_1528_v5_: C4
        nw251_ = C4()
        nw251_.ctor__((d_1526_v3_) <= (d_1526_v3_), self.f14, _dafny.MultiSet([(d_1527_v4_)[default__.safeIndex((self).f18, len(d_1527_v4_))], self.f14]))
        d_1528_v5_ = nw251_
        d_1529_v6_: str
        d_1529_v6_ = _dafny.CodePoint('t')
        d_1530_v7_: _dafny.Seq
        d_1530_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
        d_1531_v8_: _dafny.Seq
        d_1531_v8_ = _dafny.SeqWithoutIsStrInference([d_1530_v7_, d_1530_v7_])
        d_1532_v9_: _dafny.Array
        nw252_ = _dafny.Array(None, 21)
        nw252_[int(0)] = (d_1528_v5_).f19
        nw252_[int(1)] = (d_1528_v5_).f19
        nw252_[int(2)] = (720) >= ((self).f18)
        nw252_[int(3)] = self.f14
        nw252_[int(4)] = self.f14
        nw252_[int(5)] = ((self).f15).issubset((self).f15)
        nw252_[int(6)] = (d_1528_v5_).f19
        nw252_[int(7)] = self.f14
        nw252_[int(8)] = self.f14
        nw252_[int(9)] = self.f14
        nw252_[int(10)] = True
        nw252_[int(11)] = self.f14
        nw252_[int(12)] = (d_1529_v6_) in (_dafny.SeqWithoutIsStrInference([d_1529_v6_ for d_1533_i2_ in range(default__.abs(815))]))
        nw252_[int(13)] = (p1) <= (len(d_1530_v7_))
        nw252_[int(14)] = ((self).f18) < (len(d_1531_v8_))
        nw252_[int(15)] = (p0) != (p0)
        nw252_[int(16)] = not((d_1528_v5_).f19)
        nw252_[int(17)] = (p0).issubset(p0)
        nw252_[int(18)] = (d_1528_v5_).f19
        nw252_[int(19)] = (self).fm2((d_1528_v5_).f19, (d_1528_v5_).f19, globalState)
        nw252_[int(20)] = self.f14
        d_1532_v9_ = nw252_
        index276_ = default__.safeIndex(928, (d_1532_v9_).length(0))
        (d_1532_v9_)[index276_] = self.f14
        index277_ = default__.safeIndex(46, (d_1511_v0_).length(0))
        (d_1511_v0_)[index277_] = -934
        d_1534_v10_: _dafny.Array
        nw253_ = _dafny.Array(_dafny.Map({}), 25)
        d_1534_v10_ = nw253_
        d_1535_v11_: _dafny.Map
        d_1535_v11_ = _dafny.Map({(self).fm3(p1, p1, self.f14, p1, globalState): self.f14})
        index278_ = default__.safeIndex(245, (d_1534_v10_).length(0))
        (d_1534_v10_)[index278_] = d_1535_v11_
        d_1536_v12_: _dafny.MultiSet
        d_1536_v12_ = _dafny.MultiSet([-837])
        index279_ = default__.safeIndex(928, (d_1532_v9_).length(0))
        index280_ = default__.safeIndex(46, (d_1511_v0_).length(0))
        index281_ = default__.safeIndex(245, (d_1534_v10_).length(0))
        rhs253_ = (d_1536_v12_).isdisjoint((_dafny.MultiSet([p1])).set((self).f18, default__.abs((self).f18)))
        rhs254_ = (0) - ((self).f18)
        rhs255_ = (d_1528_v5_).f19
        rhs256_ = d_1535_v11_
        lhs218_ = d_1532_v9_
        lhs219_ = default__.safeIndex(928, (d_1532_v9_).length(0))
        lhs220_ = d_1511_v0_
        lhs221_ = default__.safeIndex(46, (d_1511_v0_).length(0))
        lhs222_ = self
        lhs223_ = d_1534_v10_
        lhs224_ = default__.safeIndex(245, (d_1534_v10_).length(0))
        lhs218_[lhs219_] = rhs253_
        lhs220_[lhs221_] = rhs254_
        lhs222_.f14 = rhs255_
        lhs223_[lhs224_] = rhs256_
        d_1537_v13_: _dafny.Seq
        d_1537_v13_ = _dafny.SeqWithoutIsStrInference([d_1536_v12_, _dafny.MultiSet([(self).f18])])
        d_1538_v14_: _dafny.Set
        d_1538_v14_ = _dafny.Set({(d_1537_v13_)[default__.safeIndex(p1, len(d_1537_v13_))]})
        hi5_ = default__.safeDivisionInt((d_1511_v0_)[default__.safeIndex(46, (d_1511_v0_).length(0))], len(d_1538_v14_))
        for d_1539_i3_ in range((0) - ((d_1511_v0_)[default__.safeIndex(46, (d_1511_v0_).length(0))]), hi5_):
            d_1540_v15_: _dafny.Set
            d_1540_v15_ = _dafny.Set({d_1527_v4_})
            d_1541_v17_: _dafny.Set
            d_1541_v17_ = _dafny.Set({(d_1528_v5_).fm2((d_1528_v5_).f19, self.f14, globalState)})
            d_1542_v18_: _dafny.Map
            d_1542_v18_ = _dafny.Map({(d_1532_v9_)[default__.safeIndex(928, (d_1532_v9_).length(0))]: d_1541_v17_})
            d_1543_v19_: _dafny.Map
            d_1543_v19_ = _dafny.Map({d_1526_v3_: len(((d_1542_v18_)[(d_1528_v5_).f19] if ((d_1528_v5_).f19) in (d_1542_v18_) else d_1541_v17_))})
            index282_ = default__.safeIndex(46, (d_1511_v0_).length(0))
            index283_ = default__.safeIndex(46, (d_1511_v0_).length(0))
            def iife175_():
                coll59_ = _dafny.Set()
                compr_59_: _dafny.Seq
                for compr_59_ in (d_1540_v15_).Elements:
                    d_1544_v16_: _dafny.Seq = compr_59_
                    if (d_1544_v16_) in (d_1540_v15_):
                        coll59_ = coll59_.union(_dafny.Set([d_1544_v16_]))
                return _dafny.Set(coll59_)
            rhs257_ = d_1530_v7_
            rhs258_ = (0) - (default__.safeDivisionInt(((0) - (p1)) * (d_1539_i3_), len((iife175_()
            ).intersection(d_1540_v15_))))
            rhs259_ = (self).f18
            rhs260_ = (self).fm15((d_1528_v5_).f19, ((d_1511_v0_)[default__.safeIndex(46, (d_1511_v0_).length(0))]) == (p1), ((d_1543_v19_)[d_1526_v3_] if (d_1526_v3_) in (d_1543_v19_) else p1), globalState)
            rhs261_ = (d_1532_v9_)[default__.safeIndex(928, (d_1532_v9_).length(0))]
            lhs225_ = d_1511_v0_
            lhs226_ = default__.safeIndex(46, (d_1511_v0_).length(0))
            lhs227_ = d_1511_v0_
            lhs228_ = default__.safeIndex(46, (d_1511_v0_).length(0))
            lhs229_ = globalState
            d_1530_v7_ = rhs257_
            lhs225_[lhs226_] = rhs258_
            lhs227_[lhs228_] = rhs259_
            d_1530_v7_ = rhs260_
            lhs229_.f9 = rhs261_
            d_1545_v20_: _dafny.Array
            def lambda76_(d_1546_i4_):
                return _dafny.CodePoint('g')

            init45_ = lambda76_
            nw254_ = _dafny.Array(None, 7)
            for i0_45_ in range(nw254_.length(0)):
                nw254_[i0_45_] = init45_(i0_45_)
            d_1545_v20_ = nw254_
            d_1547_v21_: _dafny.Map
            d_1547_v21_ = _dafny.Map({(d_1532_v9_)[default__.safeIndex(928, (d_1532_v9_).length(0))]: 95})
            index284_ = default__.safeIndex(812, (d_1545_v20_).length(0))
            (d_1545_v20_)[index284_] = (default__.fm25(p1, p1, False, d_1547_v21_, globalState) if (d_1528_v5_).f19 else d_1529_v6_)
            index285_ = default__.safeIndex(812, (d_1545_v20_).length(0))
            (d_1545_v20_)[index285_] = _dafny.CodePoint('w')
            d_1548_v22_: _dafny.Map
            d_1548_v22_ = _dafny.Map({d_1547_v21_: (self).fm2((d_1528_v5_).f19, (d_1528_v5_).f19, globalState)})
            d_1549_v24_: _dafny.Set
            d_1549_v24_ = _dafny.Set({d_1547_v21_})
            def iife176_():
                coll60_ = _dafny.Set()
                compr_60_: _dafny.Map
                for compr_60_ in (d_1548_v22_).keys.Elements:
                    d_1550_v23_: _dafny.Map = compr_60_
                    if (d_1550_v23_) in (d_1548_v22_):
                        coll60_ = coll60_.union(_dafny.Set([d_1550_v23_]))
                return _dafny.Set(coll60_)
            if (iife176_()
            ).issubset(d_1549_v24_):
                r1 = d_1539_i3_
                index286_ = default__.safeIndex(812, (d_1545_v20_).length(0))
                (d_1545_v20_)[index286_] = _dafny.CodePoint('h')
                index287_ = default__.safeIndex(46, (d_1511_v0_).length(0))
                rhs262_ = (d_1528_v5_).f19
                rhs263_ = not(((d_1528_v5_).f19) == (self.f14))
                rhs264_ = len(_dafny.SeqWithoutIsStrInference([len(p0) for d_1551_i5_ in range(default__.abs(720))]))
                lhs230_ = globalState
                lhs231_ = globalState
                lhs232_ = d_1511_v0_
                lhs233_ = default__.safeIndex(46, (d_1511_v0_).length(0))
                lhs230_.f9 = rhs262_
                lhs231_.f9 = rhs263_
                lhs232_[lhs233_] = rhs264_
                nw255_ = _dafny.Array(False, 29)
                (globalState).f2 = nw255_
                d_1552_v25_: _dafny.Array
                nw256_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_1552_v25_ = nw256_
                d_1553_v26_: D0
                d_1553_v26_ = D0_DC1((d_1528_v5_).f19)
                rhs265_ = (d_1511_v0_)[default__.safeIndex(46, (d_1511_v0_).length(0))]
                rhs266_ = (((d_1553_v26_).cf1) not in (d_1547_v21_) if (d_1528_v5_).f19 else ((d_1532_v9_)[default__.safeIndex(928, (d_1532_v9_).length(0))]) or ((d_1532_v9_)[default__.safeIndex(928, (d_1532_v9_).length(0))]))
                rhs267_ = ((23) * (d_1539_i3_)) - ((63) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))))
                rhs268_ = d_1552_v25_
                lhs234_ = globalState
                lhs235_ = globalState
                lhs234_.f3 = rhs265_
                r2 = rhs266_
                lhs235_.f3 = rhs267_
                d_1552_v25_ = rhs268_
            elif True:
                d_1554_v27_: _dafny.Array
                def lambda77_(d_1555_v3_):
                    def lambda78_(d_1556_i6_):
                        return (d_1555_v3_) + (d_1555_v3_)

                    return lambda78_

                init46_ = lambda77_(d_1526_v3_)
                nw257_ = _dafny.Array(None, 11)
                for i0_46_ in range(nw257_.length(0)):
                    nw257_[i0_46_] = init46_(i0_46_)
                d_1554_v27_ = nw257_
                d_1554_v27_ = d_1554_v27_
                index288_ = default__.safeIndex(928, (d_1532_v9_).length(0))
                (d_1532_v9_)[index288_] = not(not(not(True)))
                (self).f14 = not((p1) <= ((self).f18))
                (globalState).f3 = p1
                d_1511_v0_ = d_1511_v0_
            d_1557_v28_: _dafny.Seq
            d_1557_v28_ = _dafny.SeqWithoutIsStrInference([d_1532_v9_])
            d_1558_v29_: D1
            d_1558_v29_ = D1_DC5((d_1532_v9_)[default__.safeIndex(928, (d_1532_v9_).length(0))], (self).fm2(False, self.f14, globalState))
            d_1559_v30_: D11
            d_1559_v30_ = D11_DC34((self).f15, d_1536_v12_, d_1532_v9_, d_1558_v29_)
            d_1560_v31_: _dafny.Array
            nw258_ = _dafny.Array(None, 18)
            nw258_[int(0)] = d_1532_v9_
            nw258_[int(1)] = d_1532_v9_
            nw258_[int(2)] = d_1532_v9_
            nw258_[int(3)] = d_1532_v9_
            nw258_[int(4)] = d_1532_v9_
            nw258_[int(5)] = (d_1557_v28_)[default__.safeIndex(p1, len(d_1557_v28_))]
            nw258_[int(6)] = d_1532_v9_
            nw258_[int(7)] = d_1532_v9_
            nw258_[int(8)] = d_1532_v9_
            nw258_[int(9)] = d_1532_v9_
            nw258_[int(10)] = d_1532_v9_
            nw258_[int(11)] = d_1532_v9_
            nw258_[int(12)] = d_1532_v9_
            nw258_[int(13)] = d_1532_v9_
            nw258_[int(14)] = d_1532_v9_
            nw258_[int(15)] = (d_1559_v30_).cf60
            nw258_[int(16)] = d_1532_v9_
            nw258_[int(17)] = d_1532_v9_
            d_1560_v31_ = nw258_
            index289_ = default__.safeIndex(777, (d_1560_v31_).length(0))
            (d_1560_v31_)[index289_] = d_1532_v9_
            d_1561_v32_: _dafny.Map
            d_1561_v32_ = _dafny.Map({(d_1545_v20_)[default__.safeIndex(812, (d_1545_v20_).length(0))]: d_1527_v4_})
            index290_ = default__.safeIndex(777, (d_1560_v31_).length(0))
            index291_ = default__.safeIndex(928, (d_1532_v9_).length(0))
            rhs269_ = (d_1532_v9_)[default__.safeIndex(928, (d_1532_v9_).length(0))]
            rhs270_ = d_1532_v9_
            rhs271_ = (((self).f15) - ((self).f15)).isdisjoint((self).f15)
            rhs272_ = d_1561_v32_
            lhs236_ = d_1560_v31_
            lhs237_ = default__.safeIndex(777, (d_1560_v31_).length(0))
            lhs238_ = d_1532_v9_
            lhs239_ = default__.safeIndex(928, (d_1532_v9_).length(0))
            r2 = rhs269_
            lhs236_[lhs237_] = rhs270_
            lhs238_[lhs239_] = rhs271_
            d_1561_v32_ = rhs272_
        d_1562_v33_: _dafny.Map
        d_1562_v33_ = _dafny.Map({(d_1511_v0_)[default__.safeIndex(46, (d_1511_v0_).length(0))]: d_1532_v9_})
        d_1563_v34_: _dafny.MultiSet
        d_1563_v34_ = _dafny.MultiSet([d_1527_v4_, d_1527_v4_])
        d_1564_v35_: _dafny.Seq
        d_1564_v35_ = _dafny.SeqWithoutIsStrInference([d_1532_v9_])
        d_1565_v36_: _dafny.Map
        d_1565_v36_ = _dafny.Map({d_1529_v6_: d_1532_v9_})
        d_1566_v37_: _dafny.Map
        d_1566_v37_ = _dafny.Map({(d_1532_v9_)[default__.safeIndex(928, (d_1532_v9_).length(0))]: (self).f18})
        d_1567_v38_: _dafny.Array
        nw259_ = _dafny.Array(None, 24)
        nw259_[int(0)] = d_1532_v9_
        nw259_[int(1)] = d_1532_v9_
        nw259_[int(2)] = d_1532_v9_
        nw259_[int(3)] = d_1532_v9_
        nw259_[int(4)] = d_1532_v9_
        nw259_[int(5)] = ((d_1562_v33_)[(d_1528_v5_).fm17((self).f18, (d_1528_v5_).f19, d_1563_v34_, (d_1511_v0_)[default__.safeIndex(46, (d_1511_v0_).length(0))], globalState)] if ((d_1528_v5_).fm17((self).f18, (d_1528_v5_).f19, d_1563_v34_, (d_1511_v0_)[default__.safeIndex(46, (d_1511_v0_).length(0))], globalState)) in (d_1562_v33_) else d_1532_v9_)
        nw259_[int(6)] = d_1532_v9_
        nw259_[int(7)] = d_1532_v9_
        nw259_[int(8)] = (d_1564_v35_)[default__.safeIndex((d_1511_v0_)[default__.safeIndex(46, (d_1511_v0_).length(0))], len(d_1564_v35_))]
        nw259_[int(9)] = d_1532_v9_
        nw259_[int(10)] = d_1532_v9_
        nw259_[int(11)] = d_1532_v9_
        nw259_[int(12)] = d_1532_v9_
        nw259_[int(13)] = d_1532_v9_
        nw259_[int(14)] = d_1532_v9_
        nw259_[int(15)] = d_1532_v9_
        nw259_[int(16)] = d_1532_v9_
        nw259_[int(17)] = d_1532_v9_
        nw259_[int(18)] = ((d_1565_v36_)[default__.fm25(p1, p1, (d_1528_v5_).f19, d_1566_v37_, globalState)] if (default__.fm25(p1, p1, (d_1528_v5_).f19, d_1566_v37_, globalState)) in (d_1565_v36_) else d_1532_v9_)
        nw259_[int(19)] = d_1532_v9_
        nw259_[int(20)] = d_1532_v9_
        nw259_[int(21)] = d_1532_v9_
        nw259_[int(22)] = d_1532_v9_
        nw259_[int(23)] = d_1532_v9_
        d_1567_v38_ = nw259_
        d_1568_v39_: D4
        d_1568_v39_ = D4_DC12(d_1532_v9_)
        d_1569_v40_: _dafny.Map
        d_1569_v40_ = _dafny.Map({d_1568_v39_: d_1567_v38_})
        d_1570_v41_: _dafny.Seq
        d_1570_v41_ = _dafny.SeqWithoutIsStrInference([d_1567_v38_])
        d_1567_v38_ = ((d_1569_v40_)[d_1568_v39_] if (d_1568_v39_) in (d_1569_v40_) else (d_1570_v41_)[default__.safeIndex((d_1526_v3_)[default__.safeIndex(p1, len(d_1526_v3_))], len(d_1570_v41_))])
        r0 = 561
        r1 = p1
        r2 = ((0) - (p1)) > ((d_1511_v0_)[default__.safeIndex(46, (d_1511_v0_).length(0))])
        return r0, r1, r2

    @property
    def f18(self):
        return self._f18

class C6(T0):
    def  __init__(self):
        self._f14: bool = False
        self._f15: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    @property
    def f15(self):
        return self._f15
    def ctor__(self, f14, f15):
        (self).f14 = f14
        (self)._f15 = f15

    def fm2(self, p0, p1, globalState):
        def iife177_():
            coll61_ = _dafny.Set()
            compr_61_: int
            for compr_61_ in _dafny.IntegerRange(490, 573):
                d_1571_v0_: int = compr_61_
                if ((490) <= (d_1571_v0_)) and ((d_1571_v0_) < (573)):
                    coll61_ = coll61_.union(_dafny.Set([default__.safeDivisionInt(d_1571_v0_, 621)]))
            return _dafny.Set(coll61_)
        def iife178_():
            coll62_ = _dafny.Set()
            compr_62_: int
            for compr_62_ in (_dafny.Set({-224})).Elements:
                d_1572_v1_: int = compr_62_
                if (d_1572_v1_) in (_dafny.Set({-224})):
                    coll62_ = coll62_.union(_dafny.Set([(d_1572_v1_) * (544)]))
            return _dafny.Set(coll62_)
        def iife179_():
            coll63_ = _dafny.Set()
            compr_63_: int
            for compr_63_ in _dafny.IntegerRange(613, 145):
                d_1573_v2_: int = compr_63_
                if ((613) <= (d_1573_v2_)) and ((d_1573_v2_) < (145)):
                    coll63_ = coll63_.union(_dafny.Set([(d_1573_v2_) - (943)]))
            return _dafny.Set(coll63_)
        return (iife177_()
        ) != ((iife178_()
        ) - (iife179_()
        ))

    def fm3(self, p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evj"))

    def fm11(self, p0, p1, globalState):
        def iife180_():
            coll64_ = _dafny.Set()
            compr_64_: int
            for compr_64_ in (_dafny.MultiSet([(_dafny.MultiSet([len(_dafny.Map({857: (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bv")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1574_i0_ in range(default__.abs(-494))])])).cardinality}))])).cardinality, 339])).Elements:
                d_1575_v0_: int = compr_64_
                if (d_1575_v0_) in (_dafny.MultiSet([(_dafny.MultiSet([len(_dafny.Map({857: (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bv")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1574_i0_ in range(default__.abs(-494))])])).cardinality}))])).cardinality, 339])):
                    coll64_ = coll64_.union(_dafny.Set([(d_1575_v0_) - (177)]))
            return _dafny.Set(coll64_)
        return (self.f14) and ((iife180_()
        ).issubset(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([864, 476, 180, (D1_DC3(646)).cf3, 305])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njb")))})))

    def fm12(self, p0, p1, globalState):
        def iife181_():
            def iife183_():
                coll67_ = _dafny.Map()
                compr_67_: D1
                for compr_67_ in (_dafny.SeqWithoutIsStrInference([D1_DC3(-509)])).Elements:
                    d_1576_v1_: D1 = compr_67_
                    if (d_1576_v1_) in (_dafny.SeqWithoutIsStrInference([D1_DC3(-509)])):
                        coll67_[d_1576_v1_] = 121
                return _dafny.Map(coll67_)
            coll65_ = _dafny.Map()
            def iife182_():
                coll66_ = _dafny.Map()
                compr_66_: D1
                for compr_66_ in (_dafny.SeqWithoutIsStrInference([D1_DC3(-509)])).Elements:
                    d_1576_v1_: D1 = compr_66_
                    if (d_1576_v1_) in (_dafny.SeqWithoutIsStrInference([D1_DC3(-509)])):
                        coll66_[d_1576_v1_] = 121
                return _dafny.Map(coll66_)
            compr_65_: D1
            for compr_65_ in (iife182_()
            ).keys.Elements:
                d_1577_v0_: D1 = compr_65_
                if (d_1577_v0_) in (iife183_()
                ):
                    coll65_[d_1577_v0_] = (self).f15
            return _dafny.Map(coll65_)
        return ((_dafny.Map({D1_DC3(len(_dafny.SeqWithoutIsStrInference([self.f14]))): (self).f15})) | (_dafny.Map({D1_DC3(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aksalss")))): (self).f15}))) == (iife181_()
        )

    def m0(self, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        r3: bool = False
        d_1578_v0_: int
        d_1578_v0_ = -804
        hi6_ = d_1578_v0_
        for d_1579_i0_ in range(d_1578_v0_, hi6_):
            d_1580_v1_: _dafny.Seq
            d_1580_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "li"))
            d_1581_v2_: _dafny.Map
            d_1581_v2_ = _dafny.Map({len((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pf")) for d_1582_i1_ in range(default__.abs(878))])).set(default__.safeIndex(d_1578_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pf")) for d_1583_i1_ in range(default__.abs(878))]))), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1584_i2_ in range(default__.abs(343))]))): d_1580_v1_})
            d_1581_v2_ = (d_1581_v2_).set(d_1578_v0_, (d_1580_v1_ if self.f14 else d_1580_v1_))
            d_1585_v3_: D0
            d_1585_v3_ = D0_DC0(self.f14)
            if ((d_1585_v3_ if (d_1585_v3_).cf0 else d_1585_v3_)).cf0:
                d_1586_v4_: _dafny.Seq
                d_1586_v4_ = _dafny.SeqWithoutIsStrInference([len(d_1580_v1_)])
                d_1587_v5_: _dafny.MultiSet
                d_1587_v5_ = _dafny.MultiSet([d_1579_i0_])
                d_1588_v6_: _dafny.MultiSet
                d_1588_v6_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([default__.fm0(self.f14, d_1579_i0_, len(d_1586_v4_), d_1587_v5_, globalState)]) if self.f14 else _dafny.SeqWithoutIsStrInference([d_1579_i0_])), d_1586_v4_, (d_1586_v4_).set(default__.safeIndex(d_1578_v0_, len(d_1586_v4_)), d_1578_v0_)])
                d_1588_v6_ = default__.fm13((self).f15, (488) * (d_1579_i0_), globalState)
                d_1589_v7_: _dafny.Array
                nw260_ = _dafny.Array(_dafny.Seq({}), 25)
                d_1589_v7_ = nw260_
                d_1590_v8_: _dafny.Array
                nw261_ = _dafny.Array(False, 23)
                d_1590_v8_ = nw261_
                d_1591_v9_: _dafny.Seq
                d_1591_v9_ = _dafny.SeqWithoutIsStrInference([d_1590_v8_, d_1590_v8_, d_1590_v8_])
                index292_ = default__.safeIndex(777, (d_1589_v7_).length(0))
                (d_1589_v7_)[index292_] = d_1591_v9_
                index293_ = default__.safeIndex(777, (d_1589_v7_).length(0))
                (d_1589_v7_)[index293_] = _dafny.SeqWithoutIsStrInference([d_1590_v8_])
                d_1592_v10_: _dafny.Map
                d_1592_v10_ = _dafny.Map({(_dafny.MultiSet([False])) == ((self).f15): d_1586_v4_})
                d_1592_v10_ = (d_1592_v10_).set(self.f14, d_1586_v4_)
                d_1593_v11_: int
                d_1594_v12_: int
                d_1595_v13_: _dafny.MultiSet
                d_1596_v14_: bool
                out47_: int
                out48_: int
                out49_: _dafny.MultiSet
                out50_: bool
                out47_, out48_, out49_, out50_ = (self).m5(globalState)
                d_1593_v11_ = out47_
                d_1594_v12_ = out48_
                d_1595_v13_ = out49_
                d_1596_v14_ = out50_
                d_1597_v15_: _dafny.Seq
                d_1597_v15_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1598_i3_ in range(default__.abs(512))])) == (d_1580_v1_), True, d_1596_v14_, d_1596_v14_, d_1596_v14_])
                rhs273_ = (d_1597_v15_) + (d_1597_v15_)
                rhs274_ = d_1590_v8_
                rhs275_ = self.f14
                lhs240_ = globalState
                d_1597_v15_ = rhs273_
                d_1590_v8_ = rhs274_
                lhs240_.f9 = rhs275_
            elif True:
                d_1599_v16_: _dafny.Set
                d_1599_v16_ = _dafny.Set({self.f14})
                d_1600_v17_: _dafny.Seq
                d_1600_v17_ = _dafny.SeqWithoutIsStrInference([d_1599_v16_])
                d_1601_v18_: _dafny.MultiSet
                d_1601_v18_ = _dafny.MultiSet([d_1579_i0_])
                d_1602_v19_: _dafny.MultiSet
                d_1602_v19_ = _dafny.MultiSet([default__.fm0(self.f14, (0) - ((d_1601_v18_).cardinality), d_1579_i0_, d_1601_v18_, globalState), -334, d_1579_i0_, (0) - (d_1578_v0_), d_1579_i0_])
                d_1603_v20_: _dafny.Map
                d_1603_v20_ = _dafny.Map({(d_1600_v17_) == (_dafny.SeqWithoutIsStrInference([d_1599_v16_, d_1599_v16_])): (d_1602_v19_).ispropersubset(d_1601_v18_)})
                d_1604_v21_: _dafny.Map
                d_1604_v21_ = _dafny.Map({default__.fm14(globalState): d_1578_v0_})
                d_1605_v22_: _dafny.Seq
                d_1605_v22_ = _dafny.SeqWithoutIsStrInference([d_1604_v21_])
                d_1603_v20_ = (d_1603_v20_).set(not((self).fm12(self.f14, (d_1605_v22_)[default__.safeIndex(d_1578_v0_, len(d_1605_v22_))], globalState)), self.f14)
                d_1606_v23_: _dafny.Map
                d_1606_v23_ = _dafny.Map({d_1580_v1_: D0_DC0(self.f14)})
                d_1607_v24_: _dafny.Seq
                d_1607_v24_ = _dafny.SeqWithoutIsStrInference([self.f14, (self).fm12(self.f14, default__.fm48(d_1579_i0_, d_1578_v0_, self.f14, globalState), globalState), self.f14])
                d_1608_v25_: T0
                nw262_ = C5()
                nw262_.ctor__(d_1578_v0_, True, _dafny.MultiSet(d_1607_v24_))
                d_1608_v25_ = nw262_
                d_1609_v26_: _dafny.Map
                d_1609_v26_ = _dafny.Map({d_1608_v25_: d_1580_v1_})
                d_1610_v27_: _dafny.Array
                nw263_ = _dafny.Array(False, 8)
                d_1610_v27_ = nw263_
                d_1606_v23_ = (d_1606_v23_).set((((d_1609_v26_)[d_1608_v25_] if (d_1608_v25_) in (d_1609_v26_) else (D4_DC13(d_1580_v1_, d_1578_v0_, self.f14, d_1610_v27_)).cf22)) + (d_1580_v1_), d_1585_v3_)
                d_1611_v28_: _dafny.Map
                d_1611_v28_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvdqndys")): d_1579_i0_})
                r3 = ((d_1580_v1_) + (d_1580_v1_)) not in (d_1611_v28_)
                r0 = d_1578_v0_
                d_1612_v29_: _dafny.Seq
                d_1612_v29_ = _dafny.SeqWithoutIsStrInference([d_1603_v20_])
                d_1613_v30_: _dafny.MultiSet
                d_1613_v30_ = _dafny.MultiSet([d_1603_v20_])
                d_1614_v31_: _dafny.Seq
                d_1614_v31_ = _dafny.SeqWithoutIsStrInference([d_1613_v30_])
                d_1615_v32_: str
                d_1615_v32_ = _dafny.CodePoint('b')
                d_1616_v33_: _dafny.Map
                d_1616_v33_ = _dafny.Map({d_1579_i0_: d_1615_v32_})
                d_1617_v34_: _dafny.Array
                nw264_ = _dafny.Array(None, 19)
                nw264_[int(0)] = (_dafny.MultiSet(d_1612_v29_)) | (d_1613_v30_)
                nw264_[int(1)] = (d_1613_v30_).intersection(d_1613_v30_)
                nw264_[int(2)] = d_1613_v30_
                nw264_[int(3)] = d_1613_v30_
                nw264_[int(4)] = d_1613_v30_
                nw264_[int(5)] = (d_1614_v31_)[default__.safeIndex(len(d_1616_v33_), len(d_1614_v31_))]
                nw264_[int(6)] = d_1613_v30_
                nw264_[int(7)] = _dafny.MultiSet((d_1612_v29_).set(default__.safeIndex(len(d_1599_v16_), len(d_1612_v29_)), d_1603_v20_))
                nw264_[int(8)] = d_1613_v30_
                nw264_[int(9)] = (d_1613_v30_) - (d_1613_v30_)
                nw264_[int(10)] = d_1613_v30_
                nw264_[int(11)] = d_1613_v30_
                nw264_[int(12)] = d_1613_v30_
                nw264_[int(13)] = d_1613_v30_
                nw264_[int(14)] = d_1613_v30_
                nw264_[int(15)] = d_1613_v30_
                nw264_[int(16)] = d_1613_v30_
                nw264_[int(17)] = _dafny.MultiSet([d_1603_v20_, d_1603_v20_])
                nw264_[int(18)] = (d_1613_v30_).set(_dafny.Map({False: d_1608_v25_.f14}), default__.abs(len(d_1607_v24_)))
                d_1617_v34_ = nw264_
                index294_ = default__.safeIndex(765, (d_1617_v34_).length(0))
                (d_1617_v34_)[index294_] = (d_1613_v30_) | (d_1613_v30_)
                index295_ = default__.safeIndex(765, (d_1617_v34_).length(0))
                (d_1617_v34_)[index295_] = _dafny.MultiSet(d_1612_v29_)
            if self.f14:
                d_1618_v35_: _dafny.Array
                nw265_ = _dafny.Array(None, 27)
                d_1618_v35_ = nw265_
                d_1618_v35_ = d_1618_v35_
                (globalState).f9 = self.f14
                d_1619_v36_: _dafny.Seq
                d_1619_v36_ = _dafny.SeqWithoutIsStrInference([d_1578_v0_])
                (globalState).f3 = default__.safeModuloInt(d_1579_i0_, len(d_1619_v36_))
                d_1620_v37_: _dafny.Set
                d_1620_v37_ = _dafny.Set({(self).f15})
                d_1621_v38_: _dafny.Map
                d_1621_v38_ = _dafny.Map({D11_DC33(d_1620_v37_): d_1578_v0_})
                d_1622_v39_: D11
                d_1622_v39_ = D11_DC33(d_1620_v37_)
                d_1621_v38_ = (d_1621_v38_).set(d_1622_v39_, ((0) - (d_1579_i0_)) * (d_1578_v0_))
                d_1623_v40_: C1
                nw266_ = C1()
                nw266_.ctor__()
                d_1623_v40_ = nw266_
            elif True:
                d_1624_v41_: _dafny.Array
                nw267_ = _dafny.Array(None, 2)
                nw267_[int(0)] = d_1579_i0_
                nw267_[int(1)] = d_1578_v0_
                d_1624_v41_ = nw267_
                index296_ = default__.safeIndex(4, (d_1624_v41_).length(0))
                (d_1624_v41_)[index296_] = d_1578_v0_
                d_1625_v42_: _dafny.Array
                nw268_ = _dafny.Array(_dafny.CodePoint('D'), 12)
                d_1625_v42_ = nw268_
                d_1626_v43_: str
                d_1626_v43_ = _dafny.CodePoint('t')
                index297_ = default__.safeIndex(347, (d_1625_v42_).length(0))
                (d_1625_v42_)[index297_] = d_1626_v43_
                index298_ = default__.safeIndex(4, (d_1624_v41_).length(0))
                index299_ = default__.safeIndex(347, (d_1625_v42_).length(0))
                rhs276_ = (d_1578_v0_) + (d_1579_i0_)
                rhs277_ = d_1579_i0_
                rhs278_ = d_1626_v43_
                rhs279_ = (0) - (d_1579_i0_)
                lhs241_ = globalState
                lhs242_ = d_1624_v41_
                lhs243_ = default__.safeIndex(4, (d_1624_v41_).length(0))
                lhs244_ = d_1625_v42_
                lhs245_ = default__.safeIndex(347, (d_1625_v42_).length(0))
                lhs241_.f3 = rhs276_
                lhs242_[lhs243_] = rhs277_
                lhs244_[lhs245_] = rhs278_
                d_1578_v0_ = rhs279_
                d_1627_v45_: _dafny.Array
                def lambda79_(d_1628_v0_, d_1629_i0_):
                    def lambda80_(d_1630_i4_):
                        def iife184_():
                            coll68_ = _dafny.Set()
                            compr_68_: int
                            for compr_68_ in (_dafny.Set({d_1629_i0_})).Elements:
                                d_1631_v44_: int = compr_68_
                                if (d_1631_v44_) in (_dafny.Set({d_1629_i0_})):
                                    coll68_ = coll68_.union(_dafny.Set([(d_1631_v44_) - (274)]))
                            return _dafny.Set(coll68_)
                        return (iife184_()
                        ).isdisjoint(_dafny.Set({d_1628_v0_}))

                    return lambda80_

                init47_ = lambda79_(d_1578_v0_, d_1579_i0_)
                nw269_ = _dafny.Array(None, 3)
                for i0_47_ in range(nw269_.length(0)):
                    nw269_[i0_47_] = init47_(i0_47_)
                d_1627_v45_ = nw269_
                index300_ = default__.safeIndex(279, (d_1627_v45_).length(0))
                (d_1627_v45_)[index300_] = True
                d_1632_v46_: _dafny.Seq
                d_1632_v46_ = _dafny.SeqWithoutIsStrInference([self.f14, (self.f14) or (True), self.f14])
                d_1633_v47_: _dafny.Set
                d_1633_v47_ = _dafny.Set({self.f14})
                index301_ = default__.safeIndex(279, (d_1627_v45_).length(0))
                rhs280_ = self.f14
                rhs281_ = (d_1632_v46_)[default__.safeIndex(len((_dafny.Set({self.f14, True}) if self.f14 else d_1633_v47_)), len(d_1632_v46_))]
                rhs282_ = self.f14
                lhs246_ = d_1627_v45_
                lhs247_ = default__.safeIndex(279, (d_1627_v45_).length(0))
                r3 = rhs280_
                r1 = rhs281_
                lhs246_[lhs247_] = rhs282_
                d_1634_v48_: _dafny.Map
                d_1634_v48_ = _dafny.Map({(d_1625_v42_)[default__.safeIndex(347, (d_1625_v42_).length(0))]: 891})
                d_1635_v49_: _dafny.Seq
                d_1635_v49_ = _dafny.SeqWithoutIsStrInference([d_1578_v0_, d_1579_i0_])
                d_1636_v50_: _dafny.Set
                d_1636_v50_ = _dafny.Set({d_1635_v49_, d_1635_v49_, (d_1635_v49_).set(default__.safeIndex((d_1624_v41_)[default__.safeIndex(4, (d_1624_v41_).length(0))], len(d_1635_v49_)), (d_1624_v41_)[default__.safeIndex(4, (d_1624_v41_).length(0))]), d_1635_v49_})
                (globalState).f9 = not((default__.fm23((0) - (d_1578_v0_), self.f14, d_1634_v48_, d_1636_v50_, globalState) if True else False))
                d_1637_v51_: int
                d_1638_v52_: int
                d_1639_v53_: _dafny.MultiSet
                d_1640_v54_: bool
                out51_: int
                out52_: int
                out53_: _dafny.MultiSet
                out54_: bool
                out51_, out52_, out53_, out54_ = (self).m5(globalState)
                d_1637_v51_ = out51_
                d_1638_v52_ = out52_
                d_1639_v53_ = out53_
                d_1640_v54_ = out54_
                d_1580_v1_ = (_dafny.SeqWithoutIsStrInference([d_1626_v43_])) + (d_1580_v1_)
            d_1641_v55_: _dafny.Array
            nw270_ = _dafny.Array(int(0), 1)
            d_1641_v55_ = nw270_
            d_1642_v56_: _dafny.Map
            d_1642_v56_ = _dafny.Map({d_1578_v0_: d_1578_v0_})
            index302_ = default__.safeIndex(518, (d_1641_v55_).length(0))
            def iife185_():
                coll69_ = _dafny.Map()
                compr_69_: int
                for compr_69_ in _dafny.IntegerRange(964, -867):
                    d_1643_v57_: int = compr_69_
                    if ((964) <= (d_1643_v57_)) and ((d_1643_v57_) < (-867)):
                        coll69_[(d_1643_v57_) - (len(d_1580_v1_))] = d_1579_i0_
                return _dafny.Map(coll69_)
            (d_1641_v55_)[index302_] = len((d_1642_v56_) | ((iife185_()
            ).set((0) - (d_1578_v0_), d_1579_i0_)))
            d_1644_v58_: _dafny.Map
            d_1644_v58_ = _dafny.Map({d_1580_v1_: (0) - (d_1578_v0_)})
            d_1645_v59_: _dafny.MultiSet
            d_1645_v59_ = _dafny.MultiSet([len(d_1580_v1_)])
            d_1646_v60_: _dafny.Seq
            d_1646_v60_ = _dafny.SeqWithoutIsStrInference([d_1578_v0_])
            d_1647_v61_: _dafny.Set
            d_1647_v61_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([default__.fm0(self.f14, d_1579_i0_, d_1579_i0_, d_1645_v59_, globalState)]), d_1646_v60_, d_1646_v60_})
            index303_ = default__.safeIndex(518, (d_1641_v55_).length(0))
            (d_1641_v55_)[index303_] = default__.safeDivisionInt(((d_1644_v58_)[d_1580_v1_] if (d_1580_v1_) in (d_1644_v58_) else len((D13_DC41(d_1647_v61_)).cf67)), d_1579_i0_)
        d_1648_v62_: _dafny.Array
        nw271_ = _dafny.Array(_dafny.Set({}), 19)
        d_1648_v62_ = nw271_
        d_1649_v63_: _dafny.Set
        d_1649_v63_ = _dafny.Set({d_1578_v0_})
        index304_ = default__.safeIndex(906, (d_1648_v62_).length(0))
        (d_1648_v62_)[index304_] = d_1649_v63_
        index305_ = default__.safeIndex(906, (d_1648_v62_).length(0))
        (d_1648_v62_)[index305_] = d_1649_v63_
        d_1650_v64_: str
        d_1650_v64_ = _dafny.CodePoint('h')
        d_1651_v65_: _dafny.MultiSet
        d_1651_v65_ = _dafny.MultiSet([d_1650_v64_])
        d_1651_v65_ = (d_1651_v65_).intersection((d_1651_v65_) - (d_1651_v65_))
        d_1652_v66_: _dafny.Array
        nw272_ = _dafny.Array(False, 10)
        d_1652_v66_ = nw272_
        index306_ = default__.safeIndex(542, (d_1652_v66_).length(0))
        (d_1652_v66_)[index306_] = not(self.f14)
        index307_ = default__.safeIndex(542, (d_1652_v66_).length(0))
        (d_1652_v66_)[index307_] = self.f14
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1652_v66_).length(0)):
            d_1653_i5_: int = guard_loop_4_
            if (True) and (((0) <= (d_1653_i5_)) and ((d_1653_i5_) < ((d_1652_v66_).length(0)))):
                (d_1652_v66_)[(d_1653_i5_)] = ((d_1578_v0_) + (d_1578_v0_)) != (d_1578_v0_)
        d_1654_v67_: int
        d_1655_v68_: int
        d_1656_v69_: _dafny.MultiSet
        d_1657_v70_: bool
        out55_: int
        out56_: int
        out57_: _dafny.MultiSet
        out58_: bool
        out55_, out56_, out57_, out58_ = (self).m5(globalState)
        d_1654_v67_ = out55_
        d_1655_v68_ = out56_
        d_1656_v69_ = out57_
        d_1657_v70_ = out58_
        def iife186_():
            coll70_ = _dafny.Map()
            compr_70_: int
            for compr_70_ in (_dafny.SeqWithoutIsStrInference([d_1578_v0_ for d_1658_i6_ in range(default__.abs(363))])).Elements:
                d_1659_v71_: int = compr_70_
                if (d_1659_v71_) in (_dafny.SeqWithoutIsStrInference([d_1578_v0_ for d_1658_i6_ in range(default__.abs(363))])):
                    coll70_[default__.safeDivisionInt(d_1659_v71_, d_1655_v68_)] = d_1578_v0_
            return _dafny.Map(coll70_)
        r0 = ((d_1578_v0_) * (d_1655_v68_)) + (len(iife186_()
        ))
        d_1660_v72_: D1
        d_1660_v72_ = D1_DC3((0) - (d_1578_v0_))
        d_1661_v73_: D1
        d_1661_v73_ = D1_DC6(d_1660_v72_)
        pat_let_tv41_ = d_1650_v64_
        pat_let_tv42_ = d_1650_v64_
        def lambda81_(source16_):
            if source16_.is_DC4:
                d_1662___mcc_h0_ = source16_.cf4
                d_1663___mcc_h1_ = source16_.cf5
                d_1664_cf5_ = d_1663___mcc_h1_
                d_1665_cf4_ = d_1662___mcc_h0_
                return (D13_DC42(_dafny.CodePoint('r'), self.f14)).cf69
            elif source16_.is_DC5:
                d_1666___mcc_h2_ = source16_.cf6
                d_1667___mcc_h3_ = source16_.cf7
                d_1668_cf7_ = d_1667___mcc_h3_
                d_1669_cf6_ = d_1666___mcc_h2_
                return self.f14
            elif source16_.is_DC3:
                d_1670___mcc_h4_ = source16_.cf3
                d_1671_cf3_ = d_1670___mcc_h4_
                return self.f14
            elif True:
                d_1672___mcc_h5_ = source16_.cf8
                d_1673_cf8_ = d_1672___mcc_h5_
                return (_dafny.Set({pat_let_tv41_})).issubset(_dafny.Set({pat_let_tv42_}))

        r1 = lambda81_((d_1661_v73_ if True else d_1661_v73_))
        d_1674_v74_: _dafny.Map
        d_1674_v74_ = _dafny.Map({(d_1652_v66_)[default__.safeIndex(542, (d_1652_v66_).length(0))]: d_1654_v67_})
        d_1675_v75_: _dafny.Map
        d_1675_v75_ = _dafny.Map({_dafny.Map({((d_1674_v74_)[(d_1652_v66_)[default__.safeIndex(542, (d_1652_v66_).length(0))]] if ((d_1652_v66_)[default__.safeIndex(542, (d_1652_v66_).length(0))]) in (d_1674_v74_) else (0) - (d_1655_v68_)): (d_1652_v66_)[default__.safeIndex(542, (d_1652_v66_).length(0))]}): d_1578_v0_})
        r2 = d_1675_v75_
        r3 = d_1657_v70_
        return r0, r1, r2, r3

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        d_1676_v0_: _dafny.Array
        def lambda82_(d_1677_i0_):
            return self.f14

        init48_ = lambda82_
        nw273_ = _dafny.Array(None, 5)
        for i0_48_ in range(nw273_.length(0)):
            nw273_[i0_48_] = init48_(i0_48_)
        d_1676_v0_ = nw273_
        d_1678_v1_: _dafny.Seq
        d_1678_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i')])
        index308_ = default__.safeIndex(538, (d_1676_v0_).length(0))
        (d_1676_v0_)[index308_] = (len(d_1678_v1_)) > (p1)
        index309_ = default__.safeIndex(538, (d_1676_v0_).length(0))
        (d_1676_v0_)[index309_] = (self).fm2(self.f14, self.f14, globalState)
        (globalState).f9 = (d_1676_v0_)[default__.safeIndex(538, (d_1676_v0_).length(0))]
        d_1679_v2_: _dafny.Array
        nw274_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
        d_1679_v2_ = nw274_
        d_1680_v3_: _dafny.Array
        nw275_ = _dafny.Array(None, 17)
        nw275_[int(0)] = d_1679_v2_
        nw275_[int(1)] = d_1679_v2_
        nw275_[int(2)] = d_1679_v2_
        nw275_[int(3)] = d_1679_v2_
        nw275_[int(4)] = d_1679_v2_
        nw275_[int(5)] = d_1679_v2_
        nw275_[int(6)] = d_1679_v2_
        nw275_[int(7)] = d_1679_v2_
        nw275_[int(8)] = d_1679_v2_
        nw275_[int(9)] = d_1679_v2_
        nw275_[int(10)] = d_1679_v2_
        nw275_[int(11)] = d_1679_v2_
        nw275_[int(12)] = d_1679_v2_
        nw275_[int(13)] = d_1679_v2_
        nw275_[int(14)] = d_1679_v2_
        nw275_[int(15)] = d_1679_v2_
        nw275_[int(16)] = (D14_DC44(d_1679_v2_)).cf71
        d_1680_v3_ = nw275_
        index310_ = default__.safeIndex(2, (d_1680_v3_).length(0))
        (d_1680_v3_)[index310_] = d_1679_v2_
        index311_ = default__.safeIndex(2, (d_1680_v3_).length(0))
        (d_1680_v3_)[index311_] = (d_1679_v2_ if (D1_DC5(self.f14, True)).cf7 else (d_1679_v2_ if (d_1676_v0_)[default__.safeIndex(538, (d_1676_v0_).length(0))] else d_1679_v2_))
        d_1681_v4_: str
        d_1681_v4_ = _dafny.CodePoint('y')
        d_1682_v5_: _dafny.Set
        d_1682_v5_ = _dafny.Set({d_1681_v4_})
        d_1683_v6_: _dafny.MultiSet
        d_1683_v6_ = _dafny.MultiSet([p1])
        d_1682_v5_ = default__.fm32(d_1678_v1_, d_1683_v6_, globalState)
        d_1684_v7_: D4
        d_1684_v7_ = D4_DC12(d_1676_v0_)
        d_1685_v8_: _dafny.Seq
        d_1685_v8_ = _dafny.SeqWithoutIsStrInference([d_1684_v7_])
        d_1686_v9_: C4
        nw276_ = C4()
        nw276_.ctor__((d_1685_v8_) != (d_1685_v8_), (d_1676_v0_)[default__.safeIndex(538, (d_1676_v0_).length(0))], (self).f15)
        d_1686_v9_ = nw276_
        if (d_1686_v9_).f19:
            if (d_1686_v9_).f19:
                d_1676_v0_ = d_1676_v0_
                d_1687_v10_: _dafny.Seq
                d_1687_v10_ = _dafny.SeqWithoutIsStrInference([d_1678_v1_, d_1678_v1_, (d_1678_v1_) + (d_1678_v1_), d_1678_v1_])
                d_1687_v10_ = d_1687_v10_
                (globalState).f9 = (d_1686_v9_).f19
                d_1688_v11_: _dafny.Map
                d_1688_v11_ = _dafny.Map({p1: ((self).f15).cardinality})
                d_1689_v12_: _dafny.Map
                d_1689_v12_ = _dafny.Map({len(d_1678_v1_): d_1688_v11_})
                d_1690_v13_: _dafny.Map
                d_1690_v13_ = _dafny.Map({_dafny.Map({p1: not(not((d_1686_v9_).f19))}): p1})
                d_1691_v14_: D7
                d_1691_v14_ = D7_DC21(d_1690_v13_)
                d_1692_v15_: _dafny.Map
                d_1692_v15_ = _dafny.Map({d_1691_v14_: d_1682_v5_})
                d_1693_v16_: _dafny.Seq
                d_1693_v16_ = _dafny.SeqWithoutIsStrInference([d_1689_v12_, d_1689_v12_])
                rhs283_ = (((d_1692_v15_)[d_1691_v14_] if (d_1691_v14_) in (d_1692_v15_) else d_1682_v5_)).intersection(_dafny.Set({_dafny.CodePoint('f')}))
                rhs284_ = p1
                rhs285_ = not((p1) <= (len(_dafny.SeqWithoutIsStrInference([(0) - (default__.fm0((d_1686_v9_).f19, (0) - (p1), p1, d_1683_v6_, globalState)), p1]))))
                rhs286_ = (d_1693_v16_)[default__.safeIndex(default__.safeModuloInt(p1, p1), len(d_1693_v16_))]
                lhs248_ = globalState
                d_1682_v5_ = rhs283_
                lhs248_.f3 = rhs284_
                r2 = rhs285_
                d_1689_v12_ = rhs286_
                d_1694_v17_: C1
                nw277_ = C1()
                nw277_.ctor__()
                d_1694_v17_ = nw277_
            elif True:
                d_1695_v18_: D4
                d_1695_v18_ = D4_DC13(d_1678_v1_, p1, (d_1676_v0_)[default__.safeIndex(538, (d_1676_v0_).length(0))], d_1676_v0_)
                rhs287_ = ((d_1695_v18_).cf22) + (_dafny.SeqWithoutIsStrInference([d_1681_v4_ for d_1696_i1_ in range(default__.abs(292))]))
                rhs288_ = d_1678_v1_
                d_1678_v1_ = rhs287_
                d_1678_v1_ = rhs288_
                d_1697_v19_: _dafny.Map
                d_1697_v19_ = _dafny.Map({(d_1686_v9_).f19: p1})
                d_1698_v20_: _dafny.Array
                nw278_ = _dafny.Array(None, 19)
                nw278_[int(0)] = d_1681_v4_
                nw278_[int(1)] = d_1681_v4_
                nw278_[int(2)] = d_1681_v4_
                nw278_[int(3)] = d_1681_v4_
                nw278_[int(4)] = d_1681_v4_
                nw278_[int(5)] = d_1681_v4_
                nw278_[int(6)] = d_1681_v4_
                nw278_[int(7)] = (d_1678_v1_)[default__.safeIndex((d_1683_v6_).cardinality, len(d_1678_v1_))]
                nw278_[int(8)] = d_1681_v4_
                nw278_[int(9)] = d_1681_v4_
                nw278_[int(10)] = default__.fm25(p1, 129, self.f14, d_1697_v19_, globalState)
                nw278_[int(11)] = d_1681_v4_
                nw278_[int(12)] = d_1681_v4_
                nw278_[int(13)] = _dafny.CodePoint('t')
                nw278_[int(14)] = _dafny.CodePoint('g')
                nw278_[int(15)] = d_1681_v4_
                nw278_[int(16)] = d_1681_v4_
                nw278_[int(17)] = d_1681_v4_
                nw278_[int(18)] = d_1681_v4_
                d_1698_v20_ = nw278_
                index312_ = default__.safeIndex(774, (d_1698_v20_).length(0))
                (d_1698_v20_)[index312_] = d_1681_v4_
                index313_ = default__.safeIndex(774, (d_1698_v20_).length(0))
                (d_1698_v20_)[index313_] = d_1681_v4_
                d_1699_v21_: _dafny.Array
                nw279_ = _dafny.Array(_dafny.Set({}), 21)
                d_1699_v21_ = nw279_
                d_1700_v22_: _dafny.Set
                d_1700_v22_ = _dafny.Set({D8_DC25()})
                index314_ = default__.safeIndex(949, (d_1699_v21_).length(0))
                (d_1699_v21_)[index314_] = d_1700_v22_
                index315_ = default__.safeIndex(949, (d_1699_v21_).length(0))
                (d_1699_v21_)[index315_] = d_1700_v22_
                d_1701_v24_: _dafny.Map
                def iife187_():
                    coll71_ = _dafny.Set()
                    compr_71_: int
                    for compr_71_ in _dafny.IntegerRange(636, 181):
                        d_1702_v23_: int = compr_71_
                        if ((636) <= (d_1702_v23_)) and ((d_1702_v23_) < (181)):
                            coll71_ = coll71_.union(_dafny.Set([default__.safeDivisionInt(d_1702_v23_, p1)]))
                    return _dafny.Set(coll71_)
                d_1701_v24_ = _dafny.Map({((d_1698_v20_)[default__.safeIndex(774, (d_1698_v20_).length(0))]) not in (d_1678_v1_): (p0).intersection(iife187_()
                )})
                d_1701_v24_ = d_1701_v24_
                d_1703_v25_: _dafny.Map
                d_1703_v25_ = _dafny.Map({d_1678_v1_: ((self).f15).cardinality})
                d_1703_v25_ = (d_1703_v25_).set(d_1678_v1_, (len(d_1678_v1_)) - ((0) - (p1)))
            d_1704_v26_: _dafny.Map
            d_1704_v26_ = _dafny.Map({(d_1686_v9_).f19: not((p1) <= (p1))})
            d_1704_v26_ = (d_1704_v26_).set(self.f14, (p1) < (p1))
            d_1678_v1_ = _dafny.SeqWithoutIsStrInference([d_1681_v4_ for d_1705_i2_ in range(default__.abs(-802))])
            d_1706_v27_: _dafny.Array
            nw280_ = _dafny.Array(None, 5)
            nw280_[int(0)] = p0
            nw280_[int(1)] = p0
            nw280_[int(2)] = p0
            nw280_[int(3)] = p0
            nw280_[int(4)] = (p0) | (p0)
            d_1706_v27_ = nw280_
            d_1707_v28_: _dafny.Map
            d_1707_v28_ = _dafny.Map({p1: (d_1686_v9_).f19})
            d_1708_v29_: _dafny.Map
            d_1708_v29_ = _dafny.Map({len(d_1707_v28_): p1})
            d_1709_v30_: _dafny.Seq
            d_1709_v30_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1710_v31_: _dafny.Seq
            d_1710_v31_ = _dafny.SeqWithoutIsStrInference([len(d_1708_v29_), p1, default__.fm0(self.f14, p1, p1, d_1683_v6_, globalState), len(d_1709_v30_), p1])
            index316_ = default__.safeIndex(338, (d_1706_v27_).length(0))
            def iife188_():
                coll72_ = _dafny.Set()
                compr_72_: int
                for compr_72_ in (d_1709_v30_).Elements:
                    d_1711_v32_: int = compr_72_
                    if (d_1711_v32_) in (d_1709_v30_):
                        coll72_ = coll72_.union(_dafny.Set([(d_1711_v32_) + (876)]))
                return _dafny.Set(coll72_)
            (d_1706_v27_)[index316_] = (_dafny.Set({p1, p1})) | (iife188_()
            )
            d_1712_v33_: _dafny.Array
            nw281_ = _dafny.Array(int(0), 14)
            d_1712_v33_ = nw281_
            index317_ = default__.safeIndex(453, (d_1712_v33_).length(0))
            (d_1712_v33_)[index317_] = ((d_1708_v29_)[p1] if (p1) in (d_1708_v29_) else p1)
            d_1713_v34_: _dafny.Set
            d_1713_v34_ = _dafny.Set({d_1678_v1_})
            index318_ = default__.safeIndex(338, (d_1706_v27_).length(0))
            index319_ = default__.safeIndex(453, (d_1712_v33_).length(0))
            rhs289_ = _dafny.Set({p1})
            rhs290_ = p1
            rhs291_ = p1
            rhs292_ = _dafny.SeqWithoutIsStrInference([d_1681_v4_ for d_1714_i3_ in range(default__.abs(-551))])
            rhs293_ = d_1713_v34_
            lhs249_ = d_1706_v27_
            lhs250_ = default__.safeIndex(338, (d_1706_v27_).length(0))
            lhs251_ = d_1712_v33_
            lhs252_ = default__.safeIndex(453, (d_1712_v33_).length(0))
            lhs249_[lhs250_] = rhs289_
            r1 = rhs290_
            lhs251_[lhs252_] = rhs291_
            d_1678_v1_ = rhs292_
            d_1713_v34_ = rhs293_
            d_1715_v35_: _dafny.Map
            d_1715_v35_ = _dafny.Map({_dafny.Map({(d_1676_v0_)[default__.safeIndex(538, (d_1676_v0_).length(0))]: (d_1712_v33_)[default__.safeIndex(453, (d_1712_v33_).length(0))]}): d_1712_v33_})
            d_1716_v36_: _dafny.Map
            d_1716_v36_ = _dafny.Map({self.f14: (d_1712_v33_)[default__.safeIndex(453, (d_1712_v33_).length(0))]})
            d_1717_v37_: _dafny.Seq
            d_1717_v37_ = _dafny.SeqWithoutIsStrInference([((d_1715_v35_)[d_1716_v36_] if (d_1716_v36_) in (d_1715_v35_) else d_1712_v33_), d_1712_v33_])
            d_1718_v38_: D5
            d_1718_v38_ = D5_DC16(d_1712_v33_)
            index320_ = default__.safeIndex(538, (d_1676_v0_).length(0))
            rhs294_ = (D0_DC0((d_1676_v0_)[default__.safeIndex(538, (d_1676_v0_).length(0))])).cf0
            rhs295_ = _dafny.SeqWithoutIsStrInference([d_1712_v33_, d_1712_v33_, d_1712_v33_, (d_1718_v38_).cf29, d_1712_v33_])
            rhs296_ = (d_1712_v33_)[default__.safeIndex(453, (d_1712_v33_).length(0))]
            rhs297_ = not(True)
            lhs253_ = d_1676_v0_
            lhs254_ = default__.safeIndex(538, (d_1676_v0_).length(0))
            lhs255_ = globalState
            lhs253_[lhs254_] = rhs294_
            d_1717_v37_ = rhs295_
            r1 = rhs296_
            lhs255_.f9 = rhs297_
        elif True:
            d_1681_v4_ = d_1681_v4_
            d_1719_v39_: _dafny.Array
            nw282_ = _dafny.Array(_dafny.Array(None, 0), 26)
            d_1719_v39_ = nw282_
            nw283_ = _dafny.Array(_dafny.Array(None, 0), 4)
            d_1719_v39_ = nw283_
            r1 = (p1) - (p1)
            if (d_1676_v0_)[default__.safeIndex(538, (d_1676_v0_).length(0))]:
                d_1720_v40_: _dafny.Array
                def lambda83_(d_1721_p1_):
                    def lambda84_(d_1722_i4_):
                        return (d_1722_i4_) - (d_1721_p1_)

                    return lambda84_

                init49_ = lambda83_(p1)
                nw284_ = _dafny.Array(None, 3)
                for i0_49_ in range(nw284_.length(0)):
                    nw284_[i0_49_] = init49_(i0_49_)
                d_1720_v40_ = nw284_
                index321_ = default__.safeIndex(831, (d_1720_v40_).length(0))
                (d_1720_v40_)[index321_] = p1
                index322_ = default__.safeIndex(831, (d_1720_v40_).length(0))
                (d_1720_v40_)[index322_] = p1
                index323_ = default__.safeIndex(538, (d_1676_v0_).length(0))
                (d_1676_v0_)[index323_] = (d_1676_v0_)[default__.safeIndex(538, (d_1676_v0_).length(0))]
                d_1723_v41_: _dafny.Map
                d_1723_v41_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1681_v4_ for d_1724_i5_ in range(default__.abs(303))])): (d_1678_v1_ if (d_1686_v9_).f19 else d_1678_v1_)})
                d_1723_v41_ = (d_1723_v41_).set((len(d_1678_v1_)) - (len(d_1678_v1_)), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aupfiojbx")) if (d_1676_v0_)[default__.safeIndex(538, (d_1676_v0_).length(0))] else d_1678_v1_))
                d_1725_v42_: _dafny.Map
                d_1725_v42_ = _dafny.Map({(d_1676_v0_)[default__.safeIndex(538, (d_1676_v0_).length(0))]: (d_1720_v40_)[default__.safeIndex(831, (d_1720_v40_).length(0))]})
                d_1726_v43_: _dafny.Map
                d_1726_v43_ = _dafny.Map({d_1681_v4_: 720})
                d_1727_v44_: _dafny.Seq
                d_1727_v44_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_1686_v9_).f19: d_1681_v4_})), (0) - (p1)])
                d_1728_v45_: _dafny.Set
                d_1728_v45_ = _dafny.Set({d_1727_v44_, d_1727_v44_})
                d_1729_v46_: D13
                d_1729_v46_ = D13_DC41(d_1728_v45_)
                d_1730_v47_: _dafny.Map
                d_1730_v47_ = _dafny.Map({((d_1725_v42_)[(d_1686_v9_).f19] if ((d_1686_v9_).f19) in (d_1725_v42_) else (d_1720_v40_)[default__.safeIndex(831, (d_1720_v40_).length(0))]): default__.fm23(p1, (d_1686_v9_).f19, d_1726_v43_, (d_1729_v46_).cf67, globalState)})
                d_1731_v48_: C2
                nw285_ = C2()
                nw285_.ctor__(d_1730_v47_)
                d_1731_v48_ = nw285_
                d_1732_v49_: C5
                nw286_ = C5()
                nw286_.ctor__(((d_1720_v40_)[default__.safeIndex(831, (d_1720_v40_).length(0))]) * (p1), ((d_1720_v40_)[default__.safeIndex(831, (d_1720_v40_).length(0))]) > (p1), ((self).f15) | ((self).f15))
                d_1732_v49_ = nw286_
            elif True:
                (d_1686_v9_).m6(globalState)
                d_1733_v50_: _dafny.Map
                d_1733_v50_ = _dafny.Map({(d_1683_v6_).cardinality: self.f14})
                d_1734_v51_: _dafny.Seq
                d_1734_v51_ = _dafny.SeqWithoutIsStrInference([(d_1686_v9_).f19])
                d_1735_v52_: C2
                nw287_ = C2()
                nw287_.ctor__((d_1733_v50_) | ((d_1733_v50_).set(len(_dafny.SeqWithoutIsStrInference([d_1734_v51_])), (d_1686_v9_).f19)))
                d_1735_v52_ = nw287_
                d_1735_v52_ = d_1735_v52_
                d_1736_v53_: _dafny.Seq
                d_1736_v53_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_1737_v54_: C0
                nw288_ = C0()
                nw288_.ctor__(self.f14, (d_1686_v9_).f19)
                d_1737_v54_ = nw288_
                d_1738_v55_: _dafny.Set
                d_1738_v55_ = _dafny.Set({d_1737_v54_})
                d_1739_v56_: _dafny.Seq
                d_1739_v56_ = _dafny.SeqWithoutIsStrInference([d_1738_v55_, d_1738_v55_])
                index324_ = default__.safeIndex(538, (d_1676_v0_).length(0))
                (d_1676_v0_)[index324_] = (p1) != (default__.safeDivisionInt((0) - ((0) - ((d_1736_v53_)[default__.safeIndex(len(d_1678_v1_), len(d_1736_v53_))])), len(d_1739_v56_)))
                d_1740_v57_: C3
                nw289_ = C3()
                nw289_.ctor__(360, d_1678_v1_, (d_1737_v54_).f23, (self).f15)
                d_1740_v57_ = nw289_
                d_1741_v58_: _dafny.Map
                d_1741_v58_ = _dafny.Map({d_1740_v57_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfija")))})
                d_1742_v59_: D15
                d_1742_v59_ = D15_DC47(d_1741_v58_)
                r1 = (D3_DC11((0) - (p1), (d_1686_v9_).fm17(len((d_1742_v59_).cf76), self.f14, _dafny.MultiSet([d_1734_v51_]), (d_1740_v57_).f20, globalState), (d_1676_v0_)[default__.safeIndex(538, (d_1676_v0_).length(0))], (d_1740_v57_).f20)).cf17
                d_1736_v53_ = d_1736_v53_
            d_1681_v4_ = d_1681_v4_
        d_1743_v60_: _dafny.Set
        d_1743_v60_ = _dafny.Set({(d_1686_v9_).f19, True, (d_1686_v9_).f19, True, True})
        r0 = (p1 if (p1) >= (((d_1683_v6_)[len(d_1743_v60_)] if (len(d_1743_v60_)) in (d_1683_v6_) else p1)) else p1)
        r1 = ((p1) + (p1)) + (len((d_1678_v1_) + (d_1678_v1_)))
        d_1744_v61_: C0
        nw290_ = C0()
        nw290_.ctor__((d_1676_v0_)[default__.safeIndex(538, (d_1676_v0_).length(0))], not(True))
        d_1744_v61_ = nw290_
        d_1745_v62_: _dafny.Map
        d_1745_v62_ = _dafny.Map({d_1744_v61_: d_1743_v60_})
        r2 = (d_1745_v62_) == (d_1745_v62_)
        return r0, r1, r2

    def m5(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.MultiSet = _dafny.MultiSet({})
        r3: bool = False
        d_1746_v0_: _dafny.Array
        def lambda85_(d_1747_i1_):
            return self.f14

        init50_ = lambda85_
        nw291_ = _dafny.Array(None, 19)
        for i0_50_ in range(nw291_.length(0)):
            nw291_[i0_50_] = init50_(i0_50_)
        d_1746_v0_ = nw291_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1746_v0_).length(0)):
            d_1748_i0_: int = guard_loop_5_
            if (True) and (((0) <= (d_1748_i0_)) and ((d_1748_i0_) < ((d_1746_v0_).length(0)))):
                (d_1746_v0_)[(d_1748_i0_)] = (default__.safeModuloInt(-320, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vuy"))))) != (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1749_i2_ in range(default__.abs(-304))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qdea")))))
        if self.f14:
            r3 = self.f14
            r3 = self.f14
            d_1750_v1_: D0
            d_1750_v1_ = D0_DC0(self.f14)
            source17_ = d_1750_v1_
            if source17_.is_DC1:
                d_1751___mcc_h0_ = source17_.cf1
                d_1752_cf1_ = d_1751___mcc_h0_
                d_1753_v2_: int
                d_1753_v2_ = -981
                r0 = d_1753_v2_
                d_1754_v3_: _dafny.Seq
                d_1754_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikpk"))
                d_1755_v4_: _dafny.Map
                d_1755_v4_ = _dafny.Map({d_1752_cf1_: d_1754_v3_})
                d_1756_v5_: _dafny.Seq
                d_1756_v5_ = _dafny.SeqWithoutIsStrInference([d_1754_v3_, d_1754_v3_, d_1754_v3_, d_1754_v3_, d_1754_v3_])
                d_1757_v6_: str
                d_1757_v6_ = _dafny.CodePoint('f')
                d_1758_v7_: _dafny.Set
                d_1758_v7_ = _dafny.Set({d_1753_v2_})
                d_1759_v8_: D9
                d_1759_v8_ = D9_DC30(d_1757_v6_, d_1752_cf1_, len(_dafny.SeqWithoutIsStrInference([d_1757_v6_ for d_1760_i3_ in range(default__.abs(875))])), d_1758_v7_, d_1753_v2_)
                d_1761_v9_: C3
                nw292_ = C3()
                nw292_.ctor__(len((d_1755_v4_) | (d_1755_v4_)), (d_1756_v5_)[default__.safeIndex(d_1753_v2_, len(d_1756_v5_))], ((d_1759_v8_).cf51 if d_1752_cf1_ else d_1752_cf1_), (self).f15)
                d_1761_v9_ = nw292_
                d_1762_v10_: C3
                nw293_ = C3()
                nw293_.ctor__(len(_dafny.SeqWithoutIsStrInference([(d_1761_v9_).f20, (d_1761_v9_).f20, (d_1761_v9_).f20])), (d_1754_v3_) + (d_1754_v3_), True, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f14, not(d_1752_cf1_), self.f14])))
                d_1762_v10_ = nw293_
                d_1763_v11_: _dafny.Seq
                d_1763_v11_ = _dafny.SeqWithoutIsStrInference([(d_1762_v10_).f20, d_1753_v2_])
                r0 = len((d_1763_v11_) + ((d_1763_v11_) + (d_1763_v11_)))
            elif source17_.is_DC0:
                d_1764___mcc_h1_ = source17_.cf0
                d_1765_cf0_ = d_1764___mcc_h1_
                d_1766_v12_: int
                d_1766_v12_ = 83
                d_1767_v13_: C5
                nw294_ = C5()
                nw294_.ctor__((d_1766_v12_) + (d_1766_v12_), self.f14, default__.fm33(self.f14, globalState))
                d_1767_v13_ = nw294_
                d_1768_v14_: str
                d_1768_v14_ = _dafny.CodePoint('r')
                d_1769_v15_: _dafny.Map
                d_1769_v15_ = _dafny.Map({d_1768_v14_: 66})
                d_1770_v16_: _dafny.Seq
                d_1770_v16_ = _dafny.SeqWithoutIsStrInference([(d_1767_v13_).f18])
                d_1771_v17_: _dafny.Set
                d_1771_v17_ = _dafny.Set({d_1770_v16_, _dafny.SeqWithoutIsStrInference([(d_1767_v13_).f18 for d_1772_i4_ in range(default__.abs(293))]), d_1770_v16_, d_1770_v16_, d_1770_v16_})
                d_1773_v18_: C4
                nw295_ = C4()
                nw295_.ctor__(d_1765_cf0_, not (d_1765_cf0_) or (not(default__.fm23(d_1766_v12_, d_1765_cf0_, d_1769_v15_, d_1771_v17_, globalState))), (self).f15)
                d_1773_v18_ = nw295_
                d_1774_v19_: _dafny.Map
                d_1774_v19_ = _dafny.Map({d_1766_v12_: self.f14})
                d_1775_v20_: C2
                nw296_ = C2()
                nw296_.ctor__(d_1774_v19_)
                d_1775_v20_ = nw296_
                (globalState).f3 = d_1766_v12_
            elif True:
                d_1776___mcc_h2_ = source17_.cf2
                d_1777_cf2_ = d_1776___mcc_h2_
                d_1778_v21_: _dafny.Set
                d_1778_v21_ = _dafny.Set({self.f14})
                r1 = len(d_1778_v21_)
                d_1779_v22_: int
                d_1779_v22_ = 999
                d_1780_v23_: D12
                d_1780_v23_ = D12_DC38(d_1779_v22_)
                d_1780_v23_ = D12_DC38(d_1779_v22_)
                (globalState).f2 = d_1746_v0_
                d_1781_v24_: _dafny.Array
                def lambda86_(d_1782_v21_):
                    def lambda87_(d_1783_i5_):
                        return D8_DC24(d_1782_v21_)

                    return lambda87_

                init51_ = lambda86_(d_1778_v21_)
                nw297_ = _dafny.Array(None, 9)
                for i0_51_ in range(nw297_.length(0)):
                    nw297_[i0_51_] = init51_(i0_51_)
                d_1781_v24_ = nw297_
                d_1784_v25_: D8
                d_1784_v25_ = D8_DC24(d_1778_v21_)
                index325_ = default__.safeIndex(774, (d_1781_v24_).length(0))
                def iife189_(_pat_let58_0):
                    def iife190_(d_1785_dt__update__tmp_h0_):
                        def iife191_(_pat_let59_0):
                            def iife192_(d_1786_dt__update_hcf45_h0_):
                                return D8_DC24(d_1786_dt__update_hcf45_h0_)
                            return iife192_(_pat_let59_0)
                        return iife191_(_dafny.Set({self.f14}))
                    return iife190_(_pat_let58_0)
                (d_1781_v24_)[index325_] = iife189_(d_1784_v25_)
                index326_ = default__.safeIndex(774, (d_1781_v24_).length(0))
                (d_1781_v24_)[index326_] = d_1784_v25_
            d_1787_v26_: _dafny.Array
            def lambda88_(d_1788_i6_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqux"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdsulisdp")))

            init52_ = lambda88_
            nw298_ = _dafny.Array(None, 11)
            for i0_52_ in range(nw298_.length(0)):
                nw298_[i0_52_] = init52_(i0_52_)
            d_1787_v26_ = nw298_
            d_1789_v27_: _dafny.Set
            d_1789_v27_ = _dafny.Set({(self).f15})
            d_1790_v28_: D11
            d_1790_v28_ = D11_DC33(d_1789_v27_)
            pat_let_tv43_ = d_1789_v27_
            def iife193_(_pat_let60_0):
                def iife194_(d_1791_dt__update__tmp_h1_):
                    def iife195_(_pat_let61_0):
                        def iife196_(d_1792_dt__update_hcf57_h0_):
                            return D11_DC33(d_1792_dt__update_hcf57_h0_)
                        return iife196_(_pat_let61_0)
                    return iife195_(pat_let_tv43_)
                return iife194_(_pat_let60_0)
            rhs298_ = d_1787_v26_
            rhs299_ = iife193_(D11_DC33(d_1789_v27_))
            d_1787_v26_ = rhs298_
            d_1790_v28_ = rhs299_
            d_1793_v29_: C1
            nw299_ = C1()
            nw299_.ctor__()
            d_1793_v29_ = nw299_
        elif True:
            d_1794_v30_: int
            d_1794_v30_ = 140
            d_1795_v31_: _dafny.Map
            d_1795_v31_ = _dafny.Map({d_1794_v30_: 41})
            d_1795_v31_ = (d_1795_v31_).set(d_1794_v30_, (d_1794_v30_) * (d_1794_v30_))
            d_1796_v32_: _dafny.Seq
            d_1796_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shmrn"))
            if not (self.f14) or (not((d_1796_v32_) <= (d_1796_v32_))):
                d_1797_v33_: C0
                nw300_ = C0()
                nw300_.ctor__(self.f14, self.f14)
                d_1797_v33_ = nw300_
                d_1794_v30_ = 3
                d_1798_v34_: _dafny.Set
                d_1798_v34_ = _dafny.Set({default__.fm27(d_1794_v30_, self.f14, globalState)})
                d_1798_v34_ = ((d_1798_v34_) | (d_1798_v34_)) - ((d_1798_v34_ if (d_1797_v33_).f24 else d_1798_v34_))
                d_1799_v35_: str
                d_1799_v35_ = _dafny.CodePoint('n')
                d_1799_v35_ = d_1799_v35_
                d_1800_v36_: _dafny.MultiSet
                d_1800_v36_ = _dafny.MultiSet([d_1794_v30_])
                d_1801_v37_: _dafny.Seq
                d_1801_v37_ = _dafny.SeqWithoutIsStrInference([(d_1797_v33_).f23])
                d_1802_v38_: _dafny.Set
                d_1802_v38_ = _dafny.Set({(_dafny.MultiSet([self.f14, (d_1797_v33_).f24])).set(self.f14, default__.abs(len(_dafny.SeqWithoutIsStrInference([d_1799_v35_ for d_1803_i7_ in range(default__.abs(624))])))), (self).f15, _dafny.MultiSet([(d_1797_v33_).f23]), _dafny.MultiSet(d_1801_v37_)})
                rhs300_ = d_1794_v30_
                rhs301_ = ((len(d_1796_v32_)) - ((d_1800_v36_).cardinality)) + ((d_1794_v30_) - (len(d_1796_v32_)))
                rhs302_ = (d_1802_v38_).isdisjoint(d_1802_v38_)
                lhs256_ = globalState
                lhs257_ = globalState
                lhs256_.f3 = rhs300_
                r0 = rhs301_
                lhs257_.f9 = rhs302_
            elif True:
                d_1804_v39_: _dafny.Map
                d_1804_v39_ = _dafny.Map({False: d_1794_v30_})
                d_1805_v40_: str
                d_1805_v40_ = _dafny.CodePoint('w')
                d_1806_v41_: _dafny.Set
                d_1806_v41_ = _dafny.Set({default__.fm25(d_1794_v30_, d_1794_v30_, self.f14, d_1804_v39_, globalState), d_1805_v40_})
                d_1807_v42_: _dafny.Seq
                d_1807_v42_ = _dafny.SeqWithoutIsStrInference([d_1806_v41_])
                d_1808_v44_: _dafny.Map
                d_1808_v44_ = _dafny.Map({_dafny.CodePoint('e'): self.f14})
                d_1809_v45_: _dafny.Seq
                d_1809_v45_ = _dafny.SeqWithoutIsStrInference([d_1794_v30_, d_1794_v30_])
                d_1810_v46_: _dafny.Set
                d_1810_v46_ = _dafny.Set({d_1809_v45_})
                def iife197_():
                    coll73_ = _dafny.Map()
                    compr_73_: str
                    for compr_73_ in (d_1808_v44_).keys.Elements:
                        d_1811_v43_: str = compr_73_
                        if (d_1811_v43_) in (d_1808_v44_):
                            coll73_[d_1811_v43_] = len(_dafny.SeqWithoutIsStrInference([d_1794_v30_, d_1794_v30_]))
                    return _dafny.Map(coll73_)
                rhs303_ = (0) - ((0) - ((d_1794_v30_) + (658)))
                rhs304_ = not((default__.fm23(d_1794_v30_, self.f14, iife197_()
                , d_1810_v46_, globalState)) and (not (self.f14) or (self.f14)))
                rhs305_ = ((d_1795_v31_)[len(d_1796_v32_)] if (len(d_1796_v32_)) in (d_1795_v31_) else d_1794_v30_)
                rhs306_ = d_1807_v42_
                lhs258_ = globalState
                lhs259_ = globalState
                lhs258_.f3 = rhs303_
                lhs259_.f9 = rhs304_
                d_1794_v30_ = rhs305_
                d_1807_v42_ = rhs306_
                d_1812_v47_: _dafny.Set
                d_1812_v47_ = _dafny.Set({not(self.f14)})
                d_1813_v48_: _dafny.Seq
                d_1813_v48_ = _dafny.SeqWithoutIsStrInference([self.f14, (default__.fm25(d_1794_v30_, len(d_1812_v47_), self.f14, d_1804_v39_, globalState)) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpidlr"))), self.f14, self.f14])
                d_1813_v48_ = d_1813_v48_
                d_1814_v49_: D0
                d_1814_v49_ = D0_DC1(self.f14)
                d_1815_v50_: D0
                d_1815_v50_ = D0_DC2(d_1814_v49_)
                d_1816_v51_: _dafny.Map
                d_1816_v51_ = _dafny.Map({d_1815_v50_: self.f14})
                d_1817_v52_: _dafny.Map
                d_1817_v52_ = _dafny.Map({d_1794_v30_: (self).fm11(d_1816_v51_, d_1794_v30_, globalState)})
                d_1818_v53_: _dafny.Seq
                d_1818_v53_ = _dafny.SeqWithoutIsStrInference([d_1817_v52_])
                d_1819_v54_: C2
                nw301_ = C2()
                nw301_.ctor__(((d_1818_v53_)[default__.safeIndex(d_1794_v30_, len(d_1818_v53_))] if True else d_1817_v52_))
                d_1819_v54_ = nw301_
                d_1820_v55_: _dafny.MultiSet
                d_1820_v55_ = _dafny.MultiSet([d_1794_v30_, d_1794_v30_])
                r0 = default__.fm0(self.f14, d_1794_v30_, d_1794_v30_, d_1820_v55_, globalState)
                d_1821_v56_: _dafny.Array
                nw302_ = _dafny.Array(int(0), 18)
                d_1821_v56_ = nw302_
                index327_ = default__.safeIndex(407, (d_1821_v56_).length(0))
                (d_1821_v56_)[index327_] = d_1794_v30_
                index328_ = default__.safeIndex(407, (d_1821_v56_).length(0))
                (d_1821_v56_)[index328_] = (d_1794_v30_) + (d_1794_v30_)
            d_1822_v57_: _dafny.Seq
            d_1822_v57_ = _dafny.SeqWithoutIsStrInference([self.f14])
            d_1823_v58_: _dafny.Seq
            d_1823_v58_ = _dafny.SeqWithoutIsStrInference([d_1822_v57_])
            if ((_dafny.SeqWithoutIsStrInference([True, True, self.f14, self.f14])) + ((d_1823_v58_)[default__.safeIndex(654, len(d_1823_v58_))])) != ((d_1822_v57_).set(default__.safeIndex(d_1794_v30_, len(d_1822_v57_)), self.f14)):
                d_1824_v59_: _dafny.MultiSet
                d_1824_v59_ = _dafny.MultiSet([d_1794_v30_, 943, d_1794_v30_, 97])
                r1 = ((d_1824_v59_)[default__.safeModuloInt(d_1794_v30_, default__.fm0(self.f14, len(d_1796_v32_), d_1794_v30_, d_1824_v59_, globalState))] if (default__.safeModuloInt(d_1794_v30_, default__.fm0(self.f14, len(d_1796_v32_), d_1794_v30_, d_1824_v59_, globalState))) in (d_1824_v59_) else (0) - (d_1794_v30_))
                d_1825_v60_: C3
                nw303_ = C3()
                nw303_.ctor__(default__.safeModuloInt(d_1794_v30_, d_1794_v30_), d_1796_v32_, not (self.f14) or (self.f14), ((self).f15).set(self.f14, default__.abs(d_1794_v30_)))
                d_1825_v60_ = nw303_
                d_1796_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srr"))
                d_1826_v61_: _dafny.Map
                d_1826_v61_ = _dafny.Map({(d_1825_v60_).f20: (self.f14 if self.f14 else self.f14)})
                d_1826_v61_ = (d_1826_v61_).set((737 if self.f14 else d_1794_v30_), True)
                d_1827_v62_: D5
                d_1827_v62_ = D5_DC18(self.f14)
                (globalState).f9 = (d_1827_v62_).cf33
            elif True:
                (globalState).f9 = ((0) - ((d_1794_v30_) - (-320))) < (d_1794_v30_)
                d_1828_v63_: D1
                d_1828_v63_ = D1_DC5(self.f14, self.f14)
                d_1829_v64_: C5
                nw304_ = C5()
                nw304_.ctor__(len(d_1822_v57_), (d_1828_v63_).cf6, (self).f15)
                d_1829_v64_ = nw304_
                d_1830_v65_: D15
                d_1830_v65_ = D15_DC48(d_1794_v30_, d_1829_v64_, 491)
                (globalState).f3 = (d_1830_v65_).cf79
                d_1831_v66_: _dafny.Map
                d_1831_v66_ = _dafny.Map({d_1746_v0_: self.f14})
                index329_ = default__.safeIndex(935, (d_1746_v0_).length(0))
                (d_1746_v0_)[index329_] = not(((d_1831_v66_)[d_1746_v0_] if (d_1746_v0_) in (d_1831_v66_) else self.f14))
                index330_ = default__.safeIndex(935, (d_1746_v0_).length(0))
                (d_1746_v0_)[index330_] = self.f14
                d_1832_v67_: D0
                d_1832_v67_ = D0_DC0((d_1746_v0_)[default__.safeIndex(935, (d_1746_v0_).length(0))])
                (globalState).f3 = (d_1794_v30_ if (d_1832_v67_).cf0 else (0) - ((d_1829_v64_).f18))
                d_1833_v68_: D7
                d_1833_v68_ = D7_DC23(self.f14, d_1794_v30_, self.f14)
                d_1833_v68_ = d_1833_v68_
            (globalState).f2 = d_1746_v0_
            d_1834_v69_: str
            d_1834_v69_ = _dafny.CodePoint('o')
            d_1835_v70_: D2
            d_1835_v70_ = D2_DC9(self.f14, d_1794_v30_, _dafny.CodePoint('v'))
            (globalState).f9 = ((D2_DC9(self.f14, d_1794_v30_, d_1834_v69_) if self.f14 else d_1835_v70_)).cf13
        d_1836_v71_: _dafny.Array
        nw305_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
        d_1836_v71_ = nw305_
        d_1837_v72_: _dafny.Seq
        d_1837_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hau"))
        index331_ = default__.safeIndex(202, (d_1836_v71_).length(0))
        (d_1836_v71_)[index331_] = d_1837_v72_
        d_1838_v73_: int
        d_1838_v73_ = 437
        index332_ = default__.safeIndex(202, (d_1836_v71_).length(0))
        (d_1836_v71_)[index332_] = ((d_1837_v72_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1839_i8_ in range(default__.abs(51))]))) + ((default__.fm27(d_1838_v73_, self.f14, globalState)) + (d_1837_v72_))
        d_1840_v74_: _dafny.Set
        d_1840_v74_ = _dafny.Set({d_1838_v73_})
        hi7_ = d_1838_v73_
        for d_1841_i9_ in range(len(d_1840_v74_), hi7_):
            d_1842_v75_: D5
            d_1842_v75_ = D5_DC18(not(True))
            d_1843_v76_: _dafny.Map
            d_1843_v76_ = _dafny.Map({d_1841_i9_: d_1842_v75_})
            d_1843_v76_ = (d_1843_v76_).set((len(d_1840_v74_)) - (d_1838_v73_), d_1842_v75_)
            index333_ = default__.safeIndex(686, (d_1746_v0_).length(0))
            (d_1746_v0_)[index333_] = (self.f14 if self.f14 else not(self.f14))
            index334_ = default__.safeIndex(129, (d_1746_v0_).length(0))
            (d_1746_v0_)[index334_] = self.f14
            d_1844_v77_: str
            d_1844_v77_ = _dafny.CodePoint('h')
            index335_ = default__.safeIndex(686, (d_1746_v0_).length(0))
            index336_ = default__.safeIndex(129, (d_1746_v0_).length(0))
            index337_ = default__.safeIndex(202, (d_1836_v71_).length(0))
            rhs307_ = True
            rhs308_ = self.f14
            rhs309_ = ((d_1837_v72_).set(default__.safeIndex(d_1838_v73_, len(d_1837_v72_)), d_1844_v77_)) + ((d_1836_v71_)[default__.safeIndex(202, (d_1836_v71_).length(0))])
            lhs260_ = d_1746_v0_
            lhs261_ = default__.safeIndex(686, (d_1746_v0_).length(0))
            lhs262_ = d_1746_v0_
            lhs263_ = default__.safeIndex(129, (d_1746_v0_).length(0))
            lhs264_ = d_1836_v71_
            lhs265_ = default__.safeIndex(202, (d_1836_v71_).length(0))
            lhs260_[lhs261_] = rhs307_
            lhs262_[lhs263_] = rhs308_
            lhs264_[lhs265_] = rhs309_
            d_1845_v78_: _dafny.Seq
            d_1845_v78_ = _dafny.SeqWithoutIsStrInference([True, self.f14])
            (globalState).f3 = len(((d_1845_v78_) + (d_1845_v78_)) + (d_1845_v78_))
            d_1846_v79_: _dafny.Map
            d_1846_v79_ = _dafny.Map({d_1841_i9_: d_1841_i9_})
            d_1847_v80_: D9
            d_1847_v80_ = D9_DC29(d_1846_v79_)
            d_1848_v81_: D9
            d_1848_v81_ = D9_DC31(d_1847_v80_)
            d_1849_v82_: D9
            d_1849_v82_ = D9_DC31(d_1848_v81_)
            d_1850_v83_: _dafny.Seq
            d_1850_v83_ = _dafny.SeqWithoutIsStrInference([d_1849_v82_])
            d_1851_v84_: _dafny.Set
            d_1851_v84_ = _dafny.Set({d_1850_v83_})
            d_1851_v84_ = default__.fm49(d_1838_v73_, globalState)
        d_1852_v85_: _dafny.Array
        def lambda89_(d_1853_v73_):
            def lambda90_(d_1854_i10_):
                return default__.safeModuloInt(d_1854_i10_, d_1853_v73_)

            return lambda90_

        init53_ = lambda89_(d_1838_v73_)
        nw306_ = _dafny.Array(None, 22)
        for i0_53_ in range(nw306_.length(0)):
            nw306_[i0_53_] = init53_(i0_53_)
        d_1852_v85_ = nw306_
        index338_ = default__.safeIndex(239, (d_1852_v85_).length(0))
        (d_1852_v85_)[index338_] = d_1838_v73_
        index339_ = default__.safeIndex(239, (d_1852_v85_).length(0))
        (d_1852_v85_)[index339_] = (0) - (default__.safeModuloInt((d_1838_v73_ if self.f14 else d_1838_v73_), (284 if self.f14 else 269)))
        d_1855_v86_: _dafny.Seq
        d_1855_v86_ = _dafny.SeqWithoutIsStrInference([(d_1852_v85_)[default__.safeIndex(239, (d_1852_v85_).length(0))]])
        d_1856_v87_: _dafny.Map
        d_1856_v87_ = _dafny.Map({len((d_1855_v86_) + (_dafny.SeqWithoutIsStrInference([(d_1852_v85_)[default__.safeIndex(239, (d_1852_v85_).length(0))]]))): self.f14})
        d_1856_v87_ = (d_1856_v87_).set(d_1838_v73_, ((d_1852_v85_)[default__.safeIndex(239, (d_1852_v85_).length(0))]) <= (510))
        r0 = (d_1852_v85_)[default__.safeIndex(239, (d_1852_v85_).length(0))]
        r1 = (d_1838_v73_) * (d_1838_v73_)
        d_1857_v88_: _dafny.MultiSet
        d_1857_v88_ = _dafny.MultiSet([d_1838_v73_, len(d_1837_v72_), d_1838_v73_, default__.safeModuloInt(d_1838_v73_, (d_1852_v85_)[default__.safeIndex(239, (d_1852_v85_).length(0))])])
        r2 = d_1857_v88_
        d_1858_v90_: _dafny.Seq
        d_1858_v90_ = _dafny.SeqWithoutIsStrInference([d_1856_v87_])
        d_1859_v91_: D7
        def iife198_():
            coll74_ = _dafny.Map()
            compr_74_: _dafny.Map
            for compr_74_ in (d_1858_v90_).Elements:
                d_1860_v89_: _dafny.Map = compr_74_
                if (d_1860_v89_) in (d_1858_v90_):
                    coll74_[d_1860_v89_] = d_1838_v73_
            return _dafny.Map(coll74_)
        d_1859_v91_ = D7_DC21(iife198_()
)
        def lambda91_(source18_):
            if source18_.is_DC22:
                d_1861___mcc_h3_ = source18_.cf37
                d_1862___mcc_h4_ = source18_.cf38
                d_1863___mcc_h5_ = source18_.cf39
                d_1864___mcc_h6_ = source18_.cf40
                d_1865___mcc_h7_ = source18_.cf41
                d_1866_cf41_ = d_1865___mcc_h7_
                d_1867_cf40_ = d_1864___mcc_h6_
                d_1868_cf39_ = d_1863___mcc_h5_
                d_1869_cf38_ = d_1862___mcc_h4_
                d_1870_cf37_ = d_1861___mcc_h3_
                return not(self.f14)
            elif source18_.is_DC23:
                d_1871___mcc_h8_ = source18_.cf42
                d_1872___mcc_h9_ = source18_.cf43
                d_1873___mcc_h10_ = source18_.cf44
                d_1874_cf44_ = d_1873___mcc_h10_
                d_1875_cf43_ = d_1872___mcc_h9_
                d_1876_cf42_ = d_1871___mcc_h8_
                return d_1874_cf44_
            elif True:
                d_1877___mcc_h11_ = source18_.cf36
                d_1878_cf36_ = d_1877___mcc_h11_
                return self.f14

        r3 = lambda91_(d_1859_v91_)
        return r0, r1, r2, r3


class C7:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self):
        pass
        pass

    def fm9(self, p0, globalState):
        return len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kko"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddx"))))

    def fm10(self, p0, p1, p2, p3, globalState):
        return False

    def m3(self, p0, globalState):
        hi8_ = p0
        for d_1879_i0_ in range(p0, hi8_):
            d_1880_v0_: D1
            d_1880_v0_ = D1_DC3(-181)
            d_1881_v1_: _dafny.Map
            d_1881_v1_ = _dafny.Map({d_1880_v0_: d_1879_i0_})
            d_1882_v2_: _dafny.Seq
            d_1882_v2_ = _dafny.SeqWithoutIsStrInference([len(d_1881_v1_), p0])
            d_1883_v3_: str
            d_1883_v3_ = _dafny.CodePoint('j')
            d_1884_v4_: _dafny.Seq
            d_1884_v4_ = _dafny.SeqWithoutIsStrInference([d_1883_v3_])
            d_1885_v5_: _dafny.Array
            nw307_ = _dafny.Array(None, 15)
            nw307_[int(0)] = d_1882_v2_
            nw307_[int(1)] = (d_1882_v2_) + (d_1882_v2_)
            nw307_[int(2)] = d_1882_v2_
            nw307_[int(3)] = d_1882_v2_
            nw307_[int(4)] = d_1882_v2_
            nw307_[int(5)] = (d_1882_v2_).set(default__.safeIndex(d_1879_i0_, len(d_1882_v2_)), (d_1880_v0_).cf3)
            nw307_[int(6)] = _dafny.SeqWithoutIsStrInference([len(d_1884_v4_), p0])
            nw307_[int(7)] = d_1882_v2_
            nw307_[int(8)] = d_1882_v2_
            nw307_[int(9)] = (d_1882_v2_ if True else d_1882_v2_)
            nw307_[int(10)] = (d_1882_v2_).set(default__.safeIndex(d_1879_i0_, len(d_1882_v2_)), d_1879_i0_)
            nw307_[int(11)] = (d_1882_v2_) + (_dafny.SeqWithoutIsStrInference([p0, p0, len(d_1884_v4_)]))
            nw307_[int(12)] = d_1882_v2_
            nw307_[int(13)] = d_1882_v2_
            nw307_[int(14)] = d_1882_v2_
            d_1885_v5_ = nw307_
            index340_ = default__.safeIndex(789, (d_1885_v5_).length(0))
            (d_1885_v5_)[index340_] = (d_1882_v2_).set(default__.safeIndex(d_1879_i0_, len(d_1882_v2_)), p0)
            d_1886_v6_: _dafny.Array
            nw308_ = _dafny.Array(False, 26)
            d_1886_v6_ = nw308_
            d_1887_v7_: bool
            d_1887_v7_ = True
            index341_ = default__.safeIndex(752, (d_1886_v6_).length(0))
            (d_1886_v6_)[index341_] = d_1887_v7_
            index342_ = default__.safeIndex(789, (d_1885_v5_).length(0))
            index343_ = default__.safeIndex(752, (d_1886_v6_).length(0))
            rhs310_ = d_1887_v7_
            rhs311_ = ((d_1882_v2_) + (d_1882_v2_)) + ((d_1882_v2_) + (d_1882_v2_))
            rhs312_ = not(d_1887_v7_)
            lhs266_ = globalState
            lhs267_ = d_1885_v5_
            lhs268_ = default__.safeIndex(789, (d_1885_v5_).length(0))
            lhs269_ = d_1886_v6_
            lhs270_ = default__.safeIndex(752, (d_1886_v6_).length(0))
            lhs266_.f9 = rhs310_
            lhs267_[lhs268_] = rhs311_
            lhs269_[lhs270_] = rhs312_
            d_1888_v8_: _dafny.Map
            d_1888_v8_ = _dafny.Map({p0: d_1883_v3_})
            d_1889_v9_: _dafny.MultiSet
            d_1889_v9_ = _dafny.MultiSet([d_1888_v8_])
            d_1890_v10_: _dafny.Array
            nw309_ = _dafny.Array(_dafny.Map({}), 2)
            d_1890_v10_ = nw309_
            d_1891_v11_: _dafny.Map
            d_1891_v11_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(d_1886_v6_)[default__.safeIndex(752, (d_1886_v6_).length(0))]])) for d_1892_i1_ in range(default__.abs(414))]): d_1883_v3_})
            d_1893_v12_: _dafny.Map
            d_1893_v12_ = _dafny.Map({d_1887_v7_: d_1886_v6_})
            rhs313_ = (d_1889_v9_) - (d_1889_v9_)
            rhs314_ = not(d_1887_v7_)
            rhs315_ = ((d_1893_v12_)[(d_1886_v6_)[default__.safeIndex(752, (d_1886_v6_).length(0))]] if ((d_1886_v6_)[default__.safeIndex(752, (d_1886_v6_).length(0))]) in (d_1893_v12_) else d_1886_v6_)
            rhs316_ = d_1890_v10_
            rhs317_ = _dafny.Map({(d_1885_v5_)[default__.safeIndex(789, (d_1885_v5_).length(0))]: d_1883_v3_})
            lhs271_ = globalState
            d_1889_v9_ = rhs313_
            d_1887_v7_ = rhs314_
            lhs271_.f2 = rhs315_
            d_1890_v10_ = rhs316_
            d_1891_v11_ = rhs317_
            d_1884_v4_ = (d_1884_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "reyaln")))
            (globalState).f9 = (d_1879_i0_) != (d_1879_i0_)
        hi9_ = p0
        for d_1894_i2_ in range(-126, hi9_):
            d_1895_v13_: _dafny.Array
            nw310_ = _dafny.Array(D1.default()(), 2)
            d_1895_v13_ = nw310_
            d_1896_v14_: D1
            d_1896_v14_ = D1_DC4(573, p0)
            index344_ = default__.safeIndex(84, (d_1895_v13_).length(0))
            def iife199_(_pat_let62_0):
                def iife200_(d_1897_dt__update__tmp_h0_):
                    def iife201_(_pat_let63_0):
                        def iife202_(d_1898_dt__update_hcf4_h0_):
                            return D1_DC4(d_1898_dt__update_hcf4_h0_, (d_1897_dt__update__tmp_h0_).cf5)
                        return iife202_(_pat_let63_0)
                    return iife201_(d_1894_i2_)
                return iife200_(_pat_let62_0)
            (d_1895_v13_)[index344_] = iife199_(d_1896_v14_)
            index345_ = default__.safeIndex(84, (d_1895_v13_).length(0))
            (d_1895_v13_)[index345_] = d_1896_v14_
            d_1899_v15_: bool
            d_1899_v15_ = True
            d_1900_v16_: C0
            nw311_ = C0()
            nw311_.ctor__(d_1899_v15_, d_1899_v15_)
            d_1900_v16_ = nw311_
            d_1901_v17_: _dafny.Map
            d_1901_v17_ = _dafny.Map({-189: p0})
            (globalState).f9 = (d_1901_v17_) != (default__.fm14(globalState))
            d_1902_v18_: str
            d_1902_v18_ = _dafny.CodePoint('o')
            d_1903_v19_: _dafny.Map
            d_1903_v19_ = _dafny.Map({d_1902_v18_: (0) - (d_1894_i2_)})
            d_1904_v20_: _dafny.Seq
            d_1904_v20_ = _dafny.SeqWithoutIsStrInference([d_1894_i2_, d_1894_i2_])
            d_1905_v21_: _dafny.Set
            d_1905_v21_ = _dafny.Set({d_1904_v20_, d_1904_v20_})
            d_1906_v22_: _dafny.Array
            def lambda92_(d_1907_v16_):
                def lambda93_(d_1908_i3_):
                    return (d_1907_v16_).f24

                return lambda93_

            init54_ = lambda92_(d_1900_v16_)
            nw312_ = _dafny.Array(None, 6)
            for i0_54_ in range(nw312_.length(0)):
                nw312_[i0_54_] = init54_(i0_54_)
            d_1906_v22_ = nw312_
            (globalState).f2 = (d_1906_v22_ if not((d_1899_v15_ if default__.fm23(p0, d_1899_v15_, d_1903_v19_, d_1905_v21_, globalState) else False)) else d_1906_v22_)
        d_1909_v23_: _dafny.MultiSet
        d_1909_v23_ = _dafny.MultiSet([(0) - (p0), 704])
        (globalState).f9 = not(not(((d_1909_v23_).intersection(d_1909_v23_)).isdisjoint(d_1909_v23_)))
        d_1910_i4_: int
        d_1910_i4_ = 0
        with _dafny.label("4"):
            while (p0) != (p0):
                with _dafny.c_label("4"):
                    if (d_1910_i4_) >= (100):
                        raise _dafny.Break("4")
                    d_1910_i4_ = (d_1910_i4_) + (1)
                    d_1911_v24_: _dafny.Array
                    nw313_ = _dafny.Array(False, 14)
                    d_1911_v24_ = nw313_
                    d_1912_v25_: _dafny.Seq
                    d_1912_v25_ = _dafny.SeqWithoutIsStrInference([d_1911_v24_, d_1911_v24_, d_1911_v24_, d_1911_v24_, d_1911_v24_])
                    d_1913_v26_: bool
                    d_1913_v26_ = False
                    d_1912_v25_ = (d_1912_v25_ if d_1913_v26_ else d_1912_v25_)
                    d_1913_v26_ = d_1913_v26_
                    d_1914_v27_: _dafny.Seq
                    d_1914_v27_ = _dafny.SeqWithoutIsStrInference([d_1913_v26_])
                    d_1915_v28_: _dafny.Seq
                    d_1915_v28_ = _dafny.SeqWithoutIsStrInference([p0, len((d_1914_v27_).set(default__.safeIndex(p0, len(d_1914_v27_)), not(d_1913_v26_))), -573, len(_dafny.Map({d_1913_v26_: d_1913_v26_})), (0) - (p0)])
                    d_1916_v29_: _dafny.Seq
                    d_1916_v29_ = (_dafny.SeqWithoutIsStrInference([506 for d_1917_i5_ in range(default__.abs(915))])) + (d_1915_v28_)
                    source19_ = d_1916_v29_
                    d_1918___mcc_h0_ = source19_
                    d_1919_cf56_ = d_1918___mcc_h0_
                    d_1920_v30_: _dafny.Array
                    nw314_ = _dafny.Array(_dafny.CodePoint('D'), 8)
                    d_1920_v30_ = nw314_
                    d_1921_v31_: _dafny.Seq
                    d_1921_v31_ = _dafny.SeqWithoutIsStrInference([d_1920_v30_])
                    d_1922_v32_: D16
                    d_1922_v32_ = D16_DC51(d_1920_v30_)
                    d_1921_v31_ = _dafny.SeqWithoutIsStrInference([d_1920_v30_, (d_1922_v32_).cf83])
                    d_1923_v33_: _dafny.Map
                    d_1923_v33_ = _dafny.Map({d_1916_v29_: p0})
                    d_1924_v34_: _dafny.MultiSet
                    d_1924_v34_ = _dafny.MultiSet([d_1923_v33_])
                    (globalState).f9 = (d_1923_v33_) in (d_1924_v34_)
                    d_1925_v35_: _dafny.Set
                    d_1925_v35_ = _dafny.Set({d_1913_v26_})
                    (globalState).f6 = (d_1925_v35_).intersection(default__.fm41(p0, p0, globalState))
                    d_1926_v36_: _dafny.Map
                    d_1926_v36_ = _dafny.Map({p0: not(d_1913_v26_)})
                    d_1927_v37_: _dafny.Seq
                    d_1927_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
                    d_1928_v38_: _dafny.MultiSet
                    d_1928_v38_ = _dafny.MultiSet([d_1913_v26_])
                    d_1929_v39_: C3
                    nw315_ = C3()
                    nw315_.ctor__(p0, d_1927_v37_, d_1913_v26_, d_1928_v38_)
                    d_1929_v39_ = nw315_
                    d_1930_v40_: _dafny.Seq
                    d_1930_v40_ = _dafny.SeqWithoutIsStrInference([d_1929_v39_, d_1929_v39_, d_1929_v39_, d_1929_v39_, d_1929_v39_])
                    d_1931_v41_: _dafny.Map
                    d_1931_v41_ = _dafny.Map({d_1930_v40_: d_1928_v38_})
                    d_1932_v42_: _dafny.Map
                    d_1932_v42_ = _dafny.Map({d_1913_v26_: d_1931_v41_})
                    d_1933_v43_: _dafny.Map
                    d_1933_v43_ = _dafny.Map({len(d_1914_v27_): d_1931_v41_})
                    d_1926_v36_ = (d_1926_v36_).set((0) - (len((((d_1932_v42_)[d_1913_v26_] if (d_1913_v26_) in (d_1932_v42_) else ((d_1933_v43_)[p0] if (p0) in (d_1933_v43_) else _dafny.Map({d_1930_v40_: d_1928_v38_})))).set(d_1930_v40_, d_1928_v38_))), d_1913_v26_)
                    d_1934_v44_: _dafny.Set
                    d_1934_v44_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([866, p0])})
                    d_1935_v45_: D13
                    d_1935_v45_ = D13_DC41(d_1934_v44_)
                    source20_ = d_1935_v45_
                    if source20_.is_DC42:
                        d_1936___mcc_h1_ = source20_.cf68
                        d_1937___mcc_h2_ = source20_.cf69
                        d_1938_cf69_ = d_1937___mcc_h2_
                        d_1939_cf68_ = d_1936___mcc_h1_
                        d_1940_v46_: _dafny.Set
                        d_1940_v46_ = _dafny.Set({d_1938_cf69_, d_1938_cf69_, d_1938_cf69_})
                        (globalState).f6 = (d_1940_v46_ if d_1913_v26_ else d_1940_v46_)
                        d_1941_v47_: _dafny.Array
                        nw316_ = _dafny.Array(int(0), 4)
                        d_1941_v47_ = nw316_
                        index346_ = default__.safeIndex(241, (d_1941_v47_).length(0))
                        (d_1941_v47_)[index346_] = (self).fm9(p0, globalState)
                        index347_ = default__.safeIndex(241, (d_1941_v47_).length(0))
                        (d_1941_v47_)[index347_] = len((d_1912_v25_) + ((d_1912_v25_) + (d_1912_v25_)))
                        d_1942_v48_: _dafny.Map
                        d_1942_v48_ = _dafny.Map({d_1938_cf69_: (d_1914_v27_) + ((_dafny.SeqWithoutIsStrInference([d_1913_v26_])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([d_1913_v26_]))), d_1938_cf69_))})
                        d_1943_v49_: _dafny.Map
                        d_1943_v49_ = _dafny.Map({_dafny.CodePoint('h'): (d_1941_v47_)[default__.safeIndex(241, (d_1941_v47_).length(0))]})
                        d_1944_v50_: _dafny.Map
                        d_1944_v50_ = _dafny.Map({d_1938_cf69_: 202})
                        d_1945_v51_: _dafny.Seq
                        d_1945_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrpyg"))
                        d_1942_v48_ = (d_1942_v48_).set(default__.fm23((d_1941_v47_)[default__.safeIndex(241, (d_1941_v47_).length(0))], False, d_1943_v49_, _dafny.Set({d_1915_v28_, d_1915_v28_, d_1915_v28_}), globalState), ((d_1942_v48_)[d_1913_v26_] if (d_1913_v26_) in (d_1942_v48_) else _dafny.SeqWithoutIsStrInference([d_1913_v26_, d_1938_cf69_, not((self).fm10(d_1944_v50_, d_1909_v23_, False, d_1945_v51_, globalState))])))
                        d_1946_v52_: _dafny.Seq
                        d_1946_v52_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sct"))])
                        d_1947_v53_: D4
                        d_1947_v53_ = D4_DC13(_dafny.SeqWithoutIsStrInference([d_1939_cf68_ for d_1948_i8_ in range(default__.abs(326))]), len(_dafny.SeqWithoutIsStrInference([(D13_DC42(d_1939_cf68_, (D9_DC30(d_1939_cf68_, d_1938_cf69_, p0, _dafny.Set({(d_1941_v47_)[default__.safeIndex(241, (d_1941_v47_).length(0))]}), (d_1941_v47_)[default__.safeIndex(241, (d_1941_v47_).length(0))])).cf51)).cf68 for d_1949_i9_ in range(default__.abs(-776))])), d_1913_v26_, d_1911_v24_)
                        d_1950_v54_: _dafny.Array
                        nw317_ = _dafny.Array(None, 23)
                        nw317_[int(0)] = ((d_1946_v52_)[default__.safeIndex((0) - (p0), len(d_1946_v52_))]) + (d_1945_v51_)
                        nw317_[int(1)] = d_1945_v51_
                        nw317_[int(2)] = d_1945_v51_
                        nw317_[int(3)] = (d_1945_v51_) + (d_1945_v51_)
                        nw317_[int(4)] = d_1945_v51_
                        nw317_[int(5)] = (d_1945_v51_) + (d_1945_v51_)
                        nw317_[int(6)] = d_1945_v51_
                        nw317_[int(7)] = (d_1945_v51_) + (d_1945_v51_)
                        nw317_[int(8)] = (d_1945_v51_).set(default__.safeIndex((d_1941_v47_)[default__.safeIndex(241, (d_1941_v47_).length(0))], len(d_1945_v51_)), d_1939_cf68_)
                        nw317_[int(9)] = _dafny.SeqWithoutIsStrInference([d_1939_cf68_ for d_1951_i6_ in range(default__.abs(218))])
                        nw317_[int(10)] = d_1945_v51_
                        nw317_[int(11)] = _dafny.SeqWithoutIsStrInference([d_1939_cf68_ for d_1952_i7_ in range(default__.abs(804))])
                        nw317_[int(12)] = d_1945_v51_
                        nw317_[int(13)] = d_1945_v51_
                        nw317_[int(14)] = (d_1947_v53_).cf22
                        nw317_[int(15)] = _dafny.SeqWithoutIsStrInference([d_1939_cf68_ for d_1953_i10_ in range(default__.abs(325))])
                        nw317_[int(16)] = d_1945_v51_
                        nw317_[int(17)] = d_1945_v51_
                        nw317_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abhvu"))
                        nw317_[int(19)] = (d_1945_v51_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")))
                        nw317_[int(20)] = d_1945_v51_
                        nw317_[int(21)] = (d_1945_v51_ if d_1938_cf69_ else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1954_i11_ in range(default__.abs(390))]))
                        nw317_[int(22)] = d_1945_v51_
                        d_1950_v54_ = nw317_
                        index348_ = default__.safeIndex(243, (d_1950_v54_).length(0))
                        (d_1950_v54_)[index348_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
                        index349_ = default__.safeIndex(243, (d_1950_v54_).length(0))
                        (d_1950_v54_)[index349_] = (d_1945_v51_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))
                    elif source20_.is_DC41:
                        d_1955___mcc_h3_ = source20_.cf67
                        d_1956_cf67_ = d_1955___mcc_h3_
                        d_1957_v55_: _dafny.Seq
                        d_1957_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gv"))
                        d_1958_v56_: _dafny.Map
                        d_1958_v56_ = _dafny.Map({p0: len(d_1957_v55_)})
                        def iife203_():
                            coll75_ = _dafny.Map()
                            compr_75_: int
                            for compr_75_ in (d_1909_v23_).Elements:
                                d_1959_v57_: int = compr_75_
                                if (d_1959_v57_) in (d_1909_v23_):
                                    coll75_[default__.safeModuloInt(d_1959_v57_, len(d_1957_v55_))] = p0
                            return _dafny.Map(coll75_)
                        (globalState).f3 = default__.safeDivisionInt((p0 if d_1913_v26_ else p0), len((d_1958_v56_) | ((iife203_()
                        ).set(p0, p0))))
                        d_1960_v58_: str
                        d_1960_v58_ = _dafny.CodePoint('f')
                        d_1961_v59_: _dafny.Array
                        nw318_ = _dafny.Array(int(0), 13)
                        d_1961_v59_ = nw318_
                        d_1962_v60_: _dafny.Map
                        d_1962_v60_ = _dafny.Map({d_1961_v59_: d_1913_v26_})
                        d_1963_v61_: _dafny.Set
                        d_1963_v61_ = _dafny.Set({default__.fm0(d_1913_v26_, 418, len(d_1962_v60_), d_1909_v23_, globalState), len(d_1957_v55_)})
                        d_1964_v62_: _dafny.Set
                        d_1964_v62_ = _dafny.Set({D9_DC30(d_1960_v58_, d_1913_v26_, p0, d_1963_v61_, p0)})
                        d_1965_v63_: _dafny.Map
                        d_1965_v63_ = _dafny.Map({d_1960_v58_: p0})
                        d_1966_v64_: D0
                        d_1966_v64_ = D0_DC1(d_1913_v26_)
                        rhs318_ = (D9_DC30(d_1960_v58_, d_1913_v26_, p0, d_1963_v61_, p0)) in (d_1964_v62_)
                        rhs319_ = (d_1911_v24_ if default__.fm23(p0, d_1913_v26_, d_1965_v63_, default__.fm50(p0, (d_1966_v64_).cf1, p0, p0, globalState), globalState) else d_1911_v24_)
                        lhs272_ = globalState
                        d_1913_v26_ = rhs318_
                        lhs272_.f2 = rhs319_
                        d_1967_v65_: _dafny.Array
                        nw319_ = _dafny.Array(None, 13)
                        nw319_[int(0)] = d_1911_v24_
                        nw319_[int(1)] = d_1911_v24_
                        nw319_[int(2)] = d_1911_v24_
                        nw319_[int(3)] = d_1911_v24_
                        nw319_[int(4)] = d_1911_v24_
                        nw319_[int(5)] = d_1911_v24_
                        nw319_[int(6)] = d_1911_v24_
                        nw319_[int(7)] = d_1911_v24_
                        nw319_[int(8)] = d_1911_v24_
                        nw319_[int(9)] = d_1911_v24_
                        nw319_[int(10)] = d_1911_v24_
                        nw319_[int(11)] = d_1911_v24_
                        nw319_[int(12)] = d_1911_v24_
                        d_1967_v65_ = nw319_
                        d_1968_v66_: _dafny.Array
                        d_1968_v66_ = d_1967_v65_
                        d_1969_v67_: _dafny.Array
                        nw320_ = _dafny.Array(None, 3)
                        nw320_[int(0)] = (d_1967_v65_ if d_1913_v26_ else d_1967_v65_)
                        nw320_[int(1)] = (d_1968_v66_)
                        nw320_[int(2)] = d_1967_v65_
                        d_1969_v67_ = nw320_
                        index350_ = default__.safeIndex(597, (d_1969_v67_).length(0))
                        (d_1969_v67_)[index350_] = (d_1967_v65_ if d_1913_v26_ else d_1967_v65_)
                        index351_ = default__.safeIndex(597, (d_1969_v67_).length(0))
                        (d_1969_v67_)[index351_] = d_1967_v65_
                        index352_ = default__.safeIndex(857, (d_1961_v59_).length(0))
                        (d_1961_v59_)[index352_] = p0
                        d_1970_v68_: _dafny.Map
                        d_1970_v68_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "itn"))})
                        index353_ = default__.safeIndex(857, (d_1961_v59_).length(0))
                        (d_1961_v59_)[index353_] = default__.safeModuloInt(-772, len(d_1970_v68_))
                    elif True:
                        d_1971___mcc_h4_ = source20_.cf70
                        d_1972_cf70_ = d_1971___mcc_h4_
                        d_1973_v69_: C0
                        nw321_ = C0()
                        nw321_.ctor__(d_1913_v26_, not(d_1913_v26_))
                        d_1973_v69_ = nw321_
                        (globalState).f9 = (d_1973_v69_).f24
                        (globalState).f9 = ((53) - (175)) > (p0)
                        d_1974_v70_: _dafny.Array
                        nw322_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
                        d_1974_v70_ = nw322_
                        d_1975_v71_: _dafny.Array
                        nw323_ = _dafny.Array(_dafny.Set({}), 10)
                        d_1975_v71_ = nw323_
                        d_1976_v72_: _dafny.Map
                        d_1976_v72_ = _dafny.Map({d_1913_v26_: (d_1973_v69_).f23})
                        rhs320_ = d_1975_v71_
                        rhs321_ = ((d_1974_v70_ if False else d_1974_v70_) if ((d_1976_v72_)[d_1913_v26_] if (d_1913_v26_) in (d_1976_v72_) else d_1913_v26_) else d_1974_v70_)
                        lhs273_ = globalState
                        lhs273_.f11 = rhs320_
                        d_1974_v70_ = rhs321_
                    pass
            pass
        d_1977_v73_: _dafny.Seq
        d_1977_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
        d_1978_v74_: D0
        d_1978_v74_ = D0_DC2(default__.fm20(p0, len(d_1977_v73_), p0, globalState))
        d_1979_v75_: D0
        d_1979_v75_ = D0_DC2(d_1978_v74_)
        pat_let_tv44_ = p0
        def lambda94_(source21_):
            if source21_.is_DC1:
                d_1980___mcc_h5_ = source21_.cf1
                d_1981_cf1_ = d_1980___mcc_h5_
                return not(not((952) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1982_i12_ in range(default__.abs(234))])), pat_let_tv44_]))))
            elif source21_.is_DC0:
                d_1983___mcc_h6_ = source21_.cf0
                d_1984_cf0_ = d_1983___mcc_h6_
                return d_1984_cf0_
            elif True:
                d_1985___mcc_h7_ = source21_.cf2
                d_1986_cf2_ = d_1985___mcc_h7_
                return (_dafny.Map({True: False})) != (_dafny.Map({False: not(False)}))

        if lambda94_(d_1979_v75_):
            d_1987_v76_: bool
            d_1987_v76_ = False
            d_1988_v77_: _dafny.Map
            d_1988_v77_ = _dafny.Map({d_1977_v73_: d_1987_v76_})
            d_1988_v77_ = d_1988_v77_
            (globalState).f9 = (d_1987_v76_) or (d_1987_v76_)
            d_1989_v78_: _dafny.MultiSet
            d_1989_v78_ = _dafny.MultiSet([d_1987_v76_, d_1987_v76_])
            d_1990_v79_: C5
            nw324_ = C5()
            nw324_.ctor__((0) - (p0), (len(d_1977_v73_)) <= (p0), d_1989_v78_)
            d_1990_v79_ = nw324_
            d_1991_v80_: _dafny.Map
            d_1991_v80_ = _dafny.Map({True: d_1989_v78_})
            d_1992_v81_: _dafny.Map
            d_1992_v81_ = _dafny.Map({((d_1991_v80_)[not(d_1987_v76_)] if (not(d_1987_v76_)) in (d_1991_v80_) else default__.fm33(d_1987_v76_, globalState)): (d_1987_v76_ if d_1987_v76_ else d_1987_v76_)})
            d_1992_v81_ = (d_1992_v81_).set(d_1989_v78_, False)
            d_1987_v76_ = (default__.safeModuloInt((d_1990_v79_).f18, p0)) < (len(d_1977_v73_))
        elif True:
            (globalState).f3 = p0
            d_1993_v82_: str
            d_1993_v82_ = _dafny.CodePoint('l')
            d_1994_v83_: D2
            d_1994_v83_ = D2_DC9(True, p0, d_1993_v82_)
            (globalState).f9 = (d_1994_v83_).cf13
            d_1995_v84_: bool
            d_1995_v84_ = True
            d_1996_v85_: _dafny.Map
            d_1996_v85_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1993_v82_ for d_1997_i13_ in range(default__.abs(409))]): d_1995_v84_})
            d_1996_v85_ = ((d_1996_v85_) | (d_1996_v85_)) | (d_1996_v85_)
            d_1998_v86_: _dafny.MultiSet
            d_1998_v86_ = _dafny.MultiSet([d_1995_v84_, d_1995_v84_])
            d_1999_v87_: C4
            nw325_ = C4()
            nw325_.ctor__(d_1995_v84_, not(d_1995_v84_), (d_1998_v86_) - (d_1998_v86_))
            d_1999_v87_ = nw325_
            (globalState).f9 = (d_1999_v87_).f19
        d_2000_v88_: bool
        d_2000_v88_ = True
        if d_2000_v88_:
            d_2001_v89_: _dafny.MultiSet
            d_2001_v89_ = _dafny.MultiSet([False])
            d_2002_v90_: _dafny.Map
            d_2002_v90_ = _dafny.Map({d_2000_v88_: d_2000_v88_})
            d_2003_v91_: _dafny.Array
            def lambda95_(d_2004_v88_):
                def lambda96_(d_2005_i14_):
                    return d_2004_v88_

                return lambda96_

            init55_ = lambda95_(d_2000_v88_)
            nw326_ = _dafny.Array(None, 15)
            for i0_55_ in range(nw326_.length(0)):
                nw326_[i0_55_] = init55_(i0_55_)
            d_2003_v91_ = nw326_
            d_2006_v92_: D11
            d_2006_v92_ = D11_DC34(d_2001_v89_, _dafny.MultiSet([len(d_2002_v90_)]), d_2003_v91_, D1_DC5(d_2000_v88_, False))
            d_2007_v93_: _dafny.MultiSet
            d_2007_v93_ = _dafny.MultiSet([(d_2006_v92_).cf60])
            d_2008_v94_: str
            d_2008_v94_ = _dafny.CodePoint('x')
            d_2009_v95_: _dafny.Map
            d_2009_v95_ = _dafny.Map({(p0) - ((d_2007_v93_).cardinality): d_2008_v94_})
            d_2010_v96_: _dafny.Seq
            d_2010_v96_ = _dafny.SeqWithoutIsStrInference([p0, (d_2001_v89_).cardinality])
            d_2009_v95_ = (d_2009_v95_).set(len(d_2010_v96_), d_2008_v94_)
            d_2011_v97_: C0
            nw327_ = C0()
            nw327_.ctor__(d_2000_v88_, d_2000_v88_)
            d_2011_v97_ = nw327_
            d_2012_v98_: _dafny.Map
            d_2012_v98_ = _dafny.Map({p0: d_2011_v97_})
            d_2012_v98_ = d_2012_v98_
            d_2013_v99_: D14
            d_2013_v99_ = D14_DC45(d_2002_v90_, d_2011_v97_, p0)
            d_2014_v100_: D14
            d_2014_v100_ = D14_DC46(d_2013_v99_)
            pat_let_tv45_ = d_2013_v99_
            pat_let_tv46_ = d_2013_v99_
            d_2015_v101_: _dafny.Array
            nw328_ = _dafny.Array(None, 26)
            nw328_[int(0)] = d_2014_v100_
            nw328_[int(1)] = d_2014_v100_
            def iife204_(_pat_let64_0):
                def iife205_(d_2016_dt__update__tmp_h1_):
                    def iife206_(_pat_let65_0):
                        def iife207_(d_2017_dt__update_hcf75_h0_):
                            return D14_DC46(d_2017_dt__update_hcf75_h0_)
                        return iife207_(_pat_let65_0)
                    return iife206_(pat_let_tv45_)
                return iife205_(_pat_let64_0)
            nw328_[int(2)] = iife204_(d_2014_v100_)
            nw328_[int(3)] = d_2014_v100_
            nw328_[int(4)] = d_2014_v100_
            nw328_[int(5)] = d_2014_v100_
            nw328_[int(6)] = d_2014_v100_
            nw328_[int(7)] = d_2014_v100_
            nw328_[int(8)] = d_2014_v100_
            nw328_[int(9)] = d_2014_v100_
            nw328_[int(10)] = d_2014_v100_
            nw328_[int(11)] = d_2014_v100_
            nw328_[int(12)] = d_2014_v100_
            nw328_[int(13)] = d_2014_v100_
            nw328_[int(14)] = d_2014_v100_
            nw328_[int(15)] = d_2014_v100_
            nw328_[int(16)] = d_2014_v100_
            nw328_[int(17)] = D14_DC46(d_2013_v99_)
            nw328_[int(18)] = d_2014_v100_
            nw328_[int(19)] = d_2014_v100_
            nw328_[int(20)] = d_2014_v100_
            nw328_[int(21)] = (d_2014_v100_ if d_2000_v88_ else d_2014_v100_)
            nw328_[int(22)] = d_2014_v100_
            nw328_[int(23)] = d_2014_v100_
            def iife208_(_pat_let66_0):
                def iife209_(d_2018_dt__update__tmp_h2_):
                    def iife210_(_pat_let67_0):
                        def iife211_(d_2019_dt__update_hcf75_h1_):
                            return D14_DC46(d_2019_dt__update_hcf75_h1_)
                        return iife211_(_pat_let67_0)
                    return iife210_(pat_let_tv46_)
                return iife209_(_pat_let66_0)
            nw328_[int(24)] = iife208_(d_2014_v100_)
            nw328_[int(25)] = d_2014_v100_
            d_2015_v101_ = nw328_
            index354_ = default__.safeIndex(642, (d_2015_v101_).length(0))
            (d_2015_v101_)[index354_] = d_2014_v100_
            index355_ = default__.safeIndex(642, (d_2015_v101_).length(0))
            rhs322_ = (d_2011_v97_).f23
            rhs323_ = (d_2014_v100_ if (d_2011_v97_).f23 else d_2014_v100_)
            rhs324_ = default__.fm27(p0, d_2000_v88_, globalState)
            rhs325_ = p0
            lhs274_ = globalState
            lhs275_ = d_2015_v101_
            lhs276_ = default__.safeIndex(642, (d_2015_v101_).length(0))
            lhs277_ = globalState
            lhs274_.f9 = rhs322_
            lhs275_[lhs276_] = rhs323_
            d_1977_v73_ = rhs324_
            lhs277_.f3 = rhs325_
            d_2020_v102_: _dafny.Seq
            d_2020_v102_ = _dafny.SeqWithoutIsStrInference([(d_2011_v97_).f23])
            d_2021_v103_: D2
            d_2021_v103_ = D2_DC7(d_2020_v102_)
            if (_dafny.SeqWithoutIsStrInference([True])) == (((d_2021_v103_).cf9).set(default__.safeIndex(p0, len((d_2021_v103_).cf9)), (d_2011_v97_).f23)):
                d_2022_v104_: _dafny.Array
                nw329_ = _dafny.Array(_dafny.Map({}), 8)
                d_2022_v104_ = nw329_
                d_2023_v106_: _dafny.Array
                def lambda97_(d_2024_v89_, d_2025_v102_):
                    def lambda98_(d_2026_i15_):
                        def iife212_():
                            coll76_ = _dafny.Set()
                            compr_76_: _dafny.MultiSet
                            for compr_76_ in (_dafny.SeqWithoutIsStrInference([d_2024_v89_, d_2024_v89_, _dafny.MultiSet(d_2025_v102_)])).Elements:
                                d_2027_v105_: _dafny.MultiSet = compr_76_
                                if (d_2027_v105_) in (_dafny.SeqWithoutIsStrInference([d_2024_v89_, d_2024_v89_, _dafny.MultiSet(d_2025_v102_)])):
                                    coll76_ = coll76_.union(_dafny.Set([d_2027_v105_]))
                            return _dafny.Set(coll76_)
                        return D11_DC33(iife212_()
)

                    return lambda98_

                init56_ = lambda97_(d_2001_v89_, d_2020_v102_)
                nw330_ = _dafny.Array(None, 24)
                for i0_56_ in range(nw330_.length(0)):
                    nw330_[i0_56_] = init56_(i0_56_)
                d_2023_v106_ = nw330_
                d_2028_v107_: _dafny.Map
                d_2028_v107_ = _dafny.Map({d_2008_v94_: (0) - (p0)})
                d_2029_v108_: _dafny.Set
                d_2029_v108_ = _dafny.Set({d_2010_v96_})
                d_2030_v109_: _dafny.Map
                d_2030_v109_ = _dafny.Map({d_2023_v106_: default__.fm23(p0, (d_2011_v97_).f23, d_2028_v107_, d_2029_v108_, globalState)})
                index356_ = default__.safeIndex(440, (d_2022_v104_).length(0))
                (d_2022_v104_)[index356_] = d_2030_v109_
                index357_ = default__.safeIndex(440, (d_2022_v104_).length(0))
                (d_2022_v104_)[index357_] = (d_2030_v109_) | (_dafny.Map({d_2023_v106_: (d_2011_v97_).f23}))
                (globalState).f9 = not ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yacasvvoi")))) > (p0)) or (d_2000_v88_)
                d_2031_v110_: _dafny.Map
                d_2031_v110_ = _dafny.Map({default__.safeModuloInt(p0, p0): p0})
                d_2031_v110_ = (d_2031_v110_) | (d_2031_v110_)
                d_2032_v111_: D13
                d_2032_v111_ = D13_DC42(d_2008_v94_, (d_2011_v97_).f24)
                d_2033_v112_: D13
                d_2033_v112_ = D13_DC43(d_2032_v111_)
                d_2034_v113_: D13
                d_2034_v113_ = D13_DC43(D13_DC43(d_2032_v111_))
                d_2035_v114_: D13
                d_2035_v114_ = D13_DC43(d_2034_v113_)
                d_2036_v115_: D13
                d_2036_v115_ = D13_DC43(d_2033_v112_)
                d_2037_v116_: _dafny.Set
                d_2037_v116_ = _dafny.Set({p0})
                d_2038_v118_: _dafny.Map
                d_2038_v118_ = _dafny.Map({p0: _dafny.Set({p0})})
                d_2039_v119_: _dafny.Seq
                d_2039_v119_ = _dafny.SeqWithoutIsStrInference([d_1977_v73_])
                d_2040_v120_: D14
                d_2040_v120_ = D14_DC45(d_2002_v90_, d_2011_v97_, len((d_2039_v119_)[default__.safeIndex(p0, len(d_2039_v119_))]))
                d_2041_v124_: _dafny.Map
                d_2041_v124_ = _dafny.Map({p0: (d_2011_v97_).f24})
                d_2042_v126_: _dafny.Seq
                def iife213_():
                    coll77_ = _dafny.Map()
                    compr_77_: int
                    for compr_77_ in _dafny.IntegerRange(206, 40):
                        d_2043_v125_: int = compr_77_
                        if ((206) <= (d_2043_v125_)) and ((d_2043_v125_) < (40)):
                            coll77_[default__.safeDivisionInt(d_2043_v125_, -866)] = (d_2011_v97_).f23
                    return _dafny.Map(coll77_)
                d_2042_v126_ = _dafny.SeqWithoutIsStrInference([d_2041_v124_, (iife213_()
                ).set(p0, True)])
                d_2044_v128_: _dafny.Array
                nw331_ = _dafny.Array(int(0), 18)
                d_2044_v128_ = nw331_
                d_2045_v130_: _dafny.Array
                nw332_ = _dafny.Array(None, 19)
                nw332_[int(0)] = d_2037_v116_
                nw332_[int(1)] = d_2037_v116_
                def iife214_():
                    coll78_ = _dafny.Set()
                    compr_78_: int
                    for compr_78_ in _dafny.IntegerRange(778, 596):
                        d_2046_v117_: int = compr_78_
                        if ((778) <= (d_2046_v117_)) and ((d_2046_v117_) < (596)):
                            coll78_ = coll78_.union(_dafny.Set([(d_2046_v117_) * (p0)]))
                    return _dafny.Set(coll78_)
                nw332_[int(2)] = iife214_()
                
                nw332_[int(3)] = (((d_2038_v118_)[len((d_2040_v120_).cf72)] if (len((d_2040_v120_).cf72)) in (d_2038_v118_) else d_2037_v116_) if (d_2011_v97_).f23 else d_2037_v116_)
                def iife215_():
                    coll79_ = _dafny.Set()
                    compr_79_: int
                    for compr_79_ in _dafny.IntegerRange(407, -440):
                        d_2047_v121_: int = compr_79_
                        if ((407) <= (d_2047_v121_)) and ((d_2047_v121_) < (-440)):
                            coll79_ = coll79_.union(_dafny.Set([default__.safeModuloInt(d_2047_v121_, p0)]))
                    return _dafny.Set(coll79_)
                nw332_[int(4)] = iife215_()
                
                nw332_[int(5)] = d_2037_v116_
                nw332_[int(6)] = d_2037_v116_
                nw332_[int(7)] = d_2037_v116_
                def iife216_():
                    coll80_ = _dafny.Set()
                    compr_80_: int
                    for compr_80_ in _dafny.IntegerRange(-243, 713):
                        d_2048_v122_: int = compr_80_
                        if ((-243) <= (d_2048_v122_)) and ((d_2048_v122_) < (713)):
                            coll80_ = coll80_.union(_dafny.Set([(d_2048_v122_) * (p0)]))
                    return _dafny.Set(coll80_)
                nw332_[int(8)] = iife216_()
                
                def iife217_():
                    coll81_ = _dafny.Set()
                    compr_81_: int
                    for compr_81_ in _dafny.IntegerRange(437, -919):
                        d_2049_v123_: int = compr_81_
                        if ((437) <= (d_2049_v123_)) and ((d_2049_v123_) < (-919)):
                            coll81_ = coll81_.union(_dafny.Set([(d_2049_v123_) - (len(_dafny.Set({(d_2011_v97_).f23})))]))
                    return _dafny.Set(coll81_)
                nw332_[int(9)] = (iife217_()
                ) - (d_2037_v116_)
                nw332_[int(10)] = default__.fm22(len(d_1977_v73_), d_2042_v126_, (d_2011_v97_).f23, (d_2011_v97_).f24, globalState)
                def iife218_():
                    coll82_ = _dafny.Set()
                    compr_82_: int
                    for compr_82_ in _dafny.IntegerRange(-200, 104):
                        d_2050_v127_: int = compr_82_
                        if ((-200) <= (d_2050_v127_)) and ((d_2050_v127_) < (104)):
                            coll82_ = coll82_.union(_dafny.Set([(d_2050_v127_) + ((d_1909_v23_).cardinality)]))
                    return _dafny.Set(coll82_)
                nw332_[int(11)] = iife218_()
                
                nw332_[int(12)] = d_2037_v116_
                nw332_[int(13)] = d_2037_v116_
                nw332_[int(14)] = d_2037_v116_
                nw332_[int(15)] = _dafny.Set({p0, len(d_2042_v126_), default__.fm0(d_2000_v88_, (_dafny.MultiSet([d_2044_v128_])).cardinality, len(_dafny.SeqWithoutIsStrInference([(d_2011_v97_).f23, (d_2011_v97_).f24])), d_1909_v23_, globalState), p0, p0})
                def iife219_():
                    coll83_ = _dafny.Set()
                    compr_83_: int
                    for compr_83_ in _dafny.IntegerRange(311, 608):
                        d_2051_v129_: int = compr_83_
                        if ((311) <= (d_2051_v129_)) and ((d_2051_v129_) < (608)):
                            coll83_ = coll83_.union(_dafny.Set([(d_2051_v129_) * (len(_dafny.SeqWithoutIsStrInference([(d_2011_v97_).f24, (d_2011_v97_).f24])))]))
                    return _dafny.Set(coll83_)
                nw332_[int(16)] = iife219_()
                
                nw332_[int(17)] = d_2037_v116_
                nw332_[int(18)] = d_2037_v116_
                d_2045_v130_ = nw332_
                pat_let_tv47_ = d_2035_v114_
                def iife220_(_pat_let68_0):
                    def iife221_(d_2052_dt__update__tmp_h3_):
                        def iife222_(_pat_let69_0):
                            def iife223_(d_2053_dt__update_hcf70_h0_):
                                return D13_DC43(d_2053_dt__update_hcf70_h0_)
                            return iife223_(_pat_let69_0)
                        return iife222_(pat_let_tv47_)
                    return iife221_(_pat_let68_0)
                rhs326_ = d_2045_v130_
                rhs327_ = (d_2011_v97_).f23
                rhs328_ = p0
                rhs329_ = iife220_(default__.fm51(d_2000_v88_, (d_2011_v97_).f23, d_2000_v88_, len(d_2041_v124_), globalState))
                lhs278_ = globalState
                lhs279_ = globalState
                lhs280_ = globalState
                lhs278_.f11 = rhs326_
                lhs279_.f9 = rhs327_
                lhs280_.f3 = rhs328_
                d_2036_v115_ = rhs329_
                d_1977_v73_ = _dafny.SeqWithoutIsStrInference([(d_1977_v73_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_2011_v97_).f24])), len(d_1977_v73_))] for d_2054_i16_ in range(default__.abs(-323))])
            elif True:
                d_2055_v131_: _dafny.Array
                def lambda99_(d_2056_p0_):
                    def lambda100_(d_2057_i17_):
                        return (d_2057_i17_) - (d_2056_p0_)

                    return lambda100_

                init57_ = lambda99_(p0)
                nw333_ = _dafny.Array(None, 22)
                for i0_57_ in range(nw333_.length(0)):
                    nw333_[i0_57_] = init57_(i0_57_)
                d_2055_v131_ = nw333_
                index358_ = default__.safeIndex(927, (d_2055_v131_).length(0))
                (d_2055_v131_)[index358_] = p0
                index359_ = default__.safeIndex(927, (d_2055_v131_).length(0))
                (d_2055_v131_)[index359_] = p0
                d_2058_v132_: C3
                nw334_ = C3()
                nw334_.ctor__(default__.safeModuloInt(845, p0), (d_1977_v73_) + (d_1977_v73_), False, (d_2001_v89_) | (d_2001_v89_))
                d_2058_v132_ = nw334_
                d_2059_v134_: D9
                def iife224_():
                    coll84_ = _dafny.Set()
                    compr_84_: int
                    for compr_84_ in (d_2010_v96_).Elements:
                        d_2060_v133_: int = compr_84_
                        if (d_2060_v133_) in (d_2010_v96_):
                            coll84_ = coll84_.union(_dafny.Set([default__.safeModuloInt(d_2060_v133_, 979)]))
                    return _dafny.Set(coll84_)
                d_2059_v134_ = D9_DC30(d_2008_v94_, (d_2011_v97_).f23, len(d_2020_v102_), iife224_()
, (d_2055_v131_)[default__.safeIndex(927, (d_2055_v131_).length(0))])
                (globalState).f9 = (d_2059_v134_).cf51
                index360_ = default__.safeIndex(927, (d_2055_v131_).length(0))
                (d_2055_v131_)[index360_] = (d_2055_v131_)[default__.safeIndex(927, (d_2055_v131_).length(0))]
                index361_ = default__.safeIndex(28, (d_2003_v91_).length(0))
                (d_2003_v91_)[index361_] = ((d_2011_v97_).f23) and ((d_2011_v97_).f24)
                index362_ = default__.safeIndex(28, (d_2003_v91_).length(0))
                (d_2003_v91_)[index362_] = (d_2011_v97_).f24
            (globalState).f9 = (d_2011_v97_).f24
        elif True:
            (globalState).f9 = d_2000_v88_
            d_2061_v135_: _dafny.Map
            d_2061_v135_ = _dafny.Map({_dafny.CodePoint('x'): not(not(d_2000_v88_))})
            d_2062_v136_: _dafny.Seq
            d_2062_v136_ = _dafny.SeqWithoutIsStrInference([d_2061_v135_, (d_2061_v135_) | (default__.fm46(len(d_1977_v73_), d_2000_v88_, globalState)), d_2061_v135_, d_2061_v135_, d_2061_v135_])
            (globalState).f3 = len(d_2062_v136_)
            d_2063_v137_: _dafny.Array
            nw335_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
            d_2063_v137_ = nw335_
            index363_ = default__.safeIndex(984, (d_2063_v137_).length(0))
            (d_2063_v137_)[index363_] = d_1977_v73_
            d_2064_v138_: str
            d_2064_v138_ = _dafny.CodePoint('e')
            d_2065_v139_: _dafny.Array
            def lambda101_(d_2066_v88_):
                def lambda102_(d_2067_i18_):
                    return d_2066_v88_

                return lambda102_

            init58_ = lambda101_(d_2000_v88_)
            nw336_ = _dafny.Array(None, 14)
            for i0_58_ in range(nw336_.length(0)):
                nw336_[i0_58_] = init58_(i0_58_)
            d_2065_v139_ = nw336_
            d_2068_v140_: _dafny.Seq
            d_2068_v140_ = _dafny.SeqWithoutIsStrInference([d_2065_v139_])
            d_2069_v141_: _dafny.MultiSet
            d_2069_v141_ = _dafny.MultiSet([(d_2068_v140_)[default__.safeIndex(p0, len(d_2068_v140_))]])
            index364_ = default__.safeIndex(984, (d_2063_v137_).length(0))
            (d_2063_v137_)[index364_] = (((default__.fm27(p0, d_2000_v88_, globalState)).set(default__.safeIndex(p0, len(default__.fm27(p0, d_2000_v88_, globalState))), d_2064_v138_)).set(default__.safeIndex(-50, len((default__.fm27(p0, d_2000_v88_, globalState)).set(default__.safeIndex(p0, len(default__.fm27(p0, d_2000_v88_, globalState))), d_2064_v138_))), d_2064_v138_)).set(default__.safeIndex((d_2069_v141_).cardinality, len(((default__.fm27(p0, d_2000_v88_, globalState)).set(default__.safeIndex(p0, len(default__.fm27(p0, d_2000_v88_, globalState))), d_2064_v138_)).set(default__.safeIndex(-50, len((default__.fm27(p0, d_2000_v88_, globalState)).set(default__.safeIndex(p0, len(default__.fm27(p0, d_2000_v88_, globalState))), d_2064_v138_))), d_2064_v138_))), d_2064_v138_)
            d_2070_v142_: _dafny.Array
            nw337_ = _dafny.Array(int(0), 8)
            d_2070_v142_ = nw337_
            index365_ = default__.safeIndex(4, (d_2070_v142_).length(0))
            (d_2070_v142_)[index365_] = p0
            index366_ = default__.safeIndex(4, (d_2070_v142_).length(0))
            (d_2070_v142_)[index366_] = ((len((d_2063_v137_)[default__.safeIndex(984, (d_2063_v137_).length(0))])) + (p0)) - (p0)
            d_2064_v138_ = d_2064_v138_

    def m4(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        (globalState).f9 = p0
        d_2071_v0_: _dafny.Array
        def lambda103_(d_2072_i1_):
            return default__.safeDivisionInt(d_2072_i1_, 150)

        init59_ = lambda103_
        nw338_ = _dafny.Array(None, 27)
        for i0_59_ in range(nw338_.length(0)):
            nw338_[i0_59_] = init59_(i0_59_)
        d_2071_v0_ = nw338_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_2071_v0_).length(0)):
            d_2073_i0_: int = guard_loop_6_
            if (True) and (((0) <= (d_2073_i0_)) and ((d_2073_i0_) < ((d_2071_v0_).length(0)))):
                (d_2071_v0_)[(d_2073_i0_)] = (d_2073_i0_) + ((D3_DC11(463, p2, p0, len(_dafny.Map({p0: p2})))).cf18)
        d_2074_v1_: _dafny.Set
        d_2074_v1_ = _dafny.Set({p0, p0})
        d_2075_v2_: _dafny.Map
        d_2075_v2_ = _dafny.Map({(d_2074_v1_) | (default__.fm41(p2, p1, globalState)): p0})
        (globalState).f3 = len(d_2075_v2_)
        d_2076_v3_: _dafny.Array
        nw339_ = _dafny.Array(None, 2)
        nw339_[int(0)] = False
        nw339_[int(1)] = p0
        d_2076_v3_ = nw339_
        index367_ = default__.safeIndex(262, (d_2076_v3_).length(0))
        (d_2076_v3_)[index367_] = p0
        index368_ = default__.safeIndex(262, (d_2076_v3_).length(0))
        (d_2076_v3_)[index368_] = p0
        d_2077_v4_: _dafny.Map
        d_2077_v4_ = _dafny.Map({p0: (d_2076_v3_)[default__.safeIndex(262, (d_2076_v3_).length(0))]})
        index369_ = default__.safeIndex(107, (d_2071_v0_).length(0))
        (d_2071_v0_)[index369_] = (0) - (p1)
        index370_ = default__.safeIndex(262, (d_2076_v3_).length(0))
        index371_ = default__.safeIndex(107, (d_2071_v0_).length(0))
        rhs330_ = d_2077_v4_
        rhs331_ = (p1) != ((0) - ((p1) - (p2)))
        rhs332_ = p1
        rhs333_ = p2
        lhs281_ = d_2076_v3_
        lhs282_ = default__.safeIndex(262, (d_2076_v3_).length(0))
        lhs283_ = d_2071_v0_
        lhs284_ = default__.safeIndex(107, (d_2071_v0_).length(0))
        lhs285_ = globalState
        d_2077_v4_ = rhs330_
        lhs281_[lhs282_] = rhs331_
        lhs283_[lhs284_] = rhs332_
        lhs285_.f3 = rhs333_
        d_2078_v5_: str
        d_2078_v5_ = _dafny.CodePoint('c')
        d_2079_v6_: _dafny.MultiSet
        d_2079_v6_ = _dafny.MultiSet([(d_2071_v0_)[default__.safeIndex(107, (d_2071_v0_).length(0))], 441])
        d_2080_v7_: _dafny.MultiSet
        d_2080_v7_ = _dafny.MultiSet([False])
        d_2081_v8_: T0
        nw340_ = C5()
        nw340_.ctor__((d_2071_v0_)[default__.safeIndex(107, (d_2071_v0_).length(0))], p0, d_2080_v7_)
        d_2081_v8_ = nw340_
        d_2082_v9_: _dafny.MultiSet
        d_2082_v9_ = _dafny.MultiSet([d_2081_v8_, d_2081_v8_])
        d_2083_v10_: _dafny.Seq
        d_2083_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnbtloi"))
        d_2084_v11_: _dafny.Set
        d_2084_v11_ = _dafny.Set({(d_2071_v0_)[default__.safeIndex(107, (d_2071_v0_).length(0))], default__.fm0((d_2076_v3_)[default__.safeIndex(262, (d_2076_v3_).length(0))], (d_2071_v0_)[default__.safeIndex(107, (d_2071_v0_).length(0))], (d_2071_v0_)[default__.safeIndex(107, (d_2071_v0_).length(0))], (d_2079_v6_).set((d_2082_v9_).cardinality, default__.abs(p1)), globalState), len(d_2083_v10_), p1})
        index372_ = default__.safeIndex(107, (d_2071_v0_).length(0))
        index373_ = default__.safeIndex(107, (d_2071_v0_).length(0))
        rhs334_ = (d_2076_v3_)[default__.safeIndex(262, (d_2076_v3_).length(0))]
        rhs335_ = default__.safeModuloInt((d_2071_v0_)[default__.safeIndex(107, (d_2071_v0_).length(0))], len(d_2084_v11_))
        rhs336_ = d_2078_v5_
        rhs337_ = (0) - (p2)
        lhs286_ = globalState
        lhs287_ = d_2071_v0_
        lhs288_ = default__.safeIndex(107, (d_2071_v0_).length(0))
        lhs289_ = d_2071_v0_
        lhs290_ = default__.safeIndex(107, (d_2071_v0_).length(0))
        lhs286_.f9 = rhs334_
        lhs287_[lhs288_] = rhs335_
        d_2078_v5_ = rhs336_
        lhs289_[lhs290_] = rhs337_
        d_2085_v12_: _dafny.Seq
        d_2085_v12_ = _dafny.SeqWithoutIsStrInference([p0, True, True])
        r0 = (_dafny.SeqWithoutIsStrInference([(d_2076_v3_)[default__.safeIndex(262, (d_2076_v3_).length(0))]])) < (d_2085_v12_)
        return r0


class C8(T0):
    def  __init__(self):
        self._f14: bool = False
        self._f15: _dafny.MultiSet = _dafny.MultiSet({})
        self._f16: _dafny.Map = _dafny.Map({})
        self._f17: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    @property
    def f15(self):
        return self._f15
    def ctor__(self, f16, f17, f14, f15):
        (self)._f16 = f16
        (self)._f17 = f17
        (self).f14 = f14
        (self)._f15 = f15

    def fm2(self, p0, p1, globalState):
        return self.f14

    def fm3(self, p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_2086_i0_ in range(default__.abs(-738))]) if self.f14 else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_2087_i1_ in range(default__.abs(123))]))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jshowx")))

    def fm4(self, p0, p1, globalState):
        return self.f14

    def m0(self, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        r3: bool = False
        d_2088_v0_: _dafny.Array
        nw341_ = _dafny.Array(None, 8)
        nw341_[int(0)] = not(self.f14)
        nw341_[int(1)] = self.f14
        nw341_[int(2)] = self.f14
        nw341_[int(3)] = self.f14
        nw341_[int(4)] = self.f14
        nw341_[int(5)] = self.f14
        nw341_[int(6)] = self.f14
        nw341_[int(7)] = self.f14
        d_2088_v0_ = nw341_
        d_2089_v1_: _dafny.MultiSet
        d_2089_v1_ = _dafny.MultiSet([d_2088_v0_, d_2088_v0_])
        if (d_2088_v0_) not in (d_2089_v1_):
            d_2090_v2_: _dafny.Set
            d_2090_v2_ = _dafny.Set({not(self.f14), self.f14, self.f14, self.f14})
            (globalState).f9 = (d_2090_v2_) != (d_2090_v2_)
            if (True if self.f14 else self.f14):
                r3 = not(self.f14)
                d_2091_v3_: _dafny.Seq
                d_2091_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okwamfa"))
                r3 = (self).fm2(self.f14, (d_2091_v3_) == (d_2091_v3_), globalState)
                (globalState).f13 = _dafny.Set({self.f14})
                d_2092_v4_: int
                d_2092_v4_ = 431
                (globalState).f3 = (0) - (d_2092_v4_)
                d_2093_v5_: _dafny.Seq
                d_2093_v5_ = _dafny.SeqWithoutIsStrInference([d_2092_v4_])
                d_2094_v6_: _dafny.Seq
                d_2094_v6_ = _dafny.SeqWithoutIsStrInference([d_2093_v5_, d_2093_v5_, d_2093_v5_])
                index374_ = default__.safeIndex(425, (d_2088_v0_).length(0))
                (d_2088_v0_)[index374_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahnxhipsh"))) == (d_2091_v3_)
                index375_ = default__.safeIndex(425, (d_2088_v0_).length(0))
                rhs338_ = ((d_2094_v6_ if self.f14 else d_2094_v6_)) + (d_2094_v6_)
                rhs339_ = (self.f14 if ((self).f15) == ((self).f15) else self.f14)
                lhs291_ = d_2088_v0_
                lhs292_ = default__.safeIndex(425, (d_2088_v0_).length(0))
                d_2094_v6_ = rhs338_
                lhs291_[lhs292_] = rhs339_
            elif True:
                d_2095_v7_: _dafny.Array
                nw342_ = _dafny.Array(int(0), 3)
                d_2095_v7_ = nw342_
                d_2096_v8_: int
                d_2096_v8_ = 726
                index376_ = default__.safeIndex(68, (d_2095_v7_).length(0))
                (d_2095_v7_)[index376_] = default__.fm0(self.f14, 325, d_2096_v8_, _dafny.MultiSet([d_2096_v8_, 370, d_2096_v8_, d_2096_v8_, d_2096_v8_]), globalState)
                d_2097_v9_: _dafny.Seq
                d_2097_v9_ = _dafny.SeqWithoutIsStrInference([-348])
                d_2098_v10_: _dafny.Map
                d_2098_v10_ = _dafny.Map({d_2096_v8_: d_2096_v8_})
                d_2099_v11_: _dafny.Map
                d_2099_v11_ = _dafny.Map({d_2097_v9_: d_2098_v10_})
                d_2100_v12_: _dafny.MultiSet
                d_2100_v12_ = _dafny.MultiSet([d_2096_v8_, len(default__.fm5(self.f14, globalState)), len(d_2099_v11_)])
                index377_ = default__.safeIndex(68, (d_2095_v7_).length(0))
                (d_2095_v7_)[index377_] = default__.fm0(self.f14, d_2096_v8_, d_2096_v8_, d_2100_v12_, globalState)
                d_2101_v13_: D0
                d_2101_v13_ = D0_DC1(self.f14)
                d_2102_v14_: _dafny.Seq
                d_2102_v14_ = _dafny.SeqWithoutIsStrInference([(d_2101_v13_).cf1])
                d_2103_v15_: _dafny.Map
                d_2103_v15_ = _dafny.Map({d_2097_v9_: len(d_2102_v14_)})
                d_2104_v16_: _dafny.Map
                d_2104_v16_ = _dafny.Map({((d_2095_v7_)[default__.safeIndex(68, (d_2095_v7_).length(0))] if self.f14 else ((d_2103_v15_)[d_2097_v9_] if (d_2097_v9_) in (d_2103_v15_) else (d_2095_v7_)[default__.safeIndex(68, (d_2095_v7_).length(0))])): (self.f14) == (self.f14)})
                d_2105_v17_: D0
                d_2105_v17_ = D0_DC0(self.f14)
                d_2104_v16_ = (d_2104_v16_).set(default__.safeModuloInt(d_2096_v8_, d_2096_v8_), (d_2105_v17_).cf0)
                d_2106_v18_: _dafny.Map
                d_2106_v18_ = _dafny.Map({self.f14: d_2095_v7_})
                d_2107_v19_: _dafny.Set
                d_2107_v19_ = _dafny.Set({d_2096_v8_, len(d_2106_v18_), (d_2095_v7_)[default__.safeIndex(68, (d_2095_v7_).length(0))], d_2096_v8_})
                d_2108_v20_: _dafny.Seq
                d_2108_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
                d_2109_v21_: str
                out59_: str
                out59_ = (self).m2(default__.safeModuloInt(len(_dafny.Map({self.f14: len(d_2107_v19_)})), d_2096_v8_), (d_2108_v20_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fmlv"))), self.f14, not (self.f14) or (self.f14), globalState)
                d_2109_v21_ = out59_
                d_2110_v23_: _dafny.Map
                d_2110_v23_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkl")): True})
                d_2111_v24_: _dafny.Map
                def iife225_():
                    coll85_ = _dafny.Map()
                    compr_85_: _dafny.Seq
                    for compr_85_ in (d_2110_v23_).keys.Elements:
                        d_2112_v22_: _dafny.Seq = compr_85_
                        if (d_2112_v22_) in (d_2110_v23_):
                            coll85_[d_2112_v22_] = _dafny.SeqWithoutIsStrInference([d_2109_v21_ for d_2113_i0_ in range(default__.abs(549))])
                    return _dafny.Map(coll85_)
                d_2111_v24_ = _dafny.Map({D0_DC0(False): (len(iife225_()
                )) > (len(d_2108_v20_))})
                (self).f14 = ((d_2111_v24_)[d_2105_v17_] if (d_2105_v17_) in (d_2111_v24_) else (len(d_2108_v20_)) > (264))
                d_2114_v25_: _dafny.Map
                d_2114_v25_ = _dafny.Map({self.f14: 562})
                d_2115_v26_: _dafny.Map
                d_2115_v26_ = _dafny.Map({(d_2095_v7_)[default__.safeIndex(68, (d_2095_v7_).length(0))]: d_2114_v25_})
                d_2114_v25_ = ((d_2115_v26_)[(0) - ((945) - ((d_2095_v7_)[default__.safeIndex(68, (d_2095_v7_).length(0))]))] if ((0) - ((945) - ((d_2095_v7_)[default__.safeIndex(68, (d_2095_v7_).length(0))]))) in (d_2115_v26_) else (d_2114_v25_).set(self.f14, len(_dafny.Set({self.f14, self.f14}))))
            d_2116_v27_: int
            d_2116_v27_ = -905
            d_2117_v28_: _dafny.Seq
            d_2117_v28_ = _dafny.SeqWithoutIsStrInference([(d_2116_v27_ if True else 3), d_2116_v27_])
            d_2117_v28_ = d_2117_v28_
            r0 = d_2116_v27_
            r3 = not((self.f14 if self.f14 else self.f14))
        elif True:
            d_2118_v29_: _dafny.Seq
            d_2118_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vo"))
            d_2119_v30_: _dafny.Array
            def lambda104_(d_2120_i1_):
                return D0_DC0(self.f14)

            init60_ = lambda104_
            nw343_ = _dafny.Array(None, 17)
            for i0_60_ in range(nw343_.length(0)):
                nw343_[i0_60_] = init60_(i0_60_)
            d_2119_v30_ = nw343_
            d_2121_v31_: _dafny.Array
            nw344_ = _dafny.Array(None, 28)
            nw344_[int(0)] = d_2119_v30_
            nw344_[int(1)] = d_2119_v30_
            nw344_[int(2)] = d_2119_v30_
            nw344_[int(3)] = d_2119_v30_
            nw344_[int(4)] = d_2119_v30_
            nw344_[int(5)] = d_2119_v30_
            nw344_[int(6)] = d_2119_v30_
            nw344_[int(7)] = d_2119_v30_
            nw344_[int(8)] = d_2119_v30_
            nw344_[int(9)] = d_2119_v30_
            nw344_[int(10)] = d_2119_v30_
            nw344_[int(11)] = d_2119_v30_
            nw344_[int(12)] = d_2119_v30_
            nw344_[int(13)] = d_2119_v30_
            nw344_[int(14)] = d_2119_v30_
            nw344_[int(15)] = d_2119_v30_
            nw344_[int(16)] = d_2119_v30_
            nw344_[int(17)] = d_2119_v30_
            nw344_[int(18)] = d_2119_v30_
            nw344_[int(19)] = d_2119_v30_
            nw344_[int(20)] = d_2119_v30_
            nw344_[int(21)] = d_2119_v30_
            nw344_[int(22)] = d_2119_v30_
            nw344_[int(23)] = d_2119_v30_
            nw344_[int(24)] = d_2119_v30_
            nw344_[int(25)] = d_2119_v30_
            nw344_[int(26)] = d_2119_v30_
            nw344_[int(27)] = d_2119_v30_
            d_2121_v31_ = nw344_
            d_2122_v32_: _dafny.Map
            d_2122_v32_ = _dafny.Map({d_2118_v29_: d_2121_v31_})
            d_2122_v32_ = (d_2122_v32_).set(d_2118_v29_, d_2121_v31_)
            d_2123_v33_: int
            d_2123_v33_ = 411
            d_2124_v34_: _dafny.Seq
            d_2124_v34_ = _dafny.SeqWithoutIsStrInference([d_2123_v33_])
            d_2125_v35_: _dafny.Set
            d_2125_v35_ = _dafny.Set({len(d_2124_v34_), d_2123_v33_, d_2123_v33_})
            (globalState).f3 = len((default__.fm6(globalState)) - (d_2125_v35_))
            r0 = -607
            d_2126_v36_: _dafny.Array
            nw345_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
            d_2126_v36_ = nw345_
            d_2127_v37_: _dafny.Map
            d_2127_v37_ = _dafny.Map({self.f14: self.f14})
            d_2128_v38_: _dafny.Map
            d_2128_v38_ = _dafny.Map({(default__.fm7(globalState) if not(self.f14) else d_2127_v37_): d_2123_v33_})
            def iife226_():
                coll86_ = _dafny.Map()
                compr_86_: _dafny.Map
                for compr_86_ in (_dafny.SeqWithoutIsStrInference([(d_2127_v37_).set(False, self.f14) for d_2129_i2_ in range(default__.abs(646))])).Elements:
                    d_2130_v39_: _dafny.Map = compr_86_
                    if (d_2130_v39_) in (_dafny.SeqWithoutIsStrInference([(d_2127_v37_).set(False, self.f14) for d_2129_i2_ in range(default__.abs(646))])):
                        coll86_[d_2130_v39_] = d_2123_v33_
                return _dafny.Map(coll86_)
            rhs340_ = d_2126_v36_
            rhs341_ = iife226_()

            rhs342_ = (d_2118_v29_) + (d_2118_v29_)
            d_2126_v36_ = rhs340_
            d_2128_v38_ = rhs341_
            d_2118_v29_ = rhs342_
            d_2131_v40_: _dafny.Array
            nw346_ = _dafny.Array(int(0), 15)
            d_2131_v40_ = nw346_
            d_2132_v41_: _dafny.Map
            d_2132_v41_ = _dafny.Map({d_2123_v33_: d_2131_v40_})
            d_2132_v41_ = (d_2132_v41_).set(d_2123_v33_, d_2131_v40_)
        d_2133_v42_: _dafny.Set
        d_2133_v42_ = _dafny.Set({self.f14})
        d_2134_v43_: int
        d_2134_v43_ = 703
        d_2135_v44_: str
        d_2135_v44_ = _dafny.CodePoint('h')
        d_2136_v45_: _dafny.Map
        d_2136_v45_ = _dafny.Map({d_2134_v43_: d_2135_v44_})
        d_2137_v46_: _dafny.Map
        d_2137_v46_ = _dafny.Map({(_dafny.Set({self.f14})).isdisjoint(d_2133_v42_): d_2136_v45_})
        d_2137_v46_ = (d_2137_v46_).set(self.f14, _dafny.Map({d_2134_v43_: d_2135_v44_}))
        d_2138_v47_: _dafny.Seq
        d_2138_v47_ = _dafny.SeqWithoutIsStrInference([self.f14, self.f14])
        d_2139_v48_: _dafny.Map
        d_2139_v48_ = _dafny.Map({d_2138_v47_: not(self.f14)})
        d_2139_v48_ = (d_2139_v48_).set(d_2138_v47_, (self).fm2(self.f14, self.f14, globalState))
        d_2140_v49_: _dafny.Map
        d_2140_v49_ = _dafny.Map({self.f14: not(self.f14)})
        d_2141_i3_: int
        d_2141_i3_ = 0
        with _dafny.label("5"):
            while (self).fm4(d_2140_v49_, self.f14, globalState):
                with _dafny.c_label("5"):
                    if (d_2141_i3_) >= (100):
                        raise _dafny.Break("5")
                    d_2141_i3_ = (d_2141_i3_) + (1)
                    d_2142_v50_: _dafny.Map
                    d_2142_v50_ = _dafny.Map({self.f14: d_2134_v43_})
                    d_2143_v51_: _dafny.MultiSet
                    d_2143_v51_ = _dafny.MultiSet([501])
                    (self).f14 = (d_2143_v51_).ispropersubset(_dafny.MultiSet([len(d_2142_v50_)]))
                    d_2144_v52_: _dafny.Map
                    d_2144_v52_ = _dafny.Map({(d_2143_v51_) - (d_2143_v51_): self.f14})
                    d_2144_v52_ = d_2144_v52_
                    if (D0_DC0((d_2138_v47_)[default__.safeIndex(d_2134_v43_, len(d_2138_v47_))])).cf0:
                        d_2145_v53_: D1
                        d_2145_v53_ = D1_DC4(d_2134_v43_, d_2134_v43_)
                        (globalState).f9 = ((0) - (d_2134_v43_)) == ((d_2145_v53_).cf4)
                        d_2146_v54_: _dafny.Seq
                        d_2146_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfaglrkli"))
                        d_2147_v55_: _dafny.Map
                        d_2147_v55_ = _dafny.Map({self.f14: _dafny.SeqWithoutIsStrInference([d_2134_v43_, len(d_2146_v54_)])})
                        d_2148_v56_: _dafny.Seq
                        d_2148_v56_ = _dafny.SeqWithoutIsStrInference([d_2134_v43_, d_2134_v43_])
                        d_2149_v57_: _dafny.Map
                        d_2149_v57_ = _dafny.Map({self.f14: (((d_2147_v55_)[self.f14] if (self.f14) in (d_2147_v55_) else d_2148_v56_)).set(default__.safeIndex(d_2134_v43_, len(((d_2147_v55_)[self.f14] if (self.f14) in (d_2147_v55_) else d_2148_v56_))), d_2134_v43_)})
                        d_2147_v55_ = ((_dafny.Map({self.f14: _dafny.SeqWithoutIsStrInference([d_2134_v43_])})) | (default__.fm8(self.f14, self.f14, len(default__.fm6(globalState)), self.f14, globalState))) | ((d_2149_v57_ if self.f14 else d_2147_v55_))
                        d_2150_v58_: _dafny.Map
                        d_2150_v58_ = _dafny.Map({((d_2140_v49_)[self.f14] if (self.f14) in (d_2140_v49_) else self.f14): d_2145_v53_})
                        d_2150_v58_ = (d_2150_v58_).set((d_2138_v47_)[default__.safeIndex(d_2134_v43_, len(d_2138_v47_))], D1_DC4(d_2134_v43_, d_2134_v43_))
                        (globalState).f9 = self.f14
                        (globalState).f9 = False
                    elif True:
                        d_2151_v59_: _dafny.Array
                        nw347_ = _dafny.Array(D1.default()(), 15)
                        d_2151_v59_ = nw347_
                        d_2152_v60_: _dafny.Array
                        nw348_ = _dafny.Array(None, 28)
                        nw348_[int(0)] = d_2151_v59_
                        nw348_[int(1)] = d_2151_v59_
                        nw348_[int(2)] = d_2151_v59_
                        nw348_[int(3)] = d_2151_v59_
                        nw348_[int(4)] = d_2151_v59_
                        nw348_[int(5)] = d_2151_v59_
                        nw348_[int(6)] = d_2151_v59_
                        nw348_[int(7)] = d_2151_v59_
                        nw348_[int(8)] = d_2151_v59_
                        nw348_[int(9)] = d_2151_v59_
                        nw348_[int(10)] = d_2151_v59_
                        nw348_[int(11)] = (d_2151_v59_ if self.f14 else d_2151_v59_)
                        nw348_[int(12)] = d_2151_v59_
                        nw348_[int(13)] = d_2151_v59_
                        nw348_[int(14)] = d_2151_v59_
                        nw348_[int(15)] = d_2151_v59_
                        nw348_[int(16)] = d_2151_v59_
                        nw348_[int(17)] = d_2151_v59_
                        nw348_[int(18)] = d_2151_v59_
                        nw348_[int(19)] = d_2151_v59_
                        nw348_[int(20)] = d_2151_v59_
                        nw348_[int(21)] = d_2151_v59_
                        nw348_[int(22)] = d_2151_v59_
                        nw348_[int(23)] = d_2151_v59_
                        nw348_[int(24)] = d_2151_v59_
                        nw348_[int(25)] = d_2151_v59_
                        nw348_[int(26)] = d_2151_v59_
                        nw348_[int(27)] = d_2151_v59_
                        d_2152_v60_ = nw348_
                        index378_ = default__.safeIndex(33, (d_2152_v60_).length(0))
                        (d_2152_v60_)[index378_] = d_2151_v59_
                        d_2153_v61_: D1
                        d_2153_v61_ = D1_DC4(d_2134_v43_, d_2134_v43_)
                        index379_ = default__.safeIndex(33, (d_2152_v60_).length(0))
                        rhs343_ = (d_2151_v59_ if (False) and (self.f14) else d_2151_v59_)
                        rhs344_ = d_2153_v61_
                        rhs345_ = d_2134_v43_
                        lhs293_ = d_2152_v60_
                        lhs294_ = default__.safeIndex(33, (d_2152_v60_).length(0))
                        lhs293_[lhs294_] = rhs343_
                        d_2153_v61_ = rhs344_
                        d_2134_v43_ = rhs345_
                        d_2154_v62_: _dafny.Seq
                        d_2154_v62_ = _dafny.SeqWithoutIsStrInference([d_2088_v0_, d_2088_v0_, d_2088_v0_])
                        d_2154_v62_ = ((d_2154_v62_).set(default__.safeIndex(-990, len(d_2154_v62_)), d_2088_v0_)).set(default__.safeIndex(d_2134_v43_, len((d_2154_v62_).set(default__.safeIndex(-990, len(d_2154_v62_)), d_2088_v0_))), d_2088_v0_)
                        d_2155_v63_: _dafny.Seq
                        d_2155_v63_ = _dafny.SeqWithoutIsStrInference([d_2138_v47_])
                        d_2156_v64_: C6
                        nw349_ = C6()
                        nw349_.ctor__(self.f14, _dafny.MultiSet((d_2155_v63_)[default__.safeIndex(d_2134_v43_, len(d_2155_v63_))]))
                        d_2156_v64_ = nw349_
                        d_2157_v65_: _dafny.Seq
                        d_2157_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehllenung"))
                        d_2157_v65_ = d_2157_v65_
                        d_2158_v66_: C7
                        nw350_ = C7()
                        nw350_.ctor__()
                        d_2158_v66_ = nw350_
                        d_2158_v66_ = d_2158_v66_
                    d_2159_v67_: _dafny.Array
                    nw351_ = _dafny.Array(None, 22)
                    d_2159_v67_ = nw351_
                    d_2159_v67_ = d_2159_v67_
                    pass
            pass
        d_2160_v69_: _dafny.Array
        def lambda105_(d_2161_v44_, d_2162_v43_):
            def lambda106_(d_2163_i4_):
                def iife227_():
                    coll87_ = _dafny.Map()
                    compr_87_: int
                    for compr_87_ in _dafny.IntegerRange(-553, 967):
                        d_2164_v68_: int = compr_87_
                        if ((-553) <= (d_2164_v68_)) and ((d_2164_v68_) < (967)):
                            coll87_[(d_2164_v68_) * (d_2162_v43_)] = _dafny.SeqWithoutIsStrInference([d_2162_v43_ for d_2165_i5_ in range(default__.abs(636))])
                    return _dafny.Map(coll87_)
                return (iife227_()
                 if self.f14 else _dafny.Map({(_dafny.MultiSet([d_2161_v44_, _dafny.CodePoint('v')])).cardinality: (_dafny.SeqWithoutIsStrInference([d_2162_v43_ for d_2166_i6_ in range(default__.abs(281))]))}))

            return lambda106_

        init61_ = lambda105_(d_2135_v44_, d_2134_v43_)
        nw352_ = _dafny.Array(None, 10)
        for i0_61_ in range(nw352_.length(0)):
            nw352_[i0_61_] = init61_(i0_61_)
        d_2160_v69_ = nw352_
        d_2167_v70_: C1
        nw353_ = C1()
        nw353_.ctor__()
        d_2167_v70_ = nw353_
        d_2168_v71_: _dafny.Map
        d_2168_v71_ = _dafny.Map({self.f14: d_2167_v70_})
        d_2169_v72_: _dafny.Seq
        d_2169_v72_ = _dafny.SeqWithoutIsStrInference([len(d_2168_v71_)])
        d_2170_v73_: _dafny.Map
        d_2170_v73_ = _dafny.Map({d_2134_v43_: d_2169_v72_})
        index380_ = default__.safeIndex(793, (d_2160_v69_).length(0))
        (d_2160_v69_)[index380_] = (d_2170_v73_) | (d_2170_v73_)
        d_2171_v74_: _dafny.Set
        d_2171_v74_ = _dafny.Set({d_2134_v43_})
        d_2172_v75_: D9
        d_2172_v75_ = D9_DC30(d_2135_v44_, self.f14, d_2134_v43_, d_2171_v74_, d_2134_v43_)
        pat_let_tv48_ = d_2134_v43_
        pat_let_tv49_ = d_2134_v43_
        index381_ = default__.safeIndex(793, (d_2160_v69_).length(0))
        def iife228_(_pat_let70_0):
            def iife229_(d_2173_dt__update__tmp_h0_):
                def iife230_(_pat_let71_0):
                    def iife231_(d_2174_dt__update_hcf54_h0_):
                        def iife232_(_pat_let72_0):
                            def iife233_(d_2175_dt__update_hcf51_h0_):
                                def iife234_(_pat_let73_0):
                                    def iife235_(d_2176_dt__update_hcf52_h0_):
                                        return D9_DC30((d_2173_dt__update__tmp_h0_).cf50, d_2175_dt__update_hcf51_h0_, d_2176_dt__update_hcf52_h0_, (d_2173_dt__update__tmp_h0_).cf53, d_2174_dt__update_hcf54_h0_)
                                    return iife235_(_pat_let73_0)
                                return iife234_(pat_let_tv49_)
                            return iife233_(_pat_let72_0)
                        return iife232_(self.f14)
                    return iife231_(_pat_let71_0)
                return iife230_(pat_let_tv48_)
            return iife229_(_pat_let70_0)
        rhs346_ = True
        rhs347_ = (iife228_(d_2172_v75_)).cf50
        rhs348_ = default__.fm52(((0) - (d_2134_v43_)) * (d_2134_v43_), globalState)
        lhs295_ = globalState
        lhs296_ = d_2160_v69_
        lhs297_ = default__.safeIndex(793, (d_2160_v69_).length(0))
        lhs295_.f9 = rhs346_
        d_2135_v44_ = rhs347_
        lhs296_[lhs297_] = rhs348_
        d_2177_v76_: _dafny.Array
        nw354_ = _dafny.Array(int(0), 2)
        d_2177_v76_ = nw354_
        d_2178_v77_: _dafny.Seq
        d_2178_v77_ = _dafny.SeqWithoutIsStrInference([d_2177_v76_, d_2177_v76_])
        d_2177_v76_ = (d_2177_v76_ if (True if self.f14 else self.f14) else (d_2178_v77_)[default__.safeIndex(d_2134_v43_, len(d_2178_v77_))])
        r0 = len(d_2171_v74_)
        d_2179_v78_: _dafny.Map
        d_2179_v78_ = _dafny.Map({287: self.f14})
        d_2180_v79_: _dafny.Seq
        d_2180_v79_ = _dafny.SeqWithoutIsStrInference([d_2179_v78_])
        r1 = (D9_DC30(d_2135_v44_, self.f14, d_2134_v43_, default__.fm22(d_2134_v43_, d_2180_v79_, self.f14, self.f14, globalState), d_2134_v43_)).cf51
        d_2181_v80_: _dafny.Map
        d_2181_v80_ = _dafny.Map({d_2179_v78_: d_2134_v43_})
        r2 = d_2181_v80_
        r3 = not(self.f14)
        return r0, r1, r2, r3

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        d_2182_v0_: _dafny.Map
        d_2182_v0_ = _dafny.Map({p1: p1})
        d_2182_v0_ = d_2182_v0_
        d_2183_v1_: _dafny.Array
        nw355_ = _dafny.Array(None, 10)
        nw355_[int(0)] = self.f14
        nw355_[int(1)] = self.f14
        nw355_[int(2)] = self.f14
        nw355_[int(3)] = self.f14
        nw355_[int(4)] = self.f14
        nw355_[int(5)] = self.f14
        nw355_[int(6)] = False
        nw355_[int(7)] = self.f14
        nw355_[int(8)] = self.f14
        nw355_[int(9)] = self.f14
        d_2183_v1_ = nw355_
        index382_ = default__.safeIndex(984, (d_2183_v1_).length(0))
        (d_2183_v1_)[index382_] = self.f14
        index383_ = default__.safeIndex(984, (d_2183_v1_).length(0))
        (d_2183_v1_)[index383_] = not(self.f14)
        d_2184_v2_: C1
        nw356_ = C1()
        nw356_.ctor__()
        d_2184_v2_ = nw356_
        (globalState).f9 = not((p1) >= (p1))
        d_2185_v3_: _dafny.Array
        nw357_ = _dafny.Array(int(0), 11)
        d_2185_v3_ = nw357_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_2185_v3_).length(0)):
            d_2186_i0_: int = guard_loop_7_
            if (True) and (((0) <= (d_2186_i0_)) and ((d_2186_i0_) < ((d_2185_v3_).length(0)))):
                (d_2185_v3_)[(d_2186_i0_)] = default__.safeDivisionInt(d_2186_i0_, p1)
        d_2187_v4_: _dafny.MultiSet
        d_2187_v4_ = _dafny.MultiSet([(len(d_2182_v0_)) == (p1)])
        d_2187_v4_ = (self).f15
        d_2188_v5_: _dafny.Seq
        d_2188_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxh"))
        r0 = (len((d_2188_v5_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))))) * (p1)
        r1 = p1
        r2 = (d_2183_v1_)[default__.safeIndex(984, (d_2183_v1_).length(0))]
        return r0, r1, r2

    def m2(self, p0, p1, p2, p3, globalState):
        r0: str = _dafny.CodePoint('D')
        d_2189_v0_: D8
        d_2189_v0_ = D8_DC27()
        d_2189_v0_ = d_2189_v0_
        (globalState).f3 = p0
        d_2190_v1_: _dafny.Array
        nw358_ = _dafny.Array(int(0), 20)
        d_2190_v1_ = nw358_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_2190_v1_).length(0)):
            d_2191_i0_: int = guard_loop_8_
            if (True) and (((0) <= (d_2191_i0_)) and ((d_2191_i0_) < ((d_2190_v1_).length(0)))):
                (d_2190_v1_)[(d_2191_i0_)] = default__.safeDivisionInt(d_2191_i0_, p0)
        d_2192_v2_: _dafny.Seq
        d_2192_v2_ = _dafny.SeqWithoutIsStrInference([self.f14, p3, p3])
        d_2193_v3_: C3
        nw359_ = C3()
        nw359_.ctor__(p0, p1, p3, (self).f15)
        d_2193_v3_ = nw359_
        d_2194_v4_: _dafny.MultiSet
        d_2194_v4_ = _dafny.MultiSet([d_2193_v3_])
        d_2195_v5_: str
        d_2195_v5_ = _dafny.CodePoint('x')
        d_2196_v6_: _dafny.Map
        d_2196_v6_ = _dafny.Map({p0: self.f14})
        d_2197_v7_: _dafny.Map
        d_2197_v7_ = _dafny.Map({d_2195_v5_: len((d_2196_v6_).set(p0, self.f14))})
        d_2198_v8_: _dafny.Set
        d_2198_v8_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(d_2193_v3_).f20 for d_2199_i1_ in range(default__.abs(14))])})
        d_2200_v9_: _dafny.Map
        d_2200_v9_ = _dafny.Map({default__.fm23(p0, (d_2192_v2_)[default__.safeIndex((d_2194_v4_).cardinality, len(d_2192_v2_))], d_2197_v7_, d_2198_v8_, globalState): (d_2193_v3_).f20})
        d_2201_v10_: _dafny.MultiSet
        d_2201_v10_ = _dafny.MultiSet([p0])
        (globalState).f3 = ((d_2200_v9_)[not((self.f14) and (self.f14))] if (not((self.f14) and (self.f14))) in (d_2200_v9_) else (0) - (default__.fm0(self.f14, len((d_2193_v3_).f21), p0, d_2201_v10_, globalState)))
        d_2202_v11_: _dafny.Array
        nw360_ = _dafny.Array(False, 14)
        d_2202_v11_ = nw360_
        (globalState).f2 = d_2202_v11_
        hi10_ = (d_2193_v3_).f20
        for d_2203_i2_ in range(p0, hi10_):
            d_2204_v12_: C5
            nw361_ = C5()
            nw361_.ctor__(p0, p3, ((self).f15).intersection((self).f15))
            d_2204_v12_ = nw361_
            d_2205_v13_: _dafny.Map
            d_2205_v13_ = _dafny.Map({751: p1})
            d_2206_v14_: D8
            d_2206_v14_ = D8_DC26(p3, d_2205_v13_)
            d_2207_v15_: _dafny.Map
            d_2207_v15_ = _dafny.Map({p2: (d_2204_v12_).fm2(p2, p3, globalState)})
            d_2208_v16_: _dafny.Map
            d_2208_v16_ = _dafny.Map({p2: ((d_2207_v15_)[p2] if (p2) in (d_2207_v15_) else p3)})
            rhs349_ = d_2204_v12_
            rhs350_ = (0) - (((d_2204_v12_).f18 if (d_2206_v14_).cf46 else default__.safeModuloInt((d_2204_v12_).f18, (0) - (len(d_2208_v16_)))))
            lhs298_ = globalState
            d_2204_v12_ = rhs349_
            lhs298_.f3 = rhs350_
            d_2209_v17_: D4
            d_2209_v17_ = D4_DC12(d_2202_v11_)
            d_2209_v17_ = d_2209_v17_
            d_2210_v18_: T0
            nw362_ = C5()
            nw362_.ctor__(p0, p3, (self).f15)
            d_2210_v18_ = nw362_
            d_2211_v19_: _dafny.Map
            d_2211_v19_ = _dafny.Map({d_2210_v18_: p2})
            d_2211_v19_ = d_2211_v19_
            d_2212_v20_: _dafny.Array
            nw363_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_2212_v20_ = nw363_
            index384_ = default__.safeIndex(385, (d_2212_v20_).length(0))
            (d_2212_v20_)[index384_] = d_2190_v1_
            index385_ = default__.safeIndex(385, (d_2212_v20_).length(0))
            (d_2212_v20_)[index385_] = d_2190_v1_
        r0 = d_2195_v5_
        return r0

    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17
