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
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([773, len(_dafny.Map({True: 506}))]))])).Elements:
                d_0_v0_: int = compr_0_
                if (d_0_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([773, len(_dafny.Map({True: 506}))]))])):
                    coll0_ = coll0_.union(_dafny.Set([(d_0_v0_) + (357)]))
            return _dafny.Set(coll0_)
        return (iife0_()
        ) - ((_dafny.Set({766}) if True else _dafny.Set({(_dafny.MultiSet([412])).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1_i0_ in range(default__.abs(98))]))})))

    @staticmethod
    def fm1(globalState):
        return (((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D8_DC18(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khb"))))]))) | (_dafny.MultiSet([D8_DC18(-677)]))) | ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D8_DC18(-488)]))).intersection(_dafny.MultiSet([D8_DC18(len(_dafny.Set({_dafny.CodePoint('d')})))])))).cardinality

    @staticmethod
    def fm2(p0, p1, globalState):
        if True:
            return (_dafny.CodePoint('w')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhudtnhvn")))
        elif True:
            return (-773) != (-371)

    @staticmethod
    def fm4(p0, globalState):
        return 4

    @staticmethod
    def fm8(p0, globalState):
        source0_ = D7_DC15(_dafny.Map({(_dafny.MultiSet([False])).cardinality: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xtkckav"))}))
        if source0_.is_DC16:
            d_2___mcc_h0_ = source0_.cf26
            d_3___mcc_h1_ = source0_.cf27
            d_4___mcc_h2_ = source0_.cf28
            d_5___mcc_h3_ = source0_.cf29
            d_6_cf29_ = d_5___mcc_h3_
            d_7_cf28_ = d_4___mcc_h2_
            d_8_cf27_ = d_3___mcc_h1_
            d_9_cf26_ = d_2___mcc_h0_
            return D3_DC9(_dafny.Map({26: _dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_9_cf26_]))).cardinality, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wlkdfem")))), 610})}))
        elif True:
            d_10___mcc_h4_ = source0_.cf25
            d_11_cf25_ = d_10___mcc_h4_
            if False:
                def iife1_():
                    coll1_ = _dafny.Map()
                    compr_1_: int
                    for compr_1_ in _dafny.IntegerRange(974, 89):
                        d_12_v0_: int = compr_1_
                        if ((974) <= (d_12_v0_)) and ((d_12_v0_) < (89)):
                            coll1_[(d_12_v0_) * ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_13_i0_ in range(default__.abs(409))]))))] = _dafny.Set({74, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, True])), 799, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))]))})
                    return _dafny.Map(coll1_)
                return D3_DC9(iife1_()
)
            elif True:
                return D3_DC9(_dafny.Map({-232: _dafny.Set({822})}))

    @staticmethod
    def fm9(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njtrptc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))

    @staticmethod
    def fm10(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yn"))) < ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_14_i0_ in range(default__.abs(225))]) if False else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uegdnut"))))

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        if (_dafny.MultiSet([False])).isdisjoint(_dafny.MultiSet([False, True])):
            return (_dafny.Map({True: False})) | (_dafny.Map({True: False}))
        elif True:
            return _dafny.Map({False: True})

    @staticmethod
    def fm16(p0, globalState):
        return (default__.safeDivisionInt(-719, -755)) == (len((_dafny.Set({not(False)}) if False else _dafny.Set({True, not(False)}))))

    @staticmethod
    def fm17(p0, p1, p2, p3, globalState):
        def iife2_():
            def iife6_():
                def iife8_():
                    coll8_ = _dafny.Map()
                    compr_8_: str
                    for compr_8_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m')])).Elements:
                        d_16_v2_: str = compr_8_
                        if (d_16_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m')])):
                            coll8_[d_16_v2_] = (D9_DC23(True)).cf38
                    return _dafny.Map(coll8_)
                coll6_ = _dafny.Map()
                def iife7_():
                    coll7_ = _dafny.Map()
                    compr_7_: str
                    for compr_7_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m')])).Elements:
                        d_16_v2_: str = compr_7_
                        if (d_16_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m')])):
                            coll7_[d_16_v2_] = (D9_DC23(True)).cf38
                    return _dafny.Map(coll7_)
                compr_6_: str
                for compr_6_ in (iife7_()
                ).keys.Elements:
                    d_17_v1_: str = compr_6_
                    if (d_17_v1_) in (iife8_()
                    ):
                        coll6_[d_17_v1_] = len(_dafny.Set({-244, 369}))
                return _dafny.Map(coll6_)
            coll2_ = _dafny.Map()
            def iife3_():
                def iife5_():
                    coll5_ = _dafny.Map()
                    compr_5_: str
                    for compr_5_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m')])).Elements:
                        d_16_v2_: str = compr_5_
                        if (d_16_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m')])):
                            coll5_[d_16_v2_] = (D9_DC23(True)).cf38
                    return _dafny.Map(coll5_)
                coll3_ = _dafny.Map()
                def iife4_():
                    coll4_ = _dafny.Map()
                    compr_4_: str
                    for compr_4_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m')])).Elements:
                        d_16_v2_: str = compr_4_
                        if (d_16_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m')])):
                            coll4_[d_16_v2_] = (D9_DC23(True)).cf38
                    return _dafny.Map(coll4_)
                compr_3_: str
                for compr_3_ in (iife4_()
                ).keys.Elements:
                    d_17_v1_: str = compr_3_
                    if (d_17_v1_) in (iife5_()
                    ):
                        coll3_[d_17_v1_] = len(_dafny.Set({-244, 369}))
                return _dafny.Map(coll3_)
            compr_2_: str
            for compr_2_ in (iife3_()
            ).keys.Elements:
                d_18_v0_: str = compr_2_
                if (d_18_v0_) in (iife6_()
                ):
                    coll2_[d_18_v0_] = 866
            return _dafny.Map(coll2_)
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False})) for d_15_i0_ in range(default__.abs(122))]), _dafny.SeqWithoutIsStrInference([len(iife2_()
        ), len(_dafny.Map({567: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cpapaci"))))}))])])

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        return ((D0_DC2(_dafny.Set({len(_dafny.Map({_dafny.Set({739, -766}): D1_DC7(_dafny.CodePoint('q'), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")))}))}), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_19_i0_ in range(default__.abs(-981))]))).cf6) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvlldbjbx")))

    @staticmethod
    def fm21(p0, p1, globalState):
        def iife9_():
            def iife11_():
                coll11_ = _dafny.Set()
                compr_11_: int
                for compr_11_ in _dafny.IntegerRange(951, -986):
                    d_23_v0_: int = compr_11_
                    if ((951) <= (d_23_v0_)) and ((d_23_v0_) < (-986)):
                        coll11_ = coll11_.union(_dafny.Set([default__.safeDivisionInt(d_23_v0_, 28)]))
                return _dafny.Set(coll11_)
            coll9_ = _dafny.Set()
            def iife10_():
                coll10_ = _dafny.Set()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(951, -986):
                    d_21_v0_: int = compr_10_
                    if ((951) <= (d_21_v0_)) and ((d_21_v0_) < (-986)):
                        coll10_ = coll10_.union(_dafny.Set([default__.safeDivisionInt(d_21_v0_, 28)]))
                return _dafny.Set(coll10_)
            compr_9_: int
            for compr_9_ in (iife10_()
            ).Elements:
                d_22_v1_: int = compr_9_
                if (d_22_v1_) in (iife11_()
                ):
                    coll9_ = coll9_.union(_dafny.Set([default__.safeDivisionInt(d_22_v1_, -818)]))
            return _dafny.Set(coll9_)
        return D0_DC2((_dafny.Set({len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([True]))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_20_i0_ in range(default__.abs(-240))]))]))})) | (iife9_()
), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_24_i1_ in range(default__.abs(217))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oiuswc"))))

    @staticmethod
    def fm22(p0, p1, globalState):
        return (default__.safeModuloInt(len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxdj")): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypeb")))})), (0) - ((0) - (len(_dafny.Map({False: False})))))) == ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cn"))) if not(not(False)) else len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_25_i1_ in range(default__.abs(708))])) for d_26_i0_ in range(default__.abs(800))]))))

    @staticmethod
    def fm23(globalState):
        return D1_DC6(len(_dafny.Set({False})))

    @staticmethod
    def fm24(globalState):
        def iife12_():
            def iife14_():
                coll14_ = _dafny.Set()
                compr_14_: D1
                for compr_14_ in (_dafny.SeqWithoutIsStrInference([D1_DC6(len(_dafny.Map({415: False})))])).Elements:
                    d_30_v0_: D1 = compr_14_
                    if (d_30_v0_) in (_dafny.SeqWithoutIsStrInference([D1_DC6(len(_dafny.Map({415: False})))])):
                        coll14_ = coll14_.union(_dafny.Set([d_30_v0_]))
                return _dafny.Set(coll14_)
            coll12_ = _dafny.Set()
            def iife13_():
                coll13_ = _dafny.Set()
                compr_13_: D1
                for compr_13_ in (_dafny.SeqWithoutIsStrInference([D1_DC6(len(_dafny.Map({415: False})))])).Elements:
                    d_28_v0_: D1 = compr_13_
                    if (d_28_v0_) in (_dafny.SeqWithoutIsStrInference([D1_DC6(len(_dafny.Map({415: False})))])):
                        coll13_ = coll13_.union(_dafny.Set([d_28_v0_]))
                return _dafny.Set(coll13_)
            compr_12_: D1
            for compr_12_ in (iife13_()
            ).Elements:
                d_29_v1_: D1 = compr_12_
                if (d_29_v1_) in (iife14_()
                ):
                    coll12_ = coll12_.union(_dafny.Set([d_29_v1_]))
            return _dafny.Set(coll12_)
        return ((_dafny.Set({D1_DC6(-922)})) | (_dafny.Set({D1_DC6(-649), D1_DC6(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_27_i0_ in range(default__.abs(922))]))), D1_DC6((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([656, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjdg"))), len(_dafny.SeqWithoutIsStrInference([False, False]))]))).cardinality)}))).intersection(iife12_()
        )

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([not(True), True, True])) + ((_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([not(True), False])))

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        if (_dafny.SeqWithoutIsStrInference([False])) == (_dafny.SeqWithoutIsStrInference([False, not(False)])):
            return (_dafny.MultiSet([False])) | (_dafny.MultiSet([False, True, True, False]))
        elif True:
            return (_dafny.MultiSet([True])) | (_dafny.MultiSet([False, False]))

    @staticmethod
    def fm27(p0, p1, p2, globalState):
        def iife15_():
            coll15_ = _dafny.Map()
            compr_15_: int
            for compr_15_ in _dafny.IntegerRange(456, -349):
                d_31_v0_: int = compr_15_
                if ((456) <= (d_31_v0_)) and ((d_31_v0_) < (-349)):
                    coll15_[(d_31_v0_) + (131)] = D1_DC7(_dafny.CodePoint('e'), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gstwlmab")))
            return _dafny.Map(coll15_)
        return (iife15_()
        ) | (_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False, not(not(False))]))).cardinality: D1_DC7(_dafny.CodePoint('a'), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "leq")))}))

    @staticmethod
    def fm28(p0, globalState):
        return _dafny.Map({(-714) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_32_i0_ in range(default__.abs(-198))]))): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emwqfdps"))})

    @staticmethod
    def fm29(p0, p1, globalState):
        return ((_dafny.Map({True: False})) | (_dafny.Map({False: True}))) | ((_dafny.Map({False: False})) | (_dafny.Map({True: False})))

    @staticmethod
    def fm30(p0, p1, p2, globalState):
        source1_ = D1_DC6(135)
        if source1_.is_DC6:
            d_33___mcc_h0_ = source1_.cf12
            d_34_cf12_ = d_33___mcc_h0_
            return _dafny.Map({271: True})
        elif source1_.is_DC7:
            d_35___mcc_h1_ = source1_.cf13
            d_36___mcc_h2_ = source1_.cf14
            d_37_cf14_ = d_36___mcc_h2_
            d_38_cf13_ = d_35___mcc_h1_
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(556, 947):
                    d_39_v0_: int = compr_16_
                    if ((556) <= (d_39_v0_)) and ((d_39_v0_) < (947)):
                        coll16_[(d_39_v0_) + (-346)] = True
                return _dafny.Map(coll16_)
            return (iife16_()
            ) | (_dafny.Map({-892: True}))
        elif True:
            d_40___mcc_h3_ = source1_.cf11
            d_41_cf11_ = d_40___mcc_h3_
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in _dafny.IntegerRange(-751, 772):
                    d_42_v1_: int = compr_17_
                    if ((-751) <= (d_42_v1_)) and ((d_42_v1_) < (772)):
                        coll17_[(d_42_v1_) - (len(_dafny.Map({_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([False])).cardinality)]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwcp"))})))] = False
                return _dafny.Map(coll17_)
            return (iife17_()
            ) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkb"))): False}))

    @staticmethod
    def fm31(globalState):
        return ((_dafny.Set({False})).intersection(_dafny.Set({not(not(True)), True, True, False}))) | ((_dafny.Set({True}) if not(True) else _dafny.Set({True})))

    @staticmethod
    def fm32(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxmfcp"))

    @staticmethod
    def fm33(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kyandbjqq"))), len(_dafny.SeqWithoutIsStrInference([True]))]) for d_43_i0_ in range(default__.abs(840))])

    @staticmethod
    def fm36(globalState):
        return -592

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        return (_dafny.Map({not((D0_DC3(330, True, False)).cf8): False})) | ((_dafny.Map({False: True})) | (_dafny.Map({not(True): False})))

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({False})) - (_dafny.Set({True}))) - (_dafny.Set({not(False)}))

    @staticmethod
    def fm39(globalState):
        source2_ = D11_DC28()
        if source2_.is_DC28:
            return _dafny.SeqWithoutIsStrInference([D7_DC15(_dafny.Map({995: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_44_i1_ in range(default__.abs(89))])})) for d_45_i0_ in range(default__.abs(571))])
        elif True:
            d_46___mcc_h0_ = source2_.cf42
            d_47_cf42_ = d_46___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([D7_DC15(_dafny.Map({881: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqaiipv"))})) for d_48_i2_ in range(default__.abs(560))])) + (_dafny.SeqWithoutIsStrInference([D7_DC15(_dafny.Map({len(_dafny.Set({True})): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnwkwu"))})), D7_DC15(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_49_i3_ in range(default__.abs(-651))])): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_50_i4_ in range(default__.abs(491))])})), D7_DC15(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qdadesapq")))): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdilqo"))})), D7_DC15(_dafny.Map({(0) - (len(_dafny.Set({False, False, True}))): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmgqbmpj"))}))]))

    @staticmethod
    def fm40(globalState):
        return not(((0) - ((_dafny.MultiSet([(0) - (len((D14_DC31(_dafny.SeqWithoutIsStrInference([True]))).cf45))])).cardinality)) not in (_dafny.SeqWithoutIsStrInference([22])))

    @staticmethod
    def fm41(p0, globalState):
        return D11_DC28()

    @staticmethod
    def fm42(p0, globalState):
        return _dafny.CodePoint('q')

    @staticmethod
    def fm43(p0, p1, p2, p3, globalState):
        def iife18_():
            coll18_ = _dafny.Set()
            compr_18_: _dafny.Map
            for compr_18_ in (_dafny.Map({_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False])): False}): False})).keys.Elements:
                d_51_v0_: _dafny.Map = compr_18_
                if (d_51_v0_) in (_dafny.Map({_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False])): False}): False})):
                    coll18_ = coll18_.union(_dafny.Set([d_51_v0_]))
            return _dafny.Set(coll18_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(iife18_()
        ), -731]))]): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_52_i0_ in range(default__.abs(87))]))})) | ((D15_DC34(_dafny.Map({_dafny.SeqWithoutIsStrInference([923, len(_dafny.SeqWithoutIsStrInference([False]))]): 496}))).cf48)) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e')])))]): 417})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: False}))]): 310})))

    @staticmethod
    def fm44(globalState):
        return D0_DC1()

    @staticmethod
    def fm45(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([_dafny.CodePoint('c')])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p')])))) | (((D16_DC36(_dafny.MultiSet([_dafny.CodePoint('d')]))).cf50 if not(not(True)) else _dafny.MultiSet([_dafny.CodePoint('i')])))

    @staticmethod
    def fm46(p0, globalState):
        def iife19_():
            coll19_ = _dafny.Map()
            compr_19_: _dafny.Seq
            for compr_19_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_54_i2_ in range(default__.abs(-971))]))) for d_55_i1_ in range(default__.abs(270))])])).Elements:
                d_56_v0_: _dafny.Seq = compr_19_
                if (d_56_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_54_i2_ in range(default__.abs(-971))]))) for d_55_i1_ in range(default__.abs(270))])])):
                    coll19_[d_56_v0_] = 460
            return _dafny.Map(coll19_)
        return ((_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvjprk")))): 800})) | (_dafny.Map({-433: 576}))) | (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_53_i0_ in range(default__.abs(-714))])): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality}))]))): len(iife19_()
        )}))

    @staticmethod
    def fm47(p0, p1, p2, globalState):
        return D1_DC5(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([315, len((_dafny.Map({False: not(False)}))), 228, 809, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywclxcu")))])]))

    @staticmethod
    def fm48(p0, p1, p2, p3, globalState):
        if (len(_dafny.Map({True: (0) - (-356)}))) < (90):
            return D7_DC15(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))]))): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khcfpyqr"))}))
        elif True:
            return D7_DC15(_dafny.Map({len((D18_DC40(_dafny.Map({False: 699}))).cf57): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqprsunx"))}))

    @staticmethod
    def fm49(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False, True, False]) for d_57_i0_ in range(default__.abs(-298))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_58_i1_ in range(default__.abs(251))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])]))

    @staticmethod
    def fm50(p0, p1, globalState):
        return (_dafny.Map({not(not(not(not(not(False))))): 615})) | (_dafny.Map({False: -21}))

    @staticmethod
    def fm51(p0, globalState):
        source3_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))])
        d_59___mcc_h0_ = source3_
        d_60_cf19_ = d_59___mcc_h0_
        return (_dafny.Map({False: _dafny.CodePoint('r')})) | (_dafny.Map({False: _dafny.CodePoint('y')}))

    @staticmethod
    def fm52(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])), _dafny.MultiSet([False, False]), _dafny.MultiSet([True, True]), _dafny.MultiSet([False])]) if not(False) else _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True, False])]))

    @staticmethod
    def fm53(p0, p1, globalState):
        def iife20_():
            coll20_ = _dafny.Map()
            compr_20_: _dafny.MultiSet
            for compr_20_ in (_dafny.Map({_dafny.MultiSet([False, not(True)]): (_dafny.MultiSet([-591])).cardinality})).keys.Elements:
                d_61_v0_: _dafny.MultiSet = compr_20_
                if (d_61_v0_) in (_dafny.Map({_dafny.MultiSet([False, not(True)]): (_dafny.MultiSet([-591])).cardinality})):
                    coll20_[d_61_v0_] = False
            return _dafny.Map(coll20_)
        return ((_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True, True, False])): not(True)})) | (iife20_()
        )) | ((_dafny.Map({_dafny.MultiSet([False, False]): False})) | (_dafny.Map({_dafny.MultiSet([True]): True})))

    @staticmethod
    def fm54(p0, p1, p2, p3, globalState):
        def iife21_():
            def iife23_():
                coll23_ = _dafny.Map()
                compr_23_: int
                for compr_23_ in _dafny.IntegerRange(313, -556):
                    d_62_v1_: int = compr_23_
                    if ((313) <= (d_62_v1_)) and ((d_62_v1_) < (-556)):
                        coll23_[(d_62_v1_) * (-215)] = False
                return _dafny.Map(coll23_)
            coll21_ = _dafny.Map()
            def iife22_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in _dafny.IntegerRange(313, -556):
                    d_62_v1_: int = compr_22_
                    if ((313) <= (d_62_v1_)) and ((d_62_v1_) < (-556)):
                        coll22_[(d_62_v1_) * (-215)] = False
                return _dafny.Map(coll22_)
            compr_21_: int
            for compr_21_ in (iife22_()
            ).keys.Elements:
                d_63_v0_: int = compr_21_
                if (d_63_v0_) in (iife23_()
                ):
                    coll21_[default__.safeModuloInt(d_63_v0_, len(_dafny.Map({True: -921})))] = False
            return _dafny.Map(coll21_)
        source4_ = (D7_DC16(True, _dafny.Map({330: not(not(True))}), _dafny.CodePoint('h'), True) if False else D7_DC16(False, iife21_()
, _dafny.CodePoint('r'), True))
        if source4_.is_DC16:
            d_64___mcc_h0_ = source4_.cf26
            d_65___mcc_h1_ = source4_.cf27
            d_66___mcc_h2_ = source4_.cf28
            d_67___mcc_h3_ = source4_.cf29
            d_68_cf29_ = d_67___mcc_h3_
            d_69_cf28_ = d_66___mcc_h2_
            d_70_cf27_ = d_65___mcc_h1_
            d_71_cf26_ = d_64___mcc_h0_
            def iife24_():
                coll24_ = _dafny.Map()
                compr_24_: int
                for compr_24_ in _dafny.IntegerRange(259, 610):
                    d_72_v2_: int = compr_24_
                    if ((259) <= (d_72_v2_)) and ((d_72_v2_) < (610)):
                        coll24_[(d_72_v2_) + (216)] = D8_DC20(D8_DC18(len(_dafny.Map({-865: d_71_cf26_}))))
                return _dafny.Map(coll24_)
            return (_dafny.Map({_dafny.Map({(0) - (-156): D8_DC20(D8_DC19())}): _dafny.Set({d_69_cf28_})})) | (_dafny.Map({iife24_()
            : _dafny.Set({d_69_cf28_})}))
        elif True:
            d_73___mcc_h4_ = source4_.cf25
            d_74_cf25_ = d_73___mcc_h4_
            return _dafny.Map({_dafny.Map({35: D8_DC20(D8_DC20(D8_DC19()))}): _dafny.Set({_dafny.CodePoint('d')})})

    @staticmethod
    def fm55(p0, globalState):
        source5_ = D16_DC36(_dafny.MultiSet([_dafny.CodePoint('j'), _dafny.CodePoint('m')]))
        if source5_.is_DC37:
            d_75___mcc_h0_ = source5_.cf51
            d_76___mcc_h1_ = source5_.cf52
            d_77_cf52_ = d_76___mcc_h1_
            d_78_cf51_ = d_75___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([616])
        elif source5_.is_DC38:
            d_79___mcc_h2_ = source5_.cf53
            d_80___mcc_h3_ = source5_.cf54
            d_81___mcc_h4_ = source5_.cf55
            d_82_cf55_ = d_81___mcc_h4_
            d_83_cf54_ = d_80___mcc_h3_
            d_84_cf53_ = d_79___mcc_h2_
            return (_dafny.SeqWithoutIsStrInference([d_82_cf55_])) + (_dafny.SeqWithoutIsStrInference([d_82_cf55_ for d_85_i0_ in range(default__.abs(-104))]))
        elif True:
            d_86___mcc_h5_ = source5_.cf50
            d_87_cf50_ = d_86___mcc_h5_
            return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_88_i1_ in range(default__.abs(-5))])

    @staticmethod
    def fm56(p0, p1, globalState):
        def iife25_():
            coll25_ = _dafny.Map()
            compr_25_: _dafny.Set
            for compr_25_ in ((_dafny.Map({_dafny.Set({True}): False})) | (_dafny.Map({_dafny.Set({False, not(False)}): True}))).keys.Elements:
                d_89_v0_: _dafny.Set = compr_25_
                if (d_89_v0_) in ((_dafny.Map({_dafny.Set({True}): False})) | (_dafny.Map({_dafny.Set({False, not(False)}): True}))):
                    coll25_[d_89_v0_] = (D3_DC10(True, False)).cf17
            return _dafny.Map(coll25_)
        return iife25_()
        

    @staticmethod
    def m0(globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: int = int(0)
        d_90_v0_: int
        d_90_v0_ = 805
        d_91_v1_: bool
        d_91_v1_ = False
        d_92_v2_: _dafny.MultiSet
        d_92_v2_ = _dafny.MultiSet([d_91_v1_, d_91_v1_])
        d_93_v3_: D10
        d_93_v3_ = D10_DC25(d_92_v2_)
        d_94_v4_: _dafny.Seq
        d_94_v4_ = _dafny.SeqWithoutIsStrInference([d_93_v3_])
        rhs0_ = d_90_v0_
        rhs1_ = d_90_v0_
        rhs2_ = (d_90_v0_) == (len(d_94_v4_))
        lhs0_ = globalState
        lhs1_ = globalState
        r1 = rhs0_
        lhs0_.f4 = rhs1_
        lhs1_.f6 = rhs2_
        d_95_v5_: C4
        nw0_ = C4()
        nw0_.ctor__()
        d_95_v5_ = nw0_
        d_96_v6_: _dafny.Map
        d_96_v6_ = _dafny.Map({d_91_v1_: True})
        d_97_v7_: _dafny.MultiSet
        d_97_v7_ = _dafny.MultiSet([d_90_v0_])
        d_98_v8_: _dafny.MultiSet
        d_98_v8_ = _dafny.MultiSet([(d_97_v7_).cardinality])
        d_99_v9_: _dafny.Seq
        d_99_v9_ = _dafny.SeqWithoutIsStrInference([(d_98_v8_).cardinality, 32, d_90_v0_])
        d_100_v10_: _dafny.Seq
        d_100_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxtwt"))
        d_101_v11_: _dafny.Array
        nw1_ = _dafny.Array(None, 6)
        nw1_[int(0)] = len(d_99_v9_)
        nw1_[int(1)] = ((d_97_v7_)[d_90_v0_] if (d_90_v0_) in (d_97_v7_) else d_90_v0_)
        nw1_[int(2)] = d_90_v0_
        nw1_[int(3)] = len(d_100_v10_)
        nw1_[int(4)] = d_90_v0_
        nw1_[int(5)] = d_90_v0_
        d_101_v11_ = nw1_
        d_102_v12_: _dafny.Array
        def lambda0_(d_103_v1_):
            def lambda1_(d_104_i2_):
                return d_103_v1_

            return lambda1_

        init0_ = lambda0_(d_91_v1_)
        nw2_ = _dafny.Array(None, 5)
        for i0_0_ in range(nw2_.length(0)):
            nw2_[i0_0_] = init0_(i0_0_)
        d_102_v12_ = nw2_
        d_105_v13_: _dafny.Array
        nw3_ = _dafny.Array(None, 9)
        nw3_[int(0)] = d_102_v12_
        nw3_[int(1)] = d_102_v12_
        nw3_[int(2)] = d_102_v12_
        nw3_[int(3)] = d_102_v12_
        nw3_[int(4)] = d_102_v12_
        nw3_[int(5)] = d_102_v12_
        nw3_[int(6)] = d_102_v12_
        nw3_[int(7)] = d_102_v12_
        nw3_[int(8)] = d_102_v12_
        d_105_v13_ = nw3_
        d_106_v14_: D9
        d_106_v14_ = D9_DC22(d_101_v11_, d_91_v1_, d_105_v13_, d_100_v10_)
        d_107_v15_: _dafny.Seq
        d_107_v15_ = _dafny.SeqWithoutIsStrInference([default__.fm16(default__.fm11((d_106_v14_).cf35, d_91_v1_, d_91_v1_, globalState), globalState), not(not(True)), d_91_v1_])
        d_108_v16_: _dafny.Array
        nw4_ = _dafny.Array(None, 18)
        nw4_[int(0)] = (default__.fm16(d_96_v6_, globalState)) or (d_91_v1_)
        nw4_[int(1)] = d_91_v1_
        nw4_[int(2)] = not(d_91_v1_)
        nw4_[int(3)] = d_91_v1_
        nw4_[int(4)] = d_91_v1_
        nw4_[int(5)] = (_dafny.SeqWithoutIsStrInference([d_90_v0_ for d_109_i1_ in range(default__.abs(563))])) == (d_99_v9_)
        nw4_[int(6)] = False
        nw4_[int(7)] = (d_107_v15_)[default__.safeIndex(d_90_v0_, len(d_107_v15_))]
        nw4_[int(8)] = (-904) < ((d_98_v8_).cardinality)
        nw4_[int(9)] = d_91_v1_
        nw4_[int(10)] = d_91_v1_
        nw4_[int(11)] = ((d_98_v8_).set(d_90_v0_, default__.abs(len(_dafny.SeqWithoutIsStrInference([d_90_v0_ for d_110_i4_ in range(default__.abs(29))]))))).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({d_90_v0_: d_90_v0_}))) for d_111_i3_ in range(default__.abs(948))])))
        nw4_[int(12)] = True
        nw4_[int(13)] = d_91_v1_
        nw4_[int(14)] = d_91_v1_
        nw4_[int(15)] = d_91_v1_
        nw4_[int(16)] = (d_107_v15_) < (d_107_v15_)
        nw4_[int(17)] = (d_91_v1_ if d_91_v1_ else default__.fm22(False, d_90_v0_, globalState))
        d_108_v16_ = nw4_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_108_v16_).length(0)):
            d_112_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_112_i0_)) and ((d_112_i0_) < ((d_108_v16_).length(0)))):
                (d_108_v16_)[(d_112_i0_)] = d_91_v1_
        d_113_v17_: T1
        nw5_ = C7()
        nw5_.ctor__(d_108_v16_, d_91_v1_, d_90_v0_)
        d_113_v17_ = nw5_
        d_114_v18_: _dafny.Map
        d_114_v18_ = _dafny.Map({len(d_100_v10_): d_113_v17_})
        d_115_v19_: _dafny.MultiSet
        d_115_v19_ = _dafny.MultiSet([d_114_v18_])
        if ((_dafny.MultiSet([(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_113_v17_.f29])): d_113_v17_})).set(d_90_v0_, d_113_v17_)])) - (d_115_v19_)).issubset(_dafny.MultiSet([d_114_v18_, d_114_v18_])):
            index0_ = default__.safeIndex(582, (d_101_v11_).length(0))
            (d_101_v11_)[index0_] = d_113_v17_.f29
            index1_ = default__.safeIndex(582, (d_101_v11_).length(0))
            (d_101_v11_)[index1_] = default__.safeModuloInt(-587, d_113_v17_.f29)
            d_116_v20_: str
            d_116_v20_ = _dafny.CodePoint('p')
            (globalState).f25 = d_116_v20_
            d_90_v0_ = (0) - ((0) - (default__.safeDivisionInt(-565, (default__.fm1(globalState) if (d_113_v17_).f28 else (d_101_v11_)[default__.safeIndex(582, (d_101_v11_).length(0))]))))
            (globalState).f6 = (d_113_v17_).f28
            (globalState).f17 = (d_100_v10_) == (d_100_v10_)
        elif True:
            index2_ = default__.safeIndex(731, (d_101_v11_).length(0))
            (d_101_v11_)[index2_] = d_113_v17_.f29
            index3_ = default__.safeIndex(731, (d_101_v11_).length(0))
            (d_101_v11_)[index3_] = (0) - (default__.safeModuloInt(d_113_v17_.f29, (d_113_v17_.f29) - (d_90_v0_)))
            d_90_v0_ = ((d_101_v11_)[default__.safeIndex(731, (d_101_v11_).length(0))]) - (len(_dafny.Map({(d_113_v17_).f28: d_96_v6_})))
            d_117_v21_: str
            d_117_v21_ = _dafny.CodePoint('u')
            d_118_v22_: C8
            nw6_ = C8()
            nw6_.ctor__(d_117_v21_, d_108_v16_, (d_113_v17_).f28)
            d_118_v22_ = nw6_
            d_119_v23_: D1
            d_119_v23_ = D1_DC6((d_101_v11_)[default__.safeIndex(731, (d_101_v11_).length(0))])
            d_120_v24_: _dafny.MultiSet
            d_120_v24_ = _dafny.MultiSet([d_119_v23_, d_119_v23_, d_119_v23_, d_119_v23_])
            index4_ = default__.safeIndex(731, (d_101_v11_).length(0))
            rhs3_ = d_90_v0_
            rhs4_ = d_118_v22_
            rhs5_ = (default__.fm4(d_120_v24_, globalState)) >= (d_113_v17_.f29)
            lhs2_ = d_101_v11_
            lhs3_ = default__.safeIndex(731, (d_101_v11_).length(0))
            lhs4_ = globalState
            lhs2_[lhs3_] = rhs3_
            d_118_v22_ = rhs4_
            lhs4_.f17 = rhs5_
            d_121_v25_: C0
            nw7_ = C0()
            nw7_.ctor__(d_100_v10_, (d_101_v11_)[default__.safeIndex(731, (d_101_v11_).length(0))], d_108_v16_, (d_113_v17_).f28)
            d_121_v25_ = nw7_
            d_121_v25_ = d_121_v25_
            d_122_v26_: _dafny.Set
            d_122_v26_ = _dafny.Set({(d_118_v22_).f32, default__.fm42(len(d_107_v15_), globalState)})
            (globalState).f4 = ((d_101_v11_)[default__.safeIndex(731, (d_101_v11_).length(0))]) * (default__.safeModuloInt(d_113_v17_.f29, len(d_122_v26_)))
        (globalState).f6 = (D5_DC13(True, d_101_v11_, False)).cf23
        (globalState).f4 = len(_dafny.SeqWithoutIsStrInference([_dafny.Set({d_113_v17_.f29, d_113_v17_.f29}) for d_123_i5_ in range(default__.abs(291))]))
        d_124_v27_: _dafny.Map
        d_124_v27_ = _dafny.Map({(0) - (d_113_v17_.f29): (d_113_v17_).f27})
        r0 = d_124_v27_
        r1 = 820
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_125_v0_: _dafny.Array
        def lambda2_(d_126_i0_):
            return False

        init1_ = lambda2_
        nw8_ = _dafny.Array(None, 27)
        for i0_1_ in range(nw8_.length(0)):
            nw8_[i0_1_] = init1_(i0_1_)
        d_125_v0_ = nw8_
        d_127_v1_: int
        d_127_v1_ = 900
        d_128_v2_: bool
        d_128_v2_ = False
        d_129_v3_: _dafny.MultiSet
        d_129_v3_ = _dafny.MultiSet([d_128_v2_])
        d_130_v4_: _dafny.Map
        d_130_v4_ = _dafny.Map({(0) - (d_127_v1_): d_129_v3_})
        d_131_v5_: _dafny.Array
        nw9_ = _dafny.Array(_dafny.Array(None, 0), 18)
        d_131_v5_ = nw9_
        d_132_v6_: _dafny.Map
        d_132_v6_ = _dafny.Map({d_129_v3_: d_131_v5_})
        d_133_v7_: _dafny.Array
        nw10_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
        d_133_v7_ = nw10_
        d_134_v8_: _dafny.Array
        def lambda3_(d_135_v1_, d_136_v2_):
            def lambda4_(d_137_i1_):
                return default__.safeModuloInt(d_137_i1_, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(_dafny.MultiSet([d_135_v1_, d_135_v1_])).cardinality: d_136_v2_})])))

            return lambda4_

        init2_ = lambda3_(d_127_v1_, d_128_v2_)
        nw11_ = _dafny.Array(None, 3)
        for i0_2_ in range(nw11_.length(0)):
            nw11_[i0_2_] = init2_(i0_2_)
        d_134_v8_ = nw11_
        d_138_v9_: _dafny.Seq
        d_138_v9_ = _dafny.SeqWithoutIsStrInference([d_134_v8_, d_134_v8_])
        d_139_v10_: _dafny.Map
        d_139_v10_ = _dafny.Map({d_128_v2_: d_131_v5_})
        d_140_v11_: _dafny.Seq
        d_140_v11_ = _dafny.SeqWithoutIsStrInference([d_128_v2_, d_128_v2_])
        d_141_v12_: _dafny.Seq
        d_141_v12_ = _dafny.SeqWithoutIsStrInference([d_127_v1_])
        d_142_v13_: _dafny.Seq
        d_142_v13_ = _dafny.SeqWithoutIsStrInference([d_131_v5_, ((d_139_v10_)[(d_140_v11_)[default__.safeIndex(len(d_141_v12_), len(d_140_v11_))]] if ((d_140_v11_)[default__.safeIndex(len(d_141_v12_), len(d_140_v11_))]) in (d_139_v10_) else d_131_v5_)])
        d_143_v14_: _dafny.Seq
        d_143_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prvhys"))
        d_144_v15_: _dafny.Seq
        d_144_v15_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")), d_143_v14_])
        d_145_v16_: _dafny.Map
        d_145_v16_ = _dafny.Map({76: d_134_v8_})
        d_146_globalState_: GlobalState
        nw12_ = GlobalState()
        nw12_.ctor__(False, True, 888, True, 277, 131, True, 306, d_125_v0_, d_130_v4_, _dafny.MultiSet([d_127_v1_]), ((d_132_v6_)[d_129_v3_] if (d_129_v3_) in (d_132_v6_) else d_131_v5_), 321, 529, d_133_v7_, False, (d_138_v9_) + (d_138_v9_), False, (d_142_v13_)[default__.safeIndex(d_127_v1_, len(d_142_v13_))], (d_144_v15_) + (d_144_v15_), 778, 481, d_145_v16_, 726, 300, _dafny.CodePoint('u'), 792)
        d_146_globalState_ = nw12_
        d_147_v17_: _dafny.Map
        d_147_v17_ = _dafny.Map({d_128_v2_: d_127_v1_})
        d_148_v18_: _dafny.MultiSet
        d_148_v18_ = _dafny.MultiSet([d_147_v17_])
        if (d_148_v18_).issubset(_dafny.MultiSet([d_147_v17_])):
            d_149_v19_: _dafny.Map
            d_150_v20_: int
            out0_: _dafny.Map
            out1_: int
            out0_, out1_ = default__.m0(d_146_globalState_)
            d_149_v19_ = out0_
            d_150_v20_ = out1_
            d_151_v21_: _dafny.MultiSet
            d_151_v21_ = _dafny.MultiSet([(0) - (((d_129_v3_)[d_128_v2_] if (d_128_v2_) in (d_129_v3_) else d_150_v20_)), len(d_143_v14_)])
            index5_ = default__.safeIndex(5, (d_125_v0_).length(0))
            (d_125_v0_)[index5_] = d_128_v2_
            index6_ = default__.safeIndex(900, (d_125_v0_).length(0))
            (d_125_v0_)[index6_] = d_128_v2_
            d_152_v22_: _dafny.Map
            d_152_v22_ = _dafny.Map({not (d_128_v2_) or (d_128_v2_): d_128_v2_})
            d_153_v23_: _dafny.Set
            d_153_v23_ = _dafny.Set({d_127_v1_, d_127_v1_, (0) - (d_127_v1_), (d_129_v3_).cardinality})
            d_154_v24_: _dafny.Seq
            d_154_v24_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_150_v20_: d_153_v23_})])
            d_155_v25_: _dafny.Map
            d_155_v25_ = _dafny.Map({d_125_v0_: True})
            index7_ = default__.safeIndex(5, (d_125_v0_).length(0))
            index8_ = default__.safeIndex(900, (d_125_v0_).length(0))
            rhs6_ = d_151_v21_
            rhs7_ = d_128_v2_
            rhs8_ = len(default__.fm0((d_154_v24_)[default__.safeIndex(d_150_v20_, len(d_154_v24_))], ((d_155_v25_)[d_125_v0_] if (d_125_v0_) in (d_155_v25_) else d_128_v2_), d_128_v2_, d_146_globalState_))
            rhs9_ = (len((d_141_v12_) + (d_141_v12_))) <= ((d_150_v20_) - (len(d_143_v14_)))
            rhs10_ = d_152_v22_
            lhs5_ = d_125_v0_
            lhs6_ = default__.safeIndex(5, (d_125_v0_).length(0))
            lhs7_ = d_125_v0_
            lhs8_ = default__.safeIndex(900, (d_125_v0_).length(0))
            d_151_v21_ = rhs6_
            lhs5_[lhs6_] = rhs7_
            d_127_v1_ = rhs8_
            lhs7_[lhs8_] = rhs9_
            d_152_v22_ = rhs10_
            d_156_v26_: C4
            nw13_ = C4()
            nw13_.ctor__()
            d_156_v26_ = nw13_
            d_143_v14_ = ((d_143_v14_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_157_i2_ in range(default__.abs(751))]))) + (d_143_v14_)
            d_147_v17_ = (d_147_v17_).set(d_128_v2_, d_150_v20_)
        elif True:
            d_143_v14_ = d_143_v14_
            d_158_v27_: _dafny.Array
            nw14_ = _dafny.Array(_dafny.Map({}), 20)
            d_158_v27_ = nw14_
            d_158_v27_ = d_158_v27_
            d_159_v28_: _dafny.Map
            d_159_v28_ = _dafny.Map({d_134_v8_: d_143_v14_})
            d_159_v28_ = d_159_v28_
            rhs11_ = d_129_v3_
            rhs12_ = d_128_v2_
            lhs9_ = d_146_globalState_
            d_129_v3_ = rhs11_
            lhs9_.f6 = rhs12_
            d_160_v29_: _dafny.Set
            d_160_v29_ = _dafny.Set({((d_129_v3_)[d_128_v2_] if (d_128_v2_) in (d_129_v3_) else d_127_v1_)})
            def iife26_():
                coll26_ = _dafny.Set()
                compr_26_: int
                for compr_26_ in _dafny.IntegerRange(667, 735):
                    d_161_v30_: int = compr_26_
                    if ((667) <= (d_161_v30_)) and ((d_161_v30_) < (735)):
                        coll26_ = coll26_.union(_dafny.Set([default__.safeModuloInt(d_161_v30_, d_127_v1_)]))
                return _dafny.Set(coll26_)
            d_160_v29_ = ((iife26_()
            ) | (d_160_v29_) if d_128_v2_ else (d_160_v29_ if (d_140_v11_)[default__.safeIndex(d_127_v1_, len(d_140_v11_))] else d_160_v29_))
        d_162_v31_: _dafny.Set
        d_162_v31_ = _dafny.Set({True, False})
        d_163_v32_: _dafny.Seq
        d_163_v32_ = _dafny.SeqWithoutIsStrInference([d_162_v31_])
        d_164_v33_: _dafny.Seq
        d_164_v33_ = _dafny.SeqWithoutIsStrInference([D5_DC12(d_163_v32_)])
        d_165_v34_: _dafny.Seq
        d_165_v34_ = _dafny.SeqWithoutIsStrInference([d_164_v33_])
        d_166_v35_: _dafny.Seq
        d_166_v35_ = d_165_v34_
        d_167_v36_: C6
        nw15_ = C6()
        nw15_.ctor__(609, (d_125_v0_ if d_128_v2_ else d_125_v0_), (d_164_v33_) in ((d_166_v35_)))
        d_167_v36_ = nw15_
        if d_128_v2_:
            (d_146_globalState_).f4 = d_127_v1_
            d_168_v37_: _dafny.Set
            d_168_v37_ = _dafny.Set({(0) - (-185)})
            rhs13_ = (d_168_v37_).intersection(d_168_v37_)
            rhs14_ = d_125_v0_
            d_168_v37_ = rhs13_
            d_125_v0_ = rhs14_
            d_147_v17_ = (d_147_v17_).set(d_128_v2_, d_127_v1_)
            (d_146_globalState_).f13 = d_127_v1_
            (d_146_globalState_).f5 = d_127_v1_
        elif True:
            d_169_v38_: _dafny.Set
            d_169_v38_ = _dafny.Set({d_127_v1_})
            d_170_v39_: D0
            d_170_v39_ = D0_DC2(d_169_v38_, d_143_v14_)
            d_171_v40_: _dafny.Map
            d_171_v40_ = _dafny.Map({d_127_v1_: d_127_v1_})
            d_172_v41_: _dafny.MultiSet
            d_172_v41_ = _dafny.MultiSet([d_127_v1_, len((d_170_v39_).cf5), ((d_141_v12_)[default__.safeIndex(d_127_v1_, len(d_141_v12_))]) - (len(d_171_v40_)), d_127_v1_])
            index9_ = default__.safeIndex(554, (d_134_v8_).length(0))
            (d_134_v8_)[index9_] = (d_127_v1_) + ((0) - (d_127_v1_))
            d_173_v42_: _dafny.Map
            d_173_v42_ = _dafny.Map({d_127_v1_: d_169_v38_})
            d_174_v43_: _dafny.Map
            d_174_v43_ = _dafny.Map({d_127_v1_: d_140_v11_})
            pat_let_tv0_ = d_127_v1_
            pat_let_tv1_ = d_127_v1_
            pat_let_tv2_ = d_127_v1_
            pat_let_tv3_ = d_173_v42_
            pat_let_tv4_ = d_128_v2_
            pat_let_tv5_ = d_128_v2_
            pat_let_tv6_ = d_146_globalState_
            index10_ = default__.safeIndex(554, (d_134_v8_).length(0))
            def iife28_(_pat_let1_0):
                def iife29_(d_175_dt__update__tmp_h0_):
                    def iife30_(_pat_let2_0):
                        def iife31_(d_176_dt__update_hcf5_h0_):
                            return D0_DC2(d_176_dt__update_hcf5_h0_, (d_175_dt__update__tmp_h0_).cf6)
                        return iife31_(_pat_let2_0)
                    return iife30_(_dafny.Set({pat_let_tv0_, (0) - (pat_let_tv1_), pat_let_tv2_}))
                return iife29_(_pat_let1_0)
            def iife27_(_pat_let0_0):
                def iife32_(d_177_dt__update__tmp_h1_):
                    def iife33_(_pat_let3_0):
                        def iife34_(d_178_dt__update_hcf5_h1_):
                            return D0_DC2(d_178_dt__update_hcf5_h1_, (d_177_dt__update__tmp_h1_).cf6)
                        return iife34_(_pat_let3_0)
                    return iife33_(default__.fm0(pat_let_tv3_, pat_let_tv4_, pat_let_tv5_, pat_let_tv6_))
                return iife32_(_pat_let0_0)
            rhs15_ = iife27_(iife28_(d_170_v39_))
            rhs16_ = d_172_v41_
            rhs17_ = (0) - ((default__.safeModuloInt(len(((d_174_v43_)[d_127_v1_] if (d_127_v1_) in (d_174_v43_) else _dafny.SeqWithoutIsStrInference([d_128_v2_]))), len(d_140_v11_))) * ((len(d_162_v31_)) + (d_127_v1_)))
            rhs18_ = d_128_v2_
            lhs10_ = d_134_v8_
            lhs11_ = default__.safeIndex(554, (d_134_v8_).length(0))
            lhs12_ = d_146_globalState_
            d_170_v39_ = rhs15_
            d_172_v41_ = rhs16_
            lhs10_[lhs11_] = rhs17_
            lhs12_.f15 = rhs18_
            d_179_v44_: _dafny.Array
            nw16_ = _dafny.Array(_dafny.Map({}), 6)
            d_179_v44_ = nw16_
            d_179_v44_ = d_179_v44_
            index11_ = default__.safeIndex(554, (d_134_v8_).length(0))
            (d_134_v8_)[index11_] = default__.safeModuloInt(d_127_v1_, (d_134_v8_)[default__.safeIndex(554, (d_134_v8_).length(0))])
            d_180_v45_: D8
            d_180_v45_ = D8_DC19()
            d_181_v46_: str
            d_181_v46_ = _dafny.CodePoint('a')
            d_182_v47_: _dafny.Set
            d_182_v47_ = _dafny.Set({d_181_v46_, d_181_v46_, _dafny.CodePoint('p'), d_181_v46_, d_181_v46_})
            d_183_v48_: _dafny.Map
            d_183_v48_ = _dafny.Map({_dafny.Map({d_127_v1_: D8_DC20(d_180_v45_)}): (d_182_v47_ if True else d_182_v47_)})
            d_184_v49_: _dafny.Map
            d_184_v49_ = _dafny.Map({d_128_v2_: d_162_v31_})
            rhs19_ = (d_128_v2_) in ((d_184_v49_ if d_128_v2_ else d_184_v49_))
            rhs20_ = default__.fm54((d_127_v1_) < ((d_134_v8_)[default__.safeIndex(554, (d_134_v8_).length(0))]), False, (d_128_v2_) or ((d_140_v11_)[default__.safeIndex(d_127_v1_, len(d_140_v11_))]), d_128_v2_, d_146_globalState_)
            lhs13_ = d_146_globalState_
            lhs13_.f17 = rhs19_
            d_183_v48_ = rhs20_
            d_185_v50_: D0
            d_185_v50_ = D0_DC3(default__.fm36(d_146_globalState_), default__.fm2(default__.fm42(d_127_v1_, d_146_globalState_), len(d_162_v31_), d_146_globalState_), not((d_128_v2_) == (d_128_v2_)))
            pat_let_tv7_ = d_128_v2_
            pat_let_tv8_ = d_134_v8_
            pat_let_tv9_ = d_134_v8_
            pat_let_tv10_ = d_147_v17_
            def iife35_(_pat_let4_0):
                def iife36_(d_186_dt__update__tmp_h2_):
                    def iife37_(_pat_let5_0):
                        def iife38_(d_187_dt__update_hcf9_h0_):
                            def iife39_(_pat_let6_0):
                                def iife40_(d_188_dt__update_hcf8_h0_):
                                    return D0_DC3((d_186_dt__update__tmp_h2_).cf7, d_188_dt__update_hcf8_h0_, d_187_dt__update_hcf9_h0_)
                                return iife40_(_pat_let6_0)
                            return iife39_(((pat_let_tv9_)[default__.safeIndex(554, (pat_let_tv8_).length(0))]) < (len(pat_let_tv10_)))
                        return iife38_(_pat_let5_0)
                    return iife37_(pat_let_tv7_)
                return iife36_(_pat_let4_0)
            d_185_v50_ = iife35_(d_185_v50_)
        (d_146_globalState_).f6 = d_128_v2_
        (d_146_globalState_).f6 = d_128_v2_
        d_189_v51_: str
        d_189_v51_ = _dafny.CodePoint('e')
        (d_146_globalState_).f25 = d_189_v51_
        d_190_v52_: D8
        d_190_v52_ = D8_DC18(d_127_v1_)
        d_191_v53_: D8
        d_191_v53_ = D8_DC20(d_190_v52_)
        d_192_v54_: D8
        d_192_v54_ = D8_DC20(d_190_v52_)
        d_193_v55_: D8
        d_193_v55_ = D8_DC20(d_192_v54_)
        d_193_v55_ = D8_DC20(d_191_v53_)
        d_194_v56_: D0
        d_194_v56_ = D0_DC3(d_127_v1_, d_128_v2_, d_128_v2_)
        d_195_v57_: _dafny.Map
        d_195_v57_ = _dafny.Map({d_194_v56_: len(d_143_v14_)})
        d_196_v58_: C3
        nw17_ = C3()
        nw17_.ctor__(len(d_143_v14_), d_195_v57_, d_127_v1_, d_125_v0_, (d_127_v1_) != (d_127_v1_))
        d_196_v58_ = nw17_
        d_197_v59_: D10
        d_197_v59_ = D10_DC25(_dafny.MultiSet([d_128_v2_]))
        d_198_v60_: _dafny.Set
        d_198_v60_ = _dafny.Set({(d_197_v59_).cf40})
        d_199_i3_: int
        d_199_i3_ = 0
        with _dafny.label("0"):
            while (d_129_v3_) in (d_198_v60_):
                with _dafny.c_label("0"):
                    if (d_199_i3_) >= (100):
                        raise _dafny.Break("0")
                    d_199_i3_ = (d_199_i3_) + (1)
                    (d_146_globalState_).f25 = d_189_v51_
                    (d_167_v36_).m6(d_128_v2_, (d_128_v2_) and (d_128_v2_), default__.safeDivisionInt((0) - ((d_196_v58_).f34), len(d_140_v11_)), d_146_globalState_)
                    d_200_v61_: _dafny.MultiSet
                    d_200_v61_ = _dafny.MultiSet([d_127_v1_])
                    d_201_v62_: C9
                    nw18_ = C9()
                    nw18_.ctor__(d_128_v2_, d_200_v61_, d_127_v1_, d_125_v0_, d_128_v2_)
                    d_201_v62_ = nw18_
                    d_202_v63_: _dafny.Map
                    d_202_v63_ = _dafny.Map({d_127_v1_: d_201_v62_})
                    d_203_v64_: _dafny.Seq
                    d_203_v64_ = _dafny.SeqWithoutIsStrInference([d_147_v17_])
                    d_202_v63_ = (d_202_v63_).set(len(d_203_v64_), d_201_v62_)
                    d_204_v65_: _dafny.Map
                    d_204_v65_ = _dafny.Map({d_127_v1_: d_128_v2_})
                    d_205_v67_: _dafny.Set
                    def iife41_():
                        coll27_ = _dafny.Map()
                        compr_27_: int
                        for compr_27_ in ((d_201_v62_).f31).Elements:
                            d_206_v66_: int = compr_27_
                            if (d_206_v66_) in ((d_201_v62_).f31):
                                coll27_[(d_206_v66_) + ((d_196_v58_).f34)] = not(d_128_v2_)
                        return _dafny.Map(coll27_)
                    d_205_v67_ = _dafny.Set({d_204_v65_, iife41_()
                    })
                    (d_146_globalState_).f6 = ((d_205_v67_).intersection(d_205_v67_)).issubset((d_205_v67_) - (d_205_v67_))
                    pass
            pass
        d_207_v68_: _dafny.Map
        d_207_v68_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_189_v51_ for d_208_i4_ in range(default__.abs(498))])): len(_dafny.Map({(d_196_v58_).f34: d_128_v2_}))})
        d_209_v69_: _dafny.Seq
        d_209_v69_ = _dafny.SeqWithoutIsStrInference([d_167_v36_, d_167_v36_])
        d_210_v70_: bool
        d_211_v71_: bool
        d_212_v72_: str
        d_213_v73_: int
        out2_: bool
        out3_: bool
        out4_: str
        out5_: int
        out2_, out3_, out4_, out5_ = (d_196_v58_).m10((((d_207_v68_)[d_127_v1_] if (d_127_v1_) in (d_207_v68_) else len(d_209_v69_))) > ((d_196_v58_).fm14((d_196_v58_).f34, d_198_v60_, d_146_globalState_)), d_146_globalState_)
        d_210_v70_ = out2_
        d_211_v71_ = out3_
        d_212_v72_ = out4_
        d_213_v73_ = out5_
        d_214_v74_: _dafny.Set
        d_214_v74_ = _dafny.Set({d_125_v0_, d_125_v0_, d_125_v0_})
        d_215_v75_: _dafny.Map
        d_215_v75_ = _dafny.Map({d_214_v74_: (d_162_v31_).ispropersubset(d_162_v31_)})
        d_216_v76_: _dafny.Set
        d_216_v76_ = _dafny.Set({-854})
        d_217_v77_: _dafny.Map
        d_217_v77_ = _dafny.Map({d_211_v71_: d_216_v76_})
        d_218_v78_: _dafny.Map
        d_218_v78_ = _dafny.Map({d_213_v73_: d_216_v76_})
        d_215_v75_ = (d_215_v75_).set((d_214_v74_).intersection(d_214_v74_), (default__.fm0(d_218_v78_, d_210_v70_, True, d_146_globalState_)).issubset(((d_217_v77_)[d_128_v2_] if (d_128_v2_) in (d_217_v77_) else d_216_v76_)))
        if (d_140_v11_)[default__.safeIndex(d_213_v73_, len(d_140_v11_))]:
            d_127_v1_ = default__.safeDivisionInt((d_196_v58_).f34, (0) - ((d_196_v58_).f34))
            index12_ = default__.safeIndex(124, (d_134_v8_).length(0))
            (d_134_v8_)[index12_] = d_213_v73_
            index13_ = default__.safeIndex(124, (d_134_v8_).length(0))
            (d_134_v8_)[index13_] = default__.fm36(d_146_globalState_)
            if False:
                d_219_v79_: _dafny.Map
                d_219_v79_ = _dafny.Map({(d_196_v58_).f34: d_210_v70_})
                d_219_v79_ = ((d_219_v79_) | (default__.fm30(((d_219_v79_)[(d_134_v8_)[default__.safeIndex(124, (d_134_v8_).length(0))]] if ((d_134_v8_)[default__.safeIndex(124, (d_134_v8_).length(0))]) in (d_219_v79_) else d_210_v70_), d_143_v14_, d_128_v2_, d_146_globalState_)) if (564) != (d_213_v73_) else d_219_v79_)
                d_220_v80_: _dafny.Map
                d_220_v80_ = _dafny.Map({d_128_v2_: d_125_v0_})
                d_221_v81_: T0
                nw19_ = C7()
                nw19_.ctor__(d_125_v0_, d_211_v71_, d_127_v1_)
                d_221_v81_ = nw19_
                d_222_v82_: _dafny.Map
                d_222_v82_ = _dafny.Map({d_221_v81_: (d_221_v81_).f27})
                d_223_v83_: _dafny.Seq
                d_223_v83_ = _dafny.SeqWithoutIsStrInference([((d_220_v80_)[not(d_210_v70_)] if (not(d_210_v70_)) in (d_220_v80_) else ((d_222_v82_)[d_221_v81_] if (d_221_v81_) in (d_222_v82_) else d_125_v0_)), d_125_v0_, (d_221_v81_).f27, (d_221_v81_).f27])
                d_125_v0_ = (d_223_v83_)[default__.safeIndex(d_213_v73_, len(d_223_v83_))]
                (d_146_globalState_).f15 = d_210_v70_
                d_224_v84_: _dafny.Array
                nw20_ = _dafny.Array(int(0), 1)
                d_224_v84_ = nw20_
                d_225_v85_: _dafny.Seq
                d_225_v85_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_224_v84_])])
                d_226_v86_: C5
                nw21_ = C5()
                nw21_.ctor__(((d_225_v85_)[default__.safeIndex((d_196_v58_).fm14(83, d_198_v60_, d_146_globalState_), len(d_225_v85_))]) + ((d_138_v9_).set(default__.safeIndex((d_196_v58_).f34, len(d_138_v9_)), d_224_v84_)), ((d_196_v58_).f34 if d_210_v70_ else -87), d_125_v0_, d_210_v70_)
                d_226_v86_ = nw21_
                d_213_v73_ = d_213_v73_
            elif True:
                d_212_v72_ = d_212_v72_
                d_128_v2_ = not(d_128_v2_)
                d_227_v87_: C4
                nw22_ = C4()
                nw22_.ctor__()
                d_227_v87_ = nw22_
                d_228_v88_: _dafny.Map
                d_228_v88_ = _dafny.Map({(d_141_v12_) + (d_141_v12_): d_227_v87_})
                d_228_v88_ = (d_228_v88_).set(default__.fm55(d_210_v70_, d_146_globalState_), d_227_v87_)
                (d_146_globalState_).f21 = (d_196_v58_).f34
                index14_ = default__.safeIndex(690, (d_125_v0_).length(0))
                (d_125_v0_)[index14_] = d_211_v71_
                d_229_v89_: _dafny.Array
                nw23_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_229_v89_ = nw23_
                index15_ = default__.safeIndex(790, (d_229_v89_).length(0))
                (d_229_v89_)[index15_] = d_125_v0_
                d_230_v90_: D1
                d_230_v90_ = D1_DC6(318)
                d_231_v91_: _dafny.MultiSet
                d_231_v91_ = _dafny.MultiSet([D1_DC6(len((d_207_v68_).set((d_196_v58_).f34, 117))), d_230_v90_, D1_DC6(d_127_v1_)])
                index16_ = default__.safeIndex(690, (d_125_v0_).length(0))
                index17_ = default__.safeIndex(124, (d_134_v8_).length(0))
                index18_ = default__.safeIndex(124, (d_134_v8_).length(0))
                index19_ = default__.safeIndex(790, (d_229_v89_).length(0))
                rhs21_ = ((d_196_v58_).f34) == (956)
                rhs22_ = d_141_v12_
                rhs23_ = default__.fm4(d_231_v91_, d_146_globalState_)
                rhs24_ = (d_196_v58_).f34
                rhs25_ = d_125_v0_
                lhs14_ = d_125_v0_
                lhs15_ = default__.safeIndex(690, (d_125_v0_).length(0))
                lhs16_ = d_134_v8_
                lhs17_ = default__.safeIndex(124, (d_134_v8_).length(0))
                lhs18_ = d_134_v8_
                lhs19_ = default__.safeIndex(124, (d_134_v8_).length(0))
                lhs20_ = d_229_v89_
                lhs21_ = default__.safeIndex(790, (d_229_v89_).length(0))
                lhs14_[lhs15_] = rhs21_
                d_141_v12_ = rhs22_
                lhs16_[lhs17_] = rhs23_
                lhs18_[lhs19_] = rhs24_
                lhs20_[lhs21_] = rhs25_
            d_127_v1_ = d_127_v1_
            d_232_v92_: _dafny.MultiSet
            d_232_v92_ = _dafny.MultiSet([d_127_v1_])
            d_233_v93_: _dafny.MultiSet
            d_233_v93_ = _dafny.MultiSet([d_232_v92_, d_232_v92_])
            d_234_v94_: T0
            nw24_ = C2()
            nw24_.ctor__(default__.fm42(992, d_146_globalState_), d_211_v71_, d_125_v0_, d_210_v70_)
            d_234_v94_ = nw24_
            d_235_v95_: _dafny.Seq
            d_235_v95_ = _dafny.SeqWithoutIsStrInference([d_234_v94_, d_234_v94_])
            d_236_v96_: _dafny.Map
            d_236_v96_ = _dafny.Map({(d_233_v93_).set(d_232_v92_, default__.abs(len(d_235_v95_))): d_212_v72_})
            d_236_v96_ = (d_236_v96_).set(d_233_v93_, d_189_v51_)
        elif True:
            d_237_v97_: _dafny.MultiSet
            d_237_v97_ = _dafny.MultiSet([d_143_v14_])
            (d_146_globalState_).f17 = (d_237_v97_).issubset(d_237_v97_)
            d_189_v51_ = (d_212_v72_ if d_128_v2_ else d_189_v51_)
            d_238_v98_: _dafny.Array
            nw25_ = _dafny.Array(_dafny.Array(None, 0), 8)
            d_238_v98_ = nw25_
            d_239_v99_: _dafny.Map
            d_239_v99_ = _dafny.Map({d_211_v71_: d_238_v98_})
            d_240_v100_: D5
            d_240_v100_ = D5_DC13(False, d_134_v8_, d_128_v2_)
            d_238_v98_ = ((d_239_v99_)[(d_240_v100_).cf21] if ((d_240_v100_).cf21) in (d_239_v99_) else d_238_v98_)
            if default__.fm22(not((_dafny.Set({(0) - (-179), (d_196_v58_).f34})).issubset(d_216_v76_)), len((d_144_v15_)[default__.safeIndex(d_213_v73_, len(d_144_v15_))]), d_146_globalState_):
                d_241_v101_: T0
                nw26_ = C7()
                nw26_.ctor__(d_125_v0_, default__.fm2(d_212_v72_, (d_196_v58_).f34, d_146_globalState_), d_213_v73_)
                d_241_v101_ = nw26_
                d_242_v102_: T1
                nw27_ = C9()
                nw27_.ctor__(d_211_v71_, _dafny.MultiSet([(0) - (d_213_v73_)]), d_127_v1_, d_125_v0_, d_128_v2_)
                d_242_v102_ = nw27_
                d_243_v103_: _dafny.Map
                d_243_v103_ = _dafny.Map({d_242_v102_: d_241_v101_})
                rhs26_ = ((d_243_v103_)[d_242_v102_] if (d_242_v102_) in (d_243_v103_) else d_241_v101_)
                rhs27_ = d_167_v36_
                d_241_v101_ = rhs26_
                d_167_v36_ = rhs27_
                d_244_v104_: _dafny.Map
                d_244_v104_ = _dafny.Map({d_210_v70_: not(d_211_v71_)})
                d_244_v104_ = (d_244_v104_).set((d_241_v101_).f28, d_210_v70_)
                d_245_v105_: D8
                d_245_v105_ = D8_DC19()
                d_246_v106_: _dafny.Map
                d_246_v106_ = _dafny.Map({d_245_v105_: default__.fm40(d_146_globalState_)})
                d_247_v107_: _dafny.Map
                d_247_v107_ = _dafny.Map({860: d_246_v106_})
                rhs28_ = d_242_v102_.f29
                rhs29_ = not(d_211_v71_)
                rhs30_ = ((d_246_v106_) | (d_246_v106_)) | ((_dafny.Map({d_245_v105_: d_128_v2_})) | ((((d_247_v107_)[d_242_v102_.f29] if (d_242_v102_.f29) in (d_247_v107_) else d_246_v106_)).set(d_245_v105_, d_210_v70_)))
                rhs31_ = d_212_v72_
                lhs22_ = d_146_globalState_
                lhs23_ = d_146_globalState_
                lhs24_ = d_146_globalState_
                lhs22_.f4 = rhs28_
                lhs23_.f15 = rhs29_
                d_246_v106_ = rhs30_
                lhs24_.f25 = rhs31_
                d_144_v15_ = (_dafny.SeqWithoutIsStrInference([d_143_v14_ for d_248_i5_ in range(default__.abs(990))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_249_i6_ in range(default__.abs(808))]), d_143_v14_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqiqust")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uapyvflv"))]))
                (d_146_globalState_).f15 = d_211_v71_
            elif True:
                d_250_v108_: C6
                nw28_ = C6()
                nw28_.ctor__(len(d_141_v12_), d_125_v0_, d_211_v71_)
                d_250_v108_ = nw28_
                d_143_v14_ = ((d_143_v14_) + (d_143_v14_)) + (d_143_v14_)
                (d_146_globalState_).f6 = d_128_v2_
                d_251_v109_: _dafny.Set
                d_251_v109_ = _dafny.Set({_dafny.Map({d_210_v70_: (d_196_v58_).f34}), d_147_v17_, d_147_v17_, _dafny.Map({d_210_v70_: d_127_v1_}), default__.fm50(d_128_v2_, len(default__.fm31(d_146_globalState_)), d_146_globalState_)})
                d_252_v110_: _dafny.Map
                d_252_v110_ = _dafny.Map({d_251_v109_: d_134_v8_})
                d_252_v110_ = (d_252_v110_).set(d_251_v109_, d_134_v8_)
                d_253_v111_: C4
                nw29_ = C4()
                nw29_.ctor__()
                d_253_v111_ = nw29_
            d_127_v1_ = len(((d_143_v14_) + ((d_144_v15_)[default__.safeIndex(d_127_v1_, len(d_144_v15_))])) + ((d_143_v14_) + (default__.fm32((d_237_v97_).cardinality, d_210_v70_, d_146_globalState_))))
        d_254_v112_: _dafny.Map
        d_254_v112_ = _dafny.Map({d_210_v70_: d_128_v2_})
        d_254_v112_ = (d_254_v112_).set(not (not(d_211_v71_)) or (d_210_v70_), d_211_v71_)
        d_255_v113_: C0
        nw30_ = C0()
        nw30_.ctor__(d_143_v14_, d_127_v1_, d_125_v0_, d_210_v70_)
        d_255_v113_ = nw30_
        (d_146_globalState_).f21 = (default__.safeDivisionInt(d_213_v73_, d_127_v1_)) + (d_127_v1_)
        if not((d_128_v2_) or (d_211_v71_)):
            d_256_v115_: _dafny.MultiSet
            d_256_v115_ = _dafny.MultiSet([d_127_v1_])
            d_257_v116_: _dafny.Seq
            def iife42_():
                coll28_ = _dafny.Map()
                compr_28_: int
                for compr_28_ in (d_256_v115_).Elements:
                    d_258_v114_: int = compr_28_
                    if (d_258_v114_) in (d_256_v115_):
                        coll28_[(d_258_v114_) - ((d_196_v58_).f34)] = (d_140_v11_)[default__.safeIndex((0) - ((d_196_v58_).f34), len(d_140_v11_))]
                return _dafny.Map(coll28_)
            d_257_v116_ = _dafny.SeqWithoutIsStrInference([iife42_()
            ])
            d_257_v116_ = d_257_v116_
            d_207_v68_ = (d_207_v68_).set(d_213_v73_, len((d_255_v113_).f36))
            index20_ = default__.safeIndex(445, (d_125_v0_).length(0))
            (d_125_v0_)[index20_] = not(d_128_v2_)
            index21_ = default__.safeIndex(445, (d_125_v0_).length(0))
            rhs32_ = (d_144_v15_)[default__.safeIndex((d_196_v58_).f34, len(d_144_v15_))]
            rhs33_ = d_128_v2_
            rhs34_ = (d_127_v1_) >= (((d_256_v115_).cardinality) * (len(d_140_v11_)))
            rhs35_ = (_dafny.CodePoint('e')) in ((d_255_v113_).f36)
            lhs25_ = d_125_v0_
            lhs26_ = default__.safeIndex(445, (d_125_v0_).length(0))
            lhs27_ = d_146_globalState_
            d_143_v14_ = rhs32_
            lhs25_[lhs26_] = rhs33_
            d_211_v71_ = rhs34_
            lhs27_.f15 = rhs35_
            index22_ = default__.safeIndex(413, (d_134_v8_).length(0))
            (d_134_v8_)[index22_] = d_213_v73_
            d_259_v117_: D1
            d_259_v117_ = D1_DC6((d_196_v58_).f34)
            d_260_v118_: _dafny.MultiSet
            d_260_v118_ = _dafny.MultiSet([D1_DC6((d_256_v115_).cardinality), d_259_v117_, d_259_v117_, d_259_v117_])
            index23_ = default__.safeIndex(413, (d_134_v8_).length(0))
            (d_134_v8_)[index23_] = (default__.fm4(d_260_v118_, d_146_globalState_)) + (default__.safeDivisionInt(len(d_140_v11_), d_213_v73_))
            d_261_v119_: _dafny.Array
            nw31_ = _dafny.Array(D9.default()(), 11)
            d_261_v119_ = nw31_
            index24_ = default__.safeIndex(512, (d_261_v119_).length(0))
            (d_261_v119_)[index24_] = (D9_DC21(d_133_v7_) if (d_125_v0_)[default__.safeIndex(445, (d_125_v0_).length(0))] else D9_DC21(d_133_v7_))
            pat_let_tv11_ = d_133_v7_
            index25_ = default__.safeIndex(512, (d_261_v119_).length(0))
            def iife43_(_pat_let7_0):
                def iife44_(d_262_dt__update__tmp_h3_):
                    def iife45_(_pat_let8_0):
                        def iife46_(d_263_dt__update_hcf33_h0_):
                            return D9_DC21(d_263_dt__update_hcf33_h0_)
                        return iife46_(_pat_let8_0)
                    return iife45_(pat_let_tv11_)
                return iife44_(_pat_let7_0)
            rhs36_ = d_128_v2_
            rhs37_ = iife43_(D9_DC21(d_133_v7_))
            rhs38_ = d_211_v71_
            lhs28_ = d_146_globalState_
            lhs29_ = d_261_v119_
            lhs30_ = default__.safeIndex(512, (d_261_v119_).length(0))
            lhs28_.f6 = rhs36_
            lhs29_[lhs30_] = rhs37_
            d_128_v2_ = rhs38_
        elif True:
            index26_ = default__.safeIndex(758, (d_125_v0_).length(0))
            (d_125_v0_)[index26_] = d_211_v71_
            index27_ = default__.safeIndex(758, (d_125_v0_).length(0))
            (d_125_v0_)[index27_] = d_211_v71_
            d_264_v120_: _dafny.Map
            d_264_v120_ = _dafny.Map({default__.safeDivisionInt(892, (d_196_v58_).f34): d_210_v70_})
            d_264_v120_ = (d_264_v120_) | (d_264_v120_)
            (d_146_globalState_).f5 = (d_196_v58_).f34
            d_265_v121_: D0
            d_265_v121_ = D0_DC2((d_216_v76_).intersection(_dafny.Set({d_127_v1_, (d_196_v58_).f34})), d_143_v14_)
            d_265_v121_ = d_265_v121_
            d_266_v122_: _dafny.Array
            nw32_ = _dafny.Array(_dafny.Seq({}), 13)
            d_266_v122_ = nw32_
            d_267_v123_: _dafny.Map
            d_267_v123_ = _dafny.Map({d_162_v31_: d_128_v2_})
            d_268_v125_: _dafny.MultiSet
            d_268_v125_ = _dafny.MultiSet([_dafny.Set({d_211_v71_}), d_162_v31_])
            d_269_v129_: _dafny.Map
            d_269_v129_ = _dafny.Map({d_162_v31_: (d_196_v58_).f34})
            d_270_v130_: _dafny.Array
            nw33_ = _dafny.Array(None, 21)
            nw33_[int(0)] = d_267_v123_
            nw33_[int(1)] = d_267_v123_
            nw33_[int(2)] = d_267_v123_
            nw33_[int(3)] = d_267_v123_
            nw33_[int(4)] = d_267_v123_
            def iife47_():
                coll29_ = _dafny.Map()
                compr_29_: _dafny.Set
                for compr_29_ in (d_268_v125_).Elements:
                    d_271_v124_: _dafny.Set = compr_29_
                    if (d_271_v124_) in (d_268_v125_):
                        coll29_[d_271_v124_] = (d_125_v0_)[default__.safeIndex(758, (d_125_v0_).length(0))]
                return _dafny.Map(coll29_)
            nw33_[int(5)] = (iife47_()
            ) | (d_267_v123_)
            nw33_[int(6)] = default__.fm56(d_211_v71_, d_212_v72_, d_146_globalState_)
            nw33_[int(7)] = d_267_v123_
            nw33_[int(8)] = d_267_v123_
            nw33_[int(9)] = (d_267_v123_) | (d_267_v123_)
            def iife48_():
                coll30_ = _dafny.Map()
                compr_30_: _dafny.Set
                for compr_30_ in (d_268_v125_).Elements:
                    d_272_v126_: _dafny.Set = compr_30_
                    if (d_272_v126_) in (d_268_v125_):
                        coll30_[d_272_v126_] = d_128_v2_
                return _dafny.Map(coll30_)
            nw33_[int(10)] = iife48_()
            
            def iife49_():
                coll31_ = _dafny.Map()
                compr_31_: _dafny.Set
                for compr_31_ in (d_267_v123_).keys.Elements:
                    d_273_v127_: _dafny.Set = compr_31_
                    if (d_273_v127_) in (d_267_v123_):
                        coll31_[d_273_v127_] = (d_140_v11_)[default__.safeIndex(d_213_v73_, len(d_140_v11_))]
                return _dafny.Map(coll31_)
            nw33_[int(11)] = iife49_()
            
            nw33_[int(12)] = d_267_v123_
            nw33_[int(13)] = d_267_v123_
            def iife50_():
                coll32_ = _dafny.Map()
                compr_32_: _dafny.Set
                for compr_32_ in (d_269_v129_).keys.Elements:
                    d_274_v128_: _dafny.Set = compr_32_
                    if (d_274_v128_) in (d_269_v129_):
                        coll32_[d_274_v128_] = (d_125_v0_)[default__.safeIndex(758, (d_125_v0_).length(0))]
                return _dafny.Map(coll32_)
            nw33_[int(14)] = (d_267_v123_) | (iife50_()
            )
            nw33_[int(15)] = (_dafny.Map({d_162_v31_: d_128_v2_})).set(_dafny.Set({(d_125_v0_)[default__.safeIndex(758, (d_125_v0_).length(0))]}), True)
            nw33_[int(16)] = d_267_v123_
            nw33_[int(17)] = (default__.fm56(d_210_v70_, ((d_255_v113_).f36)[default__.safeIndex((d_194_v56_).cf7, len((d_255_v113_).f36))], d_146_globalState_)) | (_dafny.Map({_dafny.Set({(d_125_v0_)[default__.safeIndex(758, (d_125_v0_).length(0))], d_210_v70_, d_210_v70_}): True}))
            nw33_[int(18)] = d_267_v123_
            nw33_[int(19)] = d_267_v123_
            nw33_[int(20)] = d_267_v123_
            d_270_v130_ = nw33_
            index28_ = default__.safeIndex(896, (d_270_v130_).length(0))
            (d_270_v130_)[index28_] = d_267_v123_
            d_275_v131_: _dafny.MultiSet
            d_275_v131_ = _dafny.MultiSet([d_127_v1_, len((d_143_v14_) + (_dafny.SeqWithoutIsStrInference([d_189_v51_ for d_276_i7_ in range(default__.abs(873))])))])
            d_277_v132_: _dafny.Array
            nw34_ = _dafny.Array(_dafny.Array(None, 0), 16)
            d_277_v132_ = nw34_
            d_278_v133_: _dafny.Map
            d_278_v133_ = _dafny.Map({d_277_v132_: default__.safeModuloInt((0) - (-663), (d_196_v58_).f34)})
            d_279_v134_: D9
            d_279_v134_ = D9_DC22(d_134_v8_, d_128_v2_, d_277_v132_, (d_255_v113_).f36)
            d_280_v135_: _dafny.Seq
            d_280_v135_ = _dafny.SeqWithoutIsStrInference([d_267_v123_, d_267_v123_])
            index29_ = default__.safeIndex(896, (d_270_v130_).length(0))
            rhs39_ = (d_266_v122_ if d_211_v71_ else d_266_v122_)
            rhs40_ = ((d_278_v133_)[(d_279_v134_).cf36] if ((d_279_v134_).cf36) in (d_278_v133_) else d_213_v73_)
            rhs41_ = d_127_v1_
            rhs42_ = ((d_267_v123_) | ((d_280_v135_)[default__.safeIndex(d_127_v1_, len(d_280_v135_))])) | (_dafny.Map({_dafny.Set({d_210_v70_}): (d_125_v0_)[default__.safeIndex(758, (d_125_v0_).length(0))]}))
            rhs43_ = (d_275_v131_) - (d_275_v131_)
            lhs31_ = d_146_globalState_
            lhs32_ = d_270_v130_
            lhs33_ = default__.safeIndex(896, (d_270_v130_).length(0))
            d_266_v122_ = rhs39_
            lhs31_.f13 = rhs40_
            d_127_v1_ = rhs41_
            lhs32_[lhs33_] = rhs42_
            d_275_v131_ = rhs43_
        _dafny.print(_dafny.string_of((d_125_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v0_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_127_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_128_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v3_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_v4_) == (_dafny.Map({-900: _dafny.MultiSet([False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_132_v6_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v8_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v8_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v8_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_138_v9_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_139_v10_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_140_v11_) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_141_v12_) == (_dafny.SeqWithoutIsStrInference([900]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_142_v13_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_143_v14_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_144_v15_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prvhys"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_145_v16_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_146_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_146_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_146_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f8)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_globalState_.f9) == (_dafny.Map({-900: _dafny.MultiSet([False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f10) == (_dafny.MultiSet([900]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_146_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_146_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_146_globalState_).f16)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_146_globalState_.f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_globalState_).f19) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prvhys")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prvhys"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_globalState_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_146_globalState_.f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_146_globalState_).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_globalState_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_globalState_).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_146_globalState_.f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_globalState_).f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v17_) == (_dafny.Map({False: 820}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_v18_) == (_dafny.MultiSet([_dafny.Map({False: 900})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v31_) == (_dafny.Set({True, False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_163_v32_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True, False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v33_) == (_dafny.SeqWithoutIsStrInference([D5_DC12(_dafny.SeqWithoutIsStrInference([_dafny.Set({True, False})]))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v34_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D5_DC12(_dafny.SeqWithoutIsStrInference([_dafny.Set({True, False})]))])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_166_v35_)) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D5_DC12(_dafny.SeqWithoutIsStrInference([_dafny.Set({True, False})]))])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_189_v51_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_v52_).cf31))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_191_v53_).cf32).cf31))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_192_v54_).cf32).cf31))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_193_v55_).cf32).cf32).cf31))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_194_v56_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_194_v56_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_194_v56_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v57_) == (_dafny.Map({D0_DC3(1, False, False): 763}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v58_).f34))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_196_v58_).f35) == (_dafny.Map({D0_DC3(1, False, False): 763}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v59_).cf40) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_198_v60_) == (_dafny.Set({_dafny.MultiSet([False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_199_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_207_v68_) == (_dafny.Map({498: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_209_v69_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_210_v70_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_211_v71_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_212_v72_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_213_v73_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_214_v74_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_215_v75_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v76_) == (_dafny.Set({-854}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v77_) == (_dafny.Map({True: _dafny.Set({-854})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v78_) == (_dafny.Map({763: _dafny.Set({-854})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v112_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_255_v113_).f36).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(_dafny.Map({}), False, int(0), _dafny.CodePoint('D'), False)
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
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D0_DC4)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1)
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf5)}, {self.cf6.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC4(D0, NamedTuple('DC4', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC4({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC4) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC6(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D1_DC7)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC6(D1, NamedTuple('DC6', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC7(D1, NamedTuple('DC7', [('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC7({_dafny.string_of(self.cf13)}, {self.cf14.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC7) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

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
        return lambda: D3_DC10(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC11(D4, NamedTuple('DC11', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13(False, _dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC13(D5, NamedTuple('DC13', [('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)

class D6_DC14(D6, NamedTuple('DC14', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC16(False, _dafny.Map({}), _dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)

class D7_DC16(D7, NamedTuple('DC16', [('cf26', Any), ('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC15(D7, NamedTuple('DC15', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC18(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D8_DC17)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)

class D8_DC18(D8, NamedTuple('DC18', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC19(D8, NamedTuple('DC19', [])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19)
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC17(D8, NamedTuple('DC17', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC17({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC17) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC22(_dafny.Array(None, 0), False, _dafny.Array(None, 0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)

class D9_DC22(D9, NamedTuple('DC22', [('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {self.cf37.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC23(D9, NamedTuple('DC23', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC21(D9, NamedTuple('DC21', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC26(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D10_DC25)

class D10_DC26(D10, NamedTuple('DC26', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC25(D10, NamedTuple('DC25', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC25({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC25) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC28()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D11_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)

class D11_DC28(D11, NamedTuple('DC28', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)

class D12_DC29(D12, NamedTuple('DC29', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D13_DC30)

class D13_DC30(D13, NamedTuple('DC30', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC30({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC30) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC32()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D14_DC32)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D14_DC33)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D14_DC31)

class D14_DC32(D14, NamedTuple('DC32', [])):
    def __dafnystr__(self) -> str:
        return f'D14.DC32'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC32)
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC33(D14, NamedTuple('DC33', [('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC33({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC33) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC31(D14, NamedTuple('DC31', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC31({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC31) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC35(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D15_DC35)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D15_DC34)

class D15_DC35(D15, NamedTuple('DC35', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC35({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC35) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC34(D15, NamedTuple('DC34', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC34({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC34) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC37(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D16_DC37)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D16_DC38)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D16_DC36)

class D16_DC37(D16, NamedTuple('DC37', [('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC37({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC37) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC38(D16, NamedTuple('DC38', [('cf53', Any), ('cf54', Any), ('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC38({_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC38) and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC36(D16, NamedTuple('DC36', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC36({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC36) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D17_DC39)

class D17_DC39(D17, NamedTuple('DC39', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC39({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC39) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC41(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D18_DC41)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D18_DC40)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D18_DC42)

class D18_DC41(D18, NamedTuple('DC41', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC41({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC41) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC40(D18, NamedTuple('DC40', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC40({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC40) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC42(D18, NamedTuple('DC42', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC42({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC42) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value

class GlobalState:
    def  __init__(self):
        self.f4: int = int(0)
        self.f5: int = int(0)
        self.f6: bool = False
        self.f9: _dafny.Map = _dafny.Map({})
        self.f13: int = int(0)
        self.f15: bool = False
        self.f17: bool = False
        self.f18: _dafny.Array = _dafny.Array(None, 0)
        self.f21: int = int(0)
        self.f25: str = _dafny.CodePoint('D')
        self._f0: bool = False
        self._f1: bool = False
        self._f2: int = int(0)
        self._f3: bool = False
        self._f7: int = int(0)
        self._f8: _dafny.Array = _dafny.Array(None, 0)
        self._f10: _dafny.MultiSet = _dafny.MultiSet({})
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        self._f12: int = int(0)
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        self._f16: _dafny.Seq = _dafny.Seq({})
        self._f19: _dafny.Seq = _dafny.Seq({})
        self._f20: int = int(0)
        self._f22: _dafny.Map = _dafny.Map({})
        self._f23: int = int(0)
        self._f24: int = int(0)
        self._f26: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self).f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self).f15 = f15
        (self)._f16 = f16
        (self).f17 = f17
        (self).f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self).f21 = f21
        (self)._f22 = f22
        (self)._f23 = f23
        (self)._f24 = f24
        (self).f25 = f25
        (self)._f26 = f26

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
    def f14(self):
        return self._f14
    @property
    def f16(self):
        return self._f16
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
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
    def f26(self):
        return self._f26

class C0(T1, T0):
    def  __init__(self):
        self._f29: int = int(0)
        self._f27: _dafny.Array = _dafny.Array(None, 0)
        self._f28: bool = False
        self._f36: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f36, f29, f27, f28):
        (self)._f36 = f36
        (self).f29 = f29
        (self)._f27 = f27
        (self)._f28 = f28

    def fm15(self, globalState):
        return (self.f29) + ((0) - (self.f29))

    @property
    def f36(self):
        return self._f36

class C1(T0):
    def  __init__(self):
        self._f27: _dafny.Array = _dafny.Array(None, 0)
        self._f28: bool = False
        self._f39: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f39, f27, f28):
        (self)._f39 = f39
        (self)._f27 = f27
        (self)._f28 = f28

    def fm34(self, p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhictu"))

    def fm35(self, p0, p1, globalState):
        return D3_DC9(_dafny.Map({801: _dafny.Set({527, 65, 543})}))

    def m14(self, p0, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: _dafny.Seq = _dafny.Seq({})
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r3: _dafny.Map = _dafny.Map({})
        d_281_v0_: _dafny.Seq
        d_281_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usplmirjv"))
        (globalState).f13 = len(d_281_v0_)
        d_282_v1_: int
        d_282_v1_ = 511
        d_283_v2_: _dafny.MultiSet
        d_283_v2_ = _dafny.MultiSet([d_282_v1_, d_282_v1_, d_282_v1_])
        d_284_v3_: _dafny.Seq
        d_284_v3_ = _dafny.SeqWithoutIsStrInference([d_282_v1_])
        if ((_dafny.MultiSet(d_284_v3_)) - (d_283_v2_)).issubset((d_283_v2_).set(d_282_v1_, default__.abs(d_282_v1_))):
            index30_ = default__.safeIndex(587, ((self).f27).length(0))
            ((self).f27)[index30_] = (self).f39
            index31_ = default__.safeIndex(587, ((self).f27).length(0))
            ((self).f27)[index31_] = p0
            d_285_v4_: _dafny.Map
            d_285_v4_ = _dafny.Map({default__.safeDivisionInt((0) - ((d_283_v2_).cardinality), len(d_281_v0_)): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jj"))})
            d_285_v4_ = d_285_v4_
            d_286_v5_: _dafny.Seq
            d_286_v5_ = _dafny.SeqWithoutIsStrInference([d_281_v0_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_287_i0_ in range(default__.abs(231))])])
            d_288_v6_: _dafny.Array
            nw35_ = _dafny.Array(None, 4)
            nw35_[int(0)] = ((d_286_v5_)[default__.safeIndex(d_282_v1_, len(d_286_v5_))]) + (d_281_v0_)
            nw35_[int(1)] = d_281_v0_
            nw35_[int(2)] = d_281_v0_
            nw35_[int(3)] = d_281_v0_
            d_288_v6_ = nw35_
            d_289_v7_: D9
            d_289_v7_ = D9_DC21(d_288_v6_)
            d_288_v6_ = (d_289_v7_).cf33
            index32_ = default__.safeIndex(587, ((self).f27).length(0))
            ((self).f27)[index32_] = (d_282_v1_) <= (default__.fm36(globalState))
            d_290_v8_: _dafny.Seq
            d_290_v8_ = _dafny.SeqWithoutIsStrInference([p0, ((self).f27)[default__.safeIndex(587, ((self).f27).length(0))]])
            (globalState).f5 = default__.safeDivisionInt((len(d_290_v8_)) - (len(d_284_v3_)), len((d_281_v0_) + (d_281_v0_)))
        elif True:
            d_291_v9_: D8
            d_291_v9_ = D8_DC18(d_282_v1_)
            d_292_v10_: _dafny.Set
            d_292_v10_ = _dafny.Set({d_291_v9_, d_291_v9_, d_291_v9_, D8_DC18(d_282_v1_), d_291_v9_})
            d_292_v10_ = (d_292_v10_).intersection((d_292_v10_) - (d_292_v10_))
            d_293_v11_: _dafny.Set
            d_293_v11_ = _dafny.Set({False})
            (globalState).f5 = (783) * (len(d_293_v11_))
            (globalState).f5 = (d_282_v1_) * (d_282_v1_)
            (globalState).f13 = (d_282_v1_) * ((d_282_v1_) + (d_282_v1_))
            (globalState).f6 = p0
        (globalState).f15 = (self).f39
        d_294_v12_: C0
        nw36_ = C0()
        nw36_.ctor__(d_281_v0_, d_282_v1_, (self).f27, p0)
        d_294_v12_ = nw36_
        (globalState).f15 = not(not(p0))
        d_295_v13_: _dafny.MultiSet
        d_295_v13_ = _dafny.MultiSet([(self).f39])
        d_296_v14_: D10
        d_296_v14_ = D10_DC25(d_295_v13_)
        hi0_ = (_dafny.MultiSet([((d_295_v13_)[(self).f39] if ((self).f39) in (d_295_v13_) else d_282_v1_)])).cardinality
        for d_297_i1_ in range(((d_296_v14_).cf40).cardinality, hi0_):
            d_298_v16_: _dafny.Array
            def lambda5_(d_299_i1_, d_300_v1_):
                def lambda6_(d_301_i2_):
                    def iife51_():
                        coll33_ = _dafny.Map()
                        compr_33_: int
                        for compr_33_ in _dafny.IntegerRange(-710, 610):
                            d_302_v15_: int = compr_33_
                            if ((-710) <= (d_302_v15_)) and ((d_302_v15_) < (610)):
                                coll33_[(d_302_v15_) * (len(_dafny.Set({d_299_i1_, 954})))] = _dafny.Map({d_300_v1_: 947})
                        return _dafny.Map(coll33_)
                    return iife51_()
                    

                return lambda6_

            init3_ = lambda5_(d_297_i1_, d_282_v1_)
            nw37_ = _dafny.Array(None, 15)
            for i0_3_ in range(nw37_.length(0)):
                nw37_[i0_3_] = init3_(i0_3_)
            d_298_v16_ = nw37_
            d_303_v18_: _dafny.Map
            def iife52_():
                coll34_ = _dafny.Map()
                compr_34_: int
                for compr_34_ in _dafny.IntegerRange(-859, 455):
                    d_304_v17_: int = compr_34_
                    if ((-859) <= (d_304_v17_)) and ((d_304_v17_) < (455)):
                        coll34_[default__.safeModuloInt(d_304_v17_, d_297_i1_)] = d_297_i1_
                return _dafny.Map(coll34_)
            d_303_v18_ = _dafny.Map({d_297_i1_: ((iife52_()
            ).set(d_282_v1_, d_282_v1_)).set(((d_283_v2_)[d_297_i1_] if (d_297_i1_) in (d_283_v2_) else d_282_v1_), d_297_i1_)})
            index33_ = default__.safeIndex(283, (d_298_v16_).length(0))
            (d_298_v16_)[index33_] = d_303_v18_
            index34_ = default__.safeIndex(283, (d_298_v16_).length(0))
            (d_298_v16_)[index34_] = d_303_v18_
            (globalState).f5 = ((d_295_v13_).cardinality if (self).f39 else d_282_v1_)
            d_305_v19_: _dafny.Array
            def lambda7_(d_306_i1_):
                def lambda8_(d_307_i3_):
                    return default__.safeModuloInt(d_307_i3_, d_306_i1_)

                return lambda8_

            init4_ = lambda7_(d_297_i1_)
            nw38_ = _dafny.Array(None, 21)
            for i0_4_ in range(nw38_.length(0)):
                nw38_[i0_4_] = init4_(i0_4_)
            d_305_v19_ = nw38_
            index35_ = default__.safeIndex(165, (d_305_v19_).length(0))
            (d_305_v19_)[index35_] = d_282_v1_
            index36_ = default__.safeIndex(165, (d_305_v19_).length(0))
            (d_305_v19_)[index36_] = (default__.safeDivisionInt(len(d_284_v3_), d_282_v1_)) + (d_282_v1_)
            index37_ = default__.safeIndex(165, (d_305_v19_).length(0))
            (d_305_v19_)[index37_] = (d_305_v19_)[default__.safeIndex(165, (d_305_v19_).length(0))]
        r0 = ((d_294_v12_).f36) + (d_281_v0_)
        d_308_v20_: D9
        d_308_v20_ = D9_DC23(p0)
        pat_let_tv12_ = p0
        pat_let_tv13_ = p0
        pat_let_tv14_ = d_282_v1_
        pat_let_tv15_ = p0
        pat_let_tv16_ = p0
        pat_let_tv17_ = p0
        pat_let_tv18_ = p0
        def lambda9_(source6_):
            if source6_.is_DC22:
                d_309___mcc_h0_ = source6_.cf34
                d_310___mcc_h1_ = source6_.cf35
                d_311___mcc_h2_ = source6_.cf36
                d_312___mcc_h3_ = source6_.cf37
                d_313_cf37_ = d_312___mcc_h3_
                d_314_cf36_ = d_311___mcc_h2_
                d_315_cf35_ = d_310___mcc_h1_
                d_316_cf34_ = d_309___mcc_h0_
                return ((_dafny.SeqWithoutIsStrInference([pat_let_tv12_, pat_let_tv13_])) + (_dafny.SeqWithoutIsStrInference([(self).f28]))).set(default__.safeIndex(pat_let_tv14_, len((_dafny.SeqWithoutIsStrInference([pat_let_tv15_, pat_let_tv16_])) + (_dafny.SeqWithoutIsStrInference([(self).f28])))), (self).f28)
            elif source6_.is_DC23:
                d_317___mcc_h4_ = source6_.cf38
                d_318_cf38_ = d_317___mcc_h4_
                return (_dafny.SeqWithoutIsStrInference([pat_let_tv17_, d_318_cf38_])) + (_dafny.SeqWithoutIsStrInference([d_318_cf38_, (self).f39]))
            elif source6_.is_DC21:
                d_319___mcc_h5_ = source6_.cf33
                d_320_cf33_ = d_319___mcc_h5_
                return (_dafny.SeqWithoutIsStrInference([(self).f28])) + (_dafny.SeqWithoutIsStrInference([False]))
            elif True:
                d_321___mcc_h6_ = source6_.cf39
                d_322_cf39_ = d_321___mcc_h6_
                return (_dafny.SeqWithoutIsStrInference([pat_let_tv18_])) + (_dafny.SeqWithoutIsStrInference([(self).f28]))

        r1 = lambda9_(d_308_v20_)
        d_323_v21_: _dafny.Seq
        d_323_v21_ = _dafny.SeqWithoutIsStrInference([(d_294_v12_).f36])
        r2 = (d_323_v21_)[default__.safeIndex(default__.safeModuloInt(default__.fm36(globalState), d_282_v1_), len(d_323_v21_))]
        d_324_v22_: _dafny.Seq
        d_324_v22_ = _dafny.SeqWithoutIsStrInference([(self).f28])
        d_325_v23_: _dafny.Seq
        d_325_v23_ = _dafny.SeqWithoutIsStrInference([d_324_v22_, d_324_v22_])
        r3 = default__.fm37(p0, False, d_282_v1_, len((d_325_v23_)[default__.safeIndex(d_282_v1_, len(d_325_v23_))]), globalState)
        return r0, r1, r2, r3

    def m15(self, p0, p1, p2, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: bool = False
        d_326_v0_: _dafny.Array
        nw39_ = _dafny.Array(_dafny.Map({}), 11)
        d_326_v0_ = nw39_
        index38_ = default__.safeIndex(220, (d_326_v0_).length(0))
        (d_326_v0_)[index38_] = _dafny.Map({(self).f28: p1})
        d_327_v1_: _dafny.Set
        d_327_v1_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubsto"))})
        d_328_v2_: _dafny.Map
        d_328_v2_ = _dafny.Map({not(not((self).f39)): (p1) + ((0) - (len(d_327_v1_)))})
        index39_ = default__.safeIndex(220, (d_326_v0_).length(0))
        (d_326_v0_)[index39_] = d_328_v2_
        d_329_v3_: _dafny.Array
        nw40_ = _dafny.Array(_dafny.Map({}), 9)
        d_329_v3_ = nw40_
        d_330_v4_: D11
        d_330_v4_ = D11_DC27(d_329_v3_)
        pat_let_tv19_ = d_329_v3_
        d_331_v5_: _dafny.Array
        nw41_ = _dafny.Array(None, 11)
        nw41_[int(0)] = d_329_v3_
        nw41_[int(1)] = d_329_v3_
        nw41_[int(2)] = d_329_v3_
        nw41_[int(3)] = d_329_v3_
        nw41_[int(4)] = d_329_v3_
        def iife53_(_pat_let9_0):
            def iife54_(d_332_dt__update__tmp_h0_):
                def iife55_(_pat_let10_0):
                    def iife56_(d_333_dt__update_hcf42_h0_):
                        return D11_DC27(d_333_dt__update_hcf42_h0_)
                    return iife56_(_pat_let10_0)
                return iife55_(pat_let_tv19_)
            return iife54_(_pat_let9_0)
        nw41_[int(5)] = (iife53_(d_330_v4_)).cf42
        nw41_[int(6)] = d_329_v3_
        nw41_[int(7)] = d_329_v3_
        nw41_[int(8)] = d_329_v3_
        nw41_[int(9)] = d_329_v3_
        nw41_[int(10)] = d_329_v3_
        d_331_v5_ = nw41_
        index40_ = default__.safeIndex(77, (d_331_v5_).length(0))
        (d_331_v5_)[index40_] = d_329_v3_
        index41_ = default__.safeIndex(77, (d_331_v5_).length(0))
        (d_331_v5_)[index41_] = d_329_v3_
        d_334_v6_: _dafny.Seq
        d_334_v6_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27, (self).f27])
        d_335_v7_: _dafny.Map
        d_335_v7_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference([(self).f27])})
        d_334_v6_ = (((d_335_v7_)[p1] if (p1) in (d_335_v7_) else d_334_v6_)) + (_dafny.SeqWithoutIsStrInference([(self).f27]))
        if (self).f28:
            (globalState).f6 = (p1) == (p1)
            d_336_v8_: _dafny.Seq
            d_336_v8_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_337_v9_: _dafny.Seq
            d_337_v9_ = (d_336_v8_) + (d_336_v8_)
            d_337_v9_ = d_337_v9_
            (globalState).f15 = ((p0).intersection(p0)).ispropersubset((p0) | (p0))
            d_338_v10_: _dafny.Array
            nw42_ = _dafny.Array(_dafny.Map({}), 11)
            d_338_v10_ = nw42_
            d_339_v11_: _dafny.Map
            d_339_v11_ = _dafny.Map({p1: p1})
            index42_ = default__.safeIndex(218, (d_338_v10_).length(0))
            (d_338_v10_)[index42_] = d_339_v11_
            index43_ = default__.safeIndex(218, (d_338_v10_).length(0))
            (d_338_v10_)[index43_] = (d_339_v11_) | ((d_339_v11_) | (d_339_v11_))
            d_340_v12_: _dafny.Array
            nw43_ = _dafny.Array(int(0), 29)
            d_340_v12_ = nw43_
            d_341_v13_: _dafny.Seq
            d_341_v13_ = _dafny.SeqWithoutIsStrInference([d_340_v12_])
            d_342_v14_: _dafny.Seq
            d_342_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xuh"))
            d_343_v15_: _dafny.Map
            d_343_v15_ = _dafny.Map({((0) - (p1)) * (p1): (len(d_341_v13_)) > (len(d_342_v14_))})
            d_344_v16_: _dafny.Seq
            d_344_v16_ = _dafny.SeqWithoutIsStrInference([(self).f28, (self).f39])
            d_345_v17_: _dafny.MultiSet
            d_345_v17_ = _dafny.MultiSet([p0, default__.fm38(p1, not((self).f28), (self).f28, (self).f39, globalState)])
            d_343_v15_ = (d_343_v15_).set(len((d_344_v16_) + (d_344_v16_)), (d_345_v17_) == (d_345_v17_))
        elif True:
            d_346_v18_: _dafny.Array
            nw44_ = _dafny.Array(int(0), 1)
            d_346_v18_ = nw44_
            d_347_v19_: _dafny.MultiSet
            d_347_v19_ = _dafny.MultiSet([p1, len(_dafny.Map({(self).f28: -418}))])
            d_348_v20_: _dafny.Set
            d_348_v20_ = _dafny.Set({d_347_v19_, _dafny.MultiSet([p1, p1]), d_347_v19_})
            d_349_v21_: _dafny.Map
            d_349_v21_ = _dafny.Map({d_346_v18_: d_348_v20_})
            d_349_v21_ = (d_349_v21_).set(d_346_v18_, (d_348_v20_) | (d_348_v20_))
            if (self).f28:
                d_350_v22_: _dafny.Array
                nw45_ = _dafny.Array(_dafny.Array(None, 0), 14)
                d_350_v22_ = nw45_
                index44_ = default__.safeIndex(942, (d_350_v22_).length(0))
                (d_350_v22_)[index44_] = d_346_v18_
                index45_ = default__.safeIndex(942, (d_350_v22_).length(0))
                (d_350_v22_)[index45_] = d_346_v18_
                arr0_ = (d_350_v22_)[default__.safeIndex(942, (d_350_v22_).length(0))]
                index46_ = default__.safeIndex(756, ((d_350_v22_)[default__.safeIndex(942, (d_350_v22_).length(0))]).length(0))
                arr0_[index46_] = p1
                arr1_ = (d_350_v22_)[default__.safeIndex(942, (d_350_v22_).length(0))]
                index47_ = default__.safeIndex(756, ((d_350_v22_)[default__.safeIndex(942, (d_350_v22_).length(0))]).length(0))
                def iife57_():
                    coll35_ = _dafny.Map()
                    compr_35_: int
                    for compr_35_ in _dafny.IntegerRange(580, 746):
                        d_351_v23_: int = compr_35_
                        if ((580) <= (d_351_v23_)) and ((d_351_v23_) < (746)):
                            coll35_[default__.safeModuloInt(d_351_v23_, len(_dafny.SeqWithoutIsStrInference([p1])))] = (self).f39
                    return _dafny.Map(coll35_)
                arr1_[index47_] = (len(_dafny.Set({p1, len(iife57_()
                ), p1, (0) - (p1)}))) * (p1)
                d_352_v24_: _dafny.Seq
                d_352_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txp"))
                d_353_v25_: _dafny.Seq
                d_353_v25_ = _dafny.SeqWithoutIsStrInference([(self).f39])
                d_352_v24_ = (self).fm34((False) and (not((self).f28)), 222, _dafny.MultiSet([(self).f39, (d_353_v25_)[default__.safeIndex(((d_350_v22_)[default__.safeIndex(942, (d_350_v22_).length(0))])[default__.safeIndex(756, ((d_350_v22_)[default__.safeIndex(942, (d_350_v22_).length(0))]).length(0))], len(d_353_v25_))], p2, (self).f28, (self).f28]), (0) - (((d_350_v22_)[default__.safeIndex(942, (d_350_v22_).length(0))])[default__.safeIndex(756, ((d_350_v22_)[default__.safeIndex(942, (d_350_v22_).length(0))]).length(0))]), globalState)
                d_352_v24_ = (d_352_v24_) + (d_352_v24_)
                d_354_v26_: _dafny.Map
                d_354_v26_ = _dafny.Map({p1: (self).f28})
                d_355_v27_: _dafny.Seq
                d_355_v27_ = _dafny.SeqWithoutIsStrInference([d_347_v19_, (d_347_v19_).set(p1, default__.abs(len(d_354_v26_))), d_347_v19_])
                d_355_v27_ = (_dafny.SeqWithoutIsStrInference([(d_347_v19_).set(len(_dafny.Map({159: _dafny.MultiSet([(self).f39])})), default__.abs(p1)) for d_356_i0_ in range(default__.abs(-501))])) + (_dafny.SeqWithoutIsStrInference([d_347_v19_ for d_357_i1_ in range(default__.abs(453))]))
            elif True:
                d_358_v28_: _dafny.Array
                nw46_ = _dafny.Array(_dafny.CodePoint('D'), 27)
                d_358_v28_ = nw46_
                index48_ = default__.safeIndex(572, (d_346_v18_).length(0))
                (d_346_v18_)[index48_] = p1
                d_359_v29_: _dafny.Map
                d_359_v29_ = _dafny.Map({_dafny.MultiSet([(self).f39]): d_346_v18_})
                d_360_v30_: _dafny.Map
                d_360_v30_ = _dafny.Map({(self).f28: (self).f28})
                index49_ = default__.safeIndex(572, (d_346_v18_).length(0))
                rhs44_ = default__.safeDivisionInt(p1, p1)
                rhs45_ = d_358_v28_
                rhs46_ = (0) - (len((d_359_v29_ if p2 else (d_359_v29_) | (d_359_v29_))))
                rhs47_ = ((d_360_v30_)[not (p2) or ((self).f39)] if (not (p2) or ((self).f39)) in (d_360_v30_) else p2)
                rhs48_ = p2
                lhs34_ = globalState
                lhs35_ = d_346_v18_
                lhs36_ = default__.safeIndex(572, (d_346_v18_).length(0))
                lhs37_ = globalState
                lhs38_ = globalState
                lhs34_.f5 = rhs44_
                d_358_v28_ = rhs45_
                lhs35_[lhs36_] = rhs46_
                lhs37_.f15 = rhs47_
                lhs38_.f17 = rhs48_
                d_361_v31_: _dafny.MultiSet
                d_361_v31_ = _dafny.MultiSet([(self).f39])
                d_362_v32_: _dafny.Seq
                d_362_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
                d_363_v33_: _dafny.Map
                d_363_v33_ = _dafny.Map({p0: (d_346_v18_)[default__.safeIndex(572, (d_346_v18_).length(0))]})
                d_364_v34_: _dafny.Map
                d_364_v34_ = _dafny.Map({((d_363_v33_)[_dafny.Set({p2})] if (_dafny.Set({p2})) in (d_363_v33_) else p1): (self).f39})
                d_365_v35_: str
                d_365_v35_ = _dafny.CodePoint('i')
                d_366_v36_: C0
                nw47_ = C0()
                nw47_.ctor__((self).fm34((self).f39, (d_346_v18_)[default__.safeIndex(572, (d_346_v18_).length(0))], d_361_v31_, p1, globalState), len((d_362_v32_).set(default__.safeIndex(len(d_364_v34_), len(d_362_v32_)), d_365_v35_)), (self).f27, not(p2))
                d_366_v36_ = nw47_
                d_367_v37_: _dafny.Array
                nw48_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
                d_367_v37_ = nw48_
                d_368_v38_: _dafny.Map
                d_368_v38_ = _dafny.Map({d_367_v37_: p2})
                d_368_v38_ = (d_368_v38_).set(d_367_v37_, (self).f28)
                index50_ = default__.safeIndex(911, ((self).f27).length(0))
                ((self).f27)[index50_] = (self).f39
                index51_ = default__.safeIndex(911, ((self).f27).length(0))
                index52_ = default__.safeIndex(572, (d_346_v18_).length(0))
                rhs49_ = (d_362_v32_)[default__.safeIndex((d_346_v18_)[default__.safeIndex(572, (d_346_v18_).length(0))], len(d_362_v32_))]
                rhs50_ = (p2 if (self).f28 else not((self).f28))
                rhs51_ = (self).f28
                rhs52_ = (p1) - (((0) - (default__.fm36(globalState)) if (self).f28 else p1))
                lhs39_ = (self).f27
                lhs40_ = default__.safeIndex(911, ((self).f27).length(0))
                lhs41_ = globalState
                lhs42_ = d_346_v18_
                lhs43_ = default__.safeIndex(572, (d_346_v18_).length(0))
                d_365_v35_ = rhs49_
                lhs39_[lhs40_] = rhs50_
                lhs41_.f6 = rhs51_
                lhs42_[lhs43_] = rhs52_
                d_369_v39_: _dafny.Array
                nw49_ = _dafny.Array(None, 1)
                nw49_[int(0)] = p1
                d_369_v39_ = nw49_
                d_370_v40_: _dafny.Map
                d_370_v40_ = _dafny.Map({(0) - ((d_346_v18_)[default__.safeIndex(572, (d_346_v18_).length(0))]): d_365_v35_})
                d_371_v41_: _dafny.Seq
                d_371_v41_ = _dafny.SeqWithoutIsStrInference([p1, p1, len(d_370_v40_)])
                d_372_v42_: _dafny.Map
                d_372_v42_ = _dafny.Map({d_371_v41_: d_369_v39_})
                d_373_v43_: _dafny.Seq
                d_373_v43_ = _dafny.SeqWithoutIsStrInference([d_369_v39_, d_369_v39_, d_369_v39_])
                d_374_v44_: _dafny.Set
                d_374_v44_ = _dafny.Set({_dafny.CodePoint('k'), d_365_v35_})
                d_369_v39_ = ((d_372_v42_)[d_371_v41_] if (d_371_v41_) in (d_372_v42_) else (d_373_v43_)[default__.safeIndex(len(d_374_v44_), len(d_373_v43_))])
            (globalState).f21 = p1
            d_375_v45_: _dafny.MultiSet
            d_375_v45_ = _dafny.MultiSet([(self).f39])
            if (p1) != ((((d_375_v45_)[p2] if (p2) in (d_375_v45_) else p1) if (self).f28 else p1)):
                d_376_v46_: _dafny.Seq
                d_376_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qouh"))
                d_377_v47_: str
                d_377_v47_ = _dafny.CodePoint('c')
                d_378_v48_: C0
                nw50_ = C0()
                nw50_.ctor__(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nulmgbv"))).set(default__.safeIndex(len(d_376_v46_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nulmgbv")))), d_377_v47_)) + (d_376_v46_), p1, (self).f27, not ((self).f28) or ((self).f28))
                d_378_v48_ = nw50_
                d_379_v49_: D0
                d_379_v49_ = D0_DC1()
                d_380_v50_: D0
                d_380_v50_ = D0_DC4(d_379_v49_)
                d_380_v50_ = d_380_v50_
                (globalState).f21 = (0) - (default__.safeModuloInt(-95, p1))
                d_381_v51_: _dafny.Seq
                d_381_v51_ = _dafny.SeqWithoutIsStrInference([p2, (self).f28])
                d_382_v52_: _dafny.Map
                d_382_v52_ = _dafny.Map({len(d_381_v51_): (d_378_v48_).f36})
                d_383_v53_: D7
                d_383_v53_ = D7_DC15(d_382_v52_)
                d_384_v54_: _dafny.Seq
                d_384_v54_ = _dafny.SeqWithoutIsStrInference([d_383_v53_])
                d_384_v54_ = (d_384_v54_) + ((d_384_v54_) + (default__.fm39(globalState)))
                d_385_v55_: _dafny.Map
                d_385_v55_ = _dafny.Map({p1: p2})
                d_386_v56_: _dafny.Set
                d_386_v56_ = _dafny.Set({p1, len((d_385_v55_) | (d_385_v55_)), p1})
                index53_ = default__.safeIndex(102, ((self).f27).length(0))
                ((self).f27)[index53_] = True
                d_387_v57_: _dafny.Map
                d_387_v57_ = _dafny.Map({p1: (d_386_v56_).intersection(_dafny.Set({len(d_376_v46_)}))})
                index54_ = default__.safeIndex(102, ((self).f27).length(0))
                rhs53_ = ((d_387_v57_)[p1] if (p1) in (d_387_v57_) else d_386_v56_)
                rhs54_ = (self).f39
                rhs55_ = False
                lhs44_ = (self).f27
                lhs45_ = default__.safeIndex(102, ((self).f27).length(0))
                d_386_v56_ = rhs53_
                lhs44_[lhs45_] = rhs54_
                r1 = rhs55_
            elif True:
                d_388_v58_: _dafny.Seq
                d_388_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
                d_388_v58_ = (d_388_v58_) + (d_388_v58_)
                d_389_v59_: D8
                d_389_v59_ = D8_DC18(default__.safeModuloInt(p1, p1))
                d_389_v59_ = d_389_v59_
                (globalState).f17 = default__.fm40(globalState)
                (globalState).f5 = p1
                d_390_v60_: _dafny.Map
                d_390_v60_ = _dafny.Map({337: default__.safeDivisionInt((d_347_v19_).cardinality, p1)})
                d_391_v61_: D11
                d_391_v61_ = D11_DC28()
                d_392_v62_: T1
                nw51_ = C0()
                nw51_.ctor__(d_388_v58_, p1, (self).f27, p2)
                d_392_v62_ = nw51_
                d_393_v63_: _dafny.Map
                d_393_v63_ = _dafny.Map({d_392_v62_: d_392_v62_.f29})
                d_394_v64_: _dafny.Map
                d_394_v64_ = _dafny.Map({d_393_v63_: p2})
                d_395_v65_: str
                d_395_v65_ = _dafny.CodePoint('c')
                d_396_v66_: D0
                d_396_v66_ = D0_DC0(d_394_v64_, (self).f39, -710, d_395_v65_, (self).f28)
                index55_ = default__.safeIndex(56, ((self).f27).length(0))
                ((self).f27)[index55_] = (d_396_v66_).cf1
                index56_ = default__.safeIndex(56, ((self).f27).length(0))
                rhs56_ = d_390_v60_
                rhs57_ = (self).f28
                rhs58_ = p1
                rhs59_ = default__.fm41(d_388_v58_, globalState)
                rhs60_ = (self).f39
                lhs46_ = globalState
                lhs47_ = globalState
                lhs48_ = (self).f27
                lhs49_ = default__.safeIndex(56, ((self).f27).length(0))
                d_390_v60_ = rhs56_
                lhs46_.f17 = rhs57_
                lhs47_.f4 = rhs58_
                d_391_v61_ = rhs59_
                lhs48_[lhs49_] = rhs60_
            d_328_v2_ = _dafny.Map({(self).f28: p1})
        d_397_v67_: _dafny.Map
        d_397_v67_ = _dafny.Map({default__.fm42(p1, globalState): p1})
        d_398_v68_: str
        d_398_v68_ = _dafny.CodePoint('r')
        d_399_v69_: _dafny.Seq
        d_399_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpjwj"))
        hi1_ = ((d_397_v67_)[d_398_v68_] if (d_398_v68_) in (d_397_v67_) else len(d_399_v69_))
        for d_400_i2_ in range(p1, hi1_):
            d_401_v70_: _dafny.MultiSet
            d_401_v70_ = _dafny.MultiSet([d_400_i2_])
            d_402_v71_: _dafny.Seq
            d_402_v71_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oom"))])
            d_403_v72_: _dafny.Map
            d_403_v72_ = _dafny.Map({False: d_398_v68_})
            r0 = ((d_401_v70_) - (_dafny.MultiSet([d_400_i2_, p1, d_400_i2_, -612, p1]))).intersection(_dafny.MultiSet([len(d_402_v71_), len(_dafny.Map({(self).f39: (0) - (p1)})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "my"))), d_400_i2_, len(d_403_v72_)]))
            d_404_v73_: _dafny.Array
            def lambda10_(d_405_i3_):
                return True

            init5_ = lambda10_
            nw52_ = _dafny.Array(None, 16)
            for i0_5_ in range(nw52_.length(0)):
                nw52_[i0_5_] = init5_(i0_5_)
            d_404_v73_ = nw52_
            d_404_v73_ = (self).f27
            d_406_v74_: _dafny.Seq
            d_406_v74_ = _dafny.SeqWithoutIsStrInference([d_400_i2_, 403])
            r1 = ((len(p0) if (self).f39 else d_400_i2_)) < (len((d_406_v74_) + (d_406_v74_)))
            (globalState).f21 = 935
        d_407_v75_: D1
        d_407_v75_ = D1_DC7(d_398_v68_, d_399_v69_)
        d_408_v76_: _dafny.Array
        nw53_ = _dafny.Array(None, 18)
        nw53_[int(0)] = d_398_v68_
        nw53_[int(1)] = (d_407_v75_).cf13
        nw53_[int(2)] = d_398_v68_
        nw53_[int(3)] = default__.fm42(p1, globalState)
        nw53_[int(4)] = _dafny.CodePoint('k')
        nw53_[int(5)] = d_398_v68_
        nw53_[int(6)] = _dafny.CodePoint('g')
        nw53_[int(7)] = _dafny.CodePoint('y')
        nw53_[int(8)] = d_398_v68_
        nw53_[int(9)] = d_398_v68_
        nw53_[int(10)] = d_398_v68_
        nw53_[int(11)] = d_398_v68_
        nw53_[int(12)] = d_398_v68_
        nw53_[int(13)] = d_398_v68_
        nw53_[int(14)] = d_398_v68_
        nw53_[int(15)] = d_398_v68_
        nw53_[int(16)] = d_398_v68_
        nw53_[int(17)] = d_398_v68_
        d_408_v76_ = nw53_
        d_409_v77_: _dafny.Map
        d_409_v77_ = _dafny.Map({d_408_v76_: (d_399_v69_ if (self).f39 else d_399_v69_)})
        d_409_v77_ = d_409_v77_
        d_410_v78_: _dafny.MultiSet
        d_410_v78_ = _dafny.MultiSet([default__.safeDivisionInt(487, p1)])
        r0 = d_410_v78_
        r1 = p2
        return r0, r1

    @property
    def f39(self):
        return self._f39

class C2(T0):
    def  __init__(self):
        self._f27: _dafny.Array = _dafny.Array(None, 0)
        self._f28: bool = False
        self.f38: bool = False
        self._f37: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f37, f38, f27, f28):
        (self)._f37 = f37
        (self).f38 = f38
        (self)._f27 = f27
        (self)._f28 = f28

    def fm19(self, globalState):
        def iife58_():
            coll36_ = _dafny.Map()
            compr_36_: str
            for compr_36_ in (_dafny.Set({(self).f37, (self).f37})).Elements:
                d_411_v0_: str = compr_36_
                if (d_411_v0_) in (_dafny.Set({(self).f37, (self).f37})):
                    coll36_[d_411_v0_] = len(_dafny.Map({(self).f28: len(_dafny.SeqWithoutIsStrInference([self.f38, self.f38]))}))
            return _dafny.Map(coll36_)
        return (_dafny.Map({(self).f28: (0) - ((_dafny.MultiSet([self.f38])).cardinality)})) | ((_dafny.Map({True: len(_dafny.Set({(self).f28, self.f38}))})) | (_dafny.Map({self.f38: len(iife58_()
        )})))

    def fm20(self, globalState):
        return (_dafny.MultiSet([not (False) or (self.f38), (self).f28])).cardinality

    def m12(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        if self.f38:
            d_412_v0_: _dafny.Map
            d_412_v0_ = _dafny.Map({p0: self.f38})
            d_412_v0_ = (d_412_v0_).set((0) - ((p0) - (p0)), (self).f28)
            d_413_v1_: _dafny.Array
            nw54_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_413_v1_ = nw54_
            d_414_v2_: _dafny.Seq
            d_414_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnt"))
            d_415_v3_: T1
            nw55_ = C0()
            nw55_.ctor__(d_414_v2_, p3, (self).f27, not(self.f38))
            d_415_v3_ = nw55_
            d_416_v4_: _dafny.Seq
            d_416_v4_ = _dafny.SeqWithoutIsStrInference([d_415_v3_])
            d_417_v5_: _dafny.Map
            d_417_v5_ = _dafny.Map({d_413_v1_: (d_416_v4_) + (d_416_v4_)})
            d_417_v5_ = (d_417_v5_).set(d_413_v1_, d_416_v4_)
            (globalState).f17 = (d_415_v3_).f28
            (globalState).f5 = (len(d_414_v2_)) + (d_415_v3_.f29)
            d_414_v2_ = d_414_v2_
        elif True:
            d_418_v6_: _dafny.Map
            d_418_v6_ = _dafny.Map({p0: (self).f28})
            d_419_v7_: _dafny.Seq
            d_419_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gntq"))
            d_420_v8_: _dafny.Seq
            d_420_v8_ = _dafny.SeqWithoutIsStrInference([(0) - ((0) - (p0)), p0])
            d_421_v9_: _dafny.Array
            nw56_ = _dafny.Array(None, 18)
            nw56_[int(0)] = len((d_418_v6_) | (d_418_v6_))
            nw56_[int(1)] = 607
            nw56_[int(2)] = p3
            nw56_[int(3)] = (-478) * (p0)
            nw56_[int(4)] = (p3) + (p0)
            nw56_[int(5)] = p0
            nw56_[int(6)] = 68
            nw56_[int(7)] = 101
            nw56_[int(8)] = 996
            nw56_[int(9)] = len(d_419_v7_)
            nw56_[int(10)] = p3
            nw56_[int(11)] = (p0 if self.f38 else p0)
            nw56_[int(12)] = 43
            nw56_[int(13)] = p3
            nw56_[int(14)] = (self).fm20(globalState)
            nw56_[int(15)] = p3
            nw56_[int(16)] = p3
            nw56_[int(17)] = len(d_420_v8_)
            d_421_v9_ = nw56_
            d_421_v9_ = d_421_v9_
            (globalState).f21 = p0
            d_422_v10_: D1
            d_422_v10_ = D1_DC7(p1, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkuwjpj"))) + (d_419_v7_))
            source7_ = d_422_v10_
            if source7_.is_DC6:
                d_423___mcc_h0_ = source7_.cf12
                d_424_cf12_ = d_423___mcc_h0_
                d_425_v11_: D0
                d_425_v11_ = D0_DC1()
                d_426_v12_: D0
                d_426_v12_ = D0_DC4(d_425_v11_)
                d_426_v12_ = (d_426_v12_ if (len(d_419_v7_)) >= (p3) else d_426_v12_)
                d_427_v13_: C0
                nw57_ = C0()
                nw57_.ctor__(_dafny.SeqWithoutIsStrInference([(self).f37 for d_428_i0_ in range(default__.abs(192))]), d_424_cf12_, (self).f27, self.f38)
                d_427_v13_ = nw57_
                d_429_v14_: _dafny.Seq
                d_429_v14_ = _dafny.SeqWithoutIsStrInference([(d_424_cf12_) < (p0), (self).f28, (d_420_v8_) < (d_420_v8_), (self).f28, default__.fm22(True, 979, globalState)])
                rhs61_ = True
                rhs62_ = (default__.fm21((self).f28, p3, globalState)).cf6
                rhs63_ = True
                rhs64_ = (len(_dafny.SeqWithoutIsStrInference([(self).f28]))) == (p0)
                rhs65_ = (d_429_v14_)[default__.safeIndex(p0, len(d_429_v14_))]
                lhs50_ = globalState
                lhs51_ = globalState
                lhs52_ = globalState
                lhs50_.f15 = rhs61_
                d_419_v7_ = rhs62_
                r0 = rhs63_
                lhs51_.f6 = rhs64_
                lhs52_.f17 = rhs65_
                index57_ = default__.safeIndex(249, (d_421_v9_).length(0))
                (d_421_v9_)[index57_] = ((0) - (d_424_cf12_)) * (d_424_cf12_)
                index58_ = default__.safeIndex(249, (d_421_v9_).length(0))
                (d_421_v9_)[index58_] = default__.safeDivisionInt((0) - (d_424_cf12_), (self).fm20(globalState))
            elif source7_.is_DC7:
                d_430___mcc_h1_ = source7_.cf13
                d_431___mcc_h2_ = source7_.cf14
                d_432_cf14_ = d_431___mcc_h2_
                d_433_cf13_ = d_430___mcc_h1_
                (globalState).f25 = p1
                (globalState).f21 = p3
                index59_ = default__.safeIndex(506, (d_421_v9_).length(0))
                (d_421_v9_)[index59_] = p0
                index60_ = default__.safeIndex(506, (d_421_v9_).length(0))
                (d_421_v9_)[index60_] = default__.safeDivisionInt(p3, 887)
                d_434_v15_: _dafny.Map
                d_434_v15_ = _dafny.Map({True: default__.fm23(globalState)})
                d_435_v16_: _dafny.Set
                d_435_v16_ = _dafny.Set({(self).f28, (self).f28})
                d_436_v17_: D1
                d_436_v17_ = D1_DC6(p3)
                d_434_v15_ = (d_434_v15_).set((d_435_v16_).ispropersubset(_dafny.Set({self.f38, self.f38})), d_436_v17_)
            elif True:
                d_437___mcc_h3_ = source7_.cf11
                d_438_cf11_ = d_437___mcc_h3_
                (globalState).f13 = (self).fm20(globalState)
                d_439_v18_: _dafny.Array
                nw58_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_439_v18_ = nw58_
                index61_ = default__.safeIndex(382, (d_439_v18_).length(0))
                (d_439_v18_)[index61_] = (self).f27
                index62_ = default__.safeIndex(382, (d_439_v18_).length(0))
                (d_439_v18_)[index62_] = (self).f27
                d_440_v19_: _dafny.MultiSet
                d_440_v19_ = _dafny.MultiSet([p0, 338, p3])
                (globalState).f5 = ((0) - ((p0) - (p0))) + ((d_440_v19_).cardinality)
                index63_ = default__.safeIndex(691, ((self).f27).length(0))
                ((self).f27)[index63_] = self.f38
                d_441_v20_: _dafny.Map
                d_441_v20_ = _dafny.Map({(self).f28: p3})
                d_442_v21_: _dafny.Map
                d_442_v21_ = _dafny.Map({len(d_441_v20_): p3})
                index64_ = default__.safeIndex(691, ((self).f27).length(0))
                ((self).f27)[index64_] = (default__.fm22(True, len(_dafny.SeqWithoutIsStrInference([self.f38, self.f38])), globalState) if (self).f28 else (len(d_442_v21_)) > (len(_dafny.Set({True}))))
            d_443_v22_: _dafny.Seq
            d_443_v22_ = _dafny.SeqWithoutIsStrInference([False, (self).f28, (self).f28, self.f38])
            rhs66_ = self.f38
            rhs67_ = p0
            rhs68_ = ((self).fm20(globalState)) >= (len(d_443_v22_))
            lhs53_ = globalState
            lhs54_ = globalState
            lhs55_ = globalState
            lhs53_.f6 = rhs66_
            lhs54_.f21 = rhs67_
            lhs55_.f17 = rhs68_
            d_444_v23_: _dafny.MultiSet
            d_444_v23_ = _dafny.MultiSet([d_421_v9_])
            d_444_v23_ = d_444_v23_
        if (self).f28:
            d_445_v24_: _dafny.Map
            d_445_v24_ = _dafny.Map({((self).f28 if not(True) else self.f38): p0})
            d_446_v25_: _dafny.Set
            d_446_v25_ = _dafny.Set({p1})
            d_445_v24_ = (d_445_v24_).set((d_446_v25_).ispropersubset(d_446_v25_), p3)
            d_447_v26_: _dafny.Array
            def lambda11_(d_448_i1_):
                return (d_448_i1_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chc"))))

            init6_ = lambda11_
            nw59_ = _dafny.Array(None, 22)
            for i0_6_ in range(nw59_.length(0)):
                nw59_[i0_6_] = init6_(i0_6_)
            d_447_v26_ = nw59_
            d_447_v26_ = d_447_v26_
            d_449_v27_: D0
            d_449_v27_ = D0_DC1()
            source8_ = d_449_v27_
            if source8_.is_DC0:
                d_450___mcc_h4_ = source8_.cf0
                d_451___mcc_h5_ = source8_.cf1
                d_452___mcc_h6_ = source8_.cf2
                d_453___mcc_h7_ = source8_.cf3
                d_454___mcc_h8_ = source8_.cf4
                d_455_cf4_ = d_454___mcc_h8_
                d_456_cf3_ = d_453___mcc_h7_
                d_457_cf2_ = d_452___mcc_h6_
                d_458_cf1_ = d_451___mcc_h5_
                d_459_cf0_ = d_450___mcc_h4_
                r0 = d_458_cf1_
                d_460_v28_: _dafny.Set
                d_460_v28_ = _dafny.Set({(-739) - (p3)})
                d_461_v29_: _dafny.Map
                d_461_v29_ = _dafny.Map({d_457_cf2_: 24})
                d_462_v30_: _dafny.Map
                d_462_v30_ = _dafny.Map({d_461_v29_: d_460_v28_})
                d_463_v31_: _dafny.MultiSet
                d_463_v31_ = _dafny.MultiSet([d_455_cf4_])
                d_464_v32_: _dafny.Seq
                d_464_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqvo"))
                d_465_v33_: _dafny.Map
                d_465_v33_ = _dafny.Map({default__.fm22(not(d_455_cf4_), p0, globalState): d_464_v32_})
                d_466_v34_: D0
                d_466_v34_ = D0_DC2(d_460_v28_, d_464_v32_)
                def iife59_():
                    coll37_ = _dafny.Set()
                    compr_37_: int
                    for compr_37_ in _dafny.IntegerRange(904, 560):
                        d_467_v35_: int = compr_37_
                        if ((904) <= (d_467_v35_)) and ((d_467_v35_) < (560)):
                            coll37_ = coll37_.union(_dafny.Set([default__.safeModuloInt(d_467_v35_, d_457_cf2_)]))
                    return _dafny.Set(coll37_)
                rhs69_ = (self).f28
                rhs70_ = _dafny.Set({(p3) - (p0), p3, (0) - (((d_463_v31_).set(d_455_cf4_, default__.abs(len(((d_465_v33_)[d_455_cf4_] if (d_455_cf4_) in (d_465_v33_) else (d_466_v34_).cf6))))).cardinality)})
                rhs71_ = ((d_462_v30_).set(d_461_v29_, d_460_v28_)).set(d_461_v29_, iife59_()
                )
                lhs56_ = globalState
                lhs56_.f15 = rhs69_
                d_460_v28_ = rhs70_
                d_462_v30_ = rhs71_
                (globalState).f21 = d_457_cf2_
                d_468_v36_: D1
                d_468_v36_ = D1_DC6(p0)
                (globalState).f6 = ((default__.fm24(globalState)) - (_dafny.Set({d_468_v36_, d_468_v36_, d_468_v36_, d_468_v36_}))) == (_dafny.Set({d_468_v36_}))
            elif source8_.is_DC1:
                d_469_v37_: _dafny.Seq
                d_469_v37_ = _dafny.SeqWithoutIsStrInference([p0])
                rhs72_ = self.f38
                rhs73_ = default__.fm22((self).f28, (0) - ((d_469_v37_)[default__.safeIndex(p0, len(d_469_v37_))]), globalState)
                lhs57_ = globalState
                lhs58_ = globalState
                lhs57_.f15 = rhs72_
                lhs58_.f6 = rhs73_
                d_470_v38_: _dafny.Map
                d_470_v38_ = _dafny.Map({(self).f28: (self).f28})
                d_471_v39_: D3
                d_471_v39_ = D3_DC10(self.f38, self.f38)
                d_472_v40_: _dafny.Set
                d_472_v40_ = _dafny.Set({d_471_v39_})
                d_473_v41_: _dafny.Set
                d_473_v41_ = _dafny.Set({d_472_v40_, d_472_v40_})
                d_474_v42_: _dafny.Map
                def iife60_(_pat_let11_0):
                    def iife61_(d_475_dt__update__tmp_h0_):
                        def iife62_(_pat_let12_0):
                            def iife63_(d_476_dt__update_hcf18_h0_):
                                return D3_DC10((d_475_dt__update__tmp_h0_).cf17, d_476_dt__update_hcf18_h0_)
                            return iife63_(_pat_let12_0)
                        return iife62_((self).f28)
                    return iife61_(_pat_let11_0)
                d_474_v42_ = _dafny.Map({((d_470_v38_)[self.f38] if (self.f38) in (d_470_v38_) else self.f38): (d_473_v41_ if self.f38 else _dafny.Set({d_472_v40_, _dafny.Set({D3_DC10((self).f28, self.f38)}), d_472_v40_, _dafny.Set({d_471_v39_, d_471_v39_, d_471_v39_, iife60_(D3_DC10(not(self.f38), (self).f28)), d_471_v39_}), _dafny.Set({d_471_v39_})}))})
                d_474_v42_ = (d_474_v42_).set((self).f28, (d_473_v41_) - (d_473_v41_))
                (globalState).f5 = default__.safeModuloInt(p3, p3)
                d_477_v43_: _dafny.Seq
                d_477_v43_ = _dafny.SeqWithoutIsStrInference([False, self.f38])
                d_477_v43_ = default__.fm25(p0, (default__.fm26((self).f28, p0, _dafny.CodePoint('n'), globalState)).cardinality, default__.safeDivisionInt(p0, p3), globalState)
            elif source8_.is_DC2:
                d_478___mcc_h9_ = source8_.cf5
                d_479___mcc_h10_ = source8_.cf6
                d_480_cf6_ = d_479___mcc_h10_
                d_481_cf5_ = d_478___mcc_h9_
                d_482_v44_: _dafny.Seq
                d_482_v44_ = _dafny.SeqWithoutIsStrInference([(self).f28, (self).f28])
                d_482_v44_ = (d_482_v44_ if True else default__.fm25(len(d_480_cf6_), 689, p0, globalState))
                index65_ = default__.safeIndex(103, ((self).f27).length(0))
                ((self).f27)[index65_] = (self).f28
                index66_ = default__.safeIndex(103, ((self).f27).length(0))
                ((self).f27)[index66_] = ((self).f28) and ((_dafny.Set({557})) != (d_481_cf5_))
                d_483_v45_: _dafny.Set
                d_483_v45_ = _dafny.Set({self.f38})
                d_484_v46_: _dafny.Seq
                d_484_v46_ = _dafny.SeqWithoutIsStrInference([d_483_v45_, d_483_v45_, (d_483_v45_) | (_dafny.Set({((self).f27)[default__.safeIndex(103, ((self).f27).length(0))], not(((self).f27)[default__.safeIndex(103, ((self).f27).length(0))]), self.f38}))])
                d_485_v47_: D5
                d_485_v47_ = D5_DC12(d_484_v46_)
                d_484_v46_ = (d_485_v47_).cf20
                (globalState).f5 = (0) - (p3)
            elif source8_.is_DC3:
                d_486___mcc_h11_ = source8_.cf7
                d_487___mcc_h12_ = source8_.cf8
                d_488___mcc_h13_ = source8_.cf9
                d_489_cf9_ = d_488___mcc_h13_
                d_490_cf8_ = d_487___mcc_h12_
                d_491_cf7_ = d_486___mcc_h11_
                d_492_v48_: _dafny.Set
                d_492_v48_ = _dafny.Set({d_491_cf7_, p3, p3})
                d_493_v49_: _dafny.Map
                d_493_v49_ = _dafny.Map({d_491_cf7_: d_492_v48_})
                rhs74_ = default__.fm22(self.f38, (0) - ((p0) - (p3)), globalState)
                rhs75_ = default__.fm22((default__.fm0(d_493_v49_, False, False, globalState)).isdisjoint(d_492_v48_), default__.safeDivisionInt(p0, d_491_cf7_), globalState)
                rhs76_ = self.f38
                lhs59_ = globalState
                lhs60_ = globalState
                lhs61_ = self
                lhs59_.f17 = rhs74_
                lhs60_.f17 = rhs75_
                lhs61_.f38 = rhs76_
                d_494_v50_: _dafny.Set
                d_494_v50_ = _dafny.Set({d_490_cf8_, d_489_cf9_})
                d_495_v51_: D0
                d_495_v51_ = D0_DC1()
                d_496_v52_: D0
                d_496_v52_ = D0_DC4(d_495_v51_)
                d_497_v53_: _dafny.Map
                d_497_v53_ = _dafny.Map({d_494_v50_: d_496_v52_})
                d_497_v53_ = (d_497_v53_).set(d_494_v50_, d_496_v52_)
                (globalState).f13 = (175 if d_490_cf8_ else p0)
                d_498_v54_: _dafny.Seq
                d_498_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
                d_499_v55_: T1
                nw60_ = C0()
                nw60_.ctor__(d_498_v54_, 735, (self).f27, (self).f28)
                d_499_v55_ = nw60_
                d_500_v56_: _dafny.Map
                d_500_v56_ = _dafny.Map({d_499_v55_: d_499_v55_.f29})
                d_501_v57_: _dafny.Map
                d_501_v57_ = _dafny.Map({d_500_v56_: not(self.f38)})
                d_502_v58_: D0
                d_502_v58_ = D0_DC0(d_501_v57_, d_489_cf9_, d_499_v55_.f29, (self).f37, (self).f28)
                (globalState).f17 = (d_502_v58_).cf1
            elif True:
                d_503___mcc_h14_ = source8_.cf10
                d_504_cf10_ = d_503___mcc_h14_
                d_505_v59_: D5
                d_505_v59_ = D5_DC13((self).f28, d_447_v26_, self.f38)
                d_506_v60_: _dafny.Set
                d_506_v60_ = _dafny.Set({d_505_v59_})
                d_507_v61_: _dafny.Array
                nw61_ = _dafny.Array(None, 10)
                nw61_[int(0)] = d_506_v60_
                nw61_[int(1)] = d_506_v60_
                nw61_[int(2)] = d_506_v60_
                nw61_[int(3)] = (d_506_v60_) | (d_506_v60_)
                nw61_[int(4)] = _dafny.Set({D5_DC13(self.f38, d_447_v26_, (self).f28)})
                nw61_[int(5)] = d_506_v60_
                nw61_[int(6)] = d_506_v60_
                nw61_[int(7)] = d_506_v60_
                nw61_[int(8)] = _dafny.Set({d_505_v59_, d_505_v59_})
                nw61_[int(9)] = (_dafny.Set({d_505_v59_, d_505_v59_, d_505_v59_})) | (d_506_v60_)
                d_507_v61_ = nw61_
                index67_ = default__.safeIndex(377, (d_507_v61_).length(0))
                (d_507_v61_)[index67_] = d_506_v60_
                d_508_v62_: _dafny.Map
                d_508_v62_ = _dafny.Map({p0: _dafny.MultiSet([(self).f27, (self).f27])})
                d_509_v63_: _dafny.MultiSet
                d_509_v63_ = _dafny.MultiSet([(self).f27])
                d_510_v64_: _dafny.Set
                d_510_v64_ = _dafny.Set({p0, p3})
                index68_ = default__.safeIndex(377, (d_507_v61_).length(0))
                rhs77_ = (((d_508_v62_)[p0] if (p0) in (d_508_v62_) else (d_509_v63_).set((self).f27, default__.abs(len(d_510_v64_))))).issubset(_dafny.MultiSet([(self).f27]))
                rhs78_ = _dafny.Set({d_505_v59_})
                rhs79_ = 19
                rhs80_ = (self).f37
                rhs81_ = (self).f28
                lhs62_ = globalState
                lhs63_ = d_507_v61_
                lhs64_ = default__.safeIndex(377, (d_507_v61_).length(0))
                lhs65_ = globalState
                lhs66_ = globalState
                lhs67_ = globalState
                lhs62_.f17 = rhs77_
                lhs63_[lhs64_] = rhs78_
                lhs65_.f5 = rhs79_
                lhs66_.f25 = rhs80_
                lhs67_.f17 = rhs81_
                d_511_v65_: _dafny.MultiSet
                d_511_v65_ = _dafny.MultiSet([p1])
                d_512_v66_: _dafny.Map
                d_512_v66_ = _dafny.Map({(d_511_v65_).cardinality: self.f38})
                d_513_v68_: _dafny.Array
                nw62_ = _dafny.Array(None, 5)
                nw62_[int(0)] = d_512_v66_
                def iife64_():
                    coll38_ = _dafny.Map()
                    compr_38_: int
                    for compr_38_ in _dafny.IntegerRange(236, 663):
                        d_514_v67_: int = compr_38_
                        if ((236) <= (d_514_v67_)) and ((d_514_v67_) < (663)):
                            coll38_[default__.safeModuloInt(d_514_v67_, p3)] = self.f38
                    return _dafny.Map(coll38_)
                nw62_[int(1)] = (iife64_()
                ) | (d_512_v66_)
                nw62_[int(2)] = d_512_v66_
                nw62_[int(3)] = _dafny.Map({p3: (self).f28})
                nw62_[int(4)] = d_512_v66_
                d_513_v68_ = nw62_
                d_515_v69_: _dafny.Seq
                d_515_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "osp"))
                d_516_v70_: _dafny.Map
                d_516_v70_ = _dafny.Map({p0: D1_DC7((self).f37, d_515_v69_)})
                d_517_v71_: _dafny.Seq
                d_517_v71_ = _dafny.SeqWithoutIsStrInference([d_516_v70_])
                def iife65_():
                    coll39_ = _dafny.Set()
                    compr_39_: int
                    for compr_39_ in _dafny.IntegerRange(637, 773):
                        d_518_v72_: int = compr_39_
                        if ((637) <= (d_518_v72_)) and ((d_518_v72_) < (773)):
                            coll39_ = coll39_.union(_dafny.Set([(d_518_v72_) + (p3)]))
                    return _dafny.Set(coll39_)
                rhs82_ = iife65_()

                rhs83_ = p0
                rhs84_ = (self).f28
                rhs85_ = d_513_v68_
                rhs86_ = _dafny.SeqWithoutIsStrInference([default__.fm27(self.f38, p3, self.f38, globalState)])
                lhs68_ = globalState
                lhs69_ = globalState
                d_510_v64_ = rhs82_
                lhs68_.f21 = rhs83_
                lhs69_.f15 = rhs84_
                d_513_v68_ = rhs85_
                d_517_v71_ = rhs86_
                d_515_v69_ = d_515_v69_
                d_519_v73_: D0
                d_519_v73_ = D0_DC2(d_510_v64_, d_515_v69_)
                d_520_v74_: _dafny.Set
                d_520_v74_ = _dafny.Set({self.f38, ((self).f28) and ((self).f28), (len((d_519_v73_).cf6)) == (p0), (self).f28, not(False)})
                d_520_v74_ = d_520_v74_
            (globalState).f21 = (0) - ((0) - (p3))
            d_521_v75_: _dafny.Map
            d_522_v76_: int
            out6_: _dafny.Map
            out7_: int
            out6_, out7_ = default__.m0(globalState)
            d_521_v75_ = out6_
            d_522_v76_ = out7_
        elif True:
            if (self).f28:
                d_523_v77_: _dafny.Seq
                d_523_v77_ = _dafny.SeqWithoutIsStrInference([self.f38])
                d_524_v78_: _dafny.MultiSet
                d_524_v78_ = _dafny.MultiSet([((0) - (p3)) == (p3), (d_523_v77_)[default__.safeIndex(p0, len(d_523_v77_))], self.f38])
                d_525_v79_: _dafny.Seq
                d_525_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gio"))
                d_526_v80_: _dafny.Array
                nw63_ = _dafny.Array(int(0), 8)
                d_526_v80_ = nw63_
                d_527_v81_: _dafny.Map
                d_527_v81_ = _dafny.Map({(self.f38) or ((d_523_v77_)[default__.safeIndex(p3, len(d_523_v77_))]): (d_525_v79_).set(default__.safeIndex(474, len(d_525_v79_)), (self).f37)})
                index69_ = default__.safeIndex(422, (d_526_v80_).length(0))
                (d_526_v80_)[index69_] = p0
                d_528_v82_: _dafny.Set
                d_528_v82_ = _dafny.Set({(self).fm20(globalState)})
                d_529_v83_: _dafny.Seq
                d_529_v83_ = _dafny.SeqWithoutIsStrInference([d_528_v82_, d_528_v82_, _dafny.Set({p0}), d_528_v82_, d_528_v82_])
                index70_ = default__.safeIndex(422, (d_526_v80_).length(0))
                rhs87_ = (d_524_v78_) - ((d_524_v78_).intersection(default__.fm26((d_523_v77_)[default__.safeIndex(p3, len(d_523_v77_))], p0, p1, globalState)))
                rhs88_ = d_525_v79_
                rhs89_ = d_526_v80_
                rhs90_ = _dafny.Map({(self).f28: d_525_v79_})
                rhs91_ = (len((d_529_v83_) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({p0})])))) + (p0)
                lhs70_ = d_526_v80_
                lhs71_ = default__.safeIndex(422, (d_526_v80_).length(0))
                d_524_v78_ = rhs87_
                d_525_v79_ = rhs88_
                d_526_v80_ = rhs89_
                d_527_v81_ = rhs90_
                lhs70_[lhs71_] = rhs91_
                d_530_v84_: C0
                nw64_ = C0()
                nw64_.ctor__(d_525_v79_, (0) - (-36), (self).f27, (self).f28)
                d_530_v84_ = nw64_
                (globalState).f21 = (d_526_v80_)[default__.safeIndex(422, (d_526_v80_).length(0))]
                d_531_v85_: _dafny.Array
                d_531_v85_ = (self).f27
                d_532_v86_: C0
                nw65_ = C0()
                nw65_.ctor__((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fervo"))) + ((d_530_v84_).f36), (self).fm20(globalState), (d_531_v85_), False)
                d_532_v86_ = nw65_
                index71_ = default__.safeIndex(241, ((self).f27).length(0))
                ((self).f27)[index71_] = (self).f28
                d_533_v87_: _dafny.Map
                d_533_v87_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vo"))})
                d_534_v89_: _dafny.Map
                d_534_v89_ = _dafny.Map({p3: d_533_v87_})
                d_535_v90_: _dafny.Seq
                d_535_v90_ = _dafny.SeqWithoutIsStrInference([default__.fm28(p0, globalState)])
                d_536_v92_: _dafny.Array
                nw66_ = _dafny.Array(None, 15)
                nw66_[int(0)] = d_533_v87_
                nw66_[int(1)] = d_533_v87_
                nw66_[int(2)] = _dafny.Map({(d_526_v80_)[default__.safeIndex(422, (d_526_v80_).length(0))]: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "je"))})
                nw66_[int(3)] = d_533_v87_
                def iife66_():
                    coll40_ = _dafny.Map()
                    compr_40_: int
                    for compr_40_ in _dafny.IntegerRange(547, -19):
                        d_537_v88_: int = compr_40_
                        if ((547) <= (d_537_v88_)) and ((d_537_v88_) < (-19)):
                            coll40_[default__.safeModuloInt(d_537_v88_, 510)] = d_525_v79_
                    return _dafny.Map(coll40_)
                nw66_[int(4)] = iife66_()
                
                nw66_[int(5)] = ((d_534_v89_)[p3] if (p3) in (d_534_v89_) else (d_535_v90_)[default__.safeIndex((d_526_v80_)[default__.safeIndex(422, (d_526_v80_).length(0))], len(d_535_v90_))])
                nw66_[int(6)] = d_533_v87_
                nw66_[int(7)] = (d_533_v87_) | ((D7_DC15(d_533_v87_)).cf25)
                nw66_[int(8)] = default__.fm28(len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({self.f38: self.f38})), p3, p3])), globalState)
                nw66_[int(9)] = (d_533_v87_) | (d_533_v87_)
                def iife67_():
                    coll41_ = _dafny.Map()
                    compr_41_: int
                    for compr_41_ in _dafny.IntegerRange(568, 385):
                        d_538_v91_: int = compr_41_
                        if ((568) <= (d_538_v91_)) and ((d_538_v91_) < (385)):
                            coll41_[(d_538_v91_) + (p0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnogundpk"))
                    return _dafny.Map(coll41_)
                nw66_[int(10)] = iife67_()
                
                nw66_[int(11)] = d_533_v87_
                nw66_[int(12)] = (d_533_v87_).set(p0, d_525_v79_)
                nw66_[int(13)] = (d_533_v87_ if (self).f28 else d_533_v87_)
                nw66_[int(14)] = d_533_v87_
                d_536_v92_ = nw66_
                index72_ = default__.safeIndex(948, (d_536_v92_).length(0))
                (d_536_v92_)[index72_] = d_533_v87_
                index73_ = default__.safeIndex(241, ((self).f27).length(0))
                index74_ = default__.safeIndex(948, (d_536_v92_).length(0))
                rhs92_ = self.f38
                rhs93_ = default__.safeDivisionInt((0) - (p0), (d_526_v80_)[default__.safeIndex(422, (d_526_v80_).length(0))])
                rhs94_ = d_533_v87_
                rhs95_ = not ((self).f28) or (self.f38)
                lhs72_ = (self).f27
                lhs73_ = default__.safeIndex(241, ((self).f27).length(0))
                lhs74_ = globalState
                lhs75_ = d_536_v92_
                lhs76_ = default__.safeIndex(948, (d_536_v92_).length(0))
                lhs77_ = globalState
                lhs72_[lhs73_] = rhs92_
                lhs74_.f13 = rhs93_
                lhs75_[lhs76_] = rhs94_
                lhs77_.f15 = rhs95_
            elif True:
                d_539_v93_: D0
                d_539_v93_ = D0_DC3(p3, (self).f28, (self).f28)
                d_540_v94_: D0
                d_540_v94_ = D0_DC4(d_539_v93_)
                d_541_v95_: _dafny.MultiSet
                d_541_v95_ = _dafny.MultiSet([d_540_v94_])
                d_542_v96_: _dafny.Map
                d_542_v96_ = _dafny.Map({(self).f28: p0})
                rhs96_ = (len(d_542_v96_) if (self).f28 else p3)
                rhs97_ = self.f38
                rhs98_ = _dafny.MultiSet([d_540_v94_, d_540_v94_])
                rhs99_ = self.f38
                rhs100_ = (0) - (p3)
                lhs78_ = globalState
                lhs79_ = globalState
                lhs80_ = globalState
                lhs81_ = globalState
                lhs78_.f4 = rhs96_
                lhs79_.f17 = rhs97_
                d_541_v95_ = rhs98_
                lhs80_.f15 = rhs99_
                lhs81_.f13 = rhs100_
                d_543_v97_: _dafny.Array
                nw67_ = _dafny.Array(_dafny.Seq({}), 4)
                d_543_v97_ = nw67_
                d_544_v98_: _dafny.Map
                d_544_v98_ = _dafny.Map({p0: p0})
                d_545_v99_: _dafny.Seq
                d_545_v99_ = _dafny.SeqWithoutIsStrInference([d_544_v98_])
                d_546_v100_: _dafny.Map
                d_546_v100_ = _dafny.Map({p0: d_545_v99_})
                index75_ = default__.safeIndex(275, (d_543_v97_).length(0))
                (d_543_v97_)[index75_] = ((d_546_v100_)[p0] if (p0) in (d_546_v100_) else _dafny.SeqWithoutIsStrInference([d_544_v98_ for d_547_i2_ in range(default__.abs(468))]))
                index76_ = default__.safeIndex(275, (d_543_v97_).length(0))
                (d_543_v97_)[index76_] = (d_545_v99_) + (d_545_v99_)
                d_548_v101_: _dafny.Seq
                d_548_v101_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjcirftt"))
                d_549_v102_: D0
                out8_: D0
                out8_ = (self).m13(self.f38, (self).f28, len(default__.fm29(self.f38, True, globalState)), d_548_v101_, globalState)
                d_549_v102_ = out8_
                d_550_v103_: _dafny.Map
                d_551_v104_: int
                out9_: _dafny.Map
                out10_: int
                out9_, out10_ = default__.m0(globalState)
                d_550_v103_ = out9_
                d_551_v104_ = out10_
                d_552_v105_: C0
                nw68_ = C0()
                nw68_.ctor__((d_548_v101_) + (_dafny.SeqWithoutIsStrInference([(self).f37 for d_553_i3_ in range(default__.abs(628))])), p3, ((self).f27 if not((self).f28) else (self).f27), default__.fm22(False, p0, globalState))
                d_552_v105_ = nw68_
            d_554_v106_: _dafny.Seq
            d_554_v106_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nke"))
            d_555_v107_: C0
            nw69_ = C0()
            nw69_.ctor__(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unybyg"))) + (d_554_v106_)).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unybyg"))) + (d_554_v106_))), (self).f37), len(_dafny.SeqWithoutIsStrInference([(self).fm20(globalState)])), (self).f27, (self).f28)
            d_555_v107_ = nw69_
            r0 = default__.fm22(self.f38, p3, globalState)
            d_556_v108_: _dafny.Map
            d_556_v108_ = _dafny.Map({p0: len(_dafny.Set({not((self).f28)}))})
            (globalState).f4 = (0) - (((d_556_v108_)[p3] if (p3) in (d_556_v108_) else p0))
            if not(default__.fm22(self.f38, p3, globalState)):
                (globalState).f4 = ((self).fm20(globalState)) * (p3)
                (globalState).f17 = (self).f28
                d_557_v109_: D0
                d_557_v109_ = D0_DC3((self).fm20(globalState), self.f38, (self).f28)
                d_558_v110_: _dafny.Map
                d_558_v110_ = _dafny.Map({(d_557_v109_).cf7: (self).f28})
                d_558_v110_ = (d_558_v110_) | (default__.fm30(self.f38, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")), self.f38, globalState))
                d_559_v111_: _dafny.Map
                d_559_v111_ = _dafny.Map({self.f38: not(self.f38)})
                r0 = ((d_559_v111_)[self.f38] if (self.f38) in (d_559_v111_) else (p3) >= (p3))
                d_560_v112_: _dafny.Seq
                d_560_v112_ = _dafny.SeqWithoutIsStrInference([(not((self).f28)) or (self.f38), self.f38, (not(True)) or (not(not((self).f28)))])
                d_561_v113_: _dafny.Seq
                d_561_v113_ = _dafny.SeqWithoutIsStrInference([d_560_v112_, default__.fm25(p0, (self).fm20(globalState), p3, globalState)])
                d_560_v112_ = ((d_561_v113_)[default__.safeIndex(p0, len(d_561_v113_))]) + (_dafny.SeqWithoutIsStrInference([self.f38, (d_560_v112_)[default__.safeIndex(p3, len(d_560_v112_))], not(self.f38)]))
            elif True:
                d_562_v114_: _dafny.Map
                d_562_v114_ = _dafny.Map({p3: not((self).f28)})
                d_563_v115_: _dafny.MultiSet
                d_563_v115_ = _dafny.MultiSet([p0, p3])
                d_562_v114_ = (d_562_v114_).set(p0, ((d_563_v115_).set(p3, default__.abs(p3))).isdisjoint(d_563_v115_))
                d_564_v116_: _dafny.Array
                nw70_ = _dafny.Array(int(0), 10)
                d_564_v116_ = nw70_
                d_565_v117_: _dafny.MultiSet
                d_565_v117_ = _dafny.MultiSet([d_564_v116_, d_564_v116_, d_564_v116_])
                index77_ = default__.safeIndex(426, ((self).f27).length(0))
                ((self).f27)[index77_] = (d_564_v116_) not in (d_565_v117_)
                index78_ = default__.safeIndex(426, ((self).f27).length(0))
                ((self).f27)[index78_] = self.f38
                (globalState).f4 = (default__.safeModuloInt(p0, p0)) - (p3)
                d_566_v118_: _dafny.Seq
                d_566_v118_ = _dafny.SeqWithoutIsStrInference([((self).f27)[default__.safeIndex(426, ((self).f27).length(0))], (self).f28])
                d_567_v119_: _dafny.Array
                nw71_ = _dafny.Array(None, 28)
                nw71_[int(0)] = p1
                nw71_[int(1)] = p1
                nw71_[int(2)] = p1
                nw71_[int(3)] = (self).f37
                nw71_[int(4)] = (self).f37
                nw71_[int(5)] = (self).f37
                nw71_[int(6)] = (self).f37
                nw71_[int(7)] = _dafny.CodePoint('h')
                nw71_[int(8)] = p1
                nw71_[int(9)] = (self).f37
                nw71_[int(10)] = p1
                nw71_[int(11)] = (self).f37
                nw71_[int(12)] = (self).f37
                nw71_[int(13)] = (self).f37
                nw71_[int(14)] = ((self).f37 if self.f38 else (self).f37)
                nw71_[int(15)] = p1
                nw71_[int(16)] = ((d_555_v107_).f36)[default__.safeIndex(len(d_566_v118_), len((d_555_v107_).f36))]
                nw71_[int(17)] = (self).f37
                nw71_[int(18)] = p1
                nw71_[int(19)] = (self).f37
                nw71_[int(20)] = _dafny.CodePoint('y')
                nw71_[int(21)] = p1
                nw71_[int(22)] = (self).f37
                nw71_[int(23)] = (self).f37
                nw71_[int(24)] = (self).f37
                nw71_[int(25)] = (self).f37
                nw71_[int(26)] = (self).f37
                nw71_[int(27)] = (self).f37
                d_567_v119_ = nw71_
                d_567_v119_ = d_567_v119_
                (globalState).f4 = ((self).fm20(globalState) if (self).f28 else p0)
        d_568_v120_: D0
        d_568_v120_ = D0_DC3(p3, not((self).f28), (self).f28)
        d_569_v121_: _dafny.MultiSet
        d_569_v121_ = _dafny.MultiSet([(d_568_v120_).cf8, self.f38])
        d_570_v122_: _dafny.Map
        d_570_v122_ = _dafny.Map({(d_569_v121_).cardinality: (0) - (p0)})
        d_570_v122_ = ((d_570_v122_) | (d_570_v122_)) | (_dafny.Map({(self).fm20(globalState): p0}))
        if self.f38:
            d_571_v123_: _dafny.Seq
            d_571_v123_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkegrqm"))
            (globalState).f21 = default__.safeDivisionInt(314, len(d_571_v123_))
            d_572_v124_: _dafny.Seq
            d_572_v124_ = _dafny.SeqWithoutIsStrInference([(self).fm20(globalState)])
            (globalState).f13 = (p0) - (len(d_572_v124_))
            d_573_v125_: _dafny.Array
            nw72_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
            d_573_v125_ = nw72_
            d_573_v125_ = p2
            d_574_v126_: _dafny.Seq
            d_574_v126_ = _dafny.SeqWithoutIsStrInference([(self).f28])
            d_575_v127_: _dafny.Seq
            d_575_v127_ = _dafny.SeqWithoutIsStrInference([d_574_v126_])
            d_576_v128_: _dafny.Map
            d_576_v128_ = _dafny.Map({(d_575_v127_) + (d_575_v127_): (p3) * (len(d_571_v123_))})
            d_577_v129_: _dafny.Map
            d_577_v129_ = _dafny.Map({(self).f28: d_575_v127_})
            d_578_v130_: _dafny.Array
            def lambda12_(d_579_i4_):
                return (d_579_i4_) + (199)

            init7_ = lambda12_
            nw73_ = _dafny.Array(None, 25)
            for i0_7_ in range(nw73_.length(0)):
                nw73_[i0_7_] = init7_(i0_7_)
            d_578_v130_ = nw73_
            d_580_v131_: _dafny.MultiSet
            d_580_v131_ = _dafny.MultiSet([d_578_v130_, d_578_v130_, d_578_v130_])
            d_581_v132_: _dafny.Seq
            d_581_v132_ = _dafny.SeqWithoutIsStrInference([d_580_v131_])
            d_576_v128_ = (d_576_v128_).set((d_575_v127_) + (((d_577_v129_)[self.f38] if (self.f38) in (d_577_v129_) else _dafny.SeqWithoutIsStrInference([d_574_v126_, d_574_v126_]))), ((_dafny.MultiSet([d_578_v130_, d_578_v130_])) | ((d_581_v132_)[default__.safeIndex(len(d_571_v123_), len(d_581_v132_))])).cardinality)
            index79_ = default__.safeIndex(824, (d_578_v130_).length(0))
            (d_578_v130_)[index79_] = p3
            index80_ = default__.safeIndex(824, (d_578_v130_).length(0))
            (d_578_v130_)[index80_] = default__.safeDivisionInt(p3, -240)
        elif True:
            d_582_v133_: _dafny.Array
            def lambda13_(d_583_p0_):
                def lambda14_(d_584_i5_):
                    return (d_584_i5_) * (d_583_p0_)

                return lambda14_

            init8_ = lambda13_(p0)
            nw74_ = _dafny.Array(None, 14)
            for i0_8_ in range(nw74_.length(0)):
                nw74_[i0_8_] = init8_(i0_8_)
            d_582_v133_ = nw74_
            d_582_v133_ = d_582_v133_
            d_585_v134_: _dafny.Array
            nw75_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
            d_585_v134_ = nw75_
            d_585_v134_ = d_585_v134_
            d_586_v135_: _dafny.Array
            nw76_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_586_v135_ = nw76_
            d_586_v135_ = d_586_v135_
            (globalState).f4 = p0
            d_587_v136_: _dafny.Seq
            d_587_v136_ = _dafny.SeqWithoutIsStrInference([self.f38])
            if (len(d_587_v136_)) != (p3):
                d_588_v137_: _dafny.Map
                d_588_v137_ = _dafny.Map({True: not ((self).f28) or (self.f38)})
                d_588_v137_ = (d_588_v137_).set(False, (self).f28)
                d_589_v138_: _dafny.Array
                nw77_ = _dafny.Array(None, 7)
                nw77_[int(0)] = (d_587_v136_ if default__.fm22((self).f28, p0, globalState) else d_587_v136_)
                nw77_[int(1)] = _dafny.SeqWithoutIsStrInference([self.f38, (self).f28])
                nw77_[int(2)] = (d_587_v136_) + (d_587_v136_)
                nw77_[int(3)] = _dafny.SeqWithoutIsStrInference([self.f38, self.f38])
                nw77_[int(4)] = d_587_v136_
                nw77_[int(5)] = d_587_v136_
                nw77_[int(6)] = d_587_v136_
                d_589_v138_ = nw77_
                d_590_v139_: _dafny.Seq
                d_590_v139_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkbyosc"))
                nw78_ = _dafny.Array(None, 15)
                nw78_[int(0)] = d_587_v136_
                nw78_[int(1)] = _dafny.SeqWithoutIsStrInference([self.f38])
                nw78_[int(2)] = d_587_v136_
                nw78_[int(3)] = d_587_v136_
                nw78_[int(4)] = default__.fm25((0) - (p3), (0) - (p0), p0, globalState)
                nw78_[int(5)] = d_587_v136_
                nw78_[int(6)] = ((d_587_v136_) + (d_587_v136_)).set(default__.safeIndex(len(d_590_v139_), len((d_587_v136_) + (d_587_v136_))), (self).f28)
                nw78_[int(7)] = d_587_v136_
                nw78_[int(8)] = (d_587_v136_) + (d_587_v136_)
                nw78_[int(9)] = d_587_v136_
                nw78_[int(10)] = d_587_v136_
                nw78_[int(11)] = d_587_v136_
                nw78_[int(12)] = (d_587_v136_) + (d_587_v136_)
                nw78_[int(13)] = d_587_v136_
                nw78_[int(14)] = (_dafny.SeqWithoutIsStrInference([not(False), (self).f28, (self).f28])) + (_dafny.SeqWithoutIsStrInference([self.f38, False]))
                d_589_v138_ = nw78_
                d_591_v140_: _dafny.MultiSet
                d_591_v140_ = _dafny.MultiSet([p0, p3])
                index81_ = default__.safeIndex(580, (d_582_v133_).length(0))
                (d_582_v133_)[index81_] = (d_591_v140_).cardinality
                d_592_v141_: _dafny.Array
                nw79_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_592_v141_ = nw79_
                index82_ = default__.safeIndex(466, (d_592_v141_).length(0))
                (d_592_v141_)[index82_] = d_582_v133_
                index83_ = default__.safeIndex(580, (d_582_v133_).length(0))
                index84_ = default__.safeIndex(466, (d_592_v141_).length(0))
                rhs101_ = p0
                rhs102_ = p3
                rhs103_ = p3
                rhs104_ = len(((_dafny.SeqWithoutIsStrInference([(self).f28, (d_587_v136_)[default__.safeIndex(958, len(d_587_v136_))], not(default__.fm22((self).f28, p0, globalState))])) + (d_587_v136_) if default__.fm22((self).f28, p0, globalState) else d_587_v136_))
                rhs105_ = d_582_v133_
                lhs82_ = d_582_v133_
                lhs83_ = default__.safeIndex(580, (d_582_v133_).length(0))
                lhs84_ = globalState
                lhs85_ = globalState
                lhs86_ = globalState
                lhs87_ = d_592_v141_
                lhs88_ = default__.safeIndex(466, (d_592_v141_).length(0))
                lhs82_[lhs83_] = rhs101_
                lhs84_.f13 = rhs102_
                lhs85_.f21 = rhs103_
                lhs86_.f5 = rhs104_
                lhs87_[lhs88_] = rhs105_
                (globalState).f5 = (d_582_v133_)[default__.safeIndex(580, (d_582_v133_).length(0))]
                d_593_v142_: T1
                nw80_ = C0()
                nw80_.ctor__(d_590_v139_, (d_582_v133_)[default__.safeIndex(580, (d_582_v133_).length(0))], (self).f27, (self).f28)
                d_593_v142_ = nw80_
                d_594_v143_: _dafny.Map
                d_594_v143_ = _dafny.Map({d_593_v142_: 542})
                d_595_v144_: _dafny.Map
                d_595_v144_ = _dafny.Map({d_594_v143_: self.f38})
                d_596_v145_: _dafny.Seq
                d_596_v145_ = _dafny.SeqWithoutIsStrInference([p3])
                d_597_v146_: D0
                d_597_v146_ = D0_DC4(D0_DC0(d_595_v144_, (self).f28, (d_596_v145_)[default__.safeIndex(128, len(d_596_v145_))], (self).f37, (d_593_v142_).f28))
                rhs106_ = p3
                rhs107_ = d_597_v146_
                lhs89_ = globalState
                lhs89_.f5 = rhs106_
                d_597_v146_ = rhs107_
            elif True:
                d_598_v147_: _dafny.Map
                d_598_v147_ = _dafny.Map({not((self).f28): (self).f28})
                (globalState).f21 = default__.safeModuloInt(len(d_598_v147_), p0)
                d_599_v148_: _dafny.Seq
                d_599_v148_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27, (self).f27, (self).f27])
                d_600_v149_: _dafny.Array
                nw81_ = _dafny.Array(_dafny.Map({}), 27)
                d_600_v149_ = nw81_
                rhs108_ = d_599_v148_
                rhs109_ = (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "syuird"))), p0])), p0)) - (p3)
                rhs110_ = (0) - (p3)
                rhs111_ = d_600_v149_
                lhs90_ = globalState
                lhs91_ = globalState
                d_599_v148_ = rhs108_
                lhs90_.f21 = rhs109_
                lhs91_.f5 = rhs110_
                d_600_v149_ = rhs111_
                (globalState).f13 = p0
                (self).f38 = self.f38
                (globalState).f6 = self.f38
        d_601_v150_: _dafny.Array
        def lambda15_(d_602_p0_, d_603_p3_):
            def lambda16_(d_604_i7_):
                return (_dafny.SeqWithoutIsStrInference([D1_DC5(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_602_p0_])]))])) + (_dafny.SeqWithoutIsStrInference([D1_DC5(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_603_p3_ for d_605_i8_ in range(default__.abs(90))]), _dafny.SeqWithoutIsStrInference([d_603_p3_, d_603_p3_])])), D1_DC5(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_602_p0_ for d_606_i9_ in range(default__.abs(325))]), _dafny.SeqWithoutIsStrInference([d_602_p0_, d_602_p0_, d_603_p3_, d_603_p3_, d_603_p3_])])), D1_DC5(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_602_p0_, d_602_p0_])]))]))

            return lambda16_

        init9_ = lambda15_(p0, p3)
        nw82_ = _dafny.Array(None, 10)
        for i0_9_ in range(nw82_.length(0)):
            nw82_[i0_9_] = init9_(i0_9_)
        d_601_v150_ = nw82_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_601_v150_).length(0)):
            d_607_i6_: int = guard_loop_1_
            if (True) and (((0) <= (d_607_i6_)) and ((d_607_i6_) < ((d_601_v150_).length(0)))):
                (d_601_v150_)[(d_607_i6_)] = _dafny.SeqWithoutIsStrInference([D1_DC5(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, p0, p3])])) for d_608_i10_ in range(default__.abs(281))])
        if not((self).f28):
            (globalState).f5 = p3
            d_609_v151_: _dafny.Set
            d_609_v151_ = _dafny.Set({p3})
            def iife68_():
                coll42_ = _dafny.Set()
                compr_42_: int
                for compr_42_ in _dafny.IntegerRange(-827, 751):
                    d_610_v153_: int = compr_42_
                    if ((-827) <= (d_610_v153_)) and ((d_610_v153_) < (751)):
                        coll42_ = coll42_.union(_dafny.Set([default__.safeDivisionInt(d_610_v153_, p3)]))
                return _dafny.Set(coll42_)
            def iife69_():
                coll43_ = _dafny.Set()
                compr_43_: int
                for compr_43_ in (d_609_v151_).Elements:
                    d_611_v152_: int = compr_43_
                    if (d_611_v152_) in (d_609_v151_):
                        coll43_ = coll43_.union(_dafny.Set([(d_611_v152_) + (-414)]))
                return _dafny.Set(coll43_)
            if (iife68_()
            ).ispropersubset(iife69_()
            ):
                d_612_v154_: _dafny.Array
                def lambda17_(d_613_i11_):
                    return (d_613_i11_) + (len(_dafny.SeqWithoutIsStrInference([(self).f37 for d_614_i12_ in range(default__.abs(13))])))

                init10_ = lambda17_
                nw83_ = _dafny.Array(None, 28)
                for i0_10_ in range(nw83_.length(0)):
                    nw83_[i0_10_] = init10_(i0_10_)
                d_612_v154_ = nw83_
                index85_ = default__.safeIndex(401, (d_612_v154_).length(0))
                (d_612_v154_)[index85_] = 29
                d_615_v155_: _dafny.Seq
                d_615_v155_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrefi"))
                d_616_v156_: _dafny.Map
                d_616_v156_ = _dafny.Map({len(d_615_v155_): self.f38})
                index86_ = default__.safeIndex(759, (d_612_v154_).length(0))
                (d_612_v154_)[index86_] = len(d_616_v156_)
                d_617_v157_: _dafny.Array
                d_617_v157_ = (self).f27
                d_618_v158_: _dafny.Map
                d_618_v158_ = _dafny.Map({d_617_v157_: d_568_v120_})
                d_619_v159_: _dafny.Map
                d_619_v159_ = _dafny.Map({(_dafny.Map({d_617_v157_: d_568_v120_})) | (d_618_v158_): (p3) - (p3)})
                index87_ = default__.safeIndex(401, (d_612_v154_).length(0))
                index88_ = default__.safeIndex(759, (d_612_v154_).length(0))
                rhs112_ = p3
                rhs113_ = p0
                rhs114_ = p0
                rhs115_ = (d_619_v159_).set(d_618_v158_, p3)
                rhs116_ = p0
                lhs92_ = globalState
                lhs93_ = d_612_v154_
                lhs94_ = default__.safeIndex(401, (d_612_v154_).length(0))
                lhs95_ = d_612_v154_
                lhs96_ = default__.safeIndex(759, (d_612_v154_).length(0))
                lhs97_ = globalState
                lhs92_.f4 = rhs112_
                lhs93_[lhs94_] = rhs113_
                lhs95_[lhs96_] = rhs114_
                d_619_v159_ = rhs115_
                lhs97_.f4 = rhs116_
                d_620_v160_: _dafny.Map
                d_621_v161_: int
                out11_: _dafny.Map
                out12_: int
                out11_, out12_ = default__.m0(globalState)
                d_620_v160_ = out11_
                d_621_v161_ = out12_
                (globalState).f21 = (p3) * ((d_612_v154_)[default__.safeIndex(401, (d_612_v154_).length(0))])
                d_622_v162_: _dafny.Map
                d_622_v162_ = _dafny.Map({D1_DC6(d_621_v161_): False})
                d_623_v163_: D1
                d_623_v163_ = D1_DC6(p0)
                d_624_v164_: _dafny.Set
                d_624_v164_ = _dafny.Set({self.f38})
                (globalState).f17 = ((d_622_v162_)[d_623_v163_] if (d_623_v163_) in (d_622_v162_) else (default__.fm31(globalState)) == (d_624_v164_))
                d_625_v165_: _dafny.Array
                def lambda18_(d_626_i13_):
                    return (self).f28

                init11_ = lambda18_
                nw84_ = _dafny.Array(None, 1)
                for i0_11_ in range(nw84_.length(0)):
                    nw84_[i0_11_] = init11_(i0_11_)
                d_625_v165_ = nw84_
                d_625_v165_ = (self).f27
            elif True:
                index89_ = default__.safeIndex(209, ((self).f27).length(0))
                ((self).f27)[index89_] = (self).f28
                index90_ = default__.safeIndex(209, ((self).f27).length(0))
                ((self).f27)[index90_] = (default__.safeDivisionInt(p0, (0) - (p0))) < ((0) - (((d_570_v122_)[p3] if (p3) in (d_570_v122_) else p3)))
                d_627_v166_: _dafny.MultiSet
                d_627_v166_ = _dafny.MultiSet([_dafny.CodePoint('m')])
                d_628_v167_: _dafny.Seq
                d_628_v167_ = _dafny.SeqWithoutIsStrInference([d_627_v166_, d_627_v166_])
                d_629_v168_: _dafny.Array
                nw85_ = _dafny.Array(None, 13)
                nw85_[int(0)] = self.f38
                nw85_[int(1)] = self.f38
                nw85_[int(2)] = (self).f28
                nw85_[int(3)] = ((self).f27)[default__.safeIndex(209, ((self).f27).length(0))]
                nw85_[int(4)] = not(((self).f27)[default__.safeIndex(209, ((self).f27).length(0))])
                nw85_[int(5)] = self.f38
                nw85_[int(6)] = self.f38
                nw85_[int(7)] = True
                nw85_[int(8)] = not((self).f28)
                nw85_[int(9)] = not((self).f28)
                nw85_[int(10)] = ((self).f27)[default__.safeIndex(209, ((self).f27).length(0))]
                nw85_[int(11)] = (self).f28
                nw85_[int(12)] = self.f38
                d_629_v168_ = nw85_
                d_630_v169_: C0
                nw86_ = C0()
                nw86_.ctor__(default__.fm32(p0, ((self).f27)[default__.safeIndex(209, ((self).f27).length(0))], globalState), (0) - ((len(d_628_v167_)) - ((0) - (p0))), d_629_v168_, (self).f28)
                d_630_v169_ = nw86_
                (globalState).f13 = p0
                index91_ = default__.safeIndex(825, (p2).length(0))
                (p2)[index91_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))) + ((d_630_v169_).f36)).set(default__.safeIndex(p3, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))) + ((d_630_v169_).f36))), p1)
                index92_ = default__.safeIndex(825, (p2).length(0))
                (p2)[index92_] = (d_630_v169_).f36
                (globalState).f6 = (self.f38 if ((self).f27)[default__.safeIndex(209, ((self).f27).length(0))] else self.f38)
            d_631_v170_: _dafny.Seq
            d_631_v170_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pteeog"))
            d_632_v171_: C0
            nw87_ = C0()
            nw87_.ctor__(d_631_v170_, default__.safeDivisionInt(p3, p3), (self).f27, self.f38)
            d_632_v171_ = nw87_
            d_633_v172_: _dafny.Map
            d_633_v172_ = _dafny.Map({p0: (self).f28})
            d_634_v173_: _dafny.Seq
            d_634_v173_ = _dafny.SeqWithoutIsStrInference([default__.fm22(((d_633_v172_)[p0] if (p0) in (d_633_v172_) else self.f38), p0, globalState), (self).f28])
            if ((len(d_634_v173_)) > (695)) or ((self).f28):
                d_635_v174_: _dafny.Array
                nw88_ = _dafny.Array(None, 14)
                nw88_[int(0)] = (d_569_v121_) - (d_569_v121_)
                nw88_[int(1)] = d_569_v121_
                nw88_[int(2)] = d_569_v121_
                nw88_[int(3)] = d_569_v121_
                nw88_[int(4)] = default__.fm26((self).f28, p3, p1, globalState)
                nw88_[int(5)] = _dafny.MultiSet(d_634_v173_)
                nw88_[int(6)] = d_569_v121_
                nw88_[int(7)] = d_569_v121_
                nw88_[int(8)] = d_569_v121_
                nw88_[int(9)] = (_dafny.MultiSet(d_634_v173_)).intersection(d_569_v121_)
                nw88_[int(10)] = d_569_v121_
                nw88_[int(11)] = d_569_v121_
                nw88_[int(12)] = d_569_v121_
                nw88_[int(13)] = _dafny.MultiSet([self.f38])
                d_635_v174_ = nw88_
                index93_ = default__.safeIndex(907, (d_635_v174_).length(0))
                (d_635_v174_)[index93_] = _dafny.MultiSet([self.f38])
                index94_ = default__.safeIndex(907, (d_635_v174_).length(0))
                (d_635_v174_)[index94_] = (d_569_v121_ if self.f38 else (_dafny.MultiSet(default__.fm25(p0, p0, p0, globalState)) if (d_634_v173_)[default__.safeIndex(p0, len(d_634_v173_))] else _dafny.MultiSet([False, self.f38, (self).f28, (self).f28, (self).f28])))
                d_636_v175_: _dafny.Array
                nw89_ = _dafny.Array(int(0), 21)
                d_636_v175_ = nw89_
                d_637_v176_: _dafny.Array
                nw90_ = _dafny.Array(None, 2)
                nw90_[int(0)] = d_636_v175_
                nw90_[int(1)] = d_636_v175_
                d_637_v176_ = nw90_
                index95_ = default__.safeIndex(491, (d_637_v176_).length(0))
                (d_637_v176_)[index95_] = d_636_v175_
                index96_ = default__.safeIndex(491, (d_637_v176_).length(0))
                (d_637_v176_)[index96_] = d_636_v175_
                d_638_v177_: _dafny.Map
                d_638_v177_ = _dafny.Map({False: default__.fm22(default__.fm22((self).f28, len(d_609_v151_), globalState), p3, globalState)})
                d_639_v178_: _dafny.Set
                d_639_v178_ = _dafny.Set({d_638_v177_, d_638_v177_, d_638_v177_})
                (globalState).f4 = len((d_639_v178_) - ((_dafny.Set({d_638_v177_, d_638_v177_})) | (d_639_v178_)))
                (globalState).f17 = (p0) <= (p0)
                (globalState).f5 = len(_dafny.SeqWithoutIsStrInference([d_569_v121_ for d_640_i14_ in range(default__.abs(515))]))
            elif True:
                d_641_v179_: _dafny.Array
                nw91_ = _dafny.Array(int(0), 29)
                d_641_v179_ = nw91_
                d_642_v180_: _dafny.Map
                d_642_v180_ = _dafny.Map({p3: d_641_v179_})
                d_642_v180_ = (d_642_v180_).set(default__.safeModuloInt(p0, (0) - (p3)), d_641_v179_)
                d_643_v181_: _dafny.Seq
                d_643_v181_ = _dafny.SeqWithoutIsStrInference([p3, (0) - (p3), p3])
                d_644_v182_: _dafny.MultiSet
                d_644_v182_ = _dafny.MultiSet([(_dafny.MultiSet(d_643_v181_)).cardinality, default__.safeModuloInt(p3, p3)])
                d_645_v183_: _dafny.Array
                nw92_ = _dafny.Array(False, 14)
                d_645_v183_ = nw92_
                d_646_v184_: _dafny.Seq
                d_646_v184_ = _dafny.SeqWithoutIsStrInference([d_569_v121_])
                d_647_v185_: _dafny.Array
                nw93_ = _dafny.Array(None, 18)
                nw93_[int(0)] = d_569_v121_
                nw93_[int(1)] = d_569_v121_
                nw93_[int(2)] = (d_569_v121_) - (d_569_v121_)
                nw93_[int(3)] = (d_569_v121_) | (d_569_v121_)
                nw93_[int(4)] = (d_569_v121_) | (d_569_v121_)
                nw93_[int(5)] = (d_646_v184_)[default__.safeIndex(len(d_631_v170_), len(d_646_v184_))]
                nw93_[int(6)] = (_dafny.MultiSet([self.f38, (self).f28, default__.fm22((self).f28, p0, globalState), (self).f28])).intersection(_dafny.MultiSet(d_634_v173_))
                nw93_[int(7)] = ((_dafny.MultiSet(d_634_v173_)).set((self).f28, default__.abs(p0)) if not((self).f28) else d_569_v121_)
                nw93_[int(8)] = (d_569_v121_ if self.f38 else d_569_v121_)
                nw93_[int(9)] = d_569_v121_
                nw93_[int(10)] = _dafny.MultiSet([self.f38])
                nw93_[int(11)] = (d_569_v121_).intersection(d_569_v121_)
                nw93_[int(12)] = (d_569_v121_) | (d_569_v121_)
                nw93_[int(13)] = d_569_v121_
                nw93_[int(14)] = (d_569_v121_) - (d_569_v121_)
                nw93_[int(15)] = d_569_v121_
                nw93_[int(16)] = d_569_v121_
                nw93_[int(17)] = (d_569_v121_) - (d_569_v121_)
                d_647_v185_ = nw93_
                index97_ = default__.safeIndex(670, (d_647_v185_).length(0))
                (d_647_v185_)[index97_] = d_569_v121_
                index98_ = default__.safeIndex(670, (d_647_v185_).length(0))
                rhs117_ = ((self).f28) and (self.f38)
                rhs118_ = d_644_v182_
                rhs119_ = (d_632_v171_).fm15(globalState)
                rhs120_ = (self).f27
                rhs121_ = d_569_v121_
                lhs98_ = self
                lhs99_ = globalState
                lhs100_ = d_647_v185_
                lhs101_ = default__.safeIndex(670, (d_647_v185_).length(0))
                lhs98_.f38 = rhs117_
                d_644_v182_ = rhs118_
                lhs99_.f4 = rhs119_
                d_645_v183_ = rhs120_
                lhs100_[lhs101_] = rhs121_
                rhs122_ = p3
                rhs123_ = d_633_v172_
                lhs102_ = globalState
                lhs102_.f4 = rhs122_
                d_633_v172_ = rhs123_
                (globalState).f13 = default__.safeModuloInt((p0) - (p0), (0) - (p3))
                d_648_v186_: _dafny.Array
                def lambda19_(d_649_p1_):
                    def lambda20_(d_650_i15_):
                        return d_649_p1_

                    return lambda20_

                init12_ = lambda19_(p1)
                nw94_ = _dafny.Array(None, 9)
                for i0_12_ in range(nw94_.length(0)):
                    nw94_[i0_12_] = init12_(i0_12_)
                d_648_v186_ = nw94_
                d_651_v187_: _dafny.Seq
                d_651_v187_ = _dafny.SeqWithoutIsStrInference([d_648_v186_])
                (globalState).f21 = len(d_651_v187_)
            d_652_v188_: _dafny.Array
            def lambda21_(d_653_i16_):
                return _dafny.SeqWithoutIsStrInference([D5_DC12(_dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f28, (self).f28, self.f38}) for d_654_i18_ in range(default__.abs(458))])) for d_655_i17_ in range(default__.abs(274))])

            init13_ = lambda21_
            nw95_ = _dafny.Array(None, 24)
            for i0_13_ in range(nw95_.length(0)):
                nw95_[i0_13_] = init13_(i0_13_)
            d_652_v188_ = nw95_
            d_652_v188_ = d_652_v188_
        elif True:
            d_656_v189_: _dafny.Seq
            d_656_v189_ = _dafny.SeqWithoutIsStrInference([self.f38, (self).f28])
            d_657_v190_: _dafny.MultiSet
            d_657_v190_ = _dafny.MultiSet([len(d_656_v189_), p3])
            d_658_v191_: _dafny.Map
            d_658_v191_ = _dafny.Map({self.f38: p3})
            d_659_v192_: _dafny.Seq
            d_659_v192_ = _dafny.SeqWithoutIsStrInference([(d_657_v190_).set(len(d_658_v191_), default__.abs(p0)), d_657_v190_])
            (globalState).f6 = ((_dafny.SeqWithoutIsStrInference([d_657_v190_, (d_659_v192_)[default__.safeIndex(p3, len(d_659_v192_))]])) + (_dafny.SeqWithoutIsStrInference([d_657_v190_ for d_660_i19_ in range(default__.abs(259))]))) != (d_659_v192_)
            d_661_v193_: _dafny.Seq
            d_661_v193_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iiaprsa"))
            d_662_v194_: C0
            nw96_ = C0()
            nw96_.ctor__((d_661_v193_) + (d_661_v193_), p0, (self).f27, self.f38)
            d_662_v194_ = nw96_
            d_663_v195_: _dafny.Array
            nw97_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_663_v195_ = nw97_
            d_664_v196_: _dafny.Seq
            d_664_v196_ = _dafny.SeqWithoutIsStrInference([d_663_v195_])
            d_665_v197_: _dafny.Array
            nw98_ = _dafny.Array(None, 20)
            nw98_[int(0)] = d_663_v195_
            nw98_[int(1)] = d_663_v195_
            nw98_[int(2)] = d_663_v195_
            nw98_[int(3)] = d_663_v195_
            nw98_[int(4)] = d_663_v195_
            nw98_[int(5)] = d_663_v195_
            nw98_[int(6)] = d_663_v195_
            nw98_[int(7)] = d_663_v195_
            nw98_[int(8)] = d_663_v195_
            nw98_[int(9)] = d_663_v195_
            nw98_[int(10)] = d_663_v195_
            nw98_[int(11)] = d_663_v195_
            nw98_[int(12)] = d_663_v195_
            nw98_[int(13)] = d_663_v195_
            nw98_[int(14)] = d_663_v195_
            nw98_[int(15)] = (d_664_v196_)[default__.safeIndex(p3, len(d_664_v196_))]
            nw98_[int(16)] = ((d_664_v196_)[default__.safeIndex(p0, len(d_664_v196_))] if not(self.f38) else d_663_v195_)
            nw98_[int(17)] = d_663_v195_
            nw98_[int(18)] = d_663_v195_
            nw98_[int(19)] = d_663_v195_
            d_665_v197_ = nw98_
            index99_ = default__.safeIndex(866, (d_665_v197_).length(0))
            (d_665_v197_)[index99_] = d_663_v195_
            index100_ = default__.safeIndex(866, (d_665_v197_).length(0))
            rhs124_ = (self).f28
            rhs125_ = d_663_v195_
            rhs126_ = self.f38
            lhs103_ = globalState
            lhs104_ = d_665_v197_
            lhs105_ = default__.safeIndex(866, (d_665_v197_).length(0))
            lhs106_ = globalState
            lhs103_.f17 = rhs124_
            lhs104_[lhs105_] = rhs125_
            lhs106_.f6 = rhs126_
            d_666_v198_: _dafny.Seq
            d_666_v198_ = _dafny.SeqWithoutIsStrInference([(d_568_v120_).cf7, p3, p0])
            d_667_v199_: _dafny.Seq
            d_667_v199_ = d_666_v198_
            d_668_v200_: _dafny.Seq
            d_668_v200_ = (d_666_v198_ if self.f38 else (d_667_v199_))
            source9_ = d_667_v199_
            d_669___mcc_h15_ = source9_
            d_670_cf15_ = d_669___mcc_h15_
            d_671_v201_: _dafny.Array
            def lambda22_(d_672_p1_):
                def lambda23_(d_673_i20_):
                    return d_672_p1_

                return lambda23_

            init14_ = lambda22_(p1)
            nw99_ = _dafny.Array(None, 15)
            for i0_14_ in range(nw99_.length(0)):
                nw99_[i0_14_] = init14_(i0_14_)
            d_671_v201_ = nw99_
            d_674_v202_: _dafny.Map
            d_674_v202_ = _dafny.Map({(d_671_v201_ if False else d_671_v201_): (default__.fm32(p3, not(self.f38), globalState)) <= ((d_662_v194_).f36)})
            (globalState).f17 = ((d_674_v202_)[d_671_v201_] if (d_671_v201_) in (d_674_v202_) else (False if (self).f28 else (self).f28))
            d_675_v203_: _dafny.Set
            d_675_v203_ = _dafny.Set({(self).f37})
            d_676_v204_: _dafny.Map
            d_676_v204_ = _dafny.Map({d_675_v203_: (self).f28})
            d_677_v205_: _dafny.Map
            d_677_v205_ = _dafny.Map({(p3) >= (p3): (d_676_v204_) | (d_676_v204_)})
            d_677_v205_ = (d_677_v205_).set(((0) - ((self).fm20(globalState))) == (p3), d_676_v204_)
            d_678_v206_: _dafny.Map
            d_678_v206_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgpyl"))): p1})
            rhs127_ = default__.safeModuloInt(p3, len(d_670_cf15_))
            rhs128_ = (p0 if self.f38 else ((0) - ((self).fm20(globalState))) * (len(d_678_v206_)))
            lhs107_ = globalState
            lhs108_ = globalState
            lhs107_.f5 = rhs127_
            lhs108_.f13 = rhs128_
            (globalState).f21 = p0
            rhs129_ = p0
            rhs130_ = (d_656_v189_)[default__.safeIndex(236, len(d_656_v189_))]
            lhs109_ = globalState
            lhs109_.f13 = rhs129_
            r0 = rhs130_
        r0 = (p3) == (p3)
        return r0

    def m13(self, p0, p1, p2, p3, globalState):
        r0: D0 = D0.default()()
        d_679_i0_: int
        d_679_i0_ = 0
        with _dafny.label("1"):
            while (p1) or (not(p0)):
                with _dafny.c_label("1"):
                    if (d_679_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_679_i0_ = (d_679_i0_) + (1)
                    d_680_v0_: T1
                    nw100_ = C0()
                    nw100_.ctor__(p3, p2, (self).f27, p1)
                    d_680_v0_ = nw100_
                    d_681_v1_: _dafny.MultiSet
                    d_681_v1_ = _dafny.MultiSet([548])
                    d_682_v2_: _dafny.Map
                    d_682_v2_ = _dafny.Map({d_680_v0_: (0) - ((d_681_v1_).cardinality)})
                    d_683_v3_: _dafny.Map
                    d_683_v3_ = _dafny.Map({d_682_v2_: p0})
                    d_684_v4_: D0
                    d_684_v4_ = D0_DC0(d_683_v3_, (d_681_v1_).ispropersubset(d_681_v1_), len(_dafny.Set({p2})), (self).f37, (self).f28)
                    d_685_v5_: _dafny.Set
                    d_685_v5_ = _dafny.Set({len(p3)})
                    d_686_v8_: _dafny.Map
                    def iife70_():
                        coll44_ = _dafny.Set()
                        compr_44_: int
                        for compr_44_ in _dafny.IntegerRange(524, 662):
                            d_687_v7_: int = compr_44_
                            if ((524) <= (d_687_v7_)) and ((d_687_v7_) < (662)):
                                coll44_ = coll44_.union(_dafny.Set([(d_687_v7_) + (p2)]))
                        return _dafny.Set(coll44_)
                    d_686_v8_ = _dafny.Map({d_680_v0_.f29: iife70_()
                    })
                    d_688_v9_: _dafny.Array
                    nw101_ = _dafny.Array(None, 16)
                    nw101_[int(0)] = d_685_v5_
                    nw101_[int(1)] = (_dafny.Set({p2})) | (d_685_v5_)
                    nw101_[int(2)] = d_685_v5_
                    def iife71_():
                        coll45_ = _dafny.Set()
                        compr_45_: int
                        for compr_45_ in _dafny.IntegerRange(182, 243):
                            d_689_v6_: int = compr_45_
                            if ((182) <= (d_689_v6_)) and ((d_689_v6_) < (243)):
                                coll45_ = coll45_.union(_dafny.Set([(d_689_v6_) - (d_680_v0_.f29)]))
                        return _dafny.Set(coll45_)
                    nw101_[int(3)] = iife71_()
                    
                    nw101_[int(4)] = d_685_v5_
                    nw101_[int(5)] = (default__.fm0(d_686_v8_, p0, self.f38, globalState)) - (d_685_v5_)
                    nw101_[int(6)] = d_685_v5_
                    nw101_[int(7)] = d_685_v5_
                    nw101_[int(8)] = d_685_v5_
                    nw101_[int(9)] = d_685_v5_
                    nw101_[int(10)] = d_685_v5_
                    nw101_[int(11)] = d_685_v5_
                    nw101_[int(12)] = d_685_v5_
                    nw101_[int(13)] = _dafny.Set({-582})
                    nw101_[int(14)] = d_685_v5_
                    nw101_[int(15)] = d_685_v5_
                    d_688_v9_ = nw101_
                    d_690_v10_: _dafny.Seq
                    d_690_v10_ = _dafny.SeqWithoutIsStrInference([d_685_v5_])
                    index101_ = default__.safeIndex(55, (d_688_v9_).length(0))
                    (d_688_v9_)[index101_] = (d_690_v10_)[default__.safeIndex(256, len(d_690_v10_))]
                    d_691_v11_: _dafny.Array
                    def lambda24_(d_692_i1_):
                        return _dafny.SeqWithoutIsStrInference([(self).f37 for d_693_i2_ in range(default__.abs(503))])

                    init15_ = lambda24_
                    nw102_ = _dafny.Array(None, 11)
                    for i0_15_ in range(nw102_.length(0)):
                        nw102_[i0_15_] = init15_(i0_15_)
                    d_691_v11_ = nw102_
                    index102_ = default__.safeIndex(80, (d_691_v11_).length(0))
                    (d_691_v11_)[index102_] = p3
                    d_694_v12_: _dafny.Map
                    d_694_v12_ = _dafny.Map({True: (0) - (p2)})
                    d_695_v13_: _dafny.Seq
                    d_695_v13_ = _dafny.SeqWithoutIsStrInference([default__.fm22(p0, ((d_694_v12_)[(self).f28] if ((self).f28) in (d_694_v12_) else d_680_v0_.f29), globalState), (d_680_v0_).f28])
                    d_696_v14_: _dafny.Array
                    def lambda25_(d_697_v0_):
                        def lambda26_(d_698_i3_):
                            return (d_698_i3_) * (d_697_v0_.f29)

                        return lambda26_

                    init16_ = lambda25_(d_680_v0_)
                    nw103_ = _dafny.Array(None, 18)
                    for i0_16_ in range(nw103_.length(0)):
                        nw103_[i0_16_] = init16_(i0_16_)
                    d_696_v14_ = nw103_
                    d_699_v15_: _dafny.MultiSet
                    d_699_v15_ = _dafny.MultiSet([d_696_v14_])
                    index103_ = default__.safeIndex(55, (d_688_v9_).length(0))
                    index104_ = default__.safeIndex(80, (d_691_v11_).length(0))
                    rhs131_ = d_684_v4_
                    rhs132_ = (d_685_v5_) - (d_685_v5_)
                    rhs133_ = (p3).set(default__.safeIndex(len(d_694_v12_), len(p3)), (self).f37)
                    rhs134_ = (d_695_v13_)[default__.safeIndex(default__.safeDivisionInt(p2, (d_699_v15_).cardinality), len(d_695_v13_))]
                    lhs110_ = d_688_v9_
                    lhs111_ = default__.safeIndex(55, (d_688_v9_).length(0))
                    lhs112_ = d_691_v11_
                    lhs113_ = default__.safeIndex(80, (d_691_v11_).length(0))
                    lhs114_ = globalState
                    d_684_v4_ = rhs131_
                    lhs110_[lhs111_] = rhs132_
                    lhs112_[lhs113_] = rhs133_
                    lhs114_.f15 = rhs134_
                    d_684_v4_ = d_684_v4_
                    (globalState).f5 = (0) - (p2)
                    index105_ = default__.safeIndex(80, (d_691_v11_).length(0))
                    (d_691_v11_)[index105_] = p3
                    pass
            pass
        if p0:
            if p0:
                d_700_v16_: _dafny.Seq
                d_700_v16_ = _dafny.SeqWithoutIsStrInference([p2])
                d_701_v17_: _dafny.Seq
                d_701_v17_ = _dafny.SeqWithoutIsStrInference([len(d_700_v16_)])
                d_702_v18_: _dafny.Map
                d_702_v18_ = _dafny.Map({p1: (d_701_v17_)[default__.safeIndex(p2, len(d_701_v17_))]})
                (globalState).f15 = default__.fm22(True, (0) - (((d_702_v18_)[True] if (True) in (d_702_v18_) else p2)), globalState)
                (globalState).f17 = (d_700_v16_) != (d_700_v16_)
                def iife72_():
                    coll46_ = _dafny.Map()
                    compr_46_: int
                    for compr_46_ in _dafny.IntegerRange(-895, 596):
                        d_703_v19_: int = compr_46_
                        if ((-895) <= (d_703_v19_)) and ((d_703_v19_) < (596)):
                            coll46_[(d_703_v19_) - (p2)] = 647
                    return _dafny.Map(coll46_)
                (globalState).f15 = (p2) in (iife72_()
                )
                d_704_v20_: C0
                nw104_ = C0()
                nw104_.ctor__(p3, (p2) * (p2), (self).f27, False)
                d_704_v20_ = nw104_
                (globalState).f15 = False
            elif True:
                d_705_v21_: C0
                nw105_ = C0()
                nw105_.ctor__(p3, -948, (self).f27, p1)
                d_705_v21_ = nw105_
                d_706_v22_: _dafny.Map
                d_706_v22_ = _dafny.Map({p2: p2})
                d_707_v23_: _dafny.Seq
                d_707_v23_ = _dafny.SeqWithoutIsStrInference([default__.fm22(True, p2, globalState), (p2) == (((d_706_v22_)[p2] if (p2) in (d_706_v22_) else p2)), True])
                d_707_v23_ = (d_707_v23_).set(default__.safeIndex(p2, len(d_707_v23_)), p1)
                d_708_v24_: _dafny.Seq
                d_708_v24_ = _dafny.SeqWithoutIsStrInference([p2, p2])
                d_709_v25_: _dafny.Map
                d_709_v25_ = _dafny.Map({p0: len(d_708_v24_)})
                d_710_v26_: _dafny.Map
                d_710_v26_ = _dafny.Map({d_709_v25_: p2})
                d_711_v27_: D8
                d_711_v27_ = D8_DC17(d_710_v26_)
                d_712_v28_: _dafny.Seq
                d_712_v28_ = _dafny.SeqWithoutIsStrInference([d_710_v26_, d_710_v26_])
                d_713_v29_: _dafny.Map
                d_713_v29_ = _dafny.Map({self.f38: p0})
                d_714_v30_: _dafny.Array
                nw106_ = _dafny.Array(None, 25)
                nw106_[int(0)] = (0) - (p2)
                nw106_[int(1)] = (d_705_v21_).fm15(globalState)
                nw106_[int(2)] = p2
                nw106_[int(3)] = p2
                nw106_[int(4)] = p2
                nw106_[int(5)] = p2
                nw106_[int(6)] = p2
                nw106_[int(7)] = (p2 if False else 144)
                nw106_[int(8)] = len(default__.fm33((0) - (p2), (self).f28, globalState))
                nw106_[int(9)] = (p2) + (p2)
                nw106_[int(10)] = (0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdbbk")))) + (p2))
                nw106_[int(11)] = p2
                nw106_[int(12)] = p2
                nw106_[int(13)] = len(((d_711_v27_).cf30) | ((d_712_v28_)[default__.safeIndex((d_705_v21_).fm15(globalState), len(d_712_v28_))]))
                nw106_[int(14)] = p2
                nw106_[int(15)] = p2
                nw106_[int(16)] = p2
                nw106_[int(17)] = p2
                nw106_[int(18)] = (d_708_v24_)[default__.safeIndex(p2, len(d_708_v24_))]
                nw106_[int(19)] = p2
                nw106_[int(20)] = (0) - (len(_dafny.Map({(d_707_v23_)[default__.safeIndex(p2, len(d_707_v23_))]: (self).f28})))
                nw106_[int(21)] = p2
                nw106_[int(22)] = p2
                nw106_[int(23)] = p2
                nw106_[int(24)] = len((d_713_v29_).set(self.f38, (self).f28))
                d_714_v30_ = nw106_
                index106_ = default__.safeIndex(406, (d_714_v30_).length(0))
                (d_714_v30_)[index106_] = p2
                d_715_v31_: _dafny.MultiSet
                d_715_v31_ = _dafny.MultiSet([len(d_708_v24_), p2])
                index107_ = default__.safeIndex(406, (d_714_v30_).length(0))
                (d_714_v30_)[index107_] = default__.safeModuloInt(default__.safeModuloInt((0) - (p2), (d_715_v31_).cardinality), default__.safeModuloInt(p2, p2))
                (globalState).f4 = (d_714_v30_)[default__.safeIndex(406, (d_714_v30_).length(0))]
                d_716_v32_: T0
                nw107_ = C1()
                nw107_.ctor__(self.f38, (self).f27, p1)
                d_716_v32_ = nw107_
                d_717_v33_: _dafny.Map
                d_717_v33_ = _dafny.Map({d_716_v32_: (d_713_v29_).set(False, p0)})
                d_718_v34_: _dafny.Set
                d_718_v34_ = _dafny.Set({(d_714_v30_)[default__.safeIndex(406, (d_714_v30_).length(0))], p2, len(d_717_v33_), p2})
                (globalState).f15 = (_dafny.Set({p2, p2, (0) - ((d_714_v30_)[default__.safeIndex(406, (d_714_v30_).length(0))])})).issubset(d_718_v34_)
            d_719_v35_: _dafny.Map
            d_719_v35_ = _dafny.Map({default__.safeModuloInt((0) - (p2), p2): p0})
            d_719_v35_ = (d_719_v35_).set(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iivinjgue"))), (self).f28)
            d_720_v36_: C0
            nw108_ = C0()
            nw108_.ctor__(p3, p2, (self).f27, (self).f28)
            d_720_v36_ = nw108_
            d_720_v36_ = d_720_v36_
            d_721_v37_: _dafny.Array
            nw109_ = _dafny.Array(int(0), 17)
            d_721_v37_ = nw109_
            nw110_ = _dafny.Array(None, 9)
            nw110_[int(0)] = d_721_v37_
            nw110_[int(1)] = d_721_v37_
            nw110_[int(2)] = d_721_v37_
            nw110_[int(3)] = d_721_v37_
            nw110_[int(4)] = d_721_v37_
            nw110_[int(5)] = d_721_v37_
            nw110_[int(6)] = d_721_v37_
            nw110_[int(7)] = d_721_v37_
            nw110_[int(8)] = d_721_v37_
            (globalState).f18 = nw110_
            d_722_v38_: _dafny.Seq
            d_722_v38_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f28})])
            d_723_v39_: _dafny.Map
            d_723_v39_ = _dafny.Map({(self).f27: d_722_v38_})
            d_724_v40_: _dafny.Array
            d_724_v40_ = (self).f27
            pat_let_tv20_ = d_723_v39_
            pat_let_tv21_ = d_724_v40_
            pat_let_tv22_ = d_724_v40_
            pat_let_tv23_ = d_723_v39_
            pat_let_tv24_ = d_722_v38_
            def iife73_(_pat_let13_0):
                def iife74_(d_725_dt__update__tmp_h0_):
                    def iife75_(_pat_let14_0):
                        def iife76_(d_726_dt__update_hcf20_h0_):
                            return D5_DC12(d_726_dt__update_hcf20_h0_)
                        return iife76_(_pat_let14_0)
                    return iife75_(((pat_let_tv20_)[pat_let_tv21_] if (pat_let_tv22_) in (pat_let_tv23_) else pat_let_tv24_))
                return iife74_(_pat_let13_0)
            source10_ = iife73_(D5_DC12(d_722_v38_))
            if source10_.is_DC13:
                d_727___mcc_h0_ = source10_.cf21
                d_728___mcc_h1_ = source10_.cf22
                d_729___mcc_h2_ = source10_.cf23
                d_730_cf23_ = d_729___mcc_h2_
                d_731_cf22_ = d_728___mcc_h1_
                d_732_cf21_ = d_727___mcc_h0_
                d_733_v41_: _dafny.Seq
                d_733_v41_ = _dafny.SeqWithoutIsStrInference([p0])
                d_734_v42_: _dafny.Seq
                d_734_v42_ = _dafny.SeqWithoutIsStrInference([d_733_v41_])
                rhs135_ = d_734_v42_
                rhs136_ = p2
                lhs115_ = globalState
                d_734_v42_ = rhs135_
                lhs115_.f21 = rhs136_
                d_735_v43_: _dafny.MultiSet
                d_735_v43_ = _dafny.MultiSet([self.f38])
                d_736_v44_: _dafny.Seq
                d_736_v44_ = _dafny.SeqWithoutIsStrInference([p2])
                rhs137_ = p2
                rhs138_ = default__.safeDivisionInt(len(d_736_v44_), p2)
                rhs139_ = d_735_v43_
                lhs116_ = globalState
                lhs117_ = globalState
                lhs116_.f13 = rhs137_
                lhs117_.f5 = rhs138_
                d_735_v43_ = rhs139_
                d_737_v45_: _dafny.Seq
                d_737_v45_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27, (self).f27])
                d_738_v46_: _dafny.Seq
                d_738_v46_ = _dafny.SeqWithoutIsStrInference([((d_737_v45_) + (d_737_v45_)).set(default__.safeIndex(p2, len((d_737_v45_) + (d_737_v45_))), (self).f27), (d_737_v45_) + (d_737_v45_), d_737_v45_, (d_737_v45_) + (d_737_v45_)])
                d_737_v45_ = (d_738_v46_)[default__.safeIndex(p2, len(d_738_v46_))]
                (globalState).f6 = d_730_cf23_
            elif True:
                d_739___mcc_h3_ = source10_.cf20
                d_740_cf20_ = d_739___mcc_h3_
                d_741_v47_: _dafny.Array
                nw111_ = _dafny.Array(None, 27)
                nw111_[int(0)] = d_721_v37_
                nw111_[int(1)] = d_721_v37_
                nw111_[int(2)] = d_721_v37_
                nw111_[int(3)] = d_721_v37_
                nw111_[int(4)] = d_721_v37_
                nw111_[int(5)] = d_721_v37_
                nw111_[int(6)] = d_721_v37_
                nw111_[int(7)] = d_721_v37_
                nw111_[int(8)] = (d_721_v37_ if p0 else d_721_v37_)
                nw111_[int(9)] = d_721_v37_
                nw111_[int(10)] = d_721_v37_
                nw111_[int(11)] = d_721_v37_
                nw111_[int(12)] = d_721_v37_
                nw111_[int(13)] = d_721_v37_
                nw111_[int(14)] = d_721_v37_
                nw111_[int(15)] = d_721_v37_
                nw111_[int(16)] = d_721_v37_
                nw111_[int(17)] = d_721_v37_
                nw111_[int(18)] = d_721_v37_
                nw111_[int(19)] = d_721_v37_
                nw111_[int(20)] = d_721_v37_
                nw111_[int(21)] = d_721_v37_
                nw111_[int(22)] = d_721_v37_
                nw111_[int(23)] = d_721_v37_
                nw111_[int(24)] = d_721_v37_
                nw111_[int(25)] = d_721_v37_
                nw111_[int(26)] = d_721_v37_
                d_741_v47_ = nw111_
                d_742_v48_: _dafny.Seq
                d_742_v48_ = _dafny.SeqWithoutIsStrInference([d_721_v37_])
                index108_ = default__.safeIndex(908, (d_741_v47_).length(0))
                (d_741_v47_)[index108_] = (d_742_v48_)[default__.safeIndex((0) - (p2), len(d_742_v48_))]
                index109_ = default__.safeIndex(908, (d_741_v47_).length(0))
                (d_741_v47_)[index109_] = d_721_v37_
                d_743_v49_: C1
                nw112_ = C1()
                nw112_.ctor__(p0, (self).f27, p1)
                d_743_v49_ = nw112_
                d_744_v50_: _dafny.MultiSet
                d_744_v50_ = _dafny.MultiSet([p1])
                d_745_v51_: _dafny.Set
                d_745_v51_ = _dafny.Set({(d_743_v49_).f39})
                d_746_v52_: C0
                nw113_ = C0()
                nw113_.ctor__((d_743_v49_).fm34((d_743_v49_).f39, p2, d_744_v50_, p2, globalState), (p2) * (p2), (self).f27, (True) in (d_745_v51_))
                d_746_v52_ = nw113_
                d_747_v53_: _dafny.Array
                nw114_ = _dafny.Array(_dafny.Seq({}), 7)
                d_747_v53_ = nw114_
                d_747_v53_ = d_747_v53_
        elif True:
            (globalState).f21 = (p2) - ((p2 if p1 else p2))
            d_748_v54_: C1
            nw115_ = C1()
            nw115_.ctor__(not(self.f38), (self).f27, (self).f28)
            d_748_v54_ = nw115_
            if (d_748_v54_).f39:
                d_749_v55_: _dafny.Set
                d_749_v55_ = _dafny.Set({(d_748_v54_).f39})
                rhs140_ = (_dafny.MultiSet([(self).f28, (self).f28])).ispropersubset(_dafny.MultiSet([p0, (self).f28]))
                rhs141_ = p2
                rhs142_ = default__.fm36(globalState)
                rhs143_ = p2
                rhs144_ = not(not ((d_749_v55_).issubset(d_749_v55_)) or (p0))
                lhs118_ = globalState
                lhs119_ = globalState
                lhs120_ = globalState
                lhs121_ = globalState
                lhs122_ = globalState
                lhs118_.f6 = rhs140_
                lhs119_.f13 = rhs141_
                lhs120_.f13 = rhs142_
                lhs121_.f21 = rhs143_
                lhs122_.f15 = rhs144_
                d_750_v56_: _dafny.Set
                d_750_v56_ = _dafny.Set({840})
                d_751_v57_: _dafny.Seq
                d_751_v57_ = _dafny.SeqWithoutIsStrInference([p2, default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxsf"))), p2), len(d_750_v56_), p2])
                d_752_v58_: _dafny.Seq
                d_752_v58_ = d_751_v57_
                d_751_v57_ = (d_752_v58_)
                d_753_v59_: D8
                d_753_v59_ = D8_DC18(p2)
                (globalState).f25 = default__.fm42((d_753_v59_).cf31, globalState)
                (globalState).f15 = self.f38
                d_754_v60_: _dafny.Map
                d_754_v60_ = _dafny.Map({not((d_748_v54_).f39): (self).f27})
                d_754_v60_ = (d_754_v60_).set((d_748_v54_).f39, (self).f27)
            elif True:
                (self).f38 = True
                d_755_v61_: _dafny.Seq
                d_755_v61_ = _dafny.SeqWithoutIsStrInference([863])
                d_756_v62_: _dafny.Seq
                d_756_v62_ = d_755_v61_
                d_757_v63_: _dafny.Map
                d_757_v63_ = _dafny.Map({d_756_v62_: (d_755_v61_)[default__.safeIndex(971, len(d_755_v61_))]})
                d_758_v64_: _dafny.Map
                d_758_v64_ = _dafny.Map({p0: d_757_v63_})
                d_759_v65_: _dafny.Seq
                d_759_v65_ = _dafny.SeqWithoutIsStrInference([((d_758_v64_)[False] if (False) in (d_758_v64_) else d_757_v63_)])
                d_760_v66_: D0
                d_760_v66_ = D0_DC3(p2, (self).f28, p0)
                d_761_v68_: _dafny.MultiSet
                d_761_v68_ = _dafny.MultiSet([d_756_v62_])
                d_762_v69_: _dafny.Seq
                d_762_v69_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_763_v70_: _dafny.Map
                d_763_v70_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.Map({(d_748_v54_).f39: (d_760_v66_).cf7})).set(True, p2)])): d_757_v63_})
                d_764_v72_: _dafny.Array
                nw116_ = _dafny.Array(None, 25)
                nw116_[int(0)] = (d_759_v65_)[default__.safeIndex(p2, len(d_759_v65_))]
                nw116_[int(1)] = d_757_v63_
                nw116_[int(2)] = d_757_v63_
                nw116_[int(3)] = d_757_v63_
                nw116_[int(4)] = (d_757_v63_) | ((d_757_v63_).set(d_755_v61_, p2))
                nw116_[int(5)] = default__.fm43(p2, d_760_v66_, p2, self.f38, globalState)
                nw116_[int(6)] = d_757_v63_
                def iife77_():
                    coll47_ = _dafny.Map()
                    compr_47_: _dafny.Seq
                    for compr_47_ in (d_761_v68_).Elements:
                        d_765_v67_: _dafny.Seq = compr_47_
                        if (d_765_v67_) in (d_761_v68_):
                            coll47_[d_765_v67_] = p2
                    return _dafny.Map(coll47_)
                nw116_[int(7)] = (iife77_()
                ) | (_dafny.Map({d_756_v62_: p2}))
                nw116_[int(8)] = _dafny.Map({d_756_v62_: 592})
                nw116_[int(9)] = (d_757_v63_) | (d_757_v63_)
                nw116_[int(10)] = d_757_v63_
                nw116_[int(11)] = (_dafny.Map({d_756_v62_: p2})) | (d_757_v63_)
                nw116_[int(12)] = (d_757_v63_) | (d_757_v63_)
                nw116_[int(13)] = d_757_v63_
                nw116_[int(14)] = d_757_v63_
                nw116_[int(15)] = (d_757_v63_) | (d_757_v63_)
                nw116_[int(16)] = (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(p3), p2, len(d_755_v61_)]): len(d_762_v69_)})) | (d_757_v63_)
                nw116_[int(17)] = (d_757_v63_) | (d_757_v63_)
                nw116_[int(18)] = (d_757_v63_).set(d_756_v62_, p2)
                nw116_[int(19)] = (d_757_v63_) | (d_757_v63_)
                nw116_[int(20)] = (_dafny.Map({_dafny.SeqWithoutIsStrInference([p2, len(p3)]): p2})) | (d_757_v63_)
                nw116_[int(21)] = d_757_v63_
                def iife78_():
                    coll48_ = _dafny.Map()
                    compr_48_: _dafny.Seq
                    for compr_48_ in (d_761_v68_).Elements:
                        d_766_v71_: _dafny.Seq = compr_48_
                        if (d_766_v71_) in (d_761_v68_):
                            coll48_[d_766_v71_] = p2
                    return _dafny.Map(coll48_)
                nw116_[int(22)] = ((d_763_v70_)[p2] if (p2) in (d_763_v70_) else iife78_()
                )
                nw116_[int(23)] = (d_757_v63_) | (d_757_v63_)
                nw116_[int(24)] = (d_757_v63_) | (d_757_v63_)
                d_764_v72_ = nw116_
                index110_ = default__.safeIndex(399, (d_764_v72_).length(0))
                (d_764_v72_)[index110_] = d_757_v63_
                index111_ = default__.safeIndex(399, (d_764_v72_).length(0))
                (d_764_v72_)[index111_] = (d_757_v63_) | (d_757_v63_)
                d_767_v73_: _dafny.MultiSet
                d_767_v73_ = _dafny.MultiSet([p0, True])
                d_768_v74_: _dafny.Map
                d_768_v74_ = _dafny.Map({(self).f28: p2})
                d_769_v75_: T1
                nw117_ = C0()
                nw117_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "at")), len(d_768_v74_), (self).f27, self.f38)
                d_769_v75_ = nw117_
                d_770_v76_: _dafny.Map
                d_770_v76_ = _dafny.Map({d_769_v75_: p2})
                d_771_v77_: _dafny.Map
                d_771_v77_ = _dafny.Map({d_770_v76_: (d_748_v54_).f39})
                d_772_v78_: _dafny.Seq
                d_772_v78_ = _dafny.SeqWithoutIsStrInference([(d_769_v75_).f27])
                d_773_v79_: _dafny.Map
                d_773_v79_ = _dafny.Map({d_769_v75_.f29: p2})
                d_774_v80_: _dafny.Array
                nw118_ = _dafny.Array(None, 19)
                nw118_[int(0)] = ((0) - (p2)) + (922)
                nw118_[int(1)] = (p2 if (self).f28 else p2)
                nw118_[int(2)] = p2
                nw118_[int(3)] = default__.safeModuloInt(len(d_755_v61_), (d_767_v73_).cardinality)
                nw118_[int(4)] = p2
                nw118_[int(5)] = p2
                nw118_[int(6)] = (D0_DC0(d_771_v77_, (self).f28, len(d_772_v78_), (self).f37, p1)).cf2
                nw118_[int(7)] = d_769_v75_.f29
                nw118_[int(8)] = (d_769_v75_.f29) + (len(d_773_v79_))
                nw118_[int(9)] = p2
                nw118_[int(10)] = d_769_v75_.f29
                nw118_[int(11)] = d_769_v75_.f29
                nw118_[int(12)] = p2
                nw118_[int(13)] = default__.safeModuloInt(p2, p2)
                nw118_[int(14)] = default__.safeModuloInt(len(p3), (0) - (len(p3)))
                nw118_[int(15)] = p2
                nw118_[int(16)] = (p2) - (d_769_v75_.f29)
                nw118_[int(17)] = (((d_773_v79_)[(0) - (d_769_v75_.f29)] if ((0) - (d_769_v75_.f29)) in (d_773_v79_) else d_769_v75_.f29)) - (len(d_768_v74_))
                nw118_[int(18)] = len(d_772_v78_)
                d_774_v80_ = nw118_
                index112_ = default__.safeIndex(185, (d_774_v80_).length(0))
                (d_774_v80_)[index112_] = p2
                d_775_v81_: _dafny.Array
                nw119_ = _dafny.Array(_dafny.Set({}), 16)
                d_775_v81_ = nw119_
                d_776_v82_: _dafny.Set
                d_776_v82_ = _dafny.Set({default__.fm40(globalState), p1, not(p1)})
                index113_ = default__.safeIndex(878, (d_775_v81_).length(0))
                (d_775_v81_)[index113_] = d_776_v82_
                index114_ = default__.safeIndex(185, (d_774_v80_).length(0))
                index115_ = default__.safeIndex(878, (d_775_v81_).length(0))
                rhs145_ = (d_748_v54_).f39
                rhs146_ = p0
                rhs147_ = p2
                rhs148_ = _dafny.Set({p1})
                lhs123_ = globalState
                lhs124_ = globalState
                lhs125_ = d_774_v80_
                lhs126_ = default__.safeIndex(185, (d_774_v80_).length(0))
                lhs127_ = d_775_v81_
                lhs128_ = default__.safeIndex(878, (d_775_v81_).length(0))
                lhs123_.f17 = rhs145_
                lhs124_.f6 = rhs146_
                lhs125_[lhs126_] = rhs147_
                lhs127_[lhs128_] = rhs148_
                d_777_v83_: D3
                d_777_v83_ = D3_DC10(True, self.f38)
                d_777_v83_ = d_777_v83_
                (globalState).f15 = not(p1)
            d_778_v84_: _dafny.Seq
            d_779_v85_: _dafny.Seq
            d_780_v86_: _dafny.Seq
            d_781_v87_: _dafny.Map
            out13_: _dafny.Seq
            out14_: _dafny.Seq
            out15_: _dafny.Seq
            out16_: _dafny.Map
            out13_, out14_, out15_, out16_ = (d_748_v54_).m14(self.f38, globalState)
            d_778_v84_ = out13_
            d_779_v85_ = out14_
            d_780_v86_ = out15_
            d_781_v87_ = out16_
            d_782_v88_: _dafny.Map
            d_782_v88_ = _dafny.Map({p2: p1})
            d_782_v88_ = (d_782_v88_).set(p2, (d_748_v54_).f39)
        if not(not(not(p1))):
            if (self).f28:
                d_783_v89_: _dafny.Array
                def lambda27_(d_784_i4_):
                    return self.f38

                init17_ = lambda27_
                nw120_ = _dafny.Array(None, 20)
                for i0_17_ in range(nw120_.length(0)):
                    nw120_[i0_17_] = init17_(i0_17_)
                d_783_v89_ = nw120_
                d_783_v89_ = d_783_v89_
                d_785_v90_: C0
                nw121_ = C0()
                nw121_.ctor__(p3, p2, d_783_v89_, self.f38)
                d_785_v90_ = nw121_
                d_786_v91_: _dafny.Array
                nw122_ = _dafny.Array(int(0), 26)
                d_786_v91_ = nw122_
                index116_ = default__.safeIndex(477, (d_786_v91_).length(0))
                (d_786_v91_)[index116_] = p2
                index117_ = default__.safeIndex(477, (d_786_v91_).length(0))
                (d_786_v91_)[index117_] = (p2) * (p2)
                (globalState).f4 = (d_786_v91_)[default__.safeIndex(477, (d_786_v91_).length(0))]
                d_787_v92_: _dafny.Map
                d_788_v93_: int
                out17_: _dafny.Map
                out18_: int
                out17_, out18_ = default__.m0(globalState)
                d_787_v92_ = out17_
                d_788_v93_ = out18_
            elif True:
                d_789_v94_: _dafny.Set
                d_789_v94_ = _dafny.Set({p1, (self).f28})
                d_790_v95_: C0
                nw123_ = C0()
                nw123_.ctor__(_dafny.SeqWithoutIsStrInference([(self).f37 for d_791_i5_ in range(default__.abs(257))]), p2, (self).f27, (len(p3)) < (len(d_789_v94_)))
                d_790_v95_ = nw123_
                d_790_v95_ = d_790_v95_
                d_792_v96_: _dafny.Seq
                d_792_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cplehhjuc"))
                d_793_v97_: _dafny.Array
                nw124_ = _dafny.Array(int(0), 23)
                d_793_v97_ = nw124_
                index118_ = default__.safeIndex(335, (d_793_v97_).length(0))
                (d_793_v97_)[index118_] = 855
                d_794_v98_: _dafny.Map
                d_794_v98_ = _dafny.Map({(d_789_v94_).issubset(d_789_v94_): (d_790_v95_).fm15(globalState)})
                index119_ = default__.safeIndex(335, (d_793_v97_).length(0))
                rhs149_ = p2
                rhs150_ = (p3) + ((_dafny.SeqWithoutIsStrInference([(self).f37 for d_795_i6_ in range(default__.abs(74))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([(self).f37 for d_796_i6_ in range(default__.abs(74))]))), (self).f37))
                rhs151_ = ((d_794_v98_)[p1] if (p1) in (d_794_v98_) else p2)
                rhs152_ = p2
                lhs129_ = globalState
                lhs130_ = globalState
                lhs131_ = d_793_v97_
                lhs132_ = default__.safeIndex(335, (d_793_v97_).length(0))
                lhs129_.f4 = rhs149_
                d_792_v96_ = rhs150_
                lhs130_.f13 = rhs151_
                lhs131_[lhs132_] = rhs152_
                d_797_v99_: D9
                d_797_v99_ = D9_DC23(self.f38)
                d_798_v100_: _dafny.MultiSet
                d_798_v100_ = _dafny.MultiSet([D9_DC23(p0), d_797_v99_, d_797_v99_])
                d_799_v101_: _dafny.Seq
                d_799_v101_ = _dafny.SeqWithoutIsStrInference([self.f38])
                d_800_v102_: _dafny.Map
                d_800_v102_ = _dafny.Map({self.f38: (self).f28})
                d_801_v103_: C0
                nw125_ = C0()
                nw125_.ctor__((d_790_v95_).f36, (p2) - (((_dafny.MultiSet([d_798_v100_, d_798_v100_, _dafny.MultiSet([d_797_v99_, D9_DC23((d_799_v101_)[default__.safeIndex(p2, len(d_799_v101_))])])])).set(d_798_v100_, default__.abs((d_793_v97_)[default__.safeIndex(335, (d_793_v97_).length(0))]))).cardinality), (self).f27, (((d_800_v102_)[self.f38] if (self.f38) in (d_800_v102_) else (self).f28)) == ((self).f28))
                d_801_v103_ = nw125_
                (globalState).f21 = p2
                d_802_v104_: _dafny.Seq
                d_802_v104_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])
                (globalState).f17 = ((self).f27) not in ((d_802_v104_) + (d_802_v104_))
            d_803_v105_: D8
            d_803_v105_ = D8_DC20(D8_DC19())
            d_804_v106_: _dafny.Map
            d_804_v106_ = _dafny.Map({_dafny.Map({p0: len(p3)}): p2})
            d_805_v107_: D8
            d_805_v107_ = D8_DC17(d_804_v106_)
            pat_let_tv25_ = d_805_v107_
            d_806_v108_: _dafny.Map
            def iife79_(_pat_let15_0):
                def iife80_(d_807_dt__update__tmp_h1_):
                    def iife81_(_pat_let16_0):
                        def iife82_(d_808_dt__update_hcf32_h0_):
                            return D8_DC20(d_808_dt__update_hcf32_h0_)
                        return iife82_(_pat_let16_0)
                    return iife81_(pat_let_tv25_)
                return iife80_(_pat_let15_0)
            d_806_v108_ = _dafny.Map({iife79_(d_803_v105_): (p3) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxkphh")))})
            (globalState).f4 = len(d_806_v108_)
            if p1:
                d_809_v109_: _dafny.Map
                d_810_v110_: int
                out19_: _dafny.Map
                out20_: int
                out19_, out20_ = default__.m0(globalState)
                d_809_v109_ = out19_
                d_810_v110_ = out20_
                d_811_v111_: _dafny.Seq
                d_811_v111_ = _dafny.SeqWithoutIsStrInference([d_810_v110_])
                d_812_v112_: _dafny.Array
                nw126_ = _dafny.Array(int(0), 25)
                d_812_v112_ = nw126_
                d_813_v113_: _dafny.Seq
                d_813_v113_ = _dafny.SeqWithoutIsStrInference([d_812_v112_, d_812_v112_])
                d_814_v114_: _dafny.MultiSet
                d_814_v114_ = _dafny.MultiSet([(0) - ((d_811_v111_)[default__.safeIndex(p2, len(d_811_v111_))]), (0) - (d_810_v110_), -504, (p2) * (p2), len((d_813_v113_ if (self).f28 else d_813_v113_))])
                (globalState).f21 = ((d_814_v114_)[d_810_v110_] if (d_810_v110_) in (d_814_v114_) else d_810_v110_)
                (globalState).f15 = p0
                d_815_v115_: _dafny.Map
                d_815_v115_ = _dafny.Map({len(p3): (self).f28})
                (globalState).f21 = (len(d_815_v115_)) * (default__.fm36(globalState))
                d_816_v116_: _dafny.Array
                def lambda28_(d_817_i7_):
                    return (self).f37

                init18_ = lambda28_
                nw127_ = _dafny.Array(None, 26)
                for i0_18_ in range(nw127_.length(0)):
                    nw127_[i0_18_] = init18_(i0_18_)
                d_816_v116_ = nw127_
                index120_ = default__.safeIndex(191, (d_816_v116_).length(0))
                (d_816_v116_)[index120_] = (self).f37
                index121_ = default__.safeIndex(191, (d_816_v116_).length(0))
                (d_816_v116_)[index121_] = (self).f37
            elif True:
                d_818_v117_: _dafny.Seq
                d_818_v117_ = _dafny.SeqWithoutIsStrInference([False])
                d_819_v118_: _dafny.Map
                d_819_v118_ = _dafny.Map({d_818_v117_: (self).f28})
                d_820_v119_: _dafny.Map
                d_820_v119_ = _dafny.Map({self.f38: (d_819_v118_) | (d_819_v118_)})
                d_820_v119_ = (d_820_v119_).set((self).f28, d_819_v118_)
                (globalState).f4 = 884
                (globalState).f15 = self.f38
                d_821_v120_: _dafny.Array
                nw128_ = _dafny.Array(int(0), 1)
                d_821_v120_ = nw128_
                index122_ = default__.safeIndex(905, (d_821_v120_).length(0))
                (d_821_v120_)[index122_] = (default__.fm36(globalState)) - (p2)
                index123_ = default__.safeIndex(905, (d_821_v120_).length(0))
                (d_821_v120_)[index123_] = (len(d_818_v117_)) + (p2)
                (globalState).f13 = (d_821_v120_)[default__.safeIndex(905, (d_821_v120_).length(0))]
            (globalState).f21 = (p2) + (p2)
            d_822_v121_: _dafny.Array
            nw129_ = _dafny.Array(int(0), 22)
            d_822_v121_ = nw129_
            d_823_v122_: _dafny.Seq
            d_823_v122_ = _dafny.SeqWithoutIsStrInference([p2])
            d_824_v123_: _dafny.MultiSet
            d_824_v123_ = _dafny.MultiSet([d_823_v122_])
            d_825_v124_: _dafny.Seq
            d_825_v124_ = _dafny.SeqWithoutIsStrInference([p0, p0, p1])
            d_826_v125_: _dafny.Set
            d_826_v125_ = _dafny.Set({(d_825_v124_)[default__.safeIndex(p2, len(d_825_v124_))]})
            index124_ = default__.safeIndex(880, (d_822_v121_).length(0))
            (d_822_v121_)[index124_] = ((d_824_v123_).cardinality) + (len(_dafny.Map({p0: len(d_826_v125_)})))
            d_827_v126_: _dafny.Map
            d_827_v126_ = _dafny.Map({(self).f37: len(p3)})
            index125_ = default__.safeIndex(880, (d_822_v121_).length(0))
            (d_822_v121_)[index125_] = ((d_827_v126_)[(self).f37] if ((self).f37) in (d_827_v126_) else (0) - (p2))
        elif True:
            d_828_v127_: _dafny.Array
            nw130_ = _dafny.Array(_dafny.Map({}), 10)
            d_828_v127_ = nw130_
            d_829_v128_: D0
            d_829_v128_ = D0_DC3(p2, p1, p0)
            pat_let_tv26_ = p1
            d_830_v129_: _dafny.Array
            nw131_ = _dafny.Array(None, 26)
            nw131_[int(0)] = d_829_v128_
            nw131_[int(1)] = d_829_v128_
            nw131_[int(2)] = d_829_v128_
            nw131_[int(3)] = d_829_v128_
            nw131_[int(4)] = d_829_v128_
            def iife83_(_pat_let17_0):
                def iife84_(d_831_dt__update__tmp_h2_):
                    def iife85_(_pat_let18_0):
                        def iife86_(d_832_dt__update_hcf8_h0_):
                            return D0_DC3((d_831_dt__update__tmp_h2_).cf7, d_832_dt__update_hcf8_h0_, (d_831_dt__update__tmp_h2_).cf9)
                        return iife86_(_pat_let18_0)
                    return iife85_(pat_let_tv26_)
                return iife84_(_pat_let17_0)
            nw131_[int(5)] = iife83_(d_829_v128_)
            nw131_[int(6)] = d_829_v128_
            nw131_[int(7)] = D0_DC3(p2, p1, p1)
            nw131_[int(8)] = d_829_v128_
            nw131_[int(9)] = d_829_v128_
            nw131_[int(10)] = d_829_v128_
            nw131_[int(11)] = d_829_v128_
            nw131_[int(12)] = d_829_v128_
            nw131_[int(13)] = d_829_v128_
            nw131_[int(14)] = d_829_v128_
            nw131_[int(15)] = d_829_v128_
            nw131_[int(16)] = D0_DC3(p2, p0, p1)
            nw131_[int(17)] = d_829_v128_
            nw131_[int(18)] = D0_DC3(p2, p0, p0)
            nw131_[int(19)] = d_829_v128_
            nw131_[int(20)] = d_829_v128_
            nw131_[int(21)] = d_829_v128_
            nw131_[int(22)] = d_829_v128_
            nw131_[int(23)] = d_829_v128_
            nw131_[int(24)] = D0_DC3(p2, (self).f28, p0)
            nw131_[int(25)] = d_829_v128_
            d_830_v129_ = nw131_
            d_833_v130_: _dafny.Map
            d_833_v130_ = _dafny.Map({_dafny.Map({(0) - ((0) - (p2)): p1}): d_830_v129_})
            index126_ = default__.safeIndex(113, (d_828_v127_).length(0))
            (d_828_v127_)[index126_] = d_833_v130_
            index127_ = default__.safeIndex(113, (d_828_v127_).length(0))
            (d_828_v127_)[index127_] = d_833_v130_
            index128_ = default__.safeIndex(183, ((self).f27).length(0))
            ((self).f27)[index128_] = (self.f38) and ((self).f28)
            index129_ = default__.safeIndex(183, ((self).f27).length(0))
            ((self).f27)[index129_] = (p2) == ((0) - ((829) * (p2)))
            d_834_v131_: _dafny.Seq
            d_834_v131_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqp"))
            d_835_v132_: _dafny.Array
            nw132_ = _dafny.Array(None, 8)
            nw132_[int(0)] = p1
            nw132_[int(1)] = (self).f28
            nw132_[int(2)] = (self).f28
            nw132_[int(3)] = p1
            nw132_[int(4)] = p1
            nw132_[int(5)] = ((self).f27)[default__.safeIndex(183, ((self).f27).length(0))]
            nw132_[int(6)] = ((self).f27)[default__.safeIndex(183, ((self).f27).length(0))]
            nw132_[int(7)] = p0
            d_835_v132_ = nw132_
            d_836_v133_: T1
            nw133_ = C0()
            nw133_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")), p2, d_835_v132_, ((self).f27)[default__.safeIndex(183, ((self).f27).length(0))])
            d_836_v133_ = nw133_
            d_837_v134_: _dafny.Map
            d_837_v134_ = _dafny.Map({p1: d_836_v133_})
            rhs153_ = (p2 if p0 else (-641 if True else len(d_837_v134_)))
            rhs154_ = 187
            rhs155_ = p3
            rhs156_ = 997
            lhs133_ = globalState
            lhs134_ = globalState
            lhs135_ = globalState
            lhs133_.f5 = rhs153_
            lhs134_.f13 = rhs154_
            d_834_v131_ = rhs155_
            lhs135_.f21 = rhs156_
            d_838_v135_: _dafny.Array
            def lambda29_(d_839_v133_):
                def lambda30_(d_840_i8_):
                    return (d_840_i8_) - ((0) - (d_839_v133_.f29))

                return lambda30_

            init19_ = lambda29_(d_836_v133_)
            nw134_ = _dafny.Array(None, 16)
            for i0_19_ in range(nw134_.length(0)):
                nw134_[i0_19_] = init19_(i0_19_)
            d_838_v135_ = nw134_
            d_838_v135_ = d_838_v135_
            d_841_v136_: D8
            d_841_v136_ = D8_DC18(d_836_v133_.f29)
            source11_ = d_841_v136_
            if source11_.is_DC18:
                d_842___mcc_h4_ = source11_.cf31
                d_843_cf31_ = d_842___mcc_h4_
                (globalState).f6 = (d_829_v128_).cf8
                d_844_v137_: _dafny.MultiSet
                d_844_v137_ = _dafny.MultiSet([(d_836_v133_).f28])
                d_845_v138_: _dafny.Map
                d_845_v138_ = _dafny.Map({p2: (d_844_v137_).cardinality})
                d_846_v139_: _dafny.Set
                d_846_v139_ = _dafny.Set({((d_836_v133_).f27), d_835_v132_})
                d_845_v138_ = (d_845_v138_).set(d_836_v133_.f29, len(d_846_v139_))
                (globalState).f21 = default__.fm36(globalState)
                d_847_v140_: _dafny.Seq
                d_847_v140_ = _dafny.SeqWithoutIsStrInference([(d_836_v133_).f28, not((d_836_v133_).f28), ((self).f27)[default__.safeIndex(183, ((self).f27).length(0))], True])
                d_848_v141_: _dafny.Map
                d_848_v141_ = _dafny.Map({p3: (d_847_v140_).set(default__.safeIndex(d_836_v133_.f29, len(d_847_v140_)), p1)})
                d_849_v142_: _dafny.MultiSet
                d_849_v142_ = _dafny.MultiSet([(d_836_v133_).f27])
                rhs157_ = (d_849_v142_).issubset((d_849_v142_).intersection(d_849_v142_))
                rhs158_ = d_848_v141_
                rhs159_ = (((0) - (d_836_v133_.f29)) - (d_836_v133_.f29)) * (d_843_cf31_)
                lhs136_ = globalState
                lhs137_ = globalState
                lhs136_.f17 = rhs157_
                d_848_v141_ = rhs158_
                lhs137_.f21 = rhs159_
            elif source11_.is_DC19:
                (globalState).f13 = p2
                (d_836_v133_).f29 = p2
                d_850_v143_: C1
                nw135_ = C1()
                nw135_.ctor__((d_836_v133_).f28, (d_836_v133_).f27, p0)
                d_850_v143_ = nw135_
                (globalState).f4 = d_836_v133_.f29
            elif source11_.is_DC17:
                d_851___mcc_h5_ = source11_.cf30
                d_852_cf30_ = d_851___mcc_h5_
                d_853_v144_: _dafny.Array
                nw136_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_853_v144_ = nw136_
                index130_ = default__.safeIndex(87, (d_853_v144_).length(0))
                (d_853_v144_)[index130_] = d_838_v135_
                index131_ = default__.safeIndex(87, (d_853_v144_).length(0))
                (d_853_v144_)[index131_] = d_838_v135_
                d_854_v145_: _dafny.Map
                d_854_v145_ = _dafny.Map({p0: (d_853_v144_)[default__.safeIndex(87, (d_853_v144_).length(0))]})
                d_855_v146_: _dafny.Map
                d_855_v146_ = _dafny.Map({d_836_v133_.f29: ((d_854_v145_)[((self).f27)[default__.safeIndex(183, ((self).f27).length(0))]] if (((self).f27)[default__.safeIndex(183, ((self).f27).length(0))]) in (d_854_v145_) else d_838_v135_)})
                d_855_v146_ = (d_855_v146_).set((self).fm20(globalState), (d_853_v144_)[default__.safeIndex(87, (d_853_v144_).length(0))])
                d_856_v147_: _dafny.Seq
                d_856_v147_ = _dafny.SeqWithoutIsStrInference([(d_836_v133_).f28])
                d_857_v148_: _dafny.Set
                d_857_v148_ = _dafny.Set({len((d_856_v147_) + (d_856_v147_)), d_836_v133_.f29})
                rhs160_ = d_857_v148_
                rhs161_ = (690) - (p2)
                rhs162_ = (d_834_v131_)[default__.safeIndex(p2, len(d_834_v131_))]
                lhs138_ = globalState
                lhs139_ = globalState
                d_857_v148_ = rhs160_
                lhs138_.f5 = rhs161_
                lhs139_.f25 = rhs162_
                (globalState).f13 = (len(_dafny.Map({d_836_v133_: len(p3)}))) - (p2)
            elif True:
                d_858___mcc_h6_ = source11_.cf32
                d_859_cf32_ = d_858___mcc_h6_
                d_860_v149_: T0
                nw137_ = C1()
                nw137_.ctor__(True, (d_836_v133_).f27, p1)
                d_860_v149_ = nw137_
                d_861_v150_: _dafny.Map
                d_861_v150_ = _dafny.Map({p1: d_860_v149_})
                d_862_v151_: _dafny.Map
                d_862_v151_ = _dafny.Map({default__.safeDivisionInt(p2, d_836_v133_.f29): ((d_861_v150_)[p1] if (p1) in (d_861_v150_) else d_860_v149_)})
                d_862_v151_ = (d_862_v151_).set(d_836_v133_.f29, d_860_v149_)
                d_863_v152_: _dafny.Array
                nw138_ = _dafny.Array(_dafny.Map({}), 29)
                d_863_v152_ = nw138_
                d_864_v153_: _dafny.Set
                d_864_v153_ = _dafny.Set({d_836_v133_.f29})
                index132_ = default__.safeIndex(165, (d_863_v152_).length(0))
                (d_863_v152_)[index132_] = _dafny.Map({d_864_v153_: d_836_v133_.f29})
                d_865_v154_: _dafny.Map
                d_865_v154_ = _dafny.Map({True: (self).f28})
                d_866_v155_: _dafny.Map
                d_866_v155_ = _dafny.Map({len(d_865_v154_): d_836_v133_.f29})
                d_867_v156_: _dafny.Map
                d_867_v156_ = _dafny.Map({True: d_836_v133_.f29})
                d_868_v157_: _dafny.Map
                d_868_v157_ = _dafny.Map({(_dafny.Set({826, p2, len(d_866_v155_), p2, (0) - (p2)})) | (d_864_v153_): default__.safeModuloInt(len(d_867_v156_), p2)})
                index133_ = default__.safeIndex(165, (d_863_v152_).length(0))
                (d_863_v152_)[index133_] = d_868_v157_
                d_869_v158_: _dafny.Array
                nw139_ = _dafny.Array(None, 11)
                nw139_[int(0)] = (d_860_v149_).f27
                nw139_[int(1)] = (d_836_v133_).f27
                nw139_[int(2)] = (d_836_v133_).f27
                nw139_[int(3)] = (d_836_v133_).f27
                nw139_[int(4)] = (d_836_v133_).f27
                nw139_[int(5)] = (d_836_v133_).f27
                nw139_[int(6)] = (d_860_v149_).f27
                nw139_[int(7)] = (d_860_v149_).f27
                nw139_[int(8)] = (d_836_v133_).f27
                nw139_[int(9)] = (d_860_v149_).f27
                nw139_[int(10)] = (d_836_v133_).f27
                d_869_v158_ = nw139_
                d_870_v159_: D9
                d_870_v159_ = D9_DC22(d_838_v135_, False, d_869_v158_, d_834_v131_)
                d_867_v156_ = (d_867_v156_).set((d_870_v159_).cf35, (d_836_v133_.f29) * (d_836_v133_.f29))
                d_871_v160_: _dafny.Map
                d_871_v160_ = _dafny.Map({(d_860_v149_).f28: p3})
                d_872_v161_: _dafny.Seq
                d_872_v161_ = _dafny.SeqWithoutIsStrInference([d_834_v131_, ((d_871_v160_)[(d_836_v133_).f28] if ((d_836_v133_).f28) in (d_871_v160_) else _dafny.SeqWithoutIsStrInference([(self).f37 for d_873_i9_ in range(default__.abs(-306))]))])
                d_874_v162_: _dafny.Seq
                d_874_v162_ = _dafny.SeqWithoutIsStrInference([p2])
                d_875_v163_: D1
                d_875_v163_ = D1_DC7((self).f37, d_834_v131_)
                d_876_v164_: _dafny.Array
                nw140_ = _dafny.Array(None, 25)
                nw140_[int(0)] = d_834_v131_
                nw140_[int(1)] = default__.fm32((0) - (d_836_v133_.f29), (d_836_v133_).f28, globalState)
                nw140_[int(2)] = (d_872_v161_)[default__.safeIndex(len(d_874_v162_), len(d_872_v161_))]
                nw140_[int(3)] = p3
                nw140_[int(4)] = d_834_v131_
                nw140_[int(5)] = p3
                nw140_[int(6)] = d_834_v131_
                nw140_[int(7)] = d_834_v131_
                nw140_[int(8)] = (d_875_v163_).cf14
                nw140_[int(9)] = d_834_v131_
                nw140_[int(10)] = default__.fm32(d_836_v133_.f29, p0, globalState)
                nw140_[int(11)] = (p3) + (p3)
                nw140_[int(12)] = p3
                nw140_[int(13)] = p3
                nw140_[int(14)] = (default__.fm32(d_836_v133_.f29, (self).f28, globalState)) + (d_834_v131_)
                nw140_[int(15)] = p3
                nw140_[int(16)] = d_834_v131_
                nw140_[int(17)] = p3
                nw140_[int(18)] = d_834_v131_
                nw140_[int(19)] = p3
                nw140_[int(20)] = (d_872_v161_)[default__.safeIndex(612, len(d_872_v161_))]
                nw140_[int(21)] = d_834_v131_
                nw140_[int(22)] = p3
                nw140_[int(23)] = (d_834_v131_) + (p3)
                nw140_[int(24)] = (d_834_v131_) + (d_834_v131_)
                d_876_v164_ = nw140_
                index134_ = default__.safeIndex(792, (d_876_v164_).length(0))
                (d_876_v164_)[index134_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obsvgeys"))
                index135_ = default__.safeIndex(792, (d_876_v164_).length(0))
                (d_876_v164_)[index135_] = d_834_v131_
        d_877_v165_: _dafny.Seq
        d_877_v165_ = _dafny.SeqWithoutIsStrInference([8, p2, p2])
        d_878_v166_: _dafny.Seq
        d_878_v166_ = _dafny.SeqWithoutIsStrInference([p2 for d_879_i10_ in range(default__.abs(378))])
        d_877_v165_ = (d_877_v165_) + ((d_877_v165_ if not((self).f28) else (d_878_v166_)))
        (globalState).f6 = not (self.f38) or ((p0) or ((self).f28))
        (globalState).f4 = p2
        r0 = default__.fm44(globalState)
        return r0

    @property
    def f37(self):
        return self._f37

class C3(T1, T0):
    def  __init__(self):
        self._f29: int = int(0)
        self._f27: _dafny.Array = _dafny.Array(None, 0)
        self._f28: bool = False
        self._f34: int = int(0)
        self._f35: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f34, f35, f29, f27, f28):
        (self)._f34 = f34
        (self)._f35 = f35
        (self).f29 = f29
        (self)._f27 = f27
        (self)._f28 = f28

    def fm14(self, p0, p1, globalState):
        return ((self.f29) * ((self).f34)) - (self.f29)

    def m10(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        r2: str = _dafny.CodePoint('D')
        r3: int = int(0)
        d_880_v0_: _dafny.Seq
        d_880_v0_ = _dafny.SeqWithoutIsStrInference([(self).f27])
        (globalState).f17 = not((d_880_v0_) <= ((d_880_v0_) + (d_880_v0_)))
        d_881_v1_: _dafny.Map
        d_881_v1_ = _dafny.Map({(self).f28: (self).f28})
        d_882_v2_: C0
        nw141_ = C0()
        nw141_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phmmcxi")), (self).f34, (self).f27, default__.fm16(d_881_v1_, globalState))
        d_882_v2_ = nw141_
        d_883_v3_: _dafny.Seq
        d_883_v3_ = _dafny.SeqWithoutIsStrInference([(self).f28, (self).f28, p0])
        d_884_v4_: _dafny.Seq
        d_884_v4_ = _dafny.SeqWithoutIsStrInference([(d_883_v3_)[default__.safeIndex(self.f29, len(d_883_v3_))]])
        d_885_v5_: _dafny.Seq
        d_885_v5_ = _dafny.SeqWithoutIsStrInference([d_884_v4_])
        d_886_v6_: _dafny.Seq
        d_886_v6_ = _dafny.SeqWithoutIsStrInference([(self).f34])
        d_887_v7_: str
        d_887_v7_ = _dafny.CodePoint('p')
        d_888_v8_: T1
        nw142_ = C0()
        nw142_.ctor__((d_882_v2_).f36, self.f29, (self).f27, p0)
        d_888_v8_ = nw142_
        d_889_v9_: _dafny.Map
        d_889_v9_ = _dafny.Map({((d_882_v2_).f36).set(default__.safeIndex((d_886_v6_)[default__.safeIndex(self.f29, len(d_886_v6_))], len((d_882_v2_).f36)), d_887_v7_): d_888_v8_})
        d_890_v10_: _dafny.Set
        d_890_v10_ = _dafny.Set({((d_889_v9_)[(d_882_v2_).f36] if ((d_882_v2_).f36) in (d_889_v9_) else d_888_v8_)})
        d_891_v11_: _dafny.Map
        d_891_v11_ = _dafny.Map({d_885_v5_: d_890_v10_})
        d_891_v11_ = (d_891_v11_).set(d_885_v5_, d_890_v10_)
        d_892_v12_: _dafny.Seq
        d_892_v12_ = d_886_v6_
        d_893_v13_: _dafny.Seq
        d_893_v13_ = _dafny.SeqWithoutIsStrInference([d_886_v6_, d_886_v6_, (d_892_v12_), d_886_v6_])
        d_894_v14_: D1
        d_894_v14_ = D1_DC5(d_893_v13_)
        d_895_v15_: _dafny.MultiSet
        d_895_v15_ = _dafny.MultiSet([D1_DC5(default__.fm17(p0, (d_888_v8_).f28, len((d_882_v2_).f36), p0, globalState)), D1_DC5(d_893_v13_), d_894_v14_, D1_DC5(d_893_v13_), d_894_v14_])
        d_896_v16_: _dafny.MultiSet
        d_896_v16_ = _dafny.MultiSet([d_895_v15_])
        d_896_v16_ = ((d_896_v16_).set(d_895_v15_, default__.abs((0) - (len((d_882_v2_).f36))))).intersection(d_896_v16_)
        d_897_v17_: D0
        d_897_v17_ = D0_DC3(d_888_v8_.f29, not((self).f28), True)
        d_898_v18_: _dafny.Seq
        d_898_v18_ = _dafny.SeqWithoutIsStrInference([d_897_v17_])
        d_899_v19_: _dafny.Array
        def lambda31_(d_900_v4_):
            def lambda32_(d_901_i0_):
                return (d_901_i0_) + (len(d_900_v4_))

            return lambda32_

        init20_ = lambda31_(d_884_v4_)
        nw143_ = _dafny.Array(None, 14)
        for i0_20_ in range(nw143_.length(0)):
            nw143_[i0_20_] = init20_(i0_20_)
        d_899_v19_ = nw143_
        d_902_v20_: _dafny.Map
        d_902_v20_ = _dafny.Map({(len(d_898_v18_)) * ((self).f34): d_899_v19_})
        d_902_v20_ = d_902_v20_
        r0 = not((self).f28)
        d_903_v21_: _dafny.MultiSet
        d_903_v21_ = _dafny.MultiSet([(d_888_v8_).f28, (self).f28])
        d_904_v22_: _dafny.Seq
        d_904_v22_ = _dafny.SeqWithoutIsStrInference([d_903_v21_])
        d_905_v23_: _dafny.Seq
        d_905_v23_ = d_904_v22_
        def iife87_():
            coll49_ = _dafny.Set()
            compr_49_: _dafny.MultiSet
            for compr_49_ in ((d_905_v23_)).Elements:
                d_906_v24_: _dafny.MultiSet = compr_49_
                if (d_906_v24_) in ((d_905_v23_)):
                    coll49_ = coll49_.union(_dafny.Set([d_906_v24_]))
            return _dafny.Set(coll49_)
        r0 = ((self).fm14(len(d_886_v6_), iife87_()
        , globalState)) <= ((d_886_v6_)[default__.safeIndex((self).f34, len(d_886_v6_))])
        d_907_v25_: _dafny.Map
        d_907_v25_ = _dafny.Map({(self).f28: (self).f34})
        d_908_v26_: _dafny.Map
        d_908_v26_ = _dafny.Map({((d_907_v25_)[p0] if (p0) in (d_907_v25_) else d_888_v8_.f29): (self).f34})
        r1 = (len((d_884_v4_) + (d_883_v3_))) <= (((d_908_v26_)[159] if (159) in (d_908_v26_) else 241))
        r2 = d_887_v7_
        r3 = (self).f34
        return r0, r1, r2, r3

    def m11(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: bool = False
        d_909_v0_: _dafny.Map
        d_909_v0_ = _dafny.Map({(self).f34: (self).f28})
        d_910_v1_: _dafny.Map
        d_910_v1_ = _dafny.Map({(d_909_v0_).set((self).f34, p0): (self).f34})
        d_911_v2_: _dafny.MultiSet
        d_911_v2_ = _dafny.MultiSet([True])
        d_912_v3_: _dafny.Set
        d_912_v3_ = _dafny.Set({d_911_v2_})
        (globalState).f13 = (self).fm14((self).fm14(((d_910_v1_)[d_909_v0_] if (d_909_v0_) in (d_910_v1_) else self.f29), d_912_v3_, globalState), d_912_v3_, globalState)
        hi2_ = (self).f34
        for d_913_i0_ in range((self).f34, hi2_):
            d_914_v4_: _dafny.Seq
            d_914_v4_ = _dafny.SeqWithoutIsStrInference([d_913_i0_, (self).f34, 137, -189, d_913_i0_])
            d_915_v5_: _dafny.Seq
            d_915_v5_ = d_914_v4_
            source12_ = d_915_v5_
            d_916___mcc_h0_ = source12_
            d_917_cf15_ = d_916___mcc_h0_
            d_918_v6_: _dafny.Array
            def lambda33_(d_919_v4_):
                def lambda34_(d_920_i1_):
                    return (d_920_i1_) - (len(d_919_v4_))

                return lambda34_

            init21_ = lambda33_(d_914_v4_)
            nw144_ = _dafny.Array(None, 25)
            for i0_21_ in range(nw144_.length(0)):
                nw144_[i0_21_] = init21_(i0_21_)
            d_918_v6_ = nw144_
            index136_ = default__.safeIndex(590, (d_918_v6_).length(0))
            (d_918_v6_)[index136_] = (self).f34
            index137_ = default__.safeIndex(590, (d_918_v6_).length(0))
            (d_918_v6_)[index137_] = self.f29
            (globalState).f25 = _dafny.CodePoint('h')
            d_921_v7_: D1
            d_921_v7_ = D1_DC6(d_913_i0_)
            (globalState).f21 = (0) - ((d_921_v7_).cf12)
            d_922_v8_: _dafny.Map
            d_922_v8_ = _dafny.Map({d_921_v7_: d_913_i0_})
            d_922_v8_ = (d_922_v8_).set(d_921_v7_, self.f29)
            (globalState).f6 = (d_913_i0_) > (self.f29)
            index138_ = default__.safeIndex(874, ((self).f27).length(0))
            ((self).f27)[index138_] = p0
            index139_ = default__.safeIndex(874, ((self).f27).length(0))
            ((self).f27)[index139_] = (self).f28
            index140_ = default__.safeIndex(874, ((self).f27).length(0))
            ((self).f27)[index140_] = False
        d_923_v9_: _dafny.Array
        nw145_ = _dafny.Array(_dafny.CodePoint('D'), 19)
        d_923_v9_ = nw145_
        d_924_v10_: _dafny.Map
        d_924_v10_ = _dafny.Map({d_923_v9_: (((d_909_v0_)[self.f29] if (self.f29) in (d_909_v0_) else (self).f28)) in (_dafny.SeqWithoutIsStrInference([False]))})
        d_924_v10_ = (d_924_v10_).set(d_923_v9_, not((self).f28))
        d_925_v11_: _dafny.Map
        d_925_v11_ = _dafny.Map({True: p0})
        d_926_i2_: int
        d_926_i2_ = 0
        with _dafny.label("2"):
            while not((p0 if default__.fm16(d_925_v11_, globalState) else (self).f28)):
                with _dafny.c_label("2"):
                    if (d_926_i2_) >= (100):
                        raise _dafny.Break("2")
                    d_926_i2_ = (d_926_i2_) + (1)
                    if not(p0):
                        d_927_v12_: D3
                        d_927_v12_ = D3_DC10((self).f28, p0)
                        d_909_v0_ = (d_909_v0_).set((0) - (len(_dafny.Map({(d_927_v12_).cf18: _dafny.MultiSet([(self).f34])}))), p0)
                        d_928_v13_: str
                        d_928_v13_ = _dafny.CodePoint('e')
                        d_929_v14_: C0
                        nw146_ = C0()
                        nw146_.ctor__(default__.fm18(p0, (self).f34, d_928_v13_, globalState), (0) - ((self).f34), (self).f27, not ((self).f28) or (p0))
                        d_929_v14_ = nw146_
                        d_930_v15_: _dafny.Seq
                        d_930_v15_ = _dafny.SeqWithoutIsStrInference([p0])
                        d_931_v16_: _dafny.Map
                        d_931_v16_ = _dafny.Map({(self).f28: (_dafny.MultiSet([len(d_930_v15_)])).cardinality})
                        d_932_v17_: _dafny.Set
                        d_932_v17_ = _dafny.Set({d_931_v16_, d_931_v16_, d_931_v16_, d_931_v16_})
                        d_933_v18_: _dafny.Map
                        d_933_v18_ = _dafny.Map({704: d_932_v17_})
                        d_932_v17_ = (((d_933_v18_)[self.f29] if (self.f29) in (d_933_v18_) else _dafny.Set({d_931_v16_}))) | ((_dafny.Set({d_931_v16_, d_931_v16_, d_931_v16_})) | (d_932_v17_))
                        d_934_v19_: _dafny.Seq
                        d_934_v19_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])
                        d_935_v20_: T0
                        nw147_ = C1()
                        nw147_.ctor__((self).f28, (d_934_v19_)[default__.safeIndex(self.f29, len(d_934_v19_))], p0)
                        d_935_v20_ = nw147_
                        d_936_v21_: _dafny.Array
                        d_936_v21_ = (self).f27
                        d_937_v22_: _dafny.Map
                        d_937_v22_ = _dafny.Map({(self).f34: d_936_v21_})
                        d_938_v23_: _dafny.Set
                        d_938_v23_ = _dafny.Set({d_937_v22_, d_937_v22_})
                        rhs163_ = ((d_930_v15_) + (_dafny.SeqWithoutIsStrInference([p0]))) + (_dafny.SeqWithoutIsStrInference([(self).f28, p0]))
                        rhs164_ = not(((d_938_v23_) - (d_938_v23_)).issubset(d_938_v23_))
                        rhs165_ = len((d_929_v14_).f36)
                        rhs166_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_939_i3_ in range(default__.abs(908))])
                        rhs167_ = (d_935_v20_ if not (p0) or ((self).f28) else d_935_v20_)
                        lhs140_ = self
                        d_930_v15_ = rhs163_
                        r2 = rhs164_
                        lhs140_.f29 = rhs165_
                        r1 = rhs166_
                        d_935_v20_ = rhs167_
                        d_940_v24_: _dafny.MultiSet
                        d_940_v24_ = _dafny.MultiSet([(self).f34])
                        d_940_v24_ = (d_940_v24_).intersection(d_940_v24_)
                    elif True:
                        (globalState).f21 = ((self).f34) + ((783) + (self.f29))
                        (globalState).f6 = (self).f28
                        d_941_v25_: _dafny.Array
                        def lambda35_(d_942_v11_, d_943_p0_):
                            def lambda36_(d_944_i4_):
                                return (_dafny.Set({((d_942_v11_)[(self).f28] if ((self).f28) in (d_942_v11_) else d_943_p0_), ((d_942_v11_)[(self).f28] if ((self).f28) in (d_942_v11_) else (self).f28)})) | (_dafny.Set({(self).f28}))

                            return lambda36_

                        init22_ = lambda35_(d_925_v11_, p0)
                        nw148_ = _dafny.Array(None, 22)
                        for i0_22_ in range(nw148_.length(0)):
                            nw148_[i0_22_] = init22_(i0_22_)
                        d_941_v25_ = nw148_
                        d_945_v26_: _dafny.Set
                        d_945_v26_ = _dafny.Set({(self).f28})
                        index141_ = default__.safeIndex(549, (d_941_v25_).length(0))
                        (d_941_v25_)[index141_] = d_945_v26_
                        index142_ = default__.safeIndex(549, (d_941_v25_).length(0))
                        (d_941_v25_)[index142_] = d_945_v26_
                        (globalState).f6 = p0
                        (globalState).f13 = self.f29
                    r2 = not((self).f28)
                    index143_ = default__.safeIndex(809, ((self).f27).length(0))
                    ((self).f27)[index143_] = (p0 if p0 else (self).f28)
                    index144_ = default__.safeIndex(809, ((self).f27).length(0))
                    ((self).f27)[index144_] = p0
                    d_946_v27_: _dafny.Array
                    nw149_ = _dafny.Array(_dafny.Array(None, 0), 14)
                    d_946_v27_ = nw149_
                    d_947_v28_: _dafny.Array
                    nw150_ = _dafny.Array(_dafny.Array(None, 0), 13)
                    d_947_v28_ = nw150_
                    index145_ = default__.safeIndex(120, (d_946_v27_).length(0))
                    (d_946_v27_)[index145_] = d_947_v28_
                    d_948_v29_: _dafny.Seq
                    d_948_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djlpjj"))
                    index146_ = default__.safeIndex(120, (d_946_v27_).length(0))
                    index147_ = default__.safeIndex(809, ((self).f27).length(0))
                    rhs168_ = (0) - (len(d_948_v29_))
                    rhs169_ = d_947_v28_
                    rhs170_ = ((self).f27)[default__.safeIndex(809, ((self).f27).length(0))]
                    lhs141_ = globalState
                    lhs142_ = d_946_v27_
                    lhs143_ = default__.safeIndex(120, (d_946_v27_).length(0))
                    lhs144_ = (self).f27
                    lhs145_ = default__.safeIndex(809, ((self).f27).length(0))
                    lhs141_.f5 = rhs168_
                    lhs142_[lhs143_] = rhs169_
                    lhs144_[lhs145_] = rhs170_
                    pass
            pass
        d_949_v30_: _dafny.Seq
        d_949_v30_ = _dafny.SeqWithoutIsStrInference([p0])
        d_950_v31_: _dafny.Array
        nw151_ = _dafny.Array(None, 1)
        nw151_[int(0)] = (_dafny.SeqWithoutIsStrInference([(self).f28, p0])) + (d_949_v30_)
        d_950_v31_ = nw151_
        index148_ = default__.safeIndex(427, (d_950_v31_).length(0))
        (d_950_v31_)[index148_] = d_949_v30_
        index149_ = default__.safeIndex(427, (d_950_v31_).length(0))
        (d_950_v31_)[index149_] = (((d_949_v30_).set(default__.safeIndex(self.f29, len(d_949_v30_)), (self).f28)) + (d_949_v30_)) + (d_949_v30_)
        d_951_v32_: _dafny.Seq
        d_951_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vn"))
        d_952_v33_: _dafny.Seq
        d_952_v33_ = _dafny.SeqWithoutIsStrInference([self.f29])
        d_953_v34_: _dafny.Map
        d_953_v34_ = _dafny.Map({(0) - (len(d_951_v32_)): d_952_v33_})
        d_954_v35_: _dafny.Seq
        d_954_v35_ = ((d_953_v34_)[93] if (93) in (d_953_v34_) else d_952_v33_)
        def lambda37_(source14_):
            d_955___mcc_h4_ = source14_
            d_956_cf15_ = d_955___mcc_h4_
            return D8_DC20(D8_DC20(D8_DC20(D8_DC18((self).f34))))

        source13_ = lambda37_(d_954_v35_)
        if source13_.is_DC18:
            d_957___mcc_h1_ = source13_.cf31
            d_958_cf31_ = d_957___mcc_h1_
            d_959_v36_: D1
            d_959_v36_ = D1_DC6((self).f34)
            d_960_v37_: _dafny.Map
            d_960_v37_ = _dafny.Map({(self).f34: self.f29})
            d_961_v38_: _dafny.MultiSet
            d_961_v38_ = _dafny.MultiSet([d_958_cf31_])
            d_962_v39_: _dafny.Set
            d_962_v39_ = _dafny.Set({d_958_cf31_, (0) - (len(d_951_v32_))})
            d_963_v40_: _dafny.Map
            d_963_v40_ = _dafny.Map({d_962_v39_: p0})
            d_964_v41_: C0
            nw152_ = C0()
            nw152_.ctor__(d_951_v32_, d_958_cf31_, (self).f27, (self).f28)
            d_964_v41_ = nw152_
            d_965_v42_: _dafny.Seq
            d_965_v42_ = _dafny.SeqWithoutIsStrInference([d_964_v41_, d_964_v41_])
            d_966_v43_: str
            d_966_v43_ = _dafny.CodePoint('y')
            d_967_v44_: _dafny.Map
            d_967_v44_ = _dafny.Map({d_966_v43_: (self).f28})
            d_968_v45_: _dafny.Array
            nw153_ = _dafny.Array(None, 25)
            nw153_[int(0)] = (self).f34
            nw153_[int(1)] = (d_959_v36_).cf12
            nw153_[int(2)] = (self).fm14(len((d_960_v37_).set((d_961_v38_).cardinality, self.f29)), d_912_v3_, globalState)
            nw153_[int(3)] = default__.safeDivisionInt(len(d_951_v32_), len(d_952_v33_))
            nw153_[int(4)] = d_958_cf31_
            nw153_[int(5)] = (0) - (d_958_cf31_)
            nw153_[int(6)] = d_958_cf31_
            nw153_[int(7)] = (self).f34
            nw153_[int(8)] = (0) - (len(d_909_v0_))
            nw153_[int(9)] = default__.fm36(globalState)
            nw153_[int(10)] = d_958_cf31_
            nw153_[int(11)] = len(d_951_v32_)
            nw153_[int(12)] = (self).f34
            nw153_[int(13)] = (0) - ((self).f34)
            nw153_[int(14)] = default__.fm36(globalState)
            nw153_[int(15)] = -535
            nw153_[int(16)] = len(d_963_v40_)
            nw153_[int(17)] = (0) - (len(d_951_v32_))
            nw153_[int(18)] = d_958_cf31_
            nw153_[int(19)] = (self).f34
            nw153_[int(20)] = default__.safeDivisionInt(self.f29, len(d_925_v11_))
            nw153_[int(21)] = d_958_cf31_
            nw153_[int(22)] = (self).f34
            nw153_[int(23)] = len(d_965_v42_)
            nw153_[int(24)] = (default__.fm45(877, p0, self.f29, ((d_967_v44_)[d_966_v43_] if (d_966_v43_) in (d_967_v44_) else not(p0)), globalState)).cardinality
            d_968_v45_ = nw153_
            index150_ = default__.safeIndex(679, (d_968_v45_).length(0))
            (d_968_v45_)[index150_] = d_958_cf31_
            index151_ = default__.safeIndex(679, (d_968_v45_).length(0))
            (d_968_v45_)[index151_] = self.f29
            (globalState).f15 = (self).f28
            index152_ = default__.safeIndex(679, (d_968_v45_).length(0))
            (d_968_v45_)[index152_] = (self.f29) - ((self).f34)
            (globalState).f15 = (self).f28
        elif source13_.is_DC19:
            (globalState).f5 = default__.fm36(globalState)
            (globalState).f15 = False
            d_969_v46_: _dafny.MultiSet
            d_969_v46_ = _dafny.MultiSet([d_951_v32_, d_951_v32_])
            d_970_v47_: _dafny.Map
            d_970_v47_ = _dafny.Map({(d_969_v46_).cardinality: (self).f34})
            d_971_v48_: _dafny.Seq
            d_971_v48_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uk"))])
            d_972_v49_: str
            d_972_v49_ = _dafny.CodePoint('l')
            d_970_v47_ = default__.fm46((d_971_v48_) < (_dafny.SeqWithoutIsStrInference([d_951_v32_, default__.fm18(p0, 22, d_972_v49_, globalState)])), globalState)
            index153_ = default__.safeIndex(427, (d_950_v31_).length(0))
            (d_950_v31_)[index153_] = (d_950_v31_)[default__.safeIndex(427, (d_950_v31_).length(0))]
        elif source13_.is_DC17:
            d_973___mcc_h2_ = source13_.cf30
            d_974_cf30_ = d_973___mcc_h2_
            (globalState).f25 = default__.fm42((self.f29) - (len(d_951_v32_)), globalState)
            (globalState).f17 = default__.fm40(globalState)
            if (d_951_v32_) != (d_951_v32_):
                d_975_v50_: _dafny.Set
                d_975_v50_ = _dafny.Set({(self).f28, p0, default__.fm40(globalState)})
                d_976_v51_: _dafny.Set
                d_976_v51_ = _dafny.Set({(d_975_v50_).isdisjoint(d_975_v50_)})
                d_976_v51_ = d_976_v51_
                (globalState).f17 = default__.fm16(d_925_v11_, globalState)
                (globalState).f15 = (self).f28
                (globalState).f5 = (self.f29 if p0 else self.f29)
                d_977_v52_: _dafny.Array
                nw154_ = _dafny.Array(D9.default()(), 24)
                d_977_v52_ = nw154_
                d_978_v53_: str
                d_978_v53_ = _dafny.CodePoint('a')
                d_979_v54_: _dafny.Set
                d_979_v54_ = _dafny.Set({d_978_v53_})
                d_980_v55_: D1
                d_980_v55_ = D1_DC6(len(d_979_v54_))
                d_981_v56_: _dafny.Map
                d_981_v56_ = _dafny.Map({d_977_v52_: d_980_v55_})
                d_982_v57_: _dafny.Map
                d_982_v57_ = _dafny.Map({(597) * ((0) - ((self).f34)): (d_951_v32_)[default__.safeIndex(len(d_981_v56_), len(d_951_v32_))]})
                d_982_v57_ = (d_982_v57_).set((self).f34, (d_951_v32_)[default__.safeIndex(self.f29, len(d_951_v32_))])
            elif True:
                d_983_v58_: _dafny.Map
                d_983_v58_ = _dafny.Map({(self).f34: d_951_v32_})
                d_951_v32_ = (d_951_v32_) + ((((d_983_v58_)[len(d_949_v30_)] if (len(d_949_v30_)) in (d_983_v58_) else d_951_v32_)) + (d_951_v32_))
                d_984_v59_: _dafny.Seq
                d_984_v59_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f34: False}), _dafny.Map({self.f29: (self).f28})])
                d_985_v60_: _dafny.Seq
                d_985_v60_ = _dafny.SeqWithoutIsStrInference([d_951_v32_])
                d_986_v61_: _dafny.Map
                d_986_v61_ = _dafny.Map({(d_984_v59_)[default__.safeIndex((self).f34, len(d_984_v59_))]: d_985_v60_})
                d_986_v61_ = (d_986_v61_).set((d_909_v0_) | (d_909_v0_), _dafny.SeqWithoutIsStrInference([d_951_v32_ for d_987_i5_ in range(default__.abs(-231))]))
                d_911_v2_ = d_911_v2_
                (globalState).f13 = self.f29
                (globalState).f21 = (d_952_v33_)[default__.safeIndex(self.f29, len(d_952_v33_))]
            d_988_v62_: _dafny.Map
            d_988_v62_ = _dafny.Map({(self).f34: self.f29})
            d_989_v63_: C0
            nw155_ = C0()
            nw155_.ctor__(d_951_v32_, ((self).f34) * (len(d_988_v62_)), (self).f27, p0)
            d_989_v63_ = nw155_
        elif True:
            d_990___mcc_h3_ = source13_.cf32
            d_991_cf32_ = d_990___mcc_h3_
            d_992_v64_: C1
            nw156_ = C1()
            nw156_.ctor__(p0, (self).f27, (self).f28)
            d_992_v64_ = nw156_
            d_993_v65_: _dafny.Seq
            d_993_v65_ = _dafny.SeqWithoutIsStrInference([d_992_v64_])
            (self).f29 = (default__.safeModuloInt(self.f29, self.f29) if not((d_992_v64_) in (d_993_v65_)) else len(((d_950_v31_)[default__.safeIndex(427, (d_950_v31_).length(0))]) + (d_949_v30_)))
            (globalState).f15 = (d_992_v64_).f39
            (globalState).f15 = p0
            d_994_v66_: C0
            nw157_ = C0()
            nw157_.ctor__(d_951_v32_, (self).f34, (self).f27, (self).f28)
            d_994_v66_ = nw157_
            d_995_v67_: _dafny.Seq
            d_995_v67_ = _dafny.SeqWithoutIsStrInference([d_994_v66_])
            d_996_v68_: _dafny.Set
            d_996_v68_ = _dafny.Set({(d_995_v67_)[default__.safeIndex((self).f34, len(d_995_v67_))], d_994_v66_})
            (globalState).f15 = (_dafny.Set({d_994_v66_})).issubset(d_996_v68_)
        d_997_v69_: _dafny.Set
        d_997_v69_ = _dafny.Set({self.f29})
        r0 = len(d_997_v69_)
        d_998_v70_: str
        d_998_v70_ = _dafny.CodePoint('r')
        d_999_v71_: D1
        d_999_v71_ = D1_DC7(d_998_v70_, d_951_v32_)
        r1 = ((d_951_v32_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyjwr")))).set(default__.safeIndex(len(_dafny.Map({d_999_v71_: (self).f27})), len((d_951_v32_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyjwr"))))), d_998_v70_)
        r2 = (self).f28
        return r0, r1, r2

    @property
    def f34(self):
        return self._f34
    @property
    def f35(self):
        return self._f35

class C4:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self):
        pass
        pass

    def fm12(self, p0, p1, p2, globalState):
        source15_ = D1_DC6((_dafny.MultiSet([187, len(_dafny.SeqWithoutIsStrInference([-962 for d_1000_i0_ in range(default__.abs(55))]))])).cardinality)
        if source15_.is_DC6:
            d_1001___mcc_h0_ = source15_.cf12
            d_1002_cf12_ = d_1001___mcc_h0_
            return (_dafny.MultiSet([_dafny.Map({False: not(False)})])).intersection(_dafny.MultiSet([_dafny.Map({True: False}), _dafny.Map({True: True}), _dafny.Map({False: not(False)}), _dafny.Map({True: False})]))
        elif source15_.is_DC7:
            d_1003___mcc_h1_ = source15_.cf13
            d_1004___mcc_h2_ = source15_.cf14
            d_1005_cf14_ = d_1004___mcc_h2_
            d_1006_cf13_ = d_1003___mcc_h1_
            return _dafny.MultiSet([_dafny.Map({False: True}), _dafny.Map({False: not(not(True))})])
        elif True:
            d_1007___mcc_h3_ = source15_.cf11
            d_1008_cf11_ = d_1007___mcc_h3_
            return _dafny.MultiSet([_dafny.Map({True: False}), _dafny.Map({False: True})])

    def fm13(self, p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([True])

    def m8(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Set = _dafny.Set({})
        r3: _dafny.Seq = _dafny.Seq({})
        d_1009_v0_: _dafny.Seq
        d_1009_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxd"))
        d_1010_v1_: bool
        d_1010_v1_ = True
        d_1011_v2_: _dafny.Array
        nw158_ = _dafny.Array(None, 9)
        nw158_[int(0)] = d_1010_v1_
        nw158_[int(1)] = d_1010_v1_
        nw158_[int(2)] = d_1010_v1_
        nw158_[int(3)] = d_1010_v1_
        nw158_[int(4)] = d_1010_v1_
        nw158_[int(5)] = d_1010_v1_
        nw158_[int(6)] = d_1010_v1_
        nw158_[int(7)] = d_1010_v1_
        nw158_[int(8)] = d_1010_v1_
        d_1011_v2_ = nw158_
        d_1012_v3_: T1
        nw159_ = C0()
        nw159_.ctor__(d_1009_v0_, p1, d_1011_v2_, d_1010_v1_)
        d_1012_v3_ = nw159_
        d_1013_v4_: _dafny.Map
        d_1013_v4_ = _dafny.Map({d_1012_v3_: p1})
        d_1014_v5_: _dafny.Map
        d_1014_v5_ = _dafny.Map({d_1013_v4_: (d_1012_v3_).f28})
        d_1015_v6_: _dafny.Map
        d_1015_v6_ = _dafny.Map({d_1012_v3_.f29: default__.fm36(globalState)})
        d_1016_v7_: str
        d_1016_v7_ = _dafny.CodePoint('e')
        d_1017_v8_: D0
        d_1017_v8_ = D0_DC0(d_1014_v5_, False, len(d_1015_v6_), d_1016_v7_, not(d_1010_v1_))
        r0 = (d_1017_v8_).cf2
        d_1018_v9_: C1
        nw160_ = C1()
        nw160_.ctor__((d_1010_v1_ if (d_1012_v3_).f28 else d_1010_v1_), (d_1012_v3_).f27, d_1010_v1_)
        d_1018_v9_ = nw160_
        hi3_ = (d_1012_v3_.f29 if (d_1018_v9_).f39 else (p0)[default__.safeIndex(p1, len(p0))])
        for d_1019_i0_ in range(len(_dafny.SeqWithoutIsStrInference([len(d_1009_v0_), p1, d_1012_v3_.f29])), hi3_):
            d_1020_v10_: D0
            d_1020_v10_ = D0_DC3(default__.fm36(globalState), (d_1018_v9_).f39, (d_1012_v3_).f28)
            index154_ = default__.safeIndex(799, ((d_1012_v3_).f27).length(0))
            ((d_1012_v3_).f27)[index154_] = (d_1020_v10_).cf8
            index155_ = default__.safeIndex(799, ((d_1012_v3_).f27).length(0))
            ((d_1012_v3_).f27)[index155_] = not(d_1010_v1_)
            d_1021_v11_: _dafny.Set
            d_1021_v11_ = _dafny.Set({(d_1018_v9_).f39})
            d_1022_v12_: _dafny.MultiSet
            d_1022_v12_ = _dafny.MultiSet([(_dafny.Set({d_1010_v1_, ((d_1012_v3_).f27)[default__.safeIndex(799, ((d_1012_v3_).f27).length(0))]}) if ((d_1012_v3_).f27)[default__.safeIndex(799, ((d_1012_v3_).f27).length(0))] else d_1021_v11_)])
            d_1022_v12_ = (d_1022_v12_ if (D0_DC3(d_1012_v3_.f29, True, False)).cf9 else d_1022_v12_)
            (globalState).f25 = (d_1009_v0_)[default__.safeIndex(default__.fm36(globalState), len(d_1009_v0_))]
            index156_ = default__.safeIndex(799, ((d_1012_v3_).f27).length(0))
            ((d_1012_v3_).f27)[index156_] = ((d_1021_v11_).intersection(d_1021_v11_)).ispropersubset(d_1021_v11_)
        d_1023_i1_: int
        d_1023_i1_ = 0
        with _dafny.label("3"):
            while (d_1012_v3_).f28:
                with _dafny.c_label("3"):
                    if (d_1023_i1_) >= (100):
                        raise _dafny.Break("3")
                    d_1023_i1_ = (d_1023_i1_) + (1)
                    d_1024_v13_: _dafny.MultiSet
                    d_1024_v13_ = _dafny.MultiSet([not(not((d_1012_v3_).f28)), (d_1018_v9_).f39])
                    d_1025_v14_: _dafny.Seq
                    d_1025_v14_ = _dafny.SeqWithoutIsStrInference([d_1024_v13_])
                    d_1026_v15_: D9
                    d_1026_v15_ = D9_DC23((d_1018_v9_).f39)
                    d_1027_v16_: _dafny.Map
                    d_1027_v16_ = _dafny.Map({(d_1018_v9_).f39: d_1026_v15_})
                    d_1028_v17_: _dafny.Map
                    d_1028_v17_ = _dafny.Map({d_1012_v3_.f29: d_1027_v16_})
                    d_1029_v18_: _dafny.Map
                    d_1029_v18_ = _dafny.Map({((d_1028_v17_)[(0) - (p1)] if ((0) - (p1)) in (d_1028_v17_) else d_1027_v16_): d_1010_v1_})
                    d_1030_v20_: _dafny.Array
                    def lambda38_(d_1031_v3_, d_1032_p1_, d_1033_v1_):
                        def lambda39_(d_1034_i2_):
                            def iife88_():
                                coll50_ = _dafny.Set()
                                compr_50_: int
                                for compr_50_ in _dafny.IntegerRange(-97, 757):
                                    d_1035_v19_: int = compr_50_
                                    if ((-97) <= (d_1035_v19_)) and ((d_1035_v19_) < (757)):
                                        coll50_ = coll50_.union(_dafny.Set([(d_1035_v19_) + ((_dafny.MultiSet([d_1033_v1_])).cardinality)]))
                                return _dafny.Set(coll50_)
                            return (_dafny.Set({d_1031_v3_.f29, d_1032_p1_, d_1031_v3_.f29})).intersection(iife88_()
                            )

                        return lambda39_

                    init23_ = lambda38_(d_1012_v3_, p1, d_1010_v1_)
                    nw161_ = _dafny.Array(None, 16)
                    for i0_23_ in range(nw161_.length(0)):
                        nw161_[i0_23_] = init23_(i0_23_)
                    d_1030_v20_ = nw161_
                    d_1036_v21_: _dafny.Set
                    d_1036_v21_ = _dafny.Set({len(d_1009_v0_)})
                    index157_ = default__.safeIndex(349, (d_1030_v20_).length(0))
                    (d_1030_v20_)[index157_] = (d_1036_v21_).intersection(d_1036_v21_)
                    d_1037_v22_: D0
                    d_1037_v22_ = D0_DC3(default__.fm36(globalState), (d_1018_v9_).f39, True)
                    index158_ = default__.safeIndex(349, (d_1030_v20_).length(0))
                    rhs171_ = d_1012_v3_.f29
                    rhs172_ = d_1025_v14_
                    rhs173_ = _dafny.Map({d_1027_v16_: (d_1012_v3_.f29) != (d_1012_v3_.f29)})
                    rhs174_ = d_1036_v21_
                    rhs175_ = (d_1037_v22_).cf8
                    lhs146_ = globalState
                    lhs147_ = d_1030_v20_
                    lhs148_ = default__.safeIndex(349, (d_1030_v20_).length(0))
                    lhs149_ = globalState
                    lhs146_.f13 = rhs171_
                    d_1025_v14_ = rhs172_
                    d_1029_v18_ = rhs173_
                    lhs147_[lhs148_] = rhs174_
                    lhs149_.f17 = rhs175_
                    d_1038_v23_: _dafny.Array
                    nw162_ = _dafny.Array(D8.default()(), 10)
                    d_1038_v23_ = nw162_
                    d_1039_v24_: _dafny.MultiSet
                    d_1039_v24_ = _dafny.MultiSet([d_1038_v23_])
                    d_1039_v24_ = d_1039_v24_
                    d_1037_v22_ = D0_DC3(d_1012_v3_.f29, (d_1018_v9_).f39, (d_1012_v3_).f28)
                    (globalState).f4 = (p1) + ((p0)[default__.safeIndex(p1, len(p0))])
                    pass
            pass
        (globalState).f5 = d_1012_v3_.f29
        d_1040_v25_: _dafny.Array
        nw163_ = _dafny.Array(_dafny.Array(None, 0), 6)
        d_1040_v25_ = nw163_
        d_1041_v26_: _dafny.Map
        d_1041_v26_ = _dafny.Map({417: d_1016_v7_})
        d_1042_v27_: _dafny.Map
        d_1042_v27_ = _dafny.Map({(d_1012_v3_).f28: (d_1018_v9_).f39})
        d_1043_v28_: _dafny.Array
        nw164_ = _dafny.Array(None, 16)
        nw164_[int(0)] = _dafny.Map({len(p0): d_1016_v7_})
        nw164_[int(1)] = d_1041_v26_
        nw164_[int(2)] = d_1041_v26_
        nw164_[int(3)] = d_1041_v26_
        nw164_[int(4)] = d_1041_v26_
        nw164_[int(5)] = d_1041_v26_
        nw164_[int(6)] = d_1041_v26_
        nw164_[int(7)] = d_1041_v26_
        nw164_[int(8)] = d_1041_v26_
        nw164_[int(9)] = d_1041_v26_
        nw164_[int(10)] = d_1041_v26_
        nw164_[int(11)] = d_1041_v26_
        nw164_[int(12)] = d_1041_v26_
        nw164_[int(13)] = d_1041_v26_
        nw164_[int(14)] = d_1041_v26_
        nw164_[int(15)] = _dafny.Map({len(d_1042_v27_): d_1016_v7_})
        d_1043_v28_ = nw164_
        index159_ = default__.safeIndex(462, (d_1040_v25_).length(0))
        (d_1040_v25_)[index159_] = d_1043_v28_
        index160_ = default__.safeIndex(133, ((d_1012_v3_).f27).length(0))
        ((d_1012_v3_).f27)[index160_] = d_1010_v1_
        d_1044_v29_: _dafny.Map
        d_1044_v29_ = _dafny.Map({d_1010_v1_: d_1009_v0_})
        index161_ = default__.safeIndex(462, (d_1040_v25_).length(0))
        index162_ = default__.safeIndex(133, ((d_1012_v3_).f27).length(0))
        rhs176_ = d_1043_v28_
        rhs177_ = (p0)[default__.safeIndex(default__.safeDivisionInt(p1, len(((d_1044_v29_)[(d_1012_v3_).f28] if ((d_1012_v3_).f28) in (d_1044_v29_) else d_1009_v0_))), len(p0))]
        rhs178_ = (d_1012_v3_).f28
        lhs150_ = d_1040_v25_
        lhs151_ = default__.safeIndex(462, (d_1040_v25_).length(0))
        lhs152_ = globalState
        lhs153_ = (d_1012_v3_).f27
        lhs154_ = default__.safeIndex(133, ((d_1012_v3_).f27).length(0))
        lhs150_[lhs151_] = rhs176_
        lhs152_.f13 = rhs177_
        lhs153_[lhs154_] = rhs178_
        r0 = p1
        r1 = (d_1018_v9_).f39
        def iife89_():
            coll51_ = _dafny.Set()
            compr_51_: int
            for compr_51_ in ((_dafny.Map({-443: (d_1018_v9_).f39})).set(len(d_1009_v0_), ((d_1012_v3_).f27)[default__.safeIndex(133, ((d_1012_v3_).f27).length(0))])).keys.Elements:
                d_1045_v30_: int = compr_51_
                if (d_1045_v30_) in ((_dafny.Map({-443: (d_1018_v9_).f39})).set(len(d_1009_v0_), ((d_1012_v3_).f27)[default__.safeIndex(133, ((d_1012_v3_).f27).length(0))])):
                    coll51_ = coll51_.union(_dafny.Set([(d_1045_v30_) * ((0) - (((0) - (len(_dafny.Set({True, False, True, not(False), True})))) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r')])])).cardinality)))]))
            return _dafny.Set(coll51_)
        r2 = iife89_()
        
        r3 = ((p0).set(default__.safeIndex(d_1012_v3_.f29, len(p0)), p1)) + (p0)
        return r0, r1, r2, r3

    def m9(self, p0, p1, p2, p3, globalState):
        d_1046_v0_: _dafny.Seq
        d_1046_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "islesyilu"))
        d_1047_v1_: D8
        d_1047_v1_ = D8_DC18(p3)
        d_1048_v2_: _dafny.Array
        def lambda40_(d_1049_p2_):
            def lambda41_(d_1050_i0_):
                return d_1049_p2_

            return lambda41_

        init24_ = lambda40_(p2)
        nw165_ = _dafny.Array(None, 6)
        for i0_24_ in range(nw165_.length(0)):
            nw165_[i0_24_] = init24_(i0_24_)
        d_1048_v2_ = nw165_
        d_1051_v3_: C0
        nw166_ = C0()
        nw166_.ctor__(d_1046_v0_, (d_1047_v1_).cf31, d_1048_v2_, False)
        d_1051_v3_ = nw166_
        d_1052_v5_: _dafny.Set
        d_1052_v5_ = _dafny.Set({p0, p3})
        def iife90_():
            coll52_ = _dafny.Set()
            compr_52_: int
            for compr_52_ in _dafny.IntegerRange(623, 207):
                d_1053_v4_: int = compr_52_
                if ((623) <= (d_1053_v4_)) and ((d_1053_v4_) < (207)):
                    coll52_ = coll52_.union(_dafny.Set([default__.safeDivisionInt(d_1053_v4_, p3)]))
            return _dafny.Set(coll52_)
        if ((d_1052_v5_).ispropersubset(iife90_()
        ) if ((d_1051_v3_).f36) <= (d_1046_v0_) else not(p2)):
            d_1054_v6_: str
            d_1054_v6_ = _dafny.CodePoint('r')
            d_1055_v7_: D1
            d_1055_v7_ = D1_DC7(d_1054_v6_, (d_1051_v3_).f36)
            source16_ = (d_1055_v7_ if p2 else d_1055_v7_)
            if source16_.is_DC6:
                d_1056___mcc_h0_ = source16_.cf12
                d_1057_cf12_ = d_1056___mcc_h0_
                d_1058_v8_: C1
                nw167_ = C1()
                nw167_.ctor__(p2, d_1048_v2_, p2)
                d_1058_v8_ = nw167_
                d_1059_v9_: D1
                d_1059_v9_ = D1_DC5(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1057_cf12_, d_1057_cf12_, p3, len((d_1051_v3_).f36), p0]) for d_1060_i1_ in range(default__.abs(946))]))
                d_1061_v10_: _dafny.Map
                d_1061_v10_ = _dafny.Map({p2: (0) - (p3)})
                d_1062_v11_: _dafny.MultiSet
                d_1062_v11_ = _dafny.MultiSet([p3])
                d_1059_v9_ = default__.fm47(d_1061_v10_, d_1062_v11_, p3, globalState)
                (globalState).f15 = (d_1058_v8_).f39
                (globalState).f25 = d_1054_v6_
            elif source16_.is_DC7:
                d_1063___mcc_h1_ = source16_.cf13
                d_1064___mcc_h2_ = source16_.cf14
                d_1065_cf14_ = d_1064___mcc_h2_
                d_1066_cf13_ = d_1063___mcc_h1_
                d_1067_v12_: C1
                nw168_ = C1()
                nw168_.ctor__(False, d_1048_v2_, p2)
                d_1067_v12_ = nw168_
                d_1068_v13_: _dafny.Set
                d_1068_v13_ = _dafny.Set({d_1054_v6_, default__.fm42(p0, globalState)})
                d_1069_v14_: C0
                nw169_ = C0()
                nw169_.ctor__(d_1046_v0_, (len(_dafny.SeqWithoutIsStrInference([p2]))) - (len(d_1068_v13_)), d_1048_v2_, p2)
                d_1069_v14_ = nw169_
                d_1070_v15_: _dafny.Seq
                d_1070_v15_ = _dafny.SeqWithoutIsStrInference([p3])
                index163_ = default__.safeIndex(182, (p1).length(0))
                (p1)[index163_] = (d_1070_v15_)[default__.safeIndex(p0, len(d_1070_v15_))]
                d_1071_v16_: _dafny.Map
                d_1071_v16_ = _dafny.Map({415: p0})
                d_1072_v17_: _dafny.Seq
                d_1072_v17_ = _dafny.SeqWithoutIsStrInference([(d_1071_v16_) | (default__.fm46(not(p2), globalState)), d_1071_v16_])
                index164_ = default__.safeIndex(701, (p1).length(0))
                (p1)[index164_] = p3
                index165_ = default__.safeIndex(182, (p1).length(0))
                index166_ = default__.safeIndex(701, (p1).length(0))
                rhs179_ = default__.safeModuloInt(p0, p0)
                rhs180_ = (d_1072_v17_) + ((d_1072_v17_) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({p0: p3}), _dafny.Map({p0: -711})])))
                rhs181_ = p3
                lhs155_ = p1
                lhs156_ = default__.safeIndex(182, (p1).length(0))
                lhs157_ = p1
                lhs158_ = default__.safeIndex(701, (p1).length(0))
                lhs155_[lhs156_] = rhs179_
                d_1072_v17_ = rhs180_
                lhs157_[lhs158_] = rhs181_
                d_1066_cf13_ = d_1054_v6_
            elif True:
                d_1073___mcc_h3_ = source16_.cf11
                d_1074_cf11_ = d_1073___mcc_h3_
                d_1075_v18_: _dafny.MultiSet
                d_1075_v18_ = _dafny.MultiSet([p2, p2, p2, p2, (p3) < (p3)])
                d_1075_v18_ = ((d_1075_v18_) - (d_1075_v18_)) - (d_1075_v18_)
                (globalState).f15 = (d_1075_v18_).issubset((d_1075_v18_).intersection(d_1075_v18_))
                d_1046_v0_ = (((d_1051_v3_).f36) + ((d_1051_v3_).f36)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppri"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "clfofo"))))
                d_1076_v19_: _dafny.Set
                d_1076_v19_ = _dafny.Set({p2, p2})
                d_1077_v20_: _dafny.Map
                d_1077_v20_ = _dafny.Map({(d_1076_v19_).intersection(d_1076_v19_): False})
                d_1078_v21_: _dafny.Seq
                d_1078_v21_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1077_v20_ = (d_1077_v20_).set((d_1076_v19_) | (_dafny.Set({p2})), not((d_1078_v21_)[default__.safeIndex(p0, len(d_1078_v21_))]))
            d_1079_v22_: _dafny.Seq
            d_1079_v22_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1080_v23_: _dafny.Set
            d_1080_v23_ = _dafny.Set({p2})
            d_1081_v24_: _dafny.MultiSet
            d_1081_v24_ = _dafny.MultiSet([(d_1079_v22_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1080_v23_])), len(d_1079_v22_))]])
            d_1082_v25_: _dafny.Map
            d_1082_v25_ = _dafny.Map({((d_1081_v24_).set(p0, default__.abs(p3))).set(p3, default__.abs(p0)): d_1054_v6_})
            d_1083_v27_: _dafny.Map
            d_1083_v27_ = _dafny.Map({d_1081_v24_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wydr")))})
            def iife91_():
                coll53_ = _dafny.Map()
                compr_53_: _dafny.MultiSet
                for compr_53_ in (d_1083_v27_).keys.Elements:
                    d_1084_v26_: _dafny.MultiSet = compr_53_
                    if (d_1084_v26_) in (d_1083_v27_):
                        coll53_[d_1084_v26_] = d_1054_v6_
                return _dafny.Map(coll53_)
            d_1082_v25_ = (d_1082_v25_) | (iife91_()
            )
            d_1085_v28_: _dafny.Seq
            d_1085_v28_ = _dafny.SeqWithoutIsStrInference([p2])
            d_1086_v29_: D0
            d_1086_v29_ = D0_DC3(len(d_1085_v28_), not(p2), p2)
            d_1087_v30_: _dafny.Map
            d_1087_v30_ = _dafny.Map({d_1086_v29_: p0})
            d_1088_v31_: _dafny.Array
            d_1088_v31_ = d_1048_v2_
            d_1089_v32_: C3
            nw170_ = C3()
            nw170_.ctor__(p0, (d_1087_v30_) | (d_1087_v30_), (0) - ((p0) + (p3)), (d_1048_v2_), p2)
            d_1089_v32_ = nw170_
            index167_ = default__.safeIndex(235, (p1).length(0))
            (p1)[index167_] = (len(d_1080_v23_)) - (p3)
            index168_ = default__.safeIndex(235, (p1).length(0))
            (p1)[index168_] = (d_1089_v32_).f34
            d_1090_v33_: D1
            d_1090_v33_ = D1_DC6(p0)
            index169_ = default__.safeIndex(235, (p1).length(0))
            (p1)[index169_] = (len((_dafny.SeqWithoutIsStrInference([(d_1090_v33_).cf12])) + (d_1079_v22_))) - ((d_1089_v32_).f34)
        elif True:
            d_1091_v34_: D8
            d_1091_v34_ = D8_DC18((0) - (p3))
            d_1092_v35_: D8
            d_1092_v35_ = D8_DC20(d_1091_v34_)
            source17_ = d_1092_v35_
            if source17_.is_DC18:
                d_1093___mcc_h4_ = source17_.cf31
                d_1094_cf31_ = d_1093___mcc_h4_
                d_1095_v36_: _dafny.Seq
                d_1095_v36_ = _dafny.SeqWithoutIsStrInference([p2, not(not(p2)), not(p2), p2, p2])
                d_1096_v37_: _dafny.Map
                d_1096_v37_ = _dafny.Map({len((d_1051_v3_).f36): p2})
                (globalState).f6 = ((d_1095_v36_)[default__.safeIndex((0) - (p3), len(d_1095_v36_))] if p2 else not(((d_1096_v37_)[635] if (635) in (d_1096_v37_) else p2)))
                d_1097_v38_: _dafny.Seq
                d_1097_v38_ = _dafny.SeqWithoutIsStrInference([d_1052_v5_, d_1052_v5_, d_1052_v5_])
                d_1098_v39_: _dafny.MultiSet
                d_1098_v39_ = _dafny.MultiSet([p0, p0, p3])
                (globalState).f17 = ((d_1097_v38_)[default__.safeIndex((d_1098_v39_).cardinality, len(d_1097_v38_))]).issubset(_dafny.Set({d_1094_cf31_}))
                d_1099_v40_: _dafny.Map
                d_1100_v41_: int
                out21_: _dafny.Map
                out22_: int
                out21_, out22_ = default__.m0(globalState)
                d_1099_v40_ = out21_
                d_1100_v41_ = out22_
                d_1101_v42_: _dafny.Seq
                d_1101_v42_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1102_v43_: _dafny.Seq
                d_1102_v43_ = _dafny.SeqWithoutIsStrInference([d_1101_v42_, d_1101_v42_, d_1101_v42_, d_1101_v42_, d_1101_v42_])
                d_1103_v44_: _dafny.Seq
                d_1103_v44_ = (d_1102_v43_)[default__.safeIndex(p3, len(d_1102_v43_))]
                d_1104_v45_: str
                d_1104_v45_ = _dafny.CodePoint('p')
                d_1105_v46_: _dafny.Set
                d_1105_v46_ = _dafny.Set({default__.fm26(p2, d_1094_cf31_, d_1104_v45_, globalState)})
                d_1106_v47_: D0
                d_1106_v47_ = D0_DC3((d_1051_v3_).fm15(globalState), p2, p2)
                rhs182_ = p0
                rhs183_ = p2
                rhs184_ = ((d_1105_v46_) | (d_1105_v46_)).ispropersubset(d_1105_v46_)
                rhs185_ = d_1103_v44_
                rhs186_ = (d_1106_v47_).cf9
                lhs159_ = globalState
                lhs160_ = globalState
                lhs161_ = globalState
                lhs162_ = globalState
                lhs159_.f5 = rhs182_
                lhs160_.f6 = rhs183_
                lhs161_.f6 = rhs184_
                d_1103_v44_ = rhs185_
                lhs162_.f15 = rhs186_
            elif source17_.is_DC19:
                d_1107_v48_: str
                d_1107_v48_ = _dafny.CodePoint('r')
                d_1046_v0_ = ((d_1051_v3_).f36).set(default__.safeIndex(default__.fm36(globalState), len((d_1051_v3_).f36)), d_1107_v48_)
                d_1108_v50_: _dafny.Seq
                d_1108_v50_ = _dafny.SeqWithoutIsStrInference([d_1052_v5_])
                def iife92_():
                    coll54_ = _dafny.Map()
                    compr_54_: int
                    for compr_54_ in ((d_1108_v50_)[default__.safeIndex(len(d_1052_v5_), len(d_1108_v50_))]).Elements:
                        d_1109_v49_: int = compr_54_
                        if (d_1109_v49_) in ((d_1108_v50_)[default__.safeIndex(len(d_1052_v5_), len(d_1108_v50_))]):
                            coll54_[(d_1109_v49_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))))] = p0
                    return _dafny.Map(coll54_)
                def iife93_():
                    coll55_ = _dafny.Map()
                    compr_55_: int
                    for compr_55_ in _dafny.IntegerRange(615, 805):
                        d_1110_v51_: int = compr_55_
                        if ((615) <= (d_1110_v51_)) and ((d_1110_v51_) < (805)):
                            coll55_[default__.safeDivisionInt(d_1110_v51_, (0) - (len(_dafny.SeqWithoutIsStrInference([d_1107_v48_ for d_1111_i2_ in range(default__.abs(37))]))))] = -578
                    return _dafny.Map(coll55_)
                def iife94_():
                    coll56_ = _dafny.Map()
                    compr_56_: int
                    for compr_56_ in _dafny.IntegerRange(787, -73):
                        d_1112_v52_: int = compr_56_
                        if ((787) <= (d_1112_v52_)) and ((d_1112_v52_) < (-73)):
                            coll56_[default__.safeModuloInt(d_1112_v52_, (0) - (p3))] = p3
                    return _dafny.Map(coll56_)
                rhs187_ = default__.fm32(p0, p2, globalState)
                rhs188_ = p2
                rhs189_ = not((iife92_()
                ) == ((iife93_()
                ) | (iife94_()
                )))
                lhs163_ = globalState
                lhs164_ = globalState
                d_1046_v0_ = rhs187_
                lhs163_.f17 = rhs188_
                lhs164_.f15 = rhs189_
                d_1113_v53_: _dafny.MultiSet
                d_1113_v53_ = _dafny.MultiSet([p2, p2])
                d_1114_v54_: _dafny.Map
                d_1114_v54_ = _dafny.Map({p0: (d_1113_v53_).cardinality})
                d_1115_v56_: C1
                nw171_ = C1()
                def iife95_():
                    coll57_ = _dafny.Set()
                    compr_57_: int
                    for compr_57_ in (d_1114_v54_).keys.Elements:
                        d_1116_v55_: int = compr_57_
                        if (d_1116_v55_) in (d_1114_v54_):
                            coll57_ = coll57_.union(_dafny.Set([(d_1116_v55_) - (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True, False])).cardinality for d_1117_i3_ in range(default__.abs(665))])))]))
                    return _dafny.Set(coll57_)
                nw171_.ctor__((d_1052_v5_).ispropersubset(iife95_()
                ), d_1048_v2_, p2)
                d_1115_v56_ = nw171_
                rhs190_ = d_1115_v56_
                rhs191_ = p2
                lhs165_ = globalState
                d_1115_v56_ = rhs190_
                lhs165_.f6 = rhs191_
                (globalState).f5 = default__.safeModuloInt(default__.fm36(globalState), default__.safeDivisionInt((d_1051_v3_).fm15(globalState), p0))
            elif source17_.is_DC17:
                d_1118___mcc_h5_ = source17_.cf30
                d_1119_cf30_ = d_1118___mcc_h5_
                d_1120_v57_: _dafny.MultiSet
                d_1120_v57_ = _dafny.MultiSet([D7_DC15(_dafny.Map({len(d_1052_v5_): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgx"))}))])
                d_1121_v58_: D8
                d_1121_v58_ = D8_DC17(d_1119_cf30_)
                (globalState).f4 = ((d_1120_v57_)[default__.fm48(p3, p2, p0, d_1121_v58_, globalState)] if (default__.fm48(p3, p2, p0, d_1121_v58_, globalState)) in (d_1120_v57_) else p0)
                (globalState).f4 = default__.safeDivisionInt(default__.safeModuloInt(p3, p3), p0)
                (globalState).f5 = default__.safeDivisionInt((default__.fm36(globalState)) - (p3), default__.safeModuloInt(625, 647))
                d_1122_v59_: _dafny.Array
                def lambda42_(d_1123_i4_):
                    return _dafny.CodePoint('r')

                init25_ = lambda42_
                nw172_ = _dafny.Array(None, 23)
                for i0_25_ in range(nw172_.length(0)):
                    nw172_[i0_25_] = init25_(i0_25_)
                d_1122_v59_ = nw172_
                d_1124_v60_: str
                d_1124_v60_ = _dafny.CodePoint('u')
                index170_ = default__.safeIndex(456, (d_1122_v59_).length(0))
                (d_1122_v59_)[index170_] = d_1124_v60_
                d_1125_v61_: _dafny.Seq
                d_1125_v61_ = _dafny.SeqWithoutIsStrInference([145])
                d_1126_v62_: _dafny.MultiSet
                d_1126_v62_ = _dafny.MultiSet([((d_1125_v61_)[default__.safeIndex(p3, len(d_1125_v61_))]) < (len((d_1051_v3_).f36))])
                d_1127_v63_: _dafny.Seq
                d_1127_v63_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1128_v64_: _dafny.Map
                d_1128_v64_ = _dafny.Map({p0: p0})
                d_1129_v65_: _dafny.Seq
                d_1129_v65_ = _dafny.SeqWithoutIsStrInference([d_1125_v61_])
                d_1130_v66_: D0
                d_1130_v66_ = D0_DC3(len(d_1125_v61_), p2, p2)
                index171_ = default__.safeIndex(456, (d_1122_v59_).length(0))
                rhs192_ = default__.safeModuloInt(len((d_1127_v63_) + (d_1127_v63_)), p0)
                rhs193_ = ((d_1128_v64_)[default__.fm36(globalState)] if (default__.fm36(globalState)) in (d_1128_v64_) else len(((d_1129_v65_)[default__.safeIndex(p0, len(d_1129_v65_))]) + (_dafny.SeqWithoutIsStrInference([p0]))))
                rhs194_ = (d_1124_v60_ if p2 else d_1124_v60_)
                rhs195_ = (d_1126_v62_) | (_dafny.MultiSet([p2]))
                rhs196_ = default__.safeModuloInt((len(default__.fm0(_dafny.Map({p3: d_1052_v5_}), p2, (d_1130_v66_).cf8, globalState))) + (p3), p0)
                lhs166_ = globalState
                lhs167_ = globalState
                lhs168_ = d_1122_v59_
                lhs169_ = default__.safeIndex(456, (d_1122_v59_).length(0))
                lhs170_ = globalState
                lhs166_.f13 = rhs192_
                lhs167_.f21 = rhs193_
                lhs168_[lhs169_] = rhs194_
                d_1126_v62_ = rhs195_
                lhs170_.f5 = rhs196_
            elif True:
                d_1131___mcc_h6_ = source17_.cf32
                d_1132_cf32_ = d_1131___mcc_h6_
                d_1048_v2_ = d_1048_v2_
                (globalState).f6 = not(p2)
                d_1133_v67_: _dafny.Seq
                d_1133_v67_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1134_v68_: _dafny.Set
                d_1134_v68_ = _dafny.Set({(d_1133_v67_)[default__.safeIndex(p0, len(d_1133_v67_))], p1})
                d_1135_v69_: D0
                d_1135_v69_ = D0_DC3(p0, not(p2), p2)
                d_1136_v70_: _dafny.Map
                d_1136_v70_ = _dafny.Map({d_1135_v69_: p3})
                d_1137_v71_: T1
                nw173_ = C3()
                nw173_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfeldy"))), d_1136_v70_, p0, d_1048_v2_, p2)
                d_1137_v71_ = nw173_
                d_1138_v72_: _dafny.Map
                d_1138_v72_ = _dafny.Map({d_1137_v71_: d_1137_v71_.f29})
                d_1139_v73_: _dafny.Map
                d_1139_v73_ = _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0 for d_1140_i5_ in range(default__.abs(324))]))).cardinality: p2})
                d_1141_v74_: _dafny.Map
                d_1141_v74_ = _dafny.Map({d_1138_v72_: ((d_1139_v73_)[(0) - (d_1137_v71_.f29)] if ((0) - (d_1137_v71_.f29)) in (d_1139_v73_) else (d_1137_v71_).f28)})
                d_1142_v75_: str
                d_1142_v75_ = _dafny.CodePoint('v')
                d_1143_v76_: D0
                d_1143_v76_ = D0_DC0(d_1141_v74_, p2, d_1137_v71_.f29, d_1142_v75_, not(p2))
                index172_ = default__.safeIndex(582, (d_1048_v2_).length(0))
                (d_1048_v2_)[index172_] = not((d_1143_v76_).cf1)
                index173_ = default__.safeIndex(855, (d_1048_v2_).length(0))
                (d_1048_v2_)[index173_] = (d_1137_v71_).f28
                index174_ = default__.safeIndex(523, ((d_1137_v71_).f27).length(0))
                ((d_1137_v71_).f27)[index174_] = not(p2)
                d_1144_v77_: _dafny.Map
                d_1144_v77_ = _dafny.Map({p0: p3})
                index175_ = default__.safeIndex(582, (d_1048_v2_).length(0))
                index176_ = default__.safeIndex(855, (d_1048_v2_).length(0))
                index177_ = default__.safeIndex(523, ((d_1137_v71_).f27).length(0))
                rhs197_ = 720
                rhs198_ = (d_1134_v68_ if False else d_1134_v68_)
                rhs199_ = (d_1137_v71_).f28
                rhs200_ = default__.fm40(globalState)
                rhs201_ = (len((d_1144_v77_) | (d_1144_v77_))) >= (default__.safeModuloInt(d_1137_v71_.f29, p3))
                lhs171_ = globalState
                lhs172_ = d_1048_v2_
                lhs173_ = default__.safeIndex(582, (d_1048_v2_).length(0))
                lhs174_ = d_1048_v2_
                lhs175_ = default__.safeIndex(855, (d_1048_v2_).length(0))
                lhs176_ = (d_1137_v71_).f27
                lhs177_ = default__.safeIndex(523, ((d_1137_v71_).f27).length(0))
                lhs171_.f5 = rhs197_
                d_1134_v68_ = rhs198_
                lhs172_[lhs173_] = rhs199_
                lhs174_[lhs175_] = rhs200_
                lhs176_[lhs177_] = rhs201_
                d_1145_v78_: _dafny.MultiSet
                d_1145_v78_ = _dafny.MultiSet([(d_1048_v2_)[default__.safeIndex(582, (d_1048_v2_).length(0))]])
                d_1145_v78_ = d_1145_v78_
            (globalState).f15 = True
            (globalState).f6 = not(False)
            d_1046_v0_ = (d_1046_v0_) + (((d_1051_v3_).f36) + (d_1046_v0_))
            d_1146_v79_: _dafny.Seq
            d_1146_v79_ = _dafny.SeqWithoutIsStrInference([p2])
            (globalState).f15 = (p2) in (d_1146_v79_)
        index178_ = default__.safeIndex(830, (d_1048_v2_).length(0))
        (d_1048_v2_)[index178_] = p2
        index179_ = default__.safeIndex(830, (d_1048_v2_).length(0))
        rhs202_ = (p0) * (p0)
        rhs203_ = p0
        rhs204_ = d_1048_v2_
        rhs205_ = False
        lhs178_ = globalState
        lhs179_ = globalState
        lhs180_ = d_1048_v2_
        lhs181_ = default__.safeIndex(830, (d_1048_v2_).length(0))
        lhs178_.f5 = rhs202_
        lhs179_.f4 = rhs203_
        d_1048_v2_ = rhs204_
        lhs180_[lhs181_] = rhs205_
        d_1147_v80_: _dafny.Map
        d_1147_v80_ = _dafny.Map({_dafny.Map({(d_1048_v2_)[default__.safeIndex(830, (d_1048_v2_).length(0))]: False}): p0})
        d_1148_v81_: _dafny.Map
        d_1148_v81_ = _dafny.Map({(d_1048_v2_)[default__.safeIndex(830, (d_1048_v2_).length(0))]: not((d_1048_v2_)[default__.safeIndex(830, (d_1048_v2_).length(0))])})
        hi4_ = (0) - (-130)
        for d_1149_i6_ in range(((d_1147_v80_)[d_1148_v81_] if (d_1148_v81_) in (d_1147_v80_) else len((d_1051_v3_).f36)), hi4_):
            d_1150_v82_: _dafny.Map
            d_1150_v82_ = _dafny.Map({p2: d_1052_v5_})
            d_1151_v83_: C2
            nw174_ = C2()
            nw174_.ctor__(_dafny.CodePoint('t'), p2, d_1048_v2_, not((d_1052_v5_).isdisjoint(((d_1150_v82_)[(d_1048_v2_)[default__.safeIndex(830, (d_1048_v2_).length(0))]] if ((d_1048_v2_)[default__.safeIndex(830, (d_1048_v2_).length(0))]) in (d_1150_v82_) else d_1052_v5_))))
            d_1151_v83_ = nw174_
            d_1152_v84_: _dafny.Map
            d_1152_v84_ = _dafny.Map({True: p3})
            d_1153_v85_: _dafny.Map
            d_1153_v85_ = _dafny.Map({False: _dafny.SeqWithoutIsStrInference([d_1152_v84_])})
            d_1154_v86_: _dafny.Map
            d_1154_v86_ = _dafny.Map({len(d_1052_v5_): ((d_1153_v85_)[(d_1048_v2_)[default__.safeIndex(830, (d_1048_v2_).length(0))]] if ((d_1048_v2_)[default__.safeIndex(830, (d_1048_v2_).length(0))]) in (d_1153_v85_) else _dafny.SeqWithoutIsStrInference([d_1152_v84_, d_1152_v84_]))})
            d_1155_v87_: _dafny.Seq
            d_1155_v87_ = _dafny.SeqWithoutIsStrInference([d_1152_v84_])
            d_1154_v86_ = (d_1154_v86_).set(d_1149_i6_, d_1155_v87_)
            d_1156_v88_: _dafny.Seq
            d_1156_v88_ = _dafny.SeqWithoutIsStrInference([d_1151_v83_.f38, d_1151_v83_.f38, d_1151_v83_.f38])
            d_1157_v89_: _dafny.MultiSet
            d_1157_v89_ = _dafny.MultiSet([d_1151_v83_.f38])
            if (d_1157_v89_).ispropersubset(_dafny.MultiSet(((d_1156_v88_) + (d_1156_v88_)).set(default__.safeIndex(len((d_1051_v3_).f36), len((d_1156_v88_) + (d_1156_v88_))), (d_1048_v2_)[default__.safeIndex(830, (d_1048_v2_).length(0))]))):
                d_1158_v90_: _dafny.Seq
                d_1158_v90_ = _dafny.SeqWithoutIsStrInference([d_1048_v2_])
                index180_ = default__.safeIndex(830, (d_1048_v2_).length(0))
                index181_ = default__.safeIndex(830, (d_1048_v2_).length(0))
                rhs206_ = (d_1158_v90_) != (d_1158_v90_)
                rhs207_ = (-605) == (p0)
                rhs208_ = len(default__.fm18(p2, p3, (d_1151_v83_).f37, globalState))
                lhs182_ = d_1048_v2_
                lhs183_ = default__.safeIndex(830, (d_1048_v2_).length(0))
                lhs184_ = d_1048_v2_
                lhs185_ = default__.safeIndex(830, (d_1048_v2_).length(0))
                lhs186_ = globalState
                lhs182_[lhs183_] = rhs206_
                lhs184_[lhs185_] = rhs207_
                lhs186_.f5 = rhs208_
                d_1159_v91_: _dafny.Map
                d_1159_v91_ = _dafny.Map({d_1152_v84_: d_1149_i6_})
                d_1160_v92_: D8
                d_1160_v92_ = D8_DC17(d_1159_v91_)
                d_1161_v93_: D8
                d_1161_v93_ = D8_DC20(d_1160_v92_)
                rhs209_ = 894
                rhs210_ = d_1161_v93_
                lhs187_ = globalState
                lhs187_.f13 = rhs209_
                d_1161_v93_ = rhs210_
                (globalState).f4 = default__.safeModuloInt(p0, p3)
                d_1046_v0_ = (d_1051_v3_).f36
                d_1162_v94_: _dafny.Seq
                d_1162_v94_ = _dafny.SeqWithoutIsStrInference([d_1046_v0_])
                d_1163_v95_: _dafny.Map
                d_1163_v95_ = _dafny.Map({len((d_1162_v94_) + (d_1162_v94_)): d_1048_v2_})
                d_1048_v2_ = ((d_1163_v95_)[d_1149_i6_] if (d_1149_i6_) in (d_1163_v95_) else d_1048_v2_)
            elif True:
                d_1164_v96_: _dafny.Set
                d_1164_v96_ = _dafny.Set({True})
                d_1164_v96_ = d_1164_v96_
                d_1165_v97_: _dafny.Array
                nw175_ = _dafny.Array(_dafny.CodePoint('D'), 21)
                d_1165_v97_ = nw175_
                index182_ = default__.safeIndex(455, (d_1165_v97_).length(0))
                (d_1165_v97_)[index182_] = (d_1151_v83_).f37
                index183_ = default__.safeIndex(455, (d_1165_v97_).length(0))
                (d_1165_v97_)[index183_] = (d_1151_v83_).f37
                d_1166_v98_: _dafny.Map
                d_1166_v98_ = _dafny.Map({d_1149_i6_: d_1164_v96_})
                d_1166_v98_ = (d_1166_v98_).set(((d_1152_v84_)[d_1151_v83_.f38] if (d_1151_v83_.f38) in (d_1152_v84_) else d_1149_i6_), (d_1164_v96_) - (d_1164_v96_))
                (globalState).f21 = (default__.safeDivisionInt(p0, (0) - ((0) - (p3)))) - ((d_1151_v83_).fm20(globalState))
                d_1167_v99_: _dafny.Map
                d_1167_v99_ = _dafny.Map({d_1152_v84_: p0})
                d_1168_v100_: D8
                d_1168_v100_ = D8_DC20(D8_DC17(d_1167_v99_))
                d_1168_v100_ = d_1168_v100_
            d_1169_v101_: _dafny.Array
            nw176_ = _dafny.Array(None, 14)
            nw176_[int(0)] = (d_1051_v3_).f36
            nw176_[int(1)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "meyjff"))) + (_dafny.SeqWithoutIsStrInference([(d_1151_v83_).f37 for d_1170_i7_ in range(default__.abs(595))]))
            nw176_[int(2)] = d_1046_v0_
            nw176_[int(3)] = (d_1051_v3_).f36
            nw176_[int(4)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqcgi"))) + ((d_1051_v3_).f36)
            nw176_[int(5)] = (d_1051_v3_).f36
            nw176_[int(6)] = ((d_1051_v3_).f36).set(default__.safeIndex(p3, len((d_1051_v3_).f36)), _dafny.CodePoint('t'))
            nw176_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krlvulblo"))
            nw176_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
            nw176_[int(9)] = (d_1051_v3_).f36
            nw176_[int(10)] = d_1046_v0_
            nw176_[int(11)] = _dafny.SeqWithoutIsStrInference([(d_1151_v83_).f37 for d_1171_i8_ in range(default__.abs(525))])
            nw176_[int(12)] = (d_1046_v0_) + ((d_1051_v3_).f36)
            nw176_[int(13)] = _dafny.SeqWithoutIsStrInference([(d_1151_v83_).f37 for d_1172_i9_ in range(default__.abs(733))])
            d_1169_v101_ = nw176_
            index184_ = default__.safeIndex(215, (d_1169_v101_).length(0))
            (d_1169_v101_)[index184_] = (d_1051_v3_).f36
            d_1173_v102_: D9
            d_1173_v102_ = D9_DC21(d_1169_v101_)
            pat_let_tv27_ = d_1173_v102_
            d_1174_v103_: _dafny.Map
            def iife96_(_pat_let19_0):
                def iife97_(d_1175_dt__update__tmp_h1_):
                    def iife98_(_pat_let20_0):
                        def iife99_(d_1176_dt__update_hcf39_h0_):
                            return D9_DC24(d_1176_dt__update_hcf39_h0_)
                        return iife99_(_pat_let20_0)
                    return iife98_(pat_let_tv27_)
                return iife97_(_pat_let19_0)
            d_1174_v103_ = _dafny.Map({iife96_(D9_DC24(d_1173_v102_)): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgbvpofn"))})
            d_1177_v104_: D9
            d_1177_v104_ = D9_DC24(d_1173_v102_)
            index185_ = default__.safeIndex(215, (d_1169_v101_).length(0))
            (d_1169_v101_)[index185_] = ((d_1051_v3_).f36) + ((((d_1174_v103_)[d_1177_v104_] if (d_1177_v104_) in (d_1174_v103_) else d_1046_v0_)).set(default__.safeIndex(525, len(((d_1174_v103_)[d_1177_v104_] if (d_1177_v104_) in (d_1174_v103_) else d_1046_v0_))), (d_1151_v83_).f37))
        d_1178_v105_: _dafny.Seq
        d_1178_v105_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p3: p2})])
        (globalState).f5 = (len(d_1178_v105_)) - ((p0) + (p0))
        index186_ = default__.safeIndex(726, (p1).length(0))
        (p1)[index186_] = p3
        index187_ = default__.safeIndex(726, (p1).length(0))
        (p1)[index187_] = (0) - (default__.safeDivisionInt(p0, len(_dafny.Map({(d_1048_v2_)[default__.safeIndex(830, (d_1048_v2_).length(0))]: (d_1048_v2_)[default__.safeIndex(830, (d_1048_v2_).length(0))]}))))


class C5(T1, T0):
    def  __init__(self):
        self._f29: int = int(0)
        self._f27: _dafny.Array = _dafny.Array(None, 0)
        self._f28: bool = False
        self._f33: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f33, f29, f27, f28):
        (self)._f33 = f33
        (self).f29 = f29
        (self)._f27 = f27
        (self)._f28 = f28

    def fm6(self, p0, p1, p2, p3, globalState):
        return self.f29

    def fm7(self, p0, globalState):
        return (self.f29) * (self.f29)

    def m7(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_1179_v0_: _dafny.Array
        def lambda43_(d_1180_i0_):
            return (d_1180_i0_) * (self.f29)

        init26_ = lambda43_
        nw177_ = _dafny.Array(None, 8)
        for i0_26_ in range(nw177_.length(0)):
            nw177_[i0_26_] = init26_(i0_26_)
        d_1179_v0_ = nw177_
        d_1181_v1_: _dafny.MultiSet
        d_1181_v1_ = _dafny.MultiSet([p0, True])
        index188_ = default__.safeIndex(337, (d_1179_v0_).length(0))
        (d_1179_v0_)[index188_] = (d_1181_v1_).cardinality
        d_1182_v2_: D3
        d_1182_v2_ = D3_DC10(p0, (self).f28)
        d_1183_v3_: _dafny.Set
        d_1183_v3_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([not((self).f28)]))})
        d_1184_v4_: _dafny.Map
        d_1184_v4_ = _dafny.Map({d_1182_v2_: d_1183_v3_})
        index189_ = default__.safeIndex(337, (d_1179_v0_).length(0))
        (d_1179_v0_)[index189_] = len((d_1184_v4_ if p0 else (d_1184_v4_) | (d_1184_v4_)))
        hi5_ = (0) - (default__.safeModuloInt(self.f29, (d_1179_v0_)[default__.safeIndex(337, (d_1179_v0_).length(0))]))
        for d_1185_i1_ in range((d_1179_v0_)[default__.safeIndex(337, (d_1179_v0_).length(0))], hi5_):
            (globalState).f17 = p0
            d_1186_v8_: _dafny.MultiSet
            def iife100_():
                coll58_ = _dafny.Set()
                compr_58_: int
                for compr_58_ in _dafny.IntegerRange(-224, 926):
                    d_1187_v5_: int = compr_58_
                    if ((-224) <= (d_1187_v5_)) and ((d_1187_v5_) < (926)):
                        def iife101_():
                            def iife103_():
                                coll61_ = _dafny.Map()
                                compr_61_: str
                                for compr_61_ in (_dafny.Map({_dafny.CodePoint('s'): d_1185_i1_})).keys.Elements:
                                    d_1188_v6_: str = compr_61_
                                    if (d_1188_v6_) in (_dafny.Map({_dafny.CodePoint('s'): d_1185_i1_})):
                                        coll61_[d_1188_v6_] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1189_i2_ in range(default__.abs(853))]))
                                return _dafny.Map(coll61_)
                            coll59_ = _dafny.Set()
                            def iife102_():
                                coll60_ = _dafny.Map()
                                compr_60_: str
                                for compr_60_ in (_dafny.Map({_dafny.CodePoint('s'): d_1185_i1_})).keys.Elements:
                                    d_1188_v6_: str = compr_60_
                                    if (d_1188_v6_) in (_dafny.Map({_dafny.CodePoint('s'): d_1185_i1_})):
                                        coll60_[d_1188_v6_] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1189_i2_ in range(default__.abs(853))]))
                                return _dafny.Map(coll60_)
                            compr_59_: str
                            for compr_59_ in (iife102_()
                            ).keys.Elements:
                                d_1190_v7_: str = compr_59_
                                if (d_1190_v7_) in (iife103_()
                                ):
                                    coll59_ = coll59_.union(_dafny.Set([d_1190_v7_]))
                            return _dafny.Set(coll59_)
                        coll58_ = coll58_.union(_dafny.Set([(d_1187_v5_) * (len(iife101_()
))]))
                return _dafny.Set(coll58_)
            d_1186_v8_ = _dafny.MultiSet([iife100_()
            , d_1183_v3_, (d_1183_v3_).intersection(d_1183_v3_)])
            d_1191_v9_: _dafny.Seq
            d_1191_v9_ = _dafny.SeqWithoutIsStrInference([d_1183_v3_])
            d_1186_v8_ = ((_dafny.MultiSet(d_1191_v9_)) - (d_1186_v8_)).set(d_1183_v3_, default__.abs(((d_1179_v0_)[default__.safeIndex(337, (d_1179_v0_).length(0))]) * (self.f29)))
            d_1192_v10_: str
            d_1192_v10_ = _dafny.CodePoint('g')
            d_1193_v11_: D1
            d_1193_v11_ = D1_DC7(d_1192_v10_, default__.fm9((self).f28, globalState))
            source18_ = default__.fm8((d_1193_v11_ if True else d_1193_v11_), globalState)
            if source18_.is_DC10:
                d_1194___mcc_h0_ = source18_.cf17
                d_1195___mcc_h1_ = source18_.cf18
                d_1196_cf18_ = d_1195___mcc_h1_
                d_1197_cf17_ = d_1194___mcc_h0_
                d_1197_cf17_ = (self).f28
                d_1198_v12_: _dafny.Seq
                d_1198_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
                d_1199_v14_: _dafny.MultiSet
                d_1199_v14_ = _dafny.MultiSet([d_1198_v12_, d_1198_v12_])
                def iife104_():
                    coll62_ = _dafny.Map()
                    compr_62_: _dafny.Seq
                    for compr_62_ in (d_1199_v14_).Elements:
                        d_1200_v13_: _dafny.Seq = compr_62_
                        if (d_1200_v13_) in (d_1199_v14_):
                            coll62_[d_1200_v13_] = d_1193_v11_
                    return _dafny.Map(coll62_)
                (globalState).f4 = len((_dafny.Map({d_1198_v12_: d_1193_v11_}) if True else iife104_()
                ))
                index190_ = default__.safeIndex(337, (d_1179_v0_).length(0))
                (d_1179_v0_)[index190_] = d_1185_i1_
                d_1201_v15_: _dafny.Array
                nw178_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_1201_v15_ = nw178_
                d_1202_v16_: _dafny.Map
                d_1202_v16_ = _dafny.Map({d_1201_v15_: 594})
                (globalState).f21 = ((d_1202_v16_)[d_1201_v15_] if (d_1201_v15_) in (d_1202_v16_) else self.f29)
            elif True:
                d_1203___mcc_h2_ = source18_.cf16
                d_1204_cf16_ = d_1203___mcc_h2_
                (globalState).f25 = d_1192_v10_
                d_1205_v17_: _dafny.Map
                d_1205_v17_ = _dafny.Map({default__.fm10(_dafny.CodePoint('u'), globalState): (self).f28})
                d_1206_v18_: _dafny.Seq
                d_1206_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrlr"))
                d_1205_v17_ = default__.fm11((d_1181_v1_).ispropersubset(d_1181_v1_), not((d_1206_v18_) <= (_dafny.SeqWithoutIsStrInference([d_1192_v10_ for d_1207_i3_ in range(default__.abs(-321))]))), p0, globalState)
                d_1208_v19_: D0
                d_1208_v19_ = D0_DC3(self.f29, p0, p0)
                d_1209_v20_: _dafny.Map
                d_1209_v20_ = _dafny.Map({d_1208_v19_: d_1185_i1_})
                d_1210_v21_: C3
                nw179_ = C3()
                nw179_.ctor__((self.f29) * ((d_1179_v0_)[default__.safeIndex(337, (d_1179_v0_).length(0))]), d_1209_v20_, self.f29, (self).f27, p0)
                d_1210_v21_ = nw179_
                d_1211_v22_: _dafny.Array
                nw180_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
                d_1211_v22_ = nw180_
                index191_ = default__.safeIndex(199, (d_1211_v22_).length(0))
                (d_1211_v22_)[index191_] = d_1206_v18_
                index192_ = default__.safeIndex(199, (d_1211_v22_).length(0))
                (d_1211_v22_)[index192_] = (d_1206_v18_) + (((p2)[p0] if (p0) in (p2) else (d_1206_v18_).set(default__.safeIndex(self.f29, len(d_1206_v18_)), d_1192_v10_)))
            (globalState).f4 = (0) - (self.f29)
        d_1212_i4_: int
        d_1212_i4_ = 0
        with _dafny.label("4"):
            while p0:
                with _dafny.c_label("4"):
                    if (d_1212_i4_) >= (100):
                        raise _dafny.Break("4")
                    d_1212_i4_ = (d_1212_i4_) + (1)
                    d_1213_v23_: _dafny.MultiSet
                    d_1213_v23_ = _dafny.MultiSet([(d_1179_v0_)[default__.safeIndex(337, (d_1179_v0_).length(0))]])
                    d_1214_v24_: _dafny.Map
                    d_1214_v24_ = _dafny.Map({(d_1179_v0_)[default__.safeIndex(337, (d_1179_v0_).length(0))]: d_1179_v0_})
                    d_1215_v25_: _dafny.Map
                    d_1215_v25_ = _dafny.Map({p0: len(d_1214_v24_)})
                    d_1216_v26_: _dafny.Map
                    d_1216_v26_ = _dafny.Map({(self).f28: d_1215_v25_})
                    d_1217_v27_: _dafny.MultiSet
                    d_1217_v27_ = _dafny.MultiSet([(_dafny.Map({not(p0): (d_1179_v0_)[default__.safeIndex(337, (d_1179_v0_).length(0))]})).set(p0, len(_dafny.Map({p0: d_1213_v23_}))), d_1215_v25_, ((d_1216_v26_)[(self).f28] if ((self).f28) in (d_1216_v26_) else d_1215_v25_)])
                    (globalState).f6 = (d_1217_v27_).isdisjoint((d_1217_v27_).set(d_1215_v25_, default__.abs(922)))
                    d_1218_v28_: _dafny.Seq
                    d_1218_v28_ = _dafny.SeqWithoutIsStrInference([(d_1179_v0_)[default__.safeIndex(337, (d_1179_v0_).length(0))], (d_1179_v0_)[default__.safeIndex(337, (d_1179_v0_).length(0))], self.f29, 441])
                    d_1219_v29_: _dafny.Seq
                    d_1219_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcumuyosc"))
                    d_1220_v30_: _dafny.Seq
                    d_1220_v30_ = _dafny.SeqWithoutIsStrInference([len(default__.fm25((d_1218_v28_)[default__.safeIndex(len(d_1219_v29_), len(d_1218_v28_))], (d_1179_v0_)[default__.safeIndex(337, (d_1179_v0_).length(0))], self.f29, globalState))])
                    d_1221_v31_: _dafny.Map
                    d_1221_v31_ = _dafny.Map({p0: (self).f28})
                    d_1222_v32_: str
                    d_1222_v32_ = _dafny.CodePoint('w')
                    d_1223_v33_: C2
                    nw181_ = C2()
                    nw181_.ctor__(d_1222_v32_, p0, (self).f27, True)
                    d_1223_v33_ = nw181_
                    d_1224_v34_: _dafny.Map
                    d_1224_v34_ = _dafny.Map({d_1223_v33_: d_1222_v32_})
                    (globalState).f13 = (len(d_1218_v28_)) - (len((d_1224_v34_ if default__.fm16(d_1221_v31_, globalState) else d_1224_v34_)))
                    d_1220_v30_ = d_1220_v30_
                    d_1225_v35_: C4
                    nw182_ = C4()
                    nw182_.ctor__()
                    d_1225_v35_ = nw182_
                    pass
            pass
        d_1226_v36_: _dafny.Seq
        d_1226_v36_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_1179_v0_)[default__.safeIndex(337, (d_1179_v0_).length(0))])])
        d_1227_v37_: _dafny.Map
        d_1227_v37_ = _dafny.Map({(_dafny.Set({(d_1226_v36_)})).issubset(_dafny.Set({d_1226_v36_, d_1226_v36_})): (243) != ((d_1179_v0_)[default__.safeIndex(337, (d_1179_v0_).length(0))])})
        d_1227_v37_ = d_1227_v37_
        d_1228_v38_: _dafny.Seq
        d_1228_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
        (globalState).f5 = len(d_1228_v38_)
        (globalState).f15 = not(not (p0) or (p0))
        r0 = self.f29
        return r0

    @property
    def f33(self):
        return self._f33

class C6(T1, T0):
    def  __init__(self):
        self._f29: int = int(0)
        self._f27: _dafny.Array = _dafny.Array(None, 0)
        self._f28: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f29, f27, f28):
        (self).f29 = f29
        (self)._f27 = f27
        (self)._f28 = f28

    def fm5(self, globalState):
        return self.f29

    def m6(self, p0, p1, p2, globalState):
        (globalState).f6 = not(p0)
        d_1229_v0_: _dafny.Set
        d_1229_v0_ = _dafny.Set({p2, -885, p2})
        d_1230_v1_: str
        d_1230_v1_ = _dafny.CodePoint('g')
        d_1231_v2_: _dafny.Map
        d_1231_v2_ = _dafny.Map({self.f29: d_1230_v1_})
        d_1232_v4_: _dafny.Map
        def iife105_():
            coll63_ = _dafny.Set()
            compr_63_: int
            for compr_63_ in (d_1231_v2_).keys.Elements:
                d_1233_v3_: int = compr_63_
                if (d_1233_v3_) in (d_1231_v2_):
                    coll63_ = coll63_.union(_dafny.Set([(d_1233_v3_) - (-91)]))
            return _dafny.Set(coll63_)
        d_1232_v4_ = _dafny.Map({(0) - (self.f29): iife105_()
        })
        d_1234_v5_: D3
        d_1234_v5_ = D3_DC9(d_1232_v4_)
        d_1229_v0_ = default__.fm0((d_1234_v5_).cf16, p0, not(((self).f28) == (p1)), globalState)
        d_1235_v6_: _dafny.Map
        d_1236_v7_: int
        out23_: _dafny.Map
        out24_: int
        out23_, out24_ = default__.m0(globalState)
        d_1235_v6_ = out23_
        d_1236_v7_ = out24_
        d_1237_v8_: _dafny.Seq
        d_1237_v8_ = _dafny.SeqWithoutIsStrInference([d_1229_v0_])
        d_1238_v9_: _dafny.Seq
        d_1238_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwuoanirg"))
        d_1239_v10_: D0
        d_1239_v10_ = D0_DC2((d_1237_v8_)[default__.safeIndex(len(d_1238_v9_), len(d_1237_v8_))], d_1238_v9_)
        source19_ = d_1239_v10_
        if source19_.is_DC0:
            d_1240___mcc_h0_ = source19_.cf0
            d_1241___mcc_h1_ = source19_.cf1
            d_1242___mcc_h2_ = source19_.cf2
            d_1243___mcc_h3_ = source19_.cf3
            d_1244___mcc_h4_ = source19_.cf4
            d_1245_cf4_ = d_1244___mcc_h4_
            d_1246_cf3_ = d_1243___mcc_h3_
            d_1247_cf2_ = d_1242___mcc_h2_
            d_1248_cf1_ = d_1241___mcc_h1_
            d_1249_cf0_ = d_1240___mcc_h0_
            d_1250_v11_: _dafny.Array
            nw183_ = _dafny.Array(False, 21)
            d_1250_v11_ = nw183_
            d_1250_v11_ = (self).f27
            d_1251_v12_: _dafny.Array
            def lambda44_(d_1252_p1_, d_1253_cf4_):
                def lambda45_(d_1254_i0_):
                    return (d_1254_i0_) + (len(_dafny.Map({d_1252_p1_: d_1253_cf4_})))

                return lambda45_

            init27_ = lambda44_(p1, d_1245_cf4_)
            nw184_ = _dafny.Array(None, 21)
            for i0_27_ in range(nw184_.length(0)):
                nw184_[i0_27_] = init27_(i0_27_)
            d_1251_v12_ = nw184_
            index193_ = default__.safeIndex(40, (d_1251_v12_).length(0))
            (d_1251_v12_)[index193_] = self.f29
            index194_ = default__.safeIndex(40, (d_1251_v12_).length(0))
            (d_1251_v12_)[index194_] = self.f29
            d_1255_v13_: D0
            d_1255_v13_ = D0_DC1()
            d_1255_v13_ = d_1255_v13_
            d_1256_v14_: C2
            nw185_ = C2()
            nw185_.ctor__(d_1246_cf3_, True, d_1250_v11_, False)
            d_1256_v14_ = nw185_
        elif source19_.is_DC1:
            d_1257_v15_: _dafny.Map
            d_1257_v15_ = _dafny.Map({d_1236_v7_: self.f29})
            (globalState).f4 = (d_1236_v7_ if not((p1 if p1 else (self).f28)) else (self.f29) - (len(d_1257_v15_)))
            (globalState).f5 = -700
            d_1258_v16_: D0
            d_1258_v16_ = D0_DC3(self.f29, p0, not(p1))
            d_1259_v17_: _dafny.Seq
            d_1259_v17_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_1260_v18_: _dafny.Map
            d_1260_v18_ = _dafny.Map({d_1258_v16_: len(d_1259_v17_)})
            d_1261_v19_: C3
            nw186_ = C3()
            nw186_.ctor__(self.f29, (d_1260_v18_ if p1 else _dafny.Map({d_1258_v16_: (0) - (d_1236_v7_)})), ((0) - (p2)) + (p2), (self).f27, (self).f28)
            d_1261_v19_ = nw186_
            d_1262_v20_: _dafny.Array
            def lambda46_(d_1263_p0_, d_1264_p1_):
                def lambda47_(d_1265_i1_):
                    return default__.safeDivisionInt(d_1265_i1_, len(_dafny.Map({d_1263_p0_: d_1264_p1_})))

                return lambda47_

            init28_ = lambda46_(p0, p1)
            nw187_ = _dafny.Array(None, 16)
            for i0_28_ in range(nw187_.length(0)):
                nw187_[i0_28_] = init28_(i0_28_)
            d_1262_v20_ = nw187_
            index195_ = default__.safeIndex(897, (d_1262_v20_).length(0))
            (d_1262_v20_)[index195_] = (self.f29) - ((d_1261_v19_).f34)
            d_1266_v21_: _dafny.Seq
            d_1266_v21_ = _dafny.SeqWithoutIsStrInference([d_1262_v20_])
            d_1267_v22_: _dafny.Map
            d_1267_v22_ = _dafny.Map({(self).f28: (self).f27})
            d_1268_v23_: C5
            nw188_ = C5()
            nw188_.ctor__(d_1266_v21_, self.f29, ((d_1267_v22_)[True] if (True) in (d_1267_v22_) else (self).f27), (self).f28)
            d_1268_v23_ = nw188_
            d_1269_v24_: _dafny.MultiSet
            d_1269_v24_ = _dafny.MultiSet([(d_1266_v21_)[default__.safeIndex((d_1261_v19_).f34, len(d_1266_v21_))]])
            d_1270_v25_: _dafny.Map
            d_1270_v25_ = _dafny.Map({d_1268_v23_: (d_1269_v24_).cardinality})
            index196_ = default__.safeIndex(897, (d_1262_v20_).length(0))
            (d_1262_v20_)[index196_] = (default__.safeModuloInt((0) - (((d_1270_v25_)[d_1268_v23_] if (d_1268_v23_) in (d_1270_v25_) else d_1236_v7_)), (d_1261_v19_).f34)) - (d_1236_v7_)
        elif source19_.is_DC2:
            d_1271___mcc_h5_ = source19_.cf5
            d_1272___mcc_h6_ = source19_.cf6
            d_1273_cf6_ = d_1272___mcc_h6_
            d_1274_cf5_ = d_1271___mcc_h5_
            d_1275_v26_: _dafny.Map
            d_1276_v27_: int
            out25_: _dafny.Map
            out26_: int
            out25_, out26_ = default__.m0(globalState)
            d_1275_v26_ = out25_
            d_1276_v27_ = out26_
            d_1277_v28_: _dafny.MultiSet
            d_1277_v28_ = _dafny.MultiSet([default__.fm46(p0, globalState)])
            rhs211_ = d_1238_v9_
            rhs212_ = ((d_1277_v28_) | (d_1277_v28_)) - (d_1277_v28_)
            rhs213_ = p0
            lhs188_ = globalState
            d_1273_cf6_ = rhs211_
            d_1277_v28_ = rhs212_
            lhs188_.f6 = rhs213_
            d_1278_v29_: _dafny.Array
            nw189_ = _dafny.Array(_dafny.Seq({}), 8)
            d_1278_v29_ = nw189_
            d_1279_v30_: _dafny.MultiSet
            d_1279_v30_ = _dafny.MultiSet([d_1236_v7_])
            index197_ = default__.safeIndex(8, (d_1278_v29_).length(0))
            (d_1278_v29_)[index197_] = default__.fm49(d_1279_v30_, globalState)
            d_1280_v31_: _dafny.Seq
            d_1280_v31_ = _dafny.SeqWithoutIsStrInference([(self).f28])
            d_1281_v32_: _dafny.Seq
            d_1281_v32_ = _dafny.SeqWithoutIsStrInference([d_1280_v31_])
            index198_ = default__.safeIndex(8, (d_1278_v29_).length(0))
            (d_1278_v29_)[index198_] = (d_1281_v32_) + ((d_1281_v32_) + (d_1281_v32_))
            (globalState).f6 = p0
        elif source19_.is_DC3:
            d_1282___mcc_h7_ = source19_.cf7
            d_1283___mcc_h8_ = source19_.cf8
            d_1284___mcc_h9_ = source19_.cf9
            d_1285_cf9_ = d_1284___mcc_h9_
            d_1286_cf8_ = d_1283___mcc_h8_
            d_1287_cf7_ = d_1282___mcc_h7_
            (globalState).f13 = (default__.fm36(globalState)) + (d_1236_v7_)
            d_1288_v33_: D3
            d_1288_v33_ = D3_DC10(p1, d_1285_cf9_)
            d_1289_v34_: _dafny.Set
            d_1289_v34_ = _dafny.Set({d_1288_v33_})
            d_1290_v35_: _dafny.MultiSet
            d_1290_v35_ = _dafny.MultiSet([(self).f28, not(d_1286_cf8_), p1])
            d_1291_v36_: _dafny.Seq
            d_1291_v36_ = _dafny.SeqWithoutIsStrInference([((d_1290_v35_)[False] if (False) in (d_1290_v35_) else self.f29), d_1236_v7_])
            rhs214_ = d_1236_v7_
            rhs215_ = (_dafny.Set({d_1288_v33_}) if True else _dafny.Set({D3_DC10(d_1285_cf9_, (self).f28)}))
            rhs216_ = d_1291_v36_
            lhs189_ = globalState
            lhs189_.f21 = rhs214_
            d_1289_v34_ = rhs215_
            d_1291_v36_ = rhs216_
            (self).f29 = default__.fm36(globalState)
            d_1292_v37_: C4
            nw190_ = C4()
            nw190_.ctor__()
            d_1292_v37_ = nw190_
        elif True:
            d_1293___mcc_h10_ = source19_.cf10
            d_1294_cf10_ = d_1293___mcc_h10_
            d_1295_v38_: _dafny.Map
            d_1295_v38_ = _dafny.Map({p2: p1})
            d_1296_v39_: C0
            nw191_ = C0()
            nw191_.ctor__((d_1238_v9_) + (d_1238_v9_), (0) - ((len(d_1295_v38_) if p1 else p2)), (self).f27, (self).f28)
            d_1296_v39_ = nw191_
            d_1297_v40_: _dafny.Set
            d_1297_v40_ = _dafny.Set({p1, default__.fm40(globalState)})
            d_1298_v41_: _dafny.Seq
            d_1298_v41_ = _dafny.SeqWithoutIsStrInference([(self).f28, p0])
            d_1299_v42_: _dafny.Array
            nw192_ = _dafny.Array(None, 7)
            nw192_[int(0)] = (0) - (len((d_1297_v40_).intersection(d_1297_v40_)))
            nw192_[int(1)] = (0) - (d_1236_v7_)
            nw192_[int(2)] = (0) - (p2)
            nw192_[int(3)] = self.f29
            nw192_[int(4)] = self.f29
            nw192_[int(5)] = len(d_1298_v41_)
            nw192_[int(6)] = p2
            d_1299_v42_ = nw192_
            d_1300_v43_: _dafny.Seq
            d_1300_v43_ = _dafny.SeqWithoutIsStrInference([d_1299_v42_])
            d_1301_v44_: _dafny.MultiSet
            d_1301_v44_ = _dafny.MultiSet([d_1299_v42_, (d_1300_v43_)[default__.safeIndex(d_1236_v7_, len(d_1300_v43_))], d_1299_v42_])
            index199_ = default__.safeIndex(355, (d_1299_v42_).length(0))
            (d_1299_v42_)[index199_] = (d_1301_v44_).cardinality
            index200_ = default__.safeIndex(355, (d_1299_v42_).length(0))
            (d_1299_v42_)[index200_] = (self).fm5(globalState)
            d_1302_v45_: _dafny.Map
            d_1302_v45_ = _dafny.Map({d_1230_v1_: d_1300_v43_})
            d_1303_v46_: C5
            nw193_ = C5()
            nw193_.ctor__(((d_1302_v45_)[_dafny.CodePoint('q')] if (_dafny.CodePoint('q')) in (d_1302_v45_) else d_1300_v43_), d_1236_v7_, (self).f27, (d_1298_v41_)[default__.safeIndex((d_1299_v42_)[default__.safeIndex(355, (d_1299_v42_).length(0))], len(d_1298_v41_))])
            d_1303_v46_ = nw193_
            d_1304_v47_: _dafny.Seq
            d_1304_v47_ = _dafny.SeqWithoutIsStrInference([d_1236_v7_, (d_1296_v39_).fm15(globalState), 406, (0) - (-368), default__.fm36(globalState)])
            rhs217_ = ((d_1304_v47_)[default__.safeIndex(p2, len(d_1304_v47_))]) * (default__.safeModuloInt(d_1236_v7_, default__.fm36(globalState)))
            rhs218_ = (d_1298_v41_) <= (d_1298_v41_)
            rhs219_ = (d_1303_v46_).f33
            rhs220_ = not (not(p0)) or (p1)
            lhs190_ = globalState
            lhs191_ = globalState
            lhs192_ = globalState
            lhs190_.f13 = rhs217_
            lhs191_.f15 = rhs218_
            d_1300_v43_ = rhs219_
            lhs192_.f6 = rhs220_
        d_1305_v48_: _dafny.Map
        d_1305_v48_ = _dafny.Map({(self.f29) - (d_1236_v7_): p0})
        d_1305_v48_ = (d_1305_v48_).set(len(d_1238_v9_), (self).f28)
        d_1306_v49_: _dafny.Seq
        d_1306_v49_ = _dafny.SeqWithoutIsStrInference([p1])
        d_1307_v50_: _dafny.MultiSet
        d_1307_v50_ = _dafny.MultiSet([p0, not(p0), (self).f28, False, True])
        d_1308_v51_: _dafny.Seq
        d_1308_v51_ = _dafny.SeqWithoutIsStrInference([d_1307_v50_])
        index201_ = default__.safeIndex(359, ((self).f27).length(0))
        ((self).f27)[index201_] = (d_1306_v49_)[default__.safeIndex(len(_dafny.Set({d_1308_v51_})), len(d_1306_v49_))]
        index202_ = default__.safeIndex(359, ((self).f27).length(0))
        ((self).f27)[index202_] = (p2) == (d_1236_v7_)


class C7(T0, T1):
    def  __init__(self):
        self._f29: int = int(0)
        self._f27: _dafny.Array = _dafny.Array(None, 0)
        self._f28: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f27, f28, f29):
        (self)._f27 = f27
        (self)._f28 = f28
        (self).f29 = f29

    def m5(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_1309_v0_: str
        d_1309_v0_ = _dafny.CodePoint('e')
        d_1310_v1_: _dafny.Map
        d_1310_v1_ = _dafny.Map({(self).f28: (0) - (self.f29)})
        d_1311_v2_: _dafny.Seq
        d_1311_v2_ = _dafny.SeqWithoutIsStrInference([(0) - (((d_1310_v1_)[True] if (True) in (d_1310_v1_) else self.f29)), self.f29])
        rhs221_ = self.f29
        rhs222_ = d_1309_v0_
        rhs223_ = not(not((self).f28))
        rhs224_ = default__.safeModuloInt(self.f29, (d_1311_v2_)[default__.safeIndex(self.f29, len(d_1311_v2_))])
        lhs193_ = globalState
        lhs194_ = globalState
        r0 = rhs221_
        lhs193_.f25 = rhs222_
        lhs194_.f17 = rhs223_
        r0 = rhs224_
        d_1312_v3_: _dafny.Seq
        d_1312_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yaenitp"))
        (globalState).f17 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhmgaits"))) != ((d_1312_v3_) + (d_1312_v3_))
        (globalState).f6 = ((self).f28 if (self).f28 else ((self).f28 if (self).f28 else (self).f28))
        d_1313_v4_: _dafny.Seq
        d_1313_v4_ = d_1311_v2_
        source20_ = d_1313_v4_
        d_1314___mcc_h0_ = source20_
        d_1315_cf15_ = d_1314___mcc_h0_
        (globalState).f25 = (d_1312_v3_)[default__.safeIndex(self.f29, len(d_1312_v3_))]
        d_1309_v0_ = d_1309_v0_
        d_1316_v5_: D1
        d_1316_v5_ = D1_DC6(self.f29)
        d_1317_v6_: _dafny.MultiSet
        d_1317_v6_ = _dafny.MultiSet([d_1316_v5_, d_1316_v5_])
        d_1318_v7_: _dafny.Seq
        d_1318_v7_ = _dafny.SeqWithoutIsStrInference([d_1317_v6_])
        d_1319_v8_: _dafny.Seq
        d_1319_v8_ = _dafny.SeqWithoutIsStrInference([True])
        d_1320_v9_: _dafny.Map
        d_1320_v9_ = _dafny.Map({self.f29: d_1309_v0_})
        d_1321_v10_: _dafny.Map
        d_1321_v10_ = _dafny.Map({(self).f28: d_1320_v9_})
        d_1322_v11_: _dafny.Seq
        d_1322_v11_ = _dafny.SeqWithoutIsStrInference([((d_1321_v10_)[False] if (False) in (d_1321_v10_) else d_1320_v9_)])
        d_1323_v12_: _dafny.Map
        d_1323_v12_ = _dafny.Map({self.f29: _dafny.MultiSet([d_1316_v5_, d_1316_v5_])})
        d_1324_v13_: _dafny.Map
        d_1324_v13_ = _dafny.Map({self.f29: (self).f28})
        d_1325_v14_: _dafny.Array
        nw194_ = _dafny.Array(None, 29)
        nw194_[int(0)] = self.f29
        nw194_[int(1)] = self.f29
        nw194_[int(2)] = default__.fm4((d_1318_v7_)[default__.safeIndex(self.f29, len(d_1318_v7_))], globalState)
        nw194_[int(3)] = self.f29
        nw194_[int(4)] = len(d_1312_v3_)
        nw194_[int(5)] = self.f29
        nw194_[int(6)] = self.f29
        nw194_[int(7)] = len(d_1319_v8_)
        nw194_[int(8)] = self.f29
        nw194_[int(9)] = self.f29
        nw194_[int(10)] = ((d_1310_v1_)[True] if (True) in (d_1310_v1_) else self.f29)
        nw194_[int(11)] = self.f29
        nw194_[int(12)] = 220
        nw194_[int(13)] = self.f29
        nw194_[int(14)] = len((d_1322_v11_)[default__.safeIndex(self.f29, len(d_1322_v11_))])
        nw194_[int(15)] = (0) - ((0) - (self.f29))
        nw194_[int(16)] = self.f29
        nw194_[int(17)] = self.f29
        nw194_[int(18)] = 824
        nw194_[int(19)] = self.f29
        nw194_[int(20)] = self.f29
        nw194_[int(21)] = 20
        nw194_[int(22)] = 360
        nw194_[int(23)] = default__.fm4(((d_1323_v12_)[self.f29] if (self.f29) in (d_1323_v12_) else d_1317_v6_), globalState)
        nw194_[int(24)] = self.f29
        nw194_[int(25)] = len(d_1315_cf15_)
        nw194_[int(26)] = 660
        nw194_[int(27)] = len(d_1312_v3_)
        nw194_[int(28)] = (0) - (len(d_1324_v13_))
        d_1325_v14_ = nw194_
        d_1326_v15_: _dafny.Map
        d_1326_v15_ = _dafny.Map({(True if (self).f28 else (self).f28): d_1325_v14_})
        d_1326_v15_ = (d_1326_v15_).set(not((self).f28), d_1325_v14_)
        d_1327_v16_: _dafny.Seq
        d_1327_v16_ = _dafny.SeqWithoutIsStrInference([d_1325_v14_])
        d_1328_v17_: T1
        nw195_ = C5()
        nw195_.ctor__(d_1327_v16_, self.f29, (self).f27, not((self).f28))
        d_1328_v17_ = nw195_
        d_1329_v18_: _dafny.Map
        d_1329_v18_ = _dafny.Map({d_1328_v17_: self.f29})
        d_1330_v19_: _dafny.Map
        d_1330_v19_ = _dafny.Map({(d_1329_v18_).set(d_1328_v17_, len(d_1312_v3_)): (self).f28})
        d_1331_v20_: D0
        d_1331_v20_ = D0_DC0(d_1330_v19_, (self).f28, d_1328_v17_.f29, d_1309_v0_, (d_1328_v17_).f28)
        d_1332_v21_: _dafny.Map
        d_1332_v21_ = _dafny.Map({(self).f28: (d_1328_v17_).f28})
        d_1333_v22_: _dafny.Array
        nw196_ = _dafny.Array(None, 24)
        nw196_[int(0)] = (902) > (default__.fm4(d_1317_v6_, globalState))
        nw196_[int(1)] = not((self).f28)
        nw196_[int(2)] = (self).f28
        nw196_[int(3)] = not((self).f28)
        nw196_[int(4)] = not((self).f28)
        nw196_[int(5)] = (d_1331_v20_).cf1
        nw196_[int(6)] = (self).f28
        nw196_[int(7)] = (d_1328_v17_).f28
        nw196_[int(8)] = not((d_1328_v17_).f28)
        nw196_[int(9)] = (d_1328_v17_).f28
        nw196_[int(10)] = not(default__.fm10(d_1309_v0_, globalState))
        nw196_[int(11)] = (d_1328_v17_).f28
        nw196_[int(12)] = ((self).f28) == (((d_1332_v21_)[(self).f28] if ((self).f28) in (d_1332_v21_) else (d_1328_v17_).f28))
        nw196_[int(13)] = (d_1328_v17_).f28
        nw196_[int(14)] = (d_1328_v17_).f28
        nw196_[int(15)] = (d_1328_v17_).f28
        nw196_[int(16)] = default__.fm22((d_1328_v17_).f28, d_1328_v17_.f29, globalState)
        nw196_[int(17)] = ((self).f28 if True else True)
        nw196_[int(18)] = (self).f28
        nw196_[int(19)] = not((self).f28)
        nw196_[int(20)] = (self).f28
        nw196_[int(21)] = (d_1328_v17_).f28
        nw196_[int(22)] = (d_1328_v17_).f28
        nw196_[int(23)] = (self).f28
        d_1333_v22_ = nw196_
        d_1333_v22_ = (d_1328_v17_).f27
        d_1334_v23_: _dafny.Map
        d_1334_v23_ = _dafny.Map({D10_DC26(True): (self).f28})
        d_1335_v24_: C1
        nw197_ = C1()
        nw197_.ctor__(((d_1334_v23_)[D10_DC26((self).f28)] if (D10_DC26((self).f28)) in (d_1334_v23_) else (self).f28), (self).f27, (-978) > (self.f29))
        d_1335_v24_ = nw197_
        hi6_ = self.f29
        for d_1336_i0_ in range(self.f29, hi6_):
            r0 = (default__.safeModuloInt(894, self.f29)) - (default__.fm36(globalState))
            d_1337_v25_: _dafny.Array
            nw198_ = _dafny.Array(_dafny.Map({}), 12)
            d_1337_v25_ = nw198_
            source21_ = D11_DC27(d_1337_v25_)
            if source21_.is_DC28:
                d_1312_v3_ = d_1312_v3_
                d_1312_v3_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdswexrbt"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "am"))) + (d_1312_v3_))
                d_1338_v26_: _dafny.Set
                d_1338_v26_ = _dafny.Set({(d_1335_v24_).f39, False, (d_1335_v24_).f39})
                d_1339_v27_: D8
                d_1339_v27_ = D8_DC18(len(d_1338_v26_))
                d_1339_v27_ = d_1339_v27_
                d_1340_v28_: _dafny.MultiSet
                d_1340_v28_ = _dafny.MultiSet([(self).f28])
                d_1341_v29_: _dafny.MultiSet
                d_1341_v29_ = _dafny.MultiSet([(d_1340_v28_) - (d_1340_v28_), (d_1340_v28_ if (d_1335_v24_).f39 else d_1340_v28_), d_1340_v28_])
                rhs225_ = (0) - (d_1336_i0_)
                rhs226_ = (d_1341_v29_).cardinality
                lhs195_ = globalState
                lhs196_ = globalState
                lhs195_.f5 = rhs225_
                lhs196_.f13 = rhs226_
            elif True:
                d_1342___mcc_h1_ = source21_.cf42
                d_1343_cf42_ = d_1342___mcc_h1_
                (globalState).f25 = _dafny.CodePoint('m')
                d_1344_v30_: _dafny.MultiSet
                d_1344_v30_ = _dafny.MultiSet([self.f29])
                d_1345_v31_: _dafny.MultiSet
                d_1345_v31_ = _dafny.MultiSet([True])
                d_1346_v32_: _dafny.Map
                d_1346_v32_ = _dafny.Map({self.f29: d_1345_v31_})
                d_1347_v33_: _dafny.Seq
                d_1347_v33_ = _dafny.SeqWithoutIsStrInference([(d_1345_v31_).set((d_1335_v24_).f39, default__.abs(self.f29)), ((d_1346_v32_)[d_1336_i0_] if (d_1336_i0_) in (d_1346_v32_) else d_1345_v31_), d_1345_v31_])
                (globalState).f6 = (((d_1344_v30_)[self.f29] if (self.f29) in (d_1344_v30_) else ((d_1347_v33_)[default__.safeIndex(d_1336_i0_, len(d_1347_v33_))]).cardinality)) <= (d_1336_i0_)
                d_1348_v34_: _dafny.Array
                def lambda48_(d_1349_v3_):
                    def lambda49_(d_1350_i1_):
                        return _dafny.Map({len(d_1349_v3_): (self).f28})

                    return lambda49_

                init29_ = lambda48_(d_1312_v3_)
                nw199_ = _dafny.Array(None, 13)
                for i0_29_ in range(nw199_.length(0)):
                    nw199_[i0_29_] = init29_(i0_29_)
                d_1348_v34_ = nw199_
                nw200_ = _dafny.Array(_dafny.Map({}), 9)
                d_1348_v34_ = nw200_
                d_1351_v35_: _dafny.Array
                nw201_ = _dafny.Array(int(0), 15)
                d_1351_v35_ = nw201_
                d_1352_v36_: _dafny.Seq
                d_1352_v36_ = _dafny.SeqWithoutIsStrInference([d_1351_v35_, d_1351_v35_, d_1351_v35_])
                d_1353_v37_: _dafny.MultiSet
                d_1353_v37_ = _dafny.MultiSet([d_1351_v35_, d_1351_v35_, (d_1352_v36_)[default__.safeIndex(self.f29, len(d_1352_v36_))], d_1351_v35_, d_1351_v35_])
                (globalState).f17 = (d_1353_v37_).ispropersubset((d_1353_v37_).set(d_1351_v35_, default__.abs(self.f29)))
            d_1354_v38_: _dafny.Map
            d_1354_v38_ = _dafny.Map({default__.fm50((self).f28, self.f29, globalState): (0) - (d_1336_i0_)})
            d_1354_v38_ = (d_1354_v38_).set(d_1310_v1_, 370)
            (globalState).f13 = len(d_1311_v2_)
        r0 = (0) - (self.f29)
        r1 = (self.f29) * (self.f29)
        return r0, r1


class C8(T0):
    def  __init__(self):
        self._f27: _dafny.Array = _dafny.Array(None, 0)
        self._f28: bool = False
        self._f32: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f32, f27, f28):
        (self)._f32 = f32
        (self)._f27 = f27
        (self)._f28 = f28

    def fm3(self, p0, p1, p2, p3, globalState):
        return ((D1_DC5(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-103 for d_1355_i0_ in range(default__.abs(729))])]))).cf11) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([323]), _dafny.SeqWithoutIsStrInference([-823, (_dafny.MultiSet([164])).cardinality]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f28]))]), (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([-361])))]))]))

    def m3(self, p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_1356_v0_: T1
        nw202_ = C0()
        nw202_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yovobx")), 139, (self).f27, (self).f28)
        d_1356_v0_ = nw202_
        d_1357_v1_: _dafny.Map
        d_1357_v1_ = _dafny.Map({d_1356_v0_: d_1356_v0_.f29})
        d_1358_v2_: _dafny.Map
        d_1358_v2_ = _dafny.Map({d_1357_v1_: (d_1356_v0_).f28})
        d_1359_v3_: D0
        d_1359_v3_ = D0_DC0(d_1358_v2_, (d_1356_v0_).f28, 130, (self).f32, (d_1356_v0_).f28)
        d_1360_v4_: _dafny.Map
        d_1360_v4_ = _dafny.Map({(self).f28: (d_1359_v3_).cf1})
        d_1361_v5_: D1
        d_1361_v5_ = D1_DC6(d_1356_v0_.f29)
        d_1362_v6_: _dafny.Set
        d_1362_v6_ = _dafny.Set({default__.fm4(_dafny.MultiSet([d_1361_v5_, D1_DC6(p0), d_1361_v5_]), globalState)})
        d_1363_v7_: _dafny.Seq
        d_1363_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "upehw"))
        d_1364_v8_: D0
        d_1364_v8_ = D0_DC2(d_1362_v6_, (d_1363_v7_ if (d_1356_v0_).f28 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojwsghveo"))))
        d_1365_v9_: _dafny.Seq
        d_1365_v9_ = _dafny.SeqWithoutIsStrInference([(d_1356_v0_).f28, (self).f28, (d_1356_v0_).f28])
        rhs227_ = d_1360_v4_
        rhs228_ = ((self).f28 if True else (d_1365_v9_)[default__.safeIndex(p0, len(d_1365_v9_))])
        rhs229_ = d_1364_v8_
        d_1360_v4_ = rhs227_
        r1 = rhs228_
        d_1364_v8_ = rhs229_
        d_1360_v4_ = (d_1360_v4_).set((self).f28, (d_1356_v0_).f28)
        d_1366_i0_: int
        d_1366_i0_ = 0
        with _dafny.label("5"):
            while ((d_1356_v0_).f28) or ((self).f28):
                with _dafny.c_label("5"):
                    if (d_1366_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_1366_i0_ = (d_1366_i0_) + (1)
                    d_1367_v10_: _dafny.Array
                    def lambda50_(d_1368_i1_):
                        return (d_1368_i1_) - (124)

                    init30_ = lambda50_
                    nw203_ = _dafny.Array(None, 14)
                    for i0_30_ in range(nw203_.length(0)):
                        nw203_[i0_30_] = init30_(i0_30_)
                    d_1367_v10_ = nw203_
                    d_1367_v10_ = d_1367_v10_
                    (d_1356_v0_).f29 = d_1356_v0_.f29
                    d_1369_v11_: C4
                    nw204_ = C4()
                    nw204_.ctor__()
                    d_1369_v11_ = nw204_
                    d_1369_v11_ = d_1369_v11_
                    d_1370_v12_: _dafny.Seq
                    d_1370_v12_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_1371_v13_: _dafny.Seq
                    d_1371_v13_ = d_1370_v12_
                    rhs230_ = d_1362_v6_
                    rhs231_ = d_1371_v13_
                    rhs232_ = d_1356_v0_.f29
                    rhs233_ = (self).f28
                    lhs197_ = globalState
                    d_1362_v6_ = rhs230_
                    d_1371_v13_ = rhs231_
                    lhs197_.f13 = rhs232_
                    r1 = rhs233_
                    pass
            pass
        d_1372_v14_: _dafny.MultiSet
        d_1372_v14_ = _dafny.MultiSet([d_1356_v0_.f29])
        d_1373_v16_: _dafny.Map
        d_1373_v16_ = _dafny.Map({(d_1356_v0_).f28: p0})
        d_1374_v17_: _dafny.Map
        d_1374_v17_ = _dafny.Map({d_1373_v16_: (d_1356_v0_).f28})
        d_1375_v18_: _dafny.Set
        def iife106_():
            coll64_ = _dafny.Map()
            compr_64_: _dafny.Map
            for compr_64_ in (d_1374_v17_).keys.Elements:
                d_1376_v15_: _dafny.Map = compr_64_
                if (d_1376_v15_) in (d_1374_v17_):
                    coll64_[d_1376_v15_] = d_1373_v16_
            return _dafny.Map(coll64_)
        d_1375_v18_ = _dafny.Set({default__.fm16(d_1360_v4_, globalState), ((d_1372_v14_).cardinality) <= (len(iife106_()
        ))})
        (globalState).f13 = (0) - (len(d_1375_v18_))
        rhs234_ = (not(not((self).f28)) if (d_1356_v0_).f28 else default__.fm40(globalState))
        rhs235_ = (self).f28
        rhs236_ = d_1356_v0_.f29
        lhs198_ = globalState
        lhs199_ = globalState
        r1 = rhs234_
        lhs198_.f17 = rhs235_
        lhs199_.f21 = rhs236_
        d_1377_i2_: int
        d_1377_i2_ = 0
        with _dafny.label("6"):
            while (self).f28:
                with _dafny.c_label("6"):
                    if (d_1377_i2_) >= (100):
                        raise _dafny.Break("6")
                    d_1377_i2_ = (d_1377_i2_) + (1)
                    d_1378_v19_: C1
                    nw205_ = C1()
                    nw205_.ctor__((self).f28, (d_1356_v0_).f27, (self).f28)
                    d_1378_v19_ = nw205_
                    d_1375_v18_ = d_1375_v18_
                    d_1379_v20_: _dafny.MultiSet
                    d_1379_v20_ = _dafny.MultiSet([d_1361_v5_])
                    (d_1356_v0_).f29 = default__.fm4(d_1379_v20_, globalState)
                    if (d_1378_v19_).f39:
                        d_1380_v21_: _dafny.Map
                        d_1380_v21_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdotqlbut"))): (d_1356_v0_).f27})
                        d_1381_v22_: _dafny.Array
                        d_1381_v22_ = (d_1356_v0_).f27
                        d_1382_v23_: _dafny.Array
                        nw206_ = _dafny.Array(None, 26)
                        nw206_[int(0)] = (d_1356_v0_).f27
                        nw206_[int(1)] = (d_1356_v0_).f27
                        nw206_[int(2)] = (d_1356_v0_).f27
                        nw206_[int(3)] = (d_1356_v0_).f27
                        nw206_[int(4)] = (d_1356_v0_).f27
                        nw206_[int(5)] = (d_1356_v0_).f27
                        nw206_[int(6)] = (d_1356_v0_).f27
                        nw206_[int(7)] = (d_1356_v0_).f27
                        nw206_[int(8)] = (self).f27
                        nw206_[int(9)] = (d_1356_v0_).f27
                        nw206_[int(10)] = (d_1356_v0_).f27
                        nw206_[int(11)] = (d_1356_v0_).f27
                        nw206_[int(12)] = (self).f27
                        nw206_[int(13)] = (d_1356_v0_).f27
                        nw206_[int(14)] = (d_1356_v0_).f27
                        nw206_[int(15)] = (d_1356_v0_).f27
                        nw206_[int(16)] = (d_1356_v0_).f27
                        nw206_[int(17)] = (d_1356_v0_).f27
                        nw206_[int(18)] = (self).f27
                        nw206_[int(19)] = (d_1356_v0_).f27
                        nw206_[int(20)] = ((d_1380_v21_)[d_1356_v0_.f29] if (d_1356_v0_.f29) in (d_1380_v21_) else (d_1356_v0_).f27)
                        nw206_[int(21)] = (d_1356_v0_).f27
                        nw206_[int(22)] = (self).f27
                        nw206_[int(23)] = (d_1356_v0_).f27
                        nw206_[int(24)] = (d_1381_v22_)
                        nw206_[int(25)] = (self).f27
                        d_1382_v23_ = nw206_
                        index203_ = default__.safeIndex(52, (d_1382_v23_).length(0))
                        (d_1382_v23_)[index203_] = (d_1356_v0_).f27
                        d_1383_v24_: C4
                        nw207_ = C4()
                        nw207_.ctor__()
                        d_1383_v24_ = nw207_
                        index204_ = default__.safeIndex(52, (d_1382_v23_).length(0))
                        rhs237_ = (d_1378_v19_).f39
                        rhs238_ = (d_1356_v0_).f28
                        rhs239_ = ((d_1356_v0_).f27 if ((d_1378_v19_).f39 if (d_1356_v0_).f28 else (d_1356_v0_).f28) else (d_1356_v0_).f27)
                        rhs240_ = d_1383_v24_
                        lhs200_ = globalState
                        lhs201_ = d_1382_v23_
                        lhs202_ = default__.safeIndex(52, (d_1382_v23_).length(0))
                        r1 = rhs237_
                        lhs200_.f6 = rhs238_
                        lhs201_[lhs202_] = rhs239_
                        d_1383_v24_ = rhs240_
                        (globalState).f17 = ((d_1365_v9_)[default__.safeIndex(p0, len(d_1365_v9_))] if (self).f28 else ((d_1360_v4_)[(d_1356_v0_).f28] if ((d_1356_v0_).f28) in (d_1360_v4_) else not((d_1378_v19_).f39)))
                        d_1363_v7_ = default__.fm18(((d_1378_v19_).f39 if (d_1378_v19_).f39 else (d_1378_v19_).f39), (0) - (d_1356_v0_.f29), (self).f32, globalState)
                        d_1384_v25_: _dafny.Array
                        nw208_ = _dafny.Array(int(0), 28)
                        d_1384_v25_ = nw208_
                        d_1384_v25_ = d_1384_v25_
                        r1 = (d_1356_v0_).f28
                    elif True:
                        d_1385_v26_: _dafny.Array
                        nw209_ = _dafny.Array(int(0), 20)
                        d_1385_v26_ = nw209_
                        d_1386_v27_: _dafny.Seq
                        d_1386_v27_ = _dafny.SeqWithoutIsStrInference([d_1385_v26_, d_1385_v26_])
                        d_1387_v28_: D8
                        d_1387_v28_ = D8_DC18(-681)
                        d_1388_v29_: _dafny.Map
                        d_1388_v29_ = _dafny.Map({d_1387_v28_: (d_1356_v0_).f28})
                        d_1389_v30_: C5
                        nw210_ = C5()
                        nw210_.ctor__(d_1386_v27_, len(d_1388_v29_), (self).f27, (self).f28)
                        d_1389_v30_ = nw210_
                        d_1372_v14_ = (d_1372_v14_) | ((d_1372_v14_) | (_dafny.MultiSet([p0])))
                        (globalState).f5 = default__.safeModuloInt(d_1356_v0_.f29, (d_1389_v30_).fm7((self).f32, globalState))
                        (globalState).f15 = (d_1356_v0_).f28
                        d_1390_v31_: _dafny.Array
                        nw211_ = _dafny.Array(_dafny.CodePoint('D'), 28)
                        d_1390_v31_ = nw211_
                        index205_ = default__.safeIndex(681, (d_1390_v31_).length(0))
                        (d_1390_v31_)[index205_] = (self).f32
                        d_1391_v32_: _dafny.Map
                        d_1391_v32_ = _dafny.Map({d_1356_v0_.f29: (d_1356_v0_).f28})
                        d_1392_v33_: D3
                        d_1392_v33_ = D3_DC10((d_1378_v19_).f39, (self).f28)
                        d_1393_v34_: D5
                        d_1393_v34_ = D5_DC13((d_1356_v0_).f28, d_1385_v26_, (D7_DC16((d_1378_v19_).f39, d_1391_v32_, (self).f32, (d_1392_v33_).cf17)).cf29)
                        pat_let_tv28_ = d_1356_v0_
                        d_1394_v35_: _dafny.Array
                        nw212_ = _dafny.Array(None, 2)
                        def iife107_(_pat_let21_0):
                            def iife108_(d_1395_dt__update__tmp_h0_):
                                def iife109_(_pat_let22_0):
                                    def iife110_(d_1396_dt__update_hcf21_h0_):
                                        return D5_DC13(d_1396_dt__update_hcf21_h0_, (d_1395_dt__update__tmp_h0_).cf22, (d_1395_dt__update__tmp_h0_).cf23)
                                    return iife110_(_pat_let22_0)
                                return iife109_((pat_let_tv28_).f28)
                            return iife108_(_pat_let21_0)
                        nw212_[int(0)] = iife107_(d_1393_v34_)
                        nw212_[int(1)] = d_1393_v34_
                        d_1394_v35_ = nw212_
                        index206_ = default__.safeIndex(614, (d_1394_v35_).length(0))
                        (d_1394_v35_)[index206_] = d_1393_v34_
                        d_1397_v36_: _dafny.Seq
                        d_1397_v36_ = _dafny.SeqWithoutIsStrInference([d_1363_v7_, d_1363_v7_, d_1363_v7_, d_1363_v7_])
                        index207_ = default__.safeIndex(681, (d_1390_v31_).length(0))
                        index208_ = default__.safeIndex(614, (d_1394_v35_).length(0))
                        rhs241_ = (self).f32
                        rhs242_ = d_1393_v34_
                        rhs243_ = ((d_1397_v36_)[default__.safeIndex(len(d_1391_v32_), len(d_1397_v36_))]) + (d_1363_v7_)
                        rhs244_ = (d_1393_v34_).cf22
                        rhs245_ = (d_1375_v18_).isdisjoint(d_1375_v18_)
                        lhs203_ = d_1390_v31_
                        lhs204_ = default__.safeIndex(681, (d_1390_v31_).length(0))
                        lhs205_ = d_1394_v35_
                        lhs206_ = default__.safeIndex(614, (d_1394_v35_).length(0))
                        lhs207_ = globalState
                        lhs203_[lhs204_] = rhs241_
                        lhs205_[lhs206_] = rhs242_
                        d_1363_v7_ = rhs243_
                        d_1385_v26_ = rhs244_
                        lhs207_.f15 = rhs245_
                    pass
            pass
        d_1398_v37_: _dafny.Array
        nw213_ = _dafny.Array(int(0), 29)
        d_1398_v37_ = nw213_
        d_1399_v38_: _dafny.Seq
        d_1399_v38_ = _dafny.SeqWithoutIsStrInference([d_1398_v37_])
        r0 = d_1399_v38_
        r1 = (225) == (default__.safeDivisionInt(len(d_1363_v7_), (0) - (p0)))
        return r0, r1

    def m4(self, p0, globalState):
        d_1400_v0_: _dafny.Set
        d_1400_v0_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krl")))})
        d_1401_v1_: D0
        d_1401_v1_ = D0_DC2(d_1400_v0_, default__.fm32(p0, (self).f28, globalState))
        d_1402_v2_: _dafny.Map
        d_1402_v2_ = _dafny.Map({(self).f28: (self).f32})
        d_1403_v3_: _dafny.Array
        nw214_ = _dafny.Array(None, 6)
        nw214_[int(0)] = len((d_1401_v1_).cf6)
        nw214_[int(1)] = (0) - (p0)
        nw214_[int(2)] = p0
        nw214_[int(3)] = len(d_1402_v2_)
        nw214_[int(4)] = len(_dafny.SeqWithoutIsStrInference([(self).f32 for d_1404_i0_ in range(default__.abs(388))]))
        nw214_[int(5)] = p0
        d_1403_v3_ = nw214_
        d_1405_v4_: _dafny.Seq
        d_1405_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wuevn"))
        d_1406_v5_: _dafny.Map
        d_1406_v5_ = _dafny.Map({d_1405_v4_: d_1403_v3_})
        d_1407_v6_: _dafny.Seq
        d_1407_v6_ = _dafny.SeqWithoutIsStrInference([d_1403_v3_, d_1403_v3_, d_1403_v3_, d_1403_v3_, d_1403_v3_])
        d_1408_v7_: _dafny.Array
        nw215_ = _dafny.Array(None, 16)
        nw215_[int(0)] = (d_1403_v3_ if not((self).f28) else d_1403_v3_)
        nw215_[int(1)] = d_1403_v3_
        nw215_[int(2)] = d_1403_v3_
        nw215_[int(3)] = ((d_1406_v5_)[((d_1405_v4_).set(default__.safeIndex(p0, len(d_1405_v4_)), default__.fm42(p0, globalState))).set(default__.safeIndex(p0, len((d_1405_v4_).set(default__.safeIndex(p0, len(d_1405_v4_)), default__.fm42(p0, globalState)))), (self).f32)] if (((d_1405_v4_).set(default__.safeIndex(p0, len(d_1405_v4_)), default__.fm42(p0, globalState))).set(default__.safeIndex(p0, len((d_1405_v4_).set(default__.safeIndex(p0, len(d_1405_v4_)), default__.fm42(p0, globalState)))), (self).f32)) in (d_1406_v5_) else d_1403_v3_)
        nw215_[int(4)] = (d_1407_v6_)[default__.safeIndex(p0, len(d_1407_v6_))]
        nw215_[int(5)] = (d_1403_v3_ if (self).f28 else d_1403_v3_)
        nw215_[int(6)] = d_1403_v3_
        nw215_[int(7)] = d_1403_v3_
        nw215_[int(8)] = d_1403_v3_
        nw215_[int(9)] = d_1403_v3_
        nw215_[int(10)] = d_1403_v3_
        nw215_[int(11)] = d_1403_v3_
        nw215_[int(12)] = d_1403_v3_
        nw215_[int(13)] = d_1403_v3_
        nw215_[int(14)] = d_1403_v3_
        nw215_[int(15)] = d_1403_v3_
        d_1408_v7_ = nw215_
        index209_ = default__.safeIndex(80, (d_1408_v7_).length(0))
        (d_1408_v7_)[index209_] = d_1403_v3_
        index210_ = default__.safeIndex(80, (d_1408_v7_).length(0))
        (d_1408_v7_)[index210_] = d_1403_v3_
        if (self).f28:
            d_1409_v8_: _dafny.Set
            d_1409_v8_ = _dafny.Set({(self).f28})
            d_1410_v9_: _dafny.Seq
            d_1410_v9_ = _dafny.SeqWithoutIsStrInference([d_1409_v8_])
            d_1411_v10_: D5
            d_1411_v10_ = D5_DC12(d_1410_v9_)
            d_1412_v11_: _dafny.MultiSet
            d_1412_v11_ = _dafny.MultiSet([False, default__.fm40(globalState)])
            d_1413_v12_: _dafny.Seq
            d_1413_v12_ = _dafny.SeqWithoutIsStrInference([d_1412_v11_])
            d_1414_v13_: _dafny.Seq
            d_1414_v13_ = d_1413_v12_
            d_1415_v14_: _dafny.Seq
            d_1415_v14_ = _dafny.SeqWithoutIsStrInference([(self).f28])
            d_1416_v15_: _dafny.Map
            d_1416_v15_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([(self).f28])})
            rhs246_ = d_1411_v10_
            rhs247_ = (self).f28
            rhs248_ = d_1414_v13_
            rhs249_ = ((((d_1416_v15_)[len(_dafny.SeqWithoutIsStrInference([(self).f32 for d_1417_i1_ in range(default__.abs(167))]))] if (len(_dafny.SeqWithoutIsStrInference([(self).f32 for d_1418_i1_ in range(default__.abs(167))]))) in (d_1416_v15_) else d_1415_v14_)) + (d_1415_v14_)) + ((d_1415_v14_) + (d_1415_v14_))
            lhs208_ = globalState
            d_1411_v10_ = rhs246_
            lhs208_.f15 = rhs247_
            d_1414_v13_ = rhs248_
            d_1415_v14_ = rhs249_
            d_1419_v16_: C0
            nw216_ = C0()
            nw216_.ctor__(_dafny.SeqWithoutIsStrInference([(self).f32 for d_1420_i2_ in range(default__.abs(714))]), default__.safeDivisionInt(p0, p0), (self).f27, (self).f28)
            d_1419_v16_ = nw216_
            (globalState).f17 = (self).f28
            (globalState).f21 = len(default__.fm51((p0 if (self).f28 else len(d_1416_v15_)), globalState))
            d_1421_v17_: C1
            nw217_ = C1()
            nw217_.ctor__((self).f28, (self).f27, (p0) <= ((0) - ((0) - (p0))))
            d_1421_v17_ = nw217_
        elif True:
            d_1422_v18_: C4
            nw218_ = C4()
            nw218_.ctor__()
            d_1422_v18_ = nw218_
            arr2_ = (d_1408_v7_)[default__.safeIndex(80, (d_1408_v7_).length(0))]
            index211_ = default__.safeIndex(271, ((d_1408_v7_)[default__.safeIndex(80, (d_1408_v7_).length(0))]).length(0))
            arr2_[index211_] = p0
            arr3_ = (d_1408_v7_)[default__.safeIndex(80, (d_1408_v7_).length(0))]
            index212_ = default__.safeIndex(271, ((d_1408_v7_)[default__.safeIndex(80, (d_1408_v7_).length(0))]).length(0))
            arr3_[index212_] = p0
            d_1423_v19_: _dafny.Seq
            d_1423_v19_ = _dafny.SeqWithoutIsStrInference([True])
            d_1424_v20_: _dafny.Map
            d_1424_v20_ = _dafny.Map({p0: d_1423_v19_})
            d_1425_v21_: _dafny.Map
            d_1425_v21_ = _dafny.Map({d_1424_v20_: -922})
            d_1425_v21_ = (d_1425_v21_).set(d_1424_v20_, ((d_1408_v7_)[default__.safeIndex(80, (d_1408_v7_).length(0))])[default__.safeIndex(271, ((d_1408_v7_)[default__.safeIndex(80, (d_1408_v7_).length(0))]).length(0))])
            d_1426_v22_: _dafny.MultiSet
            d_1426_v22_ = _dafny.MultiSet([True, not(not((self).f28))])
            d_1427_v23_: _dafny.Seq
            d_1427_v23_ = _dafny.SeqWithoutIsStrInference([d_1426_v22_, d_1426_v22_])
            d_1428_v24_: _dafny.Seq
            d_1428_v24_ = d_1427_v23_
            d_1429_v25_: _dafny.Array
            nw219_ = _dafny.Array(None, 25)
            nw219_[int(0)] = d_1427_v23_
            nw219_[int(1)] = d_1428_v24_
            nw219_[int(2)] = d_1428_v24_
            nw219_[int(3)] = d_1428_v24_
            nw219_[int(4)] = d_1428_v24_
            nw219_[int(5)] = d_1428_v24_
            nw219_[int(6)] = d_1428_v24_
            nw219_[int(7)] = (d_1428_v24_ if not((self).f28) else d_1428_v24_)
            nw219_[int(8)] = d_1428_v24_
            nw219_[int(9)] = d_1427_v23_
            nw219_[int(10)] = (d_1428_v24_ if True else d_1428_v24_)
            nw219_[int(11)] = d_1428_v24_
            nw219_[int(12)] = d_1428_v24_
            nw219_[int(13)] = default__.fm52(p0, p0, ((d_1408_v7_)[default__.safeIndex(80, (d_1408_v7_).length(0))])[default__.safeIndex(271, ((d_1408_v7_)[default__.safeIndex(80, (d_1408_v7_).length(0))]).length(0))], globalState)
            nw219_[int(14)] = d_1428_v24_
            nw219_[int(15)] = d_1428_v24_
            nw219_[int(16)] = (d_1428_v24_ if (self).f28 else d_1427_v23_)
            nw219_[int(17)] = d_1428_v24_
            nw219_[int(18)] = d_1428_v24_
            nw219_[int(19)] = d_1428_v24_
            nw219_[int(20)] = d_1428_v24_
            nw219_[int(21)] = d_1428_v24_
            nw219_[int(22)] = d_1428_v24_
            nw219_[int(23)] = d_1427_v23_
            nw219_[int(24)] = d_1428_v24_
            d_1429_v25_ = nw219_
            d_1429_v25_ = d_1429_v25_
            (globalState).f4 = ((d_1408_v7_)[default__.safeIndex(80, (d_1408_v7_).length(0))])[default__.safeIndex(271, ((d_1408_v7_)[default__.safeIndex(80, (d_1408_v7_).length(0))]).length(0))]
        d_1430_v26_: _dafny.Map
        d_1430_v26_ = _dafny.Map({not((self).f28): p0})
        d_1431_v27_: _dafny.Map
        d_1431_v27_ = _dafny.Map({d_1430_v26_: _dafny.CodePoint('i')})
        d_1432_v28_: _dafny.Seq
        d_1432_v28_ = _dafny.SeqWithoutIsStrInference([(self).f28, (self).f28])
        d_1431_v27_ = (d_1431_v27_).set(_dafny.Map({(d_1432_v28_)[default__.safeIndex(p0, len(d_1432_v28_))]: p0}), (self).f32)
        hi7_ = len(_dafny.SeqWithoutIsStrInference([p0 for d_1433_i4_ in range(default__.abs(22))]))
        for d_1434_i3_ in range((p0) * (p0), hi7_):
            index213_ = default__.safeIndex(80, (d_1408_v7_).length(0))
            rhs250_ = d_1403_v3_
            rhs251_ = (d_1434_i3_ if not(True) else (p0) * (p0))
            lhs209_ = d_1408_v7_
            lhs210_ = default__.safeIndex(80, (d_1408_v7_).length(0))
            lhs211_ = globalState
            lhs209_[lhs210_] = rhs250_
            lhs211_.f13 = rhs251_
            d_1435_v29_: _dafny.Map
            d_1435_v29_ = _dafny.Map({d_1434_i3_: (self).f28})
            d_1436_v31_: _dafny.Set
            def iife111_():
                coll65_ = _dafny.Map()
                compr_65_: int
                for compr_65_ in _dafny.IntegerRange(515, 829):
                    d_1437_v30_: int = compr_65_
                    if ((515) <= (d_1437_v30_)) and ((d_1437_v30_) < (829)):
                        coll65_[(d_1437_v30_) + (p0)] = (self).f28
                return _dafny.Map(coll65_)
            d_1436_v31_ = _dafny.Set({d_1435_v29_, d_1435_v29_, (iife111_()
            ).set(d_1434_i3_, (self).f28), d_1435_v29_, d_1435_v29_})
            d_1430_v26_ = (d_1430_v26_).set((self).f28, len(d_1436_v31_))
            d_1438_v32_: D1
            d_1438_v32_ = D1_DC7((self).f32, d_1405_v4_)
            source22_ = d_1438_v32_
            if source22_.is_DC6:
                d_1439___mcc_h0_ = source22_.cf12
                d_1440_cf12_ = d_1439___mcc_h0_
                d_1441_v33_: _dafny.Map
                d_1441_v33_ = _dafny.Map({(d_1430_v26_) | (d_1430_v26_): (d_1440_cf12_) * ((0) - (d_1434_i3_))})
                d_1441_v33_ = d_1441_v33_
                d_1442_v34_: _dafny.Map
                d_1442_v34_ = _dafny.Map({(self).f32: d_1440_cf12_})
                rhs252_ = ((d_1405_v4_).set(default__.safeIndex(p0, len(d_1405_v4_)), ((d_1438_v32_).cf13 if (self).f28 else (self).f32))).set(default__.safeIndex(d_1440_cf12_, len((d_1405_v4_).set(default__.safeIndex(p0, len(d_1405_v4_)), ((d_1438_v32_).cf13 if (self).f28 else (self).f32)))), (self).f32)
                rhs253_ = not((d_1442_v34_) == (d_1442_v34_))
                lhs212_ = globalState
                d_1405_v4_ = rhs252_
                lhs212_.f15 = rhs253_
                d_1443_v35_: _dafny.Map
                d_1444_v36_: int
                out27_: _dafny.Map
                out28_: int
                out27_, out28_ = default__.m0(globalState)
                d_1443_v35_ = out27_
                d_1444_v36_ = out28_
                d_1445_v37_: C4
                nw220_ = C4()
                nw220_.ctor__()
                d_1445_v37_ = nw220_
                d_1445_v37_ = d_1445_v37_
            elif source22_.is_DC7:
                d_1446___mcc_h1_ = source22_.cf13
                d_1447___mcc_h2_ = source22_.cf14
                d_1448_cf14_ = d_1447___mcc_h2_
                d_1449_cf13_ = d_1446___mcc_h1_
                d_1450_v38_: _dafny.Map
                d_1450_v38_ = _dafny.Map({p0: p0})
                d_1451_v39_: C5
                nw221_ = C5()
                nw221_.ctor__(d_1407_v6_, len(d_1450_v38_), (self).f27, (self).f28)
                d_1451_v39_ = nw221_
                d_1452_v40_: _dafny.MultiSet
                d_1452_v40_ = _dafny.MultiSet([d_1451_v39_, d_1451_v39_])
                d_1453_v41_: _dafny.Seq
                d_1453_v41_ = _dafny.SeqWithoutIsStrInference([d_1448_cf14_, d_1448_cf14_])
                rhs254_ = not ((len(d_1405_v4_)) > (p0)) or ((d_1452_v40_) == (d_1452_v40_))
                rhs255_ = (default__.safeDivisionInt(796, len(d_1453_v41_))) + (p0)
                rhs256_ = d_1434_i3_
                rhs257_ = (0) - (d_1434_i3_)
                lhs213_ = globalState
                lhs214_ = globalState
                lhs215_ = globalState
                lhs216_ = globalState
                lhs213_.f15 = rhs254_
                lhs214_.f5 = rhs255_
                lhs215_.f5 = rhs256_
                lhs216_.f4 = rhs257_
                d_1454_v42_: _dafny.MultiSet
                d_1454_v42_ = _dafny.MultiSet([d_1405_v4_])
                d_1455_v43_: _dafny.Map
                d_1455_v43_ = _dafny.Map({len(default__.fm30((self).f28, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcwjncj")), (self).f28, globalState)): d_1454_v42_})
                d_1456_v44_: _dafny.Seq
                d_1456_v44_ = _dafny.SeqWithoutIsStrInference([d_1454_v42_])
                d_1455_v43_ = (d_1455_v43_).set(len(_dafny.SeqWithoutIsStrInference([(self).f32 for d_1457_i5_ in range(default__.abs(651))])), (d_1456_v44_)[default__.safeIndex(d_1434_i3_, len(d_1456_v44_))])
                d_1458_v45_: _dafny.Seq
                d_1459_v46_: bool
                out29_: _dafny.Seq
                out30_: bool
                out29_, out30_ = (self).m3((p0) * ((0) - (p0)), globalState)
                d_1458_v45_ = out29_
                d_1459_v46_ = out30_
                d_1460_v47_: C1
                nw222_ = C1()
                nw222_.ctor__((self).f28, (self).f27, (998) >= (65))
                d_1460_v47_ = nw222_
            elif True:
                d_1461___mcc_h3_ = source22_.cf11
                d_1462_cf11_ = d_1461___mcc_h3_
                (globalState).f15 = not((self).f28)
                (globalState).f6 = (self).f28
                index214_ = default__.safeIndex(608, (d_1403_v3_).length(0))
                (d_1403_v3_)[index214_] = p0
                index215_ = default__.safeIndex(608, (d_1403_v3_).length(0))
                (d_1403_v3_)[index215_] = ((d_1434_i3_) + (d_1434_i3_)) - (d_1434_i3_)
                d_1463_v48_: C4
                nw223_ = C4()
                nw223_.ctor__()
                d_1463_v48_ = nw223_
            d_1464_v49_: D7
            d_1464_v49_ = D7_DC16((self).f28, d_1435_v29_, (self).f32, (self).f28)
            d_1465_v50_: _dafny.Map
            d_1465_v50_ = _dafny.Map({(self).f28: False})
            if default__.fm16((d_1465_v50_ if (d_1464_v49_).cf26 else d_1465_v50_), globalState):
                d_1466_v52_: _dafny.Map
                d_1466_v52_ = _dafny.Map({p0: (self).f32})
                d_1467_v53_: _dafny.MultiSet
                d_1467_v53_ = _dafny.MultiSet([((d_1466_v52_)[d_1434_i3_] if (d_1434_i3_) in (d_1466_v52_) else (self).f32)])
                d_1468_v54_: _dafny.MultiSet
                def iife112_():
                    coll66_ = _dafny.Map()
                    compr_66_: str
                    for compr_66_ in ((d_1467_v53_).set((self).f32, default__.abs(d_1434_i3_))).Elements:
                        d_1469_v51_: str = compr_66_
                        if (d_1469_v51_) in ((d_1467_v53_).set((self).f32, default__.abs(d_1434_i3_))):
                            coll66_[d_1469_v51_] = d_1434_i3_
                    return _dafny.Map(coll66_)
                d_1468_v54_ = _dafny.MultiSet([len(iife112_()
                ), d_1434_i3_, d_1434_i3_, d_1434_i3_])
                d_1470_v55_: _dafny.Map
                d_1470_v55_ = _dafny.Map({d_1402_v2_: (d_1468_v54_).cardinality})
                (globalState).f5 = ((d_1470_v55_)[d_1402_v2_] if (d_1402_v2_) in (d_1470_v55_) else (len(d_1432_v28_)) - (d_1434_i3_))
                d_1471_v56_: _dafny.Map
                d_1471_v56_ = _dafny.Map({p0: (0) - (d_1434_i3_)})
                (globalState).f21 = default__.safeDivisionInt(((d_1471_v56_)[d_1434_i3_] if (d_1434_i3_) in (d_1471_v56_) else p0), p0)
                d_1472_v57_: D9
                d_1472_v57_ = D9_DC23(not((self).f28))
                d_1473_v58_: D9
                d_1473_v58_ = D9_DC24(d_1472_v57_)
                d_1474_v59_: D11
                d_1474_v59_ = D11_DC28()
                d_1475_v60_: _dafny.Array
                nw224_ = _dafny.Array(_dafny.Map({}), 18)
                d_1475_v60_ = nw224_
                d_1476_v65_: _dafny.Array
                nw225_ = _dafny.Array(None, 25)
                nw225_[int(0)] = (d_1435_v29_) | (d_1435_v29_)
                nw225_[int(1)] = d_1435_v29_
                nw225_[int(2)] = d_1435_v29_
                def iife113_():
                    def iife115_():
                        coll69_ = _dafny.Set()
                        compr_69_: int
                        for compr_69_ in _dafny.IntegerRange(347, 714):
                            d_1479_v62_: int = compr_69_
                            if ((347) <= (d_1479_v62_)) and ((d_1479_v62_) < (714)):
                                coll69_ = coll69_.union(_dafny.Set([(d_1479_v62_) - (p0)]))
                        return _dafny.Set(coll69_)
                    coll67_ = _dafny.Map()
                    def iife114_():
                        coll68_ = _dafny.Set()
                        compr_68_: int
                        for compr_68_ in _dafny.IntegerRange(347, 714):
                            d_1477_v62_: int = compr_68_
                            if ((347) <= (d_1477_v62_)) and ((d_1477_v62_) < (714)):
                                coll68_ = coll68_.union(_dafny.Set([(d_1477_v62_) - (p0)]))
                        return _dafny.Set(coll68_)
                    compr_67_: int
                    for compr_67_ in (iife114_()
                    ).Elements:
                        d_1478_v61_: int = compr_67_
                        if (d_1478_v61_) in (iife115_()
                        ):
                            coll67_[(d_1478_v61_) - (d_1434_i3_)] = (self).f28
                    return _dafny.Map(coll67_)
                nw225_[int(3)] = (iife113_()
                ) | (d_1435_v29_)
                nw225_[int(4)] = d_1435_v29_
                nw225_[int(5)] = _dafny.Map({d_1434_i3_: (self).f28})
                nw225_[int(6)] = d_1435_v29_
                nw225_[int(7)] = d_1435_v29_
                nw225_[int(8)] = d_1435_v29_
                nw225_[int(9)] = d_1435_v29_
                nw225_[int(10)] = (d_1435_v29_) | (d_1435_v29_)
                nw225_[int(11)] = (d_1435_v29_) | (_dafny.Map({p0: (self).f28}))
                nw225_[int(12)] = d_1435_v29_
                nw225_[int(13)] = (_dafny.Map({d_1434_i3_: (self).f28})) | ((d_1435_v29_).set(p0, (self).f28))
                nw225_[int(14)] = default__.fm30((self).f28, d_1405_v4_, (self).f28, globalState)
                nw225_[int(15)] = d_1435_v29_
                nw225_[int(16)] = (_dafny.Map({d_1434_i3_: (self).f28})) | (d_1435_v29_)
                nw225_[int(17)] = d_1435_v29_
                nw225_[int(18)] = default__.fm30((self).f28, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtw")), (self).f28, globalState)
                nw225_[int(19)] = (d_1464_v49_).cf27
                def iife116_():
                    coll70_ = _dafny.Map()
                    compr_70_: int
                    for compr_70_ in (d_1435_v29_).keys.Elements:
                        d_1480_v63_: int = compr_70_
                        if (d_1480_v63_) in (d_1435_v29_):
                            coll70_[(d_1480_v63_) * (len(_dafny.Map({(self).f28: d_1434_i3_})))] = (self).f28
                    return _dafny.Map(coll70_)
                nw225_[int(20)] = (d_1435_v29_) | (iife116_()
                )
                def iife117_():
                    coll71_ = _dafny.Map()
                    compr_71_: int
                    for compr_71_ in _dafny.IntegerRange(174, 816):
                        d_1481_v64_: int = compr_71_
                        if ((174) <= (d_1481_v64_)) and ((d_1481_v64_) < (816)):
                            coll71_[(d_1481_v64_) - ((D1_DC6(len(_dafny.Map({_dafny.MultiSet([(self).f28, (self).f28, (self).f28]): (self).f28})))).cf12)] = (D9_DC23(True)).cf38
                    return _dafny.Map(coll71_)
                nw225_[int(21)] = (iife117_()
                ) | (d_1435_v29_)
                nw225_[int(22)] = ((d_1435_v29_).set(p0, (self).f28)) | (d_1435_v29_)
                nw225_[int(23)] = d_1435_v29_
                nw225_[int(24)] = (d_1464_v49_).cf27
                d_1476_v65_ = nw225_
                index216_ = default__.safeIndex(626, (d_1476_v65_).length(0))
                (d_1476_v65_)[index216_] = d_1435_v29_
                d_1482_v66_: _dafny.Array
                d_1482_v66_ = (self).f27
                d_1483_v67_: _dafny.Array
                nw226_ = _dafny.Array(None, 28)
                nw226_[int(0)] = (self).f27
                nw226_[int(1)] = (self).f27
                nw226_[int(2)] = (self).f27
                nw226_[int(3)] = (self).f27
                nw226_[int(4)] = (self).f27
                nw226_[int(5)] = (self).f27
                nw226_[int(6)] = (self).f27
                nw226_[int(7)] = (d_1482_v66_)
                nw226_[int(8)] = (self).f27
                nw226_[int(9)] = (self).f27
                nw226_[int(10)] = (d_1482_v66_)
                nw226_[int(11)] = (d_1482_v66_)
                nw226_[int(12)] = (self).f27
                nw226_[int(13)] = (self).f27
                nw226_[int(14)] = (self).f27
                nw226_[int(15)] = (self).f27
                nw226_[int(16)] = (self).f27
                nw226_[int(17)] = (self).f27
                nw226_[int(18)] = (self).f27
                nw226_[int(19)] = (self).f27
                nw226_[int(20)] = (self).f27
                nw226_[int(21)] = (self).f27
                nw226_[int(22)] = (self).f27
                nw226_[int(23)] = (self).f27
                nw226_[int(24)] = (self).f27
                nw226_[int(25)] = (d_1482_v66_)
                nw226_[int(26)] = (self).f27
                nw226_[int(27)] = (self).f27
                d_1483_v67_ = nw226_
                index217_ = default__.safeIndex(626, (d_1476_v65_).length(0))
                rhs258_ = D9_DC24(D9_DC22((d_1408_v7_)[default__.safeIndex(80, (d_1408_v7_).length(0))], not(False), d_1483_v67_, d_1405_v4_))
                rhs259_ = d_1474_v59_
                rhs260_ = d_1475_v60_
                rhs261_ = (_dafny.Map({p0: (self).f28})) | ((d_1435_v29_) | (_dafny.Map({p0: (self).f28})))
                lhs217_ = d_1476_v65_
                lhs218_ = default__.safeIndex(626, (d_1476_v65_).length(0))
                d_1473_v58_ = rhs258_
                d_1474_v59_ = rhs259_
                d_1475_v60_ = rhs260_
                lhs217_[lhs218_] = rhs261_
                (globalState).f25 = default__.fm42(p0, globalState)
                d_1484_v68_: D0
                d_1484_v68_ = D0_DC3(p0, False, (self).f28)
                d_1485_v69_: _dafny.Map
                d_1485_v69_ = _dafny.Map({d_1484_v68_: p0})
                d_1486_v70_: _dafny.MultiSet
                d_1486_v70_ = _dafny.MultiSet([(self).f28, (self).f28, (self).f28])
                d_1487_v71_: T1
                nw227_ = C3()
                nw227_.ctor__(d_1434_i3_, d_1485_v69_, (d_1486_v70_).cardinality, (self).f27, True)
                d_1487_v71_ = nw227_
                d_1488_v72_: _dafny.Set
                d_1488_v72_ = _dafny.Set({d_1487_v71_})
                (globalState).f15 = ((d_1488_v72_) | (_dafny.Set({d_1487_v71_}))).isdisjoint(d_1488_v72_)
            elif True:
                (globalState).f15 = not ((self).f28) or ((self).f28)
                (globalState).f15 = False
                d_1489_v73_: C1
                nw228_ = C1()
                nw228_.ctor__((self).f28, (self).f27, (self).f28)
                d_1489_v73_ = nw228_
                (globalState).f21 = default__.safeDivisionInt((0) - (p0), p0)
                d_1490_v74_: _dafny.MultiSet
                d_1490_v74_ = _dafny.MultiSet([D1_DC6(d_1434_i3_)])
                (globalState).f13 = default__.fm4(d_1490_v74_, globalState)
        d_1491_v75_: _dafny.Seq
        d_1491_v75_ = _dafny.SeqWithoutIsStrInference([d_1400_v0_])
        hi8_ = p0
        for d_1492_i6_ in range(len(d_1491_v75_), hi8_):
            if (self).f28:
                d_1493_v76_: D1
                d_1493_v76_ = D1_DC6(60)
                d_1494_v77_: _dafny.MultiSet
                d_1494_v77_ = _dafny.MultiSet([d_1493_v76_, d_1493_v76_])
                (globalState).f21 = default__.fm4(d_1494_v77_, globalState)
                d_1495_v78_: _dafny.Map
                d_1495_v78_ = _dafny.Map({d_1430_v26_: p0})
                d_1496_v79_: D8
                d_1496_v79_ = D8_DC17(d_1495_v78_)
                d_1497_v80_: _dafny.Seq
                d_1497_v80_ = _dafny.SeqWithoutIsStrInference([d_1496_v79_])
                (globalState).f21 = len(d_1497_v80_)
                d_1498_v81_: C2
                nw229_ = C2()
                nw229_.ctor__((self).f32, (self).f28, (self).f27, (self).f28)
                d_1498_v81_ = nw229_
                rhs262_ = d_1498_v81_
                rhs263_ = d_1492_i6_
                lhs219_ = globalState
                d_1498_v81_ = rhs262_
                lhs219_.f13 = rhs263_
                d_1499_v82_: C5
                nw230_ = C5()
                nw230_.ctor__(d_1407_v6_, d_1492_i6_, (self).f27, ((self).f28) or (True))
                d_1499_v82_ = nw230_
                d_1500_v83_: _dafny.Map
                d_1500_v83_ = _dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f28])): d_1498_v81_.f38})
                d_1501_v84_: _dafny.MultiSet
                d_1501_v84_ = _dafny.MultiSet([d_1498_v81_.f38])
                d_1502_v86_: _dafny.Map
                d_1502_v86_ = _dafny.Map({d_1501_v84_: p0})
                d_1503_v87_: _dafny.Map
                d_1503_v87_ = _dafny.Map({d_1492_i6_: (self).f28})
                d_1504_v88_: _dafny.Map
                d_1504_v88_ = _dafny.Map({(self).f27: -288})
                d_1505_v89_: _dafny.Array
                nw231_ = _dafny.Array(None, 24)
                nw231_[int(0)] = d_1500_v83_
                nw231_[int(1)] = d_1500_v83_
                nw231_[int(2)] = (d_1500_v83_) | (d_1500_v83_)
                nw231_[int(3)] = (d_1500_v83_) | (d_1500_v83_)
                def iife118_():
                    coll72_ = _dafny.Map()
                    compr_72_: _dafny.MultiSet
                    for compr_72_ in (d_1502_v86_).keys.Elements:
                        d_1506_v85_: _dafny.MultiSet = compr_72_
                        if (d_1506_v85_) in (d_1502_v86_):
                            coll72_[d_1506_v85_] = d_1498_v81_.f38
                    return _dafny.Map(coll72_)
                nw231_[int(4)] = ((d_1500_v83_).set(d_1501_v84_, False)) | (iife118_()
                )
                nw231_[int(5)] = d_1500_v83_
                nw231_[int(6)] = (_dafny.Map({d_1501_v84_: (D7_DC16(d_1498_v81_.f38, d_1503_v87_, (d_1498_v81_).f37, d_1498_v81_.f38)).cf29})).set(d_1501_v84_, default__.fm22(d_1498_v81_.f38, d_1492_i6_, globalState))
                nw231_[int(7)] = d_1500_v83_
                nw231_[int(8)] = (d_1500_v83_ if not(d_1498_v81_.f38) else d_1500_v83_)
                nw231_[int(9)] = (d_1500_v83_).set((d_1501_v84_).set((self).f28, default__.abs(p0)), d_1498_v81_.f38)
                nw231_[int(10)] = ((d_1500_v83_).set(_dafny.MultiSet([(self).f28]), (self).f28) if not((self).f28) else d_1500_v83_)
                nw231_[int(11)] = (_dafny.Map({d_1501_v84_: not(d_1498_v81_.f38)})) | ((d_1500_v83_))
                nw231_[int(12)] = d_1500_v83_
                nw231_[int(13)] = d_1500_v83_
                nw231_[int(14)] = (d_1500_v83_).set(d_1501_v84_, (self).f28)
                nw231_[int(15)] = d_1500_v83_
                nw231_[int(16)] = d_1500_v83_
                nw231_[int(17)] = d_1500_v83_
                nw231_[int(18)] = _dafny.Map({d_1501_v84_: not(True)})
                nw231_[int(19)] = d_1500_v83_
                nw231_[int(20)] = d_1500_v83_
                nw231_[int(21)] = default__.fm53(p0, d_1401_v1_, globalState)
                nw231_[int(22)] = d_1500_v83_
                nw231_[int(23)] = default__.fm53(len(d_1504_v88_), d_1401_v1_, globalState)
                d_1505_v89_ = nw231_
                index218_ = default__.safeIndex(978, (d_1505_v89_).length(0))
                (d_1505_v89_)[index218_] = _dafny.Map({(d_1501_v84_).set((self).f28, default__.abs(d_1492_i6_)): d_1498_v81_.f38})
                index219_ = default__.safeIndex(978, (d_1505_v89_).length(0))
                (d_1505_v89_)[index219_] = (d_1500_v83_) | ((_dafny.Map({d_1501_v84_: (self).f28})) | (d_1500_v83_))
            elif True:
                index220_ = default__.safeIndex(53, ((self).f27).length(0))
                ((self).f27)[index220_] = (self).f28
                d_1507_v90_: _dafny.Set
                d_1507_v90_ = _dafny.Set({d_1432_v28_})
                index221_ = default__.safeIndex(53, ((self).f27).length(0))
                ((self).f27)[index221_] = (d_1432_v28_) not in (d_1507_v90_)
                d_1508_v91_: D0
                d_1508_v91_ = D0_DC3(p0, ((self).f27)[default__.safeIndex(53, ((self).f27).length(0))], (self).f28)
                d_1508_v91_ = d_1508_v91_
                index222_ = default__.safeIndex(760, (d_1403_v3_).length(0))
                (d_1403_v3_)[index222_] = p0
                index223_ = default__.safeIndex(760, (d_1403_v3_).length(0))
                (d_1403_v3_)[index223_] = default__.safeModuloInt(470, (0) - ((0) - ((d_1492_i6_) * ((0) - (p0)))))
                index224_ = default__.safeIndex(760, (d_1403_v3_).length(0))
                (d_1403_v3_)[index224_] = (d_1403_v3_)[default__.safeIndex(760, (d_1403_v3_).length(0))]
                (globalState).f4 = (0) - (p0)
            (globalState).f6 = not(default__.fm40(globalState))
            d_1509_v92_: C5
            nw232_ = C5()
            nw232_.ctor__(d_1407_v6_, (0) - (len(default__.fm30((self).f28, _dafny.SeqWithoutIsStrInference([(self).f32 for d_1510_i7_ in range(default__.abs(875))]), (self).f28, globalState))), ((self).f27 if (self).f28 else (self).f27), (self).f28)
            d_1509_v92_ = nw232_
            d_1511_v93_: _dafny.Map
            d_1511_v93_ = _dafny.Map({(0) - ((0) - (p0)): p0})
            d_1512_v94_: D8
            d_1512_v94_ = D8_DC18(d_1492_i6_)
            d_1513_v95_: _dafny.Seq
            d_1513_v95_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1514_v96_: C5
            nw233_ = C5()
            nw233_.ctor__(d_1407_v6_, ((d_1511_v93_)[(d_1512_v94_).cf31] if ((d_1512_v94_).cf31) in (d_1511_v93_) else (d_1513_v95_)[default__.safeIndex(139, len(d_1513_v95_))]), (self).f27, (self).f28)
            d_1514_v96_ = nw233_
        hi9_ = p0
        for d_1515_i8_ in range(default__.safeDivisionInt(p0, p0), hi9_):
            (globalState).f4 = d_1515_i8_
            (globalState).f15 = not((self).f28)
            d_1516_v97_: D0
            d_1516_v97_ = D0_DC3(len(d_1400_v0_), (self).f28, (self).f28)
            d_1517_v98_: _dafny.Map
            d_1517_v98_ = _dafny.Map({d_1516_v97_: p0})
            d_1518_v99_: D5
            d_1518_v99_ = D5_DC13(True, d_1403_v3_, (self).f28)
            d_1519_v100_: _dafny.Map
            d_1519_v100_ = _dafny.Map({p0: (d_1518_v99_).cf23})
            d_1520_v101_: D3
            d_1520_v101_ = D3_DC10((self).f28, ((d_1519_v100_)[p0] if (p0) in (d_1519_v100_) else (self).f28))
            d_1521_v102_: C3
            nw234_ = C3()
            nw234_.ctor__((p0 if (self).f28 else p0), (d_1517_v98_) | (d_1517_v98_), default__.safeModuloInt(p0, d_1515_i8_), (self).f27, (d_1520_v101_).cf17)
            d_1521_v102_ = nw234_
            rhs264_ = d_1405_v4_
            rhs265_ = (default__.safeModuloInt(d_1515_i8_, p0)) + (((d_1521_v102_).f34) * (d_1515_i8_))
            lhs220_ = globalState
            d_1405_v4_ = rhs264_
            lhs220_.f4 = rhs265_

    @property
    def f32(self):
        return self._f32

class C9(T1, T0):
    def  __init__(self):
        self._f29: int = int(0)
        self._f27: _dafny.Array = _dafny.Array(None, 0)
        self._f28: bool = False
        self._f30: bool = False
        self._f31: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f30, f31, f29, f27, f28):
        (self)._f30 = f30
        (self)._f31 = f31
        (self).f29 = f29
        (self)._f27 = f27
        (self)._f28 = f28

    def m1(self, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        r3: bool = False
        d_1522_v0_: _dafny.Array
        nw235_ = _dafny.Array(int(0), 29)
        d_1522_v0_ = nw235_
        d_1523_v1_: _dafny.Map
        d_1523_v1_ = _dafny.Map({d_1522_v0_: d_1522_v0_})
        d_1523_v1_ = (d_1523_v1_).set(d_1522_v0_, d_1522_v0_)
        d_1524_v2_: _dafny.Map
        d_1524_v2_ = _dafny.Map({(self).f28: (self).f30})
        d_1525_v3_: _dafny.Set
        d_1525_v3_ = _dafny.Set({default__.fm1(globalState)})
        d_1526_v4_: _dafny.Seq
        d_1526_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hulb"))
        d_1527_v5_: D0
        d_1527_v5_ = D0_DC3(len(d_1524_v2_), (self.f29) in (d_1525_v3_), default__.fm2((d_1526_v4_)[default__.safeIndex((0) - ((0) - (self.f29)), len(d_1526_v4_))], self.f29, globalState))
        source23_ = d_1527_v5_
        if source23_.is_DC0:
            d_1528___mcc_h0_ = source23_.cf0
            d_1529___mcc_h1_ = source23_.cf1
            d_1530___mcc_h2_ = source23_.cf2
            d_1531___mcc_h3_ = source23_.cf3
            d_1532___mcc_h4_ = source23_.cf4
            d_1533_cf4_ = d_1532___mcc_h4_
            d_1534_cf3_ = d_1531___mcc_h3_
            d_1535_cf2_ = d_1530___mcc_h2_
            d_1536_cf1_ = d_1529___mcc_h1_
            d_1537_cf0_ = d_1528___mcc_h0_
            (globalState).f5 = self.f29
            d_1538_v6_: _dafny.Map
            d_1538_v6_ = _dafny.Map({d_1536_cf1_: d_1534_cf3_})
            d_1539_v7_: _dafny.MultiSet
            d_1539_v7_ = _dafny.MultiSet([d_1536_cf1_, (self).f28, (self).f28])
            (globalState).f15 = default__.fm2(((d_1538_v6_)[False] if (False) in (d_1538_v6_) else d_1534_cf3_), ((d_1539_v7_)[(self).f28] if ((self).f28) in (d_1539_v7_) else self.f29), globalState)
            (globalState).f5 = self.f29
            d_1540_v8_: T0
            nw236_ = C2()
            nw236_.ctor__(d_1534_cf3_, not(d_1533_cf4_), (self).f27, (self).f30)
            d_1540_v8_ = nw236_
            d_1541_v9_: _dafny.Map
            d_1541_v9_ = _dafny.Map({d_1540_v8_: (self).f28})
            (globalState).f13 = default__.safeModuloInt(default__.safeDivisionInt(self.f29, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efreopb")))), len(d_1541_v9_))
        elif source23_.is_DC1:
            d_1542_v10_: _dafny.Set
            d_1542_v10_ = _dafny.Set({not((self).f30), (self).f30, default__.fm40(globalState), (self).f30, default__.fm16(d_1524_v2_, globalState)})
            index225_ = default__.safeIndex(137, (d_1522_v0_).length(0))
            (d_1522_v0_)[index225_] = len(d_1542_v10_)
            d_1543_v11_: _dafny.Seq
            d_1543_v11_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(self.f29, self.f29)])
            d_1544_v12_: C0
            nw237_ = C0()
            nw237_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1545_i0_ in range(default__.abs(501))]), self.f29, (self).f27, (self).f30)
            d_1544_v12_ = nw237_
            d_1546_v13_: _dafny.Map
            d_1546_v13_ = _dafny.Map({(self).f30: d_1544_v12_})
            d_1547_v14_: _dafny.Set
            d_1547_v14_ = _dafny.Set({d_1546_v13_, d_1546_v13_})
            index226_ = default__.safeIndex(137, (d_1522_v0_).length(0))
            rhs266_ = ((self.f29) - (self.f29) if (d_1547_v14_) == (d_1547_v14_) else self.f29)
            rhs267_ = (d_1543_v11_).set(default__.safeIndex(self.f29, len(d_1543_v11_)), self.f29)
            lhs221_ = d_1522_v0_
            lhs222_ = default__.safeIndex(137, (d_1522_v0_).length(0))
            lhs221_[lhs222_] = rhs266_
            d_1543_v11_ = rhs267_
            d_1548_v15_: _dafny.Seq
            d_1548_v15_ = _dafny.SeqWithoutIsStrInference([True])
            index227_ = default__.safeIndex(137, (d_1522_v0_).length(0))
            (d_1522_v0_)[index227_] = (0) - ((len(d_1548_v15_)) - (default__.safeDivisionInt(928, self.f29)))
            d_1549_v16_: D1
            d_1549_v16_ = D1_DC6(len(d_1524_v2_))
            (globalState).f13 = default__.safeModuloInt(self.f29, default__.fm4(_dafny.MultiSet([d_1549_v16_]), globalState))
            d_1550_v17_: _dafny.Map
            d_1550_v17_ = _dafny.Map({((((self).f31)[(d_1522_v0_)[default__.safeIndex(137, (d_1522_v0_).length(0))]] if ((d_1522_v0_)[default__.safeIndex(137, (d_1522_v0_).length(0))]) in ((self).f31) else (d_1522_v0_)[default__.safeIndex(137, (d_1522_v0_).length(0))]) if (self).f30 else self.f29): default__.fm1(globalState)})
            (globalState).f13 = ((d_1550_v17_)[self.f29] if (self.f29) in (d_1550_v17_) else (0) - (self.f29))
        elif source23_.is_DC2:
            d_1551___mcc_h5_ = source23_.cf5
            d_1552___mcc_h6_ = source23_.cf6
            d_1553_cf6_ = d_1552___mcc_h6_
            d_1554_cf5_ = d_1551___mcc_h5_
            d_1555_v18_: _dafny.Array
            nw238_ = _dafny.Array(D10.default()(), 15)
            d_1555_v18_ = nw238_
            rhs268_ = ((self.f29) - (self.f29)) + ((self.f29) - (self.f29))
            rhs269_ = d_1555_v18_
            lhs223_ = globalState
            lhs223_.f13 = rhs268_
            d_1555_v18_ = rhs269_
            d_1556_v19_: str
            d_1556_v19_ = _dafny.CodePoint('v')
            d_1557_v20_: _dafny.Seq
            d_1557_v20_ = _dafny.SeqWithoutIsStrInference([(self.f29) < (self.f29), (self).f28, (self).f28, (len(d_1526_v4_)) < (self.f29), not(((d_1526_v4_).set(default__.safeIndex(self.f29, len(d_1526_v4_)), d_1556_v19_)) < (d_1526_v4_))])
            (globalState).f15 = (d_1557_v20_)[default__.safeIndex(self.f29, len(d_1557_v20_))]
            d_1558_v21_: _dafny.Map
            d_1558_v21_ = _dafny.Map({self.f29: (self).f30})
            d_1559_v22_: _dafny.Map
            d_1559_v22_ = _dafny.Map({d_1553_cf6_: (self).f27})
            d_1558_v21_ = (d_1558_v21_).set((self.f29) * (len(d_1559_v22_)), (self).f28)
            d_1560_v23_: _dafny.Array
            nw239_ = _dafny.Array(None, 21)
            nw239_[int(0)] = (self).f27
            nw239_[int(1)] = (self).f27
            nw239_[int(2)] = (self).f27
            nw239_[int(3)] = (self).f27
            nw239_[int(4)] = (self).f27
            nw239_[int(5)] = (self).f27
            nw239_[int(6)] = (self).f27
            nw239_[int(7)] = (self).f27
            nw239_[int(8)] = (self).f27
            nw239_[int(9)] = (self).f27
            nw239_[int(10)] = (self).f27
            nw239_[int(11)] = (self).f27
            nw239_[int(12)] = (self).f27
            nw239_[int(13)] = (self).f27
            nw239_[int(14)] = (self).f27
            nw239_[int(15)] = (self).f27
            nw239_[int(16)] = (self).f27
            nw239_[int(17)] = (self).f27
            nw239_[int(18)] = (self).f27
            nw239_[int(19)] = (self).f27
            nw239_[int(20)] = (self).f27
            d_1560_v23_ = nw239_
            d_1561_v24_: D9
            d_1561_v24_ = D9_DC22(d_1522_v0_, (self).f28, d_1560_v23_, d_1553_cf6_)
            d_1526_v4_ = (((d_1561_v24_).cf37) + (d_1553_cf6_)) + (d_1526_v4_)
        elif source23_.is_DC3:
            d_1562___mcc_h7_ = source23_.cf7
            d_1563___mcc_h8_ = source23_.cf8
            d_1564___mcc_h9_ = source23_.cf9
            d_1565_cf9_ = d_1564___mcc_h9_
            d_1566_cf8_ = d_1563___mcc_h8_
            d_1567_cf7_ = d_1562___mcc_h7_
            d_1568_v25_: _dafny.Array
            nw240_ = _dafny.Array(_dafny.Map({}), 24)
            d_1568_v25_ = nw240_
            d_1568_v25_ = d_1568_v25_
            (globalState).f17 = ((self).f30 if (self).f28 else d_1565_cf9_)
            (globalState).f5 = default__.fm36(globalState)
            d_1569_v26_: _dafny.Array
            def lambda51_(d_1570_i1_):
                return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1571_i2_ in range(default__.abs(-313))])

            init31_ = lambda51_
            nw241_ = _dafny.Array(None, 2)
            for i0_31_ in range(nw241_.length(0)):
                nw241_[i0_31_] = init31_(i0_31_)
            d_1569_v26_ = nw241_
            index228_ = default__.safeIndex(774, (d_1569_v26_).length(0))
            (d_1569_v26_)[index228_] = d_1526_v4_
            d_1572_v27_: _dafny.MultiSet
            d_1572_v27_ = _dafny.MultiSet([d_1522_v0_])
            d_1573_v28_: str
            d_1573_v28_ = _dafny.CodePoint('a')
            index229_ = default__.safeIndex(774, (d_1569_v26_).length(0))
            rhs270_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fs"))) + (d_1526_v4_)) + ((d_1526_v4_) + (default__.fm18(d_1565_cf9_, self.f29, d_1573_v28_, globalState)))
            rhs271_ = d_1567_cf7_
            rhs272_ = (self).f28
            rhs273_ = (d_1572_v27_) - (d_1572_v27_)
            lhs224_ = d_1569_v26_
            lhs225_ = default__.safeIndex(774, (d_1569_v26_).length(0))
            lhs226_ = globalState
            lhs224_[lhs225_] = rhs270_
            lhs226_.f21 = rhs271_
            r2 = rhs272_
            d_1572_v27_ = rhs273_
        elif True:
            d_1574___mcc_h10_ = source23_.cf10
            d_1575_cf10_ = d_1574___mcc_h10_
            d_1576_v29_: _dafny.Seq
            d_1576_v29_ = _dafny.SeqWithoutIsStrInference([(self).f30, (self).f30, (self).f30])
            d_1577_v30_: _dafny.Seq
            d_1577_v30_ = _dafny.SeqWithoutIsStrInference([d_1522_v0_])
            d_1578_v31_: C5
            nw242_ = C5()
            nw242_.ctor__(d_1577_v30_, self.f29, (self).f27, (self).f30)
            d_1578_v31_ = nw242_
            d_1579_v32_: _dafny.Map
            d_1579_v32_ = _dafny.Map({d_1578_v31_: (_dafny.SeqWithoutIsStrInference([(self).f30, (self).f28]) if (self).f28 else d_1576_v29_)})
            d_1576_v29_ = ((d_1579_v32_)[d_1578_v31_] if (d_1578_v31_) in (d_1579_v32_) else d_1576_v29_)
            (globalState).f15 = (self).f28
            d_1580_v33_: T1
            nw243_ = C7()
            nw243_.ctor__((self).f27, (self).f28, self.f29)
            d_1580_v33_ = nw243_
            d_1581_v34_: _dafny.Set
            d_1581_v34_ = _dafny.Set({d_1526_v4_, d_1526_v4_})
            nw244_ = C6()
            nw244_.ctor__((0) - (d_1580_v33_.f29), (self).f27, (d_1581_v34_).issubset(d_1581_v34_))
            d_1580_v33_ = nw244_
            (globalState).f13 = d_1580_v33_.f29
        hi10_ = self.f29
        for d_1582_i3_ in range(len((d_1526_v4_) + (d_1526_v4_)), hi10_):
            d_1583_v35_: _dafny.Map
            d_1583_v35_ = _dafny.Map({d_1582_i3_: self.f29})
            d_1584_v36_: C0
            nw245_ = C0()
            nw245_.ctor__(d_1526_v4_, len(d_1583_v35_), (self).f27, not((self).f28))
            d_1584_v36_ = nw245_
            d_1585_v37_: _dafny.Array
            def lambda52_(d_1586_v3_):
                def lambda53_(d_1587_i4_):
                    return d_1586_v3_

                return lambda53_

            init32_ = lambda52_(d_1525_v3_)
            nw246_ = _dafny.Array(None, 13)
            for i0_32_ in range(nw246_.length(0)):
                nw246_[i0_32_] = init32_(i0_32_)
            d_1585_v37_ = nw246_
            index230_ = default__.safeIndex(329, (d_1585_v37_).length(0))
            (d_1585_v37_)[index230_] = (d_1525_v3_).intersection(d_1525_v3_)
            d_1588_v38_: _dafny.Seq
            d_1588_v38_ = _dafny.SeqWithoutIsStrInference([self.f29])
            d_1589_v39_: _dafny.Seq
            d_1589_v39_ = _dafny.SeqWithoutIsStrInference([d_1588_v38_])
            d_1590_v40_: _dafny.Map
            d_1590_v40_ = _dafny.Map({D1_DC5(d_1589_v39_): (self).f28})
            d_1591_v41_: _dafny.Seq
            d_1591_v41_ = _dafny.SeqWithoutIsStrInference([(self).f30, (self).f30, (self).f30])
            d_1592_v42_: _dafny.Seq
            d_1592_v42_ = _dafny.SeqWithoutIsStrInference([len(d_1590_v40_), d_1582_i3_, len(d_1591_v41_)])
            index231_ = default__.safeIndex(329, (d_1585_v37_).length(0))
            rhs274_ = d_1584_v36_
            rhs275_ = d_1525_v3_
            rhs276_ = ((d_1592_v42_)[default__.safeIndex(d_1582_i3_, len(d_1592_v42_))]) - (self.f29)
            lhs227_ = d_1585_v37_
            lhs228_ = default__.safeIndex(329, (d_1585_v37_).length(0))
            lhs229_ = globalState
            d_1584_v36_ = rhs274_
            lhs227_[lhs228_] = rhs275_
            lhs229_.f21 = rhs276_
            d_1592_v42_ = d_1588_v38_
            d_1593_v43_: _dafny.Array
            nw247_ = _dafny.Array(None, 7)
            nw247_[int(0)] = d_1526_v4_
            nw247_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhkbfbg"))
            nw247_[int(2)] = (d_1526_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnmatc")))
            nw247_[int(3)] = (d_1584_v36_).f36
            nw247_[int(4)] = default__.fm9((self).f28, globalState)
            nw247_[int(5)] = (d_1526_v4_ if default__.fm40(globalState) else d_1526_v4_)
            nw247_[int(6)] = d_1526_v4_
            d_1593_v43_ = nw247_
            d_1593_v43_ = d_1593_v43_
            if (self).f30:
                d_1594_v44_: _dafny.Map
                d_1594_v44_ = _dafny.Map({d_1527_v5_: 20})
                d_1595_v45_: _dafny.Map
                d_1595_v45_ = _dafny.Map({(self).f28: d_1594_v44_})
                d_1596_v46_: T1
                nw248_ = C3()
                nw248_.ctor__(len(d_1526_v4_), ((d_1595_v45_)[not((self).f28)] if (not((self).f28)) in (d_1595_v45_) else d_1594_v44_), 568, (self).f27, (self).f28)
                d_1596_v46_ = nw248_
                d_1597_v47_: _dafny.Map
                d_1597_v47_ = _dafny.Map({d_1596_v46_.f29: d_1596_v46_})
                d_1598_v48_: D1
                d_1598_v48_ = D1_DC6(self.f29)
                d_1599_v49_: _dafny.MultiSet
                d_1599_v49_ = _dafny.MultiSet([d_1598_v48_])
                d_1600_v50_: _dafny.Array
                nw249_ = _dafny.Array(None, 16)
                nw249_[int(0)] = d_1596_v46_
                nw249_[int(1)] = d_1596_v46_
                nw249_[int(2)] = (d_1596_v46_ if (self).f30 else d_1596_v46_)
                nw249_[int(3)] = d_1596_v46_
                nw249_[int(4)] = d_1596_v46_
                nw249_[int(5)] = d_1596_v46_
                nw249_[int(6)] = d_1596_v46_
                nw249_[int(7)] = d_1596_v46_
                nw249_[int(8)] = d_1596_v46_
                nw249_[int(9)] = d_1596_v46_
                nw249_[int(10)] = ((d_1597_v47_)[default__.fm4(d_1599_v49_, globalState)] if (default__.fm4(d_1599_v49_, globalState)) in (d_1597_v47_) else d_1596_v46_)
                nw249_[int(11)] = d_1596_v46_
                nw249_[int(12)] = d_1596_v46_
                nw249_[int(13)] = d_1596_v46_
                nw249_[int(14)] = d_1596_v46_
                nw249_[int(15)] = d_1596_v46_
                d_1600_v50_ = nw249_
                d_1601_v51_: _dafny.Array
                nw250_ = _dafny.Array(D0.default()(), 1)
                d_1601_v51_ = nw250_
                index232_ = default__.safeIndex(7, (d_1601_v51_).length(0))
                (d_1601_v51_)[index232_] = d_1527_v5_
                index233_ = default__.safeIndex(7, (d_1601_v51_).length(0))
                rhs277_ = d_1600_v50_
                rhs278_ = d_1527_v5_
                lhs230_ = d_1601_v51_
                lhs231_ = default__.safeIndex(7, (d_1601_v51_).length(0))
                d_1600_v50_ = rhs277_
                lhs230_[lhs231_] = rhs278_
                (globalState).f4 = ((d_1582_i3_) * ((0) - (d_1582_i3_))) + (d_1596_v46_.f29)
                d_1602_v52_: _dafny.Seq
                d_1602_v52_ = d_1588_v38_
                d_1603_v53_: _dafny.MultiSet
                d_1603_v53_ = _dafny.MultiSet([(self).f30])
                d_1588_v38_ = (((d_1602_v52_)).set(default__.safeIndex((0) - (len((d_1584_v36_).f36)), len((d_1602_v52_))), (d_1603_v53_).cardinality)) + (_dafny.SeqWithoutIsStrInference([len((d_1584_v36_).f36) for d_1604_i5_ in range(default__.abs(569))]))
                index234_ = default__.safeIndex(6, (d_1522_v0_).length(0))
                (d_1522_v0_)[index234_] = 310
                index235_ = default__.safeIndex(6, (d_1522_v0_).length(0))
                (d_1522_v0_)[index235_] = self.f29
                index236_ = default__.safeIndex(567, (d_1593_v43_).length(0))
                (d_1593_v43_)[index236_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1605_i6_ in range(default__.abs(775))])
                index237_ = default__.safeIndex(567, (d_1593_v43_).length(0))
                (d_1593_v43_)[index237_] = (d_1584_v36_).f36
            elif True:
                index238_ = default__.safeIndex(474, ((self).f27).length(0))
                ((self).f27)[index238_] = (self).f28
                index239_ = default__.safeIndex(474, ((self).f27).length(0))
                ((self).f27)[index239_] = (self).f28
                (globalState).f4 = (-30) + (len((d_1585_v37_)[default__.safeIndex(329, (d_1585_v37_).length(0))]))
                d_1606_v54_: _dafny.Array
                def lambda54_(d_1607_i7_):
                    return not(True)

                init33_ = lambda54_
                nw251_ = _dafny.Array(None, 24)
                for i0_33_ in range(nw251_.length(0)):
                    nw251_[i0_33_] = init33_(i0_33_)
                d_1606_v54_ = nw251_
                d_1606_v54_ = d_1606_v54_
                r1 = not(((self).f27)[default__.safeIndex(474, ((self).f27).length(0))])
                d_1608_v55_: _dafny.Seq
                d_1608_v55_ = _dafny.SeqWithoutIsStrInference([d_1525_v3_, d_1525_v3_, (d_1525_v3_).intersection((d_1585_v37_)[default__.safeIndex(329, (d_1585_v37_).length(0))])])
                d_1608_v55_ = _dafny.SeqWithoutIsStrInference([d_1525_v3_, (d_1585_v37_)[default__.safeIndex(329, (d_1585_v37_).length(0))]])
        d_1609_v56_: _dafny.Set
        d_1609_v56_ = _dafny.Set({(self).f28, not((self).f30), (self).f30, (self).f28, (self).f28})
        d_1609_v56_ = d_1609_v56_
        d_1610_v57_: str
        d_1610_v57_ = _dafny.CodePoint('c')
        d_1611_v58_: _dafny.Map
        d_1611_v58_ = _dafny.Map({(self).f27: d_1610_v57_})
        index240_ = default__.safeIndex(639, ((self).f27).length(0))
        ((self).f27)[index240_] = (self).f30
        d_1612_v59_: _dafny.Seq
        d_1612_v59_ = _dafny.SeqWithoutIsStrInference([not ((self).f30) or ((self).f30), default__.fm16(d_1524_v2_, globalState), ((0) - (self.f29)) < (self.f29)])
        index241_ = default__.safeIndex(639, ((self).f27).length(0))
        rhs279_ = self.f29
        rhs280_ = d_1611_v58_
        rhs281_ = (d_1612_v59_)[default__.safeIndex(self.f29, len(d_1612_v59_))]
        lhs232_ = globalState
        lhs233_ = (self).f27
        lhs234_ = default__.safeIndex(639, ((self).f27).length(0))
        lhs232_.f5 = rhs279_
        d_1611_v58_ = rhs280_
        lhs233_[lhs234_] = rhs281_
        hi11_ = (585) * (self.f29)
        for d_1613_i8_ in range(len(d_1526_v4_), hi11_):
            index242_ = default__.safeIndex(639, ((self).f27).length(0))
            rhs282_ = (self).f28
            rhs283_ = (d_1526_v4_) + (d_1526_v4_)
            rhs284_ = (self.f29) + (default__.fm36(globalState))
            rhs285_ = ((self).f27)[default__.safeIndex(639, ((self).f27).length(0))]
            lhs235_ = (self).f27
            lhs236_ = default__.safeIndex(639, ((self).f27).length(0))
            lhs237_ = globalState
            lhs238_ = globalState
            lhs235_[lhs236_] = rhs282_
            d_1526_v4_ = rhs283_
            lhs237_.f5 = rhs284_
            lhs238_.f17 = rhs285_
            d_1614_v60_: _dafny.Array
            nw252_ = _dafny.Array(False, 11)
            d_1614_v60_ = nw252_
            d_1615_v61_: _dafny.Array
            nw253_ = _dafny.Array(None, 6)
            nw253_[int(0)] = d_1614_v60_
            nw253_[int(1)] = d_1614_v60_
            nw253_[int(2)] = d_1614_v60_
            nw253_[int(3)] = d_1614_v60_
            nw253_[int(4)] = d_1614_v60_
            nw253_[int(5)] = d_1614_v60_
            d_1615_v61_ = nw253_
            d_1616_v62_: D9
            d_1616_v62_ = D9_DC22(d_1522_v0_, not((self).f30), d_1615_v61_, d_1526_v4_)
            r2 = (d_1616_v62_).cf35
            r1 = (self).f28
            d_1617_v63_: _dafny.Set
            d_1617_v63_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odjwsypnq")), d_1526_v4_})
            d_1618_v64_: _dafny.Map
            d_1618_v64_ = _dafny.Map({(self).f28: d_1617_v63_})
            d_1618_v64_ = (d_1618_v64_).set((self).f30, d_1617_v63_)
        r0 = (self.f29) - (self.f29)
        r1 = (self).f30
        r2 = ((self).f30) == ((self).f30)
        r3 = (self).f30
        return r0, r1, r2, r3

    def m2(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        r3: int = int(0)
        d_1619_v0_: _dafny.Seq
        d_1619_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfqmokab"))
        d_1619_v0_ = (default__.fm32(self.f29, (self).f30, globalState)) + (d_1619_v0_)
        (globalState).f13 = self.f29
        d_1620_v1_: _dafny.Array
        nw254_ = _dafny.Array(int(0), 4)
        d_1620_v1_ = nw254_
        d_1621_v2_: _dafny.Array
        nw255_ = _dafny.Array(_dafny.Array(None, 0), 17)
        d_1621_v2_ = nw255_
        d_1622_v3_: D9
        d_1622_v3_ = D9_DC22(d_1620_v1_, (self).f30, d_1621_v2_, d_1619_v0_)
        d_1623_v4_: _dafny.Map
        d_1623_v4_ = _dafny.Map({self.f29: self.f29})
        d_1624_v5_: _dafny.Map
        d_1624_v5_ = _dafny.Map({(self).f28: d_1623_v4_})
        d_1625_v6_: _dafny.Map
        d_1625_v6_ = _dafny.Map({len(d_1624_v5_): (self).f28})
        d_1626_v7_: _dafny.Array
        nw256_ = _dafny.Array(None, 17)
        nw256_[int(0)] = (self).f30
        nw256_[int(1)] = (self).f28
        nw256_[int(2)] = (self).f28
        nw256_[int(3)] = (d_1622_v3_).cf35
        nw256_[int(4)] = (self).f28
        nw256_[int(5)] = False
        nw256_[int(6)] = True
        nw256_[int(7)] = (self.f29) != (-922)
        nw256_[int(8)] = (self).f30
        nw256_[int(9)] = (self).f30
        nw256_[int(10)] = (self).f28
        nw256_[int(11)] = (self).f30
        nw256_[int(12)] = (D9_DC22(d_1620_v1_, (self).f28, d_1621_v2_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1627_i1_ in range(default__.abs(-72))]))).cf35
        nw256_[int(13)] = (self).f30
        nw256_[int(14)] = (self).f30
        nw256_[int(15)] = ((d_1625_v6_)[self.f29] if (self.f29) in (d_1625_v6_) else not(False))
        nw256_[int(16)] = True
        d_1626_v7_ = nw256_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1626_v7_).length(0)):
            d_1628_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_1628_i0_)) and ((d_1628_i0_) < ((d_1626_v7_).length(0)))):
                def iife119_():
                    coll73_ = _dafny.Set()
                    compr_73_: D7
                    for compr_73_ in (_dafny.Set({D7_DC16((self).f30, _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([747, self.f29, self.f29])).cardinality])): (self).f28}), _dafny.CodePoint('q'), (self).f28), D7_DC16((self).f30, d_1625_v6_, _dafny.CodePoint('k'), not((self).f28)), D7_DC16((self).f30, (d_1625_v6_).set(self.f29, False), _dafny.CodePoint('g'), (self).f30), D7_DC16((self).f28, d_1625_v6_, _dafny.CodePoint('v'), (self).f30), D7_DC16((self).f30, d_1625_v6_, _dafny.CodePoint('b'), (self).f28)})).Elements:
                        d_1629_v8_: D7 = compr_73_
                        if (d_1629_v8_) in (_dafny.Set({D7_DC16((self).f30, _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([747, self.f29, self.f29])).cardinality])): (self).f28}), _dafny.CodePoint('q'), (self).f28), D7_DC16((self).f30, d_1625_v6_, _dafny.CodePoint('k'), not((self).f28)), D7_DC16((self).f30, (d_1625_v6_).set(self.f29, False), _dafny.CodePoint('g'), (self).f30), D7_DC16((self).f28, d_1625_v6_, _dafny.CodePoint('v'), (self).f30), D7_DC16((self).f30, d_1625_v6_, _dafny.CodePoint('b'), (self).f28)})):
                            coll73_ = coll73_.union(_dafny.Set([d_1629_v8_]))
                    return _dafny.Set(coll73_)
                def iife120_():
                    coll74_ = _dafny.Map()
                    compr_74_: int
                    for compr_74_ in ((_dafny.SeqWithoutIsStrInference([self.f29 for d_1630_i2_ in range(default__.abs(388))]))).Elements:
                        d_1631_v9_: int = compr_74_
                        if (d_1631_v9_) in ((_dafny.SeqWithoutIsStrInference([self.f29 for d_1630_i2_ in range(default__.abs(388))]))):
                            coll74_[default__.safeDivisionInt(d_1631_v9_, self.f29)] = (self).f30
                    return _dafny.Map(coll74_)
                (d_1626_v7_)[(d_1628_i0_)] = ((_dafny.Set({D7_DC16((self).f30, d_1625_v6_, _dafny.CodePoint('b'), (self).f30), D7_DC16((self).f28, _dafny.Map({self.f29: (self).f30}), _dafny.CodePoint('f'), (self).f30), D7_DC16((self).f28, d_1625_v6_, _dafny.CodePoint('y'), (self).f28), D7_DC16((self).f28, _dafny.Map({self.f29: (self).f30}), _dafny.CodePoint('k'), (self).f30), D7_DC16((self).f28, d_1625_v6_, _dafny.CodePoint('p'), (self).f28)})) - (_dafny.Set({D7_DC16(True, (d_1625_v6_).set((_dafny.MultiSet([(self).f28, (self).f28])).cardinality, (self).f28), _dafny.CodePoint('d'), (self).f28), D7_DC16((self).f28, (_dafny.Map({self.f29: (self).f28})).set(558, (self).f28), _dafny.CodePoint('m'), True)}))) in ((_dafny.SeqWithoutIsStrInference([iife119_()
                ])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({D7_DC16((self).f30, (iife120_()
).set(self.f29, (self).f30), _dafny.CodePoint('k'), (self).f28), D7_DC16((self).f30, d_1625_v6_, _dafny.CodePoint('j'), (self).f28), D7_DC16((self).f28, d_1625_v6_, _dafny.CodePoint('h'), (self).f30), D7_DC16((self).f28, d_1625_v6_, _dafny.CodePoint('o'), (self).f28)})])))
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1620_v1_).length(0)):
            d_1632_i3_: int = guard_loop_3_
            if (True) and (((0) <= (d_1632_i3_)) and ((d_1632_i3_) < ((d_1620_v1_).length(0)))):
                (d_1620_v1_)[(d_1632_i3_)] = (d_1632_i3_) + (len(_dafny.Map({d_1623_v4_: (self).f28})))
        r0 = (self.f29) == (self.f29)
        d_1633_v10_: _dafny.Set
        d_1633_v10_ = _dafny.Set({self.f29})
        def iife121_():
            coll75_ = _dafny.Set()
            compr_75_: int
            for compr_75_ in (_dafny.SeqWithoutIsStrInference([self.f29, 250])).Elements:
                d_1634_v11_: int = compr_75_
                if (d_1634_v11_) in (_dafny.SeqWithoutIsStrInference([self.f29, 250])):
                    coll75_ = coll75_.union(_dafny.Set([default__.safeDivisionInt(d_1634_v11_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: False})), 44])))]))
            return _dafny.Set(coll75_)
        d_1633_v10_ = (d_1633_v10_) | (iife121_()
        )
        r0 = False
        r1 = self.f29
        r2 = not ((self).f30) or ((self.f29) != (self.f29))
        d_1635_v12_: _dafny.Map
        d_1635_v12_ = _dafny.Map({(self).f28: (self).f30})
        d_1636_v13_: _dafny.Map
        d_1636_v13_ = _dafny.Map({d_1635_v12_: (self).f28})
        d_1637_v14_: _dafny.Seq
        d_1637_v14_ = _dafny.SeqWithoutIsStrInference([self.f29])
        r3 = len((_dafny.SeqWithoutIsStrInference([len(d_1636_v13_)])) + (d_1637_v14_))
        return r0, r1, r2, r3

    @property
    def f30(self):
        return self._f30
    @property
    def f31(self):
        return self._f31
