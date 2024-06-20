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
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(555, 992):
                d_0_v0_: int = compr_0_
                if ((555) <= (d_0_v0_)) and ((d_0_v0_) < (992)):
                    coll0_[(d_0_v0_) * (-542)] = D2_DC6(_dafny.SeqWithoutIsStrInference([123, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1_i0_ in range(default__.abs(623))])), len(_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2_i1_ in range(default__.abs(161))]): True})), len(_dafny.SeqWithoutIsStrInference([False]))]))
            return _dafny.Map(coll0_)
        return (len(iife0_()
        )) >= (len((_dafny.Map({True: True})) | (_dafny.Map({False: False}))))

    @staticmethod
    def fm1(p0, p1, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: _dafny.Seq
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([(D2_DC6(_dafny.SeqWithoutIsStrInference([523, (_dafny.MultiSet([_dafny.CodePoint('p'), _dafny.CodePoint('l')])).cardinality]))).cf11, _dafny.SeqWithoutIsStrInference([156, -520])])).Elements:
                d_3_v0_: _dafny.Seq = compr_1_
                if (d_3_v0_) in (_dafny.SeqWithoutIsStrInference([(D2_DC6(_dafny.SeqWithoutIsStrInference([523, (_dafny.MultiSet([_dafny.CodePoint('p'), _dafny.CodePoint('l')])).cardinality]))).cf11, _dafny.SeqWithoutIsStrInference([156, -520])])):
                    coll1_[d_3_v0_] = not(True)
            return _dafny.Map(coll1_)
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: _dafny.Seq
            for compr_2_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([742 for d_4_i0_ in range(default__.abs(615))]): 225})).keys.Elements:
                d_5_v1_: _dafny.Seq = compr_2_
                if (d_5_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([742 for d_4_i0_ in range(default__.abs(615))]): 225})):
                    coll2_[d_5_v1_] = True
            return _dafny.Map(coll2_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([55]): True})) | (iife1_()
        )) | (iife2_()
        )

    @staticmethod
    def fm2(p0, p1, globalState):
        if (not(not(False))) or (True):
            def iife3_():
                coll3_ = _dafny.Set()
                compr_3_: str
                for compr_3_ in (_dafny.Map({_dafny.CodePoint('v'): True})).keys.Elements:
                    d_6_v0_: str = compr_3_
                    if (d_6_v0_) in (_dafny.Map({_dafny.CodePoint('v'): True})):
                        coll3_ = coll3_.union(_dafny.Set([d_6_v0_]))
                return _dafny.Set(coll3_)
            return (_dafny.MultiSet([_dafny.Set({_dafny.CodePoint('t'), _dafny.CodePoint('c'), _dafny.CodePoint('h')}), iife3_()
            ])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Set({_dafny.CodePoint('e'), _dafny.CodePoint('n')})])))
        elif True:
            def iife4_():
                coll4_ = _dafny.Set()
                compr_4_: str
                for compr_4_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_7_i0_ in range(default__.abs(953))]))).Elements:
                    d_8_v1_: str = compr_4_
                    if (d_8_v1_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_7_i0_ in range(default__.abs(953))]))):
                        coll4_ = coll4_.union(_dafny.Set([d_8_v1_]))
                return _dafny.Set(coll4_)
            def iife5_():
                coll5_ = _dafny.Set()
                compr_5_: str
                for compr_5_ in (_dafny.Map({_dafny.CodePoint('b'): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iasqoksj"))})).keys.Elements:
                    d_9_v2_: str = compr_5_
                    if (d_9_v2_) in (_dafny.Map({_dafny.CodePoint('b'): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iasqoksj"))})):
                        coll5_ = coll5_.union(_dafny.Set([d_9_v2_]))
                return _dafny.Set(coll5_)
            return _dafny.MultiSet([iife4_()
            , _dafny.Set({_dafny.CodePoint('y'), _dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('q'), _dafny.CodePoint('n')}), iife5_()
            ])

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        def iife6_():
            coll6_ = _dafny.Set()
            compr_6_: str
            for compr_6_ in (_dafny.Map({_dafny.CodePoint('i'): _dafny.CodePoint('q')})).keys.Elements:
                d_10_v0_: str = compr_6_
                if (d_10_v0_) in (_dafny.Map({_dafny.CodePoint('i'): _dafny.CodePoint('q')})):
                    coll6_ = coll6_.union(_dafny.Set([d_10_v0_]))
            return _dafny.Set(coll6_)
        return (_dafny.SeqWithoutIsStrInference([iife6_()
 for d_11_i0_ in range(default__.abs(183))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({_dafny.CodePoint('o')}) for d_12_i1_ in range(default__.abs(872))]))

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        return _dafny.CodePoint('o')

    @staticmethod
    def fm13(p0, p1, p2, globalState):
        return len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anbp"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qmbavh")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))))

    @staticmethod
    def fm16(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, False, True}))]) if not(False) else _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({True: (0) - (-169)}))) for d_13_i0_ in range(default__.abs(148))]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, not(not(True))])), (0) - ((0) - (-274)), len(_dafny.Map({525: 504})), -521])) + (_dafny.SeqWithoutIsStrInference([174, len(_dafny.Map({(0) - (len(_dafny.Set({False, False}))): _dafny.CodePoint('p')}))])))

    @staticmethod
    def fm17(p0, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jogx"))

    @staticmethod
    def fm18(p0, p1, p2, p3, globalState):
        return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): (_dafny.MultiSet([True, True])).cardinality})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxblvklw"))): 135}))

    @staticmethod
    def fm19(globalState):
        return D1_DC2((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvymmqds"))): len(_dafny.Map({not(False): not(True)}))})) | (_dafny.Map({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ygsxci")))])).cardinality: len(_dafny.SeqWithoutIsStrInference([True, False]))})))

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({not(not(False)): True}))) | ((_dafny.Map({not(True): True})) | (_dafny.Map({True: not(not(False))})))

    @staticmethod
    def fm21(p0, globalState):
        def iife7_():
            coll7_ = _dafny.Set()
            compr_7_: int
            for compr_7_ in (_dafny.SeqWithoutIsStrInference([-371])).Elements:
                d_14_v0_: int = compr_7_
                if (d_14_v0_) in (_dafny.SeqWithoutIsStrInference([-371])):
                    coll7_ = coll7_.union(_dafny.Set([(d_14_v0_) + (len(_dafny.Map({(0) - ((_dafny.MultiSet([792])).cardinality): len(_dafny.Map({True: 504}))})))]))
            return _dafny.Set(coll7_)
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkuhorvqx")))])).Elements:
                d_15_v1_: int = compr_8_
                if (d_15_v1_) in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkuhorvqx")))])):
                    coll8_[(d_15_v1_) - (-619)] = D9_DC23(_dafny.MultiSet([True]))
            return _dafny.Map(coll8_)
        return _dafny.SeqWithoutIsStrInference([(iife7_()
) | (_dafny.Set({101, len(_dafny.SeqWithoutIsStrInference([len(iife8_()
)]))})) for d_16_i0_ in range(default__.abs(105))])

    @staticmethod
    def fm22(p0, p1, globalState):
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(786, 626):
                d_18_v0_: int = compr_9_
                if ((786) <= (d_18_v0_)) and ((d_18_v0_) < (626)):
                    coll9_ = coll9_.union(_dafny.Set([default__.safeDivisionInt(d_18_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olteeivdc"))))]))
            return _dafny.Set(coll9_)
        if (_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_17_i0_ in range(default__.abs(-459))]), _dafny.SeqWithoutIsStrInference([len(iife9_()
        ), len(_dafny.SeqWithoutIsStrInference([386 for d_19_i1_ in range(default__.abs(-830))]))])})).ispropersubset(_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True}))]), _dafny.SeqWithoutIsStrInference([885, len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxe"))})), -13]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({12})), 170, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjay")))]), _dafny.SeqWithoutIsStrInference([890])})):
            if True:
                return _dafny.CodePoint('d')
            elif True:
                return _dafny.CodePoint('q')
        elif True:
            return _dafny.CodePoint('d')
        elif True:
            return _dafny.CodePoint('o')

    @staticmethod
    def fm23(p0, p1, globalState):
        if True:
            return D2_DC7(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h')]), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eurek"))), len(_dafny.Set({_dafny.CodePoint('u'), _dafny.CodePoint('d')})), True)
        elif True:
            return D2_DC7(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v')]), len(_dafny.SeqWithoutIsStrInference([not(False)])), 798, not(False))

    @staticmethod
    def fm24(p0, p1, globalState):
        return ((_dafny.Map({False: 102})) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('w')]))}))) | ((_dafny.Map({not(False): 47})) | (_dafny.Map({not(False): 686})))

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        return (_dafny.Set({False, (D9_DC25(True, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, False]))])), True)).cf41, False, True})) - (_dafny.Set({True}))

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        def iife10_():
            coll10_ = _dafny.Set()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(782, -958):
                d_20_v0_: int = compr_10_
                if ((782) <= (d_20_v0_)) and ((d_20_v0_) < (-958)):
                    coll10_ = coll10_.union(_dafny.Set([(d_20_v0_) - (86)]))
            return _dafny.Set(coll10_)
        return ((iife10_()
        ).intersection(_dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([411, 918, 498, 372, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([75 for d_21_i0_ in range(default__.abs(656))]))).cardinality]))).cardinality, 478}))).intersection(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dslxbs")))}))

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')])

    @staticmethod
    def fm31(globalState):
        return default__.safeModuloInt((len(_dafny.SeqWithoutIsStrInference([True, True]))) * (742), 652)

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjjqasxw"))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_22_i0_ in range(default__.abs(-236))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idsrwnse"))))

    @staticmethod
    def fm33(p0, p1, p2, p3, globalState):
        source0_ = D2_DC7(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('a'), _dafny.CodePoint('y')]), (0) - (-221), len(_dafny.SeqWithoutIsStrInference([625])), False)
        if source0_.is_DC7:
            d_23___mcc_h0_ = source0_.cf12
            d_24___mcc_h1_ = source0_.cf13
            d_25___mcc_h2_ = source0_.cf14
            d_26___mcc_h3_ = source0_.cf15
            d_27_cf15_ = d_26___mcc_h3_
            d_28_cf14_ = d_25___mcc_h2_
            d_29_cf13_ = d_24___mcc_h1_
            d_30_cf12_ = d_23___mcc_h0_
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: int
                for compr_11_ in _dafny.IntegerRange(109, 892):
                    d_31_v0_: int = compr_11_
                    if ((109) <= (d_31_v0_)) and ((d_31_v0_) < (892)):
                        coll11_[default__.safeDivisionInt(d_31_v0_, d_28_cf14_)] = d_27_cf15_
                return _dafny.Map(coll11_)
            return (_dafny.Map({d_29_cf13_: d_27_cf15_})) | (iife11_()
            )
        elif source0_.is_DC8:
            d_32___mcc_h4_ = source0_.cf16
            d_33___mcc_h5_ = source0_.cf17
            d_34___mcc_h6_ = source0_.cf18
            d_35_cf18_ = d_34___mcc_h6_
            d_36_cf17_ = d_33___mcc_h5_
            d_37_cf16_ = d_32___mcc_h4_
            return (_dafny.Map({(0) - (len(_dafny.Map({d_36_cf17_: 741}))): True})) | (_dafny.Map({142: d_36_cf17_}))
        elif True:
            d_38___mcc_h7_ = source0_.cf11
            d_39_cf11_ = d_38___mcc_h7_
            return _dafny.Map({138: True})

    @staticmethod
    def fm34(globalState):
        source1_ = (D11_DC30(True) if True else D11_DC30(True))
        if source1_.is_DC28:
            d_40___mcc_h0_ = source1_.cf44
            d_41_cf44_ = d_40___mcc_h0_
            return _dafny.CodePoint('y')
        elif source1_.is_DC29:
            d_42___mcc_h1_ = source1_.cf45
            d_43___mcc_h2_ = source1_.cf46
            d_44___mcc_h3_ = source1_.cf47
            d_45___mcc_h4_ = source1_.cf48
            d_46_cf48_ = d_45___mcc_h4_
            d_47_cf47_ = d_44___mcc_h3_
            d_48_cf46_ = d_43___mcc_h2_
            d_49_cf45_ = d_42___mcc_h1_
            return _dafny.CodePoint('a')
        elif source1_.is_DC30:
            d_50___mcc_h5_ = source1_.cf49
            d_51_cf49_ = d_50___mcc_h5_
            return _dafny.CodePoint('u')
        elif source1_.is_DC27:
            d_52___mcc_h6_ = source1_.cf43
            d_53_cf43_ = d_52___mcc_h6_
            return _dafny.CodePoint('p')
        elif True:
            d_54___mcc_h7_ = source1_.cf50
            d_55_cf50_ = d_54___mcc_h7_
            if False:
                return _dafny.CodePoint('a')
            elif True:
                return _dafny.CodePoint('v')

    @staticmethod
    def fm35(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(0) - (-181) for d_56_i0_ in range(default__.abs(568))])) + (_dafny.SeqWithoutIsStrInference([-182 for d_57_i1_ in range(default__.abs(-152))]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eolvsmea"))]))])) + (_dafny.SeqWithoutIsStrInference([673, 562])))

    @staticmethod
    def fm36(globalState):
        return ((_dafny.MultiSet([True])) - (_dafny.MultiSet([True]))).intersection((_dafny.MultiSet([not(True)])).intersection(_dafny.MultiSet([False])))

    @staticmethod
    def fm37(p0, globalState):
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: int
            for compr_12_ in _dafny.IntegerRange(890, 345):
                d_58_v0_: int = compr_12_
                if ((890) <= (d_58_v0_)) and ((d_58_v0_) < (345)):
                    coll12_[default__.safeModuloInt(d_58_v0_, 137)] = _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmgrm")))})
            return _dafny.Map(coll12_)
        return iife12_()
        

    @staticmethod
    def fm38(p0, globalState):
        source2_ = D5_DC15(D5_DC12(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([-443]))})))
        if source2_.is_DC13:
            def iife13_():
                def iife15_():
                    coll15_ = _dafny.Map()
                    compr_15_: _dafny.Seq
                    for compr_15_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, True])])).Elements:
                        d_59_v1_: _dafny.Seq = compr_15_
                        if (d_59_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, True])])):
                            coll15_[d_59_v1_] = -606
                    return _dafny.Map(coll15_)
                coll13_ = _dafny.Map()
                def iife14_():
                    coll14_ = _dafny.Map()
                    compr_14_: _dafny.Seq
                    for compr_14_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, True])])).Elements:
                        d_59_v1_: _dafny.Seq = compr_14_
                        if (d_59_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, True])])):
                            coll14_[d_59_v1_] = -606
                    return _dafny.Map(coll14_)
                compr_13_: _dafny.Seq
                for compr_13_ in (iife14_()
                ).keys.Elements:
                    d_60_v0_: _dafny.Seq = compr_13_
                    if (d_60_v0_) in (iife15_()
                    ):
                        coll13_[d_60_v0_] = _dafny.Map({False: len(_dafny.Map({478: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_61_i0_ in range(default__.abs(150))]))}))})
                return _dafny.Map(coll13_)
            return (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): _dafny.Map({(D11_DC28(False)).cf44: len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(D9_DC25(True, len(_dafny.Set({False})), False)).cf40])), -792]))}))})})) | (iife13_()
            )
        elif source2_.is_DC14:
            d_62___mcc_h0_ = source2_.cf23
            d_63_cf23_ = d_62___mcc_h0_
            return _dafny.Map({_dafny.SeqWithoutIsStrInference([d_63_cf23_, d_63_cf23_, d_63_cf23_]): _dafny.Map({d_63_cf23_: 683})})
        elif source2_.is_DC12:
            d_64___mcc_h1_ = source2_.cf22
            d_65_cf22_ = d_64___mcc_h1_
            return _dafny.Map({_dafny.SeqWithoutIsStrInference([False, False]): d_65_cf22_})
        elif True:
            d_66___mcc_h2_ = source2_.cf24
            d_67_cf24_ = d_66___mcc_h2_
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: _dafny.Seq
                for compr_16_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([True, False]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uevfjb"))})).keys.Elements:
                    d_68_v2_: _dafny.Seq = compr_16_
                    if (d_68_v2_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([True, False]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uevfjb"))})):
                        coll16_[d_68_v2_] = _dafny.Map({False: (0) - (-902)})
                return _dafny.Map(coll16_)
            return (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): _dafny.Map({not(False): 164})})) | (iife16_()
            )

    @staticmethod
    def fm39(p0, globalState):
        source3_ = D11_DC28(True)
        if source3_.is_DC28:
            d_69___mcc_h0_ = source3_.cf44
            d_70_cf44_ = d_69___mcc_h0_
            return (_dafny.Map({320: (0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_71_i1_ in range(default__.abs(775))])) for d_72_i0_ in range(default__.abs(189))]))).cardinality)})) | (_dafny.Map({402: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vavsi")))])).cardinality}))
        elif source3_.is_DC29:
            d_73___mcc_h1_ = source3_.cf45
            d_74___mcc_h2_ = source3_.cf46
            d_75___mcc_h3_ = source3_.cf47
            d_76___mcc_h4_ = source3_.cf48
            d_77_cf48_ = d_76___mcc_h4_
            d_78_cf47_ = d_75___mcc_h3_
            d_79_cf46_ = d_74___mcc_h2_
            d_80_cf45_ = d_73___mcc_h1_
            return _dafny.Map({d_78_cf47_: d_78_cf47_})
        elif source3_.is_DC30:
            d_81___mcc_h5_ = source3_.cf49
            d_82_cf49_ = d_81___mcc_h5_
            return _dafny.Map({109: len(_dafny.Map({d_82_cf49_: False}))})
        elif source3_.is_DC27:
            d_83___mcc_h6_ = source3_.cf43
            d_84_cf43_ = d_83___mcc_h6_
            return (_dafny.Map({80: len(_dafny.SeqWithoutIsStrInference([408]))})) | (_dafny.Map({len(_dafny.Set({False, False})): 172}))
        elif True:
            d_85___mcc_h7_ = source3_.cf50
            d_86_cf50_ = d_85___mcc_h7_
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in (_dafny.MultiSet([698])).Elements:
                    d_87_v0_: int = compr_17_
                    if (d_87_v0_) in (_dafny.MultiSet([698])):
                        coll17_[default__.safeDivisionInt(d_87_v0_, -374)] = 531
                return _dafny.Map(coll17_)
            return iife17_()
            

    @staticmethod
    def fm40(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aynax"))))])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (-406) for d_88_i0_ in range(default__.abs(387))])))

    @staticmethod
    def fm41(p0, globalState):
        source4_ = D11_DC29(-157, 677, (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.SeqWithoutIsStrInference([-188 for d_89_i0_ in range(default__.abs(48))]), _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality]))) for d_90_i1_ in range(default__.abs(953))])})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsx")))]))])).cardinality, 681)
        if source4_.is_DC28:
            d_91___mcc_h0_ = source4_.cf44
            d_92_cf44_ = d_91___mcc_h0_
            return D4_DC11(d_92_cf44_)
        elif source4_.is_DC29:
            d_93___mcc_h1_ = source4_.cf45
            d_94___mcc_h2_ = source4_.cf46
            d_95___mcc_h3_ = source4_.cf47
            d_96___mcc_h4_ = source4_.cf48
            d_97_cf48_ = d_96___mcc_h4_
            d_98_cf47_ = d_95___mcc_h3_
            d_99_cf46_ = d_94___mcc_h2_
            d_100_cf45_ = d_93___mcc_h1_
            return D4_DC11(False)
        elif source4_.is_DC30:
            d_101___mcc_h5_ = source4_.cf49
            d_102_cf49_ = d_101___mcc_h5_
            return D4_DC11(d_102_cf49_)
        elif source4_.is_DC27:
            d_103___mcc_h6_ = source4_.cf43
            d_104_cf43_ = d_103___mcc_h6_
            return D4_DC11(False)
        elif True:
            d_105___mcc_h7_ = source4_.cf50
            d_106_cf50_ = d_105___mcc_h7_
            return D4_DC11(False)

    @staticmethod
    def fm42(globalState):
        return D0_DC1((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgdv")))])).isdisjoint(_dafny.MultiSet([-820, len(_dafny.SeqWithoutIsStrInference([not(True)])), 103, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aheg"))), len(_dafny.Map({True: len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qehktwkga")))}))}))])), (-193) + (len(_dafny.SeqWithoutIsStrInference([False, not(not(not(True)))]))))

    @staticmethod
    def fm43(globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yaynwj")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_107_i0_ in range(default__.abs(188))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrbinxrmv"))])) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fmj"))])) - ((D12_DC32(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtvdgkj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxyl"))])))).cf51))

    @staticmethod
    def fm44(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([D1_DC2((D1_DC2(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True]) for d_108_i1_ in range(default__.abs(27))])): 894}))).cf3) for d_109_i0_ in range(default__.abs(855))])

    @staticmethod
    def fm45(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([False, False])) + (_dafny.SeqWithoutIsStrInference([False, False]))

    @staticmethod
    def fm46(p0, globalState):
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(821, 144):
                d_110_v0_: int = compr_18_
                if ((821) <= (d_110_v0_)) and ((d_110_v0_) < (144)):
                    coll18_[(d_110_v0_) + ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dav")))))] = -163
            return _dafny.Map(coll18_)
        def iife19_():
            coll19_ = _dafny.Map()
            compr_19_: int
            for compr_19_ in _dafny.IntegerRange(767, 440):
                d_111_v1_: int = compr_19_
                if ((767) <= (d_111_v1_)) and ((d_111_v1_) < (440)):
                    coll19_[default__.safeDivisionInt(d_111_v1_, len(_dafny.Map({len((D2_DC6(_dafny.SeqWithoutIsStrInference([185, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvwjkcv"))), 685]))).cf11): len(_dafny.Set({152, len(_dafny.Set({not(True)}))}))})))] = (_dafny.MultiSet([70, 800])).cardinality
            return _dafny.Map(coll19_)
        return _dafny.Set({iife18_()
        , (_dafny.Map({(0) - (len(iife19_()
        )): 489})) | (_dafny.Map({955: 299})), (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([383])): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yi")))})) | (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_112_i0_ in range(default__.abs(203))]))): 120}))})

    @staticmethod
    def fm47(p0, p1, globalState):
        return _dafny.Set({_dafny.CodePoint('f'), (_dafny.CodePoint('i') if True else _dafny.CodePoint('l')), _dafny.CodePoint('o')})

    @staticmethod
    def fm48(p0, p1, p2, p3, globalState):
        return D2_DC6(_dafny.SeqWithoutIsStrInference([794, -102]))

    @staticmethod
    def m0(p0, globalState):
        r0: int = int(0)
        d_113_v0_: int
        d_113_v0_ = -3
        d_114_v1_: _dafny.Set
        d_114_v1_ = _dafny.Set({d_113_v0_})
        d_115_v2_: bool
        d_115_v2_ = False
        d_116_v3_: _dafny.Seq
        d_116_v3_ = _dafny.SeqWithoutIsStrInference([True, d_115_v2_])
        d_117_v4_: C2
        nw0_ = C2()
        nw0_.ctor__(d_115_v2_)
        d_117_v4_ = nw0_
        d_118_v5_: _dafny.MultiSet
        d_118_v5_ = _dafny.MultiSet([d_117_v4_, d_117_v4_])
        d_119_v6_: C4
        nw1_ = C4()
        nw1_.ctor__(d_115_v2_, False)
        d_119_v6_ = nw1_
        d_120_v7_: _dafny.Set
        d_120_v7_ = _dafny.Set({d_119_v6_})
        d_121_v8_: _dafny.Array
        nw2_ = _dafny.Array(int(0), 23)
        d_121_v8_ = nw2_
        d_122_v9_: _dafny.Map
        d_122_v9_ = _dafny.Map({d_121_v8_: (d_119_v6_).f16})
        d_123_v10_: _dafny.Map
        d_123_v10_ = d_122_v9_
        d_124_v11_: _dafny.Map
        d_124_v11_ = _dafny.Map({(d_119_v6_).f16: d_115_v2_})
        d_125_v12_: D4
        d_125_v12_ = D4_DC11((d_119_v6_).f16)
        d_126_v13_: _dafny.Map
        d_126_v13_ = _dafny.Map({d_113_v0_: d_113_v0_})
        d_127_v14_: _dafny.Seq
        d_127_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wm"))
        d_128_v15_: str
        d_128_v15_ = _dafny.CodePoint('v')
        d_129_v16_: _dafny.MultiSet
        d_129_v16_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yyl")), (d_127_v14_).set(default__.safeIndex(d_113_v0_, len(d_127_v14_)), d_128_v15_)])
        d_130_v17_: _dafny.MultiSet
        d_130_v17_ = _dafny.MultiSet([d_116_v3_])
        d_131_v18_: _dafny.Array
        nw3_ = _dafny.Array(None, 29)
        nw3_[int(0)] = (d_113_v0_) < (len(d_114_v1_))
        nw3_[int(1)] = (d_116_v3_)[default__.safeIndex((d_118_v5_).cardinality, len(d_116_v3_))]
        nw3_[int(2)] = (d_120_v7_).ispropersubset(d_120_v7_)
        nw3_[int(3)] = d_115_v2_
        nw3_[int(4)] = not((d_119_v6_).f16)
        nw3_[int(5)] = (d_119_v6_).f16
        nw3_[int(6)] = (d_119_v6_).f16
        nw3_[int(7)] = d_115_v2_
        nw3_[int(8)] = True
        nw3_[int(9)] = d_115_v2_
        nw3_[int(10)] = d_115_v2_
        nw3_[int(11)] = d_115_v2_
        nw3_[int(12)] = (d_119_v6_).f16
        nw3_[int(13)] = ((_dafny.Map({d_121_v8_: d_115_v2_}))) == ((d_122_v9_).set(d_121_v8_, (d_119_v6_).f16))
        nw3_[int(14)] = ((d_122_v9_)[d_121_v8_] if (d_121_v8_) in (d_122_v9_) else d_115_v2_)
        nw3_[int(15)] = (d_115_v2_) == (d_115_v2_)
        nw3_[int(16)] = (default__.fm25(d_114_v1_, d_124_v11_, d_113_v0_, d_115_v2_, globalState)).isdisjoint(default__.fm25(d_114_v1_, default__.fm20(d_115_v2_, d_113_v0_, (d_119_v6_).f16, globalState), d_113_v0_, (d_119_v6_).f16, globalState))
        nw3_[int(17)] = d_115_v2_
        nw3_[int(18)] = (d_119_v6_).f16
        nw3_[int(19)] = (d_119_v6_).f16
        nw3_[int(20)] = False
        nw3_[int(21)] = (-421) >= (d_113_v0_)
        nw3_[int(22)] = (d_115_v2_ if (d_119_v6_).f16 else (d_119_v6_).f16)
        nw3_[int(23)] = not((d_125_v12_).cf21)
        nw3_[int(24)] = default__.fm0(d_126_v13_, d_113_v0_, globalState)
        nw3_[int(25)] = (d_129_v16_).issubset(_dafny.MultiSet([d_127_v14_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_132_i0_ in range(default__.abs(139))]), d_127_v14_]))
        nw3_[int(26)] = (not(False)) or (d_115_v2_)
        nw3_[int(27)] = not((d_115_v2_ if (d_119_v6_).f16 else (d_119_v6_).f16))
        nw3_[int(28)] = (d_113_v0_) <= (((d_130_v17_)[d_116_v3_] if (d_116_v3_) in (d_130_v17_) else d_113_v0_))
        d_131_v18_ = nw3_
        d_131_v18_ = d_131_v18_
        d_133_v19_: _dafny.Seq
        d_133_v19_ = _dafny.SeqWithoutIsStrInference([d_116_v3_])
        source5_ = (d_133_v19_)[default__.safeIndex(d_113_v0_, len(d_133_v19_))]
        d_134___mcc_h0_ = source5_
        d_135_cf34_ = d_134___mcc_h0_
        r0 = d_113_v0_
        d_115_v2_ = (d_119_v6_).f16
        d_136_v20_: _dafny.Array
        nw4_ = _dafny.Array(_dafny.MultiSet({}), 25)
        d_136_v20_ = nw4_
        d_137_v21_: _dafny.MultiSet
        d_137_v21_ = _dafny.MultiSet([d_128_v15_])
        index0_ = default__.safeIndex(375, (d_136_v20_).length(0))
        (d_136_v20_)[index0_] = (d_137_v21_).intersection(d_137_v21_)
        index1_ = default__.safeIndex(375, (d_136_v20_).length(0))
        rhs0_ = (d_137_v21_) | (_dafny.MultiSet([d_128_v15_, _dafny.CodePoint('s')]))
        rhs1_ = not(d_115_v2_)
        lhs0_ = d_136_v20_
        lhs1_ = default__.safeIndex(375, (d_136_v20_).length(0))
        lhs0_[lhs1_] = rhs0_
        d_115_v2_ = rhs1_
        d_115_v2_ = ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rokcgkova"))))) <= (d_113_v0_)
        d_138_v22_: D1
        d_138_v22_ = D1_DC3(104)
        d_139_i1_: int
        d_139_i1_ = 0
        with _dafny.label("0"):
            pat_let_tv0_ = d_119_v6_
            pat_let_tv1_ = d_119_v6_
            pat_let_tv2_ = d_119_v6_
            pat_let_tv3_ = d_115_v2_
            pat_let_tv4_ = d_119_v6_
            pat_let_tv5_ = d_113_v0_
            def lambda0_(source6_):
                if source6_.is_DC3:
                    d_144___mcc_h1_ = source6_.cf4
                    d_145_cf4_ = d_144___mcc_h1_
                    return True
                elif source6_.is_DC4:
                    d_146___mcc_h2_ = source6_.cf5
                    d_147___mcc_h3_ = source6_.cf6
                    d_148___mcc_h4_ = source6_.cf7
                    d_149_cf7_ = d_148___mcc_h4_
                    d_150_cf6_ = d_147___mcc_h3_
                    d_151_cf5_ = d_146___mcc_h2_
                    return (_dafny.Set({(pat_let_tv0_).f16})).issubset(_dafny.Set({(pat_let_tv1_).f16}))
                elif source6_.is_DC5:
                    d_152___mcc_h5_ = source6_.cf8
                    d_153___mcc_h6_ = source6_.cf9
                    d_154___mcc_h7_ = source6_.cf10
                    d_155_cf10_ = d_154___mcc_h7_
                    d_156_cf9_ = d_153___mcc_h6_
                    d_157_cf8_ = d_152___mcc_h5_
                    return (pat_let_tv2_).f16
                elif True:
                    d_158___mcc_h8_ = source6_.cf3
                    d_159_cf3_ = d_158___mcc_h8_
                    return (len(_dafny.Set({pat_let_tv3_, (pat_let_tv4_).f16}))) == (pat_let_tv5_)

            while lambda0_((d_138_v22_ if d_115_v2_ else d_138_v22_)):
                with _dafny.c_label("0"):
                    if (d_139_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_139_i1_ = (d_139_i1_) + (1)
                    index2_ = default__.safeIndex(275, (d_121_v8_).length(0))
                    (d_121_v8_)[index2_] = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nntti"))) + (d_127_v14_)).set(default__.safeIndex(d_113_v0_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nntti"))) + (d_127_v14_))), d_128_v15_))
                    index3_ = default__.safeIndex(275, (d_121_v8_).length(0))
                    (d_121_v8_)[index3_] = (d_113_v0_) + (d_113_v0_)
                    d_140_v23_: C4
                    nw5_ = C4()
                    nw5_.ctor__(((d_121_v8_)[default__.safeIndex(275, (d_121_v8_).length(0))]) != ((0) - ((d_121_v8_)[default__.safeIndex(275, (d_121_v8_).length(0))])), (d_114_v1_).issubset((D11_DC27(d_114_v1_)).cf43))
                    d_140_v23_ = nw5_
                    d_141_v24_: _dafny.Map
                    d_141_v24_ = _dafny.Map({(0) - (d_113_v0_): d_131_v18_})
                    d_142_v25_: D1
                    d_142_v25_ = D1_DC4(d_131_v18_, len(d_127_v14_), d_128_v15_)
                    d_131_v18_ = ((d_141_v24_)[d_113_v0_] if (d_113_v0_) in (d_141_v24_) else (d_142_v25_).cf5)
                    d_143_v26_: _dafny.Array
                    nw6_ = _dafny.Array(_dafny.Array(None, 0), 13)
                    d_143_v26_ = nw6_
                    index4_ = default__.safeIndex(475, (d_143_v26_).length(0))
                    (d_143_v26_)[index4_] = d_131_v18_
                    index5_ = default__.safeIndex(475, (d_143_v26_).length(0))
                    (d_143_v26_)[index5_] = d_131_v18_
                    pass
            pass
        d_160_v27_: _dafny.Set
        d_160_v27_ = _dafny.Set({d_124_v11_, d_124_v11_, d_124_v11_})
        d_115_v2_ = (d_160_v27_) == (d_160_v27_)
        d_161_v28_: _dafny.Array
        nw7_ = _dafny.Array(D7.default()(), 22)
        d_161_v28_ = nw7_
        d_162_v29_: D7
        d_162_v29_ = D7_DC20(-722, d_128_v15_)
        index6_ = default__.safeIndex(557, (d_161_v28_).length(0))
        (d_161_v28_)[index6_] = d_162_v29_
        index7_ = default__.safeIndex(557, (d_161_v28_).length(0))
        (d_161_v28_)[index7_] = d_162_v29_
        d_113_v0_ = d_113_v0_
        r0 = d_113_v0_
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_163_v0_: bool
        d_163_v0_ = False
        d_164_v1_: _dafny.Seq
        d_164_v1_ = _dafny.SeqWithoutIsStrInference([d_163_v0_, not(d_163_v0_), d_163_v0_, not(d_163_v0_)])
        d_165_v2_: int
        d_165_v2_ = -199
        d_166_v3_: _dafny.Array
        nw8_ = _dafny.Array(None, 27)
        nw8_[int(0)] = d_163_v0_
        nw8_[int(1)] = d_163_v0_
        nw8_[int(2)] = d_163_v0_
        nw8_[int(3)] = d_163_v0_
        nw8_[int(4)] = d_163_v0_
        nw8_[int(5)] = d_163_v0_
        nw8_[int(6)] = d_163_v0_
        nw8_[int(7)] = d_163_v0_
        nw8_[int(8)] = True
        nw8_[int(9)] = d_163_v0_
        nw8_[int(10)] = d_163_v0_
        nw8_[int(11)] = d_163_v0_
        nw8_[int(12)] = not(d_163_v0_)
        nw8_[int(13)] = True
        nw8_[int(14)] = d_163_v0_
        nw8_[int(15)] = d_163_v0_
        nw8_[int(16)] = d_163_v0_
        nw8_[int(17)] = d_163_v0_
        nw8_[int(18)] = d_163_v0_
        nw8_[int(19)] = d_163_v0_
        nw8_[int(20)] = d_163_v0_
        nw8_[int(21)] = d_163_v0_
        nw8_[int(22)] = d_163_v0_
        nw8_[int(23)] = (d_164_v1_)[default__.safeIndex(d_165_v2_, len(d_164_v1_))]
        nw8_[int(24)] = d_163_v0_
        nw8_[int(25)] = d_163_v0_
        nw8_[int(26)] = d_163_v0_
        d_166_v3_ = nw8_
        d_167_v4_: D0
        d_167_v4_ = D0_DC0(False)
        d_168_v5_: _dafny.Set
        d_168_v5_ = _dafny.Set({d_164_v1_, (_dafny.SeqWithoutIsStrInference([d_163_v0_, (d_167_v4_).cf0, not(False), d_163_v0_, d_163_v0_])).set(default__.safeIndex(d_165_v2_, len(_dafny.SeqWithoutIsStrInference([d_163_v0_, (d_167_v4_).cf0, not(False), d_163_v0_, d_163_v0_]))), False), d_164_v1_})
        d_169_v6_: _dafny.Seq
        d_169_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nftgk"))
        d_170_globalState_: GlobalState
        nw9_ = GlobalState()
        nw9_.ctor__(-586, d_166_v3_, True, False, d_164_v1_, (d_168_v5_) | (d_168_v5_), False, _dafny.MultiSet([d_165_v2_]), True, 721, d_166_v3_, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_171_i0_ in range(default__.abs(777))])) + (d_169_v6_), 604)
        d_170_globalState_ = nw9_
        d_172_v7_: D0
        d_172_v7_ = D0_DC1(d_163_v0_, default__.safeDivisionInt(d_165_v2_, d_165_v2_))
        index8_ = default__.safeIndex(243, (d_166_v3_).length(0))
        (d_166_v3_)[index8_] = (d_169_v6_) < (d_169_v6_)
        d_173_v8_: _dafny.Map
        d_173_v8_ = _dafny.Map({(0) - (d_165_v2_): 833})
        d_174_v9_: _dafny.MultiSet
        d_174_v9_ = _dafny.MultiSet([d_165_v2_])
        d_175_v10_: str
        d_175_v10_ = _dafny.CodePoint('d')
        index9_ = default__.safeIndex(243, (d_166_v3_).length(0))
        rhs2_ = d_172_v7_
        rhs3_ = default__.fm0(d_173_v8_, (0) - ((d_174_v9_).cardinality), d_170_globalState_)
        rhs4_ = (d_175_v10_) in ((d_169_v6_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qtgkj"))))
        rhs5_ = d_163_v0_
        rhs6_ = d_166_v3_
        lhs2_ = d_166_v3_
        lhs3_ = default__.safeIndex(243, (d_166_v3_).length(0))
        d_172_v7_ = rhs2_
        lhs2_[lhs3_] = rhs3_
        d_163_v0_ = rhs4_
        d_163_v0_ = rhs5_
        d_166_v3_ = rhs6_
        d_176_i1_: int
        d_176_i1_ = 0
        with _dafny.label("1"):
            while False:
                with _dafny.c_label("1"):
                    if (d_176_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_176_i1_ = (d_176_i1_) + (1)
                    d_177_v11_: _dafny.Array
                    def lambda1_(d_178_v1_):
                        def lambda2_(d_179_i2_):
                            return d_178_v1_

                        return lambda2_

                    init0_ = lambda1_(d_164_v1_)
                    nw10_ = _dafny.Array(None, 13)
                    for i0_0_ in range(nw10_.length(0)):
                        nw10_[i0_0_] = init0_(i0_0_)
                    d_177_v11_ = nw10_
                    index10_ = default__.safeIndex(17, (d_177_v11_).length(0))
                    (d_177_v11_)[index10_] = d_164_v1_
                    index11_ = default__.safeIndex(17, (d_177_v11_).length(0))
                    (d_177_v11_)[index11_] = d_164_v1_
                    d_180_v12_: _dafny.Array
                    nw11_ = _dafny.Array(_dafny.Set({}), 14)
                    d_180_v12_ = nw11_
                    d_181_v13_: _dafny.Map
                    d_181_v13_ = _dafny.Map({d_166_v3_: default__.fm0(d_173_v8_, d_165_v2_, d_170_globalState_)})
                    d_182_v14_: _dafny.Map
                    d_182_v14_ = _dafny.Map({d_163_v0_: ((d_181_v13_)[d_166_v3_] if (d_166_v3_) in (d_181_v13_) else True)})
                    d_183_v15_: _dafny.Array
                    nw12_ = _dafny.Array(int(0), 16)
                    d_183_v15_ = nw12_
                    d_184_v16_: _dafny.Set
                    d_184_v16_ = _dafny.Set({d_183_v15_, d_183_v15_})
                    d_185_v17_: _dafny.Map
                    d_185_v17_ = _dafny.Map({len(d_182_v14_): d_184_v16_})
                    index12_ = default__.safeIndex(995, (d_180_v12_).length(0))
                    (d_180_v12_)[index12_] = (((d_185_v17_)[-75] if (-75) in (d_185_v17_) else d_184_v16_)) | (d_184_v16_)
                    index13_ = default__.safeIndex(995, (d_180_v12_).length(0))
                    (d_180_v12_)[index13_] = _dafny.Set({d_183_v15_})
                    d_175_v10_ = d_175_v10_
                    d_186_v18_: _dafny.Map
                    d_186_v18_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wm")): (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]})
                    d_187_v19_: int
                    out0_: int
                    out0_ = default__.m0(d_186_v18_, d_170_globalState_)
                    d_187_v19_ = out0_
                    pass
            pass
        d_188_v20_: _dafny.Array
        def lambda3_(d_189_v2_):
            def lambda4_(d_190_i3_):
                return (d_190_i3_) + (d_189_v2_)

            return lambda4_

        init1_ = lambda3_(d_165_v2_)
        nw13_ = _dafny.Array(None, 27)
        for i0_1_ in range(nw13_.length(0)):
            nw13_[i0_1_] = init1_(i0_1_)
        d_188_v20_ = nw13_
        d_191_v21_: _dafny.Map
        d_191_v21_ = _dafny.Map({d_165_v2_: d_188_v20_})
        d_191_v21_ = d_191_v21_
        d_192_v22_: D1
        d_192_v22_ = D1_DC2(d_173_v8_)
        d_193_i4_: int
        d_193_i4_ = 0
        with _dafny.label("2"):
            while default__.fm0((d_192_v22_).cf3, d_165_v2_, d_170_globalState_):
                with _dafny.c_label("2"):
                    if (d_193_i4_) >= (100):
                        raise _dafny.Break("2")
                    d_193_i4_ = (d_193_i4_) + (1)
                    index14_ = default__.safeIndex(243, (d_166_v3_).length(0))
                    (d_166_v3_)[index14_] = default__.fm0(d_173_v8_, d_165_v2_, d_170_globalState_)
                    d_194_v23_: _dafny.Map
                    d_194_v23_ = _dafny.Map({d_165_v2_: not(d_163_v0_)})
                    d_195_v24_: _dafny.Seq
                    d_195_v24_ = _dafny.SeqWithoutIsStrInference([d_165_v2_])
                    d_196_v25_: _dafny.Map
                    d_196_v25_ = _dafny.Map({d_194_v23_: d_195_v24_})
                    d_197_v26_: _dafny.Set
                    d_197_v26_ = _dafny.Set({(d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))], d_163_v0_, d_163_v0_, not(False)})
                    d_198_v27_: _dafny.Map
                    d_198_v27_ = _dafny.Map({len(d_197_v26_): d_195_v24_})
                    d_199_v28_: _dafny.Map
                    d_199_v28_ = _dafny.Map({(d_196_v25_) | (_dafny.Map({d_194_v23_: ((d_198_v27_)[len(d_169_v6_)] if (len(d_169_v6_)) in (d_198_v27_) else d_195_v24_)})): default__.fm0(d_173_v8_, d_165_v2_, d_170_globalState_)})
                    d_199_v28_ = (d_199_v28_).set((d_196_v25_) | (_dafny.Map({d_194_v23_: d_195_v24_})), d_163_v0_)
                    index15_ = default__.safeIndex(484, (d_188_v20_).length(0))
                    def iife20_():
                        coll20_ = _dafny.Set()
                        compr_20_: int
                        for compr_20_ in _dafny.IntegerRange(448, 86):
                            d_200_v29_: int = compr_20_
                            if ((448) <= (d_200_v29_)) and ((d_200_v29_) < (86)):
                                coll20_ = coll20_.union(_dafny.Set([(d_200_v29_) * (d_165_v2_)]))
                        return _dafny.Set(coll20_)
                    (d_188_v20_)[index15_] = len(iife20_()
                    )
                    d_201_v30_: _dafny.Map
                    d_201_v30_ = _dafny.Map({d_163_v0_: len(d_195_v24_)})
                    index16_ = default__.safeIndex(484, (d_188_v20_).length(0))
                    index17_ = default__.safeIndex(243, (d_166_v3_).length(0))
                    rhs7_ = False
                    rhs8_ = d_165_v2_
                    rhs9_ = d_188_v20_
                    rhs10_ = (782) + (d_165_v2_)
                    rhs11_ = ((d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]) not in ((d_201_v30_) | (d_201_v30_))
                    lhs4_ = d_188_v20_
                    lhs5_ = default__.safeIndex(484, (d_188_v20_).length(0))
                    lhs6_ = d_166_v3_
                    lhs7_ = default__.safeIndex(243, (d_166_v3_).length(0))
                    d_163_v0_ = rhs7_
                    lhs4_[lhs5_] = rhs8_
                    d_188_v20_ = rhs9_
                    d_165_v2_ = rhs10_
                    lhs6_[lhs7_] = rhs11_
                    d_202_v31_: _dafny.Seq
                    d_202_v31_ = _dafny.SeqWithoutIsStrInference([d_166_v3_, d_166_v3_])
                    index18_ = default__.safeIndex(243, (d_166_v3_).length(0))
                    (d_166_v3_)[index18_] = (d_166_v3_) not in (d_202_v31_)
                    pass
            pass
        d_203_v32_: _dafny.Map
        d_203_v32_ = _dafny.Map({not((d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]): d_165_v2_})
        hi0_ = d_165_v2_
        for d_204_i5_ in range(len((d_203_v32_) | (d_203_v32_)), hi0_):
            d_205_v33_: _dafny.Array
            nw14_ = _dafny.Array(_dafny.Array(None, 0), 21)
            d_205_v33_ = nw14_
            d_206_v34_: _dafny.Seq
            d_206_v34_ = _dafny.SeqWithoutIsStrInference([d_165_v2_])
            d_207_v35_: _dafny.Map
            d_207_v35_ = _dafny.Map({(d_206_v34_) + (d_206_v34_): (True if d_163_v0_ else d_163_v0_)})
            d_208_v36_: _dafny.Map
            d_208_v36_ = _dafny.Map({d_188_v20_: d_165_v2_})
            index19_ = default__.safeIndex(243, (d_166_v3_).length(0))
            index20_ = default__.safeIndex(243, (d_166_v3_).length(0))
            rhs12_ = d_165_v2_
            rhs13_ = d_205_v33_
            rhs14_ = default__.fm1(False, d_172_v7_, d_170_globalState_)
            rhs15_ = ((d_208_v36_) != (d_208_v36_) if ((d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))] if (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))] else d_163_v0_) else True)
            rhs16_ = (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]
            lhs8_ = d_166_v3_
            lhs9_ = default__.safeIndex(243, (d_166_v3_).length(0))
            lhs10_ = d_166_v3_
            lhs11_ = default__.safeIndex(243, (d_166_v3_).length(0))
            d_165_v2_ = rhs12_
            d_205_v33_ = rhs13_
            d_207_v35_ = rhs14_
            lhs8_[lhs9_] = rhs15_
            lhs10_[lhs11_] = rhs16_
            d_165_v2_ = (0) - ((d_204_i5_) - (d_204_i5_))
            index21_ = default__.safeIndex(243, (d_166_v3_).length(0))
            (d_166_v3_)[index21_] = (d_165_v2_) != (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjidffht"))))
            d_163_v0_ = ((0) - (d_204_i5_)) == (len((d_169_v6_) + (_dafny.SeqWithoutIsStrInference([d_175_v10_ for d_209_i6_ in range(default__.abs(-249))]))))
        d_210_v37_: _dafny.Set
        d_210_v37_ = _dafny.Set({d_175_v10_})
        d_211_v38_: _dafny.MultiSet
        d_211_v38_ = _dafny.MultiSet([d_210_v37_])
        d_212_v39_: _dafny.MultiSet
        d_212_v39_ = _dafny.MultiSet([False, (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))], (d_164_v1_)[default__.safeIndex((d_174_v9_).cardinality, len(d_164_v1_))]])
        d_213_v41_: _dafny.Seq
        def iife21_():
            coll21_ = _dafny.Set()
            compr_21_: str
            for compr_21_ in (_dafny.Map({_dafny.CodePoint('p'): (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]})).keys.Elements:
                d_214_v40_: str = compr_21_
                if (d_214_v40_) in (_dafny.Map({_dafny.CodePoint('p'): (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]})):
                    coll21_ = coll21_.union(_dafny.Set([d_214_v40_]))
            return _dafny.Set(coll21_)
        d_213_v41_ = _dafny.SeqWithoutIsStrInference([d_210_v37_, iife21_()
        ])
        d_215_v42_: _dafny.Array
        nw15_ = _dafny.Array(None, 9)
        nw15_[int(0)] = d_211_v38_
        nw15_[int(1)] = d_211_v38_
        nw15_[int(2)] = default__.fm2(not(False), d_165_v2_, d_170_globalState_)
        nw15_[int(3)] = _dafny.MultiSet(default__.fm3((d_212_v39_).cardinality, d_165_v2_, _dafny.SeqWithoutIsStrInference([d_174_v9_, d_174_v9_]), d_175_v10_, d_170_globalState_))
        nw15_[int(4)] = d_211_v38_
        nw15_[int(5)] = d_211_v38_
        nw15_[int(6)] = _dafny.MultiSet(d_213_v41_)
        nw15_[int(7)] = (_dafny.MultiSet([d_210_v37_, d_210_v37_])).set(d_210_v37_, default__.abs(d_165_v2_))
        nw15_[int(8)] = d_211_v38_
        d_215_v42_ = nw15_
        index22_ = default__.safeIndex(457, (d_215_v42_).length(0))
        def iife22_():
            coll22_ = _dafny.Set()
            compr_22_: str
            for compr_22_ in (_dafny.Map({d_175_v10_: 613})).keys.Elements:
                d_216_v43_: str = compr_22_
                if (d_216_v43_) in (_dafny.Map({d_175_v10_: 613})):
                    coll22_ = coll22_.union(_dafny.Set([d_216_v43_]))
            return _dafny.Set(coll22_)
        (d_215_v42_)[index22_] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([iife22_()
 for d_217_i7_ in range(default__.abs(389))]))
        index23_ = default__.safeIndex(457, (d_215_v42_).length(0))
        (d_215_v42_)[index23_] = (d_211_v38_).set(d_210_v37_, default__.abs(d_165_v2_))
        hi1_ = (d_165_v2_) - (len(_dafny.Map({d_165_v2_: d_165_v2_})))
        for d_218_i8_ in range(d_165_v2_, hi1_):
            d_188_v20_ = d_188_v20_
            index24_ = default__.safeIndex(243, (d_166_v3_).length(0))
            (d_166_v3_)[index24_] = not(not((d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]))
            d_165_v2_ = d_218_i8_
            if (d_165_v2_) <= (452):
                d_165_v2_ = (len(((d_164_v1_).set(default__.safeIndex(len(d_169_v6_), len(d_164_v1_)), (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))])) + (_dafny.SeqWithoutIsStrInference([not(d_163_v0_), d_163_v0_, d_163_v0_, d_163_v0_])))) + (d_218_i8_)
                d_219_v44_: _dafny.Map
                d_219_v44_ = _dafny.Map({(0) - (d_165_v2_): (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]})
                d_220_v46_: C2
                nw16_ = C2()
                def iife23_():
                    coll23_ = _dafny.Map()
                    compr_23_: int
                    for compr_23_ in _dafny.IntegerRange(603, 37):
                        d_221_v45_: int = compr_23_
                        if ((603) <= (d_221_v45_)) and ((d_221_v45_) < (37)):
                            coll23_[(d_221_v45_) + (d_165_v2_)] = d_163_v0_
                    return _dafny.Map(coll23_)
                nw16_.ctor__((d_219_v44_) != (iife23_()
                ))
                d_220_v46_ = nw16_
                index25_ = default__.safeIndex(243, (d_166_v3_).length(0))
                (d_166_v3_)[index25_] = not(d_163_v0_)
                d_163_v0_ = (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]
                d_222_v47_: _dafny.Array
                nw17_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_222_v47_ = nw17_
                d_223_v48_: _dafny.Array
                nw18_ = _dafny.Array(None, 11)
                d_223_v48_ = nw18_
                index26_ = default__.safeIndex(852, (d_222_v47_).length(0))
                (d_222_v47_)[index26_] = d_223_v48_
                index27_ = default__.safeIndex(852, (d_222_v47_).length(0))
                rhs17_ = ((d_165_v2_) - (957)) <= (855)
                rhs18_ = d_218_i8_
                rhs19_ = ((d_173_v8_)[len((d_219_v44_) | (d_219_v44_))] if (len((d_219_v44_) | (d_219_v44_))) in (d_173_v8_) else d_218_i8_)
                rhs20_ = d_223_v48_
                lhs12_ = d_222_v47_
                lhs13_ = default__.safeIndex(852, (d_222_v47_).length(0))
                d_163_v0_ = rhs17_
                d_165_v2_ = rhs18_
                d_165_v2_ = rhs19_
                lhs12_[lhs13_] = rhs20_
            elif True:
                d_169_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tkuqi"))
                d_224_v49_: D2
                d_224_v49_ = D2_DC8(d_163_v0_, d_163_v0_, d_163_v0_)
                index28_ = default__.safeIndex(243, (d_166_v3_).length(0))
                rhs21_ = d_163_v0_
                rhs22_ = (0) - ((0) - (default__.safeModuloInt((d_218_i8_) * (479), d_165_v2_)))
                rhs23_ = (d_224_v49_).cf17
                lhs14_ = d_166_v3_
                lhs15_ = default__.safeIndex(243, (d_166_v3_).length(0))
                lhs14_[lhs15_] = rhs21_
                d_165_v2_ = rhs22_
                d_163_v0_ = rhs23_
                d_175_v10_ = _dafny.CodePoint('l')
                d_225_v50_: _dafny.Array
                nw19_ = _dafny.Array(_dafny.Set({}), 21)
                d_225_v50_ = nw19_
                d_226_v51_: _dafny.Set
                d_226_v51_ = _dafny.Set({d_165_v2_, d_165_v2_})
                index29_ = default__.safeIndex(713, (d_225_v50_).length(0))
                (d_225_v50_)[index29_] = d_226_v51_
                index30_ = default__.safeIndex(713, (d_225_v50_).length(0))
                (d_225_v50_)[index30_] = d_226_v51_
                d_227_v52_: T1
                nw20_ = C5()
                nw20_.ctor__(not(d_163_v0_), d_188_v20_)
                d_227_v52_ = nw20_
                d_227_v52_ = d_227_v52_
        d_228_i9_: int
        d_228_i9_ = 0
        with _dafny.label("3"):
            while (((d_212_v39_)[d_163_v0_] if (d_163_v0_) in (d_212_v39_) else -671)) != (d_165_v2_):
                with _dafny.c_label("3"):
                    if (d_228_i9_) >= (100):
                        raise _dafny.Break("3")
                    d_228_i9_ = (d_228_i9_) + (1)
                    index31_ = default__.safeIndex(541, (d_188_v20_).length(0))
                    (d_188_v20_)[index31_] = d_165_v2_
                    index32_ = default__.safeIndex(541, (d_188_v20_).length(0))
                    (d_188_v20_)[index32_] = ((d_203_v32_)[(d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]] if ((d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]) in (d_203_v32_) else d_165_v2_)
                    d_229_v53_: _dafny.Map
                    d_229_v53_ = _dafny.Map({d_169_v6_: (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]})
                    d_230_v54_: int
                    out1_: int
                    out1_ = default__.m0(d_229_v53_, d_170_globalState_)
                    d_230_v54_ = out1_
                    index33_ = default__.safeIndex(243, (d_166_v3_).length(0))
                    (d_166_v3_)[index33_] = d_163_v0_
                    index34_ = default__.safeIndex(541, (d_188_v20_).length(0))
                    rhs24_ = 711
                    rhs25_ = (d_230_v54_) + (-301)
                    lhs16_ = d_188_v20_
                    lhs17_ = default__.safeIndex(541, (d_188_v20_).length(0))
                    d_165_v2_ = rhs24_
                    lhs16_[lhs17_] = rhs25_
                    pass
            pass
        d_231_v55_: C3
        nw21_ = C3()
        nw21_.ctor__(d_163_v0_, not((d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]))
        d_231_v55_ = nw21_
        d_231_v55_ = (d_231_v55_ if ((d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]) not in (d_164_v1_) else d_231_v55_)
        d_232_v56_: _dafny.Seq
        d_232_v56_ = _dafny.SeqWithoutIsStrInference([d_188_v20_, d_188_v20_, (d_188_v20_ if (d_231_v55_).f19 else d_188_v20_)])
        d_232_v56_ = d_232_v56_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_188_v20_).length(0)):
            d_233_i10_: int = guard_loop_0_
            if (True) and (((0) <= (d_233_i10_)) and ((d_233_i10_) < ((d_188_v20_).length(0)))):
                (d_188_v20_)[(d_233_i10_)] = default__.safeDivisionInt(d_233_i10_, d_165_v2_)
        if d_163_v0_:
            index35_ = default__.safeIndex(243, (d_166_v3_).length(0))
            (d_166_v3_)[index35_] = True
            d_234_v57_: _dafny.Array
            out2_: _dafny.Array
            out2_ = (d_231_v55_).m1(default__.fm31(d_170_globalState_), d_170_globalState_)
            d_234_v57_ = out2_
            d_235_v58_: _dafny.Array
            def lambda5_(d_236_v55_):
                def lambda6_(d_237_i11_):
                    return _dafny.Set({(d_236_v55_).f19})

                return lambda6_

            init2_ = lambda5_(d_231_v55_)
            nw22_ = _dafny.Array(None, 29)
            for i0_2_ in range(nw22_.length(0)):
                nw22_[i0_2_] = init2_(i0_2_)
            d_235_v58_ = nw22_
            d_238_v60_: _dafny.Set
            def iife24_():
                coll24_ = _dafny.Map()
                compr_24_: int
                for compr_24_ in (d_174_v9_).Elements:
                    d_239_v59_: int = compr_24_
                    if (d_239_v59_) in (d_174_v9_):
                        coll24_[(d_239_v59_) + (d_165_v2_)] = d_165_v2_
                return _dafny.Map(coll24_)
            d_238_v60_ = _dafny.Set({(d_231_v55_).f19, (d_231_v55_).f19, (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))], not(default__.fm0(iife24_()
            , d_165_v2_, d_170_globalState_))})
            index36_ = default__.safeIndex(787, (d_235_v58_).length(0))
            (d_235_v58_)[index36_] = (d_238_v60_ if (d_164_v1_)[default__.safeIndex(d_165_v2_, len(d_164_v1_))] else _dafny.Set({(d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]}))
            index37_ = default__.safeIndex(787, (d_235_v58_).length(0))
            (d_235_v58_)[index37_] = (d_238_v60_) | (d_238_v60_)
            d_240_v61_: C3
            nw23_ = C3()
            nw23_.ctor__((d_231_v55_).f19, (len(_dafny.Set({d_165_v2_}))) > (d_165_v2_))
            d_240_v61_ = nw23_
            index38_ = default__.safeIndex(243, (d_166_v3_).length(0))
            (d_166_v3_)[index38_] = (default__.safeModuloInt(d_165_v2_, d_165_v2_)) < (len((d_238_v60_) - ((d_235_v58_)[default__.safeIndex(787, (d_235_v58_).length(0))])))
        elif True:
            d_241_v63_: _dafny.Seq
            d_241_v63_ = _dafny.SeqWithoutIsStrInference([d_165_v2_, d_165_v2_])
            pat_let_tv6_ = d_165_v2_
            d_242_v64_: _dafny.Map
            def iife25_(_pat_let0_0):
                def iife26_(d_243_dt__update__tmp_h0_):
                    def iife27_(_pat_let1_0):
                        def iife28_(d_244_dt__update_hcf2_h0_):
                            return D0_DC1((d_243_dt__update__tmp_h0_).cf1, d_244_dt__update_hcf2_h0_)
                        return iife28_(_pat_let1_0)
                    return iife27_(pat_let_tv6_)
                return iife26_(_pat_let0_0)
            def iife29_():
                coll25_ = _dafny.Map()
                compr_25_: int
                for compr_25_ in (d_241_v63_).Elements:
                    d_245_v62_: int = compr_25_
                    if (d_245_v62_) in (d_241_v63_):
                        coll25_[(d_245_v62_) - (d_165_v2_)] = d_165_v2_
                return _dafny.Map(coll25_)
            d_242_v64_ = _dafny.Map({(d_231_v55_).fm4(d_163_v0_, iife25_(d_172_v7_), len(d_203_v32_), iife29_()
            , d_170_globalState_): d_203_v32_})
            d_165_v2_ = (0) - (len((_dafny.Map({(d_231_v55_).f19: default__.fm31(d_170_globalState_)})) | (((d_242_v64_)[(d_231_v55_).f19] if ((d_231_v55_).f19) in (d_242_v64_) else d_203_v32_))))
            d_246_v65_: _dafny.Array
            nw24_ = _dafny.Array(None, 4)
            d_246_v65_ = nw24_
            d_246_v65_ = d_246_v65_
            index39_ = default__.safeIndex(243, (d_166_v3_).length(0))
            (d_166_v3_)[index39_] = (d_231_v55_).f19
            d_165_v2_ = d_165_v2_
            d_174_v9_ = d_174_v9_
        d_247_v66_: _dafny.Array
        nw25_ = _dafny.Array(_dafny.Array(None, 0), 14)
        d_247_v66_ = nw25_
        index40_ = default__.safeIndex(244, (d_247_v66_).length(0))
        (d_247_v66_)[index40_] = d_188_v20_
        index41_ = default__.safeIndex(244, (d_247_v66_).length(0))
        (d_247_v66_)[index41_] = d_188_v20_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_166_v3_).length(0)):
            d_248_i12_: int = guard_loop_1_
            if (True) and (((0) <= (d_248_i12_)) and ((d_248_i12_) < ((d_166_v3_).length(0)))):
                (d_166_v3_)[(d_248_i12_)] = not (d_163_v0_) or (((d_231_v55_).f19 if (d_231_v55_).f19 else (d_231_v55_).f19))
        d_249_v67_: D5
        d_249_v67_ = D5_DC13()
        source7_ = d_249_v67_
        if source7_.is_DC13:
            d_250_v68_: _dafny.Set
            d_250_v68_ = _dafny.Set({d_165_v2_})
            index42_ = default__.safeIndex(243, (d_166_v3_).length(0))
            rhs26_ = (d_165_v2_) - (d_165_v2_)
            rhs27_ = (d_250_v68_) == (_dafny.Set({d_165_v2_, len(d_169_v6_)}))
            lhs18_ = d_166_v3_
            lhs19_ = default__.safeIndex(243, (d_166_v3_).length(0))
            d_165_v2_ = rhs26_
            lhs18_[lhs19_] = rhs27_
            d_251_v69_: _dafny.Array
            nw26_ = _dafny.Array(D5.default()(), 8)
            d_251_v69_ = nw26_
            d_252_v70_: _dafny.Map
            d_252_v70_ = _dafny.Map({d_231_v55_: d_173_v8_})
            d_253_v71_: _dafny.Seq
            d_253_v71_ = _dafny.SeqWithoutIsStrInference([d_173_v8_])
            pat_let_tv7_ = d_165_v2_
            pat_let_tv8_ = d_165_v2_
            d_254_v72_: _dafny.Array
            nw27_ = _dafny.Array(None, 23)
            nw27_[int(0)] = ((d_252_v70_)[d_231_v55_] if (d_231_v55_) in (d_252_v70_) else d_173_v8_)
            nw27_[int(1)] = d_173_v8_
            nw27_[int(2)] = d_173_v8_
            nw27_[int(3)] = d_173_v8_
            nw27_[int(4)] = (d_173_v8_).set(d_165_v2_, d_165_v2_)
            nw27_[int(5)] = ((d_253_v71_)[default__.safeIndex(d_165_v2_, len(d_253_v71_))] if (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))] else d_173_v8_)
            nw27_[int(6)] = (d_173_v8_ if d_163_v0_ else _dafny.Map({303: 904}))
            def iife30_(_pat_let2_0):
                def iife31_(d_255_dt__update__tmp_h1_):
                    def iife32_(_pat_let3_0):
                        def iife33_(d_256_dt__update_hcf3_h0_):
                            return D1_DC2(d_256_dt__update_hcf3_h0_)
                        return iife33_(_pat_let3_0)
                    return iife32_(_dafny.Map({pat_let_tv7_: pat_let_tv8_}))
                return iife31_(_pat_let2_0)
            nw27_[int(7)] = ((iife30_(D1_DC2(d_173_v8_))).cf3) | (d_173_v8_)
            nw27_[int(8)] = d_173_v8_
            nw27_[int(9)] = default__.fm39(d_165_v2_, d_170_globalState_)
            nw27_[int(10)] = ((d_173_v8_).set(len(default__.fm28(203, d_165_v2_, not(True), d_170_globalState_)), d_165_v2_)) | (d_173_v8_)
            nw27_[int(11)] = d_173_v8_
            nw27_[int(12)] = d_173_v8_
            nw27_[int(13)] = d_173_v8_
            nw27_[int(14)] = d_173_v8_
            nw27_[int(15)] = d_173_v8_
            nw27_[int(16)] = d_173_v8_
            nw27_[int(17)] = d_173_v8_
            nw27_[int(18)] = (d_173_v8_ if d_163_v0_ else d_173_v8_)
            nw27_[int(19)] = default__.fm39(d_165_v2_, d_170_globalState_)
            nw27_[int(20)] = d_173_v8_
            nw27_[int(21)] = d_173_v8_
            nw27_[int(22)] = d_173_v8_
            d_254_v72_ = nw27_
            index43_ = default__.safeIndex(920, (d_254_v72_).length(0))
            (d_254_v72_)[index43_] = d_173_v8_
            d_257_v73_: _dafny.Seq
            d_257_v73_ = _dafny.SeqWithoutIsStrInference([d_165_v2_, d_165_v2_])
            d_258_v74_: _dafny.Set
            d_258_v74_ = _dafny.Set({d_257_v73_, d_257_v73_, d_257_v73_, d_257_v73_})
            index44_ = default__.safeIndex(318, (d_188_v20_).length(0))
            (d_188_v20_)[index44_] = len(d_258_v74_)
            d_259_v75_: _dafny.Map
            d_259_v75_ = _dafny.Map({(d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]: (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]})
            d_260_v76_: D4
            d_260_v76_ = D4_DC11((d_231_v55_).f19)
            d_261_v77_: _dafny.Set
            d_261_v77_ = _dafny.Set({(d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))], not((d_260_v76_).cf21)})
            index45_ = default__.safeIndex(920, (d_254_v72_).length(0))
            index46_ = default__.safeIndex(318, (d_188_v20_).length(0))
            rhs28_ = (d_251_v69_ if ((d_164_v1_)[default__.safeIndex(d_165_v2_, len(d_164_v1_))]) not in (d_259_v75_) else d_251_v69_)
            rhs29_ = (d_173_v8_).set((d_174_v9_).cardinality, len((d_257_v73_ if (d_231_v55_).f19 else d_257_v73_)))
            rhs30_ = (-208) - (len((d_261_v77_) | (_dafny.Set({(d_231_v55_).f19}))))
            lhs20_ = d_254_v72_
            lhs21_ = default__.safeIndex(920, (d_254_v72_).length(0))
            lhs22_ = d_188_v20_
            lhs23_ = default__.safeIndex(318, (d_188_v20_).length(0))
            d_251_v69_ = rhs28_
            lhs20_[lhs21_] = rhs29_
            lhs22_[lhs23_] = rhs30_
            if (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]:
                d_262_v78_: _dafny.Seq
                d_262_v78_ = _dafny.SeqWithoutIsStrInference([d_257_v73_, (d_257_v73_ if (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))] else d_257_v73_), d_257_v73_, d_257_v73_, d_257_v73_])
                d_262_v78_ = d_262_v78_
                d_263_v79_: C3
                nw28_ = C3()
                nw28_.ctor__(not ((d_231_v55_).f19) or ((d_231_v55_).f19), ((d_231_v55_).f19 if d_163_v0_ else (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]))
                d_263_v79_ = nw28_
                d_264_v80_: _dafny.Array
                nw29_ = _dafny.Array(_dafny.CodePoint('D'), 12)
                d_264_v80_ = nw29_
                nw30_ = _dafny.Array(_dafny.CodePoint('D'), 11)
                d_264_v80_ = nw30_
                d_166_v3_ = d_166_v3_
                d_265_v81_: C3
                nw31_ = C3()
                nw31_.ctor__((d_231_v55_).f19, (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))])
                d_265_v81_ = nw31_
            elif True:
                d_266_v82_: _dafny.Array
                out3_: _dafny.Array
                out3_ = (d_231_v55_).m1((d_188_v20_)[default__.safeIndex(318, (d_188_v20_).length(0))], d_170_globalState_)
                d_266_v82_ = out3_
                d_267_v83_: D1
                d_267_v83_ = D1_DC5(d_169_v6_, d_166_v3_, d_165_v2_)
                d_268_v84_: _dafny.Map
                d_269_v85_: bool
                d_270_v86_: bool
                d_271_v87_: _dafny.Seq
                out4_: _dafny.Map
                out5_: bool
                out6_: bool
                out7_: _dafny.Seq
                out4_, out5_, out6_, out7_ = (d_231_v55_).m2((d_188_v20_)[default__.safeIndex(318, (d_188_v20_).length(0))], (d_169_v6_).set(default__.safeIndex(-928, len(d_169_v6_)), d_175_v10_), (d_267_v83_).cf10, d_170_globalState_)
                d_268_v84_ = out4_
                d_269_v85_ = out5_
                d_270_v86_ = out6_
                d_271_v87_ = out7_
                index47_ = default__.safeIndex(318, (d_188_v20_).length(0))
                (d_188_v20_)[index47_] = default__.safeDivisionInt(((d_257_v73_)[default__.safeIndex((d_188_v20_)[default__.safeIndex(318, (d_188_v20_).length(0))], len(d_257_v73_))]) + (d_165_v2_), (0) - ((d_188_v20_)[default__.safeIndex(318, (d_188_v20_).length(0))]))
                index48_ = default__.safeIndex(318, (d_188_v20_).length(0))
                (d_188_v20_)[index48_] = (default__.fm31(d_170_globalState_)) * ((0) - (default__.safeModuloInt(((d_174_v9_).set(d_165_v2_, default__.abs((d_188_v20_)[default__.safeIndex(318, (d_188_v20_).length(0))]))).cardinality, d_165_v2_)))
                d_166_v3_ = d_166_v3_
            index49_ = default__.safeIndex(243, (d_166_v3_).length(0))
            (d_166_v3_)[index49_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lv"))) <= (d_169_v6_)
        elif source7_.is_DC14:
            d_272___mcc_h0_ = source7_.cf23
            d_273_cf23_ = d_272___mcc_h0_
            d_274_v88_: _dafny.Array
            nw32_ = _dafny.Array(None, 3)
            nw32_[int(0)] = d_166_v3_
            nw32_[int(1)] = d_166_v3_
            nw32_[int(2)] = d_166_v3_
            d_274_v88_ = nw32_
            index50_ = default__.safeIndex(37, (d_274_v88_).length(0))
            (d_274_v88_)[index50_] = (d_166_v3_ if d_273_cf23_ else d_166_v3_)
            index51_ = default__.safeIndex(37, (d_274_v88_).length(0))
            (d_274_v88_)[index51_] = d_166_v3_
            d_275_v89_: _dafny.Seq
            d_275_v89_ = _dafny.SeqWithoutIsStrInference([d_165_v2_])
            d_276_v90_: _dafny.Seq
            d_276_v90_ = _dafny.SeqWithoutIsStrInference([d_166_v3_, d_166_v3_])
            d_277_v91_: _dafny.Map
            d_277_v91_ = _dafny.Map({(d_275_v89_) + (default__.fm16(len(d_276_v90_), d_165_v2_, (d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))], d_165_v2_, d_170_globalState_)): len(d_169_v6_)})
            d_277_v91_ = (d_277_v91_).set(d_275_v89_, default__.safeModuloInt(d_165_v2_, d_165_v2_))
            d_278_v92_: _dafny.MultiSet
            d_278_v92_ = _dafny.MultiSet([d_192_v22_, d_192_v22_, d_192_v22_])
            pat_let_tv9_ = d_173_v8_
            def iife34_(_pat_let4_0):
                def iife35_(d_279_dt__update__tmp_h2_):
                    def iife36_(_pat_let5_0):
                        def iife37_(d_280_dt__update_hcf3_h1_):
                            return D1_DC2(d_280_dt__update_hcf3_h1_)
                        return iife37_(_pat_let5_0)
                    return iife36_(pat_let_tv9_)
                return iife35_(_pat_let4_0)
            d_169_v6_ = default__.fm17((_dafny.MultiSet([d_192_v22_, d_192_v22_, iife34_(d_192_v22_)])).intersection(d_278_v92_), d_170_globalState_)
            d_165_v2_ = (297) - (d_165_v2_)
        elif source7_.is_DC12:
            d_281___mcc_h1_ = source7_.cf22
            d_282_cf22_ = d_281___mcc_h1_
            d_283_v93_: C2
            nw33_ = C2()
            nw33_.ctor__(not(not(d_163_v0_)))
            d_283_v93_ = nw33_
            index52_ = default__.safeIndex(244, (d_247_v66_).length(0))
            (d_247_v66_)[index52_] = d_188_v20_
            index53_ = default__.safeIndex(244, (d_247_v66_).length(0))
            (d_247_v66_)[index53_] = d_188_v20_
            d_284_v94_: C5
            nw34_ = C5()
            nw34_.ctor__(not (True) or (d_163_v0_), d_188_v20_)
            d_284_v94_ = nw34_
        elif True:
            d_285___mcc_h2_ = source7_.cf24
            d_286_cf24_ = d_285___mcc_h2_
            if (d_231_v55_).f19:
                d_163_v0_ = (d_231_v55_).f19
                d_163_v0_ = (d_231_v55_).f19
                index54_ = default__.safeIndex(244, (d_247_v66_).length(0))
                (d_247_v66_)[index54_] = d_188_v20_
                d_287_v95_: _dafny.Array
                out8_: _dafny.Array
                out8_ = (d_231_v55_).m1(d_165_v2_, d_170_globalState_)
                d_287_v95_ = out8_
                d_163_v0_ = (d_231_v55_).f19
            elif True:
                d_288_v96_: _dafny.Map
                d_289_v97_: bool
                d_290_v98_: bool
                d_291_v99_: _dafny.Seq
                out9_: _dafny.Map
                out10_: bool
                out11_: bool
                out12_: _dafny.Seq
                out9_, out10_, out11_, out12_ = (d_231_v55_).m2(d_165_v2_, d_169_v6_, (0) - (d_165_v2_), d_170_globalState_)
                d_288_v96_ = out9_
                d_289_v97_ = out10_
                d_290_v98_ = out11_
                d_291_v99_ = out12_
                d_292_v100_: T0
                nw35_ = C5()
                nw35_.ctor__(d_289_v97_, d_188_v20_)
                d_292_v100_ = nw35_
                d_293_v101_: _dafny.Map
                d_293_v101_ = _dafny.Map({d_165_v2_: d_292_v100_})
                d_290_v98_ = not(not((((0) - (d_165_v2_)) * (len(d_293_v101_))) >= (d_165_v2_)))
                d_294_v102_: _dafny.Seq
                d_294_v102_ = _dafny.SeqWithoutIsStrInference([d_166_v3_])
                d_295_v103_: D1
                d_295_v103_ = D1_DC4((d_294_v102_)[default__.safeIndex(d_165_v2_, len(d_294_v102_))], d_165_v2_, d_175_v10_)
                d_290_v98_ = ((d_295_v103_).cf6) < (d_165_v2_)
                d_165_v2_ = default__.safeDivisionInt((d_165_v2_) + ((0) - (d_165_v2_)), d_165_v2_)
                index55_ = default__.safeIndex(617, (d_188_v20_).length(0))
                (d_188_v20_)[index55_] = (d_165_v2_) - (d_165_v2_)
                index56_ = default__.safeIndex(617, (d_188_v20_).length(0))
                (d_188_v20_)[index56_] = d_165_v2_
            if d_163_v0_:
                d_296_v104_: _dafny.Seq
                d_296_v104_ = _dafny.SeqWithoutIsStrInference([d_165_v2_, d_165_v2_, (d_174_v9_).cardinality])
                d_297_v105_: _dafny.Seq
                d_297_v105_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epfsk"))), d_165_v2_]), d_296_v104_, d_296_v104_])
                d_298_v106_: _dafny.Map
                d_299_v107_: bool
                d_300_v108_: bool
                d_301_v109_: _dafny.Seq
                out13_: _dafny.Map
                out14_: bool
                out15_: bool
                out16_: _dafny.Seq
                out13_, out14_, out15_, out16_ = (d_231_v55_).m2((len(_dafny.Set({(d_231_v55_).f19, (d_231_v55_).f19}))) + (len((d_297_v105_)[default__.safeIndex(d_165_v2_, len(d_297_v105_))])), d_169_v6_, d_165_v2_, d_170_globalState_)
                d_298_v106_ = out13_
                d_299_v107_ = out14_
                d_300_v108_ = out15_
                d_301_v109_ = out16_
                arr0_ = (d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]
                index57_ = default__.safeIndex(267, ((d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]).length(0))
                arr0_[index57_] = default__.safeModuloInt(d_165_v2_, len(d_301_v109_))
                d_302_v110_: _dafny.Array
                def lambda7_(d_303_v108_, d_304_v107_, d_305_v3_, d_306_v0_, d_307_v55_):
                    def lambda8_(d_308_i13_):
                        return (_dafny.Set({_dafny.Map({d_303_v108_: d_304_v107_}), _dafny.Map({(d_305_v3_)[default__.safeIndex(243, (d_305_v3_).length(0))]: d_306_v0_})})) | (_dafny.Set({_dafny.Map({False: d_303_v108_}), _dafny.Map({d_304_v107_: (d_307_v55_).f19})}))

                    return lambda8_

                init3_ = lambda7_(d_300_v108_, d_299_v107_, d_166_v3_, d_163_v0_, d_231_v55_)
                nw36_ = _dafny.Array(None, 23)
                for i0_3_ in range(nw36_.length(0)):
                    nw36_[i0_3_] = init3_(i0_3_)
                d_302_v110_ = nw36_
                d_309_v111_: _dafny.Map
                d_309_v111_ = _dafny.Map({d_163_v0_: not(d_299_v107_)})
                d_310_v112_: _dafny.Set
                d_310_v112_ = _dafny.Set({d_309_v111_, d_309_v111_})
                index58_ = default__.safeIndex(396, (d_302_v110_).length(0))
                (d_302_v110_)[index58_] = d_310_v112_
                d_311_v113_: _dafny.Set
                d_311_v113_ = _dafny.Set({d_165_v2_})
                arr1_ = (d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]
                index59_ = default__.safeIndex(267, ((d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]).length(0))
                index60_ = default__.safeIndex(396, (d_302_v110_).length(0))
                rhs31_ = (d_165_v2_) + (len(d_311_v113_))
                rhs32_ = d_165_v2_
                rhs33_ = d_212_v39_
                rhs34_ = d_165_v2_
                rhs35_ = d_310_v112_
                lhs24_ = (d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]
                lhs25_ = default__.safeIndex(267, ((d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]).length(0))
                lhs26_ = d_302_v110_
                lhs27_ = default__.safeIndex(396, (d_302_v110_).length(0))
                d_165_v2_ = rhs31_
                d_165_v2_ = rhs32_
                d_212_v39_ = rhs33_
                lhs24_[lhs25_] = rhs34_
                lhs26_[lhs27_] = rhs35_
                arr2_ = (d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]
                index61_ = default__.safeIndex(267, ((d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]).length(0))
                rhs36_ = ((d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))])[default__.safeIndex(267, ((d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]).length(0))]
                rhs37_ = _dafny.Map({((d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))])[default__.safeIndex(267, ((d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]).length(0))]: d_165_v2_})
                lhs28_ = (d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]
                lhs29_ = default__.safeIndex(267, ((d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]).length(0))
                lhs28_[lhs29_] = rhs36_
                d_173_v8_ = rhs37_
                d_312_v114_: _dafny.Map
                d_312_v114_ = _dafny.Map({True: d_296_v104_})
                arr3_ = (d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]
                index62_ = default__.safeIndex(267, ((d_247_v66_)[default__.safeIndex(244, (d_247_v66_).length(0))]).length(0))
                arr3_[index62_] = default__.safeModuloInt(len(d_312_v114_), default__.fm13(d_300_v108_, (d_231_v55_).f19, d_175_v10_, d_170_globalState_))
                index63_ = default__.safeIndex(243, (d_166_v3_).length(0))
                (d_166_v3_)[index63_] = ((d_231_v55_).f19) == ((d_231_v55_).f19)
            elif True:
                index64_ = default__.safeIndex(243, (d_166_v3_).length(0))
                (d_166_v3_)[index64_] = False
                d_313_v116_: _dafny.Seq
                d_313_v116_ = _dafny.SeqWithoutIsStrInference([len(d_169_v6_), d_165_v2_])
                index65_ = default__.safeIndex(243, (d_166_v3_).length(0))
                def iife38_():
                    coll26_ = _dafny.Map()
                    compr_26_: int
                    for compr_26_ in (d_313_v116_).Elements:
                        d_314_v115_: int = compr_26_
                        if (d_314_v115_) in (d_313_v116_):
                            coll26_[(d_314_v115_) + (d_165_v2_)] = ((d_174_v9_).set(272, default__.abs(d_165_v2_))).cardinality
                    return _dafny.Map(coll26_)
                (d_166_v3_)[index65_] = default__.fm0(iife38_()
                , d_165_v2_, d_170_globalState_)
                d_315_v117_: D2
                d_315_v117_ = D2_DC6(d_313_v116_)
                d_316_v118_: D2
                d_316_v118_ = D2_DC6((d_315_v117_).cf11)
                d_317_v119_: _dafny.Map
                d_317_v119_ = _dafny.Map({(0) - (d_165_v2_): d_169_v6_})
                d_318_v120_: _dafny.Map
                d_318_v120_ = _dafny.Map({(d_231_v55_).f19: d_317_v119_})
                rhs38_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "od")))
                rhs39_ = default__.fm48(d_175_v10_, (((d_318_v120_)[(d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]] if ((d_166_v3_)[default__.safeIndex(243, (d_166_v3_).length(0))]) in (d_318_v120_) else d_317_v119_)).set(d_165_v2_, d_169_v6_), d_165_v2_, (d_169_v6_) < (d_169_v6_), d_170_globalState_)
                rhs40_ = _dafny.SeqWithoutIsStrInference([d_165_v2_ for d_319_i14_ in range(default__.abs(706))])
                rhs41_ = (d_203_v32_).set((d_231_v55_).f19, 488)
                d_165_v2_ = rhs38_
                d_316_v118_ = rhs39_
                d_313_v116_ = rhs40_
                d_203_v32_ = rhs41_
                index66_ = default__.safeIndex(244, (d_247_v66_).length(0))
                nw37_ = _dafny.Array(int(0), 21)
                (d_247_v66_)[index66_] = nw37_
                d_320_v121_: T1
                nw38_ = C5()
                nw38_.ctor__(True, d_188_v20_)
                d_320_v121_ = nw38_
                d_320_v121_ = d_320_v121_
            d_321_v122_: _dafny.Array
            out17_: _dafny.Array
            out17_ = (d_231_v55_).m1((85) * (d_165_v2_), d_170_globalState_)
            d_321_v122_ = out17_
            index67_ = default__.safeIndex(243, (d_166_v3_).length(0))
            rhs42_ = d_163_v0_
            rhs43_ = (d_231_v55_).fm27(d_170_globalState_)
            lhs30_ = d_166_v3_
            lhs31_ = default__.safeIndex(243, (d_166_v3_).length(0))
            lhs30_[lhs31_] = rhs42_
            d_175_v10_ = rhs43_
        d_165_v2_ = ((0) - ((0) - (d_165_v2_))) * (d_165_v2_)
        _dafny.print(_dafny.string_of(d_163_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v1_) == (_dafny.SeqWithoutIsStrInference([False, True, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_165_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v3_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_167_v4_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v5_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference([False, True, False, True]), _dafny.SeqWithoutIsStrInference([False, False, True, False, False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_169_v6_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f1)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f4) == (_dafny.SeqWithoutIsStrInference([False, True, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f5) == (_dafny.Set({_dafny.SeqWithoutIsStrInference([False, True, False, True]), _dafny.SeqWithoutIsStrInference([False, False, True, False, False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f7) == (_dafny.MultiSet([-199]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_globalState_).f10)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_170_globalState_).f11).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_172_v7_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_172_v7_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_173_v8_) == (_dafny.Map({199: 833}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_174_v9_) == (_dafny.MultiSet([-199]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_175_v10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_176_i1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v20_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_191_v21_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_192_v22_).cf3) == (_dafny.Map({199: 833}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_193_i4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_v32_) == (_dafny.Map({True: 78001}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_210_v37_) == (_dafny.Set({_dafny.CodePoint('d')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v38_) == (_dafny.MultiSet([_dafny.Set({_dafny.CodePoint('d')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v39_) == (_dafny.MultiSet([False, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v41_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({_dafny.CodePoint('d')}), _dafny.Set({_dafny.CodePoint('p')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_215_v42_)[0]) == (_dafny.MultiSet([_dafny.Set({_dafny.CodePoint('d')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_215_v42_)[1]) == (_dafny.MultiSet([_dafny.Set({_dafny.CodePoint('d')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_215_v42_)[2]) == (_dafny.MultiSet([_dafny.Set({_dafny.CodePoint('t'), _dafny.CodePoint('c'), _dafny.CodePoint('h')}), _dafny.Set({_dafny.CodePoint('v')}), _dafny.Set({_dafny.CodePoint('e'), _dafny.CodePoint('n')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_215_v42_)[3]) == (_dafny.MultiSet([_dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('i')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')}), _dafny.Set({_dafny.CodePoint('o')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_215_v42_)[4]) == (_dafny.MultiSet([_dafny.Set({_dafny.CodePoint('d')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_215_v42_)[5]) == (_dafny.MultiSet([_dafny.Set({_dafny.CodePoint('d')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_215_v42_)[6]) == (_dafny.MultiSet([_dafny.Set({_dafny.CodePoint('d')}), _dafny.Set({_dafny.CodePoint('p')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_215_v42_)[7]) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_215_v42_)[8]) == (_dafny.MultiSet([_dafny.Set({_dafny.CodePoint('d')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_228_i9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_231_v55_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_232_v56_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v66_)[6])[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
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
        return lambda: D1_DC3(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any), ('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({self.cf8.VerbatimString(True)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC7(D2, NamedTuple('DC7', [('cf12', Any), ('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({self.cf12.VerbatimString(True)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf11 == __o.cf11
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

class D3_DC9(D3, NamedTuple('DC9', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC11(D4, NamedTuple('DC11', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13()
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
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC13(D5, NamedTuple('DC13', [])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13)
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(int(0), int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)

class D6_DC17(D6, NamedTuple('DC17', [('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC20(int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC20(D7, NamedTuple('DC20', [('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC22(D8, NamedTuple('DC22', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC24(False, False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)

class D9_DC24(D9, NamedTuple('DC24', [('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf39', Any), ('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC23(D9, NamedTuple('DC23', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)

class D10_DC26(D10, NamedTuple('DC26', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC28(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D11_DC28)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)

class D11_DC28(D11, NamedTuple('DC28', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC29(D11, NamedTuple('DC29', [('cf45', Any), ('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC31(D11, NamedTuple('DC31', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC33(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)

class D12_DC33(D12, NamedTuple('DC33', [('cf52', Any), ('cf53', Any), ('cf54', Any), ('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({self.cf52.VerbatimString(True)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC32(D12, NamedTuple('DC32', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f13(self):
        return self._f13
    @f13.setter
    def f13(self, value):
        self._f13 = value
    def m1(self, p0, globalState):
        pass

    def m2(self, p0, p1, p2, globalState):
        pass


class T1:
    pass

class GlobalState:
    def  __init__(self):
        self._f0: int = int(0)
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f2: bool = False
        self._f3: bool = False
        self._f4: _dafny.Seq = _dafny.Seq({})
        self._f5: _dafny.Set = _dafny.Set({})
        self._f6: bool = False
        self._f7: _dafny.MultiSet = _dafny.MultiSet({})
        self._f8: bool = False
        self._f9: int = int(0)
        self._f10: _dafny.Array = _dafny.Array(None, 0)
        self._f11: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f12: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12

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

class C0:
    def  __init__(self):
        self._f15: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f15):
        (self)._f15 = f15

    def fm10(self, p0, p1, p2, p3, globalState):
        return (not (False) or (False)) not in ((_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([True]))})) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([True]))})))

    @property
    def f15(self):
        return self._f15

class C1(T1):
    def  __init__(self):
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        self.f18: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f17: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f14(self):
        return self._f14
    def ctor__(self, f17, f18, f14):
        (self)._f17 = f17
        (self).f18 = f18
        (self)._f14 = f14

    def fm6(self, globalState):
        return (self).f17

    def fm7(self, p0, p1, p2, globalState):
        return -859

    def fm14(self, p0, p1, p2, p3, globalState):
        return not((len(self.f18)) >= (default__.safeDivisionInt(len(_dafny.Set({(self).f17})), 41)))

    def fm15(self, globalState):
        return (((_dafny.MultiSet([(self).f17])) - (_dafny.MultiSet([(self).f17]))) | (_dafny.MultiSet([False, (self).f17]))).cardinality

    def m5(self, p0, globalState):
        d_322_i0_: int
        d_322_i0_ = 0
        with _dafny.label("4"):
            while (self).f17:
                with _dafny.c_label("4"):
                    if (d_322_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_322_i0_ = (d_322_i0_) + (1)
                    d_323_v0_: int
                    d_323_v0_ = 669
                    d_324_v1_: bool
                    d_324_v1_ = True
                    rhs44_ = default__.safeDivisionInt(703, (d_323_v0_) * (d_323_v0_))
                    rhs45_ = (self).f17
                    rhs46_ = (d_323_v0_) <= (d_323_v0_)
                    d_323_v0_ = rhs44_
                    d_324_v1_ = rhs45_
                    d_324_v1_ = rhs46_
                    d_325_v2_: _dafny.Map
                    d_325_v2_ = _dafny.Map({(self).fm14(d_323_v0_, (self).f17, d_323_v0_, d_323_v0_, globalState): (self).f14})
                    d_326_v3_: _dafny.Map
                    d_326_v3_ = _dafny.Map({d_323_v0_: ((d_325_v2_)[(self).f17] if ((self).f17) in (d_325_v2_) else (self).f14)})
                    rhs47_ = len((_dafny.Map({d_323_v0_: (self).f14})) | (d_326_v3_))
                    rhs48_ = False
                    d_323_v0_ = rhs47_
                    d_324_v1_ = rhs48_
                    d_327_v4_: str
                    d_327_v4_ = _dafny.CodePoint('t')
                    d_327_v4_ = (self.f18)[default__.safeIndex(d_323_v0_, len(self.f18))]
                    index68_ = default__.safeIndex(585, ((self).f14).length(0))
                    ((self).f14)[index68_] = (d_323_v0_) - (d_323_v0_)
                    d_328_v5_: _dafny.Set
                    d_328_v5_ = _dafny.Set({True, d_324_v1_, (self).f17})
                    d_329_v6_: _dafny.Map
                    d_329_v6_ = _dafny.Map({(d_328_v5_) - (_dafny.Set({(self).f17, True})): ((0) - (d_323_v0_)) - (d_323_v0_)})
                    index69_ = default__.safeIndex(585, ((self).f14).length(0))
                    ((self).f14)[index69_] = len(d_329_v6_)
                    pass
            pass
        d_330_v7_: int
        d_330_v7_ = -432
        hi2_ = 98
        for d_331_i1_ in range(d_330_v7_, hi2_):
            d_332_v8_: _dafny.Map
            d_332_v8_ = _dafny.Map({(self).f17: d_331_i1_})
            d_333_v9_: _dafny.Seq
            d_333_v9_ = _dafny.SeqWithoutIsStrInference([(self).f17])
            d_334_v10_: _dafny.Map
            d_334_v10_ = _dafny.Map({d_333_v9_: d_330_v7_})
            d_335_v11_: _dafny.Seq
            d_335_v11_ = _dafny.SeqWithoutIsStrInference([d_331_i1_, len((d_334_v10_).set(_dafny.SeqWithoutIsStrInference([(self).f17, (self).f17]), d_331_i1_)), 740])
            d_336_v12_: _dafny.Map
            d_336_v12_ = _dafny.Map({not(True): (default__.fm16(d_330_v7_, len(d_332_v8_), (self).f17, d_331_i1_, globalState) if (self).f17 else d_335_v11_)})
            d_336_v12_ = (d_336_v12_).set(not ((self).f17) or ((self).f17), d_335_v11_)
            d_330_v7_ = (default__.safeDivisionInt(d_331_i1_, d_330_v7_)) * (d_330_v7_)
            d_330_v7_ = default__.safeDivisionInt(160, 5)
            d_337_v13_: _dafny.MultiSet
            d_337_v13_ = _dafny.MultiSet([d_331_i1_])
            d_338_v14_: _dafny.Set
            d_338_v14_ = _dafny.Set({(self).f17})
            d_339_v15_: _dafny.Set
            d_339_v15_ = _dafny.Set({d_338_v14_})
            d_330_v7_ = ((d_337_v13_)[(d_330_v7_) + (len(self.f18))] if ((d_330_v7_) + (len(self.f18))) in (d_337_v13_) else len(d_339_v15_))
        d_340_i2_: int
        d_340_i2_ = 0
        with _dafny.label("5"):
            while not(not((self).f17)):
                with _dafny.c_label("5"):
                    if (d_340_i2_) >= (100):
                        raise _dafny.Break("5")
                    d_340_i2_ = (d_340_i2_) + (1)
                    d_341_v16_: C0
                    nw39_ = C0()
                    nw39_.ctor__(d_330_v7_)
                    d_341_v16_ = nw39_
                    d_342_v17_: _dafny.Map
                    d_342_v17_ = _dafny.Map({d_330_v7_: (d_341_v16_).f15})
                    d_343_v18_: D1
                    d_343_v18_ = D1_DC2(d_342_v17_)
                    source8_ = d_343_v18_
                    if source8_.is_DC3:
                        d_344___mcc_h0_ = source8_.cf4
                        d_345_cf4_ = d_344___mcc_h0_
                        d_346_v19_: D1
                        d_346_v19_ = D1_DC3(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_347_i3_ in range(default__.abs(335))])))
                        d_348_v20_: _dafny.Array
                        nw40_ = _dafny.Array(None, 26)
                        nw40_[int(0)] = D1_DC3(len(self.f18))
                        nw40_[int(1)] = d_346_v19_
                        nw40_[int(2)] = d_346_v19_
                        nw40_[int(3)] = D1_DC3((d_341_v16_).f15)
                        nw40_[int(4)] = d_346_v19_
                        nw40_[int(5)] = d_346_v19_
                        nw40_[int(6)] = d_346_v19_
                        nw40_[int(7)] = d_346_v19_
                        nw40_[int(8)] = d_346_v19_
                        nw40_[int(9)] = d_346_v19_
                        nw40_[int(10)] = D1_DC3(d_345_cf4_)
                        nw40_[int(11)] = d_346_v19_
                        nw40_[int(12)] = d_346_v19_
                        nw40_[int(13)] = D1_DC3(115)
                        nw40_[int(14)] = d_346_v19_
                        nw40_[int(15)] = d_346_v19_
                        nw40_[int(16)] = d_346_v19_
                        nw40_[int(17)] = d_346_v19_
                        nw40_[int(18)] = d_346_v19_
                        nw40_[int(19)] = d_346_v19_
                        nw40_[int(20)] = d_346_v19_
                        nw40_[int(21)] = d_346_v19_
                        nw40_[int(22)] = d_346_v19_
                        nw40_[int(23)] = d_346_v19_
                        nw40_[int(24)] = d_346_v19_
                        nw40_[int(25)] = d_346_v19_
                        d_348_v20_ = nw40_
                        index70_ = default__.safeIndex(871, (d_348_v20_).length(0))
                        (d_348_v20_)[index70_] = d_346_v19_
                        d_349_v21_: bool
                        d_349_v21_ = True
                        d_350_v22_: _dafny.Seq
                        d_350_v22_ = _dafny.SeqWithoutIsStrInference([len(self.f18), 31, d_345_cf4_])
                        pat_let_tv10_ = d_350_v22_
                        pat_let_tv11_ = d_349_v21_
                        pat_let_tv12_ = d_345_cf4_
                        index71_ = default__.safeIndex(871, (d_348_v20_).length(0))
                        def iife39_(_pat_let6_0):
                            def iife40_(d_351_dt__update__tmp_h0_):
                                def iife41_(_pat_let7_0):
                                    def iife42_(d_352_dt__update_hcf4_h0_):
                                        return D1_DC3(d_352_dt__update_hcf4_h0_)
                                    return iife42_(_pat_let7_0)
                                return iife41_((len(pat_let_tv10_) if pat_let_tv11_ else pat_let_tv12_))
                            return iife40_(_pat_let6_0)
                        rhs49_ = iife39_(D1_DC3(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))))
                        rhs50_ = default__.safeDivisionInt(d_330_v7_, d_345_cf4_)
                        rhs51_ = not((self).f17)
                        rhs52_ = (d_341_v16_).f15
                        rhs53_ = (self).f17
                        lhs32_ = d_348_v20_
                        lhs33_ = default__.safeIndex(871, (d_348_v20_).length(0))
                        lhs32_[lhs33_] = rhs49_
                        d_330_v7_ = rhs50_
                        d_349_v21_ = rhs51_
                        d_345_cf4_ = rhs52_
                        d_349_v21_ = rhs53_
                        d_353_v23_: D2
                        d_353_v23_ = D2_DC7(self.f18, (d_341_v16_).f15, (d_341_v16_).f15, (self).f17)
                        d_354_v24_: _dafny.Seq
                        d_354_v24_ = _dafny.SeqWithoutIsStrInference([d_349_v21_, d_349_v21_])
                        d_355_v25_: str
                        d_355_v25_ = _dafny.CodePoint('l')
                        pat_let_tv13_ = d_342_v17_
                        d_356_v26_: _dafny.MultiSet
                        def iife43_(_pat_let8_0):
                            def iife44_(d_357_dt__update__tmp_h1_):
                                def iife45_(_pat_let9_0):
                                    def iife46_(d_358_dt__update_hcf3_h0_):
                                        return D1_DC2(d_358_dt__update_hcf3_h0_)
                                    return iife46_(_pat_let9_0)
                                return iife45_(pat_let_tv13_)
                            return iife44_(_pat_let8_0)
                        d_356_v26_ = _dafny.MultiSet([d_343_v18_, iife43_(d_343_v18_)])
                        d_359_v27_: _dafny.Array
                        nw41_ = _dafny.Array(None, 25)
                        nw41_[int(0)] = d_349_v21_
                        nw41_[int(1)] = (d_353_v23_).cf15
                        nw41_[int(2)] = d_349_v21_
                        nw41_[int(3)] = (self).f17
                        nw41_[int(4)] = d_349_v21_
                        nw41_[int(5)] = (d_330_v7_) < (len(d_354_v24_))
                        nw41_[int(6)] = (self).f17
                        nw41_[int(7)] = d_349_v21_
                        nw41_[int(8)] = d_349_v21_
                        nw41_[int(9)] = (self).f17
                        nw41_[int(10)] = (self).f17
                        nw41_[int(11)] = not(not(d_349_v21_))
                        nw41_[int(12)] = (d_355_v25_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eyusi")))
                        nw41_[int(13)] = d_349_v21_
                        nw41_[int(14)] = (d_354_v24_)[default__.safeIndex(len(d_350_v22_), len(d_354_v24_))]
                        nw41_[int(15)] = not(d_349_v21_)
                        nw41_[int(16)] = False
                        nw41_[int(17)] = (self).f17
                        nw41_[int(18)] = False
                        nw41_[int(19)] = d_349_v21_
                        nw41_[int(20)] = (d_349_v21_) == ((self).f17)
                        nw41_[int(21)] = (self).f17
                        nw41_[int(22)] = d_349_v21_
                        nw41_[int(23)] = (self).f17
                        nw41_[int(24)] = (default__.fm17(d_356_v26_, globalState)) != (self.f18)
                        d_359_v27_ = nw41_
                        index72_ = default__.safeIndex(336, (d_359_v27_).length(0))
                        (d_359_v27_)[index72_] = not(not ((self).f17) or ((self).f17))
                        index73_ = default__.safeIndex(336, (d_359_v27_).length(0))
                        (d_359_v27_)[index73_] = ((d_341_v16_).f15) == (default__.safeModuloInt(250, 848))
                        d_360_v28_: _dafny.Seq
                        d_360_v28_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_341_v16_).f15])])
                        d_361_v29_: _dafny.MultiSet
                        d_361_v29_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(d_341_v16_).f15, (d_341_v16_).f15])])
                        rhs54_ = (d_360_v28_)[default__.safeIndex((d_341_v16_).f15, len(d_360_v28_))]
                        rhs55_ = (_dafny.MultiSet([(d_350_v22_).set(default__.safeIndex(d_330_v7_, len(d_350_v22_)), (d_341_v16_).f15), default__.fm16(263, (d_341_v16_).f15, not(not((self).f17)), d_330_v7_, globalState), _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_359_v27_)[default__.safeIndex(336, (d_359_v27_).length(0))]: (self).f17})) for d_362_i4_ in range(default__.abs(523))])])).isdisjoint(d_361_v29_)
                        rhs56_ = 172
                        d_350_v22_ = rhs54_
                        d_349_v21_ = rhs55_
                        d_330_v7_ = rhs56_
                        d_363_v30_: int
                        d_364_v31_: int
                        d_365_v32_: bool
                        out18_: int
                        out19_: int
                        out20_: bool
                        out18_, out19_, out20_ = (self).m6((d_359_v27_)[default__.safeIndex(336, (d_359_v27_).length(0))], globalState)
                        d_363_v30_ = out18_
                        d_364_v31_ = out19_
                        d_365_v32_ = out20_
                    elif source8_.is_DC4:
                        d_366___mcc_h1_ = source8_.cf5
                        d_367___mcc_h2_ = source8_.cf6
                        d_368___mcc_h3_ = source8_.cf7
                        d_369_cf7_ = d_368___mcc_h3_
                        d_370_cf6_ = d_367___mcc_h2_
                        d_371_cf5_ = d_366___mcc_h1_
                        d_372_v33_: _dafny.Set
                        d_372_v33_ = _dafny.Set({(self).f17, (self).f17})
                        d_373_v34_: _dafny.Map
                        d_373_v34_ = _dafny.Map({(self.f18 if (self).f17 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkrgtsv"))): default__.safeDivisionInt(d_330_v7_, len(d_372_v33_))})
                        d_373_v34_ = (d_373_v34_).set(self.f18, (d_341_v16_).f15)
                        d_374_v35_: _dafny.Map
                        d_374_v35_ = _dafny.Map({(self).f17: (self).f17})
                        d_370_cf6_ = (0) - ((d_330_v7_) * (len(d_374_v35_)))
                        (self).f18 = ((self.f18) + (_dafny.SeqWithoutIsStrInference([d_369_cf7_ for d_375_i5_ in range(default__.abs(-952))]))).set(default__.safeIndex((d_341_v16_).f15, len((self.f18) + (_dafny.SeqWithoutIsStrInference([d_369_cf7_ for d_376_i5_ in range(default__.abs(-952))])))), d_369_cf7_)
                        d_377_v36_: bool
                        d_377_v36_ = True
                        d_377_v36_ = ((-718) + (d_370_cf6_)) <= (447)
                    elif source8_.is_DC5:
                        d_378___mcc_h4_ = source8_.cf8
                        d_379___mcc_h5_ = source8_.cf9
                        d_380___mcc_h6_ = source8_.cf10
                        d_381_cf10_ = d_380___mcc_h6_
                        d_382_cf9_ = d_379___mcc_h5_
                        d_383_cf8_ = d_378___mcc_h4_
                        d_384_v37_: _dafny.Map
                        d_384_v37_ = _dafny.Map({d_330_v7_: (self).f17})
                        d_385_v38_: _dafny.Array
                        nw42_ = _dafny.Array(None, 21)
                        nw42_[int(0)] = d_343_v18_
                        nw42_[int(1)] = D1_DC2(d_342_v17_)
                        nw42_[int(2)] = d_343_v18_
                        nw42_[int(3)] = d_343_v18_
                        nw42_[int(4)] = d_343_v18_
                        nw42_[int(5)] = d_343_v18_
                        nw42_[int(6)] = d_343_v18_
                        nw42_[int(7)] = d_343_v18_
                        nw42_[int(8)] = d_343_v18_
                        nw42_[int(9)] = D1_DC2(_dafny.Map({len(default__.fm18((self).f17, (self).f17, (d_341_v16_).f15, 348, globalState)): len(d_384_v37_)}))
                        nw42_[int(10)] = (d_343_v18_ if not(True) else d_343_v18_)
                        nw42_[int(11)] = D1_DC2(d_342_v17_)
                        nw42_[int(12)] = d_343_v18_
                        nw42_[int(13)] = d_343_v18_
                        nw42_[int(14)] = d_343_v18_
                        nw42_[int(15)] = d_343_v18_
                        nw42_[int(16)] = d_343_v18_
                        nw42_[int(17)] = d_343_v18_
                        nw42_[int(18)] = d_343_v18_
                        nw42_[int(19)] = D1_DC2(d_342_v17_)
                        nw42_[int(20)] = default__.fm19(globalState)
                        d_385_v38_ = nw42_
                        index74_ = default__.safeIndex(907, (d_385_v38_).length(0))
                        (d_385_v38_)[index74_] = d_343_v18_
                        index75_ = default__.safeIndex(907, (d_385_v38_).length(0))
                        (d_385_v38_)[index75_] = d_343_v18_
                        d_386_v39_: C0
                        nw43_ = C0()
                        nw43_.ctor__(d_330_v7_)
                        d_386_v39_ = nw43_
                        d_387_v40_: bool
                        d_387_v40_ = False
                        d_388_v41_: _dafny.Seq
                        d_388_v41_ = _dafny.SeqWithoutIsStrInference([d_387_v40_])
                        d_387_v40_ = (False) or ((d_388_v41_)[default__.safeIndex(d_381_cf10_, len(d_388_v41_))])
                        d_389_v42_: _dafny.Seq
                        d_389_v42_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_341_v16_).f15), (d_341_v16_).f15])
                        d_330_v7_ = len((d_389_v42_) + ((d_389_v42_) + (d_389_v42_)))
                    elif True:
                        d_390___mcc_h7_ = source8_.cf3
                        d_391_cf3_ = d_390___mcc_h7_
                        index76_ = default__.safeIndex(731, ((self).f14).length(0))
                        ((self).f14)[index76_] = (d_330_v7_) + ((d_341_v16_).f15)
                        index77_ = default__.safeIndex(731, ((self).f14).length(0))
                        ((self).f14)[index77_] = (d_341_v16_).f15
                        d_392_v43_: _dafny.Array
                        nw44_ = _dafny.Array(False, 25)
                        d_392_v43_ = nw44_
                        index78_ = default__.safeIndex(993, (d_392_v43_).length(0))
                        (d_392_v43_)[index78_] = (self).f17
                        d_393_v44_: bool
                        d_393_v44_ = False
                        d_394_v45_: _dafny.MultiSet
                        d_394_v45_ = _dafny.MultiSet([d_393_v44_, not((self).f17)])
                        index79_ = default__.safeIndex(993, (d_392_v43_).length(0))
                        rhs57_ = not((((d_394_v45_).cardinality) + ((d_341_v16_).f15)) != ((d_341_v16_).f15))
                        rhs58_ = ((d_341_v16_).f15) < (len(self.f18))
                        lhs34_ = d_392_v43_
                        lhs35_ = default__.safeIndex(993, (d_392_v43_).length(0))
                        lhs34_[lhs35_] = rhs57_
                        d_393_v44_ = rhs58_
                        index80_ = default__.safeIndex(993, (d_392_v43_).length(0))
                        (d_392_v43_)[index80_] = not(True)
                        d_393_v44_ = (d_392_v43_)[default__.safeIndex(993, (d_392_v43_).length(0))]
                    d_395_v46_: _dafny.Seq
                    d_395_v46_ = _dafny.SeqWithoutIsStrInference([(self).f17, not((self).f17)])
                    d_396_v47_: _dafny.Map
                    d_396_v47_ = _dafny.Map({(d_341_v16_).f15: (_dafny.SeqWithoutIsStrInference([(self).f17, (self).f17])) + (d_395_v46_)})
                    d_397_v48_: _dafny.Set
                    d_397_v48_ = _dafny.Set({d_395_v46_, _dafny.SeqWithoutIsStrInference([(self).f17, (self).f17])})
                    d_398_v49_: _dafny.MultiSet
                    d_398_v49_ = _dafny.MultiSet([True])
                    d_399_v50_: _dafny.Seq
                    d_399_v50_ = _dafny.SeqWithoutIsStrInference([len(d_397_v48_), (d_398_v49_).cardinality])
                    d_396_v47_ = (d_396_v47_).set(len(_dafny.Map({_dafny.Set({(self).f17, (self).f17}): (d_399_v50_)[default__.safeIndex(d_330_v7_, len(d_399_v50_))]})), d_395_v46_)
                    d_400_v51_: D0
                    d_400_v51_ = D0_DC0((self).f17)
                    source9_ = d_400_v51_
                    if source9_.is_DC1:
                        d_401___mcc_h8_ = source9_.cf1
                        d_402___mcc_h9_ = source9_.cf2
                        d_403_cf2_ = d_402___mcc_h9_
                        d_404_cf1_ = d_401___mcc_h8_
                        index81_ = default__.safeIndex(810, ((self).f14).length(0))
                        ((self).f14)[index81_] = default__.safeDivisionInt((d_341_v16_).f15, d_403_cf2_)
                        index82_ = default__.safeIndex(810, ((self).f14).length(0))
                        ((self).f14)[index82_] = d_403_cf2_
                        d_405_v52_: _dafny.Array
                        def lambda9_(d_406_v16_):
                            def lambda10_(d_407_i6_):
                                return (d_407_i6_) * ((d_406_v16_).f15)

                            return lambda10_

                        init4_ = lambda9_(d_341_v16_)
                        nw45_ = _dafny.Array(None, 19)
                        for i0_4_ in range(nw45_.length(0)):
                            nw45_[i0_4_] = init4_(i0_4_)
                        d_405_v52_ = nw45_
                        def lambda11_(d_408_v16_):
                            def lambda12_(d_409_i7_):
                                return (d_409_i7_) * ((d_408_v16_).f15)

                            return lambda12_

                        init5_ = lambda11_(d_341_v16_)
                        nw46_ = _dafny.Array(None, 19)
                        for i0_5_ in range(nw46_.length(0)):
                            nw46_[i0_5_] = init5_(i0_5_)
                        d_405_v52_ = nw46_
                        d_410_v53_: _dafny.Map
                        d_410_v53_ = _dafny.Map({(self).f17: d_404_cf1_})
                        d_411_v54_: _dafny.Map
                        d_411_v54_ = _dafny.Map({False: 994})
                        d_412_v55_: _dafny.Set
                        d_412_v55_ = _dafny.Set({(d_341_v16_).f15, (d_341_v16_).f15})
                        d_413_v56_: _dafny.Array
                        nw47_ = _dafny.Array(None, 25)
                        nw47_[int(0)] = d_404_cf1_
                        nw47_[int(1)] = (self).f17
                        nw47_[int(2)] = d_404_cf1_
                        nw47_[int(3)] = d_404_cf1_
                        nw47_[int(4)] = not(d_404_cf1_)
                        nw47_[int(5)] = (self).f17
                        nw47_[int(6)] = d_404_cf1_
                        nw47_[int(7)] = False
                        nw47_[int(8)] = d_404_cf1_
                        nw47_[int(9)] = (self).f17
                        nw47_[int(10)] = (self).f17
                        nw47_[int(11)] = False
                        nw47_[int(12)] = (self).f17
                        nw47_[int(13)] = False
                        nw47_[int(14)] = d_404_cf1_
                        nw47_[int(15)] = True
                        nw47_[int(16)] = (self).f17
                        nw47_[int(17)] = d_404_cf1_
                        nw47_[int(18)] = d_404_cf1_
                        nw47_[int(19)] = True
                        nw47_[int(20)] = (d_395_v46_)[default__.safeIndex(len(d_411_v54_), len(d_395_v46_))]
                        nw47_[int(21)] = (self).fm14((d_341_v16_).f15, d_404_cf1_, (d_341_v16_).f15, len(d_412_v55_), globalState)
                        nw47_[int(22)] = (self).f17
                        nw47_[int(23)] = False
                        nw47_[int(24)] = d_404_cf1_
                        d_413_v56_ = nw47_
                        d_414_v57_: _dafny.Seq
                        d_414_v57_ = _dafny.SeqWithoutIsStrInference([d_410_v53_])
                        d_415_v58_: _dafny.Map
                        d_415_v58_ = _dafny.Map({d_413_v56_: (d_414_v57_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_416_i8_ in range(default__.abs(-257))])), len(d_414_v57_))]})
                        d_417_v59_: _dafny.Map
                        d_417_v59_ = d_410_v53_
                        d_418_v60_: _dafny.Array
                        nw48_ = _dafny.Array(None, 10)
                        nw48_[int(0)] = d_410_v53_
                        nw48_[int(1)] = _dafny.Map({not(d_404_cf1_): not((self).f17)})
                        nw48_[int(2)] = (d_410_v53_) | (d_410_v53_)
                        nw48_[int(3)] = ((d_415_v58_)[d_413_v56_] if (d_413_v56_) in (d_415_v58_) else (d_410_v53_).set(d_404_cf1_, ((d_410_v53_)[(self).f17] if ((self).f17) in (d_410_v53_) else d_404_cf1_)))
                        nw48_[int(4)] = d_410_v53_
                        nw48_[int(5)] = (d_417_v59_)
                        nw48_[int(6)] = d_410_v53_
                        nw48_[int(7)] = d_410_v53_
                        nw48_[int(8)] = d_410_v53_
                        nw48_[int(9)] = (default__.fm20(d_404_cf1_, ((self).f14)[default__.safeIndex(810, ((self).f14).length(0))], d_404_cf1_, globalState)) | (d_410_v53_)
                        d_418_v60_ = nw48_
                        index83_ = default__.safeIndex(679, (d_418_v60_).length(0))
                        (d_418_v60_)[index83_] = d_410_v53_
                        index84_ = default__.safeIndex(679, (d_418_v60_).length(0))
                        (d_418_v60_)[index84_] = d_410_v53_
                        d_405_v52_ = d_405_v52_
                    elif True:
                        d_419___mcc_h10_ = source9_.cf0
                        d_420_cf0_ = d_419___mcc_h10_
                        d_421_v61_: int
                        d_422_v62_: int
                        d_423_v63_: bool
                        out21_: int
                        out22_: int
                        out23_: bool
                        out21_, out22_, out23_ = (self).m6((d_395_v46_)[default__.safeIndex((d_341_v16_).f15, len(d_395_v46_))], globalState)
                        d_421_v61_ = out21_
                        d_422_v62_ = out22_
                        d_423_v63_ = out23_
                        d_399_v50_ = (d_399_v50_) + (d_399_v50_)
                        d_424_v64_: str
                        d_424_v64_ = _dafny.CodePoint('h')
                        d_425_v65_: _dafny.Map
                        d_425_v65_ = _dafny.Map({d_421_v61_: d_424_v64_})
                        d_426_v66_: _dafny.Map
                        d_426_v66_ = _dafny.Map({(self).f14: d_425_v65_})
                        d_426_v66_ = (_dafny.Map({(self).f14: d_425_v65_})) | (d_426_v66_)
                        d_427_v67_: _dafny.Array
                        def lambda13_(d_428_v63_):
                            def lambda14_(d_429_i9_):
                                return not(d_428_v63_)

                            return lambda14_

                        init6_ = lambda13_(d_423_v63_)
                        nw49_ = _dafny.Array(None, 4)
                        for i0_6_ in range(nw49_.length(0)):
                            nw49_[i0_6_] = init6_(i0_6_)
                        d_427_v67_ = nw49_
                        d_430_v68_: _dafny.Map
                        d_430_v68_ = _dafny.Map({d_427_v67_: _dafny.Set({(self).fm6(globalState)})})
                        d_431_v69_: _dafny.Set
                        d_431_v69_ = _dafny.Set({d_423_v63_, True, d_423_v63_})
                        d_430_v68_ = (d_430_v68_).set(d_427_v67_, d_431_v69_)
                    pass
            pass
        d_432_v70_: bool
        d_432_v70_ = False
        d_432_v70_ = d_432_v70_
        d_433_v71_: int
        d_434_v72_: int
        d_435_v73_: bool
        out24_: int
        out25_: int
        out26_: bool
        out24_, out25_, out26_ = (self).m6((d_330_v7_) < (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_436_i10_ in range(default__.abs(-78))]))), globalState)
        d_433_v71_ = out24_
        d_434_v72_ = out25_
        d_435_v73_ = out26_
        d_437_v74_: C0
        nw50_ = C0()
        nw50_.ctor__((self).fm15(globalState))
        d_437_v74_ = nw50_

    def m6(self, p0, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        d_438_v0_: int
        d_438_v0_ = 848
        d_439_v1_: _dafny.Map
        d_439_v1_ = _dafny.Map({d_438_v0_: p0})
        d_440_v2_: _dafny.Map
        d_440_v2_ = _dafny.Map({not(p0): d_439_v1_})
        d_441_v4_: _dafny.Map
        def iife47_():
            coll27_ = _dafny.Map()
            compr_27_: int
            for compr_27_ in (d_439_v1_).keys.Elements:
                d_442_v3_: int = compr_27_
                if (d_442_v3_) in (d_439_v1_):
                    coll27_[(d_442_v3_) + (-153)] = (self).f17
            return _dafny.Map(coll27_)
        d_441_v4_ = _dafny.Map({d_438_v0_: len(((d_440_v2_)[(self).f17] if ((self).f17) in (d_440_v2_) else iife47_()
        ))})
        d_443_v5_: D0
        d_443_v5_ = D0_DC1(p0, d_438_v0_)
        if default__.fm0(d_441_v4_, (_dafny.MultiSet([(d_443_v5_).cf1])).cardinality, globalState):
            d_438_v0_ = d_438_v0_
            d_444_v6_: _dafny.Map
            d_444_v6_ = _dafny.Map({self.f18: p0})
            d_445_v7_: _dafny.Map
            d_445_v7_ = _dafny.Map({self.f18: d_444_v6_})
            d_446_v8_: int
            out27_: int
            out27_ = default__.m0(((d_445_v7_)[self.f18] if (self.f18) in (d_445_v7_) else _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_447_i0_ in range(default__.abs(498))]): (self).f17})), globalState)
            d_446_v8_ = out27_
            rhs59_ = d_446_v8_
            rhs60_ = p0
            r1 = rhs59_
            r2 = rhs60_
            d_448_v9_: _dafny.Array
            nw51_ = _dafny.Array(False, 29)
            d_448_v9_ = nw51_
            d_449_v10_: _dafny.MultiSet
            d_449_v10_ = _dafny.MultiSet([d_448_v9_])
            r2 = not((default__.fm0(d_441_v4_, (self).fm7(107, (d_449_v10_).cardinality, (self).f17, globalState), globalState) if p0 else (self).f17))
            d_450_v11_: C0
            nw52_ = C0()
            nw52_.ctor__(d_438_v0_)
            d_450_v11_ = nw52_
        elif True:
            d_451_v12_: _dafny.Array
            nw53_ = _dafny.Array(False, 26)
            d_451_v12_ = nw53_
            index85_ = default__.safeIndex(254, (d_451_v12_).length(0))
            (d_451_v12_)[index85_] = p0
            d_452_v13_: D4
            d_452_v13_ = D4_DC10((self).f14)
            d_453_v14_: _dafny.Array
            nw54_ = _dafny.Array(None, 17)
            nw54_[int(0)] = (self).f14
            nw54_[int(1)] = (self).f14
            nw54_[int(2)] = (d_452_v13_).cf20
            nw54_[int(3)] = (self).f14
            nw54_[int(4)] = (self).f14
            nw54_[int(5)] = (self).f14
            nw54_[int(6)] = (self).f14
            nw54_[int(7)] = (self).f14
            nw54_[int(8)] = (self).f14
            nw54_[int(9)] = (self).f14
            nw54_[int(10)] = (self).f14
            nw54_[int(11)] = (self).f14
            nw54_[int(12)] = (self).f14
            nw54_[int(13)] = (self).f14
            nw54_[int(14)] = (self).f14
            nw54_[int(15)] = (self).f14
            nw54_[int(16)] = (self).f14
            d_453_v14_ = nw54_
            d_454_v15_: _dafny.Array
            nw55_ = _dafny.Array(_dafny.Seq({}), 24)
            d_454_v15_ = nw55_
            d_455_v16_: _dafny.Seq
            d_455_v16_ = _dafny.SeqWithoutIsStrInference([(0) - (d_438_v0_), d_438_v0_, d_438_v0_])
            index86_ = default__.safeIndex(298, (d_454_v15_).length(0))
            (d_454_v15_)[index86_] = (d_455_v16_) + (d_455_v16_)
            d_456_v17_: _dafny.Array
            nw56_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
            d_456_v17_ = nw56_
            index87_ = default__.safeIndex(608, (d_456_v17_).length(0))
            (d_456_v17_)[index87_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okriy"))
            index88_ = default__.safeIndex(254, (d_451_v12_).length(0))
            index89_ = default__.safeIndex(298, (d_454_v15_).length(0))
            index90_ = default__.safeIndex(608, (d_456_v17_).length(0))
            rhs61_ = False
            rhs62_ = d_453_v14_
            rhs63_ = (d_438_v0_) * (d_438_v0_)
            rhs64_ = d_455_v16_
            rhs65_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_457_i1_ in range(default__.abs(300))])) + (self.f18)
            lhs36_ = d_451_v12_
            lhs37_ = default__.safeIndex(254, (d_451_v12_).length(0))
            lhs38_ = d_454_v15_
            lhs39_ = default__.safeIndex(298, (d_454_v15_).length(0))
            lhs40_ = d_456_v17_
            lhs41_ = default__.safeIndex(608, (d_456_v17_).length(0))
            lhs36_[lhs37_] = rhs61_
            d_453_v14_ = rhs62_
            r1 = rhs63_
            lhs38_[lhs39_] = rhs64_
            lhs40_[lhs41_] = rhs65_
            d_458_v18_: D1
            d_458_v18_ = D1_DC3(d_438_v0_)
            pat_let_tv14_ = d_438_v0_
            def iife48_(_pat_let10_0):
                def iife49_(d_459_dt__update__tmp_h0_):
                    def iife50_(_pat_let11_0):
                        def iife51_(d_460_dt__update_hcf4_h0_):
                            return D1_DC3(d_460_dt__update_hcf4_h0_)
                        return iife51_(_pat_let11_0)
                    return iife50_(pat_let_tv14_)
                return iife49_(_pat_let10_0)
            source10_ = iife48_(d_458_v18_)
            if source10_.is_DC3:
                d_461___mcc_h0_ = source10_.cf4
                d_462_cf4_ = d_461___mcc_h0_
                d_463_v19_: C0
                nw57_ = C0()
                nw57_.ctor__(default__.safeModuloInt((0) - (d_438_v0_), d_462_cf4_))
                d_463_v19_ = nw57_
                d_464_v20_: D1
                d_464_v20_ = D1_DC2(d_441_v4_)
                d_464_v20_ = d_464_v20_
                d_462_cf4_ = (d_463_v19_).f15
                d_451_v12_ = d_451_v12_
            elif source10_.is_DC4:
                d_465___mcc_h1_ = source10_.cf5
                d_466___mcc_h2_ = source10_.cf6
                d_467___mcc_h3_ = source10_.cf7
                d_468_cf7_ = d_467___mcc_h3_
                d_469_cf6_ = d_466___mcc_h2_
                d_470_cf5_ = d_465___mcc_h1_
                d_471_v21_: _dafny.Set
                d_471_v21_ = _dafny.Set({(0) - ((self).fm7(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqslcms"))).set(default__.safeIndex(d_469_cf6_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqslcms")))), d_468_cf7_)), d_438_v0_, (d_451_v12_)[default__.safeIndex(254, (d_451_v12_).length(0))], globalState)), d_469_cf6_, d_469_cf6_})
                d_472_v22_: _dafny.Seq
                d_472_v22_ = _dafny.SeqWithoutIsStrInference([d_471_v21_, d_471_v21_, d_471_v21_, d_471_v21_, d_471_v21_])
                d_473_v23_: _dafny.Map
                d_473_v23_ = _dafny.Map({True: True})
                d_474_v24_: _dafny.Map
                d_474_v24_ = _dafny.Map({(self).fm14(d_469_cf6_, (d_451_v12_)[default__.safeIndex(254, (d_451_v12_).length(0))], len(d_473_v23_), d_469_cf6_, globalState): default__.fm21(not(p0), globalState)})
                d_472_v22_ = ((d_474_v24_)[(d_451_v12_)[default__.safeIndex(254, (d_451_v12_).length(0))]] if ((d_451_v12_)[default__.safeIndex(254, (d_451_v12_).length(0))]) in (d_474_v24_) else d_472_v22_)
                d_475_v25_: _dafny.Map
                d_475_v25_ = _dafny.Map({p0: d_438_v0_})
                d_476_v26_: D1
                d_476_v26_ = D1_DC2((_dafny.Map({(_dafny.MultiSet([d_468_cf7_, d_468_cf7_])).cardinality: len(d_475_v25_)})).set(len(_dafny.SeqWithoutIsStrInference([d_438_v0_, 113])), len((d_456_v17_)[default__.safeIndex(608, (d_456_v17_).length(0))])))
                d_477_v27_: _dafny.MultiSet
                d_477_v27_ = _dafny.MultiSet([d_476_v26_])
                index91_ = default__.safeIndex(608, (d_456_v17_).length(0))
                (d_456_v17_)[index91_] = default__.fm17((d_477_v27_) | (_dafny.MultiSet([d_476_v26_, D1_DC2(_dafny.Map({d_469_cf6_: d_438_v0_})), d_476_v26_, d_476_v26_, d_476_v26_])), globalState)
                index92_ = default__.safeIndex(254, (d_451_v12_).length(0))
                (d_451_v12_)[index92_] = p0
                d_478_v28_: _dafny.Map
                d_478_v28_ = _dafny.Map({d_439_v1_: _dafny.SeqWithoutIsStrInference([(d_439_v1_).set(len(_dafny.SeqWithoutIsStrInference([(self).f17, p0, (self).f17])), (d_451_v12_)[default__.safeIndex(254, (d_451_v12_).length(0))]) for d_479_i2_ in range(default__.abs(365))])})
                d_480_v29_: _dafny.Seq
                d_480_v29_ = _dafny.SeqWithoutIsStrInference([d_439_v1_, d_439_v1_])
                d_481_v30_: _dafny.Map
                d_481_v30_ = _dafny.Map({(((d_478_v28_)[d_439_v1_] if (d_439_v1_) in (d_478_v28_) else _dafny.SeqWithoutIsStrInference([d_439_v1_ for d_482_i3_ in range(default__.abs(613))]))) != (d_480_v29_): d_443_v5_})
                d_481_v30_ = d_481_v30_
            elif source10_.is_DC5:
                d_483___mcc_h4_ = source10_.cf8
                d_484___mcc_h5_ = source10_.cf9
                d_485___mcc_h6_ = source10_.cf10
                d_486_cf10_ = d_485___mcc_h6_
                d_487_cf9_ = d_484___mcc_h5_
                d_488_cf8_ = d_483___mcc_h4_
                d_489_v31_: C0
                nw58_ = C0()
                nw58_.ctor__(d_438_v0_)
                d_489_v31_ = nw58_
                d_489_v31_ = d_489_v31_
                d_490_v32_: C0
                nw59_ = C0()
                nw59_.ctor__(523)
                d_490_v32_ = nw59_
                d_491_v33_: C0
                nw60_ = C0()
                nw60_.ctor__((d_489_v31_).f15)
                d_491_v33_ = nw60_
                r0 = d_486_cf10_
            elif True:
                d_492___mcc_h7_ = source10_.cf3
                d_493_cf3_ = d_492___mcc_h7_
                d_438_v0_ = d_438_v0_
                d_438_v0_ = ((d_438_v0_ if p0 else d_438_v0_)) - (d_438_v0_)
                d_494_v34_: str
                d_494_v34_ = _dafny.CodePoint('b')
                d_495_v35_: _dafny.Map
                d_495_v35_ = _dafny.Map({d_438_v0_: d_494_v34_})
                d_496_v36_: _dafny.Map
                d_496_v36_ = _dafny.Map({True: d_495_v35_})
                d_496_v36_ = (d_496_v36_) | (_dafny.Map({(self).f17: d_495_v35_}))
                d_497_v37_: _dafny.Array
                nw61_ = _dafny.Array(_dafny.Array(None, 0), 5)
                d_497_v37_ = nw61_
                index93_ = default__.safeIndex(898, (d_497_v37_).length(0))
                (d_497_v37_)[index93_] = d_456_v17_
                d_498_v38_: _dafny.MultiSet
                d_498_v38_ = _dafny.MultiSet([p0])
                d_499_v39_: _dafny.MultiSet
                d_499_v39_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_494_v34_ for d_500_i4_ in range(default__.abs(221))])])
                d_501_v40_: D2
                d_501_v40_ = D2_DC7((d_456_v17_)[default__.safeIndex(608, (d_456_v17_).length(0))], d_438_v0_, ((d_499_v39_)[(d_456_v17_)[default__.safeIndex(608, (d_456_v17_).length(0))]] if ((d_456_v17_)[default__.safeIndex(608, (d_456_v17_).length(0))]) in (d_499_v39_) else 204), (self).f17)
                d_502_v41_: _dafny.Map
                d_502_v41_ = _dafny.Map({d_494_v34_: d_441_v4_})
                d_503_v42_: D1
                d_503_v42_ = D1_DC2(((d_502_v41_)[d_494_v34_] if (d_494_v34_) in (d_502_v41_) else _dafny.Map({d_438_v0_: len(d_455_v16_)})))
                d_504_v43_: _dafny.MultiSet
                d_504_v43_ = _dafny.MultiSet([d_503_v42_, d_503_v42_, default__.fm19(globalState)])
                d_505_v44_: D1
                d_505_v44_ = D1_DC5(default__.fm17(d_504_v43_, globalState), d_451_v12_, d_438_v0_)
                d_506_v45_: _dafny.Set
                d_506_v45_ = _dafny.Set({False})
                d_507_v46_: _dafny.Seq
                d_507_v46_ = _dafny.SeqWithoutIsStrInference([d_506_v45_, d_506_v45_])
                index94_ = default__.safeIndex(898, (d_497_v37_).length(0))
                rhs66_ = d_456_v17_
                rhs67_ = (d_501_v40_).cf14
                rhs68_ = (d_505_v44_).cf10
                rhs69_ = (d_498_v38_).set((d_438_v0_) < (len((d_507_v46_)[default__.safeIndex(879, len(d_507_v46_))])), default__.abs(d_438_v0_))
                rhs70_ = d_438_v0_
                lhs42_ = d_497_v37_
                lhs43_ = default__.safeIndex(898, (d_497_v37_).length(0))
                lhs42_[lhs43_] = rhs66_
                r1 = rhs67_
                r1 = rhs68_
                d_498_v38_ = rhs69_
                d_438_v0_ = rhs70_
            d_451_v12_ = d_451_v12_
            d_508_v47_: _dafny.Seq
            d_508_v47_ = _dafny.SeqWithoutIsStrInference([d_451_v12_, d_451_v12_])
            d_509_v48_: str
            d_509_v48_ = _dafny.CodePoint('m')
            d_510_v49_: _dafny.Array
            nw62_ = _dafny.Array(None, 19)
            nw62_[int(0)] = (d_508_v47_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_509_v48_])), len(d_508_v47_))]
            nw62_[int(1)] = d_451_v12_
            nw62_[int(2)] = (d_451_v12_ if p0 else d_451_v12_)
            nw62_[int(3)] = d_451_v12_
            nw62_[int(4)] = d_451_v12_
            nw62_[int(5)] = d_451_v12_
            nw62_[int(6)] = d_451_v12_
            nw62_[int(7)] = d_451_v12_
            nw62_[int(8)] = d_451_v12_
            nw62_[int(9)] = d_451_v12_
            nw62_[int(10)] = (d_508_v47_)[default__.safeIndex(d_438_v0_, len(d_508_v47_))]
            nw62_[int(11)] = d_451_v12_
            nw62_[int(12)] = d_451_v12_
            nw62_[int(13)] = d_451_v12_
            nw62_[int(14)] = d_451_v12_
            nw62_[int(15)] = d_451_v12_
            nw62_[int(16)] = d_451_v12_
            nw62_[int(17)] = d_451_v12_
            nw62_[int(18)] = d_451_v12_
            d_510_v49_ = nw62_
            index95_ = default__.safeIndex(900, (d_510_v49_).length(0))
            (d_510_v49_)[index95_] = d_451_v12_
            d_511_v50_: _dafny.Map
            d_511_v50_ = _dafny.Map({not((p0 if (self).f17 else (self).f17)): (self).f17})
            index96_ = default__.safeIndex(900, (d_510_v49_).length(0))
            index97_ = default__.safeIndex(254, (d_451_v12_).length(0))
            rhs71_ = (d_508_v47_)[default__.safeIndex(d_438_v0_, len(d_508_v47_))]
            rhs72_ = ((d_511_v50_)[(p0) and ((self).f17)] if ((p0) and ((self).f17)) in (d_511_v50_) else p0)
            lhs44_ = d_510_v49_
            lhs45_ = default__.safeIndex(900, (d_510_v49_).length(0))
            lhs46_ = d_451_v12_
            lhs47_ = default__.safeIndex(254, (d_451_v12_).length(0))
            lhs44_[lhs45_] = rhs71_
            lhs46_[lhs47_] = rhs72_
            index98_ = default__.safeIndex(254, (d_451_v12_).length(0))
            (d_451_v12_)[index98_] = (d_438_v0_) > (d_438_v0_)
        d_512_i5_: int
        d_512_i5_ = 0
        with _dafny.label("6"):
            while p0:
                with _dafny.c_label("6"):
                    if (d_512_i5_) >= (100):
                        raise _dafny.Break("6")
                    d_512_i5_ = (d_512_i5_) + (1)
                    r2 = (self).f17
                    r2 = (self).f17
                    (self).f18 = ((self.f18 if False else self.f18)) + (self.f18)
                    d_513_v51_: _dafny.Seq
                    d_513_v51_ = _dafny.SeqWithoutIsStrInference([d_438_v0_, d_438_v0_, (d_438_v0_ if False else (0) - (d_438_v0_))])
                    rhs73_ = (d_438_v0_) - (d_438_v0_)
                    rhs74_ = d_438_v0_
                    rhs75_ = (self).f17
                    rhs76_ = (d_513_v51_) + (d_513_v51_)
                    d_438_v0_ = rhs73_
                    r0 = rhs74_
                    r2 = rhs75_
                    d_513_v51_ = rhs76_
                    pass
            pass
        d_514_v52_: _dafny.Array
        nw63_ = _dafny.Array(None, 4)
        nw63_[int(0)] = p0
        nw63_[int(1)] = (self).f17
        nw63_[int(2)] = (self).f17
        nw63_[int(3)] = (self).f17
        d_514_v52_ = nw63_
        d_515_v53_: str
        d_515_v53_ = _dafny.CodePoint('o')
        d_516_v54_: D1
        d_516_v54_ = D1_DC4(d_514_v52_, d_438_v0_, d_515_v53_)
        d_517_v55_: _dafny.MultiSet
        d_517_v55_ = _dafny.MultiSet([d_514_v52_, d_514_v52_])
        d_518_v56_: _dafny.Seq
        d_518_v56_ = _dafny.SeqWithoutIsStrInference([d_517_v55_, d_517_v55_, d_517_v55_])
        r0 = ((d_516_v54_).cf6) - (((d_518_v56_)[default__.safeIndex(531, len(d_518_v56_))]).cardinality)
        d_519_v57_: _dafny.Array
        def lambda15_(d_520_i6_):
            return D0_DC0(True)

        init7_ = lambda15_
        nw64_ = _dafny.Array(None, 19)
        for i0_7_ in range(nw64_.length(0)):
            nw64_[i0_7_] = init7_(i0_7_)
        d_519_v57_ = nw64_
        d_521_v58_: D0
        d_521_v58_ = D0_DC0((self).f17)
        index99_ = default__.safeIndex(989, (d_519_v57_).length(0))
        (d_519_v57_)[index99_] = d_521_v58_
        index100_ = default__.safeIndex(989, (d_519_v57_).length(0))
        (d_519_v57_)[index100_] = d_521_v58_
        r2 = (d_515_v53_) not in (self.f18)
        index101_ = default__.safeIndex(292, ((self).f14).length(0))
        ((self).f14)[index101_] = d_438_v0_
        index102_ = default__.safeIndex(292, ((self).f14).length(0))
        ((self).f14)[index102_] = -219
        d_522_v59_: _dafny.Set
        d_522_v59_ = _dafny.Set({(self).f17})
        d_523_v60_: _dafny.MultiSet
        d_523_v60_ = _dafny.MultiSet([p0, (d_522_v59_).isdisjoint(d_522_v59_), (True) or (default__.fm0(d_441_v4_, ((self).f14)[default__.safeIndex(292, ((self).f14).length(0))], globalState))])
        r0 = ((d_523_v60_)[True] if (True) in (d_523_v60_) else 870)
        r1 = ((self).f14)[default__.safeIndex(292, ((self).f14).length(0))]
        r2 = p0
        return r0, r1, r2

    @property
    def f17(self):
        return self._f17

class C2(T0):
    def  __init__(self):
        self._f13: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f13(self):
        return self._f13
    @f13.setter
    def f13(self, value):
        self._f13 = value
    def ctor__(self, f13):
        (self).f13 = f13

    def fm4(self, p0, p1, p2, p3, globalState):
        return self.f13

    def fm5(self, globalState):
        source11_ = D5_DC13()
        if source11_.is_DC13:
            return D1_DC3(614)
        elif source11_.is_DC14:
            d_524___mcc_h0_ = source11_.cf23
            d_525_cf23_ = d_524___mcc_h0_
            return D1_DC3(504)
        elif source11_.is_DC12:
            d_526___mcc_h1_ = source11_.cf22
            d_527_cf22_ = d_526___mcc_h1_
            return D1_DC3(((d_527_cf22_)[self.f13] if (self.f13) in (d_527_cf22_) else len(_dafny.Map({self.f13: self.f13}))))
        elif True:
            d_528___mcc_h2_ = source11_.cf24
            d_529_cf24_ = d_528___mcc_h2_
            return D1_DC3(len(_dafny.Set({self.f13})))

    def fm29(self, p0, p1, p2, globalState):
        def iife52_():
            coll28_ = _dafny.Map()
            compr_28_: int
            for compr_28_ in (_dafny.SeqWithoutIsStrInference([214, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqw")))])).Elements:
                d_530_v0_: int = compr_28_
                if (d_530_v0_) in (_dafny.SeqWithoutIsStrInference([214, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqw")))])):
                    coll28_[(d_530_v0_) + (24)] = (0) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([self.f13])): -428})))
            return _dafny.Map(coll28_)
        if (_dafny.Map({len(_dafny.Map({self.f13: len(_dafny.Map({248: self.f13}))})): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eaogdxqm")))})) != (iife52_()
        ):
            return D5_DC13()
        elif True:
            return D5_DC13()

    def fm30(self, p0, p1, p2, globalState):
        return self.f13

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_531_v0_: _dafny.Map
        d_531_v0_ = _dafny.Map({self.f13: self.f13})
        hi3_ = (0) - ((0) - (p0))
        for d_532_i0_ in range((len(d_531_v0_)) + (p0), hi3_):
            if True:
                d_533_v1_: int
                d_533_v1_ = 661
                d_534_v2_: str
                d_534_v2_ = _dafny.CodePoint('v')
                d_535_v3_: _dafny.Map
                d_535_v3_ = _dafny.Map({d_534_v2_: d_532_i0_})
                d_536_v4_: _dafny.Seq
                d_536_v4_ = _dafny.SeqWithoutIsStrInference([len(d_535_v3_)])
                d_533_v1_ = default__.safeDivisionInt((0) - (len(d_536_v4_)), (D2_DC7(_dafny.SeqWithoutIsStrInference([d_534_v2_, d_534_v2_]), p0, d_533_v1_, False)).cf14)
                d_537_v5_: _dafny.Map
                d_537_v5_ = _dafny.Map({p0: p0})
                d_538_v6_: _dafny.Seq
                d_538_v6_ = _dafny.SeqWithoutIsStrInference([self.f13])
                d_539_v7_: _dafny.Map
                d_539_v7_ = _dafny.Map({default__.fm31(globalState): (d_538_v6_)[default__.safeIndex(d_532_i0_, len(d_538_v6_))]})
                d_540_v8_: _dafny.Map
                d_540_v8_ = _dafny.Map({((d_537_v5_)[d_532_i0_] if (d_532_i0_) in (d_537_v5_) else default__.fm31(globalState)): d_539_v7_})
                d_540_v8_ = (d_540_v8_).set(d_532_i0_, d_539_v7_)
                d_541_v9_: _dafny.Seq
                d_541_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ioqxh"))
                d_542_v10_: _dafny.Map
                d_542_v10_ = _dafny.Map({d_541_v9_: not(self.f13)})
                d_543_v11_: int
                out28_: int
                out28_ = default__.m0(d_542_v10_, globalState)
                d_543_v11_ = out28_
                d_544_v12_: _dafny.Array
                nw65_ = _dafny.Array(int(0), 10)
                d_544_v12_ = nw65_
                index103_ = default__.safeIndex(636, (d_544_v12_).length(0))
                (d_544_v12_)[index103_] = d_543_v11_
                index104_ = default__.safeIndex(636, (d_544_v12_).length(0))
                (d_544_v12_)[index104_] = (0) - ((0) - (d_532_i0_))
                (self).f13 = self.f13
            elif True:
                (self).f13 = self.f13
                d_545_v13_: _dafny.Seq
                d_545_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnbb"))
                d_546_v14_: _dafny.Array
                nw66_ = _dafny.Array(False, 17)
                d_546_v14_ = nw66_
                d_547_v15_: str
                d_547_v15_ = _dafny.CodePoint('j')
                d_548_v16_: D1
                d_548_v16_ = D1_DC4(d_546_v14_, p0, d_547_v15_)
                d_549_v17_: _dafny.Set
                d_549_v17_ = _dafny.Set({self.f13})
                d_550_v18_: _dafny.Map
                d_550_v18_ = _dafny.Map({p0: d_549_v17_})
                d_551_v19_: _dafny.Map
                d_551_v19_ = _dafny.Map({False: d_549_v17_})
                d_552_v20_: D1
                d_552_v20_ = D1_DC5(d_545_v13_, (d_548_v16_).cf5, len(((d_550_v18_)[p0] if (p0) in (d_550_v18_) else ((d_551_v19_)[self.f13] if (self.f13) in (d_551_v19_) else d_549_v17_))))
                d_553_v21_: _dafny.Map
                d_553_v21_ = _dafny.Map({self.f13: len(d_531_v0_)})
                d_554_v22_: _dafny.Seq
                d_554_v22_ = _dafny.SeqWithoutIsStrInference([d_545_v13_])
                pat_let_tv15_ = d_545_v13_
                d_555_v23_: _dafny.Array
                nw67_ = _dafny.Array(None, 23)
                nw67_[int(0)] = d_552_v20_
                nw67_[int(1)] = d_552_v20_
                def iife53_(_pat_let12_0):
                    def iife54_(d_556_dt__update__tmp_h0_):
                        def iife55_(_pat_let13_0):
                            def iife56_(d_557_dt__update_hcf8_h0_):
                                def iife57_(_pat_let14_0):
                                    def iife58_(d_558_dt__update_hcf10_h0_):
                                        return D1_DC5(d_557_dt__update_hcf8_h0_, (d_556_dt__update__tmp_h0_).cf9, d_558_dt__update_hcf10_h0_)
                                    return iife58_(_pat_let14_0)
                                return iife57_(d_532_i0_)
                            return iife56_(_pat_let13_0)
                        return iife55_(pat_let_tv15_)
                    return iife54_(_pat_let12_0)
                nw67_[int(2)] = iife53_(d_552_v20_)
                nw67_[int(3)] = d_552_v20_
                nw67_[int(4)] = D1_DC5(d_545_v13_, d_546_v14_, d_532_i0_)
                nw67_[int(5)] = d_552_v20_
                nw67_[int(6)] = D1_DC5(_dafny.SeqWithoutIsStrInference([d_547_v15_ for d_559_i1_ in range(default__.abs(123))]), d_546_v14_, d_532_i0_)
                nw67_[int(7)] = d_552_v20_
                nw67_[int(8)] = D1_DC5(d_545_v13_, d_546_v14_, d_532_i0_)
                nw67_[int(9)] = d_552_v20_
                nw67_[int(10)] = d_552_v20_
                nw67_[int(11)] = d_552_v20_
                nw67_[int(12)] = d_552_v20_
                nw67_[int(13)] = d_552_v20_
                nw67_[int(14)] = d_552_v20_
                nw67_[int(15)] = d_552_v20_
                nw67_[int(16)] = D1_DC5(default__.fm32(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xehdabxv")), self.f13, ((d_553_v21_)[self.f13] if (self.f13) in (d_553_v21_) else p0), d_532_i0_, globalState), d_546_v14_, d_532_i0_)
                nw67_[int(17)] = D1_DC5(d_545_v13_, d_546_v14_, len(d_554_v22_))
                nw67_[int(18)] = d_552_v20_
                nw67_[int(19)] = d_552_v20_
                nw67_[int(20)] = d_552_v20_
                nw67_[int(21)] = d_552_v20_
                nw67_[int(22)] = d_552_v20_
                d_555_v23_ = nw67_
                rhs77_ = (d_555_v23_ if self.f13 else d_555_v23_)
                rhs78_ = self.f13
                lhs48_ = self
                d_555_v23_ = rhs77_
                lhs48_.f13 = rhs78_
                d_560_v24_: _dafny.Array
                def lambda16_(d_561_i0_):
                    def lambda17_(d_562_i2_):
                        return (d_562_i2_) + ((0) - (d_561_i0_))

                    return lambda17_

                init8_ = lambda16_(d_532_i0_)
                nw68_ = _dafny.Array(None, 27)
                for i0_8_ in range(nw68_.length(0)):
                    nw68_[i0_8_] = init8_(i0_8_)
                d_560_v24_ = nw68_
                index105_ = default__.safeIndex(322, (d_560_v24_).length(0))
                (d_560_v24_)[index105_] = default__.fm31(globalState)
                index106_ = default__.safeIndex(322, (d_560_v24_).length(0))
                (d_560_v24_)[index106_] = len(((_dafny.SeqWithoutIsStrInference([d_547_v15_ for d_563_i3_ in range(default__.abs(593))])) + (d_545_v13_)) + ((d_545_v13_) + (d_545_v13_)))
                index107_ = default__.safeIndex(322, (d_560_v24_).length(0))
                (d_560_v24_)[index107_] = (d_560_v24_)[default__.safeIndex(322, (d_560_v24_).length(0))]
                (self).f13 = self.f13
            d_564_v25_: _dafny.Seq
            d_564_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hujnpvv"))
            d_565_v26_: str
            d_565_v26_ = _dafny.CodePoint('w')
            (self).f13 = not((d_564_v25_) <= ((((d_564_v25_) + (d_564_v25_)).set(default__.safeIndex(d_532_i0_, len((d_564_v25_) + (d_564_v25_))), d_565_v26_)).set(default__.safeIndex(p0, len(((d_564_v25_) + (d_564_v25_)).set(default__.safeIndex(d_532_i0_, len((d_564_v25_) + (d_564_v25_))), d_565_v26_))), _dafny.CodePoint('f'))))
            d_566_v27_: int
            d_566_v27_ = 918
            d_566_v27_ = d_532_i0_
            d_566_v27_ = p0
        d_567_v28_: _dafny.Array
        nw69_ = _dafny.Array(int(0), 27)
        d_567_v28_ = nw69_
        d_568_v29_: _dafny.Map
        d_568_v29_ = _dafny.Map({d_567_v28_: d_567_v28_})
        d_568_v29_ = (_dafny.Map({d_567_v28_: d_567_v28_})) | ((d_568_v29_) | (d_568_v29_))
        d_569_v30_: C0
        nw70_ = C0()
        nw70_.ctor__(p0)
        d_569_v30_ = nw70_
        d_570_v31_: _dafny.Array
        def lambda18_(d_571_i4_):
            return self.f13

        init9_ = lambda18_
        nw71_ = _dafny.Array(None, 29)
        for i0_9_ in range(nw71_.length(0)):
            nw71_[i0_9_] = init9_(i0_9_)
        d_570_v31_ = nw71_
        d_572_v32_: D1
        d_572_v32_ = D1_DC4(d_570_v31_, (d_569_v30_).f15, _dafny.CodePoint('m'))
        d_573_v33_: _dafny.Seq
        d_573_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqogre"))
        d_574_v34_: _dafny.Map
        d_574_v34_ = _dafny.Map({(_dafny.CodePoint('t') if False else (d_572_v32_).cf7): d_573_v33_})
        d_575_v35_: str
        d_575_v35_ = _dafny.CodePoint('f')
        d_574_v34_ = (d_574_v34_).set(d_575_v35_, d_573_v33_)
        d_576_v36_: _dafny.Seq
        d_576_v36_ = _dafny.SeqWithoutIsStrInference([self.f13])
        index108_ = default__.safeIndex(501, (d_567_v28_).length(0))
        (d_567_v28_)[index108_] = len(d_576_v36_)
        index109_ = default__.safeIndex(501, (d_567_v28_).length(0))
        (d_567_v28_)[index109_] = (d_569_v30_).f15
        d_577_v37_: _dafny.MultiSet
        d_577_v37_ = _dafny.MultiSet([(d_567_v28_)[default__.safeIndex(501, (d_567_v28_).length(0))]])
        d_576_v36_ = _dafny.SeqWithoutIsStrInference([(d_577_v37_).ispropersubset(d_577_v37_), not((523) < (p0))])
        r0 = d_567_v28_
        return r0

    def m2(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        r2: bool = False
        r3: _dafny.Seq = _dafny.Seq({})
        d_578_v0_: _dafny.Set
        d_578_v0_ = _dafny.Set({p0})
        hi4_ = p2
        for d_579_i0_ in range(len(default__.fm33(self.f13, False, d_578_v0_, self.f13, globalState)), hi4_):
            d_580_v1_: _dafny.Array
            nw72_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_580_v1_ = nw72_
            d_581_v2_: _dafny.Array
            nw73_ = _dafny.Array(int(0), 25)
            d_581_v2_ = nw73_
            d_582_v3_: str
            d_582_v3_ = _dafny.CodePoint('p')
            d_583_v4_: _dafny.MultiSet
            d_583_v4_ = _dafny.MultiSet([d_582_v3_])
            index110_ = default__.safeIndex(110, (d_581_v2_).length(0))
            (d_581_v2_)[index110_] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wawnlrb"))), ((d_583_v4_)[d_582_v3_] if (d_582_v3_) in (d_583_v4_) else d_579_i0_))
            d_584_v5_: _dafny.Set
            d_584_v5_ = _dafny.Set({self.f13, self.f13, self.f13})
            d_585_v6_: _dafny.Seq
            d_585_v6_ = _dafny.SeqWithoutIsStrInference([p0, default__.fm31(globalState), len(d_584_v5_), p0])
            d_586_v7_: _dafny.MultiSet
            d_586_v7_ = _dafny.MultiSet([d_585_v6_, d_585_v6_])
            d_587_v8_: _dafny.Map
            d_587_v8_ = _dafny.Map({self.f13: d_579_i0_})
            d_588_v9_: D5
            d_588_v9_ = D5_DC12(d_587_v8_)
            d_589_v10_: _dafny.Map
            d_589_v10_ = _dafny.Map({p2: d_588_v9_})
            d_590_v11_: _dafny.Map
            d_590_v11_ = _dafny.Map({(self).fm30(d_586_v7_, p0, d_589_v10_, globalState): _dafny.CodePoint('h')})
            index111_ = default__.safeIndex(110, (d_581_v2_).length(0))
            rhs79_ = self.f13
            rhs80_ = d_580_v1_
            rhs81_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tasrqgp")))) - (len((d_590_v11_) | (d_590_v11_)))
            lhs49_ = d_581_v2_
            lhs50_ = default__.safeIndex(110, (d_581_v2_).length(0))
            r2 = rhs79_
            d_580_v1_ = rhs80_
            lhs49_[lhs50_] = rhs81_
            r1 = (((d_581_v2_)[default__.safeIndex(110, (d_581_v2_).length(0))]) <= (d_579_i0_)) or (self.f13)
            index112_ = default__.safeIndex(110, (d_581_v2_).length(0))
            (d_581_v2_)[index112_] = 62
            d_591_v12_: _dafny.Array
            nw74_ = _dafny.Array(_dafny.MultiSet({}), 18)
            d_591_v12_ = nw74_
            d_592_v13_: _dafny.Array
            nw75_ = _dafny.Array(False, 4)
            d_592_v13_ = nw75_
            d_593_v14_: _dafny.MultiSet
            d_593_v14_ = _dafny.MultiSet([d_592_v13_])
            index113_ = default__.safeIndex(766, (d_591_v12_).length(0))
            (d_591_v12_)[index113_] = d_593_v14_
            index114_ = default__.safeIndex(766, (d_591_v12_).length(0))
            (d_591_v12_)[index114_] = (d_593_v14_).set(d_592_v13_, default__.abs(d_579_i0_))
        d_594_v16_: _dafny.Map
        d_594_v16_ = _dafny.Map({self.f13: p0})
        d_595_v17_: str
        d_595_v17_ = _dafny.CodePoint('w')
        d_596_v18_: _dafny.Map
        d_596_v18_ = _dafny.Map({(p1).set(default__.safeIndex(len(d_594_v16_), len(p1)), d_595_v17_): p2})
        d_597_v19_: int
        out29_: int
        def iife59_():
            coll29_ = _dafny.Map()
            compr_29_: _dafny.Seq
            for compr_29_ in (d_596_v18_).keys.Elements:
                d_598_v15_: _dafny.Seq = compr_29_
                if (d_598_v15_) in (d_596_v18_):
                    coll29_[d_598_v15_] = self.f13
            return _dafny.Map(coll29_)
        out29_ = default__.m0(iife59_()
        , globalState)
        d_597_v19_ = out29_
        d_599_i1_: int
        d_599_i1_ = 0
        with _dafny.label("7"):
            while self.f13:
                with _dafny.c_label("7"):
                    if (d_599_i1_) >= (100):
                        raise _dafny.Break("7")
                    d_599_i1_ = (d_599_i1_) + (1)
                    d_600_v20_: _dafny.MultiSet
                    d_600_v20_ = _dafny.MultiSet([251])
                    d_601_v21_: D2
                    d_601_v21_ = D2_DC7(p1, p2, d_597_v19_, not(self.f13))
                    d_602_v22_: _dafny.Set
                    d_602_v22_ = _dafny.Set({d_595_v17_})
                    d_603_v23_: _dafny.Map
                    d_603_v23_ = _dafny.Map({len(d_602_v22_): p2})
                    d_604_v24_: _dafny.Map
                    d_604_v24_ = _dafny.Map({d_600_v20_: (0) - (d_597_v19_)})
                    d_605_v25_: D0
                    d_605_v25_ = D0_DC1(self.f13, p2)
                    d_606_v26_: _dafny.Array
                    nw76_ = _dafny.Array(None, 5)
                    nw76_[int(0)] = self.f13
                    nw76_[int(1)] = (default__.fm34(globalState)) not in (p1)
                    nw76_[int(2)] = ((d_600_v20_).set(len((d_601_v21_).cf12), default__.abs((0) - (len(d_603_v23_))))) not in (d_604_v24_)
                    nw76_[int(3)] = (self).fm4(self.f13, d_605_v25_, d_597_v19_, d_603_v23_, globalState)
                    nw76_[int(4)] = self.f13
                    d_606_v26_ = nw76_
                    index115_ = default__.safeIndex(493, (d_606_v26_).length(0))
                    (d_606_v26_)[index115_] = True
                    index116_ = default__.safeIndex(493, (d_606_v26_).length(0))
                    (d_606_v26_)[index116_] = (_dafny.SeqWithoutIsStrInference([d_595_v17_ for d_607_i2_ in range(default__.abs(106))])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))
                    d_608_v27_: _dafny.Seq
                    d_608_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
                    d_608_v27_ = d_608_v27_
                    if (d_606_v26_)[default__.safeIndex(493, (d_606_v26_).length(0))]:
                        d_609_v28_: _dafny.Array
                        nw77_ = _dafny.Array(_dafny.Seq({}), 5)
                        d_609_v28_ = nw77_
                        nw78_ = _dafny.Array(_dafny.Seq({}), 8)
                        d_609_v28_ = nw78_
                        d_597_v19_ = (default__.fm31(globalState)) * ((len(d_608_v27_)) * (default__.fm31(globalState)))
                        d_610_v29_: _dafny.Seq
                        d_610_v29_ = _dafny.SeqWithoutIsStrInference([not((d_606_v26_)[default__.safeIndex(493, (d_606_v26_).length(0))]), (d_606_v26_)[default__.safeIndex(493, (d_606_v26_).length(0))]])
                        d_611_v30_: _dafny.Set
                        d_611_v30_ = _dafny.Set({(d_606_v26_)[default__.safeIndex(493, (d_606_v26_).length(0))]})
                        d_612_v31_: _dafny.Seq
                        d_612_v31_ = _dafny.SeqWithoutIsStrInference([len(d_611_v30_), d_597_v19_])
                        d_613_v32_: D1
                        d_613_v32_ = D1_DC4(d_606_v26_, p2, _dafny.CodePoint('y'))
                        d_614_v33_: _dafny.MultiSet
                        d_614_v33_ = _dafny.MultiSet([(d_613_v32_).cf7])
                        d_615_v34_: _dafny.Array
                        nw79_ = _dafny.Array(None, 15)
                        nw79_[int(0)] = (_dafny.MultiSet(d_610_v29_)).cardinality
                        nw79_[int(1)] = d_597_v19_
                        nw79_[int(2)] = 300
                        nw79_[int(3)] = p0
                        nw79_[int(4)] = p0
                        nw79_[int(5)] = len(d_610_v29_)
                        nw79_[int(6)] = len(d_594_v16_)
                        nw79_[int(7)] = p2
                        nw79_[int(8)] = p2
                        nw79_[int(9)] = (d_612_v31_)[default__.safeIndex(((d_614_v33_)[d_595_v17_] if (d_595_v17_) in (d_614_v33_) else d_597_v19_), len(d_612_v31_))]
                        nw79_[int(10)] = d_597_v19_
                        nw79_[int(11)] = p0
                        nw79_[int(12)] = ((d_600_v20_)[d_597_v19_] if (d_597_v19_) in (d_600_v20_) else p0)
                        nw79_[int(13)] = 304
                        nw79_[int(14)] = p2
                        d_615_v34_ = nw79_
                        d_616_v35_: C1
                        nw80_ = C1()
                        nw80_.ctor__(self.f13, d_608_v27_, d_615_v34_)
                        d_616_v35_ = nw80_
                        d_597_v19_ = p0
                        d_617_v36_: _dafny.Map
                        d_617_v36_ = _dafny.Map({self.f13: d_595_v17_})
                        d_618_v37_: _dafny.Array
                        nw81_ = _dafny.Array(_dafny.Seq({}), 13)
                        d_618_v37_ = nw81_
                        index117_ = default__.safeIndex(845, (d_618_v37_).length(0))
                        (d_618_v37_)[index117_] = d_610_v29_
                        index118_ = default__.safeIndex(845, (d_618_v37_).length(0))
                        rhs82_ = -25
                        rhs83_ = _dafny.Map({(p1) == ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krd"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krd")))), _dafny.CodePoint('o'))): d_595_v17_})
                        rhs84_ = d_612_v31_
                        rhs85_ = _dafny.SeqWithoutIsStrInference([(self.f13) == (self.f13), (d_616_v35_).f17])
                        rhs86_ = d_608_v27_
                        lhs51_ = d_618_v37_
                        lhs52_ = default__.safeIndex(845, (d_618_v37_).length(0))
                        d_597_v19_ = rhs82_
                        d_617_v36_ = rhs83_
                        d_612_v31_ = rhs84_
                        lhs51_[lhs52_] = rhs85_
                        d_608_v27_ = rhs86_
                    elif True:
                        d_619_v38_: _dafny.Array
                        nw82_ = _dafny.Array(None, 3)
                        nw82_[int(0)] = p2
                        nw82_[int(1)] = d_597_v19_
                        nw82_[int(2)] = 149
                        d_619_v38_ = nw82_
                        index119_ = default__.safeIndex(825, (d_619_v38_).length(0))
                        (d_619_v38_)[index119_] = d_597_v19_
                        d_620_v39_: _dafny.Seq
                        d_620_v39_ = _dafny.SeqWithoutIsStrInference([d_597_v19_, p0, p0])
                        index120_ = default__.safeIndex(825, (d_619_v38_).length(0))
                        (d_619_v38_)[index120_] = ((p2) * (p2) if (d_606_v26_)[default__.safeIndex(493, (d_606_v26_).length(0))] else len((d_620_v39_) + (d_620_v39_)))
                        d_621_v40_: C0
                        nw83_ = C0()
                        nw83_.ctor__(d_597_v19_)
                        d_621_v40_ = nw83_
                        d_622_v41_: _dafny.Array
                        nw84_ = _dafny.Array(_dafny.Set({}), 2)
                        d_622_v41_ = nw84_
                        d_622_v41_ = d_622_v41_
                        rhs87_ = d_595_v17_
                        rhs88_ = self.f13
                        lhs53_ = self
                        d_595_v17_ = rhs87_
                        lhs53_.f13 = rhs88_
                        d_608_v27_ = (p1 if ((d_606_v26_)[default__.safeIndex(493, (d_606_v26_).length(0))] if (d_606_v26_)[default__.safeIndex(493, (d_606_v26_).length(0))] else (d_606_v26_)[default__.safeIndex(493, (d_606_v26_).length(0))]) else p1)
                    d_623_v42_: _dafny.Map
                    d_623_v42_ = _dafny.Map({d_608_v27_: d_606_v26_})
                    d_623_v42_ = (d_623_v42_).set(d_608_v27_, d_606_v26_)
                    pass
            pass
        hi5_ = 151
        for d_624_i3_ in range((0) - (len((d_578_v0_) - (d_578_v0_))), hi5_):
            d_625_v43_: _dafny.Seq
            d_625_v43_ = _dafny.SeqWithoutIsStrInference([self.f13])
            d_626_v44_: _dafny.Seq
            d_626_v44_ = _dafny.SeqWithoutIsStrInference([(d_625_v43_)[default__.safeIndex(p2, len(d_625_v43_))]])
            d_627_v45_: _dafny.Map
            d_627_v45_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([True]): self.f13})
            (self).f13 = ((d_625_v43_) in (d_627_v45_)) and (self.f13)
            d_628_v46_: _dafny.Array
            nw85_ = _dafny.Array(int(0), 27)
            d_628_v46_ = nw85_
            d_629_v47_: C1
            nw86_ = C1()
            nw86_.ctor__(self.f13, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwxskgni")), d_628_v46_)
            d_629_v47_ = nw86_
            d_630_v48_: D0
            d_630_v48_ = D0_DC0(self.f13)
            pat_let_tv16_ = d_629_v47_
            pat_let_tv17_ = d_629_v47_
            d_631_v49_: _dafny.Array
            nw87_ = _dafny.Array(None, 6)
            def iife60_(_pat_let15_0):
                def iife61_(d_632_dt__update__tmp_h0_):
                    def iife62_(_pat_let16_0):
                        def iife63_(d_633_dt__update_hcf0_h0_):
                            return D0_DC0(d_633_dt__update_hcf0_h0_)
                        return iife63_(_pat_let16_0)
                    return iife62_(False)
                return iife61_(_pat_let15_0)
            nw87_[int(0)] = iife60_(d_630_v48_)
            def iife64_(_pat_let17_0):
                def iife65_(d_634_dt__update__tmp_h1_):
                    def iife66_(_pat_let18_0):
                        def iife67_(d_635_dt__update_hcf0_h1_):
                            return D0_DC0(d_635_dt__update_hcf0_h1_)
                        return iife67_(_pat_let18_0)
                    return iife66_((pat_let_tv16_).f17)
                return iife65_(_pat_let17_0)
            nw87_[int(1)] = iife64_(d_630_v48_)
            nw87_[int(2)] = d_630_v48_
            def iife68_(_pat_let19_0):
                def iife69_(d_636_dt__update__tmp_h2_):
                    def iife70_(_pat_let20_0):
                        def iife71_(d_637_dt__update_hcf0_h2_):
                            return D0_DC0(d_637_dt__update_hcf0_h2_)
                        return iife71_(_pat_let20_0)
                    return iife70_((pat_let_tv17_).f17)
                return iife69_(_pat_let19_0)
            nw87_[int(3)] = iife68_(d_630_v48_)
            nw87_[int(4)] = d_630_v48_
            nw87_[int(5)] = D0_DC0(self.f13)
            d_631_v49_ = nw87_
            (d_629_v47_).m5(d_631_v49_, globalState)
            d_638_v50_: _dafny.Map
            d_638_v50_ = _dafny.Map({(d_629_v47_).f17: _dafny.Map({d_624_i3_: d_624_i3_})})
            d_639_v51_: _dafny.Map
            d_639_v51_ = _dafny.Map({_dafny.MultiSet(d_625_v43_): ((d_629_v47_).f17) not in (d_638_v50_)})
            d_640_v52_: _dafny.MultiSet
            d_640_v52_ = _dafny.MultiSet([True])
            d_641_v53_: _dafny.Array
            nw88_ = _dafny.Array(None, 20)
            nw88_[int(0)] = (d_629_v47_).f17
            nw88_[int(1)] = self.f13
            nw88_[int(2)] = self.f13
            nw88_[int(3)] = (d_629_v47_).f17
            nw88_[int(4)] = self.f13
            nw88_[int(5)] = not(self.f13)
            nw88_[int(6)] = not(self.f13)
            nw88_[int(7)] = not((d_640_v52_).isdisjoint(d_640_v52_))
            nw88_[int(8)] = not((d_629_v47_).f17)
            nw88_[int(9)] = (d_629_v47_).f17
            nw88_[int(10)] = self.f13
            nw88_[int(11)] = (d_629_v47_).f17
            nw88_[int(12)] = (p0) == (p0)
            nw88_[int(13)] = True
            nw88_[int(14)] = self.f13
            nw88_[int(15)] = not (self.f13) or ((d_629_v47_).f17)
            nw88_[int(16)] = (p2) > (p2)
            nw88_[int(17)] = self.f13
            nw88_[int(18)] = self.f13
            nw88_[int(19)] = not(self.f13)
            d_641_v53_ = nw88_
            index121_ = default__.safeIndex(413, (d_641_v53_).length(0))
            (d_641_v53_)[index121_] = (d_629_v47_).fm6(globalState)
            index122_ = default__.safeIndex(910, (d_641_v53_).length(0))
            (d_641_v53_)[index122_] = self.f13
            d_642_v54_: _dafny.Seq
            d_642_v54_ = _dafny.SeqWithoutIsStrInference([d_624_i3_, p2])
            d_643_v55_: _dafny.Map
            d_643_v55_ = _dafny.Map({(d_629_v47_).fm7(d_624_i3_, (0) - ((d_642_v54_)[default__.safeIndex(len(d_629_v47_.f18), len(d_642_v54_))]), self.f13, globalState): len(p1)})
            d_644_v56_: _dafny.Set
            d_644_v56_ = _dafny.Set({(d_629_v47_).fm14(d_624_i3_, (d_629_v47_).f17, len(d_643_v55_), p2, globalState)})
            index123_ = default__.safeIndex(413, (d_641_v53_).length(0))
            index124_ = default__.safeIndex(910, (d_641_v53_).length(0))
            rhs89_ = (_dafny.Map({_dafny.MultiSet([self.f13, (d_630_v48_).cf0, (d_629_v47_).f17]): (d_629_v47_).f17})) | (_dafny.Map({_dafny.MultiSet(d_626_v44_): (d_629_v47_).f17}))
            rhs90_ = True
            rhs91_ = (d_644_v56_).issubset(d_644_v56_)
            rhs92_ = not((d_629_v47_).f17)
            rhs93_ = (p1) != (d_629_v47_.f18)
            lhs54_ = d_641_v53_
            lhs55_ = default__.safeIndex(413, (d_641_v53_).length(0))
            lhs56_ = d_641_v53_
            lhs57_ = default__.safeIndex(910, (d_641_v53_).length(0))
            d_639_v51_ = rhs89_
            r1 = rhs90_
            lhs54_[lhs55_] = rhs91_
            lhs56_[lhs57_] = rhs92_
            r1 = rhs93_
        d_645_v57_: _dafny.Set
        d_645_v57_ = _dafny.Set({self.f13, not(self.f13)})
        d_646_v58_: _dafny.Map
        d_646_v58_ = _dafny.Map({self.f13: self.f13})
        d_647_v59_: _dafny.Array
        nw89_ = _dafny.Array(None, 16)
        nw89_[int(0)] = p0
        nw89_[int(1)] = p2
        nw89_[int(2)] = (0) - (p2)
        nw89_[int(3)] = p0
        nw89_[int(4)] = p0
        nw89_[int(5)] = len(d_645_v57_)
        nw89_[int(6)] = p2
        nw89_[int(7)] = p0
        nw89_[int(8)] = p2
        nw89_[int(9)] = len(p1)
        nw89_[int(10)] = d_597_v19_
        nw89_[int(11)] = p0
        nw89_[int(12)] = d_597_v19_
        nw89_[int(13)] = len(d_646_v58_)
        nw89_[int(14)] = d_597_v19_
        nw89_[int(15)] = p2
        d_647_v59_ = nw89_
        d_648_v60_: D4
        d_648_v60_ = D4_DC10((d_647_v59_ if self.f13 else d_647_v59_))
        source12_ = d_648_v60_
        if source12_.is_DC11:
            d_649___mcc_h0_ = source12_.cf21
            d_650_cf21_ = d_649___mcc_h0_
            d_651_v61_: C0
            nw90_ = C0()
            nw90_.ctor__(default__.safeDivisionInt(822, default__.fm31(globalState)))
            d_651_v61_ = nw90_
            d_652_v62_: _dafny.Map
            d_652_v62_ = _dafny.Map({default__.fm35(p0, globalState): False})
            d_653_v63_: _dafny.Array
            nw91_ = _dafny.Array(None, 14)
            nw91_[int(0)] = d_647_v59_
            nw91_[int(1)] = d_647_v59_
            nw91_[int(2)] = d_647_v59_
            nw91_[int(3)] = d_647_v59_
            nw91_[int(4)] = d_647_v59_
            nw91_[int(5)] = d_647_v59_
            nw91_[int(6)] = d_647_v59_
            nw91_[int(7)] = d_647_v59_
            nw91_[int(8)] = d_647_v59_
            nw91_[int(9)] = d_647_v59_
            nw91_[int(10)] = d_647_v59_
            nw91_[int(11)] = d_647_v59_
            nw91_[int(12)] = d_647_v59_
            nw91_[int(13)] = d_647_v59_
            d_653_v63_ = nw91_
            d_654_v64_: C1
            nw92_ = C1()
            nw92_.ctor__(self.f13, p1, d_647_v59_)
            d_654_v64_ = nw92_
            d_655_v65_: _dafny.Map
            d_655_v65_ = _dafny.Map({d_653_v63_: _dafny.Set({d_654_v64_})})
            d_656_v66_: _dafny.Set
            d_656_v66_ = _dafny.Set({d_654_v64_})
            d_657_v67_: _dafny.Map
            d_657_v67_ = _dafny.Map({(p2) - (len(d_652_v62_)): ((0) - (len(((d_655_v65_)[d_653_v63_] if (d_653_v63_) in (d_655_v65_) else d_656_v66_))) if not(d_650_cf21_) else d_597_v19_)})
            d_657_v67_ = d_657_v67_
            d_658_v68_: _dafny.Array
            nw93_ = _dafny.Array(False, 11)
            d_658_v68_ = nw93_
            index125_ = default__.safeIndex(657, (d_658_v68_).length(0))
            (d_658_v68_)[index125_] = (False if self.f13 else d_650_cf21_)
            d_659_v69_: _dafny.MultiSet
            d_659_v69_ = _dafny.MultiSet([d_658_v68_, d_658_v68_, d_658_v68_])
            index126_ = default__.safeIndex(657, (d_658_v68_).length(0))
            (d_658_v68_)[index126_] = not(((d_659_v69_).isdisjoint(d_659_v69_)) or (d_650_cf21_))
            index127_ = default__.safeIndex(657, (d_658_v68_).length(0))
            (d_658_v68_)[index127_] = (d_654_v64_).f17
        elif True:
            d_660___mcc_h1_ = source12_.cf20
            d_661_cf20_ = d_660___mcc_h1_
            d_662_v70_: _dafny.Seq
            d_662_v70_ = _dafny.SeqWithoutIsStrInference([self.f13])
            d_663_v71_: C1
            nw94_ = C1()
            nw94_.ctor__((not((d_662_v70_)[default__.safeIndex(d_597_v19_, len(d_662_v70_))])) == (self.f13), p1, d_661_cf20_)
            d_663_v71_ = nw94_
            if self.f13:
                d_664_v72_: _dafny.Seq
                d_664_v72_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqgr"))).set(default__.safeIndex(d_597_v19_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqgr")))), d_595_v17_), _dafny.SeqWithoutIsStrInference([d_595_v17_ for d_665_i4_ in range(default__.abs(-754))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uthxbo"))])
                d_666_v73_: _dafny.Map
                d_666_v73_ = _dafny.Map({(d_663_v71_).f17: d_664_v72_})
                d_667_v74_: _dafny.Seq
                d_667_v74_ = _dafny.SeqWithoutIsStrInference([len(d_645_v57_)])
                r1 = ((((d_666_v73_)[(d_663_v71_).f17] if ((d_663_v71_).f17) in (d_666_v73_) else d_664_v72_) if self.f13 else d_664_v72_)) <= (_dafny.SeqWithoutIsStrInference([default__.fm32(d_663_v71_.f18, (d_663_v71_).f17, p2, (d_667_v74_)[default__.safeIndex(len(d_594_v16_), len(d_667_v74_))], globalState)]))
                d_597_v19_ = (0) - ((d_597_v19_) * (d_597_v19_))
                d_578_v0_ = (d_578_v0_ if (d_663_v71_).f17 else d_578_v0_)
                d_597_v19_ = (len(_dafny.SeqWithoutIsStrInference([p2]))) * ((0) - (default__.safeModuloInt(len(d_596_v18_), 416)))
                d_668_v75_: C1
                nw95_ = C1()
                nw95_.ctor__((d_663_v71_).f17, p1, d_661_cf20_)
                d_668_v75_ = nw95_
            elif True:
                d_669_v76_: _dafny.Array
                def lambda19_(d_670_v17_):
                    def lambda20_(d_671_i5_):
                        return _dafny.SeqWithoutIsStrInference([d_670_v17_ for d_672_i6_ in range(default__.abs(4))])

                    return lambda20_

                init10_ = lambda19_(d_595_v17_)
                nw96_ = _dafny.Array(None, 13)
                for i0_10_ in range(nw96_.length(0)):
                    nw96_[i0_10_] = init10_(i0_10_)
                d_669_v76_ = nw96_
                d_673_v77_: _dafny.Seq
                d_673_v77_ = _dafny.SeqWithoutIsStrInference([d_669_v76_, d_669_v76_, d_669_v76_, d_669_v76_, d_669_v76_])
                d_669_v76_ = (d_673_v77_)[default__.safeIndex(len(_dafny.Map({d_597_v19_: True})), len(d_673_v77_))]
                index128_ = default__.safeIndex(114, (d_661_cf20_).length(0))
                (d_661_cf20_)[index128_] = p0
                index129_ = default__.safeIndex(114, (d_661_cf20_).length(0))
                (d_661_cf20_)[index129_] = default__.safeModuloInt(p2, p0)
                r2 = self.f13
                d_597_v19_ = (0) - ((0) - ((-336) - ((d_661_cf20_)[default__.safeIndex(114, (d_661_cf20_).length(0))])))
                d_674_v78_: _dafny.Map
                d_674_v78_ = _dafny.Map({p2: self.f13})
                d_674_v78_ = (d_674_v78_) | (d_674_v78_)
            r2 = (d_663_v71_).f17
            source13_ = (d_648_v60_ if (p2) > (d_597_v19_) else D4_DC10(d_661_cf20_))
            if source13_.is_DC11:
                d_675___mcc_h2_ = source13_.cf21
                d_676_cf21_ = d_675___mcc_h2_
                d_677_v79_: _dafny.Array
                nw97_ = _dafny.Array(_dafny.Seq({}), 2)
                d_677_v79_ = nw97_
                d_678_v80_: _dafny.Seq
                d_678_v80_ = _dafny.SeqWithoutIsStrInference([p0])
                d_679_v81_: _dafny.Seq
                d_679_v81_ = _dafny.SeqWithoutIsStrInference([(d_678_v80_)[default__.safeIndex(d_597_v19_, len(d_678_v80_))]])
                index130_ = default__.safeIndex(621, (d_677_v79_).length(0))
                (d_677_v79_)[index130_] = (d_679_v81_) + (d_678_v80_)
                d_680_v82_: _dafny.Array
                nw98_ = _dafny.Array(_dafny.Map({}), 10)
                d_680_v82_ = nw98_
                d_681_v83_: D0
                d_681_v83_ = D0_DC1((d_663_v71_).f17, p0)
                d_682_v84_: D5
                d_682_v84_ = D5_DC14(d_676_cf21_)
                d_683_v85_: D4
                d_683_v85_ = D4_DC11(self.f13)
                d_684_v86_: _dafny.Array
                nw99_ = _dafny.Array(None, 26)
                nw99_[int(0)] = (d_663_v71_).f17
                nw99_[int(1)] = self.f13
                nw99_[int(2)] = d_676_cf21_
                nw99_[int(3)] = (d_663_v71_).f17
                nw99_[int(4)] = False
                nw99_[int(5)] = d_676_cf21_
                nw99_[int(6)] = False
                nw99_[int(7)] = not(not((d_663_v71_).f17))
                nw99_[int(8)] = d_676_cf21_
                nw99_[int(9)] = False
                nw99_[int(10)] = (d_663_v71_).f17
                nw99_[int(11)] = (self).fm4((d_663_v71_).f17, d_681_v83_, (default__.fm36(globalState)).cardinality, _dafny.Map({106: d_597_v19_}), globalState)
                nw99_[int(12)] = d_676_cf21_
                nw99_[int(13)] = not((d_663_v71_).fm14(p0, (d_663_v71_).f17, d_597_v19_, (d_663_v71_).fm7(p2, len(d_663_v71_.f18), self.f13, globalState), globalState))
                nw99_[int(14)] = (d_682_v84_).cf23
                nw99_[int(15)] = self.f13
                nw99_[int(16)] = self.f13
                nw99_[int(17)] = (d_662_v70_)[default__.safeIndex(d_597_v19_, len(d_662_v70_))]
                nw99_[int(18)] = (d_663_v71_).f17
                nw99_[int(19)] = (d_663_v71_).f17
                nw99_[int(20)] = (d_663_v71_).f17
                nw99_[int(21)] = not((d_663_v71_).f17)
                nw99_[int(22)] = (d_663_v71_).fm6(globalState)
                nw99_[int(23)] = (d_663_v71_).f17
                nw99_[int(24)] = True
                nw99_[int(25)] = (d_683_v85_).cf21
                d_684_v86_ = nw99_
                index131_ = default__.safeIndex(282, (d_680_v82_).length(0))
                (d_680_v82_)[index131_] = _dafny.Map({p0: d_684_v86_})
                index132_ = default__.safeIndex(563, (d_684_v86_).length(0))
                (d_684_v86_)[index132_] = (d_663_v71_).f17
                d_685_v87_: _dafny.MultiSet
                d_685_v87_ = _dafny.MultiSet([(0) - (p2)])
                d_686_v88_: _dafny.Map
                d_686_v88_ = _dafny.Map({((d_685_v87_).cardinality) * (p2): d_684_v86_})
                index133_ = default__.safeIndex(621, (d_677_v79_).length(0))
                index134_ = default__.safeIndex(282, (d_680_v82_).length(0))
                index135_ = default__.safeIndex(563, (d_684_v86_).length(0))
                rhs94_ = default__.fm35((d_663_v71_).fm15(globalState), globalState)
                rhs95_ = d_647_v59_
                rhs96_ = d_686_v88_
                rhs97_ = (self.f13) == (self.f13)
                rhs98_ = d_661_cf20_
                lhs58_ = d_677_v79_
                lhs59_ = default__.safeIndex(621, (d_677_v79_).length(0))
                lhs60_ = d_680_v82_
                lhs61_ = default__.safeIndex(282, (d_680_v82_).length(0))
                lhs62_ = d_684_v86_
                lhs63_ = default__.safeIndex(563, (d_684_v86_).length(0))
                lhs58_[lhs59_] = rhs94_
                d_647_v59_ = rhs95_
                lhs60_[lhs61_] = rhs96_
                lhs62_[lhs63_] = rhs97_
                d_661_cf20_ = rhs98_
                d_597_v19_ = ((d_594_v16_)[self.f13] if (self.f13) in (d_594_v16_) else p0)
                d_687_v89_: _dafny.Map
                d_687_v89_ = _dafny.Map({(0) - (p2): d_661_cf20_})
                d_688_v90_: _dafny.Map
                d_688_v90_ = _dafny.Map({p2: (p0) < (len(d_687_v89_))})
                rhs99_ = (d_663_v71_).fm6(globalState)
                rhs100_ = d_645_v57_
                rhs101_ = (p2) + ((p0 if False else (d_663_v71_).fm7(d_597_v19_, p2, (d_663_v71_).f17, globalState)))
                rhs102_ = ((d_688_v90_)[default__.safeDivisionInt(p0, d_597_v19_)] if (default__.safeDivisionInt(p0, d_597_v19_)) in (d_688_v90_) else (d_685_v87_).issubset(d_685_v87_))
                rhs103_ = self.f13
                r2 = rhs99_
                d_645_v57_ = rhs100_
                d_597_v19_ = rhs101_
                d_676_cf21_ = rhs102_
                r1 = rhs103_
                d_597_v19_ = 976
            elif True:
                d_689___mcc_h3_ = source13_.cf20
                d_690_cf20_ = d_689___mcc_h3_
                d_597_v19_ = len(((d_646_v58_) | (_dafny.Map({(d_663_v71_).f17: (d_663_v71_).f17}))) | ((d_646_v58_) | (_dafny.Map({(d_663_v71_).f17: (d_663_v71_).f17}))))
                d_597_v19_ = d_597_v19_
                d_691_v91_: _dafny.MultiSet
                d_691_v91_ = _dafny.MultiSet([self.f13])
                d_597_v19_ = default__.safeDivisionInt((p2) * (((d_691_v91_).set((d_663_v71_).f17, default__.abs(p0))).cardinality), d_597_v19_)
                d_692_v92_: _dafny.Seq
                d_692_v92_ = _dafny.SeqWithoutIsStrInference([d_663_v71_])
                d_693_v93_: _dafny.Map
                d_693_v93_ = _dafny.Map({p0: 581})
                d_694_v94_: _dafny.Map
                d_694_v94_ = _dafny.Map({((d_693_v93_)[p0] if (p0) in (d_693_v93_) else ((d_693_v93_)[26] if (26) in (d_693_v93_) else p2)): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))})
                d_695_v95_: T1
                nw100_ = C1()
                nw100_.ctor__(((d_663_v71_).fm15(globalState)) <= (len(d_692_v92_)), ((d_694_v94_)[d_597_v19_] if (d_597_v19_) in (d_694_v94_) else p1), d_690_cf20_)
                d_695_v95_ = nw100_
                d_696_v96_: _dafny.Map
                d_696_v96_ = _dafny.Map({d_597_v19_: d_594_v16_})
                index136_ = default__.safeIndex(576, ((d_695_v95_).f14).length(0))
                ((d_695_v95_).f14)[index136_] = len(((_dafny.SeqWithoutIsStrInference([d_663_v71_.f18, d_663_v71_.f18, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sp")), d_663_v71_.f18, d_663_v71_.f18])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_663_v71_.f18, d_663_v71_.f18, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sp")), d_663_v71_.f18, d_663_v71_.f18]))), p1)).set(default__.safeIndex(p2, len((_dafny.SeqWithoutIsStrInference([d_663_v71_.f18, d_663_v71_.f18, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sp")), d_663_v71_.f18, d_663_v71_.f18])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_663_v71_.f18, d_663_v71_.f18, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sp")), d_663_v71_.f18, d_663_v71_.f18]))), p1))), d_663_v71_.f18))
                d_697_v97_: D7
                d_697_v97_ = D7_DC19(d_695_v95_)
                index137_ = default__.safeIndex(576, ((d_695_v95_).f14).length(0))
                rhs104_ = (d_697_v97_).cf30
                rhs105_ = default__.fm37((d_663_v71_).f17, globalState)
                rhs106_ = default__.safeDivisionInt(p2, (len(p1)) + (len(_dafny.SeqWithoutIsStrInference([d_595_v17_ for d_698_i7_ in range(default__.abs(-849))]))))
                lhs64_ = (d_695_v95_).f14
                lhs65_ = default__.safeIndex(576, ((d_695_v95_).f14).length(0))
                d_695_v95_ = rhs104_
                d_696_v96_ = rhs105_
                lhs64_[lhs65_] = rhs106_
        d_699_v98_: _dafny.Array
        def lambda21_(d_700_i8_):
            return self.f13

        init11_ = lambda21_
        nw101_ = _dafny.Array(None, 24)
        for i0_11_ in range(nw101_.length(0)):
            nw101_[i0_11_] = init11_(i0_11_)
        d_699_v98_ = nw101_
        d_701_v99_: _dafny.MultiSet
        d_701_v99_ = _dafny.MultiSet([d_699_v98_, d_699_v98_])
        (self).f13 = ((d_701_v99_) | (d_701_v99_)) != ((_dafny.MultiSet([d_699_v98_])).intersection(d_701_v99_))
        d_702_v100_: _dafny.Map
        d_702_v100_ = _dafny.Map({p2: d_699_v98_})
        r0 = (d_702_v100_) | (d_702_v100_)
        r1 = self.f13
        d_703_v101_: D2
        d_703_v101_ = D2_DC7(_dafny.SeqWithoutIsStrInference([d_595_v17_]), len(d_645_v57_), 749, self.f13)
        r2 = ((d_703_v101_ if self.f13 else d_703_v101_)).cf15
        d_704_v102_: _dafny.Seq
        d_704_v102_ = _dafny.SeqWithoutIsStrInference([self.f13])
        r3 = (d_704_v102_ if (self.f13) in (d_704_v102_) else (d_704_v102_) + (d_704_v102_))
        return r0, r1, r2, r3


class C3(T0):
    def  __init__(self):
        self._f13: bool = False
        self._f19: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f13(self):
        return self._f13
    @f13.setter
    def f13(self, value):
        self._f13 = value
    def ctor__(self, f19, f13):
        (self)._f19 = f19
        (self).f13 = f13

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife72_():
            coll30_ = _dafny.Map()
            compr_30_: int
            for compr_30_ in _dafny.IntegerRange(439, -49):
                d_705_v0_: int = compr_30_
                if ((439) <= (d_705_v0_)) and ((d_705_v0_) < (-49)):
                    coll30_[(d_705_v0_) + (-610)] = 836
            return _dafny.Map(coll30_)
        source14_ = D1_DC2(iife72_()
)
        if source14_.is_DC3:
            d_706___mcc_h0_ = source14_.cf4
            d_707_cf4_ = d_706___mcc_h0_
            return True
        elif source14_.is_DC4:
            d_708___mcc_h1_ = source14_.cf5
            d_709___mcc_h2_ = source14_.cf6
            d_710___mcc_h3_ = source14_.cf7
            d_711_cf7_ = d_710___mcc_h3_
            d_712_cf6_ = d_709___mcc_h2_
            d_713_cf5_ = d_708___mcc_h1_
            return self.f13
        elif source14_.is_DC5:
            d_714___mcc_h4_ = source14_.cf8
            d_715___mcc_h5_ = source14_.cf9
            d_716___mcc_h6_ = source14_.cf10
            d_717_cf10_ = d_716___mcc_h6_
            d_718_cf9_ = d_715___mcc_h5_
            d_719_cf8_ = d_714___mcc_h4_
            return False
        elif True:
            d_720___mcc_h7_ = source14_.cf3
            d_721_cf3_ = d_720___mcc_h7_
            return self.f13

    def fm5(self, globalState):
        if self.f13:
            return D1_DC3(len(_dafny.SeqWithoutIsStrInference([(self).f19, self.f13])))
        elif True:
            return D1_DC3(len(_dafny.Map({not(self.f13): (self).f19})))

    def fm27(self, globalState):
        return _dafny.CodePoint('d')

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        hi6_ = p0
        for d_722_i0_ in range(p0, hi6_):
            d_723_v0_: _dafny.Seq
            d_723_v0_ = _dafny.SeqWithoutIsStrInference([self.f13, (self).f19])
            (self).f13 = (d_723_v0_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_724_i1_ in range(default__.abs(625))])), len(d_723_v0_))]
            d_725_v1_: _dafny.Array
            nw102_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_725_v1_ = nw102_
            d_726_v2_: _dafny.Array
            nw103_ = _dafny.Array(False, 24)
            d_726_v2_ = nw103_
            index138_ = default__.safeIndex(376, (d_725_v1_).length(0))
            (d_725_v1_)[index138_] = d_726_v2_
            index139_ = default__.safeIndex(376, (d_725_v1_).length(0))
            def lambda22_(d_727_i2_):
                return False

            init12_ = lambda22_
            nw104_ = _dafny.Array(None, 4)
            for i0_12_ in range(nw104_.length(0)):
                nw104_[i0_12_] = init12_(i0_12_)
            (d_725_v1_)[index139_] = nw104_
            d_728_v3_: _dafny.Seq
            d_728_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rn"))
            d_728_v3_ = d_728_v3_
            if not((self).f19):
                d_729_v4_: int
                d_729_v4_ = 687
                d_730_v5_: _dafny.Set
                d_730_v5_ = _dafny.Set({(d_725_v1_)[default__.safeIndex(376, (d_725_v1_).length(0))], d_726_v2_, (d_725_v1_)[default__.safeIndex(376, (d_725_v1_).length(0))], (d_725_v1_)[default__.safeIndex(376, (d_725_v1_).length(0))]})
                d_729_v4_ = len((d_730_v5_) - (d_730_v5_))
                d_731_v6_: _dafny.Array
                def lambda23_(d_732_v3_):
                    def lambda24_(d_733_i3_):
                        return d_732_v3_

                    return lambda24_

                init13_ = lambda23_(d_728_v3_)
                nw105_ = _dafny.Array(None, 21)
                for i0_13_ in range(nw105_.length(0)):
                    nw105_[i0_13_] = init13_(i0_13_)
                d_731_v6_ = nw105_
                d_734_v7_: str
                d_734_v7_ = _dafny.CodePoint('n')
                index140_ = default__.safeIndex(759, (d_731_v6_).length(0))
                (d_731_v6_)[index140_] = (_dafny.SeqWithoutIsStrInference([(d_728_v3_)[default__.safeIndex(p0, len(d_728_v3_))] for d_735_i4_ in range(default__.abs(138))])).set(default__.safeIndex(d_729_v4_, len(_dafny.SeqWithoutIsStrInference([(d_728_v3_)[default__.safeIndex(p0, len(d_728_v3_))] for d_736_i4_ in range(default__.abs(138))]))), d_734_v7_)
                d_737_v8_: _dafny.Seq
                d_737_v8_ = _dafny.SeqWithoutIsStrInference([d_728_v3_])
                index141_ = default__.safeIndex(759, (d_731_v6_).length(0))
                (d_731_v6_)[index141_] = (d_737_v8_)[default__.safeIndex(p0, len(d_737_v8_))]
                d_738_v9_: _dafny.Map
                d_738_v9_ = _dafny.Map({(0) - (d_722_i0_): -33})
                d_729_v4_ = (len(d_738_v9_)) * (d_729_v4_)
                d_739_v10_: _dafny.Map
                d_739_v10_ = _dafny.Map({self.f13: p0})
                d_739_v10_ = (d_739_v10_).set(self.f13, -151)
                d_740_v11_: _dafny.Array
                nw106_ = _dafny.Array(int(0), 21)
                d_740_v11_ = nw106_
                d_741_v12_: C1
                nw107_ = C1()
                nw107_.ctor__((self).f19, d_728_v3_, d_740_v11_)
                d_741_v12_ = nw107_
            elif True:
                d_742_v13_: int
                d_742_v13_ = -760
                d_742_v13_ = (0) - ((len(_dafny.Map({self.f13: d_742_v13_}))) - (625))
                d_743_v14_: str
                d_743_v14_ = _dafny.CodePoint('j')
                d_744_v15_: _dafny.Map
                d_744_v15_ = _dafny.Map({d_743_v14_: (self).f19})
                (self).f13 = ((d_744_v15_)[d_743_v14_] if (d_743_v14_) in (d_744_v15_) else (self).f19)
                d_745_v16_: _dafny.Array
                def lambda25_(d_746_v13_):
                    def lambda26_(d_747_i5_):
                        return (d_747_i5_) + ((0) - (d_746_v13_))

                    return lambda26_

                init14_ = lambda25_(d_742_v13_)
                nw108_ = _dafny.Array(None, 6)
                for i0_14_ in range(nw108_.length(0)):
                    nw108_[i0_14_] = init14_(i0_14_)
                d_745_v16_ = nw108_
                index142_ = default__.safeIndex(678, (d_745_v16_).length(0))
                (d_745_v16_)[index142_] = -945
                d_748_v17_: D0
                d_748_v17_ = D0_DC1(self.f13, p0)
                d_749_v18_: T0
                nw109_ = C2()
                nw109_.ctor__(self.f13)
                d_749_v18_ = nw109_
                d_750_v19_: _dafny.Seq
                d_750_v19_ = _dafny.SeqWithoutIsStrInference([d_749_v18_])
                d_751_v20_: _dafny.Map
                d_751_v20_ = _dafny.Map({d_742_v13_: default__.fm28(len(d_750_v19_), d_742_v13_, self.f13, globalState)})
                d_752_v21_: C0
                nw110_ = C0()
                nw110_.ctor__(len(((d_751_v20_)[d_742_v13_] if (d_742_v13_) in (d_751_v20_) else d_728_v3_)))
                d_752_v21_ = nw110_
                d_753_v22_: _dafny.Map
                d_753_v22_ = _dafny.Map({(d_752_v21_).f15: len(d_728_v3_)})
                d_754_v23_: _dafny.Set
                d_754_v23_ = _dafny.Set({(self).fm4((self).f19, d_748_v17_, len(_dafny.Map({d_752_v21_: d_749_v18_.f13})), d_753_v22_, globalState)})
                index143_ = default__.safeIndex(678, (d_745_v16_).length(0))
                (d_745_v16_)[index143_] = len(d_754_v23_)
                d_755_v24_: D5
                d_755_v24_ = D5_DC14(not(False))
                index144_ = default__.safeIndex(678, (d_745_v16_).length(0))
                rhs107_ = d_755_v24_
                rhs108_ = (0) - (d_722_i0_)
                rhs109_ = self.f13
                lhs66_ = d_745_v16_
                lhs67_ = default__.safeIndex(678, (d_745_v16_).length(0))
                lhs68_ = self
                d_755_v24_ = rhs107_
                lhs66_[lhs67_] = rhs108_
                lhs68_.f13 = rhs109_
                d_756_v25_: _dafny.Map
                d_756_v25_ = _dafny.Map({(d_755_v24_).cf23: (d_745_v16_)[default__.safeIndex(678, (d_745_v16_).length(0))]})
                d_757_v26_: _dafny.MultiSet
                d_757_v26_ = _dafny.MultiSet([d_722_i0_, d_722_i0_])
                d_756_v25_ = (d_756_v25_).set((d_757_v26_).issubset(d_757_v26_), d_742_v13_)
        hi7_ = p0
        for d_758_i6_ in range(p0, hi7_):
            (self).f13 = self.f13
            d_759_v27_: D2
            d_759_v27_ = D2_DC8(self.f13, self.f13, self.f13)
            d_760_v28_: _dafny.Seq
            d_760_v28_ = _dafny.SeqWithoutIsStrInference([d_759_v27_, d_759_v27_])
            d_761_v29_: _dafny.Map
            d_761_v29_ = _dafny.Map({(self).f19: False})
            d_760_v28_ = _dafny.SeqWithoutIsStrInference([d_759_v27_, D2_DC8((self).f19, self.f13, ((d_761_v29_)[(self).f19] if ((self).f19) in (d_761_v29_) else self.f13)), d_759_v27_])
            d_762_v30_: int
            d_762_v30_ = 328
            d_762_v30_ = (0) - (default__.safeModuloInt(d_762_v30_, d_762_v30_))
            d_763_v31_: _dafny.Seq
            d_763_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orhdrkah"))
            d_764_v32_: _dafny.Array
            nw111_ = _dafny.Array(False, 2)
            d_764_v32_ = nw111_
            d_765_v33_: D6
            d_765_v33_ = D6_DC17((0) - (d_758_i6_), p0, d_764_v32_)
            d_766_v34_: D6
            d_766_v34_ = D6_DC18(d_765_v33_)
            d_767_v35_: _dafny.Map
            d_767_v35_ = _dafny.Map({d_766_v34_: self.f13})
            d_768_v36_: D1
            d_768_v36_ = D1_DC5(d_763_v31_, d_764_v32_, len(d_767_v35_))
            pat_let_tv18_ = d_762_v30_
            def iife73_(_pat_let21_0):
                def iife74_(d_769_dt__update__tmp_h0_):
                    def iife75_(_pat_let22_0):
                        def iife76_(d_770_dt__update_hcf8_h0_):
                            def iife77_(_pat_let23_0):
                                def iife78_(d_771_dt__update_hcf10_h0_):
                                    return D1_DC5(d_770_dt__update_hcf8_h0_, (d_769_dt__update__tmp_h0_).cf9, d_771_dt__update_hcf10_h0_)
                                return iife78_(_pat_let23_0)
                            return iife77_(pat_let_tv18_)
                        return iife76_(_pat_let22_0)
                    return iife75_((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))).set(default__.safeIndex(694, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")))), _dafny.CodePoint('s')))
                return iife74_(_pat_let21_0)
            d_768_v36_ = iife73_(d_768_v36_)
        d_772_v37_: D6
        d_772_v37_ = D6_DC16(default__.fm38(self.f13, globalState))
        d_773_v38_: D6
        d_773_v38_ = D6_DC18(d_772_v37_)
        d_774_v39_: D6
        d_774_v39_ = D6_DC18(d_772_v37_)
        d_774_v39_ = d_774_v39_
        d_775_v40_: str
        d_775_v40_ = _dafny.CodePoint('q')
        d_776_v41_: D7
        d_776_v41_ = D7_DC20(p0, d_775_v40_)
        d_777_v42_: D0
        d_777_v42_ = D0_DC1(not(self.f13), p0)
        d_778_v43_: _dafny.Seq
        d_778_v43_ = _dafny.SeqWithoutIsStrInference([p0])
        d_779_v44_: _dafny.Map
        d_779_v44_ = _dafny.Map({self.f13: _dafny.CodePoint('x')})
        d_780_v45_: _dafny.Array
        nw112_ = _dafny.Array(None, 16)
        nw112_[int(0)] = d_775_v40_
        nw112_[int(1)] = d_775_v40_
        nw112_[int(2)] = d_775_v40_
        nw112_[int(3)] = (d_776_v41_).cf32
        nw112_[int(4)] = _dafny.CodePoint('a')
        nw112_[int(5)] = d_775_v40_
        nw112_[int(6)] = d_775_v40_
        nw112_[int(7)] = (d_775_v40_ if self.f13 else d_775_v40_)
        nw112_[int(8)] = _dafny.CodePoint('h')
        nw112_[int(9)] = (d_775_v40_ if (self).fm4((self).f19, d_777_v42_, p0, default__.fm39(len(d_778_v43_), globalState), globalState) else d_775_v40_)
        nw112_[int(10)] = d_775_v40_
        nw112_[int(11)] = d_775_v40_
        nw112_[int(12)] = d_775_v40_
        nw112_[int(13)] = (d_775_v40_ if False else d_775_v40_)
        nw112_[int(14)] = (_dafny.CodePoint('e') if self.f13 else d_775_v40_)
        nw112_[int(15)] = ((d_779_v44_)[True] if (True) in (d_779_v44_) else d_775_v40_)
        d_780_v45_ = nw112_
        d_780_v45_ = d_780_v45_
        d_781_v46_: _dafny.Seq
        d_781_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvovm"))
        d_782_v47_: _dafny.Map
        d_782_v47_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fls"))): p0})
        d_783_v48_: _dafny.Seq
        d_783_v48_ = _dafny.SeqWithoutIsStrInference([not((self).f19)])
        d_784_v49_: _dafny.Set
        d_784_v49_ = _dafny.Set({(self).f19})
        d_785_v50_: _dafny.Array
        nw113_ = _dafny.Array(None, 6)
        nw113_[int(0)] = default__.fm0(d_782_v47_, p0, globalState)
        nw113_[int(1)] = (d_783_v48_)[default__.safeIndex(default__.fm31(globalState), len(d_783_v48_))]
        nw113_[int(2)] = (d_783_v48_)[default__.safeIndex(len(d_784_v49_), len(d_783_v48_))]
        nw113_[int(3)] = self.f13
        nw113_[int(4)] = (self).f19
        nw113_[int(5)] = self.f13
        d_785_v50_ = nw113_
        d_786_v51_: D1
        d_786_v51_ = D1_DC5((d_781_v46_) + (d_781_v46_), d_785_v50_, p0)
        source15_ = d_786_v51_
        if source15_.is_DC3:
            d_787___mcc_h0_ = source15_.cf4
            d_788_cf4_ = d_787___mcc_h0_
            d_789_v52_: D1
            d_789_v52_ = D1_DC4(d_785_v50_, (0) - (len(_dafny.SeqWithoutIsStrInference([d_788_cf4_ for d_790_i7_ in range(default__.abs(955))]))), d_775_v40_)
            d_791_v53_: _dafny.Set
            d_791_v53_ = _dafny.Set({d_789_v52_, d_789_v52_, d_789_v52_, d_789_v52_})
            d_792_v54_: _dafny.Seq
            d_792_v54_ = _dafny.SeqWithoutIsStrInference([d_791_v53_])
            if ((d_792_v54_)[default__.safeIndex(d_788_cf4_, len(d_792_v54_))]).ispropersubset((d_791_v53_) - (d_791_v53_)):
                d_793_v55_: _dafny.Map
                d_793_v55_ = _dafny.Map({not(False): p0})
                d_793_v55_ = (d_793_v55_).set(not(not (self.f13) or (self.f13)), len(_dafny.SeqWithoutIsStrInference([d_784_v49_, d_784_v49_])))
                d_794_v56_: _dafny.Array
                def lambda27_(d_795_cf4_):
                    def lambda28_(d_796_i8_):
                        return (d_796_i8_) - (d_795_cf4_)

                    return lambda28_

                init15_ = lambda27_(d_788_cf4_)
                nw114_ = _dafny.Array(None, 7)
                for i0_15_ in range(nw114_.length(0)):
                    nw114_[i0_15_] = init15_(i0_15_)
                d_794_v56_ = nw114_
                d_797_v57_: C1
                nw115_ = C1()
                nw115_.ctor__(self.f13, d_781_v46_, d_794_v56_)
                d_797_v57_ = nw115_
                d_798_v58_: _dafny.MultiSet
                d_798_v58_ = _dafny.MultiSet([(d_776_v41_).cf31])
                d_788_cf4_ = ((d_798_v58_)[d_788_cf4_] if (d_788_cf4_) in (d_798_v58_) else (-547) * (d_788_cf4_))
                d_799_v59_: _dafny.Array
                nw116_ = _dafny.Array(None, 3)
                nw116_[int(0)] = _dafny.SeqWithoutIsStrInference([d_775_v40_ for d_800_i9_ in range(default__.abs(239))])
                nw116_[int(1)] = d_781_v46_
                nw116_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
                d_799_v59_ = nw116_
                index145_ = default__.safeIndex(280, (d_799_v59_).length(0))
                (d_799_v59_)[index145_] = d_781_v46_
                index146_ = default__.safeIndex(280, (d_799_v59_).length(0))
                (d_799_v59_)[index146_] = (d_797_v57_.f18).set(default__.safeIndex(p0, len(d_797_v57_.f18)), d_775_v40_)
                rhs110_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bplas"))
                rhs111_ = d_797_v57_
                rhs112_ = d_778_v43_
                lhs69_ = d_797_v57_
                lhs69_.f18 = rhs110_
                d_797_v57_ = rhs111_
                d_778_v43_ = rhs112_
            elif True:
                d_801_v60_: _dafny.Array
                def lambda29_(d_802_cf4_):
                    def lambda30_(d_803_i10_):
                        return (d_803_i10_) * (d_802_cf4_)

                    return lambda30_

                init16_ = lambda29_(d_788_cf4_)
                nw117_ = _dafny.Array(None, 2)
                for i0_16_ in range(nw117_.length(0)):
                    nw117_[i0_16_] = init16_(i0_16_)
                d_801_v60_ = nw117_
                index147_ = default__.safeIndex(188, (d_801_v60_).length(0))
                (d_801_v60_)[index147_] = default__.fm31(globalState)
                index148_ = default__.safeIndex(188, (d_801_v60_).length(0))
                (d_801_v60_)[index148_] = p0
                d_804_v61_: _dafny.Map
                d_804_v61_ = _dafny.Map({d_801_v60_: d_801_v60_})
                d_804_v61_ = (d_804_v61_).set(d_801_v60_, (d_801_v60_ if self.f13 else d_801_v60_))
                d_805_v62_: _dafny.Set
                d_805_v62_ = _dafny.Set({(d_801_v60_)[default__.safeIndex(188, (d_801_v60_).length(0))]})
                d_806_v63_: D2
                d_806_v63_ = D2_DC7(d_781_v46_, d_788_cf4_, -989, self.f13)
                d_807_v64_: _dafny.MultiSet
                d_807_v64_ = _dafny.MultiSet([(d_805_v62_) - (_dafny.Set({(d_806_v63_).cf13, (d_801_v60_)[default__.safeIndex(188, (d_801_v60_).length(0))], (d_801_v60_)[default__.safeIndex(188, (d_801_v60_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_775_v40_ for d_808_i11_ in range(default__.abs(367))]))})), d_805_v62_])
                d_807_v64_ = d_807_v64_
                index149_ = default__.safeIndex(188, (d_801_v60_).length(0))
                (d_801_v60_)[index149_] = d_788_cf4_
                d_775_v40_ = _dafny.CodePoint('v')
            (self).f13 = (default__.fm31(globalState)) >= (p0)
            d_809_v65_: _dafny.Array
            nw118_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
            d_809_v65_ = nw118_
            index150_ = default__.safeIndex(721, (d_809_v65_).length(0))
            (d_809_v65_)[index150_] = (d_781_v46_ if self.f13 else d_781_v46_)
            index151_ = default__.safeIndex(721, (d_809_v65_).length(0))
            (d_809_v65_)[index151_] = (D2_DC7(d_781_v46_, -656, len(d_778_v43_), self.f13)).cf12
            d_810_v66_: _dafny.Array
            def lambda31_(d_811_p0_, d_812_cf4_):
                def lambda32_(d_813_i12_):
                    return _dafny.Map({d_811_p0_: d_812_cf4_})

                return lambda32_

            init17_ = lambda31_(p0, d_788_cf4_)
            nw119_ = _dafny.Array(None, 24)
            for i0_17_ in range(nw119_.length(0)):
                nw119_[i0_17_] = init17_(i0_17_)
            d_810_v66_ = nw119_
            d_810_v66_ = d_810_v66_
        elif source15_.is_DC4:
            d_814___mcc_h1_ = source15_.cf5
            d_815___mcc_h2_ = source15_.cf6
            d_816___mcc_h3_ = source15_.cf7
            d_817_cf7_ = d_816___mcc_h3_
            d_818_cf6_ = d_815___mcc_h2_
            d_819_cf5_ = d_814___mcc_h1_
            d_818_cf6_ = default__.safeModuloInt(p0, (default__.fm31(globalState) if (self).f19 else d_818_cf6_))
            d_820_v67_: _dafny.Array
            def lambda33_(d_821_i13_):
                return (d_821_i13_) * (862)

            init18_ = lambda33_
            nw120_ = _dafny.Array(None, 28)
            for i0_18_ in range(nw120_.length(0)):
                nw120_[i0_18_] = init18_(i0_18_)
            d_820_v67_ = nw120_
            d_822_v68_: _dafny.Seq
            d_822_v68_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([default__.fm31(globalState), len(d_784_v49_)]), _dafny.SeqWithoutIsStrInference([d_818_cf6_, p0]), d_778_v43_, d_778_v43_])
            d_823_v69_: _dafny.Map
            d_823_v69_ = _dafny.Map({p0: self.f13})
            d_824_v70_: _dafny.MultiSet
            d_824_v70_ = _dafny.MultiSet([self.f13, (self).f19])
            index152_ = default__.safeIndex(909, (d_820_v67_).length(0))
            (d_820_v67_)[index152_] = default__.safeDivisionInt((0) - ((0) - (((d_782_v47_)[p0] if (p0) in (d_782_v47_) else len(d_822_v68_)))), len((d_823_v69_).set((d_824_v70_).cardinality, (self).f19)))
            index153_ = default__.safeIndex(909, (d_820_v67_).length(0))
            (d_820_v67_)[index153_] = d_818_cf6_
            d_825_v71_: _dafny.Map
            d_825_v71_ = _dafny.Map({d_817_cf7_: d_785_v50_})
            d_825_v71_ = (d_825_v71_).set(d_817_cf7_, d_785_v50_)
            d_826_v72_: _dafny.MultiSet
            d_826_v72_ = _dafny.MultiSet([d_818_cf6_])
            d_827_v73_: _dafny.Map
            d_827_v73_ = _dafny.Map({(self).f19: (self).f19})
            if (_dafny.MultiSet([(0) - (len((d_827_v73_).set(self.f13, self.f13))), (d_820_v67_)[default__.safeIndex(909, (d_820_v67_).length(0))], d_818_cf6_])).issubset(d_826_v72_):
                d_828_v74_: D1
                d_828_v74_ = D1_DC4(d_785_v50_, d_818_cf6_, d_775_v40_)
                d_829_v75_: C2
                nw121_ = C2()
                nw121_.ctor__(self.f13)
                d_829_v75_ = nw121_
                d_830_v76_: _dafny.Seq
                d_830_v76_ = _dafny.SeqWithoutIsStrInference([d_829_v75_])
                index154_ = default__.safeIndex(909, (d_820_v67_).length(0))
                rhs113_ = d_817_cf7_
                rhs114_ = (d_828_v74_).cf6
                rhs115_ = ((890) * ((d_820_v67_)[default__.safeIndex(909, (d_820_v67_).length(0))])) < (948)
                rhs116_ = ((d_826_v72_) - (default__.fm40(len(d_830_v76_), self.f13, self.f13, (self).f19, globalState))) - (d_826_v72_)
                rhs117_ = (d_820_v67_)[default__.safeIndex(909, (d_820_v67_).length(0))]
                lhs70_ = self
                lhs71_ = d_820_v67_
                lhs72_ = default__.safeIndex(909, (d_820_v67_).length(0))
                d_817_cf7_ = rhs113_
                d_818_cf6_ = rhs114_
                lhs70_.f13 = rhs115_
                d_826_v72_ = rhs116_
                lhs71_[lhs72_] = rhs117_
                d_831_v77_: _dafny.Array
                nw122_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_831_v77_ = nw122_
                d_832_v78_: _dafny.Map
                d_832_v78_ = _dafny.Map({d_831_v77_: default__.fm41(p0, globalState)})
                d_832_v78_ = (d_832_v78_).set(d_831_v77_, D4_DC11(not((self).f19)))
                index155_ = default__.safeIndex(909, (d_820_v67_).length(0))
                (d_820_v67_)[index155_] = (d_820_v67_)[default__.safeIndex(909, (d_820_v67_).length(0))]
                d_833_v79_: _dafny.Map
                d_833_v79_ = _dafny.Map({d_826_v72_: (888) - (d_818_cf6_)})
                d_834_v80_: _dafny.Set
                d_834_v80_ = _dafny.Set({(d_820_v67_)[default__.safeIndex(909, (d_820_v67_).length(0))]})
                index156_ = default__.safeIndex(909, (d_820_v67_).length(0))
                (d_820_v67_)[index156_] = ((d_833_v79_)[default__.fm40((0) - ((d_820_v67_)[default__.safeIndex(909, (d_820_v67_).length(0))]), False, self.f13, self.f13, globalState)] if (default__.fm40((0) - ((d_820_v67_)[default__.safeIndex(909, (d_820_v67_).length(0))]), False, self.f13, self.f13, globalState)) in (d_833_v79_) else len(d_834_v80_))
                d_835_v81_: C2
                nw123_ = C2()
                nw123_.ctor__(self.f13)
                d_835_v81_ = nw123_
            elif True:
                index157_ = default__.safeIndex(909, (d_820_v67_).length(0))
                (d_820_v67_)[index157_] = (d_820_v67_)[default__.safeIndex(909, (d_820_v67_).length(0))]
                (self).f13 = (d_783_v48_)[default__.safeIndex((d_818_cf6_) + (len(d_781_v46_)), len(d_783_v48_))]
                (self).f13 = (p0) < (d_818_cf6_)
                d_836_v83_: _dafny.Map
                d_836_v83_ = _dafny.Map({d_781_v46_: d_775_v40_})
                d_837_v84_: _dafny.Map
                d_837_v84_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_775_v40_ for d_838_i14_ in range(default__.abs(150))]): (d_783_v48_)[default__.safeIndex(656, len(d_783_v48_))]})
                d_839_v85_: int
                out30_: int
                def iife79_():
                    coll31_ = _dafny.Map()
                    compr_31_: _dafny.Seq
                    for compr_31_ in (d_836_v83_).keys.Elements:
                        d_840_v82_: _dafny.Seq = compr_31_
                        if (d_840_v82_) in (d_836_v83_):
                            coll31_[d_840_v82_] = True
                    return _dafny.Map(coll31_)
                out30_ = default__.m0((iife79_()
                ) | (d_837_v84_), globalState)
                d_839_v85_ = out30_
                d_818_cf6_ = (0) - (p0)
        elif source15_.is_DC5:
            d_841___mcc_h4_ = source15_.cf8
            d_842___mcc_h5_ = source15_.cf9
            d_843___mcc_h6_ = source15_.cf10
            d_844_cf10_ = d_843___mcc_h6_
            d_845_cf9_ = d_842___mcc_h5_
            d_846_cf8_ = d_841___mcc_h4_
            (self).f13 = not((self).f19)
            d_847_v86_: _dafny.Map
            d_847_v86_ = _dafny.Map({808: d_785_v50_})
            d_848_v87_: _dafny.Set
            d_848_v87_ = _dafny.Set({d_845_cf9_, d_845_cf9_, ((d_847_v86_)[d_844_cf10_] if (d_844_cf10_) in (d_847_v86_) else d_845_cf9_)})
            d_848_v87_ = (_dafny.Set({d_785_v50_, d_785_v50_})) - (d_848_v87_)
            d_844_cf10_ = d_844_cf10_
            index158_ = default__.safeIndex(658, (d_785_v50_).length(0))
            (d_785_v50_)[index158_] = ((self).f19) or (self.f13)
            index159_ = default__.safeIndex(658, (d_785_v50_).length(0))
            rhs118_ = (d_845_cf9_ if not (self.f13) or (False) else d_785_v50_)
            rhs119_ = default__.safeModuloInt(d_844_cf10_, d_844_cf10_)
            rhs120_ = self.f13
            lhs73_ = d_785_v50_
            lhs74_ = default__.safeIndex(658, (d_785_v50_).length(0))
            d_845_cf9_ = rhs118_
            d_844_cf10_ = rhs119_
            lhs73_[lhs74_] = rhs120_
        elif True:
            d_849___mcc_h7_ = source15_.cf3
            d_850_cf3_ = d_849___mcc_h7_
            d_851_v88_: int
            d_851_v88_ = 120
            d_852_v89_: _dafny.Seq
            d_852_v89_ = _dafny.SeqWithoutIsStrInference([d_783_v48_])
            d_853_v90_: _dafny.Seq
            d_853_v90_ = (d_852_v89_)[default__.safeIndex(p0, len(d_852_v89_))]
            d_851_v88_ = len((d_853_v90_))
            d_854_v91_: C0
            nw124_ = C0()
            nw124_.ctor__(d_851_v88_)
            d_854_v91_ = nw124_
            d_855_v92_: _dafny.Array
            def lambda34_(d_856_v91_):
                def lambda35_(d_857_i15_):
                    return (d_857_i15_) * ((d_856_v91_).f15)

                return lambda35_

            init19_ = lambda34_(d_854_v91_)
            nw125_ = _dafny.Array(None, 6)
            for i0_19_ in range(nw125_.length(0)):
                nw125_[i0_19_] = init19_(i0_19_)
            d_855_v92_ = nw125_
            d_858_v93_: C1
            nw126_ = C1()
            nw126_.ctor__(self.f13, (d_781_v46_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyflgki"))), d_855_v92_)
            d_858_v93_ = nw126_
            d_851_v88_ = (d_854_v91_).f15
        d_859_v94_: D2
        d_859_v94_ = D2_DC8(True, self.f13, self.f13)
        def lambda36_(source16_):
            if source16_.is_DC7:
                d_860___mcc_h8_ = source16_.cf12
                d_861___mcc_h9_ = source16_.cf13
                d_862___mcc_h10_ = source16_.cf14
                d_863___mcc_h11_ = source16_.cf15
                d_864_cf15_ = d_863___mcc_h11_
                d_865_cf14_ = d_862___mcc_h10_
                d_866_cf13_ = d_861___mcc_h9_
                d_867_cf12_ = d_860___mcc_h8_
                return self.f13
            elif source16_.is_DC8:
                d_868___mcc_h12_ = source16_.cf16
                d_869___mcc_h13_ = source16_.cf17
                d_870___mcc_h14_ = source16_.cf18
                d_871_cf18_ = d_870___mcc_h14_
                d_872_cf17_ = d_869___mcc_h13_
                d_873_cf16_ = d_868___mcc_h12_
                return d_873_cf16_
            elif True:
                d_874___mcc_h15_ = source16_.cf11
                d_875_cf11_ = d_874___mcc_h15_
                return (self).f19

        (self).f13 = not(lambda36_(d_859_v94_))
        d_876_v95_: _dafny.Array
        nw127_ = _dafny.Array(int(0), 8)
        d_876_v95_ = nw127_
        d_877_v96_: D4
        d_877_v96_ = D4_DC10(d_876_v95_)
        r0 = (d_877_v96_).cf20
        return r0

    def m2(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        r2: bool = False
        r3: _dafny.Seq = _dafny.Seq({})
        d_878_i0_: int
        d_878_i0_ = 0
        with _dafny.label("8"):
            while False:
                with _dafny.c_label("8"):
                    if (d_878_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_878_i0_ = (d_878_i0_) + (1)
                    d_879_v0_: _dafny.Array
                    nw128_ = _dafny.Array(int(0), 26)
                    d_879_v0_ = nw128_
                    index160_ = default__.safeIndex(23, (d_879_v0_).length(0))
                    (d_879_v0_)[index160_] = p2
                    index161_ = default__.safeIndex(23, (d_879_v0_).length(0))
                    (d_879_v0_)[index161_] = p2
                    index162_ = default__.safeIndex(23, (d_879_v0_).length(0))
                    (d_879_v0_)[index162_] = (p0 if self.f13 else p2)
                    d_880_v1_: _dafny.Array
                    def lambda37_(d_881_i1_):
                        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f13, (self).f19]))) | (_dafny.MultiSet([self.f13]))

                    init20_ = lambda37_
                    nw129_ = _dafny.Array(None, 26)
                    for i0_20_ in range(nw129_.length(0)):
                        nw129_[i0_20_] = init20_(i0_20_)
                    d_880_v1_ = nw129_
                    d_882_v2_: _dafny.MultiSet
                    d_882_v2_ = _dafny.MultiSet([(self).f19, (self).f19])
                    index163_ = default__.safeIndex(769, (d_880_v1_).length(0))
                    (d_880_v1_)[index163_] = d_882_v2_
                    index164_ = default__.safeIndex(769, (d_880_v1_).length(0))
                    (d_880_v1_)[index164_] = (_dafny.MultiSet([True])) | (_dafny.MultiSet([self.f13]))
                    d_883_v3_: _dafny.Seq
                    d_883_v3_ = _dafny.SeqWithoutIsStrInference([p2])
                    d_884_v4_: _dafny.Map
                    d_884_v4_ = _dafny.Map({d_883_v3_: d_883_v3_})
                    (self).f13 = (d_883_v3_) not in (d_884_v4_)
                    pass
            pass
        d_885_v6_: _dafny.Seq
        d_885_v6_ = _dafny.SeqWithoutIsStrInference([785])
        d_886_v7_: _dafny.Map
        d_886_v7_ = _dafny.Map({self.f13: len(d_885_v6_)})
        d_887_v8_: _dafny.MultiSet
        d_887_v8_ = _dafny.MultiSet([len(d_886_v7_), (_dafny.MultiSet([(self).f19])).cardinality])
        d_888_v9_: _dafny.Map
        d_888_v9_ = _dafny.Map({p2: p2})
        def iife80_():
            coll32_ = _dafny.Map()
            compr_32_: int
            for compr_32_ in (d_887_v8_).Elements:
                d_889_v5_: int = compr_32_
                if (d_889_v5_) in (d_887_v8_):
                    coll32_[default__.safeDivisionInt(d_889_v5_, p2)] = p0
            return _dafny.Map(coll32_)
        r2 = (default__.fm0(iife80_()
        , len(d_888_v9_), globalState)) and ((self).f19)
        d_890_i2_: int
        d_890_i2_ = 0
        with _dafny.label("9"):
            while (p0) < (p0):
                with _dafny.c_label("9"):
                    if (d_890_i2_) >= (100):
                        raise _dafny.Break("9")
                    d_890_i2_ = (d_890_i2_) + (1)
                    d_891_v10_: D0
                    d_891_v10_ = D0_DC0(self.f13)
                    d_892_v11_: _dafny.Map
                    d_892_v11_ = _dafny.Map({p0: (d_891_v10_).cf0})
                    d_893_v13_: _dafny.Seq
                    d_893_v13_ = _dafny.SeqWithoutIsStrInference([p1, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_894_i3_ in range(default__.abs(-25))]), p1])
                    d_895_v15_: _dafny.Seq
                    d_895_v15_ = _dafny.SeqWithoutIsStrInference([self.f13, self.f13])
                    d_896_v16_: _dafny.Seq
                    d_896_v16_ = _dafny.SeqWithoutIsStrInference([d_895_v15_])
                    d_897_v17_: _dafny.Array
                    nw130_ = _dafny.Array(None, 18)
                    nw130_[int(0)] = p2
                    nw130_[int(1)] = p0
                    nw130_[int(2)] = p2
                    def iife81_():
                        def iife83_():
                            coll35_ = _dafny.Map()
                            compr_35_: _dafny.Seq
                            for compr_35_ in (d_893_v13_).Elements:
                                d_898_v12_: _dafny.Seq = compr_35_
                                if (d_898_v12_) in (d_893_v13_):
                                    coll35_[d_898_v12_] = (self).f19
                            return _dafny.Map(coll35_)
                        coll33_ = _dafny.Set()
                        def iife82_():
                            coll34_ = _dafny.Map()
                            compr_34_: _dafny.Seq
                            for compr_34_ in (d_893_v13_).Elements:
                                d_898_v12_: _dafny.Seq = compr_34_
                                if (d_898_v12_) in (d_893_v13_):
                                    coll34_[d_898_v12_] = (self).f19
                            return _dafny.Map(coll34_)
                        compr_33_: _dafny.Seq
                        for compr_33_ in (iife82_()
                        ).keys.Elements:
                            d_899_v14_: _dafny.Seq = compr_33_
                            if (d_899_v14_) in (iife83_()
                            ):
                                coll33_ = coll33_.union(_dafny.Set([d_899_v14_]))
                        return _dafny.Set(coll33_)
                    nw130_[int(3)] = len(iife81_()
                    )
                    nw130_[int(4)] = p0
                    nw130_[int(5)] = p0
                    nw130_[int(6)] = p0
                    nw130_[int(7)] = (0) - (p2)
                    nw130_[int(8)] = p2
                    nw130_[int(9)] = p2
                    nw130_[int(10)] = p0
                    nw130_[int(11)] = len(d_896_v16_)
                    nw130_[int(12)] = p2
                    nw130_[int(13)] = p2
                    nw130_[int(14)] = p0
                    nw130_[int(15)] = p2
                    nw130_[int(16)] = p2
                    nw130_[int(17)] = p0
                    d_897_v17_ = nw130_
                    d_900_v18_: C1
                    nw131_ = C1()
                    nw131_.ctor__(((d_892_v11_)[p2] if (p2) in (d_892_v11_) else default__.fm0(d_888_v9_, p0, globalState)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wceigad")), d_897_v17_)
                    d_900_v18_ = nw131_
                    d_901_v19_: str
                    d_901_v19_ = _dafny.CodePoint('p')
                    d_901_v19_ = d_901_v19_
                    (self).f13 = (d_900_v18_).f17
                    d_902_v20_: _dafny.Map
                    d_902_v20_ = _dafny.Map({d_895_v15_: (self).f19})
                    d_902_v20_ = (d_902_v20_).set(d_895_v15_, (self).f19)
                    pass
            pass
        d_903_v21_: D2
        d_903_v21_ = D2_DC6(d_885_v6_)
        pat_let_tv19_ = d_885_v6_
        pat_let_tv20_ = p2
        pat_let_tv21_ = d_885_v6_
        pat_let_tv22_ = d_887_v8_
        def lambda38_(source17_):
            if source17_.is_DC7:
                d_904___mcc_h0_ = source17_.cf12
                d_905___mcc_h1_ = source17_.cf13
                d_906___mcc_h2_ = source17_.cf14
                d_907___mcc_h3_ = source17_.cf15
                d_908_cf15_ = d_907___mcc_h3_
                d_909_cf14_ = d_906___mcc_h2_
                d_910_cf13_ = d_905___mcc_h1_
                d_911_cf12_ = d_904___mcc_h0_
                return self.f13
            elif source17_.is_DC8:
                d_912___mcc_h4_ = source17_.cf16
                d_913___mcc_h5_ = source17_.cf17
                d_914___mcc_h6_ = source17_.cf18
                d_915_cf18_ = d_914___mcc_h6_
                d_916_cf17_ = d_913___mcc_h5_
                d_917_cf16_ = d_912___mcc_h4_
                return (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g')]))) < ((pat_let_tv19_)[default__.safeIndex(pat_let_tv20_, len(pat_let_tv21_))])
            elif True:
                d_918___mcc_h7_ = source17_.cf11
                d_919_cf11_ = d_918___mcc_h7_
                return ((pat_let_tv22_).cardinality) < (75)

        if lambda38_(d_903_v21_):
            d_888_v9_ = (d_888_v9_).set(p0, 733)
            d_920_v22_: _dafny.Set
            d_920_v22_ = _dafny.Set({p0, 150})
            d_921_v23_: _dafny.Array
            def lambda39_(d_922_p0_):
                def lambda40_(d_923_i4_):
                    return (d_923_i4_) + (d_922_p0_)

                return lambda40_

            init21_ = lambda39_(p0)
            nw132_ = _dafny.Array(None, 12)
            for i0_21_ in range(nw132_.length(0)):
                nw132_[i0_21_] = init21_(i0_21_)
            d_921_v23_ = nw132_
            d_924_v24_: _dafny.Map
            d_924_v24_ = _dafny.Map({d_921_v23_: not((self).f19)})
            d_925_v25_: _dafny.Map
            d_925_v25_ = _dafny.Map({d_920_v22_: (d_924_v24_) | (d_924_v24_)})
            d_925_v25_ = (d_925_v25_).set((_dafny.Set({p0})) - (d_920_v22_), d_924_v24_)
            d_926_v27_: C1
            nw133_ = C1()
            def iife84_():
                coll36_ = _dafny.Map()
                compr_36_: int
                for compr_36_ in _dafny.IntegerRange(-765, 322):
                    d_927_v26_: int = compr_36_
                    if ((-765) <= (d_927_v26_)) and ((d_927_v26_) < (322)):
                        coll36_[(d_927_v26_) + ((d_887_v8_).cardinality)] = p0
                return _dafny.Map(coll36_)
            nw133_.ctor__(default__.fm0(iife84_()
            , (d_885_v6_)[default__.safeIndex(p2, len(d_885_v6_))], globalState), p1, d_921_v23_)
            d_926_v27_ = nw133_
            d_926_v27_ = d_926_v27_
            d_928_v28_: D4
            d_928_v28_ = D4_DC11((d_926_v27_).f17)
            d_929_v29_: _dafny.MultiSet
            d_929_v29_ = _dafny.MultiSet([d_928_v28_])
            d_930_v30_: _dafny.Seq
            d_930_v30_ = _dafny.SeqWithoutIsStrInference([d_929_v29_, d_929_v29_, _dafny.MultiSet([d_928_v28_]), d_929_v29_, d_929_v29_])
            d_931_v31_: _dafny.Array
            nw134_ = _dafny.Array(None, 12)
            nw134_[int(0)] = self.f13
            nw134_[int(1)] = False
            nw134_[int(2)] = (d_926_v27_).f17
            nw134_[int(3)] = ((d_926_v27_).f17) or ((self).f19)
            nw134_[int(4)] = (self).f19
            nw134_[int(5)] = self.f13
            nw134_[int(6)] = self.f13
            nw134_[int(7)] = self.f13
            nw134_[int(8)] = (d_926_v27_).f17
            nw134_[int(9)] = self.f13
            nw134_[int(10)] = (d_929_v29_).isdisjoint((d_930_v30_)[default__.safeIndex(p2, len(d_930_v30_))])
            nw134_[int(11)] = (d_926_v27_).f17
            d_931_v31_ = nw134_
            d_932_v32_: D5
            d_932_v32_ = D5_DC14(not((self).f19))
            index165_ = default__.safeIndex(289, (d_931_v31_).length(0))
            def iife85_(_pat_let24_0):
                def iife86_(d_933_dt__update__tmp_h0_):
                    def iife87_(_pat_let25_0):
                        def iife88_(d_934_dt__update_hcf23_h0_):
                            return D5_DC14(d_934_dt__update_hcf23_h0_)
                        return iife88_(_pat_let25_0)
                    return iife87_(self.f13)
                return iife86_(_pat_let24_0)
            (d_931_v31_)[index165_] = (iife85_(d_932_v32_)).cf23
            index166_ = default__.safeIndex(289, (d_931_v31_).length(0))
            (d_931_v31_)[index166_] = self.f13
            d_935_v33_: int
            d_935_v33_ = 470
            d_935_v33_ = 203
        elif True:
            d_936_v34_: _dafny.Map
            d_936_v34_ = _dafny.Map({self.f13: self.f13})
            rhs121_ = not(not (not(not(True))) or (not((self).f19)))
            rhs122_ = not(not(((len((d_936_v34_).set(self.f13, (self).f19))) >= (p2)) or (self.f13)))
            r2 = rhs121_
            r1 = rhs122_
            d_937_v35_: _dafny.Map
            d_937_v35_ = _dafny.Map({p1: default__.safeModuloInt(p0, p0)})
            d_938_v36_: str
            d_938_v36_ = _dafny.CodePoint('d')
            d_937_v35_ = (d_937_v35_).set(((p1).set(default__.safeIndex(p0, len(p1)), d_938_v36_)) + (p1), (0) - (p0))
            d_938_v36_ = default__.fm34(globalState)
            d_939_v38_: _dafny.Array
            def lambda41_(d_940_v36_):
                def lambda42_(d_941_i5_):
                    def iife89_():
                        coll37_ = _dafny.Map()
                        compr_37_: int
                        for compr_37_ in _dafny.IntegerRange(-622, 300):
                            d_942_v37_: int = compr_37_
                            if ((-622) <= (d_942_v37_)) and ((d_942_v37_) < (300)):
                                coll37_[default__.safeDivisionInt(d_942_v37_, len(_dafny.SeqWithoutIsStrInference([True])))] = 969
                        return _dafny.Map(coll37_)
                    return (d_941_i5_) + ((D7_DC20(len(iife89_()
), (D7_DC20(len(_dafny.SeqWithoutIsStrInference([True])), d_940_v36_)).cf32)).cf31)

                return lambda42_

            init22_ = lambda41_(d_938_v36_)
            nw135_ = _dafny.Array(None, 3)
            for i0_22_ in range(nw135_.length(0)):
                nw135_[i0_22_] = init22_(i0_22_)
            d_939_v38_ = nw135_
            d_943_v39_: _dafny.Set
            d_943_v39_ = _dafny.Set({d_939_v38_, d_939_v38_})
            rhs123_ = self.f13
            rhs124_ = not((d_943_v39_).isdisjoint((d_943_v39_) - (d_943_v39_)))
            r1 = rhs123_
            r2 = rhs124_
            d_944_v40_: _dafny.Array
            nw136_ = _dafny.Array(D7.default()(), 18)
            d_944_v40_ = nw136_
            d_945_v41_: T1
            nw137_ = C1()
            nw137_.ctor__((self).f19, p1, d_939_v38_)
            d_945_v41_ = nw137_
            d_946_v42_: D7
            d_946_v42_ = D7_DC19(d_945_v41_)
            d_947_v43_: D7
            d_947_v43_ = D7_DC21(d_946_v42_)
            index167_ = default__.safeIndex(637, (d_944_v40_).length(0))
            (d_944_v40_)[index167_] = d_947_v43_
            index168_ = default__.safeIndex(637, (d_944_v40_).length(0))
            (d_944_v40_)[index168_] = d_947_v43_
        d_948_v44_: _dafny.Seq
        d_948_v44_ = _dafny.SeqWithoutIsStrInference([self.f13])
        d_949_v45_: _dafny.Seq
        d_949_v45_ = _dafny.SeqWithoutIsStrInference([(d_948_v44_)[default__.safeIndex(p0, len(d_948_v44_))]])
        d_950_v46_: _dafny.Array
        nw138_ = _dafny.Array(None, 26)
        nw138_[int(0)] = (self).f19
        nw138_[int(1)] = self.f13
        nw138_[int(2)] = not((self).f19)
        nw138_[int(3)] = (d_949_v45_)[default__.safeIndex(p0, len(d_949_v45_))]
        nw138_[int(4)] = self.f13
        nw138_[int(5)] = (self).fm4((self).f19, default__.fm42(globalState), p0, _dafny.Map({p0: default__.fm31(globalState)}), globalState)
        nw138_[int(6)] = not((self).f19)
        nw138_[int(7)] = (p1) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_951_i6_ in range(default__.abs(379))]))
        nw138_[int(8)] = (self).f19
        nw138_[int(9)] = (self).f19
        nw138_[int(10)] = not(self.f13)
        nw138_[int(11)] = self.f13
        nw138_[int(12)] = self.f13
        nw138_[int(13)] = self.f13
        nw138_[int(14)] = self.f13
        nw138_[int(15)] = (self).f19
        nw138_[int(16)] = self.f13
        nw138_[int(17)] = (d_885_v6_) == (d_885_v6_)
        nw138_[int(18)] = self.f13
        nw138_[int(19)] = (self).f19
        nw138_[int(20)] = (self).f19
        nw138_[int(21)] = self.f13
        nw138_[int(22)] = False
        nw138_[int(23)] = self.f13
        nw138_[int(24)] = not((self).f19)
        nw138_[int(25)] = not(self.f13)
        d_950_v46_ = nw138_
        d_950_v46_ = d_950_v46_
        d_952_v47_: _dafny.Array
        nw139_ = _dafny.Array(int(0), 9)
        d_952_v47_ = nw139_
        index169_ = default__.safeIndex(662, (d_952_v47_).length(0))
        (d_952_v47_)[index169_] = p0
        index170_ = default__.safeIndex(662, (d_952_v47_).length(0))
        (d_952_v47_)[index170_] = default__.fm31(globalState)
        d_953_v48_: _dafny.Map
        d_953_v48_ = _dafny.Map({(d_952_v47_)[default__.safeIndex(662, (d_952_v47_).length(0))]: d_950_v46_})
        r0 = (d_953_v48_) | (d_953_v48_)
        r1 = ((True) or ((self).f19)) in (_dafny.MultiSet([self.f13, (self).f19]))
        r2 = (self).f19
        r3 = _dafny.SeqWithoutIsStrInference([True, ((d_952_v47_)[default__.safeIndex(662, (d_952_v47_).length(0))]) != (p2), (self).f19, (self.f13 if (self).f19 else True)])
        return r0, r1, r2, r3

    @property
    def f19(self):
        return self._f19

class C4(T0):
    def  __init__(self):
        self._f13: bool = False
        self._f16: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f13(self):
        return self._f13
    @f13.setter
    def f13(self, value):
        self._f13 = value
    def ctor__(self, f16, f13):
        (self)._f16 = f16
        (self).f13 = f13

    def fm4(self, p0, p1, p2, p3, globalState):
        return (self).f16

    def fm5(self, globalState):
        return D1_DC3(621)

    def fm12(self, p0, p1, globalState):
        return True

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        hi8_ = p0
        for d_954_i0_ in range(p0, hi8_):
            d_955_v0_: C0
            nw140_ = C0()
            nw140_.ctor__(d_954_i0_)
            d_955_v0_ = nw140_
            d_956_v1_: _dafny.Array
            def lambda43_(d_957_i1_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwoxu")) if self.f13 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxu")))

            init23_ = lambda43_
            nw141_ = _dafny.Array(None, 23)
            for i0_23_ in range(nw141_.length(0)):
                nw141_[i0_23_] = init23_(i0_23_)
            d_956_v1_ = nw141_
            d_958_v2_: _dafny.Seq
            d_958_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fylr"))
            index171_ = default__.safeIndex(314, (d_956_v1_).length(0))
            (d_956_v1_)[index171_] = d_958_v2_
            index172_ = default__.safeIndex(314, (d_956_v1_).length(0))
            (d_956_v1_)[index172_] = d_958_v2_
            d_956_v1_ = d_956_v1_
            d_959_v3_: _dafny.Seq
            d_959_v3_ = _dafny.SeqWithoutIsStrInference([(d_955_v0_).f15])
            d_960_v4_: _dafny.Seq
            d_960_v4_ = _dafny.SeqWithoutIsStrInference([self.f13])
            d_961_v5_: D0
            d_961_v5_ = D0_DC1(self.f13, p0)
            d_962_v6_: str
            d_962_v6_ = _dafny.CodePoint('k')
            d_963_v7_: _dafny.Map
            d_963_v7_ = _dafny.Map({p0: (self).f16})
            d_964_v8_: _dafny.Array
            nw142_ = _dafny.Array(None, 2)
            nw142_[int(0)] = d_963_v7_
            nw142_[int(1)] = _dafny.Map({default__.fm13((self).f16, (self).f16, d_962_v6_, globalState): self.f13})
            d_964_v8_ = nw142_
            d_965_v9_: _dafny.Map
            d_965_v9_ = _dafny.Map({d_964_v8_: (d_955_v0_).f15})
            d_966_v10_: _dafny.Array
            nw143_ = _dafny.Array(None, 9)
            nw143_[int(0)] = (0) - (p0)
            nw143_[int(1)] = len((d_960_v4_ if (self).f16 else d_960_v4_))
            nw143_[int(2)] = (d_955_v0_).f15
            nw143_[int(3)] = default__.safeDivisionInt(p0, default__.fm13((d_961_v5_).cf1, (self).f16, d_962_v6_, globalState))
            nw143_[int(4)] = ((d_965_v9_)[d_964_v8_] if (d_964_v8_) in (d_965_v9_) else p0)
            nw143_[int(5)] = 355
            nw143_[int(6)] = (p0) + ((d_955_v0_).f15)
            nw143_[int(7)] = p0
            nw143_[int(8)] = p0
            d_966_v10_ = nw143_
            index173_ = default__.safeIndex(219, (d_966_v10_).length(0))
            (d_966_v10_)[index173_] = (d_954_i0_) - (p0)
            d_967_v11_: int
            d_967_v11_ = 151
            d_968_v12_: T1
            nw144_ = C1()
            nw144_.ctor__((self).f16, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "luf")), d_966_v10_)
            d_968_v12_ = nw144_
            d_969_v13_: _dafny.MultiSet
            d_969_v13_ = _dafny.MultiSet([(self).f16, (self).f16])
            d_970_v14_: _dafny.Seq
            d_970_v14_ = _dafny.SeqWithoutIsStrInference([d_968_v12_])
            d_971_v15_: D2
            d_971_v15_ = D2_DC7(d_958_v2_, d_967_v11_, 518, (self).f16)
            index174_ = default__.safeIndex(219, (d_966_v10_).length(0))
            rhs125_ = d_959_v3_
            rhs126_ = (560) > (947)
            rhs127_ = (default__.safeDivisionInt((0) - (d_967_v11_), d_967_v11_)) - ((d_955_v0_).f15)
            rhs128_ = default__.safeDivisionInt(p0, (d_954_i0_) + ((d_968_v12_).fm7((d_955_v0_).f15, (d_969_v13_).cardinality, True, globalState)))
            rhs129_ = (d_970_v14_)[default__.safeIndex(((d_955_v0_).f15) - ((d_971_v15_).cf13), len(d_970_v14_))]
            lhs75_ = self
            lhs76_ = d_966_v10_
            lhs77_ = default__.safeIndex(219, (d_966_v10_).length(0))
            d_959_v3_ = rhs125_
            lhs75_.f13 = rhs126_
            lhs76_[lhs77_] = rhs127_
            d_967_v11_ = rhs128_
            d_968_v12_ = rhs129_
        d_972_v16_: str
        d_972_v16_ = _dafny.CodePoint('n')
        d_973_v17_: _dafny.Seq
        d_973_v17_ = _dafny.SeqWithoutIsStrInference([d_972_v16_, default__.fm22(self.f13, self.f13, globalState)])
        d_974_v18_: _dafny.Array
        nw145_ = _dafny.Array(int(0), 10)
        d_974_v18_ = nw145_
        d_975_v19_: C1
        nw146_ = C1()
        nw146_.ctor__(True, d_973_v17_, d_974_v18_)
        d_975_v19_ = nw146_
        d_976_v20_: _dafny.Seq
        d_976_v20_ = _dafny.SeqWithoutIsStrInference([d_975_v19_])
        d_977_v21_: D0
        d_977_v21_ = D0_DC1(self.f13, len(d_976_v20_))
        d_978_v22_: D2
        d_978_v22_ = D2_DC7(d_973_v17_, len(_dafny.Map({p0: self.f13})), len(_dafny.SeqWithoutIsStrInference([p0 for d_979_i2_ in range(default__.abs(18))])), not((d_977_v21_).cf1))
        d_980_v25_: _dafny.Map
        d_980_v25_ = _dafny.Map({393: d_978_v22_})
        d_981_v27_: _dafny.Map
        d_981_v27_ = _dafny.Map({len(_dafny.Map({p0: p0})): p0})
        d_982_v28_: _dafny.Array
        nw147_ = _dafny.Array(None, 29)
        nw147_[int(0)] = _dafny.Map({p0: d_978_v22_})
        def iife90_():
            def iife92_():
                coll40_ = _dafny.Map()
                compr_40_: int
                for compr_40_ in _dafny.IntegerRange(457, 35):
                    d_983_v24_: int = compr_40_
                    if ((457) <= (d_983_v24_)) and ((d_983_v24_) < (35)):
                        coll40_[(d_983_v24_) * (len(_dafny.Map({(_dafny.MultiSet([p0])).cardinality: (d_975_v19_).f17})))] = p0
                return _dafny.Map(coll40_)
            coll38_ = _dafny.Map()
            def iife91_():
                coll39_ = _dafny.Map()
                compr_39_: int
                for compr_39_ in _dafny.IntegerRange(457, 35):
                    d_983_v24_: int = compr_39_
                    if ((457) <= (d_983_v24_)) and ((d_983_v24_) < (35)):
                        coll39_[(d_983_v24_) * (len(_dafny.Map({(_dafny.MultiSet([p0])).cardinality: (d_975_v19_).f17})))] = p0
                return _dafny.Map(coll39_)
            compr_38_: int
            for compr_38_ in (iife91_()
            ).keys.Elements:
                d_984_v23_: int = compr_38_
                if (d_984_v23_) in (iife92_()
                ):
                    coll38_[default__.safeModuloInt(d_984_v23_, p0)] = d_978_v22_
            return _dafny.Map(coll38_)
        nw147_[int(1)] = iife90_()
        
        nw147_[int(2)] = d_980_v25_
        nw147_[int(3)] = _dafny.Map({p0: d_978_v22_})
        nw147_[int(4)] = d_980_v25_
        def iife93_():
            coll41_ = _dafny.Map()
            compr_41_: int
            for compr_41_ in _dafny.IntegerRange(500, -712):
                d_985_v26_: int = compr_41_
                if ((500) <= (d_985_v26_)) and ((d_985_v26_) < (-712)):
                    coll41_[default__.safeModuloInt(d_985_v26_, p0)] = d_978_v22_
            return _dafny.Map(coll41_)
        nw147_[int(5)] = iife93_()
        
        nw147_[int(6)] = d_980_v25_
        nw147_[int(7)] = d_980_v25_
        nw147_[int(8)] = d_980_v25_
        nw147_[int(9)] = d_980_v25_
        nw147_[int(10)] = (_dafny.Map({p0: d_978_v22_})).set(p0, d_978_v22_)
        nw147_[int(11)] = d_980_v25_
        nw147_[int(12)] = (_dafny.Map({p0: default__.fm23(len(d_975_v19_.f18), (self).f16, globalState)})) | (d_980_v25_)
        nw147_[int(13)] = _dafny.Map({p0: d_978_v22_})
        nw147_[int(14)] = d_980_v25_
        nw147_[int(15)] = (d_980_v25_) | (d_980_v25_)
        nw147_[int(16)] = (d_980_v25_) | (d_980_v25_)
        nw147_[int(17)] = d_980_v25_
        nw147_[int(18)] = d_980_v25_
        nw147_[int(19)] = d_980_v25_
        nw147_[int(20)] = d_980_v25_
        nw147_[int(21)] = d_980_v25_
        nw147_[int(22)] = (d_980_v25_) | (d_980_v25_)
        nw147_[int(23)] = (_dafny.Map({p0: d_978_v22_})).set(p0, d_978_v22_)
        nw147_[int(24)] = _dafny.Map({p0: d_978_v22_})
        nw147_[int(25)] = (d_980_v25_).set(p0, D2_DC7(d_975_v19_.f18, p0, len(_dafny.Set({(self).f16, (self).fm4((self).f16, d_977_v21_, p0, d_981_v27_, globalState)})), (self).f16))
        nw147_[int(26)] = d_980_v25_
        nw147_[int(27)] = (_dafny.Map({p0: d_978_v22_})) | (d_980_v25_)
        nw147_[int(28)] = d_980_v25_
        d_982_v28_ = nw147_
        index175_ = default__.safeIndex(495, (d_982_v28_).length(0))
        (d_982_v28_)[index175_] = d_980_v25_
        index176_ = default__.safeIndex(495, (d_982_v28_).length(0))
        (d_982_v28_)[index176_] = ((d_980_v25_) | (d_980_v25_)) | (_dafny.Map({p0: d_978_v22_}))
        d_986_v29_: _dafny.Map
        d_986_v29_ = _dafny.Map({(d_978_v22_).cf15: p0})
        (self).f13 = (p0) <= (len((d_986_v29_) | (d_986_v29_)))
        d_987_v30_: int
        d_987_v30_ = 888
        d_987_v30_ = d_987_v30_
        d_987_v30_ = (default__.safeModuloInt(p0, p0)) + (-301)
        hi9_ = p0
        for d_988_i3_ in range(d_987_v30_, hi9_):
            d_989_v31_: _dafny.Seq
            d_989_v31_ = _dafny.SeqWithoutIsStrInference([p0])
            (self).f13 = (len(d_986_v29_)) > ((d_989_v31_)[default__.safeIndex(d_987_v30_, len(d_989_v31_))])
            d_990_v32_: _dafny.Seq
            d_990_v32_ = _dafny.SeqWithoutIsStrInference([(d_975_v19_).f17])
            rhs130_ = (d_990_v32_) + (d_990_v32_)
            rhs131_ = (_dafny.MultiSet([((d_981_v27_)[len(d_981_v27_)] if (len(d_981_v27_)) in (d_981_v27_) else len(d_976_v20_)), d_988_i3_, p0])).cardinality
            d_990_v32_ = rhs130_
            d_987_v30_ = rhs131_
            (self).f13 = self.f13
            d_991_v33_: C1
            nw148_ = C1()
            nw148_.ctor__((d_975_v19_).fm14(p0, (d_975_v19_).f17, d_988_i3_, d_987_v30_, globalState), d_975_v19_.f18, d_974_v18_)
            d_991_v33_ = nw148_
        r0 = d_974_v18_
        return r0

    def m2(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        r2: bool = False
        r3: _dafny.Seq = _dafny.Seq({})
        d_992_v0_: _dafny.Map
        d_992_v0_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvnyliga"))})
        hi10_ = len(((d_992_v0_)[p2] if (p2) in (d_992_v0_) else p1))
        for d_993_i0_ in range(p2, hi10_):
            d_994_v1_: int
            d_994_v1_ = -188
            d_995_v2_: _dafny.MultiSet
            d_995_v2_ = _dafny.MultiSet([self.f13])
            rhs132_ = (d_995_v2_).cardinality
            rhs133_ = len((default__.fm24(d_994_v1_, self.f13, globalState)) | (_dafny.Map({self.f13: d_993_i0_})))
            rhs134_ = (self.f13) == ((self).f16)
            rhs135_ = self.f13
            lhs78_ = self
            d_994_v1_ = rhs132_
            d_994_v1_ = rhs133_
            r1 = rhs134_
            lhs78_.f13 = rhs135_
            d_996_v3_: _dafny.Array
            nw149_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_996_v3_ = nw149_
            d_996_v3_ = d_996_v3_
            d_997_v4_: _dafny.Array
            nw150_ = _dafny.Array(_dafny.CodePoint('D'), 5)
            d_997_v4_ = nw150_
            d_998_v5_: _dafny.Array
            nw151_ = _dafny.Array(_dafny.Map({}), 22)
            d_998_v5_ = nw151_
            d_999_v6_: _dafny.Map
            d_999_v6_ = _dafny.Map({p2: (0) - (d_993_i0_)})
            index177_ = default__.safeIndex(164, (d_998_v5_).length(0))
            (d_998_v5_)[index177_] = (d_999_v6_).set(p0, p2)
            d_1000_v7_: str
            d_1000_v7_ = _dafny.CodePoint('s')
            index178_ = default__.safeIndex(164, (d_998_v5_).length(0))
            rhs136_ = d_997_v4_
            rhs137_ = (d_999_v6_).set(default__.fm13((self).f16, True, d_1000_v7_, globalState), p2)
            rhs138_ = not(self.f13)
            rhs139_ = len((p1) + ((p1) + (p1)))
            rhs140_ = d_994_v1_
            lhs79_ = d_998_v5_
            lhs80_ = default__.safeIndex(164, (d_998_v5_).length(0))
            lhs81_ = self
            d_997_v4_ = rhs136_
            lhs79_[lhs80_] = rhs137_
            lhs81_.f13 = rhs138_
            d_994_v1_ = rhs139_
            d_994_v1_ = rhs140_
            d_1001_v8_: _dafny.Map
            d_1001_v8_ = _dafny.Map({((self).f16) == (not((self).f16)): (p1).set(default__.safeIndex(p2, len(p1)), d_1000_v7_)})
            d_1002_v9_: _dafny.Map
            d_1002_v9_ = _dafny.Map({(_dafny.MultiSet([d_993_i0_])).cardinality: _dafny.Set({d_994_v1_})})
            d_1003_v10_: _dafny.Array
            def lambda44_(d_1004_p2_):
                def lambda45_(d_1005_i1_):
                    return (d_1005_i1_) * (d_1004_p2_)

                return lambda45_

            init24_ = lambda44_(p2)
            nw152_ = _dafny.Array(None, 15)
            for i0_24_ in range(nw152_.length(0)):
                nw152_[i0_24_] = init24_(i0_24_)
            d_1003_v10_ = nw152_
            d_1006_v11_: _dafny.Map
            d_1006_v11_ = _dafny.Map({len(d_1002_v9_): d_1003_v10_})
            d_1001_v8_ = (d_1001_v8_).set((len(d_1006_v11_)) < (-440), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dct")))
        d_1007_v12_: _dafny.Array
        nw153_ = _dafny.Array(False, 3)
        d_1007_v12_ = nw153_
        index179_ = default__.safeIndex(780, (d_1007_v12_).length(0))
        (d_1007_v12_)[index179_] = False
        d_1008_v13_: int
        d_1008_v13_ = 539
        index180_ = default__.safeIndex(583, (d_1007_v12_).length(0))
        (d_1007_v12_)[index180_] = (self).f16
        d_1009_v14_: _dafny.Set
        d_1009_v14_ = _dafny.Set({True, self.f13})
        d_1010_v15_: _dafny.Set
        d_1010_v15_ = _dafny.Set({d_1008_v13_})
        d_1011_v16_: _dafny.Map
        d_1011_v16_ = _dafny.Map({(self).f16: self.f13})
        index181_ = default__.safeIndex(780, (d_1007_v12_).length(0))
        index182_ = default__.safeIndex(583, (d_1007_v12_).length(0))
        rhs141_ = self.f13
        rhs142_ = default__.safeDivisionInt(d_1008_v13_, -663)
        rhs143_ = ((d_1009_v14_) | (default__.fm25(d_1010_v15_, d_1011_v16_, p0, self.f13, globalState))) != ((d_1009_v14_) - (d_1009_v14_))
        rhs144_ = self.f13
        lhs82_ = d_1007_v12_
        lhs83_ = default__.safeIndex(780, (d_1007_v12_).length(0))
        lhs84_ = d_1007_v12_
        lhs85_ = default__.safeIndex(583, (d_1007_v12_).length(0))
        lhs86_ = self
        lhs82_[lhs83_] = rhs141_
        d_1008_v13_ = rhs142_
        lhs84_[lhs85_] = rhs143_
        lhs86_.f13 = rhs144_
        d_1008_v13_ = 326
        d_1012_v17_: D1
        d_1012_v17_ = D1_DC5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arq")), d_1007_v12_, 764)
        d_1013_v18_: D4
        d_1013_v18_ = D4_DC11((self).f16)
        d_1014_v19_: str
        d_1014_v19_ = _dafny.CodePoint('m')
        d_1015_v20_: _dafny.MultiSet
        d_1015_v20_ = _dafny.MultiSet([d_1008_v13_, len(d_1009_v14_), p0])
        d_1016_v21_: _dafny.Map
        d_1016_v21_ = _dafny.Map({p1: (0) - (d_1008_v13_)})
        d_1017_v22_: _dafny.Array
        def lambda46_(d_1018_i2_):
            return default__.safeDivisionInt(d_1018_i2_, 836)

        init25_ = lambda46_
        nw154_ = _dafny.Array(None, 25)
        for i0_25_ in range(nw154_.length(0)):
            nw154_[i0_25_] = init25_(i0_25_)
        d_1017_v22_ = nw154_
        d_1019_v23_: _dafny.Map
        d_1019_v23_ = _dafny.Map({(self).f16: d_1008_v13_})
        d_1020_v24_: D5
        d_1020_v24_ = D5_DC12(d_1019_v23_)
        d_1021_v25_: _dafny.Map
        d_1021_v25_ = _dafny.Map({d_1017_v22_: len((d_1020_v24_).cf22)})
        d_1022_v26_: _dafny.Seq
        d_1022_v26_ = _dafny.SeqWithoutIsStrInference([(self).f16])
        d_1023_v27_: _dafny.Seq
        d_1023_v27_ = _dafny.SeqWithoutIsStrInference([d_1008_v13_, p0])
        d_1024_v28_: _dafny.Array
        nw155_ = _dafny.Array(None, 26)
        nw155_[int(0)] = (d_1012_v17_).cf10
        nw155_[int(1)] = 168
        nw155_[int(2)] = p0
        nw155_[int(3)] = (len(p1)) + (590)
        nw155_[int(4)] = p0
        nw155_[int(5)] = default__.fm13(not((d_1013_v18_).cf21), self.f13, d_1014_v19_, globalState)
        nw155_[int(6)] = p0
        nw155_[int(7)] = (0) - (default__.safeDivisionInt((0) - ((d_1015_v20_).cardinality), len(d_1016_v21_)))
        nw155_[int(8)] = default__.safeModuloInt(default__.fm13((self).f16, self.f13, _dafny.CodePoint('d'), globalState), (0) - (p0))
        nw155_[int(9)] = p0
        nw155_[int(10)] = (0) - (default__.safeModuloInt(p2, len(default__.fm26(d_1014_v19_, p0, d_1008_v13_, globalState))))
        nw155_[int(11)] = p0
        nw155_[int(12)] = ((d_1021_v25_)[d_1017_v22_] if (d_1017_v22_) in (d_1021_v25_) else len(p1))
        nw155_[int(13)] = default__.safeDivisionInt(p2, len(d_1022_v26_))
        nw155_[int(14)] = p2
        nw155_[int(15)] = p0
        nw155_[int(16)] = 532
        nw155_[int(17)] = p2
        nw155_[int(18)] = p0
        nw155_[int(19)] = (d_1008_v13_) + (d_1008_v13_)
        nw155_[int(20)] = (707) * (p2)
        nw155_[int(21)] = (d_1008_v13_) + (d_1008_v13_)
        nw155_[int(22)] = p2
        nw155_[int(23)] = (0) - (((d_1023_v27_)[default__.safeIndex(p2, len(d_1023_v27_))]) + (p2))
        nw155_[int(24)] = 333
        nw155_[int(25)] = ((_dafny.MultiSet([d_1014_v19_, d_1014_v19_])).set(d_1014_v19_, default__.abs(d_1008_v13_))).cardinality
        d_1024_v28_ = nw155_
        index183_ = default__.safeIndex(659, (d_1017_v22_).length(0))
        (d_1017_v22_)[index183_] = (p2) * (len(d_1023_v27_))
        d_1025_v29_: _dafny.Map
        d_1025_v29_ = d_1011_v16_
        index184_ = default__.safeIndex(659, (d_1017_v22_).length(0))
        rhs145_ = p0
        rhs146_ = d_1025_v29_
        lhs87_ = d_1017_v22_
        lhs88_ = default__.safeIndex(659, (d_1017_v22_).length(0))
        lhs87_[lhs88_] = rhs145_
        d_1025_v29_ = rhs146_
        d_1026_v30_: _dafny.Array
        nw156_ = _dafny.Array(_dafny.Map({}), 17)
        d_1026_v30_ = nw156_
        d_1027_v31_: _dafny.Map
        d_1027_v31_ = _dafny.Map({d_1022_v26_: (d_1019_v23_).set((d_1007_v12_)[default__.safeIndex(780, (d_1007_v12_).length(0))], d_1008_v13_)})
        index185_ = default__.safeIndex(63, (d_1026_v30_).length(0))
        (d_1026_v30_)[index185_] = d_1027_v31_
        d_1028_v32_: D6
        d_1028_v32_ = D6_DC16(d_1027_v31_)
        d_1029_v33_: _dafny.Map
        d_1029_v33_ = _dafny.Map({(d_1007_v12_)[default__.safeIndex(780, (d_1007_v12_).length(0))]: p1})
        index186_ = default__.safeIndex(780, (d_1007_v12_).length(0))
        index187_ = default__.safeIndex(63, (d_1026_v30_).length(0))
        index188_ = default__.safeIndex(659, (d_1017_v22_).length(0))
        rhs147_ = not (self.f13) or ((d_1022_v26_)[default__.safeIndex(p2, len(d_1022_v26_))])
        rhs148_ = (d_1028_v32_).cf25
        rhs149_ = (0) - ((p0) - (((d_1017_v22_)[default__.safeIndex(659, (d_1017_v22_).length(0))]) + (len(d_1029_v33_))))
        rhs150_ = (p2) + (534)
        rhs151_ = d_1024_v28_
        lhs89_ = d_1007_v12_
        lhs90_ = default__.safeIndex(780, (d_1007_v12_).length(0))
        lhs91_ = d_1026_v30_
        lhs92_ = default__.safeIndex(63, (d_1026_v30_).length(0))
        lhs93_ = d_1017_v22_
        lhs94_ = default__.safeIndex(659, (d_1017_v22_).length(0))
        lhs89_[lhs90_] = rhs147_
        lhs91_[lhs92_] = rhs148_
        lhs93_[lhs94_] = rhs149_
        d_1008_v13_ = rhs150_
        d_1017_v22_ = rhs151_
        d_1009_v14_ = d_1009_v14_
        d_1030_v34_: _dafny.Seq
        d_1030_v34_ = _dafny.SeqWithoutIsStrInference([d_1007_v12_])
        d_1031_v35_: _dafny.Map
        d_1031_v35_ = _dafny.Map({len(d_1011_v16_): (d_1030_v34_)[default__.safeIndex(d_1008_v13_, len(d_1030_v34_))]})
        r0 = (d_1031_v35_ if (d_1007_v12_)[default__.safeIndex(780, (d_1007_v12_).length(0))] else d_1031_v35_)
        r1 = ((d_1022_v26_)[default__.safeIndex((_dafny.MultiSet(d_1022_v26_)).cardinality, len(d_1022_v26_))] if self.f13 else (d_1007_v12_)[default__.safeIndex(780, (d_1007_v12_).length(0))])
        d_1032_v36_: _dafny.MultiSet
        d_1032_v36_ = _dafny.MultiSet([(self).f16, self.f13, not((self).f16)])
        r2 = ((d_1032_v36_) - (d_1032_v36_)).issubset(d_1032_v36_)
        r3 = d_1022_v26_
        return r0, r1, r2, r3

    def m4(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_1033_v0_: D4
        d_1033_v0_ = D4_DC11(p1)
        d_1034_v1_: str
        d_1034_v1_ = _dafny.CodePoint('o')
        d_1035_v2_: _dafny.Seq
        d_1035_v2_ = _dafny.SeqWithoutIsStrInference([d_1034_v1_])
        d_1036_v3_: int
        d_1036_v3_ = 956
        d_1037_v4_: D2
        d_1037_v4_ = D2_DC7(d_1035_v2_, d_1036_v3_, 593, p0)
        pat_let_tv23_ = d_1037_v4_
        def iife94_(_pat_let26_0):
            def iife95_(d_1038_dt__update__tmp_h0_):
                def iife96_(_pat_let27_0):
                    def iife97_(d_1039_dt__update_hcf21_h0_):
                        return D4_DC11(d_1039_dt__update_hcf21_h0_)
                    return iife97_(_pat_let27_0)
                return iife96_((pat_let_tv23_).cf15)
            return iife95_(_pat_let26_0)
        source18_ = iife94_(d_1033_v0_)
        if source18_.is_DC11:
            d_1040___mcc_h0_ = source18_.cf21
            d_1041_cf21_ = d_1040___mcc_h0_
            d_1042_v5_: _dafny.Array
            nw157_ = _dafny.Array(int(0), 3)
            d_1042_v5_ = nw157_
            index189_ = default__.safeIndex(739, (d_1042_v5_).length(0))
            (d_1042_v5_)[index189_] = d_1036_v3_
            index190_ = default__.safeIndex(739, (d_1042_v5_).length(0))
            (d_1042_v5_)[index190_] = default__.safeModuloInt(default__.fm13(d_1041_cf21_, d_1041_cf21_, d_1034_v1_, globalState), (0) - ((0) - ((d_1036_v3_) * (d_1036_v3_))))
            d_1043_v6_: _dafny.Set
            d_1043_v6_ = _dafny.Set({p1, self.f13, d_1041_cf21_})
            d_1043_v6_ = d_1043_v6_
            d_1044_v7_: _dafny.Map
            d_1044_v7_ = _dafny.Map({self.f13: not(p1)})
            d_1044_v7_ = (d_1044_v7_).set(p1, d_1041_cf21_)
            d_1045_v8_: C1
            nw158_ = C1()
            nw158_.ctor__(d_1041_cf21_, (d_1035_v2_).set(default__.safeIndex(len(d_1044_v7_), len(d_1035_v2_)), _dafny.CodePoint('c')), d_1042_v5_)
            d_1045_v8_ = nw158_
        elif True:
            d_1046___mcc_h1_ = source18_.cf20
            d_1047_cf20_ = d_1046___mcc_h1_
            d_1048_v9_: _dafny.Array
            def lambda47_(d_1049_p1_):
                def lambda48_(d_1050_i0_):
                    return d_1049_p1_

                return lambda48_

            init26_ = lambda47_(p1)
            nw159_ = _dafny.Array(None, 10)
            for i0_26_ in range(nw159_.length(0)):
                nw159_[i0_26_] = init26_(i0_26_)
            d_1048_v9_ = nw159_
            d_1051_v10_: D1
            d_1051_v10_ = D1_DC4(d_1048_v9_, d_1036_v3_, default__.fm22((self).f16, p0, globalState))
            d_1052_v11_: _dafny.Array
            nw160_ = _dafny.Array(None, 28)
            nw160_[int(0)] = d_1034_v1_
            nw160_[int(1)] = d_1034_v1_
            nw160_[int(2)] = d_1034_v1_
            nw160_[int(3)] = (d_1051_v10_).cf7
            nw160_[int(4)] = d_1034_v1_
            nw160_[int(5)] = d_1034_v1_
            nw160_[int(6)] = d_1034_v1_
            nw160_[int(7)] = _dafny.CodePoint('u')
            nw160_[int(8)] = _dafny.CodePoint('k')
            nw160_[int(9)] = d_1034_v1_
            nw160_[int(10)] = d_1034_v1_
            nw160_[int(11)] = _dafny.CodePoint('o')
            nw160_[int(12)] = _dafny.CodePoint('p')
            nw160_[int(13)] = d_1034_v1_
            nw160_[int(14)] = d_1034_v1_
            nw160_[int(15)] = _dafny.CodePoint('s')
            nw160_[int(16)] = d_1034_v1_
            nw160_[int(17)] = d_1034_v1_
            nw160_[int(18)] = (d_1035_v2_)[default__.safeIndex(d_1036_v3_, len(d_1035_v2_))]
            nw160_[int(19)] = d_1034_v1_
            nw160_[int(20)] = (d_1035_v2_)[default__.safeIndex(d_1036_v3_, len(d_1035_v2_))]
            nw160_[int(21)] = d_1034_v1_
            nw160_[int(22)] = d_1034_v1_
            nw160_[int(23)] = d_1034_v1_
            nw160_[int(24)] = d_1034_v1_
            nw160_[int(25)] = d_1034_v1_
            nw160_[int(26)] = (d_1051_v10_).cf7
            nw160_[int(27)] = d_1034_v1_
            d_1052_v11_ = nw160_
            index191_ = default__.safeIndex(262, (d_1052_v11_).length(0))
            (d_1052_v11_)[index191_] = default__.fm22(self.f13, p0, globalState)
            index192_ = default__.safeIndex(262, (d_1052_v11_).length(0))
            rhs152_ = (not(((self).f16 if self.f13 else p1))) or ((p1) and ((self).f16))
            rhs153_ = d_1034_v1_
            lhs95_ = self
            lhs96_ = d_1052_v11_
            lhs97_ = default__.safeIndex(262, (d_1052_v11_).length(0))
            lhs95_.f13 = rhs152_
            lhs96_[lhs97_] = rhs153_
            d_1053_v12_: _dafny.Map
            d_1053_v12_ = _dafny.Map({d_1036_v3_: d_1036_v3_})
            d_1054_v13_: T0
            nw161_ = C2()
            nw161_.ctor__(True)
            d_1054_v13_ = nw161_
            r1 = ((len(_dafny.Map({d_1053_v12_: d_1054_v13_}))) + (len(d_1035_v2_))) > (459)
            d_1055_v14_: _dafny.Map
            d_1055_v14_ = _dafny.Map({d_1036_v3_: (self).f16})
            d_1056_v15_: _dafny.Map
            d_1056_v15_ = _dafny.Map({((d_1055_v14_)[(0) - (d_1036_v3_)] if ((0) - (d_1036_v3_)) in (d_1055_v14_) else p0): default__.safeDivisionInt(d_1036_v3_, d_1036_v3_)})
            d_1056_v15_ = (d_1056_v15_).set(d_1054_v13_.f13, default__.safeDivisionInt(d_1036_v3_, d_1036_v3_))
            d_1055_v14_ = d_1055_v14_
        (self).f13 = self.f13
        d_1057_v16_: _dafny.Seq
        d_1057_v16_ = _dafny.SeqWithoutIsStrInference([not(self.f13), p1, self.f13, (self).f16])
        d_1058_v17_: _dafny.Array
        def lambda49_(d_1059_i1_):
            return self.f13

        init27_ = lambda49_
        nw162_ = _dafny.Array(None, 11)
        for i0_27_ in range(nw162_.length(0)):
            nw162_[i0_27_] = init27_(i0_27_)
        d_1058_v17_ = nw162_
        d_1060_v18_: _dafny.Map
        d_1060_v18_ = _dafny.Map({d_1057_v16_: d_1058_v17_})
        d_1060_v18_ = (d_1060_v18_).set(_dafny.SeqWithoutIsStrInference([(d_1057_v16_)[default__.safeIndex(d_1036_v3_, len(d_1057_v16_))], (self).f16, not((self).f16)]), d_1058_v17_)
        d_1061_i2_: int
        d_1061_i2_ = 0
        with _dafny.label("10"):
            while p1:
                with _dafny.c_label("10"):
                    if (d_1061_i2_) >= (100):
                        raise _dafny.Break("10")
                    d_1061_i2_ = (d_1061_i2_) + (1)
                    if (self).f16:
                        d_1062_v19_: _dafny.Map
                        d_1062_v19_ = _dafny.Map({self.f13: d_1036_v3_})
                        d_1063_v20_: _dafny.Map
                        d_1063_v20_ = _dafny.Map({(self).f16: len(d_1062_v19_)})
                        d_1064_v21_: _dafny.Set
                        d_1064_v21_ = _dafny.Set({True, not(p0)})
                        d_1065_v22_: _dafny.Set
                        d_1065_v22_ = _dafny.Set({d_1036_v3_, len(d_1064_v21_)})
                        d_1066_v23_: _dafny.Map
                        d_1066_v23_ = _dafny.Map({self.f13: p0})
                        rhs154_ = (d_1057_v16_)[default__.safeIndex(default__.safeModuloInt(d_1036_v3_, d_1036_v3_), len(d_1057_v16_))]
                        rhs155_ = ((len(d_1062_v19_) if p0 else d_1036_v3_)) * (default__.safeDivisionInt(d_1036_v3_, d_1036_v3_))
                        rhs156_ = len(_dafny.Map({_dafny.MultiSet([d_1036_v3_, -641, len(d_1035_v2_), d_1036_v3_]): default__.fm25(d_1065_v22_, d_1066_v23_, len(_dafny.SeqWithoutIsStrInference([self.f13, False, (self).f16])), p0, globalState)}))
                        r0 = rhs154_
                        d_1036_v3_ = rhs155_
                        r2 = rhs156_
                        d_1067_v24_: D0
                        d_1067_v24_ = D0_DC1((self).f16, -634)
                        d_1068_v25_: _dafny.Map
                        d_1068_v25_ = _dafny.Map({(0) - (d_1036_v3_): d_1036_v3_})
                        d_1069_v26_: _dafny.Map
                        d_1069_v26_ = _dafny.Map({len(d_1035_v2_): len(d_1068_v25_)})
                        d_1070_v27_: _dafny.Seq
                        d_1070_v27_ = _dafny.SeqWithoutIsStrInference([d_1036_v3_])
                        d_1034_v1_ = default__.fm22((self).fm4((d_1057_v16_)[default__.safeIndex(d_1036_v3_, len(d_1057_v16_))], d_1067_v24_, d_1036_v3_, (d_1068_v25_).set(d_1036_v3_, d_1036_v3_), globalState), (d_1036_v3_) >= ((d_1070_v27_)[default__.safeIndex(d_1036_v3_, len(d_1070_v27_))]), globalState)
                        d_1071_v28_: _dafny.Array
                        nw163_ = _dafny.Array(int(0), 8)
                        d_1071_v28_ = nw163_
                        index193_ = default__.safeIndex(711, (d_1071_v28_).length(0))
                        (d_1071_v28_)[index193_] = d_1036_v3_
                        index194_ = default__.safeIndex(711, (d_1071_v28_).length(0))
                        rhs157_ = ((len(d_1035_v2_)) - (d_1036_v3_)) < (len(d_1063_v20_))
                        rhs158_ = self.f13
                        rhs159_ = d_1036_v3_
                        lhs98_ = self
                        lhs99_ = d_1071_v28_
                        lhs100_ = default__.safeIndex(711, (d_1071_v28_).length(0))
                        lhs98_.f13 = rhs157_
                        r0 = rhs158_
                        lhs99_[lhs100_] = rhs159_
                        d_1070_v27_ = _dafny.SeqWithoutIsStrInference([(d_1071_v28_)[default__.safeIndex(711, (d_1071_v28_).length(0))] for d_1072_i3_ in range(default__.abs(303))])
                        d_1073_v29_: _dafny.MultiSet
                        d_1073_v29_ = _dafny.MultiSet([((d_1071_v28_)[default__.safeIndex(711, (d_1071_v28_).length(0))]) - (d_1036_v3_), 741, (d_1071_v28_)[default__.safeIndex(711, (d_1071_v28_).length(0))], d_1036_v3_])
                        d_1074_v30_: T1
                        nw164_ = C1()
                        nw164_.ctor__(p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nldyh")), d_1071_v28_)
                        d_1074_v30_ = nw164_
                        d_1075_v31_: _dafny.Seq
                        d_1075_v31_ = _dafny.SeqWithoutIsStrInference([d_1074_v30_, d_1074_v30_])
                        d_1076_v32_: D7
                        d_1076_v32_ = D7_DC19(d_1074_v30_)
                        d_1077_v33_: _dafny.Array
                        nw165_ = _dafny.Array(None, 14)
                        nw165_[int(0)] = d_1074_v30_
                        nw165_[int(1)] = (d_1075_v31_)[default__.safeIndex(d_1036_v3_, len(d_1075_v31_))]
                        nw165_[int(2)] = d_1074_v30_
                        nw165_[int(3)] = (d_1076_v32_).cf30
                        nw165_[int(4)] = d_1074_v30_
                        nw165_[int(5)] = d_1074_v30_
                        nw165_[int(6)] = d_1074_v30_
                        nw165_[int(7)] = d_1074_v30_
                        nw165_[int(8)] = d_1074_v30_
                        nw165_[int(9)] = d_1074_v30_
                        nw165_[int(10)] = d_1074_v30_
                        nw165_[int(11)] = d_1074_v30_
                        nw165_[int(12)] = d_1074_v30_
                        nw165_[int(13)] = d_1074_v30_
                        d_1077_v33_ = nw165_
                        d_1078_v34_: _dafny.Seq
                        d_1078_v34_ = _dafny.SeqWithoutIsStrInference([d_1077_v33_])
                        rhs160_ = d_1034_v1_
                        rhs161_ = (d_1057_v16_ if p1 else d_1057_v16_)
                        rhs162_ = (d_1035_v2_) <= (d_1035_v2_)
                        rhs163_ = (d_1073_v29_).set(d_1036_v3_, default__.abs(len(d_1064_v21_)))
                        rhs164_ = d_1078_v34_
                        d_1034_v1_ = rhs160_
                        d_1057_v16_ = rhs161_
                        r1 = rhs162_
                        d_1073_v29_ = rhs163_
                        d_1078_v34_ = rhs164_
                    elif True:
                        d_1079_v35_: _dafny.Array
                        nw166_ = _dafny.Array(None, 4)
                        nw166_[int(0)] = d_1036_v3_
                        nw166_[int(1)] = 330
                        nw166_[int(2)] = d_1036_v3_
                        nw166_[int(3)] = default__.safeDivisionInt(d_1036_v3_, d_1036_v3_)
                        d_1079_v35_ = nw166_
                        d_1080_v36_: _dafny.Map
                        d_1080_v36_ = _dafny.Map({d_1079_v35_: p0})
                        index195_ = default__.safeIndex(831, (d_1079_v35_).length(0))
                        (d_1079_v35_)[index195_] = len(d_1080_v36_)
                        d_1081_v37_: _dafny.Map
                        d_1081_v37_ = _dafny.Map({self.f13: (d_1057_v16_)[default__.safeIndex(d_1036_v3_, len(d_1057_v16_))]})
                        d_1082_v38_: _dafny.Map
                        d_1082_v38_ = _dafny.Map({d_1081_v37_: d_1036_v3_})
                        index196_ = default__.safeIndex(831, (d_1079_v35_).length(0))
                        (d_1079_v35_)[index196_] = len((d_1082_v38_).set(d_1081_v37_, d_1036_v3_))
                        d_1083_v39_: C1
                        nw167_ = C1()
                        nw167_.ctor__(p0, d_1035_v2_, d_1079_v35_)
                        d_1083_v39_ = nw167_
                        d_1084_v40_: _dafny.Map
                        d_1084_v40_ = _dafny.Map({not(self.f13): len(d_1035_v2_)})
                        d_1085_v41_: _dafny.Seq
                        d_1085_v41_ = _dafny.SeqWithoutIsStrInference([d_1036_v3_, 431])
                        d_1086_v42_: _dafny.MultiSet
                        d_1086_v42_ = _dafny.MultiSet([len(d_1083_v39_.f18)])
                        d_1087_v43_: _dafny.Seq
                        d_1087_v43_ = _dafny.SeqWithoutIsStrInference([(d_1085_v41_)[default__.safeIndex(53, len(d_1085_v41_))], ((d_1086_v42_)[len(d_1035_v2_)] if (len(d_1035_v2_)) in (d_1086_v42_) else (d_1079_v35_)[default__.safeIndex(831, (d_1079_v35_).length(0))]), (d_1079_v35_)[default__.safeIndex(831, (d_1079_v35_).length(0))], d_1036_v3_])
                        rhs165_ = ((d_1035_v2_ if not((self).f16) else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiwrcowth"))).set(default__.safeIndex(len(d_1084_v40_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiwrcowth")))), d_1034_v1_))) + (d_1083_v39_.f18)
                        rhs166_ = not(not(p1))
                        rhs167_ = (d_1087_v43_)[default__.safeIndex(d_1036_v3_, len(d_1087_v43_))]
                        d_1035_v2_ = rhs165_
                        r0 = rhs166_
                        d_1036_v3_ = rhs167_
                        r2 = (d_1079_v35_)[default__.safeIndex(831, (d_1079_v35_).length(0))]
                        d_1086_v42_ = (_dafny.MultiSet(d_1085_v41_)).set(d_1036_v3_, default__.abs(((d_1079_v35_)[default__.safeIndex(831, (d_1079_v35_).length(0))]) - (31)))
                    d_1088_v44_: D1
                    d_1088_v44_ = D1_DC4(d_1058_v17_, d_1036_v3_, d_1034_v1_)
                    d_1089_v45_: _dafny.Seq
                    d_1089_v45_ = _dafny.SeqWithoutIsStrInference([d_1035_v2_, (d_1035_v2_).set(default__.safeIndex(d_1036_v3_, len(d_1035_v2_)), _dafny.CodePoint('l')), d_1035_v2_])
                    d_1090_v46_: _dafny.Array
                    nw168_ = _dafny.Array(None, 11)
                    nw168_[int(0)] = (d_1036_v3_) - ((d_1088_v44_).cf6)
                    nw168_[int(1)] = 609
                    nw168_[int(2)] = default__.safeModuloInt(d_1036_v3_, len(_dafny.SeqWithoutIsStrInference([(d_1037_v4_).cf13])))
                    nw168_[int(3)] = d_1036_v3_
                    nw168_[int(4)] = d_1036_v3_
                    nw168_[int(5)] = (0) - (d_1036_v3_)
                    nw168_[int(6)] = default__.safeDivisionInt(d_1036_v3_, len(d_1057_v16_))
                    nw168_[int(7)] = (d_1036_v3_) * (d_1036_v3_)
                    nw168_[int(8)] = (d_1036_v3_) * (d_1036_v3_)
                    nw168_[int(9)] = (0) - ((d_1036_v3_) * (len(d_1089_v45_)))
                    nw168_[int(10)] = d_1036_v3_
                    d_1090_v46_ = nw168_
                    index197_ = default__.safeIndex(646, (d_1090_v46_).length(0))
                    (d_1090_v46_)[index197_] = len(d_1035_v2_)
                    d_1091_v47_: D6
                    d_1091_v47_ = D6_DC17(d_1036_v3_, d_1036_v3_, d_1058_v17_)
                    d_1092_v48_: _dafny.Set
                    d_1092_v48_ = _dafny.Set({d_1091_v47_})
                    index198_ = default__.safeIndex(646, (d_1090_v46_).length(0))
                    (d_1090_v46_)[index198_] = len((d_1092_v48_) | (d_1092_v48_))
                    d_1093_v49_: _dafny.MultiSet
                    d_1093_v49_ = _dafny.MultiSet([d_1034_v1_])
                    d_1094_v50_: _dafny.Seq
                    d_1094_v50_ = _dafny.SeqWithoutIsStrInference([default__.fm31(globalState)])
                    rhs168_ = 120
                    rhs169_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcdkyipa")))
                    rhs170_ = _dafny.MultiSet(default__.fm28((len(d_1035_v2_) if self.f13 else (d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]), default__.fm31(globalState), (d_1094_v50_) != (d_1094_v50_), globalState))
                    rhs171_ = d_1090_v46_
                    d_1036_v3_ = rhs168_
                    d_1036_v3_ = rhs169_
                    d_1093_v49_ = rhs170_
                    d_1090_v46_ = rhs171_
                    d_1095_v51_: _dafny.Seq
                    d_1095_v51_ = _dafny.SeqWithoutIsStrInference([d_1094_v50_])
                    d_1096_v52_: D2
                    d_1096_v52_ = D2_DC6((d_1095_v51_)[default__.safeIndex(default__.fm31(globalState), len(d_1095_v51_))])
                    source19_ = d_1096_v52_
                    if source19_.is_DC7:
                        d_1097___mcc_h2_ = source19_.cf12
                        d_1098___mcc_h3_ = source19_.cf13
                        d_1099___mcc_h4_ = source19_.cf14
                        d_1100___mcc_h5_ = source19_.cf15
                        d_1101_cf15_ = d_1100___mcc_h5_
                        d_1102_cf14_ = d_1099___mcc_h4_
                        d_1103_cf13_ = d_1098___mcc_h3_
                        d_1104_cf12_ = d_1097___mcc_h2_
                        d_1105_v53_: _dafny.Map
                        d_1105_v53_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajufne"))) + (d_1035_v2_): (d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]})
                        d_1105_v53_ = d_1105_v53_
                        index199_ = default__.safeIndex(268, (d_1058_v17_).length(0))
                        (d_1058_v17_)[index199_] = default__.fm0(_dafny.Map({(d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]: (0) - (d_1036_v3_)}), (d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))], globalState)
                        d_1106_v54_: _dafny.Map
                        d_1106_v54_ = _dafny.Map({(0) - ((d_1093_v49_).cardinality): self.f13})
                        d_1107_v55_: _dafny.Set
                        d_1107_v55_ = _dafny.Set({d_1057_v16_})
                        d_1108_v56_: C2
                        nw169_ = C2()
                        nw169_.ctor__(True)
                        d_1108_v56_ = nw169_
                        d_1109_v57_: _dafny.Map
                        d_1109_v57_ = _dafny.Map({d_1108_v56_: (d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]})
                        index200_ = default__.safeIndex(646, (d_1090_v46_).length(0))
                        index201_ = default__.safeIndex(268, (d_1058_v17_).length(0))
                        rhs172_ = (d_1101_cf15_) and (p1)
                        rhs173_ = default__.safeModuloInt(d_1103_cf13_, (d_1103_cf13_) - ((d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]))
                        rhs174_ = ((d_1036_v3_) - (635)) not in (d_1094_v50_)
                        rhs175_ = ((d_1107_v55_).intersection(d_1107_v55_)).ispropersubset(_dafny.Set({d_1057_v16_, d_1057_v16_}))
                        rhs176_ = (d_1106_v54_) | ((_dafny.Map({len((d_1095_v51_)[default__.safeIndex(((d_1109_v57_)[d_1108_v56_] if (d_1108_v56_) in (d_1109_v57_) else (d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]), len(d_1095_v51_))]): p0})) | (d_1106_v54_))
                        lhs101_ = d_1090_v46_
                        lhs102_ = default__.safeIndex(646, (d_1090_v46_).length(0))
                        lhs103_ = d_1058_v17_
                        lhs104_ = default__.safeIndex(268, (d_1058_v17_).length(0))
                        r1 = rhs172_
                        lhs101_[lhs102_] = rhs173_
                        lhs103_[lhs104_] = rhs174_
                        r0 = rhs175_
                        d_1106_v54_ = rhs176_
                        index202_ = default__.safeIndex(646, (d_1090_v46_).length(0))
                        (d_1090_v46_)[index202_] = (default__.fm13(p1, (self).f16, d_1034_v1_, globalState)) - (-352)
                        d_1110_v58_: _dafny.Array
                        def lambda50_(d_1111_v16_, d_1112_v17_):
                            def lambda51_(d_1113_i4_):
                                return _dafny.SeqWithoutIsStrInference([d_1111_v16_, _dafny.SeqWithoutIsStrInference([(d_1112_v17_)[default__.safeIndex(268, (d_1112_v17_).length(0))]])])

                            return lambda51_

                        init28_ = lambda50_(d_1057_v16_, d_1058_v17_)
                        nw170_ = _dafny.Array(None, 15)
                        for i0_28_ in range(nw170_.length(0)):
                            nw170_[i0_28_] = init28_(i0_28_)
                        d_1110_v58_ = nw170_
                        d_1110_v58_ = d_1110_v58_
                    elif source19_.is_DC8:
                        d_1114___mcc_h6_ = source19_.cf16
                        d_1115___mcc_h7_ = source19_.cf17
                        d_1116___mcc_h8_ = source19_.cf18
                        d_1117_cf18_ = d_1116___mcc_h8_
                        d_1118_cf17_ = d_1115___mcc_h7_
                        d_1119_cf16_ = d_1114___mcc_h6_
                        d_1120_v59_: _dafny.Map
                        d_1120_v59_ = _dafny.Map({d_1117_cf18_: (d_1093_v49_).cardinality})
                        d_1121_v60_: _dafny.MultiSet
                        d_1121_v60_ = _dafny.MultiSet([d_1036_v3_, ((d_1120_v59_)[d_1117_cf18_] if (d_1117_cf18_) in (d_1120_v59_) else (d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]), (d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]])
                        index203_ = default__.safeIndex(646, (d_1090_v46_).length(0))
                        (d_1090_v46_)[index203_] = (d_1121_v60_).cardinality
                        d_1122_v61_: _dafny.MultiSet
                        d_1122_v61_ = _dafny.MultiSet([d_1035_v2_, (d_1035_v2_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpqdcyt"))), (d_1089_v45_)[default__.safeIndex((0) - (-673), len(d_1089_v45_))]])
                        d_1122_v61_ = ((_dafny.MultiSet(d_1089_v45_)).intersection(default__.fm43(globalState))).intersection(d_1122_v61_)
                        d_1123_v62_: _dafny.Map
                        d_1123_v62_ = _dafny.Map({d_1036_v3_: 767})
                        d_1124_v63_: D1
                        d_1124_v63_ = D1_DC2(d_1123_v62_)
                        d_1125_v64_: _dafny.Seq
                        d_1125_v64_ = _dafny.SeqWithoutIsStrInference([d_1124_v63_, d_1124_v63_])
                        d_1125_v64_ = default__.fm44(d_1036_v3_, False, 704, (d_1035_v2_) != (_dafny.SeqWithoutIsStrInference([d_1034_v1_ for d_1126_i5_ in range(default__.abs(330))])), globalState)
                        index204_ = default__.safeIndex(646, (d_1090_v46_).length(0))
                        (d_1090_v46_)[index204_] = (d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]
                    elif True:
                        d_1127___mcc_h9_ = source19_.cf11
                        d_1128_cf11_ = d_1127___mcc_h9_
                        index205_ = default__.safeIndex(646, (d_1090_v46_).length(0))
                        (d_1090_v46_)[index205_] = (d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]
                        d_1129_v65_: _dafny.Seq
                        d_1129_v65_ = _dafny.SeqWithoutIsStrInference([d_1057_v16_, d_1057_v16_])
                        d_1130_v66_: D9
                        d_1130_v66_ = D9_DC23(_dafny.MultiSet((d_1129_v65_)[default__.safeIndex(-343, len(d_1129_v65_))]))
                        d_1131_v67_: _dafny.MultiSet
                        d_1131_v67_ = _dafny.MultiSet([(self).f16])
                        d_1132_v68_: C3
                        nw171_ = C3()
                        nw171_.ctor__(not((d_1131_v67_).ispropersubset((d_1130_v66_).cf35)), self.f13)
                        d_1132_v68_ = nw171_
                        d_1133_v69_: D1
                        d_1133_v69_ = D1_DC5((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "caoyl"))).set(default__.safeIndex(435, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "caoyl")))), d_1034_v1_), d_1058_v17_, d_1036_v3_)
                        index206_ = default__.safeIndex(646, (d_1090_v46_).length(0))
                        (d_1090_v46_)[index206_] = ((d_1133_v69_).cf10) - ((0) - ((0) - ((d_1128_cf11_)[default__.safeIndex(len(d_1035_v2_), len(d_1128_cf11_))])))
                        d_1134_v70_: D5
                        d_1134_v70_ = D5_DC14(p0)
                        d_1135_v71_: D5
                        d_1135_v71_ = D5_DC15(d_1134_v70_)
                        d_1136_v72_: _dafny.MultiSet
                        d_1136_v72_ = _dafny.MultiSet([d_1036_v3_, (d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]])
                        d_1137_v73_: _dafny.MultiSet
                        d_1137_v73_ = _dafny.MultiSet([d_1136_v72_, d_1136_v72_])
                        index207_ = default__.safeIndex(770, (d_1058_v17_).length(0))
                        (d_1058_v17_)[index207_] = (d_1137_v73_).ispropersubset(d_1137_v73_)
                        d_1138_v74_: D5
                        d_1138_v74_ = D5_DC12(_dafny.Map({(self).f16: d_1036_v3_}))
                        d_1139_v75_: _dafny.Map
                        d_1139_v75_ = _dafny.Map({len(d_1035_v2_): d_1058_v17_})
                        d_1140_v76_: _dafny.Map
                        d_1140_v76_ = _dafny.Map({476: ((d_1139_v75_)[(d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]] if ((d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]) in (d_1139_v75_) else d_1058_v17_)})
                        d_1141_v77_: _dafny.Map
                        d_1141_v77_ = _dafny.Map({d_1058_v17_: d_1058_v17_})
                        d_1142_v78_: _dafny.Array
                        nw172_ = _dafny.Array(None, 21)
                        nw172_[int(0)] = d_1058_v17_
                        nw172_[int(1)] = d_1058_v17_
                        nw172_[int(2)] = d_1058_v17_
                        nw172_[int(3)] = d_1058_v17_
                        nw172_[int(4)] = d_1058_v17_
                        nw172_[int(5)] = d_1058_v17_
                        nw172_[int(6)] = (d_1133_v69_).cf9
                        nw172_[int(7)] = (d_1058_v17_ if self.f13 else d_1058_v17_)
                        nw172_[int(8)] = ((d_1140_v76_)[(d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]] if ((d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))]) in (d_1140_v76_) else d_1058_v17_)
                        nw172_[int(9)] = d_1058_v17_
                        nw172_[int(10)] = d_1058_v17_
                        nw172_[int(11)] = d_1058_v17_
                        nw172_[int(12)] = d_1058_v17_
                        nw172_[int(13)] = d_1058_v17_
                        nw172_[int(14)] = d_1058_v17_
                        nw172_[int(15)] = d_1058_v17_
                        nw172_[int(16)] = ((d_1141_v77_)[d_1058_v17_] if (d_1058_v17_) in (d_1141_v77_) else d_1058_v17_)
                        nw172_[int(17)] = d_1058_v17_
                        nw172_[int(18)] = d_1058_v17_
                        nw172_[int(19)] = d_1058_v17_
                        nw172_[int(20)] = d_1058_v17_
                        d_1142_v78_ = nw172_
                        index208_ = default__.safeIndex(119, (d_1142_v78_).length(0))
                        (d_1142_v78_)[index208_] = d_1058_v17_
                        d_1143_v79_: _dafny.Array
                        nw173_ = _dafny.Array(None, 2)
                        nw173_[int(0)] = default__.fm32(d_1035_v2_, (self).f16, d_1036_v3_, (d_1090_v46_)[default__.safeIndex(646, (d_1090_v46_).length(0))], globalState)
                        nw173_[int(1)] = d_1035_v2_
                        d_1143_v79_ = nw173_
                        index209_ = default__.safeIndex(375, (d_1143_v79_).length(0))
                        (d_1143_v79_)[index209_] = d_1035_v2_
                        d_1144_v80_: _dafny.Map
                        d_1144_v80_ = _dafny.Map({((d_1035_v2_) + (d_1035_v2_)).set(default__.safeIndex(251, len((d_1035_v2_) + (d_1035_v2_))), d_1034_v1_): not((d_1035_v2_) < (d_1035_v2_))})
                        index210_ = default__.safeIndex(770, (d_1058_v17_).length(0))
                        index211_ = default__.safeIndex(119, (d_1142_v78_).length(0))
                        index212_ = default__.safeIndex(375, (d_1143_v79_).length(0))
                        rhs177_ = d_1135_v71_
                        rhs178_ = ((d_1144_v80_)[_dafny.SeqWithoutIsStrInference([d_1034_v1_ for d_1145_i6_ in range(default__.abs(70))])] if (_dafny.SeqWithoutIsStrInference([d_1034_v1_ for d_1146_i6_ in range(default__.abs(70))])) in (d_1144_v80_) else (self).f16)
                        rhs179_ = d_1138_v74_
                        rhs180_ = d_1058_v17_
                        rhs181_ = d_1035_v2_
                        lhs105_ = d_1058_v17_
                        lhs106_ = default__.safeIndex(770, (d_1058_v17_).length(0))
                        lhs107_ = d_1142_v78_
                        lhs108_ = default__.safeIndex(119, (d_1142_v78_).length(0))
                        lhs109_ = d_1143_v79_
                        lhs110_ = default__.safeIndex(375, (d_1143_v79_).length(0))
                        d_1135_v71_ = rhs177_
                        lhs105_[lhs106_] = rhs178_
                        d_1138_v74_ = rhs179_
                        lhs107_[lhs108_] = rhs180_
                        lhs109_[lhs110_] = rhs181_
                    pass
            pass
        d_1147_v81_: _dafny.Map
        d_1147_v81_ = _dafny.Map({p1: True})
        d_1148_v82_: _dafny.Seq
        d_1148_v82_ = (d_1057_v16_).set(default__.safeIndex(len(d_1147_v81_), len(d_1057_v16_)), self.f13)
        d_1149_v83_: _dafny.Array
        nw174_ = _dafny.Array(None, 15)
        nw174_[int(0)] = d_1148_v82_
        nw174_[int(1)] = d_1148_v82_
        nw174_[int(2)] = d_1148_v82_
        nw174_[int(3)] = d_1148_v82_
        nw174_[int(4)] = d_1148_v82_
        nw174_[int(5)] = d_1148_v82_
        nw174_[int(6)] = d_1148_v82_
        nw174_[int(7)] = default__.fm45(p0, 67, globalState)
        nw174_[int(8)] = d_1148_v82_
        nw174_[int(9)] = d_1148_v82_
        nw174_[int(10)] = d_1148_v82_
        nw174_[int(11)] = default__.fm45(self.f13, d_1036_v3_, globalState)
        nw174_[int(12)] = d_1148_v82_
        nw174_[int(13)] = d_1148_v82_
        nw174_[int(14)] = d_1148_v82_
        d_1149_v83_ = nw174_
        index213_ = default__.safeIndex(992, (d_1149_v83_).length(0))
        (d_1149_v83_)[index213_] = d_1148_v82_
        index214_ = default__.safeIndex(992, (d_1149_v83_).length(0))
        (d_1149_v83_)[index214_] = d_1148_v82_
        d_1150_v84_: _dafny.Set
        d_1150_v84_ = _dafny.Set({d_1036_v3_, d_1036_v3_})
        d_1151_v85_: _dafny.Set
        d_1151_v85_ = _dafny.Set({(self).f16})
        d_1152_v86_: _dafny.Map
        d_1152_v86_ = _dafny.Map({self.f13: d_1151_v85_})
        d_1153_v87_: _dafny.Seq
        d_1153_v87_ = _dafny.SeqWithoutIsStrInference([d_1151_v85_, d_1151_v85_])
        d_1154_v88_: _dafny.MultiSet
        d_1154_v88_ = _dafny.MultiSet([len(d_1035_v2_)])
        d_1155_v89_: _dafny.Map
        d_1155_v89_ = _dafny.Map({d_1154_v88_: p1})
        d_1156_v90_: _dafny.Array
        nw175_ = _dafny.Array(None, 23)
        nw175_[int(0)] = _dafny.Set({p1, (d_1057_v16_)[default__.safeIndex(d_1036_v3_, len(d_1057_v16_))]})
        nw175_[int(1)] = (default__.fm25(d_1150_v84_, d_1147_v81_, d_1036_v3_, self.f13, globalState) if ((d_1147_v81_)[p0] if (p0) in (d_1147_v81_) else self.f13) else _dafny.Set({p1}))
        nw175_[int(2)] = ((d_1152_v86_)[not(self.f13)] if (not(self.f13)) in (d_1152_v86_) else _dafny.Set({p0}))
        nw175_[int(3)] = d_1151_v85_
        nw175_[int(4)] = (_dafny.Set({(self).f16, p0}) if True else d_1151_v85_)
        nw175_[int(5)] = d_1151_v85_
        nw175_[int(6)] = d_1151_v85_
        nw175_[int(7)] = d_1151_v85_
        nw175_[int(8)] = (d_1153_v87_)[default__.safeIndex(d_1036_v3_, len(d_1153_v87_))]
        nw175_[int(9)] = d_1151_v85_
        nw175_[int(10)] = d_1151_v85_
        nw175_[int(11)] = (d_1151_v85_ if ((d_1155_v89_)[_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([915 for d_1157_i7_ in range(default__.abs(707))]))] if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([915 for d_1158_i7_ in range(default__.abs(707))]))) in (d_1155_v89_) else p0) else d_1151_v85_)
        nw175_[int(12)] = (d_1151_v85_).intersection(d_1151_v85_)
        nw175_[int(13)] = d_1151_v85_
        nw175_[int(14)] = (_dafny.Set({p1, (self).f16}) if p1 else d_1151_v85_)
        nw175_[int(15)] = d_1151_v85_
        nw175_[int(16)] = d_1151_v85_
        nw175_[int(17)] = d_1151_v85_
        nw175_[int(18)] = (d_1151_v85_) - (d_1151_v85_)
        nw175_[int(19)] = (d_1151_v85_ if self.f13 else _dafny.Set({(self).f16}))
        nw175_[int(20)] = d_1151_v85_
        nw175_[int(21)] = _dafny.Set({True})
        nw175_[int(22)] = d_1151_v85_
        d_1156_v90_ = nw175_
        index215_ = default__.safeIndex(465, (d_1156_v90_).length(0))
        (d_1156_v90_)[index215_] = (d_1151_v85_).intersection(d_1151_v85_)
        index216_ = default__.safeIndex(465, (d_1156_v90_).length(0))
        rhs182_ = d_1058_v17_
        rhs183_ = ((d_1151_v85_) | (_dafny.Set({(d_1057_v16_)[default__.safeIndex(d_1036_v3_, len(d_1057_v16_))], self.f13, p0, (self).f16}))) | (d_1151_v85_)
        lhs111_ = d_1156_v90_
        lhs112_ = default__.safeIndex(465, (d_1156_v90_).length(0))
        d_1058_v17_ = rhs182_
        lhs111_[lhs112_] = rhs183_
        r0 = (d_1035_v2_) <= (d_1035_v2_)
        r1 = (d_1151_v85_).ispropersubset(((d_1156_v90_)[default__.safeIndex(465, (d_1156_v90_).length(0))]).intersection(d_1151_v85_))
        r2 = (len(d_1147_v81_)) + (d_1036_v3_)
        return r0, r1, r2

    @property
    def f16(self):
        return self._f16

class C5(T0, T1):
    def  __init__(self):
        self._f13: bool = False
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f13(self):
        return self._f13
    @f13.setter
    def f13(self, value):
        self._f13 = value
    @property
    def f14(self):
        return self._f14
    def ctor__(self, f13, f14):
        (self).f13 = f13
        (self)._f14 = f14

    def fm4(self, p0, p1, p2, p3, globalState):
        return ((790 if self.f13 else len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([self.f13])), len(_dafny.Map({self.f13: self.f13}))])), len(_dafny.SeqWithoutIsStrInference([-516, 198, len(_dafny.Set({_dafny.CodePoint('m'), _dafny.CodePoint('q')})), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flisrw")))), len(_dafny.Set({self.f13}))]))])))) < (len(_dafny.Map({757: -211})))

    def fm5(self, globalState):
        return D1_DC3(default__.safeModuloInt(len(_dafny.Set({self.f13})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1159_i0_ in range(default__.abs(346))]))))

    def fm6(self, globalState):
        return not((D0_DC1(self.f13, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "box"))))).cf1)

    def fm7(self, p0, p1, p2, globalState):
        return (default__.safeDivisionInt(457, -669)) - (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nuqc"))), len(_dafny.Map({_dafny.CodePoint('o'): D0_DC1(False, 114)}))))

    def fm8(self, p0, p1, globalState):
        return (_dafny.CodePoint('a')) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rq")))

    def fm9(self, p0, globalState):
        return default__.safeModuloInt((0) - ((-694 if self.f13 else (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f13, not(self.f13), self.f13]))).cardinality)), (-159) * ((_dafny.MultiSet([not(self.f13), self.f13])).cardinality))

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_1160_v0_: _dafny.Array
        nw176_ = _dafny.Array(None, 1)
        nw176_[int(0)] = not(not(self.f13))
        d_1160_v0_ = nw176_
        d_1161_v1_: D1
        d_1161_v1_ = D1_DC4(d_1160_v0_, p0, _dafny.CodePoint('i'))
        source20_ = d_1161_v1_
        if source20_.is_DC3:
            d_1162___mcc_h0_ = source20_.cf4
            d_1163_cf4_ = d_1162___mcc_h0_
            d_1164_v2_: _dafny.MultiSet
            d_1164_v2_ = _dafny.MultiSet([d_1163_cf4_])
            d_1163_cf4_ = (self).fm7((self).fm9(p0, globalState), default__.safeDivisionInt(p0, ((d_1164_v2_).set(p0, default__.abs(p0))).cardinality), self.f13, globalState)
            if not(True):
                d_1165_v3_: _dafny.Seq
                d_1165_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gljnndblh"))
                d_1166_v4_: int
                out31_: int
                out31_ = default__.m0(_dafny.Map({d_1165_v3_: self.f13}), globalState)
                d_1166_v4_ = out31_
                d_1167_v5_: _dafny.Array
                nw177_ = _dafny.Array(_dafny.Seq({}), 26)
                d_1167_v5_ = nw177_
                d_1167_v5_ = d_1167_v5_
                d_1168_v6_: C0
                nw178_ = C0()
                nw178_.ctor__((self).fm7(d_1163_cf4_, (0) - (d_1166_v4_), self.f13, globalState))
                d_1168_v6_ = nw178_
                d_1169_v7_: bool
                d_1170_v8_: int
                out32_: bool
                out33_: int
                out32_, out33_ = (self).m3(d_1163_cf4_, (d_1168_v6_).f15, (d_1168_v6_).f15, globalState)
                d_1169_v7_ = out32_
                d_1170_v8_ = out33_
                d_1165_v3_ = d_1165_v3_
            elif True:
                rhs184_ = self.f13
                rhs185_ = d_1163_cf4_
                lhs113_ = self
                lhs113_.f13 = rhs184_
                d_1163_cf4_ = rhs185_
                d_1171_v9_: _dafny.Map
                d_1171_v9_ = _dafny.Map({self.f13: (d_1164_v2_).set(p0, default__.abs(p0))})
                d_1171_v9_ = (d_1171_v9_).set(self.f13, (d_1164_v2_) - (d_1164_v2_))
                d_1163_cf4_ = p0
                index217_ = default__.safeIndex(371, ((self).f14).length(0))
                ((self).f14)[index217_] = p0
                index218_ = default__.safeIndex(371, ((self).f14).length(0))
                ((self).f14)[index218_] = d_1163_cf4_
                d_1172_v10_: _dafny.Map
                d_1172_v10_ = _dafny.Map({((self).f14)[default__.safeIndex(371, ((self).f14).length(0))]: self.f13})
                d_1163_cf4_ = len(d_1172_v10_)
            d_1173_v11_: _dafny.Seq
            d_1173_v11_ = _dafny.SeqWithoutIsStrInference([(self).f14])
            d_1174_v12_: C0
            nw179_ = C0()
            nw179_.ctor__((0) - (len(d_1173_v11_)))
            d_1174_v12_ = nw179_
            d_1175_v13_: _dafny.Seq
            d_1175_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqmf"))
            d_1176_v14_: str
            d_1176_v14_ = _dafny.CodePoint('w')
            d_1175_v13_ = ((d_1175_v13_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abddeo")))).set(default__.safeIndex((p0) * (len(d_1175_v13_)), len((d_1175_v13_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abddeo"))))), d_1176_v14_)
        elif source20_.is_DC4:
            d_1177___mcc_h1_ = source20_.cf5
            d_1178___mcc_h2_ = source20_.cf6
            d_1179___mcc_h3_ = source20_.cf7
            d_1180_cf7_ = d_1179___mcc_h3_
            d_1181_cf6_ = d_1178___mcc_h2_
            d_1182_cf5_ = d_1177___mcc_h1_
            d_1183_v15_: _dafny.MultiSet
            d_1183_v15_ = _dafny.MultiSet([self.f13, self.f13])
            d_1181_cf6_ = (d_1183_v15_).cardinality
            (self).f13 = self.f13
            d_1184_v16_: _dafny.Seq
            d_1184_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "buj"))
            d_1184_v16_ = d_1184_v16_
            d_1185_v17_: _dafny.Map
            d_1185_v17_ = _dafny.Map({d_1181_cf6_: (self).fm9((self).fm9(p0, globalState), globalState)})
            d_1186_v18_: _dafny.Map
            d_1186_v18_ = _dafny.Map({685: d_1185_v17_})
            d_1187_v19_: _dafny.Map
            d_1187_v19_ = _dafny.Map({self.f13: self.f13})
            d_1186_v18_ = (d_1186_v18_).set((0) - ((self).fm9(len(d_1187_v19_), globalState)), (d_1185_v17_).set(p0, (d_1183_v15_).cardinality))
        elif source20_.is_DC5:
            d_1188___mcc_h4_ = source20_.cf8
            d_1189___mcc_h5_ = source20_.cf9
            d_1190___mcc_h6_ = source20_.cf10
            d_1191_cf10_ = d_1190___mcc_h6_
            d_1192_cf9_ = d_1189___mcc_h5_
            d_1193_cf8_ = d_1188___mcc_h4_
            d_1194_v20_: _dafny.Array
            def lambda52_(d_1195_p0_):
                def lambda53_(d_1196_i0_):
                    return _dafny.MultiSet((D2_DC6(_dafny.SeqWithoutIsStrInference([d_1195_p0_]))).cf11)

                return lambda53_

            init29_ = lambda52_(p0)
            nw180_ = _dafny.Array(None, 25)
            for i0_29_ in range(nw180_.length(0)):
                nw180_[i0_29_] = init29_(i0_29_)
            d_1194_v20_ = nw180_
            d_1194_v20_ = d_1194_v20_
            d_1197_v21_: D0
            d_1197_v21_ = D0_DC0(self.f13)
            d_1197_v21_ = d_1197_v21_
            d_1198_v22_: _dafny.Map
            d_1198_v22_ = _dafny.Map({p0: d_1161_v1_})
            d_1199_v23_: _dafny.Map
            d_1199_v23_ = _dafny.Map({self.f13: _dafny.CodePoint('x')})
            d_1198_v22_ = (d_1198_v22_).set(len(d_1199_v23_), d_1161_v1_)
            d_1200_v24_: _dafny.Map
            d_1200_v24_ = _dafny.Map({d_1193_cf8_: default__.fm11(self.f13, p0, self.f13, globalState)})
            d_1201_v25_: str
            d_1201_v25_ = _dafny.CodePoint('b')
            d_1200_v24_ = (d_1200_v24_).set(d_1193_cf8_, d_1201_v25_)
        elif True:
            d_1202___mcc_h7_ = source20_.cf3
            d_1203_cf3_ = d_1202___mcc_h7_
            d_1204_v26_: int
            d_1204_v26_ = 283
            d_1204_v26_ = (0) - ((0) - (d_1204_v26_))
            d_1204_v26_ = p0
            d_1204_v26_ = len(_dafny.Map({d_1204_v26_: (d_1204_v26_ if not(self.f13) else d_1204_v26_)}))
            (self).f13 = self.f13
        d_1205_v27_: int
        d_1205_v27_ = -407
        d_1206_v28_: _dafny.Seq
        d_1206_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sffwtcir"))
        d_1205_v27_ = len((d_1206_v28_) + (d_1206_v28_))
        d_1207_v29_: _dafny.Array
        nw181_ = _dafny.Array(None, 18)
        d_1207_v29_ = nw181_
        d_1208_v30_: _dafny.Set
        d_1208_v30_ = _dafny.Set({self.f13})
        d_1209_v31_: _dafny.Set
        d_1209_v31_ = _dafny.Set({d_1208_v30_})
        d_1210_v32_: C0
        nw182_ = C0()
        nw182_.ctor__(len(d_1209_v31_))
        d_1210_v32_ = nw182_
        index219_ = default__.safeIndex(991, (d_1207_v29_).length(0))
        (d_1207_v29_)[index219_] = d_1210_v32_
        index220_ = default__.safeIndex(991, (d_1207_v29_).length(0))
        (d_1207_v29_)[index220_] = d_1210_v32_
        (self).f13 = True
        d_1211_v33_: C0
        nw183_ = C0()
        nw183_.ctor__(p0)
        d_1211_v33_ = nw183_
        d_1212_v34_: _dafny.Set
        d_1212_v34_ = _dafny.Set({549})
        d_1213_v35_: _dafny.Map
        d_1213_v35_ = _dafny.Map({(0) - (-689): d_1212_v34_})
        d_1214_v36_: _dafny.MultiSet
        d_1214_v36_ = _dafny.MultiSet([False, self.f13])
        d_1215_v37_: _dafny.Map
        d_1215_v37_ = _dafny.Map({d_1214_v36_: (d_1211_v33_).f15})
        d_1216_v38_: _dafny.Seq
        d_1216_v38_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1217_v39_: _dafny.Map
        d_1217_v39_ = _dafny.Map({(d_1216_v38_)[default__.safeIndex((self).fm9(d_1205_v27_, globalState), len(d_1216_v38_))]: (self).f14})
        rhs186_ = (len((((d_1213_v35_)[(d_1210_v32_).f15] if ((d_1210_v32_).f15) in (d_1213_v35_) else d_1212_v34_)) - (d_1212_v34_))) <= (len(d_1215_v37_))
        rhs187_ = d_1211_v33_
        rhs188_ = (p0) * (len(d_1217_v39_))
        lhs114_ = self
        lhs114_.f13 = rhs186_
        d_1210_v32_ = rhs187_
        d_1205_v27_ = rhs188_
        r0 = (self).f14
        return r0

    def m2(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        r2: bool = False
        r3: _dafny.Seq = _dafny.Seq({})
        d_1218_i0_: int
        d_1218_i0_ = 0
        with _dafny.label("11"):
            while self.f13:
                with _dafny.c_label("11"):
                    if (d_1218_i0_) >= (100):
                        raise _dafny.Break("11")
                    d_1218_i0_ = (d_1218_i0_) + (1)
                    d_1219_v0_: int
                    d_1219_v0_ = 21
                    d_1219_v0_ = default__.safeDivisionInt(default__.safeDivisionInt(p2, p2), (p0 if self.f13 else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1220_i1_ in range(default__.abs(545))]))))
                    d_1221_v1_: _dafny.Array
                    def lambda54_(d_1222_p1_):
                        def lambda55_(d_1223_i2_):
                            return (d_1222_p1_) + (d_1222_p1_)

                        return lambda55_

                    init30_ = lambda54_(p1)
                    nw184_ = _dafny.Array(None, 20)
                    for i0_30_ in range(nw184_.length(0)):
                        nw184_[i0_30_] = init30_(i0_30_)
                    d_1221_v1_ = nw184_
                    index221_ = default__.safeIndex(152, (d_1221_v1_).length(0))
                    (d_1221_v1_)[index221_] = p1
                    index222_ = default__.safeIndex(152, (d_1221_v1_).length(0))
                    (d_1221_v1_)[index222_] = p1
                    d_1224_v2_: _dafny.MultiSet
                    d_1224_v2_ = _dafny.MultiSet([p2])
                    d_1225_v3_: _dafny.Set
                    d_1225_v3_ = _dafny.Set({d_1224_v2_})
                    d_1226_v4_: str
                    d_1226_v4_ = _dafny.CodePoint('t')
                    d_1227_v5_: _dafny.Map
                    d_1227_v5_ = _dafny.Map({d_1225_v3_: (p1) + ((p1).set(default__.safeIndex(d_1219_v0_, len(p1)), d_1226_v4_))})
                    d_1227_v5_ = (d_1227_v5_).set(d_1225_v3_, ((d_1221_v1_)[default__.safeIndex(152, (d_1221_v1_).length(0))]).set(default__.safeIndex(p2, len((d_1221_v1_)[default__.safeIndex(152, (d_1221_v1_).length(0))])), d_1226_v4_))
                    d_1228_v6_: _dafny.Array
                    nw185_ = _dafny.Array(None, 12)
                    d_1228_v6_ = nw185_
                    d_1229_v7_: T0
                    nw186_ = C3()
                    nw186_.ctor__(self.f13, self.f13)
                    d_1229_v7_ = nw186_
                    index223_ = default__.safeIndex(838, (d_1228_v6_).length(0))
                    (d_1228_v6_)[index223_] = d_1229_v7_
                    index224_ = default__.safeIndex(838, (d_1228_v6_).length(0))
                    nw187_ = C3()
                    nw187_.ctor__(self.f13, not(self.f13))
                    (d_1228_v6_)[index224_] = nw187_
                    pass
            pass
        d_1230_v8_: C3
        nw188_ = C3()
        nw188_.ctor__((False if self.f13 else self.f13), self.f13)
        d_1230_v8_ = nw188_
        if self.f13:
            d_1231_v9_: _dafny.Seq
            d_1231_v9_ = _dafny.SeqWithoutIsStrInference([not(self.f13)])
            d_1232_v10_: _dafny.Map
            d_1232_v10_ = _dafny.Map({p2: len(d_1231_v9_)})
            d_1233_v11_: D0
            d_1233_v11_ = D0_DC1(self.f13, ((d_1232_v10_)[p2] if (p2) in (d_1232_v10_) else p0))
            source21_ = d_1233_v11_
            if source21_.is_DC1:
                d_1234___mcc_h0_ = source21_.cf1
                d_1235___mcc_h1_ = source21_.cf2
                d_1236_cf2_ = d_1235___mcc_h1_
                d_1237_cf1_ = d_1234___mcc_h0_
                (self).f13 = default__.fm0((d_1232_v10_) | (d_1232_v10_), (len(p1)) * (p0), globalState)
                d_1236_cf2_ = 207
                d_1236_cf2_ = p0
                d_1238_v12_: _dafny.Set
                d_1238_v12_ = _dafny.Set({p2, p0, p0})
                d_1239_v13_: _dafny.Map
                d_1239_v13_ = _dafny.Map({len(d_1231_v9_): (d_1230_v8_).f19})
                d_1240_v15_: _dafny.Array
                nw189_ = _dafny.Array(None, 5)
                nw189_[int(0)] = d_1238_v12_
                nw189_[int(1)] = d_1238_v12_
                nw189_[int(2)] = (d_1238_v12_).intersection(d_1238_v12_)
                nw189_[int(3)] = d_1238_v12_
                def iife98_():
                    coll42_ = _dafny.Set()
                    compr_42_: int
                    for compr_42_ in (d_1239_v13_).keys.Elements:
                        d_1241_v14_: int = compr_42_
                        if (d_1241_v14_) in (d_1239_v13_):
                            coll42_ = coll42_.union(_dafny.Set([(d_1241_v14_) * (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([False, False])})), len(_dafny.Map({True: len(_dafny.Map({False: False}))})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1242_i3_ in range(default__.abs(442))]))])))]))
                    return _dafny.Set(coll42_)
                nw189_[int(4)] = iife98_()
                
                d_1240_v15_ = nw189_
                d_1243_v16_: _dafny.Map
                d_1243_v16_ = _dafny.Map({p2: d_1230_v8_})
                index225_ = default__.safeIndex(473, (d_1240_v15_).length(0))
                (d_1240_v15_)[index225_] = _dafny.Set({len(d_1243_v16_)})
                d_1244_v17_: _dafny.Seq
                d_1244_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucptmuvvm"))
                d_1245_v18_: _dafny.Seq
                d_1245_v18_ = _dafny.SeqWithoutIsStrInference([p2])
                index226_ = default__.safeIndex(473, (d_1240_v15_).length(0))
                rhs189_ = ((d_1238_v12_) | (d_1238_v12_)) - ((_dafny.Set({d_1236_cf2_, d_1236_cf2_, len(d_1245_v18_)})).intersection(d_1238_v12_))
                rhs190_ = self.f13
                rhs191_ = (d_1230_v8_).f19
                rhs192_ = len(_dafny.Set({(p0) + (d_1236_cf2_), (0) - ((0) - (p2)), p0}))
                rhs193_ = p1
                lhs115_ = d_1240_v15_
                lhs116_ = default__.safeIndex(473, (d_1240_v15_).length(0))
                lhs115_[lhs116_] = rhs189_
                d_1237_cf1_ = rhs190_
                d_1237_cf1_ = rhs191_
                d_1236_cf2_ = rhs192_
                d_1244_v17_ = rhs193_
            elif True:
                d_1246___mcc_h2_ = source21_.cf0
                d_1247_cf0_ = d_1246___mcc_h2_
                d_1248_v19_: str
                d_1248_v19_ = _dafny.CodePoint('m')
                d_1249_v20_: _dafny.Map
                d_1249_v20_ = _dafny.Map({(d_1230_v8_).f19: (d_1230_v8_).f19})
                d_1250_v21_: _dafny.MultiSet
                d_1250_v21_ = _dafny.MultiSet([((d_1249_v20_)[d_1247_cf0_] if (d_1247_cf0_) in (d_1249_v20_) else (d_1230_v8_).f19), self.f13])
                d_1251_v22_: D2
                d_1251_v22_ = D2_DC7((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), d_1248_v19_, d_1248_v19_])).set(default__.safeIndex(len(p1), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), d_1248_v19_, d_1248_v19_]))), d_1248_v19_), (d_1250_v21_).cardinality, p2, d_1247_cf0_)
                d_1252_v23_: _dafny.Map
                d_1252_v23_ = _dafny.Map({not((d_1251_v22_).cf15): p2})
                d_1252_v23_ = (d_1252_v23_).set(not((d_1230_v8_).f19), p2)
                d_1253_v24_: _dafny.Map
                d_1253_v24_ = _dafny.Map({p2: d_1232_v10_})
                d_1254_v25_: C4
                nw190_ = C4()
                nw190_.ctor__((self).fm8(D1_DC2(((d_1253_v24_)[p2] if (p2) in (d_1253_v24_) else default__.fm18((d_1230_v8_).f19, (d_1230_v8_).f19, len(d_1252_v23_), (self).fm7(p0, p2, d_1247_cf0_, globalState), globalState))), p2, globalState), not((d_1230_v8_).f19))
                d_1254_v25_ = nw190_
                d_1255_v26_: D7
                d_1255_v26_ = D7_DC20(606, d_1248_v19_)
                d_1256_v27_: D7
                d_1256_v27_ = D7_DC21(d_1255_v26_)
                d_1257_v28_: D7
                d_1257_v28_ = D7_DC21(d_1255_v26_)
                d_1258_v29_: D7
                d_1258_v29_ = D7_DC21(d_1256_v27_)
                d_1258_v29_ = d_1258_v29_
                d_1249_v20_ = (d_1249_v20_) | ((d_1249_v20_) | (d_1249_v20_))
            d_1259_v30_: str
            d_1259_v30_ = _dafny.CodePoint('h')
            d_1260_v31_: _dafny.Map
            d_1260_v31_ = _dafny.Map({(D4_DC11((d_1230_v8_).f19)).cf21: self.f13})
            d_1261_v32_: D2
            d_1261_v32_ = D2_DC8(self.f13, False, self.f13)
            d_1259_v30_ = default__.fm11(((d_1260_v31_)[(d_1230_v8_).f19] if ((d_1230_v8_).f19) in (d_1260_v31_) else (d_1261_v32_).cf17), p2, default__.fm0(_dafny.Map({p0: -513}), p2, globalState), globalState)
            d_1262_v33_: _dafny.Array
            nw191_ = _dafny.Array(False, 29)
            d_1262_v33_ = nw191_
            index227_ = default__.safeIndex(852, (d_1262_v33_).length(0))
            (d_1262_v33_)[index227_] = (p0) <= (p2)
            d_1263_v34_: C4
            nw192_ = C4()
            nw192_.ctor__(self.f13, (d_1230_v8_).f19)
            d_1263_v34_ = nw192_
            d_1264_v35_: _dafny.Set
            d_1264_v35_ = _dafny.Set({(d_1230_v8_).f19})
            d_1265_v36_: _dafny.Map
            d_1265_v36_ = _dafny.Map({d_1263_v34_: d_1264_v35_})
            index228_ = default__.safeIndex(852, (d_1262_v33_).length(0))
            rhs194_ = (d_1230_v8_).f19
            rhs195_ = ((self).fm9(len(((d_1265_v36_)[d_1263_v34_] if (d_1263_v34_) in (d_1265_v36_) else _dafny.Set({self.f13}))), globalState)) >= (p0)
            lhs117_ = self
            lhs118_ = d_1262_v33_
            lhs119_ = default__.safeIndex(852, (d_1262_v33_).length(0))
            lhs117_.f13 = rhs194_
            lhs118_[lhs119_] = rhs195_
            d_1266_v37_: _dafny.Seq
            d_1266_v37_ = d_1231_v9_
            source22_ = d_1266_v37_
            d_1267___mcc_h3_ = source22_
            d_1268_cf34_ = d_1267___mcc_h3_
            index229_ = default__.safeIndex(852, (d_1262_v33_).length(0))
            (d_1262_v33_)[index229_] = (d_1230_v8_).f19
            index230_ = default__.safeIndex(980, ((self).f14).length(0))
            ((self).f14)[index230_] = p2
            index231_ = default__.safeIndex(980, ((self).f14).length(0))
            ((self).f14)[index231_] = p2
            d_1269_v38_: D5
            d_1269_v38_ = D5_DC14((d_1262_v33_)[default__.safeIndex(852, (d_1262_v33_).length(0))])
            d_1270_v39_: _dafny.Set
            d_1270_v39_ = _dafny.Set({D5_DC14((d_1263_v34_).f16), d_1269_v38_, d_1269_v38_, d_1269_v38_})
            d_1271_v40_: _dafny.Array
            nw193_ = _dafny.Array(None, 9)
            nw193_[int(0)] = d_1262_v33_
            nw193_[int(1)] = d_1262_v33_
            nw193_[int(2)] = d_1262_v33_
            nw193_[int(3)] = d_1262_v33_
            nw193_[int(4)] = d_1262_v33_
            nw193_[int(5)] = d_1262_v33_
            nw193_[int(6)] = d_1262_v33_
            nw193_[int(7)] = d_1262_v33_
            nw193_[int(8)] = d_1262_v33_
            d_1271_v40_ = nw193_
            d_1272_v41_: _dafny.Map
            d_1272_v41_ = _dafny.Map({d_1270_v39_: d_1271_v40_})
            d_1272_v41_ = (d_1272_v41_).set(d_1270_v39_, d_1271_v40_)
            d_1273_v42_: _dafny.Array
            def lambda56_(d_1274_v34_):
                def lambda57_(d_1275_i4_):
                    return D9_DC23(_dafny.MultiSet([(d_1274_v34_).f16]))

                return lambda57_

            init31_ = lambda56_(d_1263_v34_)
            nw194_ = _dafny.Array(None, 25)
            for i0_31_ in range(nw194_.length(0)):
                nw194_[i0_31_] = init31_(i0_31_)
            d_1273_v42_ = nw194_
            d_1276_v43_: D1
            d_1276_v43_ = D1_DC2(_dafny.Map({p0: (self).fm7(((self).f14)[default__.safeIndex(980, ((self).f14).length(0))], (_dafny.MultiSet([(0) - (p0)])).cardinality, (d_1263_v34_).f16, globalState)}))
            d_1277_v44_: _dafny.MultiSet
            d_1277_v44_ = _dafny.MultiSet([(self).fm8(d_1276_v43_, p0, globalState), (d_1263_v34_).f16])
            d_1278_v45_: D9
            d_1278_v45_ = D9_DC23(d_1277_v44_)
            index232_ = default__.safeIndex(387, (d_1273_v42_).length(0))
            (d_1273_v42_)[index232_] = d_1278_v45_
            index233_ = default__.safeIndex(852, (d_1262_v33_).length(0))
            index234_ = default__.safeIndex(980, ((self).f14).length(0))
            index235_ = default__.safeIndex(387, (d_1273_v42_).length(0))
            rhs196_ = not(False)
            rhs197_ = ((d_1277_v44_)[not(True)] if (not(True)) in (d_1277_v44_) else p2)
            rhs198_ = d_1278_v45_
            rhs199_ = d_1259_v30_
            lhs120_ = d_1262_v33_
            lhs121_ = default__.safeIndex(852, (d_1262_v33_).length(0))
            lhs122_ = (self).f14
            lhs123_ = default__.safeIndex(980, ((self).f14).length(0))
            lhs124_ = d_1273_v42_
            lhs125_ = default__.safeIndex(387, (d_1273_v42_).length(0))
            lhs120_[lhs121_] = rhs196_
            lhs122_[lhs123_] = rhs197_
            lhs124_[lhs125_] = rhs198_
            d_1259_v30_ = rhs199_
            d_1231_v9_ = d_1231_v9_
        elif True:
            d_1279_v46_: _dafny.Seq
            d_1279_v46_ = _dafny.SeqWithoutIsStrInference([p0, 627, p0, len(p1)])
            d_1280_v47_: C3
            nw195_ = C3()
            nw195_.ctor__((True if (d_1230_v8_).f19 else True), (d_1279_v46_) != (d_1279_v46_))
            d_1280_v47_ = nw195_
            d_1281_v48_: int
            d_1281_v48_ = 58
            d_1282_v49_: str
            d_1282_v49_ = _dafny.CodePoint('r')
            d_1281_v48_ = default__.fm13((d_1230_v8_).f19, self.f13, d_1282_v49_, globalState)
            d_1281_v48_ = p0
            d_1283_v50_: C3
            nw196_ = C3()
            nw196_.ctor__((self.f13) not in (_dafny.Set({False})), (p0) != ((0) - (p0)))
            d_1283_v50_ = nw196_
            d_1284_v51_: _dafny.Array
            nw197_ = _dafny.Array(_dafny.Map({}), 14)
            d_1284_v51_ = nw197_
            d_1285_v52_: C0
            nw198_ = C0()
            nw198_.ctor__(p2)
            d_1285_v52_ = nw198_
            d_1286_v53_: _dafny.Map
            d_1286_v53_ = _dafny.Map({d_1285_v52_: (d_1280_v47_).f19})
            index236_ = default__.safeIndex(611, (d_1284_v51_).length(0))
            (d_1284_v51_)[index236_] = d_1286_v53_
            index237_ = default__.safeIndex(611, (d_1284_v51_).length(0))
            (d_1284_v51_)[index237_] = ((d_1286_v53_).set(d_1285_v52_, (d_1230_v8_).f19)) | (d_1286_v53_)
        d_1287_v54_: _dafny.Seq
        d_1287_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvtqh"))
        d_1287_v54_ = p1
        index238_ = default__.safeIndex(400, ((self).f14).length(0))
        ((self).f14)[index238_] = p2
        d_1288_v55_: _dafny.MultiSet
        d_1288_v55_ = _dafny.MultiSet([p0, p0])
        d_1289_v56_: _dafny.Seq
        d_1289_v56_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, (d_1288_v55_).cardinality, p0])])
        index239_ = default__.safeIndex(400, ((self).f14).length(0))
        ((self).f14)[index239_] = len(d_1289_v56_)
        d_1290_v57_: _dafny.Map
        d_1290_v57_ = _dafny.Map({p0: p2})
        d_1291_v58_: _dafny.Map
        d_1291_v58_ = _dafny.Map({((self).f14)[default__.safeIndex(400, ((self).f14).length(0))]: default__.fm0(d_1290_v57_, (_dafny.MultiSet([self.f13])).cardinality, globalState)})
        d_1292_v59_: _dafny.Map
        d_1292_v59_ = _dafny.Map({(d_1291_v58_) | (d_1291_v58_): d_1287_v54_})
        d_1292_v59_ = (d_1292_v59_).set(d_1291_v58_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbsuyxpfb")))
        d_1293_v60_: _dafny.Array
        def lambda58_(d_1294_i5_):
            return self.f13

        init32_ = lambda58_
        nw199_ = _dafny.Array(None, 20)
        for i0_32_ in range(nw199_.length(0)):
            nw199_[i0_32_] = init32_(i0_32_)
        d_1293_v60_ = nw199_
        d_1295_v61_: _dafny.Seq
        d_1295_v61_ = _dafny.SeqWithoutIsStrInference([d_1293_v60_])
        d_1296_v62_: _dafny.Map
        d_1296_v62_ = _dafny.Map({73: (d_1295_v61_)[default__.safeIndex(p2, len(d_1295_v61_))]})
        d_1297_v63_: _dafny.Seq
        d_1297_v63_ = _dafny.SeqWithoutIsStrInference([p2])
        d_1298_v64_: D1
        d_1298_v64_ = D1_DC5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvhsai")), d_1293_v60_, p2)
        d_1299_v65_: _dafny.Seq
        d_1299_v65_ = _dafny.SeqWithoutIsStrInference([(d_1296_v62_) | (_dafny.Map({len(d_1297_v63_): (D1_DC5((d_1298_v64_).cf8, d_1293_v60_, ((self).f14)[default__.safeIndex(400, ((self).f14).length(0))])).cf9}))])
        r0 = (d_1299_v65_)[default__.safeIndex(((self).f14)[default__.safeIndex(400, ((self).f14).length(0))], len(d_1299_v65_))]
        d_1300_v66_: _dafny.Seq
        d_1300_v66_ = _dafny.SeqWithoutIsStrInference([(d_1230_v8_).f19])
        r1 = (p0) != ((len(d_1300_v66_)) - ((self).fm9(len(d_1297_v63_), globalState)))
        r2 = ((d_1230_v8_).f19 if (d_1230_v8_).f19 else self.f13)
        r3 = d_1300_v66_
        return r0, r1, r2, r3

    def m3(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: int = int(0)
        d_1301_v0_: str
        d_1301_v0_ = _dafny.CodePoint('m')
        d_1302_v1_: D7
        d_1302_v1_ = D7_DC20(p1, d_1301_v0_)
        r1 = (0) - (default__.safeDivisionInt(p1, (d_1302_v1_).cf31))
        if not(True):
            r0 = self.f13
            d_1303_v2_: _dafny.MultiSet
            d_1303_v2_ = _dafny.MultiSet([508])
            r1 = (d_1303_v2_).cardinality
            d_1304_v3_: _dafny.Array
            nw200_ = _dafny.Array(False, 13)
            d_1304_v3_ = nw200_
            d_1305_v4_: _dafny.Seq
            d_1305_v4_ = _dafny.SeqWithoutIsStrInference([self.f13, self.f13])
            d_1306_v5_: _dafny.Set
            d_1306_v5_ = _dafny.Set({len(d_1305_v4_), len(d_1305_v4_)})
            index240_ = default__.safeIndex(42, (d_1304_v3_).length(0))
            (d_1304_v3_)[index240_] = (d_1306_v5_).isdisjoint(_dafny.Set({(self).fm9(165, globalState), p1}))
            d_1307_v6_: _dafny.Array
            nw201_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
            d_1307_v6_ = nw201_
            d_1308_v7_: _dafny.Seq
            d_1308_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            index241_ = default__.safeIndex(94, (d_1307_v6_).length(0))
            (d_1307_v6_)[index241_] = (d_1308_v7_) + (d_1308_v7_)
            d_1309_v8_: D1
            d_1309_v8_ = D1_DC5(d_1308_v7_, d_1304_v3_, p0)
            index242_ = default__.safeIndex(42, (d_1304_v3_).length(0))
            index243_ = default__.safeIndex(94, (d_1307_v6_).length(0))
            rhs200_ = (d_1301_v0_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywfkh")))
            rhs201_ = (d_1309_v8_).cf8
            lhs126_ = d_1304_v3_
            lhs127_ = default__.safeIndex(42, (d_1304_v3_).length(0))
            lhs128_ = d_1307_v6_
            lhs129_ = default__.safeIndex(94, (d_1307_v6_).length(0))
            lhs126_[lhs127_] = rhs200_
            lhs128_[lhs129_] = rhs201_
            d_1308_v7_ = d_1308_v7_
            d_1310_v9_: _dafny.Array
            nw202_ = _dafny.Array(None, 3)
            nw202_[int(0)] = _dafny.CodePoint('c')
            nw202_[int(1)] = d_1301_v0_
            nw202_[int(2)] = d_1301_v0_
            d_1310_v9_ = nw202_
            index244_ = default__.safeIndex(618, (d_1310_v9_).length(0))
            (d_1310_v9_)[index244_] = d_1301_v0_
            index245_ = default__.safeIndex(618, (d_1310_v9_).length(0))
            (d_1310_v9_)[index245_] = _dafny.CodePoint('y')
        elif True:
            r1 = (-939) * (len(_dafny.SeqWithoutIsStrInference([self.f13, self.f13, True])))
            d_1311_v10_: _dafny.Seq
            d_1311_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "garuvy"))
            d_1312_v11_: _dafny.Map
            d_1312_v11_ = _dafny.Map({d_1311_v10_: not (self.f13) or ((self).fm6(globalState))})
            d_1312_v11_ = (d_1312_v11_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fg")), not(self.f13))
            d_1313_v12_: _dafny.Map
            d_1313_v12_ = _dafny.Map({p0: True})
            d_1314_v13_: _dafny.MultiSet
            d_1314_v13_ = _dafny.MultiSet([self.f13])
            d_1315_v14_: _dafny.Map
            d_1315_v14_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([self.f13])): d_1314_v13_})
            d_1313_v12_ = (d_1313_v12_).set(-5, (p0) not in (d_1315_v14_))
            d_1316_v15_: D0
            d_1316_v15_ = D0_DC0(self.f13)
            d_1316_v15_ = D0_DC0(self.f13)
            r1 = default__.safeModuloInt(p0, 898)
        d_1317_v16_: _dafny.Array
        nw203_ = _dafny.Array(int(0), 12)
        d_1317_v16_ = nw203_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1317_v16_).length(0)):
            d_1318_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_1318_i0_)) and ((d_1318_i0_) < ((d_1317_v16_).length(0)))):
                (d_1317_v16_)[(d_1318_i0_)] = (d_1318_i0_) + (p2)
        d_1319_v17_: _dafny.Seq
        d_1319_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isf"))
        hi11_ = len(d_1319_v17_)
        for d_1320_i1_ in range(p0, hi11_):
            d_1321_v18_: _dafny.Array
            def lambda59_(d_1322_i2_):
                return False

            init33_ = lambda59_
            nw204_ = _dafny.Array(None, 11)
            for i0_33_ in range(nw204_.length(0)):
                nw204_[i0_33_] = init33_(i0_33_)
            d_1321_v18_ = nw204_
            index246_ = default__.safeIndex(600, (d_1321_v18_).length(0))
            (d_1321_v18_)[index246_] = self.f13
            d_1323_v19_: _dafny.Seq
            d_1323_v19_ = _dafny.SeqWithoutIsStrInference([(p0) != (p1)])
            d_1324_v20_: _dafny.Map
            d_1324_v20_ = _dafny.Map({True: p1})
            index247_ = default__.safeIndex(600, (d_1321_v18_).length(0))
            (d_1321_v18_)[index247_] = (d_1323_v19_)[default__.safeIndex(len(d_1324_v20_), len(d_1323_v19_))]
            d_1325_v21_: _dafny.Seq
            d_1325_v21_ = _dafny.SeqWithoutIsStrInference([p2])
            index248_ = default__.safeIndex(560, ((self).f14).length(0))
            ((self).f14)[index248_] = ((d_1325_v21_)[default__.safeIndex(len(d_1319_v17_), len(d_1325_v21_))]) * (p1)
            index249_ = default__.safeIndex(560, ((self).f14).length(0))
            ((self).f14)[index249_] = (0) - (p2)
            (self).f13 = (d_1321_v18_)[default__.safeIndex(600, (d_1321_v18_).length(0))]
            d_1326_v23_: C4
            nw205_ = C4()
            nw205_.ctor__(True, self.f13)
            d_1326_v23_ = nw205_
            d_1327_v24_: _dafny.MultiSet
            d_1327_v24_ = _dafny.MultiSet([(d_1326_v23_).f16])
            d_1328_v25_: _dafny.Map
            d_1328_v25_ = _dafny.Map({d_1326_v23_: d_1327_v24_})
            d_1329_v26_: _dafny.Map
            d_1329_v26_ = _dafny.Map({len(_dafny.Map({self.f13: p2})): (0) - (len(d_1328_v25_))})
            d_1330_v27_: _dafny.Set
            def iife99_():
                coll43_ = _dafny.Map()
                compr_43_: int
                for compr_43_ in _dafny.IntegerRange(983, 402):
                    d_1331_v22_: int = compr_43_
                    if ((983) <= (d_1331_v22_)) and ((d_1331_v22_) < (402)):
                        coll43_[(d_1331_v22_) - (p0)] = p0
                return _dafny.Map(coll43_)
            d_1330_v27_ = _dafny.Set({(iife99_()
             if (d_1321_v18_)[default__.safeIndex(600, (d_1321_v18_).length(0))] else (d_1329_v26_).set(p0, 361)), _dafny.Map({len(d_1319_v17_): 726}), (d_1329_v26_).set(p0, (0) - (len(_dafny.Set({default__.fm13((d_1321_v18_)[default__.safeIndex(600, (d_1321_v18_).length(0))], (d_1321_v18_)[default__.safeIndex(600, (d_1321_v18_).length(0))], d_1301_v0_, globalState), p0, p2, len(default__.fm24(len(d_1319_v17_), (d_1321_v18_)[default__.safeIndex(600, (d_1321_v18_).length(0))], globalState))})))), _dafny.Map({p2: p2})})
            d_1332_v28_: D0
            d_1332_v28_ = D0_DC1((d_1326_v23_).f16, 356)
            index250_ = default__.safeIndex(600, (d_1321_v18_).length(0))
            rhs202_ = ((self).f14)[default__.safeIndex(560, ((self).f14).length(0))]
            rhs203_ = p0
            rhs204_ = (d_1326_v23_).fm4(((0) - (p1)) < (306), d_1332_v28_, d_1320_i1_, d_1329_v26_, globalState)
            rhs205_ = default__.fm46(p0, globalState)
            lhs130_ = d_1321_v18_
            lhs131_ = default__.safeIndex(600, (d_1321_v18_).length(0))
            r1 = rhs202_
            r1 = rhs203_
            lhs130_[lhs131_] = rhs204_
            d_1330_v27_ = rhs205_
        d_1333_i3_: int
        d_1333_i3_ = 0
        with _dafny.label("12"):
            while self.f13:
                with _dafny.c_label("12"):
                    if (d_1333_i3_) >= (100):
                        raise _dafny.Break("12")
                    d_1333_i3_ = (d_1333_i3_) + (1)
                    r1 = (0) - (default__.safeDivisionInt(p2, (-158 if True else p0)))
                    d_1334_v29_: _dafny.Map
                    d_1334_v29_ = _dafny.Map({p0: (True) == (True)})
                    d_1335_v30_: _dafny.Map
                    d_1335_v30_ = _dafny.Map({self.f13: self.f13})
                    d_1336_v31_: _dafny.Map
                    d_1336_v31_ = _dafny.Map({(0) - (len(d_1335_v30_)): p2})
                    d_1334_v29_ = (d_1334_v29_).set((p1 if default__.fm0(d_1336_v31_, p0, globalState) else p0), self.f13)
                    d_1337_v32_: D2
                    d_1337_v32_ = D2_DC7(d_1319_v17_, p1, p0, self.f13)
                    d_1338_v33_: _dafny.Seq
                    d_1338_v33_ = _dafny.SeqWithoutIsStrInference([self.f13, self.f13, self.f13, self.f13, (self).fm4(self.f13, D0_DC1(self.f13, -46), 675, d_1336_v31_, globalState)])
                    d_1339_v34_: _dafny.Seq
                    d_1339_v34_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k')])])
                    pat_let_tv24_ = p1
                    d_1340_v35_: _dafny.Array
                    nw206_ = _dafny.Array(None, 19)
                    nw206_[int(0)] = d_1337_v32_
                    nw206_[int(1)] = d_1337_v32_
                    nw206_[int(2)] = d_1337_v32_
                    nw206_[int(3)] = d_1337_v32_
                    nw206_[int(4)] = d_1337_v32_
                    def iife100_(_pat_let28_0):
                        def iife101_(d_1341_dt__update__tmp_h0_):
                            def iife102_(_pat_let29_0):
                                def iife103_(d_1342_dt__update_hcf13_h0_):
                                    return D2_DC7((d_1341_dt__update__tmp_h0_).cf12, d_1342_dt__update_hcf13_h0_, (d_1341_dt__update__tmp_h0_).cf14, (d_1341_dt__update__tmp_h0_).cf15)
                                return iife103_(_pat_let29_0)
                            return iife102_(pat_let_tv24_)
                        return iife101_(_pat_let28_0)
                    nw206_[int(5)] = iife100_(d_1337_v32_)
                    nw206_[int(6)] = D2_DC7((d_1319_v17_).set(default__.safeIndex(p2, len(d_1319_v17_)), d_1301_v0_), p1, (0) - (p0), self.f13)
                    nw206_[int(7)] = d_1337_v32_
                    nw206_[int(8)] = d_1337_v32_
                    nw206_[int(9)] = d_1337_v32_
                    nw206_[int(10)] = D2_DC7(_dafny.SeqWithoutIsStrInference([d_1301_v0_]), len(d_1338_v33_), p0, self.f13)
                    nw206_[int(11)] = d_1337_v32_
                    nw206_[int(12)] = default__.fm23(p1, self.f13, globalState)
                    nw206_[int(13)] = d_1337_v32_
                    nw206_[int(14)] = d_1337_v32_
                    nw206_[int(15)] = default__.fm23(p1, self.f13, globalState)
                    nw206_[int(16)] = D2_DC7((d_1339_v34_)[default__.safeIndex(p2, len(d_1339_v34_))], p0, p0, self.f13)
                    nw206_[int(17)] = D2_DC7(_dafny.SeqWithoutIsStrInference([d_1301_v0_, d_1301_v0_]), p2, p1, self.f13)
                    nw206_[int(18)] = d_1337_v32_
                    d_1340_v35_ = nw206_
                    index251_ = default__.safeIndex(923, (d_1340_v35_).length(0))
                    (d_1340_v35_)[index251_] = d_1337_v32_
                    index252_ = default__.safeIndex(923, (d_1340_v35_).length(0))
                    (d_1340_v35_)[index252_] = d_1337_v32_
                    (self).f13 = self.f13
                    pass
            pass
        if self.f13:
            rhs206_ = -866
            rhs207_ = self.f13
            rhs208_ = (p2) - ((p2) - (p2))
            rhs209_ = 107
            r1 = rhs206_
            r0 = rhs207_
            r1 = rhs208_
            r1 = rhs209_
            r1 = p2
            d_1343_v36_: D9
            d_1343_v36_ = D9_DC24(self.f13, (d_1301_v0_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxkq"))), d_1301_v0_)
            source23_ = d_1343_v36_
            if source23_.is_DC24:
                d_1344___mcc_h0_ = source23_.cf36
                d_1345___mcc_h1_ = source23_.cf37
                d_1346___mcc_h2_ = source23_.cf38
                d_1347_cf38_ = d_1346___mcc_h2_
                d_1348_cf37_ = d_1345___mcc_h1_
                d_1349_cf36_ = d_1344___mcc_h0_
                (self).f13 = d_1349_cf36_
                d_1350_v37_: D7
                d_1350_v37_ = D7_DC20(p1, d_1301_v0_)
                d_1351_v38_: D7
                d_1351_v38_ = D7_DC21(d_1350_v37_)
                d_1352_v39_: D7
                d_1352_v39_ = D7_DC21(d_1350_v37_)
                d_1353_v40_: D7
                d_1353_v40_ = D7_DC21(d_1352_v39_)
                d_1354_v41_: _dafny.Seq
                d_1354_v41_ = _dafny.SeqWithoutIsStrInference([d_1353_v40_, d_1353_v40_])
                d_1355_v42_: C2
                nw207_ = C2()
                nw207_.ctor__((d_1354_v41_) != (d_1354_v41_))
                d_1355_v42_ = nw207_
                d_1355_v42_ = d_1355_v42_
                d_1356_v43_: _dafny.Seq
                d_1356_v43_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_1357_v44_: _dafny.MultiSet
                d_1357_v44_ = _dafny.MultiSet([d_1349_cf36_])
                r0 = (p2) > ((d_1356_v43_)[default__.safeIndex((d_1357_v44_).cardinality, len(d_1356_v43_))])
                d_1358_v45_: _dafny.Array
                nw208_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_1358_v45_ = nw208_
                d_1359_v46_: D4
                d_1359_v46_ = D4_DC10(d_1317_v16_)
                index253_ = default__.safeIndex(738, (d_1358_v45_).length(0))
                (d_1358_v45_)[index253_] = (d_1359_v46_).cf20
                index254_ = default__.safeIndex(738, (d_1358_v45_).length(0))
                (d_1358_v45_)[index254_] = d_1317_v16_
            elif source23_.is_DC25:
                d_1360___mcc_h3_ = source23_.cf39
                d_1361___mcc_h4_ = source23_.cf40
                d_1362___mcc_h5_ = source23_.cf41
                d_1363_cf41_ = d_1362___mcc_h5_
                d_1364_cf40_ = d_1361___mcc_h4_
                d_1365_cf39_ = d_1360___mcc_h3_
                d_1366_v47_: _dafny.Map
                d_1366_v47_ = _dafny.Map({(self).fm7((0) - (p2), p0, d_1365_cf39_, globalState): p0})
                d_1367_v48_: C2
                nw209_ = C2()
                nw209_.ctor__(default__.fm0(d_1366_v47_, p0, globalState))
                d_1367_v48_ = nw209_
                d_1368_v49_: _dafny.Map
                d_1368_v49_ = _dafny.Map({d_1364_cf40_: d_1367_v48_})
                rhs210_ = (d_1363_cf41_) or (d_1365_cf39_)
                rhs211_ = (len(d_1368_v49_)) >= (p1)
                r0 = rhs210_
                d_1365_cf39_ = rhs211_
                d_1369_v50_: _dafny.Map
                d_1369_v50_ = _dafny.Map({self.f13: len(_dafny.Set({d_1365_cf39_, d_1363_cf41_}))})
                d_1370_v51_: _dafny.Seq
                d_1370_v51_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_1369_v50_)), 690, (0) - (d_1364_cf40_)])
                d_1371_v52_: _dafny.Map
                d_1371_v52_ = _dafny.Map({d_1365_cf39_: d_1365_cf39_})
                d_1372_v53_: _dafny.Map
                d_1372_v53_ = _dafny.Map({(d_1370_v51_)[default__.safeIndex(475, len(d_1370_v51_))]: d_1371_v52_})
                d_1373_v54_: _dafny.Map
                d_1373_v54_ = _dafny.Map({d_1301_v0_: d_1364_cf40_})
                r1 = (len((d_1372_v53_).set(len(d_1373_v54_), d_1371_v52_))) + (p2)
                r1 = default__.safeDivisionInt(d_1364_cf40_, d_1364_cf40_)
                index255_ = default__.safeIndex(301, (d_1317_v16_).length(0))
                (d_1317_v16_)[index255_] = (p0 if d_1365_cf39_ else 664)
                d_1374_v55_: D0
                d_1374_v55_ = D0_DC1(d_1365_cf39_, 209)
                d_1375_v56_: _dafny.Set
                d_1375_v56_ = _dafny.Set({d_1301_v0_, d_1301_v0_})
                index256_ = default__.safeIndex(301, (d_1317_v16_).length(0))
                def iife104_():
                    coll44_ = _dafny.Set()
                    compr_44_: str
                    for compr_44_ in (d_1375_v56_).Elements:
                        d_1376_v57_: str = compr_44_
                        if (d_1376_v57_) in (d_1375_v56_):
                            coll44_ = coll44_.union(_dafny.Set([d_1376_v57_]))
                    return _dafny.Set(coll44_)
                rhs212_ = default__.safeDivisionInt(len((default__.fm47(d_1374_v55_, p0, globalState)) | (iife104_()
                )), p2)
                rhs213_ = len(d_1319_v17_)
                rhs214_ = d_1369_v50_
                lhs132_ = d_1317_v16_
                lhs133_ = default__.safeIndex(301, (d_1317_v16_).length(0))
                r1 = rhs212_
                lhs132_[lhs133_] = rhs213_
                d_1369_v50_ = rhs214_
            elif True:
                d_1377___mcc_h6_ = source23_.cf35
                d_1378_cf35_ = d_1377___mcc_h6_
                d_1379_v58_: _dafny.Array
                nw210_ = _dafny.Array(False, 12)
                d_1379_v58_ = nw210_
                index257_ = default__.safeIndex(937, (d_1379_v58_).length(0))
                (d_1379_v58_)[index257_] = (self.f13 if self.f13 else self.f13)
                d_1380_v59_: C1
                nw211_ = C1()
                nw211_.ctor__(self.f13, d_1319_v17_, (self).f14)
                d_1380_v59_ = nw211_
                d_1381_v60_: _dafny.Seq
                d_1381_v60_ = _dafny.SeqWithoutIsStrInference([self.f13])
                d_1382_v61_: _dafny.Map
                d_1382_v61_ = _dafny.Map({d_1380_v59_: (_dafny.MultiSet(d_1381_v60_)) | (d_1378_cf35_)})
                d_1383_v62_: _dafny.MultiSet
                d_1383_v62_ = _dafny.MultiSet([d_1380_v59_.f18, d_1380_v59_.f18])
                d_1384_v63_: _dafny.Map
                d_1384_v63_ = _dafny.Map({p2: self.f13})
                d_1385_v64_: _dafny.Seq
                d_1385_v64_ = _dafny.SeqWithoutIsStrInference([d_1319_v17_, d_1380_v59_.f18, d_1319_v17_])
                index258_ = default__.safeIndex(937, (d_1379_v58_).length(0))
                rhs215_ = (((d_1383_v62_).set(_dafny.SeqWithoutIsStrInference([d_1301_v0_ for d_1386_i4_ in range(default__.abs(-486))]), default__.abs(len(d_1384_v63_)))).set((d_1385_v64_)[default__.safeIndex(p1, len(d_1385_v64_))], default__.abs(((d_1378_cf35_)[self.f13] if (self.f13) in (d_1378_cf35_) else p0)))).isdisjoint(d_1383_v62_)
                rhs216_ = self.f13
                rhs217_ = ((self.f13 if self.f13 else False) if (d_1380_v59_).f17 else not(self.f13))
                rhs218_ = (d_1382_v61_ if not ((d_1380_v59_).f17) or ((d_1380_v59_).f17) else (d_1382_v61_) | (d_1382_v61_))
                lhs134_ = self
                lhs135_ = self
                lhs136_ = d_1379_v58_
                lhs137_ = default__.safeIndex(937, (d_1379_v58_).length(0))
                lhs134_.f13 = rhs215_
                lhs135_.f13 = rhs216_
                lhs136_[lhs137_] = rhs217_
                d_1382_v61_ = rhs218_
                index259_ = default__.safeIndex(624, (d_1317_v16_).length(0))
                (d_1317_v16_)[index259_] = p2
                index260_ = default__.safeIndex(937, (d_1379_v58_).length(0))
                index261_ = default__.safeIndex(624, (d_1317_v16_).length(0))
                index262_ = default__.safeIndex(937, (d_1379_v58_).length(0))
                rhs219_ = (d_1379_v58_)[default__.safeIndex(937, (d_1379_v58_).length(0))]
                rhs220_ = p0
                rhs221_ = ((0) - (p2)) - (p0)
                rhs222_ = (d_1380_v59_).f17
                rhs223_ = (d_1378_cf35_) - (d_1378_cf35_)
                lhs138_ = d_1379_v58_
                lhs139_ = default__.safeIndex(937, (d_1379_v58_).length(0))
                lhs140_ = d_1317_v16_
                lhs141_ = default__.safeIndex(624, (d_1317_v16_).length(0))
                lhs142_ = d_1379_v58_
                lhs143_ = default__.safeIndex(937, (d_1379_v58_).length(0))
                lhs138_[lhs139_] = rhs219_
                lhs140_[lhs141_] = rhs220_
                r1 = rhs221_
                lhs142_[lhs143_] = rhs222_
                d_1378_cf35_ = rhs223_
                d_1387_v65_: _dafny.Set
                d_1387_v65_ = _dafny.Set({d_1319_v17_})
                d_1388_v66_: _dafny.Array
                nw212_ = _dafny.Array(int(0), 23)
                d_1388_v66_ = nw212_
                d_1389_v67_: _dafny.Map
                d_1389_v67_ = _dafny.Map({d_1380_v59_.f18: (d_1380_v59_).f17})
                index263_ = default__.safeIndex(624, (d_1317_v16_).length(0))
                def iife105_():
                    coll45_ = _dafny.Set()
                    compr_45_: _dafny.Seq
                    for compr_45_ in (d_1389_v67_).keys.Elements:
                        d_1390_v68_: _dafny.Seq = compr_45_
                        if (d_1390_v68_) in (d_1389_v67_):
                            coll45_ = coll45_.union(_dafny.Set([d_1390_v68_]))
                    return _dafny.Set(coll45_)
                rhs224_ = (default__.safeDivisionInt(p1, len(d_1381_v60_))) >= (p0)
                rhs225_ = iife105_()

                rhs226_ = d_1301_v0_
                rhs227_ = (self).f14
                rhs228_ = (d_1380_v59_).fm7(517, ((d_1317_v16_)[default__.safeIndex(624, (d_1317_v16_).length(0))]) * (p2), (d_1380_v59_).f17, globalState)
                lhs144_ = d_1317_v16_
                lhs145_ = default__.safeIndex(624, (d_1317_v16_).length(0))
                r0 = rhs224_
                d_1387_v65_ = rhs225_
                d_1301_v0_ = rhs226_
                d_1388_v66_ = rhs227_
                lhs144_[lhs145_] = rhs228_
                r0 = not(True)
            d_1391_v69_: C0
            nw213_ = C0()
            nw213_.ctor__(p0)
            d_1391_v69_ = nw213_
            d_1319_v17_ = d_1319_v17_
        elif True:
            d_1392_v70_: _dafny.Set
            d_1392_v70_ = _dafny.Set({self.f13})
            d_1393_v71_: _dafny.Map
            d_1393_v71_ = _dafny.Map({self.f13: len(d_1392_v70_)})
            d_1394_v72_: _dafny.Map
            d_1394_v72_ = _dafny.Map({d_1393_v71_: not(self.f13)})
            d_1395_v73_: _dafny.Map
            d_1395_v73_ = _dafny.Map({-410: len(d_1394_v72_)})
            d_1396_v74_: _dafny.Seq
            d_1396_v74_ = _dafny.SeqWithoutIsStrInference([-185])
            d_1397_v75_: _dafny.Set
            d_1397_v75_ = _dafny.Set({p0, len(d_1396_v74_)})
            d_1398_v76_: _dafny.Map
            d_1398_v76_ = _dafny.Map({p0: ((d_1395_v73_)[p0] if (p0) in (d_1395_v73_) else len(d_1397_v75_))})
            d_1399_v77_: _dafny.Set
            d_1399_v77_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([self.f13])})
            d_1398_v76_ = (d_1398_v76_).set(p0, default__.safeModuloInt(len(d_1399_v77_), p0))
            r1 = default__.safeModuloInt(p2, p2)
            d_1400_v78_: D9
            d_1400_v78_ = D9_DC25(self.f13, (0) - (p0), self.f13)
            d_1401_v79_: _dafny.Map
            d_1401_v79_ = _dafny.Map({self.f13: self.f13})
            (self).f13 = not(not ((default__.fm20(self.f13, p0, (d_1400_v78_).cf41, globalState)) == (d_1401_v79_)) or (False))
            (self).f13 = not(self.f13)
            d_1402_v80_: _dafny.Array
            def lambda60_(d_1403_v0_):
                def lambda61_(d_1404_i5_):
                    return d_1403_v0_

                return lambda61_

            init34_ = lambda60_(d_1301_v0_)
            nw214_ = _dafny.Array(None, 11)
            for i0_34_ in range(nw214_.length(0)):
                nw214_[i0_34_] = init34_(i0_34_)
            d_1402_v80_ = nw214_
            index264_ = default__.safeIndex(802, (d_1402_v80_).length(0))
            (d_1402_v80_)[index264_] = d_1301_v0_
            index265_ = default__.safeIndex(802, (d_1402_v80_).length(0))
            rhs229_ = 92
            rhs230_ = d_1301_v0_
            lhs146_ = d_1402_v80_
            lhs147_ = default__.safeIndex(802, (d_1402_v80_).length(0))
            r1 = rhs229_
            lhs146_[lhs147_] = rhs230_
        d_1405_v81_: _dafny.Array
        nw215_ = _dafny.Array(False, 9)
        d_1405_v81_ = nw215_
        d_1406_v82_: _dafny.MultiSet
        d_1406_v82_ = _dafny.MultiSet([d_1405_v81_])
        r0 = (d_1406_v82_).issubset(d_1406_v82_)
        d_1407_v83_: D1
        d_1407_v83_ = D1_DC5(d_1319_v17_, d_1405_v81_, p0)
        r1 = (d_1407_v83_).cf10
        return r0, r1

