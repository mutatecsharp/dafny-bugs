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
    def fm3(globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in (_dafny.Set({956})).Elements:
                d_0_v0_: int = compr_0_
                if (d_0_v0_) in (_dafny.Set({956})):
                    coll0_ = coll0_.union(_dafny.Set([(d_0_v0_) - (642)]))
            return _dafny.Set(coll0_)
        return ((0) - (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqmj"))), len(_dafny.SeqWithoutIsStrInference([True, False]))))) - (len((_dafny.Map({62: len(_dafny.SeqWithoutIsStrInference([iife0_()
 for d_1_i0_ in range(default__.abs(504))]))})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({221, len(_dafny.Map({_dafny.MultiSet([_dafny.CodePoint('x')]): len(_dafny.Set({False, False}))}))}))])): len(_dafny.Set({82, 475}))}))))

    @staticmethod
    def fm4(globalState):
        return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([((D22_DC58(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nugniqpu")), -968, (D21_DC56(False)).cf100)).cf104) or (True)]))

    @staticmethod
    def fm7(p0, p1, p2, p3, globalState):
        if not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpli"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgmxlbhfv")))):
            return (_dafny.MultiSet([923, len(_dafny.Map({(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2_i0_ in range(default__.abs(-974))])))): True}))])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([586])))
        elif True:
            return (_dafny.MultiSet([(_dafny.MultiSet([not(False), True, False])).cardinality])).intersection((D17_DC44(_dafny.MultiSet([505, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_3_i1_ in range(default__.abs(831))]))]))).cf82)

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        source0_ = D11_DC27(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([393, len(_dafny.Set({not(False), False}))]))}))
        if source0_.is_DC28:
            d_4___mcc_h0_ = source0_.cf52
            d_5___mcc_h1_ = source0_.cf53
            d_6___mcc_h2_ = source0_.cf54
            d_7___mcc_h3_ = source0_.cf55
            d_8___mcc_h4_ = source0_.cf56
            d_9_cf56_ = d_8___mcc_h4_
            d_10_cf55_ = d_7___mcc_h3_
            d_11_cf54_ = d_6___mcc_h2_
            d_12_cf53_ = d_5___mcc_h1_
            d_13_cf52_ = d_4___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([d_13_cf52_, d_9_cf56_, d_13_cf52_])
        elif source0_.is_DC29:
            d_14___mcc_h5_ = source0_.cf57
            d_15___mcc_h6_ = source0_.cf58
            d_16___mcc_h7_ = source0_.cf59
            d_17___mcc_h8_ = source0_.cf60
            d_18_cf60_ = d_17___mcc_h8_
            d_19_cf59_ = d_16___mcc_h7_
            d_20_cf58_ = d_15___mcc_h6_
            d_21_cf57_ = d_14___mcc_h5_
            return _dafny.SeqWithoutIsStrInference([d_19_cf59_, d_19_cf59_])
        elif source0_.is_DC30:
            d_22___mcc_h9_ = source0_.cf61
            d_23___mcc_h10_ = source0_.cf62
            d_24___mcc_h11_ = source0_.cf63
            d_25___mcc_h12_ = source0_.cf64
            d_26_cf64_ = d_25___mcc_h12_
            d_27_cf63_ = d_24___mcc_h11_
            d_28_cf62_ = d_23___mcc_h10_
            d_29_cf61_ = d_22___mcc_h9_
            return _dafny.SeqWithoutIsStrInference([d_26_cf64_])
        elif source0_.is_DC27:
            d_30___mcc_h13_ = source0_.cf51
            d_31_cf51_ = d_30___mcc_h13_
            return (_dafny.SeqWithoutIsStrInference([True, not(False)])) + (_dafny.SeqWithoutIsStrInference([True, True]))
        elif True:
            d_32___mcc_h14_ = source0_.cf65
            d_33_cf65_ = d_32___mcc_h14_
            return (_dafny.SeqWithoutIsStrInference([True, True, not(True), False])) + (_dafny.SeqWithoutIsStrInference([not(True), not(True), True]))

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        if (_dafny.MultiSet([len(_dafny.Set({_dafny.CodePoint('t'), _dafny.CodePoint('d'), _dafny.CodePoint('s'), _dafny.CodePoint('e')}))])).isdisjoint(_dafny.MultiSet([824, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqde")))])):
            return D1_DC4(True, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_34_i0_ in range(default__.abs(-541))])), 356)
        elif True:
            return D1_DC4(False, -745, 639)

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        return (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vis"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kiygyprti"))))) - (default__.safeDivisionInt(631, (_dafny.MultiSet([False])).cardinality))

    @staticmethod
    def fm11(globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: str
            for compr_1_ in (_dafny.MultiSet([_dafny.CodePoint('c')])).Elements:
                d_35_v0_: str = compr_1_
                if (d_35_v0_) in (_dafny.MultiSet([_dafny.CodePoint('c')])):
                    coll1_[d_35_v0_] = not((True) or (False))
            return _dafny.Map(coll1_)
        return iife1_()
        

    @staticmethod
    def fm16(globalState):
        return _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwllid")): _dafny.Set({(_dafny.MultiSet([len(_dafny.Set({False, False, False, False, False}))])).cardinality})})

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        return (-760) + ((743) - (727))

    @staticmethod
    def fm19(globalState):
        return (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False, not(False)])})) | ((_dafny.Map({False: _dafny.SeqWithoutIsStrInference([False])})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False])})))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        source1_ = _dafny.Map({False: not(False)})
        d_36___mcc_h0_ = source1_
        d_37_cf50_ = d_36___mcc_h0_
        return D1_DC3(_dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm21(globalState):
        return _dafny.SeqWithoutIsStrInference([(len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ruhfod"))): False}))) >= (len(_dafny.SeqWithoutIsStrInference([True, False]))), not(((_dafny.MultiSet([40])).cardinality) >= ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpluriwse"))), 979, len(_dafny.Map({False: D17_DC47(True, not(True), 234, 465, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vafxu"))))}))]), _dafny.MultiSet([len(_dafny.Set({830, -793})), len(_dafny.Set({not(True)})), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({(_dafny.MultiSet([False, True])).cardinality, -55, len(_dafny.SeqWithoutIsStrInference([138])), (0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vu")) for d_38_i1_ in range(default__.abs(126))]))).cardinality)})) for d_39_i0_ in range(default__.abs(442))])), 464]), _dafny.MultiSet([583])])))))), (843) == (159)])

    @staticmethod
    def fm24(p0, p1, p2, p3, globalState):
        return D0_DC0((len(_dafny.Map({_dafny.CodePoint('k'): len(_dafny.SeqWithoutIsStrInference([622]))}))) < (len(_dafny.Map({False: True}))), (924) + (79), not (True) or (True), (_dafny.MultiSet([_dafny.CodePoint('x'), _dafny.CodePoint('y')])).cardinality, (len(_dafny.SeqWithoutIsStrInference([False, not(False)]))) <= (len(_dafny.SeqWithoutIsStrInference([(D16_DC41(_dafny.CodePoint('p'), not(True))).cf78, False, True, False, True]))))

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        return 606

    @staticmethod
    def fm26(globalState):
        return ((0) - ((len(_dafny.SeqWithoutIsStrInference([True, True, True]))) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmptvjdkk")))))) - (817)

    @staticmethod
    def fm27(p0, globalState):
        return not((_dafny.SeqWithoutIsStrInference([True, False, not(True)])) not in ((_dafny.Set({_dafny.SeqWithoutIsStrInference([not(True)])})) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([True])}))))

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_40_i0_ in range(default__.abs(705))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hniuithge"))))

    @staticmethod
    def fm29(p0, p1, p2, globalState):
        def iife2_():
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: _dafny.Map
                for compr_3_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: -495}), _dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([True]))})])).Elements:
                    d_42_v1_: _dafny.Map = compr_3_
                    if (d_42_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: -495}), _dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([True]))})])):
                        coll3_[d_42_v1_] = -260
                return _dafny.Map(coll3_)
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([816])).Elements:
                d_41_v0_: int = compr_2_
                if (d_41_v0_) in (_dafny.SeqWithoutIsStrInference([816])):
                    coll2_[(d_41_v0_) - (len(iife3_()
                    ))] = False
            return _dafny.Map(coll2_)
        return ((_dafny.Set({D4_DC9(_dafny.Map({True: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality}), len(iife2_()
), False), D4_DC9(_dafny.Map({False: len(_dafny.Map({False: 196}))}), -920, False)})) | (_dafny.Set({D4_DC9(_dafny.Map({True: -591}), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spsf"))), True)}))).intersection(_dafny.Set({D4_DC9(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([True]))}), len(_dafny.Set({True, False, True})), True)}))

    @staticmethod
    def fm30(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([False, True, not(True), True, not(not(True))])) + (_dafny.SeqWithoutIsStrInference([not(not(False)), True, True, True]))

    @staticmethod
    def fm32(p0, globalState):
        return ((_dafny.MultiSet([D17_DC47(False, False, 561, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality), D17_DC47(True, True, 22, len(_dafny.Map({True: 487})), -809)])).intersection(_dafny.MultiSet([D17_DC47(True, True, -449, (0) - (-376), len(_dafny.Set({959})))]))).cardinality

    @staticmethod
    def fm33(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([(D6_DC15(_dafny.SeqWithoutIsStrInference([-960, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(D4_DC9(_dafny.Map({False: 74}), -771, False)).cf23: True})])), -756]))).cf36, _dafny.SeqWithoutIsStrInference([106])])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(_dafny.MultiSet([227])).cardinality, -37])).cardinality for d_43_i0_ in range(default__.abs(28))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.CodePoint('g')}))])]))).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([-403])]))

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([not(True), False])) + (_dafny.SeqWithoutIsStrInference([True, False, True]))) + (_dafny.SeqWithoutIsStrInference([True, not(False), False]))

    @staticmethod
    def fm37(globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_44_i0_ in range(default__.abs(-122))])

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        def iife4_():
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: _dafny.Map
                for compr_6_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Map({False: True})): -969}), _dafny.Map({526: 970}), _dafny.Map({len(_dafny.Set({True, True, True})): len(_dafny.SeqWithoutIsStrInference([True]))})]))).Elements:
                    d_45_v0_: _dafny.Map = compr_6_
                    if (d_45_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Map({False: True})): -969}), _dafny.Map({526: 970}), _dafny.Map({len(_dafny.Set({True, True, True})): len(_dafny.SeqWithoutIsStrInference([True]))})]))):
                        coll6_[d_45_v0_] = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_46_i0_ in range(default__.abs(-765))])))
                return _dafny.Map(coll6_)
            coll4_ = _dafny.Set()
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: _dafny.Map
                for compr_5_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Map({False: True})): -969}), _dafny.Map({526: 970}), _dafny.Map({len(_dafny.Set({True, True, True})): len(_dafny.SeqWithoutIsStrInference([True]))})]))).Elements:
                    d_45_v0_: _dafny.Map = compr_5_
                    if (d_45_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Map({False: True})): -969}), _dafny.Map({526: 970}), _dafny.Map({len(_dafny.Set({True, True, True})): len(_dafny.SeqWithoutIsStrInference([True]))})]))):
                        coll5_[d_45_v0_] = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_46_i0_ in range(default__.abs(-765))])))
                return _dafny.Map(coll5_)
            compr_4_: _dafny.Map
            for compr_4_ in (iife5_()
            ).keys.Elements:
                d_47_v1_: _dafny.Map = compr_4_
                if (d_47_v1_) in (iife6_()
                ):
                    coll4_ = coll4_.union(_dafny.Set([d_47_v1_]))
            return _dafny.Set(coll4_)
        return (_dafny.Set({_dafny.Map({-263: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfkjeqf")))}), _dafny.Map({-747: 409})})) - (iife4_()
        )

    @staticmethod
    def fm39(p0, globalState):
        return (_dafny.MultiSet([True])) | ((_dafny.MultiSet([False, False, False])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False), True]))))

    @staticmethod
    def fm40(p0, globalState):
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnsahm"))])

    @staticmethod
    def fm41(p0, globalState):
        source2_ = D6_DC17(D6_DC16(_dafny.Map({_dafny.Map({False: len(_dafny.Map({not(not(not(False))): (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([-899]))))}))}): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_48_i0_ in range(default__.abs(830))]))}), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iuwsmcdhj")))))
        if source2_.is_DC16:
            d_49___mcc_h0_ = source2_.cf37
            d_50___mcc_h1_ = source2_.cf38
            d_51_cf38_ = d_50___mcc_h1_
            d_52_cf37_ = d_49___mcc_h0_
            if not(True):
                def iife7_():
                    coll7_ = _dafny.Map()
                    compr_7_: int
                    for compr_7_ in _dafny.IntegerRange(609, 701):
                        d_53_v0_: int = compr_7_
                        if ((609) <= (d_53_v0_)) and ((d_53_v0_) < (701)):
                            coll7_[(d_53_v0_) - (len(_dafny.Map({True: True})))] = len(_dafny.Map({True: _dafny.CodePoint('a')}))
                    return _dafny.Map(coll7_)
                return _dafny.Set({iife7_()
                })
            elif True:
                def iife8_():
                    coll8_ = _dafny.Map()
                    compr_8_: int
                    for compr_8_ in (_dafny.SeqWithoutIsStrInference([d_51_cf38_, d_51_cf38_, d_51_cf38_])).Elements:
                        d_54_v1_: int = compr_8_
                        if (d_54_v1_) in (_dafny.SeqWithoutIsStrInference([d_51_cf38_, d_51_cf38_, d_51_cf38_])):
                            coll8_[(d_54_v1_) * (405)] = d_51_cf38_
                    return _dafny.Map(coll8_)
                def iife9_():
                    coll9_ = _dafny.Map()
                    compr_9_: int
                    for compr_9_ in _dafny.IntegerRange(207, 799):
                        d_55_v2_: int = compr_9_
                        if ((207) <= (d_55_v2_)) and ((d_55_v2_) < (799)):
                            coll9_[(d_55_v2_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "alxqvtpj"))))] = d_51_cf38_
                    return _dafny.Map(coll9_)
                return _dafny.Set({_dafny.Map({794: len(_dafny.SeqWithoutIsStrInference([False]))}), iife8_()
                , iife9_()
                , _dafny.Map({d_51_cf38_: d_51_cf38_})})
        elif source2_.is_DC15:
            d_56___mcc_h2_ = source2_.cf36
            d_57_cf36_ = d_56___mcc_h2_
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(586, 527):
                    d_58_v3_: int = compr_10_
                    if ((586) <= (d_58_v3_)) and ((d_58_v3_) < (527)):
                        coll10_[default__.safeModuloInt(d_58_v3_, (_dafny.MultiSet([False, (D12_DC33(False, False, _dafny.CodePoint('g'), (0) - (len(_dafny.SeqWithoutIsStrInference([True, False]))))).cf67])).cardinality)] = len(_dafny.Map({not(True): False}))
                return _dafny.Map(coll10_)
            return _dafny.Set({iife10_()
            })
        elif True:
            d_59___mcc_h3_ = source2_.cf39
            d_60_cf39_ = d_59___mcc_h3_
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: int
                for compr_11_ in (_dafny.MultiSet([len(_dafny.Map({True: len(_dafny.Set({_dafny.CodePoint('i')}))}))])).Elements:
                    d_61_v4_: int = compr_11_
                    if (d_61_v4_) in (_dafny.MultiSet([len(_dafny.Map({True: len(_dafny.Set({_dafny.CodePoint('i')}))}))])):
                        coll11_[(d_61_v4_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yw"))))] = -912
                return _dafny.Map(coll11_)
            return _dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqgro"))): 570}), iife11_()
            })

    @staticmethod
    def fm42(globalState):
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: _dafny.Map
            for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: len(_dafny.Map({757: False}))}), _dafny.Map({not(True): 252})])).Elements:
                d_62_v0_: _dafny.Map = compr_12_
                if (d_62_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: len(_dafny.Map({757: False}))}), _dafny.Map({not(True): 252})])):
                    coll12_[d_62_v0_] = 692
            return _dafny.Map(coll12_)
        return D6_DC17((D6_DC16(_dafny.Map({_dafny.Map({False: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([-828, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")))])): False}))}): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypkhfedox")))}), len(_dafny.Map({841: _dafny.CodePoint('l')}))) if False else D6_DC16(iife12_()
, 604)))

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        return (_dafny.Map({228: _dafny.MultiSet([False])}))

    @staticmethod
    def fm44(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(85, len(_dafny.Map({617: 394})))])

    @staticmethod
    def fm45(p0, globalState):
        return (D1_DC3(_dafny.SeqWithoutIsStrInference([True]))).cf11

    @staticmethod
    def fm46(globalState):
        return (_dafny.Map({False: False})) | ((_dafny.Map({False: False})) | (_dafny.Map({(D9_DC22(True)).cf44: True})))

    @staticmethod
    def fm47(p0, globalState):
        return (_dafny.MultiSet([False, True, False, False, True])) - ((_dafny.MultiSet([False])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False]))))

    @staticmethod
    def fm48(p0, globalState):
        return _dafny.Set({not (False) or (False)})

    @staticmethod
    def fm49(p0, p1, p2, p3, globalState):
        return 617

    @staticmethod
    def fm50(globalState):
        return _dafny.CodePoint('w')

    @staticmethod
    def fm51(globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvdhuqeu"))])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")) for d_63_i0_ in range(default__.abs(-64))])))) | ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghctjwph"))])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iruuh"))])))

    @staticmethod
    def fm52(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gysce"))

    @staticmethod
    def fm53(globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")) for d_64_i0_ in range(default__.abs(-397))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwl")) for d_65_i1_ in range(default__.abs(125))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")), (D1_DC5((D0_DC1(570, False, (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "creav")))])).cardinality, 972)).cf6, len(_dafny.Set({_dafny.CodePoint('t')})), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfryxfmyj")))).cf17, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_66_i2_ in range(default__.abs(389))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxyl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqwwuc"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvofd")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqhinq"))])))

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        return (0) - (-983)

    @staticmethod
    def fm56(p0, globalState):
        def iife13_():
            coll13_ = _dafny.Set()
            compr_13_: _dafny.Map
            for compr_13_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({102: len(_dafny.SeqWithoutIsStrInference([not(True), True, False]))}), _dafny.Map({-268: 401})])).Elements:
                d_67_v0_: _dafny.Map = compr_13_
                if (d_67_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({102: len(_dafny.SeqWithoutIsStrInference([not(True), True, False]))}), _dafny.Map({-268: 401})])):
                    coll13_ = coll13_.union(_dafny.Set([d_67_v0_]))
            return _dafny.Set(coll13_)
        return iife13_()
        

    @staticmethod
    def fm57(p0, p1, globalState):
        return ((_dafny.Map({True: (_dafny.MultiSet([not(False), True])).cardinality})) | (_dafny.Map({False: 356}))) | ((_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([False]))})) | (_dafny.Map({False: len(_dafny.Map({not(True): -933}))})))

    @staticmethod
    def fm58(p0, p1, globalState):
        def iife14_():
            coll14_ = _dafny.Map()
            compr_14_: int
            for compr_14_ in _dafny.IntegerRange(-242, 715):
                d_68_v0_: int = compr_14_
                if ((-242) <= (d_68_v0_)) and ((d_68_v0_) < (715)):
                    coll14_[(d_68_v0_) * (-904)] = 586
            return _dafny.Map(coll14_)
        return ((_dafny.Map({len(_dafny.Map({not(not(True)): False})): (_dafny.MultiSet([False])).cardinality})) | (_dafny.Map({len(_dafny.Map({323: False})): 461}))) | ((iife14_()
        ) | (_dafny.Map({len(_dafny.Set({True, True, False})): 810})))

    @staticmethod
    def fm59(p0, p1, p2, globalState):
        def iife15_():
            coll15_ = _dafny.Map()
            compr_15_: _dafny.Map
            for compr_15_ in (_dafny.MultiSet([_dafny.Map({not(not(True)): True})])).Elements:
                d_69_v0_: _dafny.Map = compr_15_
                if (d_69_v0_) in (_dafny.MultiSet([_dafny.Map({not(not(True)): True})])):
                    coll15_[d_69_v0_] = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykodhj"))): _dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.Set({True}), _dafny.Set({True, True, False})])): 716})})})
            return _dafny.Map(coll15_)
        return (_dafny.Map({_dafny.Map({False: True}): _dafny.Map({len(_dafny.Map({(_dafny.MultiSet([False, False])).cardinality: True})): _dafny.Set({_dafny.Map({241: 364})})})})) | (iife15_()
        )

    @staticmethod
    def fm60(p0, p1, p2, p3, globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icuwbr"))): not((_dafny.Set({True})).isdisjoint(_dafny.Set({not(False)})))})

    @staticmethod
    def fm61(p0, p1, globalState):
        return D0_DC0((True) or (False), 130, False, len((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.Set({False}), _dafny.Set({not(True)})]))])) + (_dafny.SeqWithoutIsStrInference([107 for d_70_i0_ in range(default__.abs(510))]))), (len(_dafny.Set({True, False}))) != (141))

    @staticmethod
    def fm62(p0, p1, p2, globalState):
        return _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vmsmhd")), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkq"))) + ((D1_DC5(True, 437, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esln")))).cf17), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obr")), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqybmrr"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxvbtvfen"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfekrrf"))})

    @staticmethod
    def fm63(p0, p1, globalState):
        source3_ = D12_DC32(_dafny.MultiSet([_dafny.CodePoint('e'), _dafny.CodePoint('w'), _dafny.CodePoint('w')]))
        if source3_.is_DC33:
            d_71___mcc_h0_ = source3_.cf67
            d_72___mcc_h1_ = source3_.cf68
            d_73___mcc_h2_ = source3_.cf69
            d_74___mcc_h3_ = source3_.cf70
            d_75_cf70_ = d_74___mcc_h3_
            d_76_cf69_ = d_73___mcc_h2_
            d_77_cf68_ = d_72___mcc_h1_
            d_78_cf67_ = d_71___mcc_h0_
            return D5_DC12((D5_DC12(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vljeskjse"))}))).cf28)
        elif True:
            d_79___mcc_h4_ = source3_.cf66
            d_80_cf66_ = d_79___mcc_h4_
            def iife16_():
                coll16_ = _dafny.Set()
                compr_16_: _dafny.Seq
                for compr_16_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jo")): True})).keys.Elements:
                    d_81_v0_: _dafny.Seq = compr_16_
                    if (d_81_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jo")): True})):
                        coll16_ = coll16_.union(_dafny.Set([d_81_v0_]))
                return _dafny.Set(coll16_)
            return D5_DC12(iife16_()
)

    @staticmethod
    def fm64(p0, p1, p2, p3, globalState):
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: _dafny.Map
            for compr_17_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): True})])).Elements:
                d_82_v0_: _dafny.Map = compr_17_
                if (d_82_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): True})])):
                    coll17_[d_82_v0_] = True
            return _dafny.Map(coll17_)
        return (_dafny.Map({_dafny.Map({False: True}): False})) | ((_dafny.Map({_dafny.Map({False: True}): True})) | (iife17_()
        ))

    @staticmethod
    def fm65(p0, p1, p2, p3, globalState):
        def iife18_():
            coll18_ = _dafny.Set()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(838, 647):
                d_83_v0_: int = compr_18_
                if ((838) <= (d_83_v0_)) and ((d_83_v0_) < (647)):
                    coll18_ = coll18_.union(_dafny.Set([(d_83_v0_) * ((_dafny.MultiSet([False, False])).cardinality)]))
            return _dafny.Set(coll18_)
        return (_dafny.MultiSet([_dafny.Set({983})])) | (_dafny.MultiSet([_dafny.Set({769}), iife18_()
        ]))

    @staticmethod
    def fm66(globalState):
        def iife19_():
            coll19_ = _dafny.Map()
            compr_19_: D4
            for compr_19_ in (_dafny.SeqWithoutIsStrInference([D4_DC10(False, True, _dafny.SeqWithoutIsStrInference([not(False)])) for d_84_i0_ in range(default__.abs(188))])).Elements:
                d_85_v0_: D4 = compr_19_
                if (d_85_v0_) in (_dafny.SeqWithoutIsStrInference([D4_DC10(False, True, _dafny.SeqWithoutIsStrInference([not(False)])) for d_84_i0_ in range(default__.abs(188))])):
                    coll19_[d_85_v0_] = 547
            return _dafny.Map(coll19_)
        return iife19_()
        

    @staticmethod
    def fm67(globalState):
        return (_dafny.SeqWithoutIsStrInference([D4_DC10(True, False, _dafny.SeqWithoutIsStrInference([False]))])) + (_dafny.SeqWithoutIsStrInference([D4_DC10(not(False), False, _dafny.SeqWithoutIsStrInference([True]))]))

    @staticmethod
    def fm68(globalState):
        def iife20_():
            coll20_ = _dafny.Map()
            compr_20_: int
            for compr_20_ in _dafny.IntegerRange(729, 7):
                d_86_v0_: int = compr_20_
                if ((729) <= (d_86_v0_)) and ((d_86_v0_) < (7)):
                    coll20_[(d_86_v0_) - (896)] = (0) - (len(_dafny.Set({True})))
            return _dafny.Map(coll20_)
        return ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.Set({968}))])) - (_dafny.MultiSet([len(_dafny.Map({True: False})), len(_dafny.Map({len(iife20_()
        ): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))}))]))).intersection(_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([649 for d_87_i0_ in range(default__.abs(906))])) + (_dafny.SeqWithoutIsStrInference([127, -812]))))

    @staticmethod
    def fm69(globalState):
        return ((_dafny.Map({D6_DC17(D6_DC16(_dafny.Map({_dafny.Map({False: -149}): -391}), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmclfmyx"))))): not(not(True))})) | (_dafny.Map({D6_DC17(D6_DC16(_dafny.Map({_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqr")))}): -809}), 793)): not(True)}))) | (_dafny.Map({D6_DC17(D6_DC17(D6_DC15(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([606 for d_88_i0_ in range(default__.abs(310))])), len(_dafny.SeqWithoutIsStrInference([-1 for d_89_i1_ in range(default__.abs(401))]))])))): True}))

    @staticmethod
    def fm70(globalState):
        return D12_DC32((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p'), _dafny.CodePoint('v')]))) - (_dafny.MultiSet([_dafny.CodePoint('m')])))

    @staticmethod
    def fm71(p0, p1, p2, p3, globalState):
        def iife21_():
            coll21_ = _dafny.Set()
            compr_21_: int
            for compr_21_ in _dafny.IntegerRange(545, 995):
                d_90_v0_: int = compr_21_
                if ((545) <= (d_90_v0_)) and ((d_90_v0_) < (995)):
                    coll21_ = coll21_.union(_dafny.Set([(d_90_v0_) + (len(_dafny.Map({True: len(_dafny.Map({True: 847}))})))]))
            return _dafny.Set(coll21_)
        def iife22_():
            coll22_ = _dafny.Set()
            compr_22_: int
            for compr_22_ in _dafny.IntegerRange(110, 637):
                d_91_v1_: int = compr_22_
                if ((110) <= (d_91_v1_)) and ((d_91_v1_) < (637)):
                    coll22_ = coll22_.union(_dafny.Set([(d_91_v1_) * (125)]))
            return _dafny.Set(coll22_)
        return D4_DC9(_dafny.Map({True: 355}), len((iife21_()
).intersection(iife22_()
)), (False) == (True))

    @staticmethod
    def fm72(globalState):
        return D17_DC45()

    @staticmethod
    def fm73(p0, globalState):
        source4_ = D4_DC9(_dafny.Map({False: 636}), -170, True)
        if source4_.is_DC9:
            d_92___mcc_h0_ = source4_.cf21
            d_93___mcc_h1_ = source4_.cf22
            d_94___mcc_h2_ = source4_.cf23
            d_95_cf23_ = d_94___mcc_h2_
            d_96_cf22_ = d_93___mcc_h1_
            d_97_cf21_ = d_92___mcc_h0_
            return D14_DC38(d_95_cf23_)
        elif source4_.is_DC10:
            d_98___mcc_h3_ = source4_.cf24
            d_99___mcc_h4_ = source4_.cf25
            d_100___mcc_h5_ = source4_.cf26
            d_101_cf26_ = d_100___mcc_h5_
            d_102_cf25_ = d_99___mcc_h4_
            d_103_cf24_ = d_98___mcc_h3_
            return D14_DC38(True)
        elif source4_.is_DC8:
            d_104___mcc_h6_ = source4_.cf20
            d_105_cf20_ = d_104___mcc_h6_
            return D14_DC38(True)
        elif True:
            d_106___mcc_h7_ = source4_.cf27
            d_107_cf27_ = d_106___mcc_h7_
            return D14_DC38(False)

    @staticmethod
    def fm74(p0, globalState):
        def iife23_():
            coll23_ = _dafny.Set()
            compr_23_: _dafny.Seq
            for compr_23_ in ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([644 for d_108_i1_ in range(default__.abs(151))]) for d_109_i0_ in range(default__.abs(-931))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-651 for d_110_i3_ in range(default__.abs(-17))])) for d_111_i2_ in range(default__.abs(681))]))]), _dafny.SeqWithoutIsStrInference([450 for d_112_i4_ in range(default__.abs(-215))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))])]))).Elements:
                d_113_v0_: _dafny.Seq = compr_23_
                if (d_113_v0_) in ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([644 for d_108_i1_ in range(default__.abs(151))]) for d_109_i0_ in range(default__.abs(-931))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-651 for d_110_i3_ in range(default__.abs(-17))])) for d_111_i2_ in range(default__.abs(681))]))]), _dafny.SeqWithoutIsStrInference([450 for d_112_i4_ in range(default__.abs(-215))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))])]))):
                    coll23_ = coll23_.union(_dafny.Set([d_113_v0_]))
            return _dafny.Set(coll23_)
        return iife23_()
        

    @staticmethod
    def fm75(p0, globalState):
        return D17_DC47(not (False) or (True), (True) in (_dafny.SeqWithoutIsStrInference([False, True])), ((_dafny.MultiSet([True])).cardinality) * (-220), -266, 716)

    @staticmethod
    def fm76(p0, p1, p2, globalState):
        def iife24_():
            coll24_ = _dafny.Set()
            compr_24_: int
            for compr_24_ in _dafny.IntegerRange(-240, 808):
                d_114_v0_: int = compr_24_
                if ((-240) <= (d_114_v0_)) and ((d_114_v0_) < (808)):
                    coll24_ = coll24_.union(_dafny.Set([(d_114_v0_) - (len(_dafny.Map({False: 166})))]))
            return _dafny.Set(coll24_)
        return iife24_()
        

    @staticmethod
    def fm77(globalState):
        return (_dafny.Map({-658: _dafny.Map({383: (0) - (-334)})}))

    @staticmethod
    def fm78(p0, p1, p2, p3, globalState):
        if False:
            if False:
                return _dafny.Map({True: _dafny.MultiSet([904, -311, 245, 818])})
            elif True:
                return _dafny.Map({not(False): _dafny.MultiSet([548])})
        elif True:
            return _dafny.Map({True: _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))])})

    @staticmethod
    def fm79(p0, p1, globalState):
        return D18_DC50((_dafny.MultiSet([D19_DC53((_dafny.MultiSet([35])).cardinality, 334, False, True)])).isdisjoint(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D19_DC53(len(_dafny.SeqWithoutIsStrInference([True])), 27, False, False), D19_DC53(-359, 889, False, False)]))), False)

    @staticmethod
    def fm80(p0, globalState):
        source5_ = D1_DC3(_dafny.SeqWithoutIsStrInference([False, not(True), True]))
        if source5_.is_DC4:
            d_115___mcc_h0_ = source5_.cf12
            d_116___mcc_h1_ = source5_.cf13
            d_117___mcc_h2_ = source5_.cf14
            d_118_cf14_ = d_117___mcc_h2_
            d_119_cf13_ = d_116___mcc_h1_
            d_120_cf12_ = d_115___mcc_h0_
            return D17_DC48(D17_DC47(True, d_120_cf12_, d_118_cf14_, d_119_cf13_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_121_i0_ in range(default__.abs(623))]))))
        elif source5_.is_DC5:
            d_122___mcc_h3_ = source5_.cf15
            d_123___mcc_h4_ = source5_.cf16
            d_124___mcc_h5_ = source5_.cf17
            d_125_cf17_ = d_124___mcc_h5_
            d_126_cf16_ = d_123___mcc_h4_
            d_127_cf15_ = d_122___mcc_h3_
            return D17_DC48(D17_DC47(not(d_127_cf15_), d_127_cf15_, 615, len(_dafny.SeqWithoutIsStrInference([438])), d_126_cf16_))
        elif True:
            d_128___mcc_h6_ = source5_.cf11
            d_129_cf11_ = d_128___mcc_h6_
            return D17_DC48(D17_DC45())

    @staticmethod
    def fm81(p0, p1, globalState):
        def iife25_():
            coll25_ = _dafny.Map()
            compr_25_: int
            for compr_25_ in _dafny.IntegerRange(215, 821):
                d_130_v0_: int = compr_25_
                if ((215) <= (d_130_v0_)) and ((d_130_v0_) < (821)):
                    coll25_[(d_130_v0_) + (833)] = _dafny.CodePoint('w')
            return _dafny.Map(coll25_)
        return iife25_()
        

    @staticmethod
    def fm82(p0, globalState):
        def iife26_():
            coll26_ = _dafny.Map()
            compr_26_: D9
            for compr_26_ in (_dafny.SeqWithoutIsStrInference([D9_DC23(False, (0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality), True)])).Elements:
                d_131_v0_: D9 = compr_26_
                if (d_131_v0_) in (_dafny.SeqWithoutIsStrInference([D9_DC23(False, (0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality), True)])):
                    coll26_[d_131_v0_] = (_dafny.MultiSet([_dafny.CodePoint('n')])).cardinality
            return _dafny.Map(coll26_)
        def iife27_():
            coll27_ = _dafny.Map()
            compr_27_: int
            for compr_27_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): (_dafny.MultiSet([False])).cardinality})).keys.Elements:
                d_132_v1_: int = compr_27_
                if (d_132_v1_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): (_dafny.MultiSet([False])).cardinality})):
                    coll27_[(d_132_v1_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_133_i1_ in range(default__.abs(874))])))] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppcgyghvm")))))]))).cardinality
            return _dafny.Map(coll27_)
        def iife28_():
            coll28_ = _dafny.Map()
            compr_28_: D9
            for compr_28_ in (_dafny.SeqWithoutIsStrInference([D9_DC23(False, -347, False)])).Elements:
                d_135_v2_: D9 = compr_28_
                if (d_135_v2_) in (_dafny.SeqWithoutIsStrInference([D9_DC23(False, -347, False)])):
                    coll28_[d_135_v2_] = 356
            return _dafny.Map(coll28_)
        return ((_dafny.SeqWithoutIsStrInference([iife26_()
        ])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({D9_DC23((D0_DC1((D9_DC24(len(_dafny.Set({_dafny.CodePoint('t')})))).cf48, False, 348, -164)).cf6, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')])), True): len(_dafny.SeqWithoutIsStrInference([False, False, not(not(False)), True, True]))})]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.Map({D9_DC23(False, len(iife27_()
), False): len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True])]))}) for d_134_i0_ in range(default__.abs(-106))])) + (_dafny.SeqWithoutIsStrInference([iife28_()
        ])))

    @staticmethod
    def fm83(p0, globalState):
        return D9_DC23((_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_136_i0_ in range(default__.abs(888))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_137_i1_ in range(default__.abs(143))])})).isdisjoint(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "titbebql")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_138_i2_ in range(default__.abs(255))])})), 702, not (False) or (True))

    @staticmethod
    def fm84(globalState):
        return D1_DC5(True, default__.safeDivisionInt((0) - (-939), 718), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tcbk")))

    @staticmethod
    def fm85(globalState):
        return _dafny.Map({False: (0) - (len(_dafny.SeqWithoutIsStrInference([not(False)])))})

    @staticmethod
    def fm86(p0, p1, p2, globalState):
        source6_ = D11_DC28(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nixa"))), 767, False, not(not(False)))
        if source6_.is_DC28:
            d_139___mcc_h0_ = source6_.cf52
            d_140___mcc_h1_ = source6_.cf53
            d_141___mcc_h2_ = source6_.cf54
            d_142___mcc_h3_ = source6_.cf55
            d_143___mcc_h4_ = source6_.cf56
            d_144_cf56_ = d_143___mcc_h4_
            d_145_cf55_ = d_142___mcc_h3_
            d_146_cf54_ = d_141___mcc_h2_
            d_147_cf53_ = d_140___mcc_h1_
            d_148_cf52_ = d_139___mcc_h0_
            def iife29_():
                coll29_ = _dafny.Set()
                compr_29_: _dafny.Map
                for compr_29_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: d_146_cf54_}) for d_149_i0_ in range(default__.abs(917))])).Elements:
                    d_150_v0_: _dafny.Map = compr_29_
                    if (d_150_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: d_146_cf54_}) for d_149_i0_ in range(default__.abs(917))])):
                        coll29_ = coll29_.union(_dafny.Set([d_150_v0_]))
                return _dafny.Set(coll29_)
            return (_dafny.Set({_dafny.Map({d_144_cf56_: len(_dafny.SeqWithoutIsStrInference([d_148_cf52_, False, False, d_144_cf56_]))})})) | (iife29_()
            )
        elif source6_.is_DC29:
            d_151___mcc_h5_ = source6_.cf57
            d_152___mcc_h6_ = source6_.cf58
            d_153___mcc_h7_ = source6_.cf59
            d_154___mcc_h8_ = source6_.cf60
            d_155_cf60_ = d_154___mcc_h8_
            d_156_cf59_ = d_153___mcc_h7_
            d_157_cf58_ = d_152___mcc_h6_
            d_158_cf57_ = d_151___mcc_h5_
            def iife30_():
                coll30_ = _dafny.Set()
                compr_30_: _dafny.Map
                for compr_30_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_156_cf59_: 984}) for d_159_i1_ in range(default__.abs(505))])).Elements:
                    d_160_v1_: _dafny.Map = compr_30_
                    if (d_160_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_156_cf59_: 984}) for d_159_i1_ in range(default__.abs(505))])):
                        coll30_ = coll30_.union(_dafny.Set([d_160_v1_]))
                return _dafny.Set(coll30_)
            return iife30_()
            
        elif source6_.is_DC30:
            d_161___mcc_h9_ = source6_.cf61
            d_162___mcc_h10_ = source6_.cf62
            d_163___mcc_h11_ = source6_.cf63
            d_164___mcc_h12_ = source6_.cf64
            d_165_cf64_ = d_164___mcc_h12_
            d_166_cf63_ = d_163___mcc_h11_
            d_167_cf62_ = d_162___mcc_h10_
            d_168_cf61_ = d_161___mcc_h9_
            return _dafny.Set({_dafny.Map({d_168_cf61_: (0) - (d_166_cf63_)})})
        elif source6_.is_DC27:
            d_169___mcc_h13_ = source6_.cf51
            d_170_cf51_ = d_169___mcc_h13_
            return _dafny.Set({_dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gp")))})})
        elif True:
            d_171___mcc_h14_ = source6_.cf65
            d_172_cf65_ = d_171___mcc_h14_
            return (_dafny.Set({_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))})})).intersection(_dafny.Set({_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([False, not(True)]))}), _dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))})}))

    @staticmethod
    def fm87(p0, p1, globalState):
        def iife31_():
            coll31_ = _dafny.Set()
            compr_31_: int
            for compr_31_ in (_dafny.Map({502: True})).keys.Elements:
                d_174_v0_: int = compr_31_
                if (d_174_v0_) in (_dafny.Map({502: True})):
                    coll31_ = coll31_.union(_dafny.Set([default__.safeDivisionInt(d_174_v0_, 710)]))
            return _dafny.Set(coll31_)
        def iife32_():
            coll32_ = _dafny.Set()
            compr_32_: int
            for compr_32_ in _dafny.IntegerRange(193, -346):
                d_175_v1_: int = compr_32_
                if ((193) <= (d_175_v1_)) and ((d_175_v1_) < (-346)):
                    def iife33_():
                        coll33_ = _dafny.Set()
                        compr_33_: int
                        for compr_33_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y')])) for d_176_i1_ in range(default__.abs(963))])).Elements:
                            d_177_v2_: int = compr_33_
                            if (d_177_v2_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y')])) for d_176_i1_ in range(default__.abs(963))])):
                                coll33_ = coll33_.union(_dafny.Set([default__.safeDivisionInt(d_177_v2_, len(_dafny.Map({554: False})))]))
                        return _dafny.Set(coll33_)
                    coll32_ = coll32_.union(_dafny.Set([default__.safeModuloInt(d_175_v1_, (0) - (len(_dafny.Map({len(_dafny.Set({(_dafny.MultiSet([len(iife33_()
)])).cardinality})): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvf"))}))))]))
            return _dafny.Set(coll32_)
        return ((_dafny.MultiSet([D11_DC27(_dafny.Set({644, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qejo")))}))])) | (_dafny.MultiSet([D11_DC27(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))), (_dafny.MultiSet([False, True, False, True])).cardinality})), D11_DC27(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_173_i0_ in range(default__.abs(514))])), -408])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")))})), D11_DC27(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet((D6_DC15(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.Map({False: False}))]))).cf36)).cardinality])]))})), D11_DC27(iife31_()
)]))).intersection((_dafny.MultiSet([D11_DC27(iife32_()
)]) if True else _dafny.MultiSet([D11_DC27(_dafny.Set({len(_dafny.Map({False: False}))})), D11_DC27(_dafny.Set({522, 424}))])))

    @staticmethod
    def fm88(p0, p1, p2, globalState):
        return D11_DC27(_dafny.Set({(_dafny.MultiSet([False])).cardinality, len(_dafny.Map({_dafny.CodePoint('t'): len(_dafny.Map({True: 41}))}))}))

    @staticmethod
    def m0(p0, p1, globalState):
        r0: bool = False
        d_178_v0_: _dafny.Seq
        d_178_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcfoxtf"))
        d_179_v1_: str
        d_179_v1_ = _dafny.CodePoint('r')
        if (default__.fm49(d_178_v0_, p1, p0, p0, globalState)) >= (len((d_178_v0_).set(default__.safeIndex((0) - (p1), len(d_178_v0_)), d_179_v1_))):
            d_180_v2_: _dafny.Map
            d_180_v2_ = _dafny.Map({p0: False})
            d_181_v3_: C4
            nw0_ = C4()
            nw0_.ctor__(p1, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")), True, p0, p1, len(d_180_v2_), p0)
            d_181_v3_ = nw0_
            d_182_v4_: D22
            d_182_v4_ = D22_DC58(d_178_v0_, len(d_178_v0_), (d_181_v3_).f24)
            d_183_v5_: D22
            d_183_v5_ = D22_DC60(d_182_v4_)
            d_184_v6_: _dafny.Map
            d_184_v6_ = _dafny.Map({_dafny.Map({True: d_181_v3_}): d_183_v5_})
            d_185_v7_: _dafny.Map
            d_185_v7_ = _dafny.Map({(d_181_v3_).f24: d_181_v3_})
            d_184_v6_ = (d_184_v6_).set(d_185_v7_, d_183_v5_)
            d_186_v8_: _dafny.Seq
            d_186_v8_ = _dafny.SeqWithoutIsStrInference([(d_181_v3_).f24, (d_181_v3_).f24, False, p0])
            d_187_v9_: C1
            nw1_ = C1()
            nw1_.ctor__((d_181_v3_).f23, (d_181_v3_).f24, p0, len(d_186_v8_))
            d_187_v9_ = nw1_
            d_188_v10_: _dafny.Array
            nw2_ = _dafny.Array(_dafny.Set({}), 7)
            d_188_v10_ = nw2_
            d_188_v10_ = d_188_v10_
            d_189_v11_: _dafny.Array
            def lambda0_(d_190_v3_):
                def lambda1_(d_191_i0_):
                    return _dafny.MultiSet([(d_190_v3_).f24, (d_190_v3_).f24])

                return lambda1_

            init0_ = lambda0_(d_181_v3_)
            nw3_ = _dafny.Array(None, 25)
            for i0_0_ in range(nw3_.length(0)):
                nw3_[i0_0_] = init0_(i0_0_)
            d_189_v11_ = nw3_
            d_192_v12_: _dafny.MultiSet
            d_192_v12_ = _dafny.MultiSet([p0, (d_181_v3_).f24])
            d_193_v13_: _dafny.MultiSet
            d_193_v13_ = d_192_v12_
            index0_ = default__.safeIndex(843, (d_189_v11_).length(0))
            (d_189_v11_)[index0_] = (d_193_v13_)
            index1_ = default__.safeIndex(843, (d_189_v11_).length(0))
            (d_189_v11_)[index1_] = (d_192_v12_ if not((d_181_v3_).f24) else d_192_v12_)
            d_194_v14_: T2
            nw4_ = C4()
            nw4_.ctor__(855, p0, d_178_v0_, (d_181_v3_).f24, p0, p1, (d_181_v3_).f23, True)
            d_194_v14_ = nw4_
            d_195_v15_: _dafny.Seq
            d_195_v15_ = _dafny.SeqWithoutIsStrInference([d_194_v14_, d_194_v14_, d_194_v14_, d_194_v14_, d_194_v14_])
            d_196_v16_: _dafny.Map
            d_196_v16_ = _dafny.Map({(d_194_v14_).f5: d_194_v14_})
            d_195_v15_ = (d_195_v15_).set(default__.safeIndex(p1, len(d_195_v15_)), ((d_196_v16_)[(d_181_v3_).f23] if ((d_181_v3_).f23) in (d_196_v16_) else d_194_v14_))
        elif True:
            d_197_v17_: _dafny.Map
            d_197_v17_ = _dafny.Map({p0: p0})
            d_198_v18_: _dafny.Map
            d_198_v18_ = _dafny.Map({d_197_v17_: p0})
            d_199_v19_: _dafny.Array
            nw5_ = _dafny.Array(None, 4)
            nw5_[int(0)] = p0
            nw5_[int(1)] = p0
            nw5_[int(2)] = p0
            nw5_[int(3)] = p0
            d_199_v19_ = nw5_
            d_200_v20_: _dafny.Map
            d_200_v20_ = _dafny.Map({p0: d_199_v19_})
            source7_ = default__.fm88(p0, ((d_198_v18_)[d_197_v17_] if (d_197_v17_) in (d_198_v18_) else False), len(d_200_v20_), globalState)
            if source7_.is_DC28:
                d_201___mcc_h0_ = source7_.cf52
                d_202___mcc_h1_ = source7_.cf53
                d_203___mcc_h2_ = source7_.cf54
                d_204___mcc_h3_ = source7_.cf55
                d_205___mcc_h4_ = source7_.cf56
                d_206_cf56_ = d_205___mcc_h4_
                d_207_cf55_ = d_204___mcc_h3_
                d_208_cf54_ = d_203___mcc_h2_
                d_209_cf53_ = d_202___mcc_h1_
                d_210_cf52_ = d_201___mcc_h0_
                d_211_v21_: C14
                nw6_ = C14()
                nw6_.ctor__(d_207_cf55_, p1)
                d_211_v21_ = nw6_
                d_212_v22_: C8
                nw7_ = C8()
                nw7_.ctor__(d_178_v0_, p1)
                d_212_v22_ = nw7_
                d_213_v23_: _dafny.Array
                nw8_ = _dafny.Array(int(0), 5)
                d_213_v23_ = nw8_
                index2_ = default__.safeIndex(783, (d_213_v23_).length(0))
                (d_213_v23_)[index2_] = len((d_212_v22_.f13).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_210_cf52_])) for d_214_i1_ in range(default__.abs(80))])), len(d_212_v22_.f13)), d_179_v1_))
                d_215_v24_: _dafny.Seq
                d_215_v24_ = _dafny.SeqWithoutIsStrInference([d_199_v19_, d_199_v19_])
                index3_ = default__.safeIndex(783, (d_213_v23_).length(0))
                (d_213_v23_)[index3_] = len((_dafny.SeqWithoutIsStrInference([d_199_v19_, d_199_v19_, d_199_v19_])) + (d_215_v24_))
                d_216_v25_: D17
                d_216_v25_ = D17_DC47(d_210_cf52_, p0, (d_213_v23_)[default__.safeIndex(783, (d_213_v23_).length(0))], (d_213_v23_)[default__.safeIndex(783, (d_213_v23_).length(0))], 513)
                d_217_v26_: T0
                nw9_ = C7()
                nw9_.ctor__((0) - (d_208_cf54_), True, d_211_v21_.f2, d_208_cf54_, p1, d_207_cf55_, d_212_v22_.f13, d_207_cf55_)
                d_217_v26_ = nw9_
                d_218_v27_: _dafny.Map
                d_218_v27_ = _dafny.Map({d_208_cf54_: d_217_v26_})
                d_219_v28_: _dafny.Map
                d_219_v28_ = _dafny.Map({d_218_v27_: p0})
                d_220_v29_: _dafny.Set
                d_220_v29_ = _dafny.Set({d_211_v21_.f2, True})
                pat_let_tv0_ = d_207_cf55_
                pat_let_tv1_ = d_211_v21_
                pat_let_tv2_ = d_207_cf55_
                d_221_v30_: _dafny.Array
                nw10_ = _dafny.Array(None, 19)
                nw10_[int(0)] = d_216_v25_
                def iife34_(_pat_let0_0):
                    def iife35_(d_222_dt__update__tmp_h0_):
                        def iife36_(_pat_let1_0):
                            def iife37_(d_223_dt__update_hcf87_h0_):
                                return D17_DC47((d_222_dt__update__tmp_h0_).cf83, (d_222_dt__update__tmp_h0_).cf84, (d_222_dt__update__tmp_h0_).cf85, (d_222_dt__update__tmp_h0_).cf86, d_223_dt__update_hcf87_h0_)
                            return iife37_(_pat_let1_0)
                        return iife36_((D1_DC4(pat_let_tv0_, -835, 796)).cf14)
                    return iife35_(_pat_let0_0)
                nw10_[int(1)] = iife34_(d_216_v25_)
                nw10_[int(2)] = d_216_v25_
                nw10_[int(3)] = D17_DC47(d_207_cf55_, d_210_cf52_, (d_212_v22_).f14, d_209_cf53_, d_209_cf53_)
                nw10_[int(4)] = d_216_v25_
                nw10_[int(5)] = d_216_v25_
                nw10_[int(6)] = d_216_v25_
                nw10_[int(7)] = D17_DC47((D0_DC2(p0, False)).cf10, d_210_cf52_, len(d_219_v28_), len(d_198_v18_), (d_213_v23_)[default__.safeIndex(783, (d_213_v23_).length(0))])
                nw10_[int(8)] = d_216_v25_
                nw10_[int(9)] = d_216_v25_
                def iife38_(_pat_let2_0):
                    def iife39_(d_224_dt__update__tmp_h1_):
                        def iife40_(_pat_let3_0):
                            def iife41_(d_225_dt__update_hcf85_h0_):
                                return D17_DC47((d_224_dt__update__tmp_h1_).cf83, (d_224_dt__update__tmp_h1_).cf84, d_225_dt__update_hcf85_h0_, (d_224_dt__update__tmp_h1_).cf86, (d_224_dt__update__tmp_h1_).cf87)
                            return iife41_(_pat_let3_0)
                        return iife40_((pat_let_tv1_).f3)
                    return iife39_(_pat_let2_0)
                nw10_[int(10)] = iife38_(d_216_v25_)
                nw10_[int(11)] = d_216_v25_
                nw10_[int(12)] = D17_DC47(d_207_cf55_, p0, (d_212_v22_).f14, len(d_220_v29_), (d_213_v23_)[default__.safeIndex(783, (d_213_v23_).length(0))])
                nw10_[int(13)] = d_216_v25_
                nw10_[int(14)] = d_216_v25_
                nw10_[int(15)] = d_216_v25_
                nw10_[int(16)] = d_216_v25_
                nw10_[int(17)] = default__.fm75(d_206_cf56_, globalState)
                def iife42_(_pat_let4_0):
                    def iife43_(d_226_dt__update__tmp_h2_):
                        def iife44_(_pat_let5_0):
                            def iife45_(d_227_dt__update_hcf84_h0_):
                                return D17_DC47((d_226_dt__update__tmp_h2_).cf83, d_227_dt__update_hcf84_h0_, (d_226_dt__update__tmp_h2_).cf85, (d_226_dt__update__tmp_h2_).cf86, (d_226_dt__update__tmp_h2_).cf87)
                            return iife45_(_pat_let5_0)
                        return iife44_(not(pat_let_tv2_))
                    return iife43_(_pat_let4_0)
                nw10_[int(18)] = iife42_(default__.fm75((d_217_v26_).f4, globalState))
                d_221_v30_ = nw10_
                d_228_v31_: T1
                nw11_ = C3()
                nw11_.ctor__(d_179_v1_, len(d_220_v29_), d_211_v21_.f2, d_210_cf52_, d_208_cf54_)
                d_228_v31_ = nw11_
                d_229_v32_: _dafny.Map
                d_229_v32_ = _dafny.Map({(d_212_v22_).f14: d_228_v31_})
                d_230_v33_: _dafny.Seq
                d_230_v33_ = _dafny.SeqWithoutIsStrInference([d_229_v32_, d_229_v32_, d_229_v32_])
                d_231_v34_: _dafny.Map
                d_231_v34_ = _dafny.Map({(d_230_v33_)[default__.safeIndex((d_217_v26_).f5, len(d_230_v33_))]: d_206_cf56_})
                d_232_v35_: _dafny.Map
                d_232_v35_ = _dafny.Map({(d_228_v31_).f5: (d_228_v31_).f15})
                rhs0_ = (d_217_v26_).f4
                rhs1_ = d_209_cf53_
                rhs2_ = ((d_231_v34_)[(_dafny.Map({len(d_232_v35_): d_228_v31_})).set((d_211_v21_).f3, d_228_v31_)] if ((_dafny.Map({len(d_232_v35_): d_228_v31_})).set((d_211_v21_).f3, d_228_v31_)) in (d_231_v34_) else True)
                rhs3_ = d_221_v30_
                r0 = rhs0_
                d_209_cf53_ = rhs1_
                d_206_cf56_ = rhs2_
                d_221_v30_ = rhs3_
            elif source7_.is_DC29:
                d_233___mcc_h5_ = source7_.cf57
                d_234___mcc_h6_ = source7_.cf58
                d_235___mcc_h7_ = source7_.cf59
                d_236___mcc_h8_ = source7_.cf60
                d_237_cf60_ = d_236___mcc_h8_
                d_238_cf59_ = d_235___mcc_h7_
                d_239_cf58_ = d_234___mcc_h6_
                d_240_cf57_ = d_233___mcc_h5_
                d_240_cf57_ = 885
                d_241_v36_: D14
                d_241_v36_ = D14_DC38(p0)
                d_241_v36_ = d_241_v36_
                d_242_v37_: C5
                nw12_ = C5()
                nw12_.ctor__()
                d_242_v37_ = nw12_
                index4_ = default__.safeIndex(40, (d_199_v19_).length(0))
                (d_199_v19_)[index4_] = d_237_cf60_
                d_243_v38_: _dafny.Array
                nw13_ = _dafny.Array(_dafny.MultiSet({}), 9)
                d_243_v38_ = nw13_
                d_244_v39_: _dafny.MultiSet
                d_244_v39_ = _dafny.MultiSet([d_178_v0_, d_178_v0_, _dafny.SeqWithoutIsStrInference([d_179_v1_ for d_245_i2_ in range(default__.abs(503))]), _dafny.SeqWithoutIsStrInference([d_179_v1_ for d_246_i3_ in range(default__.abs(582))]), d_178_v0_])
                index5_ = default__.safeIndex(313, (d_243_v38_).length(0))
                (d_243_v38_)[index5_] = d_244_v39_
                d_247_v40_: _dafny.Seq
                d_247_v40_ = _dafny.SeqWithoutIsStrInference([(d_237_cf60_ if True else p0)])
                index6_ = default__.safeIndex(40, (d_199_v19_).length(0))
                index7_ = default__.safeIndex(313, (d_243_v38_).length(0))
                rhs4_ = (d_247_v40_)[default__.safeIndex(p1, len(d_247_v40_))]
                rhs5_ = d_244_v39_
                lhs0_ = d_199_v19_
                lhs1_ = default__.safeIndex(40, (d_199_v19_).length(0))
                lhs2_ = d_243_v38_
                lhs3_ = default__.safeIndex(313, (d_243_v38_).length(0))
                lhs0_[lhs1_] = rhs4_
                lhs2_[lhs3_] = rhs5_
            elif source7_.is_DC30:
                d_248___mcc_h9_ = source7_.cf61
                d_249___mcc_h10_ = source7_.cf62
                d_250___mcc_h11_ = source7_.cf63
                d_251___mcc_h12_ = source7_.cf64
                d_252_cf64_ = d_251___mcc_h12_
                d_253_cf63_ = d_250___mcc_h11_
                d_254_cf62_ = d_249___mcc_h10_
                d_255_cf61_ = d_248___mcc_h9_
                d_256_v41_: _dafny.Array
                nw14_ = _dafny.Array(None, 3)
                d_256_v41_ = nw14_
                d_257_v42_: _dafny.Map
                d_257_v42_ = _dafny.Map({d_255_cf61_: p1})
                d_258_v43_: T0
                nw15_ = C7()
                nw15_.ctor__(606, d_255_cf61_, d_255_cf61_, d_253_cf63_, ((d_257_v42_)[False] if (False) in (d_257_v42_) else p1), True, d_178_v0_, not(not(d_255_cf61_)))
                d_258_v43_ = nw15_
                index8_ = default__.safeIndex(518, (d_256_v41_).length(0))
                (d_256_v41_)[index8_] = d_258_v43_
                d_259_v44_: D1
                d_259_v44_ = D1_DC5(d_252_cf64_, d_253_cf63_, d_178_v0_)
                index9_ = default__.safeIndex(518, (d_256_v41_).length(0))
                nw16_ = C11()
                nw16_.ctor__(d_259_v44_, not(d_252_cf64_), (d_258_v43_).f5)
                (d_256_v41_)[index9_] = nw16_
                d_260_v45_: _dafny.Seq
                d_260_v45_ = _dafny.SeqWithoutIsStrInference([d_199_v19_, d_199_v19_])
                d_260_v45_ = d_260_v45_
                d_255_cf61_ = (d_255_cf61_) in (d_257_v42_)
                d_253_cf63_ = (d_253_cf63_) + (d_253_cf63_)
            elif source7_.is_DC27:
                d_261___mcc_h13_ = source7_.cf51
                d_262_cf51_ = d_261___mcc_h13_
                d_197_v17_ = ((d_197_v17_) | (d_197_v17_)) | ((d_197_v17_ if p0 else d_197_v17_))
                d_263_v46_: _dafny.Map
                d_263_v46_ = _dafny.Map({p0: p1})
                d_264_v47_: _dafny.Seq
                d_264_v47_ = _dafny.SeqWithoutIsStrInference([d_263_v46_])
                index10_ = default__.safeIndex(643, (d_199_v19_).length(0))
                (d_199_v19_)[index10_] = (d_263_v46_) not in (d_264_v47_)
                index11_ = default__.safeIndex(643, (d_199_v19_).length(0))
                (d_199_v19_)[index11_] = not(((p1 if p0 else p1)) == (-461))
                r0 = (not((d_199_v19_)[default__.safeIndex(643, (d_199_v19_).length(0))])) and (p0)
                d_265_v48_: int
                d_265_v48_ = 242
                d_265_v48_ = p1
            elif True:
                d_266___mcc_h14_ = source7_.cf65
                d_267_cf65_ = d_266___mcc_h14_
                d_268_v49_: int
                d_268_v49_ = 18
                d_268_v49_ = p1
                r0 = p0
                d_268_v49_ = (len(d_178_v0_)) + (d_268_v49_)
                r0 = default__.fm27(d_179_v1_, globalState)
            d_269_v50_: _dafny.Seq
            d_269_v50_ = _dafny.SeqWithoutIsStrInference([p0])
            d_270_v51_: D1
            d_270_v51_ = D1_DC3(d_269_v50_)
            d_271_v52_: _dafny.Seq
            d_271_v52_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_272_v53_: D19
            d_272_v53_ = D19_DC53(p1, p1, p0, p0)
            pat_let_tv3_ = d_269_v50_
            d_273_v54_: _dafny.Array
            nw17_ = _dafny.Array(None, 21)
            def iife46_(_pat_let6_0):
                def iife47_(d_274_dt__update__tmp_h3_):
                    def iife48_(_pat_let7_0):
                        def iife49_(d_275_dt__update_hcf11_h0_):
                            return D1_DC3(d_275_dt__update_hcf11_h0_)
                        return iife49_(_pat_let7_0)
                    return iife48_(pat_let_tv3_)
                return iife47_(_pat_let6_0)
            nw17_[int(0)] = (iife46_(d_270_v51_)).cf11
            nw17_[int(1)] = (d_269_v50_) + (d_269_v50_)
            nw17_[int(2)] = d_269_v50_
            nw17_[int(3)] = d_269_v50_
            nw17_[int(4)] = d_269_v50_
            nw17_[int(5)] = (d_269_v50_) + (d_269_v50_)
            nw17_[int(6)] = (d_269_v50_) + (d_269_v50_)
            nw17_[int(7)] = d_269_v50_
            nw17_[int(8)] = (d_269_v50_) + ((d_269_v50_).set(default__.safeIndex(len(d_271_v52_), len(d_269_v50_)), p0))
            nw17_[int(9)] = d_269_v50_
            nw17_[int(10)] = d_269_v50_
            nw17_[int(11)] = d_269_v50_
            nw17_[int(12)] = d_269_v50_
            nw17_[int(13)] = _dafny.SeqWithoutIsStrInference([default__.fm27(d_179_v1_, globalState), p0, False])
            nw17_[int(14)] = (_dafny.SeqWithoutIsStrInference([p0])) + ((d_270_v51_).cf11)
            nw17_[int(15)] = d_269_v50_
            nw17_[int(16)] = d_269_v50_
            nw17_[int(17)] = d_269_v50_
            nw17_[int(18)] = d_269_v50_
            nw17_[int(19)] = default__.fm45(p1, globalState)
            nw17_[int(20)] = ((_dafny.SeqWithoutIsStrInference([p0, p0])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([p0, p0]))), (d_272_v53_).cf96)) + (d_269_v50_)
            d_273_v54_ = nw17_
            d_273_v54_ = d_273_v54_
            d_276_v55_: _dafny.Array
            nw18_ = _dafny.Array(_dafny.Map({}), 11)
            d_276_v55_ = nw18_
            d_276_v55_ = d_276_v55_
            d_277_v56_: _dafny.Set
            d_277_v56_ = _dafny.Set({p1, p1})
            d_278_v57_: _dafny.Seq
            d_278_v57_ = _dafny.SeqWithoutIsStrInference([d_277_v56_, d_277_v56_])
            d_279_v58_: C13
            nw19_ = C13()
            nw19_.ctor__(p0, (_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kuf")))}) for d_280_i4_ in range(default__.abs(316))])) < (d_278_v57_), (d_271_v52_) != (d_271_v52_), p1)
            d_279_v58_ = nw19_
            d_281_v59_: _dafny.Seq
            d_281_v59_ = _dafny.SeqWithoutIsStrInference([d_279_v58_])
            d_279_v58_ = (d_281_v59_)[default__.safeIndex((0) - ((0) - ((p1) - ((0) - (p1)))), len(d_281_v59_))]
            d_282_v60_: _dafny.Array
            nw20_ = _dafny.Array(int(0), 28)
            d_282_v60_ = nw20_
            index12_ = default__.safeIndex(148, (d_282_v60_).length(0))
            (d_282_v60_)[index12_] = (p1 if d_279_v58_.f7 else p1)
            index13_ = default__.safeIndex(148, (d_282_v60_).length(0))
            (d_282_v60_)[index13_] = default__.fm18(d_279_v58_.f6, d_279_v58_.f7, True, globalState)
        d_283_v61_: _dafny.Set
        d_283_v61_ = _dafny.Set({d_179_v1_, d_179_v1_})
        if ((d_283_v61_).issubset(d_283_v61_)) or ((p1) < (p1)):
            r0 = p0
            d_284_v62_: _dafny.Set
            d_284_v62_ = _dafny.Set({p0})
            d_285_v63_: _dafny.Array
            nw21_ = _dafny.Array(None, 19)
            nw21_[int(0)] = p0
            nw21_[int(1)] = (p1) > (p1)
            nw21_[int(2)] = p0
            nw21_[int(3)] = False
            nw21_[int(4)] = not(p0)
            nw21_[int(5)] = p0
            nw21_[int(6)] = p0
            nw21_[int(7)] = (d_284_v62_) == (d_284_v62_)
            nw21_[int(8)] = not(p0)
            nw21_[int(9)] = p0
            nw21_[int(10)] = p0
            nw21_[int(11)] = not(p0)
            nw21_[int(12)] = p0
            nw21_[int(13)] = (p0) or (True)
            nw21_[int(14)] = p0
            nw21_[int(15)] = p0
            nw21_[int(16)] = p0
            nw21_[int(17)] = p0
            nw21_[int(18)] = not(p0)
            d_285_v63_ = nw21_
            index14_ = default__.safeIndex(335, (d_285_v63_).length(0))
            (d_285_v63_)[index14_] = p0
            d_286_v64_: C2
            nw22_ = C2()
            nw22_.ctor__(d_285_v63_, default__.fm68(globalState), d_178_v0_, not(True), p0, p1)
            d_286_v64_ = nw22_
            d_287_v65_: _dafny.Seq
            d_287_v65_ = _dafny.SeqWithoutIsStrInference([d_286_v64_, d_286_v64_, d_286_v64_])
            d_288_v66_: _dafny.Seq
            d_288_v66_ = d_287_v65_
            d_289_v67_: _dafny.MultiSet
            d_289_v67_ = _dafny.MultiSet([(d_288_v66_)])
            index15_ = default__.safeIndex(335, (d_285_v63_).length(0))
            (d_285_v63_)[index15_] = (d_289_v67_).ispropersubset(d_289_v67_)
            d_290_v68_: int
            d_290_v68_ = 665
            d_290_v68_ = -774
            d_290_v68_ = 79
            d_291_v69_: _dafny.MultiSet
            d_291_v69_ = _dafny.MultiSet([(d_286_v64_).f28, (d_286_v64_).f28])
            d_291_v69_ = (d_291_v69_) | (d_291_v69_)
        elif True:
            d_292_v70_: _dafny.Map
            d_292_v70_ = _dafny.Map({d_178_v0_: p0})
            r0 = (not(((d_292_v70_)[d_178_v0_] if (d_178_v0_) in (d_292_v70_) else p0))) == (p0)
            d_178_v0_ = (d_178_v0_) + ((_dafny.SeqWithoutIsStrInference([d_179_v1_ for d_293_i5_ in range(default__.abs(19))])) + (d_178_v0_))
            d_294_v71_: _dafny.Seq
            d_294_v71_ = _dafny.SeqWithoutIsStrInference([False])
            d_295_v72_: _dafny.Map
            d_295_v72_ = _dafny.Map({p1: default__.fm55(p1, d_294_v71_, 674, globalState)})
            d_295_v72_ = (d_295_v72_).set((0) - (p1), p1)
            d_295_v72_ = (d_295_v72_ if not(p0) else _dafny.Map({p1: p1}))
            d_296_v73_: T1
            nw23_ = C7()
            nw23_.ctor__(929, p0, p0, p1, p1, default__.fm27(d_179_v1_, globalState), d_178_v0_, not(p0))
            d_296_v73_ = nw23_
            d_297_v74_: _dafny.Seq
            d_297_v74_ = _dafny.SeqWithoutIsStrInference([(d_296_v73_).f15])
            d_298_v75_: D1
            d_298_v75_ = D1_DC5(p0, len(d_297_v74_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvdop")))
            pat_let_tv4_ = d_178_v0_
            d_299_v76_: C6
            nw24_ = C6()
            def iife50_(_pat_let8_0):
                def iife51_(d_300_dt__update__tmp_h4_):
                    def iife52_(_pat_let9_0):
                        def iife53_(d_301_dt__update_hcf17_h0_):
                            return D1_DC5((d_300_dt__update__tmp_h4_).cf15, (d_300_dt__update__tmp_h4_).cf16, d_301_dt__update_hcf17_h0_)
                        return iife53_(_pat_let9_0)
                    return iife52_(pat_let_tv4_)
                return iife51_(_pat_let8_0)
            nw24_.ctor__((d_296_v73_).f5, (iife50_(d_298_v75_)).cf15, (d_296_v73_).f16, -305, (d_296_v73_).f15, not((d_296_v73_).f4))
            d_299_v76_ = nw24_
            d_302_v77_: _dafny.Map
            d_302_v77_ = _dafny.Map({p1: d_299_v76_})
            d_303_v78_: _dafny.Map
            d_303_v78_ = _dafny.Map({d_296_v73_: (d_294_v71_).set(default__.safeIndex(p1, len(d_294_v71_)), (d_294_v71_)[default__.safeIndex(len(d_302_v77_), len(d_294_v71_))])})
            d_303_v78_ = d_303_v78_
        d_304_v79_: _dafny.MultiSet
        d_304_v79_ = _dafny.MultiSet([p1])
        d_305_v80_: D6
        d_305_v80_ = D6_DC15((default__.fm44(p1, D0_DC0(p0, p1, p0, len(_dafny.SeqWithoutIsStrInference([p1, (d_304_v79_).cardinality])), p0), p0, p0, globalState)).set(default__.safeIndex(default__.fm3(globalState), len(default__.fm44(p1, D0_DC0(p0, p1, p0, len(_dafny.SeqWithoutIsStrInference([p1, (d_304_v79_).cardinality])), p0), p0, p0, globalState))), len(d_178_v0_)))
        d_306_v81_: _dafny.Map
        d_306_v81_ = _dafny.Map({p1: (0) - (p1)})
        d_307_v82_: _dafny.Seq
        d_307_v82_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_308_v83_: _dafny.MultiSet
        d_308_v83_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufxwitu"))])
        d_309_v84_: _dafny.MultiSet
        d_309_v84_ = _dafny.MultiSet([True, p0, p0])
        d_310_v85_: _dafny.Seq
        d_310_v85_ = _dafny.SeqWithoutIsStrInference([((d_308_v83_)[d_178_v0_] if (d_178_v0_) in (d_308_v83_) else (d_309_v84_).cardinality)])
        d_311_v86_: _dafny.Seq
        d_311_v86_ = _dafny.SeqWithoutIsStrInference([default__.fm25(p1, (0) - (((d_306_v81_)[len(d_307_v82_)] if (len(d_307_v82_)) in (d_306_v81_) else len(d_310_v85_))), p0, globalState)])
        pat_let_tv5_ = d_310_v85_
        def iife54_(_pat_let10_0):
            def iife55_(d_312_dt__update__tmp_h5_):
                def iife56_(_pat_let11_0):
                    def iife57_(d_313_dt__update_hcf36_h0_):
                        return D6_DC15(d_313_dt__update_hcf36_h0_)
                    return iife57_(_pat_let11_0)
                return iife56_(pat_let_tv5_)
            return iife55_(_pat_let10_0)
        source8_ = iife54_(d_305_v80_)
        if source8_.is_DC16:
            d_314___mcc_h15_ = source8_.cf37
            d_315___mcc_h16_ = source8_.cf38
            d_316_cf38_ = d_315___mcc_h16_
            d_317_cf37_ = d_314___mcc_h15_
            d_318_v87_: _dafny.Array
            nw25_ = _dafny.Array(_dafny.Seq({}), 17)
            d_318_v87_ = nw25_
            d_319_v88_: D26
            d_319_v88_ = D26_DC66(d_318_v87_)
            d_318_v87_ = (d_319_v88_).cf117
            d_320_v89_: _dafny.Array
            nw26_ = _dafny.Array(False, 23)
            d_320_v89_ = nw26_
            d_321_v90_: C2
            nw27_ = C2()
            nw27_.ctor__(d_320_v89_, (d_304_v79_).set(p1, default__.abs(d_316_cf38_)), d_178_v0_, p0, default__.fm27(d_179_v1_, globalState), d_316_cf38_)
            d_321_v90_ = nw27_
            d_322_v91_: _dafny.Map
            d_322_v91_ = _dafny.Map({d_321_v90_: p0})
            d_323_v92_: C4
            nw28_ = C4()
            nw28_.ctor__(931, p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "se")), p0, p0, p1, p1, p0)
            d_323_v92_ = nw28_
            d_324_v93_: T2
            nw29_ = C7()
            nw29_.ctor__(default__.safeDivisionInt(d_316_cf38_, len(d_322_v91_)), p0, True, ((d_309_v84_).cardinality) + (len(_dafny.Map({d_323_v92_: not(True)}))), d_316_cf38_, not (p0) or (p0), (d_178_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xakt"))), (d_323_v92_).f24)
            d_324_v93_ = nw29_
            d_324_v93_ = d_324_v93_
            d_325_v94_: D22
            d_325_v94_ = D22_DC58(d_178_v0_, (d_324_v93_).f5, (d_324_v93_).f4)
            d_326_v95_: C6
            nw30_ = C6()
            nw30_.ctor__((d_324_v93_).f5, (d_325_v94_).cf104, p0, -394, ((_dafny.MultiSet([d_324_v93_.f18]) if p0 else d_309_v84_)).cardinality, p0)
            d_326_v95_ = nw30_
            d_327_v96_: _dafny.Seq
            d_327_v96_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([((d_324_v93_).f17)[default__.safeIndex(d_326_v95_.f21, len((d_324_v93_).f17))] for d_328_i6_ in range(default__.abs(309))])) + ((d_324_v93_).f17)])
            rhs6_ = d_326_v95_
            rhs7_ = len(d_327_v96_)
            lhs4_ = d_326_v95_
            d_326_v95_ = rhs6_
            lhs4_.f21 = rhs7_
            d_329_v97_: _dafny.Set
            d_329_v97_ = _dafny.Set({(d_326_v95_.f21) - ((d_324_v93_).f5)})
            d_330_v98_: _dafny.Array
            def lambda2_(d_331_v92_):
                def lambda3_(d_332_i7_):
                    return (d_332_i7_) * ((d_331_v92_).f23)

                return lambda3_

            init1_ = lambda2_(d_323_v92_)
            nw31_ = _dafny.Array(None, 29)
            for i0_1_ in range(nw31_.length(0)):
                nw31_[i0_1_] = init1_(i0_1_)
            d_330_v98_ = nw31_
            d_333_v99_: _dafny.Map
            d_333_v99_ = _dafny.Map({True: p1})
            d_334_v100_: _dafny.Map
            d_334_v100_ = d_333_v99_
            d_335_v101_: _dafny.Map
            d_335_v101_ = _dafny.Map({d_178_v0_: (d_324_v93_).f5})
            d_336_v102_: _dafny.Map
            d_336_v102_ = _dafny.Map({(0) - (((d_335_v101_)[d_178_v0_] if (d_178_v0_) in (d_335_v101_) else d_316_cf38_)): d_330_v98_})
            rhs8_ = (d_304_v79_).ispropersubset(d_304_v79_)
            rhs9_ = d_329_v97_
            rhs10_ = ((d_336_v102_)[p1] if (p1) in (d_336_v102_) else d_330_v98_)
            rhs11_ = d_334_v100_
            rhs12_ = ((d_323_v92_).f23) != (((d_306_v81_)[d_326_v95_.f21] if (d_326_v95_.f21) in (d_306_v81_) else d_326_v95_.f21))
            lhs5_ = d_324_v93_
            lhs5_.f18 = rhs8_
            d_329_v97_ = rhs9_
            d_330_v98_ = rhs10_
            d_334_v100_ = rhs11_
            r0 = rhs12_
        elif source8_.is_DC15:
            d_337___mcc_h17_ = source8_.cf36
            d_338_cf36_ = d_337___mcc_h17_
            d_339_v103_: C5
            nw32_ = C5()
            nw32_.ctor__()
            d_339_v103_ = nw32_
            d_340_v104_: int
            d_340_v104_ = 143
            d_340_v104_ = d_340_v104_
            d_341_v105_: _dafny.Seq
            d_341_v105_ = _dafny.SeqWithoutIsStrInference([d_178_v0_])
            rhs13_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))) if False else (d_340_v104_) - (len(d_178_v0_)))
            rhs14_ = d_179_v1_
            rhs15_ = ((0) - (d_340_v104_)) + (d_340_v104_)
            rhs16_ = ((d_341_v105_)[default__.safeIndex(p1, len(d_341_v105_))] if p0 else d_178_v0_)
            d_340_v104_ = rhs13_
            d_179_v1_ = rhs14_
            d_340_v104_ = rhs15_
            d_178_v0_ = rhs16_
            r0 = p0
        elif True:
            d_342___mcc_h18_ = source8_.cf39
            d_343_cf39_ = d_342___mcc_h18_
            d_344_v106_: C5
            nw33_ = C5()
            nw33_.ctor__()
            d_344_v106_ = nw33_
            if p0:
                d_345_v107_: int
                d_345_v107_ = -428
                d_346_v108_: _dafny.Set
                d_346_v108_ = _dafny.Set({d_306_v81_})
                d_347_v109_: _dafny.Set
                d_347_v109_ = d_346_v108_
                d_348_v110_: C0
                nw34_ = C0()
                nw34_.ctor__(d_347_v109_, p1)
                d_348_v110_ = nw34_
                d_349_v111_: T2
                nw35_ = C4()
                nw35_.ctor__(p1, False, d_178_v0_, p0, p0, p1, d_345_v107_, p0)
                d_349_v111_ = nw35_
                d_350_v112_: D18
                d_350_v112_ = D18_DC49(d_349_v111_)
                d_351_v113_: D24
                d_351_v113_ = D24_DC64(False, d_348_v110_, d_350_v112_)
                d_352_v114_: _dafny.Map
                d_352_v114_ = _dafny.Map({(d_351_v113_).cf114: d_349_v111_.f18})
                d_345_v107_ = (len(d_352_v114_)) * ((0) - ((d_349_v111_).f5))
                d_345_v107_ = (d_349_v111_).f5
                d_353_v115_: _dafny.Map
                d_353_v115_ = _dafny.Map({(d_349_v111_).f5: False})
                d_354_v116_: _dafny.Set
                d_354_v116_ = _dafny.Set({((d_353_v115_)[(d_348_v110_).f27] if ((d_348_v110_).f27) in (d_353_v115_) else p0), d_349_v111_.f18})
                d_355_v117_: _dafny.Array
                nw36_ = _dafny.Array(None, 26)
                nw36_[int(0)] = (d_354_v116_) - (d_354_v116_)
                nw36_[int(1)] = d_354_v116_
                nw36_[int(2)] = d_354_v116_
                nw36_[int(3)] = d_354_v116_
                nw36_[int(4)] = (d_354_v116_).intersection(d_354_v116_)
                nw36_[int(5)] = (d_354_v116_).intersection(d_354_v116_)
                nw36_[int(6)] = d_354_v116_
                nw36_[int(7)] = (default__.fm48(911, globalState)).intersection(d_354_v116_)
                nw36_[int(8)] = d_354_v116_
                nw36_[int(9)] = d_354_v116_
                nw36_[int(10)] = (d_354_v116_).intersection(d_354_v116_)
                nw36_[int(11)] = (d_354_v116_) - (_dafny.Set({d_349_v111_.f18}))
                nw36_[int(12)] = d_354_v116_
                nw36_[int(13)] = d_354_v116_
                nw36_[int(14)] = _dafny.Set({(d_349_v111_).f4, d_349_v111_.f18})
                nw36_[int(15)] = (d_354_v116_) - (_dafny.Set({p0}))
                nw36_[int(16)] = (d_354_v116_).intersection(d_354_v116_)
                nw36_[int(17)] = d_354_v116_
                nw36_[int(18)] = d_354_v116_
                nw36_[int(19)] = d_354_v116_
                nw36_[int(20)] = d_354_v116_
                nw36_[int(21)] = _dafny.Set({(d_349_v111_).f4})
                nw36_[int(22)] = d_354_v116_
                nw36_[int(23)] = d_354_v116_
                nw36_[int(24)] = d_354_v116_
                nw36_[int(25)] = (d_354_v116_) - (d_354_v116_)
                d_355_v117_ = nw36_
                index16_ = default__.safeIndex(558, (d_355_v117_).length(0))
                (d_355_v117_)[index16_] = d_354_v116_
                d_356_v118_: _dafny.Map
                d_356_v118_ = _dafny.Map({len(d_353_v115_): d_354_v116_})
                index17_ = default__.safeIndex(558, (d_355_v117_).length(0))
                (d_355_v117_)[index17_] = (default__.fm48(d_345_v107_, globalState)) - ((d_354_v116_) | (((d_356_v118_)[d_345_v107_] if (d_345_v107_) in (d_356_v118_) else d_354_v116_)))
                d_357_v119_: C5
                nw37_ = C5()
                nw37_.ctor__()
                d_357_v119_ = nw37_
                (d_349_v111_).f18 = (d_349_v111_).f4
            elif True:
                d_358_v120_: _dafny.Array
                nw38_ = _dafny.Array(None, 7)
                nw38_[int(0)] = d_178_v0_
                nw38_[int(1)] = d_178_v0_
                nw38_[int(2)] = (d_178_v0_) + (d_178_v0_)
                nw38_[int(3)] = d_178_v0_
                nw38_[int(4)] = _dafny.SeqWithoutIsStrInference([d_179_v1_ for d_359_i8_ in range(default__.abs(117))])
                nw38_[int(5)] = (d_178_v0_) + (d_178_v0_)
                nw38_[int(6)] = _dafny.SeqWithoutIsStrInference([(d_178_v0_)[default__.safeIndex(p1, len(d_178_v0_))] for d_360_i9_ in range(default__.abs(924))])
                d_358_v120_ = nw38_
                d_358_v120_ = d_358_v120_
                d_361_v121_: _dafny.Array
                def lambda4_(d_362_p1_):
                    def lambda5_(d_363_i10_):
                        return (d_363_i10_) - (d_362_p1_)

                    return lambda5_

                init2_ = lambda4_(p1)
                nw39_ = _dafny.Array(None, 3)
                for i0_2_ in range(nw39_.length(0)):
                    nw39_[i0_2_] = init2_(i0_2_)
                d_361_v121_ = nw39_
                d_364_v122_: _dafny.Map
                d_364_v122_ = _dafny.Map({p1: d_361_v121_})
                d_364_v122_ = (d_364_v122_).set(460, d_361_v121_)
                d_365_v123_: C12
                nw40_ = C12()
                nw40_.ctor__(p1)
                d_365_v123_ = nw40_
                d_366_v124_: _dafny.Map
                d_366_v124_ = _dafny.Map({(d_365_v123_).f8: p0})
                d_367_v125_: C6
                nw41_ = C6()
                nw41_.ctor__((d_365_v123_).f8, ((d_365_v123_).f8) not in (d_366_v124_), p0, (d_365_v123_).f8, (0) - (p1), p0)
                d_367_v125_ = nw41_
                d_368_v126_: C3
                nw42_ = C3()
                nw42_.ctor__(d_179_v1_, -610, True, p0, (d_365_v123_).f8)
                d_368_v126_ = nw42_
                d_369_v127_: _dafny.Map
                d_369_v127_ = _dafny.Map({p0: d_368_v126_})
                d_370_v128_: T0
                nw43_ = C7()
                nw43_.ctor__((d_365_v123_).f8, (d_367_v125_).f22, (d_367_v125_).f22, (d_365_v123_).f8, len(d_369_v127_), p0, d_178_v0_, not(True))
                d_370_v128_ = nw43_
                d_371_v129_: _dafny.Array
                nw44_ = _dafny.Array(None, 5)
                nw44_[int(0)] = d_370_v128_
                nw44_[int(1)] = d_370_v128_
                nw44_[int(2)] = d_370_v128_
                nw44_[int(3)] = d_370_v128_
                nw44_[int(4)] = (d_370_v128_ if (d_370_v128_).f4 else d_370_v128_)
                d_371_v129_ = nw44_
                d_371_v129_ = d_371_v129_
            r0 = p0
            def iife58_():
                coll34_ = _dafny.Set()
                compr_34_: int
                for compr_34_ in _dafny.IntegerRange(-522, 959):
                    d_372_v130_: int = compr_34_
                    if ((-522) <= (d_372_v130_)) and ((d_372_v130_) < (959)):
                        coll34_ = coll34_.union(_dafny.Set([(d_372_v130_) * (len(d_306_v81_))]))
                return _dafny.Set(coll34_)
            def iife59_():
                coll35_ = _dafny.Set()
                compr_35_: int
                for compr_35_ in _dafny.IntegerRange(-42, 830):
                    d_373_v131_: int = compr_35_
                    if ((-42) <= (d_373_v131_)) and ((d_373_v131_) < (830)):
                        coll35_ = coll35_.union(_dafny.Set([(d_373_v131_) * (p1)]))
                return _dafny.Set(coll35_)
            r0 = not((iife58_()
            ).ispropersubset(iife59_()
            ))
        d_374_v132_: _dafny.Array
        nw45_ = _dafny.Array(int(0), 15)
        d_374_v132_ = nw45_
        index18_ = default__.safeIndex(524, (d_374_v132_).length(0))
        (d_374_v132_)[index18_] = (0) - (p1)
        index19_ = default__.safeIndex(524, (d_374_v132_).length(0))
        (d_374_v132_)[index19_] = len(_dafny.Set({(d_374_v132_ if not(p0) else d_374_v132_), d_374_v132_, d_374_v132_}))
        d_375_v133_: C8
        nw46_ = C8()
        nw46_.ctor__(d_178_v0_, p1)
        d_375_v133_ = nw46_
        index20_ = default__.safeIndex(524, (d_374_v132_).length(0))
        (d_374_v132_)[index20_] = (d_374_v132_)[default__.safeIndex(524, (d_374_v132_).length(0))]
        r0 = ((d_374_v132_)[default__.safeIndex(524, (d_374_v132_).length(0))]) != ((d_374_v132_)[default__.safeIndex(524, (d_374_v132_).length(0))])
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_376_v0_: bool
        d_376_v0_ = False
        d_377_v1_: _dafny.Set
        d_377_v1_ = _dafny.Set({d_376_v0_, d_376_v0_, d_376_v0_, d_376_v0_, d_376_v0_})
        d_378_globalState_: GlobalState
        nw47_ = GlobalState()
        nw47_.ctor__(d_377_v1_, True)
        d_378_globalState_ = nw47_
        d_379_v2_: int
        d_379_v2_ = -871
        pat_let_tv6_ = d_376_v0_
        pat_let_tv7_ = d_376_v0_
        def iife60_(_pat_let12_0):
            def iife61_(d_380_dt__update__tmp_h0_):
                def iife62_(_pat_let13_0):
                    def iife63_(d_381_dt__update_hcf0_h0_):
                        return D0_DC0(d_381_dt__update_hcf0_h0_, (d_380_dt__update__tmp_h0_).cf1, (d_380_dt__update__tmp_h0_).cf2, (d_380_dt__update__tmp_h0_).cf3, (d_380_dt__update__tmp_h0_).cf4)
                    return iife63_(_pat_let13_0)
                return iife62_(not (pat_let_tv6_) or (pat_let_tv7_))
            return iife61_(_pat_let12_0)
        source9_ = iife60_(D0_DC0(d_376_v0_, d_379_v2_, d_376_v0_, d_379_v2_, d_376_v0_))
        if source9_.is_DC0:
            d_382___mcc_h0_ = source9_.cf0
            d_383___mcc_h1_ = source9_.cf1
            d_384___mcc_h2_ = source9_.cf2
            d_385___mcc_h3_ = source9_.cf3
            d_386___mcc_h4_ = source9_.cf4
            d_387_cf4_ = d_386___mcc_h4_
            d_388_cf3_ = d_385___mcc_h3_
            d_389_cf2_ = d_384___mcc_h2_
            d_390_cf1_ = d_383___mcc_h1_
            d_391_cf0_ = d_382___mcc_h0_
            d_392_v3_: D1
            d_392_v3_ = D1_DC3(_dafny.SeqWithoutIsStrInference([d_391_cf0_]))
            d_393_v4_: _dafny.Seq
            d_393_v4_ = _dafny.SeqWithoutIsStrInference([not(d_391_cf0_), d_389_cf2_])
            if ((d_392_v3_).cf11) < ((d_393_v4_).set(default__.safeIndex(d_379_v2_, len(d_393_v4_)), d_391_cf0_)):
                d_388_cf3_ = d_390_cf1_
                d_391_cf0_ = (d_377_v1_).issubset(d_377_v1_)
                d_393_v4_ = d_393_v4_
                d_394_v5_: _dafny.Seq
                d_394_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awdpjwlt"))
                d_395_v6_: bool
                out0_: bool
                out0_ = default__.m0((d_394_v5_) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_396_i0_ in range(default__.abs(-482))])), d_388_cf3_, d_378_globalState_)
                d_395_v6_ = out0_
                d_397_v7_: _dafny.Array
                nw48_ = _dafny.Array(int(0), 23)
                d_397_v7_ = nw48_
                d_398_v8_: _dafny.Seq
                d_398_v8_ = _dafny.SeqWithoutIsStrInference([d_390_cf1_, d_388_cf3_])
                index21_ = default__.safeIndex(26, (d_397_v7_).length(0))
                (d_397_v7_)[index21_] = (0) - (len((d_398_v8_) + (d_398_v8_)))
                d_399_v9_: _dafny.MultiSet
                d_399_v9_ = _dafny.MultiSet([(0) - (d_379_v2_), d_390_cf1_])
                d_400_v10_: D0
                d_400_v10_ = D0_DC0(d_387_cf4_, d_390_cf1_, d_391_cf0_, ((d_399_v9_)[d_379_v2_] if (d_379_v2_) in (d_399_v9_) else 879), d_391_cf0_)
                index22_ = default__.safeIndex(26, (d_397_v7_).length(0))
                (d_397_v7_)[index22_] = (d_400_v10_).cf3
            elif True:
                d_401_v11_: _dafny.Array
                nw49_ = _dafny.Array(None, 20)
                nw49_[int(0)] = not(d_376_v0_)
                nw49_[int(1)] = d_389_cf2_
                nw49_[int(2)] = d_389_cf2_
                nw49_[int(3)] = d_389_cf2_
                nw49_[int(4)] = False
                nw49_[int(5)] = d_389_cf2_
                nw49_[int(6)] = d_387_cf4_
                nw49_[int(7)] = d_391_cf0_
                nw49_[int(8)] = d_389_cf2_
                nw49_[int(9)] = d_389_cf2_
                nw49_[int(10)] = True
                nw49_[int(11)] = d_391_cf0_
                nw49_[int(12)] = d_391_cf0_
                nw49_[int(13)] = d_389_cf2_
                nw49_[int(14)] = d_389_cf2_
                nw49_[int(15)] = d_389_cf2_
                nw49_[int(16)] = d_389_cf2_
                nw49_[int(17)] = d_389_cf2_
                nw49_[int(18)] = d_391_cf0_
                nw49_[int(19)] = not(d_387_cf4_)
                d_401_v11_ = nw49_
                d_402_v12_: C9
                nw50_ = C9()
                nw50_.ctor__(d_401_v11_)
                d_402_v12_ = nw50_
                d_403_v13_: _dafny.Seq
                d_403_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
                d_404_v14_: _dafny.Seq
                d_404_v14_ = _dafny.SeqWithoutIsStrInference([len(d_403_v13_)])
                d_387_cf4_ = ((d_404_v14_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_379_v2_ for d_405_i1_ in range(default__.abs(922))])), len(d_404_v14_))]) < (d_379_v2_)
                d_406_v15_: _dafny.Array
                nw51_ = _dafny.Array(int(0), 15)
                d_406_v15_ = nw51_
                index23_ = default__.safeIndex(743, (d_406_v15_).length(0))
                (d_406_v15_)[index23_] = d_379_v2_
                index24_ = default__.safeIndex(743, (d_406_v15_).length(0))
                (d_406_v15_)[index24_] = default__.safeModuloInt((d_390_cf1_) * (len(d_403_v13_)), d_388_cf3_)
                d_391_cf0_ = (d_393_v4_)[default__.safeIndex((d_406_v15_)[default__.safeIndex(743, (d_406_v15_).length(0))], len(d_393_v4_))]
                d_388_cf3_ = d_388_cf3_
            d_388_cf3_ = d_388_cf3_
            d_407_v16_: _dafny.Map
            d_407_v16_ = _dafny.Map({d_389_cf2_: d_388_cf3_})
            d_408_v17_: D4
            d_408_v17_ = D4_DC9(d_407_v16_, d_379_v2_, False)
            d_379_v2_ = (((d_408_v17_).cf22) * ((0) - (d_379_v2_))) * (len(d_407_v16_))
            d_379_v2_ = d_390_cf1_
        elif source9_.is_DC1:
            d_409___mcc_h5_ = source9_.cf5
            d_410___mcc_h6_ = source9_.cf6
            d_411___mcc_h7_ = source9_.cf7
            d_412___mcc_h8_ = source9_.cf8
            d_413_cf8_ = d_412___mcc_h8_
            d_414_cf7_ = d_411___mcc_h7_
            d_415_cf6_ = d_410___mcc_h6_
            d_416_cf5_ = d_409___mcc_h5_
            d_417_v18_: _dafny.Seq
            d_417_v18_ = _dafny.SeqWithoutIsStrInference([False, True])
            d_418_v19_: _dafny.Seq
            d_418_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mb"))
            d_419_v20_: str
            d_419_v20_ = _dafny.CodePoint('q')
            pat_let_tv8_ = d_413_cf8_
            def iife64_(_pat_let14_0):
                def iife65_(d_421_dt__update__tmp_h1_):
                    def iife66_(_pat_let15_0):
                        def iife67_(d_422_dt__update_hcf85_h0_):
                            return D17_DC47((d_421_dt__update__tmp_h1_).cf83, (d_421_dt__update__tmp_h1_).cf84, d_422_dt__update_hcf85_h0_, (d_421_dt__update__tmp_h1_).cf86, (d_421_dt__update__tmp_h1_).cf87)
                        return iife67_(_pat_let15_0)
                    return iife66_(pat_let_tv8_)
                return iife65_(_pat_let14_0)
            source10_ = iife64_(D17_DC47(d_415_cf6_, (d_417_v18_)[default__.safeIndex(default__.fm25(d_414_cf7_, d_416_cf5_, d_415_cf6_, d_378_globalState_), len(d_417_v18_))], d_413_cf8_, len((d_418_v19_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_414_cf7_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_420_i2_ in range(default__.abs(502))]))])), len(d_418_v19_)), d_419_v20_)), d_416_cf5_))
            if source10_.is_DC45:
                d_423_v21_: _dafny.Array
                def lambda6_(d_424_v0_, d_425_cf6_):
                    def lambda7_(d_426_i3_):
                        return _dafny.Map({d_424_v0_: not(d_425_cf6_)})

                    return lambda7_

                init3_ = lambda6_(d_376_v0_, d_415_cf6_)
                nw52_ = _dafny.Array(None, 27)
                for i0_3_ in range(nw52_.length(0)):
                    nw52_[i0_3_] = init3_(i0_3_)
                d_423_v21_ = nw52_
                d_427_v22_: _dafny.Map
                d_427_v22_ = _dafny.Map({d_418_v19_: d_423_v21_})
                d_427_v22_ = (d_427_v22_).set(d_418_v19_, d_423_v21_)
                d_428_v23_: _dafny.Array
                def lambda8_(d_429_cf6_):
                    def lambda9_(d_430_i4_):
                        return d_429_cf6_

                    return lambda9_

                init4_ = lambda8_(d_415_cf6_)
                nw53_ = _dafny.Array(None, 28)
                for i0_4_ in range(nw53_.length(0)):
                    nw53_[i0_4_] = init4_(i0_4_)
                d_428_v23_ = nw53_
                index25_ = default__.safeIndex(992, (d_428_v23_).length(0))
                (d_428_v23_)[index25_] = d_415_cf6_
                d_431_v24_: _dafny.Set
                d_431_v24_ = _dafny.Set({d_419_v20_, d_419_v20_})
                index26_ = default__.safeIndex(992, (d_428_v23_).length(0))
                (d_428_v23_)[index26_] = ((d_431_v24_) | (d_431_v24_)).issubset(d_431_v24_)
                d_432_v25_: _dafny.Array
                nw54_ = _dafny.Array(int(0), 26)
                d_432_v25_ = nw54_
                d_433_v26_: _dafny.Seq
                d_433_v26_ = _dafny.SeqWithoutIsStrInference([d_432_v25_, d_432_v25_])
                d_433_v26_ = d_433_v26_
                d_416_cf5_ = (d_413_cf8_) + (d_413_cf8_)
            elif source10_.is_DC46:
                d_434_v27_: bool
                out1_: bool
                out1_ = default__.m0(d_415_cf6_, default__.fm26(d_378_globalState_), d_378_globalState_)
                d_434_v27_ = out1_
                d_435_v28_: _dafny.MultiSet
                d_435_v28_ = _dafny.MultiSet([d_379_v2_])
                d_436_v29_: C13
                nw55_ = C13()
                nw55_.ctor__(d_434_v27_, d_376_v0_, d_376_v0_, d_413_cf8_)
                d_436_v29_ = nw55_
                d_437_v30_: _dafny.Map
                d_437_v30_ = _dafny.Map({d_436_v29_: d_413_cf8_})
                d_438_v31_: _dafny.Map
                d_438_v31_ = _dafny.Map({d_437_v30_: d_419_v20_})
                d_439_v32_: _dafny.Map
                d_439_v32_ = _dafny.Map({(d_435_v28_).cardinality: len(d_438_v31_)})
                d_439_v32_ = (d_439_v32_).set(d_414_cf7_, 178)
                d_440_v33_: _dafny.Map
                d_440_v33_ = _dafny.Map({d_419_v20_: d_414_cf7_})
                d_441_v34_: _dafny.MultiSet
                d_441_v34_ = _dafny.MultiSet([d_434_v27_, d_376_v0_])
                d_440_v33_ = (d_440_v33_).set(d_419_v20_, ((d_441_v34_)[d_415_cf6_] if (d_415_cf6_) in (d_441_v34_) else 445))
                d_442_v35_: _dafny.Array
                def lambda10_(d_443_v29_):
                    def lambda11_(d_444_i5_):
                        return not(d_443_v29_.f7)

                    return lambda11_

                init5_ = lambda10_(d_436_v29_)
                nw56_ = _dafny.Array(None, 10)
                for i0_5_ in range(nw56_.length(0)):
                    nw56_[i0_5_] = init5_(i0_5_)
                d_442_v35_ = nw56_
                d_442_v35_ = d_442_v35_
            elif source10_.is_DC47:
                d_445___mcc_h11_ = source10_.cf83
                d_446___mcc_h12_ = source10_.cf84
                d_447___mcc_h13_ = source10_.cf85
                d_448___mcc_h14_ = source10_.cf86
                d_449___mcc_h15_ = source10_.cf87
                d_450_cf87_ = d_449___mcc_h15_
                d_451_cf86_ = d_448___mcc_h14_
                d_452_cf85_ = d_447___mcc_h13_
                d_453_cf84_ = d_446___mcc_h12_
                d_454_cf83_ = d_445___mcc_h11_
                d_453_cf84_ = d_415_cf6_
                d_415_cf6_ = not(not (d_415_cf6_) or ((d_413_cf8_) not in (_dafny.Set({d_379_v2_, 329}))))
                d_455_v36_: _dafny.Map
                d_455_v36_ = _dafny.Map({d_416_cf5_: d_454_cf83_})
                d_455_v36_ = (d_455_v36_).set(d_414_cf7_, d_376_v0_)
                d_456_v37_: _dafny.Array
                def lambda12_(d_457_cf87_):
                    def lambda13_(d_458_i6_):
                        return (d_458_i6_) + (d_457_cf87_)

                    return lambda13_

                init6_ = lambda12_(d_450_cf87_)
                nw57_ = _dafny.Array(None, 17)
                for i0_6_ in range(nw57_.length(0)):
                    nw57_[i0_6_] = init6_(i0_6_)
                d_456_v37_ = nw57_
                index27_ = default__.safeIndex(38, (d_456_v37_).length(0))
                (d_456_v37_)[index27_] = d_413_cf8_
                index28_ = default__.safeIndex(38, (d_456_v37_).length(0))
                (d_456_v37_)[index28_] = d_451_cf86_
            elif source10_.is_DC44:
                d_459___mcc_h16_ = source10_.cf82
                d_460_cf82_ = d_459___mcc_h16_
                d_461_v38_: C8
                nw58_ = C8()
                nw58_.ctor__(d_418_v19_, (0) - (d_413_cf8_))
                d_461_v38_ = nw58_
                d_462_v39_: _dafny.Map
                d_462_v39_ = _dafny.Map({d_415_cf6_: (d_413_cf8_) - ((d_461_v38_).f14)})
                d_462_v39_ = (d_462_v39_).set(True, (d_413_cf8_) - (d_379_v2_))
                d_463_v40_: _dafny.Map
                d_463_v40_ = _dafny.Map({d_376_v0_: d_415_cf6_})
                d_464_v41_: _dafny.Map
                d_464_v41_ = _dafny.Map({d_463_v40_: _dafny.SeqWithoutIsStrInference([d_419_v20_ for d_465_i7_ in range(default__.abs(-667))])})
                d_466_v42_: _dafny.Map
                d_466_v42_ = _dafny.Map({d_464_v41_: (d_377_v1_).intersection(d_377_v1_)})
                d_466_v42_ = (d_466_v42_).set(d_464_v41_, d_377_v1_)
                d_467_v43_: _dafny.Set
                d_467_v43_ = _dafny.Set({(0) - (-29), d_414_cf7_})
                d_376_v0_ = ((d_467_v43_).intersection(d_467_v43_)).issubset(d_467_v43_)
            elif True:
                d_468___mcc_h17_ = source10_.cf88
                d_469_cf88_ = d_468___mcc_h17_
                d_470_v44_: D5
                d_470_v44_ = D5_DC13((d_417_v18_)[default__.safeIndex(d_413_cf8_, len(d_417_v18_))], d_376_v0_)
                d_470_v44_ = d_470_v44_
                d_471_v45_: _dafny.Array
                nw59_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_471_v45_ = nw59_
                d_472_v46_: _dafny.Map
                d_472_v46_ = _dafny.Map({d_471_v45_: len(d_417_v18_)})
                d_473_v47_: _dafny.Set
                d_473_v47_ = _dafny.Set({908, d_414_cf7_})
                d_474_v48_: bool
                out2_: bool
                out2_ = default__.m0((((d_472_v46_)[d_471_v45_] if (d_471_v45_) in (d_472_v46_) else d_413_cf8_)) >= (d_413_cf8_), default__.safeModuloInt(len(d_473_v47_), d_416_cf5_), d_378_globalState_)
                d_474_v48_ = out2_
                d_475_v49_: D1
                d_475_v49_ = D1_DC5(d_415_cf6_, d_379_v2_, d_418_v19_)
                d_476_v50_: C5
                nw60_ = C5()
                nw60_.ctor__()
                d_476_v50_ = nw60_
                d_477_v51_: _dafny.Set
                d_477_v51_ = _dafny.Set({d_476_v50_})
                d_478_v52_: T0
                nw61_ = C11()
                nw61_.ctor__(d_475_v49_, (d_476_v50_) in (d_477_v51_), d_416_cf5_)
                d_478_v52_ = nw61_
                d_478_v52_ = d_478_v52_
                d_479_v53_: C7
                nw62_ = C7()
                nw62_.ctor__(default__.safeDivisionInt((0) - ((0) - (d_416_cf5_)), 918), (d_417_v18_)[default__.safeIndex((d_478_v52_).f5, len(d_417_v18_))], (d_478_v52_).f4, 194, d_413_cf8_, d_376_v0_, (d_418_v19_) + (d_418_v19_), d_415_cf6_)
                d_479_v53_ = nw62_
            d_415_cf6_ = d_415_cf6_
            d_480_v54_: D21
            d_480_v54_ = D21_DC56(not(d_376_v0_))
            d_481_v55_: _dafny.Map
            d_481_v55_ = _dafny.Map({(d_480_v54_).cf100: d_416_cf5_})
            d_482_v56_: _dafny.Map
            d_482_v56_ = _dafny.Map({len(default__.fm48(d_414_cf7_, d_378_globalState_)): (0) - (d_413_cf8_)})
            d_483_v57_: _dafny.Seq
            d_483_v57_ = _dafny.SeqWithoutIsStrInference([d_481_v55_, d_481_v55_, d_481_v55_, (d_481_v55_) | (d_481_v55_), default__.fm57(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))): d_413_cf8_}), (0) - (((d_482_v56_)[d_416_cf5_] if (d_416_cf5_) in (d_482_v56_) else d_414_cf7_)), d_378_globalState_)])
            d_483_v57_ = d_483_v57_
            d_484_v58_: _dafny.Array
            def lambda14_(d_485_cf5_):
                def lambda15_(d_486_i8_):
                    return default__.safeDivisionInt(d_486_i8_, d_485_cf5_)

                return lambda15_

            init7_ = lambda14_(d_416_cf5_)
            nw63_ = _dafny.Array(None, 19)
            for i0_7_ in range(nw63_.length(0)):
                nw63_[i0_7_] = init7_(i0_7_)
            d_484_v58_ = nw63_
            d_487_v59_: _dafny.Map
            d_487_v59_ = _dafny.Map({d_484_v58_: d_416_cf5_})
            d_488_v60_: _dafny.Map
            d_488_v60_ = _dafny.Map({d_419_v20_: (d_413_cf8_) > ((0) - (d_416_cf5_))})
            d_489_v61_: _dafny.Array
            nw64_ = _dafny.Array(False, 19)
            d_489_v61_ = nw64_
            d_490_v62_: _dafny.Seq
            d_490_v62_ = _dafny.SeqWithoutIsStrInference([d_489_v61_])
            rhs17_ = d_487_v59_
            rhs18_ = d_418_v19_
            rhs19_ = d_488_v60_
            rhs20_ = d_490_v62_
            d_487_v59_ = rhs17_
            d_418_v19_ = rhs18_
            d_488_v60_ = rhs19_
            d_490_v62_ = rhs20_
        elif True:
            d_491___mcc_h9_ = source9_.cf9
            d_492___mcc_h10_ = source9_.cf10
            d_493_cf10_ = d_492___mcc_h10_
            d_494_cf9_ = d_491___mcc_h9_
            d_495_v63_: _dafny.Map
            d_495_v63_ = _dafny.Map({d_494_cf9_: d_493_cf10_})
            d_496_v64_: _dafny.Map
            d_496_v64_ = d_495_v63_
            d_497_v65_: _dafny.Map
            d_497_v65_ = _dafny.Map({d_496_v64_: not(not(True))})
            d_498_v67_: _dafny.Set
            d_498_v67_ = _dafny.Set({d_496_v64_})
            def iife68_():
                coll36_ = _dafny.Set()
                compr_36_: _dafny.Map
                for compr_36_ in (d_497_v65_).keys.Elements:
                    d_499_v66_: _dafny.Map = compr_36_
                    if (d_499_v66_) in (d_497_v65_):
                        coll36_ = coll36_.union(_dafny.Set([d_499_v66_]))
                return _dafny.Set(coll36_)
            d_376_v0_ = (d_498_v67_).ispropersubset((iife68_()
             if d_494_cf9_ else d_498_v67_))
            d_500_v68_: _dafny.MultiSet
            d_500_v68_ = _dafny.MultiSet([913])
            d_501_v69_: D9
            d_501_v69_ = D9_DC23(False, (d_500_v68_).cardinality, d_494_cf9_)
            d_502_v70_: _dafny.Array
            nw65_ = _dafny.Array(None, 11)
            nw65_[int(0)] = d_493_cf10_
            nw65_[int(1)] = d_494_cf9_
            nw65_[int(2)] = True
            nw65_[int(3)] = d_494_cf9_
            nw65_[int(4)] = d_376_v0_
            nw65_[int(5)] = d_494_cf9_
            nw65_[int(6)] = d_493_cf10_
            nw65_[int(7)] = d_493_cf10_
            nw65_[int(8)] = d_376_v0_
            nw65_[int(9)] = (d_501_v69_).cf45
            nw65_[int(10)] = d_493_cf10_
            d_502_v70_ = nw65_
            d_503_v71_: D4
            d_503_v71_ = D4_DC8(d_502_v70_)
            source11_ = d_503_v71_
            if source11_.is_DC9:
                d_504___mcc_h18_ = source11_.cf21
                d_505___mcc_h19_ = source11_.cf22
                d_506___mcc_h20_ = source11_.cf23
                d_507_cf23_ = d_506___mcc_h20_
                d_508_cf22_ = d_505___mcc_h19_
                d_509_cf21_ = d_504___mcc_h18_
                d_510_v72_: _dafny.Seq
                d_510_v72_ = _dafny.SeqWithoutIsStrInference([d_508_cf22_])
                d_511_v73_: C13
                nw66_ = C13()
                nw66_.ctor__(not(False), not(d_493_cf10_), d_493_cf10_, d_379_v2_)
                d_511_v73_ = nw66_
                d_512_v74_: _dafny.Map
                d_512_v74_ = _dafny.Map({(d_510_v72_) + (d_510_v72_): d_511_v73_})
                d_512_v74_ = (d_512_v74_).set((_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(d_511_v73_.f6), d_493_cf10_])).cardinality for d_513_i9_ in range(default__.abs(201))])) + (_dafny.SeqWithoutIsStrInference([(0) - (d_508_cf22_) for d_514_i10_ in range(default__.abs(16))])), d_511_v73_)
                d_508_cf22_ = len(d_510_v72_)
                d_515_v75_: T0
                nw67_ = C6()
                nw67_.ctor__(d_508_cf22_, True, d_494_cf9_, d_508_cf22_, d_508_cf22_, d_493_cf10_)
                d_515_v75_ = nw67_
                d_516_v76_: _dafny.Map
                d_516_v76_ = _dafny.Map({d_515_v75_: d_494_cf9_})
                d_517_v77_: _dafny.MultiSet
                d_517_v77_ = _dafny.MultiSet([d_516_v76_, d_516_v76_])
                d_518_v78_: _dafny.Map
                d_518_v78_ = _dafny.Map({d_494_cf9_: _dafny.SeqWithoutIsStrInference([d_516_v76_])})
                d_519_v79_: _dafny.Seq
                d_519_v79_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_515_v75_: d_376_v0_})])
                (d_511_v73_).f6 = (not((d_517_v77_).isdisjoint(_dafny.MultiSet(((d_518_v78_)[d_511_v73_.f7] if (d_511_v73_.f7) in (d_518_v78_) else d_519_v79_)))) if d_511_v73_.f7 else d_511_v73_.f7)
                d_379_v2_ = d_508_cf22_
            elif source11_.is_DC10:
                d_520___mcc_h21_ = source11_.cf24
                d_521___mcc_h22_ = source11_.cf25
                d_522___mcc_h23_ = source11_.cf26
                d_523_cf26_ = d_522___mcc_h23_
                d_524_cf25_ = d_521___mcc_h22_
                d_525_cf24_ = d_520___mcc_h21_
                d_526_v80_: str
                d_526_v80_ = _dafny.CodePoint('c')
                d_527_v81_: _dafny.Seq
                d_527_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilgab"))
                d_528_v82_: C7
                nw68_ = C7()
                nw68_.ctor__(d_379_v2_, d_376_v0_, d_494_cf9_, d_379_v2_, d_379_v2_, default__.fm27(d_526_v80_, d_378_globalState_), d_527_v81_, d_525_cf24_)
                d_528_v82_ = nw68_
                d_528_v82_ = d_528_v82_
                d_376_v0_ = d_376_v0_
                d_529_v83_: _dafny.Set
                d_529_v83_ = _dafny.Set({d_379_v2_})
                d_530_v84_: D11
                d_530_v84_ = D11_DC27(d_529_v83_)
                d_531_v85_: _dafny.MultiSet
                d_531_v85_ = _dafny.MultiSet([d_530_v84_])
                d_532_v86_: _dafny.Map
                d_532_v86_ = _dafny.Map({d_531_v85_: (d_528_v82_).f20})
                d_533_v87_: _dafny.Seq
                d_533_v87_ = _dafny.SeqWithoutIsStrInference([923, (d_528_v82_).f19, d_379_v2_, len(d_532_v86_), (d_528_v82_).f19])
                d_534_v88_: C6
                nw69_ = C6()
                nw69_.ctor__(d_379_v2_, d_376_v0_, not ((d_528_v82_).f20) or (d_493_cf10_), len(d_533_v87_), (d_528_v82_).f19, (not(not(d_525_cf24_))) and ((d_528_v82_).f20))
                d_534_v88_ = nw69_
                d_535_v89_: T2
                nw70_ = C2()
                nw70_.ctor__((D4_DC8(d_502_v70_)).cf20, d_500_v68_, d_527_v81_, d_494_cf9_, d_376_v0_, 803)
                d_535_v89_ = nw70_
                d_536_v90_: _dafny.Set
                d_536_v90_ = _dafny.Set({d_535_v89_})
                rhs21_ = (len((d_527_v81_) + (d_527_v81_))) * (default__.safeModuloInt((d_528_v82_).f19, len(d_536_v90_)))
                rhs22_ = d_502_v70_
                rhs23_ = (d_525_cf24_) and ((d_525_cf24_) or ((d_534_v88_).fm0(d_378_globalState_)))
                rhs24_ = 277
                d_379_v2_ = rhs21_
                d_502_v70_ = rhs22_
                d_376_v0_ = rhs23_
                d_379_v2_ = rhs24_
            elif source11_.is_DC8:
                d_537___mcc_h24_ = source11_.cf20
                d_538_cf20_ = d_537___mcc_h24_
                d_539_v91_: _dafny.Seq
                d_539_v91_ = _dafny.SeqWithoutIsStrInference([d_494_cf9_])
                d_540_v92_: str
                d_540_v92_ = _dafny.CodePoint('j')
                d_541_v93_: _dafny.Map
                d_541_v93_ = _dafny.Map({D11_DC29(d_379_v2_, d_538_cf20_, d_494_cf9_, d_493_cf10_): default__.fm18((d_539_v91_)[default__.safeIndex(d_379_v2_, len(d_539_v91_))], d_493_cf10_, default__.fm27(d_540_v92_, d_378_globalState_), d_378_globalState_)})
                d_542_v94_: D11
                d_542_v94_ = D11_DC29(d_379_v2_, d_538_cf20_, d_493_cf10_, d_376_v0_)
                d_543_v95_: _dafny.Seq
                d_543_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbhvii"))
                d_541_v93_ = (d_541_v93_).set(d_542_v94_, default__.safeDivisionInt(len(d_543_v95_), len(_dafny.Map({d_495_v63_: d_379_v2_}))))
                d_544_v96_: bool
                out3_: bool
                out3_ = default__.m0((d_543_v95_) < (d_543_v95_), d_379_v2_, d_378_globalState_)
                d_544_v96_ = out3_
                d_379_v2_ = 207
                index29_ = default__.safeIndex(882, (d_538_cf20_).length(0))
                (d_538_cf20_)[index29_] = d_544_v96_
                index30_ = default__.safeIndex(882, (d_538_cf20_).length(0))
                (d_538_cf20_)[index30_] = d_376_v0_
            elif True:
                d_545___mcc_h25_ = source11_.cf27
                d_546_cf27_ = d_545___mcc_h25_
                d_547_v97_: C10
                nw71_ = C10()
                nw71_.ctor__(default__.safeDivisionInt(len((D19_DC52(d_377_v1_)).cf93), d_379_v2_), d_502_v70_)
                d_547_v97_ = nw71_
                d_548_v98_: _dafny.Array
                nw72_ = _dafny.Array(D17.default()(), 27)
                d_548_v98_ = nw72_
                d_549_v99_: _dafny.Seq
                d_549_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                d_550_v100_: T2
                nw73_ = C2()
                nw73_.ctor__(d_502_v70_, d_500_v68_, d_549_v99_, d_494_cf9_, d_494_cf9_, (d_547_v97_).f10)
                d_550_v100_ = nw73_
                d_551_v101_: _dafny.Seq
                d_551_v101_ = _dafny.SeqWithoutIsStrInference([d_550_v100_.f18, (d_550_v100_).f4])
                d_552_v102_: C4
                nw74_ = C4()
                nw74_.ctor__(len((d_550_v100_).f17), (d_550_v100_).f4, (d_550_v100_).f17, d_494_cf9_, True, len(d_551_v101_), d_379_v2_, d_494_cf9_)
                d_552_v102_ = nw74_
                d_553_v103_: D17
                d_553_v103_ = D17_DC47(d_376_v0_, d_376_v0_, len(_dafny.Map({_dafny.Map({d_550_v100_: d_376_v0_}): _dafny.SeqWithoutIsStrInference([d_552_v102_, d_552_v102_, d_552_v102_])})), d_379_v2_, (d_552_v102_).f23)
                index31_ = default__.safeIndex(151, (d_548_v98_).length(0))
                (d_548_v98_)[index31_] = d_553_v103_
                index32_ = default__.safeIndex(151, (d_548_v98_).length(0))
                (d_548_v98_)[index32_] = d_553_v103_
                d_554_v104_: D14
                d_554_v104_ = D14_DC38((d_547_v97_).fm15(d_378_globalState_))
                d_555_v105_: C5
                nw75_ = C5()
                nw75_.ctor__()
                d_555_v105_ = nw75_
                d_556_v106_: _dafny.Array
                def lambda16_(d_557_v97_):
                    def lambda17_(d_558_i11_):
                        return (d_558_i11_) - ((d_557_v97_).f10)

                    return lambda17_

                init8_ = lambda16_(d_547_v97_)
                nw76_ = _dafny.Array(None, 27)
                for i0_8_ in range(nw76_.length(0)):
                    nw76_[i0_8_] = init8_(i0_8_)
                d_556_v106_ = nw76_
                d_559_v107_: _dafny.Map
                d_559_v107_ = _dafny.Map({d_555_v105_: d_556_v106_})
                rhs25_ = D14_DC38((d_559_v107_) == (d_559_v107_))
                rhs26_ = d_550_v100_.f18
                rhs27_ = ((d_550_v100_).f17) == ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrfnpk")) if True else d_549_v99_))
                rhs28_ = d_376_v0_
                rhs29_ = (d_552_v102_).f23
                d_554_v104_ = rhs25_
                d_494_cf9_ = rhs26_
                d_376_v0_ = rhs27_
                d_376_v0_ = rhs28_
                d_379_v2_ = rhs29_
                d_549_v99_ = d_549_v99_
            d_560_v108_: str
            d_560_v108_ = _dafny.CodePoint('i')
            d_561_v109_: _dafny.Set
            d_561_v109_ = _dafny.Set({d_560_v108_})
            d_562_v110_: _dafny.Seq
            d_562_v110_ = _dafny.SeqWithoutIsStrInference([len(d_561_v109_), d_379_v2_])
            d_379_v2_ = default__.safeDivisionInt((d_562_v110_)[default__.safeIndex(d_379_v2_, len(d_562_v110_))], d_379_v2_)
            d_379_v2_ = d_379_v2_
        d_563_v111_: bool
        out4_: bool
        out4_ = default__.m0(d_376_v0_, 67, d_378_globalState_)
        d_563_v111_ = out4_
        d_564_v112_: _dafny.Array
        nw77_ = _dafny.Array(_dafny.CodePoint('D'), 19)
        d_564_v112_ = nw77_
        d_565_v113_: str
        d_565_v113_ = _dafny.CodePoint('t')
        index33_ = default__.safeIndex(306, (d_564_v112_).length(0))
        (d_564_v112_)[index33_] = d_565_v113_
        index34_ = default__.safeIndex(306, (d_564_v112_).length(0))
        (d_564_v112_)[index34_] = _dafny.CodePoint('e')
        d_566_v114_: _dafny.Array
        def lambda18_(d_567_v0_):
            def lambda19_(d_568_i12_):
                return d_567_v0_

            return lambda19_

        init9_ = lambda18_(d_376_v0_)
        nw78_ = _dafny.Array(None, 8)
        for i0_9_ in range(nw78_.length(0)):
            nw78_[i0_9_] = init9_(i0_9_)
        d_566_v114_ = nw78_
        d_569_v115_: D4
        d_569_v115_ = D4_DC8(d_566_v114_)
        source12_ = d_569_v115_
        if source12_.is_DC9:
            d_570___mcc_h26_ = source12_.cf21
            d_571___mcc_h27_ = source12_.cf22
            d_572___mcc_h28_ = source12_.cf23
            d_573_cf23_ = d_572___mcc_h28_
            d_574_cf22_ = d_571___mcc_h27_
            d_575_cf21_ = d_570___mcc_h26_
            d_576_v116_: _dafny.Array
            def lambda20_(d_577_v0_, d_578_v2_, d_579_cf22_, d_580_v111_):
                def lambda21_(d_581_i13_):
                    return (_dafny.Map({(0) - (len(_dafny.Map({d_577_v0_: d_577_v0_}))): d_578_v2_})) | (_dafny.Map({d_579_cf22_: len(_dafny.Map({True: d_580_v111_}))}))

                return lambda21_

            init10_ = lambda20_(d_376_v0_, d_379_v2_, d_574_cf22_, d_563_v111_)
            nw79_ = _dafny.Array(None, 8)
            for i0_10_ in range(nw79_.length(0)):
                nw79_[i0_10_] = init10_(i0_10_)
            d_576_v116_ = nw79_
            d_582_v117_: _dafny.Map
            d_582_v117_ = _dafny.Map({len(_dafny.Set({d_574_cf22_})): len(d_575_cf21_)})
            index35_ = default__.safeIndex(590, (d_576_v116_).length(0))
            (d_576_v116_)[index35_] = d_582_v117_
            index36_ = default__.safeIndex(590, (d_576_v116_).length(0))
            (d_576_v116_)[index36_] = _dafny.Map({(0) - (d_574_cf22_): d_574_cf22_})
            d_583_v118_: _dafny.Seq
            d_583_v118_ = _dafny.SeqWithoutIsStrInference([d_573_cf23_])
            d_583_v118_ = d_583_v118_
            d_379_v2_ = d_379_v2_
            d_584_v119_: _dafny.Array
            nw80_ = _dafny.Array(int(0), 12)
            d_584_v119_ = nw80_
            d_584_v119_ = d_584_v119_
        elif source12_.is_DC10:
            d_585___mcc_h29_ = source12_.cf24
            d_586___mcc_h30_ = source12_.cf25
            d_587___mcc_h31_ = source12_.cf26
            d_588_cf26_ = d_587___mcc_h31_
            d_589_cf25_ = d_586___mcc_h30_
            d_590_cf24_ = d_585___mcc_h29_
            d_591_v120_: _dafny.Array
            nw81_ = _dafny.Array(int(0), 8)
            d_591_v120_ = nw81_
            d_592_v121_: D11
            d_592_v121_ = D11_DC30(d_590_cf24_, d_591_v120_, default__.fm3(d_378_globalState_), not(d_590_cf24_))
            d_379_v2_ = (d_592_v121_).cf63
            index37_ = default__.safeIndex(674, (d_591_v120_).length(0))
            (d_591_v120_)[index37_] = d_379_v2_
            d_593_v122_: _dafny.Seq
            d_593_v122_ = _dafny.SeqWithoutIsStrInference([d_588_cf26_])
            index38_ = default__.safeIndex(674, (d_591_v120_).length(0))
            (d_591_v120_)[index38_] = default__.fm25(default__.safeModuloInt(d_379_v2_, d_379_v2_), d_379_v2_, (d_593_v122_) < (_dafny.SeqWithoutIsStrInference([d_588_cf26_])), d_378_globalState_)
            d_594_v123_: _dafny.Seq
            d_594_v123_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "faar"))
            d_595_v124_: _dafny.Map
            d_595_v124_ = _dafny.Map({len(d_594_v123_): (d_591_v120_)[default__.safeIndex(674, (d_591_v120_).length(0))]})
            d_596_v125_: _dafny.Set
            d_596_v125_ = _dafny.Set({d_595_v124_})
            d_597_v126_: _dafny.Set
            d_597_v126_ = d_596_v125_
            d_598_v127_: C0
            nw82_ = C0()
            nw82_.ctor__(d_597_v126_, d_379_v2_)
            d_598_v127_ = nw82_
            d_379_v2_ = (d_591_v120_)[default__.safeIndex(674, (d_591_v120_).length(0))]
        elif source12_.is_DC8:
            d_599___mcc_h32_ = source12_.cf20
            d_600_cf20_ = d_599___mcc_h32_
            d_379_v2_ = d_379_v2_
            d_601_v128_: _dafny.Array
            nw83_ = _dafny.Array(_dafny.Seq({}), 8)
            d_601_v128_ = nw83_
            d_602_v129_: _dafny.Seq
            d_602_v129_ = _dafny.SeqWithoutIsStrInference([d_563_v111_])
            d_603_v130_: _dafny.Map
            d_603_v130_ = _dafny.Map({d_563_v111_: d_379_v2_})
            d_604_v131_: _dafny.Array
            nw84_ = _dafny.Array(None, 1)
            nw84_[int(0)] = len(d_603_v130_)
            d_604_v131_ = nw84_
            d_605_v132_: _dafny.Set
            d_605_v132_ = _dafny.Set({d_565_v113_})
            d_606_v133_: D5
            d_606_v133_ = D5_DC14(d_563_v111_, d_602_v129_, d_604_v131_, len(d_605_v132_), d_379_v2_)
            d_607_v134_: _dafny.Seq
            d_607_v134_ = _dafny.SeqWithoutIsStrInference([d_563_v111_, (d_606_v133_).cf31])
            index39_ = default__.safeIndex(51, (d_601_v128_).length(0))
            (d_601_v128_)[index39_] = d_602_v129_
            index40_ = default__.safeIndex(51, (d_601_v128_).length(0))
            (d_601_v128_)[index40_] = d_602_v129_
            d_608_v135_: C14
            nw85_ = C14()
            nw85_.ctor__(d_376_v0_, d_379_v2_)
            d_608_v135_ = nw85_
            d_609_v136_: _dafny.Map
            d_609_v136_ = _dafny.Map({d_604_v131_: d_608_v135_})
            d_610_v137_: _dafny.Seq
            d_610_v137_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "an"))
            d_379_v2_ = default__.safeDivisionInt(len(d_609_v136_), (len(d_610_v137_)) * (d_379_v2_))
            d_611_v138_: _dafny.Map
            d_611_v138_ = _dafny.Map({d_563_v111_: d_563_v111_})
            rhs30_ = (default__.fm27(default__.fm50(d_378_globalState_), d_378_globalState_) if d_563_v111_ else (d_379_v2_) < ((d_608_v135_).f3))
            rhs31_ = (d_608_v135_.f2 if default__.fm27((d_564_v112_)[default__.safeIndex(306, (d_564_v112_).length(0))], d_378_globalState_) else ((d_611_v138_)[d_563_v111_] if (d_563_v111_) in (d_611_v138_) else d_563_v111_))
            lhs6_ = d_608_v135_
            d_563_v111_ = rhs30_
            lhs6_.f2 = rhs31_
        elif True:
            d_612___mcc_h33_ = source12_.cf27
            d_613_cf27_ = d_612___mcc_h33_
            d_614_v139_: C9
            nw86_ = C9()
            nw86_.ctor__(d_566_v114_)
            d_614_v139_ = nw86_
            d_615_v140_: _dafny.MultiSet
            d_615_v140_ = _dafny.MultiSet([d_614_v139_, d_614_v139_, d_614_v139_, d_614_v139_])
            d_616_v141_: _dafny.Map
            d_616_v141_ = _dafny.Map({(d_615_v140_).cardinality: (d_376_v0_ if d_376_v0_ else d_376_v0_)})
            d_376_v0_ = ((d_616_v141_)[d_379_v2_] if (d_379_v2_) in (d_616_v141_) else d_563_v111_)
            d_617_v142_: _dafny.Seq
            d_617_v142_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldga"))
            d_618_v143_: _dafny.MultiSet
            d_618_v143_ = _dafny.MultiSet([d_379_v2_, len(d_617_v142_)])
            d_619_v144_: C2
            nw87_ = C2()
            nw87_.ctor__(d_566_v114_, d_618_v143_, d_617_v142_, default__.fm27(d_565_v113_, d_378_globalState_), d_563_v111_, d_379_v2_)
            d_619_v144_ = nw87_
            d_620_v145_: _dafny.Map
            d_620_v145_ = _dafny.Map({len(d_617_v142_): d_619_v144_})
            d_620_v145_ = (d_620_v145_).set(default__.safeDivisionInt(d_379_v2_, d_379_v2_), d_619_v144_)
            d_621_v146_: C14
            nw88_ = C14()
            nw88_.ctor__(d_563_v111_, d_379_v2_)
            d_621_v146_ = nw88_
            d_622_v147_: _dafny.Map
            d_622_v147_ = _dafny.Map({(d_621_v146_).f3: d_379_v2_})
            d_623_v148_: _dafny.Set
            d_623_v148_ = _dafny.Set({d_379_v2_})
            d_622_v147_ = (d_622_v147_).set(len((d_623_v148_) | (d_623_v148_)), (d_379_v2_) + (d_379_v2_))
        if d_563_v111_:
            d_624_v149_: _dafny.Array
            nw89_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_624_v149_ = nw89_
            d_624_v149_ = d_624_v149_
            d_379_v2_ = (d_379_v2_) + (len(_dafny.SeqWithoutIsStrInference([(d_564_v112_)[default__.safeIndex(306, (d_564_v112_).length(0))] for d_625_i14_ in range(default__.abs(688))])))
            d_376_v0_ = False
            d_626_v150_: _dafny.MultiSet
            d_626_v150_ = _dafny.MultiSet([771])
            d_627_v151_: _dafny.Map
            d_627_v151_ = _dafny.Map({d_566_v114_: d_376_v0_})
            d_626_v150_ = (d_626_v150_).set(len(d_627_v151_), default__.abs(d_379_v2_))
            d_566_v114_ = d_566_v114_
        elif True:
            d_628_v152_: _dafny.Seq
            d_628_v152_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbelxm"))
            d_629_v153_: T2
            nw90_ = C7()
            nw90_.ctor__(d_379_v2_, d_376_v0_, d_563_v111_, (0) - (d_379_v2_), d_379_v2_, d_563_v111_, d_628_v152_, not(d_563_v111_))
            d_629_v153_ = nw90_
            d_630_v154_: _dafny.Map
            d_630_v154_ = _dafny.Map({d_629_v153_: (d_629_v153_).f5})
            d_631_v155_: T0
            nw91_ = C4()
            nw91_.ctor__(d_379_v2_, d_563_v111_, d_628_v152_, not(d_376_v0_), d_563_v111_, 624, len(d_630_v154_), d_376_v0_)
            d_631_v155_ = nw91_
            d_632_v156_: _dafny.MultiSet
            d_632_v156_ = _dafny.MultiSet([d_631_v155_, d_631_v155_])
            d_379_v2_ = default__.safeDivisionInt((d_379_v2_) - ((d_632_v156_).cardinality), (d_631_v155_).f5)
            (d_629_v153_).f18 = False
            d_633_v157_: _dafny.MultiSet
            d_633_v157_ = _dafny.MultiSet([d_566_v114_])
            d_376_v0_ = (d_633_v157_).ispropersubset(d_633_v157_)
            d_634_v159_: _dafny.MultiSet
            d_634_v159_ = _dafny.MultiSet([(d_629_v153_).f17])
            d_635_v160_: D7
            d_635_v160_ = D7_DC19(d_379_v2_)
            d_636_v161_: _dafny.Map
            d_636_v161_ = _dafny.Map({(d_629_v153_).f17: d_635_v160_})
            d_637_v162_: _dafny.Map
            def iife69_():
                coll37_ = _dafny.Map()
                compr_37_: _dafny.Seq
                for compr_37_ in (d_634_v159_).Elements:
                    d_638_v158_: _dafny.Seq = compr_37_
                    if (d_638_v158_) in (d_634_v159_):
                        coll37_[d_638_v158_] = D7_DC19((d_629_v153_).f5)
                return _dafny.Map(coll37_)
            d_637_v162_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdjle")): (iife69_()
            ) | (d_636_v161_)})
            d_637_v162_ = (d_637_v162_).set(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_639_i15_ in range(default__.abs(5))]), d_636_v161_)
            d_640_v163_: _dafny.Seq
            d_640_v163_ = _dafny.SeqWithoutIsStrInference([d_376_v0_, (d_631_v155_).f4, d_563_v111_])
            d_640_v163_ = _dafny.SeqWithoutIsStrInference([d_629_v153_.f18, (d_629_v153_).f4, True, d_563_v111_, (d_629_v153_).f4])
        if d_376_v0_:
            d_641_v164_: _dafny.Seq
            d_641_v164_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xfj"))
            d_642_v165_: _dafny.Seq
            d_642_v165_ = _dafny.SeqWithoutIsStrInference([d_379_v2_, len(d_641_v164_)])
            d_642_v165_ = (_dafny.SeqWithoutIsStrInference([d_379_v2_, d_379_v2_])) + (d_642_v165_)
            d_563_v111_ = d_376_v0_
            d_643_v166_: _dafny.Seq
            d_643_v166_ = _dafny.SeqWithoutIsStrInference([d_563_v111_])
            d_644_v167_: T2
            nw92_ = C4()
            nw92_.ctor__(d_379_v2_, d_563_v111_, _dafny.SeqWithoutIsStrInference([d_565_v113_ for d_645_i16_ in range(default__.abs(854))]), d_376_v0_, d_563_v111_, len(d_643_v166_), d_379_v2_, d_376_v0_)
            d_644_v167_ = nw92_
            d_646_v168_: _dafny.Seq
            d_646_v168_ = _dafny.SeqWithoutIsStrInference([d_644_v167_, d_644_v167_])
            d_647_v169_: C12
            nw93_ = C12()
            nw93_.ctor__((len(d_646_v168_)) - ((d_644_v167_).f5))
            d_647_v169_ = nw93_
            d_647_v169_ = d_647_v169_
            d_642_v165_ = ((_dafny.SeqWithoutIsStrInference([(d_644_v167_).f5 for d_648_i17_ in range(default__.abs(707))])) + (_dafny.SeqWithoutIsStrInference([(d_647_v169_).f8 for d_649_i18_ in range(default__.abs(-846))]))) + (d_642_v165_)
            d_650_v170_: _dafny.Set
            d_650_v170_ = _dafny.Set({(d_644_v167_).f5, (d_644_v167_).f5, 434})
            d_651_v171_: _dafny.Map
            d_651_v171_ = _dafny.Map({(d_650_v170_) - (d_650_v170_): (d_647_v169_).f8})
            d_651_v171_ = d_651_v171_
        elif True:
            d_652_v172_: _dafny.Map
            d_652_v172_ = _dafny.Map({d_376_v0_: d_379_v2_})
            d_653_v173_: _dafny.Map
            d_653_v173_ = _dafny.Map({d_652_v172_: d_652_v172_})
            d_654_v174_: _dafny.Map
            d_654_v174_ = _dafny.Map({d_379_v2_: d_653_v173_})
            def iife70_():
                coll38_ = _dafny.Map()
                compr_38_: _dafny.Map
                for compr_38_ in (default__.fm86(d_376_v0_, d_563_v111_, d_379_v2_, d_378_globalState_)).Elements:
                    d_655_v175_: _dafny.Map = compr_38_
                    if (d_655_v175_) in (default__.fm86(d_376_v0_, d_563_v111_, d_379_v2_, d_378_globalState_)):
                        coll38_[d_655_v175_] = _dafny.Map({not(d_376_v0_): d_379_v2_})
                return _dafny.Map(coll38_)
            d_654_v174_ = (d_654_v174_).set(d_379_v2_, iife70_()
            )
            d_656_v176_: _dafny.Seq
            d_656_v176_ = _dafny.SeqWithoutIsStrInference([d_563_v111_])
            d_657_v177_: _dafny.Seq
            d_657_v177_ = _dafny.SeqWithoutIsStrInference([d_379_v2_])
            d_658_v178_: _dafny.Set
            d_658_v178_ = _dafny.Set({d_379_v2_, d_379_v2_})
            d_659_v179_: _dafny.MultiSet
            d_659_v179_ = _dafny.MultiSet([d_379_v2_, 464, d_379_v2_, (0) - (len(d_657_v177_))])
            d_660_v180_: _dafny.Array
            nw94_ = _dafny.Array(None, 15)
            nw94_[int(0)] = len(d_656_v176_)
            nw94_[int(1)] = d_379_v2_
            nw94_[int(2)] = (d_379_v2_) - (d_379_v2_)
            nw94_[int(3)] = d_379_v2_
            nw94_[int(4)] = d_379_v2_
            nw94_[int(5)] = (d_379_v2_) + (d_379_v2_)
            nw94_[int(6)] = d_379_v2_
            nw94_[int(7)] = d_379_v2_
            nw94_[int(8)] = (d_379_v2_) * (len(d_657_v177_))
            nw94_[int(9)] = len(d_658_v178_)
            nw94_[int(10)] = d_379_v2_
            nw94_[int(11)] = d_379_v2_
            nw94_[int(12)] = (d_379_v2_) + (d_379_v2_)
            nw94_[int(13)] = ((d_659_v179_)[len((d_657_v177_).set(default__.safeIndex(d_379_v2_, len(d_657_v177_)), d_379_v2_))] if (len((d_657_v177_).set(default__.safeIndex(d_379_v2_, len(d_657_v177_)), d_379_v2_))) in (d_659_v179_) else d_379_v2_)
            nw94_[int(14)] = d_379_v2_
            d_660_v180_ = nw94_
            d_661_v181_: _dafny.Seq
            d_661_v181_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahj"))
            d_662_v182_: T0
            nw95_ = C4()
            nw95_.ctor__(d_379_v2_, d_376_v0_, d_661_v181_, False, d_376_v0_, d_379_v2_, d_379_v2_, d_376_v0_)
            d_662_v182_ = nw95_
            d_663_v183_: _dafny.MultiSet
            d_663_v183_ = _dafny.MultiSet([d_662_v182_])
            index41_ = default__.safeIndex(211, (d_660_v180_).length(0))
            (d_660_v180_)[index41_] = (0) - ((d_379_v2_) - ((d_663_v183_).cardinality))
            d_664_v184_: D0
            d_664_v184_ = D0_DC0((d_662_v182_).f4, (d_659_v179_).cardinality, d_563_v111_, (d_662_v182_).f5, (d_662_v182_).f4)
            index42_ = default__.safeIndex(211, (d_660_v180_).length(0))
            (d_660_v180_)[index42_] = (0) - ((0) - (default__.fm32(d_664_v184_, d_378_globalState_)))
            index43_ = default__.safeIndex(644, (d_566_v114_).length(0))
            (d_566_v114_)[index43_] = (d_662_v182_).f4
            index44_ = default__.safeIndex(644, (d_566_v114_).length(0))
            (d_566_v114_)[index44_] = (d_662_v182_).fm0(d_378_globalState_)
            index45_ = default__.safeIndex(644, (d_566_v114_).length(0))
            (d_566_v114_)[index45_] = (d_566_v114_)[default__.safeIndex(644, (d_566_v114_).length(0))]
            (d_662_v182_).m3(d_379_v2_, d_378_globalState_)
        d_665_v185_: _dafny.Array
        def lambda22_(d_666_v2_, d_667_v0_):
            def lambda23_(d_668_i20_):
                return _dafny.Map({d_666_v2_: d_667_v0_})

            return lambda23_

        init11_ = lambda22_(d_379_v2_, d_376_v0_)
        nw96_ = _dafny.Array(None, 2)
        for i0_11_ in range(nw96_.length(0)):
            nw96_[i0_11_] = init11_(i0_11_)
        d_665_v185_ = nw96_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_665_v185_).length(0)):
            d_669_i19_: int = guard_loop_0_
            if (True) and (((0) <= (d_669_i19_)) and ((d_669_i19_) < ((d_665_v185_).length(0)))):
                (d_665_v185_)[(d_669_i19_)] = ((_dafny.Map({d_379_v2_: d_563_v111_})) | (_dafny.Map({239: d_563_v111_}))) | (_dafny.Map({d_379_v2_: d_563_v111_}))
        d_379_v2_ = (d_379_v2_) - ((d_379_v2_) - (d_379_v2_))
        d_670_v186_: _dafny.Seq
        d_670_v186_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
        d_671_v187_: D22
        d_671_v187_ = D22_DC58(d_670_v186_, (D22_DC58(d_670_v186_, d_379_v2_, d_376_v0_)).cf103, d_376_v0_)
        d_379_v2_ = (d_671_v187_).cf103
        hi0_ = (0) - (default__.safeDivisionInt((0) - (d_379_v2_), d_379_v2_))
        for d_672_i21_ in range(-800, hi0_):
            d_673_v188_: D16
            d_673_v188_ = D16_DC41(d_565_v113_, d_376_v0_)
            source13_ = d_673_v188_
            if source13_.is_DC41:
                d_674___mcc_h34_ = source13_.cf77
                d_675___mcc_h35_ = source13_.cf78
                d_676_cf78_ = d_675___mcc_h35_
                d_677_cf77_ = d_674___mcc_h34_
                d_678_v189_: bool
                out5_: bool
                out5_ = default__.m0(not (d_676_cf78_) or (False), d_672_i21_, d_378_globalState_)
                d_678_v189_ = out5_
                d_679_v190_: _dafny.Set
                d_679_v190_ = _dafny.Set({(0) - (d_672_i21_)})
                d_680_v191_: _dafny.Map
                d_680_v191_ = _dafny.Map({d_379_v2_: d_379_v2_})
                d_681_v192_: _dafny.Map
                d_681_v192_ = _dafny.Map({d_679_v190_: (d_672_i21_) < (len(d_680_v191_))})
                d_682_v193_: _dafny.Seq
                d_682_v193_ = _dafny.SeqWithoutIsStrInference([d_678_v189_])
                d_683_v194_: D4
                d_683_v194_ = D4_DC10(d_678_v189_, d_678_v189_, d_682_v193_)
                d_684_v195_: _dafny.Seq
                d_684_v195_ = _dafny.SeqWithoutIsStrInference([True, d_676_cf78_, d_563_v111_, True, (d_683_v194_).cf25])
                d_685_v196_: _dafny.Array
                nw97_ = _dafny.Array(int(0), 21)
                d_685_v196_ = nw97_
                d_686_v197_: D5
                d_686_v197_ = D5_DC14(d_376_v0_, d_682_v193_, d_685_v196_, len(d_670_v186_), d_672_i21_)
                d_681_v192_ = (d_681_v192_).set(d_679_v190_, (d_686_v197_).cf31)
                d_687_v198_: _dafny.Seq
                d_687_v198_ = _dafny.SeqWithoutIsStrInference([d_670_v186_])
                d_688_v199_: _dafny.Map
                d_688_v199_ = _dafny.Map({(0) - (len((d_687_v198_)[default__.safeIndex(d_379_v2_, len(d_687_v198_))])): D1_DC3(d_682_v193_)})
                d_689_v200_: bool
                out6_: bool
                out6_ = default__.m0(d_563_v111_, len((d_688_v199_) | (d_688_v199_)), d_378_globalState_)
                d_689_v200_ = out6_
                d_690_v201_: bool
                out7_: bool
                out7_ = default__.m0(d_678_v189_, default__.safeModuloInt(d_379_v2_, len(d_680_v191_)), d_378_globalState_)
                d_690_v201_ = out7_
            elif source13_.is_DC42:
                d_691___mcc_h36_ = source13_.cf79
                d_692___mcc_h37_ = source13_.cf80
                d_693_cf80_ = d_692___mcc_h37_
                d_694_cf79_ = d_691___mcc_h36_
                d_695_v202_: _dafny.Array
                def lambda24_(d_696_v2_):
                    def lambda25_(d_697_i22_):
                        return (d_697_i22_) + (d_696_v2_)

                    return lambda25_

                init12_ = lambda24_(d_379_v2_)
                nw98_ = _dafny.Array(None, 14)
                for i0_12_ in range(nw98_.length(0)):
                    nw98_[i0_12_] = init12_(i0_12_)
                d_695_v202_ = nw98_
                index46_ = default__.safeIndex(313, (d_695_v202_).length(0))
                (d_695_v202_)[index46_] = d_379_v2_
                d_698_v203_: _dafny.Map
                d_698_v203_ = _dafny.Map({d_566_v114_: d_672_i21_})
                index47_ = default__.safeIndex(313, (d_695_v202_).length(0))
                (d_695_v202_)[index47_] = len(((d_698_v203_) | (d_698_v203_)) | (d_698_v203_))
                index48_ = default__.safeIndex(313, (d_695_v202_).length(0))
                (d_695_v202_)[index48_] = (d_695_v202_)[default__.safeIndex(313, (d_695_v202_).length(0))]
                d_694_cf79_ = d_563_v111_
                d_699_v204_: D12
                d_699_v204_ = D12_DC33(d_693_cf80_, d_694_cf79_, _dafny.CodePoint('x'), 26)
                d_379_v2_ = (d_699_v204_).cf70
            elif source13_.is_DC40:
                d_700___mcc_h38_ = source13_.cf76
                d_701_cf76_ = d_700___mcc_h38_
                d_702_v205_: D1
                d_702_v205_ = D1_DC5(d_376_v0_, d_672_i21_, d_670_v186_)
                d_703_v206_: C11
                nw99_ = C11()
                nw99_.ctor__(d_702_v205_, d_376_v0_, len(_dafny.Map({135: d_563_v111_})))
                d_703_v206_ = nw99_
                d_704_v207_: _dafny.Seq
                d_704_v207_ = _dafny.SeqWithoutIsStrInference([d_703_v206_])
                d_704_v207_ = d_704_v207_
                d_563_v111_ = ((d_701_cf76_)[(0) - (d_379_v2_)] if ((0) - (d_379_v2_)) in (d_701_cf76_) else d_376_v0_)
                d_705_v208_: C9
                nw100_ = C9()
                nw100_.ctor__(d_566_v114_)
                d_705_v208_ = nw100_
                d_705_v208_ = d_705_v208_
                def iife71_():
                    coll39_ = _dafny.Map()
                    compr_39_: int
                    for compr_39_ in _dafny.IntegerRange(-963, 429):
                        d_706_v209_: int = compr_39_
                        if ((-963) <= (d_706_v209_)) and ((d_706_v209_) < (429)):
                            coll39_[default__.safeDivisionInt(d_706_v209_, -308)] = 193
                    return _dafny.Map(coll39_)
                d_379_v2_ = len(default__.fm57(iife71_()
                , d_672_i21_, d_378_globalState_))
            elif True:
                d_707___mcc_h39_ = source13_.cf81
                d_708_cf81_ = d_707___mcc_h39_
                d_709_v210_: _dafny.Seq
                d_709_v210_ = _dafny.SeqWithoutIsStrInference([d_566_v114_])
                d_709_v210_ = ((_dafny.SeqWithoutIsStrInference([d_566_v114_])) + (d_709_v210_)) + (d_709_v210_)
                d_670_v186_ = (d_671_v187_).cf102
                d_710_v211_: _dafny.Seq
                d_710_v211_ = _dafny.SeqWithoutIsStrInference([d_672_i21_])
                d_711_v212_: _dafny.Seq
                d_711_v212_ = _dafny.SeqWithoutIsStrInference([d_563_v111_])
                d_712_v213_: _dafny.MultiSet
                d_712_v213_ = _dafny.MultiSet([d_563_v111_])
                d_713_v214_: _dafny.Seq
                d_713_v214_ = _dafny.SeqWithoutIsStrInference([d_712_v213_])
                d_714_v215_: C14
                nw101_ = C14()
                nw101_.ctor__(True, default__.safeModuloInt(992, (default__.fm87(d_710_v211_, (d_711_v212_)[default__.safeIndex(((d_713_v214_)[default__.safeIndex(d_379_v2_, len(d_713_v214_))]).cardinality, len(d_711_v212_))], d_378_globalState_)).cardinality))
                d_714_v215_ = nw101_
                d_376_v0_ = (d_379_v2_) in (_dafny.MultiSet(d_710_v211_))
            d_715_v216_: _dafny.MultiSet
            d_715_v216_ = _dafny.MultiSet([d_672_i21_, default__.safeDivisionInt(len(d_670_v186_), d_672_i21_), d_672_i21_])
            d_716_v217_: _dafny.Seq
            d_716_v217_ = _dafny.SeqWithoutIsStrInference([not(d_563_v111_)])
            rhs32_ = d_379_v2_
            rhs33_ = d_715_v216_
            rhs34_ = d_376_v0_
            rhs35_ = (d_716_v217_)[default__.safeIndex(default__.safeDivisionInt(d_379_v2_, d_672_i21_), len(d_716_v217_))]
            rhs36_ = default__.fm48(d_672_i21_, d_378_globalState_)
            lhs7_ = d_378_globalState_
            d_379_v2_ = rhs32_
            d_715_v216_ = rhs33_
            d_563_v111_ = rhs34_
            d_376_v0_ = rhs35_
            lhs7_.f0 = rhs36_
            d_670_v186_ = default__.fm28(default__.fm37(d_378_globalState_), d_379_v2_, d_379_v2_, d_378_globalState_)
            d_717_v218_: C8
            nw102_ = C8()
            nw102_.ctor__(_dafny.SeqWithoutIsStrInference([d_565_v113_ for d_718_i23_ in range(default__.abs(64))]), d_672_i21_)
            d_717_v218_ = nw102_
            d_719_v219_: T1
            nw103_ = C4()
            nw103_.ctor__(len((d_716_v217_) + (d_716_v217_)), default__.fm27(d_565_v113_, d_378_globalState_), d_717_v218_.f13, d_376_v0_, not(d_563_v111_), (d_379_v2_) + ((d_717_v218_).f14), (d_717_v218_).f14, not(d_563_v111_))
            d_719_v219_ = nw103_
            rhs37_ = len((_dafny.SeqWithoutIsStrInference([(d_719_v219_).f16, (d_719_v219_).f4])) + (d_716_v217_))
            rhs38_ = d_717_v218_
            rhs39_ = d_719_v219_
            rhs40_ = (d_717_v218_).f14
            d_379_v2_ = rhs37_
            d_717_v218_ = rhs38_
            d_719_v219_ = rhs39_
            d_379_v2_ = rhs40_
        d_379_v2_ = d_379_v2_
        d_720_v220_: _dafny.Map
        d_720_v220_ = _dafny.Map({d_379_v2_: d_563_v111_})
        d_721_v221_: D1
        d_721_v221_ = D1_DC5(((d_720_v220_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgnnye")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgnnye")))) in (d_720_v220_) else d_563_v111_), d_379_v2_, d_670_v186_)
        d_722_v222_: D18
        d_722_v222_ = D18_DC50(True, d_376_v0_)
        d_723_v223_: _dafny.Seq
        d_723_v223_ = _dafny.SeqWithoutIsStrInference([True])
        d_724_v224_: C11
        nw104_ = C11()
        nw104_.ctor__(d_721_v221_, (default__.fm27((d_564_v112_)[default__.safeIndex(306, (d_564_v112_).length(0))], d_378_globalState_) if d_563_v111_ else (d_722_v222_).cf90), (len((d_723_v223_).set(default__.safeIndex(len(d_670_v186_), len(d_723_v223_)), (d_723_v223_)[default__.safeIndex(d_379_v2_, len(d_723_v223_))]))) + (d_379_v2_))
        d_724_v224_ = nw104_
        d_725_v225_: C9
        nw105_ = C9()
        nw105_.ctor__(d_566_v114_)
        d_725_v225_ = nw105_
        if False:
            d_726_v226_: _dafny.Array
            nw106_ = _dafny.Array(D9.default()(), 1)
            d_726_v226_ = nw106_
            d_727_v228_: T2
            nw107_ = C4()
            nw107_.ctor__(d_379_v2_, (d_724_v224_).fm0(d_378_globalState_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpnvhw")), d_563_v111_, d_563_v111_, d_379_v2_, d_379_v2_, d_563_v111_)
            d_727_v228_ = nw107_
            d_728_v229_: _dafny.Map
            d_728_v229_ = _dafny.Map({d_727_v228_: (d_727_v228_).f5})
            d_729_v230_: _dafny.Map
            d_729_v230_ = _dafny.Map({d_379_v2_: (d_727_v228_).f5})
            d_730_v231_: _dafny.Map
            d_730_v231_ = _dafny.Map({(d_727_v228_).f5: d_565_v113_})
            d_731_v232_: _dafny.Array
            nw108_ = _dafny.Array(None, 24)
            nw108_[int(0)] = ((d_728_v229_)[d_727_v228_] if (d_727_v228_) in (d_728_v229_) else d_379_v2_)
            nw108_[int(1)] = len(d_729_v230_)
            nw108_[int(2)] = (d_727_v228_).f5
            nw108_[int(3)] = (d_727_v228_).f5
            nw108_[int(4)] = d_379_v2_
            nw108_[int(5)] = d_379_v2_
            nw108_[int(6)] = (d_727_v228_).f5
            nw108_[int(7)] = len(d_730_v231_)
            nw108_[int(8)] = (d_727_v228_).f5
            nw108_[int(9)] = d_379_v2_
            nw108_[int(10)] = len(d_723_v223_)
            nw108_[int(11)] = (d_727_v228_).f5
            nw108_[int(12)] = (0) - ((d_727_v228_).f5)
            nw108_[int(13)] = (d_727_v228_).f5
            nw108_[int(14)] = d_379_v2_
            nw108_[int(15)] = d_379_v2_
            nw108_[int(16)] = (d_727_v228_).f5
            nw108_[int(17)] = d_379_v2_
            nw108_[int(18)] = (d_727_v228_).f5
            nw108_[int(19)] = (d_727_v228_).f5
            nw108_[int(20)] = d_379_v2_
            nw108_[int(21)] = 575
            nw108_[int(22)] = 551
            nw108_[int(23)] = (d_727_v228_).f5
            d_731_v232_ = nw108_
            d_732_v233_: _dafny.Map
            def iife72_():
                coll40_ = _dafny.Map()
                compr_40_: int
                for compr_40_ in _dafny.IntegerRange(76, 551):
                    d_733_v227_: int = compr_40_
                    if ((76) <= (d_733_v227_)) and ((d_733_v227_) < (551)):
                        coll40_[default__.safeDivisionInt(d_733_v227_, d_379_v2_)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_565_v113_ for d_734_i24_ in range(default__.abs(801))]))).cardinality
                return _dafny.Map(coll40_)
            d_732_v233_ = _dafny.Map({len(iife72_()
            ): d_731_v232_})
            d_735_v234_: D9
            d_735_v234_ = D9_DC21(d_732_v233_)
            index49_ = default__.safeIndex(404, (d_726_v226_).length(0))
            (d_726_v226_)[index49_] = d_735_v234_
            d_736_v235_: _dafny.Seq
            d_736_v235_ = _dafny.SeqWithoutIsStrInference([d_670_v186_])
            d_737_v236_: _dafny.Map
            d_737_v236_ = _dafny.Map({d_563_v111_: d_732_v233_})
            pat_let_tv9_ = d_737_v236_
            pat_let_tv10_ = d_727_v228_
            pat_let_tv11_ = d_727_v228_
            pat_let_tv12_ = d_737_v236_
            pat_let_tv13_ = d_732_v233_
            index50_ = default__.safeIndex(404, (d_726_v226_).length(0))
            def iife73_(_pat_let16_0):
                def iife74_(d_738_dt__update__tmp_h2_):
                    def iife75_(_pat_let17_0):
                        def iife76_(d_739_dt__update_hcf43_h0_):
                            return D9_DC21(d_739_dt__update_hcf43_h0_)
                        return iife76_(_pat_let17_0)
                    return iife75_(((pat_let_tv9_)[(pat_let_tv10_).f4] if ((pat_let_tv11_).f4) in (pat_let_tv12_) else pat_let_tv13_))
                return iife74_(_pat_let16_0)
            rhs41_ = ((d_736_v235_)[default__.safeIndex((0) - (d_379_v2_), len(d_736_v235_))]) + ((d_727_v228_).f17)
            rhs42_ = iife73_(d_735_v234_)
            rhs43_ = (d_727_v228_).f5
            lhs8_ = d_726_v226_
            lhs9_ = default__.safeIndex(404, (d_726_v226_).length(0))
            d_670_v186_ = rhs41_
            lhs8_[lhs9_] = rhs42_
            d_379_v2_ = rhs43_
            d_565_v113_ = _dafny.CodePoint('f')
            d_740_v237_: C10
            nw109_ = C10()
            nw109_.ctor__(d_379_v2_, (d_725_v225_).f12)
            d_740_v237_ = nw109_
            d_723_v223_ = _dafny.SeqWithoutIsStrInference([(d_740_v237_) in (_dafny.Map({d_740_v237_: (d_727_v228_).f5})), d_563_v111_])
            d_741_v238_: _dafny.Seq
            d_741_v238_ = _dafny.SeqWithoutIsStrInference([d_723_v223_])
            d_742_v239_: _dafny.Map
            d_742_v239_ = _dafny.Map({len(d_720_v220_): d_741_v238_})
            d_742_v239_ = (d_742_v239_).set((d_379_v2_ if not(d_376_v0_) else (d_727_v228_).f5), (d_741_v238_).set(default__.safeIndex(d_379_v2_, len(d_741_v238_)), d_723_v223_))
            if ((d_727_v228_).f4 if d_563_v111_ else d_376_v0_):
                d_376_v0_ = not(((d_727_v228_).f5) != ((len(_dafny.SeqWithoutIsStrInference([not(d_563_v111_)]))) + ((d_740_v237_).f10)))
                d_743_v240_: _dafny.Seq
                out8_: _dafny.Seq
                out8_ = (d_725_v225_).m8(not((d_727_v228_).f4), (d_727_v228_).f4, ((d_727_v228_).f5) + ((d_740_v237_).f10), d_378_globalState_)
                d_743_v240_ = out8_
                (d_727_v228_).f18 = not (d_727_v228_.f18) or ((d_565_v113_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecvsptxut"))))
                (d_727_v228_).m3(-952, d_378_globalState_)
                d_379_v2_ = default__.fm26(d_378_globalState_)
            elif True:
                d_744_v241_: C1
                nw110_ = C1()
                nw110_.ctor__(107, d_727_v228_.f18, (len(d_670_v186_)) != ((d_727_v228_).f5), ((d_727_v228_).f5) - ((d_740_v237_).f10))
                d_744_v241_ = nw110_
                d_745_v242_: _dafny.Map
                d_745_v242_ = _dafny.Map({(d_727_v228_).f4: (d_740_v237_).f10})
                d_746_v243_: _dafny.Seq
                d_746_v243_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_727_v228_).f5), len(d_745_v242_)])
                d_747_v244_: _dafny.Seq
                d_747_v244_ = _dafny.SeqWithoutIsStrInference([d_746_v243_])
                d_748_v245_: _dafny.Array
                nw111_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_748_v245_ = nw111_
                d_749_v246_: _dafny.Map
                d_749_v246_ = _dafny.Map({d_748_v245_: d_563_v111_})
                nw112_ = C1()
                nw112_.ctor__(len(d_747_v244_), default__.fm27((d_564_v112_)[default__.safeIndex(306, (d_564_v112_).length(0))], d_378_globalState_), ((d_749_v246_)[d_748_v245_] if (d_748_v245_) in (d_749_v246_) else d_727_v228_.f18), (d_727_v228_).f5)
                d_744_v241_ = nw112_
                d_563_v111_ = False
                d_750_v247_: D21
                d_750_v247_ = D21_DC56((d_727_v228_).f4)
                d_751_v248_: _dafny.Array
                nw113_ = _dafny.Array(None, 13)
                nw113_[int(0)] = d_750_v247_
                nw113_[int(1)] = D21_DC56(d_376_v0_)
                nw113_[int(2)] = d_750_v247_
                nw113_[int(3)] = d_750_v247_
                nw113_[int(4)] = d_750_v247_
                nw113_[int(5)] = d_750_v247_
                nw113_[int(6)] = d_750_v247_
                nw113_[int(7)] = d_750_v247_
                nw113_[int(8)] = d_750_v247_
                nw113_[int(9)] = d_750_v247_
                nw113_[int(10)] = d_750_v247_
                nw113_[int(11)] = d_750_v247_
                nw113_[int(12)] = d_750_v247_
                d_751_v248_ = nw113_
                d_751_v248_ = d_751_v248_
                (d_727_v228_).f18 = d_376_v0_
                d_744_v241_ = d_744_v241_
        elif True:
            d_752_v249_: _dafny.Array
            nw114_ = _dafny.Array(_dafny.Seq({}), 7)
            d_752_v249_ = nw114_
            d_753_v250_: _dafny.Seq
            d_753_v250_ = _dafny.SeqWithoutIsStrInference([d_379_v2_, (0) - (d_379_v2_)])
            d_754_v251_: D6
            d_754_v251_ = D6_DC15(d_753_v250_)
            index51_ = default__.safeIndex(472, (d_752_v249_).length(0))
            (d_752_v249_)[index51_] = ((_dafny.SeqWithoutIsStrInference([d_563_v111_, d_376_v0_, d_563_v111_, d_563_v111_, not(d_563_v111_)])) + (d_723_v223_)).set(default__.safeIndex((0) - (d_379_v2_), len((_dafny.SeqWithoutIsStrInference([d_563_v111_, d_376_v0_, d_563_v111_, d_563_v111_, not(d_563_v111_)])) + (d_723_v223_))), d_563_v111_)
            d_755_v252_: D19
            d_755_v252_ = D19_DC52(_dafny.Set({d_563_v111_, False}))
            pat_let_tv14_ = d_377_v1_
            index52_ = default__.safeIndex(472, (d_752_v249_).length(0))
            def iife77_(_pat_let18_0):
                def iife78_(d_756_dt__update__tmp_h3_):
                    def iife79_(_pat_let19_0):
                        def iife80_(d_757_dt__update_hcf93_h0_):
                            return D19_DC52(d_757_dt__update_hcf93_h0_)
                        return iife80_(_pat_let19_0)
                    return iife79_(pat_let_tv14_)
                return iife78_(_pat_let18_0)
            rhs44_ = (len((iife77_(d_755_v252_)).cf93)) != (-853)
            rhs45_ = d_752_v249_
            rhs46_ = D6_DC15(d_753_v250_)
            rhs47_ = d_723_v223_
            lhs10_ = d_752_v249_
            lhs11_ = default__.safeIndex(472, (d_752_v249_).length(0))
            d_563_v111_ = rhs44_
            d_752_v249_ = rhs45_
            d_754_v251_ = rhs46_
            lhs10_[lhs11_] = rhs47_
            d_758_v253_: _dafny.MultiSet
            d_758_v253_ = _dafny.MultiSet([(d_379_v2_) - ((0) - (d_379_v2_)), d_379_v2_])
            d_758_v253_ = d_758_v253_
            d_759_v254_: bool
            out9_: bool
            out9_ = default__.m0(d_563_v111_, (0) - ((d_379_v2_) + ((0) - (d_379_v2_))), d_378_globalState_)
            d_759_v254_ = out9_
            d_376_v0_ = (d_670_v186_) == ((d_670_v186_) + (d_670_v186_))
            d_760_v255_: C10
            nw115_ = C10()
            nw115_.ctor__((0) - ((990) + (d_379_v2_)), d_566_v114_)
            d_760_v255_ = nw115_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_566_v114_).length(0)):
            d_761_i25_: int = guard_loop_1_
            if (True) and (((0) <= (d_761_i25_)) and ((d_761_i25_) < ((d_566_v114_).length(0)))):
                (d_566_v114_)[(d_761_i25_)] = d_563_v111_
        d_563_v111_ = d_563_v111_
        _dafny.print(_dafny.string_of(d_376_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_377_v1_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_378_globalState_.f0) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_378_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_379_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_563_v111_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_564_v112_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_565_v113_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_566_v114_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_566_v114_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_566_v114_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_566_v114_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_566_v114_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_566_v114_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_566_v114_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_566_v114_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_569_v115_).cf20)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_569_v115_).cf20)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_569_v115_).cf20)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_569_v115_).cf20)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_569_v115_).cf20)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_569_v115_).cf20)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_569_v115_).cf20)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_569_v115_).cf20)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_665_v185_)[0]) == (_dafny.Map({-1: False, 239: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_665_v185_)[1]) == (_dafny.Map({-1: False, 239: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_670_v186_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_671_v187_).cf102).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_671_v187_).cf103))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_671_v187_).cf104))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_720_v220_) == (_dafny.Map({0: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_721_v221_).cf15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_721_v221_).cf16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_721_v221_).cf17).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_722_v222_).cf90))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_722_v222_).cf91))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_723_v223_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_724_v224_).f9).cf15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_724_v224_).f9).cf16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_724_v224_).f9).cf17).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_725_v225_).f12)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_725_v225_).f12)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_725_v225_).f12)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_725_v225_).f12)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_725_v225_).f12)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_725_v225_).f12)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_725_v225_).f12)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_725_v225_).f12)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(False, int(0), False, int(0), False)
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

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf5', Any), ('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(False, int(0), int(0))
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

class D1_DC4(D1, NamedTuple('DC4', [('cf12', Any), ('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {self.cf17.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC6(D2, NamedTuple('DC6', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

class D3_DC7(D3, NamedTuple('DC7', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC9(_dafny.Map({}), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC9(D4, NamedTuple('DC9', [('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf24', Any), ('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC8(D4, NamedTuple('DC8', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13(False, False)
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

class D5_DC13(D5, NamedTuple('DC13', [('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(_dafny.Map({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)

class D6_DC16(D6, NamedTuple('DC16', [('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC19(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)

class D7_DC19(D7, NamedTuple('DC19', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)

class D8_DC20(D8, NamedTuple('DC20', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC22(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)

class D9_DC22(D9, NamedTuple('DC22', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC23(D9, NamedTuple('DC23', [('cf45', Any), ('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC21(D9, NamedTuple('DC21', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf49 == __o.cf49
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

class D10_DC26(D10, NamedTuple('DC26', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC28(False, int(0), int(0), False, False)
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

class D11_DC28(D11, NamedTuple('DC28', [('cf52', Any), ('cf53', Any), ('cf54', Any), ('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28({_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC29(D11, NamedTuple('DC29', [('cf57', Any), ('cf58', Any), ('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [('cf61', Any), ('cf62', Any), ('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf61 == __o.cf61 and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC31(D11, NamedTuple('DC31', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC33(False, False, _dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)

class D12_DC33(D12, NamedTuple('DC33', [('cf67', Any), ('cf68', Any), ('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf67)}, {_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf67 == __o.cf67 and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC32(D12, NamedTuple('DC32', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC35()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)

class D13_DC35(D13, NamedTuple('DC35', [])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35)
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC34(D13, NamedTuple('DC34', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC37(D9.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D14_DC37)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D14_DC38)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)

class D14_DC37(D14, NamedTuple('DC37', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC37({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC37) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC38(D14, NamedTuple('DC38', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC38({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC38) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC36(D14, NamedTuple('DC36', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D15_DC39)

class D15_DC39(D15, NamedTuple('DC39', [('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC39({_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC39) and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC41(_dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D16_DC41)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D16_DC42)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D16_DC40)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D16_DC43)

class D16_DC41(D16, NamedTuple('DC41', [('cf77', Any), ('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC41({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC41) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC42(D16, NamedTuple('DC42', [('cf79', Any), ('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC42({_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC42) and self.cf79 == __o.cf79 and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC40(D16, NamedTuple('DC40', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC40({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC40) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC43(D16, NamedTuple('DC43', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC43({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC43) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC45()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D17_DC45)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D17_DC46)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D17_DC47)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D17_DC44)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D17_DC48)

class D17_DC45(D17, NamedTuple('DC45', [])):
    def __dafnystr__(self) -> str:
        return f'D17.DC45'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC45)
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC46(D17, NamedTuple('DC46', [])):
    def __dafnystr__(self) -> str:
        return f'D17.DC46'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC46)
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC47(D17, NamedTuple('DC47', [('cf83', Any), ('cf84', Any), ('cf85', Any), ('cf86', Any), ('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC47({_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC47) and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85 and self.cf86 == __o.cf86 and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC44(D17, NamedTuple('DC44', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC44({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC44) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC48(D17, NamedTuple('DC48', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC48({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC48) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC50(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D18_DC50)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D18_DC49)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D18_DC51)

class D18_DC50(D18, NamedTuple('DC50', [('cf90', Any), ('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC50({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC50) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC49(D18, NamedTuple('DC49', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC49({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC49) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC51(D18, NamedTuple('DC51', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC51({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC51) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC53(int(0), int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D19_DC53)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D19_DC52)

class D19_DC53(D19, NamedTuple('DC53', [('cf94', Any), ('cf95', Any), ('cf96', Any), ('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC53({_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)}, {_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC53) and self.cf94 == __o.cf94 and self.cf95 == __o.cf95 and self.cf96 == __o.cf96 and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC52(D19, NamedTuple('DC52', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC52({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC52) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D20_DC54)

class D20_DC54(D20, NamedTuple('DC54', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC54({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC54) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC56(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D21_DC56)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D21_DC55)

class D21_DC56(D21, NamedTuple('DC56', [('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC56({_dafny.string_of(self.cf100)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC56) and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC55(D21, NamedTuple('DC55', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC55({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC55) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC58(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D22_DC58)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D22_DC59)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D22_DC57)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D22_DC60)

class D22_DC58(D22, NamedTuple('DC58', [('cf102', Any), ('cf103', Any), ('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC58({self.cf102.VerbatimString(True)}, {_dafny.string_of(self.cf103)}, {_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC58) and self.cf102 == __o.cf102 and self.cf103 == __o.cf103 and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC59(D22, NamedTuple('DC59', [('cf105', Any), ('cf106', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC59({_dafny.string_of(self.cf105)}, {_dafny.string_of(self.cf106)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC59) and self.cf105 == __o.cf105 and self.cf106 == __o.cf106
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC57(D22, NamedTuple('DC57', [('cf101', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC57({_dafny.string_of(self.cf101)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC57) and self.cf101 == __o.cf101
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC60(D22, NamedTuple('DC60', [('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC60({_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC60) and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC62(_dafny.Array(None, 0), _dafny.Map({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D23_DC62)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D23_DC61)

class D23_DC62(D23, NamedTuple('DC62', [('cf109', Any), ('cf110', Any), ('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC62({_dafny.string_of(self.cf109)}, {_dafny.string_of(self.cf110)}, {_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC62) and self.cf109 == __o.cf109 and self.cf110 == __o.cf110 and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC61(D23, NamedTuple('DC61', [('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC61({_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC61) and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC64(False, None, D18.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D24_DC64)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D24_DC63)

class D24_DC64(D24, NamedTuple('DC64', [('cf113', Any), ('cf114', Any), ('cf115', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC64({_dafny.string_of(self.cf113)}, {_dafny.string_of(self.cf114)}, {_dafny.string_of(self.cf115)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC64) and self.cf113 == __o.cf113 and self.cf114 == __o.cf114 and self.cf115 == __o.cf115
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC63(D24, NamedTuple('DC63', [('cf112', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC63({_dafny.string_of(self.cf112)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC63) and self.cf112 == __o.cf112
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D25_DC65)

class D25_DC65(D25, NamedTuple('DC65', [('cf116', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC65({_dafny.string_of(self.cf116)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC65) and self.cf116 == __o.cf116
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC67(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D26_DC67)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D26_DC66)

class D26_DC67(D26, NamedTuple('DC67', [('cf118', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC67({_dafny.string_of(self.cf118)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC67) and self.cf118 == __o.cf118
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC66(D26, NamedTuple('DC66', [('cf117', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC66({_dafny.string_of(self.cf117)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC66) and self.cf117 == __o.cf117
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D27_DC68)

class D27_DC68(D27, NamedTuple('DC68', [('cf119', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC68({_dafny.string_of(self.cf119)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC68) and self.cf119 == __o.cf119
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D28_DC69)

class D28_DC69(D28, NamedTuple('DC69', [('cf120', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC69({_dafny.string_of(self.cf120)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC69) and self.cf120 == __o.cf120
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m3(self, p0, globalState):
        pass


class T1:
    pass

class T2:
    pass
    @property
    def f18(self):
        return self._f18
    @f18.setter
    def f18(self, value):
        self._f18 = value

class GlobalState:
    def  __init__(self):
        self.f0: _dafny.Set = _dafny.Set({})
        self._f1: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1):
        (self).f0 = f0
        (self)._f1 = f1

    @property
    def f1(self):
        return self._f1

class C0:
    def  __init__(self):
        self.f26: _dafny.Set = _dafny.Set({})
        self._f27: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f26, f27):
        (self).f26 = f26
        (self)._f27 = f27

    def fm36(self, p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([(self).f27, (len(_dafny.SeqWithoutIsStrInference([(0) - ((0) - ((self).f27)) for d_762_i0_ in range(default__.abs(788))]))) + (len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([True, False]))})))])

    @property
    def f27(self):
        return self._f27

class C1(T1, T0):
    def  __init__(self):
        self._f4: bool = False
        self._f5: int = int(0)
        self._f15: int = int(0)
        self._f16: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    def ctor__(self, f15, f16, f4, f5):
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f4 = f4
        (self)._f5 = f5

    def fm0(self, globalState):
        return not((self).f4)

    def m3(self, p0, globalState):
        d_763_i0_: int
        d_763_i0_ = 0
        with _dafny.label("0"):
            while (self).f16:
                with _dafny.c_label("0"):
                    if (d_763_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_763_i0_ = (d_763_i0_) + (1)
                    d_764_v0_: _dafny.Seq
                    d_764_v0_ = _dafny.SeqWithoutIsStrInference([True, (self).f4, (self).f4])
                    d_765_v1_: _dafny.Map
                    d_765_v1_ = _dafny.Map({d_764_v0_: True})
                    d_766_v2_: D1
                    d_766_v2_ = D1_DC3(default__.fm45((self).f15, globalState))
                    d_767_v3_: _dafny.Seq
                    d_767_v3_ = _dafny.SeqWithoutIsStrInference([(self).f5])
                    d_765_v1_ = _dafny.Map({((d_766_v2_).cf11).set(default__.safeIndex(len(_dafny.Map({(self).f5: d_767_v3_})), len((d_766_v2_).cf11)), (self).f16): ((self).f5) != (p0)})
                    d_768_v4_: _dafny.Set
                    d_768_v4_ = _dafny.Set({p0})
                    d_769_v5_: _dafny.Array
                    nw116_ = _dafny.Array(None, 9)
                    nw116_[int(0)] = (self).f15
                    nw116_[int(1)] = (self).f15
                    nw116_[int(2)] = p0
                    nw116_[int(3)] = p0
                    nw116_[int(4)] = len(default__.fm46(globalState))
                    nw116_[int(5)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahpmcvwc")))
                    nw116_[int(6)] = len(d_768_v4_)
                    nw116_[int(7)] = len(d_764_v0_)
                    nw116_[int(8)] = (self).f5
                    d_769_v5_ = nw116_
                    d_770_v6_: _dafny.Map
                    d_770_v6_ = _dafny.Map({(self).f5: d_769_v5_})
                    d_771_v7_: D9
                    d_771_v7_ = D9_DC21(d_770_v6_)
                    d_772_v8_: D9
                    d_772_v8_ = D9_DC25(d_771_v7_)
                    d_772_v8_ = (d_772_v8_ if False else d_772_v8_)
                    d_773_v9_: int
                    d_773_v9_ = 691
                    d_773_v9_ = -190
                    d_774_v10_: _dafny.Map
                    d_774_v10_ = _dafny.Map({(self).f4: d_767_v3_})
                    d_775_v11_: _dafny.MultiSet
                    d_775_v11_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "luv"))])
                    d_776_v12_: _dafny.Seq
                    d_776_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdpxux"))
                    d_777_v13_: _dafny.Seq
                    d_777_v13_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "muubgj")), d_776_v12_])
                    d_774_v10_ = (d_774_v10_).set((_dafny.MultiSet(d_777_v13_)).issubset(d_775_v11_), d_767_v3_)
                    pass
            pass
        d_778_v14_: _dafny.Array
        nw117_ = _dafny.Array(int(0), 24)
        d_778_v14_ = nw117_
        d_779_v15_: _dafny.Map
        d_779_v15_ = _dafny.Map({d_778_v14_: (self).f16})
        d_780_v16_: _dafny.Seq
        d_780_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sq"))
        d_781_v17_: D1
        d_781_v17_ = D1_DC5((self).f16, p0, d_780_v16_)
        d_782_v18_: _dafny.Map
        d_782_v18_ = _dafny.Map({(self).f4: (d_781_v17_).cf16})
        d_783_v19_: _dafny.Map
        d_783_v19_ = _dafny.Map({len(d_782_v18_): p0})
        d_784_v20_: _dafny.Set
        d_784_v20_ = _dafny.Set({d_783_v19_})
        d_785_v21_: _dafny.Set
        d_785_v21_ = d_784_v20_
        d_786_v22_: _dafny.MultiSet
        d_786_v22_ = _dafny.MultiSet([(self).f4])
        d_787_v23_: _dafny.Map
        d_787_v23_ = _dafny.Map({(self).f4: d_786_v22_})
        d_788_v24_: C0
        nw118_ = C0()
        nw118_.ctor__(d_785_v21_, len(d_787_v23_))
        d_788_v24_ = nw118_
        rhs48_ = d_779_v15_
        rhs49_ = (d_788_v24_ if (self).f4 else d_788_v24_)
        d_779_v15_ = rhs48_
        d_788_v24_ = rhs49_
        d_789_v25_: _dafny.Seq
        d_789_v25_ = _dafny.SeqWithoutIsStrInference([d_786_v22_, d_786_v22_, _dafny.MultiSet([(self).f4, True, (self).f16]), d_786_v22_, d_786_v22_])
        d_790_v26_: _dafny.Map
        d_790_v26_ = _dafny.Map({not((self).f4): (self).f4})
        d_791_v27_: _dafny.Seq
        d_791_v27_ = _dafny.SeqWithoutIsStrInference([(self).f16])
        d_792_v28_: _dafny.Array
        nw119_ = _dafny.Array(None, 16)
        nw119_[int(0)] = (default__.fm47((self).f4, globalState)) | (_dafny.MultiSet([(self).f4]))
        nw119_[int(1)] = d_786_v22_
        nw119_[int(2)] = d_786_v22_
        nw119_[int(3)] = d_786_v22_
        nw119_[int(4)] = d_786_v22_
        nw119_[int(5)] = _dafny.MultiSet([(self).f16, False, (self).f16, (self).f16])
        nw119_[int(6)] = default__.fm47((self).f4, globalState)
        nw119_[int(7)] = (d_789_v25_)[default__.safeIndex((self).f5, len(d_789_v25_))]
        nw119_[int(8)] = (d_786_v22_).set((self).f16, default__.abs((d_788_v24_).f27))
        nw119_[int(9)] = (d_786_v22_).set(((d_790_v26_)[(self).f4] if ((self).f4) in (d_790_v26_) else (self).f4), default__.abs(p0))
        nw119_[int(10)] = default__.fm47((self).f4, globalState)
        nw119_[int(11)] = _dafny.MultiSet((d_791_v27_) + (d_791_v27_))
        nw119_[int(12)] = d_786_v22_
        nw119_[int(13)] = d_786_v22_
        nw119_[int(14)] = d_786_v22_
        nw119_[int(15)] = (d_786_v22_) - (d_786_v22_)
        d_792_v28_ = nw119_
        index53_ = default__.safeIndex(688, (d_792_v28_).length(0))
        (d_792_v28_)[index53_] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not((d_781_v17_).cf15)]))
        d_793_v29_: int
        d_793_v29_ = 911
        d_794_v30_: _dafny.Map
        d_794_v30_ = _dafny.Map({811: not((self).f4)})
        d_795_v31_: _dafny.Set
        d_795_v31_ = _dafny.Set({not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srygicbqe"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxeqsitj")))), (self).f4, (self).f16, ((d_794_v30_)[(d_788_v24_).f27] if ((d_788_v24_).f27) in (d_794_v30_) else (self).f4), (self).f16})
        d_796_v32_: _dafny.Seq
        d_796_v32_ = _dafny.SeqWithoutIsStrInference([(self).f15, d_793_v29_, p0])
        d_797_v33_: _dafny.MultiSet
        d_797_v33_ = _dafny.MultiSet([len(d_791_v27_), len(d_796_v32_)])
        index54_ = default__.safeIndex(688, (d_792_v28_).length(0))
        rhs50_ = d_795_v31_
        rhs51_ = d_786_v22_
        rhs52_ = default__.safeModuloInt((self).f15, len((default__.fm48((d_788_v24_).f27, globalState)) | (d_795_v31_)))
        rhs53_ = (0) - ((default__.fm49(d_780_v16_, (0) - (default__.fm49(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgjpp")), (self).f5, False, (self).fm0(globalState), globalState)), False, (self).f16, globalState) if ((0) - (d_793_v29_)) not in (d_797_v33_) else (p0) - (596)))
        lhs12_ = globalState
        lhs13_ = d_792_v28_
        lhs14_ = default__.safeIndex(688, (d_792_v28_).length(0))
        lhs12_.f0 = rhs50_
        lhs13_[lhs14_] = rhs51_
        d_793_v29_ = rhs52_
        d_793_v29_ = rhs53_
        d_798_v34_: bool
        d_798_v34_ = True
        d_798_v34_ = (self).f4
        d_799_i1_: int
        d_799_i1_ = 0
        with _dafny.label("1"):
            while False:
                with _dafny.c_label("1"):
                    if (d_799_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_799_i1_ = (d_799_i1_) + (1)
                    if (self).fm0(globalState):
                        d_793_v29_ = (0) - ((d_796_v32_)[default__.safeIndex((d_788_v24_).f27, len(d_796_v32_))])
                        d_800_v35_: str
                        d_800_v35_ = _dafny.CodePoint('h')
                        d_800_v35_ = d_800_v35_
                        d_801_v36_: _dafny.Map
                        d_801_v36_ = _dafny.Map({(d_788_v24_).f27: d_782_v18_})
                        d_802_v37_: _dafny.Set
                        d_802_v37_ = _dafny.Set({(d_788_v24_).f27, (d_788_v24_).f27, 600, (d_788_v24_).f27})
                        rhs54_ = d_778_v14_
                        rhs55_ = (self).f16
                        rhs56_ = (-853 if (d_798_v34_ if (self).f4 else d_798_v34_) else (d_788_v24_).f27)
                        rhs57_ = not (not((d_802_v37_).issubset(_dafny.Set({(self).f15, d_793_v29_, len(d_796_v32_), (0) - ((self).f15), (_dafny.MultiSet([_dafny.Map({(self).f4: 279}), d_782_v18_, ((d_801_v36_)[len(d_780_v16_)] if (len(d_780_v16_)) in (d_801_v36_) else d_782_v18_)])).cardinality})))) or ((self).f4)
                        rhs58_ = ((d_792_v28_)[default__.safeIndex(688, (d_792_v28_).length(0))]).intersection(d_786_v22_)
                        d_778_v14_ = rhs54_
                        d_798_v34_ = rhs55_
                        d_793_v29_ = rhs56_
                        d_798_v34_ = rhs57_
                        d_786_v22_ = rhs58_
                        d_791_v27_ = default__.fm45((d_788_v24_).f27, globalState)
                        d_798_v34_ = not((self).f4)
                    elif True:
                        d_803_v38_: str
                        d_803_v38_ = _dafny.CodePoint('y')
                        d_793_v29_ = len((d_780_v16_).set(default__.safeIndex(len(d_796_v32_), len(d_780_v16_)), (d_803_v38_ if d_798_v34_ else d_803_v38_)))
                        d_803_v38_ = default__.fm50(globalState)
                        d_804_v39_: _dafny.Map
                        d_804_v39_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_803_v38_ for d_805_i2_ in range(default__.abs(652))]): (0) - (len(d_780_v16_))})
                        index55_ = default__.safeIndex(379, (d_778_v14_).length(0))
                        (d_778_v14_)[index55_] = len(d_804_v39_)
                        d_806_v40_: _dafny.Set
                        d_806_v40_ = _dafny.Set({740, (d_788_v24_).f27, (self).f5})
                        index56_ = default__.safeIndex(379, (d_778_v14_).length(0))
                        (d_778_v14_)[index56_] = ((d_783_v19_)[len(d_806_v40_)] if (len(d_806_v40_)) in (d_783_v19_) else (0) - (((self).f15) * (984)))
                        d_807_v41_: bool
                        out10_: bool
                        out10_ = default__.m0((self).f4, len((_dafny.SeqWithoutIsStrInference([(self).f5])) + ((_dafny.SeqWithoutIsStrInference([p0, p0])).set(default__.safeIndex(184, len(_dafny.SeqWithoutIsStrInference([p0, p0]))), p0))), globalState)
                        d_807_v41_ = out10_
                        d_808_v42_: D4
                        d_808_v42_ = D4_DC10(not(d_807_v41_), d_807_v41_, d_791_v27_)
                        d_809_v43_: D4
                        d_809_v43_ = D4_DC11(d_808_v42_)
                        d_810_v44_: _dafny.Array
                        nw120_ = _dafny.Array(None, 18)
                        nw120_[int(0)] = d_803_v38_
                        nw120_[int(1)] = d_803_v38_
                        nw120_[int(2)] = d_803_v38_
                        nw120_[int(3)] = d_803_v38_
                        nw120_[int(4)] = d_803_v38_
                        nw120_[int(5)] = d_803_v38_
                        nw120_[int(6)] = d_803_v38_
                        nw120_[int(7)] = d_803_v38_
                        nw120_[int(8)] = _dafny.CodePoint('r')
                        nw120_[int(9)] = d_803_v38_
                        nw120_[int(10)] = d_803_v38_
                        nw120_[int(11)] = d_803_v38_
                        nw120_[int(12)] = d_803_v38_
                        nw120_[int(13)] = (d_780_v16_)[default__.safeIndex(default__.fm49((d_780_v16_).set(default__.safeIndex(len(d_791_v27_), len(d_780_v16_)), _dafny.CodePoint('p')), len(d_780_v16_), (self).f4, True, globalState), len(d_780_v16_))]
                        nw120_[int(14)] = d_803_v38_
                        nw120_[int(15)] = d_803_v38_
                        nw120_[int(16)] = d_803_v38_
                        nw120_[int(17)] = _dafny.CodePoint('q')
                        d_810_v44_ = nw120_
                        d_811_v45_: _dafny.Map
                        d_811_v45_ = _dafny.Map({d_809_v43_: d_810_v44_})
                        d_812_v46_: _dafny.Array
                        def lambda26_(d_813_v24_):
                            def lambda27_(d_814_i3_):
                                return (d_814_i3_) + ((d_813_v24_).f27)

                            return lambda27_

                        init13_ = lambda26_(d_788_v24_)
                        nw121_ = _dafny.Array(None, 15)
                        for i0_13_ in range(nw121_.length(0)):
                            nw121_[i0_13_] = init13_(i0_13_)
                        d_812_v46_ = nw121_
                        rhs59_ = (d_811_v45_) | (d_811_v45_)
                        rhs60_ = d_812_v46_
                        d_811_v45_ = rhs59_
                        d_812_v46_ = rhs60_
                    d_793_v29_ = default__.safeDivisionInt(d_793_v29_, (self).f15)
                    d_793_v29_ = len(((d_791_v27_).set(default__.safeIndex((d_788_v24_).f27, len(d_791_v27_)), (self).f4)) + ((D4_DC10((self).fm0(globalState), (self).f4, default__.fm45((d_788_v24_).f27, globalState))).cf26))
                    rhs61_ = d_778_v14_
                    rhs62_ = (d_788_v24_).f27
                    d_778_v14_ = rhs61_
                    d_793_v29_ = rhs62_
                    pass
            pass
        d_815_v47_: _dafny.Array
        nw122_ = _dafny.Array(False, 5)
        d_815_v47_ = nw122_
        d_816_v48_: _dafny.MultiSet
        d_816_v48_ = _dafny.MultiSet([d_815_v47_, d_815_v47_])
        d_817_v49_: D7
        d_817_v49_ = D7_DC18(d_816_v48_)
        d_818_v50_: D4
        d_818_v50_ = D4_DC10(d_798_v34_, (self).f4, d_791_v27_)
        d_819_v51_: _dafny.Map
        d_819_v51_ = _dafny.Map({d_817_v49_: d_818_v50_})
        d_820_v52_: _dafny.MultiSet
        d_820_v52_ = _dafny.MultiSet([(d_819_v51_) | (d_819_v51_), d_819_v51_, d_819_v51_, (_dafny.Map({D7_DC18(d_816_v48_): d_818_v50_})) | (_dafny.Map({d_817_v49_: d_818_v50_}))])
        d_821_v53_: _dafny.MultiSet
        d_821_v53_ = _dafny.MultiSet([d_780_v16_])
        d_822_v54_: _dafny.Seq
        d_822_v54_ = _dafny.SeqWithoutIsStrInference([d_821_v53_])
        d_823_v55_: _dafny.Seq
        d_823_v55_ = _dafny.SeqWithoutIsStrInference([d_780_v16_])
        d_824_v56_: _dafny.Map
        d_824_v56_ = _dafny.Map({(self).f4: d_821_v53_})
        d_825_v57_: _dafny.Set
        d_825_v57_ = _dafny.Set({d_793_v29_})
        d_826_v58_: _dafny.Array
        nw123_ = _dafny.Array(None, 21)
        nw123_[int(0)] = d_821_v53_
        nw123_[int(1)] = d_821_v53_
        nw123_[int(2)] = d_821_v53_
        nw123_[int(3)] = (_dafny.MultiSet([d_780_v16_, d_780_v16_])) - ((d_822_v54_)[default__.safeIndex((self).f5, len(d_822_v54_))])
        nw123_[int(4)] = (d_821_v53_).intersection(d_821_v53_)
        nw123_[int(5)] = (d_821_v53_) | (d_821_v53_)
        nw123_[int(6)] = d_821_v53_
        nw123_[int(7)] = (d_821_v53_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mg")), default__.abs((d_788_v24_).f27))
        nw123_[int(8)] = d_821_v53_
        nw123_[int(9)] = d_821_v53_
        nw123_[int(10)] = (d_821_v53_ if d_798_v34_ else default__.fm51(globalState))
        nw123_[int(11)] = d_821_v53_
        nw123_[int(12)] = _dafny.MultiSet(d_823_v55_)
        nw123_[int(13)] = default__.fm51(globalState)
        nw123_[int(14)] = _dafny.MultiSet([d_780_v16_, d_780_v16_, d_780_v16_])
        nw123_[int(15)] = (d_821_v53_) - (d_821_v53_)
        nw123_[int(16)] = d_821_v53_
        nw123_[int(17)] = (d_821_v53_).intersection(_dafny.MultiSet([d_780_v16_]))
        nw123_[int(18)] = _dafny.MultiSet([d_780_v16_])
        nw123_[int(19)] = d_821_v53_
        nw123_[int(20)] = ((d_824_v56_)[(self).f16] if ((self).f16) in (d_824_v56_) else _dafny.MultiSet([default__.fm52((self).f16, d_825_v57_, len(d_794_v30_), (self).f16, globalState)]))
        d_826_v58_ = nw123_
        index57_ = default__.safeIndex(825, (d_826_v58_).length(0))
        (d_826_v58_)[index57_] = d_821_v53_
        index58_ = default__.safeIndex(825, (d_826_v58_).length(0))
        rhs63_ = d_820_v52_
        rhs64_ = not(d_798_v34_)
        rhs65_ = d_821_v53_
        rhs66_ = d_778_v14_
        rhs67_ = d_818_v50_
        lhs15_ = d_826_v58_
        lhs16_ = default__.safeIndex(825, (d_826_v58_).length(0))
        d_820_v52_ = rhs63_
        d_798_v34_ = rhs64_
        lhs15_[lhs16_] = rhs65_
        d_778_v14_ = rhs66_
        d_818_v50_ = rhs67_


class C2(T2, T0):
    def  __init__(self):
        self._f18: bool = False
        self._f17: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: bool = False
        self._f5: int = int(0)
        self._f28: _dafny.Array = _dafny.Array(None, 0)
        self._f29: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f18(self):
        return self._f18
    @f18.setter
    def f18(self, value):
        self._f18 = value
    @property
    def f17(self):
        return self._f17
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f28, f29, f17, f18, f4, f5):
        (self)._f28 = f28
        (self)._f29 = f29
        (self)._f17 = f17
        (self).f18 = f18
        (self)._f4 = f4
        (self)._f5 = f5

    def fm0(self, globalState):
        return (self.f18) and ((self.f18 if (self).f4 else (self).f4))

    def fm54(self, p0, globalState):
        return _dafny.SeqWithoutIsStrInference([not(self.f18)])

    def m3(self, p0, globalState):
        d_827_v0_: _dafny.Map
        d_827_v0_ = _dafny.Map({self.f18: self.f18})
        d_828_v1_: _dafny.Map
        d_828_v1_ = d_827_v0_
        d_829_v2_: _dafny.Map
        d_829_v2_ = _dafny.Map({p0: (self).f4})
        d_830_v3_: _dafny.Seq
        d_830_v3_ = _dafny.SeqWithoutIsStrInference([self.f18])
        d_831_v4_: _dafny.Array
        nw124_ = _dafny.Array(None, 10)
        nw124_[int(0)] = (p0) * (p0)
        nw124_[int(1)] = len((d_828_v1_))
        nw124_[int(2)] = (self).f5
        nw124_[int(3)] = (self).f5
        nw124_[int(4)] = (len(d_829_v2_)) - (-645)
        nw124_[int(5)] = (self).f5
        nw124_[int(6)] = default__.safeModuloInt(p0, (self).f5)
        nw124_[int(7)] = ((self).f5) - ((0) - (default__.fm55(p0, d_830_v3_, p0, globalState)))
        nw124_[int(8)] = 668
        nw124_[int(9)] = (self).f5
        d_831_v4_ = nw124_
        index59_ = default__.safeIndex(440, (d_831_v4_).length(0))
        (d_831_v4_)[index59_] = default__.safeModuloInt(p0, (self).f5)
        index60_ = default__.safeIndex(440, (d_831_v4_).length(0))
        (d_831_v4_)[index60_] = (self).f5
        index61_ = default__.safeIndex(440, (d_831_v4_).length(0))
        (d_831_v4_)[index61_] = (default__.fm55((d_831_v4_)[default__.safeIndex(440, (d_831_v4_).length(0))], d_830_v3_, p0, globalState)) - (default__.safeModuloInt(p0, (0) - ((self).f5)))
        if not ((self).f4) or (self.f18):
            d_832_v5_: _dafny.Seq
            d_832_v5_ = _dafny.SeqWithoutIsStrInference([p0, p0, len((self).f17)])
            d_832_v5_ = (d_832_v5_) + (d_832_v5_)
            (self).f18 = (self).f4
            d_831_v4_ = d_831_v4_
            d_833_v6_: str
            d_833_v6_ = _dafny.CodePoint('c')
            d_833_v6_ = d_833_v6_
            d_834_v7_: _dafny.Seq
            d_834_v7_ = _dafny.SeqWithoutIsStrInference([d_832_v5_, d_832_v5_])
            d_835_v8_: _dafny.Array
            def lambda28_(d_836_i0_):
                return (_dafny.MultiSet([(self).f4, self.f18, self.f18])) | (_dafny.MultiSet([False, self.f18, (self).f4]))

            init14_ = lambda28_
            nw125_ = _dafny.Array(None, 21)
            for i0_14_ in range(nw125_.length(0)):
                nw125_[i0_14_] = init14_(i0_14_)
            d_835_v8_ = nw125_
            index62_ = default__.safeIndex(440, (d_831_v4_).length(0))
            rhs68_ = (len(((self).f17) + ((self).f17))) * ((d_831_v4_)[default__.safeIndex(440, (d_831_v4_).length(0))])
            rhs69_ = d_834_v7_
            rhs70_ = d_835_v8_
            lhs17_ = d_831_v4_
            lhs18_ = default__.safeIndex(440, (d_831_v4_).length(0))
            lhs17_[lhs18_] = rhs68_
            d_834_v7_ = rhs69_
            d_835_v8_ = rhs70_
        elif True:
            d_837_v9_: _dafny.Set
            d_837_v9_ = _dafny.Set({(d_831_v4_)[default__.safeIndex(440, (d_831_v4_).length(0))], (self).f5, len((self).f17), (d_831_v4_)[default__.safeIndex(440, (d_831_v4_).length(0))]})
            d_838_v10_: C1
            nw126_ = C1()
            nw126_.ctor__(((d_831_v4_)[default__.safeIndex(440, (d_831_v4_).length(0))]) * (753), (self).f4, ((d_831_v4_)[default__.safeIndex(440, (d_831_v4_).length(0))]) not in (d_837_v9_), default__.fm55(len((_dafny.SeqWithoutIsStrInference([False])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([False]))), not(True))), d_830_v3_, (d_831_v4_)[default__.safeIndex(440, (d_831_v4_).length(0))], globalState))
            d_838_v10_ = nw126_
            d_839_v11_: _dafny.Map
            d_839_v11_ = _dafny.Map({65: 733})
            d_839_v11_ = (d_839_v11_).set((d_831_v4_)[default__.safeIndex(440, (d_831_v4_).length(0))], (d_831_v4_)[default__.safeIndex(440, (d_831_v4_).length(0))])
            d_840_v12_: _dafny.Map
            d_840_v12_ = _dafny.Map({(self).f4: (d_831_v4_)[default__.safeIndex(440, (d_831_v4_).length(0))]})
            index63_ = default__.safeIndex(440, (d_831_v4_).length(0))
            (d_831_v4_)[index63_] = default__.fm55((default__.fm55((0) - ((d_831_v4_)[default__.safeIndex(440, (d_831_v4_).length(0))]), _dafny.SeqWithoutIsStrInference([(d_838_v10_).fm0(globalState)]), len((self).f17), globalState)) + (len(d_840_v12_)), (d_830_v3_) + (_dafny.SeqWithoutIsStrInference([self.f18])), len((d_837_v9_).intersection(d_837_v9_)), globalState)
            d_841_v13_: C0
            nw127_ = C0()
            nw127_.ctor__(default__.fm56((self).f4, globalState), p0)
            d_841_v13_ = nw127_
            d_842_v14_: _dafny.Seq
            d_842_v14_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_831_v4_)[default__.safeIndex(440, (d_831_v4_).length(0))]), p0])
            d_843_v15_: _dafny.MultiSet
            d_843_v15_ = _dafny.MultiSet([len(d_842_v14_), (self).f5, p0])
            d_843_v15_ = d_843_v15_
        (self).f18 = (not((self).f4) if self.f18 else self.f18)
        (self).f18 = self.f18
        d_844_v16_: C0
        nw128_ = C0()
        nw128_.ctor__(default__.fm56(self.f18, globalState), p0)
        d_844_v16_ = nw128_

    @property
    def f28(self):
        return self._f28
    @property
    def f29(self):
        return self._f29

class C3(T1, T0):
    def  __init__(self):
        self._f4: bool = False
        self._f5: int = int(0)
        self._f15: int = int(0)
        self._f16: bool = False
        self._f25: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    def ctor__(self, f25, f15, f16, f4, f5):
        (self)._f25 = f25
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f4 = f4
        (self)._f5 = f5

    def fm0(self, globalState):
        return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lg")))) == (490)

    def fm35(self, p0, p1, p2, globalState):
        return len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqfkpktcs")))

    def m3(self, p0, globalState):
        d_845_v0_: _dafny.Array
        nw129_ = _dafny.Array(False, 17)
        d_845_v0_ = nw129_
        d_846_v1_: D4
        d_846_v1_ = D4_DC8(d_845_v0_)
        d_847_v2_: _dafny.MultiSet
        d_847_v2_ = _dafny.MultiSet([D4_DC8(d_845_v0_), d_846_v1_, d_846_v1_, d_846_v1_, d_846_v1_])
        hi1_ = 734
        for d_848_i0_ in range((d_847_v2_).cardinality, hi1_):
            d_849_v3_: bool
            d_849_v3_ = False
            d_849_v3_ = (self).f4
            d_849_v3_ = True
            d_845_v0_ = d_845_v0_
            d_850_v4_: _dafny.Array
            def lambda29_(d_851_i1_):
                return (d_851_i1_) + (293)

            init15_ = lambda29_
            nw130_ = _dafny.Array(None, 12)
            for i0_15_ in range(nw130_.length(0)):
                nw130_[i0_15_] = init15_(i0_15_)
            d_850_v4_ = nw130_
            d_852_v5_: _dafny.Seq
            d_852_v5_ = _dafny.SeqWithoutIsStrInference([d_850_v4_, d_850_v4_])
            d_850_v4_ = (d_852_v5_)[default__.safeIndex(d_848_i0_, len(d_852_v5_))]
        hi2_ = (self).f15
        for d_853_i2_ in range((self).f15, hi2_):
            d_854_v6_: _dafny.Map
            d_854_v6_ = _dafny.Map({(self).f15: len(default__.fm37(globalState))})
            d_855_v7_: _dafny.Set
            d_855_v7_ = _dafny.Set({(d_854_v6_).set((self).f15, 927)})
            d_856_v8_: C0
            nw131_ = C0()
            nw131_.ctor__(d_855_v7_, p0)
            d_856_v8_ = nw131_
            d_857_v9_: _dafny.Seq
            d_857_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkspqi"))
            d_858_v10_: _dafny.Set
            d_858_v10_ = _dafny.Set({len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rspby"))) + (d_857_v9_)), d_853_i2_})
            def iife81_():
                coll41_ = _dafny.Set()
                compr_41_: int
                for compr_41_ in _dafny.IntegerRange(790, -133):
                    d_859_v11_: int = compr_41_
                    if ((790) <= (d_859_v11_)) and ((d_859_v11_) < (-133)):
                        coll41_ = coll41_.union(_dafny.Set([default__.safeModuloInt(d_859_v11_, len(_dafny.SeqWithoutIsStrInference([(d_856_v8_).f27, 805, len(_dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmlptudj"))).set(default__.safeIndex((d_856_v8_).f27, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmlptudj")))), (self).f25))])), d_853_i2_, (self).f15])))]))
                return _dafny.Set(coll41_)
            d_858_v10_ = iife81_()
            
            d_860_v12_: int
            d_860_v12_ = -598
            d_861_v13_: _dafny.Map
            d_861_v13_ = _dafny.Map({((self).f16 if (self).f16 else (self).f4): d_853_i2_})
            d_862_v14_: _dafny.Set
            d_862_v14_ = _dafny.Set({(self).f4, False, (self).f4, (self).f4, True})
            d_860_v12_ = ((d_861_v13_)[(self).f16] if ((self).f16) in (d_861_v13_) else len(d_862_v14_))
            d_863_v15_: _dafny.Map
            d_863_v15_ = _dafny.Map({p0: _dafny.Map({(self).f15: (self).f15})})
            d_860_v12_ = len(d_863_v15_)
        d_864_v16_: _dafny.Set
        d_864_v16_ = _dafny.Set({(self).f5})
        d_865_v17_: _dafny.Map
        d_865_v17_ = _dafny.Map({(self).f4: len(d_864_v16_)})
        d_866_v18_: _dafny.Map
        d_866_v18_ = _dafny.Map({(d_865_v17_).set((self).f4, (0) - ((self).f15)): 481})
        d_867_v19_: _dafny.Seq
        d_867_v19_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c')])
        d_868_v20_: _dafny.Seq
        d_868_v20_ = _dafny.SeqWithoutIsStrInference([d_867_v19_, d_867_v19_])
        d_869_v21_: D6
        d_869_v21_ = D6_DC16(d_866_v18_, len(d_868_v20_))
        d_870_v22_: _dafny.Seq
        d_870_v22_ = _dafny.SeqWithoutIsStrInference([-226])
        d_871_v23_: _dafny.Map
        d_871_v23_ = _dafny.Map({(self).f4: d_870_v22_})
        d_872_v24_: _dafny.Map
        d_872_v24_ = _dafny.Map({(self).f16: False})
        d_873_v25_: _dafny.Map
        d_873_v25_ = _dafny.Map({len(((d_871_v23_)[((d_872_v24_)[(self).f4] if ((self).f4) in (d_872_v24_) else (self).f4)] if (((d_872_v24_)[(self).f4] if ((self).f4) in (d_872_v24_) else (self).f4)) in (d_871_v23_) else d_870_v22_)): d_869_v21_})
        d_874_v27_: _dafny.MultiSet
        def iife82_():
            coll42_ = _dafny.Map()
            compr_42_: int
            for compr_42_ in _dafny.IntegerRange(264, 576):
                d_875_v26_: int = compr_42_
                if ((264) <= (d_875_v26_)) and ((d_875_v26_) < (576)):
                    coll42_[default__.safeModuloInt(d_875_v26_, 112)] = d_869_v21_
            return _dafny.Map(coll42_)
        d_874_v27_ = _dafny.MultiSet([_dafny.Map({933: d_869_v21_}), d_873_v25_, iife82_()
        ])
        if (_dafny.MultiSet([d_873_v25_, _dafny.Map({782: d_869_v21_})])).ispropersubset(d_874_v27_):
            d_876_v28_: _dafny.Map
            d_876_v28_ = _dafny.Map({(self).f15: p0})
            d_877_v29_: _dafny.Set
            d_877_v29_ = _dafny.Set({d_876_v28_})
            d_878_v30_: _dafny.Set
            d_878_v30_ = d_877_v29_
            d_879_v31_: C0
            nw132_ = C0()
            nw132_.ctor__(d_878_v30_, p0)
            d_879_v31_ = nw132_
            d_880_v32_: bool
            d_880_v32_ = True
            d_880_v32_ = d_880_v32_
            d_881_v33_: _dafny.Map
            d_881_v33_ = _dafny.Map({(d_879_v31_).f27: d_880_v32_})
            d_881_v33_ = d_881_v33_
            d_880_v32_ = not((self).f4)
            d_882_v34_: int
            d_882_v34_ = 455
            rhs71_ = (-964) <= ((self).f15)
            rhs72_ = (self).f5
            d_880_v32_ = rhs71_
            d_882_v34_ = rhs72_
        elif True:
            d_883_v35_: _dafny.Set
            d_883_v35_ = default__.fm38((self).f16, (self).f4, (self).f5, (self).f16, globalState)
            d_884_v36_: C0
            nw133_ = C0()
            nw133_.ctor__(d_883_v35_, (self).f15)
            d_884_v36_ = nw133_
            d_885_v37_: _dafny.Array
            nw134_ = _dafny.Array(int(0), 17)
            d_885_v37_ = nw134_
            d_886_v38_: _dafny.Seq
            d_886_v38_ = _dafny.SeqWithoutIsStrInference([(self).f16, (self).f16])
            index64_ = default__.safeIndex(256, (d_885_v37_).length(0))
            (d_885_v37_)[index64_] = (len(d_886_v38_)) + ((d_884_v36_).f27)
            d_887_v39_: str
            d_887_v39_ = _dafny.CodePoint('h')
            index65_ = default__.safeIndex(551, (d_845_v0_).length(0))
            (d_845_v0_)[index65_] = (self).f16
            d_888_v40_: _dafny.MultiSet
            d_888_v40_ = _dafny.MultiSet([(self).f4])
            d_889_v41_: _dafny.MultiSet
            d_889_v41_ = _dafny.MultiSet([d_888_v40_])
            d_890_v42_: D4
            d_890_v42_ = D4_DC10((self).f4, (self).f16, d_886_v38_)
            d_891_v43_: _dafny.Map
            d_891_v43_ = _dafny.Map({d_889_v41_: _dafny.Map({d_885_v37_: _dafny.MultiSet([d_890_v42_])})})
            d_892_v44_: _dafny.Map
            def iife83_(_pat_let20_0):
                def iife84_(d_893_dt__update__tmp_h0_):
                    def iife85_(_pat_let21_0):
                        def iife86_(d_894_dt__update_hcf24_h0_):
                            return D4_DC10(d_894_dt__update_hcf24_h0_, (d_893_dt__update__tmp_h0_).cf25, (d_893_dt__update__tmp_h0_).cf26)
                        return iife86_(_pat_let21_0)
                    return iife85_(False)
                return iife84_(_pat_let20_0)
            d_892_v44_ = _dafny.Map({d_885_v37_: _dafny.MultiSet([D4_DC10(not(False), not((self).f16), d_886_v38_), iife83_(d_890_v42_), d_890_v42_, d_890_v42_, d_890_v42_])})
            index66_ = default__.safeIndex(256, (d_885_v37_).length(0))
            index67_ = default__.safeIndex(551, (d_845_v0_).length(0))
            rhs73_ = len(((d_891_v43_)[(d_889_v41_) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_888_v40_])))] if ((d_889_v41_) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_888_v40_])))) in (d_891_v43_) else d_892_v44_))
            rhs74_ = _dafny.SeqWithoutIsStrInference([(self).f5 for d_895_i3_ in range(default__.abs(736))])
            rhs75_ = (self).f25
            rhs76_ = (self).f4
            lhs19_ = d_885_v37_
            lhs20_ = default__.safeIndex(256, (d_885_v37_).length(0))
            lhs21_ = d_845_v0_
            lhs22_ = default__.safeIndex(551, (d_845_v0_).length(0))
            lhs19_[lhs20_] = rhs73_
            d_870_v22_ = rhs74_
            d_887_v39_ = rhs75_
            lhs21_[lhs22_] = rhs76_
            d_896_v45_: _dafny.Array
            def lambda30_(d_897_v24_):
                def lambda31_(d_898_i4_):
                    return ((d_897_v24_)[(self).f4] if ((self).f4) in (d_897_v24_) else (self).f16)

                return lambda31_

            init16_ = lambda30_(d_872_v24_)
            nw135_ = _dafny.Array(None, 2)
            for i0_16_ in range(nw135_.length(0)):
                nw135_[i0_16_] = init16_(i0_16_)
            d_896_v45_ = nw135_
            nw136_ = _dafny.Array(False, 19)
            d_896_v45_ = nw136_
            index68_ = default__.safeIndex(551, (d_845_v0_).length(0))
            (d_845_v0_)[index68_] = not((self).fm0(globalState))
            index69_ = default__.safeIndex(256, (d_885_v37_).length(0))
            (d_885_v37_)[index69_] = (self).fm35(((default__.fm39(len(d_865_v17_), globalState)).cardinality) != (p0), (self).f16, p0, globalState)
        d_899_v46_: int
        d_899_v46_ = -144
        d_900_v47_: _dafny.Array
        nw137_ = _dafny.Array(int(0), 24)
        d_900_v47_ = nw137_
        d_901_v48_: _dafny.Set
        d_901_v48_ = _dafny.Set({(self).f4})
        d_902_v49_: D9
        d_902_v49_ = D9_DC23((self).f16, (self).fm35((self).f16, (D5_DC14((self).f4, _dafny.SeqWithoutIsStrInference([(self).f16]), d_900_v47_, len(d_901_v48_), (self).f15)).cf31, len(d_867_v19_), globalState), False)
        d_899_v46_ = (-761 if not((d_902_v49_).cf47) else p0)
        d_903_v50_: _dafny.Map
        d_903_v50_ = _dafny.Map({len(d_867_v19_): (0) - (p0)})
        d_904_v51_: _dafny.Seq
        d_904_v51_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_903_v50_).set(634, p0), d_903_v50_})])
        d_905_v52_: _dafny.Set
        d_905_v52_ = (d_904_v51_)[default__.safeIndex(p0, len(d_904_v51_))]
        d_906_v53_: C0
        nw138_ = C0()
        nw138_.ctor__(d_905_v52_, -723)
        d_906_v53_ = nw138_
        d_907_v54_: _dafny.Set
        d_907_v54_ = _dafny.Set({d_903_v50_})
        d_908_v55_: C0
        nw139_ = C0()
        nw139_.ctor__(d_907_v54_, (self).f5)
        d_908_v55_ = nw139_

    def m17(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: D0 = D0.default()()
        r2: bool = False
        hi3_ = p1
        for d_909_i0_ in range((self).f5, hi3_):
            d_910_v0_: _dafny.Seq
            d_910_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
            d_911_v1_: _dafny.MultiSet
            d_911_v1_ = _dafny.MultiSet([d_910_v0_, d_910_v0_, d_910_v0_, (d_910_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjnchtlc"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qja"))])
            d_912_v2_: int
            d_912_v2_ = 305
            d_913_v3_: _dafny.MultiSet
            d_913_v3_ = _dafny.MultiSet([p3, p2])
            d_914_v4_: _dafny.Seq
            d_914_v4_ = _dafny.SeqWithoutIsStrInference([True])
            d_915_v5_: _dafny.Seq
            d_915_v5_ = _dafny.SeqWithoutIsStrInference([(d_913_v3_).set(p2, default__.abs(d_909_i0_)), (d_913_v3_ if True else d_913_v3_), d_913_v3_, _dafny.MultiSet([(self).f15, (0) - (len(d_914_v4_))]), d_913_v3_])
            rhs77_ = (d_911_v1_) | (default__.fm40(p3, globalState))
            rhs78_ = -717
            rhs79_ = _dafny.SeqWithoutIsStrInference([d_913_v3_])
            rhs80_ = ((self).f16) or ((d_910_v0_) < (default__.fm37(globalState)))
            d_911_v1_ = rhs77_
            d_912_v2_ = rhs78_
            d_915_v5_ = rhs79_
            r2 = rhs80_
            d_916_v6_: _dafny.Set
            d_916_v6_ = _dafny.Set({d_912_v2_})
            d_917_v7_: D9
            d_917_v7_ = D9_DC22((self).f4)
            d_918_v8_: C0
            nw140_ = C0()
            nw140_.ctor__(default__.fm41(d_916_v6_, globalState), ((self).fm35((self).f16, (d_917_v7_).cf44, p1, globalState)) + (len(_dafny.Map({(self).f16: ((d_911_v1_)[d_910_v0_] if (d_910_v0_) in (d_911_v1_) else 479)}))))
            d_918_v8_ = nw140_
            d_912_v2_ = default__.safeModuloInt(default__.safeModuloInt(len(d_910_v0_), p0), d_912_v2_)
            r2 = (self).f16
        d_919_v9_: _dafny.Map
        d_919_v9_ = _dafny.Map({(self).f4: (self).f5})
        d_920_v10_: _dafny.Seq
        d_920_v10_ = _dafny.SeqWithoutIsStrInference([False, (self).f4])
        d_921_v11_: _dafny.Seq
        d_921_v11_ = _dafny.SeqWithoutIsStrInference([p2])
        d_922_v12_: _dafny.Map
        d_922_v12_ = _dafny.Map({d_919_v9_: (self).fm35((self).f4, (d_920_v10_)[default__.safeIndex(p0, len(d_920_v10_))], len(d_921_v11_), globalState)})
        d_923_v13_: _dafny.Seq
        d_923_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sndnas"))
        d_924_v14_: D6
        d_924_v14_ = D6_DC16((D6_DC16(d_922_v12_, len(_dafny.SeqWithoutIsStrInference([(self).f16])))).cf37, len(d_923_v13_))
        d_925_v15_: D6
        d_925_v15_ = D6_DC17(d_924_v14_)
        source14_ = (d_925_v15_ if (703) < ((0) - (p0)) else default__.fm42(globalState))
        if source14_.is_DC16:
            d_926___mcc_h0_ = source14_.cf37
            d_927___mcc_h1_ = source14_.cf38
            d_928_cf38_ = d_927___mcc_h1_
            d_929_cf37_ = d_926___mcc_h0_
            r2 = (self).f16
            if (self).fm0(globalState):
                d_930_v16_: _dafny.Array
                nw141_ = _dafny.Array(_dafny.Map({}), 7)
                d_930_v16_ = nw141_
                d_931_v17_: _dafny.MultiSet
                d_931_v17_ = _dafny.MultiSet([(self).f16, True, (self).f4])
                index70_ = default__.safeIndex(22, (d_930_v16_).length(0))
                (d_930_v16_)[index70_] = _dafny.Map({p0: d_931_v17_})
                index71_ = default__.safeIndex(22, (d_930_v16_).length(0))
                rhs81_ = not(((d_928_cf38_) <= (950) if (self).f16 else (self).f16))
                rhs82_ = (self).f4
                rhs83_ = default__.fm43(p1, (d_921_v11_).set(default__.safeIndex(p3, len(d_921_v11_)), len(d_919_v9_)), p0, globalState)
                lhs23_ = d_930_v16_
                lhs24_ = default__.safeIndex(22, (d_930_v16_).length(0))
                r2 = rhs81_
                r2 = rhs82_
                lhs23_[lhs24_] = rhs83_
                d_932_v18_: D0
                d_932_v18_ = D0_DC0(not((self).f16), p0, (self).f16, p1, (self).f4)
                d_921_v11_ = default__.fm44(132, d_932_v18_, (self).f16, (self).f4, globalState)
                r2 = not((self).f16)
                d_933_v19_: D9
                d_933_v19_ = D9_DC24((self).f15)
                d_928_cf38_ = (d_933_v19_).cf48
                d_934_v20_: _dafny.Set
                d_934_v20_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlwkseqa"))})
                d_935_v21_: D5
                d_935_v21_ = D5_DC12(d_934_v20_)
                d_936_v22_: _dafny.MultiSet
                d_936_v22_ = _dafny.MultiSet([D5_DC12(d_934_v20_), d_935_v21_, d_935_v21_])
                d_937_v23_: _dafny.Array
                nw142_ = _dafny.Array(_dafny.Map({}), 16)
                d_937_v23_ = nw142_
                d_938_v24_: T1
                nw143_ = C1()
                nw143_.ctor__((0) - (p1), not((self).f4), (self).f16, len(d_923_v13_))
                d_938_v24_ = nw143_
                d_939_v25_: _dafny.Map
                d_939_v25_ = _dafny.Map({d_938_v24_: (d_938_v24_).f16})
                index72_ = default__.safeIndex(289, (d_937_v23_).length(0))
                (d_937_v23_)[index72_] = (_dafny.Map({d_938_v24_: (self).f16}) if (self).f4 else d_939_v25_)
                index73_ = default__.safeIndex(289, (d_937_v23_).length(0))
                rhs84_ = d_936_v22_
                rhs85_ = not(((d_938_v24_).f5) < ((0) - ((d_928_cf38_) + ((d_921_v11_)[default__.safeIndex(d_928_cf38_, len(d_921_v11_))]))))
                rhs86_ = (0) - ((len(d_920_v10_)) * (p3))
                rhs87_ = (d_938_v24_).f4
                rhs88_ = (_dafny.Map({d_938_v24_: (self).f16})) | (d_939_v25_)
                lhs25_ = d_937_v23_
                lhs26_ = default__.safeIndex(289, (d_937_v23_).length(0))
                d_936_v22_ = rhs84_
                r2 = rhs85_
                d_928_cf38_ = rhs86_
                r2 = rhs87_
                lhs25_[lhs26_] = rhs88_
            elif True:
                d_940_v26_: _dafny.Set
                d_940_v26_ = default__.fm38((self).f4, (self).f4, p1, (self).f16, globalState)
                d_941_v27_: C0
                nw144_ = C0()
                nw144_.ctor__(d_940_v26_, default__.safeModuloInt((self).f5, d_928_cf38_))
                d_941_v27_ = nw144_
                d_942_v28_: _dafny.Array
                nw145_ = _dafny.Array(False, 12)
                d_942_v28_ = nw145_
                index74_ = default__.safeIndex(0, (d_942_v28_).length(0))
                (d_942_v28_)[index74_] = not((d_920_v10_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptyokebry"))), len(d_920_v10_))])
                index75_ = default__.safeIndex(0, (d_942_v28_).length(0))
                (d_942_v28_)[index75_] = not ((self).f16) or (not ((self).f4) or ((self).f16))
                d_943_v29_: _dafny.Array
                def lambda32_(d_944_v13_):
                    def lambda33_(d_945_i1_):
                        return default__.safeDivisionInt(d_945_i1_, len(d_944_v13_))

                    return lambda33_

                init17_ = lambda32_(d_923_v13_)
                nw146_ = _dafny.Array(None, 1)
                for i0_17_ in range(nw146_.length(0)):
                    nw146_[i0_17_] = init17_(i0_17_)
                d_943_v29_ = nw146_
                d_946_v30_: _dafny.Map
                d_946_v30_ = _dafny.Map({(self).f4: (d_942_v28_)[default__.safeIndex(0, (d_942_v28_).length(0))]})
                index76_ = default__.safeIndex(237, (d_943_v29_).length(0))
                (d_943_v29_)[index76_] = len(d_946_v30_)
                index77_ = default__.safeIndex(237, (d_943_v29_).length(0))
                (d_943_v29_)[index77_] = len(d_923_v13_)
                d_920_v10_ = default__.fm45((-315) * ((d_943_v29_)[default__.safeIndex(237, (d_943_v29_).length(0))]), globalState)
                d_942_v28_ = d_942_v28_
            r2 = (not((self).f4) if (self).f4 else (False if (self).f16 else True))
            d_947_v31_: _dafny.MultiSet
            d_947_v31_ = _dafny.MultiSet([not((self).f16)])
            d_948_v32_: C1
            nw147_ = C1()
            nw147_.ctor__(((d_947_v31_)[(self).f16] if ((self).f16) in (d_947_v31_) else p2), (self).f16, (self).f4, p2)
            d_948_v32_ = nw147_
        elif source14_.is_DC15:
            d_949___mcc_h2_ = source14_.cf36
            d_950_cf36_ = d_949___mcc_h2_
            r2 = (((self).f15) + (p1)) in ((d_950_cf36_) + (d_950_cf36_))
            d_951_v33_: _dafny.Seq
            d_951_v33_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrvipmo")), d_923_v13_, (d_923_v13_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjb"))), d_923_v13_, _dafny.SeqWithoutIsStrInference([(self).f25 for d_952_i2_ in range(default__.abs(153))])])
            d_951_v33_ = (default__.fm53(globalState) if (self).f4 else (d_951_v33_) + (d_951_v33_))
            r2 = (self).f16
            d_953_v34_: _dafny.Array
            nw148_ = _dafny.Array(_dafny.Map({}), 25)
            d_953_v34_ = nw148_
            d_954_v36_: _dafny.Map
            def iife87_():
                coll43_ = _dafny.Map()
                compr_43_: int
                for compr_43_ in _dafny.IntegerRange(156, 907):
                    d_955_v35_: int = compr_43_
                    if ((156) <= (d_955_v35_)) and ((d_955_v35_) < (907)):
                        coll43_[default__.safeDivisionInt(d_955_v35_, p3)] = p2
                return _dafny.Map(coll43_)
            d_954_v36_ = _dafny.Map({iife87_()
            : _dafny.SeqWithoutIsStrInference([(self).f25 for d_956_i3_ in range(default__.abs(-616))])})
            d_957_v37_: _dafny.Map
            d_957_v37_ = _dafny.Map({181: 238})
            index78_ = default__.safeIndex(724, (d_953_v34_).length(0))
            (d_953_v34_)[index78_] = (d_954_v36_).set(d_957_v37_, d_923_v13_)
            d_958_v39_: _dafny.MultiSet
            d_958_v39_ = _dafny.MultiSet([d_957_v37_])
            index79_ = default__.safeIndex(724, (d_953_v34_).length(0))
            def iife88_():
                coll44_ = _dafny.Map()
                compr_44_: _dafny.Map
                for compr_44_ in (d_958_v39_).Elements:
                    d_959_v38_: _dafny.Map = compr_44_
                    if (d_959_v38_) in (d_958_v39_):
                        coll44_[d_959_v38_] = d_923_v13_
                return _dafny.Map(coll44_)
            (d_953_v34_)[index79_] = ((iife88_()
             if (self).f16 else d_954_v36_)) | (d_954_v36_)
        elif True:
            d_960___mcc_h3_ = source14_.cf39
            d_961_cf39_ = d_960___mcc_h3_
            d_962_v40_: bool
            out11_: bool
            out11_ = default__.m0(((self).f4 if (self).f4 else (self).f4), (0) - (p2), globalState)
            d_962_v40_ = out11_
            d_963_v41_: _dafny.Set
            d_963_v41_ = _dafny.Set({(self).f4, (self).f16})
            d_964_v42_: _dafny.MultiSet
            d_964_v42_ = _dafny.MultiSet([d_962_v40_, ((self).f4 if True else d_962_v40_), (d_963_v41_).issubset(d_963_v41_), (self).fm0(globalState), (self).f4])
            d_964_v42_ = (d_964_v42_).intersection((d_964_v42_) | (_dafny.MultiSet([d_962_v40_, d_962_v40_, d_962_v40_])))
            d_965_v43_: D0
            d_965_v43_ = D0_DC1(451, False, p2, (self).f5)
            if (d_965_v43_).cf6:
                r2 = (self).f16
                d_966_v44_: int
                d_966_v44_ = 608
                d_966_v44_ = default__.safeDivisionInt(p0, d_966_v44_)
                d_967_v45_: _dafny.Map
                d_967_v45_ = _dafny.Map({(self).f16: (self).f16})
                d_968_v46_: _dafny.Map
                d_968_v46_ = _dafny.Map({len(d_919_v9_): p3})
                d_969_v47_: D1
                d_969_v47_ = D1_DC4((self).f4, p3, p0)
                d_970_v48_: _dafny.Map
                d_970_v48_ = _dafny.Map({(self).f25: (d_969_v47_).cf14})
                d_971_v49_: _dafny.Array
                nw149_ = _dafny.Array(None, 16)
                nw149_[int(0)] = d_966_v44_
                nw149_[int(1)] = 622
                nw149_[int(2)] = ((d_921_v11_)[default__.safeIndex((self).f15, len(d_921_v11_))] if (self).f16 else p0)
                nw149_[int(3)] = (p0) + (p1)
                nw149_[int(4)] = default__.safeDivisionInt(p1, (self).fm35((self).f16, (self).f16, len(d_920_v10_), globalState))
                nw149_[int(5)] = (d_921_v11_)[default__.safeIndex(len(d_967_v45_), len(d_921_v11_))]
                nw149_[int(6)] = p1
                nw149_[int(7)] = (self).fm35((self).f4, (self).f4, 322, globalState)
                nw149_[int(8)] = p1
                nw149_[int(9)] = (self).fm35(d_962_v40_, d_962_v40_, len(d_968_v46_), globalState)
                nw149_[int(10)] = (len(d_970_v48_)) * (p1)
                nw149_[int(11)] = (0) - (p1)
                nw149_[int(12)] = (self).f5
                nw149_[int(13)] = p0
                nw149_[int(14)] = 871
                nw149_[int(15)] = ((self).f15) - (940)
                d_971_v49_ = nw149_
                index80_ = default__.safeIndex(575, (d_971_v49_).length(0))
                (d_971_v49_)[index80_] = p3
                index81_ = default__.safeIndex(575, (d_971_v49_).length(0))
                (d_971_v49_)[index81_] = p2
                d_971_v49_ = d_971_v49_
                d_972_v50_: _dafny.Array
                def lambda34_(d_973_v41_, d_974_v40_):
                    def lambda35_(d_975_i4_):
                        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_973_v41_, d_973_v41_, d_973_v41_]))).intersection(_dafny.MultiSet([_dafny.Set({d_974_v40_})]))

                    return lambda35_

                init18_ = lambda34_(d_963_v41_, d_962_v40_)
                nw150_ = _dafny.Array(None, 25)
                for i0_18_ in range(nw150_.length(0)):
                    nw150_[i0_18_] = init18_(i0_18_)
                d_972_v50_ = nw150_
                index82_ = default__.safeIndex(919, (d_972_v50_).length(0))
                (d_972_v50_)[index82_] = _dafny.MultiSet([d_963_v41_, d_963_v41_, d_963_v41_])
                d_976_v51_: _dafny.MultiSet
                d_976_v51_ = _dafny.MultiSet([d_963_v41_])
                index83_ = default__.safeIndex(919, (d_972_v50_).length(0))
                (d_972_v50_)[index83_] = (d_976_v51_).set(d_963_v41_, default__.abs(d_966_v44_))
            elif True:
                d_977_v52_: _dafny.Map
                d_977_v52_ = _dafny.Map({p1: (self).f16})
                d_962_v40_ = ((d_977_v52_)[p2] if (p2) in (d_977_v52_) else (d_920_v10_)[default__.safeIndex(p0, len(d_920_v10_))])
                d_978_v53_: _dafny.Array
                nw151_ = _dafny.Array(False, 20)
                d_978_v53_ = nw151_
                d_978_v53_ = d_978_v53_
                d_979_v54_: D1
                d_979_v54_ = D1_DC5((self).f16, (self).f5, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))
                d_980_v55_: _dafny.Map
                d_980_v55_ = _dafny.Map({d_923_v13_: (0) - (((d_979_v54_).cf16) - (601))})
                d_980_v55_ = (d_980_v55_) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f25, (self).f25]): default__.fm49(d_923_v13_, p2, (self).f4, d_962_v40_, globalState)}))
                d_981_v56_: int
                d_981_v56_ = 566
                d_982_v57_: _dafny.Map
                d_982_v57_ = _dafny.Map({d_962_v40_: not(d_962_v40_)})
                d_981_v56_ = ((self).f15) * (len(((_dafny.Map({(self).f4: p3})).set(((d_982_v57_)[not(False)] if (not(False)) in (d_982_v57_) else (self).fm0(globalState)), (0) - (p1))).set((self).f4, p1)))
                rhs89_ = d_978_v53_
                rhs90_ = True
                rhs91_ = 183
                rhs92_ = 842
                d_978_v53_ = rhs89_
                d_962_v40_ = rhs90_
                d_981_v56_ = rhs91_
                d_981_v56_ = rhs92_
            d_983_v58_: C1
            nw152_ = C1()
            nw152_.ctor__(len(d_923_v13_), not((self).f16), (self).f4, (p1) - ((0) - ((self).f5)))
            d_983_v58_ = nw152_
        if (self).f4:
            d_984_v60_: _dafny.Array
            def lambda36_(d_985_i5_):
                def iife89_():
                    coll45_ = _dafny.Map()
                    compr_45_: _dafny.MultiSet
                    for compr_45_ in (_dafny.Set({_dafny.MultiSet([False, (self).f16]), _dafny.MultiSet([(self).f4, (self).f16, (self).f4])})).Elements:
                        d_986_v59_: _dafny.MultiSet = compr_45_
                        if (d_986_v59_) in (_dafny.Set({_dafny.MultiSet([False, (self).f16]), _dafny.MultiSet([(self).f4, (self).f16, (self).f4])})):
                            coll45_[d_986_v59_] = True
                    return _dafny.Map(coll45_)
                return iife89_()
                

            init19_ = lambda36_
            nw153_ = _dafny.Array(None, 27)
            for i0_19_ in range(nw153_.length(0)):
                nw153_[i0_19_] = init19_(i0_19_)
            d_984_v60_ = nw153_
            d_987_v61_: _dafny.Map
            d_987_v61_ = _dafny.Map({_dafny.MultiSet([(self).f4, not((self).f4)]): (self).fm0(globalState)})
            index84_ = default__.safeIndex(463, (d_984_v60_).length(0))
            (d_984_v60_)[index84_] = d_987_v61_
            index85_ = default__.safeIndex(463, (d_984_v60_).length(0))
            (d_984_v60_)[index85_] = d_987_v61_
            d_923_v13_ = ((_dafny.SeqWithoutIsStrInference([(self).f25 for d_988_i6_ in range(default__.abs(872))])) + (d_923_v13_)) + (d_923_v13_)
            d_989_v62_: _dafny.MultiSet
            d_989_v62_ = _dafny.MultiSet([(self).f5, p1, p3, len(d_920_v10_)])
            d_990_v63_: _dafny.Set
            d_990_v63_ = _dafny.Set({(d_989_v62_).set(p1, default__.abs(p2))})
            d_991_v64_: _dafny.Seq
            d_991_v64_ = _dafny.SeqWithoutIsStrInference([(d_989_v62_).set((self).f15, default__.abs((self).f5)), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({(self).f4, (self).f16, True, (self).f4}))])), d_989_v62_, d_989_v62_])
            rhs93_ = not ((d_923_v13_) != (d_923_v13_)) or ((self).fm0(globalState))
            rhs94_ = _dafny.Set({(d_991_v64_)[default__.safeIndex((0) - ((self).fm35((self).f16, (self).f4, -872, globalState)), len(d_991_v64_))]})
            r2 = rhs93_
            d_990_v63_ = rhs94_
            d_992_v65_: D4
            d_992_v65_ = D4_DC10((self).f4, (self).f4, d_920_v10_)
            d_993_v66_: D4
            d_993_v66_ = D4_DC11(d_992_v65_)
            source15_ = (D4_DC11(d_992_v65_) if (self).f16 else d_993_v66_)
            if source15_.is_DC9:
                d_994___mcc_h4_ = source15_.cf21
                d_995___mcc_h5_ = source15_.cf22
                d_996___mcc_h6_ = source15_.cf23
                d_997_cf23_ = d_996___mcc_h6_
                d_998_cf22_ = d_995___mcc_h5_
                d_999_cf21_ = d_994___mcc_h4_
                d_1000_v67_: _dafny.Array
                nw154_ = _dafny.Array(_dafny.CodePoint('D'), 22)
                d_1000_v67_ = nw154_
                index86_ = default__.safeIndex(560, (d_1000_v67_).length(0))
                (d_1000_v67_)[index86_] = (self).f25
                d_1001_v68_: _dafny.Array
                nw155_ = _dafny.Array(False, 25)
                d_1001_v68_ = nw155_
                index87_ = default__.safeIndex(552, (d_1001_v68_).length(0))
                (d_1001_v68_)[index87_] = (self).f4
                index88_ = default__.safeIndex(560, (d_1000_v67_).length(0))
                index89_ = default__.safeIndex(552, (d_1001_v68_).length(0))
                rhs95_ = d_1000_v67_
                rhs96_ = (self).f25
                rhs97_ = p3
                rhs98_ = (self).fm0(globalState)
                rhs99_ = (self).f16
                lhs27_ = d_1000_v67_
                lhs28_ = default__.safeIndex(560, (d_1000_v67_).length(0))
                lhs29_ = d_1001_v68_
                lhs30_ = default__.safeIndex(552, (d_1001_v68_).length(0))
                d_1000_v67_ = rhs95_
                lhs27_[lhs28_] = rhs96_
                d_998_cf22_ = rhs97_
                r2 = rhs98_
                lhs29_[lhs30_] = rhs99_
                d_1002_v69_: _dafny.Map
                d_1002_v69_ = _dafny.Map({(d_921_v11_) + (d_921_v11_): d_1001_v68_})
                d_1001_v68_ = ((d_1002_v69_)[d_921_v11_] if (d_921_v11_) in (d_1002_v69_) else d_1001_v68_)
                d_1003_v70_: _dafny.Set
                d_1003_v70_ = _dafny.Set({(d_1001_v68_)[default__.safeIndex(552, (d_1001_v68_).length(0))]})
                d_1004_v71_: _dafny.MultiSet
                d_1004_v71_ = _dafny.MultiSet([_dafny.Set({False, (d_1001_v68_)[default__.safeIndex(552, (d_1001_v68_).length(0))], (d_1001_v68_)[default__.safeIndex(552, (d_1001_v68_).length(0))]}), d_1003_v70_])
                r2 = ((d_1004_v71_).set(d_1003_v70_, default__.abs((0) - (p1)))).issubset(d_1004_v71_)
                d_1005_v72_: _dafny.Array
                nw156_ = _dafny.Array(int(0), 7)
                d_1005_v72_ = nw156_
                index90_ = default__.safeIndex(506, (d_1005_v72_).length(0))
                (d_1005_v72_)[index90_] = p2
                d_1006_v73_: D0
                d_1006_v73_ = D0_DC1(p2, (self).f4, p3, d_998_cf22_)
                index91_ = default__.safeIndex(506, (d_1005_v72_).length(0))
                (d_1005_v72_)[index91_] = (d_1006_v73_).cf7
            elif source15_.is_DC10:
                d_1007___mcc_h7_ = source15_.cf24
                d_1008___mcc_h8_ = source15_.cf25
                d_1009___mcc_h9_ = source15_.cf26
                d_1010_cf26_ = d_1009___mcc_h9_
                d_1011_cf25_ = d_1008___mcc_h8_
                d_1012_cf24_ = d_1007___mcc_h7_
                d_1013_v74_: int
                d_1013_v74_ = 524
                d_1013_v74_ = d_1013_v74_
                d_1014_v75_: _dafny.Map
                d_1014_v75_ = _dafny.Map({d_1013_v74_: (self).f15})
                d_1015_v76_: _dafny.Set
                d_1015_v76_ = _dafny.Set({d_1011_cf25_, (self).f4})
                d_1016_v77_: _dafny.Array
                nw157_ = _dafny.Array(None, 18)
                nw157_[int(0)] = p2
                nw157_[int(1)] = ((d_1014_v75_)[p1] if (p1) in (d_1014_v75_) else (self).f15)
                nw157_[int(2)] = (self).f15
                nw157_[int(3)] = len(d_1015_v76_)
                nw157_[int(4)] = p0
                nw157_[int(5)] = p2
                nw157_[int(6)] = d_1013_v74_
                nw157_[int(7)] = (0) - (d_1013_v74_)
                nw157_[int(8)] = p2
                nw157_[int(9)] = p1
                nw157_[int(10)] = p0
                nw157_[int(11)] = (self).f5
                nw157_[int(12)] = p3
                nw157_[int(13)] = len(d_923_v13_)
                nw157_[int(14)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbhhy")))
                nw157_[int(15)] = p1
                nw157_[int(16)] = (self).f5
                nw157_[int(17)] = p2
                d_1016_v77_ = nw157_
                d_1017_v78_: _dafny.Map
                d_1017_v78_ = _dafny.Map({len(d_1014_v75_): d_1016_v77_})
                d_1018_v79_: D9
                d_1018_v79_ = D9_DC21(d_1017_v78_)
                d_1019_v80_: _dafny.Map
                d_1019_v80_ = _dafny.Map({p1: d_1018_v79_})
                rhs100_ = default__.safeDivisionInt(403, 684)
                rhs101_ = ((d_923_v13_) <= (d_923_v13_) if True else d_1011_cf25_)
                rhs102_ = d_1012_cf24_
                rhs103_ = ((self).f15) - ((self).fm35(d_1011_cf25_, (self).fm0(globalState), len(d_1019_v80_), globalState))
                d_1013_v74_ = rhs100_
                r2 = rhs101_
                d_1012_cf24_ = rhs102_
                d_1013_v74_ = rhs103_
                d_1020_v81_: D5
                d_1020_v81_ = D5_DC14((self).f4, d_920_v10_, d_1016_v77_, p3, p3)
                d_1021_v82_: _dafny.MultiSet
                d_1021_v82_ = _dafny.MultiSet([(d_1020_v81_).cf31, (self).fm0(globalState)])
                d_1022_v83_: _dafny.MultiSet
                d_1022_v83_ = d_1021_v82_
                d_1023_v84_: _dafny.MultiSet
                d_1023_v84_ = _dafny.MultiSet([d_1022_v83_, _dafny.MultiSet(d_920_v10_), d_1022_v83_])
                d_1023_v84_ = d_1023_v84_
                d_1024_v85_: _dafny.Set
                d_1024_v85_ = _dafny.Set({(0) - ((self).fm35((self).f4, d_1012_cf24_, p3, globalState))})
                d_1025_v86_: _dafny.Array
                nw158_ = _dafny.Array(None, 16)
                nw158_[int(0)] = (d_1024_v85_).isdisjoint(d_1024_v85_)
                nw158_[int(1)] = d_1012_cf24_
                nw158_[int(2)] = not (d_1012_cf24_) or (d_1012_cf24_)
                nw158_[int(3)] = ((self).f16) == (d_1011_cf25_)
                nw158_[int(4)] = (self).f4
                nw158_[int(5)] = d_1011_cf25_
                nw158_[int(6)] = d_1011_cf25_
                nw158_[int(7)] = (d_1010_cf26_)[default__.safeIndex(d_1013_v74_, len(d_1010_cf26_))]
                nw158_[int(8)] = not (not(not(d_1012_cf24_))) or (d_1011_cf25_)
                nw158_[int(9)] = (self).f4
                nw158_[int(10)] = (p0) >= (len(d_1010_cf26_))
                nw158_[int(11)] = (self).f16
                nw158_[int(12)] = (self).f4
                nw158_[int(13)] = False
                nw158_[int(14)] = ((self).f15) >= ((self).f15)
                nw158_[int(15)] = not (False) or (not((self).f16))
                d_1025_v86_ = nw158_
                index92_ = default__.safeIndex(10, (d_1025_v86_).length(0))
                (d_1025_v86_)[index92_] = (self).f16
                index93_ = default__.safeIndex(10, (d_1025_v86_).length(0))
                (d_1025_v86_)[index93_] = (self).f4
            elif source15_.is_DC8:
                d_1026___mcc_h10_ = source15_.cf20
                d_1027_cf20_ = d_1026___mcc_h10_
                d_1028_v87_: int
                d_1028_v87_ = 890
                d_1029_v88_: _dafny.Map
                d_1029_v88_ = _dafny.Map({d_1028_v87_: (self).f5})
                d_1030_v90_: _dafny.Map
                def iife90_():
                    coll46_ = _dafny.Map()
                    compr_46_: int
                    for compr_46_ in _dafny.IntegerRange(-521, 280):
                        d_1031_v89_: int = compr_46_
                        if ((-521) <= (d_1031_v89_)) and ((d_1031_v89_) < (280)):
                            coll46_[default__.safeModuloInt(d_1031_v89_, (self).f5)] = (self).f5
                    return _dafny.Map(coll46_)
                d_1030_v90_ = _dafny.Map({len(((d_1029_v88_).set(p3, 596) if (self).f4 else iife90_()
                )): (self).fm0(globalState)})
                rhs104_ = len(d_1030_v90_)
                rhs105_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prueqaxo"))
                d_1028_v87_ = rhs104_
                r0 = rhs105_
                d_1032_v91_: _dafny.Array
                def lambda37_(d_1033_i7_):
                    return default__.safeModuloInt(d_1033_i7_, (self).f15)

                init20_ = lambda37_
                nw159_ = _dafny.Array(None, 22)
                for i0_20_ in range(nw159_.length(0)):
                    nw159_[i0_20_] = init20_(i0_20_)
                d_1032_v91_ = nw159_
                index94_ = default__.safeIndex(80, (d_1032_v91_).length(0))
                (d_1032_v91_)[index94_] = ((self).f5) - (p3)
                d_1034_v92_: _dafny.Map
                d_1034_v92_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f16, (self).f4]): p3})
                index95_ = default__.safeIndex(80, (d_1032_v91_).length(0))
                rhs106_ = (((d_1034_v92_)[d_920_v10_] if (d_920_v10_) in (d_1034_v92_) else len(d_923_v13_)) if (self).f4 else 934)
                rhs107_ = ((self).f15) * (p3)
                lhs31_ = d_1032_v91_
                lhs32_ = default__.safeIndex(80, (d_1032_v91_).length(0))
                d_1028_v87_ = rhs106_
                lhs31_[lhs32_] = rhs107_
                d_1035_v93_: T2
                nw160_ = C2()
                nw160_.ctor__(d_1027_cf20_, d_989_v62_, d_923_v13_, (self).f16, not((self).f4), ((d_989_v62_) | (d_989_v62_)).cardinality)
                d_1035_v93_ = nw160_
                d_1035_v93_ = d_1035_v93_
                d_1036_v94_: C1
                nw161_ = C1()
                nw161_.ctor__(d_1028_v87_, False, ((self).f5) >= (p0), p2)
                d_1036_v94_ = nw161_
            elif True:
                d_1037___mcc_h11_ = source15_.cf27
                d_1038_cf27_ = d_1037___mcc_h11_
                d_1039_v95_: int
                d_1039_v95_ = 250
                d_1040_v96_: _dafny.Map
                d_1040_v96_ = _dafny.Map({(d_920_v10_)[default__.safeIndex(d_1039_v95_, len(d_920_v10_))]: d_923_v13_})
                d_1041_v97_: _dafny.Map
                d_1041_v97_ = _dafny.Map({len(d_923_v13_): (self).f16})
                d_1042_v99_: _dafny.Array
                nw162_ = _dafny.Array(None, 20)
                nw162_[int(0)] = (default__.fm37(globalState)) + (d_923_v13_)
                nw162_[int(1)] = (_dafny.SeqWithoutIsStrInference([(self).f25 for d_1043_i8_ in range(default__.abs(777))])) + (d_923_v13_)
                nw162_[int(2)] = d_923_v13_
                nw162_[int(3)] = (d_923_v13_) + (d_923_v13_)
                nw162_[int(4)] = ((_dafny.SeqWithoutIsStrInference([(self).f25 for d_1044_i9_ in range(default__.abs(6))])).set(default__.safeIndex(-524, len(_dafny.SeqWithoutIsStrInference([(self).f25 for d_1045_i9_ in range(default__.abs(6))]))), (self).f25)) + (d_923_v13_)
                nw162_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eigk"))
                nw162_[int(6)] = _dafny.SeqWithoutIsStrInference([(self).f25 for d_1046_i10_ in range(default__.abs(119))])
                nw162_[int(7)] = d_923_v13_
                nw162_[int(8)] = d_923_v13_
                nw162_[int(9)] = d_923_v13_
                nw162_[int(10)] = _dafny.SeqWithoutIsStrInference([(self).f25 for d_1047_i11_ in range(default__.abs(929))])
                nw162_[int(11)] = d_923_v13_
                nw162_[int(12)] = (d_923_v13_).set(default__.safeIndex((self).f15, len(d_923_v13_)), (self).f25)
                nw162_[int(13)] = (_dafny.SeqWithoutIsStrInference([(self).f25 for d_1048_i12_ in range(default__.abs(-369))])) + (_dafny.SeqWithoutIsStrInference([(self).f25 for d_1049_i13_ in range(default__.abs(911))]))
                nw162_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkgsb"))
                nw162_[int(15)] = ((d_1040_v96_)[((d_1041_v97_)[p2] if (p2) in (d_1041_v97_) else (self).f4)] if (((d_1041_v97_)[p2] if (p2) in (d_1041_v97_) else (self).f4)) in (d_1040_v96_) else d_923_v13_)
                nw162_[int(16)] = d_923_v13_
                nw162_[int(17)] = d_923_v13_
                nw162_[int(18)] = d_923_v13_
                def iife91_():
                    coll47_ = _dafny.Set()
                    compr_47_: int
                    for compr_47_ in _dafny.IntegerRange(973, 204):
                        d_1050_v98_: int = compr_47_
                        if ((973) <= (d_1050_v98_)) and ((d_1050_v98_) < (204)):
                            coll47_ = coll47_.union(_dafny.Set([(d_1050_v98_) * (p2)]))
                    return _dafny.Set(coll47_)
                nw162_[int(19)] = default__.fm52((self).f16, iife91_()
                , (self).f5, (self).f16, globalState)
                d_1042_v99_ = nw162_
                index96_ = default__.safeIndex(492, (d_1042_v99_).length(0))
                (d_1042_v99_)[index96_] = d_923_v13_
                d_1051_v100_: _dafny.Array
                def lambda38_(d_1052_v13_):
                    def lambda39_(d_1053_i14_):
                        return (d_1052_v13_) != (_dafny.SeqWithoutIsStrInference([(self).f25 for d_1054_i15_ in range(default__.abs(-249))]))

                    return lambda39_

                init21_ = lambda38_(d_923_v13_)
                nw163_ = _dafny.Array(None, 24)
                for i0_21_ in range(nw163_.length(0)):
                    nw163_[i0_21_] = init21_(i0_21_)
                d_1051_v100_ = nw163_
                index97_ = default__.safeIndex(457, (d_1051_v100_).length(0))
                (d_1051_v100_)[index97_] = (self).f4
                d_1055_v101_: D1
                d_1055_v101_ = D1_DC4((self).f16, default__.safeModuloInt((self).f5, len(d_923_v13_)), p1)
                d_1056_v102_: D1
                d_1056_v102_ = D1_DC5((self).f4, (self).f15, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pw")))
                index98_ = default__.safeIndex(492, (d_1042_v99_).length(0))
                index99_ = default__.safeIndex(457, (d_1051_v100_).length(0))
                rhs108_ = p2
                rhs109_ = (d_1056_v102_).cf17
                rhs110_ = (self).f16
                rhs111_ = (d_920_v10_)[default__.safeIndex(default__.safeDivisionInt(p2, p3), len(d_920_v10_))]
                rhs112_ = d_1055_v101_
                lhs33_ = d_1042_v99_
                lhs34_ = default__.safeIndex(492, (d_1042_v99_).length(0))
                lhs35_ = d_1051_v100_
                lhs36_ = default__.safeIndex(457, (d_1051_v100_).length(0))
                d_1039_v95_ = rhs108_
                lhs33_[lhs34_] = rhs109_
                lhs35_[lhs36_] = rhs110_
                r2 = rhs111_
                d_1055_v101_ = rhs112_
                r2 = (False) or ((self).f16)
                index100_ = default__.safeIndex(457, (d_1051_v100_).length(0))
                (d_1051_v100_)[index100_] = (self).f4
                d_1057_v103_: _dafny.Map
                d_1057_v103_ = _dafny.Map({p3: p1})
                def iife92_():
                    coll48_ = _dafny.Map()
                    compr_48_: int
                    for compr_48_ in _dafny.IntegerRange(-337, -278):
                        d_1058_v104_: int = compr_48_
                        if ((-337) <= (d_1058_v104_)) and ((d_1058_v104_) < (-278)):
                            coll48_[(d_1058_v104_) - (p2)] = (d_1042_v99_)[default__.safeIndex(492, (d_1042_v99_).length(0))]
                    return _dafny.Map(coll48_)
                d_1039_v95_ = (0) - (((d_1057_v103_)[(self).f5] if ((self).f5) in (d_1057_v103_) else len(iife92_()
                )))
            d_1059_v105_: _dafny.Map
            d_1059_v105_ = _dafny.Map({((self).f5) + ((self).f5): (self).f4})
            d_1060_v106_: _dafny.Map
            d_1060_v106_ = _dafny.Map({(self).f16: d_1059_v105_})
            d_1059_v105_ = ((d_1060_v106_)[False] if (False) in (d_1060_v106_) else d_1059_v105_)
        elif True:
            if (self).f4:
                d_1061_v107_: int
                d_1061_v107_ = 869
                d_1061_v107_ = (d_1061_v107_) + (p1)
                d_1062_v108_: _dafny.Array
                nw164_ = _dafny.Array(None, 8)
                nw164_[int(0)] = (self).f4
                nw164_[int(1)] = (self).f4
                nw164_[int(2)] = (self).f4
                nw164_[int(3)] = (self).f4
                nw164_[int(4)] = (self).f4
                nw164_[int(5)] = (d_920_v10_)[default__.safeIndex(p1, len(d_920_v10_))]
                nw164_[int(6)] = not((self).f16)
                nw164_[int(7)] = (self).f16
                d_1062_v108_ = nw164_
                d_1063_v109_: _dafny.MultiSet
                d_1063_v109_ = _dafny.MultiSet([d_1062_v108_, d_1062_v108_, d_1062_v108_])
                r2 = (d_1063_v109_).isdisjoint((d_1063_v109_) - (d_1063_v109_))
                d_1064_v110_: C1
                nw165_ = C1()
                nw165_.ctor__(449, (self).f4, not((self).f16), p2)
                d_1064_v110_ = nw165_
                r2 = (self).f16
                d_1065_v111_: _dafny.MultiSet
                d_1065_v111_ = _dafny.MultiSet([(self).f16])
                d_1066_v112_: _dafny.Array
                def lambda40_(d_1067_p1_):
                    def lambda41_(d_1068_i17_):
                        return (d_1068_i17_) * ((0) - (d_1067_p1_))

                    return lambda41_

                init22_ = lambda40_(p1)
                nw166_ = _dafny.Array(None, 2)
                for i0_22_ in range(nw166_.length(0)):
                    nw166_[i0_22_] = init22_(i0_22_)
                d_1066_v112_ = nw166_
                d_1069_v113_: _dafny.MultiSet
                d_1069_v113_ = _dafny.MultiSet([d_1066_v112_])
                d_1070_v114_: _dafny.Map
                d_1070_v114_ = _dafny.Map({(self).f4: D1_DC3(d_920_v10_)})
                d_1071_v115_: _dafny.Array
                nw167_ = _dafny.Array(None, 19)
                nw167_[int(0)] = d_921_v11_
                nw167_[int(1)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False])) for d_1072_i16_ in range(default__.abs(-80))])
                nw167_[int(2)] = (_dafny.SeqWithoutIsStrInference([(d_1065_v111_).cardinality, (d_1069_v113_).cardinality]) if (self).f4 else d_921_v11_)
                nw167_[int(3)] = (d_921_v11_ if False else d_921_v11_)
                nw167_[int(4)] = d_921_v11_
                nw167_[int(5)] = d_921_v11_
                nw167_[int(6)] = _dafny.SeqWithoutIsStrInference([p3, p2, (0) - ((self).f15), (0) - (p2), len(d_1070_v114_)])
                nw167_[int(7)] = (d_921_v11_) + (d_921_v11_)
                nw167_[int(8)] = _dafny.SeqWithoutIsStrInference([len(d_920_v10_)])
                nw167_[int(9)] = _dafny.SeqWithoutIsStrInference([p0])
                nw167_[int(10)] = d_921_v11_
                nw167_[int(11)] = _dafny.SeqWithoutIsStrInference([(self).f15 for d_1073_i18_ in range(default__.abs(421))])
                nw167_[int(12)] = _dafny.SeqWithoutIsStrInference([len(d_923_v13_)])
                nw167_[int(13)] = d_921_v11_
                nw167_[int(14)] = _dafny.SeqWithoutIsStrInference([(self).f15, p3, p0, len(d_923_v13_)])
                nw167_[int(15)] = (_dafny.SeqWithoutIsStrInference([p3 for d_1074_i19_ in range(default__.abs(-98))])) + (d_921_v11_)
                nw167_[int(16)] = (d_921_v11_) + (d_921_v11_)
                nw167_[int(17)] = d_921_v11_
                nw167_[int(18)] = d_921_v11_
                d_1071_v115_ = nw167_
                index101_ = default__.safeIndex(201, (d_1071_v115_).length(0))
                (d_1071_v115_)[index101_] = d_921_v11_
                index102_ = default__.safeIndex(374, (d_1066_v112_).length(0))
                (d_1066_v112_)[index102_] = p3
                index103_ = default__.safeIndex(991, (d_1066_v112_).length(0))
                (d_1066_v112_)[index103_] = (0) - (-354)
                d_1075_v117_: D0
                def iife93_():
                    coll49_ = _dafny.Set()
                    compr_49_: int
                    for compr_49_ in _dafny.IntegerRange(277, -425):
                        d_1076_v116_: int = compr_49_
                        if ((277) <= (d_1076_v116_)) and ((d_1076_v116_) < (-425)):
                            coll49_ = coll49_.union(_dafny.Set([(d_1076_v116_) + (p0)]))
                    return _dafny.Set(coll49_)
                d_1075_v117_ = D0_DC1(d_1061_v107_, (self).f16, len(iife93_()
), p3)
                d_1077_v118_: D0
                d_1077_v118_ = D0_DC0((self).f16, (0) - (len(_dafny.Map({(self).f25: (self).f4}))), (d_1075_v117_).cf6, p0, (self).f4)
                d_1078_v119_: _dafny.MultiSet
                d_1078_v119_ = _dafny.MultiSet([(self).fm35((self).f4, (self).f4, p0, globalState)])
                index104_ = default__.safeIndex(201, (d_1071_v115_).length(0))
                index105_ = default__.safeIndex(374, (d_1066_v112_).length(0))
                index106_ = default__.safeIndex(991, (d_1066_v112_).length(0))
                rhs113_ = ((_dafny.SeqWithoutIsStrInference([len(d_923_v13_)]) if (self).f16 else default__.fm44(d_1061_v107_, d_1077_v118_, (self).f16, (self).f16, globalState))) + (d_921_v11_)
                rhs114_ = p2
                rhs115_ = default__.fm55(p1, d_920_v10_, ((self).f15) + (((d_1078_v119_)[(self).f5] if ((self).f5) in (d_1078_v119_) else d_1061_v107_)), globalState)
                lhs37_ = d_1071_v115_
                lhs38_ = default__.safeIndex(201, (d_1071_v115_).length(0))
                lhs39_ = d_1066_v112_
                lhs40_ = default__.safeIndex(374, (d_1066_v112_).length(0))
                lhs41_ = d_1066_v112_
                lhs42_ = default__.safeIndex(991, (d_1066_v112_).length(0))
                lhs37_[lhs38_] = rhs113_
                lhs39_[lhs40_] = rhs114_
                lhs41_[lhs42_] = rhs115_
            elif True:
                d_921_v11_ = d_921_v11_
                d_1079_v120_: _dafny.Array
                def lambda42_(d_1080_i20_):
                    return not((self).f4)

                init23_ = lambda42_
                nw168_ = _dafny.Array(None, 4)
                for i0_23_ in range(nw168_.length(0)):
                    nw168_[i0_23_] = init23_(i0_23_)
                d_1079_v120_ = nw168_
                index107_ = default__.safeIndex(224, (d_1079_v120_).length(0))
                (d_1079_v120_)[index107_] = (p0) >= ((self).f5)
                index108_ = default__.safeIndex(224, (d_1079_v120_).length(0))
                (d_1079_v120_)[index108_] = True
                d_1081_v121_: _dafny.Map
                d_1081_v121_ = _dafny.Map({D1_DC5(not((self).fm0(globalState)), (self).f5, d_923_v13_): d_921_v11_})
                d_1082_v122_: D4
                d_1082_v122_ = D4_DC10(True, (self).f4, d_920_v10_)
                d_1083_v123_: _dafny.MultiSet
                d_1083_v123_ = _dafny.MultiSet([d_1082_v122_])
                d_1084_v124_: _dafny.Set
                d_1084_v124_ = _dafny.Set({(self).f16, (self).f4, not((d_1079_v120_)[default__.safeIndex(224, (d_1079_v120_).length(0))]), (d_1079_v120_)[default__.safeIndex(224, (d_1079_v120_).length(0))], (self).f16})
                d_1085_v125_: _dafny.Array
                nw169_ = _dafny.Array(None, 22)
                nw169_[int(0)] = p0
                nw169_[int(1)] = default__.safeDivisionInt(p3, (self).f5)
                nw169_[int(2)] = p1
                nw169_[int(3)] = (self).f15
                nw169_[int(4)] = p0
                nw169_[int(5)] = len(d_1081_v121_)
                nw169_[int(6)] = default__.safeModuloInt(p0, p3)
                nw169_[int(7)] = (0) - (p3)
                nw169_[int(8)] = (p3) + (80)
                nw169_[int(9)] = (0) - (default__.fm49(d_923_v13_, (d_1083_v123_).cardinality, (d_1079_v120_)[default__.safeIndex(224, (d_1079_v120_).length(0))], False, globalState))
                nw169_[int(10)] = len(d_923_v13_)
                nw169_[int(11)] = p1
                nw169_[int(12)] = p2
                nw169_[int(13)] = (self).f15
                nw169_[int(14)] = default__.safeDivisionInt((self).f15, len(_dafny.Map({(self).f25: (self).f4})))
                nw169_[int(15)] = default__.safeDivisionInt(p3, p0)
                nw169_[int(16)] = (p2) + ((self).f15)
                nw169_[int(17)] = len(d_1084_v124_)
                nw169_[int(18)] = p1
                nw169_[int(19)] = 565
                nw169_[int(20)] = p3
                nw169_[int(21)] = (self).f15
                d_1085_v125_ = nw169_
                d_1085_v125_ = d_1085_v125_
                d_1086_v126_: C1
                nw170_ = C1()
                nw170_.ctor__(p0, (default__.fm49(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sf")), p3, False, (self).f4, globalState)) == (len(d_920_v10_)), (d_1079_v120_)[default__.safeIndex(224, (d_1079_v120_).length(0))], len(d_923_v13_))
                d_1086_v126_ = nw170_
                d_1087_v127_: _dafny.Array
                nw171_ = _dafny.Array(D1.default()(), 10)
                d_1087_v127_ = nw171_
                d_1088_v128_: _dafny.Map
                d_1088_v128_ = _dafny.Map({len((d_923_v13_) + (d_923_v13_)): (self).f4})
                index109_ = default__.safeIndex(224, (d_1079_v120_).length(0))
                rhs116_ = d_1087_v127_
                rhs117_ = not(((self).f16) and ((d_1079_v120_)[default__.safeIndex(224, (d_1079_v120_).length(0))]))
                rhs118_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt((self).f15, (self).f15)])
                rhs119_ = ((d_1088_v128_)[261] if (261) in (d_1088_v128_) else (self).f16)
                lhs43_ = d_1079_v120_
                lhs44_ = default__.safeIndex(224, (d_1079_v120_).length(0))
                d_1087_v127_ = rhs116_
                r2 = rhs117_
                d_921_v11_ = rhs118_
                lhs43_[lhs44_] = rhs119_
            r2 = ((d_923_v13_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jita")))) <= ((d_923_v13_) + (d_923_v13_))
            d_1089_v129_: _dafny.Map
            d_1089_v129_ = _dafny.Map({(self).f15: 106})
            d_1090_v130_: _dafny.Set
            d_1090_v130_ = _dafny.Set({d_1089_v129_, d_1089_v129_})
            d_1091_v131_: _dafny.MultiSet
            d_1091_v131_ = _dafny.MultiSet([(self).f5, p3, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhawq")))])
            d_1092_v132_: C0
            nw172_ = C0()
            nw172_.ctor__(d_1090_v130_, ((d_1091_v131_)[len(d_923_v13_)] if (len(d_923_v13_)) in (d_1091_v131_) else p3))
            d_1092_v132_ = nw172_
            d_1093_v133_: D4
            d_1093_v133_ = D4_DC9(default__.fm57(d_1089_v129_, p2, globalState), p1, (self).f16)
            d_1094_v134_: _dafny.Map
            d_1094_v134_ = _dafny.Map({(self).f4: (d_1093_v133_ if (self).f16 else d_1093_v133_)})
            d_1094_v134_ = (d_1094_v134_) | (d_1094_v134_)
            d_1095_v135_: _dafny.MultiSet
            d_1095_v135_ = (_dafny.MultiSet(d_920_v10_)).intersection(_dafny.MultiSet(d_920_v10_))
            d_1095_v135_ = _dafny.MultiSet([(self).f16, (self).f4, (self).f16, (self).f4])
        d_1096_v136_: _dafny.Map
        d_1096_v136_ = _dafny.Map({(self).f25: False})
        d_1097_v137_: D4
        d_1097_v137_ = D4_DC10((self).f16, (d_920_v10_)[default__.safeIndex((self).f5, len(d_920_v10_))], d_920_v10_)
        d_1098_v138_: _dafny.MultiSet
        d_1098_v138_ = _dafny.MultiSet([(d_1096_v136_).set((self).f25, (self).f16), d_1096_v136_, _dafny.Map({(self).f25: (d_1097_v137_).cf25})])
        r2 = (d_1098_v138_).issubset(d_1098_v138_)
        d_1099_v139_: _dafny.Map
        d_1099_v139_ = _dafny.Map({(self).f4: d_919_v9_})
        d_1100_v140_: _dafny.Array
        nw173_ = _dafny.Array(None, 14)
        nw173_[int(0)] = d_919_v9_
        nw173_[int(1)] = d_919_v9_
        nw173_[int(2)] = (d_919_v9_) | (d_919_v9_)
        nw173_[int(3)] = d_919_v9_
        nw173_[int(4)] = d_919_v9_
        nw173_[int(5)] = d_919_v9_
        nw173_[int(6)] = d_919_v9_
        nw173_[int(7)] = d_919_v9_
        nw173_[int(8)] = d_919_v9_
        nw173_[int(9)] = (d_919_v9_) | (d_919_v9_)
        nw173_[int(10)] = ((d_1099_v139_)[(self).f16] if ((self).f16) in (d_1099_v139_) else d_919_v9_)
        nw173_[int(11)] = d_919_v9_
        nw173_[int(12)] = d_919_v9_
        nw173_[int(13)] = (d_919_v9_) | (d_919_v9_)
        d_1100_v140_ = nw173_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1100_v140_).length(0)):
            d_1101_i21_: int = guard_loop_2_
            if (True) and (((0) <= (d_1101_i21_)) and ((d_1101_i21_) < ((d_1100_v140_).length(0)))):
                (d_1100_v140_)[(d_1101_i21_)] = (d_919_v9_) | ((d_919_v9_) | (d_919_v9_))
        d_1102_v141_: int
        d_1102_v141_ = 838
        d_1103_v142_: _dafny.MultiSet
        d_1103_v142_ = _dafny.MultiSet([(self).f16, (self).f16, (p0) <= (p2), (d_1097_v137_).cf25])
        d_1102_v141_ = ((d_1103_v142_)[(self).f4] if ((self).f4) in (d_1103_v142_) else (p1 if (self).f4 else (self).f5))
        r0 = d_923_v13_
        d_1104_v143_: _dafny.Array
        def lambda43_(d_1105_i22_):
            return (self).f4

        init24_ = lambda43_
        nw174_ = _dafny.Array(None, 13)
        for i0_24_ in range(nw174_.length(0)):
            nw174_[i0_24_] = init24_(i0_24_)
        d_1104_v143_ = nw174_
        d_1106_v144_: _dafny.Seq
        d_1106_v144_ = _dafny.SeqWithoutIsStrInference([d_1104_v143_])
        d_1107_v145_: _dafny.Set
        d_1107_v145_ = _dafny.Set({p1, (self).f5})
        d_1108_v146_: _dafny.Seq
        d_1108_v146_ = _dafny.SeqWithoutIsStrInference([d_1107_v145_])
        d_1109_v147_: _dafny.Map
        d_1109_v147_ = _dafny.Map({d_1106_v144_: (d_1108_v146_)[default__.safeIndex((0) - (p2), len(d_1108_v146_))]})
        d_1110_v148_: D11
        d_1110_v148_ = D11_DC27(((d_1109_v147_)[d_1106_v144_] if (d_1106_v144_) in (d_1109_v147_) else d_1107_v145_))
        pat_let_tv15_ = p1
        def iife94_(_pat_let22_0):
            def iife95_(d_1111_dt__update__tmp_h1_):
                def iife96_(_pat_let23_0):
                    def iife97_(d_1112_dt__update_hcf7_h0_):
                        def iife98_(_pat_let24_0):
                            def iife99_(d_1113_dt__update_hcf6_h0_):
                                return D0_DC1((d_1111_dt__update__tmp_h1_).cf5, d_1113_dt__update_hcf6_h0_, d_1112_dt__update_hcf7_h0_, (d_1111_dt__update__tmp_h1_).cf8)
                            return iife99_(_pat_let24_0)
                        return iife98_((self).f16)
                    return iife97_(_pat_let23_0)
                return iife96_(pat_let_tv15_)
            return iife95_(_pat_let22_0)
        r1 = iife94_(D0_DC1(len((d_1110_v148_).cf51), (d_920_v10_)[default__.safeIndex((0) - (len(d_923_v13_)), len(d_920_v10_))], p2, d_1102_v141_))
        r2 = (self).f16
        return r0, r1, r2

    @property
    def f25(self):
        return self._f25

class C4(T2, T1, T0):
    def  __init__(self):
        self._f18: bool = False
        self._f17: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: bool = False
        self._f5: int = int(0)
        self._f15: int = int(0)
        self._f16: bool = False
        self._f23: int = int(0)
        self._f24: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f18(self):
        return self._f18
    @f18.setter
    def f18(self, value):
        self._f18 = value
    @property
    def f17(self):
        return self._f17
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    def ctor__(self, f23, f24, f17, f18, f4, f5, f15, f16):
        (self)._f23 = f23
        (self)._f24 = f24
        (self)._f17 = f17
        (self).f18 = f18
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f15 = f15
        (self)._f16 = f16

    def fm0(self, globalState):
        return (self).f4

    def fm31(self, p0, p1, p2, p3, globalState):
        return (False) or (not((self).f24))

    def m3(self, p0, globalState):
        d_1114_i0_: int
        d_1114_i0_ = 0
        with _dafny.label("2"):
            while not((self).f24):
                with _dafny.c_label("2"):
                    if (d_1114_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_1114_i0_ = (d_1114_i0_) + (1)
                    d_1115_v0_: _dafny.Array
                    def lambda44_(d_1116_i1_):
                        return (self).f16

                    init25_ = lambda44_
                    nw175_ = _dafny.Array(None, 13)
                    for i0_25_ in range(nw175_.length(0)):
                        nw175_[i0_25_] = init25_(i0_25_)
                    d_1115_v0_ = nw175_
                    index110_ = default__.safeIndex(326, (d_1115_v0_).length(0))
                    (d_1115_v0_)[index110_] = (self).f4
                    index111_ = default__.safeIndex(326, (d_1115_v0_).length(0))
                    (d_1115_v0_)[index111_] = self.f18
                    d_1117_v1_: _dafny.Set
                    d_1117_v1_ = _dafny.Set({True, self.f18})
                    index112_ = default__.safeIndex(326, (d_1115_v0_).length(0))
                    (d_1115_v0_)[index112_] = (d_1117_v1_).ispropersubset(d_1117_v1_)
                    d_1118_v2_: str
                    d_1118_v2_ = _dafny.CodePoint('n')
                    d_1118_v2_ = d_1118_v2_
                    d_1119_v3_: int
                    d_1119_v3_ = 39
                    d_1120_v4_: D0
                    d_1120_v4_ = D0_DC0(self.f18, 486, (self).f16, (self).f5, (self).f16)
                    d_1119_v3_ = default__.fm32(d_1120_v4_, globalState)
                    pass
            pass
        if (self).f4:
            d_1121_v5_: _dafny.Array
            nw176_ = _dafny.Array(int(0), 3)
            d_1121_v5_ = nw176_
            d_1122_v6_: _dafny.Map
            d_1122_v6_ = _dafny.Map({(self).f23: d_1121_v5_})
            d_1123_v7_: _dafny.Map
            d_1123_v7_ = _dafny.Map({True: _dafny.CodePoint('b')})
            d_1124_v8_: _dafny.Seq
            d_1124_v8_ = _dafny.SeqWithoutIsStrInference([len(d_1123_v7_), len(_dafny.SeqWithoutIsStrInference([(self).f24]))])
            d_1125_v9_: _dafny.Seq
            d_1125_v9_ = _dafny.SeqWithoutIsStrInference([len((D9_DC21(d_1122_v6_)).cf43), len(d_1124_v8_)])
            d_1126_v10_: _dafny.MultiSet
            d_1126_v10_ = _dafny.MultiSet([d_1125_v9_])
            d_1127_v11_: D0
            d_1127_v11_ = D0_DC0(self.f18, (self).f15, (self).f4, p0, (self).f4)
            (self).f18 = ((default__.fm33((self).f23, (self).f4, p0, False, globalState)).set(d_1125_v9_, default__.abs(default__.fm32(d_1127_v11_, globalState)))).issubset((d_1126_v10_).intersection(d_1126_v10_))
            d_1128_v12_: _dafny.Array
            nw177_ = _dafny.Array(False, 6)
            d_1128_v12_ = nw177_
            index113_ = default__.safeIndex(45, (d_1128_v12_).length(0))
            (d_1128_v12_)[index113_] = (self).f16
            index114_ = default__.safeIndex(45, (d_1128_v12_).length(0))
            (d_1128_v12_)[index114_] = (self).f4
            index115_ = default__.safeIndex(45, (d_1128_v12_).length(0))
            (d_1128_v12_)[index115_] = (self).f4
            if ((self).f4) == ((self.f18) == ((self).f4)):
                d_1129_v13_: _dafny.Seq
                d_1129_v13_ = _dafny.SeqWithoutIsStrInference([((0) - ((self).f23)) <= ((self).f23), (p0) > ((self).f15), (d_1128_v12_)[default__.safeIndex(45, (d_1128_v12_).length(0))]])
                d_1130_v14_: _dafny.Array
                nw178_ = _dafny.Array(_dafny.Seq({}), 25)
                d_1130_v14_ = nw178_
                d_1131_v15_: _dafny.Map
                d_1131_v15_ = _dafny.Map({(self).f15: self.f18})
                d_1132_v16_: str
                d_1132_v16_ = _dafny.CodePoint('d')
                d_1133_v17_: _dafny.Set
                d_1133_v17_ = _dafny.Set({(self).fm31(len(d_1131_v15_), (self).f4, (self).f4, d_1132_v16_, globalState), self.f18, (self).f24})
                index116_ = default__.safeIndex(68, (d_1130_v14_).length(0))
                (d_1130_v14_)[index116_] = (default__.fm34((self).f5, (self).f23, False, len(d_1133_v17_), globalState)).set(default__.safeIndex(p0, len(default__.fm34((self).f5, (self).f23, False, len(d_1133_v17_), globalState))), (self).f16)
                d_1134_v18_: _dafny.Map
                d_1134_v18_ = _dafny.Map({p0: (-439) + ((self).f23)})
                index117_ = default__.safeIndex(68, (d_1130_v14_).length(0))
                rhs120_ = d_1129_v13_
                rhs121_ = d_1129_v13_
                rhs122_ = d_1134_v18_
                rhs123_ = (d_1133_v17_) == (d_1133_v17_)
                lhs45_ = d_1130_v14_
                lhs46_ = default__.safeIndex(68, (d_1130_v14_).length(0))
                lhs47_ = self
                d_1129_v13_ = rhs120_
                lhs45_[lhs46_] = rhs121_
                d_1134_v18_ = rhs122_
                lhs47_.f18 = rhs123_
                d_1135_v19_: _dafny.Map
                d_1135_v19_ = _dafny.Map({not((d_1128_v12_)[default__.safeIndex(45, (d_1128_v12_).length(0))]): p0})
                d_1136_v20_: _dafny.Map
                d_1136_v20_ = _dafny.Map({d_1135_v19_: 421})
                d_1137_v21_: D6
                d_1137_v21_ = D6_DC16(d_1136_v20_, (self).f23)
                d_1138_v22_: D6
                d_1138_v22_ = D6_DC17(d_1137_v21_)
                d_1138_v22_ = d_1138_v22_
                d_1139_v23_: C3
                nw179_ = C3()
                nw179_.ctor__(d_1132_v16_, 607, (d_1128_v12_)[default__.safeIndex(45, (d_1128_v12_).length(0))], self.f18, default__.safeDivisionInt(default__.fm55((self).f15, d_1129_v13_, p0, globalState), 435))
                d_1139_v23_ = nw179_
                d_1140_v24_: _dafny.Map
                d_1140_v24_ = _dafny.Map({(self).f17: (self).f17})
                d_1140_v24_ = (d_1140_v24_).set((self).f17, default__.fm37(globalState))
                d_1141_v25_: _dafny.Array
                nw180_ = _dafny.Array(D5.default()(), 24)
                d_1141_v25_ = nw180_
                d_1141_v25_ = d_1141_v25_
            elif True:
                index118_ = default__.safeIndex(730, (d_1121_v5_).length(0))
                (d_1121_v5_)[index118_] = p0
                index119_ = default__.safeIndex(730, (d_1121_v5_).length(0))
                (d_1121_v5_)[index119_] = -327
                d_1142_v26_: _dafny.Map
                d_1142_v26_ = _dafny.Map({self.f18: p0})
                d_1143_v27_: _dafny.Seq
                d_1143_v27_ = _dafny.SeqWithoutIsStrInference([d_1142_v26_])
                index120_ = default__.safeIndex(45, (d_1128_v12_).length(0))
                (d_1128_v12_)[index120_] = ((d_1143_v27_)[default__.safeIndex(p0, len(d_1143_v27_))]) == (d_1142_v26_)
                d_1144_v28_: _dafny.Seq
                d_1144_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iywbpbre"))
                d_1145_v29_: _dafny.Array
                nw181_ = _dafny.Array(_dafny.CodePoint('D'), 6)
                d_1145_v29_ = nw181_
                d_1146_v30_: str
                d_1146_v30_ = _dafny.CodePoint('y')
                index121_ = default__.safeIndex(138, (d_1145_v29_).length(0))
                (d_1145_v29_)[index121_] = d_1146_v30_
                d_1147_v31_: _dafny.MultiSet
                d_1147_v31_ = _dafny.MultiSet([(self).fm0(globalState)])
                d_1148_v32_: _dafny.Seq
                d_1148_v32_ = _dafny.SeqWithoutIsStrInference([(d_1128_v12_)[default__.safeIndex(45, (d_1128_v12_).length(0))], self.f18, (self).f4])
                d_1149_v33_: D1
                d_1149_v33_ = D1_DC3(d_1148_v32_)
                d_1150_v34_: _dafny.Map
                d_1150_v34_ = _dafny.Map({d_1149_v33_: (self).f16})
                d_1151_v35_: _dafny.Seq
                d_1151_v35_ = _dafny.SeqWithoutIsStrInference([((d_1150_v34_)[d_1149_v33_] if (d_1149_v33_) in (d_1150_v34_) else (self).f4), True, (self).f24, (self).f16])
                index122_ = default__.safeIndex(138, (d_1145_v29_).length(0))
                index123_ = default__.safeIndex(45, (d_1128_v12_).length(0))
                index124_ = default__.safeIndex(730, (d_1121_v5_).length(0))
                rhs124_ = _dafny.SeqWithoutIsStrInference([d_1146_v30_ for d_1152_i2_ in range(default__.abs(922))])
                rhs125_ = d_1146_v30_
                rhs126_ = (self).f16
                rhs127_ = 962
                rhs128_ = _dafny.MultiSet(d_1151_v35_)
                lhs48_ = d_1145_v29_
                lhs49_ = default__.safeIndex(138, (d_1145_v29_).length(0))
                lhs50_ = d_1128_v12_
                lhs51_ = default__.safeIndex(45, (d_1128_v12_).length(0))
                lhs52_ = d_1121_v5_
                lhs53_ = default__.safeIndex(730, (d_1121_v5_).length(0))
                d_1144_v28_ = rhs124_
                lhs48_[lhs49_] = rhs125_
                lhs50_[lhs51_] = rhs126_
                lhs52_[lhs53_] = rhs127_
                d_1147_v31_ = rhs128_
                (self).f18 = (self).f24
                d_1153_v36_: _dafny.Seq
                d_1153_v36_ = _dafny.SeqWithoutIsStrInference([d_1144_v28_])
                d_1153_v36_ = ((d_1153_v36_) + (_dafny.SeqWithoutIsStrInference([(self).f17, d_1144_v28_]))) + (d_1153_v36_)
            index125_ = default__.safeIndex(45, (d_1128_v12_).length(0))
            (d_1128_v12_)[index125_] = not(not((self).f4))
        elif True:
            d_1154_v37_: _dafny.Set
            d_1154_v37_ = _dafny.Set({(self).f4})
            d_1155_v38_: _dafny.Array
            nw182_ = _dafny.Array(None, 10)
            nw182_[int(0)] = ((self).f24) == ((self).f4)
            nw182_[int(1)] = (False if (self).f4 else (self).f4)
            nw182_[int(2)] = (self).f24
            nw182_[int(3)] = (self).f16
            nw182_[int(4)] = self.f18
            nw182_[int(5)] = (d_1154_v37_).issubset(_dafny.Set({False, (self).f24}))
            nw182_[int(6)] = (self).f4
            nw182_[int(7)] = (self).f24
            nw182_[int(8)] = (self).fm0(globalState)
            nw182_[int(9)] = self.f18
            d_1155_v38_ = nw182_
            index126_ = default__.safeIndex(492, (d_1155_v38_).length(0))
            (d_1155_v38_)[index126_] = ((self).f23) < ((self).f15)
            index127_ = default__.safeIndex(492, (d_1155_v38_).length(0))
            (d_1155_v38_)[index127_] = (self).f16
            d_1156_v39_: _dafny.Seq
            d_1156_v39_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f5), (self).f23])
            (self).f18 = ((self).f15) != ((d_1156_v39_)[default__.safeIndex((self).f5, len(d_1156_v39_))])
            if (self).f16:
                d_1157_v40_: _dafny.MultiSet
                d_1157_v40_ = _dafny.MultiSet([(self).f24])
                d_1158_v41_: _dafny.MultiSet
                d_1158_v41_ = _dafny.MultiSet([(self).f23, (d_1157_v40_).cardinality])
                d_1159_v42_: str
                d_1159_v42_ = _dafny.CodePoint('b')
                d_1160_v43_: _dafny.MultiSet
                d_1160_v43_ = _dafny.MultiSet([d_1159_v42_])
                d_1161_v44_: D12
                d_1161_v44_ = D12_DC32(d_1160_v43_)
                d_1162_v45_: _dafny.Seq
                d_1162_v45_ = _dafny.SeqWithoutIsStrInference([self.f18])
                d_1163_v46_: _dafny.Seq
                d_1163_v46_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f5, default__.fm49((self).f17, (self).f5, not((d_1162_v45_)[default__.safeIndex(len(d_1162_v45_), len(d_1162_v45_))]), True, globalState)]), d_1158_v41_])
                d_1164_v47_: D9
                d_1164_v47_ = D9_DC23(self.f18, (d_1156_v39_)[default__.safeIndex((self).f23, len(d_1156_v39_))], ((d_1158_v41_).set(859, default__.abs(((d_1161_v44_).cf66).cardinality))).ispropersubset((d_1163_v46_)[default__.safeIndex(len((self).f17), len(d_1163_v46_))]))
                d_1165_v48_: _dafny.Array
                nw183_ = _dafny.Array(None, 3)
                nw183_[int(0)] = (_dafny.MultiSet([(self).f24, (self).f24])).cardinality
                nw183_[int(1)] = p0
                nw183_[int(2)] = len((self).f17)
                d_1165_v48_ = nw183_
                d_1166_v49_: D11
                d_1166_v49_ = D11_DC30((self).f4, d_1165_v48_, p0, self.f18)
                d_1167_v50_: _dafny.Map
                d_1167_v50_ = _dafny.Map({(self).f5: (self).f16})
                d_1168_v51_: _dafny.Map
                d_1168_v51_ = _dafny.Map({d_1167_v50_: (self).f16})
                def iife100_():
                    coll50_ = _dafny.Map()
                    compr_50_: int
                    for compr_50_ in _dafny.IntegerRange(225, 2):
                        d_1169_v52_: int = compr_50_
                        if ((225) <= (d_1169_v52_)) and ((d_1169_v52_) < (2)):
                            coll50_[default__.safeModuloInt(d_1169_v52_, (self).f23)] = self.f18
                    return _dafny.Map(coll50_)
                def iife101_():
                    coll51_ = _dafny.Map()
                    compr_51_: int
                    for compr_51_ in _dafny.IntegerRange(225, 2):
                        d_1170_v52_: int = compr_51_
                        if ((225) <= (d_1170_v52_)) and ((d_1170_v52_) < (2)):
                            coll51_[default__.safeModuloInt(d_1170_v52_, (self).f23)] = self.f18
                    return _dafny.Map(coll51_)
                d_1164_v47_ = D9_DC23((self).f16, (d_1166_v49_).cf63, ((d_1168_v51_)[iife100_()
] if (iife101_()
) in (d_1168_v51_) else True))
                d_1171_v53_: int
                d_1171_v53_ = -514
                d_1171_v53_ = default__.fm55((d_1156_v39_)[default__.safeIndex(len(default__.fm58(len((self).f17), d_1171_v53_, globalState)), len(d_1156_v39_))], (d_1162_v45_) + (d_1162_v45_), default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_1159_v42_ for d_1172_i3_ in range(default__.abs(395))])), 57), globalState)
                d_1173_v54_: D1
                d_1173_v54_ = D1_DC3(d_1162_v45_)
                pat_let_tv16_ = d_1162_v45_
                d_1174_v55_: _dafny.Seq
                def iife102_(_pat_let25_0):
                    def iife103_(d_1175_dt__update__tmp_h0_):
                        def iife104_(_pat_let26_0):
                            def iife105_(d_1176_dt__update_hcf11_h0_):
                                return D1_DC3(d_1176_dt__update_hcf11_h0_)
                            return iife105_(_pat_let26_0)
                        return iife104_(pat_let_tv16_)
                    return iife103_(_pat_let25_0)
                d_1174_v55_ = _dafny.SeqWithoutIsStrInference([iife102_(d_1173_v54_), D1_DC3(d_1162_v45_)])
                d_1177_v56_: _dafny.Map
                d_1177_v56_ = _dafny.Map({len((d_1156_v39_) + (_dafny.SeqWithoutIsStrInference([(0) - (len(d_1174_v55_)), (d_1157_v40_).cardinality, (self).f5]))): (self).f17})
                d_1178_v57_: _dafny.Seq
                d_1178_v57_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({len(d_1167_v50_): (self).f17})])
                d_1177_v56_ = (d_1178_v57_)[default__.safeIndex(-170, len(d_1178_v57_))]
                d_1171_v53_ = default__.fm49((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmuak")) if False else (self).f17), ((self).f15) + (p0), (self).fm31(p0, (self).f24, (self).fm31((self).f15, (self).f24, (self).f24, d_1159_v42_, globalState), d_1159_v42_, globalState), (self).f16, globalState)
                index128_ = default__.safeIndex(492, (d_1155_v38_).length(0))
                (d_1155_v38_)[index128_] = (self).f24
            elif True:
                d_1179_v58_: int
                d_1179_v58_ = 228
                d_1179_v58_ = (81 if (self).f4 else (len((self).f17)) * (d_1179_v58_))
                index129_ = default__.safeIndex(492, (d_1155_v38_).length(0))
                (d_1155_v38_)[index129_] = (self).f4
                d_1180_v59_: _dafny.Seq
                d_1180_v59_ = _dafny.SeqWithoutIsStrInference([self.f18, not(not(not(not(self.f18))))])
                d_1181_v60_: str
                d_1182_v61_: int
                d_1183_v62_: D0
                d_1184_v63_: int
                out12_: str
                out13_: int
                out14_: D0
                out15_: int
                out12_, out13_, out14_, out15_ = (self).m15((self).f16, len(d_1180_v59_), globalState)
                d_1181_v60_ = out12_
                d_1182_v61_ = out13_
                d_1183_v62_ = out14_
                d_1184_v63_ = out15_
                d_1181_v60_ = d_1181_v60_
                d_1155_v38_ = d_1155_v38_
            d_1185_v64_: str
            d_1185_v64_ = _dafny.CodePoint('a')
            d_1186_v65_: _dafny.Array
            nw184_ = _dafny.Array(None, 10)
            nw184_[int(0)] = d_1185_v64_
            nw184_[int(1)] = d_1185_v64_
            nw184_[int(2)] = d_1185_v64_
            nw184_[int(3)] = ((self).f17)[default__.safeIndex((d_1156_v39_)[default__.safeIndex(p0, len(d_1156_v39_))], len((self).f17))]
            nw184_[int(4)] = d_1185_v64_
            nw184_[int(5)] = default__.fm50(globalState)
            nw184_[int(6)] = (_dafny.CodePoint('m') if (self).f16 else d_1185_v64_)
            nw184_[int(7)] = d_1185_v64_
            nw184_[int(8)] = d_1185_v64_
            nw184_[int(9)] = d_1185_v64_
            d_1186_v65_ = nw184_
            index130_ = default__.safeIndex(558, (d_1186_v65_).length(0))
            (d_1186_v65_)[index130_] = d_1185_v64_
            index131_ = default__.safeIndex(558, (d_1186_v65_).length(0))
            (d_1186_v65_)[index131_] = ((self).f17)[default__.safeIndex((self).f15, len((self).f17))]
            d_1187_v66_: _dafny.Seq
            d_1187_v66_ = _dafny.SeqWithoutIsStrInference([(d_1186_v65_)[default__.safeIndex(558, (d_1186_v65_).length(0))], (d_1186_v65_)[default__.safeIndex(558, (d_1186_v65_).length(0))], (d_1186_v65_)[default__.safeIndex(558, (d_1186_v65_).length(0))]])
            d_1187_v66_ = d_1187_v66_
        d_1188_v67_: _dafny.Seq
        d_1188_v67_ = _dafny.SeqWithoutIsStrInference([(self).f23, (self).f15])
        d_1189_v68_: _dafny.Map
        d_1189_v68_ = _dafny.Map({True: (self).f4})
        d_1190_v69_: _dafny.Array
        nw185_ = _dafny.Array(None, 9)
        nw185_[int(0)] = (self).f24
        nw185_[int(1)] = ((d_1188_v67_)[default__.safeIndex((self).f23, len(d_1188_v67_))]) == ((self).f23)
        nw185_[int(2)] = ((d_1189_v68_)[self.f18] if (self.f18) in (d_1189_v68_) else self.f18)
        nw185_[int(3)] = self.f18
        nw185_[int(4)] = False
        nw185_[int(5)] = self.f18
        nw185_[int(6)] = self.f18
        nw185_[int(7)] = True
        nw185_[int(8)] = self.f18
        d_1190_v69_ = nw185_
        index132_ = default__.safeIndex(243, (d_1190_v69_).length(0))
        (d_1190_v69_)[index132_] = self.f18
        d_1191_v70_: int
        d_1191_v70_ = -378
        index133_ = default__.safeIndex(243, (d_1190_v69_).length(0))
        rhs129_ = (self).f4
        rhs130_ = default__.safeModuloInt((self).f15, (self).f5)
        rhs131_ = d_1190_v69_
        lhs54_ = d_1190_v69_
        lhs55_ = default__.safeIndex(243, (d_1190_v69_).length(0))
        lhs54_[lhs55_] = rhs129_
        d_1191_v70_ = rhs130_
        d_1190_v69_ = rhs131_
        d_1192_v71_: _dafny.Array
        def lambda45_(d_1193_v67_):
            def lambda46_(d_1194_i4_):
                return (d_1194_i4_) * (len(d_1193_v67_))

            return lambda46_

        init26_ = lambda45_(d_1188_v67_)
        nw186_ = _dafny.Array(None, 25)
        for i0_26_ in range(nw186_.length(0)):
            nw186_[i0_26_] = init26_(i0_26_)
        d_1192_v71_ = nw186_
        d_1195_v72_: _dafny.Map
        d_1195_v72_ = _dafny.Map({d_1192_v71_: (d_1190_v69_)[default__.safeIndex(243, (d_1190_v69_).length(0))]})
        d_1195_v72_ = (d_1195_v72_).set(d_1192_v71_, (d_1190_v69_)[default__.safeIndex(243, (d_1190_v69_).length(0))])
        if (D11_DC30(not(((d_1189_v68_)[(self).f24] if ((self).f24) in (d_1189_v68_) else (self).f4)), d_1192_v71_, (self).f23, (d_1190_v69_)[default__.safeIndex(243, (d_1190_v69_).length(0))])).cf61:
            (self).f18 = (self).f24
            d_1196_v73_: D1
            d_1196_v73_ = D1_DC5(self.f18, d_1191_v70_, (self).f17)
            pat_let_tv17_ = p0
            d_1197_v74_: _dafny.Seq
            def iife106_(_pat_let27_0):
                def iife107_(d_1198_dt__update__tmp_h1_):
                    def iife108_(_pat_let28_0):
                        def iife109_(d_1199_dt__update_hcf16_h0_):
                            return D1_DC5((d_1198_dt__update__tmp_h1_).cf15, d_1199_dt__update_hcf16_h0_, (d_1198_dt__update__tmp_h1_).cf17)
                        return iife109_(_pat_let28_0)
                    return iife108_(pat_let_tv17_)
                return iife107_(_pat_let27_0)
            d_1197_v74_ = _dafny.SeqWithoutIsStrInference([iife106_(d_1196_v73_), d_1196_v73_])
            d_1191_v70_ = len(d_1197_v74_)
            index134_ = default__.safeIndex(243, (d_1190_v69_).length(0))
            (d_1190_v69_)[index134_] = (self).f4
            d_1191_v70_ = ((0) - ((self).f23)) + ((self).f23)
            d_1200_v75_: str
            d_1200_v75_ = _dafny.CodePoint('s')
            d_1201_v76_: _dafny.Set
            d_1201_v76_ = _dafny.Set({(self).f4, (self).f24})
            d_1202_v78_: D9
            d_1202_v78_ = D9_DC22((d_1190_v69_)[default__.safeIndex(243, (d_1190_v69_).length(0))])
            d_1203_v79_: _dafny.Map
            d_1203_v79_ = _dafny.Map({d_1201_v76_: d_1202_v78_})
            d_1204_v80_: _dafny.MultiSet
            d_1204_v80_ = _dafny.MultiSet([d_1190_v69_, d_1190_v69_, d_1190_v69_])
            d_1205_v81_: _dafny.Map
            d_1205_v81_ = _dafny.Map({D7_DC18(d_1204_v80_): True})
            d_1206_v82_: D7
            d_1206_v82_ = D7_DC18(d_1204_v80_)
            index135_ = default__.safeIndex(243, (d_1190_v69_).length(0))
            index136_ = default__.safeIndex(243, (d_1190_v69_).length(0))
            def iife110_():
                coll52_ = _dafny.Map()
                compr_52_: _dafny.Set
                for compr_52_ in (d_1203_v79_).keys.Elements:
                    d_1207_v77_: _dafny.Set = compr_52_
                    if (d_1207_v77_) in (d_1203_v79_):
                        coll52_[d_1207_v77_] = _dafny.MultiSet([d_1200_v75_])
                return _dafny.Map(coll52_)
            rhs132_ = (self).fm31(default__.fm49(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amkdf")), len(d_1189_v68_), (d_1190_v69_)[default__.safeIndex(243, (d_1190_v69_).length(0))], self.f18, globalState), (d_1190_v69_)[default__.safeIndex(243, (d_1190_v69_).length(0))], (self).f24, d_1200_v75_, globalState)
            rhs133_ = (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([(self).f24])), (self).f23)) * ((self).f23)
            rhs134_ = ((self).f5) > ((len(d_1201_v76_) if (d_1190_v69_)[default__.safeIndex(243, (d_1190_v69_).length(0))] else len(iife110_()
            )))
            rhs135_ = ((d_1189_v68_)[((d_1205_v81_)[d_1206_v82_] if (d_1206_v82_) in (d_1205_v81_) else (self).f16)] if (((d_1205_v81_)[d_1206_v82_] if (d_1206_v82_) in (d_1205_v81_) else (self).f16)) in (d_1189_v68_) else (self).f4)
            rhs136_ = d_1201_v76_
            lhs56_ = d_1190_v69_
            lhs57_ = default__.safeIndex(243, (d_1190_v69_).length(0))
            lhs58_ = d_1190_v69_
            lhs59_ = default__.safeIndex(243, (d_1190_v69_).length(0))
            lhs60_ = self
            lhs61_ = globalState
            lhs56_[lhs57_] = rhs132_
            d_1191_v70_ = rhs133_
            lhs58_[lhs59_] = rhs134_
            lhs60_.f18 = rhs135_
            lhs61_.f0 = rhs136_
        elif True:
            d_1208_v83_: str
            d_1208_v83_ = _dafny.CodePoint('r')
            d_1209_v84_: _dafny.Array
            nw187_ = _dafny.Array(None, 4)
            nw187_[int(0)] = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1210_i5_ in range(default__.abs(636))])) + ((self).f17)).set(default__.safeIndex((self).f23, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1211_i5_ in range(default__.abs(636))])) + ((self).f17))), d_1208_v83_)
            nw187_[int(1)] = (self).f17
            nw187_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgtgjn"))
            nw187_[int(3)] = (self).f17
            d_1209_v84_ = nw187_
            d_1212_v85_: _dafny.Seq
            d_1212_v85_ = _dafny.SeqWithoutIsStrInference([(self).f17])
            index137_ = default__.safeIndex(179, (d_1209_v84_).length(0))
            (d_1209_v84_)[index137_] = (d_1212_v85_)[default__.safeIndex((self).f23, len(d_1212_v85_))]
            index138_ = default__.safeIndex(179, (d_1209_v84_).length(0))
            (d_1209_v84_)[index138_] = (self).f17
            d_1213_v86_: _dafny.Seq
            d_1213_v86_ = _dafny.SeqWithoutIsStrInference([(self).f16])
            d_1214_v87_: _dafny.Map
            d_1214_v87_ = _dafny.Map({d_1213_v86_: p0})
            d_1215_v88_: _dafny.Map
            d_1215_v88_ = _dafny.Map({p0: p0})
            d_1216_v89_: _dafny.MultiSet
            d_1216_v89_ = _dafny.MultiSet([d_1215_v88_, d_1215_v88_])
            d_1217_v90_: C2
            nw188_ = C2()
            nw188_.ctor__(d_1190_v69_, _dafny.MultiSet([-381, ((d_1214_v87_)[_dafny.SeqWithoutIsStrInference([(self).f16])] if (_dafny.SeqWithoutIsStrInference([(self).f16])) in (d_1214_v87_) else (self).f5)]), ((d_1209_v84_)[default__.safeIndex(179, (d_1209_v84_).length(0))]) + ((self).f17), (_dafny.MultiSet([d_1215_v88_])).ispropersubset(d_1216_v89_), (self).f16, (self).f23)
            d_1217_v90_ = nw188_
            d_1218_v91_: _dafny.Map
            d_1218_v91_ = _dafny.Map({_dafny.CodePoint('q'): d_1191_v70_})
            d_1219_v92_: _dafny.Map
            d_1219_v92_ = _dafny.Map({((self).f17)[default__.safeIndex(len(d_1218_v91_), len((self).f17))]: (self).f17})
            d_1219_v92_ = (d_1219_v92_).set(d_1208_v83_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))
            index139_ = default__.safeIndex(179, (d_1209_v84_).length(0))
            (d_1209_v84_)[index139_] = ((d_1209_v84_)[default__.safeIndex(179, (d_1209_v84_).length(0))]) + ((d_1209_v84_)[default__.safeIndex(179, (d_1209_v84_).length(0))])
            d_1220_v93_: C3
            nw189_ = C3()
            nw189_.ctor__(d_1208_v83_, len((d_1213_v86_) + (d_1213_v86_)), (self).f4, (self).f16, d_1191_v70_)
            d_1220_v93_ = nw189_
        (self).f18 = ((self).f15) == (d_1191_v70_)

    def m15(self, p0, p1, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: D0 = D0.default()()
        r3: int = int(0)
        r1 = (0) - (-326)
        r1 = default__.fm49(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1221_i0_ in range(default__.abs(815))]), p1, ((self).f17) == ((self).f17), self.f18, globalState)
        d_1222_v0_: D0
        d_1222_v0_ = D0_DC1(p1, (self).f24, p1, (self).f5)
        d_1223_v1_: _dafny.Seq
        d_1223_v1_ = _dafny.SeqWithoutIsStrInference([(self).f24, p0, (self).f16])
        def iife111_(_pat_let29_0):
            def iife112_(d_1224_dt__update__tmp_h0_):
                def iife113_(_pat_let30_0):
                    def iife114_(d_1225_dt__update_hcf6_h0_):
                        def iife115_(_pat_let31_0):
                            def iife116_(d_1226_dt__update_hcf5_h0_):
                                return D0_DC1(d_1226_dt__update_hcf5_h0_, d_1225_dt__update_hcf6_h0_, (d_1224_dt__update__tmp_h0_).cf7, (d_1224_dt__update__tmp_h0_).cf8)
                            return iife116_(_pat_let31_0)
                        return iife115_((self).f23)
                    return iife114_(_pat_let30_0)
                return iife113_((self).f16)
            return iife112_(_pat_let29_0)
        (self).f18 = ((self).f24 if (iife111_(d_1222_v0_)).cf6 else (d_1223_v1_)[default__.safeIndex((self).f15, len(d_1223_v1_))])
        d_1227_v2_: str
        d_1227_v2_ = _dafny.CodePoint('r')
        d_1228_v3_: _dafny.Array
        nw190_ = _dafny.Array(None, 24)
        nw190_[int(0)] = not(p0)
        nw190_[int(1)] = (self).f4
        nw190_[int(2)] = (self).f24
        nw190_[int(3)] = False
        nw190_[int(4)] = True
        nw190_[int(5)] = (self).f4
        nw190_[int(6)] = self.f18
        nw190_[int(7)] = self.f18
        nw190_[int(8)] = (self).f4
        nw190_[int(9)] = (self).f16
        nw190_[int(10)] = (self).fm31((0) - (p1), (self).f24, (self).f24, d_1227_v2_, globalState)
        nw190_[int(11)] = not((self).f4)
        nw190_[int(12)] = (self).f4
        nw190_[int(13)] = (self).f16
        nw190_[int(14)] = (self).f16
        nw190_[int(15)] = (self).f16
        nw190_[int(16)] = (self).f24
        nw190_[int(17)] = (D1_DC4((self).f4, 576, (self).f15)).cf12
        nw190_[int(18)] = (self).f4
        nw190_[int(19)] = p0
        nw190_[int(20)] = True
        nw190_[int(21)] = (self).f24
        nw190_[int(22)] = (self).f16
        nw190_[int(23)] = p0
        d_1228_v3_ = nw190_
        d_1229_v4_: _dafny.Array
        nw191_ = _dafny.Array(None, 1)
        nw191_[int(0)] = d_1228_v3_
        d_1229_v4_ = nw191_
        index140_ = default__.safeIndex(761, (d_1229_v4_).length(0))
        (d_1229_v4_)[index140_] = d_1228_v3_
        index141_ = default__.safeIndex(761, (d_1229_v4_).length(0))
        (d_1229_v4_)[index141_] = d_1228_v3_
        d_1230_v5_: _dafny.Seq
        d_1230_v5_ = _dafny.SeqWithoutIsStrInference([(d_1229_v4_)[default__.safeIndex(761, (d_1229_v4_).length(0))], (d_1229_v4_)[default__.safeIndex(761, (d_1229_v4_).length(0))], d_1228_v3_])
        d_1231_v6_: _dafny.Seq
        d_1231_v6_ = _dafny.SeqWithoutIsStrInference([(d_1230_v5_)[default__.safeIndex((self).f5, len(d_1230_v5_))]])
        d_1232_v7_: D13
        d_1232_v7_ = D13_DC34(d_1230_v5_)
        d_1231_v6_ = ((d_1231_v6_) + (d_1231_v6_)) + (((d_1232_v7_).cf71) + (d_1231_v6_))
        if (self).f16:
            d_1233_v8_: _dafny.Map
            d_1233_v8_ = _dafny.Map({(d_1223_v1_)[default__.safeIndex((0) - (p1), len(d_1223_v1_))]: (self).f15})
            d_1234_v9_: _dafny.Map
            d_1234_v9_ = _dafny.Map({(self).f15: 469})
            d_1235_v10_: _dafny.Map
            d_1235_v10_ = _dafny.Map({((d_1233_v8_)[(self).f4] if ((self).f4) in (d_1233_v8_) else len(d_1234_v9_)): (self).f5})
            d_1236_v11_: _dafny.MultiSet
            d_1236_v11_ = _dafny.MultiSet([(self).f17])
            d_1234_v9_ = (d_1234_v9_).set((self).f15, ((d_1236_v11_).intersection(d_1236_v11_)).cardinality)
            d_1237_v12_: D12
            d_1237_v12_ = D12_DC33(self.f18, p0, ((self).f17)[default__.safeIndex((self).f15, len((self).f17))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "la"))))
            d_1238_v13_: _dafny.Seq
            d_1238_v13_ = _dafny.SeqWithoutIsStrInference([default__.fm55((self).f23, d_1223_v1_, (d_1222_v0_).cf7, globalState), p1])
            d_1239_v14_: _dafny.Map
            d_1239_v14_ = _dafny.Map({self.f18: (self).f24})
            d_1240_v15_: _dafny.MultiSet
            d_1240_v15_ = _dafny.MultiSet([(self).f24])
            d_1241_v16_: _dafny.Set
            d_1241_v16_ = _dafny.Set({d_1235_v10_, _dafny.Map({(d_1240_v15_).cardinality: (self).f15})})
            d_1242_v17_: _dafny.Set
            d_1242_v17_ = d_1241_v16_
            d_1243_v18_: _dafny.Map
            d_1243_v18_ = _dafny.Map({(self).f23: d_1242_v17_})
            d_1244_v19_: _dafny.Map
            d_1244_v19_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oj"))).set(default__.safeIndex((self).f15, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oj")))), (d_1237_v12_).cf69): (default__.fm59((self).f15, d_1238_v13_, False, globalState)).set(d_1239_v14_, (d_1243_v18_).set(len(d_1238_v13_), d_1242_v17_))})
            d_1245_v20_: _dafny.Map
            d_1245_v20_ = _dafny.Map({d_1239_v14_: d_1243_v18_})
            d_1244_v19_ = (d_1244_v19_).set((self).f17, d_1245_v20_)
            r0 = d_1227_v2_
            d_1246_v21_: D12
            d_1246_v21_ = D12_DC32(_dafny.MultiSet([_dafny.CodePoint('j')]))
            d_1246_v21_ = d_1246_v21_
            d_1247_v22_: _dafny.MultiSet
            d_1247_v22_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(self).f23 for d_1248_i1_ in range(default__.abs(-845))])), p1, (self).f5])
            d_1249_v23_: D1
            d_1249_v23_ = D1_DC4((d_1223_v1_)[default__.safeIndex((self).f15, len(d_1223_v1_))], (self).f23, (self).f15)
            d_1250_v24_: C2
            nw192_ = C2()
            nw192_.ctor__((d_1229_v4_)[default__.safeIndex(761, (d_1229_v4_).length(0))], (d_1247_v22_) - (d_1247_v22_), ((self).f17) + (((self).f17).set(default__.safeIndex(916, len((self).f17)), d_1227_v2_)), (d_1249_v23_).cf12, not ((self).f16) or (p0), p1)
            d_1250_v24_ = nw192_
        elif True:
            d_1251_v25_: _dafny.Map
            d_1251_v25_ = _dafny.Map({p1: p1})
            d_1252_v26_: _dafny.Set
            d_1252_v26_ = _dafny.Set({d_1251_v25_})
            d_1253_v27_: _dafny.Set
            d_1253_v27_ = d_1252_v26_
            d_1254_v28_: C0
            nw193_ = C0()
            nw193_.ctor__(d_1253_v27_, p1)
            d_1254_v28_ = nw193_
            (self).f18 = False
            d_1255_v30_: _dafny.Array
            def lambda47_(d_1256_v1_):
                def lambda48_(d_1257_i2_):
                    def iife117_():
                        coll53_ = _dafny.Set()
                        compr_53_: int
                        for compr_53_ in _dafny.IntegerRange(774, 864):
                            d_1258_v29_: int = compr_53_
                            if ((774) <= (d_1258_v29_)) and ((d_1258_v29_) < (864)):
                                coll53_ = coll53_.union(_dafny.Set([default__.safeDivisionInt(d_1258_v29_, len(d_1256_v1_))]))
                        return _dafny.Set(coll53_)
                    return iife117_()
                    

                return lambda48_

            init27_ = lambda47_(d_1223_v1_)
            nw194_ = _dafny.Array(None, 17)
            for i0_27_ in range(nw194_.length(0)):
                nw194_[i0_27_] = init27_(i0_27_)
            d_1255_v30_ = nw194_
            d_1259_v31_: _dafny.Array
            nw195_ = _dafny.Array(None, 7)
            nw195_[int(0)] = d_1255_v30_
            nw195_[int(1)] = d_1255_v30_
            nw195_[int(2)] = d_1255_v30_
            nw195_[int(3)] = d_1255_v30_
            nw195_[int(4)] = d_1255_v30_
            nw195_[int(5)] = d_1255_v30_
            nw195_[int(6)] = d_1255_v30_
            d_1259_v31_ = nw195_
            index142_ = default__.safeIndex(680, (d_1259_v31_).length(0))
            (d_1259_v31_)[index142_] = d_1255_v30_
            index143_ = default__.safeIndex(680, (d_1259_v31_).length(0))
            (d_1259_v31_)[index143_] = d_1255_v30_
            d_1260_v32_: _dafny.Array
            def lambda49_(d_1261_i3_):
                return (d_1261_i3_) - (641)

            init28_ = lambda49_
            nw196_ = _dafny.Array(None, 13)
            for i0_28_ in range(nw196_.length(0)):
                nw196_[i0_28_] = init28_(i0_28_)
            d_1260_v32_ = nw196_
            d_1260_v32_ = d_1260_v32_
            d_1262_v33_: _dafny.Seq
            d_1262_v33_ = _dafny.SeqWithoutIsStrInference([d_1227_v2_, d_1227_v2_])
            d_1262_v33_ = d_1262_v33_
        r0 = d_1227_v2_
        r1 = ((self).f23) - ((self).f23)
        def iife118_(_pat_let32_0):
            def iife119_(d_1263_dt__update__tmp_h1_):
                def iife120_(_pat_let33_0):
                    def iife121_(d_1264_dt__update_hcf7_h0_):
                        def iife122_(_pat_let34_0):
                            def iife123_(d_1265_dt__update_hcf6_h1_):
                                return D0_DC1((d_1263_dt__update__tmp_h1_).cf5, d_1265_dt__update_hcf6_h1_, d_1264_dt__update_hcf7_h0_, (d_1263_dt__update__tmp_h1_).cf8)
                            return iife123_(_pat_let34_0)
                        return iife122_((self).f16)
                    return iife121_(_pat_let33_0)
                return iife120_(len((((self).f17).set(default__.safeIndex((self).f5, len((self).f17)), _dafny.CodePoint('k'))) + ((self).f17)))
            return iife119_(_pat_let32_0)
        r2 = iife118_(d_1222_v0_)
        r3 = (self).f15
        return r0, r1, r2, r3

    def m16(self, p0, p1, p2, globalState):
        d_1266_v0_: _dafny.Map
        d_1266_v0_ = _dafny.Map({(self).f16: p1})
        (self).f18 = not((p1) < (((d_1266_v0_)[(self).f4] if ((self).f4) in (d_1266_v0_) else (self).f23)))
        d_1267_v1_: int
        d_1267_v1_ = -783
        d_1268_v2_: _dafny.Array
        nw197_ = _dafny.Array(False, 28)
        d_1268_v2_ = nw197_
        d_1269_v3_: _dafny.MultiSet
        d_1269_v3_ = _dafny.MultiSet([d_1268_v2_])
        d_1270_v4_: D7
        d_1270_v4_ = D7_DC18(d_1269_v3_)
        d_1271_v5_: D0
        d_1271_v5_ = D0_DC0((self).f16, (self).f15, (self).f16, (self).f5, (self).f24)
        d_1272_v7_: _dafny.Map
        def iife124_():
            coll54_ = _dafny.Set()
            compr_54_: int
            for compr_54_ in _dafny.IntegerRange(-887, -674):
                d_1273_v6_: int = compr_54_
                if ((-887) <= (d_1273_v6_)) and ((d_1273_v6_) < (-674)):
                    coll54_ = coll54_.union(_dafny.Set([(d_1273_v6_) + (p1)]))
            return _dafny.Set(coll54_)
        d_1272_v7_ = _dafny.Map({(default__.fm32(d_1271_v5_, globalState)) * (746): iife124_()
        })
        d_1274_v8_: _dafny.Set
        d_1274_v8_ = _dafny.Set({d_1267_v1_, (self).f23, -672, p1, (_dafny.MultiSet([p1])).cardinality})
        rhs137_ = len(((d_1272_v7_)[len(d_1274_v8_)] if (len(d_1274_v8_)) in (d_1272_v7_) else (d_1274_v8_) | (d_1274_v8_)))
        rhs138_ = d_1270_v4_
        d_1267_v1_ = rhs137_
        d_1270_v4_ = rhs138_
        d_1267_v1_ = d_1267_v1_
        d_1275_v9_: _dafny.MultiSet
        d_1275_v9_ = _dafny.MultiSet([(self).f23])
        d_1276_v10_: C2
        nw198_ = C2()
        nw198_.ctor__(d_1268_v2_, d_1275_v9_, (self).f17, (self).f4, (self).f16, (0) - (-251))
        d_1276_v10_ = nw198_
        d_1277_v11_: C1
        nw199_ = C1()
        nw199_.ctor__((self).f23, ((self).f23) == (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "priockxi")))), p0, (self).f5)
        d_1277_v11_ = nw199_
        d_1278_v12_: _dafny.Array
        nw200_ = _dafny.Array(_dafny.Seq({}), 21)
        d_1278_v12_ = nw200_
        d_1279_v13_: _dafny.Seq
        d_1279_v13_ = _dafny.SeqWithoutIsStrInference([self.f18])
        index144_ = default__.safeIndex(82, (d_1278_v12_).length(0))
        (d_1278_v12_)[index144_] = d_1279_v13_
        index145_ = default__.safeIndex(82, (d_1278_v12_).length(0))
        (d_1278_v12_)[index145_] = (d_1279_v13_) + ((d_1279_v13_).set(default__.safeIndex(d_1267_v1_, len(d_1279_v13_)), (self).f4))

    @property
    def f23(self):
        return self._f23
    @property
    def f24(self):
        return self._f24

class C5:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self):
        pass
        pass

    def m13(self, p0, p1, globalState):
        d_1280_v0_: bool
        d_1280_v0_ = True
        d_1280_v0_ = (d_1280_v0_ if d_1280_v0_ else d_1280_v0_)
        d_1281_v1_: int
        d_1281_v1_ = 210
        d_1281_v1_ = default__.fm26(globalState)
        d_1282_v2_: str
        d_1282_v2_ = _dafny.CodePoint('i')
        if default__.fm27(d_1282_v2_, globalState):
            d_1280_v0_ = d_1280_v0_
            if d_1280_v0_:
                d_1283_v3_: _dafny.Array
                nw201_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_1283_v3_ = nw201_
                d_1284_v4_: _dafny.Array
                def lambda50_(d_1285_p0_):
                    def lambda51_(d_1286_i0_):
                        return default__.safeDivisionInt(d_1286_i0_, d_1285_p0_)

                    return lambda51_

                init29_ = lambda50_(p0)
                nw202_ = _dafny.Array(None, 25)
                for i0_29_ in range(nw202_.length(0)):
                    nw202_[i0_29_] = init29_(i0_29_)
                d_1284_v4_ = nw202_
                index146_ = default__.safeIndex(45, (d_1283_v3_).length(0))
                (d_1283_v3_)[index146_] = d_1284_v4_
                d_1287_v5_: _dafny.Set
                d_1287_v5_ = _dafny.Set({p0, (0) - (p0), p0, p0})
                d_1288_v6_: _dafny.Seq
                d_1288_v6_ = _dafny.SeqWithoutIsStrInference([d_1280_v0_, d_1280_v0_, not(d_1280_v0_)])
                index147_ = default__.safeIndex(45, (d_1283_v3_).length(0))
                rhs139_ = d_1280_v0_
                rhs140_ = d_1284_v4_
                rhs141_ = d_1287_v5_
                rhs142_ = (0) - (p0)
                rhs143_ = d_1288_v6_
                lhs62_ = d_1283_v3_
                lhs63_ = default__.safeIndex(45, (d_1283_v3_).length(0))
                d_1280_v0_ = rhs139_
                lhs62_[lhs63_] = rhs140_
                d_1287_v5_ = rhs141_
                d_1281_v1_ = rhs142_
                d_1288_v6_ = rhs143_
                d_1282_v2_ = d_1282_v2_
                d_1289_v7_: _dafny.Set
                d_1289_v7_ = _dafny.Set({d_1280_v0_, d_1280_v0_})
                (globalState).f0 = (d_1289_v7_) | (d_1289_v7_)
                d_1281_v1_ = default__.safeDivisionInt(p0, default__.fm26(globalState))
                d_1290_v8_: _dafny.Array
                def lambda52_(d_1291_v0_):
                    def lambda53_(d_1292_i1_):
                        return d_1291_v0_

                    return lambda53_

                init30_ = lambda52_(d_1280_v0_)
                nw203_ = _dafny.Array(None, 20)
                for i0_30_ in range(nw203_.length(0)):
                    nw203_[i0_30_] = init30_(i0_30_)
                d_1290_v8_ = nw203_
                d_1293_v9_: _dafny.MultiSet
                d_1293_v9_ = _dafny.MultiSet([d_1290_v8_, d_1290_v8_])
                d_1294_v10_: D7
                d_1294_v10_ = D7_DC18(d_1293_v9_)
                d_1280_v0_ = ((d_1293_v9_).intersection((d_1294_v10_).cf40)).ispropersubset(_dafny.MultiSet([d_1290_v8_]))
            elif True:
                d_1281_v1_ = p0
                d_1295_v11_: _dafny.Seq
                d_1295_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aalcxeg"))
                d_1296_v12_: D1
                d_1296_v12_ = D1_DC5(not(d_1280_v0_), d_1281_v1_, d_1295_v11_)
                d_1297_v13_: _dafny.Map
                d_1297_v13_ = _dafny.Map({len((d_1296_v12_).cf17): d_1295_v11_})
                d_1297_v13_ = d_1297_v13_
                d_1281_v1_ = (default__.safeDivisionInt(d_1281_v1_, p0)) - (p0)
                d_1298_v14_: _dafny.Map
                d_1298_v14_ = _dafny.Map({361: d_1281_v1_})
                d_1281_v1_ = default__.safeModuloInt((0) - (len(d_1298_v14_)), 212)
                d_1280_v0_ = (d_1281_v1_) <= (d_1281_v1_)
            d_1280_v0_ = True
            d_1299_v15_: _dafny.Set
            d_1299_v15_ = _dafny.Set({_dafny.Map({151: d_1281_v1_})})
            d_1300_v16_: _dafny.Array
            nw204_ = _dafny.Array(int(0), 20)
            d_1300_v16_ = nw204_
            d_1301_v17_: _dafny.Set
            d_1301_v17_ = _dafny.Set({d_1300_v16_})
            d_1302_v18_: _dafny.Seq
            d_1302_v18_ = _dafny.SeqWithoutIsStrInference([d_1301_v17_])
            d_1303_v19_: _dafny.Map
            d_1303_v19_ = _dafny.Map({d_1281_v1_: d_1299_v15_})
            d_1304_v20_: _dafny.Seq
            d_1304_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bx"))
            d_1305_v21_: _dafny.Set
            d_1305_v21_ = d_1299_v15_
            d_1306_v22_: _dafny.Set
            d_1306_v22_ = (d_1305_v21_)
            def iife125_():
                coll55_ = _dafny.Set()
                compr_55_: _dafny.Map
                for compr_55_ in ((d_1306_v22_)).Elements:
                    d_1307_v23_: _dafny.Map = compr_55_
                    if (d_1307_v23_) in ((d_1306_v22_)):
                        coll55_ = coll55_.union(_dafny.Set([d_1307_v23_]))
                return _dafny.Set(coll55_)
            d_1299_v15_ = (((d_1303_v19_)[len(d_1304_v20_)] if (len(d_1304_v20_)) in (d_1303_v19_) else iife125_()
            ) if not((p0) >= (len((d_1302_v18_)[default__.safeIndex(p0, len(d_1302_v18_))]))) else d_1299_v15_)
            d_1308_v24_: _dafny.Map
            d_1308_v24_ = _dafny.Map({default__.fm26(globalState): 221})
            index148_ = default__.safeIndex(225, (d_1300_v16_).length(0))
            (d_1300_v16_)[index148_] = default__.fm26(globalState)
            index149_ = default__.safeIndex(958, (d_1300_v16_).length(0))
            (d_1300_v16_)[index149_] = p0
            d_1309_v25_: _dafny.Seq
            d_1309_v25_ = _dafny.SeqWithoutIsStrInference([d_1281_v1_])
            d_1310_v26_: _dafny.Seq
            d_1310_v26_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({d_1281_v1_: p0})).set(len((d_1309_v25_).set(default__.safeIndex(859, len(d_1309_v25_)), len(_dafny.Map({p0: d_1280_v0_})))), len(d_1308_v24_)), d_1308_v24_])
            d_1311_v27_: _dafny.Set
            d_1311_v27_ = _dafny.Set({d_1280_v0_})
            index150_ = default__.safeIndex(225, (d_1300_v16_).length(0))
            index151_ = default__.safeIndex(958, (d_1300_v16_).length(0))
            rhs144_ = (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjxlrdmc"))): (0) - (len(d_1304_v20_))})) | (((d_1310_v26_)[default__.safeIndex(d_1281_v1_, len(d_1310_v26_))]).set(default__.fm26(globalState), d_1281_v1_))
            rhs145_ = default__.fm26(globalState)
            rhs146_ = (d_1311_v27_) | (d_1311_v27_)
            rhs147_ = p0
            rhs148_ = _dafny.CodePoint('h')
            lhs64_ = d_1300_v16_
            lhs65_ = default__.safeIndex(225, (d_1300_v16_).length(0))
            lhs66_ = globalState
            lhs67_ = d_1300_v16_
            lhs68_ = default__.safeIndex(958, (d_1300_v16_).length(0))
            d_1308_v24_ = rhs144_
            lhs64_[lhs65_] = rhs145_
            lhs66_.f0 = rhs146_
            lhs67_[lhs68_] = rhs147_
            d_1282_v2_ = rhs148_
        elif True:
            d_1280_v0_ = d_1280_v0_
            d_1312_v28_: _dafny.Array
            nw205_ = _dafny.Array(False, 13)
            d_1312_v28_ = nw205_
            d_1313_v29_: _dafny.MultiSet
            d_1313_v29_ = _dafny.MultiSet([d_1312_v28_, d_1312_v28_])
            d_1314_v30_: D7
            d_1314_v30_ = D7_DC18(d_1313_v29_)
            d_1315_v31_: _dafny.Seq
            d_1315_v31_ = _dafny.SeqWithoutIsStrInference([d_1314_v30_, d_1314_v30_])
            pat_let_tv18_ = d_1313_v29_
            d_1316_v32_: _dafny.Map
            def iife126_(_pat_let35_0):
                def iife127_(d_1317_dt__update__tmp_h0_):
                    def iife128_(_pat_let36_0):
                        def iife129_(d_1318_dt__update_hcf40_h0_):
                            return D7_DC18(d_1318_dt__update_hcf40_h0_)
                        return iife129_(_pat_let36_0)
                    return iife128_(pat_let_tv18_)
                return iife127_(_pat_let35_0)
            d_1316_v32_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([iife126_(d_1314_v30_)])) < (d_1315_v31_): d_1280_v0_})
            d_1319_v33_: _dafny.Seq
            d_1319_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqfowxgg"))
            d_1320_v34_: D1
            d_1320_v34_ = D1_DC5(d_1280_v0_, d_1281_v1_, d_1319_v33_)
            d_1321_v35_: _dafny.Map
            d_1321_v35_ = _dafny.Map({d_1320_v34_: d_1280_v0_})
            d_1316_v32_ = ((d_1316_v32_ if d_1280_v0_ else d_1316_v32_)).set(((d_1321_v35_)[D1_DC5(d_1280_v0_, p0, d_1319_v33_)] if (D1_DC5(d_1280_v0_, p0, d_1319_v33_)) in (d_1321_v35_) else d_1280_v0_), d_1280_v0_)
            d_1322_v36_: _dafny.Set
            d_1322_v36_ = _dafny.Set({d_1280_v0_, d_1280_v0_, d_1280_v0_})
            d_1323_v38_: _dafny.Map
            def iife130_():
                coll56_ = _dafny.Map()
                compr_56_: int
                for compr_56_ in _dafny.IntegerRange(399, 136):
                    d_1324_v37_: int = compr_56_
                    if ((399) <= (d_1324_v37_)) and ((d_1324_v37_) < (136)):
                        coll56_[(d_1324_v37_) * (len(d_1319_v33_))] = 899
                return _dafny.Map(coll56_)
            d_1323_v38_ = _dafny.Map({d_1322_v36_: iife130_()
            })
            d_1325_v39_: _dafny.Map
            d_1325_v39_ = _dafny.Map({d_1280_v0_: d_1323_v38_})
            d_1323_v38_ = ((d_1323_v38_ if d_1280_v0_ else ((d_1325_v39_)[d_1280_v0_] if (d_1280_v0_) in (d_1325_v39_) else d_1323_v38_))).set(d_1322_v36_, _dafny.Map({846: d_1281_v1_}))
            d_1326_v40_: _dafny.Map
            d_1326_v40_ = _dafny.Map({d_1280_v0_: p0})
            d_1327_v41_: _dafny.Set
            d_1327_v41_ = _dafny.Set({d_1319_v33_, default__.fm28(_dafny.SeqWithoutIsStrInference([d_1282_v2_ for d_1328_i2_ in range(default__.abs(-288))]), ((d_1326_v40_)[True] if (True) in (d_1326_v40_) else d_1281_v1_), d_1281_v1_, globalState)})
            d_1329_v42_: D5
            d_1329_v42_ = D5_DC12((_dafny.Set({d_1319_v33_, d_1319_v33_})).intersection(d_1327_v41_))
            d_1329_v42_ = d_1329_v42_
            d_1330_v43_: _dafny.Array
            nw206_ = _dafny.Array(_dafny.CodePoint('D'), 16)
            d_1330_v43_ = nw206_
            nw207_ = _dafny.Array(_dafny.CodePoint('D'), 17)
            d_1330_v43_ = nw207_
        d_1331_v44_: _dafny.Set
        d_1331_v44_ = _dafny.Set({False})
        d_1332_v45_: _dafny.Seq
        d_1332_v45_ = _dafny.SeqWithoutIsStrInference([d_1281_v1_, len(d_1331_v44_)])
        d_1333_v46_: _dafny.MultiSet
        d_1333_v46_ = _dafny.MultiSet([d_1280_v0_])
        d_1334_v47_: _dafny.Set
        d_1334_v47_ = _dafny.Set({d_1281_v1_, len(_dafny.Set({d_1280_v0_})), len((d_1332_v45_).set(default__.safeIndex(((d_1333_v46_)[d_1280_v0_] if (d_1280_v0_) in (d_1333_v46_) else 4), len(d_1332_v45_)), d_1281_v1_))})
        if (d_1334_v47_).ispropersubset((d_1334_v47_).intersection(d_1334_v47_)):
            d_1281_v1_ = default__.safeModuloInt((default__.fm26(globalState)) + (len(_dafny.SeqWithoutIsStrInference([len(d_1331_v44_) for d_1335_i3_ in range(default__.abs(-616))]))), 783)
            d_1336_v48_: _dafny.Map
            d_1336_v48_ = _dafny.Map({not(not(d_1280_v0_)): 501})
            d_1337_v49_: _dafny.Seq
            d_1337_v49_ = _dafny.SeqWithoutIsStrInference([d_1336_v48_])
            d_1338_v50_: _dafny.Map
            d_1338_v50_ = _dafny.Map({(d_1337_v49_) + (d_1337_v49_): len(d_1331_v44_)})
            d_1338_v50_ = (d_1338_v50_).set(d_1337_v49_, default__.safeDivisionInt(p0, p0))
            d_1339_v51_: _dafny.Seq
            d_1339_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrucpv"))
            d_1340_v52_: _dafny.Seq
            d_1340_v52_ = _dafny.SeqWithoutIsStrInference([d_1280_v0_, d_1280_v0_])
            d_1341_v53_: _dafny.Array
            nw208_ = _dafny.Array(None, 1)
            nw208_[int(0)] = (d_1333_v46_).cardinality
            d_1341_v53_ = nw208_
            d_1342_v54_: _dafny.Array
            nw209_ = _dafny.Array(None, 5)
            nw209_[int(0)] = (_dafny.Map({d_1280_v0_: len(d_1339_v51_)})) | (d_1336_v48_)
            nw209_[int(1)] = d_1336_v48_
            nw209_[int(2)] = (d_1336_v48_) | ((d_1336_v48_).set(d_1280_v0_, (D5_DC14(d_1280_v0_, d_1340_v52_, d_1341_v53_, p0, len(d_1334_v47_))).cf35))
            nw209_[int(3)] = d_1336_v48_
            nw209_[int(4)] = d_1336_v48_
            d_1342_v54_ = nw209_
            index152_ = default__.safeIndex(818, (d_1342_v54_).length(0))
            (d_1342_v54_)[index152_] = d_1336_v48_
            d_1343_v55_: _dafny.Set
            d_1343_v55_ = _dafny.Set({d_1341_v53_})
            index153_ = default__.safeIndex(818, (d_1342_v54_).length(0))
            (d_1342_v54_)[index153_] = _dafny.Map({(p0) != (p0): default__.safeModuloInt((0) - (len(d_1343_v55_)), d_1281_v1_)})
            d_1280_v0_ = d_1280_v0_
            d_1344_v56_: _dafny.Array
            nw210_ = _dafny.Array(_dafny.Set({}), 12)
            d_1344_v56_ = nw210_
            d_1345_v57_: D4
            d_1345_v57_ = D4_DC9(d_1336_v48_, len(_dafny.SeqWithoutIsStrInference([d_1280_v0_])), d_1280_v0_)
            d_1346_v58_: _dafny.Set
            d_1346_v58_ = _dafny.Set({d_1345_v57_, d_1345_v57_})
            d_1347_v59_: _dafny.MultiSet
            d_1347_v59_ = _dafny.MultiSet([d_1345_v57_, d_1345_v57_])
            def iife131_():
                coll57_ = _dafny.Set()
                compr_57_: D4
                for compr_57_ in (d_1347_v59_).Elements:
                    d_1348_v60_: D4 = compr_57_
                    if (d_1348_v60_) in (d_1347_v59_):
                        coll57_ = coll57_.union(_dafny.Set([d_1348_v60_]))
                return _dafny.Set(coll57_)
            rhs149_ = d_1344_v56_
            rhs150_ = not((d_1280_v0_) and (d_1280_v0_))
            rhs151_ = (d_1331_v44_).issubset(d_1331_v44_)
            rhs152_ = (iife131_()
            ).ispropersubset((default__.fm29(_dafny.CodePoint('k'), p0, d_1280_v0_, globalState)) | (d_1346_v58_))
            d_1344_v56_ = rhs149_
            d_1280_v0_ = rhs150_
            d_1280_v0_ = rhs151_
            d_1280_v0_ = rhs152_
        elif True:
            d_1349_v61_: _dafny.Seq
            d_1349_v61_ = _dafny.SeqWithoutIsStrInference([d_1280_v0_])
            d_1350_v62_: _dafny.Array
            nw211_ = _dafny.Array(None, 17)
            nw211_[int(0)] = d_1349_v61_
            nw211_[int(1)] = d_1349_v61_
            nw211_[int(2)] = d_1349_v61_
            nw211_[int(3)] = d_1349_v61_
            nw211_[int(4)] = d_1349_v61_
            nw211_[int(5)] = d_1349_v61_
            nw211_[int(6)] = d_1349_v61_
            nw211_[int(7)] = d_1349_v61_
            nw211_[int(8)] = d_1349_v61_
            nw211_[int(9)] = (d_1349_v61_) + (d_1349_v61_)
            nw211_[int(10)] = d_1349_v61_
            nw211_[int(11)] = d_1349_v61_
            nw211_[int(12)] = (d_1349_v61_).set(default__.safeIndex(181, len(d_1349_v61_)), True)
            nw211_[int(13)] = d_1349_v61_
            nw211_[int(14)] = d_1349_v61_
            nw211_[int(15)] = d_1349_v61_
            nw211_[int(16)] = default__.fm30(p0, globalState)
            d_1350_v62_ = nw211_
            index154_ = default__.safeIndex(352, (d_1350_v62_).length(0))
            (d_1350_v62_)[index154_] = d_1349_v61_
            index155_ = default__.safeIndex(352, (d_1350_v62_).length(0))
            (d_1350_v62_)[index155_] = _dafny.SeqWithoutIsStrInference([d_1280_v0_, d_1280_v0_])
            d_1351_v63_: C1
            nw212_ = C1()
            nw212_.ctor__(d_1281_v1_, d_1280_v0_, d_1280_v0_, p0)
            d_1351_v63_ = nw212_
            d_1352_v64_: _dafny.Seq
            d_1352_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdaxhmngh"))
            d_1353_v65_: _dafny.Seq
            d_1353_v65_ = _dafny.SeqWithoutIsStrInference([d_1352_v64_])
            d_1280_v0_ = (((d_1353_v65_)[default__.safeIndex(p0, len(d_1353_v65_))]) + (d_1352_v64_)) != ((d_1352_v64_) + (d_1352_v64_))
            d_1281_v1_ = (0) - ((p0) - (p0))
            d_1354_v66_: _dafny.Array
            nw213_ = _dafny.Array(False, 14)
            d_1354_v66_ = nw213_
            d_1355_v67_: _dafny.MultiSet
            d_1355_v67_ = _dafny.MultiSet([len(d_1334_v47_), 104])
            index156_ = default__.safeIndex(825, (d_1354_v66_).length(0))
            (d_1354_v66_)[index156_] = (d_1352_v64_) == ((default__.fm28(default__.fm28(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yb")), 294, p0, globalState), len(d_1352_v64_), ((d_1355_v67_)[d_1281_v1_] if (d_1281_v1_) in (d_1355_v67_) else p0), globalState)).set(default__.safeIndex((0) - (((d_1333_v46_)[d_1280_v0_] if (d_1280_v0_) in (d_1333_v46_) else len(d_1352_v64_))), len(default__.fm28(default__.fm28(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yb")), 294, p0, globalState), len(d_1352_v64_), ((d_1355_v67_)[d_1281_v1_] if (d_1281_v1_) in (d_1355_v67_) else p0), globalState))), _dafny.CodePoint('g')))
            index157_ = default__.safeIndex(825, (d_1354_v66_).length(0))
            (d_1354_v66_)[index157_] = (d_1355_v67_).ispropersubset((d_1355_v67_).set(p0, default__.abs(p0)))
        d_1356_v68_: _dafny.Array
        nw214_ = _dafny.Array(_dafny.Map({}), 13)
        d_1356_v68_ = nw214_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1356_v68_).length(0)):
            d_1357_i4_: int = guard_loop_3_
            if (True) and (((0) <= (d_1357_i4_)) and ((d_1357_i4_) < ((d_1356_v68_).length(0)))):
                (d_1356_v68_)[(d_1357_i4_)] = ((_dafny.Map({p0: 813})) | (_dafny.Map({p0: d_1281_v1_}))) | (_dafny.Map({d_1281_v1_: -503}))
        hi4_ = d_1281_v1_
        for d_1358_i5_ in range((0) - (p0), hi4_):
            if d_1280_v0_:
                d_1280_v0_ = d_1280_v0_
                d_1359_v69_: _dafny.Seq
                d_1359_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
                d_1360_v70_: _dafny.Seq
                d_1360_v70_ = _dafny.SeqWithoutIsStrInference([d_1280_v0_, d_1280_v0_])
                d_1361_v71_: C4
                nw215_ = C4()
                nw215_.ctor__(len(d_1359_v69_), (d_1360_v70_)[default__.safeIndex(410, len(d_1360_v70_))], d_1359_v69_, d_1280_v0_, d_1280_v0_, d_1358_i5_, len(((d_1332_v45_).set(default__.safeIndex(d_1281_v1_, len(d_1332_v45_)), d_1281_v1_)).set(default__.safeIndex(-402, len((d_1332_v45_).set(default__.safeIndex(d_1281_v1_, len(d_1332_v45_)), d_1281_v1_))), len(d_1331_v44_))), not(d_1280_v0_))
                d_1361_v71_ = nw215_
                d_1362_v72_: _dafny.Seq
                d_1362_v72_ = _dafny.SeqWithoutIsStrInference([d_1332_v45_])
                d_1363_v73_: _dafny.Map
                d_1363_v73_ = _dafny.Map({(d_1362_v72_)[default__.safeIndex(d_1281_v1_, len(d_1362_v72_))]: len(d_1359_v69_)})
                d_1363_v73_ = d_1363_v73_
                d_1364_v74_: D4
                d_1364_v74_ = D4_DC9(_dafny.Map({d_1280_v0_: p0}), p0, (d_1361_v71_).f24)
                d_1365_v75_: _dafny.Map
                d_1365_v75_ = _dafny.Map({d_1281_v1_: d_1358_i5_})
                d_1366_v76_: _dafny.MultiSet
                d_1366_v76_ = _dafny.MultiSet([default__.fm26(globalState), (d_1333_v46_).cardinality, p0])
                d_1367_v77_: _dafny.Array
                nw216_ = _dafny.Array(None, 28)
                nw216_[int(0)] = p0
                nw216_[int(1)] = len(d_1359_v69_)
                nw216_[int(2)] = default__.safeDivisionInt((0) - ((0) - (len(d_1359_v69_))), ((d_1365_v75_)[(d_1361_v71_).f23] if ((d_1361_v71_).f23) in (d_1365_v75_) else d_1281_v1_))
                nw216_[int(3)] = p0
                nw216_[int(4)] = d_1281_v1_
                nw216_[int(5)] = ((d_1361_v71_).f23 if (d_1361_v71_).f24 else (0) - (p0))
                nw216_[int(6)] = (d_1361_v71_).f23
                nw216_[int(7)] = ((d_1361_v71_).f23) - (d_1358_i5_)
                nw216_[int(8)] = d_1281_v1_
                nw216_[int(9)] = (d_1361_v71_).f23
                nw216_[int(10)] = 635
                nw216_[int(11)] = 787
                nw216_[int(12)] = (d_1281_v1_) - (446)
                nw216_[int(13)] = d_1281_v1_
                nw216_[int(14)] = ((d_1366_v76_)[d_1358_i5_] if (d_1358_i5_) in (d_1366_v76_) else p0)
                nw216_[int(15)] = ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([(d_1361_v71_).f24])))) if (d_1361_v71_).f24 else d_1281_v1_)
                nw216_[int(16)] = (d_1361_v71_).f23
                nw216_[int(17)] = p0
                nw216_[int(18)] = d_1281_v1_
                nw216_[int(19)] = (default__.fm55((d_1361_v71_).f23, d_1360_v70_, d_1281_v1_, globalState)) - (d_1358_i5_)
                nw216_[int(20)] = ((d_1361_v71_).f23) - ((d_1361_v71_).f23)
                nw216_[int(21)] = len((d_1332_v45_) + (d_1332_v45_))
                nw216_[int(22)] = default__.fm49(d_1359_v69_, p0, False, (d_1361_v71_).f24, globalState)
                nw216_[int(23)] = len(d_1332_v45_)
                nw216_[int(24)] = 629
                nw216_[int(25)] = (p0 if (d_1361_v71_).f24 else 470)
                nw216_[int(26)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhvybudgo")))
                nw216_[int(27)] = d_1281_v1_
                d_1367_v77_ = nw216_
                index158_ = default__.safeIndex(855, (d_1367_v77_).length(0))
                (d_1367_v77_)[index158_] = d_1358_i5_
                index159_ = default__.safeIndex(855, (d_1367_v77_).length(0))
                rhs153_ = d_1364_v74_
                rhs154_ = (d_1361_v71_).f24
                rhs155_ = not((d_1361_v71_).f24)
                rhs156_ = d_1281_v1_
                lhs69_ = d_1367_v77_
                lhs70_ = default__.safeIndex(855, (d_1367_v77_).length(0))
                d_1364_v74_ = rhs153_
                d_1280_v0_ = rhs154_
                d_1280_v0_ = rhs155_
                lhs69_[lhs70_] = rhs156_
                index160_ = default__.safeIndex(855, (d_1367_v77_).length(0))
                (d_1367_v77_)[index160_] = d_1281_v1_
            elif True:
                d_1281_v1_ = default__.fm49(_dafny.SeqWithoutIsStrInference([d_1282_v2_ for d_1368_i6_ in range(default__.abs(0))]), d_1358_i5_, not(False), (d_1280_v0_ if False else d_1280_v0_), globalState)
                d_1281_v1_ = d_1281_v1_
                d_1280_v0_ = not (d_1280_v0_) or (d_1280_v0_)
                d_1369_v78_: _dafny.Map
                d_1369_v78_ = _dafny.Map({_dafny.Set({d_1358_i5_}): d_1280_v0_})
                d_1370_v79_: _dafny.Map
                d_1370_v79_ = _dafny.Map({p0: d_1280_v0_})
                d_1371_v81_: _dafny.Map
                d_1371_v81_ = _dafny.Map({d_1280_v0_: d_1280_v0_})
                def iife132_():
                    coll58_ = _dafny.Set()
                    compr_58_: int
                    for compr_58_ in (d_1370_v79_).keys.Elements:
                        d_1372_v80_: int = compr_58_
                        if (d_1372_v80_) in (d_1370_v79_):
                            coll58_ = coll58_.union(_dafny.Set([default__.safeModuloInt(d_1372_v80_, 848)]))
                    return _dafny.Set(coll58_)
                def iife133_():
                    coll59_ = _dafny.Set()
                    compr_59_: int
                    for compr_59_ in (d_1370_v79_).keys.Elements:
                        d_1373_v80_: int = compr_59_
                        if (d_1373_v80_) in (d_1370_v79_):
                            coll59_ = coll59_.union(_dafny.Set([default__.safeModuloInt(d_1373_v80_, 848)]))
                    return _dafny.Set(coll59_)
                d_1280_v0_ = not(((d_1369_v78_)[iife132_()
                ] if (iife133_()
                ) in (d_1369_v78_) else (len(d_1371_v81_)) < (d_1281_v1_)))
                d_1374_v82_: _dafny.Array
                def lambda54_(d_1375_i7_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxeby"))

                init31_ = lambda54_
                nw217_ = _dafny.Array(None, 16)
                for i0_31_ in range(nw217_.length(0)):
                    nw217_[i0_31_] = init31_(i0_31_)
                d_1374_v82_ = nw217_
                d_1376_v83_: _dafny.Array
                def lambda55_(d_1377_v0_):
                    def lambda56_(d_1378_i8_):
                        return d_1377_v0_

                    return lambda56_

                init32_ = lambda55_(d_1280_v0_)
                nw218_ = _dafny.Array(None, 1)
                for i0_32_ in range(nw218_.length(0)):
                    nw218_[i0_32_] = init32_(i0_32_)
                d_1376_v83_ = nw218_
                d_1379_v84_: _dafny.Map
                d_1379_v84_ = _dafny.Map({d_1374_v82_: d_1376_v83_})
                d_1379_v84_ = (d_1379_v84_).set(d_1374_v82_, d_1376_v83_)
            d_1380_v85_: _dafny.Seq
            d_1380_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulcu"))
            d_1380_v85_ = default__.fm28(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flt")), p0, (0) - ((d_1332_v45_)[default__.safeIndex(d_1358_i5_, len(d_1332_v45_))]), globalState)
            if (not (d_1280_v0_) or (not(d_1280_v0_))) and (d_1280_v0_):
                d_1280_v0_ = (True if d_1280_v0_ else d_1280_v0_)
                d_1381_v86_: _dafny.Seq
                d_1381_v86_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsfss"))])
                d_1382_v87_: _dafny.Seq
                d_1382_v87_ = _dafny.SeqWithoutIsStrInference([d_1280_v0_])
                d_1383_v88_: _dafny.Array
                nw219_ = _dafny.Array(None, 19)
                nw219_[int(0)] = len(d_1331_v44_)
                nw219_[int(1)] = (0) - (d_1358_i5_)
                nw219_[int(2)] = p0
                nw219_[int(3)] = d_1358_i5_
                nw219_[int(4)] = len((d_1381_v86_) + (d_1381_v86_))
                nw219_[int(5)] = d_1281_v1_
                nw219_[int(6)] = p0
                nw219_[int(7)] = len(_dafny.SeqWithoutIsStrInference([d_1282_v2_ for d_1384_i9_ in range(default__.abs(558))]))
                nw219_[int(8)] = d_1281_v1_
                nw219_[int(9)] = (d_1281_v1_) * (p0)
                nw219_[int(10)] = p0
                nw219_[int(11)] = d_1281_v1_
                nw219_[int(12)] = len(d_1382_v87_)
                nw219_[int(13)] = (716) + (p0)
                nw219_[int(14)] = 103
                nw219_[int(15)] = p0
                nw219_[int(16)] = p0
                nw219_[int(17)] = p0
                nw219_[int(18)] = 787
                d_1383_v88_ = nw219_
                d_1385_v89_: _dafny.Array
                nw220_ = _dafny.Array(_dafny.CodePoint('D'), 18)
                d_1385_v89_ = nw220_
                index161_ = default__.safeIndex(695, (d_1385_v89_).length(0))
                (d_1385_v89_)[index161_] = d_1282_v2_
                d_1386_v90_: _dafny.Set
                d_1386_v90_ = _dafny.Set({d_1383_v88_, d_1383_v88_, d_1383_v88_})
                d_1387_v91_: _dafny.Array
                nw221_ = _dafny.Array(None, 5)
                nw221_[int(0)] = default__.fm27(d_1282_v2_, globalState)
                nw221_[int(1)] = d_1280_v0_
                nw221_[int(2)] = d_1280_v0_
                nw221_[int(3)] = d_1280_v0_
                nw221_[int(4)] = (_dafny.CodePoint('r')) in (d_1380_v85_)
                d_1387_v91_ = nw221_
                index162_ = default__.safeIndex(864, (d_1387_v91_).length(0))
                (d_1387_v91_)[index162_] = d_1280_v0_
                d_1388_v92_: _dafny.Map
                d_1388_v92_ = _dafny.Map({d_1281_v1_: default__.fm27((d_1380_v85_)[default__.safeIndex(254, len(d_1380_v85_))], globalState)})
                index163_ = default__.safeIndex(695, (d_1385_v89_).length(0))
                index164_ = default__.safeIndex(864, (d_1387_v91_).length(0))
                rhs157_ = d_1383_v88_
                rhs158_ = d_1282_v2_
                rhs159_ = ((_dafny.Map({p0: d_1280_v0_})) | (_dafny.Map({p0: d_1280_v0_}))) == (d_1388_v92_)
                rhs160_ = d_1386_v90_
                rhs161_ = d_1280_v0_
                lhs71_ = d_1385_v89_
                lhs72_ = default__.safeIndex(695, (d_1385_v89_).length(0))
                lhs73_ = d_1387_v91_
                lhs74_ = default__.safeIndex(864, (d_1387_v91_).length(0))
                d_1383_v88_ = rhs157_
                lhs71_[lhs72_] = rhs158_
                d_1280_v0_ = rhs159_
                d_1386_v90_ = rhs160_
                lhs73_[lhs74_] = rhs161_
                d_1281_v1_ = (0) - (d_1281_v1_)
                index165_ = default__.safeIndex(864, (d_1387_v91_).length(0))
                (d_1387_v91_)[index165_] = (d_1358_i5_) == ((0) - (d_1281_v1_))
                d_1389_v93_: C2
                nw222_ = C2()
                nw222_.ctor__(d_1387_v91_, _dafny.MultiSet([d_1281_v1_, len(d_1380_v85_), p0]), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ja"))) + (d_1380_v85_), (d_1387_v91_)[default__.safeIndex(864, (d_1387_v91_).length(0))], not((d_1387_v91_)[default__.safeIndex(864, (d_1387_v91_).length(0))]), 893)
                d_1389_v93_ = nw222_
            elif True:
                d_1390_v94_: _dafny.MultiSet
                d_1391_v95_: int
                d_1392_v96_: _dafny.MultiSet
                out16_: _dafny.MultiSet
                out17_: int
                out18_: _dafny.MultiSet
                out16_, out17_, out18_ = (self).m14(globalState)
                d_1390_v94_ = out16_
                d_1391_v95_ = out17_
                d_1392_v96_ = out18_
                d_1393_v97_: _dafny.Array
                nw223_ = _dafny.Array(int(0), 12)
                d_1393_v97_ = nw223_
                index166_ = default__.safeIndex(287, (d_1393_v97_).length(0))
                (d_1393_v97_)[index166_] = d_1281_v1_
                d_1394_v98_: _dafny.Map
                d_1394_v98_ = _dafny.Map({d_1281_v1_: d_1281_v1_})
                d_1395_v99_: _dafny.Seq
                d_1395_v99_ = _dafny.SeqWithoutIsStrInference([True])
                d_1396_v100_: _dafny.Map
                d_1396_v100_ = _dafny.Map({d_1380_v85_: ((d_1394_v98_)[d_1281_v1_] if (d_1281_v1_) in (d_1394_v98_) else default__.fm55(d_1281_v1_, d_1395_v99_, d_1358_i5_, globalState))})
                index167_ = default__.safeIndex(287, (d_1393_v97_).length(0))
                (d_1393_v97_)[index167_] = ((d_1396_v100_)[d_1380_v85_] if (d_1380_v85_) in (d_1396_v100_) else 694)
                d_1397_v101_: _dafny.Array
                nw224_ = _dafny.Array(_dafny.Map({}), 1)
                d_1397_v101_ = nw224_
                d_1398_v102_: _dafny.Map
                d_1398_v102_ = _dafny.Map({d_1280_v0_: d_1397_v101_})
                d_1399_v103_: D14
                d_1399_v103_ = D14_DC36(((d_1398_v102_)[d_1280_v0_] if (d_1280_v0_) in (d_1398_v102_) else d_1397_v101_))
                d_1397_v101_ = (d_1399_v103_).cf72
                d_1400_v104_: _dafny.Array
                def lambda57_(d_1401_v2_):
                    def lambda58_(d_1402_i10_):
                        return d_1401_v2_

                    return lambda58_

                init33_ = lambda57_(d_1282_v2_)
                nw225_ = _dafny.Array(None, 13)
                for i0_33_ in range(nw225_.length(0)):
                    nw225_[i0_33_] = init33_(i0_33_)
                d_1400_v104_ = nw225_
                index168_ = default__.safeIndex(384, (d_1400_v104_).length(0))
                (d_1400_v104_)[index168_] = d_1282_v2_
                d_1403_v105_: _dafny.Map
                d_1403_v105_ = _dafny.Map({p0: False})
                index169_ = default__.safeIndex(384, (d_1400_v104_).length(0))
                rhs162_ = _dafny.CodePoint('l')
                rhs163_ = default__.safeModuloInt(len((d_1403_v105_) | (d_1403_v105_)), d_1391_v95_)
                lhs75_ = d_1400_v104_
                lhs76_ = default__.safeIndex(384, (d_1400_v104_).length(0))
                lhs75_[lhs76_] = rhs162_
                d_1391_v95_ = rhs163_
                d_1404_v106_: _dafny.Map
                d_1404_v106_ = _dafny.Map({d_1280_v0_: d_1280_v0_})
                d_1405_v107_: D0
                d_1405_v107_ = D0_DC0(d_1280_v0_, len(d_1404_v106_), d_1280_v0_, p0, d_1280_v0_)
                index170_ = default__.safeIndex(287, (d_1393_v97_).length(0))
                (d_1393_v97_)[index170_] = default__.fm32(d_1405_v107_, globalState)
            d_1406_v108_: _dafny.Map
            d_1406_v108_ = _dafny.Map({d_1280_v0_: d_1280_v0_})
            d_1406_v108_ = (d_1406_v108_).set(d_1280_v0_, not(d_1280_v0_))

    def m14(self, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: _dafny.MultiSet = _dafny.MultiSet({})
        d_1407_v0_: bool
        d_1407_v0_ = True
        d_1408_v1_: _dafny.Seq
        d_1408_v1_ = _dafny.SeqWithoutIsStrInference([d_1407_v0_, d_1407_v0_])
        d_1409_v2_: int
        d_1409_v2_ = -668
        d_1410_v3_: D9
        d_1410_v3_ = D9_DC22((d_1408_v1_)[default__.safeIndex(d_1409_v2_, len(d_1408_v1_))])
        d_1411_v4_: _dafny.Seq
        d_1411_v4_ = _dafny.SeqWithoutIsStrInference([(d_1407_v0_) or (True), True, d_1407_v0_, ((d_1410_v3_).cf44 if d_1407_v0_ else d_1407_v0_), d_1407_v0_])
        d_1412_v5_: _dafny.Seq
        d_1412_v5_ = _dafny.SeqWithoutIsStrInference([d_1409_v2_])
        d_1413_v6_: _dafny.Map
        d_1413_v6_ = _dafny.Map({True: (d_1412_v5_)[default__.safeIndex(d_1409_v2_, len(d_1412_v5_))]})
        d_1414_i0_: int
        d_1414_i0_ = 0
        with _dafny.label("3"):
            while (d_1408_v1_)[default__.safeIndex(len(d_1413_v6_), len(d_1408_v1_))]:
                with _dafny.c_label("3"):
                    if (d_1414_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_1414_i0_ = (d_1414_i0_) + (1)
                    d_1415_v7_: _dafny.Array
                    nw226_ = _dafny.Array(_dafny.Seq({}), 10)
                    d_1415_v7_ = nw226_
                    d_1416_v8_: _dafny.Seq
                    d_1416_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yauicvbc"))
                    d_1417_v9_: _dafny.Seq
                    d_1417_v9_ = _dafny.SeqWithoutIsStrInference([d_1416_v8_, d_1416_v8_])
                    index171_ = default__.safeIndex(620, (d_1415_v7_).length(0))
                    (d_1415_v7_)[index171_] = d_1417_v9_
                    index172_ = default__.safeIndex(620, (d_1415_v7_).length(0))
                    (d_1415_v7_)[index172_] = (d_1417_v9_ if not (d_1407_v0_) or (d_1407_v0_) else (_dafny.SeqWithoutIsStrInference([d_1416_v8_])) + (d_1417_v9_))
                    d_1418_v10_: _dafny.Array
                    nw227_ = _dafny.Array(int(0), 19)
                    d_1418_v10_ = nw227_
                    d_1419_v11_: D5
                    d_1419_v11_ = D5_DC14(d_1407_v0_, d_1411_v4_, d_1418_v10_, (d_1409_v2_) * (d_1409_v2_), d_1409_v2_)
                    source16_ = d_1419_v11_
                    if source16_.is_DC13:
                        d_1420___mcc_h0_ = source16_.cf29
                        d_1421___mcc_h1_ = source16_.cf30
                        d_1422_cf30_ = d_1421___mcc_h1_
                        d_1423_cf29_ = d_1420___mcc_h0_
                        r1 = 140
                        d_1424_v12_: T0
                        nw228_ = C4()
                        nw228_.ctor__((0) - (d_1409_v2_), d_1407_v0_, d_1416_v8_, d_1422_cf30_, False, len(d_1416_v8_), (d_1412_v5_)[default__.safeIndex(d_1409_v2_, len(d_1412_v5_))], d_1423_cf29_)
                        d_1424_v12_ = nw228_
                        d_1425_v13_: _dafny.MultiSet
                        d_1425_v13_ = _dafny.MultiSet([d_1424_v12_])
                        r1 = (0) - (default__.safeModuloInt(((d_1425_v13_)[d_1424_v12_] if (d_1424_v12_) in (d_1425_v13_) else d_1409_v2_), (0) - ((0) - ((d_1424_v12_).f5))))
                        d_1422_cf30_ = d_1422_cf30_
                        r1 = default__.safeModuloInt((0) - ((d_1424_v12_).f5), (d_1409_v2_) * (-933))
                    elif source16_.is_DC14:
                        d_1426___mcc_h2_ = source16_.cf31
                        d_1427___mcc_h3_ = source16_.cf32
                        d_1428___mcc_h4_ = source16_.cf33
                        d_1429___mcc_h5_ = source16_.cf34
                        d_1430___mcc_h6_ = source16_.cf35
                        d_1431_cf35_ = d_1430___mcc_h6_
                        d_1432_cf34_ = d_1429___mcc_h5_
                        d_1433_cf33_ = d_1428___mcc_h4_
                        d_1434_cf32_ = d_1427___mcc_h3_
                        d_1435_cf31_ = d_1426___mcc_h2_
                        d_1409_v2_ = default__.safeDivisionInt(d_1431_cf35_, d_1409_v2_)
                        d_1436_v14_: bool
                        out19_: bool
                        out19_ = default__.m0(d_1435_cf31_, d_1431_cf35_, globalState)
                        d_1436_v14_ = out19_
                        d_1437_v15_: str
                        d_1437_v15_ = _dafny.CodePoint('s')
                        d_1435_cf31_ = not((d_1407_v0_ if default__.fm27(d_1437_v15_, globalState) else d_1436_v14_))
                        d_1409_v2_ = d_1432_cf34_
                    elif True:
                        d_1438___mcc_h7_ = source16_.cf28
                        d_1439_cf28_ = d_1438___mcc_h7_
                        index173_ = default__.safeIndex(670, (d_1418_v10_).length(0))
                        (d_1418_v10_)[index173_] = d_1409_v2_
                        d_1440_v16_: _dafny.Set
                        d_1440_v16_ = _dafny.Set({d_1407_v0_, d_1407_v0_, not(d_1407_v0_)})
                        index174_ = default__.safeIndex(670, (d_1418_v10_).length(0))
                        (d_1418_v10_)[index174_] = (default__.fm55((0) - (d_1409_v2_), default__.fm30(len(d_1440_v16_), globalState), 108, globalState)) * (len(d_1416_v8_))
                        d_1407_v0_ = ((d_1418_v10_)[default__.safeIndex(670, (d_1418_v10_).length(0))]) <= ((d_1418_v10_)[default__.safeIndex(670, (d_1418_v10_).length(0))])
                        d_1441_v17_: _dafny.MultiSet
                        d_1441_v17_ = _dafny.MultiSet([d_1407_v0_])
                        d_1442_v18_: _dafny.Set
                        d_1442_v18_ = _dafny.Set({d_1441_v17_, default__.fm47(default__.fm27(default__.fm50(globalState), globalState), globalState), d_1441_v17_, d_1441_v17_, d_1441_v17_})
                        d_1407_v0_ = ((d_1442_v18_).ispropersubset(d_1442_v18_) if d_1407_v0_ else d_1407_v0_)
                        d_1443_v19_: str
                        d_1443_v19_ = _dafny.CodePoint('g')
                        d_1444_v20_: C3
                        nw229_ = C3()
                        nw229_.ctor__(d_1443_v19_, (0) - ((d_1418_v10_)[default__.safeIndex(670, (d_1418_v10_).length(0))]), default__.fm27(d_1443_v19_, globalState), d_1407_v0_, d_1409_v2_)
                        d_1444_v20_ = nw229_
                    d_1445_v21_: _dafny.Array
                    def lambda59_(d_1446_i1_):
                        return _dafny.SeqWithoutIsStrInference([163])

                    init34_ = lambda59_
                    nw230_ = _dafny.Array(None, 3)
                    for i0_34_ in range(nw230_.length(0)):
                        nw230_[i0_34_] = init34_(i0_34_)
                    d_1445_v21_ = nw230_
                    d_1447_v22_: _dafny.Array
                    nw231_ = _dafny.Array(None, 9)
                    nw231_[int(0)] = d_1445_v21_
                    nw231_[int(1)] = d_1445_v21_
                    nw231_[int(2)] = d_1445_v21_
                    nw231_[int(3)] = d_1445_v21_
                    nw231_[int(4)] = d_1445_v21_
                    nw231_[int(5)] = d_1445_v21_
                    nw231_[int(6)] = d_1445_v21_
                    nw231_[int(7)] = d_1445_v21_
                    nw231_[int(8)] = d_1445_v21_
                    d_1447_v22_ = nw231_
                    index175_ = default__.safeIndex(858, (d_1447_v22_).length(0))
                    (d_1447_v22_)[index175_] = d_1445_v21_
                    index176_ = default__.safeIndex(858, (d_1447_v22_).length(0))
                    rhs164_ = d_1445_v21_
                    rhs165_ = d_1407_v0_
                    rhs166_ = d_1407_v0_
                    rhs167_ = d_1409_v2_
                    rhs168_ = (0) - ((d_1412_v5_)[default__.safeIndex(d_1409_v2_, len(d_1412_v5_))])
                    lhs77_ = d_1447_v22_
                    lhs78_ = default__.safeIndex(858, (d_1447_v22_).length(0))
                    lhs77_[lhs78_] = rhs164_
                    d_1407_v0_ = rhs165_
                    d_1407_v0_ = rhs166_
                    r1 = rhs167_
                    r1 = rhs168_
                    d_1448_v23_: _dafny.Array
                    nw232_ = _dafny.Array(False, 23)
                    d_1448_v23_ = nw232_
                    d_1449_v24_: str
                    d_1449_v24_ = _dafny.CodePoint('i')
                    d_1450_v25_: _dafny.Map
                    d_1450_v25_ = _dafny.Map({d_1407_v0_: d_1449_v24_})
                    d_1451_v26_: _dafny.Map
                    d_1451_v26_ = _dafny.Map({d_1448_v23_: ((d_1450_v25_)[False] if (False) in (d_1450_v25_) else d_1449_v24_)})
                    d_1451_v26_ = ((d_1451_v26_) | (d_1451_v26_)).set(d_1448_v23_, d_1449_v24_)
                    pass
            pass
        if d_1407_v0_:
            d_1452_v27_: _dafny.Seq
            d_1452_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yeuq"))
            d_1453_v28_: str
            d_1453_v28_ = _dafny.CodePoint('w')
            d_1454_v29_: _dafny.Map
            d_1454_v29_ = _dafny.Map({d_1407_v0_: ((d_1452_v27_).set(default__.safeIndex(d_1409_v2_, len(d_1452_v27_)), d_1453_v28_)) + (d_1452_v27_)})
            d_1454_v29_ = (d_1454_v29_).set(not(d_1407_v0_), d_1452_v27_)
            d_1455_v30_: _dafny.Map
            d_1455_v30_ = _dafny.Map({d_1409_v2_: d_1407_v0_})
            if ((0) - (default__.safeDivisionInt(d_1409_v2_, (0) - (len(d_1455_v30_))))) < (d_1409_v2_):
                d_1456_v32_: _dafny.Map
                d_1456_v32_ = _dafny.Map({d_1452_v27_: _dafny.SeqWithoutIsStrInference([d_1453_v28_ for d_1457_i2_ in range(default__.abs(-431))])})
                d_1458_v33_: _dafny.MultiSet
                def iife134_():
                    coll60_ = _dafny.Map()
                    compr_60_: _dafny.Seq
                    for compr_60_ in (d_1456_v32_).keys.Elements:
                        d_1459_v31_: _dafny.Seq = compr_60_
                        if (d_1459_v31_) in (d_1456_v32_):
                            coll60_[d_1459_v31_] = d_1407_v0_
                    return _dafny.Map(coll60_)
                d_1458_v33_ = _dafny.MultiSet([len(iife134_()
                ), d_1409_v2_])
                d_1460_v34_: _dafny.Map
                d_1460_v34_ = _dafny.Map({d_1407_v0_: d_1453_v28_})
                d_1461_v35_: _dafny.Map
                d_1461_v35_ = _dafny.Map({len(d_1460_v34_): d_1409_v2_})
                d_1462_v36_: D0
                d_1462_v36_ = D0_DC1(d_1409_v2_, False, d_1409_v2_, d_1409_v2_)
                d_1463_v37_: _dafny.Map
                d_1463_v37_ = _dafny.Map({d_1461_v35_: (d_1462_v36_).cf5})
                r1 = ((d_1458_v33_)[len(d_1463_v37_)] if (len(d_1463_v37_)) in (d_1458_v33_) else default__.safeModuloInt(d_1409_v2_, d_1409_v2_))
                d_1464_v38_: D9
                d_1464_v38_ = D9_DC24(d_1409_v2_)
                d_1465_v39_: _dafny.Map
                d_1465_v39_ = _dafny.Map({d_1409_v2_: d_1464_v38_})
                d_1465_v39_ = (d_1465_v39_).set(d_1409_v2_, d_1464_v38_)
                d_1409_v2_ = d_1409_v2_
                d_1407_v0_ = not (not((d_1409_v2_) <= (d_1409_v2_))) or (True)
                d_1409_v2_ = 517
            elif True:
                d_1466_v40_: _dafny.Set
                d_1466_v40_ = _dafny.Set({d_1407_v0_, d_1407_v0_})
                d_1467_v41_: _dafny.Map
                d_1467_v41_ = _dafny.Map({d_1409_v2_: len(d_1466_v40_)})
                d_1468_v43_: D4
                def iife135_():
                    coll61_ = _dafny.Set()
                    compr_61_: int
                    for compr_61_ in (default__.fm60(d_1407_v0_, not(d_1407_v0_), d_1407_v0_, (0) - (d_1409_v2_), globalState)).keys.Elements:
                        d_1469_v42_: int = compr_61_
                        if (d_1469_v42_) in (default__.fm60(d_1407_v0_, not(d_1407_v0_), d_1407_v0_, (0) - (d_1409_v2_), globalState)):
                            coll61_ = coll61_.union(_dafny.Set([(d_1469_v42_) - (len(_dafny.SeqWithoutIsStrInference([509])))]))
                    return _dafny.Set(coll61_)
                d_1468_v43_ = D4_DC9((default__.fm57(d_1467_v41_, d_1409_v2_, globalState)) | (((_dafny.Map({d_1407_v0_: d_1409_v2_})).set(d_1407_v0_, 184)).set(default__.fm27(d_1453_v28_, globalState), d_1409_v2_)), len(iife135_()
), d_1407_v0_)
                d_1468_v43_ = D4_DC9(d_1413_v6_, d_1409_v2_, True)
                d_1470_v44_: _dafny.Array
                def lambda60_(d_1471_v0_):
                    def lambda61_(d_1472_i3_):
                        return d_1471_v0_

                    return lambda61_

                init35_ = lambda60_(d_1407_v0_)
                nw233_ = _dafny.Array(None, 1)
                for i0_35_ in range(nw233_.length(0)):
                    nw233_[i0_35_] = init35_(i0_35_)
                d_1470_v44_ = nw233_
                index177_ = default__.safeIndex(754, (d_1470_v44_).length(0))
                (d_1470_v44_)[index177_] = d_1407_v0_
                index178_ = default__.safeIndex(754, (d_1470_v44_).length(0))
                (d_1470_v44_)[index178_] = False
                rhs169_ = (d_1470_v44_)[default__.safeIndex(754, (d_1470_v44_).length(0))]
                rhs170_ = d_1409_v2_
                rhs171_ = d_1409_v2_
                d_1407_v0_ = rhs169_
                r1 = rhs170_
                d_1409_v2_ = rhs171_
                d_1473_v45_: D7
                d_1473_v45_ = D7_DC19(925)
                d_1474_v46_: _dafny.Map
                d_1474_v46_ = _dafny.Map({d_1412_v5_: d_1473_v45_})
                d_1409_v2_ = len((d_1474_v46_) | (d_1474_v46_))
                d_1475_v47_: _dafny.MultiSet
                d_1475_v47_ = _dafny.MultiSet([d_1407_v0_])
                r1 = ((d_1467_v41_)[d_1409_v2_] if (d_1409_v2_) in (d_1467_v41_) else default__.safeModuloInt(d_1409_v2_, ((d_1475_v47_)[True] if (True) in (d_1475_v47_) else len(d_1466_v40_))))
            if d_1407_v0_:
                d_1476_v48_: _dafny.MultiSet
                d_1476_v48_ = _dafny.MultiSet([d_1407_v0_])
                d_1477_v49_: _dafny.MultiSet
                d_1477_v49_ = _dafny.MultiSet([d_1476_v48_])
                d_1407_v0_ = (d_1412_v5_) < ((_dafny.SeqWithoutIsStrInference([d_1409_v2_, d_1409_v2_])).set(default__.safeIndex(len(((d_1452_v27_).set(default__.safeIndex(len(d_1412_v5_), len(d_1452_v27_)), d_1453_v28_)).set(default__.safeIndex(((d_1477_v49_)[d_1476_v48_] if (d_1476_v48_) in (d_1477_v49_) else len(d_1408_v1_)), len((d_1452_v27_).set(default__.safeIndex(len(d_1412_v5_), len(d_1452_v27_)), d_1453_v28_))), d_1453_v28_)), len(_dafny.SeqWithoutIsStrInference([d_1409_v2_, d_1409_v2_]))), d_1409_v2_))
                d_1478_v50_: D6
                d_1478_v50_ = D6_DC17(D6_DC17(D6_DC15(d_1412_v5_)))
                d_1479_v51_: D6
                d_1479_v51_ = D6_DC15(d_1412_v5_)
                d_1480_v52_: D6
                d_1480_v52_ = D6_DC17(d_1479_v51_)
                d_1481_v53_: D6
                d_1481_v53_ = D6_DC17(d_1480_v52_)
                pat_let_tv19_ = d_1479_v51_
                d_1482_v54_: _dafny.Array
                nw234_ = _dafny.Array(None, 25)
                nw234_[int(0)] = d_1478_v50_
                nw234_[int(1)] = d_1478_v50_
                nw234_[int(2)] = d_1478_v50_
                nw234_[int(3)] = d_1478_v50_
                nw234_[int(4)] = d_1478_v50_
                nw234_[int(5)] = d_1478_v50_
                nw234_[int(6)] = d_1478_v50_
                nw234_[int(7)] = d_1478_v50_
                def iife136_(_pat_let37_0):
                    def iife137_(d_1483_dt__update__tmp_h0_):
                        def iife138_(_pat_let38_0):
                            def iife139_(d_1484_dt__update_hcf39_h0_):
                                return D6_DC17(d_1484_dt__update_hcf39_h0_)
                            return iife139_(_pat_let38_0)
                        return iife138_(pat_let_tv19_)
                    return iife137_(_pat_let37_0)
                nw234_[int(8)] = iife136_(d_1478_v50_)
                nw234_[int(9)] = d_1478_v50_
                nw234_[int(10)] = d_1478_v50_
                nw234_[int(11)] = D6_DC17(d_1479_v51_)
                nw234_[int(12)] = d_1478_v50_
                nw234_[int(13)] = d_1478_v50_
                nw234_[int(14)] = D6_DC17((d_1478_v50_).cf39)
                nw234_[int(15)] = default__.fm42(globalState)
                nw234_[int(16)] = d_1478_v50_
                nw234_[int(17)] = d_1478_v50_
                nw234_[int(18)] = d_1478_v50_
                nw234_[int(19)] = d_1478_v50_
                nw234_[int(20)] = D6_DC17(d_1480_v52_)
                nw234_[int(21)] = d_1478_v50_
                nw234_[int(22)] = d_1478_v50_
                nw234_[int(23)] = d_1478_v50_
                nw234_[int(24)] = d_1478_v50_
                d_1482_v54_ = nw234_
                d_1485_v55_: _dafny.Set
                d_1485_v55_ = _dafny.Set({d_1482_v54_})
                d_1486_v56_: _dafny.Seq
                d_1486_v56_ = _dafny.SeqWithoutIsStrInference([d_1485_v55_, d_1485_v55_])
                d_1487_v57_: _dafny.Array
                nw235_ = _dafny.Array(None, 24)
                nw235_[int(0)] = d_1485_v55_
                nw235_[int(1)] = (d_1485_v55_) | (d_1485_v55_)
                nw235_[int(2)] = _dafny.Set({d_1482_v54_, d_1482_v54_})
                nw235_[int(3)] = d_1485_v55_
                nw235_[int(4)] = (d_1485_v55_).intersection(d_1485_v55_)
                nw235_[int(5)] = d_1485_v55_
                nw235_[int(6)] = (d_1486_v56_)[default__.safeIndex(d_1409_v2_, len(d_1486_v56_))]
                nw235_[int(7)] = d_1485_v55_
                nw235_[int(8)] = d_1485_v55_
                nw235_[int(9)] = (d_1485_v55_) | (_dafny.Set({d_1482_v54_}))
                nw235_[int(10)] = d_1485_v55_
                nw235_[int(11)] = d_1485_v55_
                nw235_[int(12)] = d_1485_v55_
                nw235_[int(13)] = d_1485_v55_
                nw235_[int(14)] = d_1485_v55_
                nw235_[int(15)] = d_1485_v55_
                nw235_[int(16)] = d_1485_v55_
                nw235_[int(17)] = (d_1486_v56_)[default__.safeIndex(d_1409_v2_, len(d_1486_v56_))]
                nw235_[int(18)] = d_1485_v55_
                nw235_[int(19)] = d_1485_v55_
                nw235_[int(20)] = _dafny.Set({d_1482_v54_, d_1482_v54_, d_1482_v54_})
                nw235_[int(21)] = d_1485_v55_
                nw235_[int(22)] = d_1485_v55_
                nw235_[int(23)] = d_1485_v55_
                d_1487_v57_ = nw235_
                index179_ = default__.safeIndex(992, (d_1487_v57_).length(0))
                (d_1487_v57_)[index179_] = d_1485_v55_
                index180_ = default__.safeIndex(992, (d_1487_v57_).length(0))
                (d_1487_v57_)[index180_] = d_1485_v55_
                d_1488_v58_: _dafny.MultiSet
                d_1488_v58_ = _dafny.MultiSet([d_1409_v2_, d_1409_v2_, d_1409_v2_])
                d_1409_v2_ = (d_1488_v58_).cardinality
                d_1489_v59_: C4
                nw236_ = C4()
                nw236_.ctor__(len(d_1452_v27_), d_1407_v0_, ((d_1452_v27_) + (d_1452_v27_)).set(default__.safeIndex(d_1409_v2_, len((d_1452_v27_) + (d_1452_v27_))), d_1453_v28_), not((d_1407_v0_) == (d_1407_v0_)), (d_1407_v0_ if d_1407_v0_ else d_1407_v0_), d_1409_v2_, d_1409_v2_, d_1407_v0_)
                d_1489_v59_ = nw236_
                d_1452_v27_ = (((d_1452_v27_).set(default__.safeIndex(len(d_1412_v5_), len(d_1452_v27_)), d_1453_v28_)).set(default__.safeIndex((d_1489_v59_).f23, len((d_1452_v27_).set(default__.safeIndex(len(d_1412_v5_), len(d_1452_v27_)), d_1453_v28_))), d_1453_v28_)) + (d_1452_v27_)
            elif True:
                d_1407_v0_ = d_1407_v0_
                d_1490_v60_: _dafny.Array
                def lambda62_(d_1491_v0_):
                    def lambda63_(d_1492_i4_):
                        return d_1491_v0_

                    return lambda63_

                init36_ = lambda62_(d_1407_v0_)
                nw237_ = _dafny.Array(None, 3)
                for i0_36_ in range(nw237_.length(0)):
                    nw237_[i0_36_] = init36_(i0_36_)
                d_1490_v60_ = nw237_
                d_1493_v61_: _dafny.Map
                d_1493_v61_ = _dafny.Map({False: d_1490_v60_})
                d_1493_v61_ = (d_1493_v61_).set(d_1407_v0_, d_1490_v60_)
                r1 = (0) - (default__.safeModuloInt(d_1409_v2_, default__.fm49(d_1452_v27_, d_1409_v2_, d_1407_v0_, d_1407_v0_, globalState)))
                d_1494_v62_: C3
                nw238_ = C3()
                nw238_.ctor__(d_1453_v28_, len(d_1455_v30_), (d_1407_v0_) and (d_1407_v0_), d_1407_v0_, default__.safeDivisionInt(d_1409_v2_, (0) - ((0) - (((d_1413_v6_)[d_1407_v0_] if (d_1407_v0_) in (d_1413_v6_) else d_1409_v2_)))))
                d_1494_v62_ = nw238_
                d_1495_v63_: C4
                nw239_ = C4()
                nw239_.ctor__(d_1409_v2_, (True if d_1407_v0_ else d_1407_v0_), d_1452_v27_, d_1407_v0_, d_1407_v0_, (d_1409_v2_ if d_1407_v0_ else d_1409_v2_), len(d_1408_v1_), (d_1409_v2_) > (d_1409_v2_))
                d_1495_v63_ = nw239_
            d_1409_v2_ = d_1409_v2_
            d_1409_v2_ = (709) + (default__.safeDivisionInt(len(d_1452_v27_), d_1409_v2_))
        elif True:
            d_1409_v2_ = default__.fm55(d_1409_v2_, d_1408_v1_, d_1409_v2_, globalState)
            d_1409_v2_ = default__.safeModuloInt(d_1409_v2_, (d_1409_v2_) * (d_1409_v2_))
            d_1407_v0_ = d_1407_v0_
            d_1496_v64_: _dafny.Array
            def lambda64_(d_1497_v2_):
                def lambda65_(d_1498_i5_):
                    return (d_1498_i5_) - (d_1497_v2_)

                return lambda65_

            init37_ = lambda64_(d_1409_v2_)
            nw240_ = _dafny.Array(None, 7)
            for i0_37_ in range(nw240_.length(0)):
                nw240_[i0_37_] = init37_(i0_37_)
            d_1496_v64_ = nw240_
            index181_ = default__.safeIndex(13, (d_1496_v64_).length(0))
            (d_1496_v64_)[index181_] = (d_1409_v2_) - (len(_dafny.Set({d_1407_v0_, False})))
            index182_ = default__.safeIndex(13, (d_1496_v64_).length(0))
            (d_1496_v64_)[index182_] = (0) - (d_1409_v2_)
            d_1499_v65_: _dafny.Map
            d_1499_v65_ = _dafny.Map({-806: d_1407_v0_})
            d_1499_v65_ = d_1499_v65_
        if not(not (d_1407_v0_) or (default__.fm27(_dafny.CodePoint('y'), globalState))):
            r1 = d_1409_v2_
            d_1500_v66_: _dafny.Seq
            d_1500_v66_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h')])
            d_1501_v67_: C3
            nw241_ = C3()
            nw241_.ctor__((d_1500_v66_)[default__.safeIndex(d_1409_v2_, len(d_1500_v66_))], d_1409_v2_, d_1407_v0_, d_1407_v0_, d_1409_v2_)
            d_1501_v67_ = nw241_
            d_1502_v68_: _dafny.Array
            nw242_ = _dafny.Array(False, 21)
            d_1502_v68_ = nw242_
            d_1502_v68_ = d_1502_v68_
            d_1503_v69_: D0
            d_1503_v69_ = D0_DC0(False, d_1409_v2_, d_1407_v0_, d_1409_v2_, d_1407_v0_)
            d_1504_v70_: _dafny.Seq
            d_1504_v70_ = _dafny.SeqWithoutIsStrInference([d_1500_v66_])
            pat_let_tv20_ = d_1407_v0_
            pat_let_tv21_ = d_1409_v2_
            d_1505_v71_: _dafny.Array
            nw243_ = _dafny.Array(None, 19)
            nw243_[int(0)] = d_1503_v69_
            nw243_[int(1)] = d_1503_v69_
            nw243_[int(2)] = d_1503_v69_
            nw243_[int(3)] = D0_DC0(d_1407_v0_, d_1409_v2_, True, d_1409_v2_, not(d_1407_v0_))
            nw243_[int(4)] = d_1503_v69_
            nw243_[int(5)] = d_1503_v69_
            nw243_[int(6)] = (d_1503_v69_ if d_1407_v0_ else d_1503_v69_)
            def iife140_(_pat_let39_0):
                def iife141_(d_1506_dt__update__tmp_h1_):
                    def iife142_(_pat_let40_0):
                        def iife143_(d_1507_dt__update_hcf2_h0_):
                            def iife144_(_pat_let41_0):
                                def iife145_(d_1508_dt__update_hcf1_h0_):
                                    return D0_DC0((d_1506_dt__update__tmp_h1_).cf0, d_1508_dt__update_hcf1_h0_, d_1507_dt__update_hcf2_h0_, (d_1506_dt__update__tmp_h1_).cf3, (d_1506_dt__update__tmp_h1_).cf4)
                                return iife145_(_pat_let41_0)
                            return iife144_(pat_let_tv21_)
                        return iife143_(_pat_let40_0)
                    return iife142_(pat_let_tv20_)
                return iife141_(_pat_let39_0)
            nw243_[int(7)] = iife140_(d_1503_v69_)
            nw243_[int(8)] = D0_DC0(True, d_1409_v2_, d_1407_v0_, len(d_1500_v66_), True)
            nw243_[int(9)] = D0_DC0(not(d_1407_v0_), d_1409_v2_, False, d_1409_v2_, d_1407_v0_)
            nw243_[int(10)] = d_1503_v69_
            nw243_[int(11)] = D0_DC0(d_1407_v0_, d_1409_v2_, d_1407_v0_, (0) - (d_1409_v2_), True)
            nw243_[int(12)] = d_1503_v69_
            nw243_[int(13)] = default__.fm61(d_1409_v2_, not(d_1407_v0_), globalState)
            nw243_[int(14)] = d_1503_v69_
            nw243_[int(15)] = D0_DC0(d_1407_v0_, len((d_1504_v70_)[default__.safeIndex(d_1409_v2_, len(d_1504_v70_))]), d_1407_v0_, d_1409_v2_, d_1407_v0_)
            nw243_[int(16)] = d_1503_v69_
            nw243_[int(17)] = d_1503_v69_
            nw243_[int(18)] = d_1503_v69_
            d_1505_v71_ = nw243_
            index183_ = default__.safeIndex(305, (d_1505_v71_).length(0))
            (d_1505_v71_)[index183_] = d_1503_v69_
            d_1509_v72_: _dafny.MultiSet
            d_1509_v72_ = _dafny.MultiSet([d_1409_v2_, ((0) - (d_1409_v2_)) + (d_1409_v2_), (d_1409_v2_) + ((0) - (d_1409_v2_))])
            d_1510_v73_: _dafny.MultiSet
            d_1510_v73_ = _dafny.MultiSet([d_1407_v0_, d_1407_v0_])
            d_1511_v74_: _dafny.Seq
            d_1511_v74_ = _dafny.SeqWithoutIsStrInference([d_1509_v72_, d_1509_v72_])
            d_1512_v75_: _dafny.Map
            d_1512_v75_ = _dafny.Map({len(d_1408_v1_): ((d_1511_v74_)[default__.safeIndex(d_1409_v2_, len(d_1511_v74_))]).cardinality})
            d_1513_v76_: _dafny.Set
            d_1513_v76_ = _dafny.Set({d_1407_v0_, d_1407_v0_})
            d_1514_v77_: _dafny.Seq
            d_1514_v77_ = _dafny.SeqWithoutIsStrInference([d_1512_v75_, _dafny.Map({len(d_1513_v76_): d_1409_v2_}), d_1512_v75_, d_1512_v75_])
            d_1515_v79_: _dafny.Array
            nw244_ = _dafny.Array(None, 16)
            nw244_[int(0)] = (d_1409_v2_) * (len(default__.fm44((d_1510_v73_).cardinality, d_1503_v69_, d_1407_v0_, not(True), globalState)))
            nw244_[int(1)] = len(d_1514_v77_)
            nw244_[int(2)] = (d_1409_v2_) - (830)
            nw244_[int(3)] = d_1409_v2_
            def iife146_():
                coll62_ = _dafny.Map()
                compr_62_: str
                for compr_62_ in (d_1500_v66_).Elements:
                    d_1516_v78_: str = compr_62_
                    if (d_1516_v78_) in (d_1500_v66_):
                        coll62_[d_1516_v78_] = d_1409_v2_
                return _dafny.Map(coll62_)
            nw244_[int(4)] = default__.fm55(len(iife146_()
            ), d_1411_v4_, d_1409_v2_, globalState)
            nw244_[int(5)] = d_1409_v2_
            nw244_[int(6)] = default__.fm26(globalState)
            nw244_[int(7)] = len((d_1500_v66_ if not(False) else d_1500_v66_))
            nw244_[int(8)] = d_1409_v2_
            nw244_[int(9)] = d_1409_v2_
            nw244_[int(10)] = len(d_1411_v4_)
            nw244_[int(11)] = (0) - (d_1409_v2_)
            nw244_[int(12)] = d_1409_v2_
            nw244_[int(13)] = d_1409_v2_
            nw244_[int(14)] = d_1409_v2_
            nw244_[int(15)] = (len(d_1513_v76_)) + ((0) - (d_1409_v2_))
            d_1515_v79_ = nw244_
            index184_ = default__.safeIndex(42, (d_1515_v79_).length(0))
            (d_1515_v79_)[index184_] = d_1409_v2_
            index185_ = default__.safeIndex(305, (d_1505_v71_).length(0))
            index186_ = default__.safeIndex(42, (d_1515_v79_).length(0))
            rhs172_ = d_1407_v0_
            rhs173_ = d_1503_v69_
            rhs174_ = (d_1509_v72_).intersection(d_1509_v72_)
            rhs175_ = d_1407_v0_
            rhs176_ = (d_1409_v2_) * (d_1409_v2_)
            lhs79_ = d_1505_v71_
            lhs80_ = default__.safeIndex(305, (d_1505_v71_).length(0))
            lhs81_ = d_1515_v79_
            lhs82_ = default__.safeIndex(42, (d_1515_v79_).length(0))
            d_1407_v0_ = rhs172_
            lhs79_[lhs80_] = rhs173_
            d_1509_v72_ = rhs174_
            d_1407_v0_ = rhs175_
            lhs81_[lhs82_] = rhs176_
            index187_ = default__.safeIndex(42, (d_1515_v79_).length(0))
            (d_1515_v79_)[index187_] = (d_1515_v79_)[default__.safeIndex(42, (d_1515_v79_).length(0))]
        elif True:
            if d_1407_v0_:
                d_1517_v80_: str
                d_1517_v80_ = _dafny.CodePoint('t')
                d_1518_v81_: _dafny.Map
                d_1518_v81_ = _dafny.Map({True: d_1517_v80_})
                d_1519_v82_: _dafny.Map
                d_1519_v82_ = _dafny.Map({d_1518_v81_: d_1409_v2_})
                d_1407_v0_ = (_dafny.Map({not(d_1407_v0_): _dafny.CodePoint('y')})) not in (d_1519_v82_)
                d_1520_v83_: C1
                nw245_ = C1()
                nw245_.ctor__(d_1409_v2_, d_1407_v0_, d_1407_v0_, d_1409_v2_)
                d_1520_v83_ = nw245_
                d_1407_v0_ = default__.fm27(d_1517_v80_, globalState)
                d_1521_v84_: _dafny.MultiSet
                d_1521_v84_ = _dafny.MultiSet([d_1407_v0_, d_1407_v0_])
                d_1522_v85_: _dafny.MultiSet
                d_1522_v85_ = d_1521_v84_
                d_1522_v85_ = d_1522_v85_
                (d_1520_v83_).m3(d_1409_v2_, globalState)
            elif True:
                r1 = (d_1412_v5_)[default__.safeIndex((d_1409_v2_ if d_1407_v0_ else d_1409_v2_), len(d_1412_v5_))]
                d_1523_v86_: _dafny.Array
                nw246_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_1523_v86_ = nw246_
                d_1524_v87_: _dafny.Array
                def lambda66_(d_1525_i6_):
                    return True

                init38_ = lambda66_
                nw247_ = _dafny.Array(None, 11)
                for i0_38_ in range(nw247_.length(0)):
                    nw247_[i0_38_] = init38_(i0_38_)
                d_1524_v87_ = nw247_
                index188_ = default__.safeIndex(464, (d_1523_v86_).length(0))
                (d_1523_v86_)[index188_] = d_1524_v87_
                index189_ = default__.safeIndex(464, (d_1523_v86_).length(0))
                (d_1523_v86_)[index189_] = d_1524_v87_
                d_1526_v88_: _dafny.Map
                d_1526_v88_ = _dafny.Map({(d_1523_v86_)[default__.safeIndex(464, (d_1523_v86_).length(0))]: d_1407_v0_})
                d_1526_v88_ = (d_1526_v88_).set((d_1523_v86_)[default__.safeIndex(464, (d_1523_v86_).length(0))], default__.fm27(_dafny.CodePoint('t'), globalState))
                d_1409_v2_ = -354
                d_1407_v0_ = d_1407_v0_
            if d_1407_v0_:
                d_1527_v89_: _dafny.Array
                nw248_ = _dafny.Array(False, 5)
                d_1527_v89_ = nw248_
                index190_ = default__.safeIndex(194, (d_1527_v89_).length(0))
                (d_1527_v89_)[index190_] = d_1407_v0_
                d_1528_v90_: _dafny.Array
                nw249_ = _dafny.Array(int(0), 21)
                d_1528_v90_ = nw249_
                index191_ = default__.safeIndex(862, (d_1528_v90_).length(0))
                (d_1528_v90_)[index191_] = d_1409_v2_
                index192_ = default__.safeIndex(194, (d_1527_v89_).length(0))
                index193_ = default__.safeIndex(862, (d_1528_v90_).length(0))
                rhs177_ = d_1407_v0_
                rhs178_ = d_1409_v2_
                rhs179_ = False
                rhs180_ = (d_1412_v5_)[default__.safeIndex(d_1409_v2_, len(d_1412_v5_))]
                lhs83_ = d_1527_v89_
                lhs84_ = default__.safeIndex(194, (d_1527_v89_).length(0))
                lhs85_ = d_1528_v90_
                lhs86_ = default__.safeIndex(862, (d_1528_v90_).length(0))
                lhs83_[lhs84_] = rhs177_
                lhs85_[lhs86_] = rhs178_
                d_1407_v0_ = rhs179_
                r1 = rhs180_
                index194_ = default__.safeIndex(862, (d_1528_v90_).length(0))
                (d_1528_v90_)[index194_] = d_1409_v2_
                d_1409_v2_ = d_1409_v2_
                d_1529_v91_: _dafny.Array
                nw250_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                d_1529_v91_ = nw250_
                d_1530_v92_: _dafny.Seq
                d_1530_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                index195_ = default__.safeIndex(102, (d_1529_v91_).length(0))
                (d_1529_v91_)[index195_] = d_1530_v92_
                index196_ = default__.safeIndex(102, (d_1529_v91_).length(0))
                (d_1529_v91_)[index196_] = default__.fm37(globalState)
                d_1531_v94_: _dafny.MultiSet
                d_1531_v94_ = _dafny.MultiSet([(d_1527_v89_)[default__.safeIndex(194, (d_1527_v89_).length(0))], d_1407_v0_])
                d_1532_v95_: _dafny.Set
                d_1532_v95_ = _dafny.Set({(d_1527_v89_)[default__.safeIndex(194, (d_1527_v89_).length(0))]})
                d_1533_v96_: _dafny.Set
                d_1533_v96_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(d_1531_v94_).cardinality]), d_1412_v5_, d_1412_v5_, (d_1412_v5_).set(default__.safeIndex(716, len(d_1412_v5_)), len(d_1532_v95_))})
                d_1534_v98_: _dafny.Map
                d_1534_v98_ = _dafny.Map({(d_1528_v90_)[default__.safeIndex(862, (d_1528_v90_).length(0))]: len(d_1411_v4_)})
                d_1535_v99_: _dafny.Seq
                d_1535_v99_ = _dafny.SeqWithoutIsStrInference([d_1534_v98_, d_1534_v98_, d_1534_v98_])
                d_1536_v100_: _dafny.Map
                d_1536_v100_ = _dafny.Map({d_1412_v5_: d_1534_v98_})
                d_1537_v101_: str
                d_1537_v101_ = _dafny.CodePoint('n')
                index197_ = default__.safeIndex(102, (d_1529_v91_).length(0))
                def iife147_():
                    coll63_ = _dafny.Map()
                    compr_63_: _dafny.Seq
                    for compr_63_ in (d_1533_v96_).Elements:
                        d_1538_v93_: _dafny.Seq = compr_63_
                        if (d_1538_v93_) in (d_1533_v96_):
                            def iife148_():
                                coll64_ = _dafny.Map()
                                compr_64_: int
                                for compr_64_ in (_dafny.Set({-381, (d_1528_v90_)[default__.safeIndex(862, (d_1528_v90_).length(0))]})).Elements:
                                    d_1539_v97_: int = compr_64_
                                    if (d_1539_v97_) in (_dafny.Set({-381, (d_1528_v90_)[default__.safeIndex(862, (d_1528_v90_).length(0))]})):
                                        coll64_[(d_1539_v97_) * ((d_1412_v5_)[default__.safeIndex((d_1528_v90_)[default__.safeIndex(862, (d_1528_v90_).length(0))], len(d_1412_v5_))])] = 975
                                return _dafny.Map(coll64_)
                            coll63_[d_1538_v93_] = iife148_()

                    return _dafny.Map(coll63_)
                def iife149_():
                    coll65_ = _dafny.Map()
                    compr_65_: _dafny.Seq
                    for compr_65_ in (d_1533_v96_).Elements:
                        d_1540_v93_: _dafny.Seq = compr_65_
                        if (d_1540_v93_) in (d_1533_v96_):
                            def iife150_():
                                coll66_ = _dafny.Map()
                                compr_66_: int
                                for compr_66_ in (_dafny.Set({-381, (d_1528_v90_)[default__.safeIndex(862, (d_1528_v90_).length(0))]})).Elements:
                                    d_1541_v97_: int = compr_66_
                                    if (d_1541_v97_) in (_dafny.Set({-381, (d_1528_v90_)[default__.safeIndex(862, (d_1528_v90_).length(0))]})):
                                        coll66_[(d_1541_v97_) * ((d_1412_v5_)[default__.safeIndex((d_1528_v90_)[default__.safeIndex(862, (d_1528_v90_).length(0))], len(d_1412_v5_))])] = 975
                                return _dafny.Map(coll66_)
                            coll65_[d_1540_v93_] = iife150_()

                    return _dafny.Map(coll65_)
                (d_1529_v91_)[index197_] = (((d_1529_v91_)[default__.safeIndex(102, (d_1529_v91_).length(0))]).set(default__.safeIndex(len(((iife147_()
                ).set(d_1412_v5_, (d_1535_v99_)[default__.safeIndex(len((d_1529_v91_)[default__.safeIndex(102, (d_1529_v91_).length(0))]), len(d_1535_v99_))])) | (d_1536_v100_)), len((d_1529_v91_)[default__.safeIndex(102, (d_1529_v91_).length(0))])), d_1537_v101_)).set(default__.safeIndex((d_1528_v90_)[default__.safeIndex(862, (d_1528_v90_).length(0))], len(((d_1529_v91_)[default__.safeIndex(102, (d_1529_v91_).length(0))]).set(default__.safeIndex(len(((iife149_()
                ).set(d_1412_v5_, (d_1535_v99_)[default__.safeIndex(len((d_1529_v91_)[default__.safeIndex(102, (d_1529_v91_).length(0))]), len(d_1535_v99_))])) | (d_1536_v100_)), len((d_1529_v91_)[default__.safeIndex(102, (d_1529_v91_).length(0))])), d_1537_v101_))), d_1537_v101_)
            elif True:
                d_1542_v102_: _dafny.Seq
                d_1542_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucn"))
                d_1543_v103_: str
                d_1543_v103_ = _dafny.CodePoint('l')
                d_1544_v104_: _dafny.Array
                nw251_ = _dafny.Array(None, 15)
                nw251_[int(0)] = d_1407_v0_
                nw251_[int(1)] = d_1407_v0_
                nw251_[int(2)] = False
                nw251_[int(3)] = d_1407_v0_
                nw251_[int(4)] = d_1407_v0_
                nw251_[int(5)] = default__.fm27(d_1543_v103_, globalState)
                nw251_[int(6)] = d_1407_v0_
                nw251_[int(7)] = d_1407_v0_
                nw251_[int(8)] = False
                nw251_[int(9)] = d_1407_v0_
                nw251_[int(10)] = d_1407_v0_
                nw251_[int(11)] = True
                nw251_[int(12)] = d_1407_v0_
                nw251_[int(13)] = True
                nw251_[int(14)] = d_1407_v0_
                d_1544_v104_ = nw251_
                d_1545_v105_: _dafny.Map
                d_1545_v105_ = _dafny.Map({d_1542_v102_: d_1544_v104_})
                d_1546_v106_: _dafny.Map
                d_1546_v106_ = _dafny.Map({d_1545_v105_: (d_1542_v102_).set(default__.safeIndex(d_1409_v2_, len(d_1542_v102_)), d_1543_v103_)})
                d_1546_v106_ = (d_1546_v106_).set(d_1545_v105_, (d_1542_v102_) + (d_1542_v102_))
                d_1547_v107_: D5
                d_1547_v107_ = D5_DC13(d_1407_v0_, not(d_1407_v0_))
                d_1548_v108_: _dafny.Map
                d_1548_v108_ = _dafny.Map({(d_1412_v5_) != (_dafny.SeqWithoutIsStrInference([d_1409_v2_ for d_1549_i7_ in range(default__.abs(53))])): (d_1547_v107_).cf30})
                d_1548_v108_ = (d_1548_v108_).set(d_1407_v0_, d_1407_v0_)
                d_1550_v109_: D4
                d_1550_v109_ = D4_DC9(d_1413_v6_, len(d_1542_v102_), d_1407_v0_)
                d_1550_v109_ = d_1550_v109_
                d_1551_v110_: _dafny.Map
                d_1551_v110_ = _dafny.Map({d_1409_v2_: (d_1407_v0_) or (d_1407_v0_)})
                d_1551_v110_ = (d_1551_v110_).set(d_1409_v2_, d_1407_v0_)
                d_1552_v111_: _dafny.Set
                d_1552_v111_ = _dafny.Set({_dafny.Map({d_1407_v0_: d_1409_v2_}), d_1413_v6_})
                rhs181_ = (d_1409_v2_) + ((0) - (d_1409_v2_))
                rhs182_ = d_1552_v111_
                r1 = rhs181_
                d_1552_v111_ = rhs182_
            d_1553_v112_: _dafny.Set
            d_1553_v112_ = _dafny.Set({d_1409_v2_, d_1409_v2_})
            d_1554_v113_: _dafny.Seq
            d_1554_v113_ = _dafny.SeqWithoutIsStrInference([d_1553_v112_])
            rhs183_ = default__.safeModuloInt(684, (0) - (len((d_1554_v113_) + (d_1554_v113_))))
            rhs184_ = d_1407_v0_
            r1 = rhs183_
            d_1407_v0_ = rhs184_
            d_1555_v114_: D0
            d_1555_v114_ = D0_DC0(not(d_1407_v0_), d_1409_v2_, d_1407_v0_, d_1409_v2_, d_1407_v0_)
            r1 = default__.fm32(d_1555_v114_, globalState)
            source17_ = D14_DC38(not (d_1407_v0_) or (d_1407_v0_))
            if source17_.is_DC37:
                d_1556___mcc_h8_ = source17_.cf73
                d_1557_cf73_ = d_1556___mcc_h8_
                d_1558_v115_: _dafny.Map
                d_1558_v115_ = _dafny.Map({d_1409_v2_: (_dafny.MultiSet([(0) - (d_1409_v2_)])).cardinality})
                d_1559_v116_: _dafny.Seq
                d_1559_v116_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lavgxosw"))
                d_1560_v117_: _dafny.MultiSet
                d_1560_v117_ = _dafny.MultiSet([d_1409_v2_])
                d_1561_v118_: _dafny.Array
                nw252_ = _dafny.Array(None, 15)
                nw252_[int(0)] = (d_1407_v0_) or (d_1407_v0_)
                nw252_[int(1)] = not(d_1407_v0_)
                nw252_[int(2)] = d_1407_v0_
                nw252_[int(3)] = not (d_1407_v0_) or (d_1407_v0_)
                nw252_[int(4)] = not(d_1407_v0_)
                nw252_[int(5)] = d_1407_v0_
                nw252_[int(6)] = d_1407_v0_
                nw252_[int(7)] = (d_1407_v0_) and (d_1407_v0_)
                nw252_[int(8)] = not (d_1407_v0_) or (d_1407_v0_)
                nw252_[int(9)] = not((d_1559_v116_) < (d_1559_v116_))
                nw252_[int(10)] = d_1407_v0_
                nw252_[int(11)] = d_1407_v0_
                nw252_[int(12)] = (d_1409_v2_) not in (d_1560_v117_)
                nw252_[int(13)] = d_1407_v0_
                nw252_[int(14)] = d_1407_v0_
                d_1561_v118_ = nw252_
                index198_ = default__.safeIndex(0, (d_1561_v118_).length(0))
                (d_1561_v118_)[index198_] = (d_1407_v0_ if d_1407_v0_ else d_1407_v0_)
                index199_ = default__.safeIndex(0, (d_1561_v118_).length(0))
                rhs185_ = (d_1558_v115_) | (d_1558_v115_)
                rhs186_ = not (False) or ((_dafny.SeqWithoutIsStrInference([len(d_1559_v116_)])) < (d_1412_v5_))
                rhs187_ = ((19) * (len(_dafny.Set({d_1407_v0_})))) <= (d_1409_v2_)
                lhs87_ = d_1561_v118_
                lhs88_ = default__.safeIndex(0, (d_1561_v118_).length(0))
                d_1558_v115_ = rhs185_
                lhs87_[lhs88_] = rhs186_
                d_1407_v0_ = rhs187_
                d_1409_v2_ = (((_dafny.MultiSet([len(d_1559_v116_)])) - (_dafny.MultiSet([d_1409_v2_]))) - ((_dafny.MultiSet([d_1409_v2_])) | (d_1560_v117_))).cardinality
                d_1562_v119_: _dafny.Map
                d_1562_v119_ = _dafny.Map({d_1409_v2_: d_1407_v0_})
                r1 = default__.safeModuloInt((0) - (len((d_1562_v119_ if (d_1561_v118_)[default__.safeIndex(0, (d_1561_v118_).length(0))] else _dafny.Map({-989: d_1407_v0_})))), len((_dafny.SeqWithoutIsStrInference([d_1409_v2_ for d_1563_i8_ in range(default__.abs(-474))])).set(default__.safeIndex(len(_dafny.Map({d_1409_v2_: False})), len(_dafny.SeqWithoutIsStrInference([d_1409_v2_ for d_1564_i8_ in range(default__.abs(-474))]))), d_1409_v2_)))
                d_1565_v120_: bool
                out20_: bool
                out20_ = default__.m0(d_1407_v0_, len(_dafny.SeqWithoutIsStrInference([d_1409_v2_ for d_1566_i9_ in range(default__.abs(365))])), globalState)
                d_1565_v120_ = out20_
            elif source17_.is_DC38:
                d_1567___mcc_h9_ = source17_.cf74
                d_1568_cf74_ = d_1567___mcc_h9_
                d_1407_v0_ = d_1568_cf74_
                r1 = d_1409_v2_
                d_1569_v121_: _dafny.Array
                def lambda67_(d_1570_i10_):
                    return _dafny.CodePoint('s')

                init39_ = lambda67_
                nw253_ = _dafny.Array(None, 21)
                for i0_39_ in range(nw253_.length(0)):
                    nw253_[i0_39_] = init39_(i0_39_)
                d_1569_v121_ = nw253_
                d_1571_v122_: str
                d_1571_v122_ = _dafny.CodePoint('h')
                index200_ = default__.safeIndex(529, (d_1569_v121_).length(0))
                (d_1569_v121_)[index200_] = d_1571_v122_
                index201_ = default__.safeIndex(529, (d_1569_v121_).length(0))
                (d_1569_v121_)[index201_] = d_1571_v122_
                d_1407_v0_ = (d_1407_v0_) or (d_1407_v0_)
            elif True:
                d_1572___mcc_h10_ = source17_.cf72
                d_1573_cf72_ = d_1572___mcc_h10_
                d_1409_v2_ = d_1409_v2_
                d_1409_v2_ = d_1409_v2_
                d_1574_v123_: _dafny.MultiSet
                d_1574_v123_ = _dafny.MultiSet([d_1407_v0_, (d_1409_v2_) > (d_1409_v2_)])
                d_1409_v2_ = ((d_1574_v123_)[d_1407_v0_] if (d_1407_v0_) in (d_1574_v123_) else d_1409_v2_)
                d_1575_v124_: _dafny.Array
                nw254_ = _dafny.Array(None, 2)
                nw254_[int(0)] = default__.fm49(default__.fm37(globalState), d_1409_v2_, True, not(d_1407_v0_), globalState)
                nw254_[int(1)] = d_1409_v2_
                d_1575_v124_ = nw254_
                index202_ = default__.safeIndex(399, (d_1575_v124_).length(0))
                (d_1575_v124_)[index202_] = d_1409_v2_
                index203_ = default__.safeIndex(399, (d_1575_v124_).length(0))
                rhs188_ = (default__.fm55(d_1409_v2_, d_1411_v4_, d_1409_v2_, globalState) if d_1407_v0_ else d_1409_v2_)
                rhs189_ = d_1409_v2_
                rhs190_ = d_1407_v0_
                rhs191_ = d_1409_v2_
                rhs192_ = d_1409_v2_
                lhs89_ = d_1575_v124_
                lhs90_ = default__.safeIndex(399, (d_1575_v124_).length(0))
                d_1409_v2_ = rhs188_
                lhs89_[lhs90_] = rhs189_
                d_1407_v0_ = rhs190_
                r1 = rhs191_
                r1 = rhs192_
        d_1576_v125_: _dafny.Map
        d_1576_v125_ = _dafny.Map({_dafny.Map({d_1407_v0_: (_dafny.MultiSet(d_1411_v4_)).cardinality}): d_1409_v2_})
        d_1577_v126_: _dafny.Map
        d_1577_v126_ = _dafny.Map({d_1409_v2_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kikcba"))})
        d_1578_v127_: D6
        d_1578_v127_ = D6_DC16(d_1576_v125_, (0) - (len(d_1577_v126_)))
        r1 = (d_1578_v127_).cf38
        d_1579_v128_: str
        d_1579_v128_ = _dafny.CodePoint('d')
        d_1579_v128_ = d_1579_v128_
        d_1580_v129_: D12
        d_1580_v129_ = D12_DC33(d_1407_v0_, d_1407_v0_, d_1579_v128_, d_1409_v2_)
        rhs193_ = (d_1407_v0_) == ((d_1580_v129_).cf68)
        rhs194_ = d_1409_v2_
        d_1407_v0_ = rhs193_
        d_1409_v2_ = rhs194_
        d_1581_v130_: _dafny.Seq
        d_1581_v130_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkttrep"))
        d_1582_v131_: _dafny.MultiSet
        d_1582_v131_ = _dafny.MultiSet([(default__.fm52(d_1407_v0_, _dafny.Set({-499}), 164, d_1407_v0_, globalState) if d_1407_v0_ else d_1581_v130_)])
        r0 = d_1582_v131_
        r1 = d_1409_v2_
        d_1583_v132_: _dafny.MultiSet
        d_1583_v132_ = _dafny.MultiSet([d_1407_v0_, d_1407_v0_])
        r2 = (d_1583_v132_).intersection(d_1583_v132_)
        return r0, r1, r2


class C6(T0, T1):
    def  __init__(self):
        self._f4: bool = False
        self._f5: int = int(0)
        self._f15: int = int(0)
        self._f16: bool = False
        self.f21: int = int(0)
        self._f22: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    def ctor__(self, f21, f22, f4, f5, f15, f16):
        (self).f21 = f21
        (self)._f22 = f22
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f15 = f15
        (self)._f16 = f16

    def fm0(self, globalState):
        return ((self).f16) == ((self).f4)

    def m3(self, p0, globalState):
        (self).f21 = default__.safeModuloInt((self).f5, self.f21)
        d_1584_v0_: _dafny.Seq
        d_1584_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvtpgu"))
        d_1585_v1_: str
        d_1585_v1_ = _dafny.CodePoint('b')
        d_1584_v0_ = ((d_1584_v0_) + (d_1584_v0_)).set(default__.safeIndex((self).f5, len((d_1584_v0_) + (d_1584_v0_))), d_1585_v1_)
        d_1586_v2_: _dafny.MultiSet
        d_1586_v2_ = _dafny.MultiSet([(self).f22, (self).f4, (self).f4, (self).f4])
        (self).f21 = ((self).f5 if (d_1586_v2_).issubset(d_1586_v2_) else 419)
        d_1587_v3_: _dafny.Array
        def lambda68_(d_1588_i0_):
            return (_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f4: (self).f16}) for d_1589_i1_ in range(default__.abs(539))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f4: (self).f16}) for d_1590_i2_ in range(default__.abs(675))]))

        init40_ = lambda68_
        nw255_ = _dafny.Array(None, 24)
        for i0_40_ in range(nw255_.length(0)):
            nw255_[i0_40_] = init40_(i0_40_)
        d_1587_v3_ = nw255_
        d_1591_v4_: _dafny.Seq
        d_1591_v4_ = _dafny.SeqWithoutIsStrInference([True])
        d_1592_v5_: _dafny.Map
        d_1592_v5_ = _dafny.Map({(self).f16: (d_1591_v4_)[default__.safeIndex(self.f21, len(d_1591_v4_))]})
        d_1593_v6_: _dafny.Seq
        d_1593_v6_ = _dafny.SeqWithoutIsStrInference([d_1592_v5_])
        index204_ = default__.safeIndex(944, (d_1587_v3_).length(0))
        (d_1587_v3_)[index204_] = d_1593_v6_
        index205_ = default__.safeIndex(944, (d_1587_v3_).length(0))
        (d_1587_v3_)[index205_] = d_1593_v6_
        d_1594_v7_: bool
        d_1594_v7_ = True
        d_1595_v8_: _dafny.Array
        nw256_ = _dafny.Array(False, 24)
        d_1595_v8_ = nw256_
        index206_ = default__.safeIndex(58, (d_1595_v8_).length(0))
        (d_1595_v8_)[index206_] = (self).f4
        d_1596_v9_: _dafny.Array
        nw257_ = _dafny.Array(_dafny.Seq({}), 8)
        d_1596_v9_ = nw257_
        d_1597_v10_: _dafny.Seq
        d_1597_v10_ = _dafny.SeqWithoutIsStrInference([default__.fm25((self).f5, 66, (self).f22, globalState), (0) - (p0), (self).f15])
        index207_ = default__.safeIndex(625, (d_1596_v9_).length(0))
        (d_1596_v9_)[index207_] = d_1597_v10_
        d_1598_v11_: _dafny.Set
        d_1598_v11_ = _dafny.Set({(self).f16})
        index208_ = default__.safeIndex(58, (d_1595_v8_).length(0))
        index209_ = default__.safeIndex(625, (d_1596_v9_).length(0))
        rhs195_ = (d_1591_v4_)[default__.safeIndex((d_1597_v10_)[default__.safeIndex(len(d_1592_v5_), len(d_1597_v10_))], len(d_1591_v4_))]
        rhs196_ = p0
        rhs197_ = d_1598_v11_
        rhs198_ = (d_1597_v10_) <= (d_1597_v10_)
        rhs199_ = (D6_DC15(d_1597_v10_)).cf36
        lhs91_ = self
        lhs92_ = globalState
        lhs93_ = d_1595_v8_
        lhs94_ = default__.safeIndex(58, (d_1595_v8_).length(0))
        lhs95_ = d_1596_v9_
        lhs96_ = default__.safeIndex(625, (d_1596_v9_).length(0))
        d_1594_v7_ = rhs195_
        lhs91_.f21 = rhs196_
        lhs92_.f0 = rhs197_
        lhs93_[lhs94_] = rhs198_
        lhs95_[lhs96_] = rhs199_
        hi5_ = p0
        for d_1599_i3_ in range((len(d_1584_v0_)) - ((self).f5), hi5_):
            d_1600_v12_: C5
            nw258_ = C5()
            nw258_.ctor__()
            d_1600_v12_ = nw258_
            d_1601_v13_: _dafny.Array
            nw259_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_1601_v13_ = nw259_
            d_1602_v14_: _dafny.Array
            nw260_ = _dafny.Array(int(0), 22)
            d_1602_v14_ = nw260_
            index210_ = default__.safeIndex(403, (d_1601_v13_).length(0))
            (d_1601_v13_)[index210_] = d_1602_v14_
            index211_ = default__.safeIndex(403, (d_1601_v13_).length(0))
            (d_1601_v13_)[index211_] = d_1602_v14_
            d_1585_v1_ = d_1585_v1_
            d_1603_v15_: _dafny.Set
            d_1603_v15_ = _dafny.Set({217, (self).f5})
            d_1604_v16_: _dafny.Set
            d_1604_v16_ = _dafny.Set({default__.fm52((d_1595_v8_)[default__.safeIndex(58, (d_1595_v8_).length(0))], d_1603_v15_, d_1599_i3_, d_1594_v7_, globalState), d_1584_v0_, d_1584_v0_, d_1584_v0_})
            d_1605_v17_: D5
            d_1605_v17_ = D5_DC12(d_1604_v16_)
            d_1606_v18_: _dafny.Map
            d_1606_v18_ = _dafny.Map({(self).f16: (self).f5})
            pat_let_tv22_ = d_1606_v18_
            pat_let_tv23_ = globalState
            d_1607_v19_: _dafny.Array
            nw261_ = _dafny.Array(None, 13)
            nw261_[int(0)] = D5_DC12(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxs")), _dafny.SeqWithoutIsStrInference([(D12_DC33((self).f4, (self).f16, d_1585_v1_, 897)).cf69 for d_1608_i4_ in range(default__.abs(-64))])}))
            nw261_[int(1)] = d_1605_v17_
            def iife151_(_pat_let42_0):
                def iife152_(d_1609_dt__update__tmp_h0_):
                    def iife153_(_pat_let43_0):
                        def iife154_(d_1610_dt__update_hcf28_h0_):
                            return D5_DC12(d_1610_dt__update_hcf28_h0_)
                        return iife154_(_pat_let43_0)
                    return iife153_(default__.fm62((self).f4, pat_let_tv22_, (self).f5, pat_let_tv23_))
                return iife152_(_pat_let42_0)
            nw261_[int(2)] = iife151_(d_1605_v17_)
            nw261_[int(3)] = D5_DC12(default__.fm62((self).f4, d_1606_v18_, d_1599_i3_, globalState))
            nw261_[int(4)] = d_1605_v17_
            nw261_[int(5)] = d_1605_v17_
            nw261_[int(6)] = d_1605_v17_
            nw261_[int(7)] = d_1605_v17_
            nw261_[int(8)] = d_1605_v17_
            nw261_[int(9)] = d_1605_v17_
            nw261_[int(10)] = d_1605_v17_
            nw261_[int(11)] = default__.fm63(p0, (self).f22, globalState)
            nw261_[int(12)] = d_1605_v17_
            d_1607_v19_ = nw261_
            index212_ = default__.safeIndex(277, (d_1607_v19_).length(0))
            (d_1607_v19_)[index212_] = d_1605_v17_
            index213_ = default__.safeIndex(277, (d_1607_v19_).length(0))
            (d_1607_v19_)[index213_] = (d_1605_v17_ if True else d_1605_v17_)

    def m11(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        r3: int = int(0)
        r0 = (self).f4
        if (self).f16:
            r0 = (self).f22
            r2 = default__.safeModuloInt(len(_dafny.Set({(self).f4})), (self).f5)
            if (self).f22:
                d_1611_v0_: _dafny.Seq
                d_1611_v0_ = _dafny.SeqWithoutIsStrInference([not(not(not((self).f22))), (self).f4])
                d_1612_v1_: _dafny.Set
                d_1612_v1_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpiyni"))})
                d_1613_v2_: str
                d_1613_v2_ = _dafny.CodePoint('f')
                d_1614_v3_: _dafny.Map
                d_1614_v3_ = _dafny.Map({(self).f22: d_1611_v0_})
                d_1615_v4_: D1
                d_1615_v4_ = D1_DC3(d_1611_v0_)
                d_1616_v5_: _dafny.Array
                nw262_ = _dafny.Array(None, 21)
                nw262_[int(0)] = d_1611_v0_
                nw262_[int(1)] = d_1611_v0_
                nw262_[int(2)] = (d_1611_v0_) + (_dafny.SeqWithoutIsStrInference([(self).f4]))
                nw262_[int(3)] = d_1611_v0_
                nw262_[int(4)] = default__.fm34(self.f21, len(d_1612_v1_), default__.fm27(d_1613_v2_, globalState), (self).f5, globalState)
                nw262_[int(5)] = d_1611_v0_
                nw262_[int(6)] = _dafny.SeqWithoutIsStrInference([(self).f16])
                nw262_[int(7)] = ((_dafny.SeqWithoutIsStrInference([(self).f4])).set(default__.safeIndex((self).f5, len(_dafny.SeqWithoutIsStrInference([(self).f4]))), (self).f16)) + (((d_1614_v3_)[True] if (True) in (d_1614_v3_) else _dafny.SeqWithoutIsStrInference([(self).f22, (self).f22])))
                nw262_[int(8)] = (_dafny.SeqWithoutIsStrInference([(self).f4])) + (d_1611_v0_)
                nw262_[int(9)] = d_1611_v0_
                nw262_[int(10)] = (d_1611_v0_) + ((d_1615_v4_).cf11)
                nw262_[int(11)] = (d_1611_v0_).set(default__.safeIndex(self.f21, len(d_1611_v0_)), not((self).f4))
                nw262_[int(12)] = _dafny.SeqWithoutIsStrInference([(self).fm0(globalState)])
                nw262_[int(13)] = (d_1611_v0_) + (d_1611_v0_)
                nw262_[int(14)] = (d_1611_v0_) + (d_1611_v0_)
                nw262_[int(15)] = d_1611_v0_
                nw262_[int(16)] = (d_1611_v0_) + (d_1611_v0_)
                nw262_[int(17)] = (_dafny.SeqWithoutIsStrInference([(self).f16])) + (_dafny.SeqWithoutIsStrInference([(self).f22, True]))
                nw262_[int(18)] = d_1611_v0_
                nw262_[int(19)] = d_1611_v0_
                nw262_[int(20)] = default__.fm34((self).f15, self.f21, False, (self).f5, globalState)
                d_1616_v5_ = nw262_
                index214_ = default__.safeIndex(449, (d_1616_v5_).length(0))
                (d_1616_v5_)[index214_] = (d_1611_v0_) + (d_1611_v0_)
                index215_ = default__.safeIndex(449, (d_1616_v5_).length(0))
                (d_1616_v5_)[index215_] = d_1611_v0_
                r3 = default__.fm26(globalState)
                d_1617_v6_: _dafny.Map
                d_1617_v6_ = _dafny.Map({False: (self).f16})
                d_1618_v7_: _dafny.Map
                d_1618_v7_ = _dafny.Map({d_1617_v6_: (self).f4})
                d_1619_v8_: _dafny.Seq
                d_1619_v8_ = _dafny.SeqWithoutIsStrInference([self.f21])
                d_1620_v9_: C3
                nw263_ = C3()
                nw263_.ctor__(d_1613_v2_, len((default__.fm64(self.f21, True, self.f21, (self).f16, globalState) if (self).f4 else d_1618_v7_)), (self).f4, (self).f4, len((d_1619_v8_) + (d_1619_v8_)))
                d_1620_v9_ = nw263_
                d_1621_v10_: _dafny.MultiSet
                d_1621_v10_ = _dafny.MultiSet([True])
                d_1622_v11_: _dafny.Seq
                d_1622_v11_ = _dafny.SeqWithoutIsStrInference([d_1621_v10_, d_1621_v10_, d_1621_v10_, d_1621_v10_, d_1621_v10_])
                r3 = ((self).f5) * ((0) - (len(d_1622_v11_)))
                d_1623_v12_: _dafny.Map
                d_1623_v12_ = _dafny.Map({True: 903})
                d_1624_v13_: D9
                d_1624_v13_ = D9_DC23((self).f22, (0) - ((self).f15), (self).f4)
                d_1625_v14_: _dafny.Map
                d_1625_v14_ = _dafny.Map({d_1623_v12_: d_1624_v13_})
                d_1626_v15_: _dafny.Array
                nw264_ = _dafny.Array(None, 6)
                nw264_[int(0)] = (d_1625_v14_) | (d_1625_v14_)
                nw264_[int(1)] = (d_1625_v14_) | (d_1625_v14_)
                nw264_[int(2)] = d_1625_v14_
                nw264_[int(3)] = d_1625_v14_
                nw264_[int(4)] = (d_1625_v14_) | (d_1625_v14_)
                nw264_[int(5)] = (d_1625_v14_) | (d_1625_v14_)
                d_1626_v15_ = nw264_
                index216_ = default__.safeIndex(946, (d_1626_v15_).length(0))
                (d_1626_v15_)[index216_] = (d_1625_v14_) | (d_1625_v14_)
                index217_ = default__.safeIndex(946, (d_1626_v15_).length(0))
                (d_1626_v15_)[index217_] = d_1625_v14_
            elif True:
                d_1627_v16_: _dafny.Array
                nw265_ = _dafny.Array(None, 2)
                nw265_[int(0)] = (self).f4
                nw265_[int(1)] = (self).f4
                d_1627_v16_ = nw265_
                d_1628_v17_: _dafny.MultiSet
                d_1628_v17_ = _dafny.MultiSet([(self).f5])
                d_1629_v18_: _dafny.Seq
                d_1629_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "glc"))
                d_1630_v19_: C2
                nw266_ = C2()
                nw266_.ctor__(d_1627_v16_, (d_1628_v17_) - (d_1628_v17_), d_1629_v18_, (not((self).f4)) == ((self).f22), (self).f16, (self.f21 if not((self).f4) else (self).f5))
                d_1630_v19_ = nw266_
                r1 = ((self).f15 if not((self).f22) else (self).f15)
                r2 = (0) - ((0) - (-763))
                r0 = (self).f4
                d_1631_v20_: _dafny.Seq
                d_1631_v20_ = _dafny.SeqWithoutIsStrInference([self.f21, self.f21])
                d_1632_v21_: _dafny.Seq
                d_1632_v21_ = _dafny.SeqWithoutIsStrInference([d_1631_v20_])
                d_1633_v22_: D1
                d_1633_v22_ = D1_DC5((self).f4, (self).f15, d_1629_v18_)
                d_1634_v23_: _dafny.Map
                d_1634_v23_ = _dafny.Map({d_1633_v22_: d_1632_v21_})
                d_1635_v24_: str
                d_1635_v24_ = _dafny.CodePoint('h')
                rhs200_ = ((d_1634_v23_)[d_1633_v22_] if (d_1633_v22_) in (d_1634_v23_) else d_1632_v21_)
                rhs201_ = d_1627_v16_
                rhs202_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1636_i0_ in range(default__.abs(293))])) + (((d_1629_v18_) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1637_i1_ in range(default__.abs(-101))])).set(default__.safeIndex(451, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1638_i1_ in range(default__.abs(-101))]))), d_1635_v24_))).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojokswewn"))), len((d_1629_v18_) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1639_i1_ in range(default__.abs(-101))])).set(default__.safeIndex(451, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1640_i1_ in range(default__.abs(-101))]))), d_1635_v24_)))), d_1635_v24_))
                rhs203_ = d_1629_v18_
                rhs204_ = 54
                d_1632_v21_ = rhs200_
                d_1627_v16_ = rhs201_
                d_1629_v18_ = rhs202_
                d_1629_v18_ = rhs203_
                r1 = rhs204_
            d_1641_v25_: _dafny.Array
            nw267_ = _dafny.Array(D7.default()(), 1)
            d_1641_v25_ = nw267_
            d_1642_v26_: _dafny.Seq
            d_1642_v26_ = _dafny.SeqWithoutIsStrInference([d_1641_v25_, d_1641_v25_])
            d_1643_v27_: _dafny.Array
            nw268_ = _dafny.Array(None, 12)
            nw268_[int(0)] = d_1641_v25_
            nw268_[int(1)] = d_1641_v25_
            nw268_[int(2)] = d_1641_v25_
            nw268_[int(3)] = d_1641_v25_
            nw268_[int(4)] = d_1641_v25_
            nw268_[int(5)] = d_1641_v25_
            nw268_[int(6)] = (d_1641_v25_ if (self).f16 else d_1641_v25_)
            nw268_[int(7)] = d_1641_v25_
            nw268_[int(8)] = d_1641_v25_
            nw268_[int(9)] = d_1641_v25_
            nw268_[int(10)] = (d_1642_v26_)[default__.safeIndex((self).f15, len(d_1642_v26_))]
            nw268_[int(11)] = d_1641_v25_
            d_1643_v27_ = nw268_
            index218_ = default__.safeIndex(165, (d_1643_v27_).length(0))
            (d_1643_v27_)[index218_] = d_1641_v25_
            d_1644_v28_: _dafny.Array
            nw269_ = _dafny.Array(None, 29)
            nw269_[int(0)] = not(True)
            nw269_[int(1)] = (self).f16
            nw269_[int(2)] = (self).f22
            nw269_[int(3)] = (self).f16
            nw269_[int(4)] = (self).f16
            nw269_[int(5)] = (self).f22
            nw269_[int(6)] = (self).f4
            nw269_[int(7)] = (self).f16
            nw269_[int(8)] = (self).f22
            nw269_[int(9)] = (self).f16
            nw269_[int(10)] = (self).f4
            nw269_[int(11)] = (self).f22
            nw269_[int(12)] = (self).f4
            nw269_[int(13)] = (self).f16
            nw269_[int(14)] = (self).f4
            nw269_[int(15)] = not((self).f4)
            nw269_[int(16)] = (self).f22
            nw269_[int(17)] = (self).f22
            nw269_[int(18)] = (self).f22
            nw269_[int(19)] = not((self).f4)
            nw269_[int(20)] = (self).f4
            nw269_[int(21)] = (self).f16
            nw269_[int(22)] = (self).f22
            nw269_[int(23)] = False
            nw269_[int(24)] = (self).f22
            nw269_[int(25)] = (self).f4
            nw269_[int(26)] = (self).f22
            nw269_[int(27)] = (self).f4
            nw269_[int(28)] = (self).f16
            d_1644_v28_ = nw269_
            d_1645_v29_: _dafny.MultiSet
            d_1645_v29_ = _dafny.MultiSet([d_1644_v28_, d_1644_v28_])
            d_1646_v30_: D7
            d_1646_v30_ = D7_DC18(d_1645_v29_)
            index219_ = default__.safeIndex(165, (d_1643_v27_).length(0))
            nw270_ = _dafny.Array(None, 2)
            nw270_[int(0)] = d_1646_v30_
            nw270_[int(1)] = d_1646_v30_
            (d_1643_v27_)[index219_] = nw270_
            d_1647_v31_: _dafny.MultiSet
            d_1647_v31_ = _dafny.MultiSet([(self).f22, (self).f4])
            r0 = (d_1647_v31_).ispropersubset(d_1647_v31_)
        elif True:
            d_1648_v32_: _dafny.Array
            def lambda69_(d_1649_i2_):
                return (self).f16

            init41_ = lambda69_
            nw271_ = _dafny.Array(None, 19)
            for i0_41_ in range(nw271_.length(0)):
                nw271_[i0_41_] = init41_(i0_41_)
            d_1648_v32_ = nw271_
            d_1650_v33_: _dafny.Map
            d_1650_v33_ = _dafny.Map({(self).f15: (self).f4})
            d_1651_v34_: _dafny.Seq
            d_1651_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pevgm"))
            d_1652_v35_: _dafny.Seq
            d_1652_v35_ = _dafny.SeqWithoutIsStrInference([(self).f16, (self).f4])
            d_1653_v36_: C2
            nw272_ = C2()
            nw272_.ctor__(d_1648_v32_, _dafny.MultiSet([(self).f5, (0) - (self.f21), len(d_1650_v33_)]), d_1651_v34_, (self.f21) <= ((self).f5), (self).f4, (0) - (default__.safeDivisionInt((0) - (len(d_1652_v35_)), (self).f15)))
            d_1653_v36_ = nw272_
            d_1654_v37_: _dafny.Set
            d_1654_v37_ = _dafny.Set({(self).f22, (self).f16, (self).f16})
            r0 = (((self).f4) not in (d_1654_v37_) if not((self).f22) else (d_1652_v35_) <= (d_1652_v35_))
            r0 = not(((self).f15) <= (self.f21))
            d_1655_v38_: _dafny.Map
            d_1655_v38_ = _dafny.Map({not ((self).f4) or (not((self).f22)): d_1648_v32_})
            d_1655_v38_ = (d_1655_v38_).set((self).f16, (d_1653_v36_).f28)
            d_1656_v39_: _dafny.Map
            d_1656_v39_ = _dafny.Map({(self).f22: (self).f16})
            d_1657_v40_: str
            d_1657_v40_ = _dafny.CodePoint('n')
            d_1658_v41_: _dafny.Seq
            d_1658_v41_ = _dafny.SeqWithoutIsStrInference([self.f21, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqdw"))).set(default__.safeIndex(len(d_1652_v35_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqdw")))), d_1657_v40_))])
            d_1659_v42_: D11
            d_1659_v42_ = D11_DC29(self.f21, (d_1653_v36_).f28, (self).f22, ((d_1656_v39_)[True] if (True) in (d_1656_v39_) else (self).f4))
            d_1660_v44_: _dafny.Array
            nw273_ = _dafny.Array(None, 24)
            nw273_[int(0)] = self.f21
            nw273_[int(1)] = (self).f5
            nw273_[int(2)] = len(d_1656_v39_)
            nw273_[int(3)] = (0) - ((self).f5)
            nw273_[int(4)] = self.f21
            nw273_[int(5)] = len((_dafny.Map({(self).f15: self.f21})).set(len(d_1651_v34_), (0) - (self.f21)))
            nw273_[int(6)] = 420
            nw273_[int(7)] = self.f21
            nw273_[int(8)] = (self).f5
            nw273_[int(9)] = 990
            nw273_[int(10)] = (self).f15
            nw273_[int(11)] = len(d_1658_v41_)
            nw273_[int(12)] = self.f21
            nw273_[int(13)] = len(default__.fm46(globalState))
            nw273_[int(14)] = default__.fm49(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uw")), self.f21, (d_1659_v42_).cf60, (self).f4, globalState)
            def iife155_():
                coll67_ = _dafny.Set()
                compr_67_: int
                for compr_67_ in _dafny.IntegerRange(266, 501):
                    d_1661_v43_: int = compr_67_
                    if ((266) <= (d_1661_v43_)) and ((d_1661_v43_) < (501)):
                        coll67_ = coll67_.union(_dafny.Set([default__.safeDivisionInt(d_1661_v43_, len(d_1651_v34_))]))
                return _dafny.Set(coll67_)
            nw273_[int(15)] = len(iife155_()
            )
            nw273_[int(16)] = (self).f5
            nw273_[int(17)] = self.f21
            nw273_[int(18)] = (self).f5
            nw273_[int(19)] = (0) - ((self).f5)
            nw273_[int(20)] = len(d_1651_v34_)
            nw273_[int(21)] = 592
            nw273_[int(22)] = -766
            nw273_[int(23)] = self.f21
            d_1660_v44_ = nw273_
            d_1662_v45_: _dafny.Map
            d_1662_v45_ = _dafny.Map({d_1651_v34_: d_1660_v44_})
            d_1662_v45_ = (d_1662_v45_).set(d_1651_v34_, d_1660_v44_)
        d_1663_v46_: _dafny.Set
        d_1663_v46_ = _dafny.Set({(self).f5})
        d_1664_v47_: _dafny.Seq
        d_1664_v47_ = _dafny.SeqWithoutIsStrInference([(self).f4, (self).fm0(globalState), True, (self).f22])
        d_1665_v48_: C0
        nw274_ = C0()
        nw274_.ctor__(default__.fm41(d_1663_v46_, globalState), len((d_1664_v47_) + (d_1664_v47_)))
        d_1665_v48_ = nw274_
        d_1666_v49_: _dafny.Array
        def lambda70_(d_1667_i3_):
            return (self).f22

        init42_ = lambda70_
        nw275_ = _dafny.Array(None, 28)
        for i0_42_ in range(nw275_.length(0)):
            nw275_[i0_42_] = init42_(i0_42_)
        d_1666_v49_ = nw275_
        index220_ = default__.safeIndex(772, (d_1666_v49_).length(0))
        (d_1666_v49_)[index220_] = (len(_dafny.Set({(self).f5, (d_1665_v48_).f27, (self).f5}))) < ((self).f15)
        d_1668_v50_: _dafny.Seq
        d_1668_v50_ = _dafny.SeqWithoutIsStrInference([(self).f5, (0) - ((d_1665_v48_).f27)])
        d_1669_v52_: _dafny.Map
        def iife156_():
            coll68_ = _dafny.Set()
            compr_68_: int
            for compr_68_ in (d_1668_v50_).Elements:
                d_1670_v51_: int = compr_68_
                if (d_1670_v51_) in (d_1668_v50_):
                    coll68_ = coll68_.union(_dafny.Set([(d_1670_v51_) - (-912)]))
            return _dafny.Set(coll68_)
        d_1669_v52_ = _dafny.Map({(d_1663_v46_ if (self).f22 else iife156_()
        ): (self).f22})
        index221_ = default__.safeIndex(772, (d_1666_v49_).length(0))
        (d_1666_v49_)[index221_] = ((d_1669_v52_)[d_1663_v46_] if (d_1663_v46_) in (d_1669_v52_) else (self).f22)
        d_1671_v53_: D9
        d_1671_v53_ = D9_DC24((self).f5)
        source18_ = d_1671_v53_
        if source18_.is_DC22:
            d_1672___mcc_h0_ = source18_.cf44
            d_1673_cf44_ = d_1672___mcc_h0_
            if (self).f22:
                r3 = 175
                d_1674_v54_: _dafny.Map
                d_1674_v54_ = _dafny.Map({d_1673_cf44_: d_1666_v49_})
                d_1675_v55_: _dafny.Seq
                d_1675_v55_ = _dafny.SeqWithoutIsStrInference([d_1666_v49_])
                d_1676_v56_: _dafny.Array
                nw276_ = _dafny.Array(None, 14)
                nw276_[int(0)] = d_1666_v49_
                nw276_[int(1)] = d_1666_v49_
                nw276_[int(2)] = d_1666_v49_
                nw276_[int(3)] = ((d_1674_v54_)[(self).f4] if ((self).f4) in (d_1674_v54_) else d_1666_v49_)
                nw276_[int(4)] = d_1666_v49_
                nw276_[int(5)] = d_1666_v49_
                nw276_[int(6)] = d_1666_v49_
                nw276_[int(7)] = d_1666_v49_
                nw276_[int(8)] = d_1666_v49_
                nw276_[int(9)] = d_1666_v49_
                nw276_[int(10)] = (d_1675_v55_)[default__.safeIndex((self).f5, len(d_1675_v55_))]
                nw276_[int(11)] = d_1666_v49_
                nw276_[int(12)] = d_1666_v49_
                nw276_[int(13)] = d_1666_v49_
                d_1676_v56_ = nw276_
                d_1677_v57_: _dafny.Map
                d_1677_v57_ = _dafny.Map({(0) - (default__.fm55((d_1665_v48_).f27, _dafny.SeqWithoutIsStrInference([not((self).f16), (self).f22, (self).f16, d_1673_cf44_, (self).f16]), (d_1665_v48_).f27, globalState)): d_1676_v56_})
                d_1678_v58_: _dafny.MultiSet
                d_1678_v58_ = _dafny.MultiSet([d_1673_cf44_])
                d_1679_v59_: _dafny.MultiSet
                d_1679_v59_ = _dafny.MultiSet([(self).f15, (d_1678_v58_).cardinality])
                d_1677_v57_ = (d_1677_v57_).set(((d_1679_v59_)[(d_1665_v48_).f27] if ((d_1665_v48_).f27) in (d_1679_v59_) else (d_1665_v48_).f27), d_1676_v56_)
                d_1680_v60_: str
                d_1680_v60_ = _dafny.CodePoint('d')
                d_1681_v61_: _dafny.Map
                d_1681_v61_ = _dafny.Map({(self).f22: False})
                d_1682_v62_: _dafny.Map
                d_1682_v62_ = _dafny.Map({(self).f4: len(d_1681_v61_)})
                d_1683_v63_: C3
                nw277_ = C3()
                nw277_.ctor__(d_1680_v60_, (self.f21) * ((d_1665_v48_).f27), (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))], (self).f4, len((d_1682_v62_) | (d_1682_v62_)))
                d_1683_v63_ = nw277_
                d_1673_cf44_ = (self).f4
                d_1684_v64_: _dafny.Map
                d_1684_v64_ = _dafny.Map({_dafny.Map({len(d_1664_v47_): (d_1683_v63_).f25}): (self).f16})
                d_1685_v65_: _dafny.Array
                nw278_ = _dafny.Array(D11.default()(), 7)
                d_1685_v65_ = nw278_
                d_1686_v66_: _dafny.Map
                d_1686_v66_ = _dafny.Map({(0) - (len(d_1684_v64_)): d_1685_v65_})
                d_1686_v66_ = (d_1686_v66_).set((d_1665_v48_).f27, d_1685_v65_)
            elif True:
                def lambda71_(d_1687_v49_):
                    def lambda72_(d_1688_i4_):
                        return (d_1687_v49_)[default__.safeIndex(772, (d_1687_v49_).length(0))]

                    return lambda72_

                init43_ = lambda71_(d_1666_v49_)
                nw279_ = _dafny.Array(None, 22)
                for i0_43_ in range(nw279_.length(0)):
                    nw279_[i0_43_] = init43_(i0_43_)
                d_1666_v49_ = nw279_
                d_1689_v67_: _dafny.Array
                nw280_ = _dafny.Array(int(0), 2)
                d_1689_v67_ = nw280_
                index222_ = default__.safeIndex(288, (d_1689_v67_).length(0))
                (d_1689_v67_)[index222_] = ((d_1665_v48_).f27) + ((self).f5)
                index223_ = default__.safeIndex(288, (d_1689_v67_).length(0))
                (d_1689_v67_)[index223_] = default__.fm55(len(d_1668_v50_), (d_1664_v47_) + (d_1664_v47_), self.f21, globalState)
                index224_ = default__.safeIndex(288, (d_1689_v67_).length(0))
                (d_1689_v67_)[index224_] = (d_1665_v48_).f27
                d_1690_v68_: _dafny.Seq
                d_1690_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                d_1691_v69_: str
                d_1691_v69_ = _dafny.CodePoint('l')
                d_1692_v70_: _dafny.Map
                d_1692_v70_ = _dafny.Map({(d_1690_v68_).set(default__.safeIndex((self).f15, len(d_1690_v68_)), d_1691_v69_): (_dafny.MultiSet(d_1668_v50_)).cardinality})
                d_1692_v70_ = d_1692_v70_
                index225_ = default__.safeIndex(772, (d_1666_v49_).length(0))
                (d_1666_v49_)[index225_] = (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]
            d_1693_v71_: _dafny.Seq
            d_1693_v71_ = _dafny.SeqWithoutIsStrInference([d_1664_v47_])
            d_1694_v72_: D9
            d_1694_v72_ = D9_DC22(not((d_1664_v47_) in (d_1693_v71_)))
            source19_ = d_1694_v72_
            if source19_.is_DC22:
                d_1695___mcc_h7_ = source19_.cf44
                d_1696_cf44_ = d_1695___mcc_h7_
                d_1697_v73_: _dafny.Map
                d_1697_v73_ = _dafny.Map({494: 55})
                index226_ = default__.safeIndex(772, (d_1666_v49_).length(0))
                (d_1666_v49_)[index226_] = (((self).f15) < (self.f21)) == ((len(d_1697_v73_)) not in (d_1697_v73_))
                d_1698_v74_: _dafny.Map
                d_1698_v74_ = _dafny.Map({_dafny.Map({(self).f5: (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]}): (self).f4})
                d_1699_v75_: _dafny.Map
                d_1699_v75_ = _dafny.Map({(d_1665_v48_).f27: d_1696_cf44_})
                d_1700_v76_: D0
                d_1700_v76_ = D0_DC0((self).f16, (d_1665_v48_).f27, False, (d_1665_v48_).f27, ((d_1698_v74_)[d_1699_v75_] if (d_1699_v75_) in (d_1698_v74_) else d_1696_cf44_))
                d_1701_v77_: _dafny.MultiSet
                d_1701_v77_ = _dafny.MultiSet([(d_1665_v48_).f27])
                d_1702_v78_: _dafny.Map
                d_1702_v78_ = _dafny.Map({(self).f4: (self).f16})
                d_1703_v79_: D1
                d_1703_v79_ = D1_DC5((self).f22, len(d_1668_v50_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyeudyp")))
                d_1704_v80_: str
                d_1704_v80_ = _dafny.CodePoint('h')
                d_1696_cf44_ = ((d_1701_v77_ if ((d_1702_v78_)[(self).f22] if ((self).f22) in (d_1702_v78_) else d_1696_cf44_) else _dafny.MultiSet([len(d_1664_v47_), len(((d_1703_v79_).cf17).set(default__.safeIndex((d_1665_v48_).f27, len((d_1703_v79_).cf17)), d_1704_v80_)), (d_1665_v48_).f27]))).ispropersubset(((_dafny.MultiSet([default__.fm32(d_1700_v76_, globalState)])).set(self.f21, default__.abs(len(d_1663_v46_)))) | (d_1701_v77_))
                r2 = default__.safeDivisionInt(len((d_1697_v73_) | (d_1697_v73_)), ((self).f5) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjcqlkmp")))))
                rhs205_ = self.f21
                rhs206_ = (d_1665_v48_).f27
                r1 = rhs205_
                r1 = rhs206_
            elif source19_.is_DC23:
                d_1705___mcc_h8_ = source19_.cf45
                d_1706___mcc_h9_ = source19_.cf46
                d_1707___mcc_h10_ = source19_.cf47
                d_1708_cf47_ = d_1707___mcc_h10_
                d_1709_cf46_ = d_1706___mcc_h9_
                d_1710_cf45_ = d_1705___mcc_h8_
                d_1711_v81_: _dafny.Map
                d_1711_v81_ = _dafny.Map({(self).f22: (d_1665_v48_).f27})
                d_1712_v82_: _dafny.Array
                nw281_ = _dafny.Array(None, 2)
                nw281_[int(0)] = (self).f5
                nw281_[int(1)] = len(d_1711_v81_)
                d_1712_v82_ = nw281_
                d_1713_v83_: _dafny.Map
                d_1713_v83_ = _dafny.Map({(D11_DC30(d_1673_cf44_, d_1712_v82_, (0) - ((d_1665_v48_).f27), (self).f22)).cf63: (self).f22})
                d_1714_v84_: str
                d_1714_v84_ = _dafny.CodePoint('d')
                d_1715_v85_: _dafny.Seq
                d_1715_v85_ = _dafny.SeqWithoutIsStrInference([d_1714_v84_])
                d_1716_v86_: _dafny.Map
                d_1716_v86_ = _dafny.Map({d_1715_v85_: (self).fm0(globalState)})
                d_1717_v87_: _dafny.MultiSet
                d_1717_v87_ = _dafny.MultiSet([(d_1665_v48_).f27])
                d_1718_v88_: _dafny.Map
                d_1718_v88_ = _dafny.Map({d_1714_v84_: (self).f15})
                d_1719_v89_: D0
                d_1719_v89_ = D0_DC1((d_1665_v48_).f27, d_1710_cf45_, (d_1665_v48_).f27, len(d_1711_v81_))
                d_1720_v91_: _dafny.Array
                nw282_ = _dafny.Array(None, 14)
                nw282_[int(0)] = (d_1668_v50_)[default__.safeIndex((self).f5, len(d_1668_v50_))]
                nw282_[int(1)] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([(self).f4, d_1708_cf47_, d_1710_cf45_, (self).f22, (self).f4])), d_1709_cf46_)
                nw282_[int(2)] = (d_1665_v48_).f27
                nw282_[int(3)] = (d_1665_v48_).f27
                nw282_[int(4)] = (d_1665_v48_).f27
                nw282_[int(5)] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([True])), d_1709_cf46_)
                nw282_[int(6)] = default__.safeModuloInt(len(d_1713_v83_), default__.fm26(globalState))
                nw282_[int(7)] = (self).f15
                nw282_[int(8)] = (self).f15
                nw282_[int(9)] = (default__.fm65((d_1665_v48_).f27, d_1710_cf45_, (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))], d_1716_v86_, globalState)).cardinality
                nw282_[int(10)] = ((d_1717_v87_)[self.f21] if (self.f21) in (d_1717_v87_) else (d_1668_v50_)[default__.safeIndex(len(d_1718_v88_), len(d_1668_v50_))])
                nw282_[int(11)] = default__.safeDivisionInt((0) - ((d_1665_v48_).f27), (d_1719_v89_).cf8)
                def iife157_():
                    coll69_ = _dafny.Set()
                    compr_69_: _dafny.Set
                    for compr_69_ in (_dafny.Set({d_1663_v46_})).Elements:
                        d_1721_v90_: _dafny.Set = compr_69_
                        if (d_1721_v90_) in (_dafny.Set({d_1663_v46_})):
                            coll69_ = coll69_.union(_dafny.Set([d_1721_v90_]))
                    return _dafny.Set(coll69_)
                nw282_[int(12)] = (len(iife157_()
                )) * ((d_1665_v48_).f27)
                nw282_[int(13)] = default__.safeDivisionInt((0) - (len(_dafny.Map({(d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]: (d_1665_v48_).f27}))), (d_1665_v48_).f27)
                d_1720_v91_ = nw282_
                d_1722_v92_: D4
                d_1722_v92_ = D4_DC10((self).f22, (self).fm0(globalState), d_1664_v47_)
                d_1723_v93_: _dafny.Map
                d_1723_v93_ = _dafny.Map({(self).f15: len(_dafny.Map({(self).f22: (self).f16}))})
                d_1724_v94_: _dafny.Map
                d_1724_v94_ = _dafny.Map({d_1722_v92_: len((d_1713_v83_).set(len(d_1723_v93_), d_1710_cf45_))})
                d_1725_v95_: _dafny.Seq
                d_1725_v95_ = _dafny.SeqWithoutIsStrInference([d_1724_v94_])
                d_1726_v97_: _dafny.Seq
                d_1726_v97_ = _dafny.SeqWithoutIsStrInference([d_1722_v92_, d_1722_v92_])
                d_1727_v98_: D0
                d_1727_v98_ = D0_DC0((self).f4, (d_1665_v48_).f27, d_1710_cf45_, len(_dafny.SeqWithoutIsStrInference([False, d_1710_cf45_])), (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))])
                d_1728_v101_: _dafny.Array
                nw283_ = _dafny.Array(None, 28)
                nw283_[int(0)] = d_1724_v94_
                nw283_[int(1)] = d_1724_v94_
                nw283_[int(2)] = _dafny.Map({d_1722_v92_: -692})
                nw283_[int(3)] = d_1724_v94_
                nw283_[int(4)] = (d_1725_v95_)[default__.safeIndex((d_1665_v48_).f27, len(d_1725_v95_))]
                nw283_[int(5)] = d_1724_v94_
                def iife158_():
                    coll70_ = _dafny.Map()
                    compr_70_: D4
                    for compr_70_ in (d_1726_v97_).Elements:
                        d_1729_v96_: D4 = compr_70_
                        if (d_1729_v96_) in (d_1726_v97_):
                            coll70_[d_1729_v96_] = (D1_DC5((self).f4, (self).f5, d_1715_v85_)).cf16
                    return _dafny.Map(coll70_)
                nw283_[int(6)] = iife158_()
                
                nw283_[int(7)] = (d_1724_v94_) | (d_1724_v94_)
                nw283_[int(8)] = d_1724_v94_
                nw283_[int(9)] = d_1724_v94_
                nw283_[int(10)] = (default__.fm66(globalState)) | (_dafny.Map({d_1722_v92_: (d_1665_v48_).f27}))
                nw283_[int(11)] = (d_1724_v94_).set(D4_DC10(False, (self).f16, d_1664_v47_), default__.fm32(d_1727_v98_, globalState))
                nw283_[int(12)] = (d_1724_v94_) | (d_1724_v94_)
                def iife159_():
                    coll71_ = _dafny.Map()
                    compr_71_: D4
                    for compr_71_ in (default__.fm67(globalState)).Elements:
                        d_1730_v99_: D4 = compr_71_
                        if (d_1730_v99_) in (default__.fm67(globalState)):
                            coll71_[d_1730_v99_] = (d_1665_v48_).f27
                    return _dafny.Map(coll71_)
                nw283_[int(13)] = ((d_1724_v94_).set(d_1722_v92_, (self).f5)) | (iife159_()
                )
                nw283_[int(14)] = d_1724_v94_
                nw283_[int(15)] = d_1724_v94_
                nw283_[int(16)] = d_1724_v94_
                nw283_[int(17)] = d_1724_v94_
                nw283_[int(18)] = d_1724_v94_
                nw283_[int(19)] = d_1724_v94_
                nw283_[int(20)] = (d_1724_v94_) | (d_1724_v94_)
                nw283_[int(21)] = d_1724_v94_
                nw283_[int(22)] = d_1724_v94_
                nw283_[int(23)] = d_1724_v94_
                nw283_[int(24)] = _dafny.Map({d_1722_v92_: (self).f5})
                def iife160_():
                    coll72_ = _dafny.Map()
                    compr_72_: D4
                    for compr_72_ in (_dafny.SeqWithoutIsStrInference([d_1722_v92_, d_1722_v92_, d_1722_v92_])).Elements:
                        d_1731_v100_: D4 = compr_72_
                        if (d_1731_v100_) in (_dafny.SeqWithoutIsStrInference([d_1722_v92_, d_1722_v92_, d_1722_v92_])):
                            coll72_[d_1731_v100_] = (D7_DC19(-816)).cf41
                    return _dafny.Map(coll72_)
                nw283_[int(25)] = (iife160_()
                ) | (d_1724_v94_)
                nw283_[int(26)] = _dafny.Map({D4_DC10((d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))], (self).f16, d_1664_v47_): 338})
                nw283_[int(27)] = d_1724_v94_
                d_1728_v101_ = nw283_
                d_1732_v102_: D4
                d_1732_v102_ = D4_DC8((d_1666_v49_ if (self).f4 else d_1666_v49_))
                rhs207_ = (d_1712_v82_ if (self).f22 else d_1712_v82_)
                rhs208_ = (d_1728_v101_ if not (d_1710_cf45_) or (d_1710_cf45_) else d_1728_v101_)
                rhs209_ = (self).f4
                rhs210_ = d_1732_v102_
                rhs211_ = len(((d_1668_v50_) + (d_1668_v50_) if not((self).f4) else d_1668_v50_))
                d_1712_v82_ = rhs207_
                d_1728_v101_ = rhs208_
                d_1710_cf45_ = rhs209_
                d_1732_v102_ = rhs210_
                r1 = rhs211_
                d_1733_v103_: _dafny.Set
                d_1733_v103_ = _dafny.Set({d_1715_v85_})
                d_1734_v104_: _dafny.Map
                d_1734_v104_ = _dafny.Map({d_1720_v91_: not((d_1733_v103_).ispropersubset(d_1733_v103_))})
                d_1734_v104_ = (d_1734_v104_).set(d_1720_v91_, not (d_1710_cf45_) or ((d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]))
                d_1735_v105_: _dafny.Set
                d_1735_v105_ = _dafny.Set({(self).f22})
                d_1736_v106_: _dafny.Map
                d_1736_v106_ = _dafny.Map({d_1735_v105_: (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]})
                index227_ = default__.safeIndex(489, (d_1712_v82_).length(0))
                (d_1712_v82_)[index227_] = len((d_1736_v106_).set(d_1735_v105_, d_1708_cf47_))
                index228_ = default__.safeIndex(781, (d_1720_v91_).length(0))
                (d_1720_v91_)[index228_] = (d_1665_v48_).f27
                d_1737_v107_: _dafny.Map
                d_1737_v107_ = _dafny.Map({(d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]: d_1720_v91_})
                d_1738_v108_: _dafny.Map
                d_1738_v108_ = _dafny.Map({(self).f5: d_1737_v107_})
                index229_ = default__.safeIndex(489, (d_1712_v82_).length(0))
                index230_ = default__.safeIndex(781, (d_1720_v91_).length(0))
                rhs212_ = (d_1668_v50_) + ((d_1668_v50_) + (_dafny.SeqWithoutIsStrInference([(d_1665_v48_).f27 for d_1739_i5_ in range(default__.abs(992))])))
                rhs213_ = ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1709_cf46_ for d_1740_i6_ in range(default__.abs(906))]))).intersection(d_1717_v87_)).issubset((_dafny.MultiSet([(d_1665_v48_).f27])) - (d_1717_v87_))
                rhs214_ = (self).f5
                rhs215_ = len(((d_1738_v108_)[(self).f15] if ((self).f15) in (d_1738_v108_) else d_1737_v107_))
                rhs216_ = (d_1665_v48_).f27
                lhs97_ = d_1712_v82_
                lhs98_ = default__.safeIndex(489, (d_1712_v82_).length(0))
                lhs99_ = d_1720_v91_
                lhs100_ = default__.safeIndex(781, (d_1720_v91_).length(0))
                d_1668_v50_ = rhs212_
                d_1708_cf47_ = rhs213_
                r1 = rhs214_
                lhs97_[lhs98_] = rhs215_
                lhs99_[lhs100_] = rhs216_
                d_1741_v109_: C3
                nw284_ = C3()
                nw284_.ctor__(d_1714_v84_, (d_1720_v91_)[default__.safeIndex(781, (d_1720_v91_).length(0))], False, (self).f22, (d_1720_v91_)[default__.safeIndex(781, (d_1720_v91_).length(0))])
                d_1741_v109_ = nw284_
            elif source19_.is_DC24:
                d_1742___mcc_h11_ = source19_.cf48
                d_1743_cf48_ = d_1742___mcc_h11_
                d_1744_v110_: _dafny.Seq
                d_1744_v110_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                d_1745_v111_: _dafny.Array
                nw285_ = _dafny.Array(None, 17)
                nw285_[int(0)] = d_1744_v110_
                nw285_[int(1)] = d_1744_v110_
                nw285_[int(2)] = ((d_1744_v110_).set(default__.safeIndex((self).f5, len(d_1744_v110_)), _dafny.CodePoint('t'))) + (d_1744_v110_)
                nw285_[int(3)] = d_1744_v110_
                nw285_[int(4)] = d_1744_v110_
                nw285_[int(5)] = d_1744_v110_
                nw285_[int(6)] = d_1744_v110_
                nw285_[int(7)] = (d_1744_v110_ if (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))] else d_1744_v110_)
                nw285_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlnqfiwu"))
                nw285_[int(9)] = d_1744_v110_
                nw285_[int(10)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxbvss"))) + (d_1744_v110_)
                nw285_[int(11)] = d_1744_v110_
                nw285_[int(12)] = d_1744_v110_
                nw285_[int(13)] = d_1744_v110_
                nw285_[int(14)] = d_1744_v110_
                nw285_[int(15)] = d_1744_v110_
                nw285_[int(16)] = d_1744_v110_
                d_1745_v111_ = nw285_
                index231_ = default__.safeIndex(688, (d_1745_v111_).length(0))
                (d_1745_v111_)[index231_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykgekgkd"))
                index232_ = default__.safeIndex(688, (d_1745_v111_).length(0))
                (d_1745_v111_)[index232_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1746_i7_ in range(default__.abs(414))])
                d_1747_v113_: _dafny.Map
                d_1747_v113_ = _dafny.Map({(d_1665_v48_).f27: (d_1665_v48_).f27})
                d_1748_v114_: _dafny.MultiSet
                def iife161_():
                    coll73_ = _dafny.Map()
                    compr_73_: int
                    for compr_73_ in ((d_1747_v113_)).keys.Elements:
                        d_1749_v112_: int = compr_73_
                        if (d_1749_v112_) in ((d_1747_v113_)):
                            coll73_[default__.safeModuloInt(d_1749_v112_, (d_1665_v48_).f27)] = (self).f22
                    return _dafny.Map(coll73_)
                d_1748_v114_ = _dafny.MultiSet([iife161_()
                ])
                d_1750_v115_: _dafny.Map
                d_1750_v115_ = _dafny.Map({((d_1665_v48_).f27) != (861): d_1748_v114_})
                d_1750_v115_ = (d_1750_v115_).set(not((self).f16), d_1748_v114_)
                d_1751_v116_: C5
                nw286_ = C5()
                nw286_.ctor__()
                d_1751_v116_ = nw286_
                d_1752_v117_: D14
                d_1752_v117_ = D14_DC38((380) <= (966))
                d_1752_v117_ = d_1752_v117_
            elif source19_.is_DC21:
                d_1753___mcc_h12_ = source19_.cf43
                d_1754_cf43_ = d_1753___mcc_h12_
                d_1755_v118_: _dafny.Map
                d_1755_v118_ = _dafny.Map({(self).f16: d_1666_v49_})
                d_1755_v118_ = (d_1755_v118_ if (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))] else d_1755_v118_)
                d_1756_v119_: _dafny.MultiSet
                d_1756_v119_ = _dafny.MultiSet([(self).f5, (self).f15, (d_1665_v48_).f27])
                d_1757_v120_: _dafny.Map
                d_1757_v120_ = _dafny.Map({d_1756_v119_: ((d_1664_v47_).set(default__.safeIndex((self).f5, len(d_1664_v47_)), not((d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]))) == (d_1664_v47_)})
                d_1758_v121_: _dafny.Array
                nw287_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_1758_v121_ = nw287_
                d_1759_v122_: _dafny.Array
                nw288_ = _dafny.Array(_dafny.Map({}), 28)
                d_1759_v122_ = nw288_
                index233_ = default__.safeIndex(812, (d_1758_v121_).length(0))
                (d_1758_v121_)[index233_] = d_1759_v122_
                d_1760_v123_: _dafny.Seq
                d_1760_v123_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({False: (self).f4})])
                d_1761_v124_: _dafny.Map
                d_1761_v124_ = _dafny.Map({(self).f15: (d_1760_v123_)[default__.safeIndex(len(_dafny.Map({self.f21: (d_1665_v48_).f27})), len(d_1760_v123_))]})
                d_1762_v125_: D7
                d_1762_v125_ = D7_DC19(len(d_1761_v124_))
                d_1763_v126_: _dafny.Map
                d_1763_v126_ = _dafny.Map({not((self).f4): (self).f22})
                d_1764_v127_: _dafny.MultiSet
                d_1764_v127_ = _dafny.MultiSet([((d_1763_v126_)[(d_1664_v47_)[default__.safeIndex((d_1665_v48_).f27, len(d_1664_v47_))]] if ((d_1664_v47_)[default__.safeIndex((d_1665_v48_).f27, len(d_1664_v47_))]) in (d_1763_v126_) else (self).f16), d_1673_cf44_, d_1673_cf44_, (self).f4])
                index234_ = default__.safeIndex(812, (d_1758_v121_).length(0))
                rhs217_ = d_1757_v120_
                rhs218_ = default__.safeModuloInt(len(d_1664_v47_), (d_1665_v48_).f27)
                rhs219_ = (self.f21 if d_1673_cf44_ else default__.safeDivisionInt((0) - ((self).f5), (d_1764_v127_).cardinality))
                rhs220_ = d_1759_v122_
                rhs221_ = d_1762_v125_
                lhs101_ = d_1758_v121_
                lhs102_ = default__.safeIndex(812, (d_1758_v121_).length(0))
                d_1757_v120_ = rhs217_
                r3 = rhs218_
                r3 = rhs219_
                lhs101_[lhs102_] = rhs220_
                d_1762_v125_ = rhs221_
                d_1765_v128_: C5
                nw289_ = C5()
                nw289_.ctor__()
                d_1765_v128_ = nw289_
                r0 = (len((d_1664_v47_) + (d_1664_v47_))) <= ((self).f5)
            elif True:
                d_1766___mcc_h13_ = source19_.cf49
                d_1767_cf49_ = d_1766___mcc_h13_
                d_1768_v129_: _dafny.Seq
                d_1768_v129_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eibub"))
                d_1769_v130_: _dafny.Map
                d_1769_v130_ = _dafny.Map({d_1666_v49_: len(d_1768_v129_)})
                d_1770_v131_: _dafny.MultiSet
                d_1770_v131_ = _dafny.MultiSet([(self).f15, self.f21])
                index235_ = default__.safeIndex(772, (d_1666_v49_).length(0))
                rhs222_ = (_dafny.Map({d_1666_v49_: (self).f15})) != ((d_1769_v130_) | (d_1769_v130_))
                rhs223_ = (self).f16
                rhs224_ = ((d_1770_v131_) | (d_1770_v131_)).issubset(d_1770_v131_)
                lhs103_ = d_1666_v49_
                lhs104_ = default__.safeIndex(772, (d_1666_v49_).length(0))
                r0 = rhs222_
                lhs103_[lhs104_] = rhs223_
                r0 = rhs224_
                r0 = not(((d_1770_v131_) - (d_1770_v131_)).issubset(d_1770_v131_))
                d_1771_v132_: str
                d_1771_v132_ = _dafny.CodePoint('v')
                d_1771_v132_ = default__.fm50(globalState)
                d_1772_v133_: D4
                d_1772_v133_ = D4_DC9(_dafny.Map({True: self.f21}), 546, (self).f16)
                d_1773_v134_: D4
                d_1773_v134_ = D4_DC11(d_1772_v133_)
                d_1774_v135_: D4
                d_1774_v135_ = D4_DC11(d_1772_v133_)
                d_1775_v136_: _dafny.Map
                d_1775_v136_ = _dafny.Map({(d_1665_v48_).f27: d_1774_v135_})
                d_1776_v137_: _dafny.Seq
                d_1776_v137_ = _dafny.SeqWithoutIsStrInference([d_1775_v136_, d_1775_v136_])
                d_1777_v138_: _dafny.Seq
                d_1777_v138_ = _dafny.SeqWithoutIsStrInference([(d_1776_v137_)[default__.safeIndex((self).f15, len(d_1776_v137_))]])
                d_1778_v139_: D1
                d_1778_v139_ = D1_DC3(_dafny.SeqWithoutIsStrInference([d_1673_cf44_, (self).f16]))
                d_1779_v140_: _dafny.Map
                d_1779_v140_ = _dafny.Map({(d_1776_v137_) + (_dafny.SeqWithoutIsStrInference([d_1775_v136_, d_1775_v136_, d_1775_v136_, d_1775_v136_, d_1775_v136_])): d_1778_v139_})
                d_1779_v140_ = (d_1779_v140_).set((d_1776_v137_) + (_dafny.SeqWithoutIsStrInference([d_1775_v136_])), d_1778_v139_)
            d_1780_v141_: _dafny.Map
            d_1780_v141_ = _dafny.Map({(self).f15: (self).f4})
            d_1780_v141_ = (d_1780_v141_).set((0) - ((d_1665_v48_).f27), False)
            index236_ = default__.safeIndex(772, (d_1666_v49_).length(0))
            (d_1666_v49_)[index236_] = not(not(True))
        elif source18_.is_DC23:
            d_1781___mcc_h1_ = source18_.cf45
            d_1782___mcc_h2_ = source18_.cf46
            d_1783___mcc_h3_ = source18_.cf47
            d_1784_cf47_ = d_1783___mcc_h3_
            d_1785_cf46_ = d_1782___mcc_h2_
            d_1786_cf45_ = d_1781___mcc_h1_
            d_1786_cf45_ = (self).f16
            d_1787_v142_: _dafny.Set
            d_1787_v142_ = _dafny.Set({(self).f22, (self).f16})
            d_1788_v143_: C1
            nw290_ = C1()
            nw290_.ctor__(len(d_1787_v142_), (self).f4, (self).f22, 749)
            d_1788_v143_ = nw290_
            d_1789_v144_: _dafny.MultiSet
            d_1789_v144_ = _dafny.MultiSet([(self).f4])
            r3 = (773 if not((self).f4) else (d_1789_v144_).cardinality)
            d_1790_v145_: _dafny.Map
            d_1790_v145_ = _dafny.Map({d_1786_cf45_: (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]})
            d_1786_cf45_ = ((d_1790_v145_)[d_1786_cf45_] if (d_1786_cf45_) in (d_1790_v145_) else (self).f16)
        elif source18_.is_DC24:
            d_1791___mcc_h4_ = source18_.cf48
            d_1792_cf48_ = d_1791___mcc_h4_
            d_1793_v146_: _dafny.MultiSet
            d_1793_v146_ = _dafny.MultiSet([not((self).f16)])
            d_1794_v147_: _dafny.Seq
            d_1794_v147_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bb"))
            d_1795_v148_: C4
            nw291_ = C4()
            nw291_.ctor__((d_1665_v48_).f27, (_dafny.MultiSet(d_1664_v47_)) != (d_1793_v146_), d_1794_v147_, not((self).f22), (self).f16, d_1792_cf48_, (0) - (self.f21), (self).f4)
            d_1795_v148_ = nw291_
            d_1796_v149_: _dafny.Array
            def lambda73_(d_1797_v148_):
                def lambda74_(d_1798_i8_):
                    return default__.safeModuloInt(d_1798_i8_, (d_1797_v148_).f23)

                return lambda74_

            init44_ = lambda73_(d_1795_v148_)
            nw292_ = _dafny.Array(None, 29)
            for i0_44_ in range(nw292_.length(0)):
                nw292_[i0_44_] = init44_(i0_44_)
            d_1796_v149_ = nw292_
            d_1799_v150_: _dafny.Map
            d_1799_v150_ = _dafny.Map({d_1668_v50_: d_1796_v149_})
            d_1800_v151_: str
            d_1800_v151_ = _dafny.CodePoint('q')
            d_1801_v152_: _dafny.MultiSet
            d_1801_v152_ = _dafny.MultiSet([d_1800_v151_, d_1800_v151_])
            d_1802_v153_: _dafny.Map
            d_1802_v153_ = _dafny.Map({(d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]: (self).f22})
            d_1803_v154_: D11
            d_1803_v154_ = D11_DC30((d_1795_v148_).f24, ((d_1799_v150_)[_dafny.SeqWithoutIsStrInference([(d_1665_v48_).f27, (d_1801_v152_).cardinality])] if (_dafny.SeqWithoutIsStrInference([(d_1665_v48_).f27, (d_1801_v152_).cardinality])) in (d_1799_v150_) else d_1796_v149_), (self).f5, (548) < (len(d_1802_v153_)))
            source20_ = d_1803_v154_
            if source20_.is_DC28:
                d_1804___mcc_h14_ = source20_.cf52
                d_1805___mcc_h15_ = source20_.cf53
                d_1806___mcc_h16_ = source20_.cf54
                d_1807___mcc_h17_ = source20_.cf55
                d_1808___mcc_h18_ = source20_.cf56
                d_1809_cf56_ = d_1808___mcc_h18_
                d_1810_cf55_ = d_1807___mcc_h17_
                d_1811_cf54_ = d_1806___mcc_h16_
                d_1812_cf53_ = d_1805___mcc_h15_
                d_1813_cf52_ = d_1804___mcc_h14_
                d_1814_v155_: _dafny.Set
                d_1814_v155_ = _dafny.Set({d_1666_v49_, d_1666_v49_})
                d_1815_v156_: _dafny.Seq
                d_1815_v156_ = _dafny.SeqWithoutIsStrInference([d_1814_v155_])
                d_1809_cf56_ = (d_1814_v155_).ispropersubset((d_1815_v156_)[default__.safeIndex((d_1795_v148_).f23, len(d_1815_v156_))])
                d_1666_v49_ = d_1666_v49_
                d_1816_v157_: D4
                d_1816_v157_ = D4_DC10(False, (self).f4, d_1664_v47_)
                d_1817_v158_: D4
                d_1817_v158_ = D4_DC11(d_1816_v157_)
                d_1817_v158_ = D4_DC11(d_1816_v157_)
                index237_ = default__.safeIndex(717, (d_1796_v149_).length(0))
                (d_1796_v149_)[index237_] = d_1811_cf54_
                d_1818_v159_: _dafny.Map
                d_1818_v159_ = _dafny.Map({self.f21: len(d_1794_v147_)})
                d_1819_v160_: _dafny.Map
                d_1819_v160_ = _dafny.Map({d_1818_v159_: d_1664_v47_})
                index238_ = default__.safeIndex(772, (d_1666_v49_).length(0))
                index239_ = default__.safeIndex(717, (d_1796_v149_).length(0))
                rhs225_ = (d_1818_v159_) in (d_1819_v160_)
                rhs226_ = (d_1800_v151_ if ((d_1795_v148_).f23) == (d_1812_cf53_) else d_1800_v151_)
                rhs227_ = (0) - ((default__.fm26(globalState)) + (((self).f5 if (self).f4 else (d_1795_v148_).f23)))
                rhs228_ = ((self).f5) + (d_1811_cf54_)
                rhs229_ = d_1664_v47_
                lhs105_ = d_1666_v49_
                lhs106_ = default__.safeIndex(772, (d_1666_v49_).length(0))
                lhs107_ = d_1796_v149_
                lhs108_ = default__.safeIndex(717, (d_1796_v149_).length(0))
                lhs105_[lhs106_] = rhs225_
                d_1800_v151_ = rhs226_
                r2 = rhs227_
                lhs107_[lhs108_] = rhs228_
                d_1664_v47_ = rhs229_
            elif source20_.is_DC29:
                d_1820___mcc_h19_ = source20_.cf57
                d_1821___mcc_h20_ = source20_.cf58
                d_1822___mcc_h21_ = source20_.cf59
                d_1823___mcc_h22_ = source20_.cf60
                d_1824_cf60_ = d_1823___mcc_h22_
                d_1825_cf59_ = d_1822___mcc_h21_
                d_1826_cf58_ = d_1821___mcc_h20_
                d_1827_cf57_ = d_1820___mcc_h19_
                d_1828_v161_: _dafny.Array
                nw293_ = _dafny.Array(None, 16)
                nw293_[int(0)] = d_1796_v149_
                nw293_[int(1)] = d_1796_v149_
                nw293_[int(2)] = d_1796_v149_
                nw293_[int(3)] = d_1796_v149_
                nw293_[int(4)] = d_1796_v149_
                nw293_[int(5)] = d_1796_v149_
                nw293_[int(6)] = d_1796_v149_
                nw293_[int(7)] = d_1796_v149_
                nw293_[int(8)] = d_1796_v149_
                nw293_[int(9)] = d_1796_v149_
                nw293_[int(10)] = d_1796_v149_
                nw293_[int(11)] = d_1796_v149_
                nw293_[int(12)] = (d_1796_v149_ if (d_1664_v47_)[default__.safeIndex((d_1665_v48_).f27, len(d_1664_v47_))] else d_1796_v149_)
                nw293_[int(13)] = d_1796_v149_
                nw293_[int(14)] = d_1796_v149_
                nw293_[int(15)] = d_1796_v149_
                d_1828_v161_ = nw293_
                index240_ = default__.safeIndex(413, (d_1828_v161_).length(0))
                (d_1828_v161_)[index240_] = d_1796_v149_
                index241_ = default__.safeIndex(413, (d_1828_v161_).length(0))
                (d_1828_v161_)[index241_] = d_1796_v149_
                d_1824_cf60_ = ((len(d_1664_v47_)) + (len(d_1668_v50_))) != (d_1792_cf48_)
                d_1827_cf57_ = default__.fm49(d_1794_v147_, (self).f5, ((self).f4) == ((self).f22), (d_1795_v148_).f24, globalState)
                d_1829_v162_: _dafny.Map
                d_1829_v162_ = _dafny.Map({(self).f22: (d_1793_v146_).cardinality})
                d_1829_v162_ = (d_1829_v162_).set((self).f22, self.f21)
            elif source20_.is_DC30:
                d_1830___mcc_h23_ = source20_.cf61
                d_1831___mcc_h24_ = source20_.cf62
                d_1832___mcc_h25_ = source20_.cf63
                d_1833___mcc_h26_ = source20_.cf64
                d_1834_cf64_ = d_1833___mcc_h26_
                d_1835_cf63_ = d_1832___mcc_h25_
                d_1836_cf62_ = d_1831___mcc_h24_
                d_1837_cf61_ = d_1830___mcc_h23_
                d_1838_v163_: _dafny.Array
                def lambda75_(d_1839_v47_):
                    def lambda76_(d_1840_i9_):
                        return d_1839_v47_

                    return lambda76_

                init45_ = lambda75_(d_1664_v47_)
                nw294_ = _dafny.Array(None, 26)
                for i0_45_ in range(nw294_.length(0)):
                    nw294_[i0_45_] = init45_(i0_45_)
                d_1838_v163_ = nw294_
                index242_ = default__.safeIndex(125, (d_1838_v163_).length(0))
                (d_1838_v163_)[index242_] = _dafny.SeqWithoutIsStrInference([(self).f22])
                d_1841_v164_: D13
                d_1841_v164_ = D13_DC35()
                d_1842_v165_: _dafny.Map
                d_1842_v165_ = _dafny.Map({d_1841_v164_: (d_1665_v48_).f27})
                index243_ = default__.safeIndex(125, (d_1838_v163_).length(0))
                index244_ = default__.safeIndex(772, (d_1666_v49_).length(0))
                rhs230_ = d_1664_v47_
                rhs231_ = default__.fm55((d_1795_v148_).f23, d_1664_v47_, ((d_1842_v165_)[d_1841_v164_] if (d_1841_v164_) in (d_1842_v165_) else d_1792_cf48_), globalState)
                rhs232_ = 745
                rhs233_ = not((len(d_1664_v47_)) < ((d_1665_v48_).f27))
                rhs234_ = (self).f22
                lhs109_ = d_1838_v163_
                lhs110_ = default__.safeIndex(125, (d_1838_v163_).length(0))
                lhs111_ = self
                lhs112_ = d_1666_v49_
                lhs113_ = default__.safeIndex(772, (d_1666_v49_).length(0))
                lhs109_[lhs110_] = rhs230_
                r1 = rhs231_
                lhs111_.f21 = rhs232_
                lhs112_[lhs113_] = rhs233_
                r0 = rhs234_
                d_1843_v166_: _dafny.Map
                d_1843_v166_ = _dafny.Map({self.f21: d_1796_v149_})
                d_1844_v167_: D9
                d_1844_v167_ = D9_DC21((d_1843_v166_).set((d_1665_v48_).f27, d_1836_cf62_))
                d_1845_v168_: _dafny.Array
                nw295_ = _dafny.Array(_dafny.MultiSet({}), 13)
                d_1845_v168_ = nw295_
                index245_ = default__.safeIndex(313, (d_1845_v168_).length(0))
                (d_1845_v168_)[index245_] = d_1793_v146_
                d_1846_v169_: D1
                d_1846_v169_ = D1_DC3(d_1664_v47_)
                pat_let_tv24_ = d_1843_v166_
                index246_ = default__.safeIndex(313, (d_1845_v168_).length(0))
                index247_ = default__.safeIndex(772, (d_1666_v49_).length(0))
                def iife162_(_pat_let44_0):
                    def iife163_(d_1847_dt__update__tmp_h0_):
                        def iife164_(_pat_let45_0):
                            def iife165_(d_1848_dt__update_hcf43_h0_):
                                return D9_DC21(d_1848_dt__update_hcf43_h0_)
                            return iife165_(_pat_let45_0)
                        return iife164_(pat_let_tv24_)
                    return iife163_(_pat_let44_0)
                rhs235_ = not(default__.fm27(d_1800_v151_, globalState))
                rhs236_ = iife162_(d_1844_v167_)
                rhs237_ = d_1793_v146_
                rhs238_ = (self).f16
                rhs239_ = d_1846_v169_
                lhs114_ = d_1845_v168_
                lhs115_ = default__.safeIndex(313, (d_1845_v168_).length(0))
                lhs116_ = d_1666_v49_
                lhs117_ = default__.safeIndex(772, (d_1666_v49_).length(0))
                d_1837_cf61_ = rhs235_
                d_1844_v167_ = rhs236_
                lhs114_[lhs115_] = rhs237_
                lhs116_[lhs117_] = rhs238_
                d_1846_v169_ = rhs239_
                (self).f21 = (0) - ((d_1795_v148_).f23)
                (self).f21 = default__.safeDivisionInt((self.f21) * ((self).f15), (0) - (d_1835_cf63_))
            elif source20_.is_DC27:
                d_1849___mcc_h27_ = source20_.cf51
                d_1850_cf51_ = d_1849___mcc_h27_
                d_1851_v170_: _dafny.Map
                d_1851_v170_ = _dafny.Map({(self).f22: d_1800_v151_})
                d_1852_v171_: _dafny.MultiSet
                d_1852_v171_ = _dafny.MultiSet([(d_1795_v148_).f23, (d_1665_v48_).f27, len(d_1851_v170_)])
                d_1853_v172_: _dafny.Map
                d_1853_v172_ = _dafny.Map({(d_1795_v148_).f23: 331})
                d_1854_v173_: T1
                nw296_ = C3()
                nw296_.ctor__(d_1800_v151_, len(d_1850_cf51_), (d_1795_v148_).f24, (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))], (d_1665_v48_).f27)
                d_1854_v173_ = nw296_
                d_1855_v174_: _dafny.Map
                d_1855_v174_ = _dafny.Map({(self).f4: d_1854_v173_})
                d_1856_v175_: _dafny.Map
                d_1856_v175_ = _dafny.Map({len(d_1855_v174_): (d_1795_v148_).f24})
                d_1857_v176_: _dafny.Seq
                d_1857_v176_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_1795_v148_).f23]), d_1852_v171_, d_1852_v171_, d_1852_v171_, _dafny.MultiSet([len(d_1853_v172_), len(d_1856_v175_)])])
                d_1858_v177_: _dafny.Set
                d_1858_v177_ = _dafny.Set({(d_1854_v173_).f16, (self).f22})
                d_1859_v178_: C2
                nw297_ = C2()
                nw297_.ctor__(d_1666_v49_, d_1852_v171_, d_1794_v147_, (d_1852_v171_).issubset(((d_1857_v176_)[default__.safeIndex(len(d_1858_v177_), len(d_1857_v176_))]).set(self.f21, default__.abs(default__.fm49(d_1794_v147_, 976, (d_1795_v148_).f24, (self).f4, globalState)))), (d_1795_v148_).f24, ((0) - ((d_1795_v148_).f23)) * (739))
                d_1859_v178_ = nw297_
                d_1800_v151_ = default__.fm50(globalState)
                d_1860_v179_: _dafny.Map
                d_1860_v179_ = _dafny.Map({(self).f16: (d_1795_v148_).f23})
                index248_ = default__.safeIndex(613, (d_1796_v149_).length(0))
                (d_1796_v149_)[index248_] = ((d_1860_v179_)[(d_1795_v148_).f24] if ((d_1795_v148_).f24) in (d_1860_v179_) else (d_1665_v48_).f27)
                index249_ = default__.safeIndex(613, (d_1796_v149_).length(0))
                (d_1796_v149_)[index249_] = (self.f21) - (len(d_1668_v50_))
                d_1800_v151_ = d_1800_v151_
            elif True:
                d_1861___mcc_h28_ = source20_.cf65
                d_1862_cf65_ = d_1861___mcc_h28_
                d_1863_v180_: C5
                nw298_ = C5()
                nw298_.ctor__()
                d_1863_v180_ = nw298_
                d_1864_v181_: _dafny.Set
                d_1864_v181_ = _dafny.Set({d_1800_v151_, d_1800_v151_})
                r0 = (d_1864_v181_).isdisjoint(d_1864_v181_)
                d_1865_v182_: C2
                nw299_ = C2()
                nw299_.ctor__(d_1666_v49_, default__.fm68(globalState), (d_1794_v147_).set(default__.safeIndex(self.f21, len(d_1794_v147_)), d_1800_v151_), (not(False)) in (_dafny.MultiSet([(self).f4])), False, ((0) - (d_1792_cf48_)) - ((d_1665_v48_).f27))
                d_1865_v182_ = nw299_
                d_1866_v183_: D1
                d_1866_v183_ = D1_DC4(True, (d_1795_v148_).f23, (d_1665_v48_).f27)
                pat_let_tv25_ = d_1665_v48_
                def iife166_(_pat_let46_0):
                    def iife167_(d_1867_dt__update__tmp_h1_):
                        def iife168_(_pat_let47_0):
                            def iife169_(d_1868_dt__update_hcf13_h0_):
                                return D1_DC4((d_1867_dt__update__tmp_h1_).cf12, d_1868_dt__update_hcf13_h0_, (d_1867_dt__update__tmp_h1_).cf14)
                            return iife169_(_pat_let47_0)
                        return iife168_((pat_let_tv25_).f27)
                    return iife167_(_pat_let46_0)
                (d_1863_v180_).m13((d_1665_v48_).f27, iife166_(d_1866_v183_), globalState)
            r0 = (d_1795_v148_).f24
            d_1869_v184_: _dafny.Map
            d_1869_v184_ = _dafny.Map({(self).f5: d_1793_v146_})
            d_1870_v185_: _dafny.Map
            d_1870_v185_ = _dafny.Map({((d_1869_v184_)[(d_1665_v48_).f27] if ((d_1665_v48_).f27) in (d_1869_v184_) else d_1793_v146_): d_1794_v147_})
            d_1870_v185_ = d_1870_v185_
        elif source18_.is_DC21:
            d_1871___mcc_h5_ = source18_.cf43
            d_1872_cf43_ = d_1871___mcc_h5_
            d_1873_v186_: _dafny.Array
            def lambda77_(d_1874_i10_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xgxpsw"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxkbe")))

            init46_ = lambda77_
            nw300_ = _dafny.Array(None, 1)
            for i0_46_ in range(nw300_.length(0)):
                nw300_[i0_46_] = init46_(i0_46_)
            d_1873_v186_ = nw300_
            d_1875_v187_: _dafny.Seq
            d_1875_v187_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlxt"))
            index250_ = default__.safeIndex(613, (d_1873_v186_).length(0))
            (d_1873_v186_)[index250_] = d_1875_v187_
            d_1876_v188_: str
            d_1876_v188_ = _dafny.CodePoint('c')
            index251_ = default__.safeIndex(613, (d_1873_v186_).length(0))
            (d_1873_v186_)[index251_] = (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "quwjym"))) + (d_1875_v187_)).set(default__.safeIndex(820, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "quwjym"))) + (d_1875_v187_))), d_1876_v188_)) + ((d_1875_v187_) + (d_1875_v187_))
            d_1877_v189_: _dafny.Seq
            d_1877_v189_ = _dafny.SeqWithoutIsStrInference([d_1666_v49_, d_1666_v49_, d_1666_v49_])
            d_1878_v190_: _dafny.Map
            d_1878_v190_ = _dafny.Map({(d_1873_v186_)[default__.safeIndex(613, (d_1873_v186_).length(0))]: (d_1877_v189_)[default__.safeIndex((self).f5, len(d_1877_v189_))]})
            d_1878_v190_ = (d_1878_v190_).set(d_1875_v187_, d_1666_v49_)
            d_1879_v191_: _dafny.Array
            def lambda78_(d_1880_i11_):
                return default__.safeModuloInt(d_1880_i11_, (self).f5)

            init47_ = lambda78_
            nw301_ = _dafny.Array(None, 24)
            for i0_47_ in range(nw301_.length(0)):
                nw301_[i0_47_] = init47_(i0_47_)
            d_1879_v191_ = nw301_
            index252_ = default__.safeIndex(643, (d_1879_v191_).length(0))
            (d_1879_v191_)[index252_] = default__.safeModuloInt((d_1665_v48_).f27, self.f21)
            index253_ = default__.safeIndex(643, (d_1879_v191_).length(0))
            (d_1879_v191_)[index253_] = (0) - (self.f21)
            d_1881_v192_: _dafny.Map
            d_1881_v192_ = _dafny.Map({len((d_1873_v186_)[default__.safeIndex(613, (d_1873_v186_).length(0))]): (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]})
            d_1881_v192_ = (d_1881_v192_).set(default__.safeDivisionInt((self).f15, (d_1665_v48_).f27), (self).f16)
        elif True:
            d_1882___mcc_h6_ = source18_.cf49
            d_1883_cf49_ = d_1882___mcc_h6_
            d_1884_v193_: _dafny.Map
            d_1884_v193_ = _dafny.Map({False: (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]})
            source21_ = d_1884_v193_
            d_1885___mcc_h29_ = source21_
            d_1886_cf50_ = d_1885___mcc_h29_
            d_1887_v194_: _dafny.Seq
            d_1887_v194_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phoqoik"))
            d_1888_v195_: _dafny.Map
            d_1888_v195_ = _dafny.Map({(self).f5: (d_1665_v48_).f27})
            d_1889_v196_: _dafny.Map
            d_1889_v196_ = d_1888_v195_
            d_1890_v197_: _dafny.Map
            d_1890_v197_ = _dafny.Map({len(d_1887_v194_): (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]})
            d_1891_v198_: C4
            nw302_ = C4()
            nw302_.ctor__((self).f15, (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))], d_1887_v194_, (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))], (self).f4, len(_dafny.Map({d_1889_v196_: ((d_1890_v197_)[self.f21] if (self.f21) in (d_1890_v197_) else (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))])})), (self).f15, (self).f16)
            d_1891_v198_ = nw302_
            d_1892_v199_: _dafny.MultiSet
            d_1892_v199_ = _dafny.MultiSet([(d_1664_v47_)[default__.safeIndex((d_1665_v48_).f27, len(d_1664_v47_))]])
            d_1893_v200_: _dafny.MultiSet
            d_1893_v200_ = _dafny.MultiSet([(self).f5, ((d_1892_v199_)[True] if (True) in (d_1892_v199_) else self.f21)])
            d_1894_v201_: C2
            nw303_ = C2()
            nw303_.ctor__(d_1666_v49_, d_1893_v200_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwygqg")), (self.f21) > (self.f21), (d_1664_v47_) <= (d_1664_v47_), (d_1665_v48_).f27)
            d_1894_v201_ = nw303_
            d_1895_v202_: D6
            d_1895_v202_ = D6_DC15(d_1668_v50_)
            d_1896_v203_: D6
            d_1896_v203_ = D6_DC17(d_1895_v202_)
            d_1897_v204_: _dafny.Map
            d_1897_v204_ = _dafny.Map({d_1896_v203_: (d_1891_v198_).f24})
            d_1898_v205_: _dafny.Array
            nw304_ = _dafny.Array(None, 10)
            nw304_[int(0)] = default__.fm69(globalState)
            nw304_[int(1)] = d_1897_v204_
            nw304_[int(2)] = d_1897_v204_
            nw304_[int(3)] = (d_1897_v204_) | (d_1897_v204_)
            nw304_[int(4)] = d_1897_v204_
            nw304_[int(5)] = default__.fm69(globalState)
            nw304_[int(6)] = _dafny.Map({d_1896_v203_: (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]})
            nw304_[int(7)] = d_1897_v204_
            nw304_[int(8)] = (d_1897_v204_) | (d_1897_v204_)
            nw304_[int(9)] = d_1897_v204_
            d_1898_v205_ = nw304_
            d_1899_v206_: _dafny.Seq
            d_1899_v206_ = _dafny.SeqWithoutIsStrInference([d_1898_v205_, d_1898_v205_])
            d_1898_v205_ = (d_1899_v206_)[default__.safeIndex((d_1665_v48_).f27, len(d_1899_v206_))]
            d_1900_v207_: _dafny.Array
            nw305_ = _dafny.Array(_dafny.Seq({}), 29)
            d_1900_v207_ = nw305_
            d_1901_v208_: _dafny.Array
            nw306_ = _dafny.Array(int(0), 18)
            d_1901_v208_ = nw306_
            index254_ = default__.safeIndex(570, (d_1900_v207_).length(0))
            (d_1900_v207_)[index254_] = _dafny.SeqWithoutIsStrInference([d_1901_v208_])
            d_1902_v209_: _dafny.Set
            d_1902_v209_ = _dafny.Set({(self).f16})
            d_1903_v210_: _dafny.Map
            d_1903_v210_ = _dafny.Map({(d_1902_v209_).issubset(d_1902_v209_): (self).f5})
            d_1904_v211_: str
            d_1904_v211_ = _dafny.CodePoint('k')
            d_1905_v212_: _dafny.MultiSet
            d_1905_v212_ = _dafny.MultiSet([d_1904_v211_])
            d_1906_v213_: D12
            d_1906_v213_ = D12_DC32(d_1905_v212_)
            d_1907_v214_: _dafny.Seq
            d_1907_v214_ = _dafny.SeqWithoutIsStrInference([d_1892_v199_, d_1892_v199_, d_1892_v199_])
            d_1908_v215_: _dafny.Map
            d_1908_v215_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1909_i12_ in range(default__.abs(195))]): _dafny.Map({(d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]: self.f21})})
            d_1910_v216_: D5
            d_1910_v216_ = D5_DC14((self).f16, d_1664_v47_, d_1901_v208_, self.f21, -265)
            d_1911_v217_: _dafny.MultiSet
            d_1911_v217_ = _dafny.MultiSet([(d_1910_v216_).cf33, (d_1910_v216_).cf33, d_1901_v208_])
            index255_ = default__.safeIndex(570, (d_1900_v207_).length(0))
            rhs240_ = (d_1892_v199_) - ((d_1907_v214_)[default__.safeIndex((d_1891_v198_).f23, len(d_1907_v214_))])
            rhs241_ = 133
            rhs242_ = _dafny.SeqWithoutIsStrInference([d_1901_v208_, d_1901_v208_, d_1901_v208_])
            rhs243_ = ((d_1908_v215_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsxunsv"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsxunsv"))) in (d_1908_v215_) else (_dafny.Map({(self).f4: (d_1891_v198_).f23})).set((d_1891_v198_).f24, (d_1911_v217_).cardinality))
            rhs244_ = default__.fm70(globalState)
            lhs118_ = d_1900_v207_
            lhs119_ = default__.safeIndex(570, (d_1900_v207_).length(0))
            d_1892_v199_ = rhs240_
            r1 = rhs241_
            lhs118_[lhs119_] = rhs242_
            d_1903_v210_ = rhs243_
            d_1906_v213_ = rhs244_
            d_1912_v218_: _dafny.Seq
            d_1912_v218_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwsqicxej"))
            d_1913_v219_: D0
            d_1913_v219_ = D0_DC0((d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))], self.f21, (self).f4, len(d_1912_v218_), (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))])
            d_1914_v220_: _dafny.Array
            nw307_ = _dafny.Array(int(0), 15)
            d_1914_v220_ = nw307_
            d_1915_v221_: D5
            d_1915_v221_ = D5_DC14((d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))], (d_1664_v47_).set(default__.safeIndex(default__.fm32(d_1913_v219_, globalState), len(d_1664_v47_)), (self).f4), d_1914_v220_, (self).f15, (d_1665_v48_).f27)
            index256_ = default__.safeIndex(772, (d_1666_v49_).length(0))
            (d_1666_v49_)[index256_] = (d_1915_v221_).cf31
            r0 = (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))]
            d_1916_v222_: _dafny.MultiSet
            d_1916_v222_ = _dafny.MultiSet([(self).f22, (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))], (d_1666_v49_)[default__.safeIndex(772, (d_1666_v49_).length(0))], (self).f4])
            d_1663_v46_ = _dafny.Set({default__.safeDivisionInt(((d_1916_v222_)[(self).f22] if ((self).f22) in (d_1916_v222_) else (d_1665_v48_).f27), len(_dafny.Set({(self).f15, self.f21, (self).f5, (self).f5, (d_1668_v50_)[default__.safeIndex((self).f15, len(d_1668_v50_))]})))})
        d_1668_v50_ = d_1668_v50_
        r0 = ((self).f5) > (((self).f15) * ((self).f15))
        r1 = default__.safeDivisionInt(504, (0) - ((self).f5))
        r2 = self.f21
        r3 = (d_1665_v48_).f27
        return r0, r1, r2, r3

    def m12(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_1917_v0_: bool
        d_1917_v0_ = True
        d_1918_v1_: _dafny.Array
        nw308_ = _dafny.Array(_dafny.Array(None, 0), 13)
        d_1918_v1_ = nw308_
        d_1919_v2_: _dafny.Array
        def lambda79_(d_1920_v0_):
            def lambda80_(d_1921_i0_):
                return d_1920_v0_

            return lambda80_

        init48_ = lambda79_(d_1917_v0_)
        nw309_ = _dafny.Array(None, 16)
        for i0_48_ in range(nw309_.length(0)):
            nw309_[i0_48_] = init48_(i0_48_)
        d_1919_v2_ = nw309_
        rhs245_ = (self).f16
        rhs246_ = d_1918_v1_
        rhs247_ = d_1919_v2_
        d_1917_v0_ = rhs245_
        d_1918_v1_ = rhs246_
        r0 = rhs247_
        d_1922_i1_: int
        d_1922_i1_ = 0
        with _dafny.label("4"):
            while (self).f4:
                with _dafny.c_label("4"):
                    if (d_1922_i1_) >= (100):
                        raise _dafny.Break("4")
                    d_1922_i1_ = (d_1922_i1_) + (1)
                    d_1923_v3_: _dafny.Seq
                    d_1923_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvbbm"))
                    d_1924_v4_: _dafny.Seq
                    d_1924_v4_ = _dafny.SeqWithoutIsStrInference([d_1923_v3_])
                    index257_ = default__.safeIndex(916, (d_1919_v2_).length(0))
                    (d_1919_v2_)[index257_] = False
                    index258_ = default__.safeIndex(34, (d_1919_v2_).length(0))
                    (d_1919_v2_)[index258_] = (self).f16
                    d_1925_v5_: _dafny.MultiSet
                    d_1925_v5_ = _dafny.MultiSet([648, self.f21])
                    d_1926_v6_: _dafny.Set
                    d_1926_v6_ = _dafny.Set({(d_1925_v5_).cardinality, self.f21})
                    index259_ = default__.safeIndex(916, (d_1919_v2_).length(0))
                    index260_ = default__.safeIndex(34, (d_1919_v2_).length(0))
                    rhs248_ = (((d_1925_v5_) | (_dafny.MultiSet([self.f21]))).cardinality) + (self.f21)
                    rhs249_ = _dafny.SeqWithoutIsStrInference([(default__.fm37(globalState)) + (default__.fm52((self).f22, d_1926_v6_, self.f21, not(d_1917_v0_), globalState)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojj")), d_1923_v3_, (d_1923_v3_) + (d_1923_v3_), (d_1923_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edqujvfl")))])
                    rhs250_ = (self).f4
                    rhs251_ = (self).f16
                    rhs252_ = d_1917_v0_
                    lhs120_ = self
                    lhs121_ = d_1919_v2_
                    lhs122_ = default__.safeIndex(916, (d_1919_v2_).length(0))
                    lhs123_ = d_1919_v2_
                    lhs124_ = default__.safeIndex(34, (d_1919_v2_).length(0))
                    lhs120_.f21 = rhs248_
                    d_1924_v4_ = rhs249_
                    lhs121_[lhs122_] = rhs250_
                    lhs123_[lhs124_] = rhs251_
                    d_1917_v0_ = rhs252_
                    index261_ = default__.safeIndex(916, (d_1919_v2_).length(0))
                    index262_ = default__.safeIndex(916, (d_1919_v2_).length(0))
                    index263_ = default__.safeIndex(916, (d_1919_v2_).length(0))
                    rhs253_ = (d_1919_v2_)[default__.safeIndex(916, (d_1919_v2_).length(0))]
                    rhs254_ = (d_1919_v2_)[default__.safeIndex(916, (d_1919_v2_).length(0))]
                    rhs255_ = (self).f16
                    lhs125_ = d_1919_v2_
                    lhs126_ = default__.safeIndex(916, (d_1919_v2_).length(0))
                    lhs127_ = d_1919_v2_
                    lhs128_ = default__.safeIndex(916, (d_1919_v2_).length(0))
                    lhs129_ = d_1919_v2_
                    lhs130_ = default__.safeIndex(916, (d_1919_v2_).length(0))
                    lhs125_[lhs126_] = rhs253_
                    lhs127_[lhs128_] = rhs254_
                    lhs129_[lhs130_] = rhs255_
                    d_1927_v7_: C5
                    nw310_ = C5()
                    nw310_.ctor__()
                    d_1927_v7_ = nw310_
                    d_1928_v8_: _dafny.Seq
                    d_1928_v8_ = _dafny.SeqWithoutIsStrInference([len(default__.fm58((self).f5, (self).f15, globalState))])
                    (self).f21 = len(d_1928_v8_)
                    pass
            pass
        d_1929_v9_: str
        d_1929_v9_ = _dafny.CodePoint('o')
        d_1930_v10_: _dafny.Map
        d_1930_v10_ = _dafny.Map({d_1929_v9_: (self).f16})
        d_1931_v11_: _dafny.Map
        d_1931_v11_ = _dafny.Map({self.f21: len(d_1930_v10_)})
        d_1932_v12_: _dafny.Map
        d_1932_v12_ = _dafny.Map({(self).f15: (self).f4})
        d_1933_v13_: _dafny.Map
        d_1933_v13_ = _dafny.Map({d_1929_v9_: (self).f5})
        d_1934_v14_: _dafny.Set
        d_1934_v14_ = _dafny.Set({d_1919_v2_, d_1919_v2_})
        d_1935_v15_: _dafny.MultiSet
        d_1935_v15_ = _dafny.MultiSet([False, d_1917_v0_])
        d_1936_v16_: _dafny.Map
        d_1936_v16_ = _dafny.Map({d_1917_v0_: (self).f15})
        d_1937_v17_: _dafny.Set
        d_1937_v17_ = _dafny.Set({len(d_1936_v16_)})
        d_1938_v18_: _dafny.Array
        nw311_ = _dafny.Array(None, 27)
        nw311_[int(0)] = (self).f5
        nw311_[int(1)] = (self).f5
        nw311_[int(2)] = (self).f5
        nw311_[int(3)] = self.f21
        nw311_[int(4)] = self.f21
        nw311_[int(5)] = default__.safeModuloInt(len(d_1931_v11_), len(d_1932_v12_))
        nw311_[int(6)] = (0) - (default__.safeDivisionInt(((d_1933_v13_)[default__.fm50(globalState)] if (default__.fm50(globalState)) in (d_1933_v13_) else len(_dafny.Map({len(_dafny.Map({(self).f4: d_1917_v0_})): d_1917_v0_}))), self.f21))
        nw311_[int(7)] = self.f21
        nw311_[int(8)] = (0) - (self.f21)
        nw311_[int(9)] = (-264) * (len(d_1934_v14_))
        nw311_[int(10)] = self.f21
        nw311_[int(11)] = ((d_1935_v15_)[True] if (True) in (d_1935_v15_) else (self).f5)
        nw311_[int(12)] = 12
        nw311_[int(13)] = (self).f5
        nw311_[int(14)] = self.f21
        nw311_[int(15)] = (self).f15
        nw311_[int(16)] = len(d_1937_v17_)
        nw311_[int(17)] = (self).f15
        nw311_[int(18)] = ((self).f5) * (self.f21)
        nw311_[int(19)] = default__.safeDivisionInt((self).f5, (self).f5)
        nw311_[int(20)] = (self).f5
        nw311_[int(21)] = (self.f21 if (self).f4 else (self).f5)
        nw311_[int(22)] = 384
        nw311_[int(23)] = self.f21
        nw311_[int(24)] = (self).f5
        nw311_[int(25)] = ((self).f15) + (195)
        nw311_[int(26)] = (self).f5
        d_1938_v18_ = nw311_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1938_v18_).length(0)):
            d_1939_i2_: int = guard_loop_4_
            if (True) and (((0) <= (d_1939_i2_)) and ((d_1939_i2_) < ((d_1938_v18_).length(0)))):
                (d_1938_v18_)[(d_1939_i2_)] = default__.safeDivisionInt(d_1939_i2_, (self).f5)
        d_1932_v12_ = (d_1932_v12_).set(len(_dafny.SeqWithoutIsStrInference([(self).f4, not((self).f16), (self).f16])), (self).f16)
        d_1940_v20_: _dafny.Array
        nw312_ = _dafny.Array(None, 5)
        nw312_[int(0)] = default__.fm60((self).f16, not((self).f16), (self).f4, (self).f5, globalState)
        def iife170_():
            coll74_ = _dafny.Map()
            compr_74_: int
            for compr_74_ in (d_1931_v11_).keys.Elements:
                d_1941_v19_: int = compr_74_
                if (d_1941_v19_) in (d_1931_v11_):
                    coll74_[(d_1941_v19_) * ((self).f15)] = d_1917_v0_
            return _dafny.Map(coll74_)
        nw312_[int(1)] = (iife170_()
        ) | (_dafny.Map({self.f21: d_1917_v0_}))
        nw312_[int(2)] = (d_1932_v12_) | (d_1932_v12_)
        nw312_[int(3)] = default__.fm60(d_1917_v0_, (self).f16, (self).f16, self.f21, globalState)
        nw312_[int(4)] = (d_1932_v12_) | (_dafny.Map({(self).f5: (self).f16}))
        d_1940_v20_ = nw312_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1940_v20_).length(0)):
            d_1942_i3_: int = guard_loop_5_
            if (True) and (((0) <= (d_1942_i3_)) and ((d_1942_i3_) < ((d_1940_v20_).length(0)))):
                (d_1940_v20_)[(d_1942_i3_)] = (d_1932_v12_) | ((D16_DC40(d_1932_v12_)).cf76)
        d_1943_v21_: _dafny.MultiSet
        d_1943_v21_ = _dafny.MultiSet([(self).f15, self.f21])
        (self).f21 = (d_1943_v21_).cardinality
        r0 = d_1919_v2_
        return r0

    @property
    def f22(self):
        return self._f22

class C7(T0, T1, T2):
    def  __init__(self):
        self._f18: bool = False
        self._f17: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: bool = False
        self._f5: int = int(0)
        self._f15: int = int(0)
        self._f16: bool = False
        self._f19: int = int(0)
        self._f20: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f18(self):
        return self._f18
    @f18.setter
    def f18(self, value):
        self._f18 = value
    @property
    def f17(self):
        return self._f17
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    def ctor__(self, f19, f20, f4, f5, f15, f16, f17, f18):
        (self)._f19 = f19
        (self)._f20 = f20
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self).f18 = f18

    def fm0(self, globalState):
        return not((not((self).f4) if (D5_DC13((self).f20, (self).f20)).cf29 else (self).f4))

    def fm22(self, globalState):
        return _dafny.CodePoint('v')

    def fm23(self, p0, globalState):
        return default__.safeModuloInt((self).f15, (self).f19)

    def m3(self, p0, globalState):
        d_1944_v0_: _dafny.Array
        nw313_ = _dafny.Array(False, 24)
        d_1944_v0_ = nw313_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_1944_v0_).length(0)):
            d_1945_i0_: int = guard_loop_6_
            if (True) and (((0) <= (d_1945_i0_)) and ((d_1945_i0_) < ((d_1944_v0_).length(0)))):
                (d_1944_v0_)[(d_1945_i0_)] = (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eespmunij")) if (self).f4 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpnvijih"))))) == ((self).f15)
        d_1946_i1_: int
        d_1946_i1_ = 0
        with _dafny.label("5"):
            while self.f18:
                with _dafny.c_label("5"):
                    if (d_1946_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_1946_i1_ = (d_1946_i1_) + (1)
                    d_1947_v1_: _dafny.Map
                    d_1947_v1_ = _dafny.Map({(self).f4: (self).f15})
                    d_1948_v2_: D4
                    d_1948_v2_ = D4_DC9(d_1947_v1_, p0, (self).f16)
                    source22_ = (D4_DC9(d_1947_v1_, (self).f15, (self).f16) if (d_1948_v2_).cf23 else D4_DC9(d_1947_v1_, (self).f19, (self).f16))
                    if source22_.is_DC9:
                        d_1949___mcc_h0_ = source22_.cf21
                        d_1950___mcc_h1_ = source22_.cf22
                        d_1951___mcc_h2_ = source22_.cf23
                        d_1952_cf23_ = d_1951___mcc_h2_
                        d_1953_cf22_ = d_1950___mcc_h1_
                        d_1954_cf21_ = d_1949___mcc_h0_
                        (self).f18 = not(d_1952_cf23_)
                        d_1955_v3_: _dafny.Array
                        nw314_ = _dafny.Array(_dafny.Array(None, 0), 8)
                        d_1955_v3_ = nw314_
                        d_1956_v4_: _dafny.Seq
                        d_1956_v4_ = _dafny.SeqWithoutIsStrInference([d_1944_v0_, d_1944_v0_])
                        index264_ = default__.safeIndex(273, (d_1955_v3_).length(0))
                        (d_1955_v3_)[index264_] = (d_1944_v0_ if (self).f16 else (d_1956_v4_)[default__.safeIndex((self).f5, len(d_1956_v4_))])
                        index265_ = default__.safeIndex(273, (d_1955_v3_).length(0))
                        (d_1955_v3_)[index265_] = d_1944_v0_
                        d_1957_v5_: _dafny.MultiSet
                        d_1957_v5_ = _dafny.MultiSet([(self).f20, (self).f4])
                        d_1952_cf23_ = not ((not(d_1952_cf23_)) == (self.f18)) or ((d_1957_v5_).ispropersubset(_dafny.MultiSet([(self).f20])))
                        d_1958_v6_: D0
                        d_1958_v6_ = D0_DC2((self).f16, self.f18)
                        d_1959_v7_: _dafny.Map
                        d_1959_v7_ = _dafny.Map({p0: d_1958_v6_})
                        d_1959_v7_ = d_1959_v7_
                    elif source22_.is_DC10:
                        d_1960___mcc_h3_ = source22_.cf24
                        d_1961___mcc_h4_ = source22_.cf25
                        d_1962___mcc_h5_ = source22_.cf26
                        d_1963_cf26_ = d_1962___mcc_h5_
                        d_1964_cf25_ = d_1961___mcc_h4_
                        d_1965_cf24_ = d_1960___mcc_h3_
                        (self).f18 = (self).f20
                        d_1966_v8_: int
                        d_1966_v8_ = -704
                        d_1966_v8_ = (0) - (p0)
                        d_1966_v8_ = 191
                        d_1967_v9_: _dafny.Array
                        nw315_ = _dafny.Array(_dafny.Array(None, 0), 11)
                        d_1967_v9_ = nw315_
                        d_1968_v10_: _dafny.Array
                        nw316_ = _dafny.Array(int(0), 22)
                        d_1968_v10_ = nw316_
                        d_1969_v11_: _dafny.Array
                        nw317_ = _dafny.Array(None, 10)
                        nw317_[int(0)] = d_1968_v10_
                        nw317_[int(1)] = d_1968_v10_
                        nw317_[int(2)] = d_1968_v10_
                        nw317_[int(3)] = d_1968_v10_
                        nw317_[int(4)] = d_1968_v10_
                        nw317_[int(5)] = d_1968_v10_
                        nw317_[int(6)] = d_1968_v10_
                        nw317_[int(7)] = d_1968_v10_
                        nw317_[int(8)] = d_1968_v10_
                        nw317_[int(9)] = d_1968_v10_
                        d_1969_v11_ = nw317_
                        index266_ = default__.safeIndex(603, (d_1967_v9_).length(0))
                        (d_1967_v9_)[index266_] = d_1969_v11_
                        d_1970_v12_: _dafny.Map
                        d_1970_v12_ = _dafny.Map({(self).f4: d_1969_v11_})
                        d_1971_v13_: _dafny.Seq
                        d_1971_v13_ = _dafny.SeqWithoutIsStrInference([d_1969_v11_, ((d_1970_v12_)[d_1964_cf25_] if (d_1964_cf25_) in (d_1970_v12_) else d_1969_v11_), d_1969_v11_, d_1969_v11_, d_1969_v11_])
                        d_1972_v14_: _dafny.Map
                        d_1972_v14_ = _dafny.Map({(0) - ((self).f15): self.f18})
                        d_1973_v15_: _dafny.Seq
                        d_1973_v15_ = _dafny.SeqWithoutIsStrInference([len(d_1972_v14_), p0])
                        index267_ = default__.safeIndex(603, (d_1967_v9_).length(0))
                        (d_1967_v9_)[index267_] = (d_1971_v13_)[default__.safeIndex((d_1973_v15_)[default__.safeIndex(p0, len(d_1973_v15_))], len(d_1971_v13_))]
                    elif source22_.is_DC8:
                        d_1974___mcc_h6_ = source22_.cf20
                        d_1975_cf20_ = d_1974___mcc_h6_
                        d_1976_v16_: int
                        d_1976_v16_ = -990
                        d_1976_v16_ = default__.safeModuloInt(default__.safeDivisionInt((self).fm23((self).f20, globalState), 948), len((self).f17))
                        d_1977_v17_: _dafny.Map
                        d_1977_v17_ = _dafny.Map({d_1976_v16_: self.f18})
                        d_1978_v18_: str
                        d_1978_v18_ = _dafny.CodePoint('k')
                        d_1979_v19_: _dafny.MultiSet
                        d_1979_v19_ = _dafny.MultiSet([d_1978_v18_, d_1978_v18_])
                        d_1980_v20_: _dafny.Seq
                        d_1980_v20_ = _dafny.SeqWithoutIsStrInference([d_1979_v19_, d_1979_v19_, d_1979_v19_, d_1979_v19_, d_1979_v19_])
                        d_1947_v1_ = (d_1947_v1_).set(((d_1977_v17_)[len((self).f17)] if (len((self).f17)) in (d_1977_v17_) else (self).f16), len(d_1980_v20_))
                        d_1981_v21_: _dafny.Array
                        nw318_ = _dafny.Array(None, 8)
                        nw318_[int(0)] = (self).fm23((self).f16, globalState)
                        nw318_[int(1)] = len((self).f17)
                        nw318_[int(2)] = 120
                        nw318_[int(3)] = (0) - (d_1976_v16_)
                        nw318_[int(4)] = d_1976_v16_
                        nw318_[int(5)] = (self).f15
                        nw318_[int(6)] = (self).fm23((self).f20, globalState)
                        nw318_[int(7)] = (0) - ((self).fm23(not((self).f4), globalState))
                        d_1981_v21_ = nw318_
                        index268_ = default__.safeIndex(28, (d_1981_v21_).length(0))
                        (d_1981_v21_)[index268_] = (self).f15
                        index269_ = default__.safeIndex(467, (d_1975_cf20_).length(0))
                        (d_1975_cf20_)[index269_] = (self).f20
                        d_1982_v22_: _dafny.Seq
                        d_1982_v22_ = _dafny.SeqWithoutIsStrInference([(self).f4])
                        d_1983_v23_: _dafny.MultiSet
                        d_1983_v23_ = _dafny.MultiSet([True])
                        d_1984_v24_: _dafny.Seq
                        d_1984_v24_ = _dafny.SeqWithoutIsStrInference([(self).f5])
                        d_1985_v25_: _dafny.Map
                        d_1985_v25_ = _dafny.Map({d_1984_v24_: False})
                        index270_ = default__.safeIndex(28, (d_1981_v21_).length(0))
                        index271_ = default__.safeIndex(467, (d_1975_cf20_).length(0))
                        rhs256_ = (self).fm0(globalState)
                        rhs257_ = (d_1982_v22_) <= (d_1982_v22_)
                        rhs258_ = d_1976_v16_
                        rhs259_ = ((d_1983_v23_).set((self).f16, default__.abs((0) - (d_1976_v16_)))) in (_dafny.Map({d_1983_v23_: len(d_1985_v25_)}))
                        lhs131_ = self
                        lhs132_ = self
                        lhs133_ = d_1981_v21_
                        lhs134_ = default__.safeIndex(28, (d_1981_v21_).length(0))
                        lhs135_ = d_1975_cf20_
                        lhs136_ = default__.safeIndex(467, (d_1975_cf20_).length(0))
                        lhs131_.f18 = rhs256_
                        lhs132_.f18 = rhs257_
                        lhs133_[lhs134_] = rhs258_
                        lhs135_[lhs136_] = rhs259_
                        d_1986_v26_: D0
                        d_1986_v26_ = D0_DC0(self.f18, ((d_1981_v21_)[default__.safeIndex(28, (d_1981_v21_).length(0))]) + ((self).f15), (self).f4, 787, self.f18)
                        d_1986_v26_ = default__.fm24(d_1984_v24_, (self).f20, ((self).f16) == (self.f18), D0_DC0(False, d_1976_v16_, (d_1975_cf20_)[default__.safeIndex(467, (d_1975_cf20_).length(0))], (self).f15, self.f18), globalState)
                    elif True:
                        d_1987___mcc_h7_ = source22_.cf27
                        d_1988_cf27_ = d_1987___mcc_h7_
                        (self).f18 = False
                        d_1989_v27_: _dafny.Map
                        d_1989_v27_ = _dafny.Map({len(_dafny.Set({(self).f17})): (self).f5})
                        d_1989_v27_ = (d_1989_v27_).set((self).f19, 454)
                        d_1990_v28_: _dafny.Array
                        nw319_ = _dafny.Array(_dafny.Map({}), 9)
                        d_1990_v28_ = nw319_
                        d_1991_v30_: _dafny.MultiSet
                        d_1991_v30_ = _dafny.MultiSet([self.f18, (self).f4])
                        d_1992_v31_: _dafny.MultiSet
                        d_1992_v31_ = d_1991_v30_
                        d_1993_v32_: _dafny.Map
                        d_1993_v32_ = _dafny.Map({d_1992_v31_: True})
                        d_1994_v33_: str
                        d_1994_v33_ = _dafny.CodePoint('n')
                        d_1995_v34_: _dafny.Map
                        d_1995_v34_ = _dafny.Map({d_1992_v31_: d_1994_v33_})
                        index272_ = default__.safeIndex(82, (d_1990_v28_).length(0))
                        def iife171_():
                            coll75_ = _dafny.Map()
                            compr_75_: _dafny.MultiSet
                            for compr_75_ in (d_1993_v32_).keys.Elements:
                                d_1996_v29_: _dafny.MultiSet = compr_75_
                                if (d_1996_v29_) in (d_1993_v32_):
                                    coll75_[d_1996_v29_] = _dafny.CodePoint('s')
                            return _dafny.Map(coll75_)
                        (d_1990_v28_)[index272_] = (iife171_()
                        ) | (d_1995_v34_)
                        index273_ = default__.safeIndex(82, (d_1990_v28_).length(0))
                        (d_1990_v28_)[index273_] = d_1995_v34_
                        d_1997_v35_: _dafny.Array
                        def lambda81_(d_1998_p0_):
                            def lambda82_(d_1999_i2_):
                                return (d_1999_i2_) - (d_1998_p0_)

                            return lambda82_

                        init49_ = lambda81_(p0)
                        nw320_ = _dafny.Array(None, 21)
                        for i0_49_ in range(nw320_.length(0)):
                            nw320_[i0_49_] = init49_(i0_49_)
                        d_1997_v35_ = nw320_
                        d_2000_v36_: _dafny.Seq
                        d_2000_v36_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                        d_2001_v37_: _dafny.Seq
                        d_2001_v37_ = _dafny.SeqWithoutIsStrInference([p0, (d_2000_v36_)[default__.safeIndex(len((self).f17), len(d_2000_v36_))], len(_dafny.Map({p0: self.f18}))])
                        index274_ = default__.safeIndex(584, (d_1997_v35_).length(0))
                        (d_1997_v35_)[index274_] = len(d_2001_v37_)
                        index275_ = default__.safeIndex(584, (d_1997_v35_).length(0))
                        (d_1997_v35_)[index275_] = ((self).f5) - ((0) - ((self).f15))
                    (self).f18 = (self).fm0(globalState)
                    d_2002_v38_: str
                    d_2002_v38_ = _dafny.CodePoint('c')
                    d_2002_v38_ = _dafny.CodePoint('v')
                    d_2003_v39_: _dafny.MultiSet
                    d_2003_v39_ = _dafny.MultiSet([(self).f17])
                    rhs260_ = (self).fm0(globalState)
                    rhs261_ = (d_2003_v39_) - ((d_2003_v39_) - (d_2003_v39_))
                    lhs137_ = self
                    lhs137_.f18 = rhs260_
                    d_2003_v39_ = rhs261_
                    pass
            pass
        d_2004_v40_: int
        d_2004_v40_ = 432
        d_2005_v41_: _dafny.Map
        d_2005_v41_ = _dafny.Map({not (False) or (not((self).fm0(globalState))): (self).f4})
        d_2006_v42_: _dafny.MultiSet
        d_2006_v42_ = _dafny.MultiSet([(self).f19, d_2004_v40_, (self).f19, p0])
        rhs262_ = ((0) - ((self).fm23((self).f20, globalState))) * ((self).f19)
        rhs263_ = ((d_2005_v41_) | (d_2005_v41_)).set((self).f16, (self).f20)
        rhs264_ = (((self).f20) and ((self).f20)) == ((self).f4)
        rhs265_ = (d_2006_v42_).issubset(d_2006_v42_)
        rhs266_ = (self).f20
        lhs138_ = self
        lhs139_ = self
        lhs140_ = self
        d_2004_v40_ = rhs262_
        d_2005_v41_ = rhs263_
        lhs138_.f18 = rhs264_
        lhs139_.f18 = rhs265_
        lhs140_.f18 = rhs266_
        d_2004_v40_ = ((self).f19) * (907)
        d_2007_v43_: _dafny.MultiSet
        d_2007_v43_ = _dafny.MultiSet([self.f18, (self).f4, (self).f20, (self).f16, (self).f16])
        d_2008_v44_: _dafny.Seq
        d_2008_v44_ = _dafny.SeqWithoutIsStrInference([((d_2007_v43_)[(self).f4] if ((self).f4) in (d_2007_v43_) else d_2004_v40_)])
        d_2009_v45_: _dafny.Map
        d_2009_v45_ = _dafny.Map({p0: (d_2008_v44_)[default__.safeIndex(p0, len(d_2008_v44_))]})
        d_2010_v46_: _dafny.Map
        d_2010_v46_ = _dafny.Map({d_2009_v45_: (self).f16})
        d_2011_v48_: _dafny.Set
        def iife172_():
            coll76_ = _dafny.Set()
            compr_76_: _dafny.Map
            for compr_76_ in (d_2010_v46_).keys.Elements:
                d_2012_v47_: _dafny.Map = compr_76_
                if (d_2012_v47_) in (d_2010_v46_):
                    coll76_ = coll76_.union(_dafny.Set([d_2012_v47_]))
            return _dafny.Set(coll76_)
        d_2011_v48_ = iife172_()
        
        d_2013_v49_: C0
        nw321_ = C0()
        nw321_.ctor__(d_2011_v48_, (self).f5)
        d_2013_v49_ = nw321_
        d_2014_i3_: int
        d_2014_i3_ = 0
        with _dafny.label("6"):
            while False:
                with _dafny.c_label("6"):
                    if (d_2014_i3_) >= (100):
                        raise _dafny.Break("6")
                    d_2014_i3_ = (d_2014_i3_) + (1)
                    d_2015_v50_: str
                    d_2015_v50_ = _dafny.CodePoint('d')
                    d_2016_v51_: C3
                    nw322_ = C3()
                    nw322_.ctor__(d_2015_v50_, (self).f15, (self).f20, (self).f4, ((d_2013_v49_).f27 if (self).f20 else p0))
                    d_2016_v51_ = nw322_
                    d_2017_v52_: _dafny.Seq
                    d_2017_v52_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2015_v50_ for d_2018_i4_ in range(default__.abs(-595))])])
                    (self).f18 = (((d_2017_v52_)[default__.safeIndex((self).f15, len(d_2017_v52_))]) + (default__.fm37(globalState))) == ((self).f17)
                    rhs267_ = d_2008_v44_
                    rhs268_ = (self).f4
                    rhs269_ = (self).f16
                    rhs270_ = (((self).f17) < ((self).f17)) == ((self).f16)
                    lhs141_ = self
                    lhs142_ = self
                    lhs143_ = self
                    d_2008_v44_ = rhs267_
                    lhs141_.f18 = rhs268_
                    lhs142_.f18 = rhs269_
                    lhs143_.f18 = rhs270_
                    if not(not((self).f16)):
                        d_2019_v53_: C2
                        nw323_ = C2()
                        nw323_.ctor__(d_1944_v0_, d_2006_v42_, ((self).f17).set(default__.safeIndex((self).f19, len((self).f17)), d_2015_v50_), (self).f16, (self.f18) and ((self).f16), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apdvmcjtv"))))
                        d_2019_v53_ = nw323_
                        d_2020_v54_: _dafny.Set
                        d_2020_v54_ = _dafny.Set({True})
                        d_2004_v40_ = (((0) - ((self).f19)) * (len((self).f17))) + (len((_dafny.Set({(self).f4, False, (self).f16, (self).f20})).intersection(d_2020_v54_)))
                        d_2021_v55_: C2
                        nw324_ = C2()
                        nw324_.ctor__(d_1944_v0_, d_2006_v42_, (self).f17, (self).f16, self.f18, (d_2013_v49_).f27)
                        d_2021_v55_ = nw324_
                        d_2004_v40_ = (self).f19
                        (self).f18 = (self).f4
                    elif True:
                        index276_ = default__.safeIndex(997, (d_1944_v0_).length(0))
                        (d_1944_v0_)[index276_] = (self).f20
                        index277_ = default__.safeIndex(997, (d_1944_v0_).length(0))
                        (d_1944_v0_)[index277_] = not((self).f20)
                        index278_ = default__.safeIndex(997, (d_1944_v0_).length(0))
                        (d_1944_v0_)[index278_] = ((_dafny.SeqWithoutIsStrInference([(d_2016_v51_).f25 for d_2022_i5_ in range(default__.abs(601))])) + ((self).f17)) == ((self).f17)
                        d_2023_v56_: _dafny.Seq
                        d_2023_v56_ = _dafny.SeqWithoutIsStrInference([(self).f4, True])
                        d_2024_v57_: _dafny.Seq
                        d_2024_v57_ = _dafny.SeqWithoutIsStrInference([d_2023_v56_])
                        d_2025_v58_: _dafny.Seq
                        d_2025_v58_ = _dafny.SeqWithoutIsStrInference([d_2023_v56_, (d_2024_v57_)[default__.safeIndex(d_2004_v40_, len(d_2024_v57_))], _dafny.SeqWithoutIsStrInference([False])])
                        d_2026_v59_: _dafny.Map
                        d_2026_v59_ = _dafny.Map({d_2024_v57_: self.f18})
                        index279_ = default__.safeIndex(997, (d_1944_v0_).length(0))
                        (d_1944_v0_)[index279_] = (((d_2026_v59_)[d_2025_v58_] if (d_2025_v58_) in (d_2026_v59_) else (self).f20) if self.f18 else (self).f4)
                        d_2027_v60_: _dafny.Array
                        def lambda83_(d_2028_i6_):
                            return (d_2028_i6_) + ((self).f5)

                        init50_ = lambda83_
                        nw325_ = _dafny.Array(None, 8)
                        for i0_50_ in range(nw325_.length(0)):
                            nw325_[i0_50_] = init50_(i0_50_)
                        d_2027_v60_ = nw325_
                        index280_ = default__.safeIndex(915, (d_2027_v60_).length(0))
                        (d_2027_v60_)[index280_] = (d_2013_v49_).f27
                        index281_ = default__.safeIndex(915, (d_2027_v60_).length(0))
                        (d_2027_v60_)[index281_] = ((d_2013_v49_).f27) + (((0) - ((d_2013_v49_).f27)) + (d_2004_v40_))
                        d_2027_v60_ = d_2027_v60_
                    pass
            pass

    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20

class C8:
    def  __init__(self):
        self.f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f14: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self, f13, f14):
        (self).f13 = f13
        (self)._f14 = f14

    def fm17(self, p0, p1, p2, p3, globalState):
        source23_ = D0_DC0(True, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ty")))), True, len(self.f13), True)
        if source23_.is_DC0:
            d_2029___mcc_h0_ = source23_.cf0
            d_2030___mcc_h1_ = source23_.cf1
            d_2031___mcc_h2_ = source23_.cf2
            d_2032___mcc_h3_ = source23_.cf3
            d_2033___mcc_h4_ = source23_.cf4
            d_2034_cf4_ = d_2033___mcc_h4_
            d_2035_cf3_ = d_2032___mcc_h3_
            d_2036_cf2_ = d_2031___mcc_h2_
            d_2037_cf1_ = d_2030___mcc_h1_
            d_2038_cf0_ = d_2029___mcc_h0_
            return d_2038_cf0_
        elif source23_.is_DC1:
            d_2039___mcc_h5_ = source23_.cf5
            d_2040___mcc_h6_ = source23_.cf6
            d_2041___mcc_h7_ = source23_.cf7
            d_2042___mcc_h8_ = source23_.cf8
            d_2043_cf8_ = d_2042___mcc_h8_
            d_2044_cf7_ = d_2041___mcc_h7_
            d_2045_cf6_ = d_2040___mcc_h6_
            d_2046_cf5_ = d_2039___mcc_h5_
            return (d_2046_cf5_) >= ((self).f14)
        elif True:
            d_2047___mcc_h9_ = source23_.cf9
            d_2048___mcc_h10_ = source23_.cf10
            d_2049_cf10_ = d_2048___mcc_h10_
            d_2050_cf9_ = d_2047___mcc_h9_
            return ((self).f14) > ((self).f14)

    def m10(self, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        r3: _dafny.Seq = _dafny.Seq({})
        d_2051_v0_: bool
        d_2051_v0_ = True
        hi6_ = (self).f14
        for d_2052_i0_ in range(((self).f14 if not(d_2051_v0_) else (self).f14), hi6_):
            d_2053_v1_: _dafny.Map
            d_2053_v1_ = _dafny.Map({len(self.f13): (self).f14})
            d_2054_v2_: _dafny.MultiSet
            d_2054_v2_ = _dafny.MultiSet([d_2051_v0_])
            d_2055_v3_: _dafny.Array
            nw326_ = _dafny.Array(None, 7)
            nw326_[int(0)] = (self).f14
            nw326_[int(1)] = (self).f14
            nw326_[int(2)] = (self).f14
            nw326_[int(3)] = (self).f14
            nw326_[int(4)] = ((self).f14) * (len(d_2053_v1_))
            nw326_[int(5)] = (d_2054_v2_).cardinality
            nw326_[int(6)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfxpmukgb")))
            d_2055_v3_ = nw326_
            index282_ = default__.safeIndex(945, (d_2055_v3_).length(0))
            (d_2055_v3_)[index282_] = d_2052_i0_
            index283_ = default__.safeIndex(945, (d_2055_v3_).length(0))
            (d_2055_v3_)[index283_] = (self).f14
            index284_ = default__.safeIndex(945, (d_2055_v3_).length(0))
            (d_2055_v3_)[index284_] = 897
            d_2056_v4_: _dafny.Array
            nw327_ = _dafny.Array(_dafny.Map({}), 13)
            d_2056_v4_ = nw327_
            d_2057_v5_: _dafny.Array
            def lambda84_(d_2058_v0_):
                def lambda85_(d_2059_i1_):
                    return (d_2058_v0_) == (d_2058_v0_)

                return lambda85_

            init51_ = lambda84_(d_2051_v0_)
            nw328_ = _dafny.Array(None, 2)
            for i0_51_ in range(nw328_.length(0)):
                nw328_[i0_51_] = init51_(i0_51_)
            d_2057_v5_ = nw328_
            d_2060_v6_: _dafny.Set
            d_2060_v6_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([False, d_2051_v0_])})
            index285_ = default__.safeIndex(232, (d_2057_v5_).length(0))
            (d_2057_v5_)[index285_] = (d_2060_v6_).issubset(d_2060_v6_)
            d_2061_v7_: D1
            d_2061_v7_ = D1_DC5(d_2051_v0_, d_2052_i0_, self.f13)
            index286_ = default__.safeIndex(945, (d_2055_v3_).length(0))
            index287_ = default__.safeIndex(232, (d_2057_v5_).length(0))
            rhs271_ = (D0_DC1(d_2052_i0_, d_2051_v0_, (d_2055_v3_)[default__.safeIndex(945, (d_2055_v3_).length(0))], d_2052_i0_)).cf8
            rhs272_ = d_2056_v4_
            rhs273_ = (d_2061_v7_).cf15
            lhs144_ = d_2055_v3_
            lhs145_ = default__.safeIndex(945, (d_2055_v3_).length(0))
            lhs146_ = d_2057_v5_
            lhs147_ = default__.safeIndex(232, (d_2057_v5_).length(0))
            lhs144_[lhs145_] = rhs271_
            d_2056_v4_ = rhs272_
            lhs146_[lhs147_] = rhs273_
            d_2062_v8_: _dafny.Seq
            d_2062_v8_ = _dafny.SeqWithoutIsStrInference([not(d_2051_v0_)])
            d_2063_v9_: D0
            d_2063_v9_ = D0_DC0(d_2051_v0_, (self).f14, (d_2057_v5_)[default__.safeIndex(232, (d_2057_v5_).length(0))], (self).f14, (d_2057_v5_)[default__.safeIndex(232, (d_2057_v5_).length(0))])
            d_2064_v10_: _dafny.Map
            d_2064_v10_ = _dafny.Map({d_2051_v0_: d_2063_v9_})
            d_2065_v11_: _dafny.Seq
            d_2065_v11_ = _dafny.SeqWithoutIsStrInference([default__.fm18((d_2057_v5_)[default__.safeIndex(232, (d_2057_v5_).length(0))], (self).fm17(d_2062_v8_, (d_2055_v3_)[default__.safeIndex(945, (d_2055_v3_).length(0))], d_2051_v0_, -995, globalState), (d_2057_v5_)[default__.safeIndex(232, (d_2057_v5_).length(0))], globalState), (self).f14, len(d_2064_v10_)])
            d_2066_v12_: _dafny.Map
            d_2066_v12_ = _dafny.Map({(0) - (((d_2065_v11_)[default__.safeIndex((d_2055_v3_)[default__.safeIndex(945, (d_2055_v3_).length(0))], len(d_2065_v11_))]) * ((d_2055_v3_)[default__.safeIndex(945, (d_2055_v3_).length(0))])): default__.fm19(globalState)})
            d_2067_v13_: _dafny.Seq
            d_2067_v13_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2051_v0_: d_2062_v8_})])
            d_2068_v14_: _dafny.Map
            d_2068_v14_ = _dafny.Map({False: d_2062_v8_})
            d_2066_v12_ = (d_2066_v12_).set((d_2055_v3_)[default__.safeIndex(945, (d_2055_v3_).length(0))], (((d_2067_v13_)[default__.safeIndex(len(self.f13), len(d_2067_v13_))]).set(d_2051_v0_, _dafny.SeqWithoutIsStrInference([d_2051_v0_]))) | (d_2068_v14_))
        d_2069_v15_: str
        d_2069_v15_ = _dafny.CodePoint('f')
        d_2070_v16_: _dafny.MultiSet
        d_2070_v16_ = _dafny.MultiSet([d_2051_v0_, not(d_2051_v0_)])
        source24_ = default__.fm20(not(d_2051_v0_), d_2069_v15_, 430, default__.safeModuloInt((d_2070_v16_).cardinality, (0) - (len(self.f13))), globalState)
        if source24_.is_DC4:
            d_2071___mcc_h0_ = source24_.cf12
            d_2072___mcc_h1_ = source24_.cf13
            d_2073___mcc_h2_ = source24_.cf14
            d_2074_cf14_ = d_2073___mcc_h2_
            d_2075_cf13_ = d_2072___mcc_h1_
            d_2076_cf12_ = d_2071___mcc_h0_
            d_2075_cf13_ = d_2075_cf13_
            d_2077_v17_: D1
            d_2077_v17_ = D1_DC4((d_2051_v0_ if d_2076_cf12_ else d_2051_v0_), d_2075_cf13_, default__.safeDivisionInt(d_2074_cf14_, d_2075_cf13_))
            source25_ = d_2077_v17_
            if source25_.is_DC4:
                d_2078___mcc_h7_ = source25_.cf12
                d_2079___mcc_h8_ = source25_.cf13
                d_2080___mcc_h9_ = source25_.cf14
                d_2081_cf14_ = d_2080___mcc_h9_
                d_2082_cf13_ = d_2079___mcc_h8_
                d_2083_cf12_ = d_2078___mcc_h7_
                r2 = d_2051_v0_
                d_2084_v18_: _dafny.Seq
                d_2084_v18_ = _dafny.SeqWithoutIsStrInference([d_2083_cf12_, d_2083_cf12_, not(d_2051_v0_)])
                d_2085_v19_: _dafny.Array
                nw329_ = _dafny.Array(None, 10)
                nw329_[int(0)] = d_2084_v18_
                nw329_[int(1)] = d_2084_v18_
                nw329_[int(2)] = default__.fm21(globalState)
                nw329_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_2051_v0_, not(not(d_2051_v0_))])) + (_dafny.SeqWithoutIsStrInference([d_2051_v0_, d_2083_cf12_, d_2051_v0_]))
                nw329_[int(4)] = d_2084_v18_
                nw329_[int(5)] = d_2084_v18_
                nw329_[int(6)] = (d_2084_v18_).set(default__.safeIndex(d_2081_cf14_, len(d_2084_v18_)), d_2051_v0_)
                nw329_[int(7)] = d_2084_v18_
                nw329_[int(8)] = (d_2084_v18_).set(default__.safeIndex(d_2082_cf13_, len(d_2084_v18_)), d_2051_v0_)
                nw329_[int(9)] = d_2084_v18_
                d_2085_v19_ = nw329_
                index288_ = default__.safeIndex(504, (d_2085_v19_).length(0))
                (d_2085_v19_)[index288_] = (d_2084_v18_) + (d_2084_v18_)
                index289_ = default__.safeIndex(504, (d_2085_v19_).length(0))
                (d_2085_v19_)[index289_] = (_dafny.SeqWithoutIsStrInference([d_2083_cf12_])) + (_dafny.SeqWithoutIsStrInference([d_2051_v0_, True]))
                d_2082_cf13_ = d_2081_cf14_
                d_2086_v20_: _dafny.Map
                d_2086_v20_ = _dafny.Map({d_2083_cf12_: (d_2084_v18_).set(default__.safeIndex(d_2081_cf14_, len(d_2084_v18_)), d_2051_v0_)})
                d_2082_cf13_ = len(d_2086_v20_)
            elif source25_.is_DC5:
                d_2087___mcc_h10_ = source25_.cf15
                d_2088___mcc_h11_ = source25_.cf16
                d_2089___mcc_h12_ = source25_.cf17
                d_2090_cf17_ = d_2089___mcc_h12_
                d_2091_cf16_ = d_2088___mcc_h11_
                d_2092_cf15_ = d_2087___mcc_h10_
                d_2092_cf15_ = d_2092_cf15_
                d_2093_v21_: _dafny.Array
                def lambda86_(d_2094_cf15_):
                    def lambda87_(d_2095_i2_):
                        return d_2094_cf15_

                    return lambda87_

                init52_ = lambda86_(d_2092_cf15_)
                nw330_ = _dafny.Array(None, 10)
                for i0_52_ in range(nw330_.length(0)):
                    nw330_[i0_52_] = init52_(i0_52_)
                d_2093_v21_ = nw330_
                d_2096_v22_: D4
                d_2096_v22_ = D4_DC8(d_2093_v21_)
                pat_let_tv26_ = d_2093_v21_
                d_2097_v23_: _dafny.MultiSet
                def iife173_(_pat_let48_0):
                    def iife174_(d_2098_dt__update__tmp_h0_):
                        def iife175_(_pat_let49_0):
                            def iife176_(d_2099_dt__update_hcf20_h0_):
                                return D4_DC8(d_2099_dt__update_hcf20_h0_)
                            return iife176_(_pat_let49_0)
                        return iife175_(pat_let_tv26_)
                    return iife174_(_pat_let48_0)
                d_2097_v23_ = _dafny.MultiSet([iife173_(d_2096_v22_), D4_DC8(d_2093_v21_), d_2096_v22_])
                d_2097_v23_ = (d_2097_v23_) | (d_2097_v23_)
                d_2100_v24_: _dafny.Map
                d_2100_v24_ = _dafny.Map({d_2091_cf16_: d_2076_cf12_})
                d_2101_v25_: _dafny.Set
                d_2101_v25_ = _dafny.Set({d_2100_v24_})
                d_2101_v25_ = d_2101_v25_
                d_2102_v26_: _dafny.Array
                nw331_ = _dafny.Array(None, 11)
                nw331_[int(0)] = (d_2093_v21_ if d_2051_v0_ else d_2093_v21_)
                nw331_[int(1)] = d_2093_v21_
                nw331_[int(2)] = d_2093_v21_
                nw331_[int(3)] = d_2093_v21_
                nw331_[int(4)] = d_2093_v21_
                nw331_[int(5)] = d_2093_v21_
                nw331_[int(6)] = (d_2093_v21_ if not(d_2076_cf12_) else d_2093_v21_)
                nw331_[int(7)] = d_2093_v21_
                nw331_[int(8)] = (d_2093_v21_ if (D1_DC4(not(d_2076_cf12_), d_2075_cf13_, d_2091_cf16_)).cf12 else d_2093_v21_)
                nw331_[int(9)] = d_2093_v21_
                nw331_[int(10)] = d_2093_v21_
                d_2102_v26_ = nw331_
                index290_ = default__.safeIndex(990, (d_2102_v26_).length(0))
                (d_2102_v26_)[index290_] = d_2093_v21_
                index291_ = default__.safeIndex(990, (d_2102_v26_).length(0))
                (d_2102_v26_)[index291_] = d_2093_v21_
            elif True:
                d_2103___mcc_h13_ = source25_.cf11
                d_2104_cf11_ = d_2103___mcc_h13_
                d_2105_v27_: _dafny.Map
                d_2105_v27_ = _dafny.Map({d_2076_cf12_: d_2075_cf13_})
                d_2106_v28_: _dafny.Map
                d_2106_v28_ = _dafny.Map({False: d_2051_v0_})
                d_2105_v27_ = (d_2105_v27_).set((((d_2106_v28_)[d_2076_cf12_] if (d_2076_cf12_) in (d_2106_v28_) else d_2051_v0_)) == (False), d_2074_cf14_)
                d_2107_v29_: C5
                nw332_ = C5()
                nw332_.ctor__()
                d_2107_v29_ = nw332_
                d_2108_v30_: _dafny.Map
                d_2108_v30_ = _dafny.Map({(self).f14: not(((d_2106_v28_)[d_2051_v0_] if (d_2051_v0_) in (d_2106_v28_) else not(default__.fm27(d_2069_v15_, globalState))))})
                d_2109_v31_: _dafny.MultiSet
                d_2109_v31_ = _dafny.MultiSet([d_2075_cf13_])
                d_2110_v32_: D17
                d_2110_v32_ = D17_DC44(d_2109_v31_)
                d_2111_v33_: _dafny.Array
                nw333_ = _dafny.Array(None, 27)
                nw333_[int(0)] = d_2075_cf13_
                nw333_[int(1)] = d_2075_cf13_
                nw333_[int(2)] = d_2075_cf13_
                nw333_[int(3)] = len(_dafny.Set({d_2074_cf14_}))
                nw333_[int(4)] = 981
                nw333_[int(5)] = d_2074_cf14_
                nw333_[int(6)] = d_2075_cf13_
                nw333_[int(7)] = (self).f14
                nw333_[int(8)] = (self).f14
                nw333_[int(9)] = len(d_2104_cf11_)
                nw333_[int(10)] = d_2075_cf13_
                nw333_[int(11)] = (self).f14
                nw333_[int(12)] = d_2074_cf14_
                nw333_[int(13)] = d_2075_cf13_
                nw333_[int(14)] = (0) - ((self).f14)
                nw333_[int(15)] = (self).f14
                nw333_[int(16)] = 713
                nw333_[int(17)] = 834
                nw333_[int(18)] = -159
                nw333_[int(19)] = len(self.f13)
                nw333_[int(20)] = (self).f14
                nw333_[int(21)] = d_2075_cf13_
                nw333_[int(22)] = -131
                nw333_[int(23)] = d_2074_cf14_
                nw333_[int(24)] = (self).f14
                nw333_[int(25)] = 985
                nw333_[int(26)] = ((d_2110_v32_).cf82).cardinality
                d_2111_v33_ = nw333_
                d_2112_v34_: _dafny.Map
                d_2112_v34_ = _dafny.Map({((_dafny.MultiSet(d_2104_cf11_)).set(not(d_2076_cf12_), default__.abs(len(self.f13)))).cardinality: default__.fm25(d_2075_cf13_, (D11_DC30(d_2051_v0_, d_2111_v33_, d_2075_cf13_, d_2051_v0_)).cf63, d_2076_cf12_, globalState)})
                d_2113_v35_: _dafny.Map
                d_2113_v35_ = _dafny.Map({len(d_2112_v34_): _dafny.Map({(self).f14: d_2076_cf12_})})
                d_2114_v37_: _dafny.Seq
                d_2114_v37_ = _dafny.SeqWithoutIsStrInference([d_2075_cf13_])
                d_2115_v38_: _dafny.Array
                nw334_ = _dafny.Array(None, 8)
                nw334_[int(0)] = d_2108_v30_
                nw334_[int(1)] = (d_2108_v30_) | (d_2108_v30_)
                nw334_[int(2)] = ((d_2113_v35_)[d_2074_cf14_] if (d_2074_cf14_) in (d_2113_v35_) else d_2108_v30_)
                nw334_[int(3)] = d_2108_v30_
                nw334_[int(4)] = d_2108_v30_
                nw334_[int(5)] = d_2108_v30_
                def iife177_():
                    coll77_ = _dafny.Map()
                    compr_77_: int
                    for compr_77_ in (d_2114_v37_).Elements:
                        d_2116_v36_: int = compr_77_
                        if (d_2116_v36_) in (d_2114_v37_):
                            coll77_[(d_2116_v36_) * ((0) - ((self).f14))] = not(not(d_2051_v0_))
                    return _dafny.Map(coll77_)
                nw334_[int(6)] = (iife177_()
                ).set(d_2075_cf13_, d_2076_cf12_)
                nw334_[int(7)] = (default__.fm60(d_2076_cf12_, d_2076_cf12_, d_2051_v0_, (0) - ((self).f14), globalState)) | (_dafny.Map({(self).f14: d_2076_cf12_}))
                d_2115_v38_ = nw334_
                index292_ = default__.safeIndex(406, (d_2115_v38_).length(0))
                (d_2115_v38_)[index292_] = d_2108_v30_
                index293_ = default__.safeIndex(406, (d_2115_v38_).length(0))
                (d_2115_v38_)[index293_] = d_2108_v30_
                d_2112_v34_ = d_2112_v34_
            d_2117_v39_: _dafny.Array
            def lambda88_(d_2118_cf13_):
                def lambda89_(d_2119_i3_):
                    return (d_2119_i3_) - (d_2118_cf13_)

                return lambda89_

            init53_ = lambda88_(d_2075_cf13_)
            nw335_ = _dafny.Array(None, 29)
            for i0_53_ in range(nw335_.length(0)):
                nw335_[i0_53_] = init53_(i0_53_)
            d_2117_v39_ = nw335_
            index294_ = default__.safeIndex(84, (d_2117_v39_).length(0))
            (d_2117_v39_)[index294_] = d_2075_cf13_
            index295_ = default__.safeIndex(84, (d_2117_v39_).length(0))
            (d_2117_v39_)[index295_] = (self).f14
            d_2120_v40_: D16
            d_2120_v40_ = D16_DC40(_dafny.Map({(self).f14: d_2076_cf12_}))
            source26_ = d_2120_v40_
            if source26_.is_DC41:
                d_2121___mcc_h14_ = source26_.cf77
                d_2122___mcc_h15_ = source26_.cf78
                d_2123_cf78_ = d_2122___mcc_h15_
                d_2124_cf77_ = d_2121___mcc_h14_
                r3 = _dafny.SeqWithoutIsStrInference([(self).f14, d_2074_cf14_, default__.fm18(d_2076_cf12_, False, d_2076_cf12_, globalState)])
                d_2125_v41_: _dafny.Map
                d_2125_v41_ = _dafny.Map({(self).f14: (self).f14})
                d_2126_v42_: _dafny.Seq
                d_2126_v42_ = _dafny.SeqWithoutIsStrInference([d_2075_cf13_])
                d_2127_v43_: _dafny.Set
                d_2127_v43_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tehtab")))})
                d_2128_v44_: _dafny.Array
                nw336_ = _dafny.Array(False, 27)
                d_2128_v44_ = nw336_
                d_2129_v45_: D11
                d_2129_v45_ = D11_DC29((d_2126_v42_)[default__.safeIndex(len(d_2127_v43_), len(d_2126_v42_))], d_2128_v44_, d_2123_cf78_, False)
                d_2125_v41_ = (d_2125_v41_).set((d_2129_v45_).cf57, (self).f14)
                d_2130_v46_: _dafny.Array
                nw337_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_2130_v46_ = nw337_
                index296_ = default__.safeIndex(948, (d_2130_v46_).length(0))
                (d_2130_v46_)[index296_] = d_2117_v39_
                d_2131_v47_: D16
                d_2131_v47_ = D16_DC41(d_2124_cf77_, d_2051_v0_)
                index297_ = default__.safeIndex(948, (d_2130_v46_).length(0))
                rhs274_ = d_2051_v0_
                rhs275_ = (d_2131_v47_).cf77
                rhs276_ = d_2117_v39_
                lhs148_ = d_2130_v46_
                lhs149_ = default__.safeIndex(948, (d_2130_v46_).length(0))
                r1 = rhs274_
                d_2069_v15_ = rhs275_
                lhs148_[lhs149_] = rhs276_
                d_2117_v39_ = (d_2130_v46_)[default__.safeIndex(948, (d_2130_v46_).length(0))]
            elif source26_.is_DC42:
                d_2132___mcc_h16_ = source26_.cf79
                d_2133___mcc_h17_ = source26_.cf80
                d_2134_cf80_ = d_2133___mcc_h17_
                d_2135_cf79_ = d_2132___mcc_h16_
                d_2136_v48_: D4
                d_2136_v48_ = D4_DC9(_dafny.Map({True: d_2074_cf14_}), (d_2117_v39_)[default__.safeIndex(84, (d_2117_v39_).length(0))], d_2134_cf80_)
                d_2137_v49_: _dafny.Set
                d_2137_v49_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spsfcpd")), self.f13})
                d_2136_v48_ = default__.fm71((self.f13) + (self.f13), (d_2117_v39_)[default__.safeIndex(84, (d_2117_v39_).length(0))], d_2137_v49_, _dafny.CodePoint('r'), globalState)
                d_2138_v50_: _dafny.Set
                d_2138_v50_ = _dafny.Set({(d_2117_v39_)[default__.safeIndex(84, (d_2117_v39_).length(0))]})
                d_2139_v51_: C1
                nw338_ = C1()
                nw338_.ctor__(d_2074_cf14_, False, d_2134_cf80_, len((d_2138_v50_) - (d_2138_v50_)))
                d_2139_v51_ = nw338_
                d_2069_v15_ = d_2069_v15_
                d_2140_v52_: _dafny.Array
                nw339_ = _dafny.Array(False, 1)
                d_2140_v52_ = nw339_
                d_2141_v53_: _dafny.Seq
                d_2141_v53_ = _dafny.SeqWithoutIsStrInference([(d_2117_v39_)[default__.safeIndex(84, (d_2117_v39_).length(0))]])
                d_2142_v54_: D11
                d_2142_v54_ = D11_DC29(((d_2070_v16_)[d_2076_cf12_] if (d_2076_cf12_) in (d_2070_v16_) else d_2075_cf13_), d_2140_v52_, d_2051_v0_, ((self).f14) == ((d_2141_v53_)[default__.safeIndex((d_2117_v39_)[default__.safeIndex(84, (d_2117_v39_).length(0))], len(d_2141_v53_))]))
                index298_ = default__.safeIndex(84, (d_2117_v39_).length(0))
                rhs277_ = ((self).f14) - ((self).f14)
                rhs278_ = d_2142_v54_
                rhs279_ = d_2051_v0_
                rhs280_ = (self).f14
                lhs150_ = d_2117_v39_
                lhs151_ = default__.safeIndex(84, (d_2117_v39_).length(0))
                lhs150_[lhs151_] = rhs277_
                d_2142_v54_ = rhs278_
                r1 = rhs279_
                d_2075_cf13_ = rhs280_
            elif source26_.is_DC40:
                d_2143___mcc_h18_ = source26_.cf76
                d_2144_cf76_ = d_2143___mcc_h18_
                d_2145_v55_: _dafny.Array
                def lambda90_(d_2146_v0_, d_2147_cf12_):
                    def lambda91_(d_2148_i4_):
                        return _dafny.SeqWithoutIsStrInference([d_2146_v0_, d_2147_cf12_])

                    return lambda91_

                init54_ = lambda90_(d_2051_v0_, d_2076_cf12_)
                nw340_ = _dafny.Array(None, 23)
                for i0_54_ in range(nw340_.length(0)):
                    nw340_[i0_54_] = init54_(i0_54_)
                d_2145_v55_ = nw340_
                d_2149_v56_: _dafny.Seq
                d_2149_v56_ = _dafny.SeqWithoutIsStrInference([d_2051_v0_, d_2051_v0_, d_2051_v0_, d_2076_cf12_])
                index299_ = default__.safeIndex(688, (d_2145_v55_).length(0))
                (d_2145_v55_)[index299_] = d_2149_v56_
                d_2150_v58_: _dafny.Map
                d_2150_v58_ = _dafny.Map({d_2051_v0_: True})
                d_2151_v59_: _dafny.Seq
                d_2151_v59_ = _dafny.SeqWithoutIsStrInference([(d_2117_v39_)[default__.safeIndex(84, (d_2117_v39_).length(0))], len(d_2150_v58_)])
                index300_ = default__.safeIndex(84, (d_2117_v39_).length(0))
                index301_ = default__.safeIndex(688, (d_2145_v55_).length(0))
                def iife178_():
                    coll78_ = _dafny.Map()
                    compr_78_: int
                    for compr_78_ in _dafny.IntegerRange(487, 552):
                        d_2152_v57_: int = compr_78_
                        if ((487) <= (d_2152_v57_)) and ((d_2152_v57_) < (552)):
                            coll78_[(d_2152_v57_) + (d_2075_cf13_)] = (d_2149_v56_)[default__.safeIndex(d_2074_cf14_, len(d_2149_v56_))]
                    return _dafny.Map(coll78_)
                rhs281_ = ((d_2144_cf76_) | (iife178_()
                )).set(default__.safeDivisionInt((d_2117_v39_)[default__.safeIndex(84, (d_2117_v39_).length(0))], d_2074_cf14_), not((_dafny.SeqWithoutIsStrInference([-492])) == (d_2151_v59_)))
                rhs282_ = len((self.f13) + (self.f13))
                rhs283_ = (d_2149_v56_) + (_dafny.SeqWithoutIsStrInference([d_2076_cf12_, d_2076_cf12_]))
                lhs152_ = d_2117_v39_
                lhs153_ = default__.safeIndex(84, (d_2117_v39_).length(0))
                lhs154_ = d_2145_v55_
                lhs155_ = default__.safeIndex(688, (d_2145_v55_).length(0))
                d_2144_cf76_ = rhs281_
                lhs152_[lhs153_] = rhs282_
                lhs154_[lhs155_] = rhs283_
                d_2153_v60_: _dafny.Array
                def lambda92_(d_2154_v16_):
                    def lambda93_(d_2155_i5_):
                        return d_2154_v16_

                    return lambda93_

                init55_ = lambda92_(d_2070_v16_)
                nw341_ = _dafny.Array(None, 3)
                for i0_55_ in range(nw341_.length(0)):
                    nw341_[i0_55_] = init55_(i0_55_)
                d_2153_v60_ = nw341_
                index302_ = default__.safeIndex(221, (d_2153_v60_).length(0))
                (d_2153_v60_)[index302_] = d_2070_v16_
                d_2156_v61_: _dafny.Array
                nw342_ = _dafny.Array(_dafny.Seq({}), 25)
                d_2156_v61_ = nw342_
                index303_ = default__.safeIndex(221, (d_2153_v60_).length(0))
                rhs284_ = ((d_2070_v16_).intersection(d_2070_v16_)) | (d_2070_v16_)
                rhs285_ = d_2156_v61_
                lhs156_ = d_2153_v60_
                lhs157_ = default__.safeIndex(221, (d_2153_v60_).length(0))
                lhs156_[lhs157_] = rhs284_
                d_2156_v61_ = rhs285_
                d_2051_v0_ = not(d_2076_cf12_)
                d_2117_v39_ = d_2117_v39_
            elif True:
                d_2157___mcc_h19_ = source26_.cf81
                d_2158_cf81_ = d_2157___mcc_h19_
                (self).f13 = self.f13
                d_2159_v62_: _dafny.Map
                d_2159_v62_ = _dafny.Map({default__.fm27(_dafny.CodePoint('b'), globalState): d_2069_v15_})
                d_2159_v62_ = (d_2159_v62_).set(d_2051_v0_, d_2069_v15_)
                index304_ = default__.safeIndex(84, (d_2117_v39_).length(0))
                (d_2117_v39_)[index304_] = (self).f14
                d_2075_cf13_ = default__.fm26(globalState)
        elif source24_.is_DC5:
            d_2160___mcc_h3_ = source24_.cf15
            d_2161___mcc_h4_ = source24_.cf16
            d_2162___mcc_h5_ = source24_.cf17
            d_2163_cf17_ = d_2162___mcc_h5_
            d_2164_cf16_ = d_2161___mcc_h4_
            d_2165_cf15_ = d_2160___mcc_h3_
            r2 = d_2051_v0_
            d_2166_v63_: _dafny.Seq
            d_2166_v63_ = _dafny.SeqWithoutIsStrInference([d_2051_v0_])
            d_2167_v64_: D17
            d_2167_v64_ = D17_DC45()
            d_2168_v65_: _dafny.Array
            def lambda94_(d_2169_cf16_):
                def lambda95_(d_2170_i7_):
                    return _dafny.Set({_dafny.Map({d_2169_cf16_: (self).f14})})

                return lambda95_

            init56_ = lambda94_(d_2164_cf16_)
            nw343_ = _dafny.Array(None, 24)
            for i0_56_ in range(nw343_.length(0)):
                nw343_[i0_56_] = init56_(i0_56_)
            d_2168_v65_ = nw343_
            d_2171_v66_: _dafny.Map
            d_2171_v66_ = _dafny.Map({d_2167_v64_: d_2168_v65_})
            d_2172_v67_: D1
            d_2172_v67_ = D1_DC5(d_2165_cf15_, 93, _dafny.SeqWithoutIsStrInference([d_2069_v15_ for d_2173_i8_ in range(default__.abs(-868))]))
            d_2174_v68_: _dafny.Array
            nw344_ = _dafny.Array(None, 18)
            nw344_[int(0)] = d_2051_v0_
            nw344_[int(1)] = d_2165_cf15_
            nw344_[int(2)] = (_dafny.SeqWithoutIsStrInference([d_2069_v15_ for d_2175_i6_ in range(default__.abs(711))])) < (d_2163_cf17_)
            nw344_[int(3)] = default__.fm27(d_2069_v15_, globalState)
            nw344_[int(4)] = d_2051_v0_
            nw344_[int(5)] = d_2165_cf15_
            nw344_[int(6)] = d_2051_v0_
            nw344_[int(7)] = True
            nw344_[int(8)] = d_2051_v0_
            nw344_[int(9)] = d_2165_cf15_
            nw344_[int(10)] = d_2165_cf15_
            nw344_[int(11)] = (self.f13) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")))
            nw344_[int(12)] = d_2051_v0_
            nw344_[int(13)] = (self).fm17(d_2166_v63_, d_2164_cf16_, d_2051_v0_, len((d_2171_v66_).set(default__.fm72(globalState), d_2168_v65_)), globalState)
            nw344_[int(14)] = d_2165_cf15_
            nw344_[int(15)] = (d_2172_v67_).cf15
            nw344_[int(16)] = d_2165_cf15_
            nw344_[int(17)] = d_2165_cf15_
            d_2174_v68_ = nw344_
            index305_ = default__.safeIndex(435, (d_2174_v68_).length(0))
            (d_2174_v68_)[index305_] = default__.fm27(d_2069_v15_, globalState)
            d_2176_v69_: _dafny.Set
            d_2176_v69_ = _dafny.Set({not(d_2051_v0_)})
            d_2177_v70_: _dafny.Set
            d_2177_v70_ = _dafny.Set({len(d_2176_v69_), (self).f14})
            d_2178_v71_: _dafny.MultiSet
            d_2178_v71_ = _dafny.MultiSet([d_2177_v70_])
            d_2179_v72_: _dafny.Seq
            d_2179_v72_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_2051_v0_])])
            index306_ = default__.safeIndex(435, (d_2174_v68_).length(0))
            (d_2174_v68_)[index306_] = (self).fm17(d_2166_v63_, (d_2178_v71_).cardinality, d_2165_cf15_, len(d_2179_v72_), globalState)
            d_2180_v73_: _dafny.Array
            def lambda96_(d_2181_i9_):
                return (d_2181_i9_) + ((self).f14)

            init57_ = lambda96_
            nw345_ = _dafny.Array(None, 22)
            for i0_57_ in range(nw345_.length(0)):
                nw345_[i0_57_] = init57_(i0_57_)
            d_2180_v73_ = nw345_
            index307_ = default__.safeIndex(268, (d_2180_v73_).length(0))
            (d_2180_v73_)[index307_] = (d_2070_v16_).cardinality
            d_2182_v74_: _dafny.Array
            nw346_ = _dafny.Array(None, 10)
            nw346_[int(0)] = self.f13
            nw346_[int(1)] = d_2163_cf17_
            nw346_[int(2)] = default__.fm28(self.f13, (0) - (d_2164_cf16_), 450, globalState)
            nw346_[int(3)] = self.f13
            nw346_[int(4)] = d_2163_cf17_
            nw346_[int(5)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_2183_i10_ in range(default__.abs(597))])
            nw346_[int(6)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2184_i11_ in range(default__.abs(953))])).set(default__.safeIndex(default__.fm55((self).f14, _dafny.SeqWithoutIsStrInference([d_2165_cf15_]), (0) - (d_2164_cf16_), globalState), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2185_i11_ in range(default__.abs(953))]))), d_2069_v15_)
            nw346_[int(7)] = d_2163_cf17_
            nw346_[int(8)] = self.f13
            nw346_[int(9)] = self.f13
            d_2182_v74_ = nw346_
            index308_ = default__.safeIndex(152, (d_2182_v74_).length(0))
            (d_2182_v74_)[index308_] = d_2163_cf17_
            index309_ = default__.safeIndex(268, (d_2180_v73_).length(0))
            index310_ = default__.safeIndex(152, (d_2182_v74_).length(0))
            rhs286_ = d_2164_cf16_
            rhs287_ = self.f13
            rhs288_ = ((d_2174_v68_)[default__.safeIndex(435, (d_2174_v68_).length(0))]) == (not(d_2051_v0_))
            rhs289_ = _dafny.Set({(0) - (len(d_2166_v63_))})
            lhs158_ = d_2180_v73_
            lhs159_ = default__.safeIndex(268, (d_2180_v73_).length(0))
            lhs160_ = d_2182_v74_
            lhs161_ = default__.safeIndex(152, (d_2182_v74_).length(0))
            lhs158_[lhs159_] = rhs286_
            lhs160_[lhs161_] = rhs287_
            r1 = rhs288_
            d_2177_v70_ = rhs289_
            r0 = -758
        elif True:
            d_2186___mcc_h6_ = source24_.cf11
            d_2187_cf11_ = d_2186___mcc_h6_
            d_2188_v75_: _dafny.Map
            d_2188_v75_ = _dafny.Map({not(d_2051_v0_): d_2187_cf11_})
            d_2189_v76_: _dafny.Seq
            d_2189_v76_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f14, len(self.f13)])])
            rhs290_ = d_2188_v75_
            rhs291_ = (d_2189_v76_) + ((d_2189_v76_) + (d_2189_v76_))
            rhs292_ = (self).f14
            d_2188_v75_ = rhs290_
            d_2189_v76_ = rhs291_
            r0 = rhs292_
            d_2190_v77_: _dafny.Array
            nw347_ = _dafny.Array(False, 5)
            d_2190_v77_ = nw347_
            d_2191_v78_: _dafny.Map
            d_2191_v78_ = _dafny.Map({len(self.f13): d_2190_v77_})
            d_2190_v77_ = ((d_2191_v78_)[(self).f14] if ((self).f14) in (d_2191_v78_) else d_2190_v77_)
            d_2192_v79_: _dafny.Seq
            d_2192_v79_ = _dafny.SeqWithoutIsStrInference([67])
            d_2193_v80_: _dafny.Map
            d_2193_v80_ = _dafny.Map({d_2051_v0_: d_2192_v79_})
            index311_ = default__.safeIndex(475, (d_2190_v77_).length(0))
            (d_2190_v77_)[index311_] = (d_2051_v0_) not in ((d_2193_v80_).set(d_2051_v0_, d_2192_v79_))
            index312_ = default__.safeIndex(475, (d_2190_v77_).length(0))
            (d_2190_v77_)[index312_] = (((self).f14 if d_2051_v0_ else (self).f14)) < ((self).f14)
            if ((self).f14) != ((self).f14):
                d_2194_v81_: C1
                nw348_ = C1()
                nw348_.ctor__(((self).f14 if d_2051_v0_ else (self).f14), (d_2190_v77_)[default__.safeIndex(475, (d_2190_v77_).length(0))], False, (0) - ((self).f14))
                d_2194_v81_ = nw348_
                d_2195_v82_: _dafny.Array
                nw349_ = _dafny.Array(_dafny.MultiSet({}), 15)
                d_2195_v82_ = nw349_
                d_2196_v83_: _dafny.MultiSet
                d_2196_v83_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qymaacj")), self.f13])
                index313_ = default__.safeIndex(846, (d_2195_v82_).length(0))
                (d_2195_v82_)[index313_] = d_2196_v83_
                d_2197_v84_: _dafny.Seq
                d_2197_v84_ = _dafny.SeqWithoutIsStrInference([(self.f13) + (self.f13), self.f13, self.f13])
                index314_ = default__.safeIndex(846, (d_2195_v82_).length(0))
                (d_2195_v82_)[index314_] = _dafny.MultiSet(d_2197_v84_)
                d_2198_v85_: _dafny.Map
                d_2198_v85_ = _dafny.Map({(self).f14: (self).f14})
                d_2199_v86_: _dafny.Set
                d_2199_v86_ = _dafny.Set({d_2198_v85_})
                d_2200_v87_: _dafny.Set
                d_2200_v87_ = _dafny.Set({_dafny.Map({(0) - ((self).f14): (self).f14})})
                d_2199_v86_ = d_2200_v87_
                index315_ = default__.safeIndex(475, (d_2190_v77_).length(0))
                (d_2190_v77_)[index315_] = True
                d_2201_v88_: _dafny.Map
                d_2201_v88_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f14]): (self).f14})
                d_2201_v88_ = (d_2201_v88_).set(d_2192_v79_, (self).f14)
            elif True:
                d_2202_v89_: _dafny.Map
                d_2202_v89_ = _dafny.Map({_dafny.Map({d_2051_v0_: (self).f14}): (self).f14})
                d_2203_v90_: D6
                d_2203_v90_ = D6_DC16(d_2202_v89_, (self).f14)
                d_2204_v91_: _dafny.Seq
                d_2204_v91_ = _dafny.SeqWithoutIsStrInference([d_2203_v90_, d_2203_v90_, d_2203_v90_])
                r2 = (d_2203_v90_) in ((d_2204_v91_) + (d_2204_v91_))
                d_2205_v92_: C3
                nw350_ = C3()
                nw350_.ctor__(d_2069_v15_, (self).f14, (d_2190_v77_)[default__.safeIndex(475, (d_2190_v77_).length(0))], True, 238)
                d_2205_v92_ = nw350_
                d_2206_v93_: D16
                d_2206_v93_ = D16_DC42((d_2190_v77_)[default__.safeIndex(475, (d_2190_v77_).length(0))], not (d_2051_v0_) or (d_2051_v0_))
                d_2206_v93_ = d_2206_v93_
                d_2207_v94_: _dafny.MultiSet
                d_2207_v94_ = _dafny.MultiSet([(d_2205_v92_).f25])
                index316_ = default__.safeIndex(475, (d_2190_v77_).length(0))
                (d_2190_v77_)[index316_] = (_dafny.MultiSet([d_2069_v15_])).isdisjoint(d_2207_v94_)
                d_2208_v95_: _dafny.Set
                d_2208_v95_ = _dafny.Set({(d_2205_v92_).f25})
                nw351_ = C3()
                nw351_.ctor__((d_2205_v92_).f25, (self).f14, (d_2190_v77_)[default__.safeIndex(475, (d_2190_v77_).length(0))], (d_2208_v95_).issubset(d_2208_v95_), ((self).f14) - ((self).f14))
                d_2205_v92_ = nw351_
        d_2209_v96_: _dafny.Array
        nw352_ = _dafny.Array(False, 24)
        d_2209_v96_ = nw352_
        d_2210_v97_: _dafny.Map
        d_2210_v97_ = _dafny.Map({self.f13: d_2209_v96_})
        if (d_2210_v97_) != (d_2210_v97_):
            d_2211_v98_: D7
            d_2211_v98_ = D7_DC19((self).f14)
            r2 = (default__.fm73(d_2211_v98_, globalState)).cf74
            d_2212_v99_: D0
            d_2212_v99_ = D0_DC0(default__.fm27(d_2069_v15_, globalState), 663, d_2051_v0_, (self).f14, d_2051_v0_)
            pat_let_tv27_ = d_2051_v0_
            d_2213_v100_: _dafny.Map
            def iife179_(_pat_let50_0):
                def iife180_(d_2214_dt__update__tmp_h2_):
                    def iife181_(_pat_let51_0):
                        def iife182_(d_2215_dt__update_hcf4_h0_):
                            return D0_DC0((d_2214_dt__update__tmp_h2_).cf0, (d_2214_dt__update__tmp_h2_).cf1, (d_2214_dt__update__tmp_h2_).cf2, (d_2214_dt__update__tmp_h2_).cf3, d_2215_dt__update_hcf4_h0_)
                        return iife182_(_pat_let51_0)
                    return iife181_(pat_let_tv27_)
                return iife180_(_pat_let50_0)
            d_2213_v100_ = _dafny.Map({(default__.fm32(D0_DC0(d_2051_v0_, (iife179_(d_2212_v99_)).cf1, d_2051_v0_, (self).f14, default__.fm27(d_2069_v15_, globalState)), globalState) if d_2051_v0_ else 885): (self).f14})
            d_2213_v100_ = (d_2213_v100_) | (d_2213_v100_)
            d_2216_v101_: _dafny.Array
            nw353_ = _dafny.Array(D4.default()(), 14)
            d_2216_v101_ = nw353_
            d_2217_v102_: _dafny.Map
            d_2217_v102_ = _dafny.Map({d_2051_v0_: d_2216_v101_})
            rhs293_ = (not(d_2051_v0_) if d_2051_v0_ else d_2051_v0_)
            rhs294_ = (self).f14
            rhs295_ = default__.safeModuloInt((0) - (((self).f14) * (len(d_2217_v102_))), (self).f14)
            rhs296_ = (self).f14
            d_2051_v0_ = rhs293_
            r0 = rhs294_
            r0 = rhs295_
            r0 = rhs296_
            d_2218_v103_: D4
            d_2218_v103_ = D4_DC8(d_2209_v96_)
            index317_ = default__.safeIndex(392, (d_2216_v101_).length(0))
            (d_2216_v101_)[index317_] = d_2218_v103_
            index318_ = default__.safeIndex(392, (d_2216_v101_).length(0))
            (d_2216_v101_)[index318_] = (d_2218_v103_ if ((self).f14) > ((0) - (len(_dafny.Map({d_2051_v0_: d_2051_v0_})))) else d_2218_v103_)
            r0 = (self).f14
        elif True:
            index319_ = default__.safeIndex(495, (d_2209_v96_).length(0))
            (d_2209_v96_)[index319_] = d_2051_v0_
            d_2219_v104_: _dafny.Seq
            d_2219_v104_ = _dafny.SeqWithoutIsStrInference([(len(_dafny.Map({d_2051_v0_: (self).f14}))) > ((0) - ((self).f14))])
            d_2220_v105_: _dafny.Array
            nw354_ = _dafny.Array(int(0), 4)
            d_2220_v105_ = nw354_
            index320_ = default__.safeIndex(93, (d_2220_v105_).length(0))
            (d_2220_v105_)[index320_] = (0) - ((self).f14)
            d_2221_v106_: _dafny.Map
            d_2221_v106_ = _dafny.Map({len(d_2219_v104_): (self).f14})
            d_2222_v107_: _dafny.Map
            d_2222_v107_ = _dafny.Map({176: d_2051_v0_})
            index321_ = default__.safeIndex(495, (d_2209_v96_).length(0))
            index322_ = default__.safeIndex(93, (d_2220_v105_).length(0))
            rhs297_ = d_2051_v0_
            rhs298_ = len(d_2221_v106_)
            rhs299_ = (d_2051_v0_) or (default__.fm27(d_2069_v15_, globalState))
            rhs300_ = _dafny.SeqWithoutIsStrInference([((d_2222_v107_)[(self).f14] if ((self).f14) in (d_2222_v107_) else d_2051_v0_)])
            rhs301_ = (self).f14
            lhs162_ = d_2209_v96_
            lhs163_ = default__.safeIndex(495, (d_2209_v96_).length(0))
            lhs164_ = d_2220_v105_
            lhs165_ = default__.safeIndex(93, (d_2220_v105_).length(0))
            lhs162_[lhs163_] = rhs297_
            r0 = rhs298_
            r2 = rhs299_
            d_2219_v104_ = rhs300_
            lhs164_[lhs165_] = rhs301_
            d_2223_v108_: _dafny.MultiSet
            d_2223_v108_ = _dafny.MultiSet([(d_2220_v105_)[default__.safeIndex(93, (d_2220_v105_).length(0))], (0) - ((self).f14)])
            d_2224_v109_: _dafny.Seq
            d_2224_v109_ = _dafny.SeqWithoutIsStrInference([(d_2220_v105_)[default__.safeIndex(93, (d_2220_v105_).length(0))]])
            index323_ = default__.safeIndex(93, (d_2220_v105_).length(0))
            (d_2220_v105_)[index323_] = (((d_2223_v108_).set((d_2220_v105_)[default__.safeIndex(93, (d_2220_v105_).length(0))], default__.abs(len(d_2224_v109_)))).set(-259, default__.abs((d_2220_v105_)[default__.safeIndex(93, (d_2220_v105_).length(0))]))).cardinality
            d_2224_v109_ = (d_2224_v109_) + ((d_2224_v109_) + (d_2224_v109_))
            d_2051_v0_ = default__.fm27(d_2069_v15_, globalState)
            index324_ = default__.safeIndex(495, (d_2209_v96_).length(0))
            rhs302_ = d_2220_v105_
            rhs303_ = (d_2219_v104_)[default__.safeIndex(len(default__.fm53(globalState)), len(d_2219_v104_))]
            lhs166_ = d_2209_v96_
            lhs167_ = default__.safeIndex(495, (d_2209_v96_).length(0))
            d_2220_v105_ = rhs302_
            lhs166_[lhs167_] = rhs303_
        d_2225_v110_: _dafny.Map
        d_2225_v110_ = _dafny.Map({self.f13: 338})
        d_2226_v111_: _dafny.Map
        d_2226_v111_ = _dafny.Map({len(d_2225_v110_): len(self.f13)})
        d_2227_v112_: _dafny.Map
        d_2227_v112_ = d_2226_v111_
        def lambda97_(source27_):
            d_2228___mcc_h28_ = source27_
            d_2229_cf75_ = d_2228___mcc_h28_
            return False

        if lambda97_(d_2227_v112_):
            d_2230_v113_: _dafny.Set
            d_2230_v113_ = _dafny.Set({False})
            d_2231_v114_: _dafny.Seq
            d_2231_v114_ = _dafny.SeqWithoutIsStrInference([(self).f14, (0) - (len(d_2230_v113_))])
            d_2232_v115_: _dafny.Set
            d_2232_v115_ = _dafny.Set({_dafny.Map({(d_2231_v114_)[default__.safeIndex((self).f14, len(d_2231_v114_))]: (self).f14}), _dafny.Map({(self).f14: (self).f14})})
            d_2233_v116_: _dafny.Seq
            d_2233_v116_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_2226_v111_}), d_2232_v115_])
            d_2234_v117_: C0
            nw355_ = C0()
            nw355_.ctor__((d_2233_v116_)[default__.safeIndex(default__.fm49(self.f13, (D9_DC24((self).f14)).cf48, not(d_2051_v0_), False, globalState), len(d_2233_v116_))], (self).f14)
            d_2234_v117_ = nw355_
            d_2235_v118_: _dafny.Set
            d_2235_v118_ = _dafny.Set({d_2231_v114_})
            d_2236_v119_: _dafny.Seq
            d_2236_v119_ = _dafny.SeqWithoutIsStrInference([d_2231_v114_])
            d_2237_v121_: _dafny.Array
            def lambda98_(d_2238_v15_):
                def lambda99_(d_2239_i12_):
                    return _dafny.MultiSet([d_2238_v15_, d_2238_v15_])

                return lambda99_

            init58_ = lambda98_(d_2069_v15_)
            nw356_ = _dafny.Array(None, 3)
            for i0_58_ in range(nw356_.length(0)):
                nw356_[i0_58_] = init58_(i0_58_)
            d_2237_v121_ = nw356_
            d_2240_v122_: _dafny.Map
            def iife183_():
                coll79_ = _dafny.Set()
                compr_79_: _dafny.Seq
                for compr_79_ in (d_2236_v119_).Elements:
                    d_2241_v120_: _dafny.Seq = compr_79_
                    if (d_2241_v120_) in (d_2236_v119_):
                        coll79_ = coll79_.union(_dafny.Set([d_2241_v120_]))
                return _dafny.Set(coll79_)
            d_2240_v122_ = _dafny.Map({(d_2235_v118_) | (iife183_()
            ): d_2237_v121_})
            d_2240_v122_ = (d_2240_v122_).set(default__.fm74((self).f14, globalState), d_2237_v121_)
            (self).f13 = (self.f13) + (self.f13)
            if d_2051_v0_:
                d_2242_v123_: _dafny.Map
                d_2242_v123_ = _dafny.Map({(d_2051_v0_) and (d_2051_v0_): d_2051_v0_})
                d_2242_v123_ = d_2242_v123_
                d_2243_v124_: _dafny.Map
                d_2243_v124_ = _dafny.Map({(self).fm17(_dafny.SeqWithoutIsStrInference([d_2051_v0_]), (self).f14, not(d_2051_v0_), (self).f14, globalState): (0) - ((d_2234_v117_).f27)})
                d_2244_v125_: _dafny.Seq
                d_2244_v125_ = _dafny.SeqWithoutIsStrInference([d_2051_v0_, d_2051_v0_, d_2051_v0_])
                d_2245_v126_: C6
                nw357_ = C6()
                nw357_.ctor__((948 if d_2051_v0_ else (d_2234_v117_).f27), d_2051_v0_, d_2051_v0_, (self).f14, ((d_2243_v124_)[d_2051_v0_] if (d_2051_v0_) in (d_2243_v124_) else len(d_2244_v125_)), d_2051_v0_)
                d_2245_v126_ = nw357_
                d_2246_v127_: _dafny.Map
                d_2246_v127_ = _dafny.Map({d_2243_v124_: d_2245_v126_.f21})
                d_2247_v128_: D6
                d_2247_v128_ = D6_DC16(d_2246_v127_, (d_2234_v117_).f27)
                d_2248_v129_: _dafny.Set
                d_2248_v129_ = _dafny.Set({(self).f14})
                d_2249_v130_: _dafny.Map
                d_2249_v130_ = _dafny.Map({d_2248_v129_: _dafny.Map({(d_2234_v117_).f27: (self).f14})})
                pat_let_tv28_ = d_2249_v130_
                pat_let_tv29_ = d_2248_v129_
                pat_let_tv30_ = d_2248_v129_
                pat_let_tv31_ = d_2249_v130_
                pat_let_tv32_ = d_2226_v111_
                d_2250_v131_: _dafny.Seq
                def iife184_(_pat_let52_0):
                    def iife185_(d_2251_dt__update__tmp_h3_):
                        def iife186_(_pat_let53_0):
                            def iife187_(d_2252_dt__update_hcf38_h0_):
                                return D6_DC16((d_2251_dt__update__tmp_h3_).cf37, d_2252_dt__update_hcf38_h0_)
                            return iife187_(_pat_let53_0)
                        return iife186_(len(((pat_let_tv28_)[pat_let_tv29_] if (pat_let_tv30_) in (pat_let_tv31_) else pat_let_tv32_)))
                    return iife185_(_pat_let52_0)
                d_2250_v131_ = _dafny.SeqWithoutIsStrInference([iife184_(d_2247_v128_)])
                index325_ = default__.safeIndex(187, (d_2209_v96_).length(0))
                (d_2209_v96_)[index325_] = (d_2245_v126_).fm0(globalState)
                d_2253_v132_: _dafny.Seq
                d_2253_v132_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2069_v15_ for d_2254_i14_ in range(default__.abs(196))])])
                index326_ = default__.safeIndex(187, (d_2209_v96_).length(0))
                rhs304_ = self.f13
                rhs305_ = _dafny.SeqWithoutIsStrInference([(0) - (((_dafny.MultiSet([(d_2234_v117_).f27])).cardinality) + ((d_2234_v117_).f27)) for d_2255_i13_ in range(default__.abs(-1))])
                rhs306_ = d_2250_v131_
                rhs307_ = d_2070_v16_
                rhs308_ = (((self.f13) + (self.f13)).set(default__.safeIndex(602, len((self.f13) + (self.f13))), d_2069_v15_)) != ((d_2253_v132_)[default__.safeIndex((d_2234_v117_).f27, len(d_2253_v132_))])
                lhs168_ = self
                lhs169_ = d_2209_v96_
                lhs170_ = default__.safeIndex(187, (d_2209_v96_).length(0))
                lhs168_.f13 = rhs304_
                r3 = rhs305_
                d_2250_v131_ = rhs306_
                d_2070_v16_ = rhs307_
                lhs169_[lhs170_] = rhs308_
                (self).f13 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lb"))
                index327_ = default__.safeIndex(187, (d_2209_v96_).length(0))
                (d_2209_v96_)[index327_] = not ((len((self.f13).set(default__.safeIndex(-847, len(self.f13)), d_2069_v15_))) >= ((self).f14)) or ((d_2209_v96_)[default__.safeIndex(187, (d_2209_v96_).length(0))])
            elif True:
                d_2256_v133_: _dafny.Set
                d_2256_v133_ = _dafny.Set({len(d_2230_v113_), (d_2234_v117_).f27})
                d_2257_v134_: _dafny.Seq
                d_2257_v134_ = _dafny.SeqWithoutIsStrInference([d_2051_v0_])
                d_2258_v135_: _dafny.Seq
                d_2258_v135_ = _dafny.SeqWithoutIsStrInference([d_2257_v134_, d_2257_v134_])
                d_2259_v136_: _dafny.Map
                d_2259_v136_ = _dafny.Map({(self).f14: d_2051_v0_})
                d_2260_v137_: _dafny.Array
                nw358_ = _dafny.Array(None, 22)
                nw358_[int(0)] = 114
                nw358_[int(1)] = (d_2234_v117_).f27
                nw358_[int(2)] = (d_2234_v117_).f27
                nw358_[int(3)] = (d_2234_v117_).f27
                nw358_[int(4)] = default__.safeModuloInt((d_2234_v117_).f27, (d_2231_v114_)[default__.safeIndex((self).f14, len(d_2231_v114_))])
                nw358_[int(5)] = (self).f14
                nw358_[int(6)] = len(d_2256_v133_)
                nw358_[int(7)] = (self).f14
                nw358_[int(8)] = default__.safeDivisionInt((self).f14, len(self.f13))
                nw358_[int(9)] = len((d_2258_v135_)[default__.safeIndex(len(_dafny.Map({d_2069_v15_: ((d_2259_v136_)[(self).f14] if ((self).f14) in (d_2259_v136_) else False)})), len(d_2258_v135_))])
                nw358_[int(10)] = len((self.f13) + (self.f13))
                nw358_[int(11)] = (d_2234_v117_).f27
                nw358_[int(12)] = ((d_2234_v117_).f27 if d_2051_v0_ else (d_2234_v117_).f27)
                nw358_[int(13)] = len((self.f13) + (self.f13))
                nw358_[int(14)] = (d_2234_v117_).f27
                nw358_[int(15)] = (self).f14
                nw358_[int(16)] = default__.safeDivisionInt((d_2234_v117_).f27, (d_2234_v117_).f27)
                nw358_[int(17)] = (0) - (((self).f14) - ((d_2234_v117_).f27))
                nw358_[int(18)] = (self).f14
                nw358_[int(19)] = (0) - ((d_2234_v117_).f27)
                nw358_[int(20)] = (self).f14
                nw358_[int(21)] = len(self.f13)
                d_2260_v137_ = nw358_
                d_2260_v137_ = d_2260_v137_
                r1 = d_2051_v0_
                d_2261_v138_: C1
                nw359_ = C1()
                nw359_.ctor__((d_2234_v117_).f27, d_2051_v0_, d_2051_v0_, 564)
                d_2261_v138_ = nw359_
                r2 = d_2051_v0_
                d_2262_v139_: _dafny.Array
                nw360_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_2262_v139_ = nw360_
                index328_ = default__.safeIndex(503, (d_2262_v139_).length(0))
                (d_2262_v139_)[index328_] = d_2209_v96_
                index329_ = default__.safeIndex(503, (d_2262_v139_).length(0))
                (d_2262_v139_)[index329_] = d_2209_v96_
            r1 = d_2051_v0_
        elif True:
            r0 = ((self).f14) + ((self).f14)
            d_2263_v140_: D5
            d_2263_v140_ = D5_DC13(False, d_2051_v0_)
            source28_ = d_2263_v140_
            if source28_.is_DC13:
                d_2264___mcc_h20_ = source28_.cf29
                d_2265___mcc_h21_ = source28_.cf30
                d_2266_cf30_ = d_2265___mcc_h21_
                d_2267_cf29_ = d_2264___mcc_h20_
                r2 = d_2051_v0_
                d_2268_v141_: _dafny.Seq
                d_2268_v141_ = _dafny.SeqWithoutIsStrInference([(self).f14])
                r0 = (d_2268_v141_)[default__.safeIndex((self).f14, len(d_2268_v141_))]
                d_2269_v142_: _dafny.MultiSet
                d_2269_v142_ = _dafny.MultiSet([((d_2070_v16_)[d_2267_cf29_] if (d_2267_cf29_) in (d_2070_v16_) else (self).f14), -145])
                d_2270_v143_: C4
                nw361_ = C4()
                nw361_.ctor__((self).f14, not((d_2269_v142_) != (_dafny.MultiSet(d_2268_v141_))), self.f13, d_2267_cf29_, True, (self).f14, (self).f14, ((self).f14) == ((self).f14))
                d_2270_v143_ = nw361_
                d_2271_v144_: _dafny.Map
                d_2271_v144_ = _dafny.Map({d_2069_v15_: (_dafny.MultiSet([d_2266_cf30_])).cardinality})
                d_2272_v145_: _dafny.Set
                d_2272_v145_ = _dafny.Set({(self).f14, len(d_2271_v144_), (self).f14})
                d_2273_v146_: D17
                d_2273_v146_ = D17_DC47((d_2272_v145_).ispropersubset(d_2272_v145_), d_2051_v0_, default__.safeDivisionInt((d_2270_v143_).f23, (d_2270_v143_).f23), 950, default__.safeDivisionInt((d_2270_v143_).f23, (default__.fm47(True, globalState)).cardinality))
                d_2274_v147_: D16
                d_2274_v147_ = D16_DC41(d_2069_v15_, d_2051_v0_)
                d_2273_v146_ = default__.fm75((d_2274_v147_).cf78, globalState)
            elif source28_.is_DC14:
                d_2275___mcc_h22_ = source28_.cf31
                d_2276___mcc_h23_ = source28_.cf32
                d_2277___mcc_h24_ = source28_.cf33
                d_2278___mcc_h25_ = source28_.cf34
                d_2279___mcc_h26_ = source28_.cf35
                d_2280_cf35_ = d_2279___mcc_h26_
                d_2281_cf34_ = d_2278___mcc_h25_
                d_2282_cf33_ = d_2277___mcc_h24_
                d_2283_cf32_ = d_2276___mcc_h23_
                d_2284_cf31_ = d_2275___mcc_h22_
                d_2285_v148_: _dafny.Array
                def lambda100_(d_2286_i15_):
                    return _dafny.Set({-421})

                init59_ = lambda100_
                nw362_ = _dafny.Array(None, 16)
                for i0_59_ in range(nw362_.length(0)):
                    nw362_[i0_59_] = init59_(i0_59_)
                d_2285_v148_ = nw362_
                index330_ = default__.safeIndex(706, (d_2285_v148_).length(0))
                (d_2285_v148_)[index330_] = default__.fm76(d_2069_v15_, d_2069_v15_, d_2051_v0_, globalState)
                d_2287_v149_: _dafny.Set
                d_2287_v149_ = _dafny.Set({(self).f14, (0) - (default__.safeModuloInt(d_2281_cf34_, -642)), (self).f14})
                index331_ = default__.safeIndex(706, (d_2285_v148_).length(0))
                rhs309_ = (0) - ((d_2280_cf35_) + (d_2280_cf35_))
                rhs310_ = d_2287_v149_
                lhs171_ = d_2285_v148_
                lhs172_ = default__.safeIndex(706, (d_2285_v148_).length(0))
                d_2280_cf35_ = rhs309_
                lhs171_[lhs172_] = rhs310_
                d_2281_cf34_ = d_2280_cf35_
                d_2069_v15_ = d_2069_v15_
                d_2280_cf35_ = (self).f14
            elif True:
                d_2288___mcc_h27_ = source28_.cf28
                d_2289_cf28_ = d_2288___mcc_h27_
                d_2051_v0_ = d_2051_v0_
                r2 = d_2051_v0_
                d_2290_v150_: _dafny.Map
                d_2290_v150_ = _dafny.Map({(self).f14: (d_2226_v111_).set((self).f14, (self).f14)})
                d_2291_v155_: _dafny.Array
                nw363_ = _dafny.Array(None, 16)
                nw363_[int(0)] = d_2290_v150_
                nw363_[int(1)] = d_2290_v150_
                nw363_[int(2)] = default__.fm77(globalState)
                def iife188_():
                    coll80_ = _dafny.Map()
                    compr_80_: int
                    for compr_80_ in _dafny.IntegerRange(188, 987):
                        d_2292_v151_: int = compr_80_
                        if ((188) <= (d_2292_v151_)) and ((d_2292_v151_) < (987)):
                            coll80_[(d_2292_v151_) - ((self).f14)] = d_2226_v111_
                    return _dafny.Map(coll80_)
                nw363_[int(3)] = iife188_()
                
                def iife189_():
                    coll81_ = _dafny.Map()
                    compr_81_: int
                    for compr_81_ in (((_dafny.SeqWithoutIsStrInference([498, (self).f14])).set(default__.safeIndex((self).f14, len(_dafny.SeqWithoutIsStrInference([498, (self).f14]))), (self).f14)).set(default__.safeIndex((self).f14, len((_dafny.SeqWithoutIsStrInference([498, (self).f14])).set(default__.safeIndex((self).f14, len(_dafny.SeqWithoutIsStrInference([498, (self).f14]))), (self).f14))), (self).f14)).Elements:
                        d_2293_v152_: int = compr_81_
                        if (d_2293_v152_) in (((_dafny.SeqWithoutIsStrInference([498, (self).f14])).set(default__.safeIndex((self).f14, len(_dafny.SeqWithoutIsStrInference([498, (self).f14]))), (self).f14)).set(default__.safeIndex((self).f14, len((_dafny.SeqWithoutIsStrInference([498, (self).f14])).set(default__.safeIndex((self).f14, len(_dafny.SeqWithoutIsStrInference([498, (self).f14]))), (self).f14))), (self).f14)):
                            coll81_[default__.safeDivisionInt(d_2293_v152_, 752)] = (d_2226_v111_)
                    return _dafny.Map(coll81_)
                nw363_[int(4)] = iife189_()
                
                def iife190_():
                    coll82_ = _dafny.Map()
                    compr_82_: int
                    for compr_82_ in _dafny.IntegerRange(796, 986):
                        d_2294_v153_: int = compr_82_
                        if ((796) <= (d_2294_v153_)) and ((d_2294_v153_) < (986)):
                            coll82_[default__.safeDivisionInt(d_2294_v153_, len(_dafny.SeqWithoutIsStrInference([(self).f14])))] = d_2226_v111_
                    return _dafny.Map(coll82_)
                nw363_[int(5)] = iife190_()
                
                nw363_[int(6)] = (d_2290_v150_).set((self).f14, d_2226_v111_)
                nw363_[int(7)] = _dafny.Map({(self).f14: d_2226_v111_})
                nw363_[int(8)] = (d_2290_v150_) | (d_2290_v150_)
                nw363_[int(9)] = (d_2290_v150_).set((self).f14, _dafny.Map({(self).f14: (self).f14}))
                nw363_[int(10)] = (d_2290_v150_) | (d_2290_v150_)
                nw363_[int(11)] = d_2290_v150_
                nw363_[int(12)] = d_2290_v150_
                nw363_[int(13)] = d_2290_v150_
                def iife191_():
                    coll83_ = _dafny.Map()
                    compr_83_: int
                    for compr_83_ in _dafny.IntegerRange(747, 880):
                        d_2295_v154_: int = compr_83_
                        if ((747) <= (d_2295_v154_)) and ((d_2295_v154_) < (880)):
                            coll83_[(d_2295_v154_) + ((self).f14)] = d_2226_v111_
                    return _dafny.Map(coll83_)
                nw363_[int(14)] = (_dafny.Map({len(d_2226_v111_): d_2226_v111_})) | (iife191_()
                )
                nw363_[int(15)] = (d_2290_v150_).set((self).f14, d_2226_v111_)
                d_2291_v155_ = nw363_
                index332_ = default__.safeIndex(432, (d_2291_v155_).length(0))
                (d_2291_v155_)[index332_] = (d_2290_v150_) | (d_2290_v150_)
                d_2296_v156_: _dafny.Map
                d_2296_v156_ = _dafny.Map({(self).f14: d_2051_v0_})
                index333_ = default__.safeIndex(432, (d_2291_v155_).length(0))
                (d_2291_v155_)[index333_] = (d_2290_v150_) | (_dafny.Map({len(d_2296_v156_): d_2226_v111_}))
                d_2297_v157_: _dafny.Map
                d_2297_v157_ = _dafny.Map({default__.fm27(d_2069_v15_, globalState): d_2051_v0_})
                d_2298_v158_: _dafny.Array
                nw364_ = _dafny.Array(None, 4)
                nw364_[int(0)] = d_2297_v157_
                nw364_[int(1)] = d_2297_v157_
                nw364_[int(2)] = d_2297_v157_
                nw364_[int(3)] = default__.fm46(globalState)
                d_2298_v158_ = nw364_
                index334_ = default__.safeIndex(831, (d_2298_v158_).length(0))
                (d_2298_v158_)[index334_] = d_2297_v157_
                d_2299_v159_: _dafny.Array
                nw365_ = _dafny.Array(None, 19)
                nw365_[int(0)] = self.f13
                nw365_[int(1)] = _dafny.SeqWithoutIsStrInference([d_2069_v15_ for d_2300_i16_ in range(default__.abs(970))])
                nw365_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))) + (self.f13)
                nw365_[int(3)] = _dafny.SeqWithoutIsStrInference([d_2069_v15_ for d_2301_i17_ in range(default__.abs(930))])
                nw365_[int(4)] = (_dafny.SeqWithoutIsStrInference([d_2069_v15_ for d_2302_i18_ in range(default__.abs(641))])) + (self.f13)
                nw365_[int(5)] = self.f13
                nw365_[int(6)] = self.f13
                nw365_[int(7)] = (self.f13).set(default__.safeIndex((self).f14, len(self.f13)), d_2069_v15_)
                nw365_[int(8)] = ((self.f13) + (self.f13)).set(default__.safeIndex((self).f14, len((self.f13) + (self.f13))), d_2069_v15_)
                nw365_[int(9)] = (self.f13) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bopjig")))
                nw365_[int(10)] = self.f13
                nw365_[int(11)] = self.f13
                nw365_[int(12)] = self.f13
                nw365_[int(13)] = self.f13
                nw365_[int(14)] = (self.f13 if d_2051_v0_ else self.f13)
                nw365_[int(15)] = (self.f13) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kaakuatca")))
                nw365_[int(16)] = self.f13
                nw365_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgrlndl"))
                nw365_[int(18)] = (self.f13) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rk")))
                d_2299_v159_ = nw365_
                d_2303_v160_: _dafny.Seq
                d_2303_v160_ = _dafny.SeqWithoutIsStrInference([d_2051_v0_])
                d_2304_v161_: _dafny.Array
                nw366_ = _dafny.Array(int(0), 2)
                d_2304_v161_ = nw366_
                d_2305_v162_: D5
                d_2305_v162_ = D5_DC14(d_2051_v0_, d_2303_v160_, d_2304_v161_, (self).f14, (self).f14)
                index335_ = default__.safeIndex(159, (d_2299_v159_).length(0))
                (d_2299_v159_)[index335_] = default__.fm52(not(d_2051_v0_), _dafny.Set({(self).f14}), (self).f14, (d_2305_v162_).cf31, globalState)
                index336_ = default__.safeIndex(831, (d_2298_v158_).length(0))
                index337_ = default__.safeIndex(159, (d_2299_v159_).length(0))
                rhs311_ = d_2069_v15_
                rhs312_ = d_2297_v157_
                rhs313_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xqabmlu"))
                lhs173_ = d_2298_v158_
                lhs174_ = default__.safeIndex(831, (d_2298_v158_).length(0))
                lhs175_ = d_2299_v159_
                lhs176_ = default__.safeIndex(159, (d_2299_v159_).length(0))
                d_2069_v15_ = rhs311_
                lhs173_[lhs174_] = rhs312_
                lhs175_[lhs176_] = rhs313_
            d_2306_v164_: _dafny.MultiSet
            d_2306_v164_ = _dafny.MultiSet([d_2069_v15_, d_2069_v15_])
            d_2307_v165_: C1
            nw367_ = C1()
            def iife192_():
                coll84_ = _dafny.Map()
                compr_84_: str
                for compr_84_ in (d_2306_v164_).Elements:
                    d_2308_v163_: str = compr_84_
                    if (d_2308_v163_) in (d_2306_v164_):
                        coll84_[d_2308_v163_] = False
                return _dafny.Map(coll84_)
            nw367_.ctor__(len(iife192_()
            ), d_2051_v0_, d_2051_v0_, (self).f14)
            d_2307_v165_ = nw367_
            d_2309_v166_: _dafny.Seq
            d_2309_v166_ = _dafny.SeqWithoutIsStrInference([d_2307_v165_])
            d_2310_v167_: _dafny.Array
            nw368_ = _dafny.Array(None, 14)
            nw368_[int(0)] = d_2307_v165_
            nw368_[int(1)] = d_2307_v165_
            nw368_[int(2)] = d_2307_v165_
            nw368_[int(3)] = d_2307_v165_
            nw368_[int(4)] = d_2307_v165_
            nw368_[int(5)] = d_2307_v165_
            nw368_[int(6)] = d_2307_v165_
            nw368_[int(7)] = d_2307_v165_
            nw368_[int(8)] = (d_2309_v166_)[default__.safeIndex((self).f14, len(d_2309_v166_))]
            nw368_[int(9)] = d_2307_v165_
            nw368_[int(10)] = d_2307_v165_
            nw368_[int(11)] = d_2307_v165_
            nw368_[int(12)] = d_2307_v165_
            nw368_[int(13)] = d_2307_v165_
            d_2310_v167_ = nw368_
            index338_ = default__.safeIndex(564, (d_2310_v167_).length(0))
            (d_2310_v167_)[index338_] = d_2307_v165_
            d_2311_v168_: _dafny.Map
            d_2311_v168_ = _dafny.Map({d_2051_v0_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bbybbory"))) + (self.f13)})
            d_2312_v169_: _dafny.Set
            d_2312_v169_ = _dafny.Set({(self).f14})
            index339_ = default__.safeIndex(564, (d_2310_v167_).length(0))
            rhs314_ = (self).f14
            rhs315_ = d_2051_v0_
            rhs316_ = d_2307_v165_
            rhs317_ = ((d_2311_v168_)[d_2051_v0_] if (d_2051_v0_) in (d_2311_v168_) else default__.fm52(d_2051_v0_, d_2312_v169_, (self).f14, d_2051_v0_, globalState))
            lhs177_ = d_2310_v167_
            lhs178_ = default__.safeIndex(564, (d_2310_v167_).length(0))
            lhs179_ = self
            r0 = rhs314_
            d_2051_v0_ = rhs315_
            lhs177_[lhs178_] = rhs316_
            lhs179_.f13 = rhs317_
            d_2313_v170_: D17
            d_2313_v170_ = D17_DC46()
            d_2314_v171_: _dafny.Map
            d_2314_v171_ = _dafny.Map({(d_2307_v165_).fm0(globalState): d_2313_v170_})
            d_2314_v171_ = (d_2314_v171_).set(d_2051_v0_, d_2313_v170_)
            d_2315_v172_: _dafny.Array
            nw369_ = _dafny.Array(_dafny.CodePoint('D'), 2)
            d_2315_v172_ = nw369_
            index340_ = default__.safeIndex(852, (d_2315_v172_).length(0))
            (d_2315_v172_)[index340_] = d_2069_v15_
            index341_ = default__.safeIndex(852, (d_2315_v172_).length(0))
            (d_2315_v172_)[index341_] = (self.f13)[default__.safeIndex(default__.safeDivisionInt((self).f14, (self).f14), len(self.f13))]
        d_2316_v173_: _dafny.Seq
        d_2316_v173_ = _dafny.SeqWithoutIsStrInference([d_2051_v0_])
        d_2226_v111_ = (d_2226_v111_).set(len(_dafny.Map({d_2051_v0_: d_2316_v173_})), 442)
        r0 = (self).f14
        r0 = (0) - (len((_dafny.Map({(self).f14: (self).f14})) | (default__.fm58(578, (0) - ((self).f14), globalState))))
        pat_let_tv33_ = d_2051_v0_
        def lambda101_(source29_):
            d_2317___mcc_h29_ = source29_
            d_2318_cf75_ = d_2317___mcc_h29_
            return pat_let_tv33_

        r1 = lambda101_(d_2227_v112_)
        pat_let_tv34_ = d_2051_v0_
        pat_let_tv35_ = d_2051_v0_
        pat_let_tv36_ = d_2316_v173_
        pat_let_tv37_ = d_2316_v173_
        def lambda102_(source30_):
            if source30_.is_DC41:
                d_2319___mcc_h30_ = source30_.cf77
                d_2320___mcc_h31_ = source30_.cf78
                d_2321_cf78_ = d_2320___mcc_h31_
                d_2322_cf77_ = d_2319___mcc_h30_
                return pat_let_tv34_
            elif source30_.is_DC42:
                d_2323___mcc_h32_ = source30_.cf79
                d_2324___mcc_h33_ = source30_.cf80
                d_2325_cf80_ = d_2324___mcc_h33_
                d_2326_cf79_ = d_2323___mcc_h32_
                return (_dafny.Set({_dafny.Map({(self).f14: d_2326_cf79_})})).ispropersubset(_dafny.Set({_dafny.Map({(self).f14: d_2325_cf80_})}))
            elif source30_.is_DC40:
                d_2327___mcc_h34_ = source30_.cf76
                d_2328_cf76_ = d_2327___mcc_h34_
                return pat_let_tv35_
            elif True:
                d_2329___mcc_h35_ = source30_.cf81
                d_2330_cf81_ = d_2329___mcc_h35_
                return (pat_let_tv36_)[default__.safeIndex((self).f14, len(pat_let_tv37_))]

        r2 = not(lambda102_(D16_DC42(True, d_2051_v0_)))
        d_2331_v174_: _dafny.Seq
        d_2331_v174_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f14)])
        r3 = (d_2331_v174_ if d_2051_v0_ else _dafny.SeqWithoutIsStrInference([(0) - ((self).f14), (self).f14]))
        return r0, r1, r2, r3

    @property
    def f14(self):
        return self._f14

class C9:
    def  __init__(self):
        self._f12: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    def ctor__(self, f12):
        (self)._f12 = f12

    def m8(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        d_2332_v1_: _dafny.Set
        d_2332_v1_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktgtetk"))})
        d_2333_v2_: D5
        d_2333_v2_ = D5_DC12(d_2332_v1_)
        d_2334_v3_: _dafny.Seq
        d_2334_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yblbc"))
        d_2335_i0_: int
        d_2335_i0_ = 0
        with _dafny.label("7"):
            def iife193_():
                coll85_ = _dafny.Set()
                compr_85_: _dafny.Seq
                for compr_85_ in (default__.fm16(globalState)).keys.Elements:
                    d_2339_v0_: _dafny.Seq = compr_85_
                    if (d_2339_v0_) in (default__.fm16(globalState)):
                        coll85_ = coll85_.union(_dafny.Set([d_2339_v0_]))
                return _dafny.Set(coll85_)
            while (((d_2333_v2_).cf28) | (_dafny.Set({d_2334_v3_}))).ispropersubset(iife193_()
            ):
                with _dafny.c_label("7"):
                    if (d_2335_i0_) >= (100):
                        raise _dafny.Break("7")
                    d_2335_i0_ = (d_2335_i0_) + (1)
                    d_2336_v4_: int
                    d_2336_v4_ = -620
                    d_2336_v4_ = d_2336_v4_
                    d_2337_v5_: str
                    d_2337_v5_ = _dafny.CodePoint('d')
                    d_2338_v6_: C3
                    nw370_ = C3()
                    nw370_.ctor__(d_2337_v5_, d_2336_v4_, False, True, d_2336_v4_)
                    d_2338_v6_ = nw370_
                    index342_ = default__.safeIndex(392, ((self).f12).length(0))
                    ((self).f12)[index342_] = p0
                    index343_ = default__.safeIndex(392, ((self).f12).length(0))
                    ((self).f12)[index343_] = p0
                    d_2336_v4_ = d_2336_v4_
                    pass
            pass
        d_2340_v7_: int
        d_2340_v7_ = 224
        d_2340_v7_ = d_2340_v7_
        d_2341_v8_: _dafny.Array
        nw371_ = _dafny.Array(_dafny.Array(None, 0), 17)
        d_2341_v8_ = nw371_
        index344_ = default__.safeIndex(557, (d_2341_v8_).length(0))
        (d_2341_v8_)[index344_] = (self).f12
        index345_ = default__.safeIndex(557, (d_2341_v8_).length(0))
        nw372_ = _dafny.Array(False, 12)
        (d_2341_v8_)[index345_] = nw372_
        d_2342_v9_: bool
        d_2342_v9_ = True
        d_2343_v10_: _dafny.Map
        d_2343_v10_ = _dafny.Map({d_2342_v9_: d_2340_v7_})
        d_2344_v11_: D4
        d_2344_v11_ = D4_DC9(d_2343_v10_, p2, d_2342_v9_)
        d_2342_v9_ = (d_2344_v11_).cf23
        d_2345_v12_: _dafny.Array
        nw373_ = _dafny.Array(int(0), 1)
        d_2345_v12_ = nw373_
        d_2346_v14_: _dafny.Map
        def iife194_():
            coll86_ = _dafny.Set()
            compr_86_: int
            for compr_86_ in _dafny.IntegerRange(410, 526):
                d_2347_v13_: int = compr_86_
                if ((410) <= (d_2347_v13_)) and ((d_2347_v13_) < (526)):
                    coll86_ = coll86_.union(_dafny.Set([(d_2347_v13_) + (p2)]))
            return _dafny.Set(coll86_)
        d_2346_v14_ = _dafny.Map({len(iife194_()
        ): _dafny.CodePoint('w')})
        d_2348_v15_: _dafny.Seq
        d_2348_v15_ = _dafny.SeqWithoutIsStrInference([-899, p2])
        d_2349_v16_: _dafny.Map
        d_2349_v16_ = _dafny.Map({d_2348_v15_: default__.fm27(_dafny.CodePoint('a'), globalState)})
        d_2350_v17_: _dafny.Set
        d_2350_v17_ = _dafny.Set({len(d_2334_v3_), (len(d_2348_v15_)) - (len(d_2334_v3_)), len(d_2349_v16_), p2, d_2340_v7_})
        rhs318_ = p2
        rhs319_ = d_2345_v12_
        rhs320_ = (d_2346_v14_) | (d_2346_v14_)
        rhs321_ = p1
        rhs322_ = d_2350_v17_
        d_2340_v7_ = rhs318_
        d_2345_v12_ = rhs319_
        d_2346_v14_ = rhs320_
        d_2342_v9_ = rhs321_
        d_2350_v17_ = rhs322_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_2345_v12_).length(0)):
            d_2351_i1_: int = guard_loop_7_
            if (True) and (((0) <= (d_2351_i1_)) and ((d_2351_i1_) < ((d_2345_v12_).length(0)))):
                (d_2345_v12_)[(d_2351_i1_)] = default__.safeModuloInt(d_2351_i1_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_2352_i2_ in range(default__.abs(754))])))
        r0 = d_2348_v15_
        return r0

    def m9(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        r3: int = int(0)
        d_2353_v0_: int
        d_2353_v0_ = 4
        hi7_ = (d_2353_v0_) * (d_2353_v0_)
        for d_2354_i0_ in range(d_2353_v0_, hi7_):
            d_2355_v1_: bool
            d_2355_v1_ = True
            r1 = not(d_2355_v1_)
            d_2356_v2_: _dafny.Seq
            d_2356_v2_ = _dafny.SeqWithoutIsStrInference([d_2355_v1_])
            d_2356_v2_ = d_2356_v2_
            d_2357_v3_: _dafny.Seq
            d_2357_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsc"))
            d_2358_v4_: _dafny.Set
            d_2358_v4_ = _dafny.Set({d_2355_v1_})
            d_2359_v5_: _dafny.Seq
            d_2359_v5_ = _dafny.SeqWithoutIsStrInference([d_2358_v4_, d_2358_v4_, d_2358_v4_, d_2358_v4_, d_2358_v4_])
            d_2360_v6_: _dafny.Map
            d_2360_v6_ = _dafny.Map({d_2354_i0_: d_2355_v1_})
            d_2361_v7_: _dafny.Set
            d_2361_v7_ = _dafny.Set({d_2360_v6_, (d_2360_v6_).set(d_2354_i0_, d_2355_v1_)})
            d_2362_v8_: _dafny.Array
            nw374_ = _dafny.Array(None, 13)
            nw374_[int(0)] = d_2354_i0_
            nw374_[int(1)] = d_2354_i0_
            nw374_[int(2)] = d_2354_i0_
            nw374_[int(3)] = d_2354_i0_
            nw374_[int(4)] = len(_dafny.SeqWithoutIsStrInference([d_2354_i0_]))
            nw374_[int(5)] = d_2354_i0_
            nw374_[int(6)] = d_2353_v0_
            nw374_[int(7)] = d_2353_v0_
            nw374_[int(8)] = d_2354_i0_
            nw374_[int(9)] = len((d_2359_v5_)[default__.safeIndex(d_2354_i0_, len(d_2359_v5_))])
            nw374_[int(10)] = len(d_2361_v7_)
            nw374_[int(11)] = d_2354_i0_
            nw374_[int(12)] = d_2353_v0_
            d_2362_v8_ = nw374_
            d_2363_v9_: D11
            d_2363_v9_ = D11_DC30(d_2355_v1_, d_2362_v8_, d_2354_i0_, d_2355_v1_)
            d_2364_v10_: _dafny.Set
            d_2364_v10_ = _dafny.Set({(d_2363_v9_).cf63, len(d_2357_v3_)})
            d_2357_v3_ = default__.fm52(d_2355_v1_, d_2364_v10_, len((d_2356_v2_) + (d_2356_v2_)), d_2355_v1_, globalState)
            d_2365_v11_: _dafny.Set
            d_2365_v11_ = _dafny.Set({(d_2357_v3_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2366_i1_ in range(default__.abs(119))]))})
            index346_ = default__.safeIndex(458, (d_2362_v8_).length(0))
            (d_2362_v8_)[index346_] = 580
            d_2367_v12_: str
            d_2367_v12_ = _dafny.CodePoint('e')
            d_2368_v13_: _dafny.Seq
            d_2368_v13_ = _dafny.SeqWithoutIsStrInference([d_2357_v3_])
            d_2369_v15_: _dafny.MultiSet
            d_2369_v15_ = _dafny.MultiSet([d_2355_v1_, d_2355_v1_, default__.fm27(_dafny.CodePoint('d'), globalState)])
            d_2370_v16_: _dafny.MultiSet
            d_2370_v16_ = _dafny.MultiSet([(0) - ((d_2369_v15_).cardinality), d_2353_v0_, default__.fm49(d_2357_v3_, d_2353_v0_, d_2355_v1_, d_2355_v1_, globalState), d_2353_v0_, d_2353_v0_])
            index347_ = default__.safeIndex(458, (d_2362_v8_).length(0))
            def iife195_():
                coll87_ = _dafny.Set()
                compr_87_: _dafny.Seq
                for compr_87_ in (d_2368_v13_).Elements:
                    d_2371_v14_: _dafny.Seq = compr_87_
                    if (d_2371_v14_) in (d_2368_v13_):
                        coll87_ = coll87_.union(_dafny.Set([d_2371_v14_]))
                return _dafny.Set(coll87_)
            def iife196_():
                coll88_ = _dafny.Set()
                compr_88_: int
                for compr_88_ in (d_2370_v16_).Elements:
                    d_2372_v17_: int = compr_88_
                    if (d_2372_v17_) in (d_2370_v16_):
                        coll88_ = coll88_.union(_dafny.Set([(d_2372_v17_) - (len((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([True]))))]))
                return _dafny.Set(coll88_)
            rhs323_ = iife195_()

            rhs324_ = (d_2353_v0_ if (False) and (d_2355_v1_) else d_2353_v0_)
            rhs325_ = len(iife196_()
            )
            rhs326_ = d_2367_v12_
            lhs180_ = d_2362_v8_
            lhs181_ = default__.safeIndex(458, (d_2362_v8_).length(0))
            d_2365_v11_ = rhs323_
            lhs180_[lhs181_] = rhs324_
            r3 = rhs325_
            d_2367_v12_ = rhs326_
        d_2373_v18_: bool
        d_2373_v18_ = True
        d_2374_v19_: _dafny.MultiSet
        d_2374_v19_ = _dafny.MultiSet([d_2373_v18_, False])
        d_2375_v20_: _dafny.Seq
        d_2375_v20_ = _dafny.SeqWithoutIsStrInference([not(d_2373_v18_), d_2373_v18_])
        d_2376_v21_: _dafny.Map
        d_2376_v21_ = _dafny.Map({d_2375_v20_: d_2373_v18_})
        d_2377_v22_: _dafny.Seq
        d_2377_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwpsx"))
        d_2378_v24_: _dafny.Seq
        d_2378_v24_ = _dafny.SeqWithoutIsStrInference([d_2375_v20_])
        d_2379_v25_: _dafny.Map
        d_2379_v25_ = _dafny.Map({True: not(d_2373_v18_)})
        d_2380_v26_: _dafny.Seq
        d_2380_v26_ = _dafny.SeqWithoutIsStrInference([d_2353_v0_, d_2353_v0_, d_2353_v0_, 796, -951])
        d_2381_v27_: _dafny.Array
        nw375_ = _dafny.Array(None, 20)
        nw375_[int(0)] = -641
        nw375_[int(1)] = ((d_2374_v19_).cardinality) * (d_2353_v0_)
        nw375_[int(2)] = (0) - (d_2353_v0_)
        nw375_[int(3)] = (d_2353_v0_) - (len(d_2376_v21_))
        nw375_[int(4)] = d_2353_v0_
        nw375_[int(5)] = len(d_2377_v22_)
        nw375_[int(6)] = d_2353_v0_
        nw375_[int(7)] = d_2353_v0_
        nw375_[int(8)] = d_2353_v0_
        nw375_[int(9)] = len(d_2377_v22_)
        def iife197_():
            coll89_ = _dafny.Map()
            compr_89_: _dafny.Seq
            for compr_89_ in (d_2378_v24_).Elements:
                d_2382_v23_: _dafny.Seq = compr_89_
                if (d_2382_v23_) in (d_2378_v24_):
                    coll89_[d_2382_v23_] = (0) - (d_2353_v0_)
            return _dafny.Map(coll89_)
        nw375_[int(10)] = len(iife197_()
        )
        nw375_[int(11)] = d_2353_v0_
        nw375_[int(12)] = d_2353_v0_
        nw375_[int(13)] = (D7_DC19(d_2353_v0_)).cf41
        nw375_[int(14)] = (d_2353_v0_) * (810)
        nw375_[int(15)] = d_2353_v0_
        nw375_[int(16)] = len((d_2379_v25_) | (d_2379_v25_))
        nw375_[int(17)] = (0) - (-465)
        nw375_[int(18)] = d_2353_v0_
        nw375_[int(19)] = ((0) - (d_2353_v0_) if True else len(d_2380_v26_))
        d_2381_v27_ = nw375_
        index348_ = default__.safeIndex(273, (d_2381_v27_).length(0))
        (d_2381_v27_)[index348_] = (d_2353_v0_ if d_2373_v18_ else d_2353_v0_)
        index349_ = default__.safeIndex(163, (p0).length(0))
        (p0)[index349_] = d_2377_v22_
        d_2383_v29_: D1
        d_2383_v29_ = D1_DC5(d_2373_v18_, len(d_2377_v22_), d_2377_v22_)
        d_2384_v30_: _dafny.Map
        def iife198_():
            coll90_ = _dafny.Map()
            compr_90_: int
            for compr_90_ in _dafny.IntegerRange(567, 589):
                d_2385_v28_: int = compr_90_
                if ((567) <= (d_2385_v28_)) and ((d_2385_v28_) < (589)):
                    coll90_[(d_2385_v28_) * (386)] = d_2353_v0_
            return _dafny.Map(coll90_)
        d_2384_v30_ = _dafny.Map({iife198_()
        : len(((d_2383_v29_).cf17) + (d_2377_v22_))})
        d_2386_v31_: _dafny.Map
        d_2386_v31_ = _dafny.Map({d_2353_v0_: d_2353_v0_})
        index350_ = default__.safeIndex(273, (d_2381_v27_).length(0))
        index351_ = default__.safeIndex(163, (p0).length(0))
        rhs327_ = (d_2375_v20_)[default__.safeIndex(d_2353_v0_, len(d_2375_v20_))]
        rhs328_ = d_2353_v0_
        rhs329_ = ((d_2384_v30_)[(d_2386_v31_) | ((d_2386_v31_).set(d_2353_v0_, d_2353_v0_))] if ((d_2386_v31_) | ((d_2386_v31_).set(d_2353_v0_, d_2353_v0_))) in (d_2384_v30_) else d_2353_v0_)
        rhs330_ = (d_2377_v22_) + (d_2377_v22_)
        lhs182_ = d_2381_v27_
        lhs183_ = default__.safeIndex(273, (d_2381_v27_).length(0))
        lhs184_ = p0
        lhs185_ = default__.safeIndex(163, (p0).length(0))
        r1 = rhs327_
        lhs182_[lhs183_] = rhs328_
        d_2353_v0_ = rhs329_
        lhs184_[lhs185_] = rhs330_
        d_2379_v25_ = (d_2379_v25_).set(not(d_2373_v18_), d_2373_v18_)
        if d_2373_v18_:
            r3 = (d_2380_v26_)[default__.safeIndex(len((p0)[default__.safeIndex(163, (p0).length(0))]), len(d_2380_v26_))]
            d_2387_v32_: str
            d_2387_v32_ = _dafny.CodePoint('f')
            d_2388_v33_: _dafny.Map
            d_2388_v33_ = _dafny.Map({D0_DC0(d_2373_v18_, len(_dafny.SeqWithoutIsStrInference([(d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))] for d_2389_i2_ in range(default__.abs(-546))])), d_2373_v18_, len(d_2375_v20_), d_2373_v18_): d_2387_v32_})
            index352_ = default__.safeIndex(273, (d_2381_v27_).length(0))
            (d_2381_v27_)[index352_] = len((d_2388_v33_) | (d_2388_v33_))
            d_2390_v34_: _dafny.Seq
            d_2390_v34_ = _dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(163, (p0).length(0))]])
            d_2391_v35_: _dafny.MultiSet
            d_2391_v35_ = _dafny.MultiSet([(d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))], d_2353_v0_, len(d_2390_v34_), (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]])
            index353_ = default__.safeIndex(273, (d_2381_v27_).length(0))
            index354_ = default__.safeIndex(273, (d_2381_v27_).length(0))
            rhs331_ = (False if d_2373_v18_ else True)
            rhs332_ = d_2378_v24_
            rhs333_ = (len((d_2379_v25_) | (d_2379_v25_)) if d_2373_v18_ else (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))])
            rhs334_ = (d_2380_v26_)[default__.safeIndex((0) - (((_dafny.MultiSet([(d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))], 125])).intersection(d_2391_v35_)).cardinality), len(d_2380_v26_))]
            rhs335_ = (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]
            lhs186_ = d_2381_v27_
            lhs187_ = default__.safeIndex(273, (d_2381_v27_).length(0))
            lhs188_ = d_2381_v27_
            lhs189_ = default__.safeIndex(273, (d_2381_v27_).length(0))
            d_2373_v18_ = rhs331_
            d_2378_v24_ = rhs332_
            lhs186_[lhs187_] = rhs333_
            r3 = rhs334_
            lhs188_[lhs189_] = rhs335_
            d_2353_v0_ = (((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))] if d_2373_v18_ else len(d_2386_v31_))) + (len((p0)[default__.safeIndex(163, (p0).length(0))]))
            d_2392_v36_: D6
            d_2392_v36_ = D6_DC15((d_2380_v26_) + (d_2380_v26_))
            source31_ = d_2392_v36_
            if source31_.is_DC16:
                d_2393___mcc_h0_ = source31_.cf37
                d_2394___mcc_h1_ = source31_.cf38
                d_2395_cf38_ = d_2394___mcc_h1_
                d_2396_cf37_ = d_2393___mcc_h0_
                d_2397_v37_: _dafny.Set
                d_2397_v37_ = _dafny.Set({d_2373_v18_, (d_2375_v20_)[default__.safeIndex(d_2353_v0_, len(d_2375_v20_))]})
                (globalState).f0 = d_2397_v37_
                r3 = ((d_2353_v0_ if d_2373_v18_ else (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))])) + (len(d_2377_v22_))
                r2 = not(d_2373_v18_)
                d_2398_v38_: _dafny.Array
                nw376_ = _dafny.Array(None, 16)
                nw376_[int(0)] = _dafny.Map({d_2373_v18_: True})
                nw376_[int(1)] = (d_2379_v25_) | (d_2379_v25_)
                nw376_[int(2)] = d_2379_v25_
                nw376_[int(3)] = default__.fm46(globalState)
                nw376_[int(4)] = d_2379_v25_
                nw376_[int(5)] = d_2379_v25_
                nw376_[int(6)] = d_2379_v25_
                nw376_[int(7)] = d_2379_v25_
                nw376_[int(8)] = d_2379_v25_
                nw376_[int(9)] = d_2379_v25_
                nw376_[int(10)] = d_2379_v25_
                nw376_[int(11)] = ((d_2379_v25_).set(d_2373_v18_, d_2373_v18_)) | (d_2379_v25_)
                nw376_[int(12)] = d_2379_v25_
                nw376_[int(13)] = (d_2379_v25_) | (d_2379_v25_)
                nw376_[int(14)] = d_2379_v25_
                nw376_[int(15)] = d_2379_v25_
                d_2398_v38_ = nw376_
                index355_ = default__.safeIndex(80, (d_2398_v38_).length(0))
                (d_2398_v38_)[index355_] = (d_2379_v25_) | (d_2379_v25_)
                d_2399_v39_: _dafny.Seq
                d_2399_v39_ = _dafny.SeqWithoutIsStrInference([(d_2379_v25_) | (d_2379_v25_)])
                index356_ = default__.safeIndex(80, (d_2398_v38_).length(0))
                (d_2398_v38_)[index356_] = (d_2399_v39_)[default__.safeIndex(((((d_2374_v19_).set(d_2373_v18_, default__.abs(d_2353_v0_))).set(d_2373_v18_, default__.abs((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]))).cardinality) - (default__.fm25(d_2353_v0_, d_2395_cf38_, d_2373_v18_, globalState)), len(d_2399_v39_))]
            elif source31_.is_DC15:
                d_2400___mcc_h2_ = source31_.cf36
                d_2401_cf36_ = d_2400___mcc_h2_
                r0 = d_2353_v0_
                d_2374_v19_ = d_2374_v19_
                d_2402_v40_: D12
                d_2402_v40_ = D12_DC32(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2387_v32_])))
                d_2403_v41_: _dafny.Seq
                d_2403_v41_ = _dafny.SeqWithoutIsStrInference([d_2402_v40_])
                index357_ = default__.safeIndex(273, (d_2381_v27_).length(0))
                (d_2381_v27_)[index357_] = len((d_2403_v41_) + (d_2403_v41_))
                d_2386_v31_ = (d_2386_v31_) | (d_2386_v31_)
            elif True:
                d_2404___mcc_h3_ = source31_.cf39
                d_2405_cf39_ = d_2404___mcc_h3_
                d_2406_v42_: D5
                d_2406_v42_ = D5_DC14(d_2373_v18_, d_2375_v20_, d_2381_v27_, 173, (_dafny.MultiSet(d_2380_v26_)).cardinality)
                d_2407_v43_: _dafny.Map
                d_2407_v43_ = _dafny.Map({d_2373_v18_: d_2406_v42_})
                d_2408_v44_: _dafny.Map
                d_2408_v44_ = _dafny.Map({(not(d_2373_v18_)) and (d_2373_v18_): len(d_2407_v43_)})
                d_2409_v45_: D4
                d_2409_v45_ = D4_DC10(d_2373_v18_, d_2373_v18_, d_2375_v20_)
                pat_let_tv38_ = d_2375_v20_
                def iife199_(_pat_let54_0):
                    def iife200_(d_2410_dt__update__tmp_h0_):
                        def iife201_(_pat_let55_0):
                            def iife202_(d_2411_dt__update_hcf26_h0_):
                                return D4_DC10((d_2410_dt__update__tmp_h0_).cf24, (d_2410_dt__update__tmp_h0_).cf25, d_2411_dt__update_hcf26_h0_)
                            return iife202_(_pat_let55_0)
                        return iife201_(pat_let_tv38_)
                    return iife200_(_pat_let54_0)
                d_2408_v44_ = (d_2408_v44_).set((d_2375_v20_) < ((iife199_(d_2409_v45_)).cf26), d_2353_v0_)
                d_2412_v46_: _dafny.Map
                d_2412_v46_ = _dafny.Map({-474: (d_2373_v18_ if d_2373_v18_ else d_2373_v18_)})
                d_2412_v46_ = (d_2412_v46_).set(default__.fm26(globalState), d_2373_v18_)
                index358_ = default__.safeIndex(163, (p0).length(0))
                rhs336_ = (p0)[default__.safeIndex(163, (p0).length(0))]
                rhs337_ = (654) + ((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))])
                lhs190_ = p0
                lhs191_ = default__.safeIndex(163, (p0).length(0))
                lhs190_[lhs191_] = rhs336_
                r0 = rhs337_
                d_2413_v47_: _dafny.Array
                nw377_ = _dafny.Array(False, 3)
                d_2413_v47_ = nw377_
                d_2414_v48_: _dafny.Map
                d_2414_v48_ = _dafny.Map({d_2381_v27_: (self).f12})
                d_2413_v47_ = (((d_2414_v48_)[d_2381_v27_] if (d_2381_v27_) in (d_2414_v48_) else (self).f12) if d_2373_v18_ else (self).f12)
        elif True:
            if not(not(d_2373_v18_)):
                r2 = d_2373_v18_
                d_2415_v49_: C1
                nw378_ = C1()
                nw378_.ctor__(d_2353_v0_, False, False, (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))])
                d_2415_v49_ = nw378_
                d_2416_v51_: _dafny.MultiSet
                d_2416_v51_ = _dafny.MultiSet([(default__.fm47(d_2373_v18_, globalState)).set(not(d_2373_v18_), default__.abs(d_2353_v0_))])
                d_2417_v52_: _dafny.Set
                d_2417_v52_ = _dafny.Set({d_2373_v18_})
                d_2418_v53_: _dafny.MultiSet
                d_2418_v53_ = _dafny.MultiSet([d_2380_v26_, _dafny.SeqWithoutIsStrInference([len(d_2417_v52_)])])
                def iife203_():
                    coll91_ = _dafny.Map()
                    compr_91_: _dafny.MultiSet
                    for compr_91_ in (d_2416_v51_).Elements:
                        d_2419_v50_: _dafny.MultiSet = compr_91_
                        if (d_2419_v50_) in (d_2416_v51_):
                            coll91_[d_2419_v50_] = d_2373_v18_
                    return _dafny.Map(coll91_)
                r1 = not((len(iife203_()
                )) != ((d_2418_v53_).cardinality))
                (globalState).f0 = _dafny.Set({d_2373_v18_, not(True), d_2373_v18_, d_2373_v18_})
                d_2420_v54_: str
                d_2420_v54_ = _dafny.CodePoint('q')
                d_2421_v55_: D16
                d_2421_v55_ = D16_DC41(d_2420_v54_, (d_2373_v18_) or (d_2373_v18_))
                d_2421_v55_ = d_2421_v55_
            elif True:
                r2 = d_2373_v18_
                r2 = not(d_2373_v18_)
                d_2373_v18_ = True
                d_2422_v56_: _dafny.Seq
                out21_: _dafny.Seq
                out21_ = (self).m8((d_2353_v0_) != (984), d_2373_v18_, d_2353_v0_, globalState)
                d_2422_v56_ = out21_
                d_2373_v18_ = not((_dafny.CodePoint('q')) not in ((p0)[default__.safeIndex(163, (p0).length(0))]))
            r1 = (d_2373_v18_) == (not (d_2373_v18_) or (d_2373_v18_))
            d_2423_v57_: str
            d_2423_v57_ = _dafny.CodePoint('i')
            d_2424_v58_: C3
            nw379_ = C3()
            nw379_.ctor__(d_2423_v57_, (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))], not(d_2373_v18_), d_2373_v18_, (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))])
            d_2424_v58_ = nw379_
            index359_ = default__.safeIndex(736, ((self).f12).length(0))
            ((self).f12)[index359_] = d_2373_v18_
            d_2425_v59_: _dafny.Map
            d_2425_v59_ = _dafny.Map({d_2353_v0_: not(d_2373_v18_)})
            d_2426_v60_: _dafny.Map
            d_2426_v60_ = _dafny.Map({len(d_2425_v59_): d_2373_v18_})
            d_2427_v61_: _dafny.Map
            d_2427_v61_ = _dafny.Map({d_2425_v59_: default__.fm48(default__.fm49((p0)[default__.safeIndex(163, (p0).length(0))], d_2353_v0_, (d_2375_v20_)[default__.safeIndex(d_2353_v0_, len(d_2375_v20_))], d_2373_v18_, globalState), globalState)})
            d_2428_v62_: _dafny.Seq
            d_2428_v62_ = _dafny.SeqWithoutIsStrInference([d_2427_v61_, (d_2427_v61_ if d_2373_v18_ else d_2427_v61_)])
            index360_ = default__.safeIndex(736, ((self).f12).length(0))
            rhs338_ = (d_2373_v18_ if d_2373_v18_ else d_2373_v18_)
            rhs339_ = (d_2428_v62_)[default__.safeIndex((d_2380_v26_)[default__.safeIndex((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))], len(d_2380_v26_))], len(d_2428_v62_))]
            lhs192_ = (self).f12
            lhs193_ = default__.safeIndex(736, ((self).f12).length(0))
            lhs192_[lhs193_] = rhs338_
            d_2427_v61_ = rhs339_
            r0 = default__.safeModuloInt((0) - (((0) - (d_2353_v0_)) - (d_2353_v0_)), default__.fm55((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))], d_2375_v20_, (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))], globalState))
        d_2429_v63_: _dafny.Map
        d_2429_v63_ = _dafny.Map({(d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]: d_2373_v18_})
        if ((d_2429_v63_)[d_2353_v0_] if (d_2353_v0_) in (d_2429_v63_) else d_2373_v18_):
            if d_2373_v18_:
                d_2430_v64_: _dafny.Set
                d_2430_v64_ = _dafny.Set({(d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]})
                d_2431_v65_: D5
                d_2431_v65_ = D5_DC14(d_2373_v18_, _dafny.SeqWithoutIsStrInference([(d_2375_v20_)[default__.safeIndex(len(d_2430_v64_), len(d_2375_v20_))], d_2373_v18_]), d_2381_v27_, d_2353_v0_, 73)
                d_2432_v66_: D11
                d_2432_v66_ = D11_DC30(d_2373_v18_, (d_2431_v65_).cf33, (0) - ((0) - (len((p0)[default__.safeIndex(163, (p0).length(0))]))), d_2373_v18_)
                d_2433_v67_: _dafny.Map
                d_2433_v67_ = _dafny.Map({d_2432_v66_: (d_2353_v0_ if d_2373_v18_ else d_2353_v0_)})
                d_2433_v67_ = (d_2433_v67_).set(d_2432_v66_, (d_2353_v0_) * ((0) - ((0) - (d_2353_v0_))))
                r2 = d_2373_v18_
                d_2434_v68_: _dafny.Array
                nw380_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_2434_v68_ = nw380_
                index361_ = default__.safeIndex(136, (d_2434_v68_).length(0))
                (d_2434_v68_)[index361_] = (d_2381_v27_ if d_2373_v18_ else d_2381_v27_)
                index362_ = default__.safeIndex(136, (d_2434_v68_).length(0))
                (d_2434_v68_)[index362_] = d_2381_v27_
                d_2435_v69_: _dafny.Array
                nw381_ = _dafny.Array(None, 14)
                nw381_[int(0)] = (self).f12
                nw381_[int(1)] = (self).f12
                nw381_[int(2)] = (self).f12
                nw381_[int(3)] = (self).f12
                nw381_[int(4)] = (self).f12
                nw381_[int(5)] = ((self).f12 if d_2373_v18_ else (self).f12)
                nw381_[int(6)] = ((self).f12 if (d_2375_v20_)[default__.safeIndex(181, len(d_2375_v20_))] else (self).f12)
                nw381_[int(7)] = (self).f12
                nw381_[int(8)] = (self).f12
                nw381_[int(9)] = (self).f12
                nw381_[int(10)] = (self).f12
                nw381_[int(11)] = (self).f12
                nw381_[int(12)] = (self).f12
                nw381_[int(13)] = (self).f12
                d_2435_v69_ = nw381_
                d_2435_v69_ = d_2435_v69_
                d_2436_v70_: T2
                nw382_ = C7()
                nw382_.ctor__(676, d_2373_v18_, d_2373_v18_, (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))], len(d_2377_v22_), True, d_2377_v22_, False)
                d_2436_v70_ = nw382_
                d_2437_v71_: _dafny.Seq
                d_2437_v71_ = _dafny.SeqWithoutIsStrInference([d_2436_v70_])
                d_2438_v72_: D18
                d_2438_v72_ = D18_DC49(d_2436_v70_)
                d_2439_v73_: _dafny.Set
                d_2439_v73_ = _dafny.Set({(d_2436_v70_).f4})
                d_2440_v74_: _dafny.Map
                d_2440_v74_ = _dafny.Map({_dafny.Map({d_2436_v70_.f18: (D19_DC52(d_2439_v73_)).cf93}): d_2436_v70_})
                d_2441_v75_: _dafny.Map
                d_2441_v75_ = _dafny.Map({d_2373_v18_: _dafny.Set({(d_2436_v70_).f4})})
                d_2442_v76_: _dafny.Array
                nw383_ = _dafny.Array(None, 23)
                nw383_[int(0)] = d_2436_v70_
                nw383_[int(1)] = d_2436_v70_
                nw383_[int(2)] = (d_2437_v71_)[default__.safeIndex((d_2436_v70_).f5, len(d_2437_v71_))]
                nw383_[int(3)] = d_2436_v70_
                nw383_[int(4)] = d_2436_v70_
                nw383_[int(5)] = d_2436_v70_
                nw383_[int(6)] = (d_2438_v72_).cf89
                nw383_[int(7)] = d_2436_v70_
                nw383_[int(8)] = d_2436_v70_
                nw383_[int(9)] = d_2436_v70_
                nw383_[int(10)] = d_2436_v70_
                nw383_[int(11)] = d_2436_v70_
                nw383_[int(12)] = ((d_2437_v71_)[default__.safeIndex((d_2436_v70_).f5, len(d_2437_v71_))] if (d_2436_v70_).f4 else d_2436_v70_)
                nw383_[int(13)] = d_2436_v70_
                nw383_[int(14)] = d_2436_v70_
                nw383_[int(15)] = d_2436_v70_
                nw383_[int(16)] = ((d_2440_v74_)[d_2441_v75_] if (d_2441_v75_) in (d_2440_v74_) else d_2436_v70_)
                nw383_[int(17)] = d_2436_v70_
                nw383_[int(18)] = d_2436_v70_
                nw383_[int(19)] = d_2436_v70_
                nw383_[int(20)] = d_2436_v70_
                nw383_[int(21)] = d_2436_v70_
                nw383_[int(22)] = (d_2438_v72_).cf89
                d_2442_v76_ = nw383_
                nw384_ = _dafny.Array(None, 12)
                d_2442_v76_ = nw384_
            elif True:
                d_2373_v18_ = (d_2375_v20_)[default__.safeIndex((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))], len(d_2375_v20_))]
                index363_ = default__.safeIndex(273, (d_2381_v27_).length(0))
                (d_2381_v27_)[index363_] = (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]
                d_2443_v77_: _dafny.Array
                nw385_ = _dafny.Array(_dafny.MultiSet({}), 9)
                d_2443_v77_ = nw385_
                d_2444_v78_: str
                d_2444_v78_ = _dafny.CodePoint('o')
                d_2445_v79_: _dafny.Map
                d_2445_v79_ = _dafny.Map({default__.fm27(d_2444_v78_, globalState): 940})
                d_2446_v80_: _dafny.MultiSet
                d_2446_v80_ = _dafny.MultiSet([d_2445_v79_])
                index364_ = default__.safeIndex(980, (d_2443_v77_).length(0))
                (d_2443_v77_)[index364_] = d_2446_v80_
                index365_ = default__.safeIndex(980, (d_2443_v77_).length(0))
                (d_2443_v77_)[index365_] = d_2446_v80_
                d_2447_v81_: _dafny.Map
                d_2447_v81_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2373_v18_, d_2373_v18_]): d_2381_v27_})
                d_2447_v81_ = d_2447_v81_
                r2 = d_2373_v18_
            d_2448_v82_: _dafny.Array
            nw386_ = _dafny.Array(_dafny.Map({}), 7)
            d_2448_v82_ = nw386_
            index366_ = default__.safeIndex(595, (d_2448_v82_).length(0))
            (d_2448_v82_)[index366_] = d_2386_v31_
            index367_ = default__.safeIndex(595, (d_2448_v82_).length(0))
            (d_2448_v82_)[index367_] = _dafny.Map({(default__.fm18(False, d_2373_v18_, d_2373_v18_, globalState)) * ((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]): default__.safeModuloInt((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))], (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))])})
            d_2449_v83_: str
            d_2449_v83_ = _dafny.CodePoint('h')
            d_2450_v84_: D11
            d_2450_v84_ = D11_DC28(d_2373_v18_, (0) - ((0) - (len(d_2380_v26_))), d_2353_v0_, d_2373_v18_, default__.fm27(d_2449_v83_, globalState))
            d_2353_v0_ = len((d_2375_v20_) + ((d_2375_v20_) + (_dafny.SeqWithoutIsStrInference([(d_2450_v84_).cf52]))))
            r1 = (d_2373_v18_ if not (d_2373_v18_) or (not(d_2373_v18_)) else (d_2373_v18_) and (True))
            d_2451_v85_: _dafny.Map
            d_2451_v85_ = _dafny.Map({d_2449_v83_: d_2373_v18_})
            r2 = not(not (((0) - (d_2353_v0_)) == (len(_dafny.SeqWithoutIsStrInference([d_2373_v18_, d_2373_v18_, not(d_2373_v18_)])))) or (((d_2451_v85_)[_dafny.CodePoint('k')] if (_dafny.CodePoint('k')) in (d_2451_v85_) else d_2373_v18_)))
        elif True:
            d_2452_v86_: D11
            d_2452_v86_ = D11_DC29(d_2353_v0_, (self).f12, d_2373_v18_, d_2373_v18_)
            d_2453_v87_: D4
            d_2453_v87_ = D4_DC8(((d_2452_v86_).cf58 if d_2373_v18_ else (self).f12))
            source32_ = d_2453_v87_
            if source32_.is_DC9:
                d_2454___mcc_h4_ = source32_.cf21
                d_2455___mcc_h5_ = source32_.cf22
                d_2456___mcc_h6_ = source32_.cf23
                d_2457_cf23_ = d_2456___mcc_h6_
                d_2458_cf22_ = d_2455___mcc_h5_
                d_2459_cf21_ = d_2454___mcc_h4_
                d_2460_v88_: _dafny.Map
                d_2460_v88_ = _dafny.Map({d_2377_v22_: d_2386_v31_})
                d_2461_v90_: str
                d_2461_v90_ = _dafny.CodePoint('v')
                d_2462_v91_: _dafny.Map
                d_2462_v91_ = _dafny.Map({d_2457_cf23_: d_2461_v90_})
                d_2463_v92_: _dafny.Map
                d_2463_v92_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgqxsme")): d_2462_v91_})
                def iife204_():
                    coll92_ = _dafny.Map()
                    compr_92_: _dafny.Seq
                    for compr_92_ in (d_2463_v92_).keys.Elements:
                        d_2464_v89_: _dafny.Seq = compr_92_
                        if (d_2464_v89_) in (d_2463_v92_):
                            def iife205_():
                                coll93_ = _dafny.Map()
                                compr_93_: int
                                for compr_93_ in _dafny.IntegerRange(-462, 387):
                                    d_2465_v93_: int = compr_93_
                                    if ((-462) <= (d_2465_v93_)) and ((d_2465_v93_) < (387)):
                                        coll93_[default__.safeDivisionInt(d_2465_v93_, (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))])] = len(d_2459_cf21_)
                                return _dafny.Map(coll93_)
                            coll92_[d_2464_v89_] = iife205_()

                    return _dafny.Map(coll92_)
                d_2460_v88_ = (iife204_()
                ) | ((d_2460_v88_) | (d_2460_v88_))
                index368_ = default__.safeIndex(273, (d_2381_v27_).length(0))
                (d_2381_v27_)[index368_] = default__.safeModuloInt((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))], d_2353_v0_)
                d_2379_v25_ = (d_2379_v25_).set((d_2457_cf23_ if d_2457_cf23_ else d_2457_cf23_), False)
                d_2466_v94_: _dafny.Seq
                d_2466_v94_ = _dafny.SeqWithoutIsStrInference([d_2429_v63_, _dafny.Map({d_2458_cf22_: d_2457_cf23_})])
                d_2373_v18_ = (d_2429_v63_) not in (d_2466_v94_)
            elif source32_.is_DC10:
                d_2467___mcc_h7_ = source32_.cf24
                d_2468___mcc_h8_ = source32_.cf25
                d_2469___mcc_h9_ = source32_.cf26
                d_2470_cf26_ = d_2469___mcc_h9_
                d_2471_cf25_ = d_2468___mcc_h8_
                d_2472_cf24_ = d_2467___mcc_h7_
                d_2473_v95_: C1
                nw387_ = C1()
                nw387_.ctor__(((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]) * (d_2353_v0_), d_2471_cf25_, d_2471_cf25_, d_2353_v0_)
                d_2473_v95_ = nw387_
                d_2381_v27_ = d_2381_v27_
                r2 = not (not(d_2471_cf25_)) or (not(d_2472_cf24_))
                d_2474_v96_: _dafny.Array
                nw388_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
                d_2474_v96_ = nw388_
                d_2474_v96_ = p0
            elif source32_.is_DC8:
                d_2475___mcc_h10_ = source32_.cf20
                d_2476_cf20_ = d_2475___mcc_h10_
                d_2477_v97_: _dafny.Set
                d_2477_v97_ = _dafny.Set({d_2353_v0_})
                d_2478_v98_: _dafny.MultiSet
                d_2478_v98_ = _dafny.MultiSet([(0) - (len(d_2477_v97_))])
                d_2479_v99_: C2
                nw389_ = C2()
                nw389_.ctor__(d_2476_cf20_, d_2478_v98_, d_2377_v22_, d_2373_v18_, d_2373_v18_, (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))])
                d_2479_v99_ = nw389_
                rhs340_ = not((((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]) * (d_2353_v0_)) > (d_2353_v0_))
                rhs341_ = d_2373_v18_
                rhs342_ = d_2373_v18_
                d_2373_v18_ = rhs340_
                d_2373_v18_ = rhs341_
                r1 = rhs342_
                d_2480_v100_: _dafny.Array
                nw390_ = _dafny.Array(_dafny.Array(None, 0), 20)
                d_2480_v100_ = nw390_
                d_2480_v100_ = d_2480_v100_
                d_2481_v101_: _dafny.Map
                d_2481_v101_ = _dafny.Map({d_2373_v18_: d_2353_v0_})
                d_2482_v102_: _dafny.Seq
                d_2482_v102_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2373_v18_: ((d_2374_v19_)[d_2373_v18_] if (d_2373_v18_) in (d_2374_v19_) else d_2353_v0_)}), (d_2481_v101_).set(d_2373_v18_, 432), _dafny.Map({d_2373_v18_: (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]})])
                d_2482_v102_ = d_2482_v102_
            elif True:
                d_2483___mcc_h11_ = source32_.cf27
                d_2484_cf27_ = d_2483___mcc_h11_
                d_2485_v103_: _dafny.Map
                d_2485_v103_ = _dafny.Map({(d_2373_v18_) and (d_2373_v18_): d_2353_v0_})
                d_2485_v103_ = ((_dafny.Map({False: (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]})).set(d_2373_v18_, d_2353_v0_)) | (d_2485_v103_)
                d_2486_v104_: _dafny.Map
                d_2486_v104_ = _dafny.Map({234: (_dafny.Map({(d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]: d_2353_v0_})).set(d_2353_v0_, (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))])})
                d_2487_v105_: _dafny.Array
                nw391_ = _dafny.Array(None, 18)
                nw391_[int(0)] = d_2485_v103_
                nw391_[int(1)] = d_2485_v103_
                nw391_[int(2)] = (_dafny.Map({d_2373_v18_: len(d_2377_v22_)})) | (d_2485_v103_)
                nw391_[int(3)] = d_2485_v103_
                nw391_[int(4)] = (_dafny.Map({d_2373_v18_: (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]})).set(False, d_2353_v0_)
                nw391_[int(5)] = d_2485_v103_
                nw391_[int(6)] = (default__.fm57(((d_2486_v104_)[d_2353_v0_] if (d_2353_v0_) in (d_2486_v104_) else d_2386_v31_), d_2353_v0_, globalState)) | (d_2485_v103_)
                nw391_[int(7)] = (d_2485_v103_) | (d_2485_v103_)
                nw391_[int(8)] = (d_2485_v103_ if d_2373_v18_ else d_2485_v103_)
                nw391_[int(9)] = (d_2485_v103_ if d_2373_v18_ else d_2485_v103_)
                nw391_[int(10)] = d_2485_v103_
                nw391_[int(11)] = (d_2485_v103_) | (d_2485_v103_)
                nw391_[int(12)] = (d_2485_v103_) | (d_2485_v103_)
                nw391_[int(13)] = (d_2485_v103_).set(d_2373_v18_, (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))])
                nw391_[int(14)] = d_2485_v103_
                nw391_[int(15)] = (d_2485_v103_ if d_2373_v18_ else d_2485_v103_)
                nw391_[int(16)] = d_2485_v103_
                nw391_[int(17)] = d_2485_v103_
                d_2487_v105_ = nw391_
                index369_ = default__.safeIndex(674, (d_2487_v105_).length(0))
                (d_2487_v105_)[index369_] = d_2485_v103_
                d_2488_v106_: _dafny.Array
                def lambda103_(d_2489_v18_):
                    def lambda104_(d_2490_i3_):
                        return _dafny.SeqWithoutIsStrInference([d_2489_v18_, d_2489_v18_, d_2489_v18_])

                    return lambda104_

                init60_ = lambda103_(d_2373_v18_)
                nw392_ = _dafny.Array(None, 4)
                for i0_60_ in range(nw392_.length(0)):
                    nw392_[i0_60_] = init60_(i0_60_)
                d_2488_v106_ = nw392_
                index370_ = default__.safeIndex(943, (d_2488_v106_).length(0))
                (d_2488_v106_)[index370_] = d_2375_v20_
                index371_ = default__.safeIndex(674, (d_2487_v105_).length(0))
                index372_ = default__.safeIndex(943, (d_2488_v106_).length(0))
                rhs343_ = d_2485_v103_
                rhs344_ = (_dafny.SeqWithoutIsStrInference([d_2373_v18_])) + (d_2375_v20_)
                lhs194_ = d_2487_v105_
                lhs195_ = default__.safeIndex(674, (d_2487_v105_).length(0))
                lhs196_ = d_2488_v106_
                lhs197_ = default__.safeIndex(943, (d_2488_v106_).length(0))
                lhs194_[lhs195_] = rhs343_
                lhs196_[lhs197_] = rhs344_
                d_2353_v0_ = len(d_2377_v22_)
                d_2353_v0_ = (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]
            if False:
                d_2491_v107_: D17
                d_2491_v107_ = D17_DC47((d_2375_v20_) <= (d_2375_v20_), (len(_dafny.SeqWithoutIsStrInference([(d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))], -970]))) == (d_2353_v0_), 653, len(_dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(163, (p0).length(0))], (p0)[default__.safeIndex(163, (p0).length(0))], d_2377_v22_])), d_2353_v0_)
                d_2491_v107_ = d_2491_v107_
                r3 = (d_2353_v0_) * ((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))])
                r3 = ((0) - (((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]) * ((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]))) + (d_2353_v0_)
                d_2492_v108_: _dafny.Map
                d_2492_v108_ = _dafny.Map({(self).f12: d_2373_v18_})
                d_2493_v109_: _dafny.Map
                d_2493_v109_ = _dafny.Map({d_2373_v18_: d_2353_v0_})
                d_2429_v63_ = (d_2429_v63_).set(len(d_2492_v108_), (D4_DC9(d_2493_v109_, d_2353_v0_, d_2373_v18_)).cf23)
                d_2494_v110_: C1
                nw393_ = C1()
                nw393_.ctor__((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))], ((d_2379_v25_)[d_2373_v18_] if (d_2373_v18_) in (d_2379_v25_) else d_2373_v18_), d_2373_v18_, d_2353_v0_)
                d_2494_v110_ = nw393_
            elif True:
                d_2495_v111_: _dafny.Map
                d_2495_v111_ = _dafny.Map({len((d_2377_v22_) + (d_2377_v22_)): _dafny.MultiSet([(d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]])})
                d_2496_v112_: _dafny.Array
                nw394_ = _dafny.Array(None, 4)
                d_2496_v112_ = nw394_
                d_2497_v113_: D7
                d_2497_v113_ = D7_DC19(d_2353_v0_)
                d_2498_v114_: _dafny.Set
                d_2498_v114_ = _dafny.Set({d_2380_v26_, d_2380_v26_})
                rhs345_ = (d_2379_v25_) | (d_2379_v25_)
                rhs346_ = (d_2373_v18_) and ((d_2498_v114_).ispropersubset(d_2498_v114_))
                rhs347_ = ((d_2495_v111_) | (d_2495_v111_)) | (d_2495_v111_)
                rhs348_ = d_2496_v112_
                rhs349_ = (d_2497_v113_ if d_2373_v18_ else d_2497_v113_)
                d_2379_v25_ = rhs345_
                r1 = rhs346_
                d_2495_v111_ = rhs347_
                d_2496_v112_ = rhs348_
                d_2497_v113_ = rhs349_
                d_2499_v115_: _dafny.Seq
                out22_: _dafny.Seq
                out22_ = (self).m8(d_2373_v18_, d_2373_v18_, default__.safeDivisionInt(d_2353_v0_, (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]), globalState)
                d_2499_v115_ = out22_
                r0 = 432
                d_2500_v116_: _dafny.Seq
                d_2500_v116_ = _dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(163, (p0).length(0))], d_2377_v22_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbkdelo"))])
                d_2353_v0_ = len(default__.fm28(d_2377_v22_, len(d_2500_v116_), d_2353_v0_, globalState))
                d_2453_v87_ = (D4_DC8((self).f12) if d_2373_v18_ else d_2453_v87_)
            d_2501_v117_: _dafny.MultiSet
            d_2501_v117_ = _dafny.MultiSet([d_2353_v0_])
            d_2502_v118_: C2
            nw395_ = C2()
            nw395_.ctor__((self).f12, d_2501_v117_, (p0)[default__.safeIndex(163, (p0).length(0))], d_2373_v18_, d_2373_v18_, ((d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]) * (d_2353_v0_))
            d_2502_v118_ = nw395_
            d_2503_v119_: _dafny.Seq
            d_2503_v119_ = _dafny.SeqWithoutIsStrInference([d_2381_v27_, d_2381_v27_, d_2381_v27_, d_2381_v27_, d_2381_v27_])
            d_2504_v120_: _dafny.MultiSet
            d_2504_v120_ = _dafny.MultiSet([d_2381_v27_, d_2381_v27_, d_2381_v27_, d_2381_v27_, ((d_2503_v119_)[default__.safeIndex(d_2353_v0_, len(d_2503_v119_))] if d_2373_v18_ else d_2381_v27_)])
            d_2505_v121_: _dafny.Map
            d_2505_v121_ = _dafny.Map({False: d_2504_v120_})
            d_2504_v120_ = (((d_2505_v121_)[d_2373_v18_] if (d_2373_v18_) in (d_2505_v121_) else _dafny.MultiSet([d_2381_v27_]))).intersection((d_2504_v120_) - (d_2504_v120_))
            d_2506_v122_: _dafny.Seq
            d_2506_v122_ = _dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(163, (p0).length(0))], d_2377_v22_])
            d_2507_v123_: _dafny.Map
            d_2507_v123_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dn")): d_2373_v18_})
            d_2508_v124_: _dafny.Map
            d_2508_v124_ = _dafny.Map({(d_2506_v122_)[default__.safeIndex(d_2353_v0_, len(d_2506_v122_))]: d_2507_v123_})
            def iife206_():
                coll94_ = _dafny.Map()
                compr_94_: _dafny.Seq
                for compr_94_ in (d_2507_v123_).keys.Elements:
                    d_2509_v125_: _dafny.Seq = compr_94_
                    if (d_2509_v125_) in (d_2507_v123_):
                        coll94_[d_2509_v125_] = d_2373_v18_
                return _dafny.Map(coll94_)
            r2 = ((((d_2508_v124_)[(p0)[default__.safeIndex(163, (p0).length(0))]] if ((p0)[default__.safeIndex(163, (p0).length(0))]) in (d_2508_v124_) else iife206_()
            )) | (d_2507_v123_)) == (d_2507_v123_)
        d_2353_v0_ = d_2353_v0_
        r0 = d_2353_v0_
        r1 = d_2373_v18_
        d_2510_v126_: _dafny.Set
        d_2510_v126_ = _dafny.Set({d_2353_v0_})
        d_2511_v127_: _dafny.Map
        d_2511_v127_ = _dafny.Map({d_2510_v126_: (0) - (d_2353_v0_)})
        r2 = (d_2510_v126_) not in (d_2511_v127_)
        r3 = (d_2381_v27_)[default__.safeIndex(273, (d_2381_v27_).length(0))]
        return r0, r1, r2, r3

    @property
    def f12(self):
        return self._f12

class C10:
    def  __init__(self):
        self._f10: int = int(0)
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    def ctor__(self, f10, f11):
        (self)._f10 = f10
        (self)._f11 = f11

    def fm14(self, globalState):
        if True:
            return _dafny.Map({not(not(False)): 871})
        elif True:
            return _dafny.Map({False: (self).f10})

    def fm15(self, globalState):
        return not((True) == (True))

    def m6(self, p0, p1, p2, p3, globalState):
        d_2512_v0_: bool
        d_2512_v0_ = True
        d_2512_v0_ = d_2512_v0_
        if p1:
            d_2513_v1_: _dafny.MultiSet
            d_2513_v1_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cuwx")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owott"))])
            d_2513_v1_ = d_2513_v1_
            d_2512_v0_ = (self).fm15(globalState)
            d_2514_v2_: D4
            d_2514_v2_ = D4_DC8((self).f11)
            d_2515_v3_: _dafny.Array
            nw396_ = _dafny.Array(None, 22)
            nw396_[int(0)] = (self).f11
            nw396_[int(1)] = (self).f11
            nw396_[int(2)] = (self).f11
            nw396_[int(3)] = (self).f11
            nw396_[int(4)] = (self).f11
            nw396_[int(5)] = (self).f11
            nw396_[int(6)] = (self).f11
            nw396_[int(7)] = (self).f11
            nw396_[int(8)] = (self).f11
            nw396_[int(9)] = (self).f11
            nw396_[int(10)] = (d_2514_v2_).cf20
            nw396_[int(11)] = (self).f11
            nw396_[int(12)] = (self).f11
            nw396_[int(13)] = (self).f11
            nw396_[int(14)] = (self).f11
            nw396_[int(15)] = (self).f11
            nw396_[int(16)] = (self).f11
            nw396_[int(17)] = (self).f11
            nw396_[int(18)] = (self).f11
            nw396_[int(19)] = (self).f11
            nw396_[int(20)] = (self).f11
            nw396_[int(21)] = (self).f11
            d_2515_v3_ = nw396_
            d_2516_v4_: _dafny.Seq
            d_2516_v4_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11])
            index373_ = default__.safeIndex(538, (d_2515_v3_).length(0))
            (d_2515_v3_)[index373_] = (d_2516_v4_)[default__.safeIndex(p0, len(d_2516_v4_))]
            index374_ = default__.safeIndex(538, (d_2515_v3_).length(0))
            (d_2515_v3_)[index374_] = (d_2516_v4_)[default__.safeIndex(p2, len(d_2516_v4_))]
            index375_ = default__.safeIndex(180, ((self).f11).length(0))
            ((self).f11)[index375_] = p1
            index376_ = default__.safeIndex(180, ((self).f11).length(0))
            ((self).f11)[index376_] = p1
            d_2517_v5_: int
            d_2517_v5_ = -486
            d_2517_v5_ = ((d_2517_v5_) - (p0)) * (p0)
        elif True:
            d_2518_v6_: _dafny.Map
            d_2518_v6_ = _dafny.Map({p3: p3})
            d_2518_v6_ = ((_dafny.Map({p3: p3})) | (d_2518_v6_)) | (_dafny.Map({(self).f10: p0}))
            d_2519_v7_: _dafny.Set
            d_2519_v7_ = _dafny.Set({p1, d_2512_v0_})
            d_2520_v8_: _dafny.Seq
            d_2520_v8_ = _dafny.SeqWithoutIsStrInference([d_2519_v7_])
            d_2521_v9_: _dafny.Seq
            d_2521_v9_ = _dafny.SeqWithoutIsStrInference([d_2520_v8_])
            d_2522_v10_: _dafny.Seq
            d_2522_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lto"))
            d_2523_v11_: _dafny.Map
            d_2523_v11_ = _dafny.Map({p3: p1})
            d_2524_v12_: _dafny.MultiSet
            d_2524_v12_ = _dafny.MultiSet([p0])
            d_2525_v13_: C6
            nw397_ = C6()
            nw397_.ctor__(default__.safeDivisionInt((0) - (len((d_2521_v9_)[default__.safeIndex(len(d_2522_v10_), len(d_2521_v9_))])), len(d_2523_v11_)), p1, False, ((d_2524_v12_)[p3] if (p3) in (d_2524_v12_) else (self).f10), (p3) - ((self).f10), False)
            d_2525_v13_ = nw397_
            index377_ = default__.safeIndex(837, ((self).f11).length(0))
            ((self).f11)[index377_] = False
            index378_ = default__.safeIndex(578, ((self).f11).length(0))
            ((self).f11)[index378_] = not(p1)
            d_2526_v15_: _dafny.Array
            def lambda105_(d_2527_v0_, d_2528_v13_, d_2529_p0_, d_2530_p2_):
                def lambda106_(d_2531_i0_):
                    def iife207_():
                        coll95_ = _dafny.Set()
                        compr_95_: _dafny.Seq
                        for compr_95_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])])).Elements:
                            d_2532_v14_: _dafny.Seq = compr_95_
                            if (d_2532_v14_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])])):
                                coll95_ = coll95_.union(_dafny.Set([d_2532_v14_]))
                        return _dafny.Set(coll95_)
                    return (_dafny.Set({_dafny.SeqWithoutIsStrInference([d_2528_v13_.f21, d_2529_p0_]), _dafny.SeqWithoutIsStrInference([d_2528_v13_.f21]), _dafny.SeqWithoutIsStrInference([d_2530_p2_])}) if d_2527_v0_ else iife207_()
                    )

                return lambda106_

            init61_ = lambda105_(d_2512_v0_, d_2525_v13_, p0, p2)
            nw398_ = _dafny.Array(None, 17)
            for i0_61_ in range(nw398_.length(0)):
                nw398_[i0_61_] = init61_(i0_61_)
            d_2526_v15_ = nw398_
            d_2533_v17_: D0
            d_2533_v17_ = D0_DC0((d_2525_v13_).f22, p3, (d_2525_v13_).fm0(globalState), (d_2524_v12_).cardinality, d_2512_v0_)
            d_2534_v18_: _dafny.Seq
            d_2534_v18_ = _dafny.SeqWithoutIsStrInference([d_2512_v0_])
            d_2535_v19_: _dafny.Seq
            d_2535_v19_ = _dafny.SeqWithoutIsStrInference([p0, len(d_2534_v18_)])
            d_2536_v20_: _dafny.Set
            def iife208_():
                coll96_ = _dafny.Set()
                compr_96_: int
                for compr_96_ in _dafny.IntegerRange(27, 498):
                    d_2537_v16_: int = compr_96_
                    if ((27) <= (d_2537_v16_)) and ((d_2537_v16_) < (498)):
                        coll96_ = coll96_.union(_dafny.Set([default__.safeDivisionInt(d_2537_v16_, p0)]))
                return _dafny.Set(coll96_)
            d_2536_v20_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([len(iife208_()
) for d_2538_i1_ in range(default__.abs(749))]), _dafny.SeqWithoutIsStrInference([p0, (self).f10, p2]), default__.fm44(p2, d_2533_v17_, d_2512_v0_, (d_2525_v13_).f22, globalState), d_2535_v19_})
            index379_ = default__.safeIndex(87, (d_2526_v15_).length(0))
            (d_2526_v15_)[index379_] = d_2536_v20_
            d_2539_v21_: _dafny.Map
            d_2539_v21_ = _dafny.Map({d_2512_v0_: (0) - (p3)})
            index380_ = default__.safeIndex(837, ((self).f11).length(0))
            index381_ = default__.safeIndex(578, ((self).f11).length(0))
            index382_ = default__.safeIndex(87, (d_2526_v15_).length(0))
            rhs350_ = (d_2524_v12_).isdisjoint((d_2524_v12_).set(p3, default__.abs(d_2525_v13_.f21)))
            rhs351_ = not(d_2512_v0_)
            rhs352_ = (d_2534_v18_)[default__.safeIndex(len((_dafny.Map({p1: len(d_2535_v19_)})) | (d_2539_v21_)), len(d_2534_v18_))]
            rhs353_ = (d_2536_v20_) | (d_2536_v20_)
            rhs354_ = (d_2522_v10_) not in (default__.fm51(globalState))
            lhs198_ = (self).f11
            lhs199_ = default__.safeIndex(837, ((self).f11).length(0))
            lhs200_ = (self).f11
            lhs201_ = default__.safeIndex(578, ((self).f11).length(0))
            lhs202_ = d_2526_v15_
            lhs203_ = default__.safeIndex(87, (d_2526_v15_).length(0))
            d_2512_v0_ = rhs350_
            lhs198_[lhs199_] = rhs351_
            lhs200_[lhs201_] = rhs352_
            lhs202_[lhs203_] = rhs353_
            d_2512_v0_ = rhs354_
            d_2512_v0_ = not((d_2525_v13_).f22)
            d_2540_v22_: _dafny.Seq
            d_2540_v22_ = _dafny.SeqWithoutIsStrInference([d_2522_v10_, d_2522_v10_, d_2522_v10_, d_2522_v10_])
            rhs355_ = (self).f10
            rhs356_ = d_2524_v12_
            rhs357_ = _dafny.SeqWithoutIsStrInference([(d_2522_v10_) + (d_2522_v10_) for d_2541_i2_ in range(default__.abs(184))])
            lhs204_ = d_2525_v13_
            lhs204_.f21 = rhs355_
            d_2524_v12_ = rhs356_
            d_2540_v22_ = rhs357_
        if p1:
            index383_ = default__.safeIndex(19, ((self).f11).length(0))
            ((self).f11)[index383_] = d_2512_v0_
            d_2542_v23_: _dafny.Seq
            d_2542_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpgtu"))
            index384_ = default__.safeIndex(19, ((self).f11).length(0))
            ((self).f11)[index384_] = (len(d_2542_v23_)) != ((self).f10)
            d_2543_v24_: D5
            d_2543_v24_ = D5_DC13(True, p1)
            d_2544_v25_: _dafny.Map
            d_2544_v25_ = _dafny.Map({p3: d_2543_v24_})
            d_2545_v26_: _dafny.MultiSet
            d_2545_v26_ = _dafny.MultiSet([d_2544_v25_])
            index385_ = default__.safeIndex(19, ((self).f11).length(0))
            rhs358_ = p1
            rhs359_ = d_2545_v26_
            lhs205_ = (self).f11
            lhs206_ = default__.safeIndex(19, ((self).f11).length(0))
            lhs205_[lhs206_] = rhs358_
            d_2545_v26_ = rhs359_
            d_2546_v27_: _dafny.Map
            d_2546_v27_ = _dafny.Map({p2: p3})
            rhs360_ = (d_2542_v23_) + ((d_2542_v23_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))))
            rhs361_ = _dafny.Map({p3: p3})
            d_2542_v23_ = rhs360_
            d_2546_v27_ = rhs361_
            index386_ = default__.safeIndex(19, ((self).f11).length(0))
            ((self).f11)[index386_] = True
            d_2547_v28_: bool
            out23_: bool
            out23_ = default__.m0(p1, default__.fm49(d_2542_v23_, len(d_2546_v27_), True, not(True), globalState), globalState)
            d_2547_v28_ = out23_
        elif True:
            d_2548_v29_: _dafny.Seq
            d_2548_v29_ = _dafny.SeqWithoutIsStrInference([d_2512_v0_])
            d_2549_v30_: _dafny.Seq
            d_2549_v30_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2550_v31_: _dafny.Seq
            d_2550_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whdfafflx"))
            d_2551_v32_: _dafny.Map
            d_2551_v32_ = _dafny.Map({d_2549_v30_: len(d_2550_v31_)})
            d_2552_v33_: _dafny.Map
            d_2552_v33_ = _dafny.Map({d_2548_v29_: d_2551_v32_})
            d_2553_v34_: _dafny.MultiSet
            d_2553_v34_ = _dafny.MultiSet([p0])
            def iife209_():
                coll97_ = _dafny.Map()
                compr_97_: _dafny.Seq
                for compr_97_ in ((default__.fm33(p3, p1, p3, d_2512_v0_, globalState)).set(d_2549_v30_, default__.abs(p2))).Elements:
                    d_2554_v35_: _dafny.Seq = compr_97_
                    if (d_2554_v35_) in ((default__.fm33(p3, p1, p3, d_2512_v0_, globalState)).set(d_2549_v30_, default__.abs(p2))):
                        coll97_[d_2554_v35_] = p3
                return _dafny.Map(coll97_)
            d_2552_v33_ = (d_2552_v33_).set((_dafny.SeqWithoutIsStrInference([p1, p1])) + (_dafny.SeqWithoutIsStrInference([(d_2548_v29_)[default__.safeIndex(default__.fm25(default__.fm49(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tosehbr")), ((d_2553_v34_)[p2] if (p2) in (d_2553_v34_) else len(d_2550_v31_)), d_2512_v0_, d_2512_v0_, globalState), p2, p1, globalState), len(d_2548_v29_))]])), iife209_()
            )
            d_2555_v36_: D1
            d_2555_v36_ = D1_DC4(d_2512_v0_, p3, (self).f10)
            d_2555_v36_ = d_2555_v36_
            d_2556_v37_: str
            d_2556_v37_ = _dafny.CodePoint('e')
            d_2556_v37_ = d_2556_v37_
            d_2557_v38_: _dafny.Array
            nw399_ = _dafny.Array(False, 22)
            d_2557_v38_ = nw399_
            d_2558_v39_: _dafny.MultiSet
            d_2558_v39_ = _dafny.MultiSet([True, not(d_2512_v0_), d_2512_v0_])
            d_2559_v40_: _dafny.Map
            d_2559_v40_ = _dafny.Map({d_2548_v29_: False})
            d_2560_v41_: _dafny.Seq
            d_2560_v41_ = _dafny.SeqWithoutIsStrInference([d_2557_v38_, d_2557_v38_, (self).f11, d_2557_v38_])
            d_2557_v38_ = ((d_2557_v38_ if ((d_2559_v40_)[d_2548_v29_] if (d_2548_v29_) in (d_2559_v40_) else d_2512_v0_) else (self).f11) if (d_2558_v39_).issubset(_dafny.MultiSet([True])) else (d_2560_v41_)[default__.safeIndex(p2, len(d_2560_v41_))])
            d_2561_v42_: C5
            nw400_ = C5()
            nw400_.ctor__()
            d_2561_v42_ = nw400_
        d_2562_v43_: bool
        out24_: bool
        out24_ = default__.m0(p1, default__.safeModuloInt(p0, p3), globalState)
        d_2562_v43_ = out24_
        d_2563_v44_: int
        d_2563_v44_ = 336
        d_2564_v45_: _dafny.Seq
        d_2564_v45_ = _dafny.SeqWithoutIsStrInference([p2, default__.fm18(d_2512_v0_, not(p1), d_2562_v43_, globalState)])
        d_2565_v46_: _dafny.Map
        d_2565_v46_ = _dafny.Map({(self).f10: d_2562_v43_})
        d_2566_v47_: _dafny.MultiSet
        d_2566_v47_ = _dafny.MultiSet([(0) - (p2)])
        d_2567_v48_: _dafny.Seq
        d_2567_v48_ = _dafny.SeqWithoutIsStrInference([d_2564_v45_, _dafny.SeqWithoutIsStrInference([962, len((d_2565_v46_).set((d_2566_v47_).cardinality, False))]), d_2564_v45_, d_2564_v45_, d_2564_v45_])
        d_2563_v44_ = ((len(d_2567_v48_) if p1 else p2) if p1 else (d_2563_v44_) + (p3))
        d_2568_v49_: str
        d_2568_v49_ = _dafny.CodePoint('t')
        if default__.fm27(d_2568_v49_, globalState):
            d_2569_v50_: _dafny.Array
            def lambda107_(d_2570_p2_, d_2571_v0_):
                def lambda108_(d_2572_i3_):
                    return _dafny.SeqWithoutIsStrInference([d_2570_p2_, len(_dafny.SeqWithoutIsStrInference([d_2571_v0_]))])

                return lambda108_

            init62_ = lambda107_(p2, d_2512_v0_)
            nw401_ = _dafny.Array(None, 12)
            for i0_62_ in range(nw401_.length(0)):
                nw401_[i0_62_] = init62_(i0_62_)
            d_2569_v50_ = nw401_
            index387_ = default__.safeIndex(810, (d_2569_v50_).length(0))
            (d_2569_v50_)[index387_] = (d_2564_v45_) + (d_2564_v45_)
            d_2573_v51_: D6
            d_2573_v51_ = D6_DC15(d_2564_v45_)
            d_2574_v52_: _dafny.Map
            d_2574_v52_ = _dafny.Map({d_2564_v45_: p1})
            index388_ = default__.safeIndex(810, (d_2569_v50_).length(0))
            rhs362_ = (d_2564_v45_) + ((d_2573_v51_).cf36)
            rhs363_ = ((0) - (p3) if ((d_2574_v52_)[d_2564_v45_] if (d_2564_v45_) in (d_2574_v52_) else d_2562_v43_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "db"))))
            rhs364_ = d_2564_v45_
            lhs207_ = d_2569_v50_
            lhs208_ = default__.safeIndex(810, (d_2569_v50_).length(0))
            d_2564_v45_ = rhs362_
            d_2563_v44_ = rhs363_
            lhs207_[lhs208_] = rhs364_
            d_2575_v53_: _dafny.Array
            nw402_ = _dafny.Array(None, 3)
            nw402_[int(0)] = d_2563_v44_
            nw402_[int(1)] = default__.safeDivisionInt(p0, p3)
            nw402_[int(2)] = p3
            d_2575_v53_ = nw402_
            d_2575_v53_ = d_2575_v53_
            d_2576_v54_: _dafny.Map
            d_2576_v54_ = _dafny.Map({d_2512_v0_: False})
            d_2577_v55_: _dafny.MultiSet
            d_2577_v55_ = _dafny.MultiSet([p1, p1])
            d_2576_v54_ = (d_2576_v54_).set((p2) > (p3), ((d_2577_v55_).cardinality) == (((d_2569_v50_)[default__.safeIndex(810, (d_2569_v50_).length(0))])[default__.safeIndex(d_2563_v44_, len((d_2569_v50_)[default__.safeIndex(810, (d_2569_v50_).length(0))]))]))
            d_2578_v56_: _dafny.Seq
            d_2578_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khygymr"))
            d_2579_v57_: _dafny.Map
            d_2579_v57_ = _dafny.Map({not((self).fm15(globalState)): d_2578_v56_})
            d_2579_v57_ = (d_2579_v57_).set(d_2562_v43_, d_2578_v56_)
            d_2580_v58_: D18
            d_2580_v58_ = D18_DC50(p1, d_2512_v0_)
            d_2581_v59_: _dafny.Map
            d_2581_v59_ = _dafny.Map({d_2575_v53_: d_2580_v58_})
            d_2582_v60_: D18
            d_2582_v60_ = D18_DC51(((d_2581_v59_)[d_2575_v53_] if (d_2575_v53_) in (d_2581_v59_) else d_2580_v58_))
            d_2583_v61_: _dafny.MultiSet
            d_2583_v61_ = _dafny.MultiSet([d_2582_v60_])
            d_2583_v61_ = d_2583_v61_
        elif True:
            d_2512_v0_ = True
            if (p3) >= (p0):
                d_2584_v62_: _dafny.Seq
                d_2584_v62_ = _dafny.SeqWithoutIsStrInference([d_2512_v0_])
                index389_ = default__.safeIndex(901, ((self).f11).length(0))
                ((self).f11)[index389_] = (d_2566_v47_).issubset(_dafny.MultiSet([len(d_2584_v62_)]))
                d_2585_v63_: _dafny.Set
                d_2585_v63_ = _dafny.Set({235, default__.fm18(d_2512_v0_, d_2512_v0_, p1, globalState)})
                index390_ = default__.safeIndex(901, ((self).f11).length(0))
                rhs365_ = default__.safeModuloInt(default__.safeDivisionInt(p0, (d_2564_v45_)[default__.safeIndex(p0, len(d_2564_v45_))]), p3)
                rhs366_ = (d_2585_v63_).issubset(d_2585_v63_)
                rhs367_ = d_2563_v44_
                rhs368_ = (d_2584_v62_)[default__.safeIndex((p2) - ((self).f10), len(d_2584_v62_))]
                lhs209_ = (self).f11
                lhs210_ = default__.safeIndex(901, ((self).f11).length(0))
                d_2563_v44_ = rhs365_
                lhs209_[lhs210_] = rhs366_
                d_2563_v44_ = rhs367_
                d_2562_v43_ = rhs368_
                d_2563_v44_ = (0) - ((0) - ((0) - ((0) - ((936) - (p3)))))
                index391_ = default__.safeIndex(901, ((self).f11).length(0))
                ((self).f11)[index391_] = not(not(((self).f11)[default__.safeIndex(901, ((self).f11).length(0))]))
                index392_ = default__.safeIndex(901, ((self).f11).length(0))
                ((self).f11)[index392_] = d_2512_v0_
                d_2586_v64_: _dafny.Map
                d_2586_v64_ = _dafny.Map({p0: d_2566_v47_})
                def iife210_():
                    coll98_ = _dafny.Map()
                    compr_98_: int
                    for compr_98_ in (d_2564_v45_).Elements:
                        d_2587_v65_: int = compr_98_
                        if (d_2587_v65_) in (d_2564_v45_):
                            coll98_[(d_2587_v65_) + ((self).f10)] = d_2566_v47_
                    return _dafny.Map(coll98_)
                d_2586_v64_ = iife210_()
                
            elif True:
                d_2588_v66_: C9
                nw403_ = C9()
                nw403_.ctor__((self).f11)
                d_2588_v66_ = nw403_
                d_2589_v67_: _dafny.Map
                d_2589_v67_ = _dafny.Map({d_2568_v49_: (d_2564_v45_)[default__.safeIndex(len(d_2565_v46_), len(d_2564_v45_))]})
                d_2589_v67_ = (d_2589_v67_).set(d_2568_v49_, p2)
                d_2590_v68_: _dafny.Seq
                d_2590_v68_ = _dafny.SeqWithoutIsStrInference([d_2562_v43_, (p2) <= (p2)])
                d_2512_v0_ = (d_2590_v68_)[default__.safeIndex(620, len(d_2590_v68_))]
                d_2512_v0_ = p1
                d_2563_v44_ = p3
            if not(d_2512_v0_):
                d_2563_v44_ = (74) + (p3)
                d_2591_v69_: _dafny.Seq
                d_2591_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                d_2592_v70_: _dafny.Map
                d_2592_v70_ = _dafny.Map({(self).f11: len(d_2591_v69_)})
                d_2593_v71_: D1
                d_2593_v71_ = D1_DC5(p1, (0) - (len(d_2591_v69_)), d_2591_v69_)
                pat_let_tv39_ = p3
                pat_let_tv40_ = d_2512_v0_
                d_2594_v72_: _dafny.Array
                nw404_ = _dafny.Array(None, 16)
                nw404_[int(0)] = (d_2591_v69_) + (d_2591_v69_)
                nw404_[int(1)] = d_2591_v69_
                nw404_[int(2)] = d_2591_v69_
                nw404_[int(3)] = d_2591_v69_
                nw404_[int(4)] = d_2591_v69_
                nw404_[int(5)] = (d_2591_v69_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbyu")))
                nw404_[int(6)] = (d_2591_v69_) + (d_2591_v69_)
                nw404_[int(7)] = d_2591_v69_
                nw404_[int(8)] = d_2591_v69_
                nw404_[int(9)] = d_2591_v69_
                nw404_[int(10)] = d_2591_v69_
                nw404_[int(11)] = _dafny.SeqWithoutIsStrInference([d_2568_v49_ for d_2595_i4_ in range(default__.abs(905))])
                nw404_[int(12)] = d_2591_v69_
                nw404_[int(13)] = (((d_2591_v69_) + (d_2591_v69_)).set(default__.safeIndex(((d_2592_v70_)[(self).f11] if ((self).f11) in (d_2592_v70_) else d_2563_v44_), len((d_2591_v69_) + (d_2591_v69_))), d_2568_v49_)).set(default__.safeIndex((default__.fm68(globalState)).cardinality, len(((d_2591_v69_) + (d_2591_v69_)).set(default__.safeIndex(((d_2592_v70_)[(self).f11] if ((self).f11) in (d_2592_v70_) else d_2563_v44_), len((d_2591_v69_) + (d_2591_v69_))), d_2568_v49_))), d_2568_v49_)
                nw404_[int(14)] = (d_2591_v69_) + (d_2591_v69_)
                def iife211_(_pat_let56_0):
                    def iife212_(d_2596_dt__update__tmp_h0_):
                        def iife213_(_pat_let57_0):
                            def iife214_(d_2597_dt__update_hcf16_h0_):
                                def iife215_(_pat_let58_0):
                                    def iife216_(d_2598_dt__update_hcf15_h0_):
                                        return D1_DC5(d_2598_dt__update_hcf15_h0_, d_2597_dt__update_hcf16_h0_, (d_2596_dt__update__tmp_h0_).cf17)
                                    return iife216_(_pat_let58_0)
                                return iife215_(pat_let_tv40_)
                            return iife214_(_pat_let57_0)
                        return iife213_(pat_let_tv39_)
                    return iife212_(_pat_let56_0)
                nw404_[int(15)] = (iife211_(d_2593_v71_)).cf17
                d_2594_v72_ = nw404_
                index393_ = default__.safeIndex(40, (d_2594_v72_).length(0))
                (d_2594_v72_)[index393_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
                index394_ = default__.safeIndex(40, (d_2594_v72_).length(0))
                (d_2594_v72_)[index394_] = (d_2591_v69_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2599_i5_ in range(default__.abs(497))]))
                d_2600_v73_: _dafny.Map
                d_2600_v73_ = _dafny.Map({929: default__.fm26(globalState)})
                d_2601_v74_: _dafny.Map
                d_2601_v74_ = _dafny.Map({d_2600_v73_: _dafny.CodePoint('m')})
                d_2601_v74_ = (d_2601_v74_).set(d_2600_v73_, d_2568_v49_)
                d_2563_v44_ = (d_2564_v45_)[default__.safeIndex(d_2563_v44_, len(d_2564_v45_))]
                d_2602_v75_: _dafny.Seq
                d_2602_v75_ = _dafny.SeqWithoutIsStrInference([d_2562_v43_])
                d_2603_v76_: C2
                nw405_ = C2()
                nw405_.ctor__((self).f11, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2563_v44_ for d_2604_i6_ in range(default__.abs(903))])), d_2591_v69_, d_2512_v0_, d_2512_v0_, (d_2564_v45_)[default__.safeIndex(len(d_2602_v75_), len(d_2564_v45_))])
                d_2603_v76_ = nw405_
            elif True:
                rhs369_ = not(d_2512_v0_)
                rhs370_ = d_2512_v0_
                rhs371_ = default__.safeDivisionInt(default__.safeDivisionInt(p0, d_2563_v44_), len((default__.fm45(p2, globalState)).set(default__.safeIndex((0) - (p2), len(default__.fm45(p2, globalState))), d_2512_v0_)))
                rhs372_ = (p2) < ((self).f10)
                rhs373_ = d_2563_v44_
                d_2512_v0_ = rhs369_
                d_2562_v43_ = rhs370_
                d_2563_v44_ = rhs371_
                d_2562_v43_ = rhs372_
                d_2563_v44_ = rhs373_
                rhs374_ = d_2512_v0_
                rhs375_ = d_2562_v43_
                d_2562_v43_ = rhs374_
                d_2512_v0_ = rhs375_
                d_2605_v77_: _dafny.Seq
                d_2605_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awmgq"))
                d_2606_v78_: _dafny.Map
                d_2606_v78_ = _dafny.Map({d_2605_v77_: (p2) * (len(d_2605_v77_))})
                d_2606_v78_ = (d_2606_v78_).set((d_2605_v77_ if False else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecvrw"))), len(d_2565_v46_))
                d_2607_v79_: _dafny.Set
                d_2607_v79_ = _dafny.Set({d_2564_v45_})
                d_2608_v80_: _dafny.MultiSet
                d_2608_v80_ = _dafny.MultiSet([d_2512_v0_])
                d_2609_v81_: _dafny.Map
                d_2609_v81_ = _dafny.Map({True: len(_dafny.Map({d_2562_v43_: p1}))})
                d_2610_v82_: _dafny.Map
                d_2610_v82_ = _dafny.Map({p3: (self).f10})
                d_2611_v83_: _dafny.Seq
                d_2611_v83_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11])
                d_2612_v84_: _dafny.Seq
                d_2612_v84_ = _dafny.SeqWithoutIsStrInference([d_2610_v82_, d_2610_v82_])
                d_2613_v85_: _dafny.Array
                nw406_ = _dafny.Array(None, 13)
                nw406_[int(0)] = len(d_2607_v79_)
                nw406_[int(1)] = ((d_2608_v80_ if p1 else _dafny.MultiSet([p1]))).cardinality
                nw406_[int(2)] = default__.safeDivisionInt(((d_2609_v81_)[d_2562_v43_] if (d_2562_v43_) in (d_2609_v81_) else d_2563_v44_), default__.fm18(True, p1, d_2512_v0_, globalState))
                nw406_[int(3)] = (d_2563_v44_) * (p0)
                nw406_[int(4)] = (self).f10
                nw406_[int(5)] = p2
                nw406_[int(6)] = p3
                nw406_[int(7)] = p3
                nw406_[int(8)] = ((d_2610_v82_)[d_2563_v44_] if (d_2563_v44_) in (d_2610_v82_) else (0) - (len((D13_DC34(d_2611_v83_)).cf71)))
                nw406_[int(9)] = len(d_2612_v84_)
                nw406_[int(10)] = (d_2563_v44_) * ((0) - ((0) - (d_2563_v44_)))
                nw406_[int(11)] = len(_dafny.SeqWithoutIsStrInference([d_2568_v49_ for d_2614_i7_ in range(default__.abs(439))]))
                nw406_[int(12)] = p3
                d_2613_v85_ = nw406_
                d_2613_v85_ = d_2613_v85_
                d_2615_v86_: _dafny.Map
                d_2615_v86_ = _dafny.Map({True: True})
                d_2616_v87_: _dafny.Set
                d_2616_v87_ = _dafny.Set({d_2615_v86_})
                d_2616_v87_ = d_2616_v87_
            d_2617_v88_: _dafny.Array
            def lambda109_(d_2618_p2_):
                def lambda110_(d_2619_i8_):
                    return (d_2619_i8_) - (d_2618_p2_)

                return lambda110_

            init63_ = lambda109_(p2)
            nw407_ = _dafny.Array(None, 4)
            for i0_63_ in range(nw407_.length(0)):
                nw407_[i0_63_] = init63_(i0_63_)
            d_2617_v88_ = nw407_
            d_2620_v89_: _dafny.Seq
            d_2620_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfkca"))
            d_2621_v90_: _dafny.Seq
            d_2621_v90_ = _dafny.SeqWithoutIsStrInference([d_2512_v0_, p1])
            index395_ = default__.safeIndex(373, (d_2617_v88_).length(0))
            (d_2617_v88_)[index395_] = default__.fm55(len(d_2620_v89_), d_2621_v90_, p3, globalState)
            index396_ = default__.safeIndex(373, (d_2617_v88_).length(0))
            (d_2617_v88_)[index396_] = (self).f10
            if not((d_2562_v43_) or (p1)):
                index397_ = default__.safeIndex(373, (d_2617_v88_).length(0))
                (d_2617_v88_)[index397_] = ((d_2617_v88_)[default__.safeIndex(373, (d_2617_v88_).length(0))]) + (p0)
                d_2620_v89_ = d_2620_v89_
                d_2622_v91_: _dafny.Array
                def lambda111_(d_2623_v0_):
                    def lambda112_(d_2624_i9_):
                        return d_2623_v0_

                    return lambda112_

                init64_ = lambda111_(d_2512_v0_)
                nw408_ = _dafny.Array(None, 21)
                for i0_64_ in range(nw408_.length(0)):
                    nw408_[i0_64_] = init64_(i0_64_)
                d_2622_v91_ = nw408_
                d_2622_v91_ = d_2622_v91_
                d_2562_v43_ = ((self).f10) != (len((d_2564_v45_) + (d_2564_v45_)))
                nw409_ = _dafny.Array(False, 28)
                d_2622_v91_ = nw409_
            elif True:
                d_2625_v92_: _dafny.MultiSet
                d_2625_v92_ = _dafny.MultiSet([p1])
                d_2626_v93_: _dafny.Set
                d_2626_v93_ = _dafny.Set({d_2563_v44_, d_2563_v44_})
                index398_ = default__.safeIndex(373, (d_2617_v88_).length(0))
                rhs376_ = ((d_2625_v92_).cardinality) + (-681)
                rhs377_ = not((d_2620_v89_) <= ((d_2620_v89_).set(default__.safeIndex(len(d_2626_v93_), len(d_2620_v89_)), d_2568_v49_)))
                lhs211_ = d_2617_v88_
                lhs212_ = default__.safeIndex(373, (d_2617_v88_).length(0))
                lhs211_[lhs212_] = rhs376_
                d_2562_v43_ = rhs377_
                d_2627_v94_: _dafny.Map
                d_2627_v94_ = _dafny.Map({p1: d_2512_v0_})
                d_2627_v94_ = (d_2627_v94_).set(d_2562_v43_, p1)
                d_2512_v0_ = d_2512_v0_
                d_2621_v90_ = d_2621_v90_
                d_2628_v95_: D14
                d_2628_v95_ = D14_DC38(d_2562_v43_)
                d_2512_v0_ = (d_2628_v95_).cf74

    def m7(self, p0, globalState):
        d_2629_v0_: _dafny.Array
        nw410_ = _dafny.Array(int(0), 14)
        d_2629_v0_ = nw410_
        d_2630_v1_: _dafny.Map
        d_2630_v1_ = _dafny.Map({(self).f10: p0})
        d_2631_v2_: D11
        d_2631_v2_ = D11_DC30(p0, d_2629_v0_, len(d_2630_v1_), p0)
        source33_ = d_2631_v2_
        if source33_.is_DC28:
            d_2632___mcc_h0_ = source33_.cf52
            d_2633___mcc_h1_ = source33_.cf53
            d_2634___mcc_h2_ = source33_.cf54
            d_2635___mcc_h3_ = source33_.cf55
            d_2636___mcc_h4_ = source33_.cf56
            d_2637_cf56_ = d_2636___mcc_h4_
            d_2638_cf55_ = d_2635___mcc_h3_
            d_2639_cf54_ = d_2634___mcc_h2_
            d_2640_cf53_ = d_2633___mcc_h1_
            d_2641_cf52_ = d_2632___mcc_h0_
            d_2642_v3_: _dafny.MultiSet
            d_2642_v3_ = _dafny.MultiSet([d_2629_v0_])
            d_2639_cf54_ = (0) - ((((d_2642_v3_) | (d_2642_v3_)) - (d_2642_v3_)).cardinality)
            d_2643_v4_: _dafny.Set
            d_2643_v4_ = _dafny.Set({d_2640_cf53_, 954})
            d_2644_v5_: _dafny.Seq
            d_2644_v5_ = _dafny.SeqWithoutIsStrInference([d_2638_cf55_])
            d_2639_cf54_ = default__.fm55(((0) - (len(d_2643_v4_))) * (d_2640_cf53_), d_2644_v5_, d_2639_cf54_, globalState)
            if d_2637_cf56_:
                d_2645_v6_: C9
                nw411_ = C9()
                nw411_.ctor__((D11_DC29(d_2639_cf54_, (self).f11, d_2641_cf52_, d_2637_cf56_)).cf58)
                d_2645_v6_ = nw411_
                d_2646_v7_: D9
                d_2646_v7_ = D9_DC24(d_2640_cf53_)
                d_2647_v8_: _dafny.Seq
                d_2647_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ay"))
                d_2648_v9_: _dafny.Map
                d_2648_v9_ = _dafny.Map({d_2646_v7_: (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2649_i0_ in range(default__.abs(129))])) < (d_2647_v8_)})
                d_2648_v9_ = (d_2648_v9_).set(d_2646_v7_, d_2637_cf56_)
                d_2641_cf52_ = ((-607 if d_2638_cf55_ else len(d_2647_v8_))) >= (len(d_2647_v8_))
                d_2650_v10_: _dafny.Array
                nw412_ = _dafny.Array(None, 17)
                d_2650_v10_ = nw412_
                d_2651_v11_: _dafny.MultiSet
                d_2651_v11_ = _dafny.MultiSet([d_2639_cf54_, d_2640_cf53_])
                d_2652_v12_: D1
                d_2652_v12_ = D1_DC4(d_2641_cf52_, d_2639_cf54_, d_2639_cf54_)
                d_2653_v13_: C4
                nw413_ = C4()
                nw413_.ctor__(d_2639_cf54_, ((d_2630_v1_)[(d_2651_v11_).cardinality] if ((d_2651_v11_).cardinality) in (d_2630_v1_) else not(p0)), d_2647_v8_, d_2641_cf52_, p0, (d_2652_v12_).cf14, d_2640_cf53_, p0)
                d_2653_v13_ = nw413_
                index399_ = default__.safeIndex(673, (d_2650_v10_).length(0))
                (d_2650_v10_)[index399_] = d_2653_v13_
                index400_ = default__.safeIndex(673, (d_2650_v10_).length(0))
                (d_2650_v10_)[index400_] = d_2653_v13_
                d_2638_cf55_ = False
            elif True:
                d_2638_cf55_ = d_2637_cf56_
                d_2638_cf55_ = not(False)
                d_2654_v14_: str
                d_2654_v14_ = _dafny.CodePoint('f')
                d_2655_v15_: _dafny.Seq
                d_2655_v15_ = _dafny.SeqWithoutIsStrInference([d_2654_v14_])
                d_2656_v16_: D0
                d_2656_v16_ = D0_DC0(d_2641_cf52_, d_2640_cf53_, True, (self).f10, (d_2644_v5_)[default__.safeIndex(d_2639_cf54_, len(d_2644_v5_))])
                d_2657_v17_: C7
                nw414_ = C7()
                nw414_.ctor__((0) - (d_2640_cf53_), (d_2655_v15_) <= (d_2655_v15_), p0, (self).f10, default__.fm32(d_2656_v16_, globalState), d_2638_cf55_, d_2655_v15_, d_2641_cf52_)
                d_2657_v17_ = nw414_
                d_2641_cf52_ = (len(_dafny.Map({(d_2657_v17_).f19: d_2644_v5_}))) >= (default__.fm18(default__.fm27(_dafny.CodePoint('e'), globalState), d_2641_cf52_, p0, globalState))
                d_2658_v18_: _dafny.Set
                d_2658_v18_ = _dafny.Set({d_2654_v14_})
                d_2659_v19_: _dafny.Map
                d_2659_v19_ = _dafny.Map({d_2654_v14_: d_2638_cf55_})
                def iife217_():
                    coll99_ = _dafny.Set()
                    compr_99_: str
                    for compr_99_ in (d_2659_v19_).keys.Elements:
                        d_2660_v20_: str = compr_99_
                        if (d_2660_v20_) in (d_2659_v19_):
                            coll99_ = coll99_.union(_dafny.Set([d_2660_v20_]))
                    return _dafny.Set(coll99_)
                d_2658_v18_ = (d_2658_v18_) - (iife217_()
                )
            d_2638_cf55_ = d_2638_cf55_
        elif source33_.is_DC29:
            d_2661___mcc_h5_ = source33_.cf57
            d_2662___mcc_h6_ = source33_.cf58
            d_2663___mcc_h7_ = source33_.cf59
            d_2664___mcc_h8_ = source33_.cf60
            d_2665_cf60_ = d_2664___mcc_h8_
            d_2666_cf59_ = d_2663___mcc_h7_
            d_2667_cf58_ = d_2662___mcc_h6_
            d_2668_cf57_ = d_2661___mcc_h5_
            d_2669_v21_: _dafny.Map
            d_2669_v21_ = _dafny.Map({p0: d_2665_cf60_})
            if ((d_2669_v21_)[d_2665_cf60_] if (d_2665_cf60_) in (d_2669_v21_) else p0):
                d_2666_cf59_ = ((d_2630_v1_)[d_2668_cf57_] if (d_2668_cf57_) in (d_2630_v1_) else d_2666_cf59_)
                d_2665_cf60_ = p0
                d_2670_v22_: _dafny.Seq
                d_2670_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukio"))
                d_2671_v23_: _dafny.Map
                d_2671_v23_ = _dafny.Map({p0: len(d_2670_v22_)})
                d_2672_v24_: C6
                nw415_ = C6()
                nw415_.ctor__(len(d_2670_v22_), d_2666_cf59_, d_2665_cf60_, default__.fm26(globalState), 611, (d_2666_cf59_) not in (d_2671_v23_))
                d_2672_v24_ = nw415_
                d_2673_v25_: C1
                nw416_ = C1()
                nw416_.ctor__((self).f10, not(p0), d_2665_cf60_, d_2672_v24_.f21)
                d_2673_v25_ = nw416_
                d_2674_v26_: _dafny.Array
                nw417_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_2674_v26_ = nw417_
                index401_ = default__.safeIndex(857, (d_2674_v26_).length(0))
                (d_2674_v26_)[index401_] = (self).f11
                index402_ = default__.safeIndex(857, (d_2674_v26_).length(0))
                (d_2674_v26_)[index402_] = (self).f11
            elif True:
                index403_ = default__.safeIndex(563, ((self).f11).length(0))
                ((self).f11)[index403_] = p0
                d_2675_v27_: _dafny.Seq
                d_2675_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eyli"))
                d_2676_v28_: _dafny.Map
                d_2676_v28_ = _dafny.Map({(self).f10: _dafny.CodePoint('i')})
                d_2677_v29_: _dafny.Seq
                d_2677_v29_ = _dafny.SeqWithoutIsStrInference([p0, d_2665_cf60_, not(p0), (len(d_2675_v27_)) > (len(d_2676_v28_))])
                index404_ = default__.safeIndex(563, ((self).f11).length(0))
                rhs378_ = not ((d_2665_cf60_) == (d_2666_cf59_)) or (d_2666_cf59_)
                rhs379_ = (d_2677_v29_) + (d_2677_v29_)
                rhs380_ = d_2666_cf59_
                rhs381_ = d_2666_cf59_
                lhs213_ = (self).f11
                lhs214_ = default__.safeIndex(563, ((self).f11).length(0))
                lhs213_[lhs214_] = rhs378_
                d_2677_v29_ = rhs379_
                d_2665_cf60_ = rhs380_
                d_2666_cf59_ = rhs381_
                d_2675_v27_ = d_2675_v27_
                d_2678_v30_: _dafny.MultiSet
                d_2678_v30_ = _dafny.MultiSet([d_2668_cf57_, len((d_2677_v29_).set(default__.safeIndex((self).f10, len(d_2677_v29_)), default__.fm27(_dafny.CodePoint('w'), globalState))), 537, (self).f10])
                d_2679_v31_: _dafny.Map
                d_2679_v31_ = _dafny.Map({(d_2677_v29_).set(default__.safeIndex((d_2678_v30_).cardinality, len(d_2677_v29_)), d_2666_cf59_): (not(d_2665_cf60_)) or (((self).f11)[default__.safeIndex(563, ((self).f11).length(0))])})
                d_2679_v31_ = (d_2679_v31_).set(((d_2677_v29_ if d_2666_cf59_ else d_2677_v29_)).set(default__.safeIndex((d_2678_v30_).cardinality, len((d_2677_v29_ if d_2666_cf59_ else d_2677_v29_))), d_2666_cf59_), (p0 if p0 else default__.fm27(_dafny.CodePoint('j'), globalState)))
                d_2680_v32_: _dafny.Map
                d_2680_v32_ = _dafny.Map({d_2667_cf58_: (self).fm15(globalState)})
                d_2668_cf57_ = len(d_2680_v32_)
                d_2666_cf59_ = d_2666_cf59_
            index405_ = default__.safeIndex(954, (d_2667_cf58_).length(0))
            (d_2667_cf58_)[index405_] = p0
            d_2681_v33_: _dafny.MultiSet
            d_2681_v33_ = _dafny.MultiSet([p0])
            d_2682_v34_: C6
            nw418_ = C6()
            nw418_.ctor__(d_2668_cf57_, d_2666_cf59_, d_2665_cf60_, d_2668_cf57_, (d_2681_v33_).cardinality, d_2666_cf59_)
            d_2682_v34_ = nw418_
            d_2683_v35_: _dafny.Map
            d_2683_v35_ = _dafny.Map({d_2668_cf57_: d_2682_v34_})
            index406_ = default__.safeIndex(954, (d_2667_cf58_).length(0))
            (d_2667_cf58_)[index406_] = (d_2683_v35_) == (d_2683_v35_)
            index407_ = default__.safeIndex(132, (d_2629_v0_).length(0))
            (d_2629_v0_)[index407_] = d_2668_cf57_
            index408_ = default__.safeIndex(132, (d_2629_v0_).length(0))
            (d_2629_v0_)[index408_] = (len(_dafny.SeqWithoutIsStrInference([(0) - (d_2668_cf57_) for d_2684_i1_ in range(default__.abs(-57))]))) - (((self).f10) * (d_2668_cf57_))
            d_2685_v36_: _dafny.Map
            d_2685_v36_ = _dafny.Map({p0: d_2667_cf58_})
            d_2686_v38_: _dafny.Map
            def iife218_():
                coll100_ = _dafny.Set()
                compr_100_: int
                for compr_100_ in _dafny.IntegerRange(442, 13):
                    d_2687_v37_: int = compr_100_
                    if ((442) <= (d_2687_v37_)) and ((d_2687_v37_) < (13)):
                        coll100_ = coll100_.union(_dafny.Set([(d_2687_v37_) - (d_2682_v34_.f21)]))
                return _dafny.Set(coll100_)
            d_2686_v38_ = _dafny.Map({True: len(iife218_()
            )})
            d_2688_v39_: _dafny.MultiSet
            d_2688_v39_ = _dafny.MultiSet([d_2668_cf57_])
            d_2689_v40_: _dafny.Seq
            d_2689_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
            d_2690_v41_: _dafny.Set
            d_2690_v41_ = _dafny.Set({d_2666_cf59_})
            d_2691_v42_: _dafny.Seq
            d_2691_v42_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(d_2690_v41_)]), d_2688_v39_])
            d_2692_v43_: _dafny.Array
            nw419_ = _dafny.Array(None, 7)
            nw419_[int(0)] = -57
            nw419_[int(1)] = (d_2682_v34_.f21) * (d_2668_cf57_)
            nw419_[int(2)] = default__.safeDivisionInt((0) - (len(d_2685_v36_)), ((d_2686_v38_)[d_2665_cf60_] if (d_2665_cf60_) in (d_2686_v38_) else ((d_2686_v38_)[(d_2667_cf58_)[default__.safeIndex(954, (d_2667_cf58_).length(0))]] if ((d_2667_cf58_)[default__.safeIndex(954, (d_2667_cf58_).length(0))]) in (d_2686_v38_) else (d_2629_v0_)[default__.safeIndex(132, (d_2629_v0_).length(0))])))
            nw419_[int(3)] = default__.fm26(globalState)
            nw419_[int(4)] = (d_2629_v0_)[default__.safeIndex(132, (d_2629_v0_).length(0))]
            nw419_[int(5)] = (d_2688_v39_).cardinality
            nw419_[int(6)] = default__.fm49(d_2689_v40_, ((d_2691_v42_)[default__.safeIndex(d_2682_v34_.f21, len(d_2691_v42_))]).cardinality, d_2666_cf59_, (self).fm15(globalState), globalState)
            d_2692_v43_ = nw419_
            d_2693_v44_: _dafny.Seq
            d_2693_v44_ = _dafny.SeqWithoutIsStrInference([d_2692_v43_, d_2692_v43_])
            d_2692_v43_ = (d_2693_v44_)[default__.safeIndex(754, len(d_2693_v44_))]
        elif source33_.is_DC30:
            d_2694___mcc_h9_ = source33_.cf61
            d_2695___mcc_h10_ = source33_.cf62
            d_2696___mcc_h11_ = source33_.cf63
            d_2697___mcc_h12_ = source33_.cf64
            d_2698_cf64_ = d_2697___mcc_h12_
            d_2699_cf63_ = d_2696___mcc_h11_
            d_2700_cf62_ = d_2695___mcc_h10_
            d_2701_cf61_ = d_2694___mcc_h9_
            d_2699_cf63_ = d_2699_cf63_
            d_2699_cf63_ = (default__.safeModuloInt((0) - ((0) - ((0) - ((self).f10))), d_2699_cf63_)) * ((0) - (len(default__.fm60(d_2701_cf61_, p0, d_2698_cf64_, d_2699_cf63_, globalState))))
            d_2699_cf63_ = ((d_2699_cf63_) + (d_2699_cf63_)) + ((self).f10)
            d_2702_v45_: _dafny.Seq
            d_2702_v45_ = _dafny.SeqWithoutIsStrInference([(d_2701_cf61_) or (p0), d_2701_cf61_])
            d_2702_v45_ = default__.fm21(globalState)
        elif source33_.is_DC27:
            d_2703___mcc_h13_ = source33_.cf51
            d_2704_cf51_ = d_2703___mcc_h13_
            d_2705_v46_: bool
            d_2705_v46_ = False
            d_2706_v47_: _dafny.Map
            d_2706_v47_ = _dafny.Map({d_2705_v46_: (self).f10})
            d_2705_v46_ = (p0) or ((len(d_2706_v47_)) > ((self).f10))
            d_2707_v48_: _dafny.Map
            d_2707_v48_ = _dafny.Map({(self).f10: d_2629_v0_})
            d_2708_v49_: D9
            d_2708_v49_ = D9_DC21(d_2707_v48_)
            pat_let_tv41_ = d_2707_v48_
            pat_let_tv42_ = d_2707_v48_
            d_2709_v50_: _dafny.Array
            nw420_ = _dafny.Array(None, 19)
            nw420_[int(0)] = d_2708_v49_
            nw420_[int(1)] = d_2708_v49_
            nw420_[int(2)] = d_2708_v49_
            nw420_[int(3)] = d_2708_v49_
            def iife219_(_pat_let59_0):
                def iife220_(d_2710_dt__update__tmp_h0_):
                    def iife221_(_pat_let60_0):
                        def iife222_(d_2711_dt__update_hcf43_h0_):
                            return D9_DC21(d_2711_dt__update_hcf43_h0_)
                        return iife222_(_pat_let60_0)
                    return iife221_(pat_let_tv41_)
                return iife220_(_pat_let59_0)
            nw420_[int(4)] = iife219_(d_2708_v49_)
            nw420_[int(5)] = d_2708_v49_
            nw420_[int(6)] = d_2708_v49_
            nw420_[int(7)] = d_2708_v49_
            nw420_[int(8)] = d_2708_v49_
            nw420_[int(9)] = d_2708_v49_
            nw420_[int(10)] = d_2708_v49_
            nw420_[int(11)] = d_2708_v49_
            nw420_[int(12)] = D9_DC21(d_2707_v48_)
            def iife223_(_pat_let61_0):
                def iife224_(d_2712_dt__update__tmp_h1_):
                    def iife225_(_pat_let62_0):
                        def iife226_(d_2713_dt__update_hcf43_h1_):
                            return D9_DC21(d_2713_dt__update_hcf43_h1_)
                        return iife226_(_pat_let62_0)
                    return iife225_(pat_let_tv42_)
                return iife224_(_pat_let61_0)
            nw420_[int(13)] = iife223_(d_2708_v49_)
            nw420_[int(14)] = d_2708_v49_
            nw420_[int(15)] = d_2708_v49_
            nw420_[int(16)] = d_2708_v49_
            nw420_[int(17)] = d_2708_v49_
            nw420_[int(18)] = d_2708_v49_
            d_2709_v50_ = nw420_
            d_2714_v51_: _dafny.Map
            d_2714_v51_ = _dafny.Map({d_2705_v46_: (d_2709_v50_)})
            d_2714_v51_ = d_2714_v51_
            d_2715_v52_: bool
            out25_: bool
            out25_ = default__.m0(False, (self).f10, globalState)
            d_2715_v52_ = out25_
            d_2716_v53_: _dafny.Seq
            d_2716_v53_ = _dafny.SeqWithoutIsStrInference([True])
            d_2717_v54_: _dafny.Seq
            d_2717_v54_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsirk"))), (self).f10, (self).f10])
            d_2718_v55_: _dafny.Map
            d_2718_v55_ = _dafny.Map({p0: d_2715_v52_})
            d_2719_v56_: _dafny.Map
            d_2719_v56_ = _dafny.Map({p0: ((d_2718_v55_)[d_2715_v52_] if (d_2715_v52_) in (d_2718_v55_) else d_2705_v46_)})
            d_2720_v57_: D19
            d_2720_v57_ = D19_DC53((self).f10, (self).f10, ((d_2716_v53_)[default__.safeIndex((d_2717_v54_)[default__.safeIndex((self).f10, len(d_2717_v54_))], len(d_2716_v53_))]) not in (d_2718_v55_), d_2715_v52_)
            source34_ = d_2720_v57_
            if source34_.is_DC53:
                d_2721___mcc_h15_ = source34_.cf94
                d_2722___mcc_h16_ = source34_.cf95
                d_2723___mcc_h17_ = source34_.cf96
                d_2724___mcc_h18_ = source34_.cf97
                d_2725_cf97_ = d_2724___mcc_h18_
                d_2726_cf96_ = d_2723___mcc_h17_
                d_2727_cf95_ = d_2722___mcc_h16_
                d_2728_cf94_ = d_2721___mcc_h15_
                d_2729_v58_: _dafny.Seq
                d_2729_v58_ = _dafny.SeqWithoutIsStrInference([d_2706_v47_])
                index409_ = default__.safeIndex(226, ((self).f11).length(0))
                ((self).f11)[index409_] = (d_2706_v47_) not in (d_2729_v58_)
                index410_ = default__.safeIndex(226, ((self).f11).length(0))
                ((self).f11)[index410_] = not((not(p0) if d_2726_cf96_ else p0))
                d_2728_cf94_ = default__.safeModuloInt(d_2728_cf94_, (self).f10)
                d_2730_v59_: _dafny.Seq
                d_2730_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eygk"))
                d_2728_cf94_ = (default__.safeModuloInt(len(d_2730_v59_), len(d_2718_v55_))) - (632)
                d_2728_cf94_ = (self).f10
            elif True:
                d_2731___mcc_h19_ = source34_.cf93
                d_2732_cf93_ = d_2731___mcc_h19_
                d_2733_v60_: C5
                nw421_ = C5()
                nw421_.ctor__()
                d_2733_v60_ = nw421_
                d_2734_v61_: _dafny.Seq
                d_2734_v61_ = _dafny.SeqWithoutIsStrInference([d_2629_v0_])
                rhs382_ = d_2733_v60_
                rhs383_ = not(not((d_2734_v61_) != (d_2734_v61_)))
                d_2733_v60_ = rhs382_
                d_2715_v52_ = rhs383_
                d_2735_v62_: C8
                nw422_ = C8()
                nw422_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcebatx")), (self).f10)
                d_2735_v62_ = nw422_
                d_2736_v63_: _dafny.Array
                nw423_ = _dafny.Array(False, 1)
                d_2736_v63_ = nw423_
                d_2737_v64_: int
                d_2737_v64_ = 222
                d_2738_v65_: _dafny.Map
                d_2738_v65_ = _dafny.Map({d_2735_v62_.f13: (d_2735_v62_).f14})
                d_2739_v66_: str
                d_2739_v66_ = _dafny.CodePoint('e')
                d_2740_v67_: _dafny.Map
                d_2740_v67_ = _dafny.Map({p0: d_2739_v66_})
                rhs384_ = p0
                rhs385_ = (self).f11
                rhs386_ = ((((d_2706_v47_)[False] if (False) in (d_2706_v47_) else d_2737_v64_)) * (d_2737_v64_)) * (((d_2738_v65_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "niodqnk"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "niodqnk"))) in (d_2738_v65_) else len(d_2740_v67_)))
                rhs387_ = default__.safeDivisionInt((d_2735_v62_).f14, default__.fm32(D0_DC0(True, (self).f10, d_2715_v52_, (d_2735_v62_).f14, p0), globalState))
                d_2715_v52_ = rhs384_
                d_2736_v63_ = rhs385_
                d_2737_v64_ = rhs386_
                d_2737_v64_ = rhs387_
                index411_ = default__.safeIndex(975, (d_2629_v0_).length(0))
                (d_2629_v0_)[index411_] = ((d_2735_v62_).f14) + ((0) - (d_2737_v64_))
                index412_ = default__.safeIndex(975, (d_2629_v0_).length(0))
                (d_2629_v0_)[index412_] = (self).f10
        elif True:
            d_2741___mcc_h14_ = source33_.cf65
            d_2742_cf65_ = d_2741___mcc_h14_
            d_2743_v68_: int
            d_2743_v68_ = 770
            d_2743_v68_ = (self).f10
            d_2744_v69_: bool
            d_2744_v69_ = True
            d_2744_v69_ = not(p0)
            d_2745_v70_: _dafny.MultiSet
            d_2745_v70_ = _dafny.MultiSet([(self).f10])
            d_2746_v71_: str
            d_2746_v71_ = _dafny.CodePoint('p')
            d_2747_v72_: _dafny.Seq
            d_2747_v72_ = _dafny.SeqWithoutIsStrInference([d_2746_v71_])
            d_2748_v73_: _dafny.Seq
            d_2748_v73_ = _dafny.SeqWithoutIsStrInference([d_2747_v72_, d_2747_v72_])
            d_2749_v74_: _dafny.Map
            d_2749_v74_ = _dafny.Map({p0: p0})
            d_2750_v75_: C2
            nw424_ = C2()
            nw424_.ctor__(((self).f11 if d_2744_v69_ else (self).f11), d_2745_v70_, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2751_i2_ in range(default__.abs(79))]) if p0 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uod"))), (d_2743_v68_) not in (d_2630_v1_), ((d_2747_v72_).set(default__.safeIndex(len(d_2747_v72_), len(d_2747_v72_)), d_2746_v71_)) in (d_2748_v73_), len((d_2749_v74_) | (d_2749_v74_)))
            d_2750_v75_ = nw424_
            d_2744_v69_ = d_2744_v69_
        hi8_ = (self).f10
        for d_2752_i3_ in range((self).f10, hi8_):
            d_2753_v76_: int
            d_2753_v76_ = 550
            d_2754_v77_: _dafny.Map
            d_2754_v77_ = _dafny.Map({(self).fm15(globalState): default__.fm68(globalState)})
            d_2755_v78_: _dafny.Set
            d_2755_v78_ = _dafny.Set({d_2753_v76_})
            d_2753_v76_ = (0) - (len((d_2754_v77_) | (default__.fm78(False, p0, d_2755_v78_, p0, globalState))))
            d_2756_v79_: str
            d_2756_v79_ = _dafny.CodePoint('o')
            d_2753_v76_ = ((len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ten"))).set(default__.safeIndex((self).f10, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ten")))), d_2756_v79_))) + (275)) + (d_2752_i3_)
            d_2757_v80_: _dafny.Seq
            d_2757_v80_ = _dafny.SeqWithoutIsStrInference([d_2752_i3_, d_2752_i3_])
            d_2758_v81_: _dafny.Seq
            d_2758_v81_ = _dafny.SeqWithoutIsStrInference([d_2757_v80_])
            d_2758_v81_ = d_2758_v81_
            d_2756_v79_ = d_2756_v79_
        d_2759_v82_: _dafny.Array
        nw425_ = _dafny.Array(_dafny.MultiSet({}), 5)
        d_2759_v82_ = nw425_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_2759_v82_).length(0)):
            d_2760_i4_: int = guard_loop_8_
            if (True) and (((0) <= (d_2760_i4_)) and ((d_2760_i4_) < ((d_2759_v82_).length(0)))):
                (d_2759_v82_)[(d_2760_i4_)] = ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, p0, not(False)]))) | (_dafny.MultiSet([p0, p0]))).intersection((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0]))) - (_dafny.MultiSet([p0, p0])))
        d_2761_v83_: bool
        d_2761_v83_ = False
        d_2762_v84_: _dafny.MultiSet
        d_2762_v84_ = _dafny.MultiSet([d_2761_v83_, d_2761_v83_])
        d_2763_v85_: _dafny.Seq
        d_2763_v85_ = _dafny.SeqWithoutIsStrInference([d_2761_v83_])
        d_2761_v83_ = (_dafny.MultiSet(d_2763_v85_)).ispropersubset(d_2762_v84_)
        d_2764_v86_: _dafny.Array
        nw426_ = _dafny.Array(_dafny.Map({}), 23)
        d_2764_v86_ = nw426_
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, (d_2764_v86_).length(0)):
            d_2765_i5_: int = guard_loop_9_
            if (True) and (((0) <= (d_2765_i5_)) and ((d_2765_i5_) < ((d_2764_v86_).length(0)))):
                def iife227_():
                    coll101_ = _dafny.Map()
                    compr_101_: int
                    for compr_101_ in _dafny.IntegerRange(256, 377):
                        d_2766_v87_: int = compr_101_
                        if ((256) <= (d_2766_v87_)) and ((d_2766_v87_) < (377)):
                            coll101_[(d_2766_v87_) - ((self).f10)] = (self).f10
                    return _dafny.Map(coll101_)
                (d_2764_v86_)[(d_2765_i5_)] = (((_dafny.Map({(self).f10: len(_dafny.Map({d_2761_v83_: p0}))}))) | (_dafny.Map({len(_dafny.Set({d_2761_v83_})): (self).f10}))) | ((iife227_()
                ) | (_dafny.Map({(self).f10: 522})))
        d_2767_v88_: D17
        d_2767_v88_ = D17_DC46()
        source35_ = d_2767_v88_
        if source35_.is_DC45:
            d_2768_v89_: int
            d_2768_v89_ = 194
            d_2768_v89_ = d_2768_v89_
            d_2769_v90_: _dafny.MultiSet
            d_2769_v90_ = _dafny.MultiSet([(d_2768_v89_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "myjohj"))))])
            d_2770_v91_: _dafny.Array
            nw427_ = _dafny.Array(None, 12)
            nw427_[int(0)] = (self).f11
            nw427_[int(1)] = (self).f11
            nw427_[int(2)] = (self).f11
            nw427_[int(3)] = (self).f11
            nw427_[int(4)] = (self).f11
            nw427_[int(5)] = (self).f11
            nw427_[int(6)] = (self).f11
            nw427_[int(7)] = (self).f11
            nw427_[int(8)] = (self).f11
            nw427_[int(9)] = (self).f11
            nw427_[int(10)] = (self).f11
            nw427_[int(11)] = (self).f11
            d_2770_v91_ = nw427_
            index413_ = default__.safeIndex(306, (d_2770_v91_).length(0))
            (d_2770_v91_)[index413_] = (self).f11
            index414_ = default__.safeIndex(306, (d_2770_v91_).length(0))
            rhs388_ = d_2769_v90_
            rhs389_ = (self).f10
            rhs390_ = (self).f11
            rhs391_ = (self).f10
            lhs215_ = d_2770_v91_
            lhs216_ = default__.safeIndex(306, (d_2770_v91_).length(0))
            d_2769_v90_ = rhs388_
            d_2768_v89_ = rhs389_
            lhs215_[lhs216_] = rhs390_
            d_2768_v89_ = rhs391_
            arr0_ = (d_2770_v91_)[default__.safeIndex(306, (d_2770_v91_).length(0))]
            index415_ = default__.safeIndex(533, ((d_2770_v91_)[default__.safeIndex(306, (d_2770_v91_).length(0))]).length(0))
            arr0_[index415_] = not(p0)
            arr1_ = (d_2770_v91_)[default__.safeIndex(306, (d_2770_v91_).length(0))]
            index416_ = default__.safeIndex(533, ((d_2770_v91_)[default__.safeIndex(306, (d_2770_v91_).length(0))]).length(0))
            arr1_[index416_] = p0
            d_2768_v89_ = d_2768_v89_
        elif source35_.is_DC46:
            d_2763_v85_ = ((_dafny.SeqWithoutIsStrInference([not(p0), p0])) + (d_2763_v85_)) + (d_2763_v85_)
            index417_ = default__.safeIndex(694, ((self).f11).length(0))
            ((self).f11)[index417_] = True
            index418_ = default__.safeIndex(694, ((self).f11).length(0))
            ((self).f11)[index418_] = d_2761_v83_
            d_2771_v92_: _dafny.Seq
            d_2771_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kfb"))
            d_2772_v93_: _dafny.Map
            d_2772_v93_ = _dafny.Map({(self).f10: d_2771_v92_})
            d_2630_v1_ = (d_2630_v1_).set(((self).f10) + (len(d_2772_v93_)), p0)
            d_2773_v94_: str
            d_2773_v94_ = _dafny.CodePoint('m')
            d_2773_v94_ = ((d_2773_v94_ if ((self).f11)[default__.safeIndex(694, ((self).f11).length(0))] else d_2773_v94_) if (p0) == (((self).f11)[default__.safeIndex(694, ((self).f11).length(0))]) else d_2773_v94_)
        elif source35_.is_DC47:
            d_2774___mcc_h20_ = source35_.cf83
            d_2775___mcc_h21_ = source35_.cf84
            d_2776___mcc_h22_ = source35_.cf85
            d_2777___mcc_h23_ = source35_.cf86
            d_2778___mcc_h24_ = source35_.cf87
            d_2779_cf87_ = d_2778___mcc_h24_
            d_2780_cf86_ = d_2777___mcc_h23_
            d_2781_cf85_ = d_2776___mcc_h22_
            d_2782_cf84_ = d_2775___mcc_h21_
            d_2783_cf83_ = d_2774___mcc_h20_
            d_2779_cf87_ = d_2780_cf86_
            d_2784_v95_: _dafny.Map
            d_2784_v95_ = _dafny.Map({d_2782_cf84_: d_2780_cf86_})
            d_2784_v95_ = (d_2784_v95_).set(d_2761_v83_, ((self).f10) + (d_2779_cf87_))
            d_2785_v96_: D4
            d_2785_v96_ = D4_DC10(True, d_2782_cf84_, d_2763_v85_)
            d_2786_v97_: D4
            d_2786_v97_ = D4_DC11(d_2785_v96_)
            d_2787_v98_: _dafny.Seq
            d_2787_v98_ = _dafny.SeqWithoutIsStrInference([d_2786_v97_])
            d_2788_v99_: str
            d_2788_v99_ = _dafny.CodePoint('s')
            index419_ = default__.safeIndex(657, ((self).f11).length(0))
            ((self).f11)[index419_] = (d_2763_v85_)[default__.safeIndex(d_2780_cf86_, len(d_2763_v85_))]
            d_2789_v100_: _dafny.Set
            d_2789_v100_ = _dafny.Set({d_2783_cf83_})
            d_2790_v101_: _dafny.Seq
            d_2790_v101_ = _dafny.SeqWithoutIsStrInference([d_2779_cf87_, 877, len(d_2789_v100_), d_2781_cf85_, (self).f10])
            index420_ = default__.safeIndex(657, ((self).f11).length(0))
            rhs392_ = (d_2787_v98_) + (d_2787_v98_)
            rhs393_ = d_2763_v85_
            rhs394_ = d_2788_v99_
            rhs395_ = d_2782_cf84_
            rhs396_ = ((d_2790_v101_)[default__.safeIndex(d_2780_cf86_, len(d_2790_v101_))]) > (d_2780_cf86_)
            lhs217_ = (self).f11
            lhs218_ = default__.safeIndex(657, ((self).f11).length(0))
            d_2787_v98_ = rhs392_
            d_2763_v85_ = rhs393_
            d_2788_v99_ = rhs394_
            d_2761_v83_ = rhs395_
            lhs217_[lhs218_] = rhs396_
            d_2791_v102_: _dafny.Seq
            d_2791_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhkwkgnpv"))
            d_2791_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aeqlgxow"))
        elif source35_.is_DC44:
            d_2792___mcc_h25_ = source35_.cf82
            d_2793_cf82_ = d_2792___mcc_h25_
            d_2761_v83_ = d_2761_v83_
            d_2794_v103_: _dafny.Set
            d_2794_v103_ = _dafny.Set({(self).f10})
            d_2795_v104_: _dafny.Seq
            d_2795_v104_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krsasaeg"))
            d_2761_v83_ = (default__.fm52(not(d_2761_v83_), d_2794_v103_, (self).f10, p0, globalState)) == (d_2795_v104_)
            d_2796_v105_: C9
            nw428_ = C9()
            nw428_.ctor__(((self).f11 if d_2761_v83_ else (self).f11))
            d_2796_v105_ = nw428_
            index421_ = default__.safeIndex(319, ((self).f11).length(0))
            ((self).f11)[index421_] = p0
            d_2797_v106_: str
            d_2797_v106_ = _dafny.CodePoint('s')
            d_2798_v107_: _dafny.Map
            d_2798_v107_ = _dafny.Map({d_2761_v83_: default__.fm27(d_2797_v106_, globalState)})
            index422_ = default__.safeIndex(319, ((self).f11).length(0))
            ((self).f11)[index422_] = ((d_2798_v107_)[d_2761_v83_] if (d_2761_v83_) in (d_2798_v107_) else (d_2761_v83_) or (d_2761_v83_))
        elif True:
            d_2799___mcc_h26_ = source35_.cf88
            d_2800_cf88_ = d_2799___mcc_h26_
            if not(not(not(True))):
                d_2801_v108_: _dafny.Seq
                d_2801_v108_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eynw"))
                d_2801_v108_ = d_2801_v108_
                d_2802_v109_: _dafny.Map
                d_2802_v109_ = _dafny.Map({(self).f10: (self).f10})
                d_2803_v110_: _dafny.Set
                d_2803_v110_ = _dafny.Set({(self).fm15(globalState)})
                d_2761_v83_ = not((default__.safeDivisionInt(((d_2802_v109_)[(self).f10] if ((self).f10) in (d_2802_v109_) else (self).f10), (self).f10)) == (len(d_2803_v110_)))
                index423_ = default__.safeIndex(704, ((self).f11).length(0))
                ((self).f11)[index423_] = d_2761_v83_
                index424_ = default__.safeIndex(704, ((self).f11).length(0))
                ((self).f11)[index424_] = d_2761_v83_
                d_2804_v111_: int
                d_2804_v111_ = 890
                d_2805_v112_: _dafny.MultiSet
                d_2805_v112_ = _dafny.MultiSet([D11_DC30(((self).f11)[default__.safeIndex(704, ((self).f11).length(0))], d_2629_v0_, len(d_2801_v108_), ((self).f11)[default__.safeIndex(704, ((self).f11).length(0))]), d_2631_v2_, d_2631_v2_, d_2631_v2_, d_2631_v2_])
                d_2806_v113_: _dafny.MultiSet
                d_2806_v113_ = _dafny.MultiSet([(d_2805_v112_).cardinality])
                d_2804_v111_ = default__.safeModuloInt(((d_2806_v113_)[d_2804_v111_] if (d_2804_v111_) in (d_2806_v113_) else 692), d_2804_v111_)
                index425_ = default__.safeIndex(704, ((self).f11).length(0))
                rhs397_ = d_2804_v111_
                rhs398_ = ((self).f10) != ((self).f10)
                lhs219_ = (self).f11
                lhs220_ = default__.safeIndex(704, ((self).f11).length(0))
                d_2804_v111_ = rhs397_
                lhs219_[lhs220_] = rhs398_
            elif True:
                d_2807_v114_: _dafny.Array
                nw429_ = _dafny.Array(None, 7)
                nw429_[int(0)] = (d_2763_v85_) + (d_2763_v85_)
                nw429_[int(1)] = _dafny.SeqWithoutIsStrInference([d_2761_v83_, p0, p0])
                nw429_[int(2)] = (d_2763_v85_) + (d_2763_v85_)
                nw429_[int(3)] = (d_2763_v85_) + (_dafny.SeqWithoutIsStrInference([False]))
                nw429_[int(4)] = ((d_2763_v85_).set(default__.safeIndex(-117, len(d_2763_v85_)), p0)) + (_dafny.SeqWithoutIsStrInference([d_2761_v83_, p0]))
                nw429_[int(5)] = d_2763_v85_
                nw429_[int(6)] = (d_2763_v85_) + ((d_2763_v85_).set(default__.safeIndex((self).f10, len(d_2763_v85_)), d_2761_v83_))
                d_2807_v114_ = nw429_
                d_2808_v115_: _dafny.Seq
                d_2808_v115_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vmdclrqoy"))
                d_2809_v116_: _dafny.Array
                nw430_ = _dafny.Array(None, 19)
                nw430_[int(0)] = d_2808_v115_
                nw430_[int(1)] = d_2808_v115_
                nw430_[int(2)] = d_2808_v115_
                nw430_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmnbc"))
                nw430_[int(4)] = d_2808_v115_
                nw430_[int(5)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkynr"))) + (d_2808_v115_)
                nw430_[int(6)] = d_2808_v115_
                nw430_[int(7)] = (d_2808_v115_) + (d_2808_v115_)
                nw430_[int(8)] = (d_2808_v115_) + (d_2808_v115_)
                nw430_[int(9)] = d_2808_v115_
                nw430_[int(10)] = d_2808_v115_
                nw430_[int(11)] = d_2808_v115_
                nw430_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hit"))
                nw430_[int(13)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2810_i6_ in range(default__.abs(-221))])
                nw430_[int(14)] = d_2808_v115_
                nw430_[int(15)] = d_2808_v115_
                nw430_[int(16)] = (d_2808_v115_) + (d_2808_v115_)
                nw430_[int(17)] = d_2808_v115_
                nw430_[int(18)] = default__.fm37(globalState)
                d_2809_v116_ = nw430_
                index426_ = default__.safeIndex(432, (d_2809_v116_).length(0))
                (d_2809_v116_)[index426_] = d_2808_v115_
                d_2811_v117_: _dafny.MultiSet
                d_2811_v117_ = _dafny.MultiSet([d_2808_v115_, d_2808_v115_])
                d_2812_v118_: _dafny.Seq
                d_2812_v118_ = _dafny.SeqWithoutIsStrInference([(d_2811_v117_).set(d_2808_v115_, default__.abs((self).f10)), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2808_v115_ for d_2813_i7_ in range(default__.abs(213))]))])
                d_2814_v119_: C4
                nw431_ = C4()
                nw431_.ctor__((self).f10, d_2761_v83_, d_2808_v115_, (d_2811_v117_).ispropersubset((d_2812_v118_)[default__.safeIndex((self).f10, len(d_2812_v118_))]), True, (self).f10, (self).f10, True)
                d_2814_v119_ = nw431_
                d_2815_v120_: _dafny.Set
                d_2815_v120_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulavauekr"))})
                index427_ = default__.safeIndex(432, (d_2809_v116_).length(0))
                rhs399_ = (d_2814_v119_).f24
                rhs400_ = d_2807_v114_
                rhs401_ = (d_2815_v120_) == (d_2815_v120_)
                rhs402_ = d_2808_v115_
                rhs403_ = d_2814_v119_
                lhs221_ = d_2809_v116_
                lhs222_ = default__.safeIndex(432, (d_2809_v116_).length(0))
                d_2761_v83_ = rhs399_
                d_2807_v114_ = rhs400_
                d_2761_v83_ = rhs401_
                lhs221_[lhs222_] = rhs402_
                d_2814_v119_ = rhs403_
                d_2816_v121_: _dafny.Set
                d_2816_v121_ = _dafny.Set({not((d_2814_v119_).f24)})
                (globalState).f0 = d_2816_v121_
                d_2817_v122_: _dafny.Seq
                d_2817_v122_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f10)])
                d_2818_v123_: _dafny.Map
                d_2818_v123_ = _dafny.Map({(self).f10: len(d_2817_v122_)})
                d_2819_v124_: _dafny.Map
                d_2819_v124_ = (d_2818_v123_) | (d_2818_v123_)
                d_2820_v125_: int
                d_2820_v125_ = 835
                d_2821_v126_: _dafny.Map
                d_2821_v126_ = _dafny.Map({True: default__.fm25((0) - (len(d_2817_v122_)), (self).f10, not(d_2761_v83_), globalState)})
                rhs404_ = d_2819_v124_
                rhs405_ = ((d_2821_v126_)[(d_2814_v119_).f24] if ((d_2814_v119_).f24) in (d_2821_v126_) else (d_2820_v125_) - (len(d_2808_v115_)))
                rhs406_ = not(d_2761_v83_)
                d_2819_v124_ = rhs404_
                d_2820_v125_ = rhs405_
                d_2761_v83_ = rhs406_
                index428_ = default__.safeIndex(800, ((self).f11).length(0))
                ((self).f11)[index428_] = (d_2763_v85_)[default__.safeIndex(d_2820_v125_, len(d_2763_v85_))]
                index429_ = default__.safeIndex(800, ((self).f11).length(0))
                ((self).f11)[index429_] = (d_2814_v119_).f24
                index430_ = default__.safeIndex(822, (d_2629_v0_).length(0))
                (d_2629_v0_)[index430_] = len(d_2817_v122_)
                index431_ = default__.safeIndex(822, (d_2629_v0_).length(0))
                (d_2629_v0_)[index431_] = 541
            d_2822_v127_: _dafny.Seq
            d_2822_v127_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
            d_2823_v128_: C6
            nw432_ = C6()
            nw432_.ctor__(((self).f10) * ((self).f10), d_2761_v83_, p0, len(d_2822_v127_), (self).f10, d_2761_v83_)
            d_2823_v128_ = nw432_
            d_2824_v129_: _dafny.Array
            nw433_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_2824_v129_ = nw433_
            index432_ = default__.safeIndex(578, (d_2824_v129_).length(0))
            (d_2824_v129_)[index432_] = d_2629_v0_
            index433_ = default__.safeIndex(578, (d_2824_v129_).length(0))
            (d_2824_v129_)[index433_] = d_2629_v0_
            d_2825_v130_: str
            d_2825_v130_ = _dafny.CodePoint('c')
            d_2825_v130_ = d_2825_v130_

    @property
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11

class C11(T0):
    def  __init__(self):
        self._f4: bool = False
        self._f5: int = int(0)
        self._f9: D1 = D1.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C11"
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f9, f4, f5):
        (self)._f9 = f9
        (self)._f4 = f4
        (self)._f5 = f5

    def fm0(self, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewokgucoo"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))

    def fm12(self, p0, p1, p2, p3, globalState):
        if True:
            return (_dafny.MultiSet([(self).f4])).intersection((_dafny.MultiSet([(self).f4])))
        elif (self).f4:
            return _dafny.MultiSet([(self).f4])
        elif True:
            return _dafny.MultiSet([(self).f4])

    def fm13(self, globalState):
        return ((D1_DC4((self).f4, (self).f5, (self).f5)).cf14) - ((self).f5)

    def m3(self, p0, globalState):
        d_2826_i0_: int
        d_2826_i0_ = 0
        with _dafny.label("8"):
            while (self).f4:
                with _dafny.c_label("8"):
                    if (d_2826_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_2826_i0_ = (d_2826_i0_) + (1)
                    d_2827_v0_: _dafny.Array
                    nw434_ = _dafny.Array(None, 5)
                    nw434_[int(0)] = (self).f4
                    nw434_[int(1)] = (self).f4
                    nw434_[int(2)] = (self).f4
                    nw434_[int(3)] = (self).f4
                    nw434_[int(4)] = (self).f4
                    d_2827_v0_ = nw434_
                    d_2828_v1_: _dafny.Array
                    nw435_ = _dafny.Array(None, 20)
                    nw435_[int(0)] = (d_2827_v0_ if (self).f4 else d_2827_v0_)
                    nw435_[int(1)] = d_2827_v0_
                    nw435_[int(2)] = d_2827_v0_
                    nw435_[int(3)] = (d_2827_v0_ if (self).f4 else d_2827_v0_)
                    nw435_[int(4)] = (d_2827_v0_ if not((self).f4) else d_2827_v0_)
                    nw435_[int(5)] = d_2827_v0_
                    nw435_[int(6)] = d_2827_v0_
                    nw435_[int(7)] = d_2827_v0_
                    nw435_[int(8)] = (d_2827_v0_ if (self).f4 else d_2827_v0_)
                    nw435_[int(9)] = d_2827_v0_
                    nw435_[int(10)] = d_2827_v0_
                    nw435_[int(11)] = d_2827_v0_
                    nw435_[int(12)] = d_2827_v0_
                    nw435_[int(13)] = d_2827_v0_
                    nw435_[int(14)] = d_2827_v0_
                    nw435_[int(15)] = d_2827_v0_
                    nw435_[int(16)] = d_2827_v0_
                    nw435_[int(17)] = d_2827_v0_
                    nw435_[int(18)] = d_2827_v0_
                    nw435_[int(19)] = d_2827_v0_
                    d_2828_v1_ = nw435_
                    index434_ = default__.safeIndex(478, (d_2828_v1_).length(0))
                    (d_2828_v1_)[index434_] = d_2827_v0_
                    index435_ = default__.safeIndex(478, (d_2828_v1_).length(0))
                    (d_2828_v1_)[index435_] = d_2827_v0_
                    d_2829_v2_: int
                    d_2829_v2_ = 394
                    d_2830_v3_: _dafny.Seq
                    d_2830_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlvbshut"))
                    d_2831_v4_: _dafny.Array
                    nw436_ = _dafny.Array(int(0), 29)
                    d_2831_v4_ = nw436_
                    d_2832_v5_: str
                    d_2832_v5_ = _dafny.CodePoint('y')
                    d_2833_v6_: _dafny.Map
                    d_2833_v6_ = _dafny.Map({d_2831_v4_: d_2832_v5_})
                    d_2834_v7_: D1
                    d_2834_v7_ = D1_DC4((self).f4, len(_dafny.SeqWithoutIsStrInference([746 for d_2835_i1_ in range(default__.abs(-164))])), (self).f5)
                    d_2836_v8_: _dafny.Seq
                    d_2836_v8_ = _dafny.SeqWithoutIsStrInference([d_2834_v7_, d_2834_v7_])
                    d_2837_v9_: _dafny.MultiSet
                    d_2837_v9_ = _dafny.MultiSet([False])
                    d_2838_v10_: D0
                    d_2838_v10_ = D0_DC1(-293, (self).f4, 514, p0)
                    d_2839_v11_: _dafny.MultiSet
                    d_2839_v11_ = _dafny.MultiSet([p0])
                    d_2840_v12_: _dafny.Seq
                    d_2840_v12_ = _dafny.SeqWithoutIsStrInference([len(d_2830_v3_)])
                    d_2841_v13_: _dafny.Map
                    d_2841_v13_ = _dafny.Map({len(d_2840_v12_): p0})
                    d_2842_v14_: _dafny.Map
                    d_2842_v14_ = _dafny.Map({d_2841_v13_: d_2832_v5_})
                    d_2843_v15_: _dafny.Map
                    d_2843_v15_ = _dafny.Map({p0: (self).f4})
                    d_2844_v16_: _dafny.Array
                    nw437_ = _dafny.Array(None, 28)
                    nw437_[int(0)] = d_2829_v2_
                    nw437_[int(1)] = len(_dafny.SeqWithoutIsStrInference([(self).f4]))
                    nw437_[int(2)] = -810
                    nw437_[int(3)] = p0
                    nw437_[int(4)] = p0
                    nw437_[int(5)] = p0
                    nw437_[int(6)] = default__.safeModuloInt(d_2829_v2_, len(d_2830_v3_))
                    nw437_[int(7)] = d_2829_v2_
                    nw437_[int(8)] = len((d_2833_v6_) | ((d_2833_v6_).set(d_2831_v4_, d_2832_v5_)))
                    nw437_[int(9)] = len(d_2830_v3_)
                    nw437_[int(10)] = (0) - (len(d_2836_v8_))
                    nw437_[int(11)] = default__.safeModuloInt(p0, (self).f5)
                    nw437_[int(12)] = (0) - ((0) - (p0))
                    nw437_[int(13)] = (p0) * ((d_2837_v9_).cardinality)
                    nw437_[int(14)] = p0
                    nw437_[int(15)] = (0) - (p0)
                    nw437_[int(16)] = (d_2838_v10_).cf8
                    nw437_[int(17)] = d_2829_v2_
                    nw437_[int(18)] = (((d_2839_v11_)[(self).f5] if ((self).f5) in (d_2839_v11_) else p0) if (self).f4 else 708)
                    nw437_[int(19)] = 77
                    nw437_[int(20)] = d_2829_v2_
                    nw437_[int(21)] = default__.safeModuloInt(d_2829_v2_, len(d_2842_v14_))
                    nw437_[int(22)] = p0
                    nw437_[int(23)] = (d_2829_v2_ if (self).fm0(globalState) else 317)
                    nw437_[int(24)] = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_2832_v5_ for d_2845_i2_ in range(default__.abs(924))])), p0)
                    nw437_[int(25)] = len(d_2843_v15_)
                    nw437_[int(26)] = len(d_2830_v3_)
                    nw437_[int(27)] = d_2829_v2_
                    d_2844_v16_ = nw437_
                    index436_ = default__.safeIndex(182, (d_2831_v4_).length(0))
                    (d_2831_v4_)[index436_] = (0) - (len(d_2841_v13_))
                    d_2846_v17_: _dafny.MultiSet
                    d_2846_v17_ = _dafny.MultiSet([d_2840_v12_])
                    index437_ = default__.safeIndex(182, (d_2831_v4_).length(0))
                    rhs407_ = ((d_2839_v11_)[((d_2846_v17_).cardinality) * ((0) - ((self).f5))] if (((d_2846_v17_).cardinality) * ((0) - ((self).f5))) in (d_2839_v11_) else (0) - ((self).f5))
                    rhs408_ = (self).fm13(globalState)
                    lhs223_ = d_2831_v4_
                    lhs224_ = default__.safeIndex(182, (d_2831_v4_).length(0))
                    d_2829_v2_ = rhs407_
                    lhs223_[lhs224_] = rhs408_
                    d_2847_v18_: bool
                    d_2847_v18_ = True
                    d_2847_v18_ = d_2847_v18_
                    d_2829_v2_ = (d_2831_v4_)[default__.safeIndex(182, (d_2831_v4_).length(0))]
                    pass
            pass
        d_2848_v19_: bool
        d_2848_v19_ = True
        d_2848_v19_ = (self).f4
        d_2849_i3_: int
        d_2849_i3_ = 0
        with _dafny.label("9"):
            while d_2848_v19_:
                with _dafny.c_label("9"):
                    if (d_2849_i3_) >= (100):
                        raise _dafny.Break("9")
                    d_2849_i3_ = (d_2849_i3_) + (1)
                    d_2850_v20_: _dafny.Seq
                    d_2850_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbpmaa"))
                    d_2851_v21_: C8
                    nw438_ = C8()
                    nw438_.ctor__(d_2850_v20_, (self).f5)
                    d_2851_v21_ = nw438_
                    d_2852_v22_: D18
                    d_2852_v22_ = D18_DC50(True, (self).f4)
                    d_2853_v23_: D19
                    d_2853_v23_ = D19_DC53((d_2851_v21_).f14, (d_2851_v21_).f14, d_2848_v19_, (self).f4)
                    d_2854_v24_: _dafny.Map
                    d_2854_v24_ = _dafny.Map({p0: 178})
                    d_2855_v25_: _dafny.Set
                    d_2855_v25_ = _dafny.Set({(self).f4})
                    d_2856_v26_: D0
                    d_2856_v26_ = D0_DC0(d_2848_v19_, default__.fm25((d_2851_v21_).f14, p0, d_2848_v19_, globalState), (self).f4, len(d_2855_v25_), (self).f4)
                    pat_let_tv43_ = d_2848_v19_
                    d_2857_v27_: _dafny.Array
                    nw439_ = _dafny.Array(None, 23)
                    nw439_[int(0)] = D18_DC50(d_2848_v19_, True)
                    nw439_[int(1)] = d_2852_v22_
                    nw439_[int(2)] = d_2852_v22_
                    nw439_[int(3)] = d_2852_v22_
                    nw439_[int(4)] = d_2852_v22_
                    nw439_[int(5)] = d_2852_v22_
                    nw439_[int(6)] = d_2852_v22_
                    nw439_[int(7)] = d_2852_v22_
                    nw439_[int(8)] = d_2852_v22_
                    nw439_[int(9)] = d_2852_v22_
                    nw439_[int(10)] = (D18_DC50(d_2848_v19_, d_2848_v19_) if (d_2853_v23_).cf96 else d_2852_v22_)
                    nw439_[int(11)] = d_2852_v22_
                    nw439_[int(12)] = d_2852_v22_
                    nw439_[int(13)] = d_2852_v22_
                    nw439_[int(14)] = default__.fm79(d_2854_v24_, (d_2856_v26_).cf1, globalState)
                    nw439_[int(15)] = d_2852_v22_
                    def iife228_(_pat_let63_0):
                        def iife229_(d_2858_dt__update__tmp_h0_):
                            def iife230_(_pat_let64_0):
                                def iife231_(d_2859_dt__update_hcf90_h0_):
                                    return D18_DC50(d_2859_dt__update_hcf90_h0_, (d_2858_dt__update__tmp_h0_).cf91)
                                return iife231_(_pat_let64_0)
                            return iife230_(pat_let_tv43_)
                        return iife229_(_pat_let63_0)
                    nw439_[int(16)] = iife228_(d_2852_v22_)
                    nw439_[int(17)] = d_2852_v22_
                    nw439_[int(18)] = d_2852_v22_
                    nw439_[int(19)] = d_2852_v22_
                    nw439_[int(20)] = d_2852_v22_
                    nw439_[int(21)] = d_2852_v22_
                    nw439_[int(22)] = d_2852_v22_
                    d_2857_v27_ = nw439_
                    index438_ = default__.safeIndex(848, (d_2857_v27_).length(0))
                    (d_2857_v27_)[index438_] = default__.fm79(d_2854_v24_, (d_2851_v21_).f14, globalState)
                    index439_ = default__.safeIndex(848, (d_2857_v27_).length(0))
                    rhs409_ = default__.fm27((D16_DC41(default__.fm50(globalState), d_2848_v19_)).cf77, globalState)
                    rhs410_ = d_2852_v22_
                    lhs225_ = d_2857_v27_
                    lhs226_ = default__.safeIndex(848, (d_2857_v27_).length(0))
                    d_2848_v19_ = rhs409_
                    lhs225_[lhs226_] = rhs410_
                    d_2860_v28_: C5
                    nw440_ = C5()
                    nw440_.ctor__()
                    d_2860_v28_ = nw440_
                    d_2861_v29_: _dafny.Array
                    nw441_ = _dafny.Array(None, 5)
                    nw441_[int(0)] = (self).f4
                    nw441_[int(1)] = (self).f4
                    nw441_[int(2)] = (self).f4
                    nw441_[int(3)] = True
                    nw441_[int(4)] = d_2848_v19_
                    d_2861_v29_ = nw441_
                    d_2862_v30_: _dafny.Set
                    d_2862_v30_ = _dafny.Set({d_2861_v29_, d_2861_v29_, d_2861_v29_})
                    d_2863_v31_: _dafny.Seq
                    d_2863_v31_ = _dafny.SeqWithoutIsStrInference([d_2861_v29_])
                    d_2864_v32_: _dafny.Seq
                    d_2864_v32_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_2861_v29_, d_2861_v29_}), d_2862_v30_, _dafny.Set({(d_2863_v31_)[default__.safeIndex(p0, len(d_2863_v31_))]})])
                    d_2865_v33_: _dafny.Array
                    nw442_ = _dafny.Array(None, 11)
                    nw442_[int(0)] = d_2862_v30_
                    nw442_[int(1)] = _dafny.Set({d_2861_v29_, d_2861_v29_, d_2861_v29_, d_2861_v29_, d_2861_v29_})
                    nw442_[int(2)] = (d_2862_v30_).intersection(d_2862_v30_)
                    nw442_[int(3)] = d_2862_v30_
                    nw442_[int(4)] = d_2862_v30_
                    nw442_[int(5)] = (d_2864_v32_)[default__.safeIndex((d_2851_v21_).f14, len(d_2864_v32_))]
                    nw442_[int(6)] = d_2862_v30_
                    nw442_[int(7)] = d_2862_v30_
                    nw442_[int(8)] = d_2862_v30_
                    nw442_[int(9)] = (_dafny.Set({d_2861_v29_})).intersection(d_2862_v30_)
                    nw442_[int(10)] = d_2862_v30_
                    d_2865_v33_ = nw442_
                    index440_ = default__.safeIndex(719, (d_2865_v33_).length(0))
                    (d_2865_v33_)[index440_] = d_2862_v30_
                    index441_ = default__.safeIndex(719, (d_2865_v33_).length(0))
                    (d_2865_v33_)[index441_] = (d_2862_v30_ if (self).f4 else (d_2862_v30_).intersection(d_2862_v30_))
                    pass
            pass
        source36_ = default__.fm80(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_2866_i4_ in range(default__.abs(400))]), globalState)
        if source36_.is_DC45:
            d_2867_v34_: _dafny.MultiSet
            d_2867_v34_ = _dafny.MultiSet([591])
            d_2868_v35_: _dafny.Map
            d_2868_v35_ = _dafny.Map({(d_2867_v34_) | (d_2867_v34_): (self).f4})
            d_2869_v36_: str
            d_2869_v36_ = _dafny.CodePoint('w')
            d_2870_v37_: _dafny.Seq
            d_2870_v37_ = _dafny.SeqWithoutIsStrInference([d_2869_v36_, d_2869_v36_])
            d_2871_v38_: _dafny.MultiSet
            d_2871_v38_ = _dafny.MultiSet([d_2870_v37_])
            d_2848_v19_ = ((d_2868_v35_)[(d_2867_v34_).intersection(d_2867_v34_)] if ((d_2867_v34_).intersection(d_2867_v34_)) in (d_2868_v35_) else (d_2871_v38_).ispropersubset(d_2871_v38_))
            d_2872_v39_: _dafny.Array
            def lambda113_(d_2873_i5_):
                return (d_2873_i5_) - ((self).f5)

            init65_ = lambda113_
            nw443_ = _dafny.Array(None, 6)
            for i0_65_ in range(nw443_.length(0)):
                nw443_[i0_65_] = init65_(i0_65_)
            d_2872_v39_ = nw443_
            index442_ = default__.safeIndex(672, (d_2872_v39_).length(0))
            (d_2872_v39_)[index442_] = p0
            index443_ = default__.safeIndex(672, (d_2872_v39_).length(0))
            (d_2872_v39_)[index443_] = (((self).f5 if d_2848_v19_ else len(_dafny.Map({d_2848_v19_: default__.fm27(d_2869_v36_, globalState)})))) - (len((_dafny.SeqWithoutIsStrInference([d_2869_v36_])).set(default__.safeIndex(len(default__.fm48((self).f5, globalState)), len(_dafny.SeqWithoutIsStrInference([d_2869_v36_]))), d_2869_v36_)))
            if d_2848_v19_:
                d_2874_v40_: _dafny.Array
                nw444_ = _dafny.Array(_dafny.Map({}), 29)
                d_2874_v40_ = nw444_
                d_2875_v41_: _dafny.Map
                d_2875_v41_ = _dafny.Map({d_2874_v40_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skprdcjb"))})
                d_2875_v41_ = (d_2875_v41_).set(d_2874_v40_, _dafny.SeqWithoutIsStrInference([d_2869_v36_ for d_2876_i6_ in range(default__.abs(-769))]))
                index444_ = default__.safeIndex(672, (d_2872_v39_).length(0))
                (d_2872_v39_)[index444_] = (d_2872_v39_)[default__.safeIndex(672, (d_2872_v39_).length(0))]
                d_2877_v42_: _dafny.Map
                d_2877_v42_ = _dafny.Map({len(_dafny.Map({p0: (self).f5})): (self).f4})
                d_2878_v43_: C6
                nw445_ = C6()
                nw445_.ctor__((self).f5, (self).f4, (p0) != (p0), (self).f5, (0) - (default__.safeModuloInt(len(d_2877_v42_), (self).f5)), d_2848_v19_)
                d_2878_v43_ = nw445_
                (d_2878_v43_).f21 = (0) - (len((d_2870_v37_) + (d_2870_v37_)))
                d_2870_v37_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tasjbdod"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fredni")))
            elif True:
                d_2879_v44_: _dafny.Seq
                d_2879_v44_ = _dafny.SeqWithoutIsStrInference([d_2848_v19_])
                d_2880_v45_: _dafny.Map
                d_2880_v45_ = _dafny.Map({(d_2872_v39_)[default__.safeIndex(672, (d_2872_v39_).length(0))]: d_2879_v44_})
                d_2881_v46_: _dafny.Map
                d_2881_v46_ = _dafny.Map({len(d_2880_v45_): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdlq"))})
                index445_ = default__.safeIndex(672, (d_2872_v39_).length(0))
                rhs411_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmnhqaci"))) + (((d_2881_v46_)[(self).f5] if ((self).f5) in (d_2881_v46_) else d_2870_v37_))
                rhs412_ = (d_2872_v39_)[default__.safeIndex(672, (d_2872_v39_).length(0))]
                lhs227_ = d_2872_v39_
                lhs228_ = default__.safeIndex(672, (d_2872_v39_).length(0))
                d_2870_v37_ = rhs411_
                lhs227_[lhs228_] = rhs412_
                d_2882_v47_: _dafny.Array
                nw446_ = _dafny.Array(_dafny.Map({}), 20)
                d_2882_v47_ = nw446_
                d_2882_v47_ = d_2882_v47_
                d_2883_v48_: D9
                d_2883_v48_ = D9_DC22(default__.fm27(d_2869_v36_, globalState))
                index446_ = default__.safeIndex(672, (d_2872_v39_).length(0))
                rhs413_ = d_2883_v48_
                rhs414_ = p0
                lhs229_ = d_2872_v39_
                lhs230_ = default__.safeIndex(672, (d_2872_v39_).length(0))
                d_2883_v48_ = rhs413_
                lhs229_[lhs230_] = rhs414_
                d_2884_v49_: _dafny.Map
                d_2884_v49_ = _dafny.Map({d_2869_v36_: d_2879_v44_})
                d_2885_v50_: _dafny.Seq
                d_2885_v50_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2886_v51_: _dafny.MultiSet
                d_2886_v51_ = _dafny.MultiSet([((d_2884_v49_)[d_2869_v36_] if (d_2869_v36_) in (d_2884_v49_) else d_2879_v44_), d_2879_v44_, (d_2879_v44_).set(default__.safeIndex(len(d_2885_v50_), len(d_2879_v44_)), (self).f4), d_2879_v44_])
                d_2887_v52_: bool
                out26_: bool
                out26_ = default__.m0((self).f4, ((d_2886_v51_) - (_dafny.MultiSet([default__.fm45((d_2872_v39_)[default__.safeIndex(672, (d_2872_v39_).length(0))], globalState)]))).cardinality, globalState)
                d_2887_v52_ = out26_
                index447_ = default__.safeIndex(672, (d_2872_v39_).length(0))
                (d_2872_v39_)[index447_] = (d_2872_v39_)[default__.safeIndex(672, (d_2872_v39_).length(0))]
            d_2888_v53_: _dafny.Map
            d_2888_v53_ = _dafny.Map({(self).f5: (self).f5})
            index448_ = default__.safeIndex(672, (d_2872_v39_).length(0))
            (d_2872_v39_)[index448_] = ((d_2888_v53_)[(d_2872_v39_)[default__.safeIndex(672, (d_2872_v39_).length(0))]] if ((d_2872_v39_)[default__.safeIndex(672, (d_2872_v39_).length(0))]) in (d_2888_v53_) else 481)
        elif source36_.is_DC46:
            d_2889_v54_: _dafny.Map
            d_2889_v54_ = _dafny.Map({(self).f4: d_2848_v19_})
            d_2889_v54_ = (d_2889_v54_).set(True, not((self).f4))
            d_2890_v55_: int
            d_2890_v55_ = 764
            d_2890_v55_ = 30
            d_2890_v55_ = d_2890_v55_
            d_2891_v56_: _dafny.Array
            nw447_ = _dafny.Array(_dafny.Map({}), 29)
            d_2891_v56_ = nw447_
            d_2892_v57_: _dafny.Set
            d_2892_v57_ = _dafny.Set({d_2848_v19_})
            d_2893_v58_: _dafny.Array
            nw448_ = _dafny.Array(int(0), 22)
            d_2893_v58_ = nw448_
            d_2894_v59_: _dafny.Map
            d_2894_v59_ = _dafny.Map({d_2892_v57_: d_2893_v58_})
            index449_ = default__.safeIndex(449, (d_2891_v56_).length(0))
            (d_2891_v56_)[index449_] = d_2894_v59_
            index450_ = default__.safeIndex(449, (d_2891_v56_).length(0))
            (d_2891_v56_)[index450_] = d_2894_v59_
        elif source36_.is_DC47:
            d_2895___mcc_h0_ = source36_.cf83
            d_2896___mcc_h1_ = source36_.cf84
            d_2897___mcc_h2_ = source36_.cf85
            d_2898___mcc_h3_ = source36_.cf86
            d_2899___mcc_h4_ = source36_.cf87
            d_2900_cf87_ = d_2899___mcc_h4_
            d_2901_cf86_ = d_2898___mcc_h3_
            d_2902_cf85_ = d_2897___mcc_h2_
            d_2903_cf84_ = d_2896___mcc_h1_
            d_2904_cf83_ = d_2895___mcc_h0_
            d_2905_v60_: str
            d_2905_v60_ = _dafny.CodePoint('l')
            d_2848_v19_ = (not((self).f4)) and ((d_2905_v60_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adi"))))
            d_2906_v61_: _dafny.Seq
            d_2906_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ciktkcqs"))
            d_2906_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgh"))
            d_2900_cf87_ = d_2902_cf85_
            d_2907_v62_: _dafny.Seq
            d_2907_v62_ = _dafny.SeqWithoutIsStrInference([d_2900_cf87_, (self).f5, (0) - (len(d_2906_v61_))])
            d_2908_v63_: D6
            d_2908_v63_ = D6_DC15(d_2907_v62_)
            d_2908_v63_ = d_2908_v63_
        elif source36_.is_DC44:
            d_2909___mcc_h5_ = source36_.cf82
            d_2910_cf82_ = d_2909___mcc_h5_
            d_2911_v64_: str
            d_2911_v64_ = _dafny.CodePoint('x')
            d_2912_v65_: _dafny.Seq
            d_2912_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tdcnr"))
            d_2848_v19_ = (d_2911_v64_) in (d_2912_v65_)
            if (self).f4:
                d_2911_v64_ = (d_2912_v65_)[default__.safeIndex(default__.safeModuloInt((self).f5, p0), len(d_2912_v65_))]
                d_2913_v66_: _dafny.Set
                d_2913_v66_ = _dafny.Set({(d_2848_v19_) == (True)})
                (globalState).f0 = d_2913_v66_
                d_2912_v65_ = (d_2912_v65_) + (d_2912_v65_)
                d_2914_v67_: int
                d_2914_v67_ = 26
                d_2915_v68_: _dafny.MultiSet
                d_2915_v68_ = _dafny.MultiSet([d_2848_v19_])
                d_2916_v69_: _dafny.Map
                d_2916_v69_ = _dafny.Map({d_2848_v19_: (self).f4})
                d_2914_v67_ = (default__.safeDivisionInt((d_2915_v68_).cardinality, len(d_2916_v69_)) if (self).f4 else ((self).f5) + (len(d_2912_v65_)))
                d_2848_v19_ = d_2848_v19_
            elif True:
                d_2917_v70_: _dafny.Array
                def lambda114_(d_2918_i7_):
                    return (self).f4

                init66_ = lambda114_
                nw449_ = _dafny.Array(None, 19)
                for i0_66_ in range(nw449_.length(0)):
                    nw449_[i0_66_] = init66_(i0_66_)
                d_2917_v70_ = nw449_
                index451_ = default__.safeIndex(934, (d_2917_v70_).length(0))
                (d_2917_v70_)[index451_] = d_2848_v19_
                d_2919_v71_: _dafny.Map
                d_2919_v71_ = _dafny.Map({p0: default__.fm18(d_2848_v19_, d_2848_v19_, d_2848_v19_, globalState)})
                d_2920_v72_: _dafny.MultiSet
                d_2920_v72_ = _dafny.MultiSet([(self).f4, not(d_2848_v19_)])
                index452_ = default__.safeIndex(934, (d_2917_v70_).length(0))
                (d_2917_v70_)[index452_] = ((d_2920_v72_) - (d_2920_v72_)).issubset((self).fm12(d_2919_v71_, d_2848_v19_, d_2848_v19_, d_2912_v65_, globalState))
                d_2921_v73_: C5
                nw450_ = C5()
                nw450_.ctor__()
                d_2921_v73_ = nw450_
                d_2922_v74_: int
                d_2922_v74_ = -454
                d_2923_v75_: _dafny.Seq
                d_2923_v75_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f5, p0]))])
                d_2922_v74_ = default__.safeModuloInt((self).f5, (len(d_2912_v65_)) - ((0) - (len(d_2923_v75_))))
                d_2922_v74_ = (0) - (d_2922_v74_)
                d_2922_v74_ = 186
            d_2924_v76_: _dafny.Array
            def lambda115_(d_2925_i8_):
                return (self).f4

            init67_ = lambda115_
            nw451_ = _dafny.Array(None, 7)
            for i0_67_ in range(nw451_.length(0)):
                nw451_[i0_67_] = init67_(i0_67_)
            d_2924_v76_ = nw451_
            d_2926_v77_: D4
            d_2926_v77_ = D4_DC8(d_2924_v76_)
            d_2927_v78_: C10
            nw452_ = C10()
            nw452_.ctor__(p0, (d_2926_v77_).cf20)
            d_2927_v78_ = nw452_
            d_2848_v19_ = (self).f4
        elif True:
            d_2928___mcc_h6_ = source36_.cf88
            d_2929_cf88_ = d_2928___mcc_h6_
            d_2930_v79_: str
            d_2930_v79_ = _dafny.CodePoint('w')
            d_2931_v80_: _dafny.Map
            d_2931_v80_ = _dafny.Map({d_2930_v79_: (self).f5})
            d_2932_v82_: _dafny.Set
            d_2932_v82_ = _dafny.Set({d_2930_v79_})
            d_2933_v83_: _dafny.Seq
            d_2933_v83_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_2934_v84_: _dafny.Seq
            d_2934_v84_ = _dafny.SeqWithoutIsStrInference([d_2931_v80_, (d_2931_v80_).set(d_2930_v79_, (d_2933_v83_)[default__.safeIndex((self).f5, len(d_2933_v83_))])])
            d_2935_v86_: _dafny.MultiSet
            d_2935_v86_ = _dafny.MultiSet([_dafny.CodePoint('n')])
            d_2936_v87_: _dafny.Array
            nw453_ = _dafny.Array(None, 27)
            nw453_[int(0)] = d_2931_v80_
            nw453_[int(1)] = d_2931_v80_
            nw453_[int(2)] = d_2931_v80_
            nw453_[int(3)] = d_2931_v80_
            nw453_[int(4)] = _dafny.Map({d_2930_v79_: (self).f5})
            nw453_[int(5)] = _dafny.Map({d_2930_v79_: (self).f5})
            nw453_[int(6)] = d_2931_v80_
            nw453_[int(7)] = d_2931_v80_
            nw453_[int(8)] = (d_2931_v80_).set(d_2930_v79_, p0)
            nw453_[int(9)] = d_2931_v80_
            nw453_[int(10)] = d_2931_v80_
            nw453_[int(11)] = d_2931_v80_
            def iife232_():
                coll102_ = _dafny.Map()
                compr_102_: str
                for compr_102_ in (d_2932_v82_).Elements:
                    d_2937_v81_: str = compr_102_
                    if (d_2937_v81_) in (d_2932_v82_):
                        coll102_[d_2937_v81_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nn")))
                return _dafny.Map(coll102_)
            nw453_[int(12)] = iife232_()
            
            nw453_[int(13)] = d_2931_v80_
            nw453_[int(14)] = d_2931_v80_
            nw453_[int(15)] = _dafny.Map({d_2930_v79_: (self).f5})
            nw453_[int(16)] = _dafny.Map({d_2930_v79_: p0})
            nw453_[int(17)] = (d_2934_v84_)[default__.safeIndex((self).f5, len(d_2934_v84_))]
            nw453_[int(18)] = d_2931_v80_
            def iife233_():
                coll103_ = _dafny.Map()
                compr_103_: str
                for compr_103_ in (d_2935_v86_).Elements:
                    d_2938_v85_: str = compr_103_
                    if (d_2938_v85_) in (d_2935_v86_):
                        coll103_[d_2938_v85_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uohwx")))
                return _dafny.Map(coll103_)
            nw453_[int(19)] = iife233_()
            
            nw453_[int(20)] = d_2931_v80_
            nw453_[int(21)] = d_2931_v80_
            nw453_[int(22)] = d_2931_v80_
            nw453_[int(23)] = d_2931_v80_
            nw453_[int(24)] = d_2931_v80_
            nw453_[int(25)] = d_2931_v80_
            nw453_[int(26)] = d_2931_v80_
            d_2936_v87_ = nw453_
            d_2939_v88_: D14
            d_2939_v88_ = D14_DC36(d_2936_v87_)
            d_2939_v88_ = d_2939_v88_
            d_2940_v89_: _dafny.Seq
            d_2940_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onc"))
            d_2940_v89_ = _dafny.SeqWithoutIsStrInference([d_2930_v79_ for d_2941_i9_ in range(default__.abs(143))])
            d_2942_v90_: _dafny.Seq
            d_2942_v90_ = _dafny.SeqWithoutIsStrInference([d_2848_v19_])
            d_2943_v91_: _dafny.Seq
            d_2943_v91_ = _dafny.SeqWithoutIsStrInference([d_2942_v90_, _dafny.SeqWithoutIsStrInference([(self).fm0(globalState)])])
            d_2944_v92_: _dafny.Map
            d_2944_v92_ = _dafny.Map({len(d_2943_v91_): d_2848_v19_})
            d_2945_v93_: _dafny.Set
            d_2945_v93_ = _dafny.Set({d_2848_v19_, (self).f4})
            d_2946_v94_: D0
            d_2946_v94_ = D0_DC0(d_2848_v19_, (self).f5, (self).f4, len(d_2945_v93_), (self).f4)
            d_2947_v95_: _dafny.MultiSet
            d_2947_v95_ = _dafny.MultiSet([(self).f5])
            d_2948_v96_: _dafny.Array
            nw454_ = _dafny.Array(None, 12)
            nw454_[int(0)] = (self).f4
            nw454_[int(1)] = d_2848_v19_
            nw454_[int(2)] = (default__.fm50(globalState)) in (d_2940_v89_)
            nw454_[int(3)] = d_2848_v19_
            nw454_[int(4)] = ((d_2944_v92_)[(d_2946_v94_).cf3] if ((d_2946_v94_).cf3) in (d_2944_v92_) else d_2848_v19_)
            nw454_[int(5)] = not (d_2848_v19_) or (d_2848_v19_)
            nw454_[int(6)] = (self).f4
            nw454_[int(7)] = (d_2848_v19_ if True else (self).f4)
            nw454_[int(8)] = (_dafny.MultiSet([len(d_2940_v89_)])).ispropersubset(d_2947_v95_)
            nw454_[int(9)] = d_2848_v19_
            nw454_[int(10)] = d_2848_v19_
            nw454_[int(11)] = not((self).f4)
            d_2948_v96_ = nw454_
            index453_ = default__.safeIndex(14, (d_2948_v96_).length(0))
            (d_2948_v96_)[index453_] = (self).f4
            d_2949_v97_: _dafny.Set
            d_2949_v97_ = _dafny.Set({d_2948_v96_})
            index454_ = default__.safeIndex(14, (d_2948_v96_).length(0))
            (d_2948_v96_)[index454_] = (d_2949_v97_).isdisjoint(d_2949_v97_)
            index455_ = default__.safeIndex(14, (d_2948_v96_).length(0))
            (d_2948_v96_)[index455_] = (len(_dafny.Set({d_2848_v19_}))) > ((0) - (-5))
        source37_ = default__.fm41(_dafny.Set({(self).f5}), globalState)
        d_2950___mcc_h7_ = source37_
        d_2951_cf42_ = d_2950___mcc_h7_
        d_2952_v98_: _dafny.Seq
        d_2952_v98_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dutnuuy"))
        d_2953_v99_: _dafny.Map
        d_2953_v99_ = _dafny.Map({(self).f4: (self).f5})
        d_2954_v100_: _dafny.Seq
        d_2954_v100_ = _dafny.SeqWithoutIsStrInference([True, (self).f4])
        d_2955_v101_: _dafny.Map
        d_2955_v101_ = _dafny.Map({p0: (self).f5})
        d_2956_v102_: str
        d_2956_v102_ = _dafny.CodePoint('y')
        d_2957_v103_: _dafny.Seq
        d_2957_v103_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_2958_v104_: _dafny.Map
        d_2958_v104_ = _dafny.Map({d_2956_v102_: len(d_2957_v103_)})
        d_2959_v105_: _dafny.MultiSet
        d_2959_v105_ = _dafny.MultiSet([len(d_2957_v103_), p0])
        d_2960_v106_: _dafny.Seq
        d_2960_v106_ = _dafny.SeqWithoutIsStrInference([d_2957_v103_, d_2957_v103_])
        d_2961_v107_: _dafny.Map
        d_2961_v107_ = _dafny.Map({d_2955_v101_: (0) - ((self).f5)})
        d_2962_v108_: _dafny.Array
        nw455_ = _dafny.Array(None, 17)
        nw455_[int(0)] = p0
        nw455_[int(1)] = len(d_2953_v99_)
        nw455_[int(2)] = len(d_2954_v100_)
        nw455_[int(3)] = p0
        nw455_[int(4)] = len(_dafny.Set({len(d_2955_v101_)}))
        nw455_[int(5)] = -603
        nw455_[int(6)] = len(d_2958_v104_)
        nw455_[int(7)] = -312
        nw455_[int(8)] = p0
        nw455_[int(9)] = p0
        nw455_[int(10)] = (d_2959_v105_).cardinality
        nw455_[int(11)] = len(d_2952_v98_)
        nw455_[int(12)] = len((d_2960_v106_)[default__.safeIndex(p0, len(d_2960_v106_))])
        nw455_[int(13)] = len(d_2961_v107_)
        nw455_[int(14)] = (self).f5
        nw455_[int(15)] = (self).f5
        nw455_[int(16)] = (self).f5
        d_2962_v108_ = nw455_
        d_2963_v109_: _dafny.Map
        d_2963_v109_ = _dafny.Map({len(d_2952_v98_): d_2962_v108_})
        d_2964_v110_: _dafny.MultiSet
        d_2964_v110_ = _dafny.MultiSet([d_2956_v102_])
        d_2965_v111_: _dafny.Array
        nw456_ = _dafny.Array(None, 18)
        nw456_[int(0)] = d_2848_v19_
        nw456_[int(1)] = d_2848_v19_
        nw456_[int(2)] = ((self).f4) == (False)
        nw456_[int(3)] = (False) == (True)
        nw456_[int(4)] = ((self).f4 if (self).f4 else (self).f4)
        nw456_[int(5)] = (D5_DC13(d_2848_v19_, d_2848_v19_)).cf30
        nw456_[int(6)] = (self).f4
        nw456_[int(7)] = not((self).f4)
        nw456_[int(8)] = d_2848_v19_
        nw456_[int(9)] = (d_2963_v109_) != (_dafny.Map({(self).f5: d_2962_v108_}))
        nw456_[int(10)] = (p0) <= (p0)
        nw456_[int(11)] = d_2848_v19_
        nw456_[int(12)] = ((self).f5) != (((d_2964_v110_)[d_2956_v102_] if (d_2956_v102_) in (d_2964_v110_) else p0))
        nw456_[int(13)] = d_2848_v19_
        nw456_[int(14)] = (self).f4
        nw456_[int(15)] = not((self).f4)
        nw456_[int(16)] = (self).f4
        nw456_[int(17)] = d_2848_v19_
        d_2965_v111_ = nw456_
        d_2966_v112_: _dafny.MultiSet
        d_2966_v112_ = _dafny.MultiSet([(self).f4])
        index456_ = default__.safeIndex(948, (d_2965_v111_).length(0))
        (d_2965_v111_)[index456_] = (d_2966_v112_) != (d_2966_v112_)
        d_2967_v113_: int
        d_2967_v113_ = 732
        index457_ = default__.safeIndex(468, (d_2965_v111_).length(0))
        (d_2965_v111_)[index457_] = (self).f4
        d_2968_v114_: _dafny.Array
        def lambda116_(d_2969_v19_):
            def lambda117_(d_2970_i10_):
                return _dafny.Map({(self).f4: (D9_DC22(d_2969_v19_)).cf44})

            return lambda117_

        init68_ = lambda116_(d_2848_v19_)
        nw457_ = _dafny.Array(None, 10)
        for i0_68_ in range(nw457_.length(0)):
            nw457_[i0_68_] = init68_(i0_68_)
        d_2968_v114_ = nw457_
        d_2971_v115_: _dafny.Map
        d_2971_v115_ = _dafny.Map({d_2968_v114_: (self).f5})
        index458_ = default__.safeIndex(948, (d_2965_v111_).length(0))
        index459_ = default__.safeIndex(468, (d_2965_v111_).length(0))
        rhs415_ = (d_2968_v114_) not in (d_2971_v115_)
        rhs416_ = (0) - (default__.safeModuloInt(default__.safeModuloInt(d_2967_v113_, (self).f5), (_dafny.MultiSet([(self).f4, (self).f4])).cardinality))
        rhs417_ = d_2848_v19_
        lhs231_ = d_2965_v111_
        lhs232_ = default__.safeIndex(948, (d_2965_v111_).length(0))
        lhs233_ = d_2965_v111_
        lhs234_ = default__.safeIndex(468, (d_2965_v111_).length(0))
        lhs231_[lhs232_] = rhs415_
        d_2967_v113_ = rhs416_
        lhs233_[lhs234_] = rhs417_
        d_2972_v116_: _dafny.Array
        nw458_ = _dafny.Array(_dafny.CodePoint('D'), 19)
        d_2972_v116_ = nw458_
        d_2972_v116_ = d_2972_v116_
        d_2967_v113_ = (0) - (d_2967_v113_)
        index460_ = default__.safeIndex(948, (d_2965_v111_).length(0))
        rhs418_ = ((d_2951_cf42_) | (d_2951_cf42_)).intersection(d_2951_cf42_)
        rhs419_ = not(d_2848_v19_)
        rhs420_ = ((d_2955_v101_)[d_2967_v113_] if (d_2967_v113_) in (d_2955_v101_) else -789)
        lhs235_ = d_2965_v111_
        lhs236_ = default__.safeIndex(948, (d_2965_v111_).length(0))
        d_2951_cf42_ = rhs418_
        lhs235_[lhs236_] = rhs419_
        d_2967_v113_ = rhs420_
        d_2848_v19_ = (default__.safeModuloInt((self).f5, p0)) >= (795)

    @property
    def f9(self):
        return self._f9

class C12:
    def  __init__(self):
        self._f8: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C12"
    def ctor__(self, f8):
        (self)._f8 = f8

    def fm5(self, p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('i')

    def fm6(self, p0, p1, p2, globalState):
        return not(not(((False) or (True)) == ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hethahs"))) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_2973_i0_ in range(default__.abs(135))])))))

    def m5(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        r2: bool = False
        r3: _dafny.Array = _dafny.Array(None, 0)
        r2 = not (p1) or (p1)
        d_2974_v0_: _dafny.Set
        d_2974_v0_ = _dafny.Set({(self).f8})
        d_2975_v1_: _dafny.Seq
        d_2975_v1_ = _dafny.SeqWithoutIsStrInference([d_2974_v0_, d_2974_v0_])
        d_2976_v2_: bool
        out27_: bool
        out27_ = default__.m0((d_2975_v1_) <= (d_2975_v1_), (0) - (((self).f8) + ((self).f8)), globalState)
        d_2976_v2_ = out27_
        if d_2976_v2_:
            d_2977_v3_: _dafny.Seq
            d_2977_v3_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f8), (self).f8, (self).f8, (self).f8])
            d_2977_v3_ = d_2977_v3_
            d_2978_v4_: int
            d_2978_v4_ = -624
            d_2978_v4_ = ((self).f8) + (default__.safeModuloInt((self).f8, (self).f8))
            d_2979_v5_: _dafny.Seq
            d_2979_v5_ = _dafny.SeqWithoutIsStrInference([d_2976_v2_])
            source38_ = D1_DC4((not(d_2976_v2_) if (d_2979_v5_)[default__.safeIndex((self).f8, len(d_2979_v5_))] else p1), (0) - (d_2978_v4_), (self).f8)
            if source38_.is_DC4:
                d_2980___mcc_h0_ = source38_.cf12
                d_2981___mcc_h1_ = source38_.cf13
                d_2982___mcc_h2_ = source38_.cf14
                d_2983_cf14_ = d_2982___mcc_h2_
                d_2984_cf13_ = d_2981___mcc_h1_
                d_2985_cf12_ = d_2980___mcc_h0_
                d_2986_v6_: _dafny.Map
                d_2986_v6_ = _dafny.Map({d_2983_cf14_: d_2978_v4_})
                d_2986_v6_ = (d_2986_v6_).set(468, d_2984_cf13_)
                d_2987_v7_: _dafny.Map
                d_2987_v7_ = _dafny.Map({d_2985_cf12_: (self).f8})
                d_2988_v8_: _dafny.Map
                d_2988_v8_ = d_2987_v7_
                d_2989_v9_: _dafny.Seq
                d_2989_v9_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([161])])
                d_2990_v10_: _dafny.Set
                d_2990_v10_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_2983_cf14_ for d_2991_i0_ in range(default__.abs(-940))])})
                d_2992_v11_: _dafny.MultiSet
                d_2992_v11_ = _dafny.MultiSet([(0) - (-797)])
                d_2993_v12_: _dafny.Array
                nw459_ = _dafny.Array(None, 15)
                nw459_[int(0)] = _dafny.MultiSet([len((d_2988_v8_))])
                nw459_[int(1)] = _dafny.MultiSet([(self).f8, d_2978_v4_])
                nw459_[int(2)] = (d_2989_v9_)[default__.safeIndex((d_2977_v3_)[default__.safeIndex((0) - (len(d_2974_v0_)), len(d_2977_v3_))], len(d_2989_v9_))]
                nw459_[int(3)] = _dafny.MultiSet([-309])
                nw459_[int(4)] = default__.fm7(d_2976_v2_, d_2990_v10_, (self).f8, (self).f8, globalState)
                nw459_[int(5)] = d_2992_v11_
                nw459_[int(6)] = default__.fm7(d_2985_cf12_, d_2990_v10_, 440, d_2984_cf13_, globalState)
                nw459_[int(7)] = _dafny.MultiSet(d_2977_v3_)
                nw459_[int(8)] = (d_2992_v11_) | (d_2992_v11_)
                nw459_[int(9)] = (d_2992_v11_ if p1 else d_2992_v11_)
                nw459_[int(10)] = d_2992_v11_
                nw459_[int(11)] = d_2992_v11_
                nw459_[int(12)] = d_2992_v11_
                nw459_[int(13)] = d_2992_v11_
                nw459_[int(14)] = d_2992_v11_
                d_2993_v12_ = nw459_
                index461_ = default__.safeIndex(444, (d_2993_v12_).length(0))
                (d_2993_v12_)[index461_] = (d_2992_v11_ if True else (d_2992_v11_).set((self).f8, default__.abs(d_2978_v4_)))
                d_2994_v13_: _dafny.Set
                d_2994_v13_ = _dafny.Set({p1, d_2976_v2_, d_2976_v2_, not(d_2976_v2_), True})
                d_2995_v14_: _dafny.Map
                d_2995_v14_ = _dafny.Map({(d_2994_v13_ if p1 else d_2994_v13_): d_2992_v11_})
                index462_ = default__.safeIndex(444, (d_2993_v12_).length(0))
                (d_2993_v12_)[index462_] = ((d_2995_v14_)[d_2994_v13_] if (d_2994_v13_) in (d_2995_v14_) else (d_2989_v9_)[default__.safeIndex((self).f8, len(d_2989_v9_))])
                d_2996_v15_: D1
                d_2996_v15_ = D1_DC4(False, len(default__.fm8(d_2976_v2_, d_2978_v4_, 904, globalState)), d_2978_v4_)
                d_2997_v16_: _dafny.Seq
                d_2997_v16_ = _dafny.SeqWithoutIsStrInference([d_2979_v5_])
                d_2998_v17_: str
                d_2998_v17_ = _dafny.CodePoint('b')
                pat_let_tv44_ = d_2978_v4_
                d_2999_v18_: _dafny.Array
                nw460_ = _dafny.Array(None, 12)
                nw460_[int(0)] = d_2996_v15_
                nw460_[int(1)] = d_2996_v15_
                nw460_[int(2)] = d_2996_v15_
                nw460_[int(3)] = d_2996_v15_
                nw460_[int(4)] = d_2996_v15_
                nw460_[int(5)] = d_2996_v15_
                nw460_[int(6)] = d_2996_v15_
                nw460_[int(7)] = d_2996_v15_
                nw460_[int(8)] = d_2996_v15_
                def iife234_(_pat_let65_0):
                    def iife235_(d_3000_dt__update__tmp_h0_):
                        def iife236_(_pat_let66_0):
                            def iife237_(d_3001_dt__update_hcf13_h0_):
                                return D1_DC4((d_3000_dt__update__tmp_h0_).cf12, d_3001_dt__update_hcf13_h0_, (d_3000_dt__update__tmp_h0_).cf14)
                            return iife237_(_pat_let66_0)
                        return iife236_(pat_let_tv44_)
                    return iife235_(_pat_let65_0)
                nw460_[int(9)] = iife234_(d_2996_v15_)
                nw460_[int(10)] = d_2996_v15_
                nw460_[int(11)] = default__.fm9(p1, len((d_2997_v16_)[default__.safeIndex((0) - ((self).f8), len(d_2997_v16_))]), d_2978_v4_, d_2998_v17_, globalState)
                d_2999_v18_ = nw460_
                index463_ = default__.safeIndex(995, (d_2999_v18_).length(0))
                (d_2999_v18_)[index463_] = d_2996_v15_
                index464_ = default__.safeIndex(995, (d_2999_v18_).length(0))
                (d_2999_v18_)[index464_] = d_2996_v15_
                d_3002_v19_: _dafny.Map
                d_3002_v19_ = _dafny.Map({d_2977_v3_: d_2985_cf12_})
                d_3003_v20_: _dafny.Map
                d_3003_v20_ = _dafny.Map({d_3002_v19_: len(d_2979_v5_)})
                d_2983_cf14_ = ((d_3003_v20_)[d_3002_v19_] if (d_3002_v19_) in (d_3003_v20_) else d_2984_cf13_)
            elif source38_.is_DC5:
                d_3004___mcc_h3_ = source38_.cf15
                d_3005___mcc_h4_ = source38_.cf16
                d_3006___mcc_h5_ = source38_.cf17
                d_3007_cf17_ = d_3006___mcc_h5_
                d_3008_cf16_ = d_3005___mcc_h4_
                d_3009_cf15_ = d_3004___mcc_h3_
                d_3010_v21_: _dafny.Map
                d_3010_v21_ = _dafny.Map({_dafny.CodePoint('h'): (self).f8})
                d_3011_v22_: str
                d_3011_v22_ = _dafny.CodePoint('d')
                d_3010_v21_ = (d_3010_v21_).set(d_3011_v22_, default__.fm10(False, (self).f8, d_2978_v4_, globalState))
                d_3012_v23_: _dafny.Map
                d_3012_v23_ = _dafny.Map({d_2978_v4_: 717})
                d_3013_v24_: _dafny.MultiSet
                d_3013_v24_ = _dafny.MultiSet([d_3012_v23_])
                d_3009_cf15_ = ((d_3013_v24_ if d_2976_v2_ else d_3013_v24_)).isdisjoint(d_3013_v24_)
                d_3009_cf15_ = (d_2979_v5_)[default__.safeIndex(default__.fm10(not(p1), d_3008_cf16_, d_2978_v4_, globalState), len(d_2979_v5_))]
                d_3008_cf16_ = d_2978_v4_
            elif True:
                d_3014___mcc_h6_ = source38_.cf11
                d_3015_cf11_ = d_3014___mcc_h6_
                d_3016_v25_: bool
                out28_: bool
                out28_ = default__.m0(not(d_2976_v2_), (self).f8, globalState)
                d_3016_v25_ = out28_
                d_3017_v26_: _dafny.Array
                def lambda118_(d_3018_i1_):
                    return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "acq"))])

                init69_ = lambda118_
                nw461_ = _dafny.Array(None, 2)
                for i0_69_ in range(nw461_.length(0)):
                    nw461_[i0_69_] = init69_(i0_69_)
                d_3017_v26_ = nw461_
                d_3019_v27_: _dafny.Seq
                d_3019_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
                d_3020_v28_: _dafny.MultiSet
                d_3020_v28_ = _dafny.MultiSet([d_3019_v27_, d_3019_v27_])
                index465_ = default__.safeIndex(602, (d_3017_v26_).length(0))
                (d_3017_v26_)[index465_] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_3021_i2_ in range(default__.abs(866))]), d_3019_v27_, d_3019_v27_]))) - (d_3020_v28_)
                index466_ = default__.safeIndex(602, (d_3017_v26_).length(0))
                (d_3017_v26_)[index466_] = ((d_3020_v28_) | (d_3020_v28_) if d_3016_v25_ else (d_3020_v28_).intersection(d_3020_v28_))
                d_3022_v29_: _dafny.Set
                d_3022_v29_ = _dafny.Set({True, d_3016_v25_})
                d_3023_v30_: _dafny.Array
                def lambda119_(d_3024_v4_):
                    def lambda120_(d_3025_i3_):
                        return (d_3025_i3_) - (d_3024_v4_)

                    return lambda120_

                init70_ = lambda119_(d_2978_v4_)
                nw462_ = _dafny.Array(None, 11)
                for i0_70_ in range(nw462_.length(0)):
                    nw462_[i0_70_] = init70_(i0_70_)
                d_3023_v30_ = nw462_
                d_3026_v31_: _dafny.Map
                d_3026_v31_ = _dafny.Map({p1: d_3023_v30_})
                d_3027_v32_: _dafny.Seq
                d_3027_v32_ = _dafny.SeqWithoutIsStrInference([d_3023_v30_, d_3023_v30_, d_3023_v30_, d_3023_v30_, ((d_3026_v31_)[not(d_3016_v25_)] if (not(d_3016_v25_)) in (d_3026_v31_) else d_3023_v30_)])
                d_3028_v33_: _dafny.Map
                d_3028_v33_ = _dafny.Map({(len(d_3022_v29_)) > (len(d_3019_v27_)): (d_3027_v32_)[default__.safeIndex(475, len(d_3027_v32_))]})
                d_3029_v34_: str
                d_3029_v34_ = _dafny.CodePoint('a')
                d_3030_v35_: _dafny.Set
                d_3030_v35_ = _dafny.Set({d_3029_v34_})
                d_3028_v33_ = (d_3028_v33_).set((default__.fm10((self).fm6(d_3016_v25_, d_3019_v27_, d_3030_v35_, globalState), (self).f8, -371, globalState)) < ((self).f8), d_3023_v30_)
                d_3031_v36_: _dafny.Map
                d_3031_v36_ = _dafny.Map({(default__.fm10(d_3016_v25_, d_2978_v4_, d_2978_v4_, globalState)) - ((self).f8): d_3023_v30_})
                d_3031_v36_ = (d_3031_v36_).set((243) + (d_2978_v4_), d_3023_v30_)
            d_3032_v37_: _dafny.Array
            def lambda121_(d_3033_i4_):
                return default__.safeDivisionInt(d_3033_i4_, (self).f8)

            init71_ = lambda121_
            nw463_ = _dafny.Array(None, 16)
            for i0_71_ in range(nw463_.length(0)):
                nw463_[i0_71_] = init71_(i0_71_)
            d_3032_v37_ = nw463_
            d_3034_v38_: _dafny.Map
            d_3034_v38_ = _dafny.Map({d_3032_v37_: d_2978_v4_})
            d_3034_v38_ = d_3034_v38_
            d_3035_v39_: str
            d_3035_v39_ = _dafny.CodePoint('r')
            d_3036_v40_: _dafny.Map
            d_3036_v40_ = _dafny.Map({d_3035_v39_: d_2976_v2_})
            d_3037_v41_: _dafny.Map
            d_3037_v41_ = _dafny.Map({d_3036_v40_: d_2979_v5_})
            def iife238_():
                coll104_ = _dafny.Map()
                compr_104_: str
                for compr_104_ in (default__.fm11(globalState)).keys.Elements:
                    d_3038_v42_: str = compr_104_
                    if (d_3038_v42_) in (default__.fm11(globalState)):
                        coll104_[d_3038_v42_] = p1
                return _dafny.Map(coll104_)
            def iife239_():
                coll105_ = _dafny.Map()
                compr_105_: str
                for compr_105_ in (default__.fm11(globalState)).keys.Elements:
                    d_3039_v42_: str = compr_105_
                    if (d_3039_v42_) in (default__.fm11(globalState)):
                        coll105_[d_3039_v42_] = p1
                return _dafny.Map(coll105_)
            if (not(p1) if (d_2976_v2_) and (d_2976_v2_) else (len(((d_3037_v41_)[iife238_()
            ] if (iife239_()
            ) in (d_3037_v41_) else d_2979_v5_))) >= (default__.fm10(p1, (self).f8, d_2978_v4_, globalState))):
                d_3040_v43_: _dafny.Array
                nw464_ = _dafny.Array(False, 8)
                d_3040_v43_ = nw464_
                index467_ = default__.safeIndex(818, (d_3040_v43_).length(0))
                (d_3040_v43_)[index467_] = False
                index468_ = default__.safeIndex(818, (d_3040_v43_).length(0))
                (d_3040_v43_)[index468_] = p1
                d_2978_v4_ = (self).f8
                d_3041_v44_: C5
                nw465_ = C5()
                nw465_.ctor__()
                d_3041_v44_ = nw465_
                d_3042_v45_: _dafny.Map
                d_3042_v45_ = _dafny.Map({p1: d_2978_v4_})
                d_3043_v46_: _dafny.Map
                d_3043_v46_ = d_3042_v45_
                d_3043_v46_ = d_3043_v46_
                d_3040_v43_ = d_3040_v43_
            elif True:
                d_3044_v47_: D16
                d_3044_v47_ = D16_DC41(d_3035_v39_, p1)
                d_3045_v48_: _dafny.Map
                d_3045_v48_ = _dafny.Map({(0) - ((self).f8): (d_3044_v47_).cf77})
                d_3046_v49_: _dafny.Seq
                d_3046_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fd"))
                d_3047_v52_: _dafny.Seq
                d_3047_v52_ = _dafny.SeqWithoutIsStrInference([d_3045_v48_])
                d_3048_v54_: _dafny.Array
                nw466_ = _dafny.Array(None, 16)
                nw466_[int(0)] = d_3045_v48_
                nw466_[int(1)] = (_dafny.Map({597: d_3035_v39_})) | (d_3045_v48_)
                nw466_[int(2)] = d_3045_v48_
                nw466_[int(3)] = d_3045_v48_
                nw466_[int(4)] = d_3045_v48_
                nw466_[int(5)] = d_3045_v48_
                nw466_[int(6)] = (d_3045_v48_) | (default__.fm81(True, len(d_3046_v49_), globalState))
                def iife240_():
                    coll106_ = _dafny.Map()
                    compr_106_: int
                    for compr_106_ in (p0).keys.Elements:
                        d_3049_v50_: int = compr_106_
                        if (d_3049_v50_) in (p0):
                            coll106_[(d_3049_v50_) + (d_2978_v4_)] = d_3035_v39_
                    return _dafny.Map(coll106_)
                nw466_[int(7)] = iife240_()
                
                def iife241_():
                    coll107_ = _dafny.Map()
                    compr_107_: int
                    for compr_107_ in _dafny.IntegerRange(655, 413):
                        d_3050_v51_: int = compr_107_
                        if ((655) <= (d_3050_v51_)) and ((d_3050_v51_) < (413)):
                            coll107_[(d_3050_v51_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqrd"))))] = d_3035_v39_
                    return _dafny.Map(coll107_)
                nw466_[int(8)] = (d_3045_v48_) | (iife241_()
                )
                nw466_[int(9)] = d_3045_v48_
                nw466_[int(10)] = (d_3045_v48_) | ((d_3047_v52_)[default__.safeIndex((self).f8, len(d_3047_v52_))])
                nw466_[int(11)] = d_3045_v48_
                nw466_[int(12)] = (d_3045_v48_) | (d_3045_v48_)
                def iife242_():
                    coll108_ = _dafny.Map()
                    compr_108_: int
                    for compr_108_ in (_dafny.MultiSet(d_2977_v3_)).Elements:
                        d_3051_v53_: int = compr_108_
                        if (d_3051_v53_) in (_dafny.MultiSet(d_2977_v3_)):
                            coll108_[(d_3051_v53_) - ((self).f8)] = d_3035_v39_
                    return _dafny.Map(coll108_)
                nw466_[int(13)] = (iife242_()
                ) | (d_3045_v48_)
                nw466_[int(14)] = (d_3045_v48_) | (d_3045_v48_)
                nw466_[int(15)] = d_3045_v48_
                d_3048_v54_ = nw466_
                d_3052_v55_: _dafny.Map
                d_3052_v55_ = _dafny.Map({d_2976_v2_: (self).f8})
                d_3053_v56_: _dafny.Map
                d_3053_v56_ = _dafny.Map({d_2976_v2_: ((_dafny.Map({(self).f8: d_3035_v39_})).set(len(d_2974_v0_), d_3035_v39_)).set(len(d_3052_v55_), _dafny.CodePoint('w'))})
                index469_ = default__.safeIndex(35, (d_3048_v54_).length(0))
                (d_3048_v54_)[index469_] = ((d_3053_v56_)[d_2976_v2_] if (d_2976_v2_) in (d_3053_v56_) else d_3045_v48_)
                index470_ = default__.safeIndex(35, (d_3048_v54_).length(0))
                rhs421_ = (0) - (d_2978_v4_)
                rhs422_ = d_3045_v48_
                rhs423_ = d_2978_v4_
                lhs237_ = d_3048_v54_
                lhs238_ = default__.safeIndex(35, (d_3048_v54_).length(0))
                d_2978_v4_ = rhs421_
                lhs237_[lhs238_] = rhs422_
                d_2978_v4_ = rhs423_
                d_3054_v57_: _dafny.Array
                def lambda122_(d_3055_v39_):
                    def lambda123_(d_3056_i5_):
                        return d_3055_v39_

                    return lambda123_

                init72_ = lambda122_(d_3035_v39_)
                nw467_ = _dafny.Array(None, 3)
                for i0_72_ in range(nw467_.length(0)):
                    nw467_[i0_72_] = init72_(i0_72_)
                d_3054_v57_ = nw467_
                rhs424_ = d_3054_v57_
                rhs425_ = (d_2978_v4_) - (500)
                d_3054_v57_ = rhs424_
                d_2978_v4_ = rhs425_
                d_3057_v58_: _dafny.Map
                d_3057_v58_ = _dafny.Map({d_2977_v3_: (d_2979_v5_) + (d_2979_v5_)})
                d_3058_v59_: D0
                d_3058_v59_ = D0_DC0(p1, (self).f8, d_2976_v2_, len(d_3046_v49_), p1)
                d_3057_v58_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([default__.fm32(d_3058_v59_, globalState), d_2978_v4_]): _dafny.SeqWithoutIsStrInference([d_2976_v2_, p1])})
                d_3059_v60_: _dafny.MultiSet
                d_3059_v60_ = _dafny.MultiSet([(d_2979_v5_).set(default__.safeIndex(default__.fm32(d_3058_v59_, globalState), len(d_2979_v5_)), d_2976_v2_), d_2979_v5_])
                d_3060_v61_: _dafny.Map
                d_3060_v61_ = _dafny.Map({d_2978_v4_: (d_2978_v4_) > (((d_3059_v60_)[d_2979_v5_] if (d_2979_v5_) in (d_3059_v60_) else d_2978_v4_))})
                d_3060_v61_ = (d_3060_v61_).set(-845, p1)
                index471_ = default__.safeIndex(35, (d_3048_v54_).length(0))
                (d_3048_v54_)[index471_] = ((d_3048_v54_)[default__.safeIndex(35, (d_3048_v54_).length(0))]).set(d_2978_v4_, _dafny.CodePoint('g'))
        elif True:
            r1 = p1
            d_3061_v62_: _dafny.Seq
            d_3061_v62_ = _dafny.SeqWithoutIsStrInference([True, p1])
            d_3062_v63_: _dafny.Seq
            d_3062_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjw"))
            d_3063_v64_: _dafny.Seq
            d_3063_v64_ = _dafny.SeqWithoutIsStrInference([(self).f8])
            d_3064_v65_: C6
            nw468_ = C6()
            nw468_.ctor__((self).f8, d_2976_v2_, (d_3061_v62_)[default__.safeIndex((self).f8, len(d_3061_v62_))], len(d_3062_v63_), len(d_3063_v64_), d_2976_v2_)
            d_3064_v65_ = nw468_
            d_3065_v66_: _dafny.Seq
            d_3065_v66_ = _dafny.SeqWithoutIsStrInference([d_3064_v65_, d_3064_v65_, d_3064_v65_, d_3064_v65_])
            d_3066_v67_: _dafny.Array
            nw469_ = _dafny.Array(None, 21)
            nw469_[int(0)] = d_3064_v65_
            nw469_[int(1)] = d_3064_v65_
            nw469_[int(2)] = d_3064_v65_
            nw469_[int(3)] = d_3064_v65_
            nw469_[int(4)] = d_3064_v65_
            nw469_[int(5)] = (d_3065_v66_)[default__.safeIndex((self).f8, len(d_3065_v66_))]
            nw469_[int(6)] = d_3064_v65_
            nw469_[int(7)] = d_3064_v65_
            nw469_[int(8)] = d_3064_v65_
            nw469_[int(9)] = d_3064_v65_
            nw469_[int(10)] = d_3064_v65_
            nw469_[int(11)] = d_3064_v65_
            nw469_[int(12)] = d_3064_v65_
            nw469_[int(13)] = d_3064_v65_
            nw469_[int(14)] = (d_3064_v65_ if True else d_3064_v65_)
            nw469_[int(15)] = (d_3064_v65_ if p1 else d_3064_v65_)
            nw469_[int(16)] = d_3064_v65_
            nw469_[int(17)] = d_3064_v65_
            nw469_[int(18)] = d_3064_v65_
            nw469_[int(19)] = d_3064_v65_
            nw469_[int(20)] = d_3064_v65_
            d_3066_v67_ = nw469_
            index472_ = default__.safeIndex(940, (d_3066_v67_).length(0))
            (d_3066_v67_)[index472_] = d_3064_v65_
            index473_ = default__.safeIndex(940, (d_3066_v67_).length(0))
            (d_3066_v67_)[index473_] = d_3064_v65_
            d_2976_v2_ = (d_3062_v63_) == (d_3062_v63_)
            d_3067_v68_: str
            d_3067_v68_ = _dafny.CodePoint('g')
            d_3068_v69_: _dafny.Map
            d_3068_v69_ = _dafny.Map({d_3064_v65_.f21: (self).f8})
            d_3069_v70_: _dafny.MultiSet
            d_3069_v70_ = _dafny.MultiSet([len(d_3062_v63_), (self).f8, d_3064_v65_.f21])
            d_3070_v71_: _dafny.Array
            nw470_ = _dafny.Array(None, 12)
            nw470_[int(0)] = (self).f8
            nw470_[int(1)] = 597
            nw470_[int(2)] = (self).f8
            nw470_[int(3)] = (self).f8
            nw470_[int(4)] = (self).f8
            nw470_[int(5)] = ((d_3068_v69_)[(d_3069_v70_).cardinality] if ((d_3069_v70_).cardinality) in (d_3068_v69_) else (d_3069_v70_).cardinality)
            nw470_[int(6)] = d_3064_v65_.f21
            nw470_[int(7)] = len(d_3061_v62_)
            nw470_[int(8)] = 474
            nw470_[int(9)] = (self).f8
            nw470_[int(10)] = (0) - (d_3064_v65_.f21)
            nw470_[int(11)] = (self).f8
            d_3070_v71_ = nw470_
            d_3071_v72_: _dafny.Map
            d_3071_v72_ = _dafny.Map({d_3067_v68_: d_3070_v71_})
            r1 = (d_3071_v72_) != ((_dafny.Map({(d_3062_v63_)[default__.safeIndex(d_3064_v65_.f21, len(d_3062_v63_))]: d_3070_v71_})) | (d_3071_v72_))
            d_3072_v73_: T2
            nw471_ = C4()
            nw471_.ctor__(d_3064_v65_.f21, d_2976_v2_, d_3062_v63_, d_2976_v2_, not(d_2976_v2_), d_3064_v65_.f21, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))), (d_3064_v65_).f22)
            d_3072_v73_ = nw471_
            d_3073_v74_: _dafny.Seq
            d_3073_v74_ = _dafny.SeqWithoutIsStrInference([d_3072_v73_, d_3072_v73_, d_3072_v73_, d_3072_v73_, d_3072_v73_])
            d_3074_v75_: D18
            d_3074_v75_ = D18_DC49((d_3073_v74_)[default__.safeIndex((0) - (d_3064_v65_.f21), len(d_3073_v74_))])
            d_3075_v76_: _dafny.Set
            d_3075_v76_ = _dafny.Set({(d_3074_v75_).cf89})
            if not((d_3075_v76_).isdisjoint(_dafny.Set({d_3072_v73_}))):
                d_3076_v77_: C5
                nw472_ = C5()
                nw472_.ctor__()
                d_3076_v77_ = nw472_
                d_3077_v78_: _dafny.Map
                d_3077_v78_ = _dafny.Map({d_2976_v2_: d_3076_v77_})
                d_3078_v79_: D4
                d_3078_v79_ = D4_DC10(d_3072_v73_.f18, d_2976_v2_, d_3061_v62_)
                d_3079_v80_: _dafny.Map
                d_3079_v80_ = _dafny.Map({((d_3077_v78_)[(d_3064_v65_).f22] if ((d_3064_v65_).f22) in (d_3077_v78_) else d_3076_v77_): (d_3078_v79_).cf24})
                (d_3064_v65_).f21 = default__.fm49(default__.fm28((d_3072_v73_).f17, len(d_3079_v80_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnl"))), globalState), (self).f8, True, (d_3072_v73_).fm0(globalState), globalState)
                d_3080_v82_: _dafny.Seq
                d_3080_v82_ = _dafny.SeqWithoutIsStrInference([d_3068_v69_, d_3068_v69_])
                d_3081_v83_: _dafny.Set
                d_3081_v83_ = _dafny.Set({d_2976_v2_})
                d_3082_v86_: _dafny.Array
                nw473_ = _dafny.Array(None, 26)
                nw473_[int(0)] = d_3068_v69_
                nw473_[int(1)] = d_3068_v69_
                nw473_[int(2)] = _dafny.Map({(0) - ((self).f8): len(d_3062_v63_)})
                nw473_[int(3)] = d_3068_v69_
                nw473_[int(4)] = _dafny.Map({(d_3072_v73_).f5: 758})
                nw473_[int(5)] = _dafny.Map({len(d_3068_v69_): d_3064_v65_.f21})
                nw473_[int(6)] = (d_3068_v69_) | (_dafny.Map({d_3064_v65_.f21: (_dafny.MultiSet([(0) - (d_3064_v65_.f21)])).cardinality}))
                nw473_[int(7)] = default__.fm58(d_3064_v65_.f21, d_3064_v65_.f21, globalState)
                nw473_[int(8)] = d_3068_v69_
                nw473_[int(9)] = _dafny.Map({646: d_3064_v65_.f21})
                def iife243_():
                    coll109_ = _dafny.Map()
                    compr_109_: int
                    for compr_109_ in (p0).keys.Elements:
                        d_3083_v81_: int = compr_109_
                        if (d_3083_v81_) in (p0):
                            coll109_[(d_3083_v81_) + ((self).f8)] = d_3064_v65_.f21
                    return _dafny.Map(coll109_)
                nw473_[int(10)] = iife243_()
                
                nw473_[int(11)] = d_3068_v69_
                nw473_[int(12)] = _dafny.Map({default__.fm26(globalState): default__.fm25(d_3064_v65_.f21, len(d_3062_v63_), p1, globalState)})
                nw473_[int(13)] = d_3068_v69_
                nw473_[int(14)] = d_3068_v69_
                nw473_[int(15)] = _dafny.Map({315: d_3064_v65_.f21})
                nw473_[int(16)] = d_3068_v69_
                nw473_[int(17)] = (d_3068_v69_).set(d_3064_v65_.f21, (d_3063_v64_)[default__.safeIndex(d_3064_v65_.f21, len(d_3063_v64_))])
                nw473_[int(18)] = ((d_3080_v82_)[default__.safeIndex(len(d_3081_v83_), len(d_3080_v82_))]) | (d_3068_v69_)
                nw473_[int(19)] = d_3068_v69_
                def iife244_():
                    coll110_ = _dafny.Map()
                    compr_110_: int
                    for compr_110_ in (d_3063_v64_).Elements:
                        d_3084_v84_: int = compr_110_
                        if (d_3084_v84_) in (d_3063_v64_):
                            coll110_[(d_3084_v84_) - (len(d_3062_v63_))] = len(_dafny.Map({d_3072_v73_.f18: d_2976_v2_}))
                    return _dafny.Map(coll110_)
                nw473_[int(20)] = (iife244_()
                ) | (d_3068_v69_)
                nw473_[int(21)] = d_3068_v69_
                def iife245_():
                    coll111_ = _dafny.Map()
                    compr_111_: int
                    for compr_111_ in _dafny.IntegerRange(117, -710):
                        d_3085_v85_: int = compr_111_
                        if ((117) <= (d_3085_v85_)) and ((d_3085_v85_) < (-710)):
                            coll111_[default__.safeModuloInt(d_3085_v85_, (_dafny.MultiSet(d_3061_v62_)).cardinality)] = d_3064_v65_.f21
                    return _dafny.Map(coll111_)
                nw473_[int(22)] = (iife245_()
                ) | (d_3068_v69_)
                nw473_[int(23)] = d_3068_v69_
                nw473_[int(24)] = d_3068_v69_
                nw473_[int(25)] = default__.fm58(d_3064_v65_.f21, (self).f8, globalState)
                d_3082_v86_ = nw473_
                index474_ = default__.safeIndex(291, (d_3082_v86_).length(0))
                (d_3082_v86_)[index474_] = d_3068_v69_
                index475_ = default__.safeIndex(291, (d_3082_v86_).length(0))
                (d_3082_v86_)[index475_] = d_3068_v69_
                d_3086_v87_: _dafny.MultiSet
                d_3087_v88_: int
                d_3088_v89_: _dafny.MultiSet
                out29_: _dafny.MultiSet
                out30_: int
                out31_: _dafny.MultiSet
                out29_, out30_, out31_ = (d_3076_v77_).m14(globalState)
                d_3086_v87_ = out29_
                d_3087_v88_ = out30_
                d_3088_v89_ = out31_
                (globalState).f0 = (d_3081_v83_).intersection(d_3081_v83_)
                d_3089_v90_: _dafny.MultiSet
                d_3089_v90_ = _dafny.MultiSet([d_3067_v68_, d_3067_v68_])
                d_3090_v91_: _dafny.Array
                def lambda124_(d_3091_i6_):
                    return False

                init73_ = lambda124_
                nw474_ = _dafny.Array(None, 15)
                for i0_73_ in range(nw474_.length(0)):
                    nw474_[i0_73_] = init73_(i0_73_)
                d_3090_v91_ = nw474_
                d_3092_v92_: D18
                d_3092_v92_ = D18_DC50((d_3064_v65_).f22, d_2976_v2_)
                d_3093_v93_: _dafny.Array
                nw475_ = _dafny.Array(None, 11)
                nw475_[int(0)] = (d_3081_v83_).ispropersubset(d_3081_v83_)
                nw475_[int(1)] = (802) <= (d_3087_v88_)
                nw475_[int(2)] = (d_3089_v90_).ispropersubset(d_3089_v90_)
                nw475_[int(3)] = (d_3072_v73_).f4
                nw475_[int(4)] = d_2976_v2_
                nw475_[int(5)] = p1
                nw475_[int(6)] = ((p2)[d_3090_v91_] if (d_3090_v91_) in (p2) else p1)
                nw475_[int(7)] = not((d_3092_v92_).cf90)
                nw475_[int(8)] = (d_3061_v62_) == (d_3061_v62_)
                nw475_[int(9)] = d_3072_v73_.f18
                nw475_[int(10)] = p1
                d_3093_v93_ = nw475_
                d_3090_v91_ = d_3093_v93_
            elif True:
                (d_3072_v73_).f18 = d_3072_v73_.f18
                (d_3064_v65_).f21 = d_3064_v65_.f21
                d_3094_v94_: _dafny.Map
                d_3094_v94_ = _dafny.Map({default__.fm26(globalState): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfdspj"))).set(default__.safeIndex((self).f8, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfdspj")))), default__.fm50(globalState))})
                d_3094_v94_ = (d_3094_v94_).set((self).f8, _dafny.SeqWithoutIsStrInference([d_3067_v68_ for d_3095_i7_ in range(default__.abs(653))]))
                d_3096_v95_: _dafny.Array
                def lambda125_(d_3097_v68_):
                    def lambda126_(d_3098_i8_):
                        return _dafny.SeqWithoutIsStrInference([d_3097_v68_ for d_3099_i9_ in range(default__.abs(-252))])

                    return lambda126_

                init74_ = lambda125_(d_3067_v68_)
                nw476_ = _dafny.Array(None, 22)
                for i0_74_ in range(nw476_.length(0)):
                    nw476_[i0_74_] = init74_(i0_74_)
                d_3096_v95_ = nw476_
                d_3100_v96_: _dafny.Map
                d_3100_v96_ = _dafny.Map({_dafny.CodePoint('y'): False})
                d_3101_v97_: _dafny.Map
                d_3101_v97_ = _dafny.Map({(d_3072_v73_).f5: d_3070_v71_})
                d_3102_v98_: _dafny.MultiSet
                d_3102_v98_ = _dafny.MultiSet([p1, (d_3100_v96_) != (d_3100_v96_), (d_3101_v97_) != (d_3101_v97_), (d_3064_v65_).f22, ((self).fm6(d_3072_v73_.f18, d_3062_v63_, _dafny.Set({d_3067_v68_, d_3067_v68_}), globalState) if not(p1) else (d_3072_v73_).f4)])
                rhs426_ = d_3096_v95_
                rhs427_ = _dafny.MultiSet([(d_3064_v65_).f22])
                rhs428_ = ((d_3072_v73_).f5 if (_dafny.MultiSet(d_3063_v64_)).issubset(_dafny.MultiSet([d_3064_v65_.f21, (d_3072_v73_).f5, -622])) else ((d_3069_v70_)[d_3064_v65_.f21] if (d_3064_v65_.f21) in (d_3069_v70_) else d_3064_v65_.f21))
                lhs239_ = d_3064_v65_
                d_3096_v95_ = rhs426_
                d_3102_v98_ = rhs427_
                lhs239_.f21 = rhs428_
                d_3103_v99_: _dafny.Map
                d_3103_v99_ = _dafny.Map({True: (d_3072_v73_).f5})
                (d_3064_v65_).f21 = len(d_3103_v99_)
        d_3104_v100_: int
        d_3104_v100_ = -332
        d_3104_v100_ = d_3104_v100_
        d_3105_v101_: bool
        out32_: bool
        out32_ = default__.m0(d_2976_v2_, (self).f8, globalState)
        d_3105_v101_ = out32_
        if d_3105_v101_:
            d_3106_v102_: _dafny.Array
            def lambda127_(d_3107_v2_, d_3108_v101_):
                def lambda128_(d_3109_i10_):
                    return (_dafny.MultiSet([_dafny.MultiSet([True])])).ispropersubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_3107_v2_, d_3108_v101_])])))

                return lambda128_

            init75_ = lambda127_(d_2976_v2_, d_3105_v101_)
            nw477_ = _dafny.Array(None, 13)
            for i0_75_ in range(nw477_.length(0)):
                nw477_[i0_75_] = init75_(i0_75_)
            d_3106_v102_ = nw477_
            index476_ = default__.safeIndex(84, (d_3106_v102_).length(0))
            (d_3106_v102_)[index476_] = not (p1) or (p1)
            index477_ = default__.safeIndex(84, (d_3106_v102_).length(0))
            (d_3106_v102_)[index477_] = not(d_2976_v2_)
            d_3110_v103_: _dafny.Array
            nw478_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_3110_v103_ = nw478_
            d_3111_v104_: _dafny.Array
            nw479_ = _dafny.Array(None, 11)
            nw479_[int(0)] = d_3110_v103_
            nw479_[int(1)] = d_3110_v103_
            nw479_[int(2)] = d_3110_v103_
            nw479_[int(3)] = d_3110_v103_
            nw479_[int(4)] = d_3110_v103_
            nw479_[int(5)] = d_3110_v103_
            nw479_[int(6)] = d_3110_v103_
            nw479_[int(7)] = d_3110_v103_
            nw479_[int(8)] = d_3110_v103_
            nw479_[int(9)] = d_3110_v103_
            nw479_[int(10)] = d_3110_v103_
            d_3111_v104_ = nw479_
            index478_ = default__.safeIndex(595, (d_3111_v104_).length(0))
            (d_3111_v104_)[index478_] = d_3110_v103_
            d_3112_v105_: D17
            d_3112_v105_ = D17_DC46()
            d_3113_v106_: _dafny.Set
            d_3113_v106_ = _dafny.Set({p1, (d_3106_v102_)[default__.safeIndex(84, (d_3106_v102_).length(0))]})
            d_3114_v107_: _dafny.Seq
            d_3114_v107_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gqmm"))
            d_3115_v108_: T2
            nw480_ = C7()
            nw480_.ctor__((self).f8, p1, p1, len(d_3113_v106_), d_3104_v100_, d_3105_v101_, d_3114_v107_, not(False))
            d_3115_v108_ = nw480_
            d_3116_v109_: D18
            d_3116_v109_ = D18_DC49(d_3115_v108_)
            pat_let_tv45_ = d_3115_v108_
            pat_let_tv46_ = d_3115_v108_
            d_3117_v110_: _dafny.Array
            nw481_ = _dafny.Array(None, 9)
            def iife246_(_pat_let67_0):
                def iife247_(d_3118_dt__update__tmp_h1_):
                    def iife248_(_pat_let68_0):
                        def iife249_(d_3119_dt__update_hcf89_h0_):
                            return D18_DC49(d_3119_dt__update_hcf89_h0_)
                        return iife249_(_pat_let68_0)
                    return iife248_(pat_let_tv45_)
                return iife247_(_pat_let67_0)
            nw481_[int(0)] = iife246_(D18_DC49(d_3115_v108_))
            def iife250_(_pat_let69_0):
                def iife251_(d_3120_dt__update__tmp_h2_):
                    def iife252_(_pat_let70_0):
                        def iife253_(d_3121_dt__update_hcf89_h1_):
                            return D18_DC49(d_3121_dt__update_hcf89_h1_)
                        return iife253_(_pat_let70_0)
                    return iife252_(pat_let_tv46_)
                return iife251_(_pat_let69_0)
            nw481_[int(1)] = iife250_(D18_DC49(d_3115_v108_))
            nw481_[int(2)] = d_3116_v109_
            nw481_[int(3)] = d_3116_v109_
            nw481_[int(4)] = d_3116_v109_
            nw481_[int(5)] = D18_DC49(d_3115_v108_)
            nw481_[int(6)] = d_3116_v109_
            nw481_[int(7)] = d_3116_v109_
            nw481_[int(8)] = d_3116_v109_
            d_3117_v110_ = nw481_
            index479_ = default__.safeIndex(623, (d_3117_v110_).length(0))
            (d_3117_v110_)[index479_] = d_3116_v109_
            pat_let_tv47_ = d_3115_v108_
            index480_ = default__.safeIndex(595, (d_3111_v104_).length(0))
            index481_ = default__.safeIndex(623, (d_3117_v110_).length(0))
            def iife254_(_pat_let71_0):
                def iife255_(d_3122_dt__update__tmp_h3_):
                    def iife256_(_pat_let72_0):
                        def iife257_(d_3123_dt__update_hcf89_h2_):
                            return D18_DC49(d_3123_dt__update_hcf89_h2_)
                        return iife257_(_pat_let72_0)
                    return iife256_(pat_let_tv47_)
                return iife255_(_pat_let71_0)
            rhs429_ = d_3110_v103_
            rhs430_ = D17_DC46()
            rhs431_ = iife254_((d_3116_v109_ if p1 else d_3116_v109_))
            rhs432_ = d_3106_v102_
            lhs240_ = d_3111_v104_
            lhs241_ = default__.safeIndex(595, (d_3111_v104_).length(0))
            lhs242_ = d_3117_v110_
            lhs243_ = default__.safeIndex(623, (d_3117_v110_).length(0))
            lhs240_[lhs241_] = rhs429_
            d_3112_v105_ = rhs430_
            lhs242_[lhs243_] = rhs431_
            d_3106_v102_ = rhs432_
            d_3104_v100_ = d_3104_v100_
            d_3104_v100_ = 272
            d_3124_v111_: _dafny.MultiSet
            d_3124_v111_ = _dafny.MultiSet([(self).f8, (d_3115_v108_).f5])
            d_3125_v113_: _dafny.MultiSet
            def iife258_():
                coll112_ = _dafny.Set()
                compr_112_: int
                for compr_112_ in (d_3124_v111_).Elements:
                    d_3126_v112_: int = compr_112_
                    if (d_3126_v112_) in (d_3124_v111_):
                        coll112_ = coll112_.union(_dafny.Set([default__.safeModuloInt(d_3126_v112_, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwaufw"))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({885})) for d_3127_i11_ in range(default__.abs(618))]))})))]))
                return _dafny.Set(coll112_)
            d_3125_v113_ = _dafny.MultiSet([len(iife258_()
            ), (self).f8])
            d_3105_v101_ = ((d_3125_v113_).cardinality) < ((d_3115_v108_).f5)
        elif True:
            r1 = d_2976_v2_
            d_3128_v114_: _dafny.Array
            nw482_ = _dafny.Array(int(0), 6)
            d_3128_v114_ = nw482_
            index482_ = default__.safeIndex(770, (d_3128_v114_).length(0))
            (d_3128_v114_)[index482_] = ((self).f8) * (len(default__.fm37(globalState)))
            index483_ = default__.safeIndex(770, (d_3128_v114_).length(0))
            (d_3128_v114_)[index483_] = default__.safeDivisionInt((self).f8, (self).f8)
            d_3129_v115_: _dafny.Array
            nw483_ = _dafny.Array(False, 25)
            d_3129_v115_ = nw483_
            d_3130_v116_: D13
            d_3130_v116_ = D13_DC34(_dafny.SeqWithoutIsStrInference([d_3129_v115_]))
            d_3131_v117_: D13
            d_3131_v117_ = D13_DC34((d_3130_v116_).cf71)
            d_3131_v117_ = (d_3130_v116_ if p1 else d_3131_v117_)
            d_3132_v118_: _dafny.Seq
            d_3132_v118_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jelanq"))
            d_3133_v119_: T0
            nw484_ = C7()
            nw484_.ctor__(d_3104_v100_, d_3105_v101_, not(d_2976_v2_), len(_dafny.SeqWithoutIsStrInference([d_2976_v2_, False])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_3134_i12_ in range(default__.abs(90))])), d_3105_v101_, d_3132_v118_, p1)
            d_3133_v119_ = nw484_
            d_3135_v120_: _dafny.Set
            d_3135_v120_ = _dafny.Set({d_3133_v119_})
            d_3136_v121_: _dafny.Seq
            d_3136_v121_ = _dafny.SeqWithoutIsStrInference([d_3135_v120_, d_3135_v120_, d_3135_v120_, d_3135_v120_])
            d_3104_v100_ = len((_dafny.SeqWithoutIsStrInference([d_3135_v120_])) + ((_dafny.SeqWithoutIsStrInference([_dafny.Set({d_3133_v119_})])) + (d_3136_v121_)))
            def lambda129_(d_3137_v101_):
                def lambda130_(d_3138_i13_):
                    return d_3137_v101_

                return lambda130_

            init76_ = lambda129_(d_3105_v101_)
            nw485_ = _dafny.Array(None, 21)
            for i0_76_ in range(nw485_.length(0)):
                nw485_[i0_76_] = init76_(i0_76_)
            d_3129_v115_ = nw485_
        d_3139_v122_: _dafny.Seq
        d_3139_v122_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgtd"))
        d_3140_v123_: _dafny.Seq
        d_3140_v123_ = _dafny.SeqWithoutIsStrInference([d_3139_v122_])
        r0 = ((d_3140_v123_ if p1 else d_3140_v123_)).set(default__.safeIndex((self).f8, len((d_3140_v123_ if p1 else d_3140_v123_))), d_3139_v122_)
        d_3141_v124_: D0
        d_3141_v124_ = D0_DC0(p1, (self).f8, d_2976_v2_, (self).f8, d_3105_v101_)
        r1 = ((d_3104_v100_) + ((d_3141_v124_).cf3)) >= ((self).f8)
        d_3142_v125_: _dafny.Seq
        d_3142_v125_ = _dafny.SeqWithoutIsStrInference([((0) - (d_3104_v100_)) == (d_3104_v100_)])
        r2 = (d_3142_v125_)[default__.safeIndex(default__.fm18(p1, True, p1, globalState), len(d_3142_v125_))]
        d_3143_v126_: _dafny.Array
        def lambda131_(d_3144_p1_, d_3145_v100_, d_3146_v2_, d_3147_v122_, d_3148_v101_):
            def lambda132_(d_3149_i14_):
                return (_dafny.SeqWithoutIsStrInference([D1_DC5(not(d_3144_p1_), d_3145_v100_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_3150_i15_ in range(default__.abs(702))]))])) + (_dafny.SeqWithoutIsStrInference([D1_DC5((D0_DC1((self).f8, d_3146_v2_, len(_dafny.Map({True: _dafny.CodePoint('c')})), (0) - (d_3145_v100_))).cf6, (self).f8, d_3147_v122_), D1_DC5(d_3144_p1_, d_3145_v100_, d_3147_v122_), D1_DC5(d_3148_v101_, (self).f8, d_3147_v122_), D1_DC5(d_3146_v2_, d_3145_v100_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wapwl")))]))

            return lambda132_

        init77_ = lambda131_(p1, d_3104_v100_, d_2976_v2_, d_3139_v122_, d_3105_v101_)
        nw486_ = _dafny.Array(None, 3)
        for i0_77_ in range(nw486_.length(0)):
            nw486_[i0_77_] = init77_(i0_77_)
        d_3143_v126_ = nw486_
        r3 = d_3143_v126_
        return r0, r1, r2, r3

    @property
    def f8(self):
        return self._f8

class C13(T0):
    def  __init__(self):
        self._f4: bool = False
        self._f5: int = int(0)
        self.f6: bool = False
        self.f7: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C13"
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f6, f7, f4, f5):
        (self).f6 = f6
        (self).f7 = f7
        (self)._f4 = f4
        (self)._f5 = f5

    def fm0(self, globalState):
        return (self).f4

    def fm1(self, p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_3151_i0_ in range(default__.abs(943))])

    def fm2(self, p0, p1, p2, globalState):
        return not(((self).f5) == ((self).f5))

    def m3(self, p0, globalState):
        d_3152_v0_: str
        d_3152_v0_ = _dafny.CodePoint('l')
        d_3153_v1_: _dafny.Seq
        d_3153_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cegvojfny"))
        d_3154_i0_: int
        d_3154_i0_ = 0
        with _dafny.label("10"):
            while (d_3152_v0_) in ((d_3153_v1_) + (d_3153_v1_)):
                with _dafny.c_label("10"):
                    if (d_3154_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_3154_i0_ = (d_3154_i0_) + (1)
                    if not((self).f4):
                        d_3155_v2_: _dafny.Seq
                        d_3155_v2_ = _dafny.SeqWithoutIsStrInference([self.f6])
                        d_3156_v3_: D1
                        d_3156_v3_ = D1_DC3(d_3155_v2_)
                        d_3157_v4_: _dafny.Map
                        d_3157_v4_ = _dafny.Map({d_3156_v3_: d_3155_v2_})
                        d_3157_v4_ = d_3157_v4_
                        d_3158_v5_: int
                        d_3158_v5_ = 334
                        d_3158_v5_ = (0) - ((self).f5)
                        d_3159_v6_: _dafny.MultiSet
                        d_3159_v6_ = _dafny.MultiSet([p0])
                        d_3159_v6_ = d_3159_v6_
                        (self).f7 = (d_3158_v5_) >= ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vuorr")))) + ((0) - (p0)))
                        d_3158_v5_ = p0
                    elif True:
                        d_3152_v0_ = d_3152_v0_
                        d_3160_v7_: _dafny.Array
                        def lambda133_(d_3161_i1_):
                            return default__.safeModuloInt(d_3161_i1_, (self).f5)

                        init78_ = lambda133_
                        nw487_ = _dafny.Array(None, 26)
                        for i0_78_ in range(nw487_.length(0)):
                            nw487_[i0_78_] = init78_(i0_78_)
                        d_3160_v7_ = nw487_
                        index484_ = default__.safeIndex(409, (d_3160_v7_).length(0))
                        (d_3160_v7_)[index484_] = (p0) - (p0)
                        index485_ = default__.safeIndex(409, (d_3160_v7_).length(0))
                        (d_3160_v7_)[index485_] = default__.fm3(globalState)
                        d_3162_v8_: _dafny.MultiSet
                        d_3162_v8_ = _dafny.MultiSet([(self).f4])
                        d_3163_v9_: _dafny.MultiSet
                        d_3163_v9_ = _dafny.MultiSet([d_3162_v8_, d_3162_v8_, default__.fm4(globalState), d_3162_v8_, d_3162_v8_])
                        (self).f6 = ((d_3163_v9_) - (d_3163_v9_)).ispropersubset(d_3163_v9_)
                        index486_ = default__.safeIndex(409, (d_3160_v7_).length(0))
                        (d_3160_v7_)[index486_] = p0
                        d_3164_v10_: _dafny.Array
                        nw488_ = _dafny.Array(_dafny.CodePoint('D'), 27)
                        d_3164_v10_ = nw488_
                        index487_ = default__.safeIndex(764, (d_3164_v10_).length(0))
                        (d_3164_v10_)[index487_] = d_3152_v0_
                        index488_ = default__.safeIndex(764, (d_3164_v10_).length(0))
                        (d_3164_v10_)[index488_] = d_3152_v0_
                    d_3165_v11_: _dafny.Array
                    nw489_ = _dafny.Array(int(0), 9)
                    d_3165_v11_ = nw489_
                    index489_ = default__.safeIndex(229, (d_3165_v11_).length(0))
                    (d_3165_v11_)[index489_] = p0
                    d_3166_v12_: _dafny.MultiSet
                    d_3166_v12_ = _dafny.MultiSet([len(d_3153_v1_)])
                    index490_ = default__.safeIndex(229, (d_3165_v11_).length(0))
                    (d_3165_v11_)[index490_] = (d_3166_v12_).cardinality
                    d_3167_v13_: _dafny.Map
                    d_3167_v13_ = _dafny.Map({d_3165_v11_: (self).f5})
                    (self).f7 = (d_3165_v11_) in (d_3167_v13_)
                    index491_ = default__.safeIndex(229, (d_3165_v11_).length(0))
                    (d_3165_v11_)[index491_] = (((self).f5) * (p0)) + ((self).f5)
                    pass
            pass
        d_3168_v14_: _dafny.Seq
        d_3168_v14_ = _dafny.SeqWithoutIsStrInference([p0])
        d_3169_v15_: C1
        nw490_ = C1()
        nw490_.ctor__(len((d_3168_v14_) + (d_3168_v14_)), self.f7, (634) == ((self).f5), p0)
        d_3169_v15_ = nw490_
        if not(self.f7):
            d_3153_v1_ = (d_3153_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxckdy")))
            d_3170_v16_: _dafny.Seq
            d_3170_v16_ = _dafny.SeqWithoutIsStrInference([d_3153_v1_, d_3153_v1_])
            d_3171_v17_: _dafny.Set
            d_3171_v17_ = _dafny.Set({self.f7})
            d_3172_v18_: _dafny.Map
            d_3172_v18_ = _dafny.Map({d_3170_v16_: d_3171_v17_})
            d_3173_v19_: _dafny.Map
            d_3173_v19_ = _dafny.Map({not((self).f4): _dafny.SeqWithoutIsStrInference([d_3153_v1_ for d_3174_i2_ in range(default__.abs(25))])})
            d_3172_v18_ = (d_3172_v18_).set(((d_3173_v19_)[(self).f4] if ((self).f4) in (d_3173_v19_) else d_3170_v16_), d_3171_v17_)
            d_3175_v20_: int
            d_3175_v20_ = -70
            rhs433_ = (d_3152_v0_ if (self).f4 else d_3152_v0_)
            rhs434_ = p0
            d_3152_v0_ = rhs433_
            d_3175_v20_ = rhs434_
            d_3176_v21_: _dafny.Map
            d_3176_v21_ = _dafny.Map({default__.safeModuloInt(len(d_3153_v1_), len(d_3168_v14_)): (default__.fm10((self).f4, (0) - ((self).f5), 119, globalState)) < (p0)})
            rhs435_ = ((d_3176_v21_)[p0] if (p0) in (d_3176_v21_) else not((d_3171_v17_).ispropersubset(d_3171_v17_)))
            rhs436_ = (self).f5
            lhs244_ = self
            lhs244_.f6 = rhs435_
            d_3175_v20_ = rhs436_
            d_3177_v22_: _dafny.Map
            d_3177_v22_ = _dafny.Map({default__.fm50(globalState): self.f7})
            d_3177_v22_ = d_3177_v22_
        elif True:
            d_3178_v23_: D19
            d_3178_v23_ = D19_DC52(_dafny.Set({self.f6, self.f7}))
            d_3179_v24_: _dafny.Map
            d_3179_v24_ = _dafny.Map({d_3178_v23_: (0) - (p0)})
            d_3179_v24_ = d_3179_v24_
            d_3180_v25_: _dafny.Set
            d_3180_v25_ = _dafny.Set({self.f6, self.f7, self.f7})
            d_3181_v26_: _dafny.Map
            d_3181_v26_ = _dafny.Map({(d_3180_v25_).intersection(d_3180_v25_): p0})
            d_3181_v26_ = (d_3181_v26_).set((d_3180_v25_) | (_dafny.Set({(self).f4})), (0) - (p0))
            (self).f6 = (self).f4
            if ((self).f5) != (default__.fm49(d_3153_v1_, (self).f5, self.f7, (self).f4, globalState)):
                d_3182_v27_: _dafny.Array
                nw491_ = _dafny.Array(None, 2)
                nw491_[int(0)] = (self).f4
                nw491_[int(1)] = ((self).f4 if (self).f4 else (self).f4)
                d_3182_v27_ = nw491_
                index492_ = default__.safeIndex(610, (d_3182_v27_).length(0))
                (d_3182_v27_)[index492_] = not ((self).f4) or (self.f7)
                d_3183_v28_: C2
                nw492_ = C2()
                nw492_.ctor__(d_3182_v27_, (default__.fm68(globalState)).set(p0, default__.abs((0) - (len(d_3153_v1_)))), d_3153_v1_, self.f7, self.f7, (self).f5)
                d_3183_v28_ = nw492_
                d_3184_v29_: _dafny.Map
                d_3184_v29_ = _dafny.Map({self.f7: d_3183_v28_})
                d_3185_v30_: C8
                nw493_ = C8()
                nw493_.ctor__(d_3153_v1_, p0)
                d_3185_v30_ = nw493_
                d_3186_v31_: _dafny.Array
                nw494_ = _dafny.Array(None, 15)
                nw494_[int(0)] = d_3185_v30_
                nw494_[int(1)] = d_3185_v30_
                nw494_[int(2)] = d_3185_v30_
                nw494_[int(3)] = d_3185_v30_
                nw494_[int(4)] = d_3185_v30_
                nw494_[int(5)] = d_3185_v30_
                nw494_[int(6)] = d_3185_v30_
                nw494_[int(7)] = d_3185_v30_
                nw494_[int(8)] = (d_3185_v30_ if True else d_3185_v30_)
                nw494_[int(9)] = d_3185_v30_
                nw494_[int(10)] = d_3185_v30_
                nw494_[int(11)] = (d_3185_v30_ if self.f7 else d_3185_v30_)
                nw494_[int(12)] = (D21_DC55(d_3185_v30_)).cf99
                nw494_[int(13)] = d_3185_v30_
                nw494_[int(14)] = d_3185_v30_
                d_3186_v31_ = nw494_
                d_3187_v32_: _dafny.Seq
                d_3187_v32_ = _dafny.SeqWithoutIsStrInference([(self).f4, self.f6])
                d_3188_v33_: _dafny.Map
                d_3188_v33_ = _dafny.Map({(self).f5: (self).f5})
                d_3189_v34_: _dafny.Seq
                d_3189_v34_ = _dafny.SeqWithoutIsStrInference([(d_3185_v30_).fm17(d_3187_v32_, 733, (self).f4, len(d_3188_v33_), globalState), self.f7, not ((self).f4) or (self.f7)])
                d_3190_v35_: _dafny.Seq
                d_3190_v35_ = _dafny.SeqWithoutIsStrInference([d_3184_v29_, d_3184_v29_, d_3184_v29_])
                index493_ = default__.safeIndex(610, (d_3182_v27_).length(0))
                rhs437_ = (d_3187_v32_)[default__.safeIndex((p0) * ((0) - ((0) - (p0))), len(d_3187_v32_))]
                rhs438_ = ((_dafny.Map({self.f7: d_3183_v28_})) | ((d_3190_v35_)[default__.safeIndex((0) - ((d_3185_v30_).f14), len(d_3190_v35_))])) | (d_3184_v29_)
                rhs439_ = not ((self.f7 if self.f6 else self.f7)) or (((0) - (p0)) != (p0))
                rhs440_ = d_3186_v31_
                lhs245_ = d_3182_v27_
                lhs246_ = default__.safeIndex(610, (d_3182_v27_).length(0))
                lhs247_ = self
                lhs245_[lhs246_] = rhs437_
                d_3184_v29_ = rhs438_
                lhs247_.f6 = rhs439_
                d_3186_v31_ = rhs440_
                d_3153_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ayl"))
                d_3188_v33_ = (d_3188_v33_).set((self).f5, (d_3185_v30_).f14)
                d_3191_v36_: _dafny.Set
                d_3191_v36_ = _dafny.Set({d_3188_v33_, d_3188_v33_})
                d_3192_v37_: _dafny.Map
                d_3192_v37_ = _dafny.Map({False: len(d_3180_v25_)})
                d_3193_v38_: D4
                d_3193_v38_ = D4_DC9(d_3192_v37_, (0) - (len(d_3168_v14_)), self.f7)
                d_3194_v39_: C0
                nw495_ = C0()
                nw495_.ctor__(d_3191_v36_, (d_3193_v38_).cf22)
                d_3194_v39_ = nw495_
                d_3195_v40_: int
                d_3195_v40_ = -875
                d_3195_v40_ = -670
            elif True:
                d_3196_v41_: D9
                d_3196_v41_ = D9_DC23(self.f6, (self).f5, self.f6)
                d_3197_v42_: _dafny.Set
                d_3197_v42_ = _dafny.Set({(self).f5})
                d_3198_v43_: _dafny.Map
                d_3198_v43_ = _dafny.Map({d_3196_v41_: len(d_3197_v42_)})
                d_3199_v44_: _dafny.Seq
                d_3199_v44_ = _dafny.SeqWithoutIsStrInference([d_3198_v43_, _dafny.Map({d_3196_v41_: (self).f5})])
                d_3200_v45_: _dafny.Map
                d_3200_v45_ = _dafny.Map({(self).f5: ((default__.fm82(d_3152_v0_, globalState)).set(default__.safeIndex(p0, len(default__.fm82(d_3152_v0_, globalState))), _dafny.Map({d_3196_v41_: (self).f5}))) <= (d_3199_v44_)})
                d_3200_v45_ = d_3200_v45_
                d_3201_v46_: int
                d_3201_v46_ = 432
                d_3202_v47_: _dafny.Map
                d_3202_v47_ = _dafny.Map({self.f7: self.f6})
                d_3201_v46_ = len(_dafny.Set({(len(d_3202_v47_)) != (d_3201_v46_), (self).f4, ((self).f5) not in (_dafny.SeqWithoutIsStrInference([(self).f5 for d_3203_i3_ in range(default__.abs(56))]))}))
                rhs441_ = self.f6
                rhs442_ = _dafny.CodePoint('l')
                lhs248_ = self
                lhs248_.f6 = rhs441_
                d_3152_v0_ = rhs442_
                (self).f6 = (self).f4
                d_3201_v46_ = (0) - ((p0) * ((self).f5))
            d_3204_v48_: _dafny.Map
            d_3204_v48_ = _dafny.Map({self.f6: d_3153_v1_})
            d_3153_v1_ = ((d_3204_v48_)[(self).f4] if ((self).f4) in (d_3204_v48_) else d_3153_v1_)
        d_3205_v49_: _dafny.MultiSet
        d_3205_v49_ = _dafny.MultiSet([self.f6])
        d_3206_v50_: _dafny.MultiSet
        d_3206_v50_ = (d_3205_v49_) - (d_3205_v49_)
        d_3206_v50_ = d_3206_v50_
        d_3207_v51_: D12
        d_3207_v51_ = D12_DC33(self.f6, self.f7, d_3152_v0_, ((self).f5) - ((self).f5))
        source39_ = d_3207_v51_
        if source39_.is_DC33:
            d_3208___mcc_h0_ = source39_.cf67
            d_3209___mcc_h1_ = source39_.cf68
            d_3210___mcc_h2_ = source39_.cf69
            d_3211___mcc_h3_ = source39_.cf70
            d_3212_cf70_ = d_3211___mcc_h3_
            d_3213_cf69_ = d_3210___mcc_h2_
            d_3214_cf68_ = d_3209___mcc_h1_
            d_3215_cf67_ = d_3208___mcc_h0_
            d_3212_cf70_ = len(default__.fm37(globalState))
            d_3214_cf68_ = d_3214_cf68_
            d_3212_cf70_ = (d_3212_cf70_ if self.f7 else default__.safeDivisionInt((self).f5, p0))
            d_3216_v52_: C8
            nw496_ = C8()
            nw496_.ctor__(d_3153_v1_, d_3212_cf70_)
            d_3216_v52_ = nw496_
            d_3217_v53_: _dafny.Seq
            d_3217_v53_ = _dafny.SeqWithoutIsStrInference([d_3216_v52_])
            d_3218_v54_: _dafny.Map
            d_3218_v54_ = _dafny.Map({d_3217_v53_: self.f6})
            d_3219_v55_: _dafny.Map
            d_3219_v55_ = _dafny.Map({self.f6: len(d_3218_v54_)})
            d_3219_v55_ = (d_3219_v55_).set((self).f4, d_3212_cf70_)
        elif True:
            d_3220___mcc_h4_ = source39_.cf66
            d_3221_cf66_ = d_3220___mcc_h4_
            d_3152_v0_ = d_3152_v0_
            d_3222_v56_: D9
            d_3222_v56_ = D9_DC23(False, (self).f5, self.f6)
            d_3223_v57_: _dafny.MultiSet
            d_3223_v57_ = _dafny.MultiSet([407])
            d_3224_v58_: _dafny.Map
            d_3224_v58_ = _dafny.Map({p0: p0})
            d_3225_v59_: _dafny.Map
            d_3225_v59_ = d_3224_v58_
            d_3226_v60_: _dafny.Seq
            d_3226_v60_ = _dafny.SeqWithoutIsStrInference([d_3225_v59_, d_3225_v59_])
            d_3227_v61_: _dafny.Seq
            d_3227_v61_ = _dafny.SeqWithoutIsStrInference([self.f6])
            d_3228_v62_: _dafny.Array
            nw497_ = _dafny.Array(None, 22)
            nw497_[int(0)] = p0
            nw497_[int(1)] = p0
            nw497_[int(2)] = (d_3222_v56_).cf46
            nw497_[int(3)] = len(_dafny.Map({p0: (self).f4}))
            nw497_[int(4)] = default__.safeDivisionInt((self).f5, (self).f5)
            nw497_[int(5)] = 436
            nw497_[int(6)] = len(d_3168_v14_)
            nw497_[int(7)] = (d_3223_v57_).cardinality
            nw497_[int(8)] = p0
            nw497_[int(9)] = (d_3205_v49_).cardinality
            nw497_[int(10)] = (0) - ((len(d_3226_v60_)) * (-413))
            nw497_[int(11)] = p0
            nw497_[int(12)] = len((_dafny.SeqWithoutIsStrInference([(self).f4, not((self).f4), (self).f4, self.f6])) + (d_3227_v61_))
            nw497_[int(13)] = len(d_3153_v1_)
            nw497_[int(14)] = len(_dafny.Set({_dafny.Map({(self).f5: p0})}))
            nw497_[int(15)] = p0
            nw497_[int(16)] = (self).f5
            nw497_[int(17)] = (self).f5
            nw497_[int(18)] = len(((d_3153_v1_) + (d_3153_v1_)).set(default__.safeIndex(len(default__.fm48(p0, globalState)), len((d_3153_v1_) + (d_3153_v1_))), d_3152_v0_))
            nw497_[int(19)] = 192
            nw497_[int(20)] = (self).f5
            nw497_[int(21)] = p0
            d_3228_v62_ = nw497_
            index494_ = default__.safeIndex(819, (d_3228_v62_).length(0))
            (d_3228_v62_)[index494_] = 565
            d_3229_v63_: _dafny.Array
            nw498_ = _dafny.Array(False, 25)
            d_3229_v63_ = nw498_
            d_3230_v64_: _dafny.Map
            d_3230_v64_ = _dafny.Map({d_3229_v63_: p0})
            index495_ = default__.safeIndex(819, (d_3228_v62_).length(0))
            (d_3228_v62_)[index495_] = (0) - (default__.fm18(not(((self).f5) < (len(d_3230_v64_))), self.f7, self.f7, globalState))
            index496_ = default__.safeIndex(819, (d_3228_v62_).length(0))
            (d_3228_v62_)[index496_] = (_dafny.MultiSet([p0])).cardinality
            d_3231_v65_: _dafny.MultiSet
            d_3231_v65_ = _dafny.MultiSet([d_3228_v62_])
            index497_ = default__.safeIndex(819, (d_3228_v62_).length(0))
            (d_3228_v62_)[index497_] = len((_dafny.Map({d_3205_v49_: (0) - ((d_3231_v65_).cardinality)})).set((d_3205_v49_) | (d_3205_v49_), default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "go"))), (d_3228_v62_)[default__.safeIndex(819, (d_3228_v62_).length(0))])))
        d_3232_v66_: _dafny.Array
        def lambda134_(d_3233_v1_):
            def lambda135_(d_3234_i5_):
                return d_3233_v1_

            return lambda135_

        init79_ = lambda134_(d_3153_v1_)
        nw499_ = _dafny.Array(None, 1)
        for i0_79_ in range(nw499_.length(0)):
            nw499_[i0_79_] = init79_(i0_79_)
        d_3232_v66_ = nw499_
        guard_loop_10_: int
        for guard_loop_10_ in _dafny.IntegerRange(0, (d_3232_v66_).length(0)):
            d_3235_i4_: int = guard_loop_10_
            if (True) and (((0) <= (d_3235_i4_)) and ((d_3235_i4_) < ((d_3232_v66_).length(0)))):
                (d_3232_v66_)[(d_3235_i4_)] = d_3153_v1_

    def m4(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        d_3236_v0_: int
        d_3236_v0_ = 322
        d_3236_v0_ = ((self).f5) - (764)
        d_3237_v1_: _dafny.Seq
        d_3237_v1_ = _dafny.SeqWithoutIsStrInference([(self).f4, (self).f4, p3])
        d_3238_i0_: int
        d_3238_i0_ = 0
        with _dafny.label("11"):
            while not ((d_3237_v1_)[default__.safeIndex(d_3236_v0_, len(d_3237_v1_))]) or (True):
                with _dafny.c_label("11"):
                    if (d_3238_i0_) >= (100):
                        raise _dafny.Break("11")
                    d_3238_i0_ = (d_3238_i0_) + (1)
                    index498_ = default__.safeIndex(757, (p1).length(0))
                    (p1)[index498_] = p3
                    d_3239_v2_: str
                    d_3239_v2_ = _dafny.CodePoint('c')
                    index499_ = default__.safeIndex(757, (p1).length(0))
                    rhs443_ = not(self.f7)
                    rhs444_ = (d_3236_v0_) * (((p0)[False] if (False) in (p0) else d_3236_v0_))
                    rhs445_ = self.f7
                    rhs446_ = d_3239_v2_
                    lhs249_ = p1
                    lhs250_ = default__.safeIndex(757, (p1).length(0))
                    lhs251_ = self
                    lhs249_[lhs250_] = rhs443_
                    d_3236_v0_ = rhs444_
                    lhs251_.f6 = rhs445_
                    d_3239_v2_ = rhs446_
                    d_3236_v0_ = (d_3236_v0_) - ((self).f5)
                    index500_ = default__.safeIndex(757, (p1).length(0))
                    (p1)[index500_] = (self).fm2(d_3236_v0_, d_3236_v0_, (self).f5, globalState)
                    d_3240_v3_: D0
                    d_3240_v3_ = D0_DC0(p3, 546, self.f6, (0) - ((self).f5), p3)
                    d_3241_v4_: _dafny.Seq
                    d_3241_v4_ = _dafny.SeqWithoutIsStrInference([default__.fm44((self).f5, d_3240_v3_, p3, (self).f4, globalState)])
                    d_3242_v5_: _dafny.MultiSet
                    d_3242_v5_ = _dafny.MultiSet([(self).f5, d_3236_v0_])
                    d_3243_v6_: _dafny.Seq
                    d_3243_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xspnu"))
                    d_3244_v7_: _dafny.Seq
                    d_3244_v7_ = _dafny.SeqWithoutIsStrInference([d_3236_v0_, 113, len(d_3237_v1_), (self).f5, (self).f5])
                    d_3245_v8_: _dafny.Array
                    nw500_ = _dafny.Array(None, 16)
                    nw500_[int(0)] = d_3236_v0_
                    nw500_[int(1)] = d_3236_v0_
                    nw500_[int(2)] = (d_3242_v5_).cardinality
                    nw500_[int(3)] = d_3236_v0_
                    nw500_[int(4)] = (self).f5
                    nw500_[int(5)] = (0) - (d_3236_v0_)
                    nw500_[int(6)] = ((p0)[(p1)[default__.safeIndex(757, (p1).length(0))]] if ((p1)[default__.safeIndex(757, (p1).length(0))]) in (p0) else -964)
                    nw500_[int(7)] = d_3236_v0_
                    nw500_[int(8)] = d_3236_v0_
                    nw500_[int(9)] = (self).f5
                    nw500_[int(10)] = len(d_3243_v6_)
                    nw500_[int(11)] = (self).f5
                    nw500_[int(12)] = d_3236_v0_
                    nw500_[int(13)] = len(d_3243_v6_)
                    nw500_[int(14)] = len(d_3244_v7_)
                    nw500_[int(15)] = len(d_3243_v6_)
                    d_3245_v8_ = nw500_
                    d_3246_v9_: _dafny.Map
                    d_3246_v9_ = _dafny.Map({len(d_3241_v4_): d_3245_v8_})
                    d_3246_v9_ = (d_3246_v9_).set(d_3236_v0_, d_3245_v8_)
                    pass
            pass
        d_3247_v10_: C9
        nw501_ = C9()
        nw501_.ctor__(p1)
        d_3247_v10_ = nw501_
        if p2:
            d_3248_v11_: _dafny.MultiSet
            d_3248_v11_ = _dafny.MultiSet([(self).f5])
            d_3249_v12_: _dafny.Seq
            d_3249_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oyelcnck"))
            d_3250_v13_: C2
            nw502_ = C2()
            nw502_.ctor__(p1, d_3248_v11_, d_3249_v12_, not(not(p3)), (self).f4, (0) - (default__.safeModuloInt(d_3236_v0_, d_3236_v0_)))
            d_3250_v13_ = nw502_
            d_3250_v13_ = d_3250_v13_
            r1 = p2
            d_3236_v0_ = 846
            d_3251_v14_: _dafny.Seq
            d_3251_v14_ = _dafny.SeqWithoutIsStrInference([d_3236_v0_, d_3236_v0_])
            d_3252_v15_: _dafny.Array
            nw503_ = _dafny.Array(None, 8)
            nw503_[int(0)] = d_3236_v0_
            nw503_[int(1)] = (self).f5
            nw503_[int(2)] = d_3236_v0_
            nw503_[int(3)] = (self).f5
            nw503_[int(4)] = d_3236_v0_
            nw503_[int(5)] = len(d_3251_v14_)
            nw503_[int(6)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkmnv")))
            nw503_[int(7)] = d_3236_v0_
            d_3252_v15_ = nw503_
            d_3253_v16_: _dafny.Set
            d_3253_v16_ = _dafny.Set({d_3252_v15_})
            if (d_3253_v16_).ispropersubset(_dafny.Set({d_3252_v15_, d_3252_v15_, d_3252_v15_})):
                (self).f6 = p2
                (self).f7 = p3
                d_3254_v17_: C12
                nw504_ = C12()
                nw504_.ctor__(((self).f5 if p3 else -174))
                d_3254_v17_ = nw504_
                d_3255_v18_: _dafny.Array
                nw505_ = _dafny.Array(False, 7)
                d_3255_v18_ = nw505_
                rhs447_ = _dafny.SeqWithoutIsStrInference([261])
                rhs448_ = (d_3247_v10_).f12
                d_3251_v14_ = rhs447_
                d_3255_v18_ = rhs448_
                d_3256_v19_: _dafny.Set
                d_3256_v19_ = _dafny.Set({self.f7})
                d_3257_v20_: D19
                d_3257_v20_ = D19_DC52((d_3256_v19_) - (d_3256_v19_))
                d_3257_v20_ = d_3257_v20_
            elif True:
                d_3237_v1_ = d_3237_v1_
                d_3258_v21_: C12
                nw506_ = C12()
                nw506_.ctor__(len(d_3249_v12_))
                d_3258_v21_ = nw506_
                d_3236_v0_ = (0) - (default__.safeModuloInt(d_3236_v0_, default__.fm25((d_3258_v21_).f8, -994, (self).f4, globalState)))
                d_3259_v22_: D1
                d_3259_v22_ = D1_DC5(p2, (self).f5, d_3249_v12_)
                d_3260_v23_: str
                d_3260_v23_ = _dafny.CodePoint('o')
                pat_let_tv48_ = p2
                d_3261_v24_: _dafny.Array
                nw507_ = _dafny.Array(None, 10)
                nw507_[int(0)] = d_3249_v12_
                nw507_[int(1)] = d_3249_v12_
                def iife259_(_pat_let73_0):
                    def iife260_(d_3262_dt__update__tmp_h0_):
                        def iife261_(_pat_let74_0):
                            def iife262_(d_3263_dt__update_hcf15_h0_):
                                return D1_DC5(d_3263_dt__update_hcf15_h0_, (d_3262_dt__update__tmp_h0_).cf16, (d_3262_dt__update__tmp_h0_).cf17)
                            return iife262_(_pat_let74_0)
                        return iife261_(pat_let_tv48_)
                    return iife260_(_pat_let73_0)
                nw507_[int(2)] = (iife259_(d_3259_v22_)).cf17
                nw507_[int(3)] = d_3249_v12_
                nw507_[int(4)] = d_3249_v12_
                nw507_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkbywo"))
                nw507_[int(6)] = (d_3249_v12_).set(default__.safeIndex(d_3236_v0_, len(d_3249_v12_)), d_3260_v23_)
                nw507_[int(7)] = (d_3249_v12_).set(default__.safeIndex((self).f5, len(d_3249_v12_)), d_3260_v23_)
                nw507_[int(8)] = d_3249_v12_
                nw507_[int(9)] = d_3249_v12_
                d_3261_v24_ = nw507_
                d_3264_v25_: int
                d_3265_v26_: bool
                d_3266_v27_: bool
                d_3267_v28_: int
                out33_: int
                out34_: bool
                out35_: bool
                out36_: int
                out33_, out34_, out35_, out36_ = (d_3247_v10_).m9(d_3261_v24_, globalState)
                d_3264_v25_ = out33_
                d_3265_v26_ = out34_
                d_3266_v27_ = out35_
                d_3267_v28_ = out36_
                d_3260_v23_ = d_3260_v23_
            if self.f7:
                d_3268_v29_: _dafny.Map
                d_3268_v29_ = _dafny.Map({(self).f5: (self).f5})
                d_3269_v30_: _dafny.Map
                d_3269_v30_ = _dafny.Map({_dafny.CodePoint('o'): d_3236_v0_})
                d_3270_v31_: str
                d_3270_v31_ = _dafny.CodePoint('w')
                d_3271_v32_: D12
                d_3271_v32_ = D12_DC33(self.f6, self.f6, d_3270_v31_, len(d_3249_v12_))
                d_3272_v33_: _dafny.Set
                d_3272_v33_ = _dafny.Set({(d_3268_v29_).set((self).f5, (0) - ((self).f5)), (_dafny.Map({(0) - (((d_3269_v30_)[(d_3271_v32_).cf69] if ((d_3271_v32_).cf69) in (d_3269_v30_) else d_3236_v0_)): d_3236_v0_})).set(len(_dafny.Set({d_3236_v0_})), d_3236_v0_)})
                d_3273_v34_: _dafny.Set
                d_3273_v34_ = d_3272_v33_
                d_3274_v35_: C0
                nw508_ = C0()
                nw508_.ctor__(d_3273_v34_, (d_3236_v0_) * ((0) - (len(d_3251_v14_))))
                d_3274_v35_ = nw508_
                nw509_ = C0()
                nw509_.ctor__(d_3274_v35_.f26, (d_3274_v35_).f27)
                d_3274_v35_ = nw509_
                d_3275_v36_: _dafny.Set
                d_3275_v36_ = _dafny.Set({(d_3274_v35_).f27})
                d_3276_v37_: _dafny.Array
                nw510_ = _dafny.Array(None, 25)
                nw510_[int(0)] = d_3249_v12_
                nw510_[int(1)] = d_3249_v12_
                nw510_[int(2)] = d_3249_v12_
                nw510_[int(3)] = d_3249_v12_
                nw510_[int(4)] = d_3249_v12_
                nw510_[int(5)] = d_3249_v12_
                nw510_[int(6)] = (_dafny.SeqWithoutIsStrInference([d_3270_v31_ for d_3277_i1_ in range(default__.abs(523))])).set(default__.safeIndex(d_3236_v0_, len(_dafny.SeqWithoutIsStrInference([d_3270_v31_ for d_3278_i1_ in range(default__.abs(523))]))), d_3270_v31_)
                nw510_[int(7)] = (d_3249_v12_).set(default__.safeIndex((self).f5, len(d_3249_v12_)), d_3270_v31_)
                nw510_[int(8)] = d_3249_v12_
                nw510_[int(9)] = d_3249_v12_
                nw510_[int(10)] = d_3249_v12_
                nw510_[int(11)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_3279_i2_ in range(default__.abs(739))])
                nw510_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idfrp"))
                nw510_[int(13)] = d_3249_v12_
                nw510_[int(14)] = d_3249_v12_
                nw510_[int(15)] = d_3249_v12_
                nw510_[int(16)] = d_3249_v12_
                nw510_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xupvukce"))
                nw510_[int(18)] = d_3249_v12_
                nw510_[int(19)] = d_3249_v12_
                nw510_[int(20)] = d_3249_v12_
                nw510_[int(21)] = d_3249_v12_
                nw510_[int(22)] = default__.fm52(default__.fm27(d_3270_v31_, globalState), d_3275_v36_, (self).f5, p2, globalState)
                nw510_[int(23)] = d_3249_v12_
                nw510_[int(24)] = d_3249_v12_
                d_3276_v37_ = nw510_
                d_3280_v38_: int
                d_3281_v39_: bool
                d_3282_v40_: bool
                d_3283_v41_: int
                out37_: int
                out38_: bool
                out39_: bool
                out40_: int
                out37_, out38_, out39_, out40_ = (d_3247_v10_).m9(d_3276_v37_, globalState)
                d_3280_v38_ = out37_
                d_3281_v39_ = out38_
                d_3282_v40_ = out39_
                d_3283_v41_ = out40_
                d_3284_v42_: D0
                d_3284_v42_ = D0_DC0(self.f7, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdpf"))), p2, (self).f5, (self).f4)
                d_3285_v43_: _dafny.Map
                d_3285_v43_ = _dafny.Map({(default__.fm44(d_3280_v38_, D0_DC0((self).f4, len(d_3251_v14_), not(self.f7), default__.fm32(d_3284_v42_, globalState), d_3282_v40_), p3, d_3281_v39_, globalState) if self.f7 else d_3251_v14_): d_3280_v38_})
                d_3285_v43_ = (d_3285_v43_).set(d_3251_v14_, d_3283_v41_)
                d_3286_v44_: _dafny.Map
                d_3286_v44_ = _dafny.Map({len(d_3249_v12_): (d_3247_v10_).f12})
                d_3287_v45_: _dafny.Map
                d_3287_v45_ = _dafny.Map({(d_3286_v44_) | (d_3286_v44_): (d_3283_v41_) <= (-464)})
                d_3282_v40_ = ((d_3287_v45_)[_dafny.Map({len(d_3251_v14_): (d_3247_v10_).f12})] if (_dafny.Map({len(d_3251_v14_): (d_3247_v10_).f12})) in (d_3287_v45_) else self.f7)
                d_3288_v46_: _dafny.Map
                d_3288_v46_ = _dafny.Map({(d_3274_v35_).f27: d_3270_v31_})
                d_3289_v47_: _dafny.Array
                nw511_ = _dafny.Array(None, 27)
                nw511_[int(0)] = d_3270_v31_
                nw511_[int(1)] = d_3270_v31_
                nw511_[int(2)] = _dafny.CodePoint('n')
                nw511_[int(3)] = d_3270_v31_
                nw511_[int(4)] = d_3270_v31_
                nw511_[int(5)] = d_3270_v31_
                nw511_[int(6)] = (d_3249_v12_)[default__.safeIndex((default__.fm83(self.f6, globalState)).cf46, len(d_3249_v12_))]
                nw511_[int(7)] = d_3270_v31_
                nw511_[int(8)] = d_3270_v31_
                nw511_[int(9)] = d_3270_v31_
                nw511_[int(10)] = d_3270_v31_
                nw511_[int(11)] = d_3270_v31_
                nw511_[int(12)] = d_3270_v31_
                nw511_[int(13)] = d_3270_v31_
                nw511_[int(14)] = d_3270_v31_
                nw511_[int(15)] = d_3270_v31_
                nw511_[int(16)] = d_3270_v31_
                nw511_[int(17)] = d_3270_v31_
                nw511_[int(18)] = d_3270_v31_
                nw511_[int(19)] = d_3270_v31_
                nw511_[int(20)] = (d_3249_v12_)[default__.safeIndex((self).f5, len(d_3249_v12_))]
                nw511_[int(21)] = d_3270_v31_
                nw511_[int(22)] = ((d_3288_v46_)[(self).f5] if ((self).f5) in (d_3288_v46_) else d_3270_v31_)
                nw511_[int(23)] = d_3270_v31_
                nw511_[int(24)] = d_3270_v31_
                nw511_[int(25)] = d_3270_v31_
                nw511_[int(26)] = d_3270_v31_
                d_3289_v47_ = nw511_
                index501_ = default__.safeIndex(595, (d_3289_v47_).length(0))
                (d_3289_v47_)[index501_] = d_3270_v31_
                index502_ = default__.safeIndex(595, (d_3289_v47_).length(0))
                (d_3289_v47_)[index502_] = d_3270_v31_
            elif True:
                index503_ = default__.safeIndex(608, (d_3252_v15_).length(0))
                (d_3252_v15_)[index503_] = (self).f5
                index504_ = default__.safeIndex(608, (d_3252_v15_).length(0))
                (d_3252_v15_)[index504_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brfnjllm")))
                d_3290_v48_: _dafny.MultiSet
                d_3290_v48_ = _dafny.MultiSet([d_3249_v12_])
                d_3291_v49_: C11
                nw512_ = C11()
                nw512_.ctor__(default__.fm84(globalState), (d_3290_v48_).issubset(d_3290_v48_), ((self).f5) + ((self).f5))
                d_3291_v49_ = nw512_
                d_3292_v50_: _dafny.Seq
                d_3292_v50_ = _dafny.SeqWithoutIsStrInference([d_3249_v12_, d_3249_v12_, d_3249_v12_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_3293_i3_ in range(default__.abs(-439))]), d_3249_v12_])
                d_3249_v12_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "guado"))) + ((d_3292_v50_)[default__.safeIndex((self).f5, len(d_3292_v50_))])
                index505_ = default__.safeIndex(608, (d_3252_v15_).length(0))
                (d_3252_v15_)[index505_] = (0) - (default__.safeDivisionInt(-819, (d_3252_v15_)[default__.safeIndex(608, (d_3252_v15_).length(0))]))
                index506_ = default__.safeIndex(608, (d_3252_v15_).length(0))
                (d_3252_v15_)[index506_] = default__.safeDivisionInt(((d_3252_v15_)[default__.safeIndex(608, (d_3252_v15_).length(0))]) - (d_3236_v0_), default__.fm25(d_3236_v0_, len(_dafny.Set({(self).f4})), self.f6, globalState))
        elif True:
            d_3294_v51_: _dafny.Seq
            d_3294_v51_ = _dafny.SeqWithoutIsStrInference([(d_3247_v10_).f12])
            d_3295_v52_: _dafny.Seq
            d_3295_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
            d_3296_v53_: D19
            d_3296_v53_ = D19_DC53((len(d_3294_v51_)) - (len(d_3295_v52_)), (0) - (d_3236_v0_), p3, self.f7)
            index507_ = default__.safeIndex(762, ((d_3247_v10_).f12).length(0))
            ((d_3247_v10_).f12)[index507_] = self.f7
            d_3297_v54_: _dafny.Seq
            d_3297_v54_ = _dafny.SeqWithoutIsStrInference([default__.fm49(d_3295_v52_, default__.fm25((self).f5, 189, self.f6, globalState), p3, True, globalState), d_3236_v0_, 929])
            index508_ = default__.safeIndex(762, ((d_3247_v10_).f12).length(0))
            rhs449_ = (self).f5
            rhs450_ = (d_3237_v1_)[default__.safeIndex((d_3297_v54_)[default__.safeIndex((self).f5, len(d_3297_v54_))], len(d_3237_v1_))]
            rhs451_ = d_3296_v53_
            rhs452_ = p2
            lhs252_ = self
            lhs253_ = (d_3247_v10_).f12
            lhs254_ = default__.safeIndex(762, ((d_3247_v10_).f12).length(0))
            d_3236_v0_ = rhs449_
            lhs252_.f7 = rhs450_
            d_3296_v53_ = rhs451_
            lhs253_[lhs254_] = rhs452_
            d_3298_v55_: _dafny.Set
            d_3298_v55_ = _dafny.Set({False})
            d_3299_v56_: _dafny.Seq
            d_3299_v56_ = _dafny.SeqWithoutIsStrInference([(d_3298_v55_) | (d_3298_v55_)])
            index509_ = default__.safeIndex(762, ((d_3247_v10_).f12).length(0))
            rhs453_ = (((d_3299_v56_).set(default__.safeIndex(len(d_3295_v52_), len(d_3299_v56_)), _dafny.Set({False}))) + (d_3299_v56_)) + (d_3299_v56_)
            rhs454_ = self.f7
            lhs255_ = (d_3247_v10_).f12
            lhs256_ = default__.safeIndex(762, ((d_3247_v10_).f12).length(0))
            d_3299_v56_ = rhs453_
            lhs255_[lhs256_] = rhs454_
            d_3236_v0_ = (self).f5
            d_3300_v57_: _dafny.Seq
            out41_: _dafny.Seq
            out41_ = (d_3247_v10_).m8(p2, (self).f4, (d_3297_v54_)[default__.safeIndex(d_3236_v0_, len(d_3297_v54_))], globalState)
            d_3300_v57_ = out41_
            d_3236_v0_ = (d_3297_v54_)[default__.safeIndex((self).f5, len(d_3297_v54_))]
        d_3301_v58_: _dafny.Set
        d_3301_v58_ = _dafny.Set({d_3236_v0_, (self).f5})
        d_3302_v59_: D11
        d_3302_v59_ = D11_DC27(d_3301_v58_)
        d_3303_v60_: _dafny.Seq
        d_3303_v60_ = _dafny.SeqWithoutIsStrInference([(self).f5])
        d_3304_v61_: _dafny.Map
        d_3304_v61_ = _dafny.Map({self.f6: (0) - (len(default__.fm52(p3, (d_3302_v59_).cf51, len(d_3303_v60_), self.f7, globalState)))})
        d_3305_v62_: D4
        d_3305_v62_ = D4_DC9(d_3304_v61_, (self).f5, p3)
        d_3306_v63_: _dafny.Map
        d_3306_v63_ = _dafny.Map({(0) - (((self).f5 if p2 else (d_3305_v62_).cf22)): (self).f5})
        d_3307_v65_: _dafny.MultiSet
        d_3307_v65_ = _dafny.MultiSet([(self).f5])
        d_3308_v66_: _dafny.Seq
        d_3308_v66_ = _dafny.SeqWithoutIsStrInference([d_3307_v65_, d_3307_v65_, d_3307_v65_])
        d_3309_v67_: _dafny.Seq
        def iife263_():
            coll113_ = _dafny.Map()
            compr_113_: int
            for compr_113_ in ((d_3308_v66_)[default__.safeIndex((self).f5, len(d_3308_v66_))]).Elements:
                d_3310_v64_: int = compr_113_
                if (d_3310_v64_) in ((d_3308_v66_)[default__.safeIndex((self).f5, len(d_3308_v66_))]):
                    coll113_[default__.safeDivisionInt(d_3310_v64_, (self).f5)] = d_3236_v0_
            return _dafny.Map(coll113_)
        d_3309_v67_ = _dafny.SeqWithoutIsStrInference([d_3306_v63_, _dafny.Map({-397: 955}), d_3306_v63_, iife263_()
        , d_3306_v63_])
        d_3306_v63_ = (d_3309_v67_)[default__.safeIndex(len(d_3237_v1_), len(d_3309_v67_))]
        d_3311_v68_: _dafny.Seq
        d_3311_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mydmsn"))
        d_3312_v69_: str
        d_3312_v69_ = _dafny.CodePoint('y')
        d_3313_v70_: C8
        nw513_ = C8()
        nw513_.ctor__(d_3311_v68_, (self).f5)
        d_3313_v70_ = nw513_
        d_3314_v71_: _dafny.MultiSet
        d_3314_v71_ = _dafny.MultiSet([d_3313_v70_, d_3313_v70_])
        d_3315_v72_: _dafny.Map
        d_3315_v72_ = _dafny.Map({self.f7: self.f6})
        d_3316_v73_: _dafny.Array
        nw514_ = _dafny.Array(None, 20)
        nw514_[int(0)] = d_3236_v0_
        nw514_[int(1)] = (self).f5
        nw514_[int(2)] = len(d_3304_v61_)
        nw514_[int(3)] = ((_dafny.MultiSet([len(d_3306_v63_)])).set((self).f5, default__.abs(d_3236_v0_))).cardinality
        nw514_[int(4)] = 672
        nw514_[int(5)] = len(d_3304_v61_)
        nw514_[int(6)] = d_3236_v0_
        nw514_[int(7)] = d_3236_v0_
        nw514_[int(8)] = d_3236_v0_
        nw514_[int(9)] = 255
        nw514_[int(10)] = (0) - (((d_3314_v71_).set(d_3313_v70_, default__.abs((self).f5))).cardinality)
        nw514_[int(11)] = -824
        nw514_[int(12)] = len(d_3311_v68_)
        nw514_[int(13)] = (self).f5
        nw514_[int(14)] = (0) - (default__.fm10((self).f4, len(d_3315_v72_), (d_3313_v70_).f14, globalState))
        nw514_[int(15)] = default__.fm10(p2, (d_3313_v70_).f14, (self).f5, globalState)
        nw514_[int(16)] = d_3236_v0_
        nw514_[int(17)] = d_3236_v0_
        nw514_[int(18)] = (self).f5
        nw514_[int(19)] = d_3236_v0_
        d_3316_v73_ = nw514_
        d_3317_v74_: _dafny.MultiSet
        d_3317_v74_ = _dafny.MultiSet([d_3316_v73_, d_3316_v73_])
        rhs455_ = (self).fm0(globalState)
        rhs456_ = d_3311_v68_
        rhs457_ = ((_dafny.MultiSet([d_3316_v73_])).intersection(d_3317_v74_)).cardinality
        rhs458_ = d_3312_v69_
        r1 = rhs455_
        d_3311_v68_ = rhs456_
        d_3236_v0_ = rhs457_
        d_3312_v69_ = rhs458_
        d_3318_v75_: _dafny.Set
        d_3318_v75_ = _dafny.Set({_dafny.Set({self.f6})})
        d_3319_v76_: D22
        d_3319_v76_ = D22_DC57(d_3318_v75_)
        r0 = (d_3318_v75_) | ((d_3319_v76_).cf101)
        r1 = True
        return r0, r1


class C14:
    def  __init__(self):
        self.f2: bool = False
        self._f3: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C14"
    def ctor__(self, f2, f3):
        (self).f2 = f2
        (self)._f3 = f3

    def m1(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r3: int = int(0)
        d_3320_v0_: str
        d_3320_v0_ = _dafny.CodePoint('e')
        d_3321_v1_: _dafny.Seq
        d_3321_v1_ = _dafny.SeqWithoutIsStrInference([d_3320_v0_, d_3320_v0_, d_3320_v0_, d_3320_v0_])
        hi9_ = len(d_3321_v1_)
        for d_3322_i0_ in range((self).f3, hi9_):
            r3 = (self).f3
            d_3323_v2_: _dafny.Array
            nw515_ = _dafny.Array(False, 8)
            d_3323_v2_ = nw515_
            index510_ = default__.safeIndex(911, (d_3323_v2_).length(0))
            (d_3323_v2_)[index510_] = p0
            index511_ = default__.safeIndex(911, (d_3323_v2_).length(0))
            (d_3323_v2_)[index511_] = True
            d_3324_v3_: _dafny.Set
            d_3324_v3_ = _dafny.Set({808})
            d_3325_v4_: _dafny.Set
            d_3325_v4_ = _dafny.Set({len(d_3324_v3_), (self).f3})
            r2 = len(_dafny.Set({d_3325_v4_}))
            d_3326_v5_: _dafny.Seq
            d_3326_v5_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")), _dafny.SeqWithoutIsStrInference([d_3320_v0_ for d_3327_i1_ in range(default__.abs(-515))]), d_3321_v1_, _dafny.SeqWithoutIsStrInference([d_3320_v0_ for d_3328_i2_ in range(default__.abs(-130))])])
            d_3329_v6_: _dafny.Map
            d_3329_v6_ = _dafny.Map({d_3322_i0_: len(d_3326_v5_)})
            index512_ = default__.safeIndex(911, (d_3323_v2_).length(0))
            (d_3323_v2_)[index512_] = not(((d_3322_i0_) + ((self).f3)) > (len(d_3329_v6_)))
        hi10_ = (self).f3
        for d_3330_i3_ in range(default__.safeDivisionInt(512, (self).f3), hi10_):
            d_3331_v7_: _dafny.MultiSet
            d_3331_v7_ = _dafny.MultiSet([default__.fm18(p0, self.f2, self.f2, globalState)])
            d_3332_v8_: _dafny.MultiSet
            d_3332_v8_ = _dafny.MultiSet([d_3331_v7_, d_3331_v7_])
            d_3333_v9_: C4
            nw516_ = C4()
            nw516_.ctor__(d_3330_i3_, (d_3331_v7_) in (d_3332_v8_), d_3321_v1_, p0, p0, d_3330_i3_, len(d_3321_v1_), p0)
            d_3333_v9_ = nw516_
            r2 = d_3330_i3_
            d_3334_v10_: _dafny.MultiSet
            d_3334_v10_ = _dafny.MultiSet([self.f2])
            r1 = ((default__.fm4(globalState)).isdisjoint(d_3334_v10_)) == (p0)
            d_3335_v11_: D0
            d_3335_v11_ = D0_DC0((d_3333_v9_).f24, (self).f3, not((d_3333_v9_).f24), d_3330_i3_, p0)
            d_3336_v12_: _dafny.Seq
            d_3336_v12_ = _dafny.SeqWithoutIsStrInference([default__.fm25(d_3330_i3_, len(_dafny.SeqWithoutIsStrInference([d_3330_i3_, (self).f3, (d_3333_v9_).f23, d_3330_i3_])), self.f2, globalState)])
            r3 = (default__.fm32(d_3335_v11_, globalState)) - ((d_3336_v12_)[default__.safeIndex(len((d_3321_v1_).set(default__.safeIndex((self).f3, len(d_3321_v1_)), d_3320_v0_)), len(d_3336_v12_))])
        hi11_ = (self).f3
        for d_3337_i4_ in range(191, hi11_):
            r0 = self.f2
            d_3338_v13_: _dafny.MultiSet
            d_3338_v13_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ckr"))), 767])
            d_3339_v15_: _dafny.Seq
            def iife264_():
                coll114_ = _dafny.Set()
                compr_114_: str
                for compr_114_ in (d_3321_v1_).Elements:
                    d_3340_v14_: str = compr_114_
                    if (d_3340_v14_) in (d_3321_v1_):
                        coll114_ = coll114_.union(_dafny.Set([d_3340_v14_]))
                return _dafny.Set(coll114_)
            d_3339_v15_ = _dafny.SeqWithoutIsStrInference([len(iife264_()
            )])
            d_3341_v16_: _dafny.Array
            nw517_ = _dafny.Array(False, 13)
            d_3341_v16_ = nw517_
            d_3342_v17_: C2
            nw518_ = C2()
            nw518_.ctor__(d_3341_v16_, d_3338_v13_, (d_3321_v1_).set(default__.safeIndex((self).f3, len(d_3321_v1_)), d_3320_v0_), self.f2, p0, (self).f3)
            d_3342_v17_ = nw518_
            d_3343_v18_: _dafny.Map
            d_3343_v18_ = _dafny.Map({d_3342_v17_: d_3337_i4_})
            d_3344_v19_: _dafny.Map
            d_3344_v19_ = _dafny.Map({len(d_3339_v15_): d_3343_v18_})
            d_3345_v20_: _dafny.MultiSet
            d_3345_v20_ = _dafny.MultiSet([_dafny.Map({d_3342_v17_: d_3337_i4_}), d_3343_v18_])
            d_3346_v21_: D0
            d_3346_v21_ = D0_DC0(self.f2, ((0) - ((d_3338_v13_).cardinality)) - ((self).f3), (((d_3344_v19_)[d_3337_i4_] if (d_3337_i4_) in (d_3344_v19_) else d_3343_v18_)) in (d_3345_v20_), d_3337_i4_, not ((d_3342_v17_).fm0(globalState)) or (self.f2))
            source40_ = d_3346_v21_
            if source40_.is_DC0:
                d_3347___mcc_h0_ = source40_.cf0
                d_3348___mcc_h1_ = source40_.cf1
                d_3349___mcc_h2_ = source40_.cf2
                d_3350___mcc_h3_ = source40_.cf3
                d_3351___mcc_h4_ = source40_.cf4
                d_3352_cf4_ = d_3351___mcc_h4_
                d_3353_cf3_ = d_3350___mcc_h3_
                d_3354_cf2_ = d_3349___mcc_h2_
                d_3355_cf1_ = d_3348___mcc_h1_
                d_3356_cf0_ = d_3347___mcc_h0_
                d_3357_v22_: _dafny.Set
                d_3357_v22_ = _dafny.Set({d_3339_v15_})
                d_3358_v23_: C9
                nw519_ = C9()
                nw519_.ctor__((d_3342_v17_).f28)
                d_3358_v23_ = nw519_
                d_3359_v24_: _dafny.Seq
                d_3359_v24_ = _dafny.SeqWithoutIsStrInference([p0, self.f2])
                d_3360_v25_: D11
                d_3360_v25_ = D11_DC29(len(d_3321_v1_), (d_3342_v17_).f28, d_3356_cf0_, (d_3359_v24_)[default__.safeIndex(len(d_3359_v24_), len(d_3359_v24_))])
                d_3361_v26_: _dafny.Map
                d_3361_v26_ = _dafny.Map({len(d_3321_v1_): d_3338_v13_})
                d_3362_v27_: _dafny.Seq
                d_3362_v27_ = _dafny.SeqWithoutIsStrInference([d_3341_v16_])
                d_3363_v28_: D13
                d_3363_v28_ = D13_DC34(d_3362_v27_)
                d_3364_v29_: _dafny.Map
                d_3364_v29_ = _dafny.Map({d_3363_v28_: _dafny.MultiSet([len(d_3321_v1_)])})
                d_3365_v30_: _dafny.Array
                nw520_ = _dafny.Array(None, 24)
                nw520_[int(0)] = d_3338_v13_
                nw520_[int(1)] = default__.fm7(True, d_3357_v22_, len((_dafny.Map({True: d_3358_v23_})).set(d_3352_cf4_, d_3358_v23_)), (0) - (d_3337_i4_), globalState)
                nw520_[int(2)] = (d_3342_v17_).f29
                nw520_[int(3)] = (d_3342_v17_).f29
                nw520_[int(4)] = (d_3342_v17_).f29
                nw520_[int(5)] = (d_3342_v17_).f29
                nw520_[int(6)] = (d_3342_v17_).f29
                nw520_[int(7)] = ((d_3342_v17_).f29) | ((d_3342_v17_).f29)
                nw520_[int(8)] = (d_3342_v17_).f29
                nw520_[int(9)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_3337_i4_ for d_3366_i5_ in range(default__.abs(119))]))
                nw520_[int(10)] = default__.fm7(d_3352_cf4_, d_3357_v22_, (0) - (d_3355_cf1_), d_3353_cf3_, globalState)
                nw520_[int(11)] = ((d_3342_v17_).f29) | (_dafny.MultiSet(d_3339_v15_))
                nw520_[int(12)] = (_dafny.MultiSet([d_3353_cf3_, d_3337_i4_, d_3353_cf3_])).set(len(_dafny.Set({d_3352_cf4_, (d_3360_v25_).cf59})), default__.abs((self).f3))
                nw520_[int(13)] = (((d_3361_v26_)[d_3355_cf1_] if (d_3355_cf1_) in (d_3361_v26_) else _dafny.MultiSet(d_3339_v15_))) | (((d_3342_v17_).f29).set(327, default__.abs(127)))
                nw520_[int(14)] = (d_3342_v17_).f29
                nw520_[int(15)] = (d_3338_v13_) | ((d_3342_v17_).f29)
                nw520_[int(16)] = ((d_3342_v17_).f29) | (default__.fm7(d_3356_cf0_, d_3357_v22_, (self).f3, d_3355_cf1_, globalState))
                nw520_[int(17)] = ((d_3364_v29_)[d_3363_v28_] if (d_3363_v28_) in (d_3364_v29_) else ((d_3342_v17_).f29).set(103, default__.abs(default__.fm49(d_3321_v1_, d_3337_i4_, p0, self.f2, globalState))))
                nw520_[int(18)] = (d_3342_v17_).f29
                nw520_[int(19)] = d_3338_v13_
                nw520_[int(20)] = (d_3342_v17_).f29
                nw520_[int(21)] = (d_3342_v17_).f29
                nw520_[int(22)] = (d_3342_v17_).f29
                nw520_[int(23)] = (d_3342_v17_).f29
                d_3365_v30_ = nw520_
                index513_ = default__.safeIndex(41, (d_3365_v30_).length(0))
                (d_3365_v30_)[index513_] = _dafny.MultiSet([d_3355_cf1_])
                d_3367_v31_: _dafny.Seq
                d_3367_v31_ = _dafny.SeqWithoutIsStrInference([d_3338_v13_])
                index514_ = default__.safeIndex(41, (d_3365_v30_).length(0))
                rhs459_ = (d_3367_v31_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([-480, 246, (0) - (d_3355_cf1_), d_3355_cf1_, (0) - (d_3337_i4_)])), len(d_3367_v31_))]
                rhs460_ = (d_3321_v1_) < (d_3321_v1_)
                lhs257_ = d_3365_v30_
                lhs258_ = default__.safeIndex(41, (d_3365_v30_).length(0))
                lhs257_[lhs258_] = rhs459_
                d_3354_cf2_ = rhs460_
                d_3368_v32_: _dafny.Set
                d_3368_v32_ = _dafny.Set({d_3356_cf0_})
                rhs461_ = (d_3355_cf1_) * (len((d_3357_v22_ if self.f2 else d_3357_v22_)))
                rhs462_ = (d_3354_cf2_) not in (d_3368_v32_)
                d_3353_cf3_ = rhs461_
                r1 = rhs462_
                d_3354_cf2_ = p0
                d_3369_v33_: _dafny.Map
                d_3369_v33_ = _dafny.Map({d_3355_cf1_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")))})
                d_3370_v34_: _dafny.Map
                d_3370_v34_ = _dafny.Map({self.f2: d_3369_v33_})
                d_3370_v34_ = (d_3370_v34_).set(d_3356_cf0_, d_3369_v33_)
            elif source40_.is_DC1:
                d_3371___mcc_h5_ = source40_.cf5
                d_3372___mcc_h6_ = source40_.cf6
                d_3373___mcc_h7_ = source40_.cf7
                d_3374___mcc_h8_ = source40_.cf8
                d_3375_cf8_ = d_3374___mcc_h8_
                d_3376_cf7_ = d_3373___mcc_h7_
                d_3377_cf6_ = d_3372___mcc_h6_
                d_3378_cf5_ = d_3371___mcc_h5_
                r1 = (298) in ((d_3342_v17_).f29)
                d_3379_v35_: _dafny.Seq
                d_3379_v35_ = _dafny.SeqWithoutIsStrInference([d_3341_v16_, (d_3342_v17_).f28])
                d_3380_v36_: _dafny.Map
                d_3380_v36_ = _dafny.Map({d_3378_cf5_: (d_3375_cf8_) * (len(d_3379_v35_))})
                d_3380_v36_ = (d_3380_v36_).set(len((_dafny.Set({p0})) - (_dafny.Set({self.f2, d_3377_cf6_, d_3377_cf6_, d_3377_cf6_, False}))), (self).f3)
                d_3381_v37_: D19
                d_3381_v37_ = D19_DC53(d_3375_cf8_, (d_3375_cf8_) - (d_3376_cf7_), p0, p0)
                d_3381_v37_ = d_3381_v37_
                d_3382_v38_: C11
                nw521_ = C11()
                nw521_.ctor__(D1_DC5(p0, len(d_3321_v1_), d_3321_v1_), d_3377_cf6_, (self).f3)
                d_3382_v38_ = nw521_
            elif True:
                d_3383___mcc_h9_ = source40_.cf9
                d_3384___mcc_h10_ = source40_.cf10
                d_3385_cf10_ = d_3384___mcc_h10_
                d_3386_cf9_ = d_3383___mcc_h9_
                d_3387_v39_: _dafny.Array
                nw522_ = _dafny.Array(D11.default()(), 12)
                d_3387_v39_ = nw522_
                d_3388_v40_: _dafny.Set
                d_3388_v40_ = _dafny.Set({d_3337_i4_})
                d_3389_v41_: D11
                d_3389_v41_ = D11_DC27(d_3388_v40_)
                index515_ = default__.safeIndex(942, (d_3387_v39_).length(0))
                (d_3387_v39_)[index515_] = d_3389_v41_
                index516_ = default__.safeIndex(942, (d_3387_v39_).length(0))
                (d_3387_v39_)[index516_] = d_3389_v41_
                d_3386_cf9_ = (d_3321_v1_) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_3390_i6_ in range(default__.abs(701))]))
                d_3391_v42_: _dafny.Array
                def lambda136_(d_3392_cf9_, d_3393_p0_):
                    def lambda137_(d_3394_i7_):
                        return default__.safeDivisionInt(d_3394_i7_, len(_dafny.SeqWithoutIsStrInference([not(d_3392_cf9_), True, d_3393_p0_, True, d_3393_p0_])))

                    return lambda137_

                init80_ = lambda136_(d_3386_cf9_, p0)
                nw523_ = _dafny.Array(None, 4)
                for i0_80_ in range(nw523_.length(0)):
                    nw523_[i0_80_] = init80_(i0_80_)
                d_3391_v42_ = nw523_
                d_3395_v43_: _dafny.Map
                d_3395_v43_ = _dafny.Map({d_3385_cf10_: d_3386_cf9_})
                index517_ = default__.safeIndex(501, (d_3391_v42_).length(0))
                (d_3391_v42_)[index517_] = default__.safeDivisionInt(len(d_3395_v43_), (0) - (d_3337_i4_))
                index518_ = default__.safeIndex(501, (d_3391_v42_).length(0))
                (d_3391_v42_)[index518_] = 923
                index519_ = default__.safeIndex(501, (d_3391_v42_).length(0))
                (d_3391_v42_)[index519_] = (d_3391_v42_)[default__.safeIndex(501, (d_3391_v42_).length(0))]
            d_3396_v44_: _dafny.Array
            nw524_ = _dafny.Array(_dafny.Seq({}), 28)
            d_3396_v44_ = nw524_
            d_3397_v45_: _dafny.Seq
            d_3397_v45_ = _dafny.SeqWithoutIsStrInference([p0, self.f2])
            index520_ = default__.safeIndex(794, (d_3396_v44_).length(0))
            (d_3396_v44_)[index520_] = ((default__.fm34(d_3337_i4_, (self).f3, (d_3397_v45_)[default__.safeIndex(d_3337_i4_, len(d_3397_v45_))], (self).f3, globalState)).set(default__.safeIndex(d_3337_i4_, len(default__.fm34(d_3337_i4_, (self).f3, (d_3397_v45_)[default__.safeIndex(d_3337_i4_, len(d_3397_v45_))], (self).f3, globalState))), self.f2)) + (d_3397_v45_)
            index521_ = default__.safeIndex(794, (d_3396_v44_).length(0))
            (d_3396_v44_)[index521_] = d_3397_v45_
            r3 = (self).f3
        hi12_ = (self).f3
        for d_3398_i8_ in range((self).f3, hi12_):
            d_3399_v46_: _dafny.Map
            d_3399_v46_ = _dafny.Map({d_3321_v1_: True})
            r3 = len(d_3399_v46_)
            if p0:
                d_3400_v47_: C5
                nw525_ = C5()
                nw525_.ctor__()
                d_3400_v47_ = nw525_
                d_3400_v47_ = d_3400_v47_
                r1 = default__.fm27(_dafny.CodePoint('c'), globalState)
                d_3401_v48_: _dafny.Set
                d_3401_v48_ = _dafny.Set({d_3398_i8_, d_3398_i8_})
                d_3402_v49_: C3
                nw526_ = C3()
                nw526_.ctor__(d_3320_v0_, len(d_3321_v1_), (d_3401_v48_).isdisjoint(d_3401_v48_), self.f2, ((0) - (d_3398_i8_)) + ((self).f3))
                d_3402_v49_ = nw526_
                r1 = True
                d_3403_v50_: D17
                d_3403_v50_ = D17_DC45()
                d_3403_v50_ = d_3403_v50_
            elif True:
                d_3404_v51_: D9
                d_3404_v51_ = D9_DC22(p0)
                d_3405_v52_: _dafny.Map
                d_3405_v52_ = _dafny.Map({d_3398_i8_: D14_DC37(d_3404_v51_)})
                d_3406_v53_: D14
                d_3406_v53_ = D14_DC37(d_3404_v51_)
                d_3407_v54_: _dafny.Seq
                def iife265_(_pat_let75_0):
                    def iife266_(d_3408_dt__update__tmp_h0_):
                        def iife267_(_pat_let76_0):
                            def iife268_(d_3409_dt__update_hcf73_h0_):
                                return D14_DC37(d_3409_dt__update_hcf73_h0_)
                            return iife268_(_pat_let76_0)
                        return iife267_(D9_DC22(False))
                    return iife266_(_pat_let75_0)
                d_3407_v54_ = _dafny.SeqWithoutIsStrInference([(d_3405_v52_) | (_dafny.Map({(0) - ((self).f3): iife265_(D14_DC37(d_3404_v51_))})), _dafny.Map({(self).f3: d_3406_v53_})])
                d_3407_v54_ = ((d_3407_v54_) + (d_3407_v54_)) + (_dafny.SeqWithoutIsStrInference([d_3405_v52_]))
                d_3321_v1_ = ((d_3321_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdios")))).set(default__.safeIndex(192, len((d_3321_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdios"))))), d_3320_v0_)
                d_3410_v55_: _dafny.Array
                nw527_ = _dafny.Array(False, 20)
                d_3410_v55_ = nw527_
                d_3410_v55_ = d_3410_v55_
                d_3411_v56_: _dafny.Seq
                d_3411_v56_ = _dafny.SeqWithoutIsStrInference([d_3398_i8_, 113])
                r0 = (d_3411_v56_) == (d_3411_v56_)
                rhs463_ = self.f2
                rhs464_ = default__.safeDivisionInt((0) - (default__.fm10(self.f2, (self).f3, (self).f3, globalState)), d_3398_i8_)
                r1 = rhs463_
                r2 = rhs464_
            r3 = ((0) - (d_3398_i8_)) - (d_3398_i8_)
            d_3412_v57_: _dafny.Map
            d_3412_v57_ = _dafny.Map({d_3398_i8_: self.f2})
            d_3413_v58_: _dafny.Array
            nw528_ = _dafny.Array(None, 12)
            nw528_[int(0)] = self.f2
            nw528_[int(1)] = p0
            nw528_[int(2)] = (self.f2) == (p0)
            nw528_[int(3)] = self.f2
            nw528_[int(4)] = (d_3321_v1_) < (d_3321_v1_)
            nw528_[int(5)] = self.f2
            nw528_[int(6)] = not((d_3398_i8_) > ((0) - ((self).f3)))
            nw528_[int(7)] = self.f2
            nw528_[int(8)] = self.f2
            nw528_[int(9)] = not(not (self.f2) or (default__.fm27(d_3320_v0_, globalState)))
            nw528_[int(10)] = ((d_3412_v57_)[(self).f3] if ((self).f3) in (d_3412_v57_) else p0)
            nw528_[int(11)] = p0
            d_3413_v58_ = nw528_
            d_3413_v58_ = d_3413_v58_
        d_3414_v59_: D18
        d_3414_v59_ = D18_DC50(p0, p0)
        source41_ = (d_3414_v59_ if self.f2 else (d_3414_v59_ if self.f2 else d_3414_v59_))
        if source41_.is_DC50:
            d_3415___mcc_h11_ = source41_.cf90
            d_3416___mcc_h12_ = source41_.cf91
            d_3417_cf91_ = d_3416___mcc_h12_
            d_3418_cf90_ = d_3415___mcc_h11_
            d_3419_v60_: _dafny.Seq
            d_3419_v60_ = _dafny.SeqWithoutIsStrInference([self.f2])
            d_3420_v61_: _dafny.Array
            def lambda138_(d_3421_i9_):
                return (d_3421_i9_) + ((self).f3)

            init81_ = lambda138_
            nw529_ = _dafny.Array(None, 15)
            for i0_81_ in range(nw529_.length(0)):
                nw529_[i0_81_] = init81_(i0_81_)
            d_3420_v61_ = nw529_
            d_3422_v62_: _dafny.Map
            d_3422_v62_ = _dafny.Map({d_3321_v1_: self.f2})
            d_3423_v63_: _dafny.Map
            d_3423_v63_ = _dafny.Map({d_3420_v61_: len(((d_3422_v62_).set(d_3321_v1_, not(d_3418_cf90_))).set(d_3321_v1_, d_3417_cf91_))})
            d_3424_v64_: D5
            d_3424_v64_ = D5_DC14(d_3417_cf91_, d_3419_v60_, d_3420_v61_, default__.fm3(globalState), default__.fm49(d_3321_v1_, len(d_3423_v63_), False, self.f2, globalState))
            if (d_3424_v64_).cf31:
                r2 = (self).f3
                d_3425_v65_: _dafny.Map
                d_3425_v65_ = _dafny.Map({self.f2: d_3420_v61_})
                d_3425_v65_ = (d_3425_v65_).set(self.f2, d_3420_v61_)
                d_3426_v66_: _dafny.MultiSet
                d_3426_v66_ = _dafny.MultiSet([d_3417_cf91_])
                r0 = (d_3426_v66_) == (default__.fm47(d_3418_cf90_, globalState))
                d_3427_v67_: _dafny.Set
                d_3427_v67_ = _dafny.Set({True})
                (globalState).f0 = d_3427_v67_
                d_3428_v68_: C1
                nw530_ = C1()
                nw530_.ctor__((self).f3, ((d_3422_v62_)[d_3321_v1_] if (d_3321_v1_) in (d_3422_v62_) else d_3418_cf90_), self.f2, (self).f3)
                d_3428_v68_ = nw530_
                d_3429_v69_: _dafny.Map
                d_3429_v69_ = _dafny.Map({self.f2: d_3428_v68_})
                d_3430_v70_: _dafny.Set
                d_3430_v70_ = _dafny.Set({(self).f3, (self).f3, (len(d_3429_v69_) if True else (self).f3), (0) - ((self).f3), (self).f3})
                rhs465_ = (d_3430_v70_).intersection(d_3430_v70_)
                rhs466_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xhxwlnltd"))) == (d_3321_v1_)
                d_3430_v70_ = rhs465_
                d_3417_cf91_ = rhs466_
            elif True:
                d_3431_v71_: _dafny.Map
                d_3431_v71_ = _dafny.Map({-508: (self).f3})
                d_3432_v72_: T0
                nw531_ = C7()
                nw531_.ctor__((self).f3, p0, d_3418_cf90_, (0) - ((self).f3), (self).f3, d_3417_cf91_, d_3321_v1_, p0)
                d_3432_v72_ = nw531_
                d_3433_v73_: _dafny.Map
                d_3433_v73_ = _dafny.Map({d_3432_v72_: d_3321_v1_})
                d_3434_v74_: _dafny.Map
                d_3434_v74_ = _dafny.Map({(d_3432_v72_).f4: d_3432_v72_})
                d_3435_v75_: _dafny.Seq
                d_3435_v75_ = _dafny.SeqWithoutIsStrInference([d_3432_v72_, ((d_3434_v74_)[self.f2] if (self.f2) in (d_3434_v74_) else d_3432_v72_)])
                d_3321_v1_ = (((_dafny.SeqWithoutIsStrInference([d_3320_v0_ for d_3436_i10_ in range(default__.abs(877))])).set(default__.safeIndex(len(d_3431_v71_), len(_dafny.SeqWithoutIsStrInference([d_3320_v0_ for d_3437_i10_ in range(default__.abs(877))]))), d_3320_v0_) if d_3417_cf91_ else d_3321_v1_) if (default__.fm10(self.f2, (self).f3, (self).f3, globalState)) in (_dafny.MultiSet([-585, (self).f3, (self).f3, (self).f3])) else ((d_3433_v73_)[(d_3435_v75_)[default__.safeIndex((0) - ((self).f3), len(d_3435_v75_))]] if ((d_3435_v75_)[default__.safeIndex((0) - ((self).f3), len(d_3435_v75_))]) in (d_3433_v73_) else d_3321_v1_))
                d_3419_v60_ = d_3419_v60_
                d_3438_v76_: _dafny.Seq
                d_3438_v76_ = _dafny.SeqWithoutIsStrInference([default__.fm85(globalState)])
                d_3439_v77_: D22
                d_3439_v77_ = D22_DC58(d_3321_v1_, (self).f3, p0)
                d_3440_v78_: _dafny.Array
                def lambda139_(d_3441_cf91_):
                    def lambda140_(d_3442_i11_):
                        return d_3441_cf91_

                    return lambda140_

                init82_ = lambda139_(d_3417_cf91_)
                nw532_ = _dafny.Array(None, 25)
                for i0_82_ in range(nw532_.length(0)):
                    nw532_[i0_82_] = init82_(i0_82_)
                d_3440_v78_ = nw532_
                d_3443_v79_: _dafny.Map
                d_3443_v79_ = _dafny.Map({False: (self).f3})
                d_3444_v80_: _dafny.Map
                d_3444_v80_ = d_3443_v79_
                d_3445_v81_: _dafny.Set
                d_3445_v81_ = _dafny.Set({d_3444_v80_})
                rhs467_ = _dafny.SeqWithoutIsStrInference([default__.fm85(globalState), d_3444_v80_, _dafny.Map({(d_3432_v72_).f4: 298})])
                rhs468_ = d_3439_v77_
                rhs469_ = (not(p0)) == (not((d_3445_v81_).ispropersubset(d_3445_v81_)))
                rhs470_ = d_3440_v78_
                lhs259_ = self
                d_3438_v76_ = rhs467_
                d_3439_v77_ = rhs468_
                lhs259_.f2 = rhs469_
                d_3440_v78_ = rhs470_
                d_3446_v82_: C10
                nw533_ = C10()
                nw533_.ctor__(len(d_3443_v79_), d_3440_v78_)
                d_3446_v82_ = nw533_
                nw534_ = C10()
                nw534_.ctor__((d_3432_v72_).f5, (d_3446_v82_).f11)
                d_3446_v82_ = nw534_
                d_3443_v79_ = (d_3443_v79_).set(not((d_3432_v72_).f4), 96)
            d_3447_v83_: _dafny.Seq
            d_3447_v83_ = _dafny.SeqWithoutIsStrInference([(self).f3])
            d_3448_v84_: D1
            d_3448_v84_ = D1_DC5(p0, len(d_3447_v83_), d_3321_v1_)
            d_3449_v85_: C11
            nw535_ = C11()
            nw535_.ctor__(d_3448_v84_, d_3418_cf90_, (0) - ((self).f3))
            d_3449_v85_ = nw535_
            d_3449_v85_ = d_3449_v85_
            d_3450_v86_: _dafny.MultiSet
            d_3450_v86_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_3320_v0_ for d_3451_i12_ in range(default__.abs(866))])])
            if not((((d_3450_v86_).intersection(d_3450_v86_)).cardinality) >= ((self).f3)):
                d_3452_v87_: _dafny.Set
                d_3452_v87_ = _dafny.Set({(0) - ((self).f3)})
                d_3453_v88_: _dafny.Map
                d_3453_v88_ = _dafny.Map({d_3320_v0_: d_3452_v87_})
                def iife269_():
                    coll115_ = _dafny.Set()
                    compr_115_: int
                    for compr_115_ in _dafny.IntegerRange(589, 249):
                        d_3454_v89_: int = compr_115_
                        if ((589) <= (d_3454_v89_)) and ((d_3454_v89_) < (249)):
                            coll115_ = coll115_.union(_dafny.Set([(d_3454_v89_) * ((0) - ((self).f3))]))
                    return _dafny.Set(coll115_)
                d_3453_v88_ = (d_3453_v88_).set((d_3320_v0_ if d_3418_cf90_ else d_3320_v0_), iife269_()
                )
                r2 = (self).f3
                d_3455_v90_: _dafny.Seq
                d_3455_v90_ = _dafny.SeqWithoutIsStrInference([d_3419_v60_])
                d_3456_v91_: _dafny.Seq
                d_3456_v91_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([default__.fm27(d_3320_v0_, globalState)]), default__.fm8(not(d_3417_cf91_), (self).f3, (self).f3, globalState), d_3419_v60_, (d_3455_v90_)[default__.safeIndex((self).f3, len(d_3455_v90_))], d_3419_v60_])
                rhs471_ = (d_3449_v85_).fm0(globalState)
                rhs472_ = ((d_3455_v90_)[default__.safeIndex((self).f3, len(d_3455_v90_))]) + (d_3419_v60_)
                d_3417_cf91_ = rhs471_
                d_3419_v60_ = rhs472_
                d_3457_v92_: _dafny.Map
                d_3457_v92_ = _dafny.Map({self.f2: d_3447_v83_})
                (d_3449_v85_).m3(len((d_3457_v92_) | (d_3457_v92_)), globalState)
                r1 = self.f2
            elif True:
                d_3320_v0_ = d_3320_v0_
                d_3458_v93_: C7
                nw536_ = C7()
                nw536_.ctor__(default__.fm3(globalState), d_3417_cf91_, d_3417_cf91_, (0) - (((0) - ((self).f3)) * ((self).f3)), (self).f3, self.f2, d_3321_v1_, not(self.f2))
                d_3458_v93_ = nw536_
                d_3459_v94_: D16
                d_3459_v94_ = D16_DC41(d_3320_v0_, self.f2)
                d_3320_v0_ = (d_3459_v94_).cf77
                d_3460_v95_: _dafny.Seq
                d_3460_v95_ = _dafny.SeqWithoutIsStrInference([d_3419_v60_, d_3419_v60_])
                d_3461_v96_: _dafny.MultiSet
                d_3461_v96_ = _dafny.MultiSet([d_3447_v83_, d_3447_v83_, _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f3])).cardinality for d_3462_i13_ in range(default__.abs(267))])])
                d_3463_v97_: C13
                nw537_ = C13()
                nw537_.ctor__(not((d_3460_v95_) != (d_3460_v95_)), d_3418_cf90_, not((d_3461_v96_).isdisjoint(default__.fm33(default__.fm49(d_3321_v1_, (0) - ((self).f3), d_3417_cf91_, self.f2, globalState), d_3417_cf91_, (d_3458_v93_).f19, d_3417_cf91_, globalState))), (d_3458_v93_).f19)
                d_3463_v97_ = nw537_
                d_3464_v98_: _dafny.Map
                d_3464_v98_ = _dafny.Map({((d_3458_v93_).f19) * ((d_3458_v93_).f19): ((0) - ((d_3458_v93_).f19)) == ((self).f3)})
                def iife270_():
                    coll116_ = _dafny.Map()
                    compr_116_: int
                    for compr_116_ in _dafny.IntegerRange(922, -628):
                        d_3465_v99_: int = compr_116_
                        if ((922) <= (d_3465_v99_)) and ((d_3465_v99_) < (-628)):
                            coll116_[(d_3465_v99_) * (len(_dafny.Map({603: (d_3458_v93_).f19})))] = False
                    return _dafny.Map(coll116_)
                d_3464_v98_ = (default__.fm60(not(False), d_3463_v97_.f7, d_3463_v97_.f6, (self).f3, globalState)) | ((iife270_()
                ) | (d_3464_v98_))
            d_3466_v100_: _dafny.Map
            d_3466_v100_ = _dafny.Map({(self).f3: (0) - ((self).f3)})
            d_3467_v101_: T1
            nw538_ = C3()
            nw538_.ctor__(d_3320_v0_, (self).f3, p0, p0, ((d_3466_v100_)[len(d_3447_v83_)] if (len(d_3447_v83_)) in (d_3466_v100_) else (0) - ((self).f3)))
            d_3467_v101_ = nw538_
            d_3468_v102_: T2
            nw539_ = C7()
            nw539_.ctor__((self).f3, p0, (d_3467_v101_).f4, (self).f3, (d_3467_v101_).f15, p0, d_3321_v1_, d_3417_cf91_)
            d_3468_v102_ = nw539_
            d_3469_v103_: _dafny.Set
            d_3469_v103_ = _dafny.Set({(d_3467_v101_).f5, (d_3467_v101_).f15})
            d_3470_v104_: _dafny.Map
            d_3470_v104_ = _dafny.Map({d_3467_v101_: not(not(((_dafny.MultiSet([(_dafny.MultiSet([d_3468_v102_, d_3468_v102_])).cardinality])).cardinality) not in (d_3469_v103_)))})
            d_3470_v104_ = (d_3470_v104_).set(d_3467_v101_, (d_3467_v101_).f4)
        elif source41_.is_DC49:
            d_3471___mcc_h13_ = source41_.cf89
            d_3472_cf89_ = d_3471___mcc_h13_
            r1 = False
            r0 = False
            d_3473_v105_: _dafny.Array
            nw540_ = _dafny.Array(False, 5)
            d_3473_v105_ = nw540_
            d_3474_v106_: _dafny.Map
            d_3474_v106_ = _dafny.Map({(0) - ((d_3472_cf89_).f5): d_3473_v105_})
            d_3475_v107_: _dafny.Seq
            d_3475_v107_ = _dafny.SeqWithoutIsStrInference([self.f2, p0])
            d_3474_v106_ = (d_3474_v106_).set(default__.safeDivisionInt(len(d_3475_v107_), (d_3472_cf89_).f5), d_3473_v105_)
            d_3476_v108_: D1
            d_3476_v108_ = D1_DC5(p0, -969, d_3321_v1_)
            d_3477_v109_: C11
            nw541_ = C11()
            def iife271_(_pat_let77_0):
                def iife272_(d_3478_dt__update__tmp_h2_):
                    def iife273_(_pat_let78_0):
                        def iife274_(d_3479_dt__update_hcf16_h0_):
                            return D1_DC5((d_3478_dt__update__tmp_h2_).cf15, d_3479_dt__update_hcf16_h0_, (d_3478_dt__update__tmp_h2_).cf17)
                        return iife274_(_pat_let78_0)
                    return iife273_((self).f3)
                return iife272_(_pat_let77_0)
            nw541_.ctor__(iife271_(d_3476_v108_), p0, (self).f3)
            d_3477_v109_ = nw541_
        elif True:
            d_3480___mcc_h14_ = source41_.cf92
            d_3481_cf92_ = d_3480___mcc_h14_
            d_3482_v110_: D0
            d_3482_v110_ = D0_DC0(self.f2, (self).f3, self.f2, (self).f3, self.f2)
            if (d_3482_v110_).cf4:
                d_3483_v111_: _dafny.Set
                d_3483_v111_ = _dafny.Set({(self).f3, len(d_3321_v1_)})
                d_3484_v112_: _dafny.Seq
                d_3484_v112_ = _dafny.SeqWithoutIsStrInference([(self).f3, (self).f3])
                d_3485_v113_: C6
                nw542_ = C6()
                nw542_.ctor__((0) - ((self).f3), True, False, (self).f3, (0) - (len(d_3484_v112_)), p0)
                d_3485_v113_ = nw542_
                d_3486_v114_: _dafny.Map
                d_3486_v114_ = _dafny.Map({d_3485_v113_: d_3485_v113_.f21})
                d_3487_v115_: _dafny.Seq
                d_3487_v115_ = _dafny.SeqWithoutIsStrInference([len(d_3486_v114_), (self).f3, d_3485_v113_.f21, d_3485_v113_.f21, d_3485_v113_.f21])
                d_3488_v116_: C1
                nw543_ = C1()
                nw543_.ctor__(default__.safeDivisionInt((self).f3, len(d_3321_v1_)), not((d_3483_v111_).issubset(d_3483_v111_)), not (p0) or ((D12_DC33(self.f2, self.f2, d_3320_v0_, (self).f3)).cf67), default__.safeModuloInt(default__.fm25(len(d_3484_v112_), (self).f3, False, globalState), d_3485_v113_.f21))
                d_3488_v116_ = nw543_
                nw544_ = C1()
                nw544_.ctor__((self).f3, self.f2, (p0) or (not((d_3485_v113_).f22)), (0) - ((570) * (d_3485_v113_.f21)))
                d_3488_v116_ = nw544_
                d_3489_v117_: _dafny.Map
                d_3489_v117_ = _dafny.Map({p0: (self).f3})
                d_3490_v118_: _dafny.Map
                d_3490_v118_ = _dafny.Map({len(d_3489_v117_): (d_3485_v113_).f22})
                d_3491_v119_: _dafny.Map
                d_3491_v119_ = _dafny.Map({(d_3490_v118_) | (d_3490_v118_): (_dafny.SeqWithoutIsStrInference([d_3320_v0_ for d_3492_i14_ in range(default__.abs(-35))])) + (d_3321_v1_)})
                d_3491_v119_ = (d_3491_v119_).set(d_3490_v118_, d_3321_v1_)
                d_3493_v120_: _dafny.Array
                def lambda141_(d_3494_v113_):
                    def lambda142_(d_3495_i15_):
                        return default__.safeDivisionInt(d_3495_i15_, d_3494_v113_.f21)

                    return lambda142_

                init83_ = lambda141_(d_3485_v113_)
                nw545_ = _dafny.Array(None, 3)
                for i0_83_ in range(nw545_.length(0)):
                    nw545_[i0_83_] = init83_(i0_83_)
                d_3493_v120_ = nw545_
                d_3496_v121_: _dafny.Set
                d_3496_v121_ = _dafny.Set({(d_3485_v113_).f22})
                index522_ = default__.safeIndex(890, (d_3493_v120_).length(0))
                (d_3493_v120_)[index522_] = len(d_3496_v121_)
                d_3497_v122_: _dafny.Seq
                d_3497_v122_ = _dafny.SeqWithoutIsStrInference([self.f2])
                d_3498_v123_: D4
                d_3498_v123_ = D4_DC10(self.f2, self.f2, d_3497_v122_)
                d_3499_v124_: _dafny.Array
                nw546_ = _dafny.Array(None, 25)
                nw546_[int(0)] = self.f2
                nw546_[int(1)] = p0
                nw546_[int(2)] = (d_3498_v123_).cf25
                nw546_[int(3)] = self.f2
                nw546_[int(4)] = (d_3485_v113_).f22
                nw546_[int(5)] = self.f2
                nw546_[int(6)] = (d_3485_v113_).f22
                nw546_[int(7)] = p0
                nw546_[int(8)] = (d_3485_v113_).fm0(globalState)
                nw546_[int(9)] = p0
                nw546_[int(10)] = (d_3485_v113_).f22
                nw546_[int(11)] = not(p0)
                nw546_[int(12)] = p0
                nw546_[int(13)] = (d_3485_v113_).f22
                nw546_[int(14)] = self.f2
                nw546_[int(15)] = (d_3485_v113_).f22
                nw546_[int(16)] = (d_3485_v113_).f22
                nw546_[int(17)] = not(self.f2)
                nw546_[int(18)] = not(p0)
                nw546_[int(19)] = self.f2
                nw546_[int(20)] = self.f2
                nw546_[int(21)] = p0
                nw546_[int(22)] = p0
                nw546_[int(23)] = p0
                nw546_[int(24)] = default__.fm27(d_3320_v0_, globalState)
                d_3499_v124_ = nw546_
                d_3500_v125_: _dafny.Seq
                d_3500_v125_ = _dafny.SeqWithoutIsStrInference([d_3499_v124_])
                d_3501_v126_: _dafny.Seq
                d_3501_v126_ = _dafny.SeqWithoutIsStrInference([d_3498_v123_])
                d_3502_v127_: _dafny.Seq
                d_3502_v127_ = _dafny.SeqWithoutIsStrInference([d_3501_v126_])
                index523_ = default__.safeIndex(890, (d_3493_v120_).length(0))
                rhs473_ = (d_3493_v120_ if (d_3500_v125_) != (d_3500_v125_) else d_3493_v120_)
                rhs474_ = len((d_3501_v126_) + ((d_3502_v127_)[default__.safeIndex(d_3485_v113_.f21, len(d_3502_v127_))]))
                lhs260_ = d_3493_v120_
                lhs261_ = default__.safeIndex(890, (d_3493_v120_).length(0))
                d_3493_v120_ = rhs473_
                lhs260_[lhs261_] = rhs474_
                d_3503_v128_: C12
                nw547_ = C12()
                nw547_.ctor__((d_3493_v120_)[default__.safeIndex(890, (d_3493_v120_).length(0))])
                d_3503_v128_ = nw547_
                index524_ = default__.safeIndex(890, (d_3493_v120_).length(0))
                (d_3493_v120_)[index524_] = -463
            elif True:
                r1 = ((self).f3) == ((self).f3)
                d_3504_v129_: _dafny.Array
                nw548_ = _dafny.Array(False, 28)
                d_3504_v129_ = nw548_
                index525_ = default__.safeIndex(868, (d_3504_v129_).length(0))
                (d_3504_v129_)[index525_] = p0
                d_3505_v130_: _dafny.Map
                d_3505_v130_ = _dafny.Map({(self).f3: default__.fm26(globalState)})
                d_3506_v131_: _dafny.MultiSet
                d_3506_v131_ = _dafny.MultiSet([((d_3505_v130_)[(_dafny.MultiSet([d_3321_v1_, d_3321_v1_])).cardinality] if ((_dafny.MultiSet([d_3321_v1_, d_3321_v1_])).cardinality) in (d_3505_v130_) else (self).f3)])
                index526_ = default__.safeIndex(868, (d_3504_v129_).length(0))
                (d_3504_v129_)[index526_] = (d_3506_v131_).isdisjoint(d_3506_v131_)
                d_3507_v132_: _dafny.Seq
                d_3507_v132_ = _dafny.SeqWithoutIsStrInference([(self).f3, 112, (self).f3, (self).f3])
                index527_ = default__.safeIndex(868, (d_3504_v129_).length(0))
                rhs475_ = ((d_3321_v1_).set(default__.safeIndex((self).f3, len(d_3321_v1_)), d_3320_v0_) if self.f2 else d_3321_v1_)
                rhs476_ = self.f2
                rhs477_ = ((0) - ((d_3507_v132_)[default__.safeIndex((self).f3, len(d_3507_v132_))])) == ((self).f3)
                rhs478_ = self.f2
                lhs262_ = d_3504_v129_
                lhs263_ = default__.safeIndex(868, (d_3504_v129_).length(0))
                lhs264_ = self
                d_3321_v1_ = rhs475_
                lhs262_[lhs263_] = rhs476_
                r0 = rhs477_
                lhs264_.f2 = rhs478_
                d_3508_v133_: _dafny.Map
                d_3508_v133_ = _dafny.Map({481: self.f2})
                d_3509_v134_: C7
                nw549_ = C7()
                nw549_.ctor__((0) - (default__.safeModuloInt((self).f3, (self).f3)), (d_3504_v129_)[default__.safeIndex(868, (d_3504_v129_).length(0))], default__.fm27(d_3320_v0_, globalState), (self).f3, (520) * (len(d_3508_v133_)), (d_3504_v129_)[default__.safeIndex(868, (d_3504_v129_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "segwmrc")), (d_3504_v129_)[default__.safeIndex(868, (d_3504_v129_).length(0))])
                d_3509_v134_ = nw549_
                d_3510_v135_: _dafny.Array
                nw550_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_3510_v135_ = nw550_
                rhs479_ = d_3509_v134_
                rhs480_ = (self).f3
                rhs481_ = d_3510_v135_
                d_3509_v134_ = rhs479_
                r2 = rhs480_
                d_3510_v135_ = rhs481_
                d_3511_v136_: _dafny.Seq
                d_3511_v136_ = _dafny.SeqWithoutIsStrInference([d_3508_v133_])
                d_3512_v137_: _dafny.Map
                d_3512_v137_ = _dafny.Map({(self).f3: d_3321_v1_})
                d_3513_v138_: _dafny.Set
                d_3513_v138_ = _dafny.Set({(d_3509_v134_).f20})
                d_3514_v139_: _dafny.Seq
                d_3514_v139_ = _dafny.SeqWithoutIsStrInference([(d_3504_v129_)[default__.safeIndex(868, (d_3504_v129_).length(0))]])
                d_3515_v140_: _dafny.Array
                nw551_ = _dafny.Array(_dafny.CodePoint('D'), 14)
                d_3515_v140_ = nw551_
                d_3516_v141_: _dafny.Map
                d_3516_v141_ = _dafny.Map({d_3515_v140_: (d_3509_v134_).f19})
                d_3517_v142_: _dafny.Array
                nw552_ = _dafny.Array(None, 27)
                nw552_[int(0)] = len(((d_3321_v1_).set(default__.safeIndex((d_3509_v134_).f19, len(d_3321_v1_)), d_3320_v0_)) + (d_3321_v1_))
                nw552_[int(1)] = (d_3509_v134_).f19
                nw552_[int(2)] = (self).f3
                nw552_[int(3)] = -696
                nw552_[int(4)] = (self).f3
                nw552_[int(5)] = default__.safeDivisionInt((0) - ((d_3509_v134_).f19), len(d_3511_v136_))
                nw552_[int(6)] = (self).f3
                nw552_[int(7)] = (d_3509_v134_).f19
                nw552_[int(8)] = (d_3509_v134_).f19
                nw552_[int(9)] = (self).f3
                nw552_[int(10)] = (d_3509_v134_).f19
                nw552_[int(11)] = ((d_3506_v131_)[(0) - ((d_3509_v134_).f19)] if ((0) - ((d_3509_v134_).f19)) in (d_3506_v131_) else (d_3509_v134_).f19)
                nw552_[int(12)] = len(d_3512_v137_)
                nw552_[int(13)] = (self).f3
                nw552_[int(14)] = (d_3509_v134_).f19
                nw552_[int(15)] = default__.safeModuloInt(len(d_3513_v138_), (d_3509_v134_).f19)
                nw552_[int(16)] = (d_3509_v134_).f19
                nw552_[int(17)] = (self).f3
                nw552_[int(18)] = (len(d_3514_v139_)) * ((self).f3)
                nw552_[int(19)] = ((d_3509_v134_).f19 if p0 else len(d_3321_v1_))
                nw552_[int(20)] = default__.safeDivisionInt(len(d_3516_v141_), 438)
                nw552_[int(21)] = ((d_3506_v131_) - (d_3506_v131_)).cardinality
                nw552_[int(22)] = default__.safeModuloInt((d_3509_v134_).f19, (self).f3)
                nw552_[int(23)] = (d_3509_v134_).f19
                nw552_[int(24)] = 580
                nw552_[int(25)] = len(d_3321_v1_)
                nw552_[int(26)] = (0) - (((d_3509_v134_).f19) - (len(d_3508_v133_)))
                d_3517_v142_ = nw552_
                index528_ = default__.safeIndex(506, (d_3517_v142_).length(0))
                (d_3517_v142_)[index528_] = (self).f3
                index529_ = default__.safeIndex(506, (d_3517_v142_).length(0))
                (d_3517_v142_)[index529_] = (len(d_3507_v132_)) + ((0) - ((d_3509_v134_).f19))
            d_3518_v143_: _dafny.Array
            nw553_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_3518_v143_ = nw553_
            d_3519_v144_: _dafny.Array
            def lambda143_(d_3520_p0_):
                def lambda144_(d_3521_i16_):
                    return _dafny.SeqWithoutIsStrInference([d_3520_p0_])

                return lambda144_

            init84_ = lambda143_(p0)
            nw554_ = _dafny.Array(None, 8)
            for i0_84_ in range(nw554_.length(0)):
                nw554_[i0_84_] = init84_(i0_84_)
            d_3519_v144_ = nw554_
            index530_ = default__.safeIndex(104, (d_3518_v143_).length(0))
            (d_3518_v143_)[index530_] = d_3519_v144_
            d_3522_v145_: _dafny.Map
            d_3522_v145_ = _dafny.Map({p0: p0})
            d_3523_v146_: _dafny.Map
            d_3523_v146_ = _dafny.Map({d_3522_v145_: d_3519_v144_})
            index531_ = default__.safeIndex(104, (d_3518_v143_).length(0))
            (d_3518_v143_)[index531_] = ((d_3523_v146_)[d_3522_v145_] if (d_3522_v145_) in (d_3523_v146_) else d_3519_v144_)
            d_3524_v147_: _dafny.Array
            def lambda145_(d_3525_i17_):
                return _dafny.SeqWithoutIsStrInference([(self).f3, (self).f3, -572])

            init85_ = lambda145_
            nw555_ = _dafny.Array(None, 14)
            for i0_85_ in range(nw555_.length(0)):
                nw555_[i0_85_] = init85_(i0_85_)
            d_3524_v147_ = nw555_
            d_3526_v148_: _dafny.Seq
            d_3526_v148_ = _dafny.SeqWithoutIsStrInference([14, (self).f3, (self).f3])
            index532_ = default__.safeIndex(717, (d_3524_v147_).length(0))
            (d_3524_v147_)[index532_] = d_3526_v148_
            index533_ = default__.safeIndex(717, (d_3524_v147_).length(0))
            (d_3524_v147_)[index533_] = d_3526_v148_
            d_3527_v149_: _dafny.Array
            def lambda146_(d_3528_p0_, d_3529_v0_):
                def lambda147_(d_3530_i18_):
                    return (d_3529_v0_ if d_3528_p0_ else _dafny.CodePoint('e'))

                return lambda147_

            init86_ = lambda146_(p0, d_3320_v0_)
            nw556_ = _dafny.Array(None, 7)
            for i0_86_ in range(nw556_.length(0)):
                nw556_[i0_86_] = init86_(i0_86_)
            d_3527_v149_ = nw556_
            index534_ = default__.safeIndex(441, (d_3527_v149_).length(0))
            (d_3527_v149_)[index534_] = d_3320_v0_
            d_3531_v150_: _dafny.Array
            nw557_ = _dafny.Array(int(0), 18)
            d_3531_v150_ = nw557_
            index535_ = default__.safeIndex(607, (d_3531_v150_).length(0))
            (d_3531_v150_)[index535_] = (self).f3
            d_3532_v151_: _dafny.Array
            nw558_ = _dafny.Array(False, 4)
            d_3532_v151_ = nw558_
            d_3533_v152_: _dafny.MultiSet
            d_3533_v152_ = _dafny.MultiSet([(self).f3])
            d_3534_v153_: C2
            nw559_ = C2()
            nw559_.ctor__(d_3532_v151_, d_3533_v152_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksuo")), p0, self.f2, len(d_3321_v1_))
            d_3534_v153_ = nw559_
            d_3535_v154_: _dafny.Map
            d_3535_v154_ = _dafny.Map({d_3534_v153_: d_3320_v0_})
            d_3536_v155_: D11
            d_3536_v155_ = D11_DC30(((self).f3) == ((self).f3), d_3531_v150_, len(d_3535_v154_), p0)
            d_3537_v156_: _dafny.Set
            d_3537_v156_ = _dafny.Set({(self).f3})
            d_3538_v157_: _dafny.Map
            d_3538_v157_ = _dafny.Map({(self).f3: d_3522_v145_})
            d_3539_v158_: _dafny.Map
            d_3539_v158_ = _dafny.Map({d_3320_v0_: (0) - (((self).f3) + (len(d_3538_v157_)))})
            pat_let_tv49_ = d_3531_v150_
            index536_ = default__.safeIndex(441, (d_3527_v149_).length(0))
            index537_ = default__.safeIndex(607, (d_3531_v150_).length(0))
            def iife275_(_pat_let79_0):
                def iife276_(d_3540_dt__update__tmp_h3_):
                    def iife277_(_pat_let80_0):
                        def iife278_(d_3541_dt__update_hcf62_h0_):
                            return D11_DC30((d_3540_dt__update__tmp_h3_).cf61, d_3541_dt__update_hcf62_h0_, (d_3540_dt__update__tmp_h3_).cf63, (d_3540_dt__update__tmp_h3_).cf64)
                        return iife278_(_pat_let80_0)
                    return iife277_(pat_let_tv49_)
                return iife276_(_pat_let79_0)
            rhs482_ = (d_3321_v1_)[default__.safeIndex((self).f3, len(d_3321_v1_))]
            rhs483_ = ((d_3537_v156_).intersection(d_3537_v156_)) != (d_3537_v156_)
            rhs484_ = self.f2
            rhs485_ = len(d_3539_v158_)
            rhs486_ = iife275_(D11_DC30(p0, d_3531_v150_, (self).f3, not(p0)))
            lhs265_ = d_3527_v149_
            lhs266_ = default__.safeIndex(441, (d_3527_v149_).length(0))
            lhs267_ = d_3531_v150_
            lhs268_ = default__.safeIndex(607, (d_3531_v150_).length(0))
            lhs265_[lhs266_] = rhs482_
            r1 = rhs483_
            r1 = rhs484_
            lhs267_[lhs268_] = rhs485_
            d_3536_v155_ = rhs486_
        d_3542_v159_: _dafny.MultiSet
        d_3542_v159_ = _dafny.MultiSet([self.f2, self.f2, self.f2])
        d_3543_v160_: _dafny.Seq
        d_3543_v160_ = _dafny.SeqWithoutIsStrInference([(self).f3])
        d_3544_i19_: int
        d_3544_i19_ = 0
        with _dafny.label("12"):
            while ((((d_3542_v159_)[p0] if (p0) in (d_3542_v159_) else len(d_3543_v160_))) > ((self).f3) if self.f2 else p0):
                with _dafny.c_label("12"):
                    if (d_3544_i19_) >= (100):
                        raise _dafny.Break("12")
                    d_3544_i19_ = (d_3544_i19_) + (1)
                    d_3545_v161_: _dafny.Array
                    def lambda148_(d_3546_p0_):
                        def lambda149_(d_3547_i20_):
                            return d_3546_p0_

                        return lambda149_

                    init87_ = lambda148_(p0)
                    nw560_ = _dafny.Array(None, 26)
                    for i0_87_ in range(nw560_.length(0)):
                        nw560_[i0_87_] = init87_(i0_87_)
                    d_3545_v161_ = nw560_
                    index538_ = default__.safeIndex(973, (d_3545_v161_).length(0))
                    (d_3545_v161_)[index538_] = not(self.f2)
                    index539_ = default__.safeIndex(973, (d_3545_v161_).length(0))
                    (d_3545_v161_)[index539_] = p0
                    d_3548_v162_: C6
                    nw561_ = C6()
                    nw561_.ctor__((0) - (len(_dafny.SeqWithoutIsStrInference([d_3543_v160_ for d_3549_i21_ in range(default__.abs(930))]))), p0, p0, (self).f3, (self).f3, p0)
                    d_3548_v162_ = nw561_
                    d_3550_v163_: _dafny.Seq
                    d_3550_v163_ = _dafny.SeqWithoutIsStrInference([d_3548_v162_, d_3548_v162_])
                    d_3551_v164_: _dafny.Map
                    d_3551_v164_ = _dafny.Map({d_3320_v0_: d_3548_v162_})
                    d_3552_v165_: _dafny.Array
                    nw562_ = _dafny.Array(None, 16)
                    nw562_[int(0)] = d_3548_v162_
                    nw562_[int(1)] = d_3548_v162_
                    nw562_[int(2)] = (d_3548_v162_ if self.f2 else (d_3550_v163_)[default__.safeIndex(d_3548_v162_.f21, len(d_3550_v163_))])
                    nw562_[int(3)] = d_3548_v162_
                    nw562_[int(4)] = d_3548_v162_
                    nw562_[int(5)] = ((d_3551_v164_)[d_3320_v0_] if (d_3320_v0_) in (d_3551_v164_) else (d_3550_v163_)[default__.safeIndex(len(d_3321_v1_), len(d_3550_v163_))])
                    nw562_[int(6)] = d_3548_v162_
                    nw562_[int(7)] = d_3548_v162_
                    nw562_[int(8)] = d_3548_v162_
                    nw562_[int(9)] = d_3548_v162_
                    nw562_[int(10)] = d_3548_v162_
                    nw562_[int(11)] = d_3548_v162_
                    nw562_[int(12)] = d_3548_v162_
                    nw562_[int(13)] = d_3548_v162_
                    nw562_[int(14)] = d_3548_v162_
                    nw562_[int(15)] = d_3548_v162_
                    d_3552_v165_ = nw562_
                    index540_ = default__.safeIndex(679, (d_3552_v165_).length(0))
                    (d_3552_v165_)[index540_] = d_3548_v162_
                    index541_ = default__.safeIndex(679, (d_3552_v165_).length(0))
                    (d_3552_v165_)[index541_] = (d_3548_v162_ if not((d_3545_v161_)[default__.safeIndex(973, (d_3545_v161_).length(0))]) else d_3548_v162_)
                    if p0:
                        d_3553_v166_: _dafny.Seq
                        d_3553_v166_ = _dafny.SeqWithoutIsStrInference([default__.fm27(d_3320_v0_, globalState)])
                        d_3554_v167_: _dafny.Array
                        nw563_ = _dafny.Array(None, 3)
                        nw563_[int(0)] = d_3548_v162_.f21
                        nw563_[int(1)] = len((d_3553_v166_).set(default__.safeIndex((self).f3, len(d_3553_v166_)), (d_3545_v161_)[default__.safeIndex(973, (d_3545_v161_).length(0))]))
                        nw563_[int(2)] = d_3548_v162_.f21
                        d_3554_v167_ = nw563_
                        d_3555_v168_: _dafny.Map
                        d_3555_v168_ = _dafny.Map({d_3548_v162_.f21: d_3554_v167_})
                        d_3556_v169_: D9
                        d_3556_v169_ = D9_DC21(d_3555_v168_)
                        d_3557_v170_: _dafny.Array
                        nw564_ = _dafny.Array(None, 12)
                        nw564_[int(0)] = d_3556_v169_
                        nw564_[int(1)] = d_3556_v169_
                        nw564_[int(2)] = D9_DC21(d_3555_v168_)
                        nw564_[int(3)] = d_3556_v169_
                        nw564_[int(4)] = D9_DC21(d_3555_v168_)
                        nw564_[int(5)] = d_3556_v169_
                        nw564_[int(6)] = d_3556_v169_
                        nw564_[int(7)] = D9_DC21(d_3555_v168_)
                        nw564_[int(8)] = d_3556_v169_
                        nw564_[int(9)] = d_3556_v169_
                        nw564_[int(10)] = d_3556_v169_
                        nw564_[int(11)] = d_3556_v169_
                        d_3557_v170_ = nw564_
                        d_3558_v171_: _dafny.Array
                        d_3558_v171_ = d_3557_v170_
                        d_3559_v172_: _dafny.Array
                        nw565_ = _dafny.Array(None, 2)
                        nw565_[int(0)] = d_3558_v171_
                        nw565_[int(1)] = d_3558_v171_
                        d_3559_v172_ = nw565_
                        d_3560_v173_: _dafny.MultiSet
                        d_3560_v173_ = _dafny.MultiSet([d_3559_v172_, d_3559_v172_, (d_3559_v172_ if not(self.f2) else d_3559_v172_), d_3559_v172_, d_3559_v172_])
                        r3 = ((d_3560_v173_)[d_3559_v172_] if (d_3559_v172_) in (d_3560_v173_) else len(d_3321_v1_))
                        d_3561_v175_: _dafny.MultiSet
                        d_3561_v175_ = _dafny.MultiSet([d_3545_v161_])
                        d_3562_v176_: _dafny.Map
                        d_3562_v176_ = _dafny.Map({(d_3561_v175_).cardinality: d_3548_v162_.f21})
                        d_3563_v177_: _dafny.Map
                        d_3563_v177_ = _dafny.Map({d_3562_v176_: (d_3548_v162_).f22})
                        d_3564_v178_: _dafny.Map
                        def iife279_():
                            coll117_ = _dafny.Map()
                            compr_117_: int
                            for compr_117_ in _dafny.IntegerRange(214, 135):
                                d_3565_v174_: int = compr_117_
                                if ((214) <= (d_3565_v174_)) and ((d_3565_v174_) < (135)):
                                    coll117_[default__.safeDivisionInt(d_3565_v174_, d_3548_v162_.f21)] = d_3548_v162_.f21
                            return _dafny.Map(coll117_)
                        d_3564_v178_ = _dafny.Map({(iife279_()
                        ) in (d_3563_v177_): (d_3548_v162_).f22})
                        d_3566_v179_: C3
                        nw566_ = C3()
                        nw566_.ctor__(d_3320_v0_, (self).f3, not(p0), (d_3545_v161_)[default__.safeIndex(973, (d_3545_v161_).length(0))], len(d_3321_v1_))
                        d_3566_v179_ = nw566_
                        d_3567_v180_: _dafny.Set
                        d_3567_v180_ = _dafny.Set({d_3566_v179_})
                        d_3568_v181_: D0
                        d_3568_v181_ = D0_DC2(not((d_3548_v162_).f22), (d_3548_v162_).f22)
                        index542_ = default__.safeIndex(973, (d_3545_v161_).length(0))
                        index543_ = default__.safeIndex(973, (d_3545_v161_).length(0))
                        rhs487_ = (d_3548_v162_).fm0(globalState)
                        rhs488_ = ((d_3564_v178_)[(d_3567_v180_).isdisjoint(d_3567_v180_)] if ((d_3567_v180_).isdisjoint(d_3567_v180_)) in (d_3564_v178_) else (d_3568_v181_).cf9)
                        rhs489_ = d_3548_v162_.f21
                        rhs490_ = 652
                        lhs269_ = d_3545_v161_
                        lhs270_ = default__.safeIndex(973, (d_3545_v161_).length(0))
                        lhs271_ = d_3545_v161_
                        lhs272_ = default__.safeIndex(973, (d_3545_v161_).length(0))
                        lhs269_[lhs270_] = rhs487_
                        lhs271_[lhs272_] = rhs488_
                        r3 = rhs489_
                        r3 = rhs490_
                        d_3569_v182_: D11
                        d_3569_v182_ = D11_DC29(d_3548_v162_.f21, d_3545_v161_, (d_3548_v162_).f22, (d_3545_v161_)[default__.safeIndex(973, (d_3545_v161_).length(0))])
                        index544_ = default__.safeIndex(973, (d_3545_v161_).length(0))
                        (d_3545_v161_)[index544_] = (d_3569_v182_).cf60
                        d_3570_v183_: C5
                        nw567_ = C5()
                        nw567_.ctor__()
                        d_3570_v183_ = nw567_
                        d_3571_v184_: _dafny.Array
                        def lambda150_(d_3572_v161_, d_3573_v179_, d_3574_v1_):
                            def lambda151_(d_3575_i22_):
                                return D12_DC33((d_3572_v161_)[default__.safeIndex(973, (d_3572_v161_).length(0))], self.f2, (d_3573_v179_).f25, len(d_3574_v1_))

                            return lambda151_

                        init88_ = lambda150_(d_3545_v161_, d_3566_v179_, d_3321_v1_)
                        nw568_ = _dafny.Array(None, 25)
                        for i0_88_ in range(nw568_.length(0)):
                            nw568_[i0_88_] = init88_(i0_88_)
                        d_3571_v184_ = nw568_
                        d_3576_v185_: D23
                        d_3576_v185_ = D23_DC61(d_3571_v184_)
                        rhs491_ = _dafny.SeqWithoutIsStrInference([(D16_DC41(d_3320_v0_, (d_3548_v162_).f22)).cf77 for d_3577_i23_ in range(default__.abs(381))])
                        rhs492_ = ((d_3576_v185_).cf108 if False else d_3571_v184_)
                        d_3321_v1_ = rhs491_
                        d_3571_v184_ = rhs492_
                    elif True:
                        d_3578_v186_: C13
                        nw569_ = C13()
                        nw569_.ctor__(False, (d_3548_v162_).f22, (d_3545_v161_)[default__.safeIndex(973, (d_3545_v161_).length(0))], d_3548_v162_.f21)
                        d_3578_v186_ = nw569_
                        (d_3548_v162_).f21 = default__.safeDivisionInt(d_3548_v162_.f21, d_3548_v162_.f21)
                        d_3579_v187_: _dafny.Array
                        nw570_ = _dafny.Array(int(0), 10)
                        d_3579_v187_ = nw570_
                        d_3579_v187_ = d_3579_v187_
                        d_3580_v188_: C10
                        nw571_ = C10()
                        nw571_.ctor__((d_3548_v162_.f21) + ((self).f3), d_3545_v161_)
                        d_3580_v188_ = nw571_
                        d_3580_v188_ = d_3580_v188_
                        d_3581_v189_: _dafny.MultiSet
                        d_3581_v189_ = _dafny.MultiSet([(d_3543_v160_)[default__.safeIndex((self).f3, len(d_3543_v160_))]])
                        d_3582_v190_: _dafny.Map
                        d_3582_v190_ = _dafny.Map({d_3578_v186_.f7: d_3321_v1_})
                        d_3583_v191_: _dafny.Seq
                        d_3583_v191_ = _dafny.SeqWithoutIsStrInference([(d_3580_v188_).f11])
                        d_3584_v192_: _dafny.Seq
                        d_3584_v192_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, not(self.f2)])])
                        d_3585_v193_: _dafny.Set
                        d_3585_v193_ = _dafny.Set({len((d_3584_v192_)[default__.safeIndex((d_3580_v188_).f10, len(d_3584_v192_))])})
                        d_3586_v194_: _dafny.MultiSet
                        d_3586_v194_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(d_3581_v189_).cardinality])), (0) - (d_3548_v162_.f21), d_3548_v162_.f21, len((d_3582_v190_).set((d_3578_v186_).fm2(d_3548_v162_.f21, len(d_3583_v191_), (d_3543_v160_)[default__.safeIndex((self).f3, len(d_3543_v160_))], globalState), default__.fm52(default__.fm27(d_3320_v0_, globalState), d_3585_v193_, d_3548_v162_.f21, (d_3548_v162_).f22, globalState))), (d_3581_v189_).cardinality])
                        d_3587_v195_: C2
                        nw572_ = C2()
                        nw572_.ctor__(d_3545_v161_, d_3586_v194_, d_3321_v1_, d_3578_v186_.f7, d_3578_v186_.f7, (0) - (d_3548_v162_.f21))
                        d_3587_v195_ = nw572_
                    r2 = ((d_3548_v162_.f21) - (default__.fm49(d_3321_v1_, len(_dafny.SeqWithoutIsStrInference([p0, self.f2])), p0, (d_3548_v162_).f22, globalState))) - ((0) - (d_3548_v162_.f21))
                    pass
            pass
        r0 = (_dafny.CodePoint('y')) in (default__.fm28(d_3321_v1_, (0) - ((self).f3), (self).f3, globalState))
        r1 = not(self.f2)
        d_3588_v196_: _dafny.MultiSet
        d_3588_v196_ = _dafny.MultiSet([(self).f3, (self).f3, (self).f3])
        r2 = ((d_3588_v196_).cardinality) * ((0) - ((935) * ((self).f3)))
        r3 = (self).f3
        return r0, r1, r2, r3

    def m2(self, p0, globalState):
        d_3589_v0_: int
        d_3589_v0_ = 430
        d_3590_v1_: str
        d_3590_v1_ = _dafny.CodePoint('k')
        d_3591_v2_: _dafny.MultiSet
        d_3591_v2_ = _dafny.MultiSet([self.f2, default__.fm27(d_3590_v1_, globalState), self.f2])
        d_3592_v3_: D9
        d_3592_v3_ = D9_DC23(self.f2, (d_3591_v2_).cardinality, not(not(self.f2)))
        pat_let_tv50_ = p0
        pat_let_tv51_ = d_3589_v0_
        pat_let_tv52_ = p0
        pat_let_tv53_ = d_3589_v0_
        def lambda152_(source42_):
            if source42_.is_DC22:
                d_3593___mcc_h0_ = source42_.cf44
                d_3594_cf44_ = d_3593___mcc_h0_
                return pat_let_tv50_
            elif source42_.is_DC23:
                d_3595___mcc_h1_ = source42_.cf45
                d_3596___mcc_h2_ = source42_.cf46
                d_3597___mcc_h3_ = source42_.cf47
                d_3598_cf47_ = d_3597___mcc_h3_
                d_3599_cf46_ = d_3596___mcc_h2_
                d_3600_cf45_ = d_3595___mcc_h1_
                return default__.safeModuloInt(pat_let_tv51_, len(_dafny.SeqWithoutIsStrInference([d_3598_cf47_, not(d_3598_cf47_), self.f2])))
            elif source42_.is_DC24:
                d_3601___mcc_h4_ = source42_.cf48
                d_3602_cf48_ = d_3601___mcc_h4_
                return (self).f3
            elif source42_.is_DC21:
                d_3603___mcc_h5_ = source42_.cf43
                d_3604_cf43_ = d_3603___mcc_h5_
                return len((_dafny.Map({self.f2: len(_dafny.SeqWithoutIsStrInference([pat_let_tv52_]))})) | (_dafny.Map({self.f2: pat_let_tv53_})))
            elif True:
                d_3605___mcc_h6_ = source42_.cf49
                d_3606_cf49_ = d_3605___mcc_h6_
                return (self).f3

        d_3589_v0_ = lambda152_(d_3592_v3_)
        d_3607_v4_: _dafny.Map
        d_3607_v4_ = _dafny.Map({False: self.f2})
        d_3608_v5_: _dafny.Set
        d_3608_v5_ = _dafny.Set({(0) - (p0), (default__.fm68(globalState)).cardinality, p0})
        d_3609_i0_: int
        d_3609_i0_ = 0
        with _dafny.label("13"):
            while (((d_3607_v4_)[self.f2] if (self.f2) in (d_3607_v4_) else self.f2) if self.f2 else not((d_3608_v5_).isdisjoint(d_3608_v5_))):
                with _dafny.c_label("13"):
                    if (d_3609_i0_) >= (100):
                        raise _dafny.Break("13")
                    d_3609_i0_ = (d_3609_i0_) + (1)
                    d_3610_v6_: _dafny.Set
                    d_3610_v6_ = _dafny.Set({self.f2})
                    d_3611_v7_: _dafny.Seq
                    d_3611_v7_ = _dafny.SeqWithoutIsStrInference([len(d_3610_v6_)])
                    d_3612_v8_: C8
                    nw573_ = C8()
                    nw573_.ctor__(default__.fm52(self.f2, d_3608_v5_, len(d_3611_v7_), self.f2, globalState), (self).f3)
                    d_3612_v8_ = nw573_
                    d_3613_v9_: _dafny.Seq
                    d_3613_v9_ = _dafny.SeqWithoutIsStrInference([self.f2])
                    d_3614_v10_: _dafny.Seq
                    d_3614_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f2])])
                    d_3615_v11_: _dafny.Map
                    d_3615_v11_ = _dafny.Map({(d_3613_v9_) <= ((d_3614_v10_)[default__.safeIndex(p0, len(d_3614_v10_))]): default__.safeDivisionInt((0) - ((0) - (d_3589_v0_)), (d_3612_v8_).f14)})
                    d_3616_v12_: _dafny.MultiSet
                    d_3616_v12_ = _dafny.MultiSet([len(d_3612_v8_.f13), p0])
                    d_3615_v11_ = (d_3615_v11_).set(self.f2, (d_3616_v12_).cardinality)
                    d_3617_v13_: _dafny.Array
                    def lambda153_(d_3618_v12_, d_3619_v8_):
                        def lambda154_(d_3620_i1_):
                            return (D1_DC5(self.f2, (d_3618_v12_).cardinality, d_3619_v8_.f13)).cf17

                        return lambda154_

                    init89_ = lambda153_(d_3616_v12_, d_3612_v8_)
                    nw574_ = _dafny.Array(None, 18)
                    for i0_89_ in range(nw574_.length(0)):
                        nw574_[i0_89_] = init89_(i0_89_)
                    d_3617_v13_ = nw574_
                    index545_ = default__.safeIndex(16, (d_3617_v13_).length(0))
                    (d_3617_v13_)[index545_] = d_3612_v8_.f13
                    d_3621_v14_: _dafny.Map
                    d_3621_v14_ = _dafny.Map({self.f2: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rl"))})
                    d_3622_v15_: D18
                    d_3622_v15_ = D18_DC50(self.f2, self.f2)
                    d_3623_v16_: _dafny.Set
                    d_3623_v16_ = _dafny.Set({d_3590_v1_})
                    index546_ = default__.safeIndex(16, (d_3617_v13_).length(0))
                    rhs493_ = ((d_3621_v14_)[(d_3622_v15_).cf90] if ((d_3622_v15_).cf90) in (d_3621_v14_) else (d_3612_v8_.f13) + (_dafny.SeqWithoutIsStrInference([d_3590_v1_ for d_3624_i2_ in range(default__.abs(43))])))
                    rhs494_ = default__.fm18(self.f2, (d_3623_v16_).ispropersubset(d_3623_v16_), self.f2, globalState)
                    rhs495_ = default__.safeModuloInt((d_3612_v8_).f14, len((d_3608_v5_) | (d_3608_v5_)))
                    lhs273_ = d_3617_v13_
                    lhs274_ = default__.safeIndex(16, (d_3617_v13_).length(0))
                    lhs273_[lhs274_] = rhs493_
                    d_3589_v0_ = rhs494_
                    d_3589_v0_ = rhs495_
                    d_3625_v17_: _dafny.Array
                    nw575_ = _dafny.Array(False, 18)
                    d_3625_v17_ = nw575_
                    index547_ = default__.safeIndex(980, (d_3625_v17_).length(0))
                    (d_3625_v17_)[index547_] = self.f2
                    index548_ = default__.safeIndex(980, (d_3625_v17_).length(0))
                    (d_3625_v17_)[index548_] = (self.f2 if not((d_3613_v9_)[default__.safeIndex(len(d_3611_v7_), len(d_3613_v9_))]) else not(self.f2))
                    pass
            pass
        d_3626_v18_: _dafny.Map
        d_3626_v18_ = _dafny.Map({(0) - (d_3589_v0_): p0})
        hi13_ = default__.safeModuloInt(d_3589_v0_, (self).f3)
        for d_3627_i3_ in range(len((d_3626_v18_) | (d_3626_v18_)), hi13_):
            (self).f2 = self.f2
            d_3628_v19_: _dafny.Array
            nw576_ = _dafny.Array(False, 16)
            d_3628_v19_ = nw576_
            d_3629_v20_: D4
            d_3629_v20_ = D4_DC8(d_3628_v19_)
            d_3630_v21_: _dafny.Array
            nw577_ = _dafny.Array(None, 14)
            nw577_[int(0)] = d_3628_v19_
            nw577_[int(1)] = d_3628_v19_
            nw577_[int(2)] = d_3628_v19_
            nw577_[int(3)] = d_3628_v19_
            nw577_[int(4)] = d_3628_v19_
            nw577_[int(5)] = d_3628_v19_
            nw577_[int(6)] = d_3628_v19_
            nw577_[int(7)] = d_3628_v19_
            nw577_[int(8)] = d_3628_v19_
            nw577_[int(9)] = (d_3629_v20_).cf20
            nw577_[int(10)] = d_3628_v19_
            nw577_[int(11)] = d_3628_v19_
            nw577_[int(12)] = d_3628_v19_
            nw577_[int(13)] = d_3628_v19_
            d_3630_v21_ = nw577_
            index549_ = default__.safeIndex(88, (d_3630_v21_).length(0))
            (d_3630_v21_)[index549_] = d_3628_v19_
            d_3631_v22_: _dafny.Map
            d_3631_v22_ = _dafny.Map({self.f2: d_3627_i3_})
            index550_ = default__.safeIndex(88, (d_3630_v21_).length(0))
            rhs496_ = self.f2
            rhs497_ = ((296) * ((self).f3)) - (d_3627_i3_)
            rhs498_ = d_3628_v19_
            rhs499_ = ((self).f3) * ((len(d_3631_v22_)) * (d_3589_v0_))
            rhs500_ = (self).f3
            lhs275_ = self
            lhs276_ = d_3630_v21_
            lhs277_ = default__.safeIndex(88, (d_3630_v21_).length(0))
            lhs275_.f2 = rhs496_
            d_3589_v0_ = rhs497_
            lhs276_[lhs277_] = rhs498_
            d_3589_v0_ = rhs499_
            d_3589_v0_ = rhs500_
            (self).f2 = True
            d_3632_v23_: _dafny.Seq
            d_3632_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulshfw"))
            d_3632_v23_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sywftil")) if self.f2 else (d_3632_v23_) + (d_3632_v23_))
        d_3633_v24_: D16
        d_3633_v24_ = D16_DC41(d_3590_v1_, not(self.f2))
        d_3634_v25_: _dafny.MultiSet
        d_3634_v25_ = _dafny.MultiSet([(d_3633_v24_).cf77, d_3590_v1_, d_3590_v1_, d_3590_v1_])
        d_3635_v26_: _dafny.Set
        d_3635_v26_ = _dafny.Set({self.f2})
        d_3636_v27_: _dafny.Seq
        d_3636_v27_ = _dafny.SeqWithoutIsStrInference([self.f2])
        d_3637_v28_: _dafny.Seq
        d_3637_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hehv"))
        d_3638_v29_: _dafny.Seq
        d_3638_v29_ = _dafny.SeqWithoutIsStrInference([d_3589_v0_, (self).f3])
        d_3639_v30_: _dafny.MultiSet
        d_3639_v30_ = _dafny.MultiSet([d_3634_v25_, d_3634_v25_])
        d_3640_v31_: D24
        d_3640_v31_ = D24_DC63(d_3639_v30_)
        pat_let_tv54_ = d_3634_v25_
        d_3641_v32_: _dafny.MultiSet
        def iife280_(_pat_let81_0):
            def iife281_(d_3642_dt__update__tmp_h0_):
                def iife282_(_pat_let82_0):
                    def iife283_(d_3644_dt__update_hcf112_h0_):
                        return D24_DC63(d_3644_dt__update_hcf112_h0_)
                    return iife283_(_pat_let82_0)
                return iife282_(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([pat_let_tv54_ for d_3643_i4_ in range(default__.abs(716))])))
            return iife281_(_pat_let81_0)
        d_3641_v32_ = _dafny.MultiSet([(self).f3, ((iife280_(d_3640_v31_)).cf112).cardinality])
        d_3645_v33_: _dafny.Array
        nw578_ = _dafny.Array(None, 29)
        nw578_[int(0)] = not((self.f2) not in (d_3591_v2_))
        nw578_[int(1)] = self.f2
        nw578_[int(2)] = False
        nw578_[int(3)] = self.f2
        nw578_[int(4)] = not((d_3634_v25_).ispropersubset(d_3634_v25_))
        nw578_[int(5)] = self.f2
        nw578_[int(6)] = self.f2
        nw578_[int(7)] = self.f2
        nw578_[int(8)] = (d_3635_v26_).isdisjoint(d_3635_v26_)
        nw578_[int(9)] = (self.f2) == (self.f2)
        nw578_[int(10)] = self.f2
        nw578_[int(11)] = self.f2
        nw578_[int(12)] = self.f2
        nw578_[int(13)] = True
        nw578_[int(14)] = not(self.f2)
        nw578_[int(15)] = (d_3636_v27_)[default__.safeIndex(-327, len(d_3636_v27_))]
        nw578_[int(16)] = (d_3637_v28_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))
        nw578_[int(17)] = self.f2
        nw578_[int(18)] = self.f2
        nw578_[int(19)] = True
        nw578_[int(20)] = not (default__.fm27(d_3590_v1_, globalState)) or (self.f2)
        nw578_[int(21)] = not(self.f2)
        nw578_[int(22)] = (p0) >= (len(d_3638_v29_))
        nw578_[int(23)] = self.f2
        nw578_[int(24)] = self.f2
        nw578_[int(25)] = not(self.f2)
        nw578_[int(26)] = (d_3641_v32_).ispropersubset(d_3641_v32_)
        nw578_[int(27)] = default__.fm27(d_3590_v1_, globalState)
        nw578_[int(28)] = self.f2
        d_3645_v33_ = nw578_
        d_3645_v33_ = d_3645_v33_
        d_3589_v0_ = ((d_3591_v2_)[self.f2] if (self.f2) in (d_3591_v2_) else default__.fm3(globalState))
        d_3646_i5_: int
        d_3646_i5_ = 0
        with _dafny.label("14"):
            while self.f2:
                with _dafny.c_label("14"):
                    if (d_3646_i5_) >= (100):
                        raise _dafny.Break("14")
                    d_3646_i5_ = (d_3646_i5_) + (1)
                    d_3647_v34_: _dafny.Array
                    nw579_ = _dafny.Array(int(0), 18)
                    d_3647_v34_ = nw579_
                    index551_ = default__.safeIndex(457, (d_3647_v34_).length(0))
                    (d_3647_v34_)[index551_] = p0
                    index552_ = default__.safeIndex(457, (d_3647_v34_).length(0))
                    (d_3647_v34_)[index552_] = (self).f3
                    (globalState).f0 = (d_3635_v26_) | ((_dafny.Set({self.f2, self.f2})) - (d_3635_v26_))
                    if (self.f2) == (False):
                        index553_ = default__.safeIndex(457, (d_3647_v34_).length(0))
                        (d_3647_v34_)[index553_] = (self).f3
                        d_3648_v35_: _dafny.Map
                        d_3648_v35_ = _dafny.Map({self.f2: (d_3647_v34_)[default__.safeIndex(457, (d_3647_v34_).length(0))]})
                        d_3649_v36_: _dafny.Map
                        d_3649_v36_ = d_3648_v35_
                        d_3649_v36_ = d_3649_v36_
                        d_3650_v37_: C5
                        nw580_ = C5()
                        nw580_.ctor__()
                        d_3650_v37_ = nw580_
                        index554_ = default__.safeIndex(602, (d_3645_v33_).length(0))
                        (d_3645_v33_)[index554_] = (self.f2) == (self.f2)
                        index555_ = default__.safeIndex(910, (d_3645_v33_).length(0))
                        (d_3645_v33_)[index555_] = self.f2
                        d_3651_v38_: _dafny.Map
                        d_3651_v38_ = _dafny.Map({(d_3641_v32_).cardinality: self.f2})
                        d_3652_v39_: D16
                        d_3652_v39_ = D16_DC40(d_3651_v38_)
                        d_3653_v40_: _dafny.Seq
                        d_3653_v40_ = _dafny.SeqWithoutIsStrInference([(d_3636_v27_).set(default__.safeIndex(d_3589_v0_, len(d_3636_v27_)), self.f2)])
                        d_3654_v41_: _dafny.Map
                        d_3654_v41_ = _dafny.Map({d_3652_v39_: (d_3653_v40_)[default__.safeIndex(default__.fm26(globalState), len(d_3653_v40_))]})
                        d_3655_v42_: _dafny.Seq
                        d_3655_v42_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({self.f2: self.f2}), _dafny.Map({self.f2: self.f2})])
                        d_3656_v43_: _dafny.Seq
                        d_3656_v43_ = _dafny.SeqWithoutIsStrInference([d_3654_v41_, d_3654_v41_])
                        index556_ = default__.safeIndex(602, (d_3645_v33_).length(0))
                        index557_ = default__.safeIndex(910, (d_3645_v33_).length(0))
                        rhs501_ = not(self.f2)
                        rhs502_ = ((d_3655_v42_)[default__.safeIndex(d_3589_v0_, len(d_3655_v42_))]) | (_dafny.Map({self.f2: self.f2}))
                        rhs503_ = default__.fm18(False, False, self.f2, globalState)
                        rhs504_ = default__.fm27((d_3590_v1_ if self.f2 else d_3590_v1_), globalState)
                        rhs505_ = (d_3656_v43_)[default__.safeIndex(default__.safeModuloInt(p0, (self).f3), len(d_3656_v43_))]
                        lhs278_ = d_3645_v33_
                        lhs279_ = default__.safeIndex(602, (d_3645_v33_).length(0))
                        lhs280_ = d_3645_v33_
                        lhs281_ = default__.safeIndex(910, (d_3645_v33_).length(0))
                        lhs278_[lhs279_] = rhs501_
                        d_3607_v4_ = rhs502_
                        d_3589_v0_ = rhs503_
                        lhs280_[lhs281_] = rhs504_
                        d_3654_v41_ = rhs505_
                        d_3657_v44_: _dafny.Set
                        d_3657_v44_ = _dafny.Set({d_3626_v18_, _dafny.Map({p0: default__.fm3(globalState)})})
                        d_3658_v45_: _dafny.Set
                        d_3658_v45_ = d_3657_v44_
                        d_3659_v46_: C0
                        nw581_ = C0()
                        nw581_.ctor__((d_3658_v45_ if not((d_3645_v33_)[default__.safeIndex(602, (d_3645_v33_).length(0))]) else d_3658_v45_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))))
                        d_3659_v46_ = nw581_
                    elif True:
                        d_3660_v47_: _dafny.Array
                        nw582_ = _dafny.Array(_dafny.Array(None, 0), 2)
                        d_3660_v47_ = nw582_
                        d_3660_v47_ = d_3660_v47_
                        d_3661_v48_: _dafny.Map
                        d_3661_v48_ = (d_3607_v4_).set(self.f2, self.f2)
                        d_3662_v49_: _dafny.Array
                        nw583_ = _dafny.Array(None, 28)
                        nw583_[int(0)] = d_3607_v4_
                        nw583_[int(1)] = (d_3607_v4_) | (_dafny.Map({not(self.f2): self.f2}))
                        nw583_[int(2)] = d_3607_v4_
                        nw583_[int(3)] = (d_3607_v4_) | (d_3607_v4_)
                        nw583_[int(4)] = (d_3661_v48_)
                        nw583_[int(5)] = (d_3607_v4_) | (_dafny.Map({self.f2: True}))
                        nw583_[int(6)] = d_3607_v4_
                        nw583_[int(7)] = d_3607_v4_
                        nw583_[int(8)] = d_3607_v4_
                        nw583_[int(9)] = (d_3607_v4_) | ((d_3607_v4_).set(self.f2, self.f2))
                        nw583_[int(10)] = d_3607_v4_
                        nw583_[int(11)] = _dafny.Map({self.f2: self.f2})
                        nw583_[int(12)] = d_3607_v4_
                        nw583_[int(13)] = d_3607_v4_
                        nw583_[int(14)] = d_3607_v4_
                        nw583_[int(15)] = d_3607_v4_
                        nw583_[int(16)] = ((d_3607_v4_).set(self.f2, self.f2)) | (d_3607_v4_)
                        nw583_[int(17)] = d_3607_v4_
                        nw583_[int(18)] = _dafny.Map({self.f2: self.f2})
                        nw583_[int(19)] = d_3607_v4_
                        nw583_[int(20)] = d_3607_v4_
                        nw583_[int(21)] = d_3607_v4_
                        nw583_[int(22)] = d_3607_v4_
                        nw583_[int(23)] = d_3607_v4_
                        nw583_[int(24)] = d_3607_v4_
                        nw583_[int(25)] = d_3607_v4_
                        nw583_[int(26)] = d_3607_v4_
                        nw583_[int(27)] = d_3607_v4_
                        d_3662_v49_ = nw583_
                        d_3663_v50_: _dafny.Map
                        d_3663_v50_ = _dafny.Map({False: d_3662_v49_})
                        d_3662_v49_ = ((d_3663_v50_)[self.f2] if (self.f2) in (d_3663_v50_) else d_3662_v49_)
                        d_3637_v28_ = d_3637_v28_
                        d_3664_v51_: _dafny.Array
                        nw584_ = _dafny.Array(_dafny.Array(None, 0), 25)
                        d_3664_v51_ = nw584_
                        d_3665_v52_: _dafny.Array
                        nw585_ = _dafny.Array(D6.default()(), 8)
                        d_3665_v52_ = nw585_
                        index558_ = default__.safeIndex(149, (d_3664_v51_).length(0))
                        (d_3664_v51_)[index558_] = d_3665_v52_
                        index559_ = default__.safeIndex(149, (d_3664_v51_).length(0))
                        (d_3664_v51_)[index559_] = d_3665_v52_
                        d_3666_v53_: D22
                        d_3666_v53_ = D22_DC58(d_3637_v28_, (d_3647_v34_)[default__.safeIndex(457, (d_3647_v34_).length(0))], self.f2)
                        d_3637_v28_ = (d_3666_v53_).cf102
                    d_3667_v54_: C1
                    nw586_ = C1()
                    nw586_.ctor__(default__.safeModuloInt(d_3589_v0_, 711), not(not(self.f2)), self.f2, 992)
                    d_3667_v54_ = nw586_
                    d_3668_v55_: _dafny.Array
                    def lambda155_(d_3669_p0_):
                        def lambda156_(d_3670_i6_):
                            return _dafny.SeqWithoutIsStrInference([d_3669_p0_])

                        return lambda156_

                    init90_ = lambda155_(p0)
                    nw587_ = _dafny.Array(None, 8)
                    for i0_90_ in range(nw587_.length(0)):
                        nw587_[i0_90_] = init90_(i0_90_)
                    d_3668_v55_ = nw587_
                    index560_ = default__.safeIndex(700, (d_3668_v55_).length(0))
                    (d_3668_v55_)[index560_] = (d_3638_v29_) + (d_3638_v29_)
                    d_3671_v56_: _dafny.Seq
                    d_3671_v56_ = _dafny.SeqWithoutIsStrInference([d_3667_v54_])
                    index561_ = default__.safeIndex(457, (d_3647_v34_).length(0))
                    index562_ = default__.safeIndex(700, (d_3668_v55_).length(0))
                    rhs506_ = (d_3671_v56_)[default__.safeIndex(default__.fm25(d_3589_v0_, (d_3647_v34_)[default__.safeIndex(457, (d_3647_v34_).length(0))], self.f2, globalState), len(d_3671_v56_))]
                    rhs507_ = d_3633_v24_
                    rhs508_ = ((p0) - ((self).f3)) + (-82)
                    rhs509_ = ((_dafny.SeqWithoutIsStrInference([(self).f3]) if self.f2 else d_3638_v29_)) + (d_3638_v29_)
                    lhs282_ = d_3647_v34_
                    lhs283_ = default__.safeIndex(457, (d_3647_v34_).length(0))
                    lhs284_ = d_3668_v55_
                    lhs285_ = default__.safeIndex(700, (d_3668_v55_).length(0))
                    d_3667_v54_ = rhs506_
                    d_3633_v24_ = rhs507_
                    lhs282_[lhs283_] = rhs508_
                    lhs284_[lhs285_] = rhs509_
                    pass
            pass

    @property
    def f3(self):
        return self._f3
