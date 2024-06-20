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
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([-111])).Elements:
                d_0_v0_: int = compr_0_
                if (d_0_v0_) in (_dafny.SeqWithoutIsStrInference([-111])):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeModuloInt(d_0_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exuvfvup"))))]))
            return _dafny.Set(coll0_)
        return (0) - (((0) - (len((D9_DC27(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "itjb")): False})), len(_dafny.Set({True, not(False)}))]))).cf40))) + (len(iife0_()
        )))

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        return False

    @staticmethod
    def fm6(p0, globalState):
        source0_ = D15_DC46()
        if source0_.is_DC44:
            d_1___mcc_h0_ = source0_.cf71
            d_2___mcc_h1_ = source0_.cf72
            d_3___mcc_h2_ = source0_.cf73
            d_4___mcc_h3_ = source0_.cf74
            d_5___mcc_h4_ = source0_.cf75
            d_6_cf75_ = d_5___mcc_h4_
            d_7_cf74_ = d_4___mcc_h3_
            d_8_cf73_ = d_3___mcc_h2_
            d_9_cf72_ = d_2___mcc_h1_
            d_10_cf71_ = d_1___mcc_h0_
            return d_9_cf72_
        elif source0_.is_DC45:
            d_11___mcc_h5_ = source0_.cf76
            d_12___mcc_h6_ = source0_.cf77
            d_13___mcc_h7_ = source0_.cf78
            d_14___mcc_h8_ = source0_.cf79
            d_15___mcc_h9_ = source0_.cf80
            d_16_cf80_ = d_15___mcc_h9_
            d_17_cf79_ = d_14___mcc_h8_
            d_18_cf78_ = d_13___mcc_h7_
            d_19_cf77_ = d_12___mcc_h6_
            d_20_cf76_ = d_11___mcc_h5_
            return d_18_cf78_
        elif source0_.is_DC46:
            return not(True)
        elif True:
            d_21___mcc_h10_ = source0_.cf70
            d_22_cf70_ = d_21___mcc_h10_
            return (not(True)) and (not(True))

    @staticmethod
    def fm7(p0, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgxt"))

    @staticmethod
    def fm8(p0, p1, globalState):
        if (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(False), not(True), False, False, False])).cardinality for d_23_i0_ in range(default__.abs(759))])) != (_dafny.SeqWithoutIsStrInference([6])):
            return D2_DC9(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txpfrb")), (0) - ((_dafny.MultiSet([-952, len(_dafny.SeqWithoutIsStrInference([True, False, False, True, False]))])).cardinality), True, -911)
        elif True:
            return D2_DC9(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "asarcrdjh")), (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([-842]), _dafny.SeqWithoutIsStrInference([-674 for d_24_i1_ in range(default__.abs(545))]), _dafny.SeqWithoutIsStrInference([46])])).cardinality, True, 585)

    @staticmethod
    def fm9(globalState):
        return default__.safeDivisionInt(444, ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wuuxa")))))) * (-128))

    @staticmethod
    def fm10(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False]))) < (((D9_DC28(_dafny.SeqWithoutIsStrInference([False]), False, len(_dafny.Map({348: 715})), False, not(True))).cf41) + (_dafny.SeqWithoutIsStrInference([True])))

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        return D4_DC13(_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False]))).cardinality}))

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        return _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqkswi")))})

    @staticmethod
    def fm15(p0, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: _dafny.Map
            for compr_1_ in ((_dafny.MultiSet([_dafny.Map({True: 12}), _dafny.Map({True: len(_dafny.Map({True: 84}))})])) - (_dafny.MultiSet([_dafny.Map({not(False): len(_dafny.Map({242: len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awlgsvyam"))}))}))})]))).Elements:
                d_25_v0_: _dafny.Map = compr_1_
                if (d_25_v0_) in ((_dafny.MultiSet([_dafny.Map({True: 12}), _dafny.Map({True: len(_dafny.Map({True: 84}))})])) - (_dafny.MultiSet([_dafny.Map({not(False): len(_dafny.Map({242: len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awlgsvyam"))}))}))})]))):
                    coll1_ = coll1_.union(_dafny.Set([d_25_v0_]))
            return _dafny.Set(coll1_)
        return iife1_()
        

    @staticmethod
    def fm17(p0, globalState):
        return (_dafny.Set({681})) - (_dafny.Set({340}))

    @staticmethod
    def fm18(p0, p1, p2, p3, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(957, -55):
                d_26_v0_: int = compr_2_
                if ((957) <= (d_26_v0_)) and ((d_26_v0_) < (-55)):
                    coll2_[(d_26_v0_) * (len(_dafny.Map({True: True})))] = -164
            return _dafny.Map(coll2_)
        return D2_DC9(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nc")), len(iife2_()
), True, (len(_dafny.Map({False: False}))) * ((_dafny.MultiSet([not(not((D9_DC29(len((D9_DC27(_dafny.SeqWithoutIsStrInference([388 for d_27_i0_ in range(default__.abs(849))]))).cf40), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqe"))): 37})), True)).cf48))])).cardinality))

    @staticmethod
    def fm19(p0, globalState):
        return ((0) - (default__.safeDivisionInt(len(_dafny.Set({True, True, False})), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "enqyyh"))))))) > ((-519) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aabgm")))))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnjmfaxg"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ms")))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_28_i0_ in range(default__.abs(519))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_29_i1_ in range(default__.abs(734))])))

    @staticmethod
    def fm21(p0, p1, p2, globalState):
        if (True) and (False):
            return D5_DC16(_dafny.Map({True: D1_DC3(_dafny.SeqWithoutIsStrInference([True]))}))
        elif True:
            return D5_DC16(_dafny.Map({False: D1_DC3(_dafny.SeqWithoutIsStrInference([False]))}))

    @staticmethod
    def fm22(p0, p1, p2, globalState):
        return _dafny.Set({True})

    @staticmethod
    def fm23(globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(not(True))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True, False, False]), _dafny.SeqWithoutIsStrInference([True, False])]))

    @staticmethod
    def fm24(globalState):
        return D1_DC4(False)

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        source1_ = D20_DC62()
        if source1_.is_DC62:
            def iife3_():
                coll3_ = _dafny.Set()
                compr_3_: _dafny.Seq
                for compr_3_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkghxcks")))]): False})).keys.Elements:
                    d_30_v0_: _dafny.Seq = compr_3_
                    if (d_30_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkghxcks")))]): False})):
                        coll3_ = coll3_.union(_dafny.Set([d_30_v0_]))
                return _dafny.Set(coll3_)
            return iife3_()
            
        elif source1_.is_DC63:
            d_31___mcc_h0_ = source1_.cf112
            d_32_cf112_ = d_31___mcc_h0_
            return _dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_33_i1_ in range(default__.abs(710))])) for d_34_i0_ in range(default__.abs(-605))]), _dafny.SeqWithoutIsStrInference([605])})
        elif source1_.is_DC61:
            d_35___mcc_h1_ = source1_.cf111
            d_36_cf111_ = d_35___mcc_h1_
            def iife4_():
                coll4_ = _dafny.Set()
                compr_4_: _dafny.Seq
                for compr_4_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([747, len(_dafny.Map({-938: not(True)}))]), _dafny.SeqWithoutIsStrInference([-799 for d_37_i2_ in range(default__.abs(-416))])])).Elements:
                    d_38_v1_: _dafny.Seq = compr_4_
                    if (d_38_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([747, len(_dafny.Map({-938: not(True)}))]), _dafny.SeqWithoutIsStrInference([-799 for d_37_i2_ in range(default__.abs(-416))])])):
                        coll4_ = coll4_.union(_dafny.Set([d_38_v1_]))
                return _dafny.Set(coll4_)
            return (D21_DC65(iife4_()
)).cf114
        elif True:
            d_39___mcc_h2_ = source1_.cf113
            d_40_cf113_ = d_39___mcc_h2_
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_42_i4_ in range(default__.abs(156))])))])).Elements:
                    d_43_v2_: int = compr_5_
                    if (d_43_v2_) in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_42_i4_ in range(default__.abs(156))])))])):
                        coll5_[(d_43_v2_) - (367)] = -245
                return _dafny.Map(coll5_)
            return (_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True, True, True}))])})) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True})), (0) - (len(_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([530 for d_41_i3_ in range(default__.abs(788))]))]), _dafny.SeqWithoutIsStrInference([len(iife5_()
            )])})))])}))

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        if (len(_dafny.Map({not(True): (_dafny.MultiSet([True])).cardinality}))) < (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqoca")))):
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: str
                for compr_6_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c'), _dafny.CodePoint('o'), _dafny.CodePoint('j'), _dafny.CodePoint('l')])).Elements:
                    d_44_v0_: str = compr_6_
                    if (d_44_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c'), _dafny.CodePoint('o'), _dafny.CodePoint('j'), _dafny.CodePoint('l')])):
                        def iife7_():
                            coll7_ = _dafny.Set()
                            compr_7_: _dafny.Set
                            for compr_7_ in (_dafny.MultiSet([_dafny.Set({len(_dafny.Map({False: True}))})])).Elements:
                                d_45_v1_: _dafny.Set = compr_7_
                                if (d_45_v1_) in (_dafny.MultiSet([_dafny.Set({len(_dafny.Map({False: True}))})])):
                                    coll7_ = coll7_.union(_dafny.Set([d_45_v1_]))
                            return _dafny.Set(coll7_)
                        def iife8_():
                            coll8_ = _dafny.Map()
                            compr_8_: int
                            for compr_8_ in (_dafny.MultiSet([(_dafny.MultiSet([not(False), not(False)])).cardinality])).Elements:
                                d_46_v2_: int = compr_8_
                                if (d_46_v2_) in (_dafny.MultiSet([(_dafny.MultiSet([not(False), not(False)])).cardinality])):
                                    coll8_[(d_46_v2_) * (len(_dafny.Map({False: False})))] = False
                            return _dafny.Map(coll8_)
                        coll6_[d_44_v0_] = (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): len(iife7_()
                        )})), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpk"))), len(iife8_()
                        )]))])))
                return _dafny.Map(coll6_)
            return (_dafny.Map({_dafny.CodePoint('a'): len(_dafny.Set({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvqqeji")))])).cardinality}))})) | (iife6_()
            )
        elif True:
            return (_dafny.Map({_dafny.CodePoint('t'): 334})) | (_dafny.Map({_dafny.CodePoint('f'): 642}))

    @staticmethod
    def fm27(p0, p1, globalState):
        return _dafny.MultiSet([default__.safeDivisionInt(453, 892), (0) - (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_47_i0_ in range(default__.abs(472))])), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbv")))))), ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, not(False), False]))).cardinality) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uibebmq")))), len((D2_DC9(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vmqtwapqp")), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdv"))), False, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True, True, True, not(not(not(False)))]))])).cardinality]))).cardinality)).cf14)])

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({True: False}))) | (_dafny.Map({False: True}))

    @staticmethod
    def fm29(p0, p1, p2, globalState):
        return D3_DC12(D3_DC11(True, False))

    @staticmethod
    def fm31(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([D6_DC21(_dafny.Set({not(False)})), D6_DC21(_dafny.Set({not(True), True})), D6_DC21(_dafny.Set({False}))])

    @staticmethod
    def fm32(p0, p1, p2, globalState):
        return _dafny.Map({False: not(not((_dafny.MultiSet([(_dafny.MultiSet([(0) - (len(_dafny.Set({False})))])).cardinality, (0) - (-252)])).ispropersubset((D17_DC52(False, False, -932, _dafny.MultiSet([162]))).cf91)))})

    @staticmethod
    def fm33(p0, p1, p2, globalState):
        return (False) and (not(False))

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exnq"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpfxxmfqm")))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_48_i0_ in range(default__.abs(432))]))

    @staticmethod
    def fm35(p0, p1, globalState):
        return (D1_DC3(_dafny.SeqWithoutIsStrInference([True, True, False]))).cf7

    @staticmethod
    def fm36(globalState):
        return ((0) - (-821)) - ((986) - (174))

    @staticmethod
    def fm37(globalState):
        source2_ = D21_DC65(_dafny.Set({_dafny.SeqWithoutIsStrInference([452 for d_49_i0_ in range(default__.abs(146))])}))
        if source2_.is_DC66:
            d_50___mcc_h0_ = source2_.cf115
            d_51_cf115_ = d_50___mcc_h0_
            return D6_DC22(d_51_cf115_, (_dafny.MultiSet([_dafny.CodePoint('q')])).cardinality, True)
        elif source2_.is_DC67:
            d_52___mcc_h1_ = source2_.cf116
            d_53_cf116_ = d_52___mcc_h1_
            return D6_DC22(len(_dafny.Map({d_53_cf116_: d_53_cf116_})), len(_dafny.Map({D16_DC48(): d_53_cf116_})), False)
        elif True:
            d_54___mcc_h2_ = source2_.cf114
            d_55_cf114_ = d_54___mcc_h2_
            return D6_DC22((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True]))).cardinality, 548, False)

    @staticmethod
    def fm38(globalState):
        if (False) not in (_dafny.SeqWithoutIsStrInference([False, False, True])):
            return D9_DC27(_dafny.SeqWithoutIsStrInference([-945]))
        elif True:
            return D9_DC27(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t')]))]))

    @staticmethod
    def fm39(p0, p1, globalState):
        return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not (not(False)) or (False), not(not (True) or (True))]))

    @staticmethod
    def fm40(p0, globalState):
        if (_dafny.Set({30})).isdisjoint(_dafny.Set({len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfejidul"))))})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_56_i0_ in range(default__.abs(184))])), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): True}))})):
            return D0_DC2(_dafny.MultiSet([False, not(not(True)), True, not(True)]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_57_i1_ in range(default__.abs(19))]), True)
        elif True:
            return D0_DC2(_dafny.MultiSet([True, True]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_58_i2_ in range(default__.abs(481))]), False)

    @staticmethod
    def fm41(p0, p1, p2, p3, globalState):
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(-602, 331):
                d_59_v0_: int = compr_9_
                if ((-602) <= (d_59_v0_)) and ((d_59_v0_) < (331)):
                    coll9_[default__.safeDivisionInt(d_59_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqcdui"))))] = True
            return _dafny.Map(coll9_)
        return (_dafny.Map({len(_dafny.Map({_dafny.SeqWithoutIsStrInference([not(True)]): D4_DC15(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rddv"))), len(_dafny.Set({(0) - (-719), -79})))})): True})) | (iife9_()
        )

    @staticmethod
    def fm42(p0, p1, globalState):
        source3_ = D5_DC18()
        if source3_.is_DC17:
            d_60___mcc_h0_ = source3_.cf27
            d_61_cf27_ = d_60___mcc_h0_
            return (D10_DC32(_dafny.CodePoint('a'), not(not(True)))).cf52
        elif source3_.is_DC18:
            return _dafny.CodePoint('g')
        elif source3_.is_DC19:
            d_62___mcc_h1_ = source3_.cf28
            d_63___mcc_h2_ = source3_.cf29
            d_64___mcc_h3_ = source3_.cf30
            d_65_cf30_ = d_64___mcc_h3_
            d_66_cf29_ = d_63___mcc_h2_
            d_67_cf28_ = d_62___mcc_h1_
            return d_65_cf30_
        elif source3_.is_DC16:
            d_68___mcc_h4_ = source3_.cf26
            d_69_cf26_ = d_68___mcc_h4_
            return _dafny.CodePoint('y')
        elif True:
            d_70___mcc_h5_ = source3_.cf31
            d_71_cf31_ = d_70___mcc_h5_
            return _dafny.CodePoint('o')

    @staticmethod
    def fm43(p0, p1, globalState):
        if not((_dafny.MultiSet([564, len(_dafny.Map({False: len(_dafny.Map({(0) - ((0) - ((0) - (len(_dafny.Map({_dafny.CodePoint('p'): False}))))): -373}))}))])).issubset(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False, True])), 776]))):
            return D1_DC3(_dafny.SeqWithoutIsStrInference([True, True]))
        elif True:
            return D1_DC3(_dafny.SeqWithoutIsStrInference([True, False]))

    @staticmethod
    def fm44(p0, p1, p2, p3, globalState):
        if (_dafny.MultiSet([not(False), True])).ispropersubset(_dafny.MultiSet([True])):
            return (_dafny.SeqWithoutIsStrInference([(D19_DC59(False, 810)).cf109 for d_72_i0_ in range(default__.abs(723))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True}))]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([len(_dafny.Set({not(False)})), 181])

    @staticmethod
    def fm45(globalState):
        def iife10_():
            def iife12_():
                coll12_ = _dafny.Map()
                compr_12_: _dafny.Seq
                for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_73_i0_ in range(default__.abs(-250))]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({992, (D19_DC59(False, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality)).cf109, 175, len(_dafny.SeqWithoutIsStrInference([False]))}))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "druvqxil"))) for d_74_i1_ in range(default__.abs(-566))])])).Elements:
                    d_75_v1_: _dafny.Seq = compr_12_
                    if (d_75_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_73_i0_ in range(default__.abs(-250))]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({992, (D19_DC59(False, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality)).cf109, 175, len(_dafny.SeqWithoutIsStrInference([False]))}))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "druvqxil"))) for d_74_i1_ in range(default__.abs(-566))])])):
                        coll12_[d_75_v1_] = _dafny.CodePoint('s')
                return _dafny.Map(coll12_)
            coll10_ = _dafny.Map()
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: _dafny.Seq
                for compr_11_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_73_i0_ in range(default__.abs(-250))]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({992, (D19_DC59(False, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality)).cf109, 175, len(_dafny.SeqWithoutIsStrInference([False]))}))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "druvqxil"))) for d_74_i1_ in range(default__.abs(-566))])])).Elements:
                    d_75_v1_: _dafny.Seq = compr_11_
                    if (d_75_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_73_i0_ in range(default__.abs(-250))]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({992, (D19_DC59(False, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality)).cf109, 175, len(_dafny.SeqWithoutIsStrInference([False]))}))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "druvqxil"))) for d_74_i1_ in range(default__.abs(-566))])])):
                        coll11_[d_75_v1_] = _dafny.CodePoint('s')
                return _dafny.Map(coll11_)
            compr_10_: _dafny.Seq
            for compr_10_ in (iife11_()
            ).keys.Elements:
                d_76_v0_: _dafny.Seq = compr_10_
                if (d_76_v0_) in (iife12_()
                ):
                    coll10_[d_76_v0_] = 590
            return _dafny.Map(coll10_)
        def iife13_():
            coll13_ = _dafny.Map()
            compr_13_: _dafny.Seq
            for compr_13_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqmycnx"))): False})) for d_77_i2_ in range(default__.abs(96))]): _dafny.Map({_dafny.CodePoint('h'): False})})).keys.Elements:
                d_78_v2_: _dafny.Seq = compr_13_
                if (d_78_v2_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqmycnx"))): False})) for d_77_i2_ in range(default__.abs(96))]): _dafny.Map({_dafny.CodePoint('h'): False})})):
                    coll13_[d_78_v2_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjboup")))
            return _dafny.Map(coll13_)
        return ((iife10_()
        ) | (iife13_()
        )) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([513, 779, len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yncuqfdbc"))}))]): (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofyfibxcp"))))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([162, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "te"))), 286]): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dv")))})))

    @staticmethod
    def fm46(p0, p1, globalState):
        return _dafny.Map({True: (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(True), True])): False})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))): True}))})

    @staticmethod
    def fm47(p0, p1, globalState):
        def iife14_():
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(-456, 4):
                    d_81_v0_: int = compr_16_
                    if ((-456) <= (d_81_v0_)) and ((d_81_v0_) < (4)):
                        coll16_[(d_81_v0_) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({False: False})): True}))])))] = D11_DC35(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rh")), True)
                return _dafny.Map(coll16_)
            coll14_ = _dafny.Set()
            def iife15_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in _dafny.IntegerRange(-456, 4):
                    d_81_v0_: int = compr_15_
                    if ((-456) <= (d_81_v0_)) and ((d_81_v0_) < (4)):
                        coll15_[(d_81_v0_) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({False: False})): True}))])))] = D11_DC35(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rh")), True)
                return _dafny.Map(coll15_)
            compr_14_: _dafny.Map
            for compr_14_ in ((_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Set({True, False})): D11_DC35(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edl")), True)}), _dafny.Map({118: D11_DC35(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tah")), True)}), _dafny.Map({len(_dafny.Set({591})): D11_DC35(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_79_i0_ in range(default__.abs(327))]), False)})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({-837: D11_DC35(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_80_i1_ in range(default__.abs(831))]), True)}), iife15_()
            ]))).Elements:
                d_82_v1_: _dafny.Map = compr_14_
                if (d_82_v1_) in ((_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Set({True, False})): D11_DC35(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edl")), True)}), _dafny.Map({118: D11_DC35(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tah")), True)}), _dafny.Map({len(_dafny.Set({591})): D11_DC35(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_79_i0_ in range(default__.abs(327))]), False)})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({-837: D11_DC35(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_80_i1_ in range(default__.abs(831))]), True)}), iife16_()
                ]))):
                    coll14_ = coll14_.union(_dafny.Set([d_82_v1_]))
            return _dafny.Set(coll14_)
        return iife14_()
        

    @staticmethod
    def fm48(globalState):
        def iife17_():
            def iife19_():
                coll19_ = _dafny.Map()
                compr_19_: int
                for compr_19_ in (_dafny.SeqWithoutIsStrInference([-922])).Elements:
                    d_83_v1_: int = compr_19_
                    if (d_83_v1_) in (_dafny.SeqWithoutIsStrInference([-922])):
                        coll19_[(d_83_v1_) + (631)] = False
                return _dafny.Map(coll19_)
            coll17_ = _dafny.Map()
            def iife18_():
                coll18_ = _dafny.Map()
                compr_18_: int
                for compr_18_ in (_dafny.SeqWithoutIsStrInference([-922])).Elements:
                    d_83_v1_: int = compr_18_
                    if (d_83_v1_) in (_dafny.SeqWithoutIsStrInference([-922])):
                        coll18_[(d_83_v1_) + (631)] = False
                return _dafny.Map(coll18_)
            compr_17_: int
            for compr_17_ in (_dafny.Set({len(_dafny.Map({(_dafny.MultiSet([False])).cardinality: len(_dafny.Map({_dafny.CodePoint('o'): (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(iife18_()
            )]), _dafny.SeqWithoutIsStrInference([743])])).cardinality}))})), 482, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhpvmy"))), len(_dafny.Set({False, True}))})).Elements:
                d_84_v0_: int = compr_17_
                if (d_84_v0_) in (_dafny.Set({len(_dafny.Map({(_dafny.MultiSet([False])).cardinality: len(_dafny.Map({_dafny.CodePoint('o'): (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(iife19_()
                )]), _dafny.SeqWithoutIsStrInference([743])])).cardinality}))})), 482, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhpvmy"))), len(_dafny.Set({False, True}))})):
                    coll17_[(d_84_v0_) * (len(_dafny.SeqWithoutIsStrInference([False])))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")))
            return _dafny.Map(coll17_)
        return (_dafny.Map({len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krnlu")))})): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojsoe")))})) | ((iife17_()
        ) | (_dafny.Map({len(_dafny.Map({393: -765})): (_dafny.MultiSet([not(not(False)), True])).cardinality})))

    @staticmethod
    def fm49(p0, globalState):
        return D10_DC32((_dafny.CodePoint('c') if True else _dafny.CodePoint('j')), not(not (True) or (False)))

    @staticmethod
    def fm50(p0, p1, p2, p3, globalState):
        return D6_DC21((_dafny.Set({True, not(True)})).intersection(_dafny.Set({True})))

    @staticmethod
    def fm51(globalState):
        return D11_DC35((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_85_i0_ in range(default__.abs(251))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_86_i1_ in range(default__.abs(38))])), (D1_DC4(False)).cf8)

    @staticmethod
    def fm52(globalState):
        return _dafny.Map({not(True): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_87_i0_ in range(default__.abs(183))])})

    @staticmethod
    def fm53(p0, p1, p2, p3, globalState):
        return _dafny.Map({(D10_DC33(not(False), len(_dafny.SeqWithoutIsStrInference([not(True), False])))).cf54: D3_DC12(D3_DC11(not(not(False)), False))})

    @staticmethod
    def fm54(p0, p1, p2, p3, globalState):
        if False:
            return D15_DC43(_dafny.Map({142: False}))
        elif True:
            def iife20_():
                coll20_ = _dafny.Map()
                compr_20_: int
                for compr_20_ in _dafny.IntegerRange(-229, 503):
                    d_88_v0_: int = compr_20_
                    if ((-229) <= (d_88_v0_)) and ((d_88_v0_) < (503)):
                        coll20_[(d_88_v0_) + ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gslmjy")))))] = True
                return _dafny.Map(coll20_)
            return D15_DC43(iife20_()
)

    @staticmethod
    def fm55(globalState):
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sltsbvpru")), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bslnuaicl"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hl"))), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekjt"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhpp")))])

    @staticmethod
    def fm56(globalState):
        def iife21_():
            coll21_ = _dafny.Map()
            compr_21_: int
            for compr_21_ in (_dafny.SeqWithoutIsStrInference([-338])).Elements:
                d_90_v0_: int = compr_21_
                if (d_90_v0_) in (_dafny.SeqWithoutIsStrInference([-338])):
                    coll21_[(d_90_v0_) * (len(_dafny.Set({_dafny.CodePoint('v'), _dafny.CodePoint('q')})))] = False
            return _dafny.Map(coll21_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([D21_DC65(_dafny.Set({_dafny.SeqWithoutIsStrInference([-225 for d_89_i0_ in range(default__.abs(229))])}))])): True})])) + (_dafny.SeqWithoutIsStrInference([iife21_()
        ]))

    @staticmethod
    def fm57(p0, globalState):
        def iife22_():
            coll22_ = _dafny.Set()
            compr_22_: str
            for compr_22_ in (_dafny.Map({_dafny.CodePoint('c'): _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_91_i0_ in range(default__.abs(891))])): len(_dafny.SeqWithoutIsStrInference([False, False]))})})).keys.Elements:
                d_92_v0_: str = compr_22_
                if (d_92_v0_) in (_dafny.Map({_dafny.CodePoint('c'): _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_91_i0_ in range(default__.abs(891))])): len(_dafny.SeqWithoutIsStrInference([False, False]))})})):
                    coll22_ = coll22_.union(_dafny.Set([d_92_v0_]))
            return _dafny.Set(coll22_)
        return iife22_()
        

    @staticmethod
    def m10(p0, p1, p2, p3, globalState):
        if (p3 if not(p2) else p2):
            d_93_v0_: _dafny.Array
            nw0_ = _dafny.Array(int(0), 20)
            d_93_v0_ = nw0_
            d_94_v1_: _dafny.Seq
            d_94_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brux"))
            index0_ = default__.safeIndex(10, (d_93_v0_).length(0))
            (d_93_v0_)[index0_] = len(d_94_v1_)
            index1_ = default__.safeIndex(10, (d_93_v0_).length(0))
            (d_93_v0_)[index1_] = p1
            d_95_v2_: str
            d_95_v2_ = _dafny.CodePoint('s')
            d_96_v3_: _dafny.Array
            nw1_ = _dafny.Array(None, 1)
            nw1_[int(0)] = p2
            d_96_v3_ = nw1_
            index2_ = default__.safeIndex(891, (d_96_v3_).length(0))
            (d_96_v3_)[index2_] = p3
            d_97_v4_: _dafny.Seq
            d_97_v4_ = _dafny.SeqWithoutIsStrInference([p3])
            index3_ = default__.safeIndex(891, (d_96_v3_).length(0))
            index4_ = default__.safeIndex(10, (d_93_v0_).length(0))
            rhs0_ = p0
            rhs1_ = p3
            rhs2_ = ((d_93_v0_)[default__.safeIndex(10, (d_93_v0_).length(0))]) >= (default__.safeDivisionInt(default__.fm0((d_93_v0_)[default__.safeIndex(10, (d_93_v0_).length(0))], p3, (d_97_v4_)[default__.safeIndex(default__.fm36(globalState), len(d_97_v4_))], globalState), (d_93_v0_)[default__.safeIndex(10, (d_93_v0_).length(0))]))
            rhs3_ = default__.safeDivisionInt((d_93_v0_)[default__.safeIndex(10, (d_93_v0_).length(0))], p1)
            lhs0_ = d_96_v3_
            lhs1_ = default__.safeIndex(891, (d_96_v3_).length(0))
            lhs2_ = globalState
            lhs3_ = d_93_v0_
            lhs4_ = default__.safeIndex(10, (d_93_v0_).length(0))
            d_95_v2_ = rhs0_
            lhs0_[lhs1_] = rhs1_
            lhs2_.f18 = rhs2_
            lhs3_[lhs4_] = rhs3_
            nw2_ = _dafny.Array(int(0), 9)
            (globalState).f17 = nw2_
            d_98_v5_: C2
            nw3_ = C2()
            nw3_.ctor__((d_93_v0_)[default__.safeIndex(10, (d_93_v0_).length(0))], _dafny.Set({d_93_v0_, d_93_v0_}), (d_96_v3_)[default__.safeIndex(891, (d_96_v3_).length(0))], _dafny.SeqWithoutIsStrInference([d_95_v2_ for d_99_i0_ in range(default__.abs(580))]))
            d_98_v5_ = nw3_
            d_100_v6_: _dafny.Map
            d_100_v6_ = _dafny.Map({d_98_v5_: d_96_v3_})
            d_101_v7_: T0
            nw4_ = C1()
            nw4_.ctor__(p2)
            d_101_v7_ = nw4_
            d_102_v8_: C0
            nw5_ = C0()
            nw5_.ctor__(False)
            d_102_v8_ = nw5_
            d_103_v9_: _dafny.Array
            nw6_ = _dafny.Array(None, 28)
            nw6_[int(0)] = d_102_v8_
            nw6_[int(1)] = d_102_v8_
            nw6_[int(2)] = d_102_v8_
            nw6_[int(3)] = d_102_v8_
            nw6_[int(4)] = d_102_v8_
            nw6_[int(5)] = d_102_v8_
            nw6_[int(6)] = d_102_v8_
            nw6_[int(7)] = d_102_v8_
            nw6_[int(8)] = d_102_v8_
            nw6_[int(9)] = d_102_v8_
            nw6_[int(10)] = d_102_v8_
            nw6_[int(11)] = d_102_v8_
            nw6_[int(12)] = d_102_v8_
            nw6_[int(13)] = d_102_v8_
            nw6_[int(14)] = d_102_v8_
            nw6_[int(15)] = d_102_v8_
            nw6_[int(16)] = d_102_v8_
            nw6_[int(17)] = d_102_v8_
            nw6_[int(18)] = d_102_v8_
            nw6_[int(19)] = d_102_v8_
            nw6_[int(20)] = d_102_v8_
            nw6_[int(21)] = d_102_v8_
            nw6_[int(22)] = d_102_v8_
            nw6_[int(23)] = d_102_v8_
            nw6_[int(24)] = d_102_v8_
            nw6_[int(25)] = d_102_v8_
            nw6_[int(26)] = d_102_v8_
            nw6_[int(27)] = d_102_v8_
            d_103_v9_ = nw6_
            d_104_v10_: _dafny.Map
            d_104_v10_ = _dafny.Map({d_103_v9_: d_96_v3_})
            d_105_v11_: _dafny.Map
            d_105_v11_ = _dafny.Map({d_101_v7_: ((d_104_v10_)[d_103_v9_] if (d_103_v9_) in (d_104_v10_) else d_96_v3_)})
            d_100_v6_ = (d_100_v6_).set(d_98_v5_, ((d_105_v11_)[d_101_v7_] if (d_101_v7_) in (d_105_v11_) else d_96_v3_))
            d_106_v12_: _dafny.Map
            d_106_v12_ = _dafny.Map({d_102_v8_.f37: p3})
            d_106_v12_ = (d_106_v12_).set(p2, p3)
        elif True:
            (globalState).f18 = p2
            d_107_v13_: _dafny.Array
            nw7_ = _dafny.Array(False, 10)
            d_107_v13_ = nw7_
            index5_ = default__.safeIndex(560, (d_107_v13_).length(0))
            (d_107_v13_)[index5_] = (True) or (p3)
            d_108_v14_: _dafny.Map
            d_108_v14_ = _dafny.Map({True: p1})
            d_109_v15_: _dafny.Seq
            d_109_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofqjjeh"))
            d_110_v16_: _dafny.Map
            d_110_v16_ = _dafny.Map({p1: 864})
            d_111_v17_: _dafny.Map
            d_111_v17_ = _dafny.Map({default__.safeModuloInt((0) - (len(d_109_v15_)), -368): (p1) in (d_110_v16_)})
            index6_ = default__.safeIndex(560, (d_107_v13_).length(0))
            rhs4_ = ((default__.fm1(default__.fm42(p3, False, globalState), p1, p1, p3, globalState)) not in (d_108_v14_)) == ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oa")))) >= (p1))
            rhs5_ = ((d_111_v17_)[p1] if (p1) in (d_111_v17_) else p2)
            lhs5_ = d_107_v13_
            lhs6_ = default__.safeIndex(560, (d_107_v13_).length(0))
            lhs7_ = globalState
            lhs5_[lhs6_] = rhs4_
            lhs7_.f15 = rhs5_
            d_112_v18_: _dafny.Set
            d_112_v18_ = _dafny.Set({p0, p0, default__.fm42(default__.fm19(p0, globalState), p3, globalState)})
            d_112_v18_ = default__.fm57(not(p2), globalState)
            d_113_v19_: C4
            nw8_ = C4()
            nw8_.ctor__((d_107_v13_)[default__.safeIndex(560, (d_107_v13_).length(0))], p1)
            d_113_v19_ = nw8_
            index7_ = default__.safeIndex(560, (d_107_v13_).length(0))
            (d_107_v13_)[index7_] = (d_113_v19_).f33
        d_114_v20_: _dafny.Set
        d_114_v20_ = _dafny.Set({p1, p1, 906})
        d_115_v21_: _dafny.Seq
        d_115_v21_ = _dafny.SeqWithoutIsStrInference([len(d_114_v20_)])
        d_116_v22_: _dafny.Set
        d_116_v22_ = _dafny.Set({not(p3)})
        d_117_v23_: _dafny.MultiSet
        d_117_v23_ = _dafny.MultiSet([d_116_v22_, _dafny.Set({p2})])
        d_118_v24_: _dafny.Array
        nw9_ = _dafny.Array(None, 12)
        nw9_[int(0)] = (d_115_v21_) <= (d_115_v21_)
        nw9_[int(1)] = False
        nw9_[int(2)] = p2
        nw9_[int(3)] = not (p3) or (p2)
        nw9_[int(4)] = p2
        nw9_[int(5)] = p2
        nw9_[int(6)] = False
        nw9_[int(7)] = (d_117_v23_).ispropersubset(d_117_v23_)
        nw9_[int(8)] = p2
        nw9_[int(9)] = p3
        nw9_[int(10)] = (p3) not in (d_116_v22_)
        nw9_[int(11)] = default__.fm1(p0, p1, 371, p2, globalState)
        d_118_v24_ = nw9_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_118_v24_).length(0)):
            d_119_i1_: int = guard_loop_0_
            if (True) and (((0) <= (d_119_i1_)) and ((d_119_i1_) < ((d_118_v24_).length(0)))):
                (d_118_v24_)[(d_119_i1_)] = not(p3)
        (globalState).f15 = (p3 if p3 else p3)
        d_120_v25_: _dafny.Array
        nw10_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
        d_120_v25_ = nw10_
        d_120_v25_ = d_120_v25_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_120_v25_).length(0)):
            d_121_i2_: int = guard_loop_1_
            if (True) and (((0) <= (d_121_i2_)) and ((d_121_i2_) < ((d_120_v25_).length(0)))):
                (d_120_v25_)[(d_121_i2_)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mp"))
        d_122_v26_: _dafny.Seq
        d_122_v26_ = _dafny.SeqWithoutIsStrInference([p3, p3, p3, False, p2])
        index8_ = default__.safeIndex(920, (d_118_v24_).length(0))
        (d_118_v24_)[index8_] = default__.fm33(p2, len(d_116_v22_), len(d_122_v26_), globalState)
        d_123_v27_: T0
        nw11_ = C6()
        nw11_.ctor__(p2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bydkuf")))
        d_123_v27_ = nw11_
        d_124_v28_: _dafny.Map
        d_124_v28_ = _dafny.Map({p1: d_123_v27_})
        d_125_v29_: _dafny.Map
        d_125_v29_ = _dafny.Map({((d_124_v28_)[p1] if (p1) in (d_124_v28_) else d_123_v27_): p0})
        d_126_v30_: _dafny.Seq
        d_126_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "juplsu"))
        index9_ = default__.safeIndex(920, (d_118_v24_).length(0))
        rhs6_ = p1
        rhs7_ = (((d_125_v29_)[d_123_v27_] if (d_123_v27_) in (d_125_v29_) else p0)) not in (d_126_v30_)
        rhs8_ = p1
        lhs8_ = globalState
        lhs9_ = d_118_v24_
        lhs10_ = default__.safeIndex(920, (d_118_v24_).length(0))
        lhs11_ = globalState
        lhs8_.f21 = rhs6_
        lhs9_[lhs10_] = rhs7_
        lhs11_.f14 = rhs8_

    @staticmethod
    def Main(noArgsParameter__):
        d_127_v0_: bool
        d_127_v0_ = True
        d_128_v1_: _dafny.Set
        d_128_v1_ = _dafny.Set({d_127_v0_, d_127_v0_})
        d_129_v2_: _dafny.Seq
        d_129_v2_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_127_v0_}), _dafny.Set({d_127_v0_}), d_128_v1_, _dafny.Set({d_127_v0_}), d_128_v1_])
        d_130_v3_: int
        d_130_v3_ = -208
        d_131_v4_: _dafny.Set
        d_131_v4_ = _dafny.Set({d_130_v3_})
        d_132_v6_: _dafny.MultiSet
        d_132_v6_ = _dafny.MultiSet([137])
        d_133_v7_: _dafny.Seq
        def iife23_():
            coll23_ = _dafny.Map()
            compr_23_: _dafny.MultiSet
            for compr_23_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_130_v3_]) for d_134_i0_ in range(default__.abs(-496))])).Elements:
                d_135_v5_: _dafny.MultiSet = compr_23_
                if (d_135_v5_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_130_v3_]) for d_134_i0_ in range(default__.abs(-496))])):
                    coll23_[d_135_v5_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r'), _dafny.CodePoint('s')])
            return _dafny.Map(coll23_)
        d_133_v7_ = _dafny.SeqWithoutIsStrInference([len(d_131_v4_), len((iife23_()
        ).set(d_132_v6_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_136_i1_ in range(default__.abs(506))])))])
        d_137_v8_: _dafny.Seq
        d_137_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqigbg"))
        d_138_v9_: _dafny.Array
        nw12_ = _dafny.Array(int(0), 11)
        d_138_v9_ = nw12_
        d_139_v10_: _dafny.Map
        d_139_v10_ = _dafny.Map({False: d_138_v9_})
        d_140_v11_: _dafny.Map
        d_140_v11_ = _dafny.Map({d_130_v3_: (d_137_v8_)[default__.safeIndex(d_130_v3_, len(d_137_v8_))]})
        d_141_v12_: str
        d_141_v12_ = _dafny.CodePoint('l')
        d_142_v13_: _dafny.Map
        d_142_v13_ = _dafny.Map({((d_140_v11_)[d_130_v3_] if (d_130_v3_) in (d_140_v11_) else d_141_v12_): d_127_v0_})
        d_143_globalState_: GlobalState
        nw13_ = GlobalState()
        nw13_.ctor__(True, 724, 130, d_129_v2_, True, _dafny.MultiSet([d_130_v3_]), 288, d_133_v7_, 643, 276, 647, (d_137_v8_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_144_i2_ in range(default__.abs(-875))])), 741, False, 70, True, d_139_v10_, d_138_v9_, False, 723, d_137_v8_, -758, d_142_v13_, 863, True, d_137_v8_, 637, False, True)
        d_143_globalState_ = nw13_
        if not(d_127_v0_):
            d_145_v14_: D0
            d_145_v14_ = D0_DC0((d_127_v0_ if d_127_v0_ else False))
            source4_ = d_145_v14_
            if source4_.is_DC0:
                d_146___mcc_h0_ = source4_.cf0
                d_147_cf0_ = d_146___mcc_h0_
                (d_143_globalState_).f20 = (d_137_v8_) + (d_137_v8_)
                d_148_v15_: C4
                nw14_ = C4()
                nw14_.ctor__(not(d_127_v0_), d_130_v3_)
                d_148_v15_ = nw14_
                d_149_v16_: D6
                d_149_v16_ = D6_DC21(d_128_v1_)
                d_149_v16_ = D6_DC21(d_128_v1_)
                d_150_v17_: _dafny.Array
                nw15_ = _dafny.Array(None, 29)
                d_150_v17_ = nw15_
                d_150_v17_ = d_150_v17_
            elif source4_.is_DC1:
                d_151___mcc_h1_ = source4_.cf1
                d_152___mcc_h2_ = source4_.cf2
                d_153___mcc_h3_ = source4_.cf3
                d_154_cf3_ = d_153___mcc_h3_
                d_155_cf2_ = d_152___mcc_h2_
                d_156_cf1_ = d_151___mcc_h1_
                d_154_cf3_ = default__.safeDivisionInt(d_154_cf3_, d_154_cf3_)
                d_157_v18_: T1
                nw16_ = C6()
                nw16_.ctor__(d_127_v0_, d_137_v8_)
                d_157_v18_ = nw16_
                d_158_v19_: _dafny.MultiSet
                d_158_v19_ = _dafny.MultiSet([d_157_v18_, d_157_v18_, d_157_v18_, d_157_v18_, d_157_v18_])
                d_159_v20_: T0
                nw17_ = C1()
                nw17_.ctor__((_dafny.MultiSet([d_157_v18_])).ispropersubset(d_158_v19_))
                d_159_v20_ = nw17_
                d_160_v21_: _dafny.Seq
                d_160_v21_ = _dafny.SeqWithoutIsStrInference([(d_127_v0_) not in (d_128_v1_), d_127_v0_, (d_157_v18_).f29, False])
                d_161_v22_: _dafny.Map
                d_161_v22_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_141_v12_ for d_162_i3_ in range(default__.abs(714))])): not((d_157_v18_).f29)})
                d_163_v23_: D15
                d_163_v23_ = D15_DC43(d_161_v22_)
                pat_let_tv0_ = d_161_v22_
                d_164_v24_: _dafny.Array
                nw18_ = _dafny.Array(None, 13)
                nw18_[int(0)] = D15_DC43(d_161_v22_)
                nw18_[int(1)] = D15_DC43(d_161_v22_)
                nw18_[int(2)] = d_163_v23_
                nw18_[int(3)] = d_163_v23_
                nw18_[int(4)] = d_163_v23_
                nw18_[int(5)] = default__.fm54(not((d_157_v18_).f29), (D2_DC9(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oegckwcwk")), d_130_v3_, d_127_v0_, d_155_cf2_)).cf15, d_131_v4_, (0) - (d_155_cf2_), d_143_globalState_)
                nw18_[int(6)] = default__.fm54(d_127_v0_, d_130_v3_, d_131_v4_, 627, d_143_globalState_)
                nw18_[int(7)] = d_163_v23_
                nw18_[int(8)] = d_163_v23_
                def iife24_(_pat_let0_0):
                    def iife25_(d_165_dt__update__tmp_h0_):
                        def iife26_(_pat_let1_0):
                            def iife27_(d_166_dt__update_hcf70_h0_):
                                return D15_DC43(d_166_dt__update_hcf70_h0_)
                            return iife27_(_pat_let1_0)
                        return iife26_(pat_let_tv0_)
                    return iife25_(_pat_let0_0)
                nw18_[int(9)] = iife24_(d_163_v23_)
                nw18_[int(10)] = d_163_v23_
                nw18_[int(11)] = d_163_v23_
                nw18_[int(12)] = d_163_v23_
                d_164_v24_ = nw18_
                index10_ = default__.safeIndex(254, (d_164_v24_).length(0))
                (d_164_v24_)[index10_] = d_163_v23_
                index11_ = default__.safeIndex(254, (d_164_v24_).length(0))
                rhs9_ = d_159_v20_
                rhs10_ = (d_160_v21_ if d_127_v0_ else d_160_v21_)
                rhs11_ = ((d_132_v6_ if d_127_v0_ else d_132_v6_)) - (d_132_v6_)
                rhs12_ = d_163_v23_
                rhs13_ = (d_155_cf2_) * (len(d_128_v1_))
                lhs12_ = d_164_v24_
                lhs13_ = default__.safeIndex(254, (d_164_v24_).length(0))
                lhs14_ = d_143_globalState_
                d_159_v20_ = rhs9_
                d_160_v21_ = rhs10_
                d_132_v6_ = rhs11_
                lhs12_[lhs13_] = rhs12_
                lhs14_.f14 = rhs13_
                (d_143_globalState_).f14 = (0) - ((d_155_cf2_) - (d_130_v3_))
                d_167_v25_: _dafny.Array
                nw19_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_167_v25_ = nw19_
                index12_ = default__.safeIndex(595, (d_167_v25_).length(0))
                (d_167_v25_)[index12_] = d_138_v9_
                d_168_v26_: _dafny.Map
                d_168_v26_ = _dafny.Map({d_127_v0_: -824})
                d_169_v27_: _dafny.Map
                d_169_v27_ = _dafny.Map({(d_157_v18_).f29: d_137_v8_})
                index13_ = default__.safeIndex(595, (d_167_v25_).length(0))
                rhs14_ = d_138_v9_
                rhs15_ = (((d_168_v26_)[d_127_v0_] if (d_127_v0_) in (d_168_v26_) else d_130_v3_)) - ((0) - (default__.fm36(d_143_globalState_)))
                rhs16_ = len((d_137_v8_) + (((d_169_v27_)[True] if (True) in (d_169_v27_) else d_157_v18_.f30)))
                lhs15_ = d_167_v25_
                lhs16_ = default__.safeIndex(595, (d_167_v25_).length(0))
                lhs17_ = d_143_globalState_
                lhs15_[lhs16_] = rhs14_
                d_130_v3_ = rhs15_
                lhs17_.f21 = rhs16_
            elif True:
                d_170___mcc_h4_ = source4_.cf4
                d_171___mcc_h5_ = source4_.cf5
                d_172___mcc_h6_ = source4_.cf6
                d_173_cf6_ = d_172___mcc_h6_
                d_174_cf5_ = d_171___mcc_h5_
                d_175_cf4_ = d_170___mcc_h4_
                (d_143_globalState_).f13 = ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdwnsce")))) * (d_130_v3_)) == (len(default__.fm44(d_173_cf6_, d_130_v3_, not(False), d_173_cf6_, d_143_globalState_)))
                index14_ = default__.safeIndex(71, (d_138_v9_).length(0))
                (d_138_v9_)[index14_] = (0) - (d_130_v3_)
                index15_ = default__.safeIndex(71, (d_138_v9_).length(0))
                (d_138_v9_)[index15_] = (0) - ((0) - (default__.safeModuloInt(d_130_v3_, len(d_133_v7_))))
                d_176_v28_: _dafny.Array
                nw20_ = _dafny.Array(int(0), 26)
                d_176_v28_ = nw20_
                d_177_v29_: _dafny.Map
                d_177_v29_ = _dafny.Map({d_176_v28_: False})
                d_177_v29_ = (d_177_v29_) | (d_177_v29_)
                (d_143_globalState_).f1 = d_130_v3_
            d_178_v30_: _dafny.Array
            nw21_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_178_v30_ = nw21_
            d_179_v31_: _dafny.Array
            nw22_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_179_v31_ = nw22_
            index16_ = default__.safeIndex(694, (d_178_v30_).length(0))
            (d_178_v30_)[index16_] = d_179_v31_
            d_180_v32_: _dafny.Array
            def lambda0_(d_181_v3_):
                def lambda1_(d_182_i4_):
                    return _dafny.Map({d_181_v3_: not(False)})

                return lambda1_

            init0_ = lambda0_(d_130_v3_)
            nw23_ = _dafny.Array(None, 17)
            for i0_0_ in range(nw23_.length(0)):
                nw23_[i0_0_] = init0_(i0_0_)
            d_180_v32_ = nw23_
            d_183_v33_: _dafny.Map
            d_183_v33_ = _dafny.Map({d_130_v3_: d_127_v0_})
            index17_ = default__.safeIndex(691, (d_180_v32_).length(0))
            (d_180_v32_)[index17_] = d_183_v33_
            d_184_v34_: _dafny.Map
            d_184_v34_ = _dafny.Map({d_127_v0_: (0) - (d_130_v3_)})
            d_185_v35_: _dafny.Seq
            d_185_v35_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_127_v0_: d_130_v3_})])
            d_186_v36_: _dafny.Seq
            d_186_v36_ = _dafny.SeqWithoutIsStrInference([d_184_v34_, d_184_v34_, (d_185_v35_)[default__.safeIndex(284, len(d_185_v35_))]])
            d_187_v37_: _dafny.Map
            d_187_v37_ = _dafny.Map({True: d_137_v8_})
            index18_ = default__.safeIndex(694, (d_178_v30_).length(0))
            index19_ = default__.safeIndex(691, (d_180_v32_).length(0))
            rhs17_ = d_179_v31_
            rhs18_ = True
            rhs19_ = _dafny.Map({d_130_v3_: d_127_v0_})
            rhs20_ = (d_186_v36_) + (d_186_v36_)
            rhs21_ = ((d_187_v37_)[d_127_v0_] if (d_127_v0_) in (d_187_v37_) else d_137_v8_)
            lhs18_ = d_178_v30_
            lhs19_ = default__.safeIndex(694, (d_178_v30_).length(0))
            lhs20_ = d_143_globalState_
            lhs21_ = d_180_v32_
            lhs22_ = default__.safeIndex(691, (d_180_v32_).length(0))
            lhs23_ = d_143_globalState_
            lhs18_[lhs19_] = rhs17_
            lhs20_.f13 = rhs18_
            lhs21_[lhs22_] = rhs19_
            d_186_v36_ = rhs20_
            lhs23_.f11 = rhs21_
            d_188_v38_: C6
            nw24_ = C6()
            nw24_.ctor__(d_127_v0_, (d_137_v8_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onhi"))))
            d_188_v38_ = nw24_
            nw25_ = C6()
            nw25_.ctor__(not(not(d_127_v0_)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aidqxkrn")))
            d_188_v38_ = nw25_
            (d_143_globalState_).f19 = d_130_v3_
            index20_ = default__.safeIndex(764, (d_138_v9_).length(0))
            (d_138_v9_)[index20_] = len((d_133_v7_) + (d_133_v7_))
            index21_ = default__.safeIndex(764, (d_138_v9_).length(0))
            (d_138_v9_)[index21_] = d_130_v3_
        elif True:
            d_189_v39_: _dafny.Seq
            d_189_v39_ = _dafny.SeqWithoutIsStrInference([d_127_v0_])
            d_190_v40_: C2
            nw26_ = C2()
            nw26_.ctor__(d_130_v3_, _dafny.Set({d_138_v9_}), not((d_189_v39_)[default__.safeIndex(d_130_v3_, len(d_189_v39_))]), d_137_v8_)
            d_190_v40_ = nw26_
            d_191_v41_: _dafny.MultiSet
            d_191_v41_ = _dafny.MultiSet([d_137_v8_, d_137_v8_])
            d_191_v41_ = default__.fm55(d_143_globalState_)
            d_192_v42_: C0
            nw27_ = C0()
            nw27_.ctor__(d_127_v0_)
            d_192_v42_ = nw27_
            index22_ = default__.safeIndex(119, (d_138_v9_).length(0))
            (d_138_v9_)[index22_] = (_dafny.MultiSet([d_127_v0_, d_192_v42_.f37, d_192_v42_.f37])).cardinality
            index23_ = default__.safeIndex(119, (d_138_v9_).length(0))
            rhs22_ = (default__.fm27((d_132_v6_).cardinality, d_192_v42_.f37, d_143_globalState_)) | (_dafny.MultiSet([-916, d_130_v3_]))
            rhs23_ = (len(default__.fm56(d_143_globalState_))) - ((d_190_v40_).f38)
            rhs24_ = (d_190_v40_).f38
            rhs25_ = (True if not(d_192_v42_.f37) else d_127_v0_)
            lhs24_ = d_143_globalState_
            lhs25_ = d_138_v9_
            lhs26_ = default__.safeIndex(119, (d_138_v9_).length(0))
            lhs27_ = d_143_globalState_
            d_132_v6_ = rhs22_
            lhs24_.f2 = rhs23_
            lhs25_[lhs26_] = rhs24_
            lhs27_.f15 = rhs25_
            d_193_v43_: _dafny.Array
            nw28_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
            d_193_v43_ = nw28_
            d_194_v44_: _dafny.Map
            d_194_v44_ = _dafny.Map({d_127_v0_: d_193_v43_})
            d_195_v45_: _dafny.Seq
            d_196_v46_: bool
            out0_: _dafny.Seq
            out1_: bool
            out0_, out1_ = (d_190_v40_).m3((_dafny.Set({d_192_v42_.f37, True})).isdisjoint(d_128_v1_), ((d_194_v44_)[d_127_v0_] if (d_127_v0_) in (d_194_v44_) else d_193_v43_), d_130_v3_, d_143_globalState_)
            d_195_v45_ = out0_
            d_196_v46_ = out1_
        d_197_v47_: D2
        d_197_v47_ = D2_DC8(d_141_v12_, not(default__.fm1(d_141_v12_, 615, len(d_133_v7_), d_127_v0_, d_143_globalState_)))
        d_198_v48_: _dafny.Seq
        d_198_v48_ = _dafny.SeqWithoutIsStrInference([d_197_v47_, d_197_v47_, d_197_v47_, d_197_v47_, d_197_v47_])
        d_199_v49_: _dafny.Array
        def lambda2_(d_200_i5_):
            return True

        init1_ = lambda2_
        nw29_ = _dafny.Array(None, 27)
        for i0_1_ in range(nw29_.length(0)):
            nw29_[i0_1_] = init1_(i0_1_)
        d_199_v49_ = nw29_
        d_201_v50_: D9
        d_201_v50_ = D9_DC28(default__.fm35(d_198_v48_, -251, d_143_globalState_), d_127_v0_, d_130_v3_, (d_130_v3_) == ((D0_DC1(d_199_v49_, len(d_137_v8_), d_130_v3_)).cf3), (d_133_v7_) < (_dafny.SeqWithoutIsStrInference([d_130_v3_ for d_202_i6_ in range(default__.abs(60))])))
        source5_ = d_201_v50_
        if source5_.is_DC28:
            d_203___mcc_h7_ = source5_.cf41
            d_204___mcc_h8_ = source5_.cf42
            d_205___mcc_h9_ = source5_.cf43
            d_206___mcc_h10_ = source5_.cf44
            d_207___mcc_h11_ = source5_.cf45
            d_208_cf45_ = d_207___mcc_h11_
            d_209_cf44_ = d_206___mcc_h10_
            d_210_cf43_ = d_205___mcc_h9_
            d_211_cf42_ = d_204___mcc_h8_
            d_212_cf41_ = d_203___mcc_h7_
            d_213_v51_: _dafny.Array
            nw30_ = _dafny.Array(_dafny.MultiSet({}), 29)
            d_213_v51_ = nw30_
            index24_ = default__.safeIndex(900, (d_213_v51_).length(0))
            (d_213_v51_)[index24_] = d_132_v6_
            index25_ = default__.safeIndex(900, (d_213_v51_).length(0))
            (d_213_v51_)[index25_] = default__.fm27(len((d_137_v8_) + (d_137_v8_)), True, d_143_globalState_)
            (d_143_globalState_).f13 = d_209_cf44_
            d_130_v3_ = d_210_cf43_
            d_214_v52_: T1
            nw31_ = C7()
            nw31_.ctor__(d_208_cf45_, d_208_cf45_, d_208_cf45_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnfnj")))
            d_214_v52_ = nw31_
            d_215_v53_: _dafny.Set
            d_215_v53_ = _dafny.Set({d_214_v52_, d_214_v52_, d_214_v52_, d_214_v52_, d_214_v52_})
            d_216_v54_: _dafny.Set
            d_216_v54_ = _dafny.Set({d_215_v53_, d_215_v53_, d_215_v53_, d_215_v53_, d_215_v53_})
            d_217_v55_: _dafny.Map
            d_217_v55_ = _dafny.Map({False: d_216_v54_})
            d_218_v56_: D19
            d_218_v56_ = D19_DC56(d_216_v54_)
            rhs26_ = (((d_217_v55_)[d_208_cf45_] if (d_208_cf45_) in (d_217_v55_) else _dafny.Set({_dafny.Set({d_214_v52_})})) if d_209_cf44_ else (d_218_v56_).cf99)
            rhs27_ = (len((_dafny.Set({True, d_127_v0_, d_211_cf42_})) - (_dafny.Set({d_211_cf42_, d_127_v0_, d_209_cf44_})))) * (d_210_cf43_)
            lhs28_ = d_143_globalState_
            d_216_v54_ = rhs26_
            lhs28_.f14 = rhs27_
        elif source5_.is_DC29:
            d_219___mcc_h12_ = source5_.cf46
            d_220___mcc_h13_ = source5_.cf47
            d_221___mcc_h14_ = source5_.cf48
            d_222_cf48_ = d_221___mcc_h14_
            d_223_cf47_ = d_220___mcc_h13_
            d_224_cf46_ = d_219___mcc_h12_
            default__.m10(_dafny.CodePoint('x'), d_224_cf46_, (d_222_cf48_ if True else d_222_cf48_), (d_130_v3_) <= (459), d_143_globalState_)
            d_225_v57_: C7
            nw32_ = C7()
            nw32_.ctor__(d_127_v0_, d_127_v0_, d_127_v0_, d_137_v8_)
            d_225_v57_ = nw32_
            d_226_v58_: _dafny.Set
            d_226_v58_ = _dafny.Set({d_225_v57_, d_225_v57_})
            index26_ = default__.safeIndex(151, (d_138_v9_).length(0))
            (d_138_v9_)[index26_] = default__.fm0(len(d_226_v58_), (d_225_v57_).f31, d_222_cf48_, d_143_globalState_)
            index27_ = default__.safeIndex(151, (d_138_v9_).length(0))
            (d_138_v9_)[index27_] = d_130_v3_
            d_227_v59_: _dafny.Map
            d_227_v59_ = _dafny.Map({d_133_v7_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wegckpd")))})
            d_228_v60_: D11
            d_228_v60_ = D11_DC34(d_227_v59_)
            d_228_v60_ = d_228_v60_
            (d_143_globalState_).f21 = (d_138_v9_)[default__.safeIndex(151, (d_138_v9_).length(0))]
        elif True:
            d_229___mcc_h15_ = source5_.cf40
            d_230_cf40_ = d_229___mcc_h15_
            if d_127_v0_:
                d_231_v61_: _dafny.Array
                def lambda3_(d_232_v8_, d_233_v3_):
                    def lambda4_(d_234_i7_):
                        return _dafny.SeqWithoutIsStrInference([d_232_v8_, _dafny.SeqWithoutIsStrInference([(d_232_v8_)[default__.safeIndex(d_233_v3_, len(d_232_v8_))] for d_235_i8_ in range(default__.abs(225))])])

                    return lambda4_

                init2_ = lambda3_(d_137_v8_, d_130_v3_)
                nw33_ = _dafny.Array(None, 18)
                for i0_2_ in range(nw33_.length(0)):
                    nw33_[i0_2_] = init2_(i0_2_)
                d_231_v61_ = nw33_
                d_236_v62_: C6
                nw34_ = C6()
                nw34_.ctor__(d_127_v0_, d_137_v8_)
                d_236_v62_ = nw34_
                d_237_v63_: _dafny.Map
                d_237_v63_ = _dafny.Map({d_236_v62_: _dafny.SeqWithoutIsStrInference([d_141_v12_ for d_238_i9_ in range(default__.abs(628))])})
                d_239_v64_: _dafny.Seq
                d_239_v64_ = _dafny.SeqWithoutIsStrInference([d_236_v62_, d_236_v62_])
                d_240_v65_: _dafny.Seq
                d_240_v65_ = _dafny.SeqWithoutIsStrInference([((d_237_v63_)[(d_239_v64_)[default__.safeIndex(d_130_v3_, len(d_239_v64_))]] if ((d_239_v64_)[default__.safeIndex(d_130_v3_, len(d_239_v64_))]) in (d_237_v63_) else d_137_v8_), d_137_v8_, d_137_v8_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elrvfpxmc"))])
                index28_ = default__.safeIndex(121, (d_231_v61_).length(0))
                (d_231_v61_)[index28_] = d_240_v65_
                d_241_v66_: C0
                nw35_ = C0()
                nw35_.ctor__(d_127_v0_)
                d_241_v66_ = nw35_
                d_242_v67_: _dafny.Seq
                d_242_v67_ = _dafny.SeqWithoutIsStrInference([d_241_v66_, d_241_v66_, d_241_v66_, d_241_v66_, d_241_v66_])
                index29_ = default__.safeIndex(121, (d_231_v61_).length(0))
                rhs28_ = d_130_v3_
                rhs29_ = ((d_137_v8_) + (d_137_v8_)) + (d_137_v8_)
                rhs30_ = d_240_v65_
                rhs31_ = (d_242_v67_)[default__.safeIndex(d_130_v3_, len(d_242_v67_))]
                lhs29_ = d_143_globalState_
                lhs30_ = d_231_v61_
                lhs31_ = default__.safeIndex(121, (d_231_v61_).length(0))
                d_130_v3_ = rhs28_
                lhs29_.f20 = rhs29_
                lhs30_[lhs31_] = rhs30_
                d_241_v66_ = rhs31_
                d_243_v68_: D0
                d_243_v68_ = D0_DC1(d_199_v49_, -974, d_130_v3_)
                d_243_v68_ = d_243_v68_
                d_244_v69_: _dafny.Map
                d_244_v69_ = _dafny.Map({False: len(d_137_v8_)})
                d_245_v70_: C3
                nw36_ = C3()
                nw36_.ctor__((default__.fm9(d_143_globalState_)) <= (d_130_v3_), (579) - (((d_244_v69_)[d_241_v66_.f37] if (d_241_v66_.f37) in (d_244_v69_) else d_130_v3_)), d_241_v66_.f37, (d_137_v8_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogivdrrsy"))))
                d_245_v70_ = nw36_
                (d_143_globalState_).f2 = 750
                d_246_v71_: _dafny.Array
                nw37_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
                d_246_v71_ = nw37_
                index30_ = default__.safeIndex(120, (d_246_v71_).length(0))
                (d_246_v71_)[index30_] = d_137_v8_
                index31_ = default__.safeIndex(120, (d_246_v71_).length(0))
                (d_246_v71_)[index31_] = (d_137_v8_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohllvco"))).set(default__.safeIndex((d_245_v70_).f36, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohllvco")))), d_141_v12_))
            elif True:
                (d_143_globalState_).f7 = d_230_cf40_
                d_247_v72_: _dafny.Map
                d_247_v72_ = _dafny.Map({d_127_v0_: d_141_v12_})
                d_248_v73_: _dafny.Map
                d_248_v73_ = _dafny.Map({len(d_137_v8_): d_127_v0_})
                d_249_v74_: T0
                nw38_ = C6()
                nw38_.ctor__(d_127_v0_, d_137_v8_)
                d_249_v74_ = nw38_
                d_250_v75_: _dafny.Map
                d_250_v75_ = _dafny.Map({d_127_v0_: d_249_v74_})
                d_251_v76_: _dafny.Array
                nw39_ = _dafny.Array(None, 25)
                nw39_[int(0)] = d_141_v12_
                nw39_[int(1)] = d_141_v12_
                nw39_[int(2)] = ((d_247_v72_)[d_127_v0_] if (d_127_v0_) in (d_247_v72_) else d_141_v12_)
                nw39_[int(3)] = d_141_v12_
                nw39_[int(4)] = d_141_v12_
                nw39_[int(5)] = d_141_v12_
                nw39_[int(6)] = d_141_v12_
                nw39_[int(7)] = default__.fm42(d_127_v0_, d_127_v0_, d_143_globalState_)
                nw39_[int(8)] = d_141_v12_
                nw39_[int(9)] = default__.fm42(((d_248_v73_)[len(d_250_v75_)] if (len(d_250_v75_)) in (d_248_v73_) else False), False, d_143_globalState_)
                nw39_[int(10)] = d_141_v12_
                nw39_[int(11)] = d_141_v12_
                nw39_[int(12)] = d_141_v12_
                nw39_[int(13)] = _dafny.CodePoint('x')
                nw39_[int(14)] = (d_141_v12_ if d_127_v0_ else d_141_v12_)
                nw39_[int(15)] = d_141_v12_
                nw39_[int(16)] = d_141_v12_
                nw39_[int(17)] = d_141_v12_
                nw39_[int(18)] = default__.fm42(d_127_v0_, d_127_v0_, d_143_globalState_)
                nw39_[int(19)] = d_141_v12_
                nw39_[int(20)] = d_141_v12_
                nw39_[int(21)] = default__.fm42(d_127_v0_, d_127_v0_, d_143_globalState_)
                nw39_[int(22)] = d_141_v12_
                nw39_[int(23)] = d_141_v12_
                nw39_[int(24)] = d_141_v12_
                d_251_v76_ = nw39_
                index32_ = default__.safeIndex(213, (d_251_v76_).length(0))
                (d_251_v76_)[index32_] = default__.fm42(False, d_127_v0_, d_143_globalState_)
                index33_ = default__.safeIndex(721, (d_199_v49_).length(0))
                (d_199_v49_)[index33_] = d_127_v0_
                index34_ = default__.safeIndex(213, (d_251_v76_).length(0))
                index35_ = default__.safeIndex(721, (d_199_v49_).length(0))
                rhs32_ = (d_130_v3_) != (d_130_v3_)
                rhs33_ = _dafny.CodePoint('j')
                rhs34_ = (0) - (d_130_v3_)
                rhs35_ = d_127_v0_
                rhs36_ = d_127_v0_
                lhs32_ = d_143_globalState_
                lhs33_ = d_251_v76_
                lhs34_ = default__.safeIndex(213, (d_251_v76_).length(0))
                lhs35_ = d_143_globalState_
                lhs36_ = d_199_v49_
                lhs37_ = default__.safeIndex(721, (d_199_v49_).length(0))
                lhs38_ = d_143_globalState_
                lhs32_.f15 = rhs32_
                lhs33_[lhs34_] = rhs33_
                lhs35_.f2 = rhs34_
                lhs36_[lhs37_] = rhs35_
                lhs38_.f15 = rhs36_
                d_252_v77_: _dafny.Array
                nw40_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 7)
                d_252_v77_ = nw40_
                d_253_v78_: _dafny.Map
                d_253_v78_ = _dafny.Map({len((d_230_cf40_).set(default__.safeIndex(d_130_v3_, len(d_230_cf40_)), d_130_v3_)): d_252_v77_})
                d_253_v78_ = (d_253_v78_).set(777, d_252_v77_)
                d_254_v79_: C5
                nw41_ = C5()
                nw41_.ctor__(d_127_v0_, d_137_v8_)
                d_254_v79_ = nw41_
                d_255_v80_: _dafny.MultiSet
                d_255_v80_ = _dafny.MultiSet([(d_199_v49_)[default__.safeIndex(721, (d_199_v49_).length(0))]])
                (d_143_globalState_).f13 = ((d_142_v13_)[d_141_v12_] if (d_141_v12_) in (d_142_v13_) else ((d_255_v80_).set(d_127_v0_, default__.abs(d_130_v3_))).issubset(_dafny.MultiSet([d_127_v0_, d_127_v0_, d_127_v0_])))
            default__.m10(d_141_v12_, d_130_v3_, (-976) >= (default__.fm0(d_130_v3_, d_127_v0_, d_127_v0_, d_143_globalState_)), d_127_v0_, d_143_globalState_)
            d_256_v81_: _dafny.Map
            d_256_v81_ = _dafny.Map({d_127_v0_: 241})
            d_257_v82_: _dafny.Seq
            d_257_v82_ = _dafny.SeqWithoutIsStrInference([d_199_v49_, d_199_v49_, d_199_v49_])
            rhs37_ = (d_257_v82_)[default__.safeIndex(d_130_v3_, len(d_257_v82_))]
            rhs38_ = d_141_v12_
            rhs39_ = d_256_v81_
            d_199_v49_ = rhs37_
            d_141_v12_ = rhs38_
            d_256_v81_ = rhs39_
            default__.m10(d_141_v12_, (d_130_v3_) - ((0) - (d_130_v3_)), (d_130_v3_) not in (d_131_v4_), d_127_v0_, d_143_globalState_)
        (d_143_globalState_).f14 = (d_130_v3_) * ((d_130_v3_) * (302))
        if (d_130_v3_) == (d_130_v3_):
            d_130_v3_ = d_130_v3_
            index36_ = default__.safeIndex(176, (d_138_v9_).length(0))
            (d_138_v9_)[index36_] = d_130_v3_
            d_258_v83_: _dafny.Map
            d_258_v83_ = _dafny.Map({d_130_v3_: 121})
            d_259_v84_: D10
            d_259_v84_ = D10_DC32(_dafny.CodePoint('k'), d_127_v0_)
            pat_let_tv1_ = d_127_v0_
            d_260_v85_: _dafny.Array
            nw42_ = _dafny.Array(None, 12)
            nw42_[int(0)] = d_259_v84_
            nw42_[int(1)] = d_259_v84_
            nw42_[int(2)] = d_259_v84_
            nw42_[int(3)] = D10_DC32(d_141_v12_, d_127_v0_)
            def iife28_(_pat_let2_0):
                def iife29_(d_261_dt__update__tmp_h1_):
                    def iife30_(_pat_let3_0):
                        def iife31_(d_262_dt__update_hcf53_h0_):
                            return D10_DC32((d_261_dt__update__tmp_h1_).cf52, d_262_dt__update_hcf53_h0_)
                        return iife31_(_pat_let3_0)
                    return iife30_(pat_let_tv1_)
                return iife29_(_pat_let2_0)
            nw42_[int(4)] = iife28_(d_259_v84_)
            nw42_[int(5)] = d_259_v84_
            nw42_[int(6)] = D10_DC32(d_141_v12_, d_127_v0_)
            nw42_[int(7)] = d_259_v84_
            nw42_[int(8)] = D10_DC32(d_141_v12_, d_127_v0_)
            nw42_[int(9)] = d_259_v84_
            nw42_[int(10)] = D10_DC32(d_141_v12_, d_127_v0_)
            nw42_[int(11)] = d_259_v84_
            d_260_v85_ = nw42_
            d_263_v86_: C4
            nw43_ = C4()
            nw43_.ctor__(not(d_127_v0_), (D15_DC44(d_130_v3_, d_127_v0_, d_260_v85_, d_138_v9_, (0) - (d_130_v3_))).cf71)
            d_263_v86_ = nw43_
            d_264_v87_: _dafny.Map
            d_264_v87_ = _dafny.Map({d_263_v86_: d_127_v0_})
            index37_ = default__.safeIndex(176, (d_138_v9_).length(0))
            (d_138_v9_)[index37_] = ((0) - (((d_258_v83_)[d_130_v3_] if (d_130_v3_) in (d_258_v83_) else len(d_128_v1_)))) - (len(d_264_v87_))
            d_265_v88_: _dafny.Array
            nw44_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
            d_265_v88_ = nw44_
            d_265_v88_ = d_265_v88_
            d_266_v89_: _dafny.Seq
            d_267_v90_: bool
            out2_: _dafny.Seq
            out3_: bool
            out2_, out3_ = (d_263_v86_).m2(d_127_v0_, d_143_globalState_)
            d_266_v89_ = out2_
            d_267_v90_ = out3_
            d_268_v91_: C6
            nw45_ = C6()
            nw45_.ctor__((d_263_v86_).f33, d_137_v8_)
            d_268_v91_ = nw45_
            d_268_v91_ = d_268_v91_
        elif True:
            (d_143_globalState_).f15 = (d_127_v0_) == (d_127_v0_)
            d_269_v92_: _dafny.Seq
            d_269_v92_ = _dafny.SeqWithoutIsStrInference([(d_130_v3_) != (-811)])
            d_130_v3_ = len(d_269_v92_)
            d_141_v12_ = d_141_v12_
            if d_127_v0_:
                (d_143_globalState_).f1 = ((d_132_v6_)[default__.safeModuloInt(d_130_v3_, len(_dafny.SeqWithoutIsStrInference([d_141_v12_ for d_270_i10_ in range(default__.abs(394))])))] if (default__.safeModuloInt(d_130_v3_, len(_dafny.SeqWithoutIsStrInference([d_141_v12_ for d_271_i10_ in range(default__.abs(394))])))) in (d_132_v6_) else 660)
                d_272_v93_: C5
                nw46_ = C5()
                nw46_.ctor__(d_127_v0_, d_137_v8_)
                d_272_v93_ = nw46_
                d_273_v94_: C7
                nw47_ = C7()
                nw47_.ctor__(d_127_v0_, d_127_v0_, d_127_v0_, d_137_v8_)
                d_273_v94_ = nw47_
                (d_143_globalState_).f7 = d_133_v7_
                rhs40_ = (d_130_v3_) + (((d_132_v6_)[d_130_v3_] if (d_130_v3_) in (d_132_v6_) else d_130_v3_))
                rhs41_ = d_130_v3_
                lhs39_ = d_143_globalState_
                lhs40_ = d_143_globalState_
                lhs39_.f19 = rhs40_
                lhs40_.f14 = rhs41_
            elif True:
                index38_ = default__.safeIndex(858, (d_199_v49_).length(0))
                (d_199_v49_)[index38_] = default__.fm6(d_130_v3_, d_143_globalState_)
                index39_ = default__.safeIndex(858, (d_199_v49_).length(0))
                (d_199_v49_)[index39_] = not (d_127_v0_) or (d_127_v0_)
                d_274_v95_: _dafny.Array
                nw48_ = _dafny.Array(None, 5)
                nw48_[int(0)] = d_127_v0_
                nw48_[int(1)] = (d_199_v49_)[default__.safeIndex(858, (d_199_v49_).length(0))]
                nw48_[int(2)] = (d_199_v49_)[default__.safeIndex(858, (d_199_v49_).length(0))]
                nw48_[int(3)] = True
                nw48_[int(4)] = not(d_127_v0_)
                d_274_v95_ = nw48_
                d_275_v96_: _dafny.Seq
                d_275_v96_ = _dafny.SeqWithoutIsStrInference([d_274_v95_, d_274_v95_])
                (d_143_globalState_).f28 = (d_274_v95_) in (d_275_v96_)
                d_276_v97_: _dafny.Map
                d_276_v97_ = _dafny.Map({825: d_127_v0_})
                default__.m10((d_141_v12_ if default__.fm19(d_141_v12_, d_143_globalState_) else d_141_v12_), d_130_v3_, default__.fm33((d_199_v49_)[default__.safeIndex(858, (d_199_v49_).length(0))], d_130_v3_, len(d_276_v97_), d_143_globalState_), (default__.fm36(d_143_globalState_)) < (d_130_v3_), d_143_globalState_)
                d_277_v98_: _dafny.Set
                d_277_v98_ = _dafny.Set({d_138_v9_})
                d_278_v99_: C2
                nw49_ = C2()
                nw49_.ctor__(d_130_v3_, d_277_v98_, d_127_v0_, d_137_v8_)
                d_278_v99_ = nw49_
                d_279_v100_: _dafny.Set
                d_279_v100_ = _dafny.Set({d_278_v99_, d_278_v99_})
                d_280_v101_: _dafny.Map
                d_280_v101_ = _dafny.Map({(d_137_v8_ if (d_199_v49_)[default__.safeIndex(858, (d_199_v49_).length(0))] else default__.fm20(d_127_v0_, (d_199_v49_)[default__.safeIndex(858, (d_199_v49_).length(0))], d_269_v92_, len(d_279_v100_), d_143_globalState_)): d_133_v7_})
                d_280_v101_ = (d_280_v101_).set((d_137_v8_) + (d_137_v8_), (default__.fm44((d_199_v49_)[default__.safeIndex(858, (d_199_v49_).length(0))], ((d_132_v6_)[(d_278_v99_).f38] if ((d_278_v99_).f38) in (d_132_v6_) else d_130_v3_), d_127_v0_, d_127_v0_, d_143_globalState_)) + (d_133_v7_))
                d_281_v102_: _dafny.Array
                def lambda5_(d_282_i12_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ughfu"))

                init3_ = lambda5_
                nw50_ = _dafny.Array(None, 28)
                for i0_3_ in range(nw50_.length(0)):
                    nw50_[i0_3_] = init3_(i0_3_)
                d_281_v102_ = nw50_
                d_283_v103_: _dafny.Seq
                d_284_v104_: bool
                out4_: _dafny.Seq
                out5_: bool
                out4_, out5_ = (d_278_v99_).m3(default__.fm6(len(_dafny.SeqWithoutIsStrInference([(d_137_v8_)[default__.safeIndex((d_278_v99_).f38, len(d_137_v8_))] for d_285_i11_ in range(default__.abs(502))])), d_143_globalState_), d_281_v102_, (d_130_v3_) + ((d_278_v99_).f38), d_143_globalState_)
                d_283_v103_ = out4_
                d_284_v104_ = out5_
            (d_143_globalState_).f13 = d_127_v0_
        d_286_v105_: _dafny.Array
        nw51_ = _dafny.Array(None, 18)
        d_286_v105_ = nw51_
        d_287_v106_: T1
        nw52_ = C7()
        nw52_.ctor__(d_127_v0_, d_127_v0_, d_127_v0_, d_137_v8_)
        d_287_v106_ = nw52_
        index40_ = default__.safeIndex(47, (d_286_v105_).length(0))
        (d_286_v105_)[index40_] = d_287_v106_
        d_288_v107_: _dafny.Seq
        d_288_v107_ = _dafny.SeqWithoutIsStrInference([d_287_v106_])
        index41_ = default__.safeIndex(47, (d_286_v105_).length(0))
        (d_286_v105_)[index41_] = (d_288_v107_)[default__.safeIndex(default__.fm9(d_143_globalState_), len(d_288_v107_))]
        d_199_v49_ = d_199_v49_
        index42_ = default__.safeIndex(326, (d_199_v49_).length(0))
        (d_199_v49_)[index42_] = d_127_v0_
        d_289_v108_: _dafny.Seq
        d_289_v108_ = _dafny.SeqWithoutIsStrInference([d_127_v0_, (d_287_v106_).f29, d_127_v0_])
        d_290_v109_: _dafny.Map
        d_290_v109_ = _dafny.Map({default__.fm1(d_141_v12_, len(d_289_v108_), d_130_v3_, True, d_143_globalState_): (d_287_v106_).f29})
        index43_ = default__.safeIndex(326, (d_199_v49_).length(0))
        (d_199_v49_)[index43_] = not(((d_290_v109_)[(d_287_v106_).f29] if ((d_287_v106_).f29) in (d_290_v109_) else (d_287_v106_).f29))
        (d_143_globalState_).f28 = (d_287_v106_).f29
        d_291_v110_: _dafny.Array
        nw53_ = _dafny.Array(_dafny.Map({}), 6)
        d_291_v110_ = nw53_
        d_292_v111_: _dafny.Map
        d_292_v111_ = _dafny.Map({True: d_130_v3_})
        d_293_v112_: C5
        nw54_ = C5()
        nw54_.ctor__(((0) - (len((d_292_v111_).set(d_127_v0_, d_130_v3_)))) <= (d_130_v3_), d_137_v8_)
        d_293_v112_ = nw54_
        d_294_v114_: D17
        d_294_v114_ = D17_DC52((d_287_v106_).f29, (d_199_v49_)[default__.safeIndex(326, (d_199_v49_).length(0))], d_130_v3_, d_132_v6_)
        pat_let_tv2_ = d_130_v3_
        pat_let_tv3_ = d_287_v106_
        pat_let_tv4_ = d_130_v3_
        pat_let_tv5_ = d_292_v111_
        pat_let_tv6_ = d_130_v3_
        pat_let_tv7_ = d_289_v108_
        pat_let_tv8_ = d_289_v108_
        def iife32_():
            coll24_ = _dafny.Set()
            compr_24_: int
            for compr_24_ in (d_133_v7_).Elements:
                d_295_v113_: int = compr_24_
                if (d_295_v113_) in (d_133_v7_):
                    coll24_ = coll24_.union(_dafny.Set([(d_295_v113_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jm"))))]))
            return _dafny.Set(coll24_)
        def lambda6_(source6_):
            if source6_.is_DC50:
                d_296___mcc_h16_ = source6_.cf83
                d_297___mcc_h17_ = source6_.cf84
                d_298___mcc_h18_ = source6_.cf85
                d_299_cf85_ = d_298___mcc_h18_
                d_300_cf84_ = d_297___mcc_h17_
                d_301_cf83_ = d_296___mcc_h16_
                return default__.safeDivisionInt(pat_let_tv2_, len(pat_let_tv3_.f30))
            elif source6_.is_DC51:
                d_302___mcc_h19_ = source6_.cf86
                d_303___mcc_h20_ = source6_.cf87
                d_304_cf87_ = d_303___mcc_h20_
                d_305_cf86_ = d_302___mcc_h19_
                return d_305_cf86_
            elif source6_.is_DC52:
                d_306___mcc_h21_ = source6_.cf88
                d_307___mcc_h22_ = source6_.cf89
                d_308___mcc_h23_ = source6_.cf90
                d_309___mcc_h24_ = source6_.cf91
                d_310_cf91_ = d_309___mcc_h24_
                d_311_cf90_ = d_308___mcc_h23_
                d_312_cf89_ = d_307___mcc_h22_
                d_313_cf88_ = d_306___mcc_h21_
                return (pat_let_tv4_) * (len(pat_let_tv5_))
            elif source6_.is_DC49:
                d_314___mcc_h25_ = source6_.cf82
                d_315_cf82_ = d_314___mcc_h25_
                return (0) - (pat_let_tv6_)
            elif True:
                d_316___mcc_h26_ = source6_.cf92
                d_317_cf92_ = d_316___mcc_h26_
                return (_dafny.MultiSet((pat_let_tv7_ if True else pat_let_tv8_))).cardinality

        rhs42_ = d_291_v110_
        rhs43_ = not(((d_131_v4_).intersection(iife32_()
        )).issubset(_dafny.Set({d_130_v3_, len(_dafny.Set({True}))})))
        rhs44_ = (d_287_v106_).f29
        rhs45_ = lambda6_(d_294_v114_)
        rhs46_ = ((d_293_v112_ if (d_287_v106_).f29 else d_293_v112_) if not (d_127_v0_) or ((d_199_v49_)[default__.safeIndex(326, (d_199_v49_).length(0))]) else d_293_v112_)
        lhs41_ = d_143_globalState_
        lhs42_ = d_143_globalState_
        d_291_v110_ = rhs42_
        d_127_v0_ = rhs43_
        lhs41_.f18 = rhs44_
        lhs42_.f19 = rhs45_
        d_293_v112_ = rhs46_
        d_318_v115_: _dafny.Array
        def lambda7_(d_319_v8_):
            def lambda8_(d_320_i13_):
                return d_319_v8_

            return lambda8_

        init4_ = lambda7_(d_137_v8_)
        nw55_ = _dafny.Array(None, 27)
        for i0_4_ in range(nw55_.length(0)):
            nw55_[i0_4_] = init4_(i0_4_)
        d_318_v115_ = nw55_
        d_321_v116_: C7
        nw56_ = C7()
        nw56_.ctor__((d_287_v106_).f29, (d_199_v49_)[default__.safeIndex(326, (d_199_v49_).length(0))], (d_287_v106_).f29, d_287_v106_.f30)
        d_321_v116_ = nw56_
        d_322_v117_: D20
        d_322_v117_ = D20_DC61(_dafny.SeqWithoutIsStrInference([d_321_v116_]))
        d_323_v118_: _dafny.Seq
        d_324_v119_: bool
        out6_: _dafny.Seq
        out7_: bool
        out6_, out7_ = (d_293_v112_).m3((d_287_v106_).f29, d_318_v115_, (len((d_322_v117_).cf111)) - (default__.fm36(d_143_globalState_)), d_143_globalState_)
        d_323_v118_ = out6_
        d_324_v119_ = out7_
        d_325_i14_: int
        d_325_i14_ = 0
        with _dafny.label("0"):
            while ((d_293_v112_).fm4(d_143_globalState_)) >= (len(_dafny.Set({_dafny.Map({(d_321_v116_).f31: d_130_v3_})}))):
                with _dafny.c_label("0"):
                    if (d_325_i14_) >= (100):
                        raise _dafny.Break("0")
                    d_325_i14_ = (d_325_i14_) + (1)
                    (d_143_globalState_).f17 = d_138_v9_
                    d_287_v106_ = d_287_v106_
                    d_326_v120_: _dafny.Array
                    nw57_ = _dafny.Array(None, 23)
                    d_326_v120_ = nw57_
                    d_326_v120_ = d_326_v120_
                    index44_ = default__.safeIndex(626, (d_138_v9_).length(0))
                    (d_138_v9_)[index44_] = (d_130_v3_) * (d_130_v3_)
                    index45_ = default__.safeIndex(626, (d_138_v9_).length(0))
                    rhs47_ = _dafny.CodePoint('k')
                    rhs48_ = d_130_v3_
                    lhs43_ = d_138_v9_
                    lhs44_ = default__.safeIndex(626, (d_138_v9_).length(0))
                    d_141_v12_ = rhs47_
                    lhs43_[lhs44_] = rhs48_
                    pass
            pass
        d_327_v121_: C4
        nw58_ = C4()
        nw58_.ctor__(d_324_v119_, d_130_v3_)
        d_327_v121_ = nw58_
        d_327_v121_ = d_327_v121_
        d_328_v122_: D18
        d_328_v122_ = D18_DC55(d_141_v12_, d_127_v0_, (d_327_v121_).f34, d_138_v9_, d_324_v119_)
        source7_ = d_328_v122_
        if source7_.is_DC55:
            d_329___mcc_h27_ = source7_.cf94
            d_330___mcc_h28_ = source7_.cf95
            d_331___mcc_h29_ = source7_.cf96
            d_332___mcc_h30_ = source7_.cf97
            d_333___mcc_h31_ = source7_.cf98
            d_334_cf98_ = d_333___mcc_h31_
            d_335_cf97_ = d_332___mcc_h30_
            d_336_cf96_ = d_331___mcc_h29_
            d_337_cf95_ = d_330___mcc_h28_
            d_338_cf94_ = d_329___mcc_h27_
            d_127_v0_ = (d_327_v121_).f33
            d_339_v123_: _dafny.Map
            d_339_v123_ = _dafny.Map({d_133_v7_: (d_127_v0_) == ((d_327_v121_).f33)})
            (d_143_globalState_).f15 = ((d_339_v123_)[d_133_v7_] if (d_133_v7_) in (d_339_v123_) else (d_321_v116_).f31)
            d_340_v124_: _dafny.Map
            d_340_v124_ = _dafny.Map({d_199_v49_: d_289_v108_})
            d_341_v125_: D6
            d_341_v125_ = D6_DC21(d_128_v1_)
            d_342_v126_: _dafny.MultiSet
            d_342_v126_ = _dafny.MultiSet([d_341_v125_])
            d_343_v127_: _dafny.Seq
            d_343_v127_ = _dafny.SeqWithoutIsStrInference([D6_DC21(_dafny.Set({d_324_v119_})), d_341_v125_])
            d_344_v128_: _dafny.Seq
            d_344_v128_ = _dafny.SeqWithoutIsStrInference([d_342_v126_, _dafny.MultiSet(d_343_v127_), _dafny.MultiSet([d_341_v125_, d_341_v125_]), d_342_v126_, d_342_v126_])
            rhs49_ = d_340_v124_
            rhs50_ = default__.safeModuloInt((d_327_v121_).f34, len(d_133_v7_))
            rhs51_ = (d_342_v126_) == ((d_344_v128_)[default__.safeIndex(499, len(d_344_v128_))])
            lhs45_ = d_143_globalState_
            lhs46_ = d_143_globalState_
            d_340_v124_ = rhs49_
            lhs45_.f2 = rhs50_
            lhs46_.f15 = rhs51_
            rhs52_ = d_130_v3_
            rhs53_ = 734
            rhs54_ = ((d_327_v121_).f34) >= (-143)
            rhs55_ = d_131_v4_
            lhs47_ = d_143_globalState_
            lhs48_ = d_143_globalState_
            d_130_v3_ = rhs52_
            lhs47_.f2 = rhs53_
            lhs48_.f28 = rhs54_
            d_131_v4_ = rhs55_
        elif True:
            d_345___mcc_h32_ = source7_.cf93
            d_346_cf93_ = d_345___mcc_h32_
            d_347_v129_: D0
            d_347_v129_ = D0_DC0((d_287_v106_).f29)
            rhs56_ = d_347_v129_
            rhs57_ = d_199_v49_
            d_347_v129_ = rhs56_
            d_199_v49_ = rhs57_
            d_348_v130_: _dafny.Array
            nw59_ = _dafny.Array(_dafny.Array(None, 0), 4)
            d_348_v130_ = nw59_
            index46_ = default__.safeIndex(677, (d_348_v130_).length(0))
            (d_348_v130_)[index46_] = d_199_v49_
            index47_ = default__.safeIndex(677, (d_348_v130_).length(0))
            (d_348_v130_)[index47_] = d_199_v49_
            d_349_v131_: _dafny.Array
            def lambda9_(d_350_v7_):
                def lambda10_(d_351_i15_):
                    return d_350_v7_

                return lambda10_

            init5_ = lambda9_(d_133_v7_)
            nw60_ = _dafny.Array(None, 3)
            for i0_5_ in range(nw60_.length(0)):
                nw60_[i0_5_] = init5_(i0_5_)
            d_349_v131_ = nw60_
            index48_ = default__.safeIndex(856, (d_349_v131_).length(0))
            (d_349_v131_)[index48_] = d_133_v7_
            index49_ = default__.safeIndex(856, (d_349_v131_).length(0))
            (d_349_v131_)[index49_] = d_133_v7_
            d_292_v111_ = (d_292_v111_).set((d_321_v116_).f32, (d_327_v121_).f34)
        d_352_i16_: int
        d_352_i16_ = 0
        with _dafny.label("1"):
            while not ((d_321_v116_).f31) or ((d_199_v49_)[default__.safeIndex(326, (d_199_v49_).length(0))]):
                with _dafny.c_label("1"):
                    if (d_352_i16_) >= (100):
                        raise _dafny.Break("1")
                    d_352_i16_ = (d_352_i16_) + (1)
                    d_353_v132_: _dafny.Map
                    d_353_v132_ = _dafny.Map({len(d_289_v108_): d_289_v108_})
                    d_354_v133_: _dafny.Map
                    d_354_v133_ = _dafny.Map({d_132_v6_: (d_353_v132_) | (d_353_v132_)})
                    d_354_v133_ = (d_354_v133_).set((d_132_v6_).set((d_327_v121_).f34, default__.abs(d_130_v3_)), d_353_v132_)
                    d_355_v134_: D0
                    d_355_v134_ = D0_DC1(d_199_v49_, (d_327_v121_).f34, len(d_287_v106_.f30))
                    d_199_v49_ = ((d_355_v134_ if (d_287_v106_).f29 else d_355_v134_)).cf1
                    d_130_v3_ = (0) - (((d_327_v121_).f34) * (len(d_128_v1_)))
                    (d_143_globalState_).f19 = d_130_v3_
                    pass
            pass
        (d_143_globalState_).f2 = (0) - ((d_327_v121_).f34)
        d_356_v135_: _dafny.MultiSet
        d_356_v135_ = _dafny.MultiSet([d_137_v8_, _dafny.SeqWithoutIsStrInference([d_141_v12_ for d_357_i17_ in range(default__.abs(-956))])])
        d_358_v136_: bool
        d_359_v137_: bool
        out8_: bool
        out9_: bool
        out8_, out9_ = (d_327_v121_).m6(default__.fm10((d_327_v121_).f34, (d_356_v135_).cardinality, d_143_globalState_), d_323_v118_, d_137_v8_, D0_DC0(not(not(d_127_v0_))), d_143_globalState_)
        d_358_v136_ = out8_
        d_359_v137_ = out9_
        _dafny.print(_dafny.string_of(d_127_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_128_v1_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v2_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True}), _dafny.Set({True}), _dafny.Set({True}), _dafny.Set({True}), _dafny.Set({True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_130_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v4_) == (_dafny.Set({-208}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v6_) == (_dafny.MultiSet([0, -1, 10, 9, -916, -208]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_133_v7_) == (_dafny.SeqWithoutIsStrInference([1, 2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_137_v8_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v9_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v9_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v9_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_139_v10_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_140_v11_) == (_dafny.Map({-208: _dafny.CodePoint('r')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_141_v12_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_v13_) == (_dafny.Map({_dafny.CodePoint('r'): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_.f3) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True}), _dafny.Set({True}), _dafny.Set({True}), _dafny.Set({True}), _dafny.Set({True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_).f5) == (_dafny.MultiSet([-208]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_.f7) == (_dafny.SeqWithoutIsStrInference([1, 2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_143_globalState_.f11).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_globalState_.f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_143_globalState_.f16)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_.f17)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_.f17)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_.f17)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_globalState_.f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_globalState_.f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_143_globalState_.f20).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_globalState_.f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_).f22) == (_dafny.Map({_dafny.CodePoint('r'): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_143_globalState_).f25).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_globalState_.f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_197_v47_).cf12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_197_v47_).cf13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_198_v48_) == (_dafny.SeqWithoutIsStrInference([D2_DC8(_dafny.CodePoint('l'), True), D2_DC8(_dafny.CodePoint('l'), True), D2_DC8(_dafny.CodePoint('l'), True), D2_DC8(_dafny.CodePoint('l'), True), D2_DC8(_dafny.CodePoint('l'), True)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v49_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_201_v50_).cf41) == (_dafny.SeqWithoutIsStrInference([True, True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_201_v50_).cf42))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_201_v50_).cf43))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_201_v50_).cf44))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_201_v50_).cf45))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_286_v105_)[11]).f29))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_286_v105_)[11].f30).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_287_v106_).f29))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_287_v106_.f30).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_288_v107_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_289_v108_) == (_dafny.SeqWithoutIsStrInference([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v109_) == (_dafny.Map({False: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_292_v111_) == (_dafny.Map({True: -208}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v114_).cf88))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v114_).cf89))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v114_).cf90))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_294_v114_).cf91) == (_dafny.MultiSet([0, -1, 10, 9, -916, -208]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[18]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[19]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[20]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[21]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[22]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[23]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[24]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[25]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_318_v115_)[26]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_321_v116_).f31))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_321_v116_).f32))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_322_v117_).cf111)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_323_v118_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_324_v119_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_325_i14_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v121_).f33))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v121_).f34))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_328_v122_).cf94))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_328_v122_).cf95))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_328_v122_).cf96))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_328_v122_).cf97)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_328_v122_).cf97)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_328_v122_).cf97)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_328_v122_).cf98))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_352_i16_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_356_v135_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqigbg")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k')])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_358_v136_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_359_v137_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any), ('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)}, {self.cf5.VerbatimString(True)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC4(D1, NamedTuple('DC4', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(False)
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

class D2_DC7(D2, NamedTuple('DC7', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf14', Any), ('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({self.cf14.VerbatimString(True)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10
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
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)

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

class D3_DC12(D3, NamedTuple('DC12', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC14(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

class D4_DC14(D4, NamedTuple('DC14', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
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
        return lambda: D5_DC17(_dafny.Set({}))
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
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D5_DC20)

class D5_DC17(D5, NamedTuple('DC17', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC18(D5, NamedTuple('DC18', [])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18)
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC19(D5, NamedTuple('DC19', [('cf28', Any), ('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC19({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC19) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC20(D5, NamedTuple('DC20', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC20({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC20) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC22(int(0), int(0), False)
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

class D6_DC22(D6, NamedTuple('DC22', [('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC22({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC22) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC23(D6, NamedTuple('DC23', [('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC23({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC23) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC21(D6, NamedTuple('DC21', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC21({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC21) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D7_DC24)

class D7_DC24(D7, NamedTuple('DC24', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC24({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC24) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC26()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D8_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)

class D8_DC26(D8, NamedTuple('DC26', [])):
    def __dafnystr__(self) -> str:
        return f'D8.DC26'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC26)
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC28(_dafny.Seq({}), False, int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D9_DC28)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D9_DC29)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)

class D9_DC28(D9, NamedTuple('DC28', [('cf41', Any), ('cf42', Any), ('cf43', Any), ('cf44', Any), ('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC28({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC28) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC29(D9, NamedTuple('DC29', [('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC29({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC29) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC31(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D10_DC31)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D10_DC32)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D10_DC33)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D10_DC30)

class D10_DC31(D10, NamedTuple('DC31', [('cf50', Any), ('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC31({_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC31) and self.cf50 == __o.cf50 and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC32(D10, NamedTuple('DC32', [('cf52', Any), ('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC32({_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC32) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC33(D10, NamedTuple('DC33', [('cf54', Any), ('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC33({_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC33) and self.cf54 == __o.cf54 and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC30(D10, NamedTuple('DC30', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC30({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC30) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC35(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D11_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D11_DC36)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D11_DC37)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D11_DC34)

class D11_DC35(D11, NamedTuple('DC35', [('cf57', Any), ('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC35({self.cf57.VerbatimString(True)}, {_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC35) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC36(D11, NamedTuple('DC36', [('cf59', Any), ('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC36({_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC36) and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC37(D11, NamedTuple('DC37', [('cf62', Any), ('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC37({_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC37) and self.cf62 == __o.cf62 and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC34(D11, NamedTuple('DC34', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC34({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC34) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC39(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D12_DC39)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D12_DC38)

class D12_DC39(D12, NamedTuple('DC39', [('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC39({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC39) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC38(D12, NamedTuple('DC38', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC38({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC38) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC41()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D13_DC41)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D13_DC40)

class D13_DC41(D13, NamedTuple('DC41', [])):
    def __dafnystr__(self) -> str:
        return f'D13.DC41'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC41)
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC40(D13, NamedTuple('DC40', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC40({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC40) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D14_DC42)

class D14_DC42(D14, NamedTuple('DC42', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC42({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC42) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC44(int(0), False, _dafny.Array(None, 0), _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D15_DC44)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D15_DC45)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D15_DC46)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)

class D15_DC44(D15, NamedTuple('DC44', [('cf71', Any), ('cf72', Any), ('cf73', Any), ('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC44({_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC44) and self.cf71 == __o.cf71 and self.cf72 == __o.cf72 and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC45(D15, NamedTuple('DC45', [('cf76', Any), ('cf77', Any), ('cf78', Any), ('cf79', Any), ('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC45({_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC45) and self.cf76 == __o.cf76 and self.cf77 == __o.cf77 and self.cf78 == __o.cf78 and self.cf79 == __o.cf79 and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC46(D15, NamedTuple('DC46', [])):
    def __dafnystr__(self) -> str:
        return f'D15.DC46'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC46)
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC43(D15, NamedTuple('DC43', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC48()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D16_DC48)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D16_DC47)

class D16_DC48(D16, NamedTuple('DC48', [])):
    def __dafnystr__(self) -> str:
        return f'D16.DC48'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC48)
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC47(D16, NamedTuple('DC47', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC47({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC47) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC50(False, False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D17_DC50)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D17_DC51)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D17_DC52)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D17_DC49)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D17_DC53)

class D17_DC50(D17, NamedTuple('DC50', [('cf83', Any), ('cf84', Any), ('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC50({_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC50) and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC51(D17, NamedTuple('DC51', [('cf86', Any), ('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC51({_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC51) and self.cf86 == __o.cf86 and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC52(D17, NamedTuple('DC52', [('cf88', Any), ('cf89', Any), ('cf90', Any), ('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC52({_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC52) and self.cf88 == __o.cf88 and self.cf89 == __o.cf89 and self.cf90 == __o.cf90 and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC49(D17, NamedTuple('DC49', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC49({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC49) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC53(D17, NamedTuple('DC53', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC53({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC53) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC55(_dafny.CodePoint('D'), False, int(0), _dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D18_DC55)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D18_DC54)

class D18_DC55(D18, NamedTuple('DC55', [('cf94', Any), ('cf95', Any), ('cf96', Any), ('cf97', Any), ('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC55({_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)}, {_dafny.string_of(self.cf97)}, {_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC55) and self.cf94 == __o.cf94 and self.cf95 == __o.cf95 and self.cf96 == __o.cf96 and self.cf97 == __o.cf97 and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC54(D18, NamedTuple('DC54', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC54({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC54) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC57(_dafny.Array(None, 0), int(0), None, None, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D19_DC57)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D19_DC58)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D19_DC59)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D19_DC56)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D19_DC60)

class D19_DC57(D19, NamedTuple('DC57', [('cf100', Any), ('cf101', Any), ('cf102', Any), ('cf103', Any), ('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC57({_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)}, {_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)}, {_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC57) and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102 and self.cf103 == __o.cf103 and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC58(D19, NamedTuple('DC58', [('cf105', Any), ('cf106', Any), ('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC58({_dafny.string_of(self.cf105)}, {_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC58) and self.cf105 == __o.cf105 and self.cf106 == __o.cf106 and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC59(D19, NamedTuple('DC59', [('cf108', Any), ('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC59({_dafny.string_of(self.cf108)}, {_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC59) and self.cf108 == __o.cf108 and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC56(D19, NamedTuple('DC56', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC56({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC56) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC60(D19, NamedTuple('DC60', [('cf110', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC60({_dafny.string_of(self.cf110)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC60) and self.cf110 == __o.cf110
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC62()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D20_DC62)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D20_DC63)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D20_DC61)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D20_DC64)

class D20_DC62(D20, NamedTuple('DC62', [])):
    def __dafnystr__(self) -> str:
        return f'D20.DC62'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC62)
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC63(D20, NamedTuple('DC63', [('cf112', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC63({_dafny.string_of(self.cf112)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC63) and self.cf112 == __o.cf112
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC61(D20, NamedTuple('DC61', [('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC61({_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC61) and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC64(D20, NamedTuple('DC64', [('cf113', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC64({_dafny.string_of(self.cf113)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC64) and self.cf113 == __o.cf113
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC66(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D21_DC66)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D21_DC67)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D21_DC65)

class D21_DC66(D21, NamedTuple('DC66', [('cf115', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC66({_dafny.string_of(self.cf115)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC66) and self.cf115 == __o.cf115
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC67(D21, NamedTuple('DC67', [('cf116', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC67({_dafny.string_of(self.cf116)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC67) and self.cf116 == __o.cf116
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC65(D21, NamedTuple('DC65', [('cf114', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC65({_dafny.string_of(self.cf114)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC65) and self.cf114 == __o.cf114
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m2(self, p0, globalState):
        pass

    def m3(self, p0, p1, p2, globalState):
        pass


class T1:
    pass
    @property
    def f30(self):
        return self._f30
    @f30.setter
    def f30(self, value):
        self._f30 = value

class GlobalState:
    def  __init__(self):
        self.f1: int = int(0)
        self.f2: int = int(0)
        self.f3: _dafny.Seq = _dafny.Seq({})
        self.f7: _dafny.Seq = _dafny.Seq({})
        self.f11: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f13: bool = False
        self.f14: int = int(0)
        self.f15: bool = False
        self.f16: _dafny.Map = _dafny.Map({})
        self.f17: _dafny.Array = _dafny.Array(None, 0)
        self.f18: bool = False
        self.f19: int = int(0)
        self.f20: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f21: int = int(0)
        self.f28: bool = False
        self._f0: bool = False
        self._f4: bool = False
        self._f5: _dafny.MultiSet = _dafny.MultiSet({})
        self._f6: int = int(0)
        self._f8: int = int(0)
        self._f9: int = int(0)
        self._f10: int = int(0)
        self._f12: int = int(0)
        self._f22: _dafny.Map = _dafny.Map({})
        self._f23: int = int(0)
        self._f24: bool = False
        self._f25: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f26: int = int(0)
        self._f27: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27, f28):
        (self)._f0 = f0
        (self).f1 = f1
        (self).f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self).f14 = f14
        (self).f15 = f15
        (self).f16 = f16
        (self).f17 = f17
        (self).f18 = f18
        (self).f19 = f19
        (self).f20 = f20
        (self).f21 = f21
        (self)._f22 = f22
        (self)._f23 = f23
        (self)._f24 = f24
        (self)._f25 = f25
        (self)._f26 = f26
        (self)._f27 = f27
        (self).f28 = f28

    @property
    def f0(self):
        return self._f0
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
    def f10(self):
        return self._f10
    @property
    def f12(self):
        return self._f12
    @property
    def f22(self):
        return self._f22
    @property
    def f23(self):
        return self._f23
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27

class C0:
    def  __init__(self):
        self.f37: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f37):
        (self).f37 = f37

    def fm11(self, p0, globalState):
        return self.f37

    def fm12(self, globalState):
        return (-345) - ((696 if self.f37 else 343))


class C1(T0):
    def  __init__(self):
        self._f40: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f40):
        (self)._f40 = f40

    def fm30(self, p0, globalState):
        return (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovq"))), (0) - (-114), len(_dafny.Map({(self).f40: -638}))])).intersection(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True, not((self).f40)])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p'), _dafny.CodePoint('n'), _dafny.CodePoint('m'), _dafny.CodePoint('d'), _dafny.CodePoint('r')])), 141]))

    def m2(self, p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_360_v0_: _dafny.Set
        d_360_v0_ = _dafny.Set({p0})
        d_361_v1_: D6
        d_361_v1_ = D6_DC21(d_360_v0_)
        d_362_v2_: _dafny.Seq
        d_362_v2_ = _dafny.SeqWithoutIsStrInference([d_361_v1_])
        d_363_v3_: D8
        d_363_v3_ = D8_DC25(d_362_v2_)
        d_364_v4_: int
        d_364_v4_ = 651
        d_365_v5_: _dafny.MultiSet
        d_365_v5_ = _dafny.MultiSet([(d_363_v3_).cf39, default__.fm31(d_364_v4_, 241, globalState), _dafny.SeqWithoutIsStrInference([d_361_v1_ for d_366_i0_ in range(default__.abs(660))]), d_362_v2_])
        d_367_v6_: _dafny.Map
        d_367_v6_ = _dafny.Map({d_364_v4_: _dafny.Set({p0, (self).f40})})
        (globalState).f18 = not((d_365_v5_).isdisjoint(_dafny.MultiSet([(d_362_v2_).set(default__.safeIndex(140, len(d_362_v2_)), D6_DC21(((d_367_v6_)[d_364_v4_] if (d_364_v4_) in (d_367_v6_) else _dafny.Set({(self).f40}))))])))
        d_368_v7_: _dafny.Map
        d_368_v7_ = _dafny.Map({(self).f40: p0})
        d_369_v8_: _dafny.Seq
        d_369_v8_ = _dafny.SeqWithoutIsStrInference([d_364_v4_])
        d_370_v9_: _dafny.Seq
        d_370_v9_ = _dafny.SeqWithoutIsStrInference([(d_368_v7_) | (d_368_v7_), default__.fm32(p0, default__.fm33(True, -164, d_364_v4_, globalState), d_369_v8_, globalState), d_368_v7_])
        d_371_v10_: _dafny.Seq
        d_371_v10_ = _dafny.SeqWithoutIsStrInference([d_370_v9_, (d_370_v9_) + (d_370_v9_)])
        d_370_v9_ = (d_371_v10_)[default__.safeIndex((d_364_v4_) + (d_364_v4_), len(d_371_v10_))]
        d_372_v11_: _dafny.MultiSet
        d_372_v11_ = _dafny.MultiSet([_dafny.CodePoint('j')])
        d_373_v12_: _dafny.Seq
        d_373_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xo"))
        d_374_v13_: _dafny.Array
        nw61_ = _dafny.Array(None, 17)
        nw61_[int(0)] = (d_372_v11_).isdisjoint(d_372_v11_)
        nw61_[int(1)] = not((self).f40)
        nw61_[int(2)] = p0
        nw61_[int(3)] = p0
        nw61_[int(4)] = (self).f40
        nw61_[int(5)] = ((self).f40 if (self).f40 else (self).f40)
        nw61_[int(6)] = p0
        nw61_[int(7)] = p0
        nw61_[int(8)] = (self).f40
        nw61_[int(9)] = (self).f40
        nw61_[int(10)] = default__.fm33(p0, 364, -247, globalState)
        nw61_[int(11)] = p0
        nw61_[int(12)] = (self).f40
        nw61_[int(13)] = (len(_dafny.SeqWithoutIsStrInference([286]))) < (len(d_373_v12_))
        nw61_[int(14)] = (self).f40
        nw61_[int(15)] = False
        nw61_[int(16)] = (p0) or (p0)
        d_374_v13_ = nw61_
        index50_ = default__.safeIndex(950, (d_374_v13_).length(0))
        (d_374_v13_)[index50_] = (self).f40
        index51_ = default__.safeIndex(950, (d_374_v13_).length(0))
        (d_374_v13_)[index51_] = (d_373_v12_) <= (d_373_v12_)
        d_375_v14_: _dafny.Array
        nw62_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
        d_375_v14_ = nw62_
        index52_ = default__.safeIndex(705, (d_375_v14_).length(0))
        (d_375_v14_)[index52_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_376_i1_ in range(default__.abs(-162))])
        index53_ = default__.safeIndex(705, (d_375_v14_).length(0))
        (d_375_v14_)[index53_] = default__.fm34(d_364_v4_, d_364_v4_, p0, p0, globalState)
        d_377_v15_: _dafny.Seq
        d_377_v15_ = _dafny.SeqWithoutIsStrInference([(d_374_v13_)[default__.safeIndex(950, (d_374_v13_).length(0))], ((self).f40) or (not((self).f40)), ((d_374_v13_)[default__.safeIndex(950, (d_374_v13_).length(0))]) and ((self).f40), True, (self).f40])
        d_378_v16_: D4
        d_378_v16_ = D4_DC15(d_364_v4_, d_364_v4_)
        d_379_v17_: str
        d_379_v17_ = _dafny.CodePoint('i')
        d_380_v18_: D2
        d_380_v18_ = D2_DC8(d_379_v17_, (d_374_v13_)[default__.safeIndex(950, (d_374_v13_).length(0))])
        d_381_v19_: _dafny.Seq
        d_381_v19_ = _dafny.SeqWithoutIsStrInference([D2_DC8(d_379_v17_, p0), d_380_v18_, d_380_v18_])
        rhs58_ = ((d_368_v7_)[(d_374_v13_)[default__.safeIndex(950, (d_374_v13_).length(0))]] if ((d_374_v13_)[default__.safeIndex(950, (d_374_v13_).length(0))]) in (d_368_v7_) else True)
        rhs59_ = d_364_v4_
        rhs60_ = default__.fm35((d_381_v19_) + (d_381_v19_), d_364_v4_, globalState)
        rhs61_ = d_378_v16_
        lhs49_ = globalState
        lhs50_ = globalState
        lhs49_.f28 = rhs58_
        lhs50_.f19 = rhs59_
        d_377_v15_ = rhs60_
        d_378_v16_ = rhs61_
        d_382_v20_: C0
        nw63_ = C0()
        nw63_.ctor__((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrjunsmd"))) == (_dafny.SeqWithoutIsStrInference([d_379_v17_ for d_383_i2_ in range(default__.abs(106))])))
        d_382_v20_ = nw63_
        d_384_v21_: _dafny.Seq
        d_384_v21_ = _dafny.SeqWithoutIsStrInference([d_377_v15_])
        r0 = d_384_v21_
        r1 = p0
        return r0, r1

    def m3(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        d_385_v0_: _dafny.MultiSet
        d_385_v0_ = _dafny.MultiSet([(0) - (p2)])
        d_386_v1_: _dafny.Seq
        d_386_v1_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2])
        d_387_v2_: _dafny.Map
        d_387_v2_ = _dafny.Map({(d_385_v0_).intersection(_dafny.MultiSet([p2])): (d_386_v1_) + (d_386_v1_)})
        d_388_v3_: _dafny.Array
        def lambda11_(d_389_p0_):
            def lambda12_(d_390_i0_):
                return d_389_p0_

            return lambda12_

        init6_ = lambda11_(p0)
        nw64_ = _dafny.Array(None, 26)
        for i0_6_ in range(nw64_.length(0)):
            nw64_[i0_6_] = init6_(i0_6_)
        d_388_v3_ = nw64_
        index54_ = default__.safeIndex(409, (d_388_v3_).length(0))
        (d_388_v3_)[index54_] = not((self).f40)
        d_391_v4_: _dafny.Map
        d_391_v4_ = _dafny.Map({(self).f40: p0})
        d_392_v5_: D9
        d_392_v5_ = D9_DC27(_dafny.SeqWithoutIsStrInference([102 for d_393_i1_ in range(default__.abs(368))]))
        index55_ = default__.safeIndex(409, (d_388_v3_).length(0))
        rhs62_ = (self).f40
        rhs63_ = p0
        rhs64_ = ((d_387_v2_).set(_dafny.MultiSet([len(d_391_v4_), p2, p2, p2, 478]), (d_392_v5_).cf40)) | (d_387_v2_)
        rhs65_ = not (not(p0)) or (p0)
        lhs51_ = globalState
        lhs52_ = globalState
        lhs53_ = d_388_v3_
        lhs54_ = default__.safeIndex(409, (d_388_v3_).length(0))
        lhs51_.f18 = rhs62_
        lhs52_.f18 = rhs63_
        d_387_v2_ = rhs64_
        lhs53_[lhs54_] = rhs65_
        if not ((d_388_v3_)[default__.safeIndex(409, (d_388_v3_).length(0))]) or (not((self).f40)):
            d_394_v6_: D4
            d_394_v6_ = D4_DC14((0) - (p2))
            pat_let_tv9_ = p2
            def iife33_(_pat_let4_0):
                def iife34_(d_395_dt__update__tmp_h0_):
                    def iife35_(_pat_let5_0):
                        def iife36_(d_396_dt__update_hcf23_h0_):
                            return D4_DC14(d_396_dt__update_hcf23_h0_)
                        return iife36_(_pat_let5_0)
                    return iife35_(pat_let_tv9_)
                return iife34_(_pat_let4_0)
            d_394_v6_ = iife33_(d_394_v6_)
            d_397_v7_: C0
            nw65_ = C0()
            nw65_.ctor__((self).f40)
            d_397_v7_ = nw65_
            index56_ = default__.safeIndex(409, (d_388_v3_).length(0))
            (d_388_v3_)[index56_] = not(d_397_v7_.f37)
            d_398_v8_: D0
            d_398_v8_ = D0_DC1(d_388_v3_, p2, p2)
            d_399_v9_: _dafny.Array
            nw66_ = _dafny.Array(None, 9)
            nw66_[int(0)] = (d_398_v8_).cf1
            nw66_[int(1)] = d_388_v3_
            nw66_[int(2)] = (d_398_v8_).cf1
            nw66_[int(3)] = d_388_v3_
            nw66_[int(4)] = d_388_v3_
            nw66_[int(5)] = (d_388_v3_ if d_397_v7_.f37 else d_388_v3_)
            nw66_[int(6)] = d_388_v3_
            nw66_[int(7)] = d_388_v3_
            nw66_[int(8)] = d_388_v3_
            d_399_v9_ = nw66_
            index57_ = default__.safeIndex(948, (d_399_v9_).length(0))
            (d_399_v9_)[index57_] = d_388_v3_
            d_400_v10_: _dafny.Seq
            d_400_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eakffkqhw"))
            d_401_v11_: _dafny.Seq
            d_401_v11_ = _dafny.SeqWithoutIsStrInference([d_400_v10_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkbchhlud")), (d_400_v10_) + (d_400_v10_)])
            index58_ = default__.safeIndex(948, (d_399_v9_).length(0))
            rhs66_ = (d_388_v3_ if (self).f40 else d_388_v3_)
            rhs67_ = d_400_v10_
            rhs68_ = len(d_401_v11_)
            rhs69_ = d_397_v7_.f37
            rhs70_ = default__.safeModuloInt((101) * (p2), default__.safeDivisionInt((d_397_v7_).fm12(globalState), p2))
            lhs55_ = d_399_v9_
            lhs56_ = default__.safeIndex(948, (d_399_v9_).length(0))
            lhs57_ = globalState
            lhs58_ = globalState
            lhs59_ = globalState
            lhs60_ = globalState
            lhs55_[lhs56_] = rhs66_
            lhs57_.f20 = rhs67_
            lhs58_.f19 = rhs68_
            lhs59_.f28 = rhs69_
            lhs60_.f21 = rhs70_
            d_402_v12_: _dafny.Map
            d_402_v12_ = _dafny.Map({d_386_v1_: p0})
            d_403_v14_: _dafny.MultiSet
            d_403_v14_ = _dafny.MultiSet([d_386_v1_, _dafny.SeqWithoutIsStrInference([p2])])
            d_404_v15_: C0
            nw67_ = C0()
            def iife37_():
                coll25_ = _dafny.Map()
                compr_25_: _dafny.Seq
                for compr_25_ in (d_403_v14_).Elements:
                    d_405_v13_: _dafny.Seq = compr_25_
                    if (d_405_v13_) in (d_403_v14_):
                        coll25_[d_405_v13_] = not(d_397_v7_.f37)
                return _dafny.Map(coll25_)
            nw67_.ctor__((d_402_v12_) == (iife37_()
            ))
            d_404_v15_ = nw67_
        elif True:
            d_406_v16_: _dafny.Set
            d_406_v16_ = _dafny.Set({p2})
            d_407_v17_: str
            d_407_v17_ = _dafny.CodePoint('q')
            d_408_v18_: _dafny.Seq
            d_408_v18_ = _dafny.SeqWithoutIsStrInference([d_407_v17_, d_407_v17_, _dafny.CodePoint('j'), d_407_v17_])
            nw68_ = _dafny.Array(None, 21)
            nw68_[int(0)] = p2
            nw68_[int(1)] = (default__.fm36(globalState)) + (default__.fm36(globalState))
            nw68_[int(2)] = len(d_406_v16_)
            nw68_[int(3)] = default__.safeDivisionInt(p2, p2)
            nw68_[int(4)] = p2
            nw68_[int(5)] = p2
            nw68_[int(6)] = (136) - (p2)
            nw68_[int(7)] = p2
            nw68_[int(8)] = p2
            nw68_[int(9)] = p2
            nw68_[int(10)] = -137
            nw68_[int(11)] = default__.safeModuloInt(p2, 389)
            nw68_[int(12)] = default__.safeModuloInt(p2, p2)
            nw68_[int(13)] = p2
            nw68_[int(14)] = p2
            nw68_[int(15)] = p2
            nw68_[int(16)] = len((d_408_v18_) + (d_408_v18_))
            nw68_[int(17)] = (d_386_v1_)[default__.safeIndex(p2, len(d_386_v1_))]
            nw68_[int(18)] = p2
            nw68_[int(19)] = (p2 if (d_388_v3_)[default__.safeIndex(409, (d_388_v3_).length(0))] else 241)
            nw68_[int(20)] = p2
            (globalState).f17 = nw68_
            index59_ = default__.safeIndex(409, (d_388_v3_).length(0))
            (d_388_v3_)[index59_] = (p2) > (p2)
            d_409_v19_: _dafny.Set
            d_409_v19_ = _dafny.Set({(self).f40, p0, (d_388_v3_)[default__.safeIndex(409, (d_388_v3_).length(0))], (self).f40})
            d_410_v20_: _dafny.Map
            d_410_v20_ = _dafny.Map({p2: False})
            d_411_v21_: _dafny.Array
            nw69_ = _dafny.Array(None, 9)
            nw69_[int(0)] = len((d_409_v19_) - (d_409_v19_))
            nw69_[int(1)] = default__.safeDivisionInt(p2, (0) - (default__.fm36(globalState)))
            nw69_[int(2)] = (619 if (self).f40 else len((d_408_v18_).set(default__.safeIndex(p2, len(d_408_v18_)), d_407_v17_)))
            nw69_[int(3)] = (D4_DC15(p2, len(d_408_v18_))).cf24
            nw69_[int(4)] = p2
            nw69_[int(5)] = (0) - (len(d_410_v20_))
            nw69_[int(6)] = 773
            nw69_[int(7)] = p2
            nw69_[int(8)] = default__.fm36(globalState)
            d_411_v21_ = nw69_
            index60_ = default__.safeIndex(409, (d_388_v3_).length(0))
            rhs71_ = d_411_v21_
            rhs72_ = (D6_DC23((self).f40, p2)).cf36
            lhs61_ = globalState
            lhs62_ = d_388_v3_
            lhs63_ = default__.safeIndex(409, (d_388_v3_).length(0))
            lhs61_.f17 = rhs71_
            lhs62_[lhs63_] = rhs72_
            if (self).f40:
                d_412_v22_: _dafny.Set
                d_412_v22_ = _dafny.Set({d_407_v17_})
                d_413_v23_: _dafny.MultiSet
                d_413_v23_ = _dafny.MultiSet([d_412_v22_, _dafny.Set({d_407_v17_}), d_412_v22_, d_412_v22_])
                d_414_v24_: _dafny.Seq
                d_414_v24_ = _dafny.SeqWithoutIsStrInference([d_412_v22_])
                index61_ = default__.safeIndex(409, (d_388_v3_).length(0))
                (d_388_v3_)[index61_] = ((_dafny.MultiSet(d_414_v24_)).intersection(d_413_v23_)).ispropersubset(d_413_v23_)
                d_415_v25_: _dafny.Seq
                d_415_v25_ = _dafny.SeqWithoutIsStrInference([(d_409_v19_) != (_dafny.Set({(d_388_v3_)[default__.safeIndex(409, (d_388_v3_).length(0))]})), ((d_388_v3_)[default__.safeIndex(409, (d_388_v3_).length(0))]) or (p0)])
                (globalState).f13 = (d_415_v25_)[default__.safeIndex(p2, len(d_415_v25_))]
                d_416_v26_: C0
                nw70_ = C0()
                nw70_.ctor__((p2) < (981))
                d_416_v26_ = nw70_
                rhs73_ = -164
                rhs74_ = (p2) in (_dafny.Map({p2: d_410_v20_}))
                rhs75_ = p2
                lhs64_ = globalState
                lhs65_ = globalState
                lhs66_ = globalState
                lhs64_.f19 = rhs73_
                lhs65_.f28 = rhs74_
                lhs66_.f1 = rhs75_
                index62_ = default__.safeIndex(84, (d_411_v21_).length(0))
                (d_411_v21_)[index62_] = len((d_415_v25_ if (d_388_v3_)[default__.safeIndex(409, (d_388_v3_).length(0))] else d_415_v25_))
                index63_ = default__.safeIndex(84, (d_411_v21_).length(0))
                rhs76_ = d_388_v3_
                rhs77_ = (_dafny.SeqWithoutIsStrInference([d_407_v17_ for d_417_i2_ in range(default__.abs(74))]) if ((0) - (len(d_415_v25_))) >= (p2) else ((d_408_v18_).set(default__.safeIndex(p2, len(d_408_v18_)), d_407_v17_) if d_416_v26_.f37 else d_408_v18_))
                rhs78_ = -296
                rhs79_ = not(p0)
                lhs67_ = globalState
                lhs68_ = d_411_v21_
                lhs69_ = default__.safeIndex(84, (d_411_v21_).length(0))
                lhs70_ = globalState
                d_388_v3_ = rhs76_
                lhs67_.f20 = rhs77_
                lhs68_[lhs69_] = rhs78_
                lhs70_.f28 = rhs79_
            elif True:
                index64_ = default__.safeIndex(401, (d_411_v21_).length(0))
                (d_411_v21_)[index64_] = (0) - (default__.fm36(globalState))
                index65_ = default__.safeIndex(401, (d_411_v21_).length(0))
                (d_411_v21_)[index65_] = p2
                (globalState).f28 = (self).f40
                d_418_v27_: _dafny.Seq
                d_418_v27_ = _dafny.SeqWithoutIsStrInference([(default__.fm36(globalState)) >= ((0) - (-877))])
                d_419_v28_: _dafny.Map
                d_419_v28_ = _dafny.Map({(d_411_v21_)[default__.safeIndex(401, (d_411_v21_).length(0))]: d_418_v27_})
                d_418_v27_ = (d_418_v27_ if not(p0) else ((d_419_v28_)[p2] if (p2) in (d_419_v28_) else d_418_v27_))
                d_410_v20_ = _dafny.Map({(d_411_v21_)[default__.safeIndex(401, (d_411_v21_).length(0))]: (_dafny.MultiSet([p0])).issubset(_dafny.MultiSet([p0]))})
                d_420_v29_: _dafny.Array
                def lambda13_(d_421_p2_):
                    def lambda14_(d_422_i3_):
                        return (d_422_i3_) - (d_421_p2_)

                    return lambda14_

                init7_ = lambda13_(p2)
                nw71_ = _dafny.Array(None, 23)
                for i0_7_ in range(nw71_.length(0)):
                    nw71_[i0_7_] = init7_(i0_7_)
                d_420_v29_ = nw71_
                rhs80_ = d_407_v17_
                rhs81_ = d_420_v29_
                rhs82_ = (d_411_v21_)[default__.safeIndex(401, (d_411_v21_).length(0))]
                lhs71_ = globalState
                lhs72_ = globalState
                d_407_v17_ = rhs80_
                lhs71_.f17 = rhs81_
                lhs72_.f14 = rhs82_
            (globalState).f13 = p0
        d_423_v30_: D3
        d_423_v30_ = D3_DC11((d_388_v3_)[default__.safeIndex(409, (d_388_v3_).length(0))], not(p0))
        d_424_i4_: int
        d_424_i4_ = 0
        with _dafny.label("2"):
            while (d_423_v30_).cf20:
                with _dafny.c_label("2"):
                    if (d_424_i4_) >= (100):
                        raise _dafny.Break("2")
                    d_424_i4_ = (d_424_i4_) + (1)
                    d_425_v31_: _dafny.Array
                    nw72_ = _dafny.Array(int(0), 24)
                    d_425_v31_ = nw72_
                    d_426_v32_: _dafny.Map
                    d_426_v32_ = _dafny.Map({d_425_v31_: (d_392_v5_).cf40})
                    d_426_v32_ = (d_426_v32_).set(d_425_v31_, _dafny.SeqWithoutIsStrInference([p2, p2]))
                    (globalState).f18 = (d_388_v3_)[default__.safeIndex(409, (d_388_v3_).length(0))]
                    (globalState).f14 = (p2) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_427_i5_ in range(default__.abs(706))])))
                    d_428_v33_: _dafny.Seq
                    d_428_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evxiub"))
                    d_429_v34_: _dafny.MultiSet
                    d_429_v34_ = _dafny.MultiSet([False])
                    d_430_v35_: str
                    d_430_v35_ = _dafny.CodePoint('p')
                    d_431_v36_: D5
                    d_431_v36_ = D5_DC19(len(d_428_v33_), (d_386_v1_)[default__.safeIndex((d_429_v34_).cardinality, len(d_386_v1_))], d_430_v35_)
                    (globalState).f13 = (p2) != ((d_431_v36_).cf28)
                    pass
            pass
        d_432_v37_: C0
        nw73_ = C0()
        nw73_.ctor__((d_388_v3_)[default__.safeIndex(409, (d_388_v3_).length(0))])
        d_432_v37_ = nw73_
        d_433_v38_: _dafny.Map
        d_433_v38_ = _dafny.Map({p0: p2})
        d_434_v39_: _dafny.Map
        d_434_v39_ = _dafny.Map({d_433_v38_: (default__.fm37(globalState)).cf33})
        d_435_v40_: _dafny.Seq
        d_435_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bobyhh"))
        d_436_v41_: str
        d_436_v41_ = _dafny.CodePoint('t')
        rhs83_ = ((d_433_v38_) | (_dafny.Map({default__.fm33(p0, p2, 945, globalState): p2}))) not in (d_434_v39_)
        rhs84_ = d_432_v37_.f37
        rhs85_ = (((d_435_v40_).set(default__.safeIndex(len(_dafny.Map({p2: len((default__.fm38(globalState)).cf40)})), len(d_435_v40_)), d_436_v41_)) + (d_435_v40_)) <= (d_435_v40_)
        lhs73_ = globalState
        lhs74_ = globalState
        lhs73_.f18 = rhs83_
        lhs74_.f15 = rhs84_
        r1 = rhs85_
        hi0_ = p2
        for d_437_i6_ in range(p2, hi0_):
            (globalState).f1 = (p2) * (p2)
            d_438_v42_: _dafny.Array
            nw74_ = _dafny.Array(int(0), 17)
            d_438_v42_ = nw74_
            index66_ = default__.safeIndex(864, (d_438_v42_).length(0))
            (d_438_v42_)[index66_] = p2
            index67_ = default__.safeIndex(864, (d_438_v42_).length(0))
            (d_438_v42_)[index67_] = default__.fm36(globalState)
            d_439_v43_: D6
            d_439_v43_ = D6_DC22(p2, len(d_433_v38_), (d_388_v3_)[default__.safeIndex(409, (d_388_v3_).length(0))])
            d_440_v44_: _dafny.Map
            d_440_v44_ = _dafny.Map({d_388_v3_: (d_439_v43_).cf33})
            d_433_v38_ = (d_433_v38_).set((p2) < (len((d_440_v44_).set(d_388_v3_, d_437_i6_))), (d_438_v42_)[default__.safeIndex(864, (d_438_v42_).length(0))])
            d_385_v0_ = (d_385_v0_).set(len(_dafny.Map({(d_438_v42_)[default__.safeIndex(864, (d_438_v42_).length(0))]: d_433_v38_})), default__.abs(d_437_i6_))
        r0 = d_435_v40_
        r1 = not(False)
        return r0, r1

    @property
    def f40(self):
        return self._f40

class C2(T1, T0):
    def  __init__(self):
        self._f30: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f29: bool = False
        self._f38: int = int(0)
        self._f39: _dafny.Set = _dafny.Set({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f30(self):
        return self._f30
    @f30.setter
    def f30(self, value):
        self._f30 = value
    @property
    def f29(self):
        return self._f29
    def ctor__(self, f38, f39, f29, f30):
        (self)._f38 = f38
        (self)._f39 = f39
        (self)._f29 = f29
        (self).f30 = f30

    def fm16(self, p0, globalState):
        return (self).f38

    def m2(self, p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        (globalState).f28 = (self).f29
        if not ((self).f29) or (((self).f38) > ((0) - ((self).f38))):
            (globalState).f15 = (self).f29
            (globalState).f2 = ((self).f38) - ((self).f38)
            d_441_v0_: _dafny.MultiSet
            d_441_v0_ = _dafny.MultiSet([p0])
            d_442_v1_: D4
            d_442_v1_ = D4_DC15(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewafp"))), 469)
            d_443_v2_: _dafny.Array
            nw75_ = _dafny.Array(False, 23)
            d_443_v2_ = nw75_
            d_444_v3_: _dafny.MultiSet
            d_444_v3_ = _dafny.MultiSet([d_443_v2_, d_443_v2_])
            d_445_v4_: _dafny.Seq
            d_445_v4_ = _dafny.SeqWithoutIsStrInference([p0])
            d_446_v5_: _dafny.Array
            nw76_ = _dafny.Array(None, 27)
            nw76_[int(0)] = (self).f38
            nw76_[int(1)] = 838
            nw76_[int(2)] = (self).f38
            nw76_[int(3)] = default__.safeDivisionInt((self).f38, (self).f38)
            nw76_[int(4)] = 13
            nw76_[int(5)] = (self).fm16(d_441_v0_, globalState)
            nw76_[int(6)] = (self).f38
            nw76_[int(7)] = (self).f38
            nw76_[int(8)] = (self).f38
            nw76_[int(9)] = (self).f38
            nw76_[int(10)] = (0) - ((self).f38)
            nw76_[int(11)] = (d_442_v1_).cf25
            nw76_[int(12)] = (self).f38
            nw76_[int(13)] = (self).f38
            nw76_[int(14)] = default__.safeModuloInt((self).f38, len(default__.fm17((self).f38, globalState)))
            nw76_[int(15)] = (self).f38
            nw76_[int(16)] = (self).fm16(d_441_v0_, globalState)
            nw76_[int(17)] = ((default__.fm18((self).f38, (d_444_v3_).cardinality, (self).f38, len(d_445_v4_), globalState)).cf15) * ((self).f38)
            nw76_[int(18)] = (self).f38
            nw76_[int(19)] = len(d_445_v4_)
            nw76_[int(20)] = (self).f38
            nw76_[int(21)] = len(_dafny.SeqWithoutIsStrInference([(self).f29]))
            nw76_[int(22)] = default__.safeDivisionInt((self).f38, (self).f38)
            nw76_[int(23)] = ((self).f38 if p0 else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_447_i0_ in range(default__.abs(18))])))
            nw76_[int(24)] = (self).f38
            nw76_[int(25)] = (self).f38
            nw76_[int(26)] = (self).f38
            d_446_v5_ = nw76_
            index68_ = default__.safeIndex(53, (d_446_v5_).length(0))
            (d_446_v5_)[index68_] = (self).f38
            index69_ = default__.safeIndex(53, (d_446_v5_).length(0))
            (d_446_v5_)[index69_] = (self).f38
            (globalState).f14 = (self).f38
            d_448_v6_: _dafny.Seq
            d_448_v6_ = _dafny.SeqWithoutIsStrInference([d_443_v2_])
            (globalState).f1 = len(d_448_v6_)
        elif True:
            if p0:
                d_449_v7_: C0
                nw77_ = C0()
                nw77_.ctor__(p0)
                d_449_v7_ = nw77_
                d_450_v8_: _dafny.MultiSet
                d_450_v8_ = _dafny.MultiSet([d_449_v7_.f37])
                d_451_v9_: _dafny.Array
                def lambda15_(d_452_i1_):
                    return default__.safeModuloInt(d_452_i1_, (self).f38)

                init8_ = lambda15_
                nw78_ = _dafny.Array(None, 16)
                for i0_8_ in range(nw78_.length(0)):
                    nw78_[i0_8_] = init8_(i0_8_)
                d_451_v9_ = nw78_
                d_453_v10_: _dafny.Map
                d_453_v10_ = _dafny.Map({(self.f30)[default__.safeIndex((self).fm16(d_450_v8_, globalState), len(self.f30))]: d_451_v9_})
                d_454_v11_: str
                d_454_v11_ = _dafny.CodePoint('e')
                (globalState).f17 = ((d_453_v10_)[d_454_v11_] if (d_454_v11_) in (d_453_v10_) else d_451_v9_)
                d_455_v12_: _dafny.Array
                def lambda16_(d_456_p0_, d_457_v7_):
                    def lambda17_(d_458_i2_):
                        return (_dafny.SeqWithoutIsStrInference([_dafny.Set({d_456_p0_, d_456_p0_})])) < (_dafny.SeqWithoutIsStrInference([_dafny.Set({d_457_v7_.f37}), _dafny.Set({True}), _dafny.Set({d_457_v7_.f37, d_456_p0_}), _dafny.Set({d_457_v7_.f37}), _dafny.Set({(self).f29})]))

                    return lambda17_

                init9_ = lambda16_(p0, d_449_v7_)
                nw79_ = _dafny.Array(None, 2)
                for i0_9_ in range(nw79_.length(0)):
                    nw79_[i0_9_] = init9_(i0_9_)
                d_455_v12_ = nw79_
                d_455_v12_ = d_455_v12_
                (globalState).f14 = (self).f38
                d_459_v13_: _dafny.Map
                d_459_v13_ = _dafny.Map({d_454_v11_: d_454_v11_})
                d_460_v14_: _dafny.Seq
                d_460_v14_ = _dafny.SeqWithoutIsStrInference([d_459_v13_])
                d_460_v14_ = d_460_v14_
            elif True:
                (globalState).f28 = p0
                (globalState).f28 = ((self).f38) > ((self).f38)
                d_461_v15_: _dafny.MultiSet
                d_461_v15_ = _dafny.MultiSet([(self).f38, (self).f38, (self).f38, (self).f38, (self).f38])
                d_462_v16_: D3
                d_462_v16_ = D3_DC10(d_461_v15_)
                d_463_v17_: str
                d_463_v17_ = _dafny.CodePoint('i')
                d_464_v18_: C0
                nw80_ = C0()
                nw80_.ctor__((self).f29)
                d_464_v18_ = nw80_
                d_465_v19_: _dafny.Map
                d_465_v19_ = _dafny.Map({len(self.f30): d_464_v18_})
                d_466_v20_: _dafny.Set
                d_466_v20_ = _dafny.Set({len(d_465_v19_), -408})
                d_467_v21_: _dafny.Seq
                d_467_v21_ = _dafny.SeqWithoutIsStrInference([p0])
                d_468_v22_: _dafny.Seq
                d_468_v22_ = _dafny.SeqWithoutIsStrInference([len(d_466_v20_), len(_dafny.SeqWithoutIsStrInference([d_464_v18_.f37])), -607, 730, len(d_467_v21_)])
                rhs86_ = (_dafny.MultiSet([(self).f38, (self).f38, 283])).isdisjoint((_dafny.MultiSet(d_468_v22_) if default__.fm19(d_463_v17_, globalState) else d_461_v15_))
                rhs87_ = d_464_v18_.f37
                rhs88_ = d_462_v16_
                rhs89_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_469_i3_ in range(default__.abs(817))])
                lhs75_ = globalState
                lhs76_ = globalState
                lhs77_ = globalState
                lhs75_.f28 = rhs86_
                lhs76_.f15 = rhs87_
                d_462_v16_ = rhs88_
                lhs77_.f11 = rhs89_
                d_470_v23_: _dafny.MultiSet
                d_470_v23_ = _dafny.MultiSet([(self).f29, (self).f29, d_464_v18_.f37, d_464_v18_.f37, d_464_v18_.f37])
                d_471_v24_: _dafny.Map
                d_471_v24_ = _dafny.Map({(self).f29: d_470_v23_})
                d_472_v25_: _dafny.Map
                d_472_v25_ = _dafny.Map({p0: len(self.f30)})
                d_473_v26_: _dafny.Map
                d_473_v26_ = _dafny.Map({(self).f29: d_464_v18_.f37})
                d_474_v27_: _dafny.Array
                nw81_ = _dafny.Array(None, 24)
                nw81_[int(0)] = (0) - ((self).f38)
                nw81_[int(1)] = (self).f38
                nw81_[int(2)] = 106
                nw81_[int(3)] = -464
                nw81_[int(4)] = 977
                nw81_[int(5)] = (self).f38
                nw81_[int(6)] = (0) - ((self).f38)
                nw81_[int(7)] = (self).fm16(((d_471_v24_)[d_464_v18_.f37] if (d_464_v18_.f37) in (d_471_v24_) else d_470_v23_), globalState)
                nw81_[int(8)] = (self).f38
                nw81_[int(9)] = ((self).f38 if not(not((self).f29)) else (0) - ((self).f38))
                nw81_[int(10)] = (self).f38
                nw81_[int(11)] = (d_464_v18_).fm12(globalState)
                nw81_[int(12)] = (self).f38
                nw81_[int(13)] = (self).f38
                nw81_[int(14)] = len(d_472_v25_)
                nw81_[int(15)] = 754
                nw81_[int(16)] = (self).f38
                nw81_[int(17)] = (len(_dafny.Map({d_473_v26_: d_464_v18_.f37})) if (self).f29 else -789)
                nw81_[int(18)] = 788
                nw81_[int(19)] = (self).f38
                nw81_[int(20)] = (len(d_466_v20_)) + ((self).f38)
                nw81_[int(21)] = (0) - (len(d_468_v22_))
                nw81_[int(22)] = (self).fm16(d_470_v23_, globalState)
                nw81_[int(23)] = (self).f38
                d_474_v27_ = nw81_
                index70_ = default__.safeIndex(921, (d_474_v27_).length(0))
                (d_474_v27_)[index70_] = (0) - ((self).f38)
                index71_ = default__.safeIndex(327, (d_474_v27_).length(0))
                (d_474_v27_)[index71_] = (self).f38
                index72_ = default__.safeIndex(921, (d_474_v27_).length(0))
                index73_ = default__.safeIndex(327, (d_474_v27_).length(0))
                rhs90_ = len(self.f30)
                rhs91_ = (d_464_v18_).fm12(globalState)
                rhs92_ = (self).f38
                rhs93_ = d_463_v17_
                lhs78_ = d_474_v27_
                lhs79_ = default__.safeIndex(921, (d_474_v27_).length(0))
                lhs80_ = globalState
                lhs81_ = d_474_v27_
                lhs82_ = default__.safeIndex(327, (d_474_v27_).length(0))
                lhs78_[lhs79_] = rhs90_
                lhs80_.f2 = rhs91_
                lhs81_[lhs82_] = rhs92_
                d_463_v17_ = rhs93_
                d_475_v28_: _dafny.Array
                nw82_ = _dafny.Array(False, 24)
                d_475_v28_ = nw82_
                index74_ = default__.safeIndex(730, (d_475_v28_).length(0))
                (d_475_v28_)[index74_] = (self).f29
                index75_ = default__.safeIndex(730, (d_475_v28_).length(0))
                (d_475_v28_)[index75_] = d_464_v18_.f37
            (globalState).f15 = p0
            d_476_v29_: _dafny.Map
            d_476_v29_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([p0, (self).f29, p0, (self).f29, True])): not(not(True))})
            (globalState).f21 = len((d_476_v29_) | (d_476_v29_))
            if True:
                d_477_v30_: _dafny.Seq
                d_477_v30_ = _dafny.SeqWithoutIsStrInference([p0])
                d_478_v31_: _dafny.Seq
                d_478_v31_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f29]), (d_477_v30_).set(default__.safeIndex((self).f38, len(d_477_v30_)), not(p0))])
                d_479_v32_: _dafny.Seq
                d_479_v32_ = _dafny.SeqWithoutIsStrInference([(d_478_v31_)[default__.safeIndex((0) - (len(self.f30)), len(d_478_v31_))], d_477_v30_])
                d_480_v33_: str
                d_480_v33_ = _dafny.CodePoint('t')
                d_481_v34_: _dafny.Array
                nw83_ = _dafny.Array(None, 3)
                nw83_[int(0)] = (self.f30).set(default__.safeIndex(len(d_478_v31_), len(self.f30)), d_480_v33_)
                nw83_[int(1)] = (default__.fm20(True, (self).f29, d_477_v30_, len(_dafny.SeqWithoutIsStrInference([p0])), globalState)).set(default__.safeIndex((self).f38, len(default__.fm20(True, (self).f29, d_477_v30_, len(_dafny.SeqWithoutIsStrInference([p0])), globalState))), d_480_v33_)
                nw83_[int(2)] = self.f30
                d_481_v34_ = nw83_
                index76_ = default__.safeIndex(605, (d_481_v34_).length(0))
                (d_481_v34_)[index76_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sar"))
                d_482_v35_: D1
                d_482_v35_ = D1_DC3(d_477_v30_)
                pat_let_tv10_ = d_477_v30_
                d_483_v36_: _dafny.Map
                def iife38_(_pat_let6_0):
                    def iife39_(d_484_dt__update__tmp_h0_):
                        def iife40_(_pat_let7_0):
                            def iife41_(d_485_dt__update_hcf7_h0_):
                                return D1_DC3(d_485_dt__update_hcf7_h0_)
                            return iife41_(_pat_let7_0)
                        return iife40_(pat_let_tv10_)
                    return iife39_(_pat_let6_0)
                d_483_v36_ = _dafny.Map({(self).f29: iife38_(d_482_v35_)})
                d_486_v37_: _dafny.Map
                d_486_v37_ = _dafny.Map({p0: self.f30})
                d_487_v38_: _dafny.Map
                d_487_v38_ = _dafny.Map({p0: (self).f38})
                d_488_v39_: _dafny.Seq
                d_488_v39_ = _dafny.SeqWithoutIsStrInference([(self).f38, 567, len((self.f30).set(default__.safeIndex(len(d_487_v38_), len(self.f30)), d_480_v33_))])
                index77_ = default__.safeIndex(605, (d_481_v34_).length(0))
                rhs94_ = (((d_486_v37_)[(self).f29] if ((self).f29) in (d_486_v37_) else (self.f30).set(default__.safeIndex((self).f38, len(self.f30)), d_480_v33_))) + (default__.fm20(not(p0), True, _dafny.SeqWithoutIsStrInference([(self).f29, True, not((d_477_v30_)[default__.safeIndex(904, len(d_477_v30_))]), p0]), (0) - ((_dafny.MultiSet([(self).f38, (self).f38])).cardinality), globalState))
                rhs95_ = len(d_488_v39_)
                rhs96_ = ((default__.fm21(p0, (self).f29, p0, globalState)).cf26) | ((d_483_v36_) | (d_483_v36_))
                lhs83_ = d_481_v34_
                lhs84_ = default__.safeIndex(605, (d_481_v34_).length(0))
                lhs85_ = globalState
                lhs83_[lhs84_] = rhs94_
                lhs85_.f19 = rhs95_
                d_483_v36_ = rhs96_
                (globalState).f18 = False
                d_489_v40_: _dafny.Array
                nw84_ = _dafny.Array(False, 20)
                d_489_v40_ = nw84_
                index78_ = default__.safeIndex(548, (d_489_v40_).length(0))
                (d_489_v40_)[index78_] = (self).f29
                index79_ = default__.safeIndex(548, (d_489_v40_).length(0))
                (d_489_v40_)[index79_] = ((self).f38) < ((self).f38)
                d_490_v41_: _dafny.Array
                nw85_ = _dafny.Array(_dafny.CodePoint('D'), 12)
                d_490_v41_ = nw85_
                d_491_v42_: _dafny.Seq
                d_491_v42_ = _dafny.SeqWithoutIsStrInference([d_490_v41_])
                d_490_v41_ = (d_491_v42_)[default__.safeIndex(default__.safeModuloInt((self).f38, (self).f38), len(d_491_v42_))]
                (globalState).f28 = default__.fm19(d_480_v33_, globalState)
            elif True:
                d_492_v43_: str
                d_492_v43_ = _dafny.CodePoint('s')
                d_493_v44_: _dafny.MultiSet
                d_493_v44_ = _dafny.MultiSet([d_492_v43_])
                d_494_v45_: D2
                d_494_v45_ = D2_DC6(d_493_v44_)
                d_494_v45_ = d_494_v45_
                d_495_v46_: _dafny.Map
                d_495_v46_ = _dafny.Map({len(d_476_v29_): (self).f38})
                d_495_v46_ = (d_495_v46_).set(-643, (self).f38)
                d_496_v47_: _dafny.Array
                nw86_ = _dafny.Array(_dafny.MultiSet({}), 15)
                d_496_v47_ = nw86_
                d_497_v48_: _dafny.Seq
                d_497_v48_ = _dafny.SeqWithoutIsStrInference([True])
                index80_ = default__.safeIndex(760, (d_496_v47_).length(0))
                (d_496_v47_)[index80_] = _dafny.MultiSet(d_497_v48_)
                d_498_v49_: _dafny.Seq
                d_498_v49_ = _dafny.SeqWithoutIsStrInference([self.f30])
                d_499_v50_: _dafny.Map
                d_499_v50_ = _dafny.Map({(self).f29: (self).f29})
                d_500_v51_: _dafny.MultiSet
                d_500_v51_ = _dafny.MultiSet([((d_476_v29_)[(self).f38] if ((self).f38) in (d_476_v29_) else p0), (self).f29])
                d_501_v52_: _dafny.Map
                d_501_v52_ = _dafny.Map({301: _dafny.MultiSet([p0, (self).f29])})
                index81_ = default__.safeIndex(760, (d_496_v47_).length(0))
                rhs97_ = default__.safeDivisionInt(len(d_498_v49_), (len(d_499_v50_)) - ((self).f38))
                rhs98_ = (d_500_v51_) - (((d_501_v52_)[(self).f38] if ((self).f38) in (d_501_v52_) else _dafny.MultiSet(d_497_v48_)))
                lhs86_ = globalState
                lhs87_ = d_496_v47_
                lhs88_ = default__.safeIndex(760, (d_496_v47_).length(0))
                lhs86_.f1 = rhs97_
                lhs87_[lhs88_] = rhs98_
                d_502_v53_: _dafny.Seq
                d_502_v53_ = _dafny.SeqWithoutIsStrInference([d_497_v48_])
                (globalState).f11 = default__.fm20((self).f29, p0, (d_502_v53_)[default__.safeIndex((self).f38, len(d_502_v53_))], (self).f38, globalState)
                d_503_v54_: _dafny.Set
                d_503_v54_ = _dafny.Set({p0})
                d_504_v55_: _dafny.Seq
                d_504_v55_ = _dafny.SeqWithoutIsStrInference([(D6_DC21(d_503_v54_)).cf32, default__.fm22((self).f38, not(p0), d_492_v43_, globalState), d_503_v54_])
                (globalState).f19 = len((d_504_v55_)[default__.safeIndex(870, len(d_504_v55_))])
            d_505_v56_: _dafny.Map
            d_505_v56_ = _dafny.Map({(p0) == ((self).f29): (self).f38})
            d_505_v56_ = (d_505_v56_).set((self).f29, default__.safeDivisionInt((self).f38, (self).f38))
        d_506_v57_: _dafny.Array
        def lambda18_(d_507_i4_):
            return (_dafny.CodePoint('p')) in (self.f30)

        init10_ = lambda18_
        nw87_ = _dafny.Array(None, 15)
        for i0_10_ in range(nw87_.length(0)):
            nw87_[i0_10_] = init10_(i0_10_)
        d_506_v57_ = nw87_
        d_508_v58_: str
        d_508_v58_ = _dafny.CodePoint('c')
        index82_ = default__.safeIndex(991, (d_506_v57_).length(0))
        (d_506_v57_)[index82_] = not(default__.fm19(d_508_v58_, globalState))
        index83_ = default__.safeIndex(991, (d_506_v57_).length(0))
        (d_506_v57_)[index83_] = p0
        (globalState).f19 = (0) - ((-489) * ((self).f38))
        pat_let_tv11_ = p0
        pat_let_tv12_ = p0
        pat_let_tv13_ = d_506_v57_
        pat_let_tv14_ = d_506_v57_
        def lambda19_(source8_):
            if source8_.is_DC7:
                d_509___mcc_h0_ = source8_.cf11
                d_510_cf11_ = d_509___mcc_h0_
                return (self.f30) < (self.f30)
            elif source8_.is_DC8:
                d_511___mcc_h1_ = source8_.cf12
                d_512___mcc_h2_ = source8_.cf13
                d_513_cf13_ = d_512___mcc_h2_
                d_514_cf12_ = d_511___mcc_h1_
                return (pat_let_tv11_) and (d_513_cf13_)
            elif source8_.is_DC9:
                d_515___mcc_h3_ = source8_.cf14
                d_516___mcc_h4_ = source8_.cf15
                d_517___mcc_h5_ = source8_.cf16
                d_518___mcc_h6_ = source8_.cf17
                d_519_cf17_ = d_518___mcc_h6_
                d_520_cf16_ = d_517___mcc_h5_
                d_521_cf15_ = d_516___mcc_h4_
                d_522_cf14_ = d_515___mcc_h3_
                return (False) == (pat_let_tv12_)
            elif True:
                d_523___mcc_h7_ = source8_.cf10
                d_524_cf10_ = d_523___mcc_h7_
                return (pat_let_tv14_)[default__.safeIndex(991, (pat_let_tv13_).length(0))]

        (globalState).f28 = lambda19_(D2_DC9(self.f30, (self).f38, p0, (self).f38))
        d_525_v59_: _dafny.Map
        d_525_v59_ = _dafny.Map({p0: not (p0) or (p0)})
        d_525_v59_ = (d_525_v59_).set(not((self).f29), p0)
        r0 = default__.fm23(globalState)
        r1 = p0
        return r0, r1

    def m3(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        d_526_v0_: _dafny.Array
        nw88_ = _dafny.Array(_dafny.CodePoint('D'), 17)
        d_526_v0_ = nw88_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_526_v0_).length(0)):
            d_527_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_527_i0_)) and ((d_527_i0_) < ((d_526_v0_).length(0)))):
                (d_526_v0_)[(d_527_i0_)] = _dafny.CodePoint('j')
        d_528_v1_: _dafny.Set
        d_528_v1_ = _dafny.Set({(self).f29})
        d_529_v3_: _dafny.Map
        d_529_v3_ = _dafny.Map({p0: 50})
        def iife42_():
            coll26_ = _dafny.Map()
            compr_26_: int
            for compr_26_ in _dafny.IntegerRange(588, 387):
                d_530_v2_: int = compr_26_
                if ((588) <= (d_530_v2_)) and ((d_530_v2_) < (387)):
                    coll26_[default__.safeModuloInt(d_530_v2_, p2)] = len(_dafny.SeqWithoutIsStrInference([p2, (self).f38, (self).f38]))
            return _dafny.Map(coll26_)
        (globalState).f19 = ((self).f38 if (d_528_v1_).issubset(d_528_v1_) else (len(iife42_()
        )) - (len(d_529_v3_)))
        d_531_v4_: _dafny.Array
        def lambda20_(d_532_i1_):
            return default__.safeModuloInt(d_532_i1_, (0) - (len(self.f30)))

        init11_ = lambda20_
        nw89_ = _dafny.Array(None, 21)
        for i0_11_ in range(nw89_.length(0)):
            nw89_[i0_11_] = init11_(i0_11_)
        d_531_v4_ = nw89_
        d_533_v5_: _dafny.MultiSet
        d_533_v5_ = _dafny.MultiSet([d_531_v4_])
        rhs99_ = 727
        rhs100_ = d_533_v5_
        lhs89_ = globalState
        lhs89_.f19 = rhs99_
        d_533_v5_ = rhs100_
        if default__.fm19(_dafny.CodePoint('k'), globalState):
            d_534_v6_: C0
            nw90_ = C0()
            nw90_.ctor__((self).f29)
            d_534_v6_ = nw90_
            (globalState).f18 = (default__.fm24(globalState)).cf8
            index84_ = default__.safeIndex(822, (d_531_v4_).length(0))
            (d_531_v4_)[index84_] = -588
            d_535_v7_: _dafny.Seq
            d_535_v7_ = _dafny.SeqWithoutIsStrInference([(self).f29])
            d_536_v8_: _dafny.Map
            d_536_v8_ = _dafny.Map({d_535_v7_: (self).f38})
            d_537_v9_: D4
            d_537_v9_ = D4_DC13(d_536_v8_)
            index85_ = default__.safeIndex(822, (d_531_v4_).length(0))
            rhs101_ = default__.safeDivisionInt(len(self.f30), 855)
            rhs102_ = d_537_v9_
            rhs103_ = True
            lhs90_ = d_531_v4_
            lhs91_ = default__.safeIndex(822, (d_531_v4_).length(0))
            lhs92_ = globalState
            lhs90_[lhs91_] = rhs101_
            d_537_v9_ = rhs102_
            lhs92_.f18 = rhs103_
            d_538_v10_: D2
            d_538_v10_ = D2_DC9(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xe")), (d_531_v4_)[default__.safeIndex(822, (d_531_v4_).length(0))], p0, len(self.f30))
            d_539_v11_: _dafny.MultiSet
            d_539_v11_ = _dafny.MultiSet([_dafny.Set({(self).f29, (self).f29})])
            d_540_v12_: _dafny.MultiSet
            d_540_v12_ = _dafny.MultiSet([(self).f29, (_dafny.CodePoint('t')) in (self.f30), ((d_538_v10_).cf15) >= (((d_539_v11_)[d_528_v1_] if (d_528_v1_) in (d_539_v11_) else p2))])
            d_540_v12_ = d_540_v12_
            d_541_v13_: _dafny.Array
            nw91_ = _dafny.Array(False, 23)
            d_541_v13_ = nw91_
            index86_ = default__.safeIndex(115, (d_541_v13_).length(0))
            (d_541_v13_)[index86_] = (self).f29
            d_542_v14_: _dafny.Array
            nw92_ = _dafny.Array(_dafny.Set({}), 6)
            d_542_v14_ = nw92_
            index87_ = default__.safeIndex(365, (d_542_v14_).length(0))
            (d_542_v14_)[index87_] = (d_528_v1_) | (d_528_v1_)
            d_543_v15_: str
            d_543_v15_ = _dafny.CodePoint('h')
            index88_ = default__.safeIndex(822, (d_531_v4_).length(0))
            index89_ = default__.safeIndex(115, (d_541_v13_).length(0))
            index90_ = default__.safeIndex(365, (d_542_v14_).length(0))
            rhs104_ = 211
            rhs105_ = (p2) >= ((self).f38)
            rhs106_ = (d_531_v4_)[default__.safeIndex(822, (d_531_v4_).length(0))]
            rhs107_ = default__.fm19(d_543_v15_, globalState)
            rhs108_ = d_528_v1_
            lhs93_ = globalState
            lhs94_ = globalState
            lhs95_ = d_531_v4_
            lhs96_ = default__.safeIndex(822, (d_531_v4_).length(0))
            lhs97_ = d_541_v13_
            lhs98_ = default__.safeIndex(115, (d_541_v13_).length(0))
            lhs99_ = d_542_v14_
            lhs100_ = default__.safeIndex(365, (d_542_v14_).length(0))
            lhs93_.f14 = rhs104_
            lhs94_.f18 = rhs105_
            lhs95_[lhs96_] = rhs106_
            lhs97_[lhs98_] = rhs107_
            lhs99_[lhs100_] = rhs108_
        elif True:
            d_544_v16_: _dafny.Seq
            d_544_v16_ = _dafny.SeqWithoutIsStrInference([len(self.f30)])
            d_545_v17_: _dafny.Set
            d_545_v17_ = _dafny.Set({d_544_v16_, (d_544_v16_ if (self).f29 else _dafny.SeqWithoutIsStrInference([-539, p2])), d_544_v16_, d_544_v16_})
            d_545_v17_ = ((d_545_v17_) - (d_545_v17_)).intersection(default__.fm25((self).f29, p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lyyloxwhs")), globalState))
            (globalState).f1 = (self).f38
            d_546_v18_: _dafny.Array
            nw93_ = _dafny.Array(None, 17)
            nw93_[int(0)] = True
            nw93_[int(1)] = (self).f29
            nw93_[int(2)] = p0
            nw93_[int(3)] = (self).f29
            nw93_[int(4)] = p0
            nw93_[int(5)] = p0
            nw93_[int(6)] = (self).f29
            nw93_[int(7)] = (self).f29
            nw93_[int(8)] = (True if (self).f29 else (self).f29)
            nw93_[int(9)] = (p2) == (p2)
            nw93_[int(10)] = (self).f29
            nw93_[int(11)] = p0
            nw93_[int(12)] = (self).f29
            nw93_[int(13)] = (self).f29
            nw93_[int(14)] = p0
            nw93_[int(15)] = p0
            nw93_[int(16)] = p0
            d_546_v18_ = nw93_
            index91_ = default__.safeIndex(2, (d_546_v18_).length(0))
            (d_546_v18_)[index91_] = (self).f29
            d_547_v19_: str
            d_547_v19_ = _dafny.CodePoint('x')
            index92_ = default__.safeIndex(2, (d_546_v18_).length(0))
            rhs109_ = (0) - ((d_544_v16_)[default__.safeIndex(p2, len(d_544_v16_))])
            rhs110_ = (self).f38
            rhs111_ = ((p2 if (self).f29 else (self).f38)) >= (((0) - (p2)) * ((self).f38))
            rhs112_ = p0
            rhs113_ = d_547_v19_
            lhs101_ = globalState
            lhs102_ = globalState
            lhs103_ = d_546_v18_
            lhs104_ = default__.safeIndex(2, (d_546_v18_).length(0))
            lhs101_.f14 = rhs109_
            lhs102_.f14 = rhs110_
            r1 = rhs111_
            lhs103_[lhs104_] = rhs112_
            d_547_v19_ = rhs113_
            d_548_v20_: D3
            d_548_v20_ = D3_DC11((self).f29, (self).f29)
            d_549_v21_: _dafny.Map
            d_549_v21_ = _dafny.Map({d_548_v20_: d_547_v19_})
            (globalState).f19 = (default__.safeModuloInt(p2, (0) - ((self).f38))) + (len(d_549_v21_))
            index93_ = default__.safeIndex(2, (d_546_v18_).length(0))
            (d_546_v18_)[index93_] = (self).f29
        d_550_v22_: _dafny.Array
        def lambda21_(d_551_p2_):
            def lambda22_(d_552_i3_):
                return not((d_551_p2_) < (d_551_p2_))

            return lambda22_

        init12_ = lambda21_(p2)
        nw94_ = _dafny.Array(None, 11)
        for i0_12_ in range(nw94_.length(0)):
            nw94_[i0_12_] = init12_(i0_12_)
        d_550_v22_ = nw94_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_550_v22_).length(0)):
            d_553_i2_: int = guard_loop_3_
            if (True) and (((0) <= (d_553_i2_)) and ((d_553_i2_) < ((d_550_v22_).length(0)))):
                (d_550_v22_)[(d_553_i2_)] = (((_dafny.MultiSet([(D2_DC9(self.f30, (self).f38, p0, (self).f38)).cf16, (self).f29])).cardinality) + ((0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0, p0]))).cardinality))) < (474)
        if p0:
            (globalState).f20 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwxebygro"))) + (self.f30)
            d_554_v23_: _dafny.Set
            d_554_v23_ = _dafny.Set({(self).f38})
            d_555_v24_: _dafny.MultiSet
            d_555_v24_ = _dafny.MultiSet([(self).f38])
            def iife43_():
                coll27_ = _dafny.Set()
                compr_27_: int
                for compr_27_ in (d_555_v24_).Elements:
                    d_556_v25_: int = compr_27_
                    if (d_556_v25_) in (d_555_v24_):
                        coll27_ = coll27_.union(_dafny.Set([(d_556_v25_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsl"))))]))
                return _dafny.Set(coll27_)
            rhs114_ = _dafny.Set({(self).f38})
            rhs115_ = (self).f29
            rhs116_ = (d_554_v23_) | ((d_554_v23_) | (iife43_()
            ))
            lhs105_ = globalState
            d_554_v23_ = rhs114_
            lhs105_.f28 = rhs115_
            d_554_v23_ = rhs116_
            d_557_v26_: C0
            nw95_ = C0()
            nw95_.ctor__(not((self).f29))
            d_557_v26_ = nw95_
            d_558_v27_: C0
            nw96_ = C0()
            nw96_.ctor__(d_557_v26_.f37)
            d_558_v27_ = nw96_
            d_559_v28_: _dafny.Set
            d_560_v29_: int
            out10_: _dafny.Set
            out11_: int
            out10_, out11_ = (self).m8(d_550_v22_, globalState)
            d_559_v28_ = out10_
            d_560_v29_ = out11_
        elif True:
            (globalState).f19 = p2
            d_561_v30_: _dafny.Seq
            d_561_v30_ = _dafny.SeqWithoutIsStrInference([(self).f29])
            (globalState).f14 = (0) - (default__.safeDivisionInt((self).f38, len(d_561_v30_)))
            (globalState).f13 = (self).f29
            d_562_v31_: _dafny.Seq
            d_562_v31_ = _dafny.SeqWithoutIsStrInference([d_550_v22_, d_550_v22_, d_550_v22_])
            d_563_v32_: _dafny.Set
            d_563_v32_ = _dafny.Set({d_550_v22_, (d_562_v31_)[default__.safeIndex(p2, len(d_562_v31_))], d_550_v22_, d_550_v22_})
            (globalState).f13 = ((_dafny.Set({d_550_v22_})).isdisjoint(d_563_v32_)) == (p0)
            (globalState).f14 = p2
        r0 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trbmjfc"))) + (self.f30)
        r1 = (self).f29
        return r0, r1

    def m8(self, p0, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: int = int(0)
        d_564_v0_: _dafny.Seq
        d_564_v0_ = _dafny.SeqWithoutIsStrInference([(self).f29])
        d_565_v1_: _dafny.Seq
        d_565_v1_ = _dafny.SeqWithoutIsStrInference([len(d_564_v0_), (self).f38])
        d_566_v2_: _dafny.Map
        d_566_v2_ = _dafny.Map({(self).f38: ((self).fm16(_dafny.MultiSet([(self).f29]), globalState)) - ((d_565_v1_)[default__.safeIndex((self).f38, len(d_565_v1_))])})
        d_566_v2_ = (d_566_v2_).set((self).f38, (self).f38)
        d_567_v3_: _dafny.Array
        def lambda23_(d_568_i1_):
            return (self).f29

        init13_ = lambda23_
        nw97_ = _dafny.Array(None, 16)
        for i0_13_ in range(nw97_.length(0)):
            nw97_[i0_13_] = init13_(i0_13_)
        d_567_v3_ = nw97_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_567_v3_).length(0)):
            d_569_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_569_i0_)) and ((d_569_i0_) < ((d_567_v3_).length(0)))):
                (d_567_v3_)[(d_569_i0_)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_570_i2_ in range(default__.abs(269))])) == (self.f30)
        (globalState).f13 = (self).f29
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_567_v3_).length(0)):
            d_571_i3_: int = guard_loop_5_
            if (True) and (((0) <= (d_571_i3_)) and ((d_571_i3_) < ((d_567_v3_).length(0)))):
                (d_567_v3_)[(d_571_i3_)] = ((self).f38) >= (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpqtgpu"))))
        d_572_v4_: _dafny.MultiSet
        d_572_v4_ = _dafny.MultiSet([False])
        hi1_ = (self).f38
        for d_573_i4_ in range(default__.safeModuloInt((self).fm16(d_572_v4_, globalState), (self).f38), hi1_):
            d_574_v5_: _dafny.Set
            d_574_v5_ = _dafny.Set({d_572_v4_, (d_572_v4_).set((self).f29, default__.abs((self).f38))})
            (globalState).f7 = (((d_565_v1_) + (d_565_v1_)).set(default__.safeIndex((0) - ((self).f38), len((d_565_v1_) + (d_565_v1_))), len(d_574_v5_))) + (d_565_v1_)
            d_575_v6_: _dafny.MultiSet
            d_575_v6_ = _dafny.MultiSet([(self).f38])
            d_576_v7_: D3
            d_576_v7_ = D3_DC10(d_575_v6_)
            d_577_v8_: D3
            d_577_v8_ = D3_DC12(d_576_v7_)
            source9_ = d_577_v8_
            if source9_.is_DC11:
                d_578___mcc_h0_ = source9_.cf19
                d_579___mcc_h1_ = source9_.cf20
                d_580_cf20_ = d_579___mcc_h1_
                d_581_cf19_ = d_578___mcc_h0_
                d_582_v9_: _dafny.Map
                d_582_v9_ = _dafny.Map({(self).f38: _dafny.Set({(self).f38})})
                d_583_v10_: _dafny.Set
                d_583_v10_ = _dafny.Set({p0})
                d_582_v9_ = (d_582_v9_).set((0) - (d_573_i4_), _dafny.Set({d_573_i4_, len(self.f30), (self).f38, len(d_583_v10_)}))
                d_584_v11_: C0
                nw98_ = C0()
                nw98_.ctor__((self).f29)
                d_584_v11_ = nw98_
                d_566_v2_ = (_dafny.Map({-507: (self).f38})) | (d_566_v2_)
                d_585_v12_: _dafny.Map
                d_585_v12_ = _dafny.Map({d_584_v11_.f37: d_573_i4_})
                d_586_v13_: _dafny.Array
                nw99_ = _dafny.Array(None, 26)
                nw99_[int(0)] = len((d_564_v0_).set(default__.safeIndex(d_573_i4_, len(d_564_v0_)), (self).f29))
                nw99_[int(1)] = (self).f38
                nw99_[int(2)] = (_dafny.MultiSet(d_564_v0_)).cardinality
                nw99_[int(3)] = (self).f38
                nw99_[int(4)] = (d_584_v11_).fm12(globalState)
                nw99_[int(5)] = (self).f38
                nw99_[int(6)] = (self).f38
                nw99_[int(7)] = (self).f38
                nw99_[int(8)] = len(self.f30)
                nw99_[int(9)] = (self).f38
                nw99_[int(10)] = d_573_i4_
                nw99_[int(11)] = (self).f38
                nw99_[int(12)] = 670
                nw99_[int(13)] = (self).f38
                nw99_[int(14)] = (0) - ((self).f38)
                nw99_[int(15)] = ((d_585_v12_)[False] if (False) in (d_585_v12_) else (self).f38)
                nw99_[int(16)] = (0) - ((self).f38)
                nw99_[int(17)] = d_573_i4_
                nw99_[int(18)] = (self).fm16(d_572_v4_, globalState)
                nw99_[int(19)] = (self).f38
                nw99_[int(20)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_587_i5_ in range(default__.abs(698))]))
                nw99_[int(21)] = len(self.f30)
                nw99_[int(22)] = (self).f38
                nw99_[int(23)] = 329
                nw99_[int(24)] = (self).f38
                nw99_[int(25)] = d_573_i4_
                d_586_v13_ = nw99_
                d_588_v14_: _dafny.Map
                d_588_v14_ = _dafny.Map({d_586_v13_: not(d_584_v11_.f37)})
                (globalState).f15 = ((d_588_v14_)[d_586_v13_] if (d_586_v13_) in (d_588_v14_) else d_580_cf20_)
            elif source9_.is_DC10:
                d_589___mcc_h2_ = source9_.cf18
                d_590_cf18_ = d_589___mcc_h2_
                d_591_v15_: _dafny.Array
                nw100_ = _dafny.Array(int(0), 8)
                d_591_v15_ = nw100_
                d_592_v16_: _dafny.Map
                d_592_v16_ = _dafny.Map({(self).fm16(d_572_v4_, globalState): d_591_v15_})
                d_592_v16_ = (d_592_v16_).set(d_573_i4_, d_591_v15_)
                index94_ = default__.safeIndex(417, (d_567_v3_).length(0))
                (d_567_v3_)[index94_] = (self).f29
                d_593_v17_: str
                d_593_v17_ = _dafny.CodePoint('k')
                d_594_v18_: _dafny.Map
                d_594_v18_ = _dafny.Map({_dafny.CodePoint('w'): (self).f38})
                d_595_v19_: _dafny.Map
                d_595_v19_ = _dafny.Map({False: d_594_v18_})
                index95_ = default__.safeIndex(417, (d_567_v3_).length(0))
                (d_567_v3_)[index95_] = ((self).f29 if default__.fm19(d_593_v17_, globalState) else (((d_595_v19_)[(self).f29] if ((self).f29) in (d_595_v19_) else default__.fm26((self).f29, (self).f38, (self).f29, globalState))) != (d_594_v18_))
                (globalState).f21 = d_573_i4_
                d_596_v20_: D3
                d_596_v20_ = D3_DC10((d_575_v6_).set(d_573_i4_, default__.abs((self).f38)))
                d_597_v22_: _dafny.Map
                d_597_v22_ = _dafny.Map({(0) - (d_573_i4_): (self).f29})
                d_598_v23_: _dafny.Array
                nw101_ = _dafny.Array(None, 29)
                nw101_[int(0)] = _dafny.MultiSet(d_565_v1_)
                nw101_[int(1)] = (d_596_v20_).cf18
                nw101_[int(2)] = d_590_cf18_
                nw101_[int(3)] = d_590_cf18_
                nw101_[int(4)] = _dafny.MultiSet((d_565_v1_) + (d_565_v1_))
                nw101_[int(5)] = default__.fm27((self).f38, False, globalState)
                nw101_[int(6)] = d_575_v6_
                nw101_[int(7)] = d_590_cf18_
                nw101_[int(8)] = _dafny.MultiSet(d_565_v1_)
                nw101_[int(9)] = d_590_cf18_
                nw101_[int(10)] = d_575_v6_
                nw101_[int(11)] = _dafny.MultiSet([d_573_i4_, d_573_i4_])
                nw101_[int(12)] = (d_590_cf18_).set((self).f38, default__.abs(d_573_i4_))
                nw101_[int(13)] = (_dafny.MultiSet([(self).f38, d_573_i4_, d_573_i4_]) if (d_567_v3_)[default__.safeIndex(417, (d_567_v3_).length(0))] else d_590_cf18_)
                nw101_[int(14)] = (d_590_cf18_).set(d_573_i4_, default__.abs(d_573_i4_))
                nw101_[int(15)] = (d_575_v6_).intersection(d_575_v6_)
                nw101_[int(16)] = d_590_cf18_
                nw101_[int(17)] = (d_590_cf18_) - (d_575_v6_)
                nw101_[int(18)] = d_590_cf18_
                nw101_[int(19)] = default__.fm27((self).f38, (d_567_v3_)[default__.safeIndex(417, (d_567_v3_).length(0))], globalState)
                def iife44_():
                    coll28_ = _dafny.Map()
                    compr_28_: int
                    for compr_28_ in (d_597_v22_).keys.Elements:
                        d_599_v21_: int = compr_28_
                        if (d_599_v21_) in (d_597_v22_):
                            coll28_[default__.safeModuloInt(d_599_v21_, (self).f38)] = (d_590_cf18_).cardinality
                    return _dafny.Map(coll28_)
                nw101_[int(20)] = _dafny.MultiSet([len(iife44_()
                )])
                nw101_[int(21)] = (d_575_v6_) - ((d_596_v20_).cf18)
                nw101_[int(22)] = d_590_cf18_
                nw101_[int(23)] = d_590_cf18_
                nw101_[int(24)] = d_590_cf18_
                nw101_[int(25)] = _dafny.MultiSet([79])
                nw101_[int(26)] = (d_575_v6_).intersection(d_575_v6_)
                nw101_[int(27)] = _dafny.MultiSet([d_573_i4_])
                nw101_[int(28)] = _dafny.MultiSet([(self).f38, (self).f38])
                d_598_v23_ = nw101_
                index96_ = default__.safeIndex(861, (d_598_v23_).length(0))
                (d_598_v23_)[index96_] = (_dafny.MultiSet([d_573_i4_])).set(d_573_i4_, default__.abs((self).f38))
                index97_ = default__.safeIndex(861, (d_598_v23_).length(0))
                rhs117_ = d_573_i4_
                rhs118_ = (d_575_v6_) | (d_590_cf18_)
                rhs119_ = (self.f30).set(default__.safeIndex((self).f38, len(self.f30)), d_593_v17_)
                lhs106_ = globalState
                lhs107_ = d_598_v23_
                lhs108_ = default__.safeIndex(861, (d_598_v23_).length(0))
                lhs109_ = globalState
                lhs106_.f21 = rhs117_
                lhs107_[lhs108_] = rhs118_
                lhs109_.f11 = rhs119_
            elif True:
                d_600___mcc_h3_ = source9_.cf21
                d_601_cf21_ = d_600___mcc_h3_
                index98_ = default__.safeIndex(429, (d_567_v3_).length(0))
                (d_567_v3_)[index98_] = (self).f29
                index99_ = default__.safeIndex(429, (d_567_v3_).length(0))
                (d_567_v3_)[index99_] = ((not((self).f29)) in (d_572_v4_)) in ((_dafny.SeqWithoutIsStrInference([(self).f29, (self).f29])) + (d_564_v0_))
                (globalState).f13 = ((self).fm16(d_572_v4_, globalState)) == (d_573_i4_)
                (globalState).f28 = (d_567_v3_)[default__.safeIndex(429, (d_567_v3_).length(0))]
                d_602_v24_: _dafny.Array
                nw102_ = _dafny.Array(False, 24)
                d_602_v24_ = nw102_
                d_603_v25_: _dafny.Array
                nw103_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_603_v25_ = nw103_
                d_604_v26_: _dafny.Map
                d_604_v26_ = _dafny.Map({(self).f29: (d_567_v3_)[default__.safeIndex(429, (d_567_v3_).length(0))]})
                d_605_v27_: str
                d_605_v27_ = _dafny.CodePoint('v')
                d_606_v28_: _dafny.Map
                d_606_v28_ = d_604_v26_
                d_607_v29_: _dafny.Array
                nw104_ = _dafny.Array(None, 27)
                nw104_[int(0)] = d_604_v26_
                nw104_[int(1)] = default__.fm28((self).f29, d_573_i4_, (self).f38, (self).f29, globalState)
                nw104_[int(2)] = d_604_v26_
                nw104_[int(3)] = d_604_v26_
                nw104_[int(4)] = d_604_v26_
                nw104_[int(5)] = d_604_v26_
                nw104_[int(6)] = _dafny.Map({not((self).f29): (d_567_v3_)[default__.safeIndex(429, (d_567_v3_).length(0))]})
                nw104_[int(7)] = d_604_v26_
                nw104_[int(8)] = _dafny.Map({(self).f29: default__.fm19(d_605_v27_, globalState)})
                nw104_[int(9)] = d_604_v26_
                nw104_[int(10)] = d_604_v26_
                nw104_[int(11)] = d_604_v26_
                nw104_[int(12)] = d_604_v26_
                nw104_[int(13)] = d_604_v26_
                nw104_[int(14)] = d_604_v26_
                nw104_[int(15)] = d_604_v26_
                nw104_[int(16)] = d_604_v26_
                nw104_[int(17)] = (d_606_v28_)
                nw104_[int(18)] = d_604_v26_
                nw104_[int(19)] = d_604_v26_
                nw104_[int(20)] = d_604_v26_
                nw104_[int(21)] = d_604_v26_
                nw104_[int(22)] = d_604_v26_
                nw104_[int(23)] = d_604_v26_
                nw104_[int(24)] = default__.fm28((d_567_v3_)[default__.safeIndex(429, (d_567_v3_).length(0))], d_573_i4_, (self).f38, (self).f29, globalState)
                nw104_[int(25)] = default__.fm28((d_567_v3_)[default__.safeIndex(429, (d_567_v3_).length(0))], (self).f38, (self).f38, (self).f29, globalState)
                nw104_[int(26)] = d_604_v26_
                d_607_v29_ = nw104_
                index100_ = default__.safeIndex(199, (d_603_v25_).length(0))
                (d_603_v25_)[index100_] = d_607_v29_
                d_608_v30_: _dafny.Array
                nw105_ = _dafny.Array(int(0), 1)
                d_608_v30_ = nw105_
                index101_ = default__.safeIndex(389, (d_608_v30_).length(0))
                (d_608_v30_)[index101_] = (0) - (d_573_i4_)
                index102_ = default__.safeIndex(199, (d_603_v25_).length(0))
                index103_ = default__.safeIndex(389, (d_608_v30_).length(0))
                rhs120_ = d_602_v24_
                rhs121_ = d_607_v29_
                rhs122_ = (d_573_i4_) * ((self).f38)
                lhs110_ = d_603_v25_
                lhs111_ = default__.safeIndex(199, (d_603_v25_).length(0))
                lhs112_ = d_608_v30_
                lhs113_ = default__.safeIndex(389, (d_608_v30_).length(0))
                d_602_v24_ = rhs120_
                lhs110_[lhs111_] = rhs121_
                lhs112_[lhs113_] = rhs122_
            (globalState).f13 = not (True) or ((self).f29)
            (globalState).f14 = (self).f38
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_567_v3_).length(0)):
            d_609_i6_: int = guard_loop_6_
            if (True) and (((0) <= (d_609_i6_)) and ((d_609_i6_) < ((d_567_v3_).length(0)))):
                (d_567_v3_)[(d_609_i6_)] = ((self).f38) <= (len((self.f30) + (self.f30)))
        d_610_v34_: _dafny.Map
        d_610_v34_ = _dafny.Map({False: (self).f29})
        d_611_v35_: str
        d_611_v35_ = _dafny.CodePoint('q')
        d_612_v36_: _dafny.MultiSet
        d_612_v36_ = _dafny.MultiSet([d_611_v35_])
        d_613_v37_: _dafny.Set
        d_613_v37_ = _dafny.Set({len(d_610_v34_), (d_612_v36_).cardinality, (self).f38, 859})
        d_614_v38_: _dafny.Map
        d_614_v38_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_611_v35_ for d_615_i7_ in range(default__.abs(419))])): _dafny.Set({(self).f38})})
        def iife45_():
            coll29_ = _dafny.Set()
            compr_29_: int
            for compr_29_ in (d_565_v1_).Elements:
                d_616_v31_: int = compr_29_
                if (d_616_v31_) in (d_565_v1_):
                    def iife46_():
                        coll30_ = _dafny.Set()
                        compr_30_: str
                        for compr_30_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')])).Elements:
                            d_617_v32_: str = compr_30_
                            if (d_617_v32_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')])):
                                coll30_ = coll30_.union(_dafny.Set([d_617_v32_]))
                        return _dafny.Set(coll30_)
                    def iife47_():
                        coll31_ = _dafny.Map()
                        compr_31_: int
                        for compr_31_ in _dafny.IntegerRange(-483, -255):
                            d_618_v33_: int = compr_31_
                            if ((-483) <= (d_618_v33_)) and ((d_618_v33_) < (-255)):
                                coll31_[default__.safeModuloInt(d_618_v33_, len(_dafny.Set({-81})))] = True
                        return _dafny.Map(coll31_)
                    coll29_ = coll29_.union(_dafny.Set([(d_616_v31_) * ((_dafny.MultiSet([_dafny.Map({False: -479}), _dafny.Map({not(False): len(iife46_()
)}), _dafny.Map({False: 317}), _dafny.Map({True: len(iife47_()
)})])).cardinality)]))
            return _dafny.Set(coll29_)
        r0 = (iife45_()
        ).intersection((d_613_v37_).intersection(((d_614_v38_)[(self).f38] if ((self).f38) in (d_614_v38_) else d_613_v37_)))
        d_619_v39_: D3
        d_619_v39_ = D3_DC12(default__.fm29((self).f29, (self).f38, (self).f29, globalState))
        d_620_v40_: D3
        d_620_v40_ = D3_DC11((d_564_v0_)[default__.safeIndex(587, len(d_564_v0_))], (self).f29)
        pat_let_tv15_ = d_620_v40_
        def lambda24_(source10_):
            if source10_.is_DC11:
                d_621___mcc_h4_ = source10_.cf19
                d_622___mcc_h5_ = source10_.cf20
                d_623_cf20_ = d_622___mcc_h5_
                d_624_cf19_ = d_621___mcc_h4_
                return 967
            elif source10_.is_DC10:
                d_625___mcc_h6_ = source10_.cf18
                d_626_cf18_ = d_625___mcc_h6_
                return (self).f38
            elif True:
                d_627___mcc_h7_ = source10_.cf21
                d_628_cf21_ = d_627___mcc_h7_
                return (self).f38

        def iife48_(_pat_let8_0):
            def iife49_(d_629_dt__update__tmp_h0_):
                def iife50_(_pat_let9_0):
                    def iife51_(d_630_dt__update_hcf21_h0_):
                        return D3_DC12(d_630_dt__update_hcf21_h0_)
                    return iife51_(_pat_let9_0)
                return iife50_(pat_let_tv15_)
            return iife49_(_pat_let8_0)
        r1 = lambda24_(iife48_(d_619_v39_))
        return r0, r1

    def m9(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Set = _dafny.Set({})
        d_631_v0_: _dafny.MultiSet
        d_631_v0_ = _dafny.MultiSet([p1, (0) - ((self).f38), p1])
        d_632_v1_: _dafny.Map
        d_632_v1_ = _dafny.Map({False: (self).f38})
        d_633_v2_: _dafny.Seq
        d_633_v2_ = _dafny.SeqWithoutIsStrInference([(self).f38])
        d_634_v3_: _dafny.Map
        d_634_v3_ = _dafny.Map({d_632_v1_: (_dafny.MultiSet(d_633_v2_)).set(len(self.f30), default__.abs(p1))})
        d_631_v0_ = ((d_634_v3_)[d_632_v1_] if (d_632_v1_) in (d_634_v3_) else (d_631_v0_) - (d_631_v0_))
        (globalState).f13 = (self).f29
        (globalState).f13 = p2
        (globalState).f28 = not ((self).f29) or (p2)
        d_635_v4_: D0
        d_635_v4_ = D0_DC0((p3 if (self).f29 else not((self).f29)))
        d_635_v4_ = (d_635_v4_ if (self).f29 else d_635_v4_)
        d_636_v5_: _dafny.Array
        def lambda25_(d_637_p2_):
            def lambda26_(d_638_i0_):
                return d_637_p2_

            return lambda26_

        init14_ = lambda25_(p2)
        nw106_ = _dafny.Array(None, 12)
        for i0_14_ in range(nw106_.length(0)):
            nw106_[i0_14_] = init14_(i0_14_)
        d_636_v5_ = nw106_
        index104_ = default__.safeIndex(162, (d_636_v5_).length(0))
        (d_636_v5_)[index104_] = False
        d_639_v6_: str
        d_639_v6_ = _dafny.CodePoint('q')
        d_640_v7_: _dafny.Set
        d_640_v7_ = _dafny.Set({d_639_v6_})
        pat_let_tv16_ = p2
        pat_let_tv17_ = p3
        pat_let_tv18_ = p3
        index105_ = default__.safeIndex(162, (d_636_v5_).length(0))
        def lambda27_(source11_):
            if source11_.is_DC7:
                d_641___mcc_h0_ = source11_.cf11
                d_642_cf11_ = d_641___mcc_h0_
                return (pat_let_tv16_) and (d_642_cf11_)
            elif source11_.is_DC8:
                d_643___mcc_h1_ = source11_.cf12
                d_644___mcc_h2_ = source11_.cf13
                d_645_cf13_ = d_644___mcc_h2_
                d_646_cf12_ = d_643___mcc_h1_
                return d_645_cf13_
            elif source11_.is_DC9:
                d_647___mcc_h3_ = source11_.cf14
                d_648___mcc_h4_ = source11_.cf15
                d_649___mcc_h5_ = source11_.cf16
                d_650___mcc_h6_ = source11_.cf17
                d_651_cf17_ = d_650___mcc_h6_
                d_652_cf16_ = d_649___mcc_h5_
                d_653_cf15_ = d_648___mcc_h4_
                d_654_cf14_ = d_647___mcc_h3_
                return pat_let_tv17_
            elif True:
                d_655___mcc_h7_ = source11_.cf10
                d_656_cf10_ = d_655___mcc_h7_
                return not(pat_let_tv18_)

        def iife52_(_pat_let10_0):
            def iife53_(d_658_dt__update__tmp_h0_):
                def iife54_(_pat_let11_0):
                    def iife55_(d_659_dt__update_hcf14_h0_):
                        def iife56_(_pat_let12_0):
                            def iife57_(d_660_dt__update_hcf16_h0_):
                                def iife58_(_pat_let13_0):
                                    def iife59_(d_661_dt__update_hcf15_h0_):
                                        return D2_DC9(d_659_dt__update_hcf14_h0_, d_661_dt__update_hcf15_h0_, d_660_dt__update_hcf16_h0_, (d_658_dt__update__tmp_h0_).cf17)
                                    return iife59_(_pat_let13_0)
                                return iife58_(271)
                            return iife57_(_pat_let12_0)
                        return iife56_(True)
                    return iife55_(_pat_let11_0)
                return iife54_(self.f30)
            return iife53_(_pat_let10_0)
        (d_636_v5_)[index105_] = lambda27_(iife52_(D2_DC9(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_657_i1_ in range(default__.abs(716))]), (0) - (len(d_640_v7_)), p3, p1)))
        d_662_v8_: T0
        nw107_ = C1()
        nw107_.ctor__(p2)
        d_662_v8_ = nw107_
        d_663_v9_: D10
        d_663_v9_ = D10_DC30(d_662_v8_)
        d_664_v10_: _dafny.Set
        d_664_v10_ = _dafny.Set({d_662_v8_, (d_663_v9_).cf49, d_662_v8_})
        d_665_v11_: _dafny.Set
        d_665_v11_ = _dafny.Set({(d_664_v10_).intersection(d_664_v10_)})
        r0 = d_665_v11_
        return r0

    @property
    def f38(self):
        return self._f38
    @property
    def f39(self):
        return self._f39

class C3(T1, T0):
    def  __init__(self):
        self._f30: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f29: bool = False
        self._f35: bool = False
        self._f36: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f30(self):
        return self._f30
    @f30.setter
    def f30(self, value):
        self._f30 = value
    @property
    def f29(self):
        return self._f29
    def ctor__(self, f35, f36, f29, f30):
        (self)._f35 = f35
        (self)._f36 = f36
        (self)._f29 = f29
        (self).f30 = f30

    def m2(self, p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_666_v0_: _dafny.Array
        nw108_ = _dafny.Array(_dafny.Array(None, 0), 7)
        d_666_v0_ = nw108_
        d_667_v1_: _dafny.Array
        def lambda28_(d_668_i4_):
            return (self).f35

        init15_ = lambda28_
        nw109_ = _dafny.Array(None, 20)
        for i0_15_ in range(nw109_.length(0)):
            nw109_[i0_15_] = init15_(i0_15_)
        d_667_v1_ = nw109_
        d_669_v2_: _dafny.Map
        d_669_v2_ = _dafny.Map({(self).f35: d_667_v1_})
        d_670_v3_: str
        d_670_v3_ = _dafny.CodePoint('f')
        d_671_v4_: _dafny.Array
        nw110_ = _dafny.Array(None, 24)
        nw110_[int(0)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_672_i0_ in range(default__.abs(172))])
        nw110_[int(1)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_673_i1_ in range(default__.abs(395))])
        nw110_[int(2)] = self.f30
        nw110_[int(3)] = self.f30
        nw110_[int(4)] = self.f30
        nw110_[int(5)] = self.f30
        nw110_[int(6)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_674_i2_ in range(default__.abs(514))])
        nw110_[int(7)] = self.f30
        nw110_[int(8)] = self.f30
        nw110_[int(9)] = self.f30
        nw110_[int(10)] = self.f30
        nw110_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdiohjl"))
        nw110_[int(12)] = self.f30
        nw110_[int(13)] = self.f30
        nw110_[int(14)] = self.f30
        nw110_[int(15)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_675_i3_ in range(default__.abs(107))])
        nw110_[int(16)] = (self.f30).set(default__.safeIndex(len(d_669_v2_), len(self.f30)), d_670_v3_)
        nw110_[int(17)] = self.f30
        nw110_[int(18)] = self.f30
        nw110_[int(19)] = self.f30
        nw110_[int(20)] = self.f30
        nw110_[int(21)] = self.f30
        nw110_[int(22)] = (self.f30).set(default__.safeIndex((self).f36, len(self.f30)), d_670_v3_)
        nw110_[int(23)] = self.f30
        d_671_v4_ = nw110_
        index106_ = default__.safeIndex(560, (d_666_v0_).length(0))
        (d_666_v0_)[index106_] = d_671_v4_
        index107_ = default__.safeIndex(560, (d_666_v0_).length(0))
        (d_666_v0_)[index107_] = d_671_v4_
        d_676_v5_: _dafny.Seq
        d_676_v5_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        source12_ = D1_DC3((d_676_v5_) + (_dafny.SeqWithoutIsStrInference([True])))
        if source12_.is_DC4:
            d_677___mcc_h0_ = source12_.cf8
            d_678_cf8_ = d_677___mcc_h0_
            d_679_v6_: _dafny.Array
            def lambda29_(d_680_i5_):
                return (d_680_i5_) - ((self).f36)

            init16_ = lambda29_
            nw111_ = _dafny.Array(None, 7)
            for i0_16_ in range(nw111_.length(0)):
                nw111_[i0_16_] = init16_(i0_16_)
            d_679_v6_ = nw111_
            (globalState).f17 = d_679_v6_
            (globalState).f15 = p0
            d_681_v7_: _dafny.Map
            d_681_v7_ = _dafny.Map({(self).f36: d_670_v3_})
            d_682_v8_: D2
            d_682_v8_ = D2_DC8(((d_681_v7_)[(self).f36] if ((self).f36) in (d_681_v7_) else d_670_v3_), (self).f29)
            source13_ = d_682_v8_
            if source13_.is_DC7:
                d_683___mcc_h3_ = source13_.cf11
                d_684_cf11_ = d_683___mcc_h3_
                d_685_v9_: _dafny.Seq
                d_685_v9_ = _dafny.SeqWithoutIsStrInference([d_676_v5_])
                r0 = d_685_v9_
                d_686_v10_: _dafny.Array
                nw112_ = _dafny.Array(D3.default()(), 23)
                d_686_v10_ = nw112_
                d_687_v11_: _dafny.Map
                d_687_v11_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uogups"))) + (self.f30): d_686_v10_})
                d_687_v11_ = (d_687_v11_).set((self.f30) + ((self.f30).set(default__.safeIndex((self).f36, len(self.f30)), d_670_v3_)), d_686_v10_)
                r1 = d_684_cf11_
                d_688_v12_: _dafny.Set
                d_688_v12_ = _dafny.Set({d_684_cf11_, (self).f35})
                (globalState).f19 = len((d_688_v12_) - (d_688_v12_))
            elif source13_.is_DC8:
                d_689___mcc_h4_ = source13_.cf12
                d_690___mcc_h5_ = source13_.cf13
                d_691_cf13_ = d_690___mcc_h5_
                d_692_cf12_ = d_689___mcc_h4_
                d_693_v13_: D2
                d_693_v13_ = D2_DC7((self).f29)
                (globalState).f15 = (d_693_v13_).cf11
                d_694_v14_: _dafny.Map
                d_694_v14_ = _dafny.Map({22: d_667_v1_})
                d_694_v14_ = (d_694_v14_).set(len(d_676_v5_), (d_667_v1_ if (self).f29 else d_667_v1_))
                d_695_v15_: D0
                d_695_v15_ = D0_DC1(d_667_v1_, (self).f36, (self).f36)
                rhs123_ = default__.safeDivisionInt(((self).f36 if (self).f35 else default__.fm9(globalState)), (len(self.f30)) * ((self).f36))
                rhs124_ = (self).f36
                rhs125_ = (_dafny.SeqWithoutIsStrInference([(default__.fm10((self).f36, (self).f36, globalState)) == ((d_676_v5_)[default__.safeIndex((self).f36, len(d_676_v5_))])])).set(default__.safeIndex((0) - ((self).f36), len(_dafny.SeqWithoutIsStrInference([(default__.fm10((self).f36, (self).f36, globalState)) == ((d_676_v5_)[default__.safeIndex((self).f36, len(d_676_v5_))])]))), (self).f35)
                rhs126_ = (d_695_v15_).cf2
                rhs127_ = (d_693_v13_ if d_678_cf8_ else d_693_v13_)
                lhs114_ = globalState
                lhs115_ = globalState
                lhs116_ = globalState
                lhs114_.f1 = rhs123_
                lhs115_.f14 = rhs124_
                d_676_v5_ = rhs125_
                lhs116_.f21 = rhs126_
                d_693_v13_ = rhs127_
                (globalState).f1 = ((self).f36) + ((self).f36)
            elif source13_.is_DC9:
                d_696___mcc_h6_ = source13_.cf14
                d_697___mcc_h7_ = source13_.cf15
                d_698___mcc_h8_ = source13_.cf16
                d_699___mcc_h9_ = source13_.cf17
                d_700_cf17_ = d_699___mcc_h9_
                d_701_cf16_ = d_698___mcc_h8_
                d_702_cf15_ = d_697___mcc_h7_
                d_703_cf14_ = d_696___mcc_h6_
                d_670_v3_ = d_670_v3_
                (globalState).f2 = (0) - ((self).f36)
                d_704_v16_: _dafny.Array
                nw113_ = _dafny.Array(D1.default()(), 11)
                d_704_v16_ = nw113_
                d_705_v17_: D1
                d_705_v17_ = D1_DC3(d_676_v5_)
                d_706_v18_: D1
                d_706_v18_ = D1_DC5(d_705_v17_)
                index108_ = default__.safeIndex(329, (d_704_v16_).length(0))
                (d_704_v16_)[index108_] = d_706_v18_
                index109_ = default__.safeIndex(329, (d_704_v16_).length(0))
                (d_704_v16_)[index109_] = d_706_v18_
                d_707_v19_: _dafny.Map
                d_707_v19_ = _dafny.Map({d_679_v6_: (self).f36})
                (globalState).f2 = len(d_707_v19_)
            elif True:
                d_708___mcc_h10_ = source13_.cf10
                d_709_cf10_ = d_708___mcc_h10_
                (globalState).f1 = ((self).f36 if (self).f35 else ((self).f36) - ((self).f36))
                (globalState).f11 = self.f30
                d_710_v20_: C0
                nw114_ = C0()
                nw114_.ctor__(p0)
                d_710_v20_ = nw114_
                d_711_v21_: _dafny.Array
                def lambda30_(d_712_v3_):
                    def lambda31_(d_713_i6_):
                        return d_712_v3_

                    return lambda31_

                init17_ = lambda30_(d_670_v3_)
                nw115_ = _dafny.Array(None, 15)
                for i0_17_ in range(nw115_.length(0)):
                    nw115_[i0_17_] = init17_(i0_17_)
                d_711_v21_ = nw115_
                index110_ = default__.safeIndex(263, (d_711_v21_).length(0))
                (d_711_v21_)[index110_] = (self.f30)[default__.safeIndex((self).f36, len(self.f30))]
                index111_ = default__.safeIndex(263, (d_711_v21_).length(0))
                (d_711_v21_)[index111_] = d_670_v3_
            index112_ = default__.safeIndex(728, (d_679_v6_).length(0))
            (d_679_v6_)[index112_] = (self).f36
            index113_ = default__.safeIndex(728, (d_679_v6_).length(0))
            (d_679_v6_)[index113_] = len(self.f30)
        elif source12_.is_DC3:
            d_714___mcc_h1_ = source12_.cf7
            d_715_cf7_ = d_714___mcc_h1_
            d_716_v22_: _dafny.Map
            d_716_v22_ = _dafny.Map({d_715_cf7_: (0) - ((self).f36)})
            def iife60_():
                coll32_ = _dafny.Map()
                compr_32_: int
                for compr_32_ in _dafny.IntegerRange(-784, -112):
                    d_717_v23_: int = compr_32_
                    if ((-784) <= (d_717_v23_)) and ((d_717_v23_) < (-112)):
                        coll32_[default__.safeModuloInt(d_717_v23_, 471)] = (self).f36
                return _dafny.Map(coll32_)
            d_716_v22_ = ((d_716_v22_) | ((default__.fm13((self).f36, len(iife60_()
            ), (self).f36, not((self).f35), globalState)).cf22)) | (d_716_v22_)
            (globalState).f20 = self.f30
            (globalState).f19 = (self).f36
            d_718_v24_: D1
            d_718_v24_ = D1_DC4(not (p0) or (p0))
            source14_ = d_718_v24_
            if source14_.is_DC4:
                d_719___mcc_h11_ = source14_.cf8
                d_720_cf8_ = d_719___mcc_h11_
                d_721_v25_: _dafny.Array
                def lambda32_(d_722_v3_):
                    def lambda33_(d_723_i7_):
                        return d_722_v3_

                    return lambda33_

                init18_ = lambda32_(d_670_v3_)
                nw116_ = _dafny.Array(None, 6)
                for i0_18_ in range(nw116_.length(0)):
                    nw116_[i0_18_] = init18_(i0_18_)
                d_721_v25_ = nw116_
                index114_ = default__.safeIndex(615, (d_721_v25_).length(0))
                (d_721_v25_)[index114_] = d_670_v3_
                d_724_v26_: _dafny.MultiSet
                d_724_v26_ = _dafny.MultiSet([(self).f36])
                d_725_v27_: _dafny.Map
                d_725_v27_ = _dafny.Map({not((self).f35): d_724_v26_})
                d_726_v28_: D3
                d_726_v28_ = D3_DC10(((d_725_v27_)[d_720_cf8_] if (d_720_cf8_) in (d_725_v27_) else d_724_v26_))
                d_727_v29_: D3
                d_727_v29_ = D3_DC12(d_726_v28_)
                d_728_v30_: _dafny.Map
                d_728_v30_ = _dafny.Map({(774) <= ((self).f36): d_727_v29_})
                d_729_v31_: D2
                d_729_v31_ = D2_DC8(d_670_v3_, p0)
                d_730_v32_: _dafny.Array
                nw117_ = _dafny.Array(int(0), 9)
                d_730_v32_ = nw117_
                pat_let_tv19_ = d_726_v28_
                index115_ = default__.safeIndex(615, (d_721_v25_).length(0))
                def iife61_(_pat_let14_0):
                    def iife62_(d_731_dt__update__tmp_h0_):
                        def iife63_(_pat_let15_0):
                            def iife64_(d_732_dt__update_hcf21_h0_):
                                return D3_DC12(d_732_dt__update_hcf21_h0_)
                            return iife64_(_pat_let15_0)
                        return iife63_(pat_let_tv19_)
                    return iife62_(_pat_let14_0)
                rhs128_ = (d_729_v31_).cf12
                rhs129_ = (((d_728_v30_).set((self).f35, d_727_v29_)) | ((d_728_v30_).set(False, iife61_(d_727_v29_)))) | (d_728_v30_)
                rhs130_ = d_667_v1_
                rhs131_ = d_730_v32_
                lhs117_ = d_721_v25_
                lhs118_ = default__.safeIndex(615, (d_721_v25_).length(0))
                lhs119_ = globalState
                lhs117_[lhs118_] = rhs128_
                d_728_v30_ = rhs129_
                d_667_v1_ = rhs130_
                lhs119_.f17 = rhs131_
                d_733_v33_: C0
                nw118_ = C0()
                nw118_.ctor__((True if (self).f29 else p0))
                d_733_v33_ = nw118_
                d_734_v34_: _dafny.Map
                d_734_v34_ = _dafny.Map({d_720_cf8_: len(d_676_v5_)})
                d_735_v36_: _dafny.MultiSet
                d_735_v36_ = _dafny.MultiSet([_dafny.Map({d_733_v33_.f37: (self).f36}), d_734_v34_, default__.fm14((self).f35, _dafny.SeqWithoutIsStrInference([(self).f36 for d_736_i9_ in range(default__.abs(900))]), d_720_cf8_, len(_dafny.Set({d_667_v1_})), globalState), d_734_v34_])
                d_737_v38_: _dafny.Set
                d_737_v38_ = _dafny.Set({p0, d_733_v33_.f37})
                d_738_v39_: _dafny.Set
                d_738_v39_ = _dafny.Set({_dafny.Map({p0: len(d_737_v38_)})})
                d_739_v40_: _dafny.Map
                d_739_v40_ = _dafny.Map({d_734_v34_: True})
                d_740_v42_: _dafny.Map
                d_740_v42_ = _dafny.Map({d_734_v34_: (self).f36})
                d_741_v44_: _dafny.Array
                nw119_ = _dafny.Array(None, 9)
                nw119_[int(0)] = _dafny.Set({d_734_v34_, d_734_v34_})
                def iife65_():
                    coll33_ = _dafny.Set()
                    compr_33_: _dafny.Map
                    for compr_33_ in (_dafny.SeqWithoutIsStrInference([d_734_v34_ for d_742_i8_ in range(default__.abs(677))])).Elements:
                        d_743_v35_: _dafny.Map = compr_33_
                        if (d_743_v35_) in (_dafny.SeqWithoutIsStrInference([d_734_v34_ for d_742_i8_ in range(default__.abs(677))])):
                            coll33_ = coll33_.union(_dafny.Set([d_743_v35_]))
                    return _dafny.Set(coll33_)
                nw119_[int(1)] = iife65_()
                
                def iife66_():
                    coll34_ = _dafny.Set()
                    compr_34_: _dafny.Map
                    for compr_34_ in (d_735_v36_).Elements:
                        d_744_v37_: _dafny.Map = compr_34_
                        if (d_744_v37_) in (d_735_v36_):
                            coll34_ = coll34_.union(_dafny.Set([d_744_v37_]))
                    return _dafny.Set(coll34_)
                nw119_[int(2)] = iife66_()
                
                def iife67_():
                    coll35_ = _dafny.Set()
                    compr_35_: _dafny.Map
                    for compr_35_ in (d_739_v40_).keys.Elements:
                        d_745_v41_: _dafny.Map = compr_35_
                        if (d_745_v41_) in (d_739_v40_):
                            coll35_ = coll35_.union(_dafny.Set([d_745_v41_]))
                    return _dafny.Set(coll35_)
                nw119_[int(3)] = (d_738_v39_).intersection(iife67_()
                )
                nw119_[int(4)] = d_738_v39_
                def iife68_():
                    coll36_ = _dafny.Set()
                    compr_36_: _dafny.Map
                    for compr_36_ in ((d_740_v42_).set(d_734_v34_, (self).f36)).keys.Elements:
                        d_746_v43_: _dafny.Map = compr_36_
                        if (d_746_v43_) in ((d_740_v42_).set(d_734_v34_, (self).f36)):
                            coll36_ = coll36_.union(_dafny.Set([d_746_v43_]))
                    return _dafny.Set(coll36_)
                nw119_[int(5)] = iife68_()
                
                nw119_[int(6)] = d_738_v39_
                nw119_[int(7)] = _dafny.Set({d_734_v34_, d_734_v34_, (d_734_v34_).set(d_720_cf8_, len(d_737_v38_)), d_734_v34_, d_734_v34_})
                nw119_[int(8)] = (_dafny.Set({d_734_v34_})).intersection(d_738_v39_)
                d_741_v44_ = nw119_
                index116_ = default__.safeIndex(437, (d_741_v44_).length(0))
                (d_741_v44_)[index116_] = d_738_v39_
                index117_ = default__.safeIndex(437, (d_741_v44_).length(0))
                (d_741_v44_)[index117_] = default__.fm15((self).f36, globalState)
                (globalState).f19 = (0) - ((self).f36)
            elif source14_.is_DC3:
                d_747___mcc_h12_ = source14_.cf7
                d_748_cf7_ = d_747___mcc_h12_
                (globalState).f1 = 58
                d_749_v45_: _dafny.Array
                nw120_ = _dafny.Array(int(0), 20)
                d_749_v45_ = nw120_
                index118_ = default__.safeIndex(593, (d_749_v45_).length(0))
                (d_749_v45_)[index118_] = (self).f36
                index119_ = default__.safeIndex(593, (d_749_v45_).length(0))
                (d_749_v45_)[index119_] = 943
                (globalState).f21 = (d_749_v45_)[default__.safeIndex(593, (d_749_v45_).length(0))]
                (globalState).f28 = p0
            elif True:
                d_750___mcc_h13_ = source14_.cf9
                d_751_cf9_ = d_750___mcc_h13_
                d_752_v46_: _dafny.Set
                d_752_v46_ = _dafny.Set({p0, (self).f29})
                d_753_v47_: _dafny.Seq
                d_753_v47_ = _dafny.SeqWithoutIsStrInference([d_752_v46_])
                (globalState).f3 = (d_753_v47_) + (d_753_v47_)
                (globalState).f1 = (default__.fm9(globalState)) + ((self).f36)
                d_754_v48_: C0
                nw121_ = C0()
                nw121_.ctor__((False if (self).f35 else (self).f35))
                d_754_v48_ = nw121_
                index120_ = default__.safeIndex(42, (d_667_v1_).length(0))
                (d_667_v1_)[index120_] = (self).f29
                index121_ = default__.safeIndex(42, (d_667_v1_).length(0))
                (d_667_v1_)[index121_] = (self).f35
        elif True:
            d_755___mcc_h2_ = source12_.cf9
            d_756_cf9_ = d_755___mcc_h2_
            d_757_v49_: _dafny.Map
            d_757_v49_ = _dafny.Map({(self).f35: (self).f36})
            d_758_v50_: _dafny.Array
            nw122_ = _dafny.Array(None, 17)
            nw122_[int(0)] = (self).f36
            nw122_[int(1)] = len(self.f30)
            nw122_[int(2)] = (self).f36
            nw122_[int(3)] = (self).f36
            nw122_[int(4)] = (self).f36
            nw122_[int(5)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_759_i10_ in range(default__.abs(-598))]))
            nw122_[int(6)] = (self).f36
            nw122_[int(7)] = (self).f36
            nw122_[int(8)] = len(d_757_v49_)
            nw122_[int(9)] = (self).f36
            nw122_[int(10)] = len(d_676_v5_)
            nw122_[int(11)] = (self).f36
            nw122_[int(12)] = (self).f36
            nw122_[int(13)] = (self).f36
            nw122_[int(14)] = (self).f36
            nw122_[int(15)] = (self).f36
            nw122_[int(16)] = (self).f36
            d_758_v50_ = nw122_
            d_760_v51_: _dafny.Seq
            d_760_v51_ = _dafny.SeqWithoutIsStrInference([d_758_v50_, d_758_v50_])
            d_761_v52_: _dafny.Set
            d_761_v52_ = _dafny.Set({d_758_v50_, (d_760_v51_)[default__.safeIndex(default__.fm9(globalState), len(d_760_v51_))]})
            d_762_v53_: T1
            nw123_ = C2()
            nw123_.ctor__((self).f36, d_761_v52_, p0, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gt"))) + (self.f30))
            d_762_v53_ = nw123_
            d_762_v53_ = d_762_v53_
            d_763_v54_: _dafny.MultiSet
            d_763_v54_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vccvow"))])
            (globalState).f28 = (self.f30) in (d_763_v54_)
            d_764_v55_: _dafny.Array
            def lambda34_(d_765_v3_):
                def lambda35_(d_766_i11_):
                    return d_765_v3_

                return lambda35_

            init19_ = lambda34_(d_670_v3_)
            nw124_ = _dafny.Array(None, 17)
            for i0_19_ in range(nw124_.length(0)):
                nw124_[i0_19_] = init19_(i0_19_)
            d_764_v55_ = nw124_
            d_767_v56_: _dafny.Map
            d_767_v56_ = _dafny.Map({(self).f35: d_764_v55_})
            d_768_v57_: _dafny.MultiSet
            d_768_v57_ = _dafny.MultiSet([d_764_v55_, d_764_v55_, d_764_v55_, ((d_767_v56_)[p0] if (p0) in (d_767_v56_) else d_764_v55_)])
            d_768_v57_ = _dafny.MultiSet([d_764_v55_])
            (globalState).f21 = default__.safeDivisionInt((self).f36, 871)
        d_769_v58_: _dafny.Map
        d_769_v58_ = _dafny.Map({(self).f36: (self).f36})
        d_769_v58_ = (d_769_v58_).set(((self).f36) - ((self).f36), (self).f36)
        d_770_v59_: _dafny.Map
        d_770_v59_ = _dafny.Map({d_670_v3_: p0})
        d_771_v60_: _dafny.Set
        d_771_v60_ = _dafny.Set({d_770_v59_, d_770_v59_})
        d_772_v61_: _dafny.Map
        d_772_v61_ = _dafny.Map({(self).f35: (d_771_v60_) | (_dafny.Set({_dafny.Map({d_670_v3_: False}), d_770_v59_}))})
        d_772_v61_ = (d_772_v61_).set((self).f29, _dafny.Set({(d_770_v59_).set(_dafny.CodePoint('u'), (self).f35), d_770_v59_}))
        d_773_v62_: _dafny.MultiSet
        d_773_v62_ = _dafny.MultiSet([(self).f35])
        d_774_v63_: D0
        d_774_v63_ = D0_DC2(d_773_v62_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngrrpfx")), (self).f29)
        d_775_v64_: _dafny.Array
        nw125_ = _dafny.Array(None, 16)
        nw125_[int(0)] = ((d_774_v63_).cf4 if False else d_773_v62_)
        nw125_[int(1)] = ((d_773_v62_).set((self).f29, default__.abs((0) - ((self).f36)))) | (default__.fm39(d_670_v3_, (self).f36, globalState))
        nw125_[int(2)] = (d_773_v62_) - (d_773_v62_)
        nw125_[int(3)] = _dafny.MultiSet([p0])
        nw125_[int(4)] = (d_773_v62_).set(not((self).f29), default__.abs((self).f36))
        nw125_[int(5)] = _dafny.MultiSet([p0])
        nw125_[int(6)] = d_773_v62_
        nw125_[int(7)] = d_773_v62_
        nw125_[int(8)] = (d_774_v63_).cf4
        nw125_[int(9)] = _dafny.MultiSet([p0, True, p0, (self).f29, p0])
        nw125_[int(10)] = _dafny.MultiSet(d_676_v5_)
        nw125_[int(11)] = d_773_v62_
        nw125_[int(12)] = (d_774_v63_).cf4
        nw125_[int(13)] = (default__.fm39(d_670_v3_, (self).f36, globalState)).intersection((d_774_v63_).cf4)
        nw125_[int(14)] = d_773_v62_
        nw125_[int(15)] = _dafny.MultiSet([default__.fm19(_dafny.CodePoint('k'), globalState)])
        d_775_v64_ = nw125_
        d_775_v64_ = d_775_v64_
        d_776_v65_: C1
        nw126_ = C1()
        nw126_.ctor__((self).f35)
        d_776_v65_ = nw126_
        d_777_v66_: _dafny.Seq
        d_777_v66_ = _dafny.SeqWithoutIsStrInference([d_676_v5_, d_676_v5_])
        r0 = (((d_777_v66_) + (d_777_v66_)).set(default__.safeIndex((self).f36, len((d_777_v66_) + (d_777_v66_))), d_676_v5_)) + ((d_777_v66_) + (d_777_v66_))
        r1 = False
        return r0, r1

    def m3(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        d_778_v0_: _dafny.Seq
        d_778_v0_ = _dafny.SeqWithoutIsStrInference([(self).f29, (self).f29])
        if ((self).f35) not in (((d_778_v0_) + (default__.fm35(_dafny.SeqWithoutIsStrInference([D2_DC8(_dafny.CodePoint('o'), (self).f35) for d_779_i0_ in range(default__.abs(531))]), (0) - ((self).f36), globalState))).set(default__.safeIndex((self).f36, len((d_778_v0_) + (default__.fm35(_dafny.SeqWithoutIsStrInference([D2_DC8(_dafny.CodePoint('o'), (self).f35) for d_780_i0_ in range(default__.abs(531))]), (0) - ((self).f36), globalState)))), p0)):
            d_781_v1_: T0
            nw127_ = C1()
            nw127_.ctor__((self).f29)
            d_781_v1_ = nw127_
            d_782_v2_: D10
            d_782_v2_ = D10_DC30(d_781_v1_)
            (globalState).f21 = len(_dafny.SeqWithoutIsStrInference([d_782_v2_]))
            d_783_v3_: _dafny.Map
            d_783_v3_ = _dafny.Map({(self).f29: len(self.f30)})
            d_784_v4_: _dafny.Seq
            d_784_v4_ = _dafny.SeqWithoutIsStrInference([(0) - (p2), (self).f36])
            (globalState).f1 = (default__.fm9(globalState)) - (((d_783_v3_)[(self).f35] if ((self).f35) in (d_783_v3_) else (d_784_v4_)[default__.safeIndex(p2, len(d_784_v4_))]))
            d_785_v5_: _dafny.Seq
            d_786_v6_: bool
            out12_: _dafny.Seq
            out13_: bool
            out12_, out13_ = (d_781_v1_).m2(p0, globalState)
            d_785_v5_ = out12_
            d_786_v6_ = out13_
            d_787_v7_: _dafny.Array
            nw128_ = _dafny.Array(int(0), 25)
            d_787_v7_ = nw128_
            d_788_v8_: _dafny.Set
            d_788_v8_ = _dafny.Set({d_787_v7_})
            d_789_v9_: C2
            nw129_ = C2()
            nw129_.ctor__(p2, d_788_v8_, p0, default__.fm20((self).f29, not((self).f35), d_778_v0_, p2, globalState))
            d_789_v9_ = nw129_
            d_790_v10_: C1
            nw130_ = C1()
            nw130_.ctor__(True)
            d_790_v10_ = nw130_
        elif True:
            d_791_v11_: _dafny.Array
            nw131_ = _dafny.Array(_dafny.Seq({}), 8)
            d_791_v11_ = nw131_
            d_792_v12_: str
            d_792_v12_ = _dafny.CodePoint('n')
            d_793_v13_: _dafny.Seq
            d_793_v13_ = _dafny.SeqWithoutIsStrInference([(self.f30).set(default__.safeIndex(len(self.f30), len(self.f30)), d_792_v12_)])
            index122_ = default__.safeIndex(46, (d_791_v11_).length(0))
            (d_791_v11_)[index122_] = (d_793_v13_).set(default__.safeIndex(p2, len(d_793_v13_)), self.f30)
            d_794_v14_: _dafny.Array
            nw132_ = _dafny.Array(int(0), 8)
            d_794_v14_ = nw132_
            index123_ = default__.safeIndex(224, (d_794_v14_).length(0))
            (d_794_v14_)[index123_] = 62
            d_795_v15_: _dafny.MultiSet
            d_795_v15_ = _dafny.MultiSet([p0])
            d_796_v16_: _dafny.Seq
            d_796_v16_ = _dafny.SeqWithoutIsStrInference([d_795_v15_])
            d_797_v17_: _dafny.Array
            nw133_ = _dafny.Array(None, 7)
            nw133_[int(0)] = (self).f29
            nw133_[int(1)] = (self).f29
            nw133_[int(2)] = ((d_796_v16_)[default__.safeIndex(-495, len(d_796_v16_))]).ispropersubset((d_795_v15_).set(True, default__.abs(p2)))
            nw133_[int(3)] = (self).f35
            nw133_[int(4)] = p0
            nw133_[int(5)] = (self).f35
            nw133_[int(6)] = p0
            d_797_v17_ = nw133_
            d_798_v18_: _dafny.Set
            d_798_v18_ = _dafny.Set({p0, (self).f29})
            index124_ = default__.safeIndex(46, (d_791_v11_).length(0))
            index125_ = default__.safeIndex(224, (d_794_v14_).length(0))
            rhs132_ = ((d_793_v13_).set(default__.safeIndex((self).f36, len(d_793_v13_)), self.f30)) + (_dafny.SeqWithoutIsStrInference([self.f30 for d_799_i1_ in range(default__.abs(734))]))
            rhs133_ = (len((self.f30) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjkkiurju"))))) - (((self).f36 if (self).f35 else 970))
            rhs134_ = len((d_798_v18_ if False else (d_798_v18_ if p0 else _dafny.Set({p0, (self).f35}))))
            rhs135_ = d_797_v17_
            lhs120_ = d_791_v11_
            lhs121_ = default__.safeIndex(46, (d_791_v11_).length(0))
            lhs122_ = d_794_v14_
            lhs123_ = default__.safeIndex(224, (d_794_v14_).length(0))
            lhs124_ = globalState
            lhs120_[lhs121_] = rhs132_
            lhs122_[lhs123_] = rhs133_
            lhs124_.f14 = rhs134_
            d_797_v17_ = rhs135_
            d_800_v19_: _dafny.MultiSet
            d_800_v19_ = _dafny.MultiSet([(d_794_v14_)[default__.safeIndex(224, (d_794_v14_).length(0))]])
            (globalState).f13 = (d_800_v19_).ispropersubset(d_800_v19_)
            (globalState).f15 = ((self).f36) != (-139)
            d_801_v20_: _dafny.Array
            nw134_ = _dafny.Array(None, 4)
            nw134_[int(0)] = d_792_v12_
            nw134_[int(1)] = d_792_v12_
            nw134_[int(2)] = _dafny.CodePoint('r')
            nw134_[int(3)] = d_792_v12_
            d_801_v20_ = nw134_
            d_802_v21_: _dafny.Seq
            d_802_v21_ = _dafny.SeqWithoutIsStrInference([d_801_v20_, d_801_v20_])
            d_803_v22_: _dafny.Seq
            d_803_v22_ = _dafny.SeqWithoutIsStrInference([d_801_v20_, d_801_v20_, d_801_v20_, (d_802_v21_)[default__.safeIndex(p2, len(d_802_v21_))], d_801_v20_])
            d_801_v20_ = (d_803_v22_)[default__.safeIndex((p2) * ((0) - (len(d_778_v0_))), len(d_803_v22_))]
            rhs136_ = d_794_v14_
            rhs137_ = default__.fm33(((self).f35) and ((self).f29), (d_794_v14_)[default__.safeIndex(224, (d_794_v14_).length(0))], p2, globalState)
            lhs125_ = globalState
            lhs126_ = globalState
            lhs125_.f17 = rhs136_
            lhs126_.f13 = rhs137_
        d_804_v23_: _dafny.Map
        d_804_v23_ = _dafny.Map({(self).f36: (self).f36})
        d_805_v24_: str
        d_805_v24_ = _dafny.CodePoint('h')
        d_806_v25_: _dafny.Map
        d_806_v25_ = _dafny.Map({d_805_v24_: (self).f35})
        d_807_v26_: _dafny.Map
        d_807_v26_ = _dafny.Map({(self).f35: len(d_806_v25_)})
        d_808_v27_: D4
        d_808_v27_ = D4_DC15(((d_804_v23_)[-597] if (-597) in (d_804_v23_) else len(d_807_v26_)), (self).f36)
        def iife69_(_pat_let16_0):
            def iife70_(d_809_dt__update__tmp_h0_):
                def iife71_(_pat_let17_0):
                    def iife72_(d_810_dt__update_hcf25_h0_):
                        return D4_DC15((d_809_dt__update__tmp_h0_).cf24, d_810_dt__update_hcf25_h0_)
                    return iife72_(_pat_let17_0)
                return iife71_((self).f36)
            return iife70_(_pat_let16_0)
        d_808_v27_ = iife69_(D4_DC15((self).f36, len(self.f30)))
        if (self).f29:
            d_811_v28_: _dafny.Seq
            d_811_v28_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))), (self).f36])
            (globalState).f21 = len(((d_811_v28_) + (d_811_v28_)) + (_dafny.SeqWithoutIsStrInference([p2 for d_812_i2_ in range(default__.abs(875))])))
            d_813_v29_: _dafny.Array
            nw135_ = _dafny.Array(_dafny.Seq({}), 7)
            d_813_v29_ = nw135_
            d_814_v30_: _dafny.Map
            d_814_v30_ = _dafny.Map({p2: d_813_v29_})
            d_815_v31_: D10
            d_815_v31_ = D10_DC33(p0, (self).f36)
            d_814_v30_ = (d_814_v30_).set((d_815_v31_).cf55, d_813_v29_)
            (globalState).f1 = (0) - (default__.safeModuloInt(len((self.f30).set(default__.safeIndex((0) - ((self).f36), len(self.f30)), d_805_v24_)), default__.fm9(globalState)))
            r0 = self.f30
            d_816_v32_: D6
            d_816_v32_ = D6_DC22((self).f36, (self).f36, (self).f29)
            d_817_v33_: D9
            d_817_v33_ = D9_DC29((self).f36, (d_816_v32_).cf34, (self).f29)
            d_818_v34_: D9
            d_818_v34_ = D9_DC27(_dafny.SeqWithoutIsStrInference([(d_817_v33_).cf47, p2, (self).f36, (self).f36]))
            d_819_v35_: _dafny.Set
            d_819_v35_ = _dafny.Set({d_818_v34_, d_818_v34_, d_818_v34_})
            def iife73_():
                coll37_ = _dafny.Set()
                compr_37_: D9
                for compr_37_ in ((d_819_v35_) | (d_819_v35_)).Elements:
                    d_820_v36_: D9 = compr_37_
                    if (d_820_v36_) in ((d_819_v35_) | (d_819_v35_)):
                        coll37_ = coll37_.union(_dafny.Set([d_820_v36_]))
                return _dafny.Set(coll37_)
            (globalState).f14 = len(iife73_()
            )
        elif True:
            (globalState).f18 = True
            d_821_v37_: _dafny.MultiSet
            d_821_v37_ = _dafny.MultiSet([(self).f35])
            d_822_v38_: D0
            d_822_v38_ = D0_DC2((d_821_v37_).set((self).f35, default__.abs(p2)), self.f30, (self).f29)
            d_823_v39_: _dafny.Array
            nw136_ = _dafny.Array(None, 23)
            nw136_[int(0)] = p0
            nw136_[int(1)] = (self).f35
            nw136_[int(2)] = (self).f35
            nw136_[int(3)] = p0
            nw136_[int(4)] = (self).f35
            nw136_[int(5)] = (self).f35
            nw136_[int(6)] = p0
            nw136_[int(7)] = p0
            nw136_[int(8)] = (self).f29
            nw136_[int(9)] = (self).f29
            nw136_[int(10)] = (self).f29
            nw136_[int(11)] = (self).f35
            nw136_[int(12)] = (self).f29
            nw136_[int(13)] = (self).f35
            nw136_[int(14)] = (self).f29
            nw136_[int(15)] = p0
            nw136_[int(16)] = (self).f29
            nw136_[int(17)] = p0
            nw136_[int(18)] = p0
            nw136_[int(19)] = (self).f29
            nw136_[int(20)] = (self).f29
            nw136_[int(21)] = p0
            nw136_[int(22)] = (self).f35
            d_823_v39_ = nw136_
            d_824_v40_: _dafny.Seq
            d_824_v40_ = _dafny.SeqWithoutIsStrInference([d_823_v39_])
            d_825_v41_: _dafny.Map
            d_825_v41_ = _dafny.Map({(self).f29: (self).f35})
            d_826_v42_: _dafny.Array
            nw137_ = _dafny.Array(None, 15)
            nw137_[int(0)] = D0_DC2(d_821_v37_, self.f30, (self).f29)
            nw137_[int(1)] = (d_822_v38_ if (self).f29 else d_822_v38_)
            nw137_[int(2)] = d_822_v38_
            nw137_[int(3)] = d_822_v38_
            nw137_[int(4)] = default__.fm40(p0, globalState)
            nw137_[int(5)] = D0_DC2(_dafny.MultiSet(d_778_v0_), (D2_DC9(self.f30, len(d_824_v40_), (self).f29, p2)).cf14, (self).f35)
            nw137_[int(6)] = d_822_v38_
            nw137_[int(7)] = D0_DC2(default__.fm39(d_805_v24_, p2, globalState), self.f30, ((d_825_v41_)[(self).f29] if ((self).f29) in (d_825_v41_) else not(p0)))
            nw137_[int(8)] = d_822_v38_
            nw137_[int(9)] = D0_DC2(_dafny.MultiSet([(self).f35]), self.f30, p0)
            nw137_[int(10)] = D0_DC2(d_821_v37_, self.f30, (self).f29)
            nw137_[int(11)] = d_822_v38_
            nw137_[int(12)] = D0_DC2(d_821_v37_, self.f30, True)
            nw137_[int(13)] = d_822_v38_
            nw137_[int(14)] = d_822_v38_
            d_826_v42_ = nw137_
            index126_ = default__.safeIndex(517, (d_826_v42_).length(0))
            (d_826_v42_)[index126_] = D0_DC2(d_821_v37_, _dafny.SeqWithoutIsStrInference([d_805_v24_ for d_827_i3_ in range(default__.abs(926))]), p0)
            index127_ = default__.safeIndex(517, (d_826_v42_).length(0))
            (d_826_v42_)[index127_] = d_822_v38_
            index128_ = default__.safeIndex(641, (d_823_v39_).length(0))
            (d_823_v39_)[index128_] = (self).f29
            d_828_v43_: _dafny.MultiSet
            d_828_v43_ = _dafny.MultiSet([d_823_v39_])
            d_829_v44_: _dafny.Seq
            d_829_v44_ = _dafny.SeqWithoutIsStrInference([d_828_v43_, d_828_v43_, (d_828_v43_).set(d_823_v39_, default__.abs(default__.fm36(globalState)))])
            d_830_v45_: _dafny.Seq
            d_830_v45_ = _dafny.SeqWithoutIsStrInference([(self).f36, p2])
            index129_ = default__.safeIndex(641, (d_823_v39_).length(0))
            (d_823_v39_)[index129_] = ((d_829_v44_)[default__.safeIndex(len(d_830_v45_), len(d_829_v44_))]).ispropersubset(_dafny.MultiSet([d_823_v39_, d_823_v39_, d_823_v39_]))
            (globalState).f18 = not((d_823_v39_)[default__.safeIndex(641, (d_823_v39_).length(0))])
            (globalState).f21 = default__.fm9(globalState)
        d_831_v46_: _dafny.Set
        d_831_v46_ = _dafny.Set({d_805_v24_, d_805_v24_, d_805_v24_})
        d_831_v46_ = d_831_v46_
        d_832_v47_: _dafny.Array
        nw138_ = _dafny.Array(int(0), 6)
        d_832_v47_ = nw138_
        d_833_v48_: _dafny.Set
        d_833_v48_ = _dafny.Set({d_832_v47_})
        d_834_v49_: C2
        nw139_ = C2()
        nw139_.ctor__(len((self.f30) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivbrw")))), d_833_v48_, (self).f35, self.f30)
        d_834_v49_ = nw139_
        d_835_v50_: _dafny.Set
        d_835_v50_ = _dafny.Set({(self).f35})
        d_836_v51_: _dafny.Map
        d_836_v51_ = _dafny.Map({len(self.f30): (self).f29})
        hi2_ = len(d_836_v51_)
        for d_837_i4_ in range(len((d_835_v50_) - (d_835_v50_)), hi2_):
            (globalState).f13 = not (p0) or (p0)
            (globalState).f13 = default__.fm19(_dafny.CodePoint('h'), globalState)
            d_838_v52_: C0
            nw140_ = C0()
            nw140_.ctor__((self).f35)
            d_838_v52_ = nw140_
            d_839_v53_: _dafny.MultiSet
            d_839_v53_ = _dafny.MultiSet([162, (default__.fm39(d_805_v24_, 380, globalState)).cardinality])
            d_840_v55_: _dafny.Seq
            d_840_v55_ = _dafny.SeqWithoutIsStrInference([(d_834_v49_).f38])
            d_841_v56_: _dafny.Seq
            d_841_v56_ = _dafny.SeqWithoutIsStrInference([d_840_v55_, d_840_v55_])
            d_842_v57_: _dafny.Array
            nw141_ = _dafny.Array(None, 27)
            nw141_[int(0)] = (self).f29
            nw141_[int(1)] = p0
            nw141_[int(2)] = (self).f35
            def iife74_():
                coll38_ = _dafny.Map()
                compr_38_: int
                for compr_38_ in _dafny.IntegerRange(777, -289):
                    d_843_v54_: int = compr_38_
                    if ((777) <= (d_843_v54_)) and ((d_843_v54_) < (-289)):
                        coll38_[(d_843_v54_) * ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality)] = (d_834_v49_).f38
                return _dafny.Map(coll38_)
            def iife75_():
                coll39_ = _dafny.Map()
                compr_39_: int
                for compr_39_ in _dafny.IntegerRange(777, -289):
                    d_844_v54_: int = compr_39_
                    if ((777) <= (d_844_v54_)) and ((d_844_v54_) < (-289)):
                        coll39_[(d_844_v54_) * ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality)] = (d_834_v49_).f38
                return _dafny.Map(coll39_)
            nw141_[int(3)] = default__.fm33(True, (self).f36, ((d_839_v53_)[len((iife74_()
            ).set((self).f36, (self).f36))] if (len((iife75_()
            ).set((self).f36, (self).f36))) in (d_839_v53_) else (d_834_v49_).f38), globalState)
            nw141_[int(4)] = d_838_v52_.f37
            nw141_[int(5)] = (self).f35
            nw141_[int(6)] = d_838_v52_.f37
            nw141_[int(7)] = (self).f29
            nw141_[int(8)] = (d_838_v52_).fm11(True, globalState)
            nw141_[int(9)] = (len(default__.fm41(len(d_841_v56_), (self).f29, (self).f29, d_838_v52_.f37, globalState))) <= ((self).f36)
            nw141_[int(10)] = p0
            nw141_[int(11)] = True
            nw141_[int(12)] = (d_778_v0_) == (d_778_v0_)
            nw141_[int(13)] = (self).f29
            nw141_[int(14)] = d_838_v52_.f37
            nw141_[int(15)] = p0
            nw141_[int(16)] = (self).f35
            nw141_[int(17)] = (self).f29
            nw141_[int(18)] = (777) < (((d_804_v23_)[793] if (793) in (d_804_v23_) else (self).f36))
            nw141_[int(19)] = ((d_834_v49_).f38) <= (len(_dafny.Map({d_838_v52_: (self).f29})))
            nw141_[int(20)] = p0
            nw141_[int(21)] = (self).f35
            nw141_[int(22)] = (d_778_v0_)[default__.safeIndex((d_834_v49_).f38, len(d_778_v0_))]
            nw141_[int(23)] = (self).f29
            nw141_[int(24)] = (self).f29
            nw141_[int(25)] = not((self).f35)
            nw141_[int(26)] = (_dafny.SeqWithoutIsStrInference([d_805_v24_, default__.fm42(d_838_v52_.f37, d_838_v52_.f37, globalState)])) < (_dafny.SeqWithoutIsStrInference([d_805_v24_ for d_845_i5_ in range(default__.abs(688))]))
            d_842_v57_ = nw141_
            index130_ = default__.safeIndex(613, (d_842_v57_).length(0))
            (d_842_v57_)[index130_] = (self).f35
            d_846_v58_: _dafny.Array
            nw142_ = _dafny.Array(D1.default()(), 29)
            d_846_v58_ = nw142_
            d_847_v59_: D1
            d_847_v59_ = D1_DC3(d_778_v0_)
            d_848_v60_: D1
            d_848_v60_ = D1_DC3((d_847_v59_).cf7)
            d_849_v61_: D1
            d_849_v61_ = D1_DC5(d_848_v60_)
            index131_ = default__.safeIndex(34, (d_846_v58_).length(0))
            (d_846_v58_)[index131_] = d_849_v61_
            d_850_v62_: D10
            d_850_v62_ = D10_DC33(False, p2)
            index132_ = default__.safeIndex(613, (d_842_v57_).length(0))
            index133_ = default__.safeIndex(34, (d_846_v58_).length(0))
            rhs138_ = (d_838_v52_).fm11((d_850_v62_).cf54, globalState)
            rhs139_ = d_849_v61_
            lhs127_ = d_842_v57_
            lhs128_ = default__.safeIndex(613, (d_842_v57_).length(0))
            lhs129_ = d_846_v58_
            lhs130_ = default__.safeIndex(34, (d_846_v58_).length(0))
            lhs127_[lhs128_] = rhs138_
            lhs129_[lhs130_] = rhs139_
        d_851_v63_: _dafny.Seq
        d_851_v63_ = _dafny.SeqWithoutIsStrInference([self.f30])
        r0 = (d_851_v63_)[default__.safeIndex(len(self.f30), len(d_851_v63_))]
        d_852_v64_: _dafny.MultiSet
        d_852_v64_ = _dafny.MultiSet([(0) - (p2), (self).f36])
        r1 = (default__.safeDivisionInt((d_834_v49_).f38, (d_852_v64_).cardinality)) > (p2)
        return r0, r1

    @property
    def f35(self):
        return self._f35
    @property
    def f36(self):
        return self._f36

class C4(T0):
    def  __init__(self):
        self._f33: bool = False
        self._f34: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self, f33, f34):
        (self)._f33 = f33
        (self)._f34 = f34

    def fm5(self, p0, globalState):
        return default__.safeDivisionInt((self).f34, (0) - ((self).f34))

    def m2(self, p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_853_i0_: int
        d_853_i0_ = 0
        with _dafny.label("3"):
            while default__.fm6(((self).f34) * ((self).f34), globalState):
                with _dafny.c_label("3"):
                    if (d_853_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_853_i0_ = (d_853_i0_) + (1)
                    d_854_v0_: _dafny.Map
                    d_854_v0_ = _dafny.Map({(self).f34: (self).f33})
                    d_854_v0_ = (d_854_v0_).set((self).f34, (self).f33)
                    if (self).f33:
                        (globalState).f28 = p0
                        d_855_v1_: _dafny.Set
                        d_855_v1_ = _dafny.Set({p0})
                        (globalState).f2 = (self).fm5((d_855_v1_) | (_dafny.Set({default__.fm6((self).f34, globalState)})), globalState)
                        (globalState).f18 = not((self).f33)
                        d_856_v2_: _dafny.Array
                        nw143_ = _dafny.Array(int(0), 22)
                        d_856_v2_ = nw143_
                        (globalState).f17 = d_856_v2_
                        (globalState).f21 = (((self).f34) + ((self).f34)) * ((self).f34)
                    elif True:
                        d_857_v3_: _dafny.Seq
                        d_857_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yk"))
                        (globalState).f7 = _dafny.SeqWithoutIsStrInference([((self).f34) * (len(d_857_v3_))])
                        (globalState).f28 = not (p0) or (p0)
                        d_858_v4_: _dafny.Array
                        nw144_ = _dafny.Array(_dafny.Array(None, 0), 7)
                        d_858_v4_ = nw144_
                        d_859_v5_: _dafny.Array
                        def lambda36_(d_860_p0_):
                            def lambda37_(d_861_i1_):
                                return _dafny.SeqWithoutIsStrInference([(self).f33, d_860_p0_])

                            return lambda37_

                        init20_ = lambda36_(p0)
                        nw145_ = _dafny.Array(None, 12)
                        for i0_20_ in range(nw145_.length(0)):
                            nw145_[i0_20_] = init20_(i0_20_)
                        d_859_v5_ = nw145_
                        d_862_v6_: _dafny.Seq
                        d_862_v6_ = _dafny.SeqWithoutIsStrInference([d_859_v5_])
                        index134_ = default__.safeIndex(475, (d_858_v4_).length(0))
                        (d_858_v4_)[index134_] = (d_862_v6_)[default__.safeIndex((self).f34, len(d_862_v6_))]
                        index135_ = default__.safeIndex(475, (d_858_v4_).length(0))
                        (d_858_v4_)[index135_] = d_859_v5_
                        d_863_v7_: _dafny.Array
                        nw146_ = _dafny.Array(False, 11)
                        d_863_v7_ = nw146_
                        d_863_v7_ = d_863_v7_
                        (globalState).f15 = p0
                    d_864_v8_: _dafny.Seq
                    d_864_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ue"))
                    d_865_v9_: _dafny.Seq
                    d_865_v9_ = _dafny.SeqWithoutIsStrInference([d_864_v8_])
                    (globalState).f2 = len(d_865_v9_)
                    d_866_v10_: _dafny.Seq
                    d_866_v10_ = _dafny.SeqWithoutIsStrInference([(self).f33, False, (self).f33, False, p0])
                    (globalState).f15 = (((self).f33) not in (d_866_v10_) if False else (self).f33)
                    pass
            pass
        hi3_ = (self).f34
        for d_867_i2_ in range(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwevwbc"))), hi3_):
            (globalState).f15 = not(default__.fm6(d_867_i2_, globalState))
            d_868_v11_: _dafny.Map
            d_868_v11_ = _dafny.Map({748: (self).f34})
            d_868_v11_ = (d_868_v11_).set(-437, (self).f34)
            d_869_v12_: _dafny.Array
            def lambda38_(d_870_i3_):
                return (352) != ((0) - ((self).f34))

            init21_ = lambda38_
            nw147_ = _dafny.Array(None, 28)
            for i0_21_ in range(nw147_.length(0)):
                nw147_[i0_21_] = init21_(i0_21_)
            d_869_v12_ = nw147_
            d_869_v12_ = d_869_v12_
            d_871_v13_: _dafny.Seq
            d_871_v13_ = _dafny.SeqWithoutIsStrInference([d_869_v12_, d_869_v12_])
            d_872_v14_: _dafny.Array
            nw148_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_872_v14_ = nw148_
            d_873_v15_: _dafny.Array
            def lambda39_(d_874_i4_):
                return (d_874_i4_) - (537)

            init22_ = lambda39_
            nw149_ = _dafny.Array(None, 13)
            for i0_22_ in range(nw149_.length(0)):
                nw149_[i0_22_] = init22_(i0_22_)
            d_873_v15_ = nw149_
            index136_ = default__.safeIndex(20, (d_872_v14_).length(0))
            (d_872_v14_)[index136_] = d_873_v15_
            d_875_v16_: _dafny.Seq
            d_875_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvp"))
            d_876_v17_: _dafny.MultiSet
            d_876_v17_ = _dafny.MultiSet([p0])
            d_877_v18_: _dafny.Seq
            d_877_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkowf")), default__.fm7((self).f34, globalState)])
            d_878_v19_: str
            d_878_v19_ = _dafny.CodePoint('j')
            d_879_v20_: _dafny.Map
            d_879_v20_ = _dafny.Map({d_867_i2_: d_878_v19_})
            pat_let_tv20_ = globalState
            d_880_v21_: D2
            def iife76_(_pat_let18_0):
                def iife77_(d_881_dt__update__tmp_h0_):
                    def iife78_(_pat_let19_0):
                        def iife79_(d_882_dt__update_hcf5_h0_):
                            def iife80_(_pat_let20_0):
                                def iife81_(d_883_dt__update_hcf6_h0_):
                                    return D0_DC2((d_881_dt__update__tmp_h0_).cf4, d_882_dt__update_hcf5_h0_, d_883_dt__update_hcf6_h0_)
                                return iife81_(_pat_let20_0)
                            return iife80_(True)
                        return iife79_(_pat_let19_0)
                    return iife78_(default__.fm7(d_867_i2_, pat_let_tv20_))
                return iife77_(_pat_let18_0)
            d_880_v21_ = D2_DC9(d_875_v16_, len((d_868_v11_).set(len(d_875_v16_), d_867_i2_)), (iife76_(D0_DC2(d_876_v17_, (d_877_v18_)[default__.safeIndex(len((d_879_v20_).set(d_867_i2_, d_878_v19_)), len(d_877_v18_))], p0))).cf6, (self).f34)
            index137_ = default__.safeIndex(20, (d_872_v14_).length(0))
            rhs140_ = _dafny.SeqWithoutIsStrInference([(d_869_v12_ if p0 else d_869_v12_), d_869_v12_, d_869_v12_, d_869_v12_])
            rhs141_ = d_873_v15_
            rhs142_ = (d_880_v21_).cf17
            lhs131_ = d_872_v14_
            lhs132_ = default__.safeIndex(20, (d_872_v14_).length(0))
            lhs133_ = globalState
            d_871_v13_ = rhs140_
            lhs131_[lhs132_] = rhs141_
            lhs133_.f2 = rhs142_
        d_884_v22_: _dafny.Array
        def lambda40_(d_885_i5_):
            return default__.safeDivisionInt(d_885_i5_, (self).f34)

        init23_ = lambda40_
        nw150_ = _dafny.Array(None, 9)
        for i0_23_ in range(nw150_.length(0)):
            nw150_[i0_23_] = init23_(i0_23_)
        d_884_v22_ = nw150_
        index138_ = default__.safeIndex(88, (d_884_v22_).length(0))
        (d_884_v22_)[index138_] = (self).f34
        d_886_v23_: _dafny.Seq
        d_886_v23_ = _dafny.SeqWithoutIsStrInference([True])
        index139_ = default__.safeIndex(88, (d_884_v22_).length(0))
        rhs143_ = (self).f34
        rhs144_ = ((_dafny.SeqWithoutIsStrInference([(self).f33]) if p0 else _dafny.SeqWithoutIsStrInference([(self).f33]))) + (d_886_v23_)
        lhs134_ = d_884_v22_
        lhs135_ = default__.safeIndex(88, (d_884_v22_).length(0))
        lhs134_[lhs135_] = rhs143_
        d_886_v23_ = rhs144_
        d_887_v24_: _dafny.MultiSet
        d_887_v24_ = _dafny.MultiSet([(d_884_v22_)[default__.safeIndex(88, (d_884_v22_).length(0))], (d_884_v22_)[default__.safeIndex(88, (d_884_v22_).length(0))], (self).f34])
        d_888_v25_: D3
        d_888_v25_ = D3_DC10(d_887_v24_)
        (globalState).f28 = (d_887_v24_) == ((d_888_v25_).cf18)
        index140_ = default__.safeIndex(88, (d_884_v22_).length(0))
        (d_884_v22_)[index140_] = (self).f34
        def lambda41_(d_889_i6_):
            return default__.safeDivisionInt(d_889_i6_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmap"))))

        init24_ = lambda41_
        nw151_ = _dafny.Array(None, 2)
        for i0_24_ in range(nw151_.length(0)):
            nw151_[i0_24_] = init24_(i0_24_)
        d_884_v22_ = nw151_
        d_890_v26_: _dafny.Seq
        d_890_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hoalen"))
        d_891_v27_: _dafny.Seq
        d_891_v27_ = _dafny.SeqWithoutIsStrInference([d_886_v23_, (_dafny.SeqWithoutIsStrInference([False, (self).f33, p0, p0, default__.fm6(len(d_890_v26_), globalState)])) + (d_886_v23_)])
        r0 = d_891_v27_
        r1 = not((default__.safeDivisionInt((d_884_v22_)[default__.safeIndex(88, (d_884_v22_).length(0))], (d_884_v22_)[default__.safeIndex(88, (d_884_v22_).length(0))])) <= ((d_884_v22_)[default__.safeIndex(88, (d_884_v22_).length(0))]))
        return r0, r1

    def m3(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        d_892_v0_: _dafny.Array
        nw152_ = _dafny.Array(False, 11)
        d_892_v0_ = nw152_
        index141_ = default__.safeIndex(618, (d_892_v0_).length(0))
        (d_892_v0_)[index141_] = (self).f33
        index142_ = default__.safeIndex(618, (d_892_v0_).length(0))
        (d_892_v0_)[index142_] = not ((self).f33) or ((self).f33)
        (globalState).f28 = (self).f33
        d_893_v1_: _dafny.Seq
        d_893_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fckyvw"))
        d_894_v2_: D2
        d_894_v2_ = D2_DC9(d_893_v1_, p2, p0, (self).f34)
        r0 = (d_893_v1_).set(default__.safeIndex((d_894_v2_).cf15, len(d_893_v1_)), (d_893_v1_)[default__.safeIndex(p2, len(d_893_v1_))])
        d_895_v3_: _dafny.Set
        d_895_v3_ = _dafny.Set({False, (d_892_v0_)[default__.safeIndex(618, (d_892_v0_).length(0))]})
        d_894_v2_ = default__.fm8(not(((self).f34) > ((self).fm5(d_895_v3_, globalState))), (d_892_v0_)[default__.safeIndex(618, (d_892_v0_).length(0))], globalState)
        d_896_v4_: str
        d_896_v4_ = _dafny.CodePoint('u')
        d_897_v5_: _dafny.Array
        nw153_ = _dafny.Array(int(0), 26)
        d_897_v5_ = nw153_
        index143_ = default__.safeIndex(73, (d_897_v5_).length(0))
        (d_897_v5_)[index143_] = p2
        d_898_v6_: _dafny.Set
        d_898_v6_ = _dafny.Set({p2})
        d_899_v7_: _dafny.MultiSet
        d_899_v7_ = _dafny.MultiSet([p2, (self).f34, ((self).f34 if p0 else (self).f34), len(d_898_v6_), 223])
        d_900_v8_: _dafny.Seq
        d_900_v8_ = _dafny.SeqWithoutIsStrInference([(self).f33, p0, (d_892_v0_)[default__.safeIndex(618, (d_892_v0_).length(0))]])
        d_901_v9_: _dafny.Map
        d_901_v9_ = _dafny.Map({p2: p2})
        d_902_v10_: _dafny.MultiSet
        d_902_v10_ = _dafny.MultiSet([d_900_v8_])
        d_903_v11_: _dafny.Seq
        d_903_v11_ = _dafny.SeqWithoutIsStrInference([(self).f34, (d_902_v10_).cardinality, p2])
        d_904_v12_: _dafny.Map
        d_904_v12_ = _dafny.Map({(d_892_v0_)[default__.safeIndex(618, (d_892_v0_).length(0))]: _dafny.SeqWithoutIsStrInference([(0) - (p2)])})
        index144_ = default__.safeIndex(73, (d_897_v5_).length(0))
        rhs145_ = ((d_899_v7_)[(self).fm5(_dafny.Set({(d_900_v8_)[default__.safeIndex(p2, len(d_900_v8_))]}), globalState)] if ((self).fm5(_dafny.Set({(d_900_v8_)[default__.safeIndex(p2, len(d_900_v8_))]}), globalState)) in (d_899_v7_) else len(d_901_v9_))
        rhs146_ = not((d_892_v0_)[default__.safeIndex(618, (d_892_v0_).length(0))])
        rhs147_ = d_896_v4_
        rhs148_ = ((_dafny.SeqWithoutIsStrInference([(self).f34, (self).f34, p2, len(d_900_v8_), (self).f34])) + (d_903_v11_)) + (((d_904_v12_)[(d_892_v0_)[default__.safeIndex(618, (d_892_v0_).length(0))]] if ((d_892_v0_)[default__.safeIndex(618, (d_892_v0_).length(0))]) in (d_904_v12_) else d_903_v11_))
        rhs149_ = p2
        lhs136_ = globalState
        lhs137_ = globalState
        lhs138_ = globalState
        lhs139_ = d_897_v5_
        lhs140_ = default__.safeIndex(73, (d_897_v5_).length(0))
        lhs136_.f21 = rhs145_
        lhs137_.f28 = rhs146_
        d_896_v4_ = rhs147_
        lhs138_.f7 = rhs148_
        lhs139_[lhs140_] = rhs149_
        index145_ = default__.safeIndex(801, (p1).length(0))
        (p1)[index145_] = _dafny.SeqWithoutIsStrInference([(D2_DC8(d_896_v4_, (self).f33)).cf12 for d_905_i0_ in range(default__.abs(-838))])
        index146_ = default__.safeIndex(801, (p1).length(0))
        (p1)[index146_] = d_893_v1_
        r0 = (d_893_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "keatndrhj")))
        r1 = default__.fm6(82, globalState)
        return r0, r1

    def m6(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        d_906_v0_: _dafny.Array
        nw154_ = _dafny.Array(None, 29)
        nw154_[int(0)] = p0
        nw154_[int(1)] = p0
        nw154_[int(2)] = (self).f33
        nw154_[int(3)] = False
        nw154_[int(4)] = p0
        nw154_[int(5)] = p0
        nw154_[int(6)] = not(p0)
        nw154_[int(7)] = (self).f33
        nw154_[int(8)] = not(True)
        nw154_[int(9)] = (self).f33
        nw154_[int(10)] = (self).f33
        nw154_[int(11)] = p0
        nw154_[int(12)] = (self).f33
        nw154_[int(13)] = (self).f33
        nw154_[int(14)] = (self).f33
        nw154_[int(15)] = True
        nw154_[int(16)] = (self).f33
        nw154_[int(17)] = p0
        nw154_[int(18)] = (self).f33
        nw154_[int(19)] = p0
        nw154_[int(20)] = (self).f33
        nw154_[int(21)] = p0
        nw154_[int(22)] = (self).f33
        nw154_[int(23)] = p0
        nw154_[int(24)] = True
        nw154_[int(25)] = True
        nw154_[int(26)] = p0
        nw154_[int(27)] = (self).f33
        nw154_[int(28)] = p0
        d_906_v0_ = nw154_
        d_907_v1_: _dafny.Seq
        d_907_v1_ = _dafny.SeqWithoutIsStrInference([d_906_v0_])
        d_908_v2_: _dafny.MultiSet
        d_908_v2_ = _dafny.MultiSet([default__.safeModuloInt(len(d_907_v1_), (self).f34)])
        d_909_v3_: _dafny.Set
        d_909_v3_ = _dafny.Set({(self).f33})
        d_910_v4_: _dafny.Seq
        d_910_v4_ = _dafny.SeqWithoutIsStrInference([p0])
        rhs150_ = (self).fm5(d_909_v3_, globalState)
        rhs151_ = (d_910_v4_)[default__.safeIndex(978, len(d_910_v4_))]
        rhs152_ = len(p2)
        rhs153_ = not((self).f33)
        rhs154_ = d_908_v2_
        lhs141_ = globalState
        lhs142_ = globalState
        lhs143_ = globalState
        lhs141_.f2 = rhs150_
        r0 = rhs151_
        lhs142_.f21 = rhs152_
        lhs143_.f13 = rhs153_
        d_908_v2_ = rhs154_
        d_911_v5_: _dafny.Seq
        d_911_v5_ = _dafny.SeqWithoutIsStrInference([(d_910_v4_).set(default__.safeIndex((self).f34, len(d_910_v4_)), (self).f33), d_910_v4_])
        d_912_v6_: _dafny.Seq
        d_912_v6_ = _dafny.SeqWithoutIsStrInference([(d_911_v5_)[default__.safeIndex((self).f34, len(d_911_v5_))]])
        (globalState).f2 = ((len((d_911_v5_)[default__.safeIndex((self).f34, len(d_911_v5_))])) - ((self).f34) if default__.fm6((self).f34, globalState) else (self).fm5(d_909_v3_, globalState))
        d_913_i0_: int
        d_913_i0_ = 0
        with _dafny.label("4"):
            while not(p0):
                with _dafny.c_label("4"):
                    if (d_913_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_913_i0_ = (d_913_i0_) + (1)
                    (globalState).f28 = ((self).f34) <= ((self).f34)
                    d_914_v7_: _dafny.Array
                    nw155_ = _dafny.Array(_dafny.Map({}), 4)
                    d_914_v7_ = nw155_
                    d_915_v8_: str
                    d_915_v8_ = _dafny.CodePoint('k')
                    d_916_v9_: _dafny.Map
                    d_916_v9_ = _dafny.Map({d_915_v8_: (d_910_v4_)[default__.safeIndex((self).f34, len(d_910_v4_))]})
                    index147_ = default__.safeIndex(78, (d_914_v7_).length(0))
                    (d_914_v7_)[index147_] = d_916_v9_
                    d_917_v11_: _dafny.Map
                    d_917_v11_ = _dafny.Map({_dafny.CodePoint('j'): (0) - ((self).f34)})
                    index148_ = default__.safeIndex(78, (d_914_v7_).length(0))
                    def iife82_():
                        coll40_ = _dafny.Map()
                        compr_40_: str
                        for compr_40_ in (d_917_v11_).keys.Elements:
                            d_918_v10_: str = compr_40_
                            if (d_918_v10_) in (d_917_v11_):
                                coll40_[d_918_v10_] = False
                        return _dafny.Map(coll40_)
                    rhs155_ = iife82_()

                    rhs156_ = (0) - (default__.safeDivisionInt((self).f34, (self).f34))
                    rhs157_ = (self).f34
                    lhs144_ = d_914_v7_
                    lhs145_ = default__.safeIndex(78, (d_914_v7_).length(0))
                    lhs146_ = globalState
                    lhs147_ = globalState
                    lhs144_[lhs145_] = rhs155_
                    lhs146_.f2 = rhs156_
                    lhs147_.f14 = rhs157_
                    (globalState).f2 = ((self).f34) + ((self).f34)
                    d_919_v12_: _dafny.Map
                    d_919_v12_ = _dafny.Map({default__.safeModuloInt((self).f34, (self).f34): p0})
                    d_919_v12_ = (d_919_v12_).set((self).f34, p0)
                    pass
            pass
        r0 = (self).f33
        d_920_v13_: _dafny.Set
        d_920_v13_ = _dafny.Set({(self).f34, (self).f34})
        (globalState).f14 = default__.safeDivisionInt(default__.safeModuloInt((self).f34, len(d_920_v13_)), (self).f34)
        hi4_ = (self).f34
        for d_921_i1_ in range((self).f34, hi4_):
            d_922_v14_: _dafny.MultiSet
            d_922_v14_ = _dafny.MultiSet([default__.fm6((self).f34, globalState)])
            d_923_v15_: _dafny.Seq
            d_923_v15_ = _dafny.SeqWithoutIsStrInference([d_921_i1_])
            d_924_v16_: str
            d_924_v16_ = _dafny.CodePoint('p')
            d_925_v17_: _dafny.Map
            d_925_v17_ = _dafny.Map({d_921_i1_: d_924_v16_})
            d_926_v18_: _dafny.Array
            nw156_ = _dafny.Array(None, 11)
            nw156_[int(0)] = d_921_i1_
            nw156_[int(1)] = len(((_dafny.SeqWithoutIsStrInference([(self).f34, (self).f34, d_921_i1_, d_921_i1_, (0) - ((0) - (d_921_i1_))])).set(default__.safeIndex(((d_922_v14_)[True] if (True) in (d_922_v14_) else d_921_i1_), len(_dafny.SeqWithoutIsStrInference([(self).f34, (self).f34, d_921_i1_, d_921_i1_, (0) - ((0) - (d_921_i1_))]))), d_921_i1_)) + (d_923_v15_))
            nw156_[int(2)] = default__.safeDivisionInt(d_921_i1_, d_921_i1_)
            nw156_[int(3)] = d_921_i1_
            nw156_[int(4)] = (self).f34
            nw156_[int(5)] = d_921_i1_
            nw156_[int(6)] = 305
            nw156_[int(7)] = d_921_i1_
            nw156_[int(8)] = d_921_i1_
            nw156_[int(9)] = (0) - (len(d_925_v17_))
            nw156_[int(10)] = d_921_i1_
            d_926_v18_ = nw156_
            (globalState).f17 = d_926_v18_
            (globalState).f15 = not((self).f33)
            d_927_v19_: C1
            nw157_ = C1()
            nw157_.ctor__(p0)
            d_927_v19_ = nw157_
            if (d_927_v19_).f40:
                (globalState).f17 = d_926_v18_
                index149_ = default__.safeIndex(887, (d_906_v0_).length(0))
                (d_906_v0_)[index149_] = not((self).f33)
                d_928_v20_: _dafny.Map
                d_928_v20_ = _dafny.Map({d_921_i1_: (self).f34})
                index150_ = default__.safeIndex(887, (d_906_v0_).length(0))
                rhs158_ = (default__.fm6(len(d_928_v20_), globalState)) == (True)
                rhs159_ = (d_909_v3_).isdisjoint((d_909_v3_ if not((self).f33) else d_909_v3_))
                rhs160_ = ((self).f34) != ((self).f34)
                lhs148_ = globalState
                lhs149_ = d_906_v0_
                lhs150_ = default__.safeIndex(887, (d_906_v0_).length(0))
                lhs151_ = globalState
                lhs148_.f13 = rhs158_
                lhs149_[lhs150_] = rhs159_
                lhs151_.f13 = rhs160_
                d_929_v21_: _dafny.Array
                nw158_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_929_v21_ = nw158_
                d_930_v22_: _dafny.Map
                d_930_v22_ = _dafny.Map({21: default__.fm19(d_924_v16_, globalState)})
                d_931_v23_: _dafny.Map
                d_931_v23_ = _dafny.Map({not((d_927_v19_).f40): p0})
                d_932_v24_: _dafny.Array
                nw159_ = _dafny.Array(None, 28)
                nw159_[int(0)] = (d_906_v0_)[default__.safeIndex(887, (d_906_v0_).length(0))]
                nw159_[int(1)] = True
                nw159_[int(2)] = (d_910_v4_)[default__.safeIndex((self).f34, len(d_910_v4_))]
                nw159_[int(3)] = (p3).cf0
                nw159_[int(4)] = True
                nw159_[int(5)] = (d_927_v19_).f40
                nw159_[int(6)] = (d_927_v19_).f40
                nw159_[int(7)] = p0
                nw159_[int(8)] = (d_927_v19_).f40
                nw159_[int(9)] = (self).f33
                nw159_[int(10)] = ((d_930_v22_)[(self).f34] if ((self).f34) in (d_930_v22_) else (self).f33)
                nw159_[int(11)] = (self).f33
                nw159_[int(12)] = (self).f33
                nw159_[int(13)] = (d_927_v19_).f40
                nw159_[int(14)] = ((d_931_v23_)[(self).f33] if ((self).f33) in (d_931_v23_) else (d_927_v19_).f40)
                nw159_[int(15)] = (d_927_v19_).f40
                nw159_[int(16)] = (d_906_v0_)[default__.safeIndex(887, (d_906_v0_).length(0))]
                nw159_[int(17)] = (d_906_v0_)[default__.safeIndex(887, (d_906_v0_).length(0))]
                nw159_[int(18)] = p0
                nw159_[int(19)] = (d_927_v19_).f40
                nw159_[int(20)] = p0
                nw159_[int(21)] = (d_906_v0_)[default__.safeIndex(887, (d_906_v0_).length(0))]
                nw159_[int(22)] = (self).f33
                nw159_[int(23)] = (d_927_v19_).f40
                nw159_[int(24)] = (d_927_v19_).f40
                nw159_[int(25)] = p0
                nw159_[int(26)] = (d_906_v0_)[default__.safeIndex(887, (d_906_v0_).length(0))]
                nw159_[int(27)] = (d_906_v0_)[default__.safeIndex(887, (d_906_v0_).length(0))]
                d_932_v24_ = nw159_
                index151_ = default__.safeIndex(457, (d_929_v21_).length(0))
                (d_929_v21_)[index151_] = d_932_v24_
                index152_ = default__.safeIndex(457, (d_929_v21_).length(0))
                (d_929_v21_)[index152_] = d_932_v24_
                d_931_v23_ = (d_931_v23_).set((_dafny.Set({d_921_i1_})) == (d_920_v13_), True)
                d_933_v25_: _dafny.Map
                d_933_v25_ = _dafny.Map({d_926_v18_: d_921_i1_})
                d_933_v25_ = (d_933_v25_).set(d_926_v18_, (self).f34)
            elif True:
                d_925_v17_ = (d_925_v17_).set(-901, _dafny.CodePoint('y'))
                (globalState).f14 = 566
                (globalState).f28 = (d_927_v19_).f40
                (globalState).f2 = len((d_910_v4_) + (d_910_v4_))
                d_934_v26_: _dafny.Array
                def lambda42_(d_935_i2_):
                    return D8_DC26()

                init25_ = lambda42_
                nw160_ = _dafny.Array(None, 5)
                for i0_25_ in range(nw160_.length(0)):
                    nw160_[i0_25_] = init25_(i0_25_)
                d_934_v26_ = nw160_
                d_936_v27_: D8
                d_936_v27_ = D8_DC26()
                index153_ = default__.safeIndex(556, (d_934_v26_).length(0))
                (d_934_v26_)[index153_] = (d_936_v27_ if (d_927_v19_).f40 else D8_DC26())
                d_937_v28_: D2
                d_937_v28_ = D2_DC8(d_924_v16_, (self).f33)
                d_938_v29_: _dafny.Seq
                d_938_v29_ = _dafny.SeqWithoutIsStrInference([d_937_v28_])
                index154_ = default__.safeIndex(556, (d_934_v26_).length(0))
                rhs161_ = (((0) - ((d_908_v2_).cardinality)) - (len(default__.fm35(d_938_v29_, d_921_i1_, globalState)))) <= (d_921_i1_)
                rhs162_ = d_936_v27_
                rhs163_ = p0
                lhs152_ = globalState
                lhs153_ = d_934_v26_
                lhs154_ = default__.safeIndex(556, (d_934_v26_).length(0))
                lhs155_ = globalState
                lhs152_.f15 = rhs161_
                lhs153_[lhs154_] = rhs162_
                lhs155_.f13 = rhs163_
        r0 = (p0 if p0 else (self).f33)
        r1 = (self).f33
        return r0, r1

    def m7(self, p0, globalState):
        r0: bool = False
        d_939_v0_: _dafny.Array
        def lambda43_(d_940_i1_):
            return (d_940_i1_) * (len(_dafny.Map({(self).f33: (self).f33})))

        init26_ = lambda43_
        nw161_ = _dafny.Array(None, 10)
        for i0_26_ in range(nw161_.length(0)):
            nw161_[i0_26_] = init26_(i0_26_)
        d_939_v0_ = nw161_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_939_v0_).length(0)):
            d_941_i0_: int = guard_loop_7_
            if (True) and (((0) <= (d_941_i0_)) and ((d_941_i0_) < ((d_939_v0_).length(0)))):
                (d_939_v0_)[(d_941_i0_)] = default__.safeModuloInt(d_941_i0_, p0)
        d_942_v1_: _dafny.Seq
        d_942_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqds"))
        d_943_v2_: C3
        nw162_ = C3()
        nw162_.ctor__((self).f33, len(_dafny.Map({(self).f33: p0})), ((self).f34) != ((0) - ((self).f34)), d_942_v1_)
        d_943_v2_ = nw162_
        d_944_v3_: _dafny.Seq
        d_944_v3_ = _dafny.SeqWithoutIsStrInference([(self).f33])
        d_945_v4_: D1
        d_945_v4_ = D1_DC3(d_944_v3_)
        source15_ = (D1_DC3(_dafny.SeqWithoutIsStrInference([(self).f33])) if (d_943_v2_).f35 else d_945_v4_)
        if source15_.is_DC4:
            d_946___mcc_h0_ = source15_.cf8
            d_947_cf8_ = d_946___mcc_h0_
            d_948_v5_: _dafny.Array
            nw163_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_948_v5_ = nw163_
            d_949_v6_: str
            d_949_v6_ = _dafny.CodePoint('a')
            d_950_v7_: _dafny.Array
            nw164_ = _dafny.Array(None, 9)
            nw164_[int(0)] = not(False)
            nw164_[int(1)] = (d_943_v2_).f35
            nw164_[int(2)] = (d_943_v2_).f35
            nw164_[int(3)] = d_947_cf8_
            nw164_[int(4)] = False
            nw164_[int(5)] = (d_943_v2_).f35
            nw164_[int(6)] = default__.fm19(d_949_v6_, globalState)
            nw164_[int(7)] = (d_943_v2_).f35
            nw164_[int(8)] = d_947_cf8_
            d_950_v7_ = nw164_
            d_951_v8_: _dafny.Seq
            d_951_v8_ = _dafny.SeqWithoutIsStrInference([d_950_v7_, d_950_v7_])
            index155_ = default__.safeIndex(280, (d_948_v5_).length(0))
            (d_948_v5_)[index155_] = (d_951_v8_)[default__.safeIndex(-173, len(d_951_v8_))]
            d_952_v9_: _dafny.Map
            d_952_v9_ = _dafny.Map({-577: False})
            d_953_v10_: D4
            d_953_v10_ = D4_DC14(len(d_952_v9_))
            d_954_v11_: _dafny.Array
            nw165_ = _dafny.Array(None, 8)
            def iife83_(_pat_let21_0):
                def iife84_(d_955_dt__update__tmp_h0_):
                    def iife85_(_pat_let22_0):
                        def iife86_(d_956_dt__update_hcf23_h0_):
                            return D4_DC14(d_956_dt__update_hcf23_h0_)
                        return iife86_(_pat_let22_0)
                    return iife85_((0) - ((self).f34))
                return iife84_(_pat_let21_0)
            nw165_[int(0)] = iife83_(d_953_v10_)
            nw165_[int(1)] = D4_DC14((d_943_v2_).f36)
            nw165_[int(2)] = d_953_v10_
            nw165_[int(3)] = D4_DC14(len(default__.fm20((self).f33, (d_943_v2_).f35, d_944_v3_, len(_dafny.SeqWithoutIsStrInference([(d_943_v2_).f36 for d_957_i2_ in range(default__.abs(565))])), globalState)))
            nw165_[int(4)] = d_953_v10_
            nw165_[int(5)] = d_953_v10_
            nw165_[int(6)] = d_953_v10_
            nw165_[int(7)] = d_953_v10_
            d_954_v11_ = nw165_
            index156_ = default__.safeIndex(231, (d_954_v11_).length(0))
            (d_954_v11_)[index156_] = d_953_v10_
            pat_let_tv21_ = d_943_v2_
            index157_ = default__.safeIndex(280, (d_948_v5_).length(0))
            index158_ = default__.safeIndex(231, (d_954_v11_).length(0))
            def iife87_(_pat_let23_0):
                def iife88_(d_958_dt__update__tmp_h1_):
                    def iife89_(_pat_let24_0):
                        def iife90_(d_959_dt__update_hcf23_h1_):
                            return D4_DC14(d_959_dt__update_hcf23_h1_)
                        return iife90_(_pat_let24_0)
                    return iife89_((pat_let_tv21_).f36)
                return iife88_(_pat_let23_0)
            rhs164_ = d_950_v7_
            rhs165_ = iife87_(d_953_v10_)
            rhs166_ = d_947_cf8_
            lhs156_ = d_948_v5_
            lhs157_ = default__.safeIndex(280, (d_948_v5_).length(0))
            lhs158_ = d_954_v11_
            lhs159_ = default__.safeIndex(231, (d_954_v11_).length(0))
            lhs156_[lhs157_] = rhs164_
            lhs158_[lhs159_] = rhs165_
            r0 = rhs166_
            d_960_v12_: _dafny.Map
            d_960_v12_ = _dafny.Map({(self).f33: (self).f33})
            d_961_v13_: _dafny.Set
            d_961_v13_ = _dafny.Set({len(d_960_v12_), (d_943_v2_).f36})
            d_962_v14_: D9
            d_962_v14_ = D9_DC28(d_944_v3_, (d_943_v2_).f35, (d_943_v2_).f36, (d_943_v2_).f35, (d_943_v2_).f35)
            source16_ = default__.fm43(_dafny.MultiSet([len(d_961_v13_)]), (d_962_v14_).cf43, globalState)
            if source16_.is_DC4:
                d_963___mcc_h3_ = source16_.cf8
                d_964_cf8_ = d_963___mcc_h3_
                rhs167_ = (self).f34
                rhs168_ = (d_944_v3_)[default__.safeIndex((d_943_v2_).f36, len(d_944_v3_))]
                lhs160_ = globalState
                lhs161_ = globalState
                lhs160_.f19 = rhs167_
                lhs161_.f18 = rhs168_
                d_965_v15_: _dafny.Seq
                d_965_v15_ = _dafny.SeqWithoutIsStrInference([len(d_942_v1_), (d_943_v2_).f36])
                index159_ = default__.safeIndex(252, (d_939_v0_).length(0))
                (d_939_v0_)[index159_] = len((d_965_v15_) + (d_965_v15_))
                index160_ = default__.safeIndex(252, (d_939_v0_).length(0))
                (d_939_v0_)[index160_] = p0
                (globalState).f11 = d_942_v1_
                index161_ = default__.safeIndex(252, (d_939_v0_).length(0))
                (d_939_v0_)[index161_] = ((d_943_v2_).f36) - (p0)
            elif source16_.is_DC3:
                d_966___mcc_h4_ = source16_.cf7
                d_967_cf7_ = d_966___mcc_h4_
                d_961_v13_ = d_961_v13_
                d_968_v17_: _dafny.Set
                d_968_v17_ = _dafny.Set({d_947_cf8_, d_947_cf8_, (self).f33})
                d_969_v18_: _dafny.Map
                d_969_v18_ = _dafny.Map({d_968_v17_: (d_943_v2_).f35})
                d_970_v19_: _dafny.Map
                def iife91_():
                    coll41_ = _dafny.Map()
                    compr_41_: _dafny.Set
                    for compr_41_ in (d_969_v18_).keys.Elements:
                        d_971_v16_: _dafny.Set = compr_41_
                        if (d_971_v16_) in (d_969_v18_):
                            coll41_[d_971_v16_] = (d_943_v2_).f36
                    return _dafny.Map(coll41_)
                d_970_v19_ = _dafny.Map({(d_943_v2_).f36: len(iife91_()
                )})
                def iife92_():
                    coll42_ = _dafny.Map()
                    compr_42_: int
                    for compr_42_ in _dafny.IntegerRange(73, 979):
                        d_972_v20_: int = compr_42_
                        if ((73) <= (d_972_v20_)) and ((d_972_v20_) < (979)):
                            coll42_[default__.safeModuloInt(d_972_v20_, p0)] = len(_dafny.SeqWithoutIsStrInference([p0]))
                    return _dafny.Map(coll42_)
                rhs169_ = (self).f33
                rhs170_ = (iife92_()
                ).set(551, len((d_942_v1_).set(default__.safeIndex((d_943_v2_).f36, len(d_942_v1_)), d_949_v6_)))
                lhs162_ = globalState
                lhs162_.f28 = rhs169_
                d_970_v19_ = rhs170_
                d_973_v21_: _dafny.Map
                d_973_v21_ = _dafny.Map({(d_967_cf7_) + (d_967_cf7_): p0})
                d_973_v21_ = _dafny.Map({d_967_cf7_: -239})
                (globalState).f20 = ((default__.fm7((d_943_v2_).f36, globalState)) + (d_942_v1_)).set(default__.safeIndex((0) - ((0) - (((d_943_v2_).f36) * ((d_943_v2_).f36))), len((default__.fm7((d_943_v2_).f36, globalState)) + (d_942_v1_))), d_949_v6_)
            elif True:
                d_974___mcc_h5_ = source16_.cf9
                d_975_cf9_ = d_974___mcc_h5_
                d_976_v22_: C1
                nw166_ = C1()
                nw166_.ctor__((d_943_v2_).f35)
                d_976_v22_ = nw166_
                d_977_v23_: _dafny.Set
                d_977_v23_ = _dafny.Set({d_939_v0_})
                d_978_v24_: C2
                nw167_ = C2()
                nw167_.ctor__(len((d_944_v3_) + (d_944_v3_)), d_977_v23_, (d_943_v2_).f35, d_942_v1_)
                d_978_v24_ = nw167_
                (globalState).f28 = d_947_cf8_
                d_979_v25_: _dafny.Seq
                d_979_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jj"))])
                d_979_v25_ = (d_979_v25_).set(default__.safeIndex((d_943_v2_).f36, len(d_979_v25_)), (d_942_v1_) + (d_942_v1_))
            index162_ = default__.safeIndex(477, (d_950_v7_).length(0))
            (d_950_v7_)[index162_] = False
            index163_ = default__.safeIndex(477, (d_950_v7_).length(0))
            (d_950_v7_)[index163_] = (d_943_v2_).f35
            (globalState).f14 = p0
        elif source15_.is_DC3:
            d_980___mcc_h1_ = source15_.cf7
            d_981_cf7_ = d_980___mcc_h1_
            (globalState).f28 = (d_943_v2_).f35
            r0 = ((d_943_v2_).f35) == ((d_943_v2_).f35)
            d_982_v26_: _dafny.Array
            def lambda44_(d_983_v2_):
                def lambda45_(d_984_i3_):
                    return (d_983_v2_).f35

                return lambda45_

            init27_ = lambda44_(d_943_v2_)
            nw168_ = _dafny.Array(None, 24)
            for i0_27_ in range(nw168_.length(0)):
                nw168_[i0_27_] = init27_(i0_27_)
            d_982_v26_ = nw168_
            index164_ = default__.safeIndex(376, (d_982_v26_).length(0))
            (d_982_v26_)[index164_] = (d_943_v2_).f35
            d_985_v27_: _dafny.MultiSet
            d_985_v27_ = _dafny.MultiSet([d_939_v0_, d_939_v0_])
            index165_ = default__.safeIndex(376, (d_982_v26_).length(0))
            (d_982_v26_)[index165_] = ((d_985_v27_).intersection(d_985_v27_)) == (d_985_v27_)
            (globalState).f20 = (d_942_v1_) + ((d_942_v1_) + (d_942_v1_))
        elif True:
            d_986___mcc_h2_ = source15_.cf9
            d_987_cf9_ = d_986___mcc_h2_
            d_988_v28_: _dafny.Seq
            d_989_v29_: bool
            out14_: _dafny.Seq
            out15_: bool
            out14_, out15_ = (d_943_v2_).m2(not((d_943_v2_).f35), globalState)
            d_988_v28_ = out14_
            d_989_v29_ = out15_
            d_990_v30_: _dafny.MultiSet
            d_990_v30_ = _dafny.MultiSet([-282, (self).f34])
            d_991_v31_: D3
            d_991_v31_ = D3_DC10(d_990_v30_)
            d_992_v32_: _dafny.Seq
            d_992_v32_ = _dafny.SeqWithoutIsStrInference([D3_DC10(d_990_v30_), d_991_v31_, d_991_v31_])
            d_993_v33_: _dafny.Map
            d_993_v33_ = _dafny.Map({26: d_992_v32_})
            d_994_v34_: str
            d_994_v34_ = _dafny.CodePoint('e')
            def iife93_():
                coll43_ = _dafny.Map()
                compr_43_: int
                for compr_43_ in _dafny.IntegerRange(717, 624):
                    d_995_v35_: int = compr_43_
                    if ((717) <= (d_995_v35_)) and ((d_995_v35_) < (624)):
                        coll43_[(d_995_v35_) + (p0)] = _dafny.SeqWithoutIsStrInference([d_991_v31_ for d_996_i4_ in range(default__.abs(888))])
                return _dafny.Map(coll43_)
            d_993_v33_ = ((d_993_v33_ if default__.fm19(d_994_v34_, globalState) else iife93_()
            )).set((0) - (((d_943_v2_).f36) - (p0)), (d_992_v32_) + (d_992_v32_))
            d_997_v36_: T0
            nw169_ = C1()
            nw169_.ctor__((d_943_v2_).f35)
            d_997_v36_ = nw169_
            d_998_v37_: _dafny.Array
            nw170_ = _dafny.Array(None, 18)
            nw170_[int(0)] = d_997_v36_
            nw170_[int(1)] = d_997_v36_
            nw170_[int(2)] = d_997_v36_
            nw170_[int(3)] = d_997_v36_
            nw170_[int(4)] = d_997_v36_
            nw170_[int(5)] = d_997_v36_
            nw170_[int(6)] = d_997_v36_
            nw170_[int(7)] = d_997_v36_
            nw170_[int(8)] = d_997_v36_
            nw170_[int(9)] = d_997_v36_
            nw170_[int(10)] = d_997_v36_
            nw170_[int(11)] = d_997_v36_
            nw170_[int(12)] = d_997_v36_
            nw170_[int(13)] = d_997_v36_
            nw170_[int(14)] = d_997_v36_
            nw170_[int(15)] = (d_997_v36_ if default__.fm19(_dafny.CodePoint('e'), globalState) else d_997_v36_)
            nw170_[int(16)] = d_997_v36_
            nw170_[int(17)] = d_997_v36_
            d_998_v37_ = nw170_
            d_999_v38_: _dafny.Map
            d_999_v38_ = _dafny.Map({(d_943_v2_).f35: d_998_v37_})
            d_998_v37_ = ((d_999_v38_)[not(True)] if (not(True)) in (d_999_v38_) else d_998_v37_)
            if not ((d_943_v2_).f35) or (not(((d_943_v2_).f36) < (872))):
                (globalState).f2 = (d_943_v2_).f36
                d_1000_v39_: _dafny.Array
                nw171_ = _dafny.Array(False, 6)
                d_1000_v39_ = nw171_
                index166_ = default__.safeIndex(62, (d_1000_v39_).length(0))
                (d_1000_v39_)[index166_] = d_989_v29_
                d_1001_v40_: C0
                nw172_ = C0()
                nw172_.ctor__((d_943_v2_).f35)
                d_1001_v40_ = nw172_
                d_1002_v41_: _dafny.Seq
                d_1002_v41_ = _dafny.SeqWithoutIsStrInference([d_1001_v40_, d_1001_v40_, d_1001_v40_])
                d_1003_v42_: _dafny.MultiSet
                d_1003_v42_ = _dafny.MultiSet([d_1002_v41_, d_1002_v41_, (d_1002_v41_) + (d_1002_v41_)])
                index167_ = default__.safeIndex(797, (d_939_v0_).length(0))
                (d_939_v0_)[index167_] = (d_943_v2_).f36
                index168_ = default__.safeIndex(62, (d_1000_v39_).length(0))
                index169_ = default__.safeIndex(797, (d_939_v0_).length(0))
                rhs171_ = not(d_1001_v40_.f37)
                rhs172_ = d_1000_v39_
                rhs173_ = (d_944_v3_)[default__.safeIndex((d_943_v2_).f36, len(d_944_v3_))]
                rhs174_ = d_1003_v42_
                rhs175_ = (0) - ((d_943_v2_).f36)
                lhs163_ = d_1000_v39_
                lhs164_ = default__.safeIndex(62, (d_1000_v39_).length(0))
                lhs165_ = globalState
                lhs166_ = d_939_v0_
                lhs167_ = default__.safeIndex(797, (d_939_v0_).length(0))
                lhs163_[lhs164_] = rhs171_
                d_1000_v39_ = rhs172_
                lhs165_.f28 = rhs173_
                d_1003_v42_ = rhs174_
                lhs166_[lhs167_] = rhs175_
                (globalState).f19 = (d_943_v2_).f36
                d_1004_v43_: _dafny.Set
                d_1004_v43_ = _dafny.Set({(d_943_v2_).f35})
                d_1005_v44_: _dafny.Seq
                d_1005_v44_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_989_v29_})])
                d_1006_v45_: _dafny.Map
                d_1006_v45_ = _dafny.Map({(d_943_v2_).f36: d_1004_v43_})
                d_1007_v46_: _dafny.Array
                nw173_ = _dafny.Array(None, 22)
                nw173_[int(0)] = d_1004_v43_
                nw173_[int(1)] = d_1004_v43_
                nw173_[int(2)] = d_1004_v43_
                nw173_[int(3)] = d_1004_v43_
                nw173_[int(4)] = d_1004_v43_
                nw173_[int(5)] = default__.fm22((self).f34, (self).f33, d_994_v34_, globalState)
                nw173_[int(6)] = (d_1004_v43_) | (d_1004_v43_)
                nw173_[int(7)] = (_dafny.Set({True})) | (d_1004_v43_)
                nw173_[int(8)] = (d_1004_v43_).intersection(_dafny.Set({d_1001_v40_.f37, (d_943_v2_).f35, d_1001_v40_.f37}))
                nw173_[int(9)] = d_1004_v43_
                nw173_[int(10)] = _dafny.Set({(d_1000_v39_)[default__.safeIndex(62, (d_1000_v39_).length(0))]})
                nw173_[int(11)] = (d_1005_v44_)[default__.safeIndex((d_943_v2_).f36, len(d_1005_v44_))]
                nw173_[int(12)] = d_1004_v43_
                nw173_[int(13)] = d_1004_v43_
                nw173_[int(14)] = _dafny.Set({d_1001_v40_.f37})
                nw173_[int(15)] = d_1004_v43_
                nw173_[int(16)] = (default__.fm22((d_943_v2_).f36, d_989_v29_, d_994_v34_, globalState)) - (d_1004_v43_)
                nw173_[int(17)] = _dafny.Set({(d_943_v2_).f35})
                nw173_[int(18)] = _dafny.Set({d_1001_v40_.f37})
                nw173_[int(19)] = (d_1004_v43_) - (((d_1006_v45_)[(d_943_v2_).f36] if ((d_943_v2_).f36) in (d_1006_v45_) else d_1004_v43_))
                nw173_[int(20)] = d_1004_v43_
                nw173_[int(21)] = (d_1004_v43_) - (d_1004_v43_)
                d_1007_v46_ = nw173_
                d_1007_v46_ = d_1007_v46_
                (globalState).f18 = (d_943_v2_).f35
            elif True:
                (globalState).f28 = default__.fm6(-894, globalState)
                d_994_v34_ = d_994_v34_
                d_1008_v47_: _dafny.Set
                d_1008_v47_ = _dafny.Set({(d_943_v2_).f35, (d_943_v2_).f35, not((d_943_v2_).f35), (d_943_v2_).f35, (d_943_v2_).f35})
                d_1008_v47_ = d_1008_v47_
                d_1009_v48_: _dafny.Array
                def lambda46_(d_1010_v2_, d_1011_p0_):
                    def lambda47_(d_1012_i5_):
                        return (_dafny.SeqWithoutIsStrInference([(d_1010_v2_).f36])) + (_dafny.SeqWithoutIsStrInference([d_1011_p0_, 814, d_1011_p0_, (d_1010_v2_).f36, len(_dafny.Map({(d_1010_v2_).f36: (d_1010_v2_).f36}))]))

                    return lambda47_

                init28_ = lambda46_(d_943_v2_, p0)
                nw174_ = _dafny.Array(None, 18)
                for i0_28_ in range(nw174_.length(0)):
                    nw174_[i0_28_] = init28_(i0_28_)
                d_1009_v48_ = nw174_
                d_1013_v49_: _dafny.Seq
                d_1013_v49_ = _dafny.SeqWithoutIsStrInference([(self).f34, (0) - ((d_943_v2_).f36)])
                index170_ = default__.safeIndex(902, (d_1009_v48_).length(0))
                (d_1009_v48_)[index170_] = d_1013_v49_
                index171_ = default__.safeIndex(902, (d_1009_v48_).length(0))
                (d_1009_v48_)[index171_] = d_1013_v49_
                (globalState).f21 = (d_1013_v49_)[default__.safeIndex(-844, len(d_1013_v49_))]
        (globalState).f21 = (p0) + (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxjnf"))), (self).f34))
        (globalState).f21 = ((d_943_v2_).f36) * (default__.safeModuloInt((d_943_v2_).f36, (d_943_v2_).f36))
        d_1014_v50_: str
        d_1014_v50_ = _dafny.CodePoint('g')
        d_1015_v51_: _dafny.Set
        d_1015_v51_ = _dafny.Set({d_1014_v50_, default__.fm42(True, not((self).f33), globalState), (d_1014_v50_ if (self).f33 else d_1014_v50_)})
        d_1015_v51_ = _dafny.Set({d_1014_v50_, d_1014_v50_, default__.fm42((self).f33, (d_943_v2_).f35, globalState), d_1014_v50_, d_1014_v50_})
        r0 = (d_943_v2_).f35
        return r0

    @property
    def f33(self):
        return self._f33
    @property
    def f34(self):
        return self._f34

class C5(T1, T0):
    def  __init__(self):
        self._f30: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f29: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f30(self):
        return self._f30
    @f30.setter
    def f30(self, value):
        self._f30 = value
    @property
    def f29(self):
        return self._f29
    def ctor__(self, f29, f30):
        (self)._f29 = f29
        (self).f30 = f30

    def fm3(self, p0, globalState):
        if ((self).f29) == ((self).f29):
            return len((self.f30) + (self.f30))
        elif True:
            return (0) - (len(self.f30))

    def fm4(self, globalState):
        return ((((D2_DC6(_dafny.MultiSet([_dafny.CodePoint('p')]))).cf10).cardinality if (self).f29 else 255)) * (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxmebkric")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qe")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1016_i0_ in range(default__.abs(-273))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugd")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1017_i1_ in range(default__.abs(515))])])))

    def m2(self, p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_1018_v0_: _dafny.Array
        nw175_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
        d_1018_v0_ = nw175_
        d_1019_v1_: _dafny.Array
        nw176_ = _dafny.Array(None, 13)
        nw176_[int(0)] = (self).f29
        nw176_[int(1)] = (self).f29
        nw176_[int(2)] = (self).f29
        nw176_[int(3)] = not(False)
        nw176_[int(4)] = (self).f29
        nw176_[int(5)] = p0
        nw176_[int(6)] = (self).f29
        nw176_[int(7)] = p0
        nw176_[int(8)] = p0
        nw176_[int(9)] = (self).f29
        nw176_[int(10)] = (self).f29
        nw176_[int(11)] = (self).f29
        nw176_[int(12)] = (self).f29
        d_1019_v1_ = nw176_
        d_1020_v2_: int
        d_1020_v2_ = 70
        d_1021_v3_: _dafny.Seq
        d_1021_v3_ = _dafny.SeqWithoutIsStrInference([(self).f29])
        d_1022_v4_: _dafny.Map
        d_1022_v4_ = _dafny.Map({d_1018_v0_: (0) - ((len((_dafny.SeqWithoutIsStrInference([d_1019_v1_])).set(default__.safeIndex(d_1020_v2_, len(_dafny.SeqWithoutIsStrInference([d_1019_v1_]))), d_1019_v1_))) * ((self).fm3(d_1021_v3_, globalState)))})
        d_1022_v4_ = (d_1022_v4_).set(d_1018_v0_, (self).fm4(globalState))
        d_1023_v5_: _dafny.Seq
        d_1023_v5_ = _dafny.SeqWithoutIsStrInference([d_1020_v2_, d_1020_v2_])
        hi5_ = (d_1023_v5_)[default__.safeIndex(len(d_1023_v5_), len(d_1023_v5_))]
        for d_1024_i0_ in range((d_1020_v2_) + (len(_dafny.SeqWithoutIsStrInference([d_1020_v2_ for d_1028_i1_ in range(default__.abs(453))]))), hi5_):
            d_1025_v6_: _dafny.MultiSet
            d_1025_v6_ = _dafny.MultiSet([(self).f29, (self).f29])
            rhs176_ = (d_1025_v6_).issubset(d_1025_v6_)
            rhs177_ = (d_1024_i0_) < (d_1020_v2_)
            lhs168_ = globalState
            lhs169_ = globalState
            lhs168_.f18 = rhs176_
            lhs169_.f18 = rhs177_
            (globalState).f18 = False
            d_1026_v7_: _dafny.Array
            nw177_ = _dafny.Array(int(0), 22)
            d_1026_v7_ = nw177_
            d_1027_v8_: _dafny.Map
            d_1027_v8_ = _dafny.Map({(self).f29: d_1026_v7_})
            d_1027_v8_ = (d_1027_v8_).set(((d_1025_v6_).cardinality) >= (d_1020_v2_), d_1026_v7_)
            index172_ = default__.safeIndex(700, (d_1026_v7_).length(0))
            (d_1026_v7_)[index172_] = 967
            index173_ = default__.safeIndex(700, (d_1026_v7_).length(0))
            (d_1026_v7_)[index173_] = (d_1020_v2_) * (d_1020_v2_)
        (globalState).f21 = d_1020_v2_
        hi6_ = d_1020_v2_
        for d_1029_i2_ in range(d_1020_v2_, hi6_):
            d_1030_v9_: _dafny.Array
            def lambda48_(d_1031_v3_):
                def lambda49_(d_1032_i3_):
                    return _dafny.MultiSet([D1_DC3(d_1031_v3_)])

                return lambda49_

            init29_ = lambda48_(d_1021_v3_)
            nw178_ = _dafny.Array(None, 21)
            for i0_29_ in range(nw178_.length(0)):
                nw178_[i0_29_] = init29_(i0_29_)
            d_1030_v9_ = nw178_
            d_1030_v9_ = d_1030_v9_
            (globalState).f13 = p0
            index174_ = default__.safeIndex(952, (d_1019_v1_).length(0))
            (d_1019_v1_)[index174_] = (self).f29
            d_1033_v10_: _dafny.Map
            d_1033_v10_ = _dafny.Map({d_1029_i2_: p0})
            index175_ = default__.safeIndex(952, (d_1019_v1_).length(0))
            def iife94_():
                coll44_ = _dafny.Map()
                compr_44_: int
                for compr_44_ in _dafny.IntegerRange(228, 630):
                    d_1034_v11_: int = compr_44_
                    if ((228) <= (d_1034_v11_)) and ((d_1034_v11_) < (630)):
                        coll44_[(d_1034_v11_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spsbhq"))))] = d_1029_i2_
                return _dafny.Map(coll44_)
            def iife95_():
                coll45_ = _dafny.Map()
                compr_45_: int
                for compr_45_ in _dafny.IntegerRange(228, 630):
                    d_1035_v11_: int = compr_45_
                    if ((228) <= (d_1035_v11_)) and ((d_1035_v11_) < (630)):
                        coll45_[(d_1035_v11_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spsbhq"))))] = d_1029_i2_
                return _dafny.Map(coll45_)
            (d_1019_v1_)[index175_] = (((d_1033_v10_)[len((iife94_()
            ).set(416, d_1029_i2_))] if (len((iife95_()
            ).set(416, d_1029_i2_))) in (d_1033_v10_) else p0) if p0 else (d_1029_i2_) == (d_1020_v2_))
            if p0:
                d_1036_v12_: _dafny.Array
                nw179_ = _dafny.Array(int(0), 17)
                d_1036_v12_ = nw179_
                d_1037_v13_: D1
                d_1037_v13_ = D1_DC3(d_1021_v3_)
                d_1038_v14_: D1
                d_1038_v14_ = D1_DC5(d_1037_v13_)
                d_1039_v15_: D1
                d_1039_v15_ = D1_DC5(D1_DC5(d_1037_v13_))
                d_1040_v16_: D1
                d_1040_v16_ = D1_DC5(d_1037_v13_)
                d_1041_v17_: _dafny.Map
                d_1041_v17_ = _dafny.Map({d_1036_v12_: d_1040_v16_})
                d_1041_v17_ = ((d_1041_v17_ if (self).f29 else (d_1041_v17_).set(d_1036_v12_, d_1040_v16_))).set(d_1036_v12_, d_1040_v16_)
                d_1042_v18_: _dafny.Set
                d_1042_v18_ = _dafny.Set({p0, p0, True})
                (globalState).f21 = ((d_1029_i2_) + (len(d_1042_v18_))) + (556)
                d_1043_v19_: _dafny.Map
                d_1043_v19_ = _dafny.Map({d_1036_v12_: (d_1019_v1_)[default__.safeIndex(952, (d_1019_v1_).length(0))]})
                d_1044_v20_: _dafny.Map
                d_1044_v20_ = _dafny.Map({len(d_1043_v19_): d_1020_v2_})
                d_1044_v20_ = (d_1044_v20_).set(len(self.f30), d_1020_v2_)
                d_1045_v21_: _dafny.Map
                d_1045_v21_ = _dafny.Map({(self).f29: p0})
                d_1044_v20_ = (d_1044_v20_).set((d_1020_v2_) + (len((_dafny.SeqWithoutIsStrInference([len(d_1044_v20_) for d_1046_i4_ in range(default__.abs(375))])).set(default__.safeIndex(len(d_1045_v21_), len(_dafny.SeqWithoutIsStrInference([len(d_1044_v20_) for d_1047_i4_ in range(default__.abs(375))]))), d_1029_i2_))), (d_1029_i2_) - (826))
                d_1048_v22_: _dafny.Array
                def lambda50_(d_1049_p0_):
                    def lambda51_(d_1050_i5_):
                        return d_1049_p0_

                    return lambda51_

                init30_ = lambda50_(p0)
                nw180_ = _dafny.Array(None, 19)
                for i0_30_ in range(nw180_.length(0)):
                    nw180_[i0_30_] = init30_(i0_30_)
                d_1048_v22_ = nw180_
                d_1051_v23_: D0
                d_1051_v23_ = D0_DC1(d_1048_v22_, ((d_1044_v20_)[d_1029_i2_] if (d_1029_i2_) in (d_1044_v20_) else d_1029_i2_), (0) - (d_1029_i2_))
                d_1048_v22_ = (d_1051_v23_).cf1
            elif True:
                d_1052_v24_: _dafny.Array
                nw181_ = _dafny.Array(False, 20)
                d_1052_v24_ = nw181_
                d_1053_v25_: _dafny.MultiSet
                d_1053_v25_ = _dafny.MultiSet([p0])
                d_1052_v24_ = (d_1052_v24_ if (_dafny.MultiSet(d_1021_v3_)).isdisjoint(d_1053_v25_) else d_1052_v24_)
                index176_ = default__.safeIndex(952, (d_1019_v1_).length(0))
                (d_1019_v1_)[index176_] = (d_1029_i2_) != (d_1020_v2_)
                d_1054_v26_: _dafny.Map
                d_1054_v26_ = _dafny.Map({(d_1019_v1_)[default__.safeIndex(952, (d_1019_v1_).length(0))]: d_1020_v2_})
                (globalState).f21 = (0) - ((d_1020_v2_) - (default__.safeModuloInt(d_1020_v2_, ((d_1054_v26_)[(self).f29] if ((self).f29) in (d_1054_v26_) else d_1020_v2_))))
                r1 = (d_1019_v1_)[default__.safeIndex(952, (d_1019_v1_).length(0))]
                d_1055_v27_: C0
                nw182_ = C0()
                nw182_.ctor__(p0)
                d_1055_v27_ = nw182_
        d_1056_v28_: D9
        d_1056_v28_ = D9_DC27(d_1023_v5_)
        d_1057_v29_: _dafny.Array
        nw183_ = _dafny.Array(None, 6)
        nw183_[int(0)] = (d_1056_v28_).cf40
        nw183_[int(1)] = d_1023_v5_
        nw183_[int(2)] = d_1023_v5_
        nw183_[int(3)] = (d_1023_v5_) + (d_1023_v5_)
        nw183_[int(4)] = (d_1023_v5_ if (self).f29 else _dafny.SeqWithoutIsStrInference([566]))
        nw183_[int(5)] = _dafny.SeqWithoutIsStrInference([d_1020_v2_, d_1020_v2_])
        d_1057_v29_ = nw183_
        index177_ = default__.safeIndex(829, (d_1057_v29_).length(0))
        (d_1057_v29_)[index177_] = d_1023_v5_
        index178_ = default__.safeIndex(829, (d_1057_v29_).length(0))
        (d_1057_v29_)[index178_] = ((d_1023_v5_ if p0 else d_1023_v5_) if (self).f29 else d_1023_v5_)
        d_1058_v30_: _dafny.MultiSet
        d_1058_v30_ = _dafny.MultiSet([(p0) == (True), (d_1021_v3_)[default__.safeIndex(len(self.f30), len(d_1021_v3_))], (self).f29])
        d_1058_v30_ = d_1058_v30_
        d_1059_v31_: str
        d_1059_v31_ = _dafny.CodePoint('e')
        d_1060_v32_: D2
        d_1060_v32_ = D2_DC8(d_1059_v31_, (self).f29)
        d_1061_v33_: _dafny.Seq
        d_1061_v33_ = _dafny.SeqWithoutIsStrInference([d_1060_v32_])
        d_1062_v34_: _dafny.Seq
        d_1062_v34_ = _dafny.SeqWithoutIsStrInference([default__.fm35(d_1061_v33_, 654, globalState), _dafny.SeqWithoutIsStrInference([(d_1021_v3_)[default__.safeIndex(d_1020_v2_, len(d_1021_v3_))], (self).f29, p0]), d_1021_v3_])
        r0 = (d_1062_v34_) + (d_1062_v34_)
        r1 = (self).f29
        return r0, r1

    def m3(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        d_1063_v0_: D4
        d_1063_v0_ = D4_DC15(p2, p2)
        d_1064_i0_: int
        d_1064_i0_ = 0
        with _dafny.label("5"):
            pat_let_tv22_ = p0
            def lambda52_(source17_):
                if source17_.is_DC14:
                    d_1069___mcc_h0_ = source17_.cf23
                    d_1070_cf23_ = d_1069___mcc_h0_
                    return (self).f29
                elif source17_.is_DC15:
                    d_1071___mcc_h1_ = source17_.cf24
                    d_1072___mcc_h2_ = source17_.cf25
                    d_1073_cf25_ = d_1072___mcc_h2_
                    d_1074_cf24_ = d_1071___mcc_h1_
                    return pat_let_tv22_
                elif True:
                    d_1075___mcc_h3_ = source17_.cf22
                    d_1076_cf22_ = d_1075___mcc_h3_
                    return (self).f29

            while lambda52_(d_1063_v0_):
                with _dafny.c_label("5"):
                    if (d_1064_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_1064_i0_ = (d_1064_i0_) + (1)
                    rhs178_ = (self).f29
                    rhs179_ = (self).f29
                    lhs170_ = globalState
                    lhs171_ = globalState
                    lhs170_.f13 = rhs178_
                    lhs171_.f15 = rhs179_
                    d_1065_v1_: _dafny.Array
                    nw184_ = _dafny.Array(None, 12)
                    d_1065_v1_ = nw184_
                    d_1066_v2_: T0
                    nw185_ = C4()
                    nw185_.ctor__(p0, p2)
                    d_1066_v2_ = nw185_
                    index179_ = default__.safeIndex(565, (d_1065_v1_).length(0))
                    (d_1065_v1_)[index179_] = d_1066_v2_
                    index180_ = default__.safeIndex(565, (d_1065_v1_).length(0))
                    rhs180_ = self.f30
                    rhs181_ = d_1066_v2_
                    rhs182_ = default__.fm6(p2, globalState)
                    rhs183_ = p2
                    lhs172_ = d_1065_v1_
                    lhs173_ = default__.safeIndex(565, (d_1065_v1_).length(0))
                    lhs174_ = globalState
                    lhs175_ = globalState
                    r0 = rhs180_
                    lhs172_[lhs173_] = rhs181_
                    lhs174_.f18 = rhs182_
                    lhs175_.f19 = rhs183_
                    (globalState).f2 = p2
                    d_1067_v3_: D10
                    d_1067_v3_ = D10_DC33(p0, p2)
                    d_1068_v4_: C3
                    nw186_ = C3()
                    nw186_.ctor__((d_1067_v3_).cf54, p2, (p2) == (285), self.f30)
                    d_1068_v4_ = nw186_
                    pass
            pass
        d_1077_v5_: _dafny.Seq
        d_1077_v5_ = _dafny.SeqWithoutIsStrInference([p0, (self).f29])
        d_1078_v6_: _dafny.Seq
        d_1078_v6_ = _dafny.SeqWithoutIsStrInference([d_1077_v5_, d_1077_v5_])
        d_1079_v7_: D1
        d_1079_v7_ = D1_DC3((d_1078_v6_)[default__.safeIndex(p2, len(d_1078_v6_))])
        d_1080_v8_: _dafny.Map
        d_1080_v8_ = _dafny.Map({False: d_1079_v7_})
        d_1081_v9_: D5
        d_1081_v9_ = D5_DC16(d_1080_v8_)
        pat_let_tv23_ = d_1080_v8_
        def iife96_(_pat_let25_0):
            def iife97_(d_1082_dt__update__tmp_h0_):
                def iife98_(_pat_let26_0):
                    def iife99_(d_1083_dt__update_hcf26_h0_):
                        return D5_DC16(d_1083_dt__update_hcf26_h0_)
                    return iife99_(_pat_let26_0)
                return iife98_(pat_let_tv23_)
            return iife97_(_pat_let25_0)
        source18_ = iife96_(d_1081_v9_)
        if source18_.is_DC17:
            d_1084___mcc_h4_ = source18_.cf27
            d_1085_cf27_ = d_1084___mcc_h4_
            (globalState).f18 = (self).f29
            d_1086_v10_: str
            d_1086_v10_ = _dafny.CodePoint('r')
            d_1087_v11_: _dafny.MultiSet
            d_1087_v11_ = _dafny.MultiSet([(self).f29])
            (globalState).f20 = ((self.f30).set(default__.safeIndex(p2, len(self.f30)), d_1086_v10_)).set(default__.safeIndex((d_1087_v11_).cardinality, len((self.f30).set(default__.safeIndex(p2, len(self.f30)), d_1086_v10_))), (d_1086_v10_ if (self).f29 else d_1086_v10_))
            d_1088_v12_: _dafny.Map
            d_1088_v12_ = _dafny.Map({p0: not(default__.fm6(p2, globalState))})
            d_1089_v13_: _dafny.Seq
            d_1089_v13_ = _dafny.SeqWithoutIsStrInference([p2])
            d_1090_v14_: _dafny.Seq
            d_1090_v14_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({not((self).f29): 342})), len((d_1088_v12_ if (self).f29 else default__.fm32((self).f29, (self).f29, d_1089_v13_, globalState)))])
            (globalState).f7 = d_1089_v13_
            (globalState).f1 = default__.safeDivisionInt(p2, (0) - (p2))
        elif source18_.is_DC18:
            (globalState).f18 = (self).f29
            d_1091_v15_: _dafny.Seq
            d_1091_v15_ = _dafny.SeqWithoutIsStrInference([p2])
            d_1092_v16_: _dafny.Map
            d_1092_v16_ = _dafny.Map({p2: d_1091_v15_})
            d_1093_v17_: D9
            d_1093_v17_ = D9_DC27(((d_1092_v16_)[p2] if (p2) in (d_1092_v16_) else default__.fm44((self).f29, p2, p0, p0, globalState)))
            d_1093_v17_ = D9_DC27(d_1091_v15_)
            d_1094_v18_: D0
            d_1094_v18_ = D0_DC2(_dafny.MultiSet([(self).f29]), self.f30, True)
            d_1095_v19_: str
            d_1095_v19_ = _dafny.CodePoint('u')
            pat_let_tv24_ = d_1095_v19_
            pat_let_tv25_ = p2
            pat_let_tv26_ = globalState
            def iife100_(_pat_let27_0):
                def iife101_(d_1096_dt__update__tmp_h1_):
                    def iife102_(_pat_let28_0):
                        def iife103_(d_1097_dt__update_hcf4_h0_):
                            return D0_DC2(d_1097_dt__update_hcf4_h0_, (d_1096_dt__update__tmp_h1_).cf5, (d_1096_dt__update__tmp_h1_).cf6)
                        return iife103_(_pat_let28_0)
                    return iife102_(default__.fm39(pat_let_tv24_, pat_let_tv25_, pat_let_tv26_))
                return iife101_(_pat_let27_0)
            source19_ = iife100_(d_1094_v18_)
            if source19_.is_DC0:
                d_1098___mcc_h10_ = source19_.cf0
                d_1099_cf0_ = d_1098___mcc_h10_
                (globalState).f20 = self.f30
                d_1100_v20_: _dafny.Array
                nw187_ = _dafny.Array(int(0), 18)
                d_1100_v20_ = nw187_
                d_1101_v21_: C2
                nw188_ = C2()
                nw188_.ctor__(p2, (_dafny.Set({d_1100_v20_, d_1100_v20_})).intersection(_dafny.Set({d_1100_v20_, d_1100_v20_, d_1100_v20_, d_1100_v20_, d_1100_v20_})), p0, (self.f30) + (self.f30))
                d_1101_v21_ = nw188_
                d_1102_v22_: _dafny.Array
                nw189_ = _dafny.Array(False, 9)
                d_1102_v22_ = nw189_
                index181_ = default__.safeIndex(692, (d_1102_v22_).length(0))
                (d_1102_v22_)[index181_] = ((self).f29) or (d_1099_cf0_)
                index182_ = default__.safeIndex(692, (d_1102_v22_).length(0))
                (d_1102_v22_)[index182_] = d_1099_cf0_
                (globalState).f1 = (len(_dafny.Map({(0) - ((d_1101_v21_).f38): p2}))) * ((d_1101_v21_).f38)
            elif source19_.is_DC1:
                d_1103___mcc_h11_ = source19_.cf1
                d_1104___mcc_h12_ = source19_.cf2
                d_1105___mcc_h13_ = source19_.cf3
                d_1106_cf3_ = d_1105___mcc_h13_
                d_1107_cf2_ = d_1104___mcc_h12_
                d_1108_cf1_ = d_1103___mcc_h11_
                d_1109_v23_: _dafny.Array
                nw190_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_1109_v23_ = nw190_
                d_1110_v24_: _dafny.Array
                def lambda53_(d_1111_v15_, d_1112_cf2_):
                    def lambda54_(d_1113_i1_):
                        return (D11_DC34(_dafny.Map({d_1111_v15_: d_1112_cf2_}))).cf56

                    return lambda54_

                init31_ = lambda53_(d_1091_v15_, d_1107_cf2_)
                nw191_ = _dafny.Array(None, 17)
                for i0_31_ in range(nw191_.length(0)):
                    nw191_[i0_31_] = init31_(i0_31_)
                d_1110_v24_ = nw191_
                d_1114_v25_: _dafny.Map
                d_1114_v25_ = _dafny.Map({d_1091_v15_: d_1106_cf3_})
                index183_ = default__.safeIndex(962, (d_1110_v24_).length(0))
                (d_1110_v24_)[index183_] = ((d_1114_v25_).set(d_1091_v15_, d_1107_cf2_)).set((_dafny.SeqWithoutIsStrInference([d_1107_cf2_])).set(default__.safeIndex((0) - (d_1107_cf2_), len(_dafny.SeqWithoutIsStrInference([d_1107_cf2_]))), d_1106_cf3_), d_1107_cf2_)
                pat_let_tv27_ = d_1107_cf2_
                index184_ = default__.safeIndex(962, (d_1110_v24_).length(0))
                def iife104_(_pat_let29_0):
                    def iife105_(d_1115_dt__update__tmp_h2_):
                        def iife106_(_pat_let30_0):
                            def iife107_(d_1116_dt__update_hcf25_h0_):
                                return D4_DC15((d_1115_dt__update__tmp_h2_).cf24, d_1116_dt__update_hcf25_h0_)
                            return iife107_(_pat_let30_0)
                        return iife106_(pat_let_tv27_)
                    return iife105_(_pat_let29_0)
                rhs184_ = d_1109_v23_
                rhs185_ = default__.fm45(globalState)
                rhs186_ = iife104_(d_1063_v0_)
                rhs187_ = d_1107_cf2_
                rhs188_ = not(not((p2) == ((0) - (default__.fm9(globalState)))))
                lhs176_ = d_1110_v24_
                lhs177_ = default__.safeIndex(962, (d_1110_v24_).length(0))
                lhs178_ = globalState
                lhs179_ = globalState
                d_1109_v23_ = rhs184_
                lhs176_[lhs177_] = rhs185_
                d_1063_v0_ = rhs186_
                lhs178_.f21 = rhs187_
                lhs179_.f15 = rhs188_
                index185_ = default__.safeIndex(724, (d_1108_cf1_).length(0))
                (d_1108_cf1_)[index185_] = (self).f29
                index186_ = default__.safeIndex(724, (d_1108_cf1_).length(0))
                rhs189_ = d_1106_cf3_
                rhs190_ = (self).f29
                rhs191_ = p0
                rhs192_ = default__.fm19(d_1095_v19_, globalState)
                lhs180_ = globalState
                lhs181_ = d_1108_cf1_
                lhs182_ = default__.safeIndex(724, (d_1108_cf1_).length(0))
                lhs183_ = globalState
                lhs184_ = globalState
                lhs180_.f19 = rhs189_
                lhs181_[lhs182_] = rhs190_
                lhs183_.f28 = rhs191_
                lhs184_.f15 = rhs192_
                d_1107_cf2_ = (p2) - (d_1107_cf2_)
                (globalState).f28 = (d_1107_cf2_) <= (p2)
            elif True:
                d_1117___mcc_h14_ = source19_.cf4
                d_1118___mcc_h15_ = source19_.cf5
                d_1119___mcc_h16_ = source19_.cf6
                d_1120_cf6_ = d_1119___mcc_h16_
                d_1121_cf5_ = d_1118___mcc_h15_
                d_1122_cf4_ = d_1117___mcc_h14_
                rhs193_ = (0) - (p2)
                rhs194_ = not(d_1120_cf6_)
                rhs195_ = d_1122_cf4_
                lhs185_ = globalState
                lhs186_ = globalState
                lhs185_.f19 = rhs193_
                lhs186_.f18 = rhs194_
                d_1122_cf4_ = rhs195_
                d_1123_v26_: C3
                nw192_ = C3()
                nw192_.ctor__(not((self).f29), p2, (p0 if p0 else (self).f29), (D11_DC35(self.f30, d_1120_cf6_)).cf57)
                d_1123_v26_ = nw192_
                (globalState).f13 = (self).f29
                d_1124_v27_: _dafny.Array
                nw193_ = _dafny.Array(int(0), 7)
                d_1124_v27_ = nw193_
                d_1125_v28_: _dafny.Seq
                d_1125_v28_ = _dafny.SeqWithoutIsStrInference([d_1124_v27_, d_1124_v27_, d_1124_v27_, d_1124_v27_, d_1124_v27_])
                d_1126_v29_: _dafny.Array
                nw194_ = _dafny.Array(None, 28)
                nw194_[int(0)] = d_1124_v27_
                nw194_[int(1)] = d_1124_v27_
                nw194_[int(2)] = d_1124_v27_
                nw194_[int(3)] = d_1124_v27_
                nw194_[int(4)] = d_1124_v27_
                nw194_[int(5)] = d_1124_v27_
                nw194_[int(6)] = d_1124_v27_
                nw194_[int(7)] = d_1124_v27_
                nw194_[int(8)] = d_1124_v27_
                nw194_[int(9)] = d_1124_v27_
                nw194_[int(10)] = d_1124_v27_
                nw194_[int(11)] = d_1124_v27_
                nw194_[int(12)] = d_1124_v27_
                nw194_[int(13)] = d_1124_v27_
                nw194_[int(14)] = d_1124_v27_
                nw194_[int(15)] = (d_1124_v27_ if not((self).f29) else d_1124_v27_)
                nw194_[int(16)] = d_1124_v27_
                nw194_[int(17)] = d_1124_v27_
                nw194_[int(18)] = d_1124_v27_
                nw194_[int(19)] = d_1124_v27_
                nw194_[int(20)] = d_1124_v27_
                nw194_[int(21)] = d_1124_v27_
                nw194_[int(22)] = d_1124_v27_
                nw194_[int(23)] = d_1124_v27_
                nw194_[int(24)] = (d_1125_v28_)[default__.safeIndex(21, len(d_1125_v28_))]
                nw194_[int(25)] = d_1124_v27_
                nw194_[int(26)] = d_1124_v27_
                nw194_[int(27)] = d_1124_v27_
                d_1126_v29_ = nw194_
                d_1127_v30_: D12
                d_1127_v30_ = D12_DC38(d_1124_v27_)
                d_1128_v31_: _dafny.Map
                d_1128_v31_ = _dafny.Map({(d_1123_v26_).f35: (d_1127_v30_).cf64})
                nw195_ = _dafny.Array(None, 23)
                nw195_[int(0)] = d_1124_v27_
                nw195_[int(1)] = d_1124_v27_
                nw195_[int(2)] = d_1124_v27_
                nw195_[int(3)] = d_1124_v27_
                nw195_[int(4)] = d_1124_v27_
                nw195_[int(5)] = d_1124_v27_
                nw195_[int(6)] = d_1124_v27_
                nw195_[int(7)] = d_1124_v27_
                nw195_[int(8)] = d_1124_v27_
                nw195_[int(9)] = d_1124_v27_
                nw195_[int(10)] = d_1124_v27_
                nw195_[int(11)] = ((d_1128_v31_)[d_1120_cf6_] if (d_1120_cf6_) in (d_1128_v31_) else d_1124_v27_)
                nw195_[int(12)] = d_1124_v27_
                nw195_[int(13)] = d_1124_v27_
                nw195_[int(14)] = (d_1125_v28_)[default__.safeIndex(p2, len(d_1125_v28_))]
                nw195_[int(15)] = (d_1127_v30_).cf64
                nw195_[int(16)] = d_1124_v27_
                nw195_[int(17)] = d_1124_v27_
                nw195_[int(18)] = d_1124_v27_
                nw195_[int(19)] = d_1124_v27_
                nw195_[int(20)] = d_1124_v27_
                nw195_[int(21)] = d_1124_v27_
                nw195_[int(22)] = d_1124_v27_
                d_1126_v29_ = nw195_
            (globalState).f1 = p2
        elif source18_.is_DC19:
            d_1129___mcc_h5_ = source18_.cf28
            d_1130___mcc_h6_ = source18_.cf29
            d_1131___mcc_h7_ = source18_.cf30
            d_1132_cf30_ = d_1131___mcc_h7_
            d_1133_cf29_ = d_1130___mcc_h6_
            d_1134_cf28_ = d_1129___mcc_h5_
            (globalState).f15 = (p0) or ((d_1133_cf29_) <= (len(self.f30)))
            d_1135_v32_: _dafny.Array
            nw196_ = _dafny.Array(_dafny.Map({}), 28)
            d_1135_v32_ = nw196_
            d_1136_v33_: D2
            d_1136_v33_ = D2_DC7((self).f29)
            d_1137_v34_: _dafny.Map
            d_1137_v34_ = _dafny.Map({p2: (d_1136_v33_).cf11})
            d_1138_v35_: _dafny.Map
            d_1138_v35_ = _dafny.Map({True: d_1137_v34_})
            d_1139_v36_: _dafny.Map
            d_1139_v36_ = _dafny.Map({p0: d_1138_v35_})
            index187_ = default__.safeIndex(378, (d_1135_v32_).length(0))
            (d_1135_v32_)[index187_] = (((d_1139_v36_)[False] if (False) in (d_1139_v36_) else default__.fm46(_dafny.CodePoint('k'), self.f30, globalState))).set(not((self).f29), d_1137_v34_)
            d_1140_v37_: _dafny.MultiSet
            d_1140_v37_ = _dafny.MultiSet([(self).f29, (self).f29])
            d_1141_v38_: _dafny.Seq
            d_1141_v38_ = _dafny.SeqWithoutIsStrInference([(d_1140_v37_).cardinality])
            d_1142_v39_: _dafny.MultiSet
            d_1142_v39_ = _dafny.MultiSet([_dafny.Map({d_1134_cf28_: p0}), d_1137_v34_, d_1137_v34_, d_1137_v34_, d_1137_v34_])
            d_1143_v40_: _dafny.Set
            d_1143_v40_ = _dafny.Set({197})
            d_1144_v41_: _dafny.Array
            nw197_ = _dafny.Array(None, 9)
            nw197_[int(0)] = (d_1141_v38_)[default__.safeIndex(d_1134_cf28_, len(d_1141_v38_))]
            nw197_[int(1)] = d_1133_cf29_
            nw197_[int(2)] = p2
            nw197_[int(3)] = d_1134_cf28_
            nw197_[int(4)] = (self).fm4(globalState)
            nw197_[int(5)] = ((d_1142_v39_)[default__.fm41(18, p0, p0, not((self).f29), globalState)] if (default__.fm41(18, p0, p0, not((self).f29), globalState)) in (d_1142_v39_) else len(d_1143_v40_))
            nw197_[int(6)] = d_1133_cf29_
            nw197_[int(7)] = p2
            nw197_[int(8)] = d_1133_cf29_
            d_1144_v41_ = nw197_
            d_1145_v42_: _dafny.Seq
            d_1145_v42_ = _dafny.SeqWithoutIsStrInference([d_1144_v41_, d_1144_v41_])
            d_1146_v43_: _dafny.Map
            d_1146_v43_ = _dafny.Map({d_1145_v42_: len(self.f30)})
            index188_ = default__.safeIndex(378, (d_1135_v32_).length(0))
            rhs196_ = d_1138_v35_
            rhs197_ = d_1133_cf29_
            rhs198_ = (self.f30) + (self.f30)
            rhs199_ = ((d_1146_v43_)[d_1145_v42_] if (d_1145_v42_) in (d_1146_v43_) else default__.safeDivisionInt(p2, d_1134_cf28_))
            lhs187_ = d_1135_v32_
            lhs188_ = default__.safeIndex(378, (d_1135_v32_).length(0))
            lhs189_ = globalState
            lhs190_ = globalState
            lhs191_ = globalState
            lhs187_[lhs188_] = rhs196_
            lhs189_.f14 = rhs197_
            lhs190_.f11 = rhs198_
            lhs191_.f21 = rhs199_
            d_1147_v44_: C4
            nw198_ = C4()
            nw198_.ctor__(p0, p2)
            d_1147_v44_ = nw198_
            (globalState).f21 = (0) - (len(_dafny.SeqWithoutIsStrInference([(d_1077_v5_) + (d_1077_v5_) for d_1148_i2_ in range(default__.abs(832))])))
        elif source18_.is_DC16:
            d_1149___mcc_h8_ = source18_.cf26
            d_1150_cf26_ = d_1149___mcc_h8_
            d_1151_v45_: str
            d_1151_v45_ = _dafny.CodePoint('c')
            d_1152_v46_: _dafny.Array
            def lambda55_(d_1153_p2_):
                def lambda56_(d_1154_i3_):
                    return default__.safeModuloInt(d_1154_i3_, d_1153_p2_)

                return lambda56_

            init32_ = lambda55_(p2)
            nw199_ = _dafny.Array(None, 5)
            for i0_32_ in range(nw199_.length(0)):
                nw199_[i0_32_] = init32_(i0_32_)
            d_1152_v46_ = nw199_
            d_1155_v47_: _dafny.Map
            d_1155_v47_ = _dafny.Map({d_1151_v45_: d_1152_v46_})
            d_1156_v48_: _dafny.Array
            nw200_ = _dafny.Array(None, 16)
            nw200_[int(0)] = d_1155_v47_
            nw200_[int(1)] = d_1155_v47_
            nw200_[int(2)] = d_1155_v47_
            nw200_[int(3)] = d_1155_v47_
            nw200_[int(4)] = (_dafny.Map({d_1151_v45_: d_1152_v46_})).set(default__.fm42((self).f29, p0, globalState), d_1152_v46_)
            nw200_[int(5)] = d_1155_v47_
            nw200_[int(6)] = (d_1155_v47_).set(d_1151_v45_, d_1152_v46_)
            nw200_[int(7)] = d_1155_v47_
            nw200_[int(8)] = d_1155_v47_
            nw200_[int(9)] = (d_1155_v47_).set(d_1151_v45_, d_1152_v46_)
            nw200_[int(10)] = d_1155_v47_
            nw200_[int(11)] = (d_1155_v47_) | (_dafny.Map({d_1151_v45_: d_1152_v46_}))
            nw200_[int(12)] = d_1155_v47_
            nw200_[int(13)] = d_1155_v47_
            nw200_[int(14)] = d_1155_v47_
            nw200_[int(15)] = _dafny.Map({d_1151_v45_: d_1152_v46_})
            d_1156_v48_ = nw200_
            index189_ = default__.safeIndex(982, (d_1156_v48_).length(0))
            (d_1156_v48_)[index189_] = d_1155_v47_
            index190_ = default__.safeIndex(982, (d_1156_v48_).length(0))
            (d_1156_v48_)[index190_] = (d_1155_v47_).set(_dafny.CodePoint('d'), d_1152_v46_)
            (globalState).f15 = (self).f29
            (globalState).f18 = p0
            (globalState).f18 = (self).f29
        elif True:
            d_1157___mcc_h9_ = source18_.cf31
            d_1158_cf31_ = d_1157___mcc_h9_
            (globalState).f14 = p2
            d_1159_v49_: _dafny.Array
            nw201_ = _dafny.Array(False, 26)
            d_1159_v49_ = nw201_
            d_1160_v50_: _dafny.Seq
            d_1160_v50_ = _dafny.SeqWithoutIsStrInference([186, p2, len(d_1077_v5_)])
            d_1161_v51_: D0
            d_1161_v51_ = D0_DC1(d_1159_v49_, len(d_1077_v5_), (d_1160_v50_)[default__.safeIndex((self).fm4(globalState), len(d_1160_v50_))])
            d_1162_v52_: _dafny.Seq
            d_1162_v52_ = _dafny.SeqWithoutIsStrInference([d_1159_v49_])
            d_1163_v53_: _dafny.Array
            nw202_ = _dafny.Array(None, 29)
            nw202_[int(0)] = d_1159_v49_
            nw202_[int(1)] = d_1159_v49_
            nw202_[int(2)] = (d_1159_v49_ if p0 else d_1159_v49_)
            nw202_[int(3)] = d_1159_v49_
            nw202_[int(4)] = d_1159_v49_
            nw202_[int(5)] = d_1159_v49_
            nw202_[int(6)] = d_1159_v49_
            nw202_[int(7)] = d_1159_v49_
            nw202_[int(8)] = (d_1161_v51_).cf1
            nw202_[int(9)] = d_1159_v49_
            nw202_[int(10)] = d_1159_v49_
            nw202_[int(11)] = (d_1162_v52_)[default__.safeIndex(p2, len(d_1162_v52_))]
            nw202_[int(12)] = d_1159_v49_
            nw202_[int(13)] = d_1159_v49_
            nw202_[int(14)] = d_1159_v49_
            nw202_[int(15)] = d_1159_v49_
            nw202_[int(16)] = d_1159_v49_
            nw202_[int(17)] = (d_1159_v49_ if (self).f29 else d_1159_v49_)
            nw202_[int(18)] = d_1159_v49_
            nw202_[int(19)] = (d_1159_v49_ if p0 else d_1159_v49_)
            nw202_[int(20)] = d_1159_v49_
            nw202_[int(21)] = d_1159_v49_
            nw202_[int(22)] = d_1159_v49_
            nw202_[int(23)] = d_1159_v49_
            nw202_[int(24)] = d_1159_v49_
            nw202_[int(25)] = d_1159_v49_
            nw202_[int(26)] = d_1159_v49_
            nw202_[int(27)] = d_1159_v49_
            nw202_[int(28)] = d_1159_v49_
            d_1163_v53_ = nw202_
            d_1163_v53_ = (d_1163_v53_ if (self).f29 else d_1163_v53_)
            d_1164_v54_: T0
            nw203_ = C1()
            nw203_.ctor__((self).f29)
            d_1164_v54_ = nw203_
            d_1165_v55_: _dafny.Map
            d_1165_v55_ = _dafny.Map({d_1164_v54_: (self).f29})
            d_1166_v56_: _dafny.Map
            d_1166_v56_ = _dafny.Map({p2: (len(d_1165_v55_)) < (p2)})
            def iife108_():
                coll46_ = _dafny.Map()
                compr_46_: int
                for compr_46_ in (d_1160_v50_).Elements:
                    d_1167_v57_: int = compr_46_
                    if (d_1167_v57_) in (d_1160_v50_):
                        coll46_[(d_1167_v57_) - (p2)] = p0
                return _dafny.Map(coll46_)
            d_1166_v56_ = iife108_()
            
            d_1168_v58_: _dafny.Seq
            d_1169_v59_: bool
            out16_: _dafny.Seq
            out17_: bool
            out16_, out17_ = (d_1164_v54_).m2(p0, globalState)
            d_1168_v58_ = out16_
            d_1169_v59_ = out17_
        d_1170_v60_: _dafny.Map
        d_1170_v60_ = _dafny.Map({(self).f29: _dafny.Set({(self).f29})})
        if not ((p2) < (len(d_1170_v60_))) or (p0):
            (globalState).f21 = 744
            d_1171_v61_: _dafny.Map
            d_1171_v61_ = _dafny.Map({p2: p2})
            (globalState).f28 = (len(d_1171_v61_)) >= (p2)
            (globalState).f18 = (p2) >= ((p2) * (len(self.f30)))
            (globalState).f19 = (default__.safeModuloInt(p2, p2) if p0 else len(_dafny.Set({p2})))
            (globalState).f28 = p0
        elif True:
            d_1172_v62_: T0
            nw204_ = C4()
            nw204_.ctor__((self).f29, p2)
            d_1172_v62_ = nw204_
            d_1173_v63_: D10
            d_1173_v63_ = D10_DC30(d_1172_v62_)
            source20_ = d_1173_v63_
            if source20_.is_DC31:
                d_1174___mcc_h17_ = source20_.cf50
                d_1175___mcc_h18_ = source20_.cf51
                d_1176_cf51_ = d_1175___mcc_h18_
                d_1177_cf50_ = d_1174___mcc_h17_
                d_1178_v64_: D9
                d_1178_v64_ = D9_DC29(p2, p2, d_1176_cf51_)
                d_1178_v64_ = d_1178_v64_
                d_1179_v65_: C3
                nw205_ = C3()
                nw205_.ctor__(d_1176_cf51_, default__.safeDivisionInt(476, p2), (self).f29, self.f30)
                d_1179_v65_ = nw205_
                d_1180_v66_: _dafny.Map
                d_1180_v66_ = _dafny.Map({d_1177_cf50_: (d_1179_v65_).f35})
                rhs200_ = (d_1180_v66_) | ((d_1180_v66_) | (d_1180_v66_))
                rhs201_ = p2
                lhs192_ = globalState
                d_1180_v66_ = rhs200_
                lhs192_.f21 = rhs201_
                d_1181_v67_: _dafny.Array
                nw206_ = _dafny.Array(int(0), 1)
                d_1181_v67_ = nw206_
                d_1182_v68_: _dafny.Map
                d_1182_v68_ = _dafny.Map({d_1181_v67_: True})
                d_1183_v69_: str
                d_1183_v69_ = _dafny.CodePoint('n')
                d_1184_v70_: D2
                d_1184_v70_ = D2_DC8(d_1183_v69_, d_1177_cf50_)
                rhs202_ = d_1182_v68_
                rhs203_ = (default__.fm35(_dafny.SeqWithoutIsStrInference([d_1184_v70_]), p2, globalState)) + ((_dafny.SeqWithoutIsStrInference([(d_1179_v65_).f35, d_1176_cf51_])) + (d_1077_v5_))
                rhs204_ = (((self.f30).set(default__.safeIndex(p2, len(self.f30)), d_1183_v69_)) + (self.f30)) + ((self.f30) + (self.f30))
                rhs205_ = d_1077_v5_
                rhs206_ = (d_1177_cf50_) or (p0)
                lhs193_ = self
                lhs194_ = globalState
                d_1182_v68_ = rhs202_
                d_1077_v5_ = rhs203_
                lhs193_.f30 = rhs204_
                d_1077_v5_ = rhs205_
                lhs194_.f13 = rhs206_
            elif source20_.is_DC32:
                d_1185___mcc_h19_ = source20_.cf52
                d_1186___mcc_h20_ = source20_.cf53
                d_1187_cf53_ = d_1186___mcc_h20_
                d_1188_cf52_ = d_1185___mcc_h19_
                d_1189_v71_: C0
                nw207_ = C0()
                nw207_.ctor__((self).f29)
                d_1189_v71_ = nw207_
                d_1190_v72_: _dafny.Array
                nw208_ = _dafny.Array(_dafny.Map({}), 26)
                d_1190_v72_ = nw208_
                d_1190_v72_ = d_1190_v72_
                (globalState).f19 = p2
                d_1191_v73_: D9
                d_1191_v73_ = D9_DC28(_dafny.SeqWithoutIsStrInference([p0, not((self).f29)]), d_1187_cf53_, p2, d_1187_cf53_, d_1189_v71_.f37)
                d_1192_v74_: _dafny.Map
                d_1192_v74_ = _dafny.Map({p2: default__.fm42(d_1189_v71_.f37, (d_1191_v73_).cf45, globalState)})
                d_1192_v74_ = (d_1192_v74_).set((p2) * (p2), d_1188_cf52_)
            elif source20_.is_DC33:
                d_1193___mcc_h21_ = source20_.cf54
                d_1194___mcc_h22_ = source20_.cf55
                d_1195_cf55_ = d_1194___mcc_h22_
                d_1196_cf54_ = d_1193___mcc_h21_
                d_1197_v75_: C1
                nw209_ = C1()
                nw209_.ctor__(p0)
                d_1197_v75_ = nw209_
                d_1198_v76_: C1
                nw210_ = C1()
                nw210_.ctor__((self).f29)
                d_1198_v76_ = nw210_
                d_1199_v77_: _dafny.Array
                nw211_ = _dafny.Array(int(0), 12)
                d_1199_v77_ = nw211_
                index191_ = default__.safeIndex(379, (d_1199_v77_).length(0))
                (d_1199_v77_)[index191_] = p2
                index192_ = default__.safeIndex(379, (d_1199_v77_).length(0))
                (d_1199_v77_)[index192_] = p2
                (globalState).f1 = (d_1199_v77_)[default__.safeIndex(379, (d_1199_v77_).length(0))]
            elif True:
                d_1200___mcc_h23_ = source20_.cf49
                d_1201_cf49_ = d_1200___mcc_h23_
                d_1202_v78_: _dafny.MultiSet
                d_1202_v78_ = _dafny.MultiSet([p0])
                d_1202_v78_ = d_1202_v78_
                d_1203_v79_: _dafny.Array
                def lambda57_(d_1204_p2_, d_1205_v5_):
                    def lambda58_(d_1206_i4_):
                        return (d_1206_i4_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1204_p2_]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1204_p2_, d_1204_p2_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(d_1205_v5_), d_1204_p2_, d_1204_p2_, 962]))).cardinality]))])))

                    return lambda58_

                init33_ = lambda57_(p2, d_1077_v5_)
                nw212_ = _dafny.Array(None, 14)
                for i0_33_ in range(nw212_.length(0)):
                    nw212_[i0_33_] = init33_(i0_33_)
                d_1203_v79_ = nw212_
                (globalState).f17 = d_1203_v79_
                (globalState).f15 = (self).f29
                d_1207_v80_: str
                d_1207_v80_ = _dafny.CodePoint('c')
                d_1207_v80_ = d_1207_v80_
            d_1208_v81_: _dafny.Array
            nw213_ = _dafny.Array(False, 18)
            d_1208_v81_ = nw213_
            d_1209_v82_: _dafny.Array
            nw214_ = _dafny.Array(None, 5)
            nw214_[int(0)] = d_1208_v81_
            nw214_[int(1)] = d_1208_v81_
            nw214_[int(2)] = d_1208_v81_
            nw214_[int(3)] = d_1208_v81_
            nw214_[int(4)] = d_1208_v81_
            d_1209_v82_ = nw214_
            index193_ = default__.safeIndex(117, (d_1209_v82_).length(0))
            (d_1209_v82_)[index193_] = d_1208_v81_
            d_1210_v83_: _dafny.Map
            d_1210_v83_ = _dafny.Map({d_1208_v81_: default__.fm10(p2, p2, globalState)})
            d_1211_v84_: _dafny.Seq
            d_1211_v84_ = _dafny.SeqWithoutIsStrInference([len(d_1210_v83_), (0) - (len(self.f30))])
            index194_ = default__.safeIndex(117, (d_1209_v82_).length(0))
            rhs207_ = len((d_1211_v84_) + ((d_1211_v84_) + (d_1211_v84_)))
            rhs208_ = p2
            rhs209_ = d_1208_v81_
            rhs210_ = (self).f29
            lhs195_ = globalState
            lhs196_ = globalState
            lhs197_ = d_1209_v82_
            lhs198_ = default__.safeIndex(117, (d_1209_v82_).length(0))
            lhs199_ = globalState
            lhs195_.f19 = rhs207_
            lhs196_.f21 = rhs208_
            lhs197_[lhs198_] = rhs209_
            lhs199_.f15 = rhs210_
            d_1212_v85_: _dafny.MultiSet
            d_1212_v85_ = _dafny.MultiSet([(self).f29, p0])
            d_1213_v86_: _dafny.Set
            d_1213_v86_ = _dafny.Set({p2, p2, len(_dafny.SeqWithoutIsStrInference([(d_1212_v85_).cardinality])), p2, p2})
            d_1214_v87_: D8
            d_1214_v87_ = D8_DC26()
            d_1215_v88_: _dafny.Map
            d_1215_v88_ = _dafny.Map({d_1213_v86_: d_1214_v87_})
            d_1215_v88_ = (d_1215_v88_).set(d_1213_v86_, d_1214_v87_)
            d_1216_v89_: _dafny.Array
            nw215_ = _dafny.Array(D10.default()(), 14)
            d_1216_v89_ = nw215_
            d_1217_v90_: _dafny.Map
            d_1217_v90_ = _dafny.Map({default__.fm34(p2, p2, p0, True, globalState): d_1216_v89_})
            d_1218_v91_: _dafny.Map
            d_1218_v91_ = _dafny.Map({p0: p0})
            rhs211_ = len(_dafny.SeqWithoutIsStrInference([((d_1217_v90_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "strdxawaa"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "strdxawaa"))) in (d_1217_v90_) else d_1216_v89_), d_1216_v89_, d_1216_v89_, d_1216_v89_, d_1216_v89_]))
            rhs212_ = (self).f29
            rhs213_ = not (False) or ((p2) == (p2))
            rhs214_ = (0) - (p2)
            rhs215_ = (p2) * (len(_dafny.SeqWithoutIsStrInference([len(self.f30), p2, len(d_1218_v91_)])))
            lhs200_ = globalState
            lhs201_ = globalState
            lhs202_ = globalState
            lhs203_ = globalState
            lhs204_ = globalState
            lhs200_.f2 = rhs211_
            lhs201_.f28 = rhs212_
            lhs202_.f15 = rhs213_
            lhs203_.f19 = rhs214_
            lhs204_.f2 = rhs215_
            d_1219_v92_: D3
            d_1219_v92_ = D3_DC11(p0, p0)
            d_1220_v93_: D3
            d_1220_v93_ = D3_DC12(d_1219_v92_)
            d_1221_v94_: _dafny.Set
            d_1221_v94_ = _dafny.Set({d_1220_v93_})
            d_1222_v95_: _dafny.Set
            d_1222_v95_ = _dafny.Set({d_1221_v94_, d_1221_v94_})
            (globalState).f18 = (d_1222_v95_).issubset((_dafny.Set({d_1221_v94_})) | (d_1222_v95_))
        d_1223_i5_: int
        d_1223_i5_ = 0
        with _dafny.label("6"):
            while (self).f29:
                with _dafny.c_label("6"):
                    if (d_1223_i5_) >= (100):
                        raise _dafny.Break("6")
                    d_1223_i5_ = (d_1223_i5_) + (1)
                    d_1224_v96_: _dafny.Map
                    d_1224_v96_ = _dafny.Map({(self).f29: p0})
                    d_1225_v97_: D11
                    d_1225_v97_ = D11_DC37(p2, p0)
                    d_1226_v98_: _dafny.Array
                    nw216_ = _dafny.Array(None, 14)
                    nw216_[int(0)] = p0
                    nw216_[int(1)] = (self).f29
                    nw216_[int(2)] = not(not(p0))
                    nw216_[int(3)] = p0
                    nw216_[int(4)] = p0
                    nw216_[int(5)] = p0
                    nw216_[int(6)] = True
                    nw216_[int(7)] = default__.fm33(p0, p2, p2, globalState)
                    nw216_[int(8)] = (self).f29
                    nw216_[int(9)] = (self).f29
                    nw216_[int(10)] = ((d_1224_v96_)[p0] if (p0) in (d_1224_v96_) else (d_1077_v5_)[default__.safeIndex(p2, len(d_1077_v5_))])
                    nw216_[int(11)] = (d_1225_v97_).cf63
                    nw216_[int(12)] = (self).f29
                    nw216_[int(13)] = p0
                    d_1226_v98_ = nw216_
                    d_1227_v99_: _dafny.Map
                    d_1227_v99_ = _dafny.Map({p2: d_1226_v98_})
                    d_1227_v99_ = (d_1227_v99_).set(p2, d_1226_v98_)
                    d_1228_v100_: C4
                    nw217_ = C4()
                    nw217_.ctor__(True, (-450) * (p2))
                    d_1228_v100_ = nw217_
                    (globalState).f15 = not ((d_1228_v100_).f33) or ((self).f29)
                    d_1229_v101_: str
                    d_1229_v101_ = _dafny.CodePoint('y')
                    d_1230_v102_: D5
                    d_1230_v102_ = D5_DC19(453, (_dafny.MultiSet(d_1077_v5_)).cardinality, d_1229_v101_)
                    d_1229_v101_ = (d_1230_v102_).cf30
                    pass
            pass
        d_1231_v103_: D9
        d_1231_v103_ = D9_DC29(p2, p2, (self).f29)
        (globalState).f21 = (d_1231_v103_).cf46
        if p0:
            d_1232_v104_: _dafny.MultiSet
            d_1232_v104_ = _dafny.MultiSet([default__.safeDivisionInt(p2, p2)])
            d_1233_v105_: _dafny.Map
            d_1233_v105_ = _dafny.Map({p2: d_1232_v104_})
            d_1234_v106_: _dafny.Map
            d_1234_v106_ = _dafny.Map({p2: p2})
            d_1235_v107_: _dafny.Map
            d_1235_v107_ = _dafny.Map({p0: len(d_1234_v106_)})
            d_1232_v104_ = ((d_1233_v105_)[len(_dafny.Map({d_1235_v107_: not((self).f29)}))] if (len(_dafny.Map({d_1235_v107_: not((self).f29)}))) in (d_1233_v105_) else (d_1232_v104_).set(p2, default__.abs(len(self.f30))))
            d_1236_v108_: str
            d_1236_v108_ = _dafny.CodePoint('e')
            d_1236_v108_ = d_1236_v108_
            d_1237_v109_: _dafny.Array
            nw218_ = _dafny.Array(_dafny.Map({}), 25)
            d_1237_v109_ = nw218_
            d_1238_v110_: _dafny.Map
            d_1238_v110_ = _dafny.Map({p2: (self).f29})
            index195_ = default__.safeIndex(975, (d_1237_v109_).length(0))
            (d_1237_v109_)[index195_] = d_1238_v110_
            index196_ = default__.safeIndex(975, (d_1237_v109_).length(0))
            (d_1237_v109_)[index196_] = d_1238_v110_
            (globalState).f20 = ((self.f30) + ((self.f30).set(default__.safeIndex(p2, len(self.f30)), d_1236_v108_))) + ((self.f30) + (self.f30))
            d_1239_v111_: _dafny.Seq
            d_1239_v111_ = _dafny.SeqWithoutIsStrInference([self.f30])
            d_1240_v112_: _dafny.MultiSet
            d_1240_v112_ = _dafny.MultiSet([self.f30, (d_1239_v111_)[default__.safeIndex(-182, len(d_1239_v111_))]])
            d_1241_v113_: _dafny.Seq
            d_1241_v113_ = _dafny.SeqWithoutIsStrInference([d_1240_v112_])
            (globalState).f18 = ((d_1241_v113_)[default__.safeIndex(p2, len(d_1241_v113_))]).ispropersubset((d_1240_v112_) - (d_1240_v112_))
        elif True:
            d_1242_v114_: _dafny.Array
            def lambda59_(d_1243_p2_):
                def lambda60_(d_1244_i6_):
                    return (d_1244_i6_) * (d_1243_p2_)

                return lambda60_

            init34_ = lambda59_(p2)
            nw219_ = _dafny.Array(None, 28)
            for i0_34_ in range(nw219_.length(0)):
                nw219_[i0_34_] = init34_(i0_34_)
            d_1242_v114_ = nw219_
            d_1245_v115_: _dafny.Set
            d_1245_v115_ = _dafny.Set({d_1242_v114_})
            d_1246_v116_: C2
            nw220_ = C2()
            nw220_.ctor__(p2, d_1245_v115_, (self).f29, self.f30)
            d_1246_v116_ = nw220_
            d_1247_v117_: _dafny.Array
            nw221_ = _dafny.Array(False, 13)
            d_1247_v117_ = nw221_
            d_1248_v118_: _dafny.Seq
            d_1248_v118_ = _dafny.SeqWithoutIsStrInference([d_1247_v117_])
            d_1249_v119_: _dafny.Map
            d_1249_v119_ = _dafny.Map({d_1246_v116_: d_1248_v118_})
            d_1250_v120_: _dafny.Map
            d_1250_v120_ = _dafny.Map({((d_1249_v119_)[d_1246_v116_] if (d_1246_v116_) in (d_1249_v119_) else d_1248_v118_): (d_1246_v116_).f38})
            d_1250_v120_ = (d_1250_v120_).set(d_1248_v118_, 80)
            d_1251_v121_: str
            d_1251_v121_ = _dafny.CodePoint('v')
            d_1252_v122_: D1
            d_1252_v122_ = D1_DC4((self).f29)
            d_1253_v123_: _dafny.Map
            d_1253_v123_ = _dafny.Map({len(self.f30): D11_DC35((self.f30).set(default__.safeIndex((d_1246_v116_).f38, len(self.f30)), d_1251_v121_), (d_1252_v122_).cf8)})
            d_1254_v124_: _dafny.Set
            d_1254_v124_ = _dafny.Set({d_1253_v123_})
            d_1255_v125_: _dafny.MultiSet
            d_1255_v125_ = _dafny.MultiSet([p0])
            d_1256_v126_: _dafny.Set
            d_1256_v126_ = _dafny.Set({(_dafny.MultiSet([(self).f29])) == (d_1255_v125_)})
            d_1257_v127_: D11
            d_1257_v127_ = D11_DC37(len(_dafny.SeqWithoutIsStrInference([d_1251_v121_ for d_1258_i7_ in range(default__.abs(948))])), (self).f29)
            d_1259_v128_: _dafny.Seq
            d_1259_v128_ = _dafny.SeqWithoutIsStrInference([(d_1246_v116_).f38, (d_1246_v116_).f38])
            d_1260_v129_: _dafny.Map
            d_1260_v129_ = _dafny.Map({d_1259_v128_: self.f30})
            rhs216_ = True
            rhs217_ = (d_1257_v127_).cf62
            rhs218_ = default__.fm47(p0, (self).f29, globalState)
            rhs219_ = d_1256_v126_
            rhs220_ = (d_1251_v121_) in ((self.f30 if (self).f29 else ((d_1260_v129_)[d_1259_v128_] if (d_1259_v128_) in (d_1260_v129_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ysikfqdgi")))))
            lhs205_ = globalState
            lhs206_ = globalState
            lhs207_ = globalState
            lhs205_.f18 = rhs216_
            lhs206_.f2 = rhs217_
            d_1254_v124_ = rhs218_
            d_1256_v126_ = rhs219_
            lhs207_.f18 = rhs220_
            d_1261_v130_: _dafny.Array
            nw222_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_1261_v130_ = nw222_
            d_1261_v130_ = d_1261_v130_
            d_1262_v131_: C0
            nw223_ = C0()
            nw223_.ctor__(p0)
            d_1262_v131_ = nw223_
            d_1263_v132_: _dafny.Seq
            d_1263_v132_ = _dafny.SeqWithoutIsStrInference([d_1262_v131_])
            (globalState).f21 = ((d_1246_v116_).f38) + ((len(d_1077_v5_)) - (len(d_1263_v132_)))
            d_1256_v126_ = d_1256_v126_
        r0 = self.f30
        d_1264_v133_: _dafny.Map
        d_1264_v133_ = _dafny.Map({p0: d_1077_v5_})
        r1 = ((not((self).f29)) in (((d_1264_v133_)[False] if (False) in (d_1264_v133_) else d_1077_v5_)) if (self).f29 else p0)
        return r0, r1


class C6(T0, T1):
    def  __init__(self):
        self._f30: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f29: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f30(self):
        return self._f30
    @f30.setter
    def f30(self, value):
        self._f30 = value
    @property
    def f29(self):
        return self._f29
    def ctor__(self, f29, f30):
        (self)._f29 = f29
        (self).f30 = f30

    def fm2(self, p0, globalState):
        return (self).f29

    def m2(self, p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_1265_v0_: int
        d_1265_v0_ = 914
        d_1266_v1_: _dafny.MultiSet
        d_1266_v1_ = _dafny.MultiSet([d_1265_v0_, d_1265_v0_, d_1265_v0_, d_1265_v0_])
        hi7_ = d_1265_v0_
        for d_1267_i0_ in range((d_1266_v1_).cardinality, hi7_):
            (globalState).f19 = (d_1267_i0_) + ((0) - (d_1267_i0_))
            d_1268_v2_: _dafny.Array
            nw224_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_1268_v2_ = nw224_
            nw225_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_1268_v2_ = nw225_
            if not(not ((self).f29) or ((self).f29)):
                d_1269_v3_: _dafny.Array
                nw226_ = _dafny.Array(False, 12)
                d_1269_v3_ = nw226_
                d_1270_v4_: D0
                d_1270_v4_ = D0_DC1(d_1269_v3_, d_1265_v0_, d_1267_i0_)
                (globalState).f1 = (d_1270_v4_).cf3
                d_1271_v5_: _dafny.MultiSet
                d_1271_v5_ = _dafny.MultiSet([p0])
                d_1272_v6_: D0
                d_1272_v6_ = D0_DC2(d_1271_v5_, self.f30, p0)
                d_1273_v7_: _dafny.Map
                d_1273_v7_ = _dafny.Map({d_1272_v6_: d_1265_v0_})
                d_1273_v7_ = d_1273_v7_
                d_1274_v8_: C1
                nw227_ = C1()
                nw227_.ctor__(p0)
                d_1274_v8_ = nw227_
                d_1275_v9_: _dafny.Map
                d_1275_v9_ = _dafny.Map({len((self.f30) + (self.f30)): (d_1265_v0_) >= (d_1265_v0_)})
                d_1276_v10_: _dafny.Seq
                d_1276_v10_ = _dafny.SeqWithoutIsStrInference([(d_1274_v8_).f40])
                d_1275_v9_ = (d_1275_v9_).set((d_1267_i0_) * (len(d_1276_v10_)), (p0) == (p0))
                rhs221_ = default__.fm36(globalState)
                rhs222_ = p0
                lhs208_ = globalState
                lhs209_ = globalState
                lhs208_.f21 = rhs221_
                lhs209_.f13 = rhs222_
            elif True:
                d_1277_v11_: _dafny.Seq
                d_1277_v11_ = _dafny.SeqWithoutIsStrInference([p0])
                (globalState).f21 = (len(d_1277_v11_)) * (d_1265_v0_)
                r1 = not(True)
                d_1278_v12_: _dafny.Array
                nw228_ = _dafny.Array(False, 10)
                d_1278_v12_ = nw228_
                index197_ = default__.safeIndex(730, (d_1278_v12_).length(0))
                (d_1278_v12_)[index197_] = not(p0)
                index198_ = default__.safeIndex(730, (d_1278_v12_).length(0))
                (d_1278_v12_)[index198_] = not((self).f29)
                (globalState).f18 = (self.f30) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1279_i1_ in range(default__.abs(207))]))
                d_1280_v13_: C5
                nw229_ = C5()
                nw229_.ctor__(True, self.f30)
                d_1280_v13_ = nw229_
            (globalState).f14 = d_1265_v0_
        d_1281_v14_: D6
        d_1281_v14_ = D6_DC23((self).f29, d_1265_v0_)
        d_1282_v15_: _dafny.Set
        d_1282_v15_ = _dafny.Set({(self).f29})
        d_1283_v16_: D11
        d_1283_v16_ = D11_DC36(d_1281_v14_, d_1282_v15_, (self).f29)
        d_1284_v17_: _dafny.Array
        nw230_ = _dafny.Array(None, 1)
        nw230_[int(0)] = (d_1283_v16_).cf61
        d_1284_v17_ = nw230_
        d_1285_v18_: D0
        d_1285_v18_ = D0_DC1(d_1284_v17_, d_1265_v0_, d_1265_v0_)
        d_1286_v19_: _dafny.Array
        nw231_ = _dafny.Array(None, 11)
        nw231_[int(0)] = (d_1284_v17_ if (self).fm2(p0, globalState) else d_1284_v17_)
        nw231_[int(1)] = d_1284_v17_
        nw231_[int(2)] = d_1284_v17_
        nw231_[int(3)] = (d_1284_v17_ if (self).f29 else d_1284_v17_)
        nw231_[int(4)] = d_1284_v17_
        nw231_[int(5)] = d_1284_v17_
        nw231_[int(6)] = d_1284_v17_
        nw231_[int(7)] = d_1284_v17_
        nw231_[int(8)] = (d_1285_v18_).cf1
        nw231_[int(9)] = d_1284_v17_
        nw231_[int(10)] = d_1284_v17_
        d_1286_v19_ = nw231_
        index199_ = default__.safeIndex(249, (d_1286_v19_).length(0))
        (d_1286_v19_)[index199_] = d_1284_v17_
        index200_ = default__.safeIndex(249, (d_1286_v19_).length(0))
        (d_1286_v19_)[index200_] = d_1284_v17_
        d_1287_v20_: _dafny.Set
        d_1287_v20_ = _dafny.Set({d_1265_v0_, d_1265_v0_})
        d_1288_v21_: _dafny.Seq
        d_1288_v21_ = _dafny.SeqWithoutIsStrInference([p0, (d_1287_v20_).issubset(d_1287_v20_), (self).f29])
        (globalState).f13 = (d_1288_v21_)[default__.safeIndex(d_1265_v0_, len(d_1288_v21_))]
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_1284_v17_).length(0)):
            d_1289_i2_: int = guard_loop_8_
            if (True) and (((0) <= (d_1289_i2_)) and ((d_1289_i2_) < ((d_1284_v17_).length(0)))):
                (d_1284_v17_)[(d_1289_i2_)] = p0
        d_1290_v22_: _dafny.Map
        d_1290_v22_ = _dafny.Map({not (p0) or (False): d_1265_v0_})
        d_1291_v23_: _dafny.Seq
        d_1291_v23_ = _dafny.SeqWithoutIsStrInference([d_1265_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kyov")))])
        d_1290_v22_ = (d_1290_v22_) | (_dafny.Map({p0: len(d_1291_v23_)}))
        r1 = (self).f29
        d_1292_v24_: _dafny.Map
        d_1292_v24_ = _dafny.Map({99: (_dafny.SeqWithoutIsStrInference([d_1288_v21_ for d_1293_i3_ in range(default__.abs(112))])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p0, (self).f29])), len(_dafny.SeqWithoutIsStrInference([d_1288_v21_ for d_1294_i3_ in range(default__.abs(112))]))), _dafny.SeqWithoutIsStrInference([p0]))})
        d_1295_v25_: _dafny.Seq
        d_1295_v25_ = _dafny.SeqWithoutIsStrInference([default__.fm35(_dafny.SeqWithoutIsStrInference([D2_DC8(_dafny.CodePoint('d'), (self).f29) for d_1296_i4_ in range(default__.abs(948))]), d_1265_v0_, globalState), d_1288_v21_])
        r0 = (((d_1292_v24_)[d_1265_v0_] if (d_1265_v0_) in (d_1292_v24_) else d_1295_v25_)) + (_dafny.SeqWithoutIsStrInference([d_1288_v21_, d_1288_v21_]))
        r1 = p0
        return r0, r1

    def m3(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        d_1297_v0_: _dafny.Map
        d_1297_v0_ = _dafny.Map({(self).f29: not((self).f29)})
        (globalState).f15 = ((d_1297_v0_)[(self).f29] if ((self).f29) in (d_1297_v0_) else (True if p0 else (self).f29))
        d_1298_v1_: _dafny.MultiSet
        d_1298_v1_ = _dafny.MultiSet([d_1297_v0_, d_1297_v0_, _dafny.Map({(self).f29: (self).f29})])
        d_1299_v2_: _dafny.Seq
        d_1299_v2_ = _dafny.SeqWithoutIsStrInference([(d_1298_v1_).ispropersubset(d_1298_v1_), (self).f29])
        d_1300_v3_: D9
        d_1300_v3_ = D9_DC28(d_1299_v2_, p0, (0) - (p2), not(True), (self).f29)
        d_1301_v4_: _dafny.Seq
        d_1301_v4_ = _dafny.SeqWithoutIsStrInference([(d_1299_v2_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgmvdm"))), len(d_1299_v2_)), p0), _dafny.SeqWithoutIsStrInference([p0])])
        d_1299_v2_ = ((d_1299_v2_) + (((d_1300_v3_).cf41).set(default__.safeIndex(p2, len((d_1300_v3_).cf41)), p0))) + ((d_1301_v4_)[default__.safeIndex(p2, len(d_1301_v4_))])
        hi8_ = p2
        for d_1302_i0_ in range((p2) + (p2), hi8_):
            d_1299_v2_ = d_1299_v2_
            d_1303_v5_: _dafny.Array
            nw232_ = _dafny.Array(False, 7)
            d_1303_v5_ = nw232_
            index201_ = default__.safeIndex(971, (d_1303_v5_).length(0))
            (d_1303_v5_)[index201_] = True
            index202_ = default__.safeIndex(665, (d_1303_v5_).length(0))
            (d_1303_v5_)[index202_] = not ((self).f29) or (p0)
            d_1304_v6_: _dafny.Map
            d_1304_v6_ = _dafny.Map({(0) - (p2): (self).f29})
            index203_ = default__.safeIndex(971, (d_1303_v5_).length(0))
            index204_ = default__.safeIndex(665, (d_1303_v5_).length(0))
            rhs223_ = p0
            rhs224_ = (p0 if (40) != (len(d_1304_v6_)) else default__.fm6(p2, globalState))
            rhs225_ = p0
            rhs226_ = default__.fm36(globalState)
            rhs227_ = p0
            lhs210_ = globalState
            lhs211_ = globalState
            lhs212_ = d_1303_v5_
            lhs213_ = default__.safeIndex(971, (d_1303_v5_).length(0))
            lhs214_ = globalState
            lhs215_ = d_1303_v5_
            lhs216_ = default__.safeIndex(665, (d_1303_v5_).length(0))
            lhs210_.f15 = rhs223_
            lhs211_.f18 = rhs224_
            lhs212_[lhs213_] = rhs225_
            lhs214_.f2 = rhs226_
            lhs215_[lhs216_] = rhs227_
            d_1305_v7_: _dafny.Seq
            d_1305_v7_ = _dafny.SeqWithoutIsStrInference([self.f30, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1306_i1_ in range(default__.abs(814))])])
            d_1307_v8_: _dafny.Seq
            d_1307_v8_ = _dafny.SeqWithoutIsStrInference([d_1302_i0_, (0) - (p2), d_1302_i0_, p2, d_1302_i0_])
            d_1308_v9_: bool
            d_1309_v10_: _dafny.Seq
            out18_: bool
            out19_: _dafny.Seq
            out18_, out19_ = (self).m5((d_1305_v7_)[default__.safeIndex(d_1302_i0_, len(d_1305_v7_))], ((d_1303_v5_)[default__.safeIndex(971, (d_1303_v5_).length(0))] if not(p0) else p0), (len(d_1307_v8_)) - (default__.fm36(globalState)), globalState)
            d_1308_v9_ = out18_
            d_1309_v10_ = out19_
            d_1310_v11_: C0
            nw233_ = C0()
            nw233_.ctor__(False)
            d_1310_v11_ = nw233_
            d_1311_v12_: _dafny.Seq
            d_1311_v12_ = _dafny.SeqWithoutIsStrInference([d_1310_v11_, d_1310_v11_, d_1310_v11_])
            d_1312_v13_: D9
            d_1312_v13_ = D9_DC29(len(d_1311_v12_), len(_dafny.Map({p2: d_1302_i0_})), (self).f29)
            d_1297_v0_ = (d_1297_v0_).set((d_1312_v13_).cf48, (d_1310_v11_).fm11(d_1308_v9_, globalState))
        if (self).f29:
            d_1313_v14_: _dafny.Array
            nw234_ = _dafny.Array(False, 17)
            d_1313_v14_ = nw234_
            rhs228_ = (0) - (p2)
            rhs229_ = d_1313_v14_
            rhs230_ = d_1313_v14_
            lhs217_ = globalState
            lhs217_.f2 = rhs228_
            d_1313_v14_ = rhs229_
            d_1313_v14_ = rhs230_
            d_1314_v15_: C1
            nw235_ = C1()
            nw235_.ctor__((len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1315_i2_ in range(default__.abs(-156))]))) <= (p2))
            d_1314_v15_ = nw235_
            d_1316_v16_: _dafny.MultiSet
            d_1316_v16_ = _dafny.MultiSet([p2, p2])
            d_1317_v17_: _dafny.Map
            d_1317_v17_ = _dafny.Map({495: d_1316_v16_})
            d_1318_v18_: _dafny.Map
            d_1318_v18_ = _dafny.Map({p2: p0})
            d_1319_v19_: _dafny.Seq
            d_1319_v19_ = _dafny.SeqWithoutIsStrInference([p2, p2, (((d_1317_v17_)[len(d_1318_v18_)] if (len(d_1318_v18_)) in (d_1317_v17_) else _dafny.MultiSet([default__.fm36(globalState)]))).cardinality])
            if ((d_1319_v19_) + (d_1319_v19_)) <= (d_1319_v19_):
                d_1320_v20_: _dafny.Array
                nw236_ = _dafny.Array(int(0), 4)
                d_1320_v20_ = nw236_
                index205_ = default__.safeIndex(138, (d_1320_v20_).length(0))
                (d_1320_v20_)[index205_] = p2
                index206_ = default__.safeIndex(138, (d_1320_v20_).length(0))
                (d_1320_v20_)[index206_] = p2
                d_1321_v21_: _dafny.Set
                d_1321_v21_ = _dafny.Set({d_1320_v20_, d_1320_v20_})
                d_1322_v22_: C2
                nw237_ = C2()
                nw237_.ctor__((d_1320_v20_)[default__.safeIndex(138, (d_1320_v20_).length(0))], d_1321_v21_, (d_1314_v15_).f40, self.f30)
                d_1322_v22_ = nw237_
                d_1323_v23_: _dafny.Map
                d_1323_v23_ = _dafny.Map({d_1322_v22_: d_1297_v0_})
                d_1324_v24_: _dafny.Map
                d_1324_v24_ = d_1297_v0_
                d_1325_v25_: _dafny.Array
                nw238_ = _dafny.Array(None, 14)
                nw238_[int(0)] = d_1297_v0_
                nw238_[int(1)] = d_1297_v0_
                nw238_[int(2)] = d_1297_v0_
                nw238_[int(3)] = d_1297_v0_
                nw238_[int(4)] = _dafny.Map({(d_1314_v15_).f40: (d_1314_v15_).f40})
                nw238_[int(5)] = ((d_1323_v23_)[d_1322_v22_] if (d_1322_v22_) in (d_1323_v23_) else d_1297_v0_)
                nw238_[int(6)] = (d_1297_v0_) | (d_1297_v0_)
                nw238_[int(7)] = ((d_1297_v0_).set((self).f29, (self).f29)) | (d_1297_v0_)
                nw238_[int(8)] = d_1297_v0_
                nw238_[int(9)] = _dafny.Map({(d_1314_v15_).f40: (self).f29})
                nw238_[int(10)] = (d_1324_v24_)
                nw238_[int(11)] = d_1297_v0_
                nw238_[int(12)] = d_1297_v0_
                nw238_[int(13)] = d_1297_v0_
                d_1325_v25_ = nw238_
                d_1326_v26_: str
                d_1326_v26_ = _dafny.CodePoint('k')
                d_1327_v27_: D2
                d_1327_v27_ = D2_DC8(d_1326_v26_, (self).f29)
                d_1328_v28_: D5
                d_1328_v28_ = D5_DC19((d_1320_v20_)[default__.safeIndex(138, (d_1320_v20_).length(0))], len(d_1299_v2_), (d_1327_v27_).cf12)
                rhs231_ = d_1325_v25_
                rhs232_ = (d_1319_v19_).set(default__.safeIndex((d_1328_v28_).cf28, len(d_1319_v19_)), default__.safeDivisionInt((d_1322_v22_).f38, (d_1322_v22_).f38))
                rhs233_ = p2
                rhs234_ = ((self.f30) == (self.f30)) in (d_1299_v2_)
                rhs235_ = ((d_1322_v22_).f38) * ((d_1322_v22_).f38)
                lhs218_ = globalState
                lhs219_ = globalState
                lhs220_ = globalState
                lhs221_ = globalState
                d_1325_v25_ = rhs231_
                lhs218_.f7 = rhs232_
                lhs219_.f21 = rhs233_
                lhs220_.f13 = rhs234_
                lhs221_.f1 = rhs235_
                d_1329_v29_: C0
                nw239_ = C0()
                nw239_.ctor__((self).f29)
                d_1329_v29_ = nw239_
                d_1330_v30_: _dafny.Map
                d_1330_v30_ = _dafny.Map({d_1329_v29_.f37: 994})
                (globalState).f21 = default__.safeModuloInt(len(d_1330_v30_), ((d_1322_v22_).f38) - ((d_1329_v29_).fm12(globalState)))
                d_1331_v31_: _dafny.Set
                d_1331_v31_ = _dafny.Set({(-182) + (p2), (p2) - (len(self.f30)), p2, (d_1320_v20_)[default__.safeIndex(138, (d_1320_v20_).length(0))]})
                d_1331_v31_ = d_1331_v31_
            elif True:
                (globalState).f1 = p2
                d_1332_v32_: _dafny.Array
                nw240_ = _dafny.Array(_dafny.Seq({}), 10)
                d_1332_v32_ = nw240_
                d_1333_v33_: _dafny.Array
                nw241_ = _dafny.Array(None, 9)
                nw241_[int(0)] = d_1332_v32_
                nw241_[int(1)] = d_1332_v32_
                nw241_[int(2)] = d_1332_v32_
                nw241_[int(3)] = d_1332_v32_
                nw241_[int(4)] = d_1332_v32_
                nw241_[int(5)] = d_1332_v32_
                nw241_[int(6)] = d_1332_v32_
                nw241_[int(7)] = d_1332_v32_
                nw241_[int(8)] = d_1332_v32_
                d_1333_v33_ = nw241_
                index207_ = default__.safeIndex(400, (d_1333_v33_).length(0))
                (d_1333_v33_)[index207_] = d_1332_v32_
                d_1334_v34_: _dafny.Set
                d_1334_v34_ = _dafny.Set({p2, p2, p2, len(_dafny.SeqWithoutIsStrInference([p0])), len(d_1299_v2_)})
                index208_ = default__.safeIndex(400, (d_1333_v33_).length(0))
                rhs236_ = d_1332_v32_
                rhs237_ = default__.safeModuloInt(len(self.f30), p2)
                rhs238_ = (len(d_1334_v34_)) * ((p2 if not(p0) else p2))
                rhs239_ = (self).f29
                lhs222_ = d_1333_v33_
                lhs223_ = default__.safeIndex(400, (d_1333_v33_).length(0))
                lhs224_ = globalState
                lhs225_ = globalState
                lhs226_ = globalState
                lhs222_[lhs223_] = rhs236_
                lhs224_.f21 = rhs237_
                lhs225_.f21 = rhs238_
                lhs226_.f28 = rhs239_
                d_1318_v18_ = (d_1318_v18_).set(len(d_1299_v2_), (self).f29)
                (self).f30 = self.f30
                d_1335_v35_: str
                d_1335_v35_ = _dafny.CodePoint('e')
                rhs240_ = d_1313_v14_
                rhs241_ = d_1335_v35_
                rhs242_ = d_1318_v18_
                d_1313_v14_ = rhs240_
                d_1335_v35_ = rhs241_
                d_1318_v18_ = rhs242_
            d_1336_v36_: _dafny.Array
            nw242_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_1336_v36_ = nw242_
            index209_ = default__.safeIndex(196, (d_1336_v36_).length(0))
            (d_1336_v36_)[index209_] = d_1313_v14_
            index210_ = default__.safeIndex(196, (d_1336_v36_).length(0))
            (d_1336_v36_)[index210_] = d_1313_v14_
            d_1337_v37_: D3
            d_1337_v37_ = D3_DC10(d_1316_v16_)
            if (d_1316_v16_).issubset(((d_1337_v37_).cf18) | (_dafny.MultiSet([p2]))):
                d_1338_v38_: D12
                d_1338_v38_ = D12_DC39(len(d_1297_v0_), (d_1314_v15_).f40, p2)
                d_1339_v39_: D6
                d_1339_v39_ = D6_DC23(p0, p2)
                pat_let_tv28_ = d_1314_v15_
                def iife109_(_pat_let31_0):
                    def iife110_(d_1340_dt__update__tmp_h0_):
                        def iife111_(_pat_let32_0):
                            def iife112_(d_1341_dt__update_hcf36_h0_):
                                return D6_DC23(d_1341_dt__update_hcf36_h0_, (d_1340_dt__update__tmp_h0_).cf37)
                            return iife112_(_pat_let32_0)
                        return iife111_((pat_let_tv28_).f40)
                    return iife110_(_pat_let31_0)
                rhs243_ = (iife109_(d_1339_v39_)).cf37
                rhs244_ = not(not(p0))
                rhs245_ = d_1338_v38_
                rhs246_ = p2
                rhs247_ = (0) - (default__.safeDivisionInt((p2) * (-620), -108))
                lhs227_ = globalState
                lhs228_ = globalState
                lhs229_ = globalState
                lhs227_.f19 = rhs243_
                r1 = rhs244_
                d_1338_v38_ = rhs245_
                lhs228_.f19 = rhs246_
                lhs229_.f14 = rhs247_
                d_1342_v40_: _dafny.Array
                def lambda61_(d_1343_v15_, d_1344_p2_):
                    def lambda62_(d_1345_i3_):
                        return (D11_DC37(d_1344_p2_, (d_1343_v15_).f40) if (d_1343_v15_).f40 else D11_DC37(len(self.f30), False))

                    return lambda62_

                init35_ = lambda61_(d_1314_v15_, p2)
                nw243_ = _dafny.Array(None, 25)
                for i0_35_ in range(nw243_.length(0)):
                    nw243_[i0_35_] = init35_(i0_35_)
                d_1342_v40_ = nw243_
                d_1346_v41_: D11
                d_1346_v41_ = D11_DC37(p2, (self).f29)
                index211_ = default__.safeIndex(132, (d_1342_v40_).length(0))
                (d_1342_v40_)[index211_] = d_1346_v41_
                arr0_ = (d_1336_v36_)[default__.safeIndex(196, (d_1336_v36_).length(0))]
                index212_ = default__.safeIndex(590, ((d_1336_v36_)[default__.safeIndex(196, (d_1336_v36_).length(0))]).length(0))
                arr0_[index212_] = (d_1314_v15_).f40
                index213_ = default__.safeIndex(132, (d_1342_v40_).length(0))
                arr1_ = (d_1336_v36_)[default__.safeIndex(196, (d_1336_v36_).length(0))]
                index214_ = default__.safeIndex(590, ((d_1336_v36_)[default__.safeIndex(196, (d_1336_v36_).length(0))]).length(0))
                rhs248_ = d_1346_v41_
                rhs249_ = ((d_1319_v19_) + (_dafny.SeqWithoutIsStrInference([len(d_1297_v0_)]))) <= (d_1319_v19_)
                rhs250_ = (_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([(d_1314_v15_).f40, p0]))
                lhs230_ = d_1342_v40_
                lhs231_ = default__.safeIndex(132, (d_1342_v40_).length(0))
                lhs232_ = (d_1336_v36_)[default__.safeIndex(196, (d_1336_v36_).length(0))]
                lhs233_ = default__.safeIndex(590, ((d_1336_v36_)[default__.safeIndex(196, (d_1336_v36_).length(0))]).length(0))
                lhs230_[lhs231_] = rhs248_
                lhs232_[lhs233_] = rhs249_
                d_1299_v2_ = rhs250_
                (globalState).f28 = (d_1314_v15_).f40
                d_1347_v42_: str
                d_1347_v42_ = _dafny.CodePoint('f')
                d_1348_v44_: _dafny.Set
                d_1348_v44_ = _dafny.Set({p2})
                d_1349_v45_: _dafny.Seq
                def iife113_():
                    coll47_ = _dafny.Set()
                    compr_47_: int
                    for compr_47_ in _dafny.IntegerRange(644, -791):
                        d_1350_v43_: int = compr_47_
                        if ((644) <= (d_1350_v43_)) and ((d_1350_v43_) < (-791)):
                            coll47_ = coll47_.union(_dafny.Set([(d_1350_v43_) + (p2)]))
                    return _dafny.Set(coll47_)
                d_1349_v45_ = _dafny.SeqWithoutIsStrInference([iife113_()
                , d_1348_v44_, d_1348_v44_, d_1348_v44_])
                d_1347_v42_ = default__.fm42((default__.fm39(d_1347_v42_, 883, globalState)) == (default__.fm39(d_1347_v42_, (0) - (p2), globalState)), (d_1349_v45_) <= (d_1349_v45_), globalState)
                d_1351_v46_: _dafny.Array
                def lambda63_(d_1352_i4_):
                    return (d_1352_i4_) - (142)

                init36_ = lambda63_
                nw244_ = _dafny.Array(None, 22)
                for i0_36_ in range(nw244_.length(0)):
                    nw244_[i0_36_] = init36_(i0_36_)
                d_1351_v46_ = nw244_
                d_1353_v47_: D12
                d_1353_v47_ = D12_DC38(d_1351_v46_)
                (globalState).f17 = (d_1353_v47_).cf64
            elif True:
                (globalState).f20 = self.f30
                (globalState).f15 = ((p0) and ((d_1314_v15_).f40)) or ((self).f29)
                d_1354_v48_: _dafny.Map
                d_1354_v48_ = _dafny.Map({not(((0) - ((0) - (p2))) < (p2)): (p2) + (p2)})
                d_1355_v49_: _dafny.Map
                d_1355_v49_ = _dafny.Map({p2: p2})
                d_1354_v48_ = (d_1354_v48_).set(not(p0), ((d_1355_v49_)[default__.fm9(globalState)] if (default__.fm9(globalState)) in (d_1355_v49_) else p2))
                d_1356_v50_: _dafny.Array
                nw245_ = _dafny.Array(int(0), 19)
                d_1356_v50_ = nw245_
                d_1357_v51_: _dafny.Set
                d_1357_v51_ = _dafny.Set({d_1356_v50_, d_1356_v50_})
                d_1358_v52_: C3
                nw246_ = C3()
                nw246_.ctor__((d_1318_v18_) != (d_1318_v18_), (p2) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgqvib")))), (d_1357_v51_).ispropersubset(d_1357_v51_), self.f30)
                d_1358_v52_ = nw246_
                d_1358_v52_ = d_1358_v52_
                d_1356_v50_ = d_1356_v50_
        elif True:
            d_1359_v53_: _dafny.MultiSet
            d_1359_v53_ = _dafny.MultiSet([D11_DC37(p2, p0), D11_DC37(p2, p0)])
            d_1360_v54_: D11
            d_1360_v54_ = D11_DC37(p2, p0)
            d_1359_v53_ = (d_1359_v53_) - (_dafny.MultiSet([d_1360_v54_]))
            d_1361_v55_: C3
            nw247_ = C3()
            nw247_.ctor__((p0 if p0 else not(p0)), p2, (d_1299_v2_) <= (d_1299_v2_), self.f30)
            d_1361_v55_ = nw247_
            (globalState).f19 = (d_1361_v55_).f36
            d_1362_v56_: _dafny.Array
            def lambda64_(d_1363_p2_):
                def lambda65_(d_1364_i5_):
                    return (d_1364_i5_) + (d_1363_p2_)

                return lambda65_

            init37_ = lambda64_(p2)
            nw248_ = _dafny.Array(None, 12)
            for i0_37_ in range(nw248_.length(0)):
                nw248_[i0_37_] = init37_(i0_37_)
            d_1362_v56_ = nw248_
            index215_ = default__.safeIndex(728, (d_1362_v56_).length(0))
            (d_1362_v56_)[index215_] = p2
            d_1365_v57_: _dafny.Array
            nw249_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_1365_v57_ = nw249_
            d_1366_v58_: _dafny.Array
            nw250_ = _dafny.Array(False, 7)
            d_1366_v58_ = nw250_
            index216_ = default__.safeIndex(708, (d_1365_v57_).length(0))
            (d_1365_v57_)[index216_] = d_1366_v58_
            d_1367_v59_: _dafny.Map
            d_1367_v59_ = _dafny.Map({p2: (d_1361_v55_).f36})
            d_1368_v60_: D0
            d_1368_v60_ = D0_DC1(d_1366_v58_, (d_1361_v55_).f36, (d_1300_v3_).cf43)
            index217_ = default__.safeIndex(728, (d_1362_v56_).length(0))
            index218_ = default__.safeIndex(708, (d_1365_v57_).length(0))
            rhs251_ = ((p2) - ((0) - ((d_1361_v55_).f36))) - (((d_1367_v59_)[(d_1361_v55_).f36] if ((d_1361_v55_).f36) in (d_1367_v59_) else (0) - (len(_dafny.SeqWithoutIsStrInference([(d_1361_v55_).f36])))))
            rhs252_ = (d_1361_v55_).f36
            rhs253_ = (d_1368_v60_).cf1
            lhs234_ = d_1362_v56_
            lhs235_ = default__.safeIndex(728, (d_1362_v56_).length(0))
            lhs236_ = globalState
            lhs237_ = d_1365_v57_
            lhs238_ = default__.safeIndex(708, (d_1365_v57_).length(0))
            lhs234_[lhs235_] = rhs251_
            lhs236_.f19 = rhs252_
            lhs237_[lhs238_] = rhs253_
            d_1369_v61_: C4
            nw251_ = C4()
            nw251_.ctor__(p0, p2)
            d_1369_v61_ = nw251_
            d_1370_v62_: _dafny.Array
            nw252_ = _dafny.Array(None, 20)
            nw252_[int(0)] = d_1369_v61_
            nw252_[int(1)] = d_1369_v61_
            nw252_[int(2)] = d_1369_v61_
            nw252_[int(3)] = d_1369_v61_
            nw252_[int(4)] = d_1369_v61_
            nw252_[int(5)] = d_1369_v61_
            nw252_[int(6)] = d_1369_v61_
            nw252_[int(7)] = d_1369_v61_
            nw252_[int(8)] = d_1369_v61_
            nw252_[int(9)] = d_1369_v61_
            nw252_[int(10)] = d_1369_v61_
            nw252_[int(11)] = d_1369_v61_
            nw252_[int(12)] = d_1369_v61_
            nw252_[int(13)] = d_1369_v61_
            nw252_[int(14)] = d_1369_v61_
            nw252_[int(15)] = d_1369_v61_
            nw252_[int(16)] = d_1369_v61_
            nw252_[int(17)] = d_1369_v61_
            nw252_[int(18)] = (d_1369_v61_ if (d_1361_v55_).f35 else d_1369_v61_)
            nw252_[int(19)] = d_1369_v61_
            d_1370_v62_ = nw252_
            d_1371_v63_: D13
            d_1371_v63_ = D13_DC40(d_1369_v61_)
            d_1372_v64_: _dafny.Seq
            d_1372_v64_ = _dafny.SeqWithoutIsStrInference([d_1369_v61_])
            d_1373_v65_: _dafny.Map
            d_1373_v65_ = _dafny.Map({(self).f29: (d_1372_v64_)[default__.safeIndex(826, len(d_1372_v64_))]})
            d_1374_v66_: D3
            d_1374_v66_ = D3_DC11((d_1361_v55_).f35, (d_1369_v61_).f33)
            nw253_ = _dafny.Array(None, 18)
            nw253_[int(0)] = d_1369_v61_
            nw253_[int(1)] = ((d_1371_v63_).cf68 if (d_1361_v55_).f35 else d_1369_v61_)
            nw253_[int(2)] = d_1369_v61_
            nw253_[int(3)] = d_1369_v61_
            nw253_[int(4)] = ((d_1373_v65_)[(d_1369_v61_).f33] if ((d_1369_v61_).f33) in (d_1373_v65_) else d_1369_v61_)
            nw253_[int(5)] = d_1369_v61_
            nw253_[int(6)] = d_1369_v61_
            nw253_[int(7)] = d_1369_v61_
            nw253_[int(8)] = d_1369_v61_
            nw253_[int(9)] = d_1369_v61_
            nw253_[int(10)] = (d_1369_v61_ if p0 else d_1369_v61_)
            nw253_[int(11)] = (d_1369_v61_ if (self).f29 else d_1369_v61_)
            nw253_[int(12)] = d_1369_v61_
            nw253_[int(13)] = d_1369_v61_
            nw253_[int(14)] = d_1369_v61_
            nw253_[int(15)] = d_1369_v61_
            nw253_[int(16)] = (d_1369_v61_ if (d_1374_v66_).cf20 else d_1369_v61_)
            nw253_[int(17)] = d_1369_v61_
            d_1370_v62_ = nw253_
        d_1375_v68_: _dafny.Array
        def lambda66_(d_1376_p0_):
            def lambda67_(d_1377_i7_):
                def iife114_():
                    coll48_ = _dafny.Map()
                    compr_48_: str
                    for compr_48_ in (self.f30).Elements:
                        d_1378_v67_: str = compr_48_
                        if (d_1378_v67_) in (self.f30):
                            coll48_[d_1378_v67_] = False
                    return _dafny.Map(coll48_)
                return (_dafny.MultiSet([iife114_()
                ])).issubset(_dafny.MultiSet([_dafny.Map({_dafny.CodePoint('p'): d_1376_p0_})]))

            return lambda67_

        init38_ = lambda66_(p0)
        nw254_ = _dafny.Array(None, 28)
        for i0_38_ in range(nw254_.length(0)):
            nw254_[i0_38_] = init38_(i0_38_)
        d_1375_v68_ = nw254_
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, (d_1375_v68_).length(0)):
            d_1379_i6_: int = guard_loop_9_
            if (True) and (((0) <= (d_1379_i6_)) and ((d_1379_i6_) < ((d_1375_v68_).length(0)))):
                (d_1375_v68_)[(d_1379_i6_)] = (D11_DC35(self.f30, p0)).cf58
        d_1380_v69_: _dafny.Array
        nw255_ = _dafny.Array(None, 14)
        nw255_[int(0)] = self.f30
        nw255_[int(1)] = self.f30
        nw255_[int(2)] = self.f30
        nw255_[int(3)] = self.f30
        nw255_[int(4)] = (self.f30) + (self.f30)
        nw255_[int(5)] = self.f30
        nw255_[int(6)] = self.f30
        nw255_[int(7)] = self.f30
        nw255_[int(8)] = self.f30
        nw255_[int(9)] = self.f30
        nw255_[int(10)] = self.f30
        nw255_[int(11)] = (self.f30) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnvokyy")))
        nw255_[int(12)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqootstpb"))) + (self.f30)
        nw255_[int(13)] = (self.f30 if (self).f29 else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1381_i9_ in range(default__.abs(-207))]))
        d_1380_v69_ = nw255_
        guard_loop_10_: int
        for guard_loop_10_ in _dafny.IntegerRange(0, (d_1380_v69_).length(0)):
            d_1382_i8_: int = guard_loop_10_
            if (True) and (((0) <= (d_1382_i8_)) and ((d_1382_i8_) < ((d_1380_v69_).length(0)))):
                (d_1380_v69_)[(d_1382_i8_)] = (self.f30) + (self.f30)
        d_1383_v70_: _dafny.Seq
        d_1383_v70_ = _dafny.SeqWithoutIsStrInference([542])
        r0 = default__.fm34(p2, default__.fm36(globalState), p0, (d_1383_v70_) <= (d_1383_v70_), globalState)
        r1 = p0
        return r0, r1

    def m5(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        if (self).f29:
            d_1384_v0_: _dafny.Set
            d_1384_v0_ = _dafny.Set({len(_dafny.Set({p1}))})
            if ((d_1384_v0_) | (d_1384_v0_)) == (d_1384_v0_):
                d_1385_v1_: _dafny.Array
                nw256_ = _dafny.Array(D2.default()(), 13)
                d_1385_v1_ = nw256_
                d_1386_v2_: str
                d_1386_v2_ = _dafny.CodePoint('q')
                d_1387_v3_: _dafny.MultiSet
                d_1387_v3_ = _dafny.MultiSet([_dafny.CodePoint('i'), d_1386_v2_])
                d_1388_v4_: D2
                d_1388_v4_ = D2_DC6(d_1387_v3_)
                index219_ = default__.safeIndex(210, (d_1385_v1_).length(0))
                (d_1385_v1_)[index219_] = d_1388_v4_
                index220_ = default__.safeIndex(210, (d_1385_v1_).length(0))
                (d_1385_v1_)[index220_] = d_1388_v4_
                d_1389_v5_: C0
                nw257_ = C0()
                nw257_.ctor__(p1)
                d_1389_v5_ = nw257_
                d_1390_v6_: _dafny.Seq
                d_1390_v6_ = _dafny.SeqWithoutIsStrInference([not (d_1389_v5_.f37) or ((self).f29)])
                d_1390_v6_ = (d_1390_v6_).set(default__.safeIndex(p2, len(d_1390_v6_)), d_1389_v5_.f37)
                (globalState).f21 = p2
                d_1386_v2_ = default__.fm42(not (p1) or (p1), (self).f29, globalState)
            elif True:
                d_1391_v7_: _dafny.Map
                d_1391_v7_ = _dafny.Map({len(default__.fm48(globalState)): (0) - (p2)})
                d_1392_v8_: _dafny.Map
                d_1392_v8_ = d_1391_v7_
                d_1393_v9_: _dafny.Seq
                d_1393_v9_ = _dafny.SeqWithoutIsStrInference([(d_1392_v8_), d_1391_v7_])
                d_1394_v10_: str
                d_1394_v10_ = _dafny.CodePoint('k')
                d_1391_v7_ = ((d_1393_v9_)[default__.safeIndex(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rksaphr"))).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rksaphr")))), d_1394_v10_)), len(d_1393_v9_))]) | ((_dafny.Map({p2: p2})) | ((d_1391_v7_).set(p2, p2)))
                (globalState).f13 = (not((self).f29)) and ((self).f29)
                (globalState).f2 = (605 if p1 else p2)
                d_1394_v10_ = d_1394_v10_
                (globalState).f14 = p2
            d_1395_v11_: _dafny.MultiSet
            d_1395_v11_ = _dafny.MultiSet([p2])
            d_1395_v11_ = (d_1395_v11_).intersection(d_1395_v11_)
            d_1396_v12_: _dafny.Set
            d_1396_v12_ = _dafny.Set({p1})
            d_1397_v13_: _dafny.Map
            d_1397_v13_ = _dafny.Map({len((d_1396_v12_) | (d_1396_v12_)): (p2) == (p2)})
            d_1397_v13_ = d_1397_v13_
            d_1398_v14_: D15
            d_1398_v14_ = D15_DC43(d_1397_v13_)
            rhs254_ = (d_1398_v14_).cf70
            rhs255_ = (self).f29
            lhs239_ = globalState
            d_1397_v13_ = rhs254_
            lhs239_.f28 = rhs255_
            d_1399_v15_: C0
            nw258_ = C0()
            nw258_.ctor__(True)
            d_1399_v15_ = nw258_
        elif True:
            d_1400_v16_: str
            d_1400_v16_ = _dafny.CodePoint('j')
            d_1401_v17_: _dafny.Map
            d_1401_v17_ = _dafny.Map({-76: not((self).f29)})
            d_1402_v18_: D10
            d_1402_v18_ = D10_DC32(d_1400_v16_, p1)
            d_1403_v19_: _dafny.Array
            nw259_ = _dafny.Array(None, 13)
            nw259_[int(0)] = default__.fm49(d_1400_v16_, globalState)
            nw259_[int(1)] = D10_DC32(_dafny.CodePoint('j'), ((d_1401_v17_)[p2] if (p2) in (d_1401_v17_) else (self).f29))
            nw259_[int(2)] = d_1402_v18_
            nw259_[int(3)] = d_1402_v18_
            nw259_[int(4)] = d_1402_v18_
            nw259_[int(5)] = d_1402_v18_
            nw259_[int(6)] = d_1402_v18_
            nw259_[int(7)] = d_1402_v18_
            nw259_[int(8)] = d_1402_v18_
            nw259_[int(9)] = d_1402_v18_
            nw259_[int(10)] = D10_DC32(d_1400_v16_, p1)
            nw259_[int(11)] = D10_DC32(d_1400_v16_, p1)
            nw259_[int(12)] = d_1402_v18_
            d_1403_v19_ = nw259_
            d_1404_v20_: _dafny.Array
            nw260_ = _dafny.Array(int(0), 25)
            d_1404_v20_ = nw260_
            d_1405_v21_: _dafny.Seq
            d_1405_v21_ = _dafny.SeqWithoutIsStrInference([p1, (self).f29, p1])
            d_1406_v22_: D15
            d_1406_v22_ = D15_DC44((0) - (p2), p1, d_1403_v19_, d_1404_v20_, len(d_1405_v21_))
            d_1406_v22_ = (d_1406_v22_ if False else d_1406_v22_)
            d_1407_v23_: _dafny.MultiSet
            d_1407_v23_ = _dafny.MultiSet([p2])
            d_1408_v24_: _dafny.MultiSet
            d_1408_v24_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([(self).f29, p1, (self).f29])) + (default__.fm35(_dafny.SeqWithoutIsStrInference([D2_DC8(_dafny.CodePoint('n'), (self).f29) for d_1409_i0_ in range(default__.abs(-147))]), (d_1407_v23_).cardinality, globalState))])
            (globalState).f1 = ((d_1408_v24_)[(d_1405_v21_) + (d_1405_v21_)] if ((d_1405_v21_) + (d_1405_v21_)) in (d_1408_v24_) else p2)
            index221_ = default__.safeIndex(828, (d_1404_v20_).length(0))
            (d_1404_v20_)[index221_] = default__.fm9(globalState)
            index222_ = default__.safeIndex(828, (d_1404_v20_).length(0))
            (d_1404_v20_)[index222_] = (p2) - (p2)
            d_1410_v25_: C4
            nw261_ = C4()
            nw261_.ctor__(p1, p2)
            d_1410_v25_ = nw261_
            (globalState).f28 = (self).f29
        d_1411_v26_: _dafny.Array
        nw262_ = _dafny.Array(int(0), 15)
        d_1411_v26_ = nw262_
        index223_ = default__.safeIndex(628, (d_1411_v26_).length(0))
        (d_1411_v26_)[index223_] = default__.safeModuloInt(p2, (0) - (p2))
        index224_ = default__.safeIndex(628, (d_1411_v26_).length(0))
        (d_1411_v26_)[index224_] = default__.fm36(globalState)
        index225_ = default__.safeIndex(628, (d_1411_v26_).length(0))
        (d_1411_v26_)[index225_] = default__.safeModuloInt(p2, (d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))])
        hi9_ = (p2) + (p2)
        for d_1412_i1_ in range((d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))], hi9_):
            d_1413_v27_: C1
            nw263_ = C1()
            nw263_.ctor__(not((self).f29))
            d_1413_v27_ = nw263_
            d_1414_v28_: _dafny.Seq
            d_1414_v28_ = _dafny.SeqWithoutIsStrInference([self.f30])
            d_1415_v29_: _dafny.Seq
            d_1415_v29_ = _dafny.SeqWithoutIsStrInference([(d_1414_v28_)[default__.safeIndex(d_1412_i1_, len(d_1414_v28_))]])
            d_1416_v30_: _dafny.MultiSet
            d_1416_v30_ = _dafny.MultiSet([p0, self.f30, p0, p0])
            d_1417_v31_: _dafny.Map
            d_1417_v31_ = _dafny.Map({False: True})
            d_1418_v32_: str
            d_1418_v32_ = _dafny.CodePoint('k')
            d_1419_v33_: _dafny.MultiSet
            d_1419_v33_ = _dafny.MultiSet([d_1418_v32_])
            index226_ = default__.safeIndex(628, (d_1411_v26_).length(0))
            index227_ = default__.safeIndex(628, (d_1411_v26_).length(0))
            rhs256_ = d_1412_i1_
            rhs257_ = not(((_dafny.MultiSet(d_1415_v29_)).set(self.f30, default__.abs(d_1412_i1_))).issubset(d_1416_v30_))
            rhs258_ = default__.fm33(default__.fm33((d_1413_v27_).f40, (d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))], len(d_1417_v31_), globalState), ((d_1419_v33_).set(d_1418_v32_, default__.abs(p2))).cardinality, d_1412_i1_, globalState)
            rhs259_ = d_1412_i1_
            lhs240_ = d_1411_v26_
            lhs241_ = default__.safeIndex(628, (d_1411_v26_).length(0))
            lhs242_ = globalState
            lhs243_ = globalState
            lhs244_ = d_1411_v26_
            lhs245_ = default__.safeIndex(628, (d_1411_v26_).length(0))
            lhs240_[lhs241_] = rhs256_
            lhs242_.f18 = rhs257_
            lhs243_.f28 = rhs258_
            lhs244_[lhs245_] = rhs259_
            d_1420_v34_: _dafny.Set
            d_1420_v34_ = _dafny.Set({(self).f29, (d_1413_v27_).f40, p1})
            d_1421_v35_: D6
            d_1421_v35_ = D6_DC21(d_1420_v34_)
            d_1422_v36_: _dafny.Seq
            d_1422_v36_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1421_v35_ = default__.fm50((self).f29, default__.safeDivisionInt(d_1412_i1_, d_1412_i1_), (self.f30) + (p0), (d_1422_v36_ if (d_1413_v27_).f40 else _dafny.SeqWithoutIsStrInference([(self).f29])), globalState)
            d_1423_v37_: _dafny.Seq
            d_1423_v37_ = _dafny.SeqWithoutIsStrInference([default__.fm51(globalState)])
            (globalState).f21 = len(d_1423_v37_)
        d_1424_v38_: _dafny.Seq
        d_1424_v38_ = _dafny.SeqWithoutIsStrInference([True])
        d_1425_v39_: _dafny.Map
        d_1425_v39_ = _dafny.Map({(self).f29: len(d_1424_v38_)})
        hi10_ = p2
        for d_1426_i2_ in range(default__.safeModuloInt(p2, len(d_1425_v39_)), hi10_):
            d_1427_v40_: _dafny.MultiSet
            d_1427_v40_ = _dafny.MultiSet([p1])
            d_1428_v41_: str
            d_1428_v41_ = _dafny.CodePoint('l')
            d_1429_v42_: D2
            d_1429_v42_ = D2_DC8(d_1428_v41_, p1)
            d_1430_v43_: _dafny.Seq
            d_1430_v43_ = _dafny.SeqWithoutIsStrInference([d_1429_v42_, d_1429_v42_])
            d_1427_v40_ = ((_dafny.MultiSet([True, (self).f29]) if p1 else d_1427_v40_)).intersection(_dafny.MultiSet((default__.fm35(d_1430_v43_, (d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))], globalState)) + (d_1424_v38_)))
            d_1431_v44_: D9
            d_1431_v44_ = D9_DC29((d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))], d_1426_i2_, False)
            d_1432_v45_: _dafny.Map
            d_1432_v45_ = _dafny.Map({p1: (self).f29})
            (globalState).f28 = ((d_1431_v44_).cf48) in (d_1432_v45_)
            d_1433_v46_: _dafny.MultiSet
            d_1433_v46_ = _dafny.MultiSet([d_1426_i2_, (d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))], d_1426_i2_, (d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))], p2])
            d_1434_v47_: D3
            d_1434_v47_ = D3_DC10(d_1433_v46_)
            d_1435_v48_: _dafny.Map
            d_1435_v48_ = _dafny.Map({d_1426_i2_: d_1434_v47_})
            d_1435_v48_ = (d_1435_v48_).set(p2, d_1434_v47_)
            if ((d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))]) < ((d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))]):
                (globalState).f21 = d_1426_i2_
                (globalState).f28 = (self).f29
                def lambda68_(d_1436_v26_):
                    def lambda69_(d_1437_i3_):
                        return (d_1437_i3_) + ((d_1436_v26_)[default__.safeIndex(628, (d_1436_v26_).length(0))])

                    return lambda69_

                init39_ = lambda68_(d_1411_v26_)
                nw264_ = _dafny.Array(None, 21)
                for i0_39_ in range(nw264_.length(0)):
                    nw264_[i0_39_] = init39_(i0_39_)
                (globalState).f17 = nw264_
                d_1428_v41_ = d_1428_v41_
                rhs260_ = (d_1431_v44_).cf46
                rhs261_ = d_1428_v41_
                lhs246_ = globalState
                lhs246_.f19 = rhs260_
                d_1428_v41_ = rhs261_
            elif True:
                (globalState).f15 = not (not (p1) or (False)) or ((self).f29)
                (globalState).f20 = p0
                (globalState).f14 = p2
                d_1438_v49_: C0
                nw265_ = C0()
                nw265_.ctor__(p1)
                d_1438_v49_ = nw265_
                rhs262_ = d_1438_v49_
                rhs263_ = ((p1) and (default__.fm19(d_1428_v41_, globalState))) in (d_1425_v39_)
                lhs247_ = globalState
                d_1438_v49_ = rhs262_
                lhs247_.f28 = rhs263_
                d_1439_v50_: _dafny.Array
                nw266_ = _dafny.Array(False, 28)
                d_1439_v50_ = nw266_
                index228_ = default__.safeIndex(168, (d_1439_v50_).length(0))
                (d_1439_v50_)[index228_] = (p1) and (d_1438_v49_.f37)
                d_1440_v51_: _dafny.Seq
                d_1440_v51_ = _dafny.SeqWithoutIsStrInference([(d_1427_v40_).intersection(_dafny.MultiSet([p1, p1, p1]))])
                index229_ = default__.safeIndex(168, (d_1439_v50_).length(0))
                rhs264_ = d_1438_v49_.f37
                rhs265_ = d_1438_v49_.f37
                rhs266_ = d_1440_v51_
                rhs267_ = p1
                lhs248_ = globalState
                lhs249_ = d_1439_v50_
                lhs250_ = default__.safeIndex(168, (d_1439_v50_).length(0))
                lhs251_ = globalState
                lhs248_.f18 = rhs264_
                lhs249_[lhs250_] = rhs265_
                d_1440_v51_ = rhs266_
                lhs251_.f28 = rhs267_
        d_1441_v52_: _dafny.Array
        nw267_ = _dafny.Array(_dafny.Map({}), 21)
        d_1441_v52_ = nw267_
        d_1442_v53_: _dafny.Map
        d_1442_v53_ = _dafny.Map({len(d_1424_v38_): p2})
        index230_ = default__.safeIndex(647, (d_1441_v52_).length(0))
        (d_1441_v52_)[index230_] = d_1442_v53_
        d_1443_v55_: str
        d_1443_v55_ = _dafny.CodePoint('t')
        d_1444_v56_: _dafny.Map
        d_1444_v56_ = _dafny.Map({default__.fm19(d_1443_v55_, globalState): p1})
        index231_ = default__.safeIndex(647, (d_1441_v52_).length(0))
        def iife115_():
            coll49_ = _dafny.Map()
            compr_49_: int
            for compr_49_ in (_dafny.MultiSet([default__.fm36(globalState), len(d_1444_v56_), (d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))], -57])).Elements:
                d_1445_v54_: int = compr_49_
                if (d_1445_v54_) in (_dafny.MultiSet([default__.fm36(globalState), len(d_1444_v56_), (d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))], -57])):
                    coll49_[default__.safeDivisionInt(d_1445_v54_, p2)] = default__.safeDivisionInt((d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))], len(_dafny.Map({93: p2})))
            return _dafny.Map(coll49_)
        (d_1441_v52_)[index231_] = iife115_()
        
        r0 = p1
        d_1446_v57_: _dafny.MultiSet
        d_1446_v57_ = _dafny.MultiSet([(d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))]])
        r1 = default__.fm34((p2) - ((d_1411_v26_)[default__.safeIndex(628, (d_1411_v26_).length(0))]), (-80) + (p2), (self).f29, (d_1446_v57_).issubset(_dafny.MultiSet([p2, p2])), globalState)
        return r0, r1


class C7(T1, T0):
    def  __init__(self):
        self._f30: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f29: bool = False
        self._f31: bool = False
        self._f32: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f30(self):
        return self._f30
    @f30.setter
    def f30(self, value):
        self._f30 = value
    @property
    def f29(self):
        return self._f29
    def ctor__(self, f31, f32, f29, f30):
        (self)._f31 = f31
        (self)._f32 = f32
        (self)._f29 = f29
        (self).f30 = f30

    def m2(self, p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_1447_v0_: int
        d_1447_v0_ = 601
        d_1448_v1_: _dafny.Array
        nw268_ = _dafny.Array(int(0), 6)
        d_1448_v1_ = nw268_
        d_1449_v2_: _dafny.Map
        d_1449_v2_ = _dafny.Map({d_1447_v0_: d_1448_v1_})
        d_1450_v3_: _dafny.Array
        nw269_ = _dafny.Array(None, 11)
        nw269_[int(0)] = d_1447_v0_
        nw269_[int(1)] = d_1447_v0_
        nw269_[int(2)] = d_1447_v0_
        nw269_[int(3)] = d_1447_v0_
        nw269_[int(4)] = d_1447_v0_
        nw269_[int(5)] = d_1447_v0_
        nw269_[int(6)] = d_1447_v0_
        nw269_[int(7)] = len(d_1449_v2_)
        nw269_[int(8)] = d_1447_v0_
        nw269_[int(9)] = d_1447_v0_
        nw269_[int(10)] = d_1447_v0_
        d_1450_v3_ = nw269_
        d_1451_v4_: _dafny.Set
        d_1451_v4_ = _dafny.Set({d_1450_v3_, d_1450_v3_, d_1448_v1_, d_1448_v1_, d_1448_v1_})
        d_1452_v5_: C2
        nw270_ = C2()
        nw270_.ctor__((-639) + (d_1447_v0_), (d_1451_v4_) - (d_1451_v4_), p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nicg")))
        d_1452_v5_ = nw270_
        index232_ = default__.safeIndex(911, (d_1448_v1_).length(0))
        (d_1448_v1_)[index232_] = d_1447_v0_
        index233_ = default__.safeIndex(911, (d_1448_v1_).length(0))
        (d_1448_v1_)[index233_] = d_1447_v0_
        d_1453_i0_: int
        d_1453_i0_ = 0
        with _dafny.label("7"):
            while default__.fm33(False, 689, (d_1448_v1_)[default__.safeIndex(911, (d_1448_v1_).length(0))], globalState):
                with _dafny.c_label("7"):
                    if (d_1453_i0_) >= (100):
                        raise _dafny.Break("7")
                    d_1453_i0_ = (d_1453_i0_) + (1)
                    d_1454_v6_: _dafny.Set
                    d_1454_v6_ = _dafny.Set({not(p0)})
                    d_1454_v6_ = ((d_1454_v6_ if (self).f29 else d_1454_v6_)) - ((d_1454_v6_) | (d_1454_v6_))
                    d_1455_v7_: str
                    d_1455_v7_ = _dafny.CodePoint('u')
                    d_1456_v8_: _dafny.Map
                    d_1456_v8_ = _dafny.Map({d_1450_v3_: (self).f31})
                    index234_ = default__.safeIndex(911, (d_1448_v1_).length(0))
                    rhs268_ = d_1455_v7_
                    rhs269_ = d_1456_v8_
                    rhs270_ = -103
                    rhs271_ = (d_1452_v5_).f38
                    rhs272_ = not(not(default__.fm10((d_1452_v5_).f38, (d_1452_v5_).f38, globalState)))
                    lhs252_ = d_1448_v1_
                    lhs253_ = default__.safeIndex(911, (d_1448_v1_).length(0))
                    lhs254_ = globalState
                    d_1455_v7_ = rhs268_
                    d_1456_v8_ = rhs269_
                    lhs252_[lhs253_] = rhs270_
                    lhs254_.f21 = rhs271_
                    r1 = rhs272_
                    d_1457_v9_: _dafny.Seq
                    d_1457_v9_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_1458_v10_: _dafny.Map
                    d_1458_v10_ = _dafny.Map({len((self.f30) + (self.f30)): ((self).f31) in (d_1457_v9_)})
                    d_1459_v11_: D12
                    d_1459_v11_ = D12_DC39(460, default__.fm19(d_1455_v7_, globalState), (d_1452_v5_).f38)
                    d_1458_v10_ = (d_1458_v10_).set((d_1459_v11_).cf65, (p0) or (p0))
                    d_1460_v12_: _dafny.Map
                    d_1460_v12_ = _dafny.Map({(self).f32: d_1447_v0_})
                    d_1461_v13_: C4
                    nw271_ = C4()
                    nw271_.ctor__((self).f32, (d_1452_v5_).f38)
                    d_1461_v13_ = nw271_
                    d_1462_v14_: _dafny.Set
                    d_1462_v14_ = _dafny.Set({d_1461_v13_})
                    d_1460_v12_ = (d_1460_v12_).set((self).f29, len((d_1462_v14_) - (d_1462_v14_)))
                    pass
            pass
        d_1463_v15_: _dafny.Seq
        d_1463_v15_ = _dafny.SeqWithoutIsStrInference([d_1450_v3_, d_1448_v1_])
        d_1464_v16_: _dafny.Seq
        d_1464_v16_ = _dafny.SeqWithoutIsStrInference([d_1450_v3_, (d_1463_v15_)[default__.safeIndex(d_1447_v0_, len(d_1463_v15_))], d_1450_v3_])
        d_1465_v17_: _dafny.Array
        nw272_ = _dafny.Array(None, 5)
        nw272_[int(0)] = d_1464_v16_
        nw272_[int(1)] = _dafny.SeqWithoutIsStrInference([d_1450_v3_])
        nw272_[int(2)] = d_1463_v15_
        nw272_[int(3)] = (d_1463_v15_) + (d_1463_v15_)
        nw272_[int(4)] = d_1463_v15_
        d_1465_v17_ = nw272_
        index235_ = default__.safeIndex(63, (d_1465_v17_).length(0))
        (d_1465_v17_)[index235_] = (d_1464_v16_) + (d_1463_v15_)
        index236_ = default__.safeIndex(63, (d_1465_v17_).length(0))
        (d_1465_v17_)[index236_] = (d_1464_v16_) + (_dafny.SeqWithoutIsStrInference([d_1448_v1_, d_1450_v3_]))
        d_1466_v18_: _dafny.MultiSet
        d_1466_v18_ = _dafny.MultiSet([(0) - (d_1447_v0_)])
        d_1467_v19_: _dafny.Map
        d_1467_v19_ = _dafny.Map({(self).f29: d_1466_v18_})
        d_1467_v19_ = (d_1467_v19_).set((self).f32, _dafny.MultiSet([d_1447_v0_]))
        d_1468_v20_: D3
        d_1468_v20_ = D3_DC10((d_1466_v18_) | (d_1466_v18_))
        d_1468_v20_ = d_1468_v20_
        d_1469_v21_: _dafny.Seq
        d_1469_v21_ = _dafny.SeqWithoutIsStrInference([False])
        d_1470_v22_: _dafny.Seq
        d_1470_v22_ = _dafny.SeqWithoutIsStrInference([(d_1469_v21_) + (d_1469_v21_), _dafny.SeqWithoutIsStrInference([(self).f32]), d_1469_v21_])
        r0 = d_1470_v22_
        r1 = p0
        return r0, r1

    def m3(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        d_1471_v0_: _dafny.Map
        d_1471_v0_ = _dafny.Map({p0: 258})
        d_1472_v1_: _dafny.Map
        d_1472_v1_ = _dafny.Map({(self).f31: len(d_1471_v0_)})
        d_1473_v2_: _dafny.Seq
        d_1473_v2_ = _dafny.SeqWithoutIsStrInference([not((self).f31)])
        d_1474_v3_: _dafny.Seq
        d_1474_v3_ = _dafny.SeqWithoutIsStrInference([d_1471_v0_, d_1471_v0_, default__.fm14(not(p0), _dafny.SeqWithoutIsStrInference([len(d_1473_v2_)]), (self).f31, p2, globalState), d_1472_v1_, d_1471_v0_])
        d_1475_v4_: _dafny.Array
        nw273_ = _dafny.Array(None, 1)
        nw273_[int(0)] = not((d_1471_v0_) == ((d_1474_v3_)[default__.safeIndex(p2, len(d_1474_v3_))]))
        d_1475_v4_ = nw273_
        index237_ = default__.safeIndex(964, (d_1475_v4_).length(0))
        (d_1475_v4_)[index237_] = (self).f32
        index238_ = default__.safeIndex(964, (d_1475_v4_).length(0))
        (d_1475_v4_)[index238_] = (self).f31
        d_1475_v4_ = d_1475_v4_
        r1 = (self).f31
        d_1476_v5_: _dafny.Array
        nw274_ = _dafny.Array(int(0), 28)
        d_1476_v5_ = nw274_
        index239_ = default__.safeIndex(906, (d_1476_v5_).length(0))
        (d_1476_v5_)[index239_] = (0) - (p2)
        d_1477_v6_: _dafny.MultiSet
        d_1477_v6_ = _dafny.MultiSet([p2])
        d_1478_v7_: _dafny.MultiSet
        d_1478_v7_ = _dafny.MultiSet([(self).f31, (self).f31, (d_1473_v2_)[default__.safeIndex(default__.fm36(globalState), len(d_1473_v2_))], True])
        d_1479_v8_: _dafny.Map
        d_1479_v8_ = _dafny.Map({(d_1478_v7_) == (d_1478_v7_): not((p2) != (p2))})
        d_1480_v9_: D3
        d_1480_v9_ = D3_DC10(d_1477_v6_)
        index240_ = default__.safeIndex(906, (d_1476_v5_).length(0))
        rhs273_ = p0
        rhs274_ = len(d_1479_v8_)
        rhs275_ = (d_1480_v9_).cf18
        rhs276_ = (p2) * (p2)
        rhs277_ = d_1475_v4_
        lhs255_ = globalState
        lhs256_ = d_1476_v5_
        lhs257_ = default__.safeIndex(906, (d_1476_v5_).length(0))
        lhs258_ = globalState
        lhs255_.f13 = rhs273_
        lhs256_[lhs257_] = rhs274_
        d_1477_v6_ = rhs275_
        lhs258_.f19 = rhs276_
        d_1475_v4_ = rhs277_
        d_1481_v10_: _dafny.Array
        def lambda70_(d_1482_v2_):
            def lambda71_(d_1483_i0_):
                return (_dafny.Map({(self).f31: D1_DC3(d_1482_v2_)})) | (_dafny.Map({False: D1_DC3(d_1482_v2_)}))

            return lambda71_

        init40_ = lambda70_(d_1473_v2_)
        nw275_ = _dafny.Array(None, 26)
        for i0_40_ in range(nw275_.length(0)):
            nw275_[i0_40_] = init40_(i0_40_)
        d_1481_v10_ = nw275_
        d_1484_v11_: D1
        d_1484_v11_ = D1_DC3(d_1473_v2_)
        d_1485_v12_: _dafny.Map
        d_1485_v12_ = _dafny.Map({(self).f29: d_1484_v11_})
        index241_ = default__.safeIndex(799, (d_1481_v10_).length(0))
        (d_1481_v10_)[index241_] = d_1485_v12_
        index242_ = default__.safeIndex(799, (d_1481_v10_).length(0))
        (d_1481_v10_)[index242_] = d_1485_v12_
        d_1486_v13_: _dafny.Array
        nw276_ = _dafny.Array(None, 5)
        nw276_[int(0)] = p1
        nw276_[int(1)] = p1
        nw276_[int(2)] = p1
        nw276_[int(3)] = p1
        nw276_[int(4)] = p1
        d_1486_v13_ = nw276_
        index243_ = default__.safeIndex(187, (d_1486_v13_).length(0))
        (d_1486_v13_)[index243_] = p1
        d_1487_v14_: _dafny.Map
        d_1487_v14_ = _dafny.Map({self.f30: -828})
        d_1488_v16_: str
        d_1488_v16_ = _dafny.CodePoint('d')
        d_1489_v17_: _dafny.Set
        d_1489_v17_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1490_i1_ in range(default__.abs(371))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1491_i1_ in range(default__.abs(371))]))), d_1488_v16_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmynhama"))})
        d_1492_v18_: D5
        d_1492_v18_ = D5_DC17(d_1489_v17_)
        d_1493_v19_: D5
        d_1493_v19_ = D5_DC20(d_1492_v18_)
        pat_let_tv29_ = p2
        pat_let_tv30_ = d_1476_v5_
        pat_let_tv31_ = d_1476_v5_
        pat_let_tv32_ = d_1476_v5_
        pat_let_tv33_ = d_1476_v5_
        index244_ = default__.safeIndex(187, (d_1486_v13_).length(0))
        def iife116_():
            coll50_ = _dafny.Map()
            compr_50_: _dafny.Seq
            for compr_50_ in (d_1489_v17_).Elements:
                d_1494_v15_: _dafny.Seq = compr_50_
                if (d_1494_v15_) in (d_1489_v17_):
                    coll50_[d_1494_v15_] = (d_1477_v6_).cardinality
            return _dafny.Map(coll50_)
        def lambda72_(source21_):
            if source21_.is_DC17:
                d_1495___mcc_h0_ = source21_.cf27
                d_1496_cf27_ = d_1495___mcc_h0_
                return pat_let_tv29_
            elif source21_.is_DC18:
                return len(self.f30)
            elif source21_.is_DC19:
                d_1497___mcc_h1_ = source21_.cf28
                d_1498___mcc_h2_ = source21_.cf29
                d_1499___mcc_h3_ = source21_.cf30
                d_1500_cf30_ = d_1499___mcc_h3_
                d_1501_cf29_ = d_1498___mcc_h2_
                d_1502_cf28_ = d_1497___mcc_h1_
                return 538
            elif source21_.is_DC16:
                d_1503___mcc_h4_ = source21_.cf26
                d_1504_cf26_ = d_1503___mcc_h4_
                return (pat_let_tv31_)[default__.safeIndex(906, (pat_let_tv30_).length(0))]
            elif True:
                d_1505___mcc_h5_ = source21_.cf31
                d_1506_cf31_ = d_1505___mcc_h5_
                return (pat_let_tv33_)[default__.safeIndex(906, (pat_let_tv32_).length(0))]

        rhs278_ = p1
        rhs279_ = iife116_()

        rhs280_ = lambda72_(d_1493_v19_)
        lhs259_ = d_1486_v13_
        lhs260_ = default__.safeIndex(187, (d_1486_v13_).length(0))
        lhs261_ = globalState
        lhs259_[lhs260_] = rhs278_
        d_1487_v14_ = rhs279_
        lhs261_.f21 = rhs280_
        d_1507_v20_: _dafny.Seq
        d_1507_v20_ = _dafny.SeqWithoutIsStrInference([self.f30, _dafny.SeqWithoutIsStrInference([d_1488_v16_ for d_1508_i2_ in range(default__.abs(-759))]), (self.f30 if True else ((self.f30).set(default__.safeIndex(p2, len(self.f30)), _dafny.CodePoint('i'))).set(default__.safeIndex(len(d_1472_v1_), len((self.f30).set(default__.safeIndex(p2, len(self.f30)), _dafny.CodePoint('i')))), d_1488_v16_)), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjfyn"))), self.f30])
        r0 = (d_1507_v20_)[default__.safeIndex(p2, len(d_1507_v20_))]
        r1 = ((0) - ((d_1476_v5_)[default__.safeIndex(906, (d_1476_v5_).length(0))])) >= ((0) - ((p2) * (89)))
        return r0, r1

    def m4(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        r3: _dafny.Set = _dafny.Set({})
        d_1509_v0_: _dafny.Array
        def lambda73_(d_1510_i0_):
            return (self).f29

        init41_ = lambda73_
        nw277_ = _dafny.Array(None, 5)
        for i0_41_ in range(nw277_.length(0)):
            nw277_[i0_41_] = init41_(i0_41_)
        d_1509_v0_ = nw277_
        d_1511_v1_: _dafny.Set
        d_1511_v1_ = _dafny.Set({d_1509_v0_, d_1509_v0_})
        d_1512_v2_: C4
        nw278_ = C4()
        nw278_.ctor__((d_1509_v0_) not in ((D16_DC47(d_1511_v1_)).cf81), p0)
        d_1512_v2_ = nw278_
        d_1513_v3_: D4
        d_1513_v3_ = D4_DC15(p0, p0)
        d_1514_v4_: str
        d_1514_v4_ = _dafny.CodePoint('i')
        d_1515_v5_: D10
        d_1515_v5_ = D10_DC32(d_1514_v4_, (self).f32)
        d_1516_v6_: _dafny.Array
        nw279_ = _dafny.Array(None, 9)
        nw279_[int(0)] = d_1515_v5_
        nw279_[int(1)] = d_1515_v5_
        nw279_[int(2)] = d_1515_v5_
        nw279_[int(3)] = D10_DC32(d_1514_v4_, (self).f29)
        nw279_[int(4)] = d_1515_v5_
        nw279_[int(5)] = d_1515_v5_
        nw279_[int(6)] = d_1515_v5_
        nw279_[int(7)] = d_1515_v5_
        nw279_[int(8)] = D10_DC32(d_1514_v4_, (self).f32)
        d_1516_v6_ = nw279_
        d_1517_v7_: _dafny.Array
        nw280_ = _dafny.Array(int(0), 27)
        d_1517_v7_ = nw280_
        d_1518_v8_: _dafny.Map
        d_1518_v8_ = _dafny.Map({len(self.f30): not((d_1512_v2_).f33)})
        d_1519_v9_: D15
        d_1519_v9_ = D15_DC44((d_1513_v3_).cf25, (self).f32, d_1516_v6_, d_1517_v7_, len(d_1518_v8_))
        d_1520_v10_: D3
        d_1520_v10_ = D3_DC11((d_1512_v2_).f33, (self).f32)
        d_1521_v11_: _dafny.MultiSet
        d_1521_v11_ = _dafny.MultiSet([(d_1512_v2_).f34])
        d_1522_v12_: D11
        d_1522_v12_ = D11_DC37((d_1519_v9_).cf75, default__.fm33(not((D3_DC11((d_1520_v10_).cf19, (self).f32)).cf19), p0, (d_1521_v11_).cardinality, globalState))
        source22_ = d_1522_v12_
        if source22_.is_DC35:
            d_1523___mcc_h0_ = source22_.cf57
            d_1524___mcc_h1_ = source22_.cf58
            d_1525_cf58_ = d_1524___mcc_h1_
            d_1526_cf57_ = d_1523___mcc_h0_
            (globalState).f15 = d_1525_cf58_
            if (p1) >= (p1):
                (globalState).f2 = (default__.fm36(globalState)) - (default__.safeDivisionInt((d_1512_v2_).f34, (d_1512_v2_).f34))
                d_1527_v13_: _dafny.Set
                d_1527_v13_ = _dafny.Set({(d_1512_v2_).f34, (d_1512_v2_).f34})
                d_1528_v14_: D17
                d_1528_v14_ = D17_DC49(d_1527_v13_)
                d_1529_v15_: _dafny.Seq
                d_1529_v15_ = _dafny.SeqWithoutIsStrInference([(d_1512_v2_).f33])
                d_1530_v16_: _dafny.Map
                d_1530_v16_ = _dafny.Map({(d_1512_v2_).f34: (d_1512_v2_).f34})
                d_1531_v17_: _dafny.Map
                d_1531_v17_ = _dafny.Map({self.f30: d_1527_v13_})
                d_1532_v19_: _dafny.Array
                nw281_ = _dafny.Array(None, 20)
                nw281_[int(0)] = _dafny.Set({(d_1512_v2_).f34, (d_1512_v2_).f34, p1, p0, (d_1512_v2_).f34})
                nw281_[int(1)] = (d_1527_v13_).intersection(d_1527_v13_)
                nw281_[int(2)] = d_1527_v13_
                nw281_[int(3)] = (d_1528_v14_).cf82
                nw281_[int(4)] = d_1527_v13_
                nw281_[int(5)] = d_1527_v13_
                nw281_[int(6)] = (d_1527_v13_).intersection(d_1527_v13_)
                nw281_[int(7)] = (d_1527_v13_) - (d_1527_v13_)
                nw281_[int(8)] = d_1527_v13_
                nw281_[int(9)] = d_1527_v13_
                nw281_[int(10)] = d_1527_v13_
                nw281_[int(11)] = d_1527_v13_
                nw281_[int(12)] = d_1527_v13_
                nw281_[int(13)] = d_1527_v13_
                nw281_[int(14)] = (default__.fm17((d_1512_v2_).f34, globalState) if (d_1512_v2_).f33 else d_1527_v13_)
                nw281_[int(15)] = (d_1527_v13_) - (d_1527_v13_)
                nw281_[int(16)] = d_1527_v13_
                nw281_[int(17)] = (_dafny.Set({(d_1512_v2_).f34, len((_dafny.SeqWithoutIsStrInference([d_1514_v4_ for d_1533_i1_ in range(default__.abs(730))])).set(default__.safeIndex((d_1512_v2_).f34, len(_dafny.SeqWithoutIsStrInference([d_1514_v4_ for d_1534_i1_ in range(default__.abs(730))]))), d_1514_v4_)), len(d_1529_v15_), len(_dafny.Map({(d_1512_v2_).f33: d_1530_v16_})), p1})) - (((d_1531_v17_)[self.f30] if (self.f30) in (d_1531_v17_) else _dafny.Set({len(self.f30)})))
                def iife117_():
                    coll51_ = _dafny.Set()
                    compr_51_: int
                    for compr_51_ in _dafny.IntegerRange(735, 598):
                        d_1535_v18_: int = compr_51_
                        if ((735) <= (d_1535_v18_)) and ((d_1535_v18_) < (598)):
                            coll51_ = coll51_.union(_dafny.Set([default__.safeDivisionInt(d_1535_v18_, p0)]))
                    return _dafny.Set(coll51_)
                nw281_[int(18)] = (iife117_()
                ) - (d_1527_v13_)
                nw281_[int(19)] = d_1527_v13_
                d_1532_v19_ = nw281_
                index245_ = default__.safeIndex(275, (d_1532_v19_).length(0))
                (d_1532_v19_)[index245_] = (d_1527_v13_).intersection(d_1527_v13_)
                d_1536_v20_: _dafny.Map
                d_1536_v20_ = _dafny.Map({(d_1515_v5_).cf52: (d_1527_v13_) - (d_1527_v13_)})
                index246_ = default__.safeIndex(275, (d_1532_v19_).length(0))
                (d_1532_v19_)[index246_] = ((d_1536_v20_)[_dafny.CodePoint('u')] if (_dafny.CodePoint('u')) in (d_1536_v20_) else d_1527_v13_)
                d_1537_v21_: _dafny.Map
                d_1537_v21_ = _dafny.Map({d_1509_v0_: d_1514_v4_})
                d_1538_v22_: D0
                d_1538_v22_ = D0_DC1(d_1509_v0_, p0, len(d_1526_cf57_))
                d_1537_v21_ = (d_1537_v21_).set((d_1538_v22_).cf1, d_1514_v4_)
                d_1539_v23_: C5
                nw282_ = C5()
                nw282_.ctor__(((self).f32 if (self).f32 else (self).f31), d_1526_cf57_)
                d_1539_v23_ = nw282_
                d_1540_v24_: _dafny.MultiSet
                d_1540_v24_ = _dafny.MultiSet([(d_1512_v2_).f33, (d_1512_v2_).f33])
                d_1541_v25_: _dafny.Seq
                d_1541_v25_ = _dafny.SeqWithoutIsStrInference([(d_1540_v24_).cardinality])
                d_1542_v26_: C1
                nw283_ = C1()
                nw283_.ctor__((default__.fm44((d_1512_v2_).f33, len(_dafny.SeqWithoutIsStrInference([(d_1512_v2_).f34 for d_1543_i2_ in range(default__.abs(629))])), (d_1512_v2_).f33, (d_1512_v2_).f33, globalState)) != (d_1541_v25_))
                d_1542_v26_ = nw283_
            elif True:
                (globalState).f13 = ((self.f30) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnewr")))) == ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wf"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efgrsogq"))))
                (globalState).f15 = (d_1512_v2_).f33
                d_1544_v27_: _dafny.Array
                nw284_ = _dafny.Array(D2.default()(), 23)
                d_1544_v27_ = nw284_
                d_1545_v28_: D12
                d_1545_v28_ = D12_DC39(p1, (d_1512_v2_).f33, 146)
                d_1546_v29_: D2
                d_1546_v29_ = D2_DC9(self.f30, p1, d_1525_cf58_, (d_1545_v28_).cf67)
                index247_ = default__.safeIndex(556, (d_1544_v27_).length(0))
                (d_1544_v27_)[index247_] = d_1546_v29_
                index248_ = default__.safeIndex(556, (d_1544_v27_).length(0))
                (d_1544_v27_)[index248_] = d_1546_v29_
                d_1547_v30_: D17
                d_1547_v30_ = D17_DC52(default__.fm6(len(d_1526_cf57_), globalState), (d_1512_v2_).f33, (d_1512_v2_).f34, default__.fm27((d_1512_v2_).f34, True, globalState))
                index249_ = default__.safeIndex(100, (d_1517_v7_).length(0))
                (d_1517_v7_)[index249_] = default__.safeModuloInt((d_1547_v30_).cf90, (d_1547_v30_).cf90)
                index250_ = default__.safeIndex(100, (d_1517_v7_).length(0))
                (d_1517_v7_)[index250_] = -602
                d_1548_v32_: _dafny.Map
                def iife118_():
                    coll52_ = _dafny.Map()
                    compr_52_: int
                    for compr_52_ in (_dafny.SeqWithoutIsStrInference([(d_1512_v2_).f34 for d_1549_i3_ in range(default__.abs(190))])).Elements:
                        d_1550_v31_: int = compr_52_
                        if (d_1550_v31_) in (_dafny.SeqWithoutIsStrInference([(d_1512_v2_).f34 for d_1549_i3_ in range(default__.abs(190))])):
                            coll52_[default__.safeModuloInt(d_1550_v31_, (d_1512_v2_).f34)] = (d_1512_v2_).f34
                    return _dafny.Map(coll52_)
                d_1548_v32_ = _dafny.Map({622: len(iife118_()
                )})
                (globalState).f19 = (len(d_1548_v32_)) * ((d_1512_v2_).f34)
            d_1551_v33_: _dafny.Seq
            d_1551_v33_ = _dafny.SeqWithoutIsStrInference([p0])
            if not(((_dafny.SeqWithoutIsStrInference([899 for d_1552_i4_ in range(default__.abs(301))])) + (_dafny.SeqWithoutIsStrInference([(d_1512_v2_).f34 for d_1553_i5_ in range(default__.abs(-501))]))) <= (d_1551_v33_)):
                d_1514_v4_ = d_1514_v4_
                (globalState).f13 = not((d_1525_cf58_) or ((d_1526_cf57_) < (self.f30)))
                d_1521_v11_ = (d_1521_v11_).intersection(d_1521_v11_)
                index251_ = default__.safeIndex(538, (d_1517_v7_).length(0))
                (d_1517_v7_)[index251_] = p1
                index252_ = default__.safeIndex(538, (d_1517_v7_).length(0))
                rhs281_ = (d_1551_v33_)[default__.safeIndex((847) - (len(d_1551_v33_)), len(d_1551_v33_))]
                rhs282_ = ((d_1512_v2_).f33 if (d_1512_v2_).f33 else d_1525_cf58_)
                rhs283_ = ((default__.fm34(p1, (d_1512_v2_).f34, (self).f31, not(default__.fm33((self).f29, p1, 848, globalState)), globalState)) + (d_1526_cf57_) if (d_1512_v2_).f33 else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))) + (self.f30))
                rhs284_ = (d_1512_v2_).f34
                rhs285_ = d_1517_v7_
                lhs262_ = d_1517_v7_
                lhs263_ = default__.safeIndex(538, (d_1517_v7_).length(0))
                lhs264_ = globalState
                lhs265_ = self
                lhs266_ = globalState
                lhs267_ = globalState
                lhs262_[lhs263_] = rhs281_
                lhs264_.f18 = rhs282_
                lhs265_.f30 = rhs283_
                lhs266_.f19 = rhs284_
                lhs267_.f17 = rhs285_
                (globalState).f1 = (d_1512_v2_).f34
            elif True:
                index253_ = default__.safeIndex(894, (d_1509_v0_).length(0))
                (d_1509_v0_)[index253_] = False
                index254_ = default__.safeIndex(894, (d_1509_v0_).length(0))
                (d_1509_v0_)[index254_] = (p1) == ((d_1512_v2_).f34)
                d_1554_v34_: _dafny.Set
                d_1554_v34_ = _dafny.Set({(d_1512_v2_).f33, (self).f31})
                d_1555_v35_: _dafny.Seq
                d_1555_v35_ = _dafny.SeqWithoutIsStrInference([(self).f31, d_1525_cf58_])
                d_1556_v36_: _dafny.Map
                d_1556_v36_ = _dafny.Map({(d_1512_v2_).f33: default__.fm20((d_1509_v0_)[default__.safeIndex(894, (d_1509_v0_).length(0))], (d_1509_v0_)[default__.safeIndex(894, (d_1509_v0_).length(0))], d_1555_v35_, (d_1512_v2_).f34, globalState)})
                rhs286_ = (d_1521_v11_).ispropersubset(d_1521_v11_)
                rhs287_ = (self.f30) + (default__.fm34(len(d_1554_v34_), len(((d_1556_v36_)[False] if (False) in (d_1556_v36_) else d_1526_cf57_)), (self).f31, (d_1509_v0_)[default__.safeIndex(894, (d_1509_v0_).length(0))], globalState))
                rhs288_ = (d_1551_v33_)[default__.safeIndex(((d_1512_v2_).f34) + ((d_1512_v2_).f34), len(d_1551_v33_))]
                rhs289_ = p1
                lhs268_ = globalState
                lhs269_ = globalState
                lhs270_ = globalState
                lhs271_ = globalState
                lhs268_.f18 = rhs286_
                lhs269_.f11 = rhs287_
                lhs270_.f1 = rhs288_
                lhs271_.f14 = rhs289_
                d_1557_v37_: _dafny.Set
                d_1557_v37_ = _dafny.Set({d_1517_v7_})
                d_1558_v38_: _dafny.Seq
                d_1558_v38_ = _dafny.SeqWithoutIsStrInference([d_1557_v37_])
                d_1559_v39_: C2
                nw285_ = C2()
                nw285_.ctor__((d_1512_v2_).f34, (d_1558_v38_)[default__.safeIndex((d_1512_v2_).f34, len(d_1558_v38_))], (self).f29, self.f30)
                d_1559_v39_ = nw285_
                d_1560_v40_: _dafny.MultiSet
                d_1560_v40_ = _dafny.MultiSet([d_1517_v7_, d_1517_v7_])
                d_1561_v41_: _dafny.Array
                nw286_ = _dafny.Array(False, 20)
                d_1561_v41_ = nw286_
                d_1562_v42_: _dafny.MultiSet
                d_1562_v42_ = _dafny.MultiSet([d_1561_v41_, d_1561_v41_])
                index255_ = default__.safeIndex(894, (d_1509_v0_).length(0))
                (d_1509_v0_)[index255_] = ((d_1560_v40_) | ((d_1560_v40_).set(d_1517_v7_, default__.abs((d_1562_v42_).cardinality)))).issubset((_dafny.MultiSet([d_1517_v7_])) | (d_1560_v40_))
                (globalState).f2 = default__.fm9(globalState)
            d_1563_v43_: C3
            nw287_ = C3()
            nw287_.ctor__(True, ((d_1512_v2_).f34) * (p1), (self).f32, d_1526_cf57_)
            d_1563_v43_ = nw287_
        elif source22_.is_DC36:
            d_1564___mcc_h2_ = source22_.cf59
            d_1565___mcc_h3_ = source22_.cf60
            d_1566___mcc_h4_ = source22_.cf61
            d_1567_cf61_ = d_1566___mcc_h4_
            d_1568_cf60_ = d_1565___mcc_h3_
            d_1569_cf59_ = d_1564___mcc_h2_
            d_1570_v44_: _dafny.Seq
            d_1570_v44_ = _dafny.SeqWithoutIsStrInference([(d_1512_v2_).f33, True, ((d_1512_v2_).f33 if (self).f32 else (d_1512_v2_).f33), (self).f29])
            (globalState).f13 = (d_1570_v44_)[default__.safeIndex(default__.fm9(globalState), len(d_1570_v44_))]
            index256_ = default__.safeIndex(139, (d_1517_v7_).length(0))
            (d_1517_v7_)[index256_] = (d_1512_v2_).f34
            d_1571_v45_: _dafny.Set
            d_1571_v45_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_1514_v4_ for d_1572_i6_ in range(default__.abs(250))]))})
            index257_ = default__.safeIndex(139, (d_1517_v7_).length(0))
            (d_1517_v7_)[index257_] = (p0) + ((len(d_1571_v45_)) - (p1))
            d_1573_v46_: D3
            d_1573_v46_ = D3_DC10(d_1521_v11_)
            d_1574_v47_: D3
            d_1574_v47_ = D3_DC12(d_1573_v46_)
            source23_ = (d_1574_v47_ if (self).f29 else d_1574_v47_)
            if source23_.is_DC11:
                d_1575___mcc_h8_ = source23_.cf19
                d_1576___mcc_h9_ = source23_.cf20
                d_1577_cf20_ = d_1576___mcc_h9_
                d_1578_cf19_ = d_1575___mcc_h8_
                d_1579_v48_: C1
                nw288_ = C1()
                nw288_.ctor__((self).f29)
                d_1579_v48_ = nw288_
                d_1580_v49_: _dafny.Seq
                d_1580_v49_ = _dafny.SeqWithoutIsStrInference([d_1579_v48_, d_1579_v48_])
                d_1581_v50_: _dafny.Map
                d_1581_v50_ = _dafny.Map({d_1577_cf20_: _dafny.MultiSet([d_1579_v48_, d_1579_v48_, (d_1580_v49_)[default__.safeIndex((d_1512_v2_).f34, len(d_1580_v49_))], d_1579_v48_])})
                d_1582_v51_: _dafny.Map
                d_1582_v51_ = _dafny.Map({d_1581_v50_: ((d_1518_v8_)[(d_1517_v7_)[default__.safeIndex(139, (d_1517_v7_).length(0))]] if ((d_1517_v7_)[default__.safeIndex(139, (d_1517_v7_).length(0))]) in (d_1518_v8_) else not(False))})
                d_1583_v52_: _dafny.MultiSet
                d_1583_v52_ = _dafny.MultiSet([d_1579_v48_])
                d_1584_v53_: _dafny.Seq
                d_1584_v53_ = _dafny.SeqWithoutIsStrInference([p1, (0) - ((d_1517_v7_)[default__.safeIndex(139, (d_1517_v7_).length(0))]), (len(d_1570_v44_)) + (len(_dafny.Set({(d_1517_v7_)[default__.safeIndex(139, (d_1517_v7_).length(0))], (d_1512_v2_).f34, p0}))), p1])
                index258_ = default__.safeIndex(139, (d_1517_v7_).length(0))
                rhs290_ = ((d_1582_v51_)[(d_1581_v50_) | (_dafny.Map({d_1567_cf61_: d_1583_v52_}))] if ((d_1581_v50_) | (_dafny.Map({d_1567_cf61_: d_1583_v52_}))) in (d_1582_v51_) else ((self).f31) == (d_1578_cf19_))
                rhs291_ = (default__.safeModuloInt((d_1512_v2_).f34, (d_1512_v2_).f34)) - ((d_1512_v2_).f34)
                rhs292_ = d_1584_v53_
                lhs272_ = globalState
                lhs273_ = d_1517_v7_
                lhs274_ = default__.safeIndex(139, (d_1517_v7_).length(0))
                lhs275_ = globalState
                lhs272_.f28 = rhs290_
                lhs273_[lhs274_] = rhs291_
                lhs275_.f7 = rhs292_
                d_1585_v54_: _dafny.Seq
                d_1585_v54_ = _dafny.SeqWithoutIsStrInference([d_1574_v47_, default__.fm29((d_1512_v2_).f33, (d_1517_v7_)[default__.safeIndex(139, (d_1517_v7_).length(0))], d_1567_cf61_, globalState)])
                rhs293_ = not (d_1577_cf20_) or (d_1577_cf20_)
                rhs294_ = (847) != (len((d_1585_v54_) + (d_1585_v54_)))
                rhs295_ = (d_1568_cf60_) | ((d_1568_cf60_ if d_1577_cf20_ else d_1568_cf60_))
                rhs296_ = (p0) * ((d_1512_v2_).f34)
                lhs276_ = globalState
                lhs277_ = globalState
                lhs278_ = globalState
                lhs276_.f15 = rhs293_
                lhs277_.f18 = rhs294_
                d_1568_cf60_ = rhs295_
                lhs278_.f14 = rhs296_
                d_1586_v55_: _dafny.Array
                def lambda74_(d_1587_v4_):
                    def lambda75_(d_1588_i7_):
                        return d_1587_v4_

                    return lambda75_

                init42_ = lambda74_(d_1514_v4_)
                nw289_ = _dafny.Array(None, 21)
                for i0_42_ in range(nw289_.length(0)):
                    nw289_[i0_42_] = init42_(i0_42_)
                d_1586_v55_ = nw289_
                def iife119_():
                    coll53_ = _dafny.Map()
                    compr_53_: int
                    for compr_53_ in _dafny.IntegerRange(228, -272):
                        d_1589_v56_: int = compr_53_
                        if ((228) <= (d_1589_v56_)) and ((d_1589_v56_) < (-272)):
                            coll53_[(d_1589_v56_) + (p1)] = (d_1579_v48_).f40
                    return _dafny.Map(coll53_)
                rhs297_ = d_1586_v55_
                rhs298_ = not(d_1567_cf61_)
                rhs299_ = len((iife119_()
                ) | ((default__.fm41(p0, (self).f29, d_1567_cf61_, (d_1512_v2_).f33, globalState)) | (d_1518_v8_)))
                lhs279_ = globalState
                lhs280_ = globalState
                r1 = rhs297_
                lhs279_.f13 = rhs298_
                lhs280_.f2 = rhs299_
                (globalState).f15 = d_1578_cf19_
            elif source23_.is_DC10:
                d_1590___mcc_h10_ = source23_.cf18
                d_1591_cf18_ = d_1590___mcc_h10_
                d_1592_v57_: _dafny.MultiSet
                d_1592_v57_ = _dafny.MultiSet([d_1514_v4_])
                (globalState).f21 = (((d_1592_v57_).intersection(d_1592_v57_)) - (_dafny.MultiSet([d_1514_v4_]))).cardinality
                rhs300_ = ((0) - ((d_1512_v2_).f34)) * ((d_1512_v2_).f34)
                rhs301_ = (self.f30) + (self.f30)
                rhs302_ = (d_1512_v2_).fm5(_dafny.Set({(self).f31, d_1567_cf61_}), globalState)
                rhs303_ = default__.fm33(((d_1512_v2_).f34) <= ((d_1512_v2_).f34), p1, (d_1512_v2_).f34, globalState)
                rhs304_ = -554
                lhs281_ = globalState
                lhs282_ = globalState
                lhs283_ = globalState
                lhs284_ = globalState
                lhs285_ = globalState
                lhs281_.f2 = rhs300_
                lhs282_.f11 = rhs301_
                lhs283_.f21 = rhs302_
                lhs284_.f28 = rhs303_
                lhs285_.f21 = rhs304_
                rhs305_ = d_1509_v0_
                rhs306_ = (self.f30) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uef")))
                lhs286_ = globalState
                d_1509_v0_ = rhs305_
                lhs286_.f11 = rhs306_
                d_1593_v58_: _dafny.Array
                nw290_ = _dafny.Array(None, 28)
                d_1593_v58_ = nw290_
                d_1594_v59_: _dafny.Array
                nw291_ = _dafny.Array(int(0), 5)
                d_1594_v59_ = nw291_
                d_1595_v60_: _dafny.Set
                d_1595_v60_ = _dafny.Set({d_1594_v59_})
                d_1596_v61_: C2
                nw292_ = C2()
                nw292_.ctor__(len(_dafny.Map({d_1570_v44_: (self).f31})), d_1595_v60_, (self).f31, self.f30)
                d_1596_v61_ = nw292_
                index259_ = default__.safeIndex(901, (d_1593_v58_).length(0))
                (d_1593_v58_)[index259_] = d_1596_v61_
                index260_ = default__.safeIndex(901, (d_1593_v58_).length(0))
                (d_1593_v58_)[index260_] = d_1596_v61_
            elif True:
                d_1597___mcc_h11_ = source23_.cf21
                d_1598_cf21_ = d_1597___mcc_h11_
                d_1599_v62_: _dafny.Array
                def lambda76_(d_1600_v2_):
                    def lambda77_(d_1601_i8_):
                        return _dafny.Set({146, (d_1600_v2_).f34})

                    return lambda77_

                init43_ = lambda76_(d_1512_v2_)
                nw293_ = _dafny.Array(None, 28)
                for i0_43_ in range(nw293_.length(0)):
                    nw293_[i0_43_] = init43_(i0_43_)
                d_1599_v62_ = nw293_
                index261_ = default__.safeIndex(412, (d_1599_v62_).length(0))
                (d_1599_v62_)[index261_] = d_1571_v45_
                index262_ = default__.safeIndex(412, (d_1599_v62_).length(0))
                (d_1599_v62_)[index262_] = (d_1571_v45_ if (d_1512_v2_).f33 else (_dafny.Set({len(self.f30), (d_1512_v2_).f34})).intersection(d_1571_v45_))
                d_1602_v63_: T0
                nw294_ = C1()
                nw294_.ctor__(d_1567_cf61_)
                d_1602_v63_ = nw294_
                d_1603_v64_: _dafny.Seq
                d_1603_v64_ = _dafny.SeqWithoutIsStrInference([d_1602_v63_])
                d_1604_v65_: _dafny.Array
                nw295_ = _dafny.Array(None, 11)
                nw295_[int(0)] = d_1602_v63_
                nw295_[int(1)] = d_1602_v63_
                nw295_[int(2)] = (d_1603_v64_)[default__.safeIndex(p1, len(d_1603_v64_))]
                nw295_[int(3)] = d_1602_v63_
                nw295_[int(4)] = d_1602_v63_
                nw295_[int(5)] = d_1602_v63_
                nw295_[int(6)] = d_1602_v63_
                nw295_[int(7)] = (d_1602_v63_ if (self).f29 else d_1602_v63_)
                nw295_[int(8)] = d_1602_v63_
                nw295_[int(9)] = d_1602_v63_
                nw295_[int(10)] = d_1602_v63_
                d_1604_v65_ = nw295_
                index263_ = default__.safeIndex(472, (d_1604_v65_).length(0))
                (d_1604_v65_)[index263_] = d_1602_v63_
                index264_ = default__.safeIndex(667, (d_1509_v0_).length(0))
                (d_1509_v0_)[index264_] = default__.fm19(d_1514_v4_, globalState)
                d_1605_v66_: _dafny.Array
                def lambda78_(d_1606_i9_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agpxc"))

                init44_ = lambda78_
                nw296_ = _dafny.Array(None, 29)
                for i0_44_ in range(nw296_.length(0)):
                    nw296_[i0_44_] = init44_(i0_44_)
                d_1605_v66_ = nw296_
                index265_ = default__.safeIndex(67, (d_1605_v66_).length(0))
                (d_1605_v66_)[index265_] = _dafny.SeqWithoutIsStrInference([d_1514_v4_ for d_1607_i10_ in range(default__.abs(870))])
                index266_ = default__.safeIndex(472, (d_1604_v65_).length(0))
                index267_ = default__.safeIndex(667, (d_1509_v0_).length(0))
                index268_ = default__.safeIndex(67, (d_1605_v66_).length(0))
                rhs307_ = (d_1512_v2_).f33
                rhs308_ = d_1602_v63_
                rhs309_ = (d_1512_v2_).f33
                rhs310_ = -943
                rhs311_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrfecmi"))
                lhs287_ = globalState
                lhs288_ = d_1604_v65_
                lhs289_ = default__.safeIndex(472, (d_1604_v65_).length(0))
                lhs290_ = d_1509_v0_
                lhs291_ = default__.safeIndex(667, (d_1509_v0_).length(0))
                lhs292_ = globalState
                lhs293_ = d_1605_v66_
                lhs294_ = default__.safeIndex(67, (d_1605_v66_).length(0))
                lhs287_.f18 = rhs307_
                lhs288_[lhs289_] = rhs308_
                lhs290_[lhs291_] = rhs309_
                lhs292_.f21 = rhs310_
                lhs293_[lhs294_] = rhs311_
                d_1608_v67_: _dafny.Map
                d_1608_v67_ = _dafny.Map({self.f30: (d_1509_v0_)[default__.safeIndex(667, (d_1509_v0_).length(0))]})
                d_1609_v68_: _dafny.Seq
                d_1609_v68_ = _dafny.SeqWithoutIsStrInference([p0, len(default__.fm7((0) - ((d_1512_v2_).f34), globalState)), (d_1512_v2_).f34, (d_1512_v2_).f34, (d_1517_v7_)[default__.safeIndex(139, (d_1517_v7_).length(0))]])
                rhs312_ = (d_1609_v68_ if (self).f32 else _dafny.SeqWithoutIsStrInference([(d_1512_v2_).f34]))
                rhs313_ = (d_1609_v68_) + (d_1609_v68_)
                rhs314_ = _dafny.Map({self.f30: (d_1509_v0_)[default__.safeIndex(667, (d_1509_v0_).length(0))]})
                lhs295_ = globalState
                lhs296_ = globalState
                lhs295_.f7 = rhs312_
                lhs296_.f7 = rhs313_
                d_1608_v67_ = rhs314_
                d_1610_v69_: C0
                nw297_ = C0()
                nw297_.ctor__(((d_1599_v62_)[default__.safeIndex(412, (d_1599_v62_).length(0))]).isdisjoint((d_1599_v62_)[default__.safeIndex(412, (d_1599_v62_).length(0))]))
                d_1610_v69_ = nw297_
                nw298_ = C0()
                nw298_.ctor__((d_1521_v11_) == (d_1521_v11_))
                d_1610_v69_ = nw298_
            (globalState).f28 = (-874) in (d_1518_v8_)
        elif source22_.is_DC37:
            d_1611___mcc_h5_ = source22_.cf62
            d_1612___mcc_h6_ = source22_.cf63
            d_1613_cf63_ = d_1612___mcc_h6_
            d_1614_cf62_ = d_1611___mcc_h5_
            if ((d_1613_cf63_ if (d_1512_v2_).f33 else False) if (-359) <= (p0) else ((d_1512_v2_).f34) >= (p1)):
                d_1615_v70_: _dafny.Array
                def lambda79_(d_1616_v4_):
                    def lambda80_(d_1617_i11_):
                        return _dafny.SeqWithoutIsStrInference([d_1616_v4_ for d_1618_i12_ in range(default__.abs(550))])

                    return lambda80_

                init45_ = lambda79_(d_1514_v4_)
                nw299_ = _dafny.Array(None, 6)
                for i0_45_ in range(nw299_.length(0)):
                    nw299_[i0_45_] = init45_(i0_45_)
                d_1615_v70_ = nw299_
                index269_ = default__.safeIndex(274, (d_1615_v70_).length(0))
                (d_1615_v70_)[index269_] = _dafny.SeqWithoutIsStrInference([d_1514_v4_ for d_1619_i13_ in range(default__.abs(-415))])
                index270_ = default__.safeIndex(274, (d_1615_v70_).length(0))
                (d_1615_v70_)[index270_] = (self.f30) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1620_i14_ in range(default__.abs(205))]))
                d_1621_v71_: _dafny.Map
                d_1621_v71_ = _dafny.Map({(d_1512_v2_).f33: d_1514_v4_})
                d_1622_v72_: _dafny.Seq
                d_1622_v72_ = _dafny.SeqWithoutIsStrInference([d_1621_v71_])
                d_1623_v73_: D2
                d_1623_v73_ = D2_DC8(d_1514_v4_, (self).f31)
                (globalState).f18 = ((d_1518_v8_)[d_1614_cf62_] if (d_1614_cf62_) in (d_1518_v8_) else (len(((d_1622_v72_)[default__.safeIndex(p0, len(d_1622_v72_))]).set(d_1613_cf63_, (d_1623_v73_).cf12))) > (-222))
                rhs315_ = ((self.f30) + ((d_1615_v70_)[default__.safeIndex(274, (d_1615_v70_).length(0))])) + (_dafny.SeqWithoutIsStrInference([d_1514_v4_ for d_1624_i15_ in range(default__.abs(795))]))
                rhs316_ = default__.safeDivisionInt(len(_dafny.Set({len(_dafny.Map({(d_1512_v2_).f33: (0) - ((d_1512_v2_).f34)})), (d_1512_v2_).f34})), (d_1512_v2_).f34)
                lhs297_ = globalState
                lhs298_ = globalState
                lhs297_.f20 = rhs315_
                lhs298_.f1 = rhs316_
                d_1625_v74_: _dafny.Seq
                d_1625_v74_ = _dafny.SeqWithoutIsStrInference([(d_1615_v70_)[default__.safeIndex(274, (d_1615_v70_).length(0))], self.f30, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbew")), (d_1615_v70_)[default__.safeIndex(274, (d_1615_v70_).length(0))], self.f30])
                d_1518_v8_ = (d_1518_v8_).set(len((d_1625_v74_) + (d_1625_v74_)), (self).f32)
                d_1626_v75_: _dafny.Seq
                d_1626_v75_ = _dafny.SeqWithoutIsStrInference([d_1613_cf63_, (self).f31])
                (globalState).f28 = (d_1626_v75_) == ((d_1626_v75_) + (_dafny.SeqWithoutIsStrInference([(self).f32])))
            elif True:
                d_1614_cf62_ = ((d_1521_v11_)[d_1614_cf62_] if (d_1614_cf62_) in (d_1521_v11_) else p0)
                d_1627_v76_: C5
                nw300_ = C5()
                nw300_.ctor__(d_1613_cf63_, self.f30)
                d_1627_v76_ = nw300_
                (globalState).f2 = p1
                r2 = p1
                d_1628_v77_: _dafny.Map
                d_1628_v77_ = _dafny.Map({d_1613_cf63_: self.f30})
                d_1614_cf62_ = len(((default__.fm52(globalState)).set((self).f29, self.f30)) | (d_1628_v77_))
            d_1514_v4_ = _dafny.CodePoint('c')
            index271_ = default__.safeIndex(176, (d_1509_v0_).length(0))
            (d_1509_v0_)[index271_] = (self.f30) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejv")))
            index272_ = default__.safeIndex(176, (d_1509_v0_).length(0))
            (d_1509_v0_)[index272_] = default__.fm19(d_1514_v4_, globalState)
            d_1629_v78_: D11
            d_1629_v78_ = D11_DC35(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqnwtssk")), (d_1512_v2_).f33)
            d_1630_v79_: _dafny.MultiSet
            d_1630_v79_ = _dafny.MultiSet([not((d_1629_v78_).cf58)])
            d_1631_v80_: _dafny.MultiSet
            d_1631_v80_ = _dafny.MultiSet([d_1514_v4_, d_1514_v4_])
            d_1632_v81_: _dafny.Seq
            d_1632_v81_ = _dafny.SeqWithoutIsStrInference([d_1613_cf63_])
            d_1633_v82_: _dafny.Seq
            d_1633_v82_ = _dafny.SeqWithoutIsStrInference([d_1632_v81_, d_1632_v81_, d_1632_v81_, d_1632_v81_])
            d_1630_v79_ = (_dafny.MultiSet([(self).f31, (d_1512_v2_).f33])) | (((d_1630_v79_).set(d_1613_cf63_, default__.abs((d_1631_v80_).cardinality))) | (_dafny.MultiSet((d_1633_v82_)[default__.safeIndex((d_1512_v2_).f34, len(d_1633_v82_))])))
        elif True:
            d_1634___mcc_h7_ = source22_.cf56
            d_1635_cf56_ = d_1634___mcc_h7_
            d_1636_v83_: _dafny.Seq
            d_1636_v83_ = _dafny.SeqWithoutIsStrInference([(d_1512_v2_).f34])
            d_1637_v84_: D9
            d_1637_v84_ = D9_DC27(d_1636_v83_)
            d_1638_v85_: D2
            d_1638_v85_ = D2_DC9(self.f30, len((d_1637_v84_).cf40), (d_1512_v2_).f33, (d_1512_v2_).f34)
            source24_ = d_1638_v85_
            if source24_.is_DC7:
                d_1639___mcc_h12_ = source24_.cf11
                d_1640_cf11_ = d_1639___mcc_h12_
                d_1641_v86_: _dafny.Map
                d_1641_v86_ = _dafny.Map({default__.safeDivisionInt((d_1512_v2_).f34, (d_1512_v2_).f34): self.f30})
                d_1641_v86_ = (d_1641_v86_).set(p1, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1642_i16_ in range(default__.abs(663))]))
                (globalState).f13 = d_1640_cf11_
                (globalState).f1 = (0) - (p1)
                d_1643_v87_: _dafny.Seq
                d_1643_v87_ = _dafny.SeqWithoutIsStrInference([(self).f32, (self).f32, (p1) >= ((d_1512_v2_).f34)])
                (globalState).f13 = (d_1643_v87_)[default__.safeIndex(len((self.f30) + (self.f30)), len(d_1643_v87_))]
            elif source24_.is_DC8:
                d_1644___mcc_h13_ = source24_.cf12
                d_1645___mcc_h14_ = source24_.cf13
                d_1646_cf13_ = d_1645___mcc_h14_
                d_1647_cf12_ = d_1644___mcc_h13_
                d_1648_v88_: _dafny.MultiSet
                d_1648_v88_ = _dafny.MultiSet([d_1646_cf13_, (self).f31])
                d_1649_v89_: _dafny.Seq
                d_1649_v89_ = _dafny.SeqWithoutIsStrInference([(self).f32, (d_1512_v2_).f33, False])
                d_1648_v88_ = _dafny.MultiSet((d_1649_v89_) + (d_1649_v89_))
                (globalState).f2 = (d_1512_v2_).f34
                rhs317_ = ((p0) >= (p1) if d_1646_cf13_ else (self).f32)
                rhs318_ = d_1646_cf13_
                rhs319_ = (d_1512_v2_).f33
                rhs320_ = (d_1512_v2_).f34
                lhs299_ = globalState
                lhs300_ = globalState
                lhs301_ = globalState
                lhs302_ = globalState
                lhs299_.f18 = rhs317_
                lhs300_.f18 = rhs318_
                lhs301_.f13 = rhs319_
                lhs302_.f2 = rhs320_
                d_1650_v90_: _dafny.Set
                d_1650_v90_ = _dafny.Set({(d_1512_v2_).f33})
                d_1651_v91_: D6
                d_1651_v91_ = D6_DC21(d_1650_v90_)
                d_1652_v92_: _dafny.Seq
                d_1652_v92_ = _dafny.SeqWithoutIsStrInference([d_1651_v91_])
                d_1653_v93_: D8
                d_1653_v93_ = D8_DC25(d_1652_v92_)
                d_1654_v94_: D8
                d_1654_v94_ = D8_DC25((d_1653_v93_).cf39)
                d_1655_v95_: _dafny.Array
                nw301_ = _dafny.Array(None, 20)
                nw301_[int(0)] = d_1653_v93_
                nw301_[int(1)] = d_1654_v94_
                nw301_[int(2)] = d_1654_v94_
                nw301_[int(3)] = d_1653_v93_
                nw301_[int(4)] = d_1654_v94_
                nw301_[int(5)] = d_1654_v94_
                nw301_[int(6)] = D8_DC25(d_1652_v92_)
                nw301_[int(7)] = d_1654_v94_
                nw301_[int(8)] = D8_DC25(_dafny.SeqWithoutIsStrInference([d_1651_v91_, D6_DC21(d_1650_v90_)]))
                nw301_[int(9)] = d_1653_v93_
                nw301_[int(10)] = d_1654_v94_
                nw301_[int(11)] = d_1653_v93_
                nw301_[int(12)] = d_1653_v93_
                nw301_[int(13)] = d_1653_v93_
                nw301_[int(14)] = d_1653_v93_
                nw301_[int(15)] = d_1654_v94_
                nw301_[int(16)] = d_1654_v94_
                nw301_[int(17)] = d_1654_v94_
                nw301_[int(18)] = d_1654_v94_
                nw301_[int(19)] = d_1653_v93_
                d_1655_v95_ = nw301_
                d_1656_v96_: _dafny.Map
                d_1656_v96_ = _dafny.Map({d_1655_v95_: (self).f29})
                rhs321_ = len(d_1656_v96_)
                rhs322_ = p0
                lhs303_ = globalState
                lhs304_ = globalState
                lhs303_.f19 = rhs321_
                lhs304_.f21 = rhs322_
            elif source24_.is_DC9:
                d_1657___mcc_h15_ = source24_.cf14
                d_1658___mcc_h16_ = source24_.cf15
                d_1659___mcc_h17_ = source24_.cf16
                d_1660___mcc_h18_ = source24_.cf17
                d_1661_cf17_ = d_1660___mcc_h18_
                d_1662_cf16_ = d_1659___mcc_h17_
                d_1663_cf15_ = d_1658___mcc_h16_
                d_1664_cf14_ = d_1657___mcc_h15_
                d_1665_v97_: C5
                nw302_ = C5()
                nw302_.ctor__((self).f31, self.f30)
                d_1665_v97_ = nw302_
                (globalState).f13 = not(((d_1512_v2_).f33) == ((self).f29))
                rhs323_ = d_1517_v7_
                rhs324_ = d_1661_cf17_
                rhs325_ = (d_1512_v2_).f33
                lhs305_ = globalState
                lhs306_ = globalState
                lhs305_.f17 = rhs323_
                r2 = rhs324_
                lhs306_.f18 = rhs325_
                d_1666_v98_: _dafny.Array
                nw303_ = _dafny.Array(_dafny.Seq({}), 27)
                d_1666_v98_ = nw303_
                d_1666_v98_ = d_1666_v98_
            elif True:
                d_1667___mcc_h19_ = source24_.cf10
                d_1668_cf10_ = d_1667___mcc_h19_
                d_1669_v99_: _dafny.Map
                d_1669_v99_ = _dafny.Map({((d_1512_v2_).f34) >= ((d_1512_v2_).f34): p0})
                d_1670_v100_: _dafny.Seq
                d_1670_v100_ = _dafny.SeqWithoutIsStrInference([d_1509_v0_, d_1509_v0_, d_1509_v0_])
                d_1671_v101_: D11
                d_1671_v101_ = D11_DC34(default__.fm45(globalState))
                d_1672_v102_: _dafny.Seq
                d_1672_v102_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f32: p1}), d_1669_v99_, d_1669_v99_, d_1669_v99_])
                d_1673_v103_: _dafny.Map
                d_1673_v103_ = _dafny.Map({(d_1512_v2_).f34: p0})
                pat_let_tv34_ = d_1635_cf56_
                def iife120_(_pat_let33_0):
                    def iife121_(d_1674_dt__update__tmp_h0_):
                        def iife122_(_pat_let34_0):
                            def iife123_(d_1675_dt__update_hcf56_h0_):
                                return D11_DC34(d_1675_dt__update_hcf56_h0_)
                            return iife123_(_pat_let34_0)
                        return iife122_(pat_let_tv34_)
                    return iife121_(_pat_let33_0)
                rhs326_ = (d_1672_v102_)[default__.safeIndex((len(d_1673_v103_)) + ((0) - ((d_1512_v2_).f34)), len(d_1672_v102_))]
                rhs327_ = d_1670_v100_
                rhs328_ = iife120_(d_1671_v101_)
                d_1669_v99_ = rhs326_
                d_1670_v100_ = rhs327_
                d_1671_v101_ = rhs328_
                d_1676_v104_: C1
                nw304_ = C1()
                nw304_.ctor__((d_1512_v2_).f33)
                d_1676_v104_ = nw304_
                (globalState).f2 = (d_1512_v2_).f34
                d_1677_v105_: D3
                d_1677_v105_ = D3_DC11((self).f31, (self).f31)
                d_1678_v106_: D3
                d_1678_v106_ = D3_DC12(d_1677_v105_)
                d_1679_v107_: D3
                d_1679_v107_ = D3_DC12((d_1678_v106_).cf21)
                d_1680_v108_: _dafny.Map
                d_1680_v108_ = _dafny.Map({(d_1676_v104_).f40: d_1679_v107_})
                index273_ = default__.safeIndex(548, (d_1509_v0_).length(0))
                (d_1509_v0_)[index273_] = True
                d_1681_v109_: _dafny.Array
                nw305_ = _dafny.Array(_dafny.Seq({}), 15)
                d_1681_v109_ = nw305_
                d_1682_v110_: _dafny.Seq
                d_1682_v110_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_1512_v2_).f34])])
                index274_ = default__.safeIndex(233, (d_1681_v109_).length(0))
                (d_1681_v109_)[index274_] = d_1682_v110_
                d_1683_v111_: _dafny.Set
                d_1683_v111_ = _dafny.Set({(d_1512_v2_).f33})
                d_1684_v112_: _dafny.Array
                nw306_ = _dafny.Array(None, 7)
                nw306_[int(0)] = d_1683_v111_
                nw306_[int(1)] = (d_1683_v111_) - (d_1683_v111_)
                nw306_[int(2)] = (d_1683_v111_ if (d_1676_v104_).f40 else default__.fm22(default__.fm36(globalState), default__.fm33((self).f29, (d_1512_v2_).f34, p1, globalState), d_1514_v4_, globalState))
                nw306_[int(3)] = (d_1683_v111_) | (d_1683_v111_)
                nw306_[int(4)] = d_1683_v111_
                nw306_[int(5)] = d_1683_v111_
                nw306_[int(6)] = default__.fm22((d_1512_v2_).f34, (self).f31, d_1514_v4_, globalState)
                d_1684_v112_ = nw306_
                index275_ = default__.safeIndex(499, (d_1684_v112_).length(0))
                (d_1684_v112_)[index275_] = d_1683_v111_
                index276_ = default__.safeIndex(548, (d_1509_v0_).length(0))
                index277_ = default__.safeIndex(233, (d_1681_v109_).length(0))
                index278_ = default__.safeIndex(499, (d_1684_v112_).length(0))
                rhs329_ = default__.fm53(not(default__.fm6((d_1512_v2_).f34, globalState)), (d_1512_v2_).f34, self.f30, p0, globalState)
                rhs330_ = (d_1512_v2_).f33
                rhs331_ = d_1682_v110_
                rhs332_ = 384
                rhs333_ = default__.fm22((len(d_1683_v111_) if (self).f31 else (d_1512_v2_).f34), (d_1512_v2_).f33, d_1514_v4_, globalState)
                lhs307_ = d_1509_v0_
                lhs308_ = default__.safeIndex(548, (d_1509_v0_).length(0))
                lhs309_ = d_1681_v109_
                lhs310_ = default__.safeIndex(233, (d_1681_v109_).length(0))
                lhs311_ = globalState
                lhs312_ = d_1684_v112_
                lhs313_ = default__.safeIndex(499, (d_1684_v112_).length(0))
                d_1680_v108_ = rhs329_
                lhs307_[lhs308_] = rhs330_
                lhs309_[lhs310_] = rhs331_
                lhs311_.f2 = rhs332_
                lhs312_[lhs313_] = rhs333_
            index279_ = default__.safeIndex(67, (d_1509_v0_).length(0))
            (d_1509_v0_)[index279_] = (d_1512_v2_).f33
            d_1685_v113_: _dafny.Seq
            d_1685_v113_ = _dafny.SeqWithoutIsStrInference([(self).f29])
            index280_ = default__.safeIndex(67, (d_1509_v0_).length(0))
            (d_1509_v0_)[index280_] = ((d_1512_v2_).f33) and ((not((self).f29)) in (d_1685_v113_))
            (globalState).f15 = (D17_DC52((d_1509_v0_)[default__.safeIndex(67, (d_1509_v0_).length(0))], (d_1512_v2_).f33, (d_1512_v2_).f34, d_1521_v11_)).cf89
            r2 = (d_1512_v2_).fm5(_dafny.Set({(self).f29}), globalState)
        d_1686_v114_: _dafny.Array
        nw307_ = _dafny.Array(_dafny.Map({}), 19)
        d_1686_v114_ = nw307_
        d_1687_v115_: _dafny.Map
        d_1687_v115_ = _dafny.Map({(self).f29: d_1509_v0_})
        index281_ = default__.safeIndex(315, (d_1686_v114_).length(0))
        (d_1686_v114_)[index281_] = (_dafny.Map({(self).f31: d_1509_v0_})) | (d_1687_v115_)
        d_1688_v116_: _dafny.Set
        d_1688_v116_ = _dafny.Set({self.f30})
        d_1689_v117_: _dafny.Seq
        d_1689_v117_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjcvfdu")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lnab"))])
        index282_ = default__.safeIndex(315, (d_1686_v114_).length(0))
        rhs334_ = (d_1512_v2_).fm5(_dafny.Set({False}), globalState)
        rhs335_ = (d_1512_v2_).f33
        rhs336_ = d_1687_v115_
        rhs337_ = ((self.f30) + ((d_1689_v117_)[default__.safeIndex(-811, len(d_1689_v117_))])) + ((self.f30) + (self.f30))
        rhs338_ = d_1688_v116_
        lhs314_ = globalState
        lhs315_ = globalState
        lhs316_ = d_1686_v114_
        lhs317_ = default__.safeIndex(315, (d_1686_v114_).length(0))
        lhs318_ = self
        lhs314_.f19 = rhs334_
        lhs315_.f18 = rhs335_
        lhs316_[lhs317_] = rhs336_
        lhs318_.f30 = rhs337_
        d_1688_v116_ = rhs338_
        if (self).f29:
            d_1690_v118_: _dafny.MultiSet
            d_1690_v118_ = _dafny.MultiSet([d_1509_v0_])
            if ((d_1690_v118_) | (d_1690_v118_)) != ((d_1690_v118_).set(d_1509_v0_, default__.abs(len(_dafny.Map({d_1514_v4_: (self).f32}))))):
                (globalState).f20 = self.f30
                d_1691_v119_: _dafny.Array
                nw308_ = _dafny.Array(_dafny.Map({}), 14)
                d_1691_v119_ = nw308_
                d_1692_v120_: _dafny.Array
                def lambda81_(d_1693_i17_):
                    return self.f30

                init46_ = lambda81_
                nw309_ = _dafny.Array(None, 16)
                for i0_46_ in range(nw309_.length(0)):
                    nw309_[i0_46_] = init46_(i0_46_)
                d_1692_v120_ = nw309_
                d_1694_v121_: _dafny.Seq
                d_1694_v121_ = _dafny.SeqWithoutIsStrInference([p1, p0])
                d_1695_v122_: _dafny.Map
                d_1695_v122_ = _dafny.Map({d_1692_v120_: _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1514_v4_ for d_1696_i18_ in range(default__.abs(360))]): len(d_1694_v121_)})})
                d_1697_v123_: D15
                d_1697_v123_ = D15_DC45((d_1512_v2_).f34, p1, (d_1512_v2_).f33, (self).f31, d_1692_v120_)
                d_1698_v125_: _dafny.Map
                d_1698_v125_ = _dafny.Map({self.f30: len(self.f30)})
                index283_ = default__.safeIndex(159, (d_1691_v119_).length(0))
                def iife124_():
                    coll54_ = _dafny.Map()
                    compr_54_: _dafny.Seq
                    for compr_54_ in (d_1698_v125_).keys.Elements:
                        d_1699_v124_: _dafny.Seq = compr_54_
                        if (d_1699_v124_) in (d_1698_v125_):
                            coll54_[d_1699_v124_] = p0
                    return _dafny.Map(coll54_)
                (d_1691_v119_)[index283_] = ((d_1695_v122_)[(d_1697_v123_).cf80] if ((d_1697_v123_).cf80) in (d_1695_v122_) else iife124_()
                )
                index284_ = default__.safeIndex(159, (d_1691_v119_).length(0))
                (d_1691_v119_)[index284_] = d_1698_v125_
                (globalState).f28 = (d_1512_v2_).f33
                d_1700_v126_: _dafny.Map
                d_1700_v126_ = _dafny.Map({len(_dafny.Map({(d_1512_v2_).f33: (d_1512_v2_).f33})): d_1509_v0_})
                d_1701_v127_: C4
                nw310_ = C4()
                nw310_.ctor__(default__.fm33((self).f32, (d_1512_v2_).f34, p0, globalState), len((d_1700_v126_) | (d_1700_v126_)))
                d_1701_v127_ = nw310_
                d_1509_v0_ = d_1509_v0_
            elif True:
                d_1702_v128_: _dafny.Map
                d_1702_v128_ = _dafny.Map({(self).f29: (self).f29})
                d_1703_v129_: _dafny.Seq
                d_1703_v129_ = _dafny.SeqWithoutIsStrInference([(self).f31])
                d_1704_v130_: T1
                nw311_ = C5()
                nw311_.ctor__(False, default__.fm7(p0, globalState))
                d_1704_v130_ = nw311_
                d_1705_v131_: D17
                d_1705_v131_ = D17_DC50(((d_1702_v128_)[(d_1703_v129_)[default__.safeIndex((d_1512_v2_).f34, len(d_1703_v129_))]] if ((d_1703_v129_)[default__.safeIndex((d_1512_v2_).f34, len(d_1703_v129_))]) in (d_1702_v128_) else (self).f31), True, d_1704_v130_)
                d_1705_v131_ = D17_DC50((d_1512_v2_).f33, (True) and ((d_1512_v2_).f33), d_1704_v130_)
                d_1703_v129_ = d_1703_v129_
                (globalState).f13 = True
                d_1706_v133_: _dafny.Map
                d_1706_v133_ = _dafny.Map({(d_1512_v2_).f34: p1})
                d_1707_v134_: _dafny.Seq
                d_1707_v134_ = _dafny.SeqWithoutIsStrInference([(d_1512_v2_).f34, (0) - ((d_1512_v2_).f34), default__.fm36(globalState), ((d_1706_v133_)[p0] if (p0) in (d_1706_v133_) else (d_1512_v2_).f34)])
                def iife125_():
                    coll55_ = _dafny.Map()
                    compr_55_: int
                    for compr_55_ in (d_1707_v134_).Elements:
                        d_1708_v132_: int = compr_55_
                        if (d_1708_v132_) in (d_1707_v134_):
                            coll55_[default__.safeDivisionInt(d_1708_v132_, p1)] = (self).f29
                    return _dafny.Map(coll55_)
                (globalState).f15 = (d_1518_v8_) != ((iife125_()
                ) | (d_1518_v8_))
                d_1709_v135_: _dafny.Array
                nw312_ = _dafny.Array(None, 27)
                nw312_[int(0)] = d_1517_v7_
                nw312_[int(1)] = (d_1519_v9_).cf74
                nw312_[int(2)] = d_1517_v7_
                nw312_[int(3)] = d_1517_v7_
                nw312_[int(4)] = d_1517_v7_
                nw312_[int(5)] = (d_1517_v7_ if (d_1512_v2_).f33 else d_1517_v7_)
                nw312_[int(6)] = d_1517_v7_
                nw312_[int(7)] = d_1517_v7_
                nw312_[int(8)] = d_1517_v7_
                nw312_[int(9)] = d_1517_v7_
                nw312_[int(10)] = d_1517_v7_
                nw312_[int(11)] = d_1517_v7_
                nw312_[int(12)] = d_1517_v7_
                nw312_[int(13)] = d_1517_v7_
                nw312_[int(14)] = d_1517_v7_
                nw312_[int(15)] = d_1517_v7_
                nw312_[int(16)] = d_1517_v7_
                nw312_[int(17)] = d_1517_v7_
                nw312_[int(18)] = d_1517_v7_
                nw312_[int(19)] = d_1517_v7_
                nw312_[int(20)] = d_1517_v7_
                nw312_[int(21)] = d_1517_v7_
                nw312_[int(22)] = d_1517_v7_
                nw312_[int(23)] = d_1517_v7_
                nw312_[int(24)] = d_1517_v7_
                nw312_[int(25)] = d_1517_v7_
                nw312_[int(26)] = d_1517_v7_
                d_1709_v135_ = nw312_
                index285_ = default__.safeIndex(634, (d_1709_v135_).length(0))
                (d_1709_v135_)[index285_] = d_1517_v7_
                index286_ = default__.safeIndex(634, (d_1709_v135_).length(0))
                (d_1709_v135_)[index286_] = (d_1517_v7_ if (d_1512_v2_).f33 else (d_1517_v7_ if (d_1512_v2_).f33 else d_1517_v7_))
            d_1710_v136_: _dafny.Map
            d_1710_v136_ = _dafny.Map({d_1517_v7_: (d_1512_v2_).f33})
            d_1711_v137_: _dafny.Seq
            d_1711_v137_ = _dafny.SeqWithoutIsStrInference([d_1517_v7_])
            d_1712_v138_: _dafny.Map
            d_1712_v138_ = _dafny.Map({(self).f29: (d_1512_v2_).f34})
            d_1710_v136_ = (d_1710_v136_).set((d_1711_v137_)[default__.safeIndex(((d_1712_v138_)[(self).f32] if ((self).f32) in (d_1712_v138_) else p0), len(d_1711_v137_))], (True) or (not((self).f32)))
            (globalState).f11 = self.f30
            d_1713_v139_: _dafny.Set
            d_1713_v139_ = _dafny.Set({(self).f29, (self).f29, (self).f31, (self).f31, (d_1512_v2_).f33})
            index287_ = default__.safeIndex(200, (d_1517_v7_).length(0))
            (d_1517_v7_)[index287_] = (d_1512_v2_).fm5(d_1713_v139_, globalState)
            d_1714_v140_: _dafny.Seq
            d_1714_v140_ = _dafny.SeqWithoutIsStrInference([default__.fm9(globalState)])
            index288_ = default__.safeIndex(200, (d_1517_v7_).length(0))
            rhs339_ = (self).f32
            rhs340_ = (len(d_1714_v140_) if (self).f32 else (d_1512_v2_).f34)
            lhs319_ = globalState
            lhs320_ = d_1517_v7_
            lhs321_ = default__.safeIndex(200, (d_1517_v7_).length(0))
            lhs319_.f28 = rhs339_
            lhs320_[lhs321_] = rhs340_
            if ((d_1515_v5_).cf53 if (self).f31 else (self).f29):
                d_1518_v8_ = (d_1518_v8_).set(p0, (d_1512_v2_).f33)
                (globalState).f18 = (d_1512_v2_).f33
                d_1715_v141_: _dafny.Array
                nw313_ = _dafny.Array(int(0), 25)
                d_1715_v141_ = nw313_
                d_1716_v142_: D17
                d_1716_v142_ = D17_DC51(len(self.f30), d_1715_v141_)
                d_1717_v143_: _dafny.Array
                nw314_ = _dafny.Array(None, 13)
                nw314_[int(0)] = d_1715_v141_
                nw314_[int(1)] = (d_1716_v142_).cf87
                nw314_[int(2)] = d_1715_v141_
                nw314_[int(3)] = d_1715_v141_
                nw314_[int(4)] = d_1715_v141_
                nw314_[int(5)] = (d_1716_v142_).cf87
                nw314_[int(6)] = d_1715_v141_
                nw314_[int(7)] = d_1715_v141_
                nw314_[int(8)] = d_1715_v141_
                nw314_[int(9)] = d_1715_v141_
                nw314_[int(10)] = d_1715_v141_
                nw314_[int(11)] = d_1715_v141_
                nw314_[int(12)] = d_1715_v141_
                d_1717_v143_ = nw314_
                d_1718_v144_: _dafny.Seq
                d_1718_v144_ = _dafny.SeqWithoutIsStrInference([d_1717_v143_])
                d_1717_v143_ = (d_1718_v144_)[default__.safeIndex((d_1512_v2_).f34, len(d_1718_v144_))]
                index289_ = default__.safeIndex(200, (d_1517_v7_).length(0))
                rhs341_ = (d_1517_v7_)[default__.safeIndex(200, (d_1517_v7_).length(0))]
                rhs342_ = not((d_1512_v2_).f33)
                rhs343_ = default__.safeDivisionInt((d_1512_v2_).f34, (p0) - ((d_1512_v2_).f34))
                lhs322_ = globalState
                lhs323_ = globalState
                lhs324_ = d_1517_v7_
                lhs325_ = default__.safeIndex(200, (d_1517_v7_).length(0))
                lhs322_.f19 = rhs341_
                lhs323_.f18 = rhs342_
                lhs324_[lhs325_] = rhs343_
                d_1719_v145_: D0
                d_1719_v145_ = D0_DC1(d_1509_v0_, (d_1512_v2_).f34, (d_1512_v2_).f34)
                d_1720_v146_: D0
                d_1720_v146_ = D0_DC1((d_1719_v145_).cf1, p0, len(self.f30))
                d_1719_v145_ = d_1720_v146_
            elif True:
                d_1721_v147_: _dafny.Seq
                d_1721_v147_ = _dafny.SeqWithoutIsStrInference([(d_1512_v2_).f33, ((d_1512_v2_).f33) == ((d_1512_v2_).f33)])
                d_1721_v147_ = d_1721_v147_
                d_1722_v148_: _dafny.Array
                def lambda82_(d_1723_v140_):
                    def lambda83_(d_1724_i19_):
                        return (d_1723_v140_) + (d_1723_v140_)

                    return lambda83_

                init47_ = lambda82_(d_1714_v140_)
                nw315_ = _dafny.Array(None, 10)
                for i0_47_ in range(nw315_.length(0)):
                    nw315_[i0_47_] = init47_(i0_47_)
                d_1722_v148_ = nw315_
                index290_ = default__.safeIndex(496, (d_1722_v148_).length(0))
                (d_1722_v148_)[index290_] = d_1714_v140_
                index291_ = default__.safeIndex(200, (d_1517_v7_).length(0))
                index292_ = default__.safeIndex(496, (d_1722_v148_).length(0))
                rhs344_ = (self).f32
                rhs345_ = p1
                rhs346_ = ((_dafny.SeqWithoutIsStrInference([p1 for d_1725_i20_ in range(default__.abs(652))])) + (d_1714_v140_) if (d_1512_v2_).f33 else (d_1714_v140_) + (d_1714_v140_))
                rhs347_ = (d_1512_v2_).f33
                lhs326_ = globalState
                lhs327_ = d_1517_v7_
                lhs328_ = default__.safeIndex(200, (d_1517_v7_).length(0))
                lhs329_ = d_1722_v148_
                lhs330_ = default__.safeIndex(496, (d_1722_v148_).length(0))
                lhs331_ = globalState
                lhs326_.f15 = rhs344_
                lhs327_[lhs328_] = rhs345_
                lhs329_[lhs330_] = rhs346_
                lhs331_.f18 = rhs347_
                (globalState).f2 = (0) - ((d_1517_v7_)[default__.safeIndex(200, (d_1517_v7_).length(0))])
                d_1726_v149_: _dafny.Array
                nw316_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                d_1726_v149_ = nw316_
                index293_ = default__.safeIndex(87, (d_1726_v149_).length(0))
                (d_1726_v149_)[index293_] = (d_1515_v5_).cf52
                index294_ = default__.safeIndex(87, (d_1726_v149_).length(0))
                (d_1726_v149_)[index294_] = d_1514_v4_
                (globalState).f20 = (self.f30) + (self.f30)
        elif True:
            d_1514_v4_ = d_1514_v4_
            (globalState).f15 = ((self).f32 if ((d_1518_v8_)[(d_1512_v2_).f34] if ((d_1512_v2_).f34) in (d_1518_v8_) else (d_1512_v2_).f33) else (d_1512_v2_).f33)
            if (d_1512_v2_).f33:
                (globalState).f2 = (d_1512_v2_).f34
                (globalState).f1 = (p0) * ((d_1512_v2_).fm5(default__.fm22(641, (self).f31, d_1514_v4_, globalState), globalState))
                (globalState).f21 = (d_1512_v2_).f34
                (globalState).f21 = ((0) - ((d_1512_v2_).f34)) * ((d_1512_v2_).f34)
                d_1727_v150_: _dafny.Array
                nw317_ = _dafny.Array(_dafny.Seq({}), 28)
                d_1727_v150_ = nw317_
                d_1728_v151_: _dafny.Seq
                d_1728_v151_ = _dafny.SeqWithoutIsStrInference([970, (d_1512_v2_).f34])
                index295_ = default__.safeIndex(609, (d_1727_v150_).length(0))
                (d_1727_v150_)[index295_] = d_1728_v151_
                index296_ = default__.safeIndex(609, (d_1727_v150_).length(0))
                rhs348_ = d_1728_v151_
                rhs349_ = ((d_1728_v151_) + (_dafny.SeqWithoutIsStrInference([len(self.f30), 965, (d_1728_v151_)[default__.safeIndex((d_1512_v2_).f34, len(d_1728_v151_))], (d_1512_v2_).f34, default__.fm36(globalState)]))) + (d_1728_v151_)
                lhs332_ = globalState
                lhs333_ = d_1727_v150_
                lhs334_ = default__.safeIndex(609, (d_1727_v150_).length(0))
                lhs332_.f7 = rhs348_
                lhs333_[lhs334_] = rhs349_
            elif True:
                d_1729_v152_: _dafny.Seq
                d_1729_v152_ = _dafny.SeqWithoutIsStrInference([(self).f31])
                r0 = ((d_1729_v152_)[default__.safeIndex(302, len(d_1729_v152_))]) == ((self).f32)
                d_1730_v153_: _dafny.Seq
                d_1730_v153_ = _dafny.SeqWithoutIsStrInference([len(self.f30)])
                d_1731_v154_: _dafny.Map
                d_1731_v154_ = _dafny.Map({((d_1512_v2_).f33) and (default__.fm33((d_1512_v2_).f33, len(d_1730_v153_), (d_1512_v2_).f34, globalState)): (self).f32})
                d_1731_v154_ = (d_1731_v154_).set((d_1512_v2_).f33, not(((d_1512_v2_).f34) < ((d_1512_v2_).f34)))
                d_1732_v155_: _dafny.Array
                def lambda84_(d_1733_v8_, d_1734_v2_):
                    def lambda85_(d_1735_i21_):
                        return _dafny.Map({len(d_1733_v8_): len(_dafny.SeqWithoutIsStrInference([(d_1734_v2_).f34 for d_1736_i22_ in range(default__.abs(194))]))})

                    return lambda85_

                init48_ = lambda84_(d_1518_v8_, d_1512_v2_)
                nw318_ = _dafny.Array(None, 2)
                for i0_48_ in range(nw318_.length(0)):
                    nw318_[i0_48_] = init48_(i0_48_)
                d_1732_v155_ = nw318_
                d_1737_v156_: _dafny.Map
                d_1737_v156_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlk"))): (d_1521_v11_).cardinality})
                d_1738_v158_: _dafny.Map
                def iife126_():
                    coll56_ = _dafny.Map()
                    compr_56_: int
                    for compr_56_ in _dafny.IntegerRange(18, 970):
                        d_1739_v157_: int = compr_56_
                        if ((18) <= (d_1739_v157_)) and ((d_1739_v157_) < (970)):
                            coll56_[(d_1739_v157_) + (p1)] = 716
                    return _dafny.Map(coll56_)
                d_1738_v158_ = (d_1737_v156_) | (iife126_()
                )
                rhs350_ = d_1732_v155_
                rhs351_ = self.f30
                rhs352_ = not((self).f29)
                rhs353_ = (d_1738_v158_ if (self).f32 else d_1738_v158_)
                rhs354_ = ((d_1512_v2_).f34 if (self).f31 else (d_1512_v2_).f34)
                lhs335_ = globalState
                lhs336_ = globalState
                lhs337_ = globalState
                d_1732_v155_ = rhs350_
                lhs335_.f11 = rhs351_
                lhs336_.f13 = rhs352_
                d_1738_v158_ = rhs353_
                lhs337_.f14 = rhs354_
                (globalState).f15 = (False if (self).f31 else default__.fm6((d_1512_v2_).f34, globalState))
                d_1740_v159_: _dafny.Array
                def lambda86_(d_1741_i23_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "suupvv"))

                init49_ = lambda86_
                nw319_ = _dafny.Array(None, 17)
                for i0_49_ in range(nw319_.length(0)):
                    nw319_[i0_49_] = init49_(i0_49_)
                d_1740_v159_ = nw319_
                index297_ = default__.safeIndex(549, (d_1740_v159_).length(0))
                (d_1740_v159_)[index297_] = (d_1689_v117_)[default__.safeIndex(len(d_1730_v153_), len(d_1689_v117_))]
                index298_ = default__.safeIndex(549, (d_1740_v159_).length(0))
                rhs355_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
                rhs356_ = default__.fm6(len(_dafny.SeqWithoutIsStrInference([p0 for d_1742_i24_ in range(default__.abs(569))])), globalState)
                rhs357_ = self.f30
                rhs358_ = (d_1729_v152_)[default__.safeIndex(p0, len(d_1729_v152_))]
                lhs338_ = globalState
                lhs339_ = globalState
                lhs340_ = d_1740_v159_
                lhs341_ = default__.safeIndex(549, (d_1740_v159_).length(0))
                lhs342_ = globalState
                lhs338_.f20 = rhs355_
                lhs339_.f28 = rhs356_
                lhs340_[lhs341_] = rhs357_
                lhs342_.f15 = rhs358_
            d_1743_v160_: _dafny.Seq
            d_1743_v160_ = _dafny.SeqWithoutIsStrInference([p1, (d_1512_v2_).f34])
            index299_ = default__.safeIndex(991, (d_1517_v7_).length(0))
            (d_1517_v7_)[index299_] = (p1 if (self).f29 else (d_1743_v160_)[default__.safeIndex(p0, len(d_1743_v160_))])
            d_1744_v161_: _dafny.Set
            d_1744_v161_ = _dafny.Set({(d_1512_v2_).f34, (0) - (len(_dafny.SeqWithoutIsStrInference([d_1514_v4_ for d_1745_i25_ in range(default__.abs(609))]))), (d_1512_v2_).f34})
            index300_ = default__.safeIndex(991, (d_1517_v7_).length(0))
            (d_1517_v7_)[index300_] = default__.safeDivisionInt((d_1512_v2_).f34, len(d_1744_v161_))
            d_1746_v162_: C4
            nw320_ = C4()
            nw320_.ctor__((self).f29, (d_1512_v2_).f34)
            d_1746_v162_ = nw320_
        d_1747_v163_: _dafny.MultiSet
        d_1747_v163_ = _dafny.MultiSet([(self).f29, False, (self).f32])
        if ((d_1747_v163_ if False else d_1747_v163_)).ispropersubset((d_1747_v163_).intersection(d_1747_v163_)):
            r2 = (d_1512_v2_).f34
            d_1748_v164_: C3
            nw321_ = C3()
            nw321_.ctor__(not(((d_1512_v2_).f33) and ((self).f32)), p0, False, self.f30)
            d_1748_v164_ = nw321_
            d_1749_v165_: D2
            d_1749_v165_ = D2_DC8(d_1514_v4_, (d_1512_v2_).f33)
            d_1750_v166_: _dafny.Array
            nw322_ = _dafny.Array(None, 20)
            nw322_[int(0)] = d_1514_v4_
            nw322_[int(1)] = d_1514_v4_
            nw322_[int(2)] = default__.fm42((d_1512_v2_).f33, not((self).f32), globalState)
            nw322_[int(3)] = d_1514_v4_
            nw322_[int(4)] = _dafny.CodePoint('e')
            nw322_[int(5)] = d_1514_v4_
            nw322_[int(6)] = default__.fm42((self).f31, False, globalState)
            nw322_[int(7)] = d_1514_v4_
            nw322_[int(8)] = d_1514_v4_
            nw322_[int(9)] = d_1514_v4_
            nw322_[int(10)] = d_1514_v4_
            nw322_[int(11)] = d_1514_v4_
            nw322_[int(12)] = d_1514_v4_
            nw322_[int(13)] = d_1514_v4_
            nw322_[int(14)] = d_1514_v4_
            nw322_[int(15)] = (d_1749_v165_).cf12
            nw322_[int(16)] = d_1514_v4_
            nw322_[int(17)] = _dafny.CodePoint('m')
            nw322_[int(18)] = d_1514_v4_
            nw322_[int(19)] = default__.fm42((self).f31, (self).f32, globalState)
            d_1750_v166_ = nw322_
            index301_ = default__.safeIndex(53, (d_1750_v166_).length(0))
            (d_1750_v166_)[index301_] = d_1514_v4_
            index302_ = default__.safeIndex(53, (d_1750_v166_).length(0))
            (d_1750_v166_)[index302_] = d_1514_v4_
            d_1751_v167_: _dafny.Seq
            d_1751_v167_ = _dafny.SeqWithoutIsStrInference([(self).f29])
            d_1752_v168_: _dafny.Map
            d_1752_v168_ = _dafny.Map({(d_1750_v166_)[default__.safeIndex(53, (d_1750_v166_).length(0))]: len(d_1751_v167_)})
            d_1753_v169_: _dafny.Map
            d_1753_v169_ = _dafny.Map({(self.f30) + (self.f30): ((d_1752_v168_)[d_1514_v4_] if (d_1514_v4_) in (d_1752_v168_) else p1)})
            d_1754_v170_: C6
            nw323_ = C6()
            nw323_.ctor__(True, self.f30)
            d_1754_v170_ = nw323_
            d_1755_v171_: _dafny.Seq
            d_1755_v171_ = _dafny.SeqWithoutIsStrInference([d_1754_v170_])
            d_1756_v172_: _dafny.Seq
            d_1756_v172_ = _dafny.SeqWithoutIsStrInference([(d_1748_v164_).f36, len((d_1755_v171_).set(default__.safeIndex(len(self.f30), len(d_1755_v171_)), d_1754_v170_))])
            d_1753_v169_ = ((d_1753_v169_) | (d_1753_v169_)) | ((d_1753_v169_) | (_dafny.Map({self.f30: len(d_1756_v172_)})))
            d_1757_v173_: _dafny.Map
            d_1757_v173_ = _dafny.Map({(p0) * (p1): (d_1512_v2_).f34})
            pat_let_tv35_ = d_1514_v4_
            d_1758_v174_: _dafny.Seq
            def iife127_(_pat_let35_0):
                def iife128_(d_1759_dt__update__tmp_h1_):
                    def iife129_(_pat_let36_0):
                        def iife130_(d_1760_dt__update_hcf12_h0_):
                            return D2_DC8(d_1760_dt__update_hcf12_h0_, (d_1759_dt__update__tmp_h1_).cf13)
                        return iife130_(_pat_let36_0)
                    return iife129_(pat_let_tv35_)
                return iife128_(_pat_let35_0)
            d_1758_v174_ = _dafny.SeqWithoutIsStrInference([D2_DC8((d_1750_v166_)[default__.safeIndex(53, (d_1750_v166_).length(0))], (d_1748_v164_).f35), iife127_(D2_DC8((d_1750_v166_)[default__.safeIndex(53, (d_1750_v166_).length(0))], True))])
            d_1757_v173_ = (d_1757_v173_).set(default__.safeModuloInt((d_1512_v2_).f34, len(default__.fm35(d_1758_v174_, (d_1748_v164_).f36, globalState))), 980)
        elif True:
            d_1761_v175_: _dafny.Seq
            d_1761_v175_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1762_v176_: _dafny.Seq
            d_1762_v176_ = _dafny.SeqWithoutIsStrInference([265, p0, (0) - (len(d_1761_v175_))])
            (globalState).f15 = not ((self).f29) or (((d_1512_v2_).f34) in (d_1762_v176_))
            d_1763_v177_: _dafny.Array
            nw324_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_1763_v177_ = nw324_
            index303_ = default__.safeIndex(98, (d_1763_v177_).length(0))
            (d_1763_v177_)[index303_] = d_1517_v7_
            d_1764_v178_: D10
            d_1764_v178_ = D10_DC33((self).f31, (d_1512_v2_).f34)
            index304_ = default__.safeIndex(98, (d_1763_v177_).length(0))
            rhs359_ = (d_1517_v7_ if (d_1512_v2_).f33 else d_1517_v7_)
            rhs360_ = self.f30
            rhs361_ = (d_1764_v178_).cf54
            lhs343_ = d_1763_v177_
            lhs344_ = default__.safeIndex(98, (d_1763_v177_).length(0))
            lhs345_ = globalState
            lhs346_ = globalState
            lhs343_[lhs344_] = rhs359_
            lhs345_.f11 = rhs360_
            lhs346_.f15 = rhs361_
            arr2_ = (d_1763_v177_)[default__.safeIndex(98, (d_1763_v177_).length(0))]
            index305_ = default__.safeIndex(471, ((d_1763_v177_)[default__.safeIndex(98, (d_1763_v177_).length(0))]).length(0))
            arr2_[index305_] = (0) - (len(self.f30))
            d_1765_v180_: D11
            d_1765_v180_ = D11_DC36(D6_DC23(True, (d_1512_v2_).f34), _dafny.Set({False, (self).f31}), (d_1512_v2_).f33)
            d_1766_v181_: _dafny.Seq
            d_1766_v181_ = _dafny.SeqWithoutIsStrInference([d_1517_v7_])
            d_1767_v182_: D6
            d_1767_v182_ = D6_DC23((d_1512_v2_).f33, len(d_1766_v181_))
            pat_let_tv36_ = d_1767_v182_
            d_1768_v183_: _dafny.Map
            def iife131_(_pat_let37_0):
                def iife132_(d_1769_dt__update__tmp_h2_):
                    def iife133_(_pat_let38_0):
                        def iife134_(d_1770_dt__update_hcf59_h0_):
                            return D11_DC36(d_1770_dt__update_hcf59_h0_, (d_1769_dt__update__tmp_h2_).cf60, (d_1769_dt__update__tmp_h2_).cf61)
                        return iife134_(_pat_let38_0)
                    return iife133_(pat_let_tv36_)
                return iife132_(_pat_let37_0)
            d_1768_v183_ = _dafny.Map({(d_1512_v2_).f33: iife131_(d_1765_v180_)})
            d_1771_v184_: _dafny.Set
            d_1771_v184_ = _dafny.Set({d_1768_v183_, d_1768_v183_})
            arr3_ = (d_1763_v177_)[default__.safeIndex(98, (d_1763_v177_).length(0))]
            index306_ = default__.safeIndex(471, ((d_1763_v177_)[default__.safeIndex(98, (d_1763_v177_).length(0))]).length(0))
            def iife135_():
                coll57_ = _dafny.Set()
                compr_57_: _dafny.Map
                for compr_57_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f29: D11_DC36(D6_DC23((d_1512_v2_).f33, len(_dafny.SeqWithoutIsStrInference([(d_1512_v2_).f33]))), _dafny.Set({False}), (self).f29)}) for d_1772_i26_ in range(default__.abs(412))])).Elements:
                    d_1773_v179_: _dafny.Map = compr_57_
                    if (d_1773_v179_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f29: D11_DC36(D6_DC23((d_1512_v2_).f33, len(_dafny.SeqWithoutIsStrInference([(d_1512_v2_).f33]))), _dafny.Set({False}), (self).f29)}) for d_1772_i26_ in range(default__.abs(412))])):
                        coll57_ = coll57_.union(_dafny.Set([d_1773_v179_]))
                return _dafny.Set(coll57_)
            rhs362_ = (d_1512_v2_).f34
            rhs363_ = (iife135_()
            ) == (d_1771_v184_)
            lhs347_ = (d_1763_v177_)[default__.safeIndex(98, (d_1763_v177_).length(0))]
            lhs348_ = default__.safeIndex(471, ((d_1763_v177_)[default__.safeIndex(98, (d_1763_v177_).length(0))]).length(0))
            lhs349_ = globalState
            lhs347_[lhs348_] = rhs362_
            lhs349_.f13 = rhs363_
            d_1774_v185_: C5
            nw325_ = C5()
            nw325_.ctor__((d_1512_v2_).f33, self.f30)
            d_1774_v185_ = nw325_
            d_1775_v186_: _dafny.Array
            nw326_ = _dafny.Array(_dafny.Map({}), 5)
            d_1775_v186_ = nw326_
            d_1776_v187_: _dafny.Map
            d_1776_v187_ = _dafny.Map({d_1747_v163_: p0})
            index307_ = default__.safeIndex(726, (d_1775_v186_).length(0))
            (d_1775_v186_)[index307_] = d_1776_v187_
            d_1777_v188_: _dafny.Array
            nw327_ = _dafny.Array(_dafny.Set({}), 6)
            d_1777_v188_ = nw327_
            d_1778_v189_: _dafny.Map
            d_1778_v189_ = _dafny.Map({(d_1512_v2_).f34: ((d_1763_v177_)[default__.safeIndex(98, (d_1763_v177_).length(0))])[default__.safeIndex(471, ((d_1763_v177_)[default__.safeIndex(98, (d_1763_v177_).length(0))]).length(0))]})
            d_1779_v190_: _dafny.Set
            d_1779_v190_ = _dafny.Set({d_1778_v189_})
            index308_ = default__.safeIndex(586, (d_1777_v188_).length(0))
            (d_1777_v188_)[index308_] = d_1779_v190_
            d_1780_v191_: _dafny.Seq
            d_1780_v191_ = _dafny.SeqWithoutIsStrInference([d_1774_v185_, d_1774_v185_])
            arr4_ = (d_1763_v177_)[default__.safeIndex(98, (d_1763_v177_).length(0))]
            index309_ = default__.safeIndex(471, ((d_1763_v177_)[default__.safeIndex(98, (d_1763_v177_).length(0))]).length(0))
            index310_ = default__.safeIndex(726, (d_1775_v186_).length(0))
            index311_ = default__.safeIndex(586, (d_1777_v188_).length(0))
            rhs364_ = 493
            rhs365_ = (d_1780_v191_)[default__.safeIndex(p0, len(d_1780_v191_))]
            rhs366_ = d_1776_v187_
            rhs367_ = d_1779_v190_
            lhs350_ = (d_1763_v177_)[default__.safeIndex(98, (d_1763_v177_).length(0))]
            lhs351_ = default__.safeIndex(471, ((d_1763_v177_)[default__.safeIndex(98, (d_1763_v177_).length(0))]).length(0))
            lhs352_ = d_1775_v186_
            lhs353_ = default__.safeIndex(726, (d_1775_v186_).length(0))
            lhs354_ = d_1777_v188_
            lhs355_ = default__.safeIndex(586, (d_1777_v188_).length(0))
            lhs350_[lhs351_] = rhs364_
            d_1774_v185_ = rhs365_
            lhs352_[lhs353_] = rhs366_
            lhs354_[lhs355_] = rhs367_
            d_1781_v192_: C1
            nw328_ = C1()
            nw328_.ctor__((self).f31)
            d_1781_v192_ = nw328_
        d_1782_i27_: int
        d_1782_i27_ = 0
        with _dafny.label("8"):
            while (self).f31:
                with _dafny.c_label("8"):
                    if (d_1782_i27_) >= (100):
                        raise _dafny.Break("8")
                    d_1782_i27_ = (d_1782_i27_) + (1)
                    (globalState).f18 = (d_1514_v4_) in (self.f30)
                    (globalState).f18 = (self).f31
                    d_1783_v193_: _dafny.Seq
                    d_1783_v193_ = _dafny.SeqWithoutIsStrInference([(d_1512_v2_).f33])
                    (globalState).f18 = not(((d_1512_v2_).f33) in (d_1783_v193_))
                    if (d_1512_v2_).f33:
                        d_1784_v194_: C3
                        nw329_ = C3()
                        nw329_.ctor__((self).f31, (d_1512_v2_).f34, (self).f31, _dafny.SeqWithoutIsStrInference([d_1514_v4_ for d_1785_i28_ in range(default__.abs(784))]))
                        d_1784_v194_ = nw329_
                        d_1784_v194_ = d_1784_v194_
                        d_1509_v0_ = d_1509_v0_
                        index312_ = default__.safeIndex(695, (d_1509_v0_).length(0))
                        (d_1509_v0_)[index312_] = (p1) <= (len(d_1511_v1_))
                        index313_ = default__.safeIndex(695, (d_1509_v0_).length(0))
                        (d_1509_v0_)[index313_] = (((0) - (p1) if not((self).f32) else (d_1512_v2_).f34)) <= ((0) - (((d_1784_v194_).f36 if (d_1512_v2_).f33 else len(self.f30))))
                        (globalState).f1 = p1
                        d_1786_v195_: D11
                        d_1786_v195_ = D11_DC35(self.f30, (self).f32)
                        d_1787_v196_: D2
                        d_1787_v196_ = D2_DC9((d_1786_v195_).cf57, (d_1784_v194_).f36, (d_1512_v2_).f33, ((d_1521_v11_).set((d_1512_v2_).f34, default__.abs((d_1784_v194_).f36))).cardinality)
                        d_1788_v197_: _dafny.MultiSet
                        d_1788_v197_ = _dafny.MultiSet([d_1747_v163_])
                        d_1789_v198_: _dafny.Seq
                        d_1789_v198_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_1787_v196_ = default__.fm18(((d_1747_v163_)[(self).f31] if ((self).f31) in (d_1747_v163_) else (d_1512_v2_).f34), (0) - (((d_1512_v2_).f34) + ((d_1788_v197_).cardinality)), (d_1512_v2_).f34, len((d_1789_v198_) + (d_1789_v198_)), globalState)
                    elif True:
                        d_1790_v200_: _dafny.Set
                        d_1790_v200_ = _dafny.Set({p1})
                        d_1791_v201_: _dafny.Map
                        def iife136_():
                            coll58_ = _dafny.Set()
                            compr_58_: int
                            for compr_58_ in _dafny.IntegerRange(555, 896):
                                d_1792_v199_: int = compr_58_
                                if ((555) <= (d_1792_v199_)) and ((d_1792_v199_) < (896)):
                                    coll58_ = coll58_.union(_dafny.Set([default__.safeDivisionInt(d_1792_v199_, p0)]))
                            return _dafny.Set(coll58_)
                        d_1791_v201_ = _dafny.Map({(d_1512_v2_).f33: (iife136_()
                        ).isdisjoint(d_1790_v200_)})
                        d_1793_v202_: _dafny.Set
                        d_1793_v202_ = _dafny.Set({(d_1512_v2_).f33})
                        d_1791_v201_ = (d_1791_v201_).set(((self).f31) not in (d_1793_v202_), (self).f31)
                        d_1794_v203_: _dafny.Map
                        d_1794_v203_ = _dafny.Map({self.f30: (d_1783_v193_)[default__.safeIndex((d_1512_v2_).f34, len(d_1783_v193_))]})
                        d_1795_v204_: _dafny.Set
                        d_1795_v204_ = _dafny.Set({_dafny.Map({(d_1512_v2_).f33: (d_1512_v2_).f34})})
                        def iife137_():
                            coll59_ = _dafny.Map()
                            compr_59_: _dafny.Seq
                            for compr_59_ in (d_1689_v117_).Elements:
                                d_1796_v205_: _dafny.Seq = compr_59_
                                if (d_1796_v205_) in (d_1689_v117_):
                                    coll59_[d_1796_v205_] = not((d_1512_v2_).f33)
                            return _dafny.Map(coll59_)
                        rhs368_ = (d_1517_v7_ if (d_1795_v204_).ispropersubset(d_1795_v204_) else (d_1517_v7_ if False else d_1517_v7_))
                        rhs369_ = iife137_()

                        lhs356_ = globalState
                        lhs356_.f17 = rhs368_
                        d_1794_v203_ = rhs369_
                        index314_ = default__.safeIndex(359, (d_1517_v7_).length(0))
                        (d_1517_v7_)[index314_] = (d_1512_v2_).f34
                        index315_ = default__.safeIndex(359, (d_1517_v7_).length(0))
                        rhs370_ = d_1514_v4_
                        rhs371_ = (p1) + ((d_1512_v2_).f34)
                        rhs372_ = (self.f30) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtnaivbv")))
                        lhs357_ = d_1517_v7_
                        lhs358_ = default__.safeIndex(359, (d_1517_v7_).length(0))
                        lhs359_ = globalState
                        d_1514_v4_ = rhs370_
                        lhs357_[lhs358_] = rhs371_
                        lhs359_.f15 = rhs372_
                        (globalState).f21 = len(self.f30)
                        (globalState).f13 = (self).f31
                    pass
            pass
        r0 = not((self).f32)
        d_1797_v206_: _dafny.Array
        def lambda87_(d_1798_v4_):
            def lambda88_(d_1799_i29_):
                return d_1798_v4_

            return lambda88_

        init50_ = lambda87_(d_1514_v4_)
        nw330_ = _dafny.Array(None, 17)
        for i0_50_ in range(nw330_.length(0)):
            nw330_[i0_50_] = init50_(i0_50_)
        d_1797_v206_ = nw330_
        r1 = d_1797_v206_
        r2 = p0
        d_1800_v207_: _dafny.Seq
        d_1800_v207_ = _dafny.SeqWithoutIsStrInference([(d_1512_v2_).f34])
        d_1801_v208_: _dafny.Map
        d_1801_v208_ = _dafny.Map({d_1800_v207_: d_1514_v4_})
        d_1802_v209_: _dafny.Seq
        d_1802_v209_ = _dafny.SeqWithoutIsStrInference([len(d_1801_v208_), (d_1521_v11_).cardinality])
        d_1803_v210_: _dafny.Map
        d_1803_v210_ = _dafny.Map({d_1800_v207_: (d_1512_v2_).f33})
        def iife138_():
            coll60_ = _dafny.Set()
            compr_60_: _dafny.Seq
            for compr_60_ in (d_1803_v210_).keys.Elements:
                d_1804_v211_: _dafny.Seq = compr_60_
                if (d_1804_v211_) in (d_1803_v210_):
                    coll60_ = coll60_.union(_dafny.Set([d_1804_v211_]))
            return _dafny.Set(coll60_)
        r3 = iife138_()
        
        return r0, r1, r2, r3

    @property
    def f31(self):
        return self._f31
    @property
    def f32(self):
        return self._f32

class C8:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self):
        pass
        pass

    def m0(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        d_1805_v0_: bool
        d_1805_v0_ = True
        (globalState).f28 = d_1805_v0_
        d_1806_v1_: _dafny.Array
        def lambda89_(d_1807_i1_):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))

        init51_ = lambda89_
        nw331_ = _dafny.Array(None, 24)
        for i0_51_ in range(nw331_.length(0)):
            nw331_[i0_51_] = init51_(i0_51_)
        d_1806_v1_ = nw331_
        guard_loop_11_: int
        for guard_loop_11_ in _dafny.IntegerRange(0, (d_1806_v1_).length(0)):
            d_1808_i0_: int = guard_loop_11_
            if (True) and (((0) <= (d_1808_i0_)) and ((d_1808_i0_) < ((d_1806_v1_).length(0)))):
                (d_1806_v1_)[(d_1808_i0_)] = (len((_dafny.Set({False, d_1805_v0_})) | (_dafny.Set({d_1805_v0_})))) >= (p0)
        d_1809_v2_: _dafny.Array
        nw332_ = _dafny.Array(_dafny.Seq({}), 14)
        d_1809_v2_ = nw332_
        d_1810_v3_: _dafny.MultiSet
        d_1810_v3_ = _dafny.MultiSet([d_1805_v0_])
        d_1811_v4_: D0
        d_1811_v4_ = D0_DC2(d_1810_v3_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1812_i2_ in range(default__.abs(617))]), d_1805_v0_)
        d_1813_v5_: _dafny.Seq
        d_1813_v5_ = _dafny.SeqWithoutIsStrInference([d_1805_v0_])
        d_1814_v6_: _dafny.Seq
        d_1814_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnoahtmjj"))
        d_1815_v7_: D0
        d_1815_v7_ = D0_DC0(d_1805_v0_)
        pat_let_tv37_ = d_1814_v6_
        pat_let_tv38_ = d_1814_v6_
        pat_let_tv39_ = d_1814_v6_
        pat_let_tv40_ = d_1813_v5_
        pat_let_tv41_ = p0
        pat_let_tv42_ = d_1813_v5_
        pat_let_tv43_ = p3
        pat_let_tv44_ = d_1805_v0_
        def lambda90_(source25_):
            if source25_.is_DC0:
                d_1816___mcc_h0_ = source25_.cf0
                d_1817_cf0_ = d_1816___mcc_h0_
                return d_1817_cf0_
            elif source25_.is_DC1:
                d_1818___mcc_h1_ = source25_.cf1
                d_1819___mcc_h2_ = source25_.cf2
                d_1820___mcc_h3_ = source25_.cf3
                d_1821_cf3_ = d_1820___mcc_h3_
                d_1822_cf2_ = d_1819___mcc_h2_
                d_1823_cf1_ = d_1818___mcc_h1_
                return (pat_let_tv37_) < (pat_let_tv38_)
            elif True:
                d_1824___mcc_h4_ = source25_.cf4
                d_1825___mcc_h5_ = source25_.cf5
                d_1826___mcc_h6_ = source25_.cf6
                d_1827_cf6_ = d_1826___mcc_h6_
                d_1828_cf5_ = d_1825___mcc_h5_
                d_1829_cf4_ = d_1824___mcc_h4_
                return d_1827_cf6_

        def iife139_(_pat_let39_0):
            def iife140_(d_1830_dt__update__tmp_h0_):
                def iife141_(_pat_let40_0):
                    def iife142_(d_1831_dt__update_hcf5_h0_):
                        def iife143_(_pat_let41_0):
                            def iife144_(d_1832_dt__update_hcf6_h0_):
                                return D0_DC2((d_1830_dt__update__tmp_h0_).cf4, d_1831_dt__update_hcf5_h0_, d_1832_dt__update_hcf6_h0_)
                            return iife144_(_pat_let41_0)
                        return iife143_((pat_let_tv40_)[default__.safeIndex(pat_let_tv41_, len(pat_let_tv42_))])
                    return iife142_(_pat_let40_0)
                return iife141_(pat_let_tv39_)
            return iife140_(_pat_let39_0)
        def lambda91_(source26_):
            if source26_.is_DC0:
                d_1833___mcc_h7_ = source26_.cf0
                d_1834_cf0_ = d_1833___mcc_h7_
                return d_1834_cf0_
            elif source26_.is_DC1:
                d_1835___mcc_h8_ = source26_.cf1
                d_1836___mcc_h9_ = source26_.cf2
                d_1837___mcc_h10_ = source26_.cf3
                d_1838_cf3_ = d_1837___mcc_h10_
                d_1839_cf2_ = d_1836___mcc_h9_
                d_1840_cf1_ = d_1835___mcc_h8_
                return (27) >= (pat_let_tv43_)
            elif True:
                d_1841___mcc_h11_ = source26_.cf4
                d_1842___mcc_h12_ = source26_.cf5
                d_1843___mcc_h13_ = source26_.cf6
                d_1844_cf6_ = d_1843___mcc_h13_
                d_1845_cf5_ = d_1842___mcc_h12_
                d_1846_cf4_ = d_1841___mcc_h11_
                return pat_let_tv44_

        rhs373_ = not(not(lambda90_(iife139_(d_1811_v4_))))
        rhs374_ = d_1809_v2_
        rhs375_ = lambda91_(d_1815_v7_)
        rhs376_ = (d_1814_v6_) + (d_1814_v6_)
        rhs377_ = d_1805_v0_
        lhs360_ = globalState
        lhs361_ = globalState
        lhs362_ = globalState
        d_1805_v0_ = rhs373_
        d_1809_v2_ = rhs374_
        lhs360_.f13 = rhs375_
        lhs361_.f11 = rhs376_
        lhs362_.f13 = rhs377_
        d_1847_i3_: int
        d_1847_i3_ = 0
        with _dafny.label("9"):
            while not(d_1805_v0_):
                with _dafny.c_label("9"):
                    if (d_1847_i3_) >= (100):
                        raise _dafny.Break("9")
                    d_1847_i3_ = (d_1847_i3_) + (1)
                    d_1848_v8_: _dafny.Array
                    nw333_ = _dafny.Array(D0.default()(), 20)
                    d_1848_v8_ = nw333_
                    pat_let_tv45_ = p0
                    index316_ = default__.safeIndex(298, (d_1848_v8_).length(0))
                    def iife145_(_pat_let42_0):
                        def iife146_(d_1849_dt__update__tmp_h1_):
                            def iife147_(_pat_let43_0):
                                def iife148_(d_1850_dt__update_hcf3_h0_):
                                    return D0_DC1((d_1849_dt__update__tmp_h1_).cf1, (d_1849_dt__update__tmp_h1_).cf2, d_1850_dt__update_hcf3_h0_)
                                return iife148_(_pat_let43_0)
                            return iife147_(pat_let_tv45_)
                        return iife146_(_pat_let42_0)
                    (d_1848_v8_)[index316_] = iife145_(D0_DC1(d_1806_v1_, p3, default__.fm0(p3, d_1805_v0_, d_1805_v0_, globalState)))
                    d_1851_v9_: D1
                    d_1851_v9_ = D1_DC3(d_1813_v5_)
                    d_1852_v10_: D0
                    d_1852_v10_ = D0_DC1(d_1806_v1_, default__.fm0(len((d_1851_v9_).cf7), d_1805_v0_, d_1805_v0_, globalState), p0)
                    index317_ = default__.safeIndex(298, (d_1848_v8_).length(0))
                    (d_1848_v8_)[index317_] = (d_1852_v10_ if False else D0_DC1(d_1806_v1_, p0, p3))
                    d_1853_v11_: _dafny.Seq
                    d_1853_v11_ = _dafny.SeqWithoutIsStrInference([p3, p3])
                    d_1854_v12_: _dafny.Map
                    d_1854_v12_ = _dafny.Map({(d_1853_v11_)[default__.safeIndex(p0, len(d_1853_v11_))]: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jr"))})
                    d_1855_v13_: _dafny.Seq
                    d_1855_v13_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_1856_v14_: _dafny.Array
                    nw334_ = _dafny.Array(None, 21)
                    nw334_[int(0)] = (d_1855_v13_).set(default__.safeIndex(858, len(d_1855_v13_)), p1)
                    nw334_[int(1)] = (d_1855_v13_ if d_1805_v0_ else _dafny.SeqWithoutIsStrInference([p1, p1]))
                    nw334_[int(2)] = d_1855_v13_
                    nw334_[int(3)] = d_1855_v13_
                    nw334_[int(4)] = d_1855_v13_
                    nw334_[int(5)] = d_1855_v13_
                    nw334_[int(6)] = (d_1855_v13_).set(default__.safeIndex(p0, len(d_1855_v13_)), p1)
                    nw334_[int(7)] = d_1855_v13_
                    nw334_[int(8)] = d_1855_v13_
                    nw334_[int(9)] = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
                    nw334_[int(10)] = (d_1855_v13_) + (d_1855_v13_)
                    nw334_[int(11)] = d_1855_v13_
                    nw334_[int(12)] = d_1855_v13_
                    nw334_[int(13)] = d_1855_v13_
                    nw334_[int(14)] = d_1855_v13_
                    nw334_[int(15)] = d_1855_v13_
                    nw334_[int(16)] = _dafny.SeqWithoutIsStrInference([p1, p1])
                    nw334_[int(17)] = d_1855_v13_
                    nw334_[int(18)] = (d_1855_v13_ if d_1805_v0_ else d_1855_v13_)
                    nw334_[int(19)] = (d_1855_v13_) + (_dafny.SeqWithoutIsStrInference([p1]))
                    nw334_[int(20)] = d_1855_v13_
                    d_1856_v14_ = nw334_
                    index318_ = default__.safeIndex(828, (d_1856_v14_).length(0))
                    (d_1856_v14_)[index318_] = (d_1855_v13_).set(default__.safeIndex(p3, len(d_1855_v13_)), p1)
                    d_1857_v15_: str
                    d_1857_v15_ = _dafny.CodePoint('e')
                    d_1858_v16_: _dafny.Map
                    d_1858_v16_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([p3])})
                    index319_ = default__.safeIndex(828, (d_1856_v14_).length(0))
                    rhs378_ = (default__.safeDivisionInt(p0, (0) - (p3)) if d_1805_v0_ else p0)
                    rhs379_ = d_1854_v12_
                    rhs380_ = (d_1855_v13_) + (d_1855_v13_)
                    rhs381_ = p3
                    rhs382_ = (not (d_1805_v0_) or (d_1805_v0_) if default__.fm1(d_1857_v15_, len(((d_1858_v16_)[p3] if (p3) in (d_1858_v16_) else d_1853_v11_)), (0) - (p3), d_1805_v0_, globalState) else (default__.fm0(p0, d_1805_v0_, d_1805_v0_, globalState)) > (p3))
                    lhs363_ = globalState
                    lhs364_ = d_1856_v14_
                    lhs365_ = default__.safeIndex(828, (d_1856_v14_).length(0))
                    lhs366_ = globalState
                    lhs367_ = globalState
                    lhs363_.f1 = rhs378_
                    d_1854_v12_ = rhs379_
                    lhs364_[lhs365_] = rhs380_
                    lhs366_.f21 = rhs381_
                    lhs367_.f13 = rhs382_
                    d_1859_v17_: _dafny.Array
                    nw335_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
                    d_1859_v17_ = nw335_
                    index320_ = default__.safeIndex(598, (d_1859_v17_).length(0))
                    (d_1859_v17_)[index320_] = d_1814_v6_
                    index321_ = default__.safeIndex(598, (d_1859_v17_).length(0))
                    (d_1859_v17_)[index321_] = ((d_1854_v12_)[(default__.fm0((0) - (p3), d_1805_v0_, d_1805_v0_, globalState)) + (default__.fm0(74, d_1805_v0_, d_1805_v0_, globalState))] if ((default__.fm0((0) - (p3), d_1805_v0_, d_1805_v0_, globalState)) + (default__.fm0(74, d_1805_v0_, d_1805_v0_, globalState))) in (d_1854_v12_) else d_1814_v6_)
                    d_1860_v18_: _dafny.Set
                    d_1860_v18_ = _dafny.Set({d_1805_v0_})
                    rhs383_ = p1
                    rhs384_ = d_1853_v11_
                    rhs385_ = len((d_1859_v17_)[default__.safeIndex(598, (d_1859_v17_).length(0))])
                    rhs386_ = _dafny.Set({(d_1853_v11_) == (_dafny.SeqWithoutIsStrInference([-408]))})
                    lhs368_ = globalState
                    lhs369_ = globalState
                    lhs370_ = globalState
                    lhs368_.f17 = rhs383_
                    lhs369_.f7 = rhs384_
                    lhs370_.f21 = rhs385_
                    d_1860_v18_ = rhs386_
                    pass
            pass
        source27_ = d_1811_v4_
        if source27_.is_DC0:
            d_1861___mcc_h14_ = source27_.cf0
            d_1862_cf0_ = d_1861___mcc_h14_
            d_1863_v19_: C0
            nw336_ = C0()
            nw336_.ctor__((d_1810_v3_).isdisjoint(_dafny.MultiSet(d_1813_v5_)))
            d_1863_v19_ = nw336_
            d_1864_v20_: _dafny.Array
            def lambda92_(d_1865_v19_, d_1866_p3_):
                def lambda93_(d_1867_i4_):
                    return D6_DC23(d_1865_v19_.f37, (0) - (d_1866_p3_))

                return lambda93_

            init52_ = lambda92_(d_1863_v19_, p3)
            nw337_ = _dafny.Array(None, 19)
            for i0_52_ in range(nw337_.length(0)):
                nw337_[i0_52_] = init52_(i0_52_)
            d_1864_v20_ = nw337_
            d_1864_v20_ = d_1864_v20_
            (globalState).f20 = _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('s') if False else _dafny.CodePoint('g')) for d_1868_i5_ in range(default__.abs(276))])
            d_1869_v21_: str
            d_1869_v21_ = _dafny.CodePoint('r')
            d_1869_v21_ = d_1869_v21_
        elif source27_.is_DC1:
            d_1870___mcc_h15_ = source27_.cf1
            d_1871___mcc_h16_ = source27_.cf2
            d_1872___mcc_h17_ = source27_.cf3
            d_1873_cf3_ = d_1872___mcc_h17_
            d_1874_cf2_ = d_1871___mcc_h16_
            d_1875_cf1_ = d_1870___mcc_h15_
            d_1876_v22_: D17
            d_1876_v22_ = D17_DC51(p0, p1)
            if (default__.safeModuloInt((0) - (p3), (d_1876_v22_).cf86)) <= (p3):
                d_1877_v23_: str
                d_1877_v23_ = _dafny.CodePoint('f')
                d_1877_v23_ = d_1877_v23_
                index322_ = default__.safeIndex(39, (d_1875_cf1_).length(0))
                (d_1875_cf1_)[index322_] = d_1805_v0_
                index323_ = default__.safeIndex(39, (d_1875_cf1_).length(0))
                (d_1875_cf1_)[index323_] = (d_1877_v23_) in (d_1814_v6_)
                d_1878_v24_: C6
                nw338_ = C6()
                nw338_.ctor__((d_1875_cf1_)[default__.safeIndex(39, (d_1875_cf1_).length(0))], d_1814_v6_)
                d_1878_v24_ = nw338_
                d_1879_v25_: _dafny.Map
                d_1879_v25_ = _dafny.Map({d_1805_v0_: 181})
                d_1879_v25_ = (d_1879_v25_).set(d_1805_v0_, d_1873_cf3_)
                (globalState).f13 = d_1805_v0_
            elif True:
                (globalState).f13 = d_1805_v0_
                d_1880_v26_: C3
                nw339_ = C3()
                nw339_.ctor__(d_1805_v0_, d_1873_cf3_, (d_1814_v6_) == (d_1814_v6_), d_1814_v6_)
                d_1880_v26_ = nw339_
                (globalState).f21 = len((d_1814_v6_) + (d_1814_v6_))
                (globalState).f18 = default__.fm1(default__.fm42(d_1805_v0_, (d_1880_v26_).f35, globalState), (0) - (p3), 753, (d_1805_v0_) or (d_1805_v0_), globalState)
                index324_ = default__.safeIndex(15, (d_1806_v1_).length(0))
                (d_1806_v1_)[index324_] = (False) not in (_dafny.MultiSet(d_1813_v5_))
                index325_ = default__.safeIndex(15, (d_1806_v1_).length(0))
                (d_1806_v1_)[index325_] = d_1805_v0_
            d_1881_v27_: _dafny.Seq
            d_1881_v27_ = _dafny.SeqWithoutIsStrInference([p3, p0, 743, p3])
            d_1882_v28_: _dafny.Set
            d_1882_v28_ = _dafny.Set({d_1881_v27_})
            d_1883_v30_: _dafny.Seq
            d_1883_v30_ = _dafny.SeqWithoutIsStrInference([d_1881_v27_, d_1881_v27_])
            d_1884_v32_: str
            d_1884_v32_ = _dafny.CodePoint('n')
            d_1885_v33_: _dafny.Seq
            def iife149_():
                coll61_ = _dafny.Map()
                compr_61_: D0
                for compr_61_ in (_dafny.Map({d_1815_v7_: d_1884_v32_})).keys.Elements:
                    d_1886_v31_: D0 = compr_61_
                    if (d_1886_v31_) in (_dafny.Map({d_1815_v7_: d_1884_v32_})):
                        coll61_[d_1886_v31_] = True
                return _dafny.Map(coll61_)
            d_1885_v33_ = _dafny.SeqWithoutIsStrInference([(d_1883_v30_)[default__.safeIndex(len(iife149_()
            ), len(d_1883_v30_))]])
            d_1887_v35_: _dafny.Map
            d_1887_v35_ = _dafny.Map({d_1805_v0_: d_1805_v0_})
            d_1888_v36_: _dafny.Map
            d_1888_v36_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([len(d_1813_v5_), d_1874_cf2_]): len((d_1887_v35_).set(not(d_1805_v0_), d_1805_v0_))})
            d_1889_v42_: _dafny.Array
            nw340_ = _dafny.Array(None, 24)
            nw340_[int(0)] = (d_1882_v28_) - (d_1882_v28_)
            nw340_[int(1)] = d_1882_v28_
            nw340_[int(2)] = d_1882_v28_
            nw340_[int(3)] = d_1882_v28_
            def iife150_():
                coll62_ = _dafny.Set()
                compr_62_: _dafny.Seq
                for compr_62_ in (_dafny.MultiSet([d_1881_v27_, _dafny.SeqWithoutIsStrInference([p3 for d_1890_i6_ in range(default__.abs(833))])])).Elements:
                    d_1891_v29_: _dafny.Seq = compr_62_
                    if (d_1891_v29_) in (_dafny.MultiSet([d_1881_v27_, _dafny.SeqWithoutIsStrInference([p3 for d_1890_i6_ in range(default__.abs(833))])])):
                        coll62_ = coll62_.union(_dafny.Set([d_1891_v29_]))
                return _dafny.Set(coll62_)
            nw340_[int(4)] = (default__.fm25(False, d_1805_v0_, d_1814_v6_, globalState)) | (iife150_()
            )
            nw340_[int(5)] = default__.fm25(d_1805_v0_, d_1805_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")), globalState)
            def iife151_():
                coll63_ = _dafny.Set()
                compr_63_: _dafny.Seq
                for compr_63_ in (d_1883_v30_).Elements:
                    d_1892_v34_: _dafny.Seq = compr_63_
                    if (d_1892_v34_) in (d_1883_v30_):
                        coll63_ = coll63_.union(_dafny.Set([d_1892_v34_]))
                return _dafny.Set(coll63_)
            nw340_[int(6)] = iife151_()
            
            nw340_[int(7)] = d_1882_v28_
            nw340_[int(8)] = d_1882_v28_
            def iife152_():
                coll64_ = _dafny.Set()
                compr_64_: _dafny.Seq
                for compr_64_ in (d_1888_v36_).keys.Elements:
                    d_1893_v37_: _dafny.Seq = compr_64_
                    if (d_1893_v37_) in (d_1888_v36_):
                        coll64_ = coll64_.union(_dafny.Set([d_1893_v37_]))
                return _dafny.Set(coll64_)
            nw340_[int(9)] = (_dafny.Set({_dafny.SeqWithoutIsStrInference([len(d_1887_v35_)]), d_1881_v27_})).intersection(iife152_()
            )
            def iife153_():
                coll65_ = _dafny.Set()
                compr_65_: _dafny.Seq
                for compr_65_ in (d_1888_v36_).keys.Elements:
                    d_1894_v38_: _dafny.Seq = compr_65_
                    if (d_1894_v38_) in (d_1888_v36_):
                        coll65_ = coll65_.union(_dafny.Set([d_1894_v38_]))
                return _dafny.Set(coll65_)
            nw340_[int(10)] = iife153_()
            
            nw340_[int(11)] = d_1882_v28_
            nw340_[int(12)] = (d_1882_v28_).intersection(d_1882_v28_)
            nw340_[int(13)] = (d_1882_v28_) - (_dafny.Set({d_1881_v27_}))
            def iife154_():
                coll66_ = _dafny.Set()
                compr_66_: _dafny.Seq
                for compr_66_ in (d_1885_v33_).Elements:
                    d_1895_v39_: _dafny.Seq = compr_66_
                    if (d_1895_v39_) in (d_1885_v33_):
                        coll66_ = coll66_.union(_dafny.Set([d_1895_v39_]))
                return _dafny.Set(coll66_)
            nw340_[int(14)] = (iife154_()
            ) - (d_1882_v28_)
            nw340_[int(15)] = d_1882_v28_
            nw340_[int(16)] = d_1882_v28_
            nw340_[int(17)] = (d_1882_v28_).intersection(d_1882_v28_)
            nw340_[int(18)] = d_1882_v28_
            nw340_[int(19)] = (d_1882_v28_).intersection(d_1882_v28_)
            def iife155_():
                def iife157_():
                    coll69_ = _dafny.Set()
                    compr_69_: _dafny.Seq
                    for compr_69_ in (d_1883_v30_).Elements:
                        d_1898_v40_: _dafny.Seq = compr_69_
                        if (d_1898_v40_) in (d_1883_v30_):
                            coll69_ = coll69_.union(_dafny.Set([d_1898_v40_]))
                    return _dafny.Set(coll69_)
                coll67_ = _dafny.Set()
                def iife156_():
                    coll68_ = _dafny.Set()
                    compr_68_: _dafny.Seq
                    for compr_68_ in (d_1883_v30_).Elements:
                        d_1896_v40_: _dafny.Seq = compr_68_
                        if (d_1896_v40_) in (d_1883_v30_):
                            coll68_ = coll68_.union(_dafny.Set([d_1896_v40_]))
                    return _dafny.Set(coll68_)
                compr_67_: _dafny.Seq
                for compr_67_ in (iife156_()
                ).Elements:
                    d_1897_v41_: _dafny.Seq = compr_67_
                    if (d_1897_v41_) in (iife157_()
                    ):
                        coll67_ = coll67_.union(_dafny.Set([d_1897_v41_]))
                return _dafny.Set(coll67_)
            nw340_[int(20)] = iife155_()
            
            nw340_[int(21)] = _dafny.Set({d_1881_v27_, d_1881_v27_})
            nw340_[int(22)] = d_1882_v28_
            nw340_[int(23)] = d_1882_v28_
            d_1889_v42_ = nw340_
            index326_ = default__.safeIndex(515, (d_1889_v42_).length(0))
            (d_1889_v42_)[index326_] = d_1882_v28_
            index327_ = default__.safeIndex(515, (d_1889_v42_).length(0))
            (d_1889_v42_)[index327_] = d_1882_v28_
            d_1899_v43_: D9
            d_1899_v43_ = D9_DC28(d_1813_v5_, d_1805_v0_, (d_1881_v27_)[default__.safeIndex(len(_dafny.Set({d_1874_cf2_})), len(d_1881_v27_))], True, d_1805_v0_)
            (globalState).f19 = (d_1899_v43_).cf43
            (globalState).f1 = p0
        elif True:
            d_1900___mcc_h18_ = source27_.cf4
            d_1901___mcc_h19_ = source27_.cf5
            d_1902___mcc_h20_ = source27_.cf6
            d_1903_cf6_ = d_1902___mcc_h20_
            d_1904_cf5_ = d_1901___mcc_h19_
            d_1905_cf4_ = d_1900___mcc_h18_
            d_1906_v44_: _dafny.Map
            d_1906_v44_ = _dafny.Map({(_dafny.MultiSet([d_1805_v0_, True])).cardinality: d_1903_cf6_})
            d_1907_v45_: _dafny.Seq
            d_1907_v45_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1908_v46_: _dafny.Map
            d_1908_v46_ = _dafny.Map({d_1906_v44_: (d_1907_v45_)[default__.safeIndex((0) - (p0), len(d_1907_v45_))]})
            d_1909_v47_: _dafny.Map
            d_1909_v47_ = _dafny.Map({d_1903_cf6_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))})
            d_1910_v48_: _dafny.Map
            d_1910_v48_ = _dafny.Map({len(d_1909_v47_): p1})
            d_1911_v49_: _dafny.Array
            nw341_ = _dafny.Array(None, 19)
            nw341_[int(0)] = (p1 if d_1805_v0_ else p1)
            nw341_[int(1)] = p1
            nw341_[int(2)] = p1
            nw341_[int(3)] = p1
            nw341_[int(4)] = p1
            nw341_[int(5)] = ((d_1908_v46_)[d_1906_v44_] if (d_1906_v44_) in (d_1908_v46_) else p1)
            nw341_[int(6)] = p1
            nw341_[int(7)] = p1
            nw341_[int(8)] = p1
            nw341_[int(9)] = ((d_1910_v48_)[len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1912_i7_ in range(default__.abs(-917))]))] if (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1913_i7_ in range(default__.abs(-917))]))) in (d_1910_v48_) else p1)
            nw341_[int(10)] = p1
            nw341_[int(11)] = p1
            nw341_[int(12)] = ((d_1910_v48_)[default__.fm0(p3, (d_1813_v5_)[default__.safeIndex((0) - (len(d_1906_v44_)), len(d_1813_v5_))], d_1903_cf6_, globalState)] if (default__.fm0(p3, (d_1813_v5_)[default__.safeIndex((0) - (len(d_1906_v44_)), len(d_1813_v5_))], d_1903_cf6_, globalState)) in (d_1910_v48_) else p1)
            nw341_[int(13)] = p1
            nw341_[int(14)] = p1
            nw341_[int(15)] = p1
            nw341_[int(16)] = p1
            nw341_[int(17)] = (p1 if False else p1)
            nw341_[int(18)] = p1
            d_1911_v49_ = nw341_
            index328_ = default__.safeIndex(916, (d_1911_v49_).length(0))
            (d_1911_v49_)[index328_] = (p1 if d_1805_v0_ else p1)
            index329_ = default__.safeIndex(916, (d_1911_v49_).length(0))
            (d_1911_v49_)[index329_] = p1
            d_1914_v50_: str
            d_1914_v50_ = _dafny.CodePoint('u')
            d_1915_v51_: D2
            d_1915_v51_ = D2_DC8(d_1914_v50_, d_1903_cf6_)
            d_1916_v52_: _dafny.Seq
            d_1916_v52_ = _dafny.SeqWithoutIsStrInference([d_1915_v51_, d_1915_v51_])
            d_1813_v5_ = (default__.fm35(d_1916_v52_, p0, globalState)) + (d_1813_v5_)
            d_1917_v53_: C6
            nw342_ = C6()
            nw342_.ctor__(d_1805_v0_, d_1814_v6_)
            d_1917_v53_ = nw342_
            (globalState).f11 = ((d_1904_cf5_) + (d_1904_cf5_)) + (d_1814_v6_)
        if d_1805_v0_:
            pat_let_tv46_ = d_1814_v6_
            d_1918_v54_: _dafny.Map
            def iife158_(_pat_let44_0):
                def iife159_(d_1919_dt__update__tmp_h2_):
                    def iife160_(_pat_let45_0):
                        def iife161_(d_1920_dt__update_hcf5_h1_):
                            return D0_DC2((d_1919_dt__update__tmp_h2_).cf4, d_1920_dt__update_hcf5_h1_, (d_1919_dt__update__tmp_h2_).cf6)
                        return iife161_(_pat_let45_0)
                    return iife160_(pat_let_tv46_)
                return iife159_(_pat_let44_0)
            d_1918_v54_ = _dafny.Map({p1: iife158_(d_1811_v4_)})
            d_1918_v54_ = (d_1918_v54_).set(p1, d_1811_v4_)
            (globalState).f15 = not((not((default__.fm10(p0, p0, globalState) if d_1805_v0_ else not(False)))) == (d_1805_v0_))
            (globalState).f13 = d_1805_v0_
            (globalState).f1 = ((483) + (p3)) * (p0)
            d_1921_v55_: str
            d_1921_v55_ = _dafny.CodePoint('j')
            d_1922_v56_: D2
            d_1922_v56_ = D2_DC8(d_1921_v55_, d_1805_v0_)
            d_1923_v57_: _dafny.Seq
            d_1923_v57_ = _dafny.SeqWithoutIsStrInference([d_1922_v56_, D2_DC8(d_1921_v55_, d_1805_v0_)])
            d_1924_v58_: D18
            d_1924_v58_ = D18_DC54(d_1923_v57_)
            d_1925_v59_: D1
            d_1925_v59_ = D1_DC3(default__.fm35((d_1924_v58_).cf93, p0, globalState))
            pat_let_tv47_ = d_1813_v5_
            def iife162_(_pat_let46_0):
                def iife163_(d_1926_dt__update__tmp_h3_):
                    def iife164_(_pat_let47_0):
                        def iife165_(d_1927_dt__update_hcf7_h0_):
                            return D1_DC3(d_1927_dt__update_hcf7_h0_)
                        return iife165_(_pat_let47_0)
                    return iife164_(pat_let_tv47_)
                return iife163_(_pat_let46_0)
            source28_ = iife162_(d_1925_v59_)
            if source28_.is_DC4:
                d_1928___mcc_h21_ = source28_.cf8
                d_1929_cf8_ = d_1928___mcc_h21_
                index330_ = default__.safeIndex(653, (p1).length(0))
                (p1)[index330_] = (p3) + (p0)
                d_1930_v60_: _dafny.Map
                d_1930_v60_ = _dafny.Map({d_1806_v1_: p3})
                index331_ = default__.safeIndex(653, (p1).length(0))
                (p1)[index331_] = ((d_1930_v60_)[d_1806_v1_] if (d_1806_v1_) in (d_1930_v60_) else (p3) - (p0))
                d_1931_v61_: _dafny.MultiSet
                d_1931_v61_ = _dafny.MultiSet([p0, p3])
                d_1932_v62_: C7
                nw343_ = C7()
                nw343_.ctor__(d_1929_cf8_, not (d_1929_cf8_) or (default__.fm10(len(d_1814_v6_), (p1)[default__.safeIndex(653, (p1).length(0))], globalState)), (d_1931_v61_).issubset(d_1931_v61_), d_1814_v6_)
                d_1932_v62_ = nw343_
                d_1933_v63_: C1
                nw344_ = C1()
                nw344_.ctor__(True)
                d_1933_v63_ = nw344_
                d_1934_v64_: _dafny.Set
                d_1934_v64_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_1921_v55_ for d_1935_i8_ in range(default__.abs(349))])})
                d_1936_v65_: _dafny.Array
                def lambda94_(d_1937_v6_):
                    def lambda95_(d_1938_i9_):
                        return default__.safeDivisionInt(d_1938_i9_, len(d_1937_v6_))

                    return lambda95_

                init53_ = lambda94_(d_1814_v6_)
                nw345_ = _dafny.Array(None, 15)
                for i0_53_ in range(nw345_.length(0)):
                    nw345_[i0_53_] = init53_(i0_53_)
                d_1936_v65_ = nw345_
                d_1939_v66_: _dafny.Map
                d_1939_v66_ = _dafny.Map({d_1936_v65_: default__.fm6((p1)[default__.safeIndex(653, (p1).length(0))], globalState)})
                d_1929_cf8_ = ((len(d_1934_v64_)) - (len(d_1939_v66_))) == (p0)
            elif source28_.is_DC3:
                d_1940___mcc_h22_ = source28_.cf7
                d_1941_cf7_ = d_1940___mcc_h22_
                index332_ = default__.safeIndex(657, (d_1806_v1_).length(0))
                (d_1806_v1_)[index332_] = not (d_1805_v0_) or (d_1805_v0_)
                d_1942_v67_: _dafny.MultiSet
                d_1942_v67_ = _dafny.MultiSet([d_1941_cf7_])
                index333_ = default__.safeIndex(657, (d_1806_v1_).length(0))
                (d_1806_v1_)[index333_] = (d_1942_v67_).ispropersubset(_dafny.MultiSet([d_1813_v5_]))
                d_1943_v68_: _dafny.Map
                d_1943_v68_ = _dafny.Map({default__.fm1(d_1921_v55_, (0) - (p0), p0, False, globalState): d_1925_v59_})
                d_1944_v69_: D5
                d_1944_v69_ = D5_DC20(D5_DC16(d_1943_v68_))
                d_1945_v70_: _dafny.Map
                d_1945_v70_ = _dafny.Map({p3: d_1944_v69_})
                d_1945_v70_ = d_1945_v70_
                (globalState).f11 = d_1814_v6_
                d_1946_v71_: _dafny.Set
                d_1946_v71_ = _dafny.Set({d_1814_v6_})
                d_1947_v72_: D5
                d_1947_v72_ = D5_DC17(d_1946_v71_)
                d_1947_v72_ = d_1947_v72_
            elif True:
                d_1948___mcc_h23_ = source28_.cf9
                d_1949_cf9_ = d_1948___mcc_h23_
                (globalState).f28 = not((d_1811_v4_).cf6)
                index334_ = default__.safeIndex(394, (p1).length(0))
                (p1)[index334_] = p0
                index335_ = default__.safeIndex(394, (p1).length(0))
                (p1)[index335_] = 360
                d_1950_v73_: _dafny.MultiSet
                d_1950_v73_ = _dafny.MultiSet([(p1)[default__.safeIndex(394, (p1).length(0))]])
                d_1951_v74_: _dafny.Array
                nw346_ = _dafny.Array(None, 20)
                nw346_[int(0)] = (p1)[default__.safeIndex(394, (p1).length(0))]
                nw346_[int(1)] = p0
                nw346_[int(2)] = p3
                nw346_[int(3)] = p3
                nw346_[int(4)] = (p3) - ((d_1950_v73_).cardinality)
                nw346_[int(5)] = p0
                nw346_[int(6)] = p0
                nw346_[int(7)] = p0
                nw346_[int(8)] = -609
                nw346_[int(9)] = (p1)[default__.safeIndex(394, (p1).length(0))]
                nw346_[int(10)] = default__.fm0(p0, d_1805_v0_, d_1805_v0_, globalState)
                nw346_[int(11)] = (p3) + (p3)
                nw346_[int(12)] = p3
                nw346_[int(13)] = len(d_1814_v6_)
                nw346_[int(14)] = -837
                nw346_[int(15)] = p3
                nw346_[int(16)] = p3
                nw346_[int(17)] = default__.fm9(globalState)
                nw346_[int(18)] = (d_1950_v73_).cardinality
                nw346_[int(19)] = (p1)[default__.safeIndex(394, (p1).length(0))]
                d_1951_v74_ = nw346_
                d_1952_v75_: _dafny.Seq
                d_1952_v75_ = _dafny.SeqWithoutIsStrInference([d_1814_v6_])
                index336_ = default__.safeIndex(394, (p1).length(0))
                rhs387_ = d_1921_v55_
                rhs388_ = d_1951_v74_
                rhs389_ = default__.safeDivisionInt((0) - ((p1)[default__.safeIndex(394, (p1).length(0))]), len((d_1814_v6_) + (d_1814_v6_)))
                rhs390_ = ((d_1952_v75_)[default__.safeIndex(367, len(d_1952_v75_))]) + ((d_1814_v6_ if d_1805_v0_ else _dafny.SeqWithoutIsStrInference([d_1921_v55_ for d_1953_i10_ in range(default__.abs(950))])))
                lhs371_ = globalState
                lhs372_ = p1
                lhs373_ = default__.safeIndex(394, (p1).length(0))
                lhs374_ = globalState
                d_1921_v55_ = rhs387_
                lhs371_.f17 = rhs388_
                lhs372_[lhs373_] = rhs389_
                lhs374_.f11 = rhs390_
                d_1954_v76_: C7
                nw347_ = C7()
                nw347_.ctor__(d_1805_v0_, d_1805_v0_, d_1805_v0_, ((d_1814_v6_) + (_dafny.SeqWithoutIsStrInference([d_1921_v55_ for d_1955_i11_ in range(default__.abs(998))]))).set(default__.safeIndex(p0, len((d_1814_v6_) + (_dafny.SeqWithoutIsStrInference([d_1921_v55_ for d_1956_i11_ in range(default__.abs(998))])))), d_1921_v55_))
                d_1954_v76_ = nw347_
        elif True:
            d_1957_v77_: _dafny.Array
            nw348_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
            d_1957_v77_ = nw348_
            d_1958_v78_: _dafny.Map
            d_1958_v78_ = _dafny.Map({d_1957_v77_: d_1805_v0_})
            d_1959_v79_: _dafny.Seq
            d_1959_v79_ = _dafny.SeqWithoutIsStrInference([d_1957_v77_])
            if not(((d_1958_v78_)[(d_1959_v79_)[default__.safeIndex(p0, len(d_1959_v79_))]] if ((d_1959_v79_)[default__.safeIndex(p0, len(d_1959_v79_))]) in (d_1958_v78_) else d_1805_v0_)):
                (globalState).f13 = not(d_1805_v0_)
                d_1960_v80_: _dafny.Seq
                d_1960_v80_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                (globalState).f19 = len((d_1960_v80_).set(default__.safeIndex((0) - ((_dafny.MultiSet([d_1805_v0_, d_1805_v0_])).cardinality), len(d_1960_v80_)), p0))
                d_1961_v81_: _dafny.Map
                d_1961_v81_ = _dafny.Map({p1: p3})
                d_1961_v81_ = ((d_1961_v81_) | (d_1961_v81_)) | (d_1961_v81_)
                index337_ = default__.safeIndex(588, (d_1957_v77_).length(0))
                (d_1957_v77_)[index337_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rpkac"))
                index338_ = default__.safeIndex(588, (d_1957_v77_).length(0))
                (d_1957_v77_)[index338_] = d_1814_v6_
                d_1962_v82_: _dafny.MultiSet
                d_1962_v82_ = _dafny.MultiSet([d_1960_v80_])
                rhs391_ = ((d_1962_v82_).set(d_1960_v80_, default__.abs(p0))).cardinality
                rhs392_ = ((d_1814_v6_) + (d_1814_v6_)) != ((d_1814_v6_) + ((d_1957_v77_)[default__.safeIndex(588, (d_1957_v77_).length(0))]))
                rhs393_ = not(d_1805_v0_)
                lhs375_ = globalState
                lhs376_ = globalState
                lhs377_ = globalState
                lhs375_.f1 = rhs391_
                lhs376_.f15 = rhs392_
                lhs377_.f13 = rhs393_
            elif True:
                index339_ = default__.safeIndex(65, (p1).length(0))
                (p1)[index339_] = p0
                d_1963_v83_: _dafny.MultiSet
                d_1963_v83_ = _dafny.MultiSet([p3, p0])
                d_1964_v84_: _dafny.MultiSet
                d_1964_v84_ = _dafny.MultiSet([(d_1963_v83_) - (d_1963_v83_)])
                index340_ = default__.safeIndex(65, (p1).length(0))
                (p1)[index340_] = (d_1964_v84_).cardinality
                d_1965_v85_: _dafny.Array
                nw349_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_1965_v85_ = nw349_
                index341_ = default__.safeIndex(441, (d_1965_v85_).length(0))
                (d_1965_v85_)[index341_] = p2
                index342_ = default__.safeIndex(441, (d_1965_v85_).length(0))
                (d_1965_v85_)[index342_] = p2
                d_1966_v86_: _dafny.Array
                nw350_ = _dafny.Array(_dafny.Map({}), 29)
                d_1966_v86_ = nw350_
                d_1966_v86_ = d_1966_v86_
                d_1806_v1_ = d_1806_v1_
                (globalState).f13 = not(default__.fm19((d_1814_v6_)[default__.safeIndex((p1)[default__.safeIndex(65, (p1).length(0))], len(d_1814_v6_))], globalState))
            d_1967_v87_: str
            d_1967_v87_ = _dafny.CodePoint('d')
            d_1805_v0_ = default__.fm19(d_1967_v87_, globalState)
            (globalState).f13 = d_1805_v0_
            d_1968_v88_: _dafny.Set
            d_1968_v88_ = _dafny.Set({p1, p1, p1})
            d_1969_v89_: C2
            nw351_ = C2()
            nw351_.ctor__(len((d_1813_v5_) + (d_1813_v5_)), d_1968_v88_, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxlbbb")))
            d_1969_v89_ = nw351_
            d_1970_v90_: _dafny.Array
            nw352_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_1970_v90_ = nw352_
            d_1971_v91_: _dafny.Array
            nw353_ = _dafny.Array(None, 1)
            d_1971_v91_ = nw353_
            index343_ = default__.safeIndex(96, (d_1970_v90_).length(0))
            (d_1970_v90_)[index343_] = d_1971_v91_
            index344_ = default__.safeIndex(96, (d_1970_v90_).length(0))
            (d_1970_v90_)[index344_] = d_1971_v91_
        r0 = not(d_1805_v0_)
        return r0

    def m1(self, p0, globalState):
        r0: int = int(0)
        d_1972_v0_: int
        d_1972_v0_ = 553
        d_1973_v1_: _dafny.MultiSet
        d_1973_v1_ = _dafny.MultiSet([d_1972_v0_])
        d_1973_v1_ = d_1973_v1_
        d_1974_v2_: _dafny.Array
        def lambda96_(d_1975_i1_):
            return (D0_DC0(False)).cf0

        init54_ = lambda96_
        nw354_ = _dafny.Array(None, 24)
        for i0_54_ in range(nw354_.length(0)):
            nw354_[i0_54_] = init54_(i0_54_)
        d_1974_v2_ = nw354_
        guard_loop_12_: int
        for guard_loop_12_ in _dafny.IntegerRange(0, (d_1974_v2_).length(0)):
            d_1976_i0_: int = guard_loop_12_
            if (True) and (((0) <= (d_1976_i0_)) and ((d_1976_i0_) < ((d_1974_v2_).length(0)))):
                def iife166_():
                    coll70_ = _dafny.Set()
                    compr_70_: int
                    for compr_70_ in _dafny.IntegerRange(833, 56):
                        d_1977_v3_: int = compr_70_
                        if ((833) <= (d_1977_v3_)) and ((d_1977_v3_) < (56)):
                            coll70_ = coll70_.union(_dafny.Set([(d_1977_v3_) * ((_dafny.MultiSet([not(True)])).cardinality)]))
                    return _dafny.Set(coll70_)
                (d_1974_v2_)[(d_1976_i0_)] = ((_dafny.Set({852, d_1972_v0_})).intersection(_dafny.Set({d_1972_v0_}))).isdisjoint(iife166_()
                )
        hi11_ = 518
        for d_1978_i2_ in range(default__.safeModuloInt(969, d_1972_v0_), hi11_):
            (globalState).f1 = d_1972_v0_
            (globalState).f14 = d_1972_v0_
            d_1979_v4_: _dafny.Seq
            d_1979_v4_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1979_v4_ = d_1979_v4_
            (globalState).f20 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iepcaxj"))
        d_1980_v5_: bool
        d_1980_v5_ = False
        d_1981_v6_: _dafny.Seq
        d_1981_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
        d_1982_v7_: C4
        nw355_ = C4()
        nw355_.ctor__(not(d_1980_v5_), len(d_1981_v6_))
        d_1982_v7_ = nw355_
        d_1983_v8_: _dafny.MultiSet
        d_1983_v8_ = _dafny.MultiSet([d_1982_v7_, d_1982_v7_])
        d_1983_v8_ = (d_1983_v8_) - (_dafny.MultiSet([d_1982_v7_]))
        if False:
            d_1984_v9_: _dafny.Seq
            d_1984_v9_ = _dafny.SeqWithoutIsStrInference([d_1973_v1_])
            d_1985_v10_: _dafny.Seq
            d_1985_v10_ = _dafny.SeqWithoutIsStrInference([((_dafny.SeqWithoutIsStrInference([d_1973_v1_, d_1973_v1_]) if True else d_1984_v9_)).set(default__.safeIndex(d_1972_v0_, len((_dafny.SeqWithoutIsStrInference([d_1973_v1_, d_1973_v1_]) if True else d_1984_v9_))), d_1973_v1_)])
            d_1984_v9_ = (d_1985_v10_)[default__.safeIndex(len(p0), len(d_1985_v10_))]
            d_1986_v11_: _dafny.Map
            d_1986_v11_ = _dafny.Map({d_1980_v5_: (1) <= ((d_1982_v7_).f34)})
            (globalState).f28 = ((d_1986_v11_)[(d_1982_v7_).f33] if ((d_1982_v7_).f33) in (d_1986_v11_) else d_1980_v5_)
            d_1987_v12_: _dafny.Array
            nw356_ = _dafny.Array(None, 1)
            nw356_[int(0)] = d_1973_v1_
            d_1987_v12_ = nw356_
            d_1987_v12_ = d_1987_v12_
            nw357_ = _dafny.Array(False, 10)
            d_1974_v2_ = nw357_
            if ((d_1982_v7_).f33 if (d_1982_v7_).f33 else (d_1982_v7_).f33):
                r0 = (len(d_1981_v6_)) + ((d_1982_v7_).f34)
                d_1988_v13_: D2
                d_1988_v13_ = D2_DC7((d_1982_v7_).f33)
                d_1989_v14_: _dafny.Map
                d_1989_v14_ = _dafny.Map({d_1988_v13_: (d_1982_v7_).f33})
                d_1989_v14_ = (d_1989_v14_).set(d_1988_v13_, True)
                d_1990_v15_: _dafny.Array
                nw358_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_1990_v15_ = nw358_
                d_1991_v16_: str
                d_1991_v16_ = _dafny.CodePoint('h')
                d_1992_v17_: _dafny.Array
                nw359_ = _dafny.Array(None, 20)
                nw359_[int(0)] = _dafny.SeqWithoutIsStrInference([d_1991_v16_])
                nw359_[int(1)] = d_1981_v6_
                nw359_[int(2)] = _dafny.SeqWithoutIsStrInference([d_1991_v16_])
                nw359_[int(3)] = d_1981_v6_
                nw359_[int(4)] = (_dafny.SeqWithoutIsStrInference([d_1991_v16_ for d_1993_i3_ in range(default__.abs(-964))])).set(default__.safeIndex((d_1982_v7_).f34, len(_dafny.SeqWithoutIsStrInference([d_1991_v16_ for d_1994_i3_ in range(default__.abs(-964))]))), d_1991_v16_)
                nw359_[int(5)] = d_1981_v6_
                nw359_[int(6)] = d_1981_v6_
                nw359_[int(7)] = d_1981_v6_
                nw359_[int(8)] = d_1981_v6_
                nw359_[int(9)] = d_1981_v6_
                nw359_[int(10)] = d_1981_v6_
                nw359_[int(11)] = d_1981_v6_
                nw359_[int(12)] = d_1981_v6_
                nw359_[int(13)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i')])
                nw359_[int(14)] = d_1981_v6_
                nw359_[int(15)] = d_1981_v6_
                nw359_[int(16)] = d_1981_v6_
                nw359_[int(17)] = d_1981_v6_
                nw359_[int(18)] = d_1981_v6_
                nw359_[int(19)] = d_1981_v6_
                d_1992_v17_ = nw359_
                index345_ = default__.safeIndex(349, (d_1990_v15_).length(0))
                (d_1990_v15_)[index345_] = d_1992_v17_
                d_1995_v18_: _dafny.Array
                def lambda97_(d_1996_v0_):
                    def lambda98_(d_1997_i4_):
                        return (d_1997_i4_) * (d_1996_v0_)

                    return lambda98_

                init55_ = lambda97_(d_1972_v0_)
                nw360_ = _dafny.Array(None, 9)
                for i0_55_ in range(nw360_.length(0)):
                    nw360_[i0_55_] = init55_(i0_55_)
                d_1995_v18_ = nw360_
                index346_ = default__.safeIndex(801, (d_1995_v18_).length(0))
                (d_1995_v18_)[index346_] = (0) - ((p0)[default__.safeIndex(d_1972_v0_, len(p0))])
                index347_ = default__.safeIndex(349, (d_1990_v15_).length(0))
                index348_ = default__.safeIndex(801, (d_1995_v18_).length(0))
                rhs394_ = d_1992_v17_
                rhs395_ = (d_1972_v0_) > (d_1972_v0_)
                rhs396_ = -241
                lhs378_ = d_1990_v15_
                lhs379_ = default__.safeIndex(349, (d_1990_v15_).length(0))
                lhs380_ = globalState
                lhs381_ = d_1995_v18_
                lhs382_ = default__.safeIndex(801, (d_1995_v18_).length(0))
                lhs378_[lhs379_] = rhs394_
                lhs380_.f15 = rhs395_
                lhs381_[lhs382_] = rhs396_
                d_1998_v19_: _dafny.Seq
                d_1998_v19_ = _dafny.SeqWithoutIsStrInference([(d_1982_v7_).f33])
                (globalState).f18 = (d_1998_v19_)[default__.safeIndex((d_1982_v7_).f34, len(d_1998_v19_))]
                (globalState).f15 = (d_1982_v7_).f33
            elif True:
                (globalState).f14 = default__.fm0(d_1972_v0_, (d_1982_v7_).f33, ((d_1986_v11_)[(d_1982_v7_).f33] if ((d_1982_v7_).f33) in (d_1986_v11_) else (d_1982_v7_).f33), globalState)
                d_1999_v20_: str
                d_1999_v20_ = _dafny.CodePoint('i')
                d_1999_v20_ = d_1999_v20_
                d_2000_v21_: _dafny.Map
                d_2000_v21_ = _dafny.Map({d_1980_v5_: -52})
                d_2001_v22_: D10
                d_2001_v22_ = D10_DC32(d_1999_v20_, default__.fm6(((d_2000_v21_)[d_1980_v5_] if (d_1980_v5_) in (d_2000_v21_) else (d_1982_v7_).f34), globalState))
                (globalState).f28 = (d_2001_v22_).cf53
                d_2002_v23_: _dafny.Array
                nw361_ = _dafny.Array(_dafny.Set({}), 3)
                d_2002_v23_ = nw361_
                d_2003_v24_: _dafny.Set
                d_2003_v24_ = _dafny.Set({(d_1982_v7_).f33})
                index349_ = default__.safeIndex(773, (d_2002_v23_).length(0))
                (d_2002_v23_)[index349_] = d_2003_v24_
                d_2004_v25_: _dafny.Seq
                d_2004_v25_ = _dafny.SeqWithoutIsStrInference([(d_1982_v7_).f33, (d_1982_v7_).f33, d_1980_v5_])
                d_2005_v26_: _dafny.Map
                d_2005_v26_ = _dafny.Map({len(d_2004_v25_): (d_1982_v7_).f33})
                d_2006_v29_: _dafny.Array
                nw362_ = _dafny.Array(None, 21)
                nw362_[int(0)] = d_2005_v26_
                nw362_[int(1)] = d_2005_v26_
                nw362_[int(2)] = d_2005_v26_
                nw362_[int(3)] = default__.fm41((0) - ((d_1982_v7_).f34), (d_1982_v7_).f33, not((d_1982_v7_).f33), d_1980_v5_, globalState)
                def iife167_():
                    coll71_ = _dafny.Map()
                    compr_71_: int
                    for compr_71_ in (p0).Elements:
                        d_2007_v27_: int = compr_71_
                        if (d_2007_v27_) in (p0):
                            coll71_[(d_2007_v27_) + ((d_1982_v7_).f34)] = d_1980_v5_
                    return _dafny.Map(coll71_)
                nw362_[int(4)] = iife167_()
                
                nw362_[int(5)] = d_2005_v26_
                nw362_[int(6)] = (_dafny.Map({354: (d_1982_v7_).f33})).set((d_1982_v7_).f34, (d_1982_v7_).f33)
                nw362_[int(7)] = d_2005_v26_
                def iife168_():
                    coll72_ = _dafny.Map()
                    compr_72_: int
                    for compr_72_ in _dafny.IntegerRange(-623, 476):
                        d_2008_v28_: int = compr_72_
                        if ((-623) <= (d_2008_v28_)) and ((d_2008_v28_) < (476)):
                            coll72_[(d_2008_v28_) * ((d_1982_v7_).f34)] = (d_1982_v7_).f33
                    return _dafny.Map(coll72_)
                nw362_[int(8)] = iife168_()
                
                nw362_[int(9)] = d_2005_v26_
                nw362_[int(10)] = d_2005_v26_
                nw362_[int(11)] = (d_2005_v26_) | (d_2005_v26_)
                nw362_[int(12)] = d_2005_v26_
                nw362_[int(13)] = (d_2005_v26_) | (d_2005_v26_)
                nw362_[int(14)] = d_2005_v26_
                nw362_[int(15)] = d_2005_v26_
                nw362_[int(16)] = d_2005_v26_
                nw362_[int(17)] = d_2005_v26_
                nw362_[int(18)] = (_dafny.Map({(d_1982_v7_).f34: (d_1982_v7_).f33})) | (d_2005_v26_)
                nw362_[int(19)] = d_2005_v26_
                nw362_[int(20)] = (d_2005_v26_) | (d_2005_v26_)
                d_2006_v29_ = nw362_
                index350_ = default__.safeIndex(667, (d_2006_v29_).length(0))
                (d_2006_v29_)[index350_] = d_2005_v26_
                d_2009_v30_: C3
                nw363_ = C3()
                nw363_.ctor__((d_1982_v7_).f33, (d_1982_v7_).f34, False, _dafny.SeqWithoutIsStrInference([d_1999_v20_ for d_2010_i5_ in range(default__.abs(-792))]))
                d_2009_v30_ = nw363_
                d_2011_v31_: _dafny.Seq
                d_2011_v31_ = _dafny.SeqWithoutIsStrInference([d_2009_v30_, d_2009_v30_])
                index351_ = default__.safeIndex(773, (d_2002_v23_).length(0))
                index352_ = default__.safeIndex(667, (d_2006_v29_).length(0))
                rhs397_ = d_2003_v24_
                rhs398_ = (d_1973_v1_).set(default__.fm36(globalState), default__.abs((0) - (len((d_2000_v21_).set(not(False), len(d_2011_v31_))))))
                rhs399_ = (d_2004_v25_)[default__.safeIndex(d_1972_v0_, len(d_2004_v25_))]
                rhs400_ = (p0) + (p0)
                rhs401_ = (d_2005_v26_) | (d_2005_v26_)
                lhs383_ = d_2002_v23_
                lhs384_ = default__.safeIndex(773, (d_2002_v23_).length(0))
                lhs385_ = globalState
                lhs386_ = globalState
                lhs387_ = d_2006_v29_
                lhs388_ = default__.safeIndex(667, (d_2006_v29_).length(0))
                lhs383_[lhs384_] = rhs397_
                d_1973_v1_ = rhs398_
                lhs385_.f13 = rhs399_
                lhs386_.f7 = rhs400_
                lhs387_[lhs388_] = rhs401_
                d_2012_v32_: _dafny.Array
                nw364_ = _dafny.Array(None, 4)
                nw364_[int(0)] = (d_1982_v7_).f34
                nw364_[int(1)] = len(d_1981_v6_)
                nw364_[int(2)] = (d_1982_v7_).f34
                nw364_[int(3)] = (d_1982_v7_).f34
                d_2012_v32_ = nw364_
                d_2013_v33_: D12
                d_2013_v33_ = D12_DC38(d_2012_v32_)
                index353_ = default__.safeIndex(528, (d_1974_v2_).length(0))
                (d_1974_v2_)[index353_] = (d_2004_v25_)[default__.safeIndex((0) - ((d_1982_v7_).f34), len(d_2004_v25_))]
                d_2014_v34_: T0
                nw365_ = C4()
                nw365_.ctor__((d_1982_v7_).f33, len(d_1981_v6_))
                d_2014_v34_ = nw365_
                index354_ = default__.safeIndex(528, (d_1974_v2_).length(0))
                rhs402_ = (d_2013_v33_ if (d_1982_v7_).f33 else d_2013_v33_)
                rhs403_ = not(default__.fm1(_dafny.CodePoint('o'), len(_dafny.Map({d_1980_v5_: 673})), (d_1982_v7_).f34, not(d_1980_v5_), globalState))
                rhs404_ = 406
                rhs405_ = d_2014_v34_
                rhs406_ = d_2000_v21_
                lhs389_ = d_1974_v2_
                lhs390_ = default__.safeIndex(528, (d_1974_v2_).length(0))
                lhs391_ = globalState
                d_2013_v33_ = rhs402_
                lhs389_[lhs390_] = rhs403_
                lhs391_.f19 = rhs404_
                d_2014_v34_ = rhs405_
                d_2000_v21_ = rhs406_
        elif True:
            d_2015_v35_: C5
            nw366_ = C5()
            nw366_.ctor__((d_1982_v7_).f33, (d_1981_v6_) + (d_1981_v6_))
            d_2015_v35_ = nw366_
            nw367_ = C5()
            nw367_.ctor__((d_1982_v7_).f33, d_1981_v6_)
            d_2015_v35_ = nw367_
            r0 = len(d_1981_v6_)
            d_2016_v36_: T0
            nw368_ = C4()
            nw368_.ctor__(d_1980_v5_, d_1972_v0_)
            d_2016_v36_ = nw368_
            d_2017_v37_: _dafny.Map
            d_2017_v37_ = _dafny.Map({(d_1982_v7_).f34: d_2016_v36_})
            d_2018_v38_: _dafny.Array
            nw369_ = _dafny.Array(None, 24)
            nw369_[int(0)] = d_2016_v36_
            nw369_[int(1)] = d_2016_v36_
            nw369_[int(2)] = d_2016_v36_
            nw369_[int(3)] = d_2016_v36_
            nw369_[int(4)] = d_2016_v36_
            nw369_[int(5)] = d_2016_v36_
            nw369_[int(6)] = d_2016_v36_
            nw369_[int(7)] = d_2016_v36_
            nw369_[int(8)] = d_2016_v36_
            nw369_[int(9)] = d_2016_v36_
            nw369_[int(10)] = d_2016_v36_
            nw369_[int(11)] = ((d_2017_v37_)[(d_1982_v7_).f34] if ((d_1982_v7_).f34) in (d_2017_v37_) else d_2016_v36_)
            nw369_[int(12)] = d_2016_v36_
            nw369_[int(13)] = d_2016_v36_
            nw369_[int(14)] = d_2016_v36_
            nw369_[int(15)] = d_2016_v36_
            nw369_[int(16)] = d_2016_v36_
            nw369_[int(17)] = d_2016_v36_
            nw369_[int(18)] = d_2016_v36_
            nw369_[int(19)] = d_2016_v36_
            nw369_[int(20)] = d_2016_v36_
            nw369_[int(21)] = d_2016_v36_
            nw369_[int(22)] = d_2016_v36_
            nw369_[int(23)] = d_2016_v36_
            d_2018_v38_ = nw369_
            d_2019_v39_: _dafny.Map
            d_2019_v39_ = _dafny.Map({(d_1981_v6_) + (d_1981_v6_): d_2018_v38_})
            d_2019_v39_ = (d_2019_v39_).set(d_1981_v6_, (d_2018_v38_ if d_1980_v5_ else d_2018_v38_))
            if (not(((d_1982_v7_).f33 if True else (d_1982_v7_).f33))) and ((d_1982_v7_).f33):
                (globalState).f14 = default__.safeModuloInt(d_1972_v0_, (d_1982_v7_).f34)
                d_2020_v40_: _dafny.MultiSet
                d_2020_v40_ = _dafny.MultiSet([d_1980_v5_, (d_1982_v7_).f33, (d_1982_v7_).f33, (d_1982_v7_).f33, (d_1982_v7_).f33])
                (globalState).f28 = ((d_2020_v40_).intersection(d_2020_v40_)).isdisjoint((d_2020_v40_).set((d_1982_v7_).f33, default__.abs(d_1972_v0_)))
                d_2021_v41_: _dafny.Set
                d_2021_v41_ = _dafny.Set({(d_1982_v7_).f33})
                d_2022_v42_: _dafny.Map
                d_2022_v42_ = _dafny.Map({len(d_2021_v41_): len(d_1981_v6_)})
                d_2023_v43_: _dafny.Array
                nw370_ = _dafny.Array(None, 19)
                nw370_[int(0)] = ((d_1973_v1_) - (d_1973_v1_)).cardinality
                nw370_[int(1)] = (d_1982_v7_).f34
                nw370_[int(2)] = default__.safeDivisionInt((d_1982_v7_).f34, (d_1982_v7_).f34)
                nw370_[int(3)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kljunxavc")))
                nw370_[int(4)] = (d_1982_v7_).f34
                nw370_[int(5)] = (d_1982_v7_).f34
                nw370_[int(6)] = -367
                nw370_[int(7)] = d_1972_v0_
                nw370_[int(8)] = ((d_2022_v42_)[179] if (179) in (d_2022_v42_) else (d_1982_v7_).f34)
                nw370_[int(9)] = d_1972_v0_
                nw370_[int(10)] = (d_1982_v7_).f34
                nw370_[int(11)] = (d_1982_v7_).f34
                nw370_[int(12)] = (d_1982_v7_).f34
                nw370_[int(13)] = default__.fm36(globalState)
                nw370_[int(14)] = d_1972_v0_
                nw370_[int(15)] = (d_1982_v7_).f34
                nw370_[int(16)] = default__.safeDivisionInt((d_1982_v7_).f34, (d_1982_v7_).f34)
                nw370_[int(17)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xgc"))) if (d_1982_v7_).f33 else (d_1982_v7_).f34)
                nw370_[int(18)] = (0) - ((0) - ((d_1982_v7_).f34))
                d_2023_v43_ = nw370_
                (globalState).f17 = d_2023_v43_
                (globalState).f14 = default__.safeDivisionInt(d_1972_v0_, (d_1982_v7_).f34)
                index355_ = default__.safeIndex(527, (d_2023_v43_).length(0))
                (d_2023_v43_)[index355_] = (0) - (d_1972_v0_)
                index356_ = default__.safeIndex(902, (d_2023_v43_).length(0))
                (d_2023_v43_)[index356_] = d_1972_v0_
                index357_ = default__.safeIndex(527, (d_2023_v43_).length(0))
                index358_ = default__.safeIndex(902, (d_2023_v43_).length(0))
                rhs407_ = (d_1982_v7_).f34
                rhs408_ = 309
                lhs392_ = d_2023_v43_
                lhs393_ = default__.safeIndex(527, (d_2023_v43_).length(0))
                lhs394_ = d_2023_v43_
                lhs395_ = default__.safeIndex(902, (d_2023_v43_).length(0))
                lhs392_[lhs393_] = rhs407_
                lhs394_[lhs395_] = rhs408_
            elif True:
                d_2024_v44_: _dafny.Array
                def lambda99_(d_2025_v6_):
                    def lambda100_(d_2026_i6_):
                        return d_2025_v6_

                    return lambda100_

                init56_ = lambda99_(d_1981_v6_)
                nw371_ = _dafny.Array(None, 17)
                for i0_56_ in range(nw371_.length(0)):
                    nw371_[i0_56_] = init56_(i0_56_)
                d_2024_v44_ = nw371_
                index359_ = default__.safeIndex(419, (d_2024_v44_).length(0))
                (d_2024_v44_)[index359_] = d_1981_v6_
                index360_ = default__.safeIndex(419, (d_2024_v44_).length(0))
                (d_2024_v44_)[index360_] = (d_1981_v6_) + (d_1981_v6_)
                (globalState).f15 = True
                (globalState).f15 = (d_1982_v7_).f33
                (globalState).f20 = ((d_2024_v44_)[default__.safeIndex(419, (d_2024_v44_).length(0))]) + ((d_2024_v44_)[default__.safeIndex(419, (d_2024_v44_).length(0))])
                def iife169_():
                    coll73_ = _dafny.Map()
                    compr_73_: int
                    for compr_73_ in (d_1973_v1_).Elements:
                        d_2027_v45_: int = compr_73_
                        if (d_2027_v45_) in (d_1973_v1_):
                            def iife170_():
                                coll74_ = _dafny.Set()
                                compr_74_: int
                                for compr_74_ in (d_1973_v1_).Elements:
                                    d_2028_v46_: int = compr_74_
                                    if (d_2028_v46_) in (d_1973_v1_):
                                        coll74_ = coll74_.union(_dafny.Set([(d_2028_v46_) + (len(_dafny.Map({False: (_dafny.MultiSet([True])).cardinality})))]))
                                return _dafny.Set(coll74_)
                            coll73_[(d_2027_v45_) * ((d_1982_v7_).f34)] = _dafny.Map({(d_1982_v7_).f34: len(_dafny.Map({False: len(iife170_()
                            )}))})
                    return _dafny.Map(coll73_)
                (globalState).f1 = len(iife169_()
                )
            d_2029_v47_: _dafny.MultiSet
            d_2029_v47_ = _dafny.MultiSet([False])
            d_2030_v48_: _dafny.Map
            d_2030_v48_ = _dafny.Map({(d_1982_v7_).f34: d_2029_v47_})
            d_2031_v49_: _dafny.Seq
            d_2031_v49_ = _dafny.SeqWithoutIsStrInference([d_1981_v6_])
            d_2032_v50_: D0
            d_2032_v50_ = D0_DC2(((d_2030_v48_)[d_1972_v0_] if (d_1972_v0_) in (d_2030_v48_) else d_2029_v47_), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umiruj"))) + ((d_2031_v49_)[default__.safeIndex(d_1972_v0_, len(d_2031_v49_))]), (d_1982_v7_).f33)
            d_2032_v50_ = D0_DC2(_dafny.MultiSet([(d_1982_v7_).f33, not((d_1982_v7_).f33), (d_1982_v7_).f33]), d_1981_v6_, False)
        hi12_ = d_1972_v0_
        for d_2033_i7_ in range((d_1972_v0_) - ((0) - ((d_1982_v7_).f34)), hi12_):
            d_2034_v51_: _dafny.Array
            nw372_ = _dafny.Array(None, 12)
            d_2034_v51_ = nw372_
            d_2035_v52_: _dafny.Array
            def lambda101_(d_2036_v0_):
                def lambda102_(d_2037_i8_):
                    return (d_2037_i8_) + (d_2036_v0_)

                return lambda102_

            init57_ = lambda101_(d_1972_v0_)
            nw373_ = _dafny.Array(None, 12)
            for i0_57_ in range(nw373_.length(0)):
                nw373_[i0_57_] = init57_(i0_57_)
            d_2035_v52_ = nw373_
            d_2038_v53_: _dafny.Set
            d_2038_v53_ = _dafny.Set({d_2035_v52_, d_2035_v52_})
            d_2039_v54_: str
            d_2039_v54_ = _dafny.CodePoint('r')
            d_2040_v55_: C2
            nw374_ = C2()
            nw374_.ctor__((_dafny.MultiSet([d_1982_v7_])).cardinality, d_2038_v53_, d_1980_v5_, (d_1981_v6_).set(default__.safeIndex(len(d_1981_v6_), len(d_1981_v6_)), d_2039_v54_))
            d_2040_v55_ = nw374_
            index361_ = default__.safeIndex(130, (d_2034_v51_).length(0))
            (d_2034_v51_)[index361_] = d_2040_v55_
            index362_ = default__.safeIndex(130, (d_2034_v51_).length(0))
            (d_2034_v51_)[index362_] = d_2040_v55_
            d_2041_v56_: _dafny.Map
            d_2041_v56_ = _dafny.Map({d_1980_v5_: -916})
            d_2042_v57_: T1
            nw375_ = C3()
            nw375_.ctor__((d_1982_v7_).f33, 941, (d_2041_v56_) != (d_2041_v56_), d_1981_v6_)
            d_2042_v57_ = nw375_
            d_2042_v57_ = d_2042_v57_
            d_2043_v58_: _dafny.MultiSet
            d_2043_v58_ = _dafny.MultiSet([(d_2042_v57_).f29])
            d_2044_v59_: D11
            d_2044_v59_ = D11_DC37((d_2043_v58_).cardinality, d_1980_v5_)
            def iife171_(_pat_let48_0):
                def iife172_(d_2045_dt__update__tmp_h0_):
                    def iife173_(_pat_let49_0):
                        def iife174_(d_2046_dt__update_hcf63_h0_):
                            return D11_DC37((d_2045_dt__update__tmp_h0_).cf62, d_2046_dt__update_hcf63_h0_)
                        return iife174_(_pat_let49_0)
                    return iife173_(False)
                return iife172_(_pat_let48_0)
            (globalState).f2 = (iife171_(d_2044_v59_)).cf62
            d_2047_v60_: T0
            nw376_ = C1()
            nw376_.ctor__(False)
            d_2047_v60_ = nw376_
            d_2048_v61_: _dafny.Map
            d_2048_v61_ = _dafny.Map({d_2047_v60_: ((d_2040_v55_).f38) != ((d_2040_v55_).f38)})
            d_2048_v61_ = (d_2048_v61_).set(d_2047_v60_, d_1980_v5_)
        r0 = d_1972_v0_
        return r0

