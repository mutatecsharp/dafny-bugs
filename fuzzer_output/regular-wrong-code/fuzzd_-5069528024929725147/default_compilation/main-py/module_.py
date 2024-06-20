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
    def fm3(p0, p1, p2, globalState):
        return _dafny.Map({not ((D5_DC13(False, len(_dafny.Map({True: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))}))})), _dafny.SeqWithoutIsStrInference([D3_DC9(), D3_DC9(), D3_DC9()]))).cf23) or (True): (319) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "enyymaag"))))})

    @staticmethod
    def fm4(p0, globalState):
        return default__.safeDivisionInt(997, default__.safeDivisionInt((D11_DC26(True, 212)).cf49, len(_dafny.SeqWithoutIsStrInference([-486 for d_0_i0_ in range(default__.abs(897))]))))

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(942, 311):
                d_1_v0_: int = compr_0_
                if ((942) <= (d_1_v0_)) and ((d_1_v0_) < (311)):
                    coll0_ = coll0_.union(_dafny.Set([(d_1_v0_) - ((0) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pay")))])).cardinality))]))
            return _dafny.Set(coll0_)
        return _dafny.Map({_dafny.Map({len(iife0_()
        ): -559}): (_dafny.Set({True})).ispropersubset(_dafny.Set({True, True}))})

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        return ((len(_dafny.SeqWithoutIsStrInference([not(False)]))) + ((0) - ((_dafny.MultiSet([857, (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-21: False}))])))])).cardinality))) - (791)

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        if (_dafny.SeqWithoutIsStrInference([True])) == (_dafny.SeqWithoutIsStrInference([True, True])):
            return _dafny.Map({_dafny.Map({-786: -348}): D0_DC0(False)})
        elif True:
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgxwf")))])).Elements:
                    d_2_v0_: int = compr_1_
                    if (d_2_v0_) in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgxwf")))])):
                        coll1_[(d_2_v0_) + (len(_dafny.Map({False: D1_DC3(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_3_i0_ in range(default__.abs(-941))])), len(_dafny.Map({True: False}))]))})))] = -24
                return _dafny.Map(coll1_)
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in _dafny.IntegerRange(864, 366):
                    d_4_v1_: int = compr_2_
                    if ((864) <= (d_4_v1_)) and ((d_4_v1_) < (366)):
                        coll2_[(d_4_v1_) - (len(_dafny.Map({490: False})))] = len(_dafny.SeqWithoutIsStrInference([not(True), False, True]))
                return _dafny.Map(coll2_)
            return (_dafny.Map({iife1_()
            : D0_DC0(True)})) | (_dafny.Map({iife2_()
            : D0_DC0(False)}))

    @staticmethod
    def fm9(p0, globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_5_i0_ in range(default__.abs(-995))])): (_dafny.Set({True})).isdisjoint(_dafny.Set({False, True}))})

    @staticmethod
    def fm10(globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csqbv"))

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(939, 8):
                d_6_v0_: int = compr_3_
                if ((939) <= (d_6_v0_)) and ((d_6_v0_) < (8)):
                    coll3_[(d_6_v0_) + (len(_dafny.Map({_dafny.CodePoint('a'): True})))] = False
            return _dafny.Map(coll3_)
        def iife4_():
            def iife7_():
                coll7_ = _dafny.Map()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(516, 408):
                    d_7_v1_: int = compr_7_
                    if ((516) <= (d_7_v1_)) and ((d_7_v1_) < (408)):
                        coll7_[(d_7_v1_) + (531)] = False
                return _dafny.Map(coll7_)
            def iife8_():
                coll8_ = _dafny.Map()
                compr_8_: int
                for compr_8_ in (_dafny.SeqWithoutIsStrInference([-995 for d_8_i0_ in range(default__.abs(617))])).Elements:
                    d_9_v2_: int = compr_8_
                    if (d_9_v2_) in (_dafny.SeqWithoutIsStrInference([-995 for d_8_i0_ in range(default__.abs(617))])):
                        coll8_[default__.safeModuloInt(d_9_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agsobcif"))))] = not(False)
                return _dafny.Map(coll8_)
            coll4_ = _dafny.Set()
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(516, 408):
                    d_7_v1_: int = compr_5_
                    if ((516) <= (d_7_v1_)) and ((d_7_v1_) < (408)):
                        coll5_[(d_7_v1_) + (531)] = False
                return _dafny.Map(coll5_)
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: int
                for compr_6_ in (_dafny.SeqWithoutIsStrInference([-995 for d_8_i0_ in range(default__.abs(617))])).Elements:
                    d_9_v2_: int = compr_6_
                    if (d_9_v2_) in (_dafny.SeqWithoutIsStrInference([-995 for d_8_i0_ in range(default__.abs(617))])):
                        coll6_[default__.safeModuloInt(d_9_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agsobcif"))))] = not(False)
                return _dafny.Map(coll6_)
            compr_4_: _dafny.Map
            for compr_4_ in (_dafny.SeqWithoutIsStrInference([iife5_()
            , _dafny.Map({len(_dafny.Set({True, False})): True}), iife6_()
            ])).Elements:
                d_10_v3_: _dafny.Map = compr_4_
                if (d_10_v3_) in (_dafny.SeqWithoutIsStrInference([iife7_()
                , _dafny.Map({len(_dafny.Set({True, False})): True}), iife8_()
                ])):
                    coll4_ = coll4_.union(_dafny.Set([d_10_v3_]))
            return _dafny.Set(coll4_)
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: _dafny.Map
            for compr_9_ in (_dafny.MultiSet([_dafny.Map({899: False})])).Elements:
                d_11_v4_: _dafny.Map = compr_9_
                if (d_11_v4_) in (_dafny.MultiSet([_dafny.Map({899: False})])):
                    coll9_ = coll9_.union(_dafny.Set([d_11_v4_]))
            return _dafny.Set(coll9_)
        return ((_dafny.Set({iife3_()
        , _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymwukergo")))): False}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aax"))): True}), _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahojjj")))): False})})) | (iife4_()
        )) | (iife9_()
        )

    @staticmethod
    def fm12(p0, globalState):
        return _dafny.CodePoint('j')

    @staticmethod
    def fm13(globalState):
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: _dafny.Seq
            for compr_10_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([189]))]): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_12_i0_ in range(default__.abs(744))])})).keys.Elements:
                d_13_v0_: _dafny.Seq = compr_10_
                if (d_13_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([189]))]): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_12_i0_ in range(default__.abs(744))])})):
                    coll10_[d_13_v0_] = False
            return _dafny.Map(coll10_)
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: _dafny.Seq
            for compr_11_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogphrbwfb"))), (_dafny.MultiSet([407, (0) - (len(_dafny.Set({True})))])).cardinality])])).Elements:
                d_14_v1_: _dafny.Seq = compr_11_
                if (d_14_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogphrbwfb"))), (_dafny.MultiSet([407, (0) - (len(_dafny.Set({True})))])).cardinality])])):
                    coll11_[d_14_v1_] = not(True)
            return _dafny.Map(coll11_)
        return (iife10_()
        ) != (iife11_()
        )

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([D3_DC9(), D3_DC9(), D3_DC9(), D3_DC9(), D3_DC9()])

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        return D2_DC5(_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([True]))))

    @staticmethod
    def fm16(globalState):
        return ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False]))) + ((_dafny.SeqWithoutIsStrInference([not(True)])) + (_dafny.SeqWithoutIsStrInference([not(True)])))

    @staticmethod
    def fm17(globalState):
        return (_dafny.SeqWithoutIsStrInference([not(not(True)), True])) + (_dafny.SeqWithoutIsStrInference([False, True]))

    @staticmethod
    def fm18(p0, p1, p2, p3, globalState):
        return D3_DC10(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('l')]), True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlnbv")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_15_i0_ in range(default__.abs(616))]), default__.safeModuloInt(771, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vojrh"))))))

    @staticmethod
    def fm21(p0, globalState):
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: int
            for compr_12_ in (_dafny.SeqWithoutIsStrInference([(0) - ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([not(False), False])))))])).Elements:
                d_16_v0_: int = compr_12_
                if (d_16_v0_) in (_dafny.SeqWithoutIsStrInference([(0) - ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([not(False), False])))))])):
                    coll12_[(d_16_v0_) * (492)] = False
            return _dafny.Map(coll12_)
        return (_dafny.Map({-52: not(True)})) | ((iife12_()
        ) | (_dafny.Map({-583: not(False)})))

    @staticmethod
    def fm22(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfd"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkkiujvqb")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhqapw")))

    @staticmethod
    def fm23(p0, p1, p2, globalState):
        return D2_DC5((_dafny.MultiSet([False])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True]))))

    @staticmethod
    def fm24(p0, p1, globalState):
        def iife13_():
            coll13_ = _dafny.Set()
            compr_13_: _dafny.Seq
            for compr_13_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([220]), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality, 148])})).Elements:
                d_17_v0_: _dafny.Seq = compr_13_
                if (d_17_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([220]), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality, 148])})):
                    coll13_ = coll13_.union(_dafny.Set([d_17_v0_]))
            return _dafny.Set(coll13_)
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-244]))) - (_dafny.MultiSet([len(iife13_()
        )]))) - ((_dafny.MultiSet([len((_dafny.SeqWithoutIsStrInference([False]))), 207])).intersection(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True]))])).cardinality, len(_dafny.Map({True: 872}))])), len(_dafny.SeqWithoutIsStrInference([226 for d_18_i0_ in range(default__.abs(177))]))])))

    @staticmethod
    def fm29(globalState):
        return (_dafny.MultiSet([not(not(True)), True, True, not(True), not(True)])) - ((_dafny.MultiSet([True])) | (_dafny.MultiSet([False])))

    @staticmethod
    def fm30(p0, globalState):
        return _dafny.Map({len((_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([not(False), not(True), False]))): len(_dafny.Set({731}))})

    @staticmethod
    def fm31(globalState):
        return D3_DC9()

    @staticmethod
    def fm32(globalState):
        def iife14_():
            def iife15_():
                coll15_ = _dafny.Set()
                compr_15_: int
                for compr_15_ in _dafny.IntegerRange(-54, 615):
                    d_21_v1_: int = compr_15_
                    if ((-54) <= (d_21_v1_)) and ((d_21_v1_) < (615)):
                        coll15_ = coll15_.union(_dafny.Set([(d_21_v1_) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrmfs")))])))]))
                return _dafny.Set(coll15_)
            coll14_ = _dafny.Map()
            compr_14_: int
            for compr_14_ in _dafny.IntegerRange(215, 835):
                d_19_v0_: int = compr_14_
                if ((215) <= (d_19_v0_)) and ((d_19_v0_) < (835)):
                    coll14_[(d_19_v0_) * (len(iife15_()
                    ))] = len(_dafny.Set({len(_dafny.Map({False: (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gavh"))) for d_20_i0_ in range(default__.abs(778))])))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hdtu"))), (_dafny.MultiSet([len(_dafny.Map({_dafny.CodePoint('d'): True}))])).cardinality}))
            return _dafny.Map(coll14_)
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([False, not(True), True, False]): _dafny.Map({-370: (_dafny.MultiSet([True])).cardinality})})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([not(False)]): iife14_()
        }))

    @staticmethod
    def fm33(p0, p1, globalState):
        if (_dafny.SeqWithoutIsStrInference([False])) != (_dafny.SeqWithoutIsStrInference([True, True])):
            return _dafny.CodePoint('t')
        elif True:
            return _dafny.CodePoint('v')

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvnpdwre"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqmklfrag"))))

    @staticmethod
    def fm35(p0, p1, p2, globalState):
        source0_ = D5_DC14(False, 49, len(_dafny.SeqWithoutIsStrInference([990])))
        if source0_.is_DC13:
            d_22___mcc_h0_ = source0_.cf23
            d_23___mcc_h1_ = source0_.cf24
            d_24___mcc_h2_ = source0_.cf25
            d_25_cf25_ = d_24___mcc_h2_
            d_26_cf24_ = d_23___mcc_h1_
            d_27_cf23_ = d_22___mcc_h0_
            return d_25_cf25_
        elif source0_.is_DC14:
            d_28___mcc_h3_ = source0_.cf26
            d_29___mcc_h4_ = source0_.cf27
            d_30___mcc_h5_ = source0_.cf28
            d_31_cf28_ = d_30___mcc_h5_
            d_32_cf27_ = d_29___mcc_h4_
            d_33_cf26_ = d_28___mcc_h3_
            return _dafny.SeqWithoutIsStrInference([D3_DC9()])
        elif True:
            d_34___mcc_h6_ = source0_.cf22
            d_35_cf22_ = d_34___mcc_h6_
            if False:
                return _dafny.SeqWithoutIsStrInference([D3_DC9(), D3_DC9()])
            elif True:
                return _dafny.SeqWithoutIsStrInference([D3_DC9(), D3_DC9(), D3_DC9()])

    @staticmethod
    def fm36(globalState):
        if (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m')]))) in (_dafny.Set({len(_dafny.Map({673: _dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwaejs")))]))), 213, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([False]))})) for d_36_i0_ in range(default__.abs(-927))]))})})), 732, len(_dafny.Set({True}))})):
            return _dafny.Map({515: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_37_i1_ in range(default__.abs(51))])})
        elif True:
            return (_dafny.Map({(0) - ((_dafny.MultiSet([not(False), True])).cardinality): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))})) | (_dafny.Map({len(_dafny.Set({False, False})): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ay"))}))

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        return (_dafny.Set({False})) | ((_dafny.Set({False, False})) - (_dafny.Set({(D5_DC13(not(not(False)), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_38_i0_ in range(default__.abs(546))])), _dafny.SeqWithoutIsStrInference([D3_DC9()]))).cf23, False})))

    @staticmethod
    def fm38(p0, p1, globalState):
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: int
            for compr_16_ in (_dafny.MultiSet([(_dafny.MultiSet([_dafny.CodePoint('s')])).cardinality])).Elements:
                d_39_v0_: int = compr_16_
                if (d_39_v0_) in (_dafny.MultiSet([(_dafny.MultiSet([_dafny.CodePoint('s')])).cardinality])):
                    coll16_[(d_39_v0_) * ((_dafny.MultiSet([_dafny.CodePoint('q')])).cardinality)] = not(not(not(False)))
            return _dafny.Map(coll16_)
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: int
            for compr_17_ in _dafny.IntegerRange(715, 203):
                d_40_v1_: int = compr_17_
                if ((715) <= (d_40_v1_)) and ((d_40_v1_) < (203)):
                    coll17_[default__.safeDivisionInt(d_40_v1_, -514)] = True
            return _dafny.Map(coll17_)
        return (iife16_()
        ) | ((iife17_()
         if False else _dafny.Map({(_dafny.MultiSet([_dafny.CodePoint('p'), _dafny.CodePoint('m')])).cardinality: True})))

    @staticmethod
    def fm39(p0, globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference([True])) < (_dafny.SeqWithoutIsStrInference([False, False])): _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e')])): not(not(True))})})

    @staticmethod
    def fm40(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dce")))) for d_41_i0_ in range(default__.abs(524))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oeo"))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_42_i1_ in range(default__.abs(848))]))])), 157]))

    @staticmethod
    def fm41(p0, p1, p2, p3, globalState):
        return D5_DC13(not(not (not(not(True))) or (False)), default__.safeModuloInt(556, 662), _dafny.SeqWithoutIsStrInference([D3_DC9()]))

    @staticmethod
    def fm42(p0, p1, p2, globalState):
        def iife18_():
            coll18_ = _dafny.Set()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(604, 271):
                d_43_v0_: int = compr_18_
                if ((604) <= (d_43_v0_)) and ((d_43_v0_) < (271)):
                    coll18_ = coll18_.union(_dafny.Set([default__.safeModuloInt(d_43_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "po"))))]))
            return _dafny.Set(coll18_)
        return (iife18_()
        ).intersection((_dafny.Set({len(_dafny.Set({not(not(False))}))})) | (_dafny.Set({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, True])) for d_44_i0_ in range(default__.abs(946))])): False})), len(_dafny.Set({True, True}))})))

    @staticmethod
    def fm43(globalState):
        return ((_dafny.Map({False: len(_dafny.Map({False: True}))})) | (_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_45_i0_ in range(default__.abs(72))]))}))) | (_dafny.Map({True: 993}))

    @staticmethod
    def fm44(p0, globalState):
        return D5_DC14((_dafny.MultiSet([len(_dafny.Set({False}))])).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwcbpk")))]))), (len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_46_i0_ in range(default__.abs(518))])}))) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evt")))), (len(_dafny.Map({len(_dafny.Map({not(False): False})): not(False)}))) - (305))

    @staticmethod
    def fm45(p0, p1, p2, globalState):
        return _dafny.Map({_dafny.CodePoint('m'): (len(_dafny.Map({True: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False, not(True)]))).cardinality}))) != (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgmkk"))))})

    @staticmethod
    def fm46(globalState):
        def iife19_():
            coll19_ = _dafny.Set()
            compr_19_: int
            for compr_19_ in _dafny.IntegerRange(-750, 898):
                d_47_v0_: int = compr_19_
                if ((-750) <= (d_47_v0_)) and ((d_47_v0_) < (898)):
                    coll19_ = coll19_.union(_dafny.Set([(d_47_v0_) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([D3_DC9()])])).cardinality)]))
            return _dafny.Set(coll19_)
        return ((_dafny.Map({not(True): _dafny.Set({-257, 309, len(_dafny.Set({not(False)}))})})) | (_dafny.Map({False: iife19_()
        }))) | ((_dafny.Map({False: _dafny.Set({385})})) | (_dafny.Map({False: _dafny.Set({len(_dafny.SeqWithoutIsStrInference([True, False])), (_dafny.MultiSet([len(_dafny.Map({False: True}))])).cardinality})})))

    @staticmethod
    def fm47(p0, globalState):
        return _dafny.Map({D2_DC5(_dafny.MultiSet([not(False), True, not(False), True, True])): (not(True) if True else False)})

    @staticmethod
    def Main(noArgsParameter__):
        d_48_v0_: _dafny.Array
        nw0_ = _dafny.Array(int(0), 17)
        d_48_v0_ = nw0_
        d_49_v1_: _dafny.Array
        nw1_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
        d_49_v1_ = nw1_
        d_50_v2_: int
        d_50_v2_ = 548
        d_51_v3_: bool
        d_51_v3_ = False
        d_52_v4_: _dafny.Seq
        d_52_v4_ = _dafny.SeqWithoutIsStrInference([d_51_v3_, d_51_v3_])
        d_53_v5_: _dafny.MultiSet
        d_53_v5_ = _dafny.MultiSet([(0) - (d_50_v2_), len(d_52_v4_)])
        d_54_v6_: _dafny.Array
        def lambda0_(d_55_i0_):
            return _dafny.CodePoint('c')

        init0_ = lambda0_
        nw2_ = _dafny.Array(None, 9)
        for i0_0_ in range(nw2_.length(0)):
            nw2_[i0_0_] = init0_(i0_0_)
        d_54_v6_ = nw2_
        d_56_globalState_: GlobalState
        nw3_ = GlobalState()
        nw3_.ctor__(True, True, False, True, 432, -409, 541, d_48_v0_, 766, True, _dafny.CodePoint('n'), False, -713, d_49_v1_, 804, d_53_v5_, d_54_v6_)
        d_56_globalState_ = nw3_
        d_57_v7_: _dafny.Set
        d_57_v7_ = _dafny.Set({d_50_v2_})
        d_58_v8_: _dafny.Seq
        d_58_v8_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_50_v2_}), d_57_v7_, d_57_v7_])
        d_59_v9_: _dafny.Map
        d_59_v9_ = _dafny.Map({len(d_58_v8_): len(d_57_v7_)})
        d_60_v10_: _dafny.Seq
        d_60_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvxdaq"))
        d_61_v11_: str
        d_61_v11_ = _dafny.CodePoint('v')
        d_62_v12_: _dafny.Map
        d_62_v12_ = _dafny.Map({d_59_v9_: (d_60_v10_).set(default__.safeIndex(len(d_57_v7_), len(d_60_v10_)), d_61_v11_)})
        index0_ = default__.safeIndex(749, (d_49_v1_).length(0))
        (d_49_v1_)[index0_] = ((d_62_v12_)[d_59_v9_] if (d_59_v9_) in (d_62_v12_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "angrdqnc")))
        index1_ = default__.safeIndex(749, (d_49_v1_).length(0))
        rhs0_ = d_51_v3_
        rhs1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wagfakwvy"))
        rhs2_ = d_51_v3_
        lhs0_ = d_56_globalState_
        lhs1_ = d_49_v1_
        lhs2_ = default__.safeIndex(749, (d_49_v1_).length(0))
        lhs3_ = d_56_globalState_
        lhs0_.f3 = rhs0_
        lhs1_[lhs2_] = rhs1_
        lhs3_.f9 = rhs2_
        if d_51_v3_:
            d_63_v13_: _dafny.Map
            d_63_v13_ = _dafny.Map({_dafny.CodePoint('q'): d_51_v3_})
            d_64_v14_: D1
            d_64_v14_ = D1_DC4(False, _dafny.Set({_dafny.Map({d_50_v2_: d_51_v3_})}), d_50_v2_)
            d_65_v15_: _dafny.Array
            nw4_ = _dafny.Array(None, 27)
            nw4_[int(0)] = d_51_v3_
            nw4_[int(1)] = d_51_v3_
            nw4_[int(2)] = True
            nw4_[int(3)] = False
            nw4_[int(4)] = d_51_v3_
            nw4_[int(5)] = d_51_v3_
            nw4_[int(6)] = True
            nw4_[int(7)] = (d_64_v14_).cf9
            nw4_[int(8)] = not(d_51_v3_)
            nw4_[int(9)] = not(d_51_v3_)
            nw4_[int(10)] = d_51_v3_
            nw4_[int(11)] = d_51_v3_
            nw4_[int(12)] = True
            nw4_[int(13)] = d_51_v3_
            nw4_[int(14)] = d_51_v3_
            nw4_[int(15)] = d_51_v3_
            nw4_[int(16)] = d_51_v3_
            nw4_[int(17)] = d_51_v3_
            nw4_[int(18)] = d_51_v3_
            nw4_[int(19)] = d_51_v3_
            nw4_[int(20)] = d_51_v3_
            nw4_[int(21)] = default__.fm13(d_56_globalState_)
            nw4_[int(22)] = d_51_v3_
            nw4_[int(23)] = d_51_v3_
            nw4_[int(24)] = not(not(d_51_v3_))
            nw4_[int(25)] = d_51_v3_
            nw4_[int(26)] = d_51_v3_
            d_65_v15_ = nw4_
            d_66_v16_: _dafny.Set
            d_66_v16_ = _dafny.Set({d_65_v15_})
            d_67_v17_: C5
            nw5_ = C5()
            nw5_.ctor__(_dafny.SeqWithoutIsStrInference([((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))])[default__.safeIndex(len(_dafny.Map({d_50_v2_: d_51_v3_})), len((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))]))] for d_68_i1_ in range(default__.abs(446))]), d_51_v3_, default__.safeModuloInt(969, (0) - (d_50_v2_)), d_63_v13_, d_65_v15_, (d_66_v16_) - (d_66_v16_), d_51_v3_)
            d_67_v17_ = nw5_
            index2_ = default__.safeIndex(676, (d_48_v0_).length(0))
            (d_48_v0_)[index2_] = 884
            index3_ = default__.safeIndex(676, (d_48_v0_).length(0))
            (d_48_v0_)[index3_] = 7
            d_69_v18_: D11
            d_69_v18_ = D11_DC26(d_51_v3_, len(_dafny.Set({(d_48_v0_)[default__.safeIndex(676, (d_48_v0_).length(0))]})))
            d_70_v19_: D11
            d_70_v19_ = D11_DC28(d_69_v18_)
            d_71_v20_: D11
            d_71_v20_ = D11_DC28(d_70_v19_)
            d_72_v21_: D11
            d_72_v21_ = D11_DC28(d_71_v20_)
            pat_let_tv0_ = d_70_v19_
            def iife20_(_pat_let0_0):
                def iife21_(d_73_dt__update__tmp_h0_):
                    def iife22_(_pat_let1_0):
                        def iife23_(d_74_dt__update_hcf53_h0_):
                            return D11_DC28(d_74_dt__update_hcf53_h0_)
                        return iife23_(_pat_let1_0)
                    return iife22_(pat_let_tv0_)
                return iife21_(_pat_let0_0)
            d_72_v21_ = iife20_(d_72_v21_)
            if d_51_v3_:
                d_75_v22_: _dafny.Set
                d_75_v22_ = _dafny.Set({(d_67_v17_).f30})
                index4_ = default__.safeIndex(676, (d_48_v0_).length(0))
                (d_48_v0_)[index4_] = (0) - (len(((d_75_v22_) - (d_75_v22_)) - ((d_75_v22_) | (d_75_v22_))))
                d_60_v10_ = _dafny.SeqWithoutIsStrInference([d_61_v11_ for d_76_i2_ in range(default__.abs(803))])
                d_77_v23_: _dafny.Map
                d_77_v23_ = _dafny.Map({(d_57_v7_) - (d_57_v7_): (d_52_v4_ if d_51_v3_ else d_52_v4_)})
                d_77_v23_ = (d_77_v23_).set((d_57_v7_) - (d_57_v7_), _dafny.SeqWithoutIsStrInference([(d_67_v17_).f30, (d_67_v17_).f30, d_51_v3_, not(d_51_v3_), not(not((d_67_v17_).f30))]))
                d_78_v24_: D5
                d_78_v24_ = D5_DC14((d_51_v3_ if (d_67_v17_).f30 else (d_67_v17_).f30), default__.safeDivisionInt(467, d_50_v2_), (d_48_v0_)[default__.safeIndex(676, (d_48_v0_).length(0))])
                d_78_v24_ = d_78_v24_
                d_79_v25_: _dafny.Seq
                d_79_v25_ = _dafny.SeqWithoutIsStrInference([(d_48_v0_)[default__.safeIndex(676, (d_48_v0_).length(0))]])
                d_80_v26_: _dafny.MultiSet
                d_80_v26_ = _dafny.MultiSet([d_61_v11_])
                (d_56_globalState_).f14 = (d_79_v25_)[default__.safeIndex(len(((d_60_v10_) + ((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))])).set(default__.safeIndex((d_80_v26_).cardinality, len((d_60_v10_) + ((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))]))), d_61_v11_)), len(d_79_v25_))]
            elif True:
                d_81_v27_: C9
                nw6_ = C9()
                nw6_.ctor__((d_53_v5_).cardinality, d_50_v2_, _dafny.Map({d_61_v11_: (d_67_v17_).f30}), d_65_v15_, d_66_v16_, (d_67_v17_).f30)
                d_81_v27_ = nw6_
                d_82_v28_: _dafny.Map
                d_82_v28_ = _dafny.Map({d_81_v27_: (0) - (d_81_v27_.f22)})
                d_83_v29_: bool
                d_84_v30_: bool
                d_85_v31_: int
                out0_: bool
                out1_: bool
                out2_: int
                out0_, out1_, out2_ = (d_67_v17_).m0((((d_82_v28_)[d_81_v27_] if (d_81_v27_) in (d_82_v28_) else (d_48_v0_)[default__.safeIndex(676, (d_48_v0_).length(0))]) if d_51_v3_ else d_50_v2_), d_51_v3_, d_81_v27_.f22, d_56_globalState_)
                d_83_v29_ = out0_
                d_84_v30_ = out1_
                d_85_v31_ = out2_
                d_86_v32_: _dafny.Map
                d_86_v32_ = _dafny.Map({(d_67_v17_).f30: default__.fm13(d_56_globalState_)})
                d_87_v33_: _dafny.Map
                d_87_v33_ = _dafny.Map({d_59_v9_: d_85_v31_})
                d_88_v34_: _dafny.Set
                d_88_v34_ = _dafny.Set({d_84_v30_})
                d_89_v35_: D7
                d_89_v35_ = D7_DC16(d_88_v34_)
                d_86_v32_ = (d_86_v32_).set((d_81_v27_).fm1(d_87_v33_, d_56_globalState_), ((d_89_v35_).cf30).ispropersubset(d_88_v34_))
                d_90_v36_: C2
                nw7_ = C2()
                nw7_.ctor__((d_66_v16_) | (d_66_v16_), True)
                d_90_v36_ = nw7_
                d_91_v37_: _dafny.Seq
                d_91_v37_ = _dafny.SeqWithoutIsStrInference([len(d_88_v34_)])
                d_92_v38_: _dafny.Seq
                d_92_v38_ = _dafny.SeqWithoutIsStrInference([d_91_v37_, d_91_v37_, d_91_v37_, d_91_v37_])
                d_93_v39_: _dafny.Array
                nw8_ = _dafny.Array(int(0), 19)
                d_93_v39_ = nw8_
                d_94_v40_: D0
                d_94_v40_ = D0_DC2(d_84_v30_, d_50_v2_, d_93_v39_)
                d_95_v41_: _dafny.Map
                d_95_v41_ = _dafny.Map({d_83_v29_: _dafny.CodePoint('x')})
                d_96_v42_: _dafny.Array
                nw9_ = _dafny.Array(None, 25)
                nw9_[int(0)] = d_91_v37_
                nw9_[int(1)] = d_91_v37_
                nw9_[int(2)] = _dafny.SeqWithoutIsStrInference([d_81_v27_.f22, d_81_v27_.f22])
                nw9_[int(3)] = default__.fm40(not(d_84_v30_), d_84_v30_, d_56_globalState_)
                nw9_[int(4)] = d_91_v37_
                nw9_[int(5)] = d_91_v37_
                nw9_[int(6)] = d_91_v37_
                nw9_[int(7)] = (d_91_v37_) + (d_91_v37_)
                nw9_[int(8)] = (d_92_v38_)[default__.safeIndex(d_81_v27_.f22, len(d_92_v38_))]
                nw9_[int(9)] = (d_91_v37_ if d_83_v29_ else d_91_v37_)
                nw9_[int(10)] = d_91_v37_
                nw9_[int(11)] = d_91_v37_
                nw9_[int(12)] = d_91_v37_
                nw9_[int(13)] = d_91_v37_
                nw9_[int(14)] = _dafny.SeqWithoutIsStrInference([d_50_v2_, d_50_v2_, d_50_v2_])
                nw9_[int(15)] = d_91_v37_
                nw9_[int(16)] = (_dafny.SeqWithoutIsStrInference([(0) - (d_85_v31_) for d_97_i3_ in range(default__.abs(133))])) + (d_91_v37_)
                nw9_[int(17)] = d_91_v37_
                nw9_[int(18)] = d_91_v37_
                nw9_[int(19)] = _dafny.SeqWithoutIsStrInference([len(d_91_v37_) for d_98_i4_ in range(default__.abs(-558))])
                nw9_[int(20)] = d_91_v37_
                nw9_[int(21)] = (d_91_v37_) + (_dafny.SeqWithoutIsStrInference([(d_94_v40_).cf6]))
                nw9_[int(22)] = _dafny.SeqWithoutIsStrInference([d_81_v27_.f22 for d_99_i5_ in range(default__.abs(558))])
                nw9_[int(23)] = d_91_v37_
                nw9_[int(24)] = (_dafny.SeqWithoutIsStrInference([d_50_v2_, len(d_95_v41_)])) + (_dafny.SeqWithoutIsStrInference([d_81_v27_.f22]))
                d_96_v42_ = nw9_
                d_96_v42_ = d_96_v42_
                d_100_v45_: _dafny.Array
                nw10_ = _dafny.Array(None, 26)
                nw10_[int(0)] = d_59_v9_
                def iife24_():
                    coll20_ = _dafny.Map()
                    compr_20_: int
                    for compr_20_ in _dafny.IntegerRange(-832, 895):
                        d_101_v43_: int = compr_20_
                        if ((-832) <= (d_101_v43_)) and ((d_101_v43_) < (895)):
                            coll20_[default__.safeDivisionInt(d_101_v43_, d_85_v31_)] = d_50_v2_
                    return _dafny.Map(coll20_)
                nw10_[int(1)] = iife24_()
                
                nw10_[int(2)] = _dafny.Map({193: d_81_v27_.f22})
                nw10_[int(3)] = d_59_v9_
                nw10_[int(4)] = d_59_v9_
                nw10_[int(5)] = d_59_v9_
                nw10_[int(6)] = d_59_v9_
                nw10_[int(7)] = d_59_v9_
                nw10_[int(8)] = d_59_v9_
                nw10_[int(9)] = d_59_v9_
                nw10_[int(10)] = d_59_v9_
                nw10_[int(11)] = d_59_v9_
                nw10_[int(12)] = d_59_v9_
                nw10_[int(13)] = d_59_v9_
                nw10_[int(14)] = d_59_v9_
                nw10_[int(15)] = d_59_v9_
                nw10_[int(16)] = d_59_v9_
                nw10_[int(17)] = d_59_v9_
                def iife25_():
                    coll21_ = _dafny.Map()
                    compr_21_: int
                    for compr_21_ in _dafny.IntegerRange(765, 4):
                        d_102_v44_: int = compr_21_
                        if ((765) <= (d_102_v44_)) and ((d_102_v44_) < (4)):
                            coll21_[default__.safeModuloInt(d_102_v44_, d_81_v27_.f22)] = len(_dafny.SeqWithoutIsStrInference([(d_48_v0_)[default__.safeIndex(676, (d_48_v0_).length(0))]]))
                    return _dafny.Map(coll21_)
                nw10_[int(18)] = iife25_()
                
                nw10_[int(19)] = d_59_v9_
                nw10_[int(20)] = d_59_v9_
                nw10_[int(21)] = d_59_v9_
                nw10_[int(22)] = d_59_v9_
                nw10_[int(23)] = d_59_v9_
                nw10_[int(24)] = d_59_v9_
                nw10_[int(25)] = d_59_v9_
                d_100_v45_ = nw10_
                d_103_v46_: D7
                d_103_v46_ = D7_DC17(d_100_v45_, d_83_v29_, (d_48_v0_)[default__.safeIndex(676, (d_48_v0_).length(0))])
                d_104_v47_: D5
                d_104_v47_ = D5_DC14((d_88_v34_).isdisjoint(d_88_v34_), (d_103_v46_).cf33, 453)
                d_105_v48_: D3
                d_105_v48_ = D3_DC9()
                d_106_v49_: _dafny.Seq
                d_106_v49_ = _dafny.SeqWithoutIsStrInference([d_105_v48_, d_105_v48_, d_105_v48_])
                d_107_v50_: _dafny.Seq
                d_107_v50_ = d_106_v49_
                index5_ = default__.safeIndex(501, (d_65_v15_).length(0))
                (d_65_v15_)[index5_] = d_84_v30_
                index6_ = default__.safeIndex(501, (d_65_v15_).length(0))
                rhs3_ = False
                rhs4_ = default__.fm44(default__.safeDivisionInt(d_50_v2_, (d_67_v17_).fm0(930, d_61_v11_, d_56_globalState_)), d_56_globalState_)
                rhs5_ = d_107_v50_
                rhs6_ = not((d_67_v17_).f30)
                lhs4_ = d_56_globalState_
                lhs5_ = d_65_v15_
                lhs6_ = default__.safeIndex(501, (d_65_v15_).length(0))
                lhs4_.f9 = rhs3_
                d_104_v47_ = rhs4_
                d_107_v50_ = rhs5_
                lhs5_[lhs6_] = rhs6_
            d_108_v51_: _dafny.Seq
            d_108_v51_ = _dafny.SeqWithoutIsStrInference([(d_48_v0_)[default__.safeIndex(676, (d_48_v0_).length(0))]])
            d_109_v52_: C1
            nw11_ = C1()
            nw11_.ctor__()
            d_109_v52_ = nw11_
            d_110_v53_: _dafny.Array
            nw12_ = _dafny.Array(None, 28)
            nw12_[int(0)] = d_109_v52_
            nw12_[int(1)] = d_109_v52_
            nw12_[int(2)] = d_109_v52_
            nw12_[int(3)] = d_109_v52_
            nw12_[int(4)] = d_109_v52_
            nw12_[int(5)] = d_109_v52_
            nw12_[int(6)] = d_109_v52_
            nw12_[int(7)] = d_109_v52_
            nw12_[int(8)] = d_109_v52_
            nw12_[int(9)] = d_109_v52_
            nw12_[int(10)] = d_109_v52_
            nw12_[int(11)] = d_109_v52_
            nw12_[int(12)] = (d_109_v52_ if d_51_v3_ else d_109_v52_)
            nw12_[int(13)] = d_109_v52_
            nw12_[int(14)] = d_109_v52_
            nw12_[int(15)] = d_109_v52_
            nw12_[int(16)] = d_109_v52_
            nw12_[int(17)] = d_109_v52_
            nw12_[int(18)] = d_109_v52_
            nw12_[int(19)] = d_109_v52_
            nw12_[int(20)] = d_109_v52_
            nw12_[int(21)] = d_109_v52_
            nw12_[int(22)] = d_109_v52_
            nw12_[int(23)] = d_109_v52_
            nw12_[int(24)] = d_109_v52_
            nw12_[int(25)] = d_109_v52_
            nw12_[int(26)] = d_109_v52_
            nw12_[int(27)] = d_109_v52_
            d_110_v53_ = nw12_
            index7_ = default__.safeIndex(209, (d_110_v53_).length(0))
            (d_110_v53_)[index7_] = d_109_v52_
            d_111_v54_: T2
            nw13_ = C8()
            nw13_.ctor__((d_53_v5_).cardinality, d_51_v3_, d_65_v15_, (d_66_v16_) | (d_66_v16_), (d_52_v4_) <= (d_52_v4_), ((d_48_v0_)[default__.safeIndex(676, (d_48_v0_).length(0))]) + (d_50_v2_), _dafny.Map({_dafny.CodePoint('l'): (d_67_v17_).f30}))
            d_111_v54_ = nw13_
            d_112_v55_: _dafny.Seq
            d_112_v55_ = _dafny.SeqWithoutIsStrInference([d_60_v10_, (d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))], _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_113_i6_ in range(default__.abs(589))]), (d_67_v17_).f29, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "geplenmxg"))])
            d_114_v56_: _dafny.Seq
            d_114_v56_ = _dafny.SeqWithoutIsStrInference([d_53_v5_])
            index8_ = default__.safeIndex(209, (d_110_v53_).length(0))
            rhs7_ = (d_67_v17_).f30
            rhs8_ = (d_111_v54_).f20
            rhs9_ = _dafny.SeqWithoutIsStrInference([len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwcwqwn"))) + ((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))])).set(default__.safeIndex(len(d_112_v55_), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwcwqwn"))) + ((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))]))), ((d_67_v17_).f29)[default__.safeIndex(((d_59_v9_)[len(d_114_v56_)] if (len(d_114_v56_)) in (d_59_v9_) else (d_111_v54_).f20), len((d_67_v17_).f29))]))])
            rhs10_ = d_109_v52_
            rhs11_ = d_111_v54_
            lhs7_ = d_56_globalState_
            lhs8_ = d_56_globalState_
            lhs9_ = d_110_v53_
            lhs10_ = default__.safeIndex(209, (d_110_v53_).length(0))
            lhs7_.f3 = rhs7_
            lhs8_.f14 = rhs8_
            d_108_v51_ = rhs9_
            lhs9_[lhs10_] = rhs10_
            d_111_v54_ = rhs11_
        elif True:
            d_115_v57_: _dafny.Array
            def lambda1_(d_116_v3_):
                def lambda2_(d_117_i7_):
                    return d_116_v3_

                return lambda2_

            init1_ = lambda1_(d_51_v3_)
            nw14_ = _dafny.Array(None, 28)
            for i0_1_ in range(nw14_.length(0)):
                nw14_[i0_1_] = init1_(i0_1_)
            d_115_v57_ = nw14_
            index9_ = default__.safeIndex(154, (d_115_v57_).length(0))
            (d_115_v57_)[index9_] = (d_51_v3_) and (d_51_v3_)
            d_118_v58_: _dafny.Set
            d_118_v58_ = _dafny.Set({d_115_v57_})
            d_119_v59_: _dafny.Map
            d_119_v59_ = _dafny.Map({d_61_v11_: d_51_v3_})
            d_120_v60_: C8
            nw15_ = C8()
            nw15_.ctor__(d_50_v2_, d_51_v3_, d_115_v57_, d_118_v58_, d_51_v3_, d_50_v2_, d_119_v59_)
            d_120_v60_ = nw15_
            d_121_v61_: _dafny.Seq
            d_121_v61_ = _dafny.SeqWithoutIsStrInference([d_120_v60_])
            d_122_v62_: _dafny.Set
            d_122_v62_ = _dafny.Set({d_121_v61_})
            index10_ = default__.safeIndex(154, (d_115_v57_).length(0))
            (d_115_v57_)[index10_] = (d_122_v62_).ispropersubset(d_122_v62_)
            if (d_52_v4_)[default__.safeIndex(d_120_v60_.f23, len(d_52_v4_))]:
                d_123_v63_: D10
                d_123_v63_ = D10_DC22(d_57_v7_)
                d_124_v64_: _dafny.Map
                d_124_v64_ = _dafny.Map({(d_115_v57_)[default__.safeIndex(154, (d_115_v57_).length(0))]: d_123_v63_})
                pat_let_tv1_ = d_57_v7_
                d_125_v65_: int
                d_126_v66_: bool
                d_127_v67_: int
                out3_: int
                out4_: bool
                out5_: int
                def iife26_(_pat_let2_0):
                    def iife27_(d_128_dt__update__tmp_h1_):
                        def iife28_(_pat_let3_0):
                            def iife29_(d_129_dt__update_hcf42_h0_):
                                return D10_DC22(d_129_dt__update_hcf42_h0_)
                            return iife29_(_pat_let3_0)
                        return iife28_(pat_let_tv1_)
                    return iife27_(_pat_let2_0)
                out3_, out4_, out5_ = (d_120_v60_).m2(len((d_124_v64_).set(d_51_v3_, iife26_(d_123_v63_))), d_56_globalState_)
                d_125_v65_ = out3_
                d_126_v66_ = out4_
                d_127_v67_ = out5_
                index11_ = default__.safeIndex(705, (d_54_v6_).length(0))
                (d_54_v6_)[index11_] = _dafny.CodePoint('o')
                index12_ = default__.safeIndex(705, (d_54_v6_).length(0))
                (d_54_v6_)[index12_] = d_61_v11_
                d_130_v68_: _dafny.Seq
                d_130_v68_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False]), _dafny.MultiSet([False, d_51_v3_, d_126_v66_, d_126_v66_, (d_120_v60_).f24])])
                d_131_v69_: _dafny.Array
                nw16_ = _dafny.Array(_dafny.Seq({}), 16)
                d_131_v69_ = nw16_
                d_132_v70_: _dafny.Map
                d_132_v70_ = _dafny.Map({d_120_v60_.f23: d_131_v69_})
                d_133_v71_: _dafny.MultiSet
                d_133_v71_ = _dafny.MultiSet([d_131_v69_, d_131_v69_, ((d_132_v70_)[d_120_v60_.f23] if (d_120_v60_.f23) in (d_132_v70_) else d_131_v69_)])
                rhs12_ = d_127_v67_
                rhs13_ = (d_130_v68_) + (d_130_v68_)
                rhs14_ = (0) - (d_127_v67_)
                rhs15_ = d_115_v57_
                rhs16_ = ((d_133_v71_)[d_131_v69_] if (d_131_v69_) in (d_133_v71_) else d_120_v60_.f23)
                lhs11_ = d_56_globalState_
                lhs12_ = d_120_v60_
                lhs11_.f14 = rhs12_
                d_130_v68_ = rhs13_
                d_127_v67_ = rhs14_
                d_115_v57_ = rhs15_
                lhs12_.f23 = rhs16_
                d_134_v72_: _dafny.Map
                d_134_v72_ = _dafny.Map({d_125_v65_: d_53_v5_})
                d_134_v72_ = (d_134_v72_) | ((d_134_v72_) | (d_134_v72_))
                d_50_v2_ = d_120_v60_.f23
            elif True:
                d_135_v73_: T2
                nw17_ = C9()
                nw17_.ctor__(d_50_v2_, d_50_v2_, d_119_v59_, d_115_v57_, d_118_v58_, (d_120_v60_).f24)
                d_135_v73_ = nw17_
                d_136_v74_: D10
                d_136_v74_ = D10_DC23((d_120_v60_).f24, (d_120_v60_).f24, d_135_v73_)
                d_137_v75_: D10
                d_137_v75_ = D10_DC24(d_136_v74_)
                d_138_v76_: _dafny.Array
                nw18_ = _dafny.Array(None, 5)
                nw18_[int(0)] = d_137_v75_
                nw18_[int(1)] = D10_DC24(d_136_v74_)
                nw18_[int(2)] = d_137_v75_
                nw18_[int(3)] = d_137_v75_
                nw18_[int(4)] = D10_DC24(d_136_v74_)
                d_138_v76_ = nw18_
                d_139_v77_: _dafny.Map
                d_139_v77_ = _dafny.Map({d_50_v2_: d_138_v76_})
                d_140_v78_: _dafny.Array
                nw19_ = _dafny.Array(None, 5)
                nw19_[int(0)] = d_138_v76_
                nw19_[int(1)] = ((d_139_v77_)[(d_135_v73_).f20] if ((d_135_v73_).f20) in (d_139_v77_) else d_138_v76_)
                nw19_[int(2)] = d_138_v76_
                nw19_[int(3)] = d_138_v76_
                nw19_[int(4)] = d_138_v76_
                d_140_v78_ = nw19_
                d_141_v79_: D14
                d_141_v79_ = D14_DC34(d_54_v6_)
                d_142_v80_: _dafny.MultiSet
                d_142_v80_ = _dafny.MultiSet([d_141_v79_])
                d_143_v81_: _dafny.Map
                d_143_v81_ = _dafny.Map({(502) > (d_50_v2_): (d_142_v80_).set(d_141_v79_, default__.abs(d_120_v60_.f23))})
                rhs17_ = d_140_v78_
                rhs18_ = (d_120_v60_).f24
                rhs19_ = d_143_v81_
                rhs20_ = 748
                lhs13_ = d_56_globalState_
                lhs14_ = d_120_v60_
                d_140_v78_ = rhs17_
                lhs13_.f9 = rhs18_
                d_143_v81_ = rhs19_
                lhs14_.f23 = rhs20_
                d_144_v82_: _dafny.MultiSet
                d_144_v82_ = _dafny.MultiSet([(d_115_v57_)[default__.safeIndex(154, (d_115_v57_).length(0))]])
                d_145_v83_: _dafny.Array
                nw20_ = _dafny.Array(None, 7)
                nw20_[int(0)] = d_59_v9_
                nw20_[int(1)] = d_59_v9_
                nw20_[int(2)] = (d_59_v9_).set(d_120_v60_.f23, len(_dafny.Set({d_51_v3_, (d_120_v60_).f24})))
                nw20_[int(3)] = d_59_v9_
                nw20_[int(4)] = default__.fm30((d_120_v60_).f24, d_56_globalState_)
                nw20_[int(5)] = d_59_v9_
                nw20_[int(6)] = d_59_v9_
                d_145_v83_ = nw20_
                d_146_v84_: D7
                d_146_v84_ = D7_DC17(d_145_v83_, d_51_v3_, d_120_v60_.f23)
                d_147_v85_: _dafny.Seq
                d_147_v85_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_61_v11_: (d_120_v60_).f24})])
                d_148_v86_: bool
                d_149_v87_: bool
                d_150_v88_: int
                out6_: bool
                out7_: bool
                out8_: int
                out6_, out7_, out8_ = (d_120_v60_).m0(((d_144_v82_)[(d_146_v84_).cf32] if ((d_146_v84_).cf32) in (d_144_v82_) else d_50_v2_), (_dafny.SeqWithoutIsStrInference([d_135_v73_.f21 for d_151_i8_ in range(default__.abs(201))])) == (d_147_v85_), d_120_v60_.f23, d_56_globalState_)
                d_148_v86_ = out6_
                d_149_v87_ = out7_
                d_150_v88_ = out8_
                d_149_v87_ = (d_135_v73_).f18
                d_152_v89_: T0
                nw21_ = C2()
                nw21_.ctor__(_dafny.Set({d_135_v73_.f19, d_135_v73_.f19, d_135_v73_.f19, d_115_v57_, d_115_v57_}), False)
                d_152_v89_ = nw21_
                d_153_v90_: C2
                nw22_ = C2()
                nw22_.ctor__(d_118_v58_, (562) == (len(_dafny.Map({len(d_60_v10_): d_152_v89_}))))
                d_153_v90_ = nw22_
                d_154_v91_: _dafny.Seq
                d_154_v91_ = _dafny.SeqWithoutIsStrInference([d_153_v90_])
                d_153_v90_ = (d_154_v91_)[default__.safeIndex((d_50_v2_) + (len(d_60_v10_)), len(d_154_v91_))]
                d_155_v92_: bool
                d_156_v93_: bool
                d_157_v94_: int
                out9_: bool
                out10_: bool
                out11_: int
                out9_, out10_, out11_ = (d_120_v60_).m0(default__.safeDivisionInt(len((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))]), (0) - ((d_135_v73_).f20)), ((d_135_v73_).f18) and ((d_120_v60_).f24), d_50_v2_, d_56_globalState_)
                d_155_v92_ = out9_
                d_156_v93_ = out10_
                d_157_v94_ = out11_
            d_158_v95_: _dafny.Map
            d_158_v95_ = _dafny.Map({d_50_v2_: (d_115_v57_)[default__.safeIndex(154, (d_115_v57_).length(0))]})
            d_159_v96_: _dafny.MultiSet
            d_159_v96_ = _dafny.MultiSet([d_60_v10_])
            d_160_v98_: _dafny.Map
            def iife30_():
                coll22_ = _dafny.Set()
                compr_22_: _dafny.Seq
                for compr_22_ in (d_159_v96_).Elements:
                    d_161_v97_: _dafny.Seq = compr_22_
                    if (d_161_v97_) in (d_159_v96_):
                        coll22_ = coll22_.union(_dafny.Set([d_161_v97_]))
                return _dafny.Set(coll22_)
            d_160_v98_ = _dafny.Map({iife30_()
            : (d_115_v57_)[default__.safeIndex(154, (d_115_v57_).length(0))]})
            d_158_v95_ = (d_158_v95_).set(-78, ((d_160_v98_)[_dafny.Set({(d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))]})] if (_dafny.Set({(d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))]})) in (d_160_v98_) else d_51_v3_))
            (d_120_v60_).f23 = default__.safeModuloInt(d_120_v60_.f23, default__.safeDivisionInt(d_50_v2_, 797))
            d_162_v99_: C1
            nw23_ = C1()
            nw23_.ctor__()
            d_162_v99_ = nw23_
        hi0_ = d_50_v2_
        for d_163_i9_ in range(d_50_v2_, hi0_):
            d_164_v100_: _dafny.Array
            def lambda3_(d_165_v3_):
                def lambda4_(d_166_i10_):
                    return d_165_v3_

                return lambda4_

            init2_ = lambda3_(d_51_v3_)
            nw24_ = _dafny.Array(None, 27)
            for i0_2_ in range(nw24_.length(0)):
                nw24_[i0_2_] = init2_(i0_2_)
            d_164_v100_ = nw24_
            d_167_v101_: _dafny.Set
            d_167_v101_ = _dafny.Set({d_164_v100_})
            d_168_v102_: _dafny.Map
            d_168_v102_ = _dafny.Map({d_61_v11_: d_51_v3_})
            d_169_v103_: C8
            nw25_ = C8()
            nw25_.ctor__(d_50_v2_, d_51_v3_, d_164_v100_, d_167_v101_, d_51_v3_, d_163_i9_, d_168_v102_)
            d_169_v103_ = nw25_
            d_170_v104_: _dafny.Map
            d_170_v104_ = _dafny.Map({d_51_v3_: d_169_v103_})
            d_171_v105_: _dafny.Seq
            d_171_v105_ = _dafny.SeqWithoutIsStrInference([16, d_169_v103_.f23])
            d_172_v107_: _dafny.Set
            d_172_v107_ = _dafny.Set({default__.fm33(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lx"))), d_51_v3_, d_56_globalState_), d_61_v11_})
            d_173_v108_: _dafny.Array
            nw26_ = _dafny.Array(None, 7)
            nw26_[int(0)] = (d_170_v104_) == (d_170_v104_)
            def iife31_():
                coll23_ = _dafny.Set()
                compr_23_: int
                for compr_23_ in (d_171_v105_).Elements:
                    d_174_v106_: int = compr_23_
                    if (d_174_v106_) in (d_171_v105_):
                        coll23_ = coll23_.union(_dafny.Set([(d_174_v106_) - (481)]))
                return _dafny.Set(coll23_)
            nw26_[int(1)] = (iife31_()
            ).issubset(d_57_v7_)
            nw26_[int(2)] = False
            nw26_[int(3)] = False
            nw26_[int(4)] = d_51_v3_
            nw26_[int(5)] = (d_169_v103_).f24
            nw26_[int(6)] = (d_172_v107_).ispropersubset(d_172_v107_)
            d_173_v108_ = nw26_
            index13_ = default__.safeIndex(618, (d_173_v108_).length(0))
            (d_173_v108_)[index13_] = (d_169_v103_).f24
            d_175_v109_: _dafny.Map
            d_175_v109_ = _dafny.Map({d_163_i9_: d_61_v11_})
            d_176_v110_: C4
            nw27_ = C4()
            nw27_.ctor__(len(d_175_v109_), -688, d_167_v101_, True)
            d_176_v110_ = nw27_
            d_177_v111_: _dafny.Seq
            d_177_v111_ = _dafny.SeqWithoutIsStrInference([d_176_v110_])
            index14_ = default__.safeIndex(618, (d_173_v108_).length(0))
            (d_173_v108_)[index14_] = not((d_177_v111_) < ((d_177_v111_) + (d_177_v111_)))
            d_178_v112_: D3
            d_178_v112_ = D3_DC10((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))], d_51_v3_, (d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))], _dafny.SeqWithoutIsStrInference([d_61_v11_ for d_179_i11_ in range(default__.abs(613))]), (d_176_v110_).f25)
            d_180_v113_: int
            out12_: int
            out12_ = (d_169_v103_).m1(-149, (_dafny.CodePoint('k')) in ((d_178_v112_).cf18), (d_163_i9_) * (d_163_i9_), d_51_v3_, d_56_globalState_)
            d_180_v113_ = out12_
            d_51_v3_ = d_51_v3_
            (d_56_globalState_).f3 = (d_169_v103_).f24
        d_52_v4_ = d_52_v4_
        d_181_i12_: int
        d_181_i12_ = 0
        with _dafny.label("0"):
            while (d_50_v2_) == (d_50_v2_):
                with _dafny.c_label("0"):
                    if (d_181_i12_) >= (100):
                        raise _dafny.Break("0")
                    d_181_i12_ = (d_181_i12_) + (1)
                    d_182_v114_: _dafny.Map
                    d_182_v114_ = _dafny.Map({d_52_v4_: d_59_v9_})
                    source1_ = D9_DC19(d_182_v114_)
                    if source1_.is_DC20:
                        d_183___mcc_h0_ = source1_.cf36
                        d_184___mcc_h1_ = source1_.cf37
                        d_185___mcc_h2_ = source1_.cf38
                        d_186___mcc_h3_ = source1_.cf39
                        d_187___mcc_h4_ = source1_.cf40
                        d_188_cf40_ = d_187___mcc_h4_
                        d_189_cf39_ = d_186___mcc_h3_
                        d_190_cf38_ = d_185___mcc_h2_
                        d_191_cf37_ = d_184___mcc_h1_
                        d_192_cf36_ = d_183___mcc_h0_
                        d_193_v115_: _dafny.Map
                        d_193_v115_ = _dafny.Map({d_190_cf38_: ((d_53_v5_)[len(d_192_cf36_)] if (len(d_192_cf36_)) in (d_53_v5_) else d_191_cf37_)})
                        d_194_v116_: D0
                        d_194_v116_ = D0_DC0(d_190_cf38_)
                        d_195_v117_: _dafny.Array
                        def lambda5_(d_196_v3_):
                            def lambda6_(d_197_i13_):
                                return d_196_v3_

                            return lambda6_

                        init3_ = lambda5_(d_51_v3_)
                        nw28_ = _dafny.Array(None, 21)
                        for i0_3_ in range(nw28_.length(0)):
                            nw28_[i0_3_] = init3_(i0_3_)
                        d_195_v117_ = nw28_
                        d_198_v118_: _dafny.Set
                        d_198_v118_ = _dafny.Set({d_195_v117_})
                        d_199_v119_: C4
                        nw29_ = C4()
                        nw29_.ctor__(d_191_cf37_, (0) - (((0) - (default__.fm7(d_191_cf37_, d_193_v115_, d_194_v116_, d_56_globalState_))) - (d_191_cf37_)), (d_198_v118_) | (d_198_v118_), d_51_v3_)
                        d_199_v119_ = nw29_
                        d_200_v120_: _dafny.Map
                        d_200_v120_ = _dafny.Map({d_51_v3_: d_195_v117_})
                        d_200_v120_ = ((d_200_v120_) | (d_200_v120_)) | (d_200_v120_)
                        d_200_v120_ = (d_200_v120_).set(d_51_v3_, d_195_v117_)
                        d_201_v121_: C2
                        nw30_ = C2()
                        nw30_.ctor__(d_198_v118_, d_51_v3_)
                        d_201_v121_ = nw30_
                        d_202_v122_: _dafny.Map
                        d_202_v122_ = _dafny.Map({848: d_201_v121_})
                        d_203_v123_: _dafny.Map
                        d_203_v123_ = _dafny.Map({d_191_cf37_: d_202_v122_})
                        d_203_v123_ = _dafny.Map({577: d_202_v122_})
                    elif source1_.is_DC19:
                        d_204___mcc_h5_ = source1_.cf35
                        d_205_cf35_ = d_204___mcc_h5_
                        d_206_v124_: _dafny.Array
                        nw31_ = _dafny.Array(False, 28)
                        d_206_v124_ = nw31_
                        d_207_v125_: _dafny.Set
                        d_207_v125_ = _dafny.Set({d_206_v124_, d_206_v124_, d_206_v124_})
                        d_208_v126_: C3
                        nw32_ = C3()
                        nw32_.ctor__(d_206_v124_, d_207_v125_, d_51_v3_)
                        d_208_v126_ = nw32_
                        d_209_v127_: _dafny.Map
                        d_209_v127_ = _dafny.Map({d_208_v126_: d_49_v1_})
                        d_210_v128_: _dafny.Map
                        d_210_v128_ = _dafny.Map({d_51_v3_: d_49_v1_})
                        d_211_v129_: _dafny.Array
                        nw33_ = _dafny.Array(None, 11)
                        nw33_[int(0)] = d_49_v1_
                        nw33_[int(1)] = d_49_v1_
                        nw33_[int(2)] = ((d_209_v127_)[d_208_v126_] if (d_208_v126_) in (d_209_v127_) else d_49_v1_)
                        nw33_[int(3)] = d_49_v1_
                        nw33_[int(4)] = d_49_v1_
                        nw33_[int(5)] = d_49_v1_
                        nw33_[int(6)] = d_49_v1_
                        nw33_[int(7)] = ((d_210_v128_)[d_51_v3_] if (d_51_v3_) in (d_210_v128_) else d_49_v1_)
                        nw33_[int(8)] = d_49_v1_
                        nw33_[int(9)] = d_49_v1_
                        nw33_[int(10)] = d_49_v1_
                        d_211_v129_ = nw33_
                        index15_ = default__.safeIndex(450, (d_211_v129_).length(0))
                        (d_211_v129_)[index15_] = d_49_v1_
                        d_212_v130_: _dafny.Map
                        d_212_v130_ = _dafny.Map({d_52_v4_: _dafny.SeqWithoutIsStrInference([d_51_v3_, d_51_v3_, d_51_v3_])})
                        index16_ = default__.safeIndex(450, (d_211_v129_).length(0))
                        rhs21_ = (d_52_v4_) + (((d_212_v130_)[d_52_v4_] if (d_52_v4_) in (d_212_v130_) else d_52_v4_))
                        rhs22_ = d_49_v1_
                        rhs23_ = not(not(d_51_v3_))
                        lhs15_ = d_211_v129_
                        lhs16_ = default__.safeIndex(450, (d_211_v129_).length(0))
                        lhs17_ = d_56_globalState_
                        d_52_v4_ = rhs21_
                        lhs15_[lhs16_] = rhs22_
                        lhs17_.f3 = rhs23_
                        (d_56_globalState_).f14 = d_50_v2_
                        d_213_v131_: D9
                        d_213_v131_ = D9_DC19(d_205_cf35_)
                        d_214_v132_: D9
                        d_214_v132_ = D9_DC21(d_213_v131_)
                        d_215_v133_: D9
                        d_215_v133_ = D9_DC21(d_213_v131_)
                        pat_let_tv2_ = d_214_v132_
                        def iife32_(_pat_let4_0):
                            def iife33_(d_216_dt__update__tmp_h2_):
                                def iife34_(_pat_let5_0):
                                    def iife35_(d_217_dt__update_hcf41_h0_):
                                        return D9_DC21(d_217_dt__update_hcf41_h0_)
                                    return iife35_(_pat_let5_0)
                                return iife34_(pat_let_tv2_)
                            return iife33_(_pat_let4_0)
                        d_215_v133_ = iife32_(d_215_v133_)
                        d_206_v124_ = d_206_v124_
                    elif True:
                        d_218___mcc_h6_ = source1_.cf41
                        d_219_cf41_ = d_218___mcc_h6_
                        d_220_v134_: _dafny.Array
                        nw34_ = _dafny.Array(None, 6)
                        d_220_v134_ = nw34_
                        d_221_v135_: D3
                        d_221_v135_ = D3_DC9()
                        d_222_v136_: _dafny.Seq
                        d_222_v136_ = _dafny.SeqWithoutIsStrInference([d_221_v135_])
                        d_223_v137_: D5
                        d_223_v137_ = D5_DC13(d_51_v3_, d_50_v2_, d_222_v136_)
                        d_224_v138_: _dafny.Map
                        d_224_v138_ = _dafny.Map({d_61_v11_: d_51_v3_})
                        d_225_v139_: _dafny.Array
                        def lambda7_(d_226_i14_):
                            return True

                        init4_ = lambda7_
                        nw35_ = _dafny.Array(None, 29)
                        for i0_4_ in range(nw35_.length(0)):
                            nw35_[i0_4_] = init4_(i0_4_)
                        d_225_v139_ = nw35_
                        d_227_v140_: _dafny.Set
                        d_227_v140_ = _dafny.Set({d_225_v139_})
                        d_228_v141_: C7
                        nw36_ = C7()
                        nw36_.ctor__(d_50_v2_, d_224_v138_, d_225_v139_, d_227_v140_, d_51_v3_)
                        d_228_v141_ = nw36_
                        index17_ = default__.safeIndex(584, (d_220_v134_).length(0))
                        (d_220_v134_)[index17_] = (d_228_v141_ if (d_223_v137_).cf23 else d_228_v141_)
                        index18_ = default__.safeIndex(584, (d_220_v134_).length(0))
                        (d_220_v134_)[index18_] = d_228_v141_
                        index19_ = default__.safeIndex(749, (d_49_v1_).length(0))
                        (d_49_v1_)[index19_] = (d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))]
                        d_59_v9_ = (d_59_v9_).set(d_50_v2_, (d_50_v2_) + (d_50_v2_))
                        d_229_v142_: _dafny.Map
                        d_229_v142_ = _dafny.Map({d_51_v3_: d_50_v2_})
                        d_230_v143_: D0
                        d_230_v143_ = D0_DC0(not(True))
                        (d_56_globalState_).f14 = default__.fm7(d_50_v2_, d_229_v142_, d_230_v143_, d_56_globalState_)
                    d_231_v144_: _dafny.MultiSet
                    d_231_v144_ = _dafny.MultiSet([d_51_v3_])
                    d_232_v145_: _dafny.Map
                    d_232_v145_ = _dafny.Map({default__.safeDivisionInt(d_50_v2_, len((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))])): D2_DC5(d_231_v144_)})
                    d_233_v146_: D2
                    d_233_v146_ = D2_DC5(d_231_v144_)
                    d_232_v145_ = (d_232_v145_).set((0) - (d_50_v2_), d_233_v146_)
                    d_234_v147_: _dafny.Array
                    nw37_ = _dafny.Array(_dafny.Array(None, 0), 25)
                    d_234_v147_ = nw37_
                    rhs24_ = d_234_v147_
                    rhs25_ = d_51_v3_
                    rhs26_ = d_51_v3_
                    rhs27_ = d_60_v10_
                    lhs18_ = d_56_globalState_
                    lhs19_ = d_56_globalState_
                    d_234_v147_ = rhs24_
                    lhs18_.f3 = rhs25_
                    lhs19_.f9 = rhs26_
                    d_60_v10_ = rhs27_
                    d_235_v148_: _dafny.Array
                    def lambda8_(d_236_v3_):
                        def lambda9_(d_237_i15_):
                            return _dafny.Set({d_236_v3_})

                        return lambda9_

                    init5_ = lambda8_(d_51_v3_)
                    nw38_ = _dafny.Array(None, 29)
                    for i0_5_ in range(nw38_.length(0)):
                        nw38_[i0_5_] = init5_(i0_5_)
                    d_235_v148_ = nw38_
                    d_238_v149_: _dafny.Set
                    d_238_v149_ = _dafny.Set({not(d_51_v3_)})
                    index20_ = default__.safeIndex(198, (d_235_v148_).length(0))
                    (d_235_v148_)[index20_] = d_238_v149_
                    index21_ = default__.safeIndex(198, (d_235_v148_).length(0))
                    (d_235_v148_)[index21_] = d_238_v149_
                    pass
            pass
        if d_51_v3_:
            if not(((d_50_v2_ if d_51_v3_ else d_50_v2_)) <= (d_50_v2_)):
                (d_56_globalState_).f14 = ((d_53_v5_)[d_50_v2_] if (d_50_v2_) in (d_53_v5_) else d_50_v2_)
                d_239_v150_: _dafny.Seq
                d_239_v150_ = _dafny.SeqWithoutIsStrInference([d_60_v10_])
                d_239_v150_ = (d_239_v150_).set(default__.safeIndex(d_50_v2_, len(d_239_v150_)), (d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))])
                d_240_v151_: _dafny.Array
                nw39_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_240_v151_ = nw39_
                d_241_v152_: _dafny.Seq
                d_241_v152_ = _dafny.SeqWithoutIsStrInference([d_50_v2_])
                d_242_v153_: _dafny.Seq
                d_242_v153_ = _dafny.SeqWithoutIsStrInference([d_241_v152_, d_241_v152_])
                d_243_v154_: _dafny.Array
                nw40_ = _dafny.Array(None, 10)
                nw40_[int(0)] = d_241_v152_
                nw40_[int(1)] = d_241_v152_
                nw40_[int(2)] = (d_242_v153_)[default__.safeIndex(len(d_52_v4_), len(d_242_v153_))]
                nw40_[int(3)] = d_241_v152_
                nw40_[int(4)] = d_241_v152_
                nw40_[int(5)] = d_241_v152_
                nw40_[int(6)] = d_241_v152_
                nw40_[int(7)] = d_241_v152_
                nw40_[int(8)] = d_241_v152_
                nw40_[int(9)] = (d_241_v152_)
                d_243_v154_ = nw40_
                index22_ = default__.safeIndex(16, (d_240_v151_).length(0))
                (d_240_v151_)[index22_] = d_243_v154_
                index23_ = default__.safeIndex(16, (d_240_v151_).length(0))
                nw41_ = _dafny.Array(_dafny.Seq({}), 5)
                (d_240_v151_)[index23_] = nw41_
                d_244_v155_: _dafny.Array
                def lambda10_(d_245_v3_):
                    def lambda11_(d_246_i16_):
                        return d_245_v3_

                    return lambda11_

                init6_ = lambda10_(d_51_v3_)
                nw42_ = _dafny.Array(None, 28)
                for i0_6_ in range(nw42_.length(0)):
                    nw42_[i0_6_] = init6_(i0_6_)
                d_244_v155_ = nw42_
                d_247_v156_: _dafny.Set
                d_247_v156_ = _dafny.Set({d_244_v155_})
                d_248_v157_: T0
                nw43_ = C4()
                nw43_.ctor__(d_50_v2_, 978, d_247_v156_, d_51_v3_)
                d_248_v157_ = nw43_
                d_248_v157_ = d_248_v157_
                d_249_v158_: _dafny.Map
                d_249_v158_ = _dafny.Map({(d_248_v157_).f18: (d_248_v157_).f18})
                d_250_v159_: C7
                nw44_ = C7()
                nw44_.ctor__((d_241_v152_)[default__.safeIndex(-841, len(d_241_v152_))], default__.fm45(d_53_v5_, len(d_249_v158_), d_51_v3_, d_56_globalState_), d_244_v155_, d_248_v157_.f17, default__.fm13(d_56_globalState_))
                d_250_v159_ = nw44_
                d_250_v159_ = d_250_v159_
            elif True:
                d_251_v160_: _dafny.MultiSet
                d_251_v160_ = _dafny.MultiSet([default__.fm30(True, d_56_globalState_)])
                d_251_v160_ = d_251_v160_
                (d_56_globalState_).f7 = d_48_v0_
                d_252_v161_: _dafny.Array
                def lambda12_(d_253_i17_):
                    return False

                init7_ = lambda12_
                nw45_ = _dafny.Array(None, 1)
                for i0_7_ in range(nw45_.length(0)):
                    nw45_[i0_7_] = init7_(i0_7_)
                d_252_v161_ = nw45_
                index24_ = default__.safeIndex(504, (d_252_v161_).length(0))
                (d_252_v161_)[index24_] = d_51_v3_
                d_254_v162_: D0
                d_254_v162_ = D0_DC0(default__.fm13(d_56_globalState_))
                index25_ = default__.safeIndex(504, (d_252_v161_).length(0))
                (d_252_v161_)[index25_] = (default__.safeModuloInt(d_50_v2_, d_50_v2_)) < ((len((d_60_v10_).set(default__.safeIndex(d_50_v2_, len(d_60_v10_)), default__.fm12(d_254_v162_, d_56_globalState_))) if False else (0) - (d_50_v2_)))
                d_255_v163_: _dafny.MultiSet
                d_255_v163_ = _dafny.MultiSet([d_252_v161_])
                index26_ = default__.safeIndex(504, (d_252_v161_).length(0))
                (d_252_v161_)[index26_] = (d_255_v163_) != (d_255_v163_)
                (d_56_globalState_).f14 = d_50_v2_
            d_60_v10_ = ((d_60_v10_) + (default__.fm10(d_56_globalState_)) if True else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udophjwub")))
            index27_ = default__.safeIndex(749, (d_49_v1_).length(0))
            (d_49_v1_)[index27_] = (d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))]
            d_256_v164_: _dafny.Map
            d_256_v164_ = _dafny.Map({d_51_v3_: d_51_v3_})
            d_51_v3_ = (d_51_v3_ if default__.fm13(d_56_globalState_) else (len(d_256_v164_)) > (d_50_v2_))
            d_50_v2_ = d_50_v2_
        elif True:
            d_257_v165_: _dafny.Map
            d_257_v165_ = _dafny.Map({d_50_v2_: (341) <= (len((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))]))})
            d_257_v165_ = (d_257_v165_).set(len(_dafny.SeqWithoutIsStrInference([d_51_v3_])), not(((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))]) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpts")))))
            (d_56_globalState_).f14 = len((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))])
            (d_56_globalState_).f3 = (d_51_v3_) and (d_51_v3_)
            index28_ = default__.safeIndex(782, (d_54_v6_).length(0))
            (d_54_v6_)[index28_] = d_61_v11_
            index29_ = default__.safeIndex(782, (d_54_v6_).length(0))
            (d_54_v6_)[index29_] = _dafny.CodePoint('r')
            d_258_v166_: D2
            d_258_v166_ = D2_DC6(len((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))]))
            d_259_v167_: D2
            d_259_v167_ = D2_DC7(d_258_v166_)
            source2_ = d_259_v167_
            if source2_.is_DC6:
                d_260___mcc_h7_ = source2_.cf13
                d_261_cf13_ = d_260___mcc_h7_
                index30_ = default__.safeIndex(839, (d_48_v0_).length(0))
                (d_48_v0_)[index30_] = d_261_cf13_
                index31_ = default__.safeIndex(839, (d_48_v0_).length(0))
                (d_48_v0_)[index31_] = d_261_cf13_
                d_262_v168_: _dafny.Map
                d_262_v168_ = _dafny.Map({d_51_v3_: d_51_v3_})
                d_263_v169_: _dafny.Array
                def lambda13_(d_264_v3_):
                    def lambda14_(d_265_i18_):
                        return d_264_v3_

                    return lambda14_

                init8_ = lambda13_(d_51_v3_)
                nw46_ = _dafny.Array(None, 20)
                for i0_8_ in range(nw46_.length(0)):
                    nw46_[i0_8_] = init8_(i0_8_)
                d_263_v169_ = nw46_
                d_266_v170_: _dafny.Set
                d_266_v170_ = _dafny.Set({d_263_v169_, d_263_v169_})
                d_267_v171_: T0
                nw47_ = C2()
                nw47_.ctor__(d_266_v170_, d_51_v3_)
                d_267_v171_ = nw47_
                d_268_v172_: _dafny.Map
                d_268_v172_ = _dafny.Map({(d_267_v171_ if ((d_262_v168_)[d_51_v3_] if (d_51_v3_) in (d_262_v168_) else d_51_v3_) else d_267_v171_): d_50_v2_})
                d_268_v172_ = (d_268_v172_).set(d_267_v171_, (d_48_v0_)[default__.safeIndex(839, (d_48_v0_).length(0))])
                index32_ = default__.safeIndex(839, (d_48_v0_).length(0))
                (d_48_v0_)[index32_] = d_261_cf13_
                d_269_v173_: _dafny.MultiSet
                d_269_v173_ = _dafny.MultiSet([not(not(((d_257_v165_)[(0) - (d_261_cf13_)] if ((0) - (d_261_cf13_)) in (d_257_v165_) else d_51_v3_))), True, (d_267_v171_).f18, (d_267_v171_).f18])
                d_269_v173_ = ((d_269_v173_).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_51_v3_])))).intersection(d_269_v173_)
            elif source2_.is_DC5:
                d_270___mcc_h8_ = source2_.cf12
                d_271_cf12_ = d_270___mcc_h8_
                d_51_v3_ = d_51_v3_
                d_272_v174_: D0
                d_272_v174_ = D0_DC0(d_51_v3_)
                d_273_v175_: _dafny.Array
                def lambda15_(d_274_v3_):
                    def lambda16_(d_275_i19_):
                        return d_274_v3_

                    return lambda16_

                init9_ = lambda15_(d_51_v3_)
                nw48_ = _dafny.Array(None, 12)
                for i0_9_ in range(nw48_.length(0)):
                    nw48_[i0_9_] = init9_(i0_9_)
                d_273_v175_ = nw48_
                d_276_v176_: _dafny.Set
                d_276_v176_ = _dafny.Set({d_273_v175_, d_273_v175_})
                d_277_v177_: _dafny.Map
                d_277_v177_ = _dafny.Map({d_61_v11_: d_51_v3_})
                d_278_v178_: C8
                nw49_ = C8()
                def iife36_(_pat_let6_0):
                    def iife37_(d_279_dt__update__tmp_h3_):
                        def iife38_(_pat_let7_0):
                            def iife39_(d_280_dt__update_hcf0_h0_):
                                return D0_DC0(d_280_dt__update_hcf0_h0_)
                            return iife39_(_pat_let7_0)
                        return iife38_(True)
                    return iife37_(_pat_let6_0)
                nw49_.ctor__(default__.fm7(d_50_v2_, _dafny.Map({not(d_51_v3_): 274}), iife36_(d_272_v174_), d_56_globalState_), d_51_v3_, d_273_v175_, d_276_v176_, ((d_54_v6_)[default__.safeIndex(782, (d_54_v6_).length(0))]) not in (d_60_v10_), ((d_59_v9_)[d_50_v2_] if (d_50_v2_) in (d_59_v9_) else d_50_v2_), d_277_v177_)
                d_278_v178_ = nw49_
                d_281_v179_: C5
                nw50_ = C5()
                nw50_.ctor__(d_60_v10_, (d_278_v178_).f24, d_50_v2_, d_277_v177_, d_273_v175_, d_276_v176_, d_51_v3_)
                d_281_v179_ = nw50_
                d_281_v179_ = d_281_v179_
                index33_ = default__.safeIndex(475, (d_273_v175_).length(0))
                (d_273_v175_)[index33_] = (d_278_v178_).f24
                index34_ = default__.safeIndex(66, (d_273_v175_).length(0))
                (d_273_v175_)[index34_] = not((d_281_v179_).f30)
                d_282_v180_: _dafny.Map
                d_282_v180_ = _dafny.Map({(d_281_v179_).f30: (d_278_v178_).f24})
                d_283_v181_: D0
                d_283_v181_ = D0_DC1((d_49_v1_)[default__.safeIndex(749, (d_49_v1_).length(0))], len(d_282_v180_), d_278_v178_.f23, d_61_v11_)
                pat_let_tv3_ = d_54_v6_
                pat_let_tv4_ = d_54_v6_
                pat_let_tv5_ = d_50_v2_
                index35_ = default__.safeIndex(475, (d_273_v175_).length(0))
                index36_ = default__.safeIndex(66, (d_273_v175_).length(0))
                def iife40_(_pat_let8_0):
                    def iife41_(d_284_dt__update__tmp_h4_):
                        def iife42_(_pat_let9_0):
                            def iife43_(d_285_dt__update_hcf4_h0_):
                                def iife44_(_pat_let10_0):
                                    def iife45_(d_286_dt__update_hcf3_h0_):
                                        return D0_DC1((d_284_dt__update__tmp_h4_).cf1, (d_284_dt__update__tmp_h4_).cf2, d_286_dt__update_hcf3_h0_, d_285_dt__update_hcf4_h0_)
                                    return iife45_(_pat_let10_0)
                                return iife44_(pat_let_tv5_)
                            return iife43_(_pat_let9_0)
                        return iife42_((pat_let_tv4_)[default__.safeIndex(782, (pat_let_tv3_).length(0))])
                    return iife41_(_pat_let8_0)
                rhs28_ = (d_278_v178_).f24
                rhs29_ = (d_50_v2_) - ((iife40_(d_283_v181_)).cf3)
                rhs30_ = d_278_v178_.f23
                rhs31_ = not((d_278_v178_.f23) > (d_50_v2_))
                lhs20_ = d_273_v175_
                lhs21_ = default__.safeIndex(475, (d_273_v175_).length(0))
                lhs22_ = d_56_globalState_
                lhs23_ = d_273_v175_
                lhs24_ = default__.safeIndex(66, (d_273_v175_).length(0))
                lhs20_[lhs21_] = rhs28_
                d_50_v2_ = rhs29_
                lhs22_.f14 = rhs30_
                lhs23_[lhs24_] = rhs31_
            elif True:
                d_287___mcc_h9_ = source2_.cf14
                d_288_cf14_ = d_287___mcc_h9_
                d_289_v183_: _dafny.Map
                d_289_v183_ = _dafny.Map({d_61_v11_: d_50_v2_})
                d_290_v184_: _dafny.Array
                def lambda17_(d_291_v10_, d_292_v3_, d_293_v1_, d_294_v2_):
                    def lambda18_(d_295_i20_):
                        return (D3_DC10(d_291_v10_, d_292_v3_, (d_293_v1_)[default__.safeIndex(749, (d_293_v1_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "go")), d_294_v2_)).cf17

                    return lambda18_

                init10_ = lambda17_(d_60_v10_, d_51_v3_, d_49_v1_, d_50_v2_)
                nw51_ = _dafny.Array(None, 20)
                for i0_10_ in range(nw51_.length(0)):
                    nw51_[i0_10_] = init10_(i0_10_)
                d_290_v184_ = nw51_
                d_296_v185_: C7
                nw52_ = C7()
                def iife46_():
                    coll24_ = _dafny.Map()
                    compr_24_: str
                    for compr_24_ in (d_289_v183_).keys.Elements:
                        d_297_v182_: str = compr_24_
                        if (d_297_v182_) in (d_289_v183_):
                            coll24_[d_297_v182_] = False
                    return _dafny.Map(coll24_)
                nw52_.ctor__(900, iife46_()
                , d_290_v184_, _dafny.Set({d_290_v184_, d_290_v184_, d_290_v184_, d_290_v184_, d_290_v184_}), d_51_v3_)
                d_296_v185_ = nw52_
                d_296_v185_ = d_296_v185_
                d_298_v186_: _dafny.Map
                d_298_v186_ = _dafny.Map({(d_52_v4_)[default__.safeIndex(d_50_v2_, len(d_52_v4_))]: d_57_v7_})
                d_299_v187_: _dafny.Map
                d_299_v187_ = _dafny.Map({(d_54_v6_)[default__.safeIndex(782, (d_54_v6_).length(0))]: d_61_v11_})
                d_300_v188_: _dafny.Map
                d_300_v188_ = _dafny.Map({d_299_v187_: d_298_v186_})
                d_301_v189_: _dafny.MultiSet
                d_301_v189_ = _dafny.MultiSet([d_51_v3_, d_51_v3_])
                d_302_v190_: _dafny.Set
                d_302_v190_ = _dafny.Set({d_290_v184_})
                d_303_v191_: _dafny.Array
                def lambda19_(d_304_v9_):
                    def lambda20_(d_305_i21_):
                        return d_304_v9_

                    return lambda20_

                init11_ = lambda19_(d_59_v9_)
                nw53_ = _dafny.Array(None, 10)
                for i0_11_ in range(nw53_.length(0)):
                    nw53_[i0_11_] = init11_(i0_11_)
                d_303_v191_ = nw53_
                d_306_v192_: D7
                d_306_v192_ = D7_DC17(d_303_v191_, d_51_v3_, d_50_v2_)
                d_307_v193_: _dafny.Seq
                d_307_v193_ = _dafny.SeqWithoutIsStrInference([d_306_v192_])
                d_308_v194_: D12
                d_308_v194_ = D12_DC30(d_301_v189_, d_51_v3_, D11_DC25(d_302_v190_), d_307_v193_, False)
                d_309_v195_: _dafny.Array
                nw54_ = _dafny.Array(None, 14)
                nw54_[int(0)] = d_298_v186_
                nw54_[int(1)] = ((d_300_v188_)[d_299_v187_] if (d_299_v187_) in (d_300_v188_) else d_298_v186_)
                nw54_[int(2)] = (d_298_v186_ if d_51_v3_ else default__.fm46(d_56_globalState_))
                nw54_[int(3)] = (d_298_v186_).set((d_308_v194_).cf59, d_57_v7_)
                nw54_[int(4)] = (d_298_v186_) | (d_298_v186_)
                nw54_[int(5)] = d_298_v186_
                nw54_[int(6)] = d_298_v186_
                nw54_[int(7)] = d_298_v186_
                nw54_[int(8)] = d_298_v186_
                nw54_[int(9)] = (d_298_v186_).set(d_51_v3_, d_57_v7_)
                nw54_[int(10)] = d_298_v186_
                nw54_[int(11)] = (d_298_v186_) | (d_298_v186_)
                nw54_[int(12)] = (d_298_v186_ if d_51_v3_ else d_298_v186_)
                nw54_[int(13)] = d_298_v186_
                d_309_v195_ = nw54_
                index37_ = default__.safeIndex(774, (d_309_v195_).length(0))
                (d_309_v195_)[index37_] = d_298_v186_
                index38_ = default__.safeIndex(774, (d_309_v195_).length(0))
                (d_309_v195_)[index38_] = (d_298_v186_) | (d_298_v186_)
                index39_ = default__.safeIndex(782, (d_54_v6_).length(0))
                (d_54_v6_)[index39_] = _dafny.CodePoint('r')
                d_310_v196_: _dafny.Seq
                d_310_v196_ = _dafny.SeqWithoutIsStrInference([796])
                d_311_v197_: _dafny.Seq
                d_311_v197_ = _dafny.SeqWithoutIsStrInference([d_310_v196_, _dafny.SeqWithoutIsStrInference([d_50_v2_])])
                (d_56_globalState_).f14 = (0) - (default__.fm7(d_50_v2_, _dafny.Map({True: len((d_311_v197_)[default__.safeIndex((d_53_v5_).cardinality, len(d_311_v197_))])}), D0_DC0(not(default__.fm13(d_56_globalState_))), d_56_globalState_))
        d_312_v198_: _dafny.Array
        def lambda21_(d_313_v3_):
            def lambda22_(d_314_i22_):
                return d_313_v3_

            return lambda22_

        init12_ = lambda21_(d_51_v3_)
        nw55_ = _dafny.Array(None, 12)
        for i0_12_ in range(nw55_.length(0)):
            nw55_[i0_12_] = init12_(i0_12_)
        d_312_v198_ = nw55_
        d_315_v199_: C1
        nw56_ = C1()
        nw56_.ctor__()
        d_315_v199_ = nw56_
        d_316_v200_: _dafny.Map
        d_316_v200_ = _dafny.Map({d_315_v199_: d_315_v199_})
        d_317_v201_: _dafny.Seq
        d_317_v201_ = _dafny.SeqWithoutIsStrInference([len(d_316_v200_)])
        d_318_v202_: C2
        nw57_ = C2()
        nw57_.ctor__(_dafny.Set({d_312_v198_}), not((d_50_v2_) != (len(d_317_v201_))))
        d_318_v202_ = nw57_
        d_319_v203_: D3
        d_319_v203_ = D3_DC9()
        d_320_v204_: _dafny.Seq
        d_320_v204_ = _dafny.SeqWithoutIsStrInference([d_319_v203_])
        d_321_v205_: _dafny.Seq
        d_321_v205_ = d_320_v204_
        d_322_i23_: int
        d_322_i23_ = 0
        with _dafny.label("1"):
            pat_let_tv6_ = d_51_v3_
            def lambda23_(source3_):
                d_338___mcc_h10_ = source3_
                d_339_cf21_ = d_338___mcc_h10_
                return pat_let_tv6_

            while lambda23_(d_321_v205_):
                with _dafny.c_label("1"):
                    if (d_322_i23_) >= (100):
                        raise _dafny.Break("1")
                    d_322_i23_ = (d_322_i23_) + (1)
                    d_323_v206_: _dafny.Set
                    d_323_v206_ = _dafny.Set({d_48_v0_})
                    d_324_v207_: _dafny.Map
                    d_324_v207_ = _dafny.Map({d_51_v3_: (d_323_v206_).ispropersubset(d_323_v206_)})
                    rhs32_ = d_312_v198_
                    rhs33_ = d_50_v2_
                    rhs34_ = ((d_324_v207_) | (d_324_v207_)).set((d_51_v3_) and (False), d_51_v3_)
                    lhs25_ = d_56_globalState_
                    d_312_v198_ = rhs32_
                    lhs25_.f14 = rhs33_
                    d_324_v207_ = rhs34_
                    d_325_v208_: _dafny.Seq
                    d_326_v209_: D0
                    d_327_v210_: bool
                    d_328_v211_: bool
                    out13_: _dafny.Seq
                    out14_: D0
                    out15_: bool
                    out16_: bool
                    out13_, out14_, out15_, out16_ = (d_315_v199_).m5((0) - (d_50_v2_), d_56_globalState_)
                    d_325_v208_ = out13_
                    d_326_v209_ = out14_
                    d_327_v210_ = out15_
                    d_328_v211_ = out16_
                    d_329_v212_: _dafny.MultiSet
                    d_329_v212_ = _dafny.MultiSet([d_327_v210_])
                    d_329_v212_ = d_329_v212_
                    d_330_v215_: D10
                    def iife47_():
                        coll25_ = _dafny.Set()
                        compr_25_: int
                        for compr_25_ in (d_317_v201_).Elements:
                            d_331_v214_: int = compr_25_
                            if (d_331_v214_) in (d_317_v201_):
                                coll25_ = coll25_.union(_dafny.Set([default__.safeModuloInt(d_331_v214_, 238)]))
                        return _dafny.Set(coll25_)
                    d_330_v215_ = D10_DC22(iife47_()
)
                    d_332_v216_: _dafny.Map
                    d_332_v216_ = _dafny.Map({len(d_59_v9_): d_330_v215_})
                    d_333_v217_: _dafny.Seq
                    d_334_v218_: D0
                    d_335_v219_: bool
                    d_336_v220_: bool
                    out17_: _dafny.Seq
                    out18_: D0
                    out19_: bool
                    out20_: bool
                    def iife48_():
                        coll26_ = _dafny.Map()
                        compr_26_: int
                        for compr_26_ in _dafny.IntegerRange(-608, 634):
                            d_337_v213_: int = compr_26_
                            if ((-608) <= (d_337_v213_)) and ((d_337_v213_) < (634)):
                                coll26_[default__.safeModuloInt(d_337_v213_, d_50_v2_)] = D10_DC22(d_57_v7_)
                        return _dafny.Map(coll26_)
                    out17_, out18_, out19_, out20_ = (d_315_v199_).m5(len((iife48_()
                    ) | (d_332_v216_)), d_56_globalState_)
                    d_333_v217_ = out17_
                    d_334_v218_ = out18_
                    d_335_v219_ = out19_
                    d_336_v220_ = out20_
                    pass
            pass
        d_340_v221_: _dafny.Seq
        d_341_v222_: D0
        d_342_v223_: bool
        d_343_v224_: bool
        out21_: _dafny.Seq
        out22_: D0
        out23_: bool
        out24_: bool
        out21_, out22_, out23_, out24_ = (d_315_v199_).m5((d_50_v2_) - (d_50_v2_), d_56_globalState_)
        d_340_v221_ = out21_
        d_341_v222_ = out22_
        d_342_v223_ = out23_
        d_343_v224_ = out24_
        d_344_v225_: _dafny.Seq
        d_345_v226_: D0
        d_346_v227_: bool
        d_347_v228_: bool
        out25_: _dafny.Seq
        out26_: D0
        out27_: bool
        out28_: bool
        out25_, out26_, out27_, out28_ = (d_315_v199_).m5(832, d_56_globalState_)
        d_344_v225_ = out25_
        d_345_v226_ = out26_
        d_346_v227_ = out27_
        d_347_v228_ = out28_
        d_342_v223_ = (d_343_v224_ if (d_50_v2_) <= (((d_59_v9_)[d_50_v2_] if (d_50_v2_) in (d_59_v9_) else d_50_v2_)) else not(d_347_v228_))
        (d_56_globalState_).f3 = d_342_v223_
        d_348_v229_: _dafny.Seq
        d_349_v230_: D0
        d_350_v231_: bool
        d_351_v232_: bool
        out29_: _dafny.Seq
        out30_: D0
        out31_: bool
        out32_: bool
        out29_, out30_, out31_, out32_ = (d_315_v199_).m5((len(d_60_v10_)) * (d_50_v2_), d_56_globalState_)
        d_348_v229_ = out29_
        d_349_v230_ = out30_
        d_350_v231_ = out31_
        d_351_v232_ = out32_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_312_v198_).length(0)):
            d_352_i24_: int = guard_loop_0_
            if (True) and (((0) <= (d_352_i24_)) and ((d_352_i24_) < ((d_312_v198_).length(0)))):
                (d_312_v198_)[(d_352_i24_)] = (not(not(d_350_v231_)) if (d_61_v11_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vws"))) else d_350_v231_)
        d_353_v233_: _dafny.MultiSet
        d_353_v233_ = _dafny.MultiSet([d_350_v231_, d_51_v3_])
        d_354_i25_: int
        d_354_i25_ = 0
        with _dafny.label("2"):
            while (((d_353_v233_)[d_343_v224_] if (d_343_v224_) in (d_353_v233_) else len(d_52_v4_))) == (d_50_v2_):
                with _dafny.c_label("2"):
                    if (d_354_i25_) >= (100):
                        raise _dafny.Break("2")
                    d_354_i25_ = (d_354_i25_) + (1)
                    d_355_v234_: _dafny.Set
                    d_355_v234_ = _dafny.Set({d_312_v198_})
                    d_356_v235_: _dafny.Map
                    d_356_v235_ = _dafny.Map({d_61_v11_: False})
                    d_357_v236_: T2
                    nw58_ = C8()
                    nw58_.ctor__(len(d_60_v10_), d_346_v227_, d_312_v198_, d_355_v234_, d_342_v223_, d_50_v2_, d_356_v235_)
                    d_357_v236_ = nw58_
                    d_358_v237_: _dafny.Array
                    nw59_ = _dafny.Array(None, 3)
                    nw59_[int(0)] = d_357_v236_
                    nw59_[int(1)] = d_357_v236_
                    nw59_[int(2)] = d_357_v236_
                    d_358_v237_ = nw59_
                    index40_ = default__.safeIndex(733, (d_358_v237_).length(0))
                    (d_358_v237_)[index40_] = d_357_v236_
                    index41_ = default__.safeIndex(837, (d_312_v198_).length(0))
                    (d_312_v198_)[index41_] = d_51_v3_
                    index42_ = default__.safeIndex(733, (d_358_v237_).length(0))
                    index43_ = default__.safeIndex(837, (d_312_v198_).length(0))
                    rhs35_ = d_357_v236_
                    rhs36_ = (d_357_v236_).f20
                    rhs37_ = d_346_v227_
                    rhs38_ = not(d_343_v224_)
                    rhs39_ = default__.safeModuloInt(915, d_50_v2_)
                    lhs26_ = d_358_v237_
                    lhs27_ = default__.safeIndex(733, (d_358_v237_).length(0))
                    lhs28_ = d_56_globalState_
                    lhs29_ = d_312_v198_
                    lhs30_ = default__.safeIndex(837, (d_312_v198_).length(0))
                    lhs31_ = d_56_globalState_
                    lhs26_[lhs27_] = rhs35_
                    lhs28_.f14 = rhs36_
                    lhs29_[lhs30_] = rhs37_
                    lhs31_.f3 = rhs38_
                    d_50_v2_ = rhs39_
                    d_359_v238_: C3
                    nw60_ = C3()
                    nw60_.ctor__(d_357_v236_.f19, (d_355_v234_) | (d_357_v236_.f17), d_342_v223_)
                    d_359_v238_ = nw60_
                    d_360_v239_: _dafny.Array
                    nw61_ = _dafny.Array(_dafny.Set({}), 28)
                    d_360_v239_ = nw61_
                    d_361_v240_: _dafny.Set
                    d_361_v240_ = _dafny.Set({d_48_v0_})
                    index44_ = default__.safeIndex(785, (d_360_v239_).length(0))
                    (d_360_v239_)[index44_] = d_361_v240_
                    d_362_v241_: T1
                    nw62_ = C3()
                    nw62_.ctor__(d_357_v236_.f19, d_357_v236_.f17, (d_357_v236_).f18)
                    d_362_v241_ = nw62_
                    d_363_v242_: _dafny.Map
                    d_363_v242_ = _dafny.Map({d_50_v2_: d_362_v241_})
                    index45_ = default__.safeIndex(785, (d_360_v239_).length(0))
                    rhs40_ = ((d_361_v240_).intersection(d_361_v240_)) - (d_361_v240_)
                    rhs41_ = len(d_363_v242_)
                    lhs32_ = d_360_v239_
                    lhs33_ = default__.safeIndex(785, (d_360_v239_).length(0))
                    lhs32_[lhs33_] = rhs40_
                    d_50_v2_ = rhs41_
                    d_364_v243_: _dafny.Map
                    d_364_v243_ = _dafny.Map({D2_DC5(d_353_v233_): True})
                    d_365_v244_: D2
                    d_365_v244_ = D2_DC5(d_353_v233_)
                    d_366_v246_: _dafny.Seq
                    d_366_v246_ = _dafny.SeqWithoutIsStrInference([d_365_v244_])
                    d_367_v248_: _dafny.Seq
                    d_367_v248_ = _dafny.SeqWithoutIsStrInference([d_364_v243_])
                    d_368_v252_: _dafny.MultiSet
                    d_368_v252_ = _dafny.MultiSet([d_365_v244_])
                    d_369_v253_: _dafny.Array
                    nw63_ = _dafny.Array(None, 24)
                    nw63_[int(0)] = d_364_v243_
                    nw63_[int(1)] = default__.fm47(d_347_v228_, d_56_globalState_)
                    nw63_[int(2)] = (default__.fm47(d_51_v3_, d_56_globalState_)) | (_dafny.Map({d_365_v244_: d_346_v227_}))
                    nw63_[int(3)] = (d_364_v243_ if (d_357_v236_).f18 else _dafny.Map({d_365_v244_: False}))
                    nw63_[int(4)] = d_364_v243_
                    nw63_[int(5)] = (d_364_v243_) | (d_364_v243_)
                    def iife49_():
                        coll27_ = _dafny.Map()
                        compr_27_: D2
                        for compr_27_ in ((d_366_v246_).set(default__.safeIndex((d_357_v236_).f20, len(d_366_v246_)), D2_DC5(_dafny.MultiSet([not(d_51_v3_)])))).Elements:
                            d_370_v245_: D2 = compr_27_
                            if (d_370_v245_) in ((d_366_v246_).set(default__.safeIndex((d_357_v236_).f20, len(d_366_v246_)), D2_DC5(_dafny.MultiSet([not(d_51_v3_)])))):
                                coll27_[d_370_v245_] = d_343_v224_
                        return _dafny.Map(coll27_)
                    nw63_[int(6)] = iife49_()
                    
                    def iife50_():
                        coll28_ = _dafny.Map()
                        compr_28_: D2
                        for compr_28_ in (_dafny.SeqWithoutIsStrInference([d_365_v244_ for d_371_i26_ in range(default__.abs(830))])).Elements:
                            d_372_v247_: D2 = compr_28_
                            if (d_372_v247_) in (_dafny.SeqWithoutIsStrInference([d_365_v244_ for d_371_i26_ in range(default__.abs(830))])):
                                coll28_[d_372_v247_] = (d_357_v236_).f18
                        return _dafny.Map(coll28_)
                    nw63_[int(7)] = (iife50_()
                    ).set(d_365_v244_, d_51_v3_)
                    nw63_[int(8)] = d_364_v243_
                    nw63_[int(9)] = (d_364_v243_) | (d_364_v243_)
                    nw63_[int(10)] = d_364_v243_
                    nw63_[int(11)] = _dafny.Map({d_365_v244_: d_51_v3_})
                    nw63_[int(12)] = (d_364_v243_) | ((d_367_v248_)[default__.safeIndex(157, len(d_367_v248_))])
                    def iife51_():
                        coll29_ = _dafny.Map()
                        compr_29_: D2
                        for compr_29_ in (d_366_v246_).Elements:
                            d_373_v249_: D2 = compr_29_
                            if (d_373_v249_) in (d_366_v246_):
                                coll29_[d_373_v249_] = (D5_DC14(False, len(_dafny.Set({(d_357_v236_).f20, (d_357_v236_).f20})), (d_357_v236_).f20)).cf26
                        return _dafny.Map(coll29_)
                    nw63_[int(13)] = iife51_()
                    
                    def iife52_():
                        coll30_ = _dafny.Map()
                        compr_30_: D2
                        for compr_30_ in ((d_366_v246_).set(default__.safeIndex(d_50_v2_, len(d_366_v246_)), d_365_v244_)).Elements:
                            d_374_v250_: D2 = compr_30_
                            if (d_374_v250_) in ((d_366_v246_).set(default__.safeIndex(d_50_v2_, len(d_366_v246_)), d_365_v244_)):
                                coll30_[d_374_v250_] = d_351_v232_
                        return _dafny.Map(coll30_)
                    nw63_[int(14)] = iife52_()
                    
                    nw63_[int(15)] = (d_364_v243_) | (d_364_v243_)
                    nw63_[int(16)] = d_364_v243_
                    nw63_[int(17)] = default__.fm47(d_51_v3_, d_56_globalState_)
                    nw63_[int(18)] = d_364_v243_
                    def iife53_():
                        coll31_ = _dafny.Map()
                        compr_31_: D2
                        for compr_31_ in (d_368_v252_).Elements:
                            d_375_v251_: D2 = compr_31_
                            if (d_375_v251_) in (d_368_v252_):
                                coll31_[d_375_v251_] = (d_357_v236_).f18
                        return _dafny.Map(coll31_)
                    nw63_[int(19)] = iife53_()
                    
                    nw63_[int(20)] = d_364_v243_
                    nw63_[int(21)] = default__.fm47(d_343_v224_, d_56_globalState_)
                    nw63_[int(22)] = d_364_v243_
                    nw63_[int(23)] = d_364_v243_
                    d_369_v253_ = nw63_
                    index46_ = default__.safeIndex(115, (d_369_v253_).length(0))
                    def iife54_():
                        coll32_ = _dafny.Map()
                        compr_32_: D2
                        for compr_32_ in (d_368_v252_).Elements:
                            d_376_v254_: D2 = compr_32_
                            if (d_376_v254_) in (d_368_v252_):
                                coll32_[d_376_v254_] = (d_362_v241_).f18
                        return _dafny.Map(coll32_)
                    (d_369_v253_)[index46_] = iife54_()
                    
                    d_377_v255_: _dafny.Array
                    def lambda24_(d_378_v201_):
                        def lambda25_(d_379_i27_):
                            return d_378_v201_

                        return lambda25_

                    init13_ = lambda24_(d_317_v201_)
                    nw64_ = _dafny.Array(None, 25)
                    for i0_13_ in range(nw64_.length(0)):
                        nw64_[i0_13_] = init13_(i0_13_)
                    d_377_v255_ = nw64_
                    index47_ = default__.safeIndex(688, (d_377_v255_).length(0))
                    (d_377_v255_)[index47_] = (d_317_v201_) + (d_317_v201_)
                    d_380_v256_: _dafny.Map
                    d_380_v256_ = _dafny.Map({(d_357_v236_).f20: d_350_v231_})
                    index48_ = default__.safeIndex(115, (d_369_v253_).length(0))
                    index49_ = default__.safeIndex(688, (d_377_v255_).length(0))
                    rhs42_ = d_340_v221_
                    rhs43_ = (d_357_v236_).f20
                    rhs44_ = ((_dafny.Map({d_365_v244_: ((d_380_v256_)[-278] if (-278) in (d_380_v256_) else d_346_v227_)}) if False else default__.fm47(False, d_56_globalState_))) | (d_364_v243_)
                    rhs45_ = (d_317_v201_).set(default__.safeIndex((d_357_v236_).f20, len(d_317_v201_)), (0) - (d_50_v2_))
                    lhs34_ = d_56_globalState_
                    lhs35_ = d_369_v253_
                    lhs36_ = default__.safeIndex(115, (d_369_v253_).length(0))
                    lhs37_ = d_377_v255_
                    lhs38_ = default__.safeIndex(688, (d_377_v255_).length(0))
                    d_344_v225_ = rhs42_
                    lhs34_.f14 = rhs43_
                    lhs35_[lhs36_] = rhs44_
                    lhs37_[lhs38_] = rhs45_
                    pass
            pass
        (d_56_globalState_).f3 = d_351_v232_
        _dafny.print(((d_49_v1_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_50_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_51_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_52_v4_) == (_dafny.SeqWithoutIsStrInference([False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_53_v5_) == (_dafny.MultiSet([-548, 2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v6_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v6_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v6_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v6_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v6_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v6_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v6_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v6_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v6_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_56_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_56_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_56_globalState_).f13)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_56_globalState_.f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_.f15) == (_dafny.MultiSet([0, 0, 3288, 3288]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_.f16)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_.f16)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_.f16)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_.f16)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_.f16)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_.f16)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_.f16)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_.f16)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_globalState_.f16)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_57_v7_) == (_dafny.Set({548}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_58_v8_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({548}), _dafny.Set({548}), _dafny.Set({548})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_v9_) == (_dafny.Map({3: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_60_v10_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_61_v11_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_62_v12_) == (_dafny.Map({_dafny.Map({3: 1}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvxdaq"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_181_i12_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v198_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v198_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v198_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v198_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v198_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v198_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v198_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v198_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v198_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v198_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v198_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v198_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_316_v200_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_317_v201_) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_320_v204_) == (_dafny.SeqWithoutIsStrInference([D3_DC9()]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_321_v205_)) == (_dafny.SeqWithoutIsStrInference([D3_DC9()]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_322_i23_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_340_v221_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_341_v222_).cf1).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_341_v222_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_341_v222_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_341_v222_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_342_v223_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_343_v224_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_344_v225_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_345_v226_).cf1).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_345_v226_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_345_v226_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_345_v226_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_346_v227_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_347_v228_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_348_v229_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_349_v230_).cf1).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v230_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v230_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v230_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_350_v231_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_351_v232_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_353_v233_) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_354_i25_))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({self.cf1.VerbatimString(True)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf5', Any), ('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(False, _dafny.Set({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC6(D2, NamedTuple('DC6', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC9(D3, NamedTuple('DC9', [])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9)
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf16', Any), ('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({self.cf16.VerbatimString(True)}, {_dafny.string_of(self.cf17)}, {self.cf18.VerbatimString(True)}, {self.cf19.VerbatimString(True)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf15 == __o.cf15
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

class D4_DC11(D4, NamedTuple('DC11', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13(False, int(0), _dafny.Seq({}))
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

class D5_DC13(D5, NamedTuple('DC13', [('cf23', Any), ('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf22 == __o.cf22
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

class D6_DC15(D6, NamedTuple('DC15', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC17(_dafny.Array(None, 0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)

class D7_DC17(D7, NamedTuple('DC17', [('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC16(D7, NamedTuple('DC16', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf30 == __o.cf30
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

class D8_DC18(D8, NamedTuple('DC18', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC20(_dafny.Map({}), int(0), False, _dafny.Map({}), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D9_DC19)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)

class D9_DC20(D9, NamedTuple('DC20', [('cf36', Any), ('cf37', Any), ('cf38', Any), ('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC19(D9, NamedTuple('DC19', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC19({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC19) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC21(D9, NamedTuple('DC21', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC23(False, False, None)
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

class D10_DC23(D10, NamedTuple('DC23', [('cf43', Any), ('cf44', Any), ('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC22(D10, NamedTuple('DC22', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC22({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC22) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC24(D10, NamedTuple('DC24', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC24({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC24) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC26(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D11_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D11_DC28)

class D11_DC26(D11, NamedTuple('DC26', [('cf48', Any), ('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf48 == __o.cf48 and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf50', Any), ('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC25(D11, NamedTuple('DC25', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC28(D11, NamedTuple('DC28', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC30(_dafny.MultiSet({}), False, D11.default()(), _dafny.Seq({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D12_DC30)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)

class D12_DC30(D12, NamedTuple('DC30', [('cf55', Any), ('cf56', Any), ('cf57', Any), ('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC30({_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC30) and self.cf55 == __o.cf55 and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC29(D12, NamedTuple('DC29', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC32(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D13_DC32)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D13_DC31)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D13_DC33)

class D13_DC32(D13, NamedTuple('DC32', [('cf61', Any), ('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC32({_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC32) and self.cf61 == __o.cf61 and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC31(D13, NamedTuple('DC31', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC31({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC31) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC33(D13, NamedTuple('DC33', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC33({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC33) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC35()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D14_DC35)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D14_DC34)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)

class D14_DC35(D14, NamedTuple('DC35', [])):
    def __dafnystr__(self) -> str:
        return f'D14.DC35'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC35)
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC34(D14, NamedTuple('DC34', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC34({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC34) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC36(D14, NamedTuple('DC36', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f17(self):
        return self._f17
    @f17.setter
    def f17(self, value):
        self._f17 = value

class T1:
    pass
    @property
    def f19(self):
        return self._f19
    @f19.setter
    def f19(self, value):
        self._f19 = value
    def m0(self, p0, p1, p2, globalState):
        pass


class T2:
    pass
    @property
    def f21(self):
        return self._f21
    @f21.setter
    def f21(self, value):
        self._f21 = value

class GlobalState:
    def  __init__(self):
        self.f3: bool = False
        self.f7: _dafny.Array = _dafny.Array(None, 0)
        self.f9: bool = False
        self.f14: int = int(0)
        self.f15: _dafny.MultiSet = _dafny.MultiSet({})
        self.f16: _dafny.Array = _dafny.Array(None, 0)
        self._f0: bool = False
        self._f1: bool = False
        self._f2: bool = False
        self._f4: int = int(0)
        self._f5: int = int(0)
        self._f6: int = int(0)
        self._f8: int = int(0)
        self._f10: str = _dafny.CodePoint('D')
        self._f11: bool = False
        self._f12: int = int(0)
        self._f13: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f14 = f14
        (self).f15 = f15
        (self).f16 = f16

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
    def f13(self):
        return self._f13

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass


class C1:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def fm5(self, p0, p1, globalState):
        return (583) < (122)

    def m5(self, p0, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: D0 = D0.default()()
        r2: bool = False
        r3: bool = False
        d_381_v0_: str
        d_381_v0_ = _dafny.CodePoint('n')
        d_382_v2_: _dafny.Map
        d_382_v2_ = _dafny.Map({d_381_v0_: p0})
        d_383_v4_: _dafny.Seq
        def iife55_():
            coll33_ = _dafny.Map()
            compr_33_: str
            for compr_33_ in (d_382_v2_).keys.Elements:
                d_384_v1_: str = compr_33_
                if (d_384_v1_) in (d_382_v2_):
                    def iife56_():
                        coll34_ = _dafny.Map()
                        compr_34_: int
                        for compr_34_ in _dafny.IntegerRange(-472, 765):
                            d_385_v3_: int = compr_34_
                            if ((-472) <= (d_385_v3_)) and ((d_385_v3_) < (765)):
                                coll34_[(d_385_v3_) - (p0)] = True
                        return _dafny.Map(coll34_)
                    coll33_[d_384_v1_] = len(iife56_()
                    )
            return _dafny.Map(coll33_)
        d_383_v4_ = _dafny.SeqWithoutIsStrInference([((D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecidbhho")), p0, (0) - (p0), d_381_v0_)).cf4) in (iife55_()
        )])
        d_386_i0_: int
        d_386_i0_ = 0
        with _dafny.label("3"):
            while (d_383_v4_)[default__.safeIndex(p0, len(d_383_v4_))]:
                with _dafny.c_label("3"):
                    if (d_386_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_386_i0_ = (d_386_i0_) + (1)
                    d_387_v5_: _dafny.Array
                    nw65_ = _dafny.Array(False, 15)
                    d_387_v5_ = nw65_
                    d_388_v6_: bool
                    d_388_v6_ = True
                    index50_ = default__.safeIndex(736, (d_387_v5_).length(0))
                    (d_387_v5_)[index50_] = d_388_v6_
                    index51_ = default__.safeIndex(736, (d_387_v5_).length(0))
                    (d_387_v5_)[index51_] = d_388_v6_
                    if (d_387_v5_)[default__.safeIndex(736, (d_387_v5_).length(0))]:
                        index52_ = default__.safeIndex(736, (d_387_v5_).length(0))
                        rhs46_ = p0
                        rhs47_ = (d_387_v5_)[default__.safeIndex(736, (d_387_v5_).length(0))]
                        rhs48_ = d_388_v6_
                        rhs49_ = d_387_v5_
                        lhs39_ = globalState
                        lhs40_ = d_387_v5_
                        lhs41_ = default__.safeIndex(736, (d_387_v5_).length(0))
                        lhs39_.f14 = rhs46_
                        r2 = rhs47_
                        lhs40_[lhs41_] = rhs48_
                        d_387_v5_ = rhs49_
                        d_389_v7_: _dafny.Seq
                        d_389_v7_ = _dafny.SeqWithoutIsStrInference([p0])
                        (globalState).f9 = (self).fm5(p0, (d_389_v7_) + (_dafny.SeqWithoutIsStrInference([-71 for d_390_i1_ in range(default__.abs(993))])), globalState)
                        d_391_v8_: _dafny.MultiSet
                        d_391_v8_ = _dafny.MultiSet([p0])
                        d_392_v9_: D1
                        d_392_v9_ = D1_DC3(d_391_v8_)
                        (globalState).f14 = ((d_392_v9_).cf8).cardinality
                        (globalState).f14 = p0
                        d_388_v6_ = d_388_v6_
                    elif True:
                        d_393_v10_: _dafny.Seq
                        d_393_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lylpu"))
                        d_394_v11_: _dafny.Map
                        d_394_v11_ = _dafny.Map({d_393_v10_: d_393_v10_})
                        d_395_v12_: _dafny.MultiSet
                        d_395_v12_ = _dafny.MultiSet([p0, p0])
                        d_396_v13_: _dafny.Array
                        nw66_ = _dafny.Array(None, 2)
                        nw66_[int(0)] = ((d_395_v12_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kx")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kx")))) in (d_395_v12_) else p0)
                        nw66_[int(1)] = p0
                        d_396_v13_ = nw66_
                        d_397_v14_: D0
                        d_397_v14_ = D0_DC2((d_387_v5_)[default__.safeIndex(736, (d_387_v5_).length(0))], len(d_394_v11_), d_396_v13_)
                        pat_let_tv7_ = p0
                        d_398_v15_: _dafny.MultiSet
                        def iife57_(_pat_let11_0):
                            def iife58_(d_399_dt__update__tmp_h0_):
                                def iife59_(_pat_let12_0):
                                    def iife60_(d_400_dt__update_hcf6_h0_):
                                        return D0_DC2((d_399_dt__update__tmp_h0_).cf5, d_400_dt__update_hcf6_h0_, (d_399_dt__update__tmp_h0_).cf7)
                                    return iife60_(_pat_let12_0)
                                return iife59_(pat_let_tv7_)
                            return iife58_(_pat_let11_0)
                        d_398_v15_ = _dafny.MultiSet([iife57_(d_397_v14_)])
                        d_398_v15_ = ((_dafny.MultiSet([d_397_v14_, D0_DC2((d_387_v5_)[default__.safeIndex(736, (d_387_v5_).length(0))], p0, d_396_v13_), d_397_v14_])) - (_dafny.MultiSet([D0_DC2((d_387_v5_)[default__.safeIndex(736, (d_387_v5_).length(0))], (0) - (p0), d_396_v13_), d_397_v14_])) if (d_388_v6_) or ((d_387_v5_)[default__.safeIndex(736, (d_387_v5_).length(0))]) else d_398_v15_)
                        d_401_v16_: _dafny.Array
                        nw67_ = _dafny.Array(_dafny.Map({}), 2)
                        d_401_v16_ = nw67_
                        d_402_v17_: _dafny.Map
                        d_402_v17_ = _dafny.Map({140: (d_387_v5_)[default__.safeIndex(736, (d_387_v5_).length(0))]})
                        index53_ = default__.safeIndex(782, (d_401_v16_).length(0))
                        (d_401_v16_)[index53_] = d_402_v17_
                        d_403_v18_: _dafny.Seq
                        d_403_v18_ = _dafny.SeqWithoutIsStrInference([d_393_v10_])
                        index54_ = default__.safeIndex(782, (d_401_v16_).length(0))
                        (d_401_v16_)[index54_] = (d_402_v17_).set((169 if d_388_v6_ else len(d_403_v18_)), False)
                        d_404_v19_: _dafny.Seq
                        d_404_v19_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                        d_405_v20_: _dafny.Map
                        d_405_v20_ = _dafny.Map({_dafny.Map({len(d_404_v19_): p0}): (d_387_v5_)[default__.safeIndex(736, (d_387_v5_).length(0))]})
                        (globalState).f14 = len((default__.fm6((d_387_v5_)[default__.safeIndex(736, (d_387_v5_).length(0))], ((d_402_v17_)[p0] if (p0) in (d_402_v17_) else d_388_v6_), p0, globalState)) | (d_405_v20_))
                        rhs50_ = d_393_v10_
                        rhs51_ = d_388_v6_
                        lhs42_ = globalState
                        r0 = rhs50_
                        lhs42_.f3 = rhs51_
                        d_406_v21_: _dafny.Map
                        d_406_v21_ = _dafny.Map({False: (p0) <= (p0)})
                        d_406_v21_ = (d_406_v21_).set(d_388_v6_, (d_387_v5_)[default__.safeIndex(736, (d_387_v5_).length(0))])
                    d_387_v5_ = (d_387_v5_ if (d_387_v5_)[default__.safeIndex(736, (d_387_v5_).length(0))] else d_387_v5_)
                    d_407_v22_: _dafny.Set
                    d_407_v22_ = _dafny.Set({d_388_v6_})
                    d_407_v22_ = (d_407_v22_) | ((_dafny.Set({d_388_v6_, (d_387_v5_)[default__.safeIndex(736, (d_387_v5_).length(0))]})).intersection(d_407_v22_))
                    pass
            pass
        d_408_v23_: _dafny.Seq
        d_408_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmpkw"))
        d_409_v24_: _dafny.Seq
        d_409_v24_ = _dafny.SeqWithoutIsStrInference([p0])
        d_410_v25_: _dafny.MultiSet
        d_410_v25_ = _dafny.MultiSet([default__.safeModuloInt(len(d_408_v23_), len(d_409_v24_)), p0])
        d_411_v26_: bool
        d_411_v26_ = False
        d_412_v27_: _dafny.Map
        d_412_v27_ = _dafny.Map({_dafny.Map({p0: d_411_v26_}): (d_410_v25_).cardinality})
        d_413_v28_: _dafny.Map
        d_413_v28_ = _dafny.Map({d_411_v26_: d_411_v26_})
        d_414_v29_: _dafny.MultiSet
        d_414_v29_ = _dafny.MultiSet([d_413_v28_, d_413_v28_])
        (globalState).f14 = ((d_410_v25_)[(len(d_412_v27_)) + (((d_414_v29_)[d_413_v28_] if (d_413_v28_) in (d_414_v29_) else p0))] if ((len(d_412_v27_)) + (((d_414_v29_)[d_413_v28_] if (d_413_v28_) in (d_414_v29_) else p0))) in (d_410_v25_) else 162)
        hi1_ = (0) - (p0)
        for d_415_i2_ in range(p0, hi1_):
            d_416_v30_: _dafny.Map
            d_416_v30_ = _dafny.Map({d_411_v26_: d_415_i2_})
            d_417_v31_: D0
            d_417_v31_ = D0_DC0(d_411_v26_)
            d_418_v32_: _dafny.Set
            d_418_v32_ = _dafny.Set({d_411_v26_, d_411_v26_})
            d_419_v33_: D0
            d_419_v33_ = D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fegnc")), default__.fm7(default__.fm7(d_415_i2_, _dafny.Map({d_411_v26_: d_415_i2_}), d_417_v31_, globalState), d_416_v30_, d_417_v31_, globalState), d_415_i2_, d_381_v0_)
            d_420_v34_: _dafny.Map
            d_420_v34_ = _dafny.Map({p0: d_381_v0_})
            d_421_v35_: _dafny.Array
            def lambda26_(d_422_i2_):
                def lambda27_(d_423_i4_):
                    return default__.safeModuloInt(d_423_i4_, d_422_i2_)

                return lambda27_

            init14_ = lambda26_(d_415_i2_)
            nw68_ = _dafny.Array(None, 28)
            for i0_14_ in range(nw68_.length(0)):
                nw68_[i0_14_] = init14_(i0_14_)
            d_421_v35_ = nw68_
            d_424_v36_: D0
            d_424_v36_ = D0_DC2(d_411_v26_, d_415_i2_, d_421_v35_)
            pat_let_tv8_ = d_411_v26_
            d_425_v37_: _dafny.Array
            nw69_ = _dafny.Array(None, 25)
            nw69_[int(0)] = (d_415_i2_) + (d_415_i2_)
            def iife61_(_pat_let13_0):
                def iife62_(d_426_dt__update__tmp_h1_):
                    def iife63_(_pat_let14_0):
                        def iife64_(d_427_dt__update_hcf0_h0_):
                            return D0_DC0(d_427_dt__update_hcf0_h0_)
                        return iife64_(_pat_let14_0)
                    return iife63_(pat_let_tv8_)
                return iife62_(_pat_let13_0)
            nw69_[int(1)] = default__.fm7(d_415_i2_, d_416_v30_, iife61_(d_417_v31_), globalState)
            nw69_[int(2)] = p0
            nw69_[int(3)] = default__.safeDivisionInt(len(_dafny.Map({(self).fm5(len(d_418_v32_), d_409_v24_, globalState): (self).fm5((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(d_408_v23_)]) for d_428_i3_ in range(default__.abs(-495))]))).cardinality, d_409_v24_, globalState)})), d_415_i2_)
            nw69_[int(4)] = len(d_413_v28_)
            nw69_[int(5)] = len(d_416_v30_)
            nw69_[int(6)] = (p0) - (p0)
            nw69_[int(7)] = default__.safeModuloInt(669, d_415_i2_)
            nw69_[int(8)] = p0
            nw69_[int(9)] = (0) - ((p0 if d_411_v26_ else p0))
            nw69_[int(10)] = (d_419_v33_).cf2
            nw69_[int(11)] = (len(default__.fm8(d_411_v26_, d_411_v26_, ((d_420_v34_)[len(d_408_v23_)] if (len(d_408_v23_)) in (d_420_v34_) else d_381_v0_), globalState))) * (len(d_383_v4_))
            nw69_[int(12)] = (0) - ((d_424_v36_).cf6)
            nw69_[int(13)] = d_415_i2_
            nw69_[int(14)] = 193
            nw69_[int(15)] = d_415_i2_
            nw69_[int(16)] = (0) - (p0)
            nw69_[int(17)] = d_415_i2_
            nw69_[int(18)] = (d_415_i2_) * (p0)
            nw69_[int(19)] = 552
            nw69_[int(20)] = p0
            nw69_[int(21)] = default__.safeDivisionInt(p0, len(d_408_v23_))
            nw69_[int(22)] = len((d_383_v4_) + (d_383_v4_))
            nw69_[int(23)] = d_415_i2_
            nw69_[int(24)] = p0
            d_425_v37_ = nw69_
            d_429_v38_: _dafny.MultiSet
            d_429_v38_ = _dafny.MultiSet([d_411_v26_])
            d_430_v39_: D2
            d_430_v39_ = D2_DC5(d_429_v38_)
            index55_ = default__.safeIndex(750, (d_425_v37_).length(0))
            (d_425_v37_)[index55_] = default__.fm7(default__.fm7(len(d_416_v30_), d_416_v30_, d_417_v31_, globalState), _dafny.Map({d_411_v26_: ((d_430_v39_).cf12).cardinality}), d_417_v31_, globalState)
            d_431_v40_: _dafny.Set
            d_431_v40_ = _dafny.Set({p0, p0})
            d_432_v41_: _dafny.Map
            d_432_v41_ = _dafny.Map({len(d_431_v40_): len(default__.fm9(d_381_v0_, globalState))})
            index56_ = default__.safeIndex(750, (d_425_v37_).length(0))
            (d_425_v37_)[index56_] = (0) - ((d_415_i2_) - ((0) - (((d_432_v41_)[813] if (813) in (d_432_v41_) else 394))))
            r3 = not((self).fm5((d_425_v37_)[default__.safeIndex(750, (d_425_v37_).length(0))], _dafny.SeqWithoutIsStrInference([d_415_i2_]), globalState))
            (globalState).f14 = (0) - (p0)
            d_433_v42_: C0
            nw70_ = C0()
            nw70_.ctor__()
            d_433_v42_ = nw70_
        d_434_v43_: _dafny.Seq
        d_434_v43_ = _dafny.SeqWithoutIsStrInference([(default__.fm10(globalState)) + (d_408_v23_)])
        rhs52_ = d_411_v26_
        rhs53_ = default__.safeDivisionInt((d_409_v24_)[default__.safeIndex(p0, len(d_409_v24_))], p0)
        rhs54_ = ((d_410_v25_).intersection((d_410_v25_).set(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ug"))), default__.abs(p0)))) | (d_410_v25_)
        rhs55_ = d_434_v43_
        lhs43_ = globalState
        lhs44_ = globalState
        lhs45_ = globalState
        lhs43_.f9 = rhs52_
        lhs44_.f14 = rhs53_
        lhs45_.f15 = rhs54_
        d_434_v43_ = rhs55_
        d_435_v44_: _dafny.Array
        nw71_ = _dafny.Array(False, 2)
        d_435_v44_ = nw71_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_435_v44_).length(0)):
            d_436_i5_: int = guard_loop_1_
            if (True) and (((0) <= (d_436_i5_)) and ((d_436_i5_) < ((d_435_v44_).length(0)))):
                (d_435_v44_)[(d_436_i5_)] = ((p0) - ((0) - (p0))) <= (len(_dafny.Set({d_411_v26_, d_411_v26_, d_411_v26_, d_411_v26_})))
        if (d_383_v4_)[default__.safeIndex(51, len(d_383_v4_))]:
            d_411_v26_ = (d_410_v25_).issubset((_dafny.MultiSet([p0])).set(p0, default__.abs(len(d_383_v4_))))
            if not((d_411_v26_) in ((d_383_v4_) + (d_383_v4_))):
                d_437_v45_: _dafny.Map
                d_437_v45_ = _dafny.Map({len(d_409_v24_): d_411_v26_})
                d_438_v46_: D0
                d_438_v46_ = D0_DC0(((d_437_v45_)[p0] if (p0) in (d_437_v45_) else (self).fm5(p0, d_409_v24_, globalState)))
                d_439_v47_: _dafny.Map
                d_439_v47_ = _dafny.Map({p0: default__.fm7(p0, _dafny.Map({d_411_v26_: p0}), d_438_v46_, globalState)})
                d_439_v47_ = d_439_v47_
                (globalState).f3 = d_411_v26_
                d_440_v48_: _dafny.Map
                d_440_v48_ = _dafny.Map({d_411_v26_: p0})
                (globalState).f14 = default__.safeModuloInt(default__.safeModuloInt(p0, len(d_440_v48_)), p0)
                d_441_v49_: _dafny.Array
                nw72_ = _dafny.Array(int(0), 27)
                d_441_v49_ = nw72_
                index57_ = default__.safeIndex(28, (d_441_v49_).length(0))
                (d_441_v49_)[index57_] = p0
                index58_ = default__.safeIndex(28, (d_441_v49_).length(0))
                (d_441_v49_)[index58_] = p0
                index59_ = default__.safeIndex(771, (d_435_v44_).length(0))
                (d_435_v44_)[index59_] = False
                index60_ = default__.safeIndex(771, (d_435_v44_).length(0))
                (d_435_v44_)[index60_] = d_411_v26_
            elif True:
                (globalState).f3 = d_411_v26_
                d_442_v50_: _dafny.Map
                d_442_v50_ = _dafny.Map({p0: _dafny.Map({_dafny.Map({p0: not(d_411_v26_)}): p0})})
                d_443_v51_: D0
                d_443_v51_ = D0_DC0(d_411_v26_)
                d_444_v53_: _dafny.Map
                d_444_v53_ = _dafny.Map({p0: d_411_v26_})
                d_445_v54_: _dafny.Seq
                d_445_v54_ = _dafny.SeqWithoutIsStrInference([d_444_v53_])
                def iife65_():
                    coll35_ = _dafny.Map()
                    compr_35_: _dafny.Map
                    for compr_35_ in ((d_445_v54_).set(default__.safeIndex(-468, len(d_445_v54_)), d_444_v53_)).Elements:
                        d_446_v52_: _dafny.Map = compr_35_
                        if (d_446_v52_) in ((d_445_v54_).set(default__.safeIndex(-468, len(d_445_v54_)), d_444_v53_)):
                            coll35_[d_446_v52_] = p0
                    return _dafny.Map(coll35_)
                d_412_v27_ = (((d_442_v50_)[p0] if (p0) in (d_442_v50_) else d_412_v27_)) | ((iife65_()
                 if (d_443_v51_).cf0 else d_412_v27_))
                d_447_v55_: _dafny.Map
                d_447_v55_ = _dafny.Map({d_411_v26_: (d_409_v24_)[default__.safeIndex(p0, len(d_409_v24_))]})
                d_448_v56_: _dafny.Array
                nw73_ = _dafny.Array(None, 9)
                nw73_[int(0)] = 122
                nw73_[int(1)] = p0
                nw73_[int(2)] = len(d_383_v4_)
                nw73_[int(3)] = (930) - (p0)
                nw73_[int(4)] = default__.fm7(p0, d_447_v55_, d_443_v51_, globalState)
                nw73_[int(5)] = p0
                nw73_[int(6)] = (p0) * (p0)
                nw73_[int(7)] = 693
                nw73_[int(8)] = 183
                d_448_v56_ = nw73_
                index61_ = default__.safeIndex(323, (d_448_v56_).length(0))
                (d_448_v56_)[index61_] = p0
                index62_ = default__.safeIndex(323, (d_448_v56_).length(0))
                (d_448_v56_)[index62_] = default__.safeModuloInt(p0, p0)
                d_449_v57_: _dafny.Set
                d_449_v57_ = _dafny.Set({((d_447_v55_)[d_411_v26_] if (d_411_v26_) in (d_447_v55_) else p0)})
                d_449_v57_ = d_449_v57_
                d_450_v58_: C0
                nw74_ = C0()
                nw74_.ctor__()
                d_450_v58_ = nw74_
            d_451_v59_: _dafny.Map
            d_451_v59_ = _dafny.Map({d_435_v44_: d_408_v23_})
            d_451_v59_ = ((d_451_v59_) | (d_451_v59_)) | (d_451_v59_)
            (globalState).f3 = (self).fm5(655, d_409_v24_, globalState)
            (globalState).f14 = p0
        elif True:
            d_452_v60_: C0
            nw75_ = C0()
            nw75_.ctor__()
            d_452_v60_ = nw75_
            (globalState).f9 = (d_408_v23_) == ((d_408_v23_).set(default__.safeIndex(735, len(d_408_v23_)), d_381_v0_))
            d_453_v61_: _dafny.MultiSet
            d_453_v61_ = _dafny.MultiSet([d_411_v26_])
            d_454_v62_: D2
            d_454_v62_ = D2_DC5(d_453_v61_)
            d_455_v63_: D1
            d_455_v63_ = D1_DC4(d_411_v26_, default__.fm11(p0, d_411_v26_, d_454_v62_, globalState), p0)
            d_456_v64_: _dafny.Map
            d_456_v64_ = _dafny.Map({p0: (self).fm5(604, d_409_v24_, globalState)})
            d_457_v65_: _dafny.Set
            d_457_v65_ = _dafny.Set({d_456_v64_})
            d_458_v66_: _dafny.Set
            d_458_v66_ = _dafny.Set({d_455_v63_, d_455_v63_, d_455_v63_, D1_DC4(False, d_457_v65_, p0), d_455_v63_})
            d_459_v67_: _dafny.Map
            d_459_v67_ = _dafny.Map({d_411_v26_: d_458_v66_})
            rhs56_ = d_459_v67_
            rhs57_ = (p0) * (p0)
            rhs58_ = not(False)
            rhs59_ = p0
            lhs46_ = globalState
            lhs47_ = globalState
            d_459_v67_ = rhs56_
            lhs46_.f14 = rhs57_
            d_411_v26_ = rhs58_
            lhs47_.f14 = rhs59_
            (globalState).f14 = p0
            nw76_ = _dafny.Array(int(0), 7)
            (globalState).f7 = nw76_
        r0 = (d_408_v23_) + ((_dafny.SeqWithoutIsStrInference([d_381_v0_ for d_460_i6_ in range(default__.abs(737))])) + (d_408_v23_))
        d_461_v69_: _dafny.Map
        d_461_v69_ = _dafny.Map({d_411_v26_: p0})
        d_462_v70_: _dafny.Seq
        d_462_v70_ = _dafny.SeqWithoutIsStrInference([d_461_v69_, d_461_v69_])
        d_463_v71_: _dafny.MultiSet
        d_463_v71_ = _dafny.MultiSet([d_411_v26_])
        d_464_v72_: D0
        d_464_v72_ = D0_DC0(d_411_v26_)
        d_465_v73_: D0
        def iife66_():
            coll36_ = _dafny.Map()
            compr_36_: _dafny.Map
            for compr_36_ in (d_462_v70_).Elements:
                d_466_v68_: _dafny.Map = compr_36_
                if (d_466_v68_) in (d_462_v70_):
                    coll36_[d_466_v68_] = d_411_v26_
            return _dafny.Map(coll36_)
        d_465_v73_ = D0_DC1(d_408_v23_, len((iife66_()
).set(_dafny.Map({d_411_v26_: len(_dafny.Map({p0: p0}))}), not(True))), (d_463_v71_).cardinality, default__.fm12(d_464_v72_, globalState))
        r1 = d_465_v73_
        d_467_v74_: _dafny.Map
        d_467_v74_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([_dafny.Set({-428}) for d_468_i8_ in range(default__.abs(857))])})
        d_469_v75_: _dafny.Set
        d_469_v75_ = _dafny.Set({-46, p0})
        d_470_v76_: _dafny.Seq
        d_470_v76_ = _dafny.SeqWithoutIsStrInference([d_469_v75_, d_469_v75_, d_469_v75_])
        r2 = (_dafny.SeqWithoutIsStrInference([_dafny.Set({p0, len(d_383_v4_)}) for d_471_i7_ in range(default__.abs(966))])) <= (((d_467_v74_)[p0] if (p0) in (d_467_v74_) else d_470_v76_))
        r3 = d_411_v26_
        return r0, r1, r2, r3


class C2(T0):
    def  __init__(self):
        self._f17: _dafny.Set = _dafny.Set({})
        self._f18: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f17(self):
        return self._f17
    @f17.setter
    def f17(self, value):
        self._f17 = value
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f17, f18):
        (self).f17 = f17
        (self)._f18 = f18

    def fm19(self, p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f18}), _dafny.Set({(self).f18, (self).f18})])) + ((_dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f18, (self).f18, (self).f18}), _dafny.Set({(self).f18})]) if (self).f18 else _dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f18}), _dafny.Set({(self).f18}), _dafny.Set({not(False)})])))

    def fm20(self, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxunet"))


class C3(T1, T0):
    def  __init__(self):
        self._f17: _dafny.Set = _dafny.Set({})
        self._f19: _dafny.Array = _dafny.Array(None, 0)
        self._f18: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f17(self):
        return self._f17
    @f17.setter
    def f17(self, value):
        self._f17 = value
    @property
    def f19(self):
        return self._f19
    @f19.setter
    def f19(self, value):
        self._f19 = value
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f19, f17, f18):
        (self).f19 = f19
        (self).f17 = f17
        (self)._f18 = f18

    def fm0(self, p0, p1, globalState):
        return len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ooi")))

    def m0(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_472_v0_: C1
        nw77_ = C1()
        nw77_.ctor__()
        d_472_v0_ = nw77_
        (globalState).f9 = (p1 if not (p1) or (p1) else (self).f18)
        d_473_v1_: _dafny.Map
        d_473_v1_ = _dafny.Map({p2: True})
        d_474_v2_: _dafny.Set
        d_474_v2_ = _dafny.Set({d_473_v1_})
        d_475_v3_: D1
        d_475_v3_ = D1_DC4(True, d_474_v2_, 164)
        d_476_v4_: _dafny.MultiSet
        d_476_v4_ = _dafny.MultiSet([d_475_v3_])
        d_477_v5_: _dafny.Seq
        d_477_v5_ = _dafny.SeqWithoutIsStrInference([D1_DC4(p1, d_474_v2_, p2), d_475_v3_])
        d_478_v6_: _dafny.Seq
        d_478_v6_ = _dafny.SeqWithoutIsStrInference([p1, p1, (self).f18])
        d_479_v7_: _dafny.Map
        d_479_v7_ = _dafny.Map({(self).f18: _dafny.MultiSet([p0, len(d_478_v6_)])})
        d_480_v8_: _dafny.MultiSet
        d_480_v8_ = _dafny.MultiSet([p2, p2])
        d_481_v9_: _dafny.Map
        d_481_v9_ = _dafny.Map({(_dafny.MultiSet(d_477_v5_)).issubset(d_476_v4_): (((d_479_v7_)[p1] if (p1) in (d_479_v7_) else d_480_v8_)) | (d_480_v8_)})
        d_481_v9_ = d_481_v9_
        d_482_v10_: _dafny.Seq
        d_482_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
        d_483_v11_: _dafny.Array
        nw78_ = _dafny.Array(int(0), 17)
        d_483_v11_ = nw78_
        d_484_v12_: D0
        d_484_v12_ = D0_DC2(p1, len(d_482_v10_), d_483_v11_)
        d_485_v13_: str
        d_485_v13_ = _dafny.CodePoint('v')
        d_486_v14_: _dafny.Seq
        d_486_v14_ = _dafny.SeqWithoutIsStrInference([(self).fm0((0) - (p2), d_485_v13_, globalState)])
        d_487_v15_: D3
        d_487_v15_ = D3_DC9()
        d_488_v16_: _dafny.Seq
        d_488_v16_ = _dafny.SeqWithoutIsStrInference([d_487_v15_])
        d_489_v17_: _dafny.Seq
        d_489_v17_ = d_488_v16_
        d_490_v18_: D5
        d_490_v18_ = D5_DC13(p1, (d_486_v14_)[default__.safeIndex(p2, len(d_486_v14_))], d_489_v17_)
        if ((d_484_v12_ if p1 else D0_DC2(p1, (d_490_v18_).cf24, d_483_v11_))).cf5:
            r0 = (self).f18
            d_491_v19_: C0
            nw79_ = C0()
            nw79_.ctor__()
            d_491_v19_ = nw79_
            (globalState).f14 = p0
            d_492_v20_: _dafny.Map
            d_492_v20_ = _dafny.Map({not ((self).f18) or (False): (0) - (p2)})
            d_492_v20_ = d_492_v20_
            if ((0) - (p0)) != ((p0) + (len(d_482_v10_))):
                (globalState).f3 = not (p1) or ((p1) or ((self).f18))
                d_493_v21_: _dafny.Set
                d_493_v21_ = _dafny.Set({(self).f18, False, p1, p1, (self).f18})
                d_493_v21_ = (D7_DC16(d_493_v21_)).cf30
                d_494_v22_: _dafny.Array
                nw80_ = _dafny.Array(_dafny.Seq({}), 20)
                d_494_v22_ = nw80_
                d_494_v22_ = (d_494_v22_ if (self).f18 else d_494_v22_)
                d_495_v23_: _dafny.Array
                nw81_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_495_v23_ = nw81_
                d_495_v23_ = d_495_v23_
                d_496_v24_: _dafny.Map
                d_496_v24_ = _dafny.Map({(p2) >= (len(_dafny.Map({p1: p0}))): d_482_v10_})
                d_496_v24_ = (d_496_v24_).set(False, d_482_v10_)
            elif True:
                d_497_v25_: _dafny.Array
                def lambda28_(d_498_v6_):
                    def lambda29_(d_499_i0_):
                        return d_498_v6_

                    return lambda29_

                init15_ = lambda28_(d_478_v6_)
                nw82_ = _dafny.Array(None, 15)
                for i0_15_ in range(nw82_.length(0)):
                    nw82_[i0_15_] = init15_(i0_15_)
                d_497_v25_ = nw82_
                index63_ = default__.safeIndex(34, (d_497_v25_).length(0))
                (d_497_v25_)[index63_] = (d_478_v6_).set(default__.safeIndex(p2, len(d_478_v6_)), p1)
                index64_ = default__.safeIndex(34, (d_497_v25_).length(0))
                (d_497_v25_)[index64_] = d_478_v6_
                rhs60_ = False
                rhs61_ = (default__.safeModuloInt(p0, (0) - (p0))) - (p2)
                lhs48_ = globalState
                r0 = rhs60_
                lhs48_.f14 = rhs61_
                d_500_v26_: _dafny.Array
                nw83_ = _dafny.Array(_dafny.CodePoint('D'), 8)
                d_500_v26_ = nw83_
                index65_ = default__.safeIndex(634, (d_500_v26_).length(0))
                (d_500_v26_)[index65_] = d_485_v13_
                d_501_v27_: D3
                d_501_v27_ = D3_DC10(d_482_v10_, p1, d_482_v10_, (d_482_v10_) + (d_482_v10_), p2)
                index66_ = default__.safeIndex(634, (d_500_v26_).length(0))
                rhs62_ = p0
                rhs63_ = d_485_v13_
                rhs64_ = default__.safeDivisionInt((p2) - (p0), p0)
                rhs65_ = _dafny.MultiSet([p2])
                rhs66_ = d_501_v27_
                lhs49_ = globalState
                lhs50_ = d_500_v26_
                lhs51_ = default__.safeIndex(634, (d_500_v26_).length(0))
                lhs52_ = globalState
                lhs53_ = globalState
                lhs49_.f14 = rhs62_
                lhs50_[lhs51_] = rhs63_
                lhs52_.f14 = rhs64_
                lhs53_.f15 = rhs65_
                d_501_v27_ = rhs66_
                d_502_v28_: _dafny.MultiSet
                d_502_v28_ = _dafny.MultiSet([p1])
                d_503_v29_: _dafny.Map
                d_503_v29_ = _dafny.Map({(0) - (len(_dafny.Set({True, (self).f18, (self).f18}))): d_502_v28_})
                d_503_v29_ = (d_503_v29_).set(len(d_482_v10_), (d_502_v28_) - (d_502_v28_))
                d_504_v30_: _dafny.Map
                d_504_v30_ = _dafny.Map({p0: 532})
                rhs67_ = not(((p0 if p1 else p2)) in (d_504_v30_))
                rhs68_ = self.f19
                lhs54_ = globalState
                lhs55_ = self
                lhs54_.f9 = rhs67_
                lhs55_.f19 = rhs68_
        elif True:
            d_482_v10_ = d_482_v10_
            (globalState).f14 = p0
            d_485_v13_ = d_485_v13_
            d_505_v31_: _dafny.Map
            d_505_v31_ = _dafny.Map({p2: p2})
            d_506_v32_: _dafny.Seq
            d_507_v33_: D0
            d_508_v34_: bool
            d_509_v35_: bool
            out33_: _dafny.Seq
            out34_: D0
            out35_: bool
            out36_: bool
            out33_, out34_, out35_, out36_ = (d_472_v0_).m5(((d_505_v31_)[len(d_482_v10_)] if (len(d_482_v10_)) in (d_505_v31_) else p0), globalState)
            d_506_v32_ = out33_
            d_507_v33_ = out34_
            d_508_v34_ = out35_
            d_509_v35_ = out36_
            (globalState).f3 = (p0) != (p2)
        if False:
            (globalState).f3 = p1
            (globalState).f3 = (self).f18
            arr0_ = self.f19
            index67_ = default__.safeIndex(466, (self.f19).length(0))
            arr0_[index67_] = p1
            d_510_v36_: _dafny.Map
            d_510_v36_ = _dafny.Map({p2: _dafny.MultiSet([d_483_v11_, d_483_v11_, d_483_v11_])})
            d_511_v37_: _dafny.MultiSet
            d_511_v37_ = _dafny.MultiSet([d_483_v11_, d_483_v11_])
            arr1_ = self.f19
            index68_ = default__.safeIndex(466, (self.f19).length(0))
            arr1_[index68_] = ((self).f18 if (p1) or (False) else ((0) - ((((d_510_v36_)[p2] if (p2) in (d_510_v36_) else d_511_v37_)).cardinality)) != (558))
            d_512_v38_: _dafny.Map
            d_512_v38_ = _dafny.Map({len(d_482_v10_): (0) - ((0) - (p0))})
            d_513_v39_: _dafny.MultiSet
            d_513_v39_ = _dafny.MultiSet([p1])
            d_514_v40_: _dafny.Array
            nw84_ = _dafny.Array(None, 22)
            nw84_[int(0)] = True
            nw84_[int(1)] = False
            nw84_[int(2)] = (self.f19)[default__.safeIndex(466, (self.f19).length(0))]
            nw84_[int(3)] = (self.f19)[default__.safeIndex(466, (self.f19).length(0))]
            nw84_[int(4)] = (self).f18
            nw84_[int(5)] = p1
            nw84_[int(6)] = (self).f18
            nw84_[int(7)] = (self).f18
            nw84_[int(8)] = (self).f18
            nw84_[int(9)] = (self).f18
            nw84_[int(10)] = (self.f19)[default__.safeIndex(466, (self.f19).length(0))]
            nw84_[int(11)] = (self.f19)[default__.safeIndex(466, (self.f19).length(0))]
            nw84_[int(12)] = not((self.f19)[default__.safeIndex(466, (self.f19).length(0))])
            nw84_[int(13)] = (self).f18
            nw84_[int(14)] = (self).f18
            nw84_[int(15)] = (self.f19)[default__.safeIndex(466, (self.f19).length(0))]
            nw84_[int(16)] = (self).f18
            nw84_[int(17)] = p1
            nw84_[int(18)] = (self).f18
            nw84_[int(19)] = (self.f19)[default__.safeIndex(466, (self.f19).length(0))]
            nw84_[int(20)] = p1
            nw84_[int(21)] = (self).f18
            d_514_v40_ = nw84_
            d_515_v41_: _dafny.Map
            d_515_v41_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([D3_DC10(d_482_v10_, (self).f18, d_482_v10_, d_482_v10_, len(_dafny.SeqWithoutIsStrInference([664 for d_516_i1_ in range(default__.abs(337))]))), default__.fm18(_dafny.CodePoint('d'), (self.f19)[default__.safeIndex(466, (self.f19).length(0))], p2, p1, globalState), D3_DC10(d_482_v10_, p1, d_482_v10_, d_482_v10_, p0)])})
            d_517_v42_: _dafny.Map
            d_517_v42_ = _dafny.Map({True: len(d_478_v6_)})
            d_518_v44_: _dafny.Set
            def iife67_():
                coll37_ = _dafny.Map()
                compr_37_: int
                for compr_37_ in _dafny.IntegerRange(490, 49):
                    d_519_v43_: int = compr_37_
                    if ((490) <= (d_519_v43_)) and ((d_519_v43_) < (49)):
                        coll37_[default__.safeModuloInt(d_519_v43_, p0)] = (d_475_v3_).cf9
                return _dafny.Map(coll37_)
            d_518_v44_ = _dafny.Set({((d_517_v42_)[(self.f19)[default__.safeIndex(466, (self.f19).length(0))]] if ((self.f19)[default__.safeIndex(466, (self.f19).length(0))]) in (d_517_v42_) else len(iife67_()
            )), p2})
            d_520_v45_: T0
            nw85_ = C2()
            nw85_.ctor__(self.f17, (self).f18)
            d_520_v45_ = nw85_
            d_521_v46_: _dafny.Map
            d_521_v46_ = _dafny.Map({len(d_518_v44_): d_520_v45_})
            d_522_v47_: _dafny.Map
            d_522_v47_ = _dafny.Map({_dafny.Set({(d_472_v0_).fm5(len(_dafny.SeqWithoutIsStrInference([p2 for d_523_i2_ in range(default__.abs(-349))])), d_486_v14_, globalState), True}): d_521_v46_})
            nw86_ = _dafny.Array(None, 27)
            nw86_[int(0)] = (0) - (((d_512_v38_)[p0] if (p0) in (d_512_v38_) else p0))
            nw86_[int(1)] = p2
            nw86_[int(2)] = p0
            nw86_[int(3)] = (p0) * (p2)
            nw86_[int(4)] = p0
            nw86_[int(5)] = p2
            nw86_[int(6)] = p2
            nw86_[int(7)] = p2
            nw86_[int(8)] = p0
            nw86_[int(9)] = (self).fm0(len((d_482_v10_).set(default__.safeIndex(p2, len(d_482_v10_)), d_485_v13_)), d_485_v13_, globalState)
            nw86_[int(10)] = p2
            nw86_[int(11)] = len(_dafny.Map({(d_513_v39_).cardinality: d_514_v40_}))
            nw86_[int(12)] = (self).fm0(p2, d_485_v13_, globalState)
            nw86_[int(13)] = default__.safeDivisionInt(len(d_478_v6_), p0)
            nw86_[int(14)] = (0) - ((len(d_515_v41_)) - (p0))
            nw86_[int(15)] = (self).fm0(p2, d_485_v13_, globalState)
            nw86_[int(16)] = p0
            nw86_[int(17)] = p2
            nw86_[int(18)] = len(d_522_v47_)
            nw86_[int(19)] = p2
            nw86_[int(20)] = default__.safeModuloInt(len(d_482_v10_), (self).fm0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwulprhlu"))), _dafny.CodePoint('g'), globalState))
            nw86_[int(21)] = 738
            nw86_[int(22)] = p2
            nw86_[int(23)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ytgmpt")))
            nw86_[int(24)] = p2
            nw86_[int(25)] = p2
            nw86_[int(26)] = default__.safeModuloInt(p2, len(d_478_v6_))
            (globalState).f7 = nw86_
            d_524_v48_: _dafny.Seq
            d_525_v49_: D0
            d_526_v50_: bool
            d_527_v51_: bool
            out37_: _dafny.Seq
            out38_: D0
            out39_: bool
            out40_: bool
            out37_, out38_, out39_, out40_ = (d_472_v0_).m5(p2, globalState)
            d_524_v48_ = out37_
            d_525_v49_ = out38_
            d_526_v50_ = out39_
            d_527_v51_ = out40_
        elif True:
            if (self).f18:
                d_528_v52_: C2
                nw87_ = C2()
                nw87_.ctor__(self.f17, (p1 if (self).f18 else p1))
                d_528_v52_ = nw87_
                d_529_v53_: _dafny.Map
                d_529_v53_ = _dafny.Map({(self).f18: p2})
                rhs69_ = p0
                rhs70_ = ((d_528_v52_).fm20(globalState)).set(default__.safeIndex((((d_529_v53_)[p1] if (p1) in (d_529_v53_) else p0) if (self).f18 else (0) - (p0)), len((d_528_v52_).fm20(globalState))), d_485_v13_)
                lhs56_ = globalState
                lhs56_.f14 = rhs69_
                d_482_v10_ = rhs70_
                (globalState).f3 = p1
                d_530_v54_: _dafny.Map
                d_530_v54_ = _dafny.Map({(self).f18: (self).f18})
                d_531_v55_: _dafny.Map
                d_531_v55_ = _dafny.Map({(d_472_v0_).fm5(len(d_530_v54_), d_486_v14_, globalState): False})
                d_531_v55_ = (d_530_v54_) | ((d_530_v54_).set((self).f18, p1))
                (globalState).f9 = p1
            elif True:
                r2 = p0
                d_532_v56_: D5
                d_532_v56_ = D5_DC14(p1, -226, len(d_478_v6_))
                d_533_v57_: _dafny.Map
                d_533_v57_ = _dafny.Map({p2: p0})
                d_534_v58_: _dafny.Map
                d_534_v58_ = _dafny.Map({(_dafny.Map({(self).f18: d_532_v56_})) == (_dafny.Map({(self).f18: d_532_v56_})): d_533_v57_})
                d_534_v58_ = (d_534_v58_) | (d_534_v58_)
                d_535_v59_: C1
                nw88_ = C1()
                nw88_.ctor__()
                d_535_v59_ = nw88_
                d_536_v60_: _dafny.Array
                nw89_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_536_v60_ = nw89_
                d_537_v61_: _dafny.Set
                d_537_v61_ = _dafny.Set({d_485_v13_, d_485_v13_})
                d_538_v62_: _dafny.Map
                d_538_v62_ = _dafny.Map({p1: _dafny.Set({(d_482_v10_)[default__.safeIndex(p0, len(d_482_v10_))], d_485_v13_, _dafny.CodePoint('h')})})
                d_539_v63_: _dafny.Seq
                d_539_v63_ = _dafny.SeqWithoutIsStrInference([d_537_v61_, _dafny.Set({d_485_v13_, d_485_v13_})])
                d_540_v65_: _dafny.Array
                nw90_ = _dafny.Array(None, 15)
                nw90_[int(0)] = d_537_v61_
                nw90_[int(1)] = ((d_538_v62_)[True] if (True) in (d_538_v62_) else d_537_v61_)
                nw90_[int(2)] = d_537_v61_
                nw90_[int(3)] = d_537_v61_
                nw90_[int(4)] = d_537_v61_
                nw90_[int(5)] = d_537_v61_
                nw90_[int(6)] = d_537_v61_
                nw90_[int(7)] = (d_539_v63_)[default__.safeIndex((self).fm0(p0, d_485_v13_, globalState), len(d_539_v63_))]
                def iife68_():
                    coll38_ = _dafny.Set()
                    compr_38_: str
                    for compr_38_ in (d_482_v10_).Elements:
                        d_541_v64_: str = compr_38_
                        if (d_541_v64_) in (d_482_v10_):
                            coll38_ = coll38_.union(_dafny.Set([d_541_v64_]))
                    return _dafny.Set(coll38_)
                nw90_[int(8)] = iife68_()
                
                nw90_[int(9)] = d_537_v61_
                nw90_[int(10)] = d_537_v61_
                nw90_[int(11)] = d_537_v61_
                nw90_[int(12)] = d_537_v61_
                nw90_[int(13)] = d_537_v61_
                nw90_[int(14)] = d_537_v61_
                d_540_v65_ = nw90_
                index69_ = default__.safeIndex(538, (d_536_v60_).length(0))
                (d_536_v60_)[index69_] = d_540_v65_
                index70_ = default__.safeIndex(538, (d_536_v60_).length(0))
                (d_536_v60_)[index70_] = d_540_v65_
                d_542_v66_: _dafny.Set
                d_542_v66_ = _dafny.Set({p2})
                (globalState).f9 = (p0) not in (d_542_v66_)
            index71_ = default__.safeIndex(596, (d_483_v11_).length(0))
            (d_483_v11_)[index71_] = default__.safeDivisionInt(-741, p0)
            index72_ = default__.safeIndex(596, (d_483_v11_).length(0))
            (d_483_v11_)[index72_] = p2
            d_543_v67_: _dafny.Array
            def lambda30_(d_544_v10_, d_545_v11_):
                def lambda31_(d_546_i3_):
                    return (d_544_v10_).set(default__.safeIndex((d_545_v11_)[default__.safeIndex(596, (d_545_v11_).length(0))], len(d_544_v10_)), _dafny.CodePoint('q'))

                return lambda31_

            init16_ = lambda30_(d_482_v10_, d_483_v11_)
            nw91_ = _dafny.Array(None, 28)
            for i0_16_ in range(nw91_.length(0)):
                nw91_[i0_16_] = init16_(i0_16_)
            d_543_v67_ = nw91_
            index73_ = default__.safeIndex(993, (d_543_v67_).length(0))
            (d_543_v67_)[index73_] = (d_482_v10_) + (d_482_v10_)
            index74_ = default__.safeIndex(744, (d_543_v67_).length(0))
            (d_543_v67_)[index74_] = d_482_v10_
            index75_ = default__.safeIndex(993, (d_543_v67_).length(0))
            index76_ = default__.safeIndex(596, (d_483_v11_).length(0))
            index77_ = default__.safeIndex(744, (d_543_v67_).length(0))
            rhs71_ = default__.fm21(_dafny.SeqWithoutIsStrInference([(d_483_v11_)[default__.safeIndex(596, (d_483_v11_).length(0))], p0, p0, (0) - (p0), p2]), globalState)
            rhs72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nosuekq"))
            rhs73_ = (self).fm0((d_483_v11_)[default__.safeIndex(596, (d_483_v11_).length(0))], d_485_v13_, globalState)
            rhs74_ = default__.fm22(p0, globalState)
            lhs57_ = d_543_v67_
            lhs58_ = default__.safeIndex(993, (d_543_v67_).length(0))
            lhs59_ = d_483_v11_
            lhs60_ = default__.safeIndex(596, (d_483_v11_).length(0))
            lhs61_ = d_543_v67_
            lhs62_ = default__.safeIndex(744, (d_543_v67_).length(0))
            d_473_v1_ = rhs71_
            lhs57_[lhs58_] = rhs72_
            lhs59_[lhs60_] = rhs73_
            lhs61_[lhs62_] = rhs74_
            arr2_ = self.f19
            index78_ = default__.safeIndex(65, (self.f19).length(0))
            arr2_[index78_] = (self).f18
            arr3_ = self.f19
            index79_ = default__.safeIndex(65, (self.f19).length(0))
            arr3_[index79_] = not((self).f18)
            d_547_v68_: C1
            nw92_ = C1()
            nw92_.ctor__()
            d_547_v68_ = nw92_
        d_548_v69_: D2
        d_548_v69_ = D2_DC7(default__.fm23(p0, 942, (self).f18, globalState))
        d_549_v70_: D2
        d_549_v70_ = D2_DC7(d_548_v69_)
        source4_ = d_549_v70_
        if source4_.is_DC6:
            d_550___mcc_h0_ = source4_.cf13
            d_551_cf13_ = d_550___mcc_h0_
            d_552_v71_: _dafny.Array
            nw93_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_552_v71_ = nw93_
            d_553_v72_: _dafny.Array
            nw94_ = _dafny.Array(None, 3)
            nw94_[int(0)] = d_483_v11_
            nw94_[int(1)] = d_483_v11_
            nw94_[int(2)] = d_483_v11_
            d_553_v72_ = nw94_
            index80_ = default__.safeIndex(729, (d_552_v71_).length(0))
            (d_552_v71_)[index80_] = d_553_v72_
            index81_ = default__.safeIndex(729, (d_552_v71_).length(0))
            (d_552_v71_)[index81_] = d_553_v72_
            d_554_v73_: C2
            nw95_ = C2()
            nw95_.ctor__(self.f17, (self).f18)
            d_554_v73_ = nw95_
            rhs75_ = d_551_cf13_
            rhs76_ = d_554_v73_
            lhs63_ = globalState
            lhs63_.f14 = rhs75_
            d_554_v73_ = rhs76_
            rhs77_ = p1
            rhs78_ = p2
            lhs64_ = globalState
            lhs65_ = globalState
            lhs64_.f3 = rhs77_
            lhs65_.f14 = rhs78_
            (globalState).f3 = (((self).f18) and (p1)) or ((self).f18)
        elif source4_.is_DC5:
            d_555___mcc_h1_ = source4_.cf12
            d_556_cf12_ = d_555___mcc_h1_
            d_482_v10_ = (d_482_v10_) + (d_482_v10_)
            d_557_v74_: _dafny.MultiSet
            d_557_v74_ = _dafny.MultiSet([self.f19, self.f19, self.f19, self.f19])
            d_558_v75_: _dafny.Map
            d_558_v75_ = _dafny.Map({(d_486_v14_)[default__.safeIndex(p0, len(d_486_v14_))]: p0})
            (globalState).f14 = ((d_557_v74_).cardinality) * (((d_558_v75_)[p0] if (p0) in (d_558_v75_) else p0))
            if not((self).f18):
                (globalState).f9 = (self).f18
                d_559_v76_: C0
                nw96_ = C0()
                nw96_.ctor__()
                d_559_v76_ = nw96_
                (globalState).f9 = (self).f18
                d_560_v77_: _dafny.Array
                nw97_ = _dafny.Array(None, 5)
                nw97_[int(0)] = d_472_v0_
                nw97_[int(1)] = (d_472_v0_ if p1 else d_472_v0_)
                nw97_[int(2)] = d_472_v0_
                nw97_[int(3)] = d_472_v0_
                nw97_[int(4)] = d_472_v0_
                d_560_v77_ = nw97_
                d_561_v78_: _dafny.Array
                def lambda32_(d_562_i4_):
                    return _dafny.CodePoint('l')

                init17_ = lambda32_
                nw98_ = _dafny.Array(None, 14)
                for i0_17_ in range(nw98_.length(0)):
                    nw98_[i0_17_] = init17_(i0_17_)
                d_561_v78_ = nw98_
                index82_ = default__.safeIndex(465, (d_561_v78_).length(0))
                (d_561_v78_)[index82_] = (d_485_v13_ if p1 else d_485_v13_)
                d_563_v79_: _dafny.Seq
                d_563_v79_ = _dafny.SeqWithoutIsStrInference([(d_560_v77_ if (d_472_v0_).fm5(p2, _dafny.SeqWithoutIsStrInference([p0 for d_564_i5_ in range(default__.abs(914))]), globalState) else d_560_v77_)])
                index83_ = default__.safeIndex(465, (d_561_v78_).length(0))
                rhs79_ = (d_563_v79_)[default__.safeIndex(p2, len(d_563_v79_))]
                rhs80_ = d_485_v13_
                lhs66_ = d_561_v78_
                lhs67_ = default__.safeIndex(465, (d_561_v78_).length(0))
                d_560_v77_ = rhs79_
                lhs66_[lhs67_] = rhs80_
                (globalState).f3 = (self).f18
            elif True:
                index84_ = default__.safeIndex(605, (d_483_v11_).length(0))
                (d_483_v11_)[index84_] = (p0) + (205)
                d_565_v80_: _dafny.Set
                d_565_v80_ = _dafny.Set({d_482_v10_})
                d_566_v81_: _dafny.Array
                nw99_ = _dafny.Array(None, 26)
                nw99_[int(0)] = d_482_v10_
                nw99_[int(1)] = d_482_v10_
                nw99_[int(2)] = (d_482_v10_) + (d_482_v10_)
                nw99_[int(3)] = _dafny.SeqWithoutIsStrInference([d_485_v13_ for d_567_i6_ in range(default__.abs(100))])
                nw99_[int(4)] = (d_482_v10_) + (d_482_v10_)
                nw99_[int(5)] = d_482_v10_
                nw99_[int(6)] = d_482_v10_
                nw99_[int(7)] = (d_482_v10_).set(default__.safeIndex(len(d_565_v80_), len(d_482_v10_)), d_485_v13_)
                nw99_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdic"))
                nw99_[int(9)] = d_482_v10_
                nw99_[int(10)] = (default__.fm22((d_556_cf12_).cardinality, globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fe")))
                nw99_[int(11)] = d_482_v10_
                nw99_[int(12)] = d_482_v10_
                nw99_[int(13)] = d_482_v10_
                nw99_[int(14)] = (_dafny.SeqWithoutIsStrInference([d_485_v13_ for d_568_i7_ in range(default__.abs(766))])) + (d_482_v10_)
                nw99_[int(15)] = (d_482_v10_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "su")))
                nw99_[int(16)] = d_482_v10_
                nw99_[int(17)] = _dafny.SeqWithoutIsStrInference([d_485_v13_ for d_569_i8_ in range(default__.abs(571))])
                nw99_[int(18)] = d_482_v10_
                nw99_[int(19)] = d_482_v10_
                nw99_[int(20)] = d_482_v10_
                nw99_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfxv"))
                nw99_[int(22)] = _dafny.SeqWithoutIsStrInference([d_485_v13_ for d_570_i9_ in range(default__.abs(917))])
                nw99_[int(23)] = d_482_v10_
                nw99_[int(24)] = default__.fm22(p2, globalState)
                nw99_[int(25)] = (d_482_v10_) + (d_482_v10_)
                d_566_v81_ = nw99_
                index85_ = default__.safeIndex(75, (d_566_v81_).length(0))
                (d_566_v81_)[index85_] = d_482_v10_
                index86_ = default__.safeIndex(605, (d_483_v11_).length(0))
                index87_ = default__.safeIndex(75, (d_566_v81_).length(0))
                rhs81_ = (p1) and ((self).f18)
                rhs82_ = p1
                rhs83_ = (p0) - ((394 if (self).f18 else p2))
                rhs84_ = d_482_v10_
                lhs68_ = globalState
                lhs69_ = globalState
                lhs70_ = d_483_v11_
                lhs71_ = default__.safeIndex(605, (d_483_v11_).length(0))
                lhs72_ = d_566_v81_
                lhs73_ = default__.safeIndex(75, (d_566_v81_).length(0))
                lhs68_.f3 = rhs81_
                lhs69_.f3 = rhs82_
                lhs70_[lhs71_] = rhs83_
                lhs72_[lhs73_] = rhs84_
                (globalState).f14 = len(((d_566_v81_)[default__.safeIndex(75, (d_566_v81_).length(0))]) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yh"))) + ((d_566_v81_)[default__.safeIndex(75, (d_566_v81_).length(0))])))
                d_571_v82_: C2
                nw100_ = C2()
                nw100_.ctor__(self.f17, (self).f18)
                d_571_v82_ = nw100_
                d_572_v83_: _dafny.Map
                d_572_v83_ = _dafny.Map({p1: d_571_v82_})
                d_573_v84_: _dafny.Map
                d_573_v84_ = _dafny.Map({((d_572_v83_)[False] if (False) in (d_572_v83_) else d_571_v82_): (self).f18})
                d_574_v85_: _dafny.Seq
                d_574_v85_ = _dafny.SeqWithoutIsStrInference([d_573_v84_])
                index88_ = default__.safeIndex(605, (d_483_v11_).length(0))
                (d_483_v11_)[index88_] = len(((d_574_v85_) + (_dafny.SeqWithoutIsStrInference([d_573_v84_, d_573_v84_]))) + (d_574_v85_))
                d_575_v86_: D0
                d_575_v86_ = D0_DC1((d_566_v81_)[default__.safeIndex(75, (d_566_v81_).length(0))], (d_483_v11_)[default__.safeIndex(605, (d_483_v11_).length(0))], (d_483_v11_)[default__.safeIndex(605, (d_483_v11_).length(0))], ((d_566_v81_)[default__.safeIndex(75, (d_566_v81_).length(0))])[default__.safeIndex(len((d_566_v81_)[default__.safeIndex(75, (d_566_v81_).length(0))]), len((d_566_v81_)[default__.safeIndex(75, (d_566_v81_).length(0))]))])
                d_575_v86_ = (d_575_v86_ if (self).f18 else d_575_v86_)
                arr4_ = self.f19
                index89_ = default__.safeIndex(809, (self.f19).length(0))
                arr4_[index89_] = True
                d_576_v87_: _dafny.Array
                nw101_ = _dafny.Array(int(0), 24)
                d_576_v87_ = nw101_
                d_577_v88_: _dafny.Seq
                d_577_v88_ = _dafny.SeqWithoutIsStrInference([d_576_v87_, d_576_v87_])
                d_578_v89_: _dafny.Map
                d_578_v89_ = _dafny.Map({-552: d_576_v87_})
                arr5_ = self.f19
                index90_ = default__.safeIndex(809, (self.f19).length(0))
                rhs85_ = ((d_577_v88_)[default__.safeIndex((d_483_v11_)[default__.safeIndex(605, (d_483_v11_).length(0))], len(d_577_v88_))] if ((d_566_v81_)[default__.safeIndex(75, (d_566_v81_).length(0))]) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))) else ((d_578_v89_)[(d_483_v11_)[default__.safeIndex(605, (d_483_v11_).length(0))]] if ((d_483_v11_)[default__.safeIndex(605, (d_483_v11_).length(0))]) in (d_578_v89_) else d_576_v87_))
                rhs86_ = p1
                rhs87_ = (self).f18
                rhs88_ = (p1 if (self).f18 else (self).f18)
                lhs74_ = globalState
                lhs75_ = globalState
                lhs76_ = self.f19
                lhs77_ = default__.safeIndex(809, (self.f19).length(0))
                lhs78_ = globalState
                lhs74_.f7 = rhs85_
                lhs75_.f3 = rhs86_
                lhs76_[lhs77_] = rhs87_
                lhs78_.f9 = rhs88_
            index91_ = default__.safeIndex(136, (d_483_v11_).length(0))
            (d_483_v11_)[index91_] = default__.safeDivisionInt(len(_dafny.Set({len(d_478_v6_), len(d_486_v14_)})), p0)
            index92_ = default__.safeIndex(136, (d_483_v11_).length(0))
            (d_483_v11_)[index92_] = p0
        elif True:
            d_579___mcc_h2_ = source4_.cf14
            d_580_cf14_ = d_579___mcc_h2_
            d_581_v90_: _dafny.Map
            d_581_v90_ = _dafny.Map({p1: self.f19})
            d_581_v90_ = (d_581_v90_).set(((self).f18) == ((self).f18), self.f19)
            r1 = (self).f18
            r0 = not((self).f18)
            d_582_v91_: _dafny.Map
            d_582_v91_ = _dafny.Map({d_483_v11_: d_482_v10_})
            d_582_v91_ = (d_582_v91_).set(d_483_v11_, d_482_v10_)
        d_583_v92_: _dafny.Set
        d_583_v92_ = _dafny.Set({p0})
        d_584_v93_: _dafny.Seq
        d_584_v93_ = _dafny.SeqWithoutIsStrInference([d_583_v92_])
        def iife69_():
            coll39_ = _dafny.Set()
            compr_39_: _dafny.Set
            for compr_39_ in (d_584_v93_).Elements:
                d_585_v94_: _dafny.Set = compr_39_
                if (d_585_v94_) in (d_584_v93_):
                    coll39_ = coll39_.union(_dafny.Set([d_585_v94_]))
            return _dafny.Set(coll39_)
        r0 = (not(((self).f18 if (self).f18 else (self).f18)) if (d_583_v92_) in (iife69_()
        ) else p1)
        r1 = (d_472_v0_).fm5(default__.safeDivisionInt(p2, p0), (d_486_v14_) + (_dafny.SeqWithoutIsStrInference([p0])), globalState)
        r2 = p2
        return r0, r1, r2


class C4(T0):
    def  __init__(self):
        self._f17: _dafny.Set = _dafny.Set({})
        self._f18: bool = False
        self.f26: int = int(0)
        self._f25: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f17(self):
        return self._f17
    @f17.setter
    def f17(self, value):
        self._f17 = value
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f25, f26, f17, f18):
        (self)._f25 = f25
        (self).f26 = f26
        (self).f17 = f17
        (self)._f18 = f18

    def m3(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_586_v0_: bool
        d_587_v1_: _dafny.Map
        d_588_v2_: int
        out41_: bool
        out42_: _dafny.Map
        out43_: int
        out41_, out42_, out43_ = (self).m4(globalState)
        d_586_v0_ = out41_
        d_587_v1_ = out42_
        d_588_v2_ = out43_
        d_589_v3_: _dafny.Array
        def lambda33_(d_590_p0_):
            def lambda34_(d_591_i0_):
                return d_590_p0_

            return lambda34_

        init18_ = lambda33_(p0)
        nw102_ = _dafny.Array(None, 22)
        for i0_18_ in range(nw102_.length(0)):
            nw102_[i0_18_] = init18_(i0_18_)
        d_589_v3_ = nw102_
        index93_ = default__.safeIndex(324, (d_589_v3_).length(0))
        (d_589_v3_)[index93_] = p0
        index94_ = default__.safeIndex(324, (d_589_v3_).length(0))
        (d_589_v3_)[index94_] = p0
        d_592_v4_: _dafny.Array
        def lambda35_(d_593_v2_):
            def lambda36_(d_594_i2_):
                return (d_594_i2_) + (d_593_v2_)

            return lambda36_

        init19_ = lambda35_(d_588_v2_)
        nw103_ = _dafny.Array(None, 20)
        for i0_19_ in range(nw103_.length(0)):
            nw103_[i0_19_] = init19_(i0_19_)
        d_592_v4_ = nw103_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_592_v4_).length(0)):
            d_595_i1_: int = guard_loop_2_
            if (True) and (((0) <= (d_595_i1_)) and ((d_595_i1_) < ((d_592_v4_).length(0)))):
                (d_592_v4_)[(d_595_i1_)] = (d_595_i1_) * (self.f26)
        d_596_v5_: _dafny.Seq
        d_596_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
        d_597_v6_: _dafny.Seq
        d_597_v6_ = _dafny.SeqWithoutIsStrInference([d_596_v5_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvpbsjve"))])
        d_597_v6_ = (d_597_v6_).set(default__.safeIndex(d_588_v2_, len(d_597_v6_)), (d_596_v5_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiiyumo"))))
        (self).f26 = ((self).f25) * (d_588_v2_)
        (self).f26 = self.f26
        r0 = 636
        return r0

    def m4(self, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        r2: int = int(0)
        d_598_v0_: _dafny.Array
        def lambda37_(d_599_i0_):
            return (d_599_i0_) * (len(_dafny.SeqWithoutIsStrInference([self.f26, (self).f25, (self).f25])))

        init20_ = lambda37_
        nw104_ = _dafny.Array(None, 6)
        for i0_20_ in range(nw104_.length(0)):
            nw104_[i0_20_] = init20_(i0_20_)
        d_598_v0_ = nw104_
        pat_let_tv9_ = d_598_v0_
        def iife70_(_pat_let15_0):
            def iife71_(d_600_dt__update__tmp_h0_):
                def iife72_(_pat_let16_0):
                    def iife73_(d_601_dt__update_hcf6_h0_):
                        def iife74_(_pat_let17_0):
                            def iife75_(d_602_dt__update_hcf7_h0_):
                                return D0_DC2((d_600_dt__update__tmp_h0_).cf5, d_601_dt__update_hcf6_h0_, d_602_dt__update_hcf7_h0_)
                            return iife75_(_pat_let17_0)
                        return iife74_(pat_let_tv9_)
                    return iife73_(_pat_let16_0)
                return iife72_(self.f26)
            return iife71_(_pat_let15_0)
        source5_ = iife70_(D0_DC2((self).f18, (self).f25, d_598_v0_))
        if source5_.is_DC0:
            d_603___mcc_h0_ = source5_.cf0
            d_604_cf0_ = d_603___mcc_h0_
            index95_ = default__.safeIndex(268, (d_598_v0_).length(0))
            (d_598_v0_)[index95_] = default__.safeDivisionInt((self).f25, self.f26)
            index96_ = default__.safeIndex(268, (d_598_v0_).length(0))
            (d_598_v0_)[index96_] = 646
            d_605_v1_: _dafny.Map
            d_605_v1_ = _dafny.Map({((d_598_v0_)[default__.safeIndex(268, (d_598_v0_).length(0))]) >= (self.f26): False})
            d_606_v2_: _dafny.Map
            d_606_v2_ = _dafny.Map({(self).f25: (self).f18})
            d_607_v3_: _dafny.MultiSet
            d_607_v3_ = _dafny.MultiSet([d_606_v2_])
            d_605_v1_ = default__.fm3(self.f26, True, d_607_v3_, globalState)
            source6_ = D0_DC0((self).f18)
            if source6_.is_DC0:
                d_608___mcc_h8_ = source6_.cf0
                d_609_cf0_ = d_608___mcc_h8_
                d_610_v4_: _dafny.Seq
                d_610_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnhyyx"))
                d_610_v4_ = d_610_v4_
                d_611_v5_: str
                d_611_v5_ = _dafny.CodePoint('w')
                d_611_v5_ = d_611_v5_
                index97_ = default__.safeIndex(268, (d_598_v0_).length(0))
                (d_598_v0_)[index97_] = default__.safeModuloInt(default__.fm4((d_598_v0_)[default__.safeIndex(268, (d_598_v0_).length(0))], globalState), self.f26)
                d_612_v6_: _dafny.Set
                d_612_v6_ = _dafny.Set({d_609_cf0_})
                d_613_v7_: _dafny.Array
                nw105_ = _dafny.Array(None, 5)
                nw105_[int(0)] = (d_598_v0_)[default__.safeIndex(268, (d_598_v0_).length(0))]
                nw105_[int(1)] = (d_598_v0_)[default__.safeIndex(268, (d_598_v0_).length(0))]
                nw105_[int(2)] = (self).f25
                nw105_[int(3)] = self.f26
                nw105_[int(4)] = len(d_612_v6_)
                d_613_v7_ = nw105_
                d_614_v8_: D0
                d_614_v8_ = D0_DC2(d_609_cf0_, (d_598_v0_)[default__.safeIndex(268, (d_598_v0_).length(0))], d_613_v7_)
                d_615_v9_: _dafny.MultiSet
                d_615_v9_ = _dafny.MultiSet([(d_614_v8_).cf5])
                d_616_v10_: _dafny.Set
                d_616_v10_ = _dafny.Set({(d_615_v9_).cardinality})
                d_617_v11_: _dafny.Set
                d_617_v11_ = _dafny.Set({d_616_v10_, d_616_v10_, d_616_v10_})
                d_618_v12_: _dafny.Map
                d_618_v12_ = _dafny.Map({d_616_v10_: _dafny.Set({self.f26})})
                def iife76_():
                    coll40_ = _dafny.Set()
                    compr_40_: _dafny.Set
                    for compr_40_ in (d_618_v12_).keys.Elements:
                        d_619_v13_: _dafny.Set = compr_40_
                        if (d_619_v13_) in (d_618_v12_):
                            coll40_ = coll40_.union(_dafny.Set([d_619_v13_]))
                    return _dafny.Set(coll40_)
                (globalState).f9 = ((iife76_()
                ).intersection(d_617_v11_)).ispropersubset((d_617_v11_) | (d_617_v11_))
            elif source6_.is_DC1:
                d_620___mcc_h9_ = source6_.cf1
                d_621___mcc_h10_ = source6_.cf2
                d_622___mcc_h11_ = source6_.cf3
                d_623___mcc_h12_ = source6_.cf4
                d_624_cf4_ = d_623___mcc_h12_
                d_625_cf3_ = d_622___mcc_h11_
                d_626_cf2_ = d_621___mcc_h10_
                d_627_cf1_ = d_620___mcc_h9_
                d_628_v14_: _dafny.MultiSet
                d_628_v14_ = _dafny.MultiSet([(self).f18, (self).f18, not(d_604_cf0_)])
                d_628_v14_ = d_628_v14_
                index98_ = default__.safeIndex(268, (d_598_v0_).length(0))
                rhs89_ = self.f26
                rhs90_ = self.f26
                rhs91_ = (d_598_v0_)[default__.safeIndex(268, (d_598_v0_).length(0))]
                lhs79_ = d_598_v0_
                lhs80_ = default__.safeIndex(268, (d_598_v0_).length(0))
                d_626_cf2_ = rhs89_
                lhs79_[lhs80_] = rhs90_
                d_625_cf3_ = rhs91_
                d_629_v15_: _dafny.Array
                nw106_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
                d_629_v15_ = nw106_
                index99_ = default__.safeIndex(98, (d_629_v15_).length(0))
                (d_629_v15_)[index99_] = (d_627_cf1_) + (d_627_cf1_)
                d_630_v16_: _dafny.Map
                d_630_v16_ = _dafny.Map({not(d_604_cf0_): d_624_cf4_})
                index100_ = default__.safeIndex(98, (d_629_v15_).length(0))
                (d_629_v15_)[index100_] = (d_627_cf1_).set(default__.safeIndex((d_598_v0_)[default__.safeIndex(268, (d_598_v0_).length(0))], len(d_627_cf1_)), ((d_630_v16_)[(self).f18] if ((self).f18) in (d_630_v16_) else _dafny.CodePoint('p')))
                d_631_v17_: C0
                nw107_ = C0()
                nw107_.ctor__()
                d_631_v17_ = nw107_
            elif True:
                d_632___mcc_h13_ = source6_.cf5
                d_633___mcc_h14_ = source6_.cf6
                d_634___mcc_h15_ = source6_.cf7
                d_635_cf7_ = d_634___mcc_h15_
                d_636_cf6_ = d_633___mcc_h14_
                d_637_cf5_ = d_632___mcc_h13_
                d_638_v18_: _dafny.Array
                def lambda38_(d_639_cf0_):
                    def lambda39_(d_640_i1_):
                        return (d_639_cf0_ if True else d_639_cf0_)

                    return lambda39_

                init21_ = lambda38_(d_604_cf0_)
                nw108_ = _dafny.Array(None, 27)
                for i0_21_ in range(nw108_.length(0)):
                    nw108_[i0_21_] = init21_(i0_21_)
                d_638_v18_ = nw108_
                index101_ = default__.safeIndex(260, (d_638_v18_).length(0))
                (d_638_v18_)[index101_] = False
                index102_ = default__.safeIndex(260, (d_638_v18_).length(0))
                (d_638_v18_)[index102_] = d_637_cf5_
                d_641_v19_: _dafny.Array
                def lambda40_(d_642_v1_):
                    def lambda41_(d_643_i2_):
                        return (_dafny.SeqWithoutIsStrInference([d_642_v1_])) + ((D3_DC8(_dafny.SeqWithoutIsStrInference([d_642_v1_]))).cf15)

                    return lambda41_

                init22_ = lambda40_(d_605_v1_)
                nw109_ = _dafny.Array(None, 24)
                for i0_22_ in range(nw109_.length(0)):
                    nw109_[i0_22_] = init22_(i0_22_)
                d_641_v19_ = nw109_
                d_644_v20_: _dafny.Seq
                d_644_v20_ = _dafny.SeqWithoutIsStrInference([d_605_v1_])
                index103_ = default__.safeIndex(327, (d_641_v19_).length(0))
                (d_641_v19_)[index103_] = (d_644_v20_) + (_dafny.SeqWithoutIsStrInference([d_605_v1_ for d_645_i3_ in range(default__.abs(-803))]))
                index104_ = default__.safeIndex(327, (d_641_v19_).length(0))
                (d_641_v19_)[index104_] = d_644_v20_
                d_646_v21_: D0
                d_646_v21_ = D0_DC0(d_637_cf5_)
                d_647_v22_: _dafny.Map
                d_647_v22_ = _dafny.Map({default__.fm13(globalState): d_646_v21_})
                d_647_v22_ = (d_647_v22_).set(d_604_cf0_, d_646_v21_)
                (globalState).f14 = default__.safeModuloInt(d_636_cf6_, (d_636_cf6_) + ((self).f25))
            d_648_v23_: str
            d_648_v23_ = _dafny.CodePoint('c')
            d_649_v24_: _dafny.Seq
            d_649_v24_ = _dafny.SeqWithoutIsStrInference([d_648_v23_, d_648_v23_])
            source7_ = D3_DC10((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_650_i4_ in range(default__.abs(367))])) + (d_649_v24_), d_604_cf0_, (d_649_v24_) + (d_649_v24_), d_649_v24_, default__.fm4(len(d_605_v1_), globalState))
            if source7_.is_DC9:
                (globalState).f14 = default__.safeModuloInt((d_598_v0_)[default__.safeIndex(268, (d_598_v0_).length(0))], (self).f25)
                index105_ = default__.safeIndex(268, (d_598_v0_).length(0))
                (d_598_v0_)[index105_] = ((self).f25) - (((d_598_v0_)[default__.safeIndex(268, (d_598_v0_).length(0))]) * ((d_598_v0_)[default__.safeIndex(268, (d_598_v0_).length(0))]))
                d_651_v25_: _dafny.Array
                nw110_ = _dafny.Array(False, 15)
                d_651_v25_ = nw110_
                d_652_v26_: _dafny.Map
                d_652_v26_ = _dafny.Map({d_651_v25_: d_648_v23_})
                d_652_v26_ = (d_652_v26_).set(d_651_v25_, d_648_v23_)
                (globalState).f3 = (d_648_v23_) not in (_dafny.SeqWithoutIsStrInference([d_648_v23_ for d_653_i5_ in range(default__.abs(-115))]))
            elif source7_.is_DC10:
                d_654___mcc_h16_ = source7_.cf16
                d_655___mcc_h17_ = source7_.cf17
                d_656___mcc_h18_ = source7_.cf18
                d_657___mcc_h19_ = source7_.cf19
                d_658___mcc_h20_ = source7_.cf20
                d_659_cf20_ = d_658___mcc_h20_
                d_660_cf19_ = d_657___mcc_h19_
                d_661_cf18_ = d_656___mcc_h18_
                d_662_cf17_ = d_655___mcc_h17_
                d_663_cf16_ = d_654___mcc_h16_
                (globalState).f3 = d_662_cf17_
                d_664_v27_: _dafny.Array
                def lambda42_(d_665_v0_):
                    def lambda43_(d_666_i6_):
                        return (d_666_i6_) * ((d_665_v0_)[default__.safeIndex(268, (d_665_v0_).length(0))])

                    return lambda43_

                init23_ = lambda42_(d_598_v0_)
                nw111_ = _dafny.Array(None, 6)
                for i0_23_ in range(nw111_.length(0)):
                    nw111_[i0_23_] = init23_(i0_23_)
                d_664_v27_ = nw111_
                d_667_v28_: _dafny.Set
                d_667_v28_ = _dafny.Set({d_664_v27_, d_664_v27_, d_664_v27_})
                d_668_v29_: D3
                d_668_v29_ = D3_DC9()
                d_669_v30_: _dafny.Seq
                d_669_v30_ = _dafny.SeqWithoutIsStrInference([d_668_v29_])
                d_670_v31_: _dafny.Seq
                d_670_v31_ = _dafny.SeqWithoutIsStrInference([d_668_v29_ for d_671_i8_ in range(default__.abs(316))])
                d_672_v32_: _dafny.Seq
                d_672_v32_ = _dafny.SeqWithoutIsStrInference([d_669_v30_, d_669_v30_, d_669_v30_])
                d_673_v33_: _dafny.Array
                nw112_ = _dafny.Array(None, 26)
                nw112_[int(0)] = d_669_v30_
                nw112_[int(1)] = d_669_v30_
                nw112_[int(2)] = _dafny.SeqWithoutIsStrInference([d_668_v29_])
                nw112_[int(3)] = (d_669_v30_) + (d_669_v30_)
                nw112_[int(4)] = (d_669_v30_) + (d_669_v30_)
                nw112_[int(5)] = d_669_v30_
                nw112_[int(6)] = (d_669_v30_) + (_dafny.SeqWithoutIsStrInference([d_668_v29_ for d_674_i7_ in range(default__.abs(54))]))
                nw112_[int(7)] = (d_670_v31_)
                nw112_[int(8)] = (d_669_v30_) + (_dafny.SeqWithoutIsStrInference([d_668_v29_]))
                nw112_[int(9)] = (d_672_v32_)[default__.safeIndex((self).f25, len(d_672_v32_))]
                nw112_[int(10)] = d_669_v30_
                nw112_[int(11)] = d_669_v30_
                nw112_[int(12)] = _dafny.SeqWithoutIsStrInference([d_668_v29_])
                nw112_[int(13)] = d_669_v30_
                nw112_[int(14)] = d_669_v30_
                nw112_[int(15)] = (d_669_v30_) + (d_669_v30_)
                nw112_[int(16)] = d_669_v30_
                nw112_[int(17)] = (_dafny.SeqWithoutIsStrInference([d_668_v29_ for d_675_i9_ in range(default__.abs(659))])) + (d_669_v30_)
                nw112_[int(18)] = default__.fm14(d_662_cf17_, (self).f18, -325, globalState)
                nw112_[int(19)] = d_669_v30_
                nw112_[int(20)] = _dafny.SeqWithoutIsStrInference([d_668_v29_])
                nw112_[int(21)] = (d_669_v30_) + (d_669_v30_)
                nw112_[int(22)] = (d_669_v30_) + (_dafny.SeqWithoutIsStrInference([d_668_v29_, D3_DC9(), d_668_v29_]))
                nw112_[int(23)] = d_669_v30_
                nw112_[int(24)] = _dafny.SeqWithoutIsStrInference([d_668_v29_ for d_676_i10_ in range(default__.abs(118))])
                nw112_[int(25)] = d_669_v30_
                d_673_v33_ = nw112_
                index106_ = default__.safeIndex(253, (d_673_v33_).length(0))
                (d_673_v33_)[index106_] = _dafny.SeqWithoutIsStrInference([d_668_v29_ for d_677_i11_ in range(default__.abs(394))])
                index107_ = default__.safeIndex(253, (d_673_v33_).length(0))
                rhs92_ = default__.fm13(globalState)
                rhs93_ = d_667_v28_
                rhs94_ = (d_669_v30_).set(default__.safeIndex(self.f26, len(d_669_v30_)), d_668_v29_)
                lhs81_ = d_673_v33_
                lhs82_ = default__.safeIndex(253, (d_673_v33_).length(0))
                d_604_cf0_ = rhs92_
                d_667_v28_ = rhs93_
                lhs81_[lhs82_] = rhs94_
                d_662_cf17_ = not((d_662_cf17_) and ((not(d_662_cf17_) if d_604_cf0_ else d_662_cf17_)))
                d_678_v34_: _dafny.Seq
                d_678_v34_ = _dafny.SeqWithoutIsStrInference([(self).f18])
                d_679_v35_: D3
                d_679_v35_ = D3_DC10(d_649_v24_, d_604_cf0_, d_663_cf16_, d_663_cf16_, (self).f25)
                d_680_v36_: _dafny.Map
                d_680_v36_ = _dafny.Map({(d_678_v34_)[default__.safeIndex(self.f26, len(d_678_v34_))]: d_679_v35_})
                d_681_v37_: _dafny.Map
                d_681_v37_ = _dafny.Map({len(d_680_v36_): (d_598_v0_)[default__.safeIndex(268, (d_598_v0_).length(0))]})
                r2 = ((self).f25) - (len(d_681_v37_))
            elif True:
                d_682___mcc_h21_ = source7_.cf15
                d_683_cf15_ = d_682___mcc_h21_
                d_684_v38_: _dafny.Array
                nw113_ = _dafny.Array(None, 7)
                nw113_[int(0)] = (self).f25
                nw113_[int(1)] = (self).f25
                nw113_[int(2)] = (self).f25
                nw113_[int(3)] = (d_598_v0_)[default__.safeIndex(268, (d_598_v0_).length(0))]
                nw113_[int(4)] = self.f26
                nw113_[int(5)] = self.f26
                nw113_[int(6)] = 244
                d_684_v38_ = nw113_
                d_685_v39_: D0
                d_685_v39_ = D0_DC2((self).f18, (d_598_v0_)[default__.safeIndex(268, (d_598_v0_).length(0))], d_684_v38_)
                (globalState).f14 = (d_685_v39_).cf6
                (globalState).f3 = d_604_cf0_
                d_686_v40_: C0
                nw114_ = C0()
                nw114_.ctor__()
                d_686_v40_ = nw114_
                d_687_v41_: _dafny.Array
                nw115_ = _dafny.Array(False, 13)
                d_687_v41_ = nw115_
                d_688_v42_: _dafny.Map
                d_688_v42_ = _dafny.Map({(self).f18: d_687_v41_})
                d_689_v43_: _dafny.Map
                d_689_v43_ = _dafny.Map({self.f26: d_688_v42_})
                d_690_v44_: D5
                d_690_v44_ = D5_DC12(d_688_v42_)
                d_689_v43_ = (d_689_v43_).set((self).f25, (d_690_v44_).cf22)
        elif source5_.is_DC1:
            d_691___mcc_h1_ = source5_.cf1
            d_692___mcc_h2_ = source5_.cf2
            d_693___mcc_h3_ = source5_.cf3
            d_694___mcc_h4_ = source5_.cf4
            d_695_cf4_ = d_694___mcc_h4_
            d_696_cf3_ = d_693___mcc_h3_
            d_697_cf2_ = d_692___mcc_h2_
            d_698_cf1_ = d_691___mcc_h1_
            d_698_cf1_ = d_698_cf1_
            d_699_v45_: _dafny.Array
            def lambda44_(d_700_i12_):
                return (self).f18

            init24_ = lambda44_
            nw116_ = _dafny.Array(None, 15)
            for i0_24_ in range(nw116_.length(0)):
                nw116_[i0_24_] = init24_(i0_24_)
            d_699_v45_ = nw116_
            index108_ = default__.safeIndex(923, (d_699_v45_).length(0))
            (d_699_v45_)[index108_] = (self).f18
            d_701_v46_: _dafny.MultiSet
            d_701_v46_ = _dafny.MultiSet([d_598_v0_, d_598_v0_])
            d_702_v47_: _dafny.Map
            d_702_v47_ = _dafny.Map({(self).f18: d_701_v46_})
            d_703_v48_: _dafny.Seq
            d_703_v48_ = _dafny.SeqWithoutIsStrInference([d_701_v46_, d_701_v46_])
            d_704_v49_: _dafny.Array
            nw117_ = _dafny.Array(None, 20)
            nw117_[int(0)] = d_701_v46_
            nw117_[int(1)] = d_701_v46_
            nw117_[int(2)] = (d_701_v46_).set(d_598_v0_, default__.abs(self.f26))
            nw117_[int(3)] = d_701_v46_
            nw117_[int(4)] = (d_701_v46_) | (_dafny.MultiSet([d_598_v0_, d_598_v0_, d_598_v0_]))
            nw117_[int(5)] = d_701_v46_
            nw117_[int(6)] = ((d_702_v47_)[(self).f18] if ((self).f18) in (d_702_v47_) else d_701_v46_)
            nw117_[int(7)] = (d_701_v46_) - (d_701_v46_)
            nw117_[int(8)] = d_701_v46_
            nw117_[int(9)] = d_701_v46_
            nw117_[int(10)] = (d_701_v46_).intersection(d_701_v46_)
            nw117_[int(11)] = (d_703_v48_)[default__.safeIndex(d_696_cf3_, len(d_703_v48_))]
            nw117_[int(12)] = (d_701_v46_).set(d_598_v0_, default__.abs(default__.fm4(self.f26, globalState)))
            nw117_[int(13)] = d_701_v46_
            nw117_[int(14)] = _dafny.MultiSet([d_598_v0_, d_598_v0_])
            nw117_[int(15)] = (_dafny.MultiSet([d_598_v0_])) | (d_701_v46_)
            nw117_[int(16)] = d_701_v46_
            nw117_[int(17)] = d_701_v46_
            nw117_[int(18)] = d_701_v46_
            nw117_[int(19)] = d_701_v46_
            d_704_v49_ = nw117_
            index109_ = default__.safeIndex(408, (d_704_v49_).length(0))
            (d_704_v49_)[index109_] = d_701_v46_
            d_705_v50_: D0
            d_705_v50_ = D0_DC0((self).f18)
            d_706_v51_: _dafny.Seq
            d_706_v51_ = _dafny.SeqWithoutIsStrInference([(self).f18])
            index110_ = default__.safeIndex(923, (d_699_v45_).length(0))
            index111_ = default__.safeIndex(408, (d_704_v49_).length(0))
            rhs95_ = d_695_cf4_
            rhs96_ = (d_697_cf2_) < (len((_dafny.SeqWithoutIsStrInference([(self).f18])).set(default__.safeIndex(d_696_cf3_, len(_dafny.SeqWithoutIsStrInference([(self).f18]))), (d_705_v50_).cf0)))
            rhs97_ = d_695_cf4_
            rhs98_ = (d_706_v51_)[default__.safeIndex(((d_701_v46_)[d_598_v0_] if (d_598_v0_) in (d_701_v46_) else d_696_cf3_), len(d_706_v51_))]
            rhs99_ = (d_701_v46_) | (_dafny.MultiSet([d_598_v0_]))
            lhs83_ = d_699_v45_
            lhs84_ = default__.safeIndex(923, (d_699_v45_).length(0))
            lhs85_ = globalState
            lhs86_ = d_704_v49_
            lhs87_ = default__.safeIndex(408, (d_704_v49_).length(0))
            d_695_cf4_ = rhs95_
            lhs83_[lhs84_] = rhs96_
            d_695_cf4_ = rhs97_
            lhs85_.f3 = rhs98_
            lhs86_[lhs87_] = rhs99_
            if (d_699_v45_)[default__.safeIndex(923, (d_699_v45_).length(0))]:
                index112_ = default__.safeIndex(930, (d_598_v0_).length(0))
                (d_598_v0_)[index112_] = len((_dafny.SeqWithoutIsStrInference([d_695_cf4_ for d_707_i13_ in range(default__.abs(325))])) + (d_698_cf1_))
                d_708_v52_: _dafny.Set
                d_708_v52_ = _dafny.Set({(d_699_v45_)[default__.safeIndex(923, (d_699_v45_).length(0))]})
                d_709_v53_: _dafny.Map
                d_709_v53_ = _dafny.Map({d_698_cf1_: d_696_cf3_})
                d_710_v54_: _dafny.Seq
                d_710_v54_ = _dafny.SeqWithoutIsStrInference([d_697_cf2_, d_697_cf2_])
                index113_ = default__.safeIndex(923, (d_699_v45_).length(0))
                index114_ = default__.safeIndex(923, (d_699_v45_).length(0))
                index115_ = default__.safeIndex(930, (d_598_v0_).length(0))
                rhs100_ = not(((self).f25) != ((len(d_708_v52_)) * ((self).f25)))
                rhs101_ = False
                rhs102_ = (len(d_709_v53_)) != ((len(d_710_v54_)) - (536))
                rhs103_ = ((d_699_v45_)[default__.safeIndex(923, (d_699_v45_).length(0))]) and (default__.fm13(globalState))
                rhs104_ = 647
                lhs88_ = globalState
                lhs89_ = d_699_v45_
                lhs90_ = default__.safeIndex(923, (d_699_v45_).length(0))
                lhs91_ = globalState
                lhs92_ = d_699_v45_
                lhs93_ = default__.safeIndex(923, (d_699_v45_).length(0))
                lhs94_ = d_598_v0_
                lhs95_ = default__.safeIndex(930, (d_598_v0_).length(0))
                lhs88_.f3 = rhs100_
                lhs89_[lhs90_] = rhs101_
                lhs91_.f9 = rhs102_
                lhs92_[lhs93_] = rhs103_
                lhs94_[lhs95_] = rhs104_
                d_711_v55_: _dafny.MultiSet
                d_711_v55_ = _dafny.MultiSet([False, (d_699_v45_)[default__.safeIndex(923, (d_699_v45_).length(0))]])
                d_712_v56_: D2
                d_712_v56_ = D2_DC7(D2_DC5(d_711_v55_))
                d_713_v57_: D2
                d_713_v57_ = D2_DC7(d_712_v56_)
                d_713_v57_ = d_713_v57_
                d_714_v58_: _dafny.Array
                def lambda45_(d_715_i14_):
                    return (d_715_i14_) - ((self).f25)

                init25_ = lambda45_
                nw118_ = _dafny.Array(None, 14)
                for i0_25_ in range(nw118_.length(0)):
                    nw118_[i0_25_] = init25_(i0_25_)
                d_714_v58_ = nw118_
                d_716_v59_: D0
                d_716_v59_ = D0_DC2((d_699_v45_)[default__.safeIndex(923, (d_699_v45_).length(0))], d_697_cf2_, d_714_v58_)
                rhs105_ = ((d_716_v59_).cf6) >= (615)
                rhs106_ = (self).f18
                lhs96_ = globalState
                lhs97_ = globalState
                lhs96_.f9 = rhs105_
                lhs97_.f3 = rhs106_
                (globalState).f3 = (d_699_v45_)[default__.safeIndex(923, (d_699_v45_).length(0))]
                d_717_v60_: D2
                d_717_v60_ = D2_DC5(d_711_v55_)
                d_718_v61_: _dafny.Set
                d_718_v61_ = _dafny.Set({d_717_v60_, default__.fm15((d_699_v45_)[default__.safeIndex(923, (d_699_v45_).length(0))], d_698_cf1_, d_698_cf1_, globalState)})
                (globalState).f3 = not ((d_699_v45_)[default__.safeIndex(923, (d_699_v45_).length(0))]) or ((d_718_v61_).issubset(d_718_v61_))
            elif True:
                (globalState).f9 = (d_697_cf2_) > (len(_dafny.SeqWithoutIsStrInference([d_695_cf4_ for d_719_i15_ in range(default__.abs(824))])))
                d_720_v62_: _dafny.Array
                def lambda46_(d_721_v51_):
                    def lambda47_(d_722_i16_):
                        return (d_721_v51_)

                    return lambda47_

                init26_ = lambda46_(d_706_v51_)
                nw119_ = _dafny.Array(None, 6)
                for i0_26_ in range(nw119_.length(0)):
                    nw119_[i0_26_] = init26_(i0_26_)
                d_720_v62_ = nw119_
                index116_ = default__.safeIndex(630, (d_720_v62_).length(0))
                (d_720_v62_)[index116_] = ((d_706_v51_).set(default__.safeIndex(self.f26, len(d_706_v51_)), (self).f18)) + (default__.fm16(globalState))
                d_723_v63_: _dafny.Seq
                d_723_v63_ = _dafny.SeqWithoutIsStrInference([d_697_cf2_])
                index117_ = default__.safeIndex(630, (d_720_v62_).length(0))
                (d_720_v62_)[index117_] = (((d_706_v51_) + ((d_706_v51_) + ((_dafny.SeqWithoutIsStrInference([False])).set(default__.safeIndex(self.f26, len(_dafny.SeqWithoutIsStrInference([False]))), (self).f18)))).set(default__.safeIndex(len((d_723_v63_) + (d_723_v63_)), len((d_706_v51_) + ((d_706_v51_) + ((_dafny.SeqWithoutIsStrInference([False])).set(default__.safeIndex(self.f26, len(_dafny.SeqWithoutIsStrInference([False]))), (self).f18))))), False)).set(default__.safeIndex(self.f26, len(((d_706_v51_) + ((d_706_v51_) + ((_dafny.SeqWithoutIsStrInference([False])).set(default__.safeIndex(self.f26, len(_dafny.SeqWithoutIsStrInference([False]))), (self).f18)))).set(default__.safeIndex(len((d_723_v63_) + (d_723_v63_)), len((d_706_v51_) + ((d_706_v51_) + ((_dafny.SeqWithoutIsStrInference([False])).set(default__.safeIndex(self.f26, len(_dafny.SeqWithoutIsStrInference([False]))), (self).f18))))), False))), (d_699_v45_)[default__.safeIndex(923, (d_699_v45_).length(0))])
                d_724_v64_: _dafny.Set
                d_724_v64_ = _dafny.Set({default__.fm17(globalState)})
                d_725_v65_: _dafny.Seq
                d_725_v65_ = _dafny.SeqWithoutIsStrInference([not((self).f18), (d_699_v45_)[default__.safeIndex(923, (d_699_v45_).length(0))]])
                d_724_v64_ = (d_724_v64_ if (self).f18 else _dafny.Set({d_725_v65_, d_725_v65_}))
                rhs107_ = not (default__.fm13(globalState)) or ((d_699_v45_)[default__.safeIndex(923, (d_699_v45_).length(0))])
                rhs108_ = self.f26
                rhs109_ = (d_697_cf2_) - (d_696_cf3_)
                lhs98_ = globalState
                lhs99_ = globalState
                lhs100_ = globalState
                lhs98_.f9 = rhs107_
                lhs99_.f14 = rhs108_
                lhs100_.f14 = rhs109_
                (globalState).f14 = d_696_cf3_
            (globalState).f3 = (d_699_v45_)[default__.safeIndex(923, (d_699_v45_).length(0))]
        elif True:
            d_726___mcc_h5_ = source5_.cf5
            d_727___mcc_h6_ = source5_.cf6
            d_728___mcc_h7_ = source5_.cf7
            d_729_cf7_ = d_728___mcc_h7_
            d_730_cf6_ = d_727___mcc_h6_
            d_731_cf5_ = d_726___mcc_h5_
            d_732_v66_: _dafny.Set
            d_732_v66_ = _dafny.Set({(self).f18})
            d_733_v67_: _dafny.Seq
            d_733_v67_ = _dafny.SeqWithoutIsStrInference([len(d_732_v66_), (self).f25, d_730_cf6_, d_730_cf6_, d_730_cf6_])
            d_734_v68_: _dafny.Set
            d_734_v68_ = _dafny.Set({(d_733_v67_)[default__.safeIndex(d_730_cf6_, len(d_733_v67_))]})
            (globalState).f9 = (d_734_v68_).ispropersubset(d_734_v68_)
            d_735_v69_: str
            d_735_v69_ = _dafny.CodePoint('q')
            index118_ = default__.safeIndex(337, (d_729_cf7_).length(0))
            (d_729_cf7_)[index118_] = ((self).f25) * ((0) - (d_730_cf6_))
            d_736_v70_: _dafny.MultiSet
            d_736_v70_ = _dafny.MultiSet([(self).f18])
            d_737_v71_: D2
            d_737_v71_ = D2_DC5(d_736_v70_)
            d_738_v72_: D2
            d_738_v72_ = D2_DC5((d_737_v71_).cf12)
            d_739_v73_: _dafny.Seq
            d_739_v73_ = _dafny.SeqWithoutIsStrInference([(self).f18, d_731_cf5_])
            d_740_v74_: _dafny.Map
            d_740_v74_ = _dafny.Map({len(d_739_v73_): (self).f25})
            d_741_v75_: _dafny.Seq
            d_741_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ayplhylwi"))
            index119_ = default__.safeIndex(337, (d_729_cf7_).length(0))
            rhs110_ = _dafny.CodePoint('c')
            rhs111_ = -472
            rhs112_ = d_738_v72_
            rhs113_ = default__.safeModuloInt(self.f26, d_730_cf6_)
            rhs114_ = (len(d_740_v74_)) - (len((d_741_v75_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_742_i17_ in range(default__.abs(412))]))))
            lhs101_ = d_729_cf7_
            lhs102_ = default__.safeIndex(337, (d_729_cf7_).length(0))
            lhs103_ = self
            d_735_v69_ = rhs110_
            lhs101_[lhs102_] = rhs111_
            d_737_v71_ = rhs112_
            lhs103_.f26 = rhs113_
            r2 = rhs114_
            index120_ = default__.safeIndex(337, (d_729_cf7_).length(0))
            (d_729_cf7_)[index120_] = self.f26
            d_743_v76_: _dafny.Array
            nw120_ = _dafny.Array(_dafny.Seq({}), 13)
            d_743_v76_ = nw120_
            d_744_v77_: _dafny.Seq
            d_744_v77_ = _dafny.SeqWithoutIsStrInference([d_729_cf7_])
            index121_ = default__.safeIndex(526, (d_743_v76_).length(0))
            (d_743_v76_)[index121_] = (d_744_v77_ if d_731_cf5_ else d_744_v77_)
            index122_ = default__.safeIndex(526, (d_743_v76_).length(0))
            (d_743_v76_)[index122_] = (d_744_v77_) + ((d_744_v77_).set(default__.safeIndex((d_729_cf7_)[default__.safeIndex(337, (d_729_cf7_).length(0))], len(d_744_v77_)), d_598_v0_))
        d_745_v78_: _dafny.Set
        d_745_v78_ = _dafny.Set({(self).f18})
        d_746_v79_: _dafny.Seq
        d_746_v79_ = _dafny.SeqWithoutIsStrInference([self.f26, self.f26, 224])
        if ((len(d_745_v78_)) * (self.f26)) in ((d_746_v79_) + (d_746_v79_)):
            (globalState).f9 = (self).f18
            d_747_v80_: _dafny.Seq
            d_747_v80_ = _dafny.SeqWithoutIsStrInference([(self).f18])
            if ((d_747_v80_) + (d_747_v80_)) < (d_747_v80_):
                d_748_v81_: _dafny.Array
                def lambda48_(d_749_i18_):
                    return (self).f18

                init27_ = lambda48_
                nw121_ = _dafny.Array(None, 27)
                for i0_27_ in range(nw121_.length(0)):
                    nw121_[i0_27_] = init27_(i0_27_)
                d_748_v81_ = nw121_
                d_750_v82_: T1
                nw122_ = C3()
                nw122_.ctor__(d_748_v81_, _dafny.Set({d_748_v81_, d_748_v81_}), (self).f18)
                d_750_v82_ = nw122_
                (self).f26 = default__.safeModuloInt((0) - ((self).f25), (0) - (default__.safeDivisionInt((self).f25, len(_dafny.Map({(self).f18: d_750_v82_})))))
                d_751_v83_: str
                d_751_v83_ = _dafny.CodePoint('y')
                d_752_v84_: _dafny.Seq
                d_752_v84_ = _dafny.SeqWithoutIsStrInference([d_751_v83_, d_751_v83_])
                (globalState).f14 = (((_dafny.MultiSet([d_751_v83_])) | (_dafny.MultiSet(d_752_v84_))).cardinality) + ((self).f25)
                d_753_v85_: _dafny.Set
                d_753_v85_ = _dafny.Set({854, self.f26})
                d_753_v85_ = d_753_v85_
                d_754_v86_: _dafny.Map
                d_754_v86_ = _dafny.Map({self.f26: -870})
                index123_ = default__.safeIndex(735, (d_748_v81_).length(0))
                (d_748_v81_)[index123_] = (((d_754_v86_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcaginhv")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcaginhv")))) in (d_754_v86_) else -392)) != ((0) - (len(d_746_v79_)))
                index124_ = default__.safeIndex(735, (d_748_v81_).length(0))
                (d_748_v81_)[index124_] = ((self).f18) == ((self).f18)
                d_755_v88_: T0
                nw123_ = C2()
                nw123_.ctor__(d_750_v82_.f17, not((self).f18))
                d_755_v88_ = nw123_
                d_756_v89_: _dafny.Map
                d_756_v89_ = _dafny.Map({(d_748_v81_)[default__.safeIndex(735, (d_748_v81_).length(0))]: d_755_v88_})
                d_757_v90_: bool
                d_758_v91_: bool
                d_759_v92_: int
                out44_: bool
                out45_: bool
                out46_: int
                def iife77_():
                    coll41_ = _dafny.Set()
                    compr_41_: int
                    for compr_41_ in (_dafny.SeqWithoutIsStrInference([self.f26 for d_760_i19_ in range(default__.abs(-119))])).Elements:
                        d_761_v87_: int = compr_41_
                        if (d_761_v87_) in (_dafny.SeqWithoutIsStrInference([self.f26 for d_760_i19_ in range(default__.abs(-119))])):
                            coll41_ = coll41_.union(_dafny.Set([(d_761_v87_) + ((0) - (len(_dafny.Map({len(_dafny.Set({False})): _dafny.SeqWithoutIsStrInference([D3_DC9() for d_762_i20_ in range(default__.abs(33))])}))))]))
                    return _dafny.Set(coll41_)
                out44_, out45_, out46_ = (d_750_v82_).m0((self).f25, (iife77_()
                ).issubset(d_753_v85_), len((d_756_v89_ if (d_748_v81_)[default__.safeIndex(735, (d_748_v81_).length(0))] else d_756_v89_)), globalState)
                d_757_v90_ = out44_
                d_758_v91_ = out45_
                d_759_v92_ = out46_
            elif True:
                d_763_v93_: C2
                nw124_ = C2()
                nw124_.ctor__(self.f17, (self).f18)
                d_763_v93_ = nw124_
                d_764_v94_: D3
                d_764_v94_ = D3_DC9()
                d_765_v95_: _dafny.Map
                d_765_v95_ = _dafny.Map({(self).f25: d_764_v94_})
                (globalState).f3 = (_dafny.Map({len(d_745_v78_): d_764_v94_})) != (d_765_v95_)
                d_766_v96_: _dafny.Seq
                d_766_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntmfjd"))
                d_767_v97_: _dafny.Seq
                d_767_v97_ = _dafny.SeqWithoutIsStrInference([(d_766_v96_) + (d_766_v96_)])
                r2 = len((d_767_v97_)[default__.safeIndex((self).f25, len(d_767_v97_))])
                d_768_v98_: _dafny.Array
                def lambda49_(d_769_i21_):
                    return (self).f18

                init28_ = lambda49_
                nw125_ = _dafny.Array(None, 27)
                for i0_28_ in range(nw125_.length(0)):
                    nw125_[i0_28_] = init28_(i0_28_)
                d_768_v98_ = nw125_
                d_770_v99_: C3
                nw126_ = C3()
                nw126_.ctor__(d_768_v98_, self.f17, (self).f18)
                d_770_v99_ = nw126_
                d_771_v100_: _dafny.Map
                d_771_v100_ = _dafny.Map({(0) - ((self).f25): self.f26})
                d_771_v100_ = (d_771_v100_).set(self.f26, (self).f25)
            d_772_v101_: _dafny.Array
            def lambda50_(d_773_i22_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wllppav"))

            init29_ = lambda50_
            nw127_ = _dafny.Array(None, 2)
            for i0_29_ in range(nw127_.length(0)):
                nw127_[i0_29_] = init29_(i0_29_)
            d_772_v101_ = nw127_
            d_774_v102_: _dafny.Seq
            d_774_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwvqrarpx"))
            index125_ = default__.safeIndex(808, (d_772_v101_).length(0))
            (d_772_v101_)[index125_] = d_774_v102_
            index126_ = default__.safeIndex(808, (d_772_v101_).length(0))
            (d_772_v101_)[index126_] = d_774_v102_
            d_746_v79_ = (d_746_v79_ if (self).f18 else d_746_v79_)
            d_775_v103_: str
            d_775_v103_ = _dafny.CodePoint('b')
            d_775_v103_ = d_775_v103_
        elif True:
            (globalState).f3 = (self).f18
            d_776_v104_: T0
            nw128_ = C2()
            nw128_.ctor__(self.f17, not((self).f18))
            d_776_v104_ = nw128_
            d_776_v104_ = d_776_v104_
            d_777_v105_: _dafny.Array
            nw129_ = _dafny.Array(False, 17)
            d_777_v105_ = nw129_
            d_778_v106_: C3
            nw130_ = C3()
            nw130_.ctor__(d_777_v105_, self.f17, (self).f18)
            d_778_v106_ = nw130_
            (globalState).f3 = (self).f18
            (globalState).f9 = default__.fm13(globalState)
        d_779_v107_: _dafny.Set
        d_779_v107_ = _dafny.Set({(self).f25})
        index127_ = default__.safeIndex(718, (d_598_v0_).length(0))
        (d_598_v0_)[index127_] = (d_746_v79_)[default__.safeIndex(len(d_779_v107_), len(d_746_v79_))]
        index128_ = default__.safeIndex(718, (d_598_v0_).length(0))
        (d_598_v0_)[index128_] = self.f26
        d_780_i23_: int
        d_780_i23_ = 0
        with _dafny.label("4"):
            while (self).f18:
                with _dafny.c_label("4"):
                    if (d_780_i23_) >= (100):
                        raise _dafny.Break("4")
                    d_780_i23_ = (d_780_i23_) + (1)
                    d_781_v108_: _dafny.Map
                    d_781_v108_ = _dafny.Map({d_746_v79_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsdlppl"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghnl")))})
                    d_782_v109_: _dafny.Seq
                    d_782_v109_ = _dafny.SeqWithoutIsStrInference([not((self).f18), (self).f18])
                    d_781_v108_ = (d_781_v108_).set((_dafny.SeqWithoutIsStrInference([-581])) + ((d_746_v79_).set(default__.safeIndex((d_598_v0_)[default__.safeIndex(718, (d_598_v0_).length(0))], len(d_746_v79_)), len(d_782_v109_))), (self).f18)
                    d_783_v110_: _dafny.Map
                    d_783_v110_ = _dafny.Map({_dafny.Map({(self).f18: (self).f25}): (self).f18})
                    d_784_v111_: _dafny.Map
                    d_784_v111_ = _dafny.Map({(self).f18: (d_598_v0_)[default__.safeIndex(718, (d_598_v0_).length(0))]})
                    d_783_v110_ = (d_783_v110_).set((d_784_v111_) | (d_784_v111_), (self).f18)
                    d_785_v112_: _dafny.Seq
                    d_785_v112_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hn"))
                    (self).f26 = default__.safeModuloInt(default__.safeModuloInt((self).f25, self.f26), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_786_i24_ in range(default__.abs(972))])) + (d_785_v112_)))
                    d_787_v113_: _dafny.Map
                    d_787_v113_ = _dafny.Map({10: d_785_v112_})
                    d_788_v114_: _dafny.Seq
                    d_788_v114_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_785_v112_), (self).f25])])
                    d_789_v115_: str
                    d_789_v115_ = _dafny.CodePoint('c')
                    d_785_v112_ = ((d_785_v112_) + ((d_785_v112_) + (((d_787_v113_)[(d_746_v79_)[default__.safeIndex(len(d_788_v114_), len(d_746_v79_))]] if ((d_746_v79_)[default__.safeIndex(len(d_788_v114_), len(d_746_v79_))]) in (d_787_v113_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptbvpe")))))).set(default__.safeIndex((d_598_v0_)[default__.safeIndex(718, (d_598_v0_).length(0))], len((d_785_v112_) + ((d_785_v112_) + (((d_787_v113_)[(d_746_v79_)[default__.safeIndex(len(d_788_v114_), len(d_746_v79_))]] if ((d_746_v79_)[default__.safeIndex(len(d_788_v114_), len(d_746_v79_))]) in (d_787_v113_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptbvpe"))))))), (_dafny.CodePoint('p') if (self).f18 else d_789_v115_))
                    pass
            pass
        d_790_i25_: int
        d_790_i25_ = 0
        with _dafny.label("5"):
            while (self.f26) > ((self).f25):
                with _dafny.c_label("5"):
                    if (d_790_i25_) >= (100):
                        raise _dafny.Break("5")
                    d_790_i25_ = (d_790_i25_) + (1)
                    r2 = (0) - (default__.fm4(self.f26, globalState))
                    d_791_v116_: _dafny.Array
                    nw131_ = _dafny.Array(_dafny.Array(None, 0), 28)
                    d_791_v116_ = nw131_
                    d_792_v117_: _dafny.Array
                    def lambda51_(d_793_i26_):
                        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekeegfkgk"))

                    init30_ = lambda51_
                    nw132_ = _dafny.Array(None, 21)
                    for i0_30_ in range(nw132_.length(0)):
                        nw132_[i0_30_] = init30_(i0_30_)
                    d_792_v117_ = nw132_
                    index129_ = default__.safeIndex(385, (d_791_v116_).length(0))
                    (d_791_v116_)[index129_] = d_792_v117_
                    index130_ = default__.safeIndex(385, (d_791_v116_).length(0))
                    (d_791_v116_)[index130_] = d_792_v117_
                    d_794_v118_: C0
                    nw133_ = C0()
                    nw133_.ctor__()
                    d_794_v118_ = nw133_
                    index131_ = default__.safeIndex(718, (d_598_v0_).length(0))
                    (d_598_v0_)[index131_] = self.f26
                    pass
            pass
        hi2_ = (self).f25
        for d_795_i27_ in range(self.f26, hi2_):
            d_796_v119_: str
            d_796_v119_ = _dafny.CodePoint('d')
            d_796_v119_ = d_796_v119_
            d_797_v120_: _dafny.Seq
            d_797_v120_ = _dafny.SeqWithoutIsStrInference([d_796_v119_])
            d_798_v121_: _dafny.Map
            d_798_v121_ = _dafny.Map({(d_598_v0_)[default__.safeIndex(718, (d_598_v0_).length(0))]: self.f26})
            d_799_v122_: D3
            d_799_v122_ = D3_DC10(d_797_v120_, (self).f18, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yivqbw")), d_797_v120_, len(d_798_v121_))
            (globalState).f9 = (len(_dafny.Map({d_799_v122_: (self).f18}))) != (d_795_i27_)
            d_800_v123_: _dafny.Map
            d_800_v123_ = _dafny.Map({(self).f18: _dafny.SeqWithoutIsStrInference([d_796_v119_ for d_801_i28_ in range(default__.abs(-912))])})
            d_800_v123_ = _dafny.Map({(self).f18: _dafny.SeqWithoutIsStrInference([d_796_v119_ for d_802_i29_ in range(default__.abs(103))])})
            if (self).f18:
                d_803_v124_: _dafny.Map
                d_803_v124_ = _dafny.Map({(self.f26) > (len((_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])).set(default__.safeIndex(self.f26, len(_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18]))), False))): (self).f25})
                d_803_v124_ = (d_803_v124_).set(not((self).f18), (d_598_v0_)[default__.safeIndex(718, (d_598_v0_).length(0))])
                d_804_v125_: C0
                nw134_ = C0()
                nw134_.ctor__()
                d_804_v125_ = nw134_
                d_805_v126_: _dafny.Array
                nw135_ = _dafny.Array(_dafny.MultiSet({}), 27)
                d_805_v126_ = nw135_
                index132_ = default__.safeIndex(168, (d_805_v126_).length(0))
                (d_805_v126_)[index132_] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([262, len(_dafny.SeqWithoutIsStrInference([True])), self.f26, (_dafny.MultiSet([(self).f18])).cardinality])).cardinality, 715]) for d_806_i30_ in range(default__.abs(528))]))
                index133_ = default__.safeIndex(168, (d_805_v126_).length(0))
                (d_805_v126_)[index133_] = (_dafny.MultiSet([d_746_v79_])) - (_dafny.MultiSet([d_746_v79_, d_746_v79_, d_746_v79_]))
                d_807_v127_: C0
                nw136_ = C0()
                nw136_.ctor__()
                d_807_v127_ = nw136_
                (globalState).f9 = (default__.safeDivisionInt(self.f26, (d_598_v0_)[default__.safeIndex(718, (d_598_v0_).length(0))])) <= ((-298) - (len(d_745_v78_)))
            elif True:
                (globalState).f9 = not((self).f18)
                d_808_v128_: C1
                nw137_ = C1()
                nw137_.ctor__()
                d_808_v128_ = nw137_
                d_809_v129_: _dafny.Map
                d_809_v129_ = _dafny.Map({len(d_797_v120_): (self).f18})
                d_810_v130_: _dafny.Map
                d_810_v130_ = _dafny.Map({d_795_i27_: d_809_v129_})
                d_811_v131_: _dafny.Map
                d_811_v131_ = _dafny.Map({d_810_v130_: (self).f25})
                (globalState).f14 = ((d_811_v131_)[d_810_v130_] if (d_810_v130_) in (d_811_v131_) else -449)
                d_812_v132_: _dafny.Array
                nw138_ = _dafny.Array(None, 6)
                nw138_[int(0)] = d_598_v0_
                nw138_[int(1)] = d_598_v0_
                nw138_[int(2)] = d_598_v0_
                nw138_[int(3)] = d_598_v0_
                nw138_[int(4)] = d_598_v0_
                nw138_[int(5)] = d_598_v0_
                d_812_v132_ = nw138_
                d_813_v133_: _dafny.Map
                d_813_v133_ = _dafny.Map({d_812_v132_: self.f26})
                d_814_v134_: _dafny.Map
                d_814_v134_ = _dafny.Map({(self).f18: d_812_v132_})
                d_813_v133_ = (d_813_v133_).set(((d_814_v134_)[(d_808_v128_).fm5(self.f26, d_746_v79_, globalState)] if ((d_808_v128_).fm5(self.f26, d_746_v79_, globalState)) in (d_814_v134_) else d_812_v132_), (0) - ((d_795_i27_ if (self).f18 else self.f26)))
                d_815_v135_: D0
                d_815_v135_ = D0_DC2((self).f18, (self).f25, d_598_v0_)
                d_816_v136_: _dafny.MultiSet
                d_816_v136_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_796_v119_ for d_817_i31_ in range(default__.abs(81))])])
                d_818_v137_: _dafny.MultiSet
                d_818_v137_ = _dafny.MultiSet([d_795_i27_])
                d_819_v138_: _dafny.Array
                nw139_ = _dafny.Array(None, 14)
                nw139_[int(0)] = ((self).f18 if (self).f18 else (self).f18)
                nw139_[int(1)] = not ((self).f18) or ((self).f18)
                nw139_[int(2)] = (self).f18
                nw139_[int(3)] = ((_dafny.MultiSet([d_795_i27_])).set((self).f25, default__.abs(self.f26))).issubset(default__.fm24((d_815_v135_).cf5, (d_598_v0_)[default__.safeIndex(718, (d_598_v0_).length(0))], globalState))
                nw139_[int(4)] = (self).f18
                nw139_[int(5)] = (self).f18
                nw139_[int(6)] = (((d_816_v136_)[d_797_v120_] if (d_797_v120_) in (d_816_v136_) else (d_598_v0_)[default__.safeIndex(718, (d_598_v0_).length(0))])) > (len(d_797_v120_))
                nw139_[int(7)] = (self).f18
                nw139_[int(8)] = ((self).f25) > ((d_598_v0_)[default__.safeIndex(718, (d_598_v0_).length(0))])
                nw139_[int(9)] = (d_808_v128_).fm5(self.f26, d_746_v79_, globalState)
                nw139_[int(10)] = not((self).f18)
                nw139_[int(11)] = (self).f18
                nw139_[int(12)] = (self).f18
                nw139_[int(13)] = (len(d_797_v120_)) >= (((d_818_v137_)[self.f26] if (self.f26) in (d_818_v137_) else d_795_i27_))
                d_819_v138_ = nw139_
                index134_ = default__.safeIndex(295, (d_819_v138_).length(0))
                (d_819_v138_)[index134_] = (self).f18
                index135_ = default__.safeIndex(295, (d_819_v138_).length(0))
                rhs115_ = (self).f18
                rhs116_ = d_796_v119_
                lhs104_ = d_819_v138_
                lhs105_ = default__.safeIndex(295, (d_819_v138_).length(0))
                lhs104_[lhs105_] = rhs115_
                d_796_v119_ = rhs116_
        r0 = ((self).f25) == (730)
        d_820_v139_: _dafny.Array
        nw140_ = _dafny.Array(False, 11)
        d_820_v139_ = nw140_
        d_821_v140_: _dafny.Map
        d_821_v140_ = _dafny.Map({d_820_v139_: self.f26})
        r1 = d_821_v140_
        d_822_v141_: _dafny.Map
        d_822_v141_ = _dafny.Map({(self).f18: d_779_v107_})
        r2 = default__.safeModuloInt(default__.safeModuloInt((d_598_v0_)[default__.safeIndex(718, (d_598_v0_).length(0))], (d_598_v0_)[default__.safeIndex(718, (d_598_v0_).length(0))]), len(((d_822_v141_)[not((self).f18)] if (not((self).f18)) in (d_822_v141_) else d_779_v107_)))
        return r0, r1, r2

    @property
    def f25(self):
        return self._f25

class C5(T2, T1, T0):
    def  __init__(self):
        self._f17: _dafny.Set = _dafny.Set({})
        self._f19: _dafny.Array = _dafny.Array(None, 0)
        self._f21: _dafny.Map = _dafny.Map({})
        self._f20: int = int(0)
        self._f18: bool = False
        self._f29: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f30: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f17(self):
        return self._f17
    @f17.setter
    def f17(self, value):
        self._f17 = value
    @property
    def f19(self):
        return self._f19
    @f19.setter
    def f19(self, value):
        self._f19 = value
    @property
    def f21(self):
        return self._f21
    @f21.setter
    def f21(self, value):
        self._f21 = value
    @property
    def f20(self):
        return self._f20
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f29, f30, f20, f21, f19, f17, f18):
        (self)._f29 = f29
        (self)._f30 = f30
        (self)._f20 = f20
        (self).f21 = f21
        (self).f19 = f19
        (self).f17 = f17
        (self)._f18 = f18

    def fm1(self, p0, globalState):
        return ((_dafny.MultiSet([len(_dafny.Set({not((self).f30), (self).f30})), len(_dafny.SeqWithoutIsStrInference([(self).f20, (self).f20])), (self).f20, len(_dafny.Map({(self).f30: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbbbe"))})), (self).f20])).cardinality) == (((self).f20 if (self).f18 else (self).f20))

    def fm0(self, p0, p1, globalState):
        source8_ = D1_DC3(_dafny.MultiSet([(self).f20]))
        if source8_.is_DC4:
            d_823___mcc_h0_ = source8_.cf9
            d_824___mcc_h1_ = source8_.cf10
            d_825___mcc_h2_ = source8_.cf11
            d_826_cf11_ = d_825___mcc_h2_
            d_827_cf10_ = d_824___mcc_h1_
            d_828_cf9_ = d_823___mcc_h0_
            return 902
        elif True:
            d_829___mcc_h3_ = source8_.cf8
            d_830_cf8_ = d_829___mcc_h3_
            return (0) - ((self).f20)

    def m0(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r1 = (self).f18
        d_831_v0_: _dafny.Array
        nw141_ = _dafny.Array(int(0), 3)
        d_831_v0_ = nw141_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_831_v0_).length(0)):
            d_832_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_832_i0_)) and ((d_832_i0_) < ((d_831_v0_).length(0)))):
                (d_831_v0_)[(d_832_i0_)] = (d_832_i0_) * (((self).f20) - (p0))
        d_833_v1_: _dafny.Set
        d_833_v1_ = _dafny.Set({True, p1})
        d_834_i1_: int
        d_834_i1_ = 0
        with _dafny.label("6"):
            while (len(d_833_v1_)) < ((self).f20):
                with _dafny.c_label("6"):
                    if (d_834_i1_) >= (100):
                        raise _dafny.Break("6")
                    d_834_i1_ = (d_834_i1_) + (1)
                    arr6_ = self.f19
                    index136_ = default__.safeIndex(319, (self.f19).length(0))
                    arr6_[index136_] = (self).f18
                    d_835_v2_: str
                    d_835_v2_ = _dafny.CodePoint('v')
                    arr7_ = self.f19
                    index137_ = default__.safeIndex(774, (self.f19).length(0))
                    arr7_[index137_] = (False) not in (_dafny.SeqWithoutIsStrInference([(self).f18, p1, not(p1)]))
                    d_836_v3_: _dafny.Seq
                    d_836_v3_ = _dafny.SeqWithoutIsStrInference([(self).f18])
                    arr8_ = self.f19
                    index138_ = default__.safeIndex(319, (self.f19).length(0))
                    arr9_ = self.f19
                    index139_ = default__.safeIndex(774, (self.f19).length(0))
                    rhs117_ = len((self).f29)
                    rhs118_ = ((self).f30 if False else True)
                    rhs119_ = d_835_v2_
                    rhs120_ = not((d_836_v3_) < (d_836_v3_))
                    lhs106_ = globalState
                    lhs107_ = self.f19
                    lhs108_ = default__.safeIndex(319, (self.f19).length(0))
                    lhs109_ = self.f19
                    lhs110_ = default__.safeIndex(774, (self.f19).length(0))
                    lhs106_.f14 = rhs117_
                    lhs107_[lhs108_] = rhs118_
                    d_835_v2_ = rhs119_
                    lhs109_[lhs110_] = rhs120_
                    d_837_v4_: _dafny.Array
                    nw142_ = _dafny.Array(_dafny.Seq({}), 12)
                    d_837_v4_ = nw142_
                    d_838_v5_: _dafny.Map
                    d_838_v5_ = _dafny.Map({d_835_v2_: d_837_v4_})
                    d_839_v6_: _dafny.Seq
                    d_839_v6_ = _dafny.SeqWithoutIsStrInference([d_837_v4_, d_837_v4_])
                    d_840_v7_: _dafny.Map
                    d_840_v7_ = _dafny.Map({(self).f29: (d_839_v6_)[default__.safeIndex(118, len(d_839_v6_))]})
                    d_838_v5_ = (d_838_v5_).set(_dafny.CodePoint('g'), ((d_840_v7_)[(self).f29] if ((self).f29) in (d_840_v7_) else d_837_v4_))
                    d_841_v8_: _dafny.Seq
                    d_841_v8_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_842_v9_: D1
                    d_842_v9_ = D1_DC3(_dafny.MultiSet(d_841_v8_))
                    d_842_v9_ = D1_DC3(_dafny.MultiSet([(self).f20]))
                    d_843_v10_: _dafny.Array
                    def lambda52_(d_844_i2_):
                        return (self).f18

                    init31_ = lambda52_
                    nw143_ = _dafny.Array(None, 28)
                    for i0_31_ in range(nw143_.length(0)):
                        nw143_[i0_31_] = init31_(i0_31_)
                    d_843_v10_ = nw143_
                    d_843_v10_ = d_843_v10_
                    pass
            pass
        d_845_v11_: _dafny.Array
        nw144_ = _dafny.Array(_dafny.Set({}), 14)
        d_845_v11_ = nw144_
        d_846_v12_: _dafny.Map
        d_846_v12_ = _dafny.Map({(self).f20: (self).f18})
        d_847_v13_: _dafny.Set
        d_847_v13_ = _dafny.Set({d_846_v12_, d_846_v12_})
        index140_ = default__.safeIndex(511, (d_845_v11_).length(0))
        (d_845_v11_)[index140_] = d_847_v13_
        d_848_v14_: _dafny.Set
        d_848_v14_ = _dafny.Set({(self).f29})
        index141_ = default__.safeIndex(511, (d_845_v11_).length(0))
        (d_845_v11_)[index141_] = (d_847_v13_ if (d_848_v14_).ispropersubset(d_848_v14_) else d_847_v13_)
        arr10_ = self.f19
        index142_ = default__.safeIndex(575, (self.f19).length(0))
        arr10_[index142_] = (self).f30
        d_849_v15_: _dafny.Seq
        d_849_v15_ = _dafny.SeqWithoutIsStrInference([(self).f30])
        d_850_v16_: _dafny.Seq
        d_850_v16_ = d_849_v15_
        arr11_ = self.f19
        index143_ = default__.safeIndex(575, (self.f19).length(0))
        arr11_[index143_] = (((d_850_v16_)) + (d_849_v15_)) != ((d_849_v15_ if (self).f30 else d_849_v15_))
        d_851_v17_: _dafny.Map
        d_851_v17_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_852_i4_ in range(default__.abs(49))])): (0) - (len((self).f29))})
        d_853_v18_: _dafny.Map
        d_853_v18_ = _dafny.Map({(self).f29: (self).f18})
        hi3_ = (self).f20
        for d_854_i3_ in range(((d_851_v17_)[len(d_853_v18_)] if (len(d_853_v18_)) in (d_851_v17_) else len(_dafny.Map({(self).f29: p1}))), hi3_):
            d_855_v19_: _dafny.Array
            nw145_ = _dafny.Array(D5.default()(), 26)
            d_855_v19_ = nw145_
            d_856_v20_: D3
            d_856_v20_ = D3_DC10((self).f29, (self.f19)[default__.safeIndex(575, (self.f19).length(0))], (self).f29, (self).f29, len(d_851_v17_))
            d_855_v19_ = (d_855_v19_ if ((d_856_v20_).cf20) != (p0) else (d_855_v19_ if p1 else d_855_v19_))
            d_857_v22_: _dafny.Map
            def iife78_():
                coll42_ = _dafny.Map()
                compr_42_: int
                for compr_42_ in _dafny.IntegerRange(500, 409):
                    d_858_v21_: int = compr_42_
                    if ((500) <= (d_858_v21_)) and ((d_858_v21_) < (409)):
                        coll42_[(d_858_v21_) - (d_854_i3_)] = 147
                return _dafny.Map(coll42_)
            d_857_v22_ = _dafny.Map({iife78_()
            : p2})
            d_859_v23_: _dafny.MultiSet
            d_859_v23_ = _dafny.MultiSet([(self).fm1(d_857_v22_, globalState)])
            (globalState).f3 = not((((self).f30) not in (d_849_v15_)) in (d_859_v23_))
            arr12_ = self.f19
            index144_ = default__.safeIndex(575, (self.f19).length(0))
            rhs121_ = (self).f30
            rhs122_ = 652
            rhs123_ = p0
            lhs111_ = self.f19
            lhs112_ = default__.safeIndex(575, (self.f19).length(0))
            lhs113_ = globalState
            lhs114_ = globalState
            lhs111_[lhs112_] = rhs121_
            lhs113_.f14 = rhs122_
            lhs114_.f14 = rhs123_
            d_860_v24_: _dafny.Seq
            d_860_v24_ = _dafny.SeqWithoutIsStrInference([d_859_v23_, d_859_v23_, d_859_v23_])
            d_859_v23_ = (d_860_v24_)[default__.safeIndex((len(d_833_v1_)) + (len(d_833_v1_)), len(d_860_v24_))]
        r0 = (self).f30
        d_861_v25_: _dafny.Set
        d_861_v25_ = _dafny.Set({(self).f20})
        d_862_v26_: _dafny.Array
        nw146_ = _dafny.Array(D11.default()(), 29)
        d_862_v26_ = nw146_
        d_863_v27_: _dafny.Map
        d_863_v27_ = _dafny.Map({d_862_v26_: d_861_v25_})
        d_864_v28_: _dafny.MultiSet
        d_864_v28_ = _dafny.MultiSet([(self).f30, (self).f30])
        d_865_v29_: _dafny.Map
        d_865_v29_ = _dafny.Map({False: (self.f19)[default__.safeIndex(575, (self.f19).length(0))]})
        d_866_v30_: _dafny.Seq
        d_866_v30_ = _dafny.SeqWithoutIsStrInference([(self).f20])
        r1 = ((default__.fm33(p2, ((d_865_v29_)[(self).f30] if ((self).f30) in (d_865_v29_) else (self).f18), globalState)) in ((self).f29) if not((d_861_v25_).isdisjoint(((d_863_v27_)[d_862_v26_] if (d_862_v26_) in (d_863_v27_) else _dafny.Set({p0, p2, p2, ((d_864_v28_)[(self).f30] if ((self).f30) in (d_864_v28_) else p2), p0})))) else (409) in (d_866_v30_))
        r2 = (self).f20
        return r0, r1, r2

    @property
    def f29(self):
        return self._f29
    @property
    def f30(self):
        return self._f30

class C6(T2, T1, T0):
    def  __init__(self):
        self._f17: _dafny.Set = _dafny.Set({})
        self._f19: _dafny.Array = _dafny.Array(None, 0)
        self._f21: _dafny.Map = _dafny.Map({})
        self._f20: int = int(0)
        self._f18: bool = False
        self._f28: int = int(0)
        self._f27: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f17(self):
        return self._f17
    @f17.setter
    def f17(self, value):
        self._f17 = value
    @property
    def f19(self):
        return self._f19
    @f19.setter
    def f19(self, value):
        self._f19 = value
    @property
    def f21(self):
        return self._f21
    @f21.setter
    def f21(self, value):
        self._f21 = value
    @property
    def f20(self):
        return self._f20
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f27, f28, f20, f21, f19, f17, f18):
        (self)._f27 = f27
        (self)._f28 = f28
        (self)._f20 = f20
        (self).f21 = f21
        (self).f19 = f19
        (self).f17 = f17
        (self)._f18 = f18

    def fm1(self, p0, globalState):
        return not(True)

    def fm0(self, p0, p1, globalState):
        return (self).f28

    def fm27(self, globalState):
        return (self).f18

    def fm28(self, p0, p1, p2, globalState):
        return (self).f20

    def m0(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        hi4_ = (self).f28
        for d_867_i0_ in range((0) - ((p2 if p1 else p0)), hi4_):
            d_868_v0_: _dafny.Seq
            d_868_v0_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2])
            d_869_v1_: _dafny.Map
            d_869_v1_ = _dafny.Map({not(((self).f28) != (d_867_i0_)): (d_868_v0_) + (d_868_v0_)})
            d_869_v1_ = (d_869_v1_) | (d_869_v1_)
            d_870_v2_: _dafny.Seq
            d_870_v2_ = d_868_v0_
            d_871_v3_: _dafny.Seq
            d_871_v3_ = _dafny.SeqWithoutIsStrInference([(d_870_v2_), d_868_v0_])
            d_872_v4_: _dafny.Seq
            d_872_v4_ = _dafny.SeqWithoutIsStrInference([p1])
            rhs124_ = default__.safeDivisionInt(default__.safeDivisionInt(d_867_i0_, (self).f28), (len(d_871_v3_)) - (len(d_872_v4_)))
            rhs125_ = len(d_868_v0_)
            r2 = rhs124_
            r2 = rhs125_
            d_873_v5_: D0
            d_873_v5_ = D0_DC0((self).f18)
            (globalState).f3 = (d_873_v5_).cf0
            d_874_v6_: _dafny.Map
            d_874_v6_ = _dafny.Map({p1: p1})
            d_875_v7_: _dafny.Seq
            d_875_v7_ = _dafny.SeqWithoutIsStrInference([d_874_v6_, d_874_v6_, (d_874_v6_).set(p1, p1)])
            source9_ = D3_DC8(d_875_v7_)
            if source9_.is_DC9:
                d_876_v8_: _dafny.Map
                d_876_v8_ = _dafny.Map({default__.fm29(globalState): ((self).f20) * (73)})
                d_876_v8_ = (d_876_v8_).set(_dafny.MultiSet([True, p1]), (self).f20)
                d_877_v9_: _dafny.Map
                d_877_v9_ = _dafny.Map({d_872_v4_: default__.fm30((self).f18, globalState)})
                (globalState).f14 = (((self).fm28((self).f20, 784, d_877_v9_, globalState)) + ((self).f20)) + ((self).f20)
                d_878_v10_: C2
                nw147_ = C2()
                nw147_.ctor__(self.f17, (self).f18)
                d_878_v10_ = nw147_
                d_879_v11_: _dafny.Seq
                d_879_v11_ = _dafny.SeqWithoutIsStrInference([d_878_v10_])
                (globalState).f9 = (len(d_879_v11_)) >= (((0) - (p2)) * ((self).fm28((self).fm28(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrjchpu"))), d_877_v9_, globalState), (self).f28, d_877_v9_, globalState)))
                d_880_v12_: str
                d_880_v12_ = _dafny.CodePoint('i')
                d_881_v13_: _dafny.Seq
                d_881_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jovhhqucn"))
                d_882_v14_: _dafny.Seq
                d_882_v14_ = _dafny.SeqWithoutIsStrInference([d_880_v12_, d_880_v12_, (d_881_v13_)[default__.safeIndex(867, len(d_881_v13_))]])
                d_881_v13_ = d_881_v13_
            elif source9_.is_DC10:
                d_883___mcc_h0_ = source9_.cf16
                d_884___mcc_h1_ = source9_.cf17
                d_885___mcc_h2_ = source9_.cf18
                d_886___mcc_h3_ = source9_.cf19
                d_887___mcc_h4_ = source9_.cf20
                d_888_cf20_ = d_887___mcc_h4_
                d_889_cf19_ = d_886___mcc_h3_
                d_890_cf18_ = d_885___mcc_h2_
                d_891_cf17_ = d_884___mcc_h1_
                d_892_cf16_ = d_883___mcc_h0_
                d_893_v15_: C1
                nw148_ = C1()
                nw148_.ctor__()
                d_893_v15_ = nw148_
                (globalState).f9 = True
                d_894_v16_: _dafny.Set
                d_894_v16_ = _dafny.Set({(self).f20})
                arr13_ = self.f19
                index145_ = default__.safeIndex(497, (self.f19).length(0))
                def iife79_():
                    coll43_ = _dafny.Set()
                    compr_43_: int
                    for compr_43_ in _dafny.IntegerRange(-267, 467):
                        d_895_v17_: int = compr_43_
                        if ((-267) <= (d_895_v17_)) and ((d_895_v17_) < (467)):
                            coll43_ = coll43_.union(_dafny.Set([default__.safeDivisionInt(d_895_v17_, (self).f20)]))
                    return _dafny.Set(coll43_)
                arr13_[index145_] = (d_894_v16_).isdisjoint(iife79_()
                )
                arr14_ = self.f19
                index146_ = default__.safeIndex(497, (self.f19).length(0))
                arr14_[index146_] = d_891_cf17_
                d_896_v18_: _dafny.Map
                d_896_v18_ = _dafny.Map({d_867_i0_: p0})
                d_897_v19_: _dafny.Map
                d_897_v19_ = _dafny.Map({d_896_v18_: p2})
                (globalState).f3 = (self).fm1(d_897_v19_, globalState)
            elif True:
                d_898___mcc_h5_ = source9_.cf15
                d_899_cf15_ = d_898___mcc_h5_
                d_900_v20_: str
                d_900_v20_ = _dafny.CodePoint('s')
                d_901_v21_: _dafny.Set
                d_901_v21_ = _dafny.Set({d_900_v20_})
                d_902_v22_: _dafny.Seq
                d_902_v22_ = _dafny.SeqWithoutIsStrInference([D3_DC9(), default__.fm31(globalState)])
                d_903_v23_: _dafny.MultiSet
                d_903_v23_ = _dafny.MultiSet([(D5_DC13((self).f18, (self).f28, d_902_v22_)).cf24])
                rhs126_ = p1
                rhs127_ = (d_901_v21_).issubset(d_901_v21_)
                rhs128_ = not(p1)
                rhs129_ = default__.safeModuloInt(((d_903_v23_)[d_867_i0_] if (d_867_i0_) in (d_903_v23_) else (self).f20), (d_867_i0_) * (p0))
                lhs115_ = globalState
                lhs116_ = globalState
                lhs117_ = globalState
                lhs115_.f9 = rhs126_
                r1 = rhs127_
                lhs116_.f3 = rhs128_
                lhs117_.f14 = rhs129_
                d_904_v24_: T1
                nw149_ = C3()
                nw149_.ctor__(self.f19, self.f17, p1)
                d_904_v24_ = nw149_
                d_905_v25_: _dafny.Set
                d_905_v25_ = _dafny.Set({d_904_v24_, d_904_v24_})
                d_905_v25_ = d_905_v25_
                d_906_v26_: D9
                d_906_v26_ = D9_DC19(default__.fm32(globalState))
                r2 = (self).fm28((self).f20, ((self).f28) - (144), (d_906_v26_).cf35, globalState)
                d_907_v27_: _dafny.Set
                d_907_v27_ = _dafny.Set({p0, -324, p0})
                d_908_v28_: D10
                d_908_v28_ = D10_DC22(d_907_v27_)
                def iife80_():
                    coll44_ = _dafny.Set()
                    compr_44_: int
                    for compr_44_ in (d_907_v27_).Elements:
                        d_909_v29_: int = compr_44_
                        if (d_909_v29_) in (d_907_v27_):
                            coll44_ = coll44_.union(_dafny.Set([(d_909_v29_) + (len(_dafny.SeqWithoutIsStrInference([not(True), not(False)])))]))
                    return _dafny.Set(coll44_)
                d_907_v27_ = (((d_908_v28_).cf42) | (d_907_v27_)).intersection(iife80_()
                )
        d_910_i1_: int
        d_910_i1_ = 0
        with _dafny.label("7"):
            while p1:
                with _dafny.c_label("7"):
                    if (d_910_i1_) >= (100):
                        raise _dafny.Break("7")
                    d_910_i1_ = (d_910_i1_) + (1)
                    d_911_v30_: C1
                    nw150_ = C1()
                    nw150_.ctor__()
                    d_911_v30_ = nw150_
                    d_912_v31_: _dafny.Seq
                    d_912_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ert"))
                    d_913_v32_: _dafny.Map
                    d_913_v32_ = _dafny.Map({118: (self).f28})
                    d_914_v33_: str
                    d_914_v33_ = _dafny.CodePoint('c')
                    d_915_v34_: _dafny.Array
                    nw151_ = _dafny.Array(None, 26)
                    nw151_[int(0)] = p0
                    nw151_[int(1)] = (self).f20
                    nw151_[int(2)] = (self).f28
                    nw151_[int(3)] = (self).f20
                    nw151_[int(4)] = p2
                    nw151_[int(5)] = 887
                    nw151_[int(6)] = len(d_912_v31_)
                    nw151_[int(7)] = len((self).f27)
                    nw151_[int(8)] = p2
                    nw151_[int(9)] = (self).f28
                    nw151_[int(10)] = 548
                    nw151_[int(11)] = p2
                    nw151_[int(12)] = p2
                    nw151_[int(13)] = (((self).f27)[(self).f18] if ((self).f18) in ((self).f27) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkdr"))))
                    nw151_[int(14)] = (self).f28
                    nw151_[int(15)] = len(d_913_v32_)
                    nw151_[int(16)] = p0
                    nw151_[int(17)] = (self).f28
                    nw151_[int(18)] = p2
                    nw151_[int(19)] = (self).f20
                    nw151_[int(20)] = (self).fm0((self).f20, d_914_v33_, globalState)
                    nw151_[int(21)] = (self).f28
                    nw151_[int(22)] = (self).f20
                    nw151_[int(23)] = (self).f20
                    nw151_[int(24)] = p0
                    nw151_[int(25)] = (self).f20
                    d_915_v34_ = nw151_
                    d_916_v35_: _dafny.Array
                    nw152_ = _dafny.Array(None, 10)
                    nw152_[int(0)] = d_915_v34_
                    nw152_[int(1)] = d_915_v34_
                    nw152_[int(2)] = d_915_v34_
                    nw152_[int(3)] = d_915_v34_
                    nw152_[int(4)] = d_915_v34_
                    nw152_[int(5)] = d_915_v34_
                    nw152_[int(6)] = d_915_v34_
                    nw152_[int(7)] = d_915_v34_
                    nw152_[int(8)] = d_915_v34_
                    nw152_[int(9)] = (d_915_v34_ if False else d_915_v34_)
                    d_916_v35_ = nw152_
                    d_916_v35_ = d_916_v35_
                    d_917_v36_: _dafny.Map
                    d_917_v36_ = _dafny.Map({p2: p1})
                    d_917_v36_ = (d_917_v36_).set(len(d_912_v31_), True)
                    d_918_v37_: _dafny.Seq
                    d_918_v37_ = _dafny.SeqWithoutIsStrInference([True, True])
                    d_919_v38_: _dafny.Set
                    d_919_v38_ = _dafny.Set({((self).fm0(len(d_918_v37_), d_914_v33_, globalState)) - (len(d_917_v36_))})
                    d_920_v39_: D10
                    d_920_v39_ = D10_DC22(d_919_v38_)
                    d_921_v40_: _dafny.Map
                    d_921_v40_ = _dafny.Map({False: (d_920_v39_).cf42})
                    d_919_v38_ = ((d_921_v40_)[(self).f18] if ((self).f18) in (d_921_v40_) else d_919_v38_)
                    pass
            pass
        d_922_v41_: _dafny.Seq
        d_922_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ky"))
        d_923_v42_: str
        d_923_v42_ = _dafny.CodePoint('m')
        d_924_v43_: D11
        d_924_v43_ = D11_DC25(self.f17)
        d_925_v44_: C4
        nw153_ = C4()
        nw153_.ctor__(default__.safeDivisionInt(len(d_922_v41_), len((d_922_v41_).set(default__.safeIndex(-751, len(d_922_v41_)), d_923_v42_))), p2, (d_924_v43_).cf47, ((self).f28) == (-569))
        d_925_v44_ = nw153_
        if p1:
            d_926_v45_: _dafny.Array
            nw154_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_926_v45_ = nw154_
            d_927_v46_: _dafny.Array
            nw155_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
            d_927_v46_ = nw155_
            index147_ = default__.safeIndex(316, (d_926_v45_).length(0))
            (d_926_v45_)[index147_] = d_927_v46_
            d_928_v47_: _dafny.Map
            d_928_v47_ = _dafny.Map({p1: d_927_v46_})
            index148_ = default__.safeIndex(316, (d_926_v45_).length(0))
            (d_926_v45_)[index148_] = ((d_928_v47_)[p1] if (p1) in (d_928_v47_) else d_927_v46_)
            d_929_v48_: _dafny.MultiSet
            d_929_v48_ = _dafny.MultiSet([(self).f20])
            (globalState).f15 = d_929_v48_
            (d_925_v44_).f26 = (0) - (p2)
            d_927_v46_ = d_927_v46_
            d_930_v49_: _dafny.Map
            d_930_v49_ = _dafny.Map({d_925_v44_.f26: d_925_v44_.f26})
            d_931_v50_: _dafny.Map
            d_931_v50_ = _dafny.Map({d_930_v49_: (self).f20})
            d_932_v51_: _dafny.MultiSet
            d_932_v51_ = _dafny.MultiSet([(self).f18])
            d_933_v52_: T2
            nw156_ = C5()
            nw156_.ctor__(d_922_v41_, (self).fm1((d_931_v50_).set(d_930_v49_, ((d_932_v51_)[(self).f18] if ((self).f18) in (d_932_v51_) else d_925_v44_.f26)), globalState), (self).f28, self.f21, (self.f19 if p1 else self.f19), (self.f17) | (self.f17), (self).f18)
            d_933_v52_ = nw156_
            d_934_v53_: _dafny.Map
            d_934_v53_ = _dafny.Map({(d_933_v52_).f18: p1})
            arr15_ = d_933_v52_.f19
            index149_ = default__.safeIndex(449, (d_933_v52_.f19).length(0))
            arr15_[index149_] = (p0) < (len(d_934_v53_))
            d_935_v54_: _dafny.Set
            d_935_v54_ = _dafny.Set({(self).f18})
            d_936_v55_: D10
            d_936_v55_ = D10_DC23(not((d_933_v52_).f18), p1, d_933_v52_)
            arr16_ = d_933_v52_.f19
            index150_ = default__.safeIndex(449, (d_933_v52_.f19).length(0))
            rhs130_ = 130
            rhs131_ = 986
            rhs132_ = (d_936_v55_).cf45
            rhs133_ = not(((d_933_v52_).f18) or (((0) - (p0)) == (p0)))
            rhs134_ = _dafny.Set({(d_933_v52_).f18})
            lhs118_ = globalState
            lhs119_ = globalState
            lhs120_ = d_933_v52_.f19
            lhs121_ = default__.safeIndex(449, (d_933_v52_.f19).length(0))
            lhs118_.f14 = rhs130_
            lhs119_.f14 = rhs131_
            d_933_v52_ = rhs132_
            lhs120_[lhs121_] = rhs133_
            d_935_v54_ = rhs134_
        elif True:
            d_922_v41_ = d_922_v41_
            d_937_v56_: _dafny.Array
            def lambda53_(d_938_i2_):
                return _dafny.CodePoint('e')

            init32_ = lambda53_
            nw157_ = _dafny.Array(None, 19)
            for i0_32_ in range(nw157_.length(0)):
                nw157_[i0_32_] = init32_(i0_32_)
            d_937_v56_ = nw157_
            (globalState).f16 = d_937_v56_
            r2 = (0) - (p0)
            (globalState).f9 = p1
            d_939_v57_: _dafny.Seq
            d_939_v57_ = _dafny.SeqWithoutIsStrInference([(self).f18, p1])
            d_940_v58_: D3
            d_940_v58_ = D3_DC9()
            d_941_v59_: _dafny.Map
            d_941_v59_ = _dafny.Map({(d_939_v57_) + (d_939_v57_): default__.fm34(d_940_v58_, p1, d_923_v42_, (self).f18, globalState)})
            d_922_v41_ = ((d_941_v59_)[(d_939_v57_) + (d_939_v57_)] if ((d_939_v57_) + (d_939_v57_)) in (d_941_v59_) else _dafny.SeqWithoutIsStrInference([d_923_v42_ for d_942_i3_ in range(default__.abs(243))]))
        d_943_v60_: _dafny.Seq
        d_943_v60_ = _dafny.SeqWithoutIsStrInference([(d_925_v44_).f25])
        d_944_v62_: _dafny.Seq
        def iife81_():
            coll45_ = _dafny.Map()
            compr_45_: int
            for compr_45_ in _dafny.IntegerRange(-896, 31):
                d_945_v61_: int = compr_45_
                if ((-896) <= (d_945_v61_)) and ((d_945_v61_) < (31)):
                    coll45_[default__.safeDivisionInt(d_945_v61_, (_dafny.MultiSet([810, 797])).cardinality)] = d_943_v60_
            return _dafny.Map(coll45_)
        d_944_v62_ = _dafny.SeqWithoutIsStrInference([len((d_943_v60_).set(default__.safeIndex(len(iife81_()
        ), len(d_943_v60_)), p2)), (self).f20, p2])
        d_944_v62_ = _dafny.SeqWithoutIsStrInference([p0 for d_946_i4_ in range(default__.abs(-36))])
        if (self).f18:
            r1 = (self).f18
            d_947_v63_: _dafny.Seq
            d_947_v63_ = _dafny.SeqWithoutIsStrInference([(self).f18, p1, (self).f18, (self).f18, not((self).f18)])
            rhs135_ = False
            rhs136_ = ((self).f18) or (not((d_947_v63_) == (d_947_v63_)))
            rhs137_ = not(True)
            lhs122_ = globalState
            lhs123_ = globalState
            r0 = rhs135_
            lhs122_.f9 = rhs136_
            lhs123_.f3 = rhs137_
            d_948_v64_: _dafny.MultiSet
            d_948_v64_ = _dafny.MultiSet([(self).f20])
            d_949_v65_: _dafny.Set
            d_949_v65_ = _dafny.Set({d_925_v44_.f26, (d_925_v44_).f25})
            nw158_ = _dafny.Array(None, 13)
            nw158_[int(0)] = p1
            nw158_[int(1)] = (self).f18
            nw158_[int(2)] = (d_948_v64_).ispropersubset(d_948_v64_)
            nw158_[int(3)] = p1
            nw158_[int(4)] = p1
            nw158_[int(5)] = (self).f18
            nw158_[int(6)] = (self).f18
            nw158_[int(7)] = p1
            nw158_[int(8)] = p1
            nw158_[int(9)] = not((self).f18)
            nw158_[int(10)] = (self).f18
            nw158_[int(11)] = not((_dafny.Set({(self).f20})).isdisjoint(d_949_v65_))
            nw158_[int(12)] = not (not((self).f18)) or ((self).f18)
            (self).f19 = nw158_
            d_950_v66_: bool
            d_951_v67_: _dafny.Map
            d_952_v68_: int
            out47_: bool
            out48_: _dafny.Map
            out49_: int
            out47_, out48_, out49_ = (d_925_v44_).m4(globalState)
            d_950_v66_ = out47_
            d_951_v67_ = out48_
            d_952_v68_ = out49_
            d_953_v69_: _dafny.Array
            def lambda54_(d_954_v44_):
                def lambda55_(d_955_i5_):
                    return (d_955_i5_) * ((d_954_v44_).f25)

                return lambda55_

            init33_ = lambda54_(d_925_v44_)
            nw159_ = _dafny.Array(None, 17)
            for i0_33_ in range(nw159_.length(0)):
                nw159_[i0_33_] = init33_(i0_33_)
            d_953_v69_ = nw159_
            index151_ = default__.safeIndex(945, (d_953_v69_).length(0))
            (d_953_v69_)[index151_] = d_925_v44_.f26
            d_956_v70_: D0
            d_956_v70_ = D0_DC2(d_950_v66_, (d_925_v44_).f25, d_953_v69_)
            index152_ = default__.safeIndex(945, (d_953_v69_).length(0))
            rhs138_ = d_925_v44_.f26
            rhs139_ = 832
            rhs140_ = (d_943_v60_)[default__.safeIndex(default__.safeDivisionInt((self).f28, d_925_v44_.f26), len(d_943_v60_))]
            rhs141_ = (d_956_v70_).cf7
            lhs124_ = d_953_v69_
            lhs125_ = default__.safeIndex(945, (d_953_v69_).length(0))
            lhs126_ = d_925_v44_
            lhs124_[lhs125_] = rhs138_
            lhs126_.f26 = rhs139_
            d_952_v68_ = rhs140_
            d_953_v69_ = rhs141_
        elif True:
            d_957_v71_: _dafny.Array
            nw160_ = _dafny.Array(None, 2)
            d_957_v71_ = nw160_
            d_958_v72_: _dafny.Map
            d_958_v72_ = _dafny.Map({d_957_v71_: (self).f18})
            d_958_v72_ = (d_958_v72_) | (d_958_v72_)
            d_959_v73_: _dafny.Array
            nw161_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
            d_959_v73_ = nw161_
            index153_ = default__.safeIndex(254, (d_959_v73_).length(0))
            (d_959_v73_)[index153_] = (d_922_v41_) + (d_922_v41_)
            index154_ = default__.safeIndex(254, (d_959_v73_).length(0))
            (d_959_v73_)[index154_] = d_922_v41_
            d_960_v74_: _dafny.Array
            def lambda56_(d_961_p2_):
                def lambda57_(d_962_i6_):
                    return (d_962_i6_) - (d_961_p2_)

                return lambda57_

            init34_ = lambda56_(p2)
            nw162_ = _dafny.Array(None, 11)
            for i0_34_ in range(nw162_.length(0)):
                nw162_[i0_34_] = init34_(i0_34_)
            d_960_v74_ = nw162_
            index155_ = default__.safeIndex(477, (d_960_v74_).length(0))
            (d_960_v74_)[index155_] = (d_925_v44_).f25
            index156_ = default__.safeIndex(477, (d_960_v74_).length(0))
            (d_960_v74_)[index156_] = (self).f20
            (globalState).f9 = (self).f18
            index157_ = default__.safeIndex(477, (d_960_v74_).length(0))
            (d_960_v74_)[index157_] = (d_925_v44_.f26) * (default__.safeDivisionInt((self).f28, d_925_v44_.f26))
        r0 = p1
        r1 = (self).f18
        d_963_v75_: _dafny.Seq
        d_963_v75_ = _dafny.SeqWithoutIsStrInference([True])
        r2 = default__.safeDivisionInt((p0) + ((((self).f27)[(self).f18] if ((self).f18) in ((self).f27) else (self).fm0(len(d_963_v75_), d_923_v42_, globalState))), 117)
        return r0, r1, r2

    def m6(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.Seq({})
        r2: _dafny.Seq = _dafny.Seq({})
        d_964_i0_: int
        d_964_i0_ = 0
        with _dafny.label("8"):
            while (757) >= ((self).f20):
                with _dafny.c_label("8"):
                    if (d_964_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_964_i0_ = (d_964_i0_) + (1)
                    index158_ = default__.safeIndex(206, (p1).length(0))
                    (p1)[index158_] = (self).f28
                    d_965_v0_: _dafny.Array
                    nw163_ = _dafny.Array(_dafny.CodePoint('D'), 16)
                    d_965_v0_ = nw163_
                    d_966_v1_: _dafny.Set
                    d_966_v1_ = _dafny.Set({d_965_v0_, d_965_v0_})
                    index159_ = default__.safeIndex(206, (p1).length(0))
                    (p1)[index159_] = default__.safeModuloInt(698, len(d_966_v1_))
                    d_967_v2_: _dafny.Array
                    nw164_ = _dafny.Array(_dafny.Set({}), 4)
                    d_967_v2_ = nw164_
                    d_968_v3_: _dafny.Map
                    d_968_v3_ = _dafny.Map({d_967_v2_: _dafny.Map({p0: (p1)[default__.safeIndex(206, (p1).length(0))]})})
                    d_969_v4_: _dafny.Map
                    d_969_v4_ = _dafny.Map({p0: (p1)[default__.safeIndex(206, (p1).length(0))]})
                    d_968_v3_ = (d_968_v3_).set(d_967_v2_, (_dafny.Map({p3: (self).f20})) | (d_969_v4_))
                    d_970_v5_: _dafny.Map
                    d_970_v5_ = _dafny.Map({p0: (self).f18})
                    r2 = _dafny.SeqWithoutIsStrInference([(p0) <= (len(d_970_v5_)), (self).f18])
                    d_971_v6_: _dafny.Seq
                    d_971_v6_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_972_v7_: C4
                    nw165_ = C4()
                    nw165_.ctor__(p3, default__.safeModuloInt((p1)[default__.safeIndex(206, (p1).length(0))], (self).f28), self.f17, not ((self).f18) or ((D9_DC20(d_970_v5_, len(d_971_v6_), (self).f18, _dafny.Map({p0: p0}), _dafny.CodePoint('h'))).cf38))
                    d_972_v7_ = nw165_
                    pass
            pass
        d_973_i1_: int
        d_973_i1_ = 0
        with _dafny.label("9"):
            while (p0) < ((self).f28):
                with _dafny.c_label("9"):
                    if (d_973_i1_) >= (100):
                        raise _dafny.Break("9")
                    d_973_i1_ = (d_973_i1_) + (1)
                    d_974_v8_: _dafny.Seq
                    d_974_v8_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])
                    d_975_v9_: _dafny.Map
                    d_975_v9_ = _dafny.Map({d_974_v8_: _dafny.Map({len(d_974_v8_): (self).f28})})
                    d_976_v10_: _dafny.Map
                    d_976_v10_ = _dafny.Map({(((self).f27)[False] if (False) in ((self).f27) else p0): (0) - ((self).fm28(p3, 426, d_975_v9_, globalState))})
                    d_977_v11_: _dafny.Seq
                    d_977_v11_ = _dafny.SeqWithoutIsStrInference([len(d_976_v10_), (self).f28, p0])
                    d_978_v12_: _dafny.Seq
                    d_978_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bimux"))
                    source10_ = default__.fm35(default__.safeModuloInt(p0, (self).f20), (_dafny.SeqWithoutIsStrInference([p0 for d_979_i2_ in range(default__.abs(672))])) + (d_977_v11_), len(d_978_v12_), globalState)
                    d_980___mcc_h0_ = source10_
                    d_981_cf21_ = d_980___mcc_h0_
                    d_982_v13_: D2
                    d_982_v13_ = D2_DC6((self).f28)
                    (globalState).f14 = (p3) + ((d_982_v13_).cf13)
                    d_978_v12_ = d_978_v12_
                    d_983_v14_: _dafny.Map
                    d_983_v14_ = _dafny.Map({not((self).f18): (self).f18})
                    (globalState).f14 = default__.safeModuloInt((len(d_983_v14_)) + ((self).f20), (len((self).f27)) - ((self).f28))
                    arr17_ = self.f19
                    index160_ = default__.safeIndex(88, (self.f19).length(0))
                    arr17_[index160_] = False
                    d_984_v15_: _dafny.MultiSet
                    d_984_v15_ = _dafny.MultiSet([(self).f18, (self).f18, (self).f18, not((self).f18), (self).f18])
                    arr18_ = self.f19
                    index161_ = default__.safeIndex(88, (self.f19).length(0))
                    arr18_[index161_] = (d_974_v8_)[default__.safeIndex((d_984_v15_).cardinality, len(d_974_v8_))]
                    d_985_v16_: T0
                    nw166_ = C2()
                    nw166_.ctor__(self.f17, (self).f18)
                    d_985_v16_ = nw166_
                    d_986_v17_: _dafny.Map
                    d_986_v17_ = _dafny.Map({d_978_v12_: d_985_v16_})
                    (globalState).f14 = (p3 if (self).f18 else default__.safeDivisionInt(len(d_986_v17_), (self).f28))
                    d_987_v18_: str
                    d_987_v18_ = _dafny.CodePoint('x')
                    d_988_v19_: D11
                    d_988_v19_ = D11_DC27(_dafny.SeqWithoutIsStrInference([True, (d_985_v16_).f18]), (0) - ((0) - ((self).fm0((self).f20, d_987_v18_, globalState))), 334)
                    pat_let_tv10_ = d_974_v8_
                    pat_let_tv11_ = d_974_v8_
                    d_989_v20_: _dafny.Map
                    def iife82_(_pat_let18_0):
                        def iife83_(d_990_dt__update__tmp_h0_):
                            def iife84_(_pat_let19_0):
                                def iife85_(d_991_dt__update_hcf52_h0_):
                                    def iife86_(_pat_let20_0):
                                        def iife87_(d_992_dt__update_hcf50_h0_):
                                            return D11_DC27(d_992_dt__update_hcf50_h0_, (d_990_dt__update__tmp_h0_).cf51, d_991_dt__update_hcf52_h0_)
                                        return iife87_(_pat_let20_0)
                                    return iife86_(pat_let_tv11_)
                                return iife85_(_pat_let19_0)
                            return iife84_(len(pat_let_tv10_))
                        return iife83_(_pat_let18_0)
                    d_989_v20_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).fm27(globalState)]): iife82_(d_988_v19_)})
                    d_989_v20_ = (d_989_v20_).set(d_974_v8_, d_988_v19_)
                    d_993_v21_: _dafny.MultiSet
                    d_993_v21_ = _dafny.MultiSet([(self).f20])
                    d_994_v22_: _dafny.Seq
                    d_994_v22_ = _dafny.SeqWithoutIsStrInference([d_993_v21_])
                    d_994_v22_ = d_994_v22_
                    pass
            pass
        if (self).f18:
            if (self).f18:
                d_995_v23_: _dafny.Seq
                d_995_v23_ = _dafny.SeqWithoutIsStrInference([p1])
                (globalState).f7 = (d_995_v23_)[default__.safeIndex(137, len(d_995_v23_))]
                d_996_v24_: _dafny.Map
                d_996_v24_ = _dafny.Map({(self).f18: self.f19})
                (globalState).f14 = len((d_996_v24_) | ((d_996_v24_) | (d_996_v24_)))
                d_997_v25_: _dafny.Seq
                d_997_v25_ = _dafny.SeqWithoutIsStrInference([106, p3, (0) - (p0)])
                r1 = (d_997_v25_) + (d_997_v25_)
                d_998_v26_: _dafny.Seq
                d_998_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kfekrwy"))
                d_998_v26_ = d_998_v26_
                index162_ = default__.safeIndex(517, (p1).length(0))
                (p1)[index162_] = p3
                d_999_v27_: str
                d_999_v27_ = _dafny.CodePoint('k')
                index163_ = default__.safeIndex(517, (p1).length(0))
                (p1)[index163_] = (self).fm0((self).f20, d_999_v27_, globalState)
            elif True:
                d_1000_v28_: str
                d_1000_v28_ = _dafny.CodePoint('d')
                d_1001_v29_: _dafny.Seq
                d_1001_v29_ = _dafny.SeqWithoutIsStrInference([d_1000_v28_, d_1000_v28_, d_1000_v28_, _dafny.CodePoint('m')])
                d_1002_v30_: _dafny.Array
                nw167_ = _dafny.Array(None, 15)
                nw167_[int(0)] = d_1000_v28_
                nw167_[int(1)] = d_1000_v28_
                nw167_[int(2)] = d_1000_v28_
                nw167_[int(3)] = d_1000_v28_
                nw167_[int(4)] = default__.fm33(p3, (self).f18, globalState)
                nw167_[int(5)] = d_1000_v28_
                nw167_[int(6)] = d_1000_v28_
                nw167_[int(7)] = (d_1001_v29_)[default__.safeIndex((self).f28, len(d_1001_v29_))]
                nw167_[int(8)] = d_1000_v28_
                nw167_[int(9)] = d_1000_v28_
                nw167_[int(10)] = d_1000_v28_
                nw167_[int(11)] = d_1000_v28_
                nw167_[int(12)] = d_1000_v28_
                nw167_[int(13)] = d_1000_v28_
                nw167_[int(14)] = _dafny.CodePoint('l')
                d_1002_v30_ = nw167_
                (globalState).f16 = d_1002_v30_
                d_1003_v31_: _dafny.Seq
                d_1003_v31_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1004_v32_: _dafny.Map
                d_1004_v32_ = _dafny.Map({_dafny.Map({(self).f20: len(d_1003_v31_)}): p3})
                d_1000_v28_ = default__.fm33((p3) + (p0), (self).fm1(d_1004_v32_, globalState), globalState)
                (globalState).f14 = (0) - (p0)
                d_1005_v33_: _dafny.MultiSet
                d_1005_v33_ = _dafny.MultiSet([len((self).f27)])
                d_1006_v34_: _dafny.Seq
                d_1006_v34_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet((d_1003_v31_).set(default__.safeIndex((self).f28, len(d_1003_v31_)), (self).f20)), _dafny.MultiSet([len(d_1001_v29_)]), d_1005_v33_, (d_1005_v33_) - (d_1005_v33_), _dafny.MultiSet([500])])
                d_1006_v34_ = d_1006_v34_
                (globalState).f14 = (0) - ((0) - ((((self).f20) - (p3)) - ((self).f20)))
            d_1007_v35_: D3
            d_1007_v35_ = D3_DC9()
            d_1008_v36_: _dafny.Seq
            d_1008_v36_ = _dafny.SeqWithoutIsStrInference([d_1007_v35_])
            d_1009_v37_: D5
            d_1009_v37_ = D5_DC13((self).f18, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhjby"))), d_1008_v36_)
            (globalState).f9 = (d_1009_v37_).cf23
            r0 = (0) - ((self).f20)
            d_1010_v38_: _dafny.Array
            nw168_ = _dafny.Array(_dafny.Map({}), 4)
            d_1010_v38_ = nw168_
            d_1011_v39_: _dafny.Map
            d_1011_v39_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ch")): p1})
            index164_ = default__.safeIndex(596, (d_1010_v38_).length(0))
            (d_1010_v38_)[index164_] = d_1011_v39_
            d_1012_v40_: _dafny.Seq
            d_1012_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqjs"))
            d_1013_v41_: D0
            d_1013_v41_ = D0_DC2((self).f18, len(d_1012_v40_), p1)
            index165_ = default__.safeIndex(596, (d_1010_v38_).length(0))
            (d_1010_v38_)[index165_] = _dafny.Map({(d_1012_v40_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dsssjt"))): (d_1013_v41_).cf7})
            if (self).f18:
                (globalState).f14 = p3
                rhs142_ = (self).f28
                rhs143_ = p1
                rhs144_ = (self).f18
                rhs145_ = (self).f18
                lhs127_ = globalState
                lhs128_ = globalState
                lhs129_ = globalState
                r0 = rhs142_
                lhs127_.f7 = rhs143_
                lhs128_.f9 = rhs144_
                lhs129_.f9 = rhs145_
                d_1012_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydy"))
                d_1014_v42_: _dafny.Seq
                d_1014_v42_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])
                d_1015_v43_: _dafny.Map
                d_1015_v43_ = _dafny.Map({self.f19: d_1014_v42_})
                rhs146_ = d_1015_v43_
                rhs147_ = (self).f28
                rhs148_ = (self).f18
                lhs130_ = globalState
                lhs131_ = globalState
                d_1015_v43_ = rhs146_
                lhs130_.f14 = rhs147_
                lhs131_.f3 = rhs148_
                d_1016_v44_: _dafny.Array
                nw169_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_1016_v44_ = nw169_
                d_1016_v44_ = d_1016_v44_
            elif True:
                (globalState).f14 = p0
                (globalState).f9 = ((self).f18 if (self).f18 else True)
                d_1017_v45_: str
                d_1017_v45_ = _dafny.CodePoint('s')
                (self).f21 = (_dafny.Map({d_1017_v45_: (self).f18})) | (self.f21)
                index166_ = default__.safeIndex(290, (p1).length(0))
                (p1)[index166_] = (self).f20
                index167_ = default__.safeIndex(290, (p1).length(0))
                (p1)[index167_] = len((d_1012_v40_).set(default__.safeIndex((self).f20, len(d_1012_v40_)), d_1017_v45_))
                d_1018_v46_: _dafny.Array
                nw170_ = _dafny.Array(None, 21)
                nw170_[int(0)] = d_1017_v45_
                nw170_[int(1)] = d_1017_v45_
                nw170_[int(2)] = d_1017_v45_
                nw170_[int(3)] = d_1017_v45_
                nw170_[int(4)] = d_1017_v45_
                nw170_[int(5)] = d_1017_v45_
                nw170_[int(6)] = d_1017_v45_
                nw170_[int(7)] = d_1017_v45_
                nw170_[int(8)] = d_1017_v45_
                nw170_[int(9)] = _dafny.CodePoint('a')
                nw170_[int(10)] = d_1017_v45_
                nw170_[int(11)] = d_1017_v45_
                nw170_[int(12)] = d_1017_v45_
                nw170_[int(13)] = d_1017_v45_
                nw170_[int(14)] = _dafny.CodePoint('v')
                nw170_[int(15)] = (d_1012_v40_)[default__.safeIndex(len(d_1012_v40_), len(d_1012_v40_))]
                nw170_[int(16)] = d_1017_v45_
                nw170_[int(17)] = d_1017_v45_
                nw170_[int(18)] = d_1017_v45_
                nw170_[int(19)] = d_1017_v45_
                nw170_[int(20)] = d_1017_v45_
                d_1018_v46_ = nw170_
                d_1019_v47_: _dafny.Map
                d_1019_v47_ = _dafny.Map({(self).f18: d_1018_v46_})
                d_1020_v48_: _dafny.Array
                nw171_ = _dafny.Array(None, 9)
                nw171_[int(0)] = d_1018_v46_
                nw171_[int(1)] = d_1018_v46_
                nw171_[int(2)] = d_1018_v46_
                nw171_[int(3)] = ((d_1019_v47_)[(self).f18] if ((self).f18) in (d_1019_v47_) else d_1018_v46_)
                nw171_[int(4)] = d_1018_v46_
                nw171_[int(5)] = d_1018_v46_
                nw171_[int(6)] = d_1018_v46_
                nw171_[int(7)] = d_1018_v46_
                nw171_[int(8)] = d_1018_v46_
                d_1020_v48_ = nw171_
                d_1021_v49_: _dafny.Map
                d_1021_v49_ = _dafny.Map({d_1020_v48_: _dafny.Map({len(d_1012_v40_): (self).f20})})
                d_1022_v50_: _dafny.Map
                d_1022_v50_ = _dafny.Map({(self).f20: (self).f20})
                d_1021_v49_ = (d_1021_v49_).set(d_1020_v48_, d_1022_v50_)
        elif True:
            (globalState).f14 = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1023_i3_ in range(default__.abs(168))])))
            d_1024_v51_: _dafny.Seq
            d_1024_v51_ = _dafny.SeqWithoutIsStrInference([True, not((self).f18), not((self).f18), (self).f18, (self).f18])
            d_1025_v52_: _dafny.Map
            d_1025_v52_ = _dafny.Map({d_1024_v51_: False})
            d_1026_v53_: _dafny.Set
            d_1026_v53_ = _dafny.Set({d_1024_v51_})
            def iife88_():
                coll46_ = _dafny.Set()
                compr_46_: _dafny.Seq
                for compr_46_ in (d_1026_v53_).Elements:
                    d_1027_v54_: _dafny.Seq = compr_46_
                    if (d_1027_v54_) in (d_1026_v53_):
                        coll46_ = coll46_.union(_dafny.Set([d_1027_v54_]))
                return _dafny.Set(coll46_)
            (globalState).f14 = (0) - (len((d_1025_v52_).set(d_1024_v51_, not((d_1026_v53_) != (iife88_()
            )))))
            d_1028_v55_: _dafny.Set
            d_1028_v55_ = _dafny.Set({(self).f18, (self).f18, ((self).f18 if (self).f18 else (self).f18), not (False) or ((self).f18)})
            d_1029_v56_: _dafny.MultiSet
            d_1029_v56_ = _dafny.MultiSet([(self).f18, (self).f18, (self).f18])
            d_1030_v57_: _dafny.Map
            d_1030_v57_ = _dafny.Map({(d_1029_v56_).intersection(d_1029_v56_): (self).f18})
            d_1031_v58_: _dafny.Seq
            d_1031_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wupo"))
            d_1032_v59_: _dafny.Map
            d_1032_v59_ = _dafny.Map({p0: d_1031_v58_})
            rhs149_ = not(not((self).f18))
            rhs150_ = p2
            rhs151_ = len((d_1032_v59_) | (default__.fm36(globalState)))
            rhs152_ = d_1030_v57_
            lhs132_ = globalState
            lhs132_.f3 = rhs149_
            d_1028_v55_ = rhs150_
            r0 = rhs151_
            d_1030_v57_ = rhs152_
            d_1033_v60_: T2
            nw172_ = C5()
            nw172_.ctor__(d_1031_v58_, (self).f18, len(d_1024_v51_), self.f21, self.f19, self.f17, (self).f18)
            d_1033_v60_ = nw172_
            d_1034_v61_: D10
            d_1034_v61_ = D10_DC23((self).f18, (self).f18, d_1033_v60_)
            d_1035_v62_: _dafny.Map
            d_1035_v62_ = _dafny.Map({(d_1028_v55_) | (d_1028_v55_): d_1034_v61_})
            d_1035_v62_ = (d_1035_v62_).set(default__.fm37((d_1033_v60_).f18, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yf"))), (d_1033_v60_).f18, (d_1033_v60_).f18, globalState), d_1034_v61_)
            (globalState).f14 = p0
        if (self).f18:
            index168_ = default__.safeIndex(487, (p1).length(0))
            (p1)[index168_] = p0
            d_1036_v63_: _dafny.Map
            d_1036_v63_ = _dafny.Map({(-786) > (len(p2)): 202})
            d_1037_v64_: _dafny.Seq
            d_1037_v64_ = _dafny.SeqWithoutIsStrInference([(self).f18])
            d_1038_v65_: _dafny.Map
            d_1038_v65_ = _dafny.Map({(self).f28: (self).f28})
            d_1039_v66_: _dafny.Map
            d_1039_v66_ = _dafny.Map({(self).f18: d_1038_v65_})
            d_1040_v67_: _dafny.Map
            d_1040_v67_ = _dafny.Map({d_1037_v64_: ((d_1039_v66_)[(self).f18] if ((self).f18) in (d_1039_v66_) else d_1038_v65_)})
            index169_ = default__.safeIndex(487, (p1).length(0))
            rhs153_ = (p3) > ((self).fm28(p3, (self).f20, d_1040_v67_, globalState))
            rhs154_ = (self).f18
            rhs155_ = (0) - (p3)
            rhs156_ = d_1036_v63_
            lhs133_ = globalState
            lhs134_ = globalState
            lhs135_ = p1
            lhs136_ = default__.safeIndex(487, (p1).length(0))
            lhs133_.f9 = rhs153_
            lhs134_.f9 = rhs154_
            lhs135_[lhs136_] = rhs155_
            d_1036_v63_ = rhs156_
            d_1041_v68_: _dafny.Seq
            d_1041_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hifw"))
            d_1042_v69_: _dafny.MultiSet
            d_1042_v69_ = _dafny.MultiSet([D11_DC27(d_1037_v64_, (0) - (p3), len(_dafny.Set({p3})))])
            d_1043_v70_: _dafny.Seq
            d_1043_v70_ = _dafny.SeqWithoutIsStrInference([(d_1042_v69_).intersection(d_1042_v69_), (d_1042_v69_) | (d_1042_v69_)])
            d_1044_v71_: _dafny.Map
            d_1044_v71_ = _dafny.Map({(self).f18: d_1037_v64_})
            rhs157_ = d_1041_v68_
            rhs158_ = (d_1043_v70_)[default__.safeIndex(len(((d_1044_v71_)[not(not((self).f18))] if (not(not((self).f18))) in (d_1044_v71_) else d_1037_v64_)), len(d_1043_v70_))]
            rhs159_ = (self).f18
            lhs137_ = globalState
            d_1041_v68_ = rhs157_
            d_1042_v69_ = rhs158_
            lhs137_.f3 = rhs159_
            arr19_ = self.f19
            index170_ = default__.safeIndex(105, (self.f19).length(0))
            arr19_[index170_] = (self).f18
            arr20_ = self.f19
            index171_ = default__.safeIndex(105, (self.f19).length(0))
            arr20_[index171_] = not((not((self).f18) if ((self).f18 if not((self).f18) else (self).f18) else (self).f18))
            (globalState).f14 = (p0) - ((self).f28)
            d_1045_v72_: _dafny.Array
            nw173_ = _dafny.Array(False, 10)
            d_1045_v72_ = nw173_
            d_1046_v73_: _dafny.Seq
            d_1046_v73_ = _dafny.SeqWithoutIsStrInference([d_1045_v72_, d_1045_v72_, d_1045_v72_])
            d_1046_v73_ = (d_1046_v73_) + (d_1046_v73_)
        elif True:
            d_1047_v74_: C3
            nw174_ = C3()
            nw174_.ctor__(self.f19, self.f17, (self).f18)
            d_1047_v74_ = nw174_
            d_1048_v75_: _dafny.Seq
            d_1048_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikjvinf"))
            d_1049_v76_: D3
            d_1049_v76_ = D3_DC9()
            d_1050_v77_: str
            d_1050_v77_ = _dafny.CodePoint('a')
            d_1048_v75_ = (default__.fm34(d_1049_v76_, (self).f18, d_1050_v77_, (self).f18, globalState)) + ((d_1048_v75_) + (d_1048_v75_))
            d_1051_v78_: _dafny.Seq
            d_1051_v78_ = _dafny.SeqWithoutIsStrInference([True])
            d_1052_v79_: _dafny.MultiSet
            d_1052_v79_ = _dafny.MultiSet([d_1051_v78_, d_1051_v78_])
            d_1053_v80_: _dafny.Map
            d_1053_v80_ = _dafny.Map({d_1049_v76_: ((d_1052_v79_).cardinality) <= ((d_1047_v74_).fm0(p0, d_1050_v77_, globalState))})
            d_1053_v80_ = (d_1053_v80_).set(d_1049_v76_, (self).f18)
            d_1054_v81_: _dafny.Set
            d_1054_v81_ = _dafny.Set({p3, (0) - (default__.safeDivisionInt(-98, (self).f28)), len(_dafny.SeqWithoutIsStrInference([((self).f27).set((self).f18, (self).f28)]))})
            d_1054_v81_ = (d_1054_v81_) | (d_1054_v81_)
            if (d_1051_v78_)[default__.safeIndex((self).f20, len(d_1051_v78_))]:
                d_1055_v82_: _dafny.MultiSet
                d_1055_v82_ = _dafny.MultiSet([_dafny.MultiSet([(self).f18])])
                d_1056_v83_: _dafny.Map
                d_1056_v83_ = _dafny.Map({p0: (self).f18})
                d_1057_v84_: _dafny.Seq
                d_1057_v84_ = _dafny.SeqWithoutIsStrInference([default__.fm38((d_1055_v82_).cardinality, (self).f18, globalState), (d_1056_v83_).set(273, True), _dafny.Map({(self).f28: (self).f18}), ((_dafny.Map({len((self).f27): True})).set((self).f20, (self).fm27(globalState))).set((self).f28, (self).f18), d_1056_v83_])
                d_1058_v86_: _dafny.Map
                def iife89_():
                    coll47_ = _dafny.Map()
                    compr_47_: int
                    for compr_47_ in (_dafny.SeqWithoutIsStrInference([p0, (self).f28])).Elements:
                        d_1059_v85_: int = compr_47_
                        if (d_1059_v85_) in (_dafny.SeqWithoutIsStrInference([p0, (self).f28])):
                            coll47_[default__.safeModuloInt(d_1059_v85_, (self).f28)] = 218
                    return _dafny.Map(coll47_)
                d_1058_v86_ = _dafny.Map({iife89_()
                : len(d_1048_v75_)})
                d_1057_v84_ = ((d_1057_v84_) + (d_1057_v84_) if ((self).f18 if (self).f18 else (self).fm1(d_1058_v86_, globalState)) else (d_1057_v84_ if (self).f18 else d_1057_v84_))
                arr21_ = self.f19
                index172_ = default__.safeIndex(923, (self.f19).length(0))
                arr21_[index172_] = (self).f18
                arr22_ = self.f19
                index173_ = default__.safeIndex(923, (self.f19).length(0))
                arr22_[index173_] = ((self).fm27(globalState)) not in (default__.fm39(p3, globalState))
                d_1060_v87_: _dafny.MultiSet
                d_1060_v87_ = _dafny.MultiSet([(self).f18, (self).f18, (self).f18])
                arr23_ = self.f19
                index174_ = default__.safeIndex(923, (self.f19).length(0))
                rhs160_ = (d_1060_v87_) != (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([True])) + (d_1051_v78_)))
                rhs161_ = (self).fm27(globalState)
                rhs162_ = p1
                lhs138_ = globalState
                lhs139_ = self.f19
                lhs140_ = default__.safeIndex(923, (self.f19).length(0))
                lhs141_ = globalState
                lhs138_.f9 = rhs160_
                lhs139_[lhs140_] = rhs161_
                lhs141_.f7 = rhs162_
                d_1061_v88_: _dafny.Map
                d_1061_v88_ = _dafny.Map({(self.f19)[default__.safeIndex(923, (self.f19).length(0))]: True})
                rhs163_ = p1
                rhs164_ = not(((d_1061_v88_)[(not((self.f19)[default__.safeIndex(923, (self.f19).length(0))])) or ((self).f18)] if ((not((self.f19)[default__.safeIndex(923, (self.f19).length(0))])) or ((self).f18)) in (d_1061_v88_) else (self).f18))
                rhs165_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfq"))
                rhs166_ = (d_1050_v77_) in (d_1048_v75_)
                rhs167_ = (self).f18
                lhs142_ = globalState
                lhs143_ = globalState
                lhs144_ = globalState
                lhs145_ = globalState
                lhs142_.f7 = rhs163_
                lhs143_.f3 = rhs164_
                d_1048_v75_ = rhs165_
                lhs144_.f3 = rhs166_
                lhs145_.f9 = rhs167_
                d_1062_v90_: _dafny.Map
                def iife90_():
                    coll48_ = _dafny.Map()
                    compr_48_: int
                    for compr_48_ in (_dafny.Map({(self).f28: (self.f19)[default__.safeIndex(923, (self.f19).length(0))]})).keys.Elements:
                        d_1063_v89_: int = compr_48_
                        if (d_1063_v89_) in (_dafny.Map({(self).f28: (self.f19)[default__.safeIndex(923, (self.f19).length(0))]})):
                            coll48_[(d_1063_v89_) - (p0)] = p3
                    return _dafny.Map(coll48_)
                d_1062_v90_ = _dafny.Map({d_1051_v78_: iife90_()
                })
                rhs168_ = (self).f28
                rhs169_ = (self).f18
                rhs170_ = (self).fm28((self).f20, (p0) - (p0), d_1062_v90_, globalState)
                rhs171_ = (self.f19)[default__.safeIndex(923, (self.f19).length(0))]
                rhs172_ = p3
                lhs146_ = globalState
                lhs147_ = globalState
                lhs148_ = globalState
                lhs149_ = globalState
                lhs150_ = globalState
                lhs146_.f14 = rhs168_
                lhs147_.f3 = rhs169_
                lhs148_.f14 = rhs170_
                lhs149_.f3 = rhs171_
                lhs150_.f14 = rhs172_
            elif True:
                d_1064_v91_: _dafny.Seq
                d_1064_v91_ = _dafny.SeqWithoutIsStrInference([d_1048_v75_])
                d_1065_v92_: _dafny.Map
                d_1065_v92_ = _dafny.Map({(self).f18: (d_1064_v91_) < (d_1064_v91_)})
                d_1065_v92_ = (d_1065_v92_).set((self).f18, ((self).f18 if (self).f18 else (d_1051_v78_)[default__.safeIndex(p3, len(d_1051_v78_))]))
                d_1054_v81_ = d_1054_v81_
                d_1066_v93_: _dafny.Seq
                d_1066_v93_ = _dafny.SeqWithoutIsStrInference([len((_dafny.Set({False, (self).f18, (self).f18})) | (p2)), (self).f20])
                r0 = (d_1066_v93_)[default__.safeIndex(p0, len(d_1066_v93_))]
                (self).f19 = self.f19
                (globalState).f9 = ((0) - ((p3) * ((self).f28))) == (p3)
        d_1067_v94_: C2
        nw175_ = C2()
        nw175_.ctor__((self.f17) - (self.f17), (self).f18)
        d_1067_v94_ = nw175_
        d_1068_v95_: _dafny.Seq
        d_1068_v95_ = _dafny.SeqWithoutIsStrInference([(self).f18])
        d_1069_v96_: _dafny.MultiSet
        d_1069_v96_ = _dafny.MultiSet([(self).f18])
        arr24_ = self.f19
        index175_ = default__.safeIndex(1, (self.f19).length(0))
        arr24_[index175_] = (d_1068_v95_)[default__.safeIndex((d_1069_v96_).cardinality, len(d_1068_v95_))]
        arr25_ = self.f19
        index176_ = default__.safeIndex(1, (self.f19).length(0))
        arr25_[index176_] = (self).f18
        r0 = (d_1069_v96_).cardinality
        d_1070_v97_: _dafny.Seq
        d_1070_v97_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1071_v98_: _dafny.Set
        d_1071_v98_ = _dafny.Set({(self).f20})
        d_1072_v99_: _dafny.Seq
        d_1072_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dsaa"))
        d_1073_v100_: _dafny.Map
        d_1073_v100_ = _dafny.Map({len(d_1068_v95_): p3})
        d_1074_v101_: _dafny.Map
        d_1074_v101_ = _dafny.Map({d_1068_v95_: d_1073_v100_})
        d_1075_v102_: _dafny.Seq
        d_1075_v102_ = (d_1070_v97_).set(default__.safeIndex(len(d_1071_v98_), len(d_1070_v97_)), (self).fm28(len(d_1072_v99_), p3, d_1074_v101_, globalState))
        r1 = ((d_1070_v97_) + ((d_1070_v97_) + ((d_1075_v102_)))).set(default__.safeIndex(len((d_1072_v99_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))), len((d_1070_v97_) + ((d_1070_v97_) + ((d_1075_v102_))))), (self).f20)
        r2 = d_1068_v95_
        return r0, r1, r2

    def m7(self, p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        r2: bool = False
        r3: _dafny.Seq = _dafny.Seq({})
        d_1076_v0_: _dafny.Seq
        d_1076_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "beccgqvi"))
        (globalState).f14 = (0) - (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1077_i0_ in range(default__.abs(519))])) + (d_1076_v0_)))
        d_1078_v1_: _dafny.Seq
        d_1078_v1_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18, (self).f18, (self).f18, (self).f18])
        d_1079_v2_: _dafny.MultiSet
        d_1079_v2_ = _dafny.MultiSet([d_1078_v1_])
        (globalState).f3 = ((d_1079_v2_).cardinality) >= (p0)
        d_1080_v3_: str
        d_1080_v3_ = _dafny.CodePoint('j')
        d_1080_v3_ = d_1080_v3_
        d_1081_v4_: _dafny.Seq
        d_1081_v4_ = _dafny.SeqWithoutIsStrInference([len(d_1076_v0_), p0, p0])
        d_1082_v5_: _dafny.Set
        d_1082_v5_ = _dafny.Set({(self).f18})
        (globalState).f14 = (len(d_1081_v4_)) * (len((_dafny.Set({(self).f18})) | (d_1082_v5_)))
        (globalState).f14 = ((self).f28) - (len(d_1076_v0_))
        d_1083_i1_: int
        d_1083_i1_ = 0
        with _dafny.label("10"):
            while (self).f18:
                with _dafny.c_label("10"):
                    if (d_1083_i1_) >= (100):
                        raise _dafny.Break("10")
                    d_1083_i1_ = (d_1083_i1_) + (1)
                    (globalState).f9 = (self).f18
                    d_1084_v6_: _dafny.Map
                    d_1084_v6_ = _dafny.Map({(d_1076_v0_) <= (d_1076_v0_): (self).f18})
                    d_1084_v6_ = (d_1084_v6_).set((self).f18, not((self).f18))
                    d_1076_v0_ = (_dafny.SeqWithoutIsStrInference([d_1080_v3_ for d_1085_i2_ in range(default__.abs(276))]) if (self).f18 else d_1076_v0_)
                    d_1086_v7_: C1
                    nw176_ = C1()
                    nw176_.ctor__()
                    d_1086_v7_ = nw176_
                    pass
            pass
        d_1087_v8_: T0
        nw177_ = C4()
        nw177_.ctor__((self).f28, len(d_1082_v5_), self.f17, (self).f18)
        d_1087_v8_ = nw177_
        d_1088_v9_: _dafny.Seq
        d_1088_v9_ = _dafny.SeqWithoutIsStrInference([d_1087_v8_])
        d_1089_v10_: D12
        d_1089_v10_ = D12_DC29(d_1088_v9_)
        d_1090_v11_: _dafny.Seq
        d_1090_v11_ = _dafny.SeqWithoutIsStrInference([((d_1088_v9_).set(default__.safeIndex((((self).f27)[(d_1087_v8_).f18] if ((d_1087_v8_).f18) in ((self).f27) else (self).f28), len(d_1088_v9_)), d_1087_v8_)) + (d_1088_v9_), d_1088_v9_, (d_1089_v10_).cf54])
        r0 = (d_1090_v11_)[default__.safeIndex((self).f28, len(d_1090_v11_))]
        r1 = (((self).f28 if (self).f18 else p0) if (self).f18 else (self).f20)
        d_1091_v12_: D3
        d_1091_v12_ = D3_DC10(d_1076_v0_, (d_1087_v8_).f18, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hym")), _dafny.SeqWithoutIsStrInference([d_1080_v3_ for d_1092_i3_ in range(default__.abs(34))]), (0) - (p0))
        d_1093_v13_: _dafny.MultiSet
        d_1093_v13_ = _dafny.MultiSet([(d_1087_v8_).f18, (self).f18, (self).f18, (d_1091_v12_).cf17])
        r2 = (d_1093_v13_).issubset(d_1093_v13_)
        r3 = d_1081_v4_
        return r0, r1, r2, r3

    @property
    def f28(self):
        return self._f28
    @property
    def f27(self):
        return self._f27

class C7(T2, T1, T0):
    def  __init__(self):
        self._f17: _dafny.Set = _dafny.Set({})
        self._f19: _dafny.Array = _dafny.Array(None, 0)
        self._f21: _dafny.Map = _dafny.Map({})
        self._f20: int = int(0)
        self._f18: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f17(self):
        return self._f17
    @f17.setter
    def f17(self, value):
        self._f17 = value
    @property
    def f19(self):
        return self._f19
    @f19.setter
    def f19(self, value):
        self._f19 = value
    @property
    def f21(self):
        return self._f21
    @f21.setter
    def f21(self, value):
        self._f21 = value
    @property
    def f20(self):
        return self._f20
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f20, f21, f19, f17, f18):
        (self)._f20 = f20
        (self).f21 = f21
        (self).f19 = f19
        (self).f17 = f17
        (self)._f18 = f18

    def fm1(self, p0, globalState):
        return (self).f18

    def fm0(self, p0, p1, globalState):
        return (self).f20

    def fm25(self, p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cd")))) < ((0) - ((self).f20)), (self).f18, (self).f18, (self).f18])

    def fm26(self, globalState):
        return (_dafny.MultiSet([False])).intersection(_dafny.MultiSet([(self).f18, (self).f18]))

    def m0(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_1094_v0_: _dafny.Map
        d_1094_v0_ = _dafny.Map({(self).f20: p1})
        d_1095_v1_: _dafny.Set
        d_1095_v1_ = _dafny.Set({_dafny.Map({797: p1})})
        d_1096_v2_: D1
        d_1096_v2_ = D1_DC4((((d_1094_v0_)[p0] if (p0) in (d_1094_v0_) else True)) and ((self).f18), (_dafny.Set({d_1094_v0_})).intersection(d_1095_v1_), p0)
        source11_ = d_1096_v2_
        if source11_.is_DC4:
            d_1097___mcc_h0_ = source11_.cf9
            d_1098___mcc_h1_ = source11_.cf10
            d_1099___mcc_h2_ = source11_.cf11
            d_1100_cf11_ = d_1099___mcc_h2_
            d_1101_cf10_ = d_1098___mcc_h1_
            d_1102_cf9_ = d_1097___mcc_h0_
            d_1103_v3_: _dafny.Seq
            d_1103_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edctjguo"))
            d_1104_v4_: T2
            nw178_ = C5()
            nw178_.ctor__(d_1103_v3_, p1, (self).f20, self.f21, self.f19, self.f17, d_1102_cf9_)
            d_1104_v4_ = nw178_
            d_1104_v4_ = d_1104_v4_
            d_1100_cf11_ = (p0) - (d_1100_cf11_)
            d_1105_v5_: _dafny.MultiSet
            d_1105_v5_ = _dafny.MultiSet([d_1096_v2_])
            (globalState).f14 = ((d_1105_v5_)[D1_DC4(d_1102_cf9_, d_1095_v1_, 885)] if (D1_DC4(d_1102_cf9_, d_1095_v1_, 885)) in (d_1105_v5_) else (self).f20)
            d_1106_v6_: _dafny.Map
            d_1106_v6_ = _dafny.Map({(d_1100_cf11_) > (p2): self.f19})
            d_1107_v7_: _dafny.Seq
            d_1107_v7_ = _dafny.SeqWithoutIsStrInference([((d_1094_v0_)[(d_1104_v4_).f20] if ((d_1104_v4_).f20) in (d_1094_v0_) else False), (d_1104_v4_).f18, d_1102_cf9_, not(not(not((d_1104_v4_).f18))), (d_1104_v4_).f18])
            d_1106_v6_ = (d_1106_v6_).set((d_1107_v7_)[default__.safeIndex(p2, len(d_1107_v7_))], self.f19)
        elif True:
            d_1108___mcc_h3_ = source11_.cf8
            d_1109_cf8_ = d_1108___mcc_h3_
            d_1110_v8_: _dafny.Seq
            d_1110_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
            d_1111_v9_: _dafny.Map
            d_1111_v9_ = _dafny.Map({p1: (d_1110_v8_) + (d_1110_v8_)})
            d_1112_v10_: _dafny.Seq
            d_1112_v10_ = _dafny.SeqWithoutIsStrInference([d_1111_v9_])
            rhs173_ = not(not(p1))
            rhs174_ = (d_1112_v10_)[default__.safeIndex(p0, len(d_1112_v10_))]
            rhs175_ = (D2_DC6(-752)).cf13
            rhs176_ = self.f19
            rhs177_ = d_1109_cf8_
            lhs151_ = globalState
            lhs152_ = self
            lhs151_.f3 = rhs173_
            d_1111_v9_ = rhs174_
            r2 = rhs175_
            lhs152_.f19 = rhs176_
            d_1109_cf8_ = rhs177_
            d_1113_v11_: _dafny.Array
            def lambda58_(d_1114_p2_):
                def lambda59_(d_1115_i0_):
                    return default__.safeModuloInt(d_1115_i0_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_1114_p2_, len(_dafny.Set({False})), (self).f20})), d_1114_p2_])))

                return lambda59_

            init35_ = lambda58_(p2)
            nw179_ = _dafny.Array(None, 2)
            for i0_35_ in range(nw179_.length(0)):
                nw179_[i0_35_] = init35_(i0_35_)
            d_1113_v11_ = nw179_
            index177_ = default__.safeIndex(704, (d_1113_v11_).length(0))
            (d_1113_v11_)[index177_] = default__.safeModuloInt(p0, 336)
            index178_ = default__.safeIndex(704, (d_1113_v11_).length(0))
            (d_1113_v11_)[index178_] = (self).f20
            d_1116_v12_: str
            d_1116_v12_ = _dafny.CodePoint('v')
            arr26_ = self.f19
            index179_ = default__.safeIndex(379, (self.f19).length(0))
            arr26_[index179_] = p1
            d_1117_v13_: _dafny.Map
            d_1117_v13_ = _dafny.Map({(self).f20: p2})
            d_1118_v14_: _dafny.Seq
            d_1118_v14_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_1110_v8_), p0])])
            arr27_ = self.f19
            index180_ = default__.safeIndex(379, (self.f19).length(0))
            rhs178_ = (d_1116_v12_ if p1 else d_1116_v12_)
            rhs179_ = (((d_1117_v13_)[len(d_1118_v14_)] if (len(d_1118_v14_)) in (d_1117_v13_) else 249)) < (default__.safeModuloInt((d_1113_v11_)[default__.safeIndex(704, (d_1113_v11_).length(0))], (d_1113_v11_)[default__.safeIndex(704, (d_1113_v11_).length(0))]))
            lhs153_ = self.f19
            lhs154_ = default__.safeIndex(379, (self.f19).length(0))
            d_1116_v12_ = rhs178_
            lhs153_[lhs154_] = rhs179_
            d_1119_v15_: D0
            d_1119_v15_ = D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aw")), p0, 155, d_1116_v12_)
            d_1120_v16_: _dafny.Seq
            d_1120_v16_ = _dafny.SeqWithoutIsStrInference([d_1110_v8_, ((d_1119_v15_).cf1) + (d_1110_v8_), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1121_i1_ in range(default__.abs(557))]), d_1110_v8_])
            d_1120_v16_ = (d_1120_v16_) + (_dafny.SeqWithoutIsStrInference([d_1110_v8_ for d_1122_i2_ in range(default__.abs(386))]))
        d_1123_v17_: _dafny.Map
        d_1123_v17_ = _dafny.Map({p2: self.f19})
        d_1124_v18_: _dafny.Map
        d_1124_v18_ = _dafny.Map({p2: len(d_1123_v17_)})
        if (self).fm1(_dafny.Map({d_1124_v18_: p0}), globalState):
            d_1125_v19_: _dafny.Array
            nw180_ = _dafny.Array(int(0), 2)
            d_1125_v19_ = nw180_
            d_1126_v20_: _dafny.Seq
            d_1126_v20_ = _dafny.SeqWithoutIsStrInference([p0, p2, p2])
            d_1127_v21_: str
            d_1127_v21_ = _dafny.CodePoint('s')
            rhs180_ = d_1125_v19_
            rhs181_ = (self).fm0(len((d_1126_v20_) + (d_1126_v20_)), d_1127_v21_, globalState)
            rhs182_ = False
            rhs183_ = not(True)
            rhs184_ = not(not((self).f18))
            lhs155_ = globalState
            lhs156_ = globalState
            lhs157_ = globalState
            lhs158_ = globalState
            lhs159_ = globalState
            lhs155_.f7 = rhs180_
            lhs156_.f14 = rhs181_
            lhs157_.f3 = rhs182_
            lhs158_.f9 = rhs183_
            lhs159_.f9 = rhs184_
            d_1128_v22_: _dafny.Seq
            d_1128_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fm"))
            d_1129_v23_: C5
            nw181_ = C5()
            nw181_.ctor__(d_1128_v22_, not(p1), len(d_1128_v22_), self.f21, self.f19, (self.f17) - (self.f17), (self).f18)
            d_1129_v23_ = nw181_
            d_1130_v24_: _dafny.Map
            d_1130_v24_ = _dafny.Map({False: p2})
            d_1130_v24_ = d_1130_v24_
            if False:
                d_1125_v19_ = d_1125_v19_
                d_1127_v21_ = d_1127_v21_
                r1 = (p2) >= ((self).f20)
                d_1131_v25_: C2
                nw182_ = C2()
                nw182_.ctor__(self.f17, (d_1129_v23_).f30)
                d_1131_v25_ = nw182_
                index181_ = default__.safeIndex(334, (d_1125_v19_).length(0))
                (d_1125_v19_)[index181_] = (self).f20
                index182_ = default__.safeIndex(334, (d_1125_v19_).length(0))
                (d_1125_v19_)[index182_] = (self).f20
            elif True:
                arr28_ = self.f19
                index183_ = default__.safeIndex(523, (self.f19).length(0))
                arr28_[index183_] = not (False) or ((self).f18)
                arr29_ = self.f19
                index184_ = default__.safeIndex(523, (self.f19).length(0))
                arr29_[index184_] = (d_1129_v23_).f30
                d_1132_v26_: _dafny.Array
                nw183_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_1132_v26_ = nw183_
                d_1133_v27_: _dafny.Seq
                d_1133_v27_ = _dafny.SeqWithoutIsStrInference([d_1132_v26_])
                d_1134_v28_: _dafny.Map
                d_1134_v28_ = _dafny.Map({((d_1124_v18_)[p0] if (p0) in (d_1124_v18_) else (self).f20): d_1132_v26_})
                d_1135_v29_: _dafny.Map
                d_1135_v29_ = _dafny.Map({(self).f18: ((d_1134_v28_)[p2] if (p2) in (d_1134_v28_) else d_1132_v26_)})
                d_1136_v30_: _dafny.Array
                nw184_ = _dafny.Array(None, 18)
                nw184_[int(0)] = d_1132_v26_
                nw184_[int(1)] = d_1132_v26_
                nw184_[int(2)] = d_1132_v26_
                nw184_[int(3)] = d_1132_v26_
                nw184_[int(4)] = d_1132_v26_
                nw184_[int(5)] = d_1132_v26_
                nw184_[int(6)] = d_1132_v26_
                nw184_[int(7)] = d_1132_v26_
                nw184_[int(8)] = (d_1133_v27_)[default__.safeIndex(p0, len(d_1133_v27_))]
                nw184_[int(9)] = ((d_1134_v28_)[p0] if (p0) in (d_1134_v28_) else d_1132_v26_)
                nw184_[int(10)] = d_1132_v26_
                nw184_[int(11)] = d_1132_v26_
                nw184_[int(12)] = ((d_1135_v29_)[p1] if (p1) in (d_1135_v29_) else d_1132_v26_)
                nw184_[int(13)] = d_1132_v26_
                nw184_[int(14)] = d_1132_v26_
                nw184_[int(15)] = d_1132_v26_
                nw184_[int(16)] = d_1132_v26_
                nw184_[int(17)] = d_1132_v26_
                d_1136_v30_ = nw184_
                index185_ = default__.safeIndex(510, (d_1136_v30_).length(0))
                (d_1136_v30_)[index185_] = d_1132_v26_
                index186_ = default__.safeIndex(510, (d_1136_v30_).length(0))
                (d_1136_v30_)[index186_] = d_1132_v26_
                d_1125_v19_ = d_1125_v19_
                d_1137_v31_: _dafny.Map
                d_1137_v31_ = _dafny.Map({d_1124_v18_: p0})
                d_1138_v35_: _dafny.Array
                nw185_ = _dafny.Array(None, 14)
                nw185_[int(0)] = (d_1124_v18_) | (d_1124_v18_)
                nw185_[int(1)] = (d_1124_v18_) | (d_1124_v18_)
                nw185_[int(2)] = d_1124_v18_
                def iife91_():
                    coll49_ = _dafny.Map()
                    compr_49_: int
                    for compr_49_ in _dafny.IntegerRange(982, 808):
                        d_1139_v32_: int = compr_49_
                        if ((982) <= (d_1139_v32_)) and ((d_1139_v32_) < (808)):
                            coll49_[(d_1139_v32_) + (p0)] = p2
                    return _dafny.Map(coll49_)
                nw185_[int(3)] = (_dafny.Map({p2: p2}) if (d_1129_v23_).fm1(d_1137_v31_, globalState) else iife91_()
                )
                nw185_[int(4)] = d_1124_v18_
                def iife92_():
                    coll50_ = _dafny.Map()
                    compr_50_: int
                    for compr_50_ in (_dafny.MultiSet([(self).f20, p0])).Elements:
                        d_1140_v33_: int = compr_50_
                        if (d_1140_v33_) in (_dafny.MultiSet([(self).f20, p0])):
                            coll50_[default__.safeModuloInt(d_1140_v33_, p2)] = (self).f20
                    return _dafny.Map(coll50_)
                nw185_[int(5)] = (d_1124_v18_) | (iife92_()
                )
                nw185_[int(6)] = d_1124_v18_
                nw185_[int(7)] = d_1124_v18_
                nw185_[int(8)] = _dafny.Map({(self).f20: (0) - (p0)})
                nw185_[int(9)] = _dafny.Map({p0: p0})
                nw185_[int(10)] = (d_1124_v18_) | (d_1124_v18_)
                nw185_[int(11)] = d_1124_v18_
                def iife93_():
                    coll51_ = _dafny.Map()
                    compr_51_: int
                    for compr_51_ in _dafny.IntegerRange(163, 967):
                        d_1141_v34_: int = compr_51_
                        if ((163) <= (d_1141_v34_)) and ((d_1141_v34_) < (967)):
                            coll51_[(d_1141_v34_) + ((self).f20)] = p2
                    return _dafny.Map(coll51_)
                nw185_[int(12)] = (d_1124_v18_) | (iife93_()
                )
                nw185_[int(13)] = d_1124_v18_
                d_1138_v35_ = nw185_
                index187_ = default__.safeIndex(803, (d_1138_v35_).length(0))
                (d_1138_v35_)[index187_] = d_1124_v18_
                index188_ = default__.safeIndex(803, (d_1138_v35_).length(0))
                (d_1138_v35_)[index188_] = d_1124_v18_
                d_1094_v0_ = (d_1094_v0_).set(289, p1)
            d_1142_v36_: D0
            d_1142_v36_ = D0_DC2((self).f18, p0, d_1125_v19_)
            d_1143_v37_: C4
            nw186_ = C4()
            nw186_.ctor__(p0, p0, self.f17, (p0) < ((0) - ((d_1142_v36_).cf6)))
            d_1143_v37_ = nw186_
        elif True:
            d_1144_v38_: str
            d_1144_v38_ = _dafny.CodePoint('i')
            d_1145_v39_: _dafny.Seq
            d_1145_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            d_1146_v40_: _dafny.Seq
            d_1146_v40_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlsh")), d_1145_v39_, d_1145_v39_, d_1145_v39_, d_1145_v39_])
            d_1147_v41_: _dafny.Map
            d_1147_v41_ = _dafny.Map({d_1144_v38_: len((d_1146_v40_)[default__.safeIndex((self).fm0((self).f20, default__.fm33(len(d_1145_v39_), (self).f18, globalState), globalState), len(d_1146_v40_))])})
            d_1148_v42_: _dafny.Map
            d_1148_v42_ = _dafny.Map({d_1124_v18_: len(d_1147_v41_)})
            if not((self).fm1((_dafny.Map({default__.fm30(p1, globalState): (self).f20})) | (d_1148_v42_), globalState)):
                arr30_ = self.f19
                index189_ = default__.safeIndex(434, (self.f19).length(0))
                arr30_[index189_] = (True if (self).f18 else (self).f18)
                d_1149_v43_: _dafny.Map
                d_1149_v43_ = _dafny.Map({(self).f18: self.f19})
                d_1150_v44_: D5
                d_1150_v44_ = D5_DC12(d_1149_v43_)
                d_1151_v45_: _dafny.MultiSet
                d_1151_v45_ = _dafny.MultiSet([d_1150_v44_])
                d_1152_v46_: _dafny.Map
                d_1152_v46_ = _dafny.Map({p1: d_1151_v45_})
                d_1153_v47_: _dafny.MultiSet
                d_1153_v47_ = _dafny.MultiSet([p2])
                d_1154_v48_: _dafny.Map
                d_1154_v48_ = _dafny.Map({d_1152_v46_: (d_1153_v47_).cardinality})
                arr31_ = self.f19
                index190_ = default__.safeIndex(434, (self.f19).length(0))
                arr31_[index190_] = (((d_1154_v48_)[d_1152_v46_] if (d_1152_v46_) in (d_1154_v48_) else 748)) >= (p0)
                (globalState).f14 = (self).f20
                d_1155_v49_: _dafny.Array
                nw187_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                d_1155_v49_ = nw187_
                index191_ = default__.safeIndex(696, (d_1155_v49_).length(0))
                (d_1155_v49_)[index191_] = d_1144_v38_
                d_1156_v51_: _dafny.Set
                d_1156_v51_ = _dafny.Set({d_1144_v38_, d_1144_v38_, d_1144_v38_, d_1144_v38_, d_1144_v38_})
                d_1157_v52_: _dafny.MultiSet
                def iife94_():
                    coll52_ = _dafny.Set()
                    compr_52_: str
                    for compr_52_ in (self.f21).keys.Elements:
                        d_1158_v50_: str = compr_52_
                        if (d_1158_v50_) in (self.f21):
                            coll52_ = coll52_.union(_dafny.Set([d_1158_v50_]))
                    return _dafny.Set(coll52_)
                d_1157_v52_ = _dafny.MultiSet([iife94_()
                , d_1156_v51_])
                index192_ = default__.safeIndex(696, (d_1155_v49_).length(0))
                (d_1155_v49_)[index192_] = default__.fm33(((d_1157_v52_)[d_1156_v51_] if (d_1156_v51_) in (d_1157_v52_) else p2), (self.f19)[default__.safeIndex(434, (self.f19).length(0))], globalState)
                arr32_ = self.f19
                index193_ = default__.safeIndex(434, (self.f19).length(0))
                arr32_[index193_] = (self.f19)[default__.safeIndex(434, (self.f19).length(0))]
                d_1159_v53_: D11
                d_1159_v53_ = D11_DC26(p1, p2)
                d_1160_v54_: _dafny.Array
                nw188_ = _dafny.Array(None, 19)
                nw188_[int(0)] = p1
                nw188_[int(1)] = True
                nw188_[int(2)] = True
                nw188_[int(3)] = (self).f18
                nw188_[int(4)] = True
                nw188_[int(5)] = (self).f18
                nw188_[int(6)] = (self).f18
                nw188_[int(7)] = (self.f19)[default__.safeIndex(434, (self.f19).length(0))]
                nw188_[int(8)] = (self).f18
                nw188_[int(9)] = True
                nw188_[int(10)] = (self.f19)[default__.safeIndex(434, (self.f19).length(0))]
                nw188_[int(11)] = p1
                nw188_[int(12)] = not(p1)
                nw188_[int(13)] = (self).f18
                nw188_[int(14)] = (d_1159_v53_).cf48
                nw188_[int(15)] = (self).f18
                nw188_[int(16)] = p1
                nw188_[int(17)] = (self.f19)[default__.safeIndex(434, (self.f19).length(0))]
                nw188_[int(18)] = p1
                d_1160_v54_ = nw188_
                d_1161_v55_: C3
                nw189_ = C3()
                nw189_.ctor__(d_1160_v54_, self.f17, (d_1159_v53_).cf48)
                d_1161_v55_ = nw189_
                d_1161_v55_ = (d_1161_v55_ if (p0) < (p0) else d_1161_v55_)
            elif True:
                d_1162_v56_: _dafny.MultiSet
                d_1162_v56_ = _dafny.MultiSet([p2])
                d_1163_v57_: _dafny.Seq
                d_1163_v57_ = _dafny.SeqWithoutIsStrInference([(0) - (p2), p0, (0) - ((p0) + ((self).fm0(p0, d_1144_v38_, globalState))), ((d_1162_v56_) - (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(self).f18])), p0, p0, p0, p0]))).cardinality])
                d_1164_v58_: _dafny.Map
                d_1164_v58_ = _dafny.Map({p1: True})
                d_1165_v59_: _dafny.Array
                nw190_ = _dafny.Array(None, 5)
                nw190_[int(0)] = 708
                nw190_[int(1)] = (len(d_1164_v58_)) + ((self).f20)
                nw190_[int(2)] = (p0) + (p0)
                nw190_[int(3)] = ((self).f20) * (p0)
                nw190_[int(4)] = (0) - (p2)
                d_1165_v59_ = nw190_
                d_1166_v60_: _dafny.Seq
                d_1166_v60_ = _dafny.SeqWithoutIsStrInference([d_1164_v58_, d_1164_v58_])
                index194_ = default__.safeIndex(695, (d_1165_v59_).length(0))
                (d_1165_v59_)[index194_] = (0) - (len(((d_1166_v60_)[default__.safeIndex(p2, len(d_1166_v60_))]).set((self).f18, p1)))
                index195_ = default__.safeIndex(695, (d_1165_v59_).length(0))
                rhs185_ = _dafny.SeqWithoutIsStrInference([len(d_1163_v57_) for d_1167_i3_ in range(default__.abs(22))])
                rhs186_ = p0
                rhs187_ = (p2) < ((self).f20)
                lhs160_ = d_1165_v59_
                lhs161_ = default__.safeIndex(695, (d_1165_v59_).length(0))
                d_1163_v57_ = rhs185_
                lhs160_[lhs161_] = rhs186_
                r1 = rhs187_
                d_1168_v61_: D11
                d_1168_v61_ = D11_DC25(_dafny.Set({self.f19}))
                d_1169_v62_: _dafny.Set
                d_1169_v62_ = _dafny.Set({d_1168_v61_, d_1168_v61_, D11_DC25(self.f17), d_1168_v61_})
                index196_ = default__.safeIndex(695, (d_1165_v59_).length(0))
                rhs188_ = (d_1169_v62_) - (d_1169_v62_)
                rhs189_ = default__.safeModuloInt((self).f20, (d_1165_v59_)[default__.safeIndex(695, (d_1165_v59_).length(0))])
                lhs162_ = d_1165_v59_
                lhs163_ = default__.safeIndex(695, (d_1165_v59_).length(0))
                d_1169_v62_ = rhs188_
                lhs162_[lhs163_] = rhs189_
                arr33_ = self.f19
                index197_ = default__.safeIndex(240, (self.f19).length(0))
                arr33_[index197_] = (self).f18
                d_1170_v63_: C3
                nw191_ = C3()
                nw191_.ctor__(self.f19, self.f17, (self).f18)
                d_1170_v63_ = nw191_
                d_1171_v64_: _dafny.Map
                d_1171_v64_ = _dafny.Map({d_1170_v63_: (self).f20})
                arr34_ = self.f19
                index198_ = default__.safeIndex(240, (self.f19).length(0))
                arr34_[index198_] = (d_1170_v63_) not in ((d_1171_v64_) | (_dafny.Map({d_1170_v63_: (d_1162_v56_).cardinality})))
                d_1172_v65_: _dafny.Map
                d_1172_v65_ = _dafny.Map({d_1165_v59_: _dafny.Map({p1: (self).f18})})
                rhs190_ = (self.f19)[default__.safeIndex(240, (self.f19).length(0))]
                rhs191_ = ((d_1165_v59_)[default__.safeIndex(695, (d_1165_v59_).length(0))]) * ((self).f20)
                rhs192_ = d_1165_v59_
                rhs193_ = ((d_1145_v39_).set(default__.safeIndex((d_1170_v63_).fm0(len(default__.fm40((self).f18, p1, globalState)), d_1144_v38_, globalState), len(d_1145_v39_)), d_1144_v38_)).set(default__.safeIndex(len((((d_1172_v65_)[d_1165_v59_] if (d_1165_v59_) in (d_1172_v65_) else d_1164_v58_)) | (d_1164_v58_)), len((d_1145_v39_).set(default__.safeIndex((d_1170_v63_).fm0(len(default__.fm40((self).f18, p1, globalState)), d_1144_v38_, globalState), len(d_1145_v39_)), d_1144_v38_))), d_1144_v38_)
                lhs164_ = globalState
                lhs165_ = globalState
                lhs164_.f3 = rhs190_
                r2 = rhs191_
                lhs165_.f7 = rhs192_
                d_1145_v39_ = rhs193_
                r2 = default__.safeDivisionInt(p2, len(d_1163_v57_))
            d_1173_v66_: _dafny.MultiSet
            d_1173_v66_ = _dafny.MultiSet([(self).f20])
            (globalState).f15 = d_1173_v66_
            d_1174_v67_: _dafny.Array
            nw192_ = _dafny.Array(D3.default()(), 2)
            d_1174_v67_ = nw192_
            d_1175_v68_: _dafny.Map
            d_1175_v68_ = _dafny.Map({d_1174_v67_: self.f17})
            d_1176_v69_: D11
            d_1176_v69_ = D11_DC25(((d_1175_v68_)[d_1174_v67_] if (d_1174_v67_) in (d_1175_v68_) else self.f17))
            d_1177_v70_: _dafny.Array
            def lambda60_(d_1178_v0_, d_1179_p0_, d_1180_p1_, d_1181_v18_, d_1182_v38_):
                def lambda61_(d_1183_i4_):
                    return D9_DC20(d_1178_v0_, d_1179_p0_, d_1180_p1_, d_1181_v18_, d_1182_v38_)

                return lambda61_

            init36_ = lambda60_(d_1094_v0_, p0, p1, d_1124_v18_, d_1144_v38_)
            nw193_ = _dafny.Array(None, 8)
            for i0_36_ in range(nw193_.length(0)):
                nw193_[i0_36_] = init36_(i0_36_)
            d_1177_v70_ = nw193_
            d_1184_v71_: D9
            d_1184_v71_ = D9_DC20(d_1094_v0_, p0, p1, (d_1124_v18_).set(p0, p0), d_1144_v38_)
            index199_ = default__.safeIndex(523, (d_1177_v70_).length(0))
            (d_1177_v70_)[index199_] = d_1184_v71_
            d_1185_v72_: _dafny.Set
            d_1185_v72_ = _dafny.Set({(self).f18, True, p1})
            index200_ = default__.safeIndex(523, (d_1177_v70_).length(0))
            rhs194_ = ((d_1145_v39_) + (d_1145_v39_)) + (d_1145_v39_)
            rhs195_ = d_1176_v69_
            rhs196_ = ((d_1124_v18_)[551] if (551) in (d_1124_v18_) else p0)
            rhs197_ = (self).fm1(d_1148_v42_, globalState)
            rhs198_ = D9_DC20(d_1094_v0_, p2, (d_1185_v72_).issubset(_dafny.Set({p1})), d_1124_v18_, d_1144_v38_)
            lhs166_ = d_1177_v70_
            lhs167_ = default__.safeIndex(523, (d_1177_v70_).length(0))
            d_1145_v39_ = rhs194_
            d_1176_v69_ = rhs195_
            r2 = rhs196_
            r1 = rhs197_
            lhs166_[lhs167_] = rhs198_
            d_1094_v0_ = (d_1094_v0_).set(p2, (d_1173_v66_).issubset(d_1173_v66_))
            (globalState).f9 = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dca")))) >= (p2)
        r1 = p1
        d_1186_v73_: str
        d_1186_v73_ = _dafny.CodePoint('g')
        d_1187_v74_: _dafny.Seq
        d_1187_v74_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xiu"))
        d_1188_v75_: D3
        d_1188_v75_ = D3_DC10(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1189_i5_ in range(default__.abs(955))]), (d_1186_v73_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwn"))), d_1187_v74_, d_1187_v74_, len(d_1187_v74_))
        source12_ = d_1188_v75_
        if source12_.is_DC9:
            if True:
                d_1190_v76_: _dafny.Seq
                d_1190_v76_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_1191_v77_: C5
                nw194_ = C5()
                nw194_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1192_i6_ in range(default__.abs(948))]), (self).f18, (0) - ((d_1190_v76_)[default__.safeIndex((self).f20, len(d_1190_v76_))]), _dafny.Map({d_1186_v73_: (self).f18}), self.f19, self.f17, p1)
                d_1191_v77_ = nw194_
                arr35_ = self.f19
                index201_ = default__.safeIndex(238, (self.f19).length(0))
                arr35_[index201_] = (p2) <= (p0)
                d_1193_v78_: _dafny.Map
                d_1193_v78_ = _dafny.Map({d_1124_v18_: p0})
                arr36_ = self.f19
                index202_ = default__.safeIndex(238, (self.f19).length(0))
                arr36_[index202_] = (self).fm1(d_1193_v78_, globalState)
                d_1194_v79_: _dafny.Array
                def lambda62_(d_1195_v77_, d_1196_p0_, d_1197_v0_, d_1198_v76_):
                    def lambda63_(d_1199_i7_):
                        return (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_1195_v77_).f30: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_1195_v77_).f30, ((d_1197_v0_)[d_1196_p0_] if (d_1196_p0_) in (d_1197_v0_) else True)]))).cardinality}))])) + (d_1198_v76_)

                    return lambda63_

                init37_ = lambda62_(d_1191_v77_, p0, d_1094_v0_, d_1190_v76_)
                nw195_ = _dafny.Array(None, 13)
                for i0_37_ in range(nw195_.length(0)):
                    nw195_[i0_37_] = init37_(i0_37_)
                d_1194_v79_ = nw195_
                index203_ = default__.safeIndex(366, (d_1194_v79_).length(0))
                (d_1194_v79_)[index203_] = _dafny.SeqWithoutIsStrInference([940 for d_1200_i8_ in range(default__.abs(607))])
                d_1201_v80_: _dafny.Set
                d_1201_v80_ = _dafny.Set({(self).f20})
                d_1202_v81_: _dafny.Array
                nw196_ = _dafny.Array(None, 9)
                d_1202_v81_ = nw196_
                d_1203_v82_: _dafny.Seq
                d_1203_v82_ = _dafny.SeqWithoutIsStrInference([not(False)])
                index204_ = default__.safeIndex(366, (d_1194_v79_).length(0))
                rhs199_ = d_1190_v76_
                rhs200_ = d_1201_v80_
                rhs201_ = d_1202_v81_
                rhs202_ = (p0) - (default__.safeModuloInt(p0, p0))
                rhs203_ = (d_1203_v82_)[default__.safeIndex(p2, len(d_1203_v82_))]
                lhs168_ = d_1194_v79_
                lhs169_ = default__.safeIndex(366, (d_1194_v79_).length(0))
                lhs170_ = globalState
                lhs171_ = globalState
                lhs168_[lhs169_] = rhs199_
                d_1201_v80_ = rhs200_
                d_1202_v81_ = rhs201_
                lhs170_.f14 = rhs202_
                lhs171_.f9 = rhs203_
                d_1204_v83_: _dafny.Map
                d_1204_v83_ = _dafny.Map({True: (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjn"))).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjn")))), d_1186_v73_)).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjn"))).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjn")))), d_1186_v73_))), d_1186_v73_)) + (d_1187_v74_)})
                rhs204_ = ((d_1204_v83_) | (d_1204_v83_)) | (_dafny.Map({False: (d_1191_v77_).f29}))
                rhs205_ = (p0) * (p0)
                rhs206_ = (self).fm25((self).f18, p1, p2, (p1) == ((self).f18), globalState)
                rhs207_ = ((self).f18 if (self.f19)[default__.safeIndex(238, (self.f19).length(0))] else (d_1191_v77_).fm1(d_1193_v78_, globalState))
                lhs172_ = globalState
                lhs173_ = globalState
                d_1204_v83_ = rhs204_
                lhs172_.f14 = rhs205_
                d_1203_v82_ = rhs206_
                lhs173_.f3 = rhs207_
                d_1205_v84_: D3
                d_1205_v84_ = D3_DC9()
                d_1206_v85_: _dafny.Seq
                d_1206_v85_ = _dafny.SeqWithoutIsStrInference([d_1205_v84_])
                d_1207_v86_: D5
                d_1207_v86_ = D5_DC13(not((self).fm1(d_1193_v78_, globalState)), (0) - ((0) - ((self).f20)), d_1206_v85_)
                d_1208_v87_: _dafny.Array
                nw197_ = _dafny.Array(None, 8)
                nw197_[int(0)] = default__.fm41(p2, (self).f18, (d_1191_v77_).f29, d_1186_v73_, globalState)
                nw197_[int(1)] = default__.fm41(len((d_1191_v77_).f29), (self).f18, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")), _dafny.CodePoint('g'), globalState)
                nw197_[int(2)] = d_1207_v86_
                nw197_[int(3)] = d_1207_v86_
                nw197_[int(4)] = default__.fm41((self).f20, p1, _dafny.SeqWithoutIsStrInference([d_1186_v73_ for d_1209_i9_ in range(default__.abs(-486))]), _dafny.CodePoint('x'), globalState)
                nw197_[int(5)] = d_1207_v86_
                nw197_[int(6)] = d_1207_v86_
                nw197_[int(7)] = d_1207_v86_
                d_1208_v87_ = nw197_
                d_1210_v88_: _dafny.Seq
                d_1210_v88_ = d_1206_v85_
                index205_ = default__.safeIndex(268, (d_1208_v87_).length(0))
                (d_1208_v87_)[index205_] = D5_DC13(p1, p2, d_1210_v88_)
                d_1211_v89_: _dafny.Map
                d_1211_v89_ = _dafny.Map({(self).f18: (d_1191_v77_).f30})
                d_1212_v90_: _dafny.Array
                nw198_ = _dafny.Array(None, 12)
                nw198_[int(0)] = p2
                nw198_[int(1)] = (self).fm0(p0, _dafny.CodePoint('g'), globalState)
                nw198_[int(2)] = (0) - (len(d_1211_v89_))
                nw198_[int(3)] = (self).f20
                nw198_[int(4)] = (self).f20
                nw198_[int(5)] = p2
                nw198_[int(6)] = 193
                nw198_[int(7)] = len(((d_1191_v77_).f29).set(default__.safeIndex((self).f20, len((d_1191_v77_).f29)), d_1186_v73_))
                nw198_[int(8)] = p0
                nw198_[int(9)] = (self).f20
                nw198_[int(10)] = p2
                nw198_[int(11)] = 18
                d_1212_v90_ = nw198_
                d_1213_v91_: _dafny.MultiSet
                d_1213_v91_ = _dafny.MultiSet([d_1212_v90_, d_1212_v90_])
                pat_let_tv12_ = d_1213_v91_
                pat_let_tv13_ = d_1212_v90_
                pat_let_tv14_ = d_1212_v90_
                pat_let_tv15_ = d_1213_v91_
                pat_let_tv16_ = d_1210_v88_
                index206_ = default__.safeIndex(268, (d_1208_v87_).length(0))
                def iife95_(_pat_let21_0):
                    def iife96_(d_1214_dt__update__tmp_h0_):
                        def iife97_(_pat_let22_0):
                            def iife98_(d_1215_dt__update_hcf24_h0_):
                                def iife99_(_pat_let23_0):
                                    def iife100_(d_1216_dt__update_hcf25_h0_):
                                        return D5_DC13((d_1214_dt__update__tmp_h0_).cf23, d_1215_dt__update_hcf24_h0_, d_1216_dt__update_hcf25_h0_)
                                    return iife100_(_pat_let23_0)
                                return iife99_(pat_let_tv16_)
                            return iife98_(_pat_let22_0)
                        return iife97_(default__.safeModuloInt((self).f20, ((pat_let_tv12_)[pat_let_tv13_] if (pat_let_tv14_) in (pat_let_tv15_) else 913)))
                    return iife96_(_pat_let21_0)
                (d_1208_v87_)[index206_] = iife95_(d_1207_v86_)
            elif True:
                d_1187_v74_ = _dafny.SeqWithoutIsStrInference([(D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gf")), p2, -503, (D0_DC1(d_1187_v74_, (self).f20, p2, d_1186_v73_)).cf4)).cf4 for d_1217_i10_ in range(default__.abs(654))])
                arr37_ = self.f19
                index207_ = default__.safeIndex(772, (self.f19).length(0))
                arr37_[index207_] = (p0) != ((self).f20)
                arr38_ = self.f19
                index208_ = default__.safeIndex(772, (self.f19).length(0))
                arr38_[index208_] = p1
                d_1218_v92_: _dafny.Set
                d_1218_v92_ = _dafny.Set({p1})
                d_1219_v93_: _dafny.Map
                d_1219_v93_ = _dafny.Map({False: not(p1)})
                d_1220_v94_: _dafny.Map
                d_1220_v94_ = _dafny.Map({p1: len(d_1219_v93_)})
                d_1221_v95_: _dafny.Seq
                d_1221_v95_ = _dafny.SeqWithoutIsStrInference([(self).f20])
                d_1222_v96_: _dafny.Seq
                d_1222_v96_ = _dafny.SeqWithoutIsStrInference([p0, len((d_1220_v94_).set((self).f18, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])])))), len(d_1221_v95_)])
                d_1223_v97_: _dafny.Seq
                d_1223_v97_ = _dafny.SeqWithoutIsStrInference([len(d_1187_v74_), (d_1222_v96_)[default__.safeIndex((self).fm0((self).f20, d_1186_v73_, globalState), len(d_1222_v96_))]])
                d_1224_v98_: C4
                nw199_ = C4()
                nw199_.ctor__((_dafny.MultiSet([241])).cardinality, (0) - ((self).fm0(len(d_1124_v18_), d_1186_v73_, globalState)), self.f17, not((self).f18))
                d_1224_v98_ = nw199_
                d_1225_v99_: _dafny.Set
                d_1225_v99_ = _dafny.Set({len(_dafny.Map({d_1224_v98_: default__.fm37((self).f18, 384, p1, False, globalState)}))})
                d_1226_v100_: _dafny.Array
                def lambda64_(d_1227_i11_):
                    return (d_1227_i11_) + ((self).f20)

                init38_ = lambda64_
                nw200_ = _dafny.Array(None, 6)
                for i0_38_ in range(nw200_.length(0)):
                    nw200_[i0_38_] = init38_(i0_38_)
                d_1226_v100_ = nw200_
                d_1228_v101_: D0
                d_1228_v101_ = D0_DC2(p1, (d_1224_v98_).f25, d_1226_v100_)
                d_1229_v102_: _dafny.Array
                nw201_ = _dafny.Array(None, 20)
                nw201_[int(0)] = p1
                nw201_[int(1)] = p1
                nw201_[int(2)] = (self).f18
                nw201_[int(3)] = False
                nw201_[int(4)] = (self.f19)[default__.safeIndex(772, (self.f19).length(0))]
                nw201_[int(5)] = p1
                nw201_[int(6)] = (self.f19)[default__.safeIndex(772, (self.f19).length(0))]
                nw201_[int(7)] = (self.f19)[default__.safeIndex(772, (self.f19).length(0))]
                nw201_[int(8)] = (self.f19)[default__.safeIndex(772, (self.f19).length(0))]
                nw201_[int(9)] = True
                nw201_[int(10)] = True
                nw201_[int(11)] = (self.f19)[default__.safeIndex(772, (self.f19).length(0))]
                nw201_[int(12)] = (self.f19)[default__.safeIndex(772, (self.f19).length(0))]
                nw201_[int(13)] = (self.f19)[default__.safeIndex(772, (self.f19).length(0))]
                nw201_[int(14)] = (self.f19)[default__.safeIndex(772, (self.f19).length(0))]
                nw201_[int(15)] = not(False)
                nw201_[int(16)] = p1
                nw201_[int(17)] = (self).f18
                nw201_[int(18)] = (self.f19)[default__.safeIndex(772, (self.f19).length(0))]
                nw201_[int(19)] = (d_1228_v101_).cf5
                d_1229_v102_ = nw201_
                d_1230_v103_: _dafny.Seq
                d_1230_v103_ = _dafny.SeqWithoutIsStrInference([d_1229_v102_, d_1229_v102_])
                d_1231_v104_: _dafny.Array
                nw202_ = _dafny.Array(None, 6)
                nw202_[int(0)] = len(d_1225_v99_)
                nw202_[int(1)] = len(d_1222_v96_)
                nw202_[int(2)] = (d_1224_v98_).f25
                nw202_[int(3)] = (self).f20
                nw202_[int(4)] = len(d_1230_v103_)
                nw202_[int(5)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uosnx")))
                d_1231_v104_ = nw202_
                d_1232_v105_: D0
                d_1232_v105_ = D0_DC2((self.f19)[default__.safeIndex(772, (self.f19).length(0))], len(d_1218_v92_), d_1231_v104_)
                arr39_ = self.f19
                index209_ = default__.safeIndex(772, (self.f19).length(0))
                rhs208_ = p2
                rhs209_ = (d_1228_v101_).cf7
                rhs210_ = d_1218_v92_
                rhs211_ = d_1221_v95_
                rhs212_ = not((self.f19)[default__.safeIndex(772, (self.f19).length(0))])
                lhs174_ = globalState
                lhs175_ = self.f19
                lhs176_ = default__.safeIndex(772, (self.f19).length(0))
                r2 = rhs208_
                lhs174_.f7 = rhs209_
                d_1218_v92_ = rhs210_
                d_1223_v97_ = rhs211_
                lhs175_[lhs176_] = rhs212_
                arr40_ = self.f19
                index210_ = default__.safeIndex(772, (self.f19).length(0))
                arr40_[index210_] = (self).f18
                d_1233_v106_: _dafny.Array
                nw203_ = _dafny.Array(_dafny.CodePoint('D'), 12)
                d_1233_v106_ = nw203_
                index211_ = default__.safeIndex(386, (d_1233_v106_).length(0))
                (d_1233_v106_)[index211_] = d_1186_v73_
                index212_ = default__.safeIndex(386, (d_1233_v106_).length(0))
                (d_1233_v106_)[index212_] = d_1186_v73_
            (globalState).f9 = (self).f18
            d_1234_v107_: _dafny.Array
            nw204_ = _dafny.Array(int(0), 10)
            d_1234_v107_ = nw204_
            index213_ = default__.safeIndex(55, (d_1234_v107_).length(0))
            (d_1234_v107_)[index213_] = (p0) - (p0)
            d_1235_v108_: _dafny.MultiSet
            d_1235_v108_ = _dafny.MultiSet([p1])
            index214_ = default__.safeIndex(55, (d_1234_v107_).length(0))
            (d_1234_v107_)[index214_] = (d_1235_v108_).cardinality
            d_1236_v109_: _dafny.Map
            d_1236_v109_ = _dafny.Map({(self).f18: self.f19})
            d_1237_v110_: D5
            d_1237_v110_ = D5_DC12(d_1236_v109_)
            d_1238_v111_: _dafny.Seq
            d_1238_v111_ = _dafny.SeqWithoutIsStrInference([d_1237_v110_, D5_DC12(d_1236_v109_)])
            index215_ = default__.safeIndex(55, (d_1234_v107_).length(0))
            rhs213_ = p1
            rhs214_ = (p0) + (len((d_1238_v111_) + ((d_1238_v111_).set(default__.safeIndex((self).f20, len(d_1238_v111_)), d_1237_v110_))))
            rhs215_ = default__.safeModuloInt(p0, (0) - ((self).f20))
            rhs216_ = d_1186_v73_
            rhs217_ = d_1186_v73_
            lhs177_ = globalState
            lhs178_ = d_1234_v107_
            lhs179_ = default__.safeIndex(55, (d_1234_v107_).length(0))
            lhs177_.f3 = rhs213_
            lhs178_[lhs179_] = rhs214_
            r2 = rhs215_
            d_1186_v73_ = rhs216_
            d_1186_v73_ = rhs217_
        elif source12_.is_DC10:
            d_1239___mcc_h4_ = source12_.cf16
            d_1240___mcc_h5_ = source12_.cf17
            d_1241___mcc_h6_ = source12_.cf18
            d_1242___mcc_h7_ = source12_.cf19
            d_1243___mcc_h8_ = source12_.cf20
            d_1244_cf20_ = d_1243___mcc_h8_
            d_1245_cf19_ = d_1242___mcc_h7_
            d_1246_cf18_ = d_1241___mcc_h6_
            d_1247_cf17_ = d_1240___mcc_h5_
            d_1248_cf16_ = d_1239___mcc_h4_
            (globalState).f14 = default__.safeModuloInt(617, p2)
            d_1249_v112_: C0
            nw205_ = C0()
            nw205_.ctor__()
            d_1249_v112_ = nw205_
            d_1250_v113_: _dafny.Seq
            d_1250_v113_ = _dafny.SeqWithoutIsStrInference([default__.fm30(p1, globalState), d_1124_v18_, default__.fm30((self).f18, globalState)])
            d_1251_v114_: _dafny.Seq
            d_1251_v114_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(len((d_1250_v113_)[default__.safeIndex(p2, len(d_1250_v113_))]), 944), p0, (self).fm0((self).f20, d_1186_v73_, globalState)])
            d_1244_cf20_ = len(d_1251_v114_)
            d_1252_v115_: _dafny.Set
            d_1252_v115_ = _dafny.Set({d_1244_cf20_})
            d_1253_v116_: _dafny.Map
            d_1253_v116_ = _dafny.Map({d_1124_v18_: (0) - ((self).f20)})
            d_1254_v117_: _dafny.Seq
            d_1254_v117_ = _dafny.SeqWithoutIsStrInference([(self).fm1(d_1253_v116_, globalState)])
            d_1252_v115_ = (default__.fm42(_dafny.MultiSet([True]), p1, d_1254_v117_, globalState)).intersection(d_1252_v115_)
        elif True:
            d_1255___mcc_h9_ = source12_.cf15
            d_1256_cf15_ = d_1255___mcc_h9_
            (globalState).f14 = (self).f20
            d_1257_v118_: _dafny.Array
            def lambda65_(d_1258_i12_):
                return (_dafny.SeqWithoutIsStrInference([(self).f18])) + (_dafny.SeqWithoutIsStrInference([(self).f18]))

            init39_ = lambda65_
            nw206_ = _dafny.Array(None, 27)
            for i0_39_ in range(nw206_.length(0)):
                nw206_[i0_39_] = init39_(i0_39_)
            d_1257_v118_ = nw206_
            d_1259_v119_: _dafny.Map
            d_1259_v119_ = _dafny.Map({p1: p1})
            d_1260_v120_: D9
            d_1260_v120_ = D9_DC20(_dafny.Map({len(d_1259_v119_): False}), -81, (self).f18, d_1124_v18_, d_1186_v73_)
            d_1261_v121_: _dafny.Seq
            d_1261_v121_ = _dafny.SeqWithoutIsStrInference([not(p1)])
            d_1262_v122_: _dafny.Set
            d_1262_v122_ = _dafny.Set({(self).f20, p2, p0})
            d_1263_v123_: _dafny.Map
            d_1263_v123_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference([p1])})
            nw207_ = _dafny.Array(None, 17)
            nw207_[int(0)] = _dafny.SeqWithoutIsStrInference([p1, p1, p1, (d_1260_v120_).cf38])
            nw207_[int(1)] = d_1261_v121_
            nw207_[int(2)] = (_dafny.SeqWithoutIsStrInference([(self).f18])) + (d_1261_v121_)
            nw207_[int(3)] = (d_1261_v121_) + (d_1261_v121_)
            nw207_[int(4)] = (d_1261_v121_) + (d_1261_v121_)
            nw207_[int(5)] = d_1261_v121_
            nw207_[int(6)] = d_1261_v121_
            nw207_[int(7)] = (_dafny.SeqWithoutIsStrInference([p1])) + (d_1261_v121_)
            nw207_[int(8)] = (self).fm25(not((self).f18), p1, len(d_1262_v122_), (self).f18, globalState)
            nw207_[int(9)] = (d_1261_v121_) + (_dafny.SeqWithoutIsStrInference([p1, p1]))
            nw207_[int(10)] = _dafny.SeqWithoutIsStrInference([(self).f18])
            nw207_[int(11)] = d_1261_v121_
            nw207_[int(12)] = d_1261_v121_
            nw207_[int(13)] = d_1261_v121_
            nw207_[int(14)] = d_1261_v121_
            nw207_[int(15)] = d_1261_v121_
            nw207_[int(16)] = ((d_1263_v123_)[False] if (False) in (d_1263_v123_) else d_1261_v121_)
            d_1257_v118_ = nw207_
            d_1264_v124_: _dafny.Array
            nw208_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_1264_v124_ = nw208_
            index216_ = default__.safeIndex(837, (d_1264_v124_).length(0))
            (d_1264_v124_)[index216_] = self.f19
            index217_ = default__.safeIndex(837, (d_1264_v124_).length(0))
            (d_1264_v124_)[index217_] = self.f19
            d_1265_v125_: _dafny.Array
            def lambda66_(d_1266_i13_):
                return (d_1266_i13_) + (647)

            init40_ = lambda66_
            nw209_ = _dafny.Array(None, 26)
            for i0_40_ in range(nw209_.length(0)):
                nw209_[i0_40_] = init40_(i0_40_)
            d_1265_v125_ = nw209_
            (globalState).f7 = d_1265_v125_
        d_1267_v126_: _dafny.Set
        d_1267_v126_ = _dafny.Set({(p0) >= (len((self).fm25((self).f18, p1, p0, (self).f18, globalState)))})
        d_1268_v127_: _dafny.Map
        d_1268_v127_ = _dafny.Map({d_1124_v18_: len(_dafny.Map({(self).f18: p0}))})
        d_1269_v128_: _dafny.Seq
        d_1269_v128_ = _dafny.SeqWithoutIsStrInference([d_1267_v126_, _dafny.Set({p1}), d_1267_v126_, default__.fm37((self).f18, (self).f20, p1, (self).fm1(d_1268_v127_, globalState), globalState)])
        d_1267_v126_ = (d_1269_v128_)[default__.safeIndex(p0, len(d_1269_v128_))]
        d_1270_v129_: _dafny.Seq
        d_1270_v129_ = _dafny.SeqWithoutIsStrInference([(self).f18])
        d_1271_v130_: C1
        nw210_ = C1()
        nw210_.ctor__()
        d_1271_v130_ = nw210_
        d_1272_v131_: _dafny.Set
        d_1272_v131_ = _dafny.Set({d_1271_v130_, d_1271_v130_})
        d_1273_v132_: _dafny.Map
        d_1273_v132_ = _dafny.Map({(self).f18: d_1272_v131_})
        d_1274_v133_: _dafny.Map
        d_1274_v133_ = _dafny.Map({(d_1187_v74_).set(default__.safeIndex(len(d_1270_v129_), len(d_1187_v74_)), d_1186_v73_): ((d_1273_v132_)[(self).f18] if ((self).f18) in (d_1273_v132_) else d_1272_v131_)})
        d_1275_i14_: int
        d_1275_i14_ = 0
        with _dafny.label("11"):
            while ((d_1187_v74_) + (d_1187_v74_)) not in (d_1274_v133_):
                with _dafny.c_label("11"):
                    if (d_1275_i14_) >= (100):
                        raise _dafny.Break("11")
                    d_1275_i14_ = (d_1275_i14_) + (1)
                    (globalState).f14 = p2
                    d_1276_v134_: _dafny.Array
                    nw211_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                    d_1276_v134_ = nw211_
                    index218_ = default__.safeIndex(93, (d_1276_v134_).length(0))
                    (d_1276_v134_)[index218_] = d_1186_v73_
                    index219_ = default__.safeIndex(93, (d_1276_v134_).length(0))
                    (d_1276_v134_)[index219_] = d_1186_v73_
                    r2 = p2
                    d_1277_v135_: D3
                    d_1277_v135_ = D3_DC9()
                    d_1278_v136_: _dafny.Seq
                    d_1278_v136_ = _dafny.SeqWithoutIsStrInference([D3_DC9(), d_1277_v135_])
                    d_1279_v137_: _dafny.Seq
                    d_1279_v137_ = d_1278_v136_
                    source13_ = D5_DC13(p1, ((self).f20) * (p2), d_1279_v137_)
                    if source13_.is_DC13:
                        d_1280___mcc_h10_ = source13_.cf23
                        d_1281___mcc_h11_ = source13_.cf24
                        d_1282___mcc_h12_ = source13_.cf25
                        d_1283_cf25_ = d_1282___mcc_h12_
                        d_1284_cf24_ = d_1281___mcc_h11_
                        d_1285_cf23_ = d_1280___mcc_h10_
                        (globalState).f14 = -602
                        d_1286_v138_: _dafny.MultiSet
                        d_1286_v138_ = _dafny.MultiSet([(self).f20, (self).f20, p0, p2])
                        (globalState).f14 = ((d_1286_v138_)[(self).f20] if ((self).f20) in (d_1286_v138_) else p0)
                        d_1287_v139_: C0
                        nw212_ = C0()
                        nw212_.ctor__()
                        d_1287_v139_ = nw212_
                        d_1288_v140_: _dafny.Seq
                        d_1288_v140_ = _dafny.SeqWithoutIsStrInference([len(d_1094_v0_)])
                        d_1289_v141_: _dafny.Map
                        d_1289_v141_ = _dafny.Map({(d_1288_v140_)[default__.safeIndex(p0, len(d_1288_v140_))]: d_1287_v139_})
                        d_1290_v142_: T0
                        nw213_ = C2()
                        nw213_.ctor__(self.f17, d_1285_cf23_)
                        d_1290_v142_ = nw213_
                        d_1291_v143_: _dafny.Seq
                        d_1291_v143_ = _dafny.SeqWithoutIsStrInference([d_1290_v142_])
                        d_1292_v144_: _dafny.Seq
                        d_1292_v144_ = _dafny.SeqWithoutIsStrInference([d_1290_v142_, d_1290_v142_, d_1290_v142_, (d_1291_v143_)[default__.safeIndex(p2, len(d_1291_v143_))]])
                        d_1293_v145_: _dafny.Seq
                        d_1293_v145_ = _dafny.SeqWithoutIsStrInference([d_1287_v139_, d_1287_v139_])
                        d_1294_v146_: _dafny.Array
                        nw214_ = _dafny.Array(None, 29)
                        nw214_[int(0)] = d_1287_v139_
                        nw214_[int(1)] = d_1287_v139_
                        nw214_[int(2)] = d_1287_v139_
                        nw214_[int(3)] = d_1287_v139_
                        nw214_[int(4)] = d_1287_v139_
                        nw214_[int(5)] = d_1287_v139_
                        nw214_[int(6)] = d_1287_v139_
                        nw214_[int(7)] = d_1287_v139_
                        nw214_[int(8)] = d_1287_v139_
                        nw214_[int(9)] = d_1287_v139_
                        nw214_[int(10)] = d_1287_v139_
                        nw214_[int(11)] = d_1287_v139_
                        nw214_[int(12)] = d_1287_v139_
                        nw214_[int(13)] = d_1287_v139_
                        nw214_[int(14)] = d_1287_v139_
                        nw214_[int(15)] = ((d_1289_v141_)[(_dafny.MultiSet(d_1291_v143_)).cardinality] if ((_dafny.MultiSet(d_1291_v143_)).cardinality) in (d_1289_v141_) else d_1287_v139_)
                        nw214_[int(16)] = d_1287_v139_
                        nw214_[int(17)] = d_1287_v139_
                        nw214_[int(18)] = d_1287_v139_
                        nw214_[int(19)] = d_1287_v139_
                        nw214_[int(20)] = d_1287_v139_
                        nw214_[int(21)] = d_1287_v139_
                        nw214_[int(22)] = d_1287_v139_
                        nw214_[int(23)] = d_1287_v139_
                        nw214_[int(24)] = d_1287_v139_
                        nw214_[int(25)] = d_1287_v139_
                        nw214_[int(26)] = (d_1287_v139_ if (self).f18 else d_1287_v139_)
                        nw214_[int(27)] = d_1287_v139_
                        nw214_[int(28)] = (d_1293_v145_)[default__.safeIndex(-362, len(d_1293_v145_))]
                        d_1294_v146_ = nw214_
                        d_1294_v146_ = d_1294_v146_
                        d_1295_v147_: _dafny.Map
                        d_1295_v147_ = _dafny.Map({p0: d_1187_v74_})
                        d_1296_v148_: _dafny.Array
                        nw215_ = _dafny.Array(None, 15)
                        nw215_[int(0)] = p2
                        nw215_[int(1)] = (self).f20
                        nw215_[int(2)] = p2
                        nw215_[int(3)] = (self).f20
                        nw215_[int(4)] = (self).f20
                        nw215_[int(5)] = p2
                        nw215_[int(6)] = len(d_1295_v147_)
                        nw215_[int(7)] = len((d_1267_v126_).intersection(d_1267_v126_))
                        nw215_[int(8)] = d_1284_cf24_
                        nw215_[int(9)] = (self).f20
                        nw215_[int(10)] = d_1284_cf24_
                        nw215_[int(11)] = (0) - ((0) - (d_1284_cf24_))
                        nw215_[int(12)] = (d_1288_v140_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1187_v74_ for d_1297_i15_ in range(default__.abs(-78))])), len(d_1288_v140_))]
                        nw215_[int(13)] = (self).f20
                        nw215_[int(14)] = (p2 if (d_1290_v142_).f18 else (0) - (p0))
                        d_1296_v148_ = nw215_
                        index220_ = default__.safeIndex(707, (d_1296_v148_).length(0))
                        (d_1296_v148_)[index220_] = len(d_1187_v74_)
                        d_1298_v149_: _dafny.MultiSet
                        d_1298_v149_ = _dafny.MultiSet([d_1187_v74_, d_1187_v74_, d_1187_v74_])
                        index221_ = default__.safeIndex(707, (d_1296_v148_).length(0))
                        (d_1296_v148_)[index221_] = ((self).f20) * ((187) - ((d_1298_v149_).cardinality))
                    elif source13_.is_DC14:
                        d_1299___mcc_h13_ = source13_.cf26
                        d_1300___mcc_h14_ = source13_.cf27
                        d_1301___mcc_h15_ = source13_.cf28
                        d_1302_cf28_ = d_1301___mcc_h15_
                        d_1303_cf27_ = d_1300___mcc_h14_
                        d_1304_cf26_ = d_1299___mcc_h13_
                        (self).f19 = self.f19
                        arr41_ = self.f19
                        index222_ = default__.safeIndex(194, (self.f19).length(0))
                        arr41_[index222_] = (self).f18
                        arr42_ = self.f19
                        index223_ = default__.safeIndex(194, (self.f19).length(0))
                        arr42_[index223_] = True
                        d_1305_v150_: _dafny.Seq
                        d_1305_v150_ = _dafny.SeqWithoutIsStrInference([p2, (self).f20])
                        d_1306_v151_: _dafny.Map
                        d_1306_v151_ = _dafny.Map({d_1187_v74_: d_1305_v150_})
                        d_1306_v151_ = (d_1306_v151_).set((d_1187_v74_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kla"))), (d_1305_v150_) + (_dafny.SeqWithoutIsStrInference([(self).f20, p2])))
                        d_1307_v152_: _dafny.MultiSet
                        d_1307_v152_ = _dafny.MultiSet([d_1304_cf26_])
                        (globalState).f14 = (self).fm0((d_1307_v152_).cardinality, (d_1276_v134_)[default__.safeIndex(93, (d_1276_v134_).length(0))], globalState)
                    elif True:
                        d_1308___mcc_h16_ = source13_.cf22
                        d_1309_cf22_ = d_1308___mcc_h16_
                        (globalState).f14 = (0) - (p2)
                        d_1310_v153_: _dafny.Array
                        nw216_ = _dafny.Array(_dafny.Map({}), 16)
                        d_1310_v153_ = nw216_
                        d_1311_v154_: D7
                        d_1311_v154_ = D7_DC17(d_1310_v153_, (self).f18, p2)
                        d_1312_v155_: _dafny.Map
                        d_1312_v155_ = _dafny.Map({d_1277_v135_: d_1311_v154_})
                        d_1313_v156_: _dafny.Array
                        def lambda67_(d_1314_v73_, d_1315_v134_):
                            def lambda68_(d_1316_i16_):
                                return (d_1316_i16_) + (len(_dafny.Set({d_1314_v73_, (d_1315_v134_)[default__.safeIndex(93, (d_1315_v134_).length(0))]})))

                            return lambda68_

                        init41_ = lambda67_(d_1186_v73_, d_1276_v134_)
                        nw217_ = _dafny.Array(None, 23)
                        for i0_41_ in range(nw217_.length(0)):
                            nw217_[i0_41_] = init41_(i0_41_)
                        d_1313_v156_ = nw217_
                        index224_ = default__.safeIndex(642, (d_1313_v156_).length(0))
                        (d_1313_v156_)[index224_] = p0
                        index225_ = default__.safeIndex(289, (d_1313_v156_).length(0))
                        (d_1313_v156_)[index225_] = (self).f20
                        d_1317_v157_: _dafny.Seq
                        d_1317_v157_ = _dafny.SeqWithoutIsStrInference([d_1312_v155_])
                        index226_ = default__.safeIndex(642, (d_1313_v156_).length(0))
                        index227_ = default__.safeIndex(289, (d_1313_v156_).length(0))
                        rhs218_ = ((d_1317_v157_)[default__.safeIndex(len(d_1270_v129_), len(d_1317_v157_))]) | (d_1312_v155_)
                        rhs219_ = (0) - (p2)
                        rhs220_ = default__.safeModuloInt(p0, p0)
                        rhs221_ = p0
                        lhs180_ = globalState
                        lhs181_ = d_1313_v156_
                        lhs182_ = default__.safeIndex(642, (d_1313_v156_).length(0))
                        lhs183_ = d_1313_v156_
                        lhs184_ = default__.safeIndex(289, (d_1313_v156_).length(0))
                        d_1312_v155_ = rhs218_
                        lhs180_.f14 = rhs219_
                        lhs181_[lhs182_] = rhs220_
                        lhs183_[lhs184_] = rhs221_
                        d_1318_v158_: _dafny.Array
                        nw218_ = _dafny.Array(_dafny.Array(None, 0), 14)
                        d_1318_v158_ = nw218_
                        index228_ = default__.safeIndex(646, (d_1318_v158_).length(0))
                        (d_1318_v158_)[index228_] = self.f19
                        index229_ = default__.safeIndex(646, (d_1318_v158_).length(0))
                        (d_1318_v158_)[index229_] = self.f19
                        d_1319_v160_: _dafny.Map
                        def iife101_():
                            coll53_ = _dafny.Map()
                            compr_53_: int
                            for compr_53_ in (d_1094_v0_).keys.Elements:
                                d_1320_v159_: int = compr_53_
                                if (d_1320_v159_) in (d_1094_v0_):
                                    coll53_[(d_1320_v159_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1321_i17_ in range(default__.abs(116))])))] = p1
                            return _dafny.Map(coll53_)
                        d_1319_v160_ = _dafny.Map({p1: (iife101_()
                        ) | (d_1094_v0_)})
                        d_1319_v160_ = (d_1319_v160_).set(((self).f18) or ((self).f18), (d_1094_v0_) | (d_1094_v0_))
                    pass
            pass
        r0 = p1
        r1 = p1
        r2 = (0) - (((595) * (p0)) - (((self).f20) + (p2)))
        return r0, r1, r2


class C8(T1, T2, T0):
    def  __init__(self):
        self._f17: _dafny.Set = _dafny.Set({})
        self._f19: _dafny.Array = _dafny.Array(None, 0)
        self._f21: _dafny.Map = _dafny.Map({})
        self._f20: int = int(0)
        self._f18: bool = False
        self.f23: int = int(0)
        self._f24: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f17(self):
        return self._f17
    @f17.setter
    def f17(self, value):
        self._f17 = value
    @property
    def f19(self):
        return self._f19
    @f19.setter
    def f19(self, value):
        self._f19 = value
    @property
    def f21(self):
        return self._f21
    @f21.setter
    def f21(self, value):
        self._f21 = value
    @property
    def f20(self):
        return self._f20
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f23, f24, f19, f17, f18, f20, f21):
        (self).f23 = f23
        (self)._f24 = f24
        (self).f19 = f19
        (self).f17 = f17
        (self)._f18 = f18
        (self)._f20 = f20
        (self).f21 = f21

    def fm0(self, p0, p1, globalState):
        return ((len(_dafny.Map({(self).f18: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fmocyx")))}))) - (self.f23)) * (482)

    def fm1(self, p0, globalState):
        return (self).f18

    def m0(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_1322_i0_: int
        d_1322_i0_ = 0
        with _dafny.label("12"):
            while (self).f24:
                with _dafny.c_label("12"):
                    if (d_1322_i0_) >= (100):
                        raise _dafny.Break("12")
                    d_1322_i0_ = (d_1322_i0_) + (1)
                    d_1323_v0_: _dafny.Seq
                    d_1323_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgnisgrc"))
                    d_1324_v1_: _dafny.Map
                    d_1324_v1_ = _dafny.Map({(_dafny.Map({p0: p0})).set((self).f20, (self).f20): self.f23})
                    d_1323_v0_ = (d_1323_v0_ if (self).fm1(d_1324_v1_, globalState) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkcjnysf")))
                    (globalState).f3 = not((self).f18)
                    d_1325_v2_: _dafny.Array
                    nw219_ = _dafny.Array(int(0), 21)
                    d_1325_v2_ = nw219_
                    (globalState).f7 = d_1325_v2_
                    d_1326_v3_: _dafny.Array
                    nw220_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
                    d_1326_v3_ = nw220_
                    d_1326_v3_ = (d_1326_v3_ if (self).f18 else d_1326_v3_)
                    pass
            pass
        d_1327_v4_: _dafny.Map
        d_1327_v4_ = _dafny.Map({(self).f24: self.f19})
        d_1328_v5_: T0
        nw221_ = C2()
        nw221_.ctor__((self.f17) | (_dafny.Set({self.f19, ((d_1327_v4_)[p1] if (p1) in (d_1327_v4_) else self.f19), self.f19})), not(p1))
        d_1328_v5_ = nw221_
        d_1328_v5_ = d_1328_v5_
        (globalState).f9 = default__.fm13(globalState)
        d_1329_v6_: _dafny.Set
        d_1329_v6_ = _dafny.Set({(d_1328_v5_).f18, (d_1328_v5_).f18})
        d_1330_v7_: C4
        nw222_ = C4()
        nw222_.ctor__(len(d_1329_v6_), (self).f20, d_1328_v5_.f17, (self).f24)
        d_1330_v7_ = nw222_
        r1 = (self).f24
        rhs222_ = (d_1328_v5_).f18
        rhs223_ = (d_1328_v5_).f18
        lhs185_ = globalState
        lhs186_ = globalState
        lhs185_.f9 = rhs222_
        lhs186_.f9 = rhs223_
        r0 = p1
        r1 = ((d_1329_v6_) | (_dafny.Set({p1}))).isdisjoint((d_1329_v6_).intersection(d_1329_v6_))
        d_1331_v8_: _dafny.Array
        nw223_ = _dafny.Array(_dafny.CodePoint('D'), 19)
        d_1331_v8_ = nw223_
        d_1332_v9_: _dafny.MultiSet
        d_1332_v9_ = _dafny.MultiSet([d_1331_v8_])
        r2 = (((d_1332_v9_).intersection(d_1332_v9_)).cardinality) - ((self).f20)
        return r0, r1, r2

    def m1(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        (globalState).f14 = default__.safeDivisionInt(p2, (0) - (-181))
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (self.f19).length(0)):
            d_1333_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_1333_i0_)) and ((d_1333_i0_) < ((self.f19).length(0)))):
                arr43_ = self.f19
                arr43_[(d_1333_i0_)] = (self).f24
        d_1334_i1_: int
        d_1334_i1_ = 0
        with _dafny.label("13"):
            while not((self).f24):
                with _dafny.c_label("13"):
                    if (d_1334_i1_) >= (100):
                        raise _dafny.Break("13")
                    d_1334_i1_ = (d_1334_i1_) + (1)
                    d_1335_v0_: _dafny.Seq
                    d_1335_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
                    d_1335_v0_ = d_1335_v0_
                    d_1336_v1_: T2
                    nw224_ = C7()
                    nw224_.ctor__((self).f20, self.f21, self.f19, self.f17, not((self).f18))
                    d_1336_v1_ = nw224_
                    d_1337_v2_: _dafny.Map
                    d_1337_v2_ = _dafny.Map({p2: d_1336_v1_})
                    d_1338_v3_: _dafny.Seq
                    d_1338_v3_ = _dafny.SeqWithoutIsStrInference([d_1337_v2_, d_1337_v2_, d_1337_v2_, d_1337_v2_, d_1337_v2_])
                    d_1339_v4_: _dafny.Map
                    d_1339_v4_ = _dafny.Map({len(d_1338_v3_): d_1336_v1_.f19})
                    d_1340_v5_: D13
                    d_1340_v5_ = D13_DC31(d_1339_v4_)
                    d_1341_v6_: _dafny.Array
                    nw225_ = _dafny.Array(None, 10)
                    nw225_[int(0)] = d_1339_v4_
                    nw225_[int(1)] = _dafny.Map({p0: d_1336_v1_.f19})
                    nw225_[int(2)] = d_1339_v4_
                    nw225_[int(3)] = d_1339_v4_
                    nw225_[int(4)] = (d_1339_v4_) | (d_1339_v4_)
                    nw225_[int(5)] = (d_1340_v5_).cf60
                    nw225_[int(6)] = (d_1339_v4_) | (d_1339_v4_)
                    nw225_[int(7)] = d_1339_v4_
                    nw225_[int(8)] = _dafny.Map({(d_1336_v1_).f20: d_1336_v1_.f19})
                    nw225_[int(9)] = (d_1339_v4_) | (d_1339_v4_)
                    d_1341_v6_ = nw225_
                    index230_ = default__.safeIndex(440, (d_1341_v6_).length(0))
                    (d_1341_v6_)[index230_] = d_1339_v4_
                    index231_ = default__.safeIndex(440, (d_1341_v6_).length(0))
                    (d_1341_v6_)[index231_] = d_1339_v4_
                    d_1342_v7_: _dafny.Set
                    d_1342_v7_ = _dafny.Set({(self).f18})
                    d_1343_v8_: _dafny.Seq
                    d_1343_v8_ = _dafny.SeqWithoutIsStrInference([d_1342_v7_, d_1342_v7_, d_1342_v7_, d_1342_v7_])
                    d_1344_v9_: _dafny.Map
                    d_1344_v9_ = _dafny.Map({(d_1336_v1_).f18: p0})
                    d_1345_v10_: _dafny.Map
                    d_1345_v10_ = _dafny.Map({(d_1336_v1_).f20: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxjbxgjo")))})
                    d_1346_v11_: _dafny.Seq
                    d_1346_v11_ = _dafny.SeqWithoutIsStrInference([(self).f18, p1])
                    d_1347_v12_: _dafny.MultiSet
                    d_1347_v12_ = _dafny.MultiSet([d_1344_v9_])
                    d_1348_v13_: _dafny.Array
                    nw226_ = _dafny.Array(None, 15)
                    nw226_[int(0)] = p0
                    nw226_[int(1)] = len(d_1345_v10_)
                    nw226_[int(2)] = len(d_1346_v11_)
                    nw226_[int(3)] = -285
                    nw226_[int(4)] = p0
                    nw226_[int(5)] = (d_1347_v12_).cardinality
                    nw226_[int(6)] = (0) - ((self).f20)
                    nw226_[int(7)] = len(_dafny.SeqWithoutIsStrInference([p1, (self).f24, (self).f18]))
                    nw226_[int(8)] = (self).f20
                    nw226_[int(9)] = (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvsdc")))))
                    nw226_[int(10)] = p0
                    nw226_[int(11)] = p0
                    nw226_[int(12)] = p2
                    nw226_[int(13)] = self.f23
                    nw226_[int(14)] = (self).f20
                    d_1348_v13_ = nw226_
                    d_1349_v14_: D0
                    d_1349_v14_ = D0_DC2(p1, p0, d_1348_v13_)
                    d_1350_v16_: _dafny.MultiSet
                    d_1350_v16_ = _dafny.MultiSet([(d_1336_v1_).f20])
                    d_1351_v17_: _dafny.Array
                    nw227_ = _dafny.Array(None, 8)
                    nw227_[int(0)] = default__.fm7(916, d_1344_v9_, D0_DC0((d_1349_v14_).cf5), globalState)
                    nw227_[int(1)] = (self).f20
                    nw227_[int(2)] = (len(_dafny.SeqWithoutIsStrInference([p1, p3]))) - ((0) - (len(d_1344_v9_)))
                    nw227_[int(3)] = p0
                    nw227_[int(4)] = p2
                    def iife102_():
                        coll54_ = _dafny.Set()
                        compr_54_: int
                        for compr_54_ in _dafny.IntegerRange(60, 957):
                            d_1352_v15_: int = compr_54_
                            if ((60) <= (d_1352_v15_)) and ((d_1352_v15_) < (957)):
                                coll54_ = coll54_.union(_dafny.Set([(d_1352_v15_) * (len(_dafny.Set({p2, (self).f20})))]))
                        return _dafny.Set(coll54_)
                    nw227_[int(5)] = len((iife102_()
                    ) - (_dafny.Set({((d_1350_v16_).set((self).f20, default__.abs(self.f23))).cardinality})))
                    nw227_[int(6)] = (self.f23) + ((self).f20)
                    nw227_[int(7)] = default__.safeDivisionInt((0) - (p0), p2)
                    d_1351_v17_ = nw227_
                    rhs224_ = ((d_1342_v7_ if (self).f24 else (d_1343_v8_)[default__.safeIndex(self.f23, len(d_1343_v8_))])) | (d_1342_v7_)
                    rhs225_ = (self).f18
                    rhs226_ = d_1351_v17_
                    rhs227_ = (self).f18
                    lhs187_ = globalState
                    lhs188_ = globalState
                    lhs189_ = globalState
                    d_1342_v7_ = rhs224_
                    lhs187_.f9 = rhs225_
                    lhs188_.f7 = rhs226_
                    lhs189_.f9 = rhs227_
                    d_1353_v18_: C0
                    nw228_ = C0()
                    nw228_.ctor__()
                    d_1353_v18_ = nw228_
                    pass
            pass
        arr44_ = self.f19
        index232_ = default__.safeIndex(17, (self.f19).length(0))
        arr44_[index232_] = (default__.fm29(globalState)).issubset(_dafny.MultiSet([p3]))
        d_1354_v19_: D5
        d_1354_v19_ = D5_DC14(not(False), self.f23, p2)
        d_1355_v20_: _dafny.Seq
        d_1355_v20_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_1356_v21_: _dafny.Map
        d_1356_v21_ = _dafny.Map({(self).f18: ((d_1354_v19_).cf27) + (len(d_1355_v20_))})
        d_1357_v22_: _dafny.Seq
        d_1357_v22_ = _dafny.SeqWithoutIsStrInference([False])
        d_1358_v23_: _dafny.Map
        d_1358_v23_ = _dafny.Map({p0: (self).f24})
        arr45_ = self.f19
        index233_ = default__.safeIndex(17, (self.f19).length(0))
        rhs228_ = ((self).f20) < ((len(d_1357_v22_)) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dn")))))
        rhs229_ = ((p2) >= (self.f23) if p3 else (self).f24)
        rhs230_ = (_dafny.Map({((d_1358_v23_)[self.f23] if (self.f23) in (d_1358_v23_) else p3): self.f23})) | (d_1356_v21_)
        lhs190_ = self.f19
        lhs191_ = default__.safeIndex(17, (self.f19).length(0))
        lhs192_ = globalState
        lhs190_[lhs191_] = rhs228_
        lhs192_.f9 = rhs229_
        d_1356_v21_ = rhs230_
        d_1359_i2_: int
        d_1359_i2_ = 0
        with _dafny.label("14"):
            while (self.f19)[default__.safeIndex(17, (self.f19).length(0))]:
                with _dafny.c_label("14"):
                    if (d_1359_i2_) >= (100):
                        raise _dafny.Break("14")
                    d_1359_i2_ = (d_1359_i2_) + (1)
                    if (self).f18:
                        d_1360_v24_: str
                        d_1360_v24_ = _dafny.CodePoint('d')
                        d_1360_v24_ = d_1360_v24_
                        d_1361_v25_: _dafny.Array
                        def lambda69_(d_1362_i3_):
                            return (self).f18

                        init42_ = lambda69_
                        nw229_ = _dafny.Array(None, 14)
                        for i0_42_ in range(nw229_.length(0)):
                            nw229_[i0_42_] = init42_(i0_42_)
                        d_1361_v25_ = nw229_
                        d_1363_v26_: C7
                        nw230_ = C7()
                        nw230_.ctor__((self).f20, self.f21, d_1361_v25_, self.f17, (self).f24)
                        d_1363_v26_ = nw230_
                        d_1363_v26_ = d_1363_v26_
                        arr46_ = self.f19
                        index234_ = default__.safeIndex(17, (self.f19).length(0))
                        arr46_[index234_] = p1
                        d_1364_v27_: _dafny.Seq
                        d_1364_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kme"))
                        d_1365_v28_: T0
                        nw231_ = C4()
                        nw231_.ctor__(len(d_1358_v23_), p2, self.f17, p1)
                        d_1365_v28_ = nw231_
                        d_1366_v29_: _dafny.Seq
                        d_1366_v29_ = _dafny.SeqWithoutIsStrInference([d_1365_v28_, d_1365_v28_])
                        d_1367_v30_: _dafny.Array
                        nw232_ = _dafny.Array(None, 4)
                        nw232_[int(0)] = (p2) - (p2)
                        nw232_[int(1)] = (p2) - (len(d_1364_v27_))
                        nw232_[int(2)] = (d_1355_v20_)[default__.safeIndex(len(d_1366_v29_), len(d_1355_v20_))]
                        nw232_[int(3)] = default__.safeDivisionInt((self).f20, (self).f20)
                        d_1367_v30_ = nw232_
                        index235_ = default__.safeIndex(668, (d_1367_v30_).length(0))
                        (d_1367_v30_)[index235_] = p2
                        index236_ = default__.safeIndex(668, (d_1367_v30_).length(0))
                        (d_1367_v30_)[index236_] = (0) - (p0)
                        d_1358_v23_ = (d_1358_v23_).set((d_1367_v30_)[default__.safeIndex(668, (d_1367_v30_).length(0))], (p3 if (self.f19)[default__.safeIndex(17, (self.f19).length(0))] else p3))
                    elif True:
                        d_1368_v31_: str
                        d_1368_v31_ = _dafny.CodePoint('w')
                        rhs231_ = ((self).fm0((0) - (p0), d_1368_v31_, globalState)) * (p0)
                        rhs232_ = p0
                        lhs193_ = globalState
                        lhs194_ = globalState
                        lhs193_.f14 = rhs231_
                        lhs194_.f14 = rhs232_
                        (globalState).f14 = ((self).f20) + (p2)
                        d_1369_v32_: _dafny.Array
                        def lambda70_(d_1370_p2_):
                            def lambda71_(d_1371_i4_):
                                return (d_1371_i4_) - (d_1370_p2_)

                            return lambda71_

                        init43_ = lambda70_(p2)
                        nw233_ = _dafny.Array(None, 16)
                        for i0_43_ in range(nw233_.length(0)):
                            nw233_[i0_43_] = init43_(i0_43_)
                        d_1369_v32_ = nw233_
                        index237_ = default__.safeIndex(816, (d_1369_v32_).length(0))
                        (d_1369_v32_)[index237_] = self.f23
                        index238_ = default__.safeIndex(816, (d_1369_v32_).length(0))
                        (d_1369_v32_)[index238_] = (((self).f20) - (p0)) * (586)
                        index239_ = default__.safeIndex(816, (d_1369_v32_).length(0))
                        (d_1369_v32_)[index239_] = p2
                        (globalState).f3 = (True) not in (default__.fm16(globalState))
                    r0 = default__.safeModuloInt((self).f20, ((0) - (len(d_1358_v23_))) + (len(_dafny.Map({p3: (self.f19)[default__.safeIndex(17, (self.f19).length(0))]}))))
                    d_1372_v33_: _dafny.MultiSet
                    d_1372_v33_ = _dafny.MultiSet([p1])
                    d_1373_v34_: _dafny.Seq
                    d_1373_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pinsphqf"))
                    d_1374_v35_: _dafny.Array
                    nw234_ = _dafny.Array(None, 4)
                    nw234_[int(0)] = (_dafny.CodePoint('g')) in (d_1373_v34_)
                    nw234_[int(1)] = not((self).f24)
                    nw234_[int(2)] = default__.fm13(globalState)
                    nw234_[int(3)] = p1
                    d_1374_v35_ = nw234_
                    d_1375_v36_: _dafny.Map
                    d_1375_v36_ = _dafny.Map({p1: p3})
                    rhs233_ = default__.fm29(globalState)
                    rhs234_ = (self).f20
                    rhs235_ = not(((d_1375_v36_)[p1] if (p1) in (d_1375_v36_) else (self).f18))
                    rhs236_ = d_1374_v35_
                    lhs195_ = globalState
                    d_1372_v33_ = rhs233_
                    r0 = rhs234_
                    lhs195_.f3 = rhs235_
                    d_1374_v35_ = rhs236_
                    d_1376_v37_: C4
                    nw235_ = C4()
                    nw235_.ctor__(p2, len((d_1356_v21_).set(p1, p2)), self.f17, p3)
                    d_1376_v37_ = nw235_
                    d_1377_v38_: _dafny.Map
                    d_1377_v38_ = _dafny.Map({p0: d_1376_v37_})
                    d_1377_v38_ = (d_1377_v38_).set((p0) * ((self).f20), d_1376_v37_)
                    pass
            pass
        d_1378_v39_: _dafny.MultiSet
        d_1378_v39_ = _dafny.MultiSet([len(d_1356_v21_)])
        d_1379_v40_: str
        d_1379_v40_ = _dafny.CodePoint('i')
        d_1380_v41_: _dafny.Seq
        d_1380_v41_ = _dafny.SeqWithoutIsStrInference([d_1379_v40_])
        d_1381_v42_: D3
        d_1381_v42_ = D3_DC9()
        d_1382_v44_: _dafny.Set
        def iife103_():
            coll55_ = _dafny.Map()
            compr_55_: int
            for compr_55_ in _dafny.IntegerRange(662, 960):
                d_1383_v43_: int = compr_55_
                if ((662) <= (d_1383_v43_)) and ((d_1383_v43_) < (960)):
                    coll55_[(d_1383_v43_) - (p0)] = (self.f19)[default__.safeIndex(17, (self.f19).length(0))]
            return _dafny.Map(coll55_)
        d_1382_v44_ = _dafny.Set({d_1358_v23_, iife103_()
        })
        d_1384_v45_: D0
        d_1384_v45_ = D0_DC0(True)
        d_1385_v46_: D1
        d_1385_v46_ = D1_DC4(not(True), d_1382_v44_, default__.fm7((self).f20, (d_1356_v21_).set((self.f19)[default__.safeIndex(17, (self.f19).length(0))], p2), d_1384_v45_, globalState))
        d_1386_v47_: _dafny.Map
        d_1386_v47_ = _dafny.Map({(self).f24: (d_1357_v22_)[default__.safeIndex(self.f23, len(d_1357_v22_))]})
        d_1387_v48_: _dafny.Map
        d_1387_v48_ = _dafny.Map({p2: d_1379_v40_})
        d_1388_v49_: _dafny.Set
        d_1388_v49_ = _dafny.Set({p3})
        d_1389_v50_: _dafny.MultiSet
        d_1389_v50_ = _dafny.MultiSet([d_1379_v40_, ((d_1387_v48_)[len(d_1388_v49_)] if (len(d_1388_v49_)) in (d_1387_v48_) else d_1379_v40_)])
        d_1390_v51_: _dafny.Map
        d_1390_v51_ = _dafny.Map({p0: d_1357_v22_})
        d_1391_v52_: _dafny.Set
        d_1391_v52_ = _dafny.Set({len(((d_1390_v51_)[123] if (123) in (d_1390_v51_) else d_1357_v22_))})
        pat_let_tv17_ = p1
        d_1392_v53_: _dafny.Array
        nw236_ = _dafny.Array(None, 24)
        nw236_[int(0)] = ((d_1378_v39_).intersection(d_1378_v39_)).cardinality
        nw236_[int(1)] = len(d_1357_v22_)
        nw236_[int(2)] = len(d_1380_v41_)
        nw236_[int(3)] = self.f23
        nw236_[int(4)] = p0
        nw236_[int(5)] = p0
        nw236_[int(6)] = len(default__.fm34(d_1381_v42_, (self).f18, d_1379_v40_, p1, globalState))
        def iife104_(_pat_let24_0):
            def iife105_(d_1393_dt__update__tmp_h0_):
                def iife106_(_pat_let25_0):
                    def iife107_(d_1394_dt__update_hcf9_h0_):
                        return D1_DC4(d_1394_dt__update_hcf9_h0_, (d_1393_dt__update__tmp_h0_).cf10, (d_1393_dt__update__tmp_h0_).cf11)
                    return iife107_(_pat_let25_0)
                return iife106_(pat_let_tv17_)
            return iife105_(_pat_let24_0)
        nw236_[int(7)] = ((0) - (p2)) * ((iife104_(d_1385_v46_)).cf11)
        nw236_[int(8)] = (0) - (self.f23)
        nw236_[int(9)] = p2
        nw236_[int(10)] = len(d_1386_v47_)
        nw236_[int(11)] = (d_1389_v50_).cardinality
        nw236_[int(12)] = len(_dafny.Map({(self).f18: not((self.f19)[default__.safeIndex(17, (self.f19).length(0))])}))
        nw236_[int(13)] = len((_dafny.Map({p0: d_1379_v40_})) | (d_1387_v48_))
        nw236_[int(14)] = self.f23
        nw236_[int(15)] = (self).f20
        nw236_[int(16)] = p2
        nw236_[int(17)] = (self).f20
        nw236_[int(18)] = (self).f20
        nw236_[int(19)] = self.f23
        nw236_[int(20)] = len((d_1355_v20_ if not((self).f18) else d_1355_v20_))
        nw236_[int(21)] = len(d_1391_v52_)
        nw236_[int(22)] = (p0) * ((self).f20)
        nw236_[int(23)] = p0
        d_1392_v53_ = nw236_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1392_v53_).length(0)):
            d_1395_i5_: int = guard_loop_5_
            if (True) and (((0) <= (d_1395_i5_)) and ((d_1395_i5_) < ((d_1392_v53_).length(0)))):
                (d_1392_v53_)[(d_1395_i5_)] = (d_1395_i5_) * (p2)
        r0 = p2
        return r0

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: int = int(0)
        d_1396_v0_: _dafny.Set
        d_1396_v0_ = _dafny.Set({self.f23})
        d_1397_v1_: _dafny.Map
        d_1397_v1_ = _dafny.Map({self.f23: (self).f24})
        d_1398_v2_: T1
        nw237_ = C3()
        nw237_.ctor__(self.f19, _dafny.Set({self.f19, self.f19}), not(((d_1397_v1_)[p0] if (p0) in (d_1397_v1_) else False)))
        d_1398_v2_ = nw237_
        d_1399_v3_: _dafny.MultiSet
        d_1399_v3_ = _dafny.MultiSet([d_1398_v2_, d_1398_v2_])
        d_1400_v4_: _dafny.Seq
        d_1400_v4_ = _dafny.SeqWithoutIsStrInference([(d_1398_v2_).f18, default__.fm13(globalState)])
        d_1401_v5_: _dafny.Seq
        d_1401_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwffq"))
        d_1402_v6_: _dafny.Map
        d_1402_v6_ = _dafny.Map({self.f23: len(d_1401_v5_)})
        d_1403_v7_: _dafny.Map
        d_1403_v7_ = _dafny.Map({(d_1399_v3_).cardinality: (d_1400_v4_)[default__.safeIndex(len(d_1402_v6_), len(d_1400_v4_))]})
        arr47_ = self.f19
        index240_ = default__.safeIndex(79, (self.f19).length(0))
        arr47_[index240_] = (d_1396_v0_) != (_dafny.Set({(0) - (len(d_1397_v1_))}))
        arr48_ = self.f19
        index241_ = default__.safeIndex(79, (self.f19).length(0))
        arr48_[index241_] = (d_1398_v2_).f18
        d_1404_i0_: int
        d_1404_i0_ = 0
        with _dafny.label("15"):
            while True:
                with _dafny.c_label("15"):
                    if (d_1404_i0_) >= (100):
                        raise _dafny.Break("15")
                    d_1404_i0_ = (d_1404_i0_) + (1)
                    def lambda72_(d_1405_i1_):
                        return _dafny.CodePoint('r')

                    init44_ = lambda72_
                    nw238_ = _dafny.Array(None, 3)
                    for i0_44_ in range(nw238_.length(0)):
                        nw238_[i0_44_] = init44_(i0_44_)
                    (globalState).f16 = nw238_
                    if (d_1398_v2_).f18:
                        d_1406_v8_: _dafny.Array
                        nw239_ = _dafny.Array(_dafny.CodePoint('D'), 5)
                        d_1406_v8_ = nw239_
                        d_1407_v9_: str
                        d_1407_v9_ = _dafny.CodePoint('y')
                        index242_ = default__.safeIndex(190, (d_1406_v8_).length(0))
                        (d_1406_v8_)[index242_] = d_1407_v9_
                        index243_ = default__.safeIndex(190, (d_1406_v8_).length(0))
                        (d_1406_v8_)[index243_] = d_1407_v9_
                        d_1408_v10_: C3
                        nw240_ = C3()
                        nw240_.ctor__(d_1398_v2_.f19, self.f17, True)
                        d_1408_v10_ = nw240_
                        d_1409_v11_: _dafny.Map
                        d_1409_v11_ = _dafny.Map({self.f23: d_1408_v10_})
                        d_1410_v12_: D3
                        d_1410_v12_ = D3_DC9()
                        d_1411_v13_: _dafny.Seq
                        d_1411_v13_ = _dafny.SeqWithoutIsStrInference([d_1410_v12_, d_1410_v12_, d_1410_v12_])
                        d_1412_v14_: D5
                        d_1412_v14_ = D5_DC13(not((d_1398_v2_).f18), len(d_1409_v11_), d_1411_v13_)
                        d_1412_v14_ = d_1412_v14_
                        d_1413_v15_: C2
                        nw241_ = C2()
                        nw241_.ctor__(d_1398_v2_.f17, True)
                        d_1413_v15_ = nw241_
                        d_1414_v16_: _dafny.Map
                        d_1414_v16_ = _dafny.Map({(self.f23) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oed")))): d_1413_v15_})
                        d_1414_v16_ = d_1414_v16_
                        d_1415_v17_: _dafny.Array
                        nw242_ = _dafny.Array(int(0), 29)
                        d_1415_v17_ = nw242_
                        index244_ = default__.safeIndex(325, (d_1415_v17_).length(0))
                        (d_1415_v17_)[index244_] = self.f23
                        index245_ = default__.safeIndex(325, (d_1415_v17_).length(0))
                        (d_1415_v17_)[index245_] = (0) - (p0)
                        d_1415_v17_ = d_1415_v17_
                    elif True:
                        (globalState).f14 = self.f23
                        d_1416_v18_: C2
                        nw243_ = C2()
                        nw243_.ctor__((self.f17).intersection(d_1398_v2_.f17), ((self).f20) == ((self).f20))
                        d_1416_v18_ = nw243_
                        d_1417_v19_: _dafny.Map
                        d_1417_v19_ = _dafny.Map({(self.f19)[default__.safeIndex(79, (self.f19).length(0))]: (self).f20})
                        d_1418_v20_: _dafny.Seq
                        d_1418_v20_ = _dafny.SeqWithoutIsStrInference([(d_1417_v19_) | (d_1417_v19_)])
                        d_1419_v21_: _dafny.Map
                        d_1419_v21_ = _dafny.Map({d_1402_v6_: (self).f20})
                        rhs237_ = (self).fm1(d_1419_v21_, globalState)
                        rhs238_ = d_1418_v20_
                        lhs196_ = globalState
                        lhs196_.f9 = rhs237_
                        d_1418_v20_ = rhs238_
                        (d_1398_v2_).f19 = d_1398_v2_.f19
                        d_1420_v22_: _dafny.MultiSet
                        d_1420_v22_ = _dafny.MultiSet([(self).f24])
                        d_1421_v23_: _dafny.Map
                        d_1421_v23_ = _dafny.Map({True: (self).f18})
                        d_1422_v24_: D3
                        d_1422_v24_ = D3_DC10(d_1401_v5_, (d_1398_v2_).f18, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "te")), d_1401_v5_, len(d_1421_v23_))
                        d_1423_v25_: D11
                        d_1423_v25_ = D11_DC26((self.f19)[default__.safeIndex(79, (self.f19).length(0))], len(d_1401_v5_))
                        d_1424_v26_: _dafny.Array
                        nw244_ = _dafny.Array(None, 9)
                        nw244_[int(0)] = _dafny.MultiSet([(d_1398_v2_).f18])
                        nw244_[int(1)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))) | (d_1420_v22_)
                        nw244_[int(2)] = _dafny.MultiSet([(d_1422_v24_).cf17, True, (d_1398_v2_).f18, (d_1398_v2_).f18])
                        nw244_[int(3)] = d_1420_v22_
                        nw244_[int(4)] = d_1420_v22_
                        nw244_[int(5)] = (_dafny.MultiSet(d_1400_v4_)).set((d_1398_v2_).f18, default__.abs(self.f23))
                        nw244_[int(6)] = (default__.fm29(globalState)).intersection(d_1420_v22_)
                        nw244_[int(7)] = (d_1420_v22_) - (d_1420_v22_)
                        nw244_[int(8)] = _dafny.MultiSet([(d_1423_v25_).cf48])
                        d_1424_v26_ = nw244_
                        index246_ = default__.safeIndex(531, (d_1424_v26_).length(0))
                        (d_1424_v26_)[index246_] = d_1420_v22_
                        d_1425_v27_: _dafny.Seq
                        d_1425_v27_ = _dafny.SeqWithoutIsStrInference([default__.fm29(globalState), default__.fm29(globalState)])
                        index247_ = default__.safeIndex(531, (d_1424_v26_).length(0))
                        (d_1424_v26_)[index247_] = (d_1425_v27_)[default__.safeIndex(self.f23, len(d_1425_v27_))]
                    d_1397_v1_ = (d_1397_v1_).set((self).f20, (d_1398_v2_).f18)
                    d_1426_v28_: str
                    d_1426_v28_ = _dafny.CodePoint('h')
                    d_1427_v29_: _dafny.Map
                    d_1427_v29_ = _dafny.Map({(self).f20: _dafny.Set({len(d_1402_v6_), self.f23})})
                    d_1428_v30_: _dafny.Set
                    d_1428_v30_ = _dafny.Set({(self.f19)[default__.safeIndex(79, (self.f19).length(0))], (self.f19)[default__.safeIndex(79, (self.f19).length(0))]})
                    d_1429_v31_: _dafny.Seq
                    d_1429_v31_ = _dafny.SeqWithoutIsStrInference([(self).fm0(len(d_1402_v6_), d_1426_v28_, globalState), default__.safeDivisionInt((self).f20, (self).f20), default__.safeModuloInt((self).f20, len(((d_1427_v29_)[p0] if (p0) in (d_1427_v29_) else _dafny.Set({len(d_1428_v30_)}))))])
                    d_1430_v32_: _dafny.Map
                    d_1430_v32_ = _dafny.Map({(self).f24: p0})
                    d_1431_v33_: D0
                    d_1431_v33_ = D0_DC0(default__.fm13(globalState))
                    d_1432_v34_: D3
                    d_1432_v34_ = D3_DC9()
                    d_1433_v35_: _dafny.Seq
                    d_1433_v35_ = _dafny.SeqWithoutIsStrInference([d_1432_v34_, d_1432_v34_])
                    d_1434_v36_: _dafny.Seq
                    d_1434_v36_ = d_1433_v35_
                    d_1435_v38_: _dafny.Array
                    nw245_ = _dafny.Array(None, 16)
                    nw245_[int(0)] = self.f23
                    nw245_[int(1)] = default__.fm7((self).f20, d_1430_v32_, d_1431_v33_, globalState)
                    nw245_[int(2)] = (self).f20
                    nw245_[int(3)] = (p0) - (696)
                    nw245_[int(4)] = default__.safeModuloInt((0) - ((self).f20), (self).f20)
                    nw245_[int(5)] = default__.safeModuloInt(self.f23, 440)
                    nw245_[int(6)] = (0) - (p0)
                    nw245_[int(7)] = len(d_1428_v30_)
                    nw245_[int(8)] = (D5_DC13((self).f24, p0, d_1434_v36_)).cf24
                    nw245_[int(9)] = p0
                    nw245_[int(10)] = p0
                    nw245_[int(11)] = ((d_1430_v32_)[(self.f19)[default__.safeIndex(79, (self.f19).length(0))]] if ((self.f19)[default__.safeIndex(79, (self.f19).length(0))]) in (d_1430_v32_) else p0)
                    def iife108_():
                        coll56_ = _dafny.Map()
                        compr_56_: int
                        for compr_56_ in (d_1397_v1_).keys.Elements:
                            d_1436_v37_: int = compr_56_
                            if (d_1436_v37_) in (d_1397_v1_):
                                coll56_[(d_1436_v37_) - ((0) - (p0))] = d_1426_v28_
                        return _dafny.Map(coll56_)
                    nw245_[int(12)] = default__.safeDivisionInt(len(iife108_()
                    ), len(d_1400_v4_))
                    nw245_[int(13)] = default__.safeModuloInt(p0, self.f23)
                    nw245_[int(14)] = (self).f20
                    nw245_[int(15)] = ((d_1430_v32_)[(d_1398_v2_).f18] if ((d_1398_v2_).f18) in (d_1430_v32_) else (0) - (self.f23))
                    d_1435_v38_ = nw245_
                    index248_ = default__.safeIndex(352, (d_1435_v38_).length(0))
                    (d_1435_v38_)[index248_] = (self.f23) + (self.f23)
                    index249_ = default__.safeIndex(352, (d_1435_v38_).length(0))
                    rhs239_ = (d_1429_v31_) + (d_1429_v31_)
                    rhs240_ = default__.safeDivisionInt((0) - ((self).f20), (self).f20)
                    rhs241_ = p0
                    lhs197_ = d_1435_v38_
                    lhs198_ = default__.safeIndex(352, (d_1435_v38_).length(0))
                    lhs199_ = globalState
                    d_1429_v31_ = rhs239_
                    lhs197_[lhs198_] = rhs240_
                    lhs199_.f14 = rhs241_
                    pass
            pass
        d_1437_v39_: D11
        d_1437_v39_ = D11_DC27((d_1400_v4_).set(default__.safeIndex(self.f23, len(d_1400_v4_)), (self.f19)[default__.safeIndex(79, (self.f19).length(0))]), p0, (self).f20)
        d_1438_v40_: _dafny.Seq
        d_1438_v40_ = _dafny.SeqWithoutIsStrInference([d_1403_v7_])
        d_1439_v41_: _dafny.MultiSet
        d_1439_v41_ = _dafny.MultiSet([p0])
        d_1440_v42_: str
        d_1440_v42_ = _dafny.CodePoint('b')
        d_1441_v43_: _dafny.Set
        d_1441_v43_ = _dafny.Set({d_1397_v1_})
        d_1442_v44_: _dafny.Map
        d_1442_v44_ = _dafny.Map({(self).f24: False})
        d_1443_v45_: D1
        d_1443_v45_ = D1_DC4((self).f18, d_1441_v43_, len(d_1442_v44_))
        pat_let_tv18_ = d_1441_v43_
        d_1444_v46_: _dafny.Array
        nw246_ = _dafny.Array(None, 16)
        nw246_[int(0)] = (p0) + ((0) - (p0))
        nw246_[int(1)] = p0
        nw246_[int(2)] = self.f23
        nw246_[int(3)] = (0) - (((d_1402_v6_)[self.f23] if (self.f23) in (d_1402_v6_) else (self).f20))
        nw246_[int(4)] = (d_1437_v39_).cf52
        nw246_[int(5)] = (self).f20
        nw246_[int(6)] = (self).f20
        nw246_[int(7)] = (p0) + ((self).f20)
        nw246_[int(8)] = default__.safeDivisionInt(588, -132)
        nw246_[int(9)] = 323
        nw246_[int(10)] = len(_dafny.Map({(d_1438_v40_)[default__.safeIndex((self).f20, len(d_1438_v40_))]: (self).f20}))
        nw246_[int(11)] = default__.safeModuloInt((d_1439_v41_).cardinality, (self).f20)
        nw246_[int(12)] = p0
        nw246_[int(13)] = self.f23
        nw246_[int(14)] = (self).fm0(self.f23, d_1440_v42_, globalState)
        def iife109_(_pat_let26_0):
            def iife110_(d_1445_dt__update__tmp_h0_):
                def iife111_(_pat_let27_0):
                    def iife112_(d_1446_dt__update_hcf10_h0_):
                        return D1_DC4((d_1445_dt__update__tmp_h0_).cf9, d_1446_dt__update_hcf10_h0_, (d_1445_dt__update__tmp_h0_).cf11)
                    return iife112_(_pat_let27_0)
                return iife111_(pat_let_tv18_)
            return iife110_(_pat_let26_0)
        nw246_[int(15)] = ((self).f20 if (iife109_(d_1443_v45_)).cf9 else len(d_1396_v0_))
        d_1444_v46_ = nw246_
        index250_ = default__.safeIndex(610, (d_1444_v46_).length(0))
        (d_1444_v46_)[index250_] = 358
        index251_ = default__.safeIndex(610, (d_1444_v46_).length(0))
        (d_1444_v46_)[index251_] = (0) - (len((d_1401_v5_) + (d_1401_v5_)))
        rhs242_ = (self).f24
        rhs243_ = d_1400_v4_
        lhs200_ = globalState
        lhs200_.f3 = rhs242_
        d_1400_v4_ = rhs243_
        d_1447_v47_: _dafny.Map
        d_1447_v47_ = _dafny.Map({d_1440_v42_: (d_1444_v46_)[default__.safeIndex(610, (d_1444_v46_).length(0))]})
        d_1448_v48_: _dafny.Map
        d_1448_v48_ = _dafny.Map({(d_1447_v47_) | (d_1447_v47_): d_1442_v44_})
        d_1449_v50_: _dafny.Set
        d_1449_v50_ = _dafny.Set({d_1440_v42_})
        def iife113_():
            coll57_ = _dafny.Map()
            compr_57_: str
            for compr_57_ in (d_1449_v50_).Elements:
                d_1450_v49_: str = compr_57_
                if (d_1450_v49_) in (d_1449_v50_):
                    coll57_[d_1450_v49_] = (D5_DC14((d_1398_v2_).f18, len(d_1401_v5_), p0)).cf28
            return _dafny.Map(coll57_)
        d_1448_v48_ = (d_1448_v48_).set(iife113_()
        , d_1442_v44_)
        d_1451_v51_: _dafny.Array
        nw247_ = _dafny.Array(_dafny.Array(None, 0), 27)
        d_1451_v51_ = nw247_
        index252_ = default__.safeIndex(608, (d_1451_v51_).length(0))
        (d_1451_v51_)[index252_] = d_1444_v46_
        index253_ = default__.safeIndex(608, (d_1451_v51_).length(0))
        (d_1451_v51_)[index253_] = d_1444_v46_
        r0 = self.f23
        r1 = not((self).f18)
        r2 = default__.fm4(((d_1444_v46_)[default__.safeIndex(610, (d_1444_v46_).length(0))]) + (len(_dafny.SeqWithoutIsStrInference([d_1440_v42_ for d_1452_i2_ in range(default__.abs(649))]))), globalState)
        return r0, r1, r2

    @property
    def f24(self):
        return self._f24

class C9(T2, T1, T0):
    def  __init__(self):
        self._f17: _dafny.Set = _dafny.Set({})
        self._f19: _dafny.Array = _dafny.Array(None, 0)
        self._f21: _dafny.Map = _dafny.Map({})
        self._f20: int = int(0)
        self._f18: bool = False
        self.f22: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f17(self):
        return self._f17
    @f17.setter
    def f17(self, value):
        self._f17 = value
    @property
    def f19(self):
        return self._f19
    @f19.setter
    def f19(self, value):
        self._f19 = value
    @property
    def f21(self):
        return self._f21
    @f21.setter
    def f21(self, value):
        self._f21 = value
    @property
    def f20(self):
        return self._f20
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f22, f20, f21, f19, f17, f18):
        (self).f22 = f22
        (self)._f20 = f20
        (self).f21 = f21
        (self).f19 = f19
        (self).f17 = f17
        (self)._f18 = f18

    def fm1(self, p0, globalState):
        return (self).f18

    def fm0(self, p0, p1, globalState):
        return (0) - (len((_dafny.Map({(self).f18: (self).f18})) | ((_dafny.Map({True: (self).f18})) | (_dafny.Map({(self).f18: (self).f18})))))

    def fm2(self, p0, p1, p2, globalState):
        return (770) - ((self).f20)

    def m0(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_1453_v0_: str
        d_1453_v0_ = _dafny.CodePoint('t')
        d_1454_v1_: _dafny.Set
        d_1454_v1_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))), d_1453_v0_)})
        d_1454_v1_ = d_1454_v1_
        d_1455_v2_: _dafny.Seq
        d_1455_v2_ = _dafny.SeqWithoutIsStrInference([not((self).f18)])
        d_1456_v3_: _dafny.Seq
        d_1456_v3_ = _dafny.SeqWithoutIsStrInference([d_1455_v2_, d_1455_v2_])
        d_1457_v4_: _dafny.Array
        nw248_ = _dafny.Array(None, 6)
        nw248_[int(0)] = d_1456_v3_
        nw248_[int(1)] = d_1456_v3_
        nw248_[int(2)] = d_1456_v3_
        nw248_[int(3)] = d_1456_v3_
        nw248_[int(4)] = d_1456_v3_
        nw248_[int(5)] = (d_1456_v3_) + (d_1456_v3_)
        d_1457_v4_ = nw248_
        index254_ = default__.safeIndex(92, (d_1457_v4_).length(0))
        (d_1457_v4_)[index254_] = d_1456_v3_
        arr49_ = self.f19
        index255_ = default__.safeIndex(965, (self.f19).length(0))
        arr49_[index255_] = p1
        d_1458_v5_: _dafny.MultiSet
        d_1458_v5_ = _dafny.MultiSet([not(p1)])
        d_1459_v7_: _dafny.Map
        d_1459_v7_ = _dafny.Map({self.f22: (self).f18})
        d_1460_v8_: _dafny.Seq
        d_1460_v8_ = _dafny.SeqWithoutIsStrInference([p2, self.f22, p0])
        index256_ = default__.safeIndex(92, (d_1457_v4_).length(0))
        arr50_ = self.f19
        index257_ = default__.safeIndex(965, (self.f19).length(0))
        def iife114_():
            coll58_ = _dafny.Map()
            compr_58_: int
            for compr_58_ in (d_1459_v7_).keys.Elements:
                d_1462_v6_: int = compr_58_
                if (d_1462_v6_) in (d_1459_v7_):
                    coll58_[(d_1462_v6_) * ((0) - (self.f22))] = self.f22
            return _dafny.Map(coll58_)
        rhs244_ = not((p1 if (self).f18 else not((self).f18)))
        rhs245_ = not(((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f18])) if (self).f18 else d_1458_v5_)).isdisjoint(d_1458_v5_))
        rhs246_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p1, (self).f18, p1, not(p1), (self).f18]) for d_1461_i0_ in range(default__.abs(-140))])
        rhs247_ = (self).fm1(_dafny.Map({iife114_()
        : len(d_1460_v8_)}), globalState)
        lhs201_ = d_1457_v4_
        lhs202_ = default__.safeIndex(92, (d_1457_v4_).length(0))
        lhs203_ = self.f19
        lhs204_ = default__.safeIndex(965, (self.f19).length(0))
        r0 = rhs244_
        r0 = rhs245_
        lhs201_[lhs202_] = rhs246_
        lhs203_[lhs204_] = rhs247_
        d_1463_i1_: int
        d_1463_i1_ = 0
        with _dafny.label("16"):
            while (self.f19)[default__.safeIndex(965, (self.f19).length(0))]:
                with _dafny.c_label("16"):
                    if (d_1463_i1_) >= (100):
                        raise _dafny.Break("16")
                    d_1463_i1_ = (d_1463_i1_) + (1)
                    d_1464_v9_: _dafny.Array
                    def lambda73_(d_1465_v0_):
                        def lambda74_(d_1466_i2_):
                            return d_1465_v0_

                        return lambda74_

                    init45_ = lambda73_(d_1453_v0_)
                    nw249_ = _dafny.Array(None, 10)
                    for i0_45_ in range(nw249_.length(0)):
                        nw249_[i0_45_] = init45_(i0_45_)
                    d_1464_v9_ = nw249_
                    index258_ = default__.safeIndex(102, (d_1464_v9_).length(0))
                    (d_1464_v9_)[index258_] = d_1453_v0_
                    d_1467_v10_: _dafny.Seq
                    d_1467_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmfhaq"))
                    d_1468_v11_: _dafny.MultiSet
                    d_1468_v11_ = _dafny.MultiSet([d_1453_v0_])
                    d_1469_v12_: _dafny.Map
                    d_1469_v12_ = _dafny.Map({(d_1468_v11_).cardinality: p2})
                    d_1470_v13_: _dafny.Map
                    d_1470_v13_ = _dafny.Map({d_1469_v12_: p0})
                    index259_ = default__.safeIndex(102, (d_1464_v9_).length(0))
                    rhs248_ = len((d_1467_v10_) + (d_1467_v10_))
                    rhs249_ = (self).fm1(d_1470_v13_, globalState)
                    rhs250_ = d_1453_v0_
                    rhs251_ = 795
                    lhs205_ = self
                    lhs206_ = d_1464_v9_
                    lhs207_ = default__.safeIndex(102, (d_1464_v9_).length(0))
                    lhs205_.f22 = rhs248_
                    r0 = rhs249_
                    lhs206_[lhs207_] = rhs250_
                    r2 = rhs251_
                    d_1471_v14_: _dafny.Map
                    d_1471_v14_ = _dafny.Map({(self.f19)[default__.safeIndex(965, (self.f19).length(0))]: True})
                    if (not((self.f19)[default__.safeIndex(965, (self.f19).length(0))])) in (d_1471_v14_):
                        (globalState).f3 = p1
                        arr51_ = self.f19
                        index260_ = default__.safeIndex(965, (self.f19).length(0))
                        rhs252_ = self.f22
                        rhs253_ = ((d_1464_v9_)[default__.safeIndex(102, (d_1464_v9_).length(0))]) not in (d_1467_v10_)
                        rhs254_ = (self.f19)[default__.safeIndex(965, (self.f19).length(0))]
                        rhs255_ = (self).fm0((d_1460_v8_)[default__.safeIndex((self).f20, len(d_1460_v8_))], d_1453_v0_, globalState)
                        rhs256_ = ((self).f20) in (d_1469_v12_)
                        lhs208_ = globalState
                        lhs209_ = globalState
                        lhs210_ = self.f19
                        lhs211_ = default__.safeIndex(965, (self.f19).length(0))
                        r2 = rhs252_
                        lhs208_.f3 = rhs253_
                        r1 = rhs254_
                        lhs209_.f14 = rhs255_
                        lhs210_[lhs211_] = rhs256_
                        d_1467_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdrmxxf"))
                        r0 = (not(p1)) or (((self.f19)[default__.safeIndex(965, (self.f19).length(0))] if p1 else (d_1455_v2_)[default__.safeIndex((self).f20, len(d_1455_v2_))]))
                        arr52_ = self.f19
                        index261_ = default__.safeIndex(965, (self.f19).length(0))
                        arr52_[index261_] = (d_1453_v0_) not in ((d_1467_v10_) + (d_1467_v10_))
                    elif True:
                        d_1472_v15_: _dafny.Array
                        def lambda75_(d_1473_i3_):
                            return (self.f19)[default__.safeIndex(965, (self.f19).length(0))]

                        init46_ = lambda75_
                        nw250_ = _dafny.Array(None, 15)
                        for i0_46_ in range(nw250_.length(0)):
                            nw250_[i0_46_] = init46_(i0_46_)
                        d_1472_v15_ = nw250_
                        d_1474_v16_: T2
                        nw251_ = C6()
                        nw251_.ctor__(_dafny.Map({(self.f19)[default__.safeIndex(965, (self.f19).length(0))]: (self).f20}), (self).f20, (self).f20, self.f21, d_1472_v15_, self.f17, (self).f18)
                        d_1474_v16_ = nw251_
                        d_1475_v17_: _dafny.Set
                        d_1475_v17_ = _dafny.Set({d_1474_v16_, d_1474_v16_})
                        (globalState).f14 = (len(d_1467_v10_) if (self.f19)[default__.safeIndex(965, (self.f19).length(0))] else ((self).f20) * (len(d_1475_v17_)))
                        (globalState).f14 = self.f22
                        d_1476_v18_: C3
                        nw252_ = C3()
                        nw252_.ctor__(d_1472_v15_, (_dafny.Set({d_1472_v15_})) | (self.f17), ((d_1474_v16_).f20) not in (d_1459_v7_))
                        d_1476_v18_ = nw252_
                        index262_ = default__.safeIndex(102, (d_1464_v9_).length(0))
                        (d_1464_v9_)[index262_] = (d_1464_v9_)[default__.safeIndex(102, (d_1464_v9_).length(0))]
                        arr53_ = self.f19
                        index263_ = default__.safeIndex(965, (self.f19).length(0))
                        arr53_[index263_] = not((self.f19)[default__.safeIndex(965, (self.f19).length(0))])
                    if p1:
                        d_1477_v19_: _dafny.Array
                        nw253_ = _dafny.Array(False, 28)
                        d_1477_v19_ = nw253_
                        d_1478_v20_: _dafny.Map
                        d_1478_v20_ = _dafny.Map({d_1477_v19_: d_1477_v19_})
                        d_1478_v20_ = ((d_1478_v20_) | (d_1478_v20_)) | (d_1478_v20_)
                        d_1479_v21_: _dafny.Map
                        d_1479_v21_ = _dafny.Map({default__.fm21(_dafny.SeqWithoutIsStrInference([(self).f20, 231]), globalState): p2})
                        d_1479_v21_ = (d_1479_v21_).set(d_1459_v7_, p0)
                        d_1480_v22_: C2
                        nw254_ = C2()
                        nw254_.ctor__(self.f17, default__.fm13(globalState))
                        d_1480_v22_ = nw254_
                        d_1481_v23_: D11
                        d_1481_v23_ = D11_DC27((d_1455_v2_).set(default__.safeIndex((self).f20, len(d_1455_v2_)), (self).f18), p0, p2)
                        d_1482_v24_: D0
                        d_1482_v24_ = D0_DC1((d_1467_v10_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcpfqafh"))).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcpfqafh")))), (d_1464_v9_)[default__.safeIndex(102, (d_1464_v9_).length(0))])), default__.safeDivisionInt((0) - ((d_1481_v23_).cf51), p0), p0, d_1453_v0_)
                        d_1482_v24_ = d_1482_v24_
                        (globalState).f9 = p1
                    elif True:
                        d_1483_v25_: _dafny.Array
                        nw255_ = _dafny.Array(None, 11)
                        nw255_[int(0)] = (self).f18
                        nw255_[int(1)] = (self).f18
                        nw255_[int(2)] = (self).f18
                        nw255_[int(3)] = default__.fm13(globalState)
                        nw255_[int(4)] = not((self).fm1(d_1470_v13_, globalState))
                        nw255_[int(5)] = (self).f18
                        nw255_[int(6)] = False
                        nw255_[int(7)] = (self.f19)[default__.safeIndex(965, (self.f19).length(0))]
                        nw255_[int(8)] = (self).f18
                        nw255_[int(9)] = p1
                        nw255_[int(10)] = (self).f18
                        d_1483_v25_ = nw255_
                        d_1484_v27_: C8
                        nw256_ = C8()
                        def iife115_():
                            coll59_ = _dafny.Map()
                            compr_59_: str
                            for compr_59_ in (d_1468_v11_).Elements:
                                d_1485_v26_: str = compr_59_
                                if (d_1485_v26_) in (d_1468_v11_):
                                    coll59_[d_1485_v26_] = p1
                            return _dafny.Map(coll59_)
                        nw256_.ctor__(p0, not((p1) == (((d_1459_v7_)[self.f22] if (self.f22) in (d_1459_v7_) else (self).f18))), d_1483_v25_, self.f17, (self).fm1(d_1470_v13_, globalState), p2, (iife115_()
                        ) | (self.f21))
                        d_1484_v27_ = nw256_
                        d_1484_v27_ = d_1484_v27_
                        d_1483_v25_ = d_1483_v25_
                        d_1486_v28_: C0
                        nw257_ = C0()
                        nw257_.ctor__()
                        d_1486_v28_ = nw257_
                        d_1486_v28_ = d_1486_v28_
                        d_1487_v29_: _dafny.Set
                        d_1487_v29_ = _dafny.Set({(self.f19)[default__.safeIndex(965, (self.f19).length(0))]})
                        d_1488_v30_: _dafny.Seq
                        d_1488_v30_ = d_1455_v2_
                        d_1489_v31_: _dafny.MultiSet
                        d_1489_v31_ = _dafny.MultiSet([(self).f20])
                        d_1490_v32_: _dafny.Array
                        nw258_ = _dafny.Array(None, 17)
                        nw258_[int(0)] = (d_1455_v2_) + ((d_1455_v2_).set(default__.safeIndex(p0, len(d_1455_v2_)), False))
                        nw258_[int(1)] = (d_1455_v2_ if (d_1484_v27_).f24 else d_1455_v2_)
                        nw258_[int(2)] = d_1455_v2_
                        nw258_[int(3)] = (d_1455_v2_).set(default__.safeIndex(506, len(d_1455_v2_)), (d_1484_v27_).f24)
                        nw258_[int(4)] = (d_1488_v30_)
                        nw258_[int(5)] = d_1455_v2_
                        nw258_[int(6)] = d_1455_v2_
                        nw258_[int(7)] = (d_1455_v2_).set(default__.safeIndex(p0, len(d_1455_v2_)), False)
                        nw258_[int(8)] = (d_1455_v2_) + (default__.fm16(globalState))
                        nw258_[int(9)] = d_1455_v2_
                        nw258_[int(10)] = default__.fm16(globalState)
                        nw258_[int(11)] = d_1455_v2_
                        nw258_[int(12)] = (_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18, (self).f18, p1, p1])) + (d_1455_v2_)
                        nw258_[int(13)] = d_1455_v2_
                        nw258_[int(14)] = default__.fm16(globalState)
                        nw258_[int(15)] = d_1455_v2_
                        nw258_[int(16)] = ((d_1455_v2_).set(default__.safeIndex((d_1489_v31_).cardinality, len(d_1455_v2_)), (self).f18)) + (d_1455_v2_)
                        d_1490_v32_ = nw258_
                        index264_ = default__.safeIndex(2, (d_1490_v32_).length(0))
                        (d_1490_v32_)[index264_] = d_1455_v2_
                        d_1491_v33_: _dafny.Seq
                        d_1491_v33_ = _dafny.SeqWithoutIsStrInference([default__.fm37((self).f18, self.f22, (self).f18, (d_1484_v27_).f24, globalState)])
                        d_1492_v34_: _dafny.MultiSet
                        d_1492_v34_ = _dafny.MultiSet([((d_1491_v33_)[default__.safeIndex(len(d_1471_v14_), len(d_1491_v33_))]).intersection(d_1487_v29_)])
                        index265_ = default__.safeIndex(2, (d_1490_v32_).length(0))
                        rhs257_ = _dafny.Set({(len(d_1467_v10_)) == (p2)})
                        rhs258_ = (len(d_1460_v8_)) * (len(default__.fm40(p1, (self.f19)[default__.safeIndex(965, (self.f19).length(0))], globalState)))
                        rhs259_ = (d_1455_v2_) + (d_1455_v2_)
                        rhs260_ = d_1492_v34_
                        lhs212_ = self
                        lhs213_ = d_1490_v32_
                        lhs214_ = default__.safeIndex(2, (d_1490_v32_).length(0))
                        d_1487_v29_ = rhs257_
                        lhs212_.f22 = rhs258_
                        lhs213_[lhs214_] = rhs259_
                        d_1492_v34_ = rhs260_
                        d_1493_v35_: _dafny.Map
                        d_1493_v35_ = _dafny.Map({d_1483_v25_: (self).f20})
                        rhs261_ = d_1493_v35_
                        rhs262_ = self.f22
                        rhs263_ = True
                        rhs264_ = (self).f20
                        rhs265_ = (0) - ((self).f20)
                        lhs215_ = globalState
                        lhs216_ = globalState
                        lhs217_ = globalState
                        d_1493_v35_ = rhs261_
                        lhs215_.f14 = rhs262_
                        lhs216_.f9 = rhs263_
                        lhs217_.f14 = rhs264_
                        r2 = rhs265_
                    r1 = not((self).fm1((_dafny.Map({d_1469_v12_: p2})) | (d_1470_v13_), globalState))
                    pass
            pass
        d_1494_v36_: _dafny.Seq
        d_1494_v36_ = _dafny.SeqWithoutIsStrInference([d_1453_v0_])
        d_1495_v37_: _dafny.Array
        nw259_ = _dafny.Array(int(0), 15)
        d_1495_v37_ = nw259_
        d_1496_v38_: _dafny.Map
        d_1496_v38_ = _dafny.Map({d_1453_v0_: d_1495_v37_})
        d_1497_v39_: _dafny.Array
        nw260_ = _dafny.Array(_dafny.CodePoint('D'), 17)
        d_1497_v39_ = nw260_
        d_1498_v40_: _dafny.Map
        d_1498_v40_ = _dafny.Map({p0: p2})
        d_1499_v41_: _dafny.Set
        d_1499_v41_ = _dafny.Set({d_1459_v7_, d_1459_v7_})
        d_1500_v42_: D1
        d_1500_v42_ = D1_DC4((self.f19)[default__.safeIndex(965, (self.f19).length(0))], d_1499_v41_, -445)
        d_1501_v43_: _dafny.Seq
        d_1501_v43_ = _dafny.SeqWithoutIsStrInference([d_1500_v42_])
        d_1502_v44_: _dafny.MultiSet
        d_1502_v44_ = _dafny.MultiSet([(self).f20, (self).f20, ((d_1498_v40_)[self.f22] if (self.f22) in (d_1498_v40_) else p2), len(d_1501_v43_), p2])
        d_1503_v45_: _dafny.Set
        d_1503_v45_ = _dafny.Set({len(_dafny.Map({(D14_DC34(d_1497_v39_)).cf64: d_1453_v0_})), ((d_1502_v44_)[346] if (346) in (d_1502_v44_) else (self).f20)})
        d_1504_v46_: _dafny.Map
        d_1504_v46_ = _dafny.Map({d_1453_v0_: len(d_1503_v45_)})
        d_1505_v47_: _dafny.Seq
        d_1505_v47_ = _dafny.SeqWithoutIsStrInference([d_1504_v46_])
        d_1506_v48_: _dafny.Map
        d_1506_v48_ = _dafny.Map({d_1505_v47_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymwc"))})
        rhs266_ = d_1494_v36_
        rhs267_ = d_1496_v38_
        rhs268_ = p0
        rhs269_ = ((((d_1494_v36_).set(default__.safeIndex(p0, len(d_1494_v36_)), d_1453_v0_)) + ((((d_1506_v48_)[d_1505_v47_] if (d_1505_v47_) in (d_1506_v48_) else d_1494_v36_)) + (d_1494_v36_))).set(default__.safeIndex(len((d_1455_v2_) + ((d_1455_v2_).set(default__.safeIndex(100, len(d_1455_v2_)), (self.f19)[default__.safeIndex(965, (self.f19).length(0))]))), len(((d_1494_v36_).set(default__.safeIndex(p0, len(d_1494_v36_)), d_1453_v0_)) + ((((d_1506_v48_)[d_1505_v47_] if (d_1505_v47_) in (d_1506_v48_) else d_1494_v36_)) + (d_1494_v36_)))), d_1453_v0_)).set(default__.safeIndex(len(d_1494_v36_), len((((d_1494_v36_).set(default__.safeIndex(p0, len(d_1494_v36_)), d_1453_v0_)) + ((((d_1506_v48_)[d_1505_v47_] if (d_1505_v47_) in (d_1506_v48_) else d_1494_v36_)) + (d_1494_v36_))).set(default__.safeIndex(len((d_1455_v2_) + ((d_1455_v2_).set(default__.safeIndex(100, len(d_1455_v2_)), (self.f19)[default__.safeIndex(965, (self.f19).length(0))]))), len(((d_1494_v36_).set(default__.safeIndex(p0, len(d_1494_v36_)), d_1453_v0_)) + ((((d_1506_v48_)[d_1505_v47_] if (d_1505_v47_) in (d_1506_v48_) else d_1494_v36_)) + (d_1494_v36_)))), d_1453_v0_))), d_1453_v0_)
        rhs270_ = p0
        lhs218_ = globalState
        d_1494_v36_ = rhs266_
        d_1496_v38_ = rhs267_
        lhs218_.f14 = rhs268_
        d_1494_v36_ = rhs269_
        r2 = rhs270_
        d_1507_i4_: int
        d_1507_i4_ = 0
        with _dafny.label("17"):
            while not (p1) or ((self).f18):
                with _dafny.c_label("17"):
                    if (d_1507_i4_) >= (100):
                        raise _dafny.Break("17")
                    d_1507_i4_ = (d_1507_i4_) + (1)
                    (self).f22 = p2
                    d_1504_v46_ = (d_1504_v46_).set(d_1453_v0_, (self).f20)
                    d_1508_v49_: _dafny.Map
                    d_1508_v49_ = _dafny.Map({d_1455_v2_: d_1453_v0_})
                    d_1509_v50_: _dafny.Array
                    nw261_ = _dafny.Array(None, 3)
                    nw261_[int(0)] = d_1508_v49_
                    nw261_[int(1)] = d_1508_v49_
                    nw261_[int(2)] = d_1508_v49_
                    d_1509_v50_ = nw261_
                    index266_ = default__.safeIndex(82, (d_1509_v50_).length(0))
                    (d_1509_v50_)[index266_] = d_1508_v49_
                    index267_ = default__.safeIndex(82, (d_1509_v50_).length(0))
                    rhs271_ = (self).f18
                    rhs272_ = (d_1503_v45_) - (d_1503_v45_)
                    rhs273_ = _dafny.Map({d_1455_v2_: d_1453_v0_})
                    rhs274_ = (self.f19)[default__.safeIndex(965, (self.f19).length(0))]
                    lhs219_ = globalState
                    lhs220_ = d_1509_v50_
                    lhs221_ = default__.safeIndex(82, (d_1509_v50_).length(0))
                    lhs222_ = globalState
                    lhs219_.f9 = rhs271_
                    d_1503_v45_ = rhs272_
                    lhs220_[lhs221_] = rhs273_
                    lhs222_.f3 = rhs274_
                    (globalState).f3 = (p1) and (p1)
                    pass
            pass
        d_1510_i5_: int
        d_1510_i5_ = 0
        with _dafny.label("18"):
            while (self).f18:
                with _dafny.c_label("18"):
                    if (d_1510_i5_) >= (100):
                        raise _dafny.Break("18")
                    d_1510_i5_ = (d_1510_i5_) + (1)
                    r0 = not((self).f18)
                    if False:
                        d_1511_v51_: _dafny.Map
                        d_1511_v51_ = _dafny.Map({not ((self.f19)[default__.safeIndex(965, (self.f19).length(0))]) or ((self).f18): not((d_1500_v42_).cf9)})
                        d_1512_v52_: T0
                        nw262_ = C2()
                        nw262_.ctor__(self.f17, p1)
                        d_1512_v52_ = nw262_
                        d_1513_v53_: _dafny.MultiSet
                        d_1513_v53_ = _dafny.MultiSet([d_1512_v52_])
                        d_1511_v51_ = (d_1511_v51_).set((d_1500_v42_).cf9, (_dafny.MultiSet([d_1512_v52_])).isdisjoint(d_1513_v53_))
                        d_1453_v0_ = d_1453_v0_
                        d_1455_v2_ = ((d_1455_v2_).set(default__.safeIndex(len((d_1494_v36_).set(default__.safeIndex((self).f20, len(d_1494_v36_)), d_1453_v0_)), len(d_1455_v2_)), (self.f19)[default__.safeIndex(965, (self.f19).length(0))])) + ((_dafny.SeqWithoutIsStrInference([(self.f19)[default__.safeIndex(965, (self.f19).length(0))]])).set(default__.safeIndex(self.f22, len(_dafny.SeqWithoutIsStrInference([(self.f19)[default__.safeIndex(965, (self.f19).length(0))]]))), (d_1512_v52_).f18))
                        d_1514_v54_: C4
                        nw263_ = C4()
                        nw263_.ctor__(len((d_1459_v7_) | (d_1459_v7_)), p0, self.f17, (self).fm1(_dafny.Map({d_1498_v40_: p0}), globalState))
                        d_1514_v54_ = nw263_
                        (d_1514_v54_).f26 = len(((d_1455_v2_) + (default__.fm16(globalState)) if ((self).fm2((d_1512_v52_).f18, len(d_1494_v36_), default__.fm43(globalState), globalState)) >= (len(d_1494_v36_)) else d_1455_v2_))
                    elif True:
                        d_1515_v55_: D0
                        d_1515_v55_ = D0_DC0((self.f19)[default__.safeIndex(965, (self.f19).length(0))])
                        def iife116_(_pat_let28_0):
                            def iife117_(d_1516_dt__update__tmp_h0_):
                                def iife118_(_pat_let29_0):
                                    def iife119_(d_1517_dt__update_hcf0_h0_):
                                        return D0_DC0(d_1517_dt__update_hcf0_h0_)
                                    return iife119_(_pat_let29_0)
                                return iife118_((self).f18)
                            return iife117_(_pat_let28_0)
                        d_1453_v0_ = default__.fm33(default__.fm7(p2, _dafny.Map({p1: self.f22}), iife116_(d_1515_v55_), globalState), (self.f19)[default__.safeIndex(965, (self.f19).length(0))], globalState)
                        d_1518_v56_: _dafny.Array
                        nw264_ = _dafny.Array(False, 17)
                        d_1518_v56_ = nw264_
                        d_1519_v57_: C7
                        nw265_ = C7()
                        nw265_.ctor__(self.f22, self.f21, d_1518_v56_, self.f17, p1)
                        d_1519_v57_ = nw265_
                        d_1520_v58_: _dafny.Map
                        d_1520_v58_ = _dafny.Map({d_1519_v57_: (self).f20})
                        d_1520_v58_ = (d_1520_v58_).set(d_1519_v57_, len(d_1494_v36_))
                        (globalState).f3 = p1
                        d_1521_v59_: _dafny.Array
                        nw266_ = _dafny.Array(_dafny.Map({}), 7)
                        d_1521_v59_ = nw266_
                        d_1522_v60_: _dafny.Map
                        d_1522_v60_ = _dafny.Map({d_1453_v0_: d_1518_v56_})
                        d_1523_v61_: C5
                        nw267_ = C5()
                        nw267_.ctor__(d_1494_v36_, (self).f18, p0, self.f21, ((d_1522_v60_)[d_1453_v0_] if (d_1453_v0_) in (d_1522_v60_) else d_1518_v56_), _dafny.Set({d_1518_v56_}), False)
                        d_1523_v61_ = nw267_
                        d_1524_v62_: _dafny.Map
                        d_1524_v62_ = _dafny.Map({True: (self.f19)[default__.safeIndex(965, (self.f19).length(0))]})
                        index268_ = default__.safeIndex(900, (d_1521_v59_).length(0))
                        (d_1521_v59_)[index268_] = (_dafny.Map({d_1523_v61_: d_1524_v62_})).set(d_1523_v61_, d_1524_v62_)
                        d_1525_v63_: _dafny.Map
                        d_1525_v63_ = _dafny.Map({d_1458_v5_: d_1523_v61_})
                        d_1526_v64_: _dafny.Map
                        d_1526_v64_ = _dafny.Map({((d_1525_v63_)[d_1458_v5_] if (d_1458_v5_) in (d_1525_v63_) else d_1523_v61_): d_1524_v62_})
                        index269_ = default__.safeIndex(900, (d_1521_v59_).length(0))
                        (d_1521_v59_)[index269_] = d_1526_v64_
                        d_1527_v67_: _dafny.Map
                        def iife120_():
                            coll60_ = _dafny.Map()
                            compr_60_: int
                            for compr_60_ in (d_1460_v8_).Elements:
                                d_1528_v66_: int = compr_60_
                                if (d_1528_v66_) in (d_1460_v8_):
                                    coll60_[(d_1528_v66_) + (p2)] = (0) - ((self).f20)
                            return _dafny.Map(coll60_)
                        d_1527_v67_ = _dafny.Map({iife120_()
                        : (self).f20})
                        def iife121_():
                            coll61_ = _dafny.Map()
                            compr_61_: _dafny.Map
                            for compr_61_ in (d_1527_v67_).keys.Elements:
                                d_1529_v65_: _dafny.Map = compr_61_
                                if (d_1529_v65_) in (d_1527_v67_):
                                    coll61_[d_1529_v65_] = p0
                            return _dafny.Map(coll61_)
                        (globalState).f3 = (self).fm1(iife121_()
                        , globalState)
                    d_1530_v68_: _dafny.Array
                    def lambda76_(d_1531_i6_):
                        return (self.f19)[default__.safeIndex(965, (self.f19).length(0))]

                    init47_ = lambda76_
                    nw268_ = _dafny.Array(None, 8)
                    for i0_47_ in range(nw268_.length(0)):
                        nw268_[i0_47_] = init47_(i0_47_)
                    d_1530_v68_ = nw268_
                    d_1532_v70_: C8
                    nw269_ = C8()
                    def iife122_():
                        coll62_ = _dafny.Map()
                        compr_62_: str
                        for compr_62_ in (d_1504_v46_).keys.Elements:
                            d_1533_v69_: str = compr_62_
                            if (d_1533_v69_) in (d_1504_v46_):
                                coll62_[d_1533_v69_] = p1
                        return _dafny.Map(coll62_)
                    nw269_.ctor__(p0, (self.f19)[default__.safeIndex(965, (self.f19).length(0))], d_1530_v68_, self.f17, (self).f18, p2, iife122_()
                    )
                    d_1532_v70_ = nw269_
                    d_1534_v71_: _dafny.Array
                    nw270_ = _dafny.Array(None, 10)
                    nw270_[int(0)] = d_1532_v70_
                    nw270_[int(1)] = d_1532_v70_
                    nw270_[int(2)] = d_1532_v70_
                    nw270_[int(3)] = d_1532_v70_
                    nw270_[int(4)] = d_1532_v70_
                    nw270_[int(5)] = d_1532_v70_
                    nw270_[int(6)] = d_1532_v70_
                    nw270_[int(7)] = d_1532_v70_
                    nw270_[int(8)] = d_1532_v70_
                    nw270_[int(9)] = d_1532_v70_
                    d_1534_v71_ = nw270_
                    d_1535_v72_: _dafny.Set
                    d_1535_v72_ = _dafny.Set({d_1534_v71_})
                    d_1535_v72_ = d_1535_v72_
                    (self).f22 = p2
                    pass
            pass
        d_1536_v73_: _dafny.Set
        d_1536_v73_ = _dafny.Set({(self.f19)[default__.safeIndex(965, (self.f19).length(0))], (self).f18})
        r0 = (len(d_1494_v36_)) > (len(d_1536_v73_))
        r1 = not(((_dafny.MultiSet([p0]) if not((self).f18) else d_1502_v44_)).isdisjoint(d_1502_v44_))
        r2 = (p2) * (p0)
        return r0, r1, r2

