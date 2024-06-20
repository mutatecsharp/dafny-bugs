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
        if True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwfd"))
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_0_i0_ in range(default__.abs(682))])

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        return (_dafny.CodePoint('o')) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1_i0_ in range(default__.abs(844))]))

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return D0_DC1(True, not (False) or (True), _dafny.CodePoint('t'), (len(_dafny.Set({True, False}))) - (31))

    @staticmethod
    def fm5(globalState):
        return (_dafny.Set({False, True, True, False})).ispropersubset(_dafny.Set({False, False, False}))

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2_i0_ in range(default__.abs(-952))])

    @staticmethod
    def fm13(globalState):
        return _dafny.Map({(0) - ((687) * (len(_dafny.Set({_dafny.CodePoint('k')})))): True})

    @staticmethod
    def fm15(globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkklhmuw"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_3_i0_ in range(default__.abs(794))]))) + ((D6_DC11(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rw")))).cf21)

    @staticmethod
    def fm16(globalState):
        return _dafny.SeqWithoutIsStrInference([(len(_dafny.SeqWithoutIsStrInference([834, len(_dafny.Map({False: -435}))]))) <= ((0) - (len(_dafny.Map({False: _dafny.CodePoint('j')}))))])

    @staticmethod
    def fm17(p0, globalState):
        return ((_dafny.MultiSet([_dafny.CodePoint('s'), _dafny.CodePoint('u'), _dafny.CodePoint('u')]) if False else _dafny.MultiSet([_dafny.CodePoint('o')]))) | ((_dafny.MultiSet([_dafny.CodePoint('k')])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g')]))))

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(104, 923):
                d_4_v0_: int = compr_0_
                if ((104) <= (d_4_v0_)) and ((d_4_v0_) < (923)):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeModuloInt(d_4_v0_, 703)]))
            return _dafny.Set(coll0_)
        return ((_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imlqi"))), len(_dafny.Map({False: not(True)})), 881, (_dafny.MultiSet([True])).cardinality, len(_dafny.SeqWithoutIsStrInference([-937]))})).intersection(iife0_()
        )) - (_dafny.Set({(_dafny.MultiSet([not(True)])).cardinality}))

    @staticmethod
    def fm19(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([(_dafny.MultiSet([True, not(True), False, True])).cardinality])).intersection(_dafny.MultiSet([500]))) | (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([-557 for d_5_i0_ in range(default__.abs(973))]) if True else _dafny.SeqWithoutIsStrInference([len(_dafny.Map({996: True}))]))))

    @staticmethod
    def fm20(globalState):
        return (_dafny.Map({_dafny.Map({False: True}): True}))

    @staticmethod
    def fm21(globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ay"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "poxqgk"))): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xqxl"))) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_6_i0_ in range(default__.abs(-618))]))})

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_7_i0_ in range(default__.abs(878))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r')]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d')]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_8_i1_ in range(default__.abs(995))])])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('w')]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('q')])]))

    @staticmethod
    def fm23(globalState):
        return _dafny.CodePoint('j')

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        if True:
            return D0_DC0(True)
        elif True:
            return D0_DC0(True)

    @staticmethod
    def fm26(p0, p1, p2, p3, globalState):
        source0_ = D7_DC14(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvtm"))) for d_9_i0_ in range(default__.abs(813))]))
        if source0_.is_DC15:
            d_10___mcc_h0_ = source0_.cf27
            d_11___mcc_h1_ = source0_.cf28
            d_12_cf28_ = d_11___mcc_h1_
            d_13_cf27_ = d_10___mcc_h0_
            return _dafny.Map({True: True})
        elif True:
            d_14___mcc_h2_ = source0_.cf26
            d_15_cf26_ = d_14___mcc_h2_
            return _dafny.Map({not(not(True)): not(False)})

    @staticmethod
    def fm27(p0, p1, globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khmjbaxct"))): (696) * ((_dafny.MultiSet([_dafny.CodePoint('b')])).cardinality)})

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxkrw"))

    @staticmethod
    def fm29(globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(801, 233):
                d_16_v0_: int = compr_1_
                if ((801) <= (d_16_v0_)) and ((d_16_v0_) < (233)):
                    coll1_[default__.safeDivisionInt(d_16_v0_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality)] = True
            return _dafny.Map(coll1_)
        return D5_DC10(not((not(False)) and (True)), 117, _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbplxx")): iife1_()
}), not(False))

    @staticmethod
    def fm30(p0, globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "acqcse"))])) - (_dafny.MultiSet([(D6_DC11(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_17_i0_ in range(default__.abs(-477))]))).cf21, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oc")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_18_i1_ in range(default__.abs(336))])]))) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "teyeal")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcehpgj"))]))

    @staticmethod
    def fm31(p0, p1, p2, p3, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in (_dafny.MultiSet([883])).Elements:
                d_19_v0_: int = compr_2_
                if (d_19_v0_) in (_dafny.MultiSet([883])):
                    coll2_ = coll2_.union(_dafny.Set([(d_19_v0_) * (len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([True])))})))]))
            return _dafny.Set(coll2_)
        return (_dafny.Set({382})) | ((_dafny.Set({962, (0) - (-593)})).intersection(iife2_()
        ))

    @staticmethod
    def fm32(globalState):
        return (_dafny.SeqWithoutIsStrInference([D7_DC14(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([-334])): 564}))]))])) + ((_dafny.SeqWithoutIsStrInference([D7_DC14(_dafny.SeqWithoutIsStrInference([913 for d_20_i0_ in range(default__.abs(-49))]))])) + (_dafny.SeqWithoutIsStrInference([D7_DC14(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: -550})), len(_dafny.SeqWithoutIsStrInference([False, True]))])), D7_DC14(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulsmcuwk")))])), D7_DC14(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.Set({False, True}) for d_21_i1_ in range(default__.abs(150))])): 80})), 561]))).cardinality, 374, -739, len(_dafny.SeqWithoutIsStrInference([True, False]))]))])))

    @staticmethod
    def fm33(p0, p1, p2, p3, globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(577, -231):
                d_25_v0_: int = compr_3_
                if ((577) <= (d_25_v0_)) and ((d_25_v0_) < (-231)):
                    coll3_ = coll3_.union(_dafny.Set([(d_25_v0_) + (8)]))
            return _dafny.Set(coll3_)
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([D7_DC14(_dafny.SeqWithoutIsStrInference([25]))])) + (_dafny.SeqWithoutIsStrInference([D7_DC14(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_22_i0_ in range(default__.abs(-788))]))])), D7_DC14(_dafny.SeqWithoutIsStrInference([621, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_23_i1_ in range(default__.abs(909))]))])), D7_DC14(_dafny.SeqWithoutIsStrInference([782, len(_dafny.Set({len(_dafny.Set({750}))}))]))]))) + ((_dafny.SeqWithoutIsStrInference([D7_DC14(_dafny.SeqWithoutIsStrInference([134])), D7_DC14(_dafny.SeqWithoutIsStrInference([300])), D7_DC14(_dafny.SeqWithoutIsStrInference([-30, 30, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpgtqpqa"))), 159]))])) + (_dafny.SeqWithoutIsStrInference([D7_DC14(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppysr"))), len(_dafny.Set({False, True})), 636, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvfexkl"))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_24_i2_ in range(default__.abs(434))]))])), D7_DC14(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qmdcartnm")))])), D7_DC14(_dafny.SeqWithoutIsStrInference([len(iife3_()
) for d_26_i3_ in range(default__.abs(-803))])), D7_DC14(_dafny.SeqWithoutIsStrInference([-929])), D7_DC14(_dafny.SeqWithoutIsStrInference([612]))]))))

    @staticmethod
    def fm34(p0, globalState):
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ckmfb")): False}) if not(True) else _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfwhyaky")): False}))) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_27_i0_ in range(default__.abs(237))]): False})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvc")): not(True)})))

    @staticmethod
    def fm35(p0, p1, p2, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(-548, 152):
                d_28_v0_: int = compr_4_
                if ((-548) <= (d_28_v0_)) and ((d_28_v0_) < (152)):
                    coll4_[(d_28_v0_) - (896)] = (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxd")))])).cardinality
            return _dafny.Map(coll4_)
        return ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False]))])) - ((_dafny.MultiSet([257, len(_dafny.Set({False})), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False, True]))])).cardinality, 482])))) | ((_dafny.MultiSet([len(iife4_()
        )])).intersection(_dafny.MultiSet([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([395]))).cardinality])))

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt((0) - ((_dafny.MultiSet([_dafny.CodePoint('y'), _dafny.CodePoint('k'), _dafny.CodePoint('j'), _dafny.CodePoint('y')])).cardinality), 861)])

    @staticmethod
    def fm37(p0, globalState):
        return _dafny.CodePoint('x')

    @staticmethod
    def fm38(p0, globalState):
        return (_dafny.MultiSet([True])).intersection(_dafny.MultiSet([False, False, not(False)]))

    @staticmethod
    def fm39(globalState):
        return (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([True, False, False]))): _dafny.Map({False: not(True)})})) | (_dafny.Map({(_dafny.MultiSet([_dafny.CodePoint('w'), _dafny.CodePoint('j')])).cardinality: _dafny.Map({False: False})}))

    @staticmethod
    def fm40(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([not(True), True]), _dafny.MultiSet([True])])) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True), True])), _dafny.MultiSet([False, True, True, True])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False])])))

    @staticmethod
    def fm41(p0, p1, p2, p3, globalState):
        return (D14_DC26(D0_DC0(not(False)), _dafny.Set({False, True, False, False, False}))).cf44

    @staticmethod
    def fm42(p0, p1, p2, globalState):
        return D6_DC11((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxwyhsabm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gojgfnrgd"))))

    @staticmethod
    def fm43(globalState):
        return ((_dafny.Map({_dafny.CodePoint('x'): _dafny.MultiSet([212])})) | (_dafny.Map({_dafny.CodePoint('p'): _dafny.MultiSet([383, len(_dafny.Map({_dafny.CodePoint('k'): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lu")))}))])}))) | (_dafny.Map({_dafny.CodePoint('d'): _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_29_i0_ in range(default__.abs(348))])), -327])}))

    @staticmethod
    def fm44(p0, p1, p2, p3, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: _dafny.Seq
            for compr_5_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([381]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_30_i0_ in range(default__.abs(247))])})).Elements:
                d_31_v0_: _dafny.Seq = compr_5_
                if (d_31_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([381]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_30_i0_ in range(default__.abs(247))])})):
                    coll5_[d_31_v0_] = False
            return _dafny.Map(coll5_)
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([-445, -810, 358]): True})) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([777]): not(True)})) | (iife5_()
        ))

    @staticmethod
    def fm45(p0, p1, p2, p3, globalState):
        return D11_DC22(not(not (False) or (True)))

    @staticmethod
    def fm46(p0, globalState):
        return D11_DC21((_dafny.CodePoint('k') if True else _dafny.CodePoint('e')), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pldd"))), -615)

    @staticmethod
    def fm47(p0, p1, p2, globalState):
        return (_dafny.Map({False: -131})) | ((_dafny.Map({False: (_dafny.MultiSet([False])).cardinality})) | (_dafny.Map({not(not(True)): 391})))

    @staticmethod
    def fm48(p0, globalState):
        return _dafny.Set({(_dafny.MultiSet([788])).cardinality})

    @staticmethod
    def fm49(p0, p1, p2, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(249, -663):
                d_32_v0_: int = compr_6_
                if ((249) <= (d_32_v0_)) and ((d_32_v0_) < (-663)):
                    coll6_[(d_32_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahorcb"))))] = False
            return _dafny.Map(coll6_)
        return (_dafny.Map({_dafny.CodePoint('p'): True})) | (_dafny.Map({_dafny.CodePoint('o'): (D5_DC10(False, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q'), _dafny.CodePoint('b')]))).cardinality, _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "afr")): iife6_()
}), False)).cf17}))

    @staticmethod
    def fm50(p0, p1, p2, p3, globalState):
        source1_ = D14_DC27(not(True))
        if source1_.is_DC26:
            d_33___mcc_h0_ = source1_.cf43
            d_34___mcc_h1_ = source1_.cf44
            d_35_cf44_ = d_34___mcc_h1_
            d_36_cf43_ = d_33___mcc_h0_
            return D7_DC15(85, _dafny.CodePoint('d'))
        elif source1_.is_DC27:
            d_37___mcc_h2_ = source1_.cf45
            d_38_cf45_ = d_37___mcc_h2_
            return D7_DC15(len(_dafny.Set({d_38_cf45_, d_38_cf45_, d_38_cf45_, d_38_cf45_, d_38_cf45_})), _dafny.CodePoint('b'))
        elif source1_.is_DC28:
            return D7_DC15(404, _dafny.CodePoint('n'))
        elif True:
            d_39___mcc_h3_ = source1_.cf42
            d_40_cf42_ = d_39___mcc_h3_
            return D7_DC15(717, _dafny.CodePoint('m'))

    @staticmethod
    def fm51(p0, p1, p2, globalState):
        return D14_DC28()

    @staticmethod
    def fm52(p0, p1, p2, p3, globalState):
        return (_dafny.Set({D6_DC11((D6_DC11(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_41_i0_ in range(default__.abs(662))]))).cf21), D6_DC11(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqncemshr")))})).intersection(_dafny.Set({D6_DC11(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_42_i1_ in range(default__.abs(718))])), D6_DC11(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghyhpuqt")))}))

    @staticmethod
    def fm53(p0, p1, globalState):
        return D14_DC26(D0_DC0(True), _dafny.Set({False}))

    @staticmethod
    def fm54(p0, p1, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(845, 377):
                d_44_v0_: int = compr_7_
                if ((845) <= (d_44_v0_)) and ((d_44_v0_) < (377)):
                    coll7_[default__.safeModuloInt(d_44_v0_, 674)] = False
            return _dafny.Map(coll7_)
        return len((_dafny.SeqWithoutIsStrInference([-743 for d_43_i0_ in range(default__.abs(-926))])) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(iife7_()
        )})), len((D24_DC48(_dafny.Map({True: len(_dafny.Set({-110}))}))).cf67), (_dafny.MultiSet([(_dafny.MultiSet([True])).cardinality, 893])).cardinality, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True]))]))).cardinality]))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({426: not(not(False))})) for d_45_i1_ in range(default__.abs(145))]))))

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        return (_dafny.Map({_dafny.MultiSet([not(False)]): not(True)})) | (_dafny.Map({_dafny.MultiSet([True]): True}))

    @staticmethod
    def fm56(p0, globalState):
        return (_dafny.MultiSet([_dafny.MultiSet([747]), _dafny.MultiSet([-652])])) | ((_dafny.MultiSet([_dafny.MultiSet([745, 760])])) | (_dafny.MultiSet([_dafny.MultiSet([-439])])))

    @staticmethod
    def fm57(p0, globalState):
        return (_dafny.Map({False: False})) | (_dafny.Map({not(True): True}))

    @staticmethod
    def m0(p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_46_v0_: int
        d_46_v0_ = -220
        d_47_v1_: C2
        nw0_ = C2()
        nw0_.ctor__(d_46_v0_)
        d_47_v1_ = nw0_
        d_48_v2_: bool
        d_48_v2_ = True
        d_48_v2_ = default__.fm5(globalState)
        d_49_v3_: T0
        nw1_ = C2()
        nw1_.ctor__(d_46_v0_)
        d_49_v3_ = nw1_
        d_50_v4_: T0
        d_50_v4_ = d_49_v3_
        source2_ = d_50_v4_
        d_51___mcc_h0_ = source2_
        d_52_cf53_ = d_51___mcc_h0_
        d_53_v5_: _dafny.Seq
        d_53_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxgtgy"))
        d_54_v6_: _dafny.Array
        nw2_ = _dafny.Array(None, 25)
        nw2_[int(0)] = (d_50_v4_)
        nw2_[int(1)] = d_52_cf53_
        nw2_[int(2)] = d_49_v3_
        nw2_[int(3)] = d_52_cf53_
        nw2_[int(4)] = d_49_v3_
        nw2_[int(5)] = d_52_cf53_
        nw2_[int(6)] = d_49_v3_
        nw2_[int(7)] = d_52_cf53_
        nw2_[int(8)] = d_52_cf53_
        nw2_[int(9)] = d_49_v3_
        nw2_[int(10)] = d_52_cf53_
        nw2_[int(11)] = d_52_cf53_
        nw2_[int(12)] = d_49_v3_
        nw2_[int(13)] = d_49_v3_
        nw2_[int(14)] = d_52_cf53_
        nw2_[int(15)] = d_49_v3_
        nw2_[int(16)] = d_49_v3_
        nw2_[int(17)] = d_52_cf53_
        nw2_[int(18)] = d_52_cf53_
        nw2_[int(19)] = d_52_cf53_
        nw2_[int(20)] = d_49_v3_
        nw2_[int(21)] = d_49_v3_
        nw2_[int(22)] = d_52_cf53_
        nw2_[int(23)] = d_49_v3_
        nw2_[int(24)] = d_49_v3_
        d_54_v6_ = nw2_
        d_55_v7_: _dafny.Array
        def lambda0_(d_56_v0_):
            def lambda1_(d_57_i0_):
                return (d_57_i0_) - (d_56_v0_)

            return lambda1_

        init0_ = lambda0_(d_46_v0_)
        nw3_ = _dafny.Array(None, 9)
        for i0_0_ in range(nw3_.length(0)):
            nw3_[i0_0_] = init0_(i0_0_)
        d_55_v7_ = nw3_
        d_58_v8_: _dafny.Map
        d_58_v8_ = _dafny.Map({d_55_v7_: d_54_v6_})
        d_59_v9_: _dafny.Array
        nw4_ = _dafny.Array(None, 24)
        nw4_[int(0)] = d_54_v6_
        nw4_[int(1)] = d_54_v6_
        nw4_[int(2)] = d_54_v6_
        nw4_[int(3)] = d_54_v6_
        nw4_[int(4)] = d_54_v6_
        nw4_[int(5)] = d_54_v6_
        nw4_[int(6)] = d_54_v6_
        nw4_[int(7)] = d_54_v6_
        nw4_[int(8)] = d_54_v6_
        nw4_[int(9)] = d_54_v6_
        nw4_[int(10)] = d_54_v6_
        nw4_[int(11)] = d_54_v6_
        nw4_[int(12)] = d_54_v6_
        nw4_[int(13)] = d_54_v6_
        nw4_[int(14)] = d_54_v6_
        nw4_[int(15)] = d_54_v6_
        nw4_[int(16)] = d_54_v6_
        nw4_[int(17)] = d_54_v6_
        nw4_[int(18)] = ((d_58_v8_)[d_55_v7_] if (d_55_v7_) in (d_58_v8_) else d_54_v6_)
        nw4_[int(19)] = d_54_v6_
        nw4_[int(20)] = d_54_v6_
        nw4_[int(21)] = d_54_v6_
        nw4_[int(22)] = d_54_v6_
        nw4_[int(23)] = d_54_v6_
        d_59_v9_ = nw4_
        d_60_v10_: _dafny.Map
        d_60_v10_ = _dafny.Map({d_53_v5_: d_59_v9_})
        d_60_v10_ = (d_60_v10_).set(d_53_v5_, d_59_v9_)
        if not(d_48_v2_):
            d_61_v11_: _dafny.Array
            nw5_ = _dafny.Array(False, 16)
            d_61_v11_ = nw5_
            d_62_v12_: _dafny.Seq
            d_62_v12_ = _dafny.SeqWithoutIsStrInference([d_61_v11_, d_61_v11_])
            d_63_v13_: _dafny.MultiSet
            d_63_v13_ = _dafny.MultiSet([-971, d_46_v0_])
            d_64_v14_: _dafny.Set
            d_64_v14_ = _dafny.Set({(d_62_v12_) != (d_62_v12_), not((d_63_v13_).isdisjoint(d_63_v13_)), (d_47_v1_).fm8(d_53_v5_, p0, globalState), p1})
            d_65_v15_: _dafny.Array
            nw6_ = _dafny.Array(None, 13)
            nw6_[int(0)] = d_50_v4_
            nw6_[int(1)] = d_50_v4_
            nw6_[int(2)] = d_50_v4_
            nw6_[int(3)] = d_50_v4_
            nw6_[int(4)] = d_50_v4_
            nw6_[int(5)] = d_50_v4_
            nw6_[int(6)] = d_50_v4_
            nw6_[int(7)] = d_50_v4_
            nw6_[int(8)] = d_50_v4_
            nw6_[int(9)] = d_52_cf53_
            nw6_[int(10)] = d_50_v4_
            nw6_[int(11)] = d_50_v4_
            nw6_[int(12)] = d_50_v4_
            d_65_v15_ = nw6_
            rhs0_ = d_46_v0_
            rhs1_ = d_64_v14_
            rhs2_ = d_65_v15_
            d_46_v0_ = rhs0_
            d_64_v14_ = rhs1_
            d_65_v15_ = rhs2_
            d_66_v16_: _dafny.Map
            d_66_v16_ = _dafny.Map({len(d_53_v5_): not(d_48_v2_)})
            d_67_v17_: _dafny.Map
            d_67_v17_ = _dafny.Map({d_53_v5_: d_66_v16_})
            d_68_v18_: D5
            d_68_v18_ = D5_DC10(p1, d_46_v0_, d_67_v17_, p0)
            d_69_v19_: _dafny.Map
            d_69_v19_ = _dafny.Map({True: d_68_v18_})
            d_69_v19_ = (d_69_v19_ if p0 else d_69_v19_)
            index0_ = default__.safeIndex(278, (d_61_v11_).length(0))
            (d_61_v11_)[index0_] = p1
            index1_ = default__.safeIndex(278, (d_61_v11_).length(0))
            (d_61_v11_)[index1_] = p0
            d_70_v20_: C7
            nw7_ = C7()
            nw7_.ctor__()
            d_70_v20_ = nw7_
            d_71_v21_: D20
            d_71_v21_ = D20_DC38(d_70_v20_)
            d_70_v20_ = (d_71_v21_).cf57
            d_72_v22_: D11
            d_72_v22_ = D11_DC22(p1)
            d_73_v23_: _dafny.Map
            d_73_v23_ = _dafny.Map({d_72_v22_: d_46_v0_})
            d_74_v24_: _dafny.Set
            d_74_v24_ = _dafny.Set({d_73_v23_, d_73_v23_, d_73_v23_})
            index2_ = default__.safeIndex(393, (d_55_v7_).length(0))
            (d_55_v7_)[index2_] = len(d_74_v24_)
            d_75_v25_: _dafny.Map
            d_75_v25_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('c'): (d_61_v11_)[default__.safeIndex(278, (d_61_v11_).length(0))]}) for d_76_i1_ in range(default__.abs(-954))]): (d_49_v3_).fm1((d_61_v11_)[default__.safeIndex(278, (d_61_v11_).length(0))], globalState)})
            d_77_v26_: str
            d_77_v26_ = _dafny.CodePoint('l')
            d_78_v27_: _dafny.Map
            d_78_v27_ = _dafny.Map({d_77_v26_: p1})
            d_79_v28_: _dafny.Seq
            d_79_v28_ = _dafny.SeqWithoutIsStrInference([p0])
            d_80_v29_: _dafny.Seq
            d_80_v29_ = _dafny.SeqWithoutIsStrInference([d_78_v27_, default__.fm49(d_79_v28_, not(p0), d_46_v0_, globalState)])
            index3_ = default__.safeIndex(393, (d_55_v7_).length(0))
            (d_55_v7_)[index3_] = ((d_75_v25_)[d_80_v29_] if (d_80_v29_) in (d_75_v25_) else d_46_v0_)
        elif True:
            d_46_v0_ = (0) - (d_46_v0_)
            d_81_v30_: C3
            nw8_ = C3()
            nw8_.ctor__(((0) - (d_46_v0_)) + (d_46_v0_))
            d_81_v30_ = nw8_
            d_82_v31_: C4
            nw9_ = C4()
            nw9_.ctor__(d_46_v0_)
            d_82_v31_ = nw9_
            d_83_v32_: str
            d_83_v32_ = _dafny.CodePoint('u')
            d_84_v33_: _dafny.Array
            nw10_ = _dafny.Array(None, 1)
            nw10_[int(0)] = d_83_v32_
            d_84_v33_ = nw10_
            index4_ = default__.safeIndex(485, (d_84_v33_).length(0))
            (d_84_v33_)[index4_] = d_83_v32_
            index5_ = default__.safeIndex(485, (d_84_v33_).length(0))
            (d_84_v33_)[index5_] = d_83_v32_
            d_85_v34_: _dafny.Array
            nw11_ = _dafny.Array(False, 16)
            d_85_v34_ = nw11_
            d_86_v35_: _dafny.Map
            d_86_v35_ = _dafny.Map({p0: d_46_v0_})
            d_87_v36_: _dafny.MultiSet
            d_87_v36_ = _dafny.MultiSet([d_48_v2_])
            d_88_v37_: _dafny.Map
            d_88_v37_ = _dafny.Map({d_48_v2_: d_47_v1_})
            d_89_v38_: _dafny.Map
            d_89_v38_ = _dafny.Map({d_46_v0_: d_46_v0_})
            d_90_v39_: _dafny.MultiSet
            d_90_v39_ = _dafny.MultiSet([d_89_v38_])
            nw12_ = _dafny.Array(None, 18)
            nw12_[int(0)] = p0
            nw12_[int(1)] = (d_86_v35_) == (d_86_v35_)
            nw12_[int(2)] = p0
            nw12_[int(3)] = not((d_87_v36_).isdisjoint(d_87_v36_))
            nw12_[int(4)] = p0
            nw12_[int(5)] = (p1) in (d_88_v37_)
            nw12_[int(6)] = (d_48_v2_) == (d_48_v2_)
            nw12_[int(7)] = p0
            nw12_[int(8)] = d_48_v2_
            nw12_[int(9)] = (p0 if p1 else d_48_v2_)
            nw12_[int(10)] = (d_48_v2_ if d_48_v2_ else p1)
            nw12_[int(11)] = d_48_v2_
            nw12_[int(12)] = p1
            nw12_[int(13)] = default__.fm3(d_48_v2_, len(_dafny.SeqWithoutIsStrInference([d_81_v30_])), d_46_v0_, globalState)
            nw12_[int(14)] = p0
            nw12_[int(15)] = p1
            nw12_[int(16)] = (d_90_v39_).isdisjoint(d_90_v39_)
            nw12_[int(17)] = (d_87_v36_).issubset(d_87_v36_)
            d_85_v34_ = nw12_
        d_91_v40_: _dafny.Map
        d_91_v40_ = _dafny.Map({True: p1})
        d_92_v41_: _dafny.Map
        d_92_v41_ = d_91_v40_
        d_93_v42_: _dafny.Seq
        d_93_v42_ = _dafny.SeqWithoutIsStrInference([d_46_v0_])
        d_94_v43_: _dafny.MultiSet
        d_94_v43_ = _dafny.MultiSet([len((d_93_v42_).set(default__.safeIndex(len(d_53_v5_), len(d_93_v42_)), d_46_v0_))])
        d_95_v44_: _dafny.Seq
        d_95_v44_ = _dafny.SeqWithoutIsStrInference([(d_91_v40_).set(not(True), p1), d_91_v40_])
        d_96_v45_: _dafny.Map
        d_96_v45_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p1]): p0})
        d_97_v46_: _dafny.Array
        nw13_ = _dafny.Array(None, 25)
        nw13_[int(0)] = d_92_v41_
        nw13_[int(1)] = d_92_v41_
        nw13_[int(2)] = d_92_v41_
        nw13_[int(3)] = d_92_v41_
        nw13_[int(4)] = d_91_v40_
        nw13_[int(5)] = d_92_v41_
        nw13_[int(6)] = d_92_v41_
        nw13_[int(7)] = d_91_v40_
        nw13_[int(8)] = d_92_v41_
        nw13_[int(9)] = d_92_v41_
        nw13_[int(10)] = d_92_v41_
        nw13_[int(11)] = d_92_v41_
        nw13_[int(12)] = d_92_v41_
        nw13_[int(13)] = d_92_v41_
        nw13_[int(14)] = d_92_v41_
        nw13_[int(15)] = default__.fm57(d_94_v43_, globalState)
        nw13_[int(16)] = default__.fm57(d_94_v43_, globalState)
        nw13_[int(17)] = d_92_v41_
        nw13_[int(18)] = (d_95_v44_)[default__.safeIndex(d_46_v0_, len(d_95_v44_))]
        nw13_[int(19)] = _dafny.Map({((d_96_v45_)[_dafny.SeqWithoutIsStrInference([d_48_v2_, p0])] if (_dafny.SeqWithoutIsStrInference([d_48_v2_, p0])) in (d_96_v45_) else not(not(d_48_v2_))): p1})
        nw13_[int(20)] = d_92_v41_
        nw13_[int(21)] = d_91_v40_
        nw13_[int(22)] = d_92_v41_
        nw13_[int(23)] = d_92_v41_
        nw13_[int(24)] = d_92_v41_
        d_97_v46_ = nw13_
        index6_ = default__.safeIndex(336, (d_97_v46_).length(0))
        (d_97_v46_)[index6_] = d_92_v41_
        d_98_v47_: C7
        nw14_ = C7()
        nw14_.ctor__()
        d_98_v47_ = nw14_
        d_99_v48_: C1
        nw15_ = C1()
        nw15_.ctor__(d_53_v5_)
        d_99_v48_ = nw15_
        d_100_v49_: _dafny.Seq
        d_100_v49_ = _dafny.SeqWithoutIsStrInference([d_99_v48_])
        d_101_v50_: _dafny.Set
        d_101_v50_ = _dafny.Set({p0, p1, not(((d_100_v49_)[default__.safeIndex((d_47_v1_).fm1(not(p0), globalState), len(d_100_v49_))]) in (d_100_v49_)), p1})
        index7_ = default__.safeIndex(336, (d_97_v46_).length(0))
        rhs3_ = default__.fm57(d_94_v43_, globalState)
        rhs4_ = d_46_v0_
        rhs5_ = d_98_v47_
        rhs6_ = (d_101_v50_) - ((default__.fm41(d_93_v42_, p0, p1, d_48_v2_, globalState)) | (_dafny.Set({not(d_48_v2_)})))
        lhs0_ = d_97_v46_
        lhs1_ = default__.safeIndex(336, (d_97_v46_).length(0))
        lhs0_[lhs1_] = rhs3_
        d_46_v0_ = rhs4_
        d_98_v47_ = rhs5_
        d_101_v50_ = rhs6_
        d_46_v0_ = (0) - (default__.safeModuloInt(default__.safeDivisionInt(d_46_v0_, -418), d_46_v0_))
        d_102_v51_: T2
        nw16_ = C3()
        nw16_.ctor__(735)
        d_102_v51_ = nw16_
        d_103_v52_: T2
        d_103_v52_ = d_102_v51_
        source3_ = d_103_v52_
        d_104___mcc_h1_ = source3_
        d_105_cf33_ = d_104___mcc_h1_
        d_106_v53_: C6
        nw17_ = C6()
        nw17_.ctor__()
        d_106_v53_ = nw17_
        d_107_v54_: _dafny.Array
        nw18_ = _dafny.Array(None, 6)
        nw18_[int(0)] = (d_102_v51_.f2) + (d_102_v51_.f2)
        nw18_[int(1)] = d_102_v51_.f2
        nw18_[int(2)] = (d_46_v0_) * (d_105_cf33_.f2)
        nw18_[int(3)] = (997) - (-856)
        nw18_[int(4)] = 721
        nw18_[int(5)] = d_46_v0_
        d_107_v54_ = nw18_
        index8_ = default__.safeIndex(899, (d_107_v54_).length(0))
        (d_107_v54_)[index8_] = (d_102_v51_.f2) - (d_102_v51_.f2)
        index9_ = default__.safeIndex(899, (d_107_v54_).length(0))
        (d_107_v54_)[index9_] = d_105_cf33_.f2
        d_108_v55_: _dafny.Array
        def lambda2_(d_109_i2_):
            return True

        init1_ = lambda2_
        nw19_ = _dafny.Array(None, 1)
        for i0_1_ in range(nw19_.length(0)):
            nw19_[i0_1_] = init1_(i0_1_)
        d_108_v55_ = nw19_
        index10_ = default__.safeIndex(108, (d_108_v55_).length(0))
        (d_108_v55_)[index10_] = p0
        index11_ = default__.safeIndex(108, (d_108_v55_).length(0))
        (d_108_v55_)[index11_] = p0
        index12_ = default__.safeIndex(108, (d_108_v55_).length(0))
        (d_108_v55_)[index12_] = p0
        d_110_v56_: _dafny.Array
        def lambda3_(d_111_p0_):
            def lambda4_(d_112_i3_):
                return d_111_p0_

            return lambda4_

        init2_ = lambda3_(p0)
        nw20_ = _dafny.Array(None, 21)
        for i0_2_ in range(nw20_.length(0)):
            nw20_[i0_2_] = init2_(i0_2_)
        d_110_v56_ = nw20_
        pat_let_tv0_ = d_110_v56_
        pat_let_tv1_ = d_110_v56_
        def iife9_(_pat_let1_0):
            def iife10_(d_113_dt__update__tmp_h2_):
                def iife11_(_pat_let2_0):
                    def iife12_(d_114_dt__update_hcf35_h0_):
                        return D11_DC20(d_114_dt__update_hcf35_h0_)
                    return iife12_(_pat_let2_0)
                return iife11_(pat_let_tv0_)
            return iife10_(_pat_let1_0)
        def iife8_(_pat_let0_0):
            def iife13_(d_115_dt__update__tmp_h3_):
                def iife14_(_pat_let3_0):
                    def iife15_(d_116_dt__update_hcf35_h1_):
                        return D11_DC20(d_116_dt__update_hcf35_h1_)
                    return iife15_(_pat_let3_0)
                return iife14_(pat_let_tv1_)
            return iife13_(_pat_let0_0)
        source4_ = iife8_(iife9_(D11_DC20(d_110_v56_)))
        if source4_.is_DC21:
            d_117___mcc_h2_ = source4_.cf36
            d_118___mcc_h3_ = source4_.cf37
            d_119___mcc_h4_ = source4_.cf38
            d_120_cf38_ = d_119___mcc_h4_
            d_121_cf37_ = d_118___mcc_h3_
            d_122_cf36_ = d_117___mcc_h2_
            d_48_v2_ = p0
            nw21_ = C3()
            nw21_.ctor__(d_121_cf37_)
            d_102_v51_ = nw21_
            d_123_v57_: _dafny.Seq
            d_123_v57_ = _dafny.SeqWithoutIsStrInference([d_102_v51_.f2])
            d_124_v58_: _dafny.Seq
            d_124_v58_ = _dafny.SeqWithoutIsStrInference([d_48_v2_, d_48_v2_])
            d_125_v59_: _dafny.Map
            d_125_v59_ = _dafny.Map({len(d_123_v57_): d_124_v58_})
            d_126_v60_: _dafny.Map
            d_126_v60_ = _dafny.Map({d_102_v51_.f2: d_121_cf37_})
            d_127_v61_: _dafny.MultiSet
            d_127_v61_ = _dafny.MultiSet([(0) - ((0) - (d_121_cf37_)), d_102_v51_.f2, d_46_v0_, len(d_125_v59_), (0) - (((d_126_v60_)[d_102_v51_.f2] if (d_102_v51_.f2) in (d_126_v60_) else d_102_v51_.f2))])
            d_128_v62_: _dafny.Seq
            d_128_v62_ = _dafny.SeqWithoutIsStrInference([d_127_v61_, (_dafny.MultiSet([(0) - (d_121_cf37_)])).intersection(d_127_v61_), d_127_v61_])
            d_128_v62_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_121_cf37_, d_120_cf38_]) for d_129_i4_ in range(default__.abs(722))])
            (d_102_v51_).f2 = (((d_47_v1_).fm1(d_48_v2_, globalState)) - (d_46_v0_)) + (default__.safeModuloInt(d_120_cf38_, ((d_127_v61_)[d_102_v51_.f2] if (d_102_v51_.f2) in (d_127_v61_) else d_120_cf38_)))
        elif source4_.is_DC22:
            d_130___mcc_h5_ = source4_.cf39
            d_131_cf39_ = d_130___mcc_h5_
            if p0:
                d_48_v2_ = not(p0)
                d_131_cf39_ = not(default__.fm5(globalState))
                d_48_v2_ = p1
                d_132_v63_: _dafny.Array
                nw22_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_132_v63_ = nw22_
                d_133_v64_: _dafny.Array
                nw23_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_133_v64_ = nw23_
                index13_ = default__.safeIndex(219, (d_132_v63_).length(0))
                (d_132_v63_)[index13_] = d_133_v64_
                index14_ = default__.safeIndex(219, (d_132_v63_).length(0))
                nw24_ = _dafny.Array(_dafny.Array(None, 0), 2)
                (d_132_v63_)[index14_] = nw24_
                (d_102_v51_).f2 = (0) - (d_102_v51_.f2)
            elif True:
                (d_102_v51_).f2 = ((d_102_v51_.f2) * (d_102_v51_.f2)) + (d_102_v51_.f2)
                d_48_v2_ = p0
                (d_102_v51_).f2 = d_46_v0_
                index15_ = default__.safeIndex(243, (d_110_v56_).length(0))
                (d_110_v56_)[index15_] = d_48_v2_
                d_134_v65_: _dafny.Seq
                d_134_v65_ = _dafny.SeqWithoutIsStrInference([d_131_cf39_, (d_48_v2_ if p0 else d_48_v2_), not (True) or (p1)])
                d_135_v66_: _dafny.Seq
                d_135_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pekaqfyha"))
                index16_ = default__.safeIndex(243, (d_110_v56_).length(0))
                (d_110_v56_)[index16_] = (d_134_v65_)[default__.safeIndex((default__.fm54(d_131_cf39_, d_135_v66_, globalState)) + (d_46_v0_), len(d_134_v65_))]
                d_136_v67_: C6
                nw25_ = C6()
                nw25_.ctor__()
                d_136_v67_ = nw25_
                d_136_v67_ = d_136_v67_
            d_137_v68_: _dafny.Seq
            d_137_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkantehx"))
            d_138_v69_: C7
            nw26_ = C7()
            nw26_.ctor__()
            d_138_v69_ = nw26_
            d_139_v70_: D20
            d_139_v70_ = D20_DC38(d_138_v69_)
            rhs7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjadvjecw"))
            rhs8_ = d_139_v70_
            rhs9_ = d_46_v0_
            lhs2_ = d_102_v51_
            d_137_v68_ = rhs7_
            d_139_v70_ = rhs8_
            lhs2_.f2 = rhs9_
            d_140_v71_: str
            d_140_v71_ = _dafny.CodePoint('b')
            d_141_v72_: _dafny.Seq
            d_141_v72_ = _dafny.SeqWithoutIsStrInference([(d_137_v68_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))).set(default__.safeIndex(d_102_v51_.f2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")))), d_140_v71_))])
            d_141_v72_ = d_141_v72_
            d_142_v73_: _dafny.Array
            nw27_ = _dafny.Array(None, 3)
            d_142_v73_ = nw27_
            d_143_v74_: C3
            nw28_ = C3()
            nw28_.ctor__(d_102_v51_.f2)
            d_143_v74_ = nw28_
            d_144_v75_: _dafny.Seq
            d_144_v75_ = _dafny.SeqWithoutIsStrInference([d_143_v74_, d_143_v74_])
            index17_ = default__.safeIndex(179, (d_142_v73_).length(0))
            (d_142_v73_)[index17_] = (d_144_v75_)[default__.safeIndex(len(d_137_v68_), len(d_144_v75_))]
            d_145_v76_: _dafny.Set
            d_145_v76_ = _dafny.Set({d_138_v69_})
            index18_ = default__.safeIndex(179, (d_142_v73_).length(0))
            nw29_ = C3()
            nw29_.ctor__((len((_dafny.Map({len(d_145_v76_): d_46_v0_})).set(d_46_v0_, d_102_v51_.f2))) * (d_102_v51_.f2))
            (d_142_v73_)[index18_] = nw29_
        elif True:
            d_146___mcc_h6_ = source4_.cf35
            d_147_cf35_ = d_146___mcc_h6_
            d_148_v77_: str
            d_148_v77_ = _dafny.CodePoint('d')
            d_149_v78_: _dafny.Seq
            d_149_v78_ = _dafny.SeqWithoutIsStrInference([d_148_v77_])
            d_148_v77_ = (d_149_v78_)[default__.safeIndex(d_102_v51_.f2, len(d_149_v78_))]
            d_150_v79_: _dafny.Array
            nw30_ = _dafny.Array(_dafny.Map({}), 24)
            d_150_v79_ = nw30_
            d_151_v80_: _dafny.Seq
            d_151_v80_ = _dafny.SeqWithoutIsStrInference([d_46_v0_])
            d_152_v81_: _dafny.Map
            d_152_v81_ = _dafny.Map({(_dafny.MultiSet(d_151_v80_)).cardinality: d_48_v2_})
            index19_ = default__.safeIndex(873, (d_150_v79_).length(0))
            (d_150_v79_)[index19_] = d_152_v81_
            index20_ = default__.safeIndex(873, (d_150_v79_).length(0))
            (d_150_v79_)[index20_] = d_152_v81_
            d_153_v82_: _dafny.MultiSet
            d_153_v82_ = _dafny.MultiSet([d_48_v2_])
            d_154_v83_: _dafny.Map
            d_154_v83_ = _dafny.Map({d_46_v0_: (d_153_v82_).cardinality})
            d_155_v84_: _dafny.Map
            d_155_v84_ = _dafny.Map({(_dafny.Map({(0) - (d_102_v51_.f2): d_46_v0_})) | (d_154_v83_): d_147_cf35_})
            d_155_v84_ = d_155_v84_
            d_156_v85_: _dafny.Map
            d_156_v85_ = _dafny.Map({d_48_v2_: 622})
            d_157_v86_: _dafny.Map
            d_157_v86_ = _dafny.Map({D14_DC27(True): (0) - (len(d_156_v85_))})
            d_158_v87_: C2
            nw31_ = C2()
            nw31_.ctor__(len((d_157_v86_) | (d_157_v86_)))
            d_158_v87_ = nw31_
        d_159_v89_: _dafny.Set
        def iife16_():
            coll8_ = _dafny.Set()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(-646, 248):
                d_160_v88_: int = compr_8_
                if ((-646) <= (d_160_v88_)) and ((d_160_v88_) < (248)):
                    coll8_ = coll8_.union(_dafny.Set([(d_160_v88_) * (len(_dafny.SeqWithoutIsStrInference([p1, p0, p1, p0])))]))
            return _dafny.Set(coll8_)
        d_159_v89_ = iife16_()
        
        pat_let_tv2_ = p0
        def lambda5_(source5_):
            d_161___mcc_h7_ = source5_
            d_162_cf48_ = d_161___mcc_h7_
            return pat_let_tv2_

        d_48_v2_ = lambda5_(d_159_v89_)
        d_163_v90_: _dafny.Array
        nw32_ = _dafny.Array(int(0), 10)
        d_163_v90_ = nw32_
        r0 = d_163_v90_
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_164_globalState_: GlobalState
        nw33_ = GlobalState()
        nw33_.ctor__()
        d_164_globalState_ = nw33_
        d_165_v0_: bool
        d_165_v0_ = True
        d_166_v1_: _dafny.Array
        out0_: _dafny.Array
        out0_ = default__.m0(d_165_v0_, (d_165_v0_) or (True), d_164_globalState_)
        d_166_v1_ = out0_
        d_165_v0_ = not(d_165_v0_)
        if d_165_v0_:
            d_167_v2_: str
            d_167_v2_ = _dafny.CodePoint('y')
            d_167_v2_ = d_167_v2_
            d_168_v3_: _dafny.Seq
            d_168_v3_ = _dafny.SeqWithoutIsStrInference([d_167_v2_])
            d_169_v4_: _dafny.Array
            def lambda6_(d_170_i0_):
                return not(True)

            init3_ = lambda6_
            nw34_ = _dafny.Array(None, 10)
            for i0_3_ in range(nw34_.length(0)):
                nw34_[i0_3_] = init3_(i0_3_)
            d_169_v4_ = nw34_
            d_171_v5_: C5
            nw35_ = C5()
            nw35_.ctor__(len(d_168_v3_), d_169_v4_)
            d_171_v5_ = nw35_
            d_172_v6_: _dafny.Seq
            d_172_v6_ = _dafny.SeqWithoutIsStrInference([(d_171_v5_).f0, len(d_168_v3_)])
            d_173_v7_: _dafny.MultiSet
            d_173_v7_ = _dafny.MultiSet([(d_171_v5_).f0])
            rhs10_ = ((_dafny.MultiSet((d_172_v6_).set(default__.safeIndex((d_171_v5_).f0, len(d_172_v6_)), (d_171_v5_).f0))) - (d_173_v7_)).ispropersubset(d_173_v7_)
            rhs11_ = d_165_v0_
            rhs12_ = d_165_v0_
            rhs13_ = (d_167_v2_) not in (d_168_v3_)
            d_165_v0_ = rhs10_
            d_165_v0_ = rhs11_
            d_165_v0_ = rhs12_
            d_165_v0_ = rhs13_
            d_174_v8_: D4
            d_174_v8_ = D4_DC5(_dafny.Set({d_165_v0_}))
            d_175_v9_: int
            d_175_v9_ = -538
            d_176_v10_: _dafny.Seq
            d_176_v10_ = _dafny.SeqWithoutIsStrInference([d_165_v0_, d_165_v0_, d_165_v0_])
            rhs14_ = (d_165_v0_ if d_165_v0_ else (d_176_v10_)[default__.safeIndex(len(d_176_v10_), len(d_176_v10_))])
            rhs15_ = d_174_v8_
            rhs16_ = (0) - ((d_175_v9_) - (d_175_v9_))
            d_165_v0_ = rhs14_
            d_174_v8_ = rhs15_
            d_175_v9_ = rhs16_
            d_177_v11_: _dafny.Array
            out1_: _dafny.Array
            out1_ = default__.m0(d_165_v0_, default__.fm3((d_171_v5_).fm6(d_165_v0_, d_164_globalState_), d_175_v9_, (d_171_v5_).f0, d_164_globalState_), d_164_globalState_)
            d_177_v11_ = out1_
        elif True:
            d_178_v12_: int
            d_178_v12_ = 229
            d_179_v13_: _dafny.Map
            d_179_v13_ = _dafny.Map({d_178_v12_: -100})
            d_180_v14_: _dafny.MultiSet
            d_180_v14_ = _dafny.MultiSet([d_179_v13_, d_179_v13_, d_179_v13_])
            d_181_v15_: _dafny.Seq
            d_181_v15_ = _dafny.SeqWithoutIsStrInference([(d_180_v14_).cardinality])
            d_181_v15_ = d_181_v15_
            d_182_v16_: _dafny.Seq
            d_182_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
            d_183_v17_: str
            d_183_v17_ = _dafny.CodePoint('a')
            d_184_v18_: C1
            nw36_ = C1()
            nw36_.ctor__((d_182_v16_).set(default__.safeIndex(len(d_181_v15_), len(d_182_v16_)), d_183_v17_))
            d_184_v18_ = nw36_
            d_165_v0_ = (D8_DC17(d_165_v0_, d_184_v18_, d_178_v12_)).cf30
            d_185_v19_: _dafny.Seq
            d_185_v19_ = _dafny.SeqWithoutIsStrInference([d_165_v0_, d_165_v0_, d_165_v0_])
            d_186_v20_: _dafny.Array
            nw37_ = _dafny.Array(None, 3)
            nw37_[int(0)] = d_165_v0_
            nw37_[int(1)] = (d_185_v19_)[default__.safeIndex(len((d_184_v18_).f3), len(d_185_v19_))]
            nw37_[int(2)] = d_165_v0_
            d_186_v20_ = nw37_
            index21_ = default__.safeIndex(354, (d_186_v20_).length(0))
            (d_186_v20_)[index21_] = not(d_165_v0_)
            index22_ = default__.safeIndex(354, (d_186_v20_).length(0))
            (d_186_v20_)[index22_] = d_165_v0_
            d_187_v21_: D6
            d_187_v21_ = D6_DC11((d_184_v18_).f3)
            d_188_v22_: _dafny.Map
            d_188_v22_ = _dafny.Map({d_187_v21_: d_178_v12_})
            d_188_v22_ = (d_188_v22_).set(d_187_v21_, default__.safeModuloInt(d_178_v12_, (0) - (d_178_v12_)))
            d_189_v23_: C3
            nw38_ = C3()
            nw38_.ctor__(d_178_v12_)
            d_189_v23_ = nw38_
            d_190_v24_: _dafny.Seq
            d_190_v24_ = _dafny.SeqWithoutIsStrInference([d_189_v23_, d_189_v23_])
            d_191_v25_: _dafny.Array
            nw39_ = _dafny.Array(None, 25)
            nw39_[int(0)] = d_189_v23_
            nw39_[int(1)] = d_189_v23_
            nw39_[int(2)] = d_189_v23_
            nw39_[int(3)] = d_189_v23_
            nw39_[int(4)] = d_189_v23_
            nw39_[int(5)] = d_189_v23_
            nw39_[int(6)] = (d_190_v24_)[default__.safeIndex(d_178_v12_, len(d_190_v24_))]
            nw39_[int(7)] = (d_189_v23_ if d_165_v0_ else d_189_v23_)
            nw39_[int(8)] = (d_189_v23_ if True else d_189_v23_)
            nw39_[int(9)] = d_189_v23_
            nw39_[int(10)] = d_189_v23_
            nw39_[int(11)] = d_189_v23_
            nw39_[int(12)] = d_189_v23_
            nw39_[int(13)] = d_189_v23_
            nw39_[int(14)] = d_189_v23_
            nw39_[int(15)] = d_189_v23_
            nw39_[int(16)] = d_189_v23_
            nw39_[int(17)] = d_189_v23_
            nw39_[int(18)] = d_189_v23_
            nw39_[int(19)] = d_189_v23_
            nw39_[int(20)] = d_189_v23_
            nw39_[int(21)] = (d_189_v23_ if (d_189_v23_).fm8(d_182_v16_, (d_186_v20_)[default__.safeIndex(354, (d_186_v20_).length(0))], d_164_globalState_) else d_189_v23_)
            nw39_[int(22)] = d_189_v23_
            nw39_[int(23)] = d_189_v23_
            nw39_[int(24)] = d_189_v23_
            d_191_v25_ = nw39_
            d_191_v25_ = d_191_v25_
        d_192_v27_: _dafny.Set
        def iife17_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(691, 586):
                d_193_v26_: int = compr_9_
                if ((691) <= (d_193_v26_)) and ((d_193_v26_) < (586)):
                    coll9_ = coll9_.union(_dafny.Set([(d_193_v26_) + (462)]))
            return _dafny.Set(coll9_)
        d_192_v27_ = iife17_()
        
        pat_let_tv3_ = d_165_v0_
        def lambda7_(source6_):
            d_194___mcc_h0_ = source6_
            d_195_cf48_ = d_194___mcc_h0_
            return pat_let_tv3_

        if lambda7_(d_192_v27_):
            d_196_v28_: int
            d_196_v28_ = 269
            d_197_v29_: _dafny.Seq
            d_197_v29_ = _dafny.SeqWithoutIsStrInference([(d_196_v28_) * (-921)])
            d_197_v29_ = d_197_v29_
            d_198_v30_: _dafny.Map
            d_198_v30_ = _dafny.Map({d_165_v0_: not(False)})
            d_199_v31_: _dafny.MultiSet
            d_199_v31_ = _dafny.MultiSet([((d_198_v30_)[d_165_v0_] if (d_165_v0_) in (d_198_v30_) else d_165_v0_)])
            if (d_199_v31_).ispropersubset((d_199_v31_) - (d_199_v31_)):
                d_200_v32_: _dafny.Array
                nw40_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_200_v32_ = nw40_
                index23_ = default__.safeIndex(979, (d_200_v32_).length(0))
                (d_200_v32_)[index23_] = d_166_v1_
                index24_ = default__.safeIndex(979, (d_200_v32_).length(0))
                (d_200_v32_)[index24_] = d_166_v1_
                d_165_v0_ = d_165_v0_
                d_201_v33_: _dafny.Array
                nw41_ = _dafny.Array(False, 9)
                d_201_v33_ = nw41_
                index25_ = default__.safeIndex(0, (d_201_v33_).length(0))
                (d_201_v33_)[index25_] = default__.fm5(d_164_globalState_)
                d_202_v34_: _dafny.Seq
                d_202_v34_ = _dafny.SeqWithoutIsStrInference([default__.fm5(d_164_globalState_), d_165_v0_, d_165_v0_])
                d_203_v35_: _dafny.Seq
                d_203_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgfsf"))
                index26_ = default__.safeIndex(0, (d_201_v33_).length(0))
                rhs17_ = (default__.safeModuloInt(len(d_202_v34_), len(d_203_v35_))) == ((0) - (((d_199_v31_).intersection(d_199_v31_)).cardinality))
                rhs18_ = d_165_v0_
                lhs3_ = d_201_v33_
                lhs4_ = default__.safeIndex(0, (d_201_v33_).length(0))
                d_165_v0_ = rhs17_
                lhs3_[lhs4_] = rhs18_
                d_204_v36_: _dafny.Map
                d_204_v36_ = _dafny.Map({d_196_v28_: d_196_v28_})
                d_205_v37_: _dafny.Seq
                d_205_v37_ = _dafny.SeqWithoutIsStrInference([d_204_v36_])
                arr0_ = (d_200_v32_)[default__.safeIndex(979, (d_200_v32_).length(0))]
                index27_ = default__.safeIndex(882, ((d_200_v32_)[default__.safeIndex(979, (d_200_v32_).length(0))]).length(0))
                arr0_[index27_] = (len(d_197_v29_)) * (len(d_205_v37_))
                d_206_v38_: C6
                nw42_ = C6()
                nw42_.ctor__()
                d_206_v38_ = nw42_
                d_207_v39_: _dafny.Seq
                d_207_v39_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_206_v38_: d_196_v28_})])
                arr1_ = (d_200_v32_)[default__.safeIndex(979, (d_200_v32_).length(0))]
                index28_ = default__.safeIndex(882, ((d_200_v32_)[default__.safeIndex(979, (d_200_v32_).length(0))]).length(0))
                arr1_[index28_] = len(d_207_v39_)
                d_165_v0_ = (d_201_v33_)[default__.safeIndex(0, (d_201_v33_).length(0))]
            elif True:
                d_208_v40_: _dafny.Seq
                d_208_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cpimjma"))
                d_208_v40_ = d_208_v40_
                d_209_v41_: C1
                nw43_ = C1()
                nw43_.ctor__(d_208_v40_)
                d_209_v41_ = nw43_
                d_208_v40_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_210_i1_ in range(default__.abs(991))])
                d_211_v42_: _dafny.Array
                def lambda8_(d_212_i2_):
                    return not(False)

                init4_ = lambda8_
                nw44_ = _dafny.Array(None, 27)
                for i0_4_ in range(nw44_.length(0)):
                    nw44_[i0_4_] = init4_(i0_4_)
                d_211_v42_ = nw44_
                index29_ = default__.safeIndex(258, (d_211_v42_).length(0))
                (d_211_v42_)[index29_] = d_165_v0_
                index30_ = default__.safeIndex(258, (d_211_v42_).length(0))
                (d_211_v42_)[index30_] = (d_196_v28_) < (d_196_v28_)
                index31_ = default__.safeIndex(92, (d_166_v1_).length(0))
                (d_166_v1_)[index31_] = d_196_v28_
                d_213_v43_: _dafny.Array
                nw45_ = _dafny.Array(None, 17)
                d_213_v43_ = nw45_
                index32_ = default__.safeIndex(92, (d_166_v1_).length(0))
                rhs19_ = d_196_v28_
                rhs20_ = (d_196_v28_) - (default__.safeDivisionInt(d_196_v28_, (d_209_v41_).fm1(d_165_v0_, d_164_globalState_)))
                rhs21_ = d_213_v43_
                lhs5_ = d_166_v1_
                lhs6_ = default__.safeIndex(92, (d_166_v1_).length(0))
                d_196_v28_ = rhs19_
                lhs5_[lhs6_] = rhs20_
                d_213_v43_ = rhs21_
            d_214_v44_: _dafny.Array
            out2_: _dafny.Array
            out2_ = default__.m0((d_199_v31_).isdisjoint(d_199_v31_), d_165_v0_, d_164_globalState_)
            d_214_v44_ = out2_
            d_215_v45_: _dafny.Array
            out3_: _dafny.Array
            out3_ = default__.m0(d_165_v0_, d_165_v0_, d_164_globalState_)
            d_215_v45_ = out3_
            index33_ = default__.safeIndex(929, (d_215_v45_).length(0))
            (d_215_v45_)[index33_] = -497
            d_216_v46_: C6
            nw46_ = C6()
            nw46_.ctor__()
            d_216_v46_ = nw46_
            d_217_v47_: _dafny.Map
            d_217_v47_ = _dafny.Map({d_216_v46_: d_196_v28_})
            d_218_v48_: _dafny.Map
            d_218_v48_ = _dafny.Map({d_196_v28_: d_166_v1_})
            index34_ = default__.safeIndex(929, (d_215_v45_).length(0))
            (d_215_v45_)[index34_] = default__.safeModuloInt(((d_217_v47_)[d_216_v46_] if (d_216_v46_) in (d_217_v47_) else -686), len(d_218_v48_))
        elif True:
            d_219_v49_: _dafny.Seq
            d_219_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htengavdr"))
            d_220_v50_: int
            d_220_v50_ = -705
            d_221_v51_: str
            d_221_v51_ = _dafny.CodePoint('q')
            d_219_v49_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wafife"))).set(default__.safeIndex(d_220_v50_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wafife")))), d_221_v51_)
            d_222_v52_: _dafny.Seq
            d_222_v52_ = _dafny.SeqWithoutIsStrInference([d_220_v50_])
            d_223_v53_: _dafny.Set
            d_223_v53_ = _dafny.Set({d_165_v0_})
            d_224_v54_: _dafny.Map
            d_224_v54_ = _dafny.Map({(d_222_v52_)[default__.safeIndex(len(d_223_v53_), len(d_222_v52_))]: d_166_v1_})
            d_225_v55_: _dafny.Array
            nw47_ = _dafny.Array(None, 10)
            nw47_[int(0)] = ((d_224_v54_)[len(_dafny.SeqWithoutIsStrInference([d_220_v50_, d_220_v50_]))] if (len(_dafny.SeqWithoutIsStrInference([d_220_v50_, d_220_v50_]))) in (d_224_v54_) else d_166_v1_)
            nw47_[int(1)] = d_166_v1_
            nw47_[int(2)] = d_166_v1_
            nw47_[int(3)] = d_166_v1_
            nw47_[int(4)] = d_166_v1_
            nw47_[int(5)] = (d_166_v1_ if d_165_v0_ else d_166_v1_)
            nw47_[int(6)] = d_166_v1_
            nw47_[int(7)] = d_166_v1_
            nw47_[int(8)] = d_166_v1_
            nw47_[int(9)] = d_166_v1_
            d_225_v55_ = nw47_
            d_226_v56_: _dafny.Seq
            d_226_v56_ = _dafny.SeqWithoutIsStrInference([d_225_v55_, (d_225_v55_ if d_165_v0_ else d_225_v55_)])
            d_225_v55_ = (d_226_v56_)[default__.safeIndex(d_220_v50_, len(d_226_v56_))]
            d_227_v57_: _dafny.MultiSet
            d_227_v57_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_221_v51_ for d_228_i3_ in range(default__.abs(464))]), d_219_v49_])
            d_229_v58_: C0
            nw48_ = C0()
            nw48_.ctor__((_dafny.MultiSet([d_219_v49_, d_219_v49_])) != (d_227_v57_), d_165_v0_)
            d_229_v58_ = nw48_
            index35_ = default__.safeIndex(512, (d_166_v1_).length(0))
            (d_166_v1_)[index35_] = d_220_v50_
            index36_ = default__.safeIndex(512, (d_166_v1_).length(0))
            (d_166_v1_)[index36_] = d_220_v50_
            d_230_v59_: C7
            nw49_ = C7()
            nw49_.ctor__()
            d_230_v59_ = nw49_
        if d_165_v0_:
            d_231_v60_: int
            d_231_v60_ = 823
            d_232_v61_: _dafny.Seq
            d_232_v61_ = _dafny.SeqWithoutIsStrInference([d_231_v60_])
            d_233_v62_: _dafny.Seq
            d_233_v62_ = _dafny.SeqWithoutIsStrInference([(d_232_v61_).set(default__.safeIndex(d_231_v60_, len(d_232_v61_)), d_231_v60_)])
            d_234_v63_: _dafny.Map
            d_234_v63_ = _dafny.Map({d_165_v0_: len(d_232_v61_)})
            d_235_v64_: _dafny.Seq
            d_235_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjwbxhkeo"))
            d_236_v65_: _dafny.Set
            d_236_v65_ = _dafny.Set({d_231_v60_, len(d_233_v62_), ((d_234_v63_)[d_165_v0_] if (d_165_v0_) in (d_234_v63_) else d_231_v60_), d_231_v60_, (d_231_v60_ if d_165_v0_ else len(d_235_v64_))})
            d_236_v65_ = d_236_v65_
            d_237_v66_: _dafny.Array
            def lambda9_(d_238_v0_):
                def lambda10_(d_239_i4_):
                    return d_238_v0_

                return lambda10_

            init5_ = lambda9_(d_165_v0_)
            nw50_ = _dafny.Array(None, 17)
            for i0_5_ in range(nw50_.length(0)):
                nw50_[i0_5_] = init5_(i0_5_)
            d_237_v66_ = nw50_
            index37_ = default__.safeIndex(563, (d_237_v66_).length(0))
            (d_237_v66_)[index37_] = d_165_v0_
            index38_ = default__.safeIndex(320, (d_237_v66_).length(0))
            (d_237_v66_)[index38_] = d_165_v0_
            index39_ = default__.safeIndex(563, (d_237_v66_).length(0))
            index40_ = default__.safeIndex(320, (d_237_v66_).length(0))
            rhs22_ = (d_165_v0_) and (d_165_v0_)
            rhs23_ = d_165_v0_
            lhs7_ = d_237_v66_
            lhs8_ = default__.safeIndex(563, (d_237_v66_).length(0))
            lhs9_ = d_237_v66_
            lhs10_ = default__.safeIndex(320, (d_237_v66_).length(0))
            lhs7_[lhs8_] = rhs22_
            lhs9_[lhs10_] = rhs23_
            d_240_v67_: _dafny.Array
            out4_: _dafny.Array
            out4_ = default__.m0(d_165_v0_, (d_237_v66_)[default__.safeIndex(563, (d_237_v66_).length(0))], d_164_globalState_)
            d_240_v67_ = out4_
            d_241_v68_: C4
            nw51_ = C4()
            nw51_.ctor__(d_231_v60_)
            d_241_v68_ = nw51_
            d_242_v69_: _dafny.Map
            d_242_v69_ = _dafny.Map({d_241_v68_: (d_237_v66_)[default__.safeIndex(563, (d_237_v66_).length(0))]})
            d_243_v70_: _dafny.Seq
            d_243_v70_ = _dafny.SeqWithoutIsStrInference([d_242_v69_])
            d_244_v71_: _dafny.Map
            d_244_v71_ = _dafny.Map({(d_237_v66_)[default__.safeIndex(563, (d_237_v66_).length(0))]: d_242_v69_})
            d_245_v72_: _dafny.Array
            nw52_ = _dafny.Array(None, 26)
            nw52_[int(0)] = (d_242_v69_) | (d_242_v69_)
            nw52_[int(1)] = (d_242_v69_) | (d_242_v69_)
            nw52_[int(2)] = d_242_v69_
            nw52_[int(3)] = d_242_v69_
            nw52_[int(4)] = d_242_v69_
            nw52_[int(5)] = (d_242_v69_) | (d_242_v69_)
            nw52_[int(6)] = (_dafny.Map({d_241_v68_: (d_237_v66_)[default__.safeIndex(563, (d_237_v66_).length(0))]})) | (d_242_v69_)
            nw52_[int(7)] = d_242_v69_
            nw52_[int(8)] = (d_242_v69_) | (_dafny.Map({d_241_v68_: (d_241_v68_).fm8(d_235_v64_, (d_237_v66_)[default__.safeIndex(563, (d_237_v66_).length(0))], d_164_globalState_)}))
            nw52_[int(9)] = _dafny.Map({d_241_v68_: d_165_v0_})
            nw52_[int(10)] = (d_243_v70_)[default__.safeIndex(d_231_v60_, len(d_243_v70_))]
            nw52_[int(11)] = d_242_v69_
            nw52_[int(12)] = d_242_v69_
            nw52_[int(13)] = d_242_v69_
            nw52_[int(14)] = ((d_244_v71_)[d_165_v0_] if (d_165_v0_) in (d_244_v71_) else d_242_v69_)
            nw52_[int(15)] = (d_242_v69_) | (d_242_v69_)
            nw52_[int(16)] = d_242_v69_
            nw52_[int(17)] = (D18_DC33(d_242_v69_)).cf50
            nw52_[int(18)] = (d_242_v69_) | ((_dafny.Map({d_241_v68_: d_165_v0_})).set(d_241_v68_, d_165_v0_))
            nw52_[int(19)] = d_242_v69_
            nw52_[int(20)] = d_242_v69_
            nw52_[int(21)] = (d_242_v69_) | (d_242_v69_)
            nw52_[int(22)] = d_242_v69_
            nw52_[int(23)] = _dafny.Map({d_241_v68_: d_165_v0_})
            nw52_[int(24)] = d_242_v69_
            nw52_[int(25)] = ((d_242_v69_).set(d_241_v68_, False)) | (d_242_v69_)
            d_245_v72_ = nw52_
            d_245_v72_ = d_245_v72_
            if (d_237_v66_)[default__.safeIndex(563, (d_237_v66_).length(0))]:
                d_246_v73_: _dafny.Array
                nw53_ = _dafny.Array(None, 22)
                d_246_v73_ = nw53_
                d_247_v74_: _dafny.MultiSet
                d_247_v74_ = _dafny.MultiSet([d_231_v60_, d_231_v60_])
                d_248_v75_: _dafny.MultiSet
                d_248_v75_ = _dafny.MultiSet([d_247_v74_])
                d_249_v76_: _dafny.MultiSet
                d_249_v76_ = _dafny.MultiSet([(d_248_v75_).cardinality, len(d_235_v64_)])
                d_250_v77_: T1
                nw54_ = C4()
                nw54_.ctor__(d_231_v60_)
                d_250_v77_ = nw54_
                d_251_v78_: _dafny.MultiSet
                d_251_v78_ = _dafny.MultiSet([d_250_v77_])
                d_252_v79_: T0
                nw55_ = C5()
                nw55_.ctor__(((d_247_v74_)[(d_241_v68_).fm1((d_237_v66_)[default__.safeIndex(563, (d_237_v66_).length(0))], d_164_globalState_)] if ((d_241_v68_).fm1((d_237_v66_)[default__.safeIndex(563, (d_237_v66_).length(0))], d_164_globalState_)) in (d_247_v74_) else (d_251_v78_).cardinality), d_237_v66_)
                d_252_v79_ = nw55_
                index41_ = default__.safeIndex(484, (d_246_v73_).length(0))
                (d_246_v73_)[index41_] = (d_252_v79_)
                d_253_v80_: _dafny.Map
                d_253_v80_ = _dafny.Map({d_231_v60_: d_250_v77_})
                d_254_v81_: _dafny.Map
                d_254_v81_ = _dafny.Map({d_231_v60_: d_253_v80_})
                d_255_v82_: _dafny.Set
                d_255_v82_ = _dafny.Set({(d_253_v80_) | (((d_254_v81_)[d_231_v60_] if (d_231_v60_) in (d_254_v81_) else d_253_v80_)), _dafny.Map({d_231_v60_: d_250_v77_}), (d_253_v80_).set(d_231_v60_, d_250_v77_), _dafny.Map({(0) - (d_231_v60_): d_250_v77_}), _dafny.Map({d_231_v60_: d_250_v77_})})
                d_256_v83_: _dafny.Seq
                d_256_v83_ = _dafny.SeqWithoutIsStrInference([d_252_v79_])
                d_257_v84_: _dafny.Map
                d_257_v84_ = _dafny.Map({d_232_v61_: (d_256_v83_)[default__.safeIndex((0) - (-19), len(d_256_v83_))]})
                index42_ = default__.safeIndex(484, (d_246_v73_).length(0))
                rhs24_ = ((d_257_v84_)[d_232_v61_] if (d_232_v61_) in (d_257_v84_) else d_252_v79_)
                rhs25_ = d_255_v82_
                lhs11_ = d_246_v73_
                lhs12_ = default__.safeIndex(484, (d_246_v73_).length(0))
                lhs11_[lhs12_] = rhs24_
                d_255_v82_ = rhs25_
                d_258_v85_: C1
                nw56_ = C1()
                nw56_.ctor__(d_235_v64_)
                d_258_v85_ = nw56_
                d_259_v86_: _dafny.Set
                d_259_v86_ = _dafny.Set({d_258_v85_, d_258_v85_})
                d_260_v87_: _dafny.Set
                d_260_v87_ = _dafny.Set({(d_259_v86_).intersection(d_259_v86_)})
                d_260_v87_ = (d_260_v87_) - (d_260_v87_)
                d_261_v88_: _dafny.Array
                nw57_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_261_v88_ = nw57_
                d_262_v89_: _dafny.Array
                nw58_ = _dafny.Array(_dafny.Seq({}), 26)
                d_262_v89_ = nw58_
                index43_ = default__.safeIndex(268, (d_261_v88_).length(0))
                (d_261_v88_)[index43_] = d_262_v89_
                index44_ = default__.safeIndex(268, (d_261_v88_).length(0))
                (d_261_v88_)[index44_] = d_262_v89_
                d_165_v0_ = d_165_v0_
                d_263_v90_: _dafny.Array
                nw59_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_263_v90_ = nw59_
                index45_ = default__.safeIndex(532, (d_263_v90_).length(0))
                (d_263_v90_)[index45_] = d_237_v66_
                d_264_v91_: _dafny.MultiSet
                d_264_v91_ = _dafny.MultiSet([d_166_v1_])
                index46_ = default__.safeIndex(563, (d_237_v66_).length(0))
                index47_ = default__.safeIndex(532, (d_263_v90_).length(0))
                rhs26_ = ((d_264_v91_) | (d_264_v91_)).issubset(d_264_v91_)
                rhs27_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uookla"))) + ((d_258_v85_).f3)
                rhs28_ = d_231_v60_
                rhs29_ = d_237_v66_
                lhs13_ = d_237_v66_
                lhs14_ = default__.safeIndex(563, (d_237_v66_).length(0))
                lhs15_ = d_263_v90_
                lhs16_ = default__.safeIndex(532, (d_263_v90_).length(0))
                lhs13_[lhs14_] = rhs26_
                d_235_v64_ = rhs27_
                d_231_v60_ = rhs28_
                lhs15_[lhs16_] = rhs29_
            elif True:
                d_265_v92_: _dafny.MultiSet
                d_265_v92_ = _dafny.MultiSet([d_231_v60_, d_231_v60_])
                d_231_v60_ = ((d_231_v60_) - ((d_265_v92_).cardinality)) * (d_231_v60_)
                d_266_v93_: str
                d_266_v93_ = _dafny.CodePoint('y')
                d_266_v93_ = d_266_v93_
                d_266_v93_ = d_266_v93_
                d_231_v60_ = d_231_v60_
                d_231_v60_ = d_231_v60_
        elif True:
            d_267_v94_: _dafny.Seq
            d_267_v94_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlkwbx"))
            d_267_v94_ = (d_267_v94_) + ((d_267_v94_) + (d_267_v94_))
            d_268_v95_: int
            d_268_v95_ = 785
            d_269_v96_: str
            d_269_v96_ = _dafny.CodePoint('y')
            d_270_v97_: _dafny.Map
            d_270_v97_ = _dafny.Map({d_165_v0_: _dafny.MultiSet([(d_267_v94_).set(default__.safeIndex(d_268_v95_, len(d_267_v94_)), d_269_v96_), d_267_v94_])})
            d_271_v98_: _dafny.MultiSet
            d_271_v98_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rllwbgim"))])
            d_270_v97_ = (d_270_v97_).set(True, (d_271_v98_).intersection(d_271_v98_))
            d_272_v99_: _dafny.Map
            d_272_v99_ = _dafny.Map({d_165_v0_: d_165_v0_})
            d_273_v100_: _dafny.Map
            d_273_v100_ = _dafny.Map({d_165_v0_: _dafny.SeqWithoutIsStrInference([((d_272_v99_)[False] if (False) in (d_272_v99_) else not(False))])})
            d_274_v101_: _dafny.Seq
            d_274_v101_ = _dafny.SeqWithoutIsStrInference([d_165_v0_, d_165_v0_])
            d_273_v100_ = (d_273_v100_).set(d_165_v0_, (d_274_v101_) + (default__.fm16(d_164_globalState_)))
            d_275_v102_: D6
            d_275_v102_ = D6_DC11(d_267_v94_)
            d_276_v103_: _dafny.Map
            d_276_v103_ = _dafny.Map({d_275_v102_: d_268_v95_})
            d_277_v104_: _dafny.Set
            d_277_v104_ = _dafny.Set({d_276_v103_})
            d_278_v105_: _dafny.Set
            d_278_v105_ = d_277_v104_
            d_279_v106_: _dafny.Array
            nw60_ = _dafny.Array(None, 16)
            d_279_v106_ = nw60_
            d_280_v107_: _dafny.Map
            d_280_v107_ = _dafny.Map({d_278_v105_: d_279_v106_})
            d_280_v107_ = (d_280_v107_).set(d_278_v105_, d_279_v106_)
            index48_ = default__.safeIndex(493, (d_166_v1_).length(0))
            (d_166_v1_)[index48_] = d_268_v95_
            index49_ = default__.safeIndex(493, (d_166_v1_).length(0))
            (d_166_v1_)[index49_] = default__.safeModuloInt(d_268_v95_, (0) - (d_268_v95_))
        d_166_v1_ = d_166_v1_
        if d_165_v0_:
            d_281_v108_: _dafny.Seq
            d_281_v108_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "roxp"))
            d_282_v109_: D6
            d_282_v109_ = D6_DC11(d_281_v108_)
            pat_let_tv4_ = d_281_v108_
            d_283_v110_: _dafny.Set
            def iife18_(_pat_let4_0):
                def iife19_(d_284_dt__update__tmp_h0_):
                    def iife20_(_pat_let5_0):
                        def iife21_(d_285_dt__update_hcf21_h0_):
                            return D6_DC11(d_285_dt__update_hcf21_h0_)
                        return iife21_(_pat_let5_0)
                    return iife20_(pat_let_tv4_)
                return iife19_(_pat_let4_0)
            d_283_v110_ = _dafny.Set({iife18_(d_282_v109_), d_282_v109_})
            d_286_v113_: int
            d_286_v113_ = -20
            d_287_v114_: _dafny.Seq
            d_287_v114_ = _dafny.SeqWithoutIsStrInference([default__.fm5(d_164_globalState_)])
            d_288_v115_: _dafny.MultiSet
            d_288_v115_ = _dafny.MultiSet([d_282_v109_, D6_DC11(d_281_v108_)])
            d_289_v116_: D20
            d_289_v116_ = D20_DC36(d_288_v115_)
            d_290_v119_: _dafny.Map
            d_290_v119_ = _dafny.Map({d_282_v109_: d_165_v0_})
            d_291_v121_: _dafny.Array
            nw61_ = _dafny.Array(None, 14)
            nw61_[int(0)] = d_283_v110_
            def iife22_():
                def iife24_():
                    coll12_ = _dafny.Set()
                    compr_12_: D6
                    for compr_12_ in (d_283_v110_).Elements:
                        d_294_v111_: D6 = compr_12_
                        if (d_294_v111_) in (d_283_v110_):
                            coll12_ = coll12_.union(_dafny.Set([d_294_v111_]))
                    return _dafny.Set(coll12_)
                coll10_ = _dafny.Set()
                def iife23_():
                    coll11_ = _dafny.Set()
                    compr_11_: D6
                    for compr_11_ in (d_283_v110_).Elements:
                        d_292_v111_: D6 = compr_11_
                        if (d_292_v111_) in (d_283_v110_):
                            coll11_ = coll11_.union(_dafny.Set([d_292_v111_]))
                    return _dafny.Set(coll11_)
                compr_10_: D6
                for compr_10_ in (iife23_()
                ).Elements:
                    d_293_v112_: D6 = compr_10_
                    if (d_293_v112_) in (iife24_()
                    ):
                        coll10_ = coll10_.union(_dafny.Set([d_293_v112_]))
                return _dafny.Set(coll10_)
            nw61_[int(1)] = iife22_()
            
            nw61_[int(2)] = _dafny.Set({d_282_v109_, d_282_v109_, d_282_v109_, d_282_v109_, d_282_v109_})
            nw61_[int(3)] = _dafny.Set({d_282_v109_, d_282_v109_})
            nw61_[int(4)] = d_283_v110_
            nw61_[int(5)] = d_283_v110_
            nw61_[int(6)] = d_283_v110_
            nw61_[int(7)] = default__.fm52(d_165_v0_, d_286_v113_, len(default__.fm16(d_164_globalState_)), (d_287_v114_)[default__.safeIndex(d_286_v113_, len(d_287_v114_))], d_164_globalState_)
            nw61_[int(8)] = d_283_v110_
            nw61_[int(9)] = d_283_v110_
            nw61_[int(10)] = d_283_v110_
            def iife25_():
                coll13_ = _dafny.Set()
                compr_13_: D6
                for compr_13_ in ((d_289_v116_).cf54).Elements:
                    d_295_v117_: D6 = compr_13_
                    if (d_295_v117_) in ((d_289_v116_).cf54):
                        coll13_ = coll13_.union(_dafny.Set([d_295_v117_]))
                return _dafny.Set(coll13_)
            def iife26_():
                coll14_ = _dafny.Set()
                compr_14_: D6
                for compr_14_ in (_dafny.Map({d_282_v109_: d_286_v113_})).keys.Elements:
                    d_296_v118_: D6 = compr_14_
                    if (d_296_v118_) in (_dafny.Map({d_282_v109_: d_286_v113_})):
                        coll14_ = coll14_.union(_dafny.Set([d_296_v118_]))
                return _dafny.Set(coll14_)
            nw61_[int(11)] = (iife25_()
            ) | (iife26_()
            )
            nw61_[int(12)] = d_283_v110_
            def iife27_():
                coll15_ = _dafny.Set()
                compr_15_: D6
                for compr_15_ in (d_290_v119_).keys.Elements:
                    d_297_v120_: D6 = compr_15_
                    if (d_297_v120_) in (d_290_v119_):
                        coll15_ = coll15_.union(_dafny.Set([d_297_v120_]))
                return _dafny.Set(coll15_)
            nw61_[int(13)] = iife27_()
            
            d_291_v121_ = nw61_
            d_298_v122_: _dafny.Set
            d_298_v122_ = _dafny.Set({not(d_165_v0_), d_165_v0_})
            d_291_v121_ = (d_291_v121_ if (d_298_v122_).isdisjoint(d_298_v122_) else d_291_v121_)
            def iife28_():
                coll16_ = _dafny.Map()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(492, -52):
                    d_299_v123_: int = compr_16_
                    if ((492) <= (d_299_v123_)) and ((d_299_v123_) < (-52)):
                        coll16_[(d_299_v123_) + (d_286_v113_)] = d_286_v113_
                return _dafny.Map(coll16_)
            rhs30_ = (d_281_v108_) + (d_281_v108_)
            rhs31_ = (d_281_v108_) != (d_281_v108_)
            rhs32_ = (d_286_v113_) not in (iife28_()
            )
            rhs33_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgybp"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubiaagdx")))) + ((d_281_v108_) + (d_281_v108_))
            d_281_v108_ = rhs30_
            d_165_v0_ = rhs31_
            d_165_v0_ = rhs32_
            d_281_v108_ = rhs33_
            d_165_v0_ = d_165_v0_
            d_300_v124_: _dafny.Array
            nw62_ = _dafny.Array(False, 4)
            d_300_v124_ = nw62_
            d_301_v125_: _dafny.Map
            d_301_v125_ = _dafny.Map({d_286_v113_: d_300_v124_})
            d_302_v126_: _dafny.Array
            nw63_ = _dafny.Array(None, 8)
            nw63_[int(0)] = d_300_v124_
            nw63_[int(1)] = d_300_v124_
            nw63_[int(2)] = d_300_v124_
            nw63_[int(3)] = d_300_v124_
            nw63_[int(4)] = d_300_v124_
            nw63_[int(5)] = (d_300_v124_ if d_165_v0_ else d_300_v124_)
            nw63_[int(6)] = d_300_v124_
            nw63_[int(7)] = ((d_301_v125_)[d_286_v113_] if (d_286_v113_) in (d_301_v125_) else d_300_v124_)
            d_302_v126_ = nw63_
            index50_ = default__.safeIndex(852, (d_302_v126_).length(0))
            (d_302_v126_)[index50_] = d_300_v124_
            d_303_v127_: D0
            d_303_v127_ = D0_DC0(d_165_v0_)
            d_304_v128_: D14
            d_304_v128_ = D14_DC26(d_303_v127_, d_298_v122_)
            index51_ = default__.safeIndex(852, (d_302_v126_).length(0))
            rhs34_ = d_281_v108_
            rhs35_ = d_300_v124_
            rhs36_ = default__.fm53(d_165_v0_, d_165_v0_, d_164_globalState_)
            lhs17_ = d_302_v126_
            lhs18_ = default__.safeIndex(852, (d_302_v126_).length(0))
            d_281_v108_ = rhs34_
            lhs17_[lhs18_] = rhs35_
            d_304_v128_ = rhs36_
            index52_ = default__.safeIndex(870, (d_166_v1_).length(0))
            (d_166_v1_)[index52_] = d_286_v113_
            d_305_v129_: _dafny.MultiSet
            d_305_v129_ = _dafny.MultiSet([d_166_v1_, d_166_v1_])
            index53_ = default__.safeIndex(870, (d_166_v1_).length(0))
            rhs37_ = (d_305_v129_).cardinality
            rhs38_ = default__.fm54(d_165_v0_, (d_281_v108_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ef"))), d_164_globalState_)
            rhs39_ = d_281_v108_
            rhs40_ = d_166_v1_
            lhs19_ = d_166_v1_
            lhs20_ = default__.safeIndex(870, (d_166_v1_).length(0))
            lhs19_[lhs20_] = rhs37_
            d_286_v113_ = rhs38_
            d_281_v108_ = rhs39_
            d_166_v1_ = rhs40_
        elif True:
            d_165_v0_ = False
            d_306_v130_: int
            d_306_v130_ = 627
            index54_ = default__.safeIndex(746, (d_166_v1_).length(0))
            (d_166_v1_)[index54_] = d_306_v130_
            d_307_v131_: _dafny.Seq
            d_307_v131_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")))])
            index55_ = default__.safeIndex(746, (d_166_v1_).length(0))
            (d_166_v1_)[index55_] = ((d_306_v130_) - (len(d_307_v131_))) + (-306)
            d_165_v0_ = d_165_v0_
            d_165_v0_ = d_165_v0_
            index56_ = default__.safeIndex(746, (d_166_v1_).length(0))
            (d_166_v1_)[index56_] = default__.safeDivisionInt(d_306_v130_, d_306_v130_)
        d_308_v132_: int
        d_308_v132_ = 786
        d_309_v133_: _dafny.Map
        d_309_v133_ = _dafny.Map({d_165_v0_: (0) - (d_308_v132_)})
        d_309_v133_ = (d_309_v133_).set(d_165_v0_, d_308_v132_)
        d_310_v134_: _dafny.Seq
        d_310_v134_ = _dafny.SeqWithoutIsStrInference([d_308_v132_, 390])
        d_311_v135_: _dafny.Map
        d_311_v135_ = _dafny.Map({d_308_v132_: d_165_v0_})
        d_312_v136_: _dafny.Seq
        d_312_v136_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmigsl"))
        d_313_v137_: C2
        nw64_ = C2()
        nw64_.ctor__(default__.fm54(d_165_v0_, d_312_v136_, d_164_globalState_))
        d_313_v137_ = nw64_
        d_314_v138_: _dafny.Seq
        d_314_v138_ = _dafny.SeqWithoutIsStrInference([d_313_v137_, d_313_v137_, d_313_v137_, d_313_v137_, d_313_v137_])
        hi0_ = default__.safeDivisionInt(len(d_311_v135_), len(d_314_v138_))
        for d_315_i5_ in range((d_310_v134_)[default__.safeIndex(d_308_v132_, len(d_310_v134_))], hi0_):
            d_316_v139_: C6
            nw65_ = C6()
            nw65_.ctor__()
            d_316_v139_ = nw65_
            d_317_v140_: _dafny.Set
            d_317_v140_ = _dafny.Set({d_308_v132_})
            d_318_v141_: str
            d_318_v141_ = _dafny.CodePoint('s')
            d_319_v142_: D0
            d_319_v142_ = D0_DC0(d_165_v0_)
            d_320_v143_: _dafny.Set
            d_320_v143_ = _dafny.Set({d_165_v0_, d_165_v0_})
            d_321_v144_: D14
            d_321_v144_ = D14_DC26(d_319_v142_, d_320_v143_)
            d_322_v145_: _dafny.Map
            d_322_v145_ = _dafny.Map({d_165_v0_: d_165_v0_})
            d_323_v146_: _dafny.Array
            nw66_ = _dafny.Array(None, 17)
            nw66_[int(0)] = _dafny.SeqWithoutIsStrInference([(0) - (d_315_i5_)])
            nw66_[int(1)] = d_310_v134_
            nw66_[int(2)] = d_310_v134_
            nw66_[int(3)] = d_310_v134_
            nw66_[int(4)] = d_310_v134_
            nw66_[int(5)] = default__.fm36(d_165_v0_, d_318_v141_, d_308_v132_, d_164_globalState_)
            nw66_[int(6)] = (_dafny.SeqWithoutIsStrInference([d_308_v132_, len((d_321_v144_).cf44)])).set(default__.safeIndex(d_308_v132_, len(_dafny.SeqWithoutIsStrInference([d_308_v132_, len((d_321_v144_).cf44)]))), d_315_i5_)
            nw66_[int(7)] = d_310_v134_
            nw66_[int(8)] = d_310_v134_
            nw66_[int(9)] = d_310_v134_
            nw66_[int(10)] = d_310_v134_
            nw66_[int(11)] = d_310_v134_
            nw66_[int(12)] = d_310_v134_
            nw66_[int(13)] = d_310_v134_
            nw66_[int(14)] = (d_310_v134_).set(default__.safeIndex(len(d_322_v145_), len(d_310_v134_)), d_315_i5_)
            nw66_[int(15)] = (d_310_v134_).set(default__.safeIndex(d_308_v132_, len(d_310_v134_)), d_315_i5_)
            nw66_[int(16)] = d_310_v134_
            d_323_v146_ = nw66_
            d_324_v147_: bool
            d_325_v148_: _dafny.Array
            d_326_v149_: int
            out5_: bool
            out6_: _dafny.Array
            out7_: int
            out5_, out6_, out7_ = (d_313_v137_).m1(len((d_317_v140_) | (d_317_v140_)), d_323_v146_, d_165_v0_, (d_310_v134_) < (d_310_v134_), d_164_globalState_)
            d_324_v147_ = out5_
            d_325_v148_ = out6_
            d_326_v149_ = out7_
            d_324_v147_ = default__.fm5(d_164_globalState_)
            d_327_v150_: C4
            nw67_ = C4()
            nw67_.ctor__(d_315_i5_)
            d_327_v150_ = nw67_
            d_328_v151_: D21
            d_328_v151_ = D21_DC39(d_327_v150_)
            d_329_v152_: _dafny.Seq
            d_329_v152_ = _dafny.SeqWithoutIsStrInference([(d_328_v151_).cf58, d_327_v150_, d_327_v150_, d_327_v150_])
            d_327_v150_ = (d_329_v152_)[default__.safeIndex(d_308_v132_, len(d_329_v152_))]
        if (d_165_v0_ if d_165_v0_ else d_165_v0_):
            d_330_v153_: int
            d_331_v154_: int
            out8_: int
            out9_: int
            out8_, out9_ = (d_313_v137_).m3(_dafny.MultiSet(((d_310_v134_) + (_dafny.SeqWithoutIsStrInference([553]))).set(default__.safeIndex(d_308_v132_, len((d_310_v134_) + (_dafny.SeqWithoutIsStrInference([553])))), d_308_v132_)), d_166_v1_, d_165_v0_, not(d_165_v0_), d_164_globalState_)
            d_330_v153_ = out8_
            d_331_v154_ = out9_
            d_332_v155_: T2
            nw68_ = C4()
            nw68_.ctor__(d_330_v153_)
            d_332_v155_ = nw68_
            d_333_v156_: _dafny.Array
            nw69_ = _dafny.Array(None, 2)
            nw69_[int(0)] = d_332_v155_
            nw69_[int(1)] = d_332_v155_
            d_333_v156_ = nw69_
            d_333_v156_ = d_333_v156_
            d_334_v157_: C7
            nw70_ = C7()
            nw70_.ctor__()
            d_334_v157_ = nw70_
            d_335_v158_: D20
            d_335_v158_ = D20_DC38(d_334_v157_)
            d_336_v159_: _dafny.Set
            d_336_v159_ = _dafny.Set({(d_335_v158_ if d_165_v0_ else D20_DC38(d_334_v157_))})
            d_337_v160_: _dafny.MultiSet
            d_337_v160_ = _dafny.MultiSet([d_165_v0_, d_165_v0_])
            pat_let_tv5_ = d_336_v159_
            def iife29_(_pat_let6_0):
                def iife30_(d_338_dt__update__tmp_h1_):
                    def iife31_(_pat_let7_0):
                        def iife32_(d_339_dt__update_hcf62_h0_):
                            return D22_DC43(d_339_dt__update_hcf62_h0_)
                        return iife32_(_pat_let7_0)
                    return iife31_(pat_let_tv5_)
                return iife30_(_pat_let6_0)
            rhs41_ = not (d_165_v0_) or (not((d_337_v160_).issubset(d_337_v160_)))
            rhs42_ = (iife29_(D22_DC43(d_336_v159_))).cf62
            rhs43_ = d_165_v0_
            d_165_v0_ = rhs41_
            d_336_v159_ = rhs42_
            d_165_v0_ = rhs43_
            d_330_v153_ = default__.safeDivisionInt(d_331_v154_, d_308_v132_)
            d_340_v161_: _dafny.Seq
            d_340_v161_ = _dafny.SeqWithoutIsStrInference([d_165_v0_, not(d_165_v0_)])
            d_341_v162_: D22
            d_341_v162_ = D22_DC44(default__.fm41(_dafny.SeqWithoutIsStrInference([(d_337_v160_).cardinality for d_342_i6_ in range(default__.abs(930))]), d_165_v0_, (d_340_v161_)[default__.safeIndex(186, len(d_340_v161_))], d_165_v0_, d_164_globalState_))
            d_341_v162_ = d_341_v162_
        elif True:
            d_343_v163_: T2
            nw71_ = C2()
            nw71_.ctor__(d_308_v132_)
            d_343_v163_ = nw71_
            d_343_v163_ = d_343_v163_
            d_344_v164_: _dafny.Seq
            d_344_v164_ = _dafny.SeqWithoutIsStrInference([d_165_v0_])
            d_309_v133_ = (d_309_v133_).set((d_165_v0_ if d_165_v0_ else True), default__.safeModuloInt(len(d_344_v164_), d_308_v132_))
            d_345_v165_: _dafny.MultiSet
            d_345_v165_ = _dafny.MultiSet([d_165_v0_, d_165_v0_])
            d_346_v166_: _dafny.Array
            nw72_ = _dafny.Array(None, 23)
            nw72_[int(0)] = d_165_v0_
            nw72_[int(1)] = d_165_v0_
            nw72_[int(2)] = (D5_DC9(d_165_v0_, d_308_v132_, d_345_v165_)).cf14
            nw72_[int(3)] = d_165_v0_
            nw72_[int(4)] = d_165_v0_
            nw72_[int(5)] = d_165_v0_
            nw72_[int(6)] = not(d_165_v0_)
            nw72_[int(7)] = d_165_v0_
            nw72_[int(8)] = d_165_v0_
            nw72_[int(9)] = d_165_v0_
            nw72_[int(10)] = d_165_v0_
            nw72_[int(11)] = (d_344_v164_)[default__.safeIndex(d_308_v132_, len(d_344_v164_))]
            nw72_[int(12)] = d_165_v0_
            nw72_[int(13)] = d_165_v0_
            nw72_[int(14)] = d_165_v0_
            nw72_[int(15)] = d_165_v0_
            nw72_[int(16)] = d_165_v0_
            nw72_[int(17)] = True
            nw72_[int(18)] = d_165_v0_
            nw72_[int(19)] = True
            nw72_[int(20)] = not(d_165_v0_)
            nw72_[int(21)] = d_165_v0_
            nw72_[int(22)] = d_165_v0_
            d_346_v166_ = nw72_
            d_347_v167_: _dafny.Map
            d_347_v167_ = _dafny.Map({d_165_v0_: d_165_v0_})
            d_348_v168_: _dafny.Map
            d_348_v168_ = _dafny.Map({d_346_v166_: not(((default__.fm26(d_308_v132_, d_165_v0_, d_165_v0_, default__.fm5(d_164_globalState_), d_164_globalState_)).set(d_165_v0_, d_165_v0_)) == (d_347_v167_))})
            d_348_v168_ = (d_348_v168_).set(d_346_v166_, d_165_v0_)
            if d_165_v0_:
                index57_ = default__.safeIndex(538, (d_346_v166_).length(0))
                (d_346_v166_)[index57_] = d_165_v0_
                index58_ = default__.safeIndex(538, (d_346_v166_).length(0))
                (d_346_v166_)[index58_] = d_165_v0_
                d_165_v0_ = False
                index59_ = default__.safeIndex(538, (d_346_v166_).length(0))
                rhs44_ = ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bn")))) * (d_308_v132_)) != (d_343_v163_.f2)
                rhs45_ = d_343_v163_.f2
                rhs46_ = False
                lhs21_ = d_346_v166_
                lhs22_ = default__.safeIndex(538, (d_346_v166_).length(0))
                lhs23_ = d_343_v163_
                lhs21_[lhs22_] = rhs44_
                lhs23_.f2 = rhs45_
                d_165_v0_ = rhs46_
                d_349_v169_: D6
                d_349_v169_ = D6_DC12(d_165_v0_, (d_346_v166_)[default__.safeIndex(538, (d_346_v166_).length(0))])
                d_349_v169_ = d_349_v169_
                d_350_v170_: bool
                d_351_v171_: int
                out10_: bool
                out11_: int
                out10_, out11_ = (d_313_v137_).m6(d_343_v163_.f2, d_164_globalState_)
                d_350_v170_ = out10_
                d_351_v171_ = out11_
            elif True:
                d_352_v172_: D6
                d_352_v172_ = D6_DC13(True, d_308_v132_)
                d_165_v0_ = (d_352_v172_).cf24
                d_353_v173_: _dafny.Map
                d_353_v173_ = _dafny.Map({d_310_v134_: (0) - (d_343_v163_.f2)})
                d_354_v174_: _dafny.Map
                d_354_v174_ = d_353_v173_
                d_354_v174_ = d_353_v173_
                d_355_v175_: _dafny.Map
                d_355_v175_ = _dafny.Map({(d_343_v163_.f2) * (d_343_v163_.f2): d_310_v134_})
                d_310_v134_ = ((d_355_v175_)[d_308_v132_] if (d_308_v132_) in (d_355_v175_) else d_310_v134_)
                d_356_v176_: _dafny.Array
                def lambda11_(d_357_v134_):
                    def lambda12_(d_358_i7_):
                        return d_357_v134_

                    return lambda12_

                init6_ = lambda11_(d_310_v134_)
                nw73_ = _dafny.Array(None, 14)
                for i0_6_ in range(nw73_.length(0)):
                    nw73_[i0_6_] = init6_(i0_6_)
                d_356_v176_ = nw73_
                d_359_v177_: bool
                d_360_v178_: _dafny.Array
                d_361_v179_: int
                out12_: bool
                out13_: _dafny.Array
                out14_: int
                out12_, out13_, out14_ = (d_343_v163_).m1(d_343_v163_.f2, d_356_v176_, d_165_v0_, d_165_v0_, d_164_globalState_)
                d_359_v177_ = out12_
                d_360_v178_ = out13_
                d_361_v179_ = out14_
                d_362_v180_: C1
                nw74_ = C1()
                nw74_.ctor__((d_312_v136_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fifwva"))))
                d_362_v180_ = nw74_
                d_362_v180_ = d_362_v180_
            d_363_v181_: C0
            nw75_ = C0()
            nw75_.ctor__((d_343_v163_.f2) < ((0) - (d_308_v132_)), (d_344_v164_)[default__.safeIndex(d_343_v163_.f2, len(d_344_v164_))])
            d_363_v181_ = nw75_
        d_364_v182_: bool
        d_365_v183_: int
        out15_: bool
        out16_: int
        out15_, out16_ = (d_313_v137_).m6(d_308_v132_, d_164_globalState_)
        d_364_v182_ = out15_
        d_365_v183_ = out16_
        hi1_ = d_365_v183_
        for d_366_i8_ in range(d_365_v183_, hi1_):
            d_367_v184_: str
            d_367_v184_ = _dafny.CodePoint('j')
            d_368_v185_: D7
            d_368_v185_ = D7_DC15(d_308_v132_, d_367_v184_)
            d_369_v186_: T1
            nw76_ = C6()
            nw76_.ctor__()
            d_369_v186_ = nw76_
            pat_let_tv6_ = d_367_v184_
            d_370_v187_: _dafny.Map
            def iife33_(_pat_let8_0):
                def iife34_(d_371_dt__update__tmp_h2_):
                    def iife35_(_pat_let9_0):
                        def iife36_(d_372_dt__update_hcf28_h0_):
                            return D7_DC15((d_371_dt__update__tmp_h2_).cf27, d_372_dt__update_hcf28_h0_)
                        return iife36_(_pat_let9_0)
                    return iife35_(pat_let_tv6_)
                return iife34_(_pat_let8_0)
            d_370_v187_ = _dafny.Map({(iife33_(d_368_v185_)).cf27: d_369_v186_})
            d_370_v187_ = (d_370_v187_).set((0) - ((d_365_v183_ if True else d_366_i8_)), d_369_v186_)
            d_373_v188_: _dafny.Map
            d_373_v188_ = _dafny.Map({d_312_v136_: _dafny.Map({d_366_i8_: d_364_v182_})})
            d_374_v189_: _dafny.Seq
            d_374_v189_ = _dafny.SeqWithoutIsStrInference([d_364_v182_, False, d_165_v0_, False])
            d_375_v190_: D5
            d_375_v190_ = D5_DC10(d_165_v0_, d_365_v183_, (d_373_v188_) | (d_373_v188_), not ((d_374_v189_)[default__.safeIndex((d_310_v134_)[default__.safeIndex(d_365_v183_, len(d_310_v134_))], len(d_374_v189_))]) or (d_165_v0_))
            source7_ = d_375_v190_
            if source7_.is_DC9:
                d_376___mcc_h1_ = source7_.cf14
                d_377___mcc_h2_ = source7_.cf15
                d_378___mcc_h3_ = source7_.cf16
                d_379_cf16_ = d_378___mcc_h3_
                d_380_cf15_ = d_377___mcc_h2_
                d_381_cf14_ = d_376___mcc_h1_
                d_165_v0_ = False
                d_382_v191_: int
                out17_: int
                out17_ = (d_313_v137_).m2(d_308_v132_, 597, not (d_165_v0_) or (not(d_364_v182_)), d_164_globalState_)
                d_382_v191_ = out17_
                d_365_v183_ = ((d_308_v132_ if d_364_v182_ else d_365_v183_)) - ((d_369_v186_).fm1(True, d_164_globalState_))
                d_165_v0_ = (d_374_v189_)[default__.safeIndex(d_380_cf15_, len(d_374_v189_))]
            elif source7_.is_DC10:
                d_383___mcc_h4_ = source7_.cf17
                d_384___mcc_h5_ = source7_.cf18
                d_385___mcc_h6_ = source7_.cf19
                d_386___mcc_h7_ = source7_.cf20
                d_387_cf20_ = d_386___mcc_h7_
                d_388_cf19_ = d_385___mcc_h6_
                d_389_cf18_ = d_384___mcc_h5_
                d_390_cf17_ = d_383___mcc_h4_
                d_391_v192_: int
                out18_: int
                out18_ = (d_313_v137_).m2(d_308_v132_, 223, d_364_v182_, d_164_globalState_)
                d_391_v192_ = out18_
                d_311_v135_ = (d_311_v135_).set((d_366_i8_) - (d_389_cf18_), (d_165_v0_ if d_390_cf17_ else d_390_cf17_))
                d_313_v137_ = d_313_v137_
                d_313_v137_ = d_313_v137_
            elif True:
                d_392___mcc_h8_ = source7_.cf13
                d_393_cf13_ = d_392___mcc_h8_
                d_394_v193_: C0
                nw77_ = C0()
                nw77_.ctor__(True, d_165_v0_)
                d_394_v193_ = nw77_
                index60_ = default__.safeIndex(441, (d_166_v1_).length(0))
                (d_166_v1_)[index60_] = d_366_i8_
                index61_ = default__.safeIndex(441, (d_166_v1_).length(0))
                (d_166_v1_)[index61_] = default__.safeModuloInt(d_365_v183_, -579)
                d_364_v182_ = False
                d_395_v194_: _dafny.Set
                d_395_v194_ = _dafny.Set({d_394_v193_.f4, d_394_v193_.f4, d_394_v193_.f4, d_394_v193_.f4})
                d_395_v194_ = (d_395_v194_).intersection(d_395_v194_)
            d_396_v195_: C4
            nw78_ = C4()
            nw78_.ctor__(((d_309_v133_)[False] if (False) in (d_309_v133_) else 988))
            d_396_v195_ = nw78_
            nw79_ = C4()
            nw79_.ctor__(d_365_v183_)
            d_396_v195_ = nw79_
            d_397_v196_: C7
            nw80_ = C7()
            nw80_.ctor__()
            d_397_v196_ = nw80_
            d_397_v196_ = d_397_v196_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_166_v1_).length(0)):
            d_398_i9_: int = guard_loop_0_
            if (True) and (((0) <= (d_398_i9_)) and ((d_398_i9_) < ((d_166_v1_).length(0)))):
                (d_166_v1_)[(d_398_i9_)] = (d_398_i9_) + (d_308_v132_)
        if (default__.safeModuloInt(d_365_v183_, d_365_v183_)) <= (d_365_v183_):
            if d_165_v0_:
                d_399_v197_: T0
                nw81_ = C6()
                nw81_.ctor__()
                d_399_v197_ = nw81_
                d_400_v198_: T0
                d_400_v198_ = d_399_v197_
                d_401_v199_: _dafny.Array
                nw82_ = _dafny.Array(None, 15)
                nw82_[int(0)] = d_400_v198_
                nw82_[int(1)] = d_400_v198_
                nw82_[int(2)] = d_399_v197_
                nw82_[int(3)] = d_400_v198_
                nw82_[int(4)] = d_400_v198_
                nw82_[int(5)] = d_400_v198_
                nw82_[int(6)] = d_400_v198_
                nw82_[int(7)] = d_400_v198_
                nw82_[int(8)] = d_400_v198_
                nw82_[int(9)] = d_399_v197_
                nw82_[int(10)] = d_400_v198_
                nw82_[int(11)] = d_400_v198_
                nw82_[int(12)] = d_400_v198_
                nw82_[int(13)] = d_400_v198_
                nw82_[int(14)] = d_400_v198_
                d_401_v199_ = nw82_
                d_401_v199_ = d_401_v199_
                d_402_v200_: _dafny.Set
                d_402_v200_ = _dafny.Set({d_364_v182_})
                d_165_v0_ = (d_402_v200_).isdisjoint((d_402_v200_) - (d_402_v200_))
                d_403_v201_: D6
                d_403_v201_ = D6_DC11(d_312_v136_)
                d_312_v136_ = (d_312_v136_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhskj"))) + ((d_403_v201_).cf21))
                d_404_v202_: _dafny.Array
                nw83_ = _dafny.Array(_dafny.Set({}), 22)
                d_404_v202_ = nw83_
                d_404_v202_ = d_404_v202_
                d_405_v203_: _dafny.Array
                nw84_ = _dafny.Array(_dafny.Array(None, 0), 23)
                d_405_v203_ = nw84_
                index62_ = default__.safeIndex(492, (d_405_v203_).length(0))
                (d_405_v203_)[index62_] = d_166_v1_
                index63_ = default__.safeIndex(492, (d_405_v203_).length(0))
                (d_405_v203_)[index63_] = (d_166_v1_ if True else d_166_v1_)
            elif True:
                d_364_v182_ = not(((d_313_v137_).fm1(False, d_164_globalState_)) <= (default__.safeDivisionInt(d_308_v132_, 785)))
                d_406_v204_: _dafny.Array
                nw85_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_406_v204_ = nw85_
                d_407_v205_: _dafny.Array
                nw86_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_407_v205_ = nw86_
                index64_ = default__.safeIndex(601, (d_406_v204_).length(0))
                (d_406_v204_)[index64_] = d_407_v205_
                d_408_v206_: C7
                nw87_ = C7()
                nw87_.ctor__()
                d_408_v206_ = nw87_
                d_409_v207_: D20
                d_409_v207_ = D20_DC38(d_408_v206_)
                d_410_v208_: _dafny.Set
                d_410_v208_ = _dafny.Set({d_409_v207_, d_409_v207_, d_409_v207_})
                d_411_v209_: D22
                d_411_v209_ = D22_DC43(d_410_v208_)
                d_412_v210_: _dafny.Seq
                d_412_v210_ = _dafny.SeqWithoutIsStrInference([d_411_v209_])
                d_413_v211_: C4
                nw88_ = C4()
                nw88_.ctor__(d_308_v132_)
                d_413_v211_ = nw88_
                d_414_v212_: _dafny.Map
                d_414_v212_ = _dafny.Map({(d_310_v134_)[default__.safeIndex(d_308_v132_, len(d_310_v134_))]: d_413_v211_})
                d_415_v213_: _dafny.Seq
                d_415_v213_ = _dafny.SeqWithoutIsStrInference([d_413_v211_, d_413_v211_, d_413_v211_])
                d_416_v214_: _dafny.Array
                nw89_ = _dafny.Array(None, 29)
                nw89_[int(0)] = d_413_v211_
                nw89_[int(1)] = d_413_v211_
                nw89_[int(2)] = d_413_v211_
                nw89_[int(3)] = d_413_v211_
                nw89_[int(4)] = (d_413_v211_ if d_165_v0_ else d_413_v211_)
                nw89_[int(5)] = d_413_v211_
                nw89_[int(6)] = d_413_v211_
                nw89_[int(7)] = d_413_v211_
                nw89_[int(8)] = ((d_414_v212_)[d_308_v132_] if (d_308_v132_) in (d_414_v212_) else d_413_v211_)
                nw89_[int(9)] = d_413_v211_
                nw89_[int(10)] = (d_415_v213_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orqkbxy"))), len(d_415_v213_))]
                nw89_[int(11)] = d_413_v211_
                nw89_[int(12)] = d_413_v211_
                nw89_[int(13)] = d_413_v211_
                nw89_[int(14)] = d_413_v211_
                nw89_[int(15)] = d_413_v211_
                nw89_[int(16)] = d_413_v211_
                nw89_[int(17)] = d_413_v211_
                nw89_[int(18)] = d_413_v211_
                nw89_[int(19)] = d_413_v211_
                nw89_[int(20)] = d_413_v211_
                nw89_[int(21)] = d_413_v211_
                nw89_[int(22)] = d_413_v211_
                nw89_[int(23)] = d_413_v211_
                nw89_[int(24)] = d_413_v211_
                nw89_[int(25)] = d_413_v211_
                nw89_[int(26)] = d_413_v211_
                nw89_[int(27)] = d_413_v211_
                nw89_[int(28)] = d_413_v211_
                d_416_v214_ = nw89_
                index65_ = default__.safeIndex(92, (d_416_v214_).length(0))
                (d_416_v214_)[index65_] = d_413_v211_
                d_417_v215_: str
                d_417_v215_ = _dafny.CodePoint('r')
                d_418_v216_: _dafny.MultiSet
                d_418_v216_ = _dafny.MultiSet([d_364_v182_])
                d_419_v217_: _dafny.Map
                d_419_v217_ = _dafny.Map({d_418_v216_: d_364_v182_})
                d_420_v218_: D21
                d_420_v218_ = D21_DC39(d_413_v211_)
                index66_ = default__.safeIndex(601, (d_406_v204_).length(0))
                index67_ = default__.safeIndex(92, (d_416_v214_).length(0))
                rhs47_ = d_407_v205_
                rhs48_ = len((default__.fm55(True, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_421_i10_ in range(default__.abs(506))]), d_417_v215_, d_164_globalState_)) | ((_dafny.Map({d_418_v216_: True})) | (d_419_v217_)))
                rhs49_ = default__.fm54(d_165_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmnormpf")), d_164_globalState_)
                rhs50_ = d_412_v210_
                rhs51_ = (d_413_v211_ if d_364_v182_ else (d_420_v218_).cf58)
                lhs24_ = d_406_v204_
                lhs25_ = default__.safeIndex(601, (d_406_v204_).length(0))
                lhs26_ = d_416_v214_
                lhs27_ = default__.safeIndex(92, (d_416_v214_).length(0))
                lhs24_[lhs25_] = rhs47_
                d_365_v183_ = rhs48_
                d_308_v132_ = rhs49_
                d_412_v210_ = rhs50_
                lhs26_[lhs27_] = rhs51_
                d_365_v183_ = (d_413_v211_).fm1(d_165_v0_, d_164_globalState_)
                d_422_v219_: _dafny.Map
                d_422_v219_ = _dafny.Map({d_166_v1_: d_364_v182_})
                rhs52_ = d_364_v182_
                rhs53_ = d_364_v182_
                rhs54_ = ((d_422_v219_) | (d_422_v219_)) | (d_422_v219_)
                d_165_v0_ = rhs52_
                d_165_v0_ = rhs53_
                d_422_v219_ = rhs54_
                d_423_v220_: _dafny.Array
                def lambda13_(d_424_v183_, d_425_v135_, d_426_v0_):
                    def lambda14_(d_427_i11_):
                        return ((d_425_v135_)[d_424_v183_] if (d_424_v183_) in (d_425_v135_) else d_426_v0_)

                    return lambda14_

                init7_ = lambda13_(d_365_v183_, d_311_v135_, d_165_v0_)
                nw90_ = _dafny.Array(None, 12)
                for i0_7_ in range(nw90_.length(0)):
                    nw90_[i0_7_] = init7_(i0_7_)
                d_423_v220_ = nw90_
                index68_ = default__.safeIndex(92, (d_423_v220_).length(0))
                (d_423_v220_)[index68_] = d_165_v0_
                index69_ = default__.safeIndex(92, (d_423_v220_).length(0))
                (d_423_v220_)[index69_] = d_364_v182_
            d_428_v221_: _dafny.Array
            nw91_ = _dafny.Array(_dafny.MultiSet({}), 19)
            d_428_v221_ = nw91_
            d_429_v222_: T2
            nw92_ = C4()
            nw92_.ctor__(d_365_v183_)
            d_429_v222_ = nw92_
            d_430_v223_: T2
            d_430_v223_ = d_429_v222_
            index70_ = default__.safeIndex(242, (d_428_v221_).length(0))
            (d_428_v221_)[index70_] = _dafny.MultiSet([d_430_v223_, d_430_v223_])
            d_431_v224_: _dafny.Set
            d_431_v224_ = _dafny.Set({d_429_v222_.f2, d_429_v222_.f2, 658, d_429_v222_.f2, (d_365_v183_) * ((0) - ((default__.fm38(d_165_v0_, d_164_globalState_)).cardinality))})
            d_432_v225_: _dafny.MultiSet
            d_432_v225_ = _dafny.MultiSet([d_430_v223_])
            index71_ = default__.safeIndex(242, (d_428_v221_).length(0))
            rhs55_ = len(d_431_v224_)
            rhs56_ = d_432_v225_
            lhs28_ = d_428_v221_
            lhs29_ = default__.safeIndex(242, (d_428_v221_).length(0))
            d_365_v183_ = rhs55_
            lhs28_[lhs29_] = rhs56_
            d_433_v226_: D21
            d_433_v226_ = D21_DC41(d_364_v182_)
            d_434_v227_: C3
            nw93_ = C3()
            nw93_.ctor__((d_310_v134_)[default__.safeIndex(len(default__.fm36(d_364_v182_, _dafny.CodePoint('i'), d_308_v132_, d_164_globalState_)), len(d_310_v134_))])
            d_434_v227_ = nw93_
            d_435_v228_: _dafny.Map
            d_435_v228_ = _dafny.Map({(d_433_v226_).cf60: d_434_v227_})
            d_436_v229_: _dafny.Seq
            d_436_v229_ = _dafny.SeqWithoutIsStrInference([d_434_v227_])
            d_435_v228_ = (d_435_v228_).set((d_165_v0_ if d_165_v0_ else d_165_v0_), ((d_435_v228_)[False] if (False) in (d_435_v228_) else (d_436_v229_)[default__.safeIndex(d_308_v132_, len(d_436_v229_))]))
            d_437_v230_: str
            d_437_v230_ = _dafny.CodePoint('w')
            d_438_v231_: D20
            d_438_v231_ = D20_DC37(d_437_v230_, not(d_165_v0_))
            d_439_v232_: _dafny.Map
            d_439_v232_ = _dafny.Map({(0) - ((0) - (d_308_v132_)): d_438_v231_})
            d_440_v233_: _dafny.MultiSet
            d_440_v233_ = _dafny.MultiSet([d_312_v136_])
            pat_let_tv7_ = d_364_v182_
            def iife37_(_pat_let10_0):
                def iife38_(d_441_dt__update__tmp_h4_):
                    def iife39_(_pat_let11_0):
                        def iife40_(d_442_dt__update_hcf56_h0_):
                            return D20_DC37((d_441_dt__update__tmp_h4_).cf55, d_442_dt__update_hcf56_h0_)
                        return iife40_(_pat_let11_0)
                    return iife39_(pat_let_tv7_)
                return iife38_(_pat_let10_0)
            d_439_v232_ = (d_439_v232_).set(((d_440_v233_)[d_312_v136_] if (d_312_v136_) in (d_440_v233_) else d_429_v222_.f2), iife37_(d_438_v231_))
            d_443_v234_: _dafny.Array
            nw94_ = _dafny.Array(None, 9)
            nw94_[int(0)] = d_364_v182_
            nw94_[int(1)] = d_165_v0_
            nw94_[int(2)] = d_165_v0_
            nw94_[int(3)] = d_165_v0_
            nw94_[int(4)] = d_165_v0_
            nw94_[int(5)] = d_364_v182_
            nw94_[int(6)] = not(d_364_v182_)
            nw94_[int(7)] = d_364_v182_
            nw94_[int(8)] = d_364_v182_
            d_443_v234_ = nw94_
            d_444_v235_: _dafny.Map
            d_444_v235_ = _dafny.Map({d_364_v182_: d_443_v234_})
            d_445_v236_: _dafny.Array
            nw95_ = _dafny.Array(None, 16)
            nw95_[int(0)] = d_443_v234_
            nw95_[int(1)] = d_443_v234_
            nw95_[int(2)] = d_443_v234_
            nw95_[int(3)] = d_443_v234_
            nw95_[int(4)] = d_443_v234_
            nw95_[int(5)] = d_443_v234_
            nw95_[int(6)] = d_443_v234_
            nw95_[int(7)] = d_443_v234_
            nw95_[int(8)] = d_443_v234_
            nw95_[int(9)] = d_443_v234_
            nw95_[int(10)] = d_443_v234_
            nw95_[int(11)] = ((d_444_v235_)[True] if (True) in (d_444_v235_) else d_443_v234_)
            nw95_[int(12)] = d_443_v234_
            nw95_[int(13)] = d_443_v234_
            nw95_[int(14)] = d_443_v234_
            nw95_[int(15)] = (d_443_v234_ if not(d_165_v0_) else d_443_v234_)
            d_445_v236_ = nw95_
            index72_ = default__.safeIndex(525, (d_445_v236_).length(0))
            (d_445_v236_)[index72_] = d_443_v234_
            index73_ = default__.safeIndex(525, (d_445_v236_).length(0))
            (d_445_v236_)[index73_] = d_443_v234_
        elif True:
            d_446_v237_: _dafny.Array
            nw96_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 12)
            d_446_v237_ = nw96_
            d_447_v238_: _dafny.Array
            nw97_ = _dafny.Array(None, 9)
            nw97_[int(0)] = d_446_v237_
            nw97_[int(1)] = d_446_v237_
            nw97_[int(2)] = d_446_v237_
            nw97_[int(3)] = d_446_v237_
            nw97_[int(4)] = d_446_v237_
            nw97_[int(5)] = d_446_v237_
            nw97_[int(6)] = d_446_v237_
            nw97_[int(7)] = d_446_v237_
            nw97_[int(8)] = d_446_v237_
            d_447_v238_ = nw97_
            index74_ = default__.safeIndex(655, (d_447_v238_).length(0))
            (d_447_v238_)[index74_] = d_446_v237_
            d_448_v239_: _dafny.Seq
            d_448_v239_ = _dafny.SeqWithoutIsStrInference([d_446_v237_])
            index75_ = default__.safeIndex(655, (d_447_v238_).length(0))
            (d_447_v238_)[index75_] = (d_448_v239_)[default__.safeIndex(len(d_312_v136_), len(d_448_v239_))]
            d_365_v183_ = (0) - (((len(d_310_v134_)) * (d_365_v183_)) * ((_dafny.MultiSet([(default__.fm56(d_165_v0_, d_164_globalState_)).cardinality, d_365_v183_, len(d_312_v136_), d_308_v132_])).cardinality))
            d_449_v240_: str
            d_449_v240_ = _dafny.CodePoint('a')
            d_449_v240_ = d_449_v240_
            d_449_v240_ = d_449_v240_
            index76_ = default__.safeIndex(236, (d_166_v1_).length(0))
            (d_166_v1_)[index76_] = 390
            d_450_v241_: _dafny.MultiSet
            d_450_v241_ = _dafny.MultiSet([d_449_v240_, d_449_v240_])
            index77_ = default__.safeIndex(236, (d_166_v1_).length(0))
            rhs57_ = (0) - (d_365_v183_)
            rhs58_ = (d_449_v240_) not in (d_450_v241_)
            rhs59_ = (0) - (d_365_v183_)
            lhs30_ = d_166_v1_
            lhs31_ = default__.safeIndex(236, (d_166_v1_).length(0))
            d_308_v132_ = rhs57_
            d_165_v0_ = rhs58_
            lhs30_[lhs31_] = rhs59_
        d_308_v132_ = d_365_v183_
        d_451_v242_: _dafny.Array
        nw98_ = _dafny.Array(False, 28)
        d_451_v242_ = nw98_
        index78_ = default__.safeIndex(396, (d_451_v242_).length(0))
        (d_451_v242_)[index78_] = default__.fm3(d_165_v0_, d_308_v132_, d_308_v132_, d_164_globalState_)
        index79_ = default__.safeIndex(396, (d_451_v242_).length(0))
        (d_451_v242_)[index79_] = (d_312_v136_) != (d_312_v136_)
        _dafny.print(_dafny.string_of(d_165_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v1_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v1_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v1_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v1_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v1_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v1_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v1_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v1_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v1_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v1_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_192_v27_)) == (_dafny.Set({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_308_v132_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_v133_) == (_dafny.Map({False: 786, True: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_310_v134_) == (_dafny.SeqWithoutIsStrInference([786, 390]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v135_) == (_dafny.Map({786: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_312_v136_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_314_v138_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_364_v182_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_365_v183_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_451_v242_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, False, _dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

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


class D1:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC2(D1, NamedTuple('DC2', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)

class D2_DC3(D2, NamedTuple('DC3', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D3_DC4)

class D3_DC4(D3, NamedTuple('DC4', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC4({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC4) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC6(int(0), _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D4_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D4_DC5)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D4_DC7)

class D4_DC6(D4, NamedTuple('DC6', [('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC6({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC6) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC5(D4, NamedTuple('DC5', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC5({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC5) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC7(D4, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC9(False, int(0), _dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D5_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D5_DC10)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D5_DC8)

class D5_DC9(D5, NamedTuple('DC9', [('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC9({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC9) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC10(D5, NamedTuple('DC10', [('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC10({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC10) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC8(D5, NamedTuple('DC8', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC8({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC8) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC12(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D6_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D6_DC11)

class D6_DC12(D6, NamedTuple('DC12', [('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC12({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC12) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC13(D6, NamedTuple('DC13', [('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC11(D6, NamedTuple('DC11', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC11({self.cf21.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC11) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC15(int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D7_DC14)

class D7_DC15(D7, NamedTuple('DC15', [('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC14(D7, NamedTuple('DC14', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC14({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC14) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC17(False, None, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D8_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D8_DC16)

class D8_DC17(D8, NamedTuple('DC17', [('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC17({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC17) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC16(D8, NamedTuple('DC16', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC16({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC16) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D9_DC18)

class D9_DC18(D9, NamedTuple('DC18', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC18({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC18) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D10_DC19)

class D10_DC19(D10, NamedTuple('DC19', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC19({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC19) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC21(_dafny.CodePoint('D'), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D11_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D11_DC22)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D11_DC20)

class D11_DC21(D11, NamedTuple('DC21', [('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC21({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC21) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC22(D11, NamedTuple('DC22', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC22({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC22) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC20(D11, NamedTuple('DC20', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC20({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC20) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D12_DC23)

class D12_DC23(D12, NamedTuple('DC23', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC23({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC23) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D13_DC24)

class D13_DC24(D13, NamedTuple('DC24', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC24({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC24) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC26(D0.default()(), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D14_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D14_DC27)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D14_DC28)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D14_DC25)

class D14_DC26(D14, NamedTuple('DC26', [('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC26({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC26) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC27(D14, NamedTuple('DC27', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC27({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC27) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC28(D14, NamedTuple('DC28', [])):
    def __dafnystr__(self) -> str:
        return f'D14.DC28'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC28)
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC25(D14, NamedTuple('DC25', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC25({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC25) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC30(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D15_DC30)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D15_DC29)

class D15_DC30(D15, NamedTuple('DC30', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC30({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC30) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC29(D15, NamedTuple('DC29', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC29({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC29) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D16_DC31)

class D16_DC31(D16, NamedTuple('DC31', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC31({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC31) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D17_DC32)

class D17_DC32(D17, NamedTuple('DC32', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC32({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC32) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC34(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D18_DC34)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D18_DC33)

class D18_DC34(D18, NamedTuple('DC34', [('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC34({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC34) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC33(D18, NamedTuple('DC33', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC33({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC33) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D19_DC35)

class D19_DC35(D19, NamedTuple('DC35', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC35({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC35) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC37(_dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D20_DC37)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D20_DC38)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D20_DC36)

class D20_DC37(D20, NamedTuple('DC37', [('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC37({_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC37) and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC38(D20, NamedTuple('DC38', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC38({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC38) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC36(D20, NamedTuple('DC36', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC36({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC36) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC40(None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D21_DC40)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D21_DC41)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D21_DC39)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D21_DC42)

class D21_DC40(D21, NamedTuple('DC40', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC40({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC40) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC41(D21, NamedTuple('DC41', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC41({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC41) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC39(D21, NamedTuple('DC39', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC39({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC39) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC42(D21, NamedTuple('DC42', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC42({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC42) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC44(_dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D22_DC44)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D22_DC45)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D22_DC43)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D22_DC46)

class D22_DC44(D22, NamedTuple('DC44', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC44({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC44) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC45(D22, NamedTuple('DC45', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC45({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC45) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC43(D22, NamedTuple('DC43', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC43({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC43) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC46(D22, NamedTuple('DC46', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC46({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC46) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D23_DC47)

class D23_DC47(D23, NamedTuple('DC47', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC47({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC47) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC49(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D24_DC49)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D24_DC48)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D24_DC50)

class D24_DC49(D24, NamedTuple('DC49', [('cf68', Any), ('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC49({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC49) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC48(D24, NamedTuple('DC48', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC48({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC48) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC50(D24, NamedTuple('DC50', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC50({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC50) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, p1, p2, p3, globalState):
        pass

    def m2(self, p0, p1, p2, globalState):
        pass


class T1:
    pass
    def m3(self, p0, p1, p2, p3, globalState):
        pass


class T2:
    pass
    @property
    def f2(self):
        return self._f2
    @f2.setter
    def f2(self, value):
        self._f2 = value

class GlobalState:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self):
        pass
        pass


class C0:
    def  __init__(self):
        self.f4: bool = False
        self.f5: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f4, f5):
        (self).f4 = f4
        (self).f5 = f5

    def fm14(self, p0, p1, p2, globalState):
        return self.f4


class C1(T1, T0):
    def  __init__(self):
        self._f3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f3):
        (self)._f3 = f3

    def fm0(self, globalState):
        def iife41_():
            coll17_ = _dafny.Map()
            compr_17_: _dafny.Seq
            for compr_17_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality for d_453_i1_ in range(default__.abs(60))]): 352})).keys.Elements:
                d_454_v0_: _dafny.Seq = compr_17_
                if (d_454_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality for d_453_i1_ in range(default__.abs(60))]): 352})):
                    coll17_[d_454_v0_] = 894
            return _dafny.Map(coll17_)
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([len((self).f3) for d_452_i0_ in range(default__.abs(230))]): len((self).f3)})) | ((iife41_()
         if False else _dafny.Map({_dafny.SeqWithoutIsStrInference([266]): -986})))

    def fm1(self, p0, globalState):
        return (len((_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([True])))) + (len((_dafny.SeqWithoutIsStrInference([False, False]) if not(False) else _dafny.SeqWithoutIsStrInference([not(not(False)), False]))))

    def fm11(self, globalState):
        return (_dafny.Set({True})).isdisjoint((_dafny.Set({True})).intersection(_dafny.Set({False})))

    def fm12(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([-906])) + (_dafny.SeqWithoutIsStrInference([(0) - (-299), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ob"))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_455_i0_ in range(default__.abs(730))]))]))

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_456_v0_: _dafny.Array
        nw99_ = _dafny.Array(_dafny.Array(None, 0), 29)
        d_456_v0_ = nw99_
        d_457_v1_: _dafny.Array
        nw100_ = _dafny.Array(False, 19)
        d_457_v1_ = nw100_
        index80_ = default__.safeIndex(574, (d_456_v0_).length(0))
        (d_456_v0_)[index80_] = d_457_v1_
        index81_ = default__.safeIndex(574, (d_456_v0_).length(0))
        (d_456_v0_)[index81_] = d_457_v1_
        d_458_v2_: str
        d_458_v2_ = _dafny.CodePoint('t')
        d_459_v3_: int
        d_459_v3_ = 469
        d_460_v4_: D0
        d_460_v4_ = D0_DC1(p3, p3, d_458_v2_, len(_dafny.Map({(self).f3: d_459_v3_})))
        pat_let_tv8_ = d_459_v3_
        pat_let_tv9_ = d_459_v3_
        def lambda15_(source8_):
            if source8_.is_DC1:
                d_461___mcc_h0_ = source8_.cf1
                d_462___mcc_h1_ = source8_.cf2
                d_463___mcc_h2_ = source8_.cf3
                d_464___mcc_h3_ = source8_.cf4
                d_465_cf4_ = d_464___mcc_h3_
                d_466_cf3_ = d_463___mcc_h2_
                d_467_cf2_ = d_462___mcc_h1_
                d_468_cf1_ = d_461___mcc_h0_
                return pat_let_tv8_
            elif True:
                d_469___mcc_h4_ = source8_.cf0
                d_470_cf0_ = d_469___mcc_h4_
                return pat_let_tv9_

        r0 = lambda15_(d_460_v4_)
        d_471_v5_: bool
        d_471_v5_ = False
        d_472_v6_: _dafny.Array
        def lambda16_(d_473_p3_, d_474_v5_):
            def lambda17_(d_475_i0_):
                return (_dafny.MultiSet([d_473_p3_, d_474_v5_])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_473_p3_])))

            return lambda17_

        init8_ = lambda16_(p3, d_471_v5_)
        nw101_ = _dafny.Array(None, 27)
        for i0_8_ in range(nw101_.length(0)):
            nw101_[i0_8_] = init8_(i0_8_)
        d_472_v6_ = nw101_
        d_476_v7_: _dafny.MultiSet
        d_476_v7_ = _dafny.MultiSet([p2, d_471_v5_])
        index82_ = default__.safeIndex(632, (d_472_v6_).length(0))
        (d_472_v6_)[index82_] = d_476_v7_
        d_477_v8_: _dafny.Map
        d_477_v8_ = _dafny.Map({False: default__.fm13(globalState)})
        d_478_v9_: _dafny.Map
        d_478_v9_ = _dafny.Map({d_459_v3_: p3})
        pat_let_tv10_ = p3
        index83_ = default__.safeIndex(632, (d_472_v6_).length(0))
        def lambda18_(source9_):
            if source9_.is_DC1:
                d_479___mcc_h5_ = source9_.cf1
                d_480___mcc_h6_ = source9_.cf2
                d_481___mcc_h7_ = source9_.cf3
                d_482___mcc_h8_ = source9_.cf4
                d_483_cf4_ = d_482___mcc_h8_
                d_484_cf3_ = d_481___mcc_h7_
                d_485_cf2_ = d_480___mcc_h6_
                d_486_cf1_ = d_479___mcc_h5_
                return pat_let_tv10_
            elif True:
                d_487___mcc_h9_ = source9_.cf0
                d_488_cf0_ = d_487___mcc_h9_
                return d_488_cf0_

        rhs60_ = d_458_v2_
        rhs61_ = lambda18_(D0_DC1(p3, d_471_v5_, d_458_v2_, d_459_v3_))
        rhs62_ = (d_476_v7_) | (d_476_v7_)
        rhs63_ = (len(_dafny.SeqWithoutIsStrInference([((d_477_v8_)[d_471_v5_] if (d_471_v5_) in (d_477_v8_) else d_478_v9_)]))) * ((0) - (d_459_v3_))
        lhs32_ = d_472_v6_
        lhs33_ = default__.safeIndex(632, (d_472_v6_).length(0))
        d_458_v2_ = rhs60_
        d_471_v5_ = rhs61_
        lhs32_[lhs33_] = rhs62_
        r1 = rhs63_
        d_489_v10_: _dafny.Seq
        d_489_v10_ = _dafny.SeqWithoutIsStrInference([18, d_459_v3_, default__.safeDivisionInt(d_459_v3_, d_459_v3_)])
        d_490_v11_: _dafny.Seq
        d_490_v11_ = _dafny.SeqWithoutIsStrInference([d_471_v5_, True])
        d_489_v10_ = (((self).fm12(p3, d_471_v5_, (self).fm1((d_490_v11_)[default__.safeIndex(len(_dafny.Map({p3: d_459_v3_})), len(d_490_v11_))], globalState), d_471_v5_, globalState)).set(default__.safeIndex(d_459_v3_, len((self).fm12(p3, d_471_v5_, (self).fm1((d_490_v11_)[default__.safeIndex(len(_dafny.Map({p3: d_459_v3_})), len(d_490_v11_))], globalState), d_471_v5_, globalState))), d_459_v3_)) + ((self).fm12((self).fm11(globalState), p3, d_459_v3_, d_471_v5_, globalState))
        d_491_v12_: C0
        nw102_ = C0()
        nw102_.ctor__((d_459_v3_) >= (d_459_v3_), p2)
        d_491_v12_ = nw102_
        if p2:
            d_492_v13_: _dafny.Map
            d_492_v13_ = _dafny.Map({d_471_v5_: True})
            source10_ = (d_492_v13_) | (d_492_v13_)
            d_493___mcc_h10_ = source10_
            d_494_cf5_ = d_493___mcc_h10_
            d_495_v14_: _dafny.Map
            d_495_v14_ = _dafny.Map({d_459_v3_: d_459_v3_})
            d_459_v3_ = default__.safeModuloInt(d_459_v3_, ((d_495_v14_)[d_459_v3_] if (d_459_v3_) in (d_495_v14_) else 583))
            d_496_v15_: _dafny.Map
            d_496_v15_ = _dafny.Map({d_458_v2_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmd"))})
            d_496_v15_ = (d_496_v15_).set(d_458_v2_, ((self).f3) + ((self).f3))
            d_497_v16_: _dafny.Set
            d_497_v16_ = _dafny.Set({(d_456_v0_)[default__.safeIndex(574, (d_456_v0_).length(0))]})
            d_498_v17_: _dafny.Map
            d_498_v17_ = _dafny.Map({d_491_v12_.f5: len(d_497_v16_)})
            d_499_v18_: _dafny.Map
            d_499_v18_ = (d_492_v13_).set(d_491_v12_.f5, p2)
            d_500_v19_: _dafny.Seq
            d_500_v19_ = _dafny.SeqWithoutIsStrInference([d_494_cf5_, d_499_v18_])
            index84_ = default__.safeIndex(998, (p1).length(0))
            (p1)[index84_] = (len(d_498_v17_)) + (len(d_500_v19_))
            index85_ = default__.safeIndex(998, (p1).length(0))
            (p1)[index85_] = (d_459_v3_) - ((d_489_v10_)[default__.safeIndex(d_459_v3_, len(d_489_v10_))])
            r0 = d_459_v3_
            d_501_v20_: _dafny.Array
            nw103_ = _dafny.Array(int(0), 3)
            d_501_v20_ = nw103_
            d_502_v21_: _dafny.Seq
            d_502_v21_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
            d_503_v22_: _dafny.Array
            nw104_ = _dafny.Array(None, 14)
            nw104_[int(0)] = d_501_v20_
            nw104_[int(1)] = d_501_v20_
            nw104_[int(2)] = p1
            nw104_[int(3)] = (d_502_v21_)[default__.safeIndex(d_459_v3_, len(d_502_v21_))]
            nw104_[int(4)] = p1
            nw104_[int(5)] = p1
            nw104_[int(6)] = d_501_v20_
            nw104_[int(7)] = p1
            nw104_[int(8)] = p1
            nw104_[int(9)] = d_501_v20_
            nw104_[int(10)] = d_501_v20_
            nw104_[int(11)] = d_501_v20_
            nw104_[int(12)] = d_501_v20_
            nw104_[int(13)] = d_501_v20_
            d_503_v22_ = nw104_
            d_504_v23_: _dafny.Array
            d_504_v23_ = d_503_v22_
            rhs64_ = d_456_v0_
            rhs65_ = d_501_v20_
            rhs66_ = (d_504_v23_)
            rhs67_ = (d_489_v10_) + (d_489_v10_)
            d_456_v0_ = rhs64_
            d_501_v20_ = rhs65_
            d_503_v22_ = rhs66_
            d_489_v10_ = rhs67_
            d_505_v24_: _dafny.Map
            d_505_v24_ = _dafny.Map({d_459_v3_: (d_456_v0_)[default__.safeIndex(574, (d_456_v0_).length(0))]})
            index86_ = default__.safeIndex(574, (d_456_v0_).length(0))
            (d_456_v0_)[index86_] = ((d_505_v24_)[d_459_v3_] if (d_459_v3_) in (d_505_v24_) else d_457_v1_)
            d_471_v5_ = not(((d_459_v3_) * (len((self).f3))) < (default__.safeDivisionInt(d_459_v3_, d_459_v3_)))
            d_471_v5_ = not(not((d_460_v4_).cf2))
        elif True:
            (d_491_v12_).f4 = not(d_471_v5_)
            (d_491_v12_).f5 = not(d_491_v12_.f5)
            r0 = d_459_v3_
            (d_491_v12_).f4 = not(d_491_v12_.f4)
            (d_491_v12_).f4 = p3
        r0 = default__.safeModuloInt(d_459_v3_, (d_459_v3_) + (d_459_v3_))
        r1 = (d_459_v3_) * (len((self).f3))
        return r0, r1

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        d_506_v0_: _dafny.MultiSet
        d_506_v0_ = _dafny.MultiSet([p2])
        if not((d_506_v0_).ispropersubset(((d_506_v0_).set(p3, default__.abs(p0))) - (d_506_v0_))):
            d_507_v1_: _dafny.Seq
            d_507_v1_ = _dafny.SeqWithoutIsStrInference([p0])
            d_508_v2_: str
            d_508_v2_ = _dafny.CodePoint('a')
            d_509_v3_: _dafny.Map
            d_509_v3_ = _dafny.Map({d_508_v2_: _dafny.CodePoint('q')})
            d_510_v4_: _dafny.MultiSet
            d_510_v4_ = _dafny.MultiSet([len(d_509_v3_)])
            def iife42_():
                coll18_ = _dafny.Set()
                compr_18_: int
                for compr_18_ in (d_510_v4_).Elements:
                    d_511_v5_: int = compr_18_
                    if (d_511_v5_) in (d_510_v4_):
                        coll18_ = coll18_.union(_dafny.Set([(d_511_v5_) + (315)]))
                return _dafny.Set(coll18_)
            r2 = (d_507_v1_)[default__.safeIndex(default__.safeDivisionInt(len(iife42_()
            ), p0), len(d_507_v1_))]
            d_512_v6_: _dafny.Set
            d_512_v6_ = _dafny.Set({p2})
            d_513_v7_: D4
            d_513_v7_ = D4_DC5(d_512_v6_)
            d_514_v8_: C0
            nw105_ = C0()
            nw105_.ctor__(p3, ((d_513_v7_).cf8) == (d_512_v6_))
            d_514_v8_ = nw105_
            r2 = (p0 if d_514_v8_.f4 else (p0) * (p0))
            d_515_v9_: C0
            nw106_ = C0()
            nw106_.ctor__(d_514_v8_.f5, d_514_v8_.f4)
            d_515_v9_ = nw106_
            d_516_v10_: _dafny.Seq
            d_516_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gragg"))
            d_516_v10_ = default__.fm15(globalState)
        elif True:
            d_517_v11_: _dafny.Array
            nw107_ = _dafny.Array(int(0), 16)
            d_517_v11_ = nw107_
            d_518_v12_: _dafny.Seq
            d_518_v12_ = _dafny.SeqWithoutIsStrInference([d_517_v11_, d_517_v11_])
            d_519_v13_: D4
            d_519_v13_ = D4_DC6(default__.safeDivisionInt(p0, p0), (d_518_v12_)[default__.safeIndex(len((self).f3), len(d_518_v12_))], (p0) - (p0))
            source11_ = d_519_v13_
            if source11_.is_DC6:
                d_520___mcc_h0_ = source11_.cf9
                d_521___mcc_h1_ = source11_.cf10
                d_522___mcc_h2_ = source11_.cf11
                d_523_cf11_ = d_522___mcc_h2_
                d_524_cf10_ = d_521___mcc_h1_
                d_525_cf9_ = d_520___mcc_h0_
                d_526_v14_: str
                d_526_v14_ = _dafny.CodePoint('q')
                d_527_v15_: D0
                d_527_v15_ = D0_DC1(p3, p2, d_526_v14_, 971)
                index87_ = default__.safeIndex(620, (d_517_v11_).length(0))
                (d_517_v11_)[index87_] = ((d_527_v15_).cf4 if p2 else d_525_cf9_)
                d_528_v16_: _dafny.Map
                d_528_v16_ = _dafny.Map({d_523_cf11_: len((self).f3)})
                d_529_v17_: _dafny.Map
                d_529_v17_ = _dafny.Map({p0: p2})
                d_530_v18_: _dafny.Seq
                d_530_v18_ = _dafny.SeqWithoutIsStrInference([p2, p2, (self).fm11(globalState), p3])
                index88_ = default__.safeIndex(620, (d_517_v11_).length(0))
                (d_517_v11_)[index88_] = (((d_528_v16_)[(0) - (len(d_529_v17_))] if ((0) - (len(d_529_v17_))) in (d_528_v16_) else d_525_cf9_) if not (p3) or (not(False)) else len((d_530_v18_) + (d_530_v18_)))
                d_531_v19_: _dafny.Array
                nw108_ = _dafny.Array(None, 13)
                nw108_[int(0)] = p2
                nw108_[int(1)] = p2
                nw108_[int(2)] = p3
                nw108_[int(3)] = p3
                nw108_[int(4)] = p2
                nw108_[int(5)] = p3
                nw108_[int(6)] = p2
                nw108_[int(7)] = p2
                nw108_[int(8)] = (d_530_v18_)[default__.safeIndex(len(_dafny.Set({254, d_525_cf9_, len(_dafny.SeqWithoutIsStrInference([(d_506_v0_).cardinality]))})), len(d_530_v18_))]
                nw108_[int(9)] = p2
                nw108_[int(10)] = p2
                nw108_[int(11)] = p3
                nw108_[int(12)] = p3
                d_531_v19_ = nw108_
                index89_ = default__.safeIndex(44, (d_531_v19_).length(0))
                (d_531_v19_)[index89_] = (self).fm11(globalState)
                index90_ = default__.safeIndex(44, (d_531_v19_).length(0))
                (d_531_v19_)[index90_] = False
                d_532_v20_: _dafny.Seq
                d_532_v20_ = _dafny.SeqWithoutIsStrInference([((d_528_v16_)[p0] if (p0) in (d_528_v16_) else (d_517_v11_)[default__.safeIndex(620, (d_517_v11_).length(0))])])
                d_532_v20_ = d_532_v20_
                index91_ = default__.safeIndex(44, (d_531_v19_).length(0))
                (d_531_v19_)[index91_] = (d_530_v18_)[default__.safeIndex((d_517_v11_)[default__.safeIndex(620, (d_517_v11_).length(0))], len(d_530_v18_))]
            elif source11_.is_DC5:
                d_533___mcc_h3_ = source11_.cf8
                d_534_cf8_ = d_533___mcc_h3_
                d_535_v21_: C0
                nw109_ = C0()
                nw109_.ctor__(p2, p2)
                d_535_v21_ = nw109_
                d_536_v22_: _dafny.Seq
                d_536_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
                d_537_v23_: _dafny.Array
                def lambda19_(d_538_i0_):
                    return True

                init9_ = lambda19_
                nw110_ = _dafny.Array(None, 5)
                for i0_9_ in range(nw110_.length(0)):
                    nw110_[i0_9_] = init9_(i0_9_)
                d_537_v23_ = nw110_
                d_539_v24_: _dafny.Array
                nw111_ = _dafny.Array(_dafny.Set({}), 29)
                d_539_v24_ = nw111_
                d_540_v25_: str
                d_540_v25_ = _dafny.CodePoint('n')
                d_541_v26_: D0
                d_541_v26_ = D0_DC1(d_535_v21_.f5, p3, d_540_v25_, p0)
                d_542_v27_: _dafny.Set
                d_542_v27_ = _dafny.Set({d_541_v26_, D0_DC1(d_535_v21_.f5, p3, d_540_v25_, p0), d_541_v26_, d_541_v26_})
                index92_ = default__.safeIndex(552, (d_539_v24_).length(0))
                (d_539_v24_)[index92_] = d_542_v27_
                index93_ = default__.safeIndex(552, (d_539_v24_).length(0))
                rhs68_ = (_dafny.SeqWithoutIsStrInference([d_540_v25_ for d_543_i1_ in range(default__.abs(979))])) + ((self).f3)
                rhs69_ = d_537_v23_
                rhs70_ = not(((len(d_536_v22_)) * (p0)) < (p0))
                rhs71_ = d_542_v27_
                rhs72_ = d_535_v21_.f5
                lhs34_ = d_535_v21_
                lhs35_ = d_539_v24_
                lhs36_ = default__.safeIndex(552, (d_539_v24_).length(0))
                d_536_v22_ = rhs68_
                d_537_v23_ = rhs69_
                lhs34_.f4 = rhs70_
                lhs35_[lhs36_] = rhs71_
                r0 = rhs72_
                d_544_v28_: _dafny.Array
                nw112_ = _dafny.Array(_dafny.Map({}), 4)
                d_544_v28_ = nw112_
                d_544_v28_ = d_544_v28_
                r0 = False
            elif True:
                d_545___mcc_h4_ = source11_.cf12
                d_546_cf12_ = d_545___mcc_h4_
                d_547_v29_: _dafny.Seq
                d_547_v29_ = _dafny.SeqWithoutIsStrInference([p0, default__.safeModuloInt(p0, (self).fm1(p2, globalState)), default__.safeModuloInt(p0, p0), p0, p0])
                d_547_v29_ = d_547_v29_
                r0 = not(p2)
                d_548_v30_: C0
                nw113_ = C0()
                nw113_.ctor__(p2, not(p2))
                d_548_v30_ = nw113_
                d_549_v31_: _dafny.Array
                nw114_ = _dafny.Array(False, 21)
                d_549_v31_ = nw114_
                index94_ = default__.safeIndex(386, (d_549_v31_).length(0))
                (d_549_v31_)[index94_] = d_548_v30_.f5
                index95_ = default__.safeIndex(386, (d_549_v31_).length(0))
                (d_549_v31_)[index95_] = p3
            r0 = not(p3)
            d_550_v32_: _dafny.Seq
            d_550_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nutlbe"))
            d_550_v32_ = ((self).f3) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdn")))
            r2 = p0
            d_551_v33_: _dafny.Array
            nw115_ = _dafny.Array(None, 8)
            nw115_[int(0)] = p2
            nw115_[int(1)] = p2
            nw115_[int(2)] = False
            nw115_[int(3)] = (self).fm11(globalState)
            nw115_[int(4)] = p2
            nw115_[int(5)] = (False if p2 else p2)
            nw115_[int(6)] = p3
            nw115_[int(7)] = p3
            d_551_v33_ = nw115_
            index96_ = default__.safeIndex(415, (d_551_v33_).length(0))
            (d_551_v33_)[index96_] = p3
            d_552_v34_: _dafny.Seq
            d_552_v34_ = _dafny.SeqWithoutIsStrInference([p2, p2])
            d_553_v35_: C0
            nw116_ = C0()
            nw116_.ctor__(False, p2)
            d_553_v35_ = nw116_
            d_554_v36_: _dafny.Seq
            d_554_v36_ = _dafny.SeqWithoutIsStrInference([d_553_v35_])
            index97_ = default__.safeIndex(415, (d_551_v33_).length(0))
            (d_551_v33_)[index97_] = not(((d_552_v34_)[default__.safeIndex(p0, len(d_552_v34_))] if p2 else (d_554_v36_) < (_dafny.SeqWithoutIsStrInference([d_553_v35_, d_553_v35_]))))
        d_555_v37_: _dafny.MultiSet
        d_555_v37_ = _dafny.MultiSet([p0])
        hi2_ = (d_555_v37_).cardinality
        for d_556_i2_ in range(p0, hi2_):
            d_557_v38_: _dafny.Map
            d_557_v38_ = _dafny.Map({p2: d_556_i2_})
            d_558_v39_: _dafny.Array
            nw117_ = _dafny.Array(None, 12)
            nw117_[int(0)] = -17
            nw117_[int(1)] = d_556_i2_
            nw117_[int(2)] = len(_dafny.SeqWithoutIsStrInference([p2]))
            nw117_[int(3)] = d_556_i2_
            nw117_[int(4)] = p0
            nw117_[int(5)] = d_556_i2_
            nw117_[int(6)] = len(d_557_v38_)
            nw117_[int(7)] = p0
            nw117_[int(8)] = d_556_i2_
            nw117_[int(9)] = d_556_i2_
            nw117_[int(10)] = d_556_i2_
            nw117_[int(11)] = d_556_i2_
            d_558_v39_ = nw117_
            d_559_v40_: _dafny.MultiSet
            d_559_v40_ = _dafny.MultiSet([d_558_v39_, d_558_v39_, d_558_v39_])
            d_560_v41_: D5
            d_560_v41_ = D5_DC8(d_559_v40_)
            d_561_v42_: _dafny.Array
            nw118_ = _dafny.Array(None, 21)
            nw118_[int(0)] = _dafny.MultiSet([d_558_v39_, d_558_v39_])
            nw118_[int(1)] = d_559_v40_
            nw118_[int(2)] = _dafny.MultiSet([d_558_v39_, d_558_v39_, d_558_v39_])
            nw118_[int(3)] = (d_559_v40_ if p3 else d_559_v40_)
            nw118_[int(4)] = _dafny.MultiSet([d_558_v39_, d_558_v39_, d_558_v39_, d_558_v39_])
            nw118_[int(5)] = d_559_v40_
            nw118_[int(6)] = (_dafny.MultiSet([d_558_v39_])).intersection(d_559_v40_)
            nw118_[int(7)] = (d_559_v40_) - (d_559_v40_)
            nw118_[int(8)] = _dafny.MultiSet([d_558_v39_, d_558_v39_, d_558_v39_])
            nw118_[int(9)] = d_559_v40_
            nw118_[int(10)] = d_559_v40_
            nw118_[int(11)] = _dafny.MultiSet([d_558_v39_])
            nw118_[int(12)] = d_559_v40_
            nw118_[int(13)] = d_559_v40_
            nw118_[int(14)] = d_559_v40_
            nw118_[int(15)] = (d_560_v41_).cf13
            nw118_[int(16)] = d_559_v40_
            nw118_[int(17)] = d_559_v40_
            nw118_[int(18)] = d_559_v40_
            nw118_[int(19)] = ((d_559_v40_).set(d_558_v39_, default__.abs(p0))) - (d_559_v40_)
            nw118_[int(20)] = d_559_v40_
            d_561_v42_ = nw118_
            index98_ = default__.safeIndex(811, (d_561_v42_).length(0))
            (d_561_v42_)[index98_] = d_559_v40_
            index99_ = default__.safeIndex(811, (d_561_v42_).length(0))
            (d_561_v42_)[index99_] = d_559_v40_
            d_562_v43_: _dafny.Map
            d_562_v43_ = _dafny.Map({p3: (d_506_v0_).intersection(d_506_v0_)})
            d_562_v43_ = (d_562_v43_) | ((d_562_v43_).set(p2, _dafny.MultiSet((default__.fm16(globalState)).set(default__.safeIndex(p0, len(default__.fm16(globalState))), p3))))
            d_563_v44_: str
            d_563_v44_ = _dafny.CodePoint('h')
            d_564_v45_: D0
            d_564_v45_ = D0_DC1(p2, p2, d_563_v44_, p0)
            d_565_v46_: _dafny.MultiSet
            d_565_v46_ = _dafny.MultiSet([d_564_v45_])
            d_566_v47_: _dafny.MultiSet
            d_566_v47_ = _dafny.MultiSet([d_563_v44_, d_563_v44_])
            d_567_v48_: _dafny.Set
            d_567_v48_ = _dafny.Set({(not(p2) if p3 else p3), (d_566_v47_).ispropersubset(default__.fm17((self).fm11(globalState), globalState)), True, p2})
            d_568_v49_: D5
            d_568_v49_ = D5_DC9(p2, -599, _dafny.MultiSet([p2, p3]))
            pat_let_tv11_ = p2
            def iife43_(_pat_let12_0):
                def iife44_(d_569_dt__update__tmp_h0_):
                    def iife45_(_pat_let13_0):
                        def iife46_(d_570_dt__update_hcf2_h0_):
                            return D0_DC1((d_569_dt__update__tmp_h0_).cf1, d_570_dt__update_hcf2_h0_, (d_569_dt__update__tmp_h0_).cf3, (d_569_dt__update__tmp_h0_).cf4)
                        return iife46_(_pat_let13_0)
                    return iife45_(pat_let_tv11_)
                return iife44_(_pat_let12_0)
            rhs73_ = _dafny.MultiSet([iife43_(d_564_v45_)])
            rhs74_ = d_567_v48_
            rhs75_ = d_556_i2_
            rhs76_ = ((self).f3)[default__.safeIndex(((d_568_v49_).cf15 if p3 else d_556_i2_), len((self).f3))]
            rhs77_ = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_556_i2_]))).isdisjoint(d_555_v37_)
            d_565_v46_ = rhs73_
            d_567_v48_ = rhs74_
            r2 = rhs75_
            d_563_v44_ = rhs76_
            r0 = rhs77_
            d_571_v50_: _dafny.Map
            d_571_v50_ = _dafny.Map({p0: p3})
            r2 = (len(d_571_v50_) if (self).fm11(globalState) else d_556_i2_)
        d_572_v51_: _dafny.Array
        nw119_ = _dafny.Array(None, 13)
        nw119_[int(0)] = p0
        nw119_[int(1)] = 962
        nw119_[int(2)] = len((self).f3)
        nw119_[int(3)] = len((self).f3)
        nw119_[int(4)] = p0
        nw119_[int(5)] = p0
        nw119_[int(6)] = p0
        nw119_[int(7)] = p0
        nw119_[int(8)] = p0
        nw119_[int(9)] = p0
        nw119_[int(10)] = p0
        nw119_[int(11)] = p0
        nw119_[int(12)] = p0
        d_572_v51_ = nw119_
        pat_let_tv12_ = p0
        d_573_v52_: _dafny.Seq
        def iife47_(_pat_let14_0):
            def iife48_(d_574_dt__update__tmp_h1_):
                def iife49_(_pat_let15_0):
                    def iife50_(d_575_dt__update_hcf9_h0_):
                        return D4_DC6(d_575_dt__update_hcf9_h0_, (d_574_dt__update__tmp_h1_).cf10, (d_574_dt__update__tmp_h1_).cf11)
                    return iife50_(_pat_let15_0)
                return iife49_(pat_let_tv12_)
            return iife48_(_pat_let14_0)
        d_573_v52_ = _dafny.SeqWithoutIsStrInference([(iife47_(D4_DC6((d_506_v0_).cardinality, d_572_v51_, p0))).cf9, (0) - (p0), p0])
        d_576_v53_: C0
        nw120_ = C0()
        nw120_.ctor__(p2, p2)
        d_576_v53_ = nw120_
        d_577_v54_: _dafny.Set
        d_577_v54_ = _dafny.Set({d_576_v53_, d_576_v53_})
        d_578_v55_: _dafny.Map
        d_578_v55_ = _dafny.Map({p0: _dafny.Set({d_576_v53_})})
        d_579_v56_: _dafny.Set
        d_579_v56_ = _dafny.Set({p3})
        rhs78_ = (((0) - ((0) - (p0))) - (p0)) <= (len(d_573_v52_))
        rhs79_ = p0
        rhs80_ = (((d_578_v55_)[p0] if (p0) in (d_578_v55_) else d_577_v54_)).ispropersubset(d_577_v54_)
        rhs81_ = len(((d_579_v56_) | (d_579_v56_)) | (d_579_v56_))
        r0 = rhs78_
        r2 = rhs79_
        r0 = rhs80_
        r2 = rhs81_
        r0 = d_576_v53_.f5
        rhs82_ = p0
        rhs83_ = d_576_v53_.f4
        rhs84_ = ((p0) * (p0)) != ((self).fm1(p3, globalState))
        r2 = rhs82_
        r0 = rhs83_
        r0 = rhs84_
        d_580_v57_: _dafny.Map
        d_580_v57_ = _dafny.Map({p0: p3})
        d_581_v58_: _dafny.Map
        d_581_v58_ = _dafny.Map({(self).f3: d_580_v57_})
        source12_ = D5_DC10(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tt"))), d_581_v58_, True)
        if source12_.is_DC9:
            d_582___mcc_h5_ = source12_.cf14
            d_583___mcc_h6_ = source12_.cf15
            d_584___mcc_h7_ = source12_.cf16
            d_585_cf16_ = d_584___mcc_h7_
            d_586_cf15_ = d_583___mcc_h6_
            d_587_cf14_ = d_582___mcc_h5_
            d_588_v59_: D4
            d_588_v59_ = D4_DC6(len(_dafny.Set({p2})), d_572_v51_, len(_dafny.SeqWithoutIsStrInference([(d_573_v52_)[default__.safeIndex(d_586_cf15_, len(d_573_v52_))]])))
            d_589_v60_: _dafny.Map
            d_589_v60_ = _dafny.Map({p0: d_588_v59_})
            d_590_v61_: _dafny.Map
            d_590_v61_ = _dafny.Map({d_589_v60_: p0})
            d_590_v61_ = (d_590_v61_) | (d_590_v61_)
            d_591_v62_: _dafny.Seq
            d_591_v62_ = _dafny.SeqWithoutIsStrInference([not(p3), d_576_v53_.f4, (p2 if p3 else d_576_v53_.f5)])
            rhs85_ = (d_591_v62_) + (d_591_v62_)
            rhs86_ = len((self).f3)
            d_591_v62_ = rhs85_
            d_586_cf15_ = rhs86_
            d_592_v63_: _dafny.Seq
            d_592_v63_ = _dafny.SeqWithoutIsStrInference([d_506_v0_])
            d_506_v0_ = (d_506_v0_) | ((_dafny.MultiSet([d_576_v53_.f4, False])) - ((d_592_v63_)[default__.safeIndex(d_586_cf15_, len(d_592_v63_))]))
            d_593_v64_: _dafny.Array
            nw121_ = _dafny.Array(False, 23)
            d_593_v64_ = nw121_
            d_594_v65_: _dafny.MultiSet
            d_594_v65_ = _dafny.MultiSet([d_576_v53_, d_576_v53_])
            d_595_v66_: _dafny.Map
            d_595_v66_ = _dafny.Map({d_593_v64_: d_594_v65_})
            (d_576_v53_).f4 = (d_591_v62_)[default__.safeIndex((p0 if d_587_cf14_ else len(d_595_v66_)), len(d_591_v62_))]
        elif source12_.is_DC10:
            d_596___mcc_h8_ = source12_.cf17
            d_597___mcc_h9_ = source12_.cf18
            d_598___mcc_h10_ = source12_.cf19
            d_599___mcc_h11_ = source12_.cf20
            d_600_cf20_ = d_599___mcc_h11_
            d_601_cf19_ = d_598___mcc_h10_
            d_602_cf18_ = d_597___mcc_h9_
            d_603_cf17_ = d_596___mcc_h8_
            d_604_v67_: _dafny.Seq
            d_604_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwqob"))
            index100_ = default__.safeIndex(76, (d_572_v51_).length(0))
            (d_572_v51_)[index100_] = (d_573_v52_)[default__.safeIndex(len(d_579_v56_), len(d_573_v52_))]
            index101_ = default__.safeIndex(76, (d_572_v51_).length(0))
            rhs87_ = (d_604_v67_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_605_i3_ in range(default__.abs(234))]))
            rhs88_ = (827) + (d_602_cf18_)
            rhs89_ = d_602_cf18_
            lhs37_ = d_572_v51_
            lhs38_ = default__.safeIndex(76, (d_572_v51_).length(0))
            d_604_v67_ = rhs87_
            lhs37_[lhs38_] = rhs88_
            d_602_cf18_ = rhs89_
            d_600_cf20_ = d_603_cf17_
            index102_ = default__.safeIndex(76, (d_572_v51_).length(0))
            (d_572_v51_)[index102_] = (d_572_v51_)[default__.safeIndex(76, (d_572_v51_).length(0))]
            (d_576_v53_).f5 = d_576_v53_.f5
        elif True:
            d_606___mcc_h12_ = source12_.cf13
            d_607_cf13_ = d_606___mcc_h12_
            d_608_v68_: _dafny.Set
            d_608_v68_ = _dafny.Set({p0, -983})
            d_506_v0_ = (d_506_v0_).set((d_608_v68_).ispropersubset(default__.fm18(d_576_v53_.f4, p3, d_506_v0_, globalState)), default__.abs(p0))
            d_609_v69_: _dafny.Map
            d_609_v69_ = _dafny.Map({p3: d_576_v53_.f5})
            d_610_v70_: _dafny.MultiSet
            d_610_v70_ = _dafny.MultiSet([d_609_v69_, d_609_v69_])
            r2 = (((_dafny.MultiSet([(0) - ((d_610_v70_).cardinality)])).cardinality) - ((d_573_v52_)[default__.safeIndex(p0, len(d_573_v52_))])) + (p0)
            d_611_v71_: C0
            nw122_ = C0()
            nw122_.ctor__(d_576_v53_.f5, not(d_576_v53_.f5))
            d_611_v71_ = nw122_
            d_506_v0_ = _dafny.MultiSet([not (False) or (d_611_v71_.f4)])
        r0 = p3
        d_612_v72_: _dafny.Set
        d_612_v72_ = _dafny.Set({d_555_v37_, d_555_v37_})
        d_613_v73_: _dafny.Map
        d_613_v73_ = _dafny.Map({p3: p2})
        d_614_v74_: _dafny.Array
        nw123_ = _dafny.Array(None, 5)
        nw123_[int(0)] = d_613_v73_
        nw123_[int(1)] = _dafny.Map({d_576_v53_.f5: d_576_v53_.f4})
        nw123_[int(2)] = d_613_v73_
        nw123_[int(3)] = d_613_v73_
        nw123_[int(4)] = d_613_v73_
        d_614_v74_ = nw123_
        d_615_v75_: _dafny.Map
        d_615_v75_ = _dafny.Map({d_612_v72_: d_614_v74_})
        r1 = ((d_615_v75_)[_dafny.Set({d_555_v37_, _dafny.MultiSet([p0]), default__.fm19(p0, p0, d_576_v53_.f4, p0, globalState)})] if (_dafny.Set({d_555_v37_, _dafny.MultiSet([p0]), default__.fm19(p0, p0, d_576_v53_.f4, p0, globalState)})) in (d_615_v75_) else d_614_v74_)
        r2 = (0) - (((self).fm1(p2, globalState)) - ((d_573_v52_)[default__.safeIndex(p0, len(d_573_v52_))]))
        return r0, r1, r2

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_616_v0_: _dafny.Array
        nw124_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
        d_616_v0_ = nw124_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_616_v0_).length(0)):
            d_617_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_617_i0_)) and ((d_617_i0_) < ((d_616_v0_).length(0)))):
                (d_616_v0_)[(d_617_i0_)] = (self).f3
        d_618_i1_: int
        d_618_i1_ = 0
        with _dafny.label("0"):
            while p2:
                with _dafny.c_label("0"):
                    if (d_618_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_618_i1_ = (d_618_i1_) + (1)
                    d_619_v1_: bool
                    d_619_v1_ = True
                    d_619_v1_ = (self).fm11(globalState)
                    d_620_v2_: _dafny.Array
                    nw125_ = _dafny.Array(_dafny.Map({}), 18)
                    d_620_v2_ = nw125_
                    d_621_v3_: _dafny.Array
                    nw126_ = _dafny.Array(D0.default()(), 5)
                    d_621_v3_ = nw126_
                    d_622_v4_: _dafny.Map
                    d_622_v4_ = _dafny.Map({d_621_v3_: d_619_v1_})
                    index103_ = default__.safeIndex(860, (d_620_v2_).length(0))
                    (d_620_v2_)[index103_] = (d_622_v4_) | (d_622_v4_)
                    index104_ = default__.safeIndex(860, (d_620_v2_).length(0))
                    (d_620_v2_)[index104_] = d_622_v4_
                    d_623_v5_: str
                    d_623_v5_ = _dafny.CodePoint('h')
                    d_623_v5_ = ((self).f3)[default__.safeIndex(p0, len((self).f3))]
                    d_624_v6_: _dafny.Set
                    d_624_v6_ = _dafny.Set({p2, d_619_v1_})
                    d_625_v7_: _dafny.Map
                    d_625_v7_ = _dafny.Map({(d_624_v6_) | (d_624_v6_): p1})
                    d_625_v7_ = (d_625_v7_) | ((d_625_v7_) | (d_625_v7_))
                    pass
            pass
        d_626_v8_: _dafny.Seq
        d_626_v8_ = _dafny.SeqWithoutIsStrInference([True, p2])
        d_627_v9_: _dafny.Map
        d_627_v9_ = _dafny.Map({not(p2): (d_626_v8_)[default__.safeIndex(p1, len(d_626_v8_))]})
        d_628_v10_: _dafny.Seq
        d_628_v10_ = _dafny.SeqWithoutIsStrInference([d_627_v9_, d_627_v9_, d_627_v9_, d_627_v9_, d_627_v9_])
        d_628_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p2: p2})])
        d_629_v11_: bool
        d_629_v11_ = False
        d_629_v11_ = d_629_v11_
        d_629_v11_ = d_629_v11_
        d_630_v12_: _dafny.Set
        d_630_v12_ = _dafny.Set({p0})
        d_630_v12_ = d_630_v12_
        r0 = (self).fm1(p2, globalState)
        return r0

    @property
    def f3(self):
        return self._f3

class C2(T0, T2, T1):
    def  __init__(self):
        self._f2: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f2(self):
        return self._f2
    @f2.setter
    def f2(self, value):
        self._f2 = value
    def ctor__(self, f2):
        (self).f2 = f2

    def fm0(self, globalState):
        def iife51_():
            coll19_ = _dafny.Map()
            compr_19_: _dafny.Seq
            for compr_19_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([self.f2]): True})).keys.Elements:
                d_632_v0_: _dafny.Seq = compr_19_
                if (d_632_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([self.f2]): True})):
                    coll19_[d_632_v0_] = self.f2
            return _dafny.Map(coll19_)
        def iife52_():
            coll20_ = _dafny.Map()
            compr_20_: _dafny.Seq
            for compr_20_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))), len(_dafny.Set({True, True, True, not(False), False}))])])).Elements:
                d_633_v1_: _dafny.Seq = compr_20_
                if (d_633_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))), len(_dafny.Set({True, True, True, not(False), False}))])])):
                    def iife53_():
                        coll21_ = _dafny.Map()
                        compr_21_: str
                        for compr_21_ in (_dafny.MultiSet([_dafny.CodePoint('m')])).Elements:
                            d_634_v2_: str = compr_21_
                            if (d_634_v2_) in (_dafny.MultiSet([_dafny.CodePoint('m')])):
                                coll21_[d_634_v2_] = self.f2
                        return _dafny.Map(coll21_)
                    coll20_[d_633_v1_] = len(iife53_()
                    )
            return _dafny.Map(coll20_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([self.f2]): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([self.f2, self.f2])) for d_631_i0_ in range(default__.abs(491))]))).cardinality})) | (iife51_()
        )) | ((iife52_()
        ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([self.f2]))]): len(_dafny.Map({self.f2: self.f2}))})))

    def fm1(self, p0, globalState):
        return self.f2

    def fm8(self, p0, p1, globalState):
        def iife54_():
            coll22_ = _dafny.Map()
            compr_22_: int
            for compr_22_ in _dafny.IntegerRange(934, 570):
                d_636_v0_: int = compr_22_
                if ((934) <= (d_636_v0_)) and ((d_636_v0_) < (570)):
                    coll22_[default__.safeModuloInt(d_636_v0_, self.f2)] = False
            return _dafny.Map(coll22_)
        return ((_dafny.Set({len(_dafny.Set({True})), len(_dafny.Map({(D5_DC10(True, self.f2, _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_635_i0_ in range(default__.abs(436))]): iife54_()
}), True)).cf18: not(True)})), self.f2})) | (_dafny.Set({self.f2, self.f2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cawq"))), self.f2, self.f2}))) == ((_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydjxbae")))})) - (_dafny.Set({(0) - (len(_dafny.Map({_dafny.CodePoint('w'): self.f2}))), len(_dafny.Map({False: 911}))})))

    def fm24(self, p0, p1, p2, globalState):
        return _dafny.Map({(D7_DC14(_dafny.SeqWithoutIsStrInference([self.f2, len(_dafny.SeqWithoutIsStrInference([-59]))]))).cf26: self.f2})

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        d_637_v0_: _dafny.Set
        d_637_v0_ = _dafny.Set({self.f2})
        d_638_v1_: _dafny.Map
        d_638_v1_ = _dafny.Map({p3: d_637_v0_})
        d_639_v2_: _dafny.Map
        d_639_v2_ = _dafny.Map({p3: p0})
        d_640_v3_: _dafny.Seq
        d_640_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rftn"))
        d_641_v4_: _dafny.Array
        nw127_ = _dafny.Array(int(0), 17)
        d_641_v4_ = nw127_
        d_642_v5_: _dafny.Array
        nw128_ = _dafny.Array(None, 18)
        nw128_[int(0)] = p3
        nw128_[int(1)] = p3
        nw128_[int(2)] = p3
        nw128_[int(3)] = p2
        nw128_[int(4)] = p3
        nw128_[int(5)] = False
        nw128_[int(6)] = p3
        nw128_[int(7)] = p2
        nw128_[int(8)] = p3
        nw128_[int(9)] = p2
        nw128_[int(10)] = p3
        nw128_[int(11)] = p3
        nw128_[int(12)] = True
        nw128_[int(13)] = p3
        nw128_[int(14)] = p2
        nw128_[int(15)] = p3
        nw128_[int(16)] = p2
        nw128_[int(17)] = p2
        d_642_v5_ = nw128_
        d_643_v6_: _dafny.Map
        d_643_v6_ = _dafny.Map({d_642_v5_: p0})
        d_644_v7_: D4
        d_644_v7_ = D4_DC6((self).fm1(p3, globalState), d_641_v4_, len(d_643_v6_))
        d_645_v8_: D4
        d_645_v8_ = D4_DC7(d_644_v7_)
        d_646_v9_: _dafny.Seq
        d_646_v9_ = _dafny.SeqWithoutIsStrInference([d_645_v8_])
        d_647_v10_: _dafny.Seq
        d_647_v10_ = _dafny.SeqWithoutIsStrInference([p3, p2, p3, p3, p2])
        d_648_v11_: str
        d_648_v11_ = _dafny.CodePoint('e')
        d_649_v12_: _dafny.Array
        nw129_ = _dafny.Array(None, 16)
        nw129_[int(0)] = default__.safeDivisionInt(len(((d_638_v1_)[p3] if (p3) in (d_638_v1_) else d_637_v0_)), len(d_639_v2_))
        nw129_[int(1)] = len(d_640_v3_)
        nw129_[int(2)] = p0
        nw129_[int(3)] = len(d_646_v9_)
        nw129_[int(4)] = default__.safeDivisionInt(self.f2, self.f2)
        nw129_[int(5)] = len(d_647_v10_)
        nw129_[int(6)] = p0
        nw129_[int(7)] = (_dafny.MultiSet([d_648_v11_])).cardinality
        nw129_[int(8)] = (self).fm1(False, globalState)
        nw129_[int(9)] = self.f2
        nw129_[int(10)] = p0
        nw129_[int(11)] = self.f2
        nw129_[int(12)] = (0) - (self.f2)
        nw129_[int(13)] = p0
        nw129_[int(14)] = (378) + (236)
        nw129_[int(15)] = (_dafny.MultiSet(d_647_v10_)).cardinality
        d_649_v12_ = nw129_
        index105_ = default__.safeIndex(239, (d_641_v4_).length(0))
        (d_641_v4_)[index105_] = len(d_640_v3_)
        index106_ = default__.safeIndex(239, (d_641_v4_).length(0))
        rhs90_ = ((0) - (default__.safeDivisionInt(len(d_640_v3_), (0) - ((self).fm1(p2, globalState))))) - (self.f2)
        rhs91_ = (0) - (self.f2)
        lhs39_ = d_641_v4_
        lhs40_ = default__.safeIndex(239, (d_641_v4_).length(0))
        r2 = rhs90_
        lhs39_[lhs40_] = rhs91_
        if (self.f2) <= ((d_641_v4_)[default__.safeIndex(239, (d_641_v4_).length(0))]):
            d_650_v13_: _dafny.Map
            d_650_v13_ = (self).fm0(globalState)
            source13_ = d_650_v13_
            d_651___mcc_h0_ = source13_
            d_652_cf6_ = d_651___mcc_h0_
            index107_ = default__.safeIndex(239, (d_641_v4_).length(0))
            (d_641_v4_)[index107_] = (0) - ((self).fm1(p3, globalState))
            index108_ = default__.safeIndex(239, (d_641_v4_).length(0))
            (d_641_v4_)[index108_] = p0
            r0 = p2
            index109_ = default__.safeIndex(239, (d_641_v4_).length(0))
            (d_641_v4_)[index109_] = default__.safeModuloInt(p0, (d_641_v4_)[default__.safeIndex(239, (d_641_v4_).length(0))])
            d_653_v14_: _dafny.Set
            d_653_v14_ = _dafny.Set({p2})
            d_653_v14_ = ((d_653_v14_).intersection(d_653_v14_) if (p0) == (p0) else d_653_v14_)
            d_654_v15_: _dafny.Seq
            d_654_v15_ = _dafny.SeqWithoutIsStrInference([293, (d_641_v4_)[default__.safeIndex(239, (d_641_v4_).length(0))]])
            d_655_v16_: _dafny.Map
            d_655_v16_ = _dafny.Map({d_654_v15_: len(d_640_v3_)})
            d_656_v17_: _dafny.Seq
            d_656_v17_ = _dafny.SeqWithoutIsStrInference([d_650_v13_, d_655_v16_])
            r0 = (p3) and ((d_656_v17_) != (d_656_v17_))
            d_657_v18_: D4
            d_657_v18_ = D4_DC6(p0, d_649_v12_, self.f2)
            d_658_v19_: _dafny.Map
            d_658_v19_ = _dafny.Map({(d_657_v18_).cf10: p3})
            d_659_v20_: C0
            nw130_ = C0()
            nw130_.ctor__(p3, ((d_658_v19_)[d_649_v12_] if (d_649_v12_) in (d_658_v19_) else p3))
            d_659_v20_ = nw130_
            (self).f2 = (default__.safeModuloInt(p0, p0)) + ((self.f2) * (self.f2))
        elif True:
            d_660_v21_: D7
            d_660_v21_ = D7_DC15((d_641_v4_)[default__.safeIndex(239, (d_641_v4_).length(0))], d_648_v11_)
            d_661_v22_: _dafny.Map
            d_661_v22_ = _dafny.Map({d_660_v21_: p2})
            d_662_v23_: _dafny.MultiSet
            d_662_v23_ = _dafny.MultiSet([d_648_v11_])
            d_663_v24_: _dafny.MultiSet
            d_663_v24_ = _dafny.MultiSet([len(d_639_v2_)])
            d_664_v25_: _dafny.Seq
            d_664_v25_ = _dafny.SeqWithoutIsStrInference([self.f2, p0])
            index110_ = default__.safeIndex(239, (d_641_v4_).length(0))
            rhs92_ = (len(d_661_v22_)) < (default__.safeModuloInt(p0, self.f2))
            rhs93_ = default__.safeDivisionInt((self).fm1(p3, globalState), p0)
            rhs94_ = ((0) - (default__.safeModuloInt(p0, len(d_647_v10_)))) * ((d_662_v23_).cardinality)
            rhs95_ = default__.safeModuloInt((d_641_v4_)[default__.safeIndex(239, (d_641_v4_).length(0))], self.f2)
            rhs96_ = ((d_663_v24_) | ((_dafny.MultiSet(d_664_v25_)) | (d_663_v24_))).cardinality
            lhs41_ = self
            lhs42_ = self
            lhs43_ = d_641_v4_
            lhs44_ = default__.safeIndex(239, (d_641_v4_).length(0))
            r0 = rhs92_
            lhs41_.f2 = rhs93_
            r2 = rhs94_
            lhs42_.f2 = rhs95_
            lhs43_[lhs44_] = rhs96_
            index111_ = default__.safeIndex(239, (d_641_v4_).length(0))
            (d_641_v4_)[index111_] = 814
            d_647_v10_ = d_647_v10_
            d_665_v26_: _dafny.MultiSet
            d_665_v26_ = _dafny.MultiSet([p2])
            index112_ = default__.safeIndex(239, (d_641_v4_).length(0))
            (d_641_v4_)[index112_] = default__.safeDivisionInt(p0, ((d_665_v26_ if p3 else _dafny.MultiSet([p3]))).cardinality)
            d_666_v27_: _dafny.Map
            d_666_v27_ = _dafny.Map({d_640_v3_: _dafny.Map({self.f2: False})})
            index113_ = default__.safeIndex(185, (d_642_v5_).length(0))
            (d_642_v5_)[index113_] = (D5_DC10(p2, len(d_637_v0_), d_666_v27_, p2)).cf20
            index114_ = default__.safeIndex(185, (d_642_v5_).length(0))
            (d_642_v5_)[index114_] = p2
        d_642_v5_ = d_642_v5_
        d_667_v28_: _dafny.Seq
        d_667_v28_ = _dafny.SeqWithoutIsStrInference([99])
        pat_let_tv13_ = d_648_v11_
        pat_let_tv14_ = p3
        pat_let_tv15_ = p0
        pat_let_tv16_ = d_641_v4_
        pat_let_tv17_ = d_641_v4_
        def lambda20_(source14_):
            if source14_.is_DC1:
                d_668___mcc_h1_ = source14_.cf1
                d_669___mcc_h2_ = source14_.cf2
                d_670___mcc_h3_ = source14_.cf3
                d_671___mcc_h4_ = source14_.cf4
                d_672_cf4_ = d_671___mcc_h4_
                d_673_cf3_ = d_670___mcc_h3_
                d_674_cf2_ = d_669___mcc_h2_
                d_675_cf1_ = d_668___mcc_h1_
                return (0) - (default__.safeModuloInt(self.f2, len(_dafny.Map({d_672_cf4_: pat_let_tv13_}))))
            elif True:
                d_676___mcc_h5_ = source14_.cf0
                d_677_cf0_ = d_676___mcc_h5_
                return default__.safeDivisionInt(len(_dafny.Map({pat_let_tv14_: pat_let_tv15_})), (pat_let_tv17_)[default__.safeIndex(239, (pat_let_tv16_).length(0))])

        r2 = lambda20_(default__.fm25(len(d_667_v28_), (self).fm8(d_640_v3_, p2, globalState), _dafny.MultiSet([245, self.f2]), _dafny.SeqWithoutIsStrInference([p2]), globalState))
        d_678_v29_: D7
        d_678_v29_ = D7_DC14(d_667_v28_)
        d_679_v30_: _dafny.Map
        d_679_v30_ = _dafny.Map({self.f2: p0})
        d_680_v31_: _dafny.Array
        nw131_ = _dafny.Array(None, 27)
        nw131_[int(0)] = (d_667_v28_) + (d_667_v28_)
        nw131_[int(1)] = (_dafny.SeqWithoutIsStrInference([self.f2])) + (d_667_v28_)
        nw131_[int(2)] = d_667_v28_
        nw131_[int(3)] = d_667_v28_
        nw131_[int(4)] = d_667_v28_
        nw131_[int(5)] = _dafny.SeqWithoutIsStrInference([self.f2])
        nw131_[int(6)] = _dafny.SeqWithoutIsStrInference([p0 for d_681_i1_ in range(default__.abs(543))])
        nw131_[int(7)] = _dafny.SeqWithoutIsStrInference([self.f2, (self).fm1(p2, globalState)])
        nw131_[int(8)] = (d_667_v28_) + (_dafny.SeqWithoutIsStrInference([p0]))
        nw131_[int(9)] = (d_667_v28_) + (d_667_v28_)
        nw131_[int(10)] = d_667_v28_
        nw131_[int(11)] = _dafny.SeqWithoutIsStrInference([p0 for d_682_i2_ in range(default__.abs(443))])
        nw131_[int(12)] = (d_667_v28_) + (d_667_v28_)
        nw131_[int(13)] = (_dafny.SeqWithoutIsStrInference([self.f2, len(d_637_v0_), (d_641_v4_)[default__.safeIndex(239, (d_641_v4_).length(0))]])) + (d_667_v28_)
        nw131_[int(14)] = d_667_v28_
        nw131_[int(15)] = (_dafny.SeqWithoutIsStrInference([self.f2])) + (_dafny.SeqWithoutIsStrInference([self.f2, (d_667_v28_)[default__.safeIndex(self.f2, len(d_667_v28_))]]))
        nw131_[int(16)] = d_667_v28_
        nw131_[int(17)] = d_667_v28_
        nw131_[int(18)] = _dafny.SeqWithoutIsStrInference([len(d_640_v3_) for d_683_i3_ in range(default__.abs(24))])
        nw131_[int(19)] = d_667_v28_
        nw131_[int(20)] = (d_678_v29_).cf26
        nw131_[int(21)] = (_dafny.SeqWithoutIsStrInference([(d_641_v4_)[default__.safeIndex(239, (d_641_v4_).length(0))], len(d_679_v30_)])) + ((d_667_v28_).set(default__.safeIndex(p0, len(d_667_v28_)), self.f2))
        nw131_[int(22)] = d_667_v28_
        nw131_[int(23)] = d_667_v28_
        nw131_[int(24)] = _dafny.SeqWithoutIsStrInference([334 for d_684_i4_ in range(default__.abs(700))])
        nw131_[int(25)] = _dafny.SeqWithoutIsStrInference([self.f2])
        nw131_[int(26)] = d_667_v28_
        d_680_v31_ = nw131_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_680_v31_).length(0)):
            d_685_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_685_i0_)) and ((d_685_i0_) < ((d_680_v31_).length(0)))):
                (d_680_v31_)[(d_685_i0_)] = d_667_v28_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_641_v4_).length(0)):
            d_686_i5_: int = guard_loop_3_
            if (True) and (((0) <= (d_686_i5_)) and ((d_686_i5_) < ((d_641_v4_).length(0)))):
                (d_641_v4_)[(d_686_i5_)] = (d_686_i5_) + (p0)
        r0 = (not (p2) or (False)) or (p3)
        d_687_v32_: _dafny.Array
        def lambda21_(d_688_p3_):
            def lambda22_(d_689_i6_):
                return (_dafny.Map({d_688_p3_: d_688_p3_})) | (_dafny.Map({False: d_688_p3_}))

            return lambda22_

        init10_ = lambda21_(p3)
        nw132_ = _dafny.Array(None, 1)
        for i0_10_ in range(nw132_.length(0)):
            nw132_[i0_10_] = init10_(i0_10_)
        d_687_v32_ = nw132_
        r1 = d_687_v32_
        r2 = default__.safeDivisionInt(p0, (self.f2) - (p0))
        return r0, r1, r2

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        if p2:
            d_690_v0_: _dafny.Array
            nw133_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_690_v0_ = nw133_
            d_691_v1_: _dafny.Array
            d_691_v1_ = d_690_v0_
            d_691_v1_ = d_690_v0_
            d_692_v2_: _dafny.Seq
            d_692_v2_ = _dafny.SeqWithoutIsStrInference([p2])
            d_693_v3_: _dafny.Seq
            d_693_v3_ = _dafny.SeqWithoutIsStrInference([p2, (d_692_v2_)[default__.safeIndex(p1, len(d_692_v2_))]])
            d_694_v4_: _dafny.MultiSet
            d_694_v4_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvtp"))), p1, (self.f2) + (514), len((_dafny.SeqWithoutIsStrInference([p2])) + (d_693_v3_))])
            d_695_v5_: _dafny.Map
            d_695_v5_ = _dafny.Map({p2: p2})
            (self).f2 = ((d_694_v4_)[self.f2] if (self.f2) in (d_694_v4_) else len((d_695_v5_) | (d_695_v5_)))
            d_696_v6_: _dafny.Seq
            d_696_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
            d_697_v7_: _dafny.Seq
            d_697_v7_ = _dafny.SeqWithoutIsStrInference([d_696_v6_, d_696_v6_])
            d_697_v7_ = (d_697_v7_) + (_dafny.SeqWithoutIsStrInference([d_696_v6_, d_696_v6_]))
            d_698_v8_: C1
            nw134_ = C1()
            nw134_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bawn")))
            d_698_v8_ = nw134_
            def iife55_():
                coll23_ = _dafny.Set()
                compr_23_: int
                for compr_23_ in _dafny.IntegerRange(-299, 388):
                    d_699_v9_: int = compr_23_
                    if ((-299) <= (d_699_v9_)) and ((d_699_v9_) < (388)):
                        coll23_ = coll23_.union(_dafny.Set([(d_699_v9_) + (self.f2)]))
                return _dafny.Set(coll23_)
            if (len(iife55_()
            )) <= (default__.safeModuloInt((0) - (len(d_693_v3_)), self.f2)):
                d_700_v10_: C1
                nw135_ = C1()
                nw135_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlxtt")))
                d_700_v10_ = nw135_
                d_701_v11_: _dafny.Map
                d_701_v11_ = _dafny.Map({d_696_v6_: _dafny.Map({p1: p2})})
                d_702_v12_: _dafny.Array
                nw136_ = _dafny.Array(None, 1)
                nw136_[int(0)] = self.f2
                d_702_v12_ = nw136_
                d_703_v13_: _dafny.Set
                d_703_v13_ = _dafny.Set({d_702_v12_})
                d_704_v14_: C0
                nw137_ = C0()
                nw137_.ctor__(not (p2) or ((D5_DC10(p2, len(_dafny.SeqWithoutIsStrInference([(d_696_v6_)[default__.safeIndex(self.f2, len(d_696_v6_))] for d_705_i0_ in range(default__.abs(48))])), d_701_v11_, p2)).cf17), (d_703_v13_) == (d_703_v13_))
                d_704_v14_ = nw137_
                d_706_v15_: D0
                d_706_v15_ = D0_DC0(not(not(p2)))
                d_707_v16_: str
                d_707_v16_ = _dafny.CodePoint('c')
                d_708_v17_: _dafny.Array
                nw138_ = _dafny.Array(None, 5)
                nw138_[int(0)] = d_706_v15_
                nw138_[int(1)] = default__.fm25(len(default__.fm26(len(_dafny.Map({d_707_v16_: _dafny.Map({d_704_v14_.f5: p2})})), d_704_v14_.f4, d_704_v14_.f5, d_704_v14_.f5, globalState)), d_704_v14_.f4, d_694_v4_, d_693_v3_, globalState)
                nw138_[int(2)] = D0_DC0(d_704_v14_.f5)
                nw138_[int(3)] = d_706_v15_
                nw138_[int(4)] = d_706_v15_
                d_708_v17_ = nw138_
                index115_ = default__.safeIndex(410, (d_708_v17_).length(0))
                (d_708_v17_)[index115_] = d_706_v15_
                index116_ = default__.safeIndex(410, (d_708_v17_).length(0))
                (d_708_v17_)[index116_] = default__.fm25(p0, not(d_704_v14_.f4), d_694_v4_, (d_693_v3_).set(default__.safeIndex((0) - (p0), len(d_693_v3_)), not(True)), globalState)
                index117_ = default__.safeIndex(402, (d_702_v12_).length(0))
                (d_702_v12_)[index117_] = (self.f2) + (p0)
                index118_ = default__.safeIndex(402, (d_702_v12_).length(0))
                (d_702_v12_)[index118_] = (-544) + (p0)
                d_709_v18_: _dafny.Set
                d_709_v18_ = _dafny.Set({len(d_696_v6_), len(d_693_v3_)})
                d_710_v19_: _dafny.Map
                d_710_v19_ = _dafny.Map({p1: ((0) - (len(d_709_v18_))) - ((d_702_v12_)[default__.safeIndex(402, (d_702_v12_).length(0))])})
                d_711_v20_: _dafny.Seq
                d_711_v20_ = _dafny.SeqWithoutIsStrInference([(default__.fm27(len(_dafny.Map({self.f2: not(False)})), p2, globalState)) | (_dafny.Map({len(d_710_v19_): p1}))])
                index119_ = default__.safeIndex(402, (d_702_v12_).length(0))
                rhs97_ = (d_711_v20_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p2])), len(d_711_v20_))]
                rhs98_ = not(d_704_v14_.f5)
                rhs99_ = (d_702_v12_)[default__.safeIndex(402, (d_702_v12_).length(0))]
                rhs100_ = default__.safeDivisionInt(self.f2, (len((d_700_v10_).f3)) * (self.f2))
                rhs101_ = d_710_v19_
                lhs45_ = d_704_v14_
                lhs46_ = d_702_v12_
                lhs47_ = default__.safeIndex(402, (d_702_v12_).length(0))
                d_710_v19_ = rhs97_
                lhs45_.f4 = rhs98_
                r0 = rhs99_
                lhs46_[lhs47_] = rhs100_
                d_710_v19_ = rhs101_
            elif True:
                d_712_v21_: C0
                nw139_ = C0()
                nw139_.ctor__(not((p1) >= (p0)), not(p2))
                d_712_v21_ = nw139_
                d_713_v22_: str
                d_713_v22_ = _dafny.CodePoint('i')
                d_714_v23_: _dafny.Map
                d_714_v23_ = _dafny.Map({len(d_695_v5_): d_713_v22_})
                (d_712_v21_).f4 = not((((d_714_v23_)[(0) - (len((d_698_v8_).f3))] if ((0) - (len((d_698_v8_).f3))) in (d_714_v23_) else d_713_v22_)) not in ((d_698_v8_).f3))
                d_715_v24_: _dafny.Array
                nw140_ = _dafny.Array(D7.default()(), 28)
                d_715_v24_ = nw140_
                d_716_v25_: _dafny.Map
                d_716_v25_ = _dafny.Map({d_713_v22_: d_715_v24_})
                d_717_v26_: _dafny.MultiSet
                d_717_v26_ = _dafny.MultiSet([d_715_v24_, ((d_716_v25_)[_dafny.CodePoint('s')] if (_dafny.CodePoint('s')) in (d_716_v25_) else d_715_v24_)])
                d_718_v27_: _dafny.Map
                d_718_v27_ = _dafny.Map({d_717_v26_: self.f2})
                d_719_v28_: _dafny.Map
                d_719_v28_ = _dafny.Map({p1: d_712_v21_.f5})
                d_718_v27_ = (d_718_v27_).set((d_717_v26_).set(d_715_v24_, default__.abs(p1)), len(d_719_v28_))
                d_720_v29_: D0
                d_720_v29_ = D0_DC0(p2)
                d_721_v30_: _dafny.Array
                nw141_ = _dafny.Array(None, 19)
                nw141_[int(0)] = d_712_v21_.f4
                nw141_[int(1)] = (self).fm8((d_698_v8_).f3, False, globalState)
                nw141_[int(2)] = p2
                nw141_[int(3)] = d_712_v21_.f4
                nw141_[int(4)] = d_712_v21_.f5
                nw141_[int(5)] = d_712_v21_.f4
                nw141_[int(6)] = False
                nw141_[int(7)] = d_712_v21_.f5
                nw141_[int(8)] = not(not(False))
                nw141_[int(9)] = (d_712_v21_.f5) or (p2)
                nw141_[int(10)] = (p2) and (d_712_v21_.f4)
                nw141_[int(11)] = (d_720_v29_).cf0
                nw141_[int(12)] = d_712_v21_.f4
                nw141_[int(13)] = True
                nw141_[int(14)] = p2
                nw141_[int(15)] = True
                nw141_[int(16)] = (((d_694_v4_)[p0] if (p0) in (d_694_v4_) else p1)) in (d_694_v4_)
                nw141_[int(17)] = True
                nw141_[int(18)] = d_712_v21_.f4
                d_721_v30_ = nw141_
                index120_ = default__.safeIndex(771, (d_721_v30_).length(0))
                (d_721_v30_)[index120_] = (len((d_698_v8_).f3)) > (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(_dafny.MultiSet([p0])).cardinality: self.f2})) for d_722_i1_ in range(default__.abs(208))])))
                index121_ = default__.safeIndex(771, (d_721_v30_).length(0))
                (d_721_v30_)[index121_] = ((d_698_v8_).f3) < ((d_698_v8_).f3)
                d_723_v31_: _dafny.Map
                d_723_v31_ = _dafny.Map({d_721_v30_: self.f2})
                d_724_v32_: _dafny.Array
                nw142_ = _dafny.Array(None, 12)
                nw142_[int(0)] = p0
                nw142_[int(1)] = p0
                nw142_[int(2)] = len(d_719_v28_)
                nw142_[int(3)] = self.f2
                nw142_[int(4)] = 864
                nw142_[int(5)] = (0) - (self.f2)
                nw142_[int(6)] = (d_698_v8_).fm1(p2, globalState)
                nw142_[int(7)] = p1
                nw142_[int(8)] = len(d_723_v31_)
                nw142_[int(9)] = p0
                nw142_[int(10)] = 321
                nw142_[int(11)] = p0
                d_724_v32_ = nw142_
                d_725_v33_: _dafny.MultiSet
                d_725_v33_ = _dafny.MultiSet([d_724_v32_, d_724_v32_])
                d_726_v34_: D5
                d_726_v34_ = D5_DC10(d_712_v21_.f4, self.f2, _dafny.Map({(d_698_v8_).f3: d_719_v28_}), d_712_v21_.f5)
                d_727_v35_: _dafny.Map
                d_727_v35_ = _dafny.Map({(p0) * (87): _dafny.Set({d_726_v34_})})
                d_728_v36_: _dafny.Set
                d_728_v36_ = _dafny.Set({d_712_v21_.f4})
                d_729_v37_: _dafny.MultiSet
                d_729_v37_ = _dafny.MultiSet([not(d_712_v21_.f4)])
                index122_ = default__.safeIndex(771, (d_721_v30_).length(0))
                index123_ = default__.safeIndex(771, (d_721_v30_).length(0))
                rhs102_ = d_712_v21_.f4
                rhs103_ = (d_725_v33_).intersection(d_725_v33_)
                rhs104_ = ((d_696_v6_) + (default__.fm28(p2, d_712_v21_.f4, d_712_v21_.f4, d_728_v36_, globalState))) <= ((d_698_v8_).f3)
                rhs105_ = (d_727_v35_) | (d_727_v35_)
                rhs106_ = _dafny.SeqWithoutIsStrInference([(d_729_v37_).ispropersubset(d_729_v37_)])
                lhs48_ = d_721_v30_
                lhs49_ = default__.safeIndex(771, (d_721_v30_).length(0))
                lhs50_ = d_721_v30_
                lhs51_ = default__.safeIndex(771, (d_721_v30_).length(0))
                lhs48_[lhs49_] = rhs102_
                d_725_v33_ = rhs103_
                lhs50_[lhs51_] = rhs104_
                d_727_v35_ = rhs105_
                d_693_v3_ = rhs106_
        elif True:
            d_730_v38_: bool
            d_730_v38_ = True
            d_731_v39_: _dafny.Seq
            d_731_v39_ = _dafny.SeqWithoutIsStrInference([p2, p2, not(p2)])
            d_732_v41_: _dafny.Seq
            d_732_v41_ = _dafny.SeqWithoutIsStrInference([p1, p0])
            def iife56_():
                coll24_ = _dafny.Map()
                compr_24_: int
                for compr_24_ in (d_732_v41_).Elements:
                    d_734_v40_: int = compr_24_
                    if (d_734_v40_) in (d_732_v41_):
                        coll24_[default__.safeDivisionInt(d_734_v40_, (0) - (self.f2))] = d_730_v38_
                return _dafny.Map(coll24_)
            d_730_v38_ = (self).fm8(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_733_i2_ in range(default__.abs(202))]), (d_731_v39_)[default__.safeIndex(len(iife56_()
            ), len(d_731_v39_))], globalState)
            (self).f2 = default__.safeDivisionInt(p1, p0)
            d_735_v42_: _dafny.Seq
            d_735_v42_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_736_i3_ in range(default__.abs(293))])])
            d_737_v43_: C1
            nw143_ = C1()
            nw143_.ctor__((d_735_v42_)[default__.safeIndex(p1, len(d_735_v42_))])
            d_737_v43_ = nw143_
            source15_ = default__.fm29(globalState)
            if source15_.is_DC9:
                d_738___mcc_h0_ = source15_.cf14
                d_739___mcc_h1_ = source15_.cf15
                d_740___mcc_h2_ = source15_.cf16
                d_741_cf16_ = d_740___mcc_h2_
                d_742_cf15_ = d_739___mcc_h1_
                d_743_cf14_ = d_738___mcc_h0_
                d_744_v44_: _dafny.Array
                def lambda23_(d_745_i4_):
                    return D7_DC15(-344, _dafny.CodePoint('a'))

                init11_ = lambda23_
                nw144_ = _dafny.Array(None, 12)
                for i0_11_ in range(nw144_.length(0)):
                    nw144_[i0_11_] = init11_(i0_11_)
                d_744_v44_ = nw144_
                d_746_v45_: _dafny.Set
                d_746_v45_ = _dafny.Set({d_744_v44_, d_744_v44_})
                d_746_v45_ = d_746_v45_
                (self).f2 = p0
                d_742_cf15_ = len((d_737_v43_).f3)
                d_743_cf14_ = d_743_cf14_
            elif source15_.is_DC10:
                d_747___mcc_h3_ = source15_.cf17
                d_748___mcc_h4_ = source15_.cf18
                d_749___mcc_h5_ = source15_.cf19
                d_750___mcc_h6_ = source15_.cf20
                d_751_cf20_ = d_750___mcc_h6_
                d_752_cf19_ = d_749___mcc_h5_
                d_753_cf18_ = d_748___mcc_h4_
                d_754_cf17_ = d_747___mcc_h3_
                d_754_cf17_ = p2
                d_755_v46_: _dafny.Map
                d_755_v46_ = _dafny.Map({p0: p2})
                d_756_v47_: D5
                d_756_v47_ = D5_DC10(d_754_cf17_, 637, _dafny.Map({(d_737_v43_).f3: d_755_v46_}), d_754_cf17_)
                d_757_v48_: _dafny.Set
                d_757_v48_ = _dafny.Set({(d_756_v47_).cf18})
                d_758_v49_: _dafny.Array
                nw145_ = _dafny.Array(int(0), 21)
                d_758_v49_ = nw145_
                index124_ = default__.safeIndex(160, (d_758_v49_).length(0))
                (d_758_v49_)[index124_] = self.f2
                d_759_v50_: _dafny.Map
                d_759_v50_ = _dafny.Map({d_758_v49_: p2})
                index125_ = default__.safeIndex(160, (d_758_v49_).length(0))
                rhs107_ = (_dafny.Set({p1})).intersection(_dafny.Set({p0}))
                rhs108_ = len(((d_737_v43_).f3) + ((d_737_v43_).f3))
                rhs109_ = d_753_cf18_
                rhs110_ = d_758_v49_
                rhs111_ = ((d_759_v50_)[d_758_v49_] if (d_758_v49_) in (d_759_v50_) else (d_731_v39_)[default__.safeIndex(577, len(d_731_v39_))])
                lhs52_ = self
                lhs53_ = d_758_v49_
                lhs54_ = default__.safeIndex(160, (d_758_v49_).length(0))
                d_757_v48_ = rhs107_
                lhs52_.f2 = rhs108_
                lhs53_[lhs54_] = rhs109_
                d_758_v49_ = rhs110_
                d_730_v38_ = rhs111_
                r0 = self.f2
                d_760_v51_: _dafny.MultiSet
                d_760_v51_ = _dafny.MultiSet([(d_737_v43_).f3, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_761_i5_ in range(default__.abs(48))])])
                d_762_v52_: _dafny.Map
                d_762_v52_ = _dafny.Map({d_760_v51_: (p0) >= (48)})
                d_762_v52_ = (d_762_v52_).set(default__.fm30(d_754_cf17_, globalState), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okwhp")))) > (self.f2))
            elif True:
                d_763___mcc_h7_ = source15_.cf13
                d_764_cf13_ = d_763___mcc_h7_
                d_730_v38_ = d_730_v38_
                d_765_v53_: _dafny.Array
                nw146_ = _dafny.Array(_dafny.Seq({}), 23)
                d_765_v53_ = nw146_
                d_766_v54_: bool
                d_767_v55_: _dafny.Array
                d_768_v56_: int
                out19_: bool
                out20_: _dafny.Array
                out21_: int
                out19_, out20_, out21_ = (d_737_v43_).m1((d_737_v43_).fm1(False, globalState), d_765_v53_, (d_731_v39_)[default__.safeIndex(p0, len(d_731_v39_))], p2, globalState)
                d_766_v54_ = out19_
                d_767_v55_ = out20_
                d_768_v56_ = out21_
                d_769_v57_: D0
                d_769_v57_ = D0_DC0((self.f2) > (p0))
                rhs112_ = -629
                rhs113_ = (d_737_v43_).fm1(d_766_v54_, globalState)
                rhs114_ = d_769_v57_
                rhs115_ = (d_737_v43_).fm11(globalState)
                lhs55_ = self
                r0 = rhs112_
                lhs55_.f2 = rhs113_
                d_769_v57_ = rhs114_
                d_730_v38_ = rhs115_
                d_770_v58_: _dafny.Set
                d_770_v58_ = _dafny.Set({self.f2, d_768_v56_})
                d_771_v59_: _dafny.Map
                d_771_v59_ = _dafny.Map({d_730_v38_: d_768_v56_})
                d_730_v38_ = (_dafny.Set({len((d_737_v43_).f3)})).issubset((d_770_v58_).intersection(_dafny.Set({(0) - (len(d_771_v59_)), self.f2, (0) - (d_768_v56_), d_768_v56_, -225})))
            d_772_v60_: _dafny.Map
            d_772_v60_ = _dafny.Map({531: d_730_v38_})
            d_773_v61_: _dafny.Map
            d_773_v61_ = _dafny.Map({(d_737_v43_).f3: d_772_v60_})
            d_774_v62_: D5
            d_774_v62_ = D5_DC10(d_730_v38_, p1, d_773_v61_, p2)
            d_775_v63_: _dafny.MultiSet
            d_775_v63_ = _dafny.MultiSet([p2, p2, d_730_v38_, d_730_v38_])
            d_776_v64_: _dafny.Set
            d_776_v64_ = _dafny.Set({self.f2})
            d_777_v65_: _dafny.Map
            d_777_v65_ = _dafny.Map({d_730_v38_: p2})
            d_778_v66_: _dafny.Array
            nw147_ = _dafny.Array(None, 25)
            nw147_[int(0)] = p1
            nw147_[int(1)] = self.f2
            nw147_[int(2)] = (len(_dafny.Set({-236})) if p2 else p1)
            nw147_[int(3)] = self.f2
            nw147_[int(4)] = p1
            nw147_[int(5)] = default__.safeDivisionInt((d_774_v62_).cf18, self.f2)
            nw147_[int(6)] = (0) - (self.f2)
            nw147_[int(7)] = p1
            nw147_[int(8)] = (p1 if p2 else (0) - ((0) - (((d_775_v63_)[d_730_v38_] if (d_730_v38_) in (d_775_v63_) else self.f2))))
            nw147_[int(9)] = (len(_dafny.SeqWithoutIsStrInference([self.f2, 619]))) - (p1)
            nw147_[int(10)] = (241) + (-534)
            nw147_[int(11)] = self.f2
            nw147_[int(12)] = ((0) - (p0) if d_730_v38_ else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyp"))))
            nw147_[int(13)] = 276
            nw147_[int(14)] = p0
            nw147_[int(15)] = default__.safeModuloInt(len(d_732_v41_), len(d_731_v39_))
            nw147_[int(16)] = (308 if d_730_v38_ else p1)
            nw147_[int(17)] = self.f2
            nw147_[int(18)] = len(d_776_v64_)
            nw147_[int(19)] = (len(d_777_v65_)) - (p0)
            nw147_[int(20)] = (d_775_v63_).cardinality
            nw147_[int(21)] = self.f2
            nw147_[int(22)] = default__.safeModuloInt((0) - (p0), p0)
            nw147_[int(23)] = len((d_737_v43_).f3)
            nw147_[int(24)] = p0
            d_778_v66_ = nw147_
            index126_ = default__.safeIndex(89, (d_778_v66_).length(0))
            (d_778_v66_)[index126_] = default__.safeModuloInt(p0, len(d_732_v41_))
            index127_ = default__.safeIndex(89, (d_778_v66_).length(0))
            (d_778_v66_)[index127_] = p0
        d_779_v67_: _dafny.Seq
        d_779_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhsiw"))
        d_779_v67_ = d_779_v67_
        d_780_v68_: str
        d_780_v68_ = _dafny.CodePoint('d')
        d_781_v69_: _dafny.Map
        d_781_v69_ = _dafny.Map({d_780_v68_: (d_779_v67_) + (d_779_v67_)})
        d_781_v69_ = (d_781_v69_).set(d_780_v68_, d_779_v67_)
        d_782_v70_: _dafny.Seq
        d_782_v70_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2])
        d_783_v71_: _dafny.Seq
        d_783_v71_ = _dafny.SeqWithoutIsStrInference([d_779_v67_, d_779_v67_])
        d_784_v73_: _dafny.Map
        d_784_v73_ = _dafny.Map({p2: p2})
        d_785_v74_: _dafny.Seq
        d_785_v74_ = _dafny.SeqWithoutIsStrInference([len(d_784_v73_)])
        d_786_v75_: _dafny.Map
        d_786_v75_ = _dafny.Map({d_785_v74_: True})
        d_787_v76_: _dafny.Array
        nw148_ = _dafny.Array(None, 19)
        nw148_[int(0)] = (d_782_v70_)[default__.safeIndex(p1, len(d_782_v70_))]
        nw148_[int(1)] = p2
        nw148_[int(2)] = p2
        nw148_[int(3)] = p2
        nw148_[int(4)] = ((d_783_v71_)[default__.safeIndex(self.f2, len(d_783_v71_))]) <= (d_779_v67_)
        nw148_[int(5)] = (self).fm8(d_779_v67_, p2, globalState)
        nw148_[int(6)] = (self).fm8(_dafny.SeqWithoutIsStrInference([d_780_v68_ for d_788_i6_ in range(default__.abs(-707))]), p2, globalState)
        nw148_[int(7)] = p2
        nw148_[int(8)] = p2
        nw148_[int(9)] = (self.f2) != (self.f2)
        nw148_[int(10)] = p2
        def iife57_():
            coll25_ = _dafny.Map()
            compr_25_: _dafny.Seq
            for compr_25_ in (d_786_v75_).keys.Elements:
                d_789_v72_: _dafny.Seq = compr_25_
                if (d_789_v72_) in (d_786_v75_):
                    coll25_[d_789_v72_] = p2
            return _dafny.Map(coll25_)
        nw148_[int(11)] = (p1) <= (len(iife57_()
        ))
        nw148_[int(12)] = p2
        nw148_[int(13)] = (d_782_v70_)[default__.safeIndex(self.f2, len(d_782_v70_))]
        nw148_[int(14)] = (684) == (self.f2)
        nw148_[int(15)] = (p2) or (p2)
        nw148_[int(16)] = p2
        nw148_[int(17)] = p2
        nw148_[int(18)] = (p2) == (p2)
        d_787_v76_ = nw148_
        index128_ = default__.safeIndex(53, (d_787_v76_).length(0))
        (d_787_v76_)[index128_] = p2
        index129_ = default__.safeIndex(53, (d_787_v76_).length(0))
        (d_787_v76_)[index129_] = p2
        d_790_i7_: int
        d_790_i7_ = 0
        with _dafny.label("1"):
            while not (True) or (False):
                with _dafny.c_label("1"):
                    if (d_790_i7_) >= (100):
                        raise _dafny.Break("1")
                    d_790_i7_ = (d_790_i7_) + (1)
                    d_791_v77_: _dafny.Map
                    d_791_v77_ = _dafny.Map({(d_780_v68_ if (d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))] else d_780_v68_): _dafny.Map({self.f2: p0})})
                    d_792_v78_: _dafny.MultiSet
                    d_792_v78_ = _dafny.MultiSet([p2])
                    d_793_v79_: _dafny.Map
                    d_793_v79_ = _dafny.Map({self.f2: (d_792_v78_).cardinality})
                    d_791_v77_ = (d_791_v77_).set(d_780_v68_, d_793_v79_)
                    if ((d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]) and ((d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]):
                        r0 = default__.safeModuloInt(self.f2, p1)
                        r0 = (0) - (p0)
                        index130_ = default__.safeIndex(53, (d_787_v76_).length(0))
                        (d_787_v76_)[index130_] = p2
                        d_794_v81_: _dafny.Set
                        d_794_v81_ = _dafny.Set({p1, self.f2})
                        index131_ = default__.safeIndex(53, (d_787_v76_).length(0))
                        def iife58_():
                            coll26_ = _dafny.Set()
                            compr_26_: int
                            for compr_26_ in _dafny.IntegerRange(0, -727):
                                d_795_v80_: int = compr_26_
                                if ((0) <= (d_795_v80_)) and ((d_795_v80_) < (-727)):
                                    coll26_ = coll26_.union(_dafny.Set([(d_795_v80_) + (p0)]))
                            return _dafny.Set(coll26_)
                        (d_787_v76_)[index131_] = ((iife58_()
                        ) | (d_794_v81_)).issubset(default__.fm31(not((d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]), p0, p2, p0, globalState))
                        r0 = -962
                    elif True:
                        d_792_v78_ = d_792_v78_
                        d_796_v82_: _dafny.MultiSet
                        d_796_v82_ = _dafny.MultiSet([d_779_v67_])
                        d_797_v83_: D0
                        d_797_v83_ = D0_DC1((d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))], (d_796_v82_).isdisjoint(d_796_v82_), d_780_v68_, self.f2)
                        index132_ = default__.safeIndex(53, (d_787_v76_).length(0))
                        index133_ = default__.safeIndex(53, (d_787_v76_).length(0))
                        rhs116_ = d_797_v83_
                        rhs117_ = not(not((d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]))
                        rhs118_ = _dafny.CodePoint('a')
                        rhs119_ = ((d_779_v67_) + (d_779_v67_)) < (d_779_v67_)
                        lhs56_ = d_787_v76_
                        lhs57_ = default__.safeIndex(53, (d_787_v76_).length(0))
                        lhs58_ = d_787_v76_
                        lhs59_ = default__.safeIndex(53, (d_787_v76_).length(0))
                        d_797_v83_ = rhs116_
                        lhs56_[lhs57_] = rhs117_
                        d_780_v68_ = rhs118_
                        lhs58_[lhs59_] = rhs119_
                        d_798_v84_: D7
                        d_798_v84_ = D7_DC14(d_785_v74_)
                        d_799_v85_: _dafny.Map
                        d_799_v85_ = _dafny.Map({(_dafny.MultiSet(default__.fm32(globalState))).intersection(_dafny.MultiSet([d_798_v84_])): len(d_779_v67_)})
                        d_800_v86_: _dafny.Map
                        d_800_v86_ = _dafny.Map({(d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]: p1})
                        d_801_v87_: _dafny.Set
                        d_801_v87_ = _dafny.Set({len(_dafny.Set({p0, 264, len(d_779_v67_), ((d_793_v79_)[(self).fm1(p2, globalState)] if ((self).fm1(p2, globalState)) in (d_793_v79_) else len(d_800_v86_)), 462}))})
                        d_802_v88_: C1
                        nw149_ = C1()
                        nw149_.ctor__(d_779_v67_)
                        d_802_v88_ = nw149_
                        d_803_v89_: _dafny.Map
                        d_803_v89_ = _dafny.Map({p2: d_802_v88_})
                        d_804_v90_: D8
                        d_804_v90_ = D8_DC16(d_802_v88_)
                        d_799_v85_ = (d_799_v85_).set(default__.fm33(d_780_v68_, True, len(d_801_v87_), p1, globalState), len(_dafny.SeqWithoutIsStrInference([(d_803_v89_).set(p2, (d_804_v90_).cf29), d_803_v89_, d_803_v89_, (d_803_v89_).set((d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))], d_802_v88_), d_803_v89_])))
                        r0 = self.f2
                        index134_ = default__.safeIndex(53, (d_787_v76_).length(0))
                        (d_787_v76_)[index134_] = (d_797_v83_).cf1
                    if p2:
                        d_805_v91_: _dafny.Map
                        d_805_v91_ = _dafny.Map({p1: (d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]})
                        d_806_v92_: _dafny.Map
                        d_806_v92_ = _dafny.Map({self.f2: d_780_v68_})
                        d_807_v93_: _dafny.MultiSet
                        d_807_v93_ = _dafny.MultiSet([p0, len((d_806_v92_).set(self.f2, d_780_v68_)), self.f2])
                        d_805_v91_ = (d_805_v91_).set(((d_807_v93_)[self.f2] if (self.f2) in (d_807_v93_) else p1), (d_807_v93_).ispropersubset(d_807_v93_))
                        d_808_v94_: _dafny.Map
                        d_808_v94_ = _dafny.Map({(d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]: d_787_v76_})
                        d_809_v95_: _dafny.Set
                        d_809_v95_ = _dafny.Set({(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ssamx")))) + (len(d_808_v94_))})
                        def iife59_():
                            coll27_ = _dafny.Set()
                            compr_27_: int
                            for compr_27_ in _dafny.IntegerRange(-95, 219):
                                d_810_v96_: int = compr_27_
                                if ((-95) <= (d_810_v96_)) and ((d_810_v96_) < (219)):
                                    coll27_ = coll27_.union(_dafny.Set([(d_810_v96_) * (685)]))
                            return _dafny.Set(coll27_)
                        d_809_v95_ = ((d_809_v95_ if p2 else iife59_()
                        )) | ((_dafny.Set({p0, p0})) | (d_809_v95_))
                        (self).f2 = self.f2
                        (self).f2 = self.f2
                        rhs120_ = d_779_v67_
                        rhs121_ = p1
                        rhs122_ = (self).fm1((False) and ((d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]), globalState)
                        d_779_v67_ = rhs120_
                        r0 = rhs121_
                        r0 = rhs122_
                    elif True:
                        d_811_v97_: _dafny.Array
                        nw150_ = _dafny.Array(_dafny.Seq({}), 3)
                        d_811_v97_ = nw150_
                        d_812_v98_: _dafny.Array
                        nw151_ = _dafny.Array(int(0), 18)
                        d_812_v98_ = nw151_
                        d_813_v99_: _dafny.Seq
                        d_813_v99_ = _dafny.SeqWithoutIsStrInference([d_812_v98_, d_812_v98_])
                        index135_ = default__.safeIndex(598, (d_811_v97_).length(0))
                        (d_811_v97_)[index135_] = d_813_v99_
                        index136_ = default__.safeIndex(598, (d_811_v97_).length(0))
                        (d_811_v97_)[index136_] = d_813_v99_
                        d_814_v100_: D6
                        d_814_v100_ = D6_DC13((d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))], p1)
                        index137_ = default__.safeIndex(53, (d_787_v76_).length(0))
                        (d_787_v76_)[index137_] = ((d_814_v100_ if False else d_814_v100_)).cf24
                        d_815_v101_: _dafny.Map
                        d_815_v101_ = _dafny.Map({default__.safeDivisionInt(p0, p1): d_812_v98_})
                        d_816_v102_: _dafny.Map
                        d_816_v102_ = _dafny.Map({(d_782_v70_)[default__.safeIndex(self.f2, len(d_782_v70_))]: d_815_v101_})
                        d_815_v101_ = ((d_816_v102_)[(self).fm8(d_779_v67_, p2, globalState)] if ((self).fm8(d_779_v67_, p2, globalState)) in (d_816_v102_) else d_815_v101_)
                        d_779_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "weab"))
                        d_817_v103_: _dafny.Array
                        nw152_ = _dafny.Array(_dafny.Array(None, 0), 24)
                        d_817_v103_ = nw152_
                        d_818_v104_: _dafny.Map
                        d_818_v104_ = _dafny.Map({(d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]: d_812_v98_})
                        nw153_ = _dafny.Array(None, 23)
                        nw153_[int(0)] = d_812_v98_
                        nw153_[int(1)] = d_812_v98_
                        nw153_[int(2)] = d_812_v98_
                        nw153_[int(3)] = d_812_v98_
                        nw153_[int(4)] = d_812_v98_
                        nw153_[int(5)] = (d_812_v98_ if (d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))] else d_812_v98_)
                        nw153_[int(6)] = d_812_v98_
                        nw153_[int(7)] = d_812_v98_
                        nw153_[int(8)] = d_812_v98_
                        nw153_[int(9)] = d_812_v98_
                        nw153_[int(10)] = d_812_v98_
                        nw153_[int(11)] = d_812_v98_
                        nw153_[int(12)] = d_812_v98_
                        nw153_[int(13)] = d_812_v98_
                        nw153_[int(14)] = d_812_v98_
                        nw153_[int(15)] = d_812_v98_
                        nw153_[int(16)] = d_812_v98_
                        nw153_[int(17)] = ((d_818_v104_)[(d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]] if ((d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]) in (d_818_v104_) else d_812_v98_)
                        nw153_[int(18)] = d_812_v98_
                        nw153_[int(19)] = d_812_v98_
                        nw153_[int(20)] = (d_812_v98_ if (d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))] else d_812_v98_)
                        nw153_[int(21)] = d_812_v98_
                        nw153_[int(22)] = d_812_v98_
                        d_817_v103_ = nw153_
                    r0 = p1
                    pass
            pass
        d_819_v105_: C1
        nw154_ = C1()
        nw154_.ctor__(d_779_v67_)
        d_819_v105_ = nw154_
        d_820_v106_: D8
        d_820_v106_ = D8_DC16(d_819_v105_)
        source16_ = d_820_v106_
        if source16_.is_DC17:
            d_821___mcc_h8_ = source16_.cf30
            d_822___mcc_h9_ = source16_.cf31
            d_823___mcc_h10_ = source16_.cf32
            d_824_cf32_ = d_823___mcc_h10_
            d_825_cf31_ = d_822___mcc_h9_
            d_826_cf30_ = d_821___mcc_h8_
            d_827_v107_: _dafny.Map
            d_827_v107_ = d_784_v73_
            d_828_v108_: _dafny.Map
            d_828_v108_ = _dafny.Map({d_827_v107_: p2})
            d_829_v109_: _dafny.Set
            d_829_v109_ = _dafny.Set({(d_828_v108_) | ((d_828_v108_).set(d_827_v107_, not(p2))), (d_828_v108_).set(d_827_v107_, d_826_cf30_)})
            d_830_v110_: _dafny.Map
            d_830_v110_ = _dafny.Map({d_828_v108_: p1})
            def iife60_():
                coll28_ = _dafny.Set()
                compr_28_: _dafny.Map
                for compr_28_ in (d_830_v110_).keys.Elements:
                    d_831_v111_: _dafny.Map = compr_28_
                    if (d_831_v111_) in (d_830_v110_):
                        coll28_ = coll28_.union(_dafny.Set([d_831_v111_]))
                return _dafny.Set(coll28_)
            d_829_v109_ = (iife60_()
            ) - (_dafny.Set({d_828_v108_, d_828_v108_}))
            d_832_v112_: _dafny.Array
            def lambda24_(d_833_v68_):
                def lambda25_(d_834_i8_):
                    return D7_DC15(self.f2, d_833_v68_)

                return lambda25_

            init12_ = lambda24_(d_780_v68_)
            nw155_ = _dafny.Array(None, 17)
            for i0_12_ in range(nw155_.length(0)):
                nw155_[i0_12_] = init12_(i0_12_)
            d_832_v112_ = nw155_
            d_832_v112_ = (d_832_v112_ if p2 else d_832_v112_)
            d_826_cf30_ = d_826_cf30_
            d_835_v113_: C1
            nw156_ = C1()
            nw156_.ctor__(d_779_v67_)
            d_835_v113_ = nw156_
        elif True:
            d_836___mcc_h11_ = source16_.cf29
            d_837_cf29_ = d_836___mcc_h11_
            index138_ = default__.safeIndex(53, (d_787_v76_).length(0))
            (d_787_v76_)[index138_] = ((p1) > (p1)) and (not((d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]))
            index139_ = default__.safeIndex(53, (d_787_v76_).length(0))
            (d_787_v76_)[index139_] = p2
            d_838_v114_: _dafny.Set
            d_838_v114_ = _dafny.Set({self.f2})
            d_839_v115_: _dafny.Array
            nw157_ = _dafny.Array(None, 2)
            nw157_[int(0)] = 600
            nw157_[int(1)] = len(d_838_v114_)
            d_839_v115_ = nw157_
            d_840_v116_: D4
            d_840_v116_ = D4_DC6(p1, d_839_v115_, p1)
            source17_ = d_840_v116_
            if source17_.is_DC6:
                d_841___mcc_h12_ = source17_.cf9
                d_842___mcc_h13_ = source17_.cf10
                d_843___mcc_h14_ = source17_.cf11
                d_844_cf11_ = d_843___mcc_h14_
                d_845_cf10_ = d_842___mcc_h13_
                d_846_cf9_ = d_841___mcc_h12_
                index140_ = default__.safeIndex(53, (d_787_v76_).length(0))
                (d_787_v76_)[index140_] = (self).fm8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjjaoweep")), (d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))], globalState)
                d_847_v117_: _dafny.Map
                d_847_v117_ = _dafny.Map({p0: (d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]})
                d_848_v118_: C0
                nw158_ = C0()
                nw158_.ctor__(p2, ((d_847_v117_)[self.f2] if (self.f2) in (d_847_v117_) else (d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]))
                d_848_v118_ = nw158_
                d_849_v119_: _dafny.Map
                d_849_v119_ = _dafny.Map({d_848_v118_: d_784_v73_})
                d_850_v120_: _dafny.Map
                d_850_v120_ = _dafny.Map({d_784_v73_: d_785_v74_})
                d_846_cf9_ = (0) - ((len((_dafny.Map({((d_849_v119_)[d_848_v118_] if (d_848_v118_) in (d_849_v119_) else d_784_v73_): d_785_v74_})) | (d_850_v120_))) - ((0) - (self.f2)))
                d_847_v117_ = (d_847_v117_).set(default__.safeDivisionInt(p0, p0), d_848_v118_.f5)
                d_851_v121_: _dafny.Set
                d_851_v121_ = _dafny.Set({d_848_v118_.f5, (d_819_v105_).fm11(globalState)})
                d_852_v122_: _dafny.Map
                d_852_v122_ = _dafny.Map({(d_848_v118_).fm14((d_819_v105_).fm11(globalState), p2, d_838_v114_, globalState): D6_DC11(default__.fm28(True, p2, d_848_v118_.f4, d_851_v121_, globalState))})
                d_853_v124_: D6
                d_853_v124_ = D6_DC11((d_819_v105_).f3)
                def iife61_():
                    coll29_ = _dafny.Set()
                    compr_29_: int
                    for compr_29_ in _dafny.IntegerRange(461, 893):
                        d_854_v123_: int = compr_29_
                        if ((461) <= (d_854_v123_)) and ((d_854_v123_) < (893)):
                            coll29_ = coll29_.union(_dafny.Set([(d_854_v123_) - (p0)]))
                    return _dafny.Set(coll29_)
                d_852_v122_ = (d_852_v122_).set((iife61_()
                ).ispropersubset(d_838_v114_), d_853_v124_)
            elif source17_.is_DC5:
                d_855___mcc_h15_ = source17_.cf8
                d_856_cf8_ = d_855___mcc_h15_
                index141_ = default__.safeIndex(53, (d_787_v76_).length(0))
                (d_787_v76_)[index141_] = p2
                nw159_ = _dafny.Array(int(0), 22)
                d_839_v115_ = nw159_
                (self).f2 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnoyfxqrs")))
                d_857_v125_: _dafny.MultiSet
                d_857_v125_ = _dafny.MultiSet([d_839_v115_, d_839_v115_])
                d_858_v126_: D5
                d_858_v126_ = D5_DC8(d_857_v125_)
                d_859_v127_: _dafny.Map
                d_859_v127_ = _dafny.Map({p1: d_858_v126_})
                d_860_v128_: _dafny.Seq
                d_860_v128_ = _dafny.SeqWithoutIsStrInference([d_859_v127_])
                r0 = len((_dafny.SeqWithoutIsStrInference([d_859_v127_])) + ((d_860_v128_).set(default__.safeIndex(p1, len(d_860_v128_)), d_859_v127_)))
            elif True:
                d_861___mcc_h16_ = source17_.cf12
                d_862_cf12_ = d_861___mcc_h16_
                d_863_v129_: _dafny.Map
                d_863_v129_ = _dafny.Map({d_780_v68_: (len(_dafny.Map({self.f2: p0}))) + (self.f2)})
                d_863_v129_ = (d_863_v129_).set(d_780_v68_, self.f2)
                index142_ = default__.safeIndex(53, (d_787_v76_).length(0))
                (d_787_v76_)[index142_] = (d_819_v105_).fm11(globalState)
                (self).f2 = p0
                d_837_cf29_ = d_837_cf29_
            d_864_v130_: _dafny.Map
            d_864_v130_ = _dafny.Map({(d_819_v105_).f3: p0})
            index143_ = default__.safeIndex(53, (d_787_v76_).length(0))
            index144_ = default__.safeIndex(53, (d_787_v76_).length(0))
            rhs123_ = d_820_v106_
            rhs124_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yac")): self.f2})
            rhs125_ = (p0) == (p1)
            rhs126_ = d_839_v115_
            rhs127_ = (d_787_v76_)[default__.safeIndex(53, (d_787_v76_).length(0))]
            lhs60_ = d_787_v76_
            lhs61_ = default__.safeIndex(53, (d_787_v76_).length(0))
            lhs62_ = d_787_v76_
            lhs63_ = default__.safeIndex(53, (d_787_v76_).length(0))
            d_820_v106_ = rhs123_
            d_864_v130_ = rhs124_
            lhs60_[lhs61_] = rhs125_
            d_839_v115_ = rhs126_
            lhs62_[lhs63_] = rhs127_
        d_865_v131_: _dafny.Map
        d_865_v131_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_780_v68_ for d_866_i9_ in range(default__.abs(820))])), self.f2])})
        d_867_v132_: _dafny.Array
        def lambda26_(d_868_p0_):
            def lambda27_(d_869_i11_):
                return default__.safeDivisionInt(d_869_i11_, d_868_p0_)

            return lambda27_

        init13_ = lambda26_(p0)
        nw160_ = _dafny.Array(None, 21)
        for i0_13_ in range(nw160_.length(0)):
            nw160_[i0_13_] = init13_(i0_13_)
        d_867_v132_ = nw160_
        r0 = default__.safeModuloInt(default__.safeModuloInt(p1, len(_dafny.Map({len(((d_865_v131_)[p2] if (p2) in (d_865_v131_) else _dafny.SeqWithoutIsStrInference([73 for d_870_i10_ in range(default__.abs(482))]))): d_867_v132_}))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(_dafny.MultiSet([self.f2])).cardinality: (d_819_v105_).fm11(globalState)}))])))
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        if p3:
            d_871_v0_: bool
            d_871_v0_ = True
            d_871_v0_ = d_871_v0_
            d_872_v1_: _dafny.Map
            d_872_v1_ = _dafny.Map({p1: self.f2})
            d_873_v2_: _dafny.Seq
            d_873_v2_ = _dafny.SeqWithoutIsStrInference([p3])
            d_872_v1_ = (d_872_v1_).set(p1, (self.f2) - ((_dafny.MultiSet(d_873_v2_)).cardinality))
            d_874_v3_: str
            d_874_v3_ = _dafny.CodePoint('m')
            d_875_v4_: _dafny.Seq
            d_875_v4_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), d_874_v3_])
            d_876_v5_: _dafny.Map
            d_876_v5_ = _dafny.Map({p2: (d_875_v4_)[default__.safeIndex(self.f2, len(d_875_v4_))]})
            d_877_v6_: _dafny.Set
            d_877_v6_ = _dafny.Set({d_876_v5_, d_876_v5_})
            d_878_v7_: _dafny.Set
            d_878_v7_ = _dafny.Set({d_871_v0_, p3})
            d_879_v8_: _dafny.Map
            d_879_v8_ = _dafny.Map({d_876_v5_: d_878_v7_})
            d_880_v9_: _dafny.Map
            d_880_v9_ = _dafny.Map({d_877_v6_: d_879_v8_})
            d_880_v9_ = (d_880_v9_).set((d_877_v6_) - (_dafny.Set({d_876_v5_, d_876_v5_, _dafny.Map({False: d_874_v3_}), d_876_v5_, d_876_v5_})), _dafny.Map({d_876_v5_: d_878_v7_}))
            d_881_v10_: _dafny.Set
            d_881_v10_ = _dafny.Set({self.f2, self.f2})
            if (d_881_v10_).issubset(d_881_v10_):
                d_873_v2_ = _dafny.SeqWithoutIsStrInference([p2, d_871_v0_, p3, p3, not(p3)])
                d_882_v11_: _dafny.Array
                nw161_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
                d_882_v11_ = nw161_
                d_883_v12_: _dafny.Seq
                d_883_v12_ = _dafny.SeqWithoutIsStrInference([d_875_v4_, d_875_v4_, d_875_v4_])
                index145_ = default__.safeIndex(800, (d_882_v11_).length(0))
                (d_882_v11_)[index145_] = ((d_883_v12_)[default__.safeIndex(self.f2, len(d_883_v12_))]) + (d_875_v4_)
                d_884_v13_: _dafny.Array
                nw162_ = _dafny.Array(False, 26)
                d_884_v13_ = nw162_
                d_885_v14_: _dafny.MultiSet
                d_885_v14_ = _dafny.MultiSet([p2])
                d_886_v15_: _dafny.Map
                d_886_v15_ = _dafny.Map({(self).fm1(d_871_v0_, globalState): (d_885_v14_).cardinality})
                d_887_v16_: _dafny.MultiSet
                d_887_v16_ = _dafny.MultiSet([d_886_v15_, _dafny.Map({len(_dafny.SeqWithoutIsStrInference([p2, d_871_v0_])): self.f2})])
                index146_ = default__.safeIndex(616, (d_884_v13_).length(0))
                (d_884_v13_)[index146_] = (d_887_v16_).ispropersubset(d_887_v16_)
                d_888_v17_: _dafny.Seq
                d_888_v17_ = _dafny.SeqWithoutIsStrInference([((d_885_v14_)[p2] if (p2) in (d_885_v14_) else self.f2), self.f2, self.f2, 256, self.f2])
                index147_ = default__.safeIndex(800, (d_882_v11_).length(0))
                index148_ = default__.safeIndex(616, (d_884_v13_).length(0))
                rhs128_ = (d_888_v17_) <= (d_888_v17_)
                rhs129_ = d_875_v4_
                rhs130_ = (self.f2) - ((0) - (self.f2))
                rhs131_ = (self).fm1((False if True else p3), globalState)
                rhs132_ = d_871_v0_
                lhs64_ = d_882_v11_
                lhs65_ = default__.safeIndex(800, (d_882_v11_).length(0))
                lhs66_ = self
                lhs67_ = d_884_v13_
                lhs68_ = default__.safeIndex(616, (d_884_v13_).length(0))
                d_871_v0_ = rhs128_
                lhs64_[lhs65_] = rhs129_
                lhs66_.f2 = rhs130_
                r0 = rhs131_
                lhs67_[lhs68_] = rhs132_
                d_889_v18_: C0
                nw163_ = C0()
                nw163_.ctor__((self).fm8(default__.fm28((d_884_v13_)[default__.safeIndex(616, (d_884_v13_).length(0))], p3, p2, d_878_v7_, globalState), p2, globalState), (d_873_v2_)[default__.safeIndex(self.f2, len(d_873_v2_))])
                d_889_v18_ = nw163_
                d_884_v13_ = d_884_v13_
                index149_ = default__.safeIndex(616, (d_884_v13_).length(0))
                (d_884_v13_)[index149_] = ((d_884_v13_)[default__.safeIndex(616, (d_884_v13_).length(0))] if (d_884_v13_)[default__.safeIndex(616, (d_884_v13_).length(0))] else (self.f2) < (self.f2))
            elif True:
                d_890_v19_: _dafny.Array
                nw164_ = _dafny.Array(None, 26)
                d_890_v19_ = nw164_
                d_891_v20_: C0
                nw165_ = C0()
                nw165_.ctor__(p2, (d_873_v2_)[default__.safeIndex(len(d_875_v4_), len(d_873_v2_))])
                d_891_v20_ = nw165_
                index150_ = default__.safeIndex(224, (d_890_v19_).length(0))
                (d_890_v19_)[index150_] = d_891_v20_
                index151_ = default__.safeIndex(224, (d_890_v19_).length(0))
                (d_890_v19_)[index151_] = (d_891_v20_ if d_891_v20_.f4 else d_891_v20_)
                r1 = (0) - (self.f2)
                d_892_v22_: D5
                def iife62_():
                    coll30_ = _dafny.Map()
                    compr_30_: int
                    for compr_30_ in _dafny.IntegerRange(338, 202):
                        d_893_v21_: int = compr_30_
                        if ((338) <= (d_893_v21_)) and ((d_893_v21_) < (202)):
                            coll30_[default__.safeModuloInt(d_893_v21_, ((p0)[self.f2] if (self.f2) in (p0) else 198))] = d_891_v20_.f5
                    return _dafny.Map(coll30_)
                d_892_v22_ = D5_DC10((d_891_v20_).fm14(d_871_v0_, d_891_v20_.f4, d_881_v10_, globalState), self.f2, _dafny.Map({d_875_v4_: iife62_()
}), p2)
                d_892_v22_ = d_892_v22_
                d_894_v23_: C0
                nw166_ = C0()
                nw166_.ctor__(d_871_v0_, d_871_v0_)
                d_894_v23_ = nw166_
                (d_894_v23_).f4 = (d_894_v23_).fm14(d_894_v23_.f5, (_dafny.MultiSet([p2, d_871_v0_])).issubset(_dafny.MultiSet(d_873_v2_)), (d_881_v10_) - (d_881_v10_), globalState)
            d_895_v24_: _dafny.Map
            d_895_v24_ = _dafny.Map({(0) - (self.f2): d_871_v0_})
            d_896_v25_: _dafny.Map
            d_896_v25_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrjp")): d_895_v24_})
            d_897_v26_: D5
            d_897_v26_ = D5_DC10(not(p2), self.f2, d_896_v25_, p2)
            d_871_v0_ = (((d_897_v26_).cf18) <= (876) if p2 else p2)
        elif True:
            d_898_v27_: bool
            d_898_v27_ = True
            d_898_v27_ = not (p3) or (d_898_v27_)
            d_899_v28_: str
            d_899_v28_ = _dafny.CodePoint('v')
            d_900_v29_: _dafny.MultiSet
            d_900_v29_ = _dafny.MultiSet([d_899_v28_, d_899_v28_])
            d_901_v30_: _dafny.Seq
            d_901_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
            d_902_v31_: _dafny.Array
            nw167_ = _dafny.Array(None, 7)
            nw167_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmqami"))
            nw167_[int(1)] = (d_901_v30_) + (d_901_v30_)
            nw167_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tuqtafex"))
            nw167_[int(3)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))) + (d_901_v30_)
            nw167_[int(4)] = d_901_v30_
            nw167_[int(5)] = (d_901_v30_) + (d_901_v30_)
            nw167_[int(6)] = d_901_v30_
            d_902_v31_ = nw167_
            index152_ = default__.safeIndex(981, (d_902_v31_).length(0))
            (d_902_v31_)[index152_] = d_901_v30_
            d_903_v32_: _dafny.Array
            def lambda28_(d_904_p3_):
                def lambda29_(d_905_i0_):
                    return d_904_p3_

                return lambda29_

            init14_ = lambda28_(p3)
            nw168_ = _dafny.Array(None, 20)
            for i0_14_ in range(nw168_.length(0)):
                nw168_[i0_14_] = init14_(i0_14_)
            d_903_v32_ = nw168_
            d_906_v33_: _dafny.Map
            d_906_v33_ = _dafny.Map({d_903_v32_: _dafny.MultiSet([self.f2, self.f2])})
            d_907_v34_: _dafny.Map
            d_907_v34_ = _dafny.Map({self.f2: (d_900_v29_).intersection(d_900_v29_)})
            index153_ = default__.safeIndex(981, (d_902_v31_).length(0))
            rhs133_ = ((d_907_v34_)[self.f2] if (self.f2) in (d_907_v34_) else (_dafny.MultiSet([_dafny.CodePoint('u'), d_899_v28_])).set(_dafny.CodePoint('d'), default__.abs(self.f2)))
            rhs134_ = p2
            rhs135_ = d_901_v30_
            rhs136_ = d_906_v33_
            rhs137_ = (self).fm1(False, globalState)
            lhs69_ = d_902_v31_
            lhs70_ = default__.safeIndex(981, (d_902_v31_).length(0))
            lhs71_ = self
            d_900_v29_ = rhs133_
            d_898_v27_ = rhs134_
            lhs69_[lhs70_] = rhs135_
            d_906_v33_ = rhs136_
            lhs71_.f2 = rhs137_
            d_908_v35_: _dafny.Set
            d_908_v35_ = _dafny.Set({self.f2})
            d_909_v36_: _dafny.MultiSet
            d_909_v36_ = _dafny.MultiSet([not(d_898_v27_)])
            d_910_v37_: _dafny.Seq
            d_910_v37_ = _dafny.SeqWithoutIsStrInference([d_909_v36_, d_909_v36_, d_909_v36_, d_909_v36_])
            d_911_v38_: _dafny.Map
            d_911_v38_ = _dafny.Map({d_908_v35_: (d_910_v37_)[default__.safeIndex(self.f2, len(d_910_v37_))]})
            d_911_v38_ = (d_911_v38_).set(d_908_v35_, _dafny.MultiSet([(self).fm8((d_902_v31_)[default__.safeIndex(981, (d_902_v31_).length(0))], p3, globalState)]))
            d_912_v39_: _dafny.Array
            nw169_ = _dafny.Array(_dafny.Array(None, 0), 21)
            d_912_v39_ = nw169_
            d_913_v40_: _dafny.Array
            nw170_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_913_v40_ = nw170_
            index154_ = default__.safeIndex(166, (d_912_v39_).length(0))
            (d_912_v39_)[index154_] = d_913_v40_
            index155_ = default__.safeIndex(44, (d_903_v32_).length(0))
            (d_903_v32_)[index155_] = p3
            index156_ = default__.safeIndex(434, (p1).length(0))
            (p1)[index156_] = self.f2
            d_914_v41_: _dafny.Map
            d_914_v41_ = _dafny.Map({(self).fm8((d_902_v31_)[default__.safeIndex(981, (d_902_v31_).length(0))], d_898_v27_, globalState): d_913_v40_})
            index157_ = default__.safeIndex(166, (d_912_v39_).length(0))
            index158_ = default__.safeIndex(44, (d_903_v32_).length(0))
            index159_ = default__.safeIndex(434, (p1).length(0))
            rhs138_ = (self.f2) - (self.f2)
            rhs139_ = ((d_914_v41_)[d_898_v27_] if (d_898_v27_) in (d_914_v41_) else d_913_v40_)
            rhs140_ = d_898_v27_
            rhs141_ = self.f2
            lhs72_ = self
            lhs73_ = d_912_v39_
            lhs74_ = default__.safeIndex(166, (d_912_v39_).length(0))
            lhs75_ = d_903_v32_
            lhs76_ = default__.safeIndex(44, (d_903_v32_).length(0))
            lhs77_ = p1
            lhs78_ = default__.safeIndex(434, (p1).length(0))
            lhs72_.f2 = rhs138_
            lhs73_[lhs74_] = rhs139_
            lhs75_[lhs76_] = rhs140_
            lhs77_[lhs78_] = rhs141_
            d_915_v42_: C1
            nw171_ = C1()
            nw171_.ctor__(d_901_v30_)
            d_915_v42_ = nw171_
            d_916_v43_: _dafny.Map
            d_916_v43_ = _dafny.Map({d_899_v28_: d_915_v42_})
            d_916_v43_ = (d_916_v43_).set(d_899_v28_, d_915_v42_)
        d_917_v44_: _dafny.Array
        nw172_ = _dafny.Array(False, 7)
        d_917_v44_ = nw172_
        index160_ = default__.safeIndex(698, (d_917_v44_).length(0))
        (d_917_v44_)[index160_] = not(p3)
        d_918_v45_: D6
        d_918_v45_ = D6_DC13(p2, self.f2)
        d_919_v46_: _dafny.Seq
        d_919_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgoffmr"))
        index161_ = default__.safeIndex(698, (d_917_v44_).length(0))
        rhs142_ = ((d_919_v46_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rabeglte"))) if not((p3) or ((d_918_v45_).cf24)) else p2)
        rhs143_ = self.f2
        lhs79_ = d_917_v44_
        lhs80_ = default__.safeIndex(698, (d_917_v44_).length(0))
        lhs79_[lhs80_] = rhs142_
        r1 = rhs143_
        d_920_v47_: _dafny.Seq
        d_920_v47_ = _dafny.SeqWithoutIsStrInference([p3])
        d_921_v48_: _dafny.Map
        d_921_v48_ = _dafny.Map({self.f2: d_920_v47_})
        d_922_i1_: int
        d_922_i1_ = 0
        with _dafny.label("2"):
            while (((d_921_v48_)[self.f2] if (self.f2) in (d_921_v48_) else d_920_v47_)) != ((_dafny.SeqWithoutIsStrInference([p3, (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))]])) + (d_920_v47_)):
                with _dafny.c_label("2"):
                    if (d_922_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_922_i1_ = (d_922_i1_) + (1)
                    d_923_v49_: _dafny.Map
                    d_923_v49_ = _dafny.Map({self.f2: self.f2})
                    index162_ = default__.safeIndex(897, (p1).length(0))
                    (p1)[index162_] = len(_dafny.Map({d_923_v49_: (0) - (self.f2)}))
                    d_924_v50_: str
                    d_924_v50_ = _dafny.CodePoint('v')
                    d_925_v51_: _dafny.Map
                    d_925_v51_ = _dafny.Map({d_924_v50_: self.f2})
                    index163_ = default__.safeIndex(897, (p1).length(0))
                    (p1)[index163_] = len(d_925_v51_)
                    if ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfbw"))) in (default__.fm34(p3, globalState)) if not(p3) else p2):
                        index164_ = default__.safeIndex(698, (d_917_v44_).length(0))
                        (d_917_v44_)[index164_] = (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))]
                        r0 = ((p1)[default__.safeIndex(897, (p1).length(0))]) * ((0) - ((0) - ((p1)[default__.safeIndex(897, (p1).length(0))])))
                        d_926_v52_: _dafny.Array
                        nw173_ = _dafny.Array(_dafny.Array(None, 0), 25)
                        d_926_v52_ = nw173_
                        d_927_v53_: _dafny.MultiSet
                        d_927_v53_ = _dafny.MultiSet([d_924_v50_, d_924_v50_, d_924_v50_, d_924_v50_, d_924_v50_])
                        index165_ = default__.safeIndex(897, (p1).length(0))
                        index166_ = default__.safeIndex(897, (p1).length(0))
                        rhs144_ = (d_926_v52_ if not(p2) else (d_926_v52_ if p3 else d_926_v52_))
                        rhs145_ = _dafny.MultiSet(d_919_v46_)
                        rhs146_ = (self).fm1(p3, globalState)
                        rhs147_ = self.f2
                        lhs81_ = p1
                        lhs82_ = default__.safeIndex(897, (p1).length(0))
                        lhs83_ = p1
                        lhs84_ = default__.safeIndex(897, (p1).length(0))
                        d_926_v52_ = rhs144_
                        d_927_v53_ = rhs145_
                        lhs81_[lhs82_] = rhs146_
                        lhs83_[lhs84_] = rhs147_
                        r0 = (p1)[default__.safeIndex(897, (p1).length(0))]
                        d_928_v54_: _dafny.Map
                        d_928_v54_ = _dafny.Map({(d_923_v49_) | (d_923_v49_): d_920_v47_})
                        d_928_v54_ = (d_928_v54_).set(d_923_v49_, (d_920_v47_) + (_dafny.SeqWithoutIsStrInference([p3])))
                    elif True:
                        d_917_v44_ = d_917_v44_
                        d_929_v55_: _dafny.Array
                        nw174_ = _dafny.Array(int(0), 8)
                        d_929_v55_ = nw174_
                        d_930_v56_: _dafny.Map
                        d_930_v56_ = _dafny.Map({default__.safeDivisionInt(self.f2, (p0).cardinality): (d_929_v55_ if p2 else d_929_v55_)})
                        d_931_v57_: _dafny.Map
                        d_931_v57_ = _dafny.Map({p3: d_919_v46_})
                        d_932_v58_: _dafny.Map
                        d_932_v58_ = _dafny.Map({(self).fm8(_dafny.SeqWithoutIsStrInference([d_924_v50_ for d_933_i2_ in range(default__.abs(170))]), p2, globalState): len(((d_931_v57_)[(d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))]] if ((d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))]) in (d_931_v57_) else d_919_v46_))})
                        d_929_v55_ = ((d_930_v56_)[len((d_932_v58_) | (_dafny.Map({(self).fm8(default__.fm28(p2, (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))], p3, _dafny.Set({p2, True}), globalState), (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))], globalState): 766})))] if (len((d_932_v58_) | (_dafny.Map({(self).fm8(default__.fm28(p2, (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))], p3, _dafny.Set({p2, True}), globalState), (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))], globalState): 766})))) in (d_930_v56_) else (d_929_v55_ if p3 else d_929_v55_))
                        d_934_v59_: C0
                        nw175_ = C0()
                        nw175_.ctor__(((p1)[default__.safeIndex(897, (p1).length(0))]) <= (len(_dafny.SeqWithoutIsStrInference([d_924_v50_ for d_935_i3_ in range(default__.abs(577))]))), (self).fm8((d_919_v46_).set(default__.safeIndex(self.f2, len(d_919_v46_)), d_924_v50_), False, globalState))
                        d_934_v59_ = nw175_
                        (d_934_v59_).f4 = (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))]
                        d_936_v60_: _dafny.Map
                        d_936_v60_ = _dafny.Map({p2: d_934_v59_.f5})
                        (d_934_v59_).f5 = ((d_936_v60_)[not((d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))])] if (not((d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))])) in (d_936_v60_) else p3)
                    d_937_v61_: C1
                    nw176_ = C1()
                    nw176_.ctor__(((d_919_v46_ if (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))] else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvbp")))).set(default__.safeIndex((p1)[default__.safeIndex(897, (p1).length(0))], len((d_919_v46_ if (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))] else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvbp"))))), d_924_v50_))
                    d_937_v61_ = nw176_
                    d_938_v62_: C0
                    nw177_ = C0()
                    nw177_.ctor__((d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))], (-802) >= (940))
                    d_938_v62_ = nw177_
                    pass
            pass
        d_939_v63_: _dafny.Map
        d_939_v63_ = _dafny.Map({(self.f2) + (self.f2): self.f2})
        d_940_v64_: _dafny.MultiSet
        d_940_v64_ = _dafny.MultiSet([(d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))], p2])
        index167_ = default__.safeIndex(449, (p1).length(0))
        (p1)[index167_] = (125) + ((d_940_v64_).cardinality)
        d_941_v65_: D7
        d_941_v65_ = D7_DC15(len(_dafny.Set({(d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))], True})), _dafny.CodePoint('n'))
        d_942_v66_: str
        d_942_v66_ = _dafny.CodePoint('n')
        d_943_v67_: D0
        d_943_v67_ = D0_DC1(not(p2), (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))], d_942_v66_, self.f2)
        pat_let_tv18_ = d_917_v44_
        pat_let_tv19_ = d_917_v44_
        index168_ = default__.safeIndex(449, (p1).length(0))
        index169_ = default__.safeIndex(698, (d_917_v44_).length(0))
        def lambda30_(source18_):
            if source18_.is_DC1:
                d_944___mcc_h0_ = source18_.cf1
                d_945___mcc_h1_ = source18_.cf2
                d_946___mcc_h2_ = source18_.cf3
                d_947___mcc_h3_ = source18_.cf4
                d_948_cf4_ = d_947___mcc_h3_
                d_949_cf3_ = d_946___mcc_h2_
                d_950_cf2_ = d_945___mcc_h1_
                d_951_cf1_ = d_944___mcc_h0_
                return (pat_let_tv19_)[default__.safeIndex(698, (pat_let_tv18_).length(0))]
            elif True:
                d_952___mcc_h4_ = source18_.cf0
                d_953_cf0_ = d_952___mcc_h4_
                return (172) <= (self.f2)

        rhs148_ = default__.fm28((d_920_v47_)[default__.safeIndex(self.f2, len(d_920_v47_))], True, p3, _dafny.Set({False, (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))]}), globalState)
        rhs149_ = d_939_v63_
        rhs150_ = (self.f2) + (self.f2)
        rhs151_ = (d_941_v65_).cf27
        rhs152_ = lambda30_(d_943_v67_)
        lhs85_ = self
        lhs86_ = p1
        lhs87_ = default__.safeIndex(449, (p1).length(0))
        lhs88_ = d_917_v44_
        lhs89_ = default__.safeIndex(698, (d_917_v44_).length(0))
        d_919_v46_ = rhs148_
        d_939_v63_ = rhs149_
        lhs85_.f2 = rhs150_
        lhs86_[lhs87_] = rhs151_
        lhs88_[lhs89_] = rhs152_
        d_954_v68_: _dafny.Set
        d_954_v68_ = _dafny.Set({p2, (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))], True})
        hi3_ = (p1)[default__.safeIndex(449, (p1).length(0))]
        for d_955_i4_ in range(((p1)[default__.safeIndex(449, (p1).length(0))] if p2 else len(_dafny.Map({len(d_954_v68_): (p1)[default__.safeIndex(449, (p1).length(0))]}))), hi3_):
            pat_let_tv20_ = d_919_v46_
            pat_let_tv21_ = p2
            d_956_v69_: _dafny.Set
            def iife63_(_pat_let16_0):
                def iife64_(d_957_dt__update__tmp_h0_):
                    def iife65_(_pat_let17_0):
                        def iife66_(d_958_dt__update_hcf4_h0_):
                            def iife67_(_pat_let18_0):
                                def iife68_(d_959_dt__update_hcf2_h0_):
                                    return D0_DC1((d_957_dt__update__tmp_h0_).cf1, d_959_dt__update_hcf2_h0_, (d_957_dt__update__tmp_h0_).cf3, d_958_dt__update_hcf4_h0_)
                                return iife68_(_pat_let18_0)
                            return iife67_(pat_let_tv21_)
                        return iife66_(_pat_let17_0)
                    return iife65_(len(pat_let_tv20_))
                return iife64_(_pat_let16_0)
            d_956_v69_ = _dafny.Set({d_942_v66_, (iife63_(d_943_v67_)).cf3})
            index170_ = default__.safeIndex(698, (d_917_v44_).length(0))
            (d_917_v44_)[index170_] = not(((d_956_v69_).isdisjoint(_dafny.Set({_dafny.CodePoint('g')})) if (p0).ispropersubset(default__.fm35(d_942_v66_, (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))], _dafny.MultiSet([p3]), globalState)) else p3))
            index171_ = default__.safeIndex(449, (p1).length(0))
            index172_ = default__.safeIndex(698, (d_917_v44_).length(0))
            rhs153_ = (self).fm1(p3, globalState)
            rhs154_ = (d_917_v44_)[default__.safeIndex(698, (d_917_v44_).length(0))]
            lhs90_ = p1
            lhs91_ = default__.safeIndex(449, (p1).length(0))
            lhs92_ = d_917_v44_
            lhs93_ = default__.safeIndex(698, (d_917_v44_).length(0))
            lhs90_[lhs91_] = rhs153_
            lhs92_[lhs93_] = rhs154_
            d_940_v64_ = d_940_v64_
            index173_ = default__.safeIndex(698, (d_917_v44_).length(0))
            (d_917_v44_)[index173_] = ((p0).cardinality) != (self.f2)
        d_960_v70_: _dafny.MultiSet
        d_960_v70_ = _dafny.MultiSet([((d_939_v63_)[215] if (215) in (d_939_v63_) else 405), (0) - ((p1)[default__.safeIndex(449, (p1).length(0))]), ((p1)[default__.safeIndex(449, (p1).length(0))]) + ((0) - (self.f2))])
        index174_ = default__.safeIndex(698, (d_917_v44_).length(0))
        rhs155_ = p2
        rhs156_ = (p0).intersection(d_960_v70_)
        lhs94_ = d_917_v44_
        lhs95_ = default__.safeIndex(698, (d_917_v44_).length(0))
        lhs94_[lhs95_] = rhs155_
        d_960_v70_ = rhs156_
        d_961_v71_: _dafny.Set
        d_961_v71_ = _dafny.Set({(p1)[default__.safeIndex(449, (p1).length(0))]})
        d_962_v72_: _dafny.Map
        d_962_v72_ = _dafny.Map({(p1)[default__.safeIndex(449, (p1).length(0))]: p2})
        d_963_v73_: _dafny.Map
        d_963_v73_ = _dafny.Map({d_919_v46_: d_962_v72_})
        d_964_v74_: D5
        d_964_v74_ = D5_DC10(p3, (0) - ((0) - (len(d_961_v71_))), d_963_v73_, True)
        r0 = (d_964_v74_).cf18
        d_965_v75_: _dafny.Seq
        d_965_v75_ = _dafny.SeqWithoutIsStrInference([d_961_v71_])
        r1 = len((d_965_v75_)[default__.safeIndex((p1)[default__.safeIndex(449, (p1).length(0))], len(d_965_v75_))])
        return r0, r1

    def m6(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        d_966_v0_: _dafny.Array
        def lambda31_(d_967_i0_):
            return default__.safeDivisionInt(d_967_i0_, self.f2)

        init15_ = lambda31_
        nw178_ = _dafny.Array(None, 29)
        for i0_15_ in range(nw178_.length(0)):
            nw178_[i0_15_] = init15_(i0_15_)
        d_966_v0_ = nw178_
        d_968_v1_: _dafny.Set
        d_968_v1_ = _dafny.Set({p0})
        d_969_v2_: _dafny.Seq
        d_969_v2_ = _dafny.SeqWithoutIsStrInference([603])
        index175_ = default__.safeIndex(52, (d_966_v0_).length(0))
        def iife69_():
            coll31_ = _dafny.Set()
            compr_31_: int
            for compr_31_ in (d_969_v2_).Elements:
                d_970_v3_: int = compr_31_
                if (d_970_v3_) in (d_969_v2_):
                    coll31_ = coll31_.union(_dafny.Set([default__.safeDivisionInt(d_970_v3_, 916)]))
            return _dafny.Set(coll31_)
        (d_966_v0_)[index175_] = len((d_968_v1_) | (iife69_()
        ))
        index176_ = default__.safeIndex(52, (d_966_v0_).length(0))
        (d_966_v0_)[index176_] = (len((D6_DC11(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_971_i1_ in range(default__.abs(188))]))).cf21)) - (p0)
        d_972_v4_: _dafny.Array
        nw179_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
        d_972_v4_ = nw179_
        d_973_v5_: _dafny.Map
        d_973_v5_ = _dafny.Map({d_972_v4_: d_966_v0_})
        d_973_v5_ = (d_973_v5_).set(d_972_v4_, d_966_v0_)
        d_974_v6_: bool
        d_974_v6_ = True
        rhs157_ = d_974_v6_
        rhs158_ = not (d_974_v6_) or (d_974_v6_)
        rhs159_ = d_974_v6_
        r0 = rhs157_
        r0 = rhs158_
        r0 = rhs159_
        d_975_v7_: _dafny.Map
        d_975_v7_ = _dafny.Map({d_974_v6_: d_974_v6_})
        d_976_v8_: _dafny.Seq
        d_976_v8_ = _dafny.SeqWithoutIsStrInference([d_974_v6_])
        d_977_v9_: _dafny.Array
        nw180_ = _dafny.Array(None, 5)
        nw180_[int(0)] = d_974_v6_
        nw180_[int(1)] = d_974_v6_
        nw180_[int(2)] = (d_976_v8_)[default__.safeIndex((d_966_v0_)[default__.safeIndex(52, (d_966_v0_).length(0))], len(d_976_v8_))]
        nw180_[int(3)] = d_974_v6_
        nw180_[int(4)] = d_974_v6_
        d_977_v9_ = nw180_
        d_978_v10_: _dafny.Map
        d_978_v10_ = _dafny.Map({default__.safeDivisionInt(len(d_975_v7_), (self).fm1(d_974_v6_, globalState)): d_977_v9_})
        d_979_v11_: _dafny.Map
        d_979_v11_ = _dafny.Map({d_974_v6_: p0})
        d_980_v12_: _dafny.Seq
        d_980_v12_ = _dafny.SeqWithoutIsStrInference([d_979_v11_])
        d_978_v10_ = (d_978_v10_).set(len(d_980_v12_), d_977_v9_)
        r0 = d_974_v6_
        d_981_i2_: int
        d_981_i2_ = 0
        with _dafny.label("3"):
            while d_974_v6_:
                with _dafny.c_label("3"):
                    if (d_981_i2_) >= (100):
                        raise _dafny.Break("3")
                    d_981_i2_ = (d_981_i2_) + (1)
                    d_982_v13_: str
                    d_982_v13_ = _dafny.CodePoint('y')
                    d_983_v14_: _dafny.Seq
                    d_983_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxs"))
                    rhs160_ = ((d_975_v7_)[d_974_v6_] if (d_974_v6_) in (d_975_v7_) else d_974_v6_)
                    rhs161_ = not(((d_969_v2_) + (d_969_v2_)) < ((d_969_v2_) + (default__.fm36(d_974_v6_, d_982_v13_, len(_dafny.Set({d_974_v6_, (self).fm8(d_983_v14_, d_974_v6_, globalState)})), globalState))))
                    d_974_v6_ = rhs160_
                    d_974_v6_ = rhs161_
                    if not(d_974_v6_):
                        d_982_v13_ = d_982_v13_
                        d_974_v6_ = d_974_v6_
                        d_974_v6_ = d_974_v6_
                        d_984_v15_: _dafny.Set
                        d_984_v15_ = _dafny.Set({not(d_974_v6_), d_974_v6_, (p0) == (self.f2)})
                        d_984_v15_ = (d_984_v15_) | (d_984_v15_)
                        d_974_v6_ = ((_dafny.SeqWithoutIsStrInference([675])) + ((d_969_v2_).set(default__.safeIndex(p0, len(d_969_v2_)), len(d_983_v14_)))) <= ((_dafny.SeqWithoutIsStrInference([(d_966_v0_)[default__.safeIndex(52, (d_966_v0_).length(0))]])) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_976_v8_)).cardinality for d_985_i3_ in range(default__.abs(-907))])))
                    elif True:
                        index177_ = default__.safeIndex(356, (d_977_v9_).length(0))
                        (d_977_v9_)[index177_] = d_974_v6_
                        d_986_v16_: _dafny.Set
                        d_986_v16_ = _dafny.Set({d_982_v13_, default__.fm37((self).fm1(d_974_v6_, globalState), globalState)})
                        index178_ = default__.safeIndex(356, (d_977_v9_).length(0))
                        (d_977_v9_)[index178_] = not ((d_986_v16_) == (d_986_v16_)) or ((_dafny.MultiSet(d_976_v8_)).isdisjoint(_dafny.MultiSet([True])))
                        (self).f2 = (d_966_v0_)[default__.safeIndex(52, (d_966_v0_).length(0))]
                        d_987_v17_: D0
                        d_987_v17_ = D0_DC0(not((d_977_v9_)[default__.safeIndex(356, (d_977_v9_).length(0))]))
                        d_974_v6_ = ((self).fm1((d_987_v17_).cf0, globalState)) not in (d_969_v2_)
                        d_988_v18_: _dafny.Array
                        out22_: _dafny.Array
                        out22_ = default__.m0(d_974_v6_, (d_977_v9_)[default__.safeIndex(356, (d_977_v9_).length(0))], globalState)
                        d_988_v18_ = out22_
                        d_989_v19_: C1
                        nw181_ = C1()
                        nw181_.ctor__(d_983_v14_)
                        d_989_v19_ = nw181_
                        d_990_v20_: D8
                        d_990_v20_ = D8_DC17(d_974_v6_, d_989_v19_, p0)
                        d_991_v21_: _dafny.Array
                        nw182_ = _dafny.Array(None, 1)
                        nw182_[int(0)] = d_990_v20_
                        d_991_v21_ = nw182_
                        index179_ = default__.safeIndex(198, (d_991_v21_).length(0))
                        (d_991_v21_)[index179_] = d_990_v20_
                        index180_ = default__.safeIndex(198, (d_991_v21_).length(0))
                        (d_991_v21_)[index180_] = d_990_v20_
                    d_992_v22_: C0
                    nw183_ = C0()
                    nw183_.ctor__(d_974_v6_, d_974_v6_)
                    d_992_v22_ = nw183_
                    d_993_v23_: C1
                    nw184_ = C1()
                    nw184_.ctor__(d_983_v14_)
                    d_993_v23_ = nw184_
                    d_994_v24_: D8
                    d_994_v24_ = D8_DC16(d_993_v23_)
                    source19_ = (d_994_v24_ if d_974_v6_ else d_994_v24_)
                    if source19_.is_DC17:
                        d_995___mcc_h0_ = source19_.cf30
                        d_996___mcc_h1_ = source19_.cf31
                        d_997___mcc_h2_ = source19_.cf32
                        d_998_cf32_ = d_997___mcc_h2_
                        d_999_cf31_ = d_996___mcc_h1_
                        d_1000_cf30_ = d_995___mcc_h0_
                        d_1001_v25_: _dafny.Array
                        nw185_ = _dafny.Array(_dafny.Array(None, 0), 7)
                        d_1001_v25_ = nw185_
                        d_1001_v25_ = d_1001_v25_
                        d_1002_v26_: _dafny.Set
                        d_1002_v26_ = _dafny.Set({d_1000_cf30_, not(d_992_v22_.f5), d_992_v22_.f5})
                        d_1003_v27_: D4
                        d_1003_v27_ = D4_DC5(d_1002_v26_)
                        d_1004_v28_: D4
                        d_1004_v28_ = D4_DC7(d_1003_v27_)
                        d_1004_v28_ = D4_DC7(D4_DC5(_dafny.Set({True, False, d_1000_cf30_, d_992_v22_.f4, d_992_v22_.f4})))
                        d_983_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfo"))
                        d_1005_v29_: _dafny.MultiSet
                        d_1005_v29_ = _dafny.MultiSet([d_992_v22_.f5])
                        d_1005_v29_ = ((d_1005_v29_) - (d_1005_v29_)).intersection(d_1005_v29_)
                    elif True:
                        d_1006___mcc_h3_ = source19_.cf29
                        d_1007_cf29_ = d_1006___mcc_h3_
                        index181_ = default__.safeIndex(52, (d_966_v0_).length(0))
                        (d_966_v0_)[index181_] = ((d_966_v0_)[default__.safeIndex(52, (d_966_v0_).length(0))]) - ((0) - ((self.f2) * ((self).fm1(d_992_v22_.f5, globalState))))
                        d_1008_v31_: _dafny.Map
                        d_1008_v31_ = _dafny.Map({(d_993_v23_).f3: (d_966_v0_)[default__.safeIndex(52, (d_966_v0_).length(0))]})
                        d_1009_v32_: _dafny.Map
                        d_1009_v32_ = _dafny.Map({self.f2: d_992_v22_.f4})
                        d_1010_v33_: _dafny.Map
                        d_1010_v33_ = _dafny.Map({(d_993_v23_).f3: d_1009_v32_})
                        d_1011_v34_: _dafny.Seq
                        def iife70_():
                            coll32_ = _dafny.Map()
                            compr_32_: _dafny.Seq
                            for compr_32_ in (d_1008_v31_).keys.Elements:
                                d_1012_v30_: _dafny.Seq = compr_32_
                                if (d_1012_v30_) in (d_1008_v31_):
                                    coll32_[d_1012_v30_] = _dafny.Map({p0: not(d_992_v22_.f5)})
                            return _dafny.Map(coll32_)
                        d_1011_v34_ = _dafny.SeqWithoutIsStrInference([iife70_()
                        , d_1010_v33_])
                        d_1013_v35_: D5
                        d_1013_v35_ = D5_DC10((d_974_v6_) and (False), (0) - (self.f2), (d_1011_v34_)[default__.safeIndex(p0, len(d_1011_v34_))], d_992_v22_.f5)
                        rhs162_ = self.f2
                        rhs163_ = d_1013_v35_
                        r1 = rhs162_
                        d_1013_v35_ = rhs163_
                        d_1014_v36_: C1
                        nw186_ = C1()
                        nw186_.ctor__(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))).set(default__.safeIndex(len(d_976_v8_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))), d_982_v13_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "log"))))
                        d_1014_v36_ = nw186_
                        d_1015_v37_: _dafny.Map
                        d_1015_v37_ = _dafny.Map({d_992_v22_: d_992_v22_.f4})
                        d_1015_v37_ = (d_1015_v37_).set(d_992_v22_, d_992_v22_.f5)
                    pass
            pass
        r0 = not(d_974_v6_)
        d_1016_v38_: _dafny.Seq
        d_1016_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbprgidjf"))
        r1 = len(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1017_i4_ in range(default__.abs(582))])) + (d_1016_v38_) if d_974_v6_ else d_1016_v38_))
        return r0, r1


class C3(T2, T1, T0):
    def  __init__(self):
        self._f2: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f2(self):
        return self._f2
    @f2.setter
    def f2(self, value):
        self._f2 = value
    def ctor__(self, f2):
        (self).f2 = f2

    def fm8(self, p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "haphhjf"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jktmaskl")))) not in ((_dafny.Set({not(False), False})) | (_dafny.Set({not(True)})))

    def fm0(self, globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference([self.f2]) if not(False) else _dafny.SeqWithoutIsStrInference([self.f2])): self.f2})

    def fm1(self, p0, globalState):
        return ((0) - ((0) - ((self.f2 if not(False) else self.f2)))) + (default__.safeModuloInt(self.f2, self.f2))

    def fm10(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tues")) if False else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kchnvrhfl")))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1018_i0_ in range(default__.abs(693))]))

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_1019_v0_: _dafny.Map
        d_1019_v0_ = _dafny.Map({self.f2: p3})
        d_1019_v0_ = d_1019_v0_
        d_1020_v1_: _dafny.Seq
        d_1020_v1_ = _dafny.SeqWithoutIsStrInference([self.f2])
        (self).f2 = len((d_1020_v1_) + (_dafny.SeqWithoutIsStrInference([self.f2 for d_1021_i0_ in range(default__.abs(826))])))
        d_1022_v2_: str
        d_1022_v2_ = _dafny.CodePoint('a')
        d_1023_v3_: D0
        d_1023_v3_ = D0_DC1(True, not(p3), d_1022_v2_, 87)
        d_1024_i1_: int
        d_1024_i1_ = 0
        with _dafny.label("4"):
            while (((self).fm1(p2, globalState)) == (self.f2)) or ((p2 if p2 else (d_1023_v3_).cf1)):
                with _dafny.c_label("4"):
                    if (d_1024_i1_) >= (100):
                        raise _dafny.Break("4")
                    d_1024_i1_ = (d_1024_i1_) + (1)
                    d_1025_v4_: _dafny.Array
                    def lambda32_(d_1026_v2_):
                        def lambda33_(d_1027_i2_):
                            return (_dafny.SeqWithoutIsStrInference([d_1026_v2_ for d_1028_i3_ in range(default__.abs(595))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lap")))

                        return lambda33_

                    init16_ = lambda32_(d_1022_v2_)
                    nw187_ = _dafny.Array(None, 24)
                    for i0_16_ in range(nw187_.length(0)):
                        nw187_[i0_16_] = init16_(i0_16_)
                    d_1025_v4_ = nw187_
                    d_1029_v5_: _dafny.Seq
                    d_1029_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amusabog"))
                    index182_ = default__.safeIndex(174, (d_1025_v4_).length(0))
                    (d_1025_v4_)[index182_] = d_1029_v5_
                    index183_ = default__.safeIndex(174, (d_1025_v4_).length(0))
                    (d_1025_v4_)[index183_] = d_1029_v5_
                    d_1030_v6_: _dafny.Map
                    d_1030_v6_ = _dafny.Map({p1: (0) - ((self).fm1(p2, globalState))})
                    d_1031_v7_: _dafny.Seq
                    d_1031_v7_ = _dafny.SeqWithoutIsStrInference([d_1030_v6_])
                    d_1032_v8_: _dafny.Set
                    d_1032_v8_ = _dafny.Set({p2, p3})
                    d_1033_v9_: _dafny.Map
                    d_1033_v9_ = _dafny.Map({len(d_1032_v8_): d_1030_v6_})
                    d_1034_v10_: _dafny.Array
                    nw188_ = _dafny.Array(None, 19)
                    nw188_[int(0)] = (d_1031_v7_)[default__.safeIndex(self.f2, len(d_1031_v7_))]
                    nw188_[int(1)] = d_1030_v6_
                    nw188_[int(2)] = _dafny.Map({p1: self.f2})
                    nw188_[int(3)] = d_1030_v6_
                    nw188_[int(4)] = d_1030_v6_
                    nw188_[int(5)] = d_1030_v6_
                    nw188_[int(6)] = (d_1030_v6_) | (_dafny.Map({p1: self.f2}))
                    nw188_[int(7)] = (d_1030_v6_).set(p1, self.f2)
                    nw188_[int(8)] = (d_1030_v6_) | (d_1030_v6_)
                    nw188_[int(9)] = (d_1030_v6_) | (d_1030_v6_)
                    nw188_[int(10)] = (d_1030_v6_) | (_dafny.Map({p1: self.f2}))
                    nw188_[int(11)] = d_1030_v6_
                    nw188_[int(12)] = ((d_1030_v6_).set(p1, len((d_1025_v4_)[default__.safeIndex(174, (d_1025_v4_).length(0))]))) | (_dafny.Map({p1: len(d_1029_v5_)}))
                    nw188_[int(13)] = _dafny.Map({p1: self.f2})
                    nw188_[int(14)] = (d_1030_v6_) | (d_1030_v6_)
                    nw188_[int(15)] = d_1030_v6_
                    nw188_[int(16)] = d_1030_v6_
                    nw188_[int(17)] = d_1030_v6_
                    nw188_[int(18)] = ((d_1033_v9_)[self.f2] if (self.f2) in (d_1033_v9_) else d_1030_v6_)
                    d_1034_v10_ = nw188_
                    index184_ = default__.safeIndex(473, (d_1034_v10_).length(0))
                    (d_1034_v10_)[index184_] = d_1030_v6_
                    index185_ = default__.safeIndex(473, (d_1034_v10_).length(0))
                    (d_1034_v10_)[index185_] = d_1030_v6_
                    r0 = len((d_1029_v5_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))))
                    d_1035_v11_: _dafny.Map
                    d_1035_v11_ = _dafny.Map({self.f2: self.f2})
                    index186_ = default__.safeIndex(699, (p1).length(0))
                    (p1)[index186_] = (d_1020_v1_)[default__.safeIndex(len(d_1035_v11_), len(d_1020_v1_))]
                    index187_ = default__.safeIndex(699, (p1).length(0))
                    (p1)[index187_] = self.f2
                    pass
            pass
        d_1036_v12_: _dafny.Set
        d_1036_v12_ = _dafny.Set({self.f2})
        (self).f2 = len(d_1036_v12_)
        d_1037_v13_: bool
        d_1037_v13_ = False
        d_1038_v14_: _dafny.Seq
        d_1038_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unuhbnudi"))
        d_1039_v15_: _dafny.Array
        nw189_ = _dafny.Array(_dafny.Array(None, 0), 19)
        d_1039_v15_ = nw189_
        d_1040_v16_: _dafny.Set
        d_1040_v16_ = _dafny.Set({_dafny.Set({p3})})
        rhs164_ = d_1037_v13_
        rhs165_ = (d_1040_v16_).issubset(d_1040_v16_)
        rhs166_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlb"))
        rhs167_ = d_1039_v15_
        d_1037_v13_ = rhs164_
        d_1037_v13_ = rhs165_
        d_1038_v14_ = rhs166_
        d_1039_v15_ = rhs167_
        d_1041_v17_: _dafny.Array
        nw190_ = _dafny.Array(False, 25)
        d_1041_v17_ = nw190_
        index188_ = default__.safeIndex(359, (d_1041_v17_).length(0))
        (d_1041_v17_)[index188_] = d_1037_v13_
        d_1042_v18_: D0
        d_1042_v18_ = D0_DC0(p2)
        pat_let_tv22_ = p2
        pat_let_tv23_ = d_1038_v14_
        pat_let_tv24_ = d_1042_v18_
        pat_let_tv25_ = globalState
        d_1043_v19_: _dafny.MultiSet
        def iife72_(_pat_let20_0):
            def iife73_(d_1044_dt__update__tmp_h0_):
                def iife74_(_pat_let21_0):
                    def iife75_(d_1045_dt__update_hcf0_h0_):
                        return D0_DC0(d_1045_dt__update_hcf0_h0_)
                    return iife75_(_pat_let21_0)
                return iife74_(pat_let_tv22_)
            return iife73_(_pat_let20_0)
        def iife71_(_pat_let19_0):
            def iife76_(d_1046_dt__update__tmp_h1_):
                def iife77_(_pat_let22_0):
                    def iife78_(d_1047_dt__update_hcf0_h1_):
                        return D0_DC0(d_1047_dt__update_hcf0_h1_)
                    return iife78_(_pat_let22_0)
                return iife77_((self).fm8(pat_let_tv23_, (pat_let_tv24_).cf0, pat_let_tv25_))
            return iife76_(_pat_let19_0)
        d_1043_v19_ = _dafny.MultiSet([(iife71_(iife72_(d_1042_v18_))).cf0])
        index189_ = default__.safeIndex(359, (d_1041_v17_).length(0))
        (d_1041_v17_)[index189_] = (False) not in (d_1043_v19_)
        r0 = (0) - ((self.f2) * (self.f2))
        d_1048_v20_: _dafny.Seq
        d_1048_v20_ = _dafny.SeqWithoutIsStrInference([True])
        r1 = (self).fm1((d_1048_v20_)[default__.safeIndex(len((d_1038_v14_).set(default__.safeIndex(self.f2, len(d_1038_v14_)), d_1022_v2_)), len(d_1048_v20_))], globalState)
        return r0, r1

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        if p2:
            if p2:
                d_1049_v0_: _dafny.MultiSet
                d_1049_v0_ = _dafny.MultiSet([-215, (self.f2) * (446), (p0) - (p0), p0])
                d_1050_v1_: _dafny.Seq
                d_1050_v1_ = _dafny.SeqWithoutIsStrInference([d_1049_v0_, _dafny.MultiSet([p0, p0])])
                d_1049_v0_ = (d_1049_v0_).intersection((d_1050_v1_)[default__.safeIndex(self.f2, len(d_1050_v1_))])
                (self).f2 = self.f2
                d_1051_v2_: _dafny.Set
                d_1051_v2_ = _dafny.Set({p0})
                d_1052_v3_: _dafny.Map
                d_1052_v3_ = _dafny.Map({d_1051_v2_: p0})
                d_1052_v3_ = (d_1052_v3_).set(d_1051_v2_, self.f2)
                d_1053_v4_: _dafny.Map
                d_1053_v4_ = _dafny.Map({not(p3): p2})
                d_1054_v5_: _dafny.Array
                nw191_ = _dafny.Array(None, 6)
                nw191_[int(0)] = (((d_1053_v4_)[p3] if (p3) in (d_1053_v4_) else p3)) or (p2)
                nw191_[int(1)] = (p0) == (p0)
                nw191_[int(2)] = True
                nw191_[int(3)] = (p0) == (p0)
                nw191_[int(4)] = p3
                nw191_[int(5)] = not(p2)
                d_1054_v5_ = nw191_
                index190_ = default__.safeIndex(288, (d_1054_v5_).length(0))
                (d_1054_v5_)[index190_] = p2
                d_1055_v6_: _dafny.Seq
                d_1055_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecarnc"))
                index191_ = default__.safeIndex(288, (d_1054_v5_).length(0))
                (d_1054_v5_)[index191_] = (self).fm8(d_1055_v6_, p3, globalState)
                d_1056_v7_: _dafny.Map
                d_1056_v7_ = _dafny.Map({d_1055_v6_: self.f2})
                d_1056_v7_ = (d_1056_v7_).set(d_1055_v6_, default__.safeDivisionInt(p0, p0))
            elif True:
                r0 = p3
                d_1057_v8_: _dafny.Array
                nw192_ = _dafny.Array(int(0), 12)
                d_1057_v8_ = nw192_
                d_1058_v9_: _dafny.Array
                nw193_ = _dafny.Array(None, 5)
                nw193_[int(0)] = d_1057_v8_
                nw193_[int(1)] = d_1057_v8_
                nw193_[int(2)] = d_1057_v8_
                nw193_[int(3)] = d_1057_v8_
                nw193_[int(4)] = d_1057_v8_
                d_1058_v9_ = nw193_
                index192_ = default__.safeIndex(619, (d_1058_v9_).length(0))
                (d_1058_v9_)[index192_] = d_1057_v8_
                index193_ = default__.safeIndex(619, (d_1058_v9_).length(0))
                (d_1058_v9_)[index193_] = d_1057_v8_
                r2 = ((self.f2 if p3 else p0)) * (default__.safeModuloInt(p0, self.f2))
                d_1059_v10_: _dafny.Seq
                d_1059_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
                r0 = (self).fm8(d_1059_v10_, False, globalState)
                d_1060_v11_: str
                d_1060_v11_ = _dafny.CodePoint('n')
                d_1060_v11_ = d_1060_v11_
            d_1061_v12_: _dafny.Array
            def lambda34_(d_1062_i0_):
                return (d_1062_i0_) * (self.f2)

            init17_ = lambda34_
            nw194_ = _dafny.Array(None, 25)
            for i0_17_ in range(nw194_.length(0)):
                nw194_[i0_17_] = init17_(i0_17_)
            d_1061_v12_ = nw194_
            d_1061_v12_ = d_1061_v12_
            if p2:
                d_1063_v13_: _dafny.Array
                nw195_ = _dafny.Array(False, 1)
                d_1063_v13_ = nw195_
                d_1064_v14_: _dafny.MultiSet
                d_1064_v14_ = _dafny.MultiSet([p0])
                index194_ = default__.safeIndex(666, (d_1063_v13_).length(0))
                (d_1063_v13_)[index194_] = (d_1064_v14_).isdisjoint(d_1064_v14_)
                index195_ = default__.safeIndex(666, (d_1063_v13_).length(0))
                (d_1063_v13_)[index195_] = p3
                d_1065_v15_: _dafny.MultiSet
                d_1065_v15_ = _dafny.MultiSet([p3])
                d_1066_v16_: _dafny.Array
                nw196_ = _dafny.Array(_dafny.CodePoint('D'), 18)
                d_1066_v16_ = nw196_
                d_1067_v17_: _dafny.Map
                d_1067_v17_ = _dafny.Map({(d_1065_v15_) | (d_1065_v15_): d_1066_v16_})
                d_1067_v17_ = d_1067_v17_
                d_1068_v18_: _dafny.Seq
                d_1068_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgfihkea"))
                index196_ = default__.safeIndex(666, (d_1063_v13_).length(0))
                (d_1063_v13_)[index196_] = (self).fm8((d_1068_v18_) + (d_1068_v18_), (d_1063_v13_)[default__.safeIndex(666, (d_1063_v13_).length(0))], globalState)
                d_1069_v19_: C0
                nw197_ = C0()
                nw197_.ctor__(p2, not((d_1063_v13_)[default__.safeIndex(666, (d_1063_v13_).length(0))]))
                d_1069_v19_ = nw197_
                index197_ = default__.safeIndex(666, (d_1063_v13_).length(0))
                (d_1063_v13_)[index197_] = not(not(not (p3) or (p2)))
            elif True:
                d_1070_v20_: _dafny.Seq
                d_1070_v20_ = _dafny.SeqWithoutIsStrInference([p3])
                d_1071_v21_: D0
                d_1071_v21_ = D0_DC1(p3, p2, _dafny.CodePoint('u'), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tplkvlen"))))
                d_1072_v22_: str
                d_1072_v22_ = _dafny.CodePoint('h')
                pat_let_tv26_ = d_1072_v22_
                pat_let_tv27_ = p2
                d_1073_v23_: _dafny.Set
                def iife79_(_pat_let23_0):
                    def iife80_(d_1074_dt__update__tmp_h0_):
                        def iife81_(_pat_let24_0):
                            def iife82_(d_1075_dt__update_hcf3_h0_):
                                def iife83_(_pat_let25_0):
                                    def iife84_(d_1076_dt__update_hcf2_h0_):
                                        return D0_DC1((d_1074_dt__update__tmp_h0_).cf1, d_1076_dt__update_hcf2_h0_, d_1075_dt__update_hcf3_h0_, (d_1074_dt__update__tmp_h0_).cf4)
                                    return iife84_(_pat_let25_0)
                                return iife83_(pat_let_tv27_)
                            return iife82_(_pat_let24_0)
                        return iife81_(pat_let_tv26_)
                    return iife80_(_pat_let23_0)
                d_1073_v23_ = _dafny.Set({iife79_(d_1071_v21_), d_1071_v21_, d_1071_v21_})
                d_1077_v24_: _dafny.Array
                nw198_ = _dafny.Array(None, 14)
                nw198_[int(0)] = not(p2)
                nw198_[int(1)] = (175) > (p0)
                nw198_[int(2)] = (self.f2) == (self.f2)
                nw198_[int(3)] = p2
                nw198_[int(4)] = p3
                nw198_[int(5)] = p3
                nw198_[int(6)] = p3
                nw198_[int(7)] = (d_1070_v20_)[default__.safeIndex(p0, len(d_1070_v20_))]
                nw198_[int(8)] = p2
                nw198_[int(9)] = p3
                nw198_[int(10)] = p3
                nw198_[int(11)] = (_dafny.Set({d_1071_v21_, d_1071_v21_})).issubset(d_1073_v23_)
                nw198_[int(12)] = p2
                nw198_[int(13)] = p2
                d_1077_v24_ = nw198_
                index198_ = default__.safeIndex(378, (d_1077_v24_).length(0))
                (d_1077_v24_)[index198_] = not(not((self).fm8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihkkeqm")), p2, globalState)))
                index199_ = default__.safeIndex(378, (d_1077_v24_).length(0))
                (d_1077_v24_)[index199_] = p3
                d_1078_v25_: _dafny.Seq
                d_1078_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwutipmey"))
                d_1070_v20_ = _dafny.SeqWithoutIsStrInference([(len(d_1078_v25_)) >= (p0), p3, not(p3)])
                d_1079_v26_: C0
                nw199_ = C0()
                nw199_.ctor__(p3, p2)
                d_1079_v26_ = nw199_
                d_1080_v27_: _dafny.Map
                d_1080_v27_ = _dafny.Map({self.f2: self.f2})
                d_1081_v28_: _dafny.Set
                d_1081_v28_ = _dafny.Set({p0})
                d_1082_v29_: _dafny.Map
                d_1082_v29_ = _dafny.Map({self.f2: d_1081_v28_})
                d_1083_v30_: _dafny.Map
                d_1083_v30_ = _dafny.Map({p0: (d_1078_v25_).set(default__.safeIndex(len(d_1082_v29_), len(d_1078_v25_)), d_1072_v22_)})
                d_1080_v27_ = (d_1080_v27_).set((p0) * (self.f2), len(d_1083_v30_))
                r2 = ((p0) - (len(_dafny.SeqWithoutIsStrInference([p0 for d_1084_i1_ in range(default__.abs(7))]))) if d_1079_v26_.f4 else self.f2)
            d_1085_v31_: _dafny.Map
            d_1085_v31_ = _dafny.Map({p2: False})
            r0 = ((d_1085_v31_)[p3] if (p3) in (d_1085_v31_) else p2)
            d_1086_v32_: _dafny.Array
            def lambda35_(d_1087_v31_):
                def lambda36_(d_1088_i2_):
                    return d_1087_v31_

                return lambda36_

            init18_ = lambda35_(d_1085_v31_)
            nw200_ = _dafny.Array(None, 29)
            for i0_18_ in range(nw200_.length(0)):
                nw200_[i0_18_] = init18_(i0_18_)
            d_1086_v32_ = nw200_
            d_1089_v33_: _dafny.Map
            d_1089_v33_ = d_1085_v31_
            index200_ = default__.safeIndex(133, (d_1086_v32_).length(0))
            (d_1086_v32_)[index200_] = d_1089_v33_
            index201_ = default__.safeIndex(133, (d_1086_v32_).length(0))
            (d_1086_v32_)[index201_] = d_1089_v33_
        elif True:
            r2 = self.f2
            d_1090_v34_: _dafny.Seq
            d_1090_v34_ = _dafny.SeqWithoutIsStrInference([p3, p3, p2, p2])
            d_1091_v35_: _dafny.Seq
            d_1091_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdbarimii"))
            d_1092_v36_: D6
            d_1092_v36_ = D6_DC11(d_1091_v35_)
            r2 = ((121) - (len((d_1092_v36_).cf21)) if not((d_1090_v34_)[default__.safeIndex(self.f2, len(d_1090_v34_))]) else self.f2)
            d_1093_v37_: _dafny.Set
            d_1093_v37_ = _dafny.Set({self.f2})
            d_1094_v38_: _dafny.Map
            d_1094_v38_ = _dafny.Map({(d_1093_v37_) - (_dafny.Set({511, len(_dafny.Map({not(p3): True})), (_dafny.MultiSet([p0])).cardinality})): p2})
            d_1094_v38_ = d_1094_v38_
            d_1095_v39_: C0
            nw201_ = C0()
            nw201_.ctor__(False, p2)
            d_1095_v39_ = nw201_
            r0 = p3
        d_1096_v40_: _dafny.Seq
        d_1096_v40_ = _dafny.SeqWithoutIsStrInference([p3])
        if (len((d_1096_v40_) + (d_1096_v40_))) == (p0):
            d_1097_v42_: _dafny.Map
            d_1097_v42_ = _dafny.Map({p3: p3})
            d_1098_v43_: _dafny.Seq
            d_1098_v43_ = _dafny.SeqWithoutIsStrInference([d_1097_v42_, d_1097_v42_])
            d_1099_v44_: _dafny.Seq
            d_1099_v44_ = _dafny.SeqWithoutIsStrInference([d_1098_v43_, d_1098_v43_, d_1098_v43_])
            d_1100_v45_: _dafny.Map
            d_1100_v45_ = _dafny.Map({_dafny.Map({p3: p2}): p2})
            d_1101_v46_: _dafny.Set
            d_1101_v46_ = _dafny.Set({self.f2, self.f2})
            d_1102_v47_: _dafny.Map
            d_1102_v47_ = _dafny.Map({d_1101_v46_: d_1100_v45_})
            d_1103_v48_: _dafny.MultiSet
            d_1103_v48_ = _dafny.MultiSet([False, p2])
            d_1104_v51_: _dafny.Seq
            d_1104_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpieicl"))
            d_1105_v52_: _dafny.Map
            d_1105_v52_ = _dafny.Map({d_1097_v42_: d_1104_v51_})
            d_1106_v54_: _dafny.Array
            nw202_ = _dafny.Array(None, 24)
            def iife85_():
                coll33_ = _dafny.Map()
                compr_33_: _dafny.Map
                for compr_33_ in ((d_1099_v44_)[default__.safeIndex(p0, len(d_1099_v44_))]).Elements:
                    d_1107_v41_: _dafny.Map = compr_33_
                    if (d_1107_v41_) in ((d_1099_v44_)[default__.safeIndex(p0, len(d_1099_v44_))]):
                        coll33_[d_1107_v41_] = p2
                return _dafny.Map(coll33_)
            nw202_[int(0)] = iife85_()
            
            nw202_[int(1)] = (d_1100_v45_).set(d_1097_v42_, p2)
            nw202_[int(2)] = _dafny.Map({d_1097_v42_: (d_1096_v40_)[default__.safeIndex(p0, len(d_1096_v40_))]})
            nw202_[int(3)] = (d_1100_v45_).set(d_1097_v42_, p3)
            nw202_[int(4)] = d_1100_v45_
            def iife86_():
                coll34_ = _dafny.Map()
                compr_34_: _dafny.Map
                for compr_34_ in (d_1098_v43_).Elements:
                    d_1108_v49_: _dafny.Map = compr_34_
                    if (d_1108_v49_) in (d_1098_v43_):
                        coll34_[d_1108_v49_] = p3
                return _dafny.Map(coll34_)
            nw202_[int(5)] = ((d_1102_v47_)[default__.fm18(p2, p2, d_1103_v48_, globalState)] if (default__.fm18(p2, p2, d_1103_v48_, globalState)) in (d_1102_v47_) else iife86_()
            )
            nw202_[int(6)] = d_1100_v45_
            nw202_[int(7)] = d_1100_v45_
            nw202_[int(8)] = _dafny.Map({d_1097_v42_: p2})
            nw202_[int(9)] = d_1100_v45_
            nw202_[int(10)] = (d_1100_v45_).set(d_1097_v42_, p3)
            nw202_[int(11)] = d_1100_v45_
            nw202_[int(12)] = (d_1100_v45_) | (default__.fm20(globalState))
            def iife87_():
                coll35_ = _dafny.Map()
                compr_35_: _dafny.Map
                for compr_35_ in (d_1105_v52_).keys.Elements:
                    d_1109_v50_: _dafny.Map = compr_35_
                    if (d_1109_v50_) in (d_1105_v52_):
                        coll35_[d_1109_v50_] = p3
                return _dafny.Map(coll35_)
            nw202_[int(13)] = iife87_()
            
            nw202_[int(14)] = _dafny.Map({d_1097_v42_: p3})
            def iife88_():
                coll36_ = _dafny.Map()
                compr_36_: _dafny.Map
                for compr_36_ in (d_1100_v45_).keys.Elements:
                    d_1110_v53_: _dafny.Map = compr_36_
                    if (d_1110_v53_) in (d_1100_v45_):
                        coll36_[d_1110_v53_] = p3
                return _dafny.Map(coll36_)
            nw202_[int(15)] = (default__.fm20(globalState)) | (iife88_()
            )
            nw202_[int(16)] = d_1100_v45_
            nw202_[int(17)] = d_1100_v45_
            nw202_[int(18)] = (d_1100_v45_) | ((d_1100_v45_).set(d_1097_v42_, p3))
            nw202_[int(19)] = d_1100_v45_
            nw202_[int(20)] = _dafny.Map({d_1097_v42_: True})
            nw202_[int(21)] = _dafny.Map({d_1097_v42_: ((d_1097_v42_)[p3] if (p3) in (d_1097_v42_) else p3)})
            nw202_[int(22)] = (d_1100_v45_).set(d_1097_v42_, p2)
            nw202_[int(23)] = d_1100_v45_
            d_1106_v54_ = nw202_
            index202_ = default__.safeIndex(789, (d_1106_v54_).length(0))
            (d_1106_v54_)[index202_] = _dafny.Map({d_1097_v42_: (d_1096_v40_)[default__.safeIndex(len(d_1104_v51_), len(d_1096_v40_))]})
            d_1111_v56_: _dafny.Set
            d_1111_v56_ = _dafny.Set({d_1097_v42_, default__.fm21(globalState)})
            index203_ = default__.safeIndex(789, (d_1106_v54_).length(0))
            def iife89_():
                coll37_ = _dafny.Map()
                compr_37_: _dafny.Map
                for compr_37_ in ((d_1111_v56_).intersection(d_1111_v56_)).Elements:
                    d_1112_v55_: _dafny.Map = compr_37_
                    if (d_1112_v55_) in ((d_1111_v56_).intersection(d_1111_v56_)):
                        coll37_[d_1112_v55_] = ((d_1097_v42_)[p3] if (p3) in (d_1097_v42_) else p3)
                return _dafny.Map(coll37_)
            (d_1106_v54_)[index203_] = iife89_()
            
            if p3:
                (self).f2 = p0
                d_1113_v57_: _dafny.Array
                nw203_ = _dafny.Array(False, 7)
                d_1113_v57_ = nw203_
                d_1114_v58_: _dafny.Seq
                d_1114_v58_ = _dafny.SeqWithoutIsStrInference([d_1113_v57_, d_1113_v57_])
                r0 = (d_1114_v58_) <= (d_1114_v58_)
                r0 = (p0) != ((p0) + (p0))
                index204_ = default__.safeIndex(317, (d_1113_v57_).length(0))
                (d_1113_v57_)[index204_] = (not(p3)) == (p3)
                index205_ = default__.safeIndex(317, (d_1113_v57_).length(0))
                (d_1113_v57_)[index205_] = p2
                d_1115_v59_: _dafny.Map
                d_1115_v59_ = _dafny.Map({True: p0})
                d_1116_v60_: _dafny.Map
                d_1116_v60_ = _dafny.Map({p2: default__.fm15(globalState)})
                d_1117_v61_: str
                d_1117_v61_ = _dafny.CodePoint('t')
                d_1115_v59_ = _dafny.Map({(self.f2) < (len(d_1116_v60_)): len((d_1104_v51_) + ((d_1104_v51_).set(default__.safeIndex((0) - (self.f2), len(d_1104_v51_)), d_1117_v61_)))})
            elif True:
                d_1118_v62_: _dafny.Array
                def lambda37_(d_1119_v51_):
                    def lambda38_(d_1120_i3_):
                        return (_dafny.MultiSet([d_1119_v51_])).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1121_i4_ in range(default__.abs(31))])]))

                    return lambda38_

                init19_ = lambda37_(d_1104_v51_)
                nw204_ = _dafny.Array(None, 13)
                for i0_19_ in range(nw204_.length(0)):
                    nw204_[i0_19_] = init19_(i0_19_)
                d_1118_v62_ = nw204_
                d_1122_v63_: _dafny.MultiSet
                d_1122_v63_ = _dafny.MultiSet([d_1104_v51_, d_1104_v51_, d_1104_v51_])
                index206_ = default__.safeIndex(359, (d_1118_v62_).length(0))
                (d_1118_v62_)[index206_] = (default__.fm22(p2, self.f2, p2, p0, globalState)) - (d_1122_v63_)
                d_1123_v64_: _dafny.Array
                nw205_ = _dafny.Array(int(0), 16)
                d_1123_v64_ = nw205_
                d_1124_v65_: _dafny.Map
                d_1124_v65_ = _dafny.Map({d_1123_v64_: (d_1122_v63_ if p2 else d_1122_v63_)})
                index207_ = default__.safeIndex(359, (d_1118_v62_).length(0))
                (d_1118_v62_)[index207_] = ((d_1124_v65_)[d_1123_v64_] if (d_1123_v64_) in (d_1124_v65_) else d_1122_v63_)
                d_1125_v66_: C1
                nw206_ = C1()
                nw206_.ctor__((d_1104_v51_) + (d_1104_v51_))
                d_1125_v66_ = nw206_
                d_1126_v67_: str
                d_1126_v67_ = _dafny.CodePoint('q')
                d_1127_v68_: _dafny.Map
                d_1127_v68_ = _dafny.Map({p0: p2})
                d_1128_v69_: D5
                d_1128_v69_ = D5_DC9(False, p0, d_1103_v48_)
                d_1129_v70_: D6
                d_1129_v70_ = D6_DC13(p3, p0)
                d_1130_v71_: _dafny.Seq
                d_1130_v71_ = _dafny.SeqWithoutIsStrInference([d_1129_v70_, d_1129_v70_])
                d_1131_v72_: _dafny.Seq
                d_1131_v72_ = _dafny.SeqWithoutIsStrInference([self.f2, len(d_1130_v71_)])
                d_1132_v73_: D4
                d_1132_v73_ = D4_DC6((d_1125_v66_).fm1(p2, globalState), d_1123_v64_, len(d_1131_v72_))
                d_1133_v74_: _dafny.Array
                nw207_ = _dafny.Array(None, 24)
                nw207_[int(0)] = p3
                nw207_[int(1)] = p2
                nw207_[int(2)] = p2
                nw207_[int(3)] = (d_1126_v67_) in ((_dafny.SeqWithoutIsStrInference([d_1126_v67_ for d_1134_i5_ in range(default__.abs(66))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([d_1126_v67_ for d_1135_i5_ in range(default__.abs(66))]))), d_1126_v67_))
                nw207_[int(4)] = p3
                nw207_[int(5)] = p2
                nw207_[int(6)] = ((d_1127_v68_)[self.f2] if (self.f2) in (d_1127_v68_) else not(True))
                nw207_[int(7)] = (d_1125_v66_).fm11(globalState)
                nw207_[int(8)] = p2
                nw207_[int(9)] = (self.f2) <= (self.f2)
                nw207_[int(10)] = p2
                nw207_[int(11)] = (d_1128_v69_).cf14
                nw207_[int(12)] = p3
                nw207_[int(13)] = (p0) >= (p0)
                nw207_[int(14)] = p2
                nw207_[int(15)] = (not(p3) if p3 else p2)
                nw207_[int(16)] = p2
                nw207_[int(17)] = (d_1096_v40_)[default__.safeIndex((d_1132_v73_).cf11, len(d_1096_v40_))]
                nw207_[int(18)] = p3
                nw207_[int(19)] = p3
                nw207_[int(20)] = not((_dafny.SeqWithoutIsStrInference([p0, self.f2])) < (d_1131_v72_))
                nw207_[int(21)] = p2
                nw207_[int(22)] = True
                nw207_[int(23)] = p3
                d_1133_v74_ = nw207_
                index208_ = default__.safeIndex(985, (d_1133_v74_).length(0))
                (d_1133_v74_)[index208_] = p3
                index209_ = default__.safeIndex(985, (d_1133_v74_).length(0))
                (d_1133_v74_)[index209_] = p3
                d_1096_v40_ = d_1096_v40_
                d_1123_v64_ = d_1123_v64_
            d_1136_v75_: str
            d_1136_v75_ = _dafny.CodePoint('m')
            d_1136_v75_ = d_1136_v75_
            r0 = p2
            d_1137_v76_: _dafny.Map
            d_1137_v76_ = _dafny.Map({879: default__.fm19(self.f2, p0, False, self.f2, globalState)})
            d_1138_v77_: C0
            nw208_ = C0()
            nw208_.ctor__((-623) != (len(d_1137_v76_)), p2)
            d_1138_v77_ = nw208_
        elif True:
            r0 = p2
            d_1139_v78_: str
            d_1139_v78_ = _dafny.CodePoint('r')
            d_1140_v79_: _dafny.Seq
            d_1140_v79_ = _dafny.SeqWithoutIsStrInference([d_1139_v78_])
            d_1141_v80_: _dafny.Array
            nw209_ = _dafny.Array(None, 10)
            nw209_[int(0)] = d_1139_v78_
            nw209_[int(1)] = default__.fm23(globalState)
            nw209_[int(2)] = (d_1140_v79_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wa"))), len(d_1140_v79_))]
            nw209_[int(3)] = _dafny.CodePoint('h')
            nw209_[int(4)] = (d_1139_v78_ if p2 else d_1139_v78_)
            nw209_[int(5)] = d_1139_v78_
            nw209_[int(6)] = d_1139_v78_
            nw209_[int(7)] = d_1139_v78_
            nw209_[int(8)] = d_1139_v78_
            nw209_[int(9)] = d_1139_v78_
            d_1141_v80_ = nw209_
            d_1142_v81_: _dafny.Set
            d_1142_v81_ = _dafny.Set({p0, p0, len(d_1140_v79_), p0})
            rhs168_ = d_1139_v78_
            rhs169_ = d_1141_v80_
            rhs170_ = d_1142_v81_
            d_1139_v78_ = rhs168_
            d_1141_v80_ = rhs169_
            d_1142_v81_ = rhs170_
            d_1143_v82_: D0
            d_1143_v82_ = D0_DC0(p2)
            r0 = (d_1143_v82_).cf0
            if (p0) < ((self.f2 if p3 else self.f2)):
                d_1144_v83_: C1
                nw210_ = C1()
                nw210_.ctor__(d_1140_v79_)
                d_1144_v83_ = nw210_
                d_1145_v84_: _dafny.Array
                nw211_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_1145_v84_ = nw211_
                d_1146_v85_: _dafny.Array
                d_1146_v85_ = d_1145_v84_
                d_1146_v85_ = d_1146_v85_
                d_1147_v86_: D6
                d_1147_v86_ = D6_DC11((d_1140_v79_).set(default__.safeIndex(self.f2, len(d_1140_v79_)), d_1139_v78_))
                d_1148_v87_: D6
                d_1148_v87_ = D6_DC11((d_1147_v86_).cf21)
                d_1149_v88_: _dafny.Seq
                d_1149_v88_ = _dafny.SeqWithoutIsStrInference([p0])
                rhs171_ = (self).fm8(d_1140_v79_, p2, globalState)
                rhs172_ = d_1148_v87_
                rhs173_ = default__.safeDivisionInt((d_1149_v88_)[default__.safeIndex((d_1144_v83_).fm1(p2, globalState), len(d_1149_v88_))], default__.safeDivisionInt(p0, self.f2))
                lhs96_ = self
                r0 = rhs171_
                d_1148_v87_ = rhs172_
                lhs96_.f2 = rhs173_
                (self).f2 = len(((self).fm10(default__.safeDivisionInt(len(((d_1144_v83_).f3).set(default__.safeIndex(p0, len((d_1144_v83_).f3)), _dafny.CodePoint('r'))), -521), globalState)).set(default__.safeIndex(len((d_1144_v83_).f3), len((self).fm10(default__.safeDivisionInt(len(((d_1144_v83_).f3).set(default__.safeIndex(p0, len((d_1144_v83_).f3)), _dafny.CodePoint('r'))), -521), globalState))), d_1139_v78_))
                d_1150_v89_: _dafny.Map
                d_1150_v89_ = _dafny.Map({(d_1144_v83_).fm1(True, globalState): 953})
                d_1151_v90_: _dafny.Map
                d_1151_v90_ = _dafny.Map({p0: not(p2)})
                d_1152_v91_: _dafny.Array
                nw212_ = _dafny.Array(None, 24)
                nw212_[int(0)] = p0
                nw212_[int(1)] = (p0) - (p0)
                nw212_[int(2)] = ((d_1150_v89_)[(0) - (len(_dafny.SeqWithoutIsStrInference([p0])))] if ((0) - (len(_dafny.SeqWithoutIsStrInference([p0])))) in (d_1150_v89_) else 593)
                nw212_[int(3)] = self.f2
                nw212_[int(4)] = p0
                nw212_[int(5)] = p0
                nw212_[int(6)] = p0
                nw212_[int(7)] = (0) - (p0)
                nw212_[int(8)] = (d_1144_v83_).fm1(p3, globalState)
                nw212_[int(9)] = p0
                nw212_[int(10)] = (p0) + (144)
                nw212_[int(11)] = (self.f2) - (p0)
                nw212_[int(12)] = (self.f2) * (self.f2)
                nw212_[int(13)] = self.f2
                nw212_[int(14)] = p0
                nw212_[int(15)] = p0
                nw212_[int(16)] = p0
                nw212_[int(17)] = len(d_1140_v79_)
                nw212_[int(18)] = len(d_1151_v90_)
                nw212_[int(19)] = p0
                nw212_[int(20)] = (self).fm1(p3, globalState)
                nw212_[int(21)] = self.f2
                nw212_[int(22)] = self.f2
                nw212_[int(23)] = p0
                d_1152_v91_ = nw212_
                index210_ = default__.safeIndex(938, (d_1152_v91_).length(0))
                (d_1152_v91_)[index210_] = self.f2
                index211_ = default__.safeIndex(938, (d_1152_v91_).length(0))
                rhs174_ = ((self).fm1(p2, globalState)) * (default__.safeModuloInt(self.f2, len(_dafny.SeqWithoutIsStrInference([d_1139_v78_, d_1139_v78_]))))
                rhs175_ = (d_1149_v88_)[default__.safeIndex(p0, len(d_1149_v88_))]
                rhs176_ = ((d_1140_v79_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cpr")))) + (_dafny.SeqWithoutIsStrInference([d_1139_v78_ for d_1153_i6_ in range(default__.abs(451))]))
                rhs177_ = ((529) + (p0)) * (458)
                lhs97_ = d_1152_v91_
                lhs98_ = default__.safeIndex(938, (d_1152_v91_).length(0))
                r2 = rhs174_
                r2 = rhs175_
                d_1140_v79_ = rhs176_
                lhs97_[lhs98_] = rhs177_
            elif True:
                d_1154_v92_: _dafny.MultiSet
                d_1154_v92_ = _dafny.MultiSet([(p3) and (p3)])
                (self).f2 = (d_1154_v92_).cardinality
                d_1140_v79_ = _dafny.SeqWithoutIsStrInference([d_1139_v78_ for d_1155_i7_ in range(default__.abs(745))])
                d_1154_v92_ = (d_1154_v92_) - (d_1154_v92_)
                d_1156_v93_: _dafny.Array
                nw213_ = _dafny.Array(int(0), 10)
                d_1156_v93_ = nw213_
                index212_ = default__.safeIndex(57, (d_1156_v93_).length(0))
                (d_1156_v93_)[index212_] = len(d_1140_v79_)
                index213_ = default__.safeIndex(57, (d_1156_v93_).length(0))
                (d_1156_v93_)[index213_] = p0
                d_1157_v94_: C1
                nw214_ = C1()
                nw214_.ctor__(_dafny.SeqWithoutIsStrInference([d_1139_v78_ for d_1158_i8_ in range(default__.abs(995))]))
                d_1157_v94_ = nw214_
            d_1159_v95_: _dafny.Map
            d_1159_v95_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0 for d_1160_i9_ in range(default__.abs(893))]): p0})
            d_1161_v96_: _dafny.Map
            d_1161_v96_ = d_1159_v95_
            d_1162_v97_: _dafny.Map
            d_1162_v97_ = _dafny.Map({d_1161_v96_: 30})
            r0 = not(((p0) + (p0)) < (((d_1162_v97_)[d_1161_v96_] if (d_1161_v96_) in (d_1162_v97_) else p0)))
        r2 = (0) - (self.f2)
        d_1163_v98_: _dafny.Seq
        d_1163_v98_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brhukyi"))
        hi4_ = (self.f2) + (self.f2)
        for d_1164_i10_ in range(default__.safeModuloInt(self.f2, len(d_1163_v98_)), hi4_):
            if not(p3):
                d_1165_v99_: T2
                nw215_ = C2()
                nw215_.ctor__((self).fm1(p2, globalState))
                d_1165_v99_ = nw215_
                d_1166_v100_: T2
                d_1166_v100_ = d_1165_v99_
                rhs178_ = default__.safeDivisionInt(p0, self.f2)
                rhs179_ = (d_1166_v100_)
                rhs180_ = (self).fm8(d_1163_v98_, p3, globalState)
                r2 = rhs178_
                d_1165_v99_ = rhs179_
                r0 = rhs180_
                d_1167_v101_: str
                d_1167_v101_ = _dafny.CodePoint('d')
                d_1168_v102_: _dafny.MultiSet
                d_1168_v102_ = _dafny.MultiSet([p2, p2, not(p2), True, p3])
                d_1169_v103_: _dafny.Map
                d_1169_v103_ = _dafny.Map({p2: p2})
                d_1170_v104_: _dafny.Map
                d_1170_v104_ = _dafny.Map({d_1169_v103_: d_1168_v102_})
                d_1171_v105_: D5
                d_1171_v105_ = D5_DC9(True, 571, d_1168_v102_)
                d_1172_v106_: _dafny.Array
                nw216_ = _dafny.Array(None, 28)
                nw216_[int(0)] = _dafny.MultiSet([p2, p3])
                nw216_[int(1)] = d_1168_v102_
                nw216_[int(2)] = default__.fm38(p3, globalState)
                nw216_[int(3)] = (default__.fm38(p2, globalState) if p2 else _dafny.MultiSet([p2, p3]))
                nw216_[int(4)] = d_1168_v102_
                nw216_[int(5)] = d_1168_v102_
                nw216_[int(6)] = d_1168_v102_
                nw216_[int(7)] = d_1168_v102_
                nw216_[int(8)] = d_1168_v102_
                nw216_[int(9)] = d_1168_v102_
                nw216_[int(10)] = _dafny.MultiSet((d_1096_v40_) + (d_1096_v40_))
                nw216_[int(11)] = d_1168_v102_
                nw216_[int(12)] = (d_1168_v102_ if (d_1165_v99_).fm8(d_1163_v98_, p2, globalState) else _dafny.MultiSet(d_1096_v40_))
                nw216_[int(13)] = d_1168_v102_
                nw216_[int(14)] = (d_1168_v102_).set(p3, default__.abs(p0))
                nw216_[int(15)] = d_1168_v102_
                nw216_[int(16)] = d_1168_v102_
                nw216_[int(17)] = d_1168_v102_
                nw216_[int(18)] = d_1168_v102_
                nw216_[int(19)] = d_1168_v102_
                nw216_[int(20)] = (((d_1170_v104_)[d_1169_v103_] if (d_1169_v103_) in (d_1170_v104_) else d_1168_v102_)) - (d_1168_v102_)
                nw216_[int(21)] = (d_1171_v105_).cf16
                nw216_[int(22)] = d_1168_v102_
                nw216_[int(23)] = (d_1168_v102_) | (d_1168_v102_)
                nw216_[int(24)] = d_1168_v102_
                nw216_[int(25)] = d_1168_v102_
                nw216_[int(26)] = (_dafny.MultiSet([(self).fm8(d_1163_v98_, p2, globalState)])).intersection(d_1168_v102_)
                nw216_[int(27)] = _dafny.MultiSet([p2, p3])
                d_1172_v106_ = nw216_
                index214_ = default__.safeIndex(744, (d_1172_v106_).length(0))
                (d_1172_v106_)[index214_] = _dafny.MultiSet([p3, p3])
                index215_ = default__.safeIndex(744, (d_1172_v106_).length(0))
                rhs181_ = (d_1167_v101_ if p2 else d_1167_v101_)
                rhs182_ = True
                rhs183_ = d_1168_v102_
                lhs99_ = d_1172_v106_
                lhs100_ = default__.safeIndex(744, (d_1172_v106_).length(0))
                d_1167_v101_ = rhs181_
                r0 = rhs182_
                lhs99_[lhs100_] = rhs183_
                d_1173_v107_: C1
                nw217_ = C1()
                nw217_.ctor__((D6_DC11(d_1163_v98_)).cf21)
                d_1173_v107_ = nw217_
                d_1174_v108_: D8
                d_1174_v108_ = D8_DC17(p2, d_1173_v107_, self.f2)
                d_1175_v109_: _dafny.Array
                nw218_ = _dafny.Array(None, 20)
                nw218_[int(0)] = p2
                nw218_[int(1)] = p3
                nw218_[int(2)] = p3
                nw218_[int(3)] = p2
                nw218_[int(4)] = True
                nw218_[int(5)] = True
                nw218_[int(6)] = not(False)
                nw218_[int(7)] = p2
                nw218_[int(8)] = not((d_1174_v108_).cf30)
                nw218_[int(9)] = p2
                nw218_[int(10)] = p2
                nw218_[int(11)] = p2
                nw218_[int(12)] = p3
                nw218_[int(13)] = p3
                nw218_[int(14)] = p3
                nw218_[int(15)] = False
                nw218_[int(16)] = p3
                nw218_[int(17)] = p2
                nw218_[int(18)] = (d_1173_v107_).fm11(globalState)
                nw218_[int(19)] = False
                d_1175_v109_ = nw218_
                d_1176_v110_: _dafny.Set
                d_1176_v110_ = _dafny.Set({d_1175_v109_})
                d_1176_v110_ = (d_1176_v110_).intersection(d_1176_v110_)
                d_1177_v111_: C1
                nw219_ = C1()
                nw219_.ctor__(((d_1173_v107_).f3) + (d_1163_v98_))
                d_1177_v111_ = nw219_
                d_1167_v101_ = d_1167_v101_
            elif True:
                d_1096_v40_ = d_1096_v40_
                d_1178_v112_: _dafny.MultiSet
                d_1178_v112_ = _dafny.MultiSet([not(p2), p2])
                d_1178_v112_ = d_1178_v112_
                d_1179_v113_: _dafny.Map
                d_1179_v113_ = _dafny.Map({p0: default__.safeDivisionInt(self.f2, (0) - (len(d_1163_v98_)))})
                d_1180_v114_: _dafny.Set
                d_1180_v114_ = _dafny.Set({p3, p3, True})
                d_1179_v113_ = (d_1179_v113_).set((0) - ((self).fm1(p3, globalState)), len(_dafny.Map({len(d_1180_v114_): p3})))
                d_1181_v115_: _dafny.Array
                def lambda39_(d_1182_v98_):
                    def lambda40_(d_1183_i11_):
                        return d_1182_v98_

                    return lambda40_

                init20_ = lambda39_(d_1163_v98_)
                nw220_ = _dafny.Array(None, 2)
                for i0_20_ in range(nw220_.length(0)):
                    nw220_[i0_20_] = init20_(i0_20_)
                d_1181_v115_ = nw220_
                index216_ = default__.safeIndex(505, (d_1181_v115_).length(0))
                (d_1181_v115_)[index216_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tybfst"))
                index217_ = default__.safeIndex(505, (d_1181_v115_).length(0))
                (d_1181_v115_)[index217_] = d_1163_v98_
                d_1184_v116_: D0
                d_1184_v116_ = D0_DC0(p3)
                d_1185_v117_: _dafny.Seq
                d_1185_v117_ = _dafny.SeqWithoutIsStrInference([d_1184_v116_])
                d_1185_v117_ = _dafny.SeqWithoutIsStrInference([d_1184_v116_ for d_1186_i12_ in range(default__.abs(13))])
            d_1187_v118_: str
            d_1187_v118_ = _dafny.CodePoint('c')
            d_1188_v119_: _dafny.MultiSet
            d_1188_v119_ = _dafny.MultiSet([d_1187_v118_])
            d_1188_v119_ = d_1188_v119_
            d_1189_v120_: _dafny.Array
            nw221_ = _dafny.Array(None, 16)
            nw221_[int(0)] = (self).fm8(d_1163_v98_, p3, globalState)
            nw221_[int(1)] = p3
            nw221_[int(2)] = p2
            nw221_[int(3)] = p2
            nw221_[int(4)] = p2
            nw221_[int(5)] = p3
            nw221_[int(6)] = p2
            nw221_[int(7)] = p3
            nw221_[int(8)] = p3
            nw221_[int(9)] = p2
            nw221_[int(10)] = p2
            nw221_[int(11)] = p2
            nw221_[int(12)] = p2
            nw221_[int(13)] = False
            nw221_[int(14)] = p2
            nw221_[int(15)] = p2
            d_1189_v120_ = nw221_
            d_1190_v121_: _dafny.Set
            d_1190_v121_ = _dafny.Set({d_1189_v120_})
            d_1191_v122_: _dafny.Map
            d_1191_v122_ = _dafny.Map({self.f2: d_1164_i10_})
            d_1192_v123_: _dafny.Set
            d_1192_v123_ = _dafny.Set({p2})
            d_1193_v124_: _dafny.Seq
            d_1193_v124_ = _dafny.SeqWithoutIsStrInference([d_1192_v123_])
            d_1194_v125_: _dafny.Array
            nw222_ = _dafny.Array(None, 12)
            nw222_[int(0)] = d_1191_v122_
            nw222_[int(1)] = d_1191_v122_
            nw222_[int(2)] = d_1191_v122_
            nw222_[int(3)] = d_1191_v122_
            nw222_[int(4)] = (d_1191_v122_) | (_dafny.Map({p0: d_1164_i10_}))
            nw222_[int(5)] = _dafny.Map({p0: (0) - (d_1164_i10_)})
            nw222_[int(6)] = d_1191_v122_
            nw222_[int(7)] = d_1191_v122_
            nw222_[int(8)] = default__.fm27(d_1164_i10_, p2, globalState)
            nw222_[int(9)] = (d_1191_v122_).set((0) - ((0) - (len(d_1193_v124_))), p0)
            nw222_[int(10)] = d_1191_v122_
            nw222_[int(11)] = d_1191_v122_
            d_1194_v125_ = nw222_
            rhs184_ = (d_1190_v121_) - (d_1190_v121_)
            rhs185_ = not(not(not((p2) and (False))))
            rhs186_ = d_1194_v125_
            rhs187_ = p3
            rhs188_ = (0) - ((self).fm1(p2, globalState))
            lhs101_ = self
            d_1190_v121_ = rhs184_
            r0 = rhs185_
            d_1194_v125_ = rhs186_
            r0 = rhs187_
            lhs101_.f2 = rhs188_
            r0 = p2
        d_1195_v126_: _dafny.Array
        nw223_ = _dafny.Array(False, 19)
        d_1195_v126_ = nw223_
        d_1195_v126_ = d_1195_v126_
        d_1196_v127_: _dafny.Array
        nw224_ = _dafny.Array(None, 2)
        nw224_[int(0)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uokfbik")))) - (p0)
        nw224_[int(1)] = p0
        d_1196_v127_ = nw224_
        d_1196_v127_ = d_1196_v127_
        r0 = ((p2) or (p2) if p3 else p3)
        d_1197_v128_: _dafny.Map
        d_1197_v128_ = _dafny.Map({p2: True})
        d_1198_v129_: _dafny.Seq
        d_1198_v129_ = _dafny.SeqWithoutIsStrInference([d_1197_v128_])
        d_1199_v130_: _dafny.Map
        d_1199_v130_ = _dafny.Map({True: p3})
        d_1200_v131_: _dafny.Seq
        d_1200_v131_ = _dafny.SeqWithoutIsStrInference([self.f2, self.f2, p0, -887])
        d_1201_v132_: _dafny.MultiSet
        d_1201_v132_ = _dafny.MultiSet([d_1200_v131_, _dafny.SeqWithoutIsStrInference([p0]), _dafny.SeqWithoutIsStrInference([-392 for d_1202_i13_ in range(default__.abs(916))])])
        d_1203_v133_: _dafny.Map
        d_1203_v133_ = _dafny.Map({d_1201_v132_: d_1197_v128_})
        d_1204_v134_: _dafny.Array
        nw225_ = _dafny.Array(None, 18)
        nw225_[int(0)] = d_1197_v128_
        nw225_[int(1)] = (d_1197_v128_) | ((d_1198_v129_)[default__.safeIndex(self.f2, len(d_1198_v129_))])
        nw225_[int(2)] = d_1197_v128_
        nw225_[int(3)] = (_dafny.Map({p3: p2})).set(p3, p3)
        nw225_[int(4)] = d_1197_v128_
        nw225_[int(5)] = _dafny.Map({p3: p2})
        nw225_[int(6)] = d_1197_v128_
        nw225_[int(7)] = d_1197_v128_
        nw225_[int(8)] = d_1197_v128_
        nw225_[int(9)] = d_1197_v128_
        nw225_[int(10)] = (d_1197_v128_) | (d_1197_v128_)
        nw225_[int(11)] = (d_1197_v128_) | (d_1197_v128_)
        nw225_[int(12)] = d_1197_v128_
        nw225_[int(13)] = _dafny.Map({p3: p3})
        nw225_[int(14)] = _dafny.Map({True: not(p3)})
        nw225_[int(15)] = d_1197_v128_
        nw225_[int(16)] = (_dafny.Map({p2: p2})) | ((d_1199_v130_))
        nw225_[int(17)] = (d_1197_v128_) | (((d_1203_v133_)[d_1201_v132_] if (d_1201_v132_) in (d_1203_v133_) else d_1197_v128_))
        d_1204_v134_ = nw225_
        r1 = d_1204_v134_
        d_1205_v135_: _dafny.Map
        d_1205_v135_ = _dafny.Map({d_1195_v126_: d_1163_v98_})
        d_1206_v136_: str
        d_1206_v136_ = _dafny.CodePoint('p')
        r2 = len(((((d_1205_v135_)[d_1195_v126_] if (d_1195_v126_) in (d_1205_v135_) else d_1163_v98_)) + ((d_1163_v98_) + (d_1163_v98_))).set(default__.safeIndex(p0, len((((d_1205_v135_)[d_1195_v126_] if (d_1195_v126_) in (d_1205_v135_) else d_1163_v98_)) + ((d_1163_v98_) + (d_1163_v98_)))), d_1206_v136_))
        return r0, r1, r2

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_1207_v0_: _dafny.Array
        def lambda41_(d_1208_p2_):
            def lambda42_(d_1209_i0_):
                return d_1208_p2_

            return lambda42_

        init21_ = lambda41_(p2)
        nw226_ = _dafny.Array(None, 7)
        for i0_21_ in range(nw226_.length(0)):
            nw226_[i0_21_] = init21_(i0_21_)
        d_1207_v0_ = nw226_
        index218_ = default__.safeIndex(777, (d_1207_v0_).length(0))
        (d_1207_v0_)[index218_] = p2
        d_1210_v1_: bool
        d_1210_v1_ = True
        d_1211_v2_: _dafny.Map
        d_1211_v2_ = _dafny.Map({p1: d_1210_v1_})
        d_1212_v4_: _dafny.Seq
        d_1212_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "geclc"))
        index219_ = default__.safeIndex(777, (d_1207_v0_).length(0))
        def iife90_():
            coll38_ = _dafny.Map()
            compr_38_: int
            for compr_38_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1213_i2_ in range(default__.abs(-31))])) for d_1214_i1_ in range(default__.abs(500))])).Elements:
                d_1215_v3_: int = compr_38_
                if (d_1215_v3_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1213_i2_ in range(default__.abs(-31))])) for d_1214_i1_ in range(default__.abs(500))])):
                    coll38_[default__.safeDivisionInt(d_1215_v3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oner"))))] = p2
            return _dafny.Map(coll38_)
        rhs189_ = self.f2
        rhs190_ = (len((d_1211_v2_) | (((iife90_()
        ).set(p0, p2)).set(p0, p2)))) < (self.f2)
        rhs191_ = ((d_1212_v4_) + (d_1212_v4_)) <= (d_1212_v4_)
        rhs192_ = (0) - (p1)
        lhs102_ = d_1207_v0_
        lhs103_ = default__.safeIndex(777, (d_1207_v0_).length(0))
        lhs104_ = self
        r0 = rhs189_
        lhs102_[lhs103_] = rhs190_
        d_1210_v1_ = rhs191_
        lhs104_.f2 = rhs192_
        d_1216_v5_: _dafny.Set
        d_1216_v5_ = _dafny.Set({self.f2, self.f2})
        d_1217_v6_: _dafny.MultiSet
        d_1217_v6_ = _dafny.MultiSet([d_1216_v5_])
        d_1218_i3_: int
        d_1218_i3_ = 0
        with _dafny.label("5"):
            while (((d_1217_v6_)[d_1216_v5_] if (d_1216_v5_) in (d_1217_v6_) else p0)) <= ((self).fm1((d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))], globalState)):
                with _dafny.c_label("5"):
                    if (d_1218_i3_) >= (100):
                        raise _dafny.Break("5")
                    d_1218_i3_ = (d_1218_i3_) + (1)
                    d_1212_v4_ = ((d_1212_v4_) + (d_1212_v4_)) + (d_1212_v4_)
                    if p2:
                        d_1219_v7_: _dafny.Set
                        d_1219_v7_ = _dafny.Set({(d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))]})
                        d_1219_v7_ = d_1219_v7_
                        d_1220_v8_: C1
                        nw227_ = C1()
                        nw227_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxtqqto")))
                        d_1220_v8_ = nw227_
                        d_1221_v9_: _dafny.Map
                        d_1221_v9_ = _dafny.Map({len(_dafny.Set({p2, d_1210_v1_})): self.f2})
                        d_1222_v11_: _dafny.Seq
                        d_1222_v11_ = _dafny.SeqWithoutIsStrInference([d_1221_v9_])
                        d_1223_v12_: _dafny.Array
                        nw228_ = _dafny.Array(None, 9)
                        nw228_[int(0)] = d_1221_v9_
                        nw228_[int(1)] = d_1221_v9_
                        nw228_[int(2)] = d_1221_v9_
                        nw228_[int(3)] = (d_1221_v9_) | (d_1221_v9_)
                        nw228_[int(4)] = d_1221_v9_
                        nw228_[int(5)] = d_1221_v9_
                        def iife91_():
                            coll39_ = _dafny.Map()
                            compr_39_: int
                            for compr_39_ in _dafny.IntegerRange(216, 603):
                                d_1224_v10_: int = compr_39_
                                if ((216) <= (d_1224_v10_)) and ((d_1224_v10_) < (603)):
                                    coll39_[(d_1224_v10_) + (len(_dafny.SeqWithoutIsStrInference([p1])))] = p1
                            return _dafny.Map(coll39_)
                        nw228_[int(6)] = iife91_()
                        
                        nw228_[int(7)] = (d_1222_v11_)[default__.safeIndex(p0, len(d_1222_v11_))]
                        nw228_[int(8)] = d_1221_v9_
                        d_1223_v12_ = nw228_
                        index220_ = default__.safeIndex(923, (d_1223_v12_).length(0))
                        (d_1223_v12_)[index220_] = d_1221_v9_
                        index221_ = default__.safeIndex(923, (d_1223_v12_).length(0))
                        rhs193_ = d_1221_v9_
                        rhs194_ = self.f2
                        lhs105_ = d_1223_v12_
                        lhs106_ = default__.safeIndex(923, (d_1223_v12_).length(0))
                        lhs107_ = self
                        lhs105_[lhs106_] = rhs193_
                        lhs107_.f2 = rhs194_
                        d_1225_v13_: _dafny.Seq
                        d_1225_v13_ = _dafny.SeqWithoutIsStrInference([690, p0, len(((d_1220_v8_).f3) + (d_1212_v4_))])
                        d_1225_v13_ = (d_1225_v13_) + (d_1225_v13_)
                        d_1226_v14_: _dafny.Array
                        nw229_ = _dafny.Array(int(0), 16)
                        d_1226_v14_ = nw229_
                        d_1226_v14_ = d_1226_v14_
                    elif True:
                        d_1227_v15_: _dafny.Array
                        def lambda43_(d_1228_v0_):
                            def lambda44_(d_1229_i4_):
                                return _dafny.Map({(d_1228_v0_)[default__.safeIndex(777, (d_1228_v0_).length(0))]: True})

                            return lambda44_

                        init22_ = lambda43_(d_1207_v0_)
                        nw230_ = _dafny.Array(None, 7)
                        for i0_22_ in range(nw230_.length(0)):
                            nw230_[i0_22_] = init22_(i0_22_)
                        d_1227_v15_ = nw230_
                        d_1230_v16_: _dafny.Seq
                        d_1230_v16_ = _dafny.SeqWithoutIsStrInference([d_1227_v15_])
                        d_1227_v15_ = (d_1230_v16_)[default__.safeIndex(self.f2, len(d_1230_v16_))]
                        d_1231_v17_: _dafny.Array
                        def lambda45_(d_1232_i5_):
                            return (d_1232_i5_) - (771)

                        init23_ = lambda45_
                        nw231_ = _dafny.Array(None, 14)
                        for i0_23_ in range(nw231_.length(0)):
                            nw231_[i0_23_] = init23_(i0_23_)
                        d_1231_v17_ = nw231_
                        d_1233_v18_: _dafny.Seq
                        d_1233_v18_ = _dafny.SeqWithoutIsStrInference([d_1231_v17_])
                        d_1234_v19_: str
                        d_1234_v19_ = _dafny.CodePoint('i')
                        d_1235_v20_: D4
                        d_1235_v20_ = D4_DC6(p1, d_1231_v17_, p0)
                        d_1236_v21_: _dafny.Array
                        nw232_ = _dafny.Array(None, 20)
                        nw232_[int(0)] = d_1231_v17_
                        nw232_[int(1)] = (d_1231_v17_ if p2 else d_1231_v17_)
                        nw232_[int(2)] = (d_1231_v17_ if (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))] else (d_1233_v18_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udwhy"))), len(d_1233_v18_))])
                        nw232_[int(3)] = (d_1233_v18_)[default__.safeIndex(len(((d_1212_v4_).set(default__.safeIndex(p1, len(d_1212_v4_)), d_1234_v19_)).set(default__.safeIndex(self.f2, len((d_1212_v4_).set(default__.safeIndex(p1, len(d_1212_v4_)), d_1234_v19_))), d_1234_v19_)), len(d_1233_v18_))]
                        nw232_[int(4)] = d_1231_v17_
                        nw232_[int(5)] = d_1231_v17_
                        nw232_[int(6)] = d_1231_v17_
                        nw232_[int(7)] = (d_1231_v17_ if not(True) else d_1231_v17_)
                        nw232_[int(8)] = (d_1235_v20_).cf10
                        nw232_[int(9)] = d_1231_v17_
                        nw232_[int(10)] = d_1231_v17_
                        nw232_[int(11)] = d_1231_v17_
                        nw232_[int(12)] = d_1231_v17_
                        nw232_[int(13)] = d_1231_v17_
                        nw232_[int(14)] = d_1231_v17_
                        nw232_[int(15)] = d_1231_v17_
                        nw232_[int(16)] = d_1231_v17_
                        nw232_[int(17)] = d_1231_v17_
                        nw232_[int(18)] = d_1231_v17_
                        nw232_[int(19)] = d_1231_v17_
                        d_1236_v21_ = nw232_
                        index222_ = default__.safeIndex(248, (d_1236_v21_).length(0))
                        (d_1236_v21_)[index222_] = (d_1231_v17_ if d_1210_v1_ else d_1231_v17_)
                        index223_ = default__.safeIndex(248, (d_1236_v21_).length(0))
                        (d_1236_v21_)[index223_] = (d_1235_v20_).cf10
                        d_1237_v22_: _dafny.Map
                        d_1237_v22_ = _dafny.Map({p0: d_1212_v4_})
                        d_1238_v23_: _dafny.Seq
                        d_1238_v23_ = _dafny.SeqWithoutIsStrInference([self.f2, self.f2])
                        d_1239_v24_: _dafny.Seq
                        d_1239_v24_ = _dafny.SeqWithoutIsStrInference([d_1210_v1_, d_1210_v1_])
                        d_1240_v25_: _dafny.Seq
                        d_1240_v25_ = _dafny.SeqWithoutIsStrInference([(d_1238_v23_)[default__.safeIndex(len(d_1239_v24_), len(d_1238_v23_))]])
                        d_1241_v26_: _dafny.Map
                        d_1241_v26_ = _dafny.Map({p2: p0})
                        d_1212_v4_ = ((d_1237_v22_)[p1] if (p1) in (d_1237_v22_) else (d_1212_v4_).set(default__.safeIndex(len(_dafny.Set({d_1240_v25_, d_1240_v25_, _dafny.SeqWithoutIsStrInference([len(d_1241_v26_)]), d_1240_v25_})), len(d_1212_v4_)), d_1234_v19_))
                        d_1242_v27_: _dafny.Array
                        def lambda46_(d_1243_v23_, d_1244_p1_):
                            def lambda47_(d_1245_i6_):
                                return ((d_1243_v23_) + ((d_1243_v23_).set(default__.safeIndex(d_1244_p1_, len(d_1243_v23_)), self.f2))).set(default__.safeIndex(len(d_1243_v23_), len((d_1243_v23_) + ((d_1243_v23_).set(default__.safeIndex(d_1244_p1_, len(d_1243_v23_)), self.f2)))), self.f2)

                            return lambda47_

                        init24_ = lambda46_(d_1238_v23_, p1)
                        nw233_ = _dafny.Array(None, 19)
                        for i0_24_ in range(nw233_.length(0)):
                            nw233_[i0_24_] = init24_(i0_24_)
                        d_1242_v27_ = nw233_
                        index224_ = default__.safeIndex(540, (d_1242_v27_).length(0))
                        (d_1242_v27_)[index224_] = d_1238_v23_
                        d_1246_v28_: _dafny.Set
                        d_1246_v28_ = _dafny.Set({False})
                        index225_ = default__.safeIndex(540, (d_1242_v27_).length(0))
                        (d_1242_v27_)[index225_] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1246_v28_, _dafny.Set({p2})]))])
                        d_1247_v29_: C2
                        nw234_ = C2()
                        nw234_.ctor__(self.f2)
                        d_1247_v29_ = nw234_
                    index226_ = default__.safeIndex(777, (d_1207_v0_).length(0))
                    (d_1207_v0_)[index226_] = p2
                    index227_ = default__.safeIndex(777, (d_1207_v0_).length(0))
                    (d_1207_v0_)[index227_] = d_1210_v1_
                    pass
            pass
        if d_1210_v1_:
            d_1210_v1_ = d_1210_v1_
            r0 = (0) - (self.f2)
            r0 = p0
            (self).f2 = len(d_1212_v4_)
            d_1248_v30_: _dafny.Map
            d_1248_v30_ = _dafny.Map({(d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))]: -750})
            index228_ = default__.safeIndex(777, (d_1207_v0_).length(0))
            (d_1207_v0_)[index228_] = (-831) > (len(d_1248_v30_))
        elif True:
            d_1249_v31_: _dafny.Map
            d_1249_v31_ = _dafny.Map({d_1210_v1_: (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))]})
            d_1250_v32_: _dafny.Seq
            d_1250_v32_ = _dafny.SeqWithoutIsStrInference([-366, self.f2, default__.safeModuloInt(len(d_1249_v31_), len(d_1211_v2_)), p0])
            d_1250_v32_ = d_1250_v32_
            r0 = (self).fm1(p2, globalState)
            d_1251_v33_: _dafny.Set
            d_1251_v33_ = _dafny.Set({p2, d_1210_v1_})
            d_1252_v34_: _dafny.Seq
            d_1252_v34_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({True, (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))]}), d_1251_v33_, d_1251_v33_, d_1251_v33_])
            d_1253_v35_: _dafny.Map
            d_1253_v35_ = _dafny.Map({p0: (p0 if (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))] else len((d_1252_v34_)[default__.safeIndex((0) - (p0), len(d_1252_v34_))]))})
            d_1253_v35_ = (d_1253_v35_).set(p0, p0)
            d_1254_v36_: _dafny.MultiSet
            d_1254_v36_ = _dafny.MultiSet([(d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))], (p2 if True else d_1210_v1_), d_1210_v1_])
            rhs195_ = ((d_1254_v36_)[(d_1216_v5_).issubset(d_1216_v5_)] if ((d_1216_v5_).issubset(d_1216_v5_)) in (d_1254_v36_) else ((0) - (len(d_1212_v4_))) * (p1))
            rhs196_ = default__.safeModuloInt(len(d_1251_v33_), (self.f2) * (p0))
            lhs108_ = self
            lhs109_ = self
            lhs108_.f2 = rhs195_
            lhs109_.f2 = rhs196_
            d_1255_v37_: _dafny.Array
            def lambda48_(d_1256_p0_):
                def lambda49_(d_1257_i7_):
                    return (d_1257_i7_) + (d_1256_p0_)

                return lambda49_

            init25_ = lambda48_(p0)
            nw235_ = _dafny.Array(None, 25)
            for i0_25_ in range(nw235_.length(0)):
                nw235_[i0_25_] = init25_(i0_25_)
            d_1255_v37_ = nw235_
            d_1258_v38_: _dafny.MultiSet
            d_1258_v38_ = _dafny.MultiSet([993, self.f2])
            d_1259_v39_: str
            d_1259_v39_ = _dafny.CodePoint('j')
            index229_ = default__.safeIndex(777, (d_1207_v0_).length(0))
            rhs197_ = d_1210_v1_
            rhs198_ = (self).fm8((d_1212_v4_).set(default__.safeIndex(p1, len(d_1212_v4_)), d_1259_v39_), (d_1254_v36_).ispropersubset(d_1254_v36_), globalState)
            rhs199_ = d_1255_v37_
            rhs200_ = ((d_1258_v38_) | (d_1258_v38_)) | ((d_1258_v38_) | (d_1258_v38_))
            rhs201_ = (d_1250_v32_) + (d_1250_v32_)
            lhs110_ = d_1207_v0_
            lhs111_ = default__.safeIndex(777, (d_1207_v0_).length(0))
            lhs110_[lhs111_] = rhs197_
            d_1210_v1_ = rhs198_
            d_1255_v37_ = rhs199_
            d_1258_v38_ = rhs200_
            d_1250_v32_ = rhs201_
        d_1260_v40_: _dafny.Map
        d_1260_v40_ = _dafny.Map({True: 486})
        d_1261_v41_: _dafny.Set
        d_1261_v41_ = _dafny.Set({d_1260_v40_})
        hi5_ = (p1) - (p1)
        for d_1262_i8_ in range(len((_dafny.Set({(d_1260_v40_).set(p2, p1), d_1260_v40_, d_1260_v40_})) - (d_1261_v41_)), hi5_):
            d_1263_v42_: _dafny.Set
            d_1263_v42_ = _dafny.Set({(d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))]})
            if (d_1263_v42_).issubset(d_1263_v42_):
                d_1264_v43_: _dafny.Array
                nw236_ = _dafny.Array(int(0), 7)
                d_1264_v43_ = nw236_
                d_1265_v44_: _dafny.Set
                d_1265_v44_ = _dafny.Set({d_1264_v43_, d_1264_v43_, d_1264_v43_})
                d_1265_v44_ = ((_dafny.Set({d_1264_v43_}) if False else d_1265_v44_)).intersection((d_1265_v44_) | (d_1265_v44_))
                d_1266_v45_: _dafny.Array
                def lambda50_(d_1267_v2_):
                    def lambda51_(d_1268_i9_):
                        return d_1267_v2_

                    return lambda51_

                init26_ = lambda50_(d_1211_v2_)
                nw237_ = _dafny.Array(None, 10)
                for i0_26_ in range(nw237_.length(0)):
                    nw237_[i0_26_] = init26_(i0_26_)
                d_1266_v45_ = nw237_
                index230_ = default__.safeIndex(137, (d_1266_v45_).length(0))
                (d_1266_v45_)[index230_] = d_1211_v2_
                d_1269_v46_: _dafny.Map
                d_1269_v46_ = _dafny.Map({d_1212_v4_: _dafny.Map({len(d_1216_v5_): (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))]})})
                d_1270_v47_: D5
                d_1270_v47_ = D5_DC10(d_1210_v1_, p0, d_1269_v46_, d_1210_v1_)
                d_1271_v48_: _dafny.Seq
                d_1271_v48_ = _dafny.SeqWithoutIsStrInference([d_1210_v1_])
                d_1272_v49_: _dafny.MultiSet
                d_1272_v49_ = _dafny.MultiSet([d_1210_v1_, (d_1271_v48_)[default__.safeIndex(len(d_1271_v48_), len(d_1271_v48_))]])
                d_1273_v50_: str
                d_1273_v50_ = _dafny.CodePoint('e')
                d_1274_v51_: D0
                d_1274_v51_ = D0_DC1(p2, p2, d_1273_v50_, p1)
                index231_ = default__.safeIndex(137, (d_1266_v45_).length(0))
                rhs202_ = -912
                rhs203_ = (((_dafny.Map({d_1210_v1_: (d_1272_v49_).cardinality})).set((d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))], len(d_1216_v5_))) | (d_1260_v40_) if (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))] else d_1260_v40_)
                rhs204_ = ((d_1211_v2_) | (d_1211_v2_)).set((self).fm1((self).fm8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")), False, globalState), globalState), (d_1274_v51_).cf1)
                rhs205_ = d_1270_v47_
                lhs112_ = self
                lhs113_ = d_1266_v45_
                lhs114_ = default__.safeIndex(137, (d_1266_v45_).length(0))
                lhs112_.f2 = rhs202_
                d_1260_v40_ = rhs203_
                lhs113_[lhs114_] = rhs204_
                d_1270_v47_ = rhs205_
                d_1275_v52_: _dafny.Array
                nw238_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
                d_1275_v52_ = nw238_
                index232_ = default__.safeIndex(91, (d_1275_v52_).length(0))
                (d_1275_v52_)[index232_] = d_1212_v4_
                index233_ = default__.safeIndex(91, (d_1275_v52_).length(0))
                (d_1275_v52_)[index233_] = d_1212_v4_
                d_1276_v53_: C1
                nw239_ = C1()
                nw239_.ctor__(d_1212_v4_)
                d_1276_v53_ = nw239_
                d_1276_v53_ = d_1276_v53_
                d_1277_v54_: _dafny.Map
                d_1277_v54_ = _dafny.Map({self.f2: d_1271_v48_})
                d_1278_v55_: _dafny.Seq
                d_1278_v55_ = _dafny.SeqWithoutIsStrInference([(d_1271_v48_).set(default__.safeIndex(self.f2, len(d_1271_v48_)), False), ((d_1271_v48_).set(default__.safeIndex(p1, len(d_1271_v48_)), d_1210_v1_)).set(default__.safeIndex(p1, len((d_1271_v48_).set(default__.safeIndex(p1, len(d_1271_v48_)), d_1210_v1_))), not(True)), d_1271_v48_, (((d_1277_v54_)[-989] if (-989) in (d_1277_v54_) else d_1271_v48_)) + ((d_1271_v48_).set(default__.safeIndex(p0, len(d_1271_v48_)), False)), d_1271_v48_])
                d_1279_v56_: _dafny.MultiSet
                d_1279_v56_ = _dafny.MultiSet([len(d_1263_v42_), self.f2])
                d_1280_v57_: _dafny.Map
                d_1280_v57_ = _dafny.Map({d_1210_v1_: d_1210_v1_})
                d_1281_v58_: D8
                d_1281_v58_ = D8_DC17((d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))], d_1276_v53_, len(d_1280_v57_))
                d_1271_v48_ = (d_1278_v55_)[default__.safeIndex(((d_1279_v56_)[(d_1281_v58_).cf32] if ((d_1281_v58_).cf32) in (d_1279_v56_) else p0), len(d_1278_v55_))]
            elif True:
                d_1282_v59_: _dafny.MultiSet
                d_1282_v59_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsrqrrepm"))])
                index234_ = default__.safeIndex(777, (d_1207_v0_).length(0))
                (d_1207_v0_)[index234_] = (default__.fm30(d_1210_v1_, globalState)).issubset((d_1282_v59_) | (d_1282_v59_))
                d_1283_v60_: _dafny.Array
                def lambda52_(d_1284_v4_):
                    def lambda53_(d_1285_i10_):
                        return (d_1285_i10_) - (len(d_1284_v4_))

                    return lambda53_

                init27_ = lambda52_(d_1212_v4_)
                nw240_ = _dafny.Array(None, 23)
                for i0_27_ in range(nw240_.length(0)):
                    nw240_[i0_27_] = init27_(i0_27_)
                d_1283_v60_ = nw240_
                d_1283_v60_ = d_1283_v60_
                d_1286_v61_: _dafny.MultiSet
                d_1286_v61_ = _dafny.MultiSet([self.f2])
                d_1287_v62_: str
                d_1287_v62_ = _dafny.CodePoint('p')
                d_1288_v63_: _dafny.MultiSet
                d_1288_v63_ = _dafny.MultiSet([d_1210_v1_])
                d_1289_v64_: D0
                d_1289_v64_ = D0_DC1((d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))], d_1210_v1_, d_1287_v62_, p0)
                d_1290_v65_: _dafny.Map
                d_1290_v65_ = _dafny.Map({_dafny.CodePoint('v'): _dafny.MultiSet([(d_1289_v64_).cf4])})
                d_1291_v66_: _dafny.MultiSet
                d_1291_v66_ = d_1286_v61_
                d_1292_v67_: _dafny.Seq
                d_1292_v67_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(0) - (self.f2)])])
                d_1293_v68_: _dafny.Array
                nw241_ = _dafny.Array(None, 25)
                nw241_[int(0)] = (d_1286_v61_).intersection(d_1286_v61_)
                nw241_[int(1)] = d_1286_v61_
                nw241_[int(2)] = (_dafny.MultiSet([self.f2, p1])) - (d_1286_v61_)
                nw241_[int(3)] = d_1286_v61_
                nw241_[int(4)] = (default__.fm35(d_1287_v62_, d_1210_v1_, d_1288_v63_, globalState)).intersection(d_1286_v61_)
                nw241_[int(5)] = d_1286_v61_
                nw241_[int(6)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([924 for d_1294_i11_ in range(default__.abs(620))]))) | (d_1286_v61_)
                nw241_[int(7)] = _dafny.MultiSet([p0, p0, p0, p0, self.f2])
                nw241_[int(8)] = default__.fm19(712, p0, (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))], self.f2, globalState)
                nw241_[int(9)] = (d_1286_v61_) - (((d_1290_v65_)[d_1287_v62_] if (d_1287_v62_) in (d_1290_v65_) else _dafny.MultiSet([len(d_1211_v2_), p1])))
                nw241_[int(10)] = d_1286_v61_
                nw241_[int(11)] = (d_1286_v61_).intersection(d_1286_v61_)
                nw241_[int(12)] = ((d_1286_v61_).set(p1, default__.abs(p1))).intersection((d_1291_v66_))
                nw241_[int(13)] = d_1286_v61_
                nw241_[int(14)] = (d_1292_v67_)[default__.safeIndex(p0, len(d_1292_v67_))]
                nw241_[int(15)] = d_1286_v61_
                nw241_[int(16)] = d_1286_v61_
                nw241_[int(17)] = d_1286_v61_
                nw241_[int(18)] = (_dafny.MultiSet([(self).fm1(p2, globalState)]) if True else d_1286_v61_)
                nw241_[int(19)] = default__.fm35(d_1287_v62_, (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))], d_1288_v63_, globalState)
                nw241_[int(20)] = d_1286_v61_
                nw241_[int(21)] = (d_1286_v61_).set(d_1262_i8_, default__.abs(55))
                nw241_[int(22)] = _dafny.MultiSet([d_1262_i8_])
                nw241_[int(23)] = d_1286_v61_
                nw241_[int(24)] = d_1286_v61_
                d_1293_v68_ = nw241_
                index235_ = default__.safeIndex(812, (d_1293_v68_).length(0))
                (d_1293_v68_)[index235_] = (d_1286_v61_) - (d_1286_v61_)
                index236_ = default__.safeIndex(812, (d_1293_v68_).length(0))
                (d_1293_v68_)[index236_] = d_1286_v61_
                d_1287_v62_ = d_1287_v62_
                index237_ = default__.safeIndex(777, (d_1207_v0_).length(0))
                (d_1207_v0_)[index237_] = (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))]
            d_1295_v69_: str
            d_1295_v69_ = _dafny.CodePoint('s')
            d_1296_v71_: _dafny.Array
            def lambda54_(d_1297_p0_, d_1298_p1_):
                def lambda55_(d_1299_i12_):
                    def iife92_():
                        coll40_ = _dafny.Map()
                        compr_40_: int
                        for compr_40_ in _dafny.IntegerRange(895, 618):
                            d_1300_v70_: int = compr_40_
                            if ((895) <= (d_1300_v70_)) and ((d_1300_v70_) < (618)):
                                coll40_[(d_1300_v70_) * (190)] = _dafny.Map({d_1298_p1_: self.f2})
                        return _dafny.Map(coll40_)
                    return (_dafny.Map({len(_dafny.Map({len(iife92_()
                    ): _dafny.Map({(0) - (d_1298_p1_): d_1297_p0_})})): d_1297_p0_})) | (_dafny.Map({self.f2: self.f2}))

                return lambda55_

            init28_ = lambda54_(p0, p1)
            nw242_ = _dafny.Array(None, 25)
            for i0_28_ in range(nw242_.length(0)):
                nw242_[i0_28_] = init28_(i0_28_)
            d_1296_v71_ = nw242_
            d_1301_v73_: _dafny.MultiSet
            d_1301_v73_ = _dafny.MultiSet([len(d_1212_v4_)])
            d_1302_v74_: _dafny.Map
            d_1302_v74_ = _dafny.Map({((d_1301_v73_)[d_1262_i8_] if (d_1262_i8_) in (d_1301_v73_) else 47): d_1262_i8_})
            d_1303_v75_: _dafny.Seq
            d_1303_v75_ = _dafny.SeqWithoutIsStrInference([d_1302_v74_])
            index238_ = default__.safeIndex(4, (d_1296_v71_).length(0))
            def iife93_():
                coll41_ = _dafny.Map()
                compr_41_: int
                for compr_41_ in ((d_1303_v75_)[default__.safeIndex(self.f2, len(d_1303_v75_))]).keys.Elements:
                    d_1304_v72_: int = compr_41_
                    if (d_1304_v72_) in ((d_1303_v75_)[default__.safeIndex(self.f2, len(d_1303_v75_))]):
                        coll41_[(d_1304_v72_) + ((0) - (self.f2))] = p1
                return _dafny.Map(coll41_)
            (d_1296_v71_)[index238_] = iife93_()
            
            d_1305_v76_: _dafny.Map
            d_1305_v76_ = _dafny.Map({(d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))]: d_1210_v1_})
            index239_ = default__.safeIndex(4, (d_1296_v71_).length(0))
            index240_ = default__.safeIndex(777, (d_1207_v0_).length(0))
            def iife94_():
                coll42_ = _dafny.Map()
                compr_42_: int
                for compr_42_ in _dafny.IntegerRange(105, 303):
                    d_1306_v77_: int = compr_42_
                    if ((105) <= (d_1306_v77_)) and ((d_1306_v77_) < (303)):
                        coll42_[(d_1306_v77_) * (len((d_1305_v76_).set(p2, d_1210_v1_)))] = d_1262_i8_
                return _dafny.Map(coll42_)
            rhs206_ = default__.safeDivisionInt(len(d_1305_v76_), p1)
            rhs207_ = _dafny.CodePoint('u')
            rhs208_ = iife94_()

            rhs209_ = (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([169]))}) for d_1307_i13_ in range(default__.abs(993))]))) == (self.f2)
            rhs210_ = False
            lhs115_ = d_1296_v71_
            lhs116_ = default__.safeIndex(4, (d_1296_v71_).length(0))
            lhs117_ = d_1207_v0_
            lhs118_ = default__.safeIndex(777, (d_1207_v0_).length(0))
            r0 = rhs206_
            d_1295_v69_ = rhs207_
            lhs115_[lhs116_] = rhs208_
            lhs117_[lhs118_] = rhs209_
            d_1210_v1_ = rhs210_
            d_1308_v78_: _dafny.Seq
            d_1308_v78_ = _dafny.SeqWithoutIsStrInference([d_1210_v1_])
            index241_ = default__.safeIndex(777, (d_1207_v0_).length(0))
            (d_1207_v0_)[index241_] = ((d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))] if p2 else (d_1308_v78_)[default__.safeIndex(((d_1260_v40_)[(d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))]] if ((d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))]) in (d_1260_v40_) else d_1262_i8_), len(d_1308_v78_))])
            d_1309_v79_: _dafny.Map
            d_1309_v79_ = _dafny.Map({d_1212_v4_: len((d_1263_v42_).intersection(d_1263_v42_))})
            d_1309_v79_ = (d_1309_v79_).set(d_1212_v4_, 255)
        (self).f2 = p1
        d_1310_v80_: _dafny.Seq
        d_1310_v80_ = _dafny.SeqWithoutIsStrInference([p0, self.f2, p1])
        d_1311_v81_: _dafny.Seq
        d_1311_v81_ = _dafny.SeqWithoutIsStrInference([d_1310_v80_])
        if (d_1311_v81_) < (d_1311_v81_):
            d_1210_v1_ = p2
            d_1312_v82_: str
            d_1312_v82_ = _dafny.CodePoint('k')
            r0 = len((d_1212_v4_).set(default__.safeIndex(self.f2, len(d_1212_v4_)), d_1312_v82_))
            d_1313_v83_: _dafny.Array
            nw243_ = _dafny.Array(_dafny.Array(None, 0), 13)
            d_1313_v83_ = nw243_
            d_1314_v84_: _dafny.Map
            d_1314_v84_ = _dafny.Map({(self).fm8(d_1212_v4_, (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))], globalState): d_1210_v1_})
            d_1315_v85_: _dafny.Map
            d_1315_v85_ = d_1314_v84_
            d_1316_v86_: _dafny.Map
            d_1316_v86_ = _dafny.Map({self.f2: d_1315_v85_})
            d_1317_v87_: _dafny.Seq
            d_1317_v87_ = _dafny.SeqWithoutIsStrInference([d_1316_v86_])
            d_1318_v88_: _dafny.Array
            nw244_ = _dafny.Array(None, 10)
            nw244_[int(0)] = (d_1317_v87_)[default__.safeIndex(p1, len(d_1317_v87_))]
            nw244_[int(1)] = d_1316_v86_
            nw244_[int(2)] = d_1316_v86_
            nw244_[int(3)] = d_1316_v86_
            nw244_[int(4)] = d_1316_v86_
            nw244_[int(5)] = d_1316_v86_
            nw244_[int(6)] = d_1316_v86_
            nw244_[int(7)] = default__.fm39(globalState)
            nw244_[int(8)] = d_1316_v86_
            nw244_[int(9)] = _dafny.Map({449: d_1314_v84_})
            d_1318_v88_ = nw244_
            index242_ = default__.safeIndex(182, (d_1313_v83_).length(0))
            (d_1313_v83_)[index242_] = d_1318_v88_
            d_1319_v89_: _dafny.Map
            d_1319_v89_ = _dafny.Map({p2: d_1216_v5_})
            index243_ = default__.safeIndex(182, (d_1313_v83_).length(0))
            index244_ = default__.safeIndex(777, (d_1207_v0_).length(0))
            index245_ = default__.safeIndex(777, (d_1207_v0_).length(0))
            index246_ = default__.safeIndex(777, (d_1207_v0_).length(0))
            def iife95_():
                coll43_ = _dafny.Set()
                compr_43_: int
                for compr_43_ in _dafny.IntegerRange(56, 123):
                    d_1320_v90_: int = compr_43_
                    if ((56) <= (d_1320_v90_)) and ((d_1320_v90_) < (123)):
                        coll43_ = coll43_.union(_dafny.Set([default__.safeDivisionInt(d_1320_v90_, len(d_1212_v4_))]))
                return _dafny.Set(coll43_)
            rhs211_ = d_1318_v88_
            rhs212_ = (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))]
            rhs213_ = ((d_1216_v5_).intersection(d_1216_v5_)).ispropersubset((((d_1319_v89_)[d_1210_v1_] if (d_1210_v1_) in (d_1319_v89_) else iife95_()
            )) - (d_1216_v5_))
            rhs214_ = not(((0) - ((0) - (self.f2))) not in (d_1211_v2_))
            lhs119_ = d_1313_v83_
            lhs120_ = default__.safeIndex(182, (d_1313_v83_).length(0))
            lhs121_ = d_1207_v0_
            lhs122_ = default__.safeIndex(777, (d_1207_v0_).length(0))
            lhs123_ = d_1207_v0_
            lhs124_ = default__.safeIndex(777, (d_1207_v0_).length(0))
            lhs125_ = d_1207_v0_
            lhs126_ = default__.safeIndex(777, (d_1207_v0_).length(0))
            lhs119_[lhs120_] = rhs211_
            lhs121_[lhs122_] = rhs212_
            lhs123_[lhs124_] = rhs213_
            lhs125_[lhs126_] = rhs214_
            d_1210_v1_ = (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))]
            d_1321_v91_: _dafny.Array
            nw245_ = _dafny.Array(_dafny.MultiSet({}), 29)
            d_1321_v91_ = nw245_
            d_1322_v92_: _dafny.MultiSet
            d_1322_v92_ = _dafny.MultiSet([self.f2])
            d_1323_v93_: _dafny.MultiSet
            d_1323_v93_ = d_1322_v92_
            index247_ = default__.safeIndex(330, (d_1321_v91_).length(0))
            (d_1321_v91_)[index247_] = d_1323_v93_
            index248_ = default__.safeIndex(330, (d_1321_v91_).length(0))
            (d_1321_v91_)[index248_] = d_1322_v92_
        elif True:
            if ((p0) * (p0)) >= ((p0 if p2 else p0)):
                r0 = p1
                d_1324_v94_: C0
                nw246_ = C0()
                nw246_.ctor__(p2, (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))])
                d_1324_v94_ = nw246_
                d_1325_v95_: _dafny.Seq
                d_1325_v95_ = _dafny.SeqWithoutIsStrInference([d_1212_v4_])
                d_1326_v96_: _dafny.Set
                d_1326_v96_ = _dafny.Set({d_1212_v4_, (d_1325_v95_)[default__.safeIndex(self.f2, len(d_1325_v95_))]})
                d_1326_v96_ = d_1326_v96_
                d_1327_v97_: _dafny.MultiSet
                d_1327_v97_ = _dafny.MultiSet([d_1324_v94_.f4])
                d_1328_v98_: _dafny.Seq
                d_1328_v98_ = _dafny.SeqWithoutIsStrInference([d_1327_v97_])
                d_1329_v99_: _dafny.Array
                nw247_ = _dafny.Array(None, 15)
                nw247_[int(0)] = 746
                nw247_[int(1)] = p0
                nw247_[int(2)] = p0
                nw247_[int(3)] = -554
                nw247_[int(4)] = (0) - (p1)
                nw247_[int(5)] = p1
                nw247_[int(6)] = self.f2
                nw247_[int(7)] = 203
                nw247_[int(8)] = p0
                nw247_[int(9)] = self.f2
                nw247_[int(10)] = p0
                nw247_[int(11)] = self.f2
                nw247_[int(12)] = self.f2
                nw247_[int(13)] = p0
                nw247_[int(14)] = -538
                d_1329_v99_ = nw247_
                d_1330_v100_: _dafny.Map
                d_1330_v100_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcgsm")): d_1329_v99_})
                d_1331_v101_: str
                d_1331_v101_ = _dafny.CodePoint('q')
                r0 = len(((d_1328_v98_) + (d_1328_v98_)) + ((default__.fm40(len(d_1330_v100_), d_1331_v101_, (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))], p1, globalState)) + (d_1328_v98_)))
                (d_1324_v94_).f4 = (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))]
            elif True:
                d_1210_v1_ = not((861) > ((self).fm1((d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))], globalState)))
                d_1216_v5_ = d_1216_v5_
                r0 = (default__.safeModuloInt(p1, (0) - (p0))) + ((0) - (self.f2))
                index249_ = default__.safeIndex(777, (d_1207_v0_).length(0))
                (d_1207_v0_)[index249_] = p2
                d_1210_v1_ = p2
            d_1332_v102_: _dafny.Seq
            d_1332_v102_ = _dafny.SeqWithoutIsStrInference([d_1212_v4_])
            d_1333_v104_: _dafny.Set
            d_1333_v104_ = _dafny.Set({d_1212_v4_})
            index250_ = default__.safeIndex(777, (d_1207_v0_).length(0))
            def iife96_():
                coll44_ = _dafny.Set()
                compr_44_: _dafny.Seq
                for compr_44_ in (d_1332_v102_).Elements:
                    d_1334_v103_: _dafny.Seq = compr_44_
                    if (d_1334_v103_) in (d_1332_v102_):
                        coll44_ = coll44_.union(_dafny.Set([d_1334_v103_]))
                return _dafny.Set(coll44_)
            (d_1207_v0_)[index250_] = (iife96_()
            ) == (d_1333_v104_)
            d_1335_v105_: _dafny.Map
            d_1335_v105_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0, 167]): len(d_1310_v80_)})
            d_1336_v106_: _dafny.Map
            d_1336_v106_ = _dafny.Map({p1: d_1335_v105_})
            d_1337_v107_: _dafny.MultiSet
            d_1337_v107_ = _dafny.MultiSet([(d_1336_v106_) | (_dafny.Map({685: d_1335_v105_})), d_1336_v106_])
            pat_let_tv28_ = d_1310_v80_
            def iife97_():
                coll45_ = _dafny.Map()
                compr_45_: int
                for compr_45_ in _dafny.IntegerRange(517, 472):
                    d_1338_v108_: int = compr_45_
                    if ((517) <= (d_1338_v108_)) and ((d_1338_v108_) < (472)):
                        def iife98_(_pat_let26_0):
                            def iife99_(d_1339_dt__update__tmp_h1_):
                                def iife100_(_pat_let27_0):
                                    def iife101_(d_1340_dt__update_hcf6_h0_):
                                        return d_1340_dt__update_hcf6_h0_
                                    return iife101_(_pat_let27_0)
                                return iife100_(_dafny.Map({pat_let_tv28_: self.f2}))
                            return iife99_(_pat_let26_0)
                        coll45_[(d_1338_v108_) * ((d_1310_v80_)[default__.safeIndex(len(_dafny.Set({_dafny.SeqWithoutIsStrInference([p2])})), len(d_1310_v80_))])] = iife98_(d_1335_v105_)
                return _dafny.Map(coll45_)
            d_1337_v107_ = _dafny.MultiSet([d_1336_v106_, d_1336_v106_, iife97_()
            , d_1336_v106_, d_1336_v106_])
            d_1341_v109_: _dafny.Array
            nw248_ = _dafny.Array(None, 6)
            nw248_[int(0)] = p0
            nw248_[int(1)] = self.f2
            nw248_[int(2)] = (self).fm1((d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))], globalState)
            nw248_[int(3)] = ((D6_DC13(p2, self.f2)).cf25) + (825)
            nw248_[int(4)] = p1
            nw248_[int(5)] = ((d_1260_v40_)[p2] if (p2) in (d_1260_v40_) else p1)
            d_1341_v109_ = nw248_
            index251_ = default__.safeIndex(227, (d_1341_v109_).length(0))
            (d_1341_v109_)[index251_] = p1
            index252_ = default__.safeIndex(227, (d_1341_v109_).length(0))
            (d_1341_v109_)[index252_] = (p0 if not(p2) else p0)
            d_1342_v110_: str
            d_1342_v110_ = _dafny.CodePoint('a')
            d_1343_v111_: _dafny.Array
            nw249_ = _dafny.Array(None, 9)
            nw249_[int(0)] = d_1342_v110_
            nw249_[int(1)] = _dafny.CodePoint('k')
            nw249_[int(2)] = default__.fm37(p1, globalState)
            nw249_[int(3)] = d_1342_v110_
            nw249_[int(4)] = d_1342_v110_
            nw249_[int(5)] = d_1342_v110_
            nw249_[int(6)] = (D0_DC1(p2, (d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))], d_1342_v110_, self.f2)).cf3
            nw249_[int(7)] = d_1342_v110_
            nw249_[int(8)] = d_1342_v110_
            d_1343_v111_ = nw249_
            index253_ = default__.safeIndex(419, (d_1343_v111_).length(0))
            (d_1343_v111_)[index253_] = _dafny.CodePoint('b')
            index254_ = default__.safeIndex(419, (d_1343_v111_).length(0))
            (d_1343_v111_)[index254_] = d_1342_v110_
        r0 = (p0) * ((d_1310_v80_)[default__.safeIndex((self).fm1((d_1207_v0_)[default__.safeIndex(777, (d_1207_v0_).length(0))], globalState), len(d_1310_v80_))])
        return r0

    def m5(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        d_1344_v0_: _dafny.Array
        nw250_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
        d_1344_v0_ = nw250_
        d_1344_v0_ = d_1344_v0_
        d_1345_v1_: _dafny.Seq
        d_1345_v1_ = _dafny.SeqWithoutIsStrInference([self.f2])
        hi6_ = (d_1345_v1_)[default__.safeIndex(p0, len(d_1345_v1_))]
        for d_1346_i0_ in range(p0, hi6_):
            d_1347_v2_: bool
            d_1347_v2_ = True
            if (d_1347_v2_ if d_1347_v2_ else (d_1346_i0_) <= (self.f2)):
                d_1348_v3_: _dafny.Map
                d_1348_v3_ = _dafny.Map({d_1347_v2_: d_1347_v2_})
                d_1349_v4_: _dafny.Seq
                d_1349_v4_ = _dafny.SeqWithoutIsStrInference([d_1347_v2_, not(((d_1348_v3_)[d_1347_v2_] if (d_1347_v2_) in (d_1348_v3_) else d_1347_v2_)), d_1347_v2_])
                d_1348_v3_ = (d_1348_v3_).set(d_1347_v2_, (d_1349_v4_)[default__.safeIndex((0) - (self.f2), len(d_1349_v4_))])
                (self).f2 = p0
                d_1350_v5_: C0
                nw251_ = C0()
                nw251_.ctor__(d_1347_v2_, d_1347_v2_)
                d_1350_v5_ = nw251_
                d_1351_v6_: _dafny.Map
                d_1351_v6_ = _dafny.Map({d_1350_v5_.f5: (self).fm1(not(d_1350_v5_.f4), globalState)})
                (self).f2 = (0) - ((p0) + (len((d_1351_v6_) | (d_1351_v6_))))
                d_1352_v7_: _dafny.Array
                def lambda56_(d_1353_v4_):
                    def lambda57_(d_1354_i1_):
                        return default__.safeModuloInt(d_1354_i1_, len(d_1353_v4_))

                    return lambda57_

                init29_ = lambda56_(d_1349_v4_)
                nw252_ = _dafny.Array(None, 15)
                for i0_29_ in range(nw252_.length(0)):
                    nw252_[i0_29_] = init29_(i0_29_)
                d_1352_v7_ = nw252_
                d_1355_v8_: D4
                d_1355_v8_ = D4_DC6((0) - ((self).fm1(d_1350_v5_.f5, globalState)), d_1352_v7_, p0)
                index255_ = default__.safeIndex(674, (d_1352_v7_).length(0))
                (d_1352_v7_)[index255_] = (d_1355_v8_).cf9
                d_1356_v9_: _dafny.Array
                nw253_ = _dafny.Array(_dafny.Map({}), 26)
                d_1356_v9_ = nw253_
                index256_ = default__.safeIndex(790, (d_1356_v9_).length(0))
                (d_1356_v9_)[index256_] = d_1351_v6_
                d_1357_v10_: _dafny.Map
                d_1357_v10_ = _dafny.Map({d_1346_i0_: d_1350_v5_.f5})
                d_1358_v11_: _dafny.Map
                d_1358_v11_ = _dafny.Map({_dafny.Map({self.f2: d_1346_i0_}): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1359_i2_ in range(default__.abs(505))])})
                d_1360_v12_: str
                d_1360_v12_ = _dafny.CodePoint('g')
                d_1361_v13_: D7
                d_1361_v13_ = D7_DC15(self.f2, d_1360_v12_)
                d_1362_v14_: _dafny.Map
                d_1362_v14_ = _dafny.Map({len(d_1358_v11_): d_1361_v13_})
                index257_ = default__.safeIndex(674, (d_1352_v7_).length(0))
                index258_ = default__.safeIndex(790, (d_1356_v9_).length(0))
                def iife102_():
                    coll46_ = _dafny.Map()
                    compr_46_: int
                    for compr_46_ in _dafny.IntegerRange(183, 417):
                        d_1363_v15_: int = compr_46_
                        if ((183) <= (d_1363_v15_)) and ((d_1363_v15_) < (417)):
                            coll46_[(d_1363_v15_) + (p0)] = d_1361_v13_
                    return _dafny.Map(coll46_)
                rhs215_ = self.f2
                rhs216_ = ((len(d_1357_v10_)) - (-727) if d_1350_v5_.f5 else len((d_1362_v14_) | (iife102_()
                )))
                rhs217_ = default__.safeModuloInt(p0, self.f2)
                rhs218_ = ((d_1351_v6_) | (d_1351_v6_)) | (d_1351_v6_)
                rhs219_ = (d_1345_v1_) + ((d_1345_v1_) + (d_1345_v1_))
                lhs127_ = self
                lhs128_ = d_1352_v7_
                lhs129_ = default__.safeIndex(674, (d_1352_v7_).length(0))
                lhs130_ = d_1356_v9_
                lhs131_ = default__.safeIndex(790, (d_1356_v9_).length(0))
                r0 = rhs215_
                lhs127_.f2 = rhs216_
                lhs128_[lhs129_] = rhs217_
                lhs130_[lhs131_] = rhs218_
                d_1345_v1_ = rhs219_
            elif True:
                d_1364_v16_: _dafny.Array
                nw254_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_1364_v16_ = nw254_
                d_1365_v17_: _dafny.Array
                nw255_ = _dafny.Array(int(0), 17)
                d_1365_v17_ = nw255_
                index259_ = default__.safeIndex(389, (d_1364_v16_).length(0))
                (d_1364_v16_)[index259_] = d_1365_v17_
                index260_ = default__.safeIndex(389, (d_1364_v16_).length(0))
                (d_1364_v16_)[index260_] = d_1365_v17_
                d_1366_v18_: C1
                nw256_ = C1()
                nw256_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1367_i3_ in range(default__.abs(-92))]))
                d_1366_v18_ = nw256_
                d_1366_v18_ = d_1366_v18_
                d_1368_v19_: _dafny.Set
                d_1368_v19_ = _dafny.Set({d_1346_i0_})
                d_1369_v20_: _dafny.Seq
                d_1369_v20_ = _dafny.SeqWithoutIsStrInference([d_1368_v19_])
                d_1370_v21_: _dafny.Map
                d_1370_v21_ = _dafny.Map({(d_1369_v20_) == (d_1369_v20_): d_1347_v2_})
                d_1370_v21_ = (d_1370_v21_).set(not(d_1347_v2_), d_1347_v2_)
                d_1371_v22_: _dafny.Seq
                d_1371_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwqgaeod"))
                d_1371_v22_ = d_1371_v22_
                d_1372_v23_: _dafny.Seq
                d_1372_v23_ = _dafny.SeqWithoutIsStrInference([d_1347_v2_])
                index261_ = default__.safeIndex(885, (d_1365_v17_).length(0))
                (d_1365_v17_)[index261_] = len((_dafny.SeqWithoutIsStrInference([not(d_1347_v2_), d_1347_v2_])) + (d_1372_v23_))
                index262_ = default__.safeIndex(885, (d_1365_v17_).length(0))
                (d_1365_v17_)[index262_] = ((d_1346_i0_) * (867) if d_1347_v2_ else self.f2)
            d_1373_v24_: str
            d_1373_v24_ = _dafny.CodePoint('s')
            d_1373_v24_ = d_1373_v24_
            d_1374_v25_: _dafny.Seq
            d_1374_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ednpivdn"))
            d_1375_v26_: _dafny.Seq
            d_1375_v26_ = _dafny.SeqWithoutIsStrInference([d_1374_v25_, d_1374_v25_])
            d_1375_v26_ = d_1375_v26_
            (self).f2 = d_1346_i0_
        d_1376_v27_: _dafny.Seq
        d_1376_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wosyavxm"))
        d_1377_v28_: D6
        d_1377_v28_ = D6_DC11(d_1376_v27_)
        r0 = (((d_1345_v1_)[default__.safeIndex(p0, len(d_1345_v1_))]) - (p0)) - ((0) - (len((d_1377_v28_).cf21)))
        d_1378_v29_: bool
        d_1378_v29_ = True
        r0 = ((self.f2) + (self.f2) if d_1378_v29_ else self.f2)
        d_1379_v30_: _dafny.Set
        d_1379_v30_ = _dafny.Set({(d_1345_v1_)[default__.safeIndex(self.f2, len(d_1345_v1_))], -226, p0})
        d_1380_v31_: _dafny.Seq
        d_1380_v31_ = _dafny.SeqWithoutIsStrInference([d_1378_v29_])
        d_1381_v32_: _dafny.MultiSet
        d_1381_v32_ = _dafny.MultiSet([len(d_1380_v31_), self.f2])
        d_1382_v33_: _dafny.Set
        d_1382_v33_ = _dafny.Set({False, True})
        d_1383_v34_: _dafny.Map
        d_1383_v34_ = _dafny.Map({(d_1381_v32_).cardinality: d_1382_v33_})
        def iife103_():
            def iife108_():
                def iife110_():
                    coll54_ = _dafny.Map()
                    compr_54_: int
                    for compr_54_ in _dafny.IntegerRange(349, -74):
                        d_1384_v36_: int = compr_54_
                        if ((349) <= (d_1384_v36_)) and ((d_1384_v36_) < (-74)):
                            coll54_[default__.safeDivisionInt(d_1384_v36_, self.f2)] = (d_1381_v32_).cardinality
                    return _dafny.Map(coll54_)
                def iife111_():
                    coll55_ = _dafny.Map()
                    compr_55_: int
                    for compr_55_ in _dafny.IntegerRange(443, 563):
                        d_1386_v37_: int = compr_55_
                        if ((443) <= (d_1386_v37_)) and ((d_1386_v37_) < (563)):
                            coll55_[(d_1386_v37_) * (self.f2)] = p0
                    return _dafny.Map(coll55_)
                coll52_ = _dafny.Map()
                def iife109_():
                    coll53_ = _dafny.Map()
                    compr_53_: int
                    for compr_53_ in _dafny.IntegerRange(349, -74):
                        d_1384_v36_: int = compr_53_
                        if ((349) <= (d_1384_v36_)) and ((d_1384_v36_) < (-74)):
                            coll53_[default__.safeDivisionInt(d_1384_v36_, self.f2)] = (d_1381_v32_).cardinality
                    return _dafny.Map(coll53_)
                compr_52_: int
                for compr_52_ in (iife109_()
                ).keys.Elements:
                    d_1385_v35_: int = compr_52_
                    if (d_1385_v35_) in (iife110_()
                    ):
                        coll52_[default__.safeDivisionInt(d_1385_v35_, len(_dafny.Map({_dafny.CodePoint('m'): len(iife111_()
                        )})))] = d_1380_v31_
                return _dafny.Map(coll52_)
            coll47_ = _dafny.Set()
            def iife104_():
                def iife106_():
                    coll50_ = _dafny.Map()
                    compr_50_: int
                    for compr_50_ in _dafny.IntegerRange(349, -74):
                        d_1384_v36_: int = compr_50_
                        if ((349) <= (d_1384_v36_)) and ((d_1384_v36_) < (-74)):
                            coll50_[default__.safeDivisionInt(d_1384_v36_, self.f2)] = (d_1381_v32_).cardinality
                    return _dafny.Map(coll50_)
                def iife107_():
                    coll51_ = _dafny.Map()
                    compr_51_: int
                    for compr_51_ in _dafny.IntegerRange(443, 563):
                        d_1386_v37_: int = compr_51_
                        if ((443) <= (d_1386_v37_)) and ((d_1386_v37_) < (563)):
                            coll51_[(d_1386_v37_) * (self.f2)] = p0
                    return _dafny.Map(coll51_)
                coll48_ = _dafny.Map()
                def iife105_():
                    coll49_ = _dafny.Map()
                    compr_49_: int
                    for compr_49_ in _dafny.IntegerRange(349, -74):
                        d_1384_v36_: int = compr_49_
                        if ((349) <= (d_1384_v36_)) and ((d_1384_v36_) < (-74)):
                            coll49_[default__.safeDivisionInt(d_1384_v36_, self.f2)] = (d_1381_v32_).cardinality
                    return _dafny.Map(coll49_)
                compr_48_: int
                for compr_48_ in (iife105_()
                ).keys.Elements:
                    d_1385_v35_: int = compr_48_
                    if (d_1385_v35_) in (iife106_()
                    ):
                        coll48_[default__.safeDivisionInt(d_1385_v35_, len(_dafny.Map({_dafny.CodePoint('m'): len(iife107_()
                        )})))] = d_1380_v31_
                return _dafny.Map(coll48_)
            compr_47_: int
            for compr_47_ in (default__.fm31(d_1378_v29_, len(d_1383_v34_), d_1378_v29_, len(iife104_()
            ), globalState)).Elements:
                d_1387_v38_: int = compr_47_
                if (d_1387_v38_) in (default__.fm31(d_1378_v29_, len(d_1383_v34_), d_1378_v29_, len(iife108_()
                ), globalState)):
                    coll47_ = coll47_.union(_dafny.Set([(d_1387_v38_) - (len(_dafny.Map({False: len(_dafny.Map({656: len(_dafny.Set({True, True}))}))})))]))
            return _dafny.Set(coll47_)
        d_1379_v30_ = ((d_1379_v30_) - (iife103_()
        )) - ((_dafny.Set({len(d_1376_v27_)}) if d_1378_v29_ else d_1379_v30_))
        d_1383_v34_ = (d_1383_v34_).set(self.f2, d_1382_v33_)
        d_1388_v39_: C1
        nw257_ = C1()
        nw257_.ctor__(d_1376_v27_)
        d_1388_v39_ = nw257_
        d_1389_v40_: D8
        d_1389_v40_ = D8_DC16(d_1388_v39_)
        d_1390_v41_: _dafny.Map
        d_1390_v41_ = _dafny.Map({d_1389_v40_: p0})
        r0 = ((d_1390_v41_)[d_1389_v40_] if (d_1389_v40_) in (d_1390_v41_) else p0)
        r1 = (((_dafny.MultiSet([d_1378_v29_, True])).cardinality) * ((0) - (p0))) <= (self.f2)
        r2 = d_1378_v29_
        return r0, r1, r2


class C4(T0, T1, T2):
    def  __init__(self):
        self._f2: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f2(self):
        return self._f2
    @f2.setter
    def f2(self, value):
        self._f2 = value
    def ctor__(self, f2):
        (self).f2 = f2

    def fm0(self, globalState):
        source20_ = D0_DC1(False, True, _dafny.CodePoint('l'), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mephw"))))
        if source20_.is_DC1:
            d_1391___mcc_h0_ = source20_.cf1
            d_1392___mcc_h1_ = source20_.cf2
            d_1393___mcc_h2_ = source20_.cf3
            d_1394___mcc_h3_ = source20_.cf4
            d_1395_cf4_ = d_1394___mcc_h3_
            d_1396_cf3_ = d_1393___mcc_h2_
            d_1397_cf2_ = d_1392___mcc_h1_
            d_1398_cf1_ = d_1391___mcc_h0_
            return (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_1395_cf4_, len(_dafny.SeqWithoutIsStrInference([d_1396_cf3_ for d_1399_i0_ in range(default__.abs(3))])), self.f2, d_1395_cf4_, d_1395_cf4_]): len(_dafny.SeqWithoutIsStrInference([self.f2 for d_1400_i1_ in range(default__.abs(837))]))})) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([d_1395_cf4_, d_1395_cf4_]): self.f2})))
        elif True:
            d_1401___mcc_h4_ = source20_.cf0
            d_1402_cf0_ = d_1401___mcc_h4_
            def iife112_():
                coll56_ = _dafny.Map()
                compr_56_: _dafny.Seq
                for compr_56_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([self.f2 for d_1403_i2_ in range(default__.abs(530))]), _dafny.SeqWithoutIsStrInference([self.f2, self.f2, (_dafny.MultiSet([self.f2, len(_dafny.Map({self.f2: (0) - (self.f2)}))])).cardinality])})).Elements:
                    d_1404_v0_: _dafny.Seq = compr_56_
                    if (d_1404_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([self.f2 for d_1403_i2_ in range(default__.abs(530))]), _dafny.SeqWithoutIsStrInference([self.f2, self.f2, (_dafny.MultiSet([self.f2, len(_dafny.Map({self.f2: (0) - (self.f2)}))])).cardinality])})):
                        coll56_[d_1404_v0_] = 802
                return _dafny.Map(coll56_)
            return iife112_()
            

    def fm1(self, p0, globalState):
        return len(_dafny.Set({default__.safeModuloInt(929, self.f2)}))

    def fm8(self, p0, p1, globalState):
        return ((_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")))})) | (_dafny.Set({self.f2, self.f2}))).isdisjoint(_dafny.Set({self.f2, self.f2}))

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        d_1405_v0_: _dafny.Map
        d_1405_v0_ = _dafny.Map({True: not(p2)})
        d_1406_v1_: _dafny.Map
        d_1406_v1_ = d_1405_v0_
        source21_ = d_1406_v1_
        d_1407___mcc_h0_ = source21_
        d_1408_cf5_ = d_1407___mcc_h0_
        d_1409_v2_: str
        d_1409_v2_ = _dafny.CodePoint('s')
        d_1410_v3_: _dafny.Map
        d_1410_v3_ = _dafny.Map({d_1409_v2_: (self.f2) + (self.f2)})
        d_1410_v3_ = (d_1410_v3_).set(d_1409_v2_, p0)
        d_1411_v4_: _dafny.Set
        d_1411_v4_ = _dafny.Set({p0, len(default__.fm9(p2, p2, p2, self.f2, globalState))})
        d_1411_v4_ = d_1411_v4_
        (self).f2 = p0
        d_1412_v5_: C2
        nw258_ = C2()
        nw258_.ctor__(self.f2)
        d_1412_v5_ = nw258_
        d_1413_v6_: _dafny.Seq
        d_1413_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nuymyuim"))
        d_1414_v7_: _dafny.Map
        d_1414_v7_ = _dafny.Map({p0: p2})
        d_1415_v8_: D5
        d_1415_v8_ = D5_DC10(p3, len(d_1413_v6_), _dafny.Map({d_1413_v6_: d_1414_v7_}), p3)
        d_1416_v9_: _dafny.Map
        d_1416_v9_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1415_v8_, d_1415_v8_]): p3})
        d_1417_v10_: T0
        nw259_ = C2()
        nw259_.ctor__((self).fm1(p3, globalState))
        d_1417_v10_ = nw259_
        d_1418_v11_: _dafny.MultiSet
        d_1418_v11_ = _dafny.MultiSet([d_1417_v10_])
        d_1419_v12_: _dafny.Array
        nw260_ = _dafny.Array(None, 4)
        nw260_[int(0)] = p3
        nw260_[int(1)] = (len(d_1416_v9_)) != ((d_1418_v11_).cardinality)
        nw260_[int(2)] = p2
        nw260_[int(3)] = False
        d_1419_v12_ = nw260_
        d_1419_v12_ = d_1419_v12_
        d_1420_v13_: C2
        nw261_ = C2()
        nw261_.ctor__(p0)
        d_1420_v13_ = nw261_
        d_1421_v14_: _dafny.Array
        nw262_ = _dafny.Array(None, 6)
        nw262_[int(0)] = p0
        nw262_[int(1)] = len((d_1413_v6_).set(default__.safeIndex(p0, len(d_1413_v6_)), (d_1413_v6_)[default__.safeIndex(self.f2, len(d_1413_v6_))]))
        nw262_[int(2)] = self.f2
        nw262_[int(3)] = (0) - (self.f2)
        nw262_[int(4)] = self.f2
        nw262_[int(5)] = self.f2
        d_1421_v14_ = nw262_
        d_1422_v15_: _dafny.Seq
        d_1422_v15_ = _dafny.SeqWithoutIsStrInference([self.f2, p0])
        rhs220_ = d_1421_v14_
        rhs221_ = p3
        rhs222_ = (p0) + ((d_1422_v15_)[default__.safeIndex(self.f2, len(d_1422_v15_))])
        d_1421_v14_ = rhs220_
        r0 = rhs221_
        r2 = rhs222_
        d_1423_v16_: str
        d_1423_v16_ = _dafny.CodePoint('r')
        if not((d_1423_v16_) in (d_1413_v6_)):
            d_1424_v17_: C1
            nw263_ = C1()
            nw263_.ctor__(d_1413_v6_)
            d_1424_v17_ = nw263_
            pat_let_tv29_ = d_1424_v17_
            def iife113_(_pat_let28_0):
                def iife114_(d_1425_dt__update__tmp_h0_):
                    def iife115_(_pat_let29_0):
                        def iife116_(d_1426_dt__update_hcf29_h0_):
                            return D8_DC16(d_1426_dt__update_hcf29_h0_)
                        return iife116_(_pat_let29_0)
                    return iife115_(pat_let_tv29_)
                return iife114_(_pat_let28_0)
            source22_ = iife113_(D8_DC16(d_1424_v17_))
            if source22_.is_DC17:
                d_1427___mcc_h1_ = source22_.cf30
                d_1428___mcc_h2_ = source22_.cf31
                d_1429___mcc_h3_ = source22_.cf32
                d_1430_cf32_ = d_1429___mcc_h3_
                d_1431_cf31_ = d_1428___mcc_h2_
                d_1432_cf30_ = d_1427___mcc_h1_
                d_1430_cf32_ = (p0) * ((self.f2) * (d_1430_cf32_))
                d_1433_v18_: C3
                nw264_ = C3()
                nw264_.ctor__(p0)
                d_1433_v18_ = nw264_
                (self).f2 = default__.safeDivisionInt(p0, self.f2)
                d_1434_v19_: C3
                nw265_ = C3()
                nw265_.ctor__(d_1430_cf32_)
                d_1434_v19_ = nw265_
            elif True:
                d_1435___mcc_h4_ = source22_.cf29
                d_1436_cf29_ = d_1435___mcc_h4_
                d_1437_v20_: _dafny.Map
                d_1437_v20_ = _dafny.Map({self.f2: (d_1424_v17_).fm1(p2, globalState)})
                rhs223_ = (d_1437_v20_) | ((d_1437_v20_) | (d_1437_v20_))
                rhs224_ = d_1419_v12_
                rhs225_ = (d_1414_v7_) | (d_1414_v7_)
                rhs226_ = (self.f2) - (default__.safeDivisionInt(self.f2, self.f2))
                d_1437_v20_ = rhs223_
                d_1419_v12_ = rhs224_
                d_1414_v7_ = rhs225_
                r2 = rhs226_
                r2 = self.f2
                (self).f2 = self.f2
                r2 = (p0) + (self.f2)
            (self).f2 = (0) - (self.f2)
            d_1438_v21_: D6
            d_1438_v21_ = D6_DC12(p2, p3)
            d_1439_v22_: _dafny.Map
            d_1439_v22_ = _dafny.Map({(d_1438_v21_).cf23: self.f2})
            rhs227_ = (d_1413_v6_) <= (_dafny.SeqWithoutIsStrInference([d_1423_v16_ for d_1440_i0_ in range(default__.abs(940))]))
            rhs228_ = default__.safeModuloInt(self.f2, ((d_1439_v22_)[p3] if (p3) in (d_1439_v22_) else self.f2))
            r0 = rhs227_
            r2 = rhs228_
            d_1441_v23_: _dafny.Array
            nw266_ = _dafny.Array(None, 10)
            nw266_[int(0)] = d_1421_v14_
            nw266_[int(1)] = (d_1421_v14_ if p3 else d_1421_v14_)
            nw266_[int(2)] = d_1421_v14_
            nw266_[int(3)] = d_1421_v14_
            nw266_[int(4)] = d_1421_v14_
            nw266_[int(5)] = (d_1421_v14_ if p2 else d_1421_v14_)
            nw266_[int(6)] = d_1421_v14_
            nw266_[int(7)] = d_1421_v14_
            nw266_[int(8)] = d_1421_v14_
            nw266_[int(9)] = (d_1421_v14_ if p2 else d_1421_v14_)
            d_1441_v23_ = nw266_
            d_1441_v23_ = d_1441_v23_
            index263_ = default__.safeIndex(681, (d_1421_v14_).length(0))
            (d_1421_v14_)[index263_] = default__.safeDivisionInt(p0, len(d_1422_v15_))
            index264_ = default__.safeIndex(681, (d_1421_v14_).length(0))
            (d_1421_v14_)[index264_] = self.f2
        elif True:
            d_1442_v24_: _dafny.Map
            d_1442_v24_ = _dafny.Map({(0) - ((0) - (self.f2)): p0})
            d_1442_v24_ = (d_1442_v24_).set(p0, (p0) * (-273))
            (self).f2 = (d_1420_v13_).fm1(p2, globalState)
            d_1443_v25_: _dafny.Seq
            d_1443_v25_ = _dafny.SeqWithoutIsStrInference([p2])
            d_1444_v26_: _dafny.Seq
            d_1444_v26_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p3])])
            d_1445_v27_: _dafny.Seq
            d_1445_v27_ = _dafny.SeqWithoutIsStrInference([(d_1443_v25_) + ((d_1444_v26_)[default__.safeIndex(p0, len(d_1444_v26_))]), d_1443_v25_, (d_1443_v25_) + (d_1443_v25_), d_1443_v25_, _dafny.SeqWithoutIsStrInference([p3, p2, p3])])
            d_1444_v26_ = (d_1445_v27_) + (_dafny.SeqWithoutIsStrInference([d_1443_v25_ for d_1446_i1_ in range(default__.abs(384))]))
            d_1447_v28_: _dafny.Set
            d_1447_v28_ = _dafny.Set({p2})
            rhs229_ = (d_1422_v15_) + (_dafny.SeqWithoutIsStrInference([self.f2, self.f2]))
            rhs230_ = (len((d_1447_v28_).intersection(d_1447_v28_)) if p2 else 967)
            rhs231_ = p2
            rhs232_ = -227
            d_1422_v15_ = rhs229_
            r2 = rhs230_
            r0 = rhs231_
            r2 = rhs232_
            d_1448_v29_: _dafny.Map
            d_1448_v29_ = _dafny.Map({d_1421_v14_: p0})
            d_1448_v29_ = (d_1448_v29_).set(d_1421_v14_, self.f2)
        d_1449_v30_: _dafny.Seq
        d_1449_v30_ = _dafny.SeqWithoutIsStrInference([p2])
        d_1450_v31_: _dafny.Set
        d_1450_v31_ = _dafny.Set({p3, p2, p2, not(p2), p2})
        d_1413_v6_ = (default__.fm28((d_1449_v30_)[default__.safeIndex(756, len(d_1449_v30_))], p3, p2, d_1450_v31_, globalState)) + ((default__.fm28(True, not(p3), p2, default__.fm41(d_1422_v15_, p3, True, p2, globalState), globalState)).set(default__.safeIndex(p0, len(default__.fm28(True, not(p3), p2, default__.fm41(d_1422_v15_, p3, True, p2, globalState), globalState))), d_1423_v16_))
        r0 = p3
        nw267_ = _dafny.Array(_dafny.Map({}), 2)
        r1 = nw267_
        r2 = self.f2
        return r0, r1, r2

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_1451_v0_: _dafny.Seq
        d_1451_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxrn"))
        d_1452_v1_: _dafny.Seq
        d_1452_v1_ = _dafny.SeqWithoutIsStrInference([p2])
        d_1453_v2_: _dafny.MultiSet
        d_1453_v2_ = _dafny.MultiSet([(0) - (-166)])
        d_1454_v3_: _dafny.Seq
        d_1454_v3_ = _dafny.SeqWithoutIsStrInference([d_1453_v2_])
        d_1455_v4_: str
        d_1455_v4_ = _dafny.CodePoint('j')
        d_1456_v5_: _dafny.Array
        nw268_ = _dafny.Array(None, 25)
        nw268_[int(0)] = (p2) == (p2)
        nw268_[int(1)] = p2
        nw268_[int(2)] = (p0) < (p1)
        nw268_[int(3)] = p2
        nw268_[int(4)] = not(p2)
        nw268_[int(5)] = not(p2)
        nw268_[int(6)] = p2
        nw268_[int(7)] = (p2 if p2 else p2)
        nw268_[int(8)] = p2
        nw268_[int(9)] = p2
        nw268_[int(10)] = p2
        nw268_[int(11)] = (d_1451_v0_) != (d_1451_v0_)
        nw268_[int(12)] = p2
        nw268_[int(13)] = p2
        nw268_[int(14)] = (d_1452_v1_) != (d_1452_v1_)
        nw268_[int(15)] = (_dafny.SeqWithoutIsStrInference([d_1453_v2_])) <= (d_1454_v3_)
        nw268_[int(16)] = p2
        nw268_[int(17)] = p2
        nw268_[int(18)] = p2
        nw268_[int(19)] = p2
        nw268_[int(20)] = p2
        nw268_[int(21)] = p2
        nw268_[int(22)] = (d_1455_v4_) in (d_1451_v0_)
        nw268_[int(23)] = p2
        nw268_[int(24)] = p2
        d_1456_v5_ = nw268_
        d_1457_v6_: _dafny.Set
        d_1457_v6_ = _dafny.Set({p1, p1, p1})
        d_1458_v7_: _dafny.Map
        d_1458_v7_ = _dafny.Map({d_1455_v4_: d_1457_v6_})
        d_1459_v8_: _dafny.Seq
        d_1459_v8_ = _dafny.SeqWithoutIsStrInference([d_1457_v6_])
        index265_ = default__.safeIndex(507, (d_1456_v5_).length(0))
        (d_1456_v5_)[index265_] = (_dafny.SeqWithoutIsStrInference([((d_1458_v7_)[d_1455_v4_] if (d_1455_v4_) in (d_1458_v7_) else _dafny.Set({len(d_1451_v0_), len(d_1452_v1_)})), _dafny.Set({self.f2})])) <= (d_1459_v8_)
        d_1460_v9_: C2
        nw269_ = C2()
        nw269_.ctor__(p1)
        d_1460_v9_ = nw269_
        d_1461_v10_: _dafny.Map
        d_1461_v10_ = _dafny.Map({p2: d_1460_v9_})
        d_1462_v11_: _dafny.Seq
        d_1462_v11_ = _dafny.SeqWithoutIsStrInference([p0])
        index266_ = default__.safeIndex(507, (d_1456_v5_).length(0))
        (d_1456_v5_)[index266_] = (self).fm8(d_1451_v0_, not(not((len(d_1461_v10_)) <= (len(d_1462_v11_)))), globalState)
        d_1463_v12_: _dafny.Array
        nw270_ = _dafny.Array(int(0), 8)
        d_1463_v12_ = nw270_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1463_v12_).length(0)):
            d_1464_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_1464_i0_)) and ((d_1464_i0_) < ((d_1463_v12_).length(0)))):
                (d_1463_v12_)[(d_1464_i0_)] = default__.safeModuloInt(d_1464_i0_, p1)
        r0 = p1
        hi7_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsihpmbty")))
        for d_1465_i1_ in range(p0, hi7_):
            d_1466_v13_: _dafny.Map
            d_1466_v13_ = _dafny.Map({d_1460_v9_: d_1465_i1_})
            d_1466_v13_ = (d_1466_v13_) | ((d_1466_v13_) | (d_1466_v13_))
            d_1467_v14_: C1
            nw271_ = C1()
            nw271_.ctor__(d_1451_v0_)
            d_1467_v14_ = nw271_
            d_1468_v15_: D8
            d_1468_v15_ = D8_DC17(False, d_1467_v14_, d_1465_i1_)
            d_1469_v16_: _dafny.Map
            d_1469_v16_ = _dafny.Map({(d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))]: d_1455_v4_})
            pat_let_tv30_ = d_1467_v14_
            pat_let_tv31_ = d_1469_v16_
            pat_let_tv32_ = p1
            def iife117_(_pat_let30_0):
                def iife118_(d_1470_dt__update__tmp_h0_):
                    def iife119_(_pat_let31_0):
                        def iife120_(d_1471_dt__update_hcf31_h0_):
                            def iife121_(_pat_let32_0):
                                def iife122_(d_1472_dt__update_hcf30_h0_):
                                    return D8_DC17(d_1472_dt__update_hcf30_h0_, d_1471_dt__update_hcf31_h0_, (d_1470_dt__update__tmp_h0_).cf32)
                                return iife122_(_pat_let32_0)
                            return iife121_((len(pat_let_tv31_)) != (pat_let_tv32_))
                        return iife120_(_pat_let31_0)
                    return iife119_(pat_let_tv30_)
                return iife118_(_pat_let30_0)
            source23_ = iife117_(d_1468_v15_)
            if source23_.is_DC17:
                d_1473___mcc_h0_ = source23_.cf30
                d_1474___mcc_h1_ = source23_.cf31
                d_1475___mcc_h2_ = source23_.cf32
                d_1476_cf32_ = d_1475___mcc_h2_
                d_1477_cf31_ = d_1474___mcc_h1_
                d_1478_cf30_ = d_1473___mcc_h0_
                rhs233_ = d_1476_cf32_
                rhs234_ = default__.safeDivisionInt(self.f2, 401)
                r0 = rhs233_
                d_1476_cf32_ = rhs234_
                d_1476_cf32_ = (0) - (default__.safeDivisionInt(default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqghm"))), 104), default__.safeModuloInt(d_1476_cf32_, d_1465_i1_)))
                d_1479_v17_: D11
                d_1479_v17_ = D11_DC20(d_1456_v5_)
                d_1480_v18_: _dafny.Seq
                d_1480_v18_ = _dafny.SeqWithoutIsStrInference([d_1456_v5_])
                d_1481_v19_: _dafny.Array
                nw272_ = _dafny.Array(None, 22)
                nw272_[int(0)] = d_1456_v5_
                nw272_[int(1)] = (d_1479_v17_).cf35
                nw272_[int(2)] = d_1456_v5_
                nw272_[int(3)] = (d_1456_v5_ if d_1478_cf30_ else d_1456_v5_)
                nw272_[int(4)] = d_1456_v5_
                nw272_[int(5)] = d_1456_v5_
                nw272_[int(6)] = d_1456_v5_
                nw272_[int(7)] = d_1456_v5_
                nw272_[int(8)] = d_1456_v5_
                nw272_[int(9)] = d_1456_v5_
                nw272_[int(10)] = d_1456_v5_
                nw272_[int(11)] = d_1456_v5_
                nw272_[int(12)] = d_1456_v5_
                nw272_[int(13)] = d_1456_v5_
                nw272_[int(14)] = d_1456_v5_
                nw272_[int(15)] = d_1456_v5_
                nw272_[int(16)] = d_1456_v5_
                nw272_[int(17)] = d_1456_v5_
                nw272_[int(18)] = d_1456_v5_
                nw272_[int(19)] = d_1456_v5_
                nw272_[int(20)] = (d_1480_v18_)[default__.safeIndex((d_1453_v2_).cardinality, len(d_1480_v18_))]
                nw272_[int(21)] = d_1456_v5_
                d_1481_v19_ = nw272_
                index267_ = default__.safeIndex(961, (d_1481_v19_).length(0))
                (d_1481_v19_)[index267_] = d_1456_v5_
                d_1482_v20_: _dafny.Set
                d_1482_v20_ = _dafny.Set({d_1478_cf30_, (d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))]})
                d_1483_v21_: _dafny.MultiSet
                d_1483_v21_ = _dafny.MultiSet([d_1482_v20_])
                index268_ = default__.safeIndex(961, (d_1481_v19_).length(0))
                nw273_ = _dafny.Array(None, 11)
                nw273_[int(0)] = (d_1478_cf30_ if p2 else p2)
                nw273_[int(1)] = not((d_1483_v21_).ispropersubset(d_1483_v21_))
                nw273_[int(2)] = (d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))]
                nw273_[int(3)] = d_1478_cf30_
                nw273_[int(4)] = (d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))]
                nw273_[int(5)] = not((d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))])
                nw273_[int(6)] = True
                nw273_[int(7)] = d_1478_cf30_
                nw273_[int(8)] = (d_1453_v2_).issubset(d_1453_v2_)
                nw273_[int(9)] = False
                nw273_[int(10)] = p2
                (d_1481_v19_)[index268_] = nw273_
                d_1478_cf30_ = (((d_1453_v2_)[p1] if (p1) in (d_1453_v2_) else d_1465_i1_)) != (p1)
            elif True:
                d_1484___mcc_h3_ = source23_.cf29
                d_1485_cf29_ = d_1484___mcc_h3_
                d_1486_v22_: D6
                d_1486_v22_ = D6_DC13((d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))], 763)
                d_1487_v23_: _dafny.Map
                d_1487_v23_ = _dafny.Map({d_1486_v22_: d_1456_v5_})
                d_1456_v5_ = ((d_1487_v23_)[d_1486_v22_] if (d_1486_v22_) in (d_1487_v23_) else d_1456_v5_)
                r0 = p0
                def iife123_():
                    coll57_ = _dafny.Set()
                    compr_57_: int
                    for compr_57_ in _dafny.IntegerRange(96, -363):
                        d_1488_v24_: int = compr_57_
                        if ((96) <= (d_1488_v24_)) and ((d_1488_v24_) < (-363)):
                            coll57_ = coll57_.union(_dafny.Set([(d_1488_v24_) + (d_1465_i1_)]))
                    return _dafny.Set(coll57_)
                d_1457_v6_ = iife123_()
                
                (self).f2 = p1
            d_1451_v0_ = d_1451_v0_
            if (d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))]:
                index269_ = default__.safeIndex(507, (d_1456_v5_).length(0))
                (d_1456_v5_)[index269_] = p2
                d_1489_v25_: C1
                nw274_ = C1()
                nw274_.ctor__(_dafny.SeqWithoutIsStrInference([d_1455_v4_ for d_1490_i2_ in range(default__.abs(925))]))
                d_1489_v25_ = nw274_
                d_1491_v26_: _dafny.Seq
                d_1491_v26_ = _dafny.SeqWithoutIsStrInference([d_1456_v5_])
                d_1491_v26_ = d_1491_v26_
                d_1492_v27_: _dafny.Map
                d_1492_v27_ = _dafny.Map({p2: 542})
                d_1493_v28_: _dafny.Map
                d_1493_v28_ = _dafny.Map({d_1492_v27_: (d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))]})
                d_1494_v29_: _dafny.Map
                d_1494_v29_ = _dafny.Map({p2: d_1463_v12_})
                d_1495_v30_: _dafny.Array
                nw275_ = _dafny.Array(None, 24)
                nw275_[int(0)] = d_1462_v11_
                nw275_[int(1)] = d_1462_v11_
                nw275_[int(2)] = d_1462_v11_
                nw275_[int(3)] = d_1462_v11_
                nw275_[int(4)] = d_1462_v11_
                nw275_[int(5)] = (_dafny.SeqWithoutIsStrInference([p0, len(d_1452_v1_), p0])).set(default__.safeIndex(len(d_1493_v28_), len(_dafny.SeqWithoutIsStrInference([p0, len(d_1452_v1_), p0]))), -996)
                nw275_[int(6)] = _dafny.SeqWithoutIsStrInference([715])
                nw275_[int(7)] = (d_1467_v14_).fm12((d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))], (d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))], len(d_1494_v29_), (d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))], globalState)
                nw275_[int(8)] = d_1462_v11_
                nw275_[int(9)] = d_1462_v11_
                nw275_[int(10)] = d_1462_v11_
                nw275_[int(11)] = _dafny.SeqWithoutIsStrInference([d_1465_i1_])
                nw275_[int(12)] = d_1462_v11_
                nw275_[int(13)] = _dafny.SeqWithoutIsStrInference([p1, len((d_1467_v14_).f3)])
                nw275_[int(14)] = _dafny.SeqWithoutIsStrInference([451 for d_1496_i3_ in range(default__.abs(885))])
                nw275_[int(15)] = d_1462_v11_
                nw275_[int(16)] = d_1462_v11_
                nw275_[int(17)] = d_1462_v11_
                nw275_[int(18)] = _dafny.SeqWithoutIsStrInference([p0 for d_1497_i4_ in range(default__.abs(20))])
                nw275_[int(19)] = d_1462_v11_
                nw275_[int(20)] = d_1462_v11_
                nw275_[int(21)] = d_1462_v11_
                nw275_[int(22)] = d_1462_v11_
                nw275_[int(23)] = d_1462_v11_
                d_1495_v30_ = nw275_
                d_1498_v31_: bool
                d_1499_v32_: _dafny.Array
                d_1500_v33_: int
                out23_: bool
                out24_: _dafny.Array
                out25_: int
                out23_, out24_, out25_ = (d_1489_v25_).m1(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ck"))), d_1495_v30_, not(not((d_1465_i1_) > (self.f2))), False, globalState)
                d_1498_v31_ = out23_
                d_1499_v32_ = out24_
                d_1500_v33_ = out25_
                r0 = d_1465_i1_
            elif True:
                d_1501_v34_: _dafny.Set
                d_1501_v34_ = _dafny.Set({(d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))], p2})
                d_1502_v35_: D4
                d_1502_v35_ = D4_DC5((d_1501_v34_ if True else d_1501_v34_))
                d_1503_v38_: _dafny.Map
                def iife124_():
                    coll58_ = _dafny.Map()
                    compr_58_: int
                    for compr_58_ in (d_1457_v6_).Elements:
                        d_1504_v37_: int = compr_58_
                        if (d_1504_v37_) in (d_1457_v6_):
                            coll58_[(d_1504_v37_) + (445)] = (d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))]
                    return _dafny.Map(coll58_)
                d_1503_v38_ = _dafny.Map({iife124_()
                : self.f2})
                d_1505_v39_: _dafny.Map
                d_1505_v39_ = _dafny.Map({self.f2: (d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))]})
                index270_ = default__.safeIndex(507, (d_1456_v5_).length(0))
                def iife125_():
                    coll59_ = _dafny.Map()
                    compr_59_: _dafny.Map
                    for compr_59_ in (((d_1503_v38_).set(d_1505_v39_, p1)) | (d_1503_v38_)).keys.Elements:
                        d_1506_v36_: _dafny.Map = compr_59_
                        if (d_1506_v36_) in (((d_1503_v38_).set(d_1505_v39_, p1)) | (d_1503_v38_)):
                            coll59_[d_1506_v36_] = default__.safeDivisionInt(self.f2, p0)
                    return _dafny.Map(coll59_)
                rhs235_ = (d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))]
                rhs236_ = len(iife125_()
                )
                rhs237_ = d_1502_v35_
                rhs238_ = d_1455_v4_
                lhs132_ = d_1456_v5_
                lhs133_ = default__.safeIndex(507, (d_1456_v5_).length(0))
                lhs132_[lhs133_] = rhs235_
                r0 = rhs236_
                d_1502_v35_ = rhs237_
                d_1455_v4_ = rhs238_
                d_1462_v11_ = ((d_1462_v11_).set(default__.safeIndex(p0, len(d_1462_v11_)), p1)) + (d_1462_v11_)
                (self).f2 = ((d_1460_v9_).fm1((d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))], globalState)) * (p1)
                r0 = self.f2
                index271_ = default__.safeIndex(507, (d_1456_v5_).length(0))
                (d_1456_v5_)[index271_] = p2
        d_1507_v40_: _dafny.MultiSet
        d_1507_v40_ = _dafny.MultiSet([p2, p2])
        d_1507_v40_ = (d_1507_v40_) | (d_1507_v40_)
        index272_ = default__.safeIndex(507, (d_1456_v5_).length(0))
        (d_1456_v5_)[index272_] = (self).fm8((d_1451_v0_) + (d_1451_v0_), (d_1456_v5_)[default__.safeIndex(507, (d_1456_v5_).length(0))], globalState)
        d_1508_v41_: _dafny.Set
        d_1508_v41_ = _dafny.Set({d_1456_v5_, d_1456_v5_})
        r0 = default__.safeDivisionInt(self.f2, default__.safeModuloInt(len(d_1508_v41_), len(_dafny.SeqWithoutIsStrInference([self.f2]))))
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_1509_v0_: _dafny.Map
        d_1509_v0_ = _dafny.Map({p2: not(p3)})
        d_1509_v0_ = (d_1509_v0_).set(p3, p3)
        d_1510_i0_: int
        d_1510_i0_ = 0
        with _dafny.label("6"):
            while p3:
                with _dafny.c_label("6"):
                    if (d_1510_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1510_i0_ = (d_1510_i0_) + (1)
                    d_1511_v1_: str
                    d_1511_v1_ = _dafny.CodePoint('q')
                    d_1512_v2_: _dafny.Seq
                    d_1512_v2_ = _dafny.SeqWithoutIsStrInference([(391) - (438), (self.f2) * (self.f2), -19, (self.f2) + (self.f2)])
                    d_1513_v3_: _dafny.Map
                    d_1513_v3_ = _dafny.Map({p3: d_1511_v1_})
                    d_1514_v4_: _dafny.Seq
                    d_1514_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwftu"))
                    rhs239_ = ((d_1513_v3_)[(self).fm8(d_1514_v4_, True, globalState)] if ((self).fm8(d_1514_v4_, True, globalState)) in (d_1513_v3_) else _dafny.CodePoint('k'))
                    rhs240_ = (d_1512_v2_) + (d_1512_v2_)
                    rhs241_ = (self).fm1((d_1514_v4_) <= (d_1514_v4_), globalState)
                    d_1511_v1_ = rhs239_
                    d_1512_v2_ = rhs240_
                    r0 = rhs241_
                    d_1515_v5_: bool
                    d_1515_v5_ = True
                    d_1515_v5_ = p3
                    d_1516_v6_: _dafny.Array
                    out26_: _dafny.Array
                    out26_ = default__.m0(p2, d_1515_v5_, globalState)
                    d_1516_v6_ = out26_
                    d_1512_v2_ = d_1512_v2_
                    pass
            pass
        if p3:
            d_1517_v7_: bool
            d_1517_v7_ = True
            d_1518_v8_: _dafny.Seq
            d_1518_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yauevfig"))
            d_1517_v7_ = (self).fm8(d_1518_v8_, (p3 if d_1517_v7_ else p2), globalState)
            r1 = default__.safeModuloInt(self.f2, len(d_1518_v8_))
            d_1519_v9_: str
            d_1519_v9_ = _dafny.CodePoint('a')
            d_1520_v10_: _dafny.Array
            def lambda58_(d_1521_i1_):
                return D7_DC14(_dafny.SeqWithoutIsStrInference([self.f2, self.f2, len(_dafny.Set({self.f2, self.f2})), (0) - (self.f2), self.f2]))

            init30_ = lambda58_
            nw276_ = _dafny.Array(None, 10)
            for i0_30_ in range(nw276_.length(0)):
                nw276_[i0_30_] = init30_(i0_30_)
            d_1520_v10_ = nw276_
            d_1522_v11_: _dafny.Map
            d_1522_v11_ = _dafny.Map({d_1519_v9_: d_1520_v10_})
            rhs242_ = (d_1522_v11_) | (d_1522_v11_)
            rhs243_ = self.f2
            rhs244_ = p2
            rhs245_ = (0) - ((0) - (self.f2))
            d_1522_v11_ = rhs242_
            r0 = rhs243_
            d_1517_v7_ = rhs244_
            r1 = rhs245_
            d_1523_v12_: _dafny.Array
            nw277_ = _dafny.Array(False, 27)
            d_1523_v12_ = nw277_
            index273_ = default__.safeIndex(67, (d_1523_v12_).length(0))
            (d_1523_v12_)[index273_] = p2
            d_1524_v13_: T0
            nw278_ = C2()
            nw278_.ctor__(693)
            d_1524_v13_ = nw278_
            d_1525_v14_: _dafny.MultiSet
            d_1525_v14_ = _dafny.MultiSet([d_1524_v13_])
            index274_ = default__.safeIndex(67, (d_1523_v12_).length(0))
            (d_1523_v12_)[index274_] = (d_1525_v14_) == (d_1525_v14_)
            d_1526_v15_: C1
            nw279_ = C1()
            nw279_.ctor__((d_1518_v8_ if p2 else d_1518_v8_))
            d_1526_v15_ = nw279_
        elif True:
            d_1527_v16_: _dafny.Array
            nw280_ = _dafny.Array(_dafny.Array(None, 0), 8)
            d_1527_v16_ = nw280_
            d_1527_v16_ = (d_1527_v16_ if not (True) or (p2) else d_1527_v16_)
            d_1528_v17_: _dafny.MultiSet
            d_1528_v17_ = _dafny.MultiSet([p1])
            d_1529_v18_: D5
            d_1529_v18_ = D5_DC8(d_1528_v17_)
            d_1530_v19_: bool
            d_1530_v19_ = True
            d_1531_v20_: _dafny.Seq
            d_1531_v20_ = _dafny.SeqWithoutIsStrInference([d_1530_v19_])
            d_1532_v21_: C3
            nw281_ = C3()
            nw281_.ctor__(self.f2)
            d_1532_v21_ = nw281_
            d_1533_v22_: _dafny.Map
            d_1533_v22_ = _dafny.Map({d_1531_v20_: d_1532_v21_})
            d_1534_v23_: _dafny.Map
            d_1534_v23_ = _dafny.Map({self.f2: 479})
            d_1535_v24_: _dafny.MultiSet
            d_1535_v24_ = _dafny.MultiSet([self.f2, default__.safeModuloInt(self.f2, ((d_1534_v23_)[61] if (61) in (d_1534_v23_) else self.f2))])
            d_1536_v25_: D0
            d_1536_v25_ = D0_DC0(p2)
            d_1537_v26_: _dafny.Seq
            d_1537_v26_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hweyh"))])
            d_1538_v27_: _dafny.Seq
            d_1538_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljkq"))
            d_1539_v28_: _dafny.Seq
            d_1539_v28_ = _dafny.SeqWithoutIsStrInference([self.f2])
            rhs246_ = d_1529_v18_
            rhs247_ = (d_1536_v25_).cf0
            rhs248_ = d_1533_v22_
            rhs249_ = len((d_1537_v26_) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1540_i2_ in range(default__.abs(-522))]), d_1538_v27_])))
            rhs250_ = (_dafny.MultiSet(d_1539_v28_)) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f2 for d_1541_i3_ in range(default__.abs(523))])))
            d_1529_v18_ = rhs246_
            d_1530_v19_ = rhs247_
            d_1533_v22_ = rhs248_
            r1 = rhs249_
            d_1535_v24_ = rhs250_
            d_1542_v29_: C3
            nw282_ = C3()
            nw282_.ctor__(default__.safeDivisionInt(-329, (0) - (self.f2)))
            d_1542_v29_ = nw282_
            d_1543_v30_: _dafny.Array
            nw283_ = _dafny.Array(None, 11)
            nw283_[int(0)] = d_1538_v27_
            nw283_[int(1)] = d_1538_v27_
            nw283_[int(2)] = d_1538_v27_
            nw283_[int(3)] = d_1538_v27_
            nw283_[int(4)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1544_i4_ in range(default__.abs(-663))])) + (d_1538_v27_)
            nw283_[int(5)] = d_1538_v27_
            nw283_[int(6)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjfrs"))) + (d_1538_v27_)
            nw283_[int(7)] = (d_1538_v27_) + (d_1538_v27_)
            nw283_[int(8)] = d_1538_v27_
            nw283_[int(9)] = d_1538_v27_
            nw283_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "boyqhdgay"))
            d_1543_v30_ = nw283_
            index275_ = default__.safeIndex(802, (d_1543_v30_).length(0))
            (d_1543_v30_)[index275_] = d_1538_v27_
            index276_ = default__.safeIndex(802, (d_1543_v30_).length(0))
            (d_1543_v30_)[index276_] = d_1538_v27_
            d_1531_v20_ = (d_1531_v20_) + ((d_1531_v20_) + (d_1531_v20_))
        hi8_ = self.f2
        for d_1545_i5_ in range((self.f2 if p3 else self.f2), hi8_):
            d_1546_v31_: _dafny.Seq
            d_1546_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trytchnh"))
            d_1547_v32_: C1
            nw284_ = C1()
            nw284_.ctor__(d_1546_v31_)
            d_1547_v32_ = nw284_
            d_1548_v33_: D8
            d_1548_v33_ = D8_DC16(d_1547_v32_)
            d_1548_v33_ = d_1548_v33_
            r1 = d_1545_i5_
            d_1549_v34_: _dafny.Array
            def lambda59_(d_1550_p2_):
                def lambda60_(d_1551_i6_):
                    return d_1550_p2_

                return lambda60_

            init31_ = lambda59_(p2)
            nw285_ = _dafny.Array(None, 21)
            for i0_31_ in range(nw285_.length(0)):
                nw285_[i0_31_] = init31_(i0_31_)
            d_1549_v34_ = nw285_
            d_1552_v35_: _dafny.Map
            d_1552_v35_ = _dafny.Map({p3: d_1549_v34_})
            d_1552_v35_ = d_1552_v35_
            index277_ = default__.safeIndex(931, (d_1549_v34_).length(0))
            (d_1549_v34_)[index277_] = (self).fm8(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1553_i7_ in range(default__.abs(507))]), not(p2), globalState)
            index278_ = default__.safeIndex(931, (d_1549_v34_).length(0))
            (d_1549_v34_)[index278_] = p3
        d_1554_v36_: _dafny.Array
        nw286_ = _dafny.Array(_dafny.Seq({}), 17)
        d_1554_v36_ = nw286_
        d_1554_v36_ = d_1554_v36_
        d_1555_v37_: C0
        nw287_ = C0()
        nw287_.ctor__(p3, not(p2))
        d_1555_v37_ = nw287_
        r0 = default__.safeModuloInt(837, self.f2)
        r1 = (0) - (default__.safeDivisionInt(self.f2, (0) - (self.f2)))
        return r0, r1


class C5(T0):
    def  __init__(self):
        self.f1: _dafny.Array = _dafny.Array(None, 0)
        self._f0: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f0, f1):
        (self)._f0 = f0
        (self).f1 = f1

    def fm0(self, globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f0])) for d_1556_i0_ in range(default__.abs(158))])) + (_dafny.SeqWithoutIsStrInference([(self).f0 for d_1557_i1_ in range(default__.abs(432))])): (0) - ((self).f0)})

    def fm1(self, p0, globalState):
        return 426

    def fm6(self, p0, globalState):
        return True

    def fm7(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f0: _dafny.CodePoint('m')}), _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality: _dafny.CodePoint('m')})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f0: _dafny.CodePoint('b')}), _dafny.Map({(self).f0: _dafny.CodePoint('e')}), _dafny.Map({(self).f0: _dafny.CodePoint('c')}), _dafny.Map({(self).f0: _dafny.CodePoint('l')}), _dafny.Map({(self).f0: _dafny.CodePoint('j')})]))) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f0: _dafny.CodePoint('i')})]))

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        r0 = p2
        d_1558_v0_: _dafny.Seq
        d_1558_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtifm"))
        hi9_ = (self).f0
        for d_1559_i0_ in range(len(d_1558_v0_), hi9_):
            if p2:
                d_1560_v1_: _dafny.Set
                d_1560_v1_ = _dafny.Set({len(_dafny.Map({d_1559_i0_: p3}))})
                d_1561_v2_: _dafny.Map
                d_1561_v2_ = _dafny.Map({p3: p3})
                d_1562_v3_: _dafny.Map
                d_1562_v3_ = _dafny.Map({(_dafny.Set({len(d_1561_v2_)})).issubset(d_1560_v1_): (self).f0})
                rhs251_ = -735
                rhs252_ = p0
                rhs253_ = (0) - ((self).f0)
                rhs254_ = d_1562_v3_
                rhs255_ = ((self).f0) != (p0)
                r2 = rhs251_
                r2 = rhs252_
                r2 = rhs253_
                d_1562_v3_ = rhs254_
                r0 = rhs255_
                arr2_ = self.f1
                index279_ = default__.safeIndex(337, (self.f1).length(0))
                arr2_[index279_] = p2
                arr3_ = self.f1
                index280_ = default__.safeIndex(337, (self.f1).length(0))
                arr3_[index280_] = (len(d_1561_v2_)) > (d_1559_i0_)
                d_1563_v4_: _dafny.Map
                d_1563_v4_ = _dafny.Map({p3: d_1561_v2_})
                d_1564_v5_: _dafny.Map
                d_1564_v5_ = d_1561_v2_
                d_1563_v4_ = (d_1563_v4_).set(p2, (d_1564_v5_))
                arr4_ = self.f1
                index281_ = default__.safeIndex(337, (self.f1).length(0))
                arr4_[index281_] = not(True)
                d_1565_v6_: C2
                nw288_ = C2()
                nw288_.ctor__(d_1559_i0_)
                d_1565_v6_ = nw288_
            elif True:
                d_1566_v7_: _dafny.Map
                d_1566_v7_ = _dafny.Map({p2: self.f1})
                d_1566_v7_ = (d_1566_v7_).set(p2, self.f1)
                arr5_ = self.f1
                index282_ = default__.safeIndex(711, (self.f1).length(0))
                arr5_[index282_] = not (p3) or (p2)
                d_1567_v8_: _dafny.MultiSet
                d_1567_v8_ = _dafny.MultiSet([self.f1, self.f1])
                arr6_ = self.f1
                index283_ = default__.safeIndex(711, (self.f1).length(0))
                rhs256_ = ((d_1567_v8_) | (d_1567_v8_)).isdisjoint((d_1567_v8_).intersection(d_1567_v8_))
                rhs257_ = d_1558_v0_
                lhs134_ = self.f1
                lhs135_ = default__.safeIndex(711, (self.f1).length(0))
                lhs134_[lhs135_] = rhs256_
                d_1558_v0_ = rhs257_
                r0 = (self.f1)[default__.safeIndex(711, (self.f1).length(0))]
                d_1568_v9_: C0
                nw289_ = C0()
                nw289_.ctor__((self.f1)[default__.safeIndex(711, (self.f1).length(0))], p3)
                d_1568_v9_ = nw289_
                d_1569_v10_: _dafny.Array
                def lambda61_(d_1570_i0_):
                    def lambda62_(d_1571_i1_):
                        return (d_1571_i1_) * (d_1570_i0_)

                    return lambda62_

                init32_ = lambda61_(d_1559_i0_)
                nw290_ = _dafny.Array(None, 24)
                for i0_32_ in range(nw290_.length(0)):
                    nw290_[i0_32_] = init32_(i0_32_)
                d_1569_v10_ = nw290_
                index284_ = default__.safeIndex(291, (d_1569_v10_).length(0))
                (d_1569_v10_)[index284_] = default__.safeDivisionInt(-535, d_1559_i0_)
                d_1572_v11_: _dafny.Set
                d_1572_v11_ = _dafny.Set({d_1568_v9_.f5})
                d_1573_v12_: _dafny.MultiSet
                d_1573_v12_ = _dafny.MultiSet([d_1572_v11_])
                index285_ = default__.safeIndex(291, (d_1569_v10_).length(0))
                (d_1569_v10_)[index285_] = ((_dafny.MultiSet([_dafny.Set({d_1568_v9_.f5, (self.f1)[default__.safeIndex(711, (self.f1).length(0))]})])) - (d_1573_v12_)).cardinality
            d_1574_v13_: str
            d_1574_v13_ = _dafny.CodePoint('j')
            d_1575_v14_: _dafny.Array
            nw291_ = _dafny.Array(int(0), 9)
            d_1575_v14_ = nw291_
            d_1576_v15_: _dafny.Map
            d_1576_v15_ = _dafny.Map({p2: p3})
            d_1577_v16_: D6
            d_1577_v16_ = D6_DC11(_dafny.SeqWithoutIsStrInference([(d_1558_v0_)[default__.safeIndex((self).f0, len(d_1558_v0_))] for d_1578_i2_ in range(default__.abs(-155))]))
            pat_let_tv33_ = d_1558_v0_
            def iife126_(_pat_let33_0):
                def iife127_(d_1579_dt__update__tmp_h0_):
                    def iife128_(_pat_let34_0):
                        def iife129_(d_1580_dt__update_hcf21_h0_):
                            return D6_DC11(d_1580_dt__update_hcf21_h0_)
                        return iife129_(_pat_let34_0)
                    return iife128_(pat_let_tv33_)
                return iife127_(_pat_let33_0)
            rhs258_ = _dafny.CodePoint('u')
            rhs259_ = d_1575_v14_
            rhs260_ = not (p3) or (not(p2))
            rhs261_ = d_1576_v15_
            rhs262_ = (D6_DC11(d_1558_v0_) if p2 else iife126_(default__.fm42(d_1559_i0_, p3, p2, globalState)))
            d_1574_v13_ = rhs258_
            d_1575_v14_ = rhs259_
            r0 = rhs260_
            d_1576_v15_ = rhs261_
            d_1577_v16_ = rhs262_
            d_1581_v17_: C1
            nw292_ = C1()
            nw292_.ctor__(d_1558_v0_)
            d_1581_v17_ = nw292_
            d_1582_v18_: D8
            d_1582_v18_ = D8_DC16(d_1581_v17_)
            rhs263_ = d_1582_v18_
            rhs264_ = (p0) + ((0) - (d_1559_i0_))
            rhs265_ = (False if p2 else p2)
            d_1582_v18_ = rhs263_
            r2 = rhs264_
            r0 = rhs265_
            d_1583_v19_: C3
            nw293_ = C3()
            nw293_.ctor__(d_1559_i0_)
            d_1583_v19_ = nw293_
        d_1584_v20_: _dafny.Seq
        d_1584_v20_ = _dafny.SeqWithoutIsStrInference([p3])
        r2 = default__.safeModuloInt(p0, len((d_1584_v20_) + (d_1584_v20_)))
        d_1585_i3_: int
        d_1585_i3_ = 0
        with _dafny.label("7"):
            while p3:
                with _dafny.c_label("7"):
                    if (d_1585_i3_) >= (100):
                        raise _dafny.Break("7")
                    d_1585_i3_ = (d_1585_i3_) + (1)
                    d_1558_v0_ = d_1558_v0_
                    r0 = not(not (p2) or (p3))
                    d_1586_v21_: _dafny.MultiSet
                    d_1586_v21_ = _dafny.MultiSet([660, p0])
                    if (default__.safeDivisionInt(p0, (self).f0)) <= ((d_1586_v21_).cardinality):
                        d_1587_v22_: _dafny.Map
                        d_1587_v22_ = _dafny.Map({p3: d_1558_v0_})
                        d_1588_v23_: _dafny.Map
                        d_1588_v23_ = _dafny.Map({(((d_1587_v22_)[False] if (False) in (d_1587_v22_) else d_1558_v0_)).set(default__.safeIndex((self).f0, len(((d_1587_v22_)[False] if (False) in (d_1587_v22_) else d_1558_v0_))), _dafny.CodePoint('p')): d_1558_v0_})
                        d_1589_v24_: _dafny.Map
                        d_1589_v24_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "darjgmy"))): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1590_i4_ in range(default__.abs(-130))])})
                        d_1588_v23_ = (d_1588_v23_).set(((d_1589_v24_)[p0] if (p0) in (d_1589_v24_) else d_1558_v0_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "min")))
                        d_1558_v0_ = d_1558_v0_
                        r2 = default__.safeModuloInt(205, (self).f0)
                        d_1591_v25_: _dafny.Map
                        d_1591_v25_ = _dafny.Map({(_dafny.Map({p0: (self).f0})).set((self).f0, (self).f0): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ja"))})
                        d_1592_v26_: _dafny.Map
                        d_1592_v26_ = _dafny.Map({p0: (self).f0})
                        d_1591_v25_ = (d_1591_v25_).set(d_1592_v26_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxeuxp")))
                        rhs266_ = default__.safeModuloInt(p0, p0)
                        rhs267_ = p3
                        rhs268_ = p3
                        rhs269_ = (self).fm6(True, globalState)
                        r2 = rhs266_
                        r0 = rhs267_
                        r0 = rhs268_
                        r0 = rhs269_
                    elif True:
                        r2 = (self).f0
                        d_1593_v27_: C3
                        nw294_ = C3()
                        nw294_.ctor__((self).f0)
                        d_1593_v27_ = nw294_
                        d_1594_v28_: _dafny.Array
                        nw295_ = _dafny.Array(None, 8)
                        nw295_[int(0)] = default__.safeModuloInt((self).f0, p0)
                        nw295_[int(1)] = ((self).f0) - ((self).f0)
                        nw295_[int(2)] = p0
                        nw295_[int(3)] = (0) - ((self).f0)
                        nw295_[int(4)] = (self).f0
                        nw295_[int(5)] = (0) - (p0)
                        nw295_[int(6)] = (self).f0
                        nw295_[int(7)] = default__.safeDivisionInt((self).f0, ((d_1586_v21_)[p0] if (p0) in (d_1586_v21_) else p0))
                        d_1594_v28_ = nw295_
                        index286_ = default__.safeIndex(387, (d_1594_v28_).length(0))
                        (d_1594_v28_)[index286_] = 842
                        index287_ = default__.safeIndex(387, (d_1594_v28_).length(0))
                        (d_1594_v28_)[index287_] = p0
                        d_1595_v29_: _dafny.Array
                        def lambda63_(d_1596_i5_):
                            return _dafny.Map({33: 753})

                        init33_ = lambda63_
                        nw296_ = _dafny.Array(None, 2)
                        for i0_33_ in range(nw296_.length(0)):
                            nw296_[i0_33_] = init33_(i0_33_)
                        d_1595_v29_ = nw296_
                        d_1597_v30_: _dafny.Map
                        d_1597_v30_ = _dafny.Map({(d_1584_v20_)[default__.safeIndex((self).f0, len(d_1584_v20_))]: p3})
                        d_1598_v31_: _dafny.Map
                        d_1598_v31_ = _dafny.Map({(d_1595_v29_): d_1597_v30_})
                        d_1598_v31_ = (d_1598_v31_).set(d_1595_v29_, d_1597_v30_)
                        d_1599_v32_: _dafny.Map
                        d_1599_v32_ = _dafny.Map({p0: p0})
                        d_1600_v33_: _dafny.Seq
                        d_1600_v33_ = _dafny.SeqWithoutIsStrInference([(self).f0])
                        r0 = (len(_dafny.Map({((d_1586_v21_)[len(d_1599_v32_)] if (len(d_1599_v32_)) in (d_1586_v21_) else (0) - (p0)): not(p3)}))) <= (len(d_1600_v33_))
                    arr7_ = self.f1
                    index288_ = default__.safeIndex(656, (self.f1).length(0))
                    arr7_[index288_] = (p0) == (p0)
                    arr8_ = self.f1
                    index289_ = default__.safeIndex(656, (self.f1).length(0))
                    arr8_[index289_] = p2
                    pass
            pass
        d_1601_v34_: C2
        nw297_ = C2()
        nw297_.ctor__(default__.safeModuloInt((self).f0, len(d_1558_v0_)))
        d_1601_v34_ = nw297_
        d_1602_v35_: _dafny.Map
        d_1602_v35_ = _dafny.Map({p0: len(d_1558_v0_)})
        d_1603_v36_: _dafny.Map
        d_1603_v36_ = _dafny.Map({(0) - ((self).f0): p2})
        d_1604_v37_: _dafny.Seq
        d_1604_v37_ = _dafny.SeqWithoutIsStrInference([(d_1601_v34_).fm1(True, globalState), len(d_1603_v36_), p0])
        d_1605_v38_: _dafny.MultiSet
        d_1605_v38_ = _dafny.MultiSet([p2])
        d_1606_v39_: _dafny.Set
        d_1606_v39_ = _dafny.Set({p3, p2})
        d_1607_v40_: _dafny.Array
        nw298_ = _dafny.Array(None, 22)
        nw298_[int(0)] = (self).fm1(p3, globalState)
        nw298_[int(1)] = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1608_i6_ in range(default__.abs(-749))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhwnmp"))))
        nw298_[int(2)] = p0
        nw298_[int(3)] = (self).f0
        nw298_[int(4)] = (self).f0
        nw298_[int(5)] = (self).f0
        nw298_[int(6)] = (self).f0
        nw298_[int(7)] = (len(d_1602_v35_)) + (len(d_1604_v37_))
        nw298_[int(8)] = ((d_1605_v38_)[p3] if (p3) in (d_1605_v38_) else len(d_1606_v39_))
        nw298_[int(9)] = ((d_1601_v34_).fm1(p3, globalState) if p2 else p0)
        nw298_[int(10)] = (p0) - ((0) - (-8))
        nw298_[int(11)] = p0
        nw298_[int(12)] = ((0) - ((self).f0)) * ((0) - (p0))
        nw298_[int(13)] = p0
        nw298_[int(14)] = default__.safeModuloInt(len(d_1558_v0_), (self).f0)
        nw298_[int(15)] = len(d_1558_v0_)
        nw298_[int(16)] = len(d_1558_v0_)
        nw298_[int(17)] = (self).f0
        nw298_[int(18)] = (p0) - ((self).f0)
        nw298_[int(19)] = (self).f0
        nw298_[int(20)] = len(d_1558_v0_)
        nw298_[int(21)] = len(d_1604_v37_)
        d_1607_v40_ = nw298_
        d_1607_v40_ = d_1607_v40_
        d_1609_v41_: _dafny.Set
        d_1609_v41_ = _dafny.Set({p0})
        r0 = (d_1609_v41_) != ((d_1609_v41_).intersection(d_1609_v41_))
        d_1610_v42_: _dafny.Map
        d_1610_v42_ = _dafny.Map({p3: p3})
        d_1611_v43_: _dafny.Map
        d_1611_v43_ = _dafny.Map({len(d_1558_v0_): d_1610_v42_})
        nw299_ = _dafny.Array(None, 24)
        nw299_[int(0)] = d_1610_v42_
        nw299_[int(1)] = _dafny.Map({False: p3})
        nw299_[int(2)] = d_1610_v42_
        nw299_[int(3)] = d_1610_v42_
        nw299_[int(4)] = d_1610_v42_
        nw299_[int(5)] = (d_1610_v42_ if p2 else _dafny.Map({p3: True}))
        nw299_[int(6)] = d_1610_v42_
        nw299_[int(7)] = (_dafny.Map({p3: p2}) if p2 else (d_1610_v42_).set(p2, (self).fm6(False, globalState)))
        nw299_[int(8)] = _dafny.Map({not(p3): p2})
        nw299_[int(9)] = (d_1610_v42_) | (d_1610_v42_)
        nw299_[int(10)] = (d_1610_v42_).set(p3, not(p2))
        nw299_[int(11)] = d_1610_v42_
        nw299_[int(12)] = default__.fm21(globalState)
        nw299_[int(13)] = d_1610_v42_
        nw299_[int(14)] = (((d_1611_v43_)[p0] if (p0) in (d_1611_v43_) else d_1610_v42_)) | (default__.fm21(globalState))
        nw299_[int(15)] = (d_1610_v42_) | (_dafny.Map({p3: p3}))
        nw299_[int(16)] = d_1610_v42_
        nw299_[int(17)] = d_1610_v42_
        nw299_[int(18)] = d_1610_v42_
        nw299_[int(19)] = d_1610_v42_
        nw299_[int(20)] = (d_1610_v42_) | (d_1610_v42_)
        nw299_[int(21)] = d_1610_v42_
        nw299_[int(22)] = _dafny.Map({p3: p2})
        nw299_[int(23)] = d_1610_v42_
        r1 = nw299_
        d_1612_v44_: _dafny.Map
        d_1612_v44_ = _dafny.Map({(D6_DC12(p2, p3)).cf22: (0) - ((self).f0)})
        r2 = ((d_1612_v44_)[not(p3)] if (not(p3)) in (d_1612_v44_) else ((self).f0) + ((self).f0))
        return r0, r1, r2

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_1613_v0_: str
        d_1613_v0_ = _dafny.CodePoint('f')
        d_1614_v1_: _dafny.Array
        nw300_ = _dafny.Array(None, 2)
        nw300_[int(0)] = d_1613_v0_
        nw300_[int(1)] = (d_1613_v0_ if True else d_1613_v0_)
        d_1614_v1_ = nw300_
        index290_ = default__.safeIndex(990, (d_1614_v1_).length(0))
        (d_1614_v1_)[index290_] = d_1613_v0_
        d_1615_v2_: _dafny.Array
        def lambda64_(d_1616_i0_):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfbmfbdi"))

        init34_ = lambda64_
        nw301_ = _dafny.Array(None, 2)
        for i0_34_ in range(nw301_.length(0)):
            nw301_[i0_34_] = init34_(i0_34_)
        d_1615_v2_ = nw301_
        d_1617_v3_: _dafny.Seq
        d_1617_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bu"))
        index291_ = default__.safeIndex(489, (d_1615_v2_).length(0))
        (d_1615_v2_)[index291_] = d_1617_v3_
        d_1618_v4_: _dafny.Map
        d_1618_v4_ = _dafny.Map({p2: (self).f0})
        index292_ = default__.safeIndex(990, (d_1614_v1_).length(0))
        index293_ = default__.safeIndex(489, (d_1615_v2_).length(0))
        rhs270_ = _dafny.CodePoint('a')
        rhs271_ = (d_1617_v3_) + (((d_1617_v3_).set(default__.safeIndex(len(d_1617_v3_), len(d_1617_v3_)), d_1613_v0_)) + ((_dafny.SeqWithoutIsStrInference([d_1613_v0_ for d_1619_i1_ in range(default__.abs(632))])).set(default__.safeIndex(((d_1618_v4_)[p2] if (p2) in (d_1618_v4_) else p1), len(_dafny.SeqWithoutIsStrInference([d_1613_v0_ for d_1620_i1_ in range(default__.abs(632))]))), d_1613_v0_)))
        lhs136_ = d_1614_v1_
        lhs137_ = default__.safeIndex(990, (d_1614_v1_).length(0))
        lhs138_ = d_1615_v2_
        lhs139_ = default__.safeIndex(489, (d_1615_v2_).length(0))
        lhs136_[lhs137_] = rhs270_
        lhs138_[lhs139_] = rhs271_
        d_1621_v5_: D6
        d_1621_v5_ = D6_DC11((d_1615_v2_)[default__.safeIndex(489, (d_1615_v2_).length(0))])
        d_1622_v6_: _dafny.Seq
        d_1622_v6_ = _dafny.SeqWithoutIsStrInference([(d_1621_v5_).cf21, (d_1615_v2_)[default__.safeIndex(489, (d_1615_v2_).length(0))]])
        rhs272_ = d_1617_v3_
        rhs273_ = len(_dafny.Set({(d_1615_v2_)[default__.safeIndex(489, (d_1615_v2_).length(0))], (d_1617_v3_).set(default__.safeIndex(p0, len(d_1617_v3_)), (d_1614_v1_)[default__.safeIndex(990, (d_1614_v1_).length(0))]), (d_1622_v6_)[default__.safeIndex(p1, len(d_1622_v6_))], (d_1615_v2_)[default__.safeIndex(489, (d_1615_v2_).length(0))]}))
        d_1617_v3_ = rhs272_
        r0 = rhs273_
        d_1623_v7_: _dafny.Seq
        d_1623_v7_ = _dafny.SeqWithoutIsStrInference([self.f1])
        d_1624_v8_: D11
        d_1624_v8_ = D11_DC20((d_1623_v7_)[default__.safeIndex(p1, len(d_1623_v7_))])
        d_1625_v9_: _dafny.Seq
        d_1625_v9_ = _dafny.SeqWithoutIsStrInference([d_1624_v8_, d_1624_v8_])
        d_1626_v10_: _dafny.Seq
        d_1626_v10_ = _dafny.SeqWithoutIsStrInference([d_1625_v9_])
        if (d_1624_v8_) in ((((d_1626_v10_)[default__.safeIndex((self).f0, len(d_1626_v10_))]) + (d_1625_v9_)).set(default__.safeIndex((self).f0, len(((d_1626_v10_)[default__.safeIndex((self).f0, len(d_1626_v10_))]) + (d_1625_v9_))), d_1624_v8_)):
            r0 = (p0) * (-22)
            d_1627_v11_: _dafny.Array
            def lambda65_(d_1628_p0_):
                def lambda66_(d_1629_i2_):
                    return (d_1629_i2_) * (d_1628_p0_)

                return lambda66_

            init35_ = lambda65_(p0)
            nw302_ = _dafny.Array(None, 2)
            for i0_35_ in range(nw302_.length(0)):
                nw302_[i0_35_] = init35_(i0_35_)
            d_1627_v11_ = nw302_
            index294_ = default__.safeIndex(263, (d_1627_v11_).length(0))
            (d_1627_v11_)[index294_] = (859) + (p0)
            d_1630_v12_: _dafny.Set
            d_1630_v12_ = _dafny.Set({(0) - (p1), (self).f0})
            index295_ = default__.safeIndex(263, (d_1627_v11_).length(0))
            (d_1627_v11_)[index295_] = default__.safeDivisionInt((self).f0, len((d_1630_v12_ if p2 else d_1630_v12_)))
            if p2:
                d_1631_v13_: _dafny.Seq
                d_1631_v13_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1632_v14_: _dafny.Map
                d_1632_v14_ = _dafny.Map({d_1613_v0_: _dafny.MultiSet([(d_1627_v11_)[default__.safeIndex(263, (d_1627_v11_).length(0))], (d_1631_v13_)[default__.safeIndex(((d_1618_v4_)[p2] if (p2) in (d_1618_v4_) else (self).f0), len(d_1631_v13_))]])})
                d_1633_v15_: _dafny.Map
                d_1633_v15_ = _dafny.Map({d_1613_v0_: p0})
                index296_ = default__.safeIndex(263, (d_1627_v11_).length(0))
                (d_1627_v11_)[index296_] = (len((d_1632_v14_) | (default__.fm43(globalState)))) - (((d_1633_v15_)[d_1613_v0_] if (d_1613_v0_) in (d_1633_v15_) else len(d_1631_v13_)))
                index297_ = default__.safeIndex(990, (d_1614_v1_).length(0))
                (d_1614_v1_)[index297_] = default__.fm37(p1, globalState)
                index298_ = default__.safeIndex(263, (d_1627_v11_).length(0))
                (d_1627_v11_)[index298_] = (self).f0
                d_1634_v16_: _dafny.Array
                out27_: _dafny.Array
                out27_ = default__.m0(p2, p2, globalState)
                d_1634_v16_ = out27_
                r0 = (d_1627_v11_)[default__.safeIndex(263, (d_1627_v11_).length(0))]
            elif True:
                index299_ = default__.safeIndex(263, (d_1627_v11_).length(0))
                (d_1627_v11_)[index299_] = (self).f0
                d_1635_v17_: bool
                d_1635_v17_ = False
                d_1635_v17_ = d_1635_v17_
                index300_ = default__.safeIndex(263, (d_1627_v11_).length(0))
                (d_1627_v11_)[index300_] = (self).f0
                d_1636_v18_: C3
                nw303_ = C3()
                nw303_.ctor__(p0)
                d_1636_v18_ = nw303_
                d_1637_v19_: _dafny.Map
                d_1637_v19_ = _dafny.Map({p2: p2})
                d_1638_v20_: _dafny.Map
                d_1638_v20_ = d_1637_v19_
                d_1639_v21_: _dafny.Map
                d_1639_v21_ = _dafny.Map({d_1638_v20_: self.f1})
                d_1639_v21_ = (d_1639_v21_).set(d_1638_v20_, self.f1)
            d_1640_v22_: _dafny.Set
            d_1640_v22_ = _dafny.Set({not(p2), p2, p2})
            d_1641_v23_: _dafny.Map
            d_1641_v23_ = _dafny.Map({d_1640_v22_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgpvdxtv"))})
            d_1641_v23_ = (d_1641_v23_).set(d_1640_v22_, d_1617_v3_)
            d_1642_v24_: _dafny.Array
            nw304_ = _dafny.Array(_dafny.Array(None, 0), 16)
            d_1642_v24_ = nw304_
            pat_let_tv34_ = d_1617_v3_
            d_1643_v25_: _dafny.Map
            def iife130_(_pat_let35_0):
                def iife131_(d_1644_dt__update__tmp_h0_):
                    def iife132_(_pat_let36_0):
                        def iife133_(d_1645_dt__update_hcf21_h0_):
                            return D6_DC11(d_1645_dt__update_hcf21_h0_)
                        return iife133_(_pat_let36_0)
                    return iife132_(pat_let_tv34_)
                return iife131_(_pat_let35_0)
            d_1643_v25_ = _dafny.Map({(iife130_(d_1621_v5_)).cf21: d_1642_v24_})
            d_1643_v25_ = (d_1643_v25_).set((d_1617_v3_).set(default__.safeIndex(-940, len(d_1617_v3_)), d_1613_v0_), d_1642_v24_)
        elif True:
            d_1646_v26_: C0
            nw305_ = C0()
            nw305_.ctor__((self).fm6(p2, globalState), p2)
            d_1646_v26_ = nw305_
            r0 = 183
            d_1647_v27_: _dafny.Seq
            d_1647_v27_ = _dafny.SeqWithoutIsStrInference([len(d_1617_v3_), (self).f0, (self).f0])
            d_1648_v28_: C3
            nw306_ = C3()
            nw306_.ctor__((d_1647_v27_)[default__.safeIndex((self).f0, len(d_1647_v27_))])
            d_1648_v28_ = nw306_
            arr9_ = self.f1
            index301_ = default__.safeIndex(262, (self.f1).length(0))
            arr9_[index301_] = d_1646_v26_.f5
            arr10_ = self.f1
            index302_ = default__.safeIndex(262, (self.f1).length(0))
            arr10_[index302_] = p2
            if d_1646_v26_.f4:
                d_1649_v29_: C3
                nw307_ = C3()
                nw307_.ctor__(len((d_1617_v3_) + (d_1617_v3_)))
                d_1649_v29_ = nw307_
                d_1649_v29_ = (d_1649_v29_ if d_1646_v26_.f5 else d_1649_v29_)
                r0 = (self).f0
                d_1650_v30_: _dafny.MultiSet
                d_1650_v30_ = _dafny.MultiSet([False])
                r0 = ((p1) * ((d_1650_v30_).cardinality)) - ((self).f0)
                d_1651_v31_: _dafny.Array
                nw308_ = _dafny.Array(None, 14)
                d_1651_v31_ = nw308_
                d_1652_v32_: T2
                nw309_ = C4()
                nw309_.ctor__(len((d_1615_v2_)[default__.safeIndex(489, (d_1615_v2_).length(0))]))
                d_1652_v32_ = nw309_
                index303_ = default__.safeIndex(928, (d_1651_v31_).length(0))
                (d_1651_v31_)[index303_] = d_1652_v32_
                d_1653_v33_: _dafny.Array
                nw310_ = _dafny.Array(int(0), 27)
                d_1653_v33_ = nw310_
                d_1654_v34_: _dafny.Seq
                d_1654_v34_ = _dafny.SeqWithoutIsStrInference([d_1653_v33_, d_1653_v33_])
                d_1655_v35_: _dafny.Array
                nw311_ = _dafny.Array(None, 16)
                nw311_[int(0)] = d_1653_v33_
                nw311_[int(1)] = d_1653_v33_
                nw311_[int(2)] = d_1653_v33_
                nw311_[int(3)] = d_1653_v33_
                nw311_[int(4)] = d_1653_v33_
                nw311_[int(5)] = d_1653_v33_
                nw311_[int(6)] = d_1653_v33_
                nw311_[int(7)] = d_1653_v33_
                nw311_[int(8)] = d_1653_v33_
                nw311_[int(9)] = d_1653_v33_
                nw311_[int(10)] = d_1653_v33_
                nw311_[int(11)] = d_1653_v33_
                nw311_[int(12)] = d_1653_v33_
                nw311_[int(13)] = (d_1654_v34_)[default__.safeIndex(853, len(d_1654_v34_))]
                nw311_[int(14)] = d_1653_v33_
                nw311_[int(15)] = d_1653_v33_
                d_1655_v35_ = nw311_
                d_1656_v36_: _dafny.Map
                d_1656_v36_ = _dafny.Map({d_1655_v35_: d_1647_v27_})
                d_1657_v37_: _dafny.Map
                d_1657_v37_ = _dafny.Map({p0: (self).f0})
                d_1658_v38_: _dafny.Set
                d_1658_v38_ = _dafny.Set({_dafny.Map({d_1621_v5_: len(d_1657_v37_)})})
                d_1659_v39_: _dafny.Map
                d_1659_v39_ = _dafny.Map({D6_DC11(d_1617_v3_): p1})
                d_1660_v40_: D11
                d_1660_v40_ = D11_DC21(d_1613_v0_, len(_dafny.SeqWithoutIsStrInference([(self).f0 for d_1661_i3_ in range(default__.abs(11))])), p0)
                index304_ = default__.safeIndex(928, (d_1651_v31_).length(0))
                rhs274_ = (default__.fm36(d_1646_v26_.f4, _dafny.CodePoint('i'), (self).fm1(d_1646_v26_.f5, globalState), globalState)) + (((d_1656_v36_)[d_1655_v35_] if (d_1655_v35_) in (d_1656_v36_) else d_1647_v27_))
                rhs275_ = default__.safeModuloInt(p1, (self).f0)
                rhs276_ = ((d_1658_v38_ if d_1646_v26_.f5 else (_dafny.Set({d_1659_v39_, _dafny.Map({d_1621_v5_: d_1652_v32_.f2})})))).issubset(d_1658_v38_)
                rhs277_ = d_1652_v32_
                rhs278_ = (p1) - ((d_1660_v40_).cf38)
                lhs140_ = d_1646_v26_
                lhs141_ = d_1651_v31_
                lhs142_ = default__.safeIndex(928, (d_1651_v31_).length(0))
                lhs143_ = d_1652_v32_
                d_1647_v27_ = rhs274_
                r0 = rhs275_
                lhs140_.f5 = rhs276_
                lhs141_[lhs142_] = rhs277_
                lhs143_.f2 = rhs278_
                d_1662_v41_: _dafny.Set
                d_1662_v41_ = _dafny.Set({d_1646_v26_.f5})
                (d_1652_v32_).f2 = len(d_1662_v41_)
            elif True:
                d_1663_v42_: _dafny.Array
                nw312_ = _dafny.Array(int(0), 9)
                d_1663_v42_ = nw312_
                d_1664_v43_: _dafny.Seq
                d_1664_v43_ = _dafny.SeqWithoutIsStrInference([d_1663_v42_, d_1663_v42_])
                r0 = default__.safeDivisionInt(p1, (p0 if d_1646_v26_.f5 else len(d_1664_v43_)))
                d_1665_v44_: _dafny.Set
                d_1665_v44_ = _dafny.Set({d_1646_v26_.f4})
                d_1666_v45_: D4
                d_1666_v45_ = D4_DC5(d_1665_v44_)
                index305_ = default__.safeIndex(990, (d_1614_v1_).length(0))
                index306_ = default__.safeIndex(489, (d_1615_v2_).length(0))
                rhs279_ = (self).fm1(d_1646_v26_.f5, globalState)
                rhs280_ = (d_1617_v3_)[default__.safeIndex((p0) * (144), len(d_1617_v3_))]
                rhs281_ = len((_dafny.Set({d_1646_v26_.f4})) | ((d_1666_v45_).cf8))
                rhs282_ = (d_1622_v6_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f0])), len(d_1622_v6_))]
                lhs144_ = d_1614_v1_
                lhs145_ = default__.safeIndex(990, (d_1614_v1_).length(0))
                lhs146_ = d_1615_v2_
                lhs147_ = default__.safeIndex(489, (d_1615_v2_).length(0))
                r0 = rhs279_
                lhs144_[lhs145_] = rhs280_
                r0 = rhs281_
                lhs146_[lhs147_] = rhs282_
                def iife134_():
                    coll60_ = _dafny.Map()
                    compr_60_: _dafny.Seq
                    for compr_60_ in (default__.fm44(p0, (self).f0, (self).f0, d_1646_v26_.f5, globalState)).keys.Elements:
                        d_1667_v46_: _dafny.Seq = compr_60_
                        if (d_1667_v46_) in (default__.fm44(p0, (self).f0, (self).f0, d_1646_v26_.f5, globalState)):
                            coll60_[d_1667_v46_] = False
                    return _dafny.Map(coll60_)
                r0 = (p0) - (len((_dafny.Map({d_1647_v27_: d_1646_v26_.f5})) | (iife134_()
                )))
                d_1668_v47_: _dafny.Set
                d_1668_v47_ = _dafny.Set({p1})
                d_1669_v48_: _dafny.Map
                d_1669_v48_ = _dafny.Map({(d_1646_v26_.f4 if d_1646_v26_.f4 else d_1646_v26_.f5): (_dafny.Set({p0, 418})) | (d_1668_v47_)})
                d_1669_v48_ = (d_1669_v48_).set(d_1646_v26_.f4, d_1668_v47_)
                d_1670_v49_: C0
                nw313_ = C0()
                nw313_.ctor__(not((self.f1)[default__.safeIndex(262, (self.f1).length(0))]), d_1646_v26_.f5)
                d_1670_v49_ = nw313_
        hi10_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ltbe")))
        for d_1671_i4_ in range((0) - ((p1) - (len(d_1617_v3_))), hi10_):
            d_1672_v50_: C4
            nw314_ = C4()
            nw314_.ctor__(-578)
            d_1672_v50_ = nw314_
            d_1673_v51_: _dafny.Set
            d_1673_v51_ = _dafny.Set({p0, p0})
            d_1674_v53_: _dafny.Map
            def iife135_():
                coll61_ = _dafny.Set()
                compr_61_: int
                for compr_61_ in _dafny.IntegerRange(775, 402):
                    d_1675_v52_: int = compr_61_
                    if ((775) <= (d_1675_v52_)) and ((d_1675_v52_) < (402)):
                        coll61_ = coll61_.union(_dafny.Set([default__.safeDivisionInt(d_1675_v52_, p0)]))
                return _dafny.Set(coll61_)
            d_1674_v53_ = _dafny.Map({(self).fm6(p2, globalState): (d_1673_v51_).isdisjoint(iife135_()
            )})
            d_1674_v53_ = (d_1674_v53_).set((p2) == (p2), p2)
            d_1676_v54_: _dafny.Array
            nw315_ = _dafny.Array(int(0), 10)
            d_1676_v54_ = nw315_
            d_1676_v54_ = d_1676_v54_
            if p2:
                d_1677_v55_: _dafny.Map
                d_1677_v55_ = _dafny.Map({d_1613_v0_: d_1621_v5_})
                d_1678_v56_: _dafny.Map
                d_1678_v56_ = _dafny.Map({d_1677_v55_: d_1614_v1_})
                r0 = len(d_1678_v56_)
                d_1679_v57_: _dafny.Map
                d_1679_v57_ = _dafny.Map({len(d_1674_v53_): ((0) - (p1)) > (d_1671_i4_)})
                d_1679_v57_ = ((d_1679_v57_) | (d_1679_v57_)) | (d_1679_v57_)
                d_1680_v58_: _dafny.Set
                d_1680_v58_ = _dafny.Set({p2, p2, not(p2)})
                arr11_ = self.f1
                index307_ = default__.safeIndex(239, (self.f1).length(0))
                arr11_[index307_] = (d_1671_i4_) >= (len(d_1680_v58_))
                arr12_ = self.f1
                index308_ = default__.safeIndex(239, (self.f1).length(0))
                arr12_[index308_] = p2
                arr13_ = self.f1
                index309_ = default__.safeIndex(239, (self.f1).length(0))
                arr13_[index309_] = not(p2)
                arr14_ = self.f1
                index310_ = default__.safeIndex(239, (self.f1).length(0))
                rhs283_ = (self.f1)[default__.safeIndex(239, (self.f1).length(0))]
                rhs284_ = ((d_1672_v50_).fm1((self.f1)[default__.safeIndex(239, (self.f1).length(0))], globalState) if (self.f1)[default__.safeIndex(239, (self.f1).length(0))] else default__.safeDivisionInt(p0, d_1671_i4_))
                lhs148_ = self.f1
                lhs149_ = default__.safeIndex(239, (self.f1).length(0))
                lhs148_[lhs149_] = rhs283_
                r0 = rhs284_
            elif True:
                d_1681_v59_: bool
                d_1681_v59_ = False
                d_1682_v60_: _dafny.Map
                d_1682_v60_ = _dafny.Map({p0: p2})
                d_1683_v61_: _dafny.Map
                d_1683_v61_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1613_v0_ for d_1684_i5_ in range(default__.abs(191))]): d_1682_v60_})
                d_1685_v62_: D5
                d_1685_v62_ = D5_DC10(False, p1, d_1683_v61_, p2)
                d_1681_v59_ = (d_1685_v62_).cf17
                d_1686_v63_: C3
                nw316_ = C3()
                nw316_.ctor__(p1)
                d_1686_v63_ = nw316_
                d_1687_v64_: _dafny.Seq
                d_1687_v64_ = _dafny.SeqWithoutIsStrInference([d_1671_i4_])
                r0 = len(_dafny.Map({(d_1686_v63_).fm10(d_1671_i4_, globalState): d_1687_v64_}))
                r0 = -570
                d_1688_v65_: _dafny.Array
                nw317_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_1688_v65_ = nw317_
                d_1689_v66_: _dafny.Array
                nw318_ = _dafny.Array(None, 8)
                nw318_[int(0)] = D5_DC10(p2, p1, d_1683_v61_, d_1681_v59_)
                nw318_[int(1)] = d_1685_v62_
                nw318_[int(2)] = D5_DC10(p2, d_1671_i4_, d_1683_v61_, d_1681_v59_)
                nw318_[int(3)] = d_1685_v62_
                nw318_[int(4)] = D5_DC10(p2, len(_dafny.SeqWithoutIsStrInference([p1, d_1671_i4_])), d_1683_v61_, p2)
                nw318_[int(5)] = d_1685_v62_
                nw318_[int(6)] = d_1685_v62_
                nw318_[int(7)] = d_1685_v62_
                d_1689_v66_ = nw318_
                index311_ = default__.safeIndex(852, (d_1688_v65_).length(0))
                (d_1688_v65_)[index311_] = d_1689_v66_
                index312_ = default__.safeIndex(852, (d_1688_v65_).length(0))
                (d_1688_v65_)[index312_] = d_1689_v66_
        hi11_ = p1
        for d_1690_i6_ in range((58) + ((self).f0), hi11_):
            pat_let_tv35_ = d_1613_v0_
            d_1691_v67_: D0
            def iife136_(_pat_let37_0):
                def iife137_(d_1692_dt__update__tmp_h2_):
                    def iife138_(_pat_let38_0):
                        def iife139_(d_1693_dt__update_hcf3_h0_):
                            return D0_DC1((d_1692_dt__update__tmp_h2_).cf1, (d_1692_dt__update__tmp_h2_).cf2, d_1693_dt__update_hcf3_h0_, (d_1692_dt__update__tmp_h2_).cf4)
                        return iife139_(_pat_let38_0)
                    return iife138_(pat_let_tv35_)
                return iife137_(_pat_let37_0)
            d_1691_v67_ = D0_DC0((iife136_(D0_DC1(not(p2), p2, (d_1614_v1_)[default__.safeIndex(990, (d_1614_v1_).length(0))], p0))).cf2)
            d_1691_v67_ = d_1691_v67_
            d_1694_v68_: bool
            d_1694_v68_ = False
            d_1694_v68_ = p2
            d_1695_v69_: _dafny.Array
            nw319_ = _dafny.Array(_dafny.Seq({}), 18)
            d_1695_v69_ = nw319_
            d_1696_v70_: _dafny.Seq
            d_1696_v70_ = _dafny.SeqWithoutIsStrInference([d_1694_v68_, d_1694_v68_])
            d_1697_v71_: _dafny.Seq
            d_1697_v71_ = _dafny.SeqWithoutIsStrInference([d_1696_v70_])
            index313_ = default__.safeIndex(399, (d_1695_v69_).length(0))
            (d_1695_v69_)[index313_] = d_1697_v71_
            index314_ = default__.safeIndex(399, (d_1695_v69_).length(0))
            (d_1695_v69_)[index314_] = d_1697_v71_
            r0 = p0
        d_1698_i7_: int
        d_1698_i7_ = 0
        with _dafny.label("8"):
            while not(p2):
                with _dafny.c_label("8"):
                    if (d_1698_i7_) >= (100):
                        raise _dafny.Break("8")
                    d_1698_i7_ = (d_1698_i7_) + (1)
                    r0 = p1
                    if ((self).f0) == (len(d_1617_v3_)):
                        d_1699_v72_: _dafny.Array
                        def lambda67_(d_1700_p0_):
                            def lambda68_(d_1701_i8_):
                                return default__.safeModuloInt(d_1701_i8_, (0) - (d_1700_p0_))

                            return lambda68_

                        init36_ = lambda67_(p0)
                        nw320_ = _dafny.Array(None, 16)
                        for i0_36_ in range(nw320_.length(0)):
                            nw320_[i0_36_] = init36_(i0_36_)
                        d_1699_v72_ = nw320_
                        d_1702_v73_: D4
                        d_1702_v73_ = D4_DC6((self).f0, d_1699_v72_, p0)
                        d_1699_v72_ = (d_1702_v73_).cf10
                        d_1703_v74_: _dafny.Map
                        d_1703_v74_ = _dafny.Map({p2: (p0) != (331)})
                        d_1703_v74_ = (d_1703_v74_).set(p2, (d_1703_v74_) != (d_1703_v74_))
                        d_1704_v75_: _dafny.MultiSet
                        d_1704_v75_ = _dafny.MultiSet([p0, len((d_1615_v2_)[default__.safeIndex(489, (d_1615_v2_).length(0))])])
                        d_1705_v76_: _dafny.MultiSet
                        d_1705_v76_ = (d_1704_v75_) - (_dafny.MultiSet([len((d_1615_v2_)[default__.safeIndex(489, (d_1615_v2_).length(0))])]))
                        d_1705_v76_ = d_1705_v76_
                        d_1706_v77_: C2
                        nw321_ = C2()
                        nw321_.ctor__(p0)
                        d_1706_v77_ = nw321_
                        d_1707_v78_: _dafny.Seq
                        d_1707_v78_ = _dafny.SeqWithoutIsStrInference([205, p1])
                        d_1707_v78_ = d_1707_v78_
                    elif True:
                        d_1708_v79_: bool
                        d_1708_v79_ = False
                        d_1709_v80_: _dafny.Map
                        d_1709_v80_ = _dafny.Map({p0: p0})
                        d_1710_v81_: _dafny.Map
                        d_1710_v81_ = _dafny.Map({d_1709_v80_: p2})
                        d_1711_v82_: _dafny.Seq
                        d_1711_v82_ = _dafny.SeqWithoutIsStrInference([p2])
                        d_1712_v83_: _dafny.Map
                        d_1712_v83_ = _dafny.Map({_dafny.CodePoint('u'): p0})
                        rhs285_ = p2
                        rhs286_ = (d_1618_v4_) | (_dafny.Map({not((d_1711_v82_)[default__.safeIndex(p1, len(d_1711_v82_))]): len(d_1712_v83_)}))
                        rhs287_ = (_dafny.Map({d_1709_v80_: d_1708_v79_})) | (d_1710_v81_)
                        rhs288_ = d_1708_v79_
                        rhs289_ = ((len(d_1711_v82_)) + (p0)) - (p0)
                        d_1708_v79_ = rhs285_
                        d_1618_v4_ = rhs286_
                        d_1710_v81_ = rhs287_
                        d_1708_v79_ = rhs288_
                        r0 = rhs289_
                        d_1713_v84_: _dafny.MultiSet
                        d_1713_v84_ = _dafny.MultiSet([len((d_1617_v3_) + (((d_1615_v2_)[default__.safeIndex(489, (d_1615_v2_).length(0))]).set(default__.safeIndex(len(d_1709_v80_), len((d_1615_v2_)[default__.safeIndex(489, (d_1615_v2_).length(0))])), d_1613_v0_)))])
                        d_1713_v84_ = (d_1713_v84_).intersection(d_1713_v84_)
                        d_1708_v79_ = ((0) - (p0)) < (-473)
                        d_1708_v79_ = (532) < (343)
                        index315_ = default__.safeIndex(489, (d_1615_v2_).length(0))
                        (d_1615_v2_)[index315_] = d_1617_v3_
                    source24_ = d_1624_v8_
                    if source24_.is_DC21:
                        d_1714___mcc_h0_ = source24_.cf36
                        d_1715___mcc_h1_ = source24_.cf37
                        d_1716___mcc_h2_ = source24_.cf38
                        d_1717_cf38_ = d_1716___mcc_h2_
                        d_1718_cf37_ = d_1715___mcc_h1_
                        d_1719_cf36_ = d_1714___mcc_h0_
                        index316_ = default__.safeIndex(489, (d_1615_v2_).length(0))
                        (d_1615_v2_)[index316_] = d_1617_v3_
                        index317_ = default__.safeIndex(990, (d_1614_v1_).length(0))
                        (d_1614_v1_)[index317_] = ((d_1615_v2_)[default__.safeIndex(489, (d_1615_v2_).length(0))])[default__.safeIndex(d_1718_cf37_, len((d_1615_v2_)[default__.safeIndex(489, (d_1615_v2_).length(0))]))]
                        r0 = (0) - (d_1717_cf38_)
                        d_1720_v85_: bool
                        d_1720_v85_ = True
                        d_1721_v86_: _dafny.MultiSet
                        d_1721_v86_ = _dafny.MultiSet([len((d_1615_v2_)[default__.safeIndex(489, (d_1615_v2_).length(0))]), 717, p1])
                        d_1722_v87_: _dafny.MultiSet
                        d_1722_v87_ = _dafny.MultiSet([d_1721_v86_])
                        d_1720_v85_ = (d_1722_v87_) == (d_1722_v87_)
                    elif source24_.is_DC22:
                        d_1723___mcc_h3_ = source24_.cf39
                        d_1724_cf39_ = d_1723___mcc_h3_
                        index318_ = default__.safeIndex(990, (d_1614_v1_).length(0))
                        (d_1614_v1_)[index318_] = (d_1617_v3_)[default__.safeIndex(len(d_1617_v3_), len(d_1617_v3_))]
                        r0 = p0
                        d_1725_v88_: _dafny.Map
                        d_1725_v88_ = _dafny.Map({(self).f0: p2})
                        d_1724_cf39_ = (((self).fm1(not(p2), globalState) if not(p2) else (self).f0)) in ((d_1725_v88_) | (d_1725_v88_))
                        d_1725_v88_ = d_1725_v88_
                    elif True:
                        d_1726___mcc_h4_ = source24_.cf35
                        d_1727_cf35_ = d_1726___mcc_h4_
                        (self).f1 = d_1727_cf35_
                        r0 = (self).f0
                        d_1728_v89_: _dafny.Seq
                        d_1728_v89_ = _dafny.SeqWithoutIsStrInference([len((d_1615_v2_)[default__.safeIndex(489, (d_1615_v2_).length(0))])])
                        d_1729_v90_: _dafny.Map
                        d_1729_v90_ = _dafny.Map({len(d_1728_v89_): (self).fm1((self).fm6(p2, globalState), globalState)})
                        d_1730_v91_: _dafny.Seq
                        d_1730_v91_ = _dafny.SeqWithoutIsStrInference([p1, len(d_1729_v90_), p1, (self).f0])
                        r0 = default__.safeDivisionInt(p1, (d_1730_v91_)[default__.safeIndex(p1, len(d_1730_v91_))])
                        d_1731_v92_: _dafny.Set
                        d_1731_v92_ = _dafny.Set({p0})
                        rhs290_ = ((self).f0) + (len(d_1731_v92_))
                        rhs291_ = p1
                        r0 = rhs290_
                        r0 = rhs291_
                    d_1732_v93_: bool
                    d_1732_v93_ = False
                    d_1732_v93_ = not(not(True))
                    pass
            pass
        d_1733_v94_: _dafny.MultiSet
        d_1733_v94_ = _dafny.MultiSet([(self).f0, p1, p0])
        d_1734_v95_: _dafny.Set
        d_1734_v95_ = _dafny.Set({d_1613_v0_})
        r0 = ((d_1733_v94_)[(p1) * (len(d_1734_v95_))] if ((p1) * (len(d_1734_v95_))) in (d_1733_v94_) else len(d_1617_v3_))
        return r0

    @property
    def f0(self):
        return self._f0

class C6(T0, T1):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self):
        pass
        pass

    def fm0(self, globalState):
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({469: True})) for d_1735_i0_ in range(default__.abs(320))]): len(_dafny.Map({-916: False}))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: 818}))])): 937}))}))]): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxumvgapi")))}))) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([502 for d_1736_i1_ in range(default__.abs(77))]): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqtm")))}))

    def fm1(self, p0, globalState):
        return len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "up"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgl")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmdmkvq"))))

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        d_1737_v0_: _dafny.Map
        d_1737_v0_ = _dafny.Map({(689) != (p0): p0})
        d_1738_v1_: _dafny.Seq
        d_1738_v1_ = _dafny.SeqWithoutIsStrInference([p3])
        d_1739_v2_: _dafny.Map
        d_1739_v2_ = _dafny.Map({p0: p2})
        d_1740_v3_: _dafny.Map
        d_1740_v3_ = _dafny.Map({len(d_1738_v1_): len(d_1739_v2_)})
        d_1737_v0_ = (d_1737_v0_).set((p0) not in (d_1740_v3_), p0)
        d_1741_v4_: _dafny.Array
        nw322_ = _dafny.Array(int(0), 7)
        d_1741_v4_ = nw322_
        nw323_ = _dafny.Array(int(0), 13)
        d_1741_v4_ = nw323_
        if p3:
            r0 = (p0) >= (p0)
            d_1742_v5_: _dafny.Array
            nw324_ = _dafny.Array(_dafny.Array(None, 0), 4)
            d_1742_v5_ = nw324_
            d_1743_v6_: _dafny.Array
            nw325_ = _dafny.Array(False, 2)
            d_1743_v6_ = nw325_
            d_1744_v7_: _dafny.MultiSet
            d_1744_v7_ = _dafny.MultiSet([p2])
            index319_ = default__.safeIndex(894, (d_1743_v6_).length(0))
            (d_1743_v6_)[index319_] = (d_1744_v7_).issubset(d_1744_v7_)
            d_1745_v8_: _dafny.Map
            d_1745_v8_ = _dafny.Map({p2: d_1742_v5_})
            index320_ = default__.safeIndex(894, (d_1743_v6_).length(0))
            rhs292_ = ((d_1745_v8_)[p2] if (p2) in (d_1745_v8_) else d_1742_v5_)
            rhs293_ = default__.fm5(globalState)
            lhs150_ = d_1743_v6_
            lhs151_ = default__.safeIndex(894, (d_1743_v6_).length(0))
            d_1742_v5_ = rhs292_
            lhs150_[lhs151_] = rhs293_
            d_1746_v9_: _dafny.Seq
            d_1746_v9_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0])
            d_1747_v10_: _dafny.MultiSet
            d_1747_v10_ = _dafny.MultiSet([d_1746_v9_])
            d_1748_v11_: _dafny.Map
            d_1748_v11_ = _dafny.Map({d_1747_v10_: False})
            d_1748_v11_ = (d_1748_v11_).set(d_1747_v10_, default__.fm5(globalState))
            if p3:
                r0 = not(False)
                d_1749_v12_: _dafny.Array
                nw326_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
                d_1749_v12_ = nw326_
                d_1750_v13_: _dafny.Map
                d_1750_v13_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxtfidplh"))})
                d_1751_v14_: _dafny.Seq
                d_1751_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "to"))
                index321_ = default__.safeIndex(249, (d_1749_v12_).length(0))
                (d_1749_v12_)[index321_] = (((d_1750_v13_)[p0] if (p0) in (d_1750_v13_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))) + (d_1751_v14_)
                d_1752_v15_: _dafny.Array
                nw327_ = _dafny.Array(None, 7)
                nw327_[int(0)] = d_1741_v4_
                nw327_[int(1)] = d_1741_v4_
                nw327_[int(2)] = d_1741_v4_
                nw327_[int(3)] = d_1741_v4_
                nw327_[int(4)] = d_1741_v4_
                nw327_[int(5)] = (d_1741_v4_ if p2 else d_1741_v4_)
                nw327_[int(6)] = d_1741_v4_
                d_1752_v15_ = nw327_
                d_1753_v16_: _dafny.Seq
                d_1753_v16_ = _dafny.SeqWithoutIsStrInference([d_1752_v15_, d_1752_v15_])
                index322_ = default__.safeIndex(894, (d_1743_v6_).length(0))
                index323_ = default__.safeIndex(249, (d_1749_v12_).length(0))
                rhs294_ = p2
                rhs295_ = d_1751_v14_
                rhs296_ = (d_1753_v16_)[default__.safeIndex(p0, len(d_1753_v16_))]
                lhs152_ = d_1743_v6_
                lhs153_ = default__.safeIndex(894, (d_1743_v6_).length(0))
                lhs154_ = d_1749_v12_
                lhs155_ = default__.safeIndex(249, (d_1749_v12_).length(0))
                lhs152_[lhs153_] = rhs294_
                lhs154_[lhs155_] = rhs295_
                d_1752_v15_ = rhs296_
                r2 = (0) - ((d_1746_v9_)[default__.safeIndex((p0) - ((0) - (len((d_1749_v12_)[default__.safeIndex(249, (d_1749_v12_).length(0))]))), len(d_1746_v9_))])
                d_1754_v17_: str
                d_1754_v17_ = _dafny.CodePoint('r')
                d_1755_v18_: _dafny.Map
                d_1755_v18_ = _dafny.Map({d_1744_v7_: not((d_1743_v6_)[default__.safeIndex(894, (d_1743_v6_).length(0))])})
                index324_ = default__.safeIndex(894, (d_1743_v6_).length(0))
                rhs297_ = False
                rhs298_ = (d_1743_v6_)[default__.safeIndex(894, (d_1743_v6_).length(0))]
                rhs299_ = d_1754_v17_
                rhs300_ = d_1755_v18_
                lhs156_ = d_1743_v6_
                lhs157_ = default__.safeIndex(894, (d_1743_v6_).length(0))
                r0 = rhs297_
                lhs156_[lhs157_] = rhs298_
                d_1754_v17_ = rhs299_
                d_1755_v18_ = rhs300_
                d_1756_v19_: C3
                nw328_ = C3()
                nw328_.ctor__((p0) * (p0))
                d_1756_v19_ = nw328_
            elif True:
                d_1757_v20_: _dafny.Map
                d_1757_v20_ = _dafny.Map({d_1740_v3_: d_1741_v4_})
                index325_ = default__.safeIndex(894, (d_1743_v6_).length(0))
                (d_1743_v6_)[index325_] = (len(d_1757_v20_)) >= (p0)
                rhs301_ = d_1738_v1_
                rhs302_ = (p0) + ((0) - (p0))
                d_1738_v1_ = rhs301_
                r2 = rhs302_
                d_1758_v21_: _dafny.Array
                def lambda69_(d_1759_p2_):
                    def lambda70_(d_1760_i0_):
                        return (_dafny.CodePoint('c') if d_1759_p2_ else _dafny.CodePoint('u'))

                    return lambda70_

                init37_ = lambda69_(p2)
                nw329_ = _dafny.Array(None, 17)
                for i0_37_ in range(nw329_.length(0)):
                    nw329_[i0_37_] = init37_(i0_37_)
                d_1758_v21_ = nw329_
                d_1761_v22_: str
                d_1761_v22_ = _dafny.CodePoint('l')
                index326_ = default__.safeIndex(617, (d_1758_v21_).length(0))
                (d_1758_v21_)[index326_] = d_1761_v22_
                index327_ = default__.safeIndex(617, (d_1758_v21_).length(0))
                (d_1758_v21_)[index327_] = d_1761_v22_
                d_1762_v23_: _dafny.MultiSet
                d_1762_v23_ = _dafny.MultiSet([p0, p0, p0, p0])
                d_1763_v24_: _dafny.MultiSet
                d_1763_v24_ = d_1762_v23_
                d_1764_v25_: _dafny.Array
                nw330_ = _dafny.Array(None, 13)
                nw330_[int(0)] = default__.fm19(p0, p0, (d_1743_v6_)[default__.safeIndex(894, (d_1743_v6_).length(0))], p0, globalState)
                nw330_[int(1)] = d_1762_v23_
                nw330_[int(2)] = d_1762_v23_
                nw330_[int(3)] = d_1763_v24_
                nw330_[int(4)] = d_1763_v24_
                nw330_[int(5)] = d_1763_v24_
                nw330_[int(6)] = d_1763_v24_
                nw330_[int(7)] = d_1763_v24_
                nw330_[int(8)] = d_1762_v23_
                nw330_[int(9)] = d_1762_v23_
                nw330_[int(10)] = d_1763_v24_
                nw330_[int(11)] = d_1762_v23_
                nw330_[int(12)] = (d_1763_v24_ if p3 else d_1763_v24_)
                d_1764_v25_ = nw330_
                index328_ = default__.safeIndex(481, (d_1764_v25_).length(0))
                (d_1764_v25_)[index328_] = d_1763_v24_
                index329_ = default__.safeIndex(481, (d_1764_v25_).length(0))
                (d_1764_v25_)[index329_] = d_1762_v23_
                r0 = p2
            if (d_1743_v6_)[default__.safeIndex(894, (d_1743_v6_).length(0))]:
                d_1765_v26_: C4
                nw331_ = C4()
                nw331_.ctor__((0) - (p0))
                d_1765_v26_ = nw331_
                d_1766_v27_: C5
                nw332_ = C5()
                nw332_.ctor__(p0, d_1743_v6_)
                d_1766_v27_ = nw332_
                d_1767_v28_: D14
                d_1767_v28_ = D14_DC25(d_1742_v5_)
                d_1742_v5_ = (d_1767_v28_).cf42
                d_1768_v29_: C4
                nw333_ = C4()
                nw333_.ctor__((d_1766_v27_).f0)
                d_1768_v29_ = nw333_
                d_1769_v30_: _dafny.Map
                d_1769_v30_ = _dafny.Map({(d_1739_v2_) | (d_1739_v2_): p2})
                index330_ = default__.safeIndex(894, (d_1743_v6_).length(0))
                rhs303_ = d_1741_v4_
                rhs304_ = (d_1743_v6_)[default__.safeIndex(894, (d_1743_v6_).length(0))]
                rhs305_ = (0) - ((d_1765_v26_).fm1(p3, globalState))
                rhs306_ = d_1769_v30_
                rhs307_ = (d_1743_v6_)[default__.safeIndex(894, (d_1743_v6_).length(0))]
                lhs158_ = d_1743_v6_
                lhs159_ = default__.safeIndex(894, (d_1743_v6_).length(0))
                d_1741_v4_ = rhs303_
                r0 = rhs304_
                r2 = rhs305_
                d_1769_v30_ = rhs306_
                lhs158_[lhs159_] = rhs307_
            elif True:
                index331_ = default__.safeIndex(894, (d_1743_v6_).length(0))
                (d_1743_v6_)[index331_] = not(not(not(p2)))
                r2 = p0
                index332_ = default__.safeIndex(894, (d_1743_v6_).length(0))
                (d_1743_v6_)[index332_] = p2
                d_1770_v31_: _dafny.Seq
                d_1771_v32_: int
                d_1772_v33_: int
                d_1773_v34_: bool
                out28_: _dafny.Seq
                out29_: int
                out30_: int
                out31_: bool
                out28_, out29_, out30_, out31_ = (self).m4(True, (default__.fm42(p0, p2, p3, globalState)).cf21, p0, p3, globalState)
                d_1770_v31_ = out28_
                d_1771_v32_ = out29_
                d_1772_v33_ = out30_
                d_1773_v34_ = out31_
                d_1774_v35_: _dafny.Seq
                d_1774_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bevnw"))
                d_1774_v35_ = d_1774_v35_
        elif True:
            r0 = (p0) <= (p0)
            d_1775_v36_: _dafny.Array
            nw334_ = _dafny.Array(None, 2)
            nw334_[int(0)] = True
            nw334_[int(1)] = p2
            d_1775_v36_ = nw334_
            index333_ = default__.safeIndex(85, (d_1775_v36_).length(0))
            (d_1775_v36_)[index333_] = p3
            index334_ = default__.safeIndex(85, (d_1775_v36_).length(0))
            (d_1775_v36_)[index334_] = p3
            r2 = p0
            d_1776_v37_: str
            d_1776_v37_ = _dafny.CodePoint('a')
            d_1777_v38_: D0
            d_1777_v38_ = D0_DC1(not((d_1775_v36_)[default__.safeIndex(85, (d_1775_v36_).length(0))]), (d_1775_v36_)[default__.safeIndex(85, (d_1775_v36_).length(0))], d_1776_v37_, (p0) + (p0))
            d_1777_v38_ = d_1777_v38_
            d_1778_v39_: _dafny.Map
            d_1778_v39_ = _dafny.Map({True: d_1740_v3_})
            d_1779_v40_: _dafny.Map
            d_1779_v40_ = _dafny.Map({(d_1775_v36_)[default__.safeIndex(85, (d_1775_v36_).length(0))]: (d_1740_v3_) | (((d_1778_v39_)[(d_1775_v36_)[default__.safeIndex(85, (d_1775_v36_).length(0))]] if ((d_1775_v36_)[default__.safeIndex(85, (d_1775_v36_).length(0))]) in (d_1778_v39_) else _dafny.Map({(default__.fm38(p3, globalState)).cardinality: p0})))})
            d_1779_v40_ = (d_1779_v40_).set(default__.fm5(globalState), d_1740_v3_)
        hi12_ = 575
        for d_1780_i1_ in range((0) - (p0), hi12_):
            d_1781_v41_: str
            d_1781_v41_ = _dafny.CodePoint('r')
            rhs308_ = not(not((default__.fm45(p0, True, p2, d_1780_i1_, globalState)).cf39))
            rhs309_ = not(not((True) == (p3)))
            rhs310_ = d_1781_v41_
            rhs311_ = d_1780_i1_
            r0 = rhs308_
            r0 = rhs309_
            d_1781_v41_ = rhs310_
            r2 = rhs311_
            r0 = not(p2)
            r2 = p0
            d_1782_v42_: _dafny.Set
            d_1782_v42_ = _dafny.Set({p2})
            d_1783_v43_: D0
            d_1783_v43_ = D0_DC0(p3)
            d_1784_v44_: D14
            d_1784_v44_ = D14_DC26(d_1783_v43_, d_1782_v42_)
            d_1785_v45_: _dafny.Set
            d_1785_v45_ = _dafny.Set({(d_1782_v42_).intersection(d_1782_v42_), (d_1784_v44_).cf44, d_1782_v42_, _dafny.Set({False})})
            d_1785_v45_ = d_1785_v45_
        d_1786_i2_: int
        d_1786_i2_ = 0
        with _dafny.label("9"):
            while default__.fm5(globalState):
                with _dafny.c_label("9"):
                    if (d_1786_i2_) >= (100):
                        raise _dafny.Break("9")
                    d_1786_i2_ = (d_1786_i2_) + (1)
                    r2 = p0
                    d_1787_v46_: _dafny.Seq
                    d_1787_v46_ = _dafny.SeqWithoutIsStrInference([-589])
                    r0 = ((d_1787_v46_)[default__.safeIndex(p0, len(d_1787_v46_))]) == (p0)
                    r2 = default__.safeDivisionInt(((self).fm1(p2, globalState)) * (p0), p0)
                    d_1788_v47_: _dafny.MultiSet
                    d_1788_v47_ = _dafny.MultiSet([d_1741_v4_, d_1741_v4_])
                    d_1789_v48_: C0
                    nw335_ = C0()
                    nw335_.ctor__((_dafny.MultiSet([d_1741_v4_, d_1741_v4_])).ispropersubset(d_1788_v47_), p2)
                    d_1789_v48_ = nw335_
                    pass
            pass
        if (p0) >= (p0):
            d_1790_v49_: _dafny.Map
            d_1790_v49_ = _dafny.Map({not(default__.fm5(globalState)): not(p3)})
            d_1791_v50_: _dafny.Seq
            d_1791_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
            r0 = (((d_1790_v49_)[not(not(False))] if (not(not(False))) in (d_1790_v49_) else p3) if p2 else (d_1791_v50_) <= (d_1791_v50_))
            d_1792_v51_: _dafny.Set
            d_1792_v51_ = _dafny.Set({False, p3})
            r0 = ((d_1792_v51_) | (d_1792_v51_)).issubset((_dafny.Set({p3, p2, p2})) - (_dafny.Set({True, p2, p2})))
            if not ((d_1738_v1_) == (d_1738_v1_)) or (not(p3)):
                d_1793_v52_: T1
                nw336_ = C1()
                nw336_.ctor__(d_1791_v50_)
                d_1793_v52_ = nw336_
                d_1793_v52_ = d_1793_v52_
                r2 = p0
                d_1794_v53_: _dafny.Seq
                d_1795_v54_: int
                d_1796_v55_: int
                d_1797_v56_: bool
                out32_: _dafny.Seq
                out33_: int
                out34_: int
                out35_: bool
                out32_, out33_, out34_, out35_ = (self).m4(p2, d_1791_v50_, (self).fm1(p3, globalState), not((False if p3 else p3)), globalState)
                d_1794_v53_ = out32_
                d_1795_v54_ = out33_
                d_1796_v55_ = out34_
                d_1797_v56_ = out35_
                d_1797_v56_ = d_1797_v56_
                d_1798_v57_: _dafny.Map
                d_1798_v57_ = _dafny.Map({d_1738_v1_: len(_dafny.Map({d_1797_v56_: d_1741_v4_}))})
                d_1799_v58_: str
                d_1799_v58_ = _dafny.CodePoint('h')
                d_1800_v59_: _dafny.Set
                d_1800_v59_ = _dafny.Set({p0, len(d_1738_v1_), d_1796_v55_})
                d_1801_v60_: _dafny.Array
                nw337_ = _dafny.Array(None, 25)
                nw337_[int(0)] = (d_1798_v57_) == (d_1798_v57_)
                nw337_[int(1)] = not (p2) or (p3)
                nw337_[int(2)] = d_1797_v56_
                nw337_[int(3)] = (p2 if p3 else d_1797_v56_)
                nw337_[int(4)] = (not(p2)) == (False)
                nw337_[int(5)] = (p2 if not(p2) else default__.fm5(globalState))
                nw337_[int(6)] = d_1797_v56_
                nw337_[int(7)] = not(d_1797_v56_)
                nw337_[int(8)] = p2
                nw337_[int(9)] = ((d_1790_v49_)[p2] if (p2) in (d_1790_v49_) else p3)
                nw337_[int(10)] = p2
                nw337_[int(11)] = d_1797_v56_
                nw337_[int(12)] = d_1797_v56_
                nw337_[int(13)] = p3
                nw337_[int(14)] = p3
                nw337_[int(15)] = d_1797_v56_
                nw337_[int(16)] = p3
                nw337_[int(17)] = p3
                nw337_[int(18)] = (d_1799_v58_) not in (d_1791_v50_)
                nw337_[int(19)] = (_dafny.Set({d_1796_v55_})).isdisjoint(d_1800_v59_)
                nw337_[int(20)] = p2
                nw337_[int(21)] = not (p2) or (p2)
                nw337_[int(22)] = not((p3) not in (d_1738_v1_))
                nw337_[int(23)] = d_1797_v56_
                nw337_[int(24)] = p3
                d_1801_v60_ = nw337_
                index335_ = default__.safeIndex(326, (d_1801_v60_).length(0))
                (d_1801_v60_)[index335_] = True
                index336_ = default__.safeIndex(326, (d_1801_v60_).length(0))
                (d_1801_v60_)[index336_] = (d_1792_v51_).issubset(d_1792_v51_)
            elif True:
                index337_ = default__.safeIndex(9, (d_1741_v4_).length(0))
                (d_1741_v4_)[index337_] = default__.safeModuloInt(212, p0)
                d_1802_v61_: _dafny.Seq
                d_1802_v61_ = _dafny.SeqWithoutIsStrInference([d_1792_v51_])
                index338_ = default__.safeIndex(9, (d_1741_v4_).length(0))
                (d_1741_v4_)[index338_] = ((p0 if p2 else p0) if p2 else ((0) - (len(_dafny.SeqWithoutIsStrInference([p0 for d_1803_i3_ in range(default__.abs(404))]))) if p3 else len((d_1802_v61_)[default__.safeIndex(p0, len(d_1802_v61_))])))
                d_1804_v62_: _dafny.Array
                out36_: _dafny.Array
                out36_ = default__.m0(p2, p2, globalState)
                d_1804_v62_ = out36_
                d_1740_v3_ = (d_1740_v3_).set(default__.safeModuloInt(-585, (d_1741_v4_)[default__.safeIndex(9, (d_1741_v4_).length(0))]), p0)
                d_1805_v63_: _dafny.Array
                nw338_ = _dafny.Array(_dafny.Map({}), 19)
                d_1805_v63_ = nw338_
                d_1806_v64_: _dafny.Map
                d_1806_v64_ = _dafny.Map({d_1805_v63_: not(p3)})
                d_1806_v64_ = (d_1806_v64_).set(d_1805_v63_, False)
                d_1807_v65_: str
                d_1807_v65_ = _dafny.CodePoint('e')
                d_1807_v65_ = d_1807_v65_
            r2 = ((0) - (p0) if p3 else p0)
            d_1808_v66_: _dafny.MultiSet
            d_1808_v66_ = _dafny.MultiSet([default__.fm5(globalState)])
            if ((_dafny.CodePoint('q')) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1809_i4_ in range(default__.abs(767))]))) or ((((d_1808_v66_)[p2] if (p2) in (d_1808_v66_) else p0)) < (p0)):
                d_1810_v67_: D11
                d_1810_v67_ = D11_DC21(_dafny.CodePoint('s'), 65, p0)
                d_1811_v68_: str
                d_1811_v68_ = _dafny.CodePoint('i')
                d_1812_v70_: _dafny.Set
                d_1812_v70_ = _dafny.Set({(self).fm1(p3, globalState), p0, p0, p0})
                pat_let_tv36_ = d_1811_v68_
                d_1813_v71_: _dafny.Array
                nw339_ = _dafny.Array(None, 21)
                nw339_[int(0)] = d_1810_v67_
                nw339_[int(1)] = d_1810_v67_
                nw339_[int(2)] = (d_1810_v67_ if p2 else d_1810_v67_)
                nw339_[int(3)] = d_1810_v67_
                nw339_[int(4)] = d_1810_v67_
                def iife140_(_pat_let39_0):
                    def iife141_(d_1814_dt__update__tmp_h2_):
                        def iife142_(_pat_let40_0):
                            def iife143_(d_1815_dt__update_hcf36_h0_):
                                return D11_DC21(d_1815_dt__update_hcf36_h0_, (d_1814_dt__update__tmp_h2_).cf37, (d_1814_dt__update__tmp_h2_).cf38)
                            return iife143_(_pat_let40_0)
                        return iife142_(pat_let_tv36_)
                    return iife141_(_pat_let39_0)
                nw339_[int(5)] = (iife140_(d_1810_v67_) if p2 else D11_DC21(d_1811_v68_, p0, p0))
                nw339_[int(6)] = d_1810_v67_
                nw339_[int(7)] = d_1810_v67_
                nw339_[int(8)] = d_1810_v67_
                nw339_[int(9)] = D11_DC21(d_1811_v68_, p0, p0)
                nw339_[int(10)] = d_1810_v67_
                nw339_[int(11)] = d_1810_v67_
                nw339_[int(12)] = d_1810_v67_
                def iife144_():
                    coll62_ = _dafny.Set()
                    compr_62_: int
                    for compr_62_ in _dafny.IntegerRange(424, -760):
                        d_1816_v69_: int = compr_62_
                        if ((424) <= (d_1816_v69_)) and ((d_1816_v69_) < (-760)):
                            coll62_ = coll62_.union(_dafny.Set([default__.safeModuloInt(d_1816_v69_, (0) - (p0))]))
                    return _dafny.Set(coll62_)
                nw339_[int(13)] = default__.fm46(iife144_()
                , globalState)
                nw339_[int(14)] = D11_DC21(d_1811_v68_, p0, len(_dafny.Map({len(_dafny.Map({p0: len(d_1738_v1_)})): 836})))
                nw339_[int(15)] = D11_DC21(_dafny.CodePoint('o'), p0, p0)
                nw339_[int(16)] = (d_1810_v67_ if p3 else d_1810_v67_)
                nw339_[int(17)] = d_1810_v67_
                nw339_[int(18)] = d_1810_v67_
                nw339_[int(19)] = D11_DC21(d_1811_v68_, p0, p0)
                nw339_[int(20)] = default__.fm46(d_1812_v70_, globalState)
                d_1813_v71_ = nw339_
                index339_ = default__.safeIndex(484, (d_1813_v71_).length(0))
                (d_1813_v71_)[index339_] = d_1810_v67_
                index340_ = default__.safeIndex(134, (d_1741_v4_).length(0))
                (d_1741_v4_)[index340_] = p0
                d_1817_v72_: D4
                d_1817_v72_ = D4_DC6(p0, d_1741_v4_, p0)
                index341_ = default__.safeIndex(484, (d_1813_v71_).length(0))
                index342_ = default__.safeIndex(134, (d_1741_v4_).length(0))
                rhs312_ = d_1810_v67_
                rhs313_ = default__.safeModuloInt((d_1817_v72_).cf9, len(d_1740_v3_))
                lhs160_ = d_1813_v71_
                lhs161_ = default__.safeIndex(484, (d_1813_v71_).length(0))
                lhs162_ = d_1741_v4_
                lhs163_ = default__.safeIndex(134, (d_1741_v4_).length(0))
                lhs160_[lhs161_] = rhs312_
                lhs162_[lhs163_] = rhs313_
                d_1818_v73_: _dafny.Map
                d_1818_v73_ = _dafny.Map({(d_1739_v2_) | (d_1739_v2_): len(d_1791_v50_)})
                d_1818_v73_ = (d_1818_v73_).set(d_1739_v2_, 161)
                d_1819_v74_: _dafny.Array
                def lambda71_(d_1820_v68_):
                    def lambda72_(d_1821_i5_):
                        return d_1820_v68_

                    return lambda72_

                init38_ = lambda71_(d_1811_v68_)
                nw340_ = _dafny.Array(None, 9)
                for i0_38_ in range(nw340_.length(0)):
                    nw340_[i0_38_] = init38_(i0_38_)
                d_1819_v74_ = nw340_
                index343_ = default__.safeIndex(714, (d_1819_v74_).length(0))
                (d_1819_v74_)[index343_] = default__.fm23(globalState)
                index344_ = default__.safeIndex(714, (d_1819_v74_).length(0))
                (d_1819_v74_)[index344_] = default__.fm37((len(d_1791_v50_)) + ((d_1741_v4_)[default__.safeIndex(134, (d_1741_v4_).length(0))]), globalState)
                r0 = (p0) <= ((0) - (len(d_1791_v50_)))
                d_1822_v75_: _dafny.Array
                nw341_ = _dafny.Array(_dafny.Set({}), 6)
                d_1822_v75_ = nw341_
                d_1822_v75_ = (d_1822_v75_ if not(p2) else d_1822_v75_)
            elif True:
                d_1823_v76_: _dafny.Seq
                d_1823_v76_ = _dafny.SeqWithoutIsStrInference([p0])
                r0 = (default__.fm41(d_1823_v76_, p2, p2, p2, globalState)).issubset(d_1792_v51_)
                index345_ = default__.safeIndex(306, (p1).length(0))
                (p1)[index345_] = (d_1823_v76_) + (_dafny.SeqWithoutIsStrInference([p0]))
                index346_ = default__.safeIndex(306, (p1).length(0))
                (p1)[index346_] = d_1823_v76_
                r2 = (131) - (p0)
                d_1737_v0_ = (d_1737_v0_).set(((d_1739_v2_)[p0] if (p0) in (d_1739_v2_) else p2), p0)
                r0 = not(p2)
        elif True:
            r2 = (0) - (p0)
            r2 = (self).fm1(p2, globalState)
            d_1824_v77_: D0
            d_1824_v77_ = D0_DC0(p2)
            d_1825_v78_: _dafny.Set
            d_1825_v78_ = _dafny.Set({p2})
            d_1826_v79_: D14
            d_1826_v79_ = D14_DC26(d_1824_v77_, d_1825_v78_)
            d_1827_v80_: _dafny.Set
            d_1827_v80_ = _dafny.Set({d_1826_v79_, d_1826_v79_})
            d_1828_v81_: _dafny.Map
            d_1828_v81_ = _dafny.Map({p3: d_1827_v80_})
            if (d_1827_v80_).issubset(((d_1828_v81_)[not(False)] if (not(False)) in (d_1828_v81_) else d_1827_v80_)):
                index347_ = default__.safeIndex(589, (d_1741_v4_).length(0))
                (d_1741_v4_)[index347_] = (p0) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mig"))))
                d_1829_v82_: _dafny.Map
                d_1829_v82_ = _dafny.Map({d_1741_v4_: p0})
                d_1830_v83_: _dafny.Set
                d_1830_v83_ = _dafny.Set({p0})
                index348_ = default__.safeIndex(589, (d_1741_v4_).length(0))
                rhs314_ = ((self).fm1(p2, globalState)) - (p0)
                rhs315_ = len((d_1829_v82_) | ((_dafny.Map({d_1741_v4_: p0})).set(d_1741_v4_, len(d_1830_v83_))))
                rhs316_ = p3
                rhs317_ = p0
                lhs164_ = d_1741_v4_
                lhs165_ = default__.safeIndex(589, (d_1741_v4_).length(0))
                r2 = rhs314_
                lhs164_[lhs165_] = rhs315_
                r0 = rhs316_
                r2 = rhs317_
                d_1831_v84_: _dafny.Seq
                d_1831_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kaqatlgk"))
                d_1832_v85_: _dafny.Seq
                d_1832_v85_ = _dafny.SeqWithoutIsStrInference([(d_1741_v4_)[default__.safeIndex(589, (d_1741_v4_).length(0))], default__.safeModuloInt(p0, len(_dafny.Map({p0: 292}))), ((0) - (p0)) * (p0), default__.safeModuloInt(len(d_1831_v84_), 785)])
                d_1832_v85_ = _dafny.SeqWithoutIsStrInference([((d_1741_v4_)[default__.safeIndex(589, (d_1741_v4_).length(0))]) - ((d_1741_v4_)[default__.safeIndex(589, (d_1741_v4_).length(0))]), p0, p0])
                d_1833_v86_: C0
                nw342_ = C0()
                nw342_.ctor__(p3, p3)
                d_1833_v86_ = nw342_
                index349_ = default__.safeIndex(589, (d_1741_v4_).length(0))
                (d_1741_v4_)[index349_] = default__.safeModuloInt(((self).fm1(d_1833_v86_.f5, globalState) if d_1833_v86_.f5 else (d_1741_v4_)[default__.safeIndex(589, (d_1741_v4_).length(0))]), (d_1741_v4_)[default__.safeIndex(589, (d_1741_v4_).length(0))])
                d_1834_v87_: _dafny.Array
                nw343_ = _dafny.Array(int(0), 14)
                d_1834_v87_ = nw343_
                d_1834_v87_ = d_1834_v87_
            elif True:
                d_1835_v88_: _dafny.Array
                def lambda73_(d_1836_p3_):
                    def lambda74_(d_1837_i6_):
                        return d_1836_p3_

                    return lambda74_

                init39_ = lambda73_(p3)
                nw344_ = _dafny.Array(None, 12)
                for i0_39_ in range(nw344_.length(0)):
                    nw344_[i0_39_] = init39_(i0_39_)
                d_1835_v88_ = nw344_
                index350_ = default__.safeIndex(924, (d_1835_v88_).length(0))
                (d_1835_v88_)[index350_] = p3
                d_1838_v89_: _dafny.Map
                d_1838_v89_ = _dafny.Map({d_1825_v78_: p2})
                index351_ = default__.safeIndex(924, (d_1835_v88_).length(0))
                (d_1835_v88_)[index351_] = not(((d_1838_v89_)[(_dafny.Set({p2})) | (d_1825_v78_)] if ((_dafny.Set({p2})) | (d_1825_v78_)) in (d_1838_v89_) else not(default__.fm5(globalState))))
                d_1839_v90_: _dafny.Array
                def lambda75_(d_1840_i7_):
                    return _dafny.CodePoint('i')

                init40_ = lambda75_
                nw345_ = _dafny.Array(None, 28)
                for i0_40_ in range(nw345_.length(0)):
                    nw345_[i0_40_] = init40_(i0_40_)
                d_1839_v90_ = nw345_
                d_1841_v91_: str
                d_1841_v91_ = _dafny.CodePoint('w')
                index352_ = default__.safeIndex(648, (d_1839_v90_).length(0))
                (d_1839_v90_)[index352_] = d_1841_v91_
                index353_ = default__.safeIndex(648, (d_1839_v90_).length(0))
                (d_1839_v90_)[index353_] = d_1841_v91_
                index354_ = default__.safeIndex(924, (d_1835_v88_).length(0))
                (d_1835_v88_)[index354_] = p3
                d_1842_v92_: _dafny.MultiSet
                d_1842_v92_ = _dafny.MultiSet([not(p3)])
                index355_ = default__.safeIndex(924, (d_1835_v88_).length(0))
                index356_ = default__.safeIndex(924, (d_1835_v88_).length(0))
                rhs318_ = default__.fm5(globalState)
                rhs319_ = (p0) <= ((p0) + (144))
                rhs320_ = (d_1842_v92_) == (d_1842_v92_)
                lhs166_ = d_1835_v88_
                lhs167_ = default__.safeIndex(924, (d_1835_v88_).length(0))
                lhs168_ = d_1835_v88_
                lhs169_ = default__.safeIndex(924, (d_1835_v88_).length(0))
                lhs166_[lhs167_] = rhs318_
                r0 = rhs319_
                lhs168_[lhs169_] = rhs320_
                d_1843_v93_: D6
                d_1843_v93_ = D6_DC13(p2, p0)
                r0 = ((d_1843_v93_).cf24) and ((d_1835_v88_)[default__.safeIndex(924, (d_1835_v88_).length(0))])
            d_1844_v94_: _dafny.Array
            def lambda76_(d_1845_p0_):
                def lambda77_(d_1846_i8_):
                    return (_dafny.Set({len(_dafny.Map({d_1845_p0_: _dafny.CodePoint('o')})), d_1845_p0_, d_1845_p0_, d_1845_p0_, d_1845_p0_})).intersection(_dafny.Set({d_1845_p0_}))

                return lambda77_

            init41_ = lambda76_(p0)
            nw346_ = _dafny.Array(None, 22)
            for i0_41_ in range(nw346_.length(0)):
                nw346_[i0_41_] = init41_(i0_41_)
            d_1844_v94_ = nw346_
            d_1844_v94_ = d_1844_v94_
            if p2:
                r0 = default__.fm5(globalState)
                r2 = p0
                d_1847_v97_: D14
                d_1847_v97_ = D14_DC27(False)
                d_1848_v98_: _dafny.Map
                d_1848_v98_ = _dafny.Map({((d_1737_v0_)[p3] if (p3) in (d_1737_v0_) else p0): d_1847_v97_})
                d_1849_v100_: _dafny.Array
                nw347_ = _dafny.Array(None, 22)
                nw347_[int(0)] = d_1740_v3_
                nw347_[int(1)] = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([p0 for d_1850_i9_ in range(default__.abs(84))])): p0})
                nw347_[int(2)] = d_1740_v3_
                def iife145_():
                    coll63_ = _dafny.Map()
                    compr_63_: int
                    for compr_63_ in _dafny.IntegerRange(696, 369):
                        d_1851_v95_: int = compr_63_
                        if ((696) <= (d_1851_v95_)) and ((d_1851_v95_) < (369)):
                            coll63_[default__.safeDivisionInt(d_1851_v95_, len(d_1825_v78_))] = p0
                    return _dafny.Map(coll63_)
                nw347_[int(3)] = iife145_()
                
                nw347_[int(4)] = (d_1740_v3_).set(-183, p0)
                nw347_[int(5)] = d_1740_v3_
                nw347_[int(6)] = d_1740_v3_
                nw347_[int(7)] = d_1740_v3_
                nw347_[int(8)] = d_1740_v3_
                nw347_[int(9)] = d_1740_v3_
                nw347_[int(10)] = (d_1740_v3_).set(p0, p0)
                def iife146_():
                    coll64_ = _dafny.Map()
                    compr_64_: int
                    for compr_64_ in _dafny.IntegerRange(427, -816):
                        d_1852_v96_: int = compr_64_
                        if ((427) <= (d_1852_v96_)) and ((d_1852_v96_) < (-816)):
                            coll64_[(d_1852_v96_) * ((0) - (p0))] = p0
                    return _dafny.Map(coll64_)
                nw347_[int(11)] = iife146_()
                
                nw347_[int(12)] = d_1740_v3_
                nw347_[int(13)] = d_1740_v3_
                nw347_[int(14)] = d_1740_v3_
                nw347_[int(15)] = default__.fm27(p0, p3, globalState)
                nw347_[int(16)] = d_1740_v3_
                nw347_[int(17)] = _dafny.Map({len(d_1848_v98_): len(d_1825_v78_)})
                def iife147_():
                    coll65_ = _dafny.Map()
                    compr_65_: int
                    for compr_65_ in _dafny.IntegerRange(876, 548):
                        d_1853_v99_: int = compr_65_
                        if ((876) <= (d_1853_v99_)) and ((d_1853_v99_) < (548)):
                            coll65_[(d_1853_v99_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('f')])))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gg")))
                    return _dafny.Map(coll65_)
                nw347_[int(18)] = iife147_()
                
                nw347_[int(19)] = d_1740_v3_
                nw347_[int(20)] = d_1740_v3_
                nw347_[int(21)] = d_1740_v3_
                d_1849_v100_ = nw347_
                d_1854_v101_: _dafny.Array
                d_1854_v101_ = d_1849_v100_
                d_1854_v101_ = d_1854_v101_
                index357_ = default__.safeIndex(453, (d_1741_v4_).length(0))
                (d_1741_v4_)[index357_] = p0
                index358_ = default__.safeIndex(453, (d_1741_v4_).length(0))
                (d_1741_v4_)[index358_] = p0
                d_1855_v102_: _dafny.Set
                d_1855_v102_ = _dafny.Set({(d_1741_v4_)[default__.safeIndex(453, (d_1741_v4_).length(0))]})
                index359_ = default__.safeIndex(453, (d_1741_v4_).length(0))
                (d_1741_v4_)[index359_] = default__.safeDivisionInt(p0, (_dafny.MultiSet([d_1855_v102_, d_1855_v102_])).cardinality)
            elif True:
                index360_ = default__.safeIndex(917, (d_1741_v4_).length(0))
                (d_1741_v4_)[index360_] = (p0 if p2 else 835)
                d_1856_v103_: _dafny.Seq
                d_1856_v103_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1857_v104_: _dafny.Map
                d_1857_v104_ = _dafny.Map({d_1856_v103_: p3})
                index361_ = default__.safeIndex(917, (d_1741_v4_).length(0))
                (d_1741_v4_)[index361_] = len(d_1857_v104_)
                index362_ = default__.safeIndex(917, (d_1741_v4_).length(0))
                (d_1741_v4_)[index362_] = (0) - (default__.safeDivisionInt(p0, (d_1741_v4_)[default__.safeIndex(917, (d_1741_v4_).length(0))]))
                d_1858_v105_: _dafny.Array
                nw348_ = _dafny.Array(int(0), 5)
                d_1858_v105_ = nw348_
                def lambda78_(d_1859_v4_):
                    def lambda79_(d_1860_i10_):
                        return (d_1860_i10_) - ((d_1859_v4_)[default__.safeIndex(917, (d_1859_v4_).length(0))])

                    return lambda79_

                init42_ = lambda78_(d_1741_v4_)
                nw349_ = _dafny.Array(None, 14)
                for i0_42_ in range(nw349_.length(0)):
                    nw349_[i0_42_] = init42_(i0_42_)
                d_1858_v105_ = nw349_
                d_1861_v106_: _dafny.MultiSet
                d_1861_v106_ = _dafny.MultiSet([p3, not(p2)])
                d_1861_v106_ = default__.fm38((_dafny.Set({(d_1861_v106_).cardinality})).issubset(_dafny.Set({p0})), globalState)
                r0 = p3
        d_1862_v107_: D6
        d_1862_v107_ = D6_DC13(False, p0)
        pat_let_tv37_ = p2
        pat_let_tv38_ = p2
        def lambda80_(source25_):
            if source25_.is_DC12:
                d_1863___mcc_h0_ = source25_.cf22
                d_1864___mcc_h1_ = source25_.cf23
                d_1865_cf23_ = d_1864___mcc_h1_
                d_1866_cf22_ = d_1863___mcc_h0_
                return pat_let_tv37_
            elif source25_.is_DC13:
                d_1867___mcc_h2_ = source25_.cf24
                d_1868___mcc_h3_ = source25_.cf25
                d_1869_cf25_ = d_1868___mcc_h3_
                d_1870_cf24_ = d_1867___mcc_h2_
                return not(d_1870_cf24_)
            elif True:
                d_1871___mcc_h4_ = source25_.cf21
                d_1872_cf21_ = d_1871___mcc_h4_
                return pat_let_tv38_

        r0 = not(lambda80_(d_1862_v107_))
        nw350_ = _dafny.Array(_dafny.Map({}), 24)
        r1 = nw350_
        r2 = default__.safeModuloInt(p0, (self).fm1(p2, globalState))
        return r0, r1, r2

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_1873_v0_: bool
        d_1873_v0_ = False
        d_1874_v1_: _dafny.Array
        nw351_ = _dafny.Array(int(0), 21)
        d_1874_v1_ = nw351_
        index363_ = default__.safeIndex(636, (d_1874_v1_).length(0))
        (d_1874_v1_)[index363_] = -432
        index364_ = default__.safeIndex(636, (d_1874_v1_).length(0))
        rhs321_ = (p1) == ((p0) - (189))
        rhs322_ = p0
        lhs170_ = d_1874_v1_
        lhs171_ = default__.safeIndex(636, (d_1874_v1_).length(0))
        d_1873_v0_ = rhs321_
        lhs170_[lhs171_] = rhs322_
        d_1875_i0_: int
        d_1875_i0_ = 0
        with _dafny.label("10"):
            while d_1873_v0_:
                with _dafny.c_label("10"):
                    if (d_1875_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_1875_i0_ = (d_1875_i0_) + (1)
                    d_1874_v1_ = d_1874_v1_
                    d_1876_v2_: _dafny.MultiSet
                    d_1876_v2_ = _dafny.MultiSet([d_1873_v0_])
                    if ((_dafny.MultiSet([p2])).intersection(d_1876_v2_)).issubset((d_1876_v2_) | (d_1876_v2_)):
                        d_1877_v3_: str
                        d_1877_v3_ = _dafny.CodePoint('h')
                        d_1878_v4_: _dafny.MultiSet
                        d_1878_v4_ = _dafny.MultiSet([d_1877_v3_])
                        d_1879_v5_: _dafny.Seq
                        d_1879_v5_ = _dafny.SeqWithoutIsStrInference([(d_1878_v4_).cardinality, (d_1874_v1_)[default__.safeIndex(636, (d_1874_v1_).length(0))]])
                        d_1880_v6_: _dafny.Seq
                        d_1880_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dsqqc"))
                        d_1881_v7_: _dafny.Seq
                        d_1881_v7_ = _dafny.SeqWithoutIsStrInference([d_1880_v6_])
                        d_1882_v8_: _dafny.Seq
                        d_1882_v8_ = _dafny.SeqWithoutIsStrInference([d_1880_v6_, (d_1881_v7_)[default__.safeIndex((d_1874_v1_)[default__.safeIndex(636, (d_1874_v1_).length(0))], len(d_1881_v7_))]])
                        d_1883_v9_: _dafny.Map
                        d_1883_v9_ = _dafny.Map({p2: (d_1881_v7_)[default__.safeIndex(p1, len(d_1881_v7_))]})
                        d_1884_v10_: _dafny.Map
                        d_1884_v10_ = _dafny.Map({(d_1879_v5_)[default__.safeIndex(len(d_1883_v9_), len(d_1879_v5_))]: len(d_1880_v6_)})
                        d_1885_v11_: _dafny.Set
                        d_1885_v11_ = _dafny.Set({p2, d_1873_v0_, d_1873_v0_})
                        d_1886_v12_: _dafny.Seq
                        d_1886_v12_ = _dafny.SeqWithoutIsStrInference([d_1885_v11_, (d_1885_v11_) - (d_1885_v11_)])
                        index365_ = default__.safeIndex(636, (d_1874_v1_).length(0))
                        rhs323_ = _dafny.Map({(d_1874_v1_)[default__.safeIndex(636, (d_1874_v1_).length(0))]: p0})
                        rhs324_ = d_1886_v12_
                        rhs325_ = (0) - (len(default__.fm47((D11_DC21(d_1877_v3_, p1, p0)).cf36, p1, default__.safeDivisionInt((d_1874_v1_)[default__.safeIndex(636, (d_1874_v1_).length(0))], p0), globalState)))
                        rhs326_ = d_1873_v0_
                        rhs327_ = p2
                        lhs172_ = d_1874_v1_
                        lhs173_ = default__.safeIndex(636, (d_1874_v1_).length(0))
                        d_1884_v10_ = rhs323_
                        d_1886_v12_ = rhs324_
                        lhs172_[lhs173_] = rhs325_
                        d_1873_v0_ = rhs326_
                        d_1873_v0_ = rhs327_
                        d_1887_v13_: _dafny.MultiSet
                        d_1887_v13_ = _dafny.MultiSet([d_1874_v1_])
                        d_1888_v14_: D5
                        d_1888_v14_ = D5_DC8(d_1887_v13_)
                        d_1889_v15_: _dafny.Set
                        d_1889_v15_ = _dafny.Set({d_1888_v14_, d_1888_v14_})
                        d_1889_v15_ = d_1889_v15_
                        d_1890_v16_: _dafny.Map
                        d_1890_v16_ = _dafny.Map({p2: p2})
                        d_1890_v16_ = (d_1890_v16_).set(d_1873_v0_, (p0) not in (d_1884_v10_))
                        r0 = (0) - (((d_1876_v2_)[(p2) and (d_1873_v0_)] if ((p2) and (d_1873_v0_)) in (d_1876_v2_) else (0) - (p1)))
                        d_1891_v17_: _dafny.Array
                        nw352_ = _dafny.Array(False, 13)
                        d_1891_v17_ = nw352_
                        index366_ = default__.safeIndex(889, (d_1891_v17_).length(0))
                        (d_1891_v17_)[index366_] = d_1873_v0_
                        index367_ = default__.safeIndex(889, (d_1891_v17_).length(0))
                        (d_1891_v17_)[index367_] = p2
                    elif True:
                        d_1892_v18_: T2
                        nw353_ = C4()
                        nw353_.ctor__(p1)
                        d_1892_v18_ = nw353_
                        d_1893_v19_: _dafny.Map
                        d_1893_v19_ = _dafny.Map({d_1892_v18_: ((d_1874_v1_)[default__.safeIndex(636, (d_1874_v1_).length(0))]) + ((0) - (d_1892_v18_.f2))})
                        d_1894_v20_: _dafny.Seq
                        d_1894_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pftld"))
                        index368_ = default__.safeIndex(636, (d_1874_v1_).length(0))
                        (d_1874_v1_)[index368_] = ((d_1893_v19_)[d_1892_v18_] if (d_1892_v18_) in (d_1893_v19_) else default__.safeModuloInt(d_1892_v18_.f2, len(d_1894_v20_)))
                        d_1895_v21_: _dafny.Seq
                        d_1895_v21_ = _dafny.SeqWithoutIsStrInference([d_1876_v2_, d_1876_v2_])
                        d_1896_v22_: _dafny.Map
                        d_1896_v22_ = _dafny.Map({d_1873_v0_: 76})
                        d_1897_v23_: D5
                        d_1897_v23_ = D5_DC9(not(d_1873_v0_), len(d_1894_v20_), d_1876_v2_)
                        d_1898_v24_: _dafny.Set
                        d_1898_v24_ = _dafny.Set({d_1873_v0_})
                        d_1899_v25_: D4
                        d_1899_v25_ = D4_DC5(d_1898_v24_)
                        d_1900_v26_: _dafny.Set
                        d_1900_v26_ = _dafny.Set({(0) - (((d_1895_v21_)[default__.safeIndex(d_1892_v18_.f2, len(d_1895_v21_))]).cardinality), p0, len(d_1896_v22_), (d_1897_v23_).cf15, len((d_1899_v25_).cf8)})
                        d_1873_v0_ = (d_1900_v26_).ispropersubset(d_1900_v26_)
                        index369_ = default__.safeIndex(636, (d_1874_v1_).length(0))
                        (d_1874_v1_)[index369_] = (p0 if p2 else 339)
                        d_1901_v27_: _dafny.Map
                        d_1901_v27_ = _dafny.Map({False: True})
                        d_1902_v28_: _dafny.Seq
                        d_1902_v28_ = _dafny.SeqWithoutIsStrInference([not(p2), d_1873_v0_, d_1873_v0_])
                        d_1903_v29_: C0
                        nw354_ = C0()
                        nw354_.ctor__((((d_1901_v27_)[d_1873_v0_] if (d_1873_v0_) in (d_1901_v27_) else d_1873_v0_) if p2 else not(d_1873_v0_)), (d_1902_v28_)[default__.safeIndex(p1, len(d_1902_v28_))])
                        d_1903_v29_ = nw354_
                        (d_1903_v29_).f4 = (default__.safeModuloInt((d_1874_v1_)[default__.safeIndex(636, (d_1874_v1_).length(0))], p0)) != (p1)
                    d_1904_v30_: _dafny.Seq
                    d_1904_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfqqmub"))
                    d_1905_v31_: _dafny.MultiSet
                    d_1905_v31_ = _dafny.MultiSet([(d_1874_v1_)[default__.safeIndex(636, (d_1874_v1_).length(0))], p1])
                    d_1906_v32_: _dafny.Seq
                    d_1906_v32_ = _dafny.SeqWithoutIsStrInference([(d_1874_v1_)[default__.safeIndex(636, (d_1874_v1_).length(0))], p1, (d_1905_v31_).cardinality])
                    d_1907_v33_: _dafny.Map
                    d_1907_v33_ = _dafny.Map({d_1873_v0_: d_1906_v32_})
                    d_1908_v34_: _dafny.Map
                    d_1908_v34_ = _dafny.Map({(d_1876_v2_).cardinality: p2})
                    d_1909_v35_: _dafny.Map
                    d_1909_v35_ = _dafny.Map({d_1904_v30_: d_1908_v34_})
                    d_1910_v36_: D5
                    d_1910_v36_ = D5_DC10(d_1873_v0_, p1, d_1909_v35_, d_1873_v0_)
                    d_1911_v37_: C1
                    nw355_ = C1()
                    nw355_.ctor__((d_1904_v30_).set(default__.safeIndex(len(d_1907_v33_), len(d_1904_v30_)), default__.fm37((d_1910_v36_).cf18, globalState)))
                    d_1911_v37_ = nw355_
                    index370_ = default__.safeIndex(636, (d_1874_v1_).length(0))
                    (d_1874_v1_)[index370_] = ((0) - (p1) if d_1873_v0_ else p1)
                    pass
            pass
        d_1912_v38_: _dafny.Array
        def lambda81_(d_1913_i1_):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcqb"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxvrnpqcq")))

        init43_ = lambda81_
        nw356_ = _dafny.Array(None, 23)
        for i0_43_ in range(nw356_.length(0)):
            nw356_[i0_43_] = init43_(i0_43_)
        d_1912_v38_ = nw356_
        d_1914_v39_: _dafny.Seq
        d_1914_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "paxuh"))
        index371_ = default__.safeIndex(979, (d_1912_v38_).length(0))
        (d_1912_v38_)[index371_] = d_1914_v39_
        index372_ = default__.safeIndex(979, (d_1912_v38_).length(0))
        (d_1912_v38_)[index372_] = (d_1914_v39_) + (d_1914_v39_)
        index373_ = default__.safeIndex(636, (d_1874_v1_).length(0))
        (d_1874_v1_)[index373_] = p0
        d_1915_v40_: _dafny.Map
        d_1915_v40_ = _dafny.Map({len(d_1914_v39_): (_dafny.MultiSet([p2, p2])).cardinality})
        d_1916_v41_: T0
        nw357_ = C4()
        nw357_.ctor__((d_1874_v1_)[default__.safeIndex(636, (d_1874_v1_).length(0))])
        d_1916_v41_ = nw357_
        d_1917_v42_: _dafny.Seq
        d_1917_v42_ = _dafny.SeqWithoutIsStrInference([(d_1874_v1_)[default__.safeIndex(636, (d_1874_v1_).length(0))], ((d_1915_v40_)[len(_dafny.SeqWithoutIsStrInference([d_1916_v41_, d_1916_v41_]))] if (len(_dafny.SeqWithoutIsStrInference([d_1916_v41_, d_1916_v41_]))) in (d_1915_v40_) else p1)])
        d_1918_v43_: C2
        nw358_ = C2()
        nw358_.ctor__(p1)
        d_1918_v43_ = nw358_
        d_1919_v44_: _dafny.Map
        d_1919_v44_ = _dafny.Map({d_1873_v0_: d_1918_v43_})
        d_1920_v45_: _dafny.Map
        d_1920_v45_ = _dafny.Map({d_1873_v0_: (d_1917_v42_)[default__.safeIndex((0) - (len(d_1919_v44_)), len(d_1917_v42_))]})
        d_1921_v46_: _dafny.Seq
        d_1921_v46_ = _dafny.SeqWithoutIsStrInference([True])
        d_1920_v45_ = (d_1920_v45_).set((d_1921_v46_)[default__.safeIndex(p1, len(d_1921_v46_))], p1)
        d_1922_v47_: bool
        d_1923_v48_: int
        out37_: bool
        out38_: int
        out37_, out38_ = (d_1918_v43_).m6(len((d_1912_v38_)[default__.safeIndex(979, (d_1912_v38_).length(0))]), globalState)
        d_1922_v47_ = out37_
        d_1923_v48_ = out38_
        r0 = p1
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_1924_v0_: C1
        nw359_ = C1()
        nw359_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ui")))
        d_1924_v0_ = nw359_
        d_1925_v1_: D8
        d_1925_v1_ = D8_DC17(p2, d_1924_v0_, (0) - ((p0).cardinality))
        source26_ = d_1925_v1_
        if source26_.is_DC17:
            d_1926___mcc_h0_ = source26_.cf30
            d_1927___mcc_h1_ = source26_.cf31
            d_1928___mcc_h2_ = source26_.cf32
            d_1929_cf32_ = d_1928___mcc_h2_
            d_1930_cf31_ = d_1927___mcc_h1_
            d_1931_cf30_ = d_1926___mcc_h0_
            d_1932_v2_: D4
            d_1932_v2_ = D4_DC6(862, p1, d_1929_cf32_)
            d_1933_v3_: D4
            d_1933_v3_ = D4_DC7(d_1932_v2_)
            d_1934_v4_: _dafny.Seq
            d_1934_v4_ = _dafny.SeqWithoutIsStrInference([d_1932_v2_])
            d_1935_v5_: _dafny.Map
            d_1935_v5_ = _dafny.Map({not(p2): p3})
            d_1936_v6_: D4
            d_1936_v6_ = D4_DC7((d_1934_v4_)[default__.safeIndex(len(d_1935_v5_), len(d_1934_v4_))])
            d_1937_v7_: D4
            d_1937_v7_ = D4_DC7((d_1936_v6_).cf12)
            source27_ = d_1937_v7_
            if source27_.is_DC6:
                d_1938___mcc_h4_ = source27_.cf9
                d_1939___mcc_h5_ = source27_.cf10
                d_1940___mcc_h6_ = source27_.cf11
                d_1941_cf11_ = d_1940___mcc_h6_
                d_1942_cf10_ = d_1939___mcc_h5_
                d_1943_cf9_ = d_1938___mcc_h4_
                d_1944_v8_: _dafny.Map
                d_1944_v8_ = _dafny.Map({(d_1943_cf9_) <= (d_1941_cf11_): _dafny.CodePoint('u')})
                rhs328_ = p3
                rhs329_ = d_1944_v8_
                d_1931_cf30_ = rhs328_
                d_1944_v8_ = rhs329_
                d_1945_v9_: str
                d_1945_v9_ = _dafny.CodePoint('d')
                d_1946_v10_: _dafny.Array
                def lambda82_(d_1947_p3_):
                    def lambda83_(d_1948_i0_):
                        return not(d_1947_p3_)

                    return lambda83_

                init44_ = lambda82_(p3)
                nw360_ = _dafny.Array(None, 17)
                for i0_44_ in range(nw360_.length(0)):
                    nw360_[i0_44_] = init44_(i0_44_)
                d_1946_v10_ = nw360_
                d_1949_v11_: C5
                nw361_ = C5()
                nw361_.ctor__((len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wywgutya"))).set(default__.safeIndex(d_1943_cf9_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wywgutya")))), d_1945_v9_)) if d_1931_cf30_ else (0) - (d_1943_cf9_)), d_1946_v10_)
                d_1949_v11_ = nw361_
                d_1950_v12_: _dafny.Array
                def lambda84_(d_1951_v0_):
                    def lambda85_(d_1952_i1_):
                        return ((D15_DC29(_dafny.Map({(d_1951_v0_).f3: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "luwy"))}))).cf46) | (_dafny.Map({(d_1951_v0_).f3: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "avaemdf"))}))

                    return lambda85_

                init45_ = lambda84_(d_1924_v0_)
                nw362_ = _dafny.Array(None, 21)
                for i0_45_ in range(nw362_.length(0)):
                    nw362_[i0_45_] = init45_(i0_45_)
                d_1950_v12_ = nw362_
                d_1953_v13_: _dafny.Map
                d_1953_v13_ = _dafny.Map({(d_1930_cf31_).f3: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgu"))})
                index374_ = default__.safeIndex(148, (d_1950_v12_).length(0))
                (d_1950_v12_)[index374_] = d_1953_v13_
                index375_ = default__.safeIndex(148, (d_1950_v12_).length(0))
                (d_1950_v12_)[index375_] = d_1953_v13_
                d_1931_cf30_ = p3
            elif source27_.is_DC5:
                d_1954___mcc_h7_ = source27_.cf8
                d_1955_cf8_ = d_1954___mcc_h7_
                r0 = d_1929_cf32_
                d_1956_v14_: T2
                nw363_ = C4()
                nw363_.ctor__(d_1929_cf32_)
                d_1956_v14_ = nw363_
                d_1957_v15_: _dafny.Map
                d_1957_v15_ = _dafny.Map({d_1956_v14_: d_1956_v14_.f2})
                d_1958_v16_: _dafny.Seq
                d_1958_v16_ = _dafny.SeqWithoutIsStrInference([(0) - (d_1929_cf32_), d_1929_cf32_, len((d_1957_v15_).set(d_1956_v14_, d_1956_v14_.f2)), 232, d_1929_cf32_])
                d_1929_cf32_ = (d_1956_v14_.f2 if (len(d_1958_v16_)) == (d_1956_v14_.f2) else len((d_1924_v0_).f3))
                d_1959_v17_: _dafny.Set
                d_1959_v17_ = _dafny.Set({(d_1924_v0_).fm1(d_1931_cf30_, globalState)})
                rhs330_ = (d_1929_cf32_) == (d_1929_cf32_)
                rhs331_ = (d_1959_v17_).intersection(d_1959_v17_)
                d_1931_cf30_ = rhs330_
                d_1959_v17_ = rhs331_
                d_1960_v18_: _dafny.Seq
                d_1960_v18_ = _dafny.SeqWithoutIsStrInference([p3, d_1931_cf30_, not(p3), d_1931_cf30_])
                d_1961_v19_: _dafny.Seq
                d_1961_v19_ = _dafny.SeqWithoutIsStrInference([d_1960_v18_, default__.fm16(globalState)])
                rhs332_ = d_1956_v14_.f2
                rhs333_ = (0) - (d_1929_cf32_)
                rhs334_ = (d_1961_v19_) + (d_1961_v19_)
                rhs335_ = p2
                rhs336_ = False
                lhs174_ = d_1956_v14_
                r0 = rhs332_
                lhs174_.f2 = rhs333_
                d_1961_v19_ = rhs334_
                d_1931_cf30_ = rhs335_
                d_1931_cf30_ = rhs336_
            elif True:
                d_1962___mcc_h8_ = source27_.cf12
                d_1963_cf12_ = d_1962___mcc_h8_
                d_1964_v20_: _dafny.Map
                d_1964_v20_ = _dafny.Map({d_1929_cf32_: (d_1929_cf32_ if p3 else d_1929_cf32_)})
                d_1964_v20_ = (d_1964_v20_).set((default__.fm38(p3, globalState)).cardinality, d_1929_cf32_)
                d_1965_v21_: _dafny.Array
                def lambda86_(d_1966_p2_):
                    def lambda87_(d_1967_i2_):
                        return not(d_1966_p2_)

                    return lambda87_

                init46_ = lambda86_(p2)
                nw364_ = _dafny.Array(None, 26)
                for i0_46_ in range(nw364_.length(0)):
                    nw364_[i0_46_] = init46_(i0_46_)
                d_1965_v21_ = nw364_
                d_1965_v21_ = d_1965_v21_
                d_1968_v22_: _dafny.Seq
                d_1968_v22_ = _dafny.SeqWithoutIsStrInference([p3, p2])
                d_1969_v23_: _dafny.Array
                nw365_ = _dafny.Array(_dafny.Seq({}), 18)
                d_1969_v23_ = nw365_
                d_1970_v24_: bool
                d_1971_v25_: _dafny.Array
                d_1972_v26_: int
                out39_: bool
                out40_: _dafny.Array
                out41_: int
                out39_, out40_, out41_ = (d_1924_v0_).m1((len((d_1968_v22_).set(default__.safeIndex(552, len(d_1968_v22_)), p2))) - (-947), d_1969_v23_, p3, p3, globalState)
                d_1970_v24_ = out39_
                d_1971_v25_ = out40_
                d_1972_v26_ = out41_
                d_1973_v27_: _dafny.Set
                d_1973_v27_ = _dafny.Set({d_1931_cf30_})
                d_1974_v28_: _dafny.Seq
                d_1974_v28_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_1970_v24_}), d_1973_v27_, d_1973_v27_, d_1973_v27_, d_1973_v27_])
                d_1975_v29_: _dafny.Seq
                d_1975_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbfxcksb"))
                rhs337_ = ((_dafny.SeqWithoutIsStrInference([_dafny.Set({p2}), d_1973_v27_])) + (d_1974_v28_)).set(default__.safeIndex(d_1929_cf32_, len((_dafny.SeqWithoutIsStrInference([_dafny.Set({p2}), d_1973_v27_])) + (d_1974_v28_))), d_1973_v27_)
                rhs338_ = default__.fm15(globalState)
                rhs339_ = (len((d_1924_v0_).f3)) * (d_1929_cf32_)
                d_1974_v28_ = rhs337_
                d_1975_v29_ = rhs338_
                r0 = rhs339_
            r0 = len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({d_1931_cf30_})))]))
            d_1976_v30_: _dafny.Array
            nw366_ = _dafny.Array(False, 3)
            d_1976_v30_ = nw366_
            index376_ = default__.safeIndex(429, (d_1976_v30_).length(0))
            (d_1976_v30_)[index376_] = p3
            index377_ = default__.safeIndex(429, (d_1976_v30_).length(0))
            (d_1976_v30_)[index377_] = p2
            r1 = -385
        elif True:
            d_1977___mcc_h3_ = source26_.cf29
            d_1978_cf29_ = d_1977___mcc_h3_
            d_1979_v31_: _dafny.Array
            nw367_ = _dafny.Array(False, 21)
            d_1979_v31_ = nw367_
            d_1980_v32_: int
            d_1980_v32_ = 246
            index378_ = default__.safeIndex(888, (d_1979_v31_).length(0))
            (d_1979_v31_)[index378_] = (d_1980_v32_) > (len((d_1924_v0_).f3))
            d_1981_v33_: bool
            d_1981_v33_ = True
            d_1982_v34_: _dafny.Map
            d_1982_v34_ = _dafny.Map({(d_1980_v32_) >= (d_1980_v32_): not(p3)})
            d_1983_v35_: str
            d_1983_v35_ = _dafny.CodePoint('k')
            d_1984_v36_: _dafny.Map
            d_1984_v36_ = _dafny.Map({d_1983_v35_: -702})
            d_1985_v37_: _dafny.Map
            d_1985_v37_ = _dafny.Map({len(d_1984_v36_): p3})
            index379_ = default__.safeIndex(888, (d_1979_v31_).length(0))
            rhs340_ = ((d_1982_v34_)[p3] if (p3) in (d_1982_v34_) else ((d_1985_v37_)[d_1980_v32_] if (d_1980_v32_) in (d_1985_v37_) else p2))
            rhs341_ = d_1981_v33_
            lhs175_ = d_1979_v31_
            lhs176_ = default__.safeIndex(888, (d_1979_v31_).length(0))
            lhs175_[lhs176_] = rhs340_
            d_1981_v33_ = rhs341_
            index380_ = default__.safeIndex(888, (d_1979_v31_).length(0))
            (d_1979_v31_)[index380_] = not((p0).isdisjoint(p0))
            d_1986_v38_: _dafny.Seq
            d_1986_v38_ = _dafny.SeqWithoutIsStrInference([p3, p2])
            d_1987_v39_: _dafny.Array
            out42_: _dafny.Array
            out42_ = default__.m0((len(d_1986_v38_)) <= ((self).fm1(default__.fm5(globalState), globalState)), ((d_1978_cf29_).f3) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prkosnxd"))), globalState)
            d_1987_v39_ = out42_
            d_1981_v33_ = d_1981_v33_
        d_1988_v40_: bool
        d_1988_v40_ = True
        d_1989_v41_: _dafny.Array
        def lambda88_(d_1990_i3_):
            return not(True)

        init47_ = lambda88_
        nw368_ = _dafny.Array(None, 4)
        for i0_47_ in range(nw368_.length(0)):
            nw368_[i0_47_] = init47_(i0_47_)
        d_1989_v41_ = nw368_
        d_1991_v42_: _dafny.Seq
        d_1991_v42_ = _dafny.SeqWithoutIsStrInference([d_1989_v41_])
        d_1992_v43_: D14
        d_1992_v43_ = D14_DC27(p3)
        pat_let_tv39_ = d_1988_v40_
        pat_let_tv40_ = p3
        pat_let_tv41_ = d_1924_v0_
        pat_let_tv42_ = d_1924_v0_
        pat_let_tv43_ = p2
        pat_let_tv44_ = p3
        pat_let_tv45_ = p2
        def lambda89_(source28_):
            if source28_.is_DC26:
                d_1993___mcc_h9_ = source28_.cf43
                d_1994___mcc_h10_ = source28_.cf44
                d_1995_cf44_ = d_1994___mcc_h10_
                d_1996_cf43_ = d_1993___mcc_h9_
                return ((D0_DC1(pat_let_tv39_, pat_let_tv40_, _dafny.CodePoint('s'), len((pat_let_tv41_).f3))).cf4) >= (len(_dafny.Set({len((pat_let_tv42_).f3)})))
            elif source28_.is_DC27:
                d_1997___mcc_h11_ = source28_.cf45
                d_1998_cf45_ = d_1997___mcc_h11_
                return (len(_dafny.Set({d_1998_cf45_}))) >= (len(_dafny.SeqWithoutIsStrInference([d_1998_cf45_, pat_let_tv43_])))
            elif source28_.is_DC28:
                return pat_let_tv44_
            elif True:
                d_1999___mcc_h12_ = source28_.cf42
                d_2000_cf42_ = d_1999___mcc_h12_
                return pat_let_tv45_

        rhs342_ = lambda89_(d_1992_v43_)
        rhs343_ = (_dafny.SeqWithoutIsStrInference([d_1989_v41_, d_1989_v41_, d_1989_v41_, d_1989_v41_])) + (d_1991_v42_)
        d_1988_v40_ = rhs342_
        d_1991_v42_ = rhs343_
        d_2001_v44_: _dafny.Set
        d_2001_v44_ = _dafny.Set({p1, p1, p1})
        d_2002_v45_: _dafny.Seq
        d_2002_v45_ = _dafny.SeqWithoutIsStrInference([p3])
        d_2003_v46_: _dafny.Map
        d_2003_v46_ = _dafny.Map({d_2002_v45_: d_1988_v40_})
        d_2004_v47_: int
        d_2004_v47_ = 420
        d_2005_v48_: _dafny.Set
        d_2005_v48_ = _dafny.Set({len((d_2003_v46_) | (d_2003_v46_)), d_2004_v47_})
        d_2006_v49_: _dafny.Map
        d_2006_v49_ = _dafny.Map({_dafny.MultiSet([p2, not(True)]): d_2001_v44_})
        d_2007_v50_: _dafny.Map
        d_2007_v50_ = _dafny.Map({len((d_1924_v0_).f3): d_2004_v47_})
        d_2008_v51_: _dafny.Map
        d_2008_v51_ = _dafny.Map({d_1989_v41_: ((d_2007_v50_)[d_2004_v47_] if (d_2004_v47_) in (d_2007_v50_) else d_2004_v47_)})
        rhs344_ = default__.safeModuloInt(d_2004_v47_, d_2004_v47_)
        rhs345_ = (((d_2006_v49_)[(_dafny.MultiSet(d_2002_v45_)).set(p3, default__.abs(d_2004_v47_))] if ((_dafny.MultiSet(d_2002_v45_)).set(p3, default__.abs(d_2004_v47_))) in (d_2006_v49_) else d_2001_v44_)) | (d_2001_v44_)
        rhs346_ = (default__.fm48(d_1988_v40_, globalState))
        rhs347_ = len(d_2008_v51_)
        r0 = rhs344_
        d_2001_v44_ = rhs345_
        d_2005_v48_ = rhs346_
        r0 = rhs347_
        d_2009_v52_: _dafny.Seq
        d_2009_v52_ = _dafny.SeqWithoutIsStrInference([d_2004_v47_])
        d_2010_v53_: D7
        d_2010_v53_ = D7_DC14((_dafny.SeqWithoutIsStrInference([d_2004_v47_])) + (d_2009_v52_))
        source29_ = d_2010_v53_
        if source29_.is_DC15:
            d_2011___mcc_h13_ = source29_.cf27
            d_2012___mcc_h14_ = source29_.cf28
            d_2013_cf28_ = d_2012___mcc_h14_
            d_2014_cf27_ = d_2011___mcc_h13_
            index381_ = default__.safeIndex(59, (d_1989_v41_).length(0))
            (d_1989_v41_)[index381_] = d_1988_v40_
            index382_ = default__.safeIndex(59, (d_1989_v41_).length(0))
            (d_1989_v41_)[index382_] = d_1988_v40_
            d_2015_v54_: C4
            nw369_ = C4()
            nw369_.ctor__(d_2014_cf27_)
            d_2015_v54_ = nw369_
            d_2016_v55_: _dafny.Array
            def lambda90_(d_2017_p2_):
                def lambda91_(d_2018_i4_):
                    return d_2017_p2_

                return lambda91_

            init48_ = lambda90_(p2)
            nw370_ = _dafny.Array(None, 27)
            for i0_48_ in range(nw370_.length(0)):
                nw370_[i0_48_] = init48_(i0_48_)
            d_2016_v55_ = nw370_
            d_2019_v56_: _dafny.Array
            nw371_ = _dafny.Array(None, 4)
            nw371_[int(0)] = d_2016_v55_
            nw371_[int(1)] = d_2016_v55_
            nw371_[int(2)] = d_2016_v55_
            nw371_[int(3)] = d_2016_v55_
            d_2019_v56_ = nw371_
            d_2020_v57_: D14
            d_2020_v57_ = D14_DC25(d_2019_v56_)
            pat_let_tv46_ = d_2019_v56_
            def iife148_(_pat_let41_0):
                def iife149_(d_2021_dt__update__tmp_h0_):
                    def iife150_(_pat_let42_0):
                        def iife151_(d_2022_dt__update_hcf42_h0_):
                            return D14_DC25(d_2022_dt__update_hcf42_h0_)
                        return iife151_(_pat_let42_0)
                    return iife150_(pat_let_tv46_)
                return iife149_(_pat_let41_0)
            source30_ = iife148_(d_2020_v57_)
            if source30_.is_DC26:
                d_2023___mcc_h16_ = source30_.cf43
                d_2024___mcc_h17_ = source30_.cf44
                d_2025_cf44_ = d_2024___mcc_h17_
                d_2026_cf43_ = d_2023___mcc_h16_
                d_2004_v47_ = (default__.safeModuloInt(d_2014_cf27_, len(_dafny.Map({d_2004_v47_: p3})))) - (d_2004_v47_)
                r1 = (d_2004_v47_) - ((self).fm1((d_1989_v41_)[default__.safeIndex(59, (d_1989_v41_).length(0))], globalState))
                d_2027_v58_: _dafny.Map
                d_2027_v58_ = _dafny.Map({p1: p2})
                d_2027_v58_ = (d_2027_v58_).set(p1, (d_1989_v41_)[default__.safeIndex(59, (d_1989_v41_).length(0))])
                index383_ = default__.safeIndex(59, (d_1989_v41_).length(0))
                (d_1989_v41_)[index383_] = p3
            elif source30_.is_DC27:
                d_2028___mcc_h18_ = source30_.cf45
                d_2029_cf45_ = d_2028___mcc_h18_
                d_2004_v47_ = default__.safeDivisionInt(d_2014_cf27_, d_2014_cf27_)
                d_2030_v59_: C1
                nw372_ = C1()
                nw372_.ctor__((_dafny.SeqWithoutIsStrInference([d_2013_cf28_ for d_2031_i5_ in range(default__.abs(-916))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "du"))))
                d_2030_v59_ = nw372_
                r1 = default__.safeModuloInt((0) - ((d_2004_v47_) - (d_2014_cf27_)), (0) - (len(d_1991_v42_)))
                d_2032_v60_: _dafny.Array
                nw373_ = _dafny.Array(None, 6)
                nw373_[int(0)] = d_2004_v47_
                nw373_[int(1)] = d_2004_v47_
                nw373_[int(2)] = -630
                nw373_[int(3)] = d_2014_cf27_
                nw373_[int(4)] = len(_dafny.Set({d_2013_cf28_, d_2013_cf28_, d_2013_cf28_}))
                nw373_[int(5)] = 358
                d_2032_v60_ = nw373_
                d_2032_v60_ = p1
            elif source30_.is_DC28:
                d_1988_v40_ = (d_2002_v45_)[default__.safeIndex(len((d_1924_v0_).f3), len(d_2002_v45_))]
                d_2033_v61_: _dafny.MultiSet
                d_2033_v61_ = p0
                d_2034_v62_: _dafny.Map
                d_2034_v62_ = _dafny.Map({d_2033_v61_: d_2013_cf28_})
                d_2035_v63_: _dafny.Set
                d_2035_v63_ = _dafny.Set({p3, True, default__.fm5(globalState), p3})
                d_2036_v64_: D4
                d_2036_v64_ = D4_DC5(d_2035_v63_)
                d_2037_v65_: _dafny.Map
                d_2037_v65_ = _dafny.Map({d_2034_v62_: d_2036_v64_})
                d_2037_v65_ = _dafny.Map({d_2034_v62_: d_2036_v64_})
                d_2038_v66_: C1
                nw374_ = C1()
                nw374_.ctor__((d_1924_v0_).f3)
                d_2038_v66_ = nw374_
                index384_ = default__.safeIndex(59, (d_1989_v41_).length(0))
                (d_1989_v41_)[index384_] = (d_1989_v41_)[default__.safeIndex(59, (d_1989_v41_).length(0))]
            elif True:
                d_2039___mcc_h19_ = source30_.cf42
                d_2040_cf42_ = d_2039___mcc_h19_
                d_2041_v67_: _dafny.Array
                nw375_ = _dafny.Array(_dafny.Seq({}), 8)
                d_2041_v67_ = nw375_
                d_2042_v68_: bool
                d_2043_v69_: _dafny.Array
                d_2044_v70_: int
                out43_: bool
                out44_: _dafny.Array
                out45_: int
                out43_, out44_, out45_ = (d_2015_v54_).m1(d_2014_cf27_, d_2041_v67_, not((d_2002_v45_) <= (d_2002_v45_)), p2, globalState)
                d_2042_v68_ = out43_
                d_2043_v69_ = out44_
                d_2044_v70_ = out45_
                d_2002_v45_ = d_2002_v45_
                d_2045_v71_: _dafny.Seq
                d_2045_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdfxyi"))
                d_2045_v71_ = ((d_1924_v0_).f3) + (((d_1924_v0_).f3).set(default__.safeIndex(-474, len((d_1924_v0_).f3)), d_2013_cf28_))
                d_2046_v72_: C3
                nw376_ = C3()
                nw376_.ctor__(((self).fm1(p2, globalState)) + (d_2014_cf27_))
                d_2046_v72_ = nw376_
            d_1988_v40_ = (p2 if (d_1989_v41_)[default__.safeIndex(59, (d_1989_v41_).length(0))] else p3)
        elif True:
            d_2047___mcc_h15_ = source29_.cf26
            d_2048_cf26_ = d_2047___mcc_h15_
            d_2049_v73_: _dafny.Map
            d_2049_v73_ = _dafny.Map({(d_1924_v0_).f3: p2})
            d_2050_v74_: _dafny.Map
            d_2050_v74_ = _dafny.Map({d_2004_v47_: d_1988_v40_})
            d_2049_v73_ = (d_2049_v73_).set((d_1924_v0_).f3, (d_1988_v40_) and (((d_2050_v74_)[d_2004_v47_] if (d_2004_v47_) in (d_2050_v74_) else p3)))
            d_2051_v75_: C2
            nw377_ = C2()
            nw377_.ctor__((0) - (d_2004_v47_))
            d_2051_v75_ = nw377_
            d_2052_v76_: _dafny.Seq
            d_2052_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fuumxjor"))
            d_2053_v77_: str
            d_2053_v77_ = _dafny.CodePoint('x')
            d_2052_v76_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_2054_i6_ in range(default__.abs(158))])) + (((d_2052_v76_).set(default__.safeIndex(len(d_2048_cf26_), len(d_2052_v76_)), ((d_1924_v0_).f3)[default__.safeIndex(d_2004_v47_, len((d_1924_v0_).f3))])).set(default__.safeIndex(d_2004_v47_, len((d_2052_v76_).set(default__.safeIndex(len(d_2048_cf26_), len(d_2052_v76_)), ((d_1924_v0_).f3)[default__.safeIndex(d_2004_v47_, len((d_1924_v0_).f3))]))), d_2053_v77_))
            d_2004_v47_ = d_2004_v47_
        d_2055_v78_: str
        d_2055_v78_ = _dafny.CodePoint('n')
        d_2056_v79_: _dafny.Array
        nw378_ = _dafny.Array(None, 22)
        nw378_[int(0)] = (d_1924_v0_).f3
        nw378_[int(1)] = ((d_1924_v0_).f3).set(default__.safeIndex(d_2004_v47_, len((d_1924_v0_).f3)), d_2055_v78_)
        nw378_[int(2)] = (d_1924_v0_).f3
        nw378_[int(3)] = (d_1924_v0_).f3
        nw378_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxjtpgk"))
        nw378_[int(5)] = (d_1924_v0_).f3
        nw378_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcgsqra"))
        nw378_[int(7)] = (d_1924_v0_).f3
        nw378_[int(8)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2057_i7_ in range(default__.abs(-609))])
        nw378_[int(9)] = (d_1924_v0_).f3
        nw378_[int(10)] = (d_1924_v0_).f3
        nw378_[int(11)] = (d_1924_v0_).f3
        nw378_[int(12)] = (d_1924_v0_).f3
        nw378_[int(13)] = _dafny.SeqWithoutIsStrInference([d_2055_v78_ for d_2058_i8_ in range(default__.abs(365))])
        nw378_[int(14)] = (d_1924_v0_).f3
        nw378_[int(15)] = (d_1924_v0_).f3
        nw378_[int(16)] = (d_1924_v0_).f3
        nw378_[int(17)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eoy"))).set(default__.safeIndex(d_2004_v47_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eoy")))), _dafny.CodePoint('l'))
        nw378_[int(18)] = (d_1924_v0_).f3
        nw378_[int(19)] = (d_1924_v0_).f3
        nw378_[int(20)] = (d_1924_v0_).f3
        nw378_[int(21)] = (d_1924_v0_).f3
        d_2056_v79_ = nw378_
        d_2059_v80_: _dafny.Map
        d_2059_v80_ = _dafny.Map({d_2004_v47_: d_2009_v52_})
        d_2060_v81_: _dafny.Array
        nw379_ = _dafny.Array(None, 29)
        nw379_[int(0)] = d_2009_v52_
        nw379_[int(1)] = _dafny.SeqWithoutIsStrInference([61 for d_2061_i9_ in range(default__.abs(-225))])
        nw379_[int(2)] = d_2009_v52_
        nw379_[int(3)] = d_2009_v52_
        nw379_[int(4)] = d_2009_v52_
        nw379_[int(5)] = default__.fm36(p2, d_2055_v78_, len(d_2005_v48_), globalState)
        nw379_[int(6)] = d_2009_v52_
        nw379_[int(7)] = d_2009_v52_
        nw379_[int(8)] = default__.fm36(p3, _dafny.CodePoint('b'), d_2004_v47_, globalState)
        nw379_[int(9)] = d_2009_v52_
        nw379_[int(10)] = d_2009_v52_
        nw379_[int(11)] = d_2009_v52_
        nw379_[int(12)] = _dafny.SeqWithoutIsStrInference([len((d_1924_v0_).f3)])
        nw379_[int(13)] = d_2009_v52_
        nw379_[int(14)] = _dafny.SeqWithoutIsStrInference([d_2004_v47_])
        nw379_[int(15)] = d_2009_v52_
        nw379_[int(16)] = _dafny.SeqWithoutIsStrInference([d_2004_v47_ for d_2062_i10_ in range(default__.abs(-172))])
        nw379_[int(17)] = d_2009_v52_
        nw379_[int(18)] = _dafny.SeqWithoutIsStrInference([len(d_2009_v52_)])
        nw379_[int(19)] = (d_2009_v52_).set(default__.safeIndex(d_2004_v47_, len(d_2009_v52_)), 5)
        nw379_[int(20)] = d_2009_v52_
        nw379_[int(21)] = d_2009_v52_
        nw379_[int(22)] = ((d_2059_v80_)[d_2004_v47_] if (d_2004_v47_) in (d_2059_v80_) else d_2009_v52_)
        nw379_[int(23)] = _dafny.SeqWithoutIsStrInference([d_2004_v47_, d_2004_v47_])
        nw379_[int(24)] = d_2009_v52_
        nw379_[int(25)] = d_2009_v52_
        nw379_[int(26)] = _dafny.SeqWithoutIsStrInference([len(d_2007_v50_) for d_2063_i11_ in range(default__.abs(697))])
        nw379_[int(27)] = d_2009_v52_
        nw379_[int(28)] = _dafny.SeqWithoutIsStrInference([(d_1924_v0_).fm1(True, globalState)])
        d_2060_v81_ = nw379_
        d_2064_v82_: _dafny.Map
        d_2064_v82_ = _dafny.Map({d_2056_v79_: d_2060_v81_})
        d_2065_v83_: bool
        d_2066_v84_: _dafny.Array
        d_2067_v85_: int
        out46_: bool
        out47_: _dafny.Array
        out48_: int
        out46_, out47_, out48_ = (d_1924_v0_).m1(d_2004_v47_, ((d_2064_v82_)[d_2056_v79_] if (d_2056_v79_) in (d_2064_v82_) else d_2060_v81_), p2, p2, globalState)
        d_2065_v83_ = out46_
        d_2066_v84_ = out47_
        d_2067_v85_ = out48_
        d_2068_v86_: _dafny.Map
        d_2068_v86_ = _dafny.Map({d_1989_v41_: ((d_1924_v0_).f3)[default__.safeIndex(-578, len((d_1924_v0_).f3))]})
        d_2069_v87_: _dafny.Map
        d_2069_v87_ = _dafny.Map({d_2068_v86_: _dafny.Map({d_2002_v45_: d_2004_v47_})})
        d_2070_v88_: _dafny.Map
        d_2070_v88_ = _dafny.Map({d_2002_v45_: d_2004_v47_})
        d_2069_v87_ = (d_2069_v87_).set((d_2068_v86_) | (d_2068_v86_), d_2070_v88_)
        d_2071_v89_: _dafny.Set
        d_2071_v89_ = _dafny.Set({d_2010_v53_})
        r0 = len(((d_2071_v89_) | (_dafny.Set({d_2010_v53_}))) | (d_2071_v89_))
        d_2072_v90_: _dafny.Seq
        d_2072_v90_ = _dafny.SeqWithoutIsStrInference([d_2005_v48_, d_2005_v48_])
        r1 = (len(((d_1924_v0_).f3).set(default__.safeIndex(len(d_2072_v90_), len((d_1924_v0_).f3)), d_2055_v78_))) * (((0) - (d_2067_v85_) if (d_2002_v45_)[default__.safeIndex(d_2004_v47_, len(d_2002_v45_))] else (self).fm1(p2, globalState)))
        return r0, r1

    def m4(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        r2: int = int(0)
        r3: bool = False
        if p3:
            d_2073_v0_: _dafny.Array
            nw380_ = _dafny.Array(int(0), 28)
            d_2073_v0_ = nw380_
            index385_ = default__.safeIndex(756, (d_2073_v0_).length(0))
            (d_2073_v0_)[index385_] = len((p1) + (p1))
            d_2074_v1_: _dafny.Seq
            d_2074_v1_ = _dafny.SeqWithoutIsStrInference([p3, not(p3), not(p3), p0])
            index386_ = default__.safeIndex(756, (d_2073_v0_).length(0))
            (d_2073_v0_)[index386_] = len((d_2074_v1_) + (_dafny.SeqWithoutIsStrInference([p3])))
            r3 = p3
            d_2075_v2_: _dafny.Seq
            d_2075_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iawbfjbim"))
            d_2076_v3_: str
            d_2076_v3_ = _dafny.CodePoint('h')
            d_2077_v4_: _dafny.Map
            d_2077_v4_ = _dafny.Map({len(_dafny.Map({d_2076_v3_: p0})): _dafny.SeqWithoutIsStrInference([d_2076_v3_ for d_2078_i0_ in range(default__.abs(811))])})
            d_2075_v2_ = (d_2075_v2_) + (((d_2077_v4_)[(d_2073_v0_)[default__.safeIndex(756, (d_2073_v0_).length(0))]] if ((d_2073_v0_)[default__.safeIndex(756, (d_2073_v0_).length(0))]) in (d_2077_v4_) else d_2075_v2_))
            d_2079_v5_: _dafny.Array
            nw381_ = _dafny.Array(False, 25)
            d_2079_v5_ = nw381_
            index387_ = default__.safeIndex(536, (d_2079_v5_).length(0))
            (d_2079_v5_)[index387_] = not(not(True))
            d_2080_v6_: _dafny.MultiSet
            d_2080_v6_ = _dafny.MultiSet([8])
            d_2081_v7_: _dafny.MultiSet
            d_2081_v7_ = _dafny.MultiSet([d_2080_v6_])
            d_2082_v8_: _dafny.Seq
            d_2082_v8_ = _dafny.SeqWithoutIsStrInference([p2])
            index388_ = default__.safeIndex(536, (d_2079_v5_).length(0))
            (d_2079_v5_)[index388_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffeyfgi"))) < (p1) if p0 else (d_2081_v7_).issubset(_dafny.MultiSet([_dafny.MultiSet(d_2082_v8_)])))
            d_2083_v9_: _dafny.Array
            nw382_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_2083_v9_ = nw382_
            index389_ = default__.safeIndex(777, (d_2083_v9_).length(0))
            (d_2083_v9_)[index389_] = d_2073_v0_
            index390_ = default__.safeIndex(777, (d_2083_v9_).length(0))
            (d_2083_v9_)[index390_] = d_2073_v0_
        elif True:
            d_2084_v10_: _dafny.Set
            d_2084_v10_ = _dafny.Set({p2})
            d_2085_v11_: _dafny.MultiSet
            d_2085_v11_ = _dafny.MultiSet([p0])
            d_2086_v12_: D5
            d_2086_v12_ = D5_DC9(p0, p2, d_2085_v11_)
            d_2087_v13_: _dafny.Array
            nw383_ = _dafny.Array(None, 22)
            nw383_[int(0)] = p0
            nw383_[int(1)] = p3
            nw383_[int(2)] = p3
            nw383_[int(3)] = p3
            nw383_[int(4)] = p0
            nw383_[int(5)] = p3
            nw383_[int(6)] = p3
            nw383_[int(7)] = default__.fm5(globalState)
            nw383_[int(8)] = p3
            nw383_[int(9)] = p0
            nw383_[int(10)] = p3
            nw383_[int(11)] = p0
            nw383_[int(12)] = p3
            nw383_[int(13)] = p3
            nw383_[int(14)] = p0
            nw383_[int(15)] = p3
            nw383_[int(16)] = p0
            nw383_[int(17)] = p0
            nw383_[int(18)] = not((d_2086_v12_).cf14)
            nw383_[int(19)] = p3
            nw383_[int(20)] = p3
            nw383_[int(21)] = default__.fm5(globalState)
            d_2087_v13_ = nw383_
            d_2088_v14_: C5
            nw384_ = C5()
            nw384_.ctor__(len(d_2084_v10_), d_2087_v13_)
            d_2088_v14_ = nw384_
            d_2089_v15_: _dafny.Array
            nw385_ = _dafny.Array(int(0), 25)
            d_2089_v15_ = nw385_
            d_2090_v16_: _dafny.Seq
            d_2090_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yghybw"))
            rhs348_ = (d_2089_v15_ if p0 else d_2089_v15_)
            rhs349_ = p2
            rhs350_ = d_2090_v16_
            d_2089_v15_ = rhs348_
            r2 = rhs349_
            d_2090_v16_ = rhs350_
            index391_ = default__.safeIndex(471, (d_2089_v15_).length(0))
            (d_2089_v15_)[index391_] = p2
            index392_ = default__.safeIndex(471, (d_2089_v15_).length(0))
            rhs351_ = d_2088_v14_.f1
            rhs352_ = p3
            rhs353_ = p2
            lhs177_ = d_2088_v14_
            lhs178_ = d_2089_v15_
            lhs179_ = default__.safeIndex(471, (d_2089_v15_).length(0))
            lhs177_.f1 = rhs351_
            r3 = rhs352_
            lhs178_[lhs179_] = rhs353_
            if not (True) or (not (p0) or (p3)):
                index393_ = default__.safeIndex(471, (d_2089_v15_).length(0))
                (d_2089_v15_)[index393_] = (d_2088_v14_).f0
                index394_ = default__.safeIndex(471, (d_2089_v15_).length(0))
                (d_2089_v15_)[index394_] = (d_2088_v14_).f0
                d_2091_v17_: C2
                nw386_ = C2()
                nw386_.ctor__(p2)
                d_2091_v17_ = nw386_
                d_2092_v18_: _dafny.Map
                d_2092_v18_ = _dafny.Map({p3: d_2090_v16_})
                r2 = (len((default__.fm15(globalState)) + (((d_2092_v18_)[not(not(p0))] if (not(not(p0))) in (d_2092_v18_) else p1)))) * (p2)
                r3 = p0
            elif True:
                r3 = p3
                d_2093_v19_: C3
                nw387_ = C3()
                nw387_.ctor__((d_2089_v15_)[default__.safeIndex(471, (d_2089_v15_).length(0))])
                d_2093_v19_ = nw387_
                d_2094_v20_: str
                d_2094_v20_ = _dafny.CodePoint('s')
                d_2094_v20_ = d_2094_v20_
                d_2095_v21_: _dafny.Seq
                d_2095_v21_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2096_v22_: _dafny.Map
                d_2096_v22_ = _dafny.Map({p3: p3})
                r3 = (d_2095_v21_)[default__.safeIndex(len((d_2096_v22_ if p0 else d_2096_v22_)), len(d_2095_v21_))]
                d_2089_v15_ = d_2089_v15_
            r3 = (p0 if p3 else p3)
        if p0:
            d_2097_v23_: _dafny.Array
            nw388_ = _dafny.Array(int(0), 6)
            d_2097_v23_ = nw388_
            index395_ = default__.safeIndex(1, (d_2097_v23_).length(0))
            (d_2097_v23_)[index395_] = (890) + (p2)
            index396_ = default__.safeIndex(1, (d_2097_v23_).length(0))
            (d_2097_v23_)[index396_] = p2
            index397_ = default__.safeIndex(1, (d_2097_v23_).length(0))
            (d_2097_v23_)[index397_] = p2
            d_2098_v24_: _dafny.Map
            d_2098_v24_ = _dafny.Map({not(p0): p0})
            index398_ = default__.safeIndex(1, (d_2097_v23_).length(0))
            (d_2097_v23_)[index398_] = len((d_2098_v24_) | (d_2098_v24_))
            d_2099_v25_: C2
            nw389_ = C2()
            nw389_.ctor__((d_2097_v23_)[default__.safeIndex(1, (d_2097_v23_).length(0))])
            d_2099_v25_ = nw389_
            d_2100_v26_: D6
            d_2100_v26_ = D6_DC11(p1)
            source31_ = d_2100_v26_
            if source31_.is_DC12:
                d_2101___mcc_h0_ = source31_.cf22
                d_2102___mcc_h1_ = source31_.cf23
                d_2103_cf23_ = d_2102___mcc_h1_
                d_2104_cf22_ = d_2101___mcc_h0_
                d_2105_v27_: _dafny.Seq
                d_2105_v27_ = _dafny.SeqWithoutIsStrInference([d_2104_cf22_, d_2104_cf22_])
                d_2106_v28_: _dafny.Array
                nw390_ = _dafny.Array(None, 6)
                nw390_[int(0)] = default__.fm5(globalState)
                nw390_[int(1)] = d_2104_cf22_
                nw390_[int(2)] = d_2103_cf23_
                nw390_[int(3)] = p0
                nw390_[int(4)] = (d_2105_v27_)[default__.safeIndex(len(_dafny.Map({p2: d_2103_cf23_})), len(d_2105_v27_))]
                nw390_[int(5)] = d_2103_cf23_
                d_2106_v28_ = nw390_
                d_2107_v29_: C5
                nw391_ = C5()
                nw391_.ctor__(p2, (d_2106_v28_ if p3 else d_2106_v28_))
                d_2107_v29_ = nw391_
                arr15_ = d_2107_v29_.f1
                index399_ = default__.safeIndex(301, (d_2107_v29_.f1).length(0))
                arr15_[index399_] = d_2103_cf23_
                d_2108_v30_: _dafny.Map
                d_2108_v30_ = _dafny.Map({(151) - (32): p0})
                arr16_ = d_2107_v29_.f1
                index400_ = default__.safeIndex(301, (d_2107_v29_.f1).length(0))
                arr16_[index400_] = ((d_2108_v30_)[(d_2097_v23_)[default__.safeIndex(1, (d_2097_v23_).length(0))]] if ((d_2097_v23_)[default__.safeIndex(1, (d_2097_v23_).length(0))]) in (d_2108_v30_) else d_2103_cf23_)
                d_2103_cf23_ = (d_2107_v29_.f1)[default__.safeIndex(301, (d_2107_v29_.f1).length(0))]
                d_2109_v31_: str
                d_2109_v31_ = _dafny.CodePoint('l')
                d_2109_v31_ = d_2109_v31_
            elif source31_.is_DC13:
                d_2110___mcc_h2_ = source31_.cf24
                d_2111___mcc_h3_ = source31_.cf25
                d_2112_cf25_ = d_2111___mcc_h3_
                d_2113_cf24_ = d_2110___mcc_h2_
                d_2114_v32_: _dafny.Map
                d_2114_v32_ = _dafny.Map({default__.fm16(globalState): p3})
                r1 = default__.safeDivisionInt(len(d_2114_v32_), default__.safeModuloInt((d_2097_v23_)[default__.safeIndex(1, (d_2097_v23_).length(0))], 908))
                index401_ = default__.safeIndex(1, (d_2097_v23_).length(0))
                (d_2097_v23_)[index401_] = (d_2099_v25_).fm1(p0, globalState)
                d_2115_v33_: C1
                nw392_ = C1()
                nw392_.ctor__(p1)
                d_2115_v33_ = nw392_
                d_2115_v33_ = d_2115_v33_
                r3 = True
            elif True:
                d_2116___mcc_h4_ = source31_.cf21
                d_2117_cf21_ = d_2116___mcc_h4_
                rhs354_ = 853
                rhs355_ = ((d_2097_v23_)[default__.safeIndex(1, (d_2097_v23_).length(0))]) < (((d_2097_v23_)[default__.safeIndex(1, (d_2097_v23_).length(0))]) * ((d_2097_v23_)[default__.safeIndex(1, (d_2097_v23_).length(0))]))
                rhs356_ = d_2097_v23_
                rhs357_ = (((self).fm1(p0, globalState) if p0 else (d_2097_v23_)[default__.safeIndex(1, (d_2097_v23_).length(0))])) <= (default__.safeDivisionInt(len(d_2117_cf21_), p2))
                rhs358_ = p0
                r2 = rhs354_
                r3 = rhs355_
                d_2097_v23_ = rhs356_
                r3 = rhs357_
                r3 = rhs358_
                index402_ = default__.safeIndex(1, (d_2097_v23_).length(0))
                (d_2097_v23_)[index402_] = ((d_2097_v23_)[default__.safeIndex(1, (d_2097_v23_).length(0))] if p0 else (d_2097_v23_)[default__.safeIndex(1, (d_2097_v23_).length(0))])
                d_2118_v34_: _dafny.MultiSet
                d_2118_v34_ = _dafny.MultiSet([p2, p2])
                d_2119_v35_: _dafny.MultiSet
                d_2119_v35_ = d_2118_v34_
                r3 = (d_2118_v34_).ispropersubset(((d_2119_v35_)) - (d_2118_v34_))
                d_2120_v36_: _dafny.Array
                def lambda92_(d_2121_p0_):
                    def lambda93_(d_2122_i1_):
                        return d_2121_p0_

                    return lambda93_

                init49_ = lambda92_(p0)
                nw393_ = _dafny.Array(None, 29)
                for i0_49_ in range(nw393_.length(0)):
                    nw393_[i0_49_] = init49_(i0_49_)
                d_2120_v36_ = nw393_
                index403_ = default__.safeIndex(33, (d_2120_v36_).length(0))
                (d_2120_v36_)[index403_] = p0
                index404_ = default__.safeIndex(33, (d_2120_v36_).length(0))
                (d_2120_v36_)[index404_] = p3
        elif True:
            if p3:
                r3 = not(p3)
                d_2123_v37_: _dafny.Seq
                d_2123_v37_ = _dafny.SeqWithoutIsStrInference([p2])
                d_2124_v38_: _dafny.Array
                nw394_ = _dafny.Array(None, 7)
                nw394_[int(0)] = p2
                nw394_[int(1)] = p2
                nw394_[int(2)] = (0) - (p2)
                nw394_[int(3)] = default__.safeDivisionInt((self).fm1(p3, globalState), p2)
                nw394_[int(4)] = p2
                nw394_[int(5)] = p2
                nw394_[int(6)] = len((d_2123_v37_) + (d_2123_v37_))
                d_2124_v38_ = nw394_
                index405_ = default__.safeIndex(915, (d_2124_v38_).length(0))
                (d_2124_v38_)[index405_] = len((d_2123_v37_) + (d_2123_v37_))
                index406_ = default__.safeIndex(915, (d_2124_v38_).length(0))
                (d_2124_v38_)[index406_] = default__.safeModuloInt(p2, len(default__.fm41(d_2123_v37_, p0, not(default__.fm5(globalState)), p0, globalState)))
                d_2125_v39_: _dafny.Map
                d_2125_v39_ = _dafny.Map({(d_2124_v38_)[default__.safeIndex(915, (d_2124_v38_).length(0))]: _dafny.SeqWithoutIsStrInference([p2])})
                d_2125_v39_ = (d_2125_v39_).set((0) - ((d_2124_v38_)[default__.safeIndex(915, (d_2124_v38_).length(0))]), d_2123_v37_)
                d_2126_v40_: C1
                nw395_ = C1()
                nw395_.ctor__(p1)
                d_2126_v40_ = nw395_
                d_2127_v41_: T1
                nw396_ = C4()
                nw396_.ctor__((795 if default__.fm5(globalState) else (D8_DC17(p3, d_2126_v40_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycbfarff"))).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycbfarff")))), _dafny.CodePoint('a'))))).cf32))
                d_2127_v41_ = nw396_
                d_2127_v41_ = d_2127_v41_
                d_2128_v42_: _dafny.Array
                nw397_ = _dafny.Array(None, 1)
                nw397_[int(0)] = True
                d_2128_v42_ = nw397_
                d_2129_v43_: _dafny.Set
                d_2129_v43_ = _dafny.Set({d_2128_v42_})
                d_2130_v44_: C0
                nw398_ = C0()
                nw398_.ctor__((d_2129_v43_).ispropersubset(d_2129_v43_), p0)
                d_2130_v44_ = nw398_
            elif True:
                d_2131_v45_: _dafny.Map
                d_2131_v45_ = _dafny.Map({p2: (len(_dafny.Map({331: (0) - (p2)}))) - (p2)})
                d_2132_v46_: _dafny.Seq
                d_2132_v46_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2133_v47_: T2
                nw399_ = C3()
                nw399_.ctor__(p2)
                d_2133_v47_ = nw399_
                d_2134_v48_: _dafny.MultiSet
                d_2134_v48_ = _dafny.MultiSet([p0])
                d_2135_v49_: D5
                d_2135_v49_ = D5_DC9(p0, len(_dafny.Set({d_2133_v47_, d_2133_v47_})), d_2134_v48_)
                d_2136_v50_: _dafny.Seq
                d_2136_v50_ = _dafny.SeqWithoutIsStrInference([(d_2132_v46_)[default__.safeIndex(p2, len(d_2132_v46_))], p0, (d_2135_v49_).cf14, p3])
                r1 = (0) - (((d_2131_v45_)[(p2) * (p2)] if ((p2) * (p2)) in (d_2131_v45_) else len(d_2132_v46_)))
                d_2137_v51_: _dafny.Set
                d_2137_v51_ = _dafny.Set({p2})
                d_2138_v52_: _dafny.MultiSet
                d_2138_v52_ = _dafny.MultiSet([p2, (0) - (p2)])
                d_2139_v53_: _dafny.Array
                nw400_ = _dafny.Array(None, 25)
                nw400_[int(0)] = p2
                nw400_[int(1)] = ((d_2131_v45_)[p2] if (p2) in (d_2131_v45_) else len(d_2137_v51_))
                nw400_[int(2)] = p2
                nw400_[int(3)] = p2
                nw400_[int(4)] = d_2133_v47_.f2
                nw400_[int(5)] = ((d_2138_v52_)[len(default__.fm15(globalState))] if (len(default__.fm15(globalState))) in (d_2138_v52_) else p2)
                nw400_[int(6)] = (d_2134_v48_).cardinality
                nw400_[int(7)] = p2
                nw400_[int(8)] = p2
                nw400_[int(9)] = d_2133_v47_.f2
                nw400_[int(10)] = (0) - (p2)
                nw400_[int(11)] = d_2133_v47_.f2
                nw400_[int(12)] = (self).fm1(p0, globalState)
                nw400_[int(13)] = len(p1)
                nw400_[int(14)] = d_2133_v47_.f2
                nw400_[int(15)] = p2
                nw400_[int(16)] = d_2133_v47_.f2
                nw400_[int(17)] = 267
                nw400_[int(18)] = d_2133_v47_.f2
                nw400_[int(19)] = d_2133_v47_.f2
                nw400_[int(20)] = p2
                nw400_[int(21)] = (0) - (p2)
                nw400_[int(22)] = d_2133_v47_.f2
                nw400_[int(23)] = len(p1)
                nw400_[int(24)] = (d_2133_v47_).fm1(p0, globalState)
                d_2139_v53_ = nw400_
                r1 = len(_dafny.Map({not(p0): d_2139_v53_}))
                d_2140_v54_: _dafny.Seq
                d_2140_v54_ = _dafny.SeqWithoutIsStrInference([default__.fm16(globalState)])
                d_2141_v55_: _dafny.Array
                nw401_ = _dafny.Array(None, 8)
                nw401_[int(0)] = (_dafny.SeqWithoutIsStrInference([p0])) + ((d_2132_v46_).set(default__.safeIndex(d_2133_v47_.f2, len(d_2132_v46_)), p0))
                nw401_[int(1)] = (d_2132_v46_) + (d_2136_v50_)
                nw401_[int(2)] = (d_2140_v54_)[default__.safeIndex((0) - (p2), len(d_2140_v54_))]
                nw401_[int(3)] = d_2136_v50_
                nw401_[int(4)] = d_2132_v46_
                nw401_[int(5)] = d_2136_v50_
                nw401_[int(6)] = d_2132_v46_
                nw401_[int(7)] = d_2136_v50_
                d_2141_v55_ = nw401_
                index407_ = default__.safeIndex(866, (d_2141_v55_).length(0))
                (d_2141_v55_)[index407_] = default__.fm16(globalState)
                index408_ = default__.safeIndex(866, (d_2141_v55_).length(0))
                (d_2141_v55_)[index408_] = d_2132_v46_
                r3 = p0
                d_2142_v56_: _dafny.Seq
                d_2142_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndgceyxxq"))
                index409_ = default__.safeIndex(244, (d_2139_v53_).length(0))
                (d_2139_v53_)[index409_] = (self).fm1(p0, globalState)
                index410_ = default__.safeIndex(244, (d_2139_v53_).length(0))
                rhs359_ = (p1) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxrkeyvv")))
                rhs360_ = p3
                rhs361_ = d_2133_v47_.f2
                rhs362_ = d_2139_v53_
                rhs363_ = (self).fm1((d_2133_v47_).fm8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtikft")), p3, globalState), globalState)
                lhs180_ = d_2139_v53_
                lhs181_ = default__.safeIndex(244, (d_2139_v53_).length(0))
                lhs182_ = d_2133_v47_
                d_2142_v56_ = rhs359_
                r3 = rhs360_
                lhs180_[lhs181_] = rhs361_
                d_2139_v53_ = rhs362_
                lhs182_.f2 = rhs363_
            d_2143_v57_: _dafny.Array
            nw402_ = _dafny.Array(D4.default()(), 9)
            d_2143_v57_ = nw402_
            d_2144_v58_: _dafny.Map
            d_2144_v58_ = _dafny.Map({p2: True})
            d_2145_v59_: _dafny.Set
            d_2145_v59_ = _dafny.Set({len(d_2144_v58_), p2})
            d_2146_v60_: _dafny.Seq
            d_2146_v60_ = _dafny.SeqWithoutIsStrInference([410])
            d_2147_v61_: _dafny.Array
            nw403_ = _dafny.Array(None, 20)
            nw403_[int(0)] = p2
            nw403_[int(1)] = p2
            nw403_[int(2)] = p2
            nw403_[int(3)] = len(d_2145_v59_)
            nw403_[int(4)] = (0) - (p2)
            nw403_[int(5)] = p2
            nw403_[int(6)] = p2
            nw403_[int(7)] = p2
            nw403_[int(8)] = p2
            nw403_[int(9)] = len(d_2146_v60_)
            nw403_[int(10)] = (0) - (len(p1))
            nw403_[int(11)] = p2
            nw403_[int(12)] = 934
            nw403_[int(13)] = p2
            nw403_[int(14)] = len(d_2146_v60_)
            nw403_[int(15)] = p2
            nw403_[int(16)] = p2
            nw403_[int(17)] = -472
            nw403_[int(18)] = p2
            nw403_[int(19)] = -928
            d_2147_v61_ = nw403_
            d_2148_v62_: D4
            d_2148_v62_ = D4_DC6(p2, d_2147_v61_, 994)
            index411_ = default__.safeIndex(259, (d_2143_v57_).length(0))
            (d_2143_v57_)[index411_] = d_2148_v62_
            d_2149_v63_: _dafny.Array
            nw404_ = _dafny.Array(_dafny.Array(None, 0), 26)
            d_2149_v63_ = nw404_
            d_2150_v64_: _dafny.Array
            nw405_ = _dafny.Array(False, 4)
            d_2150_v64_ = nw405_
            index412_ = default__.safeIndex(439, (d_2149_v63_).length(0))
            (d_2149_v63_)[index412_] = d_2150_v64_
            d_2151_v65_: _dafny.Map
            d_2151_v65_ = _dafny.Map({p1: _dafny.CodePoint('n')})
            d_2152_v66_: _dafny.Map
            d_2152_v66_ = _dafny.Map({d_2151_v65_: (p0 if True else p3)})
            d_2153_v68_: _dafny.Map
            d_2153_v68_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fao"))})
            d_2154_v69_: _dafny.Map
            d_2154_v69_ = _dafny.Map({p1: ((d_2153_v68_)[p0] if (p0) in (d_2153_v68_) else p1)})
            index413_ = default__.safeIndex(259, (d_2143_v57_).length(0))
            index414_ = default__.safeIndex(439, (d_2149_v63_).length(0))
            def iife152_():
                coll66_ = _dafny.Map()
                compr_66_: _dafny.Seq
                for compr_66_ in (d_2154_v69_).keys.Elements:
                    d_2155_v67_: _dafny.Seq = compr_66_
                    if (d_2155_v67_) in (d_2154_v69_):
                        coll66_[d_2155_v67_] = _dafny.CodePoint('o')
                return _dafny.Map(coll66_)
            def iife153_():
                coll67_ = _dafny.Map()
                compr_67_: _dafny.Seq
                for compr_67_ in (d_2154_v69_).keys.Elements:
                    d_2156_v67_: _dafny.Seq = compr_67_
                    if (d_2156_v67_) in (d_2154_v69_):
                        coll67_[d_2156_v67_] = _dafny.CodePoint('o')
                return _dafny.Map(coll67_)
            rhs364_ = D4_DC6(len(d_2146_v60_), d_2147_v61_, p2)
            rhs365_ = d_2150_v64_
            rhs366_ = ((d_2152_v66_)[(d_2151_v65_) | (iife152_()
            )] if ((d_2151_v65_) | (iife153_()
            )) in (d_2152_v66_) else p3)
            rhs367_ = 206
            lhs183_ = d_2143_v57_
            lhs184_ = default__.safeIndex(259, (d_2143_v57_).length(0))
            lhs185_ = d_2149_v63_
            lhs186_ = default__.safeIndex(439, (d_2149_v63_).length(0))
            lhs183_[lhs184_] = rhs364_
            lhs185_[lhs186_] = rhs365_
            r3 = rhs366_
            r2 = rhs367_
            d_2157_v70_: str
            d_2157_v70_ = _dafny.CodePoint('l')
            d_2157_v70_ = d_2157_v70_
            d_2158_v71_: _dafny.MultiSet
            d_2158_v71_ = _dafny.MultiSet([len(p1)])
            d_2159_v72_: _dafny.Seq
            d_2159_v72_ = _dafny.SeqWithoutIsStrInference([d_2158_v71_])
            d_2160_v73_: _dafny.Map
            d_2160_v73_ = _dafny.Map({(d_2149_v63_)[default__.safeIndex(439, (d_2149_v63_).length(0))]: not((d_2158_v71_).issubset((d_2159_v72_)[default__.safeIndex(p2, len(d_2159_v72_))]))})
            d_2160_v73_ = d_2160_v73_
            r2 = (p2) * (119)
        d_2161_v74_: _dafny.Map
        d_2161_v74_ = _dafny.Map({66: p3})
        d_2162_v75_: _dafny.Map
        d_2162_v75_ = _dafny.Map({p1: d_2161_v74_})
        d_2163_v76_: D5
        d_2163_v76_ = D5_DC10(False, p2, (d_2162_v75_).set(p1, d_2161_v74_), p0)
        hi13_ = (d_2163_v76_).cf18
        for d_2164_i2_ in range(p2, hi13_):
            r3 = p0
            d_2165_v77_: _dafny.Array
            nw406_ = _dafny.Array(False, 8)
            d_2165_v77_ = nw406_
            d_2165_v77_ = d_2165_v77_
            if not(p3):
                d_2166_v78_: _dafny.Set
                d_2166_v78_ = _dafny.Set({d_2164_i2_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2167_i3_ in range(default__.abs(300))])), d_2164_i2_})
                rhs368_ = default__.safeModuloInt((0) - (p2), len(d_2166_v78_))
                rhs369_ = ((d_2164_i2_ if p0 else p2)) + (-664)
                r1 = rhs368_
                r2 = rhs369_
                r2 = d_2164_i2_
                r3 = False
                r2 = p2
                d_2168_v79_: _dafny.MultiSet
                d_2168_v79_ = _dafny.MultiSet([p0, p3])
                d_2169_v80_: _dafny.Map
                d_2169_v80_ = _dafny.Map({(d_2164_i2_) + (((d_2168_v79_)[p0] if (p0) in (d_2168_v79_) else d_2164_i2_)): p2})
                d_2170_v81_: _dafny.Map
                d_2170_v81_ = _dafny.Map({p0: d_2168_v79_})
                d_2169_v80_ = (d_2169_v80_).set((p2 if p3 else d_2164_i2_), ((d_2168_v79_ if default__.fm5(globalState) else ((d_2170_v81_)[p3] if (p3) in (d_2170_v81_) else d_2168_v79_))).cardinality)
            elif True:
                d_2171_v82_: _dafny.Seq
                d_2171_v82_ = _dafny.SeqWithoutIsStrInference([len(p1), d_2164_i2_])
                d_2172_v83_: _dafny.Map
                d_2172_v83_ = _dafny.Map({d_2171_v82_: p0})
                d_2173_v84_: _dafny.Map
                d_2173_v84_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")): (_dafny.SeqWithoutIsStrInference([(0) - (-663), len(d_2172_v83_), (self).fm1(True, globalState)])) != (d_2171_v82_)})
                d_2173_v84_ = (d_2173_v84_).set(p1, p3)
                index415_ = default__.safeIndex(296, (d_2165_v77_).length(0))
                (d_2165_v77_)[index415_] = not(p0)
                index416_ = default__.safeIndex(296, (d_2165_v77_).length(0))
                (d_2165_v77_)[index416_] = not (p3) or (p3)
                d_2174_v85_: C1
                nw407_ = C1()
                nw407_.ctor__((p1) + (p1))
                d_2174_v85_ = nw407_
                d_2175_v86_: _dafny.Array
                nw408_ = _dafny.Array(None, 26)
                nw408_[int(0)] = False
                nw408_[int(1)] = p3
                nw408_[int(2)] = p0
                nw408_[int(3)] = p0
                nw408_[int(4)] = p3
                nw408_[int(5)] = not(default__.fm5(globalState))
                nw408_[int(6)] = p3
                nw408_[int(7)] = (d_2165_v77_)[default__.safeIndex(296, (d_2165_v77_).length(0))]
                nw408_[int(8)] = p3
                nw408_[int(9)] = p0
                nw408_[int(10)] = p3
                nw408_[int(11)] = (d_2165_v77_)[default__.safeIndex(296, (d_2165_v77_).length(0))]
                nw408_[int(12)] = p0
                nw408_[int(13)] = (d_2165_v77_)[default__.safeIndex(296, (d_2165_v77_).length(0))]
                nw408_[int(14)] = True
                nw408_[int(15)] = p3
                nw408_[int(16)] = (d_2165_v77_)[default__.safeIndex(296, (d_2165_v77_).length(0))]
                nw408_[int(17)] = (d_2165_v77_)[default__.safeIndex(296, (d_2165_v77_).length(0))]
                nw408_[int(18)] = p0
                nw408_[int(19)] = (d_2165_v77_)[default__.safeIndex(296, (d_2165_v77_).length(0))]
                nw408_[int(20)] = (d_2165_v77_)[default__.safeIndex(296, (d_2165_v77_).length(0))]
                nw408_[int(21)] = p3
                nw408_[int(22)] = p3
                nw408_[int(23)] = (d_2165_v77_)[default__.safeIndex(296, (d_2165_v77_).length(0))]
                nw408_[int(24)] = p0
                nw408_[int(25)] = p3
                d_2175_v86_ = nw408_
                d_2176_v87_: _dafny.Map
                d_2176_v87_ = _dafny.Map({d_2175_v86_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pgdsyug"))) + ((d_2174_v85_).f3)})
                d_2177_v88_: str
                d_2177_v88_ = _dafny.CodePoint('f')
                index417_ = default__.safeIndex(296, (d_2165_v77_).length(0))
                rhs370_ = len(_dafny.SeqWithoutIsStrInference([(d_2177_v88_ if False else d_2177_v88_)]))
                rhs371_ = d_2164_i2_
                rhs372_ = p2
                rhs373_ = ((d_2176_v87_) | (d_2176_v87_) if not((d_2165_v77_)[default__.safeIndex(296, (d_2165_v77_).length(0))]) else d_2176_v87_)
                rhs374_ = p3
                lhs187_ = d_2165_v77_
                lhs188_ = default__.safeIndex(296, (d_2165_v77_).length(0))
                r2 = rhs370_
                r2 = rhs371_
                r2 = rhs372_
                d_2176_v87_ = rhs373_
                lhs187_[lhs188_] = rhs374_
                index418_ = default__.safeIndex(296, (d_2165_v77_).length(0))
                (d_2165_v77_)[index418_] = (d_2165_v77_)[default__.safeIndex(296, (d_2165_v77_).length(0))]
            r1 = (self).fm1(p3, globalState)
        d_2178_v89_: D15
        d_2178_v89_ = D15_DC30(True)
        r1 = (self).fm1((d_2178_v89_).cf47, globalState)
        d_2179_v90_: _dafny.Array
        nw409_ = _dafny.Array(int(0), 8)
        d_2179_v90_ = nw409_
        index419_ = default__.safeIndex(114, (d_2179_v90_).length(0))
        (d_2179_v90_)[index419_] = default__.safeModuloInt(138, 795)
        d_2180_v91_: _dafny.Seq
        d_2180_v91_ = _dafny.SeqWithoutIsStrInference([len(p1)])
        d_2181_v92_: _dafny.Map
        d_2181_v92_ = _dafny.Map({d_2180_v91_: p3})
        index420_ = default__.safeIndex(114, (d_2179_v90_).length(0))
        (d_2179_v90_)[index420_] = (len(d_2181_v92_)) - (p2)
        d_2182_v93_: _dafny.Array
        nw410_ = _dafny.Array(_dafny.Array(None, 0), 13)
        d_2182_v93_ = nw410_
        d_2183_v95_: _dafny.Seq
        d_2183_v95_ = _dafny.SeqWithoutIsStrInference([p1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdflwo"))])
        d_2184_v96_: D15
        def iife154_():
            coll68_ = _dafny.Map()
            compr_68_: _dafny.Seq
            for compr_68_ in ((d_2183_v95_).set(default__.safeIndex(p2, len(d_2183_v95_)), (d_2183_v95_)[default__.safeIndex((d_2179_v90_)[default__.safeIndex(114, (d_2179_v90_).length(0))], len(d_2183_v95_))])).Elements:
                d_2185_v94_: _dafny.Seq = compr_68_
                if (d_2185_v94_) in ((d_2183_v95_).set(default__.safeIndex(p2, len(d_2183_v95_)), (d_2183_v95_)[default__.safeIndex((d_2179_v90_)[default__.safeIndex(114, (d_2179_v90_).length(0))], len(d_2183_v95_))])):
                    coll68_[d_2185_v94_] = p1
            return _dafny.Map(coll68_)
        d_2184_v96_ = D15_DC29(iife154_()
)
        pat_let_tv47_ = d_2179_v90_
        pat_let_tv48_ = d_2179_v90_
        pat_let_tv49_ = d_2179_v90_
        pat_let_tv50_ = d_2179_v90_
        pat_let_tv51_ = d_2179_v90_
        pat_let_tv52_ = d_2179_v90_
        index421_ = default__.safeIndex(114, (d_2179_v90_).length(0))
        def lambda94_(source32_):
            if source32_.is_DC30:
                d_2186___mcc_h5_ = source32_.cf47
                d_2187_cf47_ = d_2186___mcc_h5_
                return default__.safeDivisionInt((pat_let_tv48_)[default__.safeIndex(114, (pat_let_tv47_).length(0))], (pat_let_tv50_)[default__.safeIndex(114, (pat_let_tv49_).length(0))])
            elif True:
                d_2188___mcc_h6_ = source32_.cf46
                d_2189_cf46_ = d_2188___mcc_h6_
                return (pat_let_tv52_)[default__.safeIndex(114, (pat_let_tv51_).length(0))]

        rhs375_ = d_2182_v93_
        rhs376_ = default__.safeDivisionInt((d_2179_v90_)[default__.safeIndex(114, (d_2179_v90_).length(0))], p2)
        rhs377_ = lambda94_(d_2184_v96_)
        rhs378_ = len(p1)
        lhs189_ = d_2179_v90_
        lhs190_ = default__.safeIndex(114, (d_2179_v90_).length(0))
        d_2182_v93_ = rhs375_
        r1 = rhs376_
        lhs189_[lhs190_] = rhs377_
        r1 = rhs378_
        r0 = ((d_2183_v95_ if not(p3) else _dafny.SeqWithoutIsStrInference([p1, p1, p1, p1, p1]))) + ((d_2183_v95_) + (d_2183_v95_))
        d_2190_v97_: _dafny.MultiSet
        d_2190_v97_ = _dafny.MultiSet([p0])
        r1 = (len(p1) if (d_2190_v97_).ispropersubset(d_2190_v97_) else p2)
        r2 = (d_2180_v91_)[default__.safeIndex((d_2179_v90_)[default__.safeIndex(114, (d_2179_v90_).length(0))], len(d_2180_v91_))]
        d_2191_v98_: str
        d_2191_v98_ = _dafny.CodePoint('b')
        d_2192_v99_: D7
        d_2192_v99_ = D7_DC15((d_2179_v90_)[default__.safeIndex(114, (d_2179_v90_).length(0))], d_2191_v98_)
        pat_let_tv53_ = p3
        pat_let_tv54_ = d_2179_v90_
        pat_let_tv55_ = d_2179_v90_
        def lambda95_(source33_):
            if source33_.is_DC15:
                d_2193___mcc_h7_ = source33_.cf27
                d_2194___mcc_h8_ = source33_.cf28
                d_2195_cf28_ = d_2194___mcc_h8_
                d_2196_cf27_ = d_2193___mcc_h7_
                return pat_let_tv53_
            elif True:
                d_2197___mcc_h9_ = source33_.cf26
                d_2198_cf26_ = d_2197___mcc_h9_
                return False

        def iife155_(_pat_let43_0):
            def iife156_(d_2199_dt__update__tmp_h0_):
                def iife157_(_pat_let44_0):
                    def iife158_(d_2200_dt__update_hcf27_h0_):
                        return D7_DC15(d_2200_dt__update_hcf27_h0_, (d_2199_dt__update__tmp_h0_).cf28)
                    return iife158_(_pat_let44_0)
                return iife157_((pat_let_tv55_)[default__.safeIndex(114, (pat_let_tv54_).length(0))])
            return iife156_(_pat_let43_0)
        r3 = lambda95_(iife155_(d_2192_v99_))
        return r0, r1, r2, r3


class C7(T1, T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self):
        pass
        pass

    def fm0(self, globalState):
        def iife159_():
            coll69_ = _dafny.Map()
            compr_69_: int
            for compr_69_ in _dafny.IntegerRange(-18, 784):
                d_2203_v0_: int = compr_69_
                if ((-18) <= (d_2203_v0_)) and ((d_2203_v0_) < (784)):
                    coll69_[default__.safeDivisionInt(d_2203_v0_, len(_dafny.SeqWithoutIsStrInference([False])))] = False
            return _dafny.Map(coll69_)
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality for d_2201_i0_ in range(default__.abs(462))]): 711})) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cpdthucj"))) for d_2202_i1_ in range(default__.abs(638))]): 618})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([280, len(iife159_()
        ), 438, -11, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2204_i2_ in range(default__.abs(813))]))]): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_2205_i3_ in range(default__.abs(-503))]))})))

    def fm1(self, p0, globalState):
        return default__.safeDivisionInt(len((_dafny.Set({not(not(False)), False, False})).intersection(_dafny.Set({True, (D0_DC0(True)).cf0}))), 312)

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_2206_v0_: _dafny.Seq
        d_2206_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsn"))
        d_2206_v0_ = ((d_2206_v0_) + (default__.fm2(globalState))) + (d_2206_v0_)
        d_2207_v1_: int
        d_2207_v1_ = 583
        r1 = (0) - (d_2207_v1_)
        d_2208_v2_: bool
        d_2208_v2_ = True
        d_2208_v2_ = d_2208_v2_
        d_2208_v2_ = d_2208_v2_
        r0 = d_2207_v1_
        d_2209_v3_: _dafny.Map
        d_2209_v3_ = _dafny.Map({d_2208_v2_: p3})
        d_2210_v4_: _dafny.Seq
        d_2210_v4_ = _dafny.SeqWithoutIsStrInference([d_2208_v2_])
        d_2211_v5_: _dafny.Seq
        d_2211_v5_ = _dafny.SeqWithoutIsStrInference([p2, ((d_2209_v3_)[True] if (True) in (d_2209_v3_) else (d_2210_v4_)[default__.safeIndex((self).fm1(p3, globalState), len(d_2210_v4_))]), p2])
        d_2212_v6_: _dafny.Array
        nw411_ = _dafny.Array(None, 22)
        nw411_[int(0)] = (p2) == (d_2208_v2_)
        nw411_[int(1)] = p3
        nw411_[int(2)] = d_2208_v2_
        nw411_[int(3)] = True
        nw411_[int(4)] = d_2208_v2_
        nw411_[int(5)] = (d_2210_v4_)[default__.safeIndex(d_2207_v1_, len(d_2210_v4_))]
        nw411_[int(6)] = d_2208_v2_
        nw411_[int(7)] = not(d_2208_v2_)
        nw411_[int(8)] = p2
        nw411_[int(9)] = not(d_2208_v2_)
        nw411_[int(10)] = True
        nw411_[int(11)] = not(p3)
        nw411_[int(12)] = (d_2208_v2_) and (p2)
        nw411_[int(13)] = d_2208_v2_
        nw411_[int(14)] = p2
        nw411_[int(15)] = default__.fm3(True, d_2207_v1_, d_2207_v1_, globalState)
        nw411_[int(16)] = (d_2210_v4_) != (d_2210_v4_)
        nw411_[int(17)] = (d_2207_v1_) < (d_2207_v1_)
        nw411_[int(18)] = not((p2 if ((d_2209_v3_)[p3] if (p3) in (d_2209_v3_) else ((d_2209_v3_)[p2] if (p2) in (d_2209_v3_) else p2)) else False))
        nw411_[int(19)] = d_2208_v2_
        nw411_[int(20)] = p2
        nw411_[int(21)] = d_2208_v2_
        d_2212_v6_ = nw411_
        d_2213_v7_: _dafny.Seq
        d_2213_v7_ = _dafny.SeqWithoutIsStrInference([d_2207_v1_, (p0).cardinality, d_2207_v1_])
        d_2214_v8_: _dafny.Set
        d_2214_v8_ = _dafny.Set({d_2207_v1_, 799})
        d_2215_v9_: _dafny.Set
        d_2215_v9_ = _dafny.Set({d_2214_v8_})
        d_2216_v10_: _dafny.Seq
        d_2216_v10_ = _dafny.SeqWithoutIsStrInference([d_2214_v8_, d_2214_v8_])
        d_2217_v12_: _dafny.Seq
        def iife160_():
            coll70_ = _dafny.Set()
            compr_70_: _dafny.Set
            for compr_70_ in (d_2216_v10_).Elements:
                d_2218_v11_: _dafny.Set = compr_70_
                if (d_2218_v11_) in (d_2216_v10_):
                    coll70_ = coll70_.union(_dafny.Set([d_2218_v11_]))
            return _dafny.Set(coll70_)
        d_2217_v12_ = _dafny.SeqWithoutIsStrInference([d_2215_v9_, _dafny.Set({d_2214_v8_}), iife160_()
        , d_2215_v9_, d_2215_v9_])
        d_2219_v13_: str
        d_2219_v13_ = _dafny.CodePoint('t')
        d_2220_v14_: _dafny.Map
        d_2220_v14_ = _dafny.Map({d_2219_v13_: p3})
        d_2221_v15_: D0
        d_2221_v15_ = D0_DC1(d_2208_v2_, p3, d_2219_v13_, (0) - (d_2207_v1_))
        d_2222_v16_: _dafny.Map
        d_2222_v16_ = _dafny.Map({d_2221_v15_: d_2208_v2_})
        nw412_ = _dafny.Array(None, 24)
        nw412_[int(0)] = (p2) or (p2)
        nw412_[int(1)] = not (d_2208_v2_) or (d_2208_v2_)
        nw412_[int(2)] = p2
        nw412_[int(3)] = d_2208_v2_
        nw412_[int(4)] = d_2208_v2_
        nw412_[int(5)] = p3
        nw412_[int(6)] = ((d_2209_v3_)[True] if (True) in (d_2209_v3_) else p2)
        nw412_[int(7)] = d_2208_v2_
        nw412_[int(8)] = True
        nw412_[int(9)] = p3
        nw412_[int(10)] = False
        nw412_[int(11)] = p2
        nw412_[int(12)] = True
        nw412_[int(13)] = d_2208_v2_
        nw412_[int(14)] = False
        nw412_[int(15)] = (d_2213_v7_) != (d_2213_v7_)
        nw412_[int(16)] = d_2208_v2_
        nw412_[int(17)] = d_2208_v2_
        nw412_[int(18)] = ((d_2217_v12_)[default__.safeIndex(len(d_2220_v14_), len(d_2217_v12_))]).ispropersubset(_dafny.Set({d_2214_v8_}))
        nw412_[int(19)] = (D0_DC1(d_2208_v2_, d_2208_v2_, d_2219_v13_, d_2207_v1_)) not in (d_2222_v16_)
        nw412_[int(20)] = (d_2214_v8_).issubset(d_2214_v8_)
        nw412_[int(21)] = p2
        nw412_[int(22)] = p3
        nw412_[int(23)] = p3
        d_2212_v6_ = nw412_
        r0 = (d_2207_v1_) * (d_2207_v1_)
        d_2223_v17_: _dafny.Map
        d_2223_v17_ = _dafny.Map({d_2212_v6_: _dafny.CodePoint('r')})
        d_2224_v18_: _dafny.MultiSet
        d_2224_v18_ = _dafny.MultiSet([d_2219_v13_, d_2219_v13_, ((d_2223_v17_)[d_2212_v6_] if (d_2212_v6_) in (d_2223_v17_) else d_2219_v13_)])
        r1 = default__.safeModuloInt((0) - ((d_2224_v18_).cardinality), default__.safeModuloInt((self).fm1(d_2208_v2_, globalState), d_2207_v1_))
        return r0, r1

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        d_2225_v1_: _dafny.MultiSet
        d_2225_v1_ = _dafny.MultiSet([p0])
        def iife161_():
            coll71_ = _dafny.Map()
            compr_71_: int
            for compr_71_ in (d_2225_v1_).Elements:
                d_2226_v0_: int = compr_71_
                if (d_2226_v0_) in (d_2225_v1_):
                    coll71_[(d_2226_v0_) * (((d_2225_v1_).set(p0, default__.abs(p0))).cardinality)] = p0
            return _dafny.Map(coll71_)
        if (p0) in (iife161_()
        ):
            d_2227_v2_: str
            d_2227_v2_ = _dafny.CodePoint('x')
            d_2228_v3_: _dafny.Set
            d_2228_v3_ = _dafny.Set({p3, p2, p2})
            d_2229_v4_: D0
            d_2229_v4_ = D0_DC1(p2, p3, d_2227_v2_, len(d_2228_v3_))
            d_2230_v5_: _dafny.MultiSet
            d_2230_v5_ = _dafny.MultiSet([d_2229_v4_, d_2229_v4_])
            d_2231_v6_: _dafny.Map
            d_2231_v6_ = _dafny.Map({p3: (d_2230_v5_).set(d_2229_v4_, default__.abs(p0))})
            d_2232_v7_: _dafny.Seq
            d_2232_v7_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_2233_v8_: _dafny.Seq
            d_2233_v8_ = _dafny.SeqWithoutIsStrInference([default__.fm4(p0, len(d_2232_v7_), p2, p2, globalState)])
            d_2231_v6_ = (d_2231_v6_).set(p2, _dafny.MultiSet(d_2233_v8_))
            d_2234_v9_: _dafny.Array
            out49_: _dafny.Array
            out49_ = default__.m0((not(p3)) == (p2), p3, globalState)
            d_2234_v9_ = out49_
            r0 = p3
            d_2235_v10_: C0
            nw413_ = C0()
            nw413_.ctor__(p3, p3)
            d_2235_v10_ = nw413_
            d_2236_v11_: _dafny.Array
            nw414_ = _dafny.Array(False, 2)
            d_2236_v11_ = nw414_
            index422_ = default__.safeIndex(956, (d_2236_v11_).length(0))
            (d_2236_v11_)[index422_] = p2
            index423_ = default__.safeIndex(956, (d_2236_v11_).length(0))
            (d_2236_v11_)[index423_] = p2
        elif True:
            d_2237_v12_: _dafny.Seq
            d_2237_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "acq"))
            d_2238_v13_: _dafny.Set
            d_2238_v13_ = _dafny.Set({p3})
            d_2239_v14_: _dafny.Map
            d_2239_v14_ = _dafny.Map({(d_2237_v12_ if p3 else d_2237_v12_): (_dafny.Set({p2})).intersection(d_2238_v13_)})
            d_2240_v15_: str
            d_2240_v15_ = _dafny.CodePoint('c')
            d_2239_v14_ = (d_2239_v14_).set((d_2237_v12_).set(default__.safeIndex((self).fm1(p2, globalState), len(d_2237_v12_)), d_2240_v15_), d_2238_v13_)
            d_2241_v16_: _dafny.Seq
            d_2241_v16_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2242_v17_: _dafny.Map
            d_2242_v17_ = _dafny.Map({d_2225_v1_: p3})
            d_2243_v18_: _dafny.Set
            d_2243_v18_ = _dafny.Set({(self).fm1(p3, globalState), p0, (0) - (len(d_2241_v16_)), len(d_2242_v17_)})
            r2 = default__.safeModuloInt(p0, len((d_2243_v18_) - (d_2243_v18_)))
            d_2244_v19_: _dafny.Map
            d_2244_v19_ = _dafny.Map({p2: p2})
            d_2245_v20_: D6
            d_2245_v20_ = D6_DC13(p2, p0)
            d_2246_v21_: _dafny.Array
            nw415_ = _dafny.Array(None, 20)
            nw415_[int(0)] = p3
            nw415_[int(1)] = p3
            nw415_[int(2)] = ((d_2244_v19_)[not(p2)] if (not(p2)) in (d_2244_v19_) else p3)
            nw415_[int(3)] = p2
            nw415_[int(4)] = default__.fm5(globalState)
            nw415_[int(5)] = ((d_2244_v19_)[True] if (True) in (d_2244_v19_) else p3)
            nw415_[int(6)] = p2
            nw415_[int(7)] = p2
            nw415_[int(8)] = (p0) != (p0)
            nw415_[int(9)] = p2
            nw415_[int(10)] = not(p2)
            nw415_[int(11)] = p3
            nw415_[int(12)] = False
            nw415_[int(13)] = p3
            nw415_[int(14)] = (p3 if p2 else not(p2))
            nw415_[int(15)] = p3
            nw415_[int(16)] = not(p2)
            nw415_[int(17)] = default__.fm3(p3, p0, p0, globalState)
            nw415_[int(18)] = p3
            nw415_[int(19)] = (d_2245_v20_).cf24
            d_2246_v21_ = nw415_
            d_2247_v22_: _dafny.Map
            d_2247_v22_ = _dafny.Map({p3: (d_2237_v12_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), d_2240_v15_, d_2240_v15_]))})
            rhs379_ = d_2246_v21_
            rhs380_ = (0) - (len(d_2247_v22_))
            d_2246_v21_ = rhs379_
            r2 = rhs380_
            d_2248_v23_: C1
            nw416_ = C1()
            nw416_.ctor__(d_2237_v12_)
            d_2248_v23_ = nw416_
            d_2249_v24_: _dafny.MultiSet
            d_2249_v24_ = _dafny.MultiSet([d_2248_v23_, d_2248_v23_])
            d_2250_v25_: _dafny.Map
            d_2250_v25_ = _dafny.Map({_dafny.Set({d_2249_v24_, d_2249_v24_, d_2249_v24_}): (d_2241_v16_) + (d_2241_v16_)})
            d_2250_v25_ = (d_2250_v25_).set(_dafny.Set({(d_2249_v24_).set(d_2248_v23_, default__.abs(p0))}), d_2241_v16_)
            if p3:
                r2 = p0
                d_2251_v26_: T1
                nw417_ = C4()
                nw417_.ctor__(p0)
                d_2251_v26_ = nw417_
                d_2252_v27_: _dafny.Set
                d_2252_v27_ = _dafny.Set({d_2251_v26_})
                d_2253_v28_: _dafny.Set
                d_2253_v28_ = d_2252_v27_
                d_2254_v29_: _dafny.Seq
                d_2254_v29_ = _dafny.SeqWithoutIsStrInference([d_2252_v27_, (_dafny.Set({d_2251_v26_})) | (d_2252_v27_), (d_2253_v28_)])
                d_2254_v29_ = (_dafny.SeqWithoutIsStrInference([_dafny.Set({d_2251_v26_, d_2251_v26_, d_2251_v26_, d_2251_v26_}), d_2252_v27_])) + (d_2254_v29_)
                r0 = ((d_2244_v19_)[default__.fm5(globalState)] if (default__.fm5(globalState)) in (d_2244_v19_) else True)
                d_2255_v30_: _dafny.Map
                d_2255_v30_ = _dafny.Map({d_2246_v21_: d_2237_v12_})
                d_2255_v30_ = d_2255_v30_
                d_2256_v31_: D14
                d_2256_v31_ = D14_DC28()
                d_2257_v32_: _dafny.Seq
                d_2257_v32_ = _dafny.SeqWithoutIsStrInference([d_2256_v31_, d_2256_v31_])
                index424_ = default__.safeIndex(665, (d_2246_v21_).length(0))
                (d_2246_v21_)[index424_] = p3
                d_2258_v33_: _dafny.Array
                nw418_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_2258_v33_ = nw418_
                index425_ = default__.safeIndex(877, (d_2258_v33_).length(0))
                (d_2258_v33_)[index425_] = d_2246_v21_
                d_2259_v34_: _dafny.Seq
                d_2259_v34_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_2256_v31_ for d_2260_i0_ in range(default__.abs(321))])) + (d_2257_v32_)])
                d_2261_v35_: _dafny.Map
                d_2261_v35_ = _dafny.Map({p3: p0})
                d_2262_v36_: _dafny.Map
                d_2262_v36_ = _dafny.Map({(p3) in (d_2261_v35_): d_2246_v21_})
                index426_ = default__.safeIndex(665, (d_2246_v21_).length(0))
                index427_ = default__.safeIndex(877, (d_2258_v33_).length(0))
                rhs381_ = (d_2259_v34_)[default__.safeIndex(p0, len(d_2259_v34_))]
                rhs382_ = p3
                rhs383_ = ((d_2262_v36_)[p2] if (p2) in (d_2262_v36_) else d_2246_v21_)
                rhs384_ = not(p3)
                lhs191_ = d_2246_v21_
                lhs192_ = default__.safeIndex(665, (d_2246_v21_).length(0))
                lhs193_ = d_2258_v33_
                lhs194_ = default__.safeIndex(877, (d_2258_v33_).length(0))
                d_2257_v32_ = rhs381_
                lhs191_[lhs192_] = rhs382_
                lhs193_[lhs194_] = rhs383_
                r0 = rhs384_
            elif True:
                d_2263_v37_: _dafny.Map
                d_2263_v37_ = _dafny.Map({p0: default__.safeModuloInt(p0, p0)})
                d_2263_v37_ = (d_2263_v37_).set(241, p0)
                d_2264_v38_: _dafny.MultiSet
                d_2264_v38_ = _dafny.MultiSet([p3])
                index428_ = default__.safeIndex(290, (d_2246_v21_).length(0))
                (d_2246_v21_)[index428_] = (default__.fm38(p2, globalState)).issubset(d_2264_v38_)
                index429_ = default__.safeIndex(290, (d_2246_v21_).length(0))
                (d_2246_v21_)[index429_] = p2
                d_2265_v39_: C4
                nw419_ = C4()
                nw419_.ctor__(p0)
                d_2265_v39_ = nw419_
                d_2266_v40_: _dafny.Seq
                d_2266_v40_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2240_v15_ for d_2267_i1_ in range(default__.abs(892))]), (d_2248_v23_).f3, d_2237_v12_, (d_2248_v23_).f3, (d_2248_v23_).f3])
                index430_ = default__.safeIndex(290, (d_2246_v21_).length(0))
                index431_ = default__.safeIndex(290, (d_2246_v21_).length(0))
                rhs385_ = d_2240_v15_
                rhs386_ = (d_2265_v39_).fm8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pinmtrsal")), p3, globalState)
                rhs387_ = not((d_2246_v21_)[default__.safeIndex(290, (d_2246_v21_).length(0))])
                rhs388_ = _dafny.SeqWithoutIsStrInference([((D6_DC11((d_2248_v23_).f3)).cf21).set(default__.safeIndex(p0, len((D6_DC11((d_2248_v23_).f3)).cf21)), d_2240_v15_) for d_2268_i2_ in range(default__.abs(525))])
                lhs195_ = d_2246_v21_
                lhs196_ = default__.safeIndex(290, (d_2246_v21_).length(0))
                lhs197_ = d_2246_v21_
                lhs198_ = default__.safeIndex(290, (d_2246_v21_).length(0))
                d_2240_v15_ = rhs385_
                lhs195_[lhs196_] = rhs386_
                lhs197_[lhs198_] = rhs387_
                d_2266_v40_ = rhs388_
                d_2269_v41_: _dafny.Seq
                d_2269_v41_ = _dafny.SeqWithoutIsStrInference([((d_2264_v38_).set(((d_2244_v19_)[p2] if (p2) in (d_2244_v19_) else p2), default__.abs(p0))) | ((d_2264_v38_).set(p3, default__.abs(669)))])
                d_2269_v41_ = d_2269_v41_
        d_2270_v42_: str
        d_2270_v42_ = _dafny.CodePoint('u')
        d_2271_v43_: _dafny.Map
        d_2271_v43_ = _dafny.Map({p2: p2})
        if (d_2270_v42_) not in ((default__.fm9(p3, not(p2), True, (self).fm1(((d_2271_v43_)[p3] if (p3) in (d_2271_v43_) else False), globalState), globalState)).set(default__.safeIndex((self).fm1(p3, globalState), len(default__.fm9(p3, not(p2), True, (self).fm1(((d_2271_v43_)[p3] if (p3) in (d_2271_v43_) else False), globalState), globalState))), d_2270_v42_)):
            d_2272_v44_: _dafny.Seq
            d_2272_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlysn"))
            d_2273_v45_: D6
            d_2273_v45_ = D6_DC11(d_2272_v44_)
            d_2274_v46_: D5
            d_2274_v46_ = D5_DC9(p2, p0, default__.fm38(p3, globalState))
            d_2275_v47_: _dafny.Seq
            d_2275_v47_ = _dafny.SeqWithoutIsStrInference([p3, (d_2272_v44_) != ((d_2273_v45_).cf21), (d_2274_v46_).cf14])
            if (d_2275_v47_)[default__.safeIndex(p0, len(d_2275_v47_))]:
                r2 = ((len(_dafny.SeqWithoutIsStrInference([p3]))) - (p0)) + (default__.safeModuloInt(p0, p0))
                d_2276_v48_: _dafny.Array
                out50_: _dafny.Array
                out50_ = default__.m0(p2, p3, globalState)
                d_2276_v48_ = out50_
                d_2277_v49_: _dafny.Map
                d_2277_v49_ = _dafny.Map({p3: p0})
                d_2278_v50_: _dafny.Seq
                d_2278_v50_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2279_v51_: _dafny.Map
                d_2279_v51_ = _dafny.Map({((d_2277_v49_)[p2] if (p2) in (d_2277_v49_) else p0): ((d_2278_v50_).set(default__.safeIndex(p0, len(d_2278_v50_)), (self).fm1(p2, globalState)) if p3 else d_2278_v50_)})
                d_2279_v51_ = (d_2279_v51_).set(p0, (d_2278_v50_) + (d_2278_v50_))
                r0 = p2
                r0 = not (p3) or (p2)
            elif True:
                r0 = (p3) == (p3)
                d_2280_v52_: C3
                nw420_ = C3()
                nw420_.ctor__(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qiesjmhyn"))) + (d_2272_v44_)))
                d_2280_v52_ = nw420_
                r0 = not(p3)
                d_2281_v53_: _dafny.Map
                d_2281_v53_ = _dafny.Map({(p0) >= (p0): p0})
                d_2281_v53_ = (d_2281_v53_).set(not(p2), (p0 if True else p0))
                r2 = p0
            d_2282_v54_: _dafny.Array
            nw421_ = _dafny.Array(None, 14)
            nw421_[int(0)] = False
            nw421_[int(1)] = p2
            nw421_[int(2)] = p3
            nw421_[int(3)] = p2
            nw421_[int(4)] = p2
            nw421_[int(5)] = default__.fm3(p2, p0, p0, globalState)
            nw421_[int(6)] = False
            nw421_[int(7)] = default__.fm3(p3, p0, 967, globalState)
            nw421_[int(8)] = default__.fm3(False, p0, p0, globalState)
            nw421_[int(9)] = p2
            nw421_[int(10)] = p3
            nw421_[int(11)] = p3
            nw421_[int(12)] = p3
            nw421_[int(13)] = p2
            d_2282_v54_ = nw421_
            d_2283_v55_: _dafny.Set
            d_2283_v55_ = _dafny.Set({d_2282_v54_, d_2282_v54_, d_2282_v54_, d_2282_v54_, d_2282_v54_})
            rhs389_ = not(not(p2))
            rhs390_ = default__.safeDivisionInt((self).fm1(p2, globalState), p0)
            rhs391_ = (_dafny.Set({d_2282_v54_})).issubset(d_2283_v55_)
            r0 = rhs389_
            r2 = rhs390_
            r0 = rhs391_
            index432_ = default__.safeIndex(640, (d_2282_v54_).length(0))
            (d_2282_v54_)[index432_] = p3
            d_2284_v56_: _dafny.Map
            d_2284_v56_ = _dafny.Map({d_2282_v54_: p3})
            index433_ = default__.safeIndex(640, (d_2282_v54_).length(0))
            rhs392_ = p3
            rhs393_ = d_2284_v56_
            rhs394_ = p0
            lhs199_ = d_2282_v54_
            lhs200_ = default__.safeIndex(640, (d_2282_v54_).length(0))
            lhs199_[lhs200_] = rhs392_
            d_2284_v56_ = rhs393_
            r2 = rhs394_
            d_2285_v57_: _dafny.MultiSet
            d_2285_v57_ = _dafny.MultiSet([(d_2282_v54_)[default__.safeIndex(640, (d_2282_v54_).length(0))], not(True), not(((d_2271_v43_)[(d_2282_v54_)[default__.safeIndex(640, (d_2282_v54_).length(0))]] if ((d_2282_v54_)[default__.safeIndex(640, (d_2282_v54_).length(0))]) in (d_2271_v43_) else p2))])
            d_2285_v57_ = (_dafny.MultiSet([p3])).intersection(d_2285_v57_)
            d_2286_v58_: C1
            nw422_ = C1()
            nw422_.ctor__((d_2272_v44_).set(default__.safeIndex(p0, len(d_2272_v44_)), d_2270_v42_))
            d_2286_v58_ = nw422_
            d_2286_v58_ = d_2286_v58_
        elif True:
            r0 = not (p2) or (p2)
            d_2287_v59_: _dafny.Array
            nw423_ = _dafny.Array(_dafny.CodePoint('D'), 21)
            d_2287_v59_ = nw423_
            d_2288_v60_: _dafny.Seq
            d_2288_v60_ = _dafny.SeqWithoutIsStrInference([True])
            index434_ = default__.safeIndex(630, (d_2287_v59_).length(0))
            (d_2287_v59_)[index434_] = default__.fm37(len(default__.fm49(d_2288_v60_, p2, p0, globalState)), globalState)
            d_2289_v61_: _dafny.Map
            d_2289_v61_ = _dafny.Map({d_2270_v42_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnctyt"))})
            d_2290_v62_: _dafny.Map
            d_2290_v62_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_2270_v42_ for d_2291_i3_ in range(default__.abs(564))])): p0})
            d_2292_v63_: T2
            nw424_ = C2()
            nw424_.ctor__((0) - (p0))
            d_2292_v63_ = nw424_
            d_2293_v64_: _dafny.Set
            d_2293_v64_ = _dafny.Set({d_2292_v63_, d_2292_v63_, d_2292_v63_})
            d_2294_v65_: _dafny.Map
            d_2294_v65_ = _dafny.Map({p3: d_2293_v64_})
            d_2295_v66_: _dafny.Seq
            d_2295_v66_ = _dafny.SeqWithoutIsStrInference([p0, 436, (self).fm1(True, globalState), len(d_2294_v65_)])
            d_2296_v67_: _dafny.Seq
            d_2296_v67_ = _dafny.SeqWithoutIsStrInference([len(d_2290_v62_), len(default__.fm15(globalState)), (d_2295_v66_)[default__.safeIndex(d_2292_v63_.f2, len(d_2295_v66_))]])
            d_2297_v68_: D11
            d_2297_v68_ = D11_DC21(d_2270_v42_, d_2292_v63_.f2, d_2292_v63_.f2)
            d_2298_v69_: D11
            d_2298_v69_ = D11_DC22(p2)
            d_2299_v70_: _dafny.MultiSet
            d_2299_v70_ = _dafny.MultiSet([d_2298_v69_, D11_DC22(p3), d_2298_v69_])
            index435_ = default__.safeIndex(630, (d_2287_v59_).length(0))
            rhs395_ = (d_2297_v68_).cf36
            rhs396_ = d_2289_v61_
            rhs397_ = (d_2295_v66_) + (d_2296_v67_)
            rhs398_ = not((d_2299_v70_).issubset(d_2299_v70_))
            lhs201_ = d_2287_v59_
            lhs202_ = default__.safeIndex(630, (d_2287_v59_).length(0))
            lhs201_[lhs202_] = rhs395_
            d_2289_v61_ = rhs396_
            d_2295_v66_ = rhs397_
            r0 = rhs398_
            d_2300_v71_: _dafny.Seq
            d_2300_v71_ = _dafny.SeqWithoutIsStrInference([d_2298_v69_])
            d_2300_v71_ = d_2300_v71_
            r2 = p0
            source34_ = d_2298_v69_
            if source34_.is_DC21:
                d_2301___mcc_h0_ = source34_.cf36
                d_2302___mcc_h1_ = source34_.cf37
                d_2303___mcc_h2_ = source34_.cf38
                d_2304_cf38_ = d_2303___mcc_h2_
                d_2305_cf37_ = d_2302___mcc_h1_
                d_2306_cf36_ = d_2301___mcc_h0_
                d_2307_v72_: _dafny.Array
                def lambda96_(d_2308_cf37_):
                    def lambda97_(d_2309_i4_):
                        return _dafny.Set({-495, d_2308_cf37_, d_2308_cf37_})

                    return lambda97_

                init50_ = lambda96_(d_2305_cf37_)
                nw425_ = _dafny.Array(None, 7)
                for i0_50_ in range(nw425_.length(0)):
                    nw425_[i0_50_] = init50_(i0_50_)
                d_2307_v72_ = nw425_
                d_2310_v74_: _dafny.Set
                d_2310_v74_ = _dafny.Set({len(d_2271_v43_)})
                index436_ = default__.safeIndex(329, (d_2307_v72_).length(0))
                def iife162_():
                    coll72_ = _dafny.Set()
                    compr_72_: int
                    for compr_72_ in _dafny.IntegerRange(-245, -526):
                        d_2311_v73_: int = compr_72_
                        if ((-245) <= (d_2311_v73_)) and ((d_2311_v73_) < (-526)):
                            coll72_ = coll72_.union(_dafny.Set([(d_2311_v73_) - (d_2292_v63_.f2)]))
                    return _dafny.Set(coll72_)
                (d_2307_v72_)[index436_] = (iife162_()
                ) | (d_2310_v74_)
                index437_ = default__.safeIndex(329, (d_2307_v72_).length(0))
                (d_2307_v72_)[index437_] = d_2310_v74_
                r0 = p2
                d_2305_cf37_ = d_2304_cf38_
                d_2312_v75_: _dafny.Seq
                d_2312_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogsgwnlyp"))
                d_2313_v76_: _dafny.Set
                d_2313_v76_ = _dafny.Set({default__.fm5(globalState)})
                d_2314_v77_: _dafny.Array
                nw426_ = _dafny.Array(None, 29)
                nw426_[int(0)] = d_2312_v75_
                nw426_[int(1)] = d_2312_v75_
                nw426_[int(2)] = default__.fm28(p2, p2, p2, d_2313_v76_, globalState)
                nw426_[int(3)] = d_2312_v75_
                nw426_[int(4)] = d_2312_v75_
                nw426_[int(5)] = d_2312_v75_
                nw426_[int(6)] = d_2312_v75_
                nw426_[int(7)] = _dafny.SeqWithoutIsStrInference([d_2306_cf36_ for d_2315_i5_ in range(default__.abs(258))])
                nw426_[int(8)] = d_2312_v75_
                nw426_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "myypb"))
                nw426_[int(10)] = (d_2312_v75_ if p2 else d_2312_v75_)
                nw426_[int(11)] = (d_2312_v75_) + (d_2312_v75_)
                nw426_[int(12)] = (d_2312_v75_ if p2 else d_2312_v75_)
                nw426_[int(13)] = d_2312_v75_
                nw426_[int(14)] = d_2312_v75_
                nw426_[int(15)] = (d_2312_v75_) + (d_2312_v75_)
                nw426_[int(16)] = _dafny.SeqWithoutIsStrInference([d_2306_cf36_ for d_2316_i6_ in range(default__.abs(167))])
                nw426_[int(17)] = d_2312_v75_
                nw426_[int(18)] = d_2312_v75_
                nw426_[int(19)] = d_2312_v75_
                nw426_[int(20)] = d_2312_v75_
                nw426_[int(21)] = d_2312_v75_
                nw426_[int(22)] = (d_2312_v75_) + (_dafny.SeqWithoutIsStrInference([d_2306_cf36_ for d_2317_i7_ in range(default__.abs(910))]))
                nw426_[int(23)] = (d_2312_v75_) + (_dafny.SeqWithoutIsStrInference([(d_2312_v75_)[default__.safeIndex(len(d_2271_v43_), len(d_2312_v75_))] for d_2318_i8_ in range(default__.abs(494))]))
                nw426_[int(24)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcnc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")))
                nw426_[int(25)] = (d_2312_v75_) + (d_2312_v75_)
                nw426_[int(26)] = d_2312_v75_
                nw426_[int(27)] = (d_2312_v75_ if p2 else d_2312_v75_)
                nw426_[int(28)] = d_2312_v75_
                d_2314_v77_ = nw426_
                index438_ = default__.safeIndex(324, (d_2314_v77_).length(0))
                (d_2314_v77_)[index438_] = d_2312_v75_
                index439_ = default__.safeIndex(324, (d_2314_v77_).length(0))
                (d_2314_v77_)[index439_] = d_2312_v75_
            elif source34_.is_DC22:
                d_2319___mcc_h3_ = source34_.cf39
                d_2320_cf39_ = d_2319___mcc_h3_
                d_2321_v78_: _dafny.Map
                d_2321_v78_ = _dafny.Map({d_2292_v63_.f2: True})
                d_2322_v79_: _dafny.Map
                d_2322_v79_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cktdo")): d_2321_v78_})
                d_2323_v80_: D5
                d_2323_v80_ = D5_DC10(p2, p0, d_2322_v79_, not(p3))
                rhs399_ = default__.fm29(globalState)
                rhs400_ = not((True) == (p3))
                d_2323_v80_ = rhs399_
                r0 = rhs400_
                d_2324_v81_: _dafny.Array
                def lambda98_(d_2325_v66_):
                    def lambda99_(d_2326_i9_):
                        return d_2325_v66_

                    return lambda99_

                init51_ = lambda98_(d_2295_v66_)
                nw427_ = _dafny.Array(None, 28)
                for i0_51_ in range(nw427_.length(0)):
                    nw427_[i0_51_] = init51_(i0_51_)
                d_2324_v81_ = nw427_
                d_2324_v81_ = d_2324_v81_
                d_2327_v82_: _dafny.Seq
                d_2327_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odck"))
                d_2328_v83_: C1
                nw428_ = C1()
                nw428_.ctor__(d_2327_v82_)
                d_2328_v83_ = nw428_
                rhs401_ = default__.safeModuloInt(p0, d_2292_v63_.f2)
                rhs402_ = d_2328_v83_
                rhs403_ = _dafny.CodePoint('j')
                lhs203_ = d_2292_v63_
                lhs203_.f2 = rhs401_
                d_2328_v83_ = rhs402_
                d_2270_v42_ = rhs403_
                d_2329_v84_: D6
                d_2329_v84_ = D6_DC11(_dafny.SeqWithoutIsStrInference([(d_2287_v59_)[default__.safeIndex(630, (d_2287_v59_).length(0))] for d_2330_i10_ in range(default__.abs(-691))]))
                d_2327_v82_ = (d_2329_v84_).cf21
            elif True:
                d_2331___mcc_h4_ = source34_.cf35
                d_2332_cf35_ = d_2331___mcc_h4_
                r0 = p3
                d_2333_v85_: _dafny.Array
                nw429_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
                d_2333_v85_ = nw429_
                d_2334_v86_: _dafny.Seq
                d_2334_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcqafja"))
                index440_ = default__.safeIndex(692, (d_2333_v85_).length(0))
                (d_2333_v85_)[index440_] = d_2334_v86_
                index441_ = default__.safeIndex(692, (d_2333_v85_).length(0))
                (d_2333_v85_)[index441_] = (_dafny.SeqWithoutIsStrInference([(d_2287_v59_)[default__.safeIndex(630, (d_2287_v59_).length(0))] for d_2335_i11_ in range(default__.abs(886))]) if p2 else d_2334_v86_)
                d_2336_v87_: _dafny.Array
                nw430_ = _dafny.Array(None, 14)
                nw430_[int(0)] = d_2292_v63_.f2
                nw430_[int(1)] = d_2292_v63_.f2
                nw430_[int(2)] = d_2292_v63_.f2
                nw430_[int(3)] = d_2292_v63_.f2
                nw430_[int(4)] = d_2292_v63_.f2
                nw430_[int(5)] = d_2292_v63_.f2
                nw430_[int(6)] = p0
                nw430_[int(7)] = d_2292_v63_.f2
                nw430_[int(8)] = -400
                nw430_[int(9)] = d_2292_v63_.f2
                nw430_[int(10)] = d_2292_v63_.f2
                nw430_[int(11)] = d_2292_v63_.f2
                nw430_[int(12)] = d_2292_v63_.f2
                nw430_[int(13)] = len((d_2333_v85_)[default__.safeIndex(692, (d_2333_v85_).length(0))])
                d_2336_v87_ = nw430_
                d_2337_v88_: _dafny.Map
                d_2337_v88_ = _dafny.Map({p0: d_2336_v87_})
                d_2338_v89_: C1
                nw431_ = C1()
                nw431_.ctor__((d_2333_v85_)[default__.safeIndex(692, (d_2333_v85_).length(0))])
                d_2338_v89_ = nw431_
                d_2339_v90_: _dafny.MultiSet
                d_2339_v90_ = _dafny.MultiSet([d_2338_v89_, d_2338_v89_])
                d_2340_v91_: _dafny.Set
                d_2340_v91_ = _dafny.Set({len(d_2337_v88_), p0, (d_2339_v90_).cardinality, (0) - (p0), -947})
                r0 = not(((_dafny.SeqWithoutIsStrInference([d_2270_v42_ for d_2341_i12_ in range(default__.abs(633))])) < (d_2334_v86_)) == ((d_2292_v63_.f2) == (len(d_2340_v91_))))
                (d_2292_v63_).f2 = default__.safeModuloInt(p0, (0) - (d_2292_v63_.f2))
        d_2342_v92_: _dafny.Array
        nw432_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 7)
        d_2342_v92_ = nw432_
        def iife163_():
            coll73_ = _dafny.Set()
            compr_73_: int
            for compr_73_ in _dafny.IntegerRange(182, 266):
                d_2343_v93_: int = compr_73_
                if ((182) <= (d_2343_v93_)) and ((d_2343_v93_) < (266)):
                    coll73_ = coll73_.union(_dafny.Set([(d_2343_v93_) - (p0)]))
            return _dafny.Set(coll73_)
        rhs404_ = len(iife163_()
        )
        rhs405_ = d_2342_v92_
        r2 = rhs404_
        d_2342_v92_ = rhs405_
        d_2344_v94_: T1
        nw433_ = C6()
        nw433_.ctor__()
        d_2344_v94_ = nw433_
        d_2345_v95_: _dafny.Seq
        d_2345_v95_ = _dafny.SeqWithoutIsStrInference([d_2270_v42_])
        d_2346_v96_: _dafny.Map
        d_2346_v96_ = _dafny.Map({p0: len((d_2345_v95_).set(default__.safeIndex(p0, len(d_2345_v95_)), d_2270_v42_))})
        d_2347_v97_: _dafny.Map
        d_2347_v97_ = _dafny.Map({d_2344_v94_: d_2346_v96_})
        d_2348_v98_: _dafny.Set
        d_2348_v98_ = _dafny.Set({p0, p0})
        d_2349_v99_: _dafny.Seq
        d_2349_v99_ = _dafny.SeqWithoutIsStrInference([True])
        d_2350_v100_: _dafny.MultiSet
        d_2350_v100_ = _dafny.MultiSet([p2])
        d_2351_v101_: _dafny.Array
        nw434_ = _dafny.Array(None, 14)
        nw434_[int(0)] = (D6_DC13(p3, p0)).cf25
        nw434_[int(1)] = (0) - ((0) - (((0) - (p0) if p3 else p0)))
        nw434_[int(2)] = p0
        nw434_[int(3)] = len(((d_2347_v97_)[d_2344_v94_] if (d_2344_v94_) in (d_2347_v97_) else d_2346_v96_))
        nw434_[int(4)] = len((d_2348_v98_).intersection(d_2348_v98_))
        nw434_[int(5)] = default__.safeDivisionInt(len(d_2349_v99_), p0)
        nw434_[int(6)] = default__.safeDivisionInt(p0, p0)
        nw434_[int(7)] = p0
        nw434_[int(8)] = (d_2225_v1_).cardinality
        nw434_[int(9)] = (d_2350_v100_).cardinality
        nw434_[int(10)] = p0
        nw434_[int(11)] = p0
        nw434_[int(12)] = p0
        nw434_[int(13)] = p0
        d_2351_v101_ = nw434_
        d_2352_v102_: _dafny.Seq
        d_2352_v102_ = _dafny.SeqWithoutIsStrInference([(self).fm1(p2, globalState), p0, p0, (0) - (p0), len(d_2349_v99_)])
        index442_ = default__.safeIndex(35, (d_2351_v101_).length(0))
        (d_2351_v101_)[index442_] = (0) - ((default__.fm50(not(p2), not(not(p3)), p2, (d_2352_v102_)[default__.safeIndex((0) - (len(d_2271_v43_)), len(d_2352_v102_))], globalState)).cf27)
        d_2353_v103_: _dafny.Map
        d_2353_v103_ = _dafny.Map({p2: p0})
        index443_ = default__.safeIndex(35, (d_2351_v101_).length(0))
        (d_2351_v101_)[index443_] = (0) - ((0) - (default__.safeDivisionInt(((d_2353_v103_)[p3] if (p3) in (d_2353_v103_) else p0), len(d_2352_v102_))))
        d_2354_v104_: _dafny.Array
        nw435_ = _dafny.Array(False, 11)
        d_2354_v104_ = nw435_
        index444_ = default__.safeIndex(511, (d_2354_v104_).length(0))
        (d_2354_v104_)[index444_] = p3
        d_2355_v105_: _dafny.Map
        d_2355_v105_ = _dafny.Map({(d_2351_v101_)[default__.safeIndex(35, (d_2351_v101_).length(0))]: d_2350_v100_})
        index445_ = default__.safeIndex(511, (d_2354_v104_).length(0))
        (d_2354_v104_)[index445_] = (d_2350_v100_).ispropersubset(((d_2355_v105_)[p0] if (p0) in (d_2355_v105_) else d_2350_v100_))
        index446_ = default__.safeIndex(35, (d_2351_v101_).length(0))
        (d_2351_v101_)[index446_] = len(d_2349_v99_)
        r0 = p3
        d_2356_v106_: _dafny.Array
        nw436_ = _dafny.Array(_dafny.Map({}), 28)
        d_2356_v106_ = nw436_
        r1 = d_2356_v106_
        r2 = (p0) * (p0)
        return r0, r1, r2

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_2357_v0_: bool
        d_2357_v0_ = True
        d_2357_v0_ = p2
        if False:
            d_2358_v1_: _dafny.Map
            d_2358_v1_ = _dafny.Map({p1: p2})
            d_2358_v1_ = (d_2358_v1_).set(default__.safeDivisionInt(-409, p0), d_2357_v0_)
            d_2359_v2_: T2
            nw437_ = C3()
            nw437_.ctor__(((0) - (p1)) + (p0))
            d_2359_v2_ = nw437_
            d_2359_v2_ = d_2359_v2_
            d_2360_v3_: _dafny.MultiSet
            d_2360_v3_ = _dafny.MultiSet([not(True), p2, d_2357_v0_, not(not(d_2357_v0_))])
            d_2361_v4_: _dafny.Seq
            d_2361_v4_ = _dafny.SeqWithoutIsStrInference([(d_2360_v3_).set(True, default__.abs(525))])
            d_2361_v4_ = _dafny.SeqWithoutIsStrInference([(d_2360_v3_) | ((d_2361_v4_)[default__.safeIndex(p0, len(d_2361_v4_))])])
            d_2362_v5_: str
            d_2362_v5_ = _dafny.CodePoint('r')
            d_2363_v6_: _dafny.Seq
            d_2363_v6_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c'), d_2362_v5_])
            d_2364_v7_: _dafny.Map
            d_2364_v7_ = _dafny.Map({(d_2363_v6_) + ((_dafny.SeqWithoutIsStrInference([d_2362_v5_, _dafny.CodePoint('k')])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([d_2362_v5_, _dafny.CodePoint('k')]))), d_2362_v5_)): d_2357_v0_})
            if ((d_2364_v7_)[(_dafny.SeqWithoutIsStrInference([d_2362_v5_ for d_2365_i0_ in range(default__.abs(307))])).set(default__.safeIndex(len(d_2363_v6_), len(_dafny.SeqWithoutIsStrInference([d_2362_v5_ for d_2366_i0_ in range(default__.abs(307))]))), d_2362_v5_)] if ((_dafny.SeqWithoutIsStrInference([d_2362_v5_ for d_2367_i0_ in range(default__.abs(307))])).set(default__.safeIndex(len(d_2363_v6_), len(_dafny.SeqWithoutIsStrInference([d_2362_v5_ for d_2368_i0_ in range(default__.abs(307))]))), d_2362_v5_)) in (d_2364_v7_) else not(p2)):
                d_2369_v8_: _dafny.Set
                d_2369_v8_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_2362_v5_ for d_2370_i1_ in range(default__.abs(227))]), d_2363_v6_, d_2363_v6_, d_2363_v6_})
                d_2371_v9_: _dafny.Array
                nw438_ = _dafny.Array(None, 5)
                nw438_[int(0)] = d_2359_v2_.f2
                nw438_[int(1)] = -808
                nw438_[int(2)] = d_2359_v2_.f2
                nw438_[int(3)] = len(d_2369_v8_)
                nw438_[int(4)] = p0
                d_2371_v9_ = nw438_
                index447_ = default__.safeIndex(446, (d_2371_v9_).length(0))
                (d_2371_v9_)[index447_] = p0
                index448_ = default__.safeIndex(446, (d_2371_v9_).length(0))
                (d_2371_v9_)[index448_] = (d_2359_v2_.f2) + (p1)
                d_2372_v10_: _dafny.Seq
                d_2372_v10_ = _dafny.SeqWithoutIsStrInference([p1])
                d_2373_v11_: _dafny.Map
                d_2373_v11_ = _dafny.Map({d_2372_v10_: d_2359_v2_.f2})
                d_2373_v11_ = d_2373_v11_
                d_2374_v12_: _dafny.Set
                d_2374_v12_ = _dafny.Set({d_2357_v0_, p2})
                rhs406_ = not(default__.fm3(p2, (len(_dafny.SeqWithoutIsStrInference([False, p2])) if p2 else (d_2371_v9_)[default__.safeIndex(446, (d_2371_v9_).length(0))]), len(default__.fm15(globalState)), globalState))
                rhs407_ = (len(d_2374_v12_)) + ((0) - (d_2359_v2_.f2))
                rhs408_ = d_2357_v0_
                rhs409_ = d_2363_v6_
                lhs204_ = d_2359_v2_
                d_2357_v0_ = rhs406_
                lhs204_.f2 = rhs407_
                d_2357_v0_ = rhs408_
                d_2363_v6_ = rhs409_
                d_2362_v5_ = d_2362_v5_
                d_2357_v0_ = d_2357_v0_
            elif True:
                d_2357_v0_ = (d_2359_v2_.f2) <= (len(d_2358_v1_))
                d_2375_v13_: C4
                nw439_ = C4()
                nw439_.ctor__(default__.safeModuloInt(p1, (d_2360_v3_).cardinality))
                d_2375_v13_ = nw439_
                d_2376_v14_: _dafny.Seq
                d_2376_v14_ = _dafny.SeqWithoutIsStrInference([p2, p2])
                d_2377_v15_: _dafny.Map
                d_2377_v15_ = _dafny.Map({d_2376_v14_: p1})
                d_2377_v15_ = (d_2377_v15_).set(d_2376_v14_, (0) - ((d_2375_v13_).fm1(d_2357_v0_, globalState)))
                d_2378_v16_: _dafny.Array
                nw440_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
                d_2378_v16_ = nw440_
                def lambda100_(d_2379_v6_):
                    def lambda101_(d_2380_i2_):
                        return d_2379_v6_

                    return lambda101_

                init52_ = lambda100_(d_2363_v6_)
                nw441_ = _dafny.Array(None, 9)
                for i0_52_ in range(nw441_.length(0)):
                    nw441_[i0_52_] = init52_(i0_52_)
                d_2378_v16_ = nw441_
                d_2381_v17_: _dafny.Seq
                d_2381_v17_ = _dafny.SeqWithoutIsStrInference([d_2359_v2_.f2, d_2359_v2_.f2])
                d_2382_v18_: _dafny.MultiSet
                d_2382_v18_ = _dafny.MultiSet([d_2381_v17_])
                d_2383_v19_: _dafny.Map
                d_2383_v19_ = _dafny.Map({(_dafny.MultiSet([d_2357_v0_])).ispropersubset(d_2360_v3_): d_2382_v18_})
                d_2383_v19_ = (d_2383_v19_).set((-163) > (d_2359_v2_.f2), _dafny.MultiSet([d_2381_v17_, d_2381_v17_, d_2381_v17_]))
            if not((p2) == ((d_2359_v2_).fm8((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvqbna"))).set(default__.safeIndex(d_2359_v2_.f2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvqbna")))), _dafny.CodePoint('k')), d_2357_v0_, globalState))):
                d_2384_v20_: _dafny.MultiSet
                d_2384_v20_ = _dafny.MultiSet([d_2359_v2_.f2])
                d_2385_v21_: _dafny.Seq
                d_2385_v21_ = _dafny.SeqWithoutIsStrInference([d_2363_v6_])
                d_2386_v22_: _dafny.Seq
                d_2386_v22_ = _dafny.SeqWithoutIsStrInference([d_2357_v0_, p2])
                d_2387_v23_: D6
                d_2387_v23_ = D6_DC12((d_2386_v22_)[default__.safeIndex(p1, len(d_2386_v22_))], d_2357_v0_)
                d_2388_v24_: _dafny.Set
                d_2388_v24_ = _dafny.Set({d_2359_v2_.f2, d_2359_v2_.f2})
                d_2389_v25_: D11
                d_2389_v25_ = D11_DC21(d_2362_v5_, d_2359_v2_.f2, p0)
                d_2390_v26_: _dafny.Array
                nw442_ = _dafny.Array(None, 10)
                nw442_[int(0)] = d_2357_v0_
                nw442_[int(1)] = p2
                nw442_[int(2)] = d_2357_v0_
                nw442_[int(3)] = not(p2)
                nw442_[int(4)] = d_2357_v0_
                nw442_[int(5)] = d_2357_v0_
                nw442_[int(6)] = d_2357_v0_
                nw442_[int(7)] = True
                nw442_[int(8)] = p2
                nw442_[int(9)] = p2
                d_2390_v26_ = nw442_
                d_2391_v27_: _dafny.Map
                d_2391_v27_ = _dafny.Map({d_2390_v26_: d_2359_v2_.f2})
                d_2392_v28_: C1
                nw443_ = C1()
                nw443_.ctor__(d_2363_v6_)
                d_2392_v28_ = nw443_
                d_2393_v29_: D8
                d_2393_v29_ = D8_DC17(d_2357_v0_, d_2392_v28_, p1)
                d_2394_v30_: _dafny.Seq
                d_2394_v30_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2395_v31_: _dafny.Array
                nw444_ = _dafny.Array(None, 26)
                nw444_[int(0)] = (d_2359_v2_.f2) + (p1)
                nw444_[int(1)] = (d_2384_v20_).cardinality
                nw444_[int(2)] = (d_2359_v2_.f2) + (d_2359_v2_.f2)
                nw444_[int(3)] = -705
                nw444_[int(4)] = d_2359_v2_.f2
                nw444_[int(5)] = d_2359_v2_.f2
                nw444_[int(6)] = len(((d_2385_v21_)[default__.safeIndex((self).fm1(p2, globalState), len(d_2385_v21_))]) + (d_2363_v6_))
                nw444_[int(7)] = len((d_2363_v6_ if p2 else d_2363_v6_))
                nw444_[int(8)] = p1
                nw444_[int(9)] = (d_2359_v2_.f2 if p2 else p0)
                nw444_[int(10)] = p0
                nw444_[int(11)] = d_2359_v2_.f2
                nw444_[int(12)] = d_2359_v2_.f2
                nw444_[int(13)] = (d_2359_v2_.f2 if (d_2387_v23_).cf23 else d_2359_v2_.f2)
                nw444_[int(14)] = (len(d_2386_v22_)) - (d_2359_v2_.f2)
                nw444_[int(15)] = d_2359_v2_.f2
                nw444_[int(16)] = (d_2359_v2_.f2) * (d_2359_v2_.f2)
                nw444_[int(17)] = d_2359_v2_.f2
                nw444_[int(18)] = (p0) * (len(d_2388_v24_))
                nw444_[int(19)] = (d_2389_v25_).cf37
                nw444_[int(20)] = default__.safeDivisionInt(d_2359_v2_.f2, p0)
                nw444_[int(21)] = default__.safeModuloInt(p0, len(d_2391_v27_))
                nw444_[int(22)] = (d_2393_v29_).cf32
                nw444_[int(23)] = default__.safeModuloInt(d_2359_v2_.f2, p0)
                nw444_[int(24)] = p0
                nw444_[int(25)] = len(d_2394_v30_)
                d_2395_v31_ = nw444_
                index449_ = default__.safeIndex(276, (d_2395_v31_).length(0))
                (d_2395_v31_)[index449_] = p0
                d_2396_v32_: _dafny.Set
                d_2396_v32_ = _dafny.Set({d_2390_v26_})
                index450_ = default__.safeIndex(276, (d_2395_v31_).length(0))
                (d_2395_v31_)[index450_] = (len(d_2396_v32_)) * ((0) - (d_2359_v2_.f2))
                d_2397_v33_: C2
                nw445_ = C2()
                nw445_.ctor__(d_2359_v2_.f2)
                d_2397_v33_ = nw445_
                (d_2359_v2_).f2 = p1
                d_2398_v34_: _dafny.Array
                nw446_ = _dafny.Array(None, 25)
                nw446_[int(0)] = _dafny.SeqWithoutIsStrInference([d_2359_v2_.f2 for d_2399_i3_ in range(default__.abs(522))])
                nw446_[int(1)] = _dafny.SeqWithoutIsStrInference([p1 for d_2400_i4_ in range(default__.abs(442))])
                nw446_[int(2)] = _dafny.SeqWithoutIsStrInference([d_2359_v2_.f2])
                nw446_[int(3)] = (_dafny.SeqWithoutIsStrInference([p0 for d_2401_i5_ in range(default__.abs(-633))])).set(default__.safeIndex(d_2359_v2_.f2, len(_dafny.SeqWithoutIsStrInference([p0 for d_2402_i5_ in range(default__.abs(-633))]))), len(d_2358_v1_))
                nw446_[int(4)] = d_2394_v30_
                nw446_[int(5)] = _dafny.SeqWithoutIsStrInference([len((d_2392_v28_).f3), p0])
                nw446_[int(6)] = d_2394_v30_
                nw446_[int(7)] = d_2394_v30_
                nw446_[int(8)] = d_2394_v30_
                nw446_[int(9)] = d_2394_v30_
                nw446_[int(10)] = _dafny.SeqWithoutIsStrInference([len(d_2394_v30_), len(d_2394_v30_)])
                nw446_[int(11)] = _dafny.SeqWithoutIsStrInference([p1 for d_2403_i6_ in range(default__.abs(961))])
                nw446_[int(12)] = _dafny.SeqWithoutIsStrInference([(d_2360_v3_).cardinality])
                nw446_[int(13)] = d_2394_v30_
                nw446_[int(14)] = d_2394_v30_
                nw446_[int(15)] = d_2394_v30_
                nw446_[int(16)] = d_2394_v30_
                nw446_[int(17)] = d_2394_v30_
                nw446_[int(18)] = _dafny.SeqWithoutIsStrInference([d_2359_v2_.f2, d_2359_v2_.f2])
                nw446_[int(19)] = _dafny.SeqWithoutIsStrInference([d_2359_v2_.f2])
                nw446_[int(20)] = _dafny.SeqWithoutIsStrInference([d_2359_v2_.f2])
                nw446_[int(21)] = _dafny.SeqWithoutIsStrInference([(d_2395_v31_)[default__.safeIndex(276, (d_2395_v31_).length(0))], (d_2397_v33_).fm1(p2, globalState)])
                nw446_[int(22)] = _dafny.SeqWithoutIsStrInference([d_2359_v2_.f2 for d_2404_i7_ in range(default__.abs(54))])
                nw446_[int(23)] = _dafny.SeqWithoutIsStrInference([d_2359_v2_.f2 for d_2405_i8_ in range(default__.abs(383))])
                nw446_[int(24)] = _dafny.SeqWithoutIsStrInference([d_2359_v2_.f2 for d_2406_i9_ in range(default__.abs(842))])
                d_2398_v34_ = nw446_
                d_2407_v35_: bool
                d_2408_v36_: _dafny.Array
                d_2409_v37_: int
                out51_: bool
                out52_: _dafny.Array
                out53_: int
                out51_, out52_, out53_ = (d_2392_v28_).m1((len(d_2394_v30_)) + (p1), d_2398_v34_, p2, p2, globalState)
                d_2407_v35_ = out51_
                d_2408_v36_ = out52_
                d_2409_v37_ = out53_
                d_2410_v38_: _dafny.Map
                d_2410_v38_ = _dafny.Map({p2: default__.safeDivisionInt((d_2359_v2_).fm1((d_2392_v28_).fm11(globalState), globalState), d_2359_v2_.f2)})
                d_2410_v38_ = (d_2410_v38_).set(d_2357_v0_, d_2409_v37_)
            elif True:
                d_2357_v0_ = not(p2)
                d_2357_v0_ = d_2357_v0_
                r0 = d_2359_v2_.f2
                d_2411_v39_: _dafny.Seq
                d_2411_v39_ = _dafny.SeqWithoutIsStrInference([p0, d_2359_v2_.f2, d_2359_v2_.f2, p0])
                d_2412_v40_: _dafny.Map
                d_2412_v40_ = _dafny.Map({d_2362_v5_: len(d_2411_v39_)})
                d_2412_v40_ = ((_dafny.Map({d_2362_v5_: p1})) | (d_2412_v40_)).set(d_2362_v5_, p0)
                d_2413_v41_: _dafny.Array
                nw447_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_2413_v41_ = nw447_
                d_2414_v43_: _dafny.Set
                d_2414_v43_ = _dafny.Set({d_2359_v2_.f2, (d_2359_v2_).fm1(True, globalState), p1})
                def iife164_():
                    coll74_ = _dafny.Set()
                    compr_74_: int
                    for compr_74_ in _dafny.IntegerRange(329, 686):
                        d_2415_v42_: int = compr_74_
                        if ((329) <= (d_2415_v42_)) and ((d_2415_v42_) < (686)):
                            coll74_ = coll74_.union(_dafny.Set([(d_2415_v42_) - (270)]))
                    return _dafny.Set(coll74_)
                rhs410_ = len((iife164_()
                ) | (d_2414_v43_))
                rhs411_ = d_2413_v41_
                rhs412_ = p0
                r0 = rhs410_
                d_2413_v41_ = rhs411_
                r0 = rhs412_
        elif True:
            d_2416_v44_: _dafny.Seq
            d_2416_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jy"))
            d_2417_v45_: _dafny.Map
            d_2417_v45_ = _dafny.Map({d_2357_v0_: p0})
            d_2418_v47_: _dafny.Seq
            d_2418_v47_ = _dafny.SeqWithoutIsStrInference([len(d_2416_v44_)])
            d_2419_v48_: _dafny.Seq
            d_2419_v48_ = _dafny.SeqWithoutIsStrInference([d_2357_v0_, True])
            d_2420_v49_: _dafny.Map
            d_2420_v49_ = _dafny.Map({p1: (d_2418_v47_)[default__.safeIndex(len(d_2419_v48_), len(d_2418_v47_))]})
            d_2421_v50_: _dafny.Set
            d_2421_v50_ = _dafny.Set({p0, len(d_2420_v49_)})
            d_2422_v51_: _dafny.Seq
            def iife165_():
                coll75_ = _dafny.Set()
                compr_75_: int
                for compr_75_ in _dafny.IntegerRange(345, 363):
                    d_2423_v46_: int = compr_75_
                    if ((345) <= (d_2423_v46_)) and ((d_2423_v46_) < (363)):
                        coll75_ = coll75_.union(_dafny.Set([(d_2423_v46_) + (p1)]))
                return _dafny.Set(coll75_)
            d_2422_v51_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({len(d_2416_v44_), p0, ((d_2417_v45_)[d_2357_v0_] if (d_2357_v0_) in (d_2417_v45_) else p1)}), iife165_()
            , d_2421_v50_])
            d_2424_v52_: D7
            d_2424_v52_ = D7_DC14(d_2418_v47_)
            d_2425_v53_: _dafny.Seq
            d_2425_v53_ = _dafny.SeqWithoutIsStrInference([d_2424_v52_, d_2424_v52_])
            d_2426_v54_: _dafny.Array
            def lambda102_(d_2427_p2_, d_2428_p1_, d_2429_v44_):
                def lambda103_(d_2430_i10_):
                    return (d_2430_i10_) + ((D5_DC10(d_2427_p2_, d_2428_p1_, _dafny.Map({d_2429_v44_: _dafny.Map({d_2428_p1_: d_2427_p2_})}), d_2427_p2_)).cf18)

                return lambda103_

            init53_ = lambda102_(p2, p1, d_2416_v44_)
            nw448_ = _dafny.Array(None, 5)
            for i0_53_ in range(nw448_.length(0)):
                nw448_[i0_53_] = init53_(i0_53_)
            d_2426_v54_ = nw448_
            d_2431_v55_: _dafny.Seq
            d_2431_v55_ = _dafny.SeqWithoutIsStrInference([d_2426_v54_])
            d_2432_v56_: _dafny.MultiSet
            d_2432_v56_ = _dafny.MultiSet([d_2357_v0_, False, p2, d_2357_v0_])
            d_2433_v57_: _dafny.Array
            nw449_ = _dafny.Array(None, 27)
            nw449_[int(0)] = (p0) > (len(d_2422_v51_))
            nw449_[int(1)] = (len(d_2425_v53_)) > ((self).fm1(d_2357_v0_, globalState))
            nw449_[int(2)] = d_2357_v0_
            nw449_[int(3)] = (d_2357_v0_) == (p2)
            nw449_[int(4)] = (True) and (not(False))
            nw449_[int(5)] = d_2357_v0_
            nw449_[int(6)] = p2
            nw449_[int(7)] = (-180) == (len(d_2431_v55_))
            nw449_[int(8)] = d_2357_v0_
            nw449_[int(9)] = not(p2)
            nw449_[int(10)] = d_2357_v0_
            nw449_[int(11)] = d_2357_v0_
            nw449_[int(12)] = (d_2357_v0_) and (False)
            nw449_[int(13)] = p2
            nw449_[int(14)] = (d_2432_v56_).issubset(d_2432_v56_)
            nw449_[int(15)] = p2
            nw449_[int(16)] = d_2357_v0_
            nw449_[int(17)] = default__.fm3(d_2357_v0_, p1, p1, globalState)
            nw449_[int(18)] = (p0) in (_dafny.SeqWithoutIsStrInference([(self).fm1(d_2357_v0_, globalState), (d_2418_v47_)[default__.safeIndex(len(d_2416_v44_), len(d_2418_v47_))]]))
            nw449_[int(19)] = (p0) >= (p0)
            nw449_[int(20)] = d_2357_v0_
            nw449_[int(21)] = False
            nw449_[int(22)] = not (p2) or (d_2357_v0_)
            nw449_[int(23)] = d_2357_v0_
            nw449_[int(24)] = p2
            nw449_[int(25)] = d_2357_v0_
            nw449_[int(26)] = (d_2416_v44_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krmp")))
            d_2433_v57_ = nw449_
            index451_ = default__.safeIndex(33, (d_2433_v57_).length(0))
            (d_2433_v57_)[index451_] = default__.fm5(globalState)
            index452_ = default__.safeIndex(33, (d_2433_v57_).length(0))
            rhs413_ = ((d_2416_v44_) + (d_2416_v44_)) + (d_2416_v44_)
            rhs414_ = p0
            rhs415_ = d_2357_v0_
            lhs205_ = d_2433_v57_
            lhs206_ = default__.safeIndex(33, (d_2433_v57_).length(0))
            d_2416_v44_ = rhs413_
            r0 = rhs414_
            lhs205_[lhs206_] = rhs415_
            if not((d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))]):
                d_2434_v58_: _dafny.Array
                nw450_ = _dafny.Array(_dafny.Array(None, 0), 14)
                d_2434_v58_ = nw450_
                index453_ = default__.safeIndex(767, (d_2434_v58_).length(0))
                (d_2434_v58_)[index453_] = d_2426_v54_
                index454_ = default__.safeIndex(767, (d_2434_v58_).length(0))
                (d_2434_v58_)[index454_] = d_2426_v54_
                d_2435_v59_: _dafny.Map
                d_2435_v59_ = _dafny.Map({(d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))]: _dafny.CodePoint('r')})
                d_2436_v60_: _dafny.Set
                d_2436_v60_ = _dafny.Set({d_2435_v59_})
                index455_ = default__.safeIndex(33, (d_2433_v57_).length(0))
                rhs416_ = p0
                rhs417_ = (d_2436_v60_) | ((_dafny.Set({d_2435_v59_})) | (d_2436_v60_))
                rhs418_ = d_2357_v0_
                rhs419_ = (p1) == (len(_dafny.Set({(d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))]})))
                lhs207_ = d_2433_v57_
                lhs208_ = default__.safeIndex(33, (d_2433_v57_).length(0))
                r0 = rhs416_
                d_2436_v60_ = rhs417_
                d_2357_v0_ = rhs418_
                lhs207_[lhs208_] = rhs419_
                d_2437_v61_: str
                d_2437_v61_ = _dafny.CodePoint('b')
                d_2438_v62_: D5
                d_2438_v62_ = D5_DC9((d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))], 69, d_2432_v56_)
                index456_ = default__.safeIndex(33, (d_2433_v57_).length(0))
                rhs420_ = (p2) and ((d_2419_v48_)[default__.safeIndex((0) - (p0), len(d_2419_v48_))])
                rhs421_ = (0) - (((d_2438_v62_).cf15 if (d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))] else p0))
                rhs422_ = d_2437_v61_
                lhs209_ = d_2433_v57_
                lhs210_ = default__.safeIndex(33, (d_2433_v57_).length(0))
                lhs209_[lhs210_] = rhs420_
                r0 = rhs421_
                d_2437_v61_ = rhs422_
                d_2439_v63_: T0
                nw451_ = C6()
                nw451_.ctor__()
                d_2439_v63_ = nw451_
                index457_ = default__.safeIndex(33, (d_2433_v57_).length(0))
                rhs423_ = d_2439_v63_
                rhs424_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ongtjwil"))) + (_dafny.SeqWithoutIsStrInference([d_2437_v61_ for d_2440_i11_ in range(default__.abs(959))]))
                rhs425_ = ((d_2419_v48_) + (d_2419_v48_)) + (_dafny.SeqWithoutIsStrInference([d_2357_v0_, False]))
                rhs426_ = (d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))]
                lhs211_ = d_2433_v57_
                lhs212_ = default__.safeIndex(33, (d_2433_v57_).length(0))
                d_2439_v63_ = rhs423_
                d_2416_v44_ = rhs424_
                d_2419_v48_ = rhs425_
                lhs211_[lhs212_] = rhs426_
                d_2441_v64_: _dafny.Array
                nw452_ = _dafny.Array(_dafny.Set({}), 4)
                d_2441_v64_ = nw452_
                index458_ = default__.safeIndex(479, (d_2441_v64_).length(0))
                (d_2441_v64_)[index458_] = d_2421_v50_
                index459_ = default__.safeIndex(33, (d_2433_v57_).length(0))
                index460_ = default__.safeIndex(479, (d_2441_v64_).length(0))
                rhs427_ = not((d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))])
                rhs428_ = (p1) <= ((p1 if d_2357_v0_ else len(d_2435_v59_)))
                rhs429_ = d_2433_v57_
                rhs430_ = _dafny.Set({353, p1, (p0) + (573)})
                rhs431_ = d_2357_v0_
                lhs213_ = d_2433_v57_
                lhs214_ = default__.safeIndex(33, (d_2433_v57_).length(0))
                lhs215_ = d_2441_v64_
                lhs216_ = default__.safeIndex(479, (d_2441_v64_).length(0))
                d_2357_v0_ = rhs427_
                lhs213_[lhs214_] = rhs428_
                d_2433_v57_ = rhs429_
                lhs215_[lhs216_] = rhs430_
                d_2357_v0_ = rhs431_
            elif True:
                d_2442_v65_: D6
                d_2442_v65_ = D6_DC13((d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))], (self).fm1(d_2357_v0_, globalState))
                d_2357_v0_ = (d_2442_v65_).cf24
                d_2443_v66_: str
                d_2443_v66_ = _dafny.CodePoint('w')
                d_2417_v45_ = default__.fm47(d_2443_v66_, len((d_2419_v48_).set(default__.safeIndex(p0, len(d_2419_v48_)), not((d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))]))), p1, globalState)
                index461_ = default__.safeIndex(33, (d_2433_v57_).length(0))
                (d_2433_v57_)[index461_] = ((d_2432_v56_) - (d_2432_v56_)).ispropersubset(d_2432_v56_)
                d_2444_v67_: _dafny.Set
                d_2444_v67_ = _dafny.Set({d_2357_v0_})
                index462_ = default__.safeIndex(33, (d_2433_v57_).length(0))
                def iife166_():
                    coll76_ = _dafny.Set()
                    compr_76_: int
                    for compr_76_ in _dafny.IntegerRange(879, 273):
                        d_2445_v68_: int = compr_76_
                        if ((879) <= (d_2445_v68_)) and ((d_2445_v68_) < (273)):
                            coll76_ = coll76_.union(_dafny.Set([default__.safeDivisionInt(d_2445_v68_, p1)]))
                    return _dafny.Set(coll76_)
                rhs432_ = (d_2416_v44_) + ((d_2416_v44_) + (d_2416_v44_))
                rhs433_ = ((_dafny.Set({True, d_2357_v0_, (d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))], d_2357_v0_, p2})).intersection(d_2444_v67_)).isdisjoint(d_2444_v67_)
                rhs434_ = not(((p0) * (len(default__.fm49(d_2419_v48_, (d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))], p0, globalState)))) < (len(iife166_()
                )))
                lhs217_ = d_2433_v57_
                lhs218_ = default__.safeIndex(33, (d_2433_v57_).length(0))
                d_2416_v44_ = rhs432_
                d_2357_v0_ = rhs433_
                lhs217_[lhs218_] = rhs434_
                d_2446_v69_: D14
                d_2446_v69_ = D14_DC28()
                d_2447_v70_: _dafny.Seq
                d_2447_v70_ = _dafny.SeqWithoutIsStrInference([d_2446_v69_, default__.fm51(d_2416_v44_, False, d_2357_v0_, globalState)])
                rhs435_ = (d_2357_v0_) or (default__.fm5(globalState))
                rhs436_ = d_2447_v70_
                d_2357_v0_ = rhs435_
                d_2447_v70_ = rhs436_
            d_2448_v71_: str
            d_2448_v71_ = _dafny.CodePoint('k')
            d_2449_v72_: _dafny.Map
            d_2449_v72_ = _dafny.Map({len(default__.fm36((d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))], d_2448_v71_, p0, globalState)): p2})
            d_2450_v73_: _dafny.Array
            def lambda104_(d_2451_v49_):
                def lambda105_(d_2452_i12_):
                    return d_2451_v49_

                return lambda105_

            init54_ = lambda104_(d_2420_v49_)
            nw453_ = _dafny.Array(None, 16)
            for i0_54_ in range(nw453_.length(0)):
                nw453_[i0_54_] = init54_(i0_54_)
            d_2450_v73_ = nw453_
            d_2453_v74_: _dafny.Array
            d_2453_v74_ = d_2450_v73_
            d_2454_v75_: _dafny.MultiSet
            d_2454_v75_ = _dafny.MultiSet([d_2453_v74_, d_2453_v74_, d_2453_v74_, d_2453_v74_])
            d_2455_v76_: _dafny.Map
            d_2455_v76_ = _dafny.Map({p0: d_2454_v75_})
            d_2456_v77_: _dafny.Map
            d_2456_v77_ = _dafny.Map({(_dafny.Map({len(d_2449_v72_): _dafny.MultiSet([d_2453_v74_])})) | (d_2455_v76_): d_2357_v0_})
            d_2456_v77_ = (d_2456_v77_).set(_dafny.Map({p0: d_2454_v75_}), not (p2) or (d_2357_v0_))
            d_2457_v78_: C0
            nw454_ = C0()
            nw454_.ctor__((d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))], (d_2433_v57_)[default__.safeIndex(33, (d_2433_v57_).length(0))])
            d_2457_v78_ = nw454_
            d_2458_v79_: _dafny.MultiSet
            d_2458_v79_ = _dafny.MultiSet([p0])
            index463_ = default__.safeIndex(33, (d_2433_v57_).length(0))
            rhs437_ = (d_2416_v44_) + ((_dafny.SeqWithoutIsStrInference([d_2448_v71_ for d_2459_i13_ in range(default__.abs(870))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgyncltxh"))))
            rhs438_ = (p1) <= ((d_2458_v79_).cardinality)
            lhs219_ = d_2433_v57_
            lhs220_ = default__.safeIndex(33, (d_2433_v57_).length(0))
            d_2416_v44_ = rhs437_
            lhs219_[lhs220_] = rhs438_
        d_2460_v80_: _dafny.Seq
        d_2460_v80_ = _dafny.SeqWithoutIsStrInference([True, default__.fm3(False, p0, p1, globalState), d_2357_v0_, p2])
        d_2461_v81_: _dafny.MultiSet
        d_2461_v81_ = _dafny.MultiSet([not(d_2357_v0_)])
        d_2462_v82_: _dafny.Array
        nw455_ = _dafny.Array(None, 10)
        nw455_[int(0)] = d_2357_v0_
        nw455_[int(1)] = True
        nw455_[int(2)] = not(not(d_2357_v0_))
        nw455_[int(3)] = d_2357_v0_
        nw455_[int(4)] = (d_2460_v80_)[default__.safeIndex((d_2461_v81_).cardinality, len(d_2460_v80_))]
        nw455_[int(5)] = p2
        nw455_[int(6)] = (d_2460_v80_) < (d_2460_v80_)
        nw455_[int(7)] = d_2357_v0_
        nw455_[int(8)] = p2
        nw455_[int(9)] = not((p2) and (p2))
        d_2462_v82_ = nw455_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_2462_v82_).length(0)):
            d_2463_i14_: int = guard_loop_5_
            if (True) and (((0) <= (d_2463_i14_)) and ((d_2463_i14_) < ((d_2462_v82_).length(0)))):
                (d_2462_v82_)[(d_2463_i14_)] = (default__.safeDivisionInt(p0, len(_dafny.SeqWithoutIsStrInference([p0])))) > (900)
        d_2464_v83_: D6
        d_2464_v83_ = D6_DC13(not(d_2357_v0_), p1)
        pat_let_tv56_ = p0
        pat_let_tv57_ = p1
        pat_let_tv58_ = p1
        def lambda106_(source35_):
            if source35_.is_DC12:
                d_2465___mcc_h0_ = source35_.cf22
                d_2466___mcc_h1_ = source35_.cf23
                d_2467_cf23_ = d_2466___mcc_h1_
                d_2468_cf22_ = d_2465___mcc_h0_
                return pat_let_tv56_
            elif source35_.is_DC13:
                d_2469___mcc_h2_ = source35_.cf24
                d_2470___mcc_h3_ = source35_.cf25
                d_2471_cf25_ = d_2470___mcc_h3_
                d_2472_cf24_ = d_2469___mcc_h2_
                return pat_let_tv57_
            elif True:
                d_2473___mcc_h4_ = source35_.cf21
                d_2474_cf21_ = d_2473___mcc_h4_
                return pat_let_tv58_

        r0 = lambda106_(d_2464_v83_)
        d_2475_v84_: _dafny.Array
        nw456_ = _dafny.Array(int(0), 5)
        d_2475_v84_ = nw456_
        index464_ = default__.safeIndex(630, (d_2475_v84_).length(0))
        (d_2475_v84_)[index464_] = default__.safeModuloInt(308, p1)
        index465_ = default__.safeIndex(630, (d_2475_v84_).length(0))
        (d_2475_v84_)[index465_] = p1
        d_2476_v85_: D4
        d_2476_v85_ = D4_DC6(p0, d_2475_v84_, p0)
        d_2476_v85_ = d_2476_v85_
        r0 = default__.safeDivisionInt((0) - ((d_2461_v81_).cardinality), (p0) + (len(_dafny.SeqWithoutIsStrInference([p0]))))
        return r0

