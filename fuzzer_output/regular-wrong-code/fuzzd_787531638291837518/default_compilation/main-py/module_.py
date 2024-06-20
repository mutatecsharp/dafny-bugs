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
    def fm5(globalState):
        if (608) != (898):
            return _dafny.CodePoint('v')
        elif True:
            return _dafny.CodePoint('l')

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([(D1_DC1(_dafny.CodePoint('d'))).cf1 for d_0_i0_ in range(default__.abs(100))])

    @staticmethod
    def fm9(p0, globalState):
        return ((_dafny.Map({True: False})) | (_dafny.Map({True: False}))) | (_dafny.Map({not(not(True)): True}))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        source0_ = D11_DC30()
        if source0_.is_DC29:
            d_1___mcc_h0_ = source0_.cf43
            d_2_cf43_ = d_1___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([True for d_3_i0_ in range(default__.abs(-222))])) + (_dafny.SeqWithoutIsStrInference([False, False, False]))
        elif source0_.is_DC30:
            return _dafny.SeqWithoutIsStrInference([not(False), not(not(not(False))), not(True), True])
        elif source0_.is_DC28:
            d_4___mcc_h1_ = source0_.cf42
            d_5_cf42_ = d_4___mcc_h1_
            return (_dafny.SeqWithoutIsStrInference([True, True, not(not(False)), True])) + (_dafny.SeqWithoutIsStrInference([False for d_6_i1_ in range(default__.abs(593))]))
        elif True:
            d_7___mcc_h2_ = source0_.cf44
            d_8_cf44_ = d_7___mcc_h2_
            return (_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm16(p0, globalState):
        return (0) - ((0) - (len(((_dafny.Set({True, False})) - (_dafny.Set({False}))).intersection(_dafny.Set({not(not(False))})))))

    @staticmethod
    def fm17(p0, p1, globalState):
        return _dafny.Map({True: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_9_i0_ in range(default__.abs(-771))]))])})

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        return ((0) - (-461)) * (((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('e'): True})) for d_10_i0_ in range(default__.abs(-343))]))).cardinality) + (518))

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(927, 999):
                d_11_v0_: int = compr_0_
                if ((927) <= (d_11_v0_)) and ((d_11_v0_) < (999)):
                    coll0_[(d_11_v0_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykwfkuyl"))))] = 298
            return _dafny.Map(coll0_)
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvab")): True}))])) + (_dafny.SeqWithoutIsStrInference([len(iife0_()
        ), 312, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))), 792]))) + (_dafny.SeqWithoutIsStrInference([921]))

    @staticmethod
    def fm22(globalState):
        return (_dafny.MultiSet([D1_DC3(D1_DC2())])) | (_dafny.MultiSet([D1_DC3(D1_DC1(_dafny.CodePoint('y'))), D1_DC3(D1_DC2()), D1_DC3(D1_DC2())]))

    @staticmethod
    def fm23(p0, globalState):
        source1_ = False
        d_12___mcc_h0_ = source1_
        d_13_cf0_ = d_12___mcc_h0_
        return D1_DC3(D1_DC2())

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        return not((len(_dafny.SeqWithoutIsStrInference([not(True), not(False)]))) != (571))

    @staticmethod
    def fm25(globalState):
        return (_dafny.Map({False: not(True)})) | ((_dafny.Map({True: False})) | (_dafny.Map({False: False})))

    @staticmethod
    def fm26(p0, p1, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(-907, -218):
                d_14_v0_: int = compr_1_
                if ((-907) <= (d_14_v0_)) and ((d_14_v0_) < (-218)):
                    coll1_ = coll1_.union(_dafny.Set([(d_14_v0_) * (-978)]))
            return _dafny.Set(coll1_)
        return D3_DC6((_dafny.Set({644})) | (iife1_()
))

    @staticmethod
    def fm27(globalState):
        return (_dafny.Set({_dafny.MultiSet([True, True, False])})).intersection(_dafny.Set({_dafny.MultiSet([True])}))

    @staticmethod
    def fm28(p0, globalState):
        return (D2_DC4(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True)])))).cf3

    @staticmethod
    def fm29(p0, p1, globalState):
        source2_ = D10_DC25(_dafny.Map({529: ((D2_DC4(_dafny.MultiSet([False]))).cf3).cardinality}))
        if source2_.is_DC26:
            d_15___mcc_h0_ = source2_.cf40
            d_16_cf40_ = d_15___mcc_h0_
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in _dafny.IntegerRange(-914, 258):
                    d_17_v0_: int = compr_2_
                    if ((-914) <= (d_17_v0_)) and ((d_17_v0_) < (258)):
                        coll2_[(d_17_v0_) - (d_16_cf40_)] = True
                return _dafny.Map(coll2_)
            return (D3_DC6(_dafny.Set({len(iife2_()
)}))).cf9
        elif source2_.is_DC25:
            d_18___mcc_h1_ = source2_.cf39
            d_19_cf39_ = d_18___mcc_h1_
            def iife3_():
                coll3_ = _dafny.Set()
                compr_3_: int
                for compr_3_ in _dafny.IntegerRange(574, 360):
                    d_20_v1_: int = compr_3_
                    if ((574) <= (d_20_v1_)) and ((d_20_v1_) < (360)):
                        def iife4_():
                            coll4_ = _dafny.Set()
                            compr_4_: int
                            for compr_4_ in _dafny.IntegerRange(772, 533):
                                d_21_v2_: int = compr_4_
                                if ((772) <= (d_21_v2_)) and ((d_21_v2_) < (533)):
                                    coll4_ = coll4_.union(_dafny.Set([(d_21_v2_) + (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True, False])).cardinality, len(_dafny.SeqWithoutIsStrInference([-308])), -914, len(_dafny.Map({False: True})), (_dafny.MultiSet([50])).cardinality])))]))
                            return _dafny.Set(coll4_)
                        coll3_ = coll3_.union(_dafny.Set([default__.safeDivisionInt(d_20_v1_, len(_dafny.SeqWithoutIsStrInference([len(iife4_()
)])))]))
                return _dafny.Set(coll3_)
            return iife3_()
            
        elif True:
            d_22___mcc_h2_ = source2_.cf41
            d_23_cf41_ = d_22___mcc_h2_
            return _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlnmqckl")))})

    @staticmethod
    def fm30(p0, globalState):
        def iife5_():
            def iife7_():
                coll7_ = _dafny.Set()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(853, 631):
                    d_26_v1_: int = compr_7_
                    if ((853) <= (d_26_v1_)) and ((d_26_v1_) < (631)):
                        coll7_ = coll7_.union(_dafny.Set([(d_26_v1_) + (len(_dafny.Set({False, False})))]))
                return _dafny.Set(coll7_)
            coll5_ = _dafny.Map()
            def iife6_():
                coll6_ = _dafny.Set()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(853, 631):
                    d_24_v1_: int = compr_6_
                    if ((853) <= (d_24_v1_)) and ((d_24_v1_) < (631)):
                        coll6_ = coll6_.union(_dafny.Set([(d_24_v1_) + (len(_dafny.Set({False, False})))]))
                return _dafny.Set(coll6_)
            compr_5_: D3
            for compr_5_ in (_dafny.Map({D3_DC6(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsyxbepi"))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(iife6_()
): len(_dafny.Set({True}))})), len(_dafny.Set({True, False, not(False)}))]))})): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "veyog"))})).keys.Elements:
                d_25_v0_: D3 = compr_5_
                if (d_25_v0_) in (_dafny.Map({D3_DC6(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsyxbepi"))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(iife7_()
): len(_dafny.Set({True}))})), len(_dafny.Set({True, False, not(False)}))]))})): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "veyog"))})):
                    coll5_[d_25_v0_] = -649
            return _dafny.Map(coll5_)
        if (720) < (len(iife5_()
        )):
            return (_dafny.SeqWithoutIsStrInference([not(True)])) + (_dafny.SeqWithoutIsStrInference([False]))
        elif not(not(True)):
            return _dafny.SeqWithoutIsStrInference([True, True, True, False, True])
        elif True:
            return _dafny.SeqWithoutIsStrInference([True, not(True)])

    @staticmethod
    def fm31(p0, p1, globalState):
        return _dafny.Set({(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([not(not(False)), False]))): len(_dafny.SeqWithoutIsStrInference([617]))})) | (_dafny.Map({55: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_27_i0_ in range(default__.abs(-936))]))})), (_dafny.Map({424: (D8_DC22(len(_dafny.SeqWithoutIsStrInference([-954, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qejrjkfg")))])))).cf35})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([360 for d_28_i1_ in range(default__.abs(140))])): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhvkev")))})), _dafny.Map({713: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_29_i2_ in range(default__.abs(108))]))})})

    @staticmethod
    def fm32(globalState):
        source3_ = D5_DC13(False, _dafny.CodePoint('y'))
        if source3_.is_DC12:
            d_30___mcc_h0_ = source3_.cf19
            d_31_cf19_ = d_30___mcc_h0_
            return D5_DC13(d_31_cf19_, _dafny.CodePoint('r'))
        elif source3_.is_DC13:
            d_32___mcc_h1_ = source3_.cf20
            d_33___mcc_h2_ = source3_.cf21
            d_34_cf21_ = d_33___mcc_h2_
            d_35_cf20_ = d_32___mcc_h1_
            return D5_DC13(d_35_cf20_, d_34_cf21_)
        elif source3_.is_DC11:
            d_36___mcc_h3_ = source3_.cf18
            d_37_cf18_ = d_36___mcc_h3_
            return D5_DC13(True, _dafny.CodePoint('f'))
        elif True:
            d_38___mcc_h4_ = source3_.cf22
            d_39_cf22_ = d_38___mcc_h4_
            return D5_DC13(False, _dafny.CodePoint('f'))

    @staticmethod
    def fm33(globalState):
        return ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqnloowd")))))) == (-548)

    @staticmethod
    def fm34(p0, p1, globalState):
        def iife8_():
            coll8_ = _dafny.Set()
            compr_8_: D6
            for compr_8_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D6_DC16(_dafny.Map({_dafny.MultiSet([True]): 699}), _dafny.SeqWithoutIsStrInference([False, True]))]))).Elements:
                d_40_v0_: D6 = compr_8_
                if (d_40_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D6_DC16(_dafny.Map({_dafny.MultiSet([True]): 699}), _dafny.SeqWithoutIsStrInference([False, True]))]))):
                    coll8_ = coll8_.union(_dafny.Set([d_40_v0_]))
            return _dafny.Set(coll8_)
        return (_dafny.Set({D6_DC16(_dafny.Map({_dafny.MultiSet([False]): len(_dafny.Map({False: 155}))}), _dafny.SeqWithoutIsStrInference([True])), D6_DC16(_dafny.Map({_dafny.MultiSet([False]): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True]))]))}), _dafny.SeqWithoutIsStrInference([False]))})).intersection(iife8_()
        )

    @staticmethod
    def fm35(p0, globalState):
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: _dafny.MultiSet
            for compr_9_ in (_dafny.MultiSet([_dafny.MultiSet([True]), _dafny.MultiSet([True])])).Elements:
                d_41_v0_: _dafny.MultiSet = compr_9_
                if (d_41_v0_) in (_dafny.MultiSet([_dafny.MultiSet([True]), _dafny.MultiSet([True])])):
                    coll9_[d_41_v0_] = len(_dafny.SeqWithoutIsStrInference([False]))
            return _dafny.Map(coll9_)
        def iife10_():
            def iife12_():
                coll12_ = _dafny.Map()
                compr_12_: _dafny.MultiSet
                for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, True]) for d_42_i0_ in range(default__.abs(83))])).Elements:
                    d_43_v2_: _dafny.MultiSet = compr_12_
                    if (d_43_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, True]) for d_42_i0_ in range(default__.abs(83))])):
                        coll12_[d_43_v2_] = 775
                return _dafny.Map(coll12_)
            coll10_ = _dafny.Map()
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: _dafny.MultiSet
                for compr_11_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, True]) for d_42_i0_ in range(default__.abs(83))])).Elements:
                    d_43_v2_: _dafny.MultiSet = compr_11_
                    if (d_43_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, True]) for d_42_i0_ in range(default__.abs(83))])):
                        coll11_[d_43_v2_] = 775
                return _dafny.Map(coll11_)
            compr_10_: _dafny.MultiSet
            for compr_10_ in (iife11_()
            ).keys.Elements:
                d_44_v1_: _dafny.MultiSet = compr_10_
                if (d_44_v1_) in (iife12_()
                ):
                    coll10_[d_44_v1_] = len(_dafny.SeqWithoutIsStrInference([True, True, False, False, not(False)]))
            return _dafny.Map(coll10_)
        return D6_DC16((iife9_()
) | (iife10_()
), _dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm36(p0, p1, globalState):
        def iife13_():
            coll13_ = _dafny.Set()
            compr_13_: _dafny.MultiSet
            for compr_13_ in (_dafny.Set({_dafny.MultiSet([D1_DC3(D1_DC2()), D1_DC3(D1_DC1(_dafny.CodePoint('f')))])})).Elements:
                d_45_v0_: _dafny.MultiSet = compr_13_
                if (d_45_v0_) in (_dafny.Set({_dafny.MultiSet([D1_DC3(D1_DC2()), D1_DC3(D1_DC1(_dafny.CodePoint('f')))])})):
                    coll13_ = coll13_.union(_dafny.Set([d_45_v0_]))
            return _dafny.Set(coll13_)
        return (iife13_()
        ) | ((_dafny.Set({_dafny.MultiSet([D1_DC3(D1_DC3(D1_DC2())), D1_DC3(D1_DC2()), D1_DC3(D1_DC2()), D1_DC3(D1_DC2())]), _dafny.MultiSet([D1_DC3(D1_DC3(D1_DC2()))])})).intersection(_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D1_DC3(D1_DC2()), D1_DC3(D1_DC1(_dafny.CodePoint('m'))), D1_DC3(D1_DC2())]))})))

    @staticmethod
    def fm37(globalState):
        if not((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npm")))) < (-818)):
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in (_dafny.Set({93, len(_dafny.SeqWithoutIsStrInference([False])), 694, 980})).Elements:
                    d_46_v0_: int = compr_14_
                    if (d_46_v0_) in (_dafny.Set({93, len(_dafny.SeqWithoutIsStrInference([False])), 694, 980})):
                        coll14_[(d_46_v0_) + (302)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_47_i0_ in range(default__.abs(184))])
                return _dafny.Map(coll14_)
            return iife14_()
            
        elif True:
            def iife15_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in (_dafny.SeqWithoutIsStrInference([732, 894])).Elements:
                    d_48_v1_: int = compr_15_
                    if (d_48_v1_) in (_dafny.SeqWithoutIsStrInference([732, 894])):
                        coll15_[(d_48_v1_) - (-626)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nf"))
                return _dafny.Map(coll15_)
            return (iife15_()
            ) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_49_i1_ in range(default__.abs(-552))])}))

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: _dafny.MultiSet
            for compr_16_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False]))])).Elements:
                d_50_v0_: _dafny.MultiSet = compr_16_
                if (d_50_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False]))])):
                    coll16_[d_50_v0_] = -237
            return _dafny.Map(coll16_)
        if (len(iife16_()
        )) < (47):
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in (_dafny.Map({(_dafny.MultiSet([not(False)])).cardinality: False})).keys.Elements:
                    d_51_v1_: int = compr_17_
                    if (d_51_v1_) in (_dafny.Map({(_dafny.MultiSet([not(False)])).cardinality: False})):
                        coll17_[(d_51_v1_) * ((0) - (-744))] = -645
                return _dafny.Map(coll17_)
            return (_dafny.Map({(0) - (-275): 295})) | (iife17_()
            )
        elif True:
            def iife18_():
                coll18_ = _dafny.Map()
                compr_18_: int
                for compr_18_ in _dafny.IntegerRange(887, -826):
                    d_52_v2_: int = compr_18_
                    if ((887) <= (d_52_v2_)) and ((d_52_v2_) < (-826)):
                        coll18_[default__.safeModuloInt(d_52_v2_, -828)] = True
                return _dafny.Map(coll18_)
            return _dafny.Map({len(iife18_()
            ): 28})

    @staticmethod
    def fm39(p0, p1, p2, p3, globalState):
        def iife19_():
            coll19_ = _dafny.Map()
            compr_19_: int
            for compr_19_ in (_dafny.Map({500: len(_dafny.Map({_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(True)])).cardinality]): True}))})).keys.Elements:
                d_53_v0_: int = compr_19_
                if (d_53_v0_) in (_dafny.Map({500: len(_dafny.Map({_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(True)])).cardinality]): True}))})):
                    coll19_[default__.safeModuloInt(d_53_v0_, -296)] = _dafny.Set({True})
            return _dafny.Map(coll19_)
        def iife20_():
            coll20_ = _dafny.Map()
            compr_20_: int
            for compr_20_ in (_dafny.SeqWithoutIsStrInference([405 for d_54_i0_ in range(default__.abs(667))])).Elements:
                d_55_v1_: int = compr_20_
                if (d_55_v1_) in (_dafny.SeqWithoutIsStrInference([405 for d_54_i0_ in range(default__.abs(667))])):
                    coll20_[(d_55_v1_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqnmhka"))))] = 965
            return _dafny.Map(coll20_)
        return ((_dafny.Map({True: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(iife19_()
        )]))).cardinality})) | (_dafny.Map({True: len(_dafny.Set({True, True}))}))) | ((_dafny.Map({False: len(iife20_()
        )})) | (_dafny.Map({not(not(not(True))): 224})))

    @staticmethod
    def m0(p0, p1, p2, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        d_56_v0_: _dafny.Seq
        d_56_v0_ = _dafny.SeqWithoutIsStrInference([-104])
        d_57_v1_: _dafny.Map
        d_57_v1_ = _dafny.Map({d_56_v0_: p2})
        d_58_v2_: D7
        d_58_v2_ = D7_DC18(p2, ((d_57_v1_)[_dafny.SeqWithoutIsStrInference([p0 for d_59_i0_ in range(default__.abs(748))])] if (_dafny.SeqWithoutIsStrInference([p0 for d_60_i0_ in range(default__.abs(748))])) in (d_57_v1_) else p1))
        source4_ = d_58_v2_
        if source4_.is_DC18:
            d_61___mcc_h0_ = source4_.cf27
            d_62___mcc_h1_ = source4_.cf28
            d_63_cf28_ = d_62___mcc_h1_
            d_64_cf27_ = d_61___mcc_h0_
            d_65_v3_: _dafny.Seq
            d_65_v3_ = _dafny.SeqWithoutIsStrInference([not(True)])
            d_66_v4_: _dafny.Map
            d_66_v4_ = _dafny.Map({D7_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlycpc"))): d_65_v3_})
            d_67_v5_: _dafny.Map
            d_67_v5_ = _dafny.Map({d_66_v4_: (True if p2 else p1)})
            d_67_v5_ = (d_67_v5_).set(d_66_v4_, d_63_cf28_)
            d_68_v6_: _dafny.Map
            d_68_v6_ = _dafny.Map({p0: d_64_cf27_})
            d_69_v7_: _dafny.Map
            d_69_v7_ = _dafny.Map({d_68_v6_: p0})
            d_70_v8_: _dafny.Seq
            d_70_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            d_71_v9_: C5
            nw0_ = C5()
            nw0_.ctor__(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_72_i1_ in range(default__.abs(911))])).set(default__.safeIndex(len(d_69_v7_), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_73_i1_ in range(default__.abs(911))]))), _dafny.CodePoint('n'))) + (d_70_v8_), (0) - ((0) - (p0)), p0, p0)
            d_71_v9_ = nw0_
            d_74_v11_: C4
            nw1_ = C4()
            def iife21_():
                coll21_ = _dafny.Map()
                compr_21_: int
                for compr_21_ in (d_68_v6_).keys.Elements:
                    d_75_v10_: int = compr_21_
                    if (d_75_v10_) in (d_68_v6_):
                        coll21_[default__.safeModuloInt(d_75_v10_, (d_71_v9_).f10)] = (d_71_v9_).f10
                return _dafny.Map(coll21_)
            nw1_.ctor__(p1, len(iife21_()
            ), len(default__.fm21(p0, (0) - (p0), p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etffgh")), globalState)))
            d_74_v11_ = nw1_
            d_76_v12_: _dafny.MultiSet
            d_76_v12_ = _dafny.MultiSet([d_74_v11_])
            d_77_v13_: _dafny.Seq
            d_77_v13_ = _dafny.SeqWithoutIsStrInference([d_74_v11_])
            d_78_v14_: _dafny.MultiSet
            d_78_v14_ = _dafny.MultiSet([d_76_v12_, d_76_v12_, _dafny.MultiSet([d_74_v11_, d_74_v11_]), _dafny.MultiSet(d_77_v13_)])
            (globalState).f3 = (_dafny.MultiSet([d_76_v12_, _dafny.MultiSet([d_74_v11_])])).issubset(d_78_v14_)
            d_79_v15_: int
            d_79_v15_ = -986
            d_79_v15_ = len(_dafny.Set({(d_71_v9_).f10, 877, p0, d_79_v15_}))
        elif source4_.is_DC19:
            d_80___mcc_h2_ = source4_.cf29
            d_81___mcc_h3_ = source4_.cf30
            d_82___mcc_h4_ = source4_.cf31
            d_83_cf31_ = d_82___mcc_h4_
            d_84_cf30_ = d_81___mcc_h3_
            d_85_cf29_ = d_80___mcc_h2_
            d_86_v16_: C3
            nw2_ = C3()
            nw2_.ctor__(p2, default__.safeModuloInt(default__.fm20(p2, p1, not(p1), globalState), d_83_cf31_), -709)
            d_86_v16_ = nw2_
            d_87_v17_: _dafny.Seq
            d_87_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtmdkqwc"))
            d_88_v18_: _dafny.Set
            d_88_v18_ = _dafny.Set({(d_86_v16_).fm13(d_87_v17_, globalState), p0})
            d_89_v19_: _dafny.Map
            d_89_v19_ = _dafny.Map({(d_88_v18_).intersection(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([p1, not(not(not(p1)))])), p0})): (True if (d_86_v16_).f12 else d_84_cf30_)})
            d_89_v19_ = (d_89_v19_).set(d_88_v18_, not (d_84_cf30_) or ((d_86_v16_).f12))
            d_83_cf31_ = d_83_cf31_
            d_84_cf30_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ruevlxjng"))) <= ((d_87_v17_) + (d_87_v17_))
        elif source4_.is_DC20:
            d_90___mcc_h5_ = source4_.cf32
            d_91___mcc_h6_ = source4_.cf33
            d_92_cf33_ = d_91___mcc_h6_
            d_93_cf32_ = d_90___mcc_h5_
            (globalState).f4 = p1
            d_94_v20_: _dafny.Array
            nw3_ = _dafny.Array(int(0), 4)
            d_94_v20_ = nw3_
            d_95_v21_: _dafny.Array
            nw4_ = _dafny.Array(_dafny.Map({}), 12)
            d_95_v21_ = nw4_
            d_96_v22_: _dafny.Seq
            d_96_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjka"))
            d_97_v23_: _dafny.Map
            d_97_v23_ = _dafny.Map({d_92_cf33_: d_96_v22_})
            index0_ = default__.safeIndex(509, (d_95_v21_).length(0))
            (d_95_v21_)[index0_] = d_97_v23_
            d_98_v24_: D9
            d_98_v24_ = D9_DC23(d_94_v20_)
            d_99_v25_: _dafny.Map
            d_99_v25_ = _dafny.Map({p2: p2})
            d_100_v26_: _dafny.Seq
            d_100_v26_ = _dafny.SeqWithoutIsStrInference([p1, not (p1) or (((d_99_v25_)[False] if (False) in (d_99_v25_) else p1))])
            d_101_v27_: _dafny.MultiSet
            d_101_v27_ = _dafny.MultiSet([-877])
            d_102_v28_: _dafny.Set
            d_102_v28_ = _dafny.Set({p0})
            d_103_v29_: C5
            nw5_ = C5()
            nw5_.ctor__(d_96_v22_, d_92_cf33_, (d_101_v27_).cardinality, len(d_102_v28_))
            d_103_v29_ = nw5_
            d_104_v30_: _dafny.Map
            d_104_v30_ = _dafny.Map({d_103_v29_: p2})
            d_105_v31_: _dafny.Map
            d_105_v31_ = _dafny.Map({default__.safeModuloInt((d_103_v29_).f10, p0): (default__.fm37(globalState)) | (default__.fm37(globalState))})
            index1_ = default__.safeIndex(509, (d_95_v21_).length(0))
            rhs0_ = p1
            rhs1_ = (d_98_v24_).cf36
            rhs2_ = (d_100_v26_)[default__.safeIndex((len(d_104_v30_)) * ((d_103_v29_).f10), len(d_100_v26_))]
            rhs3_ = ((d_105_v31_)[(0) - ((d_103_v29_).f10)] if ((0) - ((d_103_v29_).f10)) in (d_105_v31_) else (d_97_v23_) | (d_97_v23_))
            lhs0_ = globalState
            lhs1_ = globalState
            lhs2_ = d_95_v21_
            lhs3_ = default__.safeIndex(509, (d_95_v21_).length(0))
            lhs0_.f4 = rhs0_
            d_94_v20_ = rhs1_
            lhs1_.f4 = rhs2_
            lhs2_[lhs3_] = rhs3_
            d_106_v32_: _dafny.Array
            nw6_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
            d_106_v32_ = nw6_
            d_107_v33_: _dafny.Map
            d_107_v33_ = _dafny.Map({d_106_v32_: -969})
            d_107_v33_ = (d_107_v33_).set(d_106_v32_, d_92_cf33_)
            d_92_cf33_ = default__.safeModuloInt((d_92_cf33_) - (d_92_cf33_), p0)
        elif True:
            d_108___mcc_h7_ = source4_.cf26
            d_109_cf26_ = d_108___mcc_h7_
            d_110_v34_: _dafny.Seq
            d_110_v34_ = _dafny.SeqWithoutIsStrInference([not(True)])
            d_111_v35_: _dafny.Array
            nw7_ = _dafny.Array(None, 18)
            nw7_[int(0)] = len(default__.fm30(False, globalState))
            nw7_[int(1)] = 122
            nw7_[int(2)] = len(_dafny.SeqWithoutIsStrInference([True]))
            nw7_[int(3)] = default__.fm20(p1, p2, p1, globalState)
            nw7_[int(4)] = p0
            nw7_[int(5)] = p0
            nw7_[int(6)] = p0
            nw7_[int(7)] = p0
            nw7_[int(8)] = p0
            nw7_[int(9)] = p0
            nw7_[int(10)] = 674
            nw7_[int(11)] = p0
            nw7_[int(12)] = p0
            nw7_[int(13)] = (0) - (len(d_110_v34_))
            nw7_[int(14)] = (0) - (p0)
            nw7_[int(15)] = p0
            nw7_[int(16)] = p0
            nw7_[int(17)] = 885
            d_111_v35_ = nw7_
            d_112_v36_: _dafny.Seq
            d_112_v36_ = _dafny.SeqWithoutIsStrInference([d_111_v35_, d_111_v35_])
            d_113_v37_: _dafny.Map
            d_113_v37_ = _dafny.Map({p1: (d_112_v36_)[default__.safeIndex(p0, len(d_112_v36_))]})
            d_114_v38_: _dafny.Map
            d_114_v38_ = _dafny.Map({(0) - (p0): p1})
            d_113_v37_ = (d_113_v37_).set(((d_114_v38_)[p0] if (p0) in (d_114_v38_) else p1), d_111_v35_)
            (globalState).f4 = ((False if p2 else p1) if p2 else True)
            if (p0) > (240):
                d_115_v39_: C4
                nw8_ = C4()
                nw8_.ctor__(p1, len((default__.fm30(not(not(False)), globalState)).set(default__.safeIndex(p0, len(default__.fm30(not(not(False)), globalState))), p1)), (0) - (p0))
                d_115_v39_ = nw8_
                d_116_v40_: T0
                nw9_ = C1()
                nw9_.ctor__(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhnx"))))
                d_116_v40_ = nw9_
                d_117_v41_: _dafny.Seq
                d_117_v41_ = _dafny.SeqWithoutIsStrInference([d_116_v40_, d_116_v40_])
                rhs4_ = d_115_v39_
                rhs5_ = d_117_v41_
                d_115_v39_ = rhs4_
                d_117_v41_ = rhs5_
                d_118_v42_: int
                d_118_v42_ = 0
                d_118_v42_ = default__.safeModuloInt(default__.safeModuloInt((d_116_v40_).f8, (d_116_v40_).f7), (0) - ((0) - (d_118_v42_)))
                d_56_v0_ = (d_56_v0_ if p1 else d_56_v0_)
                index2_ = default__.safeIndex(581, (d_111_v35_).length(0))
                (d_111_v35_)[index2_] = (0) - (p0)
                d_119_v43_: _dafny.Map
                d_119_v43_ = _dafny.Map({p1: p0})
                index3_ = default__.safeIndex(581, (d_111_v35_).length(0))
                (d_111_v35_)[index3_] = (d_118_v42_) * (((d_119_v43_)[default__.fm24((d_115_v39_).f11, (d_116_v40_).f7, d_109_cf26_, globalState)] if (default__.fm24((d_115_v39_).f11, (d_116_v40_).f7, d_109_cf26_, globalState)) in (d_119_v43_) else d_118_v42_))
                d_120_v44_: _dafny.MultiSet
                d_120_v44_ = _dafny.MultiSet([(default__.fm28(p2, globalState)).cardinality, d_118_v42_, len(d_109_cf26_), d_118_v42_])
                d_121_v45_: _dafny.Array
                nw10_ = _dafny.Array(None, 9)
                nw10_[int(0)] = (d_115_v39_).f11
                nw10_[int(1)] = p2
                nw10_[int(2)] = (d_118_v42_) <= ((d_116_v40_).f7)
                nw10_[int(3)] = p2
                nw10_[int(4)] = not (p1) or (p1)
                nw10_[int(5)] = p2
                nw10_[int(6)] = True
                nw10_[int(7)] = (d_120_v44_) == (d_120_v44_)
                nw10_[int(8)] = (d_115_v39_).f11
                d_121_v45_ = nw10_
                d_122_v46_: D10
                d_122_v46_ = D10_DC25(default__.fm38((d_115_v39_).fm1(_dafny.Set({p2}), p0, (0) - (d_118_v42_), 554, globalState), True, default__.fm24(p2, 816, d_109_cf26_, globalState), (d_116_v40_).f8, globalState))
                d_123_v47_: _dafny.Set
                d_123_v47_ = _dafny.Set({len((d_122_v46_).cf39)})
                index4_ = default__.safeIndex(781, (d_121_v45_).length(0))
                def iife22_():
                    coll22_ = _dafny.Set()
                    compr_22_: int
                    for compr_22_ in _dafny.IntegerRange(107, 642):
                        d_124_v48_: int = compr_22_
                        if ((107) <= (d_124_v48_)) and ((d_124_v48_) < (642)):
                            coll22_ = coll22_.union(_dafny.Set([(d_124_v48_) * ((d_116_v40_).f8)]))
                    return _dafny.Set(coll22_)
                (d_121_v45_)[index4_] = (d_123_v47_).issubset(iife22_()
                )
                d_125_v49_: _dafny.Array
                nw11_ = _dafny.Array(_dafny.CodePoint('D'), 29)
                d_125_v49_ = nw11_
                index5_ = default__.safeIndex(950, (d_125_v49_).length(0))
                (d_125_v49_)[index5_] = _dafny.CodePoint('e')
                index6_ = default__.safeIndex(931, (d_121_v45_).length(0))
                (d_121_v45_)[index6_] = p1
                d_126_v50_: _dafny.Array
                def lambda0_(d_127_p1_, d_128_v40_):
                    def lambda1_(d_129_i2_):
                        return _dafny.Map({d_127_p1_: (d_128_v40_).f8})

                    return lambda1_

                init0_ = lambda0_(p1, d_116_v40_)
                nw12_ = _dafny.Array(None, 18)
                for i0_0_ in range(nw12_.length(0)):
                    nw12_[i0_0_] = init0_(i0_0_)
                d_126_v50_ = nw12_
                index7_ = default__.safeIndex(514, (d_126_v50_).length(0))
                (d_126_v50_)[index7_] = (d_119_v43_) | (d_119_v43_)
                d_130_v51_: str
                d_130_v51_ = _dafny.CodePoint('v')
                d_131_v52_: D3
                d_131_v52_ = D3_DC6(d_123_v47_)
                index8_ = default__.safeIndex(781, (d_121_v45_).length(0))
                index9_ = default__.safeIndex(950, (d_125_v49_).length(0))
                index10_ = default__.safeIndex(931, (d_121_v45_).length(0))
                index11_ = default__.safeIndex(514, (d_126_v50_).length(0))
                rhs6_ = (d_110_v34_)[default__.safeIndex((d_116_v40_).f8, len(d_110_v34_))]
                rhs7_ = d_118_v42_
                rhs8_ = d_130_v51_
                rhs9_ = p2
                rhs10_ = default__.fm39((d_116_v40_).f8, d_131_v52_, p2, (d_115_v39_).f11, globalState)
                lhs4_ = d_121_v45_
                lhs5_ = default__.safeIndex(781, (d_121_v45_).length(0))
                lhs6_ = d_125_v49_
                lhs7_ = default__.safeIndex(950, (d_125_v49_).length(0))
                lhs8_ = d_121_v45_
                lhs9_ = default__.safeIndex(931, (d_121_v45_).length(0))
                lhs10_ = d_126_v50_
                lhs11_ = default__.safeIndex(514, (d_126_v50_).length(0))
                lhs4_[lhs5_] = rhs6_
                d_118_v42_ = rhs7_
                lhs6_[lhs7_] = rhs8_
                lhs8_[lhs9_] = rhs9_
                lhs10_[lhs11_] = rhs10_
            elif True:
                index12_ = default__.safeIndex(18, (d_111_v35_).length(0))
                (d_111_v35_)[index12_] = p0
                index13_ = default__.safeIndex(18, (d_111_v35_).length(0))
                (d_111_v35_)[index13_] = default__.safeModuloInt(len(d_56_v0_), p0)
                index14_ = default__.safeIndex(18, (d_111_v35_).length(0))
                (d_111_v35_)[index14_] = (d_111_v35_)[default__.safeIndex(18, (d_111_v35_).length(0))]
                index15_ = default__.safeIndex(18, (d_111_v35_).length(0))
                (d_111_v35_)[index15_] = p0
                d_132_v53_: _dafny.Map
                d_132_v53_ = _dafny.Map({p2: p2})
                d_133_v54_: str
                d_133_v54_ = _dafny.CodePoint('c')
                d_134_v55_: _dafny.Array
                nw13_ = _dafny.Array(None, 23)
                nw13_[int(0)] = p2
                nw13_[int(1)] = p2
                nw13_[int(2)] = p2
                nw13_[int(3)] = not(default__.fm24(p2, len(d_132_v53_), _dafny.SeqWithoutIsStrInference([d_133_v54_]), globalState))
                nw13_[int(4)] = not(p1)
                nw13_[int(5)] = p2
                nw13_[int(6)] = p1
                nw13_[int(7)] = p2
                nw13_[int(8)] = p2
                nw13_[int(9)] = p2
                nw13_[int(10)] = p1
                nw13_[int(11)] = p2
                nw13_[int(12)] = p2
                nw13_[int(13)] = p2
                nw13_[int(14)] = p1
                nw13_[int(15)] = p1
                nw13_[int(16)] = True
                nw13_[int(17)] = p1
                nw13_[int(18)] = p1
                nw13_[int(19)] = p1
                nw13_[int(20)] = p2
                nw13_[int(21)] = p1
                nw13_[int(22)] = True
                d_134_v55_ = nw13_
                d_135_v56_: _dafny.Seq
                d_135_v56_ = _dafny.SeqWithoutIsStrInference([d_134_v55_])
                d_136_v57_: _dafny.MultiSet
                d_136_v57_ = _dafny.MultiSet([(d_111_v35_)[default__.safeIndex(18, (d_111_v35_).length(0))], len(d_109_cf26_)])
                index16_ = default__.safeIndex(18, (d_111_v35_).length(0))
                (d_111_v35_)[index16_] = ((D4_DC9(p0, d_135_v56_, (d_111_v35_)[default__.safeIndex(18, (d_111_v35_).length(0))], ((d_136_v57_)[((d_136_v57_)[516] if (516) in (d_136_v57_) else (d_111_v35_)[default__.safeIndex(18, (d_111_v35_).length(0))])] if (((d_136_v57_)[516] if (516) in (d_136_v57_) else (d_111_v35_)[default__.safeIndex(18, (d_111_v35_).length(0))])) in (d_136_v57_) else 584))).cf15) + ((0) - ((d_111_v35_)[default__.safeIndex(18, (d_111_v35_).length(0))]))
                d_137_v58_: _dafny.Array
                nw14_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_137_v58_ = nw14_
                d_138_v59_: _dafny.Array
                nw15_ = _dafny.Array(int(0), 18)
                d_138_v59_ = nw15_
                index17_ = default__.safeIndex(22, (d_137_v58_).length(0))
                (d_137_v58_)[index17_] = d_138_v59_
                index18_ = default__.safeIndex(22, (d_137_v58_).length(0))
                (d_137_v58_)[index18_] = d_138_v59_
            d_139_v60_: C3
            nw16_ = C3()
            nw16_.ctor__(p1, p0, p0)
            d_139_v60_ = nw16_
            d_140_v61_: D8
            d_140_v61_ = D8_DC21(d_139_v60_)
            d_141_v62_: _dafny.Map
            d_141_v62_ = _dafny.Map({d_140_v61_: p2})
            d_141_v62_ = (d_141_v62_).set(D8_DC21(d_139_v60_), False)
        d_142_v63_: str
        d_142_v63_ = _dafny.CodePoint('s')
        d_143_v64_: _dafny.Array
        nw17_ = _dafny.Array(None, 9)
        nw17_[int(0)] = not(not(p2))
        nw17_[int(1)] = p2
        nw17_[int(2)] = p1
        nw17_[int(3)] = p2
        nw17_[int(4)] = p1
        nw17_[int(5)] = True
        nw17_[int(6)] = p1
        nw17_[int(7)] = False
        nw17_[int(8)] = p1
        d_143_v64_ = nw17_
        d_144_v65_: D6
        d_144_v65_ = D6_DC15(d_143_v64_)
        d_145_v66_: _dafny.Seq
        d_145_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbthqcwcx"))
        d_146_v67_: _dafny.Map
        d_146_v67_ = _dafny.Map({d_144_v65_: d_145_v66_})
        d_147_i3_: int
        d_147_i3_ = 0
        with _dafny.label("0"):
            while (d_142_v63_) in (((d_146_v67_)[d_144_v65_] if (d_144_v65_) in (d_146_v67_) else d_145_v66_)):
                with _dafny.c_label("0"):
                    if (d_147_i3_) >= (100):
                        raise _dafny.Break("0")
                    d_147_i3_ = (d_147_i3_) + (1)
                    d_148_v68_: T0
                    nw18_ = C5()
                    nw18_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")), p0, p0, p0)
                    d_148_v68_ = nw18_
                    d_149_v69_: T1
                    nw19_ = C5()
                    nw19_.ctor__(d_145_v66_, (d_148_v68_).f8, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ks"))), (d_148_v68_).f7)
                    d_149_v69_ = nw19_
                    d_150_v70_: D11
                    d_150_v70_ = D11_DC28(d_149_v69_)
                    d_151_v71_: _dafny.Array
                    nw20_ = _dafny.Array(None, 15)
                    nw20_[int(0)] = d_149_v69_
                    nw20_[int(1)] = d_149_v69_
                    nw20_[int(2)] = d_149_v69_
                    nw20_[int(3)] = (d_150_v70_).cf42
                    nw20_[int(4)] = d_149_v69_
                    nw20_[int(5)] = d_149_v69_
                    nw20_[int(6)] = d_149_v69_
                    nw20_[int(7)] = d_149_v69_
                    nw20_[int(8)] = d_149_v69_
                    nw20_[int(9)] = d_149_v69_
                    nw20_[int(10)] = d_149_v69_
                    nw20_[int(11)] = d_149_v69_
                    nw20_[int(12)] = d_149_v69_
                    nw20_[int(13)] = d_149_v69_
                    nw20_[int(14)] = d_149_v69_
                    d_151_v71_ = nw20_
                    d_152_v72_: _dafny.Map
                    d_152_v72_ = _dafny.Map({d_148_v68_: d_151_v71_})
                    d_153_v73_: D12
                    d_153_v73_ = D12_DC32(d_152_v72_)
                    d_154_v74_: _dafny.Seq
                    d_154_v74_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_148_v68_: d_151_v71_})])
                    d_155_v75_: _dafny.Map
                    d_155_v75_ = _dafny.Map({(d_148_v68_).f8: d_152_v72_})
                    d_156_v76_: _dafny.Array
                    nw21_ = _dafny.Array(None, 21)
                    nw21_[int(0)] = d_152_v72_
                    nw21_[int(1)] = ((d_153_v73_).cf45) | (d_152_v72_)
                    nw21_[int(2)] = (d_154_v74_)[default__.safeIndex(709, len(d_154_v74_))]
                    nw21_[int(3)] = _dafny.Map({d_148_v68_: d_151_v71_})
                    nw21_[int(4)] = d_152_v72_
                    nw21_[int(5)] = _dafny.Map({d_148_v68_: d_151_v71_})
                    nw21_[int(6)] = d_152_v72_
                    nw21_[int(7)] = d_152_v72_
                    nw21_[int(8)] = d_152_v72_
                    nw21_[int(9)] = (d_152_v72_).set(d_148_v68_, d_151_v71_)
                    nw21_[int(10)] = d_152_v72_
                    nw21_[int(11)] = _dafny.Map({d_148_v68_: d_151_v71_})
                    nw21_[int(12)] = d_152_v72_
                    nw21_[int(13)] = d_152_v72_
                    nw21_[int(14)] = (d_152_v72_) | (d_152_v72_)
                    nw21_[int(15)] = d_152_v72_
                    nw21_[int(16)] = d_152_v72_
                    nw21_[int(17)] = d_152_v72_
                    nw21_[int(18)] = (d_152_v72_).set(d_148_v68_, d_151_v71_)
                    nw21_[int(19)] = ((d_152_v72_).set(d_148_v68_, d_151_v71_)).set(d_148_v68_, d_151_v71_)
                    nw21_[int(20)] = ((d_155_v75_)[default__.fm16(p2, globalState)] if (default__.fm16(p2, globalState)) in (d_155_v75_) else d_152_v72_)
                    d_156_v76_ = nw21_
                    index19_ = default__.safeIndex(145, (d_156_v76_).length(0))
                    (d_156_v76_)[index19_] = d_152_v72_
                    index20_ = default__.safeIndex(145, (d_156_v76_).length(0))
                    (d_156_v76_)[index20_] = (d_152_v72_) | ((d_152_v72_) | (d_152_v72_))
                    d_157_v77_: _dafny.Map
                    d_157_v77_ = _dafny.Map({True: d_149_v69_})
                    d_157_v77_ = (d_157_v77_).set(default__.fm24(p2, len((d_145_v66_).set(default__.safeIndex((d_148_v68_).f7, len(d_145_v66_)), d_142_v63_)), d_145_v66_, globalState), d_149_v69_)
                    nw22_ = _dafny.Array(False, 23)
                    d_143_v64_ = nw22_
                    d_158_v78_: _dafny.Map
                    d_158_v78_ = _dafny.Map({d_148_v68_: d_149_v69_})
                    d_158_v78_ = _dafny.Map({d_148_v68_: d_149_v69_})
                    pass
            pass
        index21_ = default__.safeIndex(630, (d_143_v64_).length(0))
        (d_143_v64_)[index21_] = (d_145_v66_) <= (d_145_v66_)
        index22_ = default__.safeIndex(630, (d_143_v64_).length(0))
        (d_143_v64_)[index22_] = (p0) >= (p0)
        d_159_v79_: _dafny.Array
        nw23_ = _dafny.Array(_dafny.Array(None, 0), 3)
        d_159_v79_ = nw23_
        d_160_v80_: _dafny.Array
        nw24_ = _dafny.Array(int(0), 19)
        d_160_v80_ = nw24_
        index23_ = default__.safeIndex(944, (d_159_v79_).length(0))
        (d_159_v79_)[index23_] = d_160_v80_
        index24_ = default__.safeIndex(944, (d_159_v79_).length(0))
        (d_159_v79_)[index24_] = d_160_v80_
        if (d_143_v64_)[default__.safeIndex(630, (d_143_v64_).length(0))]:
            d_161_v81_: D7
            d_161_v81_ = D7_DC19(p0, (d_143_v64_)[default__.safeIndex(630, (d_143_v64_).length(0))], p0)
            d_162_v82_: _dafny.Seq
            d_162_v82_ = _dafny.SeqWithoutIsStrInference([d_161_v81_])
            d_163_v83_: _dafny.MultiSet
            d_163_v83_ = _dafny.MultiSet([p0, p0, len(d_162_v82_), p0, len(d_145_v66_)])
            d_163_v83_ = d_163_v83_
            arr0_ = (d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]
            index25_ = default__.safeIndex(980, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))
            arr0_[index25_] = 237
            d_164_v84_: _dafny.MultiSet
            d_164_v84_ = _dafny.MultiSet([p2, (d_143_v64_)[default__.safeIndex(630, (d_143_v64_).length(0))]])
            d_165_v85_: C1
            nw25_ = C1()
            nw25_.ctor__(p0, p0)
            d_165_v85_ = nw25_
            d_166_v86_: _dafny.Map
            d_166_v86_ = _dafny.Map({d_165_v85_: p1})
            d_167_v87_: C1
            d_167_v87_ = d_165_v85_
            d_168_v88_: D9
            d_168_v88_ = D9_DC23((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))])
            d_169_v89_: _dafny.Map
            d_169_v89_ = _dafny.Map({p0: (d_168_v88_).cf36})
            d_170_v90_: _dafny.Array
            nw26_ = _dafny.Array(D9.default()(), 18)
            d_170_v90_ = nw26_
            d_171_v91_: _dafny.Map
            d_171_v91_ = _dafny.Map({p1: d_170_v90_})
            index26_ = default__.safeIndex(630, (d_143_v64_).length(0))
            index27_ = default__.safeIndex(944, (d_159_v79_).length(0))
            arr1_ = (d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]
            index28_ = default__.safeIndex(980, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))
            rhs11_ = (d_164_v84_).isdisjoint(d_164_v84_)
            rhs12_ = ((d_166_v86_)[(d_167_v87_)] if ((d_167_v87_)) in (d_166_v86_) else True)
            rhs13_ = ((d_169_v89_)[len((_dafny.Map({p2: d_170_v90_})) | (d_171_v91_))] if (len((_dafny.Map({p2: d_170_v90_})) | (d_171_v91_))) in (d_169_v89_) else ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))] if p1 else d_160_v80_))
            rhs14_ = False
            rhs15_ = (p0) - (-541)
            lhs12_ = d_143_v64_
            lhs13_ = default__.safeIndex(630, (d_143_v64_).length(0))
            lhs14_ = globalState
            lhs15_ = d_159_v79_
            lhs16_ = default__.safeIndex(944, (d_159_v79_).length(0))
            lhs17_ = globalState
            lhs18_ = (d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]
            lhs19_ = default__.safeIndex(980, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))
            lhs12_[lhs13_] = rhs11_
            lhs14_.f3 = rhs12_
            lhs15_[lhs16_] = rhs13_
            lhs17_.f4 = rhs14_
            lhs18_[lhs19_] = rhs15_
            (globalState).f4 = True
            d_172_v92_: _dafny.Map
            d_172_v92_ = _dafny.Map({((d_56_v0_).set(default__.safeIndex(-735, len(d_56_v0_)), p0)) + (_dafny.SeqWithoutIsStrInference([((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))])[default__.safeIndex(980, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))]])): p0})
            d_172_v92_ = ((d_172_v92_) | (d_172_v92_)) | (d_172_v92_)
            d_173_v93_: _dafny.Array
            nw27_ = _dafny.Array(_dafny.Seq({}), 5)
            d_173_v93_ = nw27_
            d_173_v93_ = (d_173_v93_ if True else d_173_v93_)
        elif True:
            d_174_v94_: int
            d_174_v94_ = -106
            d_175_v95_: _dafny.MultiSet
            d_175_v95_ = _dafny.MultiSet([p2])
            rhs16_ = False
            rhs17_ = d_174_v94_
            rhs18_ = ((0) - (d_174_v94_)) == (p0)
            rhs19_ = (default__.safeModuloInt(265, ((d_175_v95_)[p2] if (p2) in (d_175_v95_) else (0) - (d_174_v94_)))) >= (p0)
            lhs20_ = globalState
            lhs21_ = globalState
            lhs22_ = globalState
            lhs20_.f4 = rhs16_
            d_174_v94_ = rhs17_
            lhs21_.f4 = rhs18_
            lhs22_.f3 = rhs19_
            (globalState).f4 = p2
            d_176_v96_: _dafny.Map
            d_176_v96_ = _dafny.Map({(len(d_145_v66_)) != (p0): _dafny.MultiSet(d_56_v0_)})
            d_176_v96_ = (d_176_v96_).set(p2, _dafny.MultiSet(d_56_v0_))
            arr2_ = (d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]
            index29_ = default__.safeIndex(869, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))
            arr2_[index29_] = p0
            arr3_ = (d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]
            index30_ = default__.safeIndex(869, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))
            arr3_[index30_] = (default__.fm16(p1, globalState)) - (p0)
            d_177_v97_: D1
            d_177_v97_ = D1_DC1(_dafny.CodePoint('d'))
            source5_ = D1_DC3(d_177_v97_)
            if source5_.is_DC2:
                nw28_ = _dafny.Array(None, 7)
                nw28_[int(0)] = d_145_v66_
                nw28_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfn"))
                nw28_[int(2)] = (d_145_v66_ if not(p1) else d_145_v66_)
                nw28_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilvyw"))
                nw28_[int(4)] = d_145_v66_
                nw28_[int(5)] = d_145_v66_
                nw28_[int(6)] = d_145_v66_
                r0 = nw28_
                d_174_v94_ = d_174_v94_
                d_178_v98_: D3
                d_178_v98_ = D3_DC7(p1)
                d_178_v98_ = d_178_v98_
                d_179_v99_: _dafny.Array
                nw29_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_179_v99_ = nw29_
                arr4_ = (d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]
                index31_ = default__.safeIndex(869, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))
                rhs20_ = d_179_v99_
                rhs21_ = ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))])[default__.safeIndex(869, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))]
                lhs23_ = (d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]
                lhs24_ = default__.safeIndex(869, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))
                d_179_v99_ = rhs20_
                lhs23_[lhs24_] = rhs21_
            elif source5_.is_DC1:
                d_180___mcc_h8_ = source5_.cf1
                d_181_cf1_ = d_180___mcc_h8_
                (globalState).f3 = (d_181_cf1_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lq")))
                d_182_v100_: T1
                nw30_ = C5()
                nw30_.ctor__(d_145_v66_, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))])[default__.safeIndex(869, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))], p0, 780)
                d_182_v100_ = nw30_
                d_183_v101_: _dafny.Map
                d_183_v101_ = _dafny.Map({d_182_v100_: p0})
                d_184_v102_: _dafny.Map
                d_184_v102_ = _dafny.Map({(d_183_v101_) | (d_183_v101_): p1})
                d_184_v102_ = (d_184_v102_).set(d_183_v101_, p2)
                index32_ = default__.safeIndex(630, (d_143_v64_).length(0))
                rhs22_ = (d_143_v64_)[default__.safeIndex(630, (d_143_v64_).length(0))]
                rhs23_ = d_56_v0_
                lhs25_ = d_143_v64_
                lhs26_ = default__.safeIndex(630, (d_143_v64_).length(0))
                lhs25_[lhs26_] = rhs22_
                d_56_v0_ = rhs23_
                d_185_v103_: _dafny.Array
                nw31_ = _dafny.Array(_dafny.Seq({}), 10)
                d_185_v103_ = nw31_
                d_186_v104_: _dafny.Seq
                d_186_v104_ = _dafny.SeqWithoutIsStrInference([d_143_v64_, d_143_v64_, d_143_v64_])
                index33_ = default__.safeIndex(655, (d_185_v103_).length(0))
                (d_185_v103_)[index33_] = d_186_v104_
                index34_ = default__.safeIndex(655, (d_185_v103_).length(0))
                (d_185_v103_)[index34_] = ((d_186_v104_) + (d_186_v104_)) + (d_186_v104_)
            elif True:
                d_187___mcc_h9_ = source5_.cf2
                d_188_cf2_ = d_187___mcc_h9_
                d_189_v105_: C6
                nw32_ = C6()
                nw32_.ctor__(len(d_56_v0_), p0)
                d_189_v105_ = nw32_
                d_190_v106_: _dafny.Seq
                d_190_v106_ = _dafny.SeqWithoutIsStrInference([d_189_v105_, d_189_v105_])
                d_191_v107_: _dafny.Map
                d_191_v107_ = _dafny.Map({len(d_190_v106_): d_174_v94_})
                d_192_v108_: C4
                nw33_ = C4()
                nw33_.ctor__(p2, d_174_v94_, d_174_v94_)
                d_192_v108_ = nw33_
                d_193_v109_: _dafny.Map
                d_193_v109_ = _dafny.Map({d_192_v108_: ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))])[default__.safeIndex(869, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))]})
                d_194_v110_: C4
                nw34_ = C4()
                nw34_.ctor__(p1, p0, ((d_193_v109_)[d_192_v108_] if (d_192_v108_) in (d_193_v109_) else 325))
                d_194_v110_ = nw34_
                d_174_v94_ = (default__.safeDivisionInt(((d_191_v107_)[len(_dafny.Map({(d_143_v64_)[default__.safeIndex(630, (d_143_v64_).length(0))]: d_194_v110_}))] if (len(_dafny.Map({(d_143_v64_)[default__.safeIndex(630, (d_143_v64_).length(0))]: d_194_v110_}))) in (d_191_v107_) else len(d_56_v0_)), d_174_v94_)) + (((d_191_v107_)[64] if (64) in (d_191_v107_) else -306))
                d_195_v111_: C2
                nw35_ = C2()
                nw35_.ctor__(((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))])[default__.safeIndex(869, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))], ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))])[default__.safeIndex(869, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))])
                d_195_v111_ = nw35_
                arr5_ = (d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]
                index35_ = default__.safeIndex(869, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))
                rhs24_ = d_143_v64_
                rhs25_ = ((d_145_v66_) + (_dafny.SeqWithoutIsStrInference([d_142_v63_ for d_196_i4_ in range(default__.abs(616))]))) == (d_145_v66_)
                rhs26_ = d_174_v94_
                lhs27_ = globalState
                lhs28_ = (d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]
                lhs29_ = default__.safeIndex(869, ((d_159_v79_)[default__.safeIndex(944, (d_159_v79_).length(0))]).length(0))
                d_143_v64_ = rhs24_
                lhs27_.f3 = rhs25_
                lhs28_[lhs29_] = rhs26_
                d_197_v112_: _dafny.Array
                nw36_ = _dafny.Array(None, 19)
                d_197_v112_ = nw36_
                d_198_v113_: C3
                nw37_ = C3()
                nw37_.ctor__(p1, p0, p0)
                d_198_v113_ = nw37_
                index36_ = default__.safeIndex(994, (d_197_v112_).length(0))
                (d_197_v112_)[index36_] = d_198_v113_
                index37_ = default__.safeIndex(994, (d_197_v112_).length(0))
                (d_197_v112_)[index37_] = d_198_v113_
        d_199_v114_: _dafny.Map
        d_199_v114_ = _dafny.Map({_dafny.MultiSet([d_145_v66_]): not((d_143_v64_)[default__.safeIndex(630, (d_143_v64_).length(0))])})
        d_200_v115_: _dafny.MultiSet
        d_200_v115_ = _dafny.MultiSet([d_145_v66_])
        d_201_v116_: C0
        nw38_ = C0()
        nw38_.ctor__(d_160_v80_, p0, p0)
        d_201_v116_ = nw38_
        d_202_v117_: _dafny.Map
        d_202_v117_ = _dafny.Map({p0: d_201_v116_})
        d_203_v118_: _dafny.Map
        d_203_v118_ = _dafny.Map({p0: d_202_v117_})
        d_199_v114_ = (d_199_v114_).set(d_200_v115_, (p0) in (d_203_v118_))
        d_204_v119_: _dafny.Array
        nw39_ = _dafny.Array(None, 9)
        nw39_[int(0)] = d_145_v66_
        nw39_[int(1)] = d_145_v66_
        nw39_[int(2)] = d_145_v66_
        nw39_[int(3)] = d_145_v66_
        nw39_[int(4)] = d_145_v66_
        nw39_[int(5)] = d_145_v66_
        nw39_[int(6)] = d_145_v66_
        nw39_[int(7)] = d_145_v66_
        nw39_[int(8)] = d_145_v66_
        d_204_v119_ = nw39_
        d_205_v120_: D14
        d_205_v120_ = D14_DC35(d_204_v119_)
        r0 = (d_205_v120_).cf48
        d_206_v121_: D5
        d_206_v121_ = D5_DC13(p2, d_142_v63_)
        pat_let_tv0_ = d_145_v66_
        pat_let_tv1_ = p2
        def lambda2_(source6_):
            if source6_.is_DC12:
                d_207___mcc_h10_ = source6_.cf19
                d_208_cf19_ = d_207___mcc_h10_
                return (pat_let_tv0_) <= ((D7_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))).cf26)
            elif source6_.is_DC13:
                d_209___mcc_h11_ = source6_.cf20
                d_210___mcc_h12_ = source6_.cf21
                d_211_cf21_ = d_210___mcc_h12_
                d_212_cf20_ = d_209___mcc_h11_
                return d_212_cf20_
            elif source6_.is_DC11:
                d_213___mcc_h13_ = source6_.cf18
                d_214_cf18_ = d_213___mcc_h13_
                return pat_let_tv1_
            elif True:
                d_215___mcc_h14_ = source6_.cf22
                d_216_cf22_ = d_215___mcc_h14_
                return True

        r1 = lambda2_(d_206_v121_)
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_217_v0_: _dafny.Array
        def lambda3_(d_218_i1_):
            return False

        init1_ = lambda3_
        nw40_ = _dafny.Array(None, 18)
        for i0_1_ in range(nw40_.length(0)):
            nw40_[i0_1_] = init1_(i0_1_)
        d_217_v0_ = nw40_
        d_219_v1_: _dafny.Map
        d_219_v1_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_220_i0_ in range(default__.abs(230))]): d_217_v0_})
        d_221_v2_: _dafny.Seq
        d_221_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
        d_222_globalState_: GlobalState
        nw41_ = GlobalState()
        nw41_.ctor__(124, True, (d_219_v1_) | (d_219_v1_), False, True, 611, d_221_v2_)
        d_222_globalState_ = nw41_
        d_223_v3_: bool
        d_223_v3_ = False
        (d_222_globalState_).f3 = d_223_v3_
        d_224_v4_: int
        d_224_v4_ = 756
        d_225_i2_: int
        d_225_i2_ = 0
        with _dafny.label("1"):
            while (default__.safeModuloInt(d_224_v4_, d_224_v4_)) <= (d_224_v4_):
                with _dafny.c_label("1"):
                    if (d_225_i2_) >= (100):
                        raise _dafny.Break("1")
                    d_225_i2_ = (d_225_i2_) + (1)
                    d_226_v5_: _dafny.Array
                    d_227_v6_: bool
                    out0_: _dafny.Array
                    out1_: bool
                    out0_, out1_ = default__.m0(d_224_v4_, d_223_v3_, d_223_v3_, d_222_globalState_)
                    d_226_v5_ = out0_
                    d_227_v6_ = out1_
                    d_228_v7_: _dafny.Map
                    d_228_v7_ = _dafny.Map({True: d_227_v6_})
                    d_229_v8_: _dafny.Seq
                    d_229_v8_ = _dafny.SeqWithoutIsStrInference([d_223_v3_])
                    d_228_v7_ = (d_228_v7_).set((d_223_v3_) in (d_229_v8_), d_223_v3_)
                    d_230_v9_: _dafny.Set
                    d_230_v9_ = _dafny.Set({d_221_v2_})
                    d_231_v10_: C2
                    nw42_ = C2()
                    nw42_.ctor__((len(d_230_v9_)) - (d_224_v4_), (0) - (d_224_v4_))
                    d_231_v10_ = nw42_
                    d_228_v7_ = (d_228_v7_).set(d_223_v3_, d_223_v3_)
                    pass
            pass
        d_232_v11_: _dafny.Seq
        d_232_v11_ = _dafny.SeqWithoutIsStrInference([88])
        (d_222_globalState_).f4 = ((d_232_v11_)[default__.safeIndex((0) - (d_224_v4_), len(d_232_v11_))]) == (len(d_221_v2_))
        d_233_v12_: _dafny.Map
        d_233_v12_ = _dafny.Map({823: False})
        d_234_v13_: C4
        nw43_ = C4()
        nw43_.ctor__(d_223_v3_, len(_dafny.Map({d_223_v3_: d_223_v3_})), d_224_v4_)
        d_234_v13_ = nw43_
        d_235_v14_: _dafny.Seq
        d_235_v14_ = _dafny.SeqWithoutIsStrInference([d_234_v13_, d_234_v13_])
        d_236_v15_: _dafny.Array
        nw44_ = _dafny.Array(None, 23)
        nw44_[int(0)] = 218
        nw44_[int(1)] = len(d_221_v2_)
        nw44_[int(2)] = d_224_v4_
        nw44_[int(3)] = len(d_233_v12_)
        nw44_[int(4)] = d_224_v4_
        nw44_[int(5)] = 709
        nw44_[int(6)] = d_224_v4_
        nw44_[int(7)] = len(d_235_v14_)
        nw44_[int(8)] = d_224_v4_
        nw44_[int(9)] = d_224_v4_
        nw44_[int(10)] = d_224_v4_
        nw44_[int(11)] = len(_dafny.SeqWithoutIsStrInference([(d_234_v13_).f11, (d_234_v13_).f11]))
        nw44_[int(12)] = d_224_v4_
        nw44_[int(13)] = (0) - (d_224_v4_)
        nw44_[int(14)] = 0
        nw44_[int(15)] = d_224_v4_
        nw44_[int(16)] = d_224_v4_
        nw44_[int(17)] = d_224_v4_
        nw44_[int(18)] = d_224_v4_
        nw44_[int(19)] = d_224_v4_
        nw44_[int(20)] = d_224_v4_
        nw44_[int(21)] = d_224_v4_
        nw44_[int(22)] = d_224_v4_
        d_236_v15_ = nw44_
        d_237_v16_: _dafny.Seq
        d_237_v16_ = _dafny.SeqWithoutIsStrInference([d_236_v15_])
        d_238_v17_: C0
        nw45_ = C0()
        nw45_.ctor__((d_237_v16_)[default__.safeIndex(d_224_v4_, len(d_237_v16_))], 525, d_224_v4_)
        d_238_v17_ = nw45_
        index38_ = default__.safeIndex(913, (d_217_v0_).length(0))
        (d_217_v0_)[index38_] = not(not(not ((d_234_v13_).f11) or ((d_234_v13_).f11)))
        d_239_v18_: _dafny.Array
        nw46_ = _dafny.Array(None, 1)
        nw46_[int(0)] = d_236_v15_
        d_239_v18_ = nw46_
        index39_ = default__.safeIndex(677, (d_239_v18_).length(0))
        (d_239_v18_)[index39_] = d_238_v17_.f13
        d_240_v19_: _dafny.Seq
        d_240_v19_ = _dafny.SeqWithoutIsStrInference([(d_234_v13_).f11, (default__.fm8(True, d_232_v11_, d_223_v3_, d_222_globalState_)) == (d_221_v2_), (d_223_v3_ if not((d_234_v13_).f11) else (d_234_v13_).f11), (d_233_v12_) != (d_233_v12_), ((d_234_v13_).f11) == (d_223_v3_)])
        d_241_v20_: str
        d_241_v20_ = _dafny.CodePoint('b')
        d_242_v21_: D1
        d_242_v21_ = D1_DC1(d_241_v20_)
        d_243_v22_: D1
        d_243_v22_ = D1_DC3(D1_DC1(d_241_v20_))
        d_244_v23_: _dafny.MultiSet
        d_244_v23_ = _dafny.MultiSet([D1_DC3(d_242_v21_), d_243_v22_])
        d_245_v24_: _dafny.Set
        d_245_v24_ = _dafny.Set({d_244_v23_, d_244_v23_, _dafny.MultiSet([d_243_v22_])})
        index40_ = default__.safeIndex(913, (d_217_v0_).length(0))
        index41_ = default__.safeIndex(677, (d_239_v18_).length(0))
        rhs27_ = not((d_223_v3_ if (d_224_v4_) >= (len(d_221_v2_)) else d_223_v3_))
        rhs28_ = d_236_v15_
        rhs29_ = d_240_v19_
        rhs30_ = default__.fm20(d_223_v3_, (d_245_v24_).ispropersubset(default__.fm36(d_223_v3_, (d_234_v13_).f11, d_222_globalState_)), (d_234_v13_).f11, d_222_globalState_)
        rhs31_ = d_234_v13_
        lhs30_ = d_217_v0_
        lhs31_ = default__.safeIndex(913, (d_217_v0_).length(0))
        lhs32_ = d_239_v18_
        lhs33_ = default__.safeIndex(677, (d_239_v18_).length(0))
        lhs30_[lhs31_] = rhs27_
        lhs32_[lhs33_] = rhs28_
        d_240_v19_ = rhs29_
        d_224_v4_ = rhs30_
        d_234_v13_ = rhs31_
        d_246_v25_: _dafny.MultiSet
        d_246_v25_ = _dafny.MultiSet([d_238_v17_.f13])
        d_247_v26_: _dafny.Set
        d_247_v26_ = _dafny.Set({(d_234_v13_).f11})
        d_248_v27_: _dafny.Map
        d_248_v27_ = _dafny.Map({(d_234_v13_).f11: (d_238_v17_).fm1(d_247_v26_, d_224_v4_, d_224_v4_, d_224_v4_, d_222_globalState_)})
        d_249_v28_: _dafny.Map
        d_249_v28_ = _dafny.Map({(0) - (d_224_v4_): (d_246_v25_).set(d_238_v17_.f13, default__.abs(d_224_v4_))})
        rhs32_ = d_217_v0_
        rhs33_ = (len((d_247_v26_) - (d_247_v26_)) if ((d_248_v27_)[(d_234_v13_).f11] if ((d_234_v13_).f11) in (d_248_v27_) else (d_234_v13_).f11) else d_224_v4_)
        rhs34_ = ((d_249_v28_)[(0) - (d_224_v4_)] if ((0) - (d_224_v4_)) in (d_249_v28_) else d_246_v25_)
        d_217_v0_ = rhs32_
        d_224_v4_ = rhs33_
        d_246_v25_ = rhs34_
        d_250_v29_: T2
        nw47_ = C2()
        nw47_.ctor__(d_224_v4_, d_224_v4_)
        d_250_v29_ = nw47_
        d_251_v30_: _dafny.Map
        d_251_v30_ = _dafny.Map({d_250_v29_: d_233_v12_})
        d_252_v31_: C2
        nw48_ = C2()
        nw48_.ctor__(d_224_v4_, (356) + (len(((d_251_v30_)[d_250_v29_] if (d_250_v29_) in (d_251_v30_) else _dafny.Map({len(_dafny.Map({(d_250_v29_).f8: d_238_v17_})): (d_217_v0_)[default__.safeIndex(913, (d_217_v0_).length(0))]})))))
        d_252_v31_ = nw48_
        d_253_i3_: int
        d_253_i3_ = 0
        with _dafny.label("2"):
            while ((d_221_v2_) + (d_221_v2_)) < (_dafny.SeqWithoutIsStrInference([d_241_v20_ for d_261_i4_ in range(default__.abs(683))])):
                with _dafny.c_label("2"):
                    if (d_253_i3_) >= (100):
                        raise _dafny.Break("2")
                    d_253_i3_ = (d_253_i3_) + (1)
                    index42_ = default__.safeIndex(913, (d_217_v0_).length(0))
                    (d_217_v0_)[index42_] = not((d_234_v13_).f11)
                    d_254_v32_: _dafny.Array
                    nw49_ = _dafny.Array(_dafny.Seq({}), 17)
                    d_254_v32_ = nw49_
                    d_255_v33_: _dafny.Map
                    d_255_v33_ = _dafny.Map({_dafny.MultiSet([(d_217_v0_)[default__.safeIndex(913, (d_217_v0_).length(0))]]): (d_250_v29_).f7})
                    d_256_v34_: D6
                    d_256_v34_ = D6_DC16(d_255_v33_, d_240_v19_)
                    index43_ = default__.safeIndex(35, (d_254_v32_).length(0))
                    (d_254_v32_)[index43_] = (d_256_v34_).cf25
                    index44_ = default__.safeIndex(35, (d_254_v32_).length(0))
                    (d_254_v32_)[index44_] = d_240_v19_
                    d_257_v35_: D7
                    d_257_v35_ = D7_DC20(d_241_v20_, d_224_v4_)
                    d_258_v36_: _dafny.Map
                    d_258_v36_ = _dafny.Map({d_257_v35_: d_241_v20_})
                    d_259_v37_: C4
                    nw50_ = C4()
                    nw50_.ctor__(d_223_v3_, (0) - (len(_dafny.SeqWithoutIsStrInference([d_258_v36_, d_258_v36_]))), (d_250_v29_).f7)
                    d_259_v37_ = nw50_
                    d_221_v2_ = (((d_252_v31_).fm14(d_222_globalState_)) + (d_221_v2_) if (d_217_v0_)[default__.safeIndex(913, (d_217_v0_).length(0))] else _dafny.SeqWithoutIsStrInference([d_241_v20_ for d_260_i5_ in range(default__.abs(220))]))
                    pass
            pass
        d_262_v38_: _dafny.Array
        def lambda4_(d_263_v0_, d_264_v13_):
            def lambda5_(d_265_i6_):
                return _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_263_v0_)[default__.safeIndex(913, (d_263_v0_).length(0))], (d_264_v13_).f11})])

            return lambda5_

        init2_ = lambda4_(d_217_v0_, d_234_v13_)
        nw51_ = _dafny.Array(None, 5)
        for i0_2_ in range(nw51_.length(0)):
            nw51_[i0_2_] = init2_(i0_2_)
        d_262_v38_ = nw51_
        d_266_v39_: _dafny.Seq
        d_266_v39_ = _dafny.SeqWithoutIsStrInference([d_247_v26_])
        index45_ = default__.safeIndex(708, (d_262_v38_).length(0))
        (d_262_v38_)[index45_] = d_266_v39_
        d_267_v40_: D4
        d_267_v40_ = D4_DC10(d_224_v4_, (d_234_v13_).f11)
        index46_ = default__.safeIndex(913, (d_217_v0_).length(0))
        index47_ = default__.safeIndex(708, (d_262_v38_).length(0))
        rhs35_ = (d_267_v40_).cf17
        rhs36_ = (d_266_v39_) + (d_266_v39_)
        lhs34_ = d_217_v0_
        lhs35_ = default__.safeIndex(913, (d_217_v0_).length(0))
        lhs36_ = d_262_v38_
        lhs37_ = default__.safeIndex(708, (d_262_v38_).length(0))
        lhs34_[lhs35_] = rhs35_
        lhs36_[lhs37_] = rhs36_
        d_268_v41_: C5
        nw52_ = C5()
        nw52_.ctor__(d_221_v2_, d_224_v4_, (d_250_v29_).f8, (d_250_v29_).f7)
        d_268_v41_ = nw52_
        d_269_v42_: _dafny.Map
        d_269_v42_ = _dafny.Map({d_268_v41_: (d_268_v41_).f10})
        d_270_v43_: _dafny.MultiSet
        d_270_v43_ = _dafny.MultiSet([d_217_v0_, d_217_v0_])
        (d_222_globalState_).f4 = (d_234_v13_).fm1(d_247_v26_, (d_250_v29_).f8, ((d_269_v42_)[d_268_v41_] if (d_268_v41_) in (d_269_v42_) else d_224_v4_), ((d_270_v43_)[d_217_v0_] if (d_217_v0_) in (d_270_v43_) else len(d_232_v11_)), d_222_globalState_)
        (d_222_globalState_).f3 = not((len(default__.fm9(d_223_v3_, d_222_globalState_))) > (((0) - (len(((d_268_v41_).f9).set(default__.safeIndex(len(d_247_v26_), len((d_268_v41_).f9)), d_241_v20_)))) * ((d_250_v29_).f7)))
        hi0_ = d_224_v4_
        for d_271_i7_ in range(((d_250_v29_).f7) - (len(d_232_v11_)), hi0_):
            d_224_v4_ = d_224_v4_
            d_272_v44_: _dafny.Seq
            d_272_v44_ = _dafny.SeqWithoutIsStrInference([d_250_v29_, d_250_v29_])
            d_273_v45_: C2
            nw53_ = C2()
            nw53_.ctor__((d_250_v29_).f7, len((d_272_v44_) + (d_272_v44_)))
            d_273_v45_ = nw53_
            d_273_v45_ = d_273_v45_
            d_273_v45_ = d_273_v45_
            (d_273_v45_).m1(d_221_v2_, d_222_globalState_)
        d_274_v46_: _dafny.Array
        nw54_ = _dafny.Array(D8.default()(), 9)
        d_274_v46_ = nw54_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_274_v46_).length(0)):
            d_275_i8_: int = guard_loop_0_
            if (True) and (((0) <= (d_275_i8_)) and ((d_275_i8_) < ((d_274_v46_).length(0)))):
                (d_274_v46_)[(d_275_i8_)] = D8_DC22((0) - ((d_268_v41_).f10))
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_238_v17_.f13).length(0)):
            d_276_i9_: int = guard_loop_1_
            if (True) and (((0) <= (d_276_i9_)) and ((d_276_i9_) < ((d_238_v17_.f13).length(0)))):
                arr6_ = d_238_v17_.f13
                arr6_[(d_276_i9_)] = default__.safeDivisionInt(d_276_i9_, (d_250_v29_).f7)
        (d_222_globalState_).f3 = not(d_223_v3_)
        d_221_v2_ = d_221_v2_
        _dafny.print(_dafny.string_of((d_217_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v0_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_219_v1_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_221_v2_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_222_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_222_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_222_globalState_).f2)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_222_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_222_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_222_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_222_globalState_).f6).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_223_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_224_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_225_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_v11_) == (_dafny.SeqWithoutIsStrInference([88]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v12_) == (_dafny.Map({823: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_234_v13_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_235_v14_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_237_v16_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_.f13)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v18_)[0])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v19_) == (_dafny.SeqWithoutIsStrInference([False, False, False, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_241_v20_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_242_v21_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_243_v22_).cf2).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_244_v23_) == (_dafny.MultiSet([D1_DC3(D1_DC1(_dafny.CodePoint('b'))), D1_DC3(D1_DC1(_dafny.CodePoint('b')))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v24_) == (_dafny.Set({_dafny.MultiSet([D1_DC3(D1_DC1(_dafny.CodePoint('b'))), D1_DC3(D1_DC1(_dafny.CodePoint('b')))]), _dafny.MultiSet([D1_DC3(D1_DC1(_dafny.CodePoint('b')))])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_v25_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v26_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_248_v27_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_249_v28_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_250_v29_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_250_v29_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_251_v30_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_253_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v38_)[0]) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True, False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v38_)[1]) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True, False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v38_)[2]) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True, False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v38_)[3]) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({False}), _dafny.Set({False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v38_)[4]) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True, False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_266_v39_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_267_v40_).cf16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_267_v40_).cf17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_268_v41_).f9).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_268_v41_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_269_v42_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v43_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v46_)[0]).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v46_)[1]).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v46_)[2]).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v46_)[3]).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v46_)[4]).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v46_)[5]).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v46_)[6]).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v46_)[7]).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v46_)[8]).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: False
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

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
        return lambda: D1_DC2()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC2(D1, NamedTuple('DC2', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2)
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
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
        return lambda: D2_DC5(False, int(0), int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)

class D2_DC5(D2, NamedTuple('DC5', [('cf4', Any), ('cf5', Any), ('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC7(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)

class D3_DC7(D3, NamedTuple('DC7', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC6(D3, NamedTuple('DC6', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC9(int(0), _dafny.Seq({}), int(0), int(0))
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

class D4_DC9(D4, NamedTuple('DC9', [('cf12', Any), ('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC8(D4, NamedTuple('DC8', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC12(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC12(D5, NamedTuple('DC12', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC11(D5, NamedTuple('DC11', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(_dafny.Map({}), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)

class D6_DC16(D6, NamedTuple('DC16', [('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC18(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)

class D7_DC18(D7, NamedTuple('DC18', [('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf29', Any), ('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({self.cf26.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC22(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)

class D8_DC22(D8, NamedTuple('DC22', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC24(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)

class D9_DC24(D9, NamedTuple('DC24', [('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC23(D9, NamedTuple('DC23', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC26(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D10_DC25)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)

class D10_DC26(D10, NamedTuple('DC26', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC25(D10, NamedTuple('DC25', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC25({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC25) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC27(D10, NamedTuple('DC27', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC29(_dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D11_DC28)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)

class D11_DC29(D11, NamedTuple('DC29', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC28(D11, NamedTuple('DC28', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC31(D11, NamedTuple('DC31', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC33(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)

class D12_DC33(D12, NamedTuple('DC33', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({self.cf46.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC32(D12, NamedTuple('DC32', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)

class D13_DC34(D13, NamedTuple('DC34', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC36(False, _dafny.Seq({}), _dafny.Seq({}), _dafny.Seq({}), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D14_DC35)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D14_DC37)

class D14_DC36(D14, NamedTuple('DC36', [('cf49', Any), ('cf50', Any), ('cf51', Any), ('cf52', Any), ('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC35(D14, NamedTuple('DC35', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC35({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC35) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC37(D14, NamedTuple('DC37', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC37({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC37) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass

class T2:
    pass
    def m1(self, p0, globalState):
        pass

    def m2(self, p0, p1, p2, p3, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f3: bool = False
        self.f4: bool = False
        self._f0: int = int(0)
        self._f1: bool = False
        self._f2: _dafny.Map = _dafny.Map({})
        self._f5: int = int(0)
        self._f6: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6

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
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6

class C0(T1, T0):
    def  __init__(self):
        self._f7: int = int(0)
        self._f8: int = int(0)
        self.f13: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    def ctor__(self, f13, f7, f8):
        (self).f13 = f13
        (self)._f7 = f7
        (self)._f8 = f8

    def fm0(self, p0, globalState):
        return _dafny.Set({(_dafny.Set({False})) | (_dafny.Set({False})), (_dafny.Set({False})) - (_dafny.Set({True, not(True)})), _dafny.Set({True, False, False})})

    def fm1(self, p0, p1, p2, p3, globalState):
        def iife23_():
            coll23_ = _dafny.Set()
            compr_23_: _dafny.Seq
            for compr_23_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_277_i3_ in range(default__.abs(539))]) for d_278_i2_ in range(default__.abs(620))])).Elements:
                d_279_v0_: _dafny.Seq = compr_23_
                if (d_279_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_277_i3_ in range(default__.abs(539))]) for d_278_i2_ in range(default__.abs(620))])):
                    coll23_ = coll23_.union(_dafny.Set([d_279_v0_]))
            return _dafny.Set(coll23_)
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nup"))})) | (iife23_()
        )).issubset((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jg"))})) | (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbglft")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_280_i0_ in range(default__.abs(739))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_281_i1_ in range(default__.abs(-819))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkesmfsv"))})))


class C1(T0):
    def  __init__(self):
        self._f7: int = int(0)
        self._f8: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    def ctor__(self, f7, f8):
        (self)._f7 = f7
        (self)._f8 = f8

    def fm18(self, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jof"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hno"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_282_i0_ in range(default__.abs(20))])))

    def fm19(self, p0, p1, p2, p3, globalState):
        return (D1_DC1(_dafny.CodePoint('m'))).cf1

    def m6(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: int = int(0)
        r2: int = int(0)
        d_283_v0_: _dafny.Array
        def lambda6_(d_284_i0_):
            return ((self).f7) > ((self).f7)

        init3_ = lambda6_
        nw55_ = _dafny.Array(None, 6)
        for i0_3_ in range(nw55_.length(0)):
            nw55_[i0_3_] = init3_(i0_3_)
        d_283_v0_ = nw55_
        d_285_v1_: bool
        d_285_v1_ = False
        index48_ = default__.safeIndex(464, (d_283_v0_).length(0))
        (d_283_v0_)[index48_] = d_285_v1_
        d_286_v2_: bool
        d_286_v2_ = d_285_v1_
        d_287_v3_: _dafny.Seq
        d_287_v3_ = _dafny.SeqWithoutIsStrInference([d_285_v1_, d_285_v1_, (d_285_v1_ if d_285_v1_ else d_285_v1_)])
        d_288_v4_: _dafny.Set
        d_288_v4_ = _dafny.Set({309})
        d_289_v5_: D2
        d_289_v5_ = D2_DC5(d_285_v1_, (self).f7, len(d_288_v4_), (self).f7, d_285_v1_)
        pat_let_tv2_ = d_286_v2_
        pat_let_tv3_ = d_286_v2_
        index49_ = default__.safeIndex(464, (d_283_v0_).length(0))
        def lambda7_(source7_):
            if source7_.is_DC5:
                d_290___mcc_h0_ = source7_.cf4
                d_291___mcc_h1_ = source7_.cf5
                d_292___mcc_h2_ = source7_.cf6
                d_293___mcc_h3_ = source7_.cf7
                d_294___mcc_h4_ = source7_.cf8
                d_295_cf8_ = d_294___mcc_h4_
                d_296_cf7_ = d_293___mcc_h3_
                d_297_cf6_ = d_292___mcc_h2_
                d_298_cf5_ = d_291___mcc_h1_
                d_299_cf4_ = d_290___mcc_h0_
                return pat_let_tv2_
            elif True:
                d_300___mcc_h5_ = source7_.cf3
                d_301_cf3_ = d_300___mcc_h5_
                return pat_let_tv3_

        rhs37_ = d_285_v1_
        rhs38_ = (d_287_v3_)[default__.safeIndex((self).f8, len(d_287_v3_))]
        rhs39_ = lambda7_(d_289_v5_)
        rhs40_ = (d_286_v2_)
        lhs38_ = globalState
        lhs39_ = d_283_v0_
        lhs40_ = default__.safeIndex(464, (d_283_v0_).length(0))
        lhs41_ = globalState
        lhs38_.f4 = rhs37_
        lhs39_[lhs40_] = rhs38_
        d_286_v2_ = rhs39_
        lhs41_.f4 = rhs40_
        d_302_v6_: _dafny.Map
        d_302_v6_ = _dafny.Map({(self).f7: (self).f7})
        d_303_v7_: _dafny.Seq
        d_303_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tuyj"))
        hi1_ = 422
        for d_304_i1_ in range(((d_302_v6_)[default__.fm20(d_285_v1_, (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], True, globalState)] if (default__.fm20(d_285_v1_, (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], True, globalState)) in (d_302_v6_) else len(d_303_v7_)), hi1_):
            d_305_v8_: _dafny.Array
            nw56_ = _dafny.Array(int(0), 2)
            d_305_v8_ = nw56_
            index50_ = default__.safeIndex(310, (d_305_v8_).length(0))
            (d_305_v8_)[index50_] = (self).f7
            d_306_v9_: _dafny.Array
            def lambda8_(d_307_v6_):
                def lambda9_(d_308_i2_):
                    return _dafny.SeqWithoutIsStrInference([(0) - ((self).f7), (self).f7, ((d_307_v6_)[(self).f8] if ((self).f8) in (d_307_v6_) else (self).f8)])

                return lambda9_

            init4_ = lambda8_(d_302_v6_)
            nw57_ = _dafny.Array(None, 27)
            for i0_4_ in range(nw57_.length(0)):
                nw57_[i0_4_] = init4_(i0_4_)
            d_306_v9_ = nw57_
            index51_ = default__.safeIndex(310, (d_305_v8_).length(0))
            rhs41_ = (self).f8
            rhs42_ = (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]
            rhs43_ = d_306_v9_
            rhs44_ = (d_289_v5_).cf7
            lhs42_ = d_305_v8_
            lhs43_ = default__.safeIndex(310, (d_305_v8_).length(0))
            lhs42_[lhs43_] = rhs41_
            d_285_v1_ = rhs42_
            d_306_v9_ = rhs43_
            r2 = rhs44_
            d_309_v10_: C0
            nw58_ = C0()
            nw58_.ctor__(d_305_v8_, d_304_i1_, d_304_i1_)
            d_309_v10_ = nw58_
            d_309_v10_ = d_309_v10_
            d_310_v11_: str
            d_310_v11_ = _dafny.CodePoint('y')
            d_310_v11_ = d_310_v11_
            d_311_v12_: _dafny.Map
            d_311_v12_ = _dafny.Map({_dafny.CodePoint('w'): d_285_v1_})
            d_312_v13_: _dafny.Map
            d_312_v13_ = _dafny.Map({d_311_v12_: (self).f8})
            if (d_311_v12_) in (d_312_v13_):
                (globalState).f3 = (d_287_v3_)[default__.safeIndex((d_289_v5_).cf5, len(d_287_v3_))]
                index52_ = default__.safeIndex(310, (d_305_v8_).length(0))
                rhs45_ = d_285_v1_
                rhs46_ = (d_305_v8_)[default__.safeIndex(310, (d_305_v8_).length(0))]
                lhs44_ = globalState
                lhs45_ = d_305_v8_
                lhs46_ = default__.safeIndex(310, (d_305_v8_).length(0))
                lhs44_.f4 = rhs45_
                lhs45_[lhs46_] = rhs46_
                d_313_v14_: _dafny.Map
                d_313_v14_ = _dafny.Map({d_286_v2_: (d_305_v8_)[default__.safeIndex(310, (d_305_v8_).length(0))]})
                d_314_v15_: _dafny.Set
                d_314_v15_ = _dafny.Set({d_313_v14_})
                d_315_v16_: _dafny.Seq
                d_315_v16_ = _dafny.SeqWithoutIsStrInference([d_288_v4_])
                index53_ = default__.safeIndex(310, (d_305_v8_).length(0))
                index54_ = default__.safeIndex(464, (d_283_v0_).length(0))
                rhs47_ = (d_305_v8_)[default__.safeIndex(310, (d_305_v8_).length(0))]
                rhs48_ = d_303_v7_
                rhs49_ = d_314_v15_
                rhs50_ = d_285_v1_
                rhs51_ = (d_315_v16_)[default__.safeIndex(((self).f8) + (310), len(d_315_v16_))]
                lhs47_ = d_305_v8_
                lhs48_ = default__.safeIndex(310, (d_305_v8_).length(0))
                lhs49_ = d_283_v0_
                lhs50_ = default__.safeIndex(464, (d_283_v0_).length(0))
                lhs47_[lhs48_] = rhs47_
                d_303_v7_ = rhs48_
                d_314_v15_ = rhs49_
                lhs49_[lhs50_] = rhs50_
                d_288_v4_ = rhs51_
                (globalState).f3 = (d_304_i1_) <= ((self).f8)
                d_316_v17_: _dafny.Map
                d_316_v17_ = _dafny.Map({(d_304_i1_) >= ((self).f8): (d_305_v8_)[default__.safeIndex(310, (d_305_v8_).length(0))]})
                d_316_v17_ = d_316_v17_
            elif True:
                d_303_v7_ = (d_303_v7_) + (d_303_v7_)
                d_317_v18_: D1
                d_317_v18_ = D1_DC1(d_310_v11_)
                d_318_v19_: _dafny.Seq
                d_318_v19_ = _dafny.SeqWithoutIsStrInference([d_317_v18_, d_317_v18_, D1_DC1(_dafny.CodePoint('c'))])
                d_319_v20_: C0
                nw59_ = C0()
                nw59_.ctor__(d_305_v8_, len(default__.fm21(715, len(d_318_v19_), (self).f7, d_303_v7_, globalState)), (self).f8)
                d_319_v20_ = nw59_
                d_320_v21_: _dafny.Map
                d_320_v21_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_289_v5_]): d_285_v1_})
                d_321_v22_: _dafny.Seq
                d_321_v22_ = _dafny.SeqWithoutIsStrInference([d_289_v5_])
                d_320_v21_ = (d_320_v21_).set(d_321_v22_, False)
                index55_ = default__.safeIndex(310, (d_305_v8_).length(0))
                (d_305_v8_)[index55_] = default__.safeDivisionInt((d_305_v8_)[default__.safeIndex(310, (d_305_v8_).length(0))], default__.safeDivisionInt(default__.fm20((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], (d_285_v1_), (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], globalState), 225))
                d_322_v23_: D1
                d_322_v23_ = D1_DC1(d_310_v11_)
                d_323_v24_: D1
                d_323_v24_ = D1_DC3(d_322_v23_)
                d_324_v25_: _dafny.MultiSet
                d_324_v25_ = _dafny.MultiSet([d_323_v24_, D1_DC3(d_322_v23_), d_323_v24_, d_323_v24_])
                d_324_v25_ = default__.fm22(globalState)
        source8_ = D2_DC5((d_287_v3_)[default__.safeIndex((self).f7, len(d_287_v3_))], (self).f7, len(d_288_v4_), (self).f7, (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))])
        if source8_.is_DC5:
            d_325___mcc_h6_ = source8_.cf4
            d_326___mcc_h7_ = source8_.cf5
            d_327___mcc_h8_ = source8_.cf6
            d_328___mcc_h9_ = source8_.cf7
            d_329___mcc_h10_ = source8_.cf8
            d_330_cf8_ = d_329___mcc_h10_
            d_331_cf7_ = d_328___mcc_h9_
            d_332_cf6_ = d_327___mcc_h8_
            d_333_cf5_ = d_326___mcc_h7_
            d_334_cf4_ = d_325___mcc_h6_
            d_335_v26_: str
            d_335_v26_ = _dafny.CodePoint('q')
            d_336_v27_: _dafny.Map
            d_336_v27_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbuccmlho"))).set(default__.safeIndex((self).f8, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbuccmlho")))), d_335_v26_): (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]})
            d_337_v28_: _dafny.Map
            d_337_v28_ = _dafny.Map({(d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]: (d_336_v27_) != (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_335_v26_ for d_338_i3_ in range(default__.abs(787))]): d_334_cf4_}))})
            d_337_v28_ = (d_337_v28_).set(d_334_cf4_, d_330_cf8_)
            d_339_v29_: D1
            d_339_v29_ = D1_DC3(default__.fm23(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qf"))), globalState))
            d_340_v30_: D1
            d_340_v30_ = D1_DC3(d_339_v29_)
            d_341_v31_: _dafny.Seq
            d_341_v31_ = _dafny.SeqWithoutIsStrInference([d_333_cf5_])
            index56_ = default__.safeIndex(464, (d_283_v0_).length(0))
            rhs52_ = d_340_v30_
            rhs53_ = d_285_v1_
            rhs54_ = d_334_cf4_
            rhs55_ = ((_dafny.SeqWithoutIsStrInference([d_335_v26_ for d_342_i4_ in range(default__.abs(-405))])) + ((_dafny.SeqWithoutIsStrInference([d_335_v26_ for d_343_i5_ in range(default__.abs(602))])) + (d_303_v7_))).set(default__.safeIndex((d_331_cf7_) - (len(d_341_v31_)), len((_dafny.SeqWithoutIsStrInference([d_335_v26_ for d_344_i4_ in range(default__.abs(-405))])) + ((_dafny.SeqWithoutIsStrInference([d_335_v26_ for d_345_i5_ in range(default__.abs(602))])) + (d_303_v7_)))), d_335_v26_)
            lhs51_ = d_283_v0_
            lhs52_ = default__.safeIndex(464, (d_283_v0_).length(0))
            d_340_v30_ = rhs52_
            d_334_cf4_ = rhs53_
            lhs51_[lhs52_] = rhs54_
            d_303_v7_ = rhs55_
            d_334_cf4_ = default__.fm24((d_331_cf7_) < ((d_289_v5_).cf5), d_333_cf5_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c')]), globalState)
            index57_ = default__.safeIndex(464, (d_283_v0_).length(0))
            rhs56_ = (self).f8
            rhs57_ = not (d_334_cf4_) or (d_334_cf4_)
            rhs58_ = not(d_285_v1_)
            lhs53_ = d_283_v0_
            lhs54_ = default__.safeIndex(464, (d_283_v0_).length(0))
            d_333_cf5_ = rhs56_
            lhs53_[lhs54_] = rhs57_
            d_330_cf8_ = rhs58_
        elif True:
            d_346___mcc_h11_ = source8_.cf3
            d_347_cf3_ = d_346___mcc_h11_
            d_348_v32_: _dafny.Array
            def lambda10_(d_349_i6_):
                return (d_349_i6_) * ((self).f7)

            init5_ = lambda10_
            nw60_ = _dafny.Array(None, 6)
            for i0_5_ in range(nw60_.length(0)):
                nw60_[i0_5_] = init5_(i0_5_)
            d_348_v32_ = nw60_
            index58_ = default__.safeIndex(161, (d_348_v32_).length(0))
            (d_348_v32_)[index58_] = (self).f8
            d_350_v33_: _dafny.Set
            d_350_v33_ = _dafny.Set({(d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]})
            d_351_v34_: C0
            nw61_ = C0()
            nw61_.ctor__(d_348_v32_, len(d_303_v7_), (self).f7)
            d_351_v34_ = nw61_
            d_352_v35_: _dafny.Map
            d_352_v35_ = _dafny.Map({d_285_v1_: (d_351_v34_ if False else d_351_v34_)})
            index59_ = default__.safeIndex(464, (d_283_v0_).length(0))
            index60_ = default__.safeIndex(161, (d_348_v32_).length(0))
            rhs59_ = not(not((not(((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]) or (not((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))])))) and (d_285_v1_)))
            rhs60_ = ((d_350_v33_).intersection(d_350_v33_) if d_285_v1_ else d_350_v33_)
            rhs61_ = (0) - ((0) - (len(d_352_v35_)))
            rhs62_ = ((self).f7) + ((self).f7)
            lhs55_ = d_283_v0_
            lhs56_ = default__.safeIndex(464, (d_283_v0_).length(0))
            lhs57_ = d_348_v32_
            lhs58_ = default__.safeIndex(161, (d_348_v32_).length(0))
            lhs55_[lhs56_] = rhs59_
            r0 = rhs60_
            lhs57_[lhs58_] = rhs61_
            r2 = rhs62_
            if ((0) - ((self).f7)) != (((self).f7) * ((0) - ((self).f8))):
                index61_ = default__.safeIndex(161, (d_348_v32_).length(0))
                (d_348_v32_)[index61_] = 117
                d_353_v36_: _dafny.Seq
                d_353_v36_ = _dafny.SeqWithoutIsStrInference([d_347_cf3_])
                d_354_v37_: _dafny.Array
                nw62_ = _dafny.Array(None, 11)
                nw62_[int(0)] = ((d_347_cf3_).set((d_287_v3_)[default__.safeIndex((self).f8, len(d_287_v3_))], default__.abs((d_348_v32_)[default__.safeIndex(161, (d_348_v32_).length(0))]))) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]), d_285_v1_])))
                nw62_[int(1)] = d_347_cf3_
                nw62_[int(2)] = d_347_cf3_
                nw62_[int(3)] = _dafny.MultiSet(d_287_v3_)
                nw62_[int(4)] = d_347_cf3_
                nw62_[int(5)] = d_347_cf3_
                nw62_[int(6)] = d_347_cf3_
                nw62_[int(7)] = (d_347_cf3_) | (d_347_cf3_)
                nw62_[int(8)] = ((d_353_v36_)[default__.safeIndex((0) - ((self).f8), len(d_353_v36_))]) | (d_347_cf3_)
                nw62_[int(9)] = d_347_cf3_
                nw62_[int(10)] = d_347_cf3_
                d_354_v37_ = nw62_
                index62_ = default__.safeIndex(908, (d_354_v37_).length(0))
                (d_354_v37_)[index62_] = d_347_cf3_
                index63_ = default__.safeIndex(908, (d_354_v37_).length(0))
                (d_354_v37_)[index63_] = ((d_347_cf3_).intersection(_dafny.MultiSet([d_285_v1_]))) - (d_347_cf3_)
                index64_ = default__.safeIndex(161, (d_348_v32_).length(0))
                (d_348_v32_)[index64_] = ((self).f8) * (((0) - ((d_348_v32_)[default__.safeIndex(161, (d_348_v32_).length(0))]) if d_285_v1_ else (d_348_v32_)[default__.safeIndex(161, (d_348_v32_).length(0))]))
                d_355_v38_: str
                d_355_v38_ = _dafny.CodePoint('u')
                d_356_v39_: _dafny.Map
                d_356_v39_ = _dafny.Map({d_355_v38_: (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]})
                d_356_v39_ = (d_356_v39_).set(d_355_v38_, (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))])
                index65_ = default__.safeIndex(161, (d_348_v32_).length(0))
                (d_348_v32_)[index65_] = (d_348_v32_)[default__.safeIndex(161, (d_348_v32_).length(0))]
            elif True:
                r1 = (((0) - ((self).f8) if d_285_v1_ else (self).f8) if (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))] else (d_348_v32_)[default__.safeIndex(161, (d_348_v32_).length(0))])
                d_357_v40_: _dafny.Seq
                d_357_v40_ = _dafny.SeqWithoutIsStrInference([D1_DC2()])
                d_357_v40_ = d_357_v40_
                d_358_v41_: D2
                d_358_v41_ = D2_DC4(d_347_cf3_)
                d_358_v41_ = d_358_v41_
                d_359_v42_: C0
                nw63_ = C0()
                nw63_.ctor__(d_348_v32_, default__.safeDivisionInt(-810, (d_348_v32_)[default__.safeIndex(161, (d_348_v32_).length(0))]), ((d_348_v32_)[default__.safeIndex(161, (d_348_v32_).length(0))]) + (len(d_303_v7_)))
                d_359_v42_ = nw63_
                d_360_v43_: _dafny.Array
                nw64_ = _dafny.Array(None, 14)
                d_360_v43_ = nw64_
                index66_ = default__.safeIndex(650, (d_360_v43_).length(0))
                (d_360_v43_)[index66_] = d_351_v34_
                d_361_v44_: _dafny.Map
                d_361_v44_ = _dafny.Map({(self).f8: d_285_v1_})
                d_362_v45_: _dafny.MultiSet
                d_362_v45_ = _dafny.MultiSet([default__.fm20((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], d_285_v1_, ((d_361_v44_)[(self).f7] if ((self).f7) in (d_361_v44_) else (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]), globalState), len(d_350_v33_), (self).f8])
                index67_ = default__.safeIndex(650, (d_360_v43_).length(0))
                rhs63_ = d_359_v42_
                rhs64_ = (d_362_v45_).cardinality
                lhs59_ = d_360_v43_
                lhs60_ = default__.safeIndex(650, (d_360_v43_).length(0))
                lhs59_[lhs60_] = rhs63_
                r1 = rhs64_
            index68_ = default__.safeIndex(464, (d_283_v0_).length(0))
            (d_283_v0_)[index68_] = (d_287_v3_)[default__.safeIndex(((d_348_v32_)[default__.safeIndex(161, (d_348_v32_).length(0))] if d_285_v1_ else (d_348_v32_)[default__.safeIndex(161, (d_348_v32_).length(0))]), len(d_287_v3_))]
            d_363_v46_: _dafny.Set
            d_363_v46_ = _dafny.Set({d_286_v2_})
            d_364_v47_: _dafny.Map
            d_364_v47_ = _dafny.Map({d_363_v46_: D2_DC5((d_351_v34_).fm1(d_350_v33_, (d_348_v32_)[default__.safeIndex(161, (d_348_v32_).length(0))], (d_348_v32_)[default__.safeIndex(161, (d_348_v32_).length(0))], (self).f7, globalState), default__.fm20(d_285_v1_, (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], globalState), (self).f8, -418, d_285_v1_)})
            index69_ = default__.safeIndex(161, (d_348_v32_).length(0))
            (d_348_v32_)[index69_] = len(d_364_v47_)
        d_365_v48_: str
        d_365_v48_ = _dafny.CodePoint('x')
        d_366_v49_: D1
        d_366_v49_ = D1_DC1(d_365_v48_)
        d_367_v50_: D1
        d_367_v50_ = D1_DC3(d_366_v49_)
        source9_ = d_367_v50_
        if source9_.is_DC2:
            d_368_v51_: _dafny.Set
            d_368_v51_ = _dafny.Set({d_283_v0_})
            if ((d_368_v51_).issubset(d_368_v51_)) and (d_285_v1_):
                (globalState).f4 = d_285_v1_
                d_369_v52_: D1
                d_369_v52_ = D1_DC1((d_365_v48_ if True else (self).fm19(d_285_v1_, (d_289_v5_).cf8, (self).f8, (self).f7, globalState)))
                d_369_v52_ = d_369_v52_
                d_370_v53_: _dafny.Array
                nw65_ = _dafny.Array(int(0), 9)
                d_370_v53_ = nw65_
                d_371_v54_: _dafny.Map
                d_371_v54_ = _dafny.Map({(self).f7: d_283_v0_})
                d_372_v55_: C0
                nw66_ = C0()
                nw66_.ctor__(d_370_v53_, (0) - ((self).f7), len(d_371_v54_))
                d_372_v55_ = nw66_
                d_373_v56_: _dafny.Map
                d_373_v56_ = _dafny.Map({(d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]: (299) <= (default__.fm20(d_285_v1_, (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], globalState))})
                d_373_v56_ = default__.fm25(globalState)
                r2 = default__.safeDivisionInt(((self).f8 if d_285_v1_ else (self).f7), (self).f8)
            elif True:
                d_374_v57_: _dafny.Map
                d_374_v57_ = _dafny.Map({((self).f8) * (default__.fm20(d_285_v1_, d_285_v1_, (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], globalState)): (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]})
                d_374_v57_ = (d_374_v57_).set((self).f7, (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))])
                d_375_v58_: _dafny.Array
                nw67_ = _dafny.Array(int(0), 26)
                d_375_v58_ = nw67_
                index70_ = default__.safeIndex(299, (d_375_v58_).length(0))
                (d_375_v58_)[index70_] = (self).f8
                index71_ = default__.safeIndex(299, (d_375_v58_).length(0))
                (d_375_v58_)[index71_] = (self).f7
                d_376_v59_: _dafny.Seq
                d_376_v59_ = _dafny.SeqWithoutIsStrInference([(d_375_v58_)[default__.safeIndex(299, (d_375_v58_).length(0))]])
                d_377_v60_: _dafny.MultiSet
                d_377_v60_ = _dafny.MultiSet([(d_375_v58_)[default__.safeIndex(299, (d_375_v58_).length(0))]])
                d_378_v61_: _dafny.Map
                d_378_v61_ = _dafny.Map({(d_376_v59_).set(default__.safeIndex(default__.fm20(default__.fm24(d_285_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trxrrblv"))), _dafny.SeqWithoutIsStrInference([d_365_v48_, d_365_v48_]), globalState), (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], d_285_v1_, globalState), len(d_376_v59_)), (self).f8): (d_288_v4_).ispropersubset((default__.fm26(d_377_v60_, (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], globalState)).cf9)})
                d_378_v61_ = d_378_v61_
                r2 = (self).f8
                d_285_v1_ = (d_287_v3_)[default__.safeIndex((d_375_v58_)[default__.safeIndex(299, (d_375_v58_).length(0))], len(d_287_v3_))]
            r1 = (d_289_v5_).cf5
            (globalState).f4 = not(d_285_v1_)
            d_379_v62_: _dafny.Seq
            d_379_v62_ = _dafny.SeqWithoutIsStrInference([(self).f7, (self).f8])
            r2 = (-476) - ((d_379_v62_)[default__.safeIndex(default__.fm20((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], d_285_v1_, False, globalState), len(d_379_v62_))])
        elif source9_.is_DC1:
            d_380___mcc_h12_ = source9_.cf1
            d_381_cf1_ = d_380___mcc_h12_
            d_382_v63_: _dafny.Array
            def lambda11_(d_383_i7_):
                return (d_383_i7_) - ((self).f7)

            init6_ = lambda11_
            nw68_ = _dafny.Array(None, 27)
            for i0_6_ in range(nw68_.length(0)):
                nw68_[i0_6_] = init6_(i0_6_)
            d_382_v63_ = nw68_
            d_384_v64_: _dafny.Map
            d_384_v64_ = _dafny.Map({d_285_v1_: (0) - ((self).f7)})
            index72_ = default__.safeIndex(300, (d_382_v63_).length(0))
            (d_382_v63_)[index72_] = (len(d_384_v64_)) * ((d_289_v5_).cf7)
            index73_ = default__.safeIndex(300, (d_382_v63_).length(0))
            (d_382_v63_)[index73_] = -885
            d_385_v65_: _dafny.MultiSet
            d_385_v65_ = _dafny.MultiSet([(d_382_v63_)[default__.safeIndex(300, (d_382_v63_).length(0))], (self).f7])
            (globalState).f4 = (d_385_v65_) == (d_385_v65_)
            d_386_v66_: D3
            d_386_v66_ = D3_DC7((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))])
            d_387_v67_: _dafny.Map
            d_387_v67_ = _dafny.Map({d_285_v1_: d_386_v66_})
            d_387_v67_ = (d_387_v67_).set(((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))] if d_285_v1_ else (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]), d_386_v66_)
            if (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]:
                (globalState).f4 = (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]
                d_381_cf1_ = d_381_cf1_
                (globalState).f4 = ((957) - (350)) == ((self).f8)
                d_388_v68_: _dafny.Array
                nw69_ = _dafny.Array(None, 17)
                nw69_[int(0)] = d_382_v63_
                nw69_[int(1)] = d_382_v63_
                nw69_[int(2)] = d_382_v63_
                nw69_[int(3)] = d_382_v63_
                nw69_[int(4)] = d_382_v63_
                nw69_[int(5)] = d_382_v63_
                nw69_[int(6)] = d_382_v63_
                nw69_[int(7)] = d_382_v63_
                nw69_[int(8)] = d_382_v63_
                nw69_[int(9)] = d_382_v63_
                nw69_[int(10)] = d_382_v63_
                nw69_[int(11)] = d_382_v63_
                nw69_[int(12)] = d_382_v63_
                nw69_[int(13)] = d_382_v63_
                nw69_[int(14)] = d_382_v63_
                nw69_[int(15)] = d_382_v63_
                nw69_[int(16)] = d_382_v63_
                d_388_v68_ = nw69_
                d_389_v69_: _dafny.Seq
                d_389_v69_ = _dafny.SeqWithoutIsStrInference([d_388_v68_, d_388_v68_])
                d_390_v70_: _dafny.Map
                d_390_v70_ = _dafny.Map({(d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]: d_285_v1_})
                d_391_v71_: _dafny.Array
                nw70_ = _dafny.Array(None, 4)
                nw70_[int(0)] = d_388_v68_
                nw70_[int(1)] = (d_389_v69_)[default__.safeIndex((self).f8, len(d_389_v69_))]
                nw70_[int(2)] = d_388_v68_
                nw70_[int(3)] = (d_389_v69_)[default__.safeIndex(len(d_390_v70_), len(d_389_v69_))]
                d_391_v71_ = nw70_
                index74_ = default__.safeIndex(830, (d_391_v71_).length(0))
                (d_391_v71_)[index74_] = d_388_v68_
                d_392_v72_: D4
                d_392_v72_ = D4_DC8(d_388_v68_)
                index75_ = default__.safeIndex(830, (d_391_v71_).length(0))
                (d_391_v71_)[index75_] = (d_392_v72_).cf11
                d_393_v73_: _dafny.Map
                d_393_v73_ = _dafny.Map({(d_382_v63_)[default__.safeIndex(300, (d_382_v63_).length(0))]: True})
                index76_ = default__.safeIndex(300, (d_382_v63_).length(0))
                (d_382_v63_)[index76_] = default__.safeModuloInt((len(_dafny.SeqWithoutIsStrInference([d_285_v1_]))) - (default__.fm20(d_285_v1_, d_285_v1_, (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], globalState)), default__.safeDivisionInt((self).f8, len(d_393_v73_)))
            elif True:
                d_394_v74_: _dafny.Array
                nw71_ = _dafny.Array(None, 7)
                nw71_[int(0)] = d_303_v7_
                nw71_[int(1)] = d_303_v7_
                nw71_[int(2)] = (self).fm18(globalState)
                nw71_[int(3)] = d_303_v7_
                nw71_[int(4)] = d_303_v7_
                nw71_[int(5)] = (self).fm18(globalState)
                nw71_[int(6)] = d_303_v7_
                d_394_v74_ = nw71_
                d_395_v75_: _dafny.Map
                d_395_v75_ = _dafny.Map({d_285_v1_: d_394_v74_})
                d_395_v75_ = (d_395_v75_).set(d_285_v1_, d_394_v74_)
                d_396_v76_: _dafny.Seq
                d_396_v76_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_287_v3_)])
                d_397_v78_: _dafny.Map
                d_397_v78_ = _dafny.Map({d_285_v1_: (d_289_v5_).cf8})
                d_398_v79_: _dafny.MultiSet
                d_398_v79_ = _dafny.MultiSet([((d_397_v78_)[(d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]] if ((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]) in (d_397_v78_) else (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))])])
                d_399_v80_: _dafny.Set
                d_399_v80_ = _dafny.Set({d_398_v79_})
                index77_ = default__.safeIndex(464, (d_283_v0_).length(0))
                def iife24_():
                    coll24_ = _dafny.Set()
                    compr_24_: _dafny.MultiSet
                    for compr_24_ in (d_396_v76_).Elements:
                        d_400_v77_: _dafny.MultiSet = compr_24_
                        if (d_400_v77_) in (d_396_v76_):
                            coll24_ = coll24_.union(_dafny.Set([d_400_v77_]))
                    return _dafny.Set(coll24_)
                (d_283_v0_)[index77_] = ((iife24_()
                 if default__.fm24(d_285_v1_, (self).f8, d_303_v7_, globalState) else default__.fm27(globalState))) != (d_399_v80_)
                (globalState).f3 = ((d_287_v3_) + (d_287_v3_)) != (d_287_v3_)
                rhs65_ = (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]
                rhs66_ = (self).f8
                lhs61_ = globalState
                lhs61_.f4 = rhs65_
                r1 = rhs66_
                d_302_v6_ = ((d_302_v6_).set((d_382_v63_)[default__.safeIndex(300, (d_382_v63_).length(0))], (0) - ((self).f7))) | (d_302_v6_)
        elif True:
            d_401___mcc_h13_ = source9_.cf2
            d_402_cf2_ = d_401___mcc_h13_
            d_403_v81_: _dafny.MultiSet
            d_403_v81_ = _dafny.MultiSet([(d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]])
            d_404_v82_: _dafny.Seq
            d_404_v82_ = _dafny.SeqWithoutIsStrInference([d_403_v81_, d_403_v81_, _dafny.MultiSet([not(False), False])])
            d_405_v83_: _dafny.MultiSet
            d_405_v83_ = _dafny.MultiSet([(self).f7])
            d_406_v84_: _dafny.Array
            nw72_ = _dafny.Array(None, 24)
            nw72_[int(0)] = len(_dafny.SeqWithoutIsStrInference([(self).f7 for d_407_i8_ in range(default__.abs(970))]))
            nw72_[int(1)] = (self).f8
            nw72_[int(2)] = (self).f7
            nw72_[int(3)] = (self).f7
            nw72_[int(4)] = (self).f8
            nw72_[int(5)] = 487
            nw72_[int(6)] = (self).f8
            nw72_[int(7)] = (self).f8
            nw72_[int(8)] = len(d_303_v7_)
            nw72_[int(9)] = (self).f7
            nw72_[int(10)] = (self).f7
            nw72_[int(11)] = 912
            nw72_[int(12)] = len(_dafny.Set({(self).f8, (d_289_v5_).cf5, default__.fm20(d_285_v1_, d_285_v1_, (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], globalState)}))
            nw72_[int(13)] = (self).f7
            nw72_[int(14)] = len((d_404_v82_).set(default__.safeIndex((self).f8, len(d_404_v82_)), d_403_v81_))
            nw72_[int(15)] = (self).f7
            nw72_[int(16)] = len(d_302_v6_)
            nw72_[int(17)] = (self).f7
            nw72_[int(18)] = len((d_303_v7_).set(default__.safeIndex((self).f8, len(d_303_v7_)), d_365_v48_))
            nw72_[int(19)] = (d_405_v83_).cardinality
            nw72_[int(20)] = 609
            nw72_[int(21)] = (self).f8
            nw72_[int(22)] = (self).f8
            nw72_[int(23)] = (self).f8
            d_406_v84_ = nw72_
            d_408_v85_: _dafny.Map
            d_408_v85_ = _dafny.Map({d_285_v1_: _dafny.CodePoint('l')})
            d_409_v86_: C0
            nw73_ = C0()
            nw73_.ctor__(d_406_v84_, ((d_405_v83_)[(self).f7] if ((self).f7) in (d_405_v83_) else len(d_408_v85_)), (self).f8)
            d_409_v86_ = nw73_
            index78_ = default__.safeIndex(464, (d_283_v0_).length(0))
            index79_ = default__.safeIndex(464, (d_283_v0_).length(0))
            rhs67_ = default__.fm24((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], (self).f8, d_303_v7_, globalState)
            rhs68_ = d_285_v1_
            rhs69_ = (default__.fm20((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], True, globalState)) == ((self).f7)
            rhs70_ = _dafny.SeqWithoutIsStrInference([d_365_v48_ for d_410_i9_ in range(default__.abs(90))])
            lhs62_ = d_283_v0_
            lhs63_ = default__.safeIndex(464, (d_283_v0_).length(0))
            lhs64_ = d_283_v0_
            lhs65_ = default__.safeIndex(464, (d_283_v0_).length(0))
            lhs66_ = globalState
            lhs62_[lhs63_] = rhs67_
            lhs64_[lhs65_] = rhs68_
            lhs66_.f3 = rhs69_
            d_303_v7_ = rhs70_
            d_411_v87_: _dafny.Array
            def lambda12_(d_412_v1_, d_413_v0_, d_414_v83_):
                def lambda13_(d_415_i10_):
                    return _dafny.SeqWithoutIsStrInference([len(_dafny.Set({D1_DC2(), D1_DC2(), D1_DC2(), D1_DC2(), D1_DC2()})), (self).f8, len(_dafny.SeqWithoutIsStrInference([d_412_v1_, (d_413_v0_)[default__.safeIndex(464, (d_413_v0_).length(0))], d_412_v1_])), (self).f8, (d_414_v83_).cardinality])

                return lambda13_

            init7_ = lambda12_(d_285_v1_, d_283_v0_, d_405_v83_)
            nw74_ = _dafny.Array(None, 25)
            for i0_7_ in range(nw74_.length(0)):
                nw74_[i0_7_] = init7_(i0_7_)
            d_411_v87_ = nw74_
            d_416_v88_: _dafny.Seq
            d_416_v88_ = _dafny.SeqWithoutIsStrInference([d_283_v0_])
            index80_ = default__.safeIndex(487, (d_411_v87_).length(0))
            (d_411_v87_)[index80_] = default__.fm21(default__.fm20((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], True, globalState), (self).f7, len((d_416_v88_).set(default__.safeIndex((self).f7, len(d_416_v88_)), d_283_v0_)), d_303_v7_, globalState)
            d_417_v89_: _dafny.Map
            d_417_v89_ = _dafny.Map({(d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]: (self).f8})
            d_418_v90_: _dafny.Seq
            d_418_v90_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_417_v89_))])
            index81_ = default__.safeIndex(487, (d_411_v87_).length(0))
            (d_411_v87_)[index81_] = (_dafny.SeqWithoutIsStrInference([len(d_287_v3_), (self).f7])) + (d_418_v90_)
            if d_285_v1_:
                d_419_v91_: D2
                d_419_v91_ = D2_DC4((default__.fm28((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))], globalState)) - (_dafny.MultiSet([(d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]])))
                d_419_v91_ = d_419_v91_
                arr7_ = d_409_v86_.f13
                index82_ = default__.safeIndex(799, (d_409_v86_.f13).length(0))
                arr7_[index82_] = (0) - ((self).f8)
                arr8_ = d_409_v86_.f13
                index83_ = default__.safeIndex(799, (d_409_v86_.f13).length(0))
                arr8_[index83_] = len(_dafny.Set({d_285_v1_, (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]}))
                d_420_v92_: _dafny.MultiSet
                d_420_v92_ = _dafny.MultiSet([d_302_v6_, d_302_v6_, d_302_v6_])
                index84_ = default__.safeIndex(464, (d_283_v0_).length(0))
                (d_283_v0_)[index84_] = (d_420_v92_).ispropersubset(d_420_v92_)
                d_403_v81_ = d_403_v81_
                d_421_v93_: _dafny.Map
                d_421_v93_ = _dafny.Map({(d_409_v86_.f13)[default__.safeIndex(799, (d_409_v86_.f13).length(0))]: d_285_v1_})
                d_422_v94_: _dafny.Array
                nw75_ = _dafny.Array(_dafny.Array(None, 0), 5)
                d_422_v94_ = nw75_
                index85_ = default__.safeIndex(675, (d_422_v94_).length(0))
                (d_422_v94_)[index85_] = d_409_v86_.f13
                arr9_ = d_409_v86_.f13
                index86_ = default__.safeIndex(799, (d_409_v86_.f13).length(0))
                index87_ = default__.safeIndex(675, (d_422_v94_).length(0))
                rhs71_ = ((d_417_v89_)[((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]) and (d_285_v1_)] if (((d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]) and (d_285_v1_)) in (d_417_v89_) else default__.safeModuloInt((self).f7, (d_409_v86_.f13)[default__.safeIndex(799, (d_409_v86_.f13).length(0))]))
                rhs72_ = d_421_v93_
                rhs73_ = (self).f7
                rhs74_ = default__.safeModuloInt((self).f8, (self).f8)
                rhs75_ = d_409_v86_.f13
                lhs67_ = d_409_v86_.f13
                lhs68_ = default__.safeIndex(799, (d_409_v86_.f13).length(0))
                lhs69_ = d_422_v94_
                lhs70_ = default__.safeIndex(675, (d_422_v94_).length(0))
                lhs67_[lhs68_] = rhs71_
                d_421_v93_ = rhs72_
                r1 = rhs73_
                r1 = rhs74_
                lhs69_[lhs70_] = rhs75_
            elif True:
                r2 = (default__.safeModuloInt((self).f8, len(_dafny.Map({d_409_v86_.f13: d_285_v1_})))) + ((self).f8)
                d_423_v95_: _dafny.Map
                d_423_v95_ = _dafny.Map({(d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))]: ((_dafny.MultiSet(((d_411_v87_)[default__.safeIndex(487, (d_411_v87_).length(0))]).set(default__.safeIndex(len((D5_DC11(d_417_v89_)).cf18), len((d_411_v87_)[default__.safeIndex(487, (d_411_v87_).length(0))])), (self).f7))).cardinality) <= (len(_dafny.Set({(self).f7})))})
                d_424_v96_: _dafny.Seq
                d_424_v96_ = _dafny.SeqWithoutIsStrInference([d_423_v95_, d_423_v95_, d_423_v95_, (d_423_v95_) | (d_423_v95_)])
                d_423_v95_ = (d_424_v96_)[default__.safeIndex((self).f7, len(d_424_v96_))]
                d_425_v97_: _dafny.Array
                nw76_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
                d_425_v97_ = nw76_
                d_426_v98_: _dafny.MultiSet
                d_426_v98_ = _dafny.MultiSet([d_365_v48_])
                index88_ = default__.safeIndex(464, (d_283_v0_).length(0))
                rhs76_ = (d_426_v98_).ispropersubset(d_426_v98_)
                rhs77_ = d_425_v97_
                rhs78_ = d_285_v1_
                lhs71_ = globalState
                lhs72_ = d_283_v0_
                lhs73_ = default__.safeIndex(464, (d_283_v0_).length(0))
                lhs71_.f4 = rhs76_
                d_425_v97_ = rhs77_
                lhs72_[lhs73_] = rhs78_
                r2 = (self).f8
                r1 = (self).f8
        d_427_v99_: D4
        d_427_v99_ = D4_DC10((self).f7, (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))])
        pat_let_tv4_ = d_303_v7_
        def lambda14_(source10_):
            if source10_.is_DC9:
                d_428___mcc_h14_ = source10_.cf12
                d_429___mcc_h15_ = source10_.cf13
                d_430___mcc_h16_ = source10_.cf14
                d_431___mcc_h17_ = source10_.cf15
                d_432_cf15_ = d_431___mcc_h17_
                d_433_cf14_ = d_430___mcc_h16_
                d_434_cf13_ = d_429___mcc_h15_
                d_435_cf12_ = d_428___mcc_h14_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cradsw"))
            elif source10_.is_DC10:
                d_436___mcc_h18_ = source10_.cf16
                d_437___mcc_h19_ = source10_.cf17
                d_438_cf17_ = d_437___mcc_h19_
                d_439_cf16_ = d_436___mcc_h18_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqqe"))
            elif True:
                d_440___mcc_h20_ = source10_.cf11
                d_441_cf11_ = d_440___mcc_h20_
                return pat_let_tv4_

        r2 = len(lambda14_(d_427_v99_))
        (globalState).f4 = ((d_285_v1_ if (d_283_v0_)[default__.safeIndex(464, (d_283_v0_).length(0))] else (d_287_v3_)[default__.safeIndex((self).f7, len(d_287_v3_))])) and (d_285_v1_)
        d_442_v100_: _dafny.Set
        d_442_v100_ = _dafny.Set({d_285_v1_})
        r0 = (d_442_v100_).intersection((d_442_v100_).intersection(d_442_v100_))
        r1 = (self).f8
        d_443_v101_: _dafny.MultiSet
        d_443_v101_ = _dafny.MultiSet([d_283_v0_])
        r2 = ((self).f8) + (((d_443_v101_).cardinality if False else (self).f7))
        return r0, r1, r2


