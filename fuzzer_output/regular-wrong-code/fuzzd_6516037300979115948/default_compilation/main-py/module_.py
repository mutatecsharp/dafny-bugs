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
        return 474

    @staticmethod
    def fm1(p0, p1, globalState):
        return _dafny.Set({_dafny.CodePoint('w'), _dafny.CodePoint('e'), _dafny.CodePoint('v')})

    @staticmethod
    def fm2(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([True, True]))) + ((_dafny.SeqWithoutIsStrInference([False, not(True)]) if False else _dafny.SeqWithoutIsStrInference([not(not(not(True))), True, True, True])))

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        def lambda0_(source0_):
            if source0_.is_DC23:
                d_0___mcc_h0_ = source0_.cf26
                d_1___mcc_h1_ = source0_.cf27
                d_2_cf27_ = d_1___mcc_h1_
                d_3_cf26_ = d_0___mcc_h0_
                return (700) >= (966)
            elif source0_.is_DC22:
                d_4___mcc_h2_ = source0_.cf25
                d_5_cf25_ = d_4___mcc_h2_
                return (_dafny.SeqWithoutIsStrInference([False])) == (_dafny.SeqWithoutIsStrInference([False]))
            elif True:
                d_6___mcc_h3_ = source0_.cf28
                d_7_cf28_ = d_6___mcc_h3_
                return not((_dafny.SeqWithoutIsStrInference([79, -789, 184])) != (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))) for d_8_i0_ in range(default__.abs(158))])))

        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(947, 522):
                d_9_v0_: int = compr_0_
                if ((947) <= (d_9_v0_)) and ((d_9_v0_) < (522)):
                    coll0_[(d_9_v0_) + (150)] = True
            return _dafny.Map(coll0_)
        return not(lambda0_(D10_DC23(_dafny.CodePoint('i'), _dafny.Set({_dafny.MultiSet([len(iife0_()
)]), _dafny.MultiSet([-591])}))))

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqimjoej"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmscsr"))):
            if False:
                return D3_DC6(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvllabupe")))}))
            elif True:
                def iife1_():
                    coll1_ = _dafny.Set()
                    compr_1_: int
                    for compr_1_ in (_dafny.Map({734: -646})).keys.Elements:
                        d_10_v0_: int = compr_1_
                        if (d_10_v0_) in (_dafny.Map({734: -646})):
                            coll1_ = coll1_.union(_dafny.Set([(d_10_v0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owlus"))))]))
                    return _dafny.Set(coll1_)
                return D3_DC6(iife1_()
)
        elif True:
            return D3_DC6(_dafny.Set({len(_dafny.Set({False, not(not(True)), False, not(True)}))}))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([499 for d_11_i0_ in range(default__.abs(133))])), len(_dafny.Map({False: -86}))), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "al")))) + ((_dafny.MultiSet([True])).cardinality), 353])

    @staticmethod
    def fm11(p0, globalState):
        return _dafny.Map({(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxjqdwhx")))) + (len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qeyxpk"))}))): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([987, (0) - ((_dafny.MultiSet([(_dafny.MultiSet([False])).cardinality, 765, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bocyf"))))])).cardinality), len(_dafny.SeqWithoutIsStrInference([False]))]))})

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: _dafny.Seq
            for compr_2_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([499 for d_12_i0_ in range(default__.abs(264))])])).Elements:
                d_13_v0_: _dafny.Seq = compr_2_
                if (d_13_v0_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([499 for d_12_i0_ in range(default__.abs(264))])])):
                    coll2_[d_13_v0_] = False
            return _dafny.Map(coll2_)
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in (_dafny.MultiSet([249, 319])).Elements:
                d_14_v1_: int = compr_3_
                if (d_14_v1_) in (_dafny.MultiSet([249, 319])):
                    coll3_ = coll3_.union(_dafny.Set([default__.safeModuloInt(d_14_v1_, len(_dafny.Set({not(True)})))]))
            return _dafny.Set(coll3_)
        return (iife2_()
        ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({True: (_dafny.MultiSet([len(iife3_()
        ), 484])).cardinality})): True}))]): not(True)}))

    @staticmethod
    def fm13(p0, p1, p2, globalState):
        def iife4_():
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(-131, 215):
                    d_15_v0_: int = compr_6_
                    if ((-131) <= (d_15_v0_)) and ((d_15_v0_) < (215)):
                        coll6_[default__.safeModuloInt(d_15_v0_, (0) - (len(_dafny.SeqWithoutIsStrInference([True]))))] = _dafny.Map({False: False})
                return _dafny.Map(coll6_)
            coll4_ = _dafny.Set()
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(-131, 215):
                    d_15_v0_: int = compr_5_
                    if ((-131) <= (d_15_v0_)) and ((d_15_v0_) < (215)):
                        coll5_[default__.safeModuloInt(d_15_v0_, (0) - (len(_dafny.SeqWithoutIsStrInference([True]))))] = _dafny.Map({False: False})
                return _dafny.Map(coll5_)
            compr_4_: int
            for compr_4_ in (iife5_()
            ).keys.Elements:
                d_16_v1_: int = compr_4_
                if (d_16_v1_) in (iife6_()
                ):
                    coll4_ = coll4_.union(_dafny.Set([(d_16_v1_) + (len(_dafny.Map({False: True})))]))
            return _dafny.Set(coll4_)
        return ((len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bajccy")): len(_dafny.SeqWithoutIsStrInference([True, not(False)]))}))) < (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({209: not(False)})])))) and ((_dafny.Set({len(_dafny.Set({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')]))}))}))})).issubset(iife4_()
        ))

    @staticmethod
    def fm14(p0, p1, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(607, -661):
                d_17_v0_: int = compr_7_
                if ((607) <= (d_17_v0_)) and ((d_17_v0_) < (-661)):
                    coll7_[default__.safeDivisionInt(d_17_v0_, -242)] = len(_dafny.Map({False: 4}))
            return _dafny.Map(coll7_)
        return ((iife7_()
        ) | (_dafny.Map({len(_dafny.Map({365: (0) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qle"))): 508})))})): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_18_i0_ in range(default__.abs(231))]))}))) | (_dafny.Map({134: len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkuy"))))]))}))

    @staticmethod
    def fm15(p0, p1, p2, p3, globalState):
        return (_dafny.Set({D4_DC8(_dafny.Map({-446: True}))})) - (_dafny.Set({D4_DC8(_dafny.Map({504: not(False)}))}))

    @staticmethod
    def fm16(p0, p1, p2, globalState):
        return (_dafny.Set({314})).intersection((_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwnognu"))), -993, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fudqsvf")))})) | (_dafny.Set({10})))

    @staticmethod
    def fm17(p0, globalState):
        return _dafny.CodePoint('e')

    @staticmethod
    def fm18(globalState):
        return (_dafny.Map({True: len(_dafny.Map({_dafny.CodePoint('p'): not(not(True))}))})) | ((_dafny.Map({False: 332})) | (_dafny.Map({False: len(_dafny.Map({777: _dafny.MultiSet([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([419]))).cardinality])}))})))

    @staticmethod
    def fm21(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([679])])

    @staticmethod
    def fm22(p0, p1, globalState):
        return not(((True) == (True)) or (True))

    @staticmethod
    def fm26(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_19_i0_ in range(default__.abs(647))])

    @staticmethod
    def fm27(globalState):
        return _dafny.CodePoint('a')

    @staticmethod
    def fm28(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfhckddj"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_20_i0_ in range(default__.abs(808))]))

    @staticmethod
    def fm31(globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))

    @staticmethod
    def fm32(p0, p1, globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smuuw"))]))) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgxqmv")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_21_i0_ in range(default__.abs(-33))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kaql")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axecgxsb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))]))).intersection((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkadtamt")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_22_i1_ in range(default__.abs(90))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_23_i2_ in range(default__.abs(749))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_24_i3_ in range(default__.abs(-965))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cheyv"))])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvjngnya"))]))))

    @staticmethod
    def fm33(p0, globalState):
        source1_ = D18_DC41(_dafny.MultiSet([False, False]))
        if source1_.is_DC42:
            d_25___mcc_h0_ = source1_.cf58
            d_26___mcc_h1_ = source1_.cf59
            d_27_cf59_ = d_26___mcc_h1_
            d_28_cf58_ = d_25___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_28_cf58_]) for d_29_i0_ in range(default__.abs(833))])
        elif source1_.is_DC43:
            d_30___mcc_h2_ = source1_.cf60
            d_31___mcc_h3_ = source1_.cf61
            d_32_cf61_ = d_31___mcc_h3_
            d_33_cf60_ = d_30___mcc_h2_
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_33_cf60_]), _dafny.SeqWithoutIsStrInference([d_33_cf60_, d_33_cf60_])])
        elif source1_.is_DC41:
            d_34___mcc_h4_ = source1_.cf57
            d_35_cf57_ = d_34___mcc_h4_
            return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False, True, False])]))
        elif True:
            d_36___mcc_h5_ = source1_.cf62
            d_37_cf62_ = d_36___mcc_h5_
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False, True])])

    @staticmethod
    def fm34(p0, p1, p2, globalState):
        return (_dafny.MultiSet([True])) | ((_dafny.MultiSet([True, False, False])) | (_dafny.MultiSet([True])))

    @staticmethod
    def fm35(p0, p1, globalState):
        return _dafny.Map({not (False) or (False): _dafny.CodePoint('l')})

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        if True:
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))) - (_dafny.MultiSet([True]))
        elif True:
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).intersection(_dafny.MultiSet([True]))

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        source2_ = D10_DC24(D10_DC24(D10_DC22(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfyfaev"))]))))
        if source2_.is_DC23:
            d_38___mcc_h0_ = source2_.cf26
            d_39___mcc_h1_ = source2_.cf27
            d_40_cf27_ = d_39___mcc_h1_
            d_41_cf26_ = d_38___mcc_h0_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtdxglg"))
        elif source2_.is_DC22:
            d_42___mcc_h2_ = source2_.cf25
            d_43_cf25_ = d_42___mcc_h2_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gibnkx"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrdqk")))
        elif True:
            d_44___mcc_h3_ = source2_.cf28
            d_45_cf28_ = d_44___mcc_h3_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idi"))

    @staticmethod
    def fm38(p0, globalState):
        if (True) or (True):
            return not(False)
        elif True:
            return (_dafny.SeqWithoutIsStrInference([D8_DC18(), D8_DC18()])) == (_dafny.SeqWithoutIsStrInference([D8_DC18()]))

    @staticmethod
    def fm39(p0, p1, globalState):
        source3_ = D17_DC40(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrblsbu")), True, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_46_i0_ in range(default__.abs(913))])): True})))
        if source3_.is_DC40:
            d_47___mcc_h0_ = source3_.cf54
            d_48___mcc_h1_ = source3_.cf55
            d_49___mcc_h2_ = source3_.cf56
            d_50_cf56_ = d_49___mcc_h2_
            d_51_cf55_ = d_48___mcc_h1_
            d_52_cf54_ = d_47___mcc_h0_
            if not(d_51_cf55_):
                return _dafny.CodePoint('p')
            elif True:
                return _dafny.CodePoint('g')
        elif True:
            d_53___mcc_h3_ = source3_.cf53
            d_54_cf53_ = d_53___mcc_h3_
            return _dafny.CodePoint('r')

    @staticmethod
    def fm40(p0, p1, globalState):
        def iife8_():
            coll8_ = _dafny.Set()
            compr_8_: int
            for compr_8_ in (_dafny.Map({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgxwniy")))})): _dafny.SeqWithoutIsStrInference([False])})).keys.Elements:
                d_55_v0_: int = compr_8_
                if (d_55_v0_) in (_dafny.Map({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgxwniy")))})): _dafny.SeqWithoutIsStrInference([False])})):
                    coll8_ = coll8_.union(_dafny.Set([(d_55_v0_) * (-660)]))
            return _dafny.Set(coll8_)
        return (_dafny.Map({len(iife8_()
        ): _dafny.CodePoint('e')})) | (_dafny.Map({-543: _dafny.CodePoint('q')}))

    @staticmethod
    def fm41(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_56_i0_ in range(default__.abs(927))])), len(_dafny.SeqWithoutIsStrInference([False, False]))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([913, len(_dafny.Map({True: (0) - ((_dafny.MultiSet([_dafny.CodePoint('x'), _dafny.CodePoint('l')])).cardinality)}))]) for d_57_i1_ in range(default__.abs(307))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([691 for d_58_i2_ in range(default__.abs(918))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, True])), len(_dafny.SeqWithoutIsStrInference([False, False]))])])))

    @staticmethod
    def fm42(p0, p1, p2, p3, globalState):
        if (750) != (len(_dafny.Map({not(False): -107}))):
            return _dafny.Map({False: 746})
        elif True:
            return _dafny.Map({True: 636})

    @staticmethod
    def fm43(globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udvooqxav"))): (_dafny.Map({False: True})) | (_dafny.Map({False: not(True)}))})

    @staticmethod
    def fm44(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftcn"))) if False else len(_dafny.Map({563: True}))), (0) - (len(_dafny.Map({True: 698})))])

    @staticmethod
    def fm45(p0, p1, p2, p3, globalState):
        return ((D27_DC64(_dafny.Map({not(not(True)): True}), False, -202, _dafny.SeqWithoutIsStrInference([216, -225]), 248)).cf97) | (_dafny.Map({True: not((D20_DC47(False, 78, True)).cf65)}))

    @staticmethod
    def fm46(globalState):
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(403, 935):
                d_60_v0_: int = compr_9_
                if ((403) <= (d_60_v0_)) and ((d_60_v0_) < (935)):
                    coll9_[(d_60_v0_) * ((_dafny.MultiSet([False])).cardinality)] = not(True)
            return _dafny.Map(coll9_)
        return (_dafny.Map({len(_dafny.Map({-145: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_59_i0_ in range(default__.abs(261))])})): False})) | (iife9_()
        )

    @staticmethod
    def fm47(p0, p1, p2, p3, globalState):
        def iife10_():
            def iife12_():
                coll12_ = _dafny.Set()
                compr_12_: int
                for compr_12_ in _dafny.IntegerRange(-757, 657):
                    d_63_v1_: int = compr_12_
                    if ((-757) <= (d_63_v1_)) and ((d_63_v1_) < (657)):
                        coll12_ = coll12_.union(_dafny.Set([default__.safeDivisionInt(d_63_v1_, (0) - (-485))]))
                return _dafny.Set(coll12_)
            coll10_ = _dafny.Map()
            def iife11_():
                coll11_ = _dafny.Set()
                compr_11_: int
                for compr_11_ in _dafny.IntegerRange(-757, 657):
                    d_61_v1_: int = compr_11_
                    if ((-757) <= (d_61_v1_)) and ((d_61_v1_) < (657)):
                        coll11_ = coll11_.union(_dafny.Set([default__.safeDivisionInt(d_61_v1_, (0) - (-485))]))
                return _dafny.Set(coll11_)
            compr_10_: int
            for compr_10_ in (iife11_()
            ).Elements:
                d_62_v0_: int = compr_10_
                if (d_62_v0_) in (iife12_()
                ):
                    coll10_[default__.safeDivisionInt(d_62_v0_, -934)] = len(_dafny.SeqWithoutIsStrInference([841 for d_64_i0_ in range(default__.abs(218))]))
            return _dafny.Map(coll10_)
        def iife13_():
            coll13_ = _dafny.Map()
            compr_13_: _dafny.Map
            for compr_13_ in (_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rikrnkyb"))): len(_dafny.SeqWithoutIsStrInference([-74]))}): False})).keys.Elements:
                d_65_v2_: _dafny.Map = compr_13_
                if (d_65_v2_) in (_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rikrnkyb"))): len(_dafny.SeqWithoutIsStrInference([-74]))}): False})):
                    coll13_[d_65_v2_] = _dafny.Map({len(_dafny.Map({True: False})): -739})
            return _dafny.Map(coll13_)
        def iife14_():
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(-800, 315):
                    d_68_v4_: int = compr_16_
                    if ((-800) <= (d_68_v4_)) and ((d_68_v4_) < (315)):
                        coll16_[(d_68_v4_) + (len(_dafny.Map({57: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ynysll")))})))] = 316
                return _dafny.Map(coll16_)
            coll14_ = _dafny.Map()
            def iife15_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in _dafny.IntegerRange(-800, 315):
                    d_68_v4_: int = compr_15_
                    if ((-800) <= (d_68_v4_)) and ((d_68_v4_) < (315)):
                        coll15_[(d_68_v4_) + (len(_dafny.Map({57: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ynysll")))})))] = 316
                return _dafny.Map(coll15_)
            compr_14_: _dafny.Map
            for compr_14_ in (_dafny.Set({_dafny.Map({938: 368}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True})) for d_67_i2_ in range(default__.abs(268))])): len(_dafny.SeqWithoutIsStrInference([False]))}), iife15_()
            })).Elements:
                d_69_v3_: _dafny.Map = compr_14_
                if (d_69_v3_) in (_dafny.Set({_dafny.Map({938: 368}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True})) for d_67_i2_ in range(default__.abs(268))])): len(_dafny.SeqWithoutIsStrInference([False]))}), iife16_()
                })):
                    def iife17_():
                        coll17_ = _dafny.Map()
                        compr_17_: int
                        for compr_17_ in _dafny.IntegerRange(-76, -119):
                            d_70_v5_: int = compr_17_
                            if ((-76) <= (d_70_v5_)) and ((d_70_v5_) < (-119)):
                                coll17_[(d_70_v5_) * (556)] = 339
                        return _dafny.Map(coll17_)
                    coll14_[d_69_v3_] = iife17_()

            return _dafny.Map(coll14_)
        return ((_dafny.Map({iife10_()
        : _dafny.Map({-606: len(_dafny.Map({False: 111}))})})) | (iife13_()
        )) | ((_dafny.Map({_dafny.Map({-620: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False, True]))])).cardinality}): _dafny.Map({-390: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_66_i1_ in range(default__.abs(410))]))})})) | (iife14_()
        ))

    @staticmethod
    def fm48(p0, p1, p2, p3, globalState):
        return (_dafny.Set({len((D14_DC33(len(_dafny.Map({887: len(_dafny.Set({not(False)}))})), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dydwcktsd")), True)).cf41)})).intersection(_dafny.Set({len((D8_DC17(_dafny.SeqWithoutIsStrInference([True]))).cf19), len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.Map({(0) - ((_dafny.MultiSet([-303, 232])).cardinality): 42}))}))

    @staticmethod
    def fm49(p0, p1, p2, globalState):
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ne")): False})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_71_i0_ in range(default__.abs(-566))]): True}))

    @staticmethod
    def fm50(p0, globalState):
        return D7_DC16()

    @staticmethod
    def fm51(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({False})).intersection(_dafny.Set({True, not(True), False}))).intersection((_dafny.Set({False})) - (_dafny.Set({True, True})))

    @staticmethod
    def fm52(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: _dafny.CodePoint('d')}), _dafny.Map({False: _dafny.CodePoint('v')})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: _dafny.CodePoint('i')}), _dafny.Map({(D18_DC42(False, -61)).cf58: _dafny.CodePoint('f')}), _dafny.Map({False: _dafny.CodePoint('w')}), _dafny.Map({not(False): _dafny.CodePoint('e')})]))

    @staticmethod
    def fm53(globalState):
        return (_dafny.MultiSet([-463])) - ((_dafny.MultiSet([(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([345]))])))), len(_dafny.SeqWithoutIsStrInference([-628 for d_72_i0_ in range(default__.abs(-27))])), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({len(_dafny.Map({-429: 128})): len(_dafny.Map({False: -740}))}))), 605, 22])): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_73_i1_ in range(default__.abs(720))]))]))}))])) | (_dafny.MultiSet([229, -372, 936, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True]))])), len(_dafny.SeqWithoutIsStrInference([881 for d_74_i2_ in range(default__.abs(669))]))])))

    @staticmethod
    def fm54(p0, p1, globalState):
        def iife18_():
            coll18_ = _dafny.Set()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(306, -44):
                d_75_v0_: int = compr_18_
                if ((306) <= (d_75_v0_)) and ((d_75_v0_) < (-44)):
                    coll18_ = coll18_.union(_dafny.Set([default__.safeDivisionInt(d_75_v0_, 32)]))
            return _dafny.Set(coll18_)
        return D4_DC11((_dafny.Set({len(_dafny.Set({not(False)}))})).issubset(_dafny.Set({(_dafny.MultiSet([906, (_dafny.MultiSet([iife18_()
])).cardinality, len(_dafny.Map({not(True): 744})), 476, len(_dafny.Set({True}))])).cardinality})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efvmiht"))))

    @staticmethod
    def fm55(p0, globalState):
        return (_dafny.Set({_dafny.Set({True, False}), _dafny.Set({True, True}), _dafny.Set({False, not(not(False)), True})})) | (_dafny.Set({_dafny.Set({True, False, False})}))

    @staticmethod
    def fm56(p0, p1, p2, p3, globalState):
        def iife19_():
            coll19_ = _dafny.Set()
            compr_19_: int
            for compr_19_ in _dafny.IntegerRange(579, 136):
                d_76_v0_: int = compr_19_
                if ((579) <= (d_76_v0_)) and ((d_76_v0_) < (136)):
                    coll19_ = coll19_.union(_dafny.Set([default__.safeModuloInt(d_76_v0_, 963)]))
            return _dafny.Set(coll19_)
        return D17_DC39((_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adbfyy")))})) | (_dafny.Map({(D9_DC21(D3_DC6(_dafny.Set({965})), _dafny.CodePoint('w'), False)).cf24: len(iife19_()
)})))

    @staticmethod
    def fm57(p0, p1, p2, p3, globalState):
        def iife20_():
            coll20_ = _dafny.Map()
            compr_20_: _dafny.Map
            for compr_20_ in (_dafny.MultiSet([_dafny.Map({True: True}), _dafny.Map({True: False})])).Elements:
                d_77_v0_: _dafny.Map = compr_20_
                if (d_77_v0_) in (_dafny.MultiSet([_dafny.Map({True: True}), _dafny.Map({True: False})])):
                    coll20_[d_77_v0_] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x'), _dafny.CodePoint('b')]))
            return _dafny.Map(coll20_)
        def iife21_():
            coll21_ = _dafny.Map()
            compr_21_: _dafny.Map
            for compr_21_ in (_dafny.Map({_dafny.Map({True: True}): 524})).keys.Elements:
                d_78_v1_: _dafny.Map = compr_21_
                if (d_78_v1_) in (_dafny.Map({_dafny.Map({True: True}): 524})):
                    coll21_[d_78_v1_] = -226
            return _dafny.Map(coll21_)
        return ((iife20_()
        ) | (iife21_()
        )) | (_dafny.Map({_dafny.Map({False: False}): len(_dafny.SeqWithoutIsStrInference([(0) - (-152)]))}))

    @staticmethod
    def fm58(p0, p1, p2, globalState):
        source4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "giripcgbh"))
        d_79___mcc_h0_ = source4_
        d_80_cf0_ = d_79___mcc_h0_
        return _dafny.Map({_dafny.CodePoint('x'): _dafny.CodePoint('d')})

    @staticmethod
    def fm59(p0, globalState):
        return D18_DC41((_dafny.MultiSet([False, False]) if not(not(not(not(False)))) else _dafny.MultiSet([False])))

    @staticmethod
    def fm60(p0, globalState):
        source5_ = D14_DC33(429, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_81_i0_ in range(default__.abs(655))]), not(not(False)))
        if source5_.is_DC33:
            d_82___mcc_h0_ = source5_.cf40
            d_83___mcc_h1_ = source5_.cf41
            d_84___mcc_h2_ = source5_.cf42
            d_85_cf42_ = d_84___mcc_h2_
            d_86_cf41_ = d_83___mcc_h1_
            d_87_cf40_ = d_82___mcc_h0_
            return (_dafny.MultiSet([_dafny.Set({False, d_85_cf42_}), _dafny.Set({False}), _dafny.Set({d_85_cf42_, False, d_85_cf42_, d_85_cf42_})])) | (_dafny.MultiSet([_dafny.Set({not(d_85_cf42_)})]))
        elif source5_.is_DC34:
            d_88___mcc_h3_ = source5_.cf43
            d_89___mcc_h4_ = source5_.cf44
            d_90___mcc_h5_ = source5_.cf45
            d_91___mcc_h6_ = source5_.cf46
            d_92_cf46_ = d_91___mcc_h6_
            d_93_cf45_ = d_90___mcc_h5_
            d_94_cf44_ = d_89___mcc_h4_
            d_95_cf43_ = d_88___mcc_h3_
            return _dafny.MultiSet([_dafny.Set({d_92_cf46_, d_92_cf46_})])
        elif True:
            d_96___mcc_h7_ = source5_.cf39
            d_97_cf39_ = d_96___mcc_h7_
            return (_dafny.MultiSet([_dafny.Set({(d_97_cf39_).f26})])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Set({(d_97_cf39_).f25}) for d_98_i1_ in range(default__.abs(78))])))

    @staticmethod
    def fm61(p0, p1, globalState):
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_99_i0_ in range(default__.abs(771))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jygawouda")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnyrafmlm")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibmng")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_100_i1_ in range(default__.abs(52))])}))) - (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcwf")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnlmfd"))}))

    @staticmethod
    def fm62(p0, p1, p2, p3, globalState):
        return D15_DC36((0) - (((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhlppqcrb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrs")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gfnjf"))}), _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvqa")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pd"))})]))).intersection(_dafny.MultiSet([_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xqyvci")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcioac")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))}), _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chcl"))})]))).cardinality))

    @staticmethod
    def fm63(p0, globalState):
        return ((_dafny.Map({_dafny.CodePoint('l'): True})) | (_dafny.Map({_dafny.CodePoint('h'): not(False)}))) | ((_dafny.Map({_dafny.CodePoint('d'): False})) | (_dafny.Map({_dafny.CodePoint('k'): False})))

    @staticmethod
    def fm64(p0, globalState):
        return D12_DC30(D12_DC29(False, -217, _dafny.SeqWithoutIsStrInference([896, len(_dafny.SeqWithoutIsStrInference([False]))]), True, 673))

    @staticmethod
    def fm65(p0, p1, p2, p3, globalState):
        return ((D37_DC88(_dafny.Map({759: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g')])}))).cf135) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([994])): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_101_i0_ in range(default__.abs(263))])}))

    @staticmethod
    def fm66(globalState):
        return D1_DC1((False if not(not(False)) else False))

    @staticmethod
    def fm67(p0, globalState):
        return ((_dafny.Set({_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.Map({-227: -734})]))})})) - (_dafny.Set({_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([-333 for d_102_i0_ in range(default__.abs(145))]))}), _dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([426 for d_103_i1_ in range(default__.abs(51))]))})}))).intersection(_dafny.Set({_dafny.Map({False: 613}), _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([True]))}), _dafny.Map({True: 546})}))

    @staticmethod
    def fm68(p0, p1, globalState):
        def iife22_():
            def iife24_():
                coll24_ = _dafny.Map()
                compr_24_: str
                for compr_24_ in (_dafny.Map({_dafny.CodePoint('y'): _dafny.CodePoint('h')})).keys.Elements:
                    d_104_v0_: str = compr_24_
                    if (d_104_v0_) in (_dafny.Map({_dafny.CodePoint('y'): _dafny.CodePoint('h')})):
                        coll24_[d_104_v0_] = False
                return _dafny.Map(coll24_)
            coll22_ = _dafny.Set()
            def iife23_():
                coll23_ = _dafny.Map()
                compr_23_: str
                for compr_23_ in (_dafny.Map({_dafny.CodePoint('y'): _dafny.CodePoint('h')})).keys.Elements:
                    d_104_v0_: str = compr_23_
                    if (d_104_v0_) in (_dafny.Map({_dafny.CodePoint('y'): _dafny.CodePoint('h')})):
                        coll23_[d_104_v0_] = False
                return _dafny.Map(coll23_)
            compr_22_: str
            for compr_22_ in (iife23_()
            ).keys.Elements:
                d_105_v1_: str = compr_22_
                if (d_105_v1_) in (iife24_()
                ):
                    coll22_ = coll22_.union(_dafny.Set([d_105_v1_]))
            return _dafny.Set(coll22_)
        return D21_DC49(iife22_()
)

    @staticmethod
    def fm69(p0, p1, globalState):
        return D15_DC37((((D7_DC15(_dafny.MultiSet([(0) - (-495)]))).cf18) - (_dafny.MultiSet([len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdecsoh")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbrrbi"))})), -500, len(_dafny.Map({not(False): _dafny.CodePoint('g')})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tccob"))), len(_dafny.Set({False}))]))).cardinality, D3_DC7(False), True)

    @staticmethod
    def fm70(globalState):
        return D27_DC63(_dafny.Map({_dafny.CodePoint('u'): True}))

    @staticmethod
    def fm71(globalState):
        return D8_DC17((_dafny.SeqWithoutIsStrInference([True, not(False)])) + ((D8_DC17(_dafny.SeqWithoutIsStrInference([not(False)]))).cf19))

    @staticmethod
    def fm72(globalState):
        return ((_dafny.SeqWithoutIsStrInference([D4_DC9(D1_DC2())])) + (_dafny.SeqWithoutIsStrInference([D4_DC9(D1_DC2())]))) + ((_dafny.SeqWithoutIsStrInference([D4_DC9(D1_DC2())])) + (_dafny.SeqWithoutIsStrInference([D4_DC9(D1_DC2()), D4_DC9(D1_DC2()), D4_DC9(D1_DC2())])))

    @staticmethod
    def fm73(p0, globalState):
        return D10_DC22((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")) for d_106_i0_ in range(default__.abs(966))]))).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hogskne")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))])))

    @staticmethod
    def fm74(p0, p1, p2, globalState):
        return _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_107_i0_ in range(default__.abs(192))]): _dafny.SeqWithoutIsStrInference([True])})

    @staticmethod
    def fm75(p0, globalState):
        return _dafny.Map({_dafny.MultiSet([(_dafny.MultiSet([D12_DC29(True, len(_dafny.SeqWithoutIsStrInference([546])), _dafny.SeqWithoutIsStrInference([28, 109, -538, len(_dafny.SeqWithoutIsStrInference([True]))]), False, 308)])).cardinality]): len(_dafny.Set({True}))})

    @staticmethod
    def fm76(p0, p1, p2, globalState):
        return (_dafny.Map({not(False): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjhipg"))})) | (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgg"))}))

    @staticmethod
    def fm77(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('x'): not(False)}))]) if True else _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([-244])))])), _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwanxmr")))) for d_108_i0_ in range(default__.abs(25))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igo")))]), _dafny.SeqWithoutIsStrInference([563 for d_109_i1_ in range(default__.abs(263))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([453]))])])

    @staticmethod
    def fm78(p0, p1, p2, p3, globalState):
        return D7_DC15(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([D33_DC76(_dafny.Set({False, False}))])), -667, 794, len(_dafny.Map({len(_dafny.Set({False})): False})), 902]))

    @staticmethod
    def fm79(globalState):
        return D1_DC2()

    @staticmethod
    def m19(p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Seq = _dafny.Seq({})
        d_110_v0_: bool
        d_110_v0_ = True
        d_111_v1_: _dafny.Set
        d_111_v1_ = _dafny.Set({d_110_v0_, d_110_v0_})
        d_112_v2_: D33
        d_112_v2_ = D33_DC76(d_111_v1_)
        d_113_v4_: _dafny.Map
        d_113_v4_ = _dafny.Map({p1: d_110_v0_})
        d_114_v5_: D3
        d_114_v5_ = D3_DC7(d_110_v0_)
        d_115_v6_: _dafny.Array
        nw0_ = _dafny.Array(None, 27)
        nw0_[int(0)] = (d_112_v2_).cf115
        nw0_[int(1)] = d_111_v1_
        nw0_[int(2)] = d_111_v1_
        def iife25_():
            coll25_ = _dafny.Map()
            compr_25_: str
            for compr_25_ in (d_113_v4_).keys.Elements:
                d_116_v3_: str = compr_25_
                if (d_116_v3_) in (d_113_v4_):
                    coll25_[d_116_v3_] = d_110_v0_
            return _dafny.Map(coll25_)
        nw0_[int(3)] = (_dafny.Set({d_110_v0_})) - (default__.fm51(False, len(iife25_()
        ), p1, d_110_v0_, globalState))
        nw0_[int(4)] = (d_111_v1_).intersection(d_111_v1_)
        nw0_[int(5)] = (default__.fm51(d_110_v0_, p0, p1, d_110_v0_, globalState) if d_110_v0_ else d_111_v1_)
        nw0_[int(6)] = d_111_v1_
        nw0_[int(7)] = d_111_v1_
        nw0_[int(8)] = d_111_v1_
        nw0_[int(9)] = d_111_v1_
        nw0_[int(10)] = d_111_v1_
        nw0_[int(11)] = (d_111_v1_) | (d_111_v1_)
        nw0_[int(12)] = default__.fm51(not(d_110_v0_), 825, p1, not(not(d_110_v0_)), globalState)
        nw0_[int(13)] = _dafny.Set({(d_114_v5_).cf9})
        nw0_[int(14)] = d_111_v1_
        nw0_[int(15)] = default__.fm51(False, p0, p1, d_110_v0_, globalState)
        nw0_[int(16)] = d_111_v1_
        nw0_[int(17)] = d_111_v1_
        nw0_[int(18)] = d_111_v1_
        nw0_[int(19)] = (d_111_v1_) - (d_111_v1_)
        nw0_[int(20)] = _dafny.Set({d_110_v0_, True, d_110_v0_, d_110_v0_})
        nw0_[int(21)] = _dafny.Set({False, d_110_v0_, d_110_v0_, d_110_v0_})
        nw0_[int(22)] = (d_111_v1_).intersection(d_111_v1_)
        nw0_[int(23)] = d_111_v1_
        nw0_[int(24)] = d_111_v1_
        nw0_[int(25)] = d_111_v1_
        nw0_[int(26)] = (d_111_v1_) | (d_111_v1_)
        d_115_v6_ = nw0_
        d_115_v6_ = (D34_DC79(d_115_v6_)).cf121
        if d_110_v0_:
            d_117_v7_: _dafny.Seq
            d_117_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgwsvujkc"))
            d_118_v8_: _dafny.Map
            d_118_v8_ = _dafny.Map({p0: p0})
            d_119_v9_: _dafny.Seq
            d_119_v9_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_120_v10_: _dafny.Seq
            d_120_v10_ = _dafny.SeqWithoutIsStrInference([d_119_v9_, d_119_v9_])
            d_121_v11_: C10
            nw1_ = C10()
            nw1_.ctor__(len(d_117_v7_), p0, _dafny.Map({d_118_v8_: d_118_v8_}), d_117_v7_, d_120_v10_)
            d_121_v11_ = nw1_
            d_122_v12_: _dafny.Seq
            d_122_v12_ = _dafny.SeqWithoutIsStrInference([d_110_v0_])
            d_123_v13_: D27
            d_123_v13_ = D27_DC65(d_110_v0_, (d_117_v7_).set(default__.safeIndex(p0, len(d_117_v7_)), p1), d_122_v12_)
            d_124_v14_: _dafny.Map
            d_124_v14_ = _dafny.Map({d_123_v13_: (d_121_v11_.f13) + (p0)})
            d_125_v16_: _dafny.Map
            d_125_v16_ = _dafny.Map({D27_DC65(d_110_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")), d_122_v12_): d_110_v0_})
            d_126_v18_: _dafny.Seq
            d_126_v18_ = _dafny.SeqWithoutIsStrInference([d_123_v13_])
            def iife26_():
                coll26_ = _dafny.Map()
                compr_26_: D27
                for compr_26_ in (d_125_v16_).keys.Elements:
                    d_127_v15_: D27 = compr_26_
                    if (d_127_v15_) in (d_125_v16_):
                        coll26_[d_127_v15_] = 688
                return _dafny.Map(coll26_)
            def iife27_():
                coll27_ = _dafny.Map()
                compr_27_: D27
                for compr_27_ in (d_126_v18_).Elements:
                    d_128_v17_: D27 = compr_27_
                    if (d_128_v17_) in (d_126_v18_):
                        coll27_[d_128_v17_] = p0
                return _dafny.Map(coll27_)
            d_124_v14_ = ((iife26_()
            ) | (d_124_v14_)) | (iife27_()
            )
            d_129_v19_: _dafny.Array
            nw2_ = _dafny.Array(None, 5)
            d_129_v19_ = nw2_
            d_130_v20_: _dafny.Map
            d_130_v20_ = _dafny.Map({d_129_v19_: (0) - ((p0) * (p0))})
            d_130_v20_ = (d_130_v20_).set(d_129_v19_, (d_121_v11_.f13 if d_110_v0_ else p0))
            d_131_v21_: _dafny.Array
            nw3_ = _dafny.Array(int(0), 5)
            d_131_v21_ = nw3_
            r1 = d_131_v21_
            r0 = d_110_v0_
        elif True:
            if d_110_v0_:
                d_132_v22_: _dafny.Map
                d_132_v22_ = _dafny.Map({p0: p0})
                d_133_v23_: _dafny.Map
                d_133_v23_ = _dafny.Map({d_132_v22_: d_132_v22_})
                d_134_v24_: _dafny.Seq
                d_134_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyhayym"))
                d_135_v25_: _dafny.Seq
                d_135_v25_ = _dafny.SeqWithoutIsStrInference([len(d_134_v24_), 741])
                d_136_v26_: _dafny.Seq
                d_136_v26_ = _dafny.SeqWithoutIsStrInference([d_135_v25_])
                d_137_v27_: C3
                nw4_ = C3()
                nw4_.ctor__(d_110_v0_, p0, len(d_111_v1_), d_133_v23_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovqwg")), d_136_v26_)
                d_137_v27_ = nw4_
                d_138_v28_: _dafny.Seq
                d_138_v28_ = _dafny.SeqWithoutIsStrInference([d_110_v0_])
                d_139_v29_: _dafny.Seq
                d_139_v29_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([default__.fm38((d_137_v27_).f22, globalState), not(d_110_v0_), False]), d_138_v28_])
                d_140_v30_: _dafny.Map
                d_140_v30_ = _dafny.Map({(d_137_v27_).f22: d_110_v0_})
                d_141_v31_: _dafny.Map
                d_141_v31_ = _dafny.Map({d_137_v27_: (d_139_v29_)[default__.safeIndex(len(d_140_v30_), len(d_139_v29_))]})
                d_142_v32_: D35
                d_142_v32_ = D35_DC82(d_141_v31_)
                rhs0_ = d_110_v0_
                rhs1_ = p0
                rhs2_ = (d_142_v32_).cf125
                lhs0_ = globalState
                d_110_v0_ = rhs0_
                lhs0_.f3 = rhs1_
                d_141_v31_ = rhs2_
                d_143_v33_: _dafny.Array
                nw5_ = _dafny.Array(int(0), 22)
                d_143_v33_ = nw5_
                index0_ = default__.safeIndex(593, (d_143_v33_).length(0))
                (d_143_v33_)[index0_] = -543
                d_144_v34_: _dafny.Map
                d_144_v34_ = _dafny.Map({False: d_143_v33_})
                d_145_v35_: D7
                d_145_v35_ = D7_DC15(_dafny.MultiSet(d_135_v25_))
                d_146_v36_: _dafny.Array
                nw6_ = _dafny.Array(None, 4)
                nw6_[int(0)] = default__.fm39(d_145_v35_, (d_137_v27_).f21, globalState)
                nw6_[int(1)] = _dafny.CodePoint('a')
                nw6_[int(2)] = p1
                nw6_[int(3)] = _dafny.CodePoint('v')
                d_146_v36_ = nw6_
                d_147_v37_: _dafny.Map
                d_147_v37_ = _dafny.Map({d_146_v36_: p0})
                d_148_v38_: _dafny.Map
                d_148_v38_ = _dafny.Map({p1: D7_DC15(_dafny.MultiSet(d_135_v25_))})
                d_149_v39_: _dafny.Seq
                d_149_v39_ = _dafny.SeqWithoutIsStrInference([d_148_v38_])
                index1_ = default__.safeIndex(593, (d_143_v33_).length(0))
                rhs3_ = not ((d_137_v27_).f21) or ((d_137_v27_).f21)
                rhs4_ = (p0 if d_110_v0_ else (0) - (len(d_144_v34_)))
                rhs5_ = (d_137_v27_).f22
                rhs6_ = ((d_147_v37_)[d_146_v36_] if (d_146_v36_) in (d_147_v37_) else len((_dafny.SeqWithoutIsStrInference([_dafny.Map({p1: default__.fm78((d_137_v27_).f22, p0, p0, (d_137_v27_).f22, globalState)})])) + (d_149_v39_)))
                lhs1_ = d_143_v33_
                lhs2_ = default__.safeIndex(593, (d_143_v33_).length(0))
                lhs3_ = globalState
                lhs4_ = globalState
                r0 = rhs3_
                lhs1_[lhs2_] = rhs4_
                lhs3_.f3 = rhs5_
                lhs4_.f3 = rhs6_
                index2_ = default__.safeIndex(593, (d_143_v33_).length(0))
                (d_143_v33_)[index2_] = (0) - ((d_137_v27_).f22)
                d_150_v40_: _dafny.MultiSet
                d_150_v40_ = _dafny.MultiSet([d_110_v0_])
                (globalState).f3 = default__.safeModuloInt((d_150_v40_).cardinality, default__.safeModuloInt((d_143_v33_)[default__.safeIndex(593, (d_143_v33_).length(0))], p0))
                d_138_v28_ = (d_138_v28_) + (d_138_v28_)
            elif True:
                d_151_v42_: _dafny.Seq
                d_151_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wp"))
                d_152_v43_: _dafny.Map
                d_152_v43_ = _dafny.Map({len(d_151_v42_): p0})
                d_153_v44_: _dafny.Seq
                d_153_v44_ = _dafny.SeqWithoutIsStrInference([d_152_v43_])
                d_154_v46_: _dafny.Seq
                d_154_v46_ = _dafny.SeqWithoutIsStrInference([p0])
                d_155_v47_: _dafny.Seq
                d_155_v47_ = _dafny.SeqWithoutIsStrInference([d_154_v46_])
                d_156_v48_: C2
                nw7_ = C2()
                def iife28_():
                    coll28_ = _dafny.Map()
                    compr_28_: _dafny.Map
                    for compr_28_ in (d_153_v44_).Elements:
                        d_157_v41_: _dafny.Map = compr_28_
                        if (d_157_v41_) in (d_153_v44_):
                            def iife29_():
                                coll29_ = _dafny.Map()
                                compr_29_: int
                                for compr_29_ in _dafny.IntegerRange(731, 664):
                                    d_158_v45_: int = compr_29_
                                    if ((731) <= (d_158_v45_)) and ((d_158_v45_) < (664)):
                                        coll29_[(d_158_v45_) * (p0)] = (_dafny.MultiSet([p0, p0])).cardinality
                                return _dafny.Map(coll29_)
                            coll28_[d_157_v41_] = iife29_()

                    return _dafny.Map(coll28_)
                nw7_.ctor__(p0, iife28_()
                , d_151_v42_, d_155_v47_)
                d_156_v48_ = nw7_
                d_159_v49_: _dafny.MultiSet
                d_159_v49_ = _dafny.MultiSet([d_156_v48_])
                d_160_v50_: _dafny.Array
                nw8_ = _dafny.Array(None, 20)
                nw8_[int(0)] = (p0) <= (p0)
                nw8_[int(1)] = d_110_v0_
                nw8_[int(2)] = d_110_v0_
                nw8_[int(3)] = not(not(d_110_v0_))
                nw8_[int(4)] = d_110_v0_
                nw8_[int(5)] = (p0) > ((0) - (p0))
                nw8_[int(6)] = d_110_v0_
                nw8_[int(7)] = (not(d_110_v0_)) == (False)
                nw8_[int(8)] = (d_110_v0_ if d_110_v0_ else d_110_v0_)
                nw8_[int(9)] = d_110_v0_
                nw8_[int(10)] = d_110_v0_
                nw8_[int(11)] = not(d_110_v0_)
                nw8_[int(12)] = False
                nw8_[int(13)] = d_110_v0_
                nw8_[int(14)] = (d_159_v49_).issubset(d_159_v49_)
                nw8_[int(15)] = not (d_110_v0_) or (d_110_v0_)
                nw8_[int(16)] = d_110_v0_
                nw8_[int(17)] = default__.fm38(p0, globalState)
                nw8_[int(18)] = d_110_v0_
                nw8_[int(19)] = d_110_v0_
                d_160_v50_ = nw8_
                index3_ = default__.safeIndex(960, (d_160_v50_).length(0))
                (d_160_v50_)[index3_] = d_110_v0_
                index4_ = default__.safeIndex(960, (d_160_v50_).length(0))
                (d_160_v50_)[index4_] = d_110_v0_
                (globalState).f3 = p0
                d_161_v51_: _dafny.Map
                d_161_v51_ = _dafny.Map({False: default__.safeDivisionInt(p0, p0)})
                d_161_v51_ = (d_161_v51_).set(((d_160_v50_)[default__.safeIndex(960, (d_160_v50_).length(0))] if not(d_110_v0_) else d_110_v0_), (0) - (p0))
                d_162_v52_: _dafny.MultiSet
                d_162_v52_ = _dafny.MultiSet([(d_160_v50_)[default__.safeIndex(960, (d_160_v50_).length(0))], d_110_v0_, (d_160_v50_)[default__.safeIndex(960, (d_160_v50_).length(0))], (d_160_v50_)[default__.safeIndex(960, (d_160_v50_).length(0))]])
                d_163_v53_: _dafny.Array
                nw9_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_163_v53_ = nw9_
                d_164_v54_: C5
                nw10_ = C5()
                nw10_.ctor__(((d_162_v52_) - (default__.fm36(p0, (d_160_v50_)[default__.safeIndex(960, (d_160_v50_).length(0))], len(_dafny.Map({d_110_v0_: p0})), globalState))).cardinality, d_163_v53_)
                d_164_v54_ = nw10_
                d_165_v55_: _dafny.Map
                d_165_v55_ = _dafny.Map({d_110_v0_: d_110_v0_})
                d_166_v56_: _dafny.Array
                nw11_ = _dafny.Array(None, 25)
                nw11_[int(0)] = p0
                nw11_[int(1)] = (0) - ((d_164_v54_).f17)
                nw11_[int(2)] = (d_164_v54_).f17
                nw11_[int(3)] = (d_164_v54_).f17
                nw11_[int(4)] = (d_154_v46_)[default__.safeIndex((d_164_v54_).fm23(globalState), len(d_154_v46_))]
                nw11_[int(5)] = (0) - ((p0) + (len(_dafny.Map({not((d_160_v50_)[default__.safeIndex(960, (d_160_v50_).length(0))]): p0}))))
                nw11_[int(6)] = (d_164_v54_).f17
                nw11_[int(7)] = (0) - ((p0) - (len(d_151_v42_)))
                nw11_[int(8)] = (d_164_v54_).f17
                nw11_[int(9)] = (d_164_v54_).f17
                nw11_[int(10)] = default__.safeModuloInt(552, -753)
                nw11_[int(11)] = default__.safeDivisionInt((d_164_v54_).f17, p0)
                nw11_[int(12)] = (d_164_v54_).f17
                nw11_[int(13)] = len(d_151_v42_)
                nw11_[int(14)] = (len((d_151_v42_).set(default__.safeIndex(412, len(d_151_v42_)), p1))) - ((d_164_v54_).f17)
                nw11_[int(15)] = p0
                nw11_[int(16)] = 443
                nw11_[int(17)] = default__.safeModuloInt(p0, 329)
                nw11_[int(18)] = (p0 if ((d_165_v55_)[True] if (True) in (d_165_v55_) else not((d_160_v50_)[default__.safeIndex(960, (d_160_v50_).length(0))])) else (d_164_v54_).f17)
                nw11_[int(19)] = p0
                nw11_[int(20)] = (d_164_v54_).f17
                nw11_[int(21)] = (p0) - (p0)
                nw11_[int(22)] = p0
                nw11_[int(23)] = (d_164_v54_).f17
                nw11_[int(24)] = (d_164_v54_).f17
                d_166_v56_ = nw11_
                index5_ = default__.safeIndex(814, (d_166_v56_).length(0))
                (d_166_v56_)[index5_] = (0) - (p0)
                index6_ = default__.safeIndex(814, (d_166_v56_).length(0))
                (d_166_v56_)[index6_] = ((d_164_v54_).f17) * ((d_154_v46_)[default__.safeIndex(p0, len(d_154_v46_))])
            (globalState).f3 = p0
            d_167_v57_: _dafny.Seq
            d_167_v57_ = _dafny.SeqWithoutIsStrInference([d_110_v0_])
            d_168_v58_: _dafny.Map
            d_168_v58_ = _dafny.Map({len(d_167_v57_): p0})
            d_169_v59_: _dafny.Map
            d_169_v59_ = _dafny.Map({d_168_v58_: d_168_v58_})
            d_170_v60_: _dafny.Seq
            d_170_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrrcj"))
            d_171_v61_: _dafny.Seq
            d_171_v61_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_172_v62_: _dafny.Seq
            d_172_v62_ = _dafny.SeqWithoutIsStrInference([d_171_v61_])
            d_173_v63_: C10
            nw12_ = C10()
            nw12_.ctor__(p0, p0, d_169_v59_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))) + (d_170_v60_), (d_172_v62_) + (_dafny.SeqWithoutIsStrInference([d_171_v61_, d_171_v61_, _dafny.SeqWithoutIsStrInference([398 for d_174_i0_ in range(default__.abs(-952))]), d_171_v61_])))
            d_173_v63_ = nw12_
            r0 = default__.fm38(default__.safeDivisionInt(d_173_v63_.f13, p0), globalState)
            d_175_v64_: _dafny.Map
            d_175_v64_ = _dafny.Map({p1: d_173_v63_.f13})
            d_176_v65_: _dafny.MultiSet
            d_176_v65_ = _dafny.MultiSet([d_175_v64_, d_175_v64_, d_175_v64_])
            if ((d_176_v65_).intersection(d_176_v65_)).issubset(d_176_v65_):
                (globalState).f3 = (452 if d_110_v0_ else len(d_170_v60_))
                d_177_v66_: T0
                nw13_ = C7()
                nw13_.ctor__(d_170_v60_, d_172_v62_, d_173_v63_.f13, d_169_v59_)
                d_177_v66_ = nw13_
                d_178_v67_: D32
                d_178_v67_ = D32_DC72(d_177_v66_)
                d_179_v68_: C7
                nw14_ = C7()
                nw14_.ctor__((d_177_v66_).f9, d_172_v62_, d_173_v63_.f13, d_169_v59_)
                d_179_v68_ = nw14_
                d_180_v69_: _dafny.Seq
                d_180_v69_ = _dafny.SeqWithoutIsStrInference([d_179_v68_, d_179_v68_, d_179_v68_, d_179_v68_, d_179_v68_])
                d_181_v70_: _dafny.Array
                nw15_ = _dafny.Array(None, 27)
                nw15_[int(0)] = d_177_v66_
                nw15_[int(1)] = (d_178_v67_).cf111
                nw15_[int(2)] = d_177_v66_
                nw15_[int(3)] = d_177_v66_
                nw15_[int(4)] = d_177_v66_
                nw15_[int(5)] = d_177_v66_
                nw15_[int(6)] = d_177_v66_
                nw15_[int(7)] = d_177_v66_
                nw15_[int(8)] = d_177_v66_
                nw15_[int(9)] = d_177_v66_
                nw15_[int(10)] = (d_177_v66_ if default__.fm8(len(d_180_v69_), d_110_v0_, d_110_v0_, globalState) else d_177_v66_)
                nw15_[int(11)] = (d_177_v66_ if d_110_v0_ else d_177_v66_)
                nw15_[int(12)] = d_177_v66_
                nw15_[int(13)] = d_177_v66_
                nw15_[int(14)] = d_177_v66_
                nw15_[int(15)] = d_177_v66_
                nw15_[int(16)] = d_177_v66_
                nw15_[int(17)] = d_177_v66_
                nw15_[int(18)] = d_177_v66_
                nw15_[int(19)] = d_177_v66_
                nw15_[int(20)] = d_177_v66_
                nw15_[int(21)] = d_177_v66_
                nw15_[int(22)] = d_177_v66_
                nw15_[int(23)] = d_177_v66_
                nw15_[int(24)] = d_177_v66_
                nw15_[int(25)] = (d_177_v66_ if d_110_v0_ else d_177_v66_)
                nw15_[int(26)] = d_177_v66_
                d_181_v70_ = nw15_
                index7_ = default__.safeIndex(161, (d_181_v70_).length(0))
                (d_181_v70_)[index7_] = d_177_v66_
                d_182_v71_: _dafny.Map
                d_182_v71_ = _dafny.Map({d_110_v0_: d_173_v63_.f13})
                d_183_v72_: _dafny.Set
                d_183_v72_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_110_v0_, False])), p0, d_173_v63_.f13, default__.fm0(d_173_v63_.f13, len(d_182_v71_), globalState)})
                d_184_v76_: _dafny.Seq
                def iife30_():
                    coll30_ = _dafny.Map()
                    compr_30_: int
                    for compr_30_ in _dafny.IntegerRange(392, 66):
                        d_185_v75_: int = compr_30_
                        if ((392) <= (d_185_v75_)) and ((d_185_v75_) < (66)):
                            coll30_[default__.safeModuloInt(d_185_v75_, len(d_167_v57_))] = len(_dafny.SeqWithoutIsStrInference([D27_DC65(d_110_v0_, (d_177_v66_).f9, d_167_v57_)]))
                    return _dafny.Map(coll30_)
                d_184_v76_ = _dafny.SeqWithoutIsStrInference([iife30_()
                , d_168_v58_])
                d_186_v77_: _dafny.Seq
                def iife31_():
                    coll31_ = _dafny.Map()
                    compr_31_: int
                    for compr_31_ in (d_183_v72_).Elements:
                        d_187_v74_: int = compr_31_
                        if (d_187_v74_) in (d_183_v72_):
                            coll31_[(d_187_v74_) * (p0)] = p0
                    return _dafny.Map(coll31_)
                d_186_v77_ = _dafny.SeqWithoutIsStrInference([iife31_()
                , (d_184_v76_)[default__.safeIndex(384, len(d_184_v76_))]])
                d_188_v78_: _dafny.Map
                d_188_v78_ = d_169_v59_
                index8_ = default__.safeIndex(161, (d_181_v70_).length(0))
                nw16_ = C7()
                def iife32_():
                    coll32_ = _dafny.Map()
                    compr_32_: _dafny.Map
                    for compr_32_ in (d_186_v77_).Elements:
                        d_189_v73_: _dafny.Map = compr_32_
                        if (d_189_v73_) in (d_186_v77_):
                            coll32_[d_189_v73_] = d_168_v58_
                    return _dafny.Map(coll32_)
                nw16_.ctor__(((d_177_v66_).f9) + ((d_177_v66_).f9), (d_177_v66_).f10, len(d_183_v72_), (iife32_()
                ) | ((d_188_v78_)))
                (d_181_v70_)[index8_] = nw16_
                d_190_v79_: _dafny.Map
                d_190_v79_ = _dafny.Map({not(d_110_v0_): d_169_v59_})
                d_191_v80_: C2
                nw17_ = C2()
                nw17_.ctor__(-937, ((d_190_v79_)[d_110_v0_] if (d_110_v0_) in (d_190_v79_) else d_169_v59_), (d_177_v66_).f9, (d_172_v62_).set(default__.safeIndex(len(default__.fm2((d_177_v66_).f9, d_173_v63_.f13, globalState)), len(d_172_v62_)), _dafny.SeqWithoutIsStrInference([193 for d_192_i1_ in range(default__.abs(361))])))
                d_191_v80_ = nw17_
                (d_173_v63_).f13 = d_173_v63_.f13
                d_193_v82_: _dafny.Array
                def lambda1_(d_194_v58_, d_195_v63_, d_196_p0_):
                    def lambda2_(d_197_i2_):
                        def iife33_():
                            coll33_ = _dafny.Map()
                            compr_33_: int
                            for compr_33_ in (_dafny.SeqWithoutIsStrInference([d_195_v63_.f13 for d_198_i3_ in range(default__.abs(657))])).Elements:
                                d_199_v81_: int = compr_33_
                                if (d_199_v81_) in (_dafny.SeqWithoutIsStrInference([d_195_v63_.f13 for d_198_i3_ in range(default__.abs(657))])):
                                    coll33_[(d_199_v81_) + (d_196_p0_)] = d_195_v63_.f13
                            return _dafny.Map(coll33_)
                        return (d_194_v58_) | (iife33_()
                        )

                    return lambda2_

                init0_ = lambda1_(d_168_v58_, d_173_v63_, p0)
                nw18_ = _dafny.Array(None, 1)
                for i0_0_ in range(nw18_.length(0)):
                    nw18_[i0_0_] = init0_(i0_0_)
                d_193_v82_ = nw18_
                d_193_v82_ = (d_193_v82_ if default__.fm8(len(d_171_v61_), default__.fm8(d_173_v63_.f13, (d_167_v57_)[default__.safeIndex(d_173_v63_.f13, len(d_167_v57_))], d_110_v0_, globalState), d_110_v0_, globalState) else d_193_v82_)
            elif True:
                d_110_v0_ = d_110_v0_
                d_200_v83_: _dafny.Array
                nw19_ = _dafny.Array(None, 26)
                d_200_v83_ = nw19_
                d_201_v84_: _dafny.Map
                d_201_v84_ = _dafny.Map({d_110_v0_: p0})
                d_202_v85_: C8
                nw20_ = C8()
                nw20_.ctor__(len(d_201_v84_), not(d_110_v0_))
                d_202_v85_ = nw20_
                index9_ = default__.safeIndex(232, (d_200_v83_).length(0))
                (d_200_v83_)[index9_] = d_202_v85_
                index10_ = default__.safeIndex(232, (d_200_v83_).length(0))
                (d_200_v83_)[index10_] = d_202_v85_
                d_203_v86_: _dafny.Map
                d_203_v86_ = _dafny.Map({d_170_v60_: d_110_v0_})
                d_204_v87_: _dafny.Array
                nw21_ = _dafny.Array(False, 29)
                d_204_v87_ = nw21_
                d_205_v88_: _dafny.Map
                d_205_v88_ = _dafny.Map({d_204_v87_: d_110_v0_})
                d_206_v89_: _dafny.Array
                nw22_ = _dafny.Array(int(0), 23)
                d_206_v89_ = nw22_
                d_207_v90_: _dafny.Map
                d_207_v90_ = _dafny.Map({d_206_v89_: d_110_v0_})
                d_208_v91_: _dafny.Array
                nw23_ = _dafny.Array(None, 15)
                nw23_[int(0)] = (default__.fm69(((d_203_v86_)[_dafny.SeqWithoutIsStrInference([((D9_DC21(D3_DC6((D3_DC6(_dafny.Set({(d_202_v85_).f15, p0}))).cf8), p1, d_110_v0_)).cf23) for d_209_i4_ in range(default__.abs(-683))])] if (_dafny.SeqWithoutIsStrInference([((D9_DC21(D3_DC6((D3_DC6(_dafny.Set({(d_202_v85_).f15, p0}))).cf8), p1, d_110_v0_)).cf23) for d_210_i4_ in range(default__.abs(-683))])) in (d_203_v86_) else False), d_170_v60_, globalState)).cf51
                nw23_[int(1)] = ((d_205_v88_)[d_204_v87_] if (d_204_v87_) in (d_205_v88_) else d_110_v0_)
                nw23_[int(2)] = not((d_202_v85_.f16 if d_110_v0_ else d_202_v85_.f16))
                nw23_[int(3)] = d_202_v85_.f16
                nw23_[int(4)] = not(d_202_v85_.f16)
                nw23_[int(5)] = (d_207_v90_) != (d_207_v90_)
                nw23_[int(6)] = (p0) > ((d_202_v85_).f15)
                nw23_[int(7)] = d_202_v85_.f16
                nw23_[int(8)] = (d_202_v85_.f16 if d_202_v85_.f16 else d_202_v85_.f16)
                nw23_[int(9)] = (473) >= (d_173_v63_.f13)
                nw23_[int(10)] = not(not(d_202_v85_.f16))
                nw23_[int(11)] = d_110_v0_
                nw23_[int(12)] = d_202_v85_.f16
                nw23_[int(13)] = not(not(default__.fm22(d_110_v0_, -863, globalState)))
                nw23_[int(14)] = default__.fm8((d_202_v85_).f15, d_110_v0_, d_202_v85_.f16, globalState)
                d_208_v91_ = nw23_
                d_208_v91_ = d_208_v91_
                d_211_v92_: _dafny.Array
                nw24_ = _dafny.Array(None, 13)
                nw24_[int(0)] = p1
                nw24_[int(1)] = _dafny.CodePoint('f')
                nw24_[int(2)] = p1
                nw24_[int(3)] = p1
                nw24_[int(4)] = p1
                nw24_[int(5)] = p1
                nw24_[int(6)] = p1
                nw24_[int(7)] = _dafny.CodePoint('u')
                nw24_[int(8)] = p1
                nw24_[int(9)] = p1
                nw24_[int(10)] = p1
                nw24_[int(11)] = p1
                nw24_[int(12)] = p1
                d_211_v92_ = nw24_
                index11_ = default__.safeIndex(905, (d_211_v92_).length(0))
                (d_211_v92_)[index11_] = p1
                index12_ = default__.safeIndex(275, (d_115_v6_).length(0))
                (d_115_v6_)[index12_] = (d_111_v1_) | (d_111_v1_)
                index13_ = default__.safeIndex(298, (d_204_v87_).length(0))
                (d_204_v87_)[index13_] = False
                index14_ = default__.safeIndex(905, (d_211_v92_).length(0))
                index15_ = default__.safeIndex(275, (d_115_v6_).length(0))
                index16_ = default__.safeIndex(298, (d_204_v87_).length(0))
                rhs7_ = _dafny.CodePoint('g')
                rhs8_ = d_110_v0_
                rhs9_ = not(default__.fm38(713, globalState))
                rhs10_ = default__.fm51(not(d_110_v0_), default__.safeDivisionInt(d_173_v63_.f13, (d_202_v85_).f15), p1, d_110_v0_, globalState)
                rhs11_ = not(d_202_v85_.f16)
                lhs5_ = d_211_v92_
                lhs6_ = default__.safeIndex(905, (d_211_v92_).length(0))
                lhs7_ = d_115_v6_
                lhs8_ = default__.safeIndex(275, (d_115_v6_).length(0))
                lhs9_ = d_204_v87_
                lhs10_ = default__.safeIndex(298, (d_204_v87_).length(0))
                lhs5_[lhs6_] = rhs7_
                d_110_v0_ = rhs8_
                r0 = rhs9_
                lhs7_[lhs8_] = rhs10_
                lhs9_[lhs10_] = rhs11_
                d_212_v93_: C0
                nw25_ = C0()
                nw25_.ctor__(d_110_v0_, (d_204_v87_)[default__.safeIndex(298, (d_204_v87_).length(0))], d_170_v60_, (d_172_v62_))
                d_212_v93_ = nw25_
        d_213_i5_: int
        d_213_i5_ = 0
        with _dafny.label("0"):
            while d_110_v0_:
                with _dafny.c_label("0"):
                    if (d_213_i5_) >= (100):
                        raise _dafny.Break("0")
                    d_213_i5_ = (d_213_i5_) + (1)
                    d_214_v94_: _dafny.Map
                    d_214_v94_ = _dafny.Map({(p0) - (p0): (p0) - (p0)})
                    d_215_v95_: _dafny.Seq
                    d_215_v95_ = _dafny.SeqWithoutIsStrInference([(0) - (p0)])
                    d_214_v94_ = (d_214_v94_).set((0) - (((d_215_v95_)[default__.safeIndex(p0, len(d_215_v95_))]) * (p0)), p0)
                    d_216_v96_: _dafny.Array
                    nw26_ = _dafny.Array(int(0), 20)
                    d_216_v96_ = nw26_
                    d_217_v97_: _dafny.MultiSet
                    d_217_v97_ = _dafny.MultiSet([d_216_v96_, d_216_v96_, d_216_v96_])
                    source6_ = D15_DC35((d_217_v97_) | (d_217_v97_))
                    if source6_.is_DC36:
                        d_218___mcc_h0_ = source6_.cf48
                        d_219_cf48_ = d_218___mcc_h0_
                        d_220_v98_: _dafny.Map
                        d_220_v98_ = _dafny.Map({d_110_v0_: d_110_v0_})
                        d_221_v99_: _dafny.Map
                        d_221_v99_ = _dafny.Map({default__.fm0(d_219_cf48_, d_219_cf48_, globalState): (p0) > ((_dafny.MultiSet([d_220_v98_])).cardinality)})
                        d_221_v99_ = (d_221_v99_).set(d_219_cf48_, not (d_110_v0_) or (d_110_v0_))
                        (globalState).f3 = -141
                        d_110_v0_ = d_110_v0_
                        d_222_v100_: _dafny.MultiSet
                        d_222_v100_ = _dafny.MultiSet([d_219_cf48_])
                        d_223_v101_: _dafny.Array
                        nw27_ = _dafny.Array(None, 13)
                        nw27_[int(0)] = d_110_v0_
                        nw27_[int(1)] = (d_219_cf48_) != (d_219_cf48_)
                        nw27_[int(2)] = d_110_v0_
                        nw27_[int(3)] = (d_110_v0_) and (not(d_110_v0_))
                        nw27_[int(4)] = d_110_v0_
                        nw27_[int(5)] = default__.fm38(p0, globalState)
                        nw27_[int(6)] = d_110_v0_
                        nw27_[int(7)] = d_110_v0_
                        nw27_[int(8)] = d_110_v0_
                        nw27_[int(9)] = not((d_222_v100_).isdisjoint(d_222_v100_))
                        nw27_[int(10)] = d_110_v0_
                        nw27_[int(11)] = ((d_222_v100_).cardinality) > (d_219_cf48_)
                        nw27_[int(12)] = d_110_v0_
                        d_223_v101_ = nw27_
                        d_224_v102_: _dafny.Seq
                        d_224_v102_ = _dafny.SeqWithoutIsStrInference([False])
                        index17_ = default__.safeIndex(315, (d_223_v101_).length(0))
                        (d_223_v101_)[index17_] = not((d_224_v102_)[default__.safeIndex(318, len(d_224_v102_))])
                        index18_ = default__.safeIndex(315, (d_223_v101_).length(0))
                        (d_223_v101_)[index18_] = d_110_v0_
                    elif source6_.is_DC37:
                        d_225___mcc_h1_ = source6_.cf49
                        d_226___mcc_h2_ = source6_.cf50
                        d_227___mcc_h3_ = source6_.cf51
                        d_228_cf51_ = d_227___mcc_h3_
                        d_229_cf50_ = d_226___mcc_h2_
                        d_230_cf49_ = d_225___mcc_h1_
                        d_231_v103_: _dafny.Set
                        d_231_v103_ = _dafny.Set({-144, d_230_cf49_})
                        d_232_v104_: _dafny.Seq
                        d_232_v104_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhutlkcul"))
                        d_233_v105_: _dafny.Seq
                        d_233_v105_ = _dafny.SeqWithoutIsStrInference([d_215_v95_, d_215_v95_])
                        d_234_v106_: C0
                        nw28_ = C0()
                        nw28_.ctor__(d_110_v0_, d_228_cf51_, d_232_v104_, d_233_v105_)
                        d_234_v106_ = nw28_
                        d_235_v107_: _dafny.Map
                        d_235_v107_ = _dafny.Map({(d_231_v103_).ispropersubset(d_231_v103_): (d_232_v104_)[default__.safeIndex((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_234_v106_]))).cardinality, len(d_232_v104_))]})
                        d_236_v108_: _dafny.Seq
                        d_236_v108_ = _dafny.SeqWithoutIsStrInference([d_110_v0_, (_dafny.SeqWithoutIsStrInference([p1 for d_237_i6_ in range(default__.abs(-272))])) <= (d_232_v104_)])
                        d_238_v109_: _dafny.Array
                        nw29_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
                        d_238_v109_ = nw29_
                        index19_ = default__.safeIndex(416, (d_238_v109_).length(0))
                        (d_238_v109_)[index19_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
                        d_239_v110_: D3
                        d_239_v110_ = D3_DC6((d_231_v103_) - (d_231_v103_))
                        d_240_v111_: D27
                        d_240_v111_ = D27_DC65(d_110_v0_, d_232_v104_, d_236_v108_)
                        pat_let_tv0_ = d_231_v103_
                        index20_ = default__.safeIndex(416, (d_238_v109_).length(0))
                        def iife34_(_pat_let0_0):
                            def iife35_(d_241_dt__update__tmp_h0_):
                                def iife36_(_pat_let1_0):
                                    def iife37_(d_242_dt__update_hcf8_h0_):
                                        return D3_DC6(d_242_dt__update_hcf8_h0_)
                                    return iife37_(_pat_let1_0)
                                return iife36_(pat_let_tv0_)
                            return iife35_(_pat_let0_0)
                        rhs12_ = default__.fm35((d_230_cf49_) - (p0), d_230_cf49_, globalState)
                        rhs13_ = ((d_240_v111_).cf104) + ((d_236_v108_) + (d_236_v108_))
                        rhs14_ = (d_232_v104_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrnx"))) + ((d_232_v104_).set(default__.safeIndex(len(d_232_v104_), len(d_232_v104_)), p1)))
                        rhs15_ = iife34_(d_239_v110_)
                        lhs11_ = d_238_v109_
                        lhs12_ = default__.safeIndex(416, (d_238_v109_).length(0))
                        d_235_v107_ = rhs12_
                        d_236_v108_ = rhs13_
                        lhs11_[lhs12_] = rhs14_
                        d_239_v110_ = rhs15_
                        d_243_v112_: _dafny.Array
                        nw30_ = _dafny.Array(_dafny.CodePoint('D'), 1)
                        d_243_v112_ = nw30_
                        d_244_v113_: D6
                        d_244_v113_ = D6_DC13(d_243_v112_)
                        d_245_v114_: C4
                        nw31_ = C4()
                        nw31_.ctor__((d_234_v106_).f26, d_228_cf51_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "golpgbcr")), d_233_v105_)
                        d_245_v114_ = nw31_
                        d_246_v115_: _dafny.Seq
                        d_246_v115_ = _dafny.SeqWithoutIsStrInference([d_245_v114_, d_245_v114_])
                        rhs16_ = (len(d_246_v115_)) + (default__.safeModuloInt(len(((d_238_v109_)[default__.safeIndex(416, (d_238_v109_).length(0))]).set(default__.safeIndex(p0, len((d_238_v109_)[default__.safeIndex(416, (d_238_v109_).length(0))])), p1)), d_230_cf49_))
                        rhs17_ = D6_DC13(d_243_v112_)
                        rhs18_ = d_216_v96_
                        lhs13_ = globalState
                        lhs13_.f3 = rhs16_
                        d_244_v113_ = rhs17_
                        d_216_v96_ = rhs18_
                        d_228_cf51_ = d_228_cf51_
                        (globalState).f3 = len(default__.fm55(default__.safeModuloInt(d_230_cf49_, d_230_cf49_), globalState))
                    elif True:
                        d_247___mcc_h4_ = source6_.cf47
                        d_248_cf47_ = d_247___mcc_h4_
                        (globalState).f3 = default__.fm0(p0, default__.fm0(p0, p0, globalState), globalState)
                        (globalState).f3 = len(_dafny.Map({p0: default__.fm34(p0, default__.fm0(97, len(d_215_v95_), globalState), (_dafny.MultiSet([p0])).cardinality, globalState)}))
                        d_249_v116_: _dafny.Array
                        def lambda3_(d_250_v0_):
                            def lambda4_(d_251_i7_):
                                return d_250_v0_

                            return lambda4_

                        init1_ = lambda3_(d_110_v0_)
                        nw32_ = _dafny.Array(None, 25)
                        for i0_1_ in range(nw32_.length(0)):
                            nw32_[i0_1_] = init1_(i0_1_)
                        d_249_v116_ = nw32_
                        index21_ = default__.safeIndex(848, (d_249_v116_).length(0))
                        (d_249_v116_)[index21_] = d_110_v0_
                        d_252_v117_: D1
                        d_252_v117_ = D1_DC2()
                        d_253_v118_: _dafny.Array
                        nw33_ = _dafny.Array(None, 24)
                        nw33_[int(0)] = d_252_v117_
                        nw33_[int(1)] = d_252_v117_
                        nw33_[int(2)] = D1_DC2()
                        nw33_[int(3)] = d_252_v117_
                        nw33_[int(4)] = d_252_v117_
                        nw33_[int(5)] = d_252_v117_
                        nw33_[int(6)] = d_252_v117_
                        nw33_[int(7)] = d_252_v117_
                        nw33_[int(8)] = d_252_v117_
                        nw33_[int(9)] = d_252_v117_
                        nw33_[int(10)] = d_252_v117_
                        nw33_[int(11)] = default__.fm79(globalState)
                        nw33_[int(12)] = d_252_v117_
                        nw33_[int(13)] = d_252_v117_
                        nw33_[int(14)] = d_252_v117_
                        nw33_[int(15)] = d_252_v117_
                        nw33_[int(16)] = d_252_v117_
                        nw33_[int(17)] = d_252_v117_
                        nw33_[int(18)] = d_252_v117_
                        nw33_[int(19)] = d_252_v117_
                        nw33_[int(20)] = d_252_v117_
                        nw33_[int(21)] = d_252_v117_
                        nw33_[int(22)] = d_252_v117_
                        nw33_[int(23)] = d_252_v117_
                        d_253_v118_ = nw33_
                        d_254_v119_: _dafny.Array
                        nw34_ = _dafny.Array(None, 29)
                        nw34_[int(0)] = d_253_v118_
                        nw34_[int(1)] = d_253_v118_
                        nw34_[int(2)] = d_253_v118_
                        nw34_[int(3)] = d_253_v118_
                        nw34_[int(4)] = d_253_v118_
                        nw34_[int(5)] = d_253_v118_
                        nw34_[int(6)] = d_253_v118_
                        nw34_[int(7)] = d_253_v118_
                        nw34_[int(8)] = d_253_v118_
                        nw34_[int(9)] = d_253_v118_
                        nw34_[int(10)] = d_253_v118_
                        nw34_[int(11)] = d_253_v118_
                        nw34_[int(12)] = d_253_v118_
                        nw34_[int(13)] = d_253_v118_
                        nw34_[int(14)] = d_253_v118_
                        nw34_[int(15)] = d_253_v118_
                        nw34_[int(16)] = d_253_v118_
                        nw34_[int(17)] = d_253_v118_
                        nw34_[int(18)] = d_253_v118_
                        nw34_[int(19)] = d_253_v118_
                        nw34_[int(20)] = d_253_v118_
                        nw34_[int(21)] = d_253_v118_
                        nw34_[int(22)] = d_253_v118_
                        nw34_[int(23)] = d_253_v118_
                        nw34_[int(24)] = d_253_v118_
                        nw34_[int(25)] = d_253_v118_
                        nw34_[int(26)] = d_253_v118_
                        nw34_[int(27)] = d_253_v118_
                        nw34_[int(28)] = d_253_v118_
                        d_254_v119_ = nw34_
                        d_255_v120_: C5
                        nw35_ = C5()
                        nw35_.ctor__(p0, d_254_v119_)
                        d_255_v120_ = nw35_
                        d_256_v121_: _dafny.Map
                        d_256_v121_ = _dafny.Map({d_110_v0_: d_255_v120_})
                        index22_ = default__.safeIndex(48, (d_216_v96_).length(0))
                        (d_216_v96_)[index22_] = default__.safeDivisionInt(p0, len(d_256_v121_))
                        index23_ = default__.safeIndex(848, (d_249_v116_).length(0))
                        index24_ = default__.safeIndex(48, (d_216_v96_).length(0))
                        rhs19_ = d_110_v0_
                        rhs20_ = (d_110_v0_ if d_110_v0_ else d_110_v0_)
                        rhs21_ = d_110_v0_
                        rhs22_ = len((_dafny.Set({d_110_v0_, d_110_v0_}) if (((d_214_v94_)[(d_255_v120_).f17] if ((d_255_v120_).f17) in (d_214_v94_) else (D14_DC33(-427, _dafny.SeqWithoutIsStrInference([p1 for d_257_i8_ in range(default__.abs(13))]), d_110_v0_)).cf40)) > ((d_255_v120_).f17) else (d_111_v1_) | (d_111_v1_)))
                        lhs14_ = d_249_v116_
                        lhs15_ = default__.safeIndex(848, (d_249_v116_).length(0))
                        lhs16_ = d_216_v96_
                        lhs17_ = default__.safeIndex(48, (d_216_v96_).length(0))
                        d_110_v0_ = rhs19_
                        d_110_v0_ = rhs20_
                        lhs14_[lhs15_] = rhs21_
                        lhs16_[lhs17_] = rhs22_
                        d_111_v1_ = (d_111_v1_) - ((d_111_v1_).intersection(d_111_v1_))
                    d_258_v122_: _dafny.Array
                    nw36_ = _dafny.Array(_dafny.Seq({}), 24)
                    d_258_v122_ = nw36_
                    d_259_v125_: _dafny.Seq
                    def iife38_():
                        coll34_ = _dafny.Set()
                        compr_34_: int
                        for compr_34_ in (d_214_v94_).keys.Elements:
                            d_260_v123_: int = compr_34_
                            if (d_260_v123_) in (d_214_v94_):
                                def iife39_():
                                    coll35_ = _dafny.Map()
                                    compr_35_: int
                                    for compr_35_ in _dafny.IntegerRange(-782, 700):
                                        d_261_v124_: int = compr_35_
                                        if ((-782) <= (d_261_v124_)) and ((d_261_v124_) < (700)):
                                            coll35_[(d_261_v124_) - (700)] = False
                                    return _dafny.Map(coll35_)
                                coll34_ = coll34_.union(_dafny.Set([(d_260_v123_) * ((D21_DC50(len(iife39_()
), True, not(True), True)).cf72)]))
                        return _dafny.Set(coll34_)
                    d_259_v125_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p0}), iife38_()
                    ])
                    index25_ = default__.safeIndex(793, (d_258_v122_).length(0))
                    (d_258_v122_)[index25_] = d_259_v125_
                    index26_ = default__.safeIndex(793, (d_258_v122_).length(0))
                    (d_258_v122_)[index26_] = (d_259_v125_) + ((d_259_v125_) + (d_259_v125_))
                    d_262_v126_: _dafny.Seq
                    d_262_v126_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mnfmt"))
                    d_263_v127_: _dafny.MultiSet
                    d_263_v127_ = _dafny.MultiSet([p0, p0, p0, p0, p0])
                    d_264_v128_: _dafny.Seq
                    d_264_v128_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_263_v127_).cardinality])])
                    d_265_v129_: _dafny.Map
                    d_265_v129_ = _dafny.Map({d_214_v94_: d_214_v94_})
                    d_266_v132_: _dafny.Map
                    d_266_v132_ = _dafny.Map({d_214_v94_: p0})
                    d_267_v133_: _dafny.Seq
                    def iife40_():
                        def iife42_():
                            coll38_ = _dafny.Map()
                            compr_38_: _dafny.Map
                            for compr_38_ in (d_266_v132_).keys.Elements:
                                d_268_v131_: _dafny.Map = compr_38_
                                if (d_268_v131_) in (d_266_v132_):
                                    coll38_[d_268_v131_] = p1
                            return _dafny.Map(coll38_)
                        coll36_ = _dafny.Map()
                        def iife41_():
                            coll37_ = _dafny.Map()
                            compr_37_: _dafny.Map
                            for compr_37_ in (d_266_v132_).keys.Elements:
                                d_268_v131_: _dafny.Map = compr_37_
                                if (d_268_v131_) in (d_266_v132_):
                                    coll37_[d_268_v131_] = p1
                            return _dafny.Map(coll37_)
                        compr_36_: _dafny.Map
                        for compr_36_ in (iife41_()
                        ).keys.Elements:
                            d_269_v130_: _dafny.Map = compr_36_
                            if (d_269_v130_) in (iife42_()
                            ):
                                coll36_[d_269_v130_] = d_214_v94_
                        return _dafny.Map(coll36_)
                    d_267_v133_ = _dafny.SeqWithoutIsStrInference([iife40_()
                    , d_265_v129_, _dafny.Map({d_214_v94_: d_214_v94_})])
                    d_270_v134_: C7
                    nw37_ = C7()
                    nw37_.ctor__(d_262_v126_, (d_264_v128_) + (_dafny.SeqWithoutIsStrInference([d_215_v95_ for d_271_i9_ in range(default__.abs(302))])), p0, (d_265_v129_ if d_110_v0_ else (d_267_v133_)[default__.safeIndex(251, len(d_267_v133_))]))
                    d_270_v134_ = nw37_
                    pass
            pass
        d_272_v135_: _dafny.Array
        nw38_ = _dafny.Array(False, 26)
        d_272_v135_ = nw38_
        index27_ = default__.safeIndex(733, (d_272_v135_).length(0))
        (d_272_v135_)[index27_] = d_110_v0_
        d_273_v136_: _dafny.Array
        nw39_ = _dafny.Array(D21.default()(), 18)
        d_273_v136_ = nw39_
        d_274_v137_: _dafny.Seq
        d_274_v137_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjhnucbss"))
        d_275_v138_: _dafny.Set
        d_275_v138_ = _dafny.Set({p1, _dafny.CodePoint('e'), _dafny.CodePoint('i'), (d_274_v137_)[default__.safeIndex(p0, len(d_274_v137_))]})
        d_276_v139_: D21
        d_276_v139_ = D21_DC51(D21_DC49(d_275_v138_))
        index28_ = default__.safeIndex(364, (d_273_v136_).length(0))
        (d_273_v136_)[index28_] = d_276_v139_
        pat_let_tv1_ = p0
        pat_let_tv2_ = p0
        pat_let_tv3_ = p0
        pat_let_tv4_ = p0
        pat_let_tv5_ = d_110_v0_
        pat_let_tv6_ = globalState
        index29_ = default__.safeIndex(733, (d_272_v135_).length(0))
        index30_ = default__.safeIndex(364, (d_273_v136_).length(0))
        def iife43_(_pat_let2_0):
            def iife44_(d_277_dt__update__tmp_h1_):
                def iife45_(_pat_let3_0):
                    def iife46_(d_278_dt__update_hcf76_h0_):
                        return D21_DC51(d_278_dt__update_hcf76_h0_)
                    return iife46_(_pat_let3_0)
                return iife45_(default__.fm68(_dafny.MultiSet([pat_let_tv1_, pat_let_tv2_, pat_let_tv3_, pat_let_tv4_]), pat_let_tv5_, pat_let_tv6_))
            return iife44_(_pat_let2_0)
        rhs23_ = d_110_v0_
        rhs24_ = ((p0) - (p0) if not(d_110_v0_) else len((d_274_v137_) + (d_274_v137_)))
        rhs25_ = iife43_(d_276_v139_)
        rhs26_ = p0
        lhs18_ = d_272_v135_
        lhs19_ = default__.safeIndex(733, (d_272_v135_).length(0))
        lhs20_ = globalState
        lhs21_ = d_273_v136_
        lhs22_ = default__.safeIndex(364, (d_273_v136_).length(0))
        lhs23_ = globalState
        lhs18_[lhs19_] = rhs23_
        lhs20_.f3 = rhs24_
        lhs21_[lhs22_] = rhs25_
        lhs23_.f3 = rhs26_
        r0 = (d_272_v135_)[default__.safeIndex(733, (d_272_v135_).length(0))]
        d_279_v140_: D11
        d_279_v140_ = D11_DC26()
        d_280_v141_: D11
        d_280_v141_ = D11_DC27(d_279_v140_)
        d_281_v142_: D11
        d_281_v142_ = D11_DC27(d_280_v141_)
        d_282_v143_: D11
        d_282_v143_ = D11_DC27(d_279_v140_)
        d_283_v144_: _dafny.Array
        nw40_ = _dafny.Array(None, 23)
        d_283_v144_ = nw40_
        d_284_v145_: _dafny.MultiSet
        d_284_v145_ = _dafny.MultiSet([not((d_272_v135_)[default__.safeIndex(733, (d_272_v135_).length(0))])])
        d_285_v146_: _dafny.Seq
        d_285_v146_ = _dafny.SeqWithoutIsStrInference([(d_272_v135_)[default__.safeIndex(733, (d_272_v135_).length(0))], d_110_v0_, d_110_v0_, d_110_v0_])
        d_286_v147_: _dafny.Map
        d_286_v147_ = _dafny.Map({d_110_v0_: d_110_v0_})
        d_287_v148_: _dafny.Seq
        d_287_v148_ = _dafny.SeqWithoutIsStrInference([(0) - (p0)])
        d_288_v149_: D27
        d_288_v149_ = D27_DC64(d_286_v147_, d_110_v0_, len(d_274_v137_), d_287_v148_, p0)
        d_289_v150_: _dafny.Seq
        d_289_v150_ = _dafny.SeqWithoutIsStrInference([d_287_v148_])
        d_290_v151_: C1
        nw41_ = C1()
        nw41_.ctor__(p0, ((d_284_v145_)[(d_272_v135_)[default__.safeIndex(733, (d_272_v135_).length(0))]] if ((d_272_v135_)[default__.safeIndex(733, (d_272_v135_).length(0))]) in (d_284_v145_) else p0), p0, default__.fm47((d_272_v135_)[default__.safeIndex(733, (d_272_v135_).length(0))], not(not((d_285_v146_)[default__.safeIndex(p0, len(d_285_v146_))])), d_110_v0_, (d_288_v149_).cf98, globalState), d_274_v137_, d_289_v150_)
        d_290_v151_ = nw41_
        index31_ = default__.safeIndex(743, (d_283_v144_).length(0))
        (d_283_v144_)[index31_] = d_290_v151_
        d_291_v152_: _dafny.Array
        nw42_ = _dafny.Array(_dafny.Array(None, 0), 12)
        d_291_v152_ = nw42_
        d_292_v153_: _dafny.Array
        nw43_ = _dafny.Array(int(0), 29)
        d_292_v153_ = nw43_
        index32_ = default__.safeIndex(298, (d_291_v152_).length(0))
        (d_291_v152_)[index32_] = d_292_v153_
        d_293_v154_: _dafny.Map
        d_293_v154_ = _dafny.Map({d_284_v145_: True})
        d_294_v156_: _dafny.Set
        d_294_v156_ = _dafny.Set({d_284_v145_, d_284_v145_})
        index33_ = default__.safeIndex(733, (d_272_v135_).length(0))
        index34_ = default__.safeIndex(743, (d_283_v144_).length(0))
        index35_ = default__.safeIndex(298, (d_291_v152_).length(0))
        def iife47_():
            coll39_ = _dafny.Set()
            compr_39_: _dafny.MultiSet
            for compr_39_ in (d_293_v154_).keys.Elements:
                d_295_v155_: _dafny.MultiSet = compr_39_
                if (d_295_v155_) in (d_293_v154_):
                    coll39_ = coll39_.union(_dafny.Set([d_295_v155_]))
            return _dafny.Set(coll39_)
        rhs27_ = d_110_v0_
        rhs28_ = (d_294_v156_).ispropersubset((iife47_()
        ).intersection(d_294_v156_))
        rhs29_ = d_282_v143_
        rhs30_ = d_290_v151_
        rhs31_ = d_292_v153_
        lhs24_ = d_272_v135_
        lhs25_ = default__.safeIndex(733, (d_272_v135_).length(0))
        lhs26_ = d_283_v144_
        lhs27_ = default__.safeIndex(743, (d_283_v144_).length(0))
        lhs28_ = d_291_v152_
        lhs29_ = default__.safeIndex(298, (d_291_v152_).length(0))
        d_110_v0_ = rhs27_
        lhs24_[lhs25_] = rhs28_
        d_282_v143_ = rhs29_
        lhs26_[lhs27_] = rhs30_
        lhs28_[lhs29_] = rhs31_
        r0 = default__.fm22(True, d_290_v151_.f23, globalState)
        r1 = d_292_v153_
        d_296_v157_: _dafny.Seq
        d_296_v157_ = _dafny.SeqWithoutIsStrInference([d_274_v137_])
        r2 = d_296_v157_
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_297_v0_: int
        d_297_v0_ = -71
        d_298_v1_: _dafny.MultiSet
        d_298_v1_ = _dafny.MultiSet([d_297_v0_, d_297_v0_])
        d_299_v2_: _dafny.Seq
        d_299_v2_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_297_v0_])])
        d_300_globalState_: GlobalState
        nw44_ = GlobalState()
        nw44_.ctor__(d_298_v1_, True, -264, -657, 409, d_299_v2_, False)
        d_300_globalState_ = nw44_
        d_301_v3_: bool
        d_301_v3_ = True
        if (not(d_301_v3_)) or (d_301_v3_):
            d_302_v4_: _dafny.Seq
            d_302_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrle"))
            d_303_v5_: _dafny.Array
            nw45_ = _dafny.Array(None, 7)
            nw45_[int(0)] = (d_297_v0_ if not(True) else d_297_v0_)
            nw45_[int(1)] = (0) - (d_297_v0_)
            nw45_[int(2)] = (default__.fm0(d_297_v0_, d_297_v0_, d_300_globalState_)) - (-538)
            nw45_[int(3)] = (0) - (d_297_v0_)
            nw45_[int(4)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gharaao")))
            nw45_[int(5)] = (0) - (len(d_302_v4_))
            nw45_[int(6)] = d_297_v0_
            d_303_v5_ = nw45_
            d_304_v6_: _dafny.Map
            d_304_v6_ = _dafny.Map({d_301_v3_: d_297_v0_})
            index36_ = default__.safeIndex(566, (d_303_v5_).length(0))
            (d_303_v5_)[index36_] = len(default__.fm1(494, ((d_304_v6_)[d_301_v3_] if (d_301_v3_) in (d_304_v6_) else d_297_v0_), d_300_globalState_))
            index37_ = default__.safeIndex(566, (d_303_v5_).length(0))
            (d_303_v5_)[index37_] = default__.safeModuloInt(d_297_v0_, d_297_v0_)
            d_305_v7_: _dafny.Seq
            d_305_v7_ = d_302_v4_
            d_306_v8_: _dafny.MultiSet
            d_306_v8_ = _dafny.MultiSet([(d_305_v7_), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_307_i0_ in range(default__.abs(-940))]), d_302_v4_, d_302_v4_, d_302_v4_])
            if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhmh"))) not in (((d_306_v8_).set(d_302_v4_, default__.abs(d_297_v0_))).intersection(d_306_v8_)):
                index38_ = default__.safeIndex(566, (d_303_v5_).length(0))
                rhs32_ = ((0) - ((d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))])) - ((default__.fm0((0) - ((d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]), -494, d_300_globalState_)) + ((d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]))
                rhs33_ = d_301_v3_
                rhs34_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_308_i1_ in range(default__.abs(645))])
                lhs30_ = d_303_v5_
                lhs31_ = default__.safeIndex(566, (d_303_v5_).length(0))
                lhs30_[lhs31_] = rhs32_
                d_301_v3_ = rhs33_
                d_302_v4_ = rhs34_
                d_297_v0_ = (d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]
                d_309_v9_: _dafny.Array
                nw46_ = _dafny.Array(_dafny.Seq({}), 23)
                d_309_v9_ = nw46_
                d_310_v10_: _dafny.Seq
                d_310_v10_ = _dafny.SeqWithoutIsStrInference([d_301_v3_, not(d_301_v3_)])
                index39_ = default__.safeIndex(396, (d_309_v9_).length(0))
                (d_309_v9_)[index39_] = d_310_v10_
                index40_ = default__.safeIndex(396, (d_309_v9_).length(0))
                (d_309_v9_)[index40_] = default__.fm2(d_302_v4_, d_297_v0_, d_300_globalState_)
                d_311_v11_: _dafny.Array
                nw47_ = _dafny.Array(False, 25)
                d_311_v11_ = nw47_
                index41_ = default__.safeIndex(355, (d_311_v11_).length(0))
                (d_311_v11_)[index41_] = d_301_v3_
                index42_ = default__.safeIndex(355, (d_311_v11_).length(0))
                (d_311_v11_)[index42_] = not(True)
                d_312_v12_: D1
                d_312_v12_ = D1_DC2()
                d_313_v13_: D4
                d_313_v13_ = D4_DC9(d_312_v12_)
                d_314_v14_: _dafny.Map
                d_314_v14_ = _dafny.Map({d_313_v13_: (d_311_v11_)[default__.safeIndex(355, (d_311_v11_).length(0))]})
                d_315_v15_: _dafny.MultiSet
                d_315_v15_ = _dafny.MultiSet([d_314_v14_])
                d_316_v16_: _dafny.Map
                d_316_v16_ = _dafny.Map({(d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]: ((d_315_v15_)[d_314_v14_] if (d_314_v14_) in (d_315_v15_) else len(d_304_v6_))})
                d_317_v18_: _dafny.Map
                def iife48_():
                    coll40_ = _dafny.Set()
                    compr_40_: int
                    for compr_40_ in (d_316_v16_).keys.Elements:
                        d_318_v17_: int = compr_40_
                        if (d_318_v17_) in (d_316_v16_):
                            coll40_ = coll40_.union(_dafny.Set([default__.safeDivisionInt(d_318_v17_, (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, not(False)]), _dafny.SeqWithoutIsStrInference([True])])).cardinality)]))
                    return _dafny.Set(coll40_)
                d_317_v18_ = _dafny.Map({len(iife48_()
                ): (d_311_v11_)[default__.safeIndex(355, (d_311_v11_).length(0))]})
                d_319_v19_: C11
                nw48_ = C11()
                nw48_.ctor__(d_317_v18_, default__.safeModuloInt(d_297_v0_, (d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]))
                d_319_v19_ = nw48_
            elif True:
                (d_300_globalState_).f3 = (797) + (default__.fm0(d_297_v0_, default__.fm0(d_297_v0_, d_297_v0_, d_300_globalState_), d_300_globalState_))
                d_320_v20_: _dafny.Array
                def lambda5_(d_321_i2_):
                    return _dafny.CodePoint('i')

                init2_ = lambda5_
                nw49_ = _dafny.Array(None, 23)
                for i0_2_ in range(nw49_.length(0)):
                    nw49_[i0_2_] = init2_(i0_2_)
                d_320_v20_ = nw49_
                index43_ = default__.safeIndex(826, (d_320_v20_).length(0))
                (d_320_v20_)[index43_] = _dafny.CodePoint('s')
                d_322_v21_: str
                d_322_v21_ = _dafny.CodePoint('p')
                d_323_v22_: _dafny.Set
                d_323_v22_ = _dafny.Set({d_301_v3_})
                d_324_v23_: C7
                nw50_ = C7()
                nw50_.ctor__(d_302_v4_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_301_v3_, d_301_v3_})), (d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]]) for d_325_i3_ in range(default__.abs(-940))]), len(d_323_v22_), default__.fm47(default__.fm8((d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))], d_301_v3_, d_301_v3_, d_300_globalState_), d_301_v3_, True, d_301_v3_, d_300_globalState_))
                d_324_v23_ = nw50_
                d_326_v24_: _dafny.Set
                d_326_v24_ = _dafny.Set({d_324_v23_, d_324_v23_, d_324_v23_})
                d_327_v25_: _dafny.Seq
                d_327_v25_ = _dafny.SeqWithoutIsStrInference([d_297_v0_, (d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))], (len(d_326_v24_)) + (d_297_v0_), (d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]])
                index44_ = default__.safeIndex(826, (d_320_v20_).length(0))
                index45_ = default__.safeIndex(566, (d_303_v5_).length(0))
                rhs35_ = d_301_v3_
                rhs36_ = (d_322_v21_ if d_301_v3_ else d_322_v21_)
                rhs37_ = len(d_327_v25_)
                lhs32_ = d_320_v20_
                lhs33_ = default__.safeIndex(826, (d_320_v20_).length(0))
                lhs34_ = d_303_v5_
                lhs35_ = default__.safeIndex(566, (d_303_v5_).length(0))
                d_301_v3_ = rhs35_
                lhs32_[lhs33_] = rhs36_
                lhs34_[lhs35_] = rhs37_
                d_297_v0_ = ((_dafny.MultiSet([d_301_v3_, d_301_v3_])).cardinality) - (len(d_302_v4_))
                d_328_v26_: _dafny.Map
                d_328_v26_ = _dafny.Map({(d_320_v20_)[default__.safeIndex(826, (d_320_v20_).length(0))]: _dafny.Set({d_301_v3_, d_301_v3_, not(True), d_301_v3_})})
                d_328_v26_ = (d_328_v26_).set((d_320_v20_)[default__.safeIndex(826, (d_320_v20_).length(0))], (d_323_v22_).intersection(d_323_v22_))
                (d_300_globalState_).f3 = (d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]
            d_329_v27_: D3
            d_329_v27_ = D3_DC7(d_301_v3_)
            if (d_329_v27_).cf9:
                d_330_v28_: _dafny.Map
                d_330_v28_ = _dafny.Map({d_301_v3_: d_302_v4_})
                d_331_v29_: _dafny.Seq
                d_331_v29_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_332_i4_ in range(default__.abs(854))]), d_302_v4_])
                d_333_v30_: _dafny.Array
                nw51_ = _dafny.Array(None, 7)
                nw51_[int(0)] = d_302_v4_
                nw51_[int(1)] = (d_302_v4_) + (d_302_v4_)
                nw51_[int(2)] = (((d_330_v28_)[d_301_v3_] if (d_301_v3_) in (d_330_v28_) else (d_331_v29_)[default__.safeIndex(d_297_v0_, len(d_331_v29_))])) + (d_302_v4_)
                nw51_[int(3)] = (d_302_v4_) + (d_302_v4_)
                nw51_[int(4)] = d_302_v4_
                nw51_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmmhky"))
                nw51_[int(6)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_334_i5_ in range(default__.abs(222))])) + (d_302_v4_)
                d_333_v30_ = nw51_
                d_335_v31_: str
                d_335_v31_ = _dafny.CodePoint('f')
                index46_ = default__.safeIndex(566, (d_303_v5_).length(0))
                rhs38_ = d_333_v30_
                rhs39_ = _dafny.CodePoint('c')
                rhs40_ = (d_297_v0_) * ((d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))])
                lhs36_ = d_303_v5_
                lhs37_ = default__.safeIndex(566, (d_303_v5_).length(0))
                d_333_v30_ = rhs38_
                d_335_v31_ = rhs39_
                lhs36_[lhs37_] = rhs40_
                d_336_v32_: _dafny.Map
                d_336_v32_ = _dafny.Map({(d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]: d_297_v0_})
                d_337_v33_: _dafny.Map
                d_337_v33_ = _dafny.Map({d_336_v32_: d_336_v32_})
                d_338_v34_: _dafny.Map
                d_338_v34_ = _dafny.Map({d_297_v0_: d_302_v4_})
                d_339_v35_: C3
                nw52_ = C3()
                nw52_.ctor__(False, (0) - (d_297_v0_), len(_dafny.SeqWithoutIsStrInference([d_301_v3_, not(d_301_v3_)])), d_337_v33_, default__.fm26(d_297_v0_, d_338_v34_, d_300_globalState_), d_299_v2_)
                d_339_v35_ = nw52_
                d_340_v36_: _dafny.Set
                d_340_v36_ = _dafny.Set({d_301_v3_})
                d_301_v3_ = (d_340_v36_).ispropersubset((d_340_v36_).intersection(d_340_v36_))
                (d_300_globalState_).f3 = ((d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))] if (d_340_v36_).ispropersubset(d_340_v36_) else -443)
                d_341_v37_: _dafny.Array
                nw53_ = _dafny.Array(False, 5)
                d_341_v37_ = nw53_
                index47_ = default__.safeIndex(343, (d_341_v37_).length(0))
                (d_341_v37_)[index47_] = d_301_v3_
                d_342_v38_: _dafny.Set
                d_342_v38_ = _dafny.Set({(d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]})
                index48_ = default__.safeIndex(343, (d_341_v37_).length(0))
                (d_341_v37_)[index48_] = (not((d_339_v35_).f21)) and ((_dafny.Set({-54, d_297_v0_, (d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]})).ispropersubset(d_342_v38_))
            elif True:
                d_343_v39_: _dafny.Array
                nw54_ = _dafny.Array(None, 26)
                d_343_v39_ = nw54_
                d_344_v40_: _dafny.MultiSet
                d_344_v40_ = _dafny.MultiSet([d_343_v39_])
                d_297_v0_ = (default__.safeDivisionInt(len(d_302_v4_), d_297_v0_)) + (((d_344_v40_)[d_343_v39_] if (d_343_v39_) in (d_344_v40_) else (0) - (d_297_v0_)))
                d_301_v3_ = (d_301_v3_) or ((d_297_v0_) >= (d_297_v0_))
                d_345_v41_: _dafny.Array
                nw55_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_345_v41_ = nw55_
                d_346_v43_: _dafny.Array
                def lambda6_(d_347_v4_, d_348_v5_):
                    def lambda7_(d_349_i6_):
                        def iife49_():
                            coll41_ = _dafny.Set()
                            compr_41_: int
                            for compr_41_ in (_dafny.Map({len(_dafny.Map({len(d_347_v4_): False})): (d_348_v5_)[default__.safeIndex(566, (d_348_v5_).length(0))]})).keys.Elements:
                                d_350_v42_: int = compr_41_
                                if (d_350_v42_) in (_dafny.Map({len(_dafny.Map({len(d_347_v4_): False})): (d_348_v5_)[default__.safeIndex(566, (d_348_v5_).length(0))]})):
                                    coll41_ = coll41_.union(_dafny.Set([default__.safeModuloInt(d_350_v42_, (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([505, 543])))))]))
                            return _dafny.Set(coll41_)
                        return iife49_()
                        

                    return lambda7_

                init3_ = lambda6_(d_302_v4_, d_303_v5_)
                nw56_ = _dafny.Array(None, 17)
                for i0_3_ in range(nw56_.length(0)):
                    nw56_[i0_3_] = init3_(i0_3_)
                d_346_v43_ = nw56_
                index49_ = default__.safeIndex(122, (d_345_v41_).length(0))
                (d_345_v41_)[index49_] = d_346_v43_
                index50_ = default__.safeIndex(122, (d_345_v41_).length(0))
                (d_345_v41_)[index50_] = d_346_v43_
                d_351_v44_: str
                d_351_v44_ = _dafny.CodePoint('n')
                d_352_v45_: _dafny.Map
                d_352_v45_ = _dafny.Map({d_351_v44_: d_297_v0_})
                d_352_v45_ = (d_352_v45_).set(d_351_v44_, (d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))])
                d_353_v46_: _dafny.MultiSet
                d_353_v46_ = _dafny.MultiSet([d_301_v3_, d_301_v3_])
                d_354_v47_: _dafny.Seq
                d_354_v47_ = _dafny.SeqWithoutIsStrInference([d_301_v3_, d_301_v3_, d_301_v3_])
                d_355_v48_: _dafny.Map
                d_355_v48_ = _dafny.Map({((d_353_v46_)[d_301_v3_] if (d_301_v3_) in (d_353_v46_) else d_297_v0_): len(d_354_v47_)})
                d_356_v49_: _dafny.Map
                d_356_v49_ = _dafny.Map({d_355_v48_: _dafny.Map({(d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]: d_297_v0_})})
                d_357_v50_: C3
                nw57_ = C3()
                nw57_.ctor__(d_301_v3_, default__.safeModuloInt(d_297_v0_, (d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]), -65, d_356_v49_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vffpxyinc"))) + (d_302_v4_), (d_299_v2_) + (d_299_v2_))
                d_357_v50_ = nw57_
            d_358_v51_: _dafny.Set
            d_358_v51_ = _dafny.Set({d_301_v3_})
            d_359_v52_: _dafny.Seq
            d_359_v52_ = _dafny.SeqWithoutIsStrInference([len(default__.fm37((0) - ((d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))]), (d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))], False, len(d_358_v51_), d_300_globalState_))])
            d_360_v53_: _dafny.Array
            nw58_ = _dafny.Array(None, 6)
            nw58_[int(0)] = d_301_v3_
            nw58_[int(1)] = d_301_v3_
            nw58_[int(2)] = d_301_v3_
            nw58_[int(3)] = d_301_v3_
            nw58_[int(4)] = False
            nw58_[int(5)] = d_301_v3_
            d_360_v53_ = nw58_
            d_361_v54_: _dafny.Map
            d_361_v54_ = _dafny.Map({(d_359_v52_)[default__.safeIndex(default__.fm0((d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))], d_297_v0_, d_300_globalState_), len(d_359_v52_))]: d_360_v53_})
            d_361_v54_ = (d_361_v54_).set((d_297_v0_) - (645), (d_360_v53_ if d_301_v3_ else d_360_v53_))
            d_362_v55_: _dafny.Seq
            d_362_v55_ = _dafny.SeqWithoutIsStrInference([d_302_v4_])
            d_363_v56_: _dafny.Array
            nw59_ = _dafny.Array(None, 18)
            nw59_[int(0)] = (d_302_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")))
            nw59_[int(1)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_364_i7_ in range(default__.abs(560))])) + (d_302_v4_)
            nw59_[int(2)] = d_302_v4_
            nw59_[int(3)] = d_302_v4_
            nw59_[int(4)] = d_302_v4_
            nw59_[int(5)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_365_i8_ in range(default__.abs(709))])
            nw59_[int(6)] = default__.fm31(d_300_globalState_)
            nw59_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnsfde"))
            nw59_[int(8)] = d_302_v4_
            nw59_[int(9)] = (d_302_v4_) + (d_302_v4_)
            nw59_[int(10)] = (d_362_v55_)[default__.safeIndex((d_303_v5_)[default__.safeIndex(566, (d_303_v5_).length(0))], len(d_362_v55_))]
            nw59_[int(11)] = d_302_v4_
            nw59_[int(12)] = d_302_v4_
            nw59_[int(13)] = d_302_v4_
            nw59_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxgpcele"))
            nw59_[int(15)] = (d_302_v4_) + (d_302_v4_)
            nw59_[int(16)] = (d_362_v55_)[default__.safeIndex(d_297_v0_, len(d_362_v55_))]
            nw59_[int(17)] = d_302_v4_
            d_363_v56_ = nw59_
            d_363_v56_ = d_363_v56_
        elif True:
            d_299_v2_ = ((d_299_v2_) + (d_299_v2_)) + (d_299_v2_)
            d_366_v57_: str
            d_366_v57_ = _dafny.CodePoint('o')
            d_367_v58_: _dafny.Map
            d_367_v58_ = _dafny.Map({d_366_v57_: d_297_v0_})
            d_367_v58_ = (d_367_v58_).set(d_366_v57_, 43)
            d_368_v59_: _dafny.Array
            nw60_ = _dafny.Array(_dafny.Seq({}), 24)
            d_368_v59_ = nw60_
            d_368_v59_ = d_368_v59_
            d_369_v60_: _dafny.Map
            d_369_v60_ = _dafny.Map({d_301_v3_: d_301_v3_})
            d_370_v61_: _dafny.Seq
            d_370_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qimtf"))
            d_371_v62_: D4
            d_371_v62_ = D4_DC11(not(d_301_v3_), d_297_v0_)
            d_372_v63_: _dafny.Seq
            d_372_v63_ = _dafny.SeqWithoutIsStrInference([(len(d_370_v61_)) > ((d_371_v62_).cf14)])
            rhs41_ = (d_369_v60_).set(((d_369_v60_)[d_301_v3_] if (d_301_v3_) in (d_369_v60_) else d_301_v3_), d_301_v3_)
            rhs42_ = (d_297_v0_) * ((-51 if False else d_297_v0_))
            rhs43_ = d_372_v63_
            lhs38_ = d_300_globalState_
            d_369_v60_ = rhs41_
            lhs38_.f3 = rhs42_
            d_372_v63_ = rhs43_
            d_373_v64_: _dafny.MultiSet
            d_373_v64_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrccntqou"))])
            d_374_v65_: bool
            d_375_v66_: _dafny.Array
            d_376_v67_: _dafny.Seq
            out0_: bool
            out1_: _dafny.Array
            out2_: _dafny.Seq
            out0_, out1_, out2_ = default__.m19((d_373_v64_).cardinality, d_366_v57_, d_300_globalState_)
            d_374_v65_ = out0_
            d_375_v66_ = out1_
            d_376_v67_ = out2_
        d_301_v3_ = d_301_v3_
        d_377_v68_: _dafny.Array
        nw61_ = _dafny.Array(D1.default()(), 11)
        d_377_v68_ = nw61_
        d_378_v69_: _dafny.Map
        d_378_v69_ = _dafny.Map({d_301_v3_: d_377_v68_})
        d_379_v70_: _dafny.Array
        nw62_ = _dafny.Array(None, 23)
        nw62_[int(0)] = d_377_v68_
        nw62_[int(1)] = ((d_378_v69_)[d_301_v3_] if (d_301_v3_) in (d_378_v69_) else d_377_v68_)
        nw62_[int(2)] = d_377_v68_
        nw62_[int(3)] = d_377_v68_
        nw62_[int(4)] = d_377_v68_
        nw62_[int(5)] = d_377_v68_
        nw62_[int(6)] = d_377_v68_
        nw62_[int(7)] = d_377_v68_
        nw62_[int(8)] = d_377_v68_
        nw62_[int(9)] = d_377_v68_
        nw62_[int(10)] = d_377_v68_
        nw62_[int(11)] = d_377_v68_
        nw62_[int(12)] = d_377_v68_
        nw62_[int(13)] = d_377_v68_
        nw62_[int(14)] = d_377_v68_
        nw62_[int(15)] = d_377_v68_
        nw62_[int(16)] = d_377_v68_
        nw62_[int(17)] = d_377_v68_
        nw62_[int(18)] = d_377_v68_
        nw62_[int(19)] = d_377_v68_
        nw62_[int(20)] = d_377_v68_
        nw62_[int(21)] = d_377_v68_
        nw62_[int(22)] = d_377_v68_
        d_379_v70_ = nw62_
        d_380_v71_: C5
        nw63_ = C5()
        nw63_.ctor__(d_297_v0_, d_379_v70_)
        d_380_v71_ = nw63_
        d_380_v71_ = d_380_v71_
        d_297_v0_ = (d_380_v71_).f17
        d_381_v72_: C5
        nw64_ = C5()
        nw64_.ctor__(4, d_380_v71_.f18)
        d_381_v72_ = nw64_
        d_382_v73_: _dafny.Set
        d_382_v73_ = _dafny.Set({d_301_v3_, d_301_v3_, d_301_v3_, d_301_v3_})
        d_383_i9_: int
        d_383_i9_ = 0
        with _dafny.label("1"):
            while (d_382_v73_).isdisjoint(d_382_v73_):
                with _dafny.c_label("1"):
                    if (d_383_i9_) >= (100):
                        raise _dafny.Break("1")
                    d_383_i9_ = (d_383_i9_) + (1)
                    if d_301_v3_:
                        d_384_v74_: _dafny.Seq
                        d_384_v74_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bx"))
                        d_385_v75_: _dafny.Map
                        d_385_v75_ = _dafny.Map({d_301_v3_: d_301_v3_})
                        d_386_v76_: _dafny.Map
                        d_386_v76_ = _dafny.Map({d_301_v3_: len(default__.fm2(d_384_v74_, len(d_385_v75_), d_300_globalState_))})
                        d_387_v77_: _dafny.Array
                        def lambda8_(d_388_v3_):
                            def lambda9_(d_389_i10_):
                                return d_388_v3_

                            return lambda9_

                        init4_ = lambda8_(d_301_v3_)
                        nw65_ = _dafny.Array(None, 9)
                        for i0_4_ in range(nw65_.length(0)):
                            nw65_[i0_4_] = init4_(i0_4_)
                        d_387_v77_ = nw65_
                        d_390_v78_: _dafny.Seq
                        d_391_v79_: D1
                        d_392_v80_: bool
                        out3_: _dafny.Seq
                        out4_: D1
                        out5_: bool
                        out3_, out4_, out5_ = (d_380_v71_).m11(d_386_v76_, d_387_v77_, ((d_380_v71_).f17) + ((d_380_v71_).f17), d_300_globalState_)
                        d_390_v78_ = out3_
                        d_391_v79_ = out4_
                        d_392_v80_ = out5_
                        d_393_v83_: T0
                        nw66_ = C0()
                        nw66_.ctor__(d_392_v80_, d_392_v80_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_394_i11_ in range(default__.abs(292))]), d_299_v2_)
                        d_393_v83_ = nw66_
                        d_395_v84_: _dafny.Seq
                        d_395_v84_ = _dafny.SeqWithoutIsStrInference([d_393_v83_])
                        d_396_v85_: _dafny.Map
                        def iife50_():
                            coll42_ = _dafny.Map()
                            compr_42_: int
                            for compr_42_ in _dafny.IntegerRange(519, -280):
                                d_397_v82_: int = compr_42_
                                if ((519) <= (d_397_v82_)) and ((d_397_v82_) < (-280)):
                                    coll42_[(d_397_v82_) * ((d_298_v1_).cardinality)] = (d_381_v72_).f17
                            return _dafny.Map(coll42_)
                        d_396_v85_ = _dafny.Map({iife50_()
                        : len(d_395_v84_)})
                        d_398_v86_: _dafny.Seq
                        d_398_v86_ = d_384_v74_
                        d_399_v87_: str
                        d_399_v87_ = _dafny.CodePoint('b')
                        d_400_v88_: C1
                        nw67_ = C1()
                        def iife51_():
                            coll43_ = _dafny.Map()
                            compr_43_: _dafny.Map
                            for compr_43_ in (d_396_v85_).keys.Elements:
                                d_401_v81_: _dafny.Map = compr_43_
                                if (d_401_v81_) in (d_396_v85_):
                                    coll43_[d_401_v81_] = _dafny.Map({(d_380_v71_).f17: (d_381_v72_).f17})
                            return _dafny.Map(coll43_)
                        nw67_.ctor__((d_381_v72_).f17, (d_381_v72_).f17, (d_381_v72_).f17, iife51_()
                        , ((d_398_v86_)) + (((d_393_v83_).f9).set(default__.safeIndex((d_381_v72_).f17, len((d_393_v83_).f9)), d_399_v87_)), (d_393_v83_).f10)
                        d_400_v88_ = nw67_
                        d_402_v89_: C8
                        nw68_ = C8()
                        nw68_.ctor__((d_400_v88_).f24, d_301_v3_)
                        d_402_v89_ = nw68_
                        d_403_v90_: _dafny.Set
                        d_403_v90_ = _dafny.Set({d_399_v87_})
                        d_404_v91_: D21
                        d_404_v91_ = D21_DC49(d_403_v90_)
                        d_405_v92_: _dafny.Seq
                        d_405_v92_ = _dafny.SeqWithoutIsStrInference([D21_DC49(d_403_v90_), d_404_v91_])
                        d_406_v93_: _dafny.Map
                        d_406_v93_ = _dafny.Map({len(d_405_v92_): (d_380_v71_).f17})
                        d_407_v94_: _dafny.Map
                        d_407_v94_ = _dafny.Map({d_406_v93_: _dafny.Map({(d_400_v88_).f24: (d_380_v71_).f17})})
                        d_408_v95_: C10
                        nw69_ = C10()
                        nw69_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "juetfnoru"))), (d_380_v71_).f17, d_407_v94_, d_384_v74_, d_299_v2_)
                        d_408_v95_ = nw69_
                        (d_400_v88_).f23 = (d_381_v72_).f17
                    elif True:
                        d_409_v96_: _dafny.Seq
                        d_409_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kinwo"))
                        d_410_v97_: _dafny.Seq
                        d_410_v97_ = _dafny.SeqWithoutIsStrInference([d_409_v96_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbmwuex"))])
                        d_411_v98_: _dafny.Map
                        d_411_v98_ = _dafny.Map({len(d_410_v97_): d_297_v0_})
                        d_412_v99_: str
                        d_412_v99_ = _dafny.CodePoint('o')
                        d_413_v100_: _dafny.Map
                        d_413_v100_ = _dafny.Map({d_412_v99_: _dafny.Map({d_297_v0_: 314})})
                        d_414_v101_: _dafny.Map
                        d_414_v101_ = _dafny.Map({d_411_v98_: ((d_413_v100_)[default__.fm27(d_300_globalState_)] if (default__.fm27(d_300_globalState_)) in (d_413_v100_) else d_411_v98_)})
                        d_415_v102_: T1
                        nw70_ = C2()
                        nw70_.ctor__((d_381_v72_).f17, d_414_v101_, d_409_v96_, d_299_v2_)
                        d_415_v102_ = nw70_
                        d_416_v103_: _dafny.Seq
                        d_416_v103_ = _dafny.SeqWithoutIsStrInference([d_415_v102_])
                        d_417_v104_: _dafny.Seq
                        d_417_v104_ = _dafny.SeqWithoutIsStrInference([d_297_v0_, (d_380_v71_).f17])
                        d_418_v105_: C1
                        nw71_ = C1()
                        nw71_.ctor__((d_381_v72_).f17, (d_380_v71_).f17, len(d_416_v103_), d_415_v102_.f12, (d_415_v102_).f9, ((d_415_v102_).f10).set(default__.safeIndex((d_381_v72_).f17, len((d_415_v102_).f10)), d_417_v104_))
                        d_418_v105_ = nw71_
                        d_419_v106_: _dafny.Map
                        d_419_v106_ = _dafny.Map({d_418_v105_: ((d_380_v71_).f17 if d_301_v3_ else (d_381_v72_).f17)})
                        d_419_v106_ = (d_419_v106_).set(d_418_v105_, (d_415_v102_).fm5((d_381_v72_).f17, (d_418_v105_).f24, d_300_globalState_))
                        (d_300_globalState_).f3 = (d_381_v72_).f17
                        d_420_v108_: _dafny.Seq
                        def iife52_():
                            coll44_ = _dafny.Map()
                            compr_44_: int
                            for compr_44_ in (_dafny.Set({(d_418_v105_).f24})).Elements:
                                d_421_v107_: int = compr_44_
                                if (d_421_v107_) in (_dafny.Set({(d_418_v105_).f24})):
                                    coll44_[(d_421_v107_) - (d_418_v105_.f23)] = (d_418_v105_).f24
                            return _dafny.Map(coll44_)
                        d_420_v108_ = _dafny.SeqWithoutIsStrInference([d_415_v102_.f12, d_414_v101_, (d_415_v102_.f12).set(d_411_v98_, d_411_v98_), d_414_v101_, _dafny.Map({iife52_()
                        : d_411_v98_})])
                        d_422_v109_: C6
                        nw72_ = C6()
                        nw72_.ctor__(620, (d_420_v108_)[default__.safeIndex((d_381_v72_).f17, len(d_420_v108_))], (d_409_v96_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjurkj"))), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-532 for d_423_i12_ in range(default__.abs(901))])]))
                        d_422_v109_ = nw72_
                        (d_300_globalState_).f3 = (d_381_v72_).f17
                        d_424_v110_: _dafny.Seq
                        d_424_v110_ = _dafny.SeqWithoutIsStrInference([d_301_v3_])
                        rhs44_ = d_424_v110_
                        rhs45_ = (d_381_v72_).f17
                        lhs39_ = d_300_globalState_
                        d_424_v110_ = rhs44_
                        lhs39_.f3 = rhs45_
                    d_425_v111_: _dafny.Seq
                    d_425_v111_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scssmowol"))
                    d_426_v112_: _dafny.Array
                    nw73_ = _dafny.Array(False, 22)
                    d_426_v112_ = nw73_
                    d_427_v113_: _dafny.Map
                    d_427_v113_ = _dafny.Map({(d_381_v72_).f17: -741})
                    d_428_v115_: _dafny.Map
                    def iife53_():
                        coll45_ = _dafny.Map()
                        compr_45_: int
                        for compr_45_ in _dafny.IntegerRange(-627, -696):
                            d_429_v114_: int = compr_45_
                            if ((-627) <= (d_429_v114_)) and ((d_429_v114_) < (-696)):
                                coll45_[default__.safeModuloInt(d_429_v114_, (d_380_v71_).f17)] = (d_380_v71_).f17
                        return _dafny.Map(coll45_)
                    d_428_v115_ = _dafny.Map({d_427_v113_: iife53_()
                    })
                    d_430_v116_: _dafny.Map
                    d_430_v116_ = _dafny.Map({d_426_v112_: d_428_v115_})
                    d_431_v117_: T1
                    nw74_ = C7()
                    nw74_.ctor__(d_425_v111_, (d_299_v2_) + (d_299_v2_), (d_380_v71_).f17, ((d_430_v116_)[d_426_v112_] if (d_426_v112_) in (d_430_v116_) else d_428_v115_))
                    d_431_v117_ = nw74_
                    d_431_v117_ = (d_431_v117_ if default__.fm13(d_301_v3_, (d_381_v72_).f17, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "akih")), d_300_globalState_) else d_431_v117_)
                    d_432_v118_: _dafny.Set
                    d_432_v118_ = _dafny.Set({d_427_v113_})
                    d_433_v119_: str
                    d_433_v119_ = _dafny.CodePoint('i')
                    d_434_v120_: _dafny.Set
                    d_434_v120_ = _dafny.Set({d_433_v119_})
                    rhs46_ = _dafny.Set({(default__.fm14(d_433_v119_, d_301_v3_, d_300_globalState_)) | (d_427_v113_), d_427_v113_})
                    rhs47_ = d_434_v120_
                    rhs48_ = d_431_v117_
                    rhs49_ = (0) - ((d_380_v71_).f17)
                    rhs50_ = not(((d_380_v71_).f17) >= (len(d_425_v111_)))
                    lhs40_ = d_300_globalState_
                    d_432_v118_ = rhs46_
                    d_434_v120_ = rhs47_
                    d_431_v117_ = rhs48_
                    lhs40_.f3 = rhs49_
                    d_301_v3_ = rhs50_
                    d_425_v111_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mc"))).set(default__.safeIndex(((d_380_v71_).fm23(d_300_globalState_) if d_301_v3_ else (d_380_v71_).f17), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mc")))), _dafny.CodePoint('l'))
                    pass
            pass
        d_435_v121_: _dafny.Array
        nw75_ = _dafny.Array(int(0), 28)
        d_435_v121_ = nw75_
        d_436_v122_: _dafny.Seq
        d_436_v122_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psiuxumem"))
        d_437_v124_: _dafny.Map
        d_437_v124_ = _dafny.Map({(d_380_v71_).f17: (d_381_v72_).f17})
        d_438_v125_: _dafny.Map
        d_438_v125_ = _dafny.Map({(d_380_v71_).f17: d_301_v3_})
        d_439_v126_: _dafny.Seq
        d_439_v126_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_380_v71_).f17: d_297_v0_}), d_437_v124_, d_437_v124_, d_437_v124_, _dafny.Map({len(d_438_v125_): len(d_438_v125_)})])
        d_440_v127_: T1
        nw76_ = C7()
        def iife54_():
            coll46_ = _dafny.Map()
            compr_46_: _dafny.Map
            for compr_46_ in (d_439_v126_).Elements:
                d_441_v123_: _dafny.Map = compr_46_
                if (d_441_v123_) in (d_439_v126_):
                    coll46_[d_441_v123_] = d_437_v124_
            return _dafny.Map(coll46_)
        nw76_.ctor__(d_436_v122_, d_299_v2_, (d_380_v71_).f17, iife54_()
        )
        d_440_v127_ = nw76_
        d_442_v128_: _dafny.Map
        d_442_v128_ = _dafny.Map({d_301_v3_: d_440_v127_})
        d_443_v129_: _dafny.Map
        d_443_v129_ = _dafny.Map({d_435_v121_: len(d_442_v128_)})
        d_444_v131_: str
        d_444_v131_ = _dafny.CodePoint('v')
        d_445_v132_: _dafny.MultiSet
        d_445_v132_ = _dafny.MultiSet([d_444_v131_])
        d_446_v133_: _dafny.Map
        def iife55_():
            coll47_ = _dafny.Map()
            compr_47_: str
            for compr_47_ in (d_445_v132_).Elements:
                d_447_v130_: str = compr_47_
                if (d_447_v130_) in (d_445_v132_):
                    coll47_[d_447_v130_] = d_301_v3_
            return _dafny.Map(coll47_)
        d_446_v133_ = _dafny.Map({iife55_()
        : (d_440_v127_).f11})
        d_448_v134_: _dafny.Map
        d_448_v134_ = _dafny.Map({d_444_v131_: not(False)})
        d_443_v129_ = (d_443_v129_).set(d_435_v121_, (default__.fm10(not(d_301_v3_), (d_381_v72_).f17, (d_380_v71_).f17, ((d_446_v133_)[d_448_v134_] if (d_448_v134_) in (d_446_v133_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "si")))), d_300_globalState_)).cardinality)
        d_449_v135_: _dafny.MultiSet
        d_449_v135_ = _dafny.MultiSet([d_301_v3_])
        d_450_v136_: _dafny.Map
        d_450_v136_ = _dafny.Map({default__.fm76(d_301_v3_, d_301_v3_, d_297_v0_, d_300_globalState_): d_449_v135_})
        d_451_v137_: _dafny.Map
        d_451_v137_ = _dafny.Map({not(not(d_301_v3_)): d_436_v122_})
        if (((_dafny.MultiSet([not(d_301_v3_), d_301_v3_, d_301_v3_])) | (((d_450_v136_)[d_451_v137_] if (d_451_v137_) in (d_450_v136_) else d_449_v135_))).cardinality) > (default__.safeDivisionInt(d_297_v0_, len(d_436_v122_))):
            d_435_v121_ = d_435_v121_
            d_452_v138_: str
            d_453_v139_: _dafny.Seq
            d_454_v140_: D7
            d_455_v141_: str
            out6_: str
            out7_: _dafny.Seq
            out8_: D7
            out9_: str
            out6_, out7_, out8_, out9_ = (d_380_v71_).m12(d_300_globalState_)
            d_452_v138_ = out6_
            d_453_v139_ = out7_
            d_454_v140_ = out8_
            d_455_v141_ = out9_
            d_456_v142_: _dafny.Map
            d_456_v142_ = _dafny.Map({d_301_v3_: d_301_v3_})
            d_456_v142_ = (_dafny.Map({not(d_301_v3_): d_301_v3_})) | ((d_456_v142_) | ((d_456_v142_).set(d_301_v3_, d_301_v3_)))
            d_457_v143_: _dafny.Seq
            d_457_v143_ = _dafny.SeqWithoutIsStrInference([(d_381_v72_).f17])
            d_458_v144_: _dafny.MultiSet
            d_458_v144_ = _dafny.MultiSet([d_457_v143_])
            d_459_v145_: _dafny.Seq
            d_459_v145_ = _dafny.SeqWithoutIsStrInference([default__.fm22(d_301_v3_, len(d_436_v122_), d_300_globalState_)])
            d_460_v146_: _dafny.MultiSet
            d_460_v146_ = _dafny.MultiSet([d_459_v145_])
            rhs51_ = d_297_v0_
            rhs52_ = (((d_440_v127_).f9) + (d_436_v122_)) + (((d_440_v127_).f9) + (d_453_v139_))
            rhs53_ = default__.fm77(len(default__.fm44(True, d_297_v0_, d_301_v3_, d_300_globalState_)), (d_440_v127_).f11, (d_460_v146_).cardinality, d_453_v139_, d_300_globalState_)
            rhs54_ = default__.fm37((d_440_v127_).f11, (d_381_v72_).f17, d_301_v3_, (d_380_v71_).f17, d_300_globalState_)
            rhs55_ = not (d_301_v3_) or (d_301_v3_)
            lhs41_ = d_300_globalState_
            lhs41_.f3 = rhs51_
            d_453_v139_ = rhs52_
            d_458_v144_ = rhs53_
            d_436_v122_ = rhs54_
            d_301_v3_ = rhs55_
            d_461_v147_: _dafny.Array
            def lambda10_(d_462_v139_, d_463_v127_):
                def lambda11_(d_464_i13_):
                    return (d_462_v139_) + ((d_463_v127_).f9)

                return lambda11_

            init5_ = lambda10_(d_453_v139_, d_440_v127_)
            nw77_ = _dafny.Array(None, 11)
            for i0_5_ in range(nw77_.length(0)):
                nw77_[i0_5_] = init5_(i0_5_)
            d_461_v147_ = nw77_
            index51_ = default__.safeIndex(770, (d_461_v147_).length(0))
            (d_461_v147_)[index51_] = ((_dafny.SeqWithoutIsStrInference([d_452_v138_ for d_465_i14_ in range(default__.abs(-477))])) + (d_453_v139_)).set(default__.safeIndex(d_297_v0_, len((_dafny.SeqWithoutIsStrInference([d_452_v138_ for d_466_i14_ in range(default__.abs(-477))])) + (d_453_v139_))), d_452_v138_)
            index52_ = default__.safeIndex(770, (d_461_v147_).length(0))
            (d_461_v147_)[index52_] = d_453_v139_
        elif True:
            if d_301_v3_:
                d_467_v148_: C1
                nw78_ = C1()
                nw78_.ctor__((d_440_v127_).f11, (0) - ((d_381_v72_).f17), (d_381_v72_).f17, d_440_v127_.f12, d_436_v122_, (d_440_v127_).f10)
                d_467_v148_ = nw78_
                d_468_v149_: _dafny.Map
                d_468_v149_ = _dafny.Map({d_467_v148_: d_301_v3_})
                d_469_v150_: _dafny.Seq
                d_469_v150_ = _dafny.SeqWithoutIsStrInference([((d_468_v149_)[d_467_v148_] if (d_467_v148_) in (d_468_v149_) else d_301_v3_)])
                d_301_v3_ = not ((d_469_v150_)[default__.safeIndex(d_297_v0_, len(d_469_v150_))]) or (d_301_v3_)
                d_470_v151_: _dafny.Seq
                d_470_v151_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_301_v3_]))).cardinality, (d_440_v127_).f11])
                d_471_v152_: C1
                nw79_ = C1()
                nw79_.ctor__(d_467_v148_.f23, 797, (0) - ((d_440_v127_).f11), (default__.fm47(d_301_v3_, d_301_v3_, d_301_v3_, d_301_v3_, d_300_globalState_)) | (d_440_v127_.f12), (d_440_v127_).f9, _dafny.SeqWithoutIsStrInference([d_470_v151_, d_470_v151_]))
                d_471_v152_ = nw79_
                d_472_v153_: _dafny.Map
                d_472_v153_ = _dafny.Map({d_301_v3_: ((_dafny.MultiSet(d_469_v150_)).cardinality) != ((d_467_v148_).f24)})
                d_472_v153_ = d_472_v153_
                d_473_v154_: str
                d_474_v155_: _dafny.Seq
                d_475_v156_: D7
                d_476_v157_: str
                out10_: str
                out11_: _dafny.Seq
                out12_: D7
                out13_: str
                out10_, out11_, out12_, out13_ = (d_380_v71_).m12(d_300_globalState_)
                d_473_v154_ = out10_
                d_474_v155_ = out11_
                d_475_v156_ = out12_
                d_476_v157_ = out13_
                index53_ = default__.safeIndex(470, (d_435_v121_).length(0))
                (d_435_v121_)[index53_] = (d_381_v72_).f17
                index54_ = default__.safeIndex(470, (d_435_v121_).length(0))
                (d_435_v121_)[index54_] = ((d_440_v127_).f11 if d_301_v3_ else 23)
            elif True:
                d_301_v3_ = (default__.safeDivisionInt((d_381_v72_).f17, (d_440_v127_).f11)) < (((d_380_v71_).f17) + ((0) - ((d_381_v72_).f17)))
                d_477_v160_: C7
                nw80_ = C7()
                def iife56_():
                    def iife58_():
                        coll50_ = _dafny.Map()
                        compr_50_: _dafny.Map
                        for compr_50_ in (d_439_v126_).Elements:
                            d_478_v159_: _dafny.Map = compr_50_
                            if (d_478_v159_) in (d_439_v126_):
                                coll50_[d_478_v159_] = d_301_v3_
                        return _dafny.Map(coll50_)
                    coll48_ = _dafny.Map()
                    def iife57_():
                        coll49_ = _dafny.Map()
                        compr_49_: _dafny.Map
                        for compr_49_ in (d_439_v126_).Elements:
                            d_478_v159_: _dafny.Map = compr_49_
                            if (d_478_v159_) in (d_439_v126_):
                                coll49_[d_478_v159_] = d_301_v3_
                        return _dafny.Map(coll49_)
                    compr_48_: _dafny.Map
                    for compr_48_ in (iife57_()
                    ).keys.Elements:
                        d_479_v158_: _dafny.Map = compr_48_
                        if (d_479_v158_) in (iife58_()
                        ):
                            coll48_[d_479_v158_] = d_437_v124_
                    return _dafny.Map(coll48_)
                nw80_.ctor__((d_440_v127_).f9, (d_440_v127_).f10, len(((d_440_v127_).f9) + ((d_440_v127_).f9)), (iife56_()
                ) | (d_440_v127_.f12))
                d_477_v160_ = nw80_
                d_480_v161_: _dafny.Seq
                d_480_v161_ = _dafny.SeqWithoutIsStrInference([793])
                d_481_v162_: D12
                d_481_v162_ = D12_DC28(d_480_v161_)
                d_482_v163_: _dafny.Map
                d_482_v163_ = _dafny.Map({d_481_v162_: d_301_v3_})
                d_483_v164_: C8
                nw81_ = C8()
                nw81_.ctor__(default__.safeModuloInt((d_440_v127_).f11, (d_449_v135_).cardinality), ((d_482_v163_)[d_481_v162_] if (d_481_v162_) in (d_482_v163_) else d_301_v3_))
                d_483_v164_ = nw81_
                d_483_v164_ = d_483_v164_
                d_484_v165_: C3
                nw82_ = C3()
                nw82_.ctor__(d_301_v3_, (d_440_v127_).f11, (d_380_v71_).f17, d_440_v127_.f12, (d_440_v127_).f9, (d_440_v127_).f10)
                d_484_v165_ = nw82_
                d_485_v166_: _dafny.Map
                d_485_v166_ = _dafny.Map({d_484_v165_: d_301_v3_})
                d_486_v167_: _dafny.Seq
                d_486_v167_ = _dafny.SeqWithoutIsStrInference([d_483_v164_.f16, (d_298_v1_).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_380_v71_).f17, (d_483_v164_).f15, len(d_485_v166_)]))), (d_484_v165_).f21])
                d_486_v167_ = (d_486_v167_).set(default__.safeIndex((default__.fm0((d_380_v71_).f17, (d_380_v71_).f17, d_300_globalState_)) * ((d_381_v72_).f17), len(d_486_v167_)), default__.fm13(d_301_v3_, (d_440_v127_).f11, (d_436_v122_).set(default__.safeIndex((d_381_v72_).f17, len(d_436_v122_)), d_444_v131_), d_300_globalState_))
                d_487_v168_: _dafny.Array
                def lambda12_(d_488_i15_):
                    return True

                init6_ = lambda12_
                nw83_ = _dafny.Array(None, 2)
                for i0_6_ in range(nw83_.length(0)):
                    nw83_[i0_6_] = init6_(i0_6_)
                d_487_v168_ = nw83_
                index55_ = default__.safeIndex(720, (d_487_v168_).length(0))
                (d_487_v168_)[index55_] = (d_484_v165_).f21
                d_489_v169_: _dafny.Seq
                d_489_v169_ = _dafny.SeqWithoutIsStrInference([d_487_v168_, d_487_v168_, d_487_v168_, d_487_v168_, d_487_v168_])
                index56_ = default__.safeIndex(720, (d_487_v168_).length(0))
                (d_487_v168_)[index56_] = ((d_438_v125_)[len((d_489_v169_) + ((d_489_v169_).set(default__.safeIndex((d_440_v127_).f11, len(d_489_v169_)), d_487_v168_)))] if (len((d_489_v169_) + ((d_489_v169_).set(default__.safeIndex((d_440_v127_).f11, len(d_489_v169_)), d_487_v168_)))) in (d_438_v125_) else d_301_v3_)
            d_490_v170_: _dafny.Map
            d_490_v170_ = d_440_v127_.f12
            source7_ = d_490_v170_
            d_491___mcc_h0_ = source7_
            d_492_cf108_ = d_491___mcc_h0_
            d_493_v172_: T0
            nw84_ = C7()
            def iife59_():
                coll51_ = _dafny.Map()
                compr_51_: int
                for compr_51_ in _dafny.IntegerRange(323, 16):
                    d_494_v171_: int = compr_51_
                    if ((323) <= (d_494_v171_)) and ((d_494_v171_) < (16)):
                        coll51_[(d_494_v171_) - ((d_381_v72_).f17)] = d_297_v0_
                return _dafny.Map(coll51_)
            nw84_.ctor__(default__.fm31(d_300_globalState_), (d_440_v127_).f10, (d_440_v127_).f11, (_dafny.Map({d_437_v124_: d_437_v124_})) | (_dafny.Map({d_437_v124_: iife59_()
            })))
            d_493_v172_ = nw84_
            d_495_v173_: C1
            nw85_ = C1()
            nw85_.ctor__(len((d_440_v127_).f9), (d_440_v127_).f11, len(default__.fm2((d_440_v127_).f9, d_297_v0_, d_300_globalState_)), d_440_v127_.f12, (d_440_v127_).f9, (d_493_v172_).f10)
            d_495_v173_ = nw85_
            d_496_v174_: _dafny.Map
            d_496_v174_ = _dafny.Map({True: d_495_v173_})
            d_497_v175_: _dafny.Set
            d_497_v175_ = _dafny.Set({d_495_v173_, ((d_496_v174_)[d_301_v3_] if (d_301_v3_) in (d_496_v174_) else d_495_v173_), d_495_v173_, d_495_v173_, d_495_v173_})
            d_498_v176_: _dafny.Map
            d_498_v176_ = _dafny.Map({d_497_v175_: d_495_v173_})
            d_499_v177_: _dafny.Array
            nw86_ = _dafny.Array(None, 2)
            nw86_[int(0)] = (d_301_v3_) == (d_301_v3_)
            nw86_[int(1)] = d_301_v3_
            d_499_v177_ = nw86_
            index57_ = default__.safeIndex(94, (d_499_v177_).length(0))
            (d_499_v177_)[index57_] = not (d_301_v3_) or (True)
            index58_ = default__.safeIndex(94, (d_499_v177_).length(0))
            rhs56_ = d_493_v172_
            rhs57_ = ((d_498_v176_) | (d_498_v176_)) | (d_498_v176_)
            rhs58_ = (d_301_v3_) in (d_382_v73_)
            lhs42_ = d_499_v177_
            lhs43_ = default__.safeIndex(94, (d_499_v177_).length(0))
            d_493_v172_ = rhs56_
            d_498_v176_ = rhs57_
            lhs42_[lhs43_] = rhs58_
            d_500_v178_: _dafny.Seq
            d_500_v178_ = _dafny.SeqWithoutIsStrInference([d_440_v127_])
            d_501_v179_: _dafny.Seq
            d_501_v179_ = _dafny.SeqWithoutIsStrInference([(d_500_v178_).set(default__.safeIndex((d_381_v72_).f17, len(d_500_v178_)), d_440_v127_)])
            d_502_v180_: C1
            nw87_ = C1()
            nw87_.ctor__(len(d_501_v179_), (d_380_v71_).f17, default__.safeDivisionInt(d_297_v0_, (d_380_v71_).f17), d_492_cf108_, (d_440_v127_).f9, (d_440_v127_).f10)
            d_502_v180_ = nw87_
            d_503_v181_: D2
            d_503_v181_ = D2_DC5((d_499_v177_)[default__.safeIndex(94, (d_499_v177_).length(0))], d_499_v177_, (d_495_v173_.f23) - ((d_502_v180_).f24), 153)
            d_503_v181_ = d_503_v181_
            d_504_v182_: _dafny.Array
            def lambda13_(d_505_v127_, d_506_v71_):
                def lambda14_(d_507_i16_):
                    return _dafny.Map({(d_505_v127_).f11: (d_506_v71_).f17})

                return lambda14_

            init7_ = lambda13_(d_440_v127_, d_380_v71_)
            nw88_ = _dafny.Array(None, 3)
            for i0_7_ in range(nw88_.length(0)):
                nw88_[i0_7_] = init7_(i0_7_)
            d_504_v182_ = nw88_
            d_504_v182_ = d_504_v182_
            d_444_v131_ = ((d_440_v127_).f9)[default__.safeIndex(650, len((d_440_v127_).f9))]
            d_508_v183_: _dafny.Array
            def lambda15_(d_509_v3_):
                def lambda16_(d_510_i17_):
                    return d_509_v3_

                return lambda16_

            init8_ = lambda15_(d_301_v3_)
            nw89_ = _dafny.Array(None, 25)
            for i0_8_ in range(nw89_.length(0)):
                nw89_[i0_8_] = init8_(i0_8_)
            d_508_v183_ = nw89_
            d_511_v184_: C9
            nw90_ = C9()
            nw90_.ctor__(d_508_v183_)
            d_511_v184_ = nw90_
            d_301_v3_ = (d_301_v3_ if d_301_v3_ else default__.fm13(d_301_v3_, (d_381_v72_).f17, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvlucshy")), d_300_globalState_))
        hi0_ = d_297_v0_
        for d_512_i18_ in range((d_380_v71_).f17, hi0_):
            d_513_v185_: _dafny.Map
            d_513_v185_ = _dafny.Map({default__.fm13(d_301_v3_, (d_381_v72_).f17, _dafny.SeqWithoutIsStrInference([d_444_v131_ for d_514_i19_ in range(default__.abs(376))]), d_300_globalState_): (d_440_v127_).f11})
            d_513_v185_ = (d_513_v185_).set(d_301_v3_, (d_381_v72_).f17)
            d_297_v0_ = ((d_381_v72_).f17) * ((d_381_v72_).f17)
            (d_300_globalState_).f3 = default__.safeDivisionInt(-49, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_515_i20_ in range(default__.abs(619))])))
            d_516_v186_: _dafny.Array
            nw91_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_516_v186_ = nw91_
            index59_ = default__.safeIndex(188, (d_516_v186_).length(0))
            (d_516_v186_)[index59_] = d_435_v121_
            index60_ = default__.safeIndex(188, (d_516_v186_).length(0))
            (d_516_v186_)[index60_] = d_435_v121_
        d_517_v187_: _dafny.Array
        nw92_ = _dafny.Array(False, 13)
        d_517_v187_ = nw92_
        d_517_v187_ = d_517_v187_
        if d_301_v3_:
            d_301_v3_ = (d_301_v3_) or (True)
            d_301_v3_ = d_301_v3_
            d_518_v188_: _dafny.Seq
            d_518_v188_ = (d_440_v127_).f9
            d_519_v189_: C0
            nw93_ = C0()
            nw93_.ctor__(d_301_v3_, default__.fm22(d_301_v3_, d_297_v0_, d_300_globalState_), (d_440_v127_).f9, (d_440_v127_).f10)
            d_519_v189_ = nw93_
            d_520_v190_: _dafny.Map
            d_520_v190_ = _dafny.Map({d_518_v188_: d_519_v189_})
            d_520_v190_ = d_520_v190_
            d_297_v0_ = d_297_v0_
            d_436_v122_ = d_436_v122_
        elif True:
            d_521_v191_: _dafny.Seq
            d_521_v191_ = _dafny.SeqWithoutIsStrInference([d_301_v3_])
            d_522_v192_: _dafny.Map
            d_522_v192_ = _dafny.Map({(d_381_v72_).f17: d_521_v191_})
            d_523_v193_: _dafny.Map
            d_523_v193_ = _dafny.Map({d_301_v3_: d_301_v3_})
            d_524_v194_: _dafny.Seq
            d_524_v194_ = _dafny.SeqWithoutIsStrInference([d_523_v193_, d_523_v193_, d_523_v193_, d_523_v193_])
            d_525_v195_: _dafny.Array
            nw94_ = _dafny.Array(None, 25)
            nw94_[int(0)] = ((d_522_v192_)[(0) - ((d_381_v72_).f17)] if ((0) - ((d_381_v72_).f17)) in (d_522_v192_) else _dafny.SeqWithoutIsStrInference([not(d_301_v3_), d_301_v3_]))
            nw94_[int(1)] = (d_521_v191_) + (d_521_v191_)
            nw94_[int(2)] = (_dafny.SeqWithoutIsStrInference([d_301_v3_])) + (d_521_v191_)
            nw94_[int(3)] = d_521_v191_
            nw94_[int(4)] = (d_521_v191_) + ((d_521_v191_).set(default__.safeIndex(len((d_524_v194_)[default__.safeIndex((d_380_v71_).f17, len(d_524_v194_))]), len(d_521_v191_)), d_301_v3_))
            nw94_[int(5)] = (d_521_v191_) + (_dafny.SeqWithoutIsStrInference([d_301_v3_, True]))
            nw94_[int(6)] = (_dafny.SeqWithoutIsStrInference([d_301_v3_])) + (d_521_v191_)
            nw94_[int(7)] = d_521_v191_
            nw94_[int(8)] = (d_521_v191_) + (d_521_v191_)
            nw94_[int(9)] = d_521_v191_
            nw94_[int(10)] = d_521_v191_
            nw94_[int(11)] = d_521_v191_
            nw94_[int(12)] = d_521_v191_
            nw94_[int(13)] = d_521_v191_
            nw94_[int(14)] = d_521_v191_
            nw94_[int(15)] = d_521_v191_
            nw94_[int(16)] = d_521_v191_
            nw94_[int(17)] = (d_521_v191_) + (d_521_v191_)
            nw94_[int(18)] = ((d_521_v191_) + (default__.fm2((d_440_v127_).f9, (d_440_v127_).f11, d_300_globalState_))).set(default__.safeIndex((d_380_v71_).f17, len((d_521_v191_) + (default__.fm2((d_440_v127_).f9, (d_440_v127_).f11, d_300_globalState_)))), not(default__.fm22(d_301_v3_, (d_381_v72_).f17, d_300_globalState_)))
            nw94_[int(19)] = (_dafny.SeqWithoutIsStrInference([d_301_v3_])) + (d_521_v191_)
            nw94_[int(20)] = d_521_v191_
            nw94_[int(21)] = d_521_v191_
            nw94_[int(22)] = d_521_v191_
            nw94_[int(23)] = _dafny.SeqWithoutIsStrInference([d_301_v3_])
            nw94_[int(24)] = d_521_v191_
            d_525_v195_ = nw94_
            index61_ = default__.safeIndex(33, (d_525_v195_).length(0))
            (d_525_v195_)[index61_] = _dafny.SeqWithoutIsStrInference([d_301_v3_])
            index62_ = default__.safeIndex(33, (d_525_v195_).length(0))
            (d_525_v195_)[index62_] = d_521_v191_
            d_301_v3_ = False
            d_526_v196_: str
            d_527_v197_: _dafny.Seq
            d_528_v198_: D7
            d_529_v199_: str
            out14_: str
            out15_: _dafny.Seq
            out16_: D7
            out17_: str
            out14_, out15_, out16_, out17_ = (d_381_v72_).m12(d_300_globalState_)
            d_526_v196_ = out14_
            d_527_v197_ = out15_
            d_528_v198_ = out16_
            d_529_v199_ = out17_
            d_530_v200_: _dafny.Array
            nw95_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_530_v200_ = nw95_
            d_531_v201_: _dafny.Map
            d_531_v201_ = _dafny.Map({d_517_v187_: d_530_v200_})
            d_531_v201_ = (d_531_v201_).set(d_517_v187_, (d_530_v200_ if d_301_v3_ else d_530_v200_))
            d_532_v202_: _dafny.Array
            def lambda17_(d_533_v127_, d_534_v71_, d_535_v122_, d_536_v196_):
                def lambda18_(d_537_i21_):
                    return (_dafny.SeqWithoutIsStrInference([((d_533_v127_).f9)[default__.safeIndex((d_534_v71_).f17, len((d_533_v127_).f9))] for d_538_i22_ in range(default__.abs(661))])) + ((d_535_v122_).set(default__.safeIndex((d_534_v71_).f17, len(d_535_v122_)), d_536_v196_))

                return lambda18_

            init9_ = lambda17_(d_440_v127_, d_380_v71_, d_436_v122_, d_526_v196_)
            nw96_ = _dafny.Array(None, 4)
            for i0_9_ in range(nw96_.length(0)):
                nw96_[i0_9_] = init9_(i0_9_)
            d_532_v202_ = nw96_
            nw97_ = _dafny.Array(None, 3)
            nw97_[int(0)] = d_436_v122_
            nw97_[int(1)] = ((d_440_v127_).f9) + (d_527_v197_)
            nw97_[int(2)] = default__.fm37(222, len(d_436_v122_), d_301_v3_, (d_381_v72_).f17, d_300_globalState_)
            d_532_v202_ = nw97_
        d_539_v203_: _dafny.Array
        nw98_ = _dafny.Array(_dafny.Array(None, 0), 15)
        d_539_v203_ = nw98_
        index63_ = default__.safeIndex(471, (d_539_v203_).length(0))
        (d_539_v203_)[index63_] = d_435_v121_
        d_540_v204_: _dafny.Seq
        d_540_v204_ = _dafny.SeqWithoutIsStrInference([d_435_v121_, d_435_v121_, d_435_v121_])
        index64_ = default__.safeIndex(471, (d_539_v203_).length(0))
        (d_539_v203_)[index64_] = (d_540_v204_)[default__.safeIndex((d_380_v71_).f17, len(d_540_v204_))]
        d_541_v205_: _dafny.Map
        d_541_v205_ = _dafny.Map({d_301_v3_: d_301_v3_})
        d_542_v206_: _dafny.Array
        d_543_v207_: _dafny.Array
        out18_: _dafny.Array
        out19_: _dafny.Array
        out18_, out19_ = (d_440_v127_).m3(default__.safeModuloInt((d_381_v72_).f17, (d_380_v71_).f17), default__.fm37((d_440_v127_).fm5((d_298_v1_).cardinality, (d_380_v71_).f17, d_300_globalState_), (d_380_v71_).f17, False, len(d_541_v205_), d_300_globalState_), (d_297_v0_) - (len(_dafny.Map({d_297_v0_: d_301_v3_}))), (d_436_v122_).set(default__.safeIndex((d_380_v71_).f17, len(d_436_v122_)), default__.fm17(d_301_v3_, d_300_globalState_)), d_300_globalState_)
        d_542_v206_ = out18_
        d_543_v207_ = out19_
        d_544_v208_: _dafny.Map
        d_544_v208_ = _dafny.Map({d_301_v3_: (d_380_v71_).f17})
        d_545_v209_: _dafny.Seq
        d_546_v210_: D1
        d_547_v211_: bool
        out20_: _dafny.Seq
        out21_: D1
        out22_: bool
        out20_, out21_, out22_ = (d_380_v71_).m11((d_544_v208_ if not(d_301_v3_) else d_544_v208_), d_517_v187_, ((0) - ((d_380_v71_).f17)) + ((d_381_v72_).f17), d_300_globalState_)
        d_545_v209_ = out20_
        d_546_v210_ = out21_
        d_547_v211_ = out22_
        d_548_v212_: C1
        nw99_ = C1()
        nw99_.ctor__((d_380_v71_).f17, (d_380_v71_).f17, (d_380_v71_).f17, d_440_v127_.f12, (d_440_v127_).f9, d_299_v2_)
        d_548_v212_ = nw99_
        d_549_v213_: _dafny.Array
        nw100_ = _dafny.Array(_dafny.CodePoint('D'), 1)
        d_549_v213_ = nw100_
        index65_ = default__.safeIndex(569, (d_549_v213_).length(0))
        (d_549_v213_)[index65_] = d_444_v131_
        index66_ = default__.safeIndex(569, (d_549_v213_).length(0))
        rhs59_ = (d_548_v212_ if d_547_v211_ else d_548_v212_)
        rhs60_ = d_444_v131_
        lhs44_ = d_549_v213_
        lhs45_ = default__.safeIndex(569, (d_549_v213_).length(0))
        d_548_v212_ = rhs59_
        lhs44_[lhs45_] = rhs60_
        d_550_v214_: D4
        d_550_v214_ = D4_DC10(d_438_v125_)
        pat_let_tv7_ = d_381_v72_
        pat_let_tv8_ = d_437_v124_
        pat_let_tv9_ = d_440_v127_
        pat_let_tv10_ = d_440_v127_
        def lambda19_(source9_):
            if source9_.is_DC9:
                d_551___mcc_h6_ = source9_.cf11
                d_552_cf11_ = d_551___mcc_h6_
                return D2_DC4((pat_let_tv7_).f17)
            elif source9_.is_DC10:
                d_553___mcc_h7_ = source9_.cf12
                d_554_cf12_ = d_553___mcc_h7_
                return D2_DC4(len(pat_let_tv8_))
            elif source9_.is_DC11:
                d_555___mcc_h8_ = source9_.cf13
                d_556___mcc_h9_ = source9_.cf14
                d_557_cf14_ = d_556___mcc_h9_
                d_558_cf13_ = d_555___mcc_h8_
                return D2_DC4((pat_let_tv9_).f11)
            elif True:
                d_559___mcc_h10_ = source9_.cf10
                d_560_cf10_ = d_559___mcc_h10_
                return D2_DC4((pat_let_tv10_).f11)

        source8_ = lambda19_(d_550_v214_)
        if source8_.is_DC5:
            d_561___mcc_h1_ = source8_.cf4
            d_562___mcc_h2_ = source8_.cf5
            d_563___mcc_h3_ = source8_.cf6
            d_564___mcc_h4_ = source8_.cf7
            d_565_cf7_ = d_564___mcc_h4_
            d_566_cf6_ = d_563___mcc_h3_
            d_567_cf5_ = d_562___mcc_h2_
            d_568_cf4_ = d_561___mcc_h1_
            d_569_v215_: D17
            d_569_v215_ = D17_DC40((d_440_v127_).f9, d_301_v3_, (d_381_v72_).f17)
            d_297_v0_ = (d_569_v215_).cf56
            d_570_v216_: D11
            d_570_v216_ = D11_DC27(D11_DC26())
            source10_ = d_570_v216_
            if source10_.is_DC26:
                d_565_cf7_ = (d_440_v127_).f11
                d_568_cf4_ = (_dafny.CodePoint('m')) not in ((_dafny.SeqWithoutIsStrInference([(d_549_v213_)[default__.safeIndex(569, (d_549_v213_).length(0))] for d_571_i23_ in range(default__.abs(804))])) + ((d_440_v127_).f9))
                d_572_v217_: _dafny.Set
                d_572_v217_ = _dafny.Set({d_297_v0_, (0) - ((d_440_v127_).f11), 922})
                (d_548_v212_).f23 = (default__.safeModuloInt(d_566_cf6_, -111)) - (len(d_572_v217_))
                d_568_cf4_ = d_568_cf4_
            elif source10_.is_DC25:
                d_573___mcc_h11_ = source10_.cf29
                d_574_cf29_ = d_573___mcc_h11_
                (d_548_v212_).f23 = 908
                d_575_v218_: _dafny.Map
                d_575_v218_ = _dafny.Map({d_566_cf6_: default__.fm31(d_300_globalState_)})
                d_576_v219_: _dafny.Array
                d_577_v220_: _dafny.Array
                out23_: _dafny.Array
                out24_: _dafny.Array
                out23_, out24_ = (d_548_v212_).m3((d_380_v71_).f17, default__.fm26((d_381_v72_).f17, d_575_v218_, d_300_globalState_), d_548_v212_.f23, (d_440_v127_).f9, d_300_globalState_)
                d_576_v219_ = out23_
                d_577_v220_ = out24_
                d_578_v221_: C3
                nw101_ = C3()
                nw101_.ctor__(d_568_cf4_, d_565_cf7_, d_548_v212_.f23, d_440_v127_.f12, (d_436_v122_).set(default__.safeIndex((d_381_v72_).f17, len(d_436_v122_)), (d_549_v213_)[default__.safeIndex(569, (d_549_v213_).length(0))]), ((d_440_v127_).f10) + ((d_440_v127_).f10))
                d_578_v221_ = nw101_
                d_578_v221_ = d_578_v221_
                (d_548_v212_).f23 = (d_548_v212_).f24
            elif True:
                d_579___mcc_h12_ = source10_.cf30
                d_580_cf30_ = d_579___mcc_h12_
                d_581_v222_: bool
                d_582_v223_: _dafny.Array
                d_583_v224_: _dafny.Seq
                out25_: bool
                out26_: _dafny.Array
                out27_: _dafny.Seq
                out25_, out26_, out27_ = default__.m19(((0) - (d_548_v212_.f23) if False else d_297_v0_), d_444_v131_, d_300_globalState_)
                d_581_v222_ = out25_
                d_582_v223_ = out26_
                d_583_v224_ = out27_
                d_584_v225_: T0
                nw102_ = C0()
                nw102_.ctor__(d_568_cf4_, default__.fm13(((d_541_v205_)[d_581_v222_] if (d_581_v222_) in (d_541_v205_) else d_547_v211_), len((d_440_v127_).f9), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xay")), d_300_globalState_), (d_440_v127_).f9, (d_440_v127_).f10)
                d_584_v225_ = nw102_
                d_585_v226_: _dafny.Map
                d_585_v226_ = _dafny.Map({d_548_v212_.f23: d_584_v225_})
                d_586_v227_: _dafny.Seq
                d_587_v228_: D1
                d_588_v229_: bool
                out28_: _dafny.Seq
                out29_: D1
                out30_: bool
                out28_, out29_, out30_ = (d_380_v71_).m11(d_544_v208_, d_517_v187_, len(d_585_v226_), d_300_globalState_)
                d_586_v227_ = out28_
                d_587_v228_ = out29_
                d_588_v229_ = out30_
                d_298_v1_ = _dafny.MultiSet(default__.fm44(d_588_v229_, (0) - ((d_380_v71_).f17), d_301_v3_, d_300_globalState_))
                d_589_v230_: D18
                d_589_v230_ = D18_DC43(d_547_v211_, d_565_cf7_)
                d_590_v231_: bool
                d_591_v232_: bool
                d_592_v233_: _dafny.Array
                d_593_v234_: int
                out31_: bool
                out32_: bool
                out33_: _dafny.Array
                out34_: int
                out31_, out32_, out33_, out34_ = (d_548_v212_).m1((d_548_v212_).f24, (d_548_v212_).f24, len(_dafny.Set({d_589_v230_, d_589_v230_})), d_300_globalState_)
                d_590_v231_ = out31_
                d_591_v232_ = out32_
                d_592_v233_ = out33_
                d_593_v234_ = out34_
            index67_ = default__.safeIndex(191, (d_543_v207_).length(0))
            (d_543_v207_)[index67_] = default__.safeModuloInt((d_380_v71_).f17, (d_548_v212_).f24)
            index68_ = default__.safeIndex(191, (d_543_v207_).length(0))
            (d_543_v207_)[index68_] = (d_440_v127_).f11
            d_594_v235_: _dafny.Array
            nw103_ = _dafny.Array(_dafny.Seq({}), 6)
            d_594_v235_ = nw103_
            d_595_v236_: _dafny.Seq
            d_595_v236_ = _dafny.SeqWithoutIsStrInference([d_594_v235_, d_594_v235_])
            d_596_v237_: D18
            d_596_v237_ = D18_DC41(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False), d_547_v211_, d_301_v3_, d_301_v3_, d_547_v211_])))
            d_597_v238_: D22
            d_597_v238_ = D22_DC53((d_595_v236_)[default__.safeIndex(d_565_cf7_, len(d_595_v236_))], ((d_596_v237_).cf57).cardinality)
            source11_ = d_597_v238_
            if source11_.is_DC53:
                d_598___mcc_h13_ = source11_.cf78
                d_599___mcc_h14_ = source11_.cf79
                d_600_cf79_ = d_599___mcc_h14_
                d_601_cf78_ = d_598___mcc_h13_
                (d_300_globalState_).f3 = ((d_381_v72_).f17) - ((d_440_v127_).fm5(d_548_v212_.f23, (d_381_v72_).f17, d_300_globalState_))
                index69_ = default__.safeIndex(569, (d_549_v213_).length(0))
                (d_549_v213_)[index69_] = (d_444_v131_ if d_568_cf4_ else d_444_v131_)
                d_602_v239_: str
                d_603_v240_: _dafny.Seq
                d_604_v241_: D7
                d_605_v242_: str
                out35_: str
                out36_: _dafny.Seq
                out37_: D7
                out38_: str
                out35_, out36_, out37_, out38_ = (d_381_v72_).m12(d_300_globalState_)
                d_602_v239_ = out35_
                d_603_v240_ = out36_
                d_604_v241_ = out37_
                d_605_v242_ = out38_
                d_301_v3_ = d_301_v3_
            elif source11_.is_DC54:
                d_606___mcc_h15_ = source11_.cf80
                d_607___mcc_h16_ = source11_.cf81
                d_608_cf81_ = d_607___mcc_h16_
                d_609_cf80_ = d_606___mcc_h15_
                d_610_v243_: _dafny.Seq
                d_610_v243_ = _dafny.SeqWithoutIsStrInference([d_382_v73_, (d_382_v73_).intersection(d_382_v73_)])
                d_382_v73_ = (d_610_v243_)[default__.safeIndex(d_548_v212_.f23, len(d_610_v243_))]
                (d_300_globalState_).f3 = default__.safeDivisionInt(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gey"))) + ((d_440_v127_).f9)), d_608_cf81_)
                d_569_v215_ = d_569_v215_
                d_611_v244_: _dafny.Seq
                d_611_v244_ = _dafny.SeqWithoutIsStrInference([d_444_v131_ for d_612_i25_ in range(default__.abs(77))])
                d_613_v245_: _dafny.Seq
                d_613_v245_ = _dafny.SeqWithoutIsStrInference([(d_440_v127_).f9, ((d_440_v127_).f9).set(default__.safeIndex((d_548_v212_).f24, len((d_440_v127_).f9)), d_444_v131_), (d_611_v244_), (d_440_v127_).f9])
                d_614_v246_: _dafny.Seq
                d_614_v246_ = _dafny.SeqWithoutIsStrInference([-464])
                rhs61_ = _dafny.SeqWithoutIsStrInference([d_444_v131_ for d_615_i24_ in range(default__.abs(22))])
                rhs62_ = len((d_440_v127_).f9)
                rhs63_ = d_440_v127_
                rhs64_ = 263
                rhs65_ = ((d_613_v245_).set(default__.safeIndex(((d_449_v135_)[d_568_cf4_] if (d_568_cf4_) in (d_449_v135_) else (0) - (len(d_614_v246_))), len(d_613_v245_)), (d_440_v127_).f9)) <= (_dafny.SeqWithoutIsStrInference([(d_436_v122_).set(default__.safeIndex((d_381_v72_).f17, len(d_436_v122_)), d_444_v131_) for d_616_i26_ in range(default__.abs(959))]))
                lhs46_ = d_548_v212_
                d_436_v122_ = rhs61_
                d_609_cf80_ = rhs62_
                d_440_v127_ = rhs63_
                lhs46_.f23 = rhs64_
                d_547_v211_ = rhs65_
            elif True:
                d_617___mcc_h17_ = source11_.cf77
                d_618_cf77_ = d_617___mcc_h17_
                d_547_v211_ = d_568_cf4_
                d_619_v247_: C3
                nw104_ = C3()
                nw104_.ctor__(False, (d_440_v127_).f11, d_548_v212_.f23, d_440_v127_.f12, (d_440_v127_).f9, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_380_v71_).f17 for d_620_i28_ in range(default__.abs(182))]) for d_621_i27_ in range(default__.abs(674))]))
                d_619_v247_ = nw104_
                d_622_v248_: _dafny.Set
                d_622_v248_ = _dafny.Set({d_619_v247_})
                d_623_v249_: _dafny.Seq
                d_623_v249_ = _dafny.SeqWithoutIsStrInference([d_622_v248_])
                d_568_cf4_ = (d_622_v248_) not in (d_623_v249_)
                d_624_v250_: _dafny.Seq
                d_624_v250_ = _dafny.SeqWithoutIsStrInference([d_547_v211_])
                d_625_v251_: T0
                nw105_ = C7()
                nw105_.ctor__((d_440_v127_).f9, (d_440_v127_).f10, default__.safeModuloInt((d_440_v127_).f11, len(d_624_v250_)), d_440_v127_.f12)
                d_625_v251_ = nw105_
                d_626_v252_: C7
                nw106_ = C7()
                nw106_.ctor__((d_440_v127_).f9, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_624_v250_), d_566_cf6_]) for d_627_i29_ in range(default__.abs(551))]), d_297_v0_, d_440_v127_.f12)
                d_626_v252_ = nw106_
                d_628_v253_: D32
                d_628_v253_ = D32_DC72(d_625_v251_)
                rhs66_ = d_301_v3_
                rhs67_ = (d_628_v253_).cf111
                rhs68_ = d_440_v127_
                rhs69_ = d_626_v252_
                d_547_v211_ = rhs66_
                d_625_v251_ = rhs67_
                d_440_v127_ = rhs68_
                d_626_v252_ = rhs69_
                (d_548_v212_).f23 = (d_380_v71_).f17
        elif True:
            d_629___mcc_h5_ = source8_.cf3
            d_630_cf3_ = d_629___mcc_h5_
            d_382_v73_ = d_382_v73_
            index70_ = default__.safeIndex(31, (d_517_v187_).length(0))
            (d_517_v187_)[index70_] = d_301_v3_
            d_631_v254_: D17
            d_631_v254_ = D17_DC40(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddkrbmll")), d_547_v211_, len(d_437_v124_))
            d_632_v255_: _dafny.Map
            d_632_v255_ = _dafny.Map({d_298_v1_: ((d_449_v135_)[d_547_v211_] if (d_547_v211_) in (d_449_v135_) else (0) - ((d_440_v127_).f11))})
            index71_ = default__.safeIndex(31, (d_517_v187_).length(0))
            (d_517_v187_)[index71_] = (d_301_v3_ if not (d_547_v211_) or ((d_631_v254_).cf55) else (len(d_632_v255_)) > ((d_380_v71_).f17))
            d_633_v256_: bool
            d_634_v257_: bool
            d_635_v258_: _dafny.Array
            d_636_v259_: int
            out39_: bool
            out40_: bool
            out41_: _dafny.Array
            out42_: int
            out39_, out40_, out41_, out42_ = (d_548_v212_).m1(len(d_438_v125_), (d_381_v72_).f17, d_630_cf3_, d_300_globalState_)
            d_633_v256_ = out39_
            d_634_v257_ = out40_
            d_635_v258_ = out41_
            d_636_v259_ = out42_
            index72_ = default__.safeIndex(31, (d_517_v187_).length(0))
            (d_517_v187_)[index72_] = (d_517_v187_)[default__.safeIndex(31, (d_517_v187_).length(0))]
        _dafny.print(_dafny.string_of(d_297_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_298_v1_) == (_dafny.MultiSet([-71, -71]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_299_v2_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-71])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_globalState_).f0) == (_dafny.MultiSet([-71, -71]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_300_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_300_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_300_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_300_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_globalState_).f5) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-71])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_300_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_301_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_378_v69_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_380_v71_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_381_v72_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_382_v73_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_383_i9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_436_v122_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_437_v124_) == (_dafny.Map({-474: 4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_438_v125_) == (_dafny.Map({-474: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_439_v126_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({-474: -474}), _dafny.Map({-474: 4}), _dafny.Map({-474: 4}), _dafny.Map({-474: 4}), _dafny.Map({1: 1})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_440_v127_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_440_v127_.f12) == (_dafny.Map({_dafny.Map({-474: -474}): _dafny.Map({-474: 4}), _dafny.Map({-474: 4}): _dafny.Map({-474: 4}), _dafny.Map({1: 1}): _dafny.Map({-474: 4})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_440_v127_).f9).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_440_v127_).f10) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-71])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_442_v128_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_443_v129_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_444_v131_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_445_v132_) == (_dafny.MultiSet([_dafny.CodePoint('v')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_446_v133_) == (_dafny.Map({_dafny.Map({_dafny.CodePoint('v'): False}): -474}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_448_v134_) == (_dafny.Map({_dafny.CodePoint('v'): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_449_v135_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_450_v136_) == (_dafny.Map({_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjhipg")), False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgg"))}): _dafny.MultiSet([False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_451_v137_) == (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psiuxumem"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_517_v187_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_540_v204_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_541_v205_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v207_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v207_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_544_v208_) == (_dafny.Map({True: -474}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_545_v209_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({-474: 478, 1: 478})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_546_v210_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_547_v211_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_548_v212_.f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_548_v212_).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_549_v213_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_550_v214_).cf12) == (_dafny.Map({-474: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({self.cf0.VerbatimString(True)})'
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
        return lambda: D2_DC5(False, _dafny.Array(None, 0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)

class D2_DC5(D2, NamedTuple('DC5', [('cf4', Any), ('cf5', Any), ('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
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

class D3_DC7(D3, NamedTuple('DC7', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC6(D3, NamedTuple('DC6', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC9(D1.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)

class D4_DC9(D4, NamedTuple('DC9', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC8(D4, NamedTuple('DC8', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.CodePoint('D')
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC12(D5, NamedTuple('DC12', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC14(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)

class D6_DC14(D6, NamedTuple('DC14', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC13(D6, NamedTuple('DC13', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC16()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)

class D7_DC16(D7, NamedTuple('DC16', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC15(D7, NamedTuple('DC15', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC18()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D8_DC17)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)

class D8_DC18(D8, NamedTuple('DC18', [])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18)
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC17(D8, NamedTuple('DC17', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC17({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC17) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC19(D8, NamedTuple('DC19', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC21(D3.default()(), _dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)

class D9_DC21(D9, NamedTuple('DC21', [('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC20(D9, NamedTuple('DC20', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC23(_dafny.CodePoint('D'), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D10_DC22)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D10_DC24)

class D10_DC23(D10, NamedTuple('DC23', [('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC22(D10, NamedTuple('DC22', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC22({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC22) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC24(D10, NamedTuple('DC24', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC24({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC24) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC26()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D11_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)

class D11_DC26(D11, NamedTuple('DC26', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC25(D11, NamedTuple('DC25', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC29(False, int(0), _dafny.Seq({}), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D12_DC30)

class D12_DC29(D12, NamedTuple('DC29', [('cf32', Any), ('cf33', Any), ('cf34', Any), ('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC28(D12, NamedTuple('DC28', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC30(D12, NamedTuple('DC30', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC30({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC30) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D13_DC31)

class D13_DC31(D13, NamedTuple('DC31', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC31({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC31) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC33(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D14_DC33)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D14_DC34)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D14_DC32)

class D14_DC33(D14, NamedTuple('DC33', [('cf40', Any), ('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC33({_dafny.string_of(self.cf40)}, {self.cf41.VerbatimString(True)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC33) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC34(D14, NamedTuple('DC34', [('cf43', Any), ('cf44', Any), ('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC34({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC34) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC32(D14, NamedTuple('DC32', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC32({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC32) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC36(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D15_DC36)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D15_DC37)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D15_DC35)

class D15_DC36(D15, NamedTuple('DC36', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC36({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC36) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC37(D15, NamedTuple('DC37', [('cf49', Any), ('cf50', Any), ('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC37({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC37) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC35(D15, NamedTuple('DC35', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC35({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC35) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D16_DC38)

class D16_DC38(D16, NamedTuple('DC38', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC38({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC38) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC40(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D17_DC40)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D17_DC39)

class D17_DC40(D17, NamedTuple('DC40', [('cf54', Any), ('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC40({self.cf54.VerbatimString(True)}, {_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC40) and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC39(D17, NamedTuple('DC39', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC39({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC39) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC42(False, int(0))
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

class D18_DC42(D18, NamedTuple('DC42', [('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC42({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC42) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC43(D18, NamedTuple('DC43', [('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC43({_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC43) and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC41(D18, NamedTuple('DC41', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC41({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC41) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC44(D18, NamedTuple('DC44', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC44({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC44) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D19_DC45)

class D19_DC45(D19, NamedTuple('DC45', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC45({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC45) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC47(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D20_DC47)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D20_DC48)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D20_DC46)

class D20_DC47(D20, NamedTuple('DC47', [('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC47({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC47) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC48(D20, NamedTuple('DC48', [('cf68', Any), ('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC48({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC48) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC46(D20, NamedTuple('DC46', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC46({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC46) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC50(int(0), False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D21_DC50)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D21_DC49)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D21_DC51)

class D21_DC50(D21, NamedTuple('DC50', [('cf72', Any), ('cf73', Any), ('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC50({_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC50) and self.cf72 == __o.cf72 and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC49(D21, NamedTuple('DC49', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC49({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC49) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC51(D21, NamedTuple('DC51', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC51({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC51) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC53(_dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D22_DC53)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D22_DC54)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D22_DC52)

class D22_DC53(D22, NamedTuple('DC53', [('cf78', Any), ('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC53({_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC53) and self.cf78 == __o.cf78 and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC54(D22, NamedTuple('DC54', [('cf80', Any), ('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC54({_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC54) and self.cf80 == __o.cf80 and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC52(D22, NamedTuple('DC52', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC52({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC52) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC56(int(0), False, int(0), _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D23_DC56)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D23_DC55)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D23_DC57)

class D23_DC56(D23, NamedTuple('DC56', [('cf83', Any), ('cf84', Any), ('cf85', Any), ('cf86', Any), ('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC56({_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC56) and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85 and self.cf86 == __o.cf86 and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC55(D23, NamedTuple('DC55', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC55({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC55) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC57(D23, NamedTuple('DC57', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC57({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC57) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC59(_dafny.Seq({}), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D24_DC59)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D24_DC58)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D24_DC60)

class D24_DC59(D24, NamedTuple('DC59', [('cf90', Any), ('cf91', Any), ('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC59({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC59) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91 and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC58(D24, NamedTuple('DC58', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC58({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC58) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC60(D24, NamedTuple('DC60', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC60({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC60) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D25_DC61)

class D25_DC61(D25, NamedTuple('DC61', [('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC61({_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC61) and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D26_DC62)

class D26_DC62(D26, NamedTuple('DC62', [('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC62({_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC62) and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC64(_dafny.Map({}), False, int(0), _dafny.Seq({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D27_DC64)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D27_DC65)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D27_DC66)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D27_DC63)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D27_DC67)

class D27_DC64(D27, NamedTuple('DC64', [('cf97', Any), ('cf98', Any), ('cf99', Any), ('cf100', Any), ('cf101', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC64({_dafny.string_of(self.cf97)}, {_dafny.string_of(self.cf98)}, {_dafny.string_of(self.cf99)}, {_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC64) and self.cf97 == __o.cf97 and self.cf98 == __o.cf98 and self.cf99 == __o.cf99 and self.cf100 == __o.cf100 and self.cf101 == __o.cf101
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC65(D27, NamedTuple('DC65', [('cf102', Any), ('cf103', Any), ('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC65({_dafny.string_of(self.cf102)}, {self.cf103.VerbatimString(True)}, {_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC65) and self.cf102 == __o.cf102 and self.cf103 == __o.cf103 and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC66(D27, NamedTuple('DC66', [('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC66({_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC66) and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC63(D27, NamedTuple('DC63', [('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC63({_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC63) and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC67(D27, NamedTuple('DC67', [('cf106', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC67({_dafny.string_of(self.cf106)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC67) and self.cf106 == __o.cf106
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D28_DC68)

class D28_DC68(D28, NamedTuple('DC68', [('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC68({_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC68) and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D29_DC69)

class D29_DC69(D29, NamedTuple('DC69', [('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC69({_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC69) and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()


class D30:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D30_DC70)

class D30_DC70(D30, NamedTuple('DC70', [('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC70({_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC70) and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()


class D31:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D31_DC71)

class D31_DC71(D31, NamedTuple('DC71', [('cf110', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC71({_dafny.string_of(self.cf110)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC71) and self.cf110 == __o.cf110
    def __hash__(self) -> int:
        return super().__hash__()


class D32:
    @classmethod
    def default(cls, ):
        return lambda: D32_DC73()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D32_DC73)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D32_DC74)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D32_DC72)
    @property
    def is_DC75(self) -> bool:
        return isinstance(self, D32_DC75)

class D32_DC73(D32, NamedTuple('DC73', [])):
    def __dafnystr__(self) -> str:
        return f'D32.DC73'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC73)
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC74(D32, NamedTuple('DC74', [('cf112', Any), ('cf113', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC74({_dafny.string_of(self.cf112)}, {_dafny.string_of(self.cf113)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC74) and self.cf112 == __o.cf112 and self.cf113 == __o.cf113
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC72(D32, NamedTuple('DC72', [('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC72({_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC72) and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC75(D32, NamedTuple('DC75', [('cf114', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC75({_dafny.string_of(self.cf114)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC75) and self.cf114 == __o.cf114
    def __hash__(self) -> int:
        return super().__hash__()


class D33:
    @classmethod
    def default(cls, ):
        return lambda: D33_DC77(False, int(0), _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC77(self) -> bool:
        return isinstance(self, D33_DC77)
    @property
    def is_DC76(self) -> bool:
        return isinstance(self, D33_DC76)
    @property
    def is_DC78(self) -> bool:
        return isinstance(self, D33_DC78)

class D33_DC77(D33, NamedTuple('DC77', [('cf116', Any), ('cf117', Any), ('cf118', Any), ('cf119', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC77({_dafny.string_of(self.cf116)}, {_dafny.string_of(self.cf117)}, {_dafny.string_of(self.cf118)}, {_dafny.string_of(self.cf119)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC77) and self.cf116 == __o.cf116 and self.cf117 == __o.cf117 and self.cf118 == __o.cf118 and self.cf119 == __o.cf119
    def __hash__(self) -> int:
        return super().__hash__()

class D33_DC76(D33, NamedTuple('DC76', [('cf115', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC76({_dafny.string_of(self.cf115)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC76) and self.cf115 == __o.cf115
    def __hash__(self) -> int:
        return super().__hash__()

class D33_DC78(D33, NamedTuple('DC78', [('cf120', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC78({_dafny.string_of(self.cf120)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC78) and self.cf120 == __o.cf120
    def __hash__(self) -> int:
        return super().__hash__()


class D34:
    @classmethod
    def default(cls, ):
        return lambda: D34_DC80()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC80(self) -> bool:
        return isinstance(self, D34_DC80)
    @property
    def is_DC81(self) -> bool:
        return isinstance(self, D34_DC81)
    @property
    def is_DC79(self) -> bool:
        return isinstance(self, D34_DC79)

class D34_DC80(D34, NamedTuple('DC80', [])):
    def __dafnystr__(self) -> str:
        return f'D34.DC80'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D34_DC80)
    def __hash__(self) -> int:
        return super().__hash__()

class D34_DC81(D34, NamedTuple('DC81', [('cf122', Any), ('cf123', Any), ('cf124', Any)])):
    def __dafnystr__(self) -> str:
        return f'D34.DC81({_dafny.string_of(self.cf122)}, {_dafny.string_of(self.cf123)}, {_dafny.string_of(self.cf124)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D34_DC81) and self.cf122 == __o.cf122 and self.cf123 == __o.cf123 and self.cf124 == __o.cf124
    def __hash__(self) -> int:
        return super().__hash__()

class D34_DC79(D34, NamedTuple('DC79', [('cf121', Any)])):
    def __dafnystr__(self) -> str:
        return f'D34.DC79({_dafny.string_of(self.cf121)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D34_DC79) and self.cf121 == __o.cf121
    def __hash__(self) -> int:
        return super().__hash__()


class D35:
    @classmethod
    def default(cls, ):
        return lambda: D35_DC83(_dafny.Array(None, 0), False, int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC83(self) -> bool:
        return isinstance(self, D35_DC83)
    @property
    def is_DC84(self) -> bool:
        return isinstance(self, D35_DC84)
    @property
    def is_DC85(self) -> bool:
        return isinstance(self, D35_DC85)
    @property
    def is_DC82(self) -> bool:
        return isinstance(self, D35_DC82)
    @property
    def is_DC86(self) -> bool:
        return isinstance(self, D35_DC86)

class D35_DC83(D35, NamedTuple('DC83', [('cf126', Any), ('cf127', Any), ('cf128', Any), ('cf129', Any), ('cf130', Any)])):
    def __dafnystr__(self) -> str:
        return f'D35.DC83({_dafny.string_of(self.cf126)}, {_dafny.string_of(self.cf127)}, {_dafny.string_of(self.cf128)}, {_dafny.string_of(self.cf129)}, {_dafny.string_of(self.cf130)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D35_DC83) and self.cf126 == __o.cf126 and self.cf127 == __o.cf127 and self.cf128 == __o.cf128 and self.cf129 == __o.cf129 and self.cf130 == __o.cf130
    def __hash__(self) -> int:
        return super().__hash__()

class D35_DC84(D35, NamedTuple('DC84', [])):
    def __dafnystr__(self) -> str:
        return f'D35.DC84'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D35_DC84)
    def __hash__(self) -> int:
        return super().__hash__()

class D35_DC85(D35, NamedTuple('DC85', [('cf131', Any), ('cf132', Any)])):
    def __dafnystr__(self) -> str:
        return f'D35.DC85({_dafny.string_of(self.cf131)}, {_dafny.string_of(self.cf132)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D35_DC85) and self.cf131 == __o.cf131 and self.cf132 == __o.cf132
    def __hash__(self) -> int:
        return super().__hash__()

class D35_DC82(D35, NamedTuple('DC82', [('cf125', Any)])):
    def __dafnystr__(self) -> str:
        return f'D35.DC82({_dafny.string_of(self.cf125)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D35_DC82) and self.cf125 == __o.cf125
    def __hash__(self) -> int:
        return super().__hash__()

class D35_DC86(D35, NamedTuple('DC86', [('cf133', Any)])):
    def __dafnystr__(self) -> str:
        return f'D35.DC86({_dafny.string_of(self.cf133)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D35_DC86) and self.cf133 == __o.cf133
    def __hash__(self) -> int:
        return super().__hash__()


class D36:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC87(self) -> bool:
        return isinstance(self, D36_DC87)

class D36_DC87(D36, NamedTuple('DC87', [('cf134', Any)])):
    def __dafnystr__(self) -> str:
        return f'D36.DC87({_dafny.string_of(self.cf134)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D36_DC87) and self.cf134 == __o.cf134
    def __hash__(self) -> int:
        return super().__hash__()


class D37:
    @classmethod
    def default(cls, ):
        return lambda: D37_DC89(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC89(self) -> bool:
        return isinstance(self, D37_DC89)
    @property
    def is_DC90(self) -> bool:
        return isinstance(self, D37_DC90)
    @property
    def is_DC88(self) -> bool:
        return isinstance(self, D37_DC88)
    @property
    def is_DC91(self) -> bool:
        return isinstance(self, D37_DC91)

class D37_DC89(D37, NamedTuple('DC89', [('cf136', Any), ('cf137', Any)])):
    def __dafnystr__(self) -> str:
        return f'D37.DC89({_dafny.string_of(self.cf136)}, {_dafny.string_of(self.cf137)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D37_DC89) and self.cf136 == __o.cf136 and self.cf137 == __o.cf137
    def __hash__(self) -> int:
        return super().__hash__()

class D37_DC90(D37, NamedTuple('DC90', [('cf138', Any), ('cf139', Any), ('cf140', Any), ('cf141', Any), ('cf142', Any)])):
    def __dafnystr__(self) -> str:
        return f'D37.DC90({_dafny.string_of(self.cf138)}, {_dafny.string_of(self.cf139)}, {_dafny.string_of(self.cf140)}, {_dafny.string_of(self.cf141)}, {self.cf142.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D37_DC90) and self.cf138 == __o.cf138 and self.cf139 == __o.cf139 and self.cf140 == __o.cf140 and self.cf141 == __o.cf141 and self.cf142 == __o.cf142
    def __hash__(self) -> int:
        return super().__hash__()

class D37_DC88(D37, NamedTuple('DC88', [('cf135', Any)])):
    def __dafnystr__(self) -> str:
        return f'D37.DC88({_dafny.string_of(self.cf135)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D37_DC88) and self.cf135 == __o.cf135
    def __hash__(self) -> int:
        return super().__hash__()

class D37_DC91(D37, NamedTuple('DC91', [('cf143', Any)])):
    def __dafnystr__(self) -> str:
        return f'D37.DC91({_dafny.string_of(self.cf143)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D37_DC91) and self.cf143 == __o.cf143
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, p1, p2, globalState):
        pass


class T1:
    pass
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    def m2(self, p0, globalState):
        pass

    def m3(self, p0, p1, p2, p3, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f3: int = int(0)
        self._f0: _dafny.MultiSet = _dafny.MultiSet({})
        self._f1: bool = False
        self._f2: int = int(0)
        self._f4: int = int(0)
        self._f5: _dafny.Seq = _dafny.Seq({})
        self._f6: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
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
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6

class C0(T0):
    def  __init__(self):
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: _dafny.Seq = _dafny.Seq({})
        self._f26: bool = False
        self._f25: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f25, f26, f9, f10):
        (self)._f25 = f25
        (self)._f26 = f26
        (self)._f9 = f9
        (self)._f10 = f10

    def fm5(self, p0, p1, globalState):
        return ((_dafny.MultiSet([len((self).f9), 157, len(_dafny.Map({(self).f25: (self).f25})), 507, 207])).intersection(_dafny.MultiSet([len((self).f9), len((self).f9), (_dafny.MultiSet([(self).f9, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_637_i0_ in range(default__.abs(339))])])).cardinality]))).cardinality

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_638_i0_: int
        d_638_i0_ = 0
        with _dafny.label("2"):
            while (self).f26:
                with _dafny.c_label("2"):
                    if (d_638_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_638_i0_ = (d_638_i0_) + (1)
                    d_639_v0_: _dafny.Seq
                    d_639_v0_ = _dafny.SeqWithoutIsStrInference([(self).f25])
                    d_640_v1_: _dafny.MultiSet
                    d_640_v1_ = _dafny.MultiSet([default__.fm0(p1, len(d_639_v0_), globalState), 8])
                    d_641_v2_: _dafny.MultiSet
                    d_641_v2_ = _dafny.MultiSet([False, (self).f26])
                    d_642_v3_: _dafny.Set
                    d_642_v3_ = _dafny.Set({(d_641_v2_).cardinality})
                    d_643_v4_: _dafny.Array
                    nw107_ = _dafny.Array(None, 27)
                    nw107_[int(0)] = not(not(((self).f25) == (not((self).f25))))
                    nw107_[int(1)] = (self).f26
                    nw107_[int(2)] = (self).f25
                    nw107_[int(3)] = default__.fm38(p1, globalState)
                    nw107_[int(4)] = (self).f25
                    nw107_[int(5)] = (self).f26
                    nw107_[int(6)] = ((d_639_v0_)[default__.safeIndex(p1, len(d_639_v0_))]) == ((self).f26)
                    nw107_[int(7)] = (self).f26
                    nw107_[int(8)] = ((self).f25) and ((self).f25)
                    nw107_[int(9)] = default__.fm38(2, globalState)
                    nw107_[int(10)] = (self).f26
                    nw107_[int(11)] = not((self).f26)
                    nw107_[int(12)] = False
                    nw107_[int(13)] = (self).f26
                    nw107_[int(14)] = (self).f25
                    nw107_[int(15)] = default__.fm38(p2, globalState)
                    nw107_[int(16)] = not ((self).f25) or ((self).f25)
                    nw107_[int(17)] = (self).f25
                    nw107_[int(18)] = (False if (self).f25 else (self).f26)
                    nw107_[int(19)] = not((p2) <= ((d_640_v1_).cardinality))
                    nw107_[int(20)] = (default__.fm38(p1, globalState) if not((self).f25) else False)
                    nw107_[int(21)] = (d_642_v3_).ispropersubset(d_642_v3_)
                    nw107_[int(22)] = not ((self).f26) or (False)
                    nw107_[int(23)] = not((self).f26)
                    nw107_[int(24)] = (self).f26
                    nw107_[int(25)] = (self).f25
                    nw107_[int(26)] = (self).f26
                    d_643_v4_ = nw107_
                    d_643_v4_ = d_643_v4_
                    d_644_v5_: _dafny.Array
                    def lambda20_(d_645_p0_):
                        def lambda21_(d_646_i1_):
                            return (d_646_i1_) - (d_645_p0_)

                        return lambda21_

                    init10_ = lambda20_(p0)
                    nw108_ = _dafny.Array(None, 24)
                    for i0_10_ in range(nw108_.length(0)):
                        nw108_[i0_10_] = init10_(i0_10_)
                    d_644_v5_ = nw108_
                    d_647_v6_: _dafny.Map
                    d_647_v6_ = _dafny.Map({D3_DC6(d_642_v3_): d_644_v5_})
                    d_648_v7_: D3
                    d_648_v7_ = D3_DC6(_dafny.Set({p1, p2}))
                    d_647_v6_ = (d_647_v6_).set((D3_DC6(d_642_v3_) if (self).f25 else d_648_v7_), d_644_v5_)
                    r1 = (self).f26
                    rhs70_ = p1
                    rhs71_ = (p2) + (-169)
                    rhs72_ = (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f25])))
                    lhs47_ = globalState
                    lhs48_ = globalState
                    r3 = rhs70_
                    lhs47_.f3 = rhs71_
                    lhs48_.f3 = rhs72_
                    pass
            pass
        d_649_v8_: _dafny.Array
        nw109_ = _dafny.Array(D12.default()(), 17)
        d_649_v8_ = nw109_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_649_v8_).length(0)):
            d_650_i2_: int = guard_loop_0_
            if (True) and (((0) <= (d_650_i2_)) and ((d_650_i2_) < ((d_649_v8_).length(0)))):
                (d_649_v8_)[(d_650_i2_)] = D12_DC29((self).f26, (p0 if (self).f25 else p2), _dafny.SeqWithoutIsStrInference([(D12_DC29((self).f26, p0, _dafny.SeqWithoutIsStrInference([p1 for d_651_i3_ in range(default__.abs(130))]), True, -940)).cf36, p2, p0, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_652_i4_ in range(default__.abs(266))])), p1]), (self).f25, len(_dafny.SeqWithoutIsStrInference([943])))
        d_653_v9_: _dafny.Array
        nw110_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
        d_653_v9_ = nw110_
        index73_ = default__.safeIndex(487, (d_653_v9_).length(0))
        (d_653_v9_)[index73_] = (self).f9
        index74_ = default__.safeIndex(487, (d_653_v9_).length(0))
        (d_653_v9_)[index74_] = (self).f9
        d_654_v10_: _dafny.Map
        d_654_v10_ = _dafny.Map({(self).f26: (self).f9})
        d_655_i5_: int
        d_655_i5_ = 0
        with _dafny.label("3"):
            while not(not ((((d_654_v10_)[(self).f25] if ((self).f25) in (d_654_v10_) else (self).f9)) <= ((d_653_v9_)[default__.safeIndex(487, (d_653_v9_).length(0))])) or ((self).f25)):
                with _dafny.c_label("3"):
                    if (d_655_i5_) >= (100):
                        raise _dafny.Break("3")
                    d_655_i5_ = (d_655_i5_) + (1)
                    (globalState).f3 = p2
                    d_656_v11_: _dafny.Array
                    def lambda22_(d_657_p0_):
                        def lambda23_(d_658_i6_):
                            return (d_658_i6_) + (d_657_p0_)

                        return lambda23_

                    init11_ = lambda22_(p0)
                    nw111_ = _dafny.Array(None, 24)
                    for i0_11_ in range(nw111_.length(0)):
                        nw111_[i0_11_] = init11_(i0_11_)
                    d_656_v11_ = nw111_
                    d_659_v12_: _dafny.Map
                    d_659_v12_ = _dafny.Map({p2: d_656_v11_})
                    d_660_v13_: D6
                    d_660_v13_ = D6_DC14(d_656_v11_)
                    d_661_v14_: _dafny.Seq
                    d_661_v14_ = _dafny.SeqWithoutIsStrInference([d_656_v11_, d_656_v11_])
                    d_662_v15_: _dafny.Array
                    nw112_ = _dafny.Array(None, 28)
                    nw112_[int(0)] = d_656_v11_
                    nw112_[int(1)] = ((d_659_v12_)[p1] if (p1) in (d_659_v12_) else (d_660_v13_).cf17)
                    nw112_[int(2)] = d_656_v11_
                    nw112_[int(3)] = d_656_v11_
                    nw112_[int(4)] = d_656_v11_
                    nw112_[int(5)] = d_656_v11_
                    nw112_[int(6)] = d_656_v11_
                    nw112_[int(7)] = d_656_v11_
                    nw112_[int(8)] = d_656_v11_
                    nw112_[int(9)] = d_656_v11_
                    nw112_[int(10)] = d_656_v11_
                    nw112_[int(11)] = d_656_v11_
                    nw112_[int(12)] = d_656_v11_
                    nw112_[int(13)] = d_656_v11_
                    nw112_[int(14)] = d_656_v11_
                    nw112_[int(15)] = (d_661_v14_)[default__.safeIndex(p0, len(d_661_v14_))]
                    nw112_[int(16)] = (d_661_v14_)[default__.safeIndex(p2, len(d_661_v14_))]
                    nw112_[int(17)] = d_656_v11_
                    nw112_[int(18)] = d_656_v11_
                    nw112_[int(19)] = d_656_v11_
                    nw112_[int(20)] = d_656_v11_
                    nw112_[int(21)] = d_656_v11_
                    nw112_[int(22)] = d_656_v11_
                    nw112_[int(23)] = d_656_v11_
                    nw112_[int(24)] = d_656_v11_
                    nw112_[int(25)] = d_656_v11_
                    nw112_[int(26)] = (d_661_v14_)[default__.safeIndex(p0, len(d_661_v14_))]
                    nw112_[int(27)] = d_656_v11_
                    d_662_v15_ = nw112_
                    index75_ = default__.safeIndex(514, (d_662_v15_).length(0))
                    (d_662_v15_)[index75_] = d_656_v11_
                    index76_ = default__.safeIndex(514, (d_662_v15_).length(0))
                    (d_662_v15_)[index76_] = (d_656_v11_ if not ((self).f26) or ((self).f25) else d_656_v11_)
                    d_663_v16_: _dafny.Array
                    nw113_ = _dafny.Array(False, 2)
                    d_663_v16_ = nw113_
                    d_663_v16_ = d_663_v16_
                    r1 = not ((self).f26) or ((self).f25)
                    pass
            pass
        index77_ = default__.safeIndex(487, (d_653_v9_).length(0))
        (d_653_v9_)[index77_] = (d_653_v9_)[default__.safeIndex(487, (d_653_v9_).length(0))]
        if (self).f26:
            r1 = (p1) < (p0)
            r0 = ((self).f25 if (self).f26 else default__.fm38(p0, globalState))
            d_664_v17_: _dafny.Map
            d_664_v17_ = _dafny.Map({(self).f25: ((self).f25) and ((self).f26)})
            d_664_v17_ = (d_664_v17_) | ((d_664_v17_) | (d_664_v17_))
            d_665_v18_: _dafny.Seq
            d_665_v18_ = _dafny.SeqWithoutIsStrInference([(self).f26])
            d_666_v19_: _dafny.Map
            d_666_v19_ = _dafny.Map({default__.safeDivisionInt(len(d_665_v18_), p0): 975})
            (globalState).f3 = len(d_666_v19_)
            d_667_v20_: str
            d_667_v20_ = _dafny.CodePoint('w')
            d_668_v21_: _dafny.Seq
            d_668_v21_ = _dafny.SeqWithoutIsStrInference([p2])
            d_669_v22_: _dafny.MultiSet
            d_669_v22_ = _dafny.MultiSet([p2, (d_668_v21_)[default__.safeIndex(p0, len(d_668_v21_))]])
            d_670_v23_: D7
            d_670_v23_ = D7_DC15(d_669_v22_)
            d_667_v20_ = default__.fm39(d_670_v23_, ((0) - (p0)) in (default__.fm40(p2, p2, globalState)), globalState)
        elif True:
            r0 = (self).f25
            d_671_v24_: _dafny.Map
            d_671_v24_ = _dafny.Map({p0: (d_653_v9_)[default__.safeIndex(487, (d_653_v9_).length(0))]})
            d_672_v25_: _dafny.Seq
            d_672_v25_ = _dafny.SeqWithoutIsStrInference([(self).f9, (self).f9, ((d_671_v24_)[p2] if (p2) in (d_671_v24_) else (d_653_v9_)[default__.safeIndex(487, (d_653_v9_).length(0))])])
            r3 = default__.safeDivisionInt(len((d_672_v25_)[default__.safeIndex(p1, len(d_672_v25_))]), p0)
            d_673_v26_: _dafny.Array
            nw114_ = _dafny.Array(_dafny.Map({}), 29)
            d_673_v26_ = nw114_
            index78_ = default__.safeIndex(59, (d_673_v26_).length(0))
            def iife60_():
                coll52_ = _dafny.Map()
                compr_52_: int
                for compr_52_ in _dafny.IntegerRange(298, 124):
                    d_674_v27_: int = compr_52_
                    if ((298) <= (d_674_v27_)) and ((d_674_v27_) < (124)):
                        coll52_[default__.safeDivisionInt(d_674_v27_, 322)] = _dafny.SeqWithoutIsStrInference([p0, p1])
                return _dafny.Map(coll52_)
            (d_673_v26_)[index78_] = iife60_()
            
            d_675_v28_: _dafny.Seq
            d_675_v28_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_676_v29_: _dafny.Map
            d_676_v29_ = _dafny.Map({default__.safeModuloInt(p2, p1): d_675_v28_})
            index79_ = default__.safeIndex(59, (d_673_v26_).length(0))
            (d_673_v26_)[index79_] = d_676_v29_
            d_677_v30_: _dafny.Array
            def lambda24_(d_678_p0_):
                def lambda25_(d_679_i7_):
                    return (d_679_i7_) * (d_678_p0_)

                return lambda25_

            init12_ = lambda24_(p0)
            nw115_ = _dafny.Array(None, 18)
            for i0_12_ in range(nw115_.length(0)):
                nw115_[i0_12_] = init12_(i0_12_)
            d_677_v30_ = nw115_
            r2 = d_677_v30_
            d_680_v31_: _dafny.Array
            nw116_ = _dafny.Array(False, 5)
            d_680_v31_ = nw116_
            index80_ = default__.safeIndex(658, (d_680_v31_).length(0))
            (d_680_v31_)[index80_] = (self).f26
            index81_ = default__.safeIndex(658, (d_680_v31_).length(0))
            (d_680_v31_)[index81_] = False
        r0 = False
        r1 = not ((p2) >= ((self).fm5(p1, p1, globalState))) or (default__.fm38(len(_dafny.SeqWithoutIsStrInference([(d_653_v9_)[default__.safeIndex(487, (d_653_v9_).length(0))]])), globalState))
        d_681_v32_: _dafny.Array
        def lambda26_(d_682_p2_):
            def lambda27_(d_683_i8_):
                return (d_683_i8_) - (d_682_p2_)

            return lambda27_

        init13_ = lambda26_(p2)
        nw117_ = _dafny.Array(None, 15)
        for i0_13_ in range(nw117_.length(0)):
            nw117_[i0_13_] = init13_(i0_13_)
        d_681_v32_ = nw117_
        r2 = d_681_v32_
        d_684_v33_: _dafny.Seq
        d_684_v33_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, default__.fm38(909, globalState)])
        d_685_v34_: _dafny.Map
        d_685_v34_ = _dafny.Map({p2: (d_653_v9_)[default__.safeIndex(487, (d_653_v9_).length(0))]})
        r3 = (self).fm5(len(d_684_v33_), len(((d_685_v34_)[p2] if (p2) in (d_685_v34_) else (self).f9)), globalState)
        return r0, r1, r2, r3

    @property
    def f26(self):
        return self._f26
    @property
    def f25(self):
        return self._f25

class C1(T1, T0):
    def  __init__(self):
        self._f12: _dafny.Map = _dafny.Map({})
        self._f11: int = int(0)
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: _dafny.Seq = _dafny.Seq({})
        self.f23: int = int(0)
        self._f24: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f11(self):
        return self._f11
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f23, f24, f11, f12, f9, f10):
        (self).f23 = f23
        (self)._f24 = f24
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f9 = f9
        (self)._f10 = f10

    def fm5(self, p0, p1, globalState):
        return (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtoj"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qunejed"))))) - ((0) - (default__.safeDivisionInt((self).f24, self.f23)))

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: bool = False
        r3: int = int(0)
        d_686_v0_: bool
        d_686_v0_ = True
        r2 = d_686_v0_
        d_687_v1_: _dafny.Map
        d_687_v1_ = _dafny.Map({self.f23: d_686_v0_})
        d_688_v2_: _dafny.MultiSet
        d_688_v2_ = _dafny.MultiSet([_dafny.Map({p0: d_686_v0_}), d_687_v1_])
        d_689_v3_: _dafny.Map
        d_689_v3_ = _dafny.Map({d_686_v0_: (d_688_v2_).cardinality})
        d_690_v4_: _dafny.Map
        d_690_v4_ = _dafny.Map({((d_689_v3_)[d_686_v0_] if (d_686_v0_) in (d_689_v3_) else len((self).f9)): p0})
        d_691_v5_: _dafny.Seq
        d_691_v5_ = _dafny.SeqWithoutIsStrInference([d_686_v0_])
        d_692_v6_: _dafny.Array
        nw118_ = _dafny.Array(int(0), 27)
        d_692_v6_ = nw118_
        d_693_v7_: _dafny.Map
        d_693_v7_ = _dafny.Map({d_692_v6_: d_686_v0_})
        d_694_v8_: _dafny.Seq
        d_694_v8_ = _dafny.SeqWithoutIsStrInference([(default__.fm36((self).f24, d_686_v0_, p0, globalState)).cardinality])
        d_695_v9_: _dafny.Array
        nw119_ = _dafny.Array(None, 14)
        nw119_[int(0)] = (self).f24
        nw119_[int(1)] = p0
        nw119_[int(2)] = p0
        nw119_[int(3)] = self.f23
        nw119_[int(4)] = p0
        nw119_[int(5)] = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_696_i0_ in range(default__.abs(-329))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnqrwgvj"))))
        nw119_[int(6)] = (0) - (self.f23)
        nw119_[int(7)] = ((d_690_v4_)[(self).f24] if ((self).f24) in (d_690_v4_) else len(default__.fm35(len(d_691_v5_), self.f23, globalState)))
        nw119_[int(8)] = p0
        nw119_[int(9)] = p0
        nw119_[int(10)] = (self).f24
        nw119_[int(11)] = (self).f24
        nw119_[int(12)] = len((_dafny.Map({d_692_v6_: True})) | (d_693_v7_))
        nw119_[int(13)] = default__.safeModuloInt((self).f11, (d_694_v8_)[default__.safeIndex(self.f23, len(d_694_v8_))])
        d_695_v9_ = nw119_
        index82_ = default__.safeIndex(518, (d_695_v9_).length(0))
        (d_695_v9_)[index82_] = (self).f11
        index83_ = default__.safeIndex(518, (d_695_v9_).length(0))
        (d_695_v9_)[index83_] = (self).f24
        d_697_v10_: _dafny.Array
        nw120_ = _dafny.Array(_dafny.Set({}), 11)
        d_697_v10_ = nw120_
        d_698_v12_: _dafny.Set
        d_698_v12_ = _dafny.Set({p0})
        index84_ = default__.safeIndex(167, (d_697_v10_).length(0))
        def iife61_():
            coll53_ = _dafny.Set()
            compr_53_: int
            for compr_53_ in ((D12_DC28((_dafny.SeqWithoutIsStrInference([self.f23 for d_699_i1_ in range(default__.abs(636))])).set(default__.safeIndex((d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))], len(_dafny.SeqWithoutIsStrInference([self.f23 for d_700_i1_ in range(default__.abs(636))]))), p0))).cf31).Elements:
                d_701_v11_: int = compr_53_
                if (d_701_v11_) in ((D12_DC28((_dafny.SeqWithoutIsStrInference([self.f23 for d_699_i1_ in range(default__.abs(636))])).set(default__.safeIndex((d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))], len(_dafny.SeqWithoutIsStrInference([self.f23 for d_700_i1_ in range(default__.abs(636))]))), p0))).cf31):
                    coll53_ = coll53_.union(_dafny.Set([(d_701_v11_) * (len(_dafny.Map({not(True): len(_dafny.Set({True, not(True)}))})))]))
            return _dafny.Set(coll53_)
        (d_697_v10_)[index84_] = (iife61_()
         if d_686_v0_ else d_698_v12_)
        index85_ = default__.safeIndex(167, (d_697_v10_).length(0))
        (d_697_v10_)[index85_] = (_dafny.Set({(d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))], (0) - (p0)})).intersection((d_698_v12_) | (_dafny.Set({default__.fm0(974, len(d_689_v3_), globalState)})))
        index86_ = default__.safeIndex(518, (d_695_v9_).length(0))
        (d_695_v9_)[index86_] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "httgemkqe"))), (100) + ((self).f24))
        r2 = d_686_v0_
        d_702_v13_: _dafny.Array
        nw121_ = _dafny.Array(_dafny.CodePoint('D'), 7)
        d_702_v13_ = nw121_
        d_703_v14_: D6
        d_703_v14_ = D6_DC13(d_702_v13_)
        pat_let_tv11_ = d_702_v13_
        def iife62_(_pat_let4_0):
            def iife63_(d_704_dt__update__tmp_h0_):
                def iife64_(_pat_let5_0):
                    def iife65_(d_705_dt__update_hcf16_h0_):
                        return D6_DC13(d_705_dt__update_hcf16_h0_)
                    return iife65_(_pat_let5_0)
                return iife64_(pat_let_tv11_)
            return iife63_(_pat_let4_0)
        source12_ = iife62_(d_703_v14_)
        if source12_.is_DC14:
            d_706___mcc_h0_ = source12_.cf17
            d_707_cf17_ = d_706___mcc_h0_
            d_708_v15_: _dafny.Map
            d_708_v15_ = _dafny.Map({(self).f11: _dafny.CodePoint('y')})
            d_709_v16_: _dafny.Map
            d_709_v16_ = _dafny.Map({(d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))]: (d_708_v15_).set(self.f23, _dafny.CodePoint('t'))})
            d_709_v16_ = (d_709_v16_).set(self.f23, d_708_v15_)
            r2 = ((-313) + (self.f23)) > ((d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))])
            index87_ = default__.safeIndex(518, (d_695_v9_).length(0))
            (d_695_v9_)[index87_] = default__.safeModuloInt(((0) - (default__.fm0((d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))], (d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))], globalState))) + ((self).f24), (0) - ((self).f24))
            rhs73_ = (d_691_v5_)[default__.safeIndex(((self).f24) + ((self).f11), len(d_691_v5_))]
            rhs74_ = (d_707_cf17_ if d_686_v0_ else d_707_cf17_)
            rhs75_ = ((0) - ((len(d_694_v8_) if False else (self).f24))) * (p0)
            rhs76_ = (self).f11
            lhs49_ = globalState
            d_686_v0_ = rhs73_
            d_707_cf17_ = rhs74_
            r0 = rhs75_
            lhs49_.f3 = rhs76_
        elif True:
            d_710___mcc_h1_ = source12_.cf16
            d_711_cf16_ = d_710___mcc_h1_
            d_686_v0_ = not(((d_687_v1_)[(self).f24] if ((self).f24) in (d_687_v1_) else (d_686_v0_) or (not(d_686_v0_))))
            if d_686_v0_:
                index88_ = default__.safeIndex(518, (d_695_v9_).length(0))
                (d_695_v9_)[index88_] = 909
                d_712_v17_: _dafny.Set
                d_712_v17_ = _dafny.Set({not(d_686_v0_)})
                d_712_v17_ = d_712_v17_
                index89_ = default__.safeIndex(518, (d_695_v9_).length(0))
                rhs77_ = (self.f23) * (p0)
                rhs78_ = (self).f11
                rhs79_ = (d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))]
                lhs50_ = d_695_v9_
                lhs51_ = default__.safeIndex(518, (d_695_v9_).length(0))
                lhs52_ = globalState
                lhs50_[lhs51_] = rhs77_
                lhs52_.f3 = rhs78_
                r0 = rhs79_
                (globalState).f3 = (self).f11
                d_713_v18_: D7
                d_713_v18_ = D7_DC16()
                d_714_v19_: _dafny.Seq
                d_714_v19_ = _dafny.SeqWithoutIsStrInference([d_713_v18_])
                d_715_v20_: _dafny.Seq
                d_715_v20_ = _dafny.SeqWithoutIsStrInference([d_714_v19_, d_714_v19_])
                d_716_v21_: _dafny.Set
                d_716_v21_ = _dafny.Set({((d_715_v20_)[default__.safeIndex(self.f23, len(d_715_v20_))]).set(default__.safeIndex(p0, len((d_715_v20_)[default__.safeIndex(self.f23, len(d_715_v20_))])), d_713_v18_), _dafny.SeqWithoutIsStrInference([d_713_v18_ for d_717_i2_ in range(default__.abs(430))])})
                d_716_v21_ = d_716_v21_
            elif True:
                d_718_v22_: _dafny.Map
                d_718_v22_ = _dafny.Map({d_686_v0_: d_686_v0_})
                d_719_v23_: _dafny.Set
                d_719_v23_ = _dafny.Set({d_686_v0_, d_686_v0_})
                d_720_v24_: _dafny.Array
                nw122_ = _dafny.Array(None, 15)
                nw122_[int(0)] = d_686_v0_
                nw122_[int(1)] = True
                nw122_[int(2)] = ((d_718_v22_)[d_686_v0_] if (d_686_v0_) in (d_718_v22_) else d_686_v0_)
                nw122_[int(3)] = (d_719_v23_).ispropersubset(_dafny.Set({d_686_v0_, d_686_v0_}))
                nw122_[int(4)] = d_686_v0_
                nw122_[int(5)] = d_686_v0_
                nw122_[int(6)] = d_686_v0_
                nw122_[int(7)] = d_686_v0_
                nw122_[int(8)] = d_686_v0_
                nw122_[int(9)] = (not(d_686_v0_) if False else d_686_v0_)
                nw122_[int(10)] = d_686_v0_
                nw122_[int(11)] = d_686_v0_
                nw122_[int(12)] = d_686_v0_
                nw122_[int(13)] = ((self).f24) < ((self).f24)
                nw122_[int(14)] = d_686_v0_
                d_720_v24_ = nw122_
                index90_ = default__.safeIndex(528, (d_720_v24_).length(0))
                (d_720_v24_)[index90_] = d_686_v0_
                index91_ = default__.safeIndex(528, (d_720_v24_).length(0))
                (d_720_v24_)[index91_] = d_686_v0_
                d_686_v0_ = not (d_686_v0_) or (((d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))]) <= (self.f23))
                d_721_v25_: _dafny.Seq
                d_721_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oyenj"))
                d_721_v25_ = (self).f9
                d_722_v26_: _dafny.Map
                d_722_v26_ = _dafny.Map({d_721_v25_: d_721_v25_})
                d_723_v27_: _dafny.MultiSet
                d_723_v27_ = _dafny.MultiSet([d_686_v0_])
                d_722_v26_ = (d_722_v26_).set(default__.fm37((self).fm5(p0, len(_dafny.Map({d_692_v6_: (d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))]})), globalState), (d_723_v27_).cardinality, d_686_v0_, (d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))], globalState), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_724_i3_ in range(default__.abs(-329))])) + ((self).f9))
                index92_ = default__.safeIndex(528, (d_720_v24_).length(0))
                (d_720_v24_)[index92_] = d_686_v0_
            d_686_v0_ = (not((_dafny.SeqWithoutIsStrInference([len(d_690_v4_) for d_725_i4_ in range(default__.abs(-413))])) != (d_694_v8_)) if d_686_v0_ else d_686_v0_)
            if (d_686_v0_) == (not(d_686_v0_)):
                r3 = (d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))]
                d_726_v28_: _dafny.Set
                d_726_v28_ = _dafny.Set({(self).f9, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_727_i5_ in range(default__.abs(15))]), ((self).f9) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gps")))})
                d_728_v29_: _dafny.Seq
                d_728_v29_ = _dafny.SeqWithoutIsStrInference([(self).f9, (self).f9])
                d_726_v28_ = _dafny.Set({(d_728_v29_)[default__.safeIndex((self).f11, len(d_728_v29_))]})
                index93_ = default__.safeIndex(518, (d_695_v9_).length(0))
                (d_695_v9_)[index93_] = (0) - (((p0 if False else p0)) * ((p0 if d_686_v0_ else (self).f11)))
                d_686_v0_ = not(d_686_v0_)
                d_729_v30_: _dafny.Seq
                d_729_v30_ = ((self).f9) + ((self).f9)
                d_729_v30_ = (self).f9
            elif True:
                def iife66_():
                    coll54_ = _dafny.Set()
                    compr_54_: int
                    for compr_54_ in _dafny.IntegerRange(162, 296):
                        d_730_v31_: int = compr_54_
                        if ((162) <= (d_730_v31_)) and ((d_730_v31_) < (296)):
                            coll54_ = coll54_.union(_dafny.Set([(d_730_v31_) - ((d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))])]))
                    return _dafny.Set(coll54_)
                d_698_v12_ = (((d_697_v10_)[default__.safeIndex(167, (d_697_v10_).length(0))]).intersection(iife66_()
                )) - ((d_698_v12_).intersection(d_698_v12_))
                d_731_v32_: _dafny.Array
                def lambda28_(d_732_v9_):
                    def lambda29_(d_733_i6_):
                        return D10_DC23(_dafny.CodePoint('g'), _dafny.Set({_dafny.MultiSet([(d_732_v9_)[default__.safeIndex(518, (d_732_v9_).length(0))]])}))

                    return lambda29_

                init14_ = lambda28_(d_695_v9_)
                nw123_ = _dafny.Array(None, 24)
                for i0_14_ in range(nw123_.length(0)):
                    nw123_[i0_14_] = init14_(i0_14_)
                d_731_v32_ = nw123_
                d_734_v33_: _dafny.Map
                d_734_v33_ = _dafny.Map({(d_686_v0_ if d_686_v0_ else not(d_686_v0_)): d_731_v32_})
                d_734_v33_ = ((d_734_v33_).set(d_686_v0_, d_731_v32_) if ((d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))]) >= (len(d_691_v5_)) else d_734_v33_)
                r2 = False
                r2 = True
                d_735_v34_: D12
                d_735_v34_ = D12_DC29(d_686_v0_, (self).fm5(p0, self.f23, globalState), d_694_v8_, d_686_v0_, len((self).f9))
                d_736_v35_: _dafny.Seq
                d_736_v35_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10, _dafny.SeqWithoutIsStrInference([d_694_v8_ for d_737_i7_ in range(default__.abs(-969))]), _dafny.SeqWithoutIsStrInference([d_694_v8_, d_694_v8_])])
                d_738_v36_: _dafny.Map
                d_738_v36_ = _dafny.Map({d_686_v0_: d_686_v0_})
                d_739_v37_: C0
                nw124_ = C0()
                nw124_.ctor__((d_691_v5_)[default__.safeIndex((self).f24, len(d_691_v5_))], ((d_735_v34_).cf35 if not(d_686_v0_) else False), (self).f9, (d_736_v35_)[default__.safeIndex(len(d_738_v36_), len(d_736_v35_))])
                d_739_v37_ = nw124_
        r0 = default__.fm0((0) - (len(d_694_v8_)), default__.safeModuloInt(p0, len(default__.fm37((self).f11, len((self).f9), False, (d_695_v9_)[default__.safeIndex(518, (d_695_v9_).length(0))], globalState))), globalState)
        d_740_v38_: _dafny.Map
        d_740_v38_ = _dafny.Map({d_686_v0_: (self).f9})
        r1 = ((d_740_v38_) | (d_740_v38_)) | (d_740_v38_)
        r2 = d_686_v0_
        d_741_v39_: D4
        d_741_v39_ = D4_DC11(not(d_686_v0_), (self).f11)
        r3 = default__.fm0((self).f24, default__.safeModuloInt(p0, (d_741_v39_).cf14), globalState)
        return r0, r1, r2, r3

    def m3(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_742_v0_: _dafny.Array
        nw125_ = _dafny.Array(int(0), 8)
        d_742_v0_ = nw125_
        r1 = d_742_v0_
        d_743_v1_: bool
        d_743_v1_ = False
        d_744_v2_: D1
        d_744_v2_ = D1_DC1(d_743_v1_)
        pat_let_tv12_ = d_743_v1_
        def iife67_(_pat_let6_0):
            def iife68_(d_745_dt__update__tmp_h0_):
                def iife69_(_pat_let7_0):
                    def iife70_(d_746_dt__update_hcf1_h0_):
                        return D1_DC1(d_746_dt__update_hcf1_h0_)
                    return iife70_(_pat_let7_0)
                return iife69_(pat_let_tv12_)
            return iife68_(_pat_let6_0)
        d_744_v2_ = iife67_(D1_DC1(d_743_v1_))
        (self).f23 = default__.safeModuloInt((self).f24, self.f23)
        d_747_v3_: D7
        d_747_v3_ = D7_DC16()
        pat_let_tv13_ = d_743_v1_
        pat_let_tv14_ = d_743_v1_
        pat_let_tv15_ = d_743_v1_
        def lambda30_(source13_):
            if source13_.is_DC16:
                return False
            elif True:
                d_748___mcc_h0_ = source13_.cf18
                d_749_cf18_ = d_748___mcc_h0_
                if pat_let_tv13_:
                    return pat_let_tv14_
                elif True:
                    return pat_let_tv15_

        d_743_v1_ = lambda30_(d_747_v3_)
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_742_v0_).length(0)):
            d_750_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_750_i0_)) and ((d_750_i0_) < ((d_742_v0_).length(0)))):
                (d_742_v0_)[(d_750_i0_)] = default__.safeModuloInt(d_750_i0_, (D12_DC29(d_743_v1_, p2, _dafny.SeqWithoutIsStrInference([self.f23, p0]), d_743_v1_, self.f23)).cf33)
        if d_743_v1_:
            if default__.fm38(p2, globalState):
                d_751_v4_: _dafny.Array
                def lambda31_(d_752_i1_):
                    return ((self).f9) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrhjt")))

                init15_ = lambda31_
                nw126_ = _dafny.Array(None, 5)
                for i0_15_ in range(nw126_.length(0)):
                    nw126_[i0_15_] = init15_(i0_15_)
                d_751_v4_ = nw126_
                index94_ = default__.safeIndex(321, (d_751_v4_).length(0))
                (d_751_v4_)[index94_] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_753_i2_ in range(default__.abs(735))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmpip")))
                index95_ = default__.safeIndex(321, (d_751_v4_).length(0))
                (d_751_v4_)[index95_] = (p1) + ((self).f9)
                index96_ = default__.safeIndex(633, (d_742_v0_).length(0))
                (d_742_v0_)[index96_] = self.f23
                d_754_v5_: _dafny.Map
                d_754_v5_ = _dafny.Map({d_743_v1_: (0) - (self.f23)})
                index97_ = default__.safeIndex(633, (d_742_v0_).length(0))
                (d_742_v0_)[index97_] = (len(p1)) + (len(d_754_v5_))
                d_755_v7_: str
                d_755_v7_ = _dafny.CodePoint('p')
                d_756_v8_: _dafny.MultiSet
                d_756_v8_ = _dafny.MultiSet([(d_742_v0_)[default__.safeIndex(633, (d_742_v0_).length(0))]])
                d_757_v9_: _dafny.Seq
                d_757_v9_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([p0, (d_742_v0_)[default__.safeIndex(633, (d_742_v0_).length(0))]]), d_756_v8_])
                d_758_v10_: _dafny.Map
                d_758_v10_ = _dafny.Map({p0: d_743_v1_})
                d_759_v11_: _dafny.Set
                d_759_v11_ = _dafny.Set({d_758_v10_})
                d_760_v12_: _dafny.Map
                d_760_v12_ = _dafny.Map({p2: self.f23})
                d_761_v14_: _dafny.MultiSet
                d_761_v14_ = _dafny.MultiSet([d_760_v12_])
                d_762_v15_: _dafny.MultiSet
                d_762_v15_ = d_761_v14_
                d_763_v16_: _dafny.Array
                nw127_ = _dafny.Array(None, 20)
                def iife71_():
                    coll55_ = _dafny.Set()
                    compr_55_: int
                    for compr_55_ in (_dafny.Map({p2: d_743_v1_})).keys.Elements:
                        d_764_v6_: int = compr_55_
                        if (d_764_v6_) in (_dafny.Map({p2: d_743_v1_})):
                            coll55_ = coll55_.union(_dafny.Set([(d_764_v6_) + (len(_dafny.Set({True})))]))
                    return _dafny.Set(coll55_)
                nw127_[int(0)] = (p0) > (len(iife71_()
                ))
                nw127_[int(1)] = d_743_v1_
                nw127_[int(2)] = d_743_v1_
                nw127_[int(3)] = (d_755_v7_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldwpvxod")))
                nw127_[int(4)] = (d_744_v2_).cf1
                nw127_[int(5)] = not((d_743_v1_) and (d_743_v1_))
                nw127_[int(6)] = (d_743_v1_ if d_743_v1_ else not(d_743_v1_))
                nw127_[int(7)] = ((d_757_v9_)[default__.safeIndex(48, len(d_757_v9_))]).ispropersubset(d_756_v8_)
                nw127_[int(8)] = default__.fm38(self.f23, globalState)
                nw127_[int(9)] = not(d_743_v1_)
                nw127_[int(10)] = d_743_v1_
                nw127_[int(11)] = d_743_v1_
                nw127_[int(12)] = d_743_v1_
                nw127_[int(13)] = (d_743_v1_) or (d_743_v1_)
                nw127_[int(14)] = d_743_v1_
                nw127_[int(15)] = d_743_v1_
                nw127_[int(16)] = (d_759_v11_).issubset(d_759_v11_)
                def iife72_():
                    coll56_ = _dafny.Map()
                    compr_56_: int
                    for compr_56_ in _dafny.IntegerRange(974, -198):
                        d_765_v13_: int = compr_56_
                        if ((974) <= (d_765_v13_)) and ((d_765_v13_) < (-198)):
                            coll56_[default__.safeModuloInt(d_765_v13_, self.f23)] = (self).f11
                    return _dafny.Map(coll56_)
                nw127_[int(17)] = (_dafny.MultiSet([d_760_v12_, (iife72_()
                ).set(p0, (self).f24)])) != ((d_762_v15_))
                nw127_[int(18)] = default__.fm38(p0, globalState)
                nw127_[int(19)] = d_743_v1_
                d_763_v16_ = nw127_
                d_766_v17_: _dafny.Set
                d_766_v17_ = _dafny.Set({d_763_v16_, d_763_v16_})
                d_767_v18_: _dafny.Set
                d_767_v18_ = _dafny.Set({True})
                d_768_v19_: _dafny.Map
                d_768_v19_ = _dafny.Map({(d_766_v17_).isdisjoint(d_766_v17_): (d_767_v18_) == (d_767_v18_)})
                index98_ = default__.safeIndex(321, (d_751_v4_).length(0))
                rhs80_ = d_763_v16_
                rhs81_ = (0) - (((d_742_v0_)[default__.safeIndex(633, (d_742_v0_).length(0))]) * ((self).f11))
                rhs82_ = d_768_v19_
                rhs83_ = d_743_v1_
                rhs84_ = (_dafny.SeqWithoutIsStrInference([(p1)[default__.safeIndex(592, len(p1))] for d_769_i3_ in range(default__.abs(207))])).set(default__.safeIndex((self).f11, len(_dafny.SeqWithoutIsStrInference([(p1)[default__.safeIndex(592, len(p1))] for d_770_i3_ in range(default__.abs(207))]))), d_755_v7_)
                lhs53_ = globalState
                lhs54_ = d_751_v4_
                lhs55_ = default__.safeIndex(321, (d_751_v4_).length(0))
                d_763_v16_ = rhs80_
                lhs53_.f3 = rhs81_
                d_768_v19_ = rhs82_
                d_743_v1_ = rhs83_
                lhs54_[lhs55_] = rhs84_
                d_771_v20_: _dafny.Array
                nw128_ = _dafny.Array(int(0), 25)
                d_771_v20_ = nw128_
                r1 = d_771_v20_
                d_743_v1_ = d_743_v1_
            elif True:
                d_772_v21_: C0
                nw129_ = C0()
                nw129_.ctor__(False, False, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_773_i4_ in range(default__.abs(479))]), (self).f10)
                d_772_v21_ = nw129_
                d_774_v22_: _dafny.Map
                d_774_v22_ = _dafny.Map({d_772_v21_: default__.safeDivisionInt(910, p2)})
                d_775_v23_: D14
                d_775_v23_ = D14_DC32(d_772_v21_)
                d_776_v24_: _dafny.Seq
                d_776_v24_ = _dafny.SeqWithoutIsStrInference([d_743_v1_, (d_772_v21_).f26, (d_772_v21_).f25])
                d_774_v22_ = (d_774_v22_).set((d_775_v23_).cf39, len((d_776_v24_) + (d_776_v24_)))
                d_777_v25_: _dafny.Array
                nw130_ = _dafny.Array(_dafny.Seq({}), 10)
                d_777_v25_ = nw130_
                d_778_v26_: _dafny.MultiSet
                d_778_v26_ = _dafny.MultiSet([(d_772_v21_).f26, d_743_v1_, (d_772_v21_).f25, d_743_v1_])
                d_779_v27_: _dafny.Map
                d_779_v27_ = _dafny.Map({d_777_v25_: (d_778_v26_).ispropersubset(d_778_v26_)})
                d_779_v27_ = (d_779_v27_).set(d_777_v25_, d_743_v1_)
                (self).f23 = (((self).f11 if default__.fm38(11, globalState) else (self).f11)) * (len(((self).f9) + ((self).f9)))
                d_780_v28_: str
                d_780_v28_ = _dafny.CodePoint('g')
                d_780_v28_ = d_780_v28_
                d_781_v29_: _dafny.Array
                nw131_ = _dafny.Array(False, 25)
                d_781_v29_ = nw131_
                d_782_v30_: _dafny.Seq
                d_782_v30_ = _dafny.SeqWithoutIsStrInference([d_776_v24_])
                d_783_v31_: _dafny.MultiSet
                d_783_v31_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_743_v1_])])
                index99_ = default__.safeIndex(171, (d_781_v29_).length(0))
                (d_781_v29_)[index99_] = (_dafny.MultiSet(d_782_v30_)).isdisjoint(d_783_v31_)
                index100_ = default__.safeIndex(171, (d_781_v29_).length(0))
                (d_781_v29_)[index100_] = d_743_v1_
            d_743_v1_ = (not(not(not(d_743_v1_)))) or ((p0) == (p2))
            d_784_v32_: _dafny.Map
            d_784_v32_ = _dafny.Map({251: d_743_v1_})
            d_785_v33_: D4
            d_785_v33_ = D4_DC8(d_784_v32_)
            pat_let_tv16_ = d_784_v32_
            d_786_v34_: _dafny.Array
            nw132_ = _dafny.Array(None, 13)
            nw132_[int(0)] = d_785_v33_
            nw132_[int(1)] = d_785_v33_
            def iife73_(_pat_let8_0):
                def iife74_(d_787_dt__update__tmp_h1_):
                    def iife75_(_pat_let9_0):
                        def iife76_(d_788_dt__update_hcf10_h0_):
                            return D4_DC8(d_788_dt__update_hcf10_h0_)
                        return iife76_(_pat_let9_0)
                    return iife75_(pat_let_tv16_)
                return iife74_(_pat_let8_0)
            nw132_[int(2)] = iife73_(D4_DC8(_dafny.Map({p0: d_743_v1_})))
            nw132_[int(3)] = d_785_v33_
            nw132_[int(4)] = d_785_v33_
            nw132_[int(5)] = (d_785_v33_ if True else D4_DC8(_dafny.Map({(self).f11: d_743_v1_})))
            nw132_[int(6)] = d_785_v33_
            nw132_[int(7)] = d_785_v33_
            nw132_[int(8)] = D4_DC8(d_784_v32_)
            nw132_[int(9)] = d_785_v33_
            nw132_[int(10)] = d_785_v33_
            nw132_[int(11)] = d_785_v33_
            nw132_[int(12)] = d_785_v33_
            d_786_v34_ = nw132_
            index101_ = default__.safeIndex(928, (d_786_v34_).length(0))
            (d_786_v34_)[index101_] = d_785_v33_
            d_789_v35_: str
            d_789_v35_ = _dafny.CodePoint('i')
            d_790_v36_: _dafny.MultiSet
            d_790_v36_ = _dafny.MultiSet([d_789_v35_])
            d_791_v37_: _dafny.MultiSet
            d_791_v37_ = _dafny.MultiSet([540, self.f23])
            d_792_v38_: _dafny.Array
            nw133_ = _dafny.Array(None, 8)
            nw133_[int(0)] = d_743_v1_
            nw133_[int(1)] = d_743_v1_
            nw133_[int(2)] = default__.fm38(self.f23, globalState)
            nw133_[int(3)] = d_743_v1_
            nw133_[int(4)] = d_743_v1_
            nw133_[int(5)] = d_743_v1_
            nw133_[int(6)] = d_743_v1_
            nw133_[int(7)] = d_743_v1_
            d_792_v38_ = nw133_
            d_793_v39_: D2
            d_793_v39_ = D2_DC5(d_743_v1_, d_792_v38_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rcp"))), self.f23)
            d_794_v40_: _dafny.Array
            nw134_ = _dafny.Array(None, 19)
            nw134_[int(0)] = not(not(((self).f24) >= (p2)))
            nw134_[int(1)] = d_743_v1_
            nw134_[int(2)] = d_743_v1_
            nw134_[int(3)] = d_743_v1_
            nw134_[int(4)] = d_743_v1_
            nw134_[int(5)] = (d_790_v36_).isdisjoint(d_790_v36_)
            nw134_[int(6)] = (self.f23) not in (d_791_v37_)
            nw134_[int(7)] = True
            nw134_[int(8)] = d_743_v1_
            nw134_[int(9)] = ((self).f10) <= ((self).f10)
            nw134_[int(10)] = True
            nw134_[int(11)] = not(d_743_v1_)
            nw134_[int(12)] = d_743_v1_
            nw134_[int(13)] = d_743_v1_
            nw134_[int(14)] = d_743_v1_
            nw134_[int(15)] = d_743_v1_
            nw134_[int(16)] = d_743_v1_
            nw134_[int(17)] = d_743_v1_
            nw134_[int(18)] = (d_793_v39_).cf4
            d_794_v40_ = nw134_
            index102_ = default__.safeIndex(614, (d_794_v40_).length(0))
            (d_794_v40_)[index102_] = d_743_v1_
            d_795_v41_: _dafny.Seq
            d_795_v41_ = _dafny.SeqWithoutIsStrInference([p0, (self).f24, (self).f24])
            index103_ = default__.safeIndex(928, (d_786_v34_).length(0))
            index104_ = default__.safeIndex(614, (d_794_v40_).length(0))
            rhs85_ = d_785_v33_
            rhs86_ = d_743_v1_
            rhs87_ = p0
            rhs88_ = default__.fm38((0) - (len((d_795_v41_) + (_dafny.SeqWithoutIsStrInference([(self).f24 for d_796_i5_ in range(default__.abs(465))])))), globalState)
            lhs56_ = d_786_v34_
            lhs57_ = default__.safeIndex(928, (d_786_v34_).length(0))
            lhs58_ = self
            lhs59_ = d_794_v40_
            lhs60_ = default__.safeIndex(614, (d_794_v40_).length(0))
            lhs56_[lhs57_] = rhs85_
            d_743_v1_ = rhs86_
            lhs58_.f23 = rhs87_
            lhs59_[lhs60_] = rhs88_
            d_797_v42_: C0
            nw135_ = C0()
            nw135_.ctor__(((d_794_v40_)[default__.safeIndex(614, (d_794_v40_).length(0))] if d_743_v1_ else d_743_v1_), d_743_v1_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqwrup"))) + (p1), default__.fm41(d_743_v1_, (d_794_v40_)[default__.safeIndex(614, (d_794_v40_).length(0))], (d_794_v40_)[default__.safeIndex(614, (d_794_v40_).length(0))], globalState))
            d_797_v42_ = nw135_
            d_798_v43_: D4
            d_798_v43_ = D4_DC11((d_797_v42_).f25, p2)
            (globalState).f3 = (0) - ((d_798_v43_).cf14)
        elif True:
            if d_743_v1_:
                index105_ = default__.safeIndex(194, (d_742_v0_).length(0))
                (d_742_v0_)[index105_] = (self).f24
                index106_ = default__.safeIndex(194, (d_742_v0_).length(0))
                (d_742_v0_)[index106_] = (self).f24
                d_799_v44_: _dafny.Map
                d_799_v44_ = _dafny.Map({(self).f11: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_800_i6_ in range(default__.abs(873))]))})
                d_801_v45_: _dafny.MultiSet
                d_801_v45_ = _dafny.MultiSet([d_799_v44_, d_799_v44_, _dafny.Map({len(_dafny.Set({d_743_v1_})): len((self).f9)})])
                d_802_v46_: _dafny.Map
                d_802_v46_ = _dafny.Map({d_801_v45_: (d_742_v0_)[default__.safeIndex(194, (d_742_v0_).length(0))]})
                d_803_v47_: _dafny.MultiSet
                d_803_v47_ = _dafny.MultiSet([d_799_v44_])
                index107_ = default__.safeIndex(194, (d_742_v0_).length(0))
                (d_742_v0_)[index107_] = ((d_802_v46_)[d_803_v47_] if (d_803_v47_) in (d_802_v46_) else p2)
                d_804_v48_: _dafny.MultiSet
                d_804_v48_ = _dafny.MultiSet([self.f23, p0, self.f23])
                d_805_v49_: _dafny.MultiSet
                d_805_v49_ = _dafny.MultiSet([(d_804_v48_).cardinality, (self).f24, (self).f11, (self).f11, (self).f24])
                d_806_v50_: _dafny.Set
                d_806_v50_ = _dafny.Set({d_743_v1_})
                d_807_v51_: _dafny.Map
                d_807_v51_ = _dafny.Map({True: (d_742_v0_)[default__.safeIndex(194, (d_742_v0_).length(0))]})
                d_808_v52_: _dafny.MultiSet
                d_808_v52_ = _dafny.MultiSet([default__.fm42(d_807_v51_, False, p0, False, globalState)])
                d_809_v53_: _dafny.Seq
                d_809_v53_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_743_v1_: (self).f24}), d_807_v51_])
                d_810_v54_: str
                d_810_v54_ = _dafny.CodePoint('q')
                d_811_v55_: _dafny.Array
                nw136_ = _dafny.Array(None, 21)
                nw136_[int(0)] = 559
                nw136_[int(1)] = (d_804_v48_).cardinality
                nw136_[int(2)] = p2
                nw136_[int(3)] = (self).f24
                nw136_[int(4)] = (self).fm5(len(d_806_v50_), p0, globalState)
                nw136_[int(5)] = p2
                nw136_[int(6)] = (p0) * (p0)
                nw136_[int(7)] = default__.safeDivisionInt(((d_808_v52_)[(d_809_v53_)[default__.safeIndex(len(_dafny.Map({(self).f24: p0})), len(d_809_v53_))]] if ((d_809_v53_)[default__.safeIndex(len(_dafny.Map({(self).f24: p0})), len(d_809_v53_))]) in (d_808_v52_) else (d_742_v0_)[default__.safeIndex(194, (d_742_v0_).length(0))]), (0) - (p2))
                nw136_[int(8)] = self.f23
                nw136_[int(9)] = p2
                nw136_[int(10)] = (self).f24
                nw136_[int(11)] = self.f23
                nw136_[int(12)] = self.f23
                nw136_[int(13)] = (self).f24
                nw136_[int(14)] = p2
                nw136_[int(15)] = ((d_742_v0_)[default__.safeIndex(194, (d_742_v0_).length(0))] if d_743_v1_ else len((p1).set(default__.safeIndex(self.f23, len(p1)), d_810_v54_)))
                nw136_[int(16)] = (0) - ((self).f11)
                nw136_[int(17)] = default__.safeDivisionInt((0) - (p0), (d_742_v0_)[default__.safeIndex(194, (d_742_v0_).length(0))])
                nw136_[int(18)] = (self).f11
                nw136_[int(19)] = (self).f24
                nw136_[int(20)] = (d_742_v0_)[default__.safeIndex(194, (d_742_v0_).length(0))]
                d_811_v55_ = nw136_
                r1 = d_811_v55_
                d_812_v56_: _dafny.Map
                d_812_v56_ = _dafny.Map({d_743_v1_: d_743_v1_})
                d_813_v57_: _dafny.Map
                d_813_v57_ = _dafny.Map({((d_804_v48_)[self.f23] if (self.f23) in (d_804_v48_) else len(p1)): d_743_v1_})
                d_814_v58_: _dafny.Seq
                d_814_v58_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_743_v1_: d_743_v1_}), _dafny.Map({not(d_743_v1_): False})])
                d_815_v59_: _dafny.Map
                d_815_v59_ = _dafny.Map({len(d_813_v57_): (d_814_v58_)[default__.safeIndex(-386, len(d_814_v58_))]})
                (self).f23 = len(((_dafny.Map({324: d_812_v56_})).set(501, d_812_v56_)) | (((default__.fm43(globalState)).set((0) - (len(d_806_v50_)), d_812_v56_)) | (d_815_v59_)))
                d_816_v60_: _dafny.Seq
                d_816_v60_ = _dafny.SeqWithoutIsStrInference([p0])
                index108_ = default__.safeIndex(194, (d_742_v0_).length(0))
                (d_742_v0_)[index108_] = (p2 if default__.fm38((self).f11, globalState) else ((d_816_v60_)[default__.safeIndex(default__.fm0((d_742_v0_)[default__.safeIndex(194, (d_742_v0_).length(0))], (d_742_v0_)[default__.safeIndex(194, (d_742_v0_).length(0))], globalState), len(d_816_v60_))]) - ((self).f24))
            elif True:
                d_817_v61_: _dafny.Array
                nw137_ = _dafny.Array(False, 4)
                d_817_v61_ = nw137_
                index109_ = default__.safeIndex(898, (d_817_v61_).length(0))
                (d_817_v61_)[index109_] = not(d_743_v1_)
                index110_ = default__.safeIndex(898, (d_817_v61_).length(0))
                rhs89_ = p2
                rhs90_ = d_743_v1_
                rhs91_ = (d_817_v61_ if (d_743_v1_ if not(d_743_v1_) else d_743_v1_) else d_817_v61_)
                rhs92_ = self.f23
                lhs61_ = globalState
                lhs62_ = d_817_v61_
                lhs63_ = default__.safeIndex(898, (d_817_v61_).length(0))
                lhs64_ = globalState
                lhs61_.f3 = rhs89_
                lhs62_[lhs63_] = rhs90_
                d_817_v61_ = rhs91_
                lhs64_.f3 = rhs92_
                index111_ = default__.safeIndex(898, (d_817_v61_).length(0))
                (d_817_v61_)[index111_] = (d_817_v61_)[default__.safeIndex(898, (d_817_v61_).length(0))]
                d_818_v62_: _dafny.Seq
                d_818_v62_ = _dafny.SeqWithoutIsStrInference([p2])
                d_819_v63_: D12
                d_819_v63_ = D12_DC29(d_743_v1_, default__.fm0(p2, (self).f24, globalState), d_818_v62_, (d_817_v61_)[default__.safeIndex(898, (d_817_v61_).length(0))], (self).f11)
                d_820_v64_: _dafny.Seq
                d_820_v64_ = _dafny.SeqWithoutIsStrInference([not((d_819_v63_).cf35)])
                rhs93_ = (d_817_v61_)[default__.safeIndex(898, (d_817_v61_).length(0))]
                rhs94_ = (-77) + (self.f23)
                rhs95_ = (d_820_v64_) + ((d_820_v64_ if True else d_820_v64_))
                lhs65_ = globalState
                d_743_v1_ = rhs93_
                lhs65_.f3 = rhs94_
                d_820_v64_ = rhs95_
                d_821_v65_: _dafny.Set
                d_821_v65_ = _dafny.Set({True, d_743_v1_, (d_817_v61_)[default__.safeIndex(898, (d_817_v61_).length(0))]})
                d_822_v66_: _dafny.Array
                nw138_ = _dafny.Array(_dafny.Map({}), 7)
                d_822_v66_ = nw138_
                d_823_v67_: _dafny.Map
                d_823_v67_ = _dafny.Map({d_742_v0_: p1})
                index112_ = default__.safeIndex(237, (d_822_v66_).length(0))
                (d_822_v66_)[index112_] = d_823_v67_
                d_824_v68_: _dafny.Seq
                d_824_v68_ = _dafny.SeqWithoutIsStrInference([d_823_v67_, (d_823_v67_) | (d_823_v67_)])
                index113_ = default__.safeIndex(237, (d_822_v66_).length(0))
                rhs96_ = d_821_v65_
                rhs97_ = (d_824_v68_)[default__.safeIndex(p0, len(d_824_v68_))]
                lhs66_ = d_822_v66_
                lhs67_ = default__.safeIndex(237, (d_822_v66_).length(0))
                d_821_v65_ = rhs96_
                lhs66_[lhs67_] = rhs97_
                d_825_v69_: _dafny.Map
                d_825_v69_ = _dafny.Map({not(not((d_817_v61_)[default__.safeIndex(898, (d_817_v61_).length(0))])): (self).f24})
                (self).f23 = default__.safeModuloInt(((d_825_v69_)[d_743_v1_] if (d_743_v1_) in (d_825_v69_) else self.f23), (self).f11)
            (globalState).f3 = (self).fm5(self.f23, p2, globalState)
            (self).f23 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))
            d_826_v70_: _dafny.MultiSet
            d_826_v70_ = _dafny.MultiSet([p0])
            d_827_v71_: _dafny.Set
            d_827_v71_ = _dafny.Set({d_826_v70_})
            d_828_v72_: D10
            d_828_v72_ = D10_DC23(_dafny.CodePoint('e'), d_827_v71_)
            d_829_v73_: D10
            d_829_v73_ = D10_DC24(d_828_v72_)
            pat_let_tv17_ = d_828_v72_
            def iife77_(_pat_let10_0):
                def iife78_(d_830_dt__update__tmp_h2_):
                    def iife79_(_pat_let11_0):
                        def iife80_(d_831_dt__update_hcf28_h0_):
                            return D10_DC24(d_831_dt__update_hcf28_h0_)
                        return iife80_(_pat_let11_0)
                    return iife79_(pat_let_tv17_)
                return iife78_(_pat_let10_0)
            d_829_v73_ = iife77_(d_829_v73_)
            d_832_v74_: D14
            d_832_v74_ = D14_DC34((0) - (p0), (self).f11, default__.fm0((self).f24, p0, globalState), d_743_v1_)
            d_743_v1_ = (d_832_v74_).cf46
        d_833_v75_: _dafny.Array
        nw139_ = _dafny.Array(_dafny.Array(None, 0), 17)
        d_833_v75_ = nw139_
        r0 = d_833_v75_
        d_834_v76_: _dafny.Map
        d_834_v76_ = _dafny.Map({d_743_v1_: d_742_v0_})
        d_835_v77_: _dafny.MultiSet
        d_835_v77_ = _dafny.MultiSet([((d_834_v76_)[d_743_v1_] if (d_743_v1_) in (d_834_v76_) else d_742_v0_)])
        d_836_v78_: _dafny.Map
        d_836_v78_ = _dafny.Map({p2: d_743_v1_})
        d_837_v79_: _dafny.Set
        d_837_v79_ = _dafny.Set({p2})
        d_838_v80_: _dafny.Set
        d_838_v80_ = _dafny.Set({default__.fm37((self).f11, -876, not(d_743_v1_), self.f23, globalState), p1})
        d_839_v81_: _dafny.Seq
        d_839_v81_ = _dafny.SeqWithoutIsStrInference([len(d_838_v80_), self.f23])
        d_840_v82_: _dafny.Set
        d_840_v82_ = _dafny.Set({d_743_v1_, d_743_v1_})
        d_841_v83_: T0
        nw140_ = C0()
        nw140_.ctor__(d_743_v1_, d_743_v1_, p1, _dafny.SeqWithoutIsStrInference([d_839_v81_ for d_842_i7_ in range(default__.abs(712))]))
        d_841_v83_ = nw140_
        d_843_v84_: _dafny.Map
        d_843_v84_ = _dafny.Map({not(d_743_v1_): False})
        d_844_v85_: _dafny.Set
        d_844_v85_ = _dafny.Set({d_843_v84_, d_843_v84_, default__.fm45(657, _dafny.Map({p1: d_743_v1_}), (self).f24, d_743_v1_, globalState), d_843_v84_})
        nw141_ = _dafny.Array(None, 22)
        nw141_[int(0)] = ((d_835_v77_)[d_742_v0_] if (d_742_v0_) in (d_835_v77_) else (self).f11)
        nw141_[int(1)] = (len(d_836_v78_)) - ((self).f24)
        nw141_[int(2)] = self.f23
        nw141_[int(3)] = len(d_837_v79_)
        nw141_[int(4)] = (self).f11
        nw141_[int(5)] = (self).f24
        nw141_[int(6)] = self.f23
        nw141_[int(7)] = p2
        nw141_[int(8)] = (self).f11
        nw141_[int(9)] = 689
        nw141_[int(10)] = default__.safeDivisionInt((self).f24, (self).f24)
        nw141_[int(11)] = len((default__.fm44(d_743_v1_, len(d_839_v81_), d_743_v1_, globalState)) + (d_839_v81_))
        nw141_[int(12)] = self.f23
        nw141_[int(13)] = len(d_840_v82_)
        nw141_[int(14)] = (self).f11
        nw141_[int(15)] = len(_dafny.Map({len(p1): d_841_v83_}))
        nw141_[int(16)] = -73
        nw141_[int(17)] = default__.safeModuloInt((self).f11, p2)
        nw141_[int(18)] = default__.safeDivisionInt(492, len(d_844_v85_))
        nw141_[int(19)] = 162
        nw141_[int(20)] = default__.safeModuloInt(784, -985)
        nw141_[int(21)] = (d_841_v83_).fm5(56, (0) - (p2), globalState)
        r1 = nw141_
        return r0, r1

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_845_v0_: _dafny.Array
        def lambda32_(d_846_i0_):
            return not((_dafny.CodePoint('k')) in ((self).f9))

        init16_ = lambda32_
        nw142_ = _dafny.Array(None, 16)
        for i0_16_ in range(nw142_.length(0)):
            nw142_[i0_16_] = init16_(i0_16_)
        d_845_v0_ = nw142_
        d_847_v1_: bool
        d_847_v1_ = False
        index114_ = default__.safeIndex(138, (d_845_v0_).length(0))
        (d_845_v0_)[index114_] = d_847_v1_
        index115_ = default__.safeIndex(138, (d_845_v0_).length(0))
        (d_845_v0_)[index115_] = d_847_v1_
        d_848_v2_: _dafny.Map
        d_848_v2_ = _dafny.Map({p2: (self).f11})
        d_849_v3_: _dafny.Seq
        d_849_v3_ = _dafny.SeqWithoutIsStrInference([(d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))], d_847_v1_, (d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))], not((d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))]), d_847_v1_])
        d_850_v4_: _dafny.MultiSet
        d_850_v4_ = _dafny.MultiSet([p2, (self).f24, (self).f24, ((d_848_v2_)[len(d_849_v3_)] if (len(d_849_v3_)) in (d_848_v2_) else len(_dafny.SeqWithoutIsStrInference([p0 for d_851_i2_ in range(default__.abs(532))]))), p0])
        d_852_v5_: _dafny.Set
        d_852_v5_ = _dafny.Set({d_848_v2_})
        d_853_v6_: _dafny.Array
        nw143_ = _dafny.Array(None, 14)
        nw143_[int(0)] = p1
        nw143_[int(1)] = (d_850_v4_).cardinality
        nw143_[int(2)] = self.f23
        nw143_[int(3)] = (d_850_v4_).cardinality
        nw143_[int(4)] = (self).f24
        nw143_[int(5)] = len(d_852_v5_)
        nw143_[int(6)] = ((d_848_v2_)[default__.fm0(783, (self).f24, globalState)] if (default__.fm0(783, (self).f24, globalState)) in (d_848_v2_) else p2)
        nw143_[int(7)] = p0
        nw143_[int(8)] = (self).f11
        nw143_[int(9)] = 27
        nw143_[int(10)] = default__.safeDivisionInt((self).f11, (self).f11)
        nw143_[int(11)] = (self).f24
        nw143_[int(12)] = len((self).f9)
        nw143_[int(13)] = (p1) + (p0)
        d_853_v6_ = nw143_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_853_v6_).length(0)):
            d_854_i1_: int = guard_loop_2_
            if (True) and (((0) <= (d_854_i1_)) and ((d_854_i1_) < ((d_853_v6_).length(0)))):
                (d_853_v6_)[(d_854_i1_)] = default__.safeModuloInt(d_854_i1_, len((_dafny.Map({True: (self).f24})) | (_dafny.Map({(d_849_v3_)[default__.safeIndex(p0, len(d_849_v3_))]: (0) - ((self).f24)}))))
        d_855_v7_: D6
        d_855_v7_ = D6_DC14(d_853_v6_)
        d_856_v8_: _dafny.MultiSet
        d_856_v8_ = _dafny.MultiSet([(d_855_v7_).cf17, d_853_v6_])
        d_857_v9_: D15
        d_857_v9_ = D15_DC35(d_856_v8_)
        if ((d_856_v8_).intersection((d_857_v9_).cf47)).ispropersubset((d_856_v8_) | ((d_856_v8_).set(d_853_v6_, default__.abs(p2)))):
            d_858_v10_: C0
            nw144_ = C0()
            nw144_.ctor__(not((d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))]), d_847_v1_, (self).f9, ((self).f10) + ((self).f10))
            d_858_v10_ = nw144_
            d_855_v7_ = d_855_v7_
            d_859_v11_: _dafny.Map
            d_859_v11_ = _dafny.Map({(d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))]: (self).f24})
            d_859_v11_ = default__.fm42(d_859_v11_, (d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))], p1, (d_858_v10_).f26, globalState)
            if not(d_847_v1_):
                r3 = self.f23
                (globalState).f3 = len((self).f9)
                d_860_v12_: _dafny.MultiSet
                d_860_v12_ = _dafny.MultiSet([(self).f9, (self).f9])
                d_861_v13_: D10
                d_861_v13_ = D10_DC22(d_860_v12_)
                d_861_v13_ = d_861_v13_
                d_862_v14_: str
                d_862_v14_ = _dafny.CodePoint('d')
                d_862_v14_ = _dafny.CodePoint('h')
                d_863_v15_: _dafny.Array
                nw145_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                d_863_v15_ = nw145_
                d_864_v16_: D7
                d_864_v16_ = D7_DC15(d_850_v4_)
                index116_ = default__.safeIndex(90, (d_863_v15_).length(0))
                (d_863_v15_)[index116_] = default__.fm39(d_864_v16_, (d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))], globalState)
                index117_ = default__.safeIndex(90, (d_863_v15_).length(0))
                (d_863_v15_)[index117_] = (d_862_v14_ if (d_858_v10_).f25 else d_862_v14_)
            elif True:
                d_865_v17_: _dafny.Map
                d_865_v17_ = _dafny.Map({d_853_v6_: (self).fm5((self).f24, p2, globalState)})
                d_866_v18_: _dafny.Map
                d_866_v18_ = _dafny.Map({(d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))]: (d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))]})
                (globalState).f3 = (((self).f11) + (len(((d_865_v17_).set(d_853_v6_, len(_dafny.Map({(d_858_v10_).f25: len(d_866_v18_)})))).set(d_853_v6_, 136)))) + (431)
                d_847_v1_ = (d_858_v10_).f26
                (globalState).f3 = ((self).f24) + (p1)
                (globalState).f3 = default__.safeDivisionInt(p0, default__.safeDivisionInt(self.f23, len(d_849_v3_)))
                (self).f23 = self.f23
            d_867_v19_: _dafny.Array
            nw146_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_867_v19_ = nw146_
            d_868_v20_: _dafny.Seq
            d_868_v20_ = _dafny.SeqWithoutIsStrInference([d_867_v19_, d_867_v19_, d_867_v19_, d_867_v19_])
            d_867_v19_ = (d_868_v20_)[default__.safeIndex((self).f11, len(d_868_v20_))]
        elif True:
            r3 = (self).f24
            index118_ = default__.safeIndex(109, (d_853_v6_).length(0))
            (d_853_v6_)[index118_] = (self).f24
            index119_ = default__.safeIndex(109, (d_853_v6_).length(0))
            (d_853_v6_)[index119_] = self.f23
            d_869_v21_: _dafny.Map
            d_869_v21_ = _dafny.Map({p1: (self).f9})
            d_870_v22_: _dafny.Map
            d_870_v22_ = _dafny.Map({(d_849_v3_) == (d_849_v3_): (d_869_v21_).set((self).f11, (self).f9)})
            d_871_v23_: _dafny.Map
            d_871_v23_ = _dafny.Map({d_847_v1_: self.f23})
            d_870_v22_ = (d_870_v22_).set(((d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))]) or ((d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))]), _dafny.Map({len(d_871_v23_): (self).f9}))
            d_872_v24_: _dafny.Array
            def lambda33_(d_873_i3_):
                return _dafny.CodePoint('c')

            init17_ = lambda33_
            nw147_ = _dafny.Array(None, 7)
            for i0_17_ in range(nw147_.length(0)):
                nw147_[i0_17_] = init17_(i0_17_)
            d_872_v24_ = nw147_
            d_874_v25_: str
            d_874_v25_ = _dafny.CodePoint('i')
            index120_ = default__.safeIndex(978, (d_872_v24_).length(0))
            (d_872_v24_)[index120_] = d_874_v25_
            index121_ = default__.safeIndex(978, (d_872_v24_).length(0))
            (d_872_v24_)[index121_] = d_874_v25_
            d_875_v26_: _dafny.Seq
            d_875_v26_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(d_847_v1_)])])
            index122_ = default__.safeIndex(109, (d_853_v6_).length(0))
            (d_853_v6_)[index122_] = (0) - (len(((_dafny.SeqWithoutIsStrInference([d_849_v3_, d_849_v3_, d_849_v3_, default__.fm2(default__.fm37(p0, (d_853_v6_)[default__.safeIndex(109, (d_853_v6_).length(0))], d_847_v1_, p1, globalState), (self).f24, globalState), d_849_v3_])).set(default__.safeIndex(578, len(_dafny.SeqWithoutIsStrInference([d_849_v3_, d_849_v3_, d_849_v3_, default__.fm2(default__.fm37(p0, (d_853_v6_)[default__.safeIndex(109, (d_853_v6_).length(0))], d_847_v1_, p1, globalState), (self).f24, globalState), d_849_v3_]))), d_849_v3_)) + ((d_875_v26_) + (d_875_v26_))))
        d_876_i4_: int
        d_876_i4_ = 0
        with _dafny.label("4"):
            while not ((d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))]) or (d_847_v1_):
                with _dafny.c_label("4"):
                    if (d_876_i4_) >= (100):
                        raise _dafny.Break("4")
                    d_876_i4_ = (d_876_i4_) + (1)
                    r1 = not ((d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))]) or (d_847_v1_)
                    d_877_v27_: _dafny.Seq
                    d_877_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqd"))
                    d_878_v28_: _dafny.Array
                    nw148_ = _dafny.Array(_dafny.Map({}), 14)
                    d_878_v28_ = nw148_
                    d_879_v29_: _dafny.Seq
                    d_879_v29_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f11)])
                    d_880_v30_: _dafny.Map
                    d_880_v30_ = _dafny.Map({len(_dafny.Map({p1: d_847_v1_})): d_879_v29_})
                    index123_ = default__.safeIndex(495, (d_878_v28_).length(0))
                    (d_878_v28_)[index123_] = d_880_v30_
                    index124_ = default__.safeIndex(495, (d_878_v28_).length(0))
                    rhs98_ = (d_877_v27_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxbbyr")))
                    rhs99_ = self.f23
                    rhs100_ = d_880_v30_
                    rhs101_ = d_879_v29_
                    lhs68_ = d_878_v28_
                    lhs69_ = default__.safeIndex(495, (d_878_v28_).length(0))
                    d_877_v27_ = rhs98_
                    r3 = rhs99_
                    lhs68_[lhs69_] = rhs100_
                    d_879_v29_ = rhs101_
                    r3 = (len(d_849_v3_)) + (default__.safeModuloInt(p1, (self).f11))
                    d_881_v31_: _dafny.Map
                    d_881_v31_ = _dafny.Map({189: (d_845_v0_)[default__.safeIndex(138, (d_845_v0_).length(0))]})
                    d_881_v31_ = ((_dafny.Map({self.f23: d_847_v1_})) | (default__.fm46(globalState))).set(p2, not(d_847_v1_))
                    pass
            pass
        r3 = ((0) - (self.f23)) + (self.f23)
        r3 = ((p0) * (len(d_848_v2_)) if ((0) - (p0)) <= ((0) - (p1)) else p0)
        r0 = (self.f23) >= (self.f23)
        r1 = d_847_v1_
        r2 = d_853_v6_
        r3 = default__.safeModuloInt(default__.safeModuloInt(-951, self.f23), p2)
        return r0, r1, r2, r3

    def m17(self, p0, p1, globalState):
        d_882_v0_: bool
        d_882_v0_ = True
        d_883_v1_: _dafny.Set
        d_883_v1_ = _dafny.Set({d_882_v0_, d_882_v0_})
        d_884_v2_: _dafny.Set
        d_884_v2_ = _dafny.Set({len(d_883_v1_)})
        d_885_v3_: D4
        d_885_v3_ = D4_DC11(d_882_v0_, len(d_884_v2_))
        d_886_v4_: _dafny.Map
        d_886_v4_ = _dafny.Map({True: default__.fm38(-248, globalState)})
        d_882_v0_ = not ((d_885_v3_).cf13) or ((-852) in (_dafny.Set({len(d_886_v4_), p1})))
        d_887_v5_: C0
        nw149_ = C0()
        nw149_.ctor__(d_882_v0_, d_882_v0_, (self).f9, (self).f10)
        d_887_v5_ = nw149_
        d_888_v6_: _dafny.Map
        d_888_v6_ = _dafny.Map({self.f23: (d_887_v5_).f26})
        d_889_v7_: D4
        d_889_v7_ = D4_DC8(d_888_v6_)
        d_889_v7_ = d_889_v7_
        d_890_v8_: _dafny.Seq
        d_890_v8_ = _dafny.SeqWithoutIsStrInference([(d_887_v5_).f25])
        d_891_v10_: _dafny.Map
        d_891_v10_ = _dafny.Map({-894: (self).f9})
        d_892_v11_: _dafny.Map
        d_892_v11_ = _dafny.Map({((d_891_v10_)[p1] if (p1) in (d_891_v10_) else (self).f9): (self).f11})
        d_893_v12_: _dafny.Map
        def iife81_():
            coll57_ = _dafny.Map()
            compr_57_: _dafny.Seq
            for compr_57_ in (d_892_v11_).keys.Elements:
                d_894_v9_: _dafny.Seq = compr_57_
                if (d_894_v9_) in (d_892_v11_):
                    coll57_[d_894_v9_] = True
            return _dafny.Map(coll57_)
        d_893_v12_ = _dafny.Map({(d_890_v8_)[default__.safeIndex((self).f11, len(d_890_v8_))]: default__.fm45(p1, iife81_()
        , 52, (d_887_v5_).f25, globalState)})
        d_893_v12_ = (d_893_v12_).set((d_887_v5_).f25, d_886_v4_)
        d_895_v13_: _dafny.Seq
        d_895_v13_ = _dafny.SeqWithoutIsStrInference([(self).f24])
        d_896_v14_: _dafny.Map
        d_896_v14_ = _dafny.Map({-980: p0})
        d_897_v15_: _dafny.Map
        d_897_v15_ = _dafny.Map({d_895_v13_: len(d_896_v14_)})
        d_898_v16_: _dafny.Array
        nw150_ = _dafny.Array(None, 13)
        nw150_[int(0)] = (self).f24
        nw150_[int(1)] = 80
        nw150_[int(2)] = p0
        nw150_[int(3)] = len(_dafny.Map({(d_887_v5_).f25: default__.fm38((self).f24, globalState)}))
        nw150_[int(4)] = self.f23
        nw150_[int(5)] = (self).f11
        nw150_[int(6)] = p0
        nw150_[int(7)] = (self).f24
        nw150_[int(8)] = ((d_897_v15_)[_dafny.SeqWithoutIsStrInference([p0])] if (_dafny.SeqWithoutIsStrInference([p0])) in (d_897_v15_) else self.f23)
        nw150_[int(9)] = len((_dafny.Map({p0: (self).f9})) | (d_891_v10_))
        nw150_[int(10)] = p0
        nw150_[int(11)] = len((self).f9)
        nw150_[int(12)] = (len(d_896_v14_)) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))))
        d_898_v16_ = nw150_
        d_899_v17_: _dafny.Seq
        d_899_v17_ = _dafny.SeqWithoutIsStrInference([d_898_v16_, d_898_v16_, d_898_v16_])
        def iife82_():
            coll58_ = _dafny.Set()
            compr_58_: int
            for compr_58_ in _dafny.IntegerRange(401, 354):
                d_900_v18_: int = compr_58_
                if ((401) <= (d_900_v18_)) and ((d_900_v18_) < (354)):
                    coll58_ = coll58_.union(_dafny.Set([(d_900_v18_) + (self.f23)]))
            return _dafny.Set(coll58_)
        d_898_v16_ = (d_899_v17_)[default__.safeIndex(default__.safeDivisionInt(len(iife82_()
        ), (self).f24), len(d_899_v17_))]
        (globalState).f3 = (self.f23) + (default__.safeModuloInt(len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([len(d_884_v2_) for d_901_i0_ in range(default__.abs(516))])})), (self).f24))

    def m18(self, p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        d_902_v0_: _dafny.MultiSet
        d_902_v0_ = _dafny.MultiSet([(self).f9])
        d_903_v1_: _dafny.Seq
        d_903_v1_ = _dafny.SeqWithoutIsStrInference([self.f23])
        d_904_v2_: C0
        nw151_ = C0()
        nw151_.ctor__((d_902_v0_).ispropersubset(d_902_v0_), p1, (self).f9, ((self).f10).set(default__.safeIndex((self).f11, len((self).f10)), d_903_v1_))
        d_904_v2_ = nw151_
        r2 = p1
        hi1_ = ((self).f11) - (self.f23)
        for d_905_i0_ in range((self).f24, hi1_):
            d_906_v3_: str
            d_906_v3_ = _dafny.CodePoint('o')
            d_906_v3_ = _dafny.CodePoint('r')
            d_907_v4_: _dafny.Seq
            d_907_v4_ = _dafny.SeqWithoutIsStrInference([p1])
            d_908_v5_: _dafny.Map
            d_908_v5_ = _dafny.Map({p1: True})
            d_909_v6_: _dafny.Set
            d_909_v6_ = _dafny.Set({len(d_908_v5_)})
            d_910_v7_: _dafny.Set
            d_910_v7_ = _dafny.Set({len(d_909_v6_)})
            d_911_v8_: D9
            d_911_v8_ = D9_DC21(D3_DC6(d_909_v6_), d_906_v3_, (d_904_v2_).f25)
            d_912_v9_: _dafny.Array
            nw152_ = _dafny.Array(None, 26)
            nw152_[int(0)] = (d_904_v2_).f25
            nw152_[int(1)] = default__.fm38((self).f24, globalState)
            nw152_[int(2)] = (d_904_v2_).f25
            nw152_[int(3)] = (d_904_v2_).f26
            nw152_[int(4)] = not((d_907_v4_)[default__.safeIndex((0) - ((self).f24), len(d_907_v4_))])
            nw152_[int(5)] = (d_904_v2_).f25
            nw152_[int(6)] = (d_904_v2_).f25
            nw152_[int(7)] = (d_904_v2_).f25
            nw152_[int(8)] = (d_904_v2_).f26
            nw152_[int(9)] = False
            nw152_[int(10)] = (d_907_v4_)[default__.safeIndex(self.f23, len(d_907_v4_))]
            nw152_[int(11)] = p1
            nw152_[int(12)] = (d_904_v2_).f25
            nw152_[int(13)] = (d_904_v2_).f25
            nw152_[int(14)] = (d_911_v8_).cf24
            nw152_[int(15)] = (d_904_v2_).f26
            nw152_[int(16)] = (d_904_v2_).f25
            nw152_[int(17)] = (d_904_v2_).f25
            nw152_[int(18)] = (d_904_v2_).f25
            nw152_[int(19)] = (d_904_v2_).f25
            nw152_[int(20)] = False
            nw152_[int(21)] = p1
            nw152_[int(22)] = (d_904_v2_).f26
            nw152_[int(23)] = not((d_904_v2_).f26)
            nw152_[int(24)] = (d_904_v2_).f26
            nw152_[int(25)] = (d_904_v2_).f25
            d_912_v9_ = nw152_
            d_913_v10_: _dafny.Map
            d_913_v10_ = _dafny.Map({d_912_v9_: d_905_i0_})
            r1 = len(d_913_v10_)
            index125_ = default__.safeIndex(345, (d_912_v9_).length(0))
            (d_912_v9_)[index125_] = (d_904_v2_).f26
            d_914_v11_: _dafny.MultiSet
            d_914_v11_ = _dafny.MultiSet([(d_904_v2_).f25])
            d_915_v12_: _dafny.Seq
            d_915_v12_ = _dafny.SeqWithoutIsStrInference([d_914_v11_, _dafny.MultiSet([(d_904_v2_).f26]), d_914_v11_])
            d_916_v13_: _dafny.Map
            d_916_v13_ = _dafny.Map({len(d_915_v12_): p1})
            index126_ = default__.safeIndex(345, (d_912_v9_).length(0))
            rhs102_ = (d_904_v2_).f25
            rhs103_ = (d_904_v2_).f26
            rhs104_ = ((d_916_v13_)[default__.safeDivisionInt((self).f24, self.f23)] if (default__.safeDivisionInt((self).f24, self.f23)) in (d_916_v13_) else (d_904_v2_).f25)
            rhs105_ = (default__.fm0(self.f23, (0) - (d_905_i0_), globalState) if (d_904_v2_).f25 else (self.f23) * ((0) - (self.f23)))
            lhs70_ = d_912_v9_
            lhs71_ = default__.safeIndex(345, (d_912_v9_).length(0))
            lhs72_ = globalState
            r0 = rhs102_
            lhs70_[lhs71_] = rhs103_
            r0 = rhs104_
            lhs72_.f3 = rhs105_
            d_917_v14_: _dafny.Map
            d_917_v14_ = _dafny.Map({(d_904_v2_).f26: d_903_v1_})
            r2 = ((self).f24) != (len(d_917_v14_))
        d_918_v15_: _dafny.Array
        nw153_ = _dafny.Array(None, 1)
        nw153_[int(0)] = _dafny.CodePoint('g')
        d_918_v15_ = nw153_
        d_919_v16_: D6
        d_919_v16_ = D6_DC13(d_918_v15_)
        pat_let_tv18_ = d_918_v15_
        def iife83_(_pat_let12_0):
            def iife84_(d_920_dt__update__tmp_h0_):
                def iife85_(_pat_let13_0):
                    def iife86_(d_921_dt__update_hcf16_h0_):
                        return D6_DC13(d_921_dt__update_hcf16_h0_)
                    return iife86_(_pat_let13_0)
                return iife85_(pat_let_tv18_)
            return iife84_(_pat_let12_0)
        source14_ = iife83_(d_919_v16_)
        if source14_.is_DC14:
            d_922___mcc_h0_ = source14_.cf17
            d_923_cf17_ = d_922___mcc_h0_
            d_924_v17_: _dafny.Map
            d_924_v17_ = _dafny.Map({(d_904_v2_).f26: (self).f24})
            d_925_v18_: C0
            nw154_ = C0()
            nw154_.ctor__((d_904_v2_).f25, (d_904_v2_).f26, default__.fm37(890, (self).f24, p1, (d_904_v2_).fm5(len(d_924_v17_), (0) - ((self).f11), globalState), globalState), (self).f10)
            d_925_v18_ = nw154_
            d_926_v19_: _dafny.Map
            d_926_v19_ = _dafny.Map({self.f23: (self).f24})
            r2 = ((self).f24) > (len(d_926_v19_))
            index127_ = default__.safeIndex(867, (d_923_cf17_).length(0))
            (d_923_cf17_)[index127_] = ((self).f11) * (((d_926_v19_)[len((self).f9)] if (len((self).f9)) in (d_926_v19_) else (self).f24))
            d_927_v20_: _dafny.Array
            nw155_ = _dafny.Array(None, 1)
            nw155_[int(0)] = (d_904_v2_).f26
            d_927_v20_ = nw155_
            index128_ = default__.safeIndex(396, (d_927_v20_).length(0))
            (d_927_v20_)[index128_] = (d_925_v18_).f26
            d_928_v21_: _dafny.Map
            d_928_v21_ = _dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f9])): len(d_903_v1_)})
            index129_ = default__.safeIndex(867, (d_923_cf17_).length(0))
            index130_ = default__.safeIndex(396, (d_927_v20_).length(0))
            rhs106_ = ((d_928_v21_)[d_902_v0_] if (d_902_v0_) in (d_928_v21_) else (self).f11)
            rhs107_ = (self).f11
            rhs108_ = (d_904_v2_).f26
            rhs109_ = ((d_925_v18_).f25 if (d_925_v18_).f26 else default__.fm38((self).f24, globalState))
            rhs110_ = (-16) == (self.f23)
            lhs73_ = d_923_cf17_
            lhs74_ = default__.safeIndex(867, (d_923_cf17_).length(0))
            lhs75_ = d_927_v20_
            lhs76_ = default__.safeIndex(396, (d_927_v20_).length(0))
            r1 = rhs106_
            lhs73_[lhs74_] = rhs107_
            r0 = rhs108_
            r2 = rhs109_
            lhs75_[lhs76_] = rhs110_
            d_929_v22_: str
            d_929_v22_ = _dafny.CodePoint('s')
            d_930_v23_: _dafny.Map
            d_930_v23_ = _dafny.Map({not((d_925_v18_).f25): p1})
            d_931_v24_: _dafny.Map
            d_931_v24_ = _dafny.Map({(self).f9: d_930_v23_})
            rhs111_ = _dafny.CodePoint('l')
            rhs112_ = len((((d_931_v24_)[_dafny.SeqWithoutIsStrInference([d_929_v22_ for d_932_i1_ in range(default__.abs(357))])] if (_dafny.SeqWithoutIsStrInference([d_929_v22_ for d_933_i1_ in range(default__.abs(357))])) in (d_931_v24_) else d_930_v23_)) | ((d_930_v23_) | (_dafny.Map({(d_904_v2_).f26: (d_925_v18_).f25}))))
            lhs77_ = globalState
            d_929_v22_ = rhs111_
            lhs77_.f3 = rhs112_
        elif True:
            d_934___mcc_h1_ = source14_.cf16
            d_935_cf16_ = d_934___mcc_h1_
            d_936_v25_: _dafny.Array
            nw156_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
            d_936_v25_ = nw156_
            index131_ = default__.safeIndex(166, (d_936_v25_).length(0))
            (d_936_v25_)[index131_] = (self).f9
            index132_ = default__.safeIndex(166, (d_936_v25_).length(0))
            (d_936_v25_)[index132_] = (self).f9
            d_937_v26_: C0
            nw157_ = C0()
            nw157_.ctor__(p1, (d_904_v2_).f26, (d_936_v25_)[default__.safeIndex(166, (d_936_v25_).length(0))], ((self).f10) + ((self).f10))
            d_937_v26_ = nw157_
            d_938_v27_: _dafny.Array
            nw158_ = _dafny.Array(False, 26)
            d_938_v27_ = nw158_
            d_938_v27_ = d_938_v27_
            d_939_v28_: _dafny.Map
            d_939_v28_ = _dafny.Map({(d_904_v2_).f26: d_938_v27_})
            d_940_v29_: _dafny.Set
            d_940_v29_ = _dafny.Set({625})
            d_941_v30_: _dafny.MultiSet
            d_941_v30_ = _dafny.MultiSet([len(default__.fm2((d_936_v25_)[default__.safeIndex(166, (d_936_v25_).length(0))], len(d_940_v29_), globalState)), (self).f24, self.f23, (self).f24, len((self).f9)])
            d_939_v28_ = (d_939_v28_).set((default__.fm38((d_941_v30_).cardinality, globalState)) == ((d_904_v2_).f25), d_938_v27_)
        (self).f23 = ((self).f24) - ((self).f24)
        (self).f23 = ((d_903_v1_)[default__.safeIndex(self.f23, len(d_903_v1_))]) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_942_i2_ in range(default__.abs(862))])))
        r0 = not ((d_904_v2_).f25) or (p1)
        r1 = (self).f11
        r2 = ((d_904_v2_).f26 if (d_904_v2_).f25 else (d_904_v2_).f25)
        return r0, r1, r2

    @property
    def f24(self):
        return self._f24

class C2(T1, T0):
    def  __init__(self):
        self._f12: _dafny.Map = _dafny.Map({})
        self._f11: int = int(0)
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f11(self):
        return self._f11
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f11, f12, f9, f10):
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f9 = f9
        (self)._f10 = f10

    def fm5(self, p0, p1, globalState):
        return (0) - ((self).f11)

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: bool = False
        r3: int = int(0)
        d_943_v0_: str
        d_943_v0_ = _dafny.CodePoint('r')
        d_944_v1_: _dafny.MultiSet
        d_944_v1_ = _dafny.MultiSet([d_943_v0_])
        d_945_v2_: _dafny.Set
        d_945_v2_ = _dafny.Set({p0, p0})
        d_946_v3_: D3
        d_946_v3_ = D3_DC6(d_945_v2_)
        d_947_v4_: bool
        d_947_v4_ = True
        d_948_v5_: D9
        d_948_v5_ = D9_DC21(d_946_v3_, d_943_v0_, d_947_v4_)
        d_949_v6_: _dafny.MultiSet
        d_949_v6_ = _dafny.MultiSet([(d_948_v5_).cf24])
        d_950_i0_: int
        d_950_i0_ = 0
        with _dafny.label("5"):
            while (default__.fm34((self).f11, p0, ((d_944_v1_)[d_943_v0_] if (d_943_v0_) in (d_944_v1_) else -847), globalState)).ispropersubset(d_949_v6_):
                with _dafny.c_label("5"):
                    if (d_950_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_950_i0_ = (d_950_i0_) + (1)
                    d_951_v7_: _dafny.Array
                    nw159_ = _dafny.Array(False, 25)
                    d_951_v7_ = nw159_
                    index133_ = default__.safeIndex(492, (d_951_v7_).length(0))
                    (d_951_v7_)[index133_] = d_947_v4_
                    d_952_v8_: _dafny.Set
                    d_952_v8_ = _dafny.Set({(self).f9})
                    index134_ = default__.safeIndex(492, (d_951_v7_).length(0))
                    rhs113_ = default__.safeModuloInt(p0, p0)
                    rhs114_ = d_947_v4_
                    rhs115_ = (d_952_v8_).isdisjoint(d_952_v8_)
                    rhs116_ = d_947_v4_
                    lhs78_ = globalState
                    lhs79_ = d_951_v7_
                    lhs80_ = default__.safeIndex(492, (d_951_v7_).length(0))
                    lhs78_.f3 = rhs113_
                    lhs79_[lhs80_] = rhs114_
                    r2 = rhs115_
                    r2 = rhs116_
                    (globalState).f3 = (self).f11
                    d_953_v9_: _dafny.Array
                    nw160_ = _dafny.Array(D1.default()(), 8)
                    d_953_v9_ = nw160_
                    d_954_v10_: _dafny.Set
                    d_954_v10_ = _dafny.Set({(d_951_v7_)[default__.safeIndex(492, (d_951_v7_).length(0))]})
                    d_955_v11_: _dafny.Seq
                    d_955_v11_ = _dafny.SeqWithoutIsStrInference([d_954_v10_])
                    d_956_v12_: _dafny.Seq
                    d_956_v12_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                    d_957_v13_: _dafny.Set
                    d_957_v13_ = _dafny.Set({d_951_v7_})
                    d_958_v14_: str
                    d_958_v14_ = d_943_v0_
                    d_959_v15_: _dafny.Map
                    d_959_v15_ = _dafny.Map({(self).fm5(p0, len(d_957_v13_), globalState): d_958_v14_})
                    d_960_v16_: _dafny.Map
                    d_960_v16_ = _dafny.Map({(0) - (len((self).f9)): d_959_v15_})
                    d_961_v17_: _dafny.Map
                    d_961_v17_ = _dafny.Map({(d_951_v7_)[default__.safeIndex(492, (d_951_v7_).length(0))]: False})
                    d_962_v18_: _dafny.Array
                    nw161_ = _dafny.Array(None, 7)
                    nw161_[int(0)] = (self).f11
                    nw161_[int(1)] = len((d_955_v11_)[default__.safeIndex(len(d_956_v12_), len(d_955_v11_))])
                    nw161_[int(2)] = (self).f11
                    nw161_[int(3)] = (self).f11
                    nw161_[int(4)] = len(d_960_v16_)
                    nw161_[int(5)] = (self).f11
                    nw161_[int(6)] = len(d_961_v17_)
                    d_962_v18_ = nw161_
                    d_963_v19_: int
                    d_964_v20_: _dafny.MultiSet
                    out43_: int
                    out44_: _dafny.MultiSet
                    out43_, out44_ = (self).m15(d_953_v9_, p0, d_962_v18_, globalState)
                    d_963_v19_ = out43_
                    d_964_v20_ = out44_
                    r0 = 933
                    pass
            pass
        d_965_v21_: _dafny.MultiSet
        d_965_v21_ = _dafny.MultiSet([p0, p0])
        source15_ = D7_DC15(d_965_v21_)
        if source15_.is_DC16:
            d_966_v22_: _dafny.Map
            d_966_v22_ = _dafny.Map({((d_965_v21_)[p0] if (p0) in (d_965_v21_) else p0): p0})
            (globalState).f3 = ((p0) + (len((d_966_v22_).set((self).f11, (self).f11))) if ((self).f11) >= (p0) else (self).f11)
            d_943_v0_ = d_943_v0_
            d_947_v4_ = (d_945_v2_).isdisjoint(d_945_v2_)
            d_967_v23_: _dafny.Array
            nw162_ = _dafny.Array(int(0), 20)
            d_967_v23_ = nw162_
            d_968_v24_: _dafny.Map
            d_968_v24_ = _dafny.Map({p0: d_967_v23_})
            d_969_v25_: D11
            d_969_v25_ = D11_DC25(d_968_v24_)
            d_970_v26_: _dafny.Seq
            d_970_v26_ = _dafny.SeqWithoutIsStrInference([d_968_v24_, d_968_v24_, d_968_v24_])
            d_971_v27_: _dafny.Array
            nw163_ = _dafny.Array(None, 25)
            nw163_[int(0)] = _dafny.Map({p0: d_967_v23_})
            nw163_[int(1)] = ((d_968_v24_).set(p0, d_967_v23_)) | (d_968_v24_)
            nw163_[int(2)] = (d_969_v25_).cf29
            nw163_[int(3)] = d_968_v24_
            nw163_[int(4)] = d_968_v24_
            nw163_[int(5)] = _dafny.Map({p0: d_967_v23_})
            nw163_[int(6)] = d_968_v24_
            nw163_[int(7)] = _dafny.Map({(self).f11: d_967_v23_})
            nw163_[int(8)] = (d_968_v24_).set(p0, d_967_v23_)
            nw163_[int(9)] = (d_968_v24_) | (d_968_v24_)
            nw163_[int(10)] = (d_968_v24_) | (d_968_v24_)
            nw163_[int(11)] = (d_968_v24_) | (d_968_v24_)
            nw163_[int(12)] = _dafny.Map({(self).f11: d_967_v23_})
            nw163_[int(13)] = d_968_v24_
            nw163_[int(14)] = (d_968_v24_ if d_947_v4_ else d_968_v24_)
            nw163_[int(15)] = d_968_v24_
            nw163_[int(16)] = d_968_v24_
            nw163_[int(17)] = d_968_v24_
            nw163_[int(18)] = (d_970_v26_)[default__.safeIndex((_dafny.MultiSet([d_947_v4_, d_947_v4_])).cardinality, len(d_970_v26_))]
            nw163_[int(19)] = d_968_v24_
            nw163_[int(20)] = (d_968_v24_) | (d_968_v24_)
            nw163_[int(21)] = (d_968_v24_ if d_947_v4_ else d_968_v24_)
            nw163_[int(22)] = d_968_v24_
            nw163_[int(23)] = ((d_968_v24_).set(p0, d_967_v23_)).set(len(d_966_v22_), d_967_v23_)
            nw163_[int(24)] = d_968_v24_
            d_971_v27_ = nw163_
            index135_ = default__.safeIndex(142, (d_971_v27_).length(0))
            (d_971_v27_)[index135_] = d_968_v24_
            d_972_v28_: _dafny.Map
            d_972_v28_ = _dafny.Map({d_967_v23_: d_968_v24_})
            index136_ = default__.safeIndex(142, (d_971_v27_).length(0))
            (d_971_v27_)[index136_] = (((d_972_v28_)[d_967_v23_] if (d_967_v23_) in (d_972_v28_) else d_968_v24_)) | (d_968_v24_)
        elif True:
            d_973___mcc_h0_ = source15_.cf18
            d_974_cf18_ = d_973___mcc_h0_
            d_975_v29_: _dafny.Map
            d_975_v29_ = _dafny.Map({default__.fm0(len((self).f9), 62, globalState): (self).f11})
            d_976_v30_: _dafny.Map
            d_976_v30_ = _dafny.Map({not(d_947_v4_): (d_975_v29_) | (d_975_v29_)})
            d_975_v29_ = ((d_976_v30_)[d_947_v4_] if (d_947_v4_) in (d_976_v30_) else d_975_v29_)
            d_977_v31_: _dafny.Seq
            d_977_v31_ = _dafny.SeqWithoutIsStrInference([p0, p0, len(d_975_v29_), 461])
            d_978_v32_: _dafny.Map
            d_978_v32_ = _dafny.Map({d_947_v4_: len(d_977_v31_)})
            d_978_v32_ = _dafny.Map({d_947_v4_: (p0) * ((self).f11)})
            d_979_v33_: D7
            d_979_v33_ = D7_DC16()
            source16_ = d_979_v33_
            if source16_.is_DC16:
                d_980_v34_: _dafny.Array
                nw164_ = _dafny.Array(_dafny.Map({}), 20)
                d_980_v34_ = nw164_
                d_980_v34_ = d_980_v34_
                r3 = ((self).f11) * (414)
                r0 = p0
                d_949_v6_ = (d_949_v6_) | (d_949_v6_)
            elif True:
                d_981___mcc_h1_ = source16_.cf18
                d_982_cf18_ = d_981___mcc_h1_
                d_983_v35_: _dafny.Array
                def lambda34_(d_984_v4_):
                    def lambda35_(d_985_i1_):
                        return D4_DC11(d_984_v4_, (self).f11)

                    return lambda35_

                init18_ = lambda34_(d_947_v4_)
                nw165_ = _dafny.Array(None, 21)
                for i0_18_ in range(nw165_.length(0)):
                    nw165_[i0_18_] = init18_(i0_18_)
                d_983_v35_ = nw165_
                d_986_v36_: _dafny.Seq
                d_986_v36_ = _dafny.SeqWithoutIsStrInference([d_947_v4_, d_947_v4_])
                d_987_v37_: _dafny.Seq
                d_987_v37_ = _dafny.SeqWithoutIsStrInference([(d_986_v36_)[default__.safeIndex(-188, len(d_986_v36_))], d_947_v4_, d_947_v4_, not(d_947_v4_), d_947_v4_])
                d_988_v38_: D4
                d_988_v38_ = D4_DC11(d_947_v4_, (_dafny.MultiSet(d_987_v37_)).cardinality)
                index137_ = default__.safeIndex(194, (d_983_v35_).length(0))
                (d_983_v35_)[index137_] = d_988_v38_
                index138_ = default__.safeIndex(194, (d_983_v35_).length(0))
                (d_983_v35_)[index138_] = d_988_v38_
                d_989_v39_: _dafny.MultiSet
                d_989_v39_ = _dafny.MultiSet([(d_986_v36_ if d_947_v4_ else d_987_v37_)])
                d_990_v40_: _dafny.Array
                nw166_ = _dafny.Array(D3.default()(), 26)
                d_990_v40_ = nw166_
                d_991_v41_: D2
                d_991_v41_ = D2_DC4((0) - (p0))
                d_992_v42_: _dafny.Array
                nw167_ = _dafny.Array(None, 2)
                nw167_[int(0)] = (self).f11
                nw167_[int(1)] = (d_991_v41_).cf3
                d_992_v42_ = nw167_
                index139_ = default__.safeIndex(731, (d_992_v42_).length(0))
                (d_992_v42_)[index139_] = (880) - ((self).f11)
                index140_ = default__.safeIndex(731, (d_992_v42_).length(0))
                rhs117_ = d_989_v39_
                rhs118_ = default__.fm0((self).f11, (self).f11, globalState)
                rhs119_ = d_990_v40_
                rhs120_ = (p0) - ((self).f11)
                lhs81_ = d_992_v42_
                lhs82_ = default__.safeIndex(731, (d_992_v42_).length(0))
                d_989_v39_ = rhs117_
                r3 = rhs118_
                d_990_v40_ = rhs119_
                lhs81_[lhs82_] = rhs120_
                d_993_v43_: C0
                nw168_ = C0()
                nw168_.ctor__(d_947_v4_, d_947_v4_, ((self).f9 if d_947_v4_ else _dafny.SeqWithoutIsStrInference([d_943_v0_ for d_994_i2_ in range(default__.abs(146))])), ((self).f10) + ((self).f10))
                d_993_v43_ = nw168_
                d_995_v44_: _dafny.Map
                d_995_v44_ = _dafny.Map({_dafny.Map({p0: p0}): d_943_v0_})
                index141_ = default__.safeIndex(731, (d_992_v42_).length(0))
                rhs121_ = ((d_995_v44_)[d_975_v29_] if (d_975_v29_) in (d_995_v44_) else d_943_v0_)
                rhs122_ = (p0 if d_947_v4_ else p0)
                rhs123_ = ((d_987_v37_) + (d_987_v37_)) + (_dafny.SeqWithoutIsStrInference([(d_993_v43_).f25, (d_993_v43_).f25]))
                lhs83_ = d_992_v42_
                lhs84_ = default__.safeIndex(731, (d_992_v42_).length(0))
                d_943_v0_ = rhs121_
                lhs83_[lhs84_] = rhs122_
                d_987_v37_ = rhs123_
            d_996_v45_: _dafny.Array
            def lambda36_(d_997_v4_):
                def lambda37_(d_998_i3_):
                    return d_997_v4_

                return lambda37_

            init19_ = lambda36_(d_947_v4_)
            nw169_ = _dafny.Array(None, 19)
            for i0_19_ in range(nw169_.length(0)):
                nw169_[i0_19_] = init19_(i0_19_)
            d_996_v45_ = nw169_
            d_996_v45_ = d_996_v45_
        r2 = d_947_v4_
        d_999_v46_: _dafny.Array
        nw170_ = _dafny.Array(False, 5)
        d_999_v46_ = nw170_
        d_1000_v47_: _dafny.Set
        d_1000_v47_ = _dafny.Set({d_999_v46_, d_999_v46_})
        d_1001_v48_: C1
        nw171_ = C1()
        nw171_.ctor__(len((d_1000_v47_) | (d_1000_v47_)), p0, (self).f11, self.f12, (self).f9, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f11]) for d_1002_i4_ in range(default__.abs(-731))]))
        d_1001_v48_ = nw171_
        d_1003_v49_: D7
        d_1003_v49_ = D7_DC16()
        source17_ = d_1003_v49_
        if source17_.is_DC16:
            d_1004_v50_: _dafny.Map
            d_1004_v50_ = _dafny.Map({d_1001_v48_.f23: d_999_v46_})
            d_1005_v51_: _dafny.MultiSet
            d_1005_v51_ = _dafny.MultiSet([d_999_v46_, ((d_1004_v50_)[d_1001_v48_.f23] if (d_1001_v48_.f23) in (d_1004_v50_) else d_999_v46_), d_999_v46_])
            d_1006_v52_: _dafny.Set
            d_1006_v52_ = _dafny.Set({d_947_v4_, False, d_947_v4_, True})
            d_1007_v53_: _dafny.MultiSet
            d_1007_v53_ = _dafny.MultiSet([d_1006_v52_])
            rhs124_ = ((d_1005_v51_).intersection(d_1005_v51_)) | (d_1005_v51_)
            rhs125_ = (_dafny.MultiSet([d_1006_v52_, _dafny.Set({d_947_v4_}), d_1006_v52_])).ispropersubset(d_1007_v53_)
            rhs126_ = p0
            lhs85_ = globalState
            d_1005_v51_ = rhs124_
            d_947_v4_ = rhs125_
            lhs85_.f3 = rhs126_
            d_1008_v54_: _dafny.Seq
            d_1008_v54_ = _dafny.SeqWithoutIsStrInference([(self).f9])
            d_1009_v55_: _dafny.Seq
            d_1009_v55_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([default__.fm0((d_1001_v48_).f24, 654, globalState)])), d_965_v21_])
            d_1010_v56_: _dafny.Map
            d_1010_v56_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")): ((d_1009_v55_)[default__.safeIndex(d_1001_v48_.f23, len(d_1009_v55_))]).cardinality})
            d_1011_v57_: C1
            nw172_ = C1()
            nw172_.ctor__(len((d_1008_v54_)[default__.safeIndex(d_1001_v48_.f23, len(d_1008_v54_))]), (len(d_1010_v56_) if d_947_v4_ else (d_1001_v48_).f24), len(d_1009_v55_), default__.fm47(d_947_v4_, d_947_v4_, d_947_v4_, d_947_v4_, globalState), ((self).f9) + ((self).f9), default__.fm41(not(d_947_v4_), d_947_v4_, d_947_v4_, globalState))
            d_1011_v57_ = nw172_
            index142_ = default__.safeIndex(298, (d_999_v46_).length(0))
            (d_999_v46_)[index142_] = d_947_v4_
            d_1012_v58_: _dafny.Array
            nw173_ = _dafny.Array(int(0), 5)
            d_1012_v58_ = nw173_
            d_1013_v59_: _dafny.Seq
            d_1013_v59_ = _dafny.SeqWithoutIsStrInference([d_1012_v58_])
            index143_ = default__.safeIndex(298, (d_999_v46_).length(0))
            (d_999_v46_)[index143_] = not(not(not(not(((d_947_v4_ if d_947_v4_ else d_947_v4_) if d_947_v4_ else default__.fm38(len(d_1013_v59_), globalState))))))
            (globalState).f3 = (0) - (((d_949_v6_)[(d_999_v46_)[default__.safeIndex(298, (d_999_v46_).length(0))]] if ((d_999_v46_)[default__.safeIndex(298, (d_999_v46_).length(0))]) in (d_949_v6_) else (d_1011_v57_.f23) - (d_1001_v48_.f23)))
        elif True:
            d_1014___mcc_h2_ = source17_.cf18
            d_1015_cf18_ = d_1014___mcc_h2_
            d_1016_v60_: _dafny.Array
            nw174_ = _dafny.Array(int(0), 24)
            d_1016_v60_ = nw174_
            index144_ = default__.safeIndex(817, (d_1016_v60_).length(0))
            (d_1016_v60_)[index144_] = d_1001_v48_.f23
            index145_ = default__.safeIndex(817, (d_1016_v60_).length(0))
            rhs127_ = d_947_v4_
            rhs128_ = default__.safeDivisionInt(p0, (d_949_v6_).cardinality)
            rhs129_ = d_1001_v48_.f23
            rhs130_ = d_947_v4_
            lhs86_ = d_1001_v48_
            lhs87_ = d_1016_v60_
            lhs88_ = default__.safeIndex(817, (d_1016_v60_).length(0))
            r2 = rhs127_
            lhs86_.f23 = rhs128_
            lhs87_[lhs88_] = rhs129_
            r2 = rhs130_
            d_947_v4_ = default__.fm38((self).f11, globalState)
            index146_ = default__.safeIndex(257, (d_999_v46_).length(0))
            (d_999_v46_)[index146_] = not(d_947_v4_)
            d_1017_v61_: _dafny.Set
            d_1017_v61_ = _dafny.Set({True})
            d_1018_v62_: _dafny.Seq
            d_1018_v62_ = _dafny.SeqWithoutIsStrInference([d_1017_v61_])
            d_1019_v63_: _dafny.Map
            d_1019_v63_ = _dafny.Map({True: (d_1001_v48_).f24})
            index147_ = default__.safeIndex(257, (d_999_v46_).length(0))
            rhs131_ = d_947_v4_
            rhs132_ = default__.safeModuloInt((0) - ((0) - ((0) - ((d_1001_v48_).f24))), ((d_1001_v48_).f24) * (len((d_1018_v62_)[default__.safeIndex((d_1016_v60_)[default__.safeIndex(817, (d_1016_v60_).length(0))], len(d_1018_v62_))])))
            rhs133_ = d_947_v4_
            rhs134_ = len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(d_1019_v63_)]))]))
            lhs89_ = d_1001_v48_
            lhs90_ = d_999_v46_
            lhs91_ = default__.safeIndex(257, (d_999_v46_).length(0))
            lhs92_ = globalState
            r2 = rhs131_
            lhs89_.f23 = rhs132_
            lhs90_[lhs91_] = rhs133_
            lhs92_.f3 = rhs134_
            r2 = ((d_1001_v48_).f24) >= ((d_1001_v48_).f24)
        d_999_v46_ = d_999_v46_
        r0 = p0
        d_1020_v64_: _dafny.Map
        d_1020_v64_ = _dafny.Map({d_947_v4_: ((self).f9 if d_947_v4_ else (self).f9)})
        r1 = d_1020_v64_
        r2 = not(d_947_v4_)
        d_1021_v65_: _dafny.Map
        d_1021_v65_ = _dafny.Map({d_947_v4_: d_947_v4_})
        d_1022_v66_: _dafny.Map
        d_1022_v66_ = _dafny.Map({(self).f11: ((d_1021_v65_)[d_947_v4_] if (d_947_v4_) in (d_1021_v65_) else d_947_v4_)})
        r3 = len(d_1022_v66_)
        return r0, r1, r2, r3

    def m3(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_1023_v0_: bool
        d_1023_v0_ = True
        d_1023_v0_ = not((p1) <= ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1024_i0_ in range(default__.abs(-862))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1025_i1_ in range(default__.abs(-667))]))))
        d_1026_i2_: int
        d_1026_i2_ = 0
        with _dafny.label("6"):
            while d_1023_v0_:
                with _dafny.c_label("6"):
                    if (d_1026_i2_) >= (100):
                        raise _dafny.Break("6")
                    d_1026_i2_ = (d_1026_i2_) + (1)
                    d_1027_v1_: _dafny.MultiSet
                    d_1027_v1_ = _dafny.MultiSet([d_1023_v0_, d_1023_v0_])
                    d_1023_v0_ = not((False) in (d_1027_v1_))
                    d_1028_v2_: D1
                    d_1028_v2_ = D1_DC1(d_1023_v0_)
                    source18_ = d_1028_v2_
                    if source18_.is_DC2:
                        d_1029_v3_: _dafny.Map
                        d_1029_v3_ = _dafny.Map({(0) - (p0): p2})
                        d_1029_v3_ = (d_1029_v3_).set((0) - (p2), (self).f11)
                        d_1023_v0_ = d_1023_v0_
                        d_1030_v4_: _dafny.MultiSet
                        d_1030_v4_ = _dafny.MultiSet([471])
                        (globalState).f3 = ((0) - (p2)) * (default__.safeModuloInt((d_1030_v4_).cardinality, p0))
                        d_1031_v5_: str
                        d_1031_v5_ = _dafny.CodePoint('k')
                        d_1032_v6_: _dafny.Map
                        d_1032_v6_ = _dafny.Map({_dafny.Set({len(_dafny.Set({d_1031_v5_, d_1031_v5_, (p1)[default__.safeIndex(len((self).f9), len(p1))]}))}): _dafny.Set({not(not(d_1023_v0_))})})
                        d_1032_v6_ = ((d_1032_v6_) | (d_1032_v6_)) | (d_1032_v6_)
                    elif source18_.is_DC1:
                        d_1033___mcc_h0_ = source18_.cf1
                        d_1034_cf1_ = d_1033___mcc_h0_
                        d_1023_v0_ = d_1034_cf1_
                        d_1023_v0_ = (not(False)) == ((d_1034_cf1_ if default__.fm38(p0, globalState) else d_1023_v0_))
                        (globalState).f3 = p0
                        d_1035_v7_: _dafny.Map
                        d_1035_v7_ = _dafny.Map({not(default__.fm38(len(p1), globalState)): default__.fm37(((d_1027_v1_)[d_1023_v0_] if (d_1023_v0_) in (d_1027_v1_) else (self).f11), (d_1027_v1_).cardinality, False, p0, globalState)})
                        d_1035_v7_ = d_1035_v7_
                    elif True:
                        d_1036___mcc_h1_ = source18_.cf2
                        d_1037_cf2_ = d_1036___mcc_h1_
                        d_1023_v0_ = d_1023_v0_
                        d_1038_v8_: _dafny.Array
                        nw175_ = _dafny.Array(None, 15)
                        nw175_[int(0)] = p0
                        nw175_[int(1)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))
                        nw175_[int(2)] = p2
                        nw175_[int(3)] = p0
                        nw175_[int(4)] = p0
                        nw175_[int(5)] = -770
                        nw175_[int(6)] = p0
                        nw175_[int(7)] = p2
                        nw175_[int(8)] = (self).f11
                        nw175_[int(9)] = p2
                        nw175_[int(10)] = p0
                        nw175_[int(11)] = default__.fm0((self).f11, p2, globalState)
                        nw175_[int(12)] = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1039_i3_ in range(default__.abs(-572))])))
                        nw175_[int(13)] = (self).f11
                        nw175_[int(14)] = (self).f11
                        d_1038_v8_ = nw175_
                        d_1040_v9_: _dafny.Set
                        d_1040_v9_ = _dafny.Set({225, 659})
                        d_1041_v10_: _dafny.Map
                        d_1041_v10_ = _dafny.Map({d_1040_v9_: d_1038_v8_})
                        d_1042_v11_: _dafny.Seq
                        d_1042_v11_ = _dafny.SeqWithoutIsStrInference([d_1038_v8_, d_1038_v8_])
                        d_1043_v12_: _dafny.Array
                        nw176_ = _dafny.Array(None, 25)
                        nw176_[int(0)] = d_1038_v8_
                        nw176_[int(1)] = d_1038_v8_
                        nw176_[int(2)] = ((d_1041_v10_)[default__.fm48(d_1023_v0_, p0, 24, d_1023_v0_, globalState)] if (default__.fm48(d_1023_v0_, p0, 24, d_1023_v0_, globalState)) in (d_1041_v10_) else d_1038_v8_)
                        nw176_[int(3)] = d_1038_v8_
                        nw176_[int(4)] = d_1038_v8_
                        nw176_[int(5)] = d_1038_v8_
                        nw176_[int(6)] = d_1038_v8_
                        nw176_[int(7)] = d_1038_v8_
                        nw176_[int(8)] = d_1038_v8_
                        nw176_[int(9)] = d_1038_v8_
                        nw176_[int(10)] = d_1038_v8_
                        nw176_[int(11)] = (d_1042_v11_)[default__.safeIndex(p0, len(d_1042_v11_))]
                        nw176_[int(12)] = d_1038_v8_
                        nw176_[int(13)] = d_1038_v8_
                        nw176_[int(14)] = d_1038_v8_
                        nw176_[int(15)] = d_1038_v8_
                        nw176_[int(16)] = d_1038_v8_
                        nw176_[int(17)] = d_1038_v8_
                        nw176_[int(18)] = d_1038_v8_
                        nw176_[int(19)] = d_1038_v8_
                        nw176_[int(20)] = d_1038_v8_
                        nw176_[int(21)] = d_1038_v8_
                        nw176_[int(22)] = d_1038_v8_
                        nw176_[int(23)] = d_1038_v8_
                        nw176_[int(24)] = d_1038_v8_
                        d_1043_v12_ = nw176_
                        d_1043_v12_ = d_1043_v12_
                        d_1044_v13_: _dafny.MultiSet
                        d_1044_v13_ = _dafny.MultiSet([(self).f11, p2])
                        d_1045_v14_: D7
                        d_1045_v14_ = D7_DC15(d_1044_v13_)
                        d_1044_v13_ = (d_1045_v14_).cf18
                        d_1023_v0_ = d_1023_v0_
                    d_1046_v15_: _dafny.Map
                    d_1046_v15_ = _dafny.Map({((self).f11) + (828): d_1023_v0_})
                    d_1046_v15_ = (d_1046_v15_).set(p0, d_1023_v0_)
                    d_1023_v0_ = not(d_1023_v0_)
                    pass
            pass
        d_1047_v16_: _dafny.Array
        def lambda38_(d_1048_v0_):
            def lambda39_(d_1049_i4_):
                return _dafny.SeqWithoutIsStrInference([d_1048_v0_, not(d_1048_v0_)])

            return lambda39_

        init20_ = lambda38_(d_1023_v0_)
        nw177_ = _dafny.Array(None, 15)
        for i0_20_ in range(nw177_.length(0)):
            nw177_[i0_20_] = init20_(i0_20_)
        d_1047_v16_ = nw177_
        d_1050_v17_: _dafny.MultiSet
        d_1050_v17_ = _dafny.MultiSet([(self).f11])
        d_1051_v18_: _dafny.Seq
        d_1051_v18_ = _dafny.SeqWithoutIsStrInference([((d_1050_v17_)[len((self).f9)] if (len((self).f9)) in (d_1050_v17_) else p0), p2, p0])
        d_1052_v19_: _dafny.Set
        d_1052_v19_ = _dafny.Set({p2, 96})
        d_1053_v20_: _dafny.Map
        d_1053_v20_ = _dafny.Map({-720: (len(d_1052_v19_)) - ((0) - ((self).f11))})
        d_1054_v21_: _dafny.MultiSet
        d_1054_v21_ = _dafny.MultiSet([True])
        d_1055_v22_: _dafny.Map
        d_1055_v22_ = _dafny.Map({p0: d_1054_v21_})
        rhs135_ = d_1047_v16_
        rhs136_ = _dafny.SeqWithoutIsStrInference([p2, (self).f11, (((d_1055_v22_)[(self).f11] if ((self).f11) in (d_1055_v22_) else d_1054_v21_)).cardinality])
        rhs137_ = (d_1053_v20_).set(p2, (self).f11)
        rhs138_ = (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1056_i5_ in range(default__.abs(523))]))) - (p2)
        lhs93_ = globalState
        d_1047_v16_ = rhs135_
        d_1051_v18_ = rhs136_
        d_1053_v20_ = rhs137_
        lhs93_.f3 = rhs138_
        d_1057_v23_: _dafny.Array
        def lambda40_(d_1058_v0_):
            def lambda41_(d_1059_i6_):
                return d_1058_v0_

            return lambda41_

        init21_ = lambda40_(d_1023_v0_)
        nw178_ = _dafny.Array(None, 27)
        for i0_21_ in range(nw178_.length(0)):
            nw178_[i0_21_] = init21_(i0_21_)
        d_1057_v23_ = nw178_
        d_1060_v24_: _dafny.Map
        d_1060_v24_ = _dafny.Map({d_1057_v23_: default__.safeDivisionInt(709, (self).f11)})
        d_1060_v24_ = (d_1060_v24_).set(d_1057_v23_, (self).f11)
        d_1061_v26_: _dafny.Seq
        def iife87_():
            coll59_ = _dafny.Map()
            compr_59_: int
            for compr_59_ in _dafny.IntegerRange(678, 989):
                d_1062_v25_: int = compr_59_
                if ((678) <= (d_1062_v25_)) and ((d_1062_v25_) < (989)):
                    coll59_[(d_1062_v25_) + ((0) - (p2))] = p2
            return _dafny.Map(coll59_)
        d_1061_v26_ = _dafny.SeqWithoutIsStrInference([iife87_()
        , d_1053_v20_])
        d_1063_v27_: _dafny.Seq
        d_1063_v27_ = _dafny.SeqWithoutIsStrInference([(d_1053_v20_) | (d_1053_v20_), (d_1061_v26_)[default__.safeIndex((0) - ((self).f11), len(d_1061_v26_))]])
        d_1061_v26_ = d_1061_v26_
        d_1064_i7_: int
        d_1064_i7_ = 0
        with _dafny.label("7"):
            while (default__.fm0(p2, p0, globalState)) > (p2):
                with _dafny.c_label("7"):
                    if (d_1064_i7_) >= (100):
                        raise _dafny.Break("7")
                    d_1064_i7_ = (d_1064_i7_) + (1)
                    d_1065_v28_: D7
                    d_1065_v28_ = D7_DC15(d_1050_v17_)
                    d_1066_v29_: _dafny.Array
                    nw179_ = _dafny.Array(None, 4)
                    nw179_[int(0)] = d_1065_v28_
                    nw179_[int(1)] = d_1065_v28_
                    nw179_[int(2)] = (d_1065_v28_ if not(d_1023_v0_) else d_1065_v28_)
                    nw179_[int(3)] = d_1065_v28_
                    d_1066_v29_ = nw179_
                    index148_ = default__.safeIndex(85, (d_1066_v29_).length(0))
                    (d_1066_v29_)[index148_] = d_1065_v28_
                    index149_ = default__.safeIndex(85, (d_1066_v29_).length(0))
                    rhs139_ = d_1065_v28_
                    rhs140_ = d_1023_v0_
                    rhs141_ = (default__.fm39(D7_DC15(d_1050_v17_), d_1023_v0_, globalState)) in ((self).f9)
                    lhs94_ = d_1066_v29_
                    lhs95_ = default__.safeIndex(85, (d_1066_v29_).length(0))
                    lhs94_[lhs95_] = rhs139_
                    d_1023_v0_ = rhs140_
                    d_1023_v0_ = rhs141_
                    d_1067_v30_: _dafny.Map
                    d_1067_v30_ = _dafny.Map({-535: p1})
                    d_1068_v31_: _dafny.Set
                    d_1068_v31_ = _dafny.Set({((d_1067_v30_)[(self).f11] if ((self).f11) in (d_1067_v30_) else p1), (p1) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdxdlg"))), p1, (self).f9, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmdg"))})
                    d_1069_v33_: _dafny.Map
                    d_1069_v33_ = _dafny.Map({(self).f9: p2})
                    def iife88_():
                        coll60_ = _dafny.Set()
                        compr_60_: _dafny.Seq
                        for compr_60_ in (default__.fm49(d_1023_v0_, d_1023_v0_, d_1023_v0_, globalState)).keys.Elements:
                            d_1070_v32_: _dafny.Seq = compr_60_
                            if (d_1070_v32_) in (default__.fm49(d_1023_v0_, d_1023_v0_, d_1023_v0_, globalState)):
                                coll60_ = coll60_.union(_dafny.Set([d_1070_v32_]))
                        return _dafny.Set(coll60_)
                    rhs142_ = ((d_1068_v31_).intersection(d_1068_v31_)) - ((iife88_()
                    ) - (d_1068_v31_))
                    rhs143_ = True
                    rhs144_ = default__.safeDivisionInt((self).f11, len(d_1069_v33_))
                    lhs96_ = globalState
                    d_1068_v31_ = rhs142_
                    d_1023_v0_ = rhs143_
                    lhs96_.f3 = rhs144_
                    d_1071_v34_: _dafny.Map
                    d_1071_v34_ = _dafny.Map({d_1023_v0_: p0})
                    d_1072_v35_: _dafny.Map
                    d_1072_v35_ = _dafny.Map({d_1071_v34_: p2})
                    (globalState).f3 = ((d_1072_v35_)[d_1071_v34_] if (d_1071_v34_) in (d_1072_v35_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvmcg"))))
                    (globalState).f3 = len((self).f9)
                    pass
            pass
        d_1073_v36_: _dafny.Array
        nw180_ = _dafny.Array(None, 10)
        nw180_[int(0)] = d_1057_v23_
        nw180_[int(1)] = d_1057_v23_
        nw180_[int(2)] = d_1057_v23_
        nw180_[int(3)] = d_1057_v23_
        nw180_[int(4)] = d_1057_v23_
        nw180_[int(5)] = d_1057_v23_
        nw180_[int(6)] = d_1057_v23_
        nw180_[int(7)] = d_1057_v23_
        nw180_[int(8)] = (d_1057_v23_ if d_1023_v0_ else d_1057_v23_)
        nw180_[int(9)] = d_1057_v23_
        d_1073_v36_ = nw180_
        r0 = d_1073_v36_
        d_1074_v37_: _dafny.Array
        def lambda42_(d_1075_p2_):
            def lambda43_(d_1076_i8_):
                return (d_1076_i8_) * (d_1075_p2_)

            return lambda43_

        init22_ = lambda42_(p2)
        nw181_ = _dafny.Array(None, 24)
        for i0_22_ in range(nw181_.length(0)):
            nw181_[i0_22_] = init22_(i0_22_)
        d_1074_v37_ = nw181_
        r1 = d_1074_v37_
        return r0, r1

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_1077_v0_: bool
        d_1077_v0_ = False
        d_1078_v1_: str
        d_1078_v1_ = _dafny.CodePoint('r')
        d_1079_v2_: C1
        nw182_ = C1()
        nw182_.ctor__((p2 if d_1077_v0_ else p2), ((self).f11) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofpnkf")))), (self).f11, self.f12, ((self).f9).set(default__.safeIndex(p0, len((self).f9)), d_1078_v1_), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p2 for d_1080_i1_ in range(default__.abs(219))]) for d_1081_i0_ in range(default__.abs(98))]))
        d_1079_v2_ = nw182_
        d_1082_v3_: _dafny.Array
        def lambda44_(d_1083_v2_):
            def lambda45_(d_1084_i2_):
                return (d_1084_i2_) + (d_1083_v2_.f23)

            return lambda45_

        init23_ = lambda44_(d_1079_v2_)
        nw183_ = _dafny.Array(None, 21)
        for i0_23_ in range(nw183_.length(0)):
            nw183_[i0_23_] = init23_(i0_23_)
        d_1082_v3_ = nw183_
        d_1085_v4_: _dafny.MultiSet
        d_1085_v4_ = _dafny.MultiSet([d_1082_v3_])
        d_1086_v5_: _dafny.Map
        d_1086_v5_ = _dafny.Map({(self).f9: d_1085_v4_})
        d_1087_v6_: D15
        d_1087_v6_ = D15_DC35(((d_1086_v5_)[(self).f9] if ((self).f9) in (d_1086_v5_) else d_1085_v4_))
        source19_ = d_1087_v6_
        if source19_.is_DC36:
            d_1088___mcc_h0_ = source19_.cf48
            d_1089_cf48_ = d_1088___mcc_h0_
            d_1090_v7_: C1
            nw184_ = C1()
            nw184_.ctor__(((0) - ((d_1079_v2_).f24)) + ((0) - ((d_1079_v2_).f24)), (0) - (default__.safeModuloInt(d_1079_v2_.f23, len((self).f9))), 907, self.f12, (self).f9, (self).f10)
            d_1090_v7_ = nw184_
            r0 = True
            d_1091_v8_: _dafny.Seq
            d_1091_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pruulrdgp"))
            d_1091_v8_ = (self).f9
            d_1092_v9_: _dafny.Map
            d_1092_v9_ = _dafny.Map({not (not(not(True))) or (False): (d_1091_v8_) + (d_1091_v8_)})
            d_1093_v10_: _dafny.MultiSet
            d_1093_v10_ = _dafny.MultiSet([(d_1090_v7_).f24])
            d_1094_v11_: D7
            d_1094_v11_ = D7_DC15((d_1093_v10_).set((self).f11, default__.abs(p2)))
            d_1095_v12_: _dafny.Map
            d_1095_v12_ = _dafny.Map({default__.fm39(d_1094_v11_, d_1077_v0_, globalState): (self).f9})
            d_1092_v9_ = (d_1092_v9_).set(d_1077_v0_, (((d_1095_v12_)[d_1078_v1_] if (d_1078_v1_) in (d_1095_v12_) else d_1091_v8_) if d_1077_v0_ else d_1091_v8_))
        elif source19_.is_DC37:
            d_1096___mcc_h1_ = source19_.cf49
            d_1097___mcc_h2_ = source19_.cf50
            d_1098___mcc_h3_ = source19_.cf51
            d_1099_cf51_ = d_1098___mcc_h3_
            d_1100_cf50_ = d_1097___mcc_h2_
            d_1101_cf49_ = d_1096___mcc_h1_
            d_1102_v13_: _dafny.Array
            nw185_ = _dafny.Array(False, 18)
            d_1102_v13_ = nw185_
            index150_ = default__.safeIndex(446, (d_1102_v13_).length(0))
            (d_1102_v13_)[index150_] = d_1099_cf51_
            d_1103_v14_: _dafny.Seq
            d_1103_v14_ = _dafny.SeqWithoutIsStrInference([d_1077_v0_])
            d_1104_v15_: _dafny.Seq
            d_1104_v15_ = _dafny.SeqWithoutIsStrInference([d_1099_cf51_, (d_1103_v14_)[default__.safeIndex(487, len(d_1103_v14_))]])
            d_1105_v16_: _dafny.Map
            d_1105_v16_ = _dafny.Map({(0) - (d_1079_v2_.f23): 998})
            d_1106_v17_: _dafny.Set
            d_1106_v17_ = _dafny.Set({len(_dafny.Map({d_1099_cf51_: _dafny.CodePoint('r')})), d_1079_v2_.f23})
            d_1107_v18_: D3
            d_1107_v18_ = D3_DC6(d_1106_v17_)
            d_1108_v19_: D9
            d_1108_v19_ = D9_DC21(d_1107_v18_, d_1078_v1_, d_1099_cf51_)
            index151_ = default__.safeIndex(446, (d_1102_v13_).length(0))
            (d_1102_v13_)[index151_] = (((d_1103_v14_)[default__.safeIndex(((d_1105_v16_)[(d_1079_v2_).f24] if ((d_1079_v2_).f24) in (d_1105_v16_) else (self).f11), len(d_1103_v14_))]) and (d_1077_v0_) if True else (d_1108_v19_).cf24)
            d_1109_v20_: _dafny.Seq
            d_1109_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adaxfjygr"))
            d_1109_v20_ = (d_1109_v20_) + ((self).f9)
            (globalState).f3 = (0) - (-154)
            r1 = not(((d_1079_v2_).f24) > ((p1) * (p2)))
        elif True:
            d_1110___mcc_h4_ = source19_.cf47
            d_1111_cf47_ = d_1110___mcc_h4_
            index152_ = default__.safeIndex(326, (d_1082_v3_).length(0))
            (d_1082_v3_)[index152_] = default__.safeModuloInt(592, (d_1079_v2_).fm5(p0, p0, globalState))
            index153_ = default__.safeIndex(326, (d_1082_v3_).length(0))
            (d_1082_v3_)[index153_] = p0
            d_1112_v21_: _dafny.MultiSet
            d_1112_v21_ = _dafny.MultiSet([True, True])
            d_1113_v22_: _dafny.MultiSet
            d_1113_v22_ = _dafny.MultiSet([d_1112_v21_, d_1112_v21_, d_1112_v21_])
            d_1114_v23_: _dafny.Map
            d_1114_v23_ = _dafny.Map({True: d_1113_v22_})
            d_1115_v24_: _dafny.Seq
            d_1115_v24_ = _dafny.SeqWithoutIsStrInference([d_1112_v21_])
            d_1114_v23_ = (d_1114_v23_).set(d_1077_v0_, _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1077_v0_]) for d_1116_i3_ in range(default__.abs(240))])) + (d_1115_v24_)))
            (globalState).f3 = (d_1079_v2_).f24
            index154_ = default__.safeIndex(326, (d_1082_v3_).length(0))
            (d_1082_v3_)[index154_] = default__.safeDivisionInt(d_1079_v2_.f23, (self).f11)
        d_1117_v25_: _dafny.Array
        nw186_ = _dafny.Array(False, 9)
        d_1117_v25_ = nw186_
        index155_ = default__.safeIndex(511, (d_1117_v25_).length(0))
        (d_1117_v25_)[index155_] = (d_1077_v0_) and (d_1077_v0_)
        index156_ = default__.safeIndex(511, (d_1117_v25_).length(0))
        (d_1117_v25_)[index156_] = False
        r3 = ((0) - ((self).f11)) + (-872)
        d_1118_v26_: _dafny.MultiSet
        d_1118_v26_ = _dafny.MultiSet([d_1077_v0_, d_1077_v0_, (d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))], True])
        d_1119_v27_: _dafny.Map
        d_1119_v27_ = _dafny.Map({(d_1118_v26_).ispropersubset((d_1118_v26_).set(d_1077_v0_, default__.abs(p2))): (d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]})
        if ((d_1119_v27_)[(d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]] if ((d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]) in (d_1119_v27_) else d_1077_v0_):
            d_1120_v28_: _dafny.Seq
            d_1120_v28_ = _dafny.SeqWithoutIsStrInference([p0, p2])
            (d_1079_v2_).f23 = len(((self).f9).set(default__.safeIndex((p0) - (len((d_1120_v28_).set(default__.safeIndex((d_1079_v2_).f24, len(d_1120_v28_)), p1))), len((self).f9)), d_1078_v1_))
            d_1121_v29_: D6
            d_1121_v29_ = D6_DC14(d_1082_v3_)
            d_1121_v29_ = d_1121_v29_
            d_1122_v30_: _dafny.Map
            d_1122_v30_ = _dafny.Map({d_1078_v1_: p2})
            d_1123_v31_: _dafny.Set
            d_1123_v31_ = _dafny.Set({len((self).f9), d_1079_v2_.f23})
            d_1124_v32_: _dafny.Map
            d_1124_v32_ = _dafny.Map({(0) - ((self).fm5(len(d_1123_v31_), p0, globalState)): not(d_1077_v0_)})
            index157_ = default__.safeIndex(511, (d_1117_v25_).length(0))
            rhs145_ = True
            rhs146_ = ((d_1124_v32_)[(d_1079_v2_).f24] if ((d_1079_v2_).f24) in (d_1124_v32_) else (d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))])
            rhs147_ = d_1120_v28_
            rhs148_ = d_1082_v3_
            rhs149_ = d_1122_v30_
            lhs97_ = d_1117_v25_
            lhs98_ = default__.safeIndex(511, (d_1117_v25_).length(0))
            d_1077_v0_ = rhs145_
            lhs97_[lhs98_] = rhs146_
            d_1120_v28_ = rhs147_
            r2 = rhs148_
            d_1122_v30_ = rhs149_
            d_1125_v33_: C0
            nw187_ = C0()
            nw187_.ctor__((d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))], d_1077_v0_, ((self).f9) + ((self).f9), (self).f10)
            d_1125_v33_ = nw187_
            r3 = (d_1079_v2_).fm5(20, p2, globalState)
        elif True:
            index158_ = default__.safeIndex(302, (d_1082_v3_).length(0))
            (d_1082_v3_)[index158_] = (p0) + ((self).f11)
            d_1126_v34_: _dafny.Set
            d_1126_v34_ = _dafny.Set({(self).f9, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hubpie"))})
            d_1127_v35_: _dafny.Set
            d_1127_v35_ = _dafny.Set({p2})
            d_1128_v36_: _dafny.Map
            d_1128_v36_ = _dafny.Map({(d_1126_v34_).intersection(_dafny.Set({(self).f9, (self).f9, (self).f9})): ((0) - (len(d_1127_v35_))) * (d_1079_v2_.f23)})
            index159_ = default__.safeIndex(302, (d_1082_v3_).length(0))
            rhs150_ = ((d_1128_v36_)[(d_1126_v34_) - (d_1126_v34_)] if ((d_1126_v34_) - (d_1126_v34_)) in (d_1128_v36_) else 149)
            rhs151_ = (d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]
            lhs99_ = d_1082_v3_
            lhs100_ = default__.safeIndex(302, (d_1082_v3_).length(0))
            lhs99_[lhs100_] = rhs150_
            d_1077_v0_ = rhs151_
            index160_ = default__.safeIndex(511, (d_1117_v25_).length(0))
            (d_1117_v25_)[index160_] = not ((d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]) or ((d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))])
            d_1129_v37_: _dafny.MultiSet
            d_1129_v37_ = _dafny.MultiSet([(d_1079_v2_).f24])
            d_1130_v38_: _dafny.Seq
            d_1130_v38_ = _dafny.SeqWithoutIsStrInference([(d_1079_v2_).f24])
            d_1131_v39_: _dafny.Seq
            d_1131_v39_ = _dafny.SeqWithoutIsStrInference([d_1117_v25_])
            d_1132_v40_: _dafny.Map
            d_1132_v40_ = _dafny.Map({d_1131_v39_: d_1118_v26_})
            d_1133_v41_: _dafny.Array
            nw188_ = _dafny.Array(None, 21)
            nw188_[int(0)] = d_1079_v2_.f23
            nw188_[int(1)] = p2
            nw188_[int(2)] = ((d_1129_v37_)[(d_1130_v38_)[default__.safeIndex(d_1079_v2_.f23, len(d_1130_v38_))]] if ((d_1130_v38_)[default__.safeIndex(d_1079_v2_.f23, len(d_1130_v38_))]) in (d_1129_v37_) else (d_1082_v3_)[default__.safeIndex(302, (d_1082_v3_).length(0))])
            nw188_[int(3)] = d_1079_v2_.f23
            nw188_[int(4)] = (-86) + (232)
            nw188_[int(5)] = p0
            nw188_[int(6)] = -7
            nw188_[int(7)] = (self).f11
            nw188_[int(8)] = p1
            nw188_[int(9)] = d_1079_v2_.f23
            nw188_[int(10)] = -488
            nw188_[int(11)] = (d_1082_v3_)[default__.safeIndex(302, (d_1082_v3_).length(0))]
            nw188_[int(12)] = (self).f11
            nw188_[int(13)] = ((self).f11) - (874)
            nw188_[int(14)] = default__.safeDivisionInt(p0, (((d_1132_v40_)[d_1131_v39_] if (d_1131_v39_) in (d_1132_v40_) else d_1118_v26_)).cardinality)
            nw188_[int(15)] = p1
            nw188_[int(16)] = ((d_1079_v2_).f24) - ((self).f11)
            nw188_[int(17)] = (0) - ((d_1079_v2_).f24)
            nw188_[int(18)] = p1
            nw188_[int(19)] = (self).f11
            nw188_[int(20)] = (0) - (len(_dafny.SeqWithoutIsStrInference([d_1078_v1_, d_1078_v1_])))
            d_1133_v41_ = nw188_
            r2 = d_1133_v41_
            r3 = d_1079_v2_.f23
            if (d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]:
                index161_ = default__.safeIndex(302, (d_1082_v3_).length(0))
                (d_1082_v3_)[index161_] = (self).f11
                index162_ = default__.safeIndex(302, (d_1082_v3_).length(0))
                (d_1082_v3_)[index162_] = (len((self).f9) if (d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))] else d_1079_v2_.f23)
                index163_ = default__.safeIndex(511, (d_1117_v25_).length(0))
                (d_1117_v25_)[index163_] = default__.fm38((d_1082_v3_)[default__.safeIndex(302, (d_1082_v3_).length(0))], globalState)
                index164_ = default__.safeIndex(511, (d_1117_v25_).length(0))
                (d_1117_v25_)[index164_] = d_1077_v0_
                r1 = default__.fm38((0) - (default__.safeDivisionInt((d_1079_v2_).f24, -890)), globalState)
            elif True:
                d_1134_v42_: _dafny.Seq
                d_1134_v42_ = _dafny.SeqWithoutIsStrInference([d_1077_v0_, (d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]])
                d_1135_v43_: C0
                nw189_ = C0()
                nw189_.ctor__(d_1077_v0_, (p2) != (len(d_1134_v42_)), (self).f9, (self).f10)
                d_1135_v43_ = nw189_
                index165_ = default__.safeIndex(511, (d_1117_v25_).length(0))
                (d_1117_v25_)[index165_] = ((_dafny.MultiSet([(d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]])).intersection(d_1118_v26_)).isdisjoint((_dafny.MultiSet([d_1077_v0_, d_1077_v0_, (d_1135_v43_).f25])) - (d_1118_v26_))
                (globalState).f3 = (d_1130_v38_)[default__.safeIndex((d_1130_v38_)[default__.safeIndex(d_1079_v2_.f23, len(d_1130_v38_))], len(d_1130_v38_))]
                d_1117_v25_ = d_1117_v25_
                d_1136_v44_: _dafny.Map
                d_1136_v44_ = _dafny.Map({(d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]: d_1078_v1_})
                d_1137_v45_: _dafny.Map
                d_1137_v45_ = _dafny.Map({(d_1079_v2_.f23) <= (d_1079_v2_.f23): ((d_1136_v44_)[(d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]] if ((d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]) in (d_1136_v44_) else d_1078_v1_)})
                d_1136_v44_ = (d_1136_v44_) | (d_1136_v44_)
        d_1138_v46_: _dafny.Seq
        d_1138_v46_ = _dafny.SeqWithoutIsStrInference([d_1077_v0_, (d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]])
        rhs152_ = ((_dafny.SeqWithoutIsStrInference([d_1077_v0_, d_1077_v0_, not(d_1077_v0_), (d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]])) + (d_1138_v46_) if (d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))] else (d_1138_v46_) + (d_1138_v46_))
        rhs153_ = 511
        rhs154_ = (d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]
        rhs155_ = d_1082_v3_
        rhs156_ = False
        d_1138_v46_ = rhs152_
        r3 = rhs153_
        r1 = rhs154_
        r2 = rhs155_
        r0 = rhs156_
        r0 = (d_1117_v25_)[default__.safeIndex(511, (d_1117_v25_).length(0))]
        r1 = not((default__.safeDivisionInt(p1, -735)) >= (d_1079_v2_.f23))
        r2 = d_1082_v3_
        r3 = default__.safeDivisionInt(223, (self).f11)
        return r0, r1, r2, r3

    def m15(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        d_1139_v0_: bool
        d_1139_v0_ = False
        d_1139_v0_ = d_1139_v0_
        d_1140_v1_: _dafny.Set
        d_1140_v1_ = _dafny.Set({d_1139_v0_})
        d_1141_v2_: _dafny.Seq
        d_1141_v2_ = _dafny.SeqWithoutIsStrInference([(self).f11, len(d_1140_v1_), (self).f11])
        d_1142_v3_: _dafny.Map
        d_1142_v3_ = _dafny.Map({(self).f11: d_1139_v0_})
        d_1143_v4_: _dafny.Map
        d_1143_v4_ = _dafny.Map({True: _dafny.CodePoint('k')})
        d_1144_v5_: _dafny.Map
        d_1144_v5_ = _dafny.Map({-956: (0) - (len(d_1143_v4_))})
        d_1145_v6_: _dafny.Seq
        d_1145_v6_ = _dafny.SeqWithoutIsStrInference([(self).f11, len((_dafny.Map({len(d_1141_v2_): d_1139_v0_})) | (d_1142_v3_)), len(d_1144_v5_)])
        d_1145_v6_ = default__.fm44(d_1139_v0_, (self).f11, not(d_1139_v0_), globalState)
        d_1146_v7_: _dafny.Seq
        d_1146_v7_ = _dafny.SeqWithoutIsStrInference([d_1139_v0_, d_1139_v0_])
        d_1147_v8_: _dafny.Set
        d_1147_v8_ = _dafny.Set({(self).f11})
        d_1148_v9_: _dafny.Map
        d_1148_v9_ = _dafny.Map({d_1147_v8_: d_1146_v7_})
        d_1149_v10_: _dafny.Array
        nw190_ = _dafny.Array(None, 21)
        nw190_[int(0)] = (d_1146_v7_) + (d_1146_v7_)
        nw190_[int(1)] = default__.fm2((self).f9, len(d_1142_v3_), globalState)
        nw190_[int(2)] = d_1146_v7_
        nw190_[int(3)] = d_1146_v7_
        nw190_[int(4)] = d_1146_v7_
        nw190_[int(5)] = ((d_1148_v9_)[d_1147_v8_] if (d_1147_v8_) in (d_1148_v9_) else _dafny.SeqWithoutIsStrInference([d_1139_v0_]))
        nw190_[int(6)] = d_1146_v7_
        nw190_[int(7)] = d_1146_v7_
        nw190_[int(8)] = d_1146_v7_
        nw190_[int(9)] = _dafny.SeqWithoutIsStrInference([d_1139_v0_])
        nw190_[int(10)] = d_1146_v7_
        nw190_[int(11)] = d_1146_v7_
        nw190_[int(12)] = d_1146_v7_
        nw190_[int(13)] = d_1146_v7_
        nw190_[int(14)] = d_1146_v7_
        nw190_[int(15)] = d_1146_v7_
        nw190_[int(16)] = d_1146_v7_
        nw190_[int(17)] = ((d_1146_v7_) + (d_1146_v7_)).set(default__.safeIndex(p1, len((d_1146_v7_) + (d_1146_v7_))), d_1139_v0_)
        nw190_[int(18)] = _dafny.SeqWithoutIsStrInference([d_1139_v0_, d_1139_v0_, d_1139_v0_, d_1139_v0_, d_1139_v0_])
        nw190_[int(19)] = d_1146_v7_
        nw190_[int(20)] = d_1146_v7_
        d_1149_v10_ = nw190_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1149_v10_).length(0)):
            d_1150_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_1150_i0_)) and ((d_1150_i0_) < ((d_1149_v10_).length(0)))):
                (d_1149_v10_)[(d_1150_i0_)] = d_1146_v7_
        d_1151_v11_: D4
        d_1151_v11_ = D4_DC11(d_1139_v0_, -601)
        d_1152_v12_: _dafny.Seq
        d_1152_v12_ = _dafny.SeqWithoutIsStrInference([d_1151_v11_, d_1151_v11_, d_1151_v11_])
        d_1153_v13_: _dafny.MultiSet
        d_1153_v13_ = _dafny.MultiSet([not(d_1139_v0_), True])
        pat_let_tv19_ = d_1153_v13_
        d_1154_v14_: _dafny.MultiSet
        def iife89_(_pat_let14_0):
            def iife90_(d_1155_dt__update__tmp_h0_):
                def iife91_(_pat_let15_0):
                    def iife92_(d_1156_dt__update_hcf14_h0_):
                        return D4_DC11((d_1155_dt__update__tmp_h0_).cf13, d_1156_dt__update_hcf14_h0_)
                    return iife92_(_pat_let15_0)
                return iife91_((pat_let_tv19_).cardinality)
            return iife90_(_pat_let14_0)
        d_1154_v14_ = _dafny.MultiSet([(d_1152_v12_) + (d_1152_v12_), _dafny.SeqWithoutIsStrInference([d_1151_v11_, iife89_(D4_DC11(d_1139_v0_, len(d_1142_v3_))), d_1151_v11_]), d_1152_v12_, _dafny.SeqWithoutIsStrInference([d_1151_v11_, d_1151_v11_]), d_1152_v12_])
        d_1154_v14_ = d_1154_v14_
        d_1157_v15_: _dafny.MultiSet
        d_1157_v15_ = _dafny.MultiSet([(self).f11, p1])
        d_1158_v16_: D3
        d_1158_v16_ = D3_DC7(d_1139_v0_)
        d_1159_v17_: D15
        d_1159_v17_ = D15_DC37(((d_1157_v15_)[p1] if (p1) in (d_1157_v15_) else p1), d_1158_v16_, not(d_1139_v0_))
        d_1160_v18_: _dafny.Seq
        d_1160_v18_ = _dafny.SeqWithoutIsStrInference([(self).f9, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wniqnefmy"))])
        d_1161_v19_: D9
        d_1161_v19_ = D9_DC20(d_1160_v18_)
        d_1162_v20_: _dafny.Map
        d_1162_v20_ = _dafny.Map({d_1159_v17_: d_1161_v19_})
        d_1162_v20_ = d_1162_v20_
        d_1163_v21_: _dafny.Array
        nw191_ = _dafny.Array(D6.default()(), 2)
        d_1163_v21_ = nw191_
        d_1164_v22_: D6
        d_1164_v22_ = D6_DC14(p2)
        index166_ = default__.safeIndex(970, (d_1163_v21_).length(0))
        (d_1163_v21_)[index166_] = d_1164_v22_
        index167_ = default__.safeIndex(970, (d_1163_v21_).length(0))
        (d_1163_v21_)[index167_] = D6_DC14(p2)
        r0 = (self).f11
        r1 = (_dafny.MultiSet([d_1139_v0_, d_1139_v0_])).intersection((d_1153_v13_ if d_1139_v0_ else _dafny.MultiSet([d_1139_v0_, default__.fm38((self).f11, globalState)])))
        return r0, r1

    def m16(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Seq = _dafny.Seq({})
        r2: bool = False
        d_1165_v0_: _dafny.Seq
        d_1165_v0_ = _dafny.SeqWithoutIsStrInference([p0])
        r2 = ((p0) or ((d_1165_v0_)[default__.safeIndex(-664, len(d_1165_v0_))]) if ((self).f11) == ((self).f11) else p0)
        d_1166_v1_: _dafny.Map
        d_1166_v1_ = _dafny.Map({(self).f11: 142})
        d_1167_v2_: str
        d_1167_v2_ = _dafny.CodePoint('c')
        d_1168_v3_: _dafny.Map
        d_1168_v3_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f11]): d_1167_v2_})
        d_1169_v4_: D14
        d_1169_v4_ = D14_DC34(len(d_1168_v3_), (self).f11, (self).f11, p0)
        d_1170_v5_: _dafny.Array
        nw192_ = _dafny.Array(None, 22)
        nw192_[int(0)] = (self).f11
        nw192_[int(1)] = 270
        nw192_[int(2)] = (self).f11
        nw192_[int(3)] = ((0) - ((self).f11) if p0 else (self).f11)
        nw192_[int(4)] = (self).f11
        nw192_[int(5)] = ((self).f11) + ((self).f11)
        nw192_[int(6)] = (self).f11
        nw192_[int(7)] = 680
        nw192_[int(8)] = ((self).f11) * (len(d_1166_v1_))
        nw192_[int(9)] = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwaeq"))), (self).f11)
        nw192_[int(10)] = (self).f11
        nw192_[int(11)] = (d_1169_v4_).cf44
        nw192_[int(12)] = 362
        nw192_[int(13)] = (self).f11
        nw192_[int(14)] = (self).f11
        nw192_[int(15)] = len(d_1165_v0_)
        nw192_[int(16)] = (self).f11
        nw192_[int(17)] = (self).f11
        nw192_[int(18)] = (self).f11
        nw192_[int(19)] = ((self).f11) * ((self).f11)
        nw192_[int(20)] = len((d_1166_v1_).set((self).f11, (self).f11))
        nw192_[int(21)] = (len((self).f9)) - ((self).f11)
        d_1170_v5_ = nw192_
        index168_ = default__.safeIndex(889, (d_1170_v5_).length(0))
        (d_1170_v5_)[index168_] = (self).f11
        index169_ = default__.safeIndex(889, (d_1170_v5_).length(0))
        (d_1170_v5_)[index169_] = (0) - ((self).f11)
        d_1171_v6_: _dafny.Array
        nw193_ = _dafny.Array(None, 21)
        nw193_[int(0)] = d_1167_v2_
        nw193_[int(1)] = d_1167_v2_
        nw193_[int(2)] = (d_1167_v2_ if p0 else d_1167_v2_)
        nw193_[int(3)] = d_1167_v2_
        nw193_[int(4)] = d_1167_v2_
        nw193_[int(5)] = d_1167_v2_
        nw193_[int(6)] = d_1167_v2_
        nw193_[int(7)] = d_1167_v2_
        nw193_[int(8)] = d_1167_v2_
        nw193_[int(9)] = d_1167_v2_
        nw193_[int(10)] = d_1167_v2_
        nw193_[int(11)] = d_1167_v2_
        nw193_[int(12)] = d_1167_v2_
        nw193_[int(13)] = d_1167_v2_
        nw193_[int(14)] = d_1167_v2_
        nw193_[int(15)] = d_1167_v2_
        nw193_[int(16)] = d_1167_v2_
        nw193_[int(17)] = ((self).f9)[default__.safeIndex(404, len((self).f9))]
        nw193_[int(18)] = d_1167_v2_
        nw193_[int(19)] = d_1167_v2_
        nw193_[int(20)] = d_1167_v2_
        d_1171_v6_ = nw193_
        d_1172_v7_: _dafny.Set
        d_1172_v7_ = _dafny.Set({(d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))]})
        d_1173_v8_: _dafny.Map
        d_1173_v8_ = _dafny.Map({(d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))]: (d_1166_v1_).set((d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))], len(d_1172_v7_))})
        d_1174_v9_: _dafny.Seq
        d_1174_v9_ = _dafny.SeqWithoutIsStrInference([(self).f11])
        d_1175_v10_: T1
        nw194_ = C1()
        nw194_.ctor__(((d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))]) - ((d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))]), ((d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))]) + ((d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))]), len(_dafny.Map({p0: (d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))]})), _dafny.Map({((d_1173_v8_)[(d_1174_v9_)[default__.safeIndex((self).f11, len(d_1174_v9_))]] if ((d_1174_v9_)[default__.safeIndex((self).f11, len(d_1174_v9_))]) in (d_1173_v8_) else d_1166_v1_): d_1166_v1_}), (self).f9, (self).f10)
        d_1175_v10_ = nw194_
        index170_ = default__.safeIndex(889, (d_1170_v5_).length(0))
        rhs157_ = d_1171_v6_
        rhs158_ = ((d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))]) < (len(d_1174_v9_))
        rhs159_ = 527
        rhs160_ = d_1175_v10_
        lhs101_ = d_1170_v5_
        lhs102_ = default__.safeIndex(889, (d_1170_v5_).length(0))
        d_1171_v6_ = rhs157_
        r2 = rhs158_
        lhs101_[lhs102_] = rhs159_
        d_1175_v10_ = rhs160_
        d_1176_v11_: _dafny.Map
        d_1176_v11_ = _dafny.Map({(d_1175_v10_).f9: (d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))]})
        d_1177_v12_: _dafny.Seq
        d_1177_v12_ = (d_1175_v10_).f9
        d_1178_v13_: _dafny.Array
        d_1179_v14_: _dafny.Array
        out45_: _dafny.Array
        out46_: _dafny.Array
        out45_, out46_ = (d_1175_v10_).m3(((0) - ((self).f11)) - ((d_1175_v10_).f11), (d_1175_v10_).f9, ((d_1176_v11_)[(self).f9] if ((self).f9) in (d_1176_v11_) else (d_1175_v10_).f11), d_1177_v12_, globalState)
        d_1178_v13_ = out45_
        d_1179_v14_ = out46_
        (globalState).f3 = (d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))]
        d_1180_v15_: _dafny.Array
        nw195_ = _dafny.Array(False, 2)
        d_1180_v15_ = nw195_
        d_1181_v16_: _dafny.Seq
        d_1181_v16_ = _dafny.SeqWithoutIsStrInference([d_1180_v15_, d_1180_v15_, d_1180_v15_, d_1180_v15_])
        if not (((0) - (len(d_1181_v16_))) > ((d_1175_v10_).f11)) or (p0):
            r2 = (725) > ((self).f11)
            d_1182_v17_: D4
            d_1182_v17_ = D4_DC11(True, (d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))])
            r2 = (len(_dafny.SeqWithoutIsStrInference([d_1167_v2_ for d_1183_i0_ in range(default__.abs(997))]))) == ((d_1182_v17_).cf14)
            d_1184_v18_: bool
            d_1185_v19_: bool
            d_1186_v20_: _dafny.Array
            d_1187_v21_: int
            out47_: bool
            out48_: bool
            out49_: _dafny.Array
            out50_: int
            out47_, out48_, out49_, out50_ = (d_1175_v10_).m1(default__.safeModuloInt((_dafny.MultiSet([(self).f11])).cardinality, (self).f11), (d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))], -250, globalState)
            d_1184_v18_ = out47_
            d_1185_v19_ = out48_
            d_1186_v20_ = out49_
            d_1187_v21_ = out50_
            d_1184_v18_ = d_1185_v19_
            d_1185_v19_ = d_1184_v18_
        elif True:
            d_1166_v1_ = d_1166_v1_
            (globalState).f3 = default__.safeDivisionInt(len(default__.fm37((d_1175_v10_).f11, (d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))], p0, (d_1175_v10_).f11, globalState)), (d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))])
            (globalState).f3 = (((d_1175_v10_).f11) + ((d_1175_v10_).f11)) - (((d_1175_v10_).f11 if p0 else (d_1175_v10_).f11))
            d_1188_v22_: _dafny.MultiSet
            d_1188_v22_ = _dafny.MultiSet([default__.fm38(len((d_1175_v10_).f9), globalState), p0, p0, p0])
            d_1189_v23_: _dafny.Map
            d_1189_v23_ = _dafny.Map({p0: p0})
            r2 = not((p0) and (not((((d_1188_v22_)[not(p0)] if (not(p0)) in (d_1188_v22_) else (d_1175_v10_).f11)) == (len(d_1189_v23_)))))
            d_1190_v25_: _dafny.Seq
            d_1190_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1175_v10_).f11: (d_1175_v10_).f11})])
            d_1191_v26_: C1
            nw196_ = C1()
            def iife93_():
                coll61_ = _dafny.Map()
                compr_61_: _dafny.Map
                for compr_61_ in (d_1190_v25_).Elements:
                    d_1192_v24_: _dafny.Map = compr_61_
                    if (d_1192_v24_) in (d_1190_v25_):
                        coll61_[d_1192_v24_] = d_1166_v1_
                return _dafny.Map(coll61_)
            nw196_.ctor__((d_1175_v10_).f11, 555, (self).f11, iife93_()
            , (d_1175_v10_).f9, _dafny.SeqWithoutIsStrInference([d_1174_v9_, d_1174_v9_]))
            d_1191_v26_ = nw196_
            d_1193_v27_: _dafny.MultiSet
            d_1193_v27_ = _dafny.MultiSet([d_1191_v26_])
            d_1194_v28_: _dafny.Seq
            d_1194_v28_ = _dafny.SeqWithoutIsStrInference([d_1193_v27_])
            d_1195_v29_: _dafny.Set
            d_1195_v29_ = _dafny.Set({p0, p0})
            (globalState).f3 = ((d_1194_v28_)[default__.safeIndex(len(d_1195_v29_), len(d_1194_v28_))]).cardinality
        nw197_ = _dafny.Array(int(0), 12)
        r0 = nw197_
        r1 = default__.fm44(((d_1170_v5_)[default__.safeIndex(889, (d_1170_v5_).length(0))]) > ((d_1175_v10_).f11), (d_1175_v10_).f11, p0, globalState)
        r2 = p0
        return r0, r1, r2


class C3(T1, T0):
    def  __init__(self):
        self._f12: _dafny.Map = _dafny.Map({})
        self._f11: int = int(0)
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: _dafny.Seq = _dafny.Seq({})
        self._f21: bool = False
        self._f22: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f11(self):
        return self._f11
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f21, f22, f11, f12, f9, f10):
        (self)._f21 = f21
        (self)._f22 = f22
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f9 = f9
        (self)._f10 = f10

    def fm5(self, p0, p1, globalState):
        return len((_dafny.SeqWithoutIsStrInference([(self).f9, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvbse")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ijam")), (self).f9])) + ((D9_DC20(_dafny.SeqWithoutIsStrInference([(self).f9 for d_1196_i0_ in range(default__.abs(-728))]))).cf21))

    def fm29(self, p0, p1, p2, globalState):
        source20_ = D3_DC7((self).f21)
        if source20_.is_DC7:
            d_1197___mcc_h0_ = source20_.cf9
            d_1198_cf9_ = d_1197___mcc_h0_
            return (self).f22
        elif True:
            d_1199___mcc_h1_ = source20_.cf8
            d_1200_cf8_ = d_1199___mcc_h1_
            return (self).f11

    def fm30(self, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1201_i0_ in range(default__.abs(151))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnatxbt")), (self).f9]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyqtru"))]))

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: bool = False
        r3: int = int(0)
        d_1202_v0_: _dafny.Array
        nw198_ = _dafny.Array(int(0), 11)
        d_1202_v0_ = nw198_
        d_1203_v1_: _dafny.Seq
        d_1203_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqbpp"))
        d_1204_v2_: _dafny.Seq
        d_1204_v2_ = _dafny.SeqWithoutIsStrInference([(self).f21, (self).f21, not(False)])
        d_1205_v3_: _dafny.Seq
        d_1205_v3_ = _dafny.SeqWithoutIsStrInference([(d_1204_v2_) + (d_1204_v2_), d_1204_v2_, d_1204_v2_, d_1204_v2_, d_1204_v2_])
        d_1206_v4_: D1
        d_1206_v4_ = D1_DC2()
        d_1207_v5_: _dafny.MultiSet
        d_1207_v5_ = _dafny.MultiSet([(self).f9, default__.fm31(globalState), d_1203_v1_])
        d_1208_v6_: _dafny.Map
        d_1208_v6_ = _dafny.Map({not(False): (self).f22})
        rhs161_ = d_1202_v0_
        rhs162_ = ((default__.fm32((self).f11, True, globalState)) - (d_1207_v5_)).issubset((D10_DC22(d_1207_v5_)).cf25)
        rhs163_ = (self).f9
        rhs164_ = default__.fm33(d_1208_v6_, globalState)
        rhs165_ = (d_1206_v4_ if ((self).f11) <= ((self).f11) else d_1206_v4_)
        d_1202_v0_ = rhs161_
        r2 = rhs162_
        d_1203_v1_ = rhs163_
        d_1205_v3_ = rhs164_
        d_1206_v4_ = rhs165_
        d_1209_v7_: _dafny.Array
        def lambda46_(d_1210_i1_):
            return (not((self).f21)) == ((self).f21)

        init24_ = lambda46_
        nw199_ = _dafny.Array(None, 1)
        for i0_24_ in range(nw199_.length(0)):
            nw199_[i0_24_] = init24_(i0_24_)
        d_1209_v7_ = nw199_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1209_v7_).length(0)):
            d_1211_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_1211_i0_)) and ((d_1211_i0_) < ((d_1209_v7_).length(0)))):
                (d_1209_v7_)[(d_1211_i0_)] = ((self).f21 if (self).f21 else not((self).f21))
        if (default__.fm31(globalState)) not in (d_1207_v5_):
            r3 = (self).fm5((self).f11, 367, globalState)
            r2 = (self).f21
            d_1212_v8_: _dafny.Array
            nw200_ = _dafny.Array(_dafny.CodePoint('D'), 28)
            d_1212_v8_ = nw200_
            d_1213_v9_: D6
            d_1213_v9_ = D6_DC13(d_1212_v8_)
            source21_ = d_1213_v9_
            if source21_.is_DC14:
                d_1214___mcc_h0_ = source21_.cf17
                d_1215_cf17_ = d_1214___mcc_h0_
                r2 = (self).f21
                d_1216_v10_: D4
                d_1216_v10_ = D4_DC8(_dafny.Map({p0: (self).f21}))
                d_1217_v11_: _dafny.Array
                nw201_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_1217_v11_ = nw201_
                d_1218_v12_: _dafny.Map
                d_1218_v12_ = _dafny.Map({d_1216_v10_: d_1217_v11_})
                d_1218_v12_ = (d_1218_v12_).set(d_1216_v10_, d_1217_v11_)
                rhs166_ = ((self).f21) or ((self).f21)
                rhs167_ = (self).f11
                rhs168_ = default__.safeModuloInt((self).f22, p0)
                rhs169_ = (0) - ((self).f22)
                lhs103_ = globalState
                r2 = rhs166_
                lhs103_.f3 = rhs167_
                r0 = rhs168_
                r3 = rhs169_
                d_1219_v13_: str
                d_1219_v13_ = _dafny.CodePoint('f')
                d_1220_v14_: _dafny.Seq
                d_1220_v14_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s'), d_1219_v13_, d_1219_v13_]))])
                d_1221_v15_: C2
                nw202_ = C2()
                nw202_.ctor__((self).f22, default__.fm47((self).f21, (self).f21, not((self).f21), default__.fm38((self).f11, globalState), globalState), (d_1203_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqrfnoj"))), _dafny.SeqWithoutIsStrInference([d_1220_v14_, d_1220_v14_, d_1220_v14_, default__.fm44(True, (d_1220_v14_)[default__.safeIndex((self).f11, len(d_1220_v14_))], (self).f21, globalState)]))
                d_1221_v15_ = nw202_
            elif True:
                d_1222___mcc_h1_ = source21_.cf16
                d_1223_cf16_ = d_1222___mcc_h1_
                d_1224_v16_: _dafny.Map
                d_1224_v16_ = _dafny.Map({(self).f22: 613})
                (globalState).f3 = (0) - (((d_1224_v16_)[((self).f22) * ((0) - ((self).f22))] if (((self).f22) * ((0) - ((self).f22))) in (d_1224_v16_) else len(d_1204_v2_)))
                (globalState).f3 = (0) - ((self).f22)
                d_1225_v17_: C2
                nw203_ = C2()
                nw203_.ctor__(((self).f11) - (964), (default__.fm47((self).f21, (self).f21, (self).f21, (d_1204_v2_)[default__.safeIndex(len(d_1208_v6_), len(d_1204_v2_))], globalState)) | (default__.fm47((self).f21, (self).f21, (self).f21, (self).f21, globalState)), (self).f9, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.Set({(self).f21, (self).f21})])).cardinality]) for d_1226_i2_ in range(default__.abs(844))]))
                d_1225_v17_ = nw203_
                d_1227_v18_: _dafny.Seq
                d_1227_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f11])])
                d_1228_v19_: _dafny.Seq
                d_1228_v19_ = _dafny.SeqWithoutIsStrInference([(self).f22])
                rhs170_ = default__.fm37((0) - (p0), p0, (d_1204_v2_) < (default__.fm2((self).f9, (self).f22, globalState)), ((self).f11) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfkpr")))), globalState)
                rhs171_ = (self).f22
                rhs172_ = (d_1204_v2_)[default__.safeIndex((D15_DC36((self).f11)).cf48, len(d_1204_v2_))]
                rhs173_ = (((d_1227_v18_)[default__.safeIndex(len(d_1228_v19_), len(d_1227_v18_))]).cardinality) <= (p0)
                lhs104_ = globalState
                d_1203_v1_ = rhs170_
                lhs104_.f3 = rhs171_
                r2 = rhs172_
                r2 = rhs173_
            d_1229_v20_: str
            d_1229_v20_ = _dafny.CodePoint('g')
            d_1230_v21_: _dafny.Map
            d_1230_v21_ = _dafny.Map({d_1229_v20_: ((self).f9) + ((self).f9)})
            d_1230_v21_ = (d_1230_v21_).set(d_1229_v20_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")))
            (globalState).f3 = (self).f11
        elif True:
            if (self).f21:
                d_1231_v22_: str
                d_1231_v22_ = _dafny.CodePoint('l')
                d_1232_v23_: _dafny.Map
                d_1232_v23_ = _dafny.Map({((0) - (p0)) + ((self).f22): d_1231_v22_})
                d_1232_v23_ = (d_1232_v23_).set(default__.safeDivisionInt((self).f22, (self).f11), d_1231_v22_)
                d_1233_v24_: _dafny.Map
                d_1233_v24_ = _dafny.Map({((self).fm5(p0, p0, globalState)) * (p0): not((self).f21)})
                d_1234_v25_: _dafny.MultiSet
                d_1234_v25_ = _dafny.MultiSet([p0])
                d_1233_v24_ = (d_1233_v24_).set(((d_1234_v25_).cardinality if (self).f21 else (self).f22), (self).f21)
                d_1203_v1_ = ((d_1203_v1_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1235_i3_ in range(default__.abs(406))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1236_i4_ in range(default__.abs(587))])) + ((self).f9))
                index171_ = default__.safeIndex(818, (d_1202_v0_).length(0))
                (d_1202_v0_)[index171_] = (self).f11
                index172_ = default__.safeIndex(818, (d_1202_v0_).length(0))
                (d_1202_v0_)[index172_] = (0) - ((p0) * (default__.safeModuloInt(861, len(d_1204_v2_))))
                index173_ = default__.safeIndex(166, (d_1209_v7_).length(0))
                (d_1209_v7_)[index173_] = (self).f21
                index174_ = default__.safeIndex(166, (d_1209_v7_).length(0))
                index175_ = default__.safeIndex(818, (d_1202_v0_).length(0))
                index176_ = default__.safeIndex(818, (d_1202_v0_).length(0))
                rhs174_ = (self).f21
                rhs175_ = (d_1202_v0_)[default__.safeIndex(818, (d_1202_v0_).length(0))]
                rhs176_ = not((self).f21)
                rhs177_ = (self).f22
                rhs178_ = -161
                lhs105_ = d_1209_v7_
                lhs106_ = default__.safeIndex(166, (d_1209_v7_).length(0))
                lhs107_ = d_1202_v0_
                lhs108_ = default__.safeIndex(818, (d_1202_v0_).length(0))
                lhs109_ = globalState
                lhs110_ = d_1202_v0_
                lhs111_ = default__.safeIndex(818, (d_1202_v0_).length(0))
                lhs105_[lhs106_] = rhs174_
                lhs107_[lhs108_] = rhs175_
                r2 = rhs176_
                lhs109_.f3 = rhs177_
                lhs110_[lhs111_] = rhs178_
            elif True:
                d_1204_v2_ = (_dafny.SeqWithoutIsStrInference([not((self).f21), (self).f21])) + (d_1204_v2_)
                d_1237_v26_: str
                d_1237_v26_ = _dafny.CodePoint('d')
                d_1237_v26_ = d_1237_v26_
                d_1238_v27_: _dafny.Map
                d_1238_v27_ = _dafny.Map({(self).f11: _dafny.SeqWithoutIsStrInference([(self).f22, p0])})
                r2 = (len(((d_1204_v2_) + (d_1204_v2_)).set(default__.safeIndex(p0, len((d_1204_v2_) + (d_1204_v2_))), False))) not in (d_1238_v27_)
                d_1239_v28_: _dafny.Map
                d_1239_v28_ = _dafny.Map({(self).f21: not(True)})
                d_1240_v29_: _dafny.Array
                nw204_ = _dafny.Array(None, 3)
                nw204_[int(0)] = (d_1239_v28_) | (_dafny.Map({(self).f21: (self).f21}))
                nw204_[int(1)] = d_1239_v28_
                nw204_[int(2)] = d_1239_v28_
                d_1240_v29_ = nw204_
                index177_ = default__.safeIndex(88, (d_1240_v29_).length(0))
                (d_1240_v29_)[index177_] = _dafny.Map({(self).f21: (self).f21})
                index178_ = default__.safeIndex(88, (d_1240_v29_).length(0))
                (d_1240_v29_)[index178_] = ((d_1239_v28_) | (d_1239_v28_)) | (d_1239_v28_)
                (globalState).f3 = default__.safeModuloInt(271, default__.safeModuloInt(default__.fm0(p0, 276, globalState), (0) - (p0)))
            r0 = (self).f11
            d_1241_v30_: D14
            d_1241_v30_ = D14_DC33(len((self).f9), d_1203_v1_, (self).f21)
            d_1242_v31_: _dafny.Array
            nw205_ = _dafny.Array(None, 26)
            nw205_[int(0)] = (self).f9
            nw205_[int(1)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1243_i5_ in range(default__.abs(-152))])
            nw205_[int(2)] = (self).f9
            nw205_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypfklawbb"))
            nw205_[int(4)] = (self).f9
            nw205_[int(5)] = (self).f9
            nw205_[int(6)] = ((self).f9) + ((d_1241_v30_).cf41)
            nw205_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpjcs"))
            nw205_[int(8)] = d_1203_v1_
            nw205_[int(9)] = (self).f9
            nw205_[int(10)] = (self).f9
            nw205_[int(11)] = (default__.fm37(p0, p0, (self).f21, (self).f11, globalState)) + (d_1203_v1_)
            nw205_[int(12)] = (self).f9
            nw205_[int(13)] = (self).f9
            nw205_[int(14)] = d_1203_v1_
            nw205_[int(15)] = (default__.fm31(globalState)) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1244_i6_ in range(default__.abs(664))]))
            nw205_[int(16)] = d_1203_v1_
            nw205_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jaablnfpo"))
            nw205_[int(18)] = d_1203_v1_
            nw205_[int(19)] = ((self).f9) + (d_1203_v1_)
            nw205_[int(20)] = (d_1203_v1_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1245_i7_ in range(default__.abs(984))]))
            nw205_[int(21)] = d_1203_v1_
            nw205_[int(22)] = (self).f9
            nw205_[int(23)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlohlnv"))) + (d_1203_v1_)
            nw205_[int(24)] = (self).f9
            nw205_[int(25)] = d_1203_v1_
            d_1242_v31_ = nw205_
            index179_ = default__.safeIndex(830, (d_1242_v31_).length(0))
            (d_1242_v31_)[index179_] = (self).f9
            index180_ = default__.safeIndex(830, (d_1242_v31_).length(0))
            (d_1242_v31_)[index180_] = d_1203_v1_
            d_1246_v33_: str
            d_1246_v33_ = _dafny.CodePoint('o')
            d_1247_v34_: _dafny.Set
            d_1247_v34_ = _dafny.Set({d_1246_v33_, d_1246_v33_, d_1246_v33_, _dafny.CodePoint('k'), d_1246_v33_})
            def iife94_():
                coll62_ = _dafny.Set()
                compr_62_: str
                for compr_62_ in (d_1203_v1_).Elements:
                    d_1248_v32_: str = compr_62_
                    if (d_1248_v32_) in (d_1203_v1_):
                        coll62_ = coll62_.union(_dafny.Set([d_1248_v32_]))
                return _dafny.Set(coll62_)
            r2 = ((iife94_()
            ) != (d_1247_v34_) if (self).f21 else (self).f21)
            r2 = not((self).f21)
        r2 = not((self).f21)
        r2 = (self).f21
        d_1249_v35_: _dafny.Map
        d_1249_v35_ = _dafny.Map({(self).f11: (self).f21})
        d_1249_v35_ = (d_1249_v35_).set((self).f11, (self).f21)
        r0 = default__.safeDivisionInt((self).f11, len(_dafny.SeqWithoutIsStrInference([(self).f21, (self).f21, (d_1204_v2_)[default__.safeIndex((self).f11, len(d_1204_v2_))]])))
        d_1250_v36_: _dafny.Set
        d_1250_v36_ = _dafny.Set({(self).f21})
        d_1251_v37_: _dafny.Map
        d_1251_v37_ = _dafny.Map({not(not((d_1250_v36_).issubset(d_1250_v36_))): (self).f9})
        r1 = d_1251_v37_
        r2 = (self).f21
        d_1252_v38_: _dafny.Set
        d_1252_v38_ = _dafny.Set({33, (self).f22})
        d_1253_v39_: _dafny.MultiSet
        d_1253_v39_ = _dafny.MultiSet([d_1252_v38_, (d_1252_v38_ if True else d_1252_v38_), default__.fm48((self).f21, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvbw"))), 573, (self).f21, globalState)])
        r3 = ((d_1253_v39_)[d_1252_v38_] if (d_1252_v38_) in (d_1253_v39_) else default__.safeModuloInt((self).f22, (self).f11))
        return r0, r1, r2, r3

    def m3(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_1254_v0_: _dafny.Seq
        d_1254_v0_ = _dafny.SeqWithoutIsStrInference([len(p1)])
        d_1255_v1_: D12
        d_1255_v1_ = D12_DC28(d_1254_v0_)
        pat_let_tv20_ = d_1254_v0_
        def iife95_(_pat_let16_0):
            def iife96_(d_1256_dt__update__tmp_h0_):
                def iife97_(_pat_let17_0):
                    def iife98_(d_1257_dt__update_hcf31_h0_):
                        return D12_DC28(d_1257_dt__update_hcf31_h0_)
                    return iife98_(_pat_let17_0)
                return iife97_(pat_let_tv20_)
            return iife96_(_pat_let16_0)
        source22_ = iife95_(d_1255_v1_)
        if source22_.is_DC29:
            d_1258___mcc_h0_ = source22_.cf32
            d_1259___mcc_h1_ = source22_.cf33
            d_1260___mcc_h2_ = source22_.cf34
            d_1261___mcc_h3_ = source22_.cf35
            d_1262___mcc_h4_ = source22_.cf36
            d_1263_cf36_ = d_1262___mcc_h4_
            d_1264_cf35_ = d_1261___mcc_h3_
            d_1265_cf34_ = d_1260___mcc_h2_
            d_1266_cf33_ = d_1259___mcc_h1_
            d_1267_cf32_ = d_1258___mcc_h0_
            d_1268_v2_: _dafny.Map
            d_1268_v2_ = _dafny.Map({False: d_1267_cf32_})
            if ((d_1268_v2_)[d_1264_cf35_] if (d_1264_cf35_) in (d_1268_v2_) else not(d_1267_cf32_)):
                d_1269_v3_: _dafny.Array
                def lambda47_(d_1270_i0_):
                    return (self).f21

                init25_ = lambda47_
                nw206_ = _dafny.Array(None, 4)
                for i0_25_ in range(nw206_.length(0)):
                    nw206_[i0_25_] = init25_(i0_25_)
                d_1269_v3_ = nw206_
                index181_ = default__.safeIndex(265, (d_1269_v3_).length(0))
                (d_1269_v3_)[index181_] = d_1267_cf32_
                index182_ = default__.safeIndex(265, (d_1269_v3_).length(0))
                (d_1269_v3_)[index182_] = default__.fm38(p0, globalState)
                d_1264_cf35_ = (default__.safeDivisionInt(-154, 892)) > (p2)
                d_1267_cf32_ = not((((self).f9) + (p1)) < ((self).f9))
                (globalState).f3 = ((0) - (d_1263_cf36_)) * (145)
                d_1271_v4_: _dafny.Array
                nw207_ = _dafny.Array(int(0), 2)
                d_1271_v4_ = nw207_
                d_1272_v5_: _dafny.Map
                d_1272_v5_ = _dafny.Map({d_1264_cf35_: len((self).f9)})
                index183_ = default__.safeIndex(160, (d_1271_v4_).length(0))
                (d_1271_v4_)[index183_] = len((d_1272_v5_).set((self).f21, 50))
                index184_ = default__.safeIndex(160, (d_1271_v4_).length(0))
                (d_1271_v4_)[index184_] = (self).f22
            elif True:
                d_1273_v6_: _dafny.MultiSet
                d_1273_v6_ = _dafny.MultiSet([d_1264_cf35_, (self).f21])
                d_1274_v7_: _dafny.Map
                d_1274_v7_ = _dafny.Map({d_1267_cf32_: d_1273_v6_})
                d_1275_v8_: _dafny.Map
                d_1275_v8_ = _dafny.Map({p1: ((d_1274_v7_)[d_1264_cf35_] if (d_1264_cf35_) in (d_1274_v7_) else _dafny.MultiSet([(self).f21, (self).f21]))})
                d_1264_cf35_ = (_dafny.Map({(self).f9: _dafny.MultiSet([d_1267_cf32_])})) != (d_1275_v8_)
                d_1276_v9_: str
                d_1276_v9_ = _dafny.CodePoint('y')
                d_1276_v9_ = d_1276_v9_
                d_1264_cf35_ = d_1264_cf35_
                d_1264_cf35_ = (d_1267_cf32_) or ((self).f21)
                d_1277_v10_: _dafny.Array
                nw208_ = _dafny.Array(D10.default()(), 14)
                d_1277_v10_ = nw208_
                d_1278_v11_: _dafny.Array
                d_1278_v11_ = d_1277_v10_
                d_1279_v12_: _dafny.Seq
                d_1279_v12_ = _dafny.SeqWithoutIsStrInference([d_1277_v10_, d_1277_v10_, (d_1278_v11_)])
                d_1279_v12_ = d_1279_v12_
            d_1280_v13_: _dafny.Array
            def lambda48_(d_1281_cf35_):
                def lambda49_(d_1282_i1_):
                    return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f21, True, (self).f21]))).isdisjoint(_dafny.MultiSet([d_1281_cf35_, d_1281_cf35_, not(d_1281_cf35_)]))

                return lambda49_

            init26_ = lambda48_(d_1264_cf35_)
            nw209_ = _dafny.Array(None, 11)
            for i0_26_ in range(nw209_.length(0)):
                nw209_[i0_26_] = init26_(i0_26_)
            d_1280_v13_ = nw209_
            index185_ = default__.safeIndex(719, (d_1280_v13_).length(0))
            (d_1280_v13_)[index185_] = d_1264_cf35_
            d_1283_v14_: _dafny.Seq
            d_1283_v14_ = _dafny.SeqWithoutIsStrInference([(self).f9, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1284_i2_ in range(default__.abs(815))])])
            index186_ = default__.safeIndex(719, (d_1280_v13_).length(0))
            rhs179_ = not(d_1267_cf32_)
            rhs180_ = len(((d_1283_v14_)[default__.safeIndex((self).f22, len(d_1283_v14_))]) + (p1))
            lhs112_ = d_1280_v13_
            lhs113_ = default__.safeIndex(719, (d_1280_v13_).length(0))
            lhs112_[lhs113_] = rhs179_
            d_1263_cf36_ = rhs180_
            d_1285_v15_: C0
            nw210_ = C0()
            nw210_.ctor__(d_1267_cf32_, not(True), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1286_i3_ in range(default__.abs(563))])) + (default__.fm37((self).f22, len(d_1265_cf34_), not((d_1280_v13_)[default__.safeIndex(719, (d_1280_v13_).length(0))]), p2, globalState)), (self).f10)
            d_1285_v15_ = nw210_
            d_1287_v16_: _dafny.Seq
            d_1287_v16_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1265_cf34_, d_1265_cf34_]), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([727, p0, (d_1265_cf34_)[default__.safeIndex(d_1263_cf36_, len(d_1265_cf34_))], d_1266_cf33_]) for d_1288_i4_ in range(default__.abs(-949))]), (self).f10, (self).f10, _dafny.SeqWithoutIsStrInference([d_1265_cf34_, _dafny.SeqWithoutIsStrInference([d_1263_cf36_ for d_1289_i5_ in range(default__.abs(6))])])])
            d_1290_v17_: C0
            nw211_ = C0()
            nw211_.ctor__(d_1267_cf32_, (d_1280_v13_)[default__.safeIndex(719, (d_1280_v13_).length(0))], p1, (d_1287_v16_)[default__.safeIndex((self).f22, len(d_1287_v16_))])
            d_1290_v17_ = nw211_
        elif source22_.is_DC28:
            d_1291___mcc_h5_ = source22_.cf31
            d_1292_cf31_ = d_1291___mcc_h5_
            d_1293_v18_: bool
            d_1293_v18_ = False
            d_1294_v19_: D11
            d_1294_v19_ = D11_DC26()
            d_1295_v20_: D11
            d_1295_v20_ = D11_DC27(d_1294_v19_)
            pat_let_tv21_ = d_1294_v19_
            pat_let_tv22_ = d_1294_v19_
            pat_let_tv23_ = d_1294_v19_
            pat_let_tv24_ = d_1294_v19_
            d_1296_v21_: _dafny.Array
            nw212_ = _dafny.Array(None, 23)
            nw212_[int(0)] = (d_1295_v20_ if not((self).f21) else d_1295_v20_)
            nw212_[int(1)] = d_1295_v20_
            def iife99_(_pat_let18_0):
                def iife100_(d_1297_dt__update__tmp_h1_):
                    def iife101_(_pat_let19_0):
                        def iife102_(d_1298_dt__update_hcf30_h0_):
                            return D11_DC27(d_1298_dt__update_hcf30_h0_)
                        return iife102_(_pat_let19_0)
                    return iife101_(pat_let_tv21_)
                return iife100_(_pat_let18_0)
            nw212_[int(2)] = iife99_(D11_DC27(d_1294_v19_))
            nw212_[int(3)] = (d_1295_v20_ if not((self).f21) else d_1295_v20_)
            nw212_[int(4)] = d_1295_v20_
            nw212_[int(5)] = d_1295_v20_
            nw212_[int(6)] = d_1295_v20_
            nw212_[int(7)] = d_1295_v20_
            nw212_[int(8)] = (d_1295_v20_ if d_1293_v18_ else D11_DC27(d_1294_v19_))
            nw212_[int(9)] = d_1295_v20_
            nw212_[int(10)] = d_1295_v20_
            nw212_[int(11)] = d_1295_v20_
            nw212_[int(12)] = d_1295_v20_
            nw212_[int(13)] = d_1295_v20_
            nw212_[int(14)] = (d_1295_v20_ if d_1293_v18_ else D11_DC27(d_1294_v19_))
            nw212_[int(15)] = d_1295_v20_
            nw212_[int(16)] = d_1295_v20_
            nw212_[int(17)] = d_1295_v20_
            nw212_[int(18)] = d_1295_v20_
            def iife103_(_pat_let20_0):
                def iife104_(d_1299_dt__update__tmp_h2_):
                    def iife105_(_pat_let21_0):
                        def iife106_(d_1300_dt__update_hcf30_h1_):
                            return D11_DC27(d_1300_dt__update_hcf30_h1_)
                        return iife106_(_pat_let21_0)
                    return iife105_(pat_let_tv22_)
                return iife104_(_pat_let20_0)
            nw212_[int(19)] = iife103_(d_1295_v20_)
            def iife108_(_pat_let23_0):
                def iife109_(d_1301_dt__update__tmp_h3_):
                    def iife110_(_pat_let24_0):
                        def iife111_(d_1302_dt__update_hcf30_h2_):
                            return D11_DC27(d_1302_dt__update_hcf30_h2_)
                        return iife111_(_pat_let24_0)
                    return iife110_(pat_let_tv23_)
                return iife109_(_pat_let23_0)
            def iife107_(_pat_let22_0):
                def iife112_(d_1303_dt__update__tmp_h4_):
                    def iife113_(_pat_let25_0):
                        def iife114_(d_1304_dt__update_hcf30_h3_):
                            return D11_DC27(d_1304_dt__update_hcf30_h3_)
                        return iife114_(_pat_let25_0)
                    return iife113_(pat_let_tv24_)
                return iife112_(_pat_let22_0)
            nw212_[int(20)] = iife107_(iife108_(d_1295_v20_))
            nw212_[int(21)] = D11_DC27(d_1294_v19_)
            nw212_[int(22)] = d_1295_v20_
            d_1296_v21_ = nw212_
            index187_ = default__.safeIndex(209, (d_1296_v21_).length(0))
            (d_1296_v21_)[index187_] = d_1295_v20_
            d_1305_v22_: _dafny.Array
            nw213_ = _dafny.Array(int(0), 28)
            d_1305_v22_ = nw213_
            d_1306_v23_: _dafny.Seq
            d_1306_v23_ = _dafny.SeqWithoutIsStrInference([d_1305_v22_, d_1305_v22_, d_1305_v22_, d_1305_v22_, d_1305_v22_])
            pat_let_tv25_ = d_1294_v19_
            index188_ = default__.safeIndex(209, (d_1296_v21_).length(0))
            def iife115_(_pat_let26_0):
                def iife116_(d_1307_dt__update__tmp_h5_):
                    def iife117_(_pat_let27_0):
                        def iife118_(d_1308_dt__update_hcf30_h4_):
                            return D11_DC27(d_1308_dt__update_hcf30_h4_)
                        return iife118_(_pat_let27_0)
                    return iife117_(pat_let_tv25_)
                return iife116_(_pat_let26_0)
            rhs181_ = (d_1306_v23_) == (_dafny.SeqWithoutIsStrInference([d_1305_v22_, d_1305_v22_, d_1305_v22_, d_1305_v22_]))
            rhs182_ = p2
            rhs183_ = (self).f22
            rhs184_ = (self).f21
            rhs185_ = iife115_(d_1295_v20_)
            lhs114_ = globalState
            lhs115_ = globalState
            lhs116_ = d_1296_v21_
            lhs117_ = default__.safeIndex(209, (d_1296_v21_).length(0))
            d_1293_v18_ = rhs181_
            lhs114_.f3 = rhs182_
            lhs115_.f3 = rhs183_
            d_1293_v18_ = rhs184_
            lhs116_[lhs117_] = rhs185_
            d_1309_v24_: _dafny.MultiSet
            d_1309_v24_ = _dafny.MultiSet([(self).f21, d_1293_v18_])
            d_1310_v25_: _dafny.Set
            d_1310_v25_ = _dafny.Set({(d_1309_v24_).cardinality})
            d_1311_v27_: _dafny.Map
            d_1311_v27_ = _dafny.Map({(0) - ((self).f11): (self).f22})
            d_1312_v28_: _dafny.Set
            d_1312_v28_ = _dafny.Set({d_1311_v27_, d_1311_v27_})
            d_1313_v29_: C1
            nw214_ = C1()
            def iife119_():
                coll63_ = _dafny.Map()
                compr_63_: _dafny.Map
                for compr_63_ in (d_1312_v28_).Elements:
                    d_1314_v26_: _dafny.Map = compr_63_
                    if (d_1314_v26_) in (d_1312_v28_):
                        coll63_[d_1314_v26_] = d_1311_v27_
                return _dafny.Map(coll63_)
            nw214_.ctor__(-431, (0) - (len(d_1310_v25_)), default__.safeDivisionInt((self).f22, -143), iife119_()
            , (self).f9, (_dafny.SeqWithoutIsStrInference([d_1292_cf31_ for d_1315_i6_ in range(default__.abs(708))])) + ((self).f10))
            d_1313_v29_ = nw214_
            d_1316_v30_: _dafny.Seq
            d_1316_v30_ = _dafny.SeqWithoutIsStrInference([d_1293_v18_])
            d_1317_v31_: _dafny.Map
            d_1317_v31_ = _dafny.Map({((d_1316_v30_)[default__.safeIndex(709, len(d_1316_v30_))] if d_1293_v18_ else not(d_1293_v18_)): default__.fm38(-284, globalState)})
            d_1317_v31_ = d_1317_v31_
            d_1318_v32_: _dafny.Array
            def lambda50_(d_1319_v18_):
                def lambda51_(d_1320_i7_):
                    return d_1319_v18_

                return lambda51_

            init27_ = lambda50_(d_1293_v18_)
            nw215_ = _dafny.Array(None, 8)
            for i0_27_ in range(nw215_.length(0)):
                nw215_[i0_27_] = init27_(i0_27_)
            d_1318_v32_ = nw215_
            d_1321_v33_: D2
            d_1321_v33_ = D2_DC5((self).f21, d_1318_v32_, -739, -140)
            d_1321_v33_ = d_1321_v33_
        elif True:
            d_1322___mcc_h6_ = source22_.cf37
            d_1323_cf37_ = d_1322___mcc_h6_
            d_1324_v34_: bool
            d_1324_v34_ = False
            d_1325_v35_: _dafny.Seq
            d_1325_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ty"))
            d_1326_v36_: _dafny.Array
            nw216_ = _dafny.Array(False, 13)
            d_1326_v36_ = nw216_
            d_1327_v37_: _dafny.Set
            d_1327_v37_ = _dafny.Set({d_1326_v36_})
            rhs186_ = (d_1324_v34_) == ((d_1327_v37_).ispropersubset(d_1327_v37_))
            rhs187_ = d_1325_v35_
            d_1324_v34_ = rhs186_
            d_1325_v35_ = rhs187_
            d_1328_v38_: _dafny.Map
            d_1328_v38_ = _dafny.Map({(self).f22: _dafny.CodePoint('e')})
            d_1329_v39_: _dafny.Map
            d_1329_v39_ = _dafny.Map({len(p1): (self).f21})
            d_1330_v41_: D1
            d_1330_v41_ = D1_DC1(d_1324_v34_)
            d_1331_v42_: _dafny.Map
            d_1331_v42_ = _dafny.Map({d_1330_v41_: (d_1330_v41_).cf1})
            d_1332_v43_: _dafny.Set
            def iife120_():
                coll64_ = _dafny.Map()
                compr_64_: D1
                for compr_64_ in (d_1331_v42_).keys.Elements:
                    d_1333_v40_: D1 = compr_64_
                    if (d_1333_v40_) in (d_1331_v42_):
                        coll64_[d_1333_v40_] = _dafny.CodePoint('u')
                return _dafny.Map(coll64_)
            d_1332_v43_ = _dafny.Set({len(iife120_()
            ), (self).f11, p0, p2})
            d_1334_v44_: _dafny.MultiSet
            d_1334_v44_ = _dafny.MultiSet([(self).f11, len(d_1332_v43_)])
            d_1335_v45_: D7
            d_1335_v45_ = D7_DC15(d_1334_v44_)
            d_1328_v38_ = (d_1328_v38_).set(len(d_1329_v39_), default__.fm39(d_1335_v45_, d_1324_v34_, globalState))
            (globalState).f3 = (self).f11
            d_1336_v46_: _dafny.Array
            nw217_ = _dafny.Array(int(0), 10)
            d_1336_v46_ = nw217_
            d_1337_v47_: _dafny.Set
            d_1337_v47_ = _dafny.Set({d_1336_v46_})
            d_1338_v48_: _dafny.Map
            d_1338_v48_ = _dafny.Map({(self).f21: d_1336_v46_})
            d_1337_v47_ = (d_1337_v47_) - (_dafny.Set({((d_1338_v48_)[(self).f21] if ((self).f21) in (d_1338_v48_) else d_1336_v46_), d_1336_v46_, d_1336_v46_, (D6_DC14(d_1336_v46_)).cf17, d_1336_v46_}))
        d_1339_v49_: _dafny.Map
        d_1339_v49_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f21])): (self).f22})
        d_1340_v50_: _dafny.Seq
        d_1340_v50_ = _dafny.SeqWithoutIsStrInference([(d_1339_v49_) | ((d_1339_v49_).set(p2, p0))])
        d_1341_v51_: _dafny.Map
        d_1341_v51_ = _dafny.Map({d_1254_v0_: not(False)})
        d_1340_v50_ = _dafny.SeqWithoutIsStrInference([((d_1340_v50_)[default__.safeIndex((_dafny.MultiSet([(self).f21, ((d_1341_v51_)[d_1254_v0_] if (d_1254_v0_) in (d_1341_v51_) else (self).f21)])).cardinality, len(d_1340_v50_))]) | (d_1339_v49_), d_1339_v49_])
        d_1255_v1_ = d_1255_v1_
        d_1342_v52_: bool
        d_1342_v52_ = False
        d_1343_v53_: _dafny.Set
        d_1343_v53_ = _dafny.Set({d_1339_v49_})
        d_1344_v54_: _dafny.Seq
        d_1344_v54_ = _dafny.SeqWithoutIsStrInference([d_1343_v53_, d_1343_v53_, d_1343_v53_, d_1343_v53_, d_1343_v53_])
        d_1342_v52_ = (d_1343_v53_).isdisjoint((d_1344_v54_)[default__.safeIndex(-426, len(d_1344_v54_))])
        if (self).f21:
            d_1345_v55_: _dafny.Seq
            d_1345_v55_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1346_i8_ in range(default__.abs(-795))])])
            d_1347_v56_: D9
            d_1347_v56_ = D9_DC20(d_1345_v55_)
            d_1347_v56_ = (d_1347_v56_ if (self).f21 else D9_DC20((d_1347_v56_).cf21))
            d_1348_v57_: D1
            d_1348_v57_ = D1_DC2()
            d_1349_v58_: D1
            d_1349_v58_ = D1_DC3(d_1348_v57_)
            d_1350_v59_: D1
            d_1350_v59_ = D1_DC3(d_1349_v58_)
            d_1351_v60_: _dafny.MultiSet
            d_1351_v60_ = _dafny.MultiSet([p2])
            d_1352_v61_: _dafny.Set
            d_1352_v61_ = _dafny.Set({d_1342_v52_})
            d_1353_v62_: _dafny.Array
            nw218_ = _dafny.Array(None, 29)
            nw218_[int(0)] = ((0) - (((d_1351_v60_)[p0] if (p0) in (d_1351_v60_) else (self).f11))) - ((self).f11)
            nw218_[int(1)] = (self).f11
            nw218_[int(2)] = (self).fm5(len(d_1352_v61_), p0, globalState)
            nw218_[int(3)] = (541) + (p0)
            nw218_[int(4)] = p2
            nw218_[int(5)] = ((self).f22) - (p0)
            nw218_[int(6)] = (0) - ((self).f11)
            nw218_[int(7)] = (p0) * ((self).f11)
            nw218_[int(8)] = (self).f11
            nw218_[int(9)] = p0
            nw218_[int(10)] = (0) - (len(d_1254_v0_))
            nw218_[int(11)] = default__.safeDivisionInt(-594, (self).f11)
            nw218_[int(12)] = (0) - (p0)
            nw218_[int(13)] = p2
            nw218_[int(14)] = (self).f22
            nw218_[int(15)] = (self).f22
            nw218_[int(16)] = (self).f11
            nw218_[int(17)] = len(d_1352_v61_)
            nw218_[int(18)] = default__.safeModuloInt(default__.fm0(len((self).f9), (self).f11, globalState), (self).f22)
            nw218_[int(19)] = p0
            nw218_[int(20)] = len(p1)
            nw218_[int(21)] = (self).f22
            nw218_[int(22)] = default__.safeModuloInt(p2, (self).f11)
            nw218_[int(23)] = (self).f11
            nw218_[int(24)] = p0
            nw218_[int(25)] = (0) - (len((p1) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1354_i9_ in range(default__.abs(498))]))))
            nw218_[int(26)] = p2
            nw218_[int(27)] = default__.safeDivisionInt(2, (self).f11)
            nw218_[int(28)] = default__.fm0((self).f11, (self).f22, globalState)
            d_1353_v62_ = nw218_
            d_1355_v63_: _dafny.Seq
            d_1355_v63_ = _dafny.SeqWithoutIsStrInference([d_1353_v62_])
            rhs188_ = d_1353_v62_
            rhs189_ = d_1350_v59_
            rhs190_ = (d_1342_v52_) or ((self).f21)
            rhs191_ = len((d_1355_v63_) + (d_1355_v63_))
            lhs118_ = globalState
            r1 = rhs188_
            d_1350_v59_ = rhs189_
            d_1342_v52_ = rhs190_
            lhs118_.f3 = rhs191_
            d_1356_v64_: D7
            d_1356_v64_ = D7_DC15(d_1351_v60_)
            d_1357_v65_: str
            d_1357_v65_ = default__.fm39(d_1356_v64_, d_1342_v52_, globalState)
            source23_ = d_1357_v65_
            d_1358___mcc_h7_ = source23_
            d_1359_cf15_ = d_1358___mcc_h7_
            d_1360_v66_: _dafny.Array
            nw219_ = _dafny.Array(False, 8)
            d_1360_v66_ = nw219_
            index189_ = default__.safeIndex(583, (d_1360_v66_).length(0))
            (d_1360_v66_)[index189_] = d_1342_v52_
            index190_ = default__.safeIndex(583, (d_1360_v66_).length(0))
            (d_1360_v66_)[index190_] = (((self).f22) > ((self).f22)) and ((len(d_1254_v0_)) != (p2))
            d_1361_v67_: _dafny.Seq
            d_1361_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcdxudp"))
            d_1361_v67_ = ((p1) + ((self).f9)) + (_dafny.SeqWithoutIsStrInference([d_1359_cf15_ for d_1362_i10_ in range(default__.abs(36))]))
            index191_ = default__.safeIndex(583, (d_1360_v66_).length(0))
            (d_1360_v66_)[index191_] = (self).f21
            d_1363_v68_: _dafny.Seq
            d_1363_v68_ = _dafny.SeqWithoutIsStrInference([(d_1360_v66_)[default__.safeIndex(583, (d_1360_v66_).length(0))]])
            d_1364_v69_: D3
            d_1364_v69_ = D3_DC7((d_1363_v68_)[default__.safeIndex((self).f22, len(d_1363_v68_))])
            d_1365_v70_: D15
            d_1365_v70_ = D15_DC37(p0, d_1364_v69_, True)
            d_1366_v71_: C2
            nw220_ = C2()
            nw220_.ctor__((self).f11, (_dafny.Map({_dafny.Map({392: default__.fm0(75, p2, globalState)}): d_1339_v49_}) if (d_1365_v70_).cf51 else self.f12), ((self).f9).set(default__.safeIndex(p0, len((self).f9)), d_1359_cf15_), (self).f10)
            d_1366_v71_ = nw220_
            index192_ = default__.safeIndex(730, (d_1353_v62_).length(0))
            (d_1353_v62_)[index192_] = 769
            d_1367_v72_: _dafny.Array
            nw221_ = _dafny.Array(_dafny.Seq({}), 2)
            d_1367_v72_ = nw221_
            d_1368_v73_: str
            d_1368_v73_ = _dafny.CodePoint('w')
            index193_ = default__.safeIndex(730, (d_1353_v62_).length(0))
            rhs192_ = (p0) - (default__.safeDivisionInt(p0, p2))
            rhs193_ = d_1342_v52_
            rhs194_ = d_1367_v72_
            rhs195_ = d_1368_v73_
            rhs196_ = (self).f21
            lhs119_ = d_1353_v62_
            lhs120_ = default__.safeIndex(730, (d_1353_v62_).length(0))
            lhs119_[lhs120_] = rhs192_
            d_1342_v52_ = rhs193_
            d_1367_v72_ = rhs194_
            d_1368_v73_ = rhs195_
            d_1342_v52_ = rhs196_
            d_1369_v74_: C1
            nw222_ = C1()
            nw222_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddxoouvc"))), p2, (self).f11, (self.f12) | (self.f12), (p1) + (p1), _dafny.SeqWithoutIsStrInference([d_1254_v0_]))
            d_1369_v74_ = nw222_
        elif True:
            if (((self).f9) + ((self).f9)) < (p1):
                (globalState).f3 = (0) - (((self).f22) + ((0) - (p0)))
                d_1370_v75_: _dafny.Array
                nw223_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
                d_1370_v75_ = nw223_
                index194_ = default__.safeIndex(249, (d_1370_v75_).length(0))
                (d_1370_v75_)[index194_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1371_i11_ in range(default__.abs(121))])
                d_1372_v76_: _dafny.Array
                nw224_ = _dafny.Array(int(0), 17)
                d_1372_v76_ = nw224_
                index195_ = default__.safeIndex(605, (d_1372_v76_).length(0))
                (d_1372_v76_)[index195_] = (0) - (len(p1))
                d_1373_v77_: D4
                d_1373_v77_ = D4_DC11(d_1342_v52_, 354)
                index196_ = default__.safeIndex(249, (d_1370_v75_).length(0))
                index197_ = default__.safeIndex(605, (d_1372_v76_).length(0))
                rhs197_ = p1
                rhs198_ = len(_dafny.Set({(d_1373_v77_).cf13}))
                lhs121_ = d_1370_v75_
                lhs122_ = default__.safeIndex(249, (d_1370_v75_).length(0))
                lhs123_ = d_1372_v76_
                lhs124_ = default__.safeIndex(605, (d_1372_v76_).length(0))
                lhs121_[lhs122_] = rhs197_
                lhs123_[lhs124_] = rhs198_
                d_1342_v52_ = (d_1342_v52_ if d_1342_v52_ else d_1342_v52_)
                d_1374_v78_: C0
                nw225_ = C0()
                nw225_.ctor__(d_1342_v52_, d_1342_v52_, (self).f9, (self).f10)
                d_1374_v78_ = nw225_
                d_1375_v79_: D7
                d_1375_v79_ = D7_DC16()
                d_1376_v80_: _dafny.Array
                nw226_ = _dafny.Array(None, 16)
                nw226_[int(0)] = D7_DC16()
                nw226_[int(1)] = d_1375_v79_
                nw226_[int(2)] = d_1375_v79_
                nw226_[int(3)] = default__.fm50((self).fm29((0) - ((d_1372_v76_)[default__.safeIndex(605, (d_1372_v76_).length(0))]), d_1342_v52_, (self).f21, globalState), globalState)
                nw226_[int(4)] = d_1375_v79_
                nw226_[int(5)] = D7_DC16()
                nw226_[int(6)] = d_1375_v79_
                nw226_[int(7)] = d_1375_v79_
                nw226_[int(8)] = d_1375_v79_
                nw226_[int(9)] = d_1375_v79_
                nw226_[int(10)] = d_1375_v79_
                nw226_[int(11)] = d_1375_v79_
                nw226_[int(12)] = d_1375_v79_
                nw226_[int(13)] = d_1375_v79_
                nw226_[int(14)] = d_1375_v79_
                nw226_[int(15)] = d_1375_v79_
                d_1376_v80_ = nw226_
                d_1376_v80_ = d_1376_v80_
            elif True:
                d_1342_v52_ = d_1342_v52_
                d_1377_v81_: _dafny.Set
                d_1377_v81_ = _dafny.Set({default__.fm38(len(d_1254_v0_), globalState)})
                d_1378_v82_: _dafny.Map
                d_1378_v82_ = _dafny.Map({((self).f22) - ((0) - (len((p1).set(default__.safeIndex((0) - (p2), len(p1)), _dafny.CodePoint('e'))))): d_1377_v81_})
                d_1378_v82_ = (d_1378_v82_).set((self).f11, d_1377_v81_)
                d_1342_v52_ = (self).f21
                d_1379_v83_: D10
                d_1379_v83_ = D10_DC22(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f9])))
                d_1380_v84_: D10
                d_1380_v84_ = D10_DC24(d_1379_v83_)
                pat_let_tv26_ = d_1379_v83_
                def iife121_(_pat_let28_0):
                    def iife122_(d_1381_dt__update__tmp_h6_):
                        def iife123_(_pat_let29_0):
                            def iife124_(d_1382_dt__update_hcf28_h0_):
                                return D10_DC24(d_1382_dt__update_hcf28_h0_)
                            return iife124_(_pat_let29_0)
                        return iife123_(pat_let_tv26_)
                    return iife122_(_pat_let28_0)
                d_1380_v84_ = iife121_(d_1380_v84_)
                d_1383_v85_: C0
                nw227_ = C0()
                nw227_.ctor__((d_1377_v81_).issubset(d_1377_v81_), (self).f21, (self).f9, ((self).f10) + ((self).f10))
                d_1383_v85_ = nw227_
            d_1384_v86_: _dafny.Seq
            d_1384_v86_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycpuii"))])
            d_1385_v87_: _dafny.Set
            d_1385_v87_ = _dafny.Set({d_1342_v52_})
            d_1386_v88_: _dafny.MultiSet
            d_1386_v88_ = _dafny.MultiSet([(self).f21, d_1342_v52_])
            d_1387_v89_: _dafny.Map
            d_1387_v89_ = _dafny.Map({d_1386_v88_: d_1342_v52_})
            d_1388_v90_: _dafny.Map
            d_1388_v90_ = _dafny.Map({(self).f21: False})
            d_1389_v91_: _dafny.Array
            nw228_ = _dafny.Array(None, 23)
            nw228_[int(0)] = len((self).f9)
            nw228_[int(1)] = (self).f11
            nw228_[int(2)] = p0
            nw228_[int(3)] = ((self).f22 if d_1342_v52_ else 192)
            nw228_[int(4)] = (self).f22
            nw228_[int(5)] = (default__.fm0(len((d_1384_v86_)[default__.safeIndex((self).f11, len(d_1384_v86_))]), p0, globalState) if not(d_1342_v52_) else p0)
            nw228_[int(6)] = (self).f11
            nw228_[int(7)] = default__.safeDivisionInt(p0, (self).f11)
            nw228_[int(8)] = (len(d_1385_v87_)) + (len(d_1339_v49_))
            nw228_[int(9)] = (0) - ((self).f11)
            nw228_[int(10)] = (self).f22
            nw228_[int(11)] = (self).f22
            nw228_[int(12)] = ((self).f22) + (p2)
            nw228_[int(13)] = p2
            nw228_[int(14)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1390_i12_ in range(default__.abs(660))]))
            nw228_[int(15)] = p2
            nw228_[int(16)] = (self).f22
            nw228_[int(17)] = default__.fm0((self).f11, p0, globalState)
            nw228_[int(18)] = (self).fm29(len(d_1387_v89_), not((self).f21), default__.fm38(len(_dafny.SeqWithoutIsStrInference([(self).f22 for d_1391_i13_ in range(default__.abs(-786))])), globalState), globalState)
            nw228_[int(19)] = default__.safeDivisionInt(len(d_1388_v90_), p0)
            nw228_[int(20)] = 269
            nw228_[int(21)] = 974
            nw228_[int(22)] = default__.safeDivisionInt((self).f11, len((self).f9))
            d_1389_v91_ = nw228_
            index198_ = default__.safeIndex(870, (d_1389_v91_).length(0))
            (d_1389_v91_)[index198_] = ((self).f11) + ((self).f11)
            d_1392_v92_: _dafny.Map
            d_1392_v92_ = _dafny.Map({((d_1388_v90_)[True] if (True) in (d_1388_v90_) else (self).f21): (self).f11})
            d_1393_v93_: str
            d_1393_v93_ = _dafny.CodePoint('u')
            d_1394_v94_: _dafny.MultiSet
            d_1394_v94_ = _dafny.MultiSet([d_1393_v93_, d_1393_v93_, d_1393_v93_])
            d_1395_v95_: _dafny.MultiSet
            d_1395_v95_ = _dafny.MultiSet([p2])
            index199_ = default__.safeIndex(870, (d_1389_v91_).length(0))
            rhs199_ = (default__.fm0((self).f11, 395, globalState)) * ((self).f22)
            rhs200_ = (self).f11
            rhs201_ = (0) - (p2)
            rhs202_ = ((d_1339_v49_)[default__.safeDivisionInt((0) - (((d_1392_v92_)[d_1342_v52_] if (d_1342_v52_) in (d_1392_v92_) else len(d_1254_v0_))), p0)] if (default__.safeDivisionInt((0) - (((d_1392_v92_)[d_1342_v52_] if (d_1342_v52_) in (d_1392_v92_) else len(d_1254_v0_))), p0)) in (d_1339_v49_) else (p2) * (((d_1394_v94_)[d_1393_v93_] if (d_1393_v93_) in (d_1394_v94_) else (d_1395_v95_).cardinality)))
            lhs125_ = globalState
            lhs126_ = globalState
            lhs127_ = d_1389_v91_
            lhs128_ = default__.safeIndex(870, (d_1389_v91_).length(0))
            lhs129_ = globalState
            lhs125_.f3 = rhs199_
            lhs126_.f3 = rhs200_
            lhs127_[lhs128_] = rhs201_
            lhs129_.f3 = rhs202_
            index200_ = default__.safeIndex(870, (d_1389_v91_).length(0))
            (d_1389_v91_)[index200_] = (default__.safeModuloInt(p0, (self).f11)) - ((self).f22)
            d_1393_v93_ = (d_1393_v93_ if (956) < ((self).fm5((self).f22, (self).f11, globalState)) else d_1393_v93_)
            if default__.fm38(len(p1), globalState):
                d_1342_v52_ = (self).f21
                d_1396_v96_: C2
                nw229_ = C2()
                nw229_.ctor__(p2, self.f12, (self).f9, (self).f10)
                d_1396_v96_ = nw229_
                (globalState).f3 = (d_1389_v91_)[default__.safeIndex(870, (d_1389_v91_).length(0))]
                d_1397_v97_: _dafny.Map
                d_1397_v97_ = _dafny.Map({p2: (self).f9})
                d_1398_v98_: _dafny.Map
                d_1398_v98_ = _dafny.Map({((d_1397_v97_)[p2] if (p2) in (d_1397_v97_) else default__.fm31(globalState)): default__.fm44(d_1342_v52_, ((d_1339_v49_)[(d_1389_v91_)[default__.safeIndex(870, (d_1389_v91_).length(0))]] if ((d_1389_v91_)[default__.safeIndex(870, (d_1389_v91_).length(0))]) in (d_1339_v49_) else (self).f22), d_1342_v52_, globalState)})
                index201_ = default__.safeIndex(870, (d_1389_v91_).length(0))
                (d_1389_v91_)[index201_] = len(_dafny.Map({d_1342_v52_: len(((d_1398_v98_)[(self).f9] if ((self).f9) in (d_1398_v98_) else d_1254_v0_))}))
                d_1399_v99_: _dafny.Map
                d_1399_v99_ = _dafny.Map({d_1393_v93_: p0})
                (globalState).f3 = ((d_1399_v99_)[d_1393_v93_] if (d_1393_v93_) in (d_1399_v99_) else (len(d_1385_v87_)) - (len((self).f9)))
            elif True:
                index202_ = default__.safeIndex(870, (d_1389_v91_).length(0))
                (d_1389_v91_)[index202_] = (888) + (p0)
                d_1400_v100_: _dafny.Map
                d_1400_v100_ = _dafny.Map({(d_1342_v52_ if d_1342_v52_ else (self).f21): (p1) + ((self).f9)})
                d_1401_v101_: _dafny.Seq
                d_1401_v101_ = _dafny.SeqWithoutIsStrInference([not((self).f21)])
                d_1400_v100_ = (d_1400_v100_).set((d_1401_v101_)[default__.safeIndex(p0, len(d_1401_v101_))], (self).f9)
                d_1392_v92_ = (default__.fm42(d_1392_v92_, d_1342_v52_, len(d_1388_v90_), d_1342_v52_, globalState)).set(d_1342_v52_, ((self).f22 if (self).f21 else 912))
                d_1402_v102_: _dafny.Array
                nw230_ = _dafny.Array(None, 11)
                nw230_[int(0)] = d_1393_v93_
                nw230_[int(1)] = d_1393_v93_
                nw230_[int(2)] = d_1393_v93_
                nw230_[int(3)] = d_1393_v93_
                nw230_[int(4)] = d_1393_v93_
                nw230_[int(5)] = d_1393_v93_
                nw230_[int(6)] = d_1393_v93_
                nw230_[int(7)] = d_1393_v93_
                nw230_[int(8)] = ((self).f9)[default__.safeIndex((self).f11, len((self).f9))]
                nw230_[int(9)] = d_1393_v93_
                nw230_[int(10)] = d_1393_v93_
                d_1402_v102_ = nw230_
                index203_ = default__.safeIndex(566, (d_1402_v102_).length(0))
                (d_1402_v102_)[index203_] = d_1393_v93_
                index204_ = default__.safeIndex(566, (d_1402_v102_).length(0))
                (d_1402_v102_)[index204_] = d_1393_v93_
                (globalState).f3 = default__.safeDivisionInt((0) - (p2), ((self).f22) + (len(p1)))
        d_1403_v103_: str
        d_1403_v103_ = _dafny.CodePoint('u')
        d_1404_v104_: str
        d_1404_v104_ = d_1403_v103_
        def lambda52_(source24_):
            d_1405___mcc_h8_ = source24_
            d_1406_cf15_ = d_1405___mcc_h8_
            return (self).f21

        d_1342_v52_ = lambda52_(d_1404_v104_)
        d_1407_v105_: _dafny.Array
        def lambda53_(d_1408_i14_):
            return (self).f21

        init28_ = lambda53_
        nw231_ = _dafny.Array(None, 5)
        for i0_28_ in range(nw231_.length(0)):
            nw231_[i0_28_] = init28_(i0_28_)
        d_1407_v105_ = nw231_
        d_1409_v106_: _dafny.Array
        nw232_ = _dafny.Array(None, 3)
        nw232_[int(0)] = d_1407_v105_
        nw232_[int(1)] = d_1407_v105_
        nw232_[int(2)] = d_1407_v105_
        d_1409_v106_ = nw232_
        r0 = d_1409_v106_
        d_1410_v107_: _dafny.Seq
        d_1410_v107_ = _dafny.SeqWithoutIsStrInference([d_1342_v52_, d_1342_v52_, d_1342_v52_, d_1342_v52_])
        d_1411_v108_: _dafny.Map
        d_1411_v108_ = _dafny.Map({d_1410_v107_: p2})
        d_1412_v109_: _dafny.Seq
        d_1412_v109_ = _dafny.SeqWithoutIsStrInference([(self).f9])
        d_1413_v110_: _dafny.Map
        d_1413_v110_ = _dafny.Map({p1: (self).f11})
        d_1414_v111_: _dafny.Array
        nw233_ = _dafny.Array(None, 28)
        nw233_[int(0)] = (self).f22
        nw233_[int(1)] = default__.safeDivisionInt(p0, p2)
        nw233_[int(2)] = p2
        nw233_[int(3)] = len(d_1254_v0_)
        nw233_[int(4)] = len(_dafny.SeqWithoutIsStrInference([p2]))
        nw233_[int(5)] = default__.safeDivisionInt(((d_1339_v49_)[len(_dafny.Map({False: (self).f21}))] if (len(_dafny.Map({False: (self).f21}))) in (d_1339_v49_) else len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f21]))])) for d_1415_i15_ in range(default__.abs(88))]))), (0) - ((self).f11))
        nw233_[int(6)] = 222
        nw233_[int(7)] = (self).f11
        nw233_[int(8)] = (((d_1411_v108_)[d_1410_v107_] if (d_1410_v107_) in (d_1411_v108_) else p2) if False else p2)
        nw233_[int(9)] = (self).f11
        nw233_[int(10)] = p0
        nw233_[int(11)] = default__.fm0(p0, p2, globalState)
        nw233_[int(12)] = len((d_1412_v109_) + (d_1412_v109_))
        nw233_[int(13)] = (0) - ((self).f22)
        nw233_[int(14)] = (self).f22
        nw233_[int(15)] = p0
        nw233_[int(16)] = (self).f11
        nw233_[int(17)] = p0
        nw233_[int(18)] = ((d_1413_v110_)[(self).f9] if ((self).f9) in (d_1413_v110_) else (self).f22)
        nw233_[int(19)] = p0
        nw233_[int(20)] = p2
        nw233_[int(21)] = p0
        nw233_[int(22)] = 593
        nw233_[int(23)] = (p2) - (len((d_1340_v50_)[default__.safeIndex((self).f11, len(d_1340_v50_))]))
        nw233_[int(24)] = len(d_1410_v107_)
        nw233_[int(25)] = (d_1254_v0_)[default__.safeIndex((self).f11, len(d_1254_v0_))]
        nw233_[int(26)] = p0
        nw233_[int(27)] = p2
        d_1414_v111_ = nw233_
        r1 = d_1414_v111_
        return r0, r1

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        r1 = not (True) or ((self).f21)
        d_1416_v0_: _dafny.Array
        nw234_ = _dafny.Array(int(0), 13)
        d_1416_v0_ = nw234_
        index205_ = default__.safeIndex(676, (d_1416_v0_).length(0))
        (d_1416_v0_)[index205_] = default__.fm0(p2, p2, globalState)
        index206_ = default__.safeIndex(676, (d_1416_v0_).length(0))
        rhs203_ = (self).f21
        rhs204_ = p0
        lhs130_ = d_1416_v0_
        lhs131_ = default__.safeIndex(676, (d_1416_v0_).length(0))
        r0 = rhs203_
        lhs130_[lhs131_] = rhs204_
        hi2_ = (self).f11
        for d_1417_i0_ in range((self).f22, hi2_):
            d_1418_v1_: D6
            d_1418_v1_ = D6_DC14(d_1416_v0_)
            d_1419_v2_: _dafny.Array
            nw235_ = _dafny.Array(None, 2)
            nw235_[int(0)] = d_1418_v1_
            nw235_[int(1)] = D6_DC14(d_1416_v0_)
            d_1419_v2_ = nw235_
            d_1420_v3_: _dafny.Array
            def lambda54_(d_1421_i1_):
                return False

            init29_ = lambda54_
            nw236_ = _dafny.Array(None, 1)
            for i0_29_ in range(nw236_.length(0)):
                nw236_[i0_29_] = init29_(i0_29_)
            d_1420_v3_ = nw236_
            d_1422_v4_: _dafny.MultiSet
            d_1422_v4_ = _dafny.MultiSet([(self).f21])
            index207_ = default__.safeIndex(406, (d_1420_v3_).length(0))
            (d_1420_v3_)[index207_] = default__.fm38((d_1422_v4_).cardinality, globalState)
            index208_ = default__.safeIndex(406, (d_1420_v3_).length(0))
            rhs205_ = d_1419_v2_
            rhs206_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndeaoaq")))) != ((d_1416_v0_)[default__.safeIndex(676, (d_1416_v0_).length(0))])
            lhs132_ = d_1420_v3_
            lhs133_ = default__.safeIndex(406, (d_1420_v3_).length(0))
            d_1419_v2_ = rhs205_
            lhs132_[lhs133_] = rhs206_
            if (d_1420_v3_)[default__.safeIndex(406, (d_1420_v3_).length(0))]:
                d_1423_v5_: _dafny.Set
                d_1423_v5_ = _dafny.Set({(d_1420_v3_)[default__.safeIndex(406, (d_1420_v3_).length(0))]})
                d_1424_v6_: _dafny.Map
                d_1424_v6_ = _dafny.Map({(self).f21: True})
                d_1425_v7_: _dafny.Map
                d_1425_v7_ = _dafny.Map({(self).f21: (d_1424_v6_).set((d_1420_v3_)[default__.safeIndex(406, (d_1420_v3_).length(0))], (d_1420_v3_)[default__.safeIndex(406, (d_1420_v3_).length(0))])})
                d_1426_v8_: _dafny.Seq
                d_1426_v8_ = _dafny.SeqWithoutIsStrInference([d_1425_v7_])
                index209_ = default__.safeIndex(676, (d_1416_v0_).length(0))
                rhs207_ = (0) - ((0) - ((p1 if (p1) < (p2) else p0)))
                rhs208_ = d_1423_v5_
                rhs209_ = len(((d_1426_v8_)[default__.safeIndex(p2, len(d_1426_v8_))]).set((d_1420_v3_)[default__.safeIndex(406, (d_1420_v3_).length(0))], (d_1424_v6_) | (d_1424_v6_)))
                rhs210_ = (self).f11
                lhs134_ = globalState
                lhs135_ = d_1416_v0_
                lhs136_ = default__.safeIndex(676, (d_1416_v0_).length(0))
                lhs137_ = globalState
                lhs134_.f3 = rhs207_
                d_1423_v5_ = rhs208_
                lhs135_[lhs136_] = rhs209_
                lhs137_.f3 = rhs210_
                d_1424_v6_ = _dafny.Map({not ((self).f21) or (not(not((self).f21))): False})
                d_1427_v9_: _dafny.Map
                d_1427_v9_ = _dafny.Map({d_1424_v6_: _dafny.Set({False, (self).f21})})
                d_1428_v10_: str
                d_1428_v10_ = _dafny.CodePoint('g')
                d_1423_v5_ = ((d_1427_v9_)[d_1424_v6_] if (d_1424_v6_) in (d_1427_v9_) else default__.fm51((self).f21, p2, d_1428_v10_, (self).f21, globalState))
                d_1429_v11_: _dafny.Seq
                d_1429_v11_ = _dafny.SeqWithoutIsStrInference([(self).f21])
                index210_ = default__.safeIndex(406, (d_1420_v3_).length(0))
                (d_1420_v3_)[index210_] = default__.fm38(len(d_1429_v11_), globalState)
                d_1430_v12_: _dafny.Array
                nw237_ = _dafny.Array(_dafny.Map({}), 26)
                d_1430_v12_ = nw237_
                d_1431_v14_: _dafny.Seq
                d_1431_v14_ = _dafny.SeqWithoutIsStrInference([d_1424_v6_])
                d_1432_v15_: _dafny.Map
                def iife125_():
                    coll65_ = _dafny.Map()
                    compr_65_: _dafny.Map
                    for compr_65_ in (d_1431_v14_).Elements:
                        d_1433_v13_: _dafny.Map = compr_65_
                        if (d_1433_v13_) in (d_1431_v14_):
                            coll65_[d_1433_v13_] = p0
                    return _dafny.Map(coll65_)
                d_1432_v15_ = _dafny.Map({(d_1416_v0_)[default__.safeIndex(676, (d_1416_v0_).length(0))]: iife125_()
                })
                d_1434_v16_: _dafny.Map
                d_1434_v16_ = _dafny.Map({d_1424_v6_: len(_dafny.SeqWithoutIsStrInference([d_1428_v10_ for d_1435_i2_ in range(default__.abs(-587))]))})
                index211_ = default__.safeIndex(687, (d_1430_v12_).length(0))
                (d_1430_v12_)[index211_] = ((d_1432_v15_)[(self).f11] if ((self).f11) in (d_1432_v15_) else d_1434_v16_)
                index212_ = default__.safeIndex(687, (d_1430_v12_).length(0))
                (d_1430_v12_)[index212_] = d_1434_v16_
            elif True:
                d_1436_v17_: str
                d_1436_v17_ = _dafny.CodePoint('y')
                d_1436_v17_ = d_1436_v17_
                (globalState).f3 = p2
                r0 = (d_1420_v3_)[default__.safeIndex(406, (d_1420_v3_).length(0))]
                d_1437_v18_: _dafny.Seq
                d_1437_v18_ = _dafny.SeqWithoutIsStrInference([default__.fm0(p1, d_1417_i0_, globalState), p0, p0])
                d_1438_v19_: C2
                nw238_ = C2()
                nw238_.ctor__((self).f22, self.f12, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uarld")), _dafny.SeqWithoutIsStrInference([d_1437_v18_]))
                d_1438_v19_ = nw238_
                d_1439_v20_: D1
                d_1439_v20_ = D1_DC2()
                d_1440_v21_: D1
                d_1440_v21_ = D1_DC3(d_1439_v20_)
                d_1441_v22_: D1
                d_1441_v22_ = D1_DC3(d_1440_v21_)
                d_1442_v23_: D1
                d_1442_v23_ = D1_DC3(d_1441_v22_)
                d_1443_v24_: D1
                d_1443_v24_ = D1_DC3(d_1440_v21_)
                d_1444_v25_: _dafny.Map
                d_1444_v25_ = _dafny.Map({d_1443_v24_: d_1438_v19_})
                d_1445_v26_: _dafny.Seq
                d_1445_v26_ = _dafny.SeqWithoutIsStrInference([d_1438_v19_])
                d_1446_v27_: _dafny.Array
                nw239_ = _dafny.Array(None, 27)
                nw239_[int(0)] = d_1438_v19_
                nw239_[int(1)] = d_1438_v19_
                nw239_[int(2)] = d_1438_v19_
                nw239_[int(3)] = d_1438_v19_
                nw239_[int(4)] = d_1438_v19_
                nw239_[int(5)] = d_1438_v19_
                nw239_[int(6)] = d_1438_v19_
                nw239_[int(7)] = d_1438_v19_
                nw239_[int(8)] = d_1438_v19_
                nw239_[int(9)] = d_1438_v19_
                nw239_[int(10)] = ((d_1444_v25_)[d_1443_v24_] if (d_1443_v24_) in (d_1444_v25_) else d_1438_v19_)
                nw239_[int(11)] = d_1438_v19_
                nw239_[int(12)] = d_1438_v19_
                nw239_[int(13)] = d_1438_v19_
                nw239_[int(14)] = (d_1438_v19_ if (self).f21 else d_1438_v19_)
                nw239_[int(15)] = d_1438_v19_
                nw239_[int(16)] = d_1438_v19_
                nw239_[int(17)] = d_1438_v19_
                nw239_[int(18)] = d_1438_v19_
                nw239_[int(19)] = d_1438_v19_
                nw239_[int(20)] = d_1438_v19_
                nw239_[int(21)] = (d_1445_v26_)[default__.safeIndex(p2, len(d_1445_v26_))]
                nw239_[int(22)] = d_1438_v19_
                nw239_[int(23)] = d_1438_v19_
                nw239_[int(24)] = d_1438_v19_
                nw239_[int(25)] = d_1438_v19_
                nw239_[int(26)] = d_1438_v19_
                d_1446_v27_ = nw239_
                index213_ = default__.safeIndex(927, (d_1446_v27_).length(0))
                (d_1446_v27_)[index213_] = d_1438_v19_
                index214_ = default__.safeIndex(927, (d_1446_v27_).length(0))
                (d_1446_v27_)[index214_] = d_1438_v19_
                index215_ = default__.safeIndex(406, (d_1420_v3_).length(0))
                (d_1420_v3_)[index215_] = (d_1420_v3_)[default__.safeIndex(406, (d_1420_v3_).length(0))]
            d_1447_v28_: str
            d_1447_v28_ = _dafny.CodePoint('b')
            d_1448_v29_: str
            d_1448_v29_ = d_1447_v28_
            d_1448_v29_ = d_1447_v28_
            d_1449_v30_: _dafny.Map
            d_1449_v30_ = _dafny.Map({d_1447_v28_: (self).f21})
            r0 = ((d_1449_v30_)[d_1447_v28_] if (d_1447_v28_) in (d_1449_v30_) else (self).f21)
        d_1450_v31_: _dafny.MultiSet
        d_1450_v31_ = _dafny.MultiSet([(self).f22, p2])
        d_1451_v32_: _dafny.Map
        d_1451_v32_ = _dafny.Map({(d_1450_v31_).issubset(_dafny.MultiSet([p1, (d_1416_v0_)[default__.safeIndex(676, (d_1416_v0_).length(0))]])): (self).f22})
        d_1452_v33_: D17
        d_1452_v33_ = D17_DC39(d_1451_v32_)
        d_1451_v32_ = ((d_1452_v33_ if (self).f21 else d_1452_v33_)).cf53
        (globalState).f3 = (0) - ((-118) + ((self).fm29(len((self).f9), (self).f21, False, globalState)))
        (globalState).f3 = 800
        r0 = (self).f21
        d_1453_v34_: _dafny.Map
        d_1453_v34_ = _dafny.Map({(self).f21: (self).f21})
        d_1454_v35_: _dafny.Set
        d_1454_v35_ = _dafny.Set({len(d_1453_v34_), (0) - ((0) - (p1))})
        r1 = (d_1454_v35_).ispropersubset((d_1454_v35_ if (self).f21 else d_1454_v35_))
        d_1455_v36_: _dafny.Seq
        d_1455_v36_ = _dafny.SeqWithoutIsStrInference([(self).f21, (self).f21])
        r2 = (d_1416_v0_ if (d_1455_v36_)[default__.safeIndex((0) - (p2), len(d_1455_v36_))] else (d_1416_v0_ if (self).f21 else d_1416_v0_))
        r3 = len(_dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(len((self).f9), p0), (d_1416_v0_)[default__.safeIndex(676, (d_1416_v0_).length(0))], p1, default__.safeModuloInt(438, (d_1416_v0_)[default__.safeIndex(676, (d_1416_v0_).length(0))])]))
        return r0, r1, r2, r3

    @property
    def f21(self):
        return self._f21
    @property
    def f22(self):
        return self._f22

class C4(T0):
    def  __init__(self):
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: _dafny.Seq = _dafny.Seq({})
        self.f19: bool = False
        self.f20: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f19, f20, f9, f10):
        (self).f19 = f19
        (self).f20 = f20
        (self)._f9 = f9
        (self)._f10 = f10

    def fm5(self, p0, p1, globalState):
        return default__.safeDivisionInt((len((self).f9)) + (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({579: self.f20}), _dafny.Map({len(_dafny.Map({self.f20: self.f20})): self.f19})]))), (301) - (692))

    def fm24(self, p0, p1, p2, globalState):
        source25_ = D2_DC4(len((self).f9))
        if source25_.is_DC5:
            d_1456___mcc_h0_ = source25_.cf4
            d_1457___mcc_h1_ = source25_.cf5
            d_1458___mcc_h2_ = source25_.cf6
            d_1459___mcc_h3_ = source25_.cf7
            d_1460_cf7_ = d_1459___mcc_h3_
            d_1461_cf6_ = d_1458___mcc_h2_
            d_1462_cf5_ = d_1457___mcc_h1_
            d_1463_cf4_ = d_1456___mcc_h0_
            return d_1461_cf6_
        elif True:
            d_1464___mcc_h4_ = source25_.cf3
            d_1465_cf3_ = d_1464___mcc_h4_
            return d_1465_cf3_

    def fm25(self, p0, p1, p2, globalState):
        return not(not(self.f19))

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_1466_v0_: _dafny.Map
        d_1466_v0_ = _dafny.Map({self.f20: _dafny.Set({self.f19, self.f19, self.f19})})
        if ((d_1466_v0_) | (d_1466_v0_)) == (d_1466_v0_):
            d_1467_v1_: _dafny.Array
            def lambda55_(d_1468_i0_):
                return self.f20

            init30_ = lambda55_
            nw240_ = _dafny.Array(None, 1)
            for i0_30_ in range(nw240_.length(0)):
                nw240_[i0_30_] = init30_(i0_30_)
            d_1467_v1_ = nw240_
            d_1469_v2_: _dafny.Array
            nw241_ = _dafny.Array(None, 16)
            nw241_[int(0)] = d_1467_v1_
            nw241_[int(1)] = d_1467_v1_
            nw241_[int(2)] = d_1467_v1_
            nw241_[int(3)] = d_1467_v1_
            nw241_[int(4)] = d_1467_v1_
            nw241_[int(5)] = d_1467_v1_
            nw241_[int(6)] = d_1467_v1_
            nw241_[int(7)] = d_1467_v1_
            nw241_[int(8)] = d_1467_v1_
            nw241_[int(9)] = d_1467_v1_
            nw241_[int(10)] = d_1467_v1_
            nw241_[int(11)] = d_1467_v1_
            nw241_[int(12)] = d_1467_v1_
            nw241_[int(13)] = d_1467_v1_
            nw241_[int(14)] = d_1467_v1_
            nw241_[int(15)] = d_1467_v1_
            d_1469_v2_ = nw241_
            d_1470_v3_: bool
            d_1471_v4_: bool
            out51_: bool
            out52_: bool
            out51_, out52_ = (self).m14(d_1469_v2_, globalState)
            d_1470_v3_ = out51_
            d_1471_v4_ = out52_
            d_1472_v5_: _dafny.Array
            def lambda56_(d_1473_p1_):
                def lambda57_(d_1474_i1_):
                    return default__.safeDivisionInt(d_1474_i1_, d_1473_p1_)

                return lambda57_

            init31_ = lambda56_(p1)
            nw242_ = _dafny.Array(None, 27)
            for i0_31_ in range(nw242_.length(0)):
                nw242_[i0_31_] = init31_(i0_31_)
            d_1472_v5_ = nw242_
            d_1475_v6_: _dafny.Map
            d_1475_v6_ = _dafny.Map({d_1471_v4_: d_1472_v5_})
            d_1476_v7_: _dafny.Map
            d_1476_v7_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1470_v3_, True]): self.f20})
            d_1475_v6_ = (d_1475_v6_).set(((d_1476_v7_)[_dafny.SeqWithoutIsStrInference([d_1471_v4_, self.f19])] if (_dafny.SeqWithoutIsStrInference([d_1471_v4_, self.f19])) in (d_1476_v7_) else d_1470_v3_), (d_1472_v5_ if True else d_1472_v5_))
            d_1471_v4_ = self.f20
            d_1470_v3_ = self.f19
            index216_ = default__.safeIndex(750, (d_1467_v1_).length(0))
            (d_1467_v1_)[index216_] = d_1470_v3_
            index217_ = default__.safeIndex(750, (d_1467_v1_).length(0))
            (d_1467_v1_)[index217_] = d_1470_v3_
        elif True:
            r3 = -389
            d_1477_v8_: _dafny.Seq
            d_1477_v8_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_1478_v9_: _dafny.Array
            nw243_ = _dafny.Array(None, 23)
            nw243_[int(0)] = True
            nw243_[int(1)] = self.f20
            nw243_[int(2)] = self.f20
            nw243_[int(3)] = self.f20
            nw243_[int(4)] = self.f20
            nw243_[int(5)] = True
            nw243_[int(6)] = self.f20
            nw243_[int(7)] = self.f20
            nw243_[int(8)] = self.f19
            nw243_[int(9)] = self.f19
            nw243_[int(10)] = self.f19
            nw243_[int(11)] = self.f20
            nw243_[int(12)] = self.f20
            nw243_[int(13)] = self.f19
            nw243_[int(14)] = self.f20
            nw243_[int(15)] = True
            nw243_[int(16)] = self.f20
            nw243_[int(17)] = self.f19
            nw243_[int(18)] = (self).fm25(self.f20, self.f19, _dafny.Set({d_1477_v8_}), globalState)
            nw243_[int(19)] = not(self.f20)
            nw243_[int(20)] = self.f19
            nw243_[int(21)] = self.f20
            nw243_[int(22)] = self.f20
            d_1478_v9_ = nw243_
            d_1479_v10_: D2
            d_1479_v10_ = D2_DC5(self.f19, d_1478_v9_, p2, p2)
            pat_let_tv27_ = p1
            def iife126_(_pat_let30_0):
                def iife127_(d_1480_dt__update__tmp_h0_):
                    def iife128_(_pat_let31_0):
                        def iife129_(d_1481_dt__update_hcf7_h0_):
                            def iife130_(_pat_let32_0):
                                def iife131_(d_1482_dt__update_hcf4_h0_):
                                    return D2_DC5(d_1482_dt__update_hcf4_h0_, (d_1480_dt__update__tmp_h0_).cf5, (d_1480_dt__update__tmp_h0_).cf6, d_1481_dt__update_hcf7_h0_)
                                return iife131_(_pat_let32_0)
                            return iife130_(False)
                        return iife129_(_pat_let31_0)
                    return iife128_(default__.safeDivisionInt(947, pat_let_tv27_))
                return iife127_(_pat_let30_0)
            source26_ = iife126_(d_1479_v10_)
            if source26_.is_DC5:
                d_1483___mcc_h0_ = source26_.cf4
                d_1484___mcc_h1_ = source26_.cf5
                d_1485___mcc_h2_ = source26_.cf6
                d_1486___mcc_h3_ = source26_.cf7
                d_1487_cf7_ = d_1486___mcc_h3_
                d_1488_cf6_ = d_1485___mcc_h2_
                d_1489_cf5_ = d_1484___mcc_h1_
                d_1490_cf4_ = d_1483___mcc_h0_
                d_1491_v11_: _dafny.Seq
                d_1491_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sayhhm"))
                d_1491_v11_ = (self).f9
                d_1492_v12_: str
                d_1492_v12_ = _dafny.CodePoint('t')
                d_1493_v13_: _dafny.Seq
                d_1493_v13_ = _dafny.SeqWithoutIsStrInference([self.f19])
                d_1494_v14_: _dafny.Seq
                d_1494_v14_ = _dafny.SeqWithoutIsStrInference([d_1493_v13_])
                d_1495_v15_: _dafny.Map
                d_1495_v15_ = _dafny.Map({d_1492_v12_: (d_1494_v14_) + (_dafny.SeqWithoutIsStrInference([d_1493_v13_, _dafny.SeqWithoutIsStrInference([d_1490_cf4_, self.f20])]))})
                d_1495_v15_ = (d_1495_v15_).set(_dafny.CodePoint('h'), d_1494_v14_)
                d_1496_v16_: _dafny.Map
                d_1496_v16_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "timmqvcek"))): d_1491_v11_})
                d_1491_v11_ = (d_1491_v11_ if self.f20 else default__.fm26(p1, d_1496_v16_, globalState))
                d_1497_v17_: _dafny.MultiSet
                d_1497_v17_ = _dafny.MultiSet([p0])
                d_1498_v18_: D7
                d_1498_v18_ = D7_DC15(d_1497_v17_)
                d_1498_v18_ = (d_1498_v18_ if self.f20 else d_1498_v18_)
            elif True:
                d_1499___mcc_h4_ = source26_.cf3
                d_1500_cf3_ = d_1499___mcc_h4_
                d_1501_v19_: _dafny.Map
                d_1501_v19_ = _dafny.Map({self.f20: True})
                d_1501_v19_ = d_1501_v19_
                d_1502_v20_: str
                d_1502_v20_ = _dafny.CodePoint('u')
                d_1503_v21_: _dafny.Set
                d_1503_v21_ = _dafny.Set({self.f20})
                d_1504_v22_: _dafny.Map
                d_1504_v22_ = _dafny.Map({d_1502_v20_: p0})
                d_1505_v23_: _dafny.MultiSet
                d_1505_v23_ = _dafny.MultiSet([True])
                d_1506_v24_: _dafny.Map
                d_1506_v24_ = _dafny.Map({self.f19: d_1500_cf3_})
                d_1507_v25_: _dafny.Seq
                d_1507_v25_ = _dafny.SeqWithoutIsStrInference([self.f20, self.f19])
                d_1508_v26_: D8
                d_1508_v26_ = D8_DC17(d_1507_v25_)
                d_1509_v27_: _dafny.Array
                nw244_ = _dafny.Array(None, 21)
                nw244_[int(0)] = len((d_1503_v21_) | (d_1503_v21_))
                nw244_[int(1)] = (2) * (d_1500_cf3_)
                nw244_[int(2)] = ((d_1504_v22_)[d_1502_v20_] if (d_1502_v20_) in (d_1504_v22_) else p1)
                nw244_[int(3)] = ((d_1505_v23_)[self.f19] if (self.f19) in (d_1505_v23_) else d_1500_cf3_)
                nw244_[int(4)] = p2
                nw244_[int(5)] = p0
                nw244_[int(6)] = 153
                nw244_[int(7)] = default__.fm0(p2, len(_dafny.Map({p0: True})), globalState)
                nw244_[int(8)] = default__.safeModuloInt((0) - (((d_1506_v24_)[self.f19] if (self.f19) in (d_1506_v24_) else d_1500_cf3_)), p1)
                nw244_[int(9)] = 148
                nw244_[int(10)] = p0
                nw244_[int(11)] = ((d_1505_v23_) | (_dafny.MultiSet((d_1508_v26_).cf19))).cardinality
                nw244_[int(12)] = 578
                nw244_[int(13)] = d_1500_cf3_
                nw244_[int(14)] = p0
                nw244_[int(15)] = (p2) + (p0)
                nw244_[int(16)] = d_1500_cf3_
                nw244_[int(17)] = (p0) * (p2)
                nw244_[int(18)] = d_1500_cf3_
                nw244_[int(19)] = default__.fm0(d_1500_cf3_, p1, globalState)
                nw244_[int(20)] = d_1500_cf3_
                d_1509_v27_ = nw244_
                index218_ = default__.safeIndex(287, (d_1509_v27_).length(0))
                (d_1509_v27_)[index218_] = p2
                d_1510_v28_: _dafny.Map
                d_1510_v28_ = _dafny.Map({(0) - (p1): p0})
                index219_ = default__.safeIndex(287, (d_1509_v27_).length(0))
                rhs211_ = default__.fm27(globalState)
                rhs212_ = (0) - (((d_1510_v28_)[p2] if (p2) in (d_1510_v28_) else p0))
                lhs138_ = d_1509_v27_
                lhs139_ = default__.safeIndex(287, (d_1509_v27_).length(0))
                d_1502_v20_ = rhs211_
                lhs138_[lhs139_] = rhs212_
                (self).f19 = ((self).fm24(p2, self.f19, d_1502_v20_, globalState)) <= (-483)
                (self).f19 = ((self).f9) != ((self).f9)
            d_1511_v29_: _dafny.Set
            d_1511_v29_ = _dafny.Set({p0, p2})
            d_1512_v30_: _dafny.Array
            nw245_ = _dafny.Array(None, 7)
            nw245_[int(0)] = p0
            nw245_[int(1)] = p0
            nw245_[int(2)] = p0
            nw245_[int(3)] = len(d_1511_v29_)
            nw245_[int(4)] = p2
            nw245_[int(5)] = p0
            nw245_[int(6)] = p2
            d_1512_v30_ = nw245_
            d_1513_v31_: _dafny.Seq
            d_1513_v31_ = _dafny.SeqWithoutIsStrInference([d_1512_v30_])
            d_1514_v32_: _dafny.Map
            d_1514_v32_ = _dafny.Map({(d_1513_v31_)[default__.safeIndex(p1, len(d_1513_v31_))]: d_1511_v29_})
            d_1515_v33_: str
            d_1515_v33_ = _dafny.CodePoint('f')
            d_1516_v34_: _dafny.Set
            d_1516_v34_ = _dafny.Set({self.f19})
            d_1517_v35_: _dafny.MultiSet
            d_1517_v35_ = _dafny.MultiSet([p1])
            d_1518_v36_: D2
            d_1518_v36_ = D2_DC4(p2)
            nw246_ = _dafny.Array(None, 20)
            nw246_[int(0)] = (p0) + (p0)
            nw246_[int(1)] = len(((self).f10)[default__.safeIndex(len(((d_1514_v32_)[d_1512_v30_] if (d_1512_v30_) in (d_1514_v32_) else d_1511_v29_)), len((self).f10))])
            nw246_[int(2)] = p2
            nw246_[int(3)] = p0
            nw246_[int(4)] = p1
            nw246_[int(5)] = default__.safeModuloInt(p1, (0) - (default__.fm0(p0, (self).fm24(p0, self.f19, d_1515_v33_, globalState), globalState)))
            nw246_[int(6)] = p0
            nw246_[int(7)] = len((_dafny.Set({self.f19})) | (d_1516_v34_))
            nw246_[int(8)] = (p0) * (p0)
            nw246_[int(9)] = p0
            nw246_[int(10)] = p0
            nw246_[int(11)] = ((_dafny.MultiSet(d_1477_v8_)).intersection(d_1517_v35_)).cardinality
            nw246_[int(12)] = len((self).f9)
            nw246_[int(13)] = p1
            nw246_[int(14)] = p1
            nw246_[int(15)] = default__.safeModuloInt((d_1518_v36_).cf3, len((self).f9))
            nw246_[int(16)] = ((0) - (-106) if self.f19 else default__.fm0(len(d_1516_v34_), p1, globalState))
            nw246_[int(17)] = p1
            nw246_[int(18)] = p2
            nw246_[int(19)] = len((self).f9)
            r2 = nw246_
            index220_ = default__.safeIndex(255, (d_1512_v30_).length(0))
            (d_1512_v30_)[index220_] = p0
            index221_ = default__.safeIndex(255, (d_1512_v30_).length(0))
            (d_1512_v30_)[index221_] = p2
            d_1519_v37_: _dafny.Seq
            d_1519_v37_ = _dafny.SeqWithoutIsStrInference([self.f20])
            source27_ = default__.fm28((p2) - (p2), (d_1519_v37_) + (d_1519_v37_), globalState)
            d_1520___mcc_h5_ = source27_
            d_1521_cf0_ = d_1520___mcc_h5_
            index222_ = default__.safeIndex(63, (d_1478_v9_).length(0))
            (d_1478_v9_)[index222_] = self.f20
            index223_ = default__.safeIndex(63, (d_1478_v9_).length(0))
            (d_1478_v9_)[index223_] = self.f19
            d_1522_v38_: _dafny.Map
            d_1522_v38_ = _dafny.Map({self.f20: (d_1478_v9_)[default__.safeIndex(63, (d_1478_v9_).length(0))]})
            d_1522_v38_ = ((_dafny.Map({self.f19: self.f20})).set(self.f19, not((d_1519_v37_)[default__.safeIndex(-119, len(d_1519_v37_))]))) | (d_1522_v38_)
            d_1523_v39_: _dafny.MultiSet
            d_1523_v39_ = _dafny.MultiSet([not(not((d_1478_v9_)[default__.safeIndex(63, (d_1478_v9_).length(0))]))])
            (self).f19 = ((_dafny.MultiSet([self.f19, (d_1478_v9_)[default__.safeIndex(63, (d_1478_v9_).length(0))], (d_1478_v9_)[default__.safeIndex(63, (d_1478_v9_).length(0))], self.f20, self.f19])) - (d_1523_v39_)).issubset(d_1523_v39_)
            d_1524_v40_: _dafny.Array
            nw247_ = _dafny.Array(None, 10)
            d_1524_v40_ = nw247_
            d_1525_v42_: _dafny.Map
            d_1525_v42_ = _dafny.Map({(d_1512_v30_)[default__.safeIndex(255, (d_1512_v30_).length(0))]: p0})
            d_1526_v43_: _dafny.Map
            d_1526_v43_ = _dafny.Map({d_1525_v42_: self.f20})
            d_1527_v45_: T1
            nw248_ = C3()
            def iife132_():
                coll66_ = _dafny.Map()
                compr_66_: _dafny.Map
                for compr_66_ in (d_1526_v43_).keys.Elements:
                    d_1528_v41_: _dafny.Map = compr_66_
                    if (d_1528_v41_) in (d_1526_v43_):
                        def iife133_():
                            coll67_ = _dafny.Map()
                            compr_67_: int
                            for compr_67_ in (d_1477_v8_).Elements:
                                d_1529_v44_: int = compr_67_
                                if (d_1529_v44_) in (d_1477_v8_):
                                    coll67_[default__.safeDivisionInt(d_1529_v44_, len(d_1519_v37_))] = p2
                            return _dafny.Map(coll67_)
                        coll66_[d_1528_v41_] = iife133_()

                return _dafny.Map(coll66_)
            nw248_.ctor__(self.f20, p2, p2, iife132_()
            , (self).f9, (self).f10)
            d_1527_v45_ = nw248_
            index224_ = default__.safeIndex(511, (d_1524_v40_).length(0))
            (d_1524_v40_)[index224_] = d_1527_v45_
            index225_ = default__.safeIndex(511, (d_1524_v40_).length(0))
            (d_1524_v40_)[index225_] = d_1527_v45_
        d_1530_v46_: _dafny.Seq
        d_1530_v46_ = _dafny.SeqWithoutIsStrInference([p0, p1])
        d_1531_v47_: _dafny.Set
        d_1531_v47_ = _dafny.Set({d_1530_v46_})
        d_1532_v48_: _dafny.MultiSet
        d_1532_v48_ = _dafny.MultiSet([self.f19, (self).fm25(self.f19, self.f19, d_1531_v47_, globalState)])
        r0 = not (self.f19) or ((d_1532_v48_).isdisjoint(d_1532_v48_))
        d_1533_v49_: _dafny.Set
        d_1533_v49_ = _dafny.Set({p0, p2})
        d_1533_v49_ = (d_1533_v49_).intersection((d_1533_v49_).intersection(d_1533_v49_))
        d_1534_v50_: str
        d_1534_v50_ = _dafny.CodePoint('e')
        d_1535_v51_: _dafny.Map
        d_1535_v51_ = _dafny.Map({self.f20: d_1534_v50_})
        d_1536_v52_: _dafny.Seq
        d_1536_v52_ = _dafny.SeqWithoutIsStrInference([d_1535_v51_])
        d_1536_v52_ = ((d_1536_v52_) + (_dafny.SeqWithoutIsStrInference([d_1535_v51_, d_1535_v51_, d_1535_v51_]))) + (default__.fm52(self.f20, self.f19, (0) - (p1), False, globalState))
        d_1537_v53_: _dafny.Array
        nw249_ = _dafny.Array(None, 15)
        nw249_[int(0)] = d_1534_v50_
        nw249_[int(1)] = d_1534_v50_
        nw249_[int(2)] = d_1534_v50_
        nw249_[int(3)] = d_1534_v50_
        nw249_[int(4)] = d_1534_v50_
        nw249_[int(5)] = _dafny.CodePoint('k')
        nw249_[int(6)] = d_1534_v50_
        nw249_[int(7)] = d_1534_v50_
        nw249_[int(8)] = d_1534_v50_
        nw249_[int(9)] = d_1534_v50_
        nw249_[int(10)] = ((self).f9)[default__.safeIndex(p2, len((self).f9))]
        nw249_[int(11)] = _dafny.CodePoint('s')
        nw249_[int(12)] = d_1534_v50_
        nw249_[int(13)] = d_1534_v50_
        nw249_[int(14)] = d_1534_v50_
        d_1537_v53_ = nw249_
        d_1538_v54_: D6
        d_1538_v54_ = D6_DC13(d_1537_v53_)
        d_1539_v55_: _dafny.Map
        d_1539_v55_ = _dafny.Map({p2: d_1538_v54_})
        d_1539_v55_ = (d_1539_v55_).set(p0, d_1538_v54_)
        d_1540_v56_: _dafny.Array
        nw250_ = _dafny.Array(False, 24)
        d_1540_v56_ = nw250_
        nw251_ = _dafny.Array(False, 14)
        d_1540_v56_ = nw251_
        r0 = self.f20
        r1 = self.f20
        d_1541_v57_: _dafny.Array
        def lambda58_(d_1542_p2_):
            def lambda59_(d_1543_i2_):
                return (d_1543_i2_) * ((0) - (d_1542_p2_))

            return lambda59_

        init32_ = lambda58_(p2)
        nw252_ = _dafny.Array(None, 19)
        for i0_32_ in range(nw252_.length(0)):
            nw252_[i0_32_] = init32_(i0_32_)
        d_1541_v57_ = nw252_
        r2 = d_1541_v57_
        r3 = default__.safeModuloInt((p0) + (p1), p0)
        return r0, r1, r2, r3

    def m13(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_1544_v0_: _dafny.Map
        d_1544_v0_ = _dafny.Map({True: (0) - (p2)})
        d_1545_v1_: _dafny.Map
        d_1545_v1_ = _dafny.Map({self.f19: d_1544_v0_})
        d_1546_v2_: _dafny.Seq
        d_1546_v2_ = _dafny.SeqWithoutIsStrInference([p2])
        d_1547_v3_: _dafny.Map
        d_1547_v3_ = _dafny.Map({_dafny.Map({p1: self.f20}): p2})
        d_1548_v4_: _dafny.Set
        d_1548_v4_ = _dafny.Set({d_1546_v2_, _dafny.SeqWithoutIsStrInference([len(d_1547_v3_)])})
        (self).f19 = (len(((d_1545_v1_)[(self).fm25(not(self.f20), self.f20, d_1548_v4_, globalState)] if ((self).fm25(not(self.f20), self.f20, d_1548_v4_, globalState)) in (d_1545_v1_) else d_1544_v0_))) <= (default__.fm0(p2, default__.fm0(p2, p2, globalState), globalState))
        d_1549_v5_: C0
        nw253_ = C0()
        nw253_.ctor__((p2) != (p2), self.f19, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")))), p3), (self).f10)
        d_1549_v5_ = nw253_
        d_1550_v6_: _dafny.Map
        d_1550_v6_ = _dafny.Map({self.f20: self.f19})
        d_1550_v6_ = (d_1550_v6_).set(self.f19, p1)
        d_1551_v7_: _dafny.Seq
        d_1551_v7_ = _dafny.SeqWithoutIsStrInference([self.f20, (d_1549_v5_).f26])
        d_1552_v8_: _dafny.Map
        d_1552_v8_ = _dafny.Map({len(d_1551_v7_): (self).f9})
        d_1553_v9_: _dafny.MultiSet
        d_1553_v9_ = _dafny.MultiSet([p2])
        rhs213_ = 902
        rhs214_ = (d_1549_v5_).fm5(p2, len((d_1552_v8_).set(p2, (self).f9)), globalState)
        rhs215_ = ((D7_DC15(d_1553_v9_)).cf18).cardinality
        rhs216_ = (not(((d_1550_v6_)[self.f20] if (self.f20) in (d_1550_v6_) else self.f19))) not in (default__.fm51((d_1549_v5_).f25, p2, p3, (d_1549_v5_).f25, globalState))
        lhs140_ = globalState
        lhs141_ = globalState
        lhs142_ = self
        r0 = rhs213_
        lhs140_.f3 = rhs214_
        lhs141_.f3 = rhs215_
        lhs142_.f19 = rhs216_
        d_1554_v10_: _dafny.Array
        nw254_ = _dafny.Array(D8.default()(), 22)
        d_1554_v10_ = nw254_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1554_v10_).length(0)):
            d_1555_i0_: int = guard_loop_5_
            if (True) and (((0) <= (d_1555_i0_)) and ((d_1555_i0_) < ((d_1554_v10_).length(0)))):
                (d_1554_v10_)[(d_1555_i0_)] = D8_DC18()
        d_1556_v11_: _dafny.Set
        d_1556_v11_ = _dafny.Set({p3, p3})
        d_1557_v12_: _dafny.Map
        d_1557_v12_ = _dafny.Map({p2: d_1553_v9_})
        d_1558_v13_: _dafny.Array
        nw255_ = _dafny.Array(None, 22)
        nw255_[int(0)] = d_1553_v9_
        nw255_[int(1)] = d_1553_v9_
        nw255_[int(2)] = (_dafny.MultiSet([p2])).intersection(((d_1557_v12_)[p2] if (p2) in (d_1557_v12_) else d_1553_v9_))
        nw255_[int(3)] = d_1553_v9_
        nw255_[int(4)] = default__.fm53(globalState)
        nw255_[int(5)] = d_1553_v9_
        nw255_[int(6)] = d_1553_v9_
        nw255_[int(7)] = d_1553_v9_
        nw255_[int(8)] = (d_1553_v9_).intersection(d_1553_v9_)
        nw255_[int(9)] = d_1553_v9_
        nw255_[int(10)] = d_1553_v9_
        nw255_[int(11)] = d_1553_v9_
        nw255_[int(12)] = d_1553_v9_
        nw255_[int(13)] = d_1553_v9_
        nw255_[int(14)] = _dafny.MultiSet([p2, 593, p2])
        nw255_[int(15)] = (d_1553_v9_) - (d_1553_v9_)
        nw255_[int(16)] = d_1553_v9_
        nw255_[int(17)] = d_1553_v9_
        nw255_[int(18)] = d_1553_v9_
        nw255_[int(19)] = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([p2 for d_1559_i1_ in range(default__.abs(868))])) + (_dafny.SeqWithoutIsStrInference([p2, p2, p2])))
        nw255_[int(20)] = d_1553_v9_
        nw255_[int(21)] = d_1553_v9_
        d_1558_v13_ = nw255_
        index226_ = default__.safeIndex(675, (d_1558_v13_).length(0))
        (d_1558_v13_)[index226_] = _dafny.MultiSet(d_1546_v2_)
        d_1560_v14_: _dafny.Array
        def lambda60_(d_1561_i2_):
            return (d_1561_i2_) + (len((self).f9))

        init33_ = lambda60_
        nw256_ = _dafny.Array(None, 23)
        for i0_33_ in range(nw256_.length(0)):
            nw256_[i0_33_] = init33_(i0_33_)
        d_1560_v14_ = nw256_
        d_1562_v15_: _dafny.Map
        d_1562_v15_ = _dafny.Map({p2: len(d_1551_v7_)})
        d_1563_v16_: _dafny.Map
        d_1563_v16_ = _dafny.Map({d_1562_v15_: _dafny.Map({(d_1553_v9_).cardinality: p2})})
        d_1564_v17_: C2
        nw257_ = C2()
        nw257_.ctor__(p2, d_1563_v16_, (self).f9, (self).f10)
        d_1564_v17_ = nw257_
        index227_ = default__.safeIndex(675, (d_1558_v13_).length(0))
        rhs217_ = (d_1556_v11_ if (d_1549_v5_).f25 else _dafny.Set({p3}))
        rhs218_ = ((_dafny.SeqWithoutIsStrInference([(d_1549_v5_).f26, (d_1549_v5_).f25])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([(d_1549_v5_).f26, (d_1549_v5_).f25]))), p1)) == (d_1551_v7_)
        rhs219_ = _dafny.MultiSet([default__.fm0(802, p2, globalState), (0) - (p2)])
        rhs220_ = d_1560_v14_
        rhs221_ = d_1564_v17_
        lhs143_ = self
        lhs144_ = d_1558_v13_
        lhs145_ = default__.safeIndex(675, (d_1558_v13_).length(0))
        d_1556_v11_ = rhs217_
        lhs143_.f19 = rhs218_
        lhs144_[lhs145_] = rhs219_
        d_1560_v14_ = rhs220_
        d_1564_v17_ = rhs221_
        r0 = (0) - (p2)
        return r0

    def m14(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        d_1565_v0_: int
        d_1565_v0_ = 371
        d_1566_v1_: _dafny.Set
        d_1566_v1_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_1565_v0_ for d_1567_i1_ in range(default__.abs(261))])})
        hi3_ = (d_1565_v0_ if (self).fm25(self.f19, self.f20, d_1566_v1_, globalState) else -159)
        for d_1568_i0_ in range((len((self).f9)) * (d_1565_v0_), hi3_):
            d_1569_v2_: _dafny.Seq
            d_1569_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yj"))
            d_1569_v2_ = d_1569_v2_
            d_1570_v3_: D14
            d_1570_v3_ = D14_DC33(d_1565_v0_, (self).f9, False)
            def iife134_(_pat_let33_0):
                def iife135_(d_1571_dt__update__tmp_h0_):
                    def iife136_(_pat_let34_0):
                        def iife137_(d_1572_dt__update_hcf40_h0_):
                            def iife138_(_pat_let35_0):
                                def iife139_(d_1573_dt__update_hcf42_h0_):
                                    return D14_DC33(d_1572_dt__update_hcf40_h0_, (d_1571_dt__update__tmp_h0_).cf41, d_1573_dt__update_hcf42_h0_)
                                return iife139_(_pat_let35_0)
                            return iife138_(self.f19)
                        return iife137_(_pat_let34_0)
                    return iife136_(d_1568_i0_)
                return iife135_(_pat_let33_0)
            d_1569_v2_ = ((iife134_(d_1570_v3_)).cf41) + ((self).f9)
            d_1574_v4_: D4
            d_1574_v4_ = D4_DC8(default__.fm46(globalState))
            source28_ = d_1574_v4_
            if source28_.is_DC9:
                d_1575___mcc_h0_ = source28_.cf11
                d_1576_cf11_ = d_1575___mcc_h0_
                d_1577_v5_: _dafny.Map
                d_1577_v5_ = _dafny.Map({d_1565_v0_: self.f19})
                d_1578_v6_: _dafny.Seq
                d_1578_v6_ = _dafny.SeqWithoutIsStrInference([default__.fm46(globalState), d_1577_v5_, (d_1577_v5_) | (d_1577_v5_), d_1577_v5_])
                rhs222_ = ((self).f9) + (d_1569_v2_)
                rhs223_ = not (self.f19) or (not(self.f20))
                rhs224_ = d_1578_v6_
                rhs225_ = self.f19
                d_1569_v2_ = rhs222_
                r0 = rhs223_
                d_1578_v6_ = rhs224_
                r1 = rhs225_
                (globalState).f3 = d_1565_v0_
                (self).f19 = not((default__.safeDivisionInt(-213, d_1568_i0_)) > ((920 if self.f20 else 313)))
                d_1579_v7_: _dafny.Seq
                d_1579_v7_ = _dafny.SeqWithoutIsStrInference([d_1568_i0_])
                d_1580_v8_: _dafny.Array
                nw258_ = _dafny.Array(int(0), 27)
                d_1580_v8_ = nw258_
                d_1581_v9_: D6
                d_1581_v9_ = D6_DC14(d_1580_v8_)
                d_1582_v10_: _dafny.MultiSet
                d_1582_v10_ = _dafny.MultiSet([d_1581_v9_])
                d_1583_v11_: _dafny.Map
                d_1583_v11_ = _dafny.Map({d_1568_i0_: (d_1582_v10_).cardinality})
                d_1584_v12_: _dafny.Map
                d_1584_v12_ = _dafny.Map({d_1568_i0_: d_1580_v8_})
                d_1585_v13_: _dafny.Map
                d_1585_v13_ = _dafny.Map({(d_1583_v11_).set(d_1565_v0_, d_1568_i0_): _dafny.Map({len(d_1579_v7_): d_1568_i0_})})
                d_1586_v14_: C3
                nw259_ = C3()
                nw259_.ctor__((d_1579_v7_) < (_dafny.SeqWithoutIsStrInference([d_1568_i0_ for d_1587_i2_ in range(default__.abs(166))])), (len(d_1583_v11_)) * (d_1568_i0_), len((d_1584_v12_) | (d_1584_v12_)), d_1585_v13_, d_1569_v2_, ((self).f10) + ((self).f10))
                d_1586_v14_ = nw259_
            elif source28_.is_DC10:
                d_1588___mcc_h1_ = source28_.cf12
                d_1589_cf12_ = d_1588___mcc_h1_
                (self).f20 = self.f20
                (globalState).f3 = default__.safeDivisionInt(d_1565_v0_, default__.safeDivisionInt((0) - (d_1565_v0_), (0) - (d_1568_i0_)))
                d_1590_v15_: _dafny.Array
                def lambda61_(d_1591_i3_):
                    return (_dafny.CodePoint('a') if self.f20 else _dafny.CodePoint('s'))

                init34_ = lambda61_
                nw260_ = _dafny.Array(None, 15)
                for i0_34_ in range(nw260_.length(0)):
                    nw260_[i0_34_] = init34_(i0_34_)
                d_1590_v15_ = nw260_
                d_1590_v15_ = (d_1590_v15_ if not (self.f19) or (self.f20) else d_1590_v15_)
                d_1592_v16_: str
                d_1592_v16_ = _dafny.CodePoint('s')
                d_1592_v16_ = d_1592_v16_
            elif source28_.is_DC11:
                d_1593___mcc_h2_ = source28_.cf13
                d_1594___mcc_h3_ = source28_.cf14
                d_1595_cf14_ = d_1594___mcc_h3_
                d_1596_cf13_ = d_1593___mcc_h2_
                d_1597_v17_: _dafny.MultiSet
                d_1597_v17_ = _dafny.MultiSet([d_1565_v0_, 417, d_1595_cf14_, len((self).f9)])
                d_1598_v18_: _dafny.Map
                d_1598_v18_ = _dafny.Map({d_1597_v17_: d_1568_i0_})
                d_1599_v19_: _dafny.Map
                d_1599_v19_ = _dafny.Map({d_1598_v18_: self.f20})
                d_1599_v19_ = (d_1599_v19_).set((d_1598_v18_) | (d_1598_v18_), False)
                d_1600_v20_: _dafny.Array
                def lambda62_(d_1601_cf13_):
                    def lambda63_(d_1602_i4_):
                        return (_dafny.SeqWithoutIsStrInference([self.f20])) == (_dafny.SeqWithoutIsStrInference([self.f19, d_1601_cf13_, self.f19, self.f19]))

                    return lambda63_

                init35_ = lambda62_(d_1596_cf13_)
                nw261_ = _dafny.Array(None, 19)
                for i0_35_ in range(nw261_.length(0)):
                    nw261_[i0_35_] = init35_(i0_35_)
                d_1600_v20_ = nw261_
                index228_ = default__.safeIndex(85, (d_1600_v20_).length(0))
                (d_1600_v20_)[index228_] = d_1596_cf13_
                d_1603_v21_: _dafny.Set
                d_1603_v21_ = _dafny.Set({self.f19, d_1596_cf13_})
                index229_ = default__.safeIndex(85, (d_1600_v20_).length(0))
                (d_1600_v20_)[index229_] = (d_1603_v21_).isdisjoint((d_1603_v21_) | (_dafny.Set({True, self.f19, self.f19, self.f19})))
                d_1604_v22_: _dafny.Array
                nw262_ = _dafny.Array(_dafny.Seq({}), 17)
                d_1604_v22_ = nw262_
                d_1604_v22_ = d_1604_v22_
                d_1605_v23_: _dafny.Array
                nw263_ = _dafny.Array(int(0), 4)
                d_1605_v23_ = nw263_
                d_1606_v24_: _dafny.Map
                d_1606_v24_ = _dafny.Map({d_1605_v23_: self.f19})
                (globalState).f3 = len((d_1606_v24_) | (d_1606_v24_))
            elif True:
                d_1607___mcc_h4_ = source28_.cf10
                d_1608_cf10_ = d_1607___mcc_h4_
                d_1569_v2_ = (d_1569_v2_) + ((d_1569_v2_) + ((self).f9))
                (globalState).f3 = len((self).f9)
                (globalState).f3 = (d_1568_i0_) - (d_1568_i0_)
                d_1609_v25_: D4
                d_1609_v25_ = D4_DC11(self.f20, d_1568_i0_)
                d_1610_v26_: _dafny.Seq
                d_1610_v26_ = _dafny.SeqWithoutIsStrInference([124])
                d_1611_v27_: _dafny.Map
                d_1611_v27_ = _dafny.Map({not(self.f19): d_1568_i0_})
                pat_let_tv28_ = d_1610_v26_
                d_1612_v28_: _dafny.Array
                nw264_ = _dafny.Array(None, 22)
                nw264_[int(0)] = default__.fm54(d_1568_i0_, d_1565_v0_, globalState)
                nw264_[int(1)] = d_1609_v25_
                def iife140_(_pat_let36_0):
                    def iife141_(d_1613_dt__update__tmp_h1_):
                        def iife142_(_pat_let37_0):
                            def iife143_(d_1614_dt__update_hcf14_h0_):
                                return D4_DC11((d_1613_dt__update__tmp_h1_).cf13, d_1614_dt__update_hcf14_h0_)
                            return iife143_(_pat_let37_0)
                        return iife142_(len(pat_let_tv28_))
                    return iife141_(_pat_let36_0)
                nw264_[int(2)] = iife140_(d_1609_v25_)
                nw264_[int(3)] = d_1609_v25_
                nw264_[int(4)] = d_1609_v25_
                nw264_[int(5)] = D4_DC11(self.f20, d_1568_i0_)
                nw264_[int(6)] = d_1609_v25_
                nw264_[int(7)] = d_1609_v25_
                nw264_[int(8)] = D4_DC11(self.f19, len(d_1611_v27_))
                nw264_[int(9)] = d_1609_v25_
                def iife144_(_pat_let38_0):
                    def iife145_(d_1615_dt__update__tmp_h2_):
                        def iife146_(_pat_let39_0):
                            def iife147_(d_1616_dt__update_hcf13_h0_):
                                return D4_DC11(d_1616_dt__update_hcf13_h0_, (d_1615_dt__update__tmp_h2_).cf14)
                            return iife147_(_pat_let39_0)
                        return iife146_(self.f20)
                    return iife145_(_pat_let38_0)
                nw264_[int(10)] = iife144_(d_1609_v25_)
                nw264_[int(11)] = d_1609_v25_
                nw264_[int(12)] = D4_DC11(self.f20, d_1565_v0_)
                nw264_[int(13)] = d_1609_v25_
                nw264_[int(14)] = D4_DC11(self.f19, d_1568_i0_)
                nw264_[int(15)] = d_1609_v25_
                nw264_[int(16)] = d_1609_v25_
                nw264_[int(17)] = d_1609_v25_
                nw264_[int(18)] = d_1609_v25_
                nw264_[int(19)] = d_1609_v25_
                nw264_[int(20)] = d_1609_v25_
                nw264_[int(21)] = d_1609_v25_
                d_1612_v28_ = nw264_
                d_1617_v29_: _dafny.Map
                d_1617_v29_ = _dafny.Map({d_1612_v28_: d_1569_v2_})
                d_1618_v30_: _dafny.Seq
                d_1618_v30_ = _dafny.SeqWithoutIsStrInference([d_1569_v2_, d_1569_v2_])
                d_1619_v31_: _dafny.Array
                nw265_ = _dafny.Array(None, 10)
                nw265_[int(0)] = ((self).f9) + ((self).f9)
                nw265_[int(1)] = d_1569_v2_
                nw265_[int(2)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1620_i5_ in range(default__.abs(736))])) + (((d_1617_v29_)[d_1612_v28_] if (d_1612_v28_) in (d_1617_v29_) else d_1569_v2_))
                nw265_[int(3)] = d_1569_v2_
                nw265_[int(4)] = d_1569_v2_
                nw265_[int(5)] = ((d_1618_v30_)[default__.safeIndex(555, len(d_1618_v30_))]) + ((self).f9)
                nw265_[int(6)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1621_i6_ in range(default__.abs(587))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))
                nw265_[int(7)] = d_1569_v2_
                nw265_[int(8)] = (self).f9
                nw265_[int(9)] = d_1569_v2_
                d_1619_v31_ = nw265_
                d_1622_v32_: _dafny.Array
                def lambda64_(d_1623_i7_):
                    return self.f20

                init36_ = lambda64_
                nw266_ = _dafny.Array(None, 25)
                for i0_36_ in range(nw266_.length(0)):
                    nw266_[i0_36_] = init36_(i0_36_)
                d_1622_v32_ = nw266_
                rhs226_ = d_1619_v31_
                rhs227_ = d_1622_v32_
                d_1619_v31_ = rhs226_
                d_1622_v32_ = rhs227_
            d_1624_v33_: _dafny.Seq
            d_1624_v33_ = _dafny.SeqWithoutIsStrInference([(self).f9])
            d_1625_v34_: _dafny.Map
            d_1625_v34_ = _dafny.Map({d_1565_v0_: (d_1624_v33_)[default__.safeIndex(d_1568_i0_, len(d_1624_v33_))]})
            d_1626_v35_: _dafny.Map
            d_1626_v35_ = _dafny.Map({d_1569_v2_: self.f19})
            d_1627_v36_: _dafny.Set
            d_1627_v36_ = _dafny.Set({d_1565_v0_, d_1565_v0_, len(d_1626_v35_), -636})
            d_1628_v37_: _dafny.Seq
            d_1628_v37_ = _dafny.SeqWithoutIsStrInference([d_1627_v36_])
            d_1629_v38_: _dafny.Set
            d_1629_v38_ = _dafny.Set({d_1569_v2_})
            d_1630_v39_: str
            d_1630_v39_ = _dafny.CodePoint('i')
            d_1631_v40_: _dafny.Map
            d_1631_v40_ = _dafny.Map({d_1565_v0_: d_1630_v39_})
            d_1625_v34_ = (d_1625_v34_).set((0) - ((len(d_1628_v37_)) - (len(d_1629_v38_))), ((((self).f9).set(default__.safeIndex(d_1565_v0_, len((self).f9)), d_1630_v39_)).set(default__.safeIndex(d_1565_v0_, len(((self).f9).set(default__.safeIndex(d_1565_v0_, len((self).f9)), d_1630_v39_))), ((d_1631_v40_)[d_1565_v0_] if (d_1565_v0_) in (d_1631_v40_) else d_1630_v39_))) + ((self).f9))
        d_1632_i8_: int
        d_1632_i8_ = 0
        with _dafny.label("8"):
            while False:
                with _dafny.c_label("8"):
                    if (d_1632_i8_) >= (100):
                        raise _dafny.Break("8")
                    d_1632_i8_ = (d_1632_i8_) + (1)
                    if not(self.f20):
                        d_1633_v41_: _dafny.Seq
                        d_1633_v41_ = _dafny.SeqWithoutIsStrInference([self.f20, self.f19, True, self.f19])
                        d_1634_v42_: _dafny.Map
                        d_1634_v42_ = _dafny.Map({d_1565_v0_: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1635_i9_ in range(default__.abs(349))]))})
                        d_1636_v43_: _dafny.Array
                        nw267_ = _dafny.Array(None, 18)
                        nw267_[int(0)] = self.f20
                        nw267_[int(1)] = self.f20
                        nw267_[int(2)] = self.f19
                        nw267_[int(3)] = False
                        nw267_[int(4)] = self.f19
                        nw267_[int(5)] = self.f20
                        nw267_[int(6)] = self.f19
                        nw267_[int(7)] = self.f20
                        nw267_[int(8)] = (self).fm25(False, self.f20, d_1566_v1_, globalState)
                        nw267_[int(9)] = self.f20
                        nw267_[int(10)] = not(self.f20)
                        nw267_[int(11)] = self.f19
                        nw267_[int(12)] = self.f19
                        nw267_[int(13)] = self.f20
                        nw267_[int(14)] = self.f19
                        nw267_[int(15)] = self.f20
                        nw267_[int(16)] = self.f20
                        nw267_[int(17)] = (d_1633_v41_)[default__.safeIndex(len(d_1634_v42_), len(d_1633_v41_))]
                        d_1636_v43_ = nw267_
                        d_1637_v44_: _dafny.Map
                        d_1637_v44_ = _dafny.Map({d_1565_v0_: d_1636_v43_})
                        index230_ = default__.safeIndex(530, (d_1636_v43_).length(0))
                        (d_1636_v43_)[index230_] = (True) or (self.f19)
                        index231_ = default__.safeIndex(530, (d_1636_v43_).length(0))
                        rhs228_ = d_1637_v44_
                        rhs229_ = True
                        rhs230_ = 299
                        lhs146_ = d_1636_v43_
                        lhs147_ = default__.safeIndex(530, (d_1636_v43_).length(0))
                        lhs148_ = globalState
                        d_1637_v44_ = rhs228_
                        lhs146_[lhs147_] = rhs229_
                        lhs148_.f3 = rhs230_
                        d_1638_v45_: _dafny.Set
                        d_1638_v45_ = _dafny.Set({d_1565_v0_, (len((self).f9)) * (d_1565_v0_), 450})
                        d_1565_v0_ = (0) - (len(d_1638_v45_))
                        d_1639_v46_: _dafny.Seq
                        d_1639_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
                        d_1640_v47_: _dafny.MultiSet
                        d_1640_v47_ = _dafny.MultiSet([-696])
                        d_1641_v48_: _dafny.Map
                        d_1641_v48_ = _dafny.Map({self.f19: d_1640_v47_})
                        d_1642_v49_: _dafny.Map
                        d_1642_v49_ = _dafny.Map({d_1565_v0_: (self).f9})
                        d_1639_v46_ = default__.fm26((0) - (len(d_1641_v48_)), d_1642_v49_, globalState)
                        d_1643_v50_: _dafny.Array
                        nw268_ = _dafny.Array(_dafny.Map({}), 14)
                        d_1643_v50_ = nw268_
                        d_1644_v51_: _dafny.Map
                        d_1644_v51_ = _dafny.Map({d_1565_v0_: (d_1633_v41_)[default__.safeIndex(d_1565_v0_, len(d_1633_v41_))]})
                        d_1645_v52_: D7
                        d_1645_v52_ = D7_DC15(d_1640_v47_)
                        d_1646_v53_: str
                        d_1646_v53_ = default__.fm39(d_1645_v52_, (d_1636_v43_)[default__.safeIndex(530, (d_1636_v43_).length(0))], globalState)
                        d_1647_v54_: str
                        d_1647_v54_ = _dafny.CodePoint('i')
                        d_1648_v55_: int
                        out53_: int
                        out53_ = (self).m13(d_1643_v50_, (((d_1644_v51_)[d_1565_v0_] if (d_1565_v0_) in (d_1644_v51_) else self.f19) if self.f20 else (d_1636_v43_)[default__.safeIndex(530, (d_1636_v43_).length(0))]), d_1565_v0_, (d_1647_v54_), globalState)
                        d_1648_v55_ = out53_
                        d_1649_v57_: _dafny.Seq
                        def iife148_():
                            coll68_ = _dafny.Set()
                            compr_68_: int
                            for compr_68_ in _dafny.IntegerRange(415, 236):
                                d_1650_v56_: int = compr_68_
                                if ((415) <= (d_1650_v56_)) and ((d_1650_v56_) < (236)):
                                    coll68_ = coll68_.union(_dafny.Set([(d_1650_v56_) - (d_1565_v0_)]))
                            return _dafny.Set(coll68_)
                        d_1649_v57_ = _dafny.SeqWithoutIsStrInference([(len(d_1633_v41_)) + (len(iife148_()
                        ))])
                        d_1649_v57_ = ((d_1649_v57_ if self.f20 else (d_1649_v57_) + (d_1649_v57_))).set(default__.safeIndex(d_1648_v55_, len((d_1649_v57_ if self.f20 else (d_1649_v57_) + (d_1649_v57_)))), default__.safeModuloInt(d_1648_v55_, d_1648_v55_))
                    elif True:
                        d_1651_v58_: _dafny.Array
                        nw269_ = _dafny.Array(False, 12)
                        d_1651_v58_ = nw269_
                        index232_ = default__.safeIndex(76, (p0).length(0))
                        (p0)[index232_] = d_1651_v58_
                        index233_ = default__.safeIndex(76, (p0).length(0))
                        (p0)[index233_] = d_1651_v58_
                        d_1652_v59_: _dafny.Map
                        d_1652_v59_ = _dafny.Map({d_1565_v0_: True})
                        d_1653_v60_: _dafny.MultiSet
                        d_1653_v60_ = _dafny.MultiSet([d_1565_v0_, d_1565_v0_, d_1565_v0_, d_1565_v0_, len(d_1652_v59_)])
                        (globalState).f3 = default__.safeDivisionInt(d_1565_v0_, default__.safeModuloInt(len((self).f9), (d_1653_v60_).cardinality))
                        d_1654_v61_: _dafny.Seq
                        d_1654_v61_ = _dafny.SeqWithoutIsStrInference([d_1565_v0_])
                        d_1655_v62_: _dafny.Map
                        d_1655_v62_ = _dafny.Map({d_1565_v0_: (_dafny.MultiSet([self.f20])).cardinality})
                        d_1656_v63_: _dafny.Map
                        d_1656_v63_ = _dafny.Map({d_1655_v62_: d_1655_v62_})
                        d_1657_v64_: _dafny.Seq
                        d_1657_v64_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1654_v61_, _dafny.SeqWithoutIsStrInference([d_1565_v0_ for d_1658_i10_ in range(default__.abs(120))])])])
                        d_1659_v65_: _dafny.Map
                        d_1659_v65_ = _dafny.Map({self.f19: d_1652_v59_})
                        d_1660_v66_: C3
                        nw270_ = C3()
                        nw270_.ctor__((False) and (self.f19), d_1565_v0_, default__.safeDivisionInt(d_1565_v0_, (d_1654_v61_)[default__.safeIndex(d_1565_v0_, len(d_1654_v61_))]), d_1656_v63_, (self).f9, ((d_1657_v64_)[default__.safeIndex(len(d_1659_v65_), len(d_1657_v64_))]).set(default__.safeIndex(len((self).f9), len((d_1657_v64_)[default__.safeIndex(len(d_1659_v65_), len(d_1657_v64_))])), d_1654_v61_))
                        d_1660_v66_ = nw270_
                        d_1661_v67_: str
                        d_1661_v67_ = _dafny.CodePoint('l')
                        d_1662_v68_: _dafny.Map
                        d_1662_v68_ = _dafny.Map({d_1661_v67_: self.f19})
                        arr0_ = (p0)[default__.safeIndex(76, (p0).length(0))]
                        index234_ = default__.safeIndex(407, ((p0)[default__.safeIndex(76, (p0).length(0))]).length(0))
                        arr0_[index234_] = (d_1660_v66_).f21
                        d_1663_v69_: _dafny.MultiSet
                        d_1663_v69_ = _dafny.MultiSet([d_1661_v67_])
                        arr1_ = (p0)[default__.safeIndex(76, (p0).length(0))]
                        index235_ = default__.safeIndex(407, ((p0)[default__.safeIndex(76, (p0).length(0))]).length(0))
                        rhs231_ = _dafny.Map({d_1661_v67_: False})
                        rhs232_ = d_1565_v0_
                        rhs233_ = ((_dafny.MultiSet([d_1661_v67_, d_1661_v67_])) - (d_1663_v69_)).issubset((d_1663_v69_).intersection(d_1663_v69_))
                        lhs149_ = globalState
                        lhs150_ = (p0)[default__.safeIndex(76, (p0).length(0))]
                        lhs151_ = default__.safeIndex(407, ((p0)[default__.safeIndex(76, (p0).length(0))]).length(0))
                        d_1662_v68_ = rhs231_
                        lhs149_.f3 = rhs232_
                        lhs150_[lhs151_] = rhs233_
                        arr2_ = (p0)[default__.safeIndex(76, (p0).length(0))]
                        index236_ = default__.safeIndex(407, ((p0)[default__.safeIndex(76, (p0).length(0))]).length(0))
                        arr2_[index236_] = not((d_1660_v66_).f21)
                    d_1565_v0_ = default__.safeDivisionInt(d_1565_v0_, d_1565_v0_)
                    if self.f19:
                        (self).f19 = self.f20
                        d_1664_v70_: _dafny.Seq
                        d_1664_v70_ = _dafny.SeqWithoutIsStrInference([self.f20, self.f19, self.f19])
                        d_1665_v71_: str
                        d_1665_v71_ = _dafny.CodePoint('u')
                        d_1666_v72_: _dafny.Map
                        d_1666_v72_ = _dafny.Map({d_1665_v71_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrc"))})
                        d_1667_v73_: _dafny.Seq
                        d_1667_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knlhj"))
                        rhs234_ = not(not(not((d_1664_v70_)[default__.safeIndex((len(((d_1666_v72_)[d_1665_v71_] if (d_1665_v71_) in (d_1666_v72_) else (d_1667_v73_))) if False else 271), len(d_1664_v70_))])))
                        rhs235_ = len(d_1566_v1_)
                        rhs236_ = default__.fm0(d_1565_v0_, (d_1565_v0_) - (91), globalState)
                        lhs152_ = globalState
                        lhs153_ = globalState
                        r1 = rhs234_
                        lhs152_.f3 = rhs235_
                        lhs153_.f3 = rhs236_
                        d_1668_v75_: _dafny.Map
                        d_1668_v75_ = _dafny.Map({d_1565_v0_: self.f20})
                        d_1669_v76_: _dafny.Map
                        def iife149_():
                            coll69_ = _dafny.Map()
                            compr_69_: D4
                            for compr_69_ in (_dafny.Map({D4_DC10(d_1668_v75_): len(_dafny.Map({len((self).f9): self.f19}))})).keys.Elements:
                                d_1670_v74_: D4 = compr_69_
                                if (d_1670_v74_) in (_dafny.Map({D4_DC10(d_1668_v75_): len(_dafny.Map({len((self).f9): self.f19}))})):
                                    coll69_[d_1670_v74_] = d_1565_v0_
                            return _dafny.Map(coll69_)
                        d_1669_v76_ = _dafny.Map({d_1565_v0_: len(iife149_()
                        )})
                        d_1671_v77_: _dafny.Map
                        d_1671_v77_ = _dafny.Map({d_1665_v71_: self.f19})
                        d_1672_v78_: _dafny.Seq
                        d_1672_v78_ = _dafny.SeqWithoutIsStrInference([len(d_1669_v76_), len(d_1671_v77_)])
                        d_1673_v79_: _dafny.Array
                        nw271_ = _dafny.Array(None, 11)
                        nw271_[int(0)] = (D12_DC29(self.f19, 111, d_1672_v78_, self.f20, -271)).cf35
                        nw271_[int(1)] = self.f19
                        nw271_[int(2)] = self.f19
                        nw271_[int(3)] = not(self.f20)
                        nw271_[int(4)] = self.f19
                        nw271_[int(5)] = self.f19
                        nw271_[int(6)] = self.f20
                        nw271_[int(7)] = self.f19
                        nw271_[int(8)] = self.f19
                        nw271_[int(9)] = self.f20
                        nw271_[int(10)] = True
                        d_1673_v79_ = nw271_
                        d_1674_v80_: _dafny.Map
                        d_1674_v80_ = _dafny.Map({d_1673_v79_: d_1673_v79_})
                        d_1674_v80_ = (d_1674_v80_).set(d_1673_v79_, d_1673_v79_)
                        d_1675_v81_: _dafny.Array
                        def lambda65_(d_1676_i11_):
                            return (self).f9

                        init37_ = lambda65_
                        nw272_ = _dafny.Array(None, 3)
                        for i0_37_ in range(nw272_.length(0)):
                            nw272_[i0_37_] = init37_(i0_37_)
                        d_1675_v81_ = nw272_
                        index237_ = default__.safeIndex(969, (d_1675_v81_).length(0))
                        (d_1675_v81_)[index237_] = (self).f9
                        index238_ = default__.safeIndex(969, (d_1675_v81_).length(0))
                        (d_1675_v81_)[index238_] = (self).f9
                        d_1677_v82_: _dafny.Array
                        def lambda66_(d_1678_v71_):
                            def lambda67_(d_1679_i12_):
                                return d_1678_v71_

                            return lambda67_

                        init38_ = lambda66_(d_1665_v71_)
                        nw273_ = _dafny.Array(None, 29)
                        for i0_38_ in range(nw273_.length(0)):
                            nw273_[i0_38_] = init38_(i0_38_)
                        d_1677_v82_ = nw273_
                        index239_ = default__.safeIndex(729, (d_1677_v82_).length(0))
                        (d_1677_v82_)[index239_] = ((self).f9)[default__.safeIndex(d_1565_v0_, len((self).f9))]
                        index240_ = default__.safeIndex(729, (d_1677_v82_).length(0))
                        (d_1677_v82_)[index240_] = d_1665_v71_
                    elif True:
                        d_1680_v83_: _dafny.Seq
                        d_1680_v83_ = _dafny.SeqWithoutIsStrInference([d_1565_v0_])
                        (self).f20 = (d_1680_v83_) != (d_1680_v83_)
                        d_1681_v84_: str
                        d_1681_v84_ = _dafny.CodePoint('a')
                        d_1682_v85_: _dafny.Map
                        d_1682_v85_ = _dafny.Map({self.f20: self.f20})
                        d_1683_v86_: _dafny.Seq
                        d_1683_v86_ = _dafny.SeqWithoutIsStrInference([self.f19, False, self.f19, self.f20, False])
                        r0 = ((((d_1682_v85_)[self.f19] if (self.f19) in (d_1682_v85_) else (self).fm25(self.f19, self.f19, _dafny.Set({d_1680_v83_}), globalState))) == ((d_1683_v86_)[default__.safeIndex(d_1565_v0_, len(d_1683_v86_))]) if (default__.fm38((self).fm24((0) - (d_1565_v0_), self.f20, d_1681_v84_, globalState), globalState)) or (True) else self.f20)
                        d_1684_v87_: _dafny.MultiSet
                        d_1684_v87_ = _dafny.MultiSet([self.f19])
                        d_1685_v88_: _dafny.Set
                        d_1685_v88_ = _dafny.Set({d_1684_v87_})
                        r1 = not ((d_1685_v88_).issubset(d_1685_v88_)) or (self.f19)
                        (self).f20 = (self.f19) == (self.f20)
                        (self).f19 = (self.f20) == (self.f19)
                    d_1686_v89_: _dafny.Seq
                    d_1686_v89_ = _dafny.SeqWithoutIsStrInference([d_1565_v0_, d_1565_v0_])
                    d_1687_v90_: _dafny.Array
                    nw274_ = _dafny.Array(False, 26)
                    d_1687_v90_ = nw274_
                    d_1688_v91_: D2
                    d_1688_v91_ = D2_DC5(self.f19, d_1687_v90_, len(_dafny.SeqWithoutIsStrInference([self.f19, self.f19])), 538)
                    d_1689_v92_: D12
                    d_1689_v92_ = D12_DC28((d_1686_v89_).set(default__.safeIndex(d_1565_v0_, len(d_1686_v89_)), d_1565_v0_))
                    d_1690_v93_: D12
                    d_1690_v93_ = D12_DC29(self.f20, len((self).f9), d_1686_v89_, self.f20, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1691_i14_ in range(default__.abs(967))])))
                    d_1692_v94_: _dafny.Map
                    d_1692_v94_ = _dafny.Map({(D17_DC40((self).f9, self.f19, len((self).f9))).cf55: self.f19})
                    d_1693_v95_: _dafny.Set
                    d_1693_v95_ = _dafny.Set({self.f20, self.f20, self.f20})
                    d_1694_v96_: _dafny.Array
                    nw275_ = _dafny.Array(None, 24)
                    nw275_[int(0)] = d_1686_v89_
                    nw275_[int(1)] = d_1686_v89_
                    nw275_[int(2)] = (d_1686_v89_) + (d_1686_v89_)
                    nw275_[int(3)] = d_1686_v89_
                    nw275_[int(4)] = d_1686_v89_
                    nw275_[int(5)] = d_1686_v89_
                    nw275_[int(6)] = default__.fm44(self.f19, d_1565_v0_, (d_1688_v91_).cf4, globalState)
                    nw275_[int(7)] = d_1686_v89_
                    nw275_[int(8)] = _dafny.SeqWithoutIsStrInference([d_1565_v0_, d_1565_v0_, (0) - (d_1565_v0_), d_1565_v0_])
                    nw275_[int(9)] = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_1565_v0_})) for d_1695_i13_ in range(default__.abs(267))])
                    nw275_[int(10)] = d_1686_v89_
                    nw275_[int(11)] = d_1686_v89_
                    nw275_[int(12)] = d_1686_v89_
                    nw275_[int(13)] = (d_1689_v92_).cf31
                    nw275_[int(14)] = d_1686_v89_
                    nw275_[int(15)] = (d_1686_v89_) + (d_1686_v89_)
                    nw275_[int(16)] = d_1686_v89_
                    nw275_[int(17)] = d_1686_v89_
                    nw275_[int(18)] = (d_1686_v89_) + (d_1686_v89_)
                    nw275_[int(19)] = d_1686_v89_
                    nw275_[int(20)] = d_1686_v89_
                    nw275_[int(21)] = (d_1686_v89_) + (d_1686_v89_)
                    nw275_[int(22)] = ((((d_1690_v93_).cf34).set(default__.safeIndex(len(d_1692_v94_), len((d_1690_v93_).cf34)), d_1565_v0_)) + (d_1686_v89_)).set(default__.safeIndex(d_1565_v0_, len((((d_1690_v93_).cf34).set(default__.safeIndex(len(d_1692_v94_), len((d_1690_v93_).cf34)), d_1565_v0_)) + (d_1686_v89_))), len(d_1693_v95_))
                    nw275_[int(23)] = d_1686_v89_
                    d_1694_v96_ = nw275_
                    d_1696_v97_: _dafny.Map
                    d_1696_v97_ = _dafny.Map({d_1565_v0_: self.f20})
                    nw276_ = _dafny.Array(None, 8)
                    nw276_[int(0)] = d_1686_v89_
                    nw276_[int(1)] = d_1686_v89_
                    nw276_[int(2)] = (_dafny.SeqWithoutIsStrInference([884 for d_1697_i15_ in range(default__.abs(-200))])) + (d_1686_v89_)
                    nw276_[int(3)] = _dafny.SeqWithoutIsStrInference([d_1565_v0_, d_1565_v0_, d_1565_v0_, len(d_1696_v97_)])
                    nw276_[int(4)] = d_1686_v89_
                    nw276_[int(5)] = d_1686_v89_
                    nw276_[int(6)] = d_1686_v89_
                    nw276_[int(7)] = default__.fm44(not(self.f19), d_1565_v0_, self.f20, globalState)
                    d_1694_v96_ = nw276_
                    pass
            pass
        (globalState).f3 = d_1565_v0_
        d_1698_v98_: _dafny.MultiSet
        d_1698_v98_ = _dafny.MultiSet([d_1565_v0_, len((self).f9)])
        d_1699_v99_: _dafny.Seq
        d_1699_v99_ = _dafny.SeqWithoutIsStrInference([True, self.f20, self.f20])
        d_1700_v100_: _dafny.Array
        nw277_ = _dafny.Array(None, 24)
        nw277_[int(0)] = (d_1699_v99_)[default__.safeIndex(d_1565_v0_, len(d_1699_v99_))]
        nw277_[int(1)] = self.f20
        nw277_[int(2)] = self.f19
        nw277_[int(3)] = self.f19
        nw277_[int(4)] = self.f19
        nw277_[int(5)] = self.f19
        nw277_[int(6)] = self.f19
        nw277_[int(7)] = self.f19
        nw277_[int(8)] = self.f19
        nw277_[int(9)] = self.f19
        nw277_[int(10)] = self.f19
        nw277_[int(11)] = self.f19
        nw277_[int(12)] = self.f19
        nw277_[int(13)] = self.f19
        nw277_[int(14)] = self.f20
        nw277_[int(15)] = self.f20
        nw277_[int(16)] = True
        nw277_[int(17)] = self.f20
        nw277_[int(18)] = self.f19
        nw277_[int(19)] = self.f19
        nw277_[int(20)] = self.f20
        nw277_[int(21)] = True
        nw277_[int(22)] = self.f19
        nw277_[int(23)] = self.f20
        d_1700_v100_ = nw277_
        d_1701_v101_: str
        d_1701_v101_ = _dafny.CodePoint('l')
        d_1702_v102_: D2
        d_1702_v102_ = D2_DC5(self.f20, d_1700_v100_, len(default__.fm51(self.f19, d_1565_v0_, d_1701_v101_, self.f20, globalState)), d_1565_v0_)
        d_1565_v0_ = len(((d_1699_v99_ if self.f19 else d_1699_v99_) if ((d_1698_v98_).cardinality) != (len(d_1699_v99_)) else ((default__.fm2((self).f9, d_1565_v0_, globalState)) + (d_1699_v99_)).set(default__.safeIndex(d_1565_v0_, len((default__.fm2((self).f9, d_1565_v0_, globalState)) + (d_1699_v99_))), (d_1702_v102_).cf4)))
        d_1703_v103_: D14
        d_1703_v103_ = D14_DC33(d_1565_v0_, (self).f9, self.f19)
        if ((d_1703_v103_).cf41) <= ((self).f9):
            index241_ = default__.safeIndex(960, (d_1700_v100_).length(0))
            (d_1700_v100_)[index241_] = True
            index242_ = default__.safeIndex(960, (d_1700_v100_).length(0))
            (d_1700_v100_)[index242_] = not((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhhyskph")))) == (d_1565_v0_))
            d_1704_v104_: _dafny.Map
            d_1704_v104_ = _dafny.Map({d_1565_v0_: d_1565_v0_})
            d_1705_v105_: _dafny.Map
            d_1705_v105_ = _dafny.Map({d_1704_v104_: d_1704_v104_})
            d_1706_v106_: C1
            nw278_ = C1()
            nw278_.ctor__(default__.fm0(d_1565_v0_, (d_1698_v98_).cardinality, globalState), default__.fm0(d_1565_v0_, d_1565_v0_, globalState), d_1565_v0_, d_1705_v105_, ((self).f9 if self.f20 else (self).f9), (self).f10)
            d_1706_v106_ = nw278_
            (globalState).f3 = (self).fm5(d_1565_v0_, default__.safeDivisionInt(((d_1698_v98_)[d_1565_v0_] if (d_1565_v0_) in (d_1698_v98_) else len((self).f9)), (d_1706_v106_).f24), globalState)
            d_1707_v107_: _dafny.Array
            nw279_ = _dafny.Array(False, 10)
            d_1707_v107_ = nw279_
            pat_let_tv29_ = d_1707_v107_
            pat_let_tv30_ = d_1706_v106_
            pat_let_tv31_ = d_1565_v0_
            pat_let_tv32_ = globalState
            def iife150_(_pat_let40_0):
                def iife151_(d_1708_dt__update__tmp_h4_):
                    def iife152_(_pat_let41_0):
                        def iife153_(d_1709_dt__update_hcf5_h0_):
                            def iife154_(_pat_let42_0):
                                def iife155_(d_1710_dt__update_hcf7_h0_):
                                    return D2_DC5((d_1708_dt__update__tmp_h4_).cf4, d_1709_dt__update_hcf5_h0_, (d_1708_dt__update__tmp_h4_).cf6, d_1710_dt__update_hcf7_h0_)
                                return iife155_(_pat_let42_0)
                            return iife154_((self).fm5((pat_let_tv30_).f24, pat_let_tv31_, pat_let_tv32_))
                        return iife153_(_pat_let41_0)
                    return iife152_(pat_let_tv29_)
                return iife151_(_pat_let40_0)
            d_1707_v107_ = (iife150_(d_1702_v102_)).cf5
            index243_ = default__.safeIndex(960, (d_1700_v100_).length(0))
            (d_1700_v100_)[index243_] = self.f19
        elif True:
            (globalState).f3 = (0) - (len((self).f9))
            d_1701_v101_ = d_1701_v101_
            d_1711_v108_: _dafny.Array
            nw280_ = _dafny.Array(None, 6)
            nw280_[int(0)] = d_1565_v0_
            nw280_[int(1)] = default__.safeDivisionInt(d_1565_v0_, d_1565_v0_)
            nw280_[int(2)] = (0) - (len((self).f9))
            nw280_[int(3)] = d_1565_v0_
            nw280_[int(4)] = len(d_1699_v99_)
            nw280_[int(5)] = len((self).f9)
            d_1711_v108_ = nw280_
            index244_ = default__.safeIndex(950, (d_1711_v108_).length(0))
            (d_1711_v108_)[index244_] = d_1565_v0_
            d_1712_v109_: _dafny.Seq
            d_1712_v109_ = _dafny.SeqWithoutIsStrInference([631, d_1565_v0_])
            index245_ = default__.safeIndex(950, (d_1711_v108_).length(0))
            (d_1711_v108_)[index245_] = (len(d_1712_v109_)) + (d_1565_v0_)
            (self).f20 = (_dafny.CodePoint('b')) not in ((self).f9)
            d_1713_v110_: _dafny.Map
            d_1713_v110_ = _dafny.Map({self.f19: d_1565_v0_})
            d_1714_v111_: _dafny.Seq
            d_1714_v111_ = _dafny.SeqWithoutIsStrInference([d_1713_v110_, d_1713_v110_, d_1713_v110_])
            r0 = ((d_1713_v110_) not in (d_1714_v111_) if not(False) else ((0) - ((d_1711_v108_)[default__.safeIndex(950, (d_1711_v108_).length(0))])) <= ((d_1711_v108_)[default__.safeIndex(950, (d_1711_v108_).length(0))]))
        d_1715_v112_: D3
        d_1715_v112_ = D3_DC7(False)
        d_1716_v113_: D15
        d_1716_v113_ = D15_DC37((d_1702_v102_).cf6, d_1715_v112_, False)
        d_1717_v114_: _dafny.MultiSet
        d_1717_v114_ = _dafny.MultiSet([d_1716_v113_, d_1716_v113_, d_1716_v113_, d_1716_v113_, d_1716_v113_])
        (globalState).f3 = default__.safeDivisionInt(d_1565_v0_, ((d_1717_v114_)[D15_DC37(d_1565_v0_, d_1715_v112_, self.f20)] if (D15_DC37(d_1565_v0_, d_1715_v112_, self.f20)) in (d_1717_v114_) else d_1565_v0_))
        d_1718_v115_: _dafny.Seq
        d_1718_v115_ = _dafny.SeqWithoutIsStrInference([d_1565_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpcmhno"))), d_1565_v0_])
        d_1719_v116_: _dafny.Map
        d_1719_v116_ = _dafny.Map({d_1718_v115_: self.f20})
        r0 = ((d_1719_v116_)[_dafny.SeqWithoutIsStrInference([d_1565_v0_ for d_1720_i16_ in range(default__.abs(780))])] if (_dafny.SeqWithoutIsStrInference([d_1565_v0_ for d_1721_i16_ in range(default__.abs(780))])) in (d_1719_v116_) else self.f19)
        d_1722_v117_: _dafny.Seq
        d_1722_v117_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1701_v101_ for d_1723_i17_ in range(default__.abs(649))]), (self).f9, default__.fm31(globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "re"))])
        r1 = (self.f19 if self.f20 else default__.fm38(len(d_1722_v117_), globalState))
        return r0, r1


class C5:
    def  __init__(self):
        self.f18: _dafny.Array = _dafny.Array(None, 0)
        self._f17: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f17, f18):
        (self)._f17 = f17
        (self).f18 = f18

    def fm23(self, globalState):
        return default__.safeModuloInt((self).f17, default__.safeDivisionInt((self).f17, -570))

    def m11(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: D1 = D1.default()()
        r2: bool = False
        hi4_ = (self).f17
        for d_1724_i0_ in range(552, hi4_):
            d_1725_v0_: bool
            d_1725_v0_ = True
            d_1726_v1_: _dafny.Seq
            d_1726_v1_ = _dafny.SeqWithoutIsStrInference([d_1725_v0_])
            d_1727_v2_: _dafny.Set
            d_1727_v2_ = _dafny.Set({default__.fm38(d_1724_i0_, globalState)})
            d_1728_v3_: _dafny.Map
            d_1728_v3_ = _dafny.Map({p2: d_1724_i0_})
            d_1729_v4_: _dafny.Map
            d_1729_v4_ = _dafny.Map({d_1728_v3_: d_1728_v3_})
            d_1730_v5_: str
            d_1730_v5_ = _dafny.CodePoint('t')
            d_1731_v6_: _dafny.Seq
            d_1731_v6_ = _dafny.SeqWithoutIsStrInference([d_1730_v5_])
            d_1732_v7_: _dafny.Map
            d_1732_v7_ = _dafny.Map({p2: d_1731_v6_})
            d_1733_v8_: _dafny.Seq
            d_1733_v8_ = _dafny.SeqWithoutIsStrInference([p2])
            d_1734_v9_: _dafny.Seq
            d_1734_v9_ = _dafny.SeqWithoutIsStrInference([d_1733_v8_, d_1733_v8_, _dafny.SeqWithoutIsStrInference([p2, p2])])
            d_1735_v10_: C3
            nw281_ = C3()
            nw281_.ctor__((d_1726_v1_) == (d_1726_v1_), d_1724_i0_, (d_1724_i0_) - (len(d_1727_v2_)), d_1729_v4_, default__.fm26(p2, d_1732_v7_, globalState), (d_1734_v9_) + (d_1734_v9_))
            d_1735_v10_ = nw281_
            d_1736_v11_: D15
            d_1736_v11_ = D15_DC36(877)
            d_1737_v12_: D15
            d_1737_v12_ = D15_DC37((d_1736_v11_).cf48, D3_DC7(d_1725_v0_), (d_1735_v10_).f21)
            if ((d_1737_v12_).cf49) > (p2):
                (globalState).f3 = ((d_1735_v10_).f22) * ((0) - ((d_1735_v10_).f22))
                d_1738_v13_: int
                d_1739_v14_: _dafny.Map
                d_1740_v15_: bool
                d_1741_v16_: int
                out54_: int
                out55_: _dafny.Map
                out56_: bool
                out57_: int
                out54_, out55_, out56_, out57_ = (d_1735_v10_).m2((self).f17, globalState)
                d_1738_v13_ = out54_
                d_1739_v14_ = out55_
                d_1740_v15_ = out56_
                d_1741_v16_ = out57_
                d_1742_v17_: _dafny.Set
                d_1742_v17_ = _dafny.Set({d_1727_v2_, d_1727_v2_})
                d_1743_v18_: _dafny.MultiSet
                d_1743_v18_ = _dafny.MultiSet([(d_1742_v17_) - (d_1742_v17_), d_1742_v17_])
                d_1744_v19_: _dafny.Seq
                d_1744_v19_ = _dafny.SeqWithoutIsStrInference([d_1742_v17_, d_1742_v17_])
                d_1745_v20_: _dafny.MultiSet
                d_1745_v20_ = _dafny.MultiSet([len(p0), len(d_1731_v6_)])
                (globalState).f3 = ((d_1743_v18_)[((d_1744_v19_)[default__.safeIndex(d_1724_i0_, len(d_1744_v19_))]).intersection(default__.fm55((d_1745_v20_).cardinality, globalState))] if (((d_1744_v19_)[default__.safeIndex(d_1724_i0_, len(d_1744_v19_))]).intersection(default__.fm55((d_1745_v20_).cardinality, globalState))) in (d_1743_v18_) else (d_1735_v10_).f22)
                d_1746_v21_: _dafny.Array
                def lambda68_(d_1747_p2_):
                    def lambda69_(d_1748_i1_):
                        return (d_1748_i1_) * (d_1747_p2_)

                    return lambda69_

                init39_ = lambda68_(p2)
                nw282_ = _dafny.Array(None, 9)
                for i0_39_ in range(nw282_.length(0)):
                    nw282_[i0_39_] = init39_(i0_39_)
                d_1746_v21_ = nw282_
                d_1746_v21_ = d_1746_v21_
                d_1749_v22_: _dafny.Map
                d_1749_v22_ = _dafny.Map({(d_1735_v10_).f22: not(not((d_1735_v10_).f21))})
                d_1725_v0_ = ((d_1749_v22_)[p2] if (p2) in (d_1749_v22_) else (d_1735_v10_).f21)
            elif True:
                d_1725_v0_ = d_1725_v0_
                d_1750_v23_: _dafny.MultiSet
                d_1750_v23_ = _dafny.MultiSet([(d_1735_v10_).f21])
                r2 = ((d_1750_v23_).intersection(d_1750_v23_)) != (d_1750_v23_)
                r2 = (d_1735_v10_).f21
                d_1751_v24_: _dafny.Map
                d_1751_v24_ = _dafny.Map({default__.fm0(d_1724_i0_, 826, globalState): p1})
                d_1752_v25_: _dafny.Array
                nw283_ = _dafny.Array(None, 6)
                nw283_[int(0)] = p1
                nw283_[int(1)] = (((d_1751_v24_)[(d_1735_v10_).f22] if ((d_1735_v10_).f22) in (d_1751_v24_) else p1) if default__.fm38((d_1735_v10_).f22, globalState) else p1)
                nw283_[int(2)] = p1
                nw283_[int(3)] = p1
                nw283_[int(4)] = p1
                nw283_[int(5)] = p1
                d_1752_v25_ = nw283_
                d_1752_v25_ = d_1752_v25_
                d_1753_v26_: D12
                d_1753_v26_ = D12_DC29((d_1735_v10_).f21, d_1724_i0_, d_1733_v8_, (d_1735_v10_).f21, (self).f17)
                d_1725_v0_ = ((d_1725_v0_ if d_1725_v0_ else False) if (d_1753_v26_).cf35 else not(d_1725_v0_))
            r2 = (d_1735_v10_).f21
            d_1754_v27_: _dafny.MultiSet
            d_1754_v27_ = _dafny.MultiSet([_dafny.CodePoint('p'), default__.fm27(globalState)])
            d_1754_v27_ = (d_1754_v27_) | (d_1754_v27_)
        d_1755_v28_: bool
        d_1755_v28_ = True
        d_1756_v29_: D3
        d_1756_v29_ = D3_DC7(d_1755_v28_)
        source29_ = d_1756_v29_
        if source29_.is_DC7:
            d_1757___mcc_h0_ = source29_.cf9
            d_1758_cf9_ = d_1757___mcc_h0_
            d_1759_v30_: _dafny.Array
            def lambda70_(d_1760_i2_):
                return default__.safeModuloInt(d_1760_i2_, (self).f17)

            init40_ = lambda70_
            nw284_ = _dafny.Array(None, 1)
            for i0_40_ in range(nw284_.length(0)):
                nw284_[i0_40_] = init40_(i0_40_)
            d_1759_v30_ = nw284_
            d_1761_v31_: _dafny.Seq
            d_1761_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydfoo"))
            d_1762_v32_: _dafny.Map
            d_1762_v32_ = _dafny.Map({len(d_1761_v31_): default__.fm38(p2, globalState)})
            index246_ = default__.safeIndex(978, (d_1759_v30_).length(0))
            (d_1759_v30_)[index246_] = default__.safeModuloInt(p2, len(d_1762_v32_))
            index247_ = default__.safeIndex(978, (d_1759_v30_).length(0))
            (d_1759_v30_)[index247_] = 916
            r2 = default__.fm38(((d_1759_v30_)[default__.safeIndex(978, (d_1759_v30_).length(0))]) - ((self).f17), globalState)
            d_1763_v33_: str
            d_1763_v33_ = _dafny.CodePoint('v')
            d_1764_v34_: _dafny.Map
            d_1764_v34_ = _dafny.Map({(d_1759_v30_)[default__.safeIndex(978, (d_1759_v30_).length(0))]: d_1763_v33_})
            d_1765_v35_: _dafny.Map
            d_1765_v35_ = _dafny.Map({p2: p2})
            d_1766_v36_: _dafny.Seq
            d_1766_v36_ = _dafny.SeqWithoutIsStrInference([464, (self).f17])
            d_1767_v37_: _dafny.Seq
            d_1767_v37_ = _dafny.SeqWithoutIsStrInference([(0) - (p2), (d_1766_v36_)[default__.safeIndex((d_1759_v30_)[default__.safeIndex(978, (d_1759_v30_).length(0))], len(d_1766_v36_))]])
            d_1768_v38_: C2
            nw285_ = C2()
            nw285_.ctor__((d_1759_v30_)[default__.safeIndex(978, (d_1759_v30_).length(0))], _dafny.Map({d_1765_v35_: d_1765_v35_}), d_1761_v31_, (_dafny.SeqWithoutIsStrInference([d_1767_v37_, _dafny.SeqWithoutIsStrInference([883])])).set(default__.safeIndex((d_1759_v30_)[default__.safeIndex(978, (d_1759_v30_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_1767_v37_, _dafny.SeqWithoutIsStrInference([883])]))), _dafny.SeqWithoutIsStrInference([(0) - (len(d_1765_v35_)), len(d_1766_v36_)])))
            d_1768_v38_ = nw285_
            d_1769_v39_: _dafny.Set
            d_1769_v39_ = _dafny.Set({d_1768_v38_})
            d_1770_v40_: _dafny.Map
            d_1770_v40_ = _dafny.Map({d_1764_v34_: d_1769_v39_})
            d_1770_v40_ = (d_1770_v40_).set(d_1764_v34_, _dafny.Set({d_1768_v38_}))
            d_1771_v41_: _dafny.Set
            d_1771_v41_ = _dafny.Set({d_1767_v37_})
            d_1771_v41_ = (d_1771_v41_) | ((d_1771_v41_) | (d_1771_v41_))
        elif True:
            d_1772___mcc_h1_ = source29_.cf8
            d_1773_cf8_ = d_1772___mcc_h1_
            d_1774_v42_: _dafny.Array
            nw286_ = _dafny.Array(int(0), 26)
            d_1774_v42_ = nw286_
            d_1775_v43_: _dafny.Set
            d_1775_v43_ = _dafny.Set({d_1774_v42_, d_1774_v42_})
            d_1775_v43_ = (d_1775_v43_) - (d_1775_v43_)
            d_1776_v44_: _dafny.Map
            d_1776_v44_ = _dafny.Map({d_1755_v28_: default__.fm38(p2, globalState)})
            d_1777_v45_: _dafny.Seq
            d_1777_v45_ = _dafny.SeqWithoutIsStrInference([((d_1776_v44_).set(d_1755_v28_, d_1755_v28_)) | (d_1776_v44_)])
            rhs237_ = (d_1777_v45_) + (d_1777_v45_)
            rhs238_ = d_1755_v28_
            d_1777_v45_ = rhs237_
            d_1755_v28_ = rhs238_
            d_1778_v46_: _dafny.Seq
            d_1778_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
            d_1779_v47_: _dafny.Map
            d_1779_v47_ = _dafny.Map({d_1778_v46_: p2})
            (globalState).f3 = len(((_dafny.Map({d_1778_v46_: p2})).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")), p2)) | (d_1779_v47_))
            d_1780_v48_: _dafny.Array
            nw287_ = _dafny.Array(_dafny.Seq({}), 13)
            d_1780_v48_ = nw287_
            d_1781_v49_: _dafny.Seq
            d_1781_v49_ = _dafny.SeqWithoutIsStrInference([d_1755_v28_, not(not(True)), d_1755_v28_])
            index248_ = default__.safeIndex(971, (d_1780_v48_).length(0))
            (d_1780_v48_)[index248_] = d_1781_v49_
            index249_ = default__.safeIndex(971, (d_1780_v48_).length(0))
            (d_1780_v48_)[index249_] = default__.fm2((d_1778_v46_) + (d_1778_v46_), p2, globalState)
        hi5_ = len(_dafny.SeqWithoutIsStrInference([p2]))
        for d_1782_i3_ in range(630, hi5_):
            d_1783_v50_: _dafny.Map
            d_1783_v50_ = _dafny.Map({672: d_1755_v28_})
            d_1784_v51_: D4
            d_1784_v51_ = D4_DC10((d_1783_v50_) | (default__.fm46(globalState)))
            d_1784_v51_ = D4_DC10(d_1783_v50_)
            d_1785_v52_: _dafny.Map
            d_1785_v52_ = _dafny.Map({d_1755_v28_: d_1755_v28_})
            d_1786_v53_: _dafny.Seq
            d_1786_v53_ = _dafny.SeqWithoutIsStrInference([d_1782_i3_])
            d_1787_v54_: _dafny.Set
            d_1787_v54_ = _dafny.Set({(d_1786_v53_)[default__.safeIndex(p2, len(d_1786_v53_))]})
            d_1788_v55_: _dafny.Map
            d_1788_v55_ = _dafny.Map({d_1785_v52_: d_1787_v54_})
            d_1788_v55_ = (d_1788_v55_).set(d_1785_v52_, d_1787_v54_)
            d_1789_v56_: str
            d_1789_v56_ = _dafny.CodePoint('n')
            d_1789_v56_ = d_1789_v56_
            d_1790_v57_: D15
            d_1790_v57_ = D15_DC36((self).f17)
            d_1791_v58_: _dafny.Map
            d_1791_v58_ = _dafny.Map({(self).f17: ((d_1790_v57_).cf48) + (p2)})
            d_1792_v59_: D15
            d_1792_v59_ = D15_DC37((self).f17, d_1756_v29_, d_1755_v28_)
            d_1793_v60_: _dafny.Array
            nw288_ = _dafny.Array(D10.default()(), 29)
            d_1793_v60_ = nw288_
            d_1794_v61_: _dafny.Array
            d_1794_v61_ = d_1793_v60_
            d_1795_v62_: _dafny.Seq
            d_1795_v62_ = _dafny.SeqWithoutIsStrInference([d_1789_v56_])
            d_1796_v63_: _dafny.Map
            d_1796_v63_ = _dafny.Map({d_1794_v61_: len(d_1795_v62_)})
            d_1791_v58_ = (d_1791_v58_).set(d_1782_i3_, default__.fm0((d_1792_v59_).cf49, len(d_1796_v63_), globalState))
        pat_let_tv33_ = p2
        pat_let_tv34_ = d_1755_v28_
        pat_let_tv35_ = d_1755_v28_
        def lambda71_(source30_):
            if source30_.is_DC40:
                d_1797___mcc_h2_ = source30_.cf54
                d_1798___mcc_h3_ = source30_.cf55
                d_1799___mcc_h4_ = source30_.cf56
                d_1800_cf56_ = d_1799___mcc_h4_
                d_1801_cf55_ = d_1798___mcc_h3_
                d_1802_cf54_ = d_1797___mcc_h2_
                return pat_let_tv33_
            elif True:
                d_1803___mcc_h5_ = source30_.cf53
                d_1804_cf53_ = d_1803___mcc_h5_
                return (len(_dafny.Map({pat_let_tv34_: pat_let_tv35_}))) + ((self).f17)

        (globalState).f3 = lambda71_(default__.fm56((self).f17, (self).f17, d_1755_v28_, (self).f17, globalState))
        d_1805_v64_: _dafny.Array
        nw289_ = _dafny.Array(False, 19)
        d_1805_v64_ = nw289_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_1805_v64_).length(0)):
            d_1806_i4_: int = guard_loop_6_
            if (True) and (((0) <= (d_1806_i4_)) and ((d_1806_i4_) < ((d_1805_v64_).length(0)))):
                (d_1805_v64_)[(d_1806_i4_)] = d_1755_v28_
        if (d_1755_v28_) or (d_1755_v28_):
            (globalState).f3 = (p2 if d_1755_v28_ else (-900) - ((self).f17))
            d_1807_v65_: _dafny.Seq
            d_1807_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdqlc"))
            d_1808_v66_: _dafny.Seq
            d_1808_v66_ = _dafny.SeqWithoutIsStrInference([(self).f17, (self).f17])
            d_1809_v67_: _dafny.Seq
            d_1809_v67_ = _dafny.SeqWithoutIsStrInference([d_1808_v66_])
            d_1810_v68_: _dafny.Seq
            d_1810_v68_ = _dafny.SeqWithoutIsStrInference([d_1809_v67_])
            d_1811_v69_: C4
            nw290_ = C4()
            nw290_.ctor__(not(d_1755_v28_), d_1755_v28_, (d_1807_v65_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1812_i5_ in range(default__.abs(901))])), (d_1810_v68_)[default__.safeIndex(len(p0), len(d_1810_v68_))])
            d_1811_v69_ = nw290_
            (d_1811_v69_).f19 = (p2) >= ((0) - (p2))
            d_1813_v70_: C4
            nw291_ = C4()
            nw291_.ctor__(d_1811_v69_.f19, d_1811_v69_.f20, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))) + (d_1807_v65_), d_1809_v67_)
            d_1813_v70_ = nw291_
            d_1814_v71_: _dafny.Array
            nw292_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_1814_v71_ = nw292_
            d_1815_v72_: bool
            d_1816_v73_: bool
            out58_: bool
            out59_: bool
            out58_, out59_ = (d_1811_v69_).m14(d_1814_v71_, globalState)
            d_1815_v72_ = out58_
            d_1816_v73_ = out59_
        elif True:
            d_1817_v75_: _dafny.MultiSet
            d_1817_v75_ = _dafny.MultiSet([p2])
            d_1818_v76_: _dafny.Map
            d_1818_v76_ = _dafny.Map({len(p0): (self).f17})
            d_1819_v78_: _dafny.Map
            d_1819_v78_ = _dafny.Map({(self).f17: d_1755_v28_})
            d_1820_v81_: _dafny.Array
            nw293_ = _dafny.Array(None, 7)
            def iife156_():
                coll70_ = _dafny.Map()
                compr_70_: int
                for compr_70_ in (d_1817_v75_).Elements:
                    d_1821_v74_: int = compr_70_
                    if (d_1821_v74_) in (d_1817_v75_):
                        coll70_[default__.safeModuloInt(d_1821_v74_, (self).f17)] = p2
                return _dafny.Map(coll70_)
            nw293_[int(0)] = iife156_()
            
            nw293_[int(1)] = d_1818_v76_
            nw293_[int(2)] = d_1818_v76_
            nw293_[int(3)] = d_1818_v76_
            nw293_[int(4)] = (d_1818_v76_) | (d_1818_v76_)
            nw293_[int(5)] = d_1818_v76_
            def iife157_():
                def iife158_():
                    def iife160_():
                        coll74_ = _dafny.Map()
                        compr_74_: int
                        for compr_74_ in (_dafny.Map({p2: _dafny.CodePoint('w')})).keys.Elements:
                            d_1823_v80_: int = compr_74_
                            if (d_1823_v80_) in (_dafny.Map({p2: _dafny.CodePoint('w')})):
                                coll74_[(d_1823_v80_) + ((self).f17)] = d_1755_v28_
                        return _dafny.Map(coll74_)
                    coll72_ = _dafny.Map()
                    def iife159_():
                        coll73_ = _dafny.Map()
                        compr_73_: int
                        for compr_73_ in (_dafny.Map({p2: _dafny.CodePoint('w')})).keys.Elements:
                            d_1823_v80_: int = compr_73_
                            if (d_1823_v80_) in (_dafny.Map({p2: _dafny.CodePoint('w')})):
                                coll73_[(d_1823_v80_) + ((self).f17)] = d_1755_v28_
                        return _dafny.Map(coll73_)
                    compr_72_: int
                    for compr_72_ in (iife159_()
                    ).keys.Elements:
                        d_1824_v79_: int = compr_72_
                        if (d_1824_v79_) in (iife160_()
                        ):
                            coll72_[(d_1824_v79_) * (p2)] = d_1755_v28_
                    return _dafny.Map(coll72_)
                coll71_ = _dafny.Map()
                compr_71_: int
                for compr_71_ in (d_1819_v78_).keys.Elements:
                    d_1822_v77_: int = compr_71_
                    if (d_1822_v77_) in (d_1819_v78_):
                        coll71_[default__.safeModuloInt(d_1822_v77_, len(iife158_()
                        ))] = p2
                return _dafny.Map(coll71_)
            nw293_[int(6)] = (iife157_()
            ) | (_dafny.Map({p2: p2}))
            d_1820_v81_ = nw293_
            index250_ = default__.safeIndex(667, (p1).length(0))
            (p1)[index250_] = d_1755_v28_
            d_1825_v83_: _dafny.Map
            def iife161_():
                coll75_ = _dafny.Set()
                compr_75_: int
                for compr_75_ in _dafny.IntegerRange(-653, 756):
                    d_1826_v82_: int = compr_75_
                    if ((-653) <= (d_1826_v82_)) and ((d_1826_v82_) < (756)):
                        coll75_ = coll75_.union(_dafny.Set([(d_1826_v82_) * (p2)]))
                return _dafny.Set(coll75_)
            d_1825_v83_ = _dafny.Map({(0) - (len(iife161_()
            )): d_1820_v81_})
            d_1827_v84_: _dafny.Seq
            d_1827_v84_ = _dafny.SeqWithoutIsStrInference([((d_1825_v83_)[(self).f17] if ((self).f17) in (d_1825_v83_) else d_1820_v81_)])
            index251_ = default__.safeIndex(667, (p1).length(0))
            rhs239_ = default__.fm38(default__.fm0(len(p0), (self).f17, globalState), globalState)
            rhs240_ = (d_1827_v84_)[default__.safeIndex(-945, len(d_1827_v84_))]
            rhs241_ = not(not (d_1755_v28_) or (d_1755_v28_))
            lhs154_ = p1
            lhs155_ = default__.safeIndex(667, (p1).length(0))
            d_1755_v28_ = rhs239_
            d_1820_v81_ = rhs240_
            lhs154_[lhs155_] = rhs241_
            d_1828_v85_: _dafny.Set
            d_1828_v85_ = _dafny.Set({p2})
            index252_ = default__.safeIndex(667, (p1).length(0))
            (p1)[index252_] = ((p1)[default__.safeIndex(667, (p1).length(0))] if d_1755_v28_ else (p2) in (d_1828_v85_))
            d_1829_v86_: _dafny.Array
            nw294_ = _dafny.Array(_dafny.Map({}), 27)
            d_1829_v86_ = nw294_
            d_1830_v88_: D4
            def iife162_():
                coll76_ = _dafny.Map()
                compr_76_: int
                for compr_76_ in _dafny.IntegerRange(975, -392):
                    d_1831_v87_: int = compr_76_
                    if ((975) <= (d_1831_v87_)) and ((d_1831_v87_) < (-392)):
                        coll76_[(d_1831_v87_) - ((_dafny.MultiSet([_dafny.CodePoint('v'), _dafny.CodePoint('w'), _dafny.CodePoint('q')])).cardinality)] = (p1)[default__.safeIndex(667, (p1).length(0))]
                return _dafny.Map(coll76_)
            d_1830_v88_ = D4_DC10(iife162_()
)
            d_1832_v89_: _dafny.Map
            d_1832_v89_ = _dafny.Map({d_1829_v86_: d_1830_v88_})
            d_1832_v89_ = (d_1832_v89_).set(d_1829_v86_, d_1830_v88_)
            index253_ = default__.safeIndex(667, (p1).length(0))
            (p1)[index253_] = False
            r2 = d_1755_v28_
        d_1833_v90_: _dafny.Map
        d_1833_v90_ = _dafny.Map({(self).f17: p2})
        d_1834_v91_: _dafny.Map
        d_1834_v91_ = _dafny.Map({d_1755_v28_: d_1755_v28_})
        d_1835_v92_: _dafny.Seq
        d_1835_v92_ = _dafny.SeqWithoutIsStrInference([(d_1833_v90_).set(len(d_1834_v91_), p2)])
        r0 = d_1835_v92_
        d_1836_v93_: D1
        d_1836_v93_ = D1_DC1(d_1755_v28_)
        r1 = d_1836_v93_
        d_1837_v94_: _dafny.Seq
        d_1837_v94_ = _dafny.SeqWithoutIsStrInference([d_1755_v28_])
        d_1838_v95_: _dafny.Seq
        d_1838_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twpkxkkm"))
        d_1839_v96_: _dafny.Seq
        d_1839_v96_ = _dafny.SeqWithoutIsStrInference([default__.fm0(len(p0), (self).f17, globalState)])
        d_1840_v97_: _dafny.Seq
        d_1840_v97_ = _dafny.SeqWithoutIsStrInference([d_1839_v96_])
        d_1841_v98_: C1
        nw295_ = C1()
        nw295_.ctor__(43, p2, p2, default__.fm47(d_1755_v28_, d_1755_v28_, not(d_1755_v28_), (d_1837_v94_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aw"))), len(d_1837_v94_))], globalState), d_1838_v95_, d_1840_v97_)
        d_1841_v98_ = nw295_
        d_1842_v99_: _dafny.MultiSet
        d_1842_v99_ = _dafny.MultiSet([d_1841_v98_])
        d_1843_v100_: _dafny.Map
        d_1843_v100_ = _dafny.Map({d_1755_v28_: d_1842_v99_})
        r2 = not((d_1843_v100_) != (d_1843_v100_))
        return r0, r1, r2

    def m12(self, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: D7 = D7.default()()
        r3: str = _dafny.CodePoint('D')
        d_1844_v0_: _dafny.Seq
        d_1844_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pupyabfjy"))
        d_1845_v1_: bool
        d_1845_v1_ = True
        d_1846_v2_: _dafny.Map
        d_1846_v2_ = _dafny.Map({(self).f17: d_1845_v1_})
        d_1847_v3_: _dafny.Map
        d_1847_v3_ = _dafny.Map({(self).f17: 875})
        d_1848_v4_: _dafny.MultiSet
        d_1848_v4_ = _dafny.MultiSet([d_1845_v1_])
        d_1849_v5_: _dafny.Map
        d_1849_v5_ = _dafny.Map({d_1848_v4_: (self).f17})
        d_1850_v6_: _dafny.Seq
        d_1850_v6_ = _dafny.SeqWithoutIsStrInference([(self).f17, len(d_1849_v5_), (self).f17, (self).f17])
        d_1851_v7_: _dafny.Array
        nw296_ = _dafny.Array(None, 18)
        nw296_[int(0)] = (0) - (len(d_1844_v0_))
        nw296_[int(1)] = (0) - (default__.safeModuloInt((0) - ((self).f17), (self).f17))
        nw296_[int(2)] = (self).f17
        nw296_[int(3)] = (self).f17
        nw296_[int(4)] = (self).f17
        nw296_[int(5)] = default__.safeDivisionInt((self).f17, (self).f17)
        nw296_[int(6)] = (self).f17
        nw296_[int(7)] = ((self).f17 if d_1845_v1_ else (self).f17)
        nw296_[int(8)] = ((self).f17) - ((self).f17)
        nw296_[int(9)] = (self).f17
        nw296_[int(10)] = ((self).f17) - ((self).f17)
        nw296_[int(11)] = (self).f17
        nw296_[int(12)] = default__.safeModuloInt((self).f17, (0) - ((self).f17))
        nw296_[int(13)] = (self).f17
        nw296_[int(14)] = ((self).f17) * ((self).f17)
        nw296_[int(15)] = (self).f17
        nw296_[int(16)] = default__.fm0(len(d_1846_v2_), len(d_1847_v3_), globalState)
        nw296_[int(17)] = (d_1850_v6_)[default__.safeIndex((self).f17, len(d_1850_v6_))]
        d_1851_v7_ = nw296_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_1851_v7_).length(0)):
            d_1852_i0_: int = guard_loop_7_
            if (True) and (((0) <= (d_1852_i0_)) and ((d_1852_i0_) < ((d_1851_v7_).length(0)))):
                (d_1851_v7_)[(d_1852_i0_)] = default__.safeModuloInt(d_1852_i0_, (self).f17)
        d_1853_v8_: D6
        d_1853_v8_ = D6_DC14(d_1851_v7_)
        d_1854_v9_: _dafny.Map
        d_1854_v9_ = _dafny.Map({d_1853_v8_: default__.safeModuloInt((self).f17, (self).f17)})
        d_1854_v9_ = (d_1854_v9_).set(D6_DC14(d_1851_v7_), len(d_1844_v0_))
        d_1855_v10_: _dafny.Seq
        d_1855_v10_ = _dafny.SeqWithoutIsStrInference([True, d_1845_v1_])
        d_1848_v4_ = _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([d_1845_v1_])) + (d_1855_v10_)).set(default__.safeIndex(((d_1847_v3_)[(self).f17] if ((self).f17) in (d_1847_v3_) else (self).f17), len((_dafny.SeqWithoutIsStrInference([d_1845_v1_])) + (d_1855_v10_))), d_1845_v1_))
        d_1856_v11_: D3
        d_1856_v11_ = D3_DC7(False)
        d_1857_v12_: _dafny.Set
        d_1857_v12_ = _dafny.Set({True, (d_1856_v11_).cf9, d_1845_v1_})
        (globalState).f3 = len(d_1857_v12_)
        index254_ = default__.safeIndex(110, (d_1851_v7_).length(0))
        (d_1851_v7_)[index254_] = (self).fm23(globalState)
        index255_ = default__.safeIndex(110, (d_1851_v7_).length(0))
        (d_1851_v7_)[index255_] = (len(_dafny.SeqWithoutIsStrInference([default__.fm38(len(d_1847_v3_), globalState)]))) + (default__.fm0((self).f17, len(d_1844_v0_), globalState))
        d_1858_v13_: _dafny.Map
        d_1858_v13_ = _dafny.Map({d_1847_v3_: d_1847_v3_})
        d_1859_v14_: _dafny.Seq
        d_1859_v14_ = _dafny.SeqWithoutIsStrInference([d_1850_v6_])
        d_1860_v15_: _dafny.Seq
        d_1860_v15_ = _dafny.SeqWithoutIsStrInference([(d_1859_v14_)[default__.safeIndex(820, len(d_1859_v14_))], _dafny.SeqWithoutIsStrInference([len(d_1847_v3_)]), d_1850_v6_])
        d_1861_v16_: C2
        nw297_ = C2()
        nw297_.ctor__((d_1851_v7_)[default__.safeIndex(110, (d_1851_v7_).length(0))], d_1858_v13_, d_1844_v0_, d_1860_v15_)
        d_1861_v16_ = nw297_
        r0 = _dafny.CodePoint('s')
        r1 = d_1844_v0_
        d_1862_v17_: D14
        d_1862_v17_ = D14_DC33((d_1851_v7_)[default__.safeIndex(110, (d_1851_v7_).length(0))], d_1844_v0_, d_1845_v1_)
        d_1863_v18_: _dafny.Map
        d_1863_v18_ = _dafny.Map({d_1862_v17_: d_1845_v1_})
        pat_let_tv36_ = d_1845_v1_
        def lambda72_(source31_):
            if source31_.is_DC9:
                d_1864___mcc_h0_ = source31_.cf11
                d_1865_cf11_ = d_1864___mcc_h0_
                return D7_DC16()
            elif source31_.is_DC10:
                d_1866___mcc_h1_ = source31_.cf12
                d_1867_cf12_ = d_1866___mcc_h1_
                return D7_DC16()
            elif source31_.is_DC11:
                d_1868___mcc_h2_ = source31_.cf13
                d_1869___mcc_h3_ = source31_.cf14
                d_1870_cf14_ = d_1869___mcc_h3_
                d_1871_cf13_ = d_1868___mcc_h2_
                if pat_let_tv36_:
                    return D7_DC16()
                elif True:
                    return D7_DC16()
            elif True:
                d_1872___mcc_h4_ = source31_.cf10
                d_1873_cf10_ = d_1872___mcc_h4_
                return D7_DC16()

        r2 = lambda72_(D4_DC11(d_1845_v1_, len(d_1863_v18_)))
        r3 = default__.fm27(globalState)
        return r0, r1, r2, r3

    @property
    def f17(self):
        return self._f17

class C6(T1, T0):
    def  __init__(self):
        self._f12: _dafny.Map = _dafny.Map({})
        self._f11: int = int(0)
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f11(self):
        return self._f11
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f11, f12, f9, f10):
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f9 = f9
        (self)._f10 = f10

    def fm5(self, p0, p1, globalState):
        return (((_dafny.MultiSet([(self).f11, (self).f11])) | (_dafny.MultiSet([(self).f11]))).intersection(((D7_DC15(_dafny.MultiSet([(self).f11, (self).f11, (self).f11, (self).f11, (self).f11]))).cf18) | (_dafny.MultiSet([(self).f11, (self).f11])))).cardinality

    def fm19(self, globalState):
        return D4_DC8((_dafny.Map({(self).f11: (D3_DC7(True)).cf9})) | (_dafny.Map({(self).f11: False})))

    def fm20(self, p0, p1, p2, globalState):
        return _dafny.Map({((self).f11) <= ((self).f11): (self).f11})

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: bool = False
        r3: int = int(0)
        d_1874_v0_: _dafny.Seq
        d_1874_v0_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1875_v1_: _dafny.MultiSet
        d_1875_v1_ = _dafny.MultiSet([d_1874_v0_, d_1874_v0_])
        if ((d_1875_v1_) | (_dafny.MultiSet(default__.fm21((self).f9, globalState)))).issubset((d_1875_v1_).intersection(d_1875_v1_)):
            d_1876_v2_: _dafny.Array
            nw298_ = _dafny.Array(_dafny.Map({}), 16)
            d_1876_v2_ = nw298_
            d_1877_v3_: _dafny.Seq
            d_1877_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "svrqwnup"))
            d_1878_v4_: bool
            d_1878_v4_ = True
            d_1879_v5_: _dafny.MultiSet
            d_1879_v5_ = _dafny.MultiSet([default__.fm22(d_1878_v4_, (self).f11, globalState), d_1878_v4_])
            rhs242_ = ((self).f11) + ((self).f11)
            rhs243_ = d_1876_v2_
            rhs244_ = (self).f9
            rhs245_ = (d_1879_v5_).ispropersubset(d_1879_v5_)
            lhs156_ = globalState
            lhs156_.f3 = rhs242_
            d_1876_v2_ = rhs243_
            d_1877_v3_ = rhs244_
            r2 = rhs245_
            d_1880_v6_: C4
            nw299_ = C4()
            nw299_.ctor__((d_1878_v4_ if d_1878_v4_ else d_1878_v4_), d_1878_v4_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mn")), ((self).f10) + ((self).f10))
            d_1880_v6_ = nw299_
            d_1881_v7_: _dafny.Set
            d_1881_v7_ = _dafny.Set({len(_dafny.Map({d_1878_v4_: d_1874_v0_})), p0})
            d_1882_v8_: _dafny.Map
            d_1882_v8_ = _dafny.Map({(0) - ((self).f11): (len(d_1881_v7_)) + (len(_dafny.Map({p0: (self).f11})))})
            d_1882_v8_ = d_1882_v8_
            (globalState).f3 = (0) - ((self).f11)
            d_1883_v9_: _dafny.Seq
            d_1883_v9_ = _dafny.SeqWithoutIsStrInference([d_1880_v6_.f19])
            d_1884_v10_: _dafny.Set
            d_1884_v10_ = _dafny.Set({d_1883_v9_})
            d_1885_v11_: _dafny.Map
            d_1885_v11_ = _dafny.Map({p0: d_1878_v4_})
            d_1886_v12_: _dafny.Set
            d_1886_v12_ = _dafny.Set({d_1880_v6_.f19})
            d_1887_v13_: C2
            nw300_ = C2()
            nw300_.ctor__(len(d_1885_v11_), self.f12, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxc")), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_1886_v12_)]), _dafny.SeqWithoutIsStrInference([default__.fm0((self).f11, p0, globalState)])]))
            d_1887_v13_ = nw300_
            d_1888_v14_: _dafny.Map
            d_1888_v14_ = _dafny.Map({d_1887_v13_: False})
            d_1889_v15_: str
            d_1889_v15_ = _dafny.CodePoint('e')
            d_1890_v16_: _dafny.Map
            d_1890_v16_ = _dafny.Map({(len(d_1884_v10_)) * (len(default__.fm57(d_1878_v4_, ((d_1888_v14_)[d_1887_v13_] if (d_1887_v13_) in (d_1888_v14_) else not(d_1880_v6_.f19)), d_1889_v15_, p0, globalState))): d_1880_v6_.f19})
            d_1891_v17_: D1
            d_1891_v17_ = D1_DC1(d_1880_v6_.f19)
            d_1885_v11_ = (d_1885_v11_).set(p0, (not(d_1878_v4_) if d_1878_v4_ else (d_1891_v17_).cf1))
        elif True:
            d_1892_v18_: D8
            d_1892_v18_ = D8_DC18()
            d_1893_v19_: D8
            d_1893_v19_ = D8_DC19(D8_DC19(d_1892_v18_))
            source32_ = d_1893_v19_
            if source32_.is_DC18:
                d_1894_v20_: bool
                d_1894_v20_ = False
                r2 = not(d_1894_v20_)
                d_1894_v20_ = not(d_1894_v20_)
                d_1895_v21_: _dafny.Array
                def lambda73_(d_1896_p0_):
                    def lambda74_(d_1897_i0_):
                        return _dafny.SeqWithoutIsStrInference([d_1896_p0_, d_1896_p0_])

                    return lambda74_

                init41_ = lambda73_(p0)
                nw301_ = _dafny.Array(None, 8)
                for i0_41_ in range(nw301_.length(0)):
                    nw301_[i0_41_] = init41_(i0_41_)
                d_1895_v21_ = nw301_
                index256_ = default__.safeIndex(887, (d_1895_v21_).length(0))
                (d_1895_v21_)[index256_] = d_1874_v0_
                index257_ = default__.safeIndex(887, (d_1895_v21_).length(0))
                (d_1895_v21_)[index257_] = _dafny.SeqWithoutIsStrInference([(self).f11, p0, (self).f11, (784 if d_1894_v20_ else (self).f11)])
                d_1898_v22_: _dafny.Map
                d_1898_v22_ = _dafny.Map({d_1894_v20_: (self).f11})
                d_1898_v22_ = (d_1898_v22_).set(d_1894_v20_, default__.safeDivisionInt((self).f11, p0))
            elif source32_.is_DC17:
                d_1899___mcc_h0_ = source32_.cf19
                d_1900_cf19_ = d_1899___mcc_h0_
                d_1901_v23_: _dafny.Map
                d_1901_v23_ = _dafny.Map({p0: (self).f9})
                d_1902_v24_: str
                d_1902_v24_ = _dafny.CodePoint('s')
                d_1903_v25_: _dafny.Seq
                d_1903_v25_ = _dafny.SeqWithoutIsStrInference([(self).f9, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxr"))])
                d_1904_v26_: bool
                d_1904_v26_ = True
                d_1905_v27_: _dafny.Map
                d_1905_v27_ = _dafny.Map({d_1904_v26_: -918})
                d_1906_v28_: _dafny.Array
                nw302_ = _dafny.Array(None, 14)
                nw302_[int(0)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1907_i1_ in range(default__.abs(658))])
                nw302_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fixhi"))
                nw302_[int(2)] = default__.fm26(462, d_1901_v23_, globalState)
                nw302_[int(3)] = ((self).f9).set(default__.safeIndex(p0, len((self).f9)), d_1902_v24_)
                nw302_[int(4)] = ((self).f9) + ((self).f9)
                nw302_[int(5)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dquiue"))) + ((d_1903_v25_)[default__.safeIndex(p0, len(d_1903_v25_))])
                nw302_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "weygk"))
                nw302_[int(7)] = ((d_1903_v25_)[default__.safeIndex(len(d_1905_v27_), len(d_1903_v25_))]) + ((self).f9)
                nw302_[int(8)] = ((self).f9) + (default__.fm31(globalState))
                nw302_[int(9)] = (self).f9
                nw302_[int(10)] = _dafny.SeqWithoutIsStrInference([d_1902_v24_ for d_1908_i2_ in range(default__.abs(-608))])
                nw302_[int(11)] = ((self).f9) + ((self).f9)
                nw302_[int(12)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjuipbpmt"))).set(default__.safeIndex((self).f11, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjuipbpmt")))), d_1902_v24_)
                nw302_[int(13)] = (self).f9
                d_1906_v28_ = nw302_
                nw303_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
                d_1906_v28_ = nw303_
                d_1909_v29_: _dafny.Array
                nw304_ = _dafny.Array(int(0), 26)
                d_1909_v29_ = nw304_
                index258_ = default__.safeIndex(953, (d_1909_v29_).length(0))
                (d_1909_v29_)[index258_] = p0
                index259_ = default__.safeIndex(953, (d_1909_v29_).length(0))
                (d_1909_v29_)[index259_] = -797
                d_1910_v31_: _dafny.Set
                d_1910_v31_ = _dafny.Set({d_1904_v26_})
                d_1911_v32_: _dafny.Map
                d_1911_v32_ = _dafny.Map({d_1904_v26_: d_1904_v26_})
                d_1912_v33_: _dafny.Map
                d_1912_v33_ = _dafny.Map({len(d_1910_v31_): len(d_1911_v32_)})
                d_1913_v34_: _dafny.Set
                d_1913_v34_ = _dafny.Set({d_1912_v33_})
                d_1914_v35_: C3
                nw305_ = C3()
                def iife163_():
                    coll77_ = _dafny.Map()
                    compr_77_: _dafny.Map
                    for compr_77_ in (d_1913_v34_).Elements:
                        d_1915_v30_: _dafny.Map = compr_77_
                        if (d_1915_v30_) in (d_1913_v34_):
                            coll77_[d_1915_v30_] = d_1912_v33_
                    return _dafny.Map(coll77_)
                nw305_.ctor__(d_1904_v26_, -490, (0) - ((0) - ((0) - ((self).f11))), (iife163_()
                ) | (self.f12), (self).f9, (_dafny.SeqWithoutIsStrInference([d_1874_v0_ for d_1916_i3_ in range(default__.abs(140))])) + ((self).f10))
                d_1914_v35_ = nw305_
                d_1917_v36_: _dafny.Array
                nw306_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_1917_v36_ = nw306_
                d_1917_v36_ = (d_1917_v36_ if d_1904_v26_ else d_1917_v36_)
            elif True:
                d_1918___mcc_h1_ = source32_.cf20
                d_1919_cf20_ = d_1918___mcc_h1_
                d_1920_v37_: bool
                d_1920_v37_ = False
                r2 = d_1920_v37_
                d_1921_v38_: _dafny.Map
                d_1921_v38_ = _dafny.Map({d_1920_v37_: (self).f11})
                d_1921_v38_ = ((_dafny.Map({d_1920_v37_: p0})) | (d_1921_v38_)).set(d_1920_v37_, (self).f11)
                d_1922_v39_: _dafny.Array
                nw307_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_1922_v39_ = nw307_
                d_1923_v40_: _dafny.Array
                nw308_ = _dafny.Array(_dafny.Map({}), 2)
                d_1923_v40_ = nw308_
                index260_ = default__.safeIndex(325, (d_1922_v39_).length(0))
                (d_1922_v39_)[index260_] = d_1923_v40_
                d_1924_v41_: _dafny.Map
                d_1924_v41_ = _dafny.Map({p0: d_1920_v37_})
                d_1925_v42_: _dafny.Map
                d_1925_v42_ = _dafny.Map({True: d_1924_v41_})
                index261_ = default__.safeIndex(325, (d_1922_v39_).length(0))
                rhs246_ = ((p0) - ((self).f11)) not in (d_1874_v0_)
                rhs247_ = d_1920_v37_
                rhs248_ = d_1923_v40_
                rhs249_ = (default__.safeDivisionInt(len(d_1925_v42_), (self).f11)) * ((self).f11)
                lhs157_ = d_1922_v39_
                lhs158_ = default__.safeIndex(325, (d_1922_v39_).length(0))
                lhs159_ = globalState
                d_1920_v37_ = rhs246_
                d_1920_v37_ = rhs247_
                lhs157_[lhs158_] = rhs248_
                lhs159_.f3 = rhs249_
                d_1926_v43_: _dafny.MultiSet
                d_1926_v43_ = _dafny.MultiSet([-267, (d_1874_v0_)[default__.safeIndex((self).f11, len(d_1874_v0_))], p0, -835])
                d_1926_v43_ = (d_1926_v43_).intersection(d_1926_v43_)
            d_1927_v44_: _dafny.Seq
            d_1927_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlljumshm"))
            d_1927_v44_ = d_1927_v44_
            d_1928_v45_: bool
            d_1928_v45_ = False
            if (d_1928_v45_ if True else d_1928_v45_):
                r2 = d_1928_v45_
                r2 = (p0) < ((self).f11)
                d_1929_v46_: _dafny.Array
                def lambda75_(d_1930_v45_):
                    def lambda76_(d_1931_i4_):
                        return d_1930_v45_

                    return lambda76_

                init42_ = lambda75_(d_1928_v45_)
                nw309_ = _dafny.Array(None, 22)
                for i0_42_ in range(nw309_.length(0)):
                    nw309_[i0_42_] = init42_(i0_42_)
                d_1929_v46_ = nw309_
                nw310_ = _dafny.Array(False, 2)
                d_1929_v46_ = nw310_
                r2 = ((0) - (p0)) < ((self).f11)
                r2 = (False) and (d_1928_v45_)
            elif True:
                d_1932_v47_: _dafny.Map
                d_1932_v47_ = _dafny.Map({d_1928_v45_: d_1928_v45_})
                d_1933_v48_: _dafny.Set
                d_1933_v48_ = _dafny.Set({d_1928_v45_, d_1928_v45_, (d_1928_v45_) == (((d_1932_v47_)[False] if (False) in (d_1932_v47_) else False)), d_1928_v45_, d_1928_v45_})
                (globalState).f3 = len(d_1933_v48_)
                d_1934_v49_: str
                d_1934_v49_ = _dafny.CodePoint('g')
                d_1934_v49_ = d_1934_v49_
                d_1935_v50_: C0
                nw311_ = C0()
                nw311_.ctor__((d_1928_v45_) and (d_1928_v45_), d_1928_v45_, ((D17_DC40(d_1927_v44_, not(d_1928_v45_), p0)).cf54) + (d_1927_v44_), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, 296, (0) - (p0), (self).f11, (self).f11])]))
                d_1935_v50_ = nw311_
                d_1936_v51_: _dafny.Set
                d_1936_v51_ = _dafny.Set({(self).f11})
                d_1928_v45_ = ((D9_DC21(D3_DC6(d_1936_v51_), d_1934_v49_, (d_1935_v50_).f26)).cf24) == ((d_1874_v0_) <= (d_1874_v0_))
                d_1937_v52_: _dafny.Array
                nw312_ = _dafny.Array(int(0), 21)
                d_1937_v52_ = nw312_
                d_1938_v53_: _dafny.Set
                d_1938_v53_ = _dafny.Set({_dafny.Set({d_1928_v45_}), d_1933_v48_})
                index262_ = default__.safeIndex(724, (d_1937_v52_).length(0))
                (d_1937_v52_)[index262_] = (len(d_1938_v53_)) * ((self).f11)
                d_1939_v54_: _dafny.Seq
                d_1939_v54_ = _dafny.SeqWithoutIsStrInference([d_1928_v45_, d_1928_v45_, (d_1935_v50_).f26, (d_1935_v50_).f25])
                index263_ = default__.safeIndex(724, (d_1937_v52_).length(0))
                rhs250_ = 419
                rhs251_ = ((self).f11) - (788)
                rhs252_ = (d_1933_v48_) | ((d_1933_v48_) | (d_1933_v48_))
                rhs253_ = default__.fm0((self).f11, len((d_1939_v54_) + (d_1939_v54_)), globalState)
                lhs160_ = globalState
                lhs161_ = globalState
                lhs162_ = d_1937_v52_
                lhs163_ = default__.safeIndex(724, (d_1937_v52_).length(0))
                lhs160_.f3 = rhs250_
                lhs161_.f3 = rhs251_
                d_1933_v48_ = rhs252_
                lhs162_[lhs163_] = rhs253_
            d_1940_v55_: _dafny.Array
            nw313_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_1940_v55_ = nw313_
            d_1941_v56_: _dafny.Seq
            d_1941_v56_ = _dafny.SeqWithoutIsStrInference([d_1940_v55_, d_1940_v55_])
            d_1942_v57_: _dafny.Map
            d_1942_v57_ = _dafny.Map({d_1928_v45_: d_1940_v55_})
            d_1943_v58_: _dafny.Array
            nw314_ = _dafny.Array(None, 4)
            nw314_[int(0)] = (d_1941_v56_)[default__.safeIndex(p0, len(d_1941_v56_))]
            nw314_[int(1)] = d_1940_v55_
            nw314_[int(2)] = ((d_1942_v57_)[False] if (False) in (d_1942_v57_) else d_1940_v55_)
            nw314_[int(3)] = (d_1940_v55_ if d_1928_v45_ else d_1940_v55_)
            d_1943_v58_ = nw314_
            index264_ = default__.safeIndex(543, (d_1943_v58_).length(0))
            (d_1943_v58_)[index264_] = d_1940_v55_
            d_1944_v59_: _dafny.Map
            d_1944_v59_ = _dafny.Map({(self).f9: d_1940_v55_})
            index265_ = default__.safeIndex(543, (d_1943_v58_).length(0))
            (d_1943_v58_)[index265_] = ((d_1944_v59_)[d_1927_v44_] if (d_1927_v44_) in (d_1944_v59_) else d_1940_v55_)
            d_1945_v60_: D12
            d_1945_v60_ = D12_DC28(default__.fm44(d_1928_v45_, len((self).f9), d_1928_v45_, globalState))
            d_1946_v61_: _dafny.Map
            d_1946_v61_ = _dafny.Map({d_1945_v60_: ((self).f11) + (p0)})
            d_1946_v61_ = d_1946_v61_
        d_1947_v62_: _dafny.Array
        def lambda77_(d_1948_i5_):
            return False

        init43_ = lambda77_
        nw315_ = _dafny.Array(None, 13)
        for i0_43_ in range(nw315_.length(0)):
            nw315_[i0_43_] = init43_(i0_43_)
        d_1947_v62_ = nw315_
        d_1949_v63_: _dafny.Map
        d_1949_v63_ = _dafny.Map({(self).f11: d_1947_v62_})
        d_1950_v64_: _dafny.Map
        d_1950_v64_ = _dafny.Map({len(d_1949_v63_): p0})
        d_1950_v64_ = (d_1950_v64_).set((self).f11, (self).f11)
        d_1951_v65_: str
        d_1951_v65_ = _dafny.CodePoint('c')
        d_1952_v66_: _dafny.Seq
        d_1952_v66_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktkt"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktkt")))), d_1951_v65_), _dafny.SeqWithoutIsStrInference([d_1951_v65_ for d_1953_i6_ in range(default__.abs(201))])])
        d_1954_v67_: _dafny.Map
        d_1954_v67_ = _dafny.Map({(self).f9: (d_1952_v66_)[default__.safeIndex((self).f11, len(d_1952_v66_))]})
        d_1954_v67_ = (d_1954_v67_).set(default__.fm31(globalState), (self).f9)
        r3 = (self).fm5(p0, default__.safeModuloInt(p0, len(d_1874_v0_)), globalState)
        d_1955_v69_: bool
        d_1955_v69_ = True
        d_1956_v70_: _dafny.Map
        d_1956_v70_ = _dafny.Map({d_1950_v64_: (_dafny.MultiSet([d_1955_v69_, d_1955_v69_, d_1955_v69_])).cardinality})
        nw316_ = _dafny.Array(None, 11)
        def iife164_():
            coll78_ = _dafny.Map()
            compr_78_: int
            for compr_78_ in _dafny.IntegerRange(-288, -774):
                d_1957_v68_: int = compr_78_
                if ((-288) <= (d_1957_v68_)) and ((d_1957_v68_) < (-774)):
                    coll78_[default__.safeModuloInt(d_1957_v68_, 615)] = p0
            return _dafny.Map(coll78_)
        nw316_[int(0)] = (_dafny.Map({iife164_()
        : (self).f11})) == (d_1956_v70_)
        nw316_[int(1)] = d_1955_v69_
        nw316_[int(2)] = d_1955_v69_
        nw316_[int(3)] = d_1955_v69_
        nw316_[int(4)] = False
        nw316_[int(5)] = d_1955_v69_
        nw316_[int(6)] = not(d_1955_v69_)
        nw316_[int(7)] = d_1955_v69_
        nw316_[int(8)] = default__.fm38(p0, globalState)
        nw316_[int(9)] = (d_1955_v69_) and (d_1955_v69_)
        nw316_[int(10)] = d_1955_v69_
        d_1947_v62_ = nw316_
        d_1958_v71_: D3
        d_1958_v71_ = D3_DC7(d_1955_v69_)
        d_1959_v72_: D15
        d_1959_v72_ = D15_DC37((self).f11, d_1958_v71_, d_1955_v69_)
        source33_ = d_1959_v72_
        if source33_.is_DC36:
            d_1960___mcc_h2_ = source33_.cf48
            d_1961_cf48_ = d_1960___mcc_h2_
            d_1962_v73_: _dafny.MultiSet
            d_1962_v73_ = _dafny.MultiSet([d_1955_v69_, d_1955_v69_])
            d_1963_v74_: _dafny.MultiSet
            d_1963_v74_ = _dafny.MultiSet([(d_1962_v73_).intersection(d_1962_v73_)])
            d_1964_v75_: _dafny.Seq
            d_1964_v75_ = _dafny.SeqWithoutIsStrInference([d_1955_v69_, d_1955_v69_, d_1955_v69_])
            d_1963_v74_ = (_dafny.MultiSet([_dafny.MultiSet(d_1964_v75_), default__.fm34(len(d_1874_v0_), (0) - ((self).f11), p0, globalState), d_1962_v73_]) if d_1955_v69_ else (_dafny.MultiSet([d_1962_v73_, _dafny.MultiSet(d_1964_v75_), _dafny.MultiSet(d_1964_v75_)])).intersection(d_1963_v74_))
            r2 = d_1955_v69_
            r2 = False
            d_1965_v76_: _dafny.Set
            d_1965_v76_ = _dafny.Set({_dafny.MultiSet(d_1874_v0_)})
            d_1966_v77_: D10
            d_1966_v77_ = D10_DC23(d_1951_v65_, d_1965_v76_)
            d_1967_v78_: _dafny.MultiSet
            d_1967_v78_ = _dafny.MultiSet([d_1961_cf48_, (self).f11])
            pat_let_tv37_ = d_1965_v76_
            d_1968_v80_: _dafny.Array
            nw317_ = _dafny.Array(None, 21)
            nw317_[int(0)] = D10_DC23(d_1951_v65_, d_1965_v76_)
            nw317_[int(1)] = d_1966_v77_
            nw317_[int(2)] = d_1966_v77_
            nw317_[int(3)] = d_1966_v77_
            nw317_[int(4)] = d_1966_v77_
            nw317_[int(5)] = d_1966_v77_
            nw317_[int(6)] = d_1966_v77_
            nw317_[int(7)] = d_1966_v77_
            nw317_[int(8)] = d_1966_v77_
            def iife165_():
                coll79_ = _dafny.Set()
                compr_79_: _dafny.MultiSet
                for compr_79_ in (_dafny.Map({d_1967_v78_: d_1955_v69_})).keys.Elements:
                    d_1969_v79_: _dafny.MultiSet = compr_79_
                    if (d_1969_v79_) in (_dafny.Map({d_1967_v78_: d_1955_v69_})):
                        coll79_ = coll79_.union(_dafny.Set([d_1969_v79_]))
                return _dafny.Set(coll79_)
            nw317_[int(9)] = D10_DC23(d_1951_v65_, iife165_()
)
            nw317_[int(10)] = d_1966_v77_
            def iife166_(_pat_let43_0):
                def iife167_(d_1970_dt__update__tmp_h0_):
                    def iife168_(_pat_let44_0):
                        def iife169_(d_1971_dt__update_hcf27_h0_):
                            return D10_DC23((d_1970_dt__update__tmp_h0_).cf26, d_1971_dt__update_hcf27_h0_)
                        return iife169_(_pat_let44_0)
                    return iife168_(pat_let_tv37_)
                return iife167_(_pat_let43_0)
            nw317_[int(11)] = iife166_(d_1966_v77_)
            nw317_[int(12)] = d_1966_v77_
            nw317_[int(13)] = D10_DC23(d_1951_v65_, d_1965_v76_)
            nw317_[int(14)] = D10_DC23(d_1951_v65_, (D10_DC23(d_1951_v65_, d_1965_v76_)).cf27)
            nw317_[int(15)] = D10_DC23(d_1951_v65_, d_1965_v76_)
            nw317_[int(16)] = d_1966_v77_
            nw317_[int(17)] = d_1966_v77_
            nw317_[int(18)] = d_1966_v77_
            nw317_[int(19)] = d_1966_v77_
            nw317_[int(20)] = d_1966_v77_
            d_1968_v80_ = nw317_
            d_1972_v81_: _dafny.Array
            d_1972_v81_ = d_1968_v80_
            d_1973_v82_: _dafny.Array
            nw318_ = _dafny.Array(None, 1)
            nw318_[int(0)] = d_1972_v81_
            d_1973_v82_ = nw318_
            d_1974_v83_: _dafny.Map
            d_1974_v83_ = _dafny.Map({p0: d_1955_v69_})
            rhs254_ = d_1973_v82_
            rhs255_ = d_1974_v83_
            rhs256_ = d_1955_v69_
            d_1973_v82_ = rhs254_
            d_1974_v83_ = rhs255_
            r2 = rhs256_
        elif source33_.is_DC37:
            d_1975___mcc_h3_ = source33_.cf49
            d_1976___mcc_h4_ = source33_.cf50
            d_1977___mcc_h5_ = source33_.cf51
            d_1978_cf51_ = d_1977___mcc_h5_
            d_1979_cf50_ = d_1976___mcc_h4_
            d_1980_cf49_ = d_1975___mcc_h3_
            d_1981_v84_: C3
            nw319_ = C3()
            nw319_.ctor__((d_1980_cf49_) != (-515), p0, d_1980_cf49_, ((self.f12).set(d_1950_v64_, _dafny.Map({d_1980_cf49_: p0}))) | ((self.f12).set(d_1950_v64_, (d_1950_v64_).set(d_1980_cf49_, p0))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nit")), (self).f10)
            d_1981_v84_ = nw319_
            d_1981_v84_ = d_1981_v84_
            d_1951_v65_ = d_1951_v65_
            d_1982_v85_: _dafny.Array
            def lambda78_(d_1983_v64_):
                def lambda79_(d_1984_i7_):
                    return (d_1983_v64_) | (d_1983_v64_)

                return lambda79_

            init44_ = lambda78_(d_1950_v64_)
            nw320_ = _dafny.Array(None, 9)
            for i0_44_ in range(nw320_.length(0)):
                nw320_[i0_44_] = init44_(i0_44_)
            d_1982_v85_ = nw320_
            index266_ = default__.safeIndex(205, (d_1982_v85_).length(0))
            (d_1982_v85_)[index266_] = d_1950_v64_
            d_1985_v86_: _dafny.Map
            d_1985_v86_ = _dafny.Map({d_1947_v62_: d_1980_cf49_})
            index267_ = default__.safeIndex(205, (d_1982_v85_).length(0))
            rhs257_ = ((d_1985_v86_) | (d_1985_v86_)) == (_dafny.Map({d_1947_v62_: d_1980_cf49_}))
            rhs258_ = d_1950_v64_
            rhs259_ = d_1947_v62_
            lhs164_ = d_1982_v85_
            lhs165_ = default__.safeIndex(205, (d_1982_v85_).length(0))
            r2 = rhs257_
            lhs164_[lhs165_] = rhs258_
            d_1947_v62_ = rhs259_
            d_1986_v87_: _dafny.Array
            def lambda80_(d_1987_p0_):
                def lambda81_(d_1988_i8_):
                    return (d_1988_i8_) + (d_1987_p0_)

                return lambda81_

            init45_ = lambda80_(p0)
            nw321_ = _dafny.Array(None, 25)
            for i0_45_ in range(nw321_.length(0)):
                nw321_[i0_45_] = init45_(i0_45_)
            d_1986_v87_ = nw321_
            d_1986_v87_ = d_1986_v87_
        elif True:
            d_1989___mcc_h6_ = source33_.cf47
            d_1990_cf47_ = d_1989___mcc_h6_
            r2 = d_1955_v69_
            d_1991_v88_: D9
            d_1991_v88_ = D9_DC21(D3_DC6(_dafny.Set({p0})), d_1951_v65_, d_1955_v69_)
            d_1992_v89_: _dafny.Map
            d_1992_v89_ = _dafny.Map({969: d_1955_v69_})
            d_1993_v90_: _dafny.Map
            d_1993_v90_ = _dafny.Map({(d_1991_v88_).cf24: d_1992_v89_})
            r3 = len((d_1993_v90_) | (d_1993_v90_))
            (globalState).f3 = (self).f11
            d_1994_v91_: _dafny.Seq
            d_1994_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpaie"))
            d_1994_v91_ = (_dafny.SeqWithoutIsStrInference([d_1951_v65_ for d_1995_i9_ in range(default__.abs(696))])).set(default__.safeIndex(((self).f11) + (len(_dafny.Map({d_1955_v69_: p0}))), len(_dafny.SeqWithoutIsStrInference([d_1951_v65_ for d_1996_i9_ in range(default__.abs(696))]))), d_1951_v65_)
        d_1997_v92_: _dafny.Array
        def lambda82_(d_1998_i10_):
            return D1_DC2()

        init46_ = lambda82_
        nw322_ = _dafny.Array(None, 3)
        for i0_46_ in range(nw322_.length(0)):
            nw322_[i0_46_] = init46_(i0_46_)
        d_1997_v92_ = nw322_
        d_1999_v93_: _dafny.Array
        nw323_ = _dafny.Array(None, 24)
        nw323_[int(0)] = d_1997_v92_
        nw323_[int(1)] = d_1997_v92_
        nw323_[int(2)] = d_1997_v92_
        nw323_[int(3)] = d_1997_v92_
        nw323_[int(4)] = d_1997_v92_
        nw323_[int(5)] = d_1997_v92_
        nw323_[int(6)] = d_1997_v92_
        nw323_[int(7)] = d_1997_v92_
        nw323_[int(8)] = d_1997_v92_
        nw323_[int(9)] = d_1997_v92_
        nw323_[int(10)] = d_1997_v92_
        nw323_[int(11)] = d_1997_v92_
        nw323_[int(12)] = d_1997_v92_
        nw323_[int(13)] = d_1997_v92_
        nw323_[int(14)] = d_1997_v92_
        nw323_[int(15)] = d_1997_v92_
        nw323_[int(16)] = d_1997_v92_
        nw323_[int(17)] = d_1997_v92_
        nw323_[int(18)] = d_1997_v92_
        nw323_[int(19)] = d_1997_v92_
        nw323_[int(20)] = d_1997_v92_
        nw323_[int(21)] = d_1997_v92_
        nw323_[int(22)] = d_1997_v92_
        nw323_[int(23)] = d_1997_v92_
        d_1999_v93_ = nw323_
        d_2000_v94_: C5
        nw324_ = C5()
        nw324_.ctor__(699, d_1999_v93_)
        d_2000_v94_ = nw324_
        d_2001_v95_: _dafny.MultiSet
        d_2001_v95_ = _dafny.MultiSet([d_2000_v94_])
        r0 = default__.safeModuloInt((self).f11, (d_2001_v95_).cardinality)
        d_2002_v96_: _dafny.Map
        d_2002_v96_ = _dafny.Map({d_1955_v69_: (self).f9})
        r1 = (d_2002_v96_).set(((d_2000_v94_).f17) <= ((self).f11), (self).f9)
        r2 = d_1955_v69_
        r3 = 663
        return r0, r1, r2, r3

    def m3(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_2003_v0_: bool
        d_2003_v0_ = True
        d_2004_v1_: _dafny.Map
        d_2004_v1_ = _dafny.Map({d_2003_v0_: not(d_2003_v0_)})
        d_2005_v2_: C0
        nw325_ = C0()
        nw325_.ctor__(False, (True) in (d_2004_v1_), p1, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p2 for d_2006_i1_ in range(default__.abs(480))]) for d_2007_i0_ in range(default__.abs(96))]))
        d_2005_v2_ = nw325_
        d_2008_v3_: _dafny.Array
        nw326_ = _dafny.Array(int(0), 19)
        d_2008_v3_ = nw326_
        rhs260_ = len(default__.fm31(globalState))
        rhs261_ = d_2005_v2_
        rhs262_ = d_2008_v3_
        lhs166_ = globalState
        lhs166_.f3 = rhs260_
        d_2005_v2_ = rhs261_
        r1 = rhs262_
        if (d_2005_v2_).f25:
            d_2009_v4_: _dafny.Map
            d_2009_v4_ = _dafny.Map({(d_2005_v2_).f25: p2})
            d_2010_v5_: _dafny.MultiSet
            d_2010_v5_ = _dafny.MultiSet([len(d_2009_v4_), p2, (self).f11, (self).f11])
            (globalState).f3 = default__.safeModuloInt((d_2010_v5_).cardinality, (self).f11)
            d_2011_v6_: _dafny.Seq
            d_2011_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqdygrsf"))
            rhs263_ = (self).f11
            rhs264_ = d_2011_v6_
            lhs167_ = globalState
            lhs167_.f3 = rhs263_
            d_2011_v6_ = rhs264_
            d_2012_v7_: _dafny.Array
            nw327_ = _dafny.Array(D11.default()(), 4)
            d_2012_v7_ = nw327_
            d_2013_v8_: _dafny.Map
            d_2013_v8_ = _dafny.Map({(d_2005_v2_).f26: d_2012_v7_})
            (globalState).f3 = (len(d_2013_v8_) if not((d_2005_v2_).f26) else p2)
            d_2014_v9_: _dafny.Map
            d_2014_v9_ = _dafny.Map({(self).f11: ((self).f11) + (default__.fm0(p0, p2, globalState))})
            d_2014_v9_ = (d_2014_v9_).set(p0, (self).f11)
            d_2015_v10_: _dafny.Seq
            d_2015_v10_ = _dafny.SeqWithoutIsStrInference([(d_2005_v2_).f25])
            d_2016_v11_: _dafny.Seq
            d_2016_v11_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2017_i2_ in range(default__.abs(-914))]), d_2011_v6_])
            d_2003_v0_ = not ((d_2005_v2_).f26) or ((d_2015_v10_)[default__.safeIndex(len(d_2016_v11_), len(d_2015_v10_))])
        elif True:
            d_2018_v12_: _dafny.Seq
            d_2018_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tn"))
            d_2018_v12_ = d_2018_v12_
            d_2003_v0_ = not((p0) == ((self).f11))
            d_2019_v13_: _dafny.Map
            d_2019_v13_ = _dafny.Map({p0: (d_2005_v2_).f26})
            if ((d_2019_v13_)[p0] if (p0) in (d_2019_v13_) else not((d_2005_v2_).f26)):
                index268_ = default__.safeIndex(42, (d_2008_v3_).length(0))
                (d_2008_v3_)[index268_] = len(_dafny.SeqWithoutIsStrInference([(d_2005_v2_).f26]))
                d_2020_v14_: _dafny.Map
                d_2020_v14_ = _dafny.Map({d_2003_v0_: len(d_2004_v1_)})
                d_2021_v18_: _dafny.MultiSet
                def iife170_():
                    coll80_ = _dafny.Set()
                    compr_80_: int
                    for compr_80_ in _dafny.IntegerRange(371, -450):
                        d_2022_v16_: int = compr_80_
                        if ((371) <= (d_2022_v16_)) and ((d_2022_v16_) < (-450)):
                            coll80_ = coll80_.union(_dafny.Set([(d_2022_v16_) - (p0)]))
                    return _dafny.Set(coll80_)
                def iife171_():
                    coll81_ = _dafny.Set()
                    compr_81_: int
                    for compr_81_ in (_dafny.SeqWithoutIsStrInference([p2, p2])).Elements:
                        d_2023_v17_: int = compr_81_
                        if (d_2023_v17_) in (_dafny.SeqWithoutIsStrInference([p2, p2])):
                            coll81_ = coll81_.union(_dafny.Set([default__.safeModuloInt(d_2023_v17_, 451)]))
                    return _dafny.Set(coll81_)
                d_2021_v18_ = _dafny.MultiSet([iife170_()
                , iife171_()
                ])
                d_2024_v19_: _dafny.Map
                d_2024_v19_ = _dafny.Map({len((self).f9): d_2020_v14_})
                d_2025_v20_: _dafny.Array
                nw328_ = _dafny.Array(False, 25)
                d_2025_v20_ = nw328_
                d_2026_v21_: D2
                d_2026_v21_ = D2_DC5((d_2005_v2_).f25, d_2025_v20_, p2, len(_dafny.SeqWithoutIsStrInference([p2])))
                d_2027_v22_: _dafny.Seq
                d_2027_v22_ = _dafny.SeqWithoutIsStrInference([d_2003_v0_, d_2003_v0_])
                d_2028_v23_: _dafny.Array
                nw329_ = _dafny.Array(None, 27)
                nw329_[int(0)] = (_dafny.Map({(d_2005_v2_).f25: p2})) | (d_2020_v14_)
                nw329_[int(1)] = d_2020_v14_
                def iife172_():
                    coll82_ = _dafny.Map()
                    compr_82_: _dafny.Set
                    for compr_82_ in (d_2021_v18_).Elements:
                        d_2029_v15_: _dafny.Set = compr_82_
                        if (d_2029_v15_) in (d_2021_v18_):
                            coll82_[d_2029_v15_] = p0
                    return _dafny.Map(coll82_)
                nw329_[int(2)] = ((d_2020_v14_).set((d_2005_v2_).f26, len(iife172_()
                ))) | (default__.fm42(d_2020_v14_, False, (self).f11, (d_2005_v2_).f26, globalState))
                nw329_[int(3)] = d_2020_v14_
                nw329_[int(4)] = d_2020_v14_
                nw329_[int(5)] = ((d_2024_v19_)[p2] if (p2) in (d_2024_v19_) else d_2020_v14_)
                nw329_[int(6)] = d_2020_v14_
                nw329_[int(7)] = (d_2020_v14_).set((d_2005_v2_).f25, (0) - ((0) - (p0)))
                nw329_[int(8)] = d_2020_v14_
                nw329_[int(9)] = d_2020_v14_
                nw329_[int(10)] = _dafny.Map({(d_2026_v21_).cf4: (self).f11})
                nw329_[int(11)] = d_2020_v14_
                nw329_[int(12)] = d_2020_v14_
                nw329_[int(13)] = d_2020_v14_
                nw329_[int(14)] = default__.fm42(d_2020_v14_, (d_2005_v2_).f26, (self).f11, not((d_2005_v2_).f25), globalState)
                nw329_[int(15)] = (d_2020_v14_).set((d_2005_v2_).f25, p2)
                nw329_[int(16)] = (d_2020_v14_).set(d_2003_v0_, -297)
                nw329_[int(17)] = (_dafny.Map({(d_2027_v22_)[default__.safeIndex(p0, len(d_2027_v22_))]: (self).fm5((self).f11, (self).f11, globalState)})) | (d_2020_v14_)
                nw329_[int(18)] = ((self).fm20((d_2005_v2_).f25, p0, (self).f11, globalState)) | (default__.fm42(d_2020_v14_, True, p0, d_2003_v0_, globalState))
                nw329_[int(19)] = d_2020_v14_
                nw329_[int(20)] = d_2020_v14_
                nw329_[int(21)] = d_2020_v14_
                nw329_[int(22)] = (d_2020_v14_) | (d_2020_v14_)
                nw329_[int(23)] = _dafny.Map({False: (self).f11})
                nw329_[int(24)] = (_dafny.Map({(d_2005_v2_).f25: 780})) | (d_2020_v14_)
                nw329_[int(25)] = d_2020_v14_
                nw329_[int(26)] = (d_2020_v14_) | (d_2020_v14_)
                d_2028_v23_ = nw329_
                index269_ = default__.safeIndex(333, (d_2028_v23_).length(0))
                (d_2028_v23_)[index269_] = ((d_2020_v14_).set((d_2005_v2_).f26, p2)) | (d_2020_v14_)
                index270_ = default__.safeIndex(42, (d_2008_v3_).length(0))
                index271_ = default__.safeIndex(333, (d_2028_v23_).length(0))
                rhs265_ = not(((d_2005_v2_).f25) and ((d_2005_v2_).f25))
                rhs266_ = 576
                rhs267_ = (d_2020_v14_) | (d_2020_v14_)
                lhs168_ = d_2008_v3_
                lhs169_ = default__.safeIndex(42, (d_2008_v3_).length(0))
                lhs170_ = d_2028_v23_
                lhs171_ = default__.safeIndex(333, (d_2028_v23_).length(0))
                d_2003_v0_ = rhs265_
                lhs168_[lhs169_] = rhs266_
                lhs170_[lhs171_] = rhs267_
                (globalState).f3 = p2
                (globalState).f3 = (d_2008_v3_)[default__.safeIndex(42, (d_2008_v3_).length(0))]
                d_2003_v0_ = (p2) <= ((self).f11)
                d_2003_v0_ = (d_2005_v2_).f25
            elif True:
                d_2003_v0_ = (d_2005_v2_).f26
                d_2030_v24_: _dafny.Array
                nw330_ = _dafny.Array(False, 12)
                d_2030_v24_ = nw330_
                index272_ = default__.safeIndex(569, (d_2030_v24_).length(0))
                (d_2030_v24_)[index272_] = (d_2005_v2_).f26
                index273_ = default__.safeIndex(569, (d_2030_v24_).length(0))
                (d_2030_v24_)[index273_] = ((self).f11) == ((112) - ((self).f11))
                d_2031_v25_: str
                d_2031_v25_ = _dafny.CodePoint('u')
                d_2032_v26_: _dafny.MultiSet
                d_2032_v26_ = _dafny.MultiSet([p0])
                d_2033_v27_: _dafny.MultiSet
                d_2033_v27_ = _dafny.MultiSet([d_2032_v26_])
                d_2031_v25_ = (d_2031_v25_ if (d_2033_v27_).isdisjoint(_dafny.MultiSet([d_2032_v26_])) else d_2031_v25_)
                index274_ = default__.safeIndex(569, (d_2030_v24_).length(0))
                rhs268_ = (d_2005_v2_).f26
                rhs269_ = (d_2005_v2_).f26
                rhs270_ = (_dafny.CodePoint('r')) in ((self).f9)
                lhs172_ = d_2030_v24_
                lhs173_ = default__.safeIndex(569, (d_2030_v24_).length(0))
                lhs172_[lhs173_] = rhs268_
                d_2003_v0_ = rhs269_
                d_2003_v0_ = rhs270_
                index275_ = default__.safeIndex(569, (d_2030_v24_).length(0))
                rhs271_ = (d_2005_v2_).f26
                rhs272_ = (d_2005_v2_).f26
                lhs174_ = d_2030_v24_
                lhs175_ = default__.safeIndex(569, (d_2030_v24_).length(0))
                d_2003_v0_ = rhs271_
                lhs174_[lhs175_] = rhs272_
            d_2034_v28_: _dafny.Seq
            d_2034_v28_ = _dafny.SeqWithoutIsStrInference([(d_2005_v2_).f25])
            d_2034_v28_ = (((d_2034_v28_) + (_dafny.SeqWithoutIsStrInference([d_2003_v0_]))) + ((_dafny.SeqWithoutIsStrInference([(d_2005_v2_).f26])).set(default__.safeIndex((self).f11, len(_dafny.SeqWithoutIsStrInference([(d_2005_v2_).f26]))), d_2003_v0_))).set(default__.safeIndex(0, len(((d_2034_v28_) + (_dafny.SeqWithoutIsStrInference([d_2003_v0_]))) + ((_dafny.SeqWithoutIsStrInference([(d_2005_v2_).f26])).set(default__.safeIndex((self).f11, len(_dafny.SeqWithoutIsStrInference([(d_2005_v2_).f26]))), d_2003_v0_)))), (d_2005_v2_).f26)
            d_2003_v0_ = d_2003_v0_
        d_2035_v29_: _dafny.Array
        nw331_ = _dafny.Array(_dafny.CodePoint('D'), 27)
        d_2035_v29_ = nw331_
        d_2036_v30_: D6
        d_2036_v30_ = D6_DC13(d_2035_v29_)
        d_2036_v30_ = d_2036_v30_
        d_2037_v31_: D1
        d_2037_v31_ = D1_DC2()
        d_2038_v32_: _dafny.Seq
        d_2038_v32_ = _dafny.SeqWithoutIsStrInference([p2])
        d_2039_v33_: _dafny.Map
        d_2039_v33_ = _dafny.Map({d_2038_v32_: p2})
        rhs273_ = not(d_2003_v0_)
        rhs274_ = d_2037_v31_
        rhs275_ = ((d_2039_v33_)[(_dafny.SeqWithoutIsStrInference([p0 for d_2040_i3_ in range(default__.abs(746))])) + (d_2038_v32_)] if ((_dafny.SeqWithoutIsStrInference([p0 for d_2041_i3_ in range(default__.abs(746))])) + (d_2038_v32_)) in (d_2039_v33_) else p0)
        lhs176_ = globalState
        d_2003_v0_ = rhs273_
        d_2037_v31_ = rhs274_
        lhs176_.f3 = rhs275_
        d_2042_v34_: _dafny.Set
        d_2042_v34_ = _dafny.Set({518, p0, len(p1), (d_2038_v32_)[default__.safeIndex(p2, len(d_2038_v32_))]})
        d_2043_i4_: int
        d_2043_i4_ = 0
        with _dafny.label("9"):
            while ((p2) + (p2)) < ((len(d_2042_v34_)) * ((self).fm5(p0, (self).f11, globalState))):
                with _dafny.c_label("9"):
                    if (d_2043_i4_) >= (100):
                        raise _dafny.Break("9")
                    d_2043_i4_ = (d_2043_i4_) + (1)
                    if not(not((d_2005_v2_).f25)):
                        d_2044_v35_: bool
                        d_2045_v36_: bool
                        d_2046_v37_: _dafny.Array
                        d_2047_v38_: int
                        out60_: bool
                        out61_: bool
                        out62_: _dafny.Array
                        out63_: int
                        out60_, out61_, out62_, out63_ = (d_2005_v2_).m1(p0, (p0) - ((self).f11), default__.safeDivisionInt((0) - (p2), p2), globalState)
                        d_2044_v35_ = out60_
                        d_2045_v36_ = out61_
                        d_2046_v37_ = out62_
                        d_2047_v38_ = out63_
                        d_2048_v39_: C3
                        nw332_ = C3()
                        nw332_.ctor__((d_2005_v2_).f25, d_2047_v38_, p0, (self.f12) | (self.f12), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2049_i5_ in range(default__.abs(230))]) if True else (self).f9), ((self).f10 if (d_2005_v2_).f25 else (self).f10))
                        d_2048_v39_ = nw332_
                        d_2050_v40_: _dafny.Map
                        d_2050_v40_ = _dafny.Map({(self).f11: (self).f11})
                        d_2051_v41_: C1
                        nw333_ = C1()
                        nw333_.ctor__((d_2048_v39_).fm29((d_2048_v39_).fm29(len(d_2050_v40_), (d_2005_v2_).f25, True, globalState), d_2003_v0_, not(d_2044_v35_), globalState), -996, (0) - ((self).f11), self.f12, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2052_i6_ in range(default__.abs(304))]), (self).f10)
                        d_2051_v41_ = nw333_
                        d_2053_v42_: _dafny.MultiSet
                        d_2053_v42_ = _dafny.MultiSet([d_2008_v3_, d_2046_v37_])
                        d_2054_v43_: D15
                        d_2054_v43_ = D15_DC35(d_2053_v42_)
                        d_2055_v44_: _dafny.Map
                        d_2055_v44_ = _dafny.Map({d_2054_v43_: d_2051_v41_.f23})
                        d_2056_v45_: _dafny.Seq
                        d_2056_v45_ = _dafny.SeqWithoutIsStrInference([d_2053_v42_])
                        d_2057_v46_: _dafny.Seq
                        d_2057_v46_ = _dafny.SeqWithoutIsStrInference([(d_2056_v45_)[default__.safeIndex(d_2047_v38_, len(d_2056_v45_))], _dafny.MultiSet([d_2046_v37_, d_2046_v37_])])
                        pat_let_tv38_ = d_2056_v45_
                        pat_let_tv39_ = d_2048_v39_
                        pat_let_tv40_ = d_2056_v45_
                        def iife173_(_pat_let45_0):
                            def iife174_(d_2058_dt__update__tmp_h0_):
                                def iife175_(_pat_let46_0):
                                    def iife176_(d_2059_dt__update_hcf47_h0_):
                                        return D15_DC35(d_2059_dt__update_hcf47_h0_)
                                    return iife176_(_pat_let46_0)
                                return iife175_((pat_let_tv38_)[default__.safeIndex((pat_let_tv39_).f22, len(pat_let_tv40_))])
                            return iife174_(_pat_let45_0)
                        d_2055_v44_ = (d_2055_v44_).set(iife173_(d_2054_v43_), d_2047_v38_)
                        d_2060_v47_: _dafny.Seq
                        d_2060_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apogbv"))
                        d_2061_v48_: _dafny.Seq
                        d_2061_v48_ = _dafny.SeqWithoutIsStrInference([(d_2045_v36_) and (d_2044_v35_)])
                        d_2062_v49_: _dafny.Array
                        nw334_ = _dafny.Array(None, 8)
                        nw334_[int(0)] = d_2038_v32_
                        nw334_[int(1)] = d_2038_v32_
                        nw334_[int(2)] = d_2038_v32_
                        nw334_[int(3)] = d_2038_v32_
                        nw334_[int(4)] = d_2038_v32_
                        nw334_[int(5)] = (d_2038_v32_ if (d_2005_v2_).f25 else _dafny.SeqWithoutIsStrInference([72, p0, (d_2051_v41_).f24, d_2047_v38_]))
                        nw334_[int(6)] = _dafny.SeqWithoutIsStrInference([(0) - ((d_2048_v39_).f22) for d_2063_i7_ in range(default__.abs(370))])
                        nw334_[int(7)] = (d_2038_v32_) + (d_2038_v32_)
                        d_2062_v49_ = nw334_
                        index276_ = default__.safeIndex(624, (d_2062_v49_).length(0))
                        (d_2062_v49_)[index276_] = ((self).f10)[default__.safeIndex((0) - (p2), len((self).f10))]
                        index277_ = default__.safeIndex(624, (d_2062_v49_).length(0))
                        rhs276_ = d_2047_v38_
                        rhs277_ = p1
                        rhs278_ = (0) - ((self).f11)
                        rhs279_ = _dafny.SeqWithoutIsStrInference([(d_2061_v48_)[default__.safeIndex(d_2051_v41_.f23, len(d_2061_v48_))]])
                        rhs280_ = ((((d_2038_v32_).set(default__.safeIndex((self).fm5(d_2051_v41_.f23, p2, globalState), len(d_2038_v32_)), 358)) + (d_2038_v32_)) + ((d_2038_v32_) + (d_2038_v32_))).set(default__.safeIndex(p2, len((((d_2038_v32_).set(default__.safeIndex((self).fm5(d_2051_v41_.f23, p2, globalState), len(d_2038_v32_)), 358)) + (d_2038_v32_)) + ((d_2038_v32_) + (d_2038_v32_)))), (d_2005_v2_).fm5(-8, (d_2048_v39_).f22, globalState))
                        lhs177_ = globalState
                        lhs178_ = d_2062_v49_
                        lhs179_ = default__.safeIndex(624, (d_2062_v49_).length(0))
                        lhs177_.f3 = rhs276_
                        d_2060_v47_ = rhs277_
                        d_2047_v38_ = rhs278_
                        d_2061_v48_ = rhs279_
                        lhs178_[lhs179_] = rhs280_
                    elif True:
                        d_2064_v50_: _dafny.Seq
                        d_2064_v50_ = _dafny.SeqWithoutIsStrInference([(d_2005_v2_).f25, (d_2005_v2_).f26])
                        d_2065_v51_: _dafny.Map
                        d_2065_v51_ = _dafny.Map({default__.fm0(p2, len((d_2064_v50_).set(default__.safeIndex(p2, len(d_2064_v50_)), (d_2005_v2_).f26)), globalState): 660})
                        d_2066_v52_: _dafny.Map
                        d_2066_v52_ = _dafny.Map({(d_2005_v2_).f25: (self).f9})
                        rhs281_ = (((d_2065_v51_)[p0] if (p0) in (d_2065_v51_) else p0)) < ((self).f11)
                        rhs282_ = default__.safeDivisionInt(len(d_2004_v1_), ((self).f11 if (d_2005_v2_).f26 else len(((d_2066_v52_)[(d_2005_v2_).f26] if ((d_2005_v2_).f26) in (d_2066_v52_) else p1))))
                        rhs283_ = (p2) - ((self).f11)
                        lhs180_ = globalState
                        lhs181_ = globalState
                        d_2003_v0_ = rhs281_
                        lhs180_.f3 = rhs282_
                        lhs181_.f3 = rhs283_
                        d_2067_v53_: _dafny.Seq
                        d_2067_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrevsuo"))
                        d_2067_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgtr"))
                        d_2003_v0_ = not(((d_2004_v1_) == (_dafny.Map({(d_2005_v2_).f25: (d_2005_v2_).f26}))) not in (default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uaqikok")), p2, globalState)))
                        rhs284_ = (d_2042_v34_).intersection((d_2042_v34_ if (d_2005_v2_).f26 else _dafny.Set({(self).f11})))
                        rhs285_ = d_2003_v0_
                        d_2042_v34_ = rhs284_
                        d_2003_v0_ = rhs285_
                        d_2068_v54_: str
                        d_2068_v54_ = _dafny.CodePoint('j')
                        d_2069_v55_: _dafny.MultiSet
                        d_2069_v55_ = _dafny.MultiSet([-269])
                        d_2070_v56_: _dafny.Set
                        d_2070_v56_ = _dafny.Set({d_2069_v55_})
                        d_2071_v57_: D10
                        d_2071_v57_ = D10_DC23(d_2068_v54_, d_2070_v56_)
                        d_2072_v58_: _dafny.Map
                        d_2072_v58_ = _dafny.Map({d_2071_v57_: (d_2005_v2_).f25})
                        d_2072_v58_ = (d_2072_v58_).set(d_2071_v57_, not((d_2005_v2_).f26))
                    d_2073_v59_: _dafny.Array
                    nw335_ = _dafny.Array(_dafny.Array(None, 0), 28)
                    d_2073_v59_ = nw335_
                    d_2074_v60_: _dafny.Array
                    nw336_ = _dafny.Array(False, 10)
                    d_2074_v60_ = nw336_
                    index278_ = default__.safeIndex(105, (d_2073_v59_).length(0))
                    (d_2073_v59_)[index278_] = d_2074_v60_
                    index279_ = default__.safeIndex(105, (d_2073_v59_).length(0))
                    (d_2073_v59_)[index279_] = d_2074_v60_
                    d_2075_v61_: _dafny.Map
                    d_2075_v61_ = _dafny.Map({976: D4_DC11((d_2005_v2_).f25, len(d_2038_v32_))})
                    d_2076_v62_: _dafny.Map
                    d_2076_v62_ = _dafny.Map({(d_2005_v2_).f25: d_2075_v61_})
                    d_2076_v62_ = (d_2076_v62_) | (d_2076_v62_)
                    d_2077_v63_: _dafny.MultiSet
                    d_2077_v63_ = _dafny.MultiSet([d_2003_v0_, (d_2005_v2_).f26, (d_2005_v2_).f25])
                    d_2003_v0_ = ((d_2077_v63_).issubset(d_2077_v63_)) and (not(d_2003_v0_))
                    pass
            pass
        d_2078_v64_: _dafny.Map
        d_2078_v64_ = _dafny.Map({(self).f11: d_2003_v0_})
        d_2079_v65_: D4
        d_2079_v65_ = D4_DC8((d_2078_v64_) | (d_2078_v64_))
        d_2079_v65_ = (D4_DC8(d_2078_v64_) if not((d_2005_v2_).f26) else d_2079_v65_)
        d_2080_v66_: _dafny.Array
        nw337_ = _dafny.Array(_dafny.Array(None, 0), 14)
        d_2080_v66_ = nw337_
        r0 = d_2080_v66_
        r1 = d_2008_v3_
        return r0, r1

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_2081_v0_: bool
        d_2081_v0_ = True
        d_2082_v1_: _dafny.Map
        d_2082_v1_ = _dafny.Map({(self).f11: d_2081_v0_})
        d_2082_v1_ = (d_2082_v1_).set((self).f11, d_2081_v0_)
        d_2083_v2_: _dafny.Array
        nw338_ = _dafny.Array(_dafny.Array(None, 0), 25)
        d_2083_v2_ = nw338_
        d_2084_v3_: C5
        nw339_ = C5()
        nw339_.ctor__((0) - (p0), d_2083_v2_)
        d_2084_v3_ = nw339_
        d_2085_v4_: _dafny.Map
        d_2085_v4_ = _dafny.Map({p0: d_2084_v3_})
        r3 = len(d_2085_v4_)
        d_2086_v5_: _dafny.Seq
        d_2086_v5_ = _dafny.SeqWithoutIsStrInference([324])
        r0 = default__.fm38(len((d_2086_v5_) + (d_2086_v5_)), globalState)
        d_2087_i0_: int
        d_2087_i0_ = 0
        with _dafny.label("10"):
            while d_2081_v0_:
                with _dafny.c_label("10"):
                    if (d_2087_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_2087_i0_ = (d_2087_i0_) + (1)
                    d_2088_v6_: _dafny.Array
                    def lambda83_(d_2089_v0_):
                        def lambda84_(d_2090_i1_):
                            return d_2089_v0_

                        return lambda84_

                    init47_ = lambda83_(d_2081_v0_)
                    nw340_ = _dafny.Array(None, 16)
                    for i0_47_ in range(nw340_.length(0)):
                        nw340_[i0_47_] = init47_(i0_47_)
                    d_2088_v6_ = nw340_
                    d_2088_v6_ = d_2088_v6_
                    d_2088_v6_ = (d_2088_v6_ if d_2081_v0_ else d_2088_v6_)
                    d_2091_v7_: _dafny.Array
                    nw341_ = _dafny.Array(int(0), 8)
                    d_2091_v7_ = nw341_
                    d_2092_v8_: _dafny.Set
                    d_2092_v8_ = _dafny.Set({d_2091_v7_, d_2091_v7_})
                    if ((d_2092_v8_ if d_2081_v0_ else d_2092_v8_)).ispropersubset(d_2092_v8_):
                        d_2093_v9_: D14
                        d_2093_v9_ = D14_DC34((d_2086_v5_)[default__.safeIndex((self).f11, len(d_2086_v5_))], len((self).f9), 300, d_2081_v0_)
                        pat_let_tv41_ = d_2081_v0_
                        d_2094_v10_: _dafny.Seq
                        def iife177_(_pat_let47_0):
                            def iife178_(d_2095_dt__update__tmp_h0_):
                                def iife179_(_pat_let48_0):
                                    def iife180_(d_2096_dt__update_hcf46_h0_):
                                        return D14_DC34((d_2095_dt__update__tmp_h0_).cf43, (d_2095_dt__update__tmp_h0_).cf44, (d_2095_dt__update__tmp_h0_).cf45, d_2096_dt__update_hcf46_h0_)
                                    return iife180_(_pat_let48_0)
                                return iife179_(pat_let_tv41_)
                            return iife178_(_pat_let47_0)
                        d_2094_v10_ = _dafny.SeqWithoutIsStrInference([(iife177_(d_2093_v9_)).cf46])
                        r3 = (0) - (default__.safeModuloInt(len(d_2094_v10_), (p1 if d_2081_v0_ else (d_2084_v3_).f17)))
                        index280_ = default__.safeIndex(682, (d_2091_v7_).length(0))
                        (d_2091_v7_)[index280_] = p1
                        d_2097_v11_: _dafny.Seq
                        d_2097_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
                        d_2098_v12_: _dafny.MultiSet
                        d_2098_v12_ = _dafny.MultiSet([d_2088_v6_, d_2088_v6_])
                        d_2099_v13_: _dafny.MultiSet
                        d_2099_v13_ = _dafny.MultiSet([_dafny.Map({(d_2084_v3_).f17: d_2081_v0_}), d_2082_v1_, _dafny.Map({(self).f11: d_2081_v0_})])
                        index281_ = default__.safeIndex(682, (d_2091_v7_).length(0))
                        rhs286_ = (0) - (((0) - ((d_2084_v3_).f17)) - (p2))
                        rhs287_ = ((self).f9) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lel"))).set(default__.safeIndex(((d_2099_v13_)[d_2082_v1_] if (d_2082_v1_) in (d_2099_v13_) else p0), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lel")))), _dafny.CodePoint('w')))
                        rhs288_ = d_2098_v12_
                        lhs182_ = d_2091_v7_
                        lhs183_ = default__.safeIndex(682, (d_2091_v7_).length(0))
                        lhs182_[lhs183_] = rhs286_
                        d_2097_v11_ = rhs287_
                        d_2098_v12_ = rhs288_
                        d_2081_v0_ = d_2081_v0_
                        d_2088_v6_ = (D2_DC5(d_2081_v0_, d_2088_v6_, (d_2091_v7_)[default__.safeIndex(682, (d_2091_v7_).length(0))], (d_2091_v7_)[default__.safeIndex(682, (d_2091_v7_).length(0))])).cf5
                        d_2097_v11_ = (self).f9
                    elif True:
                        d_2100_v14_: _dafny.MultiSet
                        d_2100_v14_ = _dafny.MultiSet([(d_2084_v3_).f17])
                        d_2101_v15_: D7
                        d_2101_v15_ = D7_DC15(d_2100_v14_)
                        d_2102_v16_: _dafny.Set
                        d_2102_v16_ = _dafny.Set({default__.fm39(d_2101_v15_, d_2081_v0_, globalState)})
                        rhs289_ = d_2088_v6_
                        rhs290_ = p1
                        rhs291_ = d_2081_v0_
                        rhs292_ = (d_2102_v16_).isdisjoint(d_2102_v16_)
                        lhs184_ = globalState
                        d_2088_v6_ = rhs289_
                        lhs184_.f3 = rhs290_
                        r1 = rhs291_
                        r0 = rhs292_
                        r3 = (0) - (-726)
                        d_2103_v17_: _dafny.Map
                        d_2103_v17_ = _dafny.Map({(self).f9: (d_2084_v3_).f17})
                        d_2081_v0_ = ((_dafny.Map({(self).f9: p0})) | (d_2103_v17_)) != (d_2103_v17_)
                        d_2104_v18_: str
                        d_2104_v18_ = _dafny.CodePoint('p')
                        d_2104_v18_ = d_2104_v18_
                        r1 = True
                    d_2105_v19_: _dafny.Array
                    def lambda85_(d_2106_v3_):
                        def lambda86_(d_2107_i2_):
                            return _dafny.MultiSet([(d_2106_v3_).f17, (d_2106_v3_).f17])

                        return lambda86_

                    init48_ = lambda85_(d_2084_v3_)
                    nw342_ = _dafny.Array(None, 26)
                    for i0_48_ in range(nw342_.length(0)):
                        nw342_[i0_48_] = init48_(i0_48_)
                    d_2105_v19_ = nw342_
                    d_2105_v19_ = (d_2105_v19_ if d_2081_v0_ else d_2105_v19_)
                    pass
            pass
        d_2108_i3_: int
        d_2108_i3_ = 0
        with _dafny.label("11"):
            while d_2081_v0_:
                with _dafny.c_label("11"):
                    if (d_2108_i3_) >= (100):
                        raise _dafny.Break("11")
                    d_2108_i3_ = (d_2108_i3_) + (1)
                    d_2082_v1_ = d_2082_v1_
                    (globalState).f3 = ((0) - ((0) - ((p2) * ((d_2084_v3_).f17)))) + ((d_2084_v3_).f17)
                    d_2109_v20_: _dafny.Array
                    nw343_ = _dafny.Array(_dafny.Array(None, 0), 18)
                    d_2109_v20_ = nw343_
                    d_2109_v20_ = (d_2109_v20_ if False else d_2109_v20_)
                    d_2110_v21_: _dafny.Array
                    def lambda87_(d_2111_p2_):
                        def lambda88_(d_2112_i4_):
                            return default__.safeDivisionInt(d_2112_i4_, d_2111_p2_)

                        return lambda88_

                    init49_ = lambda87_(p2)
                    nw344_ = _dafny.Array(None, 22)
                    for i0_49_ in range(nw344_.length(0)):
                        nw344_[i0_49_] = init49_(i0_49_)
                    d_2110_v21_ = nw344_
                    r2 = d_2110_v21_
                    pass
            pass
        if not((len((d_2082_v1_).set(p0, d_2081_v0_))) == ((p1) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: d_2081_v0_})) for d_2113_i5_ in range(default__.abs(699))]))))):
            d_2114_v22_: _dafny.Array
            nw345_ = _dafny.Array(_dafny.Seq({}), 9)
            d_2114_v22_ = nw345_
            d_2115_v23_: _dafny.Seq
            d_2115_v23_ = _dafny.SeqWithoutIsStrInference([not(d_2081_v0_), True])
            index282_ = default__.safeIndex(45, (d_2114_v22_).length(0))
            (d_2114_v22_)[index282_] = (d_2115_v23_) + (d_2115_v23_)
            d_2116_v24_: _dafny.Seq
            d_2116_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
            index283_ = default__.safeIndex(45, (d_2114_v22_).length(0))
            def iife181_():
                coll83_ = _dafny.Set()
                compr_83_: int
                for compr_83_ in _dafny.IntegerRange(407, 43):
                    d_2117_v25_: int = compr_83_
                    if ((407) <= (d_2117_v25_)) and ((d_2117_v25_) < (43)):
                        coll83_ = coll83_.union(_dafny.Set([default__.safeDivisionInt(d_2117_v25_, (d_2086_v5_)[default__.safeIndex(p2, len(d_2086_v5_))])]))
                return _dafny.Set(coll83_)
            rhs293_ = (d_2115_v23_).set(default__.safeIndex((d_2084_v3_).f17, len(d_2115_v23_)), d_2081_v0_)
            rhs294_ = d_2086_v5_
            rhs295_ = default__.safeModuloInt(len(iife181_()
            ), len(d_2116_v24_))
            rhs296_ = default__.fm31(globalState)
            lhs185_ = d_2114_v22_
            lhs186_ = default__.safeIndex(45, (d_2114_v22_).length(0))
            lhs187_ = globalState
            lhs185_[lhs186_] = rhs293_
            d_2086_v5_ = rhs294_
            lhs187_.f3 = rhs295_
            d_2116_v24_ = rhs296_
            d_2118_v26_: C5
            nw346_ = C5()
            nw346_.ctor__((d_2084_v3_).f17, d_2084_v3_.f18)
            d_2118_v26_ = nw346_
            d_2119_v27_: _dafny.MultiSet
            d_2119_v27_ = _dafny.MultiSet([d_2081_v0_])
            d_2081_v0_ = (d_2119_v27_).ispropersubset((D18_DC41(d_2119_v27_)).cf57)
            d_2081_v0_ = d_2081_v0_
            (globalState).f3 = (0) - (((d_2084_v3_).f17) - ((0) - (p1)))
        elif True:
            d_2120_v28_: _dafny.Set
            d_2120_v28_ = _dafny.Set({p2, p1})
            d_2121_v29_: D3
            d_2121_v29_ = D3_DC6(d_2120_v28_)
            d_2122_v30_: str
            d_2122_v30_ = _dafny.CodePoint('h')
            source34_ = D9_DC21(d_2121_v29_, d_2122_v30_, ((d_2084_v3_).f17) > (default__.fm0(p0, 120, globalState)))
            if source34_.is_DC21:
                d_2123___mcc_h0_ = source34_.cf22
                d_2124___mcc_h1_ = source34_.cf23
                d_2125___mcc_h2_ = source34_.cf24
                d_2126_cf24_ = d_2125___mcc_h2_
                d_2127_cf23_ = d_2124___mcc_h1_
                d_2128_cf22_ = d_2123___mcc_h0_
                r3 = p0
                d_2129_v31_: _dafny.Array
                nw347_ = _dafny.Array(False, 27)
                d_2129_v31_ = nw347_
                index284_ = default__.safeIndex(608, (d_2129_v31_).length(0))
                (d_2129_v31_)[index284_] = d_2126_cf24_
                index285_ = default__.safeIndex(608, (d_2129_v31_).length(0))
                (d_2129_v31_)[index285_] = True
                d_2130_v32_: _dafny.Seq
                d_2130_v32_ = _dafny.SeqWithoutIsStrInference([(d_2126_cf24_ if not((d_2129_v31_)[default__.safeIndex(608, (d_2129_v31_).length(0))]) else d_2081_v0_)])
                d_2130_v32_ = _dafny.SeqWithoutIsStrInference([(32) > ((d_2084_v3_).f17)])
                d_2131_v33_: C5
                nw348_ = C5()
                nw348_.ctor__(len(d_2086_v5_), d_2083_v2_)
                d_2131_v33_ = nw348_
            elif True:
                d_2132___mcc_h3_ = source34_.cf21
                d_2133_cf21_ = d_2132___mcc_h3_
                d_2134_v34_: _dafny.Map
                d_2134_v34_ = _dafny.Map({d_2122_v30_: not(d_2081_v0_)})
                d_2135_v35_: _dafny.MultiSet
                d_2135_v35_ = _dafny.MultiSet([len(d_2134_v34_), 823])
                r0 = ((default__.fm53(globalState)).set((d_2084_v3_).f17, default__.abs((self).f11))).ispropersubset((d_2135_v35_).intersection(d_2135_v35_))
                d_2136_v36_: _dafny.Map
                d_2136_v36_ = _dafny.Map({(d_2084_v3_).f17: d_2086_v5_})
                d_2137_v37_: _dafny.Seq
                d_2137_v37_ = _dafny.SeqWithoutIsStrInference([d_2081_v0_])
                d_2138_v38_: _dafny.Map
                d_2138_v38_ = _dafny.Map({(d_2084_v3_).f17: (self).f11})
                d_2139_v40_: _dafny.Map
                d_2139_v40_ = _dafny.Map({d_2081_v0_: p0})
                d_2140_v41_: _dafny.Array
                nw349_ = _dafny.Array(None, 23)
                nw349_[int(0)] = (0) - ((d_2084_v3_).f17)
                nw349_[int(1)] = (d_2084_v3_).f17
                nw349_[int(2)] = (d_2084_v3_).f17
                nw349_[int(3)] = (d_2084_v3_).f17
                nw349_[int(4)] = 613
                nw349_[int(5)] = (d_2084_v3_).f17
                nw349_[int(6)] = p0
                nw349_[int(7)] = (self).fm5(233, (d_2084_v3_).f17, globalState)
                nw349_[int(8)] = (d_2084_v3_).f17
                nw349_[int(9)] = len((self).f9)
                nw349_[int(10)] = (d_2084_v3_).f17
                nw349_[int(11)] = ((d_2135_v35_)[(self).f11] if ((self).f11) in (d_2135_v35_) else default__.fm0(676, (d_2084_v3_).f17, globalState))
                nw349_[int(12)] = len(d_2137_v37_)
                nw349_[int(13)] = (self).f11
                nw349_[int(14)] = len(d_2086_v5_)
                nw349_[int(15)] = (d_2084_v3_).f17
                nw349_[int(16)] = len((self).f9)
                nw349_[int(17)] = len(default__.fm31(globalState))
                def iife182_():
                    coll84_ = _dafny.Set()
                    compr_84_: int
                    for compr_84_ in _dafny.IntegerRange(557, 890):
                        d_2141_v39_: int = compr_84_
                        if ((557) <= (d_2141_v39_)) and ((d_2141_v39_) < (890)):
                            coll84_ = coll84_.union(_dafny.Set([(d_2141_v39_) + (len(_dafny.Map({(d_2084_v3_).f17: len(d_2138_v38_)})))]))
                    return _dafny.Set(coll84_)
                nw349_[int(18)] = (0) - (((d_2138_v38_)[174] if (174) in (d_2138_v38_) else len(iife182_()
                )))
                nw349_[int(19)] = 599
                nw349_[int(20)] = default__.fm0(len(d_2139_v40_), (d_2084_v3_).f17, globalState)
                nw349_[int(21)] = p0
                nw349_[int(22)] = (0) - (len(_dafny.Map({len((self).f9): d_2122_v30_})))
                d_2140_v41_ = nw349_
                d_2142_v42_: _dafny.Map
                d_2142_v42_ = _dafny.Map({len(d_2136_v36_): d_2140_v41_})
                r2 = ((d_2142_v42_)[(len(d_2086_v5_)) - (p0)] if ((len(d_2086_v5_)) - (p0)) in (d_2142_v42_) else d_2140_v41_)
                (globalState).f3 = (0) - (p1)
                d_2143_v43_: D6
                d_2143_v43_ = D6_DC14(d_2140_v41_)
                d_2143_v43_ = d_2143_v43_
            d_2144_v44_: D18
            d_2144_v44_ = D18_DC41((_dafny.MultiSet([d_2081_v0_, d_2081_v0_])).set(d_2081_v0_, default__.abs((self).f11)))
            source35_ = d_2144_v44_
            if source35_.is_DC42:
                d_2145___mcc_h4_ = source35_.cf58
                d_2146___mcc_h5_ = source35_.cf59
                d_2147_cf59_ = d_2146___mcc_h5_
                d_2148_cf58_ = d_2145___mcc_h4_
                d_2149_v45_: _dafny.Array
                nw350_ = _dafny.Array(int(0), 6)
                d_2149_v45_ = nw350_
                d_2150_v46_: _dafny.Array
                nw351_ = _dafny.Array(None, 9)
                nw351_[int(0)] = d_2149_v45_
                nw351_[int(1)] = (d_2149_v45_ if d_2148_cf58_ else d_2149_v45_)
                nw351_[int(2)] = d_2149_v45_
                nw351_[int(3)] = d_2149_v45_
                nw351_[int(4)] = d_2149_v45_
                nw351_[int(5)] = d_2149_v45_
                nw351_[int(6)] = d_2149_v45_
                nw351_[int(7)] = d_2149_v45_
                nw351_[int(8)] = d_2149_v45_
                d_2150_v46_ = nw351_
                index286_ = default__.safeIndex(754, (d_2150_v46_).length(0))
                (d_2150_v46_)[index286_] = d_2149_v45_
                d_2151_v47_: _dafny.Seq
                d_2151_v47_ = _dafny.SeqWithoutIsStrInference([d_2120_v28_])
                index287_ = default__.safeIndex(754, (d_2150_v46_).length(0))
                rhs297_ = (d_2151_v47_)[default__.safeIndex((len(d_2120_v28_)) * ((0) - ((d_2084_v3_).f17)), len(d_2151_v47_))]
                rhs298_ = d_2149_v45_
                rhs299_ = (d_2084_v3_).f17
                rhs300_ = p2
                lhs188_ = d_2150_v46_
                lhs189_ = default__.safeIndex(754, (d_2150_v46_).length(0))
                lhs190_ = globalState
                lhs191_ = globalState
                d_2120_v28_ = rhs297_
                lhs188_[lhs189_] = rhs298_
                lhs190_.f3 = rhs299_
                lhs191_.f3 = rhs300_
                d_2152_v48_: C5
                nw352_ = C5()
                nw352_.ctor__(p1, d_2084_v3_.f18)
                d_2152_v48_ = nw352_
                d_2153_v49_: _dafny.Array
                nw353_ = _dafny.Array(_dafny.Set({}), 18)
                d_2153_v49_ = nw353_
                d_2153_v49_ = d_2153_v49_
                d_2154_v50_: _dafny.Seq
                d_2154_v50_ = _dafny.SeqWithoutIsStrInference([d_2081_v0_])
                d_2155_v51_: _dafny.Seq
                d_2155_v51_ = _dafny.SeqWithoutIsStrInference([(d_2150_v46_)[default__.safeIndex(754, (d_2150_v46_).length(0))]])
                d_2156_v52_: _dafny.Seq
                d_2156_v52_ = _dafny.SeqWithoutIsStrInference([d_2149_v45_, (d_2155_v51_)[default__.safeIndex((d_2084_v3_).f17, len(d_2155_v51_))], d_2149_v45_])
                d_2157_v53_: _dafny.Map
                d_2157_v53_ = _dafny.Map({((d_2152_v48_).f17 if d_2081_v0_ else len(d_2154_v50_)): (d_2155_v51_)[default__.safeIndex(p0, len(d_2155_v51_))]})
                d_2158_v54_: D11
                d_2158_v54_ = D11_DC25(d_2157_v53_)
                d_2159_v55_: _dafny.MultiSet
                d_2159_v55_ = _dafny.MultiSet([(d_2152_v48_).f17])
                d_2160_v56_: _dafny.Set
                d_2160_v56_ = _dafny.Set({d_2148_cf58_})
                rhs301_ = d_2148_cf58_
                rhs302_ = 383
                rhs303_ = (_dafny.Map({(d_2084_v3_).f17: (d_2150_v46_)[default__.safeIndex(754, (d_2150_v46_).length(0))]})) | ((d_2158_v54_).cf29)
                rhs304_ = (_dafny.MultiSet([len(d_2160_v56_)])).ispropersubset(d_2159_v55_)
                rhs305_ = d_2147_cf59_
                lhs192_ = globalState
                lhs193_ = globalState
                r0 = rhs301_
                lhs192_.f3 = rhs302_
                d_2157_v53_ = rhs303_
                r0 = rhs304_
                lhs193_.f3 = rhs305_
            elif source35_.is_DC43:
                d_2161___mcc_h6_ = source35_.cf60
                d_2162___mcc_h7_ = source35_.cf61
                d_2163_cf61_ = d_2162___mcc_h7_
                d_2164_cf60_ = d_2161___mcc_h6_
                (globalState).f3 = (0) - ((893) + (p1))
                d_2165_v57_: D14
                d_2165_v57_ = D14_DC34((0) - ((d_2084_v3_).f17), -861, len(_dafny.SeqWithoutIsStrInference([d_2122_v30_, d_2122_v30_])), d_2081_v0_)
                d_2166_v58_: C0
                nw354_ = C0()
                nw354_.ctor__(d_2164_cf60_, (d_2165_v57_).cf46, _dafny.SeqWithoutIsStrInference([d_2122_v30_ for d_2167_i6_ in range(default__.abs(-902))]), (self).f10)
                d_2166_v58_ = nw354_
                d_2168_v59_: _dafny.Set
                d_2168_v59_ = _dafny.Set({d_2164_cf60_, d_2081_v0_})
                d_2163_cf61_ = ((d_2084_v3_).f17 if (d_2168_v59_) == (_dafny.Set({d_2164_cf60_, (d_2166_v58_).f25, (d_2166_v58_).f25, default__.fm22(d_2081_v0_, (0) - ((d_2084_v3_).f17), globalState)})) else (d_2084_v3_).f17)
                d_2169_v60_: _dafny.MultiSet
                d_2169_v60_ = _dafny.MultiSet([(d_2084_v3_).f17, d_2163_cf61_, p2])
                d_2170_v61_: _dafny.MultiSet
                d_2170_v61_ = _dafny.MultiSet([default__.fm38(p2, globalState)])
                d_2171_v63_: _dafny.Map
                d_2171_v63_ = _dafny.Map({d_2122_v30_: default__.fm38((self).f11, globalState)})
                d_2172_v65_: _dafny.Map
                d_2172_v65_ = _dafny.Map({(d_2166_v58_).f25: (self).f11})
                d_2173_v66_: _dafny.Array
                nw355_ = _dafny.Array(None, 29)
                nw355_[int(0)] = p2
                nw355_[int(1)] = (d_2084_v3_).f17
                nw355_[int(2)] = (d_2169_v60_).cardinality
                nw355_[int(3)] = p0
                nw355_[int(4)] = 858
                nw355_[int(5)] = p0
                nw355_[int(6)] = d_2163_cf61_
                nw355_[int(7)] = 784
                nw355_[int(8)] = (d_2170_v61_).cardinality
                nw355_[int(9)] = (d_2084_v3_).f17
                nw355_[int(10)] = (d_2084_v3_).f17
                nw355_[int(11)] = (d_2084_v3_).f17
                nw355_[int(12)] = p2
                def iife183_():
                    coll85_ = _dafny.Map()
                    compr_85_: str
                    for compr_85_ in (d_2171_v63_).keys.Elements:
                        d_2174_v62_: str = compr_85_
                        if (d_2174_v62_) in (d_2171_v63_):
                            coll85_[d_2174_v62_] = _dafny.CodePoint('l')
                    return _dafny.Map(coll85_)
                nw355_[int(13)] = len((iife183_()
                ) | (default__.fm58(d_2164_cf60_, d_2169_v60_, (0) - (p0), globalState)))
                nw355_[int(14)] = (p2) + ((d_2084_v3_).f17)
                nw355_[int(15)] = p2
                nw355_[int(16)] = p2
                nw355_[int(17)] = (self).f11
                nw355_[int(18)] = (d_2084_v3_).f17
                nw355_[int(19)] = (d_2084_v3_).f17
                nw355_[int(20)] = p1
                nw355_[int(21)] = (301) - (786)
                nw355_[int(22)] = ((d_2084_v3_).f17) * ((d_2084_v3_).f17)
                def iife184_():
                    coll86_ = _dafny.Set()
                    compr_86_: int
                    for compr_86_ in _dafny.IntegerRange(826, 596):
                        d_2175_v64_: int = compr_86_
                        if ((826) <= (d_2175_v64_)) and ((d_2175_v64_) < (596)):
                            coll86_ = coll86_.union(_dafny.Set([default__.safeModuloInt(d_2175_v64_, len(d_2086_v5_))]))
                    return _dafny.Set(coll86_)
                nw355_[int(23)] = (0) - (len(iife184_()
                ))
                nw355_[int(24)] = ((d_2084_v3_).f17) + (p1)
                nw355_[int(25)] = (0) - ((364) - (len(_dafny.SeqWithoutIsStrInference([d_2164_cf60_]))))
                nw355_[int(26)] = len(d_2172_v65_)
                nw355_[int(27)] = (len((self).f9)) - ((d_2084_v3_).f17)
                nw355_[int(28)] = default__.safeModuloInt(d_2163_cf61_, default__.fm0(454, (d_2084_v3_).f17, globalState))
                d_2173_v66_ = nw355_
                index288_ = default__.safeIndex(39, (d_2173_v66_).length(0))
                (d_2173_v66_)[index288_] = ((d_2084_v3_).f17) * (p0)
                d_2176_v67_: _dafny.Array
                nw356_ = _dafny.Array(None, 2)
                nw356_[int(0)] = (d_2166_v58_).f25
                nw356_[int(1)] = (d_2166_v58_).f25
                d_2176_v67_ = nw356_
                d_2177_v68_: _dafny.Set
                d_2177_v68_ = _dafny.Set({d_2176_v67_, d_2176_v67_})
                index289_ = default__.safeIndex(39, (d_2173_v66_).length(0))
                rhs306_ = len(_dafny.Map({d_2086_v5_: d_2164_cf60_}))
                rhs307_ = len(d_2177_v68_)
                rhs308_ = False
                rhs309_ = default__.fm22((p0) <= (d_2163_cf61_), p0, globalState)
                lhs194_ = d_2173_v66_
                lhs195_ = default__.safeIndex(39, (d_2173_v66_).length(0))
                r3 = rhs306_
                lhs194_[lhs195_] = rhs307_
                r1 = rhs308_
                r0 = rhs309_
            elif source35_.is_DC41:
                d_2178___mcc_h8_ = source35_.cf57
                d_2179_cf57_ = d_2178___mcc_h8_
                (globalState).f3 = 639
                d_2081_v0_ = (not (d_2081_v0_) or (d_2081_v0_) if d_2081_v0_ else d_2081_v0_)
                d_2180_v69_: _dafny.Array
                def lambda89_(d_2181_v0_):
                    def lambda90_(d_2182_i7_):
                        return (_dafny.SeqWithoutIsStrInference([d_2181_v0_])) + (_dafny.SeqWithoutIsStrInference([d_2181_v0_, d_2181_v0_, d_2181_v0_]))

                    return lambda90_

                init50_ = lambda89_(d_2081_v0_)
                nw357_ = _dafny.Array(None, 8)
                for i0_50_ in range(nw357_.length(0)):
                    nw357_[i0_50_] = init50_(i0_50_)
                d_2180_v69_ = nw357_
                d_2183_v70_: _dafny.Seq
                d_2183_v70_ = _dafny.SeqWithoutIsStrInference([d_2081_v0_])
                d_2184_v71_: _dafny.Seq
                d_2184_v71_ = _dafny.SeqWithoutIsStrInference([(d_2183_v70_)[default__.safeIndex(p1, len(d_2183_v70_))], d_2081_v0_])
                index290_ = default__.safeIndex(600, (d_2180_v69_).length(0))
                (d_2180_v69_)[index290_] = d_2184_v71_
                d_2185_v72_: _dafny.Map
                d_2185_v72_ = _dafny.Map({d_2179_cf57_: (d_2084_v3_).f17})
                d_2186_v73_: _dafny.MultiSet
                d_2186_v73_ = _dafny.MultiSet([p0, (d_2084_v3_).fm23(globalState), default__.fm0(len(d_2185_v72_), p1, globalState), len(d_2183_v70_), len((self).f9)])
                index291_ = default__.safeIndex(600, (d_2180_v69_).length(0))
                (d_2180_v69_)[index291_] = (d_2183_v70_).set(default__.safeIndex(((d_2186_v73_)[(d_2084_v3_).f17] if ((d_2084_v3_).f17) in (d_2186_v73_) else (0) - ((d_2084_v3_).f17)), len(d_2183_v70_)), not(d_2081_v0_))
                d_2187_v74_: _dafny.Array
                nw358_ = _dafny.Array(int(0), 4)
                d_2187_v74_ = nw358_
                d_2188_v75_: _dafny.Array
                def lambda91_(d_2189_v30_):
                    def lambda92_(d_2190_i8_):
                        return d_2189_v30_

                    return lambda92_

                init51_ = lambda91_(d_2122_v30_)
                nw359_ = _dafny.Array(None, 7)
                for i0_51_ in range(nw359_.length(0)):
                    nw359_[i0_51_] = init51_(i0_51_)
                d_2188_v75_ = nw359_
                d_2191_v76_: _dafny.Set
                d_2191_v76_ = _dafny.Set({d_2188_v75_})
                rhs310_ = d_2187_v74_
                rhs311_ = len(((d_2191_v76_).intersection(_dafny.Set({d_2188_v75_}))) | ((d_2191_v76_).intersection(d_2191_v76_)))
                lhs196_ = globalState
                r2 = rhs310_
                lhs196_.f3 = rhs311_
            elif True:
                d_2192___mcc_h9_ = source35_.cf62
                d_2193_cf62_ = d_2192___mcc_h9_
                d_2122_v30_ = d_2122_v30_
                d_2194_v77_: _dafny.Array
                def lambda93_(d_2195_v3_):
                    def lambda94_(d_2196_i9_):
                        return (d_2196_i9_) - ((d_2195_v3_).f17)

                    return lambda94_

                init52_ = lambda93_(d_2084_v3_)
                nw360_ = _dafny.Array(None, 20)
                for i0_52_ in range(nw360_.length(0)):
                    nw360_[i0_52_] = init52_(i0_52_)
                d_2194_v77_ = nw360_
                index292_ = default__.safeIndex(163, (d_2194_v77_).length(0))
                (d_2194_v77_)[index292_] = (d_2084_v3_).f17
                index293_ = default__.safeIndex(163, (d_2194_v77_).length(0))
                (d_2194_v77_)[index293_] = (d_2084_v3_).f17
                index294_ = default__.safeIndex(163, (d_2194_v77_).length(0))
                rhs312_ = d_2194_v77_
                rhs313_ = d_2081_v0_
                rhs314_ = d_2081_v0_
                rhs315_ = (((d_2084_v3_).f17 if d_2081_v0_ else p1)) * (p2)
                lhs197_ = d_2194_v77_
                lhs198_ = default__.safeIndex(163, (d_2194_v77_).length(0))
                r2 = rhs312_
                r1 = rhs313_
                r1 = rhs314_
                lhs197_[lhs198_] = rhs315_
                d_2197_v78_: _dafny.Array
                nw361_ = _dafny.Array(None, 14)
                nw361_[int(0)] = False
                nw361_[int(1)] = False
                nw361_[int(2)] = not(d_2081_v0_)
                nw361_[int(3)] = d_2081_v0_
                nw361_[int(4)] = d_2081_v0_
                nw361_[int(5)] = not(d_2081_v0_)
                nw361_[int(6)] = d_2081_v0_
                nw361_[int(7)] = d_2081_v0_
                nw361_[int(8)] = d_2081_v0_
                nw361_[int(9)] = d_2081_v0_
                nw361_[int(10)] = d_2081_v0_
                nw361_[int(11)] = d_2081_v0_
                nw361_[int(12)] = d_2081_v0_
                nw361_[int(13)] = d_2081_v0_
                d_2197_v78_ = nw361_
                d_2198_v79_: _dafny.MultiSet
                d_2198_v79_ = _dafny.MultiSet([d_2197_v78_, d_2197_v78_, d_2197_v78_, d_2197_v78_])
                d_2199_v80_: _dafny.Map
                d_2199_v80_ = _dafny.Map({(d_2084_v3_).f17: d_2194_v77_})
                d_2200_v81_: D11
                d_2200_v81_ = D11_DC25(d_2199_v80_)
                rhs316_ = (d_2198_v79_).intersection(d_2198_v79_)
                rhs317_ = d_2200_v81_
                d_2198_v79_ = rhs316_
                d_2200_v81_ = rhs317_
            d_2201_v82_: _dafny.Seq
            d_2201_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydfxhmfrd"))
            d_2201_v82_ = (self).f9
            r1 = (_dafny.SeqWithoutIsStrInference([d_2122_v30_ for d_2202_i10_ in range(default__.abs(-426))])) != (((self).f9) + (_dafny.SeqWithoutIsStrInference([d_2122_v30_ for d_2203_i11_ in range(default__.abs(53))])))
            d_2204_v84_: _dafny.Seq
            d_2204_v84_ = _dafny.SeqWithoutIsStrInference([d_2081_v0_])
            d_2205_v85_: _dafny.MultiSet
            d_2205_v85_ = _dafny.MultiSet([len(_dafny.Map({_dafny.MultiSet(d_2204_v84_): d_2081_v0_})), p1, len(d_2086_v5_), (0) - (p1), 888])
            d_2206_v86_: _dafny.Seq
            d_2206_v86_ = d_2201_v82_
            d_2207_v87_: _dafny.Map
            def iife185_():
                coll87_ = _dafny.Map()
                compr_87_: int
                for compr_87_ in (d_2205_v85_).Elements:
                    d_2208_v83_: int = compr_87_
                    if (d_2208_v83_) in (d_2205_v85_):
                        coll87_[default__.safeModuloInt(d_2208_v83_, 533)] = False
                return _dafny.Map(coll87_)
            d_2207_v87_ = _dafny.Map({len(iife185_()
            ): d_2206_v86_})
            d_2207_v87_ = (d_2207_v87_).set((p1) * (740), d_2206_v86_)
        r0 = d_2081_v0_
        d_2209_v88_: _dafny.Seq
        d_2209_v88_ = _dafny.SeqWithoutIsStrInference([(self).f9])
        d_2210_v89_: D9
        d_2210_v89_ = D9_DC20(d_2209_v88_)
        pat_let_tv42_ = d_2081_v0_
        def lambda95_(source36_):
            if source36_.is_DC21:
                d_2211___mcc_h10_ = source36_.cf22
                d_2212___mcc_h11_ = source36_.cf23
                d_2213___mcc_h12_ = source36_.cf24
                d_2214_cf24_ = d_2213___mcc_h12_
                d_2215_cf23_ = d_2212___mcc_h11_
                d_2216_cf22_ = d_2211___mcc_h10_
                return d_2214_cf24_
            elif True:
                d_2217___mcc_h13_ = source36_.cf21
                d_2218_cf21_ = d_2217___mcc_h13_
                return pat_let_tv42_

        r1 = not(lambda95_(d_2210_v89_))
        d_2219_v90_: _dafny.Array
        def lambda96_(d_2220_p0_):
            def lambda97_(d_2221_i12_):
                return (d_2221_i12_) + (d_2220_p0_)

            return lambda97_

        init53_ = lambda96_(p0)
        nw362_ = _dafny.Array(None, 20)
        for i0_53_ in range(nw362_.length(0)):
            nw362_[i0_53_] = init53_(i0_53_)
        d_2219_v90_ = nw362_
        d_2222_v91_: _dafny.Map
        d_2222_v91_ = _dafny.Map({default__.safeModuloInt(p2, p1): d_2219_v90_})
        d_2223_v92_: _dafny.Seq
        d_2223_v92_ = _dafny.SeqWithoutIsStrInference([d_2081_v0_])
        r2 = ((d_2222_v91_)[len(d_2223_v92_)] if (len(d_2223_v92_)) in (d_2222_v91_) else d_2219_v90_)
        r3 = p0
        return r0, r1, r2, r3

    def m10(self, p0, globalState):
        r0: int = int(0)
        d_2224_v0_: bool
        d_2224_v0_ = False
        d_2225_v1_: _dafny.Seq
        d_2225_v1_ = _dafny.SeqWithoutIsStrInference([d_2224_v0_])
        d_2226_v2_: _dafny.Array
        nw363_ = _dafny.Array(None, 20)
        nw363_[int(0)] = default__.fm38((self).f11, globalState)
        nw363_[int(1)] = d_2224_v0_
        nw363_[int(2)] = d_2224_v0_
        nw363_[int(3)] = d_2224_v0_
        nw363_[int(4)] = True
        nw363_[int(5)] = not(d_2224_v0_)
        nw363_[int(6)] = default__.fm38((self).f11, globalState)
        nw363_[int(7)] = d_2224_v0_
        nw363_[int(8)] = d_2224_v0_
        nw363_[int(9)] = d_2224_v0_
        nw363_[int(10)] = d_2224_v0_
        nw363_[int(11)] = True
        nw363_[int(12)] = True
        nw363_[int(13)] = d_2224_v0_
        nw363_[int(14)] = default__.fm22(d_2224_v0_, (self).f11, globalState)
        nw363_[int(15)] = d_2224_v0_
        nw363_[int(16)] = False
        nw363_[int(17)] = d_2224_v0_
        nw363_[int(18)] = d_2224_v0_
        nw363_[int(19)] = True
        d_2226_v2_ = nw363_
        d_2227_v3_: _dafny.Map
        d_2227_v3_ = _dafny.Map({d_2226_v2_: d_2224_v0_})
        d_2228_v4_: _dafny.Map
        d_2228_v4_ = _dafny.Map({len(d_2227_v3_): d_2224_v0_})
        if not((d_2225_v1_)[default__.safeIndex(len(d_2228_v4_), len(d_2225_v1_))]):
            if d_2224_v0_:
                d_2229_v5_: _dafny.Array
                nw364_ = _dafny.Array(D4.default()(), 14)
                d_2229_v5_ = nw364_
                index295_ = default__.safeIndex(265, (d_2229_v5_).length(0))
                (d_2229_v5_)[index295_] = D4_DC11(d_2224_v0_, (self).f11)
                d_2230_v6_: _dafny.Map
                d_2230_v6_ = _dafny.Map({(d_2225_v1_)[default__.safeIndex((self).f11, len(d_2225_v1_))]: 630})
                d_2231_v7_: _dafny.Seq
                d_2231_v7_ = _dafny.SeqWithoutIsStrInference([len(d_2228_v4_)])
                d_2232_v8_: _dafny.Map
                d_2232_v8_ = _dafny.Map({d_2230_v6_: d_2231_v7_})
                index296_ = default__.safeIndex(265, (d_2229_v5_).length(0))
                (d_2229_v5_)[index296_] = D4_DC11(default__.fm38(len(d_2230_v6_), globalState), len(d_2232_v8_))
                d_2226_v2_ = d_2226_v2_
                (globalState).f3 = ((self).f11) * ((self).f11)
                d_2233_v9_: _dafny.Seq
                d_2233_v9_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxl")), (self).f9])
                d_2234_v10_: str
                d_2234_v10_ = _dafny.CodePoint('d')
                d_2235_v11_: D10
                d_2235_v11_ = D10_DC22(_dafny.MultiSet([(self).f9, (d_2233_v9_)[default__.safeIndex(len(_dafny.Map({d_2224_v0_: d_2234_v10_})), len(d_2233_v9_))]]))
                d_2236_v12_: _dafny.Set
                d_2236_v12_ = _dafny.Set({((d_2227_v3_)[d_2226_v2_] if (d_2226_v2_) in (d_2227_v3_) else d_2224_v0_)})
                d_2237_v13_: _dafny.MultiSet
                d_2237_v13_ = _dafny.MultiSet([d_2236_v12_])
                rhs318_ = d_2224_v0_
                rhs319_ = (self).f11
                rhs320_ = d_2235_v11_
                rhs321_ = (((d_2237_v13_).intersection(d_2237_v13_)).intersection(d_2237_v13_)).cardinality
                rhs322_ = (0) - ((self).f11)
                lhs199_ = globalState
                lhs200_ = globalState
                d_2224_v0_ = rhs318_
                lhs199_.f3 = rhs319_
                d_2235_v11_ = rhs320_
                lhs200_.f3 = rhs321_
                r0 = rhs322_
                d_2238_v14_: _dafny.Array
                nw365_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_2238_v14_ = nw365_
                d_2239_v15_: C5
                nw366_ = C5()
                nw366_.ctor__((self).f11, d_2238_v14_)
                d_2239_v15_ = nw366_
            elif True:
                d_2224_v0_ = d_2224_v0_
                d_2240_v16_: _dafny.Seq
                d_2240_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qa"))
                d_2241_v17_: _dafny.Seq
                d_2241_v17_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                rhs323_ = ((d_2241_v17_)[default__.safeIndex(-301, len(d_2241_v17_))] if True else (self).f11)
                rhs324_ = (self).f9
                lhs201_ = globalState
                lhs201_.f3 = rhs323_
                d_2240_v16_ = rhs324_
                d_2242_v18_: _dafny.Map
                d_2242_v18_ = _dafny.Map({d_2224_v0_: default__.fm38((self).f11, globalState)})
                d_2243_v19_: _dafny.Map
                d_2243_v19_ = _dafny.Map({d_2224_v0_: d_2242_v18_})
                d_2243_v19_ = (d_2243_v19_) | ((d_2243_v19_).set(d_2224_v0_, d_2242_v18_))
                (globalState).f3 = (self).f11
                r0 = ((self).f11) + ((self).f11)
            d_2244_v20_: _dafny.Array
            nw367_ = _dafny.Array(int(0), 15)
            d_2244_v20_ = nw367_
            index297_ = default__.safeIndex(627, (d_2244_v20_).length(0))
            (d_2244_v20_)[index297_] = default__.safeDivisionInt((self).f11, (self).f11)
            index298_ = default__.safeIndex(627, (d_2244_v20_).length(0))
            (d_2244_v20_)[index298_] = default__.safeDivisionInt(default__.safeDivisionInt(426, (self).f11), (0) - ((self).f11))
            d_2228_v4_ = d_2228_v4_
            d_2245_v21_: _dafny.Map
            d_2245_v21_ = _dafny.Map({default__.safeModuloInt((self).f11, -134): d_2244_v20_})
            d_2246_v22_: _dafny.Seq
            d_2246_v22_ = _dafny.SeqWithoutIsStrInference([d_2244_v20_])
            d_2245_v21_ = (d_2245_v21_).set((self).f11, (d_2246_v22_)[default__.safeIndex((d_2244_v20_)[default__.safeIndex(627, (d_2244_v20_).length(0))], len(d_2246_v22_))])
            d_2247_v23_: _dafny.Seq
            d_2247_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tll"))
            d_2247_v23_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxxnkdro"))) + ((d_2247_v23_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbuite"))))
        elif True:
            d_2248_v24_: _dafny.Array
            nw368_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
            d_2248_v24_ = nw368_
            index299_ = default__.safeIndex(499, (d_2248_v24_).length(0))
            (d_2248_v24_)[index299_] = default__.fm37((self).f11, len(d_2228_v4_), d_2224_v0_, (self).f11, globalState)
            index300_ = default__.safeIndex(499, (d_2248_v24_).length(0))
            (d_2248_v24_)[index300_] = (self).f9
            d_2249_v25_: str
            d_2249_v25_ = _dafny.CodePoint('a')
            d_2250_v26_: D14
            d_2250_v26_ = D14_DC33((self).f11, (d_2248_v24_)[default__.safeIndex(499, (d_2248_v24_).length(0))], d_2224_v0_)
            d_2251_v27_: _dafny.MultiSet
            d_2251_v27_ = _dafny.MultiSet([D14_DC33((self).f11, ((self).f9).set(default__.safeIndex((self).f11, len((self).f9)), d_2249_v25_), True), d_2250_v26_, d_2250_v26_])
            d_2252_v28_: _dafny.Map
            d_2252_v28_ = _dafny.Map({d_2224_v0_: d_2251_v27_})
            d_2252_v28_ = d_2252_v28_
            d_2253_v29_: _dafny.Seq
            d_2253_v29_ = _dafny.SeqWithoutIsStrInference([840, (self).f11, (0) - ((self).f11), (self).f11])
            d_2254_v30_: _dafny.MultiSet
            d_2254_v30_ = _dafny.MultiSet([(self).f11])
            d_2255_v31_: _dafny.Map
            d_2255_v31_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([(self).f11 for d_2256_i0_ in range(default__.abs(216))])) <= (d_2253_v29_): ((d_2254_v30_) | (d_2254_v30_)).cardinality})
            d_2255_v31_ = (d_2255_v31_).set(d_2224_v0_, default__.fm0((self).f11, (self).f11, globalState))
            d_2249_v25_ = _dafny.CodePoint('l')
            d_2257_v32_: _dafny.Seq
            d_2257_v32_ = _dafny.SeqWithoutIsStrInference([default__.fm59(d_2253_v29_, globalState)])
            d_2257_v32_ = d_2257_v32_
        d_2258_v33_: _dafny.Array
        nw369_ = _dafny.Array(D14.default()(), 10)
        d_2258_v33_ = nw369_
        d_2259_v34_: _dafny.Array
        nw370_ = _dafny.Array(None, 15)
        nw370_[int(0)] = d_2258_v33_
        nw370_[int(1)] = d_2258_v33_
        nw370_[int(2)] = (d_2258_v33_ if not(d_2224_v0_) else d_2258_v33_)
        nw370_[int(3)] = d_2258_v33_
        nw370_[int(4)] = d_2258_v33_
        nw370_[int(5)] = d_2258_v33_
        nw370_[int(6)] = d_2258_v33_
        nw370_[int(7)] = (d_2258_v33_ if d_2224_v0_ else d_2258_v33_)
        nw370_[int(8)] = d_2258_v33_
        nw370_[int(9)] = d_2258_v33_
        nw370_[int(10)] = d_2258_v33_
        nw370_[int(11)] = d_2258_v33_
        nw370_[int(12)] = d_2258_v33_
        nw370_[int(13)] = d_2258_v33_
        nw370_[int(14)] = d_2258_v33_
        d_2259_v34_ = nw370_
        index301_ = default__.safeIndex(247, (d_2259_v34_).length(0))
        (d_2259_v34_)[index301_] = d_2258_v33_
        d_2260_v35_: _dafny.Seq
        d_2260_v35_ = _dafny.SeqWithoutIsStrInference([d_2258_v33_])
        d_2261_v36_: _dafny.Array
        d_2261_v36_ = (d_2260_v35_)[default__.safeIndex((self).f11, len(d_2260_v35_))]
        index302_ = default__.safeIndex(247, (d_2259_v34_).length(0))
        (d_2259_v34_)[index302_] = ((d_2261_v36_) if not(default__.fm38((self).f11, globalState)) else d_2258_v33_)
        d_2262_v37_: _dafny.Array
        def lambda98_(d_2263_i1_):
            return default__.safeDivisionInt(d_2263_i1_, len((self).f9))

        init54_ = lambda98_
        nw371_ = _dafny.Array(None, 8)
        for i0_54_ in range(nw371_.length(0)):
            nw371_[i0_54_] = init54_(i0_54_)
        d_2262_v37_ = nw371_
        d_2264_v38_: _dafny.Seq
        d_2264_v38_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfudaygy"))])
        index303_ = default__.safeIndex(284, (d_2262_v37_).length(0))
        (d_2262_v37_)[index303_] = len((d_2264_v38_)[default__.safeIndex((self).f11, len(d_2264_v38_))])
        index304_ = default__.safeIndex(284, (d_2262_v37_).length(0))
        def lambda99_(source37_):
            d_2265___mcc_h0_ = source37_
            d_2266_cf0_ = d_2265___mcc_h0_
            return (self).f9

        (d_2262_v37_)[index304_] = len(lambda99_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbttwbn"))))
        index305_ = default__.safeIndex(369, (d_2226_v2_).length(0))
        (d_2226_v2_)[index305_] = d_2224_v0_
        d_2267_v39_: _dafny.Seq
        d_2267_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xugdjgxwd"))
        d_2268_v40_: D9
        d_2268_v40_ = D9_DC20(_dafny.SeqWithoutIsStrInference([d_2267_v39_]))
        d_2269_v41_: _dafny.MultiSet
        d_2269_v41_ = _dafny.MultiSet([(d_2262_v37_)[default__.safeIndex(284, (d_2262_v37_).length(0))]])
        d_2270_v42_: _dafny.Map
        d_2270_v42_ = _dafny.Map({d_2224_v0_: d_2269_v41_})
        d_2271_v43_: _dafny.Map
        d_2271_v43_ = _dafny.Map({((d_2270_v42_)[d_2224_v0_] if (d_2224_v0_) in (d_2270_v42_) else d_2269_v41_): d_2226_v2_})
        d_2272_v45_: _dafny.Seq
        def iife186_():
            coll88_ = _dafny.Set()
            compr_88_: int
            for compr_88_ in _dafny.IntegerRange(973, 522):
                d_2273_v44_: int = compr_88_
                if ((973) <= (d_2273_v44_)) and ((d_2273_v44_) < (522)):
                    coll88_ = coll88_.union(_dafny.Set([(d_2273_v44_) * ((self).f11)]))
            return _dafny.Set(coll88_)
        d_2272_v45_ = _dafny.SeqWithoutIsStrInference([(d_2271_v43_).set((d_2269_v41_).set(len(iife186_()
        ), default__.abs(96)), d_2226_v2_), d_2271_v43_, d_2271_v43_])
        index306_ = default__.safeIndex(369, (d_2226_v2_).length(0))
        rhs325_ = ((d_2264_v38_) + (_dafny.SeqWithoutIsStrInference([d_2267_v39_]))) <= (d_2264_v38_)
        rhs326_ = ((d_2268_v40_).cf21).set(default__.safeIndex(len((d_2272_v45_)[default__.safeIndex((self).f11, len(d_2272_v45_))]), len((d_2268_v40_).cf21)), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2274_i2_ in range(default__.abs(476))]))
        rhs327_ = _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('m') if False else _dafny.CodePoint('y')) for d_2275_i3_ in range(default__.abs(255))])
        lhs202_ = d_2226_v2_
        lhs203_ = default__.safeIndex(369, (d_2226_v2_).length(0))
        lhs202_[lhs203_] = rhs325_
        d_2264_v38_ = rhs326_
        d_2267_v39_ = rhs327_
        d_2276_v46_: _dafny.Seq
        d_2276_v46_ = _dafny.SeqWithoutIsStrInference([(d_2262_v37_)[default__.safeIndex(284, (d_2262_v37_).length(0))], (0) - (len((self).f9))])
        d_2277_v47_: C3
        nw372_ = C3()
        nw372_.ctor__(d_2224_v0_, (d_2262_v37_)[default__.safeIndex(284, (d_2262_v37_).length(0))], len(d_2276_v46_), self.f12, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtbgn")), (self).f10)
        d_2277_v47_ = nw372_
        d_2278_v48_: _dafny.Set
        d_2278_v48_ = _dafny.Set({(0) - ((self).f11)})
        d_2279_v49_: D3
        d_2279_v49_ = D3_DC6(d_2278_v48_)
        d_2280_v50_: str
        d_2280_v50_ = _dafny.CodePoint('u')
        d_2281_v51_: D9
        d_2281_v51_ = D9_DC21(d_2279_v49_, d_2280_v50_, d_2224_v0_)
        index307_ = default__.safeIndex(284, (d_2262_v37_).length(0))
        index308_ = default__.safeIndex(284, (d_2262_v37_).length(0))
        index309_ = default__.safeIndex(284, (d_2262_v37_).length(0))
        index310_ = default__.safeIndex(284, (d_2262_v37_).length(0))
        rhs328_ = (self).f11
        rhs329_ = (self).f11
        rhs330_ = ((d_2262_v37_)[default__.safeIndex(284, (d_2262_v37_).length(0))] if (d_2281_v51_).cf24 else (self).f11)
        rhs331_ = 762
        rhs332_ = (self).f11
        lhs204_ = d_2262_v37_
        lhs205_ = default__.safeIndex(284, (d_2262_v37_).length(0))
        lhs206_ = d_2262_v37_
        lhs207_ = default__.safeIndex(284, (d_2262_v37_).length(0))
        lhs208_ = d_2262_v37_
        lhs209_ = default__.safeIndex(284, (d_2262_v37_).length(0))
        lhs210_ = d_2262_v37_
        lhs211_ = default__.safeIndex(284, (d_2262_v37_).length(0))
        lhs212_ = globalState
        lhs204_[lhs205_] = rhs328_
        lhs206_[lhs207_] = rhs329_
        lhs208_[lhs209_] = rhs330_
        lhs210_[lhs211_] = rhs331_
        lhs212_.f3 = rhs332_
        d_2282_v52_: D7
        d_2282_v52_ = D7_DC15(_dafny.MultiSet([232]))
        r0 = ((self).f11) + (((d_2282_v52_).cf18).cardinality)
        return r0


class C7(T0, T1):
    def  __init__(self):
        self._f12: _dafny.Map = _dafny.Map({})
        self._f11: int = int(0)
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f11(self):
        return self._f11
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f9, f10, f11, f12):
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self).f12 = f12

    def fm5(self, p0, p1, globalState):
        return (self).f11

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_2283_v0_: str
        d_2283_v0_ = _dafny.CodePoint('v')
        d_2283_v0_ = d_2283_v0_
        d_2284_v1_: bool
        d_2284_v1_ = True
        d_2285_v2_: _dafny.Seq
        d_2285_v2_ = _dafny.SeqWithoutIsStrInference([d_2284_v1_, True, d_2284_v1_])
        if default__.fm13((d_2285_v2_)[default__.safeIndex((self).f11, len(d_2285_v2_))], len(_dafny.SeqWithoutIsStrInference([p1 for d_2286_i0_ in range(default__.abs(-222))])), ((self).f9) + ((self).f9), globalState):
            d_2287_v3_: _dafny.Map
            d_2287_v3_ = _dafny.Map({not(not(((self).f9) < ((self).f9))): p2})
            d_2287_v3_ = (d_2287_v3_).set(d_2284_v1_, (self).f11)
            d_2288_v4_: _dafny.Map
            d_2288_v4_ = _dafny.Map({(0) - (p0): len((self).f9)})
            d_2289_v5_: _dafny.Seq
            d_2289_v5_ = _dafny.SeqWithoutIsStrInference([d_2288_v4_, d_2288_v4_, d_2288_v4_, d_2288_v4_, d_2288_v4_])
            r3 = (0) - (len((d_2288_v4_) | ((default__.fm14(((self).f9)[default__.safeIndex(p2, len((self).f9))], d_2284_v1_, globalState) if d_2284_v1_ else (d_2289_v5_)[default__.safeIndex(p2, len(d_2289_v5_))]))))
            d_2290_v6_: _dafny.Array
            nw373_ = _dafny.Array(False, 7)
            d_2290_v6_ = nw373_
            d_2291_v7_: _dafny.Map
            d_2291_v7_ = _dafny.Map({(d_2290_v6_ if not(d_2284_v1_) else d_2290_v6_): d_2284_v1_})
            d_2292_v8_: _dafny.Map
            d_2292_v8_ = _dafny.Map({(self).f11: d_2284_v1_})
            d_2293_v9_: _dafny.Set
            d_2293_v9_ = _dafny.Set({D4_DC8(d_2292_v8_)})
            d_2291_v7_ = (d_2291_v7_).set(d_2290_v6_, (d_2293_v9_).issubset(default__.fm15(252, (self).f11, d_2284_v1_, d_2283_v0_, globalState)))
            d_2294_v10_: D2
            d_2294_v10_ = D2_DC5(not(not(d_2284_v1_)), d_2290_v6_, len(d_2292_v8_), len((self).f9))
            d_2295_v11_: _dafny.Map
            d_2295_v11_ = _dafny.Map({(D4_DC11(not(not(d_2284_v1_)), len((self).f9))).cf13: d_2294_v10_})
            d_2295_v11_ = d_2295_v11_
            d_2283_v0_ = d_2283_v0_
        elif True:
            d_2296_v12_: _dafny.Array
            nw374_ = _dafny.Array(False, 24)
            d_2296_v12_ = nw374_
            d_2297_v13_: _dafny.Array
            nw375_ = _dafny.Array(None, 17)
            nw375_[int(0)] = d_2296_v12_
            nw375_[int(1)] = d_2296_v12_
            nw375_[int(2)] = d_2296_v12_
            nw375_[int(3)] = d_2296_v12_
            nw375_[int(4)] = d_2296_v12_
            nw375_[int(5)] = d_2296_v12_
            nw375_[int(6)] = d_2296_v12_
            nw375_[int(7)] = d_2296_v12_
            nw375_[int(8)] = d_2296_v12_
            nw375_[int(9)] = (d_2296_v12_ if d_2284_v1_ else d_2296_v12_)
            nw375_[int(10)] = d_2296_v12_
            nw375_[int(11)] = d_2296_v12_
            nw375_[int(12)] = d_2296_v12_
            nw375_[int(13)] = d_2296_v12_
            nw375_[int(14)] = d_2296_v12_
            nw375_[int(15)] = d_2296_v12_
            nw375_[int(16)] = d_2296_v12_
            d_2297_v13_ = nw375_
            index311_ = default__.safeIndex(534, (d_2297_v13_).length(0))
            (d_2297_v13_)[index311_] = d_2296_v12_
            index312_ = default__.safeIndex(29, (d_2296_v12_).length(0))
            (d_2296_v12_)[index312_] = (d_2284_v1_ if (d_2285_v2_)[default__.safeIndex(p2, len(d_2285_v2_))] else d_2284_v1_)
            d_2298_v14_: _dafny.Seq
            d_2298_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "utiomrrs"))
            index313_ = default__.safeIndex(534, (d_2296_v12_).length(0))
            (d_2296_v12_)[index313_] = (d_2298_v14_) < (d_2298_v14_)
            index314_ = default__.safeIndex(534, (d_2297_v13_).length(0))
            index315_ = default__.safeIndex(29, (d_2296_v12_).length(0))
            index316_ = default__.safeIndex(534, (d_2296_v12_).length(0))
            rhs333_ = d_2296_v12_
            rhs334_ = d_2284_v1_
            rhs335_ = (self).f9
            rhs336_ = not(True)
            lhs213_ = d_2297_v13_
            lhs214_ = default__.safeIndex(534, (d_2297_v13_).length(0))
            lhs215_ = d_2296_v12_
            lhs216_ = default__.safeIndex(29, (d_2296_v12_).length(0))
            lhs217_ = d_2296_v12_
            lhs218_ = default__.safeIndex(534, (d_2296_v12_).length(0))
            lhs213_[lhs214_] = rhs333_
            lhs215_[lhs216_] = rhs334_
            d_2298_v14_ = rhs335_
            lhs217_[lhs218_] = rhs336_
            if (d_2296_v12_)[default__.safeIndex(29, (d_2296_v12_).length(0))]:
                index317_ = default__.safeIndex(29, (d_2296_v12_).length(0))
                (d_2296_v12_)[index317_] = (d_2283_v0_) not in ((d_2298_v14_) + (d_2298_v14_))
                d_2299_v15_: _dafny.Array
                nw376_ = _dafny.Array(_dafny.MultiSet({}), 25)
                d_2299_v15_ = nw376_
                d_2300_v16_: _dafny.MultiSet
                d_2300_v16_ = _dafny.MultiSet([(d_2296_v12_)[default__.safeIndex(29, (d_2296_v12_).length(0))]])
                index318_ = default__.safeIndex(269, (d_2299_v15_).length(0))
                (d_2299_v15_)[index318_] = (d_2300_v16_).intersection(d_2300_v16_)
                index319_ = default__.safeIndex(269, (d_2299_v15_).length(0))
                (d_2299_v15_)[index319_] = (d_2300_v16_).intersection(_dafny.MultiSet([(d_2296_v12_)[default__.safeIndex(29, (d_2296_v12_).length(0))], True]))
                (globalState).f3 = default__.fm0(len((self).f9), len(((self).f9).set(default__.safeIndex(p2, len((self).f9)), d_2283_v0_)), globalState)
                d_2301_v17_: _dafny.Array
                nw377_ = _dafny.Array(None, 10)
                nw377_[int(0)] = (0) - (p1)
                nw377_[int(1)] = p0
                nw377_[int(2)] = p1
                nw377_[int(3)] = 222
                nw377_[int(4)] = default__.fm0(p2, len(d_2298_v14_), globalState)
                nw377_[int(5)] = p0
                nw377_[int(6)] = p2
                nw377_[int(7)] = (self).f11
                nw377_[int(8)] = (self).f11
                nw377_[int(9)] = p0
                d_2301_v17_ = nw377_
                d_2302_v18_: _dafny.Seq
                d_2302_v18_ = _dafny.SeqWithoutIsStrInference([p2])
                d_2303_v19_: _dafny.Map
                d_2303_v19_ = _dafny.Map({d_2284_v1_: d_2283_v0_})
                (self).m8((self).f11, d_2301_v17_, (d_2302_v18_)[default__.safeIndex(len(d_2303_v19_), len(d_2302_v18_))], (d_2296_v12_)[default__.safeIndex(29, (d_2296_v12_).length(0))], globalState)
                r1 = True
            elif True:
                d_2304_v20_: D3
                d_2304_v20_ = D3_DC6(default__.fm16(len((self).f9), d_2284_v1_, p0, globalState))
                def iife187_():
                    coll89_ = _dafny.Set()
                    compr_89_: int
                    for compr_89_ in (_dafny.SeqWithoutIsStrInference([p0])).Elements:
                        d_2305_v21_: int = compr_89_
                        if (d_2305_v21_) in (_dafny.SeqWithoutIsStrInference([p0])):
                            coll89_ = coll89_.union(_dafny.Set([(d_2305_v21_) + (301)]))
                    return _dafny.Set(coll89_)
                d_2304_v20_ = (d_2304_v20_ if (d_2296_v12_)[default__.safeIndex(29, (d_2296_v12_).length(0))] else D3_DC6(iife187_()
))
                index320_ = default__.safeIndex(29, (d_2296_v12_).length(0))
                rhs337_ = ((_dafny.SeqWithoutIsStrInference([d_2283_v0_ for d_2306_i1_ in range(default__.abs(666))])) + (d_2298_v14_)) <= ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_2307_i2_ in range(default__.abs(-25))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmptpg"))))
                rhs338_ = len((self).f9)
                rhs339_ = (self).f11
                lhs219_ = d_2296_v12_
                lhs220_ = default__.safeIndex(29, (d_2296_v12_).length(0))
                lhs221_ = globalState
                lhs222_ = globalState
                lhs219_[lhs220_] = rhs337_
                lhs221_.f3 = rhs338_
                lhs222_.f3 = rhs339_
                r3 = default__.safeModuloInt((p2) - (len(_dafny.SeqWithoutIsStrInference([len(default__.fm14(d_2283_v0_, d_2284_v1_, globalState))]))), (self).f11)
                d_2308_v22_: _dafny.Map
                d_2308_v22_ = _dafny.Map({p0: len(d_2298_v14_)})
                d_2309_v23_: _dafny.Seq
                d_2309_v23_ = _dafny.SeqWithoutIsStrInference([((d_2308_v22_)[p2] if (p2) in (d_2308_v22_) else p2), (self).f11, p2, (self).f11])
                r0 = ((d_2309_v23_)[default__.safeIndex(p0, len(d_2309_v23_))]) > (len((d_2298_v14_) + (d_2298_v14_)))
                (globalState).f3 = -380
            r3 = default__.safeModuloInt(p0, (p0) + (p1))
            d_2310_v24_: _dafny.Map
            d_2310_v24_ = _dafny.Map({p2: (self).f11})
            (globalState).f3 = (len((d_2310_v24_) | (d_2310_v24_))) + ((self).f11)
            rhs340_ = p1
            rhs341_ = p0
            lhs223_ = globalState
            lhs224_ = globalState
            lhs223_.f3 = rhs340_
            lhs224_.f3 = rhs341_
        d_2311_v25_: _dafny.Array
        nw378_ = _dafny.Array(_dafny.Seq({}), 8)
        d_2311_v25_ = nw378_
        index321_ = default__.safeIndex(332, (d_2311_v25_).length(0))
        (d_2311_v25_)[index321_] = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_2284_v1_: d_2284_v1_})) for d_2312_i3_ in range(default__.abs(529))])
        d_2313_v26_: _dafny.Set
        d_2313_v26_ = _dafny.Set({p0, p1})
        d_2314_v27_: _dafny.Seq
        d_2314_v27_ = _dafny.SeqWithoutIsStrInference([len(d_2313_v26_), (self).f11, p2])
        index322_ = default__.safeIndex(332, (d_2311_v25_).length(0))
        (d_2311_v25_)[index322_] = d_2314_v27_
        d_2315_v28_: str
        d_2315_v28_ = _dafny.CodePoint('w')
        d_2316_v29_: _dafny.Array
        nw379_ = _dafny.Array(None, 11)
        nw379_[int(0)] = _dafny.CodePoint('i')
        nw379_[int(1)] = d_2283_v0_
        nw379_[int(2)] = d_2283_v0_
        nw379_[int(3)] = d_2283_v0_
        nw379_[int(4)] = d_2283_v0_
        nw379_[int(5)] = (d_2315_v28_)
        nw379_[int(6)] = d_2283_v0_
        nw379_[int(7)] = _dafny.CodePoint('e')
        nw379_[int(8)] = _dafny.CodePoint('w')
        nw379_[int(9)] = d_2283_v0_
        nw379_[int(10)] = d_2283_v0_
        d_2316_v29_ = nw379_
        d_2317_v30_: _dafny.Map
        d_2317_v30_ = _dafny.Map({d_2283_v0_: d_2316_v29_})
        d_2318_v31_: _dafny.Seq
        d_2318_v31_ = _dafny.SeqWithoutIsStrInference([d_2316_v29_, d_2316_v29_])
        d_2319_v32_: _dafny.Array
        nw380_ = _dafny.Array(None, 21)
        nw380_[int(0)] = d_2316_v29_
        nw380_[int(1)] = (D6_DC13(d_2316_v29_)).cf16
        nw380_[int(2)] = ((d_2317_v30_)[d_2283_v0_] if (d_2283_v0_) in (d_2317_v30_) else (d_2318_v31_)[default__.safeIndex(431, len(d_2318_v31_))])
        nw380_[int(3)] = d_2316_v29_
        nw380_[int(4)] = d_2316_v29_
        nw380_[int(5)] = d_2316_v29_
        nw380_[int(6)] = d_2316_v29_
        nw380_[int(7)] = (d_2318_v31_)[default__.safeIndex(p0, len(d_2318_v31_))]
        nw380_[int(8)] = d_2316_v29_
        nw380_[int(9)] = d_2316_v29_
        nw380_[int(10)] = d_2316_v29_
        nw380_[int(11)] = d_2316_v29_
        nw380_[int(12)] = d_2316_v29_
        nw380_[int(13)] = d_2316_v29_
        nw380_[int(14)] = d_2316_v29_
        nw380_[int(15)] = d_2316_v29_
        nw380_[int(16)] = d_2316_v29_
        nw380_[int(17)] = d_2316_v29_
        nw380_[int(18)] = d_2316_v29_
        nw380_[int(19)] = d_2316_v29_
        nw380_[int(20)] = d_2316_v29_
        d_2319_v32_ = nw380_
        index323_ = default__.safeIndex(515, (d_2319_v32_).length(0))
        (d_2319_v32_)[index323_] = d_2316_v29_
        d_2320_v33_: _dafny.Array
        nw381_ = _dafny.Array(None, 4)
        nw381_[int(0)] = d_2284_v1_
        nw381_[int(1)] = d_2284_v1_
        nw381_[int(2)] = d_2284_v1_
        nw381_[int(3)] = d_2284_v1_
        d_2320_v33_ = nw381_
        index324_ = default__.safeIndex(402, (d_2320_v33_).length(0))
        (d_2320_v33_)[index324_] = d_2284_v1_
        index325_ = default__.safeIndex(515, (d_2319_v32_).length(0))
        index326_ = default__.safeIndex(402, (d_2320_v33_).length(0))
        rhs342_ = d_2316_v29_
        rhs343_ = (_dafny.SeqWithoutIsStrInference([p1, p2])) <= (d_2314_v27_)
        rhs344_ = not((_dafny.CodePoint('f')) not in (((self).f9 if d_2284_v1_ else ((self).f9).set(default__.safeIndex(p1, len((self).f9)), ((self).f9)[default__.safeIndex(p1, len((self).f9))]))))
        lhs225_ = d_2319_v32_
        lhs226_ = default__.safeIndex(515, (d_2319_v32_).length(0))
        lhs227_ = d_2320_v33_
        lhs228_ = default__.safeIndex(402, (d_2320_v33_).length(0))
        lhs225_[lhs226_] = rhs342_
        lhs227_[lhs228_] = rhs343_
        r1 = rhs344_
        r3 = p1
        index327_ = default__.safeIndex(402, (d_2320_v33_).length(0))
        (d_2320_v33_)[index327_] = True
        r0 = ((self).f11) != ((self).f11)
        d_2321_v34_: _dafny.Seq
        d_2321_v34_ = _dafny.SeqWithoutIsStrInference([(self).f9, (self).f9, (self).f9, (self).f9])
        r1 = ((_dafny.SeqWithoutIsStrInference([(self).f9 for d_2322_i4_ in range(default__.abs(584))])) + (_dafny.SeqWithoutIsStrInference([(self).f9]))) <= ((_dafny.SeqWithoutIsStrInference([(self).f9 for d_2323_i5_ in range(default__.abs(67))])) + (d_2321_v34_))
        d_2324_v35_: _dafny.Array
        nw382_ = _dafny.Array(int(0), 26)
        d_2324_v35_ = nw382_
        r2 = (d_2324_v35_ if d_2284_v1_ else d_2324_v35_)
        pat_let_tv43_ = p1
        def iife188_(_pat_let49_0):
            def iife189_(d_2325_dt__update__tmp_h0_):
                def iife190_(_pat_let50_0):
                    def iife191_(d_2326_dt__update_hcf3_h0_):
                        return D2_DC4(d_2326_dt__update_hcf3_h0_)
                    return iife191_(_pat_let50_0)
                return iife190_(pat_let_tv43_)
            return iife189_(_pat_let49_0)
        r3 = (iife188_(D2_DC4(449))).cf3
        return r0, r1, r2, r3

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: bool = False
        r3: int = int(0)
        d_2327_v0_: bool
        d_2327_v0_ = False
        d_2328_v1_: _dafny.Map
        d_2328_v1_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2327_v0_, d_2327_v0_]): d_2327_v0_})
        d_2329_v2_: _dafny.Map
        d_2329_v2_ = _dafny.Map({d_2327_v0_: (self).f11})
        d_2328_v1_ = (d_2328_v1_).set(default__.fm2((self).f9, ((d_2329_v2_)[d_2327_v0_] if (d_2327_v0_) in (d_2329_v2_) else p0), globalState), d_2327_v0_)
        hi6_ = p0
        for d_2330_i0_ in range(p0, hi6_):
            d_2331_v3_: _dafny.Array
            nw383_ = _dafny.Array(False, 11)
            d_2331_v3_ = nw383_
            index328_ = default__.safeIndex(706, (d_2331_v3_).length(0))
            (d_2331_v3_)[index328_] = d_2327_v0_
            d_2332_v4_: str
            d_2332_v4_ = _dafny.CodePoint('r')
            d_2333_v5_: _dafny.Seq
            d_2333_v5_ = _dafny.SeqWithoutIsStrInference([d_2327_v0_])
            d_2334_v6_: _dafny.Map
            d_2334_v6_ = _dafny.Map({d_2327_v0_: d_2327_v0_})
            d_2335_v7_: _dafny.Map
            d_2335_v7_ = _dafny.Map({d_2327_v0_: d_2332_v4_})
            d_2336_v8_: _dafny.Array
            nw384_ = _dafny.Array(None, 13)
            nw384_[int(0)] = d_2332_v4_
            nw384_[int(1)] = d_2332_v4_
            nw384_[int(2)] = d_2332_v4_
            nw384_[int(3)] = d_2332_v4_
            nw384_[int(4)] = d_2332_v4_
            nw384_[int(5)] = d_2332_v4_
            nw384_[int(6)] = d_2332_v4_
            nw384_[int(7)] = default__.fm17(d_2327_v0_, globalState)
            nw384_[int(8)] = d_2332_v4_
            nw384_[int(9)] = (d_2332_v4_ if (d_2333_v5_)[default__.safeIndex((self).f11, len(d_2333_v5_))] else default__.fm17(((d_2334_v6_)[d_2327_v0_] if (d_2327_v0_) in (d_2334_v6_) else d_2327_v0_), globalState))
            nw384_[int(10)] = ((d_2335_v7_)[d_2327_v0_] if (d_2327_v0_) in (d_2335_v7_) else _dafny.CodePoint('l'))
            nw384_[int(11)] = d_2332_v4_
            nw384_[int(12)] = d_2332_v4_
            d_2336_v8_ = nw384_
            index329_ = default__.safeIndex(336, (d_2336_v8_).length(0))
            (d_2336_v8_)[index329_] = d_2332_v4_
            d_2337_v9_: _dafny.Array
            def lambda100_(d_2338_i1_):
                return default__.safeDivisionInt(d_2338_i1_, (self).f11)

            init55_ = lambda100_
            nw385_ = _dafny.Array(None, 24)
            for i0_55_ in range(nw385_.length(0)):
                nw385_[i0_55_] = init55_(i0_55_)
            d_2337_v9_ = nw385_
            index330_ = default__.safeIndex(747, (d_2337_v9_).length(0))
            (d_2337_v9_)[index330_] = (d_2330_i0_) * (325)
            d_2339_v10_: _dafny.MultiSet
            d_2339_v10_ = _dafny.MultiSet([(self).f11, (self).f11, default__.safeModuloInt(p0, d_2330_i0_)])
            d_2340_v11_: _dafny.Seq
            d_2340_v11_ = _dafny.SeqWithoutIsStrInference([d_2330_i0_])
            index331_ = default__.safeIndex(706, (d_2331_v3_).length(0))
            index332_ = default__.safeIndex(336, (d_2336_v8_).length(0))
            index333_ = default__.safeIndex(747, (d_2337_v9_).length(0))
            rhs345_ = ((d_2334_v6_)[d_2327_v0_] if (d_2327_v0_) in (d_2334_v6_) else d_2327_v0_)
            rhs346_ = (not(d_2327_v0_)) or (not((d_2327_v0_ if d_2327_v0_ else d_2327_v0_)))
            rhs347_ = d_2332_v4_
            rhs348_ = not ((d_2333_v5_)[default__.safeIndex(d_2330_i0_, len(d_2333_v5_))]) or (not (d_2327_v0_) or (d_2327_v0_))
            rhs349_ = ((d_2339_v10_)[p0] if (p0) in (d_2339_v10_) else (0) - (((d_2340_v11_)[default__.safeIndex(p0, len(d_2340_v11_))]) * ((self).f11)))
            lhs229_ = d_2331_v3_
            lhs230_ = default__.safeIndex(706, (d_2331_v3_).length(0))
            lhs231_ = d_2336_v8_
            lhs232_ = default__.safeIndex(336, (d_2336_v8_).length(0))
            lhs233_ = d_2337_v9_
            lhs234_ = default__.safeIndex(747, (d_2337_v9_).length(0))
            r2 = rhs345_
            lhs229_[lhs230_] = rhs346_
            lhs231_[lhs232_] = rhs347_
            r2 = rhs348_
            lhs233_[lhs234_] = rhs349_
            d_2341_v12_: _dafny.Array
            def lambda101_(d_2342_i2_):
                return _dafny.SeqWithoutIsStrInference([478 for d_2343_i3_ in range(default__.abs(416))])

            init56_ = lambda101_
            nw386_ = _dafny.Array(None, 28)
            for i0_56_ in range(nw386_.length(0)):
                nw386_[i0_56_] = init56_(i0_56_)
            d_2341_v12_ = nw386_
            d_2344_v13_: _dafny.Map
            d_2344_v13_ = _dafny.Map({d_2341_v12_: (len(_dafny.Map({d_2331_v3_: 220}))) < ((self).f11)})
            d_2344_v13_ = (d_2344_v13_).set(d_2341_v12_, default__.fm13(d_2327_v0_, -347, (self).f9, globalState))
            d_2337_v9_ = d_2337_v9_
            d_2345_v14_: _dafny.Seq
            d_2345_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbueos"))
            d_2346_v15_: _dafny.Map
            d_2346_v15_ = _dafny.Map({d_2327_v0_: ((self).f10)[default__.safeIndex(len(_dafny.Set({(d_2331_v3_)[default__.safeIndex(706, (d_2331_v3_).length(0))], (d_2331_v3_)[default__.safeIndex(706, (d_2331_v3_).length(0))], True, (d_2331_v3_)[default__.safeIndex(706, (d_2331_v3_).length(0))]})), len((self).f10))]})
            index334_ = default__.safeIndex(706, (d_2331_v3_).length(0))
            index335_ = default__.safeIndex(706, (d_2331_v3_).length(0))
            rhs350_ = d_2345_v14_
            rhs351_ = _dafny.MultiSet(((d_2346_v15_)[default__.fm13(d_2327_v0_, (0) - (d_2330_i0_), d_2345_v14_, globalState)] if (default__.fm13(d_2327_v0_, (0) - (d_2330_i0_), d_2345_v14_, globalState)) in (d_2346_v15_) else _dafny.SeqWithoutIsStrInference([(d_2337_v9_)[default__.safeIndex(747, (d_2337_v9_).length(0))] for d_2347_i4_ in range(default__.abs(-402))])))
            rhs352_ = (d_2345_v14_) < (d_2345_v14_)
            rhs353_ = (not (True) or ((d_2331_v3_)[default__.safeIndex(706, (d_2331_v3_).length(0))]) if True else False)
            rhs354_ = (d_2330_i0_) + (877)
            lhs235_ = d_2331_v3_
            lhs236_ = default__.safeIndex(706, (d_2331_v3_).length(0))
            lhs237_ = d_2331_v3_
            lhs238_ = default__.safeIndex(706, (d_2331_v3_).length(0))
            d_2345_v14_ = rhs350_
            d_2339_v10_ = rhs351_
            lhs235_[lhs236_] = rhs352_
            lhs237_[lhs238_] = rhs353_
            r3 = rhs354_
        d_2348_i5_: int
        d_2348_i5_ = 0
        with _dafny.label("12"):
            while d_2327_v0_:
                with _dafny.c_label("12"):
                    if (d_2348_i5_) >= (100):
                        raise _dafny.Break("12")
                    d_2348_i5_ = (d_2348_i5_) + (1)
                    d_2349_v16_: _dafny.Map
                    d_2349_v16_ = _dafny.Map({p0: d_2329_v2_})
                    d_2350_v17_: _dafny.Array
                    nw387_ = _dafny.Array(None, 12)
                    nw387_[int(0)] = (d_2329_v2_) | (d_2329_v2_)
                    nw387_[int(1)] = (d_2329_v2_) | (_dafny.Map({d_2327_v0_: (self).f11}))
                    nw387_[int(2)] = d_2329_v2_
                    nw387_[int(3)] = default__.fm18(globalState)
                    nw387_[int(4)] = d_2329_v2_
                    nw387_[int(5)] = d_2329_v2_
                    nw387_[int(6)] = (_dafny.Map({d_2327_v0_: p0})) | (_dafny.Map({d_2327_v0_: (self).f11}))
                    nw387_[int(7)] = d_2329_v2_
                    nw387_[int(8)] = (d_2329_v2_ if d_2327_v0_ else d_2329_v2_)
                    nw387_[int(9)] = (d_2329_v2_ if d_2327_v0_ else d_2329_v2_)
                    nw387_[int(10)] = (d_2329_v2_) | ((_dafny.Map({d_2327_v0_: p0})).set(d_2327_v0_, (0) - ((self).f11)))
                    nw387_[int(11)] = ((d_2349_v16_)[(self).f11] if ((self).f11) in (d_2349_v16_) else d_2329_v2_)
                    d_2350_v17_ = nw387_
                    index336_ = default__.safeIndex(144, (d_2350_v17_).length(0))
                    (d_2350_v17_)[index336_] = (default__.fm18(globalState)) | (d_2329_v2_)
                    index337_ = default__.safeIndex(144, (d_2350_v17_).length(0))
                    (d_2350_v17_)[index337_] = (d_2329_v2_ if d_2327_v0_ else _dafny.Map({d_2327_v0_: (self).f11}))
                    d_2351_v18_: _dafny.Array
                    def lambda102_(d_2352_v0_):
                        def lambda103_(d_2353_i6_):
                            return d_2352_v0_

                        return lambda103_

                    init57_ = lambda102_(d_2327_v0_)
                    nw388_ = _dafny.Array(None, 26)
                    for i0_57_ in range(nw388_.length(0)):
                        nw388_[i0_57_] = init57_(i0_57_)
                    d_2351_v18_ = nw388_
                    d_2354_v19_: str
                    d_2354_v19_ = _dafny.CodePoint('k')
                    d_2355_v20_: _dafny.Map
                    d_2355_v20_ = _dafny.Map({d_2354_v19_: (self).f11})
                    index338_ = default__.safeIndex(676, (d_2351_v18_).length(0))
                    (d_2351_v18_)[index338_] = default__.fm13(d_2327_v0_, len(d_2355_v20_), (self).f9, globalState)
                    d_2356_v21_: _dafny.Seq
                    d_2356_v21_ = _dafny.SeqWithoutIsStrInference([d_2327_v0_, d_2327_v0_, d_2327_v0_])
                    index339_ = default__.safeIndex(676, (d_2351_v18_).length(0))
                    (d_2351_v18_)[index339_] = (d_2356_v21_)[default__.safeIndex((self).f11, len(d_2356_v21_))]
                    d_2357_v22_: _dafny.Seq
                    d_2357_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejrycfmmc"))
                    d_2357_v22_ = (((self).f9) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsspel")))) + ((self).f9)
                    (globalState).f3 = len((d_2329_v2_) | ((d_2329_v2_).set((d_2351_v18_)[default__.safeIndex(676, (d_2351_v18_).length(0))], p0)))
                    pass
            pass
        d_2358_v23_: _dafny.Array
        def lambda104_(d_2359_i7_):
            return ((self).f9) <= ((self).f9)

        init58_ = lambda104_
        nw389_ = _dafny.Array(None, 6)
        for i0_58_ in range(nw389_.length(0)):
            nw389_[i0_58_] = init58_(i0_58_)
        d_2358_v23_ = nw389_
        index340_ = default__.safeIndex(941, (d_2358_v23_).length(0))
        (d_2358_v23_)[index340_] = d_2327_v0_
        index341_ = default__.safeIndex(941, (d_2358_v23_).length(0))
        (d_2358_v23_)[index341_] = d_2327_v0_
        d_2327_v0_ = d_2327_v0_
        hi7_ = p0
        for d_2360_i8_ in range(p0, hi7_):
            d_2361_v24_: _dafny.Seq
            d_2361_v24_ = _dafny.SeqWithoutIsStrInference([d_2329_v2_])
            d_2362_v25_: _dafny.Map
            d_2362_v25_ = _dafny.Map({d_2358_v23_: len(d_2361_v24_)})
            index342_ = default__.safeIndex(941, (d_2358_v23_).length(0))
            index343_ = default__.safeIndex(941, (d_2358_v23_).length(0))
            rhs355_ = default__.fm0(p0, p0, globalState)
            rhs356_ = p0
            rhs357_ = (d_2358_v23_) not in (d_2362_v25_)
            rhs358_ = (d_2358_v23_)[default__.safeIndex(941, (d_2358_v23_).length(0))]
            lhs239_ = globalState
            lhs240_ = d_2358_v23_
            lhs241_ = default__.safeIndex(941, (d_2358_v23_).length(0))
            lhs242_ = d_2358_v23_
            lhs243_ = default__.safeIndex(941, (d_2358_v23_).length(0))
            r3 = rhs355_
            lhs239_.f3 = rhs356_
            lhs240_[lhs241_] = rhs357_
            lhs242_[lhs243_] = rhs358_
            (globalState).f3 = ((d_2360_i8_) + ((self).f11)) - (p0)
            d_2363_v26_: _dafny.Array
            nw390_ = _dafny.Array(int(0), 28)
            d_2363_v26_ = nw390_
            def lambda105_(d_2364_i9_):
                return (d_2364_i9_) - (-587)

            init59_ = lambda105_
            nw391_ = _dafny.Array(None, 27)
            for i0_59_ in range(nw391_.length(0)):
                nw391_[i0_59_] = init59_(i0_59_)
            d_2363_v26_ = nw391_
            r2 = d_2327_v0_
        r0 = (p0) * (default__.safeDivisionInt(p0, p0))
        d_2365_v27_: _dafny.Array
        nw392_ = _dafny.Array(int(0), 18)
        d_2365_v27_ = nw392_
        d_2366_v28_: _dafny.Map
        d_2366_v28_ = _dafny.Map({d_2327_v0_: (self).f9})
        d_2367_v29_: _dafny.Map
        d_2367_v29_ = _dafny.Map({d_2365_v27_: d_2366_v28_})
        r1 = ((d_2367_v29_)[d_2365_v27_] if (d_2365_v27_) in (d_2367_v29_) else (d_2366_v28_).set(d_2327_v0_, (self).f9))
        r2 = (d_2358_v23_)[default__.safeIndex(941, (d_2358_v23_).length(0))]
        r3 = (self).f11
        return r0, r1, r2, r3

    def m3(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_2368_v0_: bool
        d_2368_v0_ = False
        (globalState).f3 = ((p0 if not(d_2368_v0_) else p2) if d_2368_v0_ else p2)
        d_2369_v1_: _dafny.Array
        nw393_ = _dafny.Array(int(0), 13)
        d_2369_v1_ = nw393_
        d_2370_v2_: _dafny.Map
        d_2370_v2_ = _dafny.Map({d_2368_v0_: (0) - ((0) - (p0))})
        index344_ = default__.safeIndex(331, (d_2369_v1_).length(0))
        (d_2369_v1_)[index344_] = (0) - (len(d_2370_v2_))
        d_2371_v3_: _dafny.Seq
        d_2371_v3_ = _dafny.SeqWithoutIsStrInference([p0])
        d_2372_v4_: _dafny.Map
        d_2372_v4_ = _dafny.Map({p0: d_2368_v0_})
        index345_ = default__.safeIndex(331, (d_2369_v1_).length(0))
        rhs359_ = (p2) in (_dafny.Map({len((self).f9): p3}))
        rhs360_ = (p2) - (((self).f11) + ((self).f11))
        rhs361_ = d_2368_v0_
        rhs362_ = (self).f11
        rhs363_ = ((d_2371_v3_)[default__.safeIndex(len(d_2372_v4_), len(d_2371_v3_))]) - (357)
        lhs244_ = globalState
        lhs245_ = globalState
        lhs246_ = d_2369_v1_
        lhs247_ = default__.safeIndex(331, (d_2369_v1_).length(0))
        d_2368_v0_ = rhs359_
        lhs244_.f3 = rhs360_
        d_2368_v0_ = rhs361_
        lhs245_.f3 = rhs362_
        lhs246_[lhs247_] = rhs363_
        if (d_2371_v3_) != ((d_2371_v3_) + (d_2371_v3_)):
            d_2373_v5_: _dafny.Array
            nw394_ = _dafny.Array(_dafny.CodePoint('D'), 22)
            d_2373_v5_ = nw394_
            d_2374_v6_: str
            d_2374_v6_ = _dafny.CodePoint('p')
            index346_ = default__.safeIndex(690, (d_2373_v5_).length(0))
            (d_2373_v5_)[index346_] = d_2374_v6_
            index347_ = default__.safeIndex(690, (d_2373_v5_).length(0))
            (d_2373_v5_)[index347_] = d_2374_v6_
            d_2375_v7_: _dafny.Array
            nw395_ = _dafny.Array(False, 20)
            d_2375_v7_ = nw395_
            index348_ = default__.safeIndex(388, (d_2375_v7_).length(0))
            (d_2375_v7_)[index348_] = d_2368_v0_
            d_2376_v8_: _dafny.Set
            d_2376_v8_ = _dafny.Set({d_2368_v0_})
            index349_ = default__.safeIndex(388, (d_2375_v7_).length(0))
            (d_2375_v7_)[index349_] = not(((d_2376_v8_).ispropersubset(d_2376_v8_) if d_2368_v0_ else d_2368_v0_))
            d_2377_v9_: _dafny.MultiSet
            d_2377_v9_ = _dafny.MultiSet([(d_2369_v1_)[default__.safeIndex(331, (d_2369_v1_).length(0))], 38, (d_2369_v1_)[default__.safeIndex(331, (d_2369_v1_).length(0))]])
            d_2378_v10_: int
            d_2379_v11_: bool
            out64_: int
            out65_: bool
            out64_, out65_ = (self).m9(p0, ((d_2377_v9_)[(0) - (p0)] if ((0) - (p0)) in (d_2377_v9_) else p2), globalState)
            d_2378_v10_ = out64_
            d_2379_v11_ = out65_
            d_2380_v12_: _dafny.Seq
            d_2380_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uraswbovw"))
            d_2381_v13_: _dafny.Seq
            d_2381_v13_ = _dafny.SeqWithoutIsStrInference([d_2379_v11_])
            d_2382_v14_: _dafny.Map
            d_2382_v14_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([default__.fm13(d_2368_v0_, (d_2369_v1_)[default__.safeIndex(331, (d_2369_v1_).length(0))], _dafny.SeqWithoutIsStrInference([d_2374_v6_ for d_2383_i0_ in range(default__.abs(184))]), globalState)]): d_2368_v0_})
            d_2384_v15_: _dafny.Seq
            d_2384_v15_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([(d_2381_v13_)[default__.safeIndex(p0, len(d_2381_v13_))], (d_2375_v7_)[default__.safeIndex(388, (d_2375_v7_).length(0))], False, d_2379_v11_, not(d_2379_v11_)])) not in (d_2382_v14_)])
            d_2385_v16_: _dafny.Map
            d_2385_v16_ = _dafny.Map({(self).f11: d_2375_v7_})
            rhs364_ = not((d_2384_v15_)[default__.safeIndex(len((d_2385_v16_).set(default__.fm0(p2, p2, globalState), d_2375_v7_)), len(d_2384_v15_))])
            rhs365_ = d_2378_v10_
            rhs366_ = (p1) + (p1)
            d_2379_v11_ = rhs364_
            d_2378_v10_ = rhs365_
            d_2380_v12_ = rhs366_
            d_2386_v17_: _dafny.Map
            d_2386_v17_ = _dafny.Map({((d_2372_v4_)[p0] if (p0) in (d_2372_v4_) else d_2368_v0_): (d_2375_v7_)[default__.safeIndex(388, (d_2375_v7_).length(0))]})
            (self).m8((self).f11, d_2369_v1_, len(d_2386_v17_), d_2379_v11_, globalState)
        elif True:
            d_2387_v18_: _dafny.Map
            d_2387_v18_ = _dafny.Map({d_2369_v1_: p2})
            d_2388_v19_: _dafny.Set
            d_2388_v19_ = _dafny.Set({(self).fm5(p0, (d_2369_v1_)[default__.safeIndex(331, (d_2369_v1_).length(0))], globalState), (d_2369_v1_)[default__.safeIndex(331, (d_2369_v1_).length(0))]})
            (globalState).f3 = ((d_2387_v18_)[d_2369_v1_] if (d_2369_v1_) in (d_2387_v18_) else len(d_2388_v19_))
            d_2389_v20_: C4
            nw396_ = C4()
            nw396_.ctor__((p0) < (default__.fm0(p2, p2, globalState)), d_2368_v0_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgdbmbs")) if True else (self).f9), _dafny.SeqWithoutIsStrInference([d_2371_v3_]))
            d_2389_v20_ = nw396_
            d_2390_v21_: _dafny.Array
            nw397_ = _dafny.Array(False, 18)
            d_2390_v21_ = nw397_
            index350_ = default__.safeIndex(747, (d_2390_v21_).length(0))
            (d_2390_v21_)[index350_] = d_2389_v20_.f19
            d_2391_v22_: _dafny.MultiSet
            d_2391_v22_ = _dafny.MultiSet([d_2389_v20_.f19, False])
            index351_ = default__.safeIndex(747, (d_2390_v21_).length(0))
            (d_2390_v21_)[index351_] = (d_2391_v22_).ispropersubset(_dafny.MultiSet([True]))
            index352_ = default__.safeIndex(331, (d_2369_v1_).length(0))
            (d_2369_v1_)[index352_] = (d_2369_v1_)[default__.safeIndex(331, (d_2369_v1_).length(0))]
            d_2392_v23_: _dafny.Set
            d_2392_v23_ = _dafny.Set({d_2369_v1_})
            if ((d_2392_v23_).intersection(d_2392_v23_)).ispropersubset(d_2392_v23_):
                (self).m8(p2, d_2369_v1_, ((d_2391_v22_).cardinality if d_2368_v0_ else p0), d_2389_v20_.f20, globalState)
                (globalState).f3 = p2
                d_2393_v24_: _dafny.Array
                def lambda106_(d_2394_i1_):
                    return (self).f9

                init60_ = lambda106_
                nw398_ = _dafny.Array(None, 9)
                for i0_60_ in range(nw398_.length(0)):
                    nw398_[i0_60_] = init60_(i0_60_)
                d_2393_v24_ = nw398_
                index353_ = default__.safeIndex(101, (d_2393_v24_).length(0))
                (d_2393_v24_)[index353_] = (self).f9
                index354_ = default__.safeIndex(101, (d_2393_v24_).length(0))
                (d_2393_v24_)[index354_] = p1
                d_2395_v25_: _dafny.Array
                nw399_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_2395_v25_ = nw399_
                d_2396_v26_: D20
                d_2396_v26_ = D20_DC46(d_2395_v25_)
                d_2395_v25_ = (d_2396_v26_).cf64
                index355_ = default__.safeIndex(331, (d_2369_v1_).length(0))
                (d_2369_v1_)[index355_] = default__.safeModuloInt(p0, (d_2369_v1_)[default__.safeIndex(331, (d_2369_v1_).length(0))])
            elif True:
                d_2397_v27_: _dafny.MultiSet
                d_2397_v27_ = _dafny.MultiSet([_dafny.MultiSet([(d_2390_v21_)[default__.safeIndex(747, (d_2390_v21_).length(0))], d_2389_v20_.f20, (d_2390_v21_)[default__.safeIndex(747, (d_2390_v21_).length(0))]]), d_2391_v22_, d_2391_v22_])
                d_2397_v27_ = d_2397_v27_
                (d_2389_v20_).f19 = d_2389_v20_.f19
                d_2398_v28_: _dafny.Seq
                d_2398_v28_ = _dafny.SeqWithoutIsStrInference([default__.fm22(d_2368_v0_, -55, globalState)])
                d_2399_v29_: _dafny.Seq
                d_2399_v29_ = _dafny.SeqWithoutIsStrInference([d_2398_v28_])
                d_2399_v29_ = d_2399_v29_
                (d_2389_v20_).f19 = ((d_2369_v1_)[default__.safeIndex(331, (d_2369_v1_).length(0))]) == ((0) - (len(default__.fm31(globalState))))
                (globalState).f3 = -147
        d_2400_v30_: _dafny.Array
        def lambda107_(d_2401_v0_):
            def lambda108_(d_2402_i2_):
                return not((d_2401_v0_ if d_2401_v0_ else d_2401_v0_))

            return lambda108_

        init61_ = lambda107_(d_2368_v0_)
        nw400_ = _dafny.Array(None, 18)
        for i0_61_ in range(nw400_.length(0)):
            nw400_[i0_61_] = init61_(i0_61_)
        d_2400_v30_ = nw400_
        nw401_ = _dafny.Array(False, 27)
        d_2400_v30_ = nw401_
        d_2403_v31_: D12
        d_2403_v31_ = D12_DC29(d_2368_v0_, (self).f11, d_2371_v3_, d_2368_v0_, 317)
        d_2404_v32_: str
        d_2404_v32_ = _dafny.CodePoint('y')
        d_2405_v33_: _dafny.Array
        nw402_ = _dafny.Array(None, 15)
        nw402_[int(0)] = ((self).f9).set(default__.safeIndex((self).f11, len((self).f9)), _dafny.CodePoint('v'))
        nw402_[int(1)] = p1
        nw402_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktgqa"))
        nw402_[int(3)] = p1
        nw402_[int(4)] = (self).f9
        nw402_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fc"))
        nw402_[int(6)] = ((self).f9).set(default__.safeIndex(len((d_2403_v31_).cf34), len((self).f9)), d_2404_v32_)
        nw402_[int(7)] = p1
        nw402_[int(8)] = (self).f9
        nw402_[int(9)] = p1
        nw402_[int(10)] = p1
        nw402_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ioore"))
        nw402_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sylgk"))
        nw402_[int(13)] = (self).f9
        nw402_[int(14)] = (self).f9
        d_2405_v33_ = nw402_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_2405_v33_).length(0)):
            d_2406_i3_: int = guard_loop_8_
            if (True) and (((0) <= (d_2406_i3_)) and ((d_2406_i3_) < ((d_2405_v33_).length(0)))):
                (d_2405_v33_)[(d_2406_i3_)] = ((self).f9) + ((((self).f9) + (p1)).set(default__.safeIndex(len(p1), len(((self).f9) + (p1))), d_2404_v32_))
        d_2407_v34_: _dafny.Array
        nw403_ = _dafny.Array(D1.default()(), 11)
        d_2407_v34_ = nw403_
        d_2408_v35_: _dafny.Array
        nw404_ = _dafny.Array(None, 1)
        nw404_[int(0)] = d_2407_v34_
        d_2408_v35_ = nw404_
        d_2409_v36_: C5
        nw405_ = C5()
        nw405_.ctor__((955 if not(d_2368_v0_) else len(p1)), d_2408_v35_)
        d_2409_v36_ = nw405_
        d_2410_v37_: _dafny.Array
        nw406_ = _dafny.Array(_dafny.Array(None, 0), 9)
        d_2410_v37_ = nw406_
        r0 = d_2410_v37_
        r1 = d_2369_v1_
        return r0, r1

    def m8(self, p0, p1, p2, p3, globalState):
        d_2411_v0_: bool
        d_2411_v0_ = True
        d_2411_v0_ = d_2411_v0_
        (globalState).f3 = p0
        d_2412_i0_: int
        d_2412_i0_ = 0
        with _dafny.label("13"):
            while d_2411_v0_:
                with _dafny.c_label("13"):
                    if (d_2412_i0_) >= (100):
                        raise _dafny.Break("13")
                    d_2412_i0_ = (d_2412_i0_) + (1)
                    d_2411_v0_ = d_2411_v0_
                    d_2413_v1_: _dafny.Array
                    def lambda109_(d_2414_p3_):
                        def lambda110_(d_2415_i1_):
                            return d_2414_p3_

                        return lambda110_

                    init62_ = lambda109_(p3)
                    nw407_ = _dafny.Array(None, 22)
                    for i0_62_ in range(nw407_.length(0)):
                        nw407_[i0_62_] = init62_(i0_62_)
                    d_2413_v1_ = nw407_
                    d_2416_v2_: _dafny.Seq
                    d_2416_v2_ = _dafny.SeqWithoutIsStrInference([d_2413_v1_])
                    d_2417_v3_: C2
                    nw408_ = C2()
                    nw408_.ctor__(len(d_2416_v2_), (self.f12) | (self.f12), (self).f9, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f11]) for d_2418_i2_ in range(default__.abs(46))]))
                    d_2417_v3_ = nw408_
                    index356_ = default__.safeIndex(116, (p1).length(0))
                    (p1)[index356_] = 803
                    d_2419_v4_: _dafny.Map
                    d_2419_v4_ = _dafny.Map({p3: (self).f11})
                    d_2420_v5_: _dafny.Map
                    d_2420_v5_ = _dafny.Map({914: d_2411_v0_})
                    d_2421_v6_: D2
                    d_2421_v6_ = D2_DC5(not(d_2411_v0_), d_2413_v1_, ((d_2419_v4_)[p3] if (p3) in (d_2419_v4_) else len(d_2420_v5_)), p2)
                    d_2422_v7_: _dafny.Map
                    d_2422_v7_ = _dafny.Map({756: d_2413_v1_})
                    pat_let_tv44_ = d_2422_v7_
                    pat_let_tv45_ = p0
                    index357_ = default__.safeIndex(116, (p1).length(0))
                    def iife192_(_pat_let51_0):
                        def iife193_(d_2423_dt__update__tmp_h0_):
                            def iife194_(_pat_let52_0):
                                def iife195_(d_2424_dt__update_hcf6_h0_):
                                    return D2_DC5((d_2423_dt__update__tmp_h0_).cf4, (d_2423_dt__update__tmp_h0_).cf5, d_2424_dt__update_hcf6_h0_, (d_2423_dt__update__tmp_h0_).cf7)
                                return iife195_(_pat_let52_0)
                            return iife194_((len(pat_let_tv44_)) - (pat_let_tv45_))
                        return iife193_(_pat_let51_0)
                    rhs367_ = d_2417_v3_
                    rhs368_ = True
                    rhs369_ = d_2413_v1_
                    rhs370_ = p2
                    rhs371_ = iife192_(D2_DC5(default__.fm38(p2, globalState), d_2413_v1_, 976, (self).f11))
                    lhs248_ = p1
                    lhs249_ = default__.safeIndex(116, (p1).length(0))
                    d_2417_v3_ = rhs367_
                    d_2411_v0_ = rhs368_
                    d_2413_v1_ = rhs369_
                    lhs248_[lhs249_] = rhs370_
                    d_2421_v6_ = rhs371_
                    d_2425_v8_: _dafny.Seq
                    d_2425_v8_ = _dafny.SeqWithoutIsStrInference([(self).f11, p0])
                    d_2425_v8_ = d_2425_v8_
                    d_2426_v9_: _dafny.Array
                    nw409_ = _dafny.Array(_dafny.Array(None, 0), 27)
                    d_2426_v9_ = nw409_
                    d_2427_v10_: str
                    d_2427_v10_ = _dafny.CodePoint('i')
                    d_2428_v11_: str
                    d_2428_v11_ = d_2427_v10_
                    d_2429_v12_: _dafny.MultiSet
                    d_2429_v12_ = _dafny.MultiSet([(p1)[default__.safeIndex(116, (p1).length(0))]])
                    d_2430_v13_: _dafny.Set
                    d_2430_v13_ = _dafny.Set({d_2429_v12_})
                    d_2431_v14_: D10
                    d_2431_v14_ = D10_DC23(d_2427_v10_, d_2430_v13_)
                    d_2432_v15_: _dafny.Map
                    d_2432_v15_ = _dafny.Map({(d_2428_v11_): (d_2431_v14_).cf26})
                    rhs372_ = ((0) - (p2)) == (((self).f11) * ((self).f11))
                    rhs373_ = (d_2426_v9_ if True else d_2426_v9_)
                    rhs374_ = (-875 if d_2411_v0_ else len(_dafny.SeqWithoutIsStrInference([d_2427_v10_ for d_2433_i3_ in range(default__.abs(-459))])))
                    rhs375_ = d_2432_v15_
                    lhs250_ = globalState
                    d_2411_v0_ = rhs372_
                    d_2426_v9_ = rhs373_
                    lhs250_.f3 = rhs374_
                    d_2432_v15_ = rhs375_
                    pass
            pass
        d_2434_v16_: _dafny.Array
        nw410_ = _dafny.Array(_dafny.CodePoint('D'), 22)
        d_2434_v16_ = nw410_
        d_2435_v17_: str
        d_2435_v17_ = _dafny.CodePoint('t')
        index358_ = default__.safeIndex(57, (d_2434_v16_).length(0))
        (d_2434_v16_)[index358_] = (d_2435_v17_ if True else d_2435_v17_)
        d_2436_v18_: str
        d_2436_v18_ = d_2435_v17_
        index359_ = default__.safeIndex(57, (d_2434_v16_).length(0))
        (d_2434_v16_)[index359_] = (d_2436_v18_)
        d_2437_v19_: _dafny.Array
        nw411_ = _dafny.Array(False, 16)
        d_2437_v19_ = nw411_
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, (d_2437_v19_).length(0)):
            d_2438_i4_: int = guard_loop_9_
            if (True) and (((0) <= (d_2438_i4_)) and ((d_2438_i4_) < ((d_2437_v19_).length(0)))):
                (d_2437_v19_)[(d_2438_i4_)] = p3
        if p3:
            d_2439_v20_: _dafny.MultiSet
            d_2439_v20_ = _dafny.MultiSet([p3, p3])
            d_2440_v21_: _dafny.Map
            d_2440_v21_ = _dafny.Map({d_2439_v20_: p3})
            d_2440_v21_ = ((d_2440_v21_).set(_dafny.MultiSet([p3]), not(p3))) | ((d_2440_v21_) | (_dafny.Map({_dafny.MultiSet([d_2411_v0_]): p3})))
            d_2411_v0_ = p3
            d_2441_v22_: _dafny.Array
            def lambda111_(d_2442_i5_):
                return (self).f9

            init63_ = lambda111_
            nw412_ = _dafny.Array(None, 29)
            for i0_63_ in range(nw412_.length(0)):
                nw412_[i0_63_] = init63_(i0_63_)
            d_2441_v22_ = nw412_
            d_2443_v23_: _dafny.Map
            d_2443_v23_ = _dafny.Map({p3: d_2441_v22_})
            d_2443_v23_ = (d_2443_v23_).set(d_2411_v0_, d_2441_v22_)
            d_2444_v24_: _dafny.Set
            d_2444_v24_ = _dafny.Set({d_2411_v0_})
            index360_ = default__.safeIndex(287, (p1).length(0))
            (p1)[index360_] = len(d_2444_v24_)
            d_2445_v25_: _dafny.Seq
            d_2445_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymacyooe"))
            index361_ = default__.safeIndex(98, (d_2437_v19_).length(0))
            (d_2437_v19_)[index361_] = d_2411_v0_
            d_2446_v26_: _dafny.Seq
            d_2446_v26_ = _dafny.SeqWithoutIsStrInference([p3, d_2411_v0_, not(p3)])
            index362_ = default__.safeIndex(287, (p1).length(0))
            index363_ = default__.safeIndex(98, (d_2437_v19_).length(0))
            rhs376_ = ((self).f11 if p3 else p0)
            rhs377_ = not (p3) or (p3)
            rhs378_ = ((self).f9) + (d_2445_v25_)
            rhs379_ = True
            rhs380_ = _dafny.MultiSet(d_2446_v26_)
            lhs251_ = p1
            lhs252_ = default__.safeIndex(287, (p1).length(0))
            lhs253_ = d_2437_v19_
            lhs254_ = default__.safeIndex(98, (d_2437_v19_).length(0))
            lhs251_[lhs252_] = rhs376_
            d_2411_v0_ = rhs377_
            d_2445_v25_ = rhs378_
            lhs253_[lhs254_] = rhs379_
            d_2439_v20_ = rhs380_
            d_2447_v27_: C0
            nw413_ = C0()
            nw413_.ctor__((d_2437_v19_)[default__.safeIndex(98, (d_2437_v19_).length(0))], True, _dafny.SeqWithoutIsStrInference([(d_2434_v16_)[default__.safeIndex(57, (d_2434_v16_).length(0))] for d_2448_i6_ in range(default__.abs(447))]), (self).f10)
            d_2447_v27_ = nw413_
            d_2449_v28_: D14
            d_2449_v28_ = D14_DC32(d_2447_v27_)
            source38_ = d_2449_v28_
            if source38_.is_DC33:
                d_2450___mcc_h0_ = source38_.cf40
                d_2451___mcc_h1_ = source38_.cf41
                d_2452___mcc_h2_ = source38_.cf42
                d_2453_cf42_ = d_2452___mcc_h2_
                d_2454_cf41_ = d_2451___mcc_h1_
                d_2455_cf40_ = d_2450___mcc_h0_
                d_2456_v29_: _dafny.Array
                nw414_ = _dafny.Array(False, 20)
                d_2456_v29_ = nw414_
                d_2457_v30_: _dafny.Map
                d_2457_v30_ = _dafny.Map({(d_2447_v27_).f26: p2})
                d_2458_v31_: _dafny.Map
                d_2458_v31_ = _dafny.Map({d_2456_v29_: default__.safeDivisionInt((self).f11, ((d_2457_v30_)[True] if (True) in (d_2457_v30_) else d_2455_cf40_))})
                d_2458_v31_ = (d_2458_v31_).set(d_2456_v29_, default__.safeDivisionInt(((d_2439_v20_)[d_2453_cf42_] if (d_2453_cf42_) in (d_2439_v20_) else p2), d_2455_cf40_))
                d_2459_v32_: _dafny.Seq
                d_2459_v32_ = _dafny.SeqWithoutIsStrInference([((d_2457_v30_)[d_2453_cf42_] if (d_2453_cf42_) in (d_2457_v30_) else p0), (p1)[default__.safeIndex(287, (p1).length(0))], (self).f11])
                index364_ = default__.safeIndex(287, (p1).length(0))
                (p1)[index364_] = (len(d_2459_v32_)) + ((p1)[default__.safeIndex(287, (p1).length(0))])
                index365_ = default__.safeIndex(287, (p1).length(0))
                (p1)[index365_] = (p1)[default__.safeIndex(287, (p1).length(0))]
                d_2455_cf40_ = p0
            elif source38_.is_DC34:
                d_2460___mcc_h3_ = source38_.cf43
                d_2461___mcc_h4_ = source38_.cf44
                d_2462___mcc_h5_ = source38_.cf45
                d_2463___mcc_h6_ = source38_.cf46
                d_2464_cf46_ = d_2463___mcc_h6_
                d_2465_cf45_ = d_2462___mcc_h5_
                d_2466_cf44_ = d_2461___mcc_h4_
                d_2467_cf43_ = d_2460___mcc_h3_
                index366_ = default__.safeIndex(98, (d_2437_v19_).length(0))
                (d_2437_v19_)[index366_] = not((d_2447_v27_).f26)
                d_2468_v33_: _dafny.Map
                d_2468_v33_ = _dafny.Map({d_2466_cf44_: (self).f11})
                d_2469_v34_: _dafny.MultiSet
                d_2469_v34_ = _dafny.MultiSet([p2])
                d_2470_v35_: _dafny.Map
                d_2470_v35_ = _dafny.Map({((self).f11) < ((0) - (((d_2468_v33_)[(d_2469_v34_).cardinality] if ((d_2469_v34_).cardinality) in (d_2468_v33_) else d_2467_cf43_))): (0) - (len(_dafny.Map({D12_DC28(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, (d_2447_v27_).f26, not(True), (d_2447_v27_).f25])), d_2466_cf44_])): (d_2437_v19_)[default__.safeIndex(98, (d_2437_v19_).length(0))]})))})
                d_2471_v36_: _dafny.Set
                d_2471_v36_ = _dafny.Set({(p1)[default__.safeIndex(287, (p1).length(0))], (self).f11})
                d_2470_v35_ = (d_2470_v35_).set(default__.fm38((p1)[default__.safeIndex(287, (p1).length(0))], globalState), default__.safeModuloInt(d_2466_cf44_, len(d_2471_v36_)))
                d_2472_v37_: D18
                d_2472_v37_ = D18_DC42(d_2464_cf46_, p0)
                d_2473_v38_: int
                d_2474_v39_: bool
                out66_: int
                out67_: bool
                out66_, out67_ = (self).m9(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hskkwpm"))), (d_2465_cf45_ if (d_2472_v37_).cf58 else d_2467_cf43_), globalState)
                d_2473_v38_ = out66_
                d_2474_v39_ = out67_
                index367_ = default__.safeIndex(287, (p1).length(0))
                (p1)[index367_] = d_2466_cf44_
            elif True:
                d_2475___mcc_h7_ = source38_.cf39
                d_2476_cf39_ = d_2475___mcc_h7_
                index368_ = default__.safeIndex(287, (p1).length(0))
                (p1)[index368_] = (-836) - ((p1)[default__.safeIndex(287, (p1).length(0))])
                (globalState).f3 = p0
                d_2477_v40_: _dafny.Array
                def lambda112_(d_2478_p2_):
                    def lambda113_(d_2479_i7_):
                        return (d_2479_i7_) + (d_2478_p2_)

                    return lambda113_

                init64_ = lambda112_(p2)
                nw415_ = _dafny.Array(None, 11)
                for i0_64_ in range(nw415_.length(0)):
                    nw415_[i0_64_] = init64_(i0_64_)
                d_2477_v40_ = nw415_
                d_2480_v41_: D6
                d_2480_v41_ = D6_DC14(d_2477_v40_)
                d_2477_v40_ = (d_2477_v40_ if (d_2447_v27_).f26 else (d_2480_v41_).cf17)
                d_2435_v17_ = default__.fm17((d_2447_v27_).f26, globalState)
        elif True:
            d_2481_v42_: _dafny.MultiSet
            d_2481_v42_ = _dafny.MultiSet([(d_2434_v16_)[default__.safeIndex(57, (d_2434_v16_).length(0))]])
            d_2411_v0_ = not(((d_2481_v42_).intersection(_dafny.MultiSet((self).f9))).ispropersubset(_dafny.MultiSet([d_2435_v17_, d_2435_v17_, d_2435_v17_])))
            d_2482_v43_: _dafny.Set
            d_2482_v43_ = _dafny.Set({p2})
            d_2483_v44_: _dafny.Seq
            d_2483_v44_ = _dafny.SeqWithoutIsStrInference([p0, (0) - (p2), default__.safeDivisionInt((self).f11, (self).fm5(len(d_2482_v43_), p2, globalState))])
            d_2483_v44_ = d_2483_v44_
            d_2484_v45_: _dafny.Map
            d_2484_v45_ = _dafny.Map({(self).f11: p2})
            d_2485_v46_: C2
            nw416_ = C2()
            nw416_.ctor__(p0, _dafny.Map({d_2484_v45_: d_2484_v45_}), (self).f9, (self).f10)
            d_2485_v46_ = nw416_
            d_2486_v47_: _dafny.Seq
            d_2486_v47_ = _dafny.SeqWithoutIsStrInference([d_2485_v46_])
            d_2487_v48_: _dafny.Array
            nw417_ = _dafny.Array(None, 7)
            nw417_[int(0)] = d_2485_v46_
            nw417_[int(1)] = d_2485_v46_
            nw417_[int(2)] = (d_2486_v47_)[default__.safeIndex((self).f11, len(d_2486_v47_))]
            nw417_[int(3)] = d_2485_v46_
            nw417_[int(4)] = d_2485_v46_
            nw417_[int(5)] = d_2485_v46_
            nw417_[int(6)] = d_2485_v46_
            d_2487_v48_ = nw417_
            d_2487_v48_ = d_2487_v48_
            d_2488_v49_: _dafny.Map
            d_2488_v49_ = _dafny.Map({((self).f11) - (p2): p3})
            d_2488_v49_ = (d_2488_v49_).set((self).fm5(p2, len((self).f9), globalState), ((self).f11) not in (d_2483_v44_))
            d_2411_v0_ = not(True)

    def m9(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        d_2489_v0_: bool
        d_2489_v0_ = False
        d_2490_v1_: _dafny.Map
        d_2490_v1_ = _dafny.Map({not(d_2489_v0_): 63})
        d_2491_v2_: _dafny.MultiSet
        d_2491_v2_ = _dafny.MultiSet([d_2489_v0_])
        d_2492_v3_: str
        d_2492_v3_ = _dafny.CodePoint('p')
        d_2493_v4_: D1
        d_2493_v4_ = D1_DC1(False)
        d_2494_v5_: _dafny.Seq
        d_2494_v5_ = _dafny.SeqWithoutIsStrInference([d_2493_v4_])
        d_2495_v6_: _dafny.MultiSet
        d_2495_v6_ = _dafny.MultiSet([p1])
        d_2496_v7_: _dafny.Map
        d_2496_v7_ = _dafny.Map({p1: d_2489_v0_})
        d_2497_v8_: _dafny.Seq
        d_2497_v8_ = _dafny.SeqWithoutIsStrInference([p0, p0, len(d_2496_v7_), len(default__.fm2(_dafny.SeqWithoutIsStrInference([d_2492_v3_ for d_2498_i1_ in range(default__.abs(431))]), (self).f11, globalState))])
        d_2499_v9_: _dafny.Seq
        d_2499_v9_ = _dafny.SeqWithoutIsStrInference([True])
        d_2500_v10_: _dafny.Array
        nw418_ = _dafny.Array(None, 19)
        nw418_[int(0)] = (0) - (p1)
        nw418_[int(1)] = ((self).f11) - (p1)
        nw418_[int(2)] = len((d_2490_v1_) | (d_2490_v1_))
        nw418_[int(3)] = (0) - (-806)
        nw418_[int(4)] = p1
        nw418_[int(5)] = len((self).f9)
        nw418_[int(6)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len((self).f9) for d_2501_i0_ in range(default__.abs(326))]))).cardinality
        nw418_[int(7)] = p1
        nw418_[int(8)] = default__.safeDivisionInt(((d_2491_v2_)[True] if (True) in (d_2491_v2_) else (self).f11), (self).f11)
        nw418_[int(9)] = (self).f11
        nw418_[int(10)] = (len(_dafny.Map({d_2492_v3_: d_2494_v5_})) if d_2489_v0_ else len((self).f9))
        nw418_[int(11)] = ((d_2495_v6_) | (_dafny.MultiSet(d_2497_v8_))).cardinality
        nw418_[int(12)] = default__.safeModuloInt(-110, ((d_2491_v2_)[(d_2499_v9_)[default__.safeIndex((self).f11, len(d_2499_v9_))]] if ((d_2499_v9_)[default__.safeIndex((self).f11, len(d_2499_v9_))]) in (d_2491_v2_) else 368))
        nw418_[int(13)] = (p0) * (len(_dafny.Set({p0})))
        nw418_[int(14)] = p1
        nw418_[int(15)] = (p1) * ((self).f11)
        nw418_[int(16)] = (self).f11
        nw418_[int(17)] = default__.safeModuloInt(p1, len((self).f9))
        nw418_[int(18)] = (p0) + (p0)
        d_2500_v10_ = nw418_
        index369_ = default__.safeIndex(149, (d_2500_v10_).length(0))
        (d_2500_v10_)[index369_] = p0
        d_2502_v11_: D1
        d_2502_v11_ = D1_DC2()
        d_2503_v12_: D4
        d_2503_v12_ = D4_DC9(d_2502_v11_)
        pat_let_tv46_ = p0
        pat_let_tv47_ = p0
        pat_let_tv48_ = p1
        pat_let_tv49_ = p0
        index370_ = default__.safeIndex(149, (d_2500_v10_).length(0))
        def lambda114_(source39_):
            if source39_.is_DC9:
                d_2504___mcc_h0_ = source39_.cf11
                d_2505_cf11_ = d_2504___mcc_h0_
                return (pat_let_tv46_) * (len(_dafny.SeqWithoutIsStrInference([(self).f9])))
            elif source39_.is_DC10:
                d_2506___mcc_h1_ = source39_.cf12
                d_2507_cf12_ = d_2506___mcc_h1_
                return pat_let_tv47_
            elif source39_.is_DC11:
                d_2508___mcc_h2_ = source39_.cf13
                d_2509___mcc_h3_ = source39_.cf14
                d_2510_cf14_ = d_2509___mcc_h3_
                d_2511_cf13_ = d_2508___mcc_h2_
                return (0) - (pat_let_tv48_)
            elif True:
                d_2512___mcc_h4_ = source39_.cf10
                d_2513_cf10_ = d_2512___mcc_h4_
                return (pat_let_tv49_) * ((0) - (len((self).f9)))

        rhs381_ = (default__.fm0(478, (self).f11, globalState) if d_2489_v0_ else p0)
        rhs382_ = lambda114_(d_2503_v12_)
        lhs255_ = d_2500_v10_
        lhs256_ = default__.safeIndex(149, (d_2500_v10_).length(0))
        r0 = rhs381_
        lhs255_[lhs256_] = rhs382_
        d_2514_v13_: _dafny.Map
        d_2514_v13_ = _dafny.Map({d_2492_v3_: 227})
        d_2515_v14_: D14
        d_2515_v14_ = D14_DC33(((d_2514_v13_)[d_2492_v3_] if (d_2492_v3_) in (d_2514_v13_) else (0) - ((0) - ((d_2500_v10_)[default__.safeIndex(149, (d_2500_v10_).length(0))]))), (self).f9, False)
        d_2516_v15_: C1
        nw419_ = C1()
        nw419_.ctor__(-983, (0) - ((d_2500_v10_)[default__.safeIndex(149, (d_2500_v10_).length(0))]), p1, default__.fm47(d_2489_v0_, d_2489_v0_, False, False, globalState), (self).f9, default__.fm21((d_2515_v14_).cf41, globalState))
        d_2516_v15_ = nw419_
        pat_let_tv50_ = d_2492_v3_
        index371_ = default__.safeIndex(149, (d_2500_v10_).length(0))
        def lambda115_(source40_):
            if source40_.is_DC42:
                d_2517___mcc_h5_ = source40_.cf58
                d_2518___mcc_h6_ = source40_.cf59
                d_2519_cf59_ = d_2518___mcc_h6_
                d_2520_cf58_ = d_2517___mcc_h5_
                return ((self).f9) + (_dafny.SeqWithoutIsStrInference([pat_let_tv50_ for d_2521_i2_ in range(default__.abs(375))]))
            elif source40_.is_DC43:
                d_2522___mcc_h7_ = source40_.cf60
                d_2523___mcc_h8_ = source40_.cf61
                d_2524_cf61_ = d_2523___mcc_h8_
                d_2525_cf60_ = d_2522___mcc_h7_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "maluwbg"))
            elif source40_.is_DC41:
                d_2526___mcc_h9_ = source40_.cf57
                d_2527_cf57_ = d_2526___mcc_h9_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phbkdec"))
            elif True:
                d_2528___mcc_h10_ = source40_.cf62
                d_2529_cf62_ = d_2528___mcc_h10_
                return ((self).f9) + ((self).f9)

        (d_2500_v10_)[index371_] = len(lambda115_(D18_DC42(d_2489_v0_, ((D18_DC41(d_2491_v2_)).cf57).cardinality)))
        d_2530_v16_: _dafny.Seq
        d_2530_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
        rhs383_ = ((self).f9).set(default__.safeIndex(default__.safeModuloInt(p0, (self).f11), len((self).f9)), d_2492_v3_)
        rhs384_ = -225
        d_2530_v16_ = rhs383_
        r0 = rhs384_
        if (276) >= (d_2516_v15_.f23):
            d_2531_v17_: _dafny.Set
            d_2531_v17_ = _dafny.Set({(d_2500_v10_)[default__.safeIndex(149, (d_2500_v10_).length(0))], (self).f11})
            d_2532_v18_: _dafny.Set
            d_2532_v18_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_2489_v0_, d_2489_v0_])), len(d_2531_v17_)})
            d_2533_v19_: _dafny.Set
            d_2533_v19_ = _dafny.Set({len(d_2532_v18_)})
            d_2534_v20_: D4
            d_2534_v20_ = D4_DC11((d_2533_v19_).ispropersubset(d_2533_v19_), p1)
            source41_ = d_2534_v20_
            if source41_.is_DC9:
                d_2535___mcc_h11_ = source41_.cf11
                d_2536_cf11_ = d_2535___mcc_h11_
                (d_2516_v15_).f23 = (d_2516_v15_).f24
                d_2489_v0_ = False
                r0 = (d_2497_v8_)[default__.safeIndex((0) - ((len(d_2499_v9_)) * (len(d_2530_v16_))), len(d_2497_v8_))]
                r1 = (len(d_2530_v16_)) == (len(default__.fm31(globalState)))
            elif source41_.is_DC10:
                d_2537___mcc_h12_ = source41_.cf12
                d_2538_cf12_ = d_2537___mcc_h12_
                d_2539_v21_: D15
                d_2539_v21_ = D15_DC36((0) - (((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_2500_v10_)[default__.safeIndex(149, (d_2500_v10_).length(0))] for d_2540_i3_ in range(default__.abs(974))]))).cardinality) + (p0)))
                d_2539_v21_ = d_2539_v21_
                index372_ = default__.safeIndex(149, (d_2500_v10_).length(0))
                (d_2500_v10_)[index372_] = default__.safeDivisionInt(((d_2495_v6_)[(d_2516_v15_).f24] if ((d_2516_v15_).f24) in (d_2495_v6_) else d_2516_v15_.f23), d_2516_v15_.f23)
                d_2541_v22_: C6
                nw420_ = C6()
                nw420_.ctor__((self).f11, (self.f12) | (self.f12), (d_2530_v16_) + ((d_2530_v16_).set(default__.safeIndex((d_2516_v15_).f24, len(d_2530_v16_)), d_2492_v3_)), (self).f10)
                d_2541_v22_ = nw420_
                d_2542_v23_: _dafny.Set
                d_2542_v23_ = _dafny.Set({d_2489_v0_})
                d_2543_v24_: _dafny.MultiSet
                d_2543_v24_ = _dafny.MultiSet([d_2542_v23_])
                d_2543_v24_ = default__.fm60(default__.fm0(p0, 954, globalState), globalState)
            elif source41_.is_DC11:
                d_2544___mcc_h13_ = source41_.cf13
                d_2545___mcc_h14_ = source41_.cf14
                d_2546_cf14_ = d_2545___mcc_h14_
                d_2547_cf13_ = d_2544___mcc_h13_
                d_2489_v0_ = (d_2489_v0_) and (d_2489_v0_)
                d_2499_v9_ = (d_2499_v9_) + (d_2499_v9_)
                index373_ = default__.safeIndex(149, (d_2500_v10_).length(0))
                (d_2500_v10_)[index373_] = d_2516_v15_.f23
                d_2547_cf13_ = d_2547_cf13_
            elif True:
                d_2548___mcc_h15_ = source41_.cf10
                d_2549_cf10_ = d_2548___mcc_h15_
                d_2550_v25_: _dafny.Array
                def lambda116_(d_2551_v0_):
                    def lambda117_(d_2552_i4_):
                        return d_2551_v0_

                    return lambda117_

                init65_ = lambda116_(d_2489_v0_)
                nw421_ = _dafny.Array(None, 16)
                for i0_65_ in range(nw421_.length(0)):
                    nw421_[i0_65_] = init65_(i0_65_)
                d_2550_v25_ = nw421_
                index374_ = default__.safeIndex(276, (d_2550_v25_).length(0))
                (d_2550_v25_)[index374_] = d_2489_v0_
                index375_ = default__.safeIndex(276, (d_2550_v25_).length(0))
                (d_2550_v25_)[index375_] = ((True) == (False) if False else d_2489_v0_)
                index376_ = default__.safeIndex(276, (d_2550_v25_).length(0))
                (d_2550_v25_)[index376_] = d_2489_v0_
                r0 = (0) - (default__.safeModuloInt((self).f11, 126))
                d_2553_v26_: _dafny.Set
                d_2553_v26_ = _dafny.Set({default__.fm38((self).f11, globalState)})
                d_2554_v27_: _dafny.Array
                nw422_ = _dafny.Array(None, 2)
                nw422_[int(0)] = d_2553_v26_
                nw422_[int(1)] = (_dafny.Set({(d_2550_v25_)[default__.safeIndex(276, (d_2550_v25_).length(0))], (d_2550_v25_)[default__.safeIndex(276, (d_2550_v25_).length(0))]})).intersection(default__.fm51(d_2489_v0_, (d_2516_v15_).f24, d_2492_v3_, (d_2550_v25_)[default__.safeIndex(276, (d_2550_v25_).length(0))], globalState))
                d_2554_v27_ = nw422_
                d_2554_v27_ = d_2554_v27_
            d_2555_v28_: _dafny.Array
            nw423_ = _dafny.Array(None, 9)
            nw423_[int(0)] = (d_2497_v8_).set(default__.safeIndex((d_2500_v10_)[default__.safeIndex(149, (d_2500_v10_).length(0))], len(d_2497_v8_)), d_2516_v15_.f23)
            nw423_[int(1)] = _dafny.SeqWithoutIsStrInference([698])
            nw423_[int(2)] = d_2497_v8_
            nw423_[int(3)] = d_2497_v8_
            nw423_[int(4)] = d_2497_v8_
            nw423_[int(5)] = _dafny.SeqWithoutIsStrInference([d_2516_v15_.f23, (d_2516_v15_).f24])
            nw423_[int(6)] = ((self).f10)[default__.safeIndex((d_2516_v15_).f24, len((self).f10))]
            nw423_[int(7)] = d_2497_v8_
            nw423_[int(8)] = d_2497_v8_
            d_2555_v28_ = nw423_
            d_2555_v28_ = d_2555_v28_
            (globalState).f3 = (d_2516_v15_).f24
            d_2556_v29_: D18
            d_2556_v29_ = D18_DC41(d_2491_v2_)
            rhs385_ = (0) - (d_2516_v15_.f23)
            rhs386_ = (d_2556_v29_).cf57
            lhs257_ = d_2516_v15_
            lhs257_.f23 = rhs385_
            d_2491_v2_ = rhs386_
            d_2557_v30_: C3
            nw424_ = C3()
            nw424_.ctor__(((d_2496_v7_)[d_2516_v15_.f23] if (d_2516_v15_.f23) in (d_2496_v7_) else (d_2499_v9_)[default__.safeIndex((d_2500_v10_)[default__.safeIndex(149, (d_2500_v10_).length(0))], len(d_2499_v9_))]), d_2516_v15_.f23, (0) - (d_2516_v15_.f23), self.f12, (self).f9, _dafny.SeqWithoutIsStrInference([d_2497_v8_ for d_2558_i5_ in range(default__.abs(-601))]))
            d_2557_v30_ = nw424_
        elif True:
            index377_ = default__.safeIndex(149, (d_2500_v10_).length(0))
            (d_2500_v10_)[index377_] = p0
            d_2491_v2_ = d_2491_v2_
            d_2559_v31_: _dafny.Array
            nw425_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_2559_v31_ = nw425_
            d_2560_v32_: C5
            nw426_ = C5()
            nw426_.ctor__((p0) + ((d_2516_v15_).f24), d_2559_v31_)
            d_2560_v32_ = nw426_
            d_2561_v33_: _dafny.Map
            d_2561_v33_ = _dafny.Map({d_2489_v0_: d_2489_v0_})
            d_2562_v34_: _dafny.Set
            d_2562_v34_ = _dafny.Set({(d_2561_v33_) | ((d_2561_v33_).set(d_2489_v0_, d_2489_v0_))})
            d_2562_v34_ = d_2562_v34_
            d_2563_v35_: _dafny.Map
            d_2563_v35_ = _dafny.Map({((d_2561_v33_).set(d_2489_v0_, d_2489_v0_)) | (d_2561_v33_): default__.fm0((d_2500_v10_)[default__.safeIndex(149, (d_2500_v10_).length(0))], (d_2516_v15_).f24, globalState)})
            d_2563_v35_ = (d_2563_v35_).set(_dafny.Map({d_2489_v0_: d_2489_v0_}), (self).f11)
        hi8_ = (d_2500_v10_)[default__.safeIndex(149, (d_2500_v10_).length(0))]
        for d_2564_i6_ in range(d_2516_v15_.f23, hi8_):
            index378_ = default__.safeIndex(149, (d_2500_v10_).length(0))
            (d_2500_v10_)[index378_] = d_2516_v15_.f23
            d_2565_v36_: _dafny.Array
            nw427_ = _dafny.Array(D1.default()(), 2)
            d_2565_v36_ = nw427_
            d_2566_v37_: _dafny.Seq
            d_2566_v37_ = _dafny.SeqWithoutIsStrInference([d_2565_v36_])
            d_2567_v38_: _dafny.Set
            d_2567_v38_ = _dafny.Set({(self).f9, (self).f9})
            d_2568_v39_: _dafny.Array
            nw428_ = _dafny.Array(None, 15)
            nw428_[int(0)] = (d_2566_v37_)[default__.safeIndex(len(d_2567_v38_), len(d_2566_v37_))]
            nw428_[int(1)] = d_2565_v36_
            nw428_[int(2)] = d_2565_v36_
            nw428_[int(3)] = d_2565_v36_
            nw428_[int(4)] = d_2565_v36_
            nw428_[int(5)] = d_2565_v36_
            nw428_[int(6)] = d_2565_v36_
            nw428_[int(7)] = d_2565_v36_
            nw428_[int(8)] = d_2565_v36_
            nw428_[int(9)] = d_2565_v36_
            nw428_[int(10)] = d_2565_v36_
            nw428_[int(11)] = d_2565_v36_
            nw428_[int(12)] = d_2565_v36_
            nw428_[int(13)] = d_2565_v36_
            nw428_[int(14)] = d_2565_v36_
            d_2568_v39_ = nw428_
            d_2569_v40_: C5
            nw429_ = C5()
            nw429_.ctor__(((d_2516_v15_).f24 if d_2489_v0_ else p1), d_2568_v39_)
            d_2569_v40_ = nw429_
            d_2570_v41_: _dafny.Map
            d_2570_v41_ = _dafny.Map({d_2492_v3_: d_2489_v0_})
            d_2571_v42_: _dafny.Map
            d_2571_v42_ = _dafny.Map({((d_2516_v15_).f24) - ((self).f11): d_2570_v41_})
            d_2571_v42_ = (d_2571_v42_).set(default__.safeModuloInt((self).f11, p0), d_2570_v41_)
            d_2572_v43_: _dafny.Array
            def lambda118_(d_2573_v0_):
                def lambda119_(d_2574_i7_):
                    return d_2573_v0_

                return lambda119_

            init66_ = lambda118_(d_2489_v0_)
            nw430_ = _dafny.Array(None, 3)
            for i0_66_ in range(nw430_.length(0)):
                nw430_[i0_66_] = init66_(i0_66_)
            d_2572_v43_ = nw430_
            d_2575_v44_: _dafny.Map
            d_2575_v44_ = _dafny.Map({p0: d_2572_v43_})
            d_2576_v45_: _dafny.Map
            d_2576_v45_ = _dafny.Map({(d_2516_v15_).f24: len(d_2575_v44_)})
            (globalState).f3 = (d_2516_v15_).fm5((d_2516_v15_).f24, len(d_2576_v45_), globalState)
        r0 = (d_2516_v15_).f24
        d_2577_v46_: D12
        d_2577_v46_ = D12_DC28(d_2497_v8_)
        pat_let_tv51_ = d_2489_v0_
        pat_let_tv52_ = d_2490_v1_
        pat_let_tv53_ = d_2516_v15_
        def lambda120_(source42_):
            if source42_.is_DC29:
                d_2578___mcc_h16_ = source42_.cf32
                d_2579___mcc_h17_ = source42_.cf33
                d_2580___mcc_h18_ = source42_.cf34
                d_2581___mcc_h19_ = source42_.cf35
                d_2582___mcc_h20_ = source42_.cf36
                d_2583_cf36_ = d_2582___mcc_h20_
                d_2584_cf35_ = d_2581___mcc_h19_
                d_2585_cf34_ = d_2580___mcc_h18_
                d_2586_cf33_ = d_2579___mcc_h17_
                d_2587_cf32_ = d_2578___mcc_h16_
                return pat_let_tv51_
            elif source42_.is_DC28:
                d_2588___mcc_h21_ = source42_.cf31
                d_2589_cf31_ = d_2588___mcc_h21_
                return True
            elif True:
                d_2590___mcc_h22_ = source42_.cf37
                d_2591_cf37_ = d_2590___mcc_h22_
                return (len(pat_let_tv52_)) <= ((0) - (pat_let_tv53_.f23))

        r1 = lambda120_(d_2577_v46_)
        return r0, r1


class C8:
    def  __init__(self):
        self.f16: bool = False
        self._f15: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self, f15, f16):
        (self)._f15 = f15
        (self).f16 = f16

    def m6(self, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        r2: int = int(0)
        if self.f16:
            (globalState).f3 = (self).f15
            if self.f16:
                (self).f16 = self.f16
                d_2592_v0_: _dafny.Array
                def lambda121_(d_2593_i0_):
                    return _dafny.Set({len(_dafny.SeqWithoutIsStrInference([self.f16, self.f16, self.f16])), (self).f15, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2594_i1_ in range(default__.abs(-696))]))})

                init67_ = lambda121_
                nw431_ = _dafny.Array(None, 17)
                for i0_67_ in range(nw431_.length(0)):
                    nw431_[i0_67_] = init67_(i0_67_)
                d_2592_v0_ = nw431_
                d_2595_v1_: _dafny.Set
                d_2595_v1_ = _dafny.Set({(self).f15})
                index379_ = default__.safeIndex(266, (d_2592_v0_).length(0))
                (d_2592_v0_)[index379_] = d_2595_v1_
                index380_ = default__.safeIndex(266, (d_2592_v0_).length(0))
                (d_2592_v0_)[index380_] = d_2595_v1_
                d_2596_v2_: _dafny.Seq
                d_2596_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "polvaf"))
                d_2597_v3_: _dafny.Seq
                d_2597_v3_ = _dafny.SeqWithoutIsStrInference([(self).f15, (self).f15, (self).f15])
                d_2598_v4_: _dafny.Map
                d_2598_v4_ = _dafny.Map({d_2596_v2_: d_2597_v3_})
                d_2599_v5_: _dafny.Set
                d_2599_v5_ = _dafny.Set({self.f16, self.f16})
                d_2600_v6_: _dafny.Map
                d_2600_v6_ = _dafny.Map({(self).f15: default__.fm10(self.f16, len(d_2599_v5_), default__.fm0(-106, (self).f15, globalState), (self).f15, globalState)})
                d_2601_v7_: _dafny.Array
                nw432_ = _dafny.Array(False, 28)
                d_2601_v7_ = nw432_
                d_2602_v8_: _dafny.Seq
                d_2602_v8_ = _dafny.SeqWithoutIsStrInference([d_2601_v7_, d_2601_v7_, d_2601_v7_, d_2601_v7_, d_2601_v7_])
                index381_ = default__.safeIndex(266, (d_2592_v0_).length(0))
                rhs387_ = d_2598_v4_
                rhs388_ = (default__.fm0((self).f15, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_2603_i2_ in range(default__.abs(481))])), globalState)) * (((self).f15) - ((self).f15))
                rhs389_ = default__.fm11(len(d_2602_v8_), globalState)
                rhs390_ = ((d_2592_v0_)[default__.safeIndex(266, (d_2592_v0_).length(0))]).intersection(d_2595_v1_)
                lhs258_ = globalState
                lhs259_ = d_2592_v0_
                lhs260_ = default__.safeIndex(266, (d_2592_v0_).length(0))
                d_2598_v4_ = rhs387_
                lhs258_.f3 = rhs388_
                d_2600_v6_ = rhs389_
                lhs259_[lhs260_] = rhs390_
                d_2604_v9_: _dafny.Map
                d_2604_v9_ = _dafny.Map({len(d_2596_v2_): (13) <= ((self).f15)})
                d_2605_v11_: _dafny.Set
                def iife196_():
                    coll90_ = _dafny.Set()
                    compr_90_: int
                    for compr_90_ in _dafny.IntegerRange(871, -47):
                        d_2606_v10_: int = compr_90_
                        if ((871) <= (d_2606_v10_)) and ((d_2606_v10_) < (-47)):
                            coll90_ = coll90_.union(_dafny.Set([(d_2606_v10_) + (-542)]))
                    return _dafny.Set(coll90_)
                d_2605_v11_ = _dafny.Set({iife196_()
                , (d_2592_v0_)[default__.safeIndex(266, (d_2592_v0_).length(0))]})
                d_2607_v12_: D4
                d_2607_v12_ = D4_DC10(d_2604_v9_)
                rhs391_ = (d_2605_v11_).ispropersubset(d_2605_v11_)
                rhs392_ = (d_2607_v12_).cf12
                rhs393_ = (136) > ((self).f15)
                lhs261_ = self
                r1 = rhs391_
                d_2604_v9_ = rhs392_
                lhs261_.f16 = rhs393_
                (self).f16 = (self.f16) or (self.f16)
            elif True:
                d_2608_v13_: _dafny.Map
                d_2608_v13_ = _dafny.Map({(self).f15: self.f16})
                (globalState).f3 = default__.fm0((self).f15, len(d_2608_v13_), globalState)
                d_2609_v14_: _dafny.Array
                def lambda122_(d_2610_i3_):
                    return True

                init68_ = lambda122_
                nw433_ = _dafny.Array(None, 23)
                for i0_68_ in range(nw433_.length(0)):
                    nw433_[i0_68_] = init68_(i0_68_)
                d_2609_v14_ = nw433_
                d_2611_v15_: _dafny.Seq
                d_2611_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xsrqyjxt"))
                rhs394_ = len((d_2611_v15_) + ((d_2611_v15_) + (d_2611_v15_)))
                rhs395_ = self.f16
                rhs396_ = (self).f15
                rhs397_ = self.f16
                rhs398_ = d_2609_v14_
                lhs262_ = globalState
                lhs262_.f3 = rhs394_
                r1 = rhs395_
                r2 = rhs396_
                r1 = rhs397_
                d_2609_v14_ = rhs398_
                d_2612_v16_: _dafny.Array
                nw434_ = _dafny.Array(_dafny.MultiSet({}), 23)
                d_2612_v16_ = nw434_
                d_2613_v17_: D2
                d_2613_v17_ = D2_DC4((self).f15)
                d_2614_v18_: _dafny.Map
                d_2614_v18_ = _dafny.Map({(self).f15: (self).f15})
                d_2615_v19_: _dafny.MultiSet
                d_2615_v19_ = _dafny.MultiSet([len(d_2614_v18_), (self).f15])
                d_2616_v20_: _dafny.Set
                def iife197_(_pat_let53_0):
                    def iife198_(d_2617_dt__update__tmp_h0_):
                        def iife199_(_pat_let54_0):
                            def iife200_(d_2618_dt__update_hcf3_h0_):
                                return D2_DC4(d_2618_dt__update_hcf3_h0_)
                            return iife200_(_pat_let54_0)
                        return iife199_(-670)
                    return iife198_(_pat_let53_0)
                d_2616_v20_ = _dafny.Set({d_2613_v17_, D2_DC4((d_2615_v19_).cardinality), iife197_(d_2613_v17_), d_2613_v17_, d_2613_v17_})
                d_2619_v21_: _dafny.Set
                d_2619_v21_ = _dafny.Set({self.f16, self.f16})
                d_2620_v22_: _dafny.Seq
                d_2620_v22_ = _dafny.SeqWithoutIsStrInference([len(d_2616_v20_), len(d_2619_v21_)])
                d_2621_v23_: _dafny.Map
                d_2621_v23_ = _dafny.Map({d_2620_v22_: self.f16})
                d_2622_v24_: _dafny.Map
                d_2622_v24_ = _dafny.Map({d_2612_v16_: d_2621_v23_})
                d_2623_v25_: str
                d_2623_v25_ = _dafny.CodePoint('r')
                d_2622_v24_ = (d_2622_v24_).set(d_2612_v16_, default__.fm12(d_2611_v15_, d_2623_v25_, (self).f15, globalState))
                r1 = not(self.f16)
                d_2624_v26_: _dafny.Array
                nw435_ = _dafny.Array(None, 3)
                nw435_[int(0)] = (self).f15
                nw435_[int(1)] = (self).f15
                nw435_[int(2)] = (self).f15
                d_2624_v26_ = nw435_
                d_2625_v27_: _dafny.MultiSet
                d_2625_v27_ = _dafny.MultiSet([d_2624_v26_])
                d_2626_v28_: D2
                d_2626_v28_ = D2_DC5(self.f16, d_2609_v14_, (self).f15, (self).f15)
                d_2627_v29_: _dafny.Map
                d_2627_v29_ = _dafny.Map({self.f16: (d_2620_v22_)[default__.safeIndex(len(d_2620_v22_), len(d_2620_v22_))]})
                rhs399_ = (len(d_2611_v15_)) + (436)
                rhs400_ = (((self).f15) + ((d_2620_v22_)[default__.safeIndex((self).f15, len(d_2620_v22_))])) + ((d_2625_v27_).cardinality)
                rhs401_ = default__.safeModuloInt(default__.safeDivisionInt((d_2626_v28_).cf7, (0) - ((self).f15)), (self).f15)
                rhs402_ = (self.f16) in (d_2627_v29_)
                lhs263_ = globalState
                lhs264_ = globalState
                lhs263_.f3 = rhs399_
                r2 = rhs400_
                lhs264_.f3 = rhs401_
                r1 = rhs402_
            d_2628_v30_: _dafny.MultiSet
            d_2628_v30_ = _dafny.MultiSet([self.f16, self.f16])
            d_2629_v31_: _dafny.Seq
            d_2629_v31_ = _dafny.SeqWithoutIsStrInference([self.f16, self.f16])
            d_2630_v32_: _dafny.Map
            d_2630_v32_ = _dafny.Map({(d_2628_v30_).cardinality: (d_2629_v31_)[default__.safeIndex((self).f15, len(d_2629_v31_))]})
            d_2631_v33_: _dafny.Map
            d_2631_v33_ = _dafny.Map({(d_2630_v32_) == (d_2630_v32_): (self.f16) == (self.f16)})
            r1 = ((d_2631_v33_)[self.f16] if (self.f16) in (d_2631_v33_) else self.f16)
            (globalState).f3 = (self).f15
            d_2630_v32_ = (d_2630_v32_).set((self).f15, self.f16)
        elif True:
            d_2632_v34_: _dafny.Map
            d_2632_v34_ = _dafny.Map({self.f16: ((self).f15) + ((self).f15)})
            d_2632_v34_ = (d_2632_v34_).set(self.f16, (0) - (((0) - ((self).f15)) - (-647)))
            d_2633_v35_: D2
            d_2633_v35_ = D2_DC4((self).f15)
            def iife201_(_pat_let55_0):
                def iife202_(d_2634_dt__update__tmp_h1_):
                    def iife203_(_pat_let56_0):
                        def iife204_(d_2635_dt__update_hcf3_h1_):
                            return D2_DC4(d_2635_dt__update_hcf3_h1_)
                        return iife204_(_pat_let56_0)
                    return iife203_((self).f15)
                return iife202_(_pat_let55_0)
            source43_ = iife201_(d_2633_v35_)
            if source43_.is_DC5:
                d_2636___mcc_h0_ = source43_.cf4
                d_2637___mcc_h1_ = source43_.cf5
                d_2638___mcc_h2_ = source43_.cf6
                d_2639___mcc_h3_ = source43_.cf7
                d_2640_cf7_ = d_2639___mcc_h3_
                d_2641_cf6_ = d_2638___mcc_h2_
                d_2642_cf5_ = d_2637___mcc_h1_
                d_2643_cf4_ = d_2636___mcc_h0_
                d_2644_v36_: _dafny.Map
                d_2644_v36_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_2643_cf4_])): d_2641_cf6_})
                d_2645_v37_: _dafny.Map
                d_2645_v37_ = _dafny.Map({d_2644_v36_: d_2644_v36_})
                d_2646_v38_: _dafny.Seq
                d_2646_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yd"))
                d_2647_v39_: C6
                nw436_ = C6()
                nw436_.ctor__(163, d_2645_v37_, d_2646_v38_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2640_cf7_ for d_2648_i5_ in range(default__.abs(215))]) for d_2649_i4_ in range(default__.abs(525))]))
                d_2647_v39_ = nw436_
                d_2650_v40_: _dafny.Array
                def lambda123_(d_2651_cf4_):
                    def lambda124_(d_2652_i6_):
                        return ((D8_DC17(_dafny.SeqWithoutIsStrInference([self.f16, self.f16]))).cf19) + (_dafny.SeqWithoutIsStrInference([d_2651_cf4_, self.f16]))

                    return lambda124_

                init69_ = lambda123_(d_2643_cf4_)
                nw437_ = _dafny.Array(None, 7)
                for i0_69_ in range(nw437_.length(0)):
                    nw437_[i0_69_] = init69_(i0_69_)
                d_2650_v40_ = nw437_
                d_2650_v40_ = d_2650_v40_
                d_2653_v41_: _dafny.Seq
                d_2653_v41_ = _dafny.SeqWithoutIsStrInference([d_2646_v38_])
                d_2654_v42_: C2
                nw438_ = C2()
                nw438_.ctor__(d_2640_cf7_, d_2645_v37_, (d_2653_v41_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([self.f16, self.f16, self.f16, d_2643_cf4_])), len(d_2653_v41_))], default__.fm21(d_2646_v38_, globalState))
                d_2654_v42_ = nw438_
                d_2655_v43_: _dafny.Seq
                d_2655_v43_ = _dafny.SeqWithoutIsStrInference([d_2641_cf6_, (0) - (d_2640_cf7_)])
                d_2656_v44_: _dafny.MultiSet
                d_2656_v44_ = _dafny.MultiSet([(self).f15])
                r1 = ((d_2656_v44_) | (_dafny.MultiSet(d_2655_v43_))).ispropersubset(_dafny.MultiSet((d_2655_v43_) + (d_2655_v43_)))
            elif True:
                d_2657___mcc_h4_ = source43_.cf3
                d_2658_cf3_ = d_2657___mcc_h4_
                d_2659_v45_: _dafny.Seq
                d_2659_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fd"))
                d_2660_v46_: _dafny.Seq
                d_2660_v46_ = _dafny.SeqWithoutIsStrInference([d_2658_cf3_])
                d_2661_v47_: _dafny.Seq
                d_2661_v47_ = _dafny.SeqWithoutIsStrInference([d_2660_v46_])
                d_2662_v48_: C4
                nw439_ = C4()
                nw439_.ctor__((self.f16) == (False), self.f16, (d_2659_v45_) + (d_2659_v45_), (d_2661_v47_) + (d_2661_v47_))
                d_2662_v48_ = nw439_
                d_2663_v49_: _dafny.Array
                nw440_ = _dafny.Array(int(0), 16)
                d_2663_v49_ = nw440_
                d_2664_v50_: _dafny.Map
                d_2664_v50_ = _dafny.Map({d_2663_v49_: d_2662_v48_.f19})
                (globalState).f3 = len((d_2664_v50_ if d_2662_v48_.f20 else (_dafny.Map({d_2663_v49_: d_2662_v48_.f20})) | (d_2664_v50_)))
                d_2659_v45_ = d_2659_v45_
                d_2665_v51_: D11
                d_2665_v51_ = D11_DC26()
                d_2666_v52_: D11
                d_2666_v52_ = D11_DC27(d_2665_v51_)
                d_2667_v53_: D11
                d_2667_v53_ = D11_DC27(d_2665_v51_)
                d_2668_v54_: _dafny.Seq
                d_2668_v54_ = _dafny.SeqWithoutIsStrInference([d_2662_v48_.f20])
                d_2667_v53_ = (D11_DC27(d_2665_v51_) if (d_2668_v54_) <= (_dafny.SeqWithoutIsStrInference([d_2662_v48_.f19])) else d_2667_v53_)
            if (((D14_DC34((self).f15, (0) - ((self).f15), (self).f15, self.f16)).cf46 if self.f16 else False)) and (self.f16):
                d_2669_v55_: _dafny.Set
                d_2669_v55_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bikixvxep"))})
                d_2670_v56_: str
                d_2670_v56_ = _dafny.CodePoint('x')
                r1 = not ((self.f16 if self.f16 else self.f16)) or (not((d_2669_v55_).isdisjoint(default__.fm61(d_2670_v56_, True, globalState))))
                d_2671_v57_: _dafny.Array
                nw441_ = _dafny.Array(_dafny.Map({}), 9)
                d_2671_v57_ = nw441_
                d_2671_v57_ = d_2671_v57_
                d_2672_v58_: _dafny.Seq
                d_2672_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eloimt"))
                d_2673_v59_: _dafny.Seq
                d_2673_v59_ = _dafny.SeqWithoutIsStrInference([self.f16])
                d_2674_v60_: D7
                d_2674_v60_ = D7_DC16()
                d_2675_v61_: _dafny.Map
                d_2675_v61_ = _dafny.Map({(self).f15: d_2674_v60_})
                rhs403_ = self.f16
                rhs404_ = not(not(((self).f15) > ((self).f15)))
                rhs405_ = default__.safeModuloInt((self).f15, (default__.fm62(default__.fm13(self.f16, (self).f15, d_2672_v58_, globalState), len(d_2672_v58_), d_2673_v59_, (self).f15, globalState)).cf48)
                rhs406_ = (d_2675_v61_) == ((d_2675_v61_) | (d_2675_v61_))
                lhs265_ = self
                lhs266_ = self
                lhs267_ = globalState
                lhs268_ = self
                lhs265_.f16 = rhs403_
                lhs266_.f16 = rhs404_
                lhs267_.f3 = rhs405_
                lhs268_.f16 = rhs406_
                d_2676_v62_: D17
                d_2676_v62_ = D17_DC40(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvbh")), self.f16, (self).f15)
                pat_let_tv54_ = d_2672_v58_
                d_2677_v63_: D4
                def iife205_(_pat_let57_0):
                    def iife206_(d_2678_dt__update__tmp_h2_):
                        def iife207_(_pat_let58_0):
                            def iife208_(d_2679_dt__update_hcf54_h0_):
                                return D17_DC40(d_2679_dt__update_hcf54_h0_, (d_2678_dt__update__tmp_h2_).cf55, (d_2678_dt__update__tmp_h2_).cf56)
                            return iife208_(_pat_let58_0)
                        return iife207_(pat_let_tv54_)
                    return iife206_(_pat_let57_0)
                d_2677_v63_ = D4_DC11(self.f16, len((iife205_(d_2676_v62_)).cf54))
                d_2680_v64_: _dafny.Set
                d_2680_v64_ = _dafny.Set({600, 993, (self).f15})
                d_2677_v63_ = D4_DC11(self.f16, len(d_2680_v64_))
                d_2632_v34_ = (d_2632_v34_).set(self.f16, ((self).f15) * ((self).f15))
            elif True:
                (self).f16 = False
                r1 = self.f16
                d_2681_v65_: str
                d_2681_v65_ = _dafny.CodePoint('r')
                d_2682_v66_: _dafny.Seq
                d_2682_v66_ = _dafny.SeqWithoutIsStrInference([d_2681_v65_])
                d_2683_v67_: _dafny.Map
                d_2683_v67_ = _dafny.Map({(self).f15: d_2682_v66_})
                d_2684_v68_: _dafny.MultiSet
                d_2684_v68_ = _dafny.MultiSet([self.f16])
                d_2685_v69_: D20
                d_2685_v69_ = D20_DC47(self.f16, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2686_i7_ in range(default__.abs(78))])), self.f16)
                d_2687_v70_: _dafny.Map
                d_2687_v70_ = _dafny.Map({len(d_2682_v66_): self.f16})
                d_2688_v71_: _dafny.Array
                nw442_ = _dafny.Array(None, 24)
                nw442_[int(0)] = self.f16
                nw442_[int(1)] = self.f16
                nw442_[int(2)] = self.f16
                nw442_[int(3)] = ((self).f15) < (len(default__.fm26((self).f15, d_2683_v67_, globalState)))
                nw442_[int(4)] = (_dafny.MultiSet([self.f16, self.f16, self.f16, not(self.f16)])).ispropersubset(d_2684_v68_)
                nw442_[int(5)] = not(default__.fm13(self.f16, (self).f15, d_2682_v66_, globalState))
                nw442_[int(6)] = self.f16
                nw442_[int(7)] = self.f16
                nw442_[int(8)] = self.f16
                nw442_[int(9)] = self.f16
                nw442_[int(10)] = default__.fm13(self.f16, (self).f15, d_2682_v66_, globalState)
                nw442_[int(11)] = not(self.f16)
                nw442_[int(12)] = self.f16
                nw442_[int(13)] = (d_2685_v69_).cf65
                nw442_[int(14)] = self.f16
                nw442_[int(15)] = self.f16
                nw442_[int(16)] = ((d_2687_v70_)[(0) - ((self).f15)] if ((0) - ((self).f15)) in (d_2687_v70_) else self.f16)
                nw442_[int(17)] = (self.f16 if False else self.f16)
                nw442_[int(18)] = self.f16
                nw442_[int(19)] = self.f16
                nw442_[int(20)] = self.f16
                nw442_[int(21)] = self.f16
                nw442_[int(22)] = self.f16
                nw442_[int(23)] = self.f16
                d_2688_v71_ = nw442_
                index382_ = default__.safeIndex(771, (d_2688_v71_).length(0))
                (d_2688_v71_)[index382_] = self.f16
                index383_ = default__.safeIndex(771, (d_2688_v71_).length(0))
                (d_2688_v71_)[index383_] = not(self.f16)
                d_2632_v34_ = (d_2632_v34_).set((d_2688_v71_)[default__.safeIndex(771, (d_2688_v71_).length(0))], (self).f15)
                d_2689_v72_: D3
                d_2689_v72_ = D3_DC7(False)
                d_2690_v73_: _dafny.Seq
                d_2690_v73_ = _dafny.SeqWithoutIsStrInference([not(not ((d_2688_v71_)[default__.safeIndex(771, (d_2688_v71_).length(0))]) or ((d_2688_v71_)[default__.safeIndex(771, (d_2688_v71_).length(0))])), (-879) <= (len(_dafny.SeqWithoutIsStrInference([(self).f15 for d_2691_i8_ in range(default__.abs(700))]))), (d_2689_v72_).cf9, self.f16, not(((d_2688_v71_)[default__.safeIndex(771, (d_2688_v71_).length(0))]) and (self.f16))])
                r1 = not((d_2690_v73_)[default__.safeIndex((self).f15, len(d_2690_v73_))])
            if (default__.safeModuloInt((self).f15, (self).f15)) == ((self).f15):
                d_2692_v74_: _dafny.Map
                d_2692_v74_ = _dafny.Map({self.f16: self.f16})
                d_2632_v34_ = (d_2632_v34_).set(((self).f15) >= (len(d_2692_v74_)), (self).f15)
                d_2693_v75_: str
                d_2693_v75_ = _dafny.CodePoint('u')
                d_2694_v76_: _dafny.Map
                d_2694_v76_ = _dafny.Map({d_2693_v75_: not (self.f16) or (self.f16)})
                d_2694_v76_ = (d_2694_v76_).set(d_2693_v75_, default__.fm38((self).f15, globalState))
                d_2695_v77_: _dafny.Map
                d_2695_v77_ = _dafny.Map({((_dafny.MultiSet([(self).f15])).set((self).f15, default__.abs((self).f15))) - (default__.fm10(((d_2692_v74_)[self.f16] if (self.f16) in (d_2692_v74_) else default__.fm22(self.f16, (0) - ((self).f15), globalState)), (self).f15, (self).f15, len(_dafny.SeqWithoutIsStrInference([d_2693_v75_ for d_2696_i9_ in range(default__.abs(944))])), globalState)): (0) - ((self).f15)})
                d_2697_v78_: _dafny.MultiSet
                d_2697_v78_ = _dafny.MultiSet([(self).f15, (self).f15, 986])
                d_2695_v77_ = (d_2695_v77_).set(d_2697_v78_, len(default__.fm31(globalState)))
                d_2698_v79_: _dafny.MultiSet
                d_2698_v79_ = _dafny.MultiSet([self.f16])
                d_2699_v80_: _dafny.Array
                nw443_ = _dafny.Array(False, 13)
                d_2699_v80_ = nw443_
                d_2700_v81_: _dafny.Map
                d_2700_v81_ = _dafny.Map({self.f16: d_2699_v80_})
                d_2701_v82_: _dafny.Map
                d_2701_v82_ = _dafny.Map({(self).f15: ((self).f15) <= ((self).f15)})
                index384_ = default__.safeIndex(484, (d_2699_v80_).length(0))
                (d_2699_v80_)[index384_] = self.f16
                index385_ = default__.safeIndex(484, (d_2699_v80_).length(0))
                rhs407_ = (d_2698_v79_ if not (self.f16) or (False) else d_2698_v79_)
                rhs408_ = _dafny.Map({((0) - ((self).f15)) > ((self).f15): d_2699_v80_})
                rhs409_ = d_2701_v82_
                rhs410_ = ((self).f15 if self.f16 else (self).f15)
                rhs411_ = (self.f16 if True else self.f16)
                lhs269_ = d_2699_v80_
                lhs270_ = default__.safeIndex(484, (d_2699_v80_).length(0))
                d_2698_v79_ = rhs407_
                d_2700_v81_ = rhs408_
                d_2701_v82_ = rhs409_
                r2 = rhs410_
                lhs269_[lhs270_] = rhs411_
                d_2702_v83_: _dafny.Array
                nw444_ = _dafny.Array(int(0), 9)
                d_2702_v83_ = nw444_
                d_2703_v84_: _dafny.Array
                nw445_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_2703_v84_ = nw445_
                d_2704_v85_: D20
                d_2704_v85_ = D20_DC46(d_2703_v84_)
                pat_let_tv55_ = d_2703_v84_
                index386_ = default__.safeIndex(484, (d_2699_v80_).length(0))
                def iife209_(_pat_let59_0):
                    def iife210_(d_2705_dt__update__tmp_h3_):
                        def iife211_(_pat_let60_0):
                            def iife212_(d_2706_dt__update_hcf64_h0_):
                                return D20_DC46(d_2706_dt__update_hcf64_h0_)
                            return iife212_(_pat_let60_0)
                        return iife211_(pat_let_tv55_)
                    return iife210_(_pat_let59_0)
                rhs412_ = self.f16
                rhs413_ = d_2702_v83_
                rhs414_ = d_2693_v75_
                rhs415_ = iife209_(D20_DC46(d_2703_v84_))
                rhs416_ = (d_2699_v80_)[default__.safeIndex(484, (d_2699_v80_).length(0))]
                lhs271_ = d_2699_v80_
                lhs272_ = default__.safeIndex(484, (d_2699_v80_).length(0))
                r1 = rhs412_
                d_2702_v83_ = rhs413_
                d_2693_v75_ = rhs414_
                d_2704_v85_ = rhs415_
                lhs271_[lhs272_] = rhs416_
            elif True:
                d_2707_v86_: _dafny.Array
                def lambda125_(d_2708_i10_):
                    return self.f16

                init70_ = lambda125_
                nw446_ = _dafny.Array(None, 17)
                for i0_70_ in range(nw446_.length(0)):
                    nw446_[i0_70_] = init70_(i0_70_)
                d_2707_v86_ = nw446_
                index387_ = default__.safeIndex(817, (d_2707_v86_).length(0))
                (d_2707_v86_)[index387_] = self.f16
                index388_ = default__.safeIndex(817, (d_2707_v86_).length(0))
                (d_2707_v86_)[index388_] = self.f16
                d_2709_v87_: _dafny.Array
                def lambda126_(d_2710_v86_):
                    def lambda127_(d_2711_i11_):
                        return _dafny.MultiSet([len(_dafny.Map({self.f16: (d_2710_v86_)[default__.safeIndex(817, (d_2710_v86_).length(0))]})), (self).f15])

                    return lambda127_

                init71_ = lambda126_(d_2707_v86_)
                nw447_ = _dafny.Array(None, 27)
                for i0_71_ in range(nw447_.length(0)):
                    nw447_[i0_71_] = init71_(i0_71_)
                d_2709_v87_ = nw447_
                d_2712_v88_: _dafny.MultiSet
                d_2712_v88_ = _dafny.MultiSet([d_2709_v87_])
                d_2713_v89_: _dafny.Array
                nw448_ = _dafny.Array(None, 5)
                nw448_[int(0)] = (self).f15
                nw448_[int(1)] = (self).f15
                nw448_[int(2)] = (self).f15
                nw448_[int(3)] = (self).f15
                nw448_[int(4)] = (self).f15
                d_2713_v89_ = nw448_
                d_2714_v90_: _dafny.MultiSet
                d_2714_v90_ = _dafny.MultiSet([d_2713_v89_])
                d_2715_v91_: D15
                d_2715_v91_ = D15_DC35(d_2714_v90_)
                d_2716_v92_: _dafny.Seq
                d_2716_v92_ = _dafny.SeqWithoutIsStrInference([((d_2715_v91_).cf47).cardinality, (self).f15])
                (globalState).f3 = (0) - (default__.safeDivisionInt(default__.safeModuloInt(((d_2712_v88_)[d_2709_v87_] if (d_2709_v87_) in (d_2712_v88_) else (d_2716_v92_)[default__.safeIndex(499, len(d_2716_v92_))]), (self).f15), (157) + ((self).f15)))
                r2 = (self).f15
                d_2717_v93_: str
                d_2717_v93_ = _dafny.CodePoint('w')
                d_2718_v94_: _dafny.Map
                d_2718_v94_ = _dafny.Map({d_2717_v93_: (d_2707_v86_)[default__.safeIndex(817, (d_2707_v86_).length(0))]})
                d_2719_v95_: _dafny.Seq
                d_2719_v95_ = _dafny.SeqWithoutIsStrInference([default__.fm63(self.f16, globalState), d_2718_v94_, d_2718_v94_])
                index389_ = default__.safeIndex(135, (d_2713_v89_).length(0))
                (d_2713_v89_)[index389_] = len((d_2719_v95_)[default__.safeIndex((self).f15, len(d_2719_v95_))])
                index390_ = default__.safeIndex(135, (d_2713_v89_).length(0))
                (d_2713_v89_)[index390_] = (self).f15
                d_2713_v89_ = d_2713_v89_
            d_2720_v96_: _dafny.Array
            nw449_ = _dafny.Array(int(0), 11)
            d_2720_v96_ = nw449_
            d_2720_v96_ = d_2720_v96_
        d_2721_v97_: _dafny.Seq
        d_2721_v97_ = _dafny.SeqWithoutIsStrInference([self.f16])
        d_2722_v98_: _dafny.Map
        d_2722_v98_ = _dafny.Map({(self).f15: d_2721_v97_})
        d_2723_v99_: D11
        d_2723_v99_ = D11_DC26()
        d_2724_v100_: D11
        d_2724_v100_ = D11_DC27(d_2723_v99_)
        d_2725_v101_: _dafny.Map
        d_2725_v101_ = _dafny.Map({d_2724_v100_: len(d_2721_v97_)})
        d_2726_v102_: _dafny.MultiSet
        d_2726_v102_ = _dafny.MultiSet([(self).f15])
        d_2727_v103_: D20
        d_2727_v103_ = D20_DC47(((self).f15) > (len(((d_2722_v98_)[820] if (820) in (d_2722_v98_) else d_2721_v97_))), len(d_2725_v101_), (d_2726_v102_).issubset(d_2726_v102_))
        source44_ = d_2727_v103_
        if source44_.is_DC47:
            d_2728___mcc_h5_ = source44_.cf65
            d_2729___mcc_h6_ = source44_.cf66
            d_2730___mcc_h7_ = source44_.cf67
            d_2731_cf67_ = d_2730___mcc_h7_
            d_2732_cf66_ = d_2729___mcc_h6_
            d_2733_cf65_ = d_2728___mcc_h5_
            d_2734_v104_: _dafny.Seq
            d_2734_v104_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eaacvu"))
            d_2735_v105_: _dafny.Seq
            d_2735_v105_ = _dafny.SeqWithoutIsStrInference([(self).f15, (self).f15])
            d_2736_v106_: _dafny.Seq
            d_2736_v106_ = _dafny.SeqWithoutIsStrInference([d_2735_v105_])
            d_2737_v108_: _dafny.Map
            d_2737_v108_ = _dafny.Map({len(d_2734_v104_): (self).f15})
            d_2738_v109_: _dafny.Map
            def iife213_():
                coll91_ = _dafny.Map()
                compr_91_: int
                for compr_91_ in (_dafny.Map({d_2732_cf66_: (self).f15})).keys.Elements:
                    d_2739_v107_: int = compr_91_
                    if (d_2739_v107_) in (_dafny.Map({d_2732_cf66_: (self).f15})):
                        coll91_[default__.safeDivisionInt(d_2739_v107_, (self).f15)] = d_2732_cf66_
                return _dafny.Map(coll91_)
            d_2738_v109_ = _dafny.Map({iife213_()
            : d_2737_v108_})
            d_2740_v110_: C7
            nw450_ = C7()
            nw450_.ctor__(d_2734_v104_, d_2736_v106_, (self).f15, d_2738_v109_)
            d_2740_v110_ = nw450_
            d_2741_v111_: _dafny.Map
            d_2741_v111_ = _dafny.Map({(d_2721_v97_)[default__.safeIndex(d_2732_cf66_, len(d_2721_v97_))]: d_2740_v110_})
            d_2742_v112_: _dafny.Map
            d_2742_v112_ = _dafny.Map({len(d_2741_v111_): (self).f15})
            d_2743_v113_: _dafny.Seq
            d_2743_v113_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({self.f16: d_2731_cf67_})])
            d_2737_v108_ = (d_2737_v108_).set(295, len(d_2743_v113_))
            d_2731_cf67_ = not(d_2731_cf67_)
            d_2744_v114_: _dafny.Array
            nw451_ = _dafny.Array(_dafny.Array(None, 0), 8)
            d_2744_v114_ = nw451_
            d_2745_v115_: _dafny.Array
            nw452_ = _dafny.Array(False, 3)
            d_2745_v115_ = nw452_
            index391_ = default__.safeIndex(789, (d_2744_v114_).length(0))
            (d_2744_v114_)[index391_] = d_2745_v115_
            d_2746_v116_: _dafny.Array
            nw453_ = _dafny.Array(int(0), 7)
            d_2746_v116_ = nw453_
            d_2747_v117_: _dafny.Map
            d_2747_v117_ = _dafny.Map({d_2746_v116_: d_2745_v115_})
            d_2748_v118_: _dafny.Map
            d_2748_v118_ = _dafny.Map({(d_2734_v104_) < (d_2734_v104_): d_2745_v115_})
            index392_ = default__.safeIndex(789, (d_2744_v114_).length(0))
            rhs417_ = ((d_2748_v118_)[(False) == (d_2731_cf67_)] if ((False) == (d_2731_cf67_)) in (d_2748_v118_) else d_2745_v115_)
            rhs418_ = ((d_2747_v117_) | (d_2747_v117_)).set(d_2746_v116_, d_2745_v115_)
            rhs419_ = d_2745_v115_
            lhs273_ = d_2744_v114_
            lhs274_ = default__.safeIndex(789, (d_2744_v114_).length(0))
            lhs273_[lhs274_] = rhs417_
            d_2747_v117_ = rhs418_
            d_2745_v115_ = rhs419_
            (self).f16 = d_2733_cf65_
        elif source44_.is_DC48:
            d_2749___mcc_h8_ = source44_.cf68
            d_2750___mcc_h9_ = source44_.cf69
            d_2751___mcc_h10_ = source44_.cf70
            d_2752_cf70_ = d_2751___mcc_h10_
            d_2753_cf69_ = d_2750___mcc_h9_
            d_2754_cf68_ = d_2749___mcc_h8_
            d_2754_cf68_ = False
            (globalState).f3 = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlffyw"))), (self).f15)
            d_2755_v119_: _dafny.Map
            d_2755_v119_ = _dafny.Map({282: (self).f15})
            d_2756_v120_: _dafny.Map
            d_2756_v120_ = _dafny.Map({d_2755_v119_: d_2755_v119_})
            d_2757_v122_: _dafny.Seq
            d_2757_v122_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f15 for d_2758_i12_ in range(default__.abs(882))]))])
            d_2759_v123_: _dafny.Seq
            d_2759_v123_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbkk"))
            d_2760_v124_: _dafny.Seq
            d_2760_v124_ = _dafny.SeqWithoutIsStrInference([d_2757_v122_])
            d_2761_v125_: C2
            nw454_ = C2()
            def iife214_():
                coll92_ = _dafny.Map()
                compr_92_: int
                for compr_92_ in (d_2757_v122_).Elements:
                    d_2762_v121_: int = compr_92_
                    if (d_2762_v121_) in (d_2757_v122_):
                        coll92_[default__.safeModuloInt(d_2762_v121_, (self).f15)] = (_dafny.MultiSet([(d_2721_v97_)[default__.safeIndex(787, len(d_2721_v97_))], self.f16])).cardinality
                return _dafny.Map(coll92_)
            nw454_.ctor__((0) - ((self).f15), ((d_2756_v120_).set(_dafny.Map({652: (self).f15}), d_2755_v119_)) | (_dafny.Map({d_2755_v119_: iife214_()
            })), d_2759_v123_, d_2760_v124_)
            d_2761_v125_ = nw454_
            d_2763_v126_: _dafny.Map
            d_2763_v126_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ay")): (self).f15})
            d_2764_v127_: int
            d_2765_v128_: _dafny.Map
            d_2766_v129_: bool
            d_2767_v130_: int
            out68_: int
            out69_: _dafny.Map
            out70_: bool
            out71_: int
            out68_, out69_, out70_, out71_ = (d_2761_v125_).m2(len(d_2763_v126_), globalState)
            d_2764_v127_ = out68_
            d_2765_v128_ = out69_
            d_2766_v129_ = out70_
            d_2767_v130_ = out71_
        elif True:
            d_2768___mcc_h11_ = source44_.cf64
            d_2769_cf64_ = d_2768___mcc_h11_
            r1 = self.f16
            (self).f16 = self.f16
            if self.f16:
                d_2770_v131_: _dafny.Array
                nw455_ = _dafny.Array(D12.default()(), 11)
                d_2770_v131_ = nw455_
                d_2771_v132_: _dafny.Seq
                d_2771_v132_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ot"))
                d_2772_v133_: _dafny.Seq
                d_2772_v133_ = _dafny.SeqWithoutIsStrInference([(self).f15, (self).f15])
                d_2773_v134_: _dafny.Map
                d_2773_v134_ = _dafny.Map({self.f16: d_2772_v133_})
                d_2774_v135_: D12
                d_2774_v135_ = D12_DC29(self.f16, len(d_2771_v132_), ((d_2773_v134_)[self.f16] if (self.f16) in (d_2773_v134_) else _dafny.SeqWithoutIsStrInference([(0) - ((self).f15) for d_2775_i13_ in range(default__.abs(308))])), self.f16, (self).f15)
                index393_ = default__.safeIndex(575, (d_2770_v131_).length(0))
                (d_2770_v131_)[index393_] = D12_DC30(d_2774_v135_)
                d_2776_v136_: _dafny.Array
                nw456_ = _dafny.Array(_dafny.Set({}), 18)
                d_2776_v136_ = nw456_
                index394_ = default__.safeIndex(563, (d_2776_v136_).length(0))
                (d_2776_v136_)[index394_] = _dafny.Set({self.f16, (d_2721_v97_)[default__.safeIndex(586, len(d_2721_v97_))], self.f16, self.f16})
                d_2777_v137_: D12
                d_2777_v137_ = D12_DC30(d_2774_v135_)
                d_2778_v138_: _dafny.Set
                d_2778_v138_ = _dafny.Set({self.f16})
                index395_ = default__.safeIndex(575, (d_2770_v131_).length(0))
                index396_ = default__.safeIndex(563, (d_2776_v136_).length(0))
                rhs420_ = self.f16
                rhs421_ = ((self).f15) * (((self).f15) * ((self).f15))
                rhs422_ = d_2777_v137_
                rhs423_ = d_2772_v133_
                rhs424_ = d_2778_v138_
                lhs275_ = self
                lhs276_ = globalState
                lhs277_ = d_2770_v131_
                lhs278_ = default__.safeIndex(575, (d_2770_v131_).length(0))
                lhs279_ = d_2776_v136_
                lhs280_ = default__.safeIndex(563, (d_2776_v136_).length(0))
                lhs275_.f16 = rhs420_
                lhs276_.f3 = rhs421_
                lhs277_[lhs278_] = rhs422_
                d_2772_v133_ = rhs423_
                lhs279_[lhs280_] = rhs424_
                d_2779_v139_: str
                d_2779_v139_ = _dafny.CodePoint('j')
                d_2780_v140_: _dafny.Map
                d_2780_v140_ = _dafny.Map({d_2771_v132_: d_2779_v139_})
                d_2780_v140_ = (d_2780_v140_).set((d_2771_v132_ if (d_2721_v97_)[default__.safeIndex((self).f15, len(d_2721_v97_))] else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "un"))), d_2779_v139_)
                r1 = not(not((d_2779_v139_) in ((d_2771_v132_).set(default__.safeIndex(len(d_2771_v132_), len(d_2771_v132_)), d_2779_v139_))))
                d_2781_v141_: _dafny.Set
                d_2781_v141_ = _dafny.Set({(self).f15, (self).f15})
                (globalState).f3 = len((d_2781_v141_) - (_dafny.Set({(self).f15, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "myncocmov")))})))
                d_2782_v142_: _dafny.Array
                nw457_ = _dafny.Array(_dafny.Map({}), 5)
                d_2782_v142_ = nw457_
                d_2783_v143_: D12
                d_2783_v143_ = D12_DC28(d_2772_v133_)
                d_2784_v144_: _dafny.Map
                d_2784_v144_ = _dafny.Map({(0) - ((self).f15): d_2783_v143_})
                d_2785_v145_: _dafny.Map
                d_2785_v145_ = _dafny.Map({(self).f15: d_2784_v144_})
                index397_ = default__.safeIndex(313, (d_2782_v142_).length(0))
                (d_2782_v142_)[index397_] = d_2785_v145_
                d_2786_v146_: _dafny.MultiSet
                d_2786_v146_ = _dafny.MultiSet([self.f16, self.f16, True, self.f16, self.f16])
                d_2787_v147_: D18
                d_2787_v147_ = D18_DC43(default__.fm22(self.f16, (d_2786_v146_).cardinality, globalState), (self).f15)
                d_2788_v148_: D18
                d_2788_v148_ = D18_DC41(d_2786_v146_)
                d_2789_v149_: _dafny.Map
                d_2789_v149_ = _dafny.Map({((d_2787_v147_).cf60 if self.f16 else default__.fm38(-159, globalState)): (d_2786_v146_).intersection((d_2788_v148_).cf57)})
                d_2790_v150_: _dafny.Map
                d_2790_v150_ = _dafny.Map({self.f16: d_2726_v102_})
                index398_ = default__.safeIndex(313, (d_2782_v142_).length(0))
                rhs425_ = self.f16
                rhs426_ = d_2785_v145_
                rhs427_ = d_2789_v149_
                rhs428_ = ((d_2726_v102_) | (d_2726_v102_)).intersection(((d_2790_v150_)[self.f16] if (self.f16) in (d_2790_v150_) else d_2726_v102_))
                rhs429_ = self.f16
                lhs281_ = d_2782_v142_
                lhs282_ = default__.safeIndex(313, (d_2782_v142_).length(0))
                r1 = rhs425_
                lhs281_[lhs282_] = rhs426_
                d_2789_v149_ = rhs427_
                d_2726_v102_ = rhs428_
                r1 = rhs429_
            elif True:
                d_2791_v151_: _dafny.Map
                d_2791_v151_ = _dafny.Map({(self).f15: True})
                d_2792_v152_: _dafny.Set
                d_2792_v152_ = _dafny.Set({(self).f15, (self).f15})
                d_2793_v153_: _dafny.Map
                d_2793_v153_ = _dafny.Map({(d_2721_v97_)[default__.safeIndex(len(d_2791_v151_), len(d_2721_v97_))]: d_2792_v152_})
                d_2793_v153_ = d_2793_v153_
                d_2794_v154_: _dafny.Array
                def lambda128_(d_2795_i14_):
                    return _dafny.Map({self.f16: self.f16})

                init72_ = lambda128_
                nw458_ = _dafny.Array(None, 14)
                for i0_72_ in range(nw458_.length(0)):
                    nw458_[i0_72_] = init72_(i0_72_)
                d_2794_v154_ = nw458_
                d_2796_v155_: _dafny.Map
                d_2796_v155_ = _dafny.Map({self.f16: True})
                index399_ = default__.safeIndex(733, (d_2794_v154_).length(0))
                (d_2794_v154_)[index399_] = d_2796_v155_
                d_2797_v156_: _dafny.Seq
                d_2797_v156_ = _dafny.SeqWithoutIsStrInference([d_2796_v155_, d_2796_v155_, d_2796_v155_, d_2796_v155_])
                index400_ = default__.safeIndex(733, (d_2794_v154_).length(0))
                (d_2794_v154_)[index400_] = (d_2797_v156_)[default__.safeIndex((0) - ((self).f15), len(d_2797_v156_))]
                (self).f16 = self.f16
                d_2798_v157_: _dafny.Seq
                d_2798_v157_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                d_2799_v158_: D18
                d_2799_v158_ = D18_DC42(self.f16, (self).f15)
                d_2800_v159_: _dafny.Set
                d_2800_v159_ = _dafny.Set({_dafny.CodePoint('x')})
                d_2801_v160_: _dafny.Array
                nw459_ = _dafny.Array(None, 28)
                nw459_[int(0)] = self.f16
                nw459_[int(1)] = self.f16
                nw459_[int(2)] = self.f16
                nw459_[int(3)] = ((self).f15) == (-566)
                nw459_[int(4)] = ((self).f15) != (len(d_2798_v157_))
                nw459_[int(5)] = not(not(self.f16))
                nw459_[int(6)] = (d_2799_v158_).cf58
                nw459_[int(7)] = self.f16
                nw459_[int(8)] = self.f16
                nw459_[int(9)] = (not(self.f16)) or (self.f16)
                nw459_[int(10)] = self.f16
                nw459_[int(11)] = self.f16
                nw459_[int(12)] = self.f16
                nw459_[int(13)] = self.f16
                nw459_[int(14)] = self.f16
                nw459_[int(15)] = self.f16
                nw459_[int(16)] = self.f16
                nw459_[int(17)] = self.f16
                nw459_[int(18)] = self.f16
                nw459_[int(19)] = self.f16
                nw459_[int(20)] = (d_2800_v159_).ispropersubset((D21_DC49(d_2800_v159_)).cf71)
                nw459_[int(21)] = self.f16
                nw459_[int(22)] = not(self.f16)
                nw459_[int(23)] = self.f16
                nw459_[int(24)] = self.f16
                nw459_[int(25)] = self.f16
                nw459_[int(26)] = (d_2726_v102_).issubset(d_2726_v102_)
                nw459_[int(27)] = True
                d_2801_v160_ = nw459_
                index401_ = default__.safeIndex(551, (d_2801_v160_).length(0))
                (d_2801_v160_)[index401_] = False
                index402_ = default__.safeIndex(551, (d_2801_v160_).length(0))
                (d_2801_v160_)[index402_] = self.f16
                d_2802_v161_: _dafny.Map
                d_2802_v161_ = _dafny.Map({self.f16: (self).f15})
                d_2802_v161_ = (d_2802_v161_).set(self.f16, 596)
            d_2803_v162_: _dafny.MultiSet
            d_2803_v162_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxvac")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylcoxs"))])
            d_2804_v163_: D10
            d_2804_v163_ = D10_DC22(d_2803_v162_)
            source45_ = d_2804_v163_
            if source45_.is_DC23:
                d_2805___mcc_h12_ = source45_.cf26
                d_2806___mcc_h13_ = source45_.cf27
                d_2807_cf27_ = d_2806___mcc_h13_
                d_2808_cf26_ = d_2805___mcc_h12_
                d_2809_v164_: _dafny.Seq
                d_2809_v164_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ac"))
                d_2810_v165_: _dafny.Map
                d_2810_v165_ = _dafny.Map({self.f16: self.f16})
                d_2811_v166_: _dafny.Seq
                d_2811_v166_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f15 for d_2812_i15_ in range(default__.abs(516))])), len(d_2810_v165_)])
                d_2813_v167_: _dafny.Seq
                d_2813_v167_ = _dafny.SeqWithoutIsStrInference([d_2811_v166_])
                d_2814_v168_: C4
                nw460_ = C4()
                nw460_.ctor__(self.f16, default__.fm38(218, globalState), d_2809_v164_, d_2813_v167_)
                d_2814_v168_ = nw460_
                d_2815_v169_: _dafny.Set
                d_2815_v169_ = _dafny.Set({d_2814_v168_})
                d_2816_v170_: _dafny.Map
                d_2816_v170_ = _dafny.Map({(self).f15: (self).f15})
                d_2817_v171_: _dafny.Map
                d_2817_v171_ = _dafny.Map({d_2816_v170_: (_dafny.Map({(self).f15: (self).f15})).set((self).f15, (self).f15)})
                d_2818_v172_: C3
                nw461_ = C3()
                nw461_.ctor__(self.f16, len(d_2815_v169_), (self).f15, d_2817_v171_, d_2809_v164_, (default__.fm21(d_2809_v164_, globalState)) + (d_2813_v167_))
                d_2818_v172_ = nw461_
                r1 = self.f16
                d_2819_v173_: _dafny.Array
                nw462_ = _dafny.Array(False, 25)
                d_2819_v173_ = nw462_
                index403_ = default__.safeIndex(102, (d_2819_v173_).length(0))
                (d_2819_v173_)[index403_] = False
                index404_ = default__.safeIndex(102, (d_2819_v173_).length(0))
                (d_2819_v173_)[index404_] = d_2814_v168_.f20
                d_2820_v174_: bool
                d_2821_v175_: bool
                d_2822_v176_: _dafny.Array
                d_2823_v177_: int
                out72_: bool
                out73_: bool
                out74_: _dafny.Array
                out75_: int
                out72_, out73_, out74_, out75_ = (d_2818_v172_).m1((self).f15, (d_2818_v172_).f22, (0) - ((d_2818_v172_).f22), globalState)
                d_2820_v174_ = out72_
                d_2821_v175_ = out73_
                d_2822_v176_ = out74_
                d_2823_v177_ = out75_
            elif source45_.is_DC22:
                d_2824___mcc_h14_ = source45_.cf25
                d_2825_cf25_ = d_2824___mcc_h14_
                (self).f16 = self.f16
                (self).f16 = (806) < ((self).f15)
                d_2826_v178_: D11
                d_2826_v178_ = D11_DC26()
                d_2827_v179_: D12
                d_2827_v179_ = D12_DC30(default__.fm64(d_2826_v178_, globalState))
                d_2828_v180_: _dafny.MultiSet
                d_2828_v180_ = _dafny.MultiSet([True, self.f16, self.f16])
                d_2829_v181_: _dafny.Seq
                d_2829_v181_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dygg"))
                d_2830_v182_: D12
                d_2830_v182_ = D12_DC29(self.f16, ((d_2828_v180_)[self.f16] if (self.f16) in (d_2828_v180_) else len(d_2829_v181_)), _dafny.SeqWithoutIsStrInference([693]), not(self.f16), (self).f15)
                d_2827_v179_ = D12_DC30(d_2830_v182_)
                d_2831_v183_: _dafny.Map
                d_2831_v183_ = _dafny.Map({(0) - ((0) - ((self).f15)): (self).f15})
                d_2832_v184_: _dafny.Map
                d_2832_v184_ = _dafny.Map({(self).f15: self.f16})
                d_2831_v183_ = ((d_2831_v183_).set((self).f15, len(d_2832_v184_))) | (_dafny.Map({219: (self).f15}))
            elif True:
                d_2833___mcc_h15_ = source45_.cf28
                d_2834_cf28_ = d_2833___mcc_h15_
                d_2835_v185_: _dafny.Seq
                d_2835_v185_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uu"))
                d_2836_v186_: _dafny.Seq
                d_2836_v186_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                d_2835_v185_ = default__.fm26(((d_2836_v186_)[default__.safeIndex((self).f15, len(d_2836_v186_))]) * ((self).f15), (_dafny.Map({(self).f15: d_2835_v185_})) | (default__.fm65((self).f15, (_dafny.MultiSet([(self).f15, (self).f15, (self).f15, (self).f15])).cardinality, self.f16, self.f16, globalState)), globalState)
                r2 = (((self).f15) + ((self).f15)) * (745)
                d_2837_v187_: D2
                d_2837_v187_ = D2_DC4((_dafny.MultiSet(d_2721_v97_)).cardinality)
                d_2838_v188_: _dafny.Set
                d_2838_v188_ = _dafny.Set({self.f16, self.f16})
                d_2839_v189_: _dafny.Map
                d_2839_v189_ = _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hepep")))})
                d_2840_v190_: _dafny.Map
                d_2840_v190_ = _dafny.Map({(self).f15: self.f16})
                d_2841_v191_: D12
                d_2841_v191_ = D12_DC29(self.f16, len(d_2840_v190_), d_2836_v186_, self.f16, (self).f15)
                d_2842_v192_: _dafny.Seq
                d_2842_v192_ = _dafny.SeqWithoutIsStrInference([d_2838_v188_, _dafny.Set({self.f16, self.f16})])
                d_2843_v193_: _dafny.Array
                nw463_ = _dafny.Array(None, 20)
                nw463_[int(0)] = (self.f16 if self.f16 else self.f16)
                nw463_[int(1)] = ((self).f15) < ((0) - ((d_2837_v187_).cf3))
                nw463_[int(2)] = self.f16
                nw463_[int(3)] = self.f16
                nw463_[int(4)] = self.f16
                nw463_[int(5)] = ((d_2721_v97_)[default__.safeIndex((self).f15, len(d_2721_v97_))]) or (self.f16)
                nw463_[int(6)] = (d_2838_v188_).issubset(d_2838_v188_)
                nw463_[int(7)] = self.f16
                nw463_[int(8)] = self.f16
                nw463_[int(9)] = (d_2721_v97_)[default__.safeIndex((self).f15, len(d_2721_v97_))]
                nw463_[int(10)] = (728) >= ((self).f15)
                nw463_[int(11)] = default__.fm38(len(d_2839_v189_), globalState)
                nw463_[int(12)] = self.f16
                nw463_[int(13)] = True
                nw463_[int(14)] = self.f16
                nw463_[int(15)] = (d_2841_v191_).cf32
                nw463_[int(16)] = not(not(self.f16))
                nw463_[int(17)] = (d_2726_v102_).issubset(_dafny.MultiSet([(self).f15]))
                nw463_[int(18)] = (self.f16) in ((d_2842_v192_)[default__.safeIndex((0) - (default__.fm0(476, (self).f15, globalState)), len(d_2842_v192_))])
                nw463_[int(19)] = self.f16
                d_2843_v193_ = nw463_
                index405_ = default__.safeIndex(960, (d_2843_v193_).length(0))
                (d_2843_v193_)[index405_] = self.f16
                d_2844_v194_: D17
                d_2844_v194_ = D17_DC40(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyeonp")), self.f16, 70)
                index406_ = default__.safeIndex(960, (d_2843_v193_).length(0))
                (d_2843_v193_)[index406_] = not(default__.fm13(self.f16, (d_2726_v102_).cardinality, (d_2835_v185_) + ((d_2844_v194_).cf54), globalState))
                d_2845_v195_: _dafny.Array
                nw464_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_2845_v195_ = nw464_
                d_2846_v196_: _dafny.Array
                def lambda129_(d_2847_i16_):
                    return default__.safeModuloInt(d_2847_i16_, (self).f15)

                init73_ = lambda129_
                nw465_ = _dafny.Array(None, 28)
                for i0_73_ in range(nw465_.length(0)):
                    nw465_[i0_73_] = init73_(i0_73_)
                d_2846_v196_ = nw465_
                index407_ = default__.safeIndex(481, (d_2845_v195_).length(0))
                (d_2845_v195_)[index407_] = d_2846_v196_
                index408_ = default__.safeIndex(481, (d_2845_v195_).length(0))
                (d_2845_v195_)[index408_] = (d_2846_v196_ if (d_2843_v193_)[default__.safeIndex(960, (d_2843_v193_).length(0))] else d_2846_v196_)
        d_2848_v197_: _dafny.Array
        nw466_ = _dafny.Array(int(0), 20)
        d_2848_v197_ = nw466_
        index409_ = default__.safeIndex(136, (d_2848_v197_).length(0))
        (d_2848_v197_)[index409_] = (self).f15
        index410_ = default__.safeIndex(136, (d_2848_v197_).length(0))
        (d_2848_v197_)[index410_] = (self).f15
        d_2849_v198_: _dafny.Seq
        d_2849_v198_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqgss"))
        hi9_ = (default__.fm0(len(d_2849_v198_), (self).f15, globalState)) * ((self).f15)
        for d_2850_i17_ in range(((d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))]) * (default__.fm0((d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))], (self).f15, globalState)), hi9_):
            d_2851_v199_: _dafny.Array
            def lambda130_(d_2852_i18_):
                return self.f16

            init74_ = lambda130_
            nw467_ = _dafny.Array(None, 27)
            for i0_74_ in range(nw467_.length(0)):
                nw467_[i0_74_] = init74_(i0_74_)
            d_2851_v199_ = nw467_
            index411_ = default__.safeIndex(503, (d_2851_v199_).length(0))
            (d_2851_v199_)[index411_] = not (False) or (self.f16)
            index412_ = default__.safeIndex(503, (d_2851_v199_).length(0))
            (d_2851_v199_)[index412_] = False
            d_2853_v200_: _dafny.Map
            d_2853_v200_ = _dafny.Map({(d_2851_v199_)[default__.safeIndex(503, (d_2851_v199_).length(0))]: (self).f15})
            d_2854_v201_: D17
            d_2854_v201_ = D17_DC39(((d_2853_v200_).set(self.f16, (self).f15)) | (d_2853_v200_))
            source46_ = d_2854_v201_
            if source46_.is_DC40:
                d_2855___mcc_h16_ = source46_.cf54
                d_2856___mcc_h17_ = source46_.cf55
                d_2857___mcc_h18_ = source46_.cf56
                d_2858_cf56_ = d_2857___mcc_h18_
                d_2859_cf55_ = d_2856___mcc_h17_
                d_2860_cf54_ = d_2855___mcc_h16_
                d_2861_v202_: _dafny.Array
                nw468_ = _dafny.Array(_dafny.Array(None, 0), 23)
                d_2861_v202_ = nw468_
                d_2862_v203_: _dafny.Array
                nw469_ = _dafny.Array(None, 3)
                nw469_[int(0)] = d_2721_v97_
                nw469_[int(1)] = default__.fm2(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2863_i19_ in range(default__.abs(280))]), len(d_2860_cf54_), globalState)
                nw469_[int(2)] = d_2721_v97_
                d_2862_v203_ = nw469_
                index413_ = default__.safeIndex(753, (d_2861_v202_).length(0))
                (d_2861_v202_)[index413_] = d_2862_v203_
                index414_ = default__.safeIndex(753, (d_2861_v202_).length(0))
                (d_2861_v202_)[index414_] = d_2862_v203_
                d_2864_v204_: _dafny.Set
                d_2864_v204_ = _dafny.Set({(d_2851_v199_)[default__.safeIndex(503, (d_2851_v199_).length(0))]})
                d_2865_v205_: _dafny.Map
                d_2865_v205_ = _dafny.Map({(default__.fm66(globalState)).cf1: (d_2864_v204_).isdisjoint(_dafny.Set({True, self.f16}))})
                d_2865_v205_ = (d_2865_v205_).set(self.f16, (d_2859_cf55_) == (self.f16))
                d_2866_v206_: bool
                d_2867_v207_: bool
                d_2868_v208_: _dafny.Set
                out76_: bool
                out77_: bool
                out78_: _dafny.Set
                out76_, out77_, out78_ = (self).m7(globalState)
                d_2866_v206_ = out76_
                d_2867_v207_ = out77_
                d_2868_v208_ = out78_
                d_2869_v209_: _dafny.Array
                def lambda131_(d_2870_cf56_, d_2871_i17_, d_2872_v205_, d_2873_v207_):
                    def lambda132_(d_2874_i20_):
                        return (_dafny.SeqWithoutIsStrInference([d_2870_cf56_, len(_dafny.Map({d_2871_i17_: (self).f15})), len(d_2872_v205_), d_2870_cf56_, (_dafny.MultiSet([d_2873_v207_])).cardinality])) + (_dafny.SeqWithoutIsStrInference([-944]))

                    return lambda132_

                init75_ = lambda131_(d_2858_cf56_, d_2850_i17_, d_2865_v205_, d_2867_v207_)
                nw470_ = _dafny.Array(None, 17)
                for i0_75_ in range(nw470_.length(0)):
                    nw470_[i0_75_] = init75_(i0_75_)
                d_2869_v209_ = nw470_
                d_2875_v210_: _dafny.Seq
                d_2875_v210_ = _dafny.SeqWithoutIsStrInference([d_2858_cf56_, (d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))]])
                index415_ = default__.safeIndex(191, (d_2869_v209_).length(0))
                (d_2869_v209_)[index415_] = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(d_2859_cf55_)])).cardinality, len(d_2875_v210_)])
                index416_ = default__.safeIndex(191, (d_2869_v209_).length(0))
                (d_2869_v209_)[index416_] = (_dafny.SeqWithoutIsStrInference([(d_2875_v210_)[default__.safeIndex((d_2875_v210_)[default__.safeIndex(d_2850_i17_, len(d_2875_v210_))], len(d_2875_v210_))], (self).f15])).set(default__.safeIndex(d_2858_cf56_, len(_dafny.SeqWithoutIsStrInference([(d_2875_v210_)[default__.safeIndex((d_2875_v210_)[default__.safeIndex(d_2850_i17_, len(d_2875_v210_))], len(d_2875_v210_))], (self).f15]))), -800)
            elif True:
                d_2876___mcc_h19_ = source46_.cf53
                d_2877_cf53_ = d_2876___mcc_h19_
                d_2878_v211_: _dafny.Map
                d_2878_v211_ = _dafny.Map({(d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))]: self.f16})
                d_2879_v212_: str
                d_2879_v212_ = _dafny.CodePoint('w')
                d_2880_v213_: _dafny.Seq
                d_2880_v213_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(((d_2726_v102_)[len(d_2878_v211_)] if (len(d_2878_v211_)) in (d_2726_v102_) else len((d_2849_v198_).set(default__.safeIndex((d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))], len(d_2849_v198_)), d_2879_v212_))), (d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))]), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))), default__.safeModuloInt(d_2850_i17_, d_2850_i17_), (self).f15, default__.safeDivisionInt((0) - (len(d_2721_v97_)), 996)])
                d_2881_v214_: _dafny.Map
                d_2881_v214_ = _dafny.Map({(self).f15: (d_2726_v102_).cardinality})
                d_2882_v215_: _dafny.Map
                d_2882_v215_ = _dafny.Map({d_2881_v214_: d_2881_v214_})
                d_2883_v216_: T0
                nw471_ = C6()
                nw471_.ctor__((d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))], d_2882_v215_, d_2849_v198_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f15, (self).f15]) for d_2884_i21_ in range(default__.abs(467))]))
                d_2883_v216_ = nw471_
                d_2885_v217_: _dafny.Map
                d_2885_v217_ = _dafny.Map({self.f16: d_2883_v216_})
                d_2886_v218_: _dafny.Set
                d_2886_v218_ = _dafny.Set({default__.fm27(globalState), d_2879_v212_})
                d_2887_v219_: _dafny.Seq
                d_2887_v219_ = _dafny.SeqWithoutIsStrInference([d_2886_v218_, _dafny.Set({d_2879_v212_}), d_2886_v218_, d_2886_v218_])
                d_2880_v213_ = (_dafny.SeqWithoutIsStrInference([default__.safeModuloInt((self).f15, (0) - ((0) - (len(d_2885_v217_)))), d_2850_i17_, (0) - (d_2850_i17_)])).set(default__.safeIndex(len((d_2887_v219_) + (d_2887_v219_)), len(_dafny.SeqWithoutIsStrInference([default__.safeModuloInt((self).f15, (0) - ((0) - (len(d_2885_v217_)))), d_2850_i17_, (0) - (d_2850_i17_)]))), default__.fm0((d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))], (d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))], globalState))
                index417_ = default__.safeIndex(503, (d_2851_v199_).length(0))
                (d_2851_v199_)[index417_] = (d_2851_v199_)[default__.safeIndex(503, (d_2851_v199_).length(0))]
                d_2888_v220_: _dafny.Map
                d_2888_v220_ = _dafny.Map({d_2881_v214_: d_2880_v213_})
                d_2889_v222_: _dafny.Map
                d_2889_v222_ = _dafny.Map({d_2881_v214_: (d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))]})
                def iife215_():
                    coll93_ = _dafny.Map()
                    compr_93_: _dafny.Map
                    for compr_93_ in (d_2889_v222_).keys.Elements:
                        d_2890_v221_: _dafny.Map = compr_93_
                        if (d_2890_v221_) in (d_2889_v222_):
                            coll93_[d_2890_v221_] = (d_2880_v213_).set(default__.safeIndex((d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))], len(d_2880_v213_)), d_2850_i17_)
                    return _dafny.Map(coll93_)
                d_2888_v220_ = (iife215_()
                 if not (not(self.f16)) or (self.f16) else d_2888_v220_)
                r2 = default__.safeDivisionInt((d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))], 528)
            d_2891_v223_: str
            d_2891_v223_ = _dafny.CodePoint('o')
            d_2892_v224_: _dafny.Seq
            d_2892_v224_ = _dafny.SeqWithoutIsStrInference([(default__.fm53(globalState)).cardinality])
            d_2893_v225_: _dafny.Map
            d_2893_v225_ = _dafny.Map({(d_2851_v199_)[default__.safeIndex(503, (d_2851_v199_).length(0))]: d_2892_v224_})
            index418_ = default__.safeIndex(503, (d_2851_v199_).length(0))
            rhs430_ = d_2891_v223_
            rhs431_ = self.f16
            rhs432_ = d_2849_v198_
            rhs433_ = d_2893_v225_
            rhs434_ = d_2848_v197_
            lhs283_ = d_2851_v199_
            lhs284_ = default__.safeIndex(503, (d_2851_v199_).length(0))
            d_2891_v223_ = rhs430_
            lhs283_[lhs284_] = rhs431_
            d_2849_v198_ = rhs432_
            d_2893_v225_ = rhs433_
            d_2848_v197_ = rhs434_
            d_2894_v226_: _dafny.MultiSet
            d_2894_v226_ = _dafny.MultiSet([not(True)])
            d_2895_v227_: _dafny.Map
            d_2895_v227_ = _dafny.Map({(d_2851_v199_)[default__.safeIndex(503, (d_2851_v199_).length(0))]: d_2894_v226_})
            d_2896_v228_: _dafny.MultiSet
            d_2896_v228_ = _dafny.MultiSet([_dafny.CodePoint('n'), d_2891_v223_])
            d_2897_v229_: _dafny.Map
            d_2897_v229_ = _dafny.Map({(d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))]: (d_2896_v228_).cardinality})
            d_2898_v230_: _dafny.Map
            d_2898_v230_ = _dafny.Map({(d_2897_v229_).set((d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))], (self).f15): d_2897_v229_})
            d_2899_v231_: _dafny.Map
            d_2899_v231_ = _dafny.Map({(_dafny.MultiSet(d_2849_v198_)).cardinality: _dafny.SeqWithoutIsStrInference([d_2891_v223_ for d_2900_i22_ in range(default__.abs(-729))])})
            d_2901_v232_: _dafny.Seq
            d_2901_v232_ = _dafny.SeqWithoutIsStrInference([d_2892_v224_, _dafny.SeqWithoutIsStrInference([-176]), d_2892_v224_, d_2892_v224_, d_2892_v224_])
            d_2902_v233_: C2
            nw472_ = C2()
            nw472_.ctor__(len(d_2895_v227_), d_2898_v230_, (((d_2899_v231_)[(0) - (d_2850_i17_)] if ((0) - (d_2850_i17_)) in (d_2899_v231_) else (_dafny.SeqWithoutIsStrInference([d_2891_v223_ for d_2903_i23_ in range(default__.abs(127))])).set(default__.safeIndex(d_2850_i17_, len(_dafny.SeqWithoutIsStrInference([d_2891_v223_ for d_2904_i23_ in range(default__.abs(127))]))), d_2891_v223_))) + (d_2849_v198_), d_2901_v232_)
            d_2902_v233_ = nw472_
            d_2902_v233_ = d_2902_v233_
        d_2848_v197_ = d_2848_v197_
        r2 = (d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))]
        d_2905_v234_: _dafny.Map
        d_2905_v234_ = _dafny.Map({self.f16: default__.fm0((self).f15, (d_2848_v197_)[default__.safeIndex(136, (d_2848_v197_).length(0))], globalState)})
        r0 = (default__.fm18(globalState)) | (d_2905_v234_)
        r1 = self.f16
        r2 = (self).f15
        return r0, r1, r2

    def m7(self, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Set = _dafny.Set({})
        d_2906_i0_: int
        d_2906_i0_ = 0
        with _dafny.label("14"):
            while self.f16:
                with _dafny.c_label("14"):
                    if (d_2906_i0_) >= (100):
                        raise _dafny.Break("14")
                    d_2906_i0_ = (d_2906_i0_) + (1)
                    d_2907_v0_: _dafny.Set
                    d_2907_v0_ = _dafny.Set({(self).f15})
                    d_2908_v1_: D3
                    d_2908_v1_ = D3_DC6(d_2907_v0_)
                    d_2909_v2_: str
                    d_2909_v2_ = _dafny.CodePoint('e')
                    d_2910_v3_: _dafny.Map
                    d_2910_v3_ = _dafny.Map({d_2908_v1_: default__.fm51(self.f16, (self).f15, d_2909_v2_, self.f16, globalState)})
                    d_2910_v3_ = d_2910_v3_
                    d_2911_v4_: _dafny.Array
                    nw473_ = _dafny.Array(_dafny.MultiSet({}), 21)
                    d_2911_v4_ = nw473_
                    d_2911_v4_ = d_2911_v4_
                    d_2912_v5_: _dafny.Map
                    d_2912_v5_ = _dafny.Map({(self).f15: (self).f15})
                    d_2913_v7_: _dafny.Map
                    def iife216_():
                        coll94_ = _dafny.Map()
                        compr_94_: int
                        for compr_94_ in _dafny.IntegerRange(672, -954):
                            d_2914_v6_: int = compr_94_
                            if ((672) <= (d_2914_v6_)) and ((d_2914_v6_) < (-954)):
                                coll94_[default__.safeModuloInt(d_2914_v6_, (self).f15)] = (self).f15
                        return _dafny.Map(coll94_)
                    d_2913_v7_ = _dafny.Map({d_2912_v5_: iife216_()
                    })
                    d_2915_v8_: _dafny.Seq
                    d_2915_v8_ = _dafny.SeqWithoutIsStrInference([len((_dafny.Map({self.f16: (self).f15})).set(self.f16, len(_dafny.Map({(self).f15: (self).f15}))))])
                    d_2916_v9_: _dafny.Seq
                    d_2916_v9_ = _dafny.SeqWithoutIsStrInference([d_2915_v8_, _dafny.SeqWithoutIsStrInference([(self).f15])])
                    d_2917_v10_: T1
                    nw474_ = C2()
                    nw474_.ctor__((self).f15, d_2913_v7_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")), d_2916_v9_)
                    d_2917_v10_ = nw474_
                    d_2918_v11_: _dafny.Seq
                    d_2918_v11_ = _dafny.SeqWithoutIsStrInference([d_2917_v10_])
                    d_2919_v12_: _dafny.Seq
                    d_2919_v12_ = _dafny.SeqWithoutIsStrInference([d_2918_v11_, d_2918_v11_])
                    r0 = (d_2918_v11_) in (_dafny.MultiSet(d_2919_v12_))
                    d_2920_v13_: _dafny.Array
                    nw475_ = _dafny.Array(_dafny.Set({}), 19)
                    d_2920_v13_ = nw475_
                    d_2920_v13_ = d_2920_v13_
                    pass
            pass
        d_2921_v14_: _dafny.Array
        nw476_ = _dafny.Array(_dafny.Seq({}), 10)
        d_2921_v14_ = nw476_
        guard_loop_10_: int
        for guard_loop_10_ in _dafny.IntegerRange(0, (d_2921_v14_).length(0)):
            d_2922_i1_: int = guard_loop_10_
            if (True) and (((0) <= (d_2922_i1_)) and ((d_2922_i1_) < ((d_2921_v14_).length(0)))):
                (d_2921_v14_)[(d_2922_i1_)] = _dafny.SeqWithoutIsStrInference([D7_DC15(_dafny.MultiSet([len(_dafny.Map({False: (self).f15})), (self).f15, (self).f15]))])
        hi10_ = (0) - ((self).f15)
        for d_2923_i2_ in range(447, hi10_):
            (globalState).f3 = default__.fm0((self).f15, (self).f15, globalState)
            d_2924_v15_: _dafny.Seq
            d_2924_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yoosm"))
            d_2925_v16_: _dafny.Seq
            d_2925_v16_ = _dafny.SeqWithoutIsStrInference([718])
            d_2926_v17_: _dafny.Map
            d_2926_v17_ = _dafny.Map({d_2923_i2_: (self).f15})
            d_2927_v18_: _dafny.Map
            d_2927_v18_ = _dafny.Map({d_2926_v17_: d_2926_v17_})
            d_2928_v19_: C7
            nw477_ = C7()
            nw477_.ctor__(d_2924_v15_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([233]), d_2925_v16_]), (0) - (((self).f15) - (d_2923_i2_)), d_2927_v18_)
            d_2928_v19_ = nw477_
            d_2929_v20_: _dafny.Set
            d_2929_v20_ = _dafny.Set({self.f16})
            (globalState).f3 = (((self).f15) * (-57)) - (len((d_2929_v20_).intersection(d_2929_v20_)))
            d_2930_v21_: str
            d_2930_v21_ = _dafny.CodePoint('t')
            d_2930_v21_ = (d_2924_v15_)[default__.safeIndex(d_2923_i2_, len(d_2924_v15_))]
        d_2931_v22_: _dafny.Seq
        d_2931_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqwke"))
        d_2932_v23_: str
        d_2932_v23_ = _dafny.CodePoint('r')
        (globalState).f3 = len((d_2931_v22_).set(default__.safeIndex((self).f15, len(d_2931_v22_)), d_2932_v23_))
        d_2933_v24_: _dafny.Array
        def lambda133_(d_2934_i3_):
            return self.f16

        init76_ = lambda133_
        nw478_ = _dafny.Array(None, 10)
        for i0_76_ in range(nw478_.length(0)):
            nw478_[i0_76_] = init76_(i0_76_)
        d_2933_v24_ = nw478_
        d_2933_v24_ = d_2933_v24_
        d_2935_v25_: _dafny.Set
        d_2935_v25_ = _dafny.Set({(self).f15, 967})
        d_2936_v26_: D18
        d_2936_v26_ = D18_DC43(self.f16, default__.fm0(len(d_2935_v25_), (0) - ((self).f15), globalState))
        source47_ = d_2936_v26_
        if source47_.is_DC42:
            d_2937___mcc_h0_ = source47_.cf58
            d_2938___mcc_h1_ = source47_.cf59
            d_2939_cf59_ = d_2938___mcc_h1_
            d_2940_cf58_ = d_2937___mcc_h0_
            r0 = (d_2939_cf59_) > (default__.safeModuloInt(882, (self).f15))
            d_2941_v27_: _dafny.Array
            def lambda134_(d_2942_i4_):
                return D1_DC2()

            init77_ = lambda134_
            nw479_ = _dafny.Array(None, 9)
            for i0_77_ in range(nw479_.length(0)):
                nw479_[i0_77_] = init77_(i0_77_)
            d_2941_v27_ = nw479_
            d_2943_v28_: _dafny.Array
            nw480_ = _dafny.Array(None, 9)
            nw480_[int(0)] = d_2941_v27_
            nw480_[int(1)] = d_2941_v27_
            nw480_[int(2)] = d_2941_v27_
            nw480_[int(3)] = d_2941_v27_
            nw480_[int(4)] = d_2941_v27_
            nw480_[int(5)] = d_2941_v27_
            nw480_[int(6)] = d_2941_v27_
            nw480_[int(7)] = d_2941_v27_
            nw480_[int(8)] = d_2941_v27_
            d_2943_v28_ = nw480_
            d_2944_v29_: C5
            nw481_ = C5()
            nw481_.ctor__(default__.safeModuloInt((self).f15, 154), d_2943_v28_)
            d_2944_v29_ = nw481_
            d_2944_v29_ = d_2944_v29_
            if self.f16:
                d_2945_v30_: _dafny.Set
                d_2945_v30_ = _dafny.Set({self.f16, self.f16})
                d_2946_v31_: _dafny.Seq
                d_2946_v31_ = _dafny.SeqWithoutIsStrInference([len(d_2945_v30_)])
                rhs435_ = not(self.f16)
                rhs436_ = not((d_2940_cf58_ if (len(d_2931_v22_)) < (len(d_2946_v31_)) else default__.fm38((self).f15, globalState)))
                d_2940_cf58_ = rhs435_
                d_2940_cf58_ = rhs436_
                d_2947_v32_: _dafny.Map
                d_2947_v32_ = _dafny.Map({self.f16: (d_2944_v29_).f17})
                d_2948_v33_: _dafny.Map
                d_2948_v33_ = _dafny.Map({(d_2944_v29_).f17: -526})
                d_2949_v34_: _dafny.Map
                d_2949_v34_ = _dafny.Map({d_2948_v33_: d_2948_v33_})
                d_2950_v35_: _dafny.Seq
                d_2950_v35_ = _dafny.SeqWithoutIsStrInference([d_2946_v31_, d_2946_v31_])
                d_2951_v36_: T1
                nw482_ = C1()
                nw482_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhxaeriv"))), (self).f15, len(d_2931_v22_), ((d_2949_v34_).set(d_2948_v33_, d_2948_v33_)).set(d_2948_v33_, (_dafny.Map({(self).f15: (d_2944_v29_).f17})).set((d_2944_v29_).f17, len(d_2931_v22_))), d_2931_v22_, d_2950_v35_)
                d_2951_v36_ = nw482_
                d_2952_v37_: _dafny.Map
                d_2952_v37_ = _dafny.Map({(d_2947_v32_).set(self.f16, (d_2944_v29_).f17): d_2951_v36_})
                d_2952_v37_ = (d_2952_v37_).set(default__.fm18(globalState), d_2951_v36_)
                d_2953_v38_: _dafny.Seq
                d_2953_v38_ = _dafny.SeqWithoutIsStrInference([self.f16, self.f16])
                d_2954_v39_: C6
                nw483_ = C6()
                nw483_.ctor__(len(d_2953_v38_), (_dafny.Map({d_2948_v33_: d_2948_v33_})) | (_dafny.Map({d_2948_v33_: d_2948_v33_})), ((d_2951_v36_).f9) + ((d_2951_v36_).f9), (d_2951_v36_).f10)
                d_2954_v39_ = nw483_
                d_2954_v39_ = d_2954_v39_
                r0 = not((D20_DC47(d_2940_cf58_, d_2939_cf59_, d_2940_cf58_)).cf65)
                d_2955_v40_: D2
                d_2955_v40_ = D2_DC4(len((d_2951_v36_).f9))
                d_2956_v41_: _dafny.Seq
                d_2956_v41_ = _dafny.SeqWithoutIsStrInference([d_2955_v40_])
                d_2956_v41_ = (d_2956_v41_) + (_dafny.SeqWithoutIsStrInference([d_2955_v40_, d_2955_v40_, d_2955_v40_, d_2955_v40_]))
            elif True:
                d_2957_v42_: _dafny.Seq
                d_2957_v42_ = _dafny.SeqWithoutIsStrInference([(d_2944_v29_).f17])
                d_2958_v43_: _dafny.Array
                nw484_ = _dafny.Array(None, 3)
                nw484_[int(0)] = (d_2957_v42_)[default__.safeIndex((0) - (d_2939_cf59_), len(d_2957_v42_))]
                nw484_[int(1)] = (d_2944_v29_).f17
                nw484_[int(2)] = (self).f15
                d_2958_v43_ = nw484_
                index419_ = default__.safeIndex(604, (d_2958_v43_).length(0))
                (d_2958_v43_)[index419_] = default__.safeDivisionInt(-80, (self).f15)
                d_2959_v44_: _dafny.MultiSet
                d_2959_v44_ = _dafny.MultiSet([d_2931_v22_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxaahboxc")), d_2931_v22_, default__.fm31(globalState), d_2931_v22_])
                d_2960_v45_: D10
                d_2960_v45_ = D10_DC22(d_2959_v44_)
                d_2961_v46_: _dafny.Seq
                d_2961_v46_ = _dafny.SeqWithoutIsStrInference([default__.fm38(d_2939_cf59_, globalState), self.f16, self.f16, self.f16, d_2940_cf58_])
                pat_let_tv56_ = d_2959_v44_
                pat_let_tv57_ = d_2944_v29_
                index420_ = default__.safeIndex(604, (d_2958_v43_).length(0))
                def iife217_(_pat_let61_0):
                    def iife218_(d_2962_dt__update__tmp_h0_):
                        def iife219_(_pat_let62_0):
                            def iife220_(d_2963_dt__update_hcf25_h0_):
                                return D10_DC22(d_2963_dt__update_hcf25_h0_)
                            return iife220_(_pat_let62_0)
                        return iife219_((pat_let_tv56_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgkcyyhh")), default__.abs((pat_let_tv57_).f17)))
                    return iife218_(_pat_let61_0)
                rhs437_ = (d_2939_cf59_ if not(d_2940_cf58_) else ((self).f15) - (len(d_2961_v46_)))
                rhs438_ = d_2939_cf59_
                rhs439_ = d_2940_cf58_
                rhs440_ = iife217_(d_2960_v45_)
                rhs441_ = (d_2944_v29_).f17
                lhs285_ = d_2958_v43_
                lhs286_ = default__.safeIndex(604, (d_2958_v43_).length(0))
                lhs287_ = globalState
                lhs288_ = globalState
                lhs285_[lhs286_] = rhs437_
                lhs287_.f3 = rhs438_
                r0 = rhs439_
                d_2960_v45_ = rhs440_
                lhs288_.f3 = rhs441_
                d_2964_v47_: _dafny.Map
                d_2964_v47_ = _dafny.Map({(d_2944_v29_).f17: _dafny.SeqWithoutIsStrInference([d_2931_v22_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdhpwa"))])})
                d_2965_v48_: _dafny.Map
                d_2965_v48_ = _dafny.Map({d_2958_v43_: False})
                d_2964_v47_ = (d_2964_v47_).set(len(d_2965_v48_), _dafny.SeqWithoutIsStrInference([d_2931_v22_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))]))
                (globalState).f3 = d_2939_cf59_
                d_2966_v49_: D15
                d_2966_v49_ = D15_DC36((self).f15)
                d_2939_cf59_ = (d_2966_v49_).cf48
                r1 = False
            d_2967_v50_: _dafny.Array
            nw485_ = _dafny.Array(int(0), 11)
            d_2967_v50_ = nw485_
            d_2968_v51_: _dafny.Set
            d_2968_v51_ = _dafny.Set({d_2967_v50_, d_2967_v50_})
            d_2969_v52_: _dafny.Seq
            d_2969_v52_ = _dafny.SeqWithoutIsStrInference([(self).f15, d_2939_cf59_, d_2939_cf59_])
            d_2970_v53_: _dafny.Seq
            d_2970_v53_ = _dafny.SeqWithoutIsStrInference([d_2969_v52_])
            d_2971_v54_: C4
            nw486_ = C4()
            nw486_.ctor__((d_2968_v51_).issubset(_dafny.Set({d_2967_v50_})), self.f16, d_2931_v22_, d_2970_v53_)
            d_2971_v54_ = nw486_
        elif source47_.is_DC43:
            d_2972___mcc_h2_ = source47_.cf60
            d_2973___mcc_h3_ = source47_.cf61
            d_2974_cf61_ = d_2973___mcc_h3_
            d_2975_cf60_ = d_2972___mcc_h2_
            d_2976_v55_: _dafny.Seq
            d_2976_v55_ = _dafny.SeqWithoutIsStrInference([d_2975_cf60_])
            d_2976_v55_ = d_2976_v55_
            d_2977_v56_: D3
            d_2977_v56_ = D3_DC7(default__.fm38(d_2974_cf61_, globalState))
            source48_ = d_2977_v56_
            if source48_.is_DC7:
                d_2978___mcc_h6_ = source48_.cf9
                d_2979_cf9_ = d_2978___mcc_h6_
                d_2980_v57_: _dafny.MultiSet
                d_2980_v57_ = _dafny.MultiSet([_dafny.CodePoint('s')])
                d_2981_v59_: _dafny.Map
                def iife221_():
                    coll95_ = _dafny.Set()
                    compr_95_: str
                    for compr_95_ in (d_2980_v57_).Elements:
                        d_2982_v58_: str = compr_95_
                        if (d_2982_v58_) in (d_2980_v57_):
                            coll95_ = coll95_.union(_dafny.Set([d_2982_v58_]))
                    return _dafny.Set(coll95_)
                d_2981_v59_ = _dafny.Map({d_2932_v23_: default__.safeDivisionInt(len(iife221_()
                ), 258)})
                d_2981_v59_ = (d_2981_v59_).set(d_2932_v23_, d_2974_cf61_)
                d_2983_v60_: D3
                d_2983_v60_ = D3_DC6(_dafny.Set({730, d_2974_cf61_}))
                d_2984_v61_: _dafny.Map
                d_2984_v61_ = _dafny.Map({d_2933_v24_: d_2983_v60_})
                d_2984_v61_ = (d_2984_v61_).set(d_2933_v24_, d_2983_v60_)
                d_2985_v62_: _dafny.Array
                def lambda135_(d_2986_cf61_):
                    def lambda136_(d_2987_i5_):
                        return _dafny.SeqWithoutIsStrInference([d_2986_cf61_, (self).f15])

                    return lambda136_

                init78_ = lambda135_(d_2974_cf61_)
                nw487_ = _dafny.Array(None, 20)
                for i0_78_ in range(nw487_.length(0)):
                    nw487_[i0_78_] = init78_(i0_78_)
                d_2985_v62_ = nw487_
                d_2988_v63_: _dafny.Map
                d_2988_v63_ = _dafny.Map({d_2974_cf61_: d_2931_v22_})
                d_2989_v64_: _dafny.Seq
                d_2989_v64_ = _dafny.SeqWithoutIsStrInference([len(((d_2988_v63_)[(self).f15] if ((self).f15) in (d_2988_v63_) else d_2931_v22_)), (self).f15])
                index421_ = default__.safeIndex(37, (d_2985_v62_).length(0))
                (d_2985_v62_)[index421_] = d_2989_v64_
                index422_ = default__.safeIndex(37, (d_2985_v62_).length(0))
                (d_2985_v62_)[index422_] = d_2989_v64_
                d_2990_v65_: _dafny.Map
                d_2990_v65_ = _dafny.Map({(self).f15: len(d_2931_v22_)})
                d_2991_v66_: D2
                d_2991_v66_ = D2_DC4(d_2974_cf61_)
                d_2992_v67_: _dafny.MultiSet
                d_2992_v67_ = _dafny.MultiSet([(d_2991_v66_).cf3, d_2974_cf61_, (self).f15, (self).f15, d_2974_cf61_])
                d_2990_v65_ = (d_2990_v65_).set(d_2974_cf61_, (d_2992_v67_).cardinality)
            elif True:
                d_2993___mcc_h7_ = source48_.cf8
                d_2994_cf8_ = d_2993___mcc_h7_
                d_2995_v68_: _dafny.MultiSet
                d_2995_v68_ = _dafny.MultiSet([len(d_2976_v55_)])
                d_2996_v69_: _dafny.Set
                d_2996_v69_ = _dafny.Set({d_2995_v68_, d_2995_v68_})
                d_2997_v70_: D10
                d_2997_v70_ = D10_DC23(d_2932_v23_, d_2996_v69_)
                d_2998_v71_: _dafny.Map
                d_2998_v71_ = _dafny.Map({d_2997_v70_: False})
                d_2999_v72_: _dafny.MultiSet
                d_2999_v72_ = _dafny.MultiSet([d_2998_v71_])
                d_3000_v73_: _dafny.MultiSet
                d_3000_v73_ = _dafny.MultiSet([self.f16, (((d_2999_v72_)[d_2998_v71_] if (d_2998_v71_) in (d_2999_v72_) else -462)) == (d_2974_cf61_), d_2975_cf60_])
                d_3000_v73_ = _dafny.MultiSet((D8_DC17(((D8_DC17(d_2976_v55_)).cf19).set(default__.safeIndex((self).f15, len((D8_DC17(d_2976_v55_)).cf19)), d_2975_cf60_))).cf19)
                (globalState).f3 = default__.safeDivisionInt(default__.safeModuloInt(d_2974_cf61_, (self).f15), ((d_3000_v73_) - (_dafny.MultiSet(d_2976_v55_))).cardinality)
                d_3001_v74_: _dafny.Map
                d_3001_v74_ = _dafny.Map({-217: _dafny.SeqWithoutIsStrInference([d_2932_v23_, d_2932_v23_, _dafny.CodePoint('h')])})
                d_2931_v22_ = default__.fm26(len(d_2931_v22_), (d_3001_v74_).set(588, d_2931_v22_), globalState)
                d_2976_v55_ = (_dafny.SeqWithoutIsStrInference([self.f16])) + (d_2976_v55_)
            d_2974_cf61_ = 579
            r0 = self.f16
        elif source47_.is_DC41:
            d_3002___mcc_h4_ = source47_.cf57
            d_3003_cf57_ = d_3002___mcc_h4_
            d_3004_v75_: _dafny.Seq
            d_3004_v75_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({False: self.f16})])
            d_3005_v76_: _dafny.Map
            d_3005_v76_ = _dafny.Map({self.f16: not(self.f16)})
            (globalState).f3 = default__.safeModuloInt((self).f15, len(((d_3004_v75_)[default__.safeIndex(len(d_2931_v22_), len(d_3004_v75_))]) | (d_3005_v76_)))
            d_3006_v77_: _dafny.Map
            d_3006_v77_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_3007_i6_ in range(default__.abs(430))]): (self).f15})
            (globalState).f3 = ((self).f15) + ((0) - (((d_3006_v77_)[d_2931_v22_] if (d_2931_v22_) in (d_3006_v77_) else (0) - ((self).f15))))
            d_3008_v78_: _dafny.Map
            d_3008_v78_ = _dafny.Map({d_2932_v23_: (self).f15})
            d_3008_v78_ = (d_3008_v78_).set(d_2932_v23_, (self).f15)
            d_3009_v79_: _dafny.Seq
            d_3009_v79_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_3010_v80_: _dafny.Set
            d_3010_v80_ = _dafny.Set({d_3005_v76_, d_3005_v76_})
            d_3011_v81_: _dafny.Array
            nw488_ = _dafny.Array(_dafny.CodePoint('D'), 27)
            d_3011_v81_ = nw488_
            d_3012_v82_: D20
            d_3012_v82_ = D20_DC48(self.f16, d_3010_v80_, d_3011_v81_)
            d_3013_v84_: _dafny.Seq
            d_3013_v84_ = _dafny.SeqWithoutIsStrInference([(self).f15, len(d_3009_v79_), (self).f15, 544, len(d_2931_v22_)])
            d_3014_v85_: _dafny.Map
            d_3014_v85_ = _dafny.Map({self.f16: (self).f15})
            d_3015_v86_: _dafny.Map
            d_3015_v86_ = _dafny.Map({len(d_3014_v85_): (self).f15})
            d_3016_v87_: _dafny.Map
            def iife222_():
                coll96_ = _dafny.Map()
                compr_96_: int
                for compr_96_ in (d_3013_v84_).Elements:
                    d_3017_v83_: int = compr_96_
                    if (d_3017_v83_) in (d_3013_v84_):
                        coll96_[(d_3017_v83_) + (len(d_3005_v76_))] = (self).f15
                return _dafny.Map(coll96_)
            d_3016_v87_ = _dafny.Map({iife222_()
            : d_3015_v86_})
            d_3018_v88_: D12
            d_3018_v88_ = D12_DC29(self.f16, (self).f15, d_3013_v84_, self.f16, (self).f15)
            d_3019_v89_: _dafny.Seq
            d_3019_v89_ = _dafny.SeqWithoutIsStrInference([d_3013_v84_, ((d_3018_v88_).cf34).set(default__.safeIndex(len(d_3014_v85_), len((d_3018_v88_).cf34)), (self).f15)])
            d_3020_v90_: C1
            nw489_ = C1()
            nw489_.ctor__((self).f15, (self).f15, (self).f15, d_3016_v87_, _dafny.SeqWithoutIsStrInference([d_2932_v23_ for d_3021_i7_ in range(default__.abs(243))]), d_3019_v89_)
            d_3020_v90_ = nw489_
            d_3022_v91_: _dafny.Set
            d_3022_v91_ = _dafny.Set({d_3020_v90_})
            d_3023_v92_: _dafny.Seq
            d_3023_v92_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_3020_v90_}), d_3022_v91_, d_3022_v91_, d_3022_v91_, _dafny.Set({d_3020_v90_, d_3020_v90_, d_3020_v90_})])
            rhs442_ = _dafny.SeqWithoutIsStrInference([default__.fm38((_dafny.MultiSet([(self).f15, len(d_2935_v25_)])).cardinality, globalState), default__.fm38((self).f15, globalState), (d_3012_v82_).cf68])
            rhs443_ = d_2933_v24_
            rhs444_ = default__.safeModuloInt((self).f15, len(_dafny.SeqWithoutIsStrInference([self.f16])))
            rhs445_ = ((self).f15) + (len(d_3023_v92_))
            rhs446_ = default__.safeDivisionInt(d_3020_v90_.f23, (self).f15)
            lhs289_ = globalState
            lhs290_ = globalState
            lhs291_ = globalState
            d_3009_v79_ = rhs442_
            d_2933_v24_ = rhs443_
            lhs289_.f3 = rhs444_
            lhs290_.f3 = rhs445_
            lhs291_.f3 = rhs446_
        elif True:
            d_3024___mcc_h5_ = source47_.cf62
            d_3025_cf62_ = d_3024___mcc_h5_
            (globalState).f3 = (self).f15
            (globalState).f3 = default__.fm0((default__.fm0((self).f15, (self).f15, globalState)) * ((self).f15), 999, globalState)
            d_3026_v93_: _dafny.Map
            d_3026_v93_ = _dafny.Map({(self).f15: d_2931_v22_})
            d_3027_v94_: _dafny.Seq
            d_3027_v94_ = _dafny.SeqWithoutIsStrInference([len(default__.fm26((self).f15, d_3026_v93_, globalState))])
            d_3028_v95_: _dafny.MultiSet
            d_3028_v95_ = _dafny.MultiSet([(self).f15])
            d_3029_v96_: _dafny.Seq
            d_3029_v96_ = _dafny.SeqWithoutIsStrInference([d_3028_v95_])
            d_3030_v97_: _dafny.Seq
            d_3030_v97_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_3031_v98_: _dafny.Seq
            d_3031_v98_ = _dafny.SeqWithoutIsStrInference([(d_3030_v97_)[default__.safeIndex(len(d_2931_v22_), len(d_3030_v97_))], self.f16])
            d_3032_v99_: _dafny.Map
            d_3032_v99_ = _dafny.Map({((d_3029_v96_)[default__.safeIndex(len(d_3030_v97_), len(d_3029_v96_))]).cardinality: 215})
            d_3033_v100_: _dafny.Map
            d_3033_v100_ = _dafny.Map({d_3032_v99_: _dafny.Map({(self).f15: (self).f15})})
            d_3034_v101_: C7
            nw490_ = C7()
            nw490_.ctor__((d_2931_v22_) + (d_2931_v22_), _dafny.SeqWithoutIsStrInference([d_3027_v94_, d_3027_v94_]), (self).f15, (d_3033_v100_) | (d_3033_v100_))
            d_3034_v101_ = nw490_
            d_3035_v102_: int
            d_3036_v103_: _dafny.Map
            d_3037_v104_: bool
            d_3038_v105_: int
            out79_: int
            out80_: _dafny.Map
            out81_: bool
            out82_: int
            out79_, out80_, out81_, out82_ = (d_3034_v101_).m2((self).f15, globalState)
            d_3035_v102_ = out79_
            d_3036_v103_ = out80_
            d_3037_v104_ = out81_
            d_3038_v105_ = out82_
        d_3039_v106_: _dafny.Map
        d_3039_v106_ = _dafny.Map({(self).f15: (self).f15})
        d_3040_v107_: _dafny.Seq
        d_3040_v107_ = _dafny.SeqWithoutIsStrInference([self.f16])
        d_3041_v108_: _dafny.Map
        d_3041_v108_ = _dafny.Map({self.f16: len(d_3040_v107_)})
        d_3042_v109_: _dafny.MultiSet
        d_3042_v109_ = _dafny.MultiSet([(self).f15])
        d_3043_v110_: D7
        d_3043_v110_ = D7_DC15(d_3042_v109_)
        r0 = (((self).f15) != (((d_3039_v106_)[(self).f15] if ((self).f15) in (d_3039_v106_) else len((d_3041_v108_).set(self.f16, (self).f15))))) or (((d_3043_v110_).cf18).ispropersubset(d_3042_v109_))
        r1 = self.f16
        d_3044_v111_: _dafny.Seq
        d_3044_v111_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))])
        r2 = default__.fm67((_dafny.SeqWithoutIsStrInference([d_2932_v23_ for d_3045_i8_ in range(default__.abs(-871))])) + ((d_3044_v111_)[default__.safeIndex((self).f15, len(d_3044_v111_))]), globalState)
        return r0, r1, r2

    @property
    def f15(self):
        return self._f15

class C9:
    def  __init__(self):
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    def ctor__(self, f14):
        (self)._f14 = f14

    def fm7(self, p0, p1, globalState):
        return default__.safeModuloInt((0) - (default__.safeModuloInt(len(_dafny.Map({not((D1_DC1(False)).cf1): False})), 907)), len((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False]))))

    def m5(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: int = int(0)
        d_3046_v0_: int
        d_3046_v0_ = 128
        hi11_ = d_3046_v0_
        for d_3047_i0_ in range(default__.safeDivisionInt(-583, d_3046_v0_), hi11_):
            if p2:
                d_3048_v1_: D1
                d_3048_v1_ = D1_DC2()
                d_3049_v2_: str
                d_3049_v2_ = _dafny.CodePoint('s')
                d_3050_v3_: _dafny.Map
                d_3050_v3_ = _dafny.Map({d_3049_v2_: 43})
                d_3051_v4_: _dafny.Map
                d_3051_v4_ = _dafny.Map({len(d_3050_v3_): p2})
                d_3052_v5_: D1
                d_3052_v5_ = D1_DC1(p1)
                rhs447_ = d_3048_v1_
                rhs448_ = not((((d_3051_v4_)[444] if (444) in (d_3051_v4_) else not(p2))) == ((d_3052_v5_).cf1))
                rhs449_ = p1
                d_3048_v1_ = rhs447_
                r0 = rhs448_
                r0 = rhs449_
                d_3046_v0_ = d_3046_v0_
                d_3053_v6_: _dafny.Map
                d_3053_v6_ = _dafny.Map({p1: d_3047_i0_})
                d_3054_v7_: _dafny.Map
                d_3054_v7_ = _dafny.Map({d_3046_v0_: d_3046_v0_})
                d_3055_v8_: D2
                d_3055_v8_ = D2_DC5(p1, (self).f14, d_3046_v0_, len(d_3054_v7_))
                d_3046_v0_ = default__.safeModuloInt(((d_3053_v6_)[True] if (True) in (d_3053_v6_) else d_3047_i0_), (d_3055_v8_).cf6)
                r1 = d_3046_v0_
                (globalState).f3 = ((0) - (len(_dafny.Set({d_3046_v0_})))) * (d_3047_i0_)
            elif True:
                d_3056_v9_: _dafny.MultiSet
                d_3056_v9_ = _dafny.MultiSet([(self).f14])
                d_3057_v10_: _dafny.Set
                d_3057_v10_ = _dafny.Set({d_3047_i0_})
                d_3058_v11_: D3
                d_3058_v11_ = D3_DC6(d_3057_v10_)
                d_3059_v12_: _dafny.MultiSet
                d_3059_v12_ = _dafny.MultiSet([d_3057_v10_, d_3057_v10_, (d_3058_v11_).cf8, d_3057_v10_])
                r1 = default__.safeDivisionInt(d_3046_v0_, (d_3046_v0_ if default__.fm8((d_3056_v9_).cardinality, p1, p2, globalState) else (d_3059_v12_).cardinality))
                d_3060_v13_: _dafny.Map
                d_3060_v13_ = _dafny.Map({default__.safeDivisionInt(d_3046_v0_, d_3047_i0_): (self).f14})
                d_3060_v13_ = (d_3060_v13_).set(d_3047_i0_, (self).f14)
                r0 = p1
                d_3046_v0_ = d_3047_i0_
                index423_ = default__.safeIndex(984, ((self).f14).length(0))
                ((self).f14)[index423_] = p1
                index424_ = default__.safeIndex(984, ((self).f14).length(0))
                ((self).f14)[index424_] = p2
            index425_ = default__.safeIndex(158, ((self).f14).length(0))
            ((self).f14)[index425_] = False
            d_3061_v14_: str
            d_3061_v14_ = _dafny.CodePoint('l')
            index426_ = default__.safeIndex(158, ((self).f14).length(0))
            rhs450_ = (0) - (-891)
            rhs451_ = (d_3061_v14_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yc")))
            rhs452_ = ((d_3046_v0_) + (d_3047_i0_) if p2 else d_3046_v0_)
            lhs292_ = (self).f14
            lhs293_ = default__.safeIndex(158, ((self).f14).length(0))
            r1 = rhs450_
            lhs292_[lhs293_] = rhs451_
            r1 = rhs452_
            index427_ = default__.safeIndex(158, ((self).f14).length(0))
            ((self).f14)[index427_] = False
            index428_ = default__.safeIndex(158, ((self).f14).length(0))
            ((self).f14)[index428_] = ((self).f14)[default__.safeIndex(158, ((self).f14).length(0))]
        d_3062_v15_: _dafny.Set
        d_3062_v15_ = _dafny.Set({(self).f14, (self).f14})
        d_3062_v15_ = (_dafny.Set({(self).f14})) - (d_3062_v15_)
        d_3063_v16_: D1
        d_3063_v16_ = D1_DC1(p1)
        d_3064_v17_: _dafny.Map
        d_3064_v17_ = _dafny.Map({p2: 45})
        d_3065_v18_: _dafny.Set
        d_3065_v18_ = _dafny.Set({((d_3064_v17_)[p1] if (p1) in (d_3064_v17_) else d_3046_v0_)})
        d_3066_v19_: D3
        d_3066_v19_ = D3_DC6(d_3065_v18_)
        d_3067_v20_: str
        d_3067_v20_ = _dafny.CodePoint('c')
        d_3068_v21_: _dafny.Seq
        d_3068_v21_ = _dafny.SeqWithoutIsStrInference([d_3066_v19_, d_3066_v19_, default__.fm9((d_3063_v16_).cf1, not(p2), p2, d_3067_v20_, globalState), d_3066_v19_])
        pat_let_tv58_ = d_3068_v21_
        pat_let_tv59_ = d_3046_v0_
        pat_let_tv60_ = d_3068_v21_
        pat_let_tv61_ = d_3066_v19_
        pat_let_tv62_ = d_3066_v19_
        def iife223_(_pat_let63_0):
            def iife224_(d_3069_dt__update__tmp_h0_):
                def iife225_(_pat_let64_0):
                    def iife226_(d_3070_dt__update_hcf1_h0_):
                        return D1_DC1(d_3070_dt__update_hcf1_h0_)
                    return iife226_(_pat_let64_0)
                return iife225_(((pat_let_tv58_).set(default__.safeIndex((0) - (pat_let_tv59_), len(pat_let_tv60_)), pat_let_tv61_)) != (_dafny.SeqWithoutIsStrInference([pat_let_tv62_])))
            return iife224_(_pat_let63_0)
        source49_ = iife223_(d_3063_v16_)
        if source49_.is_DC2:
            (globalState).f3 = d_3046_v0_
            d_3071_v22_: _dafny.MultiSet
            d_3071_v22_ = _dafny.MultiSet([p1, p1])
            d_3072_v23_: _dafny.Map
            d_3072_v23_ = _dafny.Map({d_3046_v0_: (d_3071_v22_).cardinality})
            d_3073_v24_: _dafny.Seq
            d_3073_v24_ = _dafny.SeqWithoutIsStrInference([d_3046_v0_, d_3046_v0_, d_3046_v0_])
            d_3074_v25_: _dafny.Seq
            d_3074_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_3046_v0_, 787, d_3046_v0_]), d_3073_v24_])
            d_3075_v26_: C1
            nw491_ = C1()
            nw491_.ctor__(default__.safeDivisionInt(d_3046_v0_, ((d_3072_v23_)[d_3046_v0_] if (d_3046_v0_) in (d_3072_v23_) else d_3046_v0_)), len((p0 if not(p1) else (p0).set(default__.safeIndex((d_3073_v24_)[default__.safeIndex(d_3046_v0_, len(d_3073_v24_))], len(p0)), p1))), (0) - ((0) - (d_3046_v0_)), default__.fm47(p2, p2, p1, p1, globalState), _dafny.SeqWithoutIsStrInference([d_3067_v20_ for d_3076_i1_ in range(default__.abs(-636))]), d_3074_v25_)
            d_3075_v26_ = nw491_
            if (d_3046_v0_) < (d_3075_v26_.f23):
                d_3077_v27_: _dafny.Map
                d_3077_v27_ = _dafny.Map({d_3067_v20_: p2})
                d_3078_v28_: _dafny.MultiSet
                d_3078_v28_ = _dafny.MultiSet([d_3046_v0_, len((d_3077_v27_).set(d_3067_v20_, p2)), d_3075_v26_.f23, d_3075_v26_.f23])
                d_3079_v29_: _dafny.Map
                d_3079_v29_ = _dafny.Map({default__.fm68(d_3078_v28_, p1, globalState): p0})
                d_3080_v30_: D21
                d_3080_v30_ = D21_DC49(_dafny.Set({d_3067_v20_}))
                d_3079_v29_ = (d_3079_v29_).set(d_3080_v30_, p0)
                d_3081_v31_: _dafny.Seq
                d_3081_v31_ = _dafny.SeqWithoutIsStrInference([d_3064_v17_])
                index429_ = default__.safeIndex(790, ((self).f14).length(0))
                ((self).f14)[index429_] = (default__.fm38(len(((d_3081_v31_)[default__.safeIndex((d_3075_v26_).f24, len(d_3081_v31_))]).set(p2, (d_3075_v26_).f24)), globalState) if True else True)
                d_3082_v32_: _dafny.Array
                nw492_ = _dafny.Array(int(0), 12)
                d_3082_v32_ = nw492_
                d_3083_v33_: _dafny.Map
                d_3083_v33_ = _dafny.Map({(d_3075_v26_).f24: d_3082_v32_})
                d_3084_v34_: _dafny.Set
                d_3084_v34_ = _dafny.Set({((d_3083_v33_)[d_3075_v26_.f23] if (d_3075_v26_.f23) in (d_3083_v33_) else d_3082_v32_)})
                index430_ = default__.safeIndex(790, ((self).f14).length(0))
                ((self).f14)[index430_] = ((d_3084_v34_).intersection(d_3084_v34_)).ispropersubset((d_3084_v34_ if p2 else (D22_DC52(d_3084_v34_)).cf77))
                d_3085_v35_: _dafny.Map
                d_3085_v35_ = _dafny.Map({((self).f14)[default__.safeIndex(790, ((self).f14).length(0))]: p1})
                d_3085_v35_ = (d_3085_v35_).set(default__.fm38((d_3075_v26_).f24, globalState), p1)
                d_3086_v36_: _dafny.Array
                nw493_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_3086_v36_ = nw493_
                d_3087_v37_: C5
                nw494_ = C5()
                nw494_.ctor__((d_3075_v26_).f24, d_3086_v36_)
                d_3087_v37_ = nw494_
                d_3088_v38_: _dafny.Map
                d_3088_v38_ = _dafny.Map({(d_3075_v26_).f24: d_3087_v37_})
                d_3089_v39_: _dafny.Seq
                d_3089_v39_ = _dafny.SeqWithoutIsStrInference([d_3088_v38_])
                d_3090_v40_: _dafny.Map
                d_3090_v40_ = _dafny.Map({p2: (d_3089_v39_)[default__.safeIndex(d_3075_v26_.f23, len(d_3089_v39_))]})
                index431_ = default__.safeIndex(790, ((self).f14).length(0))
                ((self).f14)[index431_] = ((d_3075_v26_).f24) == ((0) - ((len((d_3090_v40_).set(p1, d_3088_v38_))) * (len(d_3073_v24_))))
                d_3091_v41_: _dafny.Array
                nw495_ = _dafny.Array(_dafny.Seq({}), 4)
                d_3091_v41_ = nw495_
                d_3092_v42_: D11
                d_3092_v42_ = D11_DC25(d_3083_v33_)
                d_3093_v43_: D11
                d_3093_v43_ = D11_DC27(d_3092_v42_)
                d_3094_v44_: _dafny.Map
                d_3094_v44_ = _dafny.Map({d_3093_v43_: d_3067_v20_})
                d_3095_v45_: _dafny.Seq
                d_3095_v45_ = _dafny.SeqWithoutIsStrInference([d_3094_v44_, d_3094_v44_])
                index432_ = default__.safeIndex(613, (d_3091_v41_).length(0))
                (d_3091_v41_)[index432_] = (d_3095_v45_) + (d_3095_v45_)
                index433_ = default__.safeIndex(613, (d_3091_v41_).length(0))
                (d_3091_v41_)[index433_] = d_3095_v45_
            elif True:
                index434_ = default__.safeIndex(599, ((self).f14).length(0))
                ((self).f14)[index434_] = p1
                d_3096_v46_: D21
                d_3096_v46_ = D21_DC50(d_3075_v26_.f23, p2, p2, p1)
                d_3097_v47_: D18
                d_3097_v47_ = D18_DC42((d_3096_v46_).cf74, (d_3075_v26_).f24)
                d_3098_v48_: _dafny.Set
                d_3098_v48_ = _dafny.Set({False, (d_3097_v47_).cf58})
                index435_ = default__.safeIndex(599, ((self).f14).length(0))
                ((self).f14)[index435_] = ((d_3098_v48_).intersection(_dafny.Set({default__.fm38((d_3075_v26_).f24, globalState), p1}))).issubset(default__.fm51(True, d_3075_v26_.f23, d_3067_v20_, not(p1), globalState))
                d_3046_v0_ = ((d_3072_v23_)[len(d_3073_v24_)] if (len(d_3073_v24_)) in (d_3072_v23_) else (0) - (default__.safeModuloInt(d_3075_v26_.f23, d_3075_v26_.f23)))
                d_3099_v49_: _dafny.Map
                d_3099_v49_ = _dafny.Map({776: ((self).f14)[default__.safeIndex(599, ((self).f14).length(0))]})
                d_3100_v50_: D4
                d_3100_v50_ = D4_DC10(d_3099_v49_)
                d_3101_v51_: _dafny.Set
                d_3101_v51_ = _dafny.Set({(d_3100_v50_).cf12, _dafny.Map({d_3075_v26_.f23: p1})})
                d_3102_v52_: D23
                d_3102_v52_ = D23_DC55(d_3101_v51_)
                d_3101_v51_ = (d_3102_v52_).cf82
                r0 = ((self).f14)[default__.safeIndex(599, ((self).f14).length(0))]
                d_3103_v53_: C8
                nw496_ = C8()
                nw496_.ctor__((537 if p1 else (d_3075_v26_).f24), (len(p0)) != (d_3046_v0_))
                d_3103_v53_ = nw496_
            d_3104_v54_: D1
            d_3104_v54_ = D1_DC2()
            source50_ = d_3104_v54_
            if source50_.is_DC2:
                d_3105_v55_: _dafny.Array
                nw497_ = _dafny.Array(_dafny.Array(None, 0), 20)
                d_3105_v55_ = nw497_
                d_3106_v56_: _dafny.Seq
                d_3106_v56_ = _dafny.SeqWithoutIsStrInference([d_3105_v55_, d_3105_v55_, d_3105_v55_])
                d_3107_v57_: _dafny.Array
                nw498_ = _dafny.Array(None, 12)
                nw498_[int(0)] = d_3105_v55_
                nw498_[int(1)] = d_3105_v55_
                nw498_[int(2)] = d_3105_v55_
                nw498_[int(3)] = d_3105_v55_
                nw498_[int(4)] = d_3105_v55_
                nw498_[int(5)] = d_3105_v55_
                nw498_[int(6)] = (d_3106_v56_)[default__.safeIndex(d_3075_v26_.f23, len(d_3106_v56_))]
                nw498_[int(7)] = (d_3105_v55_ if True else d_3105_v55_)
                nw498_[int(8)] = d_3105_v55_
                nw498_[int(9)] = d_3105_v55_
                nw498_[int(10)] = d_3105_v55_
                nw498_[int(11)] = d_3105_v55_
                d_3107_v57_ = nw498_
                d_3108_v58_: D24
                d_3108_v58_ = D24_DC58(d_3105_v55_)
                index436_ = default__.safeIndex(931, (d_3107_v57_).length(0))
                (d_3107_v57_)[index436_] = (d_3108_v58_).cf89
                index437_ = default__.safeIndex(931, (d_3107_v57_).length(0))
                (d_3107_v57_)[index437_] = (d_3105_v55_ if p2 else (d_3108_v58_).cf89)
                d_3109_v59_: _dafny.Array
                def lambda137_(d_3110_p2_):
                    def lambda138_(d_3111_i2_):
                        return d_3110_p2_

                    return lambda138_

                init79_ = lambda137_(p2)
                nw499_ = _dafny.Array(None, 8)
                for i0_79_ in range(nw499_.length(0)):
                    nw499_[i0_79_] = init79_(i0_79_)
                d_3109_v59_ = nw499_
                d_3109_v59_ = d_3109_v59_
                d_3112_v60_: _dafny.Seq
                d_3112_v60_ = _dafny.SeqWithoutIsStrInference([(self).f14])
                d_3113_v61_: _dafny.Seq
                d_3113_v61_ = _dafny.SeqWithoutIsStrInference([(d_3112_v60_)[default__.safeIndex((d_3075_v26_).f24, len(d_3112_v60_))], d_3109_v59_, (self).f14, d_3109_v59_])
                d_3114_v62_: _dafny.Seq
                d_3114_v62_ = _dafny.SeqWithoutIsStrInference([(self).f14])
                d_3113_v61_ = (d_3114_v62_)
                d_3115_v63_: _dafny.Set
                d_3115_v63_ = _dafny.Set({p2, p1, p1, p2, False})
                rhs453_ = (d_3075_v26_).f24
                rhs454_ = not((len((d_3115_v63_) | (default__.fm51(p2, -611, d_3067_v20_, p1, globalState)))) >= (d_3046_v0_))
                d_3046_v0_ = rhs453_
                r0 = rhs454_
            elif source50_.is_DC1:
                d_3116___mcc_h2_ = source50_.cf1
                d_3117_cf1_ = d_3116___mcc_h2_
                (globalState).f3 = (0) - (default__.safeDivisionInt((_dafny.MultiSet(default__.fm44(True, (d_3075_v26_).f24, p2, globalState))).cardinality, -941))
                (d_3075_v26_).f23 = (d_3075_v26_).f24
                index438_ = default__.safeIndex(549, ((self).f14).length(0))
                ((self).f14)[index438_] = p1
                index439_ = default__.safeIndex(549, ((self).f14).length(0))
                ((self).f14)[index439_] = not(not(p2))
                d_3118_v64_: _dafny.Array
                def lambda139_(d_3119_v20_):
                    def lambda140_(d_3120_i3_):
                        return d_3119_v20_

                    return lambda140_

                init80_ = lambda139_(d_3067_v20_)
                nw500_ = _dafny.Array(None, 15)
                for i0_80_ in range(nw500_.length(0)):
                    nw500_[i0_80_] = init80_(i0_80_)
                d_3118_v64_ = nw500_
                index440_ = default__.safeIndex(494, (d_3118_v64_).length(0))
                (d_3118_v64_)[index440_] = d_3067_v20_
                index441_ = default__.safeIndex(494, (d_3118_v64_).length(0))
                (d_3118_v64_)[index441_] = d_3067_v20_
            elif True:
                d_3121___mcc_h3_ = source50_.cf2
                d_3122_cf2_ = d_3121___mcc_h3_
                (d_3075_v26_).f23 = (d_3075_v26_).fm5(((d_3075_v26_).f24 if not(default__.fm38((d_3075_v26_).f24, globalState)) else d_3046_v0_), len((_dafny.SeqWithoutIsStrInference([p2, p2, p1, p1, p1])) + (p0)), globalState)
                d_3123_v65_: _dafny.Seq
                d_3123_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
                index442_ = default__.safeIndex(316, ((self).f14).length(0))
                ((self).f14)[index442_] = ((d_3062_v15_)).issubset(d_3062_v15_)
                d_3124_v66_: _dafny.Array
                nw501_ = _dafny.Array(int(0), 13)
                d_3124_v66_ = nw501_
                index443_ = default__.safeIndex(316, ((self).f14).length(0))
                rhs455_ = d_3123_v65_
                rhs456_ = default__.fm22((p0)[default__.safeIndex((d_3075_v26_).f24, len(p0))], d_3075_v26_.f23, globalState)
                rhs457_ = d_3124_v66_
                lhs294_ = (self).f14
                lhs295_ = default__.safeIndex(316, ((self).f14).length(0))
                d_3123_v65_ = rhs455_
                lhs294_[lhs295_] = rhs456_
                d_3124_v66_ = rhs457_
                (d_3075_v26_).f23 = ((0) - ((d_3075_v26_).f24)) * ((d_3075_v26_).f24)
                index444_ = default__.safeIndex(316, ((self).f14).length(0))
                ((self).f14)[index444_] = p2
        elif source49_.is_DC1:
            d_3125___mcc_h0_ = source49_.cf1
            d_3126_cf1_ = d_3125___mcc_h0_
            d_3127_v67_: _dafny.Seq
            d_3127_v67_ = _dafny.SeqWithoutIsStrInference([d_3046_v0_])
            d_3128_v68_: _dafny.Map
            d_3128_v68_ = _dafny.Map({default__.fm8(d_3046_v0_, d_3126_cf1_, p1, globalState): ((d_3127_v67_)[default__.safeIndex(d_3046_v0_, len(d_3127_v67_))]) == (d_3046_v0_)})
            d_3128_v68_ = (d_3128_v68_).set(p2, p2)
            d_3129_v69_: _dafny.Map
            d_3129_v69_ = _dafny.Map({_dafny.CodePoint('x'): d_3046_v0_})
            d_3130_v70_: _dafny.Seq
            d_3130_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lejmoqt"))
            (globalState).f3 = (len(d_3129_v69_)) - (default__.safeDivisionInt(d_3046_v0_, (self).fm7(len(default__.fm45(d_3046_v0_, _dafny.Map({d_3130_v70_: d_3126_cf1_}), (0) - (d_3046_v0_), True, globalState)), d_3046_v0_, globalState)))
            d_3131_v71_: _dafny.Seq
            d_3131_v71_ = _dafny.SeqWithoutIsStrInference([True])
            d_3131_v71_ = d_3131_v71_
            d_3132_v72_: _dafny.Array
            def lambda141_(d_3133_v0_):
                def lambda142_(d_3134_i4_):
                    return default__.safeModuloInt(d_3134_i4_, d_3133_v0_)

                return lambda142_

            init81_ = lambda141_(d_3046_v0_)
            nw502_ = _dafny.Array(None, 22)
            for i0_81_ in range(nw502_.length(0)):
                nw502_[i0_81_] = init81_(i0_81_)
            d_3132_v72_ = nw502_
            index445_ = default__.safeIndex(433, (d_3132_v72_).length(0))
            (d_3132_v72_)[index445_] = ((d_3064_v17_)[p1] if (p1) in (d_3064_v17_) else 93)
            index446_ = default__.safeIndex(433, (d_3132_v72_).length(0))
            (d_3132_v72_)[index446_] = d_3046_v0_
        elif True:
            d_3135___mcc_h1_ = source49_.cf2
            d_3136_cf2_ = d_3135___mcc_h1_
            d_3137_v73_: D20
            d_3137_v73_ = D20_DC47(False, d_3046_v0_, p2)
            d_3138_v74_: _dafny.Map
            d_3138_v74_ = _dafny.Map({d_3046_v0_: (d_3137_v73_).cf66})
            d_3139_v75_: _dafny.Map
            d_3139_v75_ = _dafny.Map({d_3138_v74_: _dafny.Map({d_3046_v0_: d_3046_v0_})})
            d_3140_v76_: _dafny.Seq
            d_3140_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvfqohe"))
            d_3141_v77_: _dafny.Seq
            d_3141_v77_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_3046_v0_ for d_3142_i6_ in range(default__.abs(381))]))])
            d_3143_v78_: _dafny.Seq
            d_3143_v78_ = _dafny.SeqWithoutIsStrInference([d_3046_v0_, d_3046_v0_, 532, d_3046_v0_, len(d_3141_v77_)])
            d_3144_v79_: _dafny.Seq
            d_3144_v79_ = _dafny.SeqWithoutIsStrInference([d_3143_v78_])
            d_3145_v80_: C2
            nw503_ = C2()
            nw503_.ctor__(d_3046_v0_, d_3139_v75_, d_3140_v76_, (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_3046_v0_]) for d_3146_i5_ in range(default__.abs(-361))])) + (d_3144_v79_))
            d_3145_v80_ = nw503_
            if not(p2):
                d_3147_v81_: _dafny.Array
                nw504_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_3147_v81_ = nw504_
                d_3148_v82_: _dafny.Array
                nw505_ = _dafny.Array(int(0), 17)
                d_3148_v82_ = nw505_
                index447_ = default__.safeIndex(279, (d_3147_v81_).length(0))
                (d_3147_v81_)[index447_] = d_3148_v82_
                index448_ = default__.safeIndex(279, (d_3147_v81_).length(0))
                nw506_ = _dafny.Array(int(0), 16)
                (d_3147_v81_)[index448_] = nw506_
                r0 = p2
                d_3046_v0_ = d_3046_v0_
                d_3149_v83_: _dafny.Map
                d_3149_v83_ = _dafny.Map({default__.fm8(d_3046_v0_, p2, p2, globalState): (self).f14})
                d_3150_v84_: _dafny.MultiSet
                d_3150_v84_ = _dafny.MultiSet([(self).f14, (((d_3149_v83_)[False] if (False) in (d_3149_v83_) else (self).f14) if p2 else (self).f14), (self).f14])
                rhs458_ = (d_3150_v84_) - (d_3150_v84_)
                rhs459_ = d_3046_v0_
                lhs296_ = globalState
                d_3150_v84_ = rhs458_
                lhs296_.f3 = rhs459_
                r0 = False
            elif True:
                d_3151_v85_: _dafny.Map
                d_3151_v85_ = _dafny.Map({((0) - (-538)) * (16): p1})
                d_3151_v85_ = (d_3151_v85_).set((0) - ((0) - (d_3046_v0_)), p1)
                d_3152_v86_: C0
                nw507_ = C0()
                nw507_.ctor__(not(p2), not(p1), d_3140_v76_, d_3144_v79_)
                d_3152_v86_ = nw507_
                d_3153_v87_: D14
                d_3153_v87_ = D14_DC32(d_3152_v86_)
                d_3154_v88_: D14
                d_3154_v88_ = D14_DC32((d_3153_v87_).cf39)
                d_3153_v87_ = d_3153_v87_
                d_3155_v89_: _dafny.Array
                d_3156_v90_: _dafny.Array
                out83_: _dafny.Array
                out84_: _dafny.Array
                out83_, out84_ = (d_3145_v80_).m3(d_3046_v0_, d_3140_v76_, (len((d_3140_v76_).set(default__.safeIndex(d_3046_v0_, len(d_3140_v76_)), d_3067_v20_)) if (d_3152_v86_).f26 else (_dafny.MultiSet([d_3065_v18_, d_3065_v18_])).cardinality), d_3140_v76_, globalState)
                d_3155_v89_ = out83_
                d_3156_v90_ = out84_
                d_3140_v76_ = ((d_3140_v76_ if (d_3152_v86_).f25 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvnrttc"))) if p1 else d_3140_v76_)
                (globalState).f3 = len(d_3064_v17_)
            d_3046_v0_ = ((0) - (d_3046_v0_)) * (d_3046_v0_)
            d_3157_v91_: _dafny.Map
            d_3157_v91_ = _dafny.Map({p2: p1})
            d_3158_v92_: _dafny.MultiSet
            d_3158_v92_ = _dafny.MultiSet([len(d_3157_v91_), d_3046_v0_, d_3046_v0_, 962, (0) - (d_3046_v0_)])
            r0 = (d_3158_v92_).ispropersubset(d_3158_v92_)
        d_3159_v94_: _dafny.Map
        d_3159_v94_ = _dafny.Map({d_3046_v0_: d_3046_v0_})
        d_3160_v95_: _dafny.MultiSet
        d_3160_v95_ = _dafny.MultiSet([690, (0) - (d_3046_v0_)])
        d_3161_v96_: _dafny.Set
        d_3161_v96_ = _dafny.Set({d_3159_v94_, _dafny.Map({d_3046_v0_: ((d_3160_v95_)[d_3046_v0_] if (d_3046_v0_) in (d_3160_v95_) else d_3046_v0_)})})
        d_3162_v97_: _dafny.Seq
        d_3162_v97_ = _dafny.SeqWithoutIsStrInference([d_3046_v0_, d_3046_v0_, -451])
        d_3163_v98_: _dafny.Seq
        d_3163_v98_ = _dafny.SeqWithoutIsStrInference([d_3162_v97_])
        d_3164_v99_: T1
        nw508_ = C1()
        def iife227_():
            coll97_ = _dafny.Map()
            compr_97_: _dafny.Map
            for compr_97_ in (d_3161_v96_).Elements:
                d_3165_v93_: _dafny.Map = compr_97_
                if (d_3165_v93_) in (d_3161_v96_):
                    coll97_[d_3165_v93_] = d_3159_v94_
            return _dafny.Map(coll97_)
        nw508_.ctor__(d_3046_v0_, d_3046_v0_, d_3046_v0_, iife227_()
        , _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "joqrcovc")), d_3163_v98_)
        d_3164_v99_ = nw508_
        d_3166_v100_: _dafny.Seq
        d_3166_v100_ = _dafny.SeqWithoutIsStrInference([d_3164_v99_])
        d_3167_v101_: _dafny.Map
        d_3167_v101_ = _dafny.Map({(self).f14: d_3166_v100_})
        d_3167_v101_ = (d_3167_v101_).set((self).f14, d_3166_v100_)
        if (default__.fm69(default__.fm22(p1, d_3046_v0_, globalState), (d_3164_v99_).f9, globalState)).cf51:
            d_3168_v102_: _dafny.Map
            d_3168_v102_ = _dafny.Map({d_3164_v99_: d_3046_v0_})
            d_3169_v103_: D24
            d_3169_v103_ = D24_DC59(d_3162_v97_, len(d_3168_v102_), len(d_3064_v17_))
            d_3170_v104_: D24
            d_3170_v104_ = D24_DC60(d_3169_v103_)
            d_3170_v104_ = d_3170_v104_
            r0 = (p2) or ((816) < ((d_3162_v97_)[default__.safeIndex(len((d_3164_v99_).f9), len(d_3162_v97_))]))
            d_3067_v20_ = d_3067_v20_
            if default__.fm38((d_3164_v99_).f11, globalState):
                d_3171_v105_: _dafny.Array
                nw509_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
                d_3171_v105_ = nw509_
                d_3171_v105_ = (d_3171_v105_ if p1 else d_3171_v105_)
                index449_ = default__.safeIndex(495, (d_3171_v105_).length(0))
                (d_3171_v105_)[index449_] = (d_3164_v99_).f9
                index450_ = default__.safeIndex(495, (d_3171_v105_).length(0))
                (d_3171_v105_)[index450_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhtvkeqx"))
                d_3046_v0_ = (len((d_3164_v99_).f9)) + ((d_3164_v99_).f11)
                d_3172_v106_: C8
                nw510_ = C8()
                nw510_.ctor__(d_3046_v0_, p1)
                d_3172_v106_ = nw510_
                d_3173_v107_: _dafny.Array
                nw511_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_3173_v107_ = nw511_
                d_3174_v108_: _dafny.MultiSet
                d_3174_v108_ = _dafny.MultiSet([p1])
                d_3175_v109_: _dafny.Array
                nw512_ = _dafny.Array(None, 16)
                nw512_[int(0)] = (d_3164_v99_).f11
                nw512_[int(1)] = 316
                nw512_[int(2)] = (d_3164_v99_).f11
                nw512_[int(3)] = (d_3164_v99_).f11
                nw512_[int(4)] = (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxtvrec")), (d_3164_v99_).f9])).cardinality
                nw512_[int(5)] = (d_3164_v99_).f11
                nw512_[int(6)] = len(p0)
                nw512_[int(7)] = d_3046_v0_
                nw512_[int(8)] = (d_3164_v99_).f11
                nw512_[int(9)] = (d_3174_v108_).cardinality
                nw512_[int(10)] = (d_3172_v106_).f15
                nw512_[int(11)] = (d_3164_v99_).f11
                nw512_[int(12)] = len(p0)
                nw512_[int(13)] = -673
                nw512_[int(14)] = 846
                nw512_[int(15)] = len((d_3164_v99_).f9)
                d_3175_v109_ = nw512_
                index451_ = default__.safeIndex(508, (d_3173_v107_).length(0))
                (d_3173_v107_)[index451_] = d_3175_v109_
                index452_ = default__.safeIndex(508, (d_3173_v107_).length(0))
                (d_3173_v107_)[index452_] = d_3175_v109_
            elif True:
                r0 = p2
                d_3176_v110_: _dafny.MultiSet
                d_3176_v110_ = _dafny.MultiSet([p2, True])
                d_3177_v111_: _dafny.Map
                d_3177_v111_ = _dafny.Map({(d_3164_v99_).f11: d_3176_v110_})
                d_3177_v111_ = (d_3177_v111_).set((0) - ((d_3164_v99_).f11), _dafny.MultiSet([p1]))
                d_3178_v112_: _dafny.MultiSet
                d_3178_v112_ = _dafny.MultiSet([(self).f14])
                d_3179_v113_: _dafny.Seq
                d_3179_v113_ = _dafny.SeqWithoutIsStrInference([d_3178_v112_, d_3178_v112_])
                r0 = ((d_3179_v113_)[default__.safeIndex((d_3164_v99_).f11, len(d_3179_v113_))]).issubset(((d_3178_v112_).set((self).f14, default__.abs((d_3164_v99_).f11))) - (d_3178_v112_))
                (globalState).f3 = (0) - ((self).fm7((d_3164_v99_).f11, (0) - ((d_3164_v99_).f11), globalState))
                d_3180_v114_: _dafny.Map
                d_3180_v114_ = _dafny.Map({d_3067_v20_: p2})
                d_3181_v116_: _dafny.Array
                nw513_ = _dafny.Array(None, 24)
                nw513_[int(0)] = (d_3164_v99_).f11
                nw513_[int(1)] = d_3046_v0_
                nw513_[int(2)] = 345
                nw513_[int(3)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfylrpmlt")))
                nw513_[int(4)] = ((d_3160_v95_) | (d_3160_v95_)).cardinality
                nw513_[int(5)] = (len((d_3164_v99_).f9)) * ((d_3164_v99_).f11)
                nw513_[int(6)] = (d_3046_v0_) + ((d_3164_v99_).f11)
                nw513_[int(7)] = (0) - ((0) - ((d_3164_v99_).f11))
                nw513_[int(8)] = -48
                nw513_[int(9)] = (d_3164_v99_).f11
                nw513_[int(10)] = default__.safeModuloInt(88, (d_3164_v99_).f11)
                nw513_[int(11)] = d_3046_v0_
                nw513_[int(12)] = 855
                nw513_[int(13)] = (d_3046_v0_) + ((d_3164_v99_).f11)
                nw513_[int(14)] = 583
                nw513_[int(15)] = default__.safeDivisionInt(len((d_3180_v114_).set(d_3067_v20_, p2)), len(default__.fm37(542, (d_3164_v99_).f11, not(p1), (0) - (d_3046_v0_), globalState)))
                nw513_[int(16)] = (0) - ((d_3164_v99_).f11)
                nw513_[int(17)] = d_3046_v0_
                nw513_[int(18)] = (d_3164_v99_).f11
                nw513_[int(19)] = -948
                nw513_[int(20)] = 713
                nw513_[int(21)] = d_3046_v0_
                def iife228_():
                    coll98_ = _dafny.Map()
                    compr_98_: int
                    for compr_98_ in _dafny.IntegerRange(864, 598):
                        d_3182_v115_: int = compr_98_
                        if ((864) <= (d_3182_v115_)) and ((d_3182_v115_) < (598)):
                            coll98_[(d_3182_v115_) * (d_3046_v0_)] = (0) - (d_3046_v0_)
                    return _dafny.Map(coll98_)
                nw513_[int(22)] = len((iife228_()
                 if p2 else d_3159_v94_))
                nw513_[int(23)] = default__.safeDivisionInt((d_3164_v99_).f11, (0) - ((d_3160_v95_).cardinality))
                d_3181_v116_ = nw513_
                index453_ = default__.safeIndex(777, (d_3181_v116_).length(0))
                (d_3181_v116_)[index453_] = (d_3164_v99_).f11
                index454_ = default__.safeIndex(777, (d_3181_v116_).length(0))
                (d_3181_v116_)[index454_] = d_3046_v0_
            d_3183_v117_: _dafny.Array
            nw514_ = _dafny.Array(None, 11)
            nw514_[int(0)] = (d_3164_v99_).f11
            nw514_[int(1)] = (d_3164_v99_).f11
            nw514_[int(2)] = (d_3164_v99_).f11
            nw514_[int(3)] = (d_3164_v99_).f11
            nw514_[int(4)] = (d_3164_v99_).f11
            nw514_[int(5)] = d_3046_v0_
            nw514_[int(6)] = len((d_3164_v99_).f9)
            nw514_[int(7)] = (d_3160_v95_).cardinality
            nw514_[int(8)] = (d_3164_v99_).f11
            nw514_[int(9)] = (d_3164_v99_).f11
            nw514_[int(10)] = (d_3164_v99_).f11
            d_3183_v117_ = nw514_
            d_3184_v118_: _dafny.Map
            d_3184_v118_ = _dafny.Map({d_3183_v117_: (d_3164_v99_).f9})
            if default__.fm13(not(p1), 235, ((d_3184_v118_)[d_3183_v117_] if (d_3183_v117_) in (d_3184_v118_) else (d_3164_v99_).f9), globalState):
                d_3185_v119_: _dafny.Map
                d_3185_v119_ = _dafny.Map({d_3067_v20_: p1})
                d_3186_v120_: _dafny.MultiSet
                d_3186_v120_ = _dafny.MultiSet([d_3185_v119_])
                d_3187_v121_: _dafny.Seq
                d_3187_v121_ = default__.fm31(globalState)
                d_3188_v122_: _dafny.Array
                d_3189_v123_: _dafny.Array
                out85_: _dafny.Array
                out86_: _dafny.Array
                out85_, out86_ = (d_3164_v99_).m3((0) - ((d_3186_v120_).cardinality), default__.fm31(globalState), (d_3164_v99_).f11, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocxakoddq")) if True else d_3187_v121_), globalState)
                d_3188_v122_ = out85_
                d_3189_v123_ = out86_
                d_3190_v124_: _dafny.Array
                nw515_ = _dafny.Array(None, 25)
                nw515_[int(0)] = ((d_3164_v99_).f9) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_3191_i7_ in range(default__.abs(757))]))
                nw515_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vuctiil"))
                nw515_[int(2)] = (_dafny.SeqWithoutIsStrInference([d_3067_v20_ for d_3192_i8_ in range(default__.abs(-456))])) + ((d_3164_v99_).f9)
                nw515_[int(3)] = (d_3164_v99_).f9
                nw515_[int(4)] = (d_3164_v99_).f9
                nw515_[int(5)] = (((d_3164_v99_).f9).set(default__.safeIndex((0) - ((d_3162_v97_)[default__.safeIndex((d_3164_v99_).f11, len(d_3162_v97_))]), len((d_3164_v99_).f9)), _dafny.CodePoint('i')) if p2 else (d_3164_v99_).f9)
                nw515_[int(6)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_3193_i9_ in range(default__.abs(148))])) + ((d_3164_v99_).f9)
                nw515_[int(7)] = (d_3164_v99_).f9
                nw515_[int(8)] = default__.fm26((d_3164_v99_).f11, _dafny.Map({(d_3164_v99_).f11: ((d_3164_v99_).f9).set(default__.safeIndex((_dafny.MultiSet([884])).cardinality, len((d_3164_v99_).f9)), d_3067_v20_)}), globalState)
                nw515_[int(9)] = (d_3164_v99_).f9
                nw515_[int(10)] = (d_3164_v99_).f9
                nw515_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eensas"))
                nw515_[int(12)] = (d_3164_v99_).f9
                nw515_[int(13)] = (d_3164_v99_).f9
                nw515_[int(14)] = (_dafny.SeqWithoutIsStrInference([d_3067_v20_ for d_3194_i10_ in range(default__.abs(938))])) + ((d_3164_v99_).f9)
                nw515_[int(15)] = ((d_3164_v99_).f9) + ((d_3164_v99_).f9)
                nw515_[int(16)] = ((d_3164_v99_).f9) + ((d_3164_v99_).f9)
                nw515_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmqrjxe"))
                nw515_[int(18)] = _dafny.SeqWithoutIsStrInference([d_3067_v20_ for d_3195_i11_ in range(default__.abs(234))])
                nw515_[int(19)] = (d_3164_v99_).f9
                nw515_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mg"))
                nw515_[int(21)] = (d_3164_v99_).f9
                nw515_[int(22)] = (d_3164_v99_).f9
                nw515_[int(23)] = (d_3164_v99_).f9
                nw515_[int(24)] = (d_3164_v99_).f9
                d_3190_v124_ = nw515_
                index455_ = default__.safeIndex(121, (d_3190_v124_).length(0))
                (d_3190_v124_)[index455_] = (d_3164_v99_).f9
                d_3196_v125_: _dafny.Set
                d_3196_v125_ = _dafny.Set({p1})
                d_3197_v126_: _dafny.Seq
                d_3197_v126_ = _dafny.SeqWithoutIsStrInference([(d_3196_v125_).issubset(d_3196_v125_), p1, p2])
                d_3198_v127_: D8
                d_3198_v127_ = D8_DC17((d_3197_v126_).set(default__.safeIndex(len(d_3162_v97_), len(d_3197_v126_)), p1))
                index456_ = default__.safeIndex(121, (d_3190_v124_).length(0))
                rhs460_ = default__.safeDivisionInt((d_3164_v99_).f11, (d_3164_v99_).f11)
                rhs461_ = (((0) - (len(d_3162_v97_))) != (d_3046_v0_)) in (_dafny.Set({not(p1)}))
                rhs462_ = (default__.fm31(globalState)).set(default__.safeIndex((d_3164_v99_).f11, len(default__.fm31(globalState))), _dafny.CodePoint('j'))
                rhs463_ = ((_dafny.SeqWithoutIsStrInference([False, p2]) if p2 else (d_3198_v127_).cf19)) + (d_3197_v126_)
                lhs297_ = globalState
                lhs298_ = d_3190_v124_
                lhs299_ = default__.safeIndex(121, (d_3190_v124_).length(0))
                lhs297_.f3 = rhs460_
                r0 = rhs461_
                lhs298_[lhs299_] = rhs462_
                d_3197_v126_ = rhs463_
                d_3199_v128_: _dafny.Map
                d_3199_v128_ = _dafny.Map({(d_3164_v99_).f11: p2})
                d_3199_v128_ = _dafny.Map({default__.safeModuloInt((d_3164_v99_).f11, d_3046_v0_): p2})
                d_3200_v129_: _dafny.MultiSet
                d_3200_v129_ = _dafny.MultiSet([not(p2)])
                d_3201_v130_: D12
                d_3201_v130_ = D12_DC29(p2, (d_3200_v129_).cardinality, ((d_3164_v99_).f10)[default__.safeIndex((d_3164_v99_).f11, len((d_3164_v99_).f10))], p2, len(p0))
                r0 = (d_3201_v130_).cf35
                d_3202_v131_: C8
                nw516_ = C8()
                nw516_.ctor__((d_3164_v99_).f11, p1)
                d_3202_v131_ = nw516_
            elif True:
                d_3203_v132_: D8
                d_3203_v132_ = D8_DC18()
                d_3204_v133_: D8
                d_3204_v133_ = D8_DC19(d_3203_v132_)
                d_3205_v134_: _dafny.Map
                d_3205_v134_ = _dafny.Map({d_3183_v117_: d_3204_v133_})
                d_3206_v135_: D6
                d_3206_v135_ = D6_DC14(d_3183_v117_)
                d_3205_v134_ = (d_3205_v134_).set((d_3206_v135_).cf17, d_3204_v133_)
                d_3207_v136_: _dafny.Seq
                d_3207_v136_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "da"))
                d_3207_v136_ = ((d_3164_v99_).f9) + ((d_3164_v99_).f9)
                (globalState).f3 = (d_3164_v99_).f11
                d_3208_v137_: _dafny.Seq
                d_3208_v137_ = _dafny.SeqWithoutIsStrInference([(d_3164_v99_).f9, (d_3164_v99_).f9])
                d_3209_v138_: C1
                nw517_ = C1()
                nw517_.ctor__((d_3164_v99_).f11, default__.safeDivisionInt(d_3046_v0_, len((d_3208_v137_)[default__.safeIndex((d_3164_v99_).f11, len(d_3208_v137_))])), ((d_3164_v99_).f11) + ((d_3164_v99_).f11), (_dafny.Map({d_3159_v94_: d_3159_v94_})) | (d_3164_v99_.f12), ((d_3164_v99_).f9) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfeh"))), d_3163_v98_)
                d_3209_v138_ = nw517_
                (globalState).f3 = -90
        elif True:
            d_3210_v139_: D20
            d_3210_v139_ = D20_DC47(p1, (d_3164_v99_).f11, p1)
            d_3211_v140_: _dafny.Seq
            d_3211_v140_ = _dafny.SeqWithoutIsStrInference([(d_3164_v99_).f10, d_3163_v98_])
            d_3212_v141_: C0
            nw518_ = C0()
            nw518_.ctor__(((0) - (d_3046_v0_)) == ((d_3164_v99_).f11), (len(d_3166_v100_)) != ((d_3210_v139_).cf66), (((d_3164_v99_).f9).set(default__.safeIndex(d_3046_v0_, len((d_3164_v99_).f9)), d_3067_v20_)) + (((d_3164_v99_).f9).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvgf"))), len((d_3164_v99_).f9)), d_3067_v20_)), (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_3164_v99_).f11 for d_3213_i12_ in range(default__.abs(-434))]), _dafny.SeqWithoutIsStrInference([len(d_3064_v17_), d_3046_v0_, (d_3164_v99_).f11, (d_3164_v99_).f11]), _dafny.SeqWithoutIsStrInference([len(d_3064_v17_), (0) - ((d_3164_v99_).f11), len((d_3164_v99_).f9), (d_3164_v99_).f11, (d_3164_v99_).f11])])) + ((d_3211_v140_)[default__.safeIndex((d_3164_v99_).f11, len(d_3211_v140_))]))
            d_3212_v141_ = nw518_
            index457_ = default__.safeIndex(389, ((self).f14).length(0))
            ((self).f14)[index457_] = (d_3212_v141_).f25
            d_3214_v142_: _dafny.Seq
            d_3214_v142_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xhfe"))
            index458_ = default__.safeIndex(389, ((self).f14).length(0))
            rhs464_ = (d_3046_v0_) != (((d_3212_v141_).fm5((d_3164_v99_).f11, 106, globalState) if p2 else 359))
            rhs465_ = (d_3067_v20_) in ((d_3214_v142_) + (d_3214_v142_))
            rhs466_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nj"))) + (((d_3164_v99_).f9) + ((d_3164_v99_).f9))
            lhs300_ = (self).f14
            lhs301_ = default__.safeIndex(389, ((self).f14).length(0))
            r0 = rhs464_
            lhs300_[lhs301_] = rhs465_
            d_3214_v142_ = rhs466_
            (globalState).f3 = (d_3164_v99_).f11
            d_3215_v143_: _dafny.Set
            d_3215_v143_ = _dafny.Set({(d_3212_v141_).f26, (d_3212_v141_).f25, (d_3212_v141_).f26, (p2) == (not(p2)), p2})
            rhs467_ = ((0) - (default__.safeModuloInt(len(d_3159_v94_), (d_3164_v99_).f11))) == (-798)
            rhs468_ = (0) - (len(d_3215_v143_))
            rhs469_ = (((d_3160_v95_).cardinality) * ((d_3164_v99_).f11) if ((self).f14)[default__.safeIndex(389, ((self).f14).length(0))] else 332)
            rhs470_ = (d_3212_v141_).f26
            lhs302_ = globalState
            lhs303_ = globalState
            r0 = rhs467_
            lhs302_.f3 = rhs468_
            lhs303_.f3 = rhs469_
            r0 = rhs470_
            index459_ = default__.safeIndex(389, ((self).f14).length(0))
            ((self).f14)[index459_] = ((self).f14)[default__.safeIndex(389, ((self).f14).length(0))]
        hi12_ = d_3046_v0_
        for d_3216_i13_ in range(399, hi12_):
            index460_ = default__.safeIndex(777, ((self).f14).length(0))
            ((self).f14)[index460_] = p1
            d_3217_v144_: D14
            d_3217_v144_ = D14_DC34(d_3046_v0_, (d_3164_v99_).f11, (d_3164_v99_).f11, p1)
            index461_ = default__.safeIndex(777, ((self).f14).length(0))
            rhs471_ = (d_3217_v144_).cf46
            rhs472_ = p2
            lhs304_ = (self).f14
            lhs305_ = default__.safeIndex(777, ((self).f14).length(0))
            r0 = rhs471_
            lhs304_[lhs305_] = rhs472_
            d_3064_v17_ = (d_3064_v17_).set((d_3160_v95_).issubset(d_3160_v95_), d_3046_v0_)
            d_3218_v145_: _dafny.Array
            def lambda143_(d_3219_v0_):
                def lambda144_(d_3220_i14_):
                    return default__.safeDivisionInt(d_3220_i14_, d_3219_v0_)

                return lambda144_

            init82_ = lambda143_(d_3046_v0_)
            nw519_ = _dafny.Array(None, 13)
            for i0_82_ in range(nw519_.length(0)):
                nw519_[i0_82_] = init82_(i0_82_)
            d_3218_v145_ = nw519_
            d_3221_v147_: _dafny.Set
            def iife229_():
                coll99_ = _dafny.Map()
                compr_99_: int
                for compr_99_ in _dafny.IntegerRange(839, 865):
                    d_3222_v146_: int = compr_99_
                    if ((839) <= (d_3222_v146_)) and ((d_3222_v146_) < (865)):
                        coll99_[(d_3222_v146_) + (d_3216_i13_)] = p2
                return _dafny.Map(coll99_)
            d_3221_v147_ = _dafny.Set({iife229_()
            })
            d_3223_v148_: _dafny.Map
            d_3223_v148_ = _dafny.Map({(0) - (d_3046_v0_): d_3221_v147_})
            d_3224_v149_: _dafny.Map
            d_3224_v149_ = _dafny.Map({(d_3164_v99_).f11: p1})
            d_3225_v150_: _dafny.Map
            d_3225_v150_ = _dafny.Map({d_3224_v149_: len(d_3162_v97_)})
            d_3226_v152_: D23
            def iife230_():
                coll100_ = _dafny.Set()
                compr_100_: _dafny.Map
                for compr_100_ in (d_3225_v150_).keys.Elements:
                    d_3227_v151_: _dafny.Map = compr_100_
                    if (d_3227_v151_) in (d_3225_v150_):
                        coll100_ = coll100_.union(_dafny.Set([d_3227_v151_]))
                return _dafny.Set(coll100_)
            d_3226_v152_ = D23_DC55(((d_3223_v148_)[len((d_3224_v149_).set((d_3164_v99_).f11, p1))] if (len((d_3224_v149_).set((d_3164_v99_).f11, p1))) in (d_3223_v148_) else iife230_()
))
            pat_let_tv63_ = d_3221_v147_
            pat_let_tv64_ = d_3221_v147_
            d_3228_v153_: _dafny.Array
            nw520_ = _dafny.Array(None, 22)
            nw520_[int(0)] = D23_DC55(d_3221_v147_)
            nw520_[int(1)] = d_3226_v152_
            nw520_[int(2)] = d_3226_v152_
            def iife232_(_pat_let66_0):
                def iife233_(d_3229_dt__update__tmp_h1_):
                    def iife234_(_pat_let67_0):
                        def iife235_(d_3230_dt__update_hcf82_h0_):
                            return D23_DC55(d_3230_dt__update_hcf82_h0_)
                        return iife235_(_pat_let67_0)
                    return iife234_(pat_let_tv63_)
                return iife233_(_pat_let66_0)
            def iife231_(_pat_let65_0):
                def iife236_(d_3231_dt__update__tmp_h2_):
                    def iife237_(_pat_let68_0):
                        def iife238_(d_3232_dt__update_hcf82_h1_):
                            return D23_DC55(d_3232_dt__update_hcf82_h1_)
                        return iife238_(_pat_let68_0)
                    return iife237_(pat_let_tv64_)
                return iife236_(_pat_let65_0)
            nw520_[int(3)] = iife231_(iife232_(d_3226_v152_))
            nw520_[int(4)] = d_3226_v152_
            nw520_[int(5)] = D23_DC55(_dafny.Set({d_3224_v149_}))
            nw520_[int(6)] = d_3226_v152_
            nw520_[int(7)] = D23_DC55(d_3221_v147_)
            nw520_[int(8)] = d_3226_v152_
            nw520_[int(9)] = d_3226_v152_
            nw520_[int(10)] = d_3226_v152_
            nw520_[int(11)] = D23_DC55(d_3221_v147_)
            nw520_[int(12)] = d_3226_v152_
            nw520_[int(13)] = d_3226_v152_
            nw520_[int(14)] = d_3226_v152_
            nw520_[int(15)] = d_3226_v152_
            nw520_[int(16)] = d_3226_v152_
            nw520_[int(17)] = d_3226_v152_
            nw520_[int(18)] = d_3226_v152_
            nw520_[int(19)] = d_3226_v152_
            nw520_[int(20)] = d_3226_v152_
            nw520_[int(21)] = d_3226_v152_
            d_3228_v153_ = nw520_
            index462_ = default__.safeIndex(45, (d_3228_v153_).length(0))
            (d_3228_v153_)[index462_] = (d_3226_v152_ if False else d_3226_v152_)
            index463_ = default__.safeIndex(45, (d_3228_v153_).length(0))
            rhs473_ = d_3218_v145_
            rhs474_ = d_3226_v152_
            rhs475_ = (0) - (len(((((d_3164_v99_).f9) + ((d_3164_v99_).f9)).set(default__.safeIndex((d_3164_v99_).f11, len(((d_3164_v99_).f9) + ((d_3164_v99_).f9))), d_3067_v20_)) + ((d_3164_v99_).f9)))
            lhs306_ = d_3228_v153_
            lhs307_ = default__.safeIndex(45, (d_3228_v153_).length(0))
            lhs308_ = globalState
            d_3218_v145_ = rhs473_
            lhs306_[lhs307_] = rhs474_
            lhs308_.f3 = rhs475_
            d_3233_v154_: _dafny.Map
            d_3233_v154_ = _dafny.Map({d_3067_v20_: ((self).f14)[default__.safeIndex(777, ((self).f14).length(0))]})
            d_3233_v154_ = (D27_DC63(d_3233_v154_)).cf96
        r0 = p1
        r1 = ((d_3164_v99_).f11) + (d_3046_v0_)
        return r0, r1

    @property
    def f14(self):
        return self._f14

class C10(T1, T0):
    def  __init__(self):
        self._f12: _dafny.Map = _dafny.Map({})
        self._f11: int = int(0)
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: _dafny.Seq = _dafny.Seq({})
        self.f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f11(self):
        return self._f11
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f13, f11, f12, f9, f10):
        (self).f13 = f13
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f9 = f9
        (self)._f10 = f10

    def fm5(self, p0, p1, globalState):
        if (-918) <= (len(_dafny.Set({(self).f9}))):
            return len((_dafny.Set({True, False, True, False})) | (_dafny.Set({True})))
        elif True:
            return (self).f11
        elif True:
            return self.f13

    def fm6(self, globalState):
        if not(((self).f11) <= (self.f13)):
            return ((self).f9)[default__.safeIndex(106, len((self).f9))]
        elif True:
            return _dafny.CodePoint('k')

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: bool = False
        r3: int = int(0)
        d_3234_v0_: bool
        d_3234_v0_ = False
        if d_3234_v0_:
            d_3235_v1_: _dafny.Array
            def lambda145_(d_3236_i0_):
                return _dafny.SeqWithoutIsStrInference([-813])

            init83_ = lambda145_
            nw521_ = _dafny.Array(None, 6)
            for i0_83_ in range(nw521_.length(0)):
                nw521_[i0_83_] = init83_(i0_83_)
            d_3235_v1_ = nw521_
            d_3237_v2_: _dafny.Map
            d_3237_v2_ = _dafny.Map({d_3234_v0_: d_3235_v1_})
            d_3237_v2_ = (d_3237_v2_).set(d_3234_v0_, d_3235_v1_)
            d_3238_v3_: _dafny.Seq
            d_3238_v3_ = _dafny.SeqWithoutIsStrInference([d_3234_v0_, d_3234_v0_])
            d_3239_v4_: _dafny.Map
            d_3239_v4_ = _dafny.Map({len(d_3238_v3_): d_3234_v0_})
            d_3240_v5_: _dafny.Map
            d_3240_v5_ = _dafny.Map({(self).f11: p0})
            d_3241_v6_: _dafny.Map
            d_3241_v6_ = _dafny.Map({(self).fm5(self.f13, len(d_3239_v4_), globalState): d_3240_v5_})
            d_3242_v7_: _dafny.MultiSet
            d_3242_v7_ = _dafny.MultiSet([(self.f13) in (d_3241_v6_)])
            d_3243_v8_: _dafny.Map
            d_3243_v8_ = _dafny.Map({d_3234_v0_: d_3234_v0_})
            d_3242_v7_ = _dafny.MultiSet([((self).f11) < (p0), d_3234_v0_, d_3234_v0_, ((d_3243_v8_).set(d_3234_v0_, d_3234_v0_)) not in (_dafny.Map({_dafny.Map({d_3234_v0_: d_3234_v0_}): d_3234_v0_})), ((self).f11) != (p0)])
            d_3244_v9_: C6
            nw522_ = C6()
            nw522_.ctor__((self).f11, self.f12, default__.fm37(self.f13, (self).f11, d_3234_v0_, p0, globalState), (self).f10)
            d_3244_v9_ = nw522_
            d_3234_v0_ = d_3234_v0_
            (globalState).f3 = p0
        elif True:
            d_3245_v10_: str
            d_3245_v10_ = _dafny.CodePoint('e')
            d_3246_v11_: _dafny.Seq
            d_3246_v11_ = _dafny.SeqWithoutIsStrInference([d_3234_v0_])
            d_3247_v12_: D7
            d_3247_v12_ = D7_DC15(_dafny.MultiSet([len(d_3246_v11_)]))
            d_3248_v13_: _dafny.Set
            d_3248_v13_ = _dafny.Set({(d_3247_v12_).cf18})
            d_3249_v14_: D10
            d_3249_v14_ = D10_DC23(d_3245_v10_, d_3248_v13_)
            d_3250_v15_: _dafny.Array
            nw523_ = _dafny.Array(None, 26)
            nw523_[int(0)] = D10_DC23(d_3245_v10_, d_3248_v13_)
            nw523_[int(1)] = d_3249_v14_
            nw523_[int(2)] = d_3249_v14_
            nw523_[int(3)] = d_3249_v14_
            nw523_[int(4)] = d_3249_v14_
            nw523_[int(5)] = d_3249_v14_
            nw523_[int(6)] = d_3249_v14_
            nw523_[int(7)] = d_3249_v14_
            nw523_[int(8)] = d_3249_v14_
            nw523_[int(9)] = d_3249_v14_
            nw523_[int(10)] = D10_DC23(d_3245_v10_, _dafny.Set({_dafny.MultiSet([p0])}))
            nw523_[int(11)] = d_3249_v14_
            nw523_[int(12)] = d_3249_v14_
            nw523_[int(13)] = D10_DC23(d_3245_v10_, d_3248_v13_)
            nw523_[int(14)] = d_3249_v14_
            nw523_[int(15)] = d_3249_v14_
            nw523_[int(16)] = d_3249_v14_
            nw523_[int(17)] = d_3249_v14_
            nw523_[int(18)] = d_3249_v14_
            nw523_[int(19)] = d_3249_v14_
            nw523_[int(20)] = d_3249_v14_
            nw523_[int(21)] = d_3249_v14_
            nw523_[int(22)] = D10_DC23(d_3245_v10_, d_3248_v13_)
            nw523_[int(23)] = d_3249_v14_
            nw523_[int(24)] = d_3249_v14_
            nw523_[int(25)] = d_3249_v14_
            d_3250_v15_ = nw523_
            d_3251_v16_: _dafny.Array
            d_3251_v16_ = d_3250_v15_
            d_3251_v16_ = d_3251_v16_
            d_3252_v17_: D18
            d_3252_v17_ = D18_DC42(d_3234_v0_, (len(d_3246_v11_) if d_3234_v0_ else (self).f11))
            d_3253_v18_: _dafny.Map
            d_3253_v18_ = _dafny.Map({d_3234_v0_: _dafny.SeqWithoutIsStrInference([d_3234_v0_])})
            rhs476_ = (d_3252_v17_ if not(not (d_3234_v0_) or (False)) else d_3252_v17_)
            rhs477_ = (p0) + ((self).f11)
            rhs478_ = d_3253_v18_
            rhs479_ = d_3234_v0_
            rhs480_ = (d_3234_v0_) == (d_3234_v0_)
            lhs309_ = self
            d_3252_v17_ = rhs476_
            lhs309_.f13 = rhs477_
            d_3253_v18_ = rhs478_
            d_3234_v0_ = rhs479_
            r2 = rhs480_
            d_3254_v19_: _dafny.Map
            d_3254_v19_ = _dafny.Map({d_3234_v0_: d_3234_v0_})
            d_3255_v20_: D8
            d_3255_v20_ = D8_DC17(_dafny.SeqWithoutIsStrInference([d_3234_v0_, d_3234_v0_, d_3234_v0_, d_3234_v0_, ((d_3254_v19_)[d_3234_v0_] if (d_3234_v0_) in (d_3254_v19_) else d_3234_v0_)]))
            source51_ = d_3255_v20_
            if source51_.is_DC18:
                d_3256_v21_: _dafny.Seq
                d_3256_v21_ = _dafny.SeqWithoutIsStrInference([d_3245_v10_, d_3245_v10_, d_3245_v10_, _dafny.CodePoint('g')])
                d_3256_v21_ = d_3256_v21_
                d_3257_v22_: _dafny.Array
                nw524_ = _dafny.Array(int(0), 8)
                d_3257_v22_ = nw524_
                d_3257_v22_ = d_3257_v22_
                d_3258_v23_: _dafny.Array
                nw525_ = _dafny.Array(D6.default()(), 3)
                d_3258_v23_ = nw525_
                d_3259_v24_: _dafny.Map
                d_3259_v24_ = _dafny.Map({_dafny.MultiSet([d_3234_v0_, d_3234_v0_]): d_3234_v0_})
                d_3260_v25_: _dafny.Map
                d_3260_v25_ = _dafny.Map({d_3259_v24_: d_3258_v23_})
                d_3261_v26_: _dafny.MultiSet
                d_3261_v26_ = _dafny.MultiSet([d_3234_v0_])
                d_3258_v23_ = ((d_3260_v25_)[_dafny.Map({d_3261_v26_: d_3234_v0_})] if (_dafny.Map({d_3261_v26_: d_3234_v0_})) in (d_3260_v25_) else d_3258_v23_)
                d_3262_v27_: _dafny.Map
                d_3262_v27_ = _dafny.Map({default__.fm22(d_3234_v0_, (self).f11, globalState): (self).f11})
                d_3262_v27_ = d_3262_v27_
            elif source51_.is_DC17:
                d_3263___mcc_h0_ = source51_.cf19
                d_3264_cf19_ = d_3263___mcc_h0_
                d_3265_v28_: _dafny.Seq
                d_3265_v28_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                r2 = (d_3265_v28_) <= ((d_3265_v28_) + (d_3265_v28_))
                d_3266_v29_: _dafny.MultiSet
                d_3266_v29_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_3245_v10_ for d_3267_i1_ in range(default__.abs(372))])])
                d_3268_v30_: D10
                d_3268_v30_ = D10_DC22(d_3266_v29_)
                d_3269_v31_: D10
                d_3269_v31_ = D10_DC24(d_3268_v30_)
                d_3269_v31_ = D10_DC24(d_3268_v30_)
                (globalState).f3 = (self).f11
                d_3245_v10_ = d_3245_v10_
            elif True:
                d_3270___mcc_h1_ = source51_.cf20
                d_3271_cf20_ = d_3270___mcc_h1_
                r2 = d_3234_v0_
                d_3272_v32_: _dafny.Array
                nw526_ = _dafny.Array(int(0), 10)
                d_3272_v32_ = nw526_
                d_3234_v0_ = not (((self).f9) < (default__.fm37(len((d_3246_v11_).set(default__.safeIndex(default__.fm0(p0, len(_dafny.Map({d_3272_v32_: len(_dafny.SeqWithoutIsStrInference([p0 for d_3273_i2_ in range(default__.abs(504))]))})), globalState), len(d_3246_v11_)), d_3234_v0_)), len(d_3246_v11_), d_3234_v0_, p0, globalState))) or (d_3234_v0_)
                index464_ = default__.safeIndex(729, (d_3272_v32_).length(0))
                (d_3272_v32_)[index464_] = (0) - (default__.safeDivisionInt((self).f11, p0))
                index465_ = default__.safeIndex(729, (d_3272_v32_).length(0))
                (d_3272_v32_)[index465_] = self.f13
                d_3274_v33_: D6
                d_3274_v33_ = D6_DC14(d_3272_v32_)
                d_3275_v34_: _dafny.Array
                nw527_ = _dafny.Array(None, 23)
                nw527_[int(0)] = d_3272_v32_
                nw527_[int(1)] = d_3272_v32_
                nw527_[int(2)] = d_3272_v32_
                nw527_[int(3)] = d_3272_v32_
                nw527_[int(4)] = d_3272_v32_
                nw527_[int(5)] = d_3272_v32_
                nw527_[int(6)] = d_3272_v32_
                nw527_[int(7)] = d_3272_v32_
                nw527_[int(8)] = d_3272_v32_
                nw527_[int(9)] = d_3272_v32_
                nw527_[int(10)] = d_3272_v32_
                nw527_[int(11)] = d_3272_v32_
                nw527_[int(12)] = d_3272_v32_
                nw527_[int(13)] = d_3272_v32_
                nw527_[int(14)] = d_3272_v32_
                nw527_[int(15)] = d_3272_v32_
                nw527_[int(16)] = d_3272_v32_
                nw527_[int(17)] = d_3272_v32_
                nw527_[int(18)] = d_3272_v32_
                nw527_[int(19)] = d_3272_v32_
                nw527_[int(20)] = d_3272_v32_
                nw527_[int(21)] = d_3272_v32_
                nw527_[int(22)] = (d_3274_v33_).cf17
                d_3275_v34_ = nw527_
                d_3276_v35_: _dafny.Map
                d_3276_v35_ = _dafny.Map({d_3234_v0_: d_3275_v34_})
                d_3276_v35_ = (d_3276_v35_).set(d_3234_v0_, d_3275_v34_)
            d_3277_v36_: _dafny.Seq
            d_3277_v36_ = _dafny.SeqWithoutIsStrInference([p0, (self).f11])
            (globalState).f3 = len((d_3277_v36_) + (d_3277_v36_))
            d_3278_v37_: _dafny.Array
            def lambda146_(d_3279_v0_):
                def lambda147_(d_3280_i3_):
                    return d_3279_v0_

                return lambda147_

            init84_ = lambda146_(d_3234_v0_)
            nw528_ = _dafny.Array(None, 11)
            for i0_84_ in range(nw528_.length(0)):
                nw528_[i0_84_] = init84_(i0_84_)
            d_3278_v37_ = nw528_
            index466_ = default__.safeIndex(300, (d_3278_v37_).length(0))
            (d_3278_v37_)[index466_] = d_3234_v0_
            d_3281_v38_: _dafny.Map
            d_3281_v38_ = _dafny.Map({(self).f11: d_3246_v11_})
            d_3282_v39_: _dafny.Set
            d_3282_v39_ = _dafny.Set({(len(d_3281_v38_)) - ((0) - (p0))})
            d_3283_v40_: _dafny.Set
            d_3283_v40_ = _dafny.Set({d_3254_v19_})
            d_3284_v41_: _dafny.Array
            def lambda148_(d_3285_v10_):
                def lambda149_(d_3286_i4_):
                    return d_3285_v10_

                return lambda149_

            init85_ = lambda148_(d_3245_v10_)
            nw529_ = _dafny.Array(None, 18)
            for i0_85_ in range(nw529_.length(0)):
                nw529_[i0_85_] = init85_(i0_85_)
            d_3284_v41_ = nw529_
            d_3287_v42_: D20
            d_3287_v42_ = D20_DC48(d_3234_v0_, d_3283_v40_, d_3284_v41_)
            pat_let_tv65_ = d_3284_v41_
            pat_let_tv66_ = d_3283_v40_
            d_3288_v43_: _dafny.Seq
            def iife239_(_pat_let69_0):
                def iife240_(d_3289_dt__update__tmp_h0_):
                    def iife241_(_pat_let70_0):
                        def iife242_(d_3290_dt__update_hcf70_h0_):
                            def iife243_(_pat_let71_0):
                                def iife244_(d_3291_dt__update_hcf69_h0_):
                                    return D20_DC48((d_3289_dt__update__tmp_h0_).cf68, d_3291_dt__update_hcf69_h0_, d_3290_dt__update_hcf70_h0_)
                                return iife244_(_pat_let71_0)
                            return iife243_(pat_let_tv66_)
                        return iife242_(_pat_let70_0)
                    return iife241_(pat_let_tv65_)
                return iife240_(_pat_let69_0)
            d_3288_v43_ = _dafny.SeqWithoutIsStrInference([(d_3287_v42_ if d_3234_v0_ else D20_DC48(d_3234_v0_, _dafny.Set({d_3254_v19_, d_3254_v19_}), d_3284_v41_)), iife239_(d_3287_v42_), d_3287_v42_])
            index467_ = default__.safeIndex(300, (d_3278_v37_).length(0))
            rhs481_ = ((d_3234_v0_) or (d_3234_v0_)) or (not(d_3234_v0_))
            rhs482_ = d_3282_v39_
            rhs483_ = d_3288_v43_
            rhs484_ = ((self).f9) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")))
            rhs485_ = d_3234_v0_
            lhs310_ = d_3278_v37_
            lhs311_ = default__.safeIndex(300, (d_3278_v37_).length(0))
            lhs310_[lhs311_] = rhs481_
            d_3282_v39_ = rhs482_
            d_3288_v43_ = rhs483_
            d_3234_v0_ = rhs484_
            r2 = rhs485_
        if (d_3234_v0_) and (default__.fm22(d_3234_v0_, self.f13, globalState)):
            r2 = d_3234_v0_
            d_3292_v44_: _dafny.Set
            d_3292_v44_ = _dafny.Set({p0, len(_dafny.Set({d_3234_v0_}))})
            d_3293_v45_: _dafny.MultiSet
            d_3293_v45_ = _dafny.MultiSet([d_3234_v0_])
            d_3294_v46_: _dafny.MultiSet
            d_3294_v46_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([p0, self.f13, p0, len(d_3292_v44_), self.f13])), ((d_3293_v45_)[d_3234_v0_] if (d_3234_v0_) in (d_3293_v45_) else p0)])
            d_3295_v47_: str
            d_3295_v47_ = _dafny.CodePoint('t')
            d_3296_v48_: _dafny.Map
            d_3296_v48_ = _dafny.Map({d_3295_v47_: d_3234_v0_})
            d_3297_v49_: _dafny.Map
            d_3297_v49_ = _dafny.Map({d_3234_v0_: ((d_3296_v48_)[_dafny.CodePoint('u')] if (_dafny.CodePoint('u')) in (d_3296_v48_) else d_3234_v0_)})
            d_3298_v50_: T0
            nw530_ = C7()
            nw530_.ctor__((self).f9, (self).f10, (0) - (self.f13), self.f12)
            d_3298_v50_ = nw530_
            d_3299_v51_: _dafny.Seq
            d_3299_v51_ = _dafny.SeqWithoutIsStrInference([d_3298_v50_])
            d_3300_v52_: _dafny.Seq
            d_3300_v52_ = _dafny.SeqWithoutIsStrInference([len(d_3299_v51_)])
            d_3301_v53_: _dafny.Array
            nw531_ = _dafny.Array(None, 20)
            nw531_[int(0)] = -759
            nw531_[int(1)] = (0) - ((d_3294_v46_).cardinality)
            nw531_[int(2)] = p0
            nw531_[int(3)] = default__.safeModuloInt(p0, self.f13)
            nw531_[int(4)] = default__.safeModuloInt((d_3294_v46_).cardinality, (0) - ((self).f11))
            nw531_[int(5)] = (p0) + ((_dafny.MultiSet([d_3234_v0_])).cardinality)
            nw531_[int(6)] = default__.safeDivisionInt(len(d_3297_v49_), p0)
            nw531_[int(7)] = p0
            nw531_[int(8)] = default__.safeDivisionInt(907, 408)
            nw531_[int(9)] = self.f13
            nw531_[int(10)] = p0
            nw531_[int(11)] = self.f13
            nw531_[int(12)] = (0) - (len((self).f9))
            nw531_[int(13)] = (self).f11
            nw531_[int(14)] = len((d_3300_v52_) + (d_3300_v52_))
            nw531_[int(15)] = self.f13
            nw531_[int(16)] = len(d_3300_v52_)
            nw531_[int(17)] = self.f13
            nw531_[int(18)] = self.f13
            nw531_[int(19)] = ((self).f11) * ((self).f11)
            d_3301_v53_ = nw531_
            def lambda150_(d_3302_p0_):
                def lambda151_(d_3303_i5_):
                    return (d_3303_i5_) + (d_3302_p0_)

                return lambda151_

            init86_ = lambda150_(p0)
            nw532_ = _dafny.Array(None, 14)
            for i0_86_ in range(nw532_.length(0)):
                nw532_[i0_86_] = init86_(i0_86_)
            d_3301_v53_ = nw532_
            d_3292_v44_ = d_3292_v44_
            d_3304_v54_: _dafny.Set
            d_3304_v54_ = _dafny.Set({d_3234_v0_})
            d_3234_v0_ = (d_3304_v54_).isdisjoint((d_3304_v54_) | (d_3304_v54_))
            d_3305_v55_: D3
            d_3305_v55_ = D3_DC6(d_3292_v44_)
            d_3306_v56_: D9
            d_3306_v56_ = D9_DC21(d_3305_v55_, d_3295_v47_, d_3234_v0_)
            rhs486_ = d_3292_v44_
            rhs487_ = (self).f11
            rhs488_ = (default__.fm61(d_3295_v47_, d_3234_v0_, globalState)).isdisjoint(default__.fm61((d_3306_v56_).cf23, d_3234_v0_, globalState))
            lhs312_ = self
            d_3292_v44_ = rhs486_
            lhs312_.f13 = rhs487_
            r2 = rhs488_
        elif True:
            d_3307_v57_: str
            d_3307_v57_ = _dafny.CodePoint('t')
            d_3308_v58_: _dafny.Map
            d_3308_v58_ = _dafny.Map({D27_DC63(_dafny.Map({d_3307_v57_: True})): d_3234_v0_})
            d_3309_v59_: _dafny.Map
            d_3309_v59_ = _dafny.Map({d_3307_v57_: d_3234_v0_})
            d_3310_v60_: D27
            d_3310_v60_ = D27_DC63(d_3309_v59_)
            d_3308_v58_ = (d_3308_v58_).set((default__.fm70(globalState) if d_3234_v0_ else d_3310_v60_), d_3234_v0_)
            d_3311_v61_: _dafny.Map
            d_3311_v61_ = _dafny.Map({d_3234_v0_: d_3234_v0_})
            d_3312_v62_: _dafny.Map
            d_3312_v62_ = _dafny.Map({(self).f9: ((d_3311_v61_)[not(d_3234_v0_)] if (not(d_3234_v0_)) in (d_3311_v61_) else d_3234_v0_)})
            d_3313_v63_: C0
            nw533_ = C0()
            nw533_.ctor__(((d_3312_v62_)[(self).f9] if ((self).f9) in (d_3312_v62_) else d_3234_v0_), not (True) or (d_3234_v0_), (self).f9, (self).f10)
            d_3313_v63_ = nw533_
            d_3314_v64_: _dafny.Seq
            d_3314_v64_ = _dafny.SeqWithoutIsStrInference([d_3234_v0_])
            d_3315_v65_: _dafny.Set
            d_3315_v65_ = _dafny.Set({315, self.f13})
            d_3316_v66_: _dafny.MultiSet
            d_3316_v66_ = _dafny.MultiSet([(d_3314_v64_).set(default__.safeIndex(p0, len(d_3314_v64_)), default__.fm22(d_3234_v0_, len(d_3315_v65_), globalState)), d_3314_v64_])
            d_3317_v67_: _dafny.Map
            d_3317_v67_ = _dafny.Map({(d_3316_v66_) != (d_3316_v66_): ((self).f10) + ((self).f10)})
            d_3317_v67_ = (d_3317_v67_).set(d_3234_v0_, ((self).f10) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0 for d_3318_i7_ in range(default__.abs(414))]) for d_3319_i6_ in range(default__.abs(856))])))
            d_3320_v68_: _dafny.Array
            def lambda152_(d_3321_v64_):
                def lambda153_(d_3322_i8_):
                    return default__.safeModuloInt(d_3322_i8_, (0) - (len(d_3321_v64_)))

                return lambda153_

            init87_ = lambda152_(d_3314_v64_)
            nw534_ = _dafny.Array(None, 14)
            for i0_87_ in range(nw534_.length(0)):
                nw534_[i0_87_] = init87_(i0_87_)
            d_3320_v68_ = nw534_
            d_3323_v69_: D6
            d_3323_v69_ = D6_DC14(d_3320_v68_)
            d_3324_v70_: _dafny.MultiSet
            d_3324_v70_ = _dafny.MultiSet([(d_3323_v69_).cf17])
            d_3325_v71_: D15
            d_3325_v71_ = D15_DC35(d_3324_v70_)
            d_3326_v72_: _dafny.Map
            d_3326_v72_ = _dafny.Map({d_3325_v71_: p0})
            d_3326_v72_ = (d_3326_v72_).set(d_3325_v71_, default__.safeDivisionInt(self.f13, self.f13))
            d_3327_v73_: _dafny.Array
            nw535_ = _dafny.Array(None, 6)
            nw535_[int(0)] = (d_3313_v63_).f26
            nw535_[int(1)] = (d_3313_v63_).f25
            nw535_[int(2)] = d_3234_v0_
            nw535_[int(3)] = default__.fm8(self.f13, (d_3313_v63_).f25, (d_3313_v63_).f25, globalState)
            nw535_[int(4)] = (d_3313_v63_).f25
            nw535_[int(5)] = (d_3313_v63_).f26
            d_3327_v73_ = nw535_
            d_3328_v74_: C9
            nw536_ = C9()
            nw536_.ctor__(d_3327_v73_)
            d_3328_v74_ = nw536_
        d_3329_v76_: _dafny.MultiSet
        def iife245_():
            coll101_ = _dafny.Set()
            compr_101_: int
            for compr_101_ in _dafny.IntegerRange(999, 56):
                d_3330_v75_: int = compr_101_
                if ((999) <= (d_3330_v75_)) and ((d_3330_v75_) < (56)):
                    coll101_ = coll101_.union(_dafny.Set([default__.safeDivisionInt(d_3330_v75_, len(_dafny.Set({-722})))]))
            return _dafny.Set(coll101_)
        d_3329_v76_ = _dafny.MultiSet([len(iife245_()
        ), p0])
        if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0 for d_3331_i9_ in range(default__.abs(660))]))).issubset(d_3329_v76_):
            d_3332_v77_: _dafny.Array
            nw537_ = _dafny.Array(False, 20)
            d_3332_v77_ = nw537_
            d_3333_v78_: _dafny.Array
            nw538_ = _dafny.Array(_dafny.CodePoint('D'), 26)
            d_3333_v78_ = nw538_
            d_3334_v79_: D6
            d_3334_v79_ = D6_DC13(d_3333_v78_)
            d_3335_v80_: _dafny.Map
            d_3335_v80_ = _dafny.Map({self.f13: d_3334_v79_})
            index468_ = default__.safeIndex(243, (d_3332_v77_).length(0))
            (d_3332_v77_)[index468_] = not((len(d_3335_v80_)) <= (self.f13))
            d_3336_v81_: _dafny.Seq
            d_3336_v81_ = _dafny.SeqWithoutIsStrInference([d_3234_v0_])
            index469_ = default__.safeIndex(243, (d_3332_v77_).length(0))
            (d_3332_v77_)[index469_] = (_dafny.MultiSet(d_3336_v81_)).issubset(_dafny.MultiSet([d_3234_v0_]))
            r3 = (107) - (p0)
            d_3337_v82_: _dafny.Seq
            d_3337_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqfxdl"))
            d_3338_v83_: _dafny.Array
            nw539_ = _dafny.Array(None, 8)
            nw539_[int(0)] = d_3337_v82_
            nw539_[int(1)] = (self).f9
            nw539_[int(2)] = (self).f9
            nw539_[int(3)] = (d_3337_v82_ if d_3234_v0_ else (self).f9)
            nw539_[int(4)] = (self).f9
            nw539_[int(5)] = ((self).f9) + ((self).f9)
            nw539_[int(6)] = (self).f9
            nw539_[int(7)] = ((self).f9) + ((self).f9)
            d_3338_v83_ = nw539_
            d_3339_v84_: _dafny.Map
            d_3339_v84_ = _dafny.Map({d_3234_v0_: (self).f11})
            d_3340_v85_: _dafny.Map
            d_3340_v85_ = _dafny.Map({((d_3339_v84_)[(d_3332_v77_)[default__.safeIndex(243, (d_3332_v77_).length(0))]] if ((d_3332_v77_)[default__.safeIndex(243, (d_3332_v77_).length(0))]) in (d_3339_v84_) else p0): (d_3332_v77_)[default__.safeIndex(243, (d_3332_v77_).length(0))]})
            d_3341_v86_: str
            d_3341_v86_ = _dafny.CodePoint('t')
            d_3342_v87_: _dafny.MultiSet
            d_3342_v87_ = _dafny.MultiSet([d_3341_v86_, d_3341_v86_])
            index470_ = default__.safeIndex(51, (d_3338_v83_).length(0))
            (d_3338_v83_)[index470_] = ((self).f9 if ((d_3340_v85_)[((d_3342_v87_)[d_3341_v86_] if (d_3341_v86_) in (d_3342_v87_) else p0)] if (((d_3342_v87_)[d_3341_v86_] if (d_3341_v86_) in (d_3342_v87_) else p0)) in (d_3340_v85_) else (d_3332_v77_)[default__.safeIndex(243, (d_3332_v77_).length(0))]) else d_3337_v82_)
            index471_ = default__.safeIndex(51, (d_3338_v83_).length(0))
            rhs489_ = d_3337_v82_
            rhs490_ = d_3336_v81_
            rhs491_ = d_3338_v83_
            rhs492_ = (d_3337_v82_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccvj")))
            rhs493_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdqwgmr"))
            lhs313_ = d_3338_v83_
            lhs314_ = default__.safeIndex(51, (d_3338_v83_).length(0))
            d_3337_v82_ = rhs489_
            d_3336_v81_ = rhs490_
            d_3338_v83_ = rhs491_
            d_3337_v82_ = rhs492_
            lhs313_[lhs314_] = rhs493_
            d_3343_v88_: _dafny.Seq
            d_3343_v88_ = _dafny.SeqWithoutIsStrInference([p0])
            d_3344_v89_: T1
            nw540_ = C6()
            nw540_.ctor__(len(d_3337_v82_), self.f12, (self).f9, _dafny.SeqWithoutIsStrInference([d_3343_v88_]))
            d_3344_v89_ = nw540_
            rhs494_ = (D22_DC54(p0, p0)).cf80
            rhs495_ = d_3344_v89_
            lhs315_ = globalState
            lhs315_.f3 = rhs494_
            d_3344_v89_ = rhs495_
            d_3345_v90_: _dafny.Map
            d_3345_v90_ = _dafny.Map({default__.fm0((d_3344_v89_).fm5((d_3343_v88_)[default__.safeIndex(self.f13, len(d_3343_v88_))], len(d_3339_v84_), globalState), (d_3344_v89_).f11, globalState): self.f13})
            d_3345_v90_ = (d_3345_v90_).set(default__.safeDivisionInt(900, len(d_3343_v88_)), p0)
        elif True:
            d_3346_v91_: _dafny.Seq
            d_3346_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "giaoojal"))
            d_3347_v92_: _dafny.Map
            d_3347_v92_ = _dafny.Map({default__.fm8(self.f13, d_3234_v0_, False, globalState): d_3346_v91_})
            d_3346_v91_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ijtdy"))) + (((d_3347_v92_)[d_3234_v0_] if (d_3234_v0_) in (d_3347_v92_) else (self).f9))
            (globalState).f3 = (self).f11
            d_3348_v93_: _dafny.Seq
            d_3348_v93_ = _dafny.SeqWithoutIsStrInference([self.f13])
            r0 = default__.safeModuloInt(len((d_3348_v93_) + (default__.fm44(d_3234_v0_, p0, d_3234_v0_, globalState))), -926)
            (globalState).f3 = (0) - (p0)
            d_3349_v94_: _dafny.Array
            nw541_ = _dafny.Array(False, 10)
            d_3349_v94_ = nw541_
            d_3350_v95_: _dafny.Seq
            d_3350_v95_ = _dafny.SeqWithoutIsStrInference([d_3349_v94_, d_3349_v94_])
            if (d_3350_v95_) < (d_3350_v95_):
                (self).f13 = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([d_3234_v0_])), (0) - (p0))
                r2 = not(d_3234_v0_)
                d_3234_v0_ = not(d_3234_v0_)
                d_3351_v96_: D7
                d_3351_v96_ = D7_DC15(d_3329_v76_)
                d_3351_v96_ = d_3351_v96_
                d_3352_v97_: _dafny.Map
                d_3352_v97_ = _dafny.Map({d_3234_v0_: (0) - ((self).f11)})
                d_3353_v98_: _dafny.Set
                d_3353_v98_ = _dafny.Set({d_3352_v97_})
                d_3354_v99_: _dafny.Map
                d_3354_v99_ = _dafny.Map({True: d_3353_v98_})
                d_3354_v99_ = (d_3354_v99_).set(d_3234_v0_, d_3353_v98_)
            elif True:
                (self).f13 = (p0) * (p0)
                r0 = (p0) + (p0)
                d_3355_v100_: _dafny.MultiSet
                d_3355_v100_ = _dafny.MultiSet([(self.f13) > ((0) - (self.f13)), d_3234_v0_])
                d_3355_v100_ = (d_3355_v100_) | (d_3355_v100_)
                d_3356_v101_: _dafny.Map
                d_3356_v101_ = _dafny.Map({492: d_3234_v0_})
                rhs496_ = (D14_DC34(self.f13, default__.fm0(self.f13, p0, globalState), p0, d_3234_v0_)).cf46
                rhs497_ = d_3356_v101_
                rhs498_ = d_3234_v0_
                d_3234_v0_ = rhs496_
                d_3356_v101_ = rhs497_
                d_3234_v0_ = rhs498_
                d_3357_v102_: _dafny.Map
                d_3357_v102_ = _dafny.Map({d_3234_v0_: not(d_3234_v0_)})
                d_3234_v0_ = (((0) - (len(d_3357_v102_))) - ((self).f11)) != (default__.safeDivisionInt(p0, 991))
        d_3358_v103_: D8
        d_3358_v103_ = D8_DC19(default__.fm71(globalState))
        d_3358_v103_ = d_3358_v103_
        d_3359_v104_: D1
        d_3359_v104_ = D1_DC1(d_3234_v0_)
        source52_ = d_3359_v104_
        if source52_.is_DC2:
            if d_3234_v0_:
                d_3360_v105_: _dafny.Seq
                d_3360_v105_ = _dafny.SeqWithoutIsStrInference([(self).f9, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "leuaeks"))])
                d_3361_v106_: _dafny.Seq
                d_3361_v106_ = _dafny.SeqWithoutIsStrInference([len((self).f9), (len(d_3360_v105_)) * ((self).f11)])
                (globalState).f3 = (d_3361_v106_)[default__.safeIndex((self).f11, len(d_3361_v106_))]
                d_3362_v107_: _dafny.Seq
                d_3362_v107_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emiqaayyi"))
                d_3363_v108_: D14
                d_3363_v108_ = D14_DC33((self).f11, d_3362_v107_, d_3234_v0_)
                d_3362_v107_ = (d_3363_v108_).cf41
                d_3364_v109_: _dafny.Map
                d_3364_v109_ = _dafny.Map({(self).f11: self.f13})
                d_3364_v109_ = (d_3364_v109_).set(p0, default__.safeModuloInt(self.f13, p0))
                d_3234_v0_ = d_3234_v0_
                d_3365_v110_: _dafny.Array
                def lambda154_(d_3366_v105_):
                    def lambda155_(d_3367_i10_):
                        return d_3366_v105_

                    return lambda155_

                init88_ = lambda154_(d_3360_v105_)
                nw542_ = _dafny.Array(None, 9)
                for i0_88_ in range(nw542_.length(0)):
                    nw542_[i0_88_] = init88_(i0_88_)
                d_3365_v110_ = nw542_
                index472_ = default__.safeIndex(995, (d_3365_v110_).length(0))
                (d_3365_v110_)[index472_] = d_3360_v105_
                index473_ = default__.safeIndex(995, (d_3365_v110_).length(0))
                (d_3365_v110_)[index473_] = d_3360_v105_
            elif True:
                d_3368_v111_: _dafny.MultiSet
                d_3368_v111_ = _dafny.MultiSet([(self).f9])
                d_3369_v112_: D10
                d_3369_v112_ = D10_DC22(d_3368_v111_)
                d_3370_v113_: _dafny.Map
                d_3370_v113_ = _dafny.Map({d_3369_v112_: d_3234_v0_})
                d_3371_v114_: _dafny.Map
                d_3371_v114_ = _dafny.Map({d_3370_v113_: (self).f9})
                r0 = len((d_3371_v114_) | ((d_3371_v114_ if d_3234_v0_ else d_3371_v114_)))
                d_3372_v115_: _dafny.Seq
                d_3372_v115_ = _dafny.SeqWithoutIsStrInference([-391])
                d_3373_v116_: _dafny.Seq
                d_3373_v116_ = _dafny.SeqWithoutIsStrInference([d_3372_v115_, _dafny.SeqWithoutIsStrInference([p0, p0])])
                d_3373_v116_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_3234_v0_, False, d_3234_v0_, not(d_3234_v0_)])).cardinality for d_3374_i11_ in range(default__.abs(520))]), (d_3372_v115_) + (_dafny.SeqWithoutIsStrInference([self.f13, (self).f11, (self).f11])), (_dafny.SeqWithoutIsStrInference([p0 for d_3375_i12_ in range(default__.abs(574))])) + (d_3372_v115_), d_3372_v115_])
                (self).m4(p0, globalState)
                d_3234_v0_ = d_3234_v0_
                (globalState).f3 = ((self).f11) - (((self).f11 if d_3234_v0_ else len((self).f9)))
            d_3376_v117_: C0
            nw543_ = C0()
            nw543_.ctor__(d_3234_v0_, (not(d_3234_v0_) if d_3234_v0_ else d_3234_v0_), (self).f9, (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0]) for d_3377_i13_ in range(default__.abs(-177))])) + ((self).f10))
            d_3376_v117_ = nw543_
            d_3378_v118_: _dafny.Array
            nw544_ = _dafny.Array(None, 10)
            nw544_[int(0)] = not(False)
            nw544_[int(1)] = (d_3376_v117_).f26
            nw544_[int(2)] = False
            nw544_[int(3)] = d_3234_v0_
            nw544_[int(4)] = (d_3376_v117_).f25
            nw544_[int(5)] = (d_3376_v117_).f26
            nw544_[int(6)] = not((d_3376_v117_).f25)
            nw544_[int(7)] = (d_3376_v117_).f25
            nw544_[int(8)] = default__.fm13((d_3376_v117_).f25, self.f13, (self).f9, globalState)
            nw544_[int(9)] = (d_3376_v117_).f26
            d_3378_v118_ = nw544_
            d_3379_v119_: _dafny.Map
            d_3379_v119_ = _dafny.Map({d_3378_v118_: (d_3376_v117_).f26})
            d_3379_v119_ = d_3379_v119_
            d_3380_v120_: _dafny.Map
            d_3380_v120_ = _dafny.Map({(d_3376_v117_).f26: (d_3376_v117_).f26})
            d_3381_v121_: _dafny.Seq
            d_3381_v121_ = _dafny.SeqWithoutIsStrInference([d_3380_v120_])
            d_3382_v122_: _dafny.Set
            d_3382_v122_ = _dafny.Set({(d_3376_v117_).f25, d_3234_v0_, (d_3376_v117_).f25})
            d_3383_v123_: D20
            d_3383_v123_ = D20_DC47((d_3376_v117_).f25, len((d_3381_v121_)[default__.safeIndex(len(d_3382_v122_), len(d_3381_v121_))]), (d_3376_v117_).f26)
            source53_ = d_3383_v123_
            if source53_.is_DC47:
                d_3384___mcc_h4_ = source53_.cf65
                d_3385___mcc_h5_ = source53_.cf66
                d_3386___mcc_h6_ = source53_.cf67
                d_3387_cf67_ = d_3386___mcc_h6_
                d_3388_cf66_ = d_3385___mcc_h5_
                d_3389_cf65_ = d_3384___mcc_h4_
                r3 = (self).f11
                d_3390_v124_: _dafny.Seq
                d_3390_v124_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                d_3391_v125_: _dafny.Set
                d_3391_v125_ = _dafny.Set({d_3388_cf66_})
                d_3392_v126_: D3
                d_3392_v126_ = D3_DC6(d_3391_v125_)
                d_3393_v127_: _dafny.Map
                d_3393_v127_ = _dafny.Map({d_3387_cf67_: d_3392_v126_})
                d_3394_v128_: _dafny.Seq
                d_3394_v128_ = _dafny.SeqWithoutIsStrInference([(d_3376_v117_).f25, d_3387_cf67_, (d_3376_v117_).f26])
                d_3395_v129_: _dafny.Seq
                d_3395_v129_ = _dafny.SeqWithoutIsStrInference([(d_3394_v128_)[default__.safeIndex(366, len(d_3394_v128_))]])
                d_3396_v130_: str
                d_3396_v130_ = _dafny.CodePoint('q')
                d_3397_v131_: _dafny.Array
                nw545_ = _dafny.Array(None, 16)
                nw545_[int(0)] = (0) - (self.f13)
                nw545_[int(1)] = p0
                nw545_[int(2)] = len(d_3394_v128_)
                nw545_[int(3)] = (self).f11
                nw545_[int(4)] = 785
                nw545_[int(5)] = p0
                nw545_[int(6)] = p0
                nw545_[int(7)] = (self).f11
                nw545_[int(8)] = self.f13
                nw545_[int(9)] = self.f13
                nw545_[int(10)] = d_3388_cf66_
                nw545_[int(11)] = (0) - (p0)
                nw545_[int(12)] = len(_dafny.Map({d_3396_v130_: self.f13}))
                nw545_[int(13)] = p0
                nw545_[int(14)] = d_3388_cf66_
                nw545_[int(15)] = (self).f11
                d_3397_v131_ = nw545_
                d_3398_v132_: _dafny.Map
                d_3398_v132_ = _dafny.Map({d_3397_v131_: d_3390_v124_})
                d_3399_v133_: _dafny.Array
                nw546_ = _dafny.Array(None, 18)
                nw546_[int(0)] = (d_3390_v124_ if (d_3376_v117_).f25 else d_3390_v124_)
                nw546_[int(1)] = (d_3390_v124_) + (_dafny.SeqWithoutIsStrInference([len((self).f9) for d_3400_i14_ in range(default__.abs(-136))]))
                nw546_[int(2)] = (d_3390_v124_) + (_dafny.SeqWithoutIsStrInference([len(d_3393_v127_), (self).f11]))
                nw546_[int(3)] = ((d_3398_v132_)[d_3397_v131_] if (d_3397_v131_) in (d_3398_v132_) else d_3390_v124_)
                nw546_[int(4)] = (d_3390_v124_) + (((self).f10)[default__.safeIndex(len(d_3390_v124_), len((self).f10))])
                nw546_[int(5)] = d_3390_v124_
                nw546_[int(6)] = d_3390_v124_
                nw546_[int(7)] = d_3390_v124_
                nw546_[int(8)] = d_3390_v124_
                nw546_[int(9)] = (d_3390_v124_) + (d_3390_v124_)
                nw546_[int(10)] = _dafny.SeqWithoutIsStrInference([p0, p0])
                nw546_[int(11)] = d_3390_v124_
                nw546_[int(12)] = d_3390_v124_
                nw546_[int(13)] = d_3390_v124_
                nw546_[int(14)] = d_3390_v124_
                nw546_[int(15)] = _dafny.SeqWithoutIsStrInference([p0, p0])
                nw546_[int(16)] = d_3390_v124_
                nw546_[int(17)] = (d_3390_v124_) + (_dafny.SeqWithoutIsStrInference([d_3388_cf66_, p0]))
                d_3399_v133_ = nw546_
                index474_ = default__.safeIndex(865, (d_3399_v133_).length(0))
                (d_3399_v133_)[index474_] = (d_3390_v124_) + (d_3390_v124_)
                index475_ = default__.safeIndex(865, (d_3399_v133_).length(0))
                (d_3399_v133_)[index475_] = ((d_3390_v124_).set(default__.safeIndex((self).f11, len(d_3390_v124_)), (self).f11)) + (d_3390_v124_)
                d_3234_v0_ = d_3234_v0_
                d_3401_v134_: _dafny.Set
                d_3401_v134_ = _dafny.Set({d_3397_v131_})
                d_3402_v135_: D22
                d_3402_v135_ = D22_DC52((d_3401_v134_).intersection(_dafny.Set({d_3397_v131_})))
                d_3402_v135_ = d_3402_v135_
            elif source53_.is_DC48:
                d_3403___mcc_h7_ = source53_.cf68
                d_3404___mcc_h8_ = source53_.cf69
                d_3405___mcc_h9_ = source53_.cf70
                d_3406_cf70_ = d_3405___mcc_h9_
                d_3407_cf69_ = d_3404___mcc_h8_
                d_3408_cf68_ = d_3403___mcc_h7_
                d_3409_v136_: D2
                d_3409_v136_ = D2_DC5(d_3408_cf68_, d_3378_v118_, -576, self.f13)
                d_3409_v136_ = d_3409_v136_
                d_3410_v137_: _dafny.Array
                def lambda156_(d_3411_v76_):
                    def lambda157_(d_3412_i15_):
                        return d_3411_v76_

                    return lambda157_

                init89_ = lambda156_(d_3329_v76_)
                nw547_ = _dafny.Array(None, 13)
                for i0_89_ in range(nw547_.length(0)):
                    nw547_[i0_89_] = init89_(i0_89_)
                d_3410_v137_ = nw547_
                d_3413_v138_: _dafny.Seq
                d_3413_v138_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11])
                index476_ = default__.safeIndex(974, (d_3410_v137_).length(0))
                (d_3410_v137_)[index476_] = _dafny.MultiSet(d_3413_v138_)
                d_3414_v139_: D17
                d_3414_v139_ = D17_DC40((self).f9, (d_3376_v117_).f25, (0) - (self.f13))
                d_3415_v140_: D14
                d_3415_v140_ = D14_DC34(p0, len(d_3413_v138_), (self).f11, d_3234_v0_)
                d_3416_v141_: _dafny.Map
                d_3416_v141_ = _dafny.Map({d_3414_v139_: (d_3415_v140_).cf44})
                index477_ = default__.safeIndex(974, (d_3410_v137_).length(0))
                (d_3410_v137_)[index477_] = _dafny.MultiSet([len(d_3416_v141_)])
                d_3417_v142_: _dafny.Map
                d_3417_v142_ = _dafny.Map({default__.fm72(globalState): self.f13})
                d_3418_v143_: D1
                d_3418_v143_ = D1_DC2()
                d_3419_v144_: D4
                d_3419_v144_ = D4_DC9(d_3418_v143_)
                d_3417_v142_ = (d_3417_v142_).set(_dafny.SeqWithoutIsStrInference([d_3419_v144_]), p0)
                d_3420_v145_: _dafny.Seq
                d_3420_v145_ = _dafny.SeqWithoutIsStrInference([d_3234_v0_])
                r3 = (self.f13 if (d_3420_v145_)[default__.safeIndex((self).f11, len(d_3420_v145_))] else (self).f11)
            elif True:
                d_3421___mcc_h10_ = source53_.cf64
                d_3422_cf64_ = d_3421___mcc_h10_
                d_3423_v146_: _dafny.Seq
                d_3423_v146_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "huqeb"))
                d_3423_v146_ = (self).f9
                d_3424_v147_: T1
                nw548_ = C1()
                nw548_.ctor__(default__.safeModuloInt(p0, (self).f11), default__.safeDivisionInt(self.f13, 82), p0, self.f12, d_3423_v146_, ((self).f10) + ((self).f10))
                d_3424_v147_ = nw548_
                d_3424_v147_ = d_3424_v147_
                d_3425_v148_: _dafny.Map
                d_3425_v148_ = _dafny.Map({True: 746})
                d_3426_v149_: _dafny.Seq
                d_3426_v149_ = _dafny.SeqWithoutIsStrInference([len(d_3425_v148_)])
                (self).f13 = (d_3426_v149_)[default__.safeIndex((d_3426_v149_)[default__.safeIndex((self).f11, len(d_3426_v149_))], len(d_3426_v149_))]
                d_3427_v150_: _dafny.Set
                d_3427_v150_ = _dafny.Set({self.f13, (d_3424_v147_).f11, p0, (d_3424_v147_).f11, p0})
                d_3428_v151_: bool
                d_3429_v152_: bool
                d_3430_v153_: _dafny.Array
                d_3431_v154_: int
                out87_: bool
                out88_: bool
                out89_: _dafny.Array
                out90_: int
                out87_, out88_, out89_, out90_ = (d_3424_v147_).m1(self.f13, (d_3424_v147_).f11, len(d_3427_v150_), globalState)
                d_3428_v151_ = out87_
                d_3429_v152_ = out88_
                d_3430_v153_ = out89_
                d_3431_v154_ = out90_
        elif source52_.is_DC1:
            d_3432___mcc_h2_ = source52_.cf1
            d_3433_cf1_ = d_3432___mcc_h2_
            d_3434_v155_: _dafny.Array
            def lambda158_(d_3435_p0_):
                def lambda159_(d_3436_i16_):
                    return (d_3436_i16_) + (d_3435_p0_)

                return lambda159_

            init90_ = lambda158_(p0)
            nw549_ = _dafny.Array(None, 3)
            for i0_90_ in range(nw549_.length(0)):
                nw549_[i0_90_] = init90_(i0_90_)
            d_3434_v155_ = nw549_
            d_3437_v156_: _dafny.Map
            d_3437_v156_ = _dafny.Map({(self).f11: d_3434_v155_})
            d_3437_v156_ = (d_3437_v156_).set(len(((self).f9) + ((self).f9)), d_3434_v155_)
            d_3438_v157_: str
            d_3438_v157_ = _dafny.CodePoint('c')
            d_3439_v158_: _dafny.Seq
            d_3439_v158_ = _dafny.SeqWithoutIsStrInference([(self).f9])
            d_3440_v160_: _dafny.Set
            def iife246_():
                coll102_ = _dafny.Set()
                compr_102_: _dafny.Seq
                for compr_102_ in (d_3439_v158_).Elements:
                    d_3441_v159_: _dafny.Seq = compr_102_
                    if (d_3441_v159_) in (d_3439_v158_):
                        coll102_ = coll102_.union(_dafny.Set([d_3441_v159_]))
                return _dafny.Set(coll102_)
            d_3440_v160_ = iife246_()
            
            def iife247_():
                coll103_ = _dafny.Set()
                compr_103_: _dafny.Seq
                for compr_103_ in (d_3439_v158_).Elements:
                    d_3442_v161_: _dafny.Seq = compr_103_
                    if (d_3442_v161_) in (d_3439_v158_):
                        coll103_ = coll103_.union(_dafny.Set([d_3442_v161_]))
                return _dafny.Set(coll103_)
            d_3438_v157_ = (_dafny.CodePoint('p') if (iife247_()
            ).issubset((d_3440_v160_)) else d_3438_v157_)
            d_3443_v162_: _dafny.Set
            d_3443_v162_ = _dafny.Set({(self).f11})
            if ((_dafny.Set({p0})).isdisjoint(d_3443_v162_)) or (d_3234_v0_):
                d_3444_v163_: _dafny.MultiSet
                d_3444_v163_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxctkfr"))])
                d_3445_v164_: D10
                d_3445_v164_ = D10_DC22((d_3444_v163_).set((self).f9, default__.abs(self.f13)))
                d_3446_v165_: _dafny.Seq
                d_3446_v165_ = _dafny.SeqWithoutIsStrInference([d_3234_v0_])
                d_3447_v166_: _dafny.Array
                nw550_ = _dafny.Array(None, 3)
                nw550_[int(0)] = d_3445_v164_
                nw550_[int(1)] = d_3445_v164_
                nw550_[int(2)] = default__.fm73((d_3446_v165_)[default__.safeIndex(len(d_3446_v165_), len(d_3446_v165_))], globalState)
                d_3447_v166_ = nw550_
                d_3448_v167_: _dafny.Set
                d_3448_v167_ = _dafny.Set({not(default__.fm8(self.f13, True, (d_3446_v165_)[default__.safeIndex((self).f11, len(d_3446_v165_))], globalState))})
                d_3449_v168_: _dafny.Array
                nw551_ = _dafny.Array(None, 20)
                nw551_[int(0)] = (196) not in (default__.fm53(globalState))
                nw551_[int(1)] = (d_3433_cf1_ if False else d_3234_v0_)
                nw551_[int(2)] = d_3433_cf1_
                nw551_[int(3)] = d_3433_cf1_
                nw551_[int(4)] = d_3234_v0_
                nw551_[int(5)] = d_3433_cf1_
                nw551_[int(6)] = (d_3438_v157_) in ((self).f9)
                nw551_[int(7)] = d_3433_cf1_
                nw551_[int(8)] = d_3234_v0_
                nw551_[int(9)] = True
                nw551_[int(10)] = (d_3234_v0_) not in (d_3446_v165_)
                nw551_[int(11)] = default__.fm38(default__.fm0(self.f13, p0, globalState), globalState)
                nw551_[int(12)] = d_3433_cf1_
                nw551_[int(13)] = (d_3448_v167_).issubset(d_3448_v167_)
                nw551_[int(14)] = d_3433_cf1_
                nw551_[int(15)] = d_3433_cf1_
                nw551_[int(16)] = d_3433_cf1_
                nw551_[int(17)] = d_3433_cf1_
                nw551_[int(18)] = d_3234_v0_
                nw551_[int(19)] = d_3234_v0_
                d_3449_v168_ = nw551_
                d_3450_v169_: _dafny.Seq
                d_3450_v169_ = _dafny.SeqWithoutIsStrInference([self.f13, p0, len(d_3446_v165_), (_dafny.MultiSet([d_3433_cf1_])).cardinality, (self).f11])
                index478_ = default__.safeIndex(990, (d_3449_v168_).length(0))
                (d_3449_v168_)[index478_] = (d_3450_v169_) < (d_3450_v169_)
                index479_ = default__.safeIndex(990, (d_3449_v168_).length(0))
                rhs499_ = (self).f11
                rhs500_ = self.f13
                rhs501_ = d_3447_v166_
                rhs502_ = self.f13
                rhs503_ = (False) == (d_3234_v0_)
                lhs316_ = globalState
                lhs317_ = globalState
                lhs318_ = d_3449_v168_
                lhs319_ = default__.safeIndex(990, (d_3449_v168_).length(0))
                lhs316_.f3 = rhs499_
                lhs317_.f3 = rhs500_
                d_3447_v166_ = rhs501_
                r0 = rhs502_
                lhs318_[lhs319_] = rhs503_
                d_3451_v170_: _dafny.Map
                d_3451_v170_ = _dafny.Map({d_3446_v165_: d_3434_v155_})
                d_3452_v171_: D6
                d_3452_v171_ = D6_DC14(d_3434_v155_)
                pat_let_tv67_ = d_3434_v155_
                def iife248_(_pat_let72_0):
                    def iife249_(d_3453_dt__update__tmp_h1_):
                        def iife250_(_pat_let73_0):
                            def iife251_(d_3454_dt__update_hcf17_h0_):
                                return D6_DC14(d_3454_dt__update_hcf17_h0_)
                            return iife251_(_pat_let73_0)
                        return iife250_(pat_let_tv67_)
                    return iife249_(_pat_let72_0)
                d_3451_v170_ = (d_3451_v170_).set(d_3446_v165_, (iife248_(d_3452_v171_)).cf17)
                index480_ = default__.safeIndex(990, (d_3449_v168_).length(0))
                (d_3449_v168_)[index480_] = not(d_3234_v0_)
                d_3455_v172_: _dafny.Array
                def lambda160_(d_3456_v0_):
                    def lambda161_(d_3457_i17_):
                        return (D17_DC39(_dafny.Map({d_3456_v0_: (self).f11}))).cf53

                    return lambda161_

                init91_ = lambda160_(d_3234_v0_)
                nw552_ = _dafny.Array(None, 28)
                for i0_91_ in range(nw552_.length(0)):
                    nw552_[i0_91_] = init91_(i0_91_)
                d_3455_v172_ = nw552_
                d_3458_v173_: _dafny.Map
                d_3458_v173_ = _dafny.Map({d_3234_v0_: p0})
                index481_ = default__.safeIndex(841, (d_3455_v172_).length(0))
                (d_3455_v172_)[index481_] = default__.fm42(d_3458_v173_, d_3234_v0_, p0, (d_3449_v168_)[default__.safeIndex(990, (d_3449_v168_).length(0))], globalState)
                index482_ = default__.safeIndex(841, (d_3455_v172_).length(0))
                (d_3455_v172_)[index482_] = (d_3458_v173_) | (d_3458_v173_)
                index483_ = default__.safeIndex(990, (d_3449_v168_).length(0))
                (d_3449_v168_)[index483_] = d_3234_v0_
            elif True:
                d_3459_v174_: _dafny.Array
                nw553_ = _dafny.Array(None, 27)
                d_3459_v174_ = nw553_
                d_3234_v0_ = (_dafny.MultiSet([d_3459_v174_])).ispropersubset(_dafny.MultiSet([d_3459_v174_]))
                d_3460_v175_: _dafny.Array
                def lambda162_(d_3461_v157_):
                    def lambda163_(d_3462_i18_):
                        return d_3461_v157_

                    return lambda163_

                init92_ = lambda162_(d_3438_v157_)
                nw554_ = _dafny.Array(None, 8)
                for i0_92_ in range(nw554_.length(0)):
                    nw554_[i0_92_] = init92_(i0_92_)
                d_3460_v175_ = nw554_
                d_3460_v175_ = (d_3460_v175_ if d_3433_cf1_ else d_3460_v175_)
                d_3433_cf1_ = d_3234_v0_
                d_3463_v176_: _dafny.Map
                d_3463_v176_ = _dafny.Map({p0: p0})
                d_3464_v177_: _dafny.Map
                d_3464_v177_ = _dafny.Map({(d_3463_v176_) | (d_3463_v176_): d_3433_cf1_})
                d_3464_v177_ = d_3464_v177_
                d_3465_v178_: _dafny.Map
                d_3465_v178_ = _dafny.Map({(0) - (((self).f11) * (self.f13)): d_3433_cf1_})
                d_3465_v178_ = (d_3465_v178_).set((self).fm5(p0, (self).f11, globalState), d_3234_v0_)
            d_3466_v179_: C0
            nw555_ = C0()
            nw555_.ctor__(d_3433_cf1_, d_3234_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iplbj")), (self).f10)
            d_3466_v179_ = nw555_
            d_3467_v180_: D14
            d_3467_v180_ = D14_DC32(d_3466_v179_)
            pat_let_tv68_ = d_3466_v179_
            d_3468_v181_: _dafny.MultiSet
            def iife252_(_pat_let74_0):
                def iife253_(d_3469_dt__update__tmp_h2_):
                    def iife254_(_pat_let75_0):
                        def iife255_(d_3470_dt__update_hcf39_h0_):
                            return D14_DC32(d_3470_dt__update_hcf39_h0_)
                        return iife255_(_pat_let75_0)
                    return iife254_(pat_let_tv68_)
                return iife253_(_pat_let74_0)
            d_3468_v181_ = _dafny.MultiSet([iife252_(D14_DC32(d_3466_v179_)), (d_3467_v180_ if (d_3466_v179_).f26 else d_3467_v180_)])
            d_3468_v181_ = (d_3468_v181_) | (d_3468_v181_)
        elif True:
            d_3471___mcc_h3_ = source52_.cf2
            d_3472_cf2_ = d_3471___mcc_h3_
            d_3473_v182_: D3
            d_3473_v182_ = D3_DC7(not(d_3234_v0_))
            d_3474_v183_: _dafny.Array
            nw556_ = _dafny.Array(None, 29)
            nw556_[int(0)] = d_3234_v0_
            nw556_[int(1)] = d_3234_v0_
            nw556_[int(2)] = d_3234_v0_
            nw556_[int(3)] = d_3234_v0_
            nw556_[int(4)] = d_3234_v0_
            nw556_[int(5)] = d_3234_v0_
            nw556_[int(6)] = d_3234_v0_
            nw556_[int(7)] = d_3234_v0_
            nw556_[int(8)] = d_3234_v0_
            nw556_[int(9)] = d_3234_v0_
            nw556_[int(10)] = False
            nw556_[int(11)] = d_3234_v0_
            nw556_[int(12)] = d_3234_v0_
            nw556_[int(13)] = d_3234_v0_
            nw556_[int(14)] = d_3234_v0_
            nw556_[int(15)] = d_3234_v0_
            nw556_[int(16)] = d_3234_v0_
            nw556_[int(17)] = d_3234_v0_
            nw556_[int(18)] = d_3234_v0_
            nw556_[int(19)] = True
            nw556_[int(20)] = d_3234_v0_
            nw556_[int(21)] = d_3234_v0_
            nw556_[int(22)] = d_3234_v0_
            nw556_[int(23)] = d_3234_v0_
            nw556_[int(24)] = d_3234_v0_
            nw556_[int(25)] = d_3234_v0_
            nw556_[int(26)] = default__.fm22(d_3234_v0_, p0, globalState)
            nw556_[int(27)] = d_3234_v0_
            nw556_[int(28)] = True
            d_3474_v183_ = nw556_
            d_3475_v184_: D2
            d_3475_v184_ = D2_DC5(d_3234_v0_, d_3474_v183_, default__.fm0(self.f13, p0, globalState), (self).f11)
            d_3473_v182_ = (d_3473_v182_ if (d_3475_v184_).cf4 else d_3473_v182_)
            (self).f13 = (default__.safeDivisionInt(p0, p0)) - (915)
            d_3476_v185_: _dafny.Set
            d_3476_v185_ = _dafny.Set({(self).f9})
            d_3234_v0_ = (d_3476_v185_).isdisjoint(d_3476_v185_)
            r2 = d_3234_v0_
        d_3477_v186_: C1
        nw557_ = C1()
        nw557_.ctor__(p0, default__.fm0(-834, self.f13, globalState), self.f13, self.f12, (self).f9, (self).f10)
        d_3477_v186_ = nw557_
        d_3478_v187_: _dafny.Map
        d_3478_v187_ = _dafny.Map({d_3477_v186_: d_3234_v0_})
        rhs504_ = d_3234_v0_
        rhs505_ = (d_3234_v0_ if not(d_3234_v0_) else (d_3477_v186_) not in (d_3478_v187_))
        d_3234_v0_ = rhs504_
        r2 = rhs505_
        r0 = (self).f11
        d_3479_v188_: _dafny.Map
        d_3479_v188_ = _dafny.Map({d_3234_v0_: (self).f9})
        r1 = d_3479_v188_
        r2 = d_3234_v0_
        r3 = (0) - (p0)
        return r0, r1, r2, r3

    def m3(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_3480_v0_: _dafny.Array
        nw558_ = _dafny.Array(_dafny.CodePoint('D'), 21)
        d_3480_v0_ = nw558_
        guard_loop_11_: int
        for guard_loop_11_ in _dafny.IntegerRange(0, (d_3480_v0_).length(0)):
            d_3481_i0_: int = guard_loop_11_
            if (True) and (((0) <= (d_3481_i0_)) and ((d_3481_i0_) < ((d_3480_v0_).length(0)))):
                (d_3480_v0_)[(d_3481_i0_)] = _dafny.CodePoint('c')
        d_3482_v1_: bool
        d_3482_v1_ = False
        d_3483_v2_: _dafny.Map
        d_3483_v2_ = _dafny.Map({not(d_3482_v1_): d_3482_v1_})
        d_3484_v3_: _dafny.Map
        d_3484_v3_ = _dafny.Map({self.f13: d_3483_v2_})
        d_3482_v1_ = (d_3482_v1_) or (not((((d_3484_v3_)[len(p1)] if (len(p1)) in (d_3484_v3_) else _dafny.Map({d_3482_v1_: True}))) == ((d_3483_v2_).set(d_3482_v1_, d_3482_v1_))))
        d_3485_v4_: _dafny.Seq
        d_3485_v4_ = _dafny.SeqWithoutIsStrInference([(self).f11])
        d_3486_v5_: C7
        nw559_ = C7()
        nw559_.ctor__((self).f9, (((self).f10) + (default__.fm41(d_3482_v1_, d_3482_v1_, d_3482_v1_, globalState))).set(default__.safeIndex(len(p1), len(((self).f10) + (default__.fm41(d_3482_v1_, d_3482_v1_, d_3482_v1_, globalState)))), d_3485_v4_), (p0 if not(d_3482_v1_) else (self).f11), (self.f12))
        d_3486_v5_ = nw559_
        (globalState).f3 = p2
        d_3487_v6_: _dafny.Seq
        d_3487_v6_ = _dafny.SeqWithoutIsStrInference([d_3482_v1_])
        d_3488_v7_: C3
        nw560_ = C3()
        nw560_.ctor__(d_3482_v1_, len(d_3487_v6_), p0, self.f12, (self).f9, (self).f10)
        d_3488_v7_ = nw560_
        d_3489_v8_: str
        d_3489_v8_ = _dafny.CodePoint('f')
        pat_let_tv69_ = p0
        def lambda164_(source54_):
            d_3490___mcc_h0_ = source54_
            d_3491_cf15_ = d_3490___mcc_h0_
            return (len((self).f9)) * (pat_let_tv69_)

        (globalState).f3 = lambda164_(d_3489_v8_)
        d_3492_v9_: _dafny.Array
        nw561_ = _dafny.Array(False, 5)
        d_3492_v9_ = nw561_
        d_3493_v10_: _dafny.Array
        nw562_ = _dafny.Array(None, 12)
        nw562_[int(0)] = d_3492_v9_
        nw562_[int(1)] = d_3492_v9_
        nw562_[int(2)] = d_3492_v9_
        nw562_[int(3)] = d_3492_v9_
        nw562_[int(4)] = d_3492_v9_
        nw562_[int(5)] = d_3492_v9_
        nw562_[int(6)] = d_3492_v9_
        nw562_[int(7)] = d_3492_v9_
        nw562_[int(8)] = d_3492_v9_
        nw562_[int(9)] = d_3492_v9_
        nw562_[int(10)] = d_3492_v9_
        nw562_[int(11)] = d_3492_v9_
        d_3493_v10_ = nw562_
        d_3494_v11_: _dafny.Seq
        d_3494_v11_ = _dafny.SeqWithoutIsStrInference([d_3493_v10_, d_3493_v10_, d_3493_v10_, d_3493_v10_])
        r0 = (d_3494_v11_)[default__.safeIndex(p2, len(d_3494_v11_))]
        nw563_ = _dafny.Array(None, 5)
        nw563_[int(0)] = p2
        nw563_[int(1)] = p0
        nw563_[int(2)] = (p2) - (p2)
        nw563_[int(3)] = (p0) + (-841)
        nw563_[int(4)] = p2
        r1 = nw563_
        return r0, r1

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_3495_v0_: C7
        nw564_ = C7()
        nw564_.ctor__((self).f9, (self).f10, p1, self.f12)
        d_3495_v0_ = nw564_
        d_3496_v1_: _dafny.Seq
        d_3496_v1_ = _dafny.SeqWithoutIsStrInference([d_3495_v0_])
        d_3497_i0_: int
        d_3497_i0_ = 0
        with _dafny.label("15"):
            while not(((d_3496_v1_)[default__.safeIndex(self.f13, len(d_3496_v1_))]) not in (d_3496_v1_)):
                with _dafny.c_label("15"):
                    if (d_3497_i0_) >= (100):
                        raise _dafny.Break("15")
                    d_3497_i0_ = (d_3497_i0_) + (1)
                    d_3498_v2_: bool
                    d_3498_v2_ = False
                    d_3499_v3_: D20
                    d_3499_v3_ = D20_DC47(d_3498_v2_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_3500_i1_ in range(default__.abs(-192))])), d_3498_v2_)
                    d_3501_v4_: _dafny.Seq
                    d_3501_v4_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                    d_3502_v5_: _dafny.Map
                    d_3502_v5_ = _dafny.Map({d_3499_v3_: (_dafny.MultiSet(d_3501_v4_)).cardinality})
                    r3 = ((d_3502_v5_)[d_3499_v3_] if (d_3499_v3_) in (d_3502_v5_) else (0) - (p1))
                    d_3503_v6_: _dafny.Array
                    nw565_ = _dafny.Array(int(0), 14)
                    d_3503_v6_ = nw565_
                    r2 = d_3503_v6_
                    d_3504_v7_: C6
                    nw566_ = C6()
                    nw566_.ctor__((self).f11, self.f12, (self).f9, ((self).f10) + (_dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([101 for d_3505_i2_ in range(default__.abs(799))])).set(default__.safeIndex((self).f11, len(_dafny.SeqWithoutIsStrInference([101 for d_3506_i2_ in range(default__.abs(799))]))), (self).f11), d_3501_v4_, d_3501_v4_])))
                    d_3504_v7_ = nw566_
                    d_3507_v8_: str
                    d_3507_v8_ = _dafny.CodePoint('v')
                    d_3508_v9_: str
                    d_3508_v9_ = d_3507_v8_
                    d_3509_v10_: _dafny.MultiSet
                    d_3509_v10_ = _dafny.MultiSet([d_3507_v8_, d_3508_v9_])
                    d_3510_v11_: _dafny.Seq
                    d_3510_v11_ = _dafny.SeqWithoutIsStrInference([d_3508_v9_, d_3508_v9_, d_3508_v9_, d_3507_v8_, d_3508_v9_])
                    d_3509_v10_ = _dafny.MultiSet((d_3510_v11_) + (d_3510_v11_))
                    pass
            pass
        d_3511_v12_: bool
        d_3511_v12_ = False
        d_3512_v13_: _dafny.Map
        d_3512_v13_ = _dafny.Map({default__.safeDivisionInt((self).f11, p2): d_3511_v12_})
        d_3512_v13_ = (d_3512_v13_).set(p0, d_3511_v12_)
        d_3513_v14_: str
        d_3513_v14_ = _dafny.CodePoint('m')
        d_3514_v15_: _dafny.Seq
        d_3514_v15_ = _dafny.SeqWithoutIsStrInference([(0) - ((0) - (p1))])
        d_3515_v16_: T0
        nw567_ = C7()
        nw567_.ctor__(((self).f9).set(default__.safeIndex(-631, len((self).f9)), d_3513_v14_), (self).f10, (d_3514_v15_)[default__.safeIndex(749, len(d_3514_v15_))], self.f12)
        d_3515_v16_ = nw567_
        d_3516_v17_: _dafny.Seq
        d_3516_v17_ = _dafny.SeqWithoutIsStrInference([d_3515_v16_, d_3515_v16_])
        d_3517_v18_: _dafny.Seq
        d_3517_v18_ = _dafny.SeqWithoutIsStrInference([d_3515_v16_, d_3515_v16_, (d_3516_v17_)[default__.safeIndex(p0, len(d_3516_v17_))], (d_3515_v16_ if d_3511_v12_ else d_3515_v16_)])
        d_3517_v18_ = (d_3517_v18_ if d_3511_v12_ else (d_3517_v18_) + (_dafny.SeqWithoutIsStrInference([d_3515_v16_])))
        d_3518_v19_: D15
        d_3518_v19_ = D15_DC36(-535)
        d_3518_v19_ = d_3518_v19_
        hi13_ = p1
        for d_3519_i3_ in range(self.f13, hi13_):
            d_3520_v20_: D17
            d_3520_v20_ = D17_DC40((d_3515_v16_).f9, d_3511_v12_, len((d_3515_v16_).f9))
            d_3521_v21_: _dafny.Set
            d_3521_v21_ = _dafny.Set({True, d_3511_v12_, (d_3520_v20_).cf55})
            d_3521_v21_ = d_3521_v21_
            (self).f13 = self.f13
            d_3522_v22_: _dafny.Map
            d_3522_v22_ = _dafny.Map({d_3511_v12_: d_3511_v12_})
            d_3523_v23_: _dafny.Set
            d_3523_v23_ = _dafny.Set({d_3522_v22_, d_3522_v22_})
            d_3524_v24_: _dafny.Array
            nw568_ = _dafny.Array(_dafny.CodePoint('D'), 24)
            d_3524_v24_ = nw568_
            d_3525_v25_: D20
            d_3525_v25_ = D20_DC48(d_3511_v12_, d_3523_v23_, d_3524_v24_)
            d_3526_v26_: _dafny.Array
            nw569_ = _dafny.Array(None, 28)
            nw569_[int(0)] = d_3511_v12_
            nw569_[int(1)] = d_3511_v12_
            nw569_[int(2)] = d_3511_v12_
            nw569_[int(3)] = not(d_3511_v12_)
            nw569_[int(4)] = d_3511_v12_
            nw569_[int(5)] = d_3511_v12_
            nw569_[int(6)] = d_3511_v12_
            nw569_[int(7)] = d_3511_v12_
            nw569_[int(8)] = d_3511_v12_
            nw569_[int(9)] = d_3511_v12_
            nw569_[int(10)] = d_3511_v12_
            nw569_[int(11)] = True
            nw569_[int(12)] = d_3511_v12_
            nw569_[int(13)] = d_3511_v12_
            nw569_[int(14)] = d_3511_v12_
            nw569_[int(15)] = d_3511_v12_
            nw569_[int(16)] = d_3511_v12_
            nw569_[int(17)] = d_3511_v12_
            nw569_[int(18)] = (d_3525_v25_).cf68
            nw569_[int(19)] = not(d_3511_v12_)
            nw569_[int(20)] = d_3511_v12_
            nw569_[int(21)] = d_3511_v12_
            nw569_[int(22)] = d_3511_v12_
            nw569_[int(23)] = d_3511_v12_
            nw569_[int(24)] = not(d_3511_v12_)
            nw569_[int(25)] = d_3511_v12_
            nw569_[int(26)] = True
            nw569_[int(27)] = d_3511_v12_
            d_3526_v26_ = nw569_
            d_3527_v27_: _dafny.Map
            d_3527_v27_ = _dafny.Map({d_3526_v26_: (d_3515_v16_).f9})
            d_3527_v27_ = (d_3527_v27_).set(d_3526_v26_, ((d_3515_v16_).f9) + ((d_3515_v16_).f9))
            index484_ = default__.safeIndex(978, (d_3526_v26_).length(0))
            (d_3526_v26_)[index484_] = d_3511_v12_
            index485_ = default__.safeIndex(978, (d_3526_v26_).length(0))
            (d_3526_v26_)[index485_] = d_3511_v12_
        d_3528_v28_: C8
        nw570_ = C8()
        nw570_.ctor__(len((self).f9), d_3511_v12_)
        d_3528_v28_ = nw570_
        r0 = d_3528_v28_.f16
        d_3529_v29_: D20
        d_3529_v29_ = D20_DC47(d_3511_v12_, (d_3528_v28_).f15, d_3528_v28_.f16)
        d_3530_v30_: _dafny.Array
        nw571_ = _dafny.Array(_dafny.Seq({}), 25)
        d_3530_v30_ = nw571_
        d_3531_v31_: D22
        d_3531_v31_ = D22_DC53(d_3530_v30_, (self).fm5(p1, p0, globalState))
        d_3532_v32_: _dafny.Array
        nw572_ = _dafny.Array(None, 27)
        nw572_[int(0)] = d_3528_v28_.f16
        nw572_[int(1)] = d_3511_v12_
        nw572_[int(2)] = d_3511_v12_
        nw572_[int(3)] = True
        nw572_[int(4)] = d_3511_v12_
        nw572_[int(5)] = not(d_3511_v12_)
        nw572_[int(6)] = True
        nw572_[int(7)] = not(d_3511_v12_)
        nw572_[int(8)] = (d_3529_v29_).cf65
        nw572_[int(9)] = d_3528_v28_.f16
        nw572_[int(10)] = d_3528_v28_.f16
        nw572_[int(11)] = d_3528_v28_.f16
        nw572_[int(12)] = d_3528_v28_.f16
        nw572_[int(13)] = d_3511_v12_
        nw572_[int(14)] = d_3528_v28_.f16
        nw572_[int(15)] = d_3511_v12_
        nw572_[int(16)] = d_3528_v28_.f16
        nw572_[int(17)] = d_3511_v12_
        nw572_[int(18)] = d_3528_v28_.f16
        nw572_[int(19)] = d_3528_v28_.f16
        nw572_[int(20)] = d_3511_v12_
        nw572_[int(21)] = d_3528_v28_.f16
        nw572_[int(22)] = False
        nw572_[int(23)] = ((d_3512_v13_)[937] if (937) in (d_3512_v13_) else d_3511_v12_)
        nw572_[int(24)] = d_3528_v28_.f16
        nw572_[int(25)] = default__.fm13(d_3528_v28_.f16, (d_3531_v31_).cf79, (self).f9, globalState)
        nw572_[int(26)] = d_3511_v12_
        d_3532_v32_ = nw572_
        d_3533_v33_: _dafny.Map
        d_3533_v33_ = _dafny.Map({d_3532_v32_: (769) == (742)})
        r1 = ((d_3533_v33_)[d_3532_v32_] if (d_3532_v32_) in (d_3533_v33_) else not(d_3528_v28_.f16))
        d_3534_v34_: _dafny.Map
        d_3534_v34_ = _dafny.Map({d_3528_v28_.f16: False})
        d_3535_v35_: _dafny.Map
        d_3535_v35_ = _dafny.Map({d_3534_v34_: not(d_3511_v12_)})
        d_3536_v36_: _dafny.MultiSet
        d_3536_v36_ = _dafny.MultiSet([(0) - ((d_3528_v28_).f15)])
        d_3537_v37_: _dafny.Map
        d_3537_v37_ = _dafny.Map({self.f13: p1})
        d_3538_v38_: _dafny.Array
        nw573_ = _dafny.Array(None, 20)
        nw573_[int(0)] = len(d_3535_v35_)
        nw573_[int(1)] = self.f13
        nw573_[int(2)] = len((d_3515_v16_).f9)
        nw573_[int(3)] = (default__.fm0((d_3528_v28_).f15, p0, globalState)) + (p0)
        nw573_[int(4)] = (d_3528_v28_).f15
        nw573_[int(5)] = p2
        nw573_[int(6)] = 33
        nw573_[int(7)] = ((d_3536_v36_).set((self).f11, default__.abs(self.f13))).cardinality
        nw573_[int(8)] = (self.f13) - ((d_3528_v28_).f15)
        nw573_[int(9)] = 159
        nw573_[int(10)] = (p2) * ((self).fm5(len(default__.fm1(p0, p2, globalState)), len(_dafny.Map({d_3511_v12_: d_3511_v12_})), globalState))
        nw573_[int(11)] = p0
        nw573_[int(12)] = p1
        nw573_[int(13)] = (self.f13 if d_3528_v28_.f16 else p0)
        nw573_[int(14)] = (self).f11
        nw573_[int(15)] = p1
        nw573_[int(16)] = ((self).f11) * (len(d_3514_v15_))
        nw573_[int(17)] = self.f13
        nw573_[int(18)] = ((d_3537_v37_)[(0) - (p0)] if ((0) - (p0)) in (d_3537_v37_) else (d_3528_v28_).f15)
        nw573_[int(19)] = (d_3528_v28_).f15
        d_3538_v38_ = nw573_
        r2 = d_3538_v38_
        r3 = default__.fm0(p2, (self).f11, globalState)
        return r0, r1, r2, r3

    def m4(self, p0, globalState):
        d_3539_v0_: bool
        d_3539_v0_ = False
        if d_3539_v0_:
            if True:
                d_3540_v1_: _dafny.Seq
                d_3540_v1_ = _dafny.SeqWithoutIsStrInference([default__.fm38(p0, globalState), d_3539_v0_])
                d_3541_v2_: _dafny.MultiSet
                d_3541_v2_ = _dafny.MultiSet([d_3539_v0_, d_3539_v0_])
                d_3542_v3_: _dafny.Map
                d_3542_v3_ = _dafny.Map({(d_3541_v2_).cardinality: d_3539_v0_})
                d_3543_v4_: _dafny.Map
                d_3543_v4_ = _dafny.Map({self.f13: len(d_3542_v3_)})
                d_3544_v5_: _dafny.Map
                d_3544_v5_ = _dafny.Map({d_3540_v1_: (d_3543_v4_) | (_dafny.Map({p0: len((self).f9)}))})
                d_3545_v6_: _dafny.Array
                def lambda165_(d_3546_v0_, d_3547_p0_):
                    def lambda166_(d_3548_i0_):
                        return not((d_3546_v0_ if d_3546_v0_ else (D20_DC47(d_3546_v0_, d_3547_p0_, d_3546_v0_)).cf67))

                    return lambda166_

                init93_ = lambda165_(d_3539_v0_, p0)
                nw574_ = _dafny.Array(None, 29)
                for i0_93_ in range(nw574_.length(0)):
                    nw574_[i0_93_] = init93_(i0_93_)
                d_3545_v6_ = nw574_
                index486_ = default__.safeIndex(533, (d_3545_v6_).length(0))
                (d_3545_v6_)[index486_] = (d_3540_v1_) < (d_3540_v1_)
                d_3549_v7_: _dafny.Array
                nw575_ = _dafny.Array(None, 11)
                d_3549_v7_ = nw575_
                d_3550_v8_: _dafny.MultiSet
                d_3550_v8_ = _dafny.MultiSet([d_3549_v7_])
                d_3551_v9_: _dafny.Seq
                d_3551_v9_ = _dafny.SeqWithoutIsStrInference([len((self).f9), ((d_3550_v8_).cardinality) - ((self).f11), 18, len(d_3540_v1_)])
                d_3552_v10_: str
                d_3552_v10_ = _dafny.CodePoint('g')
                d_3553_v11_: D12
                d_3553_v11_ = D12_DC29(d_3539_v0_, len(_dafny.Map({d_3539_v0_: d_3552_v10_})), d_3551_v9_, d_3539_v0_, p0)
                d_3554_v12_: _dafny.MultiSet
                d_3554_v12_ = _dafny.MultiSet([d_3545_v6_, d_3545_v6_, d_3545_v6_, d_3545_v6_, d_3545_v6_])
                d_3555_v13_: _dafny.Seq
                d_3555_v13_ = _dafny.SeqWithoutIsStrInference([d_3554_v12_])
                index487_ = default__.safeIndex(533, (d_3545_v6_).length(0))
                rhs506_ = (d_3544_v5_) | (_dafny.Map({(_dafny.SeqWithoutIsStrInference([d_3539_v0_])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([d_3539_v0_]))), d_3539_v0_): d_3543_v4_}))
                rhs507_ = (d_3553_v11_).cf36
                rhs508_ = default__.safeDivisionInt(p0, (0) - (self.f13))
                rhs509_ = (d_3554_v12_).issubset((d_3555_v13_)[default__.safeIndex(p0, len(d_3555_v13_))])
                rhs510_ = _dafny.SeqWithoutIsStrInference([p0, (self).fm5((self).f11, p0, globalState), self.f13])
                lhs320_ = globalState
                lhs321_ = globalState
                lhs322_ = d_3545_v6_
                lhs323_ = default__.safeIndex(533, (d_3545_v6_).length(0))
                d_3544_v5_ = rhs506_
                lhs320_.f3 = rhs507_
                lhs321_.f3 = rhs508_
                lhs322_[lhs323_] = rhs509_
                d_3551_v9_ = rhs510_
                (globalState).f3 = self.f13
                d_3545_v6_ = d_3545_v6_
                d_3556_v14_: _dafny.Array
                nw576_ = _dafny.Array(_dafny.MultiSet({}), 14)
                d_3556_v14_ = nw576_
                def lambda167_(d_3557_v1_):
                    def lambda168_(d_3558_i1_):
                        return _dafny.MultiSet(d_3557_v1_)

                    return lambda168_

                init94_ = lambda167_(d_3540_v1_)
                nw577_ = _dafny.Array(None, 10)
                for i0_94_ in range(nw577_.length(0)):
                    nw577_[i0_94_] = init94_(i0_94_)
                d_3556_v14_ = nw577_
                d_3559_v15_: _dafny.Array
                nw578_ = _dafny.Array(int(0), 8)
                d_3559_v15_ = nw578_
                d_3560_v16_: _dafny.Set
                d_3560_v16_ = _dafny.Set({d_3552_v10_})
                d_3561_v17_: _dafny.MultiSet
                d_3561_v17_ = _dafny.MultiSet([len(d_3560_v16_), (self).f11])
                index488_ = default__.safeIndex(561, (d_3559_v15_).length(0))
                (d_3559_v15_)[index488_] = default__.safeModuloInt(((d_3561_v17_)[self.f13] if (self.f13) in (d_3561_v17_) else (0) - (p0)), self.f13)
                index489_ = default__.safeIndex(561, (d_3559_v15_).length(0))
                (d_3559_v15_)[index489_] = (-85) * (len(d_3543_v4_))
            elif True:
                d_3562_v18_: _dafny.Seq
                d_3562_v18_ = _dafny.SeqWithoutIsStrInference([d_3539_v0_])
                d_3563_v19_: _dafny.Seq
                d_3563_v19_ = _dafny.SeqWithoutIsStrInference([len(d_3562_v18_), (self).f11])
                d_3564_v20_: D18
                d_3564_v20_ = D18_DC42(d_3539_v0_, 617)
                d_3565_v21_: _dafny.Set
                d_3565_v21_ = _dafny.Set({(d_3564_v20_).cf59})
                d_3566_v22_: C4
                nw579_ = C4()
                nw579_.ctor__(d_3539_v0_, d_3539_v0_, ((self).f9) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owf"))), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f11, (self).f11]), (d_3563_v19_).set(default__.safeIndex(len(d_3565_v21_), len(d_3563_v19_)), p0), (default__.fm44(not(d_3539_v0_), self.f13, d_3539_v0_, globalState)).set(default__.safeIndex(p0, len(default__.fm44(not(d_3539_v0_), self.f13, d_3539_v0_, globalState))), p0)]))
                d_3566_v22_ = nw579_
                (d_3566_v22_).f20 = d_3566_v22_.f19
                d_3539_v0_ = d_3566_v22_.f20
                (d_3566_v22_).f20 = default__.fm22(d_3539_v0_, 787, globalState)
                d_3567_v23_: str
                d_3567_v23_ = _dafny.CodePoint('b')
                d_3568_v24_: _dafny.Map
                d_3568_v24_ = _dafny.Map({self.f13: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))})
                d_3569_v25_: _dafny.Array
                nw580_ = _dafny.Array(None, 26)
                nw580_[int(0)] = d_3567_v23_
                nw580_[int(1)] = _dafny.CodePoint('y')
                nw580_[int(2)] = d_3567_v23_
                nw580_[int(3)] = (d_3567_v23_ if (D20_DC47(d_3566_v22_.f20, (0) - (len(d_3568_v24_)), False)).cf67 else d_3567_v23_)
                nw580_[int(4)] = d_3567_v23_
                nw580_[int(5)] = d_3567_v23_
                nw580_[int(6)] = d_3567_v23_
                nw580_[int(7)] = d_3567_v23_
                nw580_[int(8)] = (d_3567_v23_ if d_3566_v22_.f20 else d_3567_v23_)
                nw580_[int(9)] = d_3567_v23_
                nw580_[int(10)] = d_3567_v23_
                nw580_[int(11)] = d_3567_v23_
                nw580_[int(12)] = _dafny.CodePoint('l')
                nw580_[int(13)] = d_3567_v23_
                nw580_[int(14)] = ((self).f9)[default__.safeIndex((0) - ((self).f11), len((self).f9))]
                nw580_[int(15)] = d_3567_v23_
                nw580_[int(16)] = d_3567_v23_
                nw580_[int(17)] = d_3567_v23_
                nw580_[int(18)] = d_3567_v23_
                nw580_[int(19)] = d_3567_v23_
                nw580_[int(20)] = d_3567_v23_
                nw580_[int(21)] = d_3567_v23_
                nw580_[int(22)] = (d_3567_v23_ if d_3539_v0_ else d_3567_v23_)
                nw580_[int(23)] = d_3567_v23_
                nw580_[int(24)] = d_3567_v23_
                nw580_[int(25)] = (self).fm6(globalState)
                d_3569_v25_ = nw580_
                d_3569_v25_ = d_3569_v25_
            (self).f13 = p0
            d_3570_v26_: str
            d_3570_v26_ = _dafny.CodePoint('g')
            d_3570_v26_ = d_3570_v26_
            d_3571_v27_: _dafny.Map
            d_3571_v27_ = _dafny.Map({-643: (self).f11})
            d_3572_v28_: _dafny.Array
            nw581_ = _dafny.Array(D1.default()(), 24)
            d_3572_v28_ = nw581_
            d_3573_v29_: _dafny.Array
            nw582_ = _dafny.Array(None, 4)
            nw582_[int(0)] = d_3572_v28_
            nw582_[int(1)] = d_3572_v28_
            nw582_[int(2)] = d_3572_v28_
            nw582_[int(3)] = d_3572_v28_
            d_3573_v29_ = nw582_
            d_3574_v30_: C5
            nw583_ = C5()
            nw583_.ctor__(((d_3571_v27_)[(self).f11] if ((self).f11) in (d_3571_v27_) else (self).f11), d_3573_v29_)
            d_3574_v30_ = nw583_
            d_3575_v31_: _dafny.Seq
            d_3575_v31_ = _dafny.SeqWithoutIsStrInference([d_3539_v0_, d_3539_v0_])
            d_3576_v32_: _dafny.Map
            d_3576_v32_ = _dafny.Map({d_3575_v31_: d_3539_v0_})
            d_3577_v33_: D17
            d_3577_v33_ = D17_DC40((self).f9, d_3539_v0_, self.f13)
            d_3576_v32_ = (d_3576_v32_).set(_dafny.SeqWithoutIsStrInference([d_3539_v0_, d_3539_v0_, (d_3577_v33_).cf55, d_3539_v0_, d_3539_v0_]), d_3539_v0_)
        elif True:
            d_3578_v34_: _dafny.Map
            d_3578_v34_ = _dafny.Map({self.f13: True})
            if ((d_3578_v34_)[((self).f11) - (p0)] if (((self).f11) - (p0)) in (d_3578_v34_) else d_3539_v0_):
                d_3579_v35_: str
                d_3579_v35_ = _dafny.CodePoint('a')
                d_3580_v36_: _dafny.Map
                d_3580_v36_ = _dafny.Map({(default__.fm31(globalState)).set(default__.safeIndex(len(((self).f9).set(default__.safeIndex(p0, len((self).f9)), d_3579_v35_)), len(default__.fm31(globalState))), d_3579_v35_): (((self).f9) + ((self).f9)).set(default__.safeIndex(p0, len(((self).f9) + ((self).f9))), d_3579_v35_)})
                d_3581_v37_: _dafny.Seq
                d_3581_v37_ = _dafny.SeqWithoutIsStrInference([False])
                rhs511_ = len((self).f9)
                rhs512_ = d_3580_v36_
                rhs513_ = ((len(d_3581_v37_)) * ((self).f11)) * ((self.f13) - ((0) - ((self).f11)))
                lhs324_ = globalState
                lhs325_ = globalState
                lhs324_.f3 = rhs511_
                d_3580_v36_ = rhs512_
                lhs325_.f3 = rhs513_
                d_3582_v38_: C2
                nw584_ = C2()
                nw584_.ctor__(default__.safeModuloInt((self).f11, self.f13), self.f12, (self).f9, default__.fm41(d_3539_v0_, d_3539_v0_, False, globalState))
                d_3582_v38_ = nw584_
                d_3583_v39_: _dafny.Array
                def lambda169_(d_3584_p0_):
                    def lambda170_(d_3585_i2_):
                        return _dafny.MultiSet([d_3584_p0_, d_3584_p0_])

                    return lambda170_

                init95_ = lambda169_(p0)
                nw585_ = _dafny.Array(None, 5)
                for i0_95_ in range(nw585_.length(0)):
                    nw585_[i0_95_] = init95_(i0_95_)
                d_3583_v39_ = nw585_
                d_3586_v40_: _dafny.MultiSet
                d_3586_v40_ = _dafny.MultiSet([(self).f11])
                index490_ = default__.safeIndex(127, (d_3583_v39_).length(0))
                (d_3583_v39_)[index490_] = d_3586_v40_
                index491_ = default__.safeIndex(127, (d_3583_v39_).length(0))
                (d_3583_v39_)[index491_] = (d_3586_v40_) - (d_3586_v40_)
                d_3587_v41_: _dafny.Seq
                d_3587_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyv"))
                d_3588_v42_: _dafny.Set
                d_3588_v42_ = _dafny.Set({p0, p0})
                d_3589_v43_: _dafny.Map
                d_3589_v43_ = _dafny.Map({default__.fm22(False, (self).f11, globalState): d_3588_v42_})
                d_3590_v44_: D18
                d_3590_v44_ = D18_DC42(d_3539_v0_, self.f13)
                d_3591_v45_: _dafny.Seq
                d_3591_v45_ = _dafny.SeqWithoutIsStrInference([((d_3589_v43_)[d_3539_v0_] if (d_3539_v0_) in (d_3589_v43_) else d_3588_v42_), _dafny.Set({self.f13, (self).f11, self.f13, (d_3590_v44_).cf59, len((self).f9)})])
                d_3592_v46_: D8
                d_3592_v46_ = D8_DC18()
                d_3593_v47_: _dafny.Seq
                d_3593_v47_ = _dafny.SeqWithoutIsStrInference([d_3592_v46_, d_3592_v46_])
                d_3594_v48_: _dafny.Array
                nw586_ = _dafny.Array(False, 16)
                d_3594_v48_ = nw586_
                d_3595_v49_: _dafny.Map
                d_3595_v49_ = _dafny.Map({d_3539_v0_: d_3594_v48_})
                d_3596_v50_: _dafny.MultiSet
                d_3596_v50_ = _dafny.MultiSet([d_3539_v0_, d_3539_v0_])
                d_3597_v51_: _dafny.Seq
                d_3597_v51_ = _dafny.SeqWithoutIsStrInference([(d_3596_v50_).cardinality])
                d_3598_v52_: _dafny.Array
                nw587_ = _dafny.Array(None, 17)
                nw587_[int(0)] = self.f13
                nw587_[int(1)] = (0) - ((self).f11)
                nw587_[int(2)] = p0
                nw587_[int(3)] = self.f13
                nw587_[int(4)] = p0
                nw587_[int(5)] = len(_dafny.SeqWithoutIsStrInference([self.f13 for d_3599_i3_ in range(default__.abs(148))]))
                nw587_[int(6)] = p0
                nw587_[int(7)] = p0
                nw587_[int(8)] = len(d_3587_v41_)
                nw587_[int(9)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ak")))
                nw587_[int(10)] = (len(d_3591_v45_)) - (len(d_3593_v47_))
                nw587_[int(11)] = 172
                nw587_[int(12)] = (len(d_3595_v49_) if not(d_3539_v0_) else (d_3597_v51_)[default__.safeIndex((self).f11, len(d_3597_v51_))])
                nw587_[int(13)] = self.f13
                nw587_[int(14)] = len((d_3587_v41_) + ((self).f9))
                nw587_[int(15)] = default__.fm0(self.f13, (self).f11, globalState)
                nw587_[int(16)] = (self).f11
                d_3598_v52_ = nw587_
                index492_ = default__.safeIndex(795, (d_3598_v52_).length(0))
                (d_3598_v52_)[index492_] = (self).f11
                index493_ = default__.safeIndex(795, (d_3598_v52_).length(0))
                rhs514_ = self.f13
                rhs515_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdtkw"))) + (d_3587_v41_)) + (d_3587_v41_)
                rhs516_ = 48
                lhs326_ = globalState
                lhs327_ = d_3598_v52_
                lhs328_ = default__.safeIndex(795, (d_3598_v52_).length(0))
                lhs326_.f3 = rhs514_
                d_3587_v41_ = rhs515_
                lhs327_[lhs328_] = rhs516_
                d_3600_v53_: C4
                nw588_ = C4()
                nw588_.ctor__(d_3539_v0_, d_3539_v0_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pgtfor"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbwtbw"))), (self).f10)
                d_3600_v53_ = nw588_
            elif True:
                (globalState).f3 = p0
                (globalState).f3 = (0) - (p0)
                (self).f13 = (self).f11
                d_3601_v54_: _dafny.MultiSet
                d_3601_v54_ = _dafny.MultiSet([d_3539_v0_])
                d_3602_v55_: _dafny.Set
                d_3602_v55_ = _dafny.Set({self.f13, (self).f11})
                d_3603_v56_: _dafny.Seq
                d_3603_v56_ = _dafny.SeqWithoutIsStrInference([d_3602_v55_])
                d_3604_v57_: str
                d_3604_v57_ = _dafny.CodePoint('u')
                d_3605_v58_: _dafny.MultiSet
                d_3605_v58_ = _dafny.MultiSet([(self).f11, (self).f11, (self).f11])
                d_3606_v59_: _dafny.Map
                d_3606_v59_ = _dafny.Map({d_3604_v57_: (d_3605_v58_).cardinality})
                d_3607_v60_: _dafny.Map
                d_3607_v60_ = _dafny.Map({d_3602_v55_: p0})
                d_3608_v62_: _dafny.Map
                d_3608_v62_ = _dafny.Map({not(d_3539_v0_): (self).f9})
                d_3609_v63_: _dafny.Array
                nw589_ = _dafny.Array(None, 23)
                nw589_[int(0)] = not (d_3539_v0_) or (d_3539_v0_)
                nw589_[int(1)] = not(d_3539_v0_)
                nw589_[int(2)] = d_3539_v0_
                nw589_[int(3)] = d_3539_v0_
                nw589_[int(4)] = ((_dafny.MultiSet([not(False)])).set(not(d_3539_v0_), default__.abs(p0))).isdisjoint(d_3601_v54_)
                nw589_[int(5)] = (d_3603_v56_) == (d_3603_v56_)
                nw589_[int(6)] = (True) or (d_3539_v0_)
                nw589_[int(7)] = d_3539_v0_
                nw589_[int(8)] = (d_3606_v59_) != ((d_3606_v59_).set(_dafny.CodePoint('i'), p0))
                nw589_[int(9)] = d_3539_v0_
                nw589_[int(10)] = not(False)
                nw589_[int(11)] = d_3539_v0_
                nw589_[int(12)] = d_3539_v0_
                nw589_[int(13)] = d_3539_v0_
                nw589_[int(14)] = (d_3539_v0_) == (False)
                nw589_[int(15)] = d_3539_v0_
                nw589_[int(16)] = ((self).f9) != (((self).f9).set(default__.safeIndex(self.f13, len((self).f9)), d_3604_v57_))
                def iife256_():
                    coll104_ = _dafny.Set()
                    compr_104_: _dafny.Set
                    for compr_104_ in (d_3607_v60_).keys.Elements:
                        d_3610_v61_: _dafny.Set = compr_104_
                        if (d_3610_v61_) in (d_3607_v60_):
                            coll104_ = coll104_.union(_dafny.Set([d_3610_v61_]))
                    return _dafny.Set(coll104_)
                nw589_[int(17)] = (iife256_()
                ).isdisjoint(_dafny.Set({d_3602_v55_, d_3602_v55_, d_3602_v55_, d_3602_v55_, d_3602_v55_}))
                nw589_[int(18)] = default__.fm13(d_3539_v0_, p0, ((d_3608_v62_)[d_3539_v0_] if (d_3539_v0_) in (d_3608_v62_) else (self).f9), globalState)
                nw589_[int(19)] = True
                nw589_[int(20)] = d_3539_v0_
                nw589_[int(21)] = d_3539_v0_
                nw589_[int(22)] = True
                d_3609_v63_ = nw589_
                index494_ = default__.safeIndex(652, (d_3609_v63_).length(0))
                (d_3609_v63_)[index494_] = d_3539_v0_
                d_3611_v64_: C4
                nw590_ = C4()
                nw590_.ctor__(True, d_3539_v0_, (self).f9, (self).f10)
                d_3611_v64_ = nw590_
                d_3612_v65_: _dafny.Map
                d_3612_v65_ = _dafny.Map({d_3611_v64_: d_3611_v64_.f20})
                d_3613_v66_: C4
                d_3613_v66_ = d_3611_v64_
                d_3614_v67_: _dafny.Seq
                d_3614_v67_ = _dafny.SeqWithoutIsStrInference([not(((d_3612_v65_)[(d_3613_v66_)] if ((d_3613_v66_)) in (d_3612_v65_) else d_3611_v64_.f20))])
                d_3615_v68_: C7
                nw591_ = C7()
                nw591_.ctor__((self).f9, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f13 for d_3616_i5_ in range(default__.abs(-506))]) for d_3617_i4_ in range(default__.abs(563))]), self.f13, self.f12)
                d_3615_v68_ = nw591_
                d_3618_v69_: _dafny.Map
                d_3618_v69_ = _dafny.Map({(self).f11: d_3615_v68_})
                d_3619_v70_: _dafny.Map
                d_3619_v70_ = _dafny.Map({d_3539_v0_: d_3539_v0_})
                d_3620_v71_: _dafny.Array
                nw592_ = _dafny.Array(None, 11)
                nw592_[int(0)] = d_3615_v68_
                nw592_[int(1)] = d_3615_v68_
                nw592_[int(2)] = d_3615_v68_
                nw592_[int(3)] = d_3615_v68_
                nw592_[int(4)] = d_3615_v68_
                nw592_[int(5)] = d_3615_v68_
                nw592_[int(6)] = d_3615_v68_
                nw592_[int(7)] = d_3615_v68_
                nw592_[int(8)] = d_3615_v68_
                nw592_[int(9)] = ((d_3618_v69_)[len(d_3619_v70_)] if (len(d_3619_v70_)) in (d_3618_v69_) else d_3615_v68_)
                nw592_[int(10)] = d_3615_v68_
                d_3620_v71_ = nw592_
                d_3621_v72_: _dafny.MultiSet
                d_3621_v72_ = _dafny.MultiSet([d_3620_v71_, d_3620_v71_])
                index495_ = default__.safeIndex(652, (d_3609_v63_).length(0))
                rhs517_ = p0
                rhs518_ = ((d_3578_v34_)[(self).f11] if ((self).f11) in (d_3578_v34_) else d_3611_v64_.f19)
                rhs519_ = self.f13
                rhs520_ = (d_3614_v67_).set(default__.safeIndex(self.f13, len(d_3614_v67_)), d_3539_v0_)
                rhs521_ = (_dafny.MultiSet([d_3620_v71_])).ispropersubset(d_3621_v72_)
                lhs329_ = globalState
                lhs330_ = d_3609_v63_
                lhs331_ = default__.safeIndex(652, (d_3609_v63_).length(0))
                lhs332_ = globalState
                lhs333_ = d_3611_v64_
                lhs329_.f3 = rhs517_
                lhs330_[lhs331_] = rhs518_
                lhs332_.f3 = rhs519_
                d_3614_v67_ = rhs520_
                lhs333_.f20 = rhs521_
                index496_ = default__.safeIndex(652, (d_3609_v63_).length(0))
                (d_3609_v63_)[index496_] = d_3611_v64_.f20
            d_3622_v73_: _dafny.Set
            d_3622_v73_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwnppnol")), (self).f9, (self).f9})
            d_3623_v74_: _dafny.Map
            d_3623_v74_ = _dafny.Map({d_3539_v0_: len(d_3622_v73_)})
            d_3624_v75_: _dafny.Seq
            d_3624_v75_ = _dafny.SeqWithoutIsStrInference([d_3623_v74_, (d_3623_v74_) | (d_3623_v74_), d_3623_v74_])
            d_3624_v75_ = ((d_3624_v75_) + (d_3624_v75_)) + (d_3624_v75_)
            d_3625_v76_: _dafny.Seq
            d_3625_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjmg"))
            d_3625_v76_ = (self).f9
            d_3626_v77_: _dafny.Array
            nw593_ = _dafny.Array(False, 12)
            d_3626_v77_ = nw593_
            index497_ = default__.safeIndex(532, (d_3626_v77_).length(0))
            (d_3626_v77_)[index497_] = d_3539_v0_
            index498_ = default__.safeIndex(532, (d_3626_v77_).length(0))
            (d_3626_v77_)[index498_] = d_3539_v0_
            d_3627_v78_: D18
            d_3627_v78_ = D18_DC42(d_3539_v0_, p0)
            d_3628_v79_: _dafny.Seq
            d_3628_v79_ = _dafny.SeqWithoutIsStrInference([d_3627_v78_, d_3627_v78_])
            d_3629_v80_: _dafny.Map
            d_3629_v80_ = _dafny.Map({(d_3626_v77_)[default__.safeIndex(532, (d_3626_v77_).length(0))]: d_3539_v0_})
            d_3628_v79_ = (d_3628_v79_).set(default__.safeIndex((self.f13) + (p0), len(d_3628_v79_)), D18_DC42(d_3539_v0_, len((d_3629_v80_).set(default__.fm38(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_3630_i6_ in range(default__.abs(396))])), globalState), (d_3626_v77_)[default__.safeIndex(532, (d_3626_v77_).length(0))]))))
        d_3631_v81_: str
        d_3631_v81_ = _dafny.CodePoint('h')
        d_3632_v82_: _dafny.Map
        d_3632_v82_ = _dafny.Map({True: d_3539_v0_})
        d_3633_v83_: D27
        d_3633_v83_ = D27_DC64(d_3632_v82_, d_3539_v0_, p0, _dafny.SeqWithoutIsStrInference([self.f13 for d_3634_i8_ in range(default__.abs(-575))]), (self).f11)
        d_3635_v84_: _dafny.Array
        nw594_ = _dafny.Array(None, 24)
        nw594_[int(0)] = d_3539_v0_
        nw594_[int(1)] = d_3539_v0_
        nw594_[int(2)] = False
        nw594_[int(3)] = default__.fm22(d_3539_v0_, (self).f11, globalState)
        nw594_[int(4)] = d_3539_v0_
        nw594_[int(5)] = not (d_3539_v0_) or (not(d_3539_v0_))
        nw594_[int(6)] = d_3539_v0_
        nw594_[int(7)] = not (d_3539_v0_) or (d_3539_v0_)
        nw594_[int(8)] = False
        nw594_[int(9)] = (d_3631_v81_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pj")))
        nw594_[int(10)] = d_3539_v0_
        nw594_[int(11)] = d_3539_v0_
        nw594_[int(12)] = default__.fm13(d_3539_v0_, p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")), globalState)
        nw594_[int(13)] = d_3539_v0_
        nw594_[int(14)] = d_3539_v0_
        nw594_[int(15)] = d_3539_v0_
        nw594_[int(16)] = True
        nw594_[int(17)] = d_3539_v0_
        nw594_[int(18)] = d_3539_v0_
        nw594_[int(19)] = d_3539_v0_
        nw594_[int(20)] = ((_dafny.MultiSet([p0])).cardinality) != ((d_3633_v83_).cf101)
        nw594_[int(21)] = d_3539_v0_
        nw594_[int(22)] = d_3539_v0_
        nw594_[int(23)] = d_3539_v0_
        d_3635_v84_ = nw594_
        guard_loop_12_: int
        for guard_loop_12_ in _dafny.IntegerRange(0, (d_3635_v84_).length(0)):
            d_3636_i7_: int = guard_loop_12_
            if (True) and (((0) <= (d_3636_i7_)) and ((d_3636_i7_) < ((d_3635_v84_).length(0)))):
                def iife257_():
                    coll105_ = _dafny.Map()
                    compr_105_: int
                    for compr_105_ in _dafny.IntegerRange(-201, 481):
                        d_3639_v85_: int = compr_105_
                        if ((-201) <= (d_3639_v85_)) and ((d_3639_v85_) < (481)):
                            coll105_[(d_3639_v85_) + (len(((self).f9).set(default__.safeIndex(382, len((self).f9)), d_3631_v81_)))] = d_3539_v0_
                    return _dafny.Map(coll105_)
                (d_3635_v84_)[(d_3636_i7_)] = (len((_dafny.Map({self.f13: _dafny.Map({D10_DC22(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_3631_v81_ for d_3637_i9_ in range(default__.abs(187))]), ((self).f9).set(default__.safeIndex((0) - (self.f13), len((self).f9)), d_3631_v81_), (self).f9])): len(_dafny.Set({d_3539_v0_, d_3539_v0_, d_3539_v0_}))})})) | (_dafny.Map({self.f13: _dafny.Map({D10_DC22(_dafny.MultiSet([(self).f9, _dafny.SeqWithoutIsStrInference([d_3631_v81_ for d_3638_i10_ in range(default__.abs(833))])])): p0})})))) <= (len((iife257_()
                ) | (_dafny.Map({150: d_3539_v0_}))))
        d_3640_v86_: _dafny.Seq
        d_3640_v86_ = _dafny.SeqWithoutIsStrInference([d_3539_v0_])
        d_3640_v86_ = d_3640_v86_
        d_3641_i11_: int
        d_3641_i11_ = 0
        with _dafny.label("16"):
            while True:
                with _dafny.c_label("16"):
                    if (d_3641_i11_) >= (100):
                        raise _dafny.Break("16")
                    d_3641_i11_ = (d_3641_i11_) + (1)
                    d_3635_v84_ = d_3635_v84_
                    nw595_ = _dafny.Array(False, 23)
                    d_3635_v84_ = nw595_
                    d_3539_v0_ = default__.fm38((0) - (p0), globalState)
                    (globalState).f3 = (self).f11
                    pass
            pass
        if ((p0) <= (self.f13) if d_3539_v0_ else d_3539_v0_):
            d_3642_v87_: _dafny.Array
            def lambda171_(d_3643_i12_):
                return default__.safeModuloInt(d_3643_i12_, (self).f11)

            init96_ = lambda171_
            nw596_ = _dafny.Array(None, 1)
            for i0_96_ in range(nw596_.length(0)):
                nw596_[i0_96_] = init96_(i0_96_)
            d_3642_v87_ = nw596_
            index499_ = default__.safeIndex(990, (d_3642_v87_).length(0))
            (d_3642_v87_)[index499_] = p0
            index500_ = default__.safeIndex(668, (d_3642_v87_).length(0))
            (d_3642_v87_)[index500_] = self.f13
            d_3644_v88_: _dafny.MultiSet
            d_3644_v88_ = _dafny.MultiSet([(self.f13) * ((0) - ((self).f11))])
            d_3645_v89_: D8
            d_3645_v89_ = D8_DC18()
            index501_ = default__.safeIndex(990, (d_3642_v87_).length(0))
            index502_ = default__.safeIndex(668, (d_3642_v87_).length(0))
            rhs522_ = (d_3539_v0_) and (not (d_3539_v0_) or (d_3539_v0_))
            rhs523_ = self.f13
            rhs524_ = ((0) - (p0)) - (p0)
            rhs525_ = len(default__.fm74(d_3645_v89_, _dafny.MultiSet((d_3640_v86_) + (d_3640_v86_)), (d_3640_v86_)[default__.safeIndex((self).f11, len(d_3640_v86_))], globalState))
            rhs526_ = (d_3644_v88_) | (d_3644_v88_)
            lhs334_ = self
            lhs335_ = d_3642_v87_
            lhs336_ = default__.safeIndex(990, (d_3642_v87_).length(0))
            lhs337_ = d_3642_v87_
            lhs338_ = default__.safeIndex(668, (d_3642_v87_).length(0))
            d_3539_v0_ = rhs522_
            lhs334_.f13 = rhs523_
            lhs335_[lhs336_] = rhs524_
            lhs337_[lhs338_] = rhs525_
            d_3644_v88_ = rhs526_
            (self).f13 = p0
            (globalState).f3 = (self).f11
            d_3646_v90_: D3
            d_3646_v90_ = D3_DC7(d_3539_v0_)
            pat_let_tv70_ = p0
            pat_let_tv71_ = globalState
            def iife258_(_pat_let76_0):
                def iife259_(d_3647_dt__update__tmp_h0_):
                    def iife260_(_pat_let77_0):
                        def iife261_(d_3648_dt__update_hcf9_h0_):
                            return D3_DC7(d_3648_dt__update_hcf9_h0_)
                        return iife261_(_pat_let77_0)
                    return iife260_((default__.fm0(pat_let_tv70_, (self).f11, pat_let_tv71_)) >= ((self).f11))
                return iife259_(_pat_let76_0)
            source55_ = iife258_(d_3646_v90_)
            if source55_.is_DC7:
                d_3649___mcc_h0_ = source55_.cf9
                d_3650_cf9_ = d_3649___mcc_h0_
                d_3651_v91_: _dafny.Array
                nw597_ = _dafny.Array(None, 27)
                d_3651_v91_ = nw597_
                d_3652_v92_: _dafny.Set
                d_3652_v92_ = _dafny.Set({d_3650_cf9_})
                d_3653_v93_: _dafny.Map
                d_3653_v93_ = _dafny.Map({p0: len(d_3652_v92_)})
                d_3654_v94_: T1
                nw598_ = C7()
                nw598_.ctor__((self).f9, (self).f10, 200, _dafny.Map({_dafny.Map({p0: (d_3642_v87_)[default__.safeIndex(990, (d_3642_v87_).length(0))]}): d_3653_v93_}))
                d_3654_v94_ = nw598_
                index503_ = default__.safeIndex(993, (d_3651_v91_).length(0))
                (d_3651_v91_)[index503_] = d_3654_v94_
                index504_ = default__.safeIndex(993, (d_3651_v91_).length(0))
                (d_3651_v91_)[index504_] = d_3654_v94_
                d_3655_v95_: _dafny.Map
                d_3655_v95_ = _dafny.Map({False: p0})
                d_3539_v0_ = (default__.safeDivisionInt(self.f13, (d_3642_v87_)[default__.safeIndex(990, (d_3642_v87_).length(0))])) != ((self).fm5((self).fm5(len((self).f9), len(d_3655_v95_), globalState), p0, globalState))
                d_3656_v96_: _dafny.Seq
                d_3656_v96_ = _dafny.SeqWithoutIsStrInference([(d_3654_v94_).f11])
                d_3657_v97_: _dafny.Map
                d_3657_v97_ = _dafny.Map({499: d_3539_v0_})
                d_3658_v98_: _dafny.Array
                nw599_ = _dafny.Array(None, 14)
                nw599_[int(0)] = d_3656_v96_
                nw599_[int(1)] = d_3656_v96_
                nw599_[int(2)] = (d_3656_v96_) + (d_3656_v96_)
                nw599_[int(3)] = d_3656_v96_
                nw599_[int(4)] = _dafny.SeqWithoutIsStrInference([(d_3642_v87_)[default__.safeIndex(990, (d_3642_v87_).length(0))], (d_3654_v94_).fm5((d_3654_v94_).f11, 577, globalState)])
                nw599_[int(5)] = d_3656_v96_
                nw599_[int(6)] = _dafny.SeqWithoutIsStrInference([(d_3642_v87_)[default__.safeIndex(990, (d_3642_v87_).length(0))], p0, len(d_3657_v97_), self.f13])
                nw599_[int(7)] = _dafny.SeqWithoutIsStrInference([(self).f11, 622, p0])
                nw599_[int(8)] = d_3656_v96_
                nw599_[int(9)] = d_3656_v96_
                nw599_[int(10)] = _dafny.SeqWithoutIsStrInference([len(d_3657_v97_), len(_dafny.SeqWithoutIsStrInference([d_3631_v81_ for d_3659_i13_ in range(default__.abs(587))])), (self).f11])
                nw599_[int(11)] = d_3656_v96_
                nw599_[int(12)] = d_3656_v96_
                nw599_[int(13)] = d_3656_v96_
                d_3658_v98_ = nw599_
                index505_ = default__.safeIndex(377, (d_3658_v98_).length(0))
                (d_3658_v98_)[index505_] = d_3656_v96_
                d_3660_v99_: _dafny.MultiSet
                d_3660_v99_ = _dafny.MultiSet([d_3650_cf9_, d_3650_cf9_])
                index506_ = default__.safeIndex(377, (d_3658_v98_).length(0))
                (d_3658_v98_)[index506_] = _dafny.SeqWithoutIsStrInference([(d_3642_v87_)[default__.safeIndex(990, (d_3642_v87_).length(0))], self.f13, (d_3642_v87_)[default__.safeIndex(990, (d_3642_v87_).length(0))], ((0) - (len(d_3655_v95_)) if d_3539_v0_ else (_dafny.MultiSet(d_3656_v96_)).cardinality), (0) - (default__.safeDivisionInt(((d_3660_v99_)[not(d_3650_cf9_)] if (not(d_3650_cf9_)) in (d_3660_v99_) else self.f13), (d_3654_v94_).f11))])
                index507_ = default__.safeIndex(990, (d_3642_v87_).length(0))
                (d_3642_v87_)[index507_] = len(default__.fm44(d_3650_cf9_, -183, not(not(((d_3642_v87_)[default__.safeIndex(990, (d_3642_v87_).length(0))]) != ((d_3642_v87_)[default__.safeIndex(990, (d_3642_v87_).length(0))]))), globalState))
            elif True:
                d_3661___mcc_h1_ = source55_.cf8
                d_3662_cf8_ = d_3661___mcc_h1_
                d_3663_v100_: _dafny.Seq
                d_3663_v100_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frgutkyq")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_3664_i15_ in range(default__.abs(449))])])
                d_3665_v101_: _dafny.Array
                nw600_ = _dafny.Array(None, 15)
                nw600_[int(0)] = (self).f9
                nw600_[int(1)] = _dafny.SeqWithoutIsStrInference([d_3631_v81_ for d_3666_i14_ in range(default__.abs(-274))])
                nw600_[int(2)] = (d_3663_v100_)[default__.safeIndex(517, len(d_3663_v100_))]
                nw600_[int(3)] = (self).f9
                nw600_[int(4)] = (self).f9
                nw600_[int(5)] = ((self).f9) + ((self).f9)
                nw600_[int(6)] = (self).f9
                nw600_[int(7)] = (self).f9
                nw600_[int(8)] = default__.fm37(-898, (d_3642_v87_)[default__.safeIndex(990, (d_3642_v87_).length(0))], d_3539_v0_, p0, globalState)
                nw600_[int(9)] = ((self).f9).set(default__.safeIndex((self).f11, len((self).f9)), d_3631_v81_)
                nw600_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbprxsgwh"))
                nw600_[int(11)] = (self).f9
                nw600_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "leunq"))
                nw600_[int(13)] = (_dafny.SeqWithoutIsStrInference([d_3631_v81_ for d_3667_i16_ in range(default__.abs(6))])) + ((self).f9)
                nw600_[int(14)] = (self).f9
                d_3665_v101_ = nw600_
                d_3665_v101_ = d_3665_v101_
                d_3668_v102_: _dafny.Seq
                d_3668_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvcmat"))
                d_3668_v102_ = d_3668_v102_
                d_3669_v103_: _dafny.Seq
                d_3669_v103_ = _dafny.SeqWithoutIsStrInference([len(d_3668_v102_)])
                d_3670_v104_: _dafny.Map
                d_3670_v104_ = _dafny.Map({d_3669_v103_: self.f13})
                index508_ = default__.safeIndex(990, (d_3642_v87_).length(0))
                (d_3642_v87_)[index508_] = (0) - ((default__.safeModuloInt(len(d_3663_v100_), (self).f11)) - (len(d_3670_v104_)))
                d_3671_v105_: _dafny.Set
                d_3671_v105_ = _dafny.Set({d_3642_v87_, d_3642_v87_, d_3642_v87_, d_3642_v87_})
                d_3671_v105_ = ((d_3671_v105_).intersection(d_3671_v105_)).intersection(d_3671_v105_)
            index509_ = default__.safeIndex(990, (d_3642_v87_).length(0))
            index510_ = default__.safeIndex(990, (d_3642_v87_).length(0))
            rhs527_ = default__.safeModuloInt(p0, (d_3642_v87_)[default__.safeIndex(990, (d_3642_v87_).length(0))])
            rhs528_ = ((p0) * (len(d_3632_v82_))) + (default__.safeDivisionInt(p0, (d_3642_v87_)[default__.safeIndex(990, (d_3642_v87_).length(0))]))
            rhs529_ = d_3539_v0_
            rhs530_ = False
            rhs531_ = (len((self).f9) if d_3539_v0_ else ((d_3644_v88_)[self.f13] if (self.f13) in (d_3644_v88_) else (self).f11))
            lhs339_ = d_3642_v87_
            lhs340_ = default__.safeIndex(990, (d_3642_v87_).length(0))
            lhs341_ = self
            lhs342_ = d_3642_v87_
            lhs343_ = default__.safeIndex(990, (d_3642_v87_).length(0))
            lhs339_[lhs340_] = rhs527_
            lhs341_.f13 = rhs528_
            d_3539_v0_ = rhs529_
            d_3539_v0_ = rhs530_
            lhs342_[lhs343_] = rhs531_
        elif True:
            d_3631_v81_ = d_3631_v81_
            d_3672_v106_: _dafny.Seq
            d_3672_v106_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idid"))
            d_3672_v106_ = default__.fm31(globalState)
            (globalState).f3 = ((self).f11) * (self.f13)
            d_3673_v107_: _dafny.Set
            d_3673_v107_ = _dafny.Set({(self).f9, (self).f9, d_3672_v106_})
            d_3539_v0_ = (((self).f11) + (self.f13)) == ((0) - (len(d_3673_v107_)))
            (self).f13 = (self).f11
        if d_3539_v0_:
            d_3674_v108_: _dafny.MultiSet
            d_3674_v108_ = _dafny.MultiSet([d_3631_v81_])
            d_3674_v108_ = (_dafny.MultiSet([d_3631_v81_, d_3631_v81_, d_3631_v81_])) | (d_3674_v108_)
            d_3675_v109_: _dafny.Seq
            d_3675_v109_ = _dafny.SeqWithoutIsStrInference([792])
            d_3676_v110_: _dafny.Set
            d_3676_v110_ = _dafny.Set({d_3539_v0_, d_3539_v0_, False, d_3539_v0_, d_3539_v0_})
            d_3677_v111_: _dafny.Map
            d_3677_v111_ = _dafny.Map({d_3675_v109_: d_3676_v110_})
            d_3677_v111_ = (d_3677_v111_).set((d_3633_v83_).cf100, (_dafny.Set({d_3539_v0_, False})) | (d_3676_v110_))
            d_3678_v112_: _dafny.MultiSet
            d_3678_v112_ = _dafny.MultiSet([673, (self).f11, self.f13, (self).f11])
            (globalState).f3 = ((d_3678_v112_)[((0) - (p0)) - ((0) - (self.f13))] if (((0) - (p0)) - ((0) - (self.f13))) in (d_3678_v112_) else len(d_3675_v109_))
            d_3679_v113_: _dafny.Map
            d_3679_v113_ = _dafny.Map({p0: (d_3539_v0_ if d_3539_v0_ else d_3539_v0_)})
            d_3679_v113_ = d_3679_v113_
            d_3680_v114_: _dafny.Seq
            d_3680_v114_ = _dafny.SeqWithoutIsStrInference([d_3635_v84_])
            d_3635_v84_ = (d_3680_v114_)[default__.safeIndex(p0, len(d_3680_v114_))]
        elif True:
            d_3681_v115_: _dafny.MultiSet
            d_3681_v115_ = _dafny.MultiSet([self.f13])
            d_3682_v116_: _dafny.Map
            d_3682_v116_ = _dafny.Map({d_3681_v115_: (self).f11})
            d_3683_v117_: _dafny.Set
            d_3683_v117_ = _dafny.Set({(self).f9, (self).f9})
            d_3684_v118_: _dafny.Set
            d_3684_v118_ = _dafny.Set({d_3683_v117_})
            d_3685_v119_: _dafny.Seq
            d_3685_v119_ = _dafny.SeqWithoutIsStrInference([p0, default__.fm0((self).f11, self.f13, globalState)])
            d_3686_v121_: _dafny.Set
            d_3686_v121_ = _dafny.Set({d_3681_v115_})
            d_3687_v122_: D10
            d_3687_v122_ = D10_DC23(d_3631_v81_, d_3686_v121_)
            d_3688_v123_: _dafny.Seq
            def iife262_():
                coll106_ = _dafny.Map()
                compr_106_: _dafny.MultiSet
                for compr_106_ in ((d_3687_v122_).cf27).Elements:
                    d_3689_v120_: _dafny.MultiSet = compr_106_
                    if (d_3689_v120_) in ((d_3687_v122_).cf27):
                        coll106_[d_3689_v120_] = (self).f11
                return _dafny.Map(coll106_)
            d_3688_v123_ = _dafny.SeqWithoutIsStrInference([iife262_()
            ])
            d_3690_v124_: _dafny.Array
            nw601_ = _dafny.Array(None, 15)
            nw601_[int(0)] = (d_3682_v116_ if d_3539_v0_ else _dafny.Map({d_3681_v115_: p0}))
            nw601_[int(1)] = (d_3682_v116_) | (_dafny.Map({_dafny.MultiSet([p0]): len(d_3640_v86_)}))
            nw601_[int(2)] = _dafny.Map({d_3681_v115_: len(d_3684_v118_)})
            nw601_[int(3)] = _dafny.Map({_dafny.MultiSet([0]): p0})
            nw601_[int(4)] = (d_3682_v116_) | (d_3682_v116_)
            nw601_[int(5)] = d_3682_v116_
            nw601_[int(6)] = d_3682_v116_
            nw601_[int(7)] = (d_3682_v116_) | (d_3682_v116_)
            nw601_[int(8)] = _dafny.Map({d_3681_v115_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ayic")))})
            nw601_[int(9)] = default__.fm75((0) - (len(_dafny.Map({d_3539_v0_: (0) - ((d_3685_v119_)[default__.safeIndex((self).fm5(p0, (self).f11, globalState), len(d_3685_v119_))])}))), globalState)
            nw601_[int(10)] = (d_3682_v116_) | (d_3682_v116_)
            nw601_[int(11)] = (d_3688_v123_)[default__.safeIndex((self).f11, len(d_3688_v123_))]
            nw601_[int(12)] = _dafny.Map({_dafny.MultiSet([-893]): len(_dafny.SeqWithoutIsStrInference([d_3631_v81_ for d_3691_i17_ in range(default__.abs(943))]))})
            nw601_[int(13)] = d_3682_v116_
            nw601_[int(14)] = d_3682_v116_
            d_3690_v124_ = nw601_
            index511_ = default__.safeIndex(581, (d_3690_v124_).length(0))
            (d_3690_v124_)[index511_] = (d_3682_v116_) | ((_dafny.Map({d_3681_v115_: (self).f11})).set(_dafny.MultiSet(d_3685_v119_), p0))
            d_3692_v125_: _dafny.Map
            d_3692_v125_ = _dafny.Map({d_3681_v115_: self.f13})
            index512_ = default__.safeIndex(581, (d_3690_v124_).length(0))
            (d_3690_v124_)[index512_] = (d_3692_v125_)
            d_3693_v126_: _dafny.Map
            d_3693_v126_ = _dafny.Map({(self).f11: d_3539_v0_})
            d_3694_v127_: D4
            d_3694_v127_ = D4_DC8(d_3693_v126_)
            d_3695_v128_: _dafny.Map
            d_3695_v128_ = _dafny.Map({d_3694_v127_: (d_3681_v115_).cardinality})
            d_3539_v0_ = not (d_3539_v0_) or ((d_3694_v127_) in (d_3695_v128_))
            (self).f13 = 196
            index513_ = default__.safeIndex(752, (d_3635_v84_).length(0))
            (d_3635_v84_)[index513_] = d_3539_v0_
            index514_ = default__.safeIndex(752, (d_3635_v84_).length(0))
            (d_3635_v84_)[index514_] = (p0) == (self.f13)
            d_3696_v129_: _dafny.Array
            nw602_ = _dafny.Array(None, 5)
            nw602_[int(0)] = d_3631_v81_
            nw602_[int(1)] = d_3631_v81_
            nw602_[int(2)] = d_3631_v81_
            nw602_[int(3)] = d_3631_v81_
            nw602_[int(4)] = d_3631_v81_
            d_3696_v129_ = nw602_
            index515_ = default__.safeIndex(150, (d_3696_v129_).length(0))
            (d_3696_v129_)[index515_] = d_3631_v81_
            index516_ = default__.safeIndex(150, (d_3696_v129_).length(0))
            (d_3696_v129_)[index516_] = d_3631_v81_


class C11:
    def  __init__(self):
        self.f7: _dafny.Map = _dafny.Map({})
        self._f8: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C11"
    def ctor__(self, f7, f8):
        (self).f7 = f7
        (self)._f8 = f8

    def fm3(self, p0, p1, p2, p3, globalState):
        return (self).f8

    def fm4(self, p0, p1, p2, p3, globalState):
        return (753) * ((self).f8)

    def m0(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Map = _dafny.Map({})
        index517_ = default__.safeIndex(4, (p2).length(0))
        (p2)[index517_] = ((self).f8) - (p1)
        d_3697_v0_: _dafny.Seq
        d_3697_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vq"))
        index518_ = default__.safeIndex(4, (p2).length(0))
        (p2)[index518_] = len(d_3697_v0_)
        d_3698_v1_: bool
        d_3698_v1_ = False
        d_3698_v1_ = not(p3)
        d_3699_v2_: _dafny.Array
        nw603_ = _dafny.Array(False, 8)
        d_3699_v2_ = nw603_
        index519_ = default__.safeIndex(679, (d_3699_v2_).length(0))
        (d_3699_v2_)[index519_] = not(not (p3) or (not(p3)))
        index520_ = default__.safeIndex(679, (d_3699_v2_).length(0))
        (d_3699_v2_)[index520_] = d_3698_v1_
        (globalState).f3 = (0) - (p1)
        index521_ = default__.safeIndex(679, (d_3699_v2_).length(0))
        (d_3699_v2_)[index521_] = p0
        d_3700_v3_: _dafny.Seq
        d_3700_v3_ = _dafny.SeqWithoutIsStrInference([True, p3])
        d_3701_v4_: _dafny.Seq
        d_3701_v4_ = d_3697_v0_
        hi14_ = (self).f8
        for d_3702_i0_ in range((len(d_3700_v3_)) - (len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftg")): d_3701_v4_}))), hi14_):
            d_3703_v5_: _dafny.Set
            d_3703_v5_ = _dafny.Set({p3})
            d_3704_v6_: _dafny.Seq
            d_3704_v6_ = _dafny.SeqWithoutIsStrInference([d_3703_v5_, d_3703_v5_, d_3703_v5_, _dafny.Set({p0})])
            d_3705_v7_: _dafny.Map
            d_3705_v7_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkyu"))): (0) - ((p2)[default__.safeIndex(4, (p2).length(0))])})
            d_3706_v8_: D2
            d_3706_v8_ = D2_DC4(p1)
            d_3707_v9_: _dafny.Map
            d_3707_v9_ = _dafny.Map({d_3705_v7_: _dafny.Map({default__.fm0((d_3706_v8_).cf3, d_3702_i0_, globalState): d_3702_i0_})})
            d_3708_v10_: str
            d_3708_v10_ = _dafny.CodePoint('y')
            d_3709_v11_: _dafny.Seq
            d_3709_v11_ = _dafny.SeqWithoutIsStrInference([(self).f8, (0) - ((0) - ((self).f8)), (self).f8])
            d_3710_v12_: _dafny.Seq
            d_3710_v12_ = _dafny.SeqWithoutIsStrInference([d_3709_v11_, d_3709_v11_, default__.fm44((d_3699_v2_)[default__.safeIndex(679, (d_3699_v2_).length(0))], (d_3706_v8_).cf3, p3, globalState), _dafny.SeqWithoutIsStrInference([(p2)[default__.safeIndex(4, (p2).length(0))] for d_3711_i1_ in range(default__.abs(-320))]), d_3709_v11_])
            d_3712_v13_: C10
            nw604_ = C10()
            nw604_.ctor__(len((d_3704_v6_)[default__.safeIndex((self).f8, len(d_3704_v6_))]), d_3702_i0_, d_3707_v9_, (d_3697_v0_).set(default__.safeIndex((self).f8, len(d_3697_v0_)), d_3708_v10_), d_3710_v12_)
            d_3712_v13_ = nw604_
            index522_ = default__.safeIndex(4, (p2).length(0))
            (p2)[index522_] = len(_dafny.SeqWithoutIsStrInference([d_3702_i0_]))
            if (len(d_3697_v0_)) >= (((self).f8) + ((p2)[default__.safeIndex(4, (p2).length(0))])):
                d_3698_v1_ = (not(not(d_3698_v1_)) if (d_3699_v2_)[default__.safeIndex(679, (d_3699_v2_).length(0))] else default__.fm22(d_3698_v1_, p1, globalState))
                d_3697_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "svtrbss"))
                index523_ = default__.safeIndex(679, (d_3699_v2_).length(0))
                (d_3699_v2_)[index523_] = (default__.fm2(d_3697_v0_, 361, globalState)) != (_dafny.SeqWithoutIsStrInference([not(d_3698_v1_), p3]))
                d_3713_v14_: _dafny.Array
                nw605_ = _dafny.Array(None, 17)
                nw605_[int(0)] = d_3709_v11_
                nw605_[int(1)] = d_3709_v11_
                nw605_[int(2)] = d_3709_v11_
                nw605_[int(3)] = d_3709_v11_
                nw605_[int(4)] = d_3709_v11_
                nw605_[int(5)] = d_3709_v11_
                nw605_[int(6)] = d_3709_v11_
                nw605_[int(7)] = _dafny.SeqWithoutIsStrInference([254])
                nw605_[int(8)] = d_3709_v11_
                nw605_[int(9)] = d_3709_v11_
                nw605_[int(10)] = d_3709_v11_
                nw605_[int(11)] = default__.fm44(d_3698_v1_, (self).f8, d_3698_v1_, globalState)
                nw605_[int(12)] = d_3709_v11_
                nw605_[int(13)] = _dafny.SeqWithoutIsStrInference([(p2)[default__.safeIndex(4, (p2).length(0))] for d_3714_i2_ in range(default__.abs(95))])
                nw605_[int(14)] = _dafny.SeqWithoutIsStrInference([d_3712_v13_.f13])
                nw605_[int(15)] = d_3709_v11_
                nw605_[int(16)] = d_3709_v11_
                d_3713_v14_ = nw605_
                d_3715_v15_: D22
                d_3715_v15_ = D22_DC53(d_3713_v14_, (self).f8)
                (globalState).f3 = (d_3715_v15_).cf79
                d_3716_v16_: _dafny.Map
                d_3716_v16_ = _dafny.Map({p3: d_3702_i0_})
                d_3717_v17_: _dafny.MultiSet
                d_3717_v17_ = _dafny.MultiSet([len(d_3716_v16_)])
                d_3718_v18_: _dafny.Set
                d_3718_v18_ = _dafny.Set({925, (0) - (d_3712_v13_.f13), d_3702_i0_, (self).f8})
                index524_ = default__.safeIndex(679, (d_3699_v2_).length(0))
                (d_3699_v2_)[index524_] = ((default__.fm16((p2)[default__.safeIndex(4, (p2).length(0))], d_3698_v1_, (d_3717_v17_).cardinality, globalState) if (d_3699_v2_)[default__.safeIndex(679, (d_3699_v2_).length(0))] else d_3718_v18_)).isdisjoint(d_3718_v18_)
            elif True:
                d_3719_v19_: C2
                nw606_ = C2()
                nw606_.ctor__((p2)[default__.safeIndex(4, (p2).length(0))], d_3707_v9_, d_3697_v0_, d_3710_v12_)
                d_3719_v19_ = nw606_
                d_3720_v20_: _dafny.Array
                nw607_ = _dafny.Array(None, 10)
                nw607_[int(0)] = (d_3708_v10_ if d_3698_v1_ else _dafny.CodePoint('u'))
                nw607_[int(1)] = d_3708_v10_
                nw607_[int(2)] = d_3708_v10_
                nw607_[int(3)] = d_3708_v10_
                nw607_[int(4)] = d_3708_v10_
                nw607_[int(5)] = d_3708_v10_
                nw607_[int(6)] = d_3708_v10_
                nw607_[int(7)] = d_3708_v10_
                nw607_[int(8)] = d_3708_v10_
                nw607_[int(9)] = d_3708_v10_
                d_3720_v20_ = nw607_
                index525_ = default__.safeIndex(18, (d_3720_v20_).length(0))
                (d_3720_v20_)[index525_] = d_3708_v10_
                index526_ = default__.safeIndex(18, (d_3720_v20_).length(0))
                index527_ = default__.safeIndex(4, (p2).length(0))
                rhs532_ = d_3708_v10_
                rhs533_ = d_3712_v13_.f13
                rhs534_ = default__.safeDivisionInt((p2)[default__.safeIndex(4, (p2).length(0))], (d_3702_i0_) * ((self).f8))
                rhs535_ = ((self).f8) * ((self).f8)
                lhs344_ = d_3720_v20_
                lhs345_ = default__.safeIndex(18, (d_3720_v20_).length(0))
                lhs346_ = d_3712_v13_
                lhs347_ = p2
                lhs348_ = default__.safeIndex(4, (p2).length(0))
                lhs349_ = globalState
                lhs344_[lhs345_] = rhs532_
                lhs346_.f13 = rhs533_
                lhs347_[lhs348_] = rhs534_
                lhs349_.f3 = rhs535_
                d_3721_v21_: C9
                nw608_ = C9()
                nw608_.ctor__(d_3699_v2_)
                d_3721_v21_ = nw608_
                d_3697_v0_ = d_3697_v0_
                index528_ = default__.safeIndex(679, (d_3699_v2_).length(0))
                (d_3699_v2_)[index528_] = d_3698_v1_
            d_3722_v22_: _dafny.Map
            d_3722_v22_ = _dafny.Map({d_3698_v1_: not(p3)})
            d_3723_v23_: D18
            d_3723_v23_ = D18_DC43((d_3700_v3_)[default__.safeIndex(len((d_3697_v0_).set(default__.safeIndex(178, len(d_3697_v0_)), d_3708_v10_)), len(d_3700_v3_))], d_3702_i0_)
            index529_ = default__.safeIndex(679, (d_3699_v2_).length(0))
            (d_3699_v2_)[index529_] = ((d_3722_v22_)[(d_3723_v23_).cf60] if ((d_3723_v23_).cf60) in (d_3722_v22_) else d_3698_v1_)
        r0 = ((self.f7) | (self.f7)) | (self.f7)
        return r0

    @property
    def f8(self):
        return self._f8