class C2(T0, T2, T1):
    def  __init__(self):
        self._f7: int = int(0)
        self._f8: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    def ctor__(self, f7, f8):
        (self)._f7 = f7
        (self)._f8 = f8

    def fm2(self, p0, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikasy"))

    def fm0(self, p0, globalState):
        return _dafny.Set({(_dafny.Set({False, False, True, False})) - (_dafny.Set({False})), (_dafny.Set({False, True})) - (_dafny.Set({True, False, True, True})), (_dafny.Set({not(False), True})).intersection(_dafny.Set({(False)}))})

    def fm1(self, p0, p1, p2, p3, globalState):
        return ((False if False else True))

    def fm14(self, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qx"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cv")))

    def fm15(self, globalState):
        return False

    def m1(self, p0, globalState):
        d_444_v0_: _dafny.Array
        def lambda15_(d_445_i0_):
            return ((self).f8) < ((self).f8)

        init8_ = lambda15_
        nw77_ = _dafny.Array(None, 19)
        for i0_8_ in range(nw77_.length(0)):
            nw77_[i0_8_] = init8_(i0_8_)
        d_444_v0_ = nw77_
        d_444_v0_ = d_444_v0_
        d_446_v1_: bool
        d_446_v1_ = False
        d_447_v2_: D2
        d_447_v2_ = D2_DC5(d_446_v1_, (self).f8, (self).f7, (self).f7, d_446_v1_)
        hi2_ = (d_447_v2_).cf5
        for d_448_i1_ in range((self).f7, hi2_):
            d_449_v3_: _dafny.Seq
            d_449_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsl"))
            d_449_v3_ = (p0) + (p0)
            d_450_v4_: _dafny.Array
            nw78_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_450_v4_ = nw78_
            nw79_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_450_v4_ = nw79_
            d_451_v5_: _dafny.Array
            def lambda16_(d_452_i1_):
                def lambda17_(d_453_i2_):
                    return (d_453_i2_) + (d_452_i1_)

                return lambda17_

            init9_ = lambda16_(d_448_i1_)
            nw80_ = _dafny.Array(None, 7)
            for i0_9_ in range(nw80_.length(0)):
                nw80_[i0_9_] = init9_(i0_9_)
            d_451_v5_ = nw80_
            d_454_v6_: _dafny.Map
            d_454_v6_ = _dafny.Map({375: (self).f7})
            d_455_v7_: C0
            nw81_ = C0()
            nw81_.ctor__(d_451_v5_, len((d_449_v3_) + (p0)), len(d_454_v6_))
            d_455_v7_ = nw81_
            if (860) < (len((self).fm2(d_446_v1_, globalState))):
                d_456_v8_: int
                d_456_v8_ = -417
                d_456_v8_ = (-602) - ((self).f8)
                d_456_v8_ = d_456_v8_
                (globalState).f4 = d_446_v1_
                (globalState).f3 = (d_446_v1_ if d_446_v1_ else (d_448_i1_) >= (-965))
                index89_ = default__.safeIndex(310, (d_444_v0_).length(0))
                (d_444_v0_)[index89_] = d_446_v1_
                index90_ = default__.safeIndex(310, (d_444_v0_).length(0))
                (d_444_v0_)[index90_] = not(d_446_v1_)
            elif True:
                d_457_v9_: _dafny.MultiSet
                d_457_v9_ = _dafny.MultiSet([(self).f8])
                d_458_v10_: _dafny.Seq
                d_458_v10_ = _dafny.SeqWithoutIsStrInference([d_457_v9_, d_457_v9_, d_457_v9_])
                d_459_v11_: _dafny.Map
                d_459_v11_ = _dafny.Map({d_444_v0_: (self).f7})
                d_460_v12_: _dafny.Map
                d_460_v12_ = _dafny.Map({len(d_459_v11_): d_454_v6_})
                d_461_v13_: _dafny.Map
                d_461_v13_ = _dafny.Map({len(d_458_v10_): (d_460_v12_) | (d_460_v12_)})
                d_461_v13_ = (d_461_v13_).set(default__.fm16(d_446_v1_, globalState), (d_460_v12_).set(d_448_i1_, d_454_v6_))
                d_462_v14_: _dafny.Seq
                d_462_v14_ = _dafny.SeqWithoutIsStrInference([(self).f8])
                d_463_v15_: _dafny.Seq
                d_463_v15_ = _dafny.SeqWithoutIsStrInference([d_462_v14_])
                d_464_v16_: _dafny.MultiSet
                d_464_v16_ = _dafny.MultiSet([d_446_v1_, d_446_v1_, not(d_446_v1_), ((d_463_v15_)[default__.safeIndex((d_447_v2_).cf5, len(d_463_v15_))]) == (d_462_v14_)])
                d_465_v17_: str
                d_465_v17_ = _dafny.CodePoint('a')
                d_466_v18_: _dafny.Map
                d_466_v18_ = _dafny.Map({d_446_v1_: len((d_449_v3_).set(default__.safeIndex((0) - (d_448_i1_), len(d_449_v3_)), d_465_v17_))})
                d_467_v19_: _dafny.Seq
                d_467_v19_ = _dafny.SeqWithoutIsStrInference([True, not(not(d_446_v1_)), d_446_v1_])
                rhs79_ = False
                rhs80_ = ((d_464_v16_) | (d_464_v16_)) | (d_464_v16_)
                rhs81_ = (d_466_v18_) | (d_466_v18_)
                rhs82_ = d_467_v19_
                lhs74_ = globalState
                lhs74_.f3 = rhs79_
                d_464_v16_ = rhs80_
                d_466_v18_ = rhs81_
                d_467_v19_ = rhs82_
                d_468_v20_: C0
                nw82_ = C0()
                nw82_.ctor__(d_455_v7_.f13, (self).f8, (self).f7)
                d_468_v20_ = nw82_
                d_469_v21_: _dafny.Array
                nw83_ = _dafny.Array(_dafny.CodePoint('D'), 2)
                d_469_v21_ = nw83_
                index91_ = default__.safeIndex(361, (d_469_v21_).length(0))
                (d_469_v21_)[index91_] = (p0)[default__.safeIndex(default__.fm16(False, globalState), len(p0))]
                index92_ = default__.safeIndex(361, (d_469_v21_).length(0))
                (d_469_v21_)[index92_] = d_465_v17_
                d_470_v22_: C0
                nw84_ = C0()
                nw84_.ctor__((d_455_v7_.f13 if d_446_v1_ else d_451_v5_), default__.safeModuloInt(d_448_i1_, len(_dafny.SeqWithoutIsStrInference([(d_469_v21_)[default__.safeIndex(361, (d_469_v21_).length(0))] for d_471_i3_ in range(default__.abs(286))]))), default__.fm16(d_446_v1_, globalState))
                d_470_v22_ = nw84_
        d_472_v24_: _dafny.MultiSet
        d_472_v24_ = _dafny.MultiSet([d_446_v1_, d_446_v1_])
        d_473_v25_: _dafny.Set
        d_473_v25_ = _dafny.Set({(d_472_v24_).cardinality, (self).f8})
        d_474_v26_: _dafny.Array
        nw85_ = _dafny.Array(None, 2)
        def iife25_():
            coll25_ = _dafny.Set()
            compr_25_: int
            for compr_25_ in _dafny.IntegerRange(695, -840):
                d_475_v23_: int = compr_25_
                if ((695) <= (d_475_v23_)) and ((d_475_v23_) < (-840)):
                    coll25_ = coll25_.union(_dafny.Set([(d_475_v23_) * ((0) - ((self).f7))]))
            return _dafny.Set(coll25_)
        nw85_[int(0)] = iife25_()
        
        nw85_[int(1)] = d_473_v25_
        d_474_v26_ = nw85_
        d_476_v27_: _dafny.Map
        d_476_v27_ = _dafny.Map({(0) - ((0) - ((0) - ((self).f8))): d_474_v26_})
        d_477_v28_: _dafny.Array
        nw86_ = _dafny.Array(None, 11)
        nw86_[int(0)] = (d_474_v26_ if d_446_v1_ else d_474_v26_)
        nw86_[int(1)] = ((d_476_v27_)[(self).f8] if ((self).f8) in (d_476_v27_) else d_474_v26_)
        nw86_[int(2)] = d_474_v26_
        nw86_[int(3)] = d_474_v26_
        nw86_[int(4)] = d_474_v26_
        nw86_[int(5)] = d_474_v26_
        nw86_[int(6)] = d_474_v26_
        nw86_[int(7)] = d_474_v26_
        nw86_[int(8)] = d_474_v26_
        nw86_[int(9)] = d_474_v26_
        nw86_[int(10)] = d_474_v26_
        d_477_v28_ = nw86_
        index93_ = default__.safeIndex(734, (d_477_v28_).length(0))
        (d_477_v28_)[index93_] = d_474_v26_
        index94_ = default__.safeIndex(734, (d_477_v28_).length(0))
        (d_477_v28_)[index94_] = d_474_v26_
        d_478_v29_: _dafny.MultiSet
        d_478_v29_ = _dafny.MultiSet([715, (0) - ((self).f7)])
        d_479_v30_: _dafny.Map
        d_479_v30_ = _dafny.Map({d_446_v1_: d_446_v1_})
        d_480_v31_: _dafny.Array
        nw87_ = _dafny.Array(None, 25)
        nw87_[int(0)] = len(_dafny.Map({(d_478_v29_).cardinality: (p0).set(default__.safeIndex((d_472_v24_).cardinality, len(p0)), _dafny.CodePoint('y'))}))
        nw87_[int(1)] = (self).f8
        nw87_[int(2)] = len(d_473_v25_)
        nw87_[int(3)] = 285
        nw87_[int(4)] = (self).f7
        nw87_[int(5)] = (self).f8
        nw87_[int(6)] = (self).f8
        nw87_[int(7)] = (self).f8
        nw87_[int(8)] = len(p0)
        nw87_[int(9)] = (self).f7
        nw87_[int(10)] = (self).f8
        nw87_[int(11)] = (self).f8
        nw87_[int(12)] = 891
        nw87_[int(13)] = default__.fm16(d_446_v1_, globalState)
        nw87_[int(14)] = (0) - ((self).f7)
        nw87_[int(15)] = (self).f8
        nw87_[int(16)] = (self).f7
        nw87_[int(17)] = len(p0)
        nw87_[int(18)] = (0) - (len(d_479_v30_))
        nw87_[int(19)] = (self).f8
        nw87_[int(20)] = (self).f7
        nw87_[int(21)] = (self).f7
        nw87_[int(22)] = (self).f7
        nw87_[int(23)] = (self).f7
        nw87_[int(24)] = len(p0)
        d_480_v31_ = nw87_
        d_481_v32_: C0
        nw88_ = C0()
        nw88_.ctor__(d_480_v31_, (self).f7, (self).f7)
        d_481_v32_ = nw88_
        (globalState).f4 = d_446_v1_
        d_482_v33_: C0
        nw89_ = C0()
        nw89_.ctor__(d_481_v32_.f13, (0) - ((self).f8), ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trnq")), p0, p0, p0, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_483_i4_ in range(default__.abs(53))])])).set(p0, default__.abs((self).f7))).cardinality)
        d_482_v33_ = nw89_

    def m2(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        r2: bool = False
        d_484_v0_: _dafny.Array
        def lambda18_(d_485_i0_):
            return default__.safeDivisionInt(d_485_i0_, (self).f8)

        init10_ = lambda18_
        nw90_ = _dafny.Array(None, 2)
        for i0_10_ in range(nw90_.length(0)):
            nw90_[i0_10_] = init10_(i0_10_)
        d_484_v0_ = nw90_
        d_486_v1_: C0
        nw91_ = C0()
        nw91_.ctor__(d_484_v0_, default__.safeModuloInt(p0, (self).f8), 765)
        d_486_v1_ = nw91_
        (globalState).f4 = p1
        d_487_v2_: _dafny.Seq
        d_487_v2_ = _dafny.SeqWithoutIsStrInference([True, p3])
        d_488_v3_: D2
        d_488_v3_ = D2_DC4(_dafny.MultiSet(d_487_v2_))
        source11_ = d_488_v3_
        if source11_.is_DC5:
            d_489___mcc_h0_ = source11_.cf4
            d_490___mcc_h1_ = source11_.cf5
            d_491___mcc_h2_ = source11_.cf6
            d_492___mcc_h3_ = source11_.cf7
            d_493___mcc_h4_ = source11_.cf8
            d_494_cf8_ = d_493___mcc_h4_
            d_495_cf7_ = d_492___mcc_h3_
            d_496_cf6_ = d_491___mcc_h2_
            d_497_cf5_ = d_490___mcc_h1_
            d_498_cf4_ = d_489___mcc_h0_
            if not(((self).f8) < (d_497_cf5_)):
                (globalState).f4 = d_498_cf4_
                r2 = (d_494_cf8_ if p3 else (d_494_cf8_ if d_494_cf8_ else False))
                (globalState).f3 = p1
                d_499_v4_: str
                d_499_v4_ = _dafny.CodePoint('c')
                d_500_v5_: _dafny.Map
                d_500_v5_ = _dafny.Map({(p3 if True else False): d_499_v4_})
                d_500_v5_ = (d_500_v5_).set(p1, _dafny.CodePoint('e'))
                d_501_v6_: _dafny.Seq
                d_501_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksccjtf"))
                d_501_v6_ = d_501_v6_
            elif True:
                d_496_cf6_ = default__.safeModuloInt(len((self).fm14(globalState)), d_495_cf7_)
                d_502_v7_: str
                d_502_v7_ = _dafny.CodePoint('g')
                rhs83_ = d_502_v7_
                rhs84_ = d_487_v2_
                d_502_v7_ = rhs83_
                r0 = rhs84_
                d_503_v8_: _dafny.Array
                nw92_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
                d_503_v8_ = nw92_
                d_504_v9_: _dafny.Seq
                d_504_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucud"))
                index95_ = default__.safeIndex(902, (d_503_v8_).length(0))
                (d_503_v8_)[index95_] = (d_504_v9_).set(default__.safeIndex(448, len(d_504_v9_)), d_502_v7_)
                index96_ = default__.safeIndex(902, (d_503_v8_).length(0))
                (d_503_v8_)[index96_] = (d_504_v9_) + (_dafny.SeqWithoutIsStrInference([d_502_v7_ for d_505_i1_ in range(default__.abs(-60))]))
                r0 = (_dafny.SeqWithoutIsStrInference([p3])) + ((d_487_v2_) + (_dafny.SeqWithoutIsStrInference([p3])))
                d_506_v10_: _dafny.Array
                nw93_ = _dafny.Array(False, 13)
                d_506_v10_ = nw93_
                index97_ = default__.safeIndex(272, (d_506_v10_).length(0))
                (d_506_v10_)[index97_] = d_494_cf8_
                d_507_v11_: _dafny.Seq
                d_507_v11_ = _dafny.SeqWithoutIsStrInference([d_496_cf6_])
                d_508_v12_: _dafny.Map
                d_508_v12_ = _dafny.Map({d_498_cf4_: d_507_v11_})
                index98_ = default__.safeIndex(272, (d_506_v10_).length(0))
                rhs85_ = p3
                rhs86_ = p1
                rhs87_ = (d_496_cf6_ if p1 else d_496_cf6_)
                rhs88_ = d_507_v11_
                rhs89_ = default__.fm17(d_502_v7_, d_498_cf4_, globalState)
                lhs75_ = d_506_v10_
                lhs76_ = default__.safeIndex(272, (d_506_v10_).length(0))
                lhs75_[lhs76_] = rhs85_
                r2 = rhs86_
                d_496_cf6_ = rhs87_
                d_507_v11_ = rhs88_
                d_508_v12_ = rhs89_
            if p1:
                d_509_v13_: _dafny.Seq
                d_509_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
                d_509_v13_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_510_i2_ in range(default__.abs(-528))])) + (d_509_v13_)
                d_511_v14_: _dafny.Map
                d_511_v14_ = _dafny.Map({(self).f7: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lb"))})
                d_511_v14_ = (d_511_v14_).set(d_496_cf6_, d_509_v13_)
                d_512_v15_: T0
                nw94_ = C1()
                nw94_.ctor__((self).f7, (self).f8)
                d_512_v15_ = nw94_
                d_512_v15_ = (d_512_v15_ if p1 else d_512_v15_)
                d_513_v16_: _dafny.Set
                d_513_v16_ = _dafny.Set({d_494_cf8_})
                d_514_v17_: _dafny.Map
                d_514_v17_ = _dafny.Map({(d_512_v15_).f7: 833})
                (globalState).f3 = (d_486_v1_).fm1(d_513_v16_, (self).f8, (0) - (len(d_514_v17_)), (d_512_v15_).f8, globalState)
                d_515_v18_: _dafny.Map
                d_515_v18_ = _dafny.Map({p3: (d_512_v15_).f8})
                def iife26_():
                    coll26_ = _dafny.Map()
                    compr_26_: int
                    for compr_26_ in _dafny.IntegerRange(437, -935):
                        d_516_v19_: int = compr_26_
                        if ((437) <= (d_516_v19_)) and ((d_516_v19_) < (-935)):
                            coll26_[default__.safeModuloInt(d_516_v19_, d_497_cf5_)] = p3
                    return _dafny.Map(coll26_)
                d_515_v18_ = (d_515_v18_).set(p1, (len(iife26_()
                )) - (p0))
            elif True:
                d_517_v20_: _dafny.Map
                d_517_v20_ = _dafny.Map({not(p3): (d_497_cf5_) < (p2)})
                d_517_v20_ = (d_517_v20_).set(d_494_cf8_, d_494_cf8_)
                (globalState).f4 = (d_496_cf6_) > (default__.safeModuloInt(915, d_495_cf7_))
                d_518_v21_: str
                d_518_v21_ = _dafny.CodePoint('v')
                d_519_v22_: _dafny.Seq
                d_519_v22_ = _dafny.SeqWithoutIsStrInference([d_518_v21_])
                arr10_ = d_486_v1_.f13
                index99_ = default__.safeIndex(904, (d_486_v1_.f13).length(0))
                arr10_[index99_] = len((d_519_v22_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')])))
                d_520_v23_: _dafny.Map
                d_520_v23_ = _dafny.Map({not(d_494_cf8_): p0})
                d_521_v24_: _dafny.MultiSet
                d_521_v24_ = _dafny.MultiSet([_dafny.Map({p3: len(d_519_v22_)})])
                arr11_ = d_486_v1_.f13
                index100_ = default__.safeIndex(904, (d_486_v1_.f13).length(0))
                arr11_[index100_] = (((d_520_v23_)[p3] if (p3) in (d_520_v23_) else (d_521_v24_).cardinality) if d_494_cf8_ else ((self).f7) - (p2))
                d_517_v20_ = (d_517_v20_).set(False, (len(d_519_v22_)) == (d_495_cf7_))
                d_522_v25_: _dafny.Seq
                d_522_v25_ = _dafny.SeqWithoutIsStrInference([d_497_cf5_, d_495_cf7_])
                d_523_v26_: _dafny.MultiSet
                d_523_v26_ = _dafny.MultiSet([d_522_v25_])
                d_523_v26_ = _dafny.MultiSet([d_522_v25_, d_522_v25_])
            d_497_cf5_ = d_495_cf7_
            (globalState).f4 = d_498_cf4_
        elif True:
            d_524___mcc_h5_ = source11_.cf3
            d_525_cf3_ = d_524___mcc_h5_
            r1 = (self).f8
            d_526_v27_: str
            d_526_v27_ = _dafny.CodePoint('y')
            d_527_v28_: _dafny.Set
            d_527_v28_ = _dafny.Set({d_526_v27_, d_526_v27_, _dafny.CodePoint('d')})
            d_528_v29_: _dafny.Seq
            d_528_v29_ = _dafny.SeqWithoutIsStrInference([d_527_v28_])
            (globalState).f4 = (p2) >= (len(d_528_v29_))
            d_529_v30_: _dafny.Array
            nw95_ = _dafny.Array(False, 18)
            d_529_v30_ = nw95_
            d_530_v31_: _dafny.Seq
            d_530_v31_ = _dafny.SeqWithoutIsStrInference([d_529_v30_])
            d_531_v32_: D4
            d_531_v32_ = D4_DC9(p0, d_530_v31_, (self).f7, p0)
            d_532_v33_: _dafny.Array
            nw96_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_532_v33_ = nw96_
            d_533_v34_: D4
            d_533_v34_ = D4_DC8(d_532_v33_)
            d_534_v35_: _dafny.Map
            d_534_v35_ = _dafny.Map({p1: d_533_v34_})
            pat_let_tv5_ = p0
            pat_let_tv6_ = d_534_v35_
            def iife28_(_pat_let1_0):
                def iife29_(d_535_dt__update__tmp_h0_):
                    def iife30_(_pat_let2_0):
                        def iife31_(d_536_dt__update_hcf5_h0_):
                            def iife32_(_pat_let3_0):
                                def iife33_(d_537_dt__update_hcf8_h0_):
                                    return D2_DC5((d_535_dt__update__tmp_h0_).cf4, d_536_dt__update_hcf5_h0_, (d_535_dt__update__tmp_h0_).cf6, (d_535_dt__update__tmp_h0_).cf7, d_537_dt__update_hcf8_h0_)
                                return iife33_(_pat_let3_0)
                            return iife32_(False)
                        return iife31_(_pat_let2_0)
                    return iife30_(pat_let_tv5_)
                return iife29_(_pat_let1_0)
            def iife27_(_pat_let0_0):
                def iife34_(d_538_dt__update__tmp_h1_):
                    def iife35_(_pat_let4_0):
                        def iife36_(d_539_dt__update_hcf6_h0_):
                            return D2_DC5((d_538_dt__update__tmp_h1_).cf4, (d_538_dt__update__tmp_h1_).cf5, d_539_dt__update_hcf6_h0_, (d_538_dt__update__tmp_h1_).cf7, (d_538_dt__update__tmp_h1_).cf8)
                        return iife36_(_pat_let4_0)
                    return iife35_((0) - ((0) - (len(pat_let_tv6_))))
                return iife34_(_pat_let0_0)
            source12_ = iife27_(iife28_(D2_DC5(True, (self).f8, 475, (d_531_v32_).cf14, p3)))
            if source12_.is_DC5:
                d_540___mcc_h6_ = source12_.cf4
                d_541___mcc_h7_ = source12_.cf5
                d_542___mcc_h8_ = source12_.cf6
                d_543___mcc_h9_ = source12_.cf7
                d_544___mcc_h10_ = source12_.cf8
                d_545_cf8_ = d_544___mcc_h10_
                d_546_cf7_ = d_543___mcc_h9_
                d_547_cf6_ = d_542___mcc_h8_
                d_548_cf5_ = d_541___mcc_h7_
                d_549_cf4_ = d_540___mcc_h6_
                d_550_v36_: _dafny.Map
                d_550_v36_ = _dafny.Map({not(d_545_cf8_): d_486_v1_.f13})
                d_551_v37_: _dafny.Set
                d_551_v37_ = _dafny.Set({len(d_550_v36_)})
                d_552_v38_: D3
                d_552_v38_ = D3_DC6((d_551_v37_) | (d_551_v37_))
                d_552_v38_ = d_552_v38_
                (globalState).f4 = p3
                d_553_v39_: _dafny.Set
                d_553_v39_ = _dafny.Set({p1})
                index101_ = default__.safeIndex(84, (d_529_v30_).length(0))
                (d_529_v30_)[index101_] = True
                d_554_v40_: _dafny.Seq
                d_554_v40_ = _dafny.SeqWithoutIsStrInference([d_553_v39_, d_553_v39_, _dafny.Set({(d_487_v2_)[default__.safeIndex((self).f7, len(d_487_v2_))]})])
                index102_ = default__.safeIndex(84, (d_529_v30_).length(0))
                rhs90_ = (d_554_v40_)[default__.safeIndex((len(d_487_v2_)) + (d_546_cf7_), len(d_554_v40_))]
                rhs91_ = (((self).f8) == (p0) if p3 else p1)
                lhs77_ = d_529_v30_
                lhs78_ = default__.safeIndex(84, (d_529_v30_).length(0))
                d_553_v39_ = rhs90_
                lhs77_[lhs78_] = rhs91_
                d_547_cf6_ = (self).f8
            elif True:
                d_555___mcc_h11_ = source12_.cf3
                d_556_cf3_ = d_555___mcc_h11_
                (globalState).f4 = ((self).f7) > (p2)
                r1 = (0) - (-380)
                d_557_v41_: _dafny.Array
                nw97_ = _dafny.Array(_dafny.Set({}), 5)
                d_557_v41_ = nw97_
                d_558_v44_: _dafny.Set
                d_558_v44_ = _dafny.Set({p0})
                index103_ = default__.safeIndex(757, (d_557_v41_).length(0))
                def iife37_():
                    def iife39_():
                        coll29_ = _dafny.Set()
                        compr_29_: int
                        for compr_29_ in _dafny.IntegerRange(-205, 46):
                            d_561_v42_: int = compr_29_
                            if ((-205) <= (d_561_v42_)) and ((d_561_v42_) < (46)):
                                coll29_ = coll29_.union(_dafny.Set([default__.safeDivisionInt(d_561_v42_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vudgnl"))))]))
                        return _dafny.Set(coll29_)
                    coll27_ = _dafny.Set()
                    def iife38_():
                        coll28_ = _dafny.Set()
                        compr_28_: int
                        for compr_28_ in _dafny.IntegerRange(-205, 46):
                            d_559_v42_: int = compr_28_
                            if ((-205) <= (d_559_v42_)) and ((d_559_v42_) < (46)):
                                coll28_ = coll28_.union(_dafny.Set([default__.safeDivisionInt(d_559_v42_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vudgnl"))))]))
                        return _dafny.Set(coll28_)
                    compr_27_: int
                    for compr_27_ in (iife38_()
                    ).Elements:
                        d_560_v43_: int = compr_27_
                        if (d_560_v43_) in (iife39_()
                        ):
                            coll27_ = coll27_.union(_dafny.Set([default__.safeDivisionInt(d_560_v43_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_562_i3_ in range(default__.abs(432))])))]))
                    return _dafny.Set(coll27_)
                (d_557_v41_)[index103_] = (iife37_()
                ) - (d_558_v44_)
                d_563_v46_: D3
                d_563_v46_ = D3_DC6(_dafny.Set({p2}))
                pat_let_tv7_ = p2
                pat_let_tv8_ = d_530_v31_
                index104_ = default__.safeIndex(757, (d_557_v41_).length(0))
                def iife40_():
                    coll30_ = _dafny.Set()
                    compr_30_: str
                    for compr_30_ in (_dafny.SeqWithoutIsStrInference([d_526_v27_ for d_564_i4_ in range(default__.abs(986))])).Elements:
                        d_565_v45_: str = compr_30_
                        if (d_565_v45_) in (_dafny.SeqWithoutIsStrInference([d_526_v27_ for d_564_i4_ in range(default__.abs(986))])):
                            coll30_ = coll30_.union(_dafny.Set([d_565_v45_]))
                    return _dafny.Set(coll30_)
                def iife41_(_pat_let5_0):
                    def iife42_(d_566_dt__update__tmp_h2_):
                        def iife43_(_pat_let6_0):
                            def iife44_(d_567_dt__update_hcf14_h0_):
                                return D4_DC9((d_566_dt__update__tmp_h2_).cf12, (d_566_dt__update__tmp_h2_).cf13, d_567_dt__update_hcf14_h0_, (d_566_dt__update__tmp_h2_).cf15)
                            return iife44_(_pat_let6_0)
                        return iife43_((pat_let_tv7_) - (len(pat_let_tv8_)))
                    return iife42_(_pat_let5_0)
                rhs92_ = (d_527_v28_).issubset((d_527_v28_).intersection(iife40_()
                ))
                rhs93_ = d_487_v2_
                rhs94_ = (self).f7
                rhs95_ = iife41_(d_531_v32_)
                rhs96_ = (d_563_v46_).cf9
                lhs79_ = d_557_v41_
                lhs80_ = default__.safeIndex(757, (d_557_v41_).length(0))
                r2 = rhs92_
                r0 = rhs93_
                r1 = rhs94_
                d_531_v32_ = rhs95_
                lhs79_[lhs80_] = rhs96_
                d_529_v30_ = d_529_v30_
            d_568_v47_: _dafny.MultiSet
            d_568_v47_ = _dafny.MultiSet([p0, p2])
            d_569_v48_: C1
            nw98_ = C1()
            nw98_.ctor__((0) - ((0) - (default__.fm16(True, globalState))), (0) - (len(default__.fm29(p1, d_568_v47_, globalState))))
            d_569_v48_ = nw98_
        d_570_v49_: _dafny.MultiSet
        d_570_v49_ = _dafny.MultiSet([p1])
        d_571_v50_: _dafny.Seq
        d_571_v50_ = _dafny.SeqWithoutIsStrInference([default__.fm20(p3, not(False), p3, globalState), ((d_570_v49_)[p3] if (p3) in (d_570_v49_) else (0) - ((self).f7))])
        hi3_ = p0
        for d_572_i5_ in range(len((d_571_v50_ if p3 else d_571_v50_)), hi3_):
            (globalState).f3 = p1
            d_573_v51_: _dafny.Seq
            d_573_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvldbus"))
            d_574_v52_: _dafny.Map
            d_574_v52_ = _dafny.Map({p3: len(d_573_v51_)})
            d_575_v53_: _dafny.MultiSet
            d_575_v53_ = _dafny.MultiSet([p0])
            d_574_v52_ = (d_574_v52_).set(p3, ((d_575_v53_)[p2] if (p2) in (d_575_v53_) else p0))
            d_576_v54_: _dafny.Array
            def lambda19_(d_577_p3_):
                def lambda20_(d_578_i6_):
                    return d_577_p3_

                return lambda20_

            init11_ = lambda19_(p3)
            nw99_ = _dafny.Array(None, 13)
            for i0_11_ in range(nw99_.length(0)):
                nw99_[i0_11_] = init11_(i0_11_)
            d_576_v54_ = nw99_
            d_579_v55_: D4
            d_579_v55_ = D4_DC9(p2, _dafny.SeqWithoutIsStrInference([d_576_v54_, d_576_v54_, d_576_v54_, d_576_v54_, d_576_v54_]), 688, (self).f8)
            arr12_ = d_486_v1_.f13
            index105_ = default__.safeIndex(503, (d_486_v1_.f13).length(0))
            arr12_[index105_] = default__.safeDivisionInt(((d_574_v52_)[p1] if (p1) in (d_574_v52_) else p2), (d_579_v55_).cf14)
            arr13_ = d_486_v1_.f13
            index106_ = default__.safeIndex(503, (d_486_v1_.f13).length(0))
            arr13_[index106_] = (p0) * ((self).f7)
            d_580_v56_: _dafny.Map
            d_580_v56_ = _dafny.Map({len(d_573_v51_): d_484_v0_})
            d_581_v57_: _dafny.Map
            d_581_v57_ = _dafny.Map({d_580_v56_: d_575_v53_})
            d_581_v57_ = (d_581_v57_).set(d_580_v56_, _dafny.MultiSet((d_571_v50_ if p1 else _dafny.SeqWithoutIsStrInference([(d_575_v53_).cardinality, (self).f7, -717, 620, 742]))))
        r1 = default__.fm16((self).fm15(globalState), globalState)
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_484_v0_).length(0)):
            d_582_i7_: int = guard_loop_2_
            if (True) and (((0) <= (d_582_i7_)) and ((d_582_i7_) < ((d_484_v0_).length(0)))):
                (d_484_v0_)[(d_582_i7_)] = default__.safeModuloInt(d_582_i7_, (self).f8)
        d_583_v58_: _dafny.Set
        d_583_v58_ = _dafny.Set({p3})
        r0 = (d_487_v2_).set(default__.safeIndex((0) - ((self).f7), len(d_487_v2_)), ((self).f7) != (len(d_583_v58_)))
        d_584_v59_: _dafny.Set
        d_584_v59_ = _dafny.Set({p0})
        d_585_v60_: D3
        d_585_v60_ = D3_DC6(d_584_v59_)
        pat_let_tv9_ = p0
        pat_let_tv10_ = p2
        def lambda21_(source13_):
            if source13_.is_DC7:
                d_586___mcc_h12_ = source13_.cf10
                d_587_cf10_ = d_586___mcc_h12_
                if d_587_cf10_:
                    return pat_let_tv9_
                elif True:
                    return (_dafny.MultiSet([pat_let_tv10_, -808, (self).f8, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjemwxaw"))), (self).f8])).cardinality
            elif True:
                d_588___mcc_h13_ = source13_.cf9
                d_589_cf9_ = d_588___mcc_h13_
                return 332

        r1 = lambda21_(d_585_v60_)
        d_590_v61_: _dafny.Seq
        d_590_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtyih"))
        r2 = (default__.fm20(False, p3, p1, globalState)) != (len(d_590_v61_))
        return r0, r1, r2


class C3(T2, T0, T1):
    def  __init__(self):
        self._f7: int = int(0)
        self._f8: int = int(0)
        self._f12: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    def ctor__(self, f12, f7, f8):
        (self)._f12 = f12
        (self)._f7 = f7
        (self)._f8 = f8

    def fm2(self, p0, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skp"))

    def fm0(self, p0, globalState):
        source14_ = (self).f12
        d_591___mcc_h0_ = source14_
        d_592_cf0_ = d_591___mcc_h0_
        if (self).f12:
            return _dafny.Set({_dafny.Set({False, d_592_cf0_})})
        elif True:
            return _dafny.Set({_dafny.Set({d_592_cf0_}), _dafny.Set({(self).f12})})

    def fm1(self, p0, p1, p2, p3, globalState):
        def iife45_():
            coll31_ = _dafny.Map()
            compr_31_: int
            for compr_31_ in _dafny.IntegerRange(234, 342):
                d_593_v0_: int = compr_31_
                if ((234) <= (d_593_v0_)) and ((d_593_v0_) < (342)):
                    coll31_[(d_593_v0_) - (906)] = (self).f12
            return _dafny.Map(coll31_)
        return (default__.safeModuloInt(len(_dafny.Set({(self).f12, (self).f12, (self).f12, (self).f12, not((self).f12)})), len(iife45_()
        ))) >= (((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f8, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xutwsx"))), (self).f8]))).cardinality) - ((self).f8))

    def fm12(self, globalState):
        return (self).f12

    def fm13(self, p0, globalState):
        return (self).f8

    def m1(self, p0, globalState):
        d_594_v0_: _dafny.Seq
        d_594_v0_ = _dafny.SeqWithoutIsStrInference([(self).f12, not((self).f12), (self).f12])
        hi4_ = (0) - ((self).f7)
        for d_595_i0_ in range((0) - (len((d_594_v0_) + (d_594_v0_))), hi4_):
            d_596_v1_: bool
            d_596_v1_ = (self).f12
            if not((d_595_i0_) > ((_dafny.MultiSet([d_596_v1_])).cardinality)):
                d_597_v2_: _dafny.Array
                def lambda22_(d_598_i0_):
                    def lambda23_(d_599_i1_):
                        return default__.safeModuloInt(d_599_i1_, d_598_i0_)

                    return lambda23_

                init12_ = lambda22_(d_595_i0_)
                nw100_ = _dafny.Array(None, 22)
                for i0_12_ in range(nw100_.length(0)):
                    nw100_[i0_12_] = init12_(i0_12_)
                d_597_v2_ = nw100_
                d_600_v3_: _dafny.MultiSet
                d_600_v3_ = _dafny.MultiSet([(self).f7])
                d_601_v4_: _dafny.Map
                d_601_v4_ = _dafny.Map({d_597_v2_: d_600_v3_})
                d_602_v5_: _dafny.Array
                def lambda24_(d_603_i2_):
                    return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_604_i3_ in range(default__.abs(735))]))) != (_dafny.MultiSet([_dafny.CodePoint('e')]))

                init13_ = lambda24_
                nw101_ = _dafny.Array(None, 29)
                for i0_13_ in range(nw101_.length(0)):
                    nw101_[i0_13_] = init13_(i0_13_)
                d_602_v5_ = nw101_
                index107_ = default__.safeIndex(259, (d_602_v5_).length(0))
                (d_602_v5_)[index107_] = (self).f12
                d_605_v6_: _dafny.Set
                d_605_v6_ = _dafny.Set({(self).f7, (self).f8})
                d_606_v7_: _dafny.Seq
                d_606_v7_ = _dafny.SeqWithoutIsStrInference([len(d_605_v6_)])
                index108_ = default__.safeIndex(259, (d_602_v5_).length(0))
                rhs97_ = (d_601_v4_) | (d_601_v4_)
                rhs98_ = ((d_606_v7_) + (_dafny.SeqWithoutIsStrInference([(self).f7]))) != ((d_606_v7_) + (d_606_v7_))
                rhs99_ = (self).f12
                rhs100_ = (self).f12
                lhs81_ = globalState
                lhs82_ = globalState
                lhs83_ = d_602_v5_
                lhs84_ = default__.safeIndex(259, (d_602_v5_).length(0))
                d_601_v4_ = rhs97_
                lhs81_.f3 = rhs98_
                lhs82_.f4 = rhs99_
                lhs83_[lhs84_] = rhs100_
                d_607_v8_: str
                d_607_v8_ = _dafny.CodePoint('r')
                d_608_v9_: D1
                d_608_v9_ = D1_DC1((D1_DC1(_dafny.CodePoint('j'))).cf1)
                d_607_v8_ = (d_608_v9_).cf1
                (globalState).f3 = (d_602_v5_)[default__.safeIndex(259, (d_602_v5_).length(0))]
                (globalState).f3 = (d_594_v0_)[default__.safeIndex(d_595_i0_, len(d_594_v0_))]
                index109_ = default__.safeIndex(259, (d_602_v5_).length(0))
                (d_602_v5_)[index109_] = (self).f12
            elif True:
                (globalState).f4 = ((self).f12) and (((self).f12 if (self).f12 else (self).f12))
                d_609_v10_: _dafny.MultiSet
                d_609_v10_ = _dafny.MultiSet([d_594_v0_])
                (globalState).f3 = (d_609_v10_).ispropersubset((d_609_v10_) - (d_609_v10_))
                d_610_v11_: _dafny.Array
                def lambda25_(d_611_i4_):
                    return (self).f12

                init14_ = lambda25_
                nw102_ = _dafny.Array(None, 28)
                for i0_14_ in range(nw102_.length(0)):
                    nw102_[i0_14_] = init14_(i0_14_)
                d_610_v11_ = nw102_
                d_610_v11_ = d_610_v11_
                d_612_v12_: _dafny.Seq
                d_612_v12_ = _dafny.SeqWithoutIsStrInference([d_594_v0_])
                d_612_v12_ = d_612_v12_
                d_613_v13_: int
                d_613_v13_ = 658
                d_613_v13_ = ((self).f7) + (default__.safeModuloInt((self).f8, (self).f7))
            d_614_v14_: str
            d_614_v14_ = _dafny.CodePoint('g')
            d_614_v14_ = d_614_v14_
            (globalState).f4 = (self).f12
            d_615_v15_: _dafny.Map
            d_615_v15_ = _dafny.Map({(self).f8: default__.safeModuloInt(d_595_i0_, (_dafny.MultiSet([(self).f12])).cardinality)})
            d_615_v15_ = d_615_v15_
        d_616_v16_: _dafny.MultiSet
        d_616_v16_ = _dafny.MultiSet([(self).f8])
        d_617_v17_: _dafny.Array
        d_618_v18_: bool
        out2_: _dafny.Array
        out3_: bool
        out2_, out3_ = default__.m0((self).f7, (self).fm1(_dafny.Set({(self).f12, (self).f12}), -562, (self).f7, len(_dafny.Map({len(_dafny.Set({893})): (self).f8})), globalState), (d_616_v16_).issubset(d_616_v16_), globalState)
        d_617_v17_ = out2_
        d_618_v18_ = out3_
        d_619_v19_: str
        d_619_v19_ = _dafny.CodePoint('b')
        d_620_v20_: _dafny.Map
        d_620_v20_ = _dafny.Map({(self).fm12(globalState): d_619_v19_})
        d_620_v20_ = d_620_v20_
        d_621_i5_: int
        d_621_i5_ = 0
        with _dafny.label("3"):
            while (self).f12:
                with _dafny.c_label("3"):
                    if (d_621_i5_) >= (100):
                        raise _dafny.Break("3")
                    d_621_i5_ = (d_621_i5_) + (1)
                    d_622_v21_: _dafny.Seq
                    d_622_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xronwep"))
                    rhs101_ = p0
                    rhs102_ = (self).f12
                    d_622_v21_ = rhs101_
                    d_618_v18_ = rhs102_
                    d_618_v18_ = (self).f12
                    d_623_v22_: int
                    d_623_v22_ = 771
                    d_624_v23_: _dafny.Array
                    nw103_ = _dafny.Array(False, 1)
                    d_624_v23_ = nw103_
                    d_625_v24_: D1
                    d_625_v24_ = D1_DC1(d_619_v19_)
                    d_626_v25_: _dafny.MultiSet
                    d_626_v25_ = _dafny.MultiSet([d_625_v24_])
                    d_627_v26_: _dafny.Seq
                    d_627_v26_ = _dafny.SeqWithoutIsStrInference([d_626_v25_])
                    rhs103_ = (0) - (default__.safeModuloInt(643, 663))
                    rhs104_ = d_624_v23_
                    rhs105_ = d_619_v19_
                    rhs106_ = (p0) + (d_622_v21_)
                    rhs107_ = (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_625_v24_, d_625_v24_])) for d_628_i6_ in range(default__.abs(763))])) + (d_627_v26_)
                    d_623_v22_ = rhs103_
                    d_624_v23_ = rhs104_
                    d_619_v19_ = rhs105_
                    d_622_v21_ = rhs106_
                    d_627_v26_ = rhs107_
                    d_629_v27_: _dafny.Map
                    d_629_v27_ = _dafny.Map({d_618_v18_: d_618_v18_})
                    d_630_v28_: _dafny.Map
                    d_630_v28_ = _dafny.Map({(d_629_v27_) | (d_629_v27_): d_624_v23_})
                    d_630_v28_ = (d_630_v28_).set((d_629_v27_) | (d_629_v27_), d_624_v23_)
                    pass
            pass
        d_631_v29_: _dafny.MultiSet
        d_631_v29_ = _dafny.MultiSet([False])
        d_632_v30_: _dafny.Seq
        d_632_v30_ = _dafny.SeqWithoutIsStrInference([(d_631_v29_).cardinality])
        hi5_ = ((self).f7) * ((self).f7)
        for d_633_i7_ in range((d_632_v30_)[default__.safeIndex((self).f8, len(d_632_v30_))], hi5_):
            d_634_v31_: _dafny.Array
            d_635_v32_: bool
            out4_: _dafny.Array
            out5_: bool
            out4_, out5_ = default__.m0(((self).f7) + ((0) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvjdbcbj")), _dafny.SeqWithoutIsStrInference([d_619_v19_ for d_636_i8_ in range(default__.abs(-435))])])).cardinality)), d_618_v18_, (d_618_v18_) == (True), globalState)
            d_634_v31_ = out4_
            d_635_v32_ = out5_
            (globalState).f3 = (d_635_v32_)
            d_594_v0_ = (d_594_v0_) + ((d_594_v0_) + (d_594_v0_))
            rhs108_ = d_632_v30_
            rhs109_ = (self).f12
            lhs85_ = globalState
            d_632_v30_ = rhs108_
            lhs85_.f3 = rhs109_
        d_637_v33_: _dafny.Seq
        d_637_v33_ = _dafny.SeqWithoutIsStrInference([(d_594_v0_ if d_618_v18_ else d_594_v0_)])
        d_594_v0_ = (d_637_v33_)[default__.safeIndex(default__.safeDivisionInt((self).f8, (self).f7), len(d_637_v33_))]

    def m2(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        r2: bool = False
        d_638_v0_: _dafny.Seq
        d_638_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwwkgew"))
        r1 = default__.safeDivisionInt(267, len(d_638_v0_))
        d_639_v1_: str
        d_639_v1_ = _dafny.CodePoint('m')
        d_640_v2_: _dafny.Array
        def lambda26_(d_641_p2_):
            def lambda27_(d_642_i0_):
                return ((0) - (d_641_p2_)) > ((self).f8)

            return lambda27_

        init15_ = lambda26_(p2)
        nw104_ = _dafny.Array(None, 20)
        for i0_15_ in range(nw104_.length(0)):
            nw104_[i0_15_] = init15_(i0_15_)
        d_640_v2_ = nw104_
        rhs110_ = _dafny.CodePoint('y')
        rhs111_ = d_640_v2_
        rhs112_ = (0) - (len(d_638_v0_))
        rhs113_ = 255
        d_639_v1_ = rhs110_
        d_640_v2_ = rhs111_
        r1 = rhs112_
        r1 = rhs113_
        r2 = p3
        d_643_v3_: _dafny.Seq
        d_643_v3_ = _dafny.SeqWithoutIsStrInference([(self).f12, True, p1])
        d_644_v4_: _dafny.MultiSet
        d_644_v4_ = _dafny.MultiSet([not(p1), (d_643_v3_)[default__.safeIndex(p2, len(d_643_v3_))]])
        d_645_v5_: D2
        d_645_v5_ = D2_DC4(d_644_v4_)
        pat_let_tv11_ = d_644_v4_
        def iife46_(_pat_let7_0):
            def iife47_(d_646_dt__update__tmp_h0_):
                def iife48_(_pat_let8_0):
                    def iife49_(d_647_dt__update_hcf3_h0_):
                        return D2_DC4(d_647_dt__update_hcf3_h0_)
                    return iife49_(_pat_let8_0)
                return iife48_(pat_let_tv11_)
            return iife47_(_pat_let7_0)
        d_644_v4_ = (iife46_(d_645_v5_)).cf3
        hi6_ = (0) - ((self).f7)
        for d_648_i1_ in range(p2, hi6_):
            r1 = (0) - (default__.safeModuloInt(default__.safeModuloInt(p0, -574), 444))
            d_649_v6_: bool
            d_649_v6_ = (self).f12
            source15_ = d_649_v6_
            d_650___mcc_h0_ = source15_
            d_651_cf0_ = d_650___mcc_h0_
            r1 = p0
            r1 = (self).f8
            d_652_v7_: _dafny.Map
            d_652_v7_ = _dafny.Map({(self).f8: (self).fm13(d_638_v0_, globalState)})
            r1 = (0) - ((0) - (default__.safeModuloInt(len(d_652_v7_), len((_dafny.SeqWithoutIsStrInference([p1])) + (d_643_v3_)))))
            d_653_v8_: T2
            nw105_ = C2()
            nw105_.ctor__(p2, p2)
            d_653_v8_ = nw105_
            d_653_v8_ = d_653_v8_
            d_654_v9_: D5
            d_654_v9_ = D5_DC12(not((self).fm12(globalState)))
            d_655_v10_: _dafny.MultiSet
            d_655_v10_ = _dafny.MultiSet([-703, p0, len(d_638_v0_)])
            d_656_v11_: _dafny.Seq
            d_656_v11_ = _dafny.SeqWithoutIsStrInference([d_648_i1_])
            rhs114_ = ((self).f7) < (len(default__.fm30(p3, globalState)))
            rhs115_ = not((p3) or ((self).f12))
            rhs116_ = d_654_v9_
            rhs117_ = (d_655_v10_) - ((d_655_v10_) | (_dafny.MultiSet(d_656_v11_)))
            lhs86_ = globalState
            r2 = rhs114_
            lhs86_.f3 = rhs115_
            d_654_v9_ = rhs116_
            d_655_v10_ = rhs117_
            (globalState).f4 = not((False if p1 else not(p1)))
        d_657_v12_: _dafny.Array
        def lambda28_(d_658_v3_, d_659_p1_):
            def lambda29_(d_660_i3_):
                return default__.safeModuloInt(d_660_i3_, len((d_658_v3_).set(default__.safeIndex((self).f7, len(d_658_v3_)), d_659_p1_)))

            return lambda29_

        init16_ = lambda28_(d_643_v3_, p1)
        nw106_ = _dafny.Array(None, 24)
        for i0_16_ in range(nw106_.length(0)):
            nw106_[i0_16_] = init16_(i0_16_)
        d_657_v12_ = nw106_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_657_v12_).length(0)):
            d_661_i2_: int = guard_loop_3_
            if (True) and (((0) <= (d_661_i2_)) and ((d_661_i2_) < ((d_657_v12_).length(0)))):
                (d_657_v12_)[(d_661_i2_)] = default__.safeModuloInt(d_661_i2_, (p2 if (self).f12 else (0) - ((self).f8)))
        r0 = d_643_v3_
        r1 = (p2) * ((self).f7)
        r2 = p3
        return r0, r1, r2

    @property
    def f12(self):
        return self._f12

class C4(T1, T0):
    def  __init__(self):
        self._f7: int = int(0)
        self._f8: int = int(0)
        self._f11: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    def ctor__(self, f11, f7, f8):
        (self)._f11 = f11
        (self)._f7 = f7
        (self)._f8 = f8

    def fm0(self, p0, globalState):
        def iife50_():
            def iife52_():
                coll34_ = _dafny.Set()
                compr_34_: _dafny.Set
                for compr_34_ in (_dafny.MultiSet([_dafny.Set({(self).f11}), _dafny.Set({(self).f11, (self).f11, (self).f11})])).Elements:
                    d_664_v0_: _dafny.Set = compr_34_
                    if (d_664_v0_) in (_dafny.MultiSet([_dafny.Set({(self).f11}), _dafny.Set({(self).f11, (self).f11, (self).f11})])):
                        coll34_ = coll34_.union(_dafny.Set([d_664_v0_]))
                return _dafny.Set(coll34_)
            coll32_ = _dafny.Set()
            def iife51_():
                coll33_ = _dafny.Set()
                compr_33_: _dafny.Set
                for compr_33_ in (_dafny.MultiSet([_dafny.Set({(self).f11}), _dafny.Set({(self).f11, (self).f11, (self).f11})])).Elements:
                    d_662_v0_: _dafny.Set = compr_33_
                    if (d_662_v0_) in (_dafny.MultiSet([_dafny.Set({(self).f11}), _dafny.Set({(self).f11, (self).f11, (self).f11})])):
                        coll33_ = coll33_.union(_dafny.Set([d_662_v0_]))
                return _dafny.Set(coll33_)
            compr_32_: _dafny.Set
            for compr_32_ in (iife51_()
            ).Elements:
                d_663_v1_: _dafny.Set = compr_32_
                if (d_663_v1_) in (iife52_()
                ):
                    coll32_ = coll32_.union(_dafny.Set([d_663_v1_]))
            return _dafny.Set(coll32_)
        return ((_dafny.Set({_dafny.Set({(self).f11})})) | (_dafny.Set({_dafny.Set({(self).f11, (self).f11}), _dafny.Set({(self).f11, False, False, (self).f11})}))) | ((_dafny.Set({_dafny.Set({(self).f11})})) | (iife50_()
        ))

    def fm1(self, p0, p1, p2, p3, globalState):
        return (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_665_i0_ in range(default__.abs(480))]))) < ((self).f7)

    def fm11(self, globalState):
        if (self).f11:
            return (self).f8
        elif True:
            return len((_dafny.SeqWithoutIsStrInference([(self).f7, (self).f7])) + (_dafny.SeqWithoutIsStrInference([453 for d_666_i0_ in range(default__.abs(702))])))

    def m5(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        r2: bool = False
        r3: bool = False
        r1 = p0
        d_667_v0_: C2
        nw107_ = C2()
        nw107_.ctor__(537, -183)
        d_667_v0_ = nw107_
        d_668_v1_: C2
        nw108_ = C2()
        nw108_.ctor__(p2, default__.safeDivisionInt((self).f7, (self).f7))
        d_668_v1_ = nw108_
        d_669_i0_: int
        d_669_i0_ = 0
        with _dafny.label("4"):
            while (p0 if (self).f11 else p0):
                with _dafny.c_label("4"):
                    if (d_669_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_669_i0_ = (d_669_i0_) + (1)
                    d_670_v2_: int
                    d_670_v2_ = 241
                    d_670_v2_ = p2
                    d_671_v3_: C1
                    nw109_ = C1()
                    nw109_.ctor__(default__.safeModuloInt((self).f7, 83), default__.safeModuloInt((self).f7, default__.fm16(True, globalState)))
                    d_671_v3_ = nw109_
                    d_672_v4_: _dafny.Array
                    nw110_ = _dafny.Array(_dafny.Array(None, 0), 9)
                    d_672_v4_ = nw110_
                    d_672_v4_ = d_672_v4_
                    d_673_v5_: _dafny.Map
                    d_673_v5_ = _dafny.Map({p0: p2})
                    d_674_v6_: _dafny.Map
                    d_674_v6_ = _dafny.Map({d_673_v5_: _dafny.SeqWithoutIsStrInference([p0, (self).f11])})
                    d_670_v2_ = default__.safeDivisionInt((0) - (p2), default__.safeDivisionInt(d_670_v2_, len(d_674_v6_)))
                    pass
            pass
        source16_ = default__.fm26(_dafny.MultiSet([p2]), (self).f11, globalState)
        if source16_.is_DC7:
            d_675___mcc_h0_ = source16_.cf10
            d_676_cf10_ = d_675___mcc_h0_
            d_677_v7_: int
            d_677_v7_ = 184
            d_678_v8_: _dafny.Seq
            d_678_v8_ = _dafny.SeqWithoutIsStrInference([(self).f11, False])
            d_679_v9_: _dafny.Seq
            d_679_v9_ = _dafny.SeqWithoutIsStrInference([d_678_v8_, d_678_v8_])
            d_677_v7_ = (default__.safeModuloInt((0) - ((self).f8), d_677_v7_)) * (len((d_679_v9_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhbhuxmc"))), len(d_679_v9_))]))
            d_680_v10_: _dafny.MultiSet
            d_680_v10_ = _dafny.MultiSet([(d_677_v7_) * (p3)])
            d_681_v11_: _dafny.Seq
            d_681_v11_ = _dafny.SeqWithoutIsStrInference([p1])
            d_677_v7_ = (0) - (((d_680_v10_)[p2] if (p2) in (d_680_v10_) else len((d_681_v11_) + (d_681_v11_))))
            d_681_v11_ = d_681_v11_
            d_682_v12_: _dafny.Seq
            d_682_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "md"))
            d_683_v13_: _dafny.Array
            nw111_ = _dafny.Array(None, 10)
            nw111_[int(0)] = (self).f8
            nw111_[int(1)] = p1
            nw111_[int(2)] = (self).f8
            nw111_[int(3)] = -318
            nw111_[int(4)] = p2
            nw111_[int(5)] = d_677_v7_
            nw111_[int(6)] = len(d_682_v12_)
            nw111_[int(7)] = (d_681_v11_)[default__.safeIndex(d_677_v7_, len(d_681_v11_))]
            nw111_[int(8)] = (self).f8
            nw111_[int(9)] = p3
            d_683_v13_ = nw111_
            d_684_v14_: _dafny.Array
            nw112_ = _dafny.Array(None, 1)
            nw112_[int(0)] = d_683_v13_
            d_684_v14_ = nw112_
            index110_ = default__.safeIndex(304, (d_684_v14_).length(0))
            (d_684_v14_)[index110_] = d_683_v13_
            index111_ = default__.safeIndex(304, (d_684_v14_).length(0))
            (d_684_v14_)[index111_] = d_683_v13_
        elif True:
            d_685___mcc_h1_ = source16_.cf9
            d_686_cf9_ = d_685___mcc_h1_
            d_687_v15_: _dafny.Seq
            d_687_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
            d_688_v16_: _dafny.Set
            d_688_v16_ = _dafny.Set({d_687_v15_})
            d_689_v17_: _dafny.Array
            def lambda30_(d_690_p0_):
                def lambda31_(d_691_i1_):
                    return d_690_p0_

                return lambda31_

            init17_ = lambda30_(p0)
            nw113_ = _dafny.Array(None, 25)
            for i0_17_ in range(nw113_.length(0)):
                nw113_[i0_17_] = init17_(i0_17_)
            d_689_v17_ = nw113_
            d_692_v18_: _dafny.Map
            d_692_v18_ = _dafny.Map({d_687_v15_: d_689_v17_})
            d_693_v19_: int
            d_693_v19_ = -109
            d_694_v20_: _dafny.Seq
            d_694_v20_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekpn")): d_689_v17_}), d_692_v18_])
            d_695_v21_: _dafny.Set
            d_695_v21_ = _dafny.Set({p0})
            d_696_v22_: _dafny.MultiSet
            d_696_v22_ = _dafny.MultiSet([(self).f11])
            rhs118_ = (d_688_v16_) - (d_688_v16_)
            rhs119_ = (d_694_v20_)[default__.safeIndex(len(d_695_v21_), len(d_694_v20_))]
            rhs120_ = ((d_696_v22_).intersection(d_696_v22_)).cardinality
            d_688_v16_ = rhs118_
            d_692_v18_ = rhs119_
            d_693_v19_ = rhs120_
            r3 = not((self).f11)
            d_697_v23_: _dafny.Seq
            d_697_v23_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_698_v24_: _dafny.Seq
            d_698_v24_ = _dafny.SeqWithoutIsStrInference([80, (self).f7, p1])
            (globalState).f4 = (((d_697_v23_) + (d_697_v23_)).set(default__.safeIndex(len(d_698_v24_), len((d_697_v23_) + (d_697_v23_))), (self).f11)) == (d_697_v23_)
            d_697_v23_ = (d_697_v23_).set(default__.safeIndex(607, len(d_697_v23_)), p0)
        d_699_v25_: int
        d_699_v25_ = 321
        d_699_v25_ = p2
        r0 = p0
        r1 = (self).f11
        r2 = not (p0) or ((self).f11)
        r3 = not((self).f11)
        return r0, r1, r2, r3

    @property
    def f11(self):
        return self._f11

class C5(T0, T1):
    def  __init__(self):
        self._f7: int = int(0)
        self._f8: int = int(0)
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    def ctor__(self, f9, f10, f7, f8):
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f7 = f7
        (self)._f8 = f8

    def fm0(self, p0, globalState):
        return ((_dafny.Set({_dafny.Set({False, False})})) - (_dafny.Set({_dafny.Set({True, True}), _dafny.Set({False, True})}))) | (_dafny.Set({_dafny.Set({False}), _dafny.Set({True, False, not(True), False, True})}))

    def fm1(self, p0, p1, p2, p3, globalState):
        return (False) and ((_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.SeqWithoutIsStrInference([(self).f7])), (self).f10, (self).f8, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10]))).cardinality, len((self).f9)}), _dafny.Set({(self).f10})])) != (_dafny.SeqWithoutIsStrInference([_dafny.Set({64})])))

    def fm6(self, p0, p1, p2, globalState):
        return (self).f8

    def fm7(self, p0, globalState):
        return ((self).f10) * ((self).f7)

    def m4(self, p0, globalState):
        r0: int = int(0)
        d_700_v0_: _dafny.Array
        def lambda32_(d_701_i0_):
            return (d_701_i0_) * ((self).f7)

        init18_ = lambda32_
        nw114_ = _dafny.Array(None, 11)
        for i0_18_ in range(nw114_.length(0)):
            nw114_[i0_18_] = init18_(i0_18_)
        d_700_v0_ = nw114_
        index112_ = default__.safeIndex(993, (d_700_v0_).length(0))
        (d_700_v0_)[index112_] = (self).f7
        d_702_v1_: bool
        d_702_v1_ = False
        index113_ = default__.safeIndex(993, (d_700_v0_).length(0))
        (d_700_v0_)[index113_] = (self).fm6(_dafny.CodePoint('a'), default__.fm8(d_702_v1_, _dafny.SeqWithoutIsStrInference([(self).f8, p0, p0]), d_702_v1_, globalState), d_702_v1_, globalState)
        d_703_v2_: _dafny.Seq
        d_703_v2_ = _dafny.SeqWithoutIsStrInference([d_702_v1_, d_702_v1_, d_702_v1_])
        d_704_v3_: _dafny.Map
        d_704_v3_ = _dafny.Map({(self).f9: d_702_v1_})
        d_705_v4_: str
        d_705_v4_ = _dafny.CodePoint('g')
        d_706_v5_: _dafny.Map
        d_706_v5_ = _dafny.Map({p0: len((self).f9)})
        d_707_v6_: bool
        d_707_v6_ = d_702_v1_
        d_708_v7_: _dafny.Set
        d_708_v7_ = _dafny.Set({(self).f8, len((self).f9)})
        d_709_v8_: _dafny.Map
        d_709_v8_ = _dafny.Map({d_702_v1_: False})
        d_710_v9_: _dafny.MultiSet
        d_710_v9_ = _dafny.MultiSet([default__.fm9(d_702_v1_, globalState)])
        d_711_v10_: _dafny.Set
        d_711_v10_ = _dafny.Set({d_702_v1_})
        d_712_v11_: _dafny.MultiSet
        d_712_v11_ = _dafny.MultiSet([d_702_v1_, d_702_v1_])
        d_713_v12_: _dafny.Seq
        d_713_v12_ = _dafny.SeqWithoutIsStrInference([(self).f7])
        d_714_v13_: _dafny.Map
        d_714_v13_ = _dafny.Map({((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(d_713_v12_)]))).set(p0, default__.abs((self).f7))).set(p0, default__.abs((self).f8)): d_702_v1_})
        d_715_v14_: _dafny.Array
        nw115_ = _dafny.Array(None, 27)
        nw115_[int(0)] = d_702_v1_
        nw115_[int(1)] = d_702_v1_
        nw115_[int(2)] = (d_703_v2_)[default__.safeIndex((self).f8, len(d_703_v2_))]
        nw115_[int(3)] = ((d_704_v3_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twaaoc"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twaaoc"))) in (d_704_v3_) else d_702_v1_)
        nw115_[int(4)] = d_702_v1_
        nw115_[int(5)] = d_702_v1_
        nw115_[int(6)] = d_702_v1_
        nw115_[int(7)] = d_702_v1_
        nw115_[int(8)] = d_702_v1_
        nw115_[int(9)] = ((0) - ((self).fm6(d_705_v4_, (self).f9, d_702_v1_, globalState))) < ((0) - ((_dafny.MultiSet([d_702_v1_])).cardinality))
        nw115_[int(10)] = not(d_702_v1_)
        nw115_[int(11)] = d_702_v1_
        nw115_[int(12)] = (len(d_703_v2_)) not in (d_706_v5_)
        nw115_[int(13)] = (d_707_v6_)
        nw115_[int(14)] = d_702_v1_
        nw115_[int(15)] = d_702_v1_
        nw115_[int(16)] = (d_708_v7_).issubset(d_708_v7_)
        nw115_[int(17)] = (d_709_v8_) not in (d_710_v9_)
        nw115_[int(18)] = d_702_v1_
        nw115_[int(19)] = d_702_v1_
        nw115_[int(20)] = (d_702_v1_) or (d_702_v1_)
        nw115_[int(21)] = not((d_702_v1_) and (d_702_v1_))
        nw115_[int(22)] = (d_711_v10_).isdisjoint(d_711_v10_)
        nw115_[int(23)] = (d_712_v11_).isdisjoint(_dafny.MultiSet([d_702_v1_, d_702_v1_]))
        nw115_[int(24)] = False
        nw115_[int(25)] = d_702_v1_
        nw115_[int(26)] = ((d_714_v13_)[_dafny.MultiSet([(d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))], (d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))], p0, (self).f7])] if (_dafny.MultiSet([(d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))], (d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))], p0, (self).f7])) in (d_714_v13_) else not(d_702_v1_))
        d_715_v14_ = nw115_
        index114_ = default__.safeIndex(888, (d_715_v14_).length(0))
        (d_715_v14_)[index114_] = (d_702_v1_ if d_702_v1_ else not(d_702_v1_))
        d_716_v15_: _dafny.MultiSet
        d_716_v15_ = _dafny.MultiSet([(self).f8])
        index115_ = default__.safeIndex(993, (d_700_v0_).length(0))
        index116_ = default__.safeIndex(888, (d_715_v14_).length(0))
        rhs121_ = (self).fm7(default__.safeModuloInt(((d_716_v15_).set((d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))], default__.abs((self).f7))).cardinality, p0), globalState)
        rhs122_ = (d_703_v2_)[default__.safeIndex((self).f7, len(d_703_v2_))]
        lhs87_ = d_700_v0_
        lhs88_ = default__.safeIndex(993, (d_700_v0_).length(0))
        lhs89_ = d_715_v14_
        lhs90_ = default__.safeIndex(888, (d_715_v14_).length(0))
        lhs87_[lhs88_] = rhs121_
        lhs89_[lhs90_] = rhs122_
        d_717_v16_: _dafny.Seq
        d_717_v16_ = _dafny.SeqWithoutIsStrInference([d_700_v0_, d_700_v0_])
        index117_ = default__.safeIndex(993, (d_700_v0_).length(0))
        (d_700_v0_)[index117_] = len(((d_717_v16_) + (d_717_v16_)).set(default__.safeIndex((len(_dafny.Map({(d_715_v14_)[default__.safeIndex(888, (d_715_v14_).length(0))]: (d_715_v14_)[default__.safeIndex(888, (d_715_v14_).length(0))]}))) + ((self).f8), len((d_717_v16_) + (d_717_v16_))), d_700_v0_))
        if d_702_v1_:
            d_718_v17_: _dafny.Map
            d_718_v17_ = _dafny.Map({d_705_v4_: (default__.fm10((self).f9, (d_715_v14_)[default__.safeIndex(888, (d_715_v14_).length(0))], (0) - (p0), (self).f7, globalState)).set(default__.safeIndex((self).f7, len(default__.fm10((self).f9, (d_715_v14_)[default__.safeIndex(888, (d_715_v14_).length(0))], (0) - (p0), (self).f7, globalState))), d_702_v1_)})
            d_719_v18_: _dafny.Seq
            d_719_v18_ = _dafny.SeqWithoutIsStrInference([d_707_v6_, d_707_v6_])
            d_718_v17_ = (d_718_v17_).set(d_705_v4_, d_719_v18_)
            d_702_v1_ = ((d_709_v8_)[(d_715_v14_)[default__.safeIndex(888, (d_715_v14_).length(0))]] if ((d_715_v14_)[default__.safeIndex(888, (d_715_v14_).length(0))]) in (d_709_v8_) else not((d_715_v14_)[default__.safeIndex(888, (d_715_v14_).length(0))]))
            index118_ = default__.safeIndex(888, (d_715_v14_).length(0))
            (d_715_v14_)[index118_] = d_702_v1_
            d_720_v19_: _dafny.Map
            d_720_v19_ = _dafny.Map({d_702_v1_: d_715_v14_})
            d_721_v20_: _dafny.Array
            nw116_ = _dafny.Array(None, 15)
            nw116_[int(0)] = d_715_v14_
            nw116_[int(1)] = d_715_v14_
            nw116_[int(2)] = d_715_v14_
            nw116_[int(3)] = d_715_v14_
            nw116_[int(4)] = ((d_720_v19_)[d_702_v1_] if (d_702_v1_) in (d_720_v19_) else d_715_v14_)
            nw116_[int(5)] = d_715_v14_
            nw116_[int(6)] = d_715_v14_
            nw116_[int(7)] = d_715_v14_
            nw116_[int(8)] = d_715_v14_
            nw116_[int(9)] = d_715_v14_
            nw116_[int(10)] = d_715_v14_
            nw116_[int(11)] = d_715_v14_
            nw116_[int(12)] = ((d_720_v19_)[True] if (True) in (d_720_v19_) else d_715_v14_)
            nw116_[int(13)] = d_715_v14_
            nw116_[int(14)] = d_715_v14_
            d_721_v20_ = nw116_
            index119_ = default__.safeIndex(888, (d_721_v20_).length(0))
            (d_721_v20_)[index119_] = d_715_v14_
            index120_ = default__.safeIndex(888, (d_721_v20_).length(0))
            nw117_ = _dafny.Array(False, 23)
            (d_721_v20_)[index120_] = nw117_
            d_715_v14_ = (d_721_v20_)[default__.safeIndex(888, (d_721_v20_).length(0))]
        elif True:
            d_722_v22_: _dafny.Set
            d_722_v22_ = _dafny.Set({(self).f9})
            index121_ = default__.safeIndex(888, (d_715_v14_).length(0))
            def iife53_():
                coll35_ = _dafny.Set()
                compr_35_: _dafny.Seq
                for compr_35_ in (_dafny.SeqWithoutIsStrInference([((self).f9).set(default__.safeIndex((d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))], len((self).f9)), _dafny.CodePoint('f')), (self).f9])).Elements:
                    d_723_v21_: _dafny.Seq = compr_35_
                    if (d_723_v21_) in (_dafny.SeqWithoutIsStrInference([((self).f9).set(default__.safeIndex((d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))], len((self).f9)), _dafny.CodePoint('f')), (self).f9])):
                        coll35_ = coll35_.union(_dafny.Set([d_723_v21_]))
                return _dafny.Set(coll35_)
            rhs123_ = (self).f7
            rhs124_ = (d_722_v22_).ispropersubset(iife53_()
            )
            lhs91_ = d_715_v14_
            lhs92_ = default__.safeIndex(888, (d_715_v14_).length(0))
            r0 = rhs123_
            lhs91_[lhs92_] = rhs124_
            d_724_v23_: C3
            nw118_ = C3()
            nw118_.ctor__((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xs"))) <= ((self).f9), (self).f7, (d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))])
            d_724_v23_ = nw118_
            d_725_v24_: C1
            nw119_ = C1()
            nw119_.ctor__(((self).f7) + ((self).f7), ((0) - (len(d_713_v12_))) - ((self).f7))
            d_725_v24_ = nw119_
            if d_702_v1_:
                d_716_v15_ = _dafny.MultiSet([-599, (self).f8, (self).fm7((d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))], globalState)])
                d_726_v25_: C0
                nw120_ = C0()
                nw120_.ctor__(d_700_v0_, 991, (self).f10)
                d_726_v25_ = nw120_
                d_727_v26_: _dafny.Seq
                d_727_v26_ = _dafny.SeqWithoutIsStrInference([d_726_v25_])
                d_728_v27_: _dafny.Set
                d_728_v27_ = _dafny.Set({d_727_v26_})
                (globalState).f3 = (default__.safeDivisionInt((0) - (len(d_728_v27_)), p0)) >= (default__.safeModuloInt(default__.fm16(d_702_v1_, globalState), (d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))]))
                index122_ = default__.safeIndex(993, (d_700_v0_).length(0))
                (d_700_v0_)[index122_] = default__.safeModuloInt((d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))], 965)
                index123_ = default__.safeIndex(993, (d_700_v0_).length(0))
                (d_700_v0_)[index123_] = 602
                d_729_v28_: C4
                nw121_ = C4()
                nw121_.ctor__((d_724_v23_).f12, (self).f10, p0)
                d_729_v28_ = nw121_
            elif True:
                d_713_v12_ = (d_713_v12_) + (d_713_v12_)
                d_730_v29_: C1
                nw122_ = C1()
                nw122_.ctor__((self).f8, (0) - ((self).fm7((self).f8, globalState)))
                d_730_v29_ = nw122_
                d_731_v30_: _dafny.Map
                d_731_v30_ = _dafny.Map({(d_724_v23_).f12: len(d_703_v2_)})
                (globalState).f4 = ((d_731_v30_).set((d_724_v23_).f12, (self).f10)) != (d_731_v30_)
                (globalState).f4 = (d_715_v14_)[default__.safeIndex(888, (d_715_v14_).length(0))]
                index124_ = default__.safeIndex(993, (d_700_v0_).length(0))
                (d_700_v0_)[index124_] = (self).f10
            d_732_v31_: _dafny.Map
            d_732_v31_ = _dafny.Map({(d_716_v15_) | (d_716_v15_): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvhdv"))})
            d_732_v31_ = (d_732_v31_).set(_dafny.MultiSet(d_713_v12_), ((self).f9) + ((self).f9))
        d_733_v32_: C4
        nw123_ = C4()
        nw123_.ctor__((d_715_v14_)[default__.safeIndex(888, (d_715_v14_).length(0))], (d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))], (self).f10)
        d_733_v32_ = nw123_
        d_734_v33_: _dafny.Seq
        d_734_v33_ = _dafny.SeqWithoutIsStrInference([d_733_v32_, d_733_v32_])
        d_734_v33_ = d_734_v33_
        hi7_ = (self).f7
        for d_735_i1_ in range(p0, hi7_):
            d_736_v34_: _dafny.Seq
            d_736_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjdnlna"))
            d_737_v35_: _dafny.Array
            nw124_ = _dafny.Array(_dafny.Array(None, 0), 14)
            d_737_v35_ = nw124_
            index125_ = default__.safeIndex(690, (d_737_v35_).length(0))
            (d_737_v35_)[index125_] = d_715_v14_
            d_738_v36_: _dafny.Array
            nw125_ = _dafny.Array(None, 2)
            nw125_[int(0)] = _dafny.CodePoint('f')
            nw125_[int(1)] = d_705_v4_
            d_738_v36_ = nw125_
            index126_ = default__.safeIndex(652, (d_738_v36_).length(0))
            (d_738_v36_)[index126_] = d_705_v4_
            d_739_v37_: _dafny.Map
            d_739_v37_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0, (d_716_v15_).cardinality]): (d_715_v14_)[default__.safeIndex(888, (d_715_v14_).length(0))]})
            index127_ = default__.safeIndex(993, (d_700_v0_).length(0))
            index128_ = default__.safeIndex(690, (d_737_v35_).length(0))
            index129_ = default__.safeIndex(652, (d_738_v36_).length(0))
            rhs125_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olimw"))
            rhs126_ = len((d_739_v37_) | (d_739_v37_))
            rhs127_ = ((self).f8) < (default__.safeModuloInt((self).f8, (self).f8))
            rhs128_ = d_715_v14_
            rhs129_ = d_705_v4_
            lhs93_ = d_700_v0_
            lhs94_ = default__.safeIndex(993, (d_700_v0_).length(0))
            lhs95_ = globalState
            lhs96_ = d_737_v35_
            lhs97_ = default__.safeIndex(690, (d_737_v35_).length(0))
            lhs98_ = d_738_v36_
            lhs99_ = default__.safeIndex(652, (d_738_v36_).length(0))
            d_736_v34_ = rhs125_
            lhs93_[lhs94_] = rhs126_
            lhs95_.f3 = rhs127_
            lhs96_[lhs97_] = rhs128_
            lhs98_[lhs99_] = rhs129_
            d_740_v38_: D3
            d_740_v38_ = D3_DC7((d_733_v32_).f11)
            source17_ = d_740_v38_
            if source17_.is_DC7:
                d_741___mcc_h0_ = source17_.cf10
                d_742_cf10_ = d_741___mcc_h0_
                d_743_v39_: _dafny.Set
                d_743_v39_ = _dafny.Set({d_706_v5_, d_706_v5_, d_706_v5_, d_706_v5_, d_706_v5_})
                d_744_v41_: _dafny.Array
                nw126_ = _dafny.Array(None, 15)
                nw126_[int(0)] = default__.fm31((self).f8, _dafny.MultiSet(default__.fm30(d_742_cf10_, globalState)), globalState)
                nw126_[int(1)] = d_743_v39_
                def iife54_():
                    coll36_ = _dafny.Set()
                    compr_36_: _dafny.Map
                    for compr_36_ in (_dafny.SeqWithoutIsStrInference([(d_706_v5_).set(len(_dafny.SeqWithoutIsStrInference([d_735_i1_])), p0) for d_745_i2_ in range(default__.abs(386))])).Elements:
                        d_746_v40_: _dafny.Map = compr_36_
                        if (d_746_v40_) in (_dafny.SeqWithoutIsStrInference([(d_706_v5_).set(len(_dafny.SeqWithoutIsStrInference([d_735_i1_])), p0) for d_745_i2_ in range(default__.abs(386))])):
                            coll36_ = coll36_.union(_dafny.Set([d_746_v40_]))
                    return _dafny.Set(coll36_)
                nw126_[int(2)] = iife54_()
                
                nw126_[int(3)] = d_743_v39_
                nw126_[int(4)] = d_743_v39_
                nw126_[int(5)] = d_743_v39_
                nw126_[int(6)] = d_743_v39_
                nw126_[int(7)] = (d_743_v39_).intersection(d_743_v39_)
                nw126_[int(8)] = d_743_v39_
                nw126_[int(9)] = (default__.fm31((d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))], d_712_v11_, globalState)) - (d_743_v39_)
                nw126_[int(10)] = _dafny.Set({_dafny.Map({(self).f8: (0) - ((self).f10)})})
                nw126_[int(11)] = d_743_v39_
                nw126_[int(12)] = d_743_v39_
                nw126_[int(13)] = d_743_v39_
                nw126_[int(14)] = default__.fm31((self).f7, _dafny.MultiSet(d_703_v2_), globalState)
                d_744_v41_ = nw126_
                index130_ = default__.safeIndex(196, (d_744_v41_).length(0))
                (d_744_v41_)[index130_] = d_743_v39_
                d_747_v42_: _dafny.Seq
                d_747_v42_ = _dafny.SeqWithoutIsStrInference([(d_743_v39_).intersection(d_743_v39_)])
                index131_ = default__.safeIndex(196, (d_744_v41_).length(0))
                index132_ = default__.safeIndex(993, (d_700_v0_).length(0))
                index133_ = default__.safeIndex(652, (d_738_v36_).length(0))
                rhs130_ = (self).f9
                rhs131_ = ((self).f8) < ((self).f10)
                rhs132_ = (d_747_v42_)[default__.safeIndex((0) - ((self).f10), len(d_747_v42_))]
                rhs133_ = ((self).f7) + (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cniy"))), len(((self).f9).set(default__.safeIndex((0) - ((self).f8), len((self).f9)), (d_738_v36_)[default__.safeIndex(652, (d_738_v36_).length(0))]))))
                rhs134_ = d_705_v4_
                lhs100_ = d_744_v41_
                lhs101_ = default__.safeIndex(196, (d_744_v41_).length(0))
                lhs102_ = d_700_v0_
                lhs103_ = default__.safeIndex(993, (d_700_v0_).length(0))
                lhs104_ = d_738_v36_
                lhs105_ = default__.safeIndex(652, (d_738_v36_).length(0))
                d_736_v34_ = rhs130_
                d_702_v1_ = rhs131_
                lhs100_[lhs101_] = rhs132_
                lhs102_[lhs103_] = rhs133_
                lhs104_[lhs105_] = rhs134_
                d_700_v0_ = d_700_v0_
                index134_ = default__.safeIndex(888, (d_715_v14_).length(0))
                rhs135_ = True
                rhs136_ = d_700_v0_
                rhs137_ = _dafny.SeqWithoutIsStrInference([(d_713_v12_)[default__.safeIndex((d_716_v15_).cardinality, len(d_713_v12_))]])
                rhs138_ = ((d_733_v32_).f11) == ((d_715_v14_)[default__.safeIndex(888, (d_715_v14_).length(0))])
                rhs139_ = d_710_v9_
                lhs106_ = d_715_v14_
                lhs107_ = default__.safeIndex(888, (d_715_v14_).length(0))
                lhs106_[lhs107_] = rhs135_
                d_700_v0_ = rhs136_
                d_713_v12_ = rhs137_
                d_742_cf10_ = rhs138_
                d_710_v9_ = rhs139_
                d_704_v3_ = (d_704_v3_).set((self).f9, (d_715_v14_)[default__.safeIndex(888, (d_715_v14_).length(0))])
            elif True:
                d_748___mcc_h1_ = source17_.cf9
                d_749_cf9_ = d_748___mcc_h1_
                (globalState).f4 = d_702_v1_
                d_750_v43_: _dafny.Map
                d_750_v43_ = _dafny.Map({d_708_v7_: ((self).f9) + (d_736_v34_)})
                d_736_v34_ = ((d_750_v43_)[(d_749_cf9_) | (d_749_cf9_)] if ((d_749_cf9_) | (d_749_cf9_)) in (d_750_v43_) else d_736_v34_)
                index135_ = default__.safeIndex(993, (d_700_v0_).length(0))
                (d_700_v0_)[index135_] = d_735_i1_
                (globalState).f4 = (d_733_v32_).f11
            d_751_v44_: D1
            d_751_v44_ = D1_DC1(_dafny.CodePoint('u'))
            d_752_v45_: D1
            d_752_v45_ = D1_DC3(d_751_v44_)
            d_753_v46_: bool
            d_754_v47_: bool
            d_755_v48_: bool
            d_756_v49_: bool
            out6_: bool
            out7_: bool
            out8_: bool
            out9_: bool
            out6_, out7_, out8_, out9_ = (d_733_v32_).m5((d_733_v32_).f11, len(d_713_v12_), 677, len(_dafny.Set({d_752_v45_})), globalState)
            d_753_v46_ = out6_
            d_754_v47_ = out7_
            d_755_v48_ = out8_
            d_756_v49_ = out9_
            r0 = default__.safeDivisionInt(len((d_709_v8_) | (d_709_v8_)), (d_700_v0_)[default__.safeIndex(993, (d_700_v0_).length(0))])
        r0 = (self).f8
        return r0

    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10

class C6(T2, T1, T0):
    def  __init__(self):
        self._f7: int = int(0)
        self._f8: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    def ctor__(self, f7, f8):
        (self)._f7 = f7
        (self)._f8 = f8

    def fm2(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lysobld"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_757_i0_ in range(default__.abs(732))]))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exbusrxu")))

    def fm0(self, p0, globalState):
        return (_dafny.Set({_dafny.Set({False, not(not(True)), not(False)})})).intersection((_dafny.Set({_dafny.Set({True, False})})) - (_dafny.Set({_dafny.Set({False}), _dafny.Set({False, True, False, True, True})})))

    def fm1(self, p0, p1, p2, p3, globalState):
        return (970) <= ((self).f7)

    def fm3(self, p0, p1, p2, p3, globalState):
        if (not(False)):
            return -377
        elif True:
            return len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sv")))

    def fm4(self, p0, p1, globalState):
        return (self).f7

    def m1(self, p0, globalState):
        d_758_v0_: int
        d_758_v0_ = -479
        d_759_v1_: _dafny.Seq
        d_759_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hndygtt"))
        d_760_v2_: bool
        d_760_v2_ = True
        d_761_v3_: _dafny.MultiSet
        d_761_v3_ = _dafny.MultiSet([d_760_v2_])
        rhs140_ = d_760_v2_
        rhs141_ = (((self).f8) * (((d_761_v3_)[d_760_v2_] if (d_760_v2_) in (d_761_v3_) else (self).f8)) if d_760_v2_ else ((self).f7) + ((self).f8))
        rhs142_ = (d_759_v1_).set(default__.safeIndex(default__.safeDivisionInt((self).f7, d_758_v0_), len(d_759_v1_)), default__.fm5(globalState))
        lhs108_ = globalState
        lhs108_.f3 = rhs140_
        d_758_v0_ = rhs141_
        d_759_v1_ = rhs142_
        d_762_v4_: _dafny.Array
        nw127_ = _dafny.Array(int(0), 28)
        d_762_v4_ = nw127_
        d_763_v5_: _dafny.MultiSet
        d_763_v5_ = _dafny.MultiSet([len(d_759_v1_)])
        d_764_v6_: _dafny.Map
        d_764_v6_ = _dafny.Map({d_760_v2_: d_763_v5_})
        d_765_v7_: _dafny.Seq
        d_765_v7_ = _dafny.SeqWithoutIsStrInference([(0) - ((((d_764_v6_)[d_760_v2_] if (d_760_v2_) in (d_764_v6_) else d_763_v5_)).cardinality), (self).f7])
        d_766_v8_: T1
        nw128_ = C0()
        nw128_.ctor__(d_762_v4_, (self).f8, default__.safeModuloInt((self).f8, (d_765_v7_)[default__.safeIndex((self).f7, len(d_765_v7_))]))
        d_766_v8_ = nw128_
        d_766_v8_ = d_766_v8_
        d_767_v9_: _dafny.Seq
        d_767_v9_ = _dafny.SeqWithoutIsStrInference([(d_763_v5_).set(d_758_v0_, default__.abs(d_758_v0_))])
        d_768_i0_: int
        d_768_i0_ = 0
        with _dafny.label("5"):
            while (d_767_v9_) < (_dafny.SeqWithoutIsStrInference([d_763_v5_, d_763_v5_])):
                with _dafny.c_label("5"):
                    if (d_768_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_768_i0_ = (d_768_i0_) + (1)
                    if ((p0) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_769_i1_ in range(default__.abs(938))])) if d_760_v2_ else (d_760_v2_ if d_760_v2_ else d_760_v2_)):
                        d_770_v10_: str
                        d_770_v10_ = _dafny.CodePoint('r')
                        d_759_v1_ = ((d_759_v1_) + (d_759_v1_)).set(default__.safeIndex(default__.safeDivisionInt((d_766_v8_).f7, (self).f7), len((d_759_v1_) + (d_759_v1_))), d_770_v10_)
                        d_771_v11_: _dafny.Array
                        def lambda33_(d_772_p0_, d_773_v1_):
                            def lambda34_(d_774_i2_):
                                return (d_772_p0_) == (d_773_v1_)

                            return lambda34_

                        init19_ = lambda33_(p0, d_759_v1_)
                        nw129_ = _dafny.Array(None, 26)
                        for i0_19_ in range(nw129_.length(0)):
                            nw129_[i0_19_] = init19_(i0_19_)
                        d_771_v11_ = nw129_
                        d_775_v12_: D5
                        d_775_v12_ = D5_DC13(True, d_770_v10_)
                        d_776_v13_: D6
                        d_776_v13_ = D6_DC15(d_771_v11_)
                        rhs143_ = (d_776_v13_).cf23
                        rhs144_ = D5_DC13(d_760_v2_, d_770_v10_)
                        d_771_v11_ = rhs143_
                        d_775_v12_ = rhs144_
                        d_777_v14_: _dafny.Array
                        nw130_ = _dafny.Array(D2.default()(), 9)
                        d_777_v14_ = nw130_
                        d_778_v15_: _dafny.Map
                        d_778_v15_ = _dafny.Map({d_760_v2_: d_777_v14_})
                        d_777_v14_ = ((d_778_v15_)[d_760_v2_] if (d_760_v2_) in (d_778_v15_) else d_777_v14_)
                        d_779_v16_: _dafny.Set
                        d_779_v16_ = _dafny.Set({d_770_v10_, d_770_v10_})
                        d_779_v16_ = (d_779_v16_) | (d_779_v16_)
                        d_780_v17_: _dafny.Set
                        d_780_v17_ = _dafny.Set({default__.fm8(d_760_v2_, d_765_v7_, d_760_v2_, globalState), (p0) + (p0)})
                        d_781_v18_: _dafny.Map
                        d_781_v18_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_770_v10_ for d_782_i3_ in range(default__.abs(-863))]): _dafny.SeqWithoutIsStrInference([d_760_v2_])})
                        def iife55_():
                            coll37_ = _dafny.Set()
                            compr_37_: _dafny.Seq
                            for compr_37_ in (d_781_v18_).keys.Elements:
                                d_783_v19_: _dafny.Seq = compr_37_
                                if (d_783_v19_) in (d_781_v18_):
                                    coll37_ = coll37_.union(_dafny.Set([d_783_v19_]))
                            return _dafny.Set(coll37_)
                        d_780_v17_ = ((d_780_v17_) - (_dafny.Set({d_759_v1_}))) | (iife55_()
                        )
                    elif True:
                        index136_ = default__.safeIndex(65, (d_762_v4_).length(0))
                        (d_762_v4_)[index136_] = ((d_761_v3_)[not(d_760_v2_)] if (not(d_760_v2_)) in (d_761_v3_) else (self).f7)
                        index137_ = default__.safeIndex(65, (d_762_v4_).length(0))
                        (d_762_v4_)[index137_] = default__.safeDivisionInt(default__.fm20(d_760_v2_, d_760_v2_, d_760_v2_, globalState), (d_766_v8_).f8)
                        d_784_v20_: D7
                        d_784_v20_ = D7_DC17(p0)
                        d_785_v21_: _dafny.Seq
                        d_785_v21_ = _dafny.SeqWithoutIsStrInference([d_759_v1_, default__.fm8(d_760_v2_, d_765_v7_, d_760_v2_, globalState)])
                        d_786_v22_: _dafny.Set
                        d_786_v22_ = _dafny.Set({d_759_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))})
                        d_787_v23_: _dafny.Array
                        nw131_ = _dafny.Array(None, 21)
                        nw131_[int(0)] = d_759_v1_
                        nw131_[int(1)] = d_759_v1_
                        nw131_[int(2)] = d_759_v1_
                        nw131_[int(3)] = (d_759_v1_) + (d_759_v1_)
                        nw131_[int(4)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_788_i4_ in range(default__.abs(362))])
                        nw131_[int(5)] = (d_759_v1_) + (p0)
                        nw131_[int(6)] = (d_784_v20_).cf26
                        nw131_[int(7)] = (p0) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_789_i5_ in range(default__.abs(10))]))
                        nw131_[int(8)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_790_i6_ in range(default__.abs(805))])
                        nw131_[int(9)] = d_759_v1_
                        nw131_[int(10)] = p0
                        nw131_[int(11)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_791_i7_ in range(default__.abs(519))])
                        nw131_[int(12)] = ((d_784_v20_).cf26) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klcneosa")))
                        nw131_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbsipip"))
                        nw131_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wudxpkey"))
                        nw131_[int(15)] = p0
                        nw131_[int(16)] = (d_759_v1_) + ((d_785_v21_)[default__.safeIndex(len(d_786_v22_), len(d_785_v21_))])
                        nw131_[int(17)] = d_759_v1_
                        nw131_[int(18)] = (d_759_v1_) + (d_759_v1_)
                        nw131_[int(19)] = (p0).set(default__.safeIndex((d_763_v5_).cardinality, len(p0)), default__.fm5(globalState))
                        nw131_[int(20)] = d_759_v1_
                        d_787_v23_ = nw131_
                        index138_ = default__.safeIndex(153, (d_787_v23_).length(0))
                        (d_787_v23_)[index138_] = d_759_v1_
                        d_792_v24_: _dafny.Seq
                        d_792_v24_ = _dafny.SeqWithoutIsStrInference([d_760_v2_])
                        d_793_v25_: _dafny.Map
                        d_793_v25_ = _dafny.Map({d_758_v0_: (D5_DC12(d_760_v2_)).cf19})
                        d_794_v26_: _dafny.Map
                        d_794_v26_ = _dafny.Map({len(d_792_v24_): len(d_793_v25_)})
                        d_795_v27_: str
                        d_795_v27_ = _dafny.CodePoint('j')
                        index139_ = default__.safeIndex(153, (d_787_v23_).length(0))
                        (d_787_v23_)[index139_] = ((d_759_v1_).set(default__.safeIndex(len(d_794_v26_), len(d_759_v1_)), d_795_v27_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhe")))
                        d_796_v28_: _dafny.Array
                        nw132_ = _dafny.Array(False, 12)
                        d_796_v28_ = nw132_
                        index140_ = default__.safeIndex(300, (d_796_v28_).length(0))
                        (d_796_v28_)[index140_] = d_760_v2_
                        index141_ = default__.safeIndex(300, (d_796_v28_).length(0))
                        (d_796_v28_)[index141_] = (d_760_v2_ if not(not (d_760_v2_) or (d_760_v2_)) else d_760_v2_)
                        d_797_v29_: _dafny.Array
                        def lambda35_(d_798_v24_):
                            def lambda36_(d_799_i8_):
                                return _dafny.SeqWithoutIsStrInference([d_798_v24_])

                            return lambda36_

                        init20_ = lambda35_(d_792_v24_)
                        nw133_ = _dafny.Array(None, 9)
                        for i0_20_ in range(nw133_.length(0)):
                            nw133_[i0_20_] = init20_(i0_20_)
                        d_797_v29_ = nw133_
                        d_800_v30_: _dafny.Seq
                        d_800_v30_ = _dafny.SeqWithoutIsStrInference([d_792_v24_, d_792_v24_])
                        index142_ = default__.safeIndex(565, (d_797_v29_).length(0))
                        (d_797_v29_)[index142_] = (d_800_v30_) + (d_800_v30_)
                        d_801_v31_: D5
                        d_801_v31_ = D5_DC13((d_796_v28_)[default__.safeIndex(300, (d_796_v28_).length(0))], d_795_v27_)
                        d_802_v32_: _dafny.Map
                        d_802_v32_ = _dafny.Map({d_801_v31_: d_795_v27_})
                        d_803_v33_: D7
                        d_803_v33_ = D7_DC20(d_795_v27_, len((d_802_v32_ if True else d_802_v32_)))
                        index143_ = default__.safeIndex(65, (d_762_v4_).length(0))
                        index144_ = default__.safeIndex(565, (d_797_v29_).length(0))
                        rhs145_ = (d_766_v8_).f7
                        rhs146_ = d_800_v30_
                        rhs147_ = d_760_v2_
                        rhs148_ = D7_DC20(d_795_v27_, (0) - ((0) - ((d_766_v8_).f8)))
                        rhs149_ = not((d_796_v28_)[default__.safeIndex(300, (d_796_v28_).length(0))])
                        lhs109_ = d_762_v4_
                        lhs110_ = default__.safeIndex(65, (d_762_v4_).length(0))
                        lhs111_ = d_797_v29_
                        lhs112_ = default__.safeIndex(565, (d_797_v29_).length(0))
                        lhs113_ = globalState
                        lhs114_ = globalState
                        lhs109_[lhs110_] = rhs145_
                        lhs111_[lhs112_] = rhs146_
                        lhs113_.f3 = rhs147_
                        d_803_v33_ = rhs148_
                        lhs114_.f4 = rhs149_
                        d_758_v0_ = (d_766_v8_).f7
                    d_804_v34_: str
                    d_804_v34_ = _dafny.CodePoint('k')
                    d_805_v35_: D5
                    d_805_v35_ = D5_DC13(d_760_v2_, d_804_v34_)
                    source18_ = d_805_v35_
                    if source18_.is_DC12:
                        d_806___mcc_h0_ = source18_.cf19
                        d_807_cf19_ = d_806___mcc_h0_
                        index145_ = default__.safeIndex(170, (d_762_v4_).length(0))
                        (d_762_v4_)[index145_] = (0) - ((d_766_v8_).f7)
                        index146_ = default__.safeIndex(170, (d_762_v4_).length(0))
                        (d_762_v4_)[index146_] = (self).f8
                        d_758_v0_ = (d_766_v8_).f7
                        d_808_v36_: _dafny.Map
                        d_808_v36_ = _dafny.Map({d_807_cf19_: (d_762_v4_)[default__.safeIndex(170, (d_762_v4_).length(0))]})
                        index147_ = default__.safeIndex(170, (d_762_v4_).length(0))
                        (d_762_v4_)[index147_] = default__.safeDivisionInt(len(d_808_v36_), (d_762_v4_)[default__.safeIndex(170, (d_762_v4_).length(0))])
                        d_809_v37_: _dafny.Array
                        def lambda37_(d_810_v2_):
                            def lambda38_(d_811_i9_):
                                return d_810_v2_

                            return lambda38_

                        init21_ = lambda37_(d_760_v2_)
                        nw134_ = _dafny.Array(None, 19)
                        for i0_21_ in range(nw134_.length(0)):
                            nw134_[i0_21_] = init21_(i0_21_)
                        d_809_v37_ = nw134_
                        d_812_v38_: _dafny.Seq
                        d_812_v38_ = _dafny.SeqWithoutIsStrInference([d_809_v37_])
                        d_813_v39_: D4
                        d_813_v39_ = D4_DC9((self).f8, d_812_v38_, (0) - ((self).f8), (0) - ((d_766_v8_).f8))
                        pat_let_tv12_ = d_758_v0_
                        pat_let_tv13_ = d_758_v0_
                        d_814_v40_: _dafny.Map
                        def iife56_(_pat_let9_0):
                            def iife57_(d_815_dt__update__tmp_h0_):
                                def iife58_(_pat_let10_0):
                                    def iife59_(d_816_dt__update_hcf15_h0_):
                                        def iife60_(_pat_let11_0):
                                            def iife61_(d_817_dt__update_hcf12_h0_):
                                                return D4_DC9(d_817_dt__update_hcf12_h0_, (d_815_dt__update__tmp_h0_).cf13, (d_815_dt__update__tmp_h0_).cf14, d_816_dt__update_hcf15_h0_)
                                            return iife61_(_pat_let11_0)
                                        return iife60_(pat_let_tv13_)
                                    return iife59_(_pat_let10_0)
                                return iife58_(pat_let_tv12_)
                            return iife57_(_pat_let9_0)
                        d_814_v40_ = _dafny.Map({iife56_(d_813_v39_): (d_766_v8_).f7})
                        d_818_v41_: _dafny.Map
                        d_818_v41_ = _dafny.Map({(self).f8: d_807_cf19_})
                        d_814_v40_ = (d_814_v40_).set(d_813_v39_, len(d_818_v41_))
                    elif source18_.is_DC13:
                        d_819___mcc_h1_ = source18_.cf20
                        d_820___mcc_h2_ = source18_.cf21
                        d_821_cf21_ = d_820___mcc_h2_
                        d_822_cf20_ = d_819___mcc_h1_
                        d_823_v42_: _dafny.Set
                        d_823_v42_ = _dafny.Set({not(d_822_cf20_)})
                        d_824_v43_: _dafny.Set
                        d_824_v43_ = _dafny.Set({_dafny.Set({not(d_760_v2_)}), d_823_v42_, (d_823_v42_) - (d_823_v42_)})
                        d_824_v43_ = (_dafny.Set({d_823_v42_, d_823_v42_, d_823_v42_, _dafny.Set({d_822_cf20_})})) | (d_824_v43_)
                        d_825_v44_: _dafny.Array
                        nw135_ = _dafny.Array(None, 15)
                        nw135_[int(0)] = d_822_cf20_
                        nw135_[int(1)] = d_822_cf20_
                        nw135_[int(2)] = d_760_v2_
                        nw135_[int(3)] = d_760_v2_
                        nw135_[int(4)] = d_822_cf20_
                        nw135_[int(5)] = d_822_cf20_
                        nw135_[int(6)] = True
                        nw135_[int(7)] = True
                        nw135_[int(8)] = d_822_cf20_
                        nw135_[int(9)] = d_760_v2_
                        nw135_[int(10)] = d_822_cf20_
                        nw135_[int(11)] = True
                        nw135_[int(12)] = d_822_cf20_
                        nw135_[int(13)] = True
                        nw135_[int(14)] = d_822_cf20_
                        d_825_v44_ = nw135_
                        d_826_v45_: _dafny.MultiSet
                        d_826_v45_ = _dafny.MultiSet([d_825_v44_])
                        (globalState).f3 = (d_826_v45_).issubset((d_826_v45_ if d_822_cf20_ else d_826_v45_))
                        index148_ = default__.safeIndex(172, (d_762_v4_).length(0))
                        (d_762_v4_)[index148_] = (self).f8
                        index149_ = default__.safeIndex(172, (d_762_v4_).length(0))
                        (d_762_v4_)[index149_] = (0) - (((0) - ((d_761_v3_).cardinality)) * ((d_766_v8_).f8))
                        d_827_v46_: _dafny.Array
                        nw136_ = _dafny.Array(None, 17)
                        nw136_[int(0)] = d_825_v44_
                        nw136_[int(1)] = d_825_v44_
                        nw136_[int(2)] = d_825_v44_
                        nw136_[int(3)] = d_825_v44_
                        nw136_[int(4)] = d_825_v44_
                        nw136_[int(5)] = d_825_v44_
                        nw136_[int(6)] = d_825_v44_
                        nw136_[int(7)] = d_825_v44_
                        nw136_[int(8)] = d_825_v44_
                        nw136_[int(9)] = d_825_v44_
                        nw136_[int(10)] = d_825_v44_
                        nw136_[int(11)] = d_825_v44_
                        nw136_[int(12)] = d_825_v44_
                        nw136_[int(13)] = d_825_v44_
                        nw136_[int(14)] = d_825_v44_
                        nw136_[int(15)] = d_825_v44_
                        nw136_[int(16)] = d_825_v44_
                        d_827_v46_ = nw136_
                        index150_ = default__.safeIndex(149, (d_827_v46_).length(0))
                        (d_827_v46_)[index150_] = d_825_v44_
                        index151_ = default__.safeIndex(149, (d_827_v46_).length(0))
                        (d_827_v46_)[index151_] = (d_825_v44_ if (d_823_v42_).ispropersubset(_dafny.Set({d_822_cf20_, d_760_v2_, d_760_v2_, d_760_v2_, d_822_cf20_})) else d_825_v44_)
                    elif source18_.is_DC11:
                        d_828___mcc_h3_ = source18_.cf18
                        d_829_cf18_ = d_828___mcc_h3_
                        d_830_v47_: C5
                        nw137_ = C5()
                        nw137_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gk")), (d_766_v8_).f8, (0) - ((self).fm3(d_760_v2_, d_760_v2_, d_760_v2_, d_760_v2_, globalState)), (d_766_v8_).f7)
                        d_830_v47_ = nw137_
                        d_830_v47_ = d_830_v47_
                        d_831_v48_: _dafny.Seq
                        d_831_v48_ = _dafny.SeqWithoutIsStrInference([d_765_v7_])
                        d_831_v48_ = d_831_v48_
                        d_832_v49_: T2
                        nw138_ = C3()
                        nw138_.ctor__(d_760_v2_, default__.safeDivisionInt(len(p0), (d_766_v8_).f7), (0) - ((d_766_v8_).f8))
                        d_832_v49_ = nw138_
                        nw139_ = C2()
                        nw139_.ctor__(922, (0) - ((d_766_v8_).f8))
                        d_832_v49_ = nw139_
                        d_833_v50_: int
                        d_834_v51_: int
                        d_835_v52_: bool
                        out10_: int
                        out11_: int
                        out12_: bool
                        out10_, out11_, out12_ = (self).m3(((self).f7) > ((d_832_v49_).f8), d_758_v0_, (d_762_v4_ if d_760_v2_ else d_762_v4_), (self).fm4(d_760_v2_, d_760_v2_, globalState), globalState)
                        d_833_v50_ = out10_
                        d_834_v51_ = out11_
                        d_835_v52_ = out12_
                    elif True:
                        d_836___mcc_h4_ = source18_.cf22
                        d_837_cf22_ = d_836___mcc_h4_
                        d_838_v53_: D1
                        d_838_v53_ = D1_DC2()
                        d_839_v54_: _dafny.Seq
                        d_839_v54_ = _dafny.SeqWithoutIsStrInference([d_838_v53_])
                        d_839_v54_ = d_839_v54_
                        d_840_v55_: _dafny.Map
                        d_840_v55_ = _dafny.Map({d_804_v34_: (d_766_v8_).f8})
                        d_758_v0_ = len(d_840_v55_)
                        d_758_v0_ = d_758_v0_
                        d_758_v0_ = len(p0)
                    d_841_v56_: _dafny.Map
                    d_841_v56_ = _dafny.Map({d_758_v0_: d_760_v2_})
                    d_842_v57_: _dafny.Map
                    d_842_v57_ = _dafny.Map({d_841_v56_: d_762_v4_})
                    d_843_v58_: _dafny.Map
                    d_843_v58_ = _dafny.Map({d_762_v4_: d_804_v34_})
                    d_844_v59_: _dafny.Map
                    d_844_v59_ = _dafny.Map({d_804_v34_: d_762_v4_})
                    d_845_v60_: _dafny.Array
                    nw140_ = _dafny.Array(None, 14)
                    nw140_[int(0)] = _dafny.Map({((d_842_v57_)[_dafny.Map({len(d_765_v7_): d_760_v2_})] if (_dafny.Map({len(d_765_v7_): d_760_v2_})) in (d_842_v57_) else d_762_v4_): d_804_v34_})
                    nw140_[int(1)] = d_843_v58_
                    nw140_[int(2)] = d_843_v58_
                    nw140_[int(3)] = (d_843_v58_) | (d_843_v58_)
                    nw140_[int(4)] = d_843_v58_
                    nw140_[int(5)] = d_843_v58_
                    nw140_[int(6)] = _dafny.Map({d_762_v4_: d_804_v34_})
                    nw140_[int(7)] = (d_843_v58_).set(((d_844_v59_)[d_804_v34_] if (d_804_v34_) in (d_844_v59_) else d_762_v4_), d_804_v34_)
                    nw140_[int(8)] = (d_843_v58_) | (_dafny.Map({d_762_v4_: d_804_v34_}))
                    nw140_[int(9)] = d_843_v58_
                    nw140_[int(10)] = (d_843_v58_ if d_760_v2_ else d_843_v58_)
                    nw140_[int(11)] = (_dafny.Map({d_762_v4_: d_804_v34_})) | (d_843_v58_)
                    nw140_[int(12)] = d_843_v58_
                    nw140_[int(13)] = _dafny.Map({d_762_v4_: d_804_v34_})
                    d_845_v60_ = nw140_
                    index152_ = default__.safeIndex(75, (d_845_v60_).length(0))
                    (d_845_v60_)[index152_] = _dafny.Map({d_762_v4_: d_804_v34_})
                    index153_ = default__.safeIndex(75, (d_845_v60_).length(0))
                    (d_845_v60_)[index153_] = d_843_v58_
                    d_846_v61_: _dafny.Seq
                    d_846_v61_ = _dafny.SeqWithoutIsStrInference([d_762_v4_, d_762_v4_])
                    d_847_v62_: _dafny.Array
                    nw141_ = _dafny.Array(None, 17)
                    nw141_[int(0)] = d_762_v4_
                    nw141_[int(1)] = d_762_v4_
                    nw141_[int(2)] = d_762_v4_
                    nw141_[int(3)] = d_762_v4_
                    nw141_[int(4)] = d_762_v4_
                    nw141_[int(5)] = d_762_v4_
                    nw141_[int(6)] = d_762_v4_
                    nw141_[int(7)] = d_762_v4_
                    nw141_[int(8)] = d_762_v4_
                    nw141_[int(9)] = d_762_v4_
                    nw141_[int(10)] = d_762_v4_
                    nw141_[int(11)] = (d_846_v61_)[default__.safeIndex((self).f8, len(d_846_v61_))]
                    nw141_[int(12)] = d_762_v4_
                    nw141_[int(13)] = d_762_v4_
                    nw141_[int(14)] = d_762_v4_
                    nw141_[int(15)] = d_762_v4_
                    nw141_[int(16)] = d_762_v4_
                    d_847_v62_ = nw141_
                    d_848_v63_: D4
                    d_848_v63_ = D4_DC8(d_847_v62_)
                    source19_ = d_848_v63_
                    if source19_.is_DC9:
                        d_849___mcc_h5_ = source19_.cf12
                        d_850___mcc_h6_ = source19_.cf13
                        d_851___mcc_h7_ = source19_.cf14
                        d_852___mcc_h8_ = source19_.cf15
                        d_853_cf15_ = d_852___mcc_h8_
                        d_854_cf14_ = d_851___mcc_h7_
                        d_855_cf13_ = d_850___mcc_h6_
                        d_856_cf12_ = d_849___mcc_h5_
                        d_857_v64_: _dafny.Seq
                        d_857_v64_ = _dafny.SeqWithoutIsStrInference([d_759_v1_, _dafny.SeqWithoutIsStrInference([d_804_v34_, d_804_v34_]), d_759_v1_])
                        d_854_cf14_ = default__.safeModuloInt(((d_766_v8_).f7) * ((d_761_v3_).cardinality), (((d_761_v3_)[d_760_v2_] if (d_760_v2_) in (d_761_v3_) else (d_766_v8_).f8)) + (len(d_857_v64_)))
                        d_858_v65_: _dafny.Map
                        d_858_v65_ = _dafny.Map({d_760_v2_: 647})
                        d_859_v66_: _dafny.Set
                        d_859_v66_ = _dafny.Set({(d_766_v8_).f8, default__.safeDivisionInt((d_766_v8_).f7, 901), (0) - ((0) - (((d_858_v65_)[d_760_v2_] if (d_760_v2_) in (d_858_v65_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcam")))))), 330})
                        d_859_v66_ = ((d_859_v66_).intersection(_dafny.Set({(d_766_v8_).f7}))) - (d_859_v66_)
                        d_860_v67_: _dafny.Array
                        nw142_ = _dafny.Array(False, 20)
                        d_860_v67_ = nw142_
                        d_860_v67_ = (d_860_v67_ if (d_760_v2_) == (d_760_v2_) else d_860_v67_)
                        index154_ = default__.safeIndex(29, (d_860_v67_).length(0))
                        (d_860_v67_)[index154_] = d_760_v2_
                        index155_ = default__.safeIndex(29, (d_860_v67_).length(0))
                        (d_860_v67_)[index155_] = not (d_760_v2_) or (d_760_v2_)
                    elif source19_.is_DC10:
                        d_861___mcc_h9_ = source19_.cf16
                        d_862___mcc_h10_ = source19_.cf17
                        d_863_cf17_ = d_862___mcc_h10_
                        d_864_cf16_ = d_861___mcc_h9_
                        d_865_v68_: C5
                        nw143_ = C5()
                        nw143_.ctor__(d_759_v1_, 386, (d_763_v5_).cardinality, (d_766_v8_).f7)
                        d_865_v68_ = nw143_
                        d_866_v69_: _dafny.Map
                        d_866_v69_ = _dafny.Map({d_864_cf16_: d_865_v68_})
                        d_867_v70_: _dafny.Map
                        d_867_v70_ = _dafny.Map({(self).f8: d_866_v69_})
                        d_867_v70_ = (d_867_v70_).set((d_766_v8_).f7, (_dafny.Map({(self).f7: d_865_v68_})).set(358, d_865_v68_))
                        d_868_v71_: _dafny.Map
                        d_868_v71_ = _dafny.Map({not(d_760_v2_): (d_865_v68_).f9})
                        d_869_v72_: _dafny.Map
                        d_869_v72_ = _dafny.Map({True: d_760_v2_})
                        d_870_v73_: D7
                        d_870_v73_ = D7_DC20(d_804_v34_, (d_865_v68_).f10)
                        d_871_v74_: _dafny.Seq
                        d_871_v74_ = _dafny.SeqWithoutIsStrInference([d_870_v73_])
                        d_872_v75_: C0
                        nw144_ = C0()
                        nw144_.ctor__(d_762_v4_, 741, (_dafny.MultiSet([d_863_cf17_])).cardinality)
                        d_872_v75_ = nw144_
                        d_873_v76_: _dafny.Map
                        d_873_v76_ = _dafny.Map({d_762_v4_: d_872_v75_})
                        d_874_v77_: _dafny.Array
                        nw145_ = _dafny.Array(None, 26)
                        nw145_[int(0)] = d_760_v2_
                        nw145_[int(1)] = d_760_v2_
                        nw145_[int(2)] = not(((d_865_v68_).f9) <= (((d_868_v71_)[d_863_cf17_] if (d_863_cf17_) in (d_868_v71_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejyl")))))
                        nw145_[int(3)] = d_760_v2_
                        nw145_[int(4)] = False
                        nw145_[int(5)] = ((self).f8) >= (-50)
                        nw145_[int(6)] = (D5_DC12(d_760_v2_)).cf19
                        nw145_[int(7)] = (d_863_cf17_) and (d_863_cf17_)
                        nw145_[int(8)] = (d_760_v2_) == (d_760_v2_)
                        nw145_[int(9)] = (False) in (d_869_v72_)
                        nw145_[int(10)] = False
                        nw145_[int(11)] = not((_dafny.CodePoint('q')) not in (d_759_v1_))
                        nw145_[int(12)] = (d_871_v74_) < (d_871_v74_)
                        nw145_[int(13)] = d_760_v2_
                        nw145_[int(14)] = ((d_841_v56_)[(d_766_v8_).f7] if ((d_766_v8_).f7) in (d_841_v56_) else ((d_841_v56_)[-911] if (-911) in (d_841_v56_) else d_760_v2_))
                        nw145_[int(15)] = d_863_cf17_
                        nw145_[int(16)] = d_760_v2_
                        nw145_[int(17)] = d_760_v2_
                        nw145_[int(18)] = (default__.fm32(globalState)).cf20
                        nw145_[int(19)] = d_863_cf17_
                        nw145_[int(20)] = not(d_863_cf17_)
                        nw145_[int(21)] = (391) >= ((d_766_v8_).f7)
                        nw145_[int(22)] = d_863_cf17_
                        nw145_[int(23)] = (d_762_v4_) in (d_873_v76_)
                        nw145_[int(24)] = d_760_v2_
                        nw145_[int(25)] = d_863_cf17_
                        d_874_v77_ = nw145_
                        d_874_v77_ = d_874_v77_
                        (globalState).f4 = d_760_v2_
                        d_875_v78_: _dafny.Set
                        d_875_v78_ = _dafny.Set({d_759_v1_})
                        index156_ = default__.safeIndex(432, (d_874_v77_).length(0))
                        (d_874_v77_)[index156_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhndgwb"))) in (d_875_v78_)
                        index157_ = default__.safeIndex(814, (d_762_v4_).length(0))
                        (d_762_v4_)[index157_] = default__.safeModuloInt((0) - (((d_763_v5_)[(d_766_v8_).f8] if ((d_766_v8_).f8) in (d_763_v5_) else 718)), d_864_cf16_)
                        index158_ = default__.safeIndex(432, (d_874_v77_).length(0))
                        index159_ = default__.safeIndex(814, (d_762_v4_).length(0))
                        rhs150_ = not((default__.safeDivisionInt((d_766_v8_).f7, (d_766_v8_).f8)) < ((d_766_v8_).f8))
                        rhs151_ = default__.fm5(globalState)
                        rhs152_ = (d_766_v8_).f8
                        rhs153_ = (d_766_v8_).f8
                        rhs154_ = (d_864_cf16_) == ((d_766_v8_).f7)
                        lhs115_ = d_874_v77_
                        lhs116_ = default__.safeIndex(432, (d_874_v77_).length(0))
                        lhs117_ = d_762_v4_
                        lhs118_ = default__.safeIndex(814, (d_762_v4_).length(0))
                        lhs119_ = globalState
                        lhs115_[lhs116_] = rhs150_
                        d_804_v34_ = rhs151_
                        d_864_cf16_ = rhs152_
                        lhs117_[lhs118_] = rhs153_
                        lhs119_.f4 = rhs154_
                    elif True:
                        d_876___mcc_h11_ = source19_.cf11
                        d_877_cf11_ = d_876___mcc_h11_
                        d_878_v79_: _dafny.Set
                        d_878_v79_ = _dafny.Set({(d_766_v8_).f8})
                        d_879_v80_: _dafny.Map
                        d_879_v80_ = _dafny.Map({(d_759_v1_)[default__.safeIndex((self).f7, len(d_759_v1_))]: d_878_v79_})
                        d_878_v79_ = ((d_879_v80_)[d_804_v34_] if (d_804_v34_) in (d_879_v80_) else d_878_v79_)
                        d_758_v0_ = ((d_766_v8_).f7) - ((d_766_v8_).f7)
                        d_880_v81_: _dafny.MultiSet
                        d_880_v81_ = _dafny.MultiSet([d_804_v34_])
                        (globalState).f4 = (d_760_v2_ if d_760_v2_ else (d_804_v34_) not in (d_880_v81_))
                        d_881_v82_: _dafny.Array
                        nw146_ = _dafny.Array(False, 12)
                        d_881_v82_ = nw146_
                        d_882_v83_: _dafny.Array
                        nw147_ = _dafny.Array(None, 4)
                        nw147_[int(0)] = d_881_v82_
                        nw147_[int(1)] = (d_881_v82_ if d_760_v2_ else d_881_v82_)
                        nw147_[int(2)] = d_881_v82_
                        nw147_[int(3)] = d_881_v82_
                        d_882_v83_ = nw147_
                        index160_ = default__.safeIndex(381, (d_882_v83_).length(0))
                        (d_882_v83_)[index160_] = d_881_v82_
                        d_883_v84_: _dafny.Map
                        d_883_v84_ = _dafny.Map({not(d_760_v2_): d_760_v2_})
                        index161_ = default__.safeIndex(381, (d_882_v83_).length(0))
                        rhs155_ = d_762_v4_
                        rhs156_ = 151
                        rhs157_ = d_881_v82_
                        rhs158_ = d_758_v0_
                        rhs159_ = ((d_883_v84_)[d_760_v2_] if (d_760_v2_) in (d_883_v84_) else d_760_v2_)
                        lhs120_ = d_882_v83_
                        lhs121_ = default__.safeIndex(381, (d_882_v83_).length(0))
                        lhs122_ = globalState
                        d_762_v4_ = rhs155_
                        d_758_v0_ = rhs156_
                        lhs120_[lhs121_] = rhs157_
                        d_758_v0_ = rhs158_
                        lhs122_.f4 = rhs159_
                    pass
            pass
        d_884_v85_: str
        d_884_v85_ = _dafny.CodePoint('s')
        d_885_v86_: _dafny.Map
        d_885_v86_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_760_v2_: d_760_v2_})) for d_886_i10_ in range(default__.abs(276))])) + (d_765_v7_): d_884_v85_})
        d_885_v86_ = (d_885_v86_).set(_dafny.SeqWithoutIsStrInference([(d_766_v8_).f7 for d_887_i11_ in range(default__.abs(-981))]), d_884_v85_)
        d_888_v87_: _dafny.Map
        d_888_v87_ = _dafny.Map({d_758_v0_: (d_766_v8_).f8})
        d_888_v87_ = (d_888_v87_).set((self).f8, (-460) * ((d_766_v8_).f7))
        d_889_v88_: _dafny.Map
        d_889_v88_ = _dafny.Map({d_760_v2_: d_888_v87_})
        d_890_v89_: _dafny.Seq
        d_890_v89_ = _dafny.SeqWithoutIsStrInference([False, d_760_v2_, d_760_v2_])
        d_889_v88_ = (d_889_v88_).set((d_890_v89_)[default__.safeIndex((self).f7, len(d_890_v89_))], (d_888_v87_) | (d_888_v87_))

    def m2(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        r2: bool = False
        if p3:
            (globalState).f4 = True
            d_891_v0_: _dafny.Seq
            d_891_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjrtfrpd"))
            d_891_v0_ = (d_891_v0_) + (d_891_v0_)
            (globalState).f4 = ((self).f7) != (p0)
            d_892_v1_: str
            d_892_v1_ = _dafny.CodePoint('q')
            d_893_v2_: D5
            d_893_v2_ = D5_DC14(D5_DC13(p3, d_892_v1_))
            d_893_v2_ = (d_893_v2_ if p3 else d_893_v2_)
            d_894_v3_: _dafny.Array
            def lambda39_(d_895_i0_):
                return default__.safeModuloInt(d_895_i0_, (self).f7)

            init22_ = lambda39_
            nw148_ = _dafny.Array(None, 16)
            for i0_22_ in range(nw148_.length(0)):
                nw148_[i0_22_] = init22_(i0_22_)
            d_894_v3_ = nw148_
            index162_ = default__.safeIndex(840, (d_894_v3_).length(0))
            (d_894_v3_)[index162_] = p2
            index163_ = default__.safeIndex(840, (d_894_v3_).length(0))
            (d_894_v3_)[index163_] = (self).f7
        elif True:
            d_896_v4_: _dafny.MultiSet
            d_896_v4_ = _dafny.MultiSet([p3])
            d_897_v5_: _dafny.Map
            d_897_v5_ = _dafny.Map({d_896_v4_: p2})
            d_898_v6_: _dafny.MultiSet
            d_898_v6_ = _dafny.MultiSet([-253])
            d_899_v7_: _dafny.Seq
            d_899_v7_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('h')])
            d_900_v8_: _dafny.Seq
            d_900_v8_ = _dafny.SeqWithoutIsStrInference([default__.fm24(p3, (0) - ((d_896_v4_).cardinality), d_899_v7_, globalState)])
            d_901_v9_: D6
            d_901_v9_ = D6_DC16((d_897_v5_).set(d_896_v4_, (d_898_v6_).cardinality), d_900_v8_)
            source20_ = d_901_v9_
            if source20_.is_DC16:
                d_902___mcc_h0_ = source20_.cf24
                d_903___mcc_h1_ = source20_.cf25
                d_904_cf25_ = d_903___mcc_h1_
                d_905_cf24_ = d_902___mcc_h0_
                d_906_v10_: str
                d_906_v10_ = _dafny.CodePoint('o')
                d_907_v11_: D7
                d_907_v11_ = D7_DC20(d_906_v10_, (self).f7)
                r1 = (d_907_v11_).cf33
                d_908_v12_: T1
                nw149_ = C4()
                nw149_.ctor__(p1, p0, (self).f7)
                d_908_v12_ = nw149_
                d_909_v13_: _dafny.MultiSet
                d_909_v13_ = _dafny.MultiSet([d_908_v12_])
                d_910_v14_: _dafny.Map
                d_910_v14_ = _dafny.Map({False: p1})
                d_911_v15_: _dafny.Map
                d_911_v15_ = _dafny.Map({(d_908_v12_).f8: p0})
                d_912_v16_: _dafny.Map
                d_912_v16_ = _dafny.Map({d_899_v7_: ((d_896_v4_)[p1] if (p1) in (d_896_v4_) else p2)})
                d_913_v17_: _dafny.Seq
                d_913_v17_ = _dafny.SeqWithoutIsStrInference([len(d_904_cf25_), ((d_912_v16_)[d_899_v7_] if (d_899_v7_) in (d_912_v16_) else (d_908_v12_).f7), (d_908_v12_).f7, (d_908_v12_).f7])
                d_914_v18_: _dafny.Set
                d_914_v18_ = _dafny.Set({True, p3, p3})
                d_915_v19_: _dafny.Array
                nw150_ = _dafny.Array(None, 18)
                nw150_[int(0)] = 886
                nw150_[int(1)] = (self).f8
                nw150_[int(2)] = default__.safeDivisionInt(p2, (0) - (p2))
                nw150_[int(3)] = p0
                nw150_[int(4)] = (self).f7
                nw150_[int(5)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yiww")))
                nw150_[int(6)] = (d_909_v13_).cardinality
                nw150_[int(7)] = (self).f7
                nw150_[int(8)] = ((d_908_v12_).f7) - (636)
                nw150_[int(9)] = (self).f8
                nw150_[int(10)] = (d_908_v12_).f7
                nw150_[int(11)] = (self).f7
                nw150_[int(12)] = (d_908_v12_).f8
                nw150_[int(13)] = p0
                nw150_[int(14)] = (len(d_910_v14_)) + (p0)
                nw150_[int(15)] = (((d_896_v4_)[not(p1)] if (not(p1)) in (d_896_v4_) else ((d_911_v15_)[(d_908_v12_).f7] if ((d_908_v12_).f7) in (d_911_v15_) else (self).f8))) + ((d_913_v17_)[default__.safeIndex(len(d_914_v18_), len(d_913_v17_))])
                nw150_[int(16)] = ((self).f8 if p1 else -589)
                nw150_[int(17)] = default__.safeDivisionInt(len(d_904_cf25_), 755)
                d_915_v19_ = nw150_
                index164_ = default__.safeIndex(248, (d_915_v19_).length(0))
                (d_915_v19_)[index164_] = default__.safeDivisionInt(p0, p2)
                index165_ = default__.safeIndex(248, (d_915_v19_).length(0))
                (d_915_v19_)[index165_] = (self).f8
                d_916_v20_: int
                d_917_v21_: int
                d_918_v22_: bool
                out13_: int
                out14_: int
                out15_: bool
                out13_, out14_, out15_ = (self).m3(p3, len(_dafny.Map({689: len(d_899_v7_)})), d_915_v19_, (self).f8, globalState)
                d_916_v20_ = out13_
                d_917_v21_ = out14_
                d_918_v22_ = out15_
                d_919_v23_: int
                d_920_v24_: int
                d_921_v25_: bool
                out16_: int
                out17_: int
                out18_: bool
                out16_, out17_, out18_ = (self).m3(False, ((self).f8) - ((d_908_v12_).f7), d_915_v19_, (0) - ((self).f7), globalState)
                d_919_v23_ = out16_
                d_920_v24_ = out17_
                d_921_v25_ = out18_
            elif True:
                d_922___mcc_h2_ = source20_.cf23
                d_923_cf23_ = d_922___mcc_h2_
                r1 = (self).f7
                d_924_v26_: str
                d_924_v26_ = _dafny.CodePoint('e')
                d_925_v27_: D7
                d_925_v27_ = D7_DC20((d_924_v26_ if p3 else d_924_v26_), (self).f8)
                d_926_v28_: D5
                d_926_v28_ = D5_DC13(p1, d_924_v26_)
                d_927_v29_: _dafny.Array
                def lambda40_(d_928_p0_):
                    def lambda41_(d_929_i1_):
                        return (d_929_i1_) * (d_928_p0_)

                    return lambda41_

                init23_ = lambda40_(p0)
                nw151_ = _dafny.Array(None, 8)
                for i0_23_ in range(nw151_.length(0)):
                    nw151_[i0_23_] = init23_(i0_23_)
                d_927_v29_ = nw151_
                d_930_v30_: _dafny.Map
                d_930_v30_ = _dafny.Map({d_926_v28_: d_927_v29_})
                d_931_v31_: _dafny.Seq
                d_931_v31_ = _dafny.SeqWithoutIsStrInference([d_930_v30_])
                rhs160_ = ((d_931_v31_)[default__.safeIndex((self).f7, len(d_931_v31_))]) == (((d_930_v30_).set(d_926_v28_, d_927_v29_)).set(d_926_v28_, d_927_v29_))
                rhs161_ = (p1 if p3 else p1)
                rhs162_ = D7_DC20(d_924_v26_, (0) - (p0))
                lhs123_ = globalState
                lhs124_ = globalState
                lhs123_.f3 = rhs160_
                lhs124_.f3 = rhs161_
                d_925_v27_ = rhs162_
                index166_ = default__.safeIndex(844, (d_923_cf23_).length(0))
                (d_923_cf23_)[index166_] = p3
                index167_ = default__.safeIndex(844, (d_923_cf23_).length(0))
                (d_923_cf23_)[index167_] = not(p3)
                r1 = (len((_dafny.SeqWithoutIsStrInference([d_924_v26_ for d_932_i2_ in range(default__.abs(873))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_933_i3_ in range(default__.abs(19))])))) * (-398)
            d_934_v32_: _dafny.Map
            d_934_v32_ = _dafny.Map({p3: 662})
            d_935_v33_: D5
            d_935_v33_ = D5_DC11(d_934_v32_)
            source21_ = d_935_v33_
            if source21_.is_DC12:
                d_936___mcc_h3_ = source21_.cf19
                d_937_cf19_ = d_936___mcc_h3_
                d_938_v34_: _dafny.Seq
                d_938_v34_ = _dafny.SeqWithoutIsStrInference([d_901_v9_])
                d_939_v35_: D3
                d_939_v35_ = D3_DC6(_dafny.Set({-311, p0, (self).f7, (self).f8, (self).f8}))
                d_940_v36_: _dafny.Set
                d_940_v36_ = _dafny.Set({p2, (self).f8})
                pat_let_tv14_ = d_940_v36_
                d_941_v37_: _dafny.Map
                def iife62_(_pat_let12_0):
                    def iife63_(d_942_dt__update__tmp_h0_):
                        def iife64_(_pat_let13_0):
                            def iife65_(d_943_dt__update_hcf9_h0_):
                                return D3_DC6(d_943_dt__update_hcf9_h0_)
                            return iife65_(_pat_let13_0)
                        return iife64_(pat_let_tv14_)
                    return iife63_(_pat_let12_0)
                d_941_v37_ = _dafny.Map({default__.safeModuloInt(len(d_938_v34_), (self).f8): iife62_(d_939_v35_)})
                d_941_v37_ = (d_941_v37_).set(((0) - (p0)) * (p0), d_939_v35_)
                d_899_v7_ = d_899_v7_
                (globalState).f4 = p1
                r1 = -788
            elif source21_.is_DC13:
                d_944___mcc_h4_ = source21_.cf20
                d_945___mcc_h5_ = source21_.cf21
                d_946_cf21_ = d_945___mcc_h5_
                d_947_cf20_ = d_944___mcc_h4_
                d_948_v38_: _dafny.Array
                nw152_ = _dafny.Array(None, 8)
                nw152_[int(0)] = ((self).f7) < ((self).f7)
                nw152_[int(1)] = (p3 if p3 else d_947_cf20_)
                nw152_[int(2)] = p1
                nw152_[int(3)] = True
                nw152_[int(4)] = d_947_cf20_
                nw152_[int(5)] = d_947_cf20_
                nw152_[int(6)] = d_947_cf20_
                nw152_[int(7)] = d_947_cf20_
                d_948_v38_ = nw152_
                rhs163_ = p0
                rhs164_ = d_948_v38_
                rhs165_ = (p2) - (default__.safeDivisionInt(p2, (self).fm3(d_947_cf20_, p1, p3, True, globalState)))
                r1 = rhs163_
                d_948_v38_ = rhs164_
                r1 = rhs165_
                d_949_v39_: _dafny.Seq
                d_949_v39_ = _dafny.SeqWithoutIsStrInference([d_948_v38_])
                d_950_v40_: _dafny.Array
                nw153_ = _dafny.Array(_dafny.Seq({}), 5)
                d_950_v40_ = nw153_
                d_951_v41_: _dafny.Seq
                d_951_v41_ = _dafny.SeqWithoutIsStrInference([(self).f8])
                index168_ = default__.safeIndex(310, (d_950_v40_).length(0))
                (d_950_v40_)[index168_] = d_951_v41_
                index169_ = default__.safeIndex(310, (d_950_v40_).length(0))
                rhs166_ = (d_949_v39_ if p1 else d_949_v39_)
                rhs167_ = ((d_951_v41_) + (_dafny.SeqWithoutIsStrInference([(self).f7]))) + ((d_951_v41_) + (d_951_v41_))
                rhs168_ = d_899_v7_
                rhs169_ = d_946_cf21_
                lhs125_ = d_950_v40_
                lhs126_ = default__.safeIndex(310, (d_950_v40_).length(0))
                d_949_v39_ = rhs166_
                lhs125_[lhs126_] = rhs167_
                d_899_v7_ = rhs168_
                d_946_cf21_ = rhs169_
                d_952_v42_: C1
                nw154_ = C1()
                nw154_.ctor__((self).f7, (0) - (p0))
                d_952_v42_ = nw154_
                d_953_v43_: _dafny.Map
                d_953_v43_ = _dafny.Map({p2: not(p3)})
                d_954_v44_: _dafny.Set
                d_954_v44_ = _dafny.Set({(self).f7, p2, p0, (0) - (len(d_953_v43_))})
                d_955_v45_: _dafny.Map
                d_955_v45_ = _dafny.Map({d_952_v42_: len(d_954_v44_)})
                d_955_v45_ = (d_955_v45_).set(d_952_v42_, p0)
                d_956_v46_: _dafny.Array
                def lambda42_(d_957_i4_):
                    return default__.safeModuloInt(d_957_i4_, (self).f8)

                init24_ = lambda42_
                nw155_ = _dafny.Array(None, 23)
                for i0_24_ in range(nw155_.length(0)):
                    nw155_[i0_24_] = init24_(i0_24_)
                d_956_v46_ = nw155_
                index170_ = default__.safeIndex(684, (d_956_v46_).length(0))
                (d_956_v46_)[index170_] = default__.safeDivisionInt(p2, (self).f7)
                d_958_v47_: _dafny.Map
                d_958_v47_ = _dafny.Map({(self).f7: len(d_899_v7_)})
                d_959_v48_: _dafny.Map
                d_959_v48_ = _dafny.Map({p1: d_958_v47_})
                index171_ = default__.safeIndex(684, (d_956_v46_).length(0))
                (d_956_v46_)[index171_] = (default__.safeModuloInt(p0, len(((d_959_v48_)[p3] if (p3) in (d_959_v48_) else d_958_v47_)))) + ((self).f8)
            elif source21_.is_DC11:
                d_960___mcc_h6_ = source21_.cf18
                d_961_cf18_ = d_960___mcc_h6_
                (globalState).f3 = p3
                d_962_v49_: str
                d_962_v49_ = _dafny.CodePoint('j')
                d_963_v50_: _dafny.MultiSet
                d_963_v50_ = _dafny.MultiSet([d_962_v49_])
                d_964_v51_: _dafny.Seq
                d_964_v51_ = _dafny.SeqWithoutIsStrInference([(d_963_v50_).cardinality, p0, p2])
                d_965_v52_: C5
                nw156_ = C5()
                nw156_.ctor__(d_899_v7_, (self).f7, (d_964_v51_)[default__.safeIndex(p0, len(d_964_v51_))], (self).f8)
                d_965_v52_ = nw156_
                r1 = (self).f7
                r1 = p2
            elif True:
                d_966___mcc_h7_ = source21_.cf22
                d_967_cf22_ = d_966___mcc_h7_
                d_934_v32_ = (d_934_v32_).set(True, (0) - ((self).f7))
                r1 = ((self).f8) + (((self).f8 if p3 else (self).f8))
                d_968_v53_: _dafny.Array
                def lambda43_(d_969_i5_):
                    return default__.safeModuloInt(d_969_i5_, 648)

                init25_ = lambda43_
                nw157_ = _dafny.Array(None, 7)
                for i0_25_ in range(nw157_.length(0)):
                    nw157_[i0_25_] = init25_(i0_25_)
                d_968_v53_ = nw157_
                index172_ = default__.safeIndex(659, (d_968_v53_).length(0))
                (d_968_v53_)[index172_] = len(d_899_v7_)
                index173_ = default__.safeIndex(659, (d_968_v53_).length(0))
                (d_968_v53_)[index173_] = p0
                index174_ = default__.safeIndex(659, (d_968_v53_).length(0))
                (d_968_v53_)[index174_] = 758
            d_970_v54_: D1
            d_970_v54_ = D1_DC3(default__.fm23((self).f8, globalState))
            source22_ = d_970_v54_
            if source22_.is_DC2:
                d_971_v55_: _dafny.Array
                def lambda44_(d_972_p1_):
                    def lambda45_(d_973_i6_):
                        return d_972_p1_

                    return lambda45_

                init26_ = lambda44_(p1)
                nw158_ = _dafny.Array(None, 15)
                for i0_26_ in range(nw158_.length(0)):
                    nw158_[i0_26_] = init26_(i0_26_)
                d_971_v55_ = nw158_
                index175_ = default__.safeIndex(715, (d_971_v55_).length(0))
                (d_971_v55_)[index175_] = p1
                d_974_v56_: _dafny.Array
                nw159_ = _dafny.Array(int(0), 24)
                d_974_v56_ = nw159_
                index176_ = default__.safeIndex(870, (d_974_v56_).length(0))
                (d_974_v56_)[index176_] = p2
                d_975_v57_: C1
                nw160_ = C1()
                nw160_.ctor__((self).f7, (self).f8)
                d_975_v57_ = nw160_
                d_976_v58_: _dafny.Seq
                d_976_v58_ = _dafny.SeqWithoutIsStrInference([d_975_v57_])
                index177_ = default__.safeIndex(715, (d_971_v55_).length(0))
                index178_ = default__.safeIndex(870, (d_974_v56_).length(0))
                rhs170_ = (p1) and ((d_976_v58_) != (d_976_v58_))
                rhs171_ = (p0) < (len(d_900_v8_))
                rhs172_ = default__.fm16(not(p1), globalState)
                rhs173_ = (p2) + (738)
                lhs127_ = d_971_v55_
                lhs128_ = default__.safeIndex(715, (d_971_v55_).length(0))
                lhs129_ = globalState
                lhs130_ = d_974_v56_
                lhs131_ = default__.safeIndex(870, (d_974_v56_).length(0))
                lhs127_[lhs128_] = rhs170_
                lhs129_.f4 = rhs171_
                r1 = rhs172_
                lhs130_[lhs131_] = rhs173_
                d_977_v59_: _dafny.Map
                d_977_v59_ = _dafny.Map({(((d_896_v4_)[(d_971_v55_)[default__.safeIndex(715, (d_971_v55_).length(0))]] if ((d_971_v55_)[default__.safeIndex(715, (d_971_v55_).length(0))]) in (d_896_v4_) else (self).f8)) - ((self).f7): p0})
                d_977_v59_ = (d_977_v59_).set(default__.safeDivisionInt(594, 147), p2)
                index179_ = default__.safeIndex(870, (d_974_v56_).length(0))
                (d_974_v56_)[index179_] = len((d_899_v7_) + ((d_899_v7_ if (d_971_v55_)[default__.safeIndex(715, (d_971_v55_).length(0))] else d_899_v7_)))
                index180_ = default__.safeIndex(870, (d_974_v56_).length(0))
                (d_974_v56_)[index180_] = (p2) * ((d_974_v56_)[default__.safeIndex(870, (d_974_v56_).length(0))])
            elif source22_.is_DC1:
                d_978___mcc_h8_ = source22_.cf1
                d_979_cf1_ = d_978___mcc_h8_
                d_980_v60_: _dafny.Map
                d_980_v60_ = _dafny.Map({(self).f8: p1})
                d_981_v61_: _dafny.Map
                d_981_v61_ = _dafny.Map({(self).f8: len(_dafny.Map({p2: p3}))})
                d_982_v62_: _dafny.Seq
                d_982_v62_ = _dafny.SeqWithoutIsStrInference([d_981_v61_, d_981_v61_])
                d_983_v63_: _dafny.Map
                d_983_v63_ = _dafny.Map({d_899_v7_: p3})
                d_984_v64_: _dafny.Array
                nw161_ = _dafny.Array(None, 27)
                nw161_[int(0)] = (self).f8
                nw161_[int(1)] = len(d_980_v60_)
                nw161_[int(2)] = default__.fm16(p1, globalState)
                nw161_[int(3)] = default__.fm20(p3, p1, p1, globalState)
                nw161_[int(4)] = (self).f7
                nw161_[int(5)] = (0) - ((self).f7)
                nw161_[int(6)] = (self).f7
                nw161_[int(7)] = p0
                nw161_[int(8)] = len(d_899_v7_)
                nw161_[int(9)] = len(d_982_v62_)
                nw161_[int(10)] = len(d_900_v8_)
                nw161_[int(11)] = (0) - ((self).f8)
                nw161_[int(12)] = 500
                nw161_[int(13)] = p0
                nw161_[int(14)] = p2
                nw161_[int(15)] = 532
                nw161_[int(16)] = p2
                nw161_[int(17)] = p2
                nw161_[int(18)] = len(d_899_v7_)
                nw161_[int(19)] = p0
                nw161_[int(20)] = (self).f7
                nw161_[int(21)] = (self).f8
                nw161_[int(22)] = (self).f8
                nw161_[int(23)] = len(d_983_v63_)
                nw161_[int(24)] = 274
                nw161_[int(25)] = p2
                nw161_[int(26)] = p0
                d_984_v64_ = nw161_
                d_985_v65_: _dafny.Seq
                d_985_v65_ = _dafny.SeqWithoutIsStrInference([d_984_v64_, d_984_v64_, d_984_v64_, d_984_v64_, d_984_v64_])
                d_986_v66_: _dafny.Array
                nw162_ = _dafny.Array(None, 10)
                nw162_[int(0)] = d_984_v64_
                nw162_[int(1)] = d_984_v64_
                nw162_[int(2)] = d_984_v64_
                nw162_[int(3)] = d_984_v64_
                nw162_[int(4)] = d_984_v64_
                nw162_[int(5)] = d_984_v64_
                nw162_[int(6)] = (d_985_v65_)[default__.safeIndex((0) - (p2), len(d_985_v65_))]
                nw162_[int(7)] = d_984_v64_
                nw162_[int(8)] = d_984_v64_
                nw162_[int(9)] = d_984_v64_
                d_986_v66_ = nw162_
                index181_ = default__.safeIndex(861, (d_986_v66_).length(0))
                (d_986_v66_)[index181_] = d_984_v64_
                index182_ = default__.safeIndex(861, (d_986_v66_).length(0))
                (d_986_v66_)[index182_] = d_984_v64_
                d_987_v67_: _dafny.Set
                d_987_v67_ = _dafny.Set({d_899_v7_})
                (globalState).f4 = ((d_987_v67_) - (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwlofrqbi"))}))).issubset((d_987_v67_ if (d_900_v8_)[default__.safeIndex(p0, len(d_900_v8_))] else d_987_v67_))
                d_988_v68_: _dafny.Map
                d_988_v68_ = _dafny.Map({p1: d_899_v7_})
                index183_ = default__.safeIndex(155, (d_984_v64_).length(0))
                (d_984_v64_)[index183_] = len(d_988_v68_)
                d_989_v69_: _dafny.Seq
                d_989_v69_ = _dafny.SeqWithoutIsStrInference([p0])
                d_990_v70_: _dafny.Map
                d_990_v70_ = _dafny.Map({d_979_cf1_: d_899_v7_})
                d_991_v71_: _dafny.Map
                d_991_v71_ = _dafny.Map({p3: (d_986_v66_ if p3 else d_986_v66_)})
                index184_ = default__.safeIndex(155, (d_984_v64_).length(0))
                rhs174_ = ((0) - ((self).f8)) + ((self).f7)
                rhs175_ = d_989_v69_
                rhs176_ = (p1 if (d_979_cf1_) not in (d_899_v7_) else ((d_980_v60_)[(self).f8] if ((self).f8) in (d_980_v60_) else not(False)))
                rhs177_ = default__.safeDivisionInt((self).f7, len(d_990_v70_))
                rhs178_ = ((d_991_v71_)[True] if (True) in (d_991_v71_) else d_986_v66_)
                lhs132_ = d_984_v64_
                lhs133_ = default__.safeIndex(155, (d_984_v64_).length(0))
                lhs134_ = globalState
                lhs132_[lhs133_] = rhs174_
                d_989_v69_ = rhs175_
                lhs134_.f3 = rhs176_
                r1 = rhs177_
                d_986_v66_ = rhs178_
                index185_ = default__.safeIndex(155, (d_984_v64_).length(0))
                (d_984_v64_)[index185_] = ((0) - (default__.safeDivisionInt((0) - ((self).f7), (self).f8))) * (646)
            elif True:
                d_992___mcc_h9_ = source22_.cf2
                d_993_cf2_ = d_992___mcc_h9_
                d_994_v72_: _dafny.Set
                d_994_v72_ = _dafny.Set({p2})
                d_994_v72_ = d_994_v72_
                d_995_v73_: T0
                nw163_ = C5()
                nw163_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "svx")), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_996_i7_ in range(default__.abs(532))])), 135, (self).f8)
                d_995_v73_ = nw163_
                d_997_v74_: _dafny.Map
                d_997_v74_ = _dafny.Map({d_899_v7_: d_995_v73_})
                d_998_v75_: _dafny.Seq
                d_998_v75_ = _dafny.SeqWithoutIsStrInference([((d_997_v74_)[(self).fm2(p3, globalState)] if ((self).fm2(p3, globalState)) in (d_997_v74_) else d_995_v73_), d_995_v73_])
                d_999_v76_: C2
                nw164_ = C2()
                nw164_.ctor__(len(d_998_v75_), default__.fm20(p3, False, p1, globalState))
                d_999_v76_ = nw164_
                d_900_v8_ = d_900_v8_
                d_1000_v77_: _dafny.Array
                nw165_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_1000_v77_ = nw165_
                d_1001_v78_: _dafny.Array
                nw166_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_1001_v78_ = nw166_
                index186_ = default__.safeIndex(418, (d_1000_v77_).length(0))
                (d_1000_v77_)[index186_] = d_1001_v78_
                d_1002_v79_: _dafny.Array
                nw167_ = _dafny.Array(_dafny.Seq({}), 16)
                d_1002_v79_ = nw167_
                d_1003_v80_: D7
                d_1003_v80_ = D7_DC19(p0, p1, (self).f7)
                d_1004_v81_: _dafny.Map
                d_1004_v81_ = _dafny.Map({((d_995_v73_).f8) - (p2): ((d_1003_v80_).cf31) <= (p2)})
                index187_ = default__.safeIndex(418, (d_1000_v77_).length(0))
                rhs179_ = d_1001_v78_
                rhs180_ = d_899_v7_
                rhs181_ = d_1002_v79_
                rhs182_ = p3
                rhs183_ = d_1004_v81_
                lhs135_ = d_1000_v77_
                lhs136_ = default__.safeIndex(418, (d_1000_v77_).length(0))
                lhs137_ = globalState
                lhs135_[lhs136_] = rhs179_
                d_899_v7_ = rhs180_
                d_1002_v79_ = rhs181_
                lhs137_.f3 = rhs182_
                d_1004_v81_ = rhs183_
            r1 = p2
            d_1005_v82_: _dafny.Array
            nw168_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_1005_v82_ = nw168_
            d_1006_v83_: _dafny.Array
            nw169_ = _dafny.Array(False, 19)
            d_1006_v83_ = nw169_
            d_1007_v84_: _dafny.Map
            d_1007_v84_ = _dafny.Map({p3: d_1006_v83_})
            d_1008_v85_: str
            d_1008_v85_ = _dafny.CodePoint('w')
            d_1009_v86_: _dafny.Array
            nw170_ = _dafny.Array(None, 14)
            nw170_[int(0)] = ((d_898_v6_)[p0] if (p0) in (d_898_v6_) else (self).f8)
            nw170_[int(1)] = (self).f7
            nw170_[int(2)] = p0
            nw170_[int(3)] = len(d_1007_v84_)
            nw170_[int(4)] = (self).f7
            nw170_[int(5)] = (0) - (p0)
            nw170_[int(6)] = p2
            nw170_[int(7)] = (self).fm4(True, True, globalState)
            nw170_[int(8)] = p2
            nw170_[int(9)] = default__.fm20(p1, p1, p3, globalState)
            nw170_[int(10)] = (self).f7
            nw170_[int(11)] = 965
            nw170_[int(12)] = ((_dafny.MultiSet([d_1008_v85_])).set(d_1008_v85_, default__.abs(len(d_899_v7_)))).cardinality
            nw170_[int(13)] = p0
            d_1009_v86_ = nw170_
            index188_ = default__.safeIndex(933, (d_1005_v82_).length(0))
            (d_1005_v82_)[index188_] = d_1009_v86_
            d_1010_v87_: D2
            d_1010_v87_ = D2_DC4(d_896_v4_)
            d_1011_v88_: _dafny.MultiSet
            d_1011_v88_ = _dafny.MultiSet([d_1010_v87_, d_1010_v87_, d_1010_v87_])
            d_1012_v89_: _dafny.Map
            d_1012_v89_ = _dafny.Map({(self).f7: p3})
            index189_ = default__.safeIndex(933, (d_1005_v82_).length(0))
            rhs184_ = (-197 if p1 else default__.safeDivisionInt((self).f8, len(d_1012_v89_)))
            rhs185_ = d_1009_v86_
            rhs186_ = (_dafny.MultiSet([d_1010_v87_])) | (d_1011_v88_)
            rhs187_ = ((self).f8) + ((self).f7)
            lhs138_ = d_1005_v82_
            lhs139_ = default__.safeIndex(933, (d_1005_v82_).length(0))
            r1 = rhs184_
            lhs138_[lhs139_] = rhs185_
            d_1011_v88_ = rhs186_
            r1 = rhs187_
        if p3:
            d_1013_v90_: str
            d_1013_v90_ = _dafny.CodePoint('h')
            d_1013_v90_ = d_1013_v90_
            r1 = ((self).f7 if not(True) else (p0 if p1 else (self).f7))
            d_1014_v91_: _dafny.Array
            def lambda46_(d_1015_i8_):
                return (d_1015_i8_) * (-317)

            init27_ = lambda46_
            nw171_ = _dafny.Array(None, 8)
            for i0_27_ in range(nw171_.length(0)):
                nw171_[i0_27_] = init27_(i0_27_)
            d_1014_v91_ = nw171_
            index190_ = default__.safeIndex(552, (d_1014_v91_).length(0))
            (d_1014_v91_)[index190_] = 629
            d_1016_v92_: _dafny.MultiSet
            d_1016_v92_ = _dafny.MultiSet([p2, p2])
            d_1017_v94_: _dafny.Set
            def iife66_():
                coll38_ = _dafny.Set()
                compr_38_: int
                for compr_38_ in _dafny.IntegerRange(224, 478):
                    d_1018_v93_: int = compr_38_
                    if ((224) <= (d_1018_v93_)) and ((d_1018_v93_) < (478)):
                        coll38_ = coll38_.union(_dafny.Set([(d_1018_v93_) + (701)]))
                return _dafny.Set(coll38_)
            d_1017_v94_ = _dafny.Set({default__.fm29(p3, d_1016_v92_, globalState), (iife66_()
            ).intersection(_dafny.Set({-460, (self).f8, (0) - (p2)}))})
            d_1019_v95_: _dafny.Seq
            d_1019_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctt"))
            d_1020_v96_: _dafny.Set
            d_1020_v96_ = _dafny.Set({788})
            index191_ = default__.safeIndex(552, (d_1014_v91_).length(0))
            rhs188_ = (0) - (default__.safeModuloInt((self).f7, len(d_1019_v95_)))
            rhs189_ = ((d_1017_v94_) - (_dafny.Set({d_1020_v96_}))) - (d_1017_v94_)
            rhs190_ = p0
            lhs140_ = d_1014_v91_
            lhs141_ = default__.safeIndex(552, (d_1014_v91_).length(0))
            lhs140_[lhs141_] = rhs188_
            d_1017_v94_ = rhs189_
            r1 = rhs190_
            r1 = ((self).fm3(p3, p1, p3, False, globalState)) * ((-226) * ((d_1014_v91_)[default__.safeIndex(552, (d_1014_v91_).length(0))]))
            d_1021_v97_: _dafny.Seq
            d_1021_v97_ = _dafny.SeqWithoutIsStrInference([p1, (D5_DC13(default__.fm24(p1, (self).f7, d_1019_v95_, globalState), d_1013_v90_)).cf20, not(p3)])
            d_1022_v98_: T2
            nw172_ = C2()
            nw172_.ctor__((self).f8, p2)
            d_1022_v98_ = nw172_
            d_1023_v99_: _dafny.Map
            d_1023_v99_ = _dafny.Map({len(d_1020_v96_): d_1022_v98_})
            d_1024_v100_: _dafny.Array
            nw173_ = _dafny.Array(None, 4)
            nw173_[int(0)] = p1
            nw173_[int(1)] = (d_1021_v97_)[default__.safeIndex(len(d_1023_v99_), len(d_1021_v97_))]
            nw173_[int(2)] = p3
            nw173_[int(3)] = p3
            d_1024_v100_ = nw173_
            index192_ = default__.safeIndex(225, (d_1024_v100_).length(0))
            (d_1024_v100_)[index192_] = p1
            index193_ = default__.safeIndex(225, (d_1024_v100_).length(0))
            (d_1024_v100_)[index193_] = (d_1021_v97_)[default__.safeIndex(731, len(d_1021_v97_))]
        elif True:
            d_1025_v101_: _dafny.Map
            d_1025_v101_ = _dafny.Map({False: p1})
            r1 = (len((d_1025_v101_).set(p3, p1))) * ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymjehvl")))) + ((self).f7))
            d_1026_v102_: _dafny.Seq
            d_1026_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uaehms"))
            d_1027_v103_: _dafny.Seq
            d_1027_v103_ = _dafny.SeqWithoutIsStrInference([d_1026_v102_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1028_i9_ in range(default__.abs(274))]), d_1026_v102_])
            if ((d_1026_v102_) + (d_1026_v102_)) != (((d_1027_v103_)[default__.safeIndex((self).f7, len(d_1027_v103_))]) + (d_1026_v102_)):
                r1 = (self).f8
                d_1029_v104_: _dafny.Seq
                d_1029_v104_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1030_v105_: str
                d_1030_v105_ = _dafny.CodePoint('i')
                d_1031_v106_: _dafny.Seq
                d_1031_v106_ = _dafny.SeqWithoutIsStrInference([(self).f8])
                d_1032_v107_: _dafny.Array
                nw174_ = _dafny.Array(None, 10)
                nw174_[int(0)] = len((d_1029_v104_).set(default__.safeIndex(p2, len(d_1029_v104_)), p1))
                nw174_[int(1)] = (D7_DC20(d_1030_v105_, len(d_1027_v103_))).cf33
                nw174_[int(2)] = (self).f7
                nw174_[int(3)] = (p2) * (len(_dafny.Map({not(p3): p1})))
                nw174_[int(4)] = (p2 if p3 else (self).f7)
                nw174_[int(5)] = (len(d_1031_v106_)) + ((0) - ((self).f8))
                nw174_[int(6)] = p2
                nw174_[int(7)] = (545) - ((0) - (-723))
                nw174_[int(8)] = p2
                nw174_[int(9)] = (self).f7
                d_1032_v107_ = nw174_
                d_1033_v108_: _dafny.Map
                d_1033_v108_ = _dafny.Map({d_1030_v105_: p3})
                d_1034_v109_: C5
                nw175_ = C5()
                nw175_.ctor__(d_1026_v102_, (self).f7, (self).f7, (0) - (len(d_1033_v108_)))
                d_1034_v109_ = nw175_
                d_1035_v110_: _dafny.Map
                d_1035_v110_ = _dafny.Map({(self).f7: d_1034_v109_})
                index194_ = default__.safeIndex(17, (d_1032_v107_).length(0))
                (d_1032_v107_)[index194_] = ((self).f8) - (len((d_1035_v110_).set((self).f8, d_1034_v109_)))
                index195_ = default__.safeIndex(17, (d_1032_v107_).length(0))
                rhs191_ = (self).f7
                rhs192_ = (self).f8
                lhs142_ = d_1032_v107_
                lhs143_ = default__.safeIndex(17, (d_1032_v107_).length(0))
                lhs142_[lhs143_] = rhs191_
                r1 = rhs192_
                d_1030_v105_ = d_1030_v105_
                d_1036_v111_: _dafny.Map
                d_1036_v111_ = _dafny.Map({_dafny.CodePoint('k'): (self).f8})
                d_1037_v113_: _dafny.Set
                d_1037_v113_ = _dafny.Set({default__.fm5(globalState)})
                d_1038_v114_: _dafny.Seq
                def iife67_():
                    coll39_ = _dafny.Set()
                    compr_39_: str
                    for compr_39_ in (d_1036_v111_).keys.Elements:
                        d_1039_v112_: str = compr_39_
                        if (d_1039_v112_) in (d_1036_v111_):
                            coll39_ = coll39_.union(_dafny.Set([d_1039_v112_]))
                    return _dafny.Set(coll39_)
                d_1038_v114_ = _dafny.SeqWithoutIsStrInference([(_dafny.Set({d_1030_v105_})) | (iife67_()
                ), d_1037_v113_])
                d_1038_v114_ = (d_1038_v114_) + ((d_1038_v114_) + (d_1038_v114_))
                index196_ = default__.safeIndex(17, (d_1032_v107_).length(0))
                (d_1032_v107_)[index196_] = ((self).f7) + ((self).f7)
            elif True:
                d_1040_v115_: _dafny.Seq
                d_1040_v115_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1041_v116_: D2
                d_1041_v116_ = D2_DC5(p1, (self).f8, len((d_1040_v115_).set(default__.safeIndex(len(d_1040_v115_), len(d_1040_v115_)), p1)), (self).f7, default__.fm24(p3, 998, d_1026_v102_, globalState))
                d_1042_v117_: _dafny.Set
                d_1042_v117_ = _dafny.Set({p3})
                d_1043_v118_: _dafny.Array
                nw176_ = _dafny.Array(None, 17)
                nw176_[int(0)] = not(p3)
                nw176_[int(1)] = p1
                nw176_[int(2)] = not(not((p3 if (d_1041_v116_).cf4 else (self).fm1(d_1042_v117_, (self).f7, -719, (0) - ((0) - (p0)), globalState))))
                nw176_[int(3)] = p3
                nw176_[int(4)] = p3
                nw176_[int(5)] = p3
                nw176_[int(6)] = (d_1026_v102_) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1044_i10_ in range(default__.abs(279))]))
                nw176_[int(7)] = p1
                nw176_[int(8)] = p1
                nw176_[int(9)] = not (p3) or (True)
                nw176_[int(10)] = p3
                nw176_[int(11)] = p3
                nw176_[int(12)] = p1
                nw176_[int(13)] = not (p1) or (True)
                nw176_[int(14)] = p3
                nw176_[int(15)] = p3
                nw176_[int(16)] = p3
                d_1043_v118_ = nw176_
                index197_ = default__.safeIndex(271, (d_1043_v118_).length(0))
                (d_1043_v118_)[index197_] = True
                index198_ = default__.safeIndex(271, (d_1043_v118_).length(0))
                (d_1043_v118_)[index198_] = ((self).f8) != (((self).f8) + (p2))
                d_1045_v119_: _dafny.Set
                d_1045_v119_ = _dafny.Set({(0) - ((0) - (len(d_1025_v101_))), p2})
                r1 = (p0) * (len(d_1045_v119_))
                d_1046_v120_: _dafny.Map
                d_1046_v120_ = _dafny.Map({p0: (self).f7})
                d_1047_v121_: C1
                nw177_ = C1()
                nw177_.ctor__(((d_1046_v120_)[p0] if (p0) in (d_1046_v120_) else (self).f7), ((self).f7 if not(p3) else (self).fm3(p1, True, (d_1043_v118_)[default__.safeIndex(271, (d_1043_v118_).length(0))], p1, globalState)))
                d_1047_v121_ = nw177_
                d_1048_v122_: D5
                d_1048_v122_ = D5_DC13(p1, _dafny.CodePoint('q'))
                d_1049_v123_: _dafny.Map
                d_1049_v123_ = _dafny.Map({default__.safeModuloInt(default__.fm16(p3, globalState), p0): _dafny.SeqWithoutIsStrInference([d_1048_v122_, d_1048_v122_, default__.fm32(globalState), default__.fm32(globalState)])})
                d_1050_v124_: _dafny.MultiSet
                d_1050_v124_ = _dafny.MultiSet([p3])
                d_1051_v125_: _dafny.MultiSet
                d_1051_v125_ = _dafny.MultiSet([((d_1050_v124_)[p3] if (p3) in (d_1050_v124_) else (self).f7)])
                d_1052_v126_: _dafny.Seq
                d_1052_v126_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt((d_1051_v125_).cardinality, 827), (0) - ((-451) + (p0)), p0, default__.safeModuloInt((self).f7, p2)])
                index199_ = default__.safeIndex(271, (d_1043_v118_).length(0))
                rhs193_ = (False if p1 else p1)
                rhs194_ = d_1026_v102_
                rhs195_ = d_1049_v123_
                rhs196_ = (d_1043_v118_)[default__.safeIndex(271, (d_1043_v118_).length(0))]
                rhs197_ = d_1052_v126_
                lhs144_ = globalState
                lhs145_ = d_1043_v118_
                lhs146_ = default__.safeIndex(271, (d_1043_v118_).length(0))
                lhs144_.f4 = rhs193_
                d_1026_v102_ = rhs194_
                d_1049_v123_ = rhs195_
                lhs145_[lhs146_] = rhs196_
                d_1052_v126_ = rhs197_
                r1 = (len((d_1026_v102_) + (d_1026_v102_))) + (((d_1051_v125_)[default__.fm20(p3, (d_1043_v118_)[default__.safeIndex(271, (d_1043_v118_).length(0))], p1, globalState)] if (default__.fm20(p3, (d_1043_v118_)[default__.safeIndex(271, (d_1043_v118_).length(0))], p1, globalState)) in (d_1051_v125_) else len(d_1040_v115_)))
            d_1053_v127_: C3
            nw178_ = C3()
            nw178_.ctor__(not(p3), (self).f7, p2)
            d_1053_v127_ = nw178_
            d_1054_v128_: _dafny.Array
            nw179_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
            d_1054_v128_ = nw179_
            index200_ = default__.safeIndex(414, (d_1054_v128_).length(0))
            (d_1054_v128_)[index200_] = d_1026_v102_
            index201_ = default__.safeIndex(414, (d_1054_v128_).length(0))
            (d_1054_v128_)[index201_] = (d_1026_v102_) + (d_1026_v102_)
            d_1055_v129_: _dafny.Set
            d_1055_v129_ = _dafny.Set({(d_1053_v127_).f12})
            d_1056_v130_: _dafny.Seq
            d_1056_v130_ = _dafny.SeqWithoutIsStrInference([p1, p3])
            d_1057_v131_: _dafny.Array
            nw180_ = _dafny.Array(None, 18)
            nw180_[int(0)] = len(d_1026_v102_)
            nw180_[int(1)] = (0) - (default__.safeModuloInt((self).f8, p2))
            nw180_[int(2)] = p0
            nw180_[int(3)] = p2
            nw180_[int(4)] = p2
            nw180_[int(5)] = (0) - ((self).f8)
            nw180_[int(6)] = (self).f7
            nw180_[int(7)] = (self).f8
            nw180_[int(8)] = len(d_1055_v129_)
            nw180_[int(9)] = len(d_1056_v130_)
            nw180_[int(10)] = -703
            nw180_[int(11)] = default__.safeModuloInt(p2, len(d_1055_v129_))
            nw180_[int(12)] = len((d_1053_v127_).fm2(True, globalState))
            nw180_[int(13)] = p0
            nw180_[int(14)] = ((self).f7) + (p2)
            nw180_[int(15)] = (self).f8
            nw180_[int(16)] = (125) - ((self).f8)
            nw180_[int(17)] = p2
            d_1057_v131_ = nw180_
            d_1057_v131_ = d_1057_v131_
        if (p0) >= (p2):
            d_1058_v132_: _dafny.Seq
            d_1058_v132_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1059_v133_: _dafny.Seq
            d_1059_v133_ = _dafny.SeqWithoutIsStrInference([d_1058_v132_, _dafny.SeqWithoutIsStrInference([431 for d_1060_i11_ in range(default__.abs(774))])])
            d_1061_v134_: _dafny.Map
            d_1061_v134_ = _dafny.Map({(self).f7: len(d_1059_v133_)})
            d_1061_v134_ = (d_1061_v134_).set(-267, ((self).f8) + (p0))
            d_1062_v135_: _dafny.MultiSet
            d_1062_v135_ = _dafny.MultiSet([p1, p3, p3, not(p3), not(p1)])
            d_1063_v136_: _dafny.Map
            d_1063_v136_ = _dafny.Map({d_1062_v135_: p1})
            d_1063_v136_ = (d_1063_v136_).set(d_1062_v135_, ((self).f8) < ((self).f7))
            d_1064_v137_: _dafny.Map
            d_1064_v137_ = _dafny.Map({(self).f8: p3})
            d_1064_v137_ = (d_1064_v137_).set((0) - ((self).f7), True)
            (globalState).f4 = p3
            (globalState).f3 = (p0) > ((self).f7)
        elif True:
            d_1065_v138_: _dafny.Seq
            d_1065_v138_ = _dafny.SeqWithoutIsStrInference([p2])
            d_1066_v139_: _dafny.Seq
            d_1066_v139_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxgvl"))
            d_1067_v140_: _dafny.Map
            d_1067_v140_ = _dafny.Map({len(d_1065_v138_): len(d_1066_v139_)})
            d_1068_v141_: _dafny.Set
            d_1068_v141_ = _dafny.Set({p1, p1, p1, p3, p3})
            d_1069_v142_: _dafny.MultiSet
            d_1069_v142_ = _dafny.MultiSet([d_1068_v141_])
            d_1070_v143_: _dafny.Map
            d_1070_v143_ = _dafny.Map({not(False): True})
            d_1071_v144_: _dafny.MultiSet
            d_1071_v144_ = _dafny.MultiSet([p2])
            d_1072_v145_: _dafny.MultiSet
            d_1072_v145_ = _dafny.MultiSet([d_1071_v144_, d_1071_v144_])
            d_1073_v146_: _dafny.Array
            nw181_ = _dafny.Array(None, 9)
            nw181_[int(0)] = (self).f7
            nw181_[int(1)] = (0) - ((self).f8)
            nw181_[int(2)] = (self).f7
            nw181_[int(3)] = 34
            nw181_[int(4)] = ((d_1072_v145_)[d_1071_v144_] if (d_1071_v144_) in (d_1072_v145_) else p0)
            nw181_[int(5)] = len(_dafny.Map({p1: not(p3)}))
            nw181_[int(6)] = p0
            nw181_[int(7)] = len(d_1066_v139_)
            nw181_[int(8)] = p2
            d_1073_v146_ = nw181_
            d_1074_v147_: _dafny.Map
            d_1074_v147_ = _dafny.Map({d_1073_v146_: p3})
            d_1075_v148_: _dafny.Array
            nw182_ = _dafny.Array(None, 26)
            nw182_[int(0)] = p2
            nw182_[int(1)] = ((d_1067_v140_)[p0] if (p0) in (d_1067_v140_) else p0)
            nw182_[int(2)] = (self).f8
            nw182_[int(3)] = (self).f8
            nw182_[int(4)] = 201
            nw182_[int(5)] = p2
            nw182_[int(6)] = ((d_1069_v142_)[d_1068_v141_] if (d_1068_v141_) in (d_1069_v142_) else len(d_1070_v143_))
            nw182_[int(7)] = p2
            nw182_[int(8)] = -434
            nw182_[int(9)] = (d_1065_v138_)[default__.safeIndex(p0, len(d_1065_v138_))]
            nw182_[int(10)] = (self).fm3(p3, p1, p1, ((d_1070_v143_)[p3] if (p3) in (d_1070_v143_) else p1), globalState)
            nw182_[int(11)] = p2
            nw182_[int(12)] = (self).f7
            nw182_[int(13)] = p2
            nw182_[int(14)] = len(d_1065_v138_)
            nw182_[int(15)] = (self).f7
            nw182_[int(16)] = (self).f7
            nw182_[int(17)] = len(d_1065_v138_)
            nw182_[int(18)] = (0) - (len(d_1074_v147_))
            nw182_[int(19)] = p0
            nw182_[int(20)] = p0
            nw182_[int(21)] = p2
            nw182_[int(22)] = p2
            nw182_[int(23)] = default__.fm20(p3, True, False, globalState)
            nw182_[int(24)] = p2
            nw182_[int(25)] = p0
            d_1075_v148_ = nw182_
            d_1076_v149_: C0
            nw183_ = C0()
            nw183_.ctor__(d_1073_v146_, (self).f7, (self).f8)
            d_1076_v149_ = nw183_
            d_1077_v150_: _dafny.Seq
            d_1077_v150_ = _dafny.SeqWithoutIsStrInference([p1])
            r0 = (d_1077_v150_) + (d_1077_v150_)
            d_1078_v151_: _dafny.Map
            d_1078_v151_ = _dafny.Map({((self).f7) < ((self).f7): (self).f7})
            d_1078_v151_ = (d_1078_v151_).set(p3, (p2) - ((self).f8))
            r1 = p2
            arr14_ = d_1076_v149_.f13
            index202_ = default__.safeIndex(37, (d_1076_v149_.f13).length(0))
            arr14_[index202_] = p2
            arr15_ = d_1076_v149_.f13
            index203_ = default__.safeIndex(37, (d_1076_v149_.f13).length(0))
            arr15_[index203_] = p2
        if p1:
            d_1079_v152_: str
            d_1079_v152_ = _dafny.CodePoint('e')
            d_1080_v153_: _dafny.Seq
            d_1080_v153_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vlqsggw"))
            d_1081_v154_: _dafny.Seq
            d_1081_v154_ = _dafny.SeqWithoutIsStrInference([(d_1080_v153_)[default__.safeIndex(p0, len(d_1080_v153_))], d_1079_v152_, default__.fm5(globalState)])
            d_1079_v152_ = (d_1079_v152_ if p1 else (d_1081_v154_)[default__.safeIndex((self).f8, len(d_1081_v154_))])
            d_1082_v155_: _dafny.Array
            def lambda47_(d_1083_i12_):
                return default__.safeDivisionInt(d_1083_i12_, (self).f8)

            init28_ = lambda47_
            nw184_ = _dafny.Array(None, 13)
            for i0_28_ in range(nw184_.length(0)):
                nw184_[i0_28_] = init28_(i0_28_)
            d_1082_v155_ = nw184_
            d_1084_v156_: _dafny.Map
            d_1084_v156_ = _dafny.Map({d_1082_v155_: p2})
            d_1084_v156_ = (d_1084_v156_).set(d_1082_v155_, ((self).f8 if not(p3) else p0))
            r1 = p0
            if p3:
                d_1085_v157_: C2
                nw185_ = C2()
                nw185_.ctor__((0) - (len(d_1081_v154_)), (self).f8)
                d_1085_v157_ = nw185_
                (globalState).f3 = p1
                d_1082_v155_ = d_1082_v155_
                d_1086_v158_: _dafny.Seq
                d_1086_v158_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1087_v159_: _dafny.Seq
                d_1087_v159_ = _dafny.SeqWithoutIsStrInference([d_1086_v158_])
                d_1088_v160_: _dafny.Seq
                d_1088_v160_ = _dafny.SeqWithoutIsStrInference([(self).f8, p0, len((d_1087_v159_)[default__.safeIndex(p0, len(d_1087_v159_))])])
                d_1089_v161_: _dafny.MultiSet
                d_1089_v161_ = _dafny.MultiSet([p0, (self).f8, (0) - (p0)])
                index204_ = default__.safeIndex(918, (d_1082_v155_).length(0))
                (d_1082_v155_)[index204_] = (0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikupdc")))) * (((d_1089_v161_)[(0) - (p2)] if ((0) - (p2)) in (d_1089_v161_) else len(d_1088_v160_))))
                d_1090_v162_: _dafny.Map
                d_1090_v162_ = _dafny.Map({len(d_1086_v158_): (d_1088_v160_) + (d_1088_v160_)})
                d_1091_v163_: _dafny.Map
                d_1091_v163_ = _dafny.Map({p1: p3})
                d_1092_v164_: _dafny.MultiSet
                d_1092_v164_ = _dafny.MultiSet([(d_1087_v159_)[default__.safeIndex(len(d_1091_v163_), len(d_1087_v159_))]])
                index205_ = default__.safeIndex(918, (d_1082_v155_).length(0))
                rhs198_ = (self).f7
                rhs199_ = ((d_1090_v162_)[p0] if (p0) in (d_1090_v162_) else d_1088_v160_)
                rhs200_ = (((0) - ((d_1092_v164_).cardinality)) + (p2)) - ((self).f7)
                lhs147_ = d_1082_v155_
                lhs148_ = default__.safeIndex(918, (d_1082_v155_).length(0))
                r1 = rhs198_
                d_1088_v160_ = rhs199_
                lhs147_[lhs148_] = rhs200_
                (globalState).f4 = not(p1)
            elif True:
                d_1093_v165_: _dafny.Set
                d_1093_v165_ = _dafny.Set({p3, p1, p1})
                d_1094_v166_: _dafny.Set
                d_1094_v166_ = _dafny.Set({p3, p1, (self).fm1(d_1093_v165_, p2, -90, len(d_1080_v153_), globalState), not(p1), p3})
                d_1095_v167_: D1
                d_1095_v167_ = D1_DC2()
                d_1096_v168_: D1
                d_1096_v168_ = D1_DC3(d_1095_v167_)
                d_1097_v169_: D1
                d_1097_v169_ = D1_DC3(d_1096_v168_)
                d_1098_v170_: D1
                d_1098_v170_ = D1_DC3(d_1097_v169_)
                d_1099_v171_: _dafny.MultiSet
                d_1099_v171_ = _dafny.MultiSet([D1_DC3(d_1097_v169_), d_1098_v170_])
                d_1100_v172_: _dafny.Array
                nw186_ = _dafny.Array(None, 15)
                nw186_[int(0)] = not(p1)
                nw186_[int(1)] = p3
                nw186_[int(2)] = p1
                nw186_[int(3)] = p3
                nw186_[int(4)] = (self).fm1(d_1093_v165_, p2, ((d_1099_v171_)[d_1098_v170_] if (d_1098_v170_) in (d_1099_v171_) else p0), (self).f7, globalState)
                nw186_[int(5)] = p1
                nw186_[int(6)] = p3
                nw186_[int(7)] = p1
                nw186_[int(8)] = p3
                nw186_[int(9)] = p3
                nw186_[int(10)] = p3
                nw186_[int(11)] = p3
                nw186_[int(12)] = p3
                nw186_[int(13)] = p1
                nw186_[int(14)] = p1
                d_1100_v172_ = nw186_
                d_1101_v173_: D6
                d_1101_v173_ = D6_DC15(d_1100_v172_)
                rhs201_ = d_1101_v173_
                rhs202_ = (d_1100_v172_ if p1 else d_1100_v172_)
                rhs203_ = (self).fm4(p3, p3, globalState)
                rhs204_ = (self).f7
                d_1101_v173_ = rhs201_
                d_1100_v172_ = rhs202_
                r1 = rhs203_
                r1 = rhs204_
                index206_ = default__.safeIndex(781, (d_1100_v172_).length(0))
                (d_1100_v172_)[index206_] = False
                d_1102_v174_: bool
                d_1102_v174_ = p3
                d_1103_v175_: _dafny.Array
                nw187_ = _dafny.Array(None, 16)
                nw187_[int(0)] = p1
                nw187_[int(1)] = d_1102_v174_
                nw187_[int(2)] = p1
                nw187_[int(3)] = p1
                nw187_[int(4)] = d_1102_v174_
                nw187_[int(5)] = d_1102_v174_
                nw187_[int(6)] = d_1102_v174_
                nw187_[int(7)] = default__.fm33(globalState)
                nw187_[int(8)] = default__.fm33(globalState)
                nw187_[int(9)] = d_1102_v174_
                nw187_[int(10)] = d_1102_v174_
                nw187_[int(11)] = d_1102_v174_
                nw187_[int(12)] = d_1102_v174_
                nw187_[int(13)] = d_1102_v174_
                nw187_[int(14)] = d_1102_v174_
                nw187_[int(15)] = d_1102_v174_
                d_1103_v175_ = nw187_
                index207_ = default__.safeIndex(781, (d_1100_v172_).length(0))
                rhs205_ = not (p1) or (((self).fm1(d_1094_v166_, (self).f8, (self).f8, (self).f8, globalState)) or (p3))
                rhs206_ = d_1103_v175_
                lhs149_ = d_1100_v172_
                lhs150_ = default__.safeIndex(781, (d_1100_v172_).length(0))
                lhs149_[lhs150_] = rhs205_
                d_1103_v175_ = rhs206_
                d_1104_v176_: _dafny.Map
                d_1104_v176_ = _dafny.Map({p3: p2})
                d_1105_v177_: _dafny.Seq
                d_1105_v177_ = _dafny.SeqWithoutIsStrInference([p3])
                index208_ = default__.safeIndex(781, (d_1100_v172_).length(0))
                index209_ = default__.safeIndex(781, (d_1100_v172_).length(0))
                rhs207_ = (((d_1104_v176_)[p3] if (p3) in (d_1104_v176_) else p0)) >= (default__.safeDivisionInt(p2, (self).f8))
                rhs208_ = (d_1105_v177_)[default__.safeIndex(p2, len(d_1105_v177_))]
                lhs151_ = d_1100_v172_
                lhs152_ = default__.safeIndex(781, (d_1100_v172_).length(0))
                lhs153_ = d_1100_v172_
                lhs154_ = default__.safeIndex(781, (d_1100_v172_).length(0))
                lhs151_[lhs152_] = rhs207_
                lhs153_[lhs154_] = rhs208_
                d_1106_v178_: _dafny.Seq
                d_1106_v178_ = _dafny.SeqWithoutIsStrInference([(d_1081_v154_).set(default__.safeIndex((self).f8, len(d_1081_v154_)), d_1079_v152_), d_1081_v154_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bp")), d_1080_v153_, d_1080_v153_])
                d_1107_v179_: _dafny.Map
                d_1107_v179_ = _dafny.Map({_dafny.MultiSet([d_1080_v153_, (d_1106_v178_)[default__.safeIndex(p2, len(d_1106_v178_))]]): d_1100_v172_})
                d_1108_v180_: _dafny.MultiSet
                d_1108_v180_ = _dafny.MultiSet([(d_1081_v154_).set(default__.safeIndex(p2, len(d_1081_v154_)), d_1079_v152_)])
                d_1109_v181_: _dafny.Set
                d_1109_v181_ = _dafny.Set({((d_1107_v179_)[d_1108_v180_] if (d_1108_v180_) in (d_1107_v179_) else d_1100_v172_)})
                d_1110_v182_: D5
                d_1110_v182_ = D5_DC11(d_1104_v176_)
                rhs209_ = (d_1109_v181_) | (d_1109_v181_)
                rhs210_ = d_1110_v182_
                rhs211_ = (d_1105_v177_)[default__.safeIndex(len((_dafny.SeqWithoutIsStrInference([d_1079_v152_ for d_1111_i13_ in range(default__.abs(205))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_1079_v152_ for d_1112_i13_ in range(default__.abs(205))]))), _dafny.CodePoint('o'))), len(d_1105_v177_))]
                rhs212_ = p2
                lhs155_ = globalState
                d_1109_v181_ = rhs209_
                d_1110_v182_ = rhs210_
                lhs155_.f4 = rhs211_
                r1 = rhs212_
                d_1105_v177_ = d_1105_v177_
            r1 = p2
        elif True:
            (globalState).f4 = p3
            (globalState).f3 = (((self).f8) > (-20)) == (p1)
            r1 = default__.safeDivisionInt(193, p0)
            d_1113_v183_: _dafny.Array
            nw188_ = _dafny.Array(False, 25)
            d_1113_v183_ = nw188_
            d_1114_v184_: D6
            d_1114_v184_ = D6_DC15(d_1113_v183_)
            d_1113_v183_ = (d_1114_v184_).cf23
            d_1115_v185_: _dafny.Seq
            d_1115_v185_ = _dafny.SeqWithoutIsStrInference([p1, p3, True])
            d_1116_v186_: D6
            d_1116_v186_ = D6_DC16((default__.fm35(p2, globalState)).cf24, d_1115_v185_)
            d_1117_v187_: _dafny.Map
            d_1117_v187_ = _dafny.Map({_dafny.MultiSet([p1]): p0})
            d_1118_v188_: _dafny.Set
            d_1118_v188_ = _dafny.Set({D6_DC16(d_1117_v187_, d_1115_v185_)})
            d_1119_v189_: _dafny.Seq
            d_1119_v189_ = _dafny.SeqWithoutIsStrInference([default__.fm34(p1, p3, globalState), (_dafny.Set({d_1116_v186_, d_1116_v186_})).intersection(_dafny.Set({d_1116_v186_, d_1116_v186_})), d_1118_v188_, d_1118_v188_, (d_1118_v188_) | (d_1118_v188_)])
            d_1119_v189_ = _dafny.SeqWithoutIsStrInference([d_1118_v188_ for d_1120_i14_ in range(default__.abs(705))])
        d_1121_v190_: _dafny.Array
        nw189_ = _dafny.Array(False, 24)
        d_1121_v190_ = nw189_
        index210_ = default__.safeIndex(351, (d_1121_v190_).length(0))
        (d_1121_v190_)[index210_] = p3
        index211_ = default__.safeIndex(351, (d_1121_v190_).length(0))
        (d_1121_v190_)[index211_] = p3
        r1 = p2
        d_1122_v191_: _dafny.Seq
        d_1122_v191_ = _dafny.SeqWithoutIsStrInference([p3])
        r0 = d_1122_v191_
        d_1123_v192_: D5
        d_1123_v192_ = D5_DC11((_dafny.Map({(d_1121_v190_)[default__.safeIndex(351, (d_1121_v190_).length(0))]: p2})).set(p1, (self).f8))
        pat_let_tv15_ = p0
        pat_let_tv16_ = d_1121_v190_
        pat_let_tv17_ = d_1121_v190_
        pat_let_tv18_ = p0
        pat_let_tv19_ = p3
        def lambda48_(source23_):
            if source23_.is_DC12:
                d_1124___mcc_h10_ = source23_.cf19
                d_1125_cf19_ = d_1124___mcc_h10_
                return (0) - (pat_let_tv15_)
            elif source23_.is_DC13:
                d_1126___mcc_h11_ = source23_.cf20
                d_1127___mcc_h12_ = source23_.cf21
                d_1128_cf21_ = d_1127___mcc_h12_
                d_1129_cf20_ = d_1126___mcc_h11_
                return 862
            elif source23_.is_DC11:
                d_1130___mcc_h13_ = source23_.cf18
                d_1131_cf18_ = d_1130___mcc_h13_
                return (self).f8
            elif True:
                d_1132___mcc_h14_ = source23_.cf22
                d_1133_cf22_ = d_1132___mcc_h14_
                return default__.safeModuloInt((D2_DC5((D3_DC7((pat_let_tv17_)[default__.safeIndex(351, (pat_let_tv16_).length(0))])).cf10, (self).f7, (self).f8, pat_let_tv18_, (pat_let_tv19_))).cf7, -174)

        r1 = (0) - (lambda48_(d_1123_v192_))
        r2 = not(p3)
        return r0, r1, r2

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        d_1134_v0_: C2
        nw190_ = C2()
        nw190_.ctor__((self).f8, 687)
        d_1134_v0_ = nw190_
        d_1135_v1_: _dafny.Set
        d_1135_v1_ = _dafny.Set({d_1134_v0_, d_1134_v0_, d_1134_v0_, d_1134_v0_, d_1134_v0_})
        d_1136_v2_: _dafny.Map
        d_1136_v2_ = _dafny.Map({(self).f8: (d_1135_v1_).intersection(d_1135_v1_)})
        d_1135_v1_ = ((d_1136_v2_)[p1] if (p1) in (d_1136_v2_) else d_1135_v1_)
        d_1137_v3_: _dafny.Seq
        d_1137_v3_ = _dafny.SeqWithoutIsStrInference([(self).f8])
        r0 = (d_1137_v3_)[default__.safeIndex(p3, len(d_1137_v3_))]
        d_1138_v4_: T1
        nw191_ = C0()
        nw191_.ctor__(p2, (self).f8, (default__.fm20(p0, p0, p0, globalState)) + ((self).f8))
        d_1138_v4_ = nw191_
        d_1138_v4_ = d_1138_v4_
        d_1139_v5_: D4
        d_1139_v5_ = D4_DC10((0) - ((self).f7), ((self).f8) > (p3))
        d_1140_v6_: _dafny.Array
        def lambda49_(d_1141_p0_):
            def lambda50_(d_1142_i0_):
                return (_dafny.MultiSet([d_1141_p0_])) - (_dafny.MultiSet([d_1141_p0_, d_1141_p0_, d_1141_p0_, d_1141_p0_, d_1141_p0_]))

            return lambda50_

        init29_ = lambda49_(p0)
        nw192_ = _dafny.Array(None, 11)
        for i0_29_ in range(nw192_.length(0)):
            nw192_[i0_29_] = init29_(i0_29_)
        d_1140_v6_ = nw192_
        rhs213_ = d_1139_v5_
        rhs214_ = d_1140_v6_
        d_1139_v5_ = rhs213_
        d_1140_v6_ = rhs214_
        hi8_ = -776
        for d_1143_i1_ in range(307, hi8_):
            d_1144_v7_: D7
            d_1144_v7_ = D7_DC19(p1, p0, (d_1138_v4_).f7)
            d_1145_v8_: _dafny.Array
            def lambda51_(d_1146_p0_):
                def lambda52_(d_1147_i2_):
                    return d_1146_p0_

                return lambda52_

            init30_ = lambda51_(p0)
            nw193_ = _dafny.Array(None, 2)
            for i0_30_ in range(nw193_.length(0)):
                nw193_[i0_30_] = init30_(i0_30_)
            d_1145_v8_ = nw193_
            d_1148_v9_: _dafny.MultiSet
            d_1148_v9_ = _dafny.MultiSet([d_1145_v8_, d_1145_v8_, d_1145_v8_, d_1145_v8_])
            d_1149_v10_: _dafny.MultiSet
            d_1149_v10_ = _dafny.MultiSet([len(_dafny.Map({d_1138_v4_: p0}))])
            d_1150_v11_: C4
            nw194_ = C4()
            nw194_.ctor__((d_1149_v10_).ispropersubset(_dafny.MultiSet([(d_1144_v7_).cf29, (d_1148_v9_).cardinality])), (d_1138_v4_).f8, (d_1138_v4_).f8)
            d_1150_v11_ = nw194_
            d_1151_v12_: str
            d_1151_v12_ = _dafny.CodePoint('y')
            d_1151_v12_ = d_1151_v12_
            d_1152_v13_: _dafny.Seq
            d_1152_v13_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_1153_v14_: _dafny.Seq
            d_1153_v14_ = _dafny.SeqWithoutIsStrInference([d_1152_v13_])
            d_1154_v15_: C3
            nw195_ = C3()
            nw195_.ctor__((d_1150_v11_).f11, (d_1138_v4_).f7, (self).f7)
            d_1154_v15_ = nw195_
            d_1155_v16_: _dafny.Set
            d_1155_v16_ = _dafny.Set({d_1154_v15_})
            d_1156_v17_: C3
            nw196_ = C3()
            nw196_.ctor__(not((d_1152_v13_) == ((d_1153_v14_)[default__.safeIndex(p1, len(d_1153_v14_))])), 17, (len(d_1155_v16_)) + (p1))
            d_1156_v17_ = nw196_
            d_1157_v18_: _dafny.Seq
            d_1157_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eu"))
            d_1158_v19_: D8
            d_1158_v19_ = D8_DC21(d_1154_v15_)
            rhs215_ = (d_1158_v19_).cf34
            rhs216_ = (d_1138_v4_).f7
            rhs217_ = ((d_1156_v17_).fm2((d_1156_v17_).f12, globalState)).set(default__.safeIndex((d_1138_v4_).f8, len((d_1156_v17_).fm2((d_1156_v17_).f12, globalState))), (d_1151_v12_ if (d_1154_v15_).f12 else d_1151_v12_))
            d_1156_v17_ = rhs215_
            r0 = rhs216_
            d_1157_v18_ = rhs217_
            r2 = (d_1154_v15_).f12
        d_1159_v20_: str
        d_1159_v20_ = _dafny.CodePoint('a')
        d_1160_v21_: _dafny.Set
        d_1160_v21_ = _dafny.Set({False})
        d_1161_v22_: _dafny.Map
        d_1161_v22_ = _dafny.Map({not(True): d_1160_v21_})
        d_1162_v23_: _dafny.Map
        d_1162_v23_ = _dafny.Map({(d_1138_v4_).f7: ((d_1161_v22_)[True] if (True) in (d_1161_v22_) else d_1160_v21_)})
        d_1163_v24_: _dafny.Seq
        d_1163_v24_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_1164_v25_: _dafny.Map
        d_1164_v25_ = _dafny.Map({d_1159_v20_: ((d_1162_v23_)[len(d_1163_v24_)] if (len(d_1163_v24_)) in (d_1162_v23_) else _dafny.Set({not(p0)}))})
        d_1164_v25_ = (d_1164_v25_).set(d_1159_v20_, d_1160_v21_)
        r0 = 872
        r1 = (d_1138_v4_).f7
        r2 = p0
        return r0, r1, r2

