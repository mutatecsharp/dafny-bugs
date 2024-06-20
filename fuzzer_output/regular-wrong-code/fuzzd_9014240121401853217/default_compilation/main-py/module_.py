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
        if (True) in (_dafny.SeqWithoutIsStrInference([False, not(True)])):
            return ((_dafny.MultiSet([-675])).cardinality) * (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))), 873, 829, len(_dafny.Map({_dafny.CodePoint('b'): True}))})))
        elif True:
            return ((0) - ((_dafny.MultiSet([431])).cardinality)) * (590)

    @staticmethod
    def fm1(p0, globalState):
        return not(False)

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return (_dafny.Map({True: 914})) | ((_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_0_i0_ in range(default__.abs(-832))]))})) | (_dafny.Map({True: (0) - (len(_dafny.SeqWithoutIsStrInference([True, False])))})))

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        return D0_DC1((_dafny.Set({True, False})).issubset(_dafny.Set({True})))

    @staticmethod
    def fm6(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(716, 849):
                d_1_v0_: int = compr_0_
                if ((716) <= (d_1_v0_)) and ((d_1_v0_) < (849)):
                    coll0_[default__.safeDivisionInt(d_1_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ja"))))] = -249
            return _dafny.Map(coll0_)
        return _dafny.Set({len((iife0_()
        ) | (_dafny.Map({125: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")))})))})

    @staticmethod
    def fm7(globalState):
        return (_dafny.MultiSet([True])) | ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False), True]))) - (_dafny.MultiSet([True])))

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khrapsoyo"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))

    @staticmethod
    def fm9(p0, globalState):
        return _dafny.CodePoint('r')

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(410, 725):
                d_2_v0_: int = compr_1_
                if ((410) <= (d_2_v0_)) and ((d_2_v0_) < (725)):
                    coll1_ = coll1_.union(_dafny.Set([(d_2_v0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsnmoy"))))]))
            return _dafny.Set(coll1_)
        def iife2_():
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(291, 256):
                    d_3_v1_: int = compr_4_
                    if ((291) <= (d_3_v1_)) and ((d_3_v1_) < (256)):
                        coll4_[default__.safeDivisionInt(d_3_v1_, len(_dafny.Set({-295, (0) - ((_dafny.MultiSet([-593, 360])).cardinality)})))] = True
                return _dafny.Map(coll4_)
            coll2_ = _dafny.Set()
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in _dafny.IntegerRange(291, 256):
                    d_3_v1_: int = compr_3_
                    if ((291) <= (d_3_v1_)) and ((d_3_v1_) < (256)):
                        coll3_[default__.safeDivisionInt(d_3_v1_, len(_dafny.Set({-295, (0) - ((_dafny.MultiSet([-593, 360])).cardinality)})))] = True
                return _dafny.Map(coll3_)
            compr_2_: int
            for compr_2_ in (iife3_()
            ).keys.Elements:
                d_4_v2_: int = compr_2_
                if (d_4_v2_) in (iife4_()
                ):
                    coll2_ = coll2_.union(_dafny.Set([(d_4_v2_) * (208)]))
            return _dafny.Set(coll2_)
        return _dafny.SeqWithoutIsStrInference([(iife1_()
        ).issubset(iife2_()
        ), False, (478) not in (_dafny.Set({-756, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdw"))): _dafny.CodePoint('m')}))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smixm"))) for d_5_i1_ in range(default__.abs(497))])), 338])) for d_6_i0_ in range(default__.abs(-709))]))})), True])

    @staticmethod
    def fm13(globalState):
        def iife5_():
            def iife7_():
                coll7_ = _dafny.Map()
                compr_7_: _dafny.Map
                for compr_7_ in (_dafny.Set({_dafny.Map({True: False})})).Elements:
                    d_7_v0_: _dafny.Map = compr_7_
                    if (d_7_v0_) in (_dafny.Set({_dafny.Map({True: False})})):
                        coll7_[d_7_v0_] = True
                return _dafny.Map(coll7_)
            coll5_ = _dafny.Set()
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: _dafny.Map
                for compr_6_ in (_dafny.Set({_dafny.Map({True: False})})).Elements:
                    d_7_v0_: _dafny.Map = compr_6_
                    if (d_7_v0_) in (_dafny.Set({_dafny.Map({True: False})})):
                        coll6_[d_7_v0_] = True
                return _dafny.Map(coll6_)
            compr_5_: _dafny.Map
            for compr_5_ in (iife6_()
            ).keys.Elements:
                d_8_v1_: _dafny.Map = compr_5_
                if (d_8_v1_) in (iife7_()
                ):
                    coll5_ = coll5_.union(_dafny.Set([d_8_v1_]))
            return _dafny.Set(coll5_)
        return (_dafny.Set({_dafny.Map({True: False}), _dafny.Map({True: not(False)}), _dafny.Map({True: False})})) | ((_dafny.Set({_dafny.Map({True: False})})) - (iife5_()
        ))

    @staticmethod
    def fm14(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dybqgrbwg"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxkb"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_9_i0_ in range(default__.abs(756))])))

    @staticmethod
    def fm15(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([-520 for d_10_i0_ in range(default__.abs(300))]))) + (_dafny.SeqWithoutIsStrInference([421, len(_dafny.Set({False})), 469]))

    @staticmethod
    def fm16(globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(874, 620):
                d_11_v0_: int = compr_8_
                if ((874) <= (d_11_v0_)) and ((d_11_v0_) < (620)):
                    coll8_[(d_11_v0_) * (751)] = -828
            return _dafny.Map(coll8_)
        source0_ = D13_DC38(_dafny.Map({False: len(iife8_()
)}))
        if source0_.is_DC39:
            if False:
                return _dafny.CodePoint('h')
            elif True:
                return _dafny.CodePoint('t')
        elif source0_.is_DC40:
            d_12___mcc_h0_ = source0_.cf56
            d_13___mcc_h1_ = source0_.cf57
            d_14___mcc_h2_ = source0_.cf58
            d_15_cf58_ = d_14___mcc_h2_
            d_16_cf57_ = d_13___mcc_h1_
            d_17_cf56_ = d_12___mcc_h0_
            return _dafny.CodePoint('p')
        elif True:
            d_18___mcc_h3_ = source0_.cf55
            d_19_cf55_ = d_18___mcc_h3_
            return _dafny.CodePoint('u')

    @staticmethod
    def fm17(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({False})) - (_dafny.Set({True}))) - ((_dafny.Set({False})) | (_dafny.Set({True})))

    @staticmethod
    def fm19(globalState):
        return _dafny.Map({(822 if False else -43): _dafny.CodePoint('t')})

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        if (_dafny.SeqWithoutIsStrInference([_dafny.Set({not(False)}), _dafny.Set({False}), _dafny.Set({True}), _dafny.Set({not(False)})])) != (_dafny.SeqWithoutIsStrInference([_dafny.Set({True, not(not(True))})])):
            return _dafny.CodePoint('a')
        elif True:
            return _dafny.CodePoint('p')

    @staticmethod
    def fm21(globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_20_i0_ in range(default__.abs(791))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "curdiyset")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))

    @staticmethod
    def fm22(p0, globalState):
        source1_ = D12_DC37(D12_DC36(True, -997, False))
        if source1_.is_DC34:
            d_21___mcc_h0_ = source1_.cf48
            d_22___mcc_h1_ = source1_.cf49
            d_23___mcc_h2_ = source1_.cf50
            d_24_cf50_ = d_23___mcc_h2_
            d_25_cf49_ = d_22___mcc_h1_
            d_26_cf48_ = d_21___mcc_h0_
            return _dafny.Map({(0) - (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_27_i0_ in range(default__.abs(26))])), 192}))): True})
        elif source1_.is_DC35:
            return (_dafny.Map({882: False})) | (_dafny.Map({144: True}))
        elif source1_.is_DC36:
            d_28___mcc_h3_ = source1_.cf51
            d_29___mcc_h4_ = source1_.cf52
            d_30___mcc_h5_ = source1_.cf53
            d_31_cf53_ = d_30___mcc_h5_
            d_32_cf52_ = d_29___mcc_h4_
            d_33_cf51_ = d_28___mcc_h3_
            return _dafny.Map({d_32_cf52_: d_33_cf51_})
        elif source1_.is_DC33:
            d_34___mcc_h6_ = source1_.cf47
            d_35_cf47_ = d_34___mcc_h6_
            def iife9_():
                coll9_ = _dafny.Map()
                compr_9_: int
                for compr_9_ in _dafny.IntegerRange(997, 164):
                    d_36_v0_: int = compr_9_
                    if ((997) <= (d_36_v0_)) and ((d_36_v0_) < (164)):
                        coll9_[(d_36_v0_) - (492)] = False
                return _dafny.Map(coll9_)
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(766, 670):
                    d_37_v1_: int = compr_10_
                    if ((766) <= (d_37_v1_)) and ((d_37_v1_) < (670)):
                        coll10_[(d_37_v1_) * (len(d_35_cf47_))] = not(not(True))
                return _dafny.Map(coll10_)
            return (iife9_()
            ) | (iife10_()
            )
        elif True:
            d_38___mcc_h7_ = source1_.cf54
            d_39_cf54_ = d_38___mcc_h7_
            return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "slrpor"))): False})) | (_dafny.Map({765: False}))

    @staticmethod
    def fm23(p0, p1, globalState):
        def iife11_():
            coll11_ = _dafny.Set()
            compr_11_: _dafny.Map
            for compr_11_ in (_dafny.Map({_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([not(False)]))): len(_dafny.Set({False, False, False, False}))}): True})).keys.Elements:
                d_40_v0_: _dafny.Map = compr_11_
                if (d_40_v0_) in (_dafny.Map({_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([not(False)]))): len(_dafny.Set({False, False, False, False}))}): True})):
                    coll11_ = coll11_.union(_dafny.Set([d_40_v0_]))
            return _dafny.Set(coll11_)
        return (_dafny.MultiSet([(_dafny.MultiSet([True])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgn")))])) - ((_dafny.MultiSet([(0) - ((0) - (-41)), len(_dafny.Map({False: False})), 458])) - (_dafny.MultiSet([len(iife11_()
        )])))

    @staticmethod
    def fm24(p0, globalState):
        return _dafny.Map({not (not(False)) or ((D9_DC24(not(not(True)), False, True)).cf35): _dafny.Map({True: 855})})

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + ((_dafny.SeqWithoutIsStrInference([False, True])) + (_dafny.SeqWithoutIsStrInference([(D12_DC34(D0_DC2(D0_DC1(False)), False, True)).cf50, True, not(True), False])))

    @staticmethod
    def fm26(p0, p1, p2, p3, globalState):
        source2_ = D3_DC10((D3_DC10(D3_DC9(528))).cf12)
        if source2_.is_DC9:
            d_41___mcc_h0_ = source2_.cf11
            d_42_cf11_ = d_41___mcc_h0_
            def iife12_():
                coll12_ = _dafny.Map()
                compr_12_: _dafny.Map
                for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): True})])).Elements:
                    d_43_v0_: _dafny.Map = compr_12_
                    if (d_43_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): True})])):
                        coll12_[d_43_v0_] = len(_dafny.Map({d_42_cf11_: (0) - (len(_dafny.SeqWithoutIsStrInference([False, True, True, True])))}))
                return _dafny.Map(coll12_)
            return D2_DC7(-394, iife12_()
, 890)
        elif source2_.is_DC8:
            d_44___mcc_h1_ = source2_.cf10
            d_45_cf10_ = d_44___mcc_h1_
            return D2_DC7(338, _dafny.Map({_dafny.Map({False: False}): -77}), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_46_i0_ in range(default__.abs(373))])))
        elif True:
            d_47___mcc_h2_ = source2_.cf12
            d_48_cf12_ = d_47___mcc_h2_
            def iife13_():
                coll13_ = _dafny.Map()
                compr_13_: _dafny.Map
                for compr_13_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: False}), _dafny.Map({True: True})])).Elements:
                    d_50_v1_: _dafny.Map = compr_13_
                    if (d_50_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: False}), _dafny.Map({True: True})])):
                        coll13_[d_50_v1_] = 762
                return _dafny.Map(coll13_)
            return D2_DC7(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_49_i1_ in range(default__.abs(115))])), iife13_()
, 107)

    @staticmethod
    def fm27(globalState):
        if not (False) or (not(False)):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_51_i0_ in range(default__.abs(455))]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_52_i1_ in range(default__.abs(838))])

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        return _dafny.Set({_dafny.CodePoint('r')})

    @staticmethod
    def fm29(globalState):
        if not(not((not(False)) and (True))):
            return _dafny.CodePoint('i')
        elif True:
            return _dafny.CodePoint('n')

    @staticmethod
    def fm30(globalState):
        source3_ = (D11_DC30() if False else D11_DC30())
        if source3_.is_DC30:
            return _dafny.Set({-94, 271})
        elif source3_.is_DC31:
            d_53___mcc_h0_ = source3_.cf44
            d_54___mcc_h1_ = source3_.cf45
            d_55_cf45_ = d_54___mcc_h1_
            d_56_cf44_ = d_53___mcc_h0_
            return (_dafny.Set({d_56_cf44_})).intersection(_dafny.Set({(0) - (len(_dafny.Set({False})))}))
        elif source3_.is_DC29:
            d_57___mcc_h2_ = source3_.cf43
            d_58_cf43_ = d_57___mcc_h2_
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(591, 801):
                    d_59_v0_: int = compr_14_
                    if ((591) <= (d_59_v0_)) and ((d_59_v0_) < (801)):
                        coll14_[(d_59_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aydffd"))))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikjyjyifl")))
                return _dafny.Map(coll14_)
            return (_dafny.Set({413})) | (_dafny.Set({-448, len(iife14_()
            )}))
        elif True:
            d_60___mcc_h3_ = source3_.cf46
            d_61_cf46_ = d_60___mcc_h3_
            return (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ba")))})).intersection(_dafny.Set({len(_dafny.Set({False, False})), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ryykijq"))): _dafny.CodePoint('t')}))}))

    @staticmethod
    def fm31(globalState):
        def iife15_():
            coll15_ = _dafny.Map()
            compr_15_: int
            for compr_15_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(0) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_62_i1_ in range(default__.abs(280))]))])).cardinality): not(False)})) for d_63_i0_ in range(default__.abs(-981))])).Elements:
                d_64_v0_: int = compr_15_
                if (d_64_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(0) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_62_i1_ in range(default__.abs(280))]))])).cardinality): not(False)})) for d_63_i0_ in range(default__.abs(-981))])):
                    coll15_[(d_64_v0_) + (806)] = True
            return _dafny.Map(coll15_)
        return ((_dafny.MultiSet([D0_DC2(D0_DC0(_dafny.Map({(_dafny.MultiSet([_dafny.MultiSet([-475])])).cardinality: False}))), D0_DC2(D0_DC2(D0_DC0(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgwt"))): True})))), D0_DC2(D0_DC2(D0_DC0(iife15_()
)))])).intersection(_dafny.MultiSet([D0_DC2(D0_DC0(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([True]))): False})))]))).intersection((_dafny.MultiSet([D0_DC2(D0_DC2(D0_DC1(True))), D0_DC2(D0_DC1(False)), D0_DC2(D0_DC2(D0_DC1(True)))])).intersection(_dafny.MultiSet([D0_DC2(D0_DC2(D0_DC2(D0_DC2(D0_DC1(True)))))])))

    @staticmethod
    def fm32(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([D3_DC9((0) - (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(True)])).cardinality, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_65_i0_ in range(default__.abs(754))])): False}))])))), D3_DC9((0) - ((_dafny.MultiSet([(_dafny.MultiSet([False])).cardinality, 678, 813, len(_dafny.Set({973, 62, (0) - (-221), len(_dafny.Map({True: not(True)})), len(_dafny.SeqWithoutIsStrInference([True]))}))])).cardinality))])) + (_dafny.SeqWithoutIsStrInference([D3_DC9(-407)]))

    @staticmethod
    def fm33(p0, p1, globalState):
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pptit"))): D8_DC21(False, False, _dafny.SeqWithoutIsStrInference([False, True]), True, 249)})) | (_dafny.Map({len(_dafny.Map({-630: -844})): D8_DC21(True, True, _dafny.SeqWithoutIsStrInference([False]), not(False), -398)}))) | ((_dafny.Map({123: D8_DC21(True, True, _dafny.SeqWithoutIsStrInference([False]), False, len(_dafny.SeqWithoutIsStrInference([48])))})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ska"))): D8_DC21(True, not(not(False)), _dafny.SeqWithoutIsStrInference([False, True, False, False, True]), True, -541)})))

    @staticmethod
    def fm34(p0, p1, p2, globalState):
        return (_dafny.Map({_dafny.Set({False, False}): len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([False, True])}))})) | (_dafny.Map({_dafny.Set({not(True), not(False), False, False, not(True)}): 299}))

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        d_66_v0_: _dafny.Map
        d_66_v0_ = _dafny.Map({p0: _dafny.MultiSet([p0, p0, p0, False, p0])})
        d_67_v1_: D0
        d_67_v1_ = D0_DC1(p0)
        d_68_v2_: _dafny.MultiSet
        d_68_v2_ = _dafny.MultiSet([p0, (d_67_v1_).cf1])
        d_69_v3_: D10
        d_69_v3_ = D10_DC27(p0, 921, p1, ((d_66_v0_)[not(not(True))] if (not(not(True))) in (d_66_v0_) else d_68_v2_))
        d_70_v4_: _dafny.Seq
        d_70_v4_ = _dafny.SeqWithoutIsStrInference([p0])
        d_71_v5_: _dafny.Set
        d_71_v5_ = _dafny.Set({(0) - (len((d_70_v4_).set(default__.safeIndex(p2, len(d_70_v4_)), not(p0)))), p2, p2, 947})
        if (default__.safeModuloInt((0) - (((p3)[True] if (True) in (p3) else ((d_69_v3_).cf41).cardinality)), p2)) in (d_71_v5_):
            (globalState).f0 = p0
            d_72_v6_: C2
            nw0_ = C2()
            nw0_.ctor__(p0)
            d_72_v6_ = nw0_
            d_72_v6_ = d_72_v6_
            d_73_v7_: _dafny.Seq
            d_73_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "osmkakap"))
            d_73_v7_ = (_dafny.SeqWithoutIsStrInference([p1 for d_74_i0_ in range(default__.abs(993))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "baefwmqnu")))
            d_75_v8_: _dafny.MultiSet
            d_75_v8_ = _dafny.MultiSet([p1, p1])
            d_76_v9_: _dafny.Array
            def lambda0_(d_77_v6_):
                def lambda1_(d_78_i1_):
                    return not((d_77_v6_).f13)

                return lambda1_

            init0_ = lambda0_(d_72_v6_)
            nw1_ = _dafny.Array(None, 24)
            for i0_0_ in range(nw1_.length(0)):
                nw1_[i0_0_] = init0_(i0_0_)
            d_76_v9_ = nw1_
            d_79_v10_: C0
            nw2_ = C0()
            nw2_.ctor__((p2) != (p2), (D4_DC13((d_75_v8_).cardinality, d_76_v9_)).cf20)
            d_79_v10_ = nw2_
            d_80_v11_: _dafny.Array
            nw3_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
            d_80_v11_ = nw3_
            index0_ = default__.safeIndex(776, (d_80_v11_).length(0))
            (d_80_v11_)[index0_] = default__.fm14(d_73_v7_, p2, globalState)
            index1_ = default__.safeIndex(776, (d_80_v11_).length(0))
            (d_80_v11_)[index1_] = (d_73_v7_) + (d_73_v7_)
        elif True:
            d_81_v12_: _dafny.Map
            d_81_v12_ = _dafny.Map({p2: not ((d_69_v3_).cf38) or (p0)})
            d_81_v12_ = d_81_v12_
            d_82_v13_: _dafny.Seq
            d_82_v13_ = _dafny.SeqWithoutIsStrInference([p2])
            d_83_v14_: _dafny.Seq
            d_83_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxwm"))
            rhs0_ = d_82_v13_
            rhs1_ = not(p0)
            rhs2_ = p0
            rhs3_ = len(d_83_v14_)
            lhs0_ = globalState
            lhs1_ = globalState
            lhs2_ = globalState
            d_82_v13_ = rhs0_
            lhs0_.f0 = rhs1_
            lhs1_.f0 = rhs2_
            lhs2_.f5 = rhs3_
            d_84_v15_: _dafny.Map
            d_84_v15_ = _dafny.Map({(0) - (default__.fm0(len(d_83_v14_), p2, len(d_83_v14_), p0, globalState)): len(_dafny.SeqWithoutIsStrInference([p2 for d_85_i2_ in range(default__.abs(561))]))})
            d_84_v15_ = (d_84_v15_).set(len(_dafny.Set({False, p0})), len(d_81_v12_))
            d_86_v16_: _dafny.Array
            def lambda2_(d_87_p0_):
                def lambda3_(d_88_i3_):
                    return d_87_p0_

                return lambda3_

            init1_ = lambda2_(p0)
            nw4_ = _dafny.Array(None, 23)
            for i0_1_ in range(nw4_.length(0)):
                nw4_[i0_1_] = init1_(i0_1_)
            d_86_v16_ = nw4_
            index2_ = default__.safeIndex(634, (d_86_v16_).length(0))
            (d_86_v16_)[index2_] = p0
            index3_ = default__.safeIndex(634, (d_86_v16_).length(0))
            (d_86_v16_)[index3_] = (p2) > (p2)
            index4_ = default__.safeIndex(634, (d_86_v16_).length(0))
            (d_86_v16_)[index4_] = (p0) not in (d_70_v4_)
        (globalState).f2 = p0
        if default__.fm1(p0, globalState):
            d_89_v17_: _dafny.Array
            nw5_ = _dafny.Array(False, 14)
            d_89_v17_ = nw5_
            index5_ = default__.safeIndex(884, (d_89_v17_).length(0))
            (d_89_v17_)[index5_] = False
            index6_ = default__.safeIndex(884, (d_89_v17_).length(0))
            (d_89_v17_)[index6_] = not(False)
            if p0:
                d_90_v18_: _dafny.Array
                def lambda4_(d_91_p2_):
                    def lambda5_(d_92_i4_):
                        return default__.safeModuloInt(d_92_i4_, d_91_p2_)

                    return lambda5_

                init2_ = lambda4_(p2)
                nw6_ = _dafny.Array(None, 5)
                for i0_2_ in range(nw6_.length(0)):
                    nw6_[i0_2_] = init2_(i0_2_)
                d_90_v18_ = nw6_
                d_93_v19_: _dafny.Map
                d_93_v19_ = _dafny.Map({p2: d_90_v18_})
                d_94_v20_: _dafny.Map
                d_94_v20_ = _dafny.Map({True: p0})
                d_95_v21_: _dafny.Seq
                d_95_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikfpgwiid"))
                d_96_v22_: T1
                nw7_ = C3()
                nw7_.ctor__(p2)
                d_96_v22_ = nw7_
                d_97_v23_: _dafny.Map
                d_97_v23_ = _dafny.Map({d_96_v22_: p2})
                d_98_v24_: C2
                nw8_ = C2()
                nw8_.ctor__(p0)
                d_98_v24_ = nw8_
                d_99_v25_: _dafny.Map
                d_99_v25_ = _dafny.Map({p2: d_98_v24_})
                d_100_v26_: _dafny.Map
                d_100_v26_ = _dafny.Map({_dafny.CodePoint('a'): len(d_99_v25_)})
                d_101_v27_: _dafny.Map
                d_101_v27_ = _dafny.Map({d_89_v17_: 806})
                d_102_v28_: _dafny.Set
                d_102_v28_ = _dafny.Set({(d_89_v17_)[default__.safeIndex(884, (d_89_v17_).length(0))], p0})
                d_103_v29_: _dafny.Array
                nw9_ = _dafny.Array(None, 21)
                nw9_[int(0)] = p2
                nw9_[int(1)] = p2
                nw9_[int(2)] = (p2) - (len(d_93_v19_))
                nw9_[int(3)] = p2
                nw9_[int(4)] = p2
                nw9_[int(5)] = p2
                nw9_[int(6)] = p2
                nw9_[int(7)] = p2
                nw9_[int(8)] = ((p3)[p0] if (p0) in (p3) else (0) - (default__.fm0(len(d_94_v20_), p2, len(d_95_v21_), p0, globalState)))
                nw9_[int(9)] = len(d_97_v23_)
                nw9_[int(10)] = (0) - (p2)
                nw9_[int(11)] = p2
                nw9_[int(12)] = p2
                nw9_[int(13)] = default__.safeDivisionInt(p2, (0) - (p2))
                nw9_[int(14)] = ((d_100_v26_)[p1] if (p1) in (d_100_v26_) else p2)
                nw9_[int(15)] = p2
                nw9_[int(16)] = (0) - (p2)
                nw9_[int(17)] = p2
                nw9_[int(18)] = (0) - (len((d_101_v27_) | (d_101_v27_)))
                nw9_[int(19)] = len(d_102_v28_)
                nw9_[int(20)] = p2
                d_103_v29_ = nw9_
                d_104_v30_: C0
                nw10_ = C0()
                nw10_.ctor__((d_89_v17_)[default__.safeIndex(884, (d_89_v17_).length(0))], d_89_v17_)
                d_104_v30_ = nw10_
                d_105_v31_: _dafny.Set
                d_105_v31_ = _dafny.Set({d_104_v30_})
                d_106_v32_: _dafny.Map
                d_106_v32_ = _dafny.Map({(_dafny.MultiSet([p2, len(d_105_v31_)])).cardinality: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbcudgtk"))) for d_107_i5_ in range(default__.abs(592))])})
                d_108_v33_: _dafny.Seq
                d_108_v33_ = _dafny.SeqWithoutIsStrInference([p2])
                index7_ = default__.safeIndex(420, (d_103_v29_).length(0))
                (d_103_v29_)[index7_] = len((((d_106_v32_)[p2] if (p2) in (d_106_v32_) else _dafny.SeqWithoutIsStrInference([552]))) + (d_108_v33_))
                d_109_v34_: _dafny.Array
                def lambda6_(d_110_p1_):
                    def lambda7_(d_111_i6_):
                        return _dafny.SeqWithoutIsStrInference([d_110_p1_ for d_112_i7_ in range(default__.abs(209))])

                    return lambda7_

                init3_ = lambda6_(p1)
                nw11_ = _dafny.Array(None, 25)
                for i0_3_ in range(nw11_.length(0)):
                    nw11_[i0_3_] = init3_(i0_3_)
                d_109_v34_ = nw11_
                index8_ = default__.safeIndex(420, (d_103_v29_).length(0))
                rhs4_ = p2
                rhs5_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agjcjgg"))) < ((d_95_v21_) + (_dafny.SeqWithoutIsStrInference([p1 for d_113_i8_ in range(default__.abs(944))])))
                rhs6_ = p2
                rhs7_ = d_109_v34_
                lhs3_ = d_103_v29_
                lhs4_ = default__.safeIndex(420, (d_103_v29_).length(0))
                lhs5_ = globalState
                lhs6_ = globalState
                lhs3_[lhs4_] = rhs4_
                lhs5_.f2 = rhs5_
                lhs6_.f5 = rhs6_
                d_109_v34_ = rhs7_
                index9_ = default__.safeIndex(420, (d_103_v29_).length(0))
                rhs8_ = p0
                rhs9_ = default__.fm0((d_103_v29_)[default__.safeIndex(420, (d_103_v29_).length(0))], (d_103_v29_)[default__.safeIndex(420, (d_103_v29_).length(0))], p2, p0, globalState)
                rhs10_ = ((d_103_v29_)[default__.safeIndex(420, (d_103_v29_).length(0))]) * ((d_103_v29_)[default__.safeIndex(420, (d_103_v29_).length(0))])
                lhs7_ = globalState
                lhs8_ = globalState
                lhs9_ = d_103_v29_
                lhs10_ = default__.safeIndex(420, (d_103_v29_).length(0))
                lhs7_.f2 = rhs8_
                lhs8_.f5 = rhs9_
                lhs9_[lhs10_] = rhs10_
                d_114_v35_: C4
                nw12_ = C4()
                nw12_.ctor__((d_103_v29_)[default__.safeIndex(420, (d_103_v29_).length(0))])
                d_114_v35_ = nw12_
                d_115_v36_: _dafny.Map
                d_115_v36_ = _dafny.Map({p2: d_114_v35_})
                d_116_v37_: _dafny.Seq
                d_116_v37_ = _dafny.SeqWithoutIsStrInference([d_114_v35_, ((d_115_v36_)[len(d_102_v28_)] if (len(d_102_v28_)) in (d_115_v36_) else d_114_v35_), d_114_v35_])
                d_116_v37_ = (d_116_v37_) + ((d_116_v37_).set(default__.safeIndex(p2, len(d_116_v37_)), d_114_v35_))
                d_117_v38_: str
                d_117_v38_ = _dafny.CodePoint('s')
                d_117_v38_ = (p1 if (d_98_v24_).f13 else p1)
                (globalState).f0 = (True) not in (p3)
            elif True:
                d_118_v39_: _dafny.Array
                nw13_ = _dafny.Array(None, 11)
                d_118_v39_ = nw13_
                d_119_v40_: _dafny.Seq
                d_119_v40_ = _dafny.SeqWithoutIsStrInference([p2])
                d_120_v41_: C2
                nw14_ = C2()
                nw14_.ctor__((d_70_v4_)[default__.safeIndex((d_119_v40_)[default__.safeIndex(p2, len(d_119_v40_))], len(d_70_v4_))])
                d_120_v41_ = nw14_
                index10_ = default__.safeIndex(618, (d_118_v39_).length(0))
                (d_118_v39_)[index10_] = d_120_v41_
                index11_ = default__.safeIndex(618, (d_118_v39_).length(0))
                nw15_ = C2()
                nw15_.ctor__((d_120_v41_).f13)
                (d_118_v39_)[index11_] = nw15_
                (globalState).f2 = p0
                d_121_v42_: T0
                nw16_ = C2()
                nw16_.ctor__(p0)
                d_121_v42_ = nw16_
                d_122_v43_: _dafny.Map
                d_122_v43_ = _dafny.Map({d_121_v42_: p2})
                d_123_v44_: _dafny.Map
                d_123_v44_ = _dafny.Map({len(d_122_v43_): (d_89_v17_)[default__.safeIndex(884, (d_89_v17_).length(0))]})
                d_124_v45_: D0
                d_124_v45_ = D0_DC0(d_123_v44_)
                d_68_v2_ = _dafny.MultiSet([not((False) and ((d_89_v17_)[default__.safeIndex(884, (d_89_v17_).length(0))])), not((d_120_v41_).fm4(d_124_v45_, d_124_v45_, d_68_v2_, globalState))])
                d_125_v46_: _dafny.Seq
                d_125_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "colybgq"))
                d_126_v47_: _dafny.Map
                d_126_v47_ = _dafny.Map({d_125_v46_: (d_119_v40_) + (_dafny.SeqWithoutIsStrInference([652]))})
                d_127_v48_: _dafny.Seq
                d_127_v48_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p2 for d_128_i9_ in range(default__.abs(131))])])
                d_126_v47_ = (d_126_v47_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqy")), (d_127_v48_)[default__.safeIndex(-912, len(d_127_v48_))])
                d_129_v49_: C3
                nw17_ = C3()
                nw17_.ctor__(p2)
                d_129_v49_ = nw17_
            d_130_v50_: C1
            nw18_ = C1()
            nw18_.ctor__()
            d_130_v50_ = nw18_
            (globalState).f0 = (d_68_v2_).issubset(d_68_v2_)
            d_131_v51_: D7
            d_131_v51_ = D7_DC18()
            d_132_v52_: _dafny.Seq
            d_132_v52_ = _dafny.SeqWithoutIsStrInference([d_131_v51_])
            d_133_v53_: _dafny.Seq
            d_133_v53_ = _dafny.SeqWithoutIsStrInference([d_132_v52_])
            d_134_v54_: _dafny.Map
            d_134_v54_ = _dafny.Map({(d_89_v17_)[default__.safeIndex(884, (d_89_v17_).length(0))]: d_133_v53_})
            index12_ = default__.safeIndex(884, (d_89_v17_).length(0))
            (d_89_v17_)[index12_] = (len(((d_134_v54_)[p0] if (p0) in (d_134_v54_) else d_133_v53_))) <= ((0) - ((-279 if True else p2)))
        elif True:
            d_135_v55_: _dafny.Map
            d_135_v55_ = _dafny.Map({p2: p0})
            d_136_v56_: C4
            nw19_ = C4()
            nw19_.ctor__(len(d_135_v55_))
            d_136_v56_ = nw19_
            d_137_v57_: _dafny.Array
            nw20_ = _dafny.Array(None, 15)
            nw20_[int(0)] = d_136_v56_
            nw20_[int(1)] = d_136_v56_
            nw20_[int(2)] = d_136_v56_
            nw20_[int(3)] = (d_136_v56_ if not(p0) else d_136_v56_)
            nw20_[int(4)] = d_136_v56_
            nw20_[int(5)] = d_136_v56_
            nw20_[int(6)] = d_136_v56_
            nw20_[int(7)] = d_136_v56_
            nw20_[int(8)] = d_136_v56_
            nw20_[int(9)] = d_136_v56_
            nw20_[int(10)] = d_136_v56_
            nw20_[int(11)] = d_136_v56_
            nw20_[int(12)] = d_136_v56_
            nw20_[int(13)] = (d_136_v56_ if p0 else d_136_v56_)
            nw20_[int(14)] = d_136_v56_
            d_137_v57_ = nw20_
            index13_ = default__.safeIndex(755, (d_137_v57_).length(0))
            (d_137_v57_)[index13_] = d_136_v56_
            d_138_v58_: _dafny.Seq
            d_138_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlsst"))
            d_139_v59_: _dafny.Map
            d_139_v59_ = _dafny.Map({p0: len(d_138_v58_)})
            d_140_v60_: _dafny.Array
            nw21_ = _dafny.Array(int(0), 13)
            d_140_v60_ = nw21_
            index14_ = default__.safeIndex(504, (d_140_v60_).length(0))
            (d_140_v60_)[index14_] = (0) - ((d_136_v56_).f9)
            d_141_v61_: _dafny.Array
            def lambda8_(d_142_p0_):
                def lambda9_(d_143_i10_):
                    return d_142_p0_

                return lambda9_

            init4_ = lambda8_(p0)
            nw22_ = _dafny.Array(None, 23)
            for i0_4_ in range(nw22_.length(0)):
                nw22_[i0_4_] = init4_(i0_4_)
            d_141_v61_ = nw22_
            d_144_v62_: D4
            d_144_v62_ = D4_DC13(p2, d_141_v61_)
            d_145_v63_: C4
            d_145_v63_ = d_136_v56_
            d_146_v64_: C4
            d_146_v64_ = (d_145_v63_)
            d_147_v65_: _dafny.Map
            d_147_v65_ = _dafny.Map({p0: _dafny.CodePoint('y')})
            d_148_v66_: D2
            d_148_v66_ = D2_DC6(p0)
            d_149_v67_: C2
            nw23_ = C2()
            nw23_.ctor__(p0)
            d_149_v67_ = nw23_
            d_150_v68_: _dafny.Set
            d_150_v68_ = _dafny.Set({d_149_v67_, d_149_v67_, d_149_v67_})
            index15_ = default__.safeIndex(755, (d_137_v57_).length(0))
            index16_ = default__.safeIndex(504, (d_140_v60_).length(0))
            rhs11_ = (d_145_v63_)
            rhs12_ = d_139_v59_
            rhs13_ = len(default__.fm8(p0, d_147_v65_, (d_148_v66_).cf6, globalState))
            rhs14_ = d_144_v62_
            rhs15_ = ((d_150_v68_).isdisjoint(d_150_v68_)) and ((d_149_v67_).f13)
            lhs11_ = d_137_v57_
            lhs12_ = default__.safeIndex(755, (d_137_v57_).length(0))
            lhs13_ = d_140_v60_
            lhs14_ = default__.safeIndex(504, (d_140_v60_).length(0))
            lhs15_ = globalState
            lhs11_[lhs12_] = rhs11_
            d_139_v59_ = rhs12_
            lhs13_[lhs14_] = rhs13_
            d_144_v62_ = rhs14_
            lhs15_.f2 = rhs15_
            d_151_v69_: C0
            nw24_ = C0()
            nw24_.ctor__(p0, d_141_v61_)
            d_151_v69_ = nw24_
            d_152_v70_: T0
            nw25_ = C1()
            nw25_.ctor__()
            d_152_v70_ = nw25_
            d_153_v71_: _dafny.Set
            d_153_v71_ = _dafny.Set({d_152_v70_, d_152_v70_, d_152_v70_, d_152_v70_})
            d_154_v72_: _dafny.Seq
            d_154_v72_ = _dafny.SeqWithoutIsStrInference([len(d_153_v71_), p2])
            index17_ = default__.safeIndex(504, (d_140_v60_).length(0))
            (d_140_v60_)[index17_] = (0) - (default__.safeDivisionInt(default__.safeModuloInt((d_136_v56_).f9, p2), len(d_154_v72_)))
            d_135_v55_ = d_135_v55_
            d_155_v74_: _dafny.Seq
            d_155_v74_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_151_v69_).f10, (d_149_v67_).f13])])
            d_156_v75_: _dafny.Map
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: _dafny.Seq
                for compr_16_ in (d_155_v74_).Elements:
                    d_157_v73_: _dafny.Seq = compr_16_
                    if (d_157_v73_) in (d_155_v74_):
                        coll16_[d_157_v73_] = len(d_154_v72_)
                return _dafny.Map(coll16_)
            d_156_v75_ = _dafny.Map({iife16_()
            : d_140_v60_})
            d_158_v76_: _dafny.Map
            d_158_v76_ = _dafny.Map({d_70_v4_: (d_140_v60_)[default__.safeIndex(504, (d_140_v60_).length(0))]})
            d_140_v60_ = ((d_156_v75_)[(d_158_v76_) | (d_158_v76_)] if ((d_158_v76_) | (d_158_v76_)) in (d_156_v75_) else d_140_v60_)
        d_159_v77_: _dafny.Array
        def lambda10_(d_160_p0_):
            def lambda11_(d_161_i11_):
                return default__.safeDivisionInt(d_161_i11_, len(_dafny.Set({d_160_p0_})))

            return lambda11_

        init5_ = lambda10_(p0)
        nw26_ = _dafny.Array(None, 20)
        for i0_5_ in range(nw26_.length(0)):
            nw26_[i0_5_] = init5_(i0_5_)
        d_159_v77_ = nw26_
        d_162_v78_: _dafny.Map
        d_162_v78_ = _dafny.Map({p0: d_159_v77_})
        (globalState).f5 = default__.safeModuloInt((p2) * (len(d_162_v78_)), default__.fm0(p2, default__.fm0(818, p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmpivcli"))), p0, globalState), p2, p0, globalState))
        d_163_v79_: C4
        nw27_ = C4()
        nw27_.ctor__((p2) - (p2))
        d_163_v79_ = nw27_
        (globalState).f0 = False

    @staticmethod
    def Main(noArgsParameter__):
        d_164_v0_: int
        d_164_v0_ = 292
        d_165_v1_: _dafny.Seq
        d_165_v1_ = _dafny.SeqWithoutIsStrInference([d_164_v0_, d_164_v0_])
        d_166_v2_: bool
        d_166_v2_ = False
        d_167_v3_: _dafny.Array
        nw28_ = _dafny.Array(None, 3)
        nw28_[int(0)] = d_164_v0_
        nw28_[int(1)] = (d_165_v1_)[default__.safeIndex(len(_dafny.Map({d_166_v2_: d_166_v2_})), len(d_165_v1_))]
        nw28_[int(2)] = d_164_v0_
        d_167_v3_ = nw28_
        d_168_v4_: _dafny.MultiSet
        d_168_v4_ = _dafny.MultiSet([d_167_v3_])
        d_169_v5_: _dafny.Seq
        d_169_v5_ = _dafny.SeqWithoutIsStrInference([True])
        d_170_v6_: _dafny.Map
        d_170_v6_ = _dafny.Map({len(d_169_v5_): d_166_v2_})
        d_171_v7_: _dafny.Map
        d_171_v7_ = _dafny.Map({d_164_v0_: _dafny.MultiSet([d_170_v6_])})
        d_172_v9_: D0
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: int
            for compr_17_ in _dafny.IntegerRange(388, 693):
                d_173_v8_: int = compr_17_
                if ((388) <= (d_173_v8_)) and ((d_173_v8_) < (693)):
                    coll17_[default__.safeDivisionInt(d_173_v8_, d_164_v0_)] = d_166_v2_
            return _dafny.Map(coll17_)
        d_172_v9_ = D0_DC0(iife17_()
)
        d_174_v10_: _dafny.MultiSet
        d_174_v10_ = _dafny.MultiSet([d_170_v6_, (d_172_v9_).cf0])
        d_175_v11_: _dafny.Set
        d_175_v11_ = _dafny.Set({d_166_v2_, d_166_v2_})
        d_176_globalState_: GlobalState
        nw29_ = GlobalState()
        nw29_.ctor__(True, -616, False, d_168_v4_, True, 663, (((d_171_v7_)[d_164_v0_] if (d_164_v0_) in (d_171_v7_) else d_174_v10_)) | (d_174_v10_), False, d_175_v11_)
        d_176_globalState_ = nw29_
        d_177_v12_: _dafny.Seq
        d_177_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiwqmuo"))
        (d_176_globalState_).f2 = (default__.safeModuloInt(d_164_v0_, (0) - (d_164_v0_))) < (default__.fm0(len(d_177_v12_), d_164_v0_, d_164_v0_, d_166_v2_, d_176_globalState_))
        d_178_i0_: int
        d_178_i0_ = 0
        with _dafny.label("0"):
            while d_166_v2_:
                with _dafny.c_label("0"):
                    if (d_178_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_178_i0_ = (d_178_i0_) + (1)
                    (d_176_globalState_).f0 = default__.fm1((d_164_v0_) <= (311), d_176_globalState_)
                    index18_ = default__.safeIndex(871, (d_167_v3_).length(0))
                    (d_167_v3_)[index18_] = d_164_v0_
                    index19_ = default__.safeIndex(871, (d_167_v3_).length(0))
                    rhs16_ = d_164_v0_
                    rhs17_ = 618
                    lhs16_ = d_167_v3_
                    lhs17_ = default__.safeIndex(871, (d_167_v3_).length(0))
                    lhs16_[lhs17_] = rhs16_
                    d_164_v0_ = rhs17_
                    d_179_v13_: str
                    d_179_v13_ = _dafny.CodePoint('y')
                    d_180_v14_: _dafny.Map
                    d_180_v14_ = _dafny.Map({d_166_v2_: len(d_177_v12_)})
                    default__.m0(not((d_166_v2_) == (d_166_v2_)), d_179_v13_, (0) - (((d_180_v14_)[d_166_v2_] if (d_166_v2_) in (d_180_v14_) else (d_167_v3_)[default__.safeIndex(871, (d_167_v3_).length(0))])), d_180_v14_, d_176_globalState_)
                    d_181_v15_: _dafny.Map
                    d_181_v15_ = _dafny.Map({(0) - (d_164_v0_): d_165_v1_})
                    d_181_v15_ = d_181_v15_
                    pass
            pass
        d_182_v17_: _dafny.Map
        d_182_v17_ = _dafny.Map({d_166_v2_: len(d_177_v12_)})
        def iife18_():
            coll18_ = _dafny.Set()
            compr_18_: str
            for compr_18_ in (d_177_v12_).Elements:
                d_183_v16_: str = compr_18_
                if (d_183_v16_) in (d_177_v12_):
                    coll18_ = coll18_.union(_dafny.Set([d_183_v16_]))
            return _dafny.Set(coll18_)
        default__.m0(True, (d_177_v12_)[default__.safeIndex(default__.fm0(len(iife18_()
        ), (0) - (d_164_v0_), d_164_v0_, d_166_v2_, d_176_globalState_), len(d_177_v12_))], (default__.fm0(d_164_v0_, d_164_v0_, ((d_182_v17_)[d_166_v2_] if (d_166_v2_) in (d_182_v17_) else len(d_177_v12_)), d_166_v2_, d_176_globalState_)) - (d_164_v0_), default__.fm2(d_166_v2_, d_166_v2_, d_166_v2_, d_182_v17_, d_176_globalState_), d_176_globalState_)
        d_184_v18_: str
        d_184_v18_ = _dafny.CodePoint('n')
        default__.m0(d_166_v2_, d_184_v18_, default__.fm0(d_164_v0_, d_164_v0_, d_164_v0_, d_166_v2_, d_176_globalState_), _dafny.Map({d_166_v2_: default__.fm0(d_164_v0_, d_164_v0_, len(d_177_v12_), True, d_176_globalState_)}), d_176_globalState_)
        hi0_ = -752
        for d_185_i1_ in range(d_164_v0_, hi0_):
            d_164_v0_ = d_185_i1_
            (d_176_globalState_).f0 = d_166_v2_
            default__.m0(d_166_v2_, d_184_v18_, d_185_i1_, d_182_v17_, d_176_globalState_)
            d_186_v19_: _dafny.MultiSet
            d_186_v19_ = _dafny.MultiSet([d_185_i1_, d_164_v0_, d_185_i1_])
            d_186_v19_ = d_186_v19_
        d_187_v20_: D0
        d_187_v20_ = D0_DC1(d_166_v2_)
        d_188_v21_: _dafny.Map
        d_188_v21_ = _dafny.Map({d_166_v2_: d_187_v20_})
        d_188_v21_ = (d_188_v21_).set(d_166_v2_, D0_DC1(d_166_v2_))
        source4_ = d_187_v20_
        if source4_.is_DC1:
            d_189___mcc_h0_ = source4_.cf1
            d_190_cf1_ = d_189___mcc_h0_
            d_191_v22_: _dafny.Map
            d_191_v22_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrvcjj")): d_164_v0_})
            d_192_v23_: _dafny.MultiSet
            d_192_v23_ = _dafny.MultiSet([len(d_191_v22_)])
            d_193_v25_: D0
            d_193_v25_ = D0_DC0(d_170_v6_)
            d_194_v26_: D0
            d_194_v26_ = D0_DC2(d_193_v25_)
            d_195_v27_: _dafny.Map
            d_195_v27_ = _dafny.Map({d_194_v26_: d_177_v12_})
            def iife19_():
                coll19_ = _dafny.Map()
                compr_19_: D0
                for compr_19_ in (d_195_v27_).keys.Elements:
                    d_196_v24_: D0 = compr_19_
                    if (d_196_v24_) in (d_195_v27_):
                        coll19_[d_196_v24_] = (0) - (d_164_v0_)
                return _dafny.Map(coll19_)
            def iife20_():
                coll20_ = _dafny.Map()
                compr_20_: D0
                for compr_20_ in (d_195_v27_).keys.Elements:
                    d_197_v24_: D0 = compr_20_
                    if (d_197_v24_) in (d_195_v27_):
                        coll20_[d_197_v24_] = (0) - (d_164_v0_)
                return _dafny.Map(coll20_)
            rhs18_ = ((d_192_v23_)[(0) - ((0) - (len(iife19_()
            )))] if ((0) - ((0) - (len(iife20_()
            )))) in (d_192_v23_) else default__.fm0(d_164_v0_, d_164_v0_, d_164_v0_, d_190_cf1_, d_176_globalState_))
            rhs19_ = default__.safeDivisionInt(550, default__.safeDivisionInt(d_164_v0_, d_164_v0_))
            lhs18_ = d_176_globalState_
            lhs18_.f5 = rhs18_
            d_164_v0_ = rhs19_
            d_198_v28_: _dafny.Seq
            d_198_v28_ = _dafny.SeqWithoutIsStrInference([d_175_v11_])
            (d_176_globalState_).f0 = not(((d_190_cf1_) or (default__.fm1(d_190_cf1_, d_176_globalState_))) in ((d_175_v11_).intersection((d_198_v28_)[default__.safeIndex(d_164_v0_, len(d_198_v28_))])))
            d_177_v12_ = d_177_v12_
            (d_176_globalState_).f2 = d_166_v2_
        elif source4_.is_DC0:
            d_199___mcc_h1_ = source4_.cf0
            d_200_cf0_ = d_199___mcc_h1_
            default__.m0(not(d_166_v2_), d_184_v18_, d_164_v0_, d_182_v17_, d_176_globalState_)
            default__.m0((not(d_166_v2_)) == (d_166_v2_), d_184_v18_, d_164_v0_, d_182_v17_, d_176_globalState_)
            d_177_v12_ = (d_177_v12_) + (d_177_v12_)
            (d_176_globalState_).f5 = d_164_v0_
        elif True:
            d_201___mcc_h2_ = source4_.cf2
            d_202_cf2_ = d_201___mcc_h2_
            default__.m0(d_166_v2_, d_184_v18_, (0) - ((d_164_v0_) - (d_164_v0_)), d_182_v17_, d_176_globalState_)
            pat_let_tv0_ = d_166_v2_
            d_203_v29_: _dafny.MultiSet
            def iife21_(_pat_let0_0):
                def iife22_(d_204_dt__update__tmp_h0_):
                    def iife23_(_pat_let1_0):
                        def iife24_(d_205_dt__update_hcf1_h0_):
                            return D0_DC1(d_205_dt__update_hcf1_h0_)
                        return iife24_(_pat_let1_0)
                    return iife23_(pat_let_tv0_)
                return iife22_(_pat_let0_0)
            d_203_v29_ = _dafny.MultiSet([d_187_v20_, d_187_v20_, iife21_(d_187_v20_)])
            d_206_v30_: _dafny.Set
            d_206_v30_ = _dafny.Set({d_167_v3_})
            (d_176_globalState_).f2 = ((_dafny.MultiSet([D0_DC1((d_169_v5_)[default__.safeIndex(len(d_182_v17_), len(d_169_v5_))])]) if d_166_v2_ else d_203_v29_)) == ((((d_203_v29_).set(d_187_v20_, default__.abs(d_164_v0_))).set(default__.fm3((d_187_v20_).cf1, d_164_v0_, len(d_206_v30_), d_164_v0_, d_176_globalState_), default__.abs(d_164_v0_)) if False else d_203_v29_))
            d_170_v6_ = (d_170_v6_).set(d_164_v0_, ((0) - (d_164_v0_)) < (d_164_v0_))
            d_207_v31_: C1
            nw30_ = C1()
            nw30_.ctor__()
            d_207_v31_ = nw30_
        if False:
            (d_176_globalState_).f2 = d_166_v2_
            (d_176_globalState_).f0 = not(d_166_v2_)
            if d_166_v2_:
                d_208_v32_: _dafny.Map
                d_208_v32_ = _dafny.Map({d_164_v0_: d_167_v3_})
                d_209_v33_: _dafny.MultiSet
                d_209_v33_ = _dafny.MultiSet([d_166_v2_, default__.fm1(d_166_v2_, d_176_globalState_), d_166_v2_])
                d_208_v32_ = (d_208_v32_).set((d_209_v33_).cardinality, d_167_v3_)
                d_165_v1_ = d_165_v1_
                d_210_v34_: D4
                d_210_v34_ = D4_DC12(len(d_177_v12_), d_164_v0_, d_177_v12_, 350, len(d_169_v5_))
                d_182_v17_ = (d_182_v17_).set((d_177_v12_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvye"))), (d_210_v34_).cf17)
                d_211_v35_: _dafny.Array
                nw31_ = _dafny.Array(False, 24)
                d_211_v35_ = nw31_
                index20_ = default__.safeIndex(765, (d_211_v35_).length(0))
                (d_211_v35_)[index20_] = d_166_v2_
                index21_ = default__.safeIndex(765, (d_211_v35_).length(0))
                (d_211_v35_)[index21_] = d_166_v2_
                (d_176_globalState_).f2 = d_166_v2_
            elif True:
                (d_176_globalState_).f2 = d_166_v2_
                d_164_v0_ = len((d_177_v12_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "poavktqf"))))
                d_212_v36_: _dafny.Array
                def lambda12_(d_213_v6_):
                    def lambda13_(d_214_i2_):
                        return d_213_v6_

                    return lambda13_

                init6_ = lambda12_(d_170_v6_)
                nw32_ = _dafny.Array(None, 3)
                for i0_6_ in range(nw32_.length(0)):
                    nw32_[i0_6_] = init6_(i0_6_)
                d_212_v36_ = nw32_
                index22_ = default__.safeIndex(47, (d_212_v36_).length(0))
                (d_212_v36_)[index22_] = _dafny.Map({d_164_v0_: d_166_v2_})
                d_215_v37_: D8
                d_215_v37_ = D8_DC21(d_166_v2_, d_166_v2_, ((d_169_v5_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcd"))), len(d_169_v5_)), d_166_v2_)).set(default__.safeIndex(-577, len((d_169_v5_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcd"))), len(d_169_v5_)), d_166_v2_))), d_166_v2_), d_166_v2_, d_164_v0_)
                index23_ = default__.safeIndex(47, (d_212_v36_).length(0))
                (d_212_v36_)[index23_] = (d_170_v6_).set(d_164_v0_, (d_215_v37_).cf27)
                d_169_v5_ = (((d_169_v5_).set(default__.safeIndex(d_164_v0_, len(d_169_v5_)), d_166_v2_)) + (d_169_v5_)) + (d_169_v5_)
                d_216_v38_: T1
                nw33_ = C4()
                nw33_.ctor__(d_164_v0_)
                d_216_v38_ = nw33_
                d_217_v39_: _dafny.Seq
                d_217_v39_ = _dafny.SeqWithoutIsStrInference([d_216_v38_, d_216_v38_])
                default__.m0(d_166_v2_, d_184_v18_, default__.safeDivisionInt(d_164_v0_, default__.fm0(d_164_v0_, d_164_v0_, d_164_v0_, d_166_v2_, d_176_globalState_)), (d_182_v17_).set(d_166_v2_, len(d_217_v39_)), d_176_globalState_)
            d_218_v40_: _dafny.Map
            d_218_v40_ = _dafny.Map({d_164_v0_: D8_DC21(d_166_v2_, d_166_v2_, d_169_v5_, d_166_v2_, d_164_v0_)})
            d_218_v40_ = default__.fm33(d_166_v2_, d_184_v18_, d_176_globalState_)
            d_219_v41_: _dafny.Array
            def lambda14_(d_220_v18_):
                def lambda15_(d_221_i3_):
                    return d_220_v18_

                return lambda15_

            init7_ = lambda14_(d_184_v18_)
            nw34_ = _dafny.Array(None, 25)
            for i0_7_ in range(nw34_.length(0)):
                nw34_[i0_7_] = init7_(i0_7_)
            d_219_v41_ = nw34_
            index24_ = default__.safeIndex(349, (d_219_v41_).length(0))
            (d_219_v41_)[index24_] = d_184_v18_
            index25_ = default__.safeIndex(349, (d_219_v41_).length(0))
            (d_219_v41_)[index25_] = d_184_v18_
        elif True:
            d_222_v42_: _dafny.Map
            d_222_v42_ = _dafny.Map({d_164_v0_: d_164_v0_})
            d_222_v42_ = (d_222_v42_).set(d_164_v0_, d_164_v0_)
            d_164_v0_ = (0) - (len(d_182_v17_))
            d_223_v43_: C1
            nw35_ = C1()
            nw35_.ctor__()
            d_223_v43_ = nw35_
            d_223_v43_ = d_223_v43_
            (d_176_globalState_).f2 = d_166_v2_
            (d_176_globalState_).f2 = d_166_v2_
        (d_176_globalState_).f2 = (d_164_v0_) < (d_164_v0_)
        d_224_v44_: _dafny.Map
        d_224_v44_ = _dafny.Map({d_177_v12_: d_164_v0_})
        d_224_v44_ = (d_224_v44_).set(d_177_v12_, default__.safeDivisionInt(d_164_v0_, d_164_v0_))
        default__.m0(d_166_v2_, d_184_v18_, d_164_v0_, (d_182_v17_) | (d_182_v17_), d_176_globalState_)
        d_225_v45_: _dafny.MultiSet
        d_225_v45_ = _dafny.MultiSet([d_166_v2_])
        (d_176_globalState_).f2 = (d_225_v45_).isdisjoint(d_225_v45_)
        if d_166_v2_:
            d_226_v46_: _dafny.MultiSet
            d_226_v46_ = _dafny.MultiSet([d_184_v18_, d_184_v18_])
            if (default__.fm0(((d_226_v46_)[d_184_v18_] if (d_184_v18_) in (d_226_v46_) else d_164_v0_), len(d_165_v1_), d_164_v0_, d_166_v2_, d_176_globalState_)) <= (d_164_v0_):
                default__.m0(d_166_v2_, d_184_v18_, d_164_v0_, d_182_v17_, d_176_globalState_)
                d_227_v47_: T1
                nw36_ = C3()
                nw36_.ctor__(997)
                d_227_v47_ = nw36_
                d_228_v48_: _dafny.Set
                d_228_v48_ = _dafny.Set({d_227_v47_})
                d_229_v49_: D8
                d_229_v49_ = D8_DC20((d_228_v48_).intersection(d_228_v48_))
                d_229_v49_ = (d_229_v49_ if (d_164_v0_) < (d_164_v0_) else d_229_v49_)
                d_230_v50_: _dafny.Array
                nw37_ = _dafny.Array(None, 25)
                nw37_[int(0)] = d_166_v2_
                nw37_[int(1)] = d_166_v2_
                nw37_[int(2)] = d_166_v2_
                nw37_[int(3)] = d_166_v2_
                nw37_[int(4)] = d_166_v2_
                nw37_[int(5)] = d_166_v2_
                nw37_[int(6)] = False
                nw37_[int(7)] = d_166_v2_
                nw37_[int(8)] = True
                nw37_[int(9)] = True
                nw37_[int(10)] = d_166_v2_
                nw37_[int(11)] = d_166_v2_
                nw37_[int(12)] = True
                nw37_[int(13)] = d_166_v2_
                nw37_[int(14)] = d_166_v2_
                nw37_[int(15)] = d_166_v2_
                nw37_[int(16)] = (d_227_v47_).fm4(d_172_v9_, D0_DC0(_dafny.Map({793: d_166_v2_})), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_166_v2_])), d_176_globalState_)
                nw37_[int(17)] = d_166_v2_
                nw37_[int(18)] = d_166_v2_
                nw37_[int(19)] = d_166_v2_
                nw37_[int(20)] = d_166_v2_
                nw37_[int(21)] = d_166_v2_
                nw37_[int(22)] = d_166_v2_
                nw37_[int(23)] = d_166_v2_
                nw37_[int(24)] = d_166_v2_
                d_230_v50_ = nw37_
                d_231_v51_: _dafny.Map
                d_231_v51_ = _dafny.Map({d_165_v1_: d_230_v50_})
                d_232_v52_: _dafny.Map
                d_232_v52_ = _dafny.Map({d_231_v51_: (D9_DC24(d_166_v2_, False, d_166_v2_)).cf34})
                d_232_v52_ = (d_232_v52_).set(d_231_v51_, (d_166_v2_) in (d_175_v11_))
                (d_176_globalState_).f0 = True
                (d_176_globalState_).f2 = default__.fm1(((_dafny.SeqWithoutIsStrInference([d_184_v18_ for d_233_i4_ in range(default__.abs(671))])).set(default__.safeIndex(d_164_v0_, len(_dafny.SeqWithoutIsStrInference([d_184_v18_ for d_234_i4_ in range(default__.abs(671))]))), d_184_v18_)) < (d_177_v12_), d_176_globalState_)
            elif True:
                rhs20_ = default__.fm0(d_164_v0_, d_164_v0_, (0) - (d_164_v0_), not(d_166_v2_), d_176_globalState_)
                rhs21_ = ((d_226_v46_ if not(d_166_v2_) else d_226_v46_)).issubset(_dafny.MultiSet([d_184_v18_]))
                lhs19_ = d_176_globalState_
                d_164_v0_ = rhs20_
                lhs19_.f2 = rhs21_
                d_235_v53_: _dafny.Map
                d_235_v53_ = _dafny.Map({d_166_v2_: d_166_v2_})
                d_236_v54_: D2
                d_236_v54_ = D2_DC7(len(d_177_v12_), _dafny.Map({d_235_v53_: d_164_v0_}), d_164_v0_)
                d_237_v55_: _dafny.MultiSet
                d_237_v55_ = _dafny.MultiSet([d_236_v54_])
                d_238_v56_: _dafny.Map
                d_238_v56_ = _dafny.Map({((d_237_v55_)[d_236_v54_] if (d_236_v54_) in (d_237_v55_) else d_164_v0_): d_164_v0_})
                d_239_v57_: T0
                nw38_ = C1()
                nw38_.ctor__()
                d_239_v57_ = nw38_
                d_240_v58_: _dafny.Set
                d_240_v58_ = _dafny.Set({d_239_v57_})
                d_238_v56_ = (d_238_v56_).set(default__.safeModuloInt(d_164_v0_, d_164_v0_), len((d_240_v58_).intersection(d_240_v58_)))
                d_167_v3_ = d_167_v3_
                (d_176_globalState_).f5 = d_164_v0_
                d_241_v59_: T1
                nw39_ = C3()
                nw39_.ctor__(d_164_v0_)
                d_241_v59_ = nw39_
                d_242_v60_: D2
                d_242_v60_ = D2_DC5(d_241_v59_)
                d_243_v61_: _dafny.Map
                d_243_v61_ = _dafny.Map({_dafny.MultiSet(d_169_v5_): (d_242_v60_).cf5})
                d_244_v62_: D9
                d_244_v62_ = D9_DC24(d_166_v2_, d_166_v2_, d_166_v2_)
                d_245_v63_: _dafny.Array
                nw40_ = _dafny.Array(None, 15)
                nw40_[int(0)] = d_241_v59_
                nw40_[int(1)] = (d_241_v59_ if d_166_v2_ else d_241_v59_)
                nw40_[int(2)] = d_241_v59_
                nw40_[int(3)] = ((d_243_v61_)[d_225_v45_] if (d_225_v45_) in (d_243_v61_) else d_241_v59_)
                nw40_[int(4)] = d_241_v59_
                nw40_[int(5)] = d_241_v59_
                nw40_[int(6)] = d_241_v59_
                nw40_[int(7)] = (d_241_v59_ if (d_244_v62_).cf35 else d_241_v59_)
                nw40_[int(8)] = d_241_v59_
                nw40_[int(9)] = d_241_v59_
                nw40_[int(10)] = d_241_v59_
                nw40_[int(11)] = d_241_v59_
                nw40_[int(12)] = d_241_v59_
                nw40_[int(13)] = d_241_v59_
                nw40_[int(14)] = d_241_v59_
                d_245_v63_ = nw40_
                index26_ = default__.safeIndex(122, (d_245_v63_).length(0))
                (d_245_v63_)[index26_] = d_241_v59_
                d_246_v64_: D13
                d_246_v64_ = D13_DC38(d_182_v17_)
                d_247_v65_: _dafny.MultiSet
                d_247_v65_ = _dafny.MultiSet([d_246_v64_, d_246_v64_])
                d_248_v66_: _dafny.Array
                nw41_ = _dafny.Array(None, 3)
                nw41_[int(0)] = not(d_166_v2_)
                nw41_[int(1)] = (d_247_v65_).ispropersubset(d_247_v65_)
                nw41_[int(2)] = d_166_v2_
                d_248_v66_ = nw41_
                index27_ = default__.safeIndex(556, (d_248_v66_).length(0))
                (d_248_v66_)[index27_] = d_166_v2_
                d_249_v67_: C3
                nw42_ = C3()
                nw42_.ctor__(d_164_v0_)
                d_249_v67_ = nw42_
                index28_ = default__.safeIndex(122, (d_245_v63_).length(0))
                index29_ = default__.safeIndex(556, (d_248_v66_).length(0))
                rhs22_ = d_241_v59_
                rhs23_ = d_166_v2_
                rhs24_ = d_249_v67_
                lhs20_ = d_245_v63_
                lhs21_ = default__.safeIndex(122, (d_245_v63_).length(0))
                lhs22_ = d_248_v66_
                lhs23_ = default__.safeIndex(556, (d_248_v66_).length(0))
                lhs20_[lhs21_] = rhs22_
                lhs22_[lhs23_] = rhs23_
                d_249_v67_ = rhs24_
            d_250_v68_: _dafny.Array
            nw43_ = _dafny.Array(None, 5)
            d_250_v68_ = nw43_
            d_251_v69_: T0
            nw44_ = C2()
            nw44_.ctor__(d_166_v2_)
            d_251_v69_ = nw44_
            index30_ = default__.safeIndex(929, (d_250_v68_).length(0))
            (d_250_v68_)[index30_] = d_251_v69_
            index31_ = default__.safeIndex(929, (d_250_v68_).length(0))
            rhs25_ = (d_177_v12_ if d_166_v2_ else _dafny.SeqWithoutIsStrInference([(D10_DC27(d_166_v2_, d_164_v0_, d_184_v18_, d_225_v45_)).cf40 for d_252_i5_ in range(default__.abs(938))]))
            rhs26_ = d_251_v69_
            lhs24_ = d_250_v68_
            lhs25_ = default__.safeIndex(929, (d_250_v68_).length(0))
            d_177_v12_ = rhs25_
            lhs24_[lhs25_] = rhs26_
            (d_176_globalState_).f0 = (d_166_v2_ if (d_164_v0_) <= (d_164_v0_) else d_166_v2_)
            (d_176_globalState_).f5 = (0) - (d_164_v0_)
            (d_176_globalState_).f5 = len((default__.fm34(d_164_v0_, d_164_v0_, not(d_166_v2_), d_176_globalState_)).set(d_175_v11_, -582))
        elif True:
            default__.m0(not (d_166_v2_) or (d_166_v2_), d_184_v18_, d_164_v0_, d_182_v17_, d_176_globalState_)
            d_166_v2_ = (d_164_v0_) not in (d_165_v1_)
            d_253_v70_: _dafny.Array
            nw45_ = _dafny.Array(None, 28)
            nw45_[int(0)] = False
            nw45_[int(1)] = d_166_v2_
            nw45_[int(2)] = d_166_v2_
            nw45_[int(3)] = d_166_v2_
            nw45_[int(4)] = d_166_v2_
            nw45_[int(5)] = d_166_v2_
            nw45_[int(6)] = d_166_v2_
            nw45_[int(7)] = d_166_v2_
            nw45_[int(8)] = d_166_v2_
            nw45_[int(9)] = d_166_v2_
            nw45_[int(10)] = d_166_v2_
            nw45_[int(11)] = d_166_v2_
            nw45_[int(12)] = d_166_v2_
            nw45_[int(13)] = not(d_166_v2_)
            nw45_[int(14)] = default__.fm1(not(d_166_v2_), d_176_globalState_)
            nw45_[int(15)] = d_166_v2_
            nw45_[int(16)] = d_166_v2_
            nw45_[int(17)] = d_166_v2_
            nw45_[int(18)] = d_166_v2_
            nw45_[int(19)] = d_166_v2_
            nw45_[int(20)] = d_166_v2_
            nw45_[int(21)] = d_166_v2_
            nw45_[int(22)] = d_166_v2_
            nw45_[int(23)] = not(d_166_v2_)
            nw45_[int(24)] = d_166_v2_
            nw45_[int(25)] = False
            nw45_[int(26)] = True
            nw45_[int(27)] = d_166_v2_
            d_253_v70_ = nw45_
            d_254_v71_: C0
            nw46_ = C0()
            nw46_.ctor__(True, d_253_v70_)
            d_254_v71_ = nw46_
            d_255_v72_: C2
            nw47_ = C2()
            nw47_.ctor__((d_254_v71_).f10)
            d_255_v72_ = nw47_
            d_256_v73_: _dafny.Array
            nw48_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_256_v73_ = nw48_
            d_257_v74_: bool
            d_258_v75_: int
            d_259_v76_: bool
            d_260_v77_: bool
            out0_: bool
            out1_: int
            out2_: bool
            out3_: bool
            out0_, out1_, out2_, out3_ = (d_255_v72_).m1(d_164_v0_, d_256_v73_, ((_dafny.MultiSet(d_169_v5_)) | (_dafny.MultiSet(d_169_v5_))).cardinality, d_164_v0_, d_176_globalState_)
            d_257_v74_ = out0_
            d_258_v75_ = out1_
            d_259_v76_ = out2_
            d_260_v77_ = out3_
        d_261_v78_: _dafny.Array
        nw49_ = _dafny.Array(False, 10)
        d_261_v78_ = nw49_
        d_262_v79_: C0
        nw50_ = C0()
        nw50_.ctor__(d_166_v2_, d_261_v78_)
        d_262_v79_ = nw50_
        d_263_v80_: _dafny.Set
        d_263_v80_ = _dafny.Set({d_262_v79_})
        d_264_v81_: _dafny.Map
        d_264_v81_ = _dafny.Map({d_166_v2_: d_263_v80_})
        d_265_v82_: _dafny.Map
        d_265_v82_ = _dafny.Map({len(((d_264_v81_)[(d_262_v79_).fm4(d_172_v9_, D0_DC0(d_170_v6_), d_225_v45_, d_176_globalState_)] if ((d_262_v79_).fm4(d_172_v9_, D0_DC0(d_170_v6_), d_225_v45_, d_176_globalState_)) in (d_264_v81_) else d_263_v80_)): d_164_v0_})
        d_266_v83_: T1
        nw51_ = C0()
        nw51_.ctor__((d_262_v79_).f10, (d_262_v79_).f11)
        d_266_v83_ = nw51_
        d_267_v84_: _dafny.Map
        d_267_v84_ = _dafny.Map({d_266_v83_: d_164_v0_})
        d_268_v85_: _dafny.Map
        d_268_v85_ = _dafny.Map({d_166_v2_: d_267_v84_})
        d_265_v82_ = (d_265_v82_).set((d_164_v0_ if (d_262_v79_).f10 else default__.fm0(d_164_v0_, d_164_v0_, len(d_268_v85_), (d_262_v79_).f10, d_176_globalState_)), d_164_v0_)
        d_269_v86_: _dafny.Array
        def lambda16_(d_270_v2_, d_271_v0_):
            def lambda17_(d_272_i6_):
                return D13_DC38(_dafny.Map({d_270_v2_: d_271_v0_}))

            return lambda17_

        init8_ = lambda16_(d_166_v2_, d_164_v0_)
        nw52_ = _dafny.Array(None, 10)
        for i0_8_ in range(nw52_.length(0)):
            nw52_[i0_8_] = init8_(i0_8_)
        d_269_v86_ = nw52_
        d_273_v87_: _dafny.Set
        d_273_v87_ = _dafny.Set({d_269_v86_})
        if (d_273_v87_).isdisjoint(d_273_v87_):
            d_175_v11_ = default__.fm17(not((d_262_v79_).f10), (0) - (d_164_v0_), (d_262_v79_).f10, d_164_v0_, d_176_globalState_)
            d_274_v88_: C4
            nw53_ = C4()
            nw53_.ctor__(d_164_v0_)
            d_274_v88_ = nw53_
            d_275_v89_: _dafny.Seq
            d_275_v89_ = _dafny.SeqWithoutIsStrInference([d_274_v88_])
            d_276_v90_: _dafny.Set
            d_276_v90_ = _dafny.Set({len(d_275_v89_), (d_274_v88_).f9, (d_274_v88_).f9})
            d_277_v91_: _dafny.Map
            d_277_v91_ = _dafny.Map({d_262_v79_: len(d_276_v90_)})
            d_277_v91_ = (d_277_v91_).set(d_262_v79_, (0) - (d_164_v0_))
            nw54_ = _dafny.Array(False, 13)
            d_261_v78_ = nw54_
            d_278_v92_: D3
            d_278_v92_ = D3_DC8(d_177_v12_)
            (d_176_globalState_).f2 = ((d_278_v92_).cf10) <= ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_279_i7_ in range(default__.abs(956))])) + (d_177_v12_))
            d_280_v93_: C1
            nw55_ = C1()
            nw55_.ctor__()
            d_280_v93_ = nw55_
            d_281_v94_: _dafny.Map
            d_281_v94_ = _dafny.Map({d_280_v93_: d_166_v2_})
            d_164_v0_ = (len(d_281_v94_)) * ((d_274_v88_).f9)
        elif True:
            (d_176_globalState_).f2 = (d_164_v0_) != (len((d_177_v12_).set(default__.safeIndex(d_164_v0_, len(d_177_v12_)), d_184_v18_)))
            d_282_v95_: C4
            nw56_ = C4()
            nw56_.ctor__(d_164_v0_)
            d_282_v95_ = nw56_
            d_166_v2_ = d_166_v2_
            d_283_v96_: C1
            nw57_ = C1()
            nw57_.ctor__()
            d_283_v96_ = nw57_
            (d_176_globalState_).f5 = (d_164_v0_) + (d_164_v0_)
        d_284_v97_: C1
        nw58_ = C1()
        nw58_.ctor__()
        d_284_v97_ = nw58_
        d_285_v98_: _dafny.Seq
        d_285_v98_ = _dafny.SeqWithoutIsStrInference([d_284_v97_, d_284_v97_])
        d_284_v97_ = (d_285_v98_)[default__.safeIndex(d_164_v0_, len(d_285_v98_))]
        _dafny.print(_dafny.string_of(d_164_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v1_) == (_dafny.SeqWithoutIsStrInference([292, 292]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_166_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_167_v3_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_167_v3_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_167_v3_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v4_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_v5_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_v6_) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v7_) == (_dafny.Map({292: _dafny.MultiSet([_dafny.Map({1: False})])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v9_).cf0) == (_dafny.Map({1: False, 2: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_174_v10_) == (_dafny.MultiSet([_dafny.Map({1: False}), _dafny.Map({1: False, 2: False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v11_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_176_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_176_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_globalState_.f3).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_176_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_176_globalState_).f6) == (_dafny.MultiSet([_dafny.Map({1: False}), _dafny.Map({1: False}), _dafny.Map({1: False, 2: False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_globalState_.f8) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_177_v12_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_178_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_182_v17_) == (_dafny.Map({False: 7}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_184_v18_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v20_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v21_) == (_dafny.Map({False: D0_DC1(False)}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_224_v44_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiwqmuo")): 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_225_v45_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v79_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_263_v80_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_264_v81_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_265_v82_) == (_dafny.Map({1: -1, -1: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_267_v84_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_268_v85_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_269_v86_)[0]).cf55) == (_dafny.Map({True: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_269_v86_)[1]).cf55) == (_dafny.Map({True: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_269_v86_)[2]).cf55) == (_dafny.Map({True: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_269_v86_)[3]).cf55) == (_dafny.Map({True: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_269_v86_)[4]).cf55) == (_dafny.Map({True: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_269_v86_)[5]).cf55) == (_dafny.Map({True: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_269_v86_)[6]).cf55) == (_dafny.Map({True: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_269_v86_)[7]).cf55) == (_dafny.Map({True: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_269_v86_)[8]).cf55) == (_dafny.Map({True: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_269_v86_)[9]).cf55) == (_dafny.Map({True: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_273_v87_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_285_v98_)))
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
        return lambda: D1_DC4(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC6(D2, NamedTuple('DC6', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC9(D3, NamedTuple('DC9', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({self.cf10.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC12(int(0), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC12(D4, NamedTuple('DC12', [('cf14', Any), ('cf15', Any), ('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {self.cf16.VerbatimString(True)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC15(D5, NamedTuple('DC15', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC16(D6, NamedTuple('DC16', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC18()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)

class D7_DC18(D7, NamedTuple('DC18', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC21(False, False, _dafny.Seq({}), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC21(D8, NamedTuple('DC21', [('cf26', Any), ('cf27', Any), ('cf28', Any), ('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC24(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)

class D9_DC24(D9, NamedTuple('DC24', [('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC23(D9, NamedTuple('DC23', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC27(False, int(0), _dafny.CodePoint('D'), _dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)

class D10_DC27(D10, NamedTuple('DC27', [('cf38', Any), ('cf39', Any), ('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC26(D10, NamedTuple('DC26', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC30()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)

class D11_DC30(D11, NamedTuple('DC30', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC31(D11, NamedTuple('DC31', [('cf44', Any), ('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf44 == __o.cf44 and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC29(D11, NamedTuple('DC29', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC32(D11, NamedTuple('DC32', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC34(D0.default()(), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D12_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D12_DC36)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D12_DC37)

class D12_DC34(D12, NamedTuple('DC34', [('cf48', Any), ('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC35(D12, NamedTuple('DC35', [])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35)
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC36(D12, NamedTuple('DC36', [('cf51', Any), ('cf52', Any), ('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC33(D12, NamedTuple('DC33', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC37(D12, NamedTuple('DC37', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC37({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC37) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC39()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D13_DC39)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D13_DC40)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D13_DC38)

class D13_DC39(D13, NamedTuple('DC39', [])):
    def __dafnystr__(self) -> str:
        return f'D13.DC39'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC39)
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC40(D13, NamedTuple('DC40', [('cf56', Any), ('cf57', Any), ('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC40({_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC40) and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC38(D13, NamedTuple('DC38', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC38({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC38) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D14_DC41)

class D14_DC41(D14, NamedTuple('DC41', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC41({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC41) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, p1, p2, p3, globalState):
        pass


class T1:
    pass

class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self.f2: bool = False
        self.f3: _dafny.MultiSet = _dafny.MultiSet({})
        self.f5: int = int(0)
        self.f8: _dafny.Set = _dafny.Set({})
        self._f1: int = int(0)
        self._f4: bool = False
        self._f6: _dafny.MultiSet = _dafny.MultiSet({})
        self._f7: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8):
        (self).f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8

    @property
    def f1(self):
        return self._f1
    @property
    def f4(self):
        return self._f4
    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7

class C0(T1, T0):
    def  __init__(self):
        self._f10: bool = False
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f10, f11):
        (self)._f10 = f10
        (self)._f11 = f11

    def fm4(self, p0, p1, p2, globalState):
        return (self).f10

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        r3: bool = False
        (globalState).f2 = (self).f10
        d_286_v0_: _dafny.Seq
        d_286_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "asdhidii"))
        d_287_i0_: int
        d_287_i0_ = 0
        with _dafny.label("1"):
            while (d_286_v0_) <= ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqahgg"))) + (d_286_v0_)):
                with _dafny.c_label("1"):
                    if (d_287_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_287_i0_ = (d_287_i0_) + (1)
                    d_288_v1_: D0
                    d_288_v1_ = D0_DC1((self).f10)
                    d_289_v2_: _dafny.Map
                    d_289_v2_ = _dafny.Map({((d_288_v1_).cf1) == ((self).f10): (self).f10})
                    d_289_v2_ = (d_289_v2_).set((self).f10, (self).f10)
                    d_290_v3_: _dafny.Array
                    nw59_ = _dafny.Array(int(0), 27)
                    d_290_v3_ = nw59_
                    index32_ = default__.safeIndex(128, (d_290_v3_).length(0))
                    (d_290_v3_)[index32_] = p2
                    index33_ = default__.safeIndex(128, (d_290_v3_).length(0))
                    (d_290_v3_)[index33_] = p3
                    d_291_v4_: _dafny.Seq
                    d_291_v4_ = _dafny.SeqWithoutIsStrInference([(self).f10])
                    rhs27_ = (self).f10
                    rhs28_ = (d_291_v4_) + (d_291_v4_)
                    rhs29_ = (self).f10
                    rhs30_ = (self).f10
                    lhs26_ = globalState
                    lhs27_ = globalState
                    r0 = rhs27_
                    d_291_v4_ = rhs28_
                    lhs26_.f2 = rhs29_
                    lhs27_.f2 = rhs30_
                    (globalState).f5 = default__.safeModuloInt(default__.safeDivisionInt((d_290_v3_)[default__.safeIndex(128, (d_290_v3_).length(0))], len(d_286_v0_)), p3)
                    pass
            pass
        d_292_v5_: _dafny.Array
        nw60_ = _dafny.Array(_dafny.Map({}), 17)
        d_292_v5_ = nw60_
        d_293_v6_: D1
        d_293_v6_ = D1_DC3(d_292_v5_)
        d_294_v7_: _dafny.Map
        d_294_v7_ = _dafny.Map({(self).f10: d_292_v5_})
        d_295_v8_: _dafny.Seq
        d_295_v8_ = _dafny.SeqWithoutIsStrInference([d_292_v5_, d_292_v5_])
        d_296_v9_: _dafny.Array
        nw61_ = _dafny.Array(None, 21)
        nw61_[int(0)] = d_292_v5_
        nw61_[int(1)] = d_292_v5_
        nw61_[int(2)] = d_292_v5_
        nw61_[int(3)] = d_292_v5_
        nw61_[int(4)] = d_292_v5_
        nw61_[int(5)] = d_292_v5_
        nw61_[int(6)] = d_292_v5_
        nw61_[int(7)] = (d_293_v6_).cf3
        nw61_[int(8)] = d_292_v5_
        nw61_[int(9)] = d_292_v5_
        nw61_[int(10)] = d_292_v5_
        nw61_[int(11)] = d_292_v5_
        nw61_[int(12)] = ((d_294_v7_)[(self).f10] if ((self).f10) in (d_294_v7_) else (d_295_v8_)[default__.safeIndex(517, len(d_295_v8_))])
        nw61_[int(13)] = d_292_v5_
        nw61_[int(14)] = (d_292_v5_ if (self).f10 else d_292_v5_)
        nw61_[int(15)] = d_292_v5_
        nw61_[int(16)] = d_292_v5_
        nw61_[int(17)] = d_292_v5_
        nw61_[int(18)] = d_292_v5_
        nw61_[int(19)] = d_292_v5_
        nw61_[int(20)] = d_292_v5_
        d_296_v9_ = nw61_
        index34_ = default__.safeIndex(88, (d_296_v9_).length(0))
        (d_296_v9_)[index34_] = d_292_v5_
        index35_ = default__.safeIndex(88, (d_296_v9_).length(0))
        rhs31_ = d_292_v5_
        rhs32_ = default__.fm0(len(d_286_v0_), p2, p0, ((self).f10) == (not((self).f10)), globalState)
        rhs33_ = p0
        lhs28_ = d_296_v9_
        lhs29_ = default__.safeIndex(88, (d_296_v9_).length(0))
        lhs30_ = globalState
        lhs28_[lhs29_] = rhs31_
        lhs30_.f5 = rhs32_
        r1 = rhs33_
        d_297_v10_: _dafny.Set
        d_297_v10_ = _dafny.Set({(self).f10})
        d_298_v11_: _dafny.Set
        d_298_v11_ = _dafny.Set({len(d_297_v10_)})
        d_298_v11_ = default__.fm6(globalState)
        if not(True):
            (globalState).f0 = ((self).f10 if (self).f10 else not(False))
            d_299_v12_: _dafny.Map
            d_299_v12_ = _dafny.Map({(self).f10: not((self).f10)})
            d_300_v13_: _dafny.Seq
            d_300_v13_ = _dafny.SeqWithoutIsStrInference([(((d_299_v12_)[(self).f10] if ((self).f10) in (d_299_v12_) else (self).f10)) == ((self).f10)])
            d_300_v13_ = (_dafny.SeqWithoutIsStrInference([(self).f10])) + ((d_300_v13_) + (d_300_v13_))
            d_301_v14_: _dafny.Array
            def lambda18_(d_302_p3_):
                def lambda19_(d_303_i1_):
                    return default__.safeModuloInt(d_303_i1_, d_302_p3_)

                return lambda19_

            init9_ = lambda18_(p3)
            nw62_ = _dafny.Array(None, 7)
            for i0_9_ in range(nw62_.length(0)):
                nw62_[i0_9_] = init9_(i0_9_)
            d_301_v14_ = nw62_
            d_301_v14_ = d_301_v14_
            r1 = p0
            (globalState).f0 = (self).f10
        elif True:
            d_304_v15_: _dafny.MultiSet
            d_304_v15_ = _dafny.MultiSet([(self).f10])
            (globalState).f2 = (d_304_v15_).ispropersubset((default__.fm7(globalState)) | ((d_304_v15_).set((self).f10, default__.abs(p3))))
            r0 = True
            d_305_v16_: _dafny.Array
            def lambda20_(d_306_p0_):
                def lambda21_(d_307_i2_):
                    return (d_307_i2_) - (d_306_p0_)

                return lambda21_

            init10_ = lambda20_(p0)
            nw63_ = _dafny.Array(None, 13)
            for i0_10_ in range(nw63_.length(0)):
                nw63_[i0_10_] = init10_(i0_10_)
            d_305_v16_ = nw63_
            d_305_v16_ = d_305_v16_
            d_308_v17_: _dafny.Seq
            d_308_v17_ = _dafny.SeqWithoutIsStrInference([False])
            index36_ = default__.safeIndex(773, ((self).f11).length(0))
            ((self).f11)[index36_] = ((d_308_v17_).set(default__.safeIndex(p2, len(d_308_v17_)), (d_308_v17_)[default__.safeIndex(len(d_308_v17_), len(d_308_v17_))])) <= (d_308_v17_)
            index37_ = default__.safeIndex(773, ((self).f11).length(0))
            ((self).f11)[index37_] = (p3) >= (p2)
            if (self).f10:
                (globalState).f5 = default__.fm0(p3, p3, p2, ((self).f11)[default__.safeIndex(773, ((self).f11).length(0))], globalState)
                (globalState).f0 = default__.fm1(False, globalState)
                d_309_v18_: _dafny.Map
                d_309_v18_ = _dafny.Map({p3: ((self).f11)[default__.safeIndex(773, ((self).f11).length(0))]})
                d_310_v19_: _dafny.Map
                d_310_v19_ = _dafny.Map({len(d_309_v18_): default__.safeDivisionInt(p2, 682)})
                d_310_v19_ = (d_310_v19_).set(p0, ((d_310_v19_)[(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "weoxwqp"))))] if ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "weoxwqp"))))) in (d_310_v19_) else p2))
                (globalState).f2 = (self).f10
                d_311_v20_: _dafny.Map
                d_311_v20_ = _dafny.Map({d_297_v10_: (p3) == (p3)})
                d_311_v20_ = d_311_v20_
            elif True:
                d_312_v21_: str
                d_312_v21_ = _dafny.CodePoint('h')
                d_312_v21_ = d_312_v21_
                d_313_v22_: _dafny.Map
                d_313_v22_ = _dafny.Map({p0: p3})
                d_314_v23_: _dafny.Map
                d_314_v23_ = _dafny.Map({((self).f11)[default__.safeIndex(773, ((self).f11).length(0))]: (self).f10})
                d_315_v24_: _dafny.Map
                d_315_v24_ = _dafny.Map({((self).f11)[default__.safeIndex(773, ((self).f11).length(0))]: (((d_313_v22_)[default__.fm0(p2, len(d_314_v23_), p2, (self).f10, globalState)] if (default__.fm0(p2, len(d_314_v23_), p2, (self).f10, globalState)) in (d_313_v22_) else p0)) * ((0) - (p0))})
                d_315_v24_ = (d_315_v24_).set((self).f10, (p0) - (p0))
                d_316_v25_: _dafny.Map
                d_316_v25_ = _dafny.Map({p3: (self).f10})
                index38_ = default__.safeIndex(317, (d_305_v16_).length(0))
                (d_305_v16_)[index38_] = -985
                d_317_v26_: D0
                d_317_v26_ = D0_DC0(d_316_v25_)
                d_318_v27_: _dafny.Set
                d_318_v27_ = _dafny.Set({d_312_v21_})
                pat_let_tv1_ = d_296_v9_
                pat_let_tv2_ = d_296_v9_
                index39_ = default__.safeIndex(317, (d_305_v16_).length(0))
                def iife25_(_pat_let2_0):
                    def iife26_(d_319_dt__update__tmp_h0_):
                        def iife27_(_pat_let3_0):
                            def iife28_(d_320_dt__update_hcf3_h0_):
                                return D1_DC3(d_320_dt__update_hcf3_h0_)
                            return iife28_(_pat_let3_0)
                        return iife27_((pat_let_tv2_)[default__.safeIndex(88, (pat_let_tv1_).length(0))])
                    return iife26_(_pat_let2_0)
                rhs34_ = (d_317_v26_).cf0
                rhs35_ = 699
                rhs36_ = iife25_(d_293_v6_)
                rhs37_ = ((self).f10 if (d_318_v27_).ispropersubset(d_318_v27_) else ((self).f11)[default__.safeIndex(773, ((self).f11).length(0))])
                lhs31_ = d_305_v16_
                lhs32_ = default__.safeIndex(317, (d_305_v16_).length(0))
                lhs33_ = globalState
                d_316_v25_ = rhs34_
                lhs31_[lhs32_] = rhs35_
                d_293_v6_ = rhs36_
                lhs33_.f0 = rhs37_
                d_321_v28_: _dafny.Map
                d_321_v28_ = _dafny.Map({D1_DC4(p3): p2})
                d_322_v29_: D1
                d_322_v29_ = D1_DC4((d_305_v16_)[default__.safeIndex(317, (d_305_v16_).length(0))])
                d_323_v30_: _dafny.MultiSet
                d_323_v30_ = _dafny.MultiSet([675])
                d_321_v28_ = (d_321_v28_).set(d_322_v29_, default__.fm0(492, p2, (d_323_v30_).cardinality, False, globalState))
                (globalState).f2 = True
        d_324_v31_: _dafny.Seq
        d_324_v31_ = _dafny.SeqWithoutIsStrInference([p0])
        d_325_v32_: _dafny.Map
        d_325_v32_ = _dafny.Map({(d_324_v31_)[default__.safeIndex(p0, len(d_324_v31_))]: (0) - (p0)})
        r1 = (len(d_325_v32_)) - (((0) - (p3)) - (p3))
        r0 = (self).f10
        d_326_v33_: _dafny.Map
        d_326_v33_ = _dafny.Map({p0: True})
        d_327_v34_: D0
        d_327_v34_ = D0_DC0(d_326_v33_)
        d_328_v35_: _dafny.MultiSet
        d_328_v35_ = _dafny.MultiSet([(self).f10, (self).f10])
        d_329_v36_: _dafny.MultiSet
        d_329_v36_ = _dafny.MultiSet([(self).f10, (self).fm4(d_327_v34_, d_327_v34_, d_328_v35_, globalState), (self).f10])
        d_330_v37_: _dafny.Map
        d_330_v37_ = _dafny.Map({(self).fm4(d_327_v34_, d_327_v34_, (d_328_v35_).set((self).f10, default__.abs(p0)), globalState): p3})
        r1 = ((d_330_v37_)[(self).f10] if ((self).f10) in (d_330_v37_) else p3)
        r2 = ((self).f10) and (((d_326_v33_)[p3] if (p3) in (d_326_v33_) else (self).fm4(d_327_v34_, d_327_v34_, d_329_v36_, globalState)))
        d_331_v38_: str
        d_331_v38_ = _dafny.CodePoint('w')
        d_332_v39_: _dafny.Map
        d_332_v39_ = _dafny.Map({(self).f10: d_331_v38_})
        r3 = (default__.fm8((self).f10, d_332_v39_, (self).f10, globalState)) < (d_286_v0_)
        return r0, r1, r2, r3

    @property
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11

class C1(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def fm4(self, p0, p1, p2, globalState):
        source5_ = D1_DC4(-738)
        if source5_.is_DC4:
            d_333___mcc_h0_ = source5_.cf4
            d_334_cf4_ = d_333___mcc_h0_
            return True
        elif True:
            d_335___mcc_h1_ = source5_.cf3
            d_336_cf3_ = d_335___mcc_h1_
            return True

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        r3: bool = False
        d_337_v0_: bool
        d_337_v0_ = True
        d_338_v1_: _dafny.Array
        nw64_ = _dafny.Array(False, 9)
        d_338_v1_ = nw64_
        d_339_v2_: C0
        nw65_ = C0()
        nw65_.ctor__(d_337_v0_, d_338_v1_)
        d_339_v2_ = nw65_
        index40_ = default__.safeIndex(154, ((d_339_v2_).f11).length(0))
        ((d_339_v2_).f11)[index40_] = True
        d_340_v3_: _dafny.Map
        d_340_v3_ = _dafny.Map({p2: (d_339_v2_).f10})
        d_341_v4_: D0
        d_341_v4_ = D0_DC0(d_340_v3_)
        d_342_v5_: D0
        d_342_v5_ = D0_DC0((d_341_v4_).cf0)
        d_343_v6_: _dafny.MultiSet
        d_343_v6_ = _dafny.MultiSet([default__.fm1(d_337_v0_, globalState)])
        d_344_v7_: _dafny.Map
        d_344_v7_ = _dafny.Map({p2: d_343_v6_})
        index41_ = default__.safeIndex(154, ((d_339_v2_).f11).length(0))
        ((d_339_v2_).f11)[index41_] = (d_339_v2_).fm4(d_341_v4_, d_342_v5_, ((d_344_v7_)[p3] if (p3) in (d_344_v7_) else d_343_v6_), globalState)
        d_345_v8_: _dafny.Seq
        d_345_v8_ = _dafny.SeqWithoutIsStrInference([p0])
        d_346_v9_: _dafny.Seq
        d_346_v9_ = d_345_v8_
        (globalState).f5 = default__.fm0(p3, p2, 474, ((d_346_v9_)) <= (d_345_v8_), globalState)
        hi1_ = p0
        for d_347_i0_ in range(p2, hi1_):
            d_348_v10_: D7
            d_348_v10_ = D7_DC17(_dafny.SeqWithoutIsStrInference([(d_339_v2_).f10]))
            d_349_v11_: _dafny.Seq
            d_349_v11_ = _dafny.SeqWithoutIsStrInference([((d_339_v2_).f11)[default__.safeIndex(154, ((d_339_v2_).f11).length(0))], (d_339_v2_).f10, ((d_339_v2_).f11)[default__.safeIndex(154, ((d_339_v2_).f11).length(0))], (self).fm4(D0_DC0(default__.fm22((d_339_v2_).f10, globalState)), d_341_v4_, d_343_v6_, globalState), True])
            d_350_v12_: _dafny.Set
            d_350_v12_ = _dafny.Set({d_349_v11_})
            if ((d_348_v10_).cf23) in (d_350_v12_):
                (globalState).f0 = (D2_DC6((d_339_v2_).f10)).cf6
                r2 = (p3) >= (p0)
                d_351_v13_: _dafny.Map
                d_351_v13_ = _dafny.Map({_dafny.Map({d_337_v0_: d_337_v0_}): p0})
                d_352_v14_: D2
                d_352_v14_ = D2_DC7(len(d_349_v11_), d_351_v13_, d_347_i0_)
                (globalState).f5 = (d_352_v14_).cf9
                d_353_v15_: _dafny.MultiSet
                d_353_v15_ = _dafny.MultiSet([493])
                d_353_v15_ = default__.fm23((d_353_v15_).set(default__.fm0(p0, p3, p2, (d_339_v2_).f10, globalState), default__.abs(p3)), p3, globalState)
                rhs38_ = p2
                rhs39_ = d_337_v0_
                rhs40_ = (d_339_v2_).f11
                rhs41_ = d_338_v1_
                rhs42_ = len((d_340_v3_) | (d_340_v3_))
                lhs34_ = globalState
                r1 = rhs38_
                lhs34_.f0 = rhs39_
                d_338_v1_ = rhs40_
                d_338_v1_ = rhs41_
                r1 = rhs42_
            elif True:
                d_354_v16_: _dafny.Map
                d_354_v16_ = _dafny.Map({p3: d_338_v1_})
                d_354_v16_ = ((d_354_v16_) | (d_354_v16_)) | (d_354_v16_)
                r2 = d_337_v0_
                d_355_v17_: T1
                nw66_ = C0()
                nw66_.ctor__((d_339_v2_).f10, (d_339_v2_).f11)
                d_355_v17_ = nw66_
                d_356_v18_: _dafny.Set
                d_356_v18_ = _dafny.Set({d_355_v17_})
                d_357_v19_: D8
                d_357_v19_ = D8_DC20(d_356_v18_)
                rhs43_ = ((d_339_v2_).f11 if not((d_339_v2_).f10) else (d_339_v2_).f11)
                rhs44_ = ((d_356_v18_).issubset((d_357_v19_).cf25) if not((d_339_v2_).f10) else (d_339_v2_).f10)
                rhs45_ = (d_339_v2_).f10
                lhs35_ = globalState
                lhs36_ = globalState
                d_338_v1_ = rhs43_
                lhs35_.f2 = rhs44_
                lhs36_.f2 = rhs45_
                d_358_v20_: C0
                nw67_ = C0()
                nw67_.ctor__(((d_339_v2_).f11)[default__.safeIndex(154, ((d_339_v2_).f11).length(0))], (d_339_v2_).f11)
                d_358_v20_ = nw67_
                r0 = ((d_339_v2_).f10) and (True)
            d_359_v21_: C0
            nw68_ = C0()
            nw68_.ctor__(((d_339_v2_).f11)[default__.safeIndex(154, ((d_339_v2_).f11).length(0))], (d_339_v2_).f11)
            d_359_v21_ = nw68_
            (globalState).f5 = (_dafny.MultiSet([(d_339_v2_).f10])).cardinality
            d_360_v22_: _dafny.Array
            nw69_ = _dafny.Array(int(0), 6)
            d_360_v22_ = nw69_
            d_361_v23_: _dafny.Array
            d_361_v23_ = d_360_v22_
            d_362_v24_: _dafny.Map
            d_362_v24_ = _dafny.Map({(d_361_v23_ if (d_339_v2_).f10 else d_361_v23_): (D4_DC13(p3, d_338_v1_)).cf19})
            d_362_v24_ = (d_362_v24_).set(d_361_v23_, default__.safeDivisionInt(p0, p3))
        d_363_v25_: D2
        d_363_v25_ = D2_DC6(False)
        d_364_i1_: int
        d_364_i1_ = 0
        with _dafny.label("2"):
            while (d_363_v25_).cf6:
                with _dafny.c_label("2"):
                    if (d_364_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_364_i1_ = (d_364_i1_) + (1)
                    d_365_v26_: C0
                    nw70_ = C0()
                    nw70_.ctor__(d_337_v0_, d_338_v1_)
                    d_365_v26_ = nw70_
                    (globalState).f2 = (default__.fm0(p0, p0, p2, (d_339_v2_).f10, globalState)) != (len(d_340_v3_))
                    d_340_v3_ = (d_340_v3_).set(p3, ((d_339_v2_).f11)[default__.safeIndex(154, ((d_339_v2_).f11).length(0))])
                    d_366_v27_: str
                    d_366_v27_ = _dafny.CodePoint('m')
                    d_366_v27_ = d_366_v27_
                    pass
            pass
        d_367_v28_: _dafny.Seq
        d_367_v28_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_345_v8_)])
        pat_let_tv3_ = d_337_v0_
        pat_let_tv4_ = p0
        pat_let_tv5_ = d_339_v2_
        pat_let_tv6_ = d_339_v2_
        pat_let_tv7_ = d_339_v2_
        pat_let_tv8_ = d_339_v2_
        pat_let_tv9_ = p0
        def lambda22_(source6_):
            if source6_.is_DC1:
                d_368___mcc_h0_ = source6_.cf1
                d_369_cf1_ = d_368___mcc_h0_
                return 527
            elif source6_.is_DC0:
                d_370___mcc_h1_ = source6_.cf0
                d_371_cf0_ = d_370___mcc_h1_
                return (len(_dafny.Set({pat_let_tv3_}))) * (pat_let_tv4_)
            elif True:
                d_372___mcc_h2_ = source6_.cf2
                d_373_cf2_ = d_372___mcc_h2_
                return default__.safeModuloInt(len(_dafny.Map({((pat_let_tv6_).f11)[default__.safeIndex(154, ((pat_let_tv5_).f11).length(0))]: ((pat_let_tv8_).f11)[default__.safeIndex(154, ((pat_let_tv7_).f11).length(0))]})), pat_let_tv9_)

        rhs46_ = not(not (d_337_v0_) or (d_337_v0_))
        rhs47_ = default__.safeDivisionInt(p3, 652)
        rhs48_ = ((d_343_v6_) | (d_343_v6_)).isdisjoint((d_343_v6_).set(d_337_v0_, default__.abs((0) - ((_dafny.MultiSet(d_367_v28_)).cardinality))))
        rhs49_ = lambda22_(d_341_v4_)
        rhs50_ = default__.safeModuloInt(default__.safeDivisionInt(p0, p2), (0) - (default__.safeDivisionInt(p3, p2)))
        lhs37_ = globalState
        r3 = rhs46_
        lhs37_.f5 = rhs47_
        d_337_v0_ = rhs48_
        r1 = rhs49_
        r1 = rhs50_
        d_374_v29_: _dafny.Map
        d_374_v29_ = _dafny.Map({((d_339_v2_).f11)[default__.safeIndex(154, ((d_339_v2_).f11).length(0))]: False})
        d_375_v30_: _dafny.Map
        d_375_v30_ = _dafny.Map({d_374_v29_: p0})
        d_376_v31_: _dafny.Seq
        d_376_v31_ = _dafny.SeqWithoutIsStrInference([((d_339_v2_).f11)[default__.safeIndex(154, ((d_339_v2_).f11).length(0))]])
        d_377_v32_: D2
        d_377_v32_ = D2_DC7(p3, d_375_v30_, len(d_376_v31_))
        r0 = (p3) > ((d_377_v32_).cf9)
        r1 = p2
        r2 = (True) not in ((_dafny.Map({d_337_v0_: len(d_376_v31_)})) | (_dafny.Map({(self).fm4(D0_DC0(d_340_v3_), d_342_v5_, d_343_v6_, globalState): default__.fm0(p2, p0, p2, ((d_339_v2_).f11)[default__.safeIndex(154, ((d_339_v2_).f11).length(0))], globalState)})))
        r3 = d_337_v0_
        return r0, r1, r2, r3

    def m6(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        (globalState).f0 = p2
        d_378_v0_: _dafny.Seq
        d_378_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igpdhpy"))
        d_379_v1_: _dafny.Set
        d_379_v1_ = _dafny.Set({d_378_v0_})
        d_379_v1_ = d_379_v1_
        if p2:
            d_380_v2_: D2
            d_380_v2_ = D2_DC6(p2)
            d_381_v3_: _dafny.Seq
            d_381_v3_ = _dafny.SeqWithoutIsStrInference([(d_380_v2_).cf6, p2, p0, p2, not(p0)])
            d_382_v4_: _dafny.MultiSet
            d_382_v4_ = _dafny.MultiSet([default__.safeDivisionInt(len(d_381_v3_), p1)])
            d_383_v5_: _dafny.Map
            d_383_v5_ = _dafny.Map({p2: p0})
            d_384_v6_: _dafny.Array
            nw71_ = _dafny.Array(None, 11)
            nw71_[int(0)] = p0
            nw71_[int(1)] = (((d_383_v5_)[p2] if (p2) in (d_383_v5_) else p0)) and (p2)
            nw71_[int(2)] = (p0 if p2 else p2)
            nw71_[int(3)] = True
            nw71_[int(4)] = p2
            nw71_[int(5)] = p2
            nw71_[int(6)] = p0
            nw71_[int(7)] = p0
            nw71_[int(8)] = p0
            nw71_[int(9)] = p2
            nw71_[int(10)] = p2
            d_384_v6_ = nw71_
            d_385_v7_: _dafny.Map
            d_385_v7_ = _dafny.Map({(d_381_v3_)[default__.safeIndex((d_382_v4_).cardinality, len(d_381_v3_))]: d_381_v3_})
            rhs51_ = d_382_v4_
            rhs52_ = len((d_378_v0_ if p0 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "inuynen"))))
            rhs53_ = not((not(p0)) and (p2))
            rhs54_ = ((p1) * (p1)) >= (len(d_385_v7_))
            rhs55_ = d_384_v6_
            lhs38_ = globalState
            lhs39_ = globalState
            lhs40_ = globalState
            d_382_v4_ = rhs51_
            lhs38_.f5 = rhs52_
            lhs39_.f0 = rhs53_
            lhs40_.f0 = rhs54_
            d_384_v6_ = rhs55_
            d_386_v8_: _dafny.Array
            def lambda23_(d_387_p1_):
                def lambda24_(d_388_i0_):
                    return (d_388_i0_) * (d_387_p1_)

                return lambda24_

            init11_ = lambda23_(p1)
            nw72_ = _dafny.Array(None, 10)
            for i0_11_ in range(nw72_.length(0)):
                nw72_[i0_11_] = init11_(i0_11_)
            d_386_v8_ = nw72_
            index42_ = default__.safeIndex(572, (d_386_v8_).length(0))
            (d_386_v8_)[index42_] = p1
            d_389_v10_: _dafny.Array
            def lambda25_(d_390_p0_, d_391_p1_):
                def lambda26_(d_392_i1_):
                    def iife29_():
                        coll21_ = _dafny.Set()
                        compr_21_: int
                        for compr_21_ in (_dafny.SeqWithoutIsStrInference([d_391_p1_])).Elements:
                            d_393_v9_: int = compr_21_
                            if (d_393_v9_) in (_dafny.SeqWithoutIsStrInference([d_391_p1_])):
                                coll21_ = coll21_.union(_dafny.Set([default__.safeModuloInt(d_393_v9_, 718)]))
                        return _dafny.Set(coll21_)
                    return (_dafny.Map({d_390_p0_: _dafny.Map({d_390_p0_: len(_dafny.Set({_dafny.Set({d_391_p1_}), iife29_()
                    }))})})) | (_dafny.Map({d_390_p0_: _dafny.Map({d_390_p0_: d_391_p1_})}))

                return lambda26_

            init12_ = lambda25_(p0, p1)
            nw73_ = _dafny.Array(None, 6)
            for i0_12_ in range(nw73_.length(0)):
                nw73_[i0_12_] = init12_(i0_12_)
            d_389_v10_ = nw73_
            d_394_v11_: _dafny.Map
            d_394_v11_ = _dafny.Map({d_384_v6_: p0})
            d_395_v12_: D9
            d_395_v12_ = D9_DC23(d_394_v11_)
            index43_ = default__.safeIndex(659, (d_389_v10_).length(0))
            (d_389_v10_)[index43_] = default__.fm24(len((d_395_v12_).cf32), globalState)
            d_396_v13_: _dafny.Set
            d_396_v13_ = _dafny.Set({default__.fm25(((d_383_v5_)[p2] if (p2) in (d_383_v5_) else p2), _dafny.CodePoint('n'), p0, globalState), d_381_v3_, d_381_v3_})
            d_397_v14_: _dafny.Map
            d_397_v14_ = _dafny.Map({False: p1})
            d_398_v15_: _dafny.Map
            d_398_v15_ = _dafny.Map({p0: d_397_v14_})
            index44_ = default__.safeIndex(572, (d_386_v8_).length(0))
            index45_ = default__.safeIndex(659, (d_389_v10_).length(0))
            rhs56_ = len(d_396_v13_)
            rhs57_ = d_381_v3_
            rhs58_ = (d_398_v15_) | (d_398_v15_)
            lhs41_ = d_386_v8_
            lhs42_ = default__.safeIndex(572, (d_386_v8_).length(0))
            lhs43_ = d_389_v10_
            lhs44_ = default__.safeIndex(659, (d_389_v10_).length(0))
            lhs41_[lhs42_] = rhs56_
            d_381_v3_ = rhs57_
            lhs43_[lhs44_] = rhs58_
            d_399_v16_: _dafny.Seq
            d_399_v16_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bodi"))), p1, p1, -446, (d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))]])
            d_400_v17_: _dafny.Set
            d_400_v17_ = _dafny.Set({not(default__.fm1(p2, globalState))})
            if (((d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))]) - ((d_399_v16_)[default__.safeIndex(p1, len(d_399_v16_))])) != ((p1) - ((0) - (len(d_400_v17_)))):
                d_401_v18_: D4
                d_401_v18_ = D4_DC13((d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))], d_384_v6_)
                d_401_v18_ = D4_DC13((d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))], d_384_v6_)
                (globalState).f2 = ((0) - (p1)) < (((d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))]) * (len(d_383_v5_)))
                d_402_v19_: _dafny.Map
                d_402_v19_ = _dafny.Map({(d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))]: d_384_v6_})
                d_402_v19_ = (d_402_v19_).set(p1, d_384_v6_)
                d_403_v20_: _dafny.Map
                d_403_v20_ = _dafny.Map({(d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))]: not(not((p2) and (p0)))})
                pat_let_tv10_ = p2
                def iife30_(_pat_let4_0):
                    def iife31_(d_404_dt__update__tmp_h0_):
                        def iife32_(_pat_let5_0):
                            def iife33_(d_405_dt__update_hcf6_h0_):
                                return D2_DC6(d_405_dt__update_hcf6_h0_)
                            return iife33_(_pat_let5_0)
                        return iife32_(pat_let_tv10_)
                    return iife31_(_pat_let4_0)
                rhs59_ = ((d_403_v20_)[default__.fm0(len(d_400_v17_), (d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))], -119, p2, globalState)] if (default__.fm0(len(d_400_v17_), (d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))], -119, p2, globalState)) in (d_403_v20_) else p0)
                rhs60_ = 437
                rhs61_ = (iife30_(d_380_v2_)).cf6
                rhs62_ = (d_378_v0_) + ((d_378_v0_) + (d_378_v0_))
                lhs45_ = globalState
                lhs46_ = globalState
                lhs47_ = globalState
                lhs45_.f2 = rhs59_
                lhs46_.f5 = rhs60_
                lhs47_.f0 = rhs61_
                d_378_v0_ = rhs62_
                d_406_v21_: C0
                nw74_ = C0()
                nw74_.ctor__(p0, d_384_v6_)
                d_406_v21_ = nw74_
            elif True:
                index46_ = default__.safeIndex(481, (d_384_v6_).length(0))
                (d_384_v6_)[index46_] = p2
                d_407_v22_: _dafny.Array
                d_407_v22_ = d_386_v8_
                index47_ = default__.safeIndex(481, (d_384_v6_).length(0))
                index48_ = default__.safeIndex(572, (d_386_v8_).length(0))
                rhs63_ = d_384_v6_
                rhs64_ = (d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))]
                rhs65_ = p0
                rhs66_ = p1
                rhs67_ = (d_407_v22_)
                lhs48_ = globalState
                lhs49_ = d_384_v6_
                lhs50_ = default__.safeIndex(481, (d_384_v6_).length(0))
                lhs51_ = d_386_v8_
                lhs52_ = default__.safeIndex(572, (d_386_v8_).length(0))
                d_384_v6_ = rhs63_
                lhs48_.f5 = rhs64_
                lhs49_[lhs50_] = rhs65_
                lhs51_[lhs52_] = rhs66_
                d_386_v8_ = rhs67_
                d_408_v23_: _dafny.MultiSet
                d_408_v23_ = _dafny.MultiSet([((d_380_v2_).cf6 if p0 else (d_381_v3_)[default__.safeIndex((d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))], len(d_381_v3_))])])
                d_409_v24_: D0
                d_409_v24_ = D0_DC0(default__.fm22(p2, globalState))
                rhs68_ = (0) - ((d_408_v23_).cardinality)
                rhs69_ = default__.safeDivisionInt((0) - ((default__.fm26((d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))], 901, False, (self).fm4(d_409_v24_, d_409_v24_, d_408_v23_, globalState), globalState)).cf9), p1)
                lhs53_ = globalState
                lhs54_ = globalState
                lhs53_.f5 = rhs68_
                lhs54_.f5 = rhs69_
                d_410_v26_: _dafny.Map
                d_410_v26_ = _dafny.Map({len(d_378_v0_): d_382_v4_})
                d_411_v27_: _dafny.Map
                def iife34_():
                    coll22_ = _dafny.Map()
                    compr_22_: int
                    for compr_22_ in _dafny.IntegerRange(850, 991):
                        d_412_v25_: int = compr_22_
                        if ((850) <= (d_412_v25_)) and ((d_412_v25_) < (991)):
                            coll22_[default__.safeModuloInt(d_412_v25_, 416)] = 132
                    return _dafny.Map(coll22_)
                d_411_v27_ = _dafny.Map({(((d_382_v4_)[p1] if (p1) in (d_382_v4_) else (d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))])) * (len(iife34_()
                )): (((d_410_v26_)[(_dafny.MultiSet([(d_384_v6_)[default__.safeIndex(481, (d_384_v6_).length(0))], p0])).cardinality] if ((_dafny.MultiSet([(d_384_v6_)[default__.safeIndex(481, (d_384_v6_).length(0))], p0])).cardinality) in (d_410_v26_) else d_382_v4_)).intersection(d_382_v4_)})
                d_411_v27_ = (d_411_v27_).set((p1) - (p1), d_382_v4_)
                index49_ = default__.safeIndex(481, (d_384_v6_).length(0))
                (d_384_v6_)[index49_] = (default__.safeModuloInt(p1, p1)) > (((d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))]) + (p1))
                d_413_v28_: _dafny.Array
                def lambda27_(d_414_p2_, d_415_p1_, d_416_v14_):
                    def lambda28_(d_417_i2_):
                        return (_dafny.Map({d_414_p2_: d_415_p1_}) if False else d_416_v14_)

                    return lambda28_

                init13_ = lambda27_(p2, p1, d_397_v14_)
                nw75_ = _dafny.Array(None, 1)
                for i0_13_ in range(nw75_.length(0)):
                    nw75_[i0_13_] = init13_(i0_13_)
                d_413_v28_ = nw75_
                d_418_v29_: _dafny.Array
                nw76_ = _dafny.Array(False, 8)
                d_418_v29_ = nw76_
                index50_ = default__.safeIndex(572, (d_386_v8_).length(0))
                rhs70_ = d_413_v28_
                rhs71_ = (default__.fm0((d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))], p1, len(d_378_v0_), not(p0), globalState)) + (len(d_378_v0_))
                rhs72_ = (d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))]
                rhs73_ = (p1) + (len(_dafny.SeqWithoutIsStrInference([d_399_v16_, d_399_v16_])))
                rhs74_ = d_418_v29_
                lhs55_ = d_386_v8_
                lhs56_ = default__.safeIndex(572, (d_386_v8_).length(0))
                lhs57_ = globalState
                d_413_v28_ = rhs70_
                lhs55_[lhs56_] = rhs71_
                r0 = rhs72_
                lhs57_.f5 = rhs73_
                d_418_v29_ = rhs74_
            d_419_v30_: _dafny.Map
            d_419_v30_ = _dafny.Map({_dafny.Map({False: p2}): (d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))]})
            d_420_v31_: D2
            d_420_v31_ = D2_DC7(p1, d_419_v30_, (d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))])
            d_421_v32_: _dafny.Array
            nw77_ = _dafny.Array(None, 7)
            nw77_[int(0)] = d_420_v31_
            nw77_[int(1)] = d_420_v31_
            nw77_[int(2)] = d_420_v31_
            nw77_[int(3)] = default__.fm26((d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))], (d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))], p2, not(True), globalState)
            nw77_[int(4)] = D2_DC7(p1, d_419_v30_, 587)
            nw77_[int(5)] = D2_DC7((d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))], d_419_v30_, len(d_378_v0_))
            nw77_[int(6)] = default__.fm26(p1, p1, not(not(p0)), p2, globalState)
            d_421_v32_ = nw77_
            index51_ = default__.safeIndex(368, (d_421_v32_).length(0))
            (d_421_v32_)[index51_] = D2_DC7((d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))], d_419_v30_, p1)
            index52_ = default__.safeIndex(368, (d_421_v32_).length(0))
            (d_421_v32_)[index52_] = d_420_v31_
            d_422_v33_: D4
            d_422_v33_ = D4_DC14()
            d_423_v34_: _dafny.Seq
            d_423_v34_ = _dafny.SeqWithoutIsStrInference([d_422_v33_, D4_DC14(), d_422_v33_])
            d_424_v35_: str
            d_424_v35_ = _dafny.CodePoint('s')
            d_425_v36_: _dafny.Map
            d_425_v36_ = _dafny.Map({p2: d_384_v6_})
            d_426_v37_: _dafny.Set
            d_426_v37_ = _dafny.Set({(d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))], (len(d_425_v36_)) * ((0) - ((d_386_v8_)[default__.safeIndex(572, (d_386_v8_).length(0))])), 350, p1, -914})
            rhs75_ = _dafny.SeqWithoutIsStrInference([d_422_v33_, d_422_v33_, D4_DC14(), d_422_v33_, d_422_v33_])
            rhs76_ = default__.fm27(globalState)
            rhs77_ = d_424_v35_
            rhs78_ = (d_426_v37_) - ((d_426_v37_) - (d_426_v37_))
            rhs79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cps"))
            d_423_v34_ = rhs75_
            d_378_v0_ = rhs76_
            d_424_v35_ = rhs77_
            d_426_v37_ = rhs78_
            d_378_v0_ = rhs79_
        elif True:
            d_427_v38_: _dafny.Map
            d_427_v38_ = _dafny.Map({p0: p0})
            d_428_v39_: D0
            d_428_v39_ = D0_DC1(p2)
            d_429_v40_: _dafny.Map
            d_429_v40_ = _dafny.Map({len(d_427_v38_): d_428_v39_})
            d_430_v41_: D0
            d_430_v41_ = D0_DC2(((d_429_v40_)[p1] if (p1) in (d_429_v40_) else d_428_v39_))
            d_430_v41_ = d_430_v41_
            d_431_v42_: _dafny.Map
            d_431_v42_ = _dafny.Map({(0) - ((941) + (p1)): (0) - (p1)})
            d_432_v43_: _dafny.Set
            d_432_v43_ = _dafny.Set({p0, p2, p0, p2})
            d_431_v42_ = (d_431_v42_).set(len(d_432_v43_), len(_dafny.Set({p2})))
            d_433_v45_: _dafny.Set
            d_433_v45_ = _dafny.Set({p1, 584})
            def iife35_():
                coll23_ = _dafny.Set()
                compr_23_: int
                for compr_23_ in _dafny.IntegerRange(508, 914):
                    d_434_v44_: int = compr_23_
                    if ((508) <= (d_434_v44_)) and ((d_434_v44_) < (914)):
                        coll23_ = coll23_.union(_dafny.Set([(d_434_v44_) + ((0) - (p1))]))
                return _dafny.Set(coll23_)
            r2 = ((d_433_v45_) | (d_433_v45_)).ispropersubset((iife35_()
            ) | (d_433_v45_))
            r0 = p1
            d_435_v47_: _dafny.Seq
            def iife36_():
                coll24_ = _dafny.Set()
                compr_24_: int
                for compr_24_ in _dafny.IntegerRange(807, 152):
                    d_436_v46_: int = compr_24_
                    if ((807) <= (d_436_v46_)) and ((d_436_v46_) < (152)):
                        coll24_ = coll24_.union(_dafny.Set([(d_436_v46_) * (len(_dafny.SeqWithoutIsStrInference([p0])))]))
                return _dafny.Set(coll24_)
            d_435_v47_ = _dafny.SeqWithoutIsStrInference([(iife36_()
            ) - (d_433_v45_)])
            (globalState).f5 = len((d_435_v47_)[default__.safeIndex(len(d_378_v0_), len(d_435_v47_))])
        (globalState).f0 = (p0 if p2 else p0)
        d_437_v48_: _dafny.Array
        nw78_ = _dafny.Array(int(0), 14)
        d_437_v48_ = nw78_
        index53_ = default__.safeIndex(294, (d_437_v48_).length(0))
        (d_437_v48_)[index53_] = p1
        index54_ = default__.safeIndex(294, (d_437_v48_).length(0))
        (d_437_v48_)[index54_] = p1
        r0 = p1
        d_438_v49_: _dafny.Map
        d_438_v49_ = _dafny.Map({p2: (d_437_v48_)[default__.safeIndex(294, (d_437_v48_).length(0))]})
        d_439_v50_: _dafny.Seq
        d_439_v50_ = _dafny.SeqWithoutIsStrInference([321])
        d_440_v51_: _dafny.Map
        d_440_v51_ = _dafny.Map({(d_439_v50_)[default__.safeIndex(p1, len(d_439_v50_))]: False})
        r0 = len(default__.fm2(p0, True, p2, (d_438_v49_) | (_dafny.Map({((d_440_v51_)[(d_437_v48_)[default__.safeIndex(294, (d_437_v48_).length(0))]] if ((d_437_v48_)[default__.safeIndex(294, (d_437_v48_).length(0))]) in (d_440_v51_) else p0): p1})), globalState))
        r1 = p0
        r2 = not((p2) or ((p2 if p0 else p0)))
        return r0, r1, r2

    def m7(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: int = int(0)
        d_441_v0_: str
        d_441_v0_ = _dafny.CodePoint('u')
        d_442_v1_: D0
        d_442_v1_ = D0_DC1(p1)
        d_443_v2_: D2
        d_443_v2_ = D2_DC6((d_442_v1_).cf1)
        d_444_v3_: _dafny.Map
        d_444_v3_ = _dafny.Map({p3: False})
        d_445_v4_: _dafny.Map
        d_445_v4_ = _dafny.Map({p1: ((d_444_v3_)[p3] if (p3) in (d_444_v3_) else p1)})
        d_446_v5_: _dafny.Map
        d_446_v5_ = _dafny.Map({d_445_v4_: p3})
        d_447_v6_: _dafny.MultiSet
        d_447_v6_ = _dafny.MultiSet([(0) - (p3)])
        d_448_v7_: _dafny.Array
        nw79_ = _dafny.Array(None, 17)
        nw79_[int(0)] = p0
        nw79_[int(1)] = p1
        nw79_[int(2)] = p1
        nw79_[int(3)] = p0
        nw79_[int(4)] = p0
        nw79_[int(5)] = (_dafny.Set({d_441_v0_})).ispropersubset(_dafny.Set({d_441_v0_, d_441_v0_}))
        nw79_[int(6)] = True
        nw79_[int(7)] = p0
        nw79_[int(8)] = p1
        nw79_[int(9)] = (p0) and ((d_443_v2_).cf6)
        nw79_[int(10)] = (d_445_v4_) in (d_446_v5_)
        nw79_[int(11)] = False
        nw79_[int(12)] = ((d_447_v6_).set(p3, default__.abs(p3))).isdisjoint(_dafny.MultiSet([p3]))
        nw79_[int(13)] = p1
        nw79_[int(14)] = p1
        nw79_[int(15)] = p0
        nw79_[int(16)] = p0
        d_448_v7_ = nw79_
        index55_ = default__.safeIndex(160, (d_448_v7_).length(0))
        (d_448_v7_)[index55_] = p0
        d_449_v8_: D3
        d_449_v8_ = D3_DC9(p3)
        pat_let_tv11_ = p1
        pat_let_tv12_ = p0
        pat_let_tv13_ = p1
        index56_ = default__.safeIndex(160, (d_448_v7_).length(0))
        def lambda29_(source7_):
            if source7_.is_DC9:
                d_450___mcc_h0_ = source7_.cf11
                d_451_cf11_ = d_450___mcc_h0_
                return pat_let_tv11_
            elif source7_.is_DC8:
                d_452___mcc_h1_ = source7_.cf10
                d_453_cf10_ = d_452___mcc_h1_
                return pat_let_tv12_
            elif True:
                d_454___mcc_h2_ = source7_.cf12
                d_455_cf12_ = d_454___mcc_h2_
                return pat_let_tv13_

        (d_448_v7_)[index56_] = lambda29_(d_449_v8_)
        d_456_v9_: D4
        d_456_v9_ = D4_DC11(d_448_v7_)
        source8_ = d_456_v9_
        if source8_.is_DC12:
            d_457___mcc_h3_ = source8_.cf14
            d_458___mcc_h4_ = source8_.cf15
            d_459___mcc_h5_ = source8_.cf16
            d_460___mcc_h6_ = source8_.cf17
            d_461___mcc_h7_ = source8_.cf18
            d_462_cf18_ = d_461___mcc_h7_
            d_463_cf17_ = d_460___mcc_h6_
            d_464_cf16_ = d_459___mcc_h5_
            d_465_cf15_ = d_458___mcc_h4_
            d_466_cf14_ = d_457___mcc_h3_
            d_464_cf16_ = ((d_464_cf16_) + (d_464_cf16_)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "taufebrr"))) + (d_464_cf16_))
            r3 = d_466_cf14_
            d_467_v10_: _dafny.Set
            d_467_v10_ = _dafny.Set({d_441_v0_})
            d_468_v13_: _dafny.Map
            d_468_v13_ = _dafny.Map({d_441_v0_: d_466_cf14_})
            d_469_v15_: _dafny.MultiSet
            d_469_v15_ = _dafny.MultiSet([d_441_v0_, d_441_v0_])
            d_470_v17_: _dafny.Map
            d_470_v17_ = _dafny.Map({d_441_v0_: p1})
            d_471_v19_: _dafny.Array
            nw80_ = _dafny.Array(None, 22)
            def iife37_():
                def iife39_():
                    coll27_ = _dafny.Set()
                    compr_27_: str
                    for compr_27_ in (d_467_v10_).Elements:
                        d_474_v11_: str = compr_27_
                        if (d_474_v11_) in (d_467_v10_):
                            coll27_ = coll27_.union(_dafny.Set([d_474_v11_]))
                    return _dafny.Set(coll27_)
                coll25_ = _dafny.Set()
                def iife38_():
                    coll26_ = _dafny.Set()
                    compr_26_: str
                    for compr_26_ in (d_467_v10_).Elements:
                        d_472_v11_: str = compr_26_
                        if (d_472_v11_) in (d_467_v10_):
                            coll26_ = coll26_.union(_dafny.Set([d_472_v11_]))
                    return _dafny.Set(coll26_)
                compr_25_: str
                for compr_25_ in (iife38_()
                ).Elements:
                    d_473_v12_: str = compr_25_
                    if (d_473_v12_) in (iife39_()
                    ):
                        coll25_ = coll25_.union(_dafny.Set([d_473_v12_]))
                return _dafny.Set(coll25_)
            nw80_[int(0)] = (iife37_()
            ) | (d_467_v10_)
            nw80_[int(1)] = d_467_v10_
            nw80_[int(2)] = (d_467_v10_).intersection(default__.fm28((d_447_v6_).cardinality, d_463_cf17_, not(p1), globalState))
            nw80_[int(3)] = _dafny.Set({d_441_v0_, default__.fm29(globalState), d_441_v0_, d_441_v0_})
            nw80_[int(4)] = d_467_v10_
            nw80_[int(5)] = (_dafny.Set({d_441_v0_, d_441_v0_})).intersection(d_467_v10_)
            def iife40_():
                coll28_ = _dafny.Set()
                compr_28_: str
                for compr_28_ in (d_468_v13_).keys.Elements:
                    d_475_v14_: str = compr_28_
                    if (d_475_v14_) in (d_468_v13_):
                        coll28_ = coll28_.union(_dafny.Set([d_475_v14_]))
                return _dafny.Set(coll28_)
            nw80_[int(6)] = (iife40_()
             if True else _dafny.Set({d_441_v0_}))
            nw80_[int(7)] = d_467_v10_
            nw80_[int(8)] = d_467_v10_
            nw80_[int(9)] = (d_467_v10_).intersection(d_467_v10_)
            nw80_[int(10)] = d_467_v10_
            def iife41_():
                coll29_ = _dafny.Set()
                compr_29_: str
                for compr_29_ in (d_469_v15_).Elements:
                    d_476_v16_: str = compr_29_
                    if (d_476_v16_) in (d_469_v15_):
                        coll29_ = coll29_.union(_dafny.Set([d_476_v16_]))
                return _dafny.Set(coll29_)
            nw80_[int(11)] = (iife41_()
            ) - (d_467_v10_)
            nw80_[int(12)] = d_467_v10_
            nw80_[int(13)] = d_467_v10_
            nw80_[int(14)] = d_467_v10_
            nw80_[int(15)] = d_467_v10_
            nw80_[int(16)] = d_467_v10_
            nw80_[int(17)] = d_467_v10_
            def iife42_():
                coll30_ = _dafny.Set()
                compr_30_: str
                for compr_30_ in (d_470_v17_).keys.Elements:
                    d_477_v18_: str = compr_30_
                    if (d_477_v18_) in (d_470_v17_):
                        coll30_ = coll30_.union(_dafny.Set([d_477_v18_]))
                return _dafny.Set(coll30_)
            nw80_[int(18)] = (d_467_v10_).intersection(iife42_()
            )
            nw80_[int(19)] = d_467_v10_
            nw80_[int(20)] = _dafny.Set({d_441_v0_})
            nw80_[int(21)] = _dafny.Set({d_441_v0_})
            d_471_v19_ = nw80_
            index57_ = default__.safeIndex(275, (d_471_v19_).length(0))
            (d_471_v19_)[index57_] = default__.fm28(d_463_cf17_, len(d_464_cf16_), (d_448_v7_)[default__.safeIndex(160, (d_448_v7_).length(0))], globalState)
            index58_ = default__.safeIndex(275, (d_471_v19_).length(0))
            (d_471_v19_)[index58_] = d_467_v10_
            r3 = 340
        elif source8_.is_DC13:
            d_478___mcc_h8_ = source8_.cf19
            d_479___mcc_h9_ = source8_.cf20
            d_480_cf20_ = d_479___mcc_h9_
            d_481_cf19_ = d_478___mcc_h8_
            d_482_v20_: C0
            nw81_ = C0()
            nw81_.ctor__((p2).issubset(p2), d_448_v7_)
            d_482_v20_ = nw81_
            d_483_v21_: _dafny.Array
            nw82_ = _dafny.Array(int(0), 26)
            d_483_v21_ = nw82_
            d_483_v21_ = d_483_v21_
            d_481_cf19_ = d_481_cf19_
            index59_ = default__.safeIndex(371, (d_483_v21_).length(0))
            (d_483_v21_)[index59_] = p3
            index60_ = default__.safeIndex(371, (d_483_v21_).length(0))
            (d_483_v21_)[index60_] = d_481_cf19_
        elif source8_.is_DC14:
            d_484_v22_: _dafny.Seq
            d_484_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "guejok"))
            d_485_v23_: _dafny.Array
            nw83_ = _dafny.Array(int(0), 23)
            d_485_v23_ = nw83_
            d_486_v24_: _dafny.Map
            d_486_v24_ = _dafny.Map({p1: d_485_v23_})
            d_487_v25_: D4
            d_487_v25_ = D4_DC12(default__.fm0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isuftd"))), p3, p3, True, globalState), p3, d_484_v22_, (0) - (len((d_486_v24_).set(True, d_485_v23_))), 612)
            r3 = len((d_487_v25_).cf16)
            d_488_v26_: _dafny.Array
            def lambda30_(d_489_v0_):
                def lambda31_(d_490_i0_):
                    return (_dafny.SeqWithoutIsStrInference([d_489_v0_ for d_491_i1_ in range(default__.abs(-92))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elvverckl")))

                return lambda31_

            init14_ = lambda30_(d_441_v0_)
            nw84_ = _dafny.Array(None, 24)
            for i0_14_ in range(nw84_.length(0)):
                nw84_[i0_14_] = init14_(i0_14_)
            d_488_v26_ = nw84_
            d_492_v27_: _dafny.Seq
            d_492_v27_ = _dafny.SeqWithoutIsStrInference([d_484_v22_, d_484_v22_, d_484_v22_])
            d_493_v28_: _dafny.Seq
            d_493_v28_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_494_v29_: D8
            d_494_v29_ = D8_DC21(p0, p1, d_493_v28_, True, p3)
            d_495_v30_: _dafny.Set
            d_495_v30_ = _dafny.Set({(d_493_v28_)[default__.safeIndex(len((d_494_v29_).cf28), len(d_493_v28_))]})
            index61_ = default__.safeIndex(380, (d_488_v26_).length(0))
            (d_488_v26_)[index61_] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_496_i2_ in range(default__.abs(638))])) + ((d_492_v27_)[default__.safeIndex(len(d_495_v30_), len(d_492_v27_))])
            index62_ = default__.safeIndex(380, (d_488_v26_).length(0))
            (d_488_v26_)[index62_] = d_484_v22_
            d_497_v31_: D0
            d_497_v31_ = D0_DC0(d_444_v3_)
            d_498_v32_: _dafny.Seq
            d_498_v32_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, p0]))])
            d_499_v33_: C0
            nw85_ = C0()
            nw85_.ctor__((self).fm4(d_497_v31_, D0_DC0(_dafny.Map({728: (d_448_v7_)[default__.safeIndex(160, (d_448_v7_).length(0))]})), (d_498_v32_)[default__.safeIndex(811, len(d_498_v32_))], globalState), d_448_v7_)
            d_499_v33_ = nw85_
            d_500_v34_: C0
            nw86_ = C0()
            nw86_.ctor__((d_443_v2_).cf6, (d_499_v33_).f11)
            d_500_v34_ = nw86_
        elif True:
            d_501___mcc_h10_ = source8_.cf13
            d_502_cf13_ = d_501___mcc_h10_
            (globalState).f5 = p3
            d_503_v35_: C0
            nw87_ = C0()
            nw87_.ctor__(p0, d_502_cf13_)
            d_503_v35_ = nw87_
            d_504_v36_: _dafny.Array
            def lambda32_(d_505_p3_):
                def lambda33_(d_506_i3_):
                    return (d_506_i3_) - (d_505_p3_)

                return lambda33_

            init15_ = lambda32_(p3)
            nw88_ = _dafny.Array(None, 12)
            for i0_15_ in range(nw88_.length(0)):
                nw88_[i0_15_] = init15_(i0_15_)
            d_504_v36_ = nw88_
            d_507_v37_: _dafny.Seq
            d_507_v37_ = _dafny.SeqWithoutIsStrInference([d_504_v36_])
            d_508_v38_: D10
            d_508_v38_ = D10_DC26(d_507_v37_)
            d_507_v37_ = ((d_508_v38_).cf37) + (d_507_v37_)
            nw89_ = C0()
            nw89_.ctor__((p3) >= (p3), (d_503_v35_).f11)
            d_503_v35_ = nw89_
        r3 = default__.safeDivisionInt(p3, p3)
        d_509_v39_: _dafny.Array
        nw90_ = _dafny.Array(int(0), 20)
        d_509_v39_ = nw90_
        d_510_v40_: _dafny.Array
        d_510_v40_ = d_509_v39_
        source9_ = d_510_v40_
        d_511___mcc_h11_ = source9_
        d_512_cf21_ = d_511___mcc_h11_
        d_513_v41_: _dafny.MultiSet
        d_513_v41_ = _dafny.MultiSet([d_441_v0_])
        d_513_v41_ = ((d_513_v41_) | (d_513_v41_)) | ((d_513_v41_).set(d_441_v0_, default__.abs(p3)))
        (globalState).f5 = p3
        d_514_v43_: _dafny.Seq
        d_514_v43_ = _dafny.SeqWithoutIsStrInference([p3])
        d_515_v44_: _dafny.Seq
        def iife43_():
            coll31_ = _dafny.Map()
            compr_31_: int
            for compr_31_ in (d_514_v43_).Elements:
                d_516_v42_: int = compr_31_
                if (d_516_v42_) in (d_514_v43_):
                    coll31_[default__.safeDivisionInt(d_516_v42_, -970)] = p1
            return _dafny.Map(coll31_)
        d_515_v44_ = _dafny.SeqWithoutIsStrInference([(d_444_v3_) | (iife43_()
        ), _dafny.Map({p3: (d_448_v7_)[default__.safeIndex(160, (d_448_v7_).length(0))]}), d_444_v3_])
        d_515_v44_ = ((d_515_v44_) + (d_515_v44_)).set(default__.safeIndex(default__.fm0(p3, p3, p3, (d_448_v7_)[default__.safeIndex(160, (d_448_v7_).length(0))], globalState), len((d_515_v44_) + (d_515_v44_))), default__.fm22(p0, globalState))
        d_517_v45_: _dafny.Array
        def lambda34_(d_518_i4_):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipgn"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhyhwmg")))

        init16_ = lambda34_
        nw91_ = _dafny.Array(None, 9)
        for i0_16_ in range(nw91_.length(0)):
            nw91_[i0_16_] = init16_(i0_16_)
        d_517_v45_ = nw91_
        d_519_v46_: _dafny.Seq
        d_519_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knmngv"))
        index63_ = default__.safeIndex(544, (d_517_v45_).length(0))
        (d_517_v45_)[index63_] = d_519_v46_
        index64_ = default__.safeIndex(544, (d_517_v45_).length(0))
        (d_517_v45_)[index64_] = default__.fm27(globalState)
        d_520_v47_: _dafny.Array
        nw92_ = _dafny.Array(_dafny.CodePoint('D'), 14)
        d_520_v47_ = nw92_
        index65_ = default__.safeIndex(153, (d_520_v47_).length(0))
        (d_520_v47_)[index65_] = d_441_v0_
        index66_ = default__.safeIndex(153, (d_520_v47_).length(0))
        (d_520_v47_)[index66_] = d_441_v0_
        d_521_v48_: _dafny.Map
        d_521_v48_ = _dafny.Map({p3: -115})
        d_522_v49_: D0
        d_522_v49_ = D0_DC0(d_444_v3_)
        rhs80_ = (d_521_v48_ if p0 else d_521_v48_)
        rhs81_ = ((d_445_v4_)[not((p1) == (p1))] if (not((p1) == (p1))) in (d_445_v4_) else (self).fm4(d_522_v49_, d_522_v49_, p2, globalState))
        rhs82_ = d_448_v7_
        d_521_v48_ = rhs80_
        r1 = rhs81_
        d_448_v7_ = rhs82_
        d_523_v50_: _dafny.Map
        d_523_v50_ = _dafny.Map({(d_448_v7_)[default__.safeIndex(160, (d_448_v7_).length(0))]: -467})
        r0 = (d_523_v50_).set(p0, p3)
        r1 = (d_448_v7_)[default__.safeIndex(160, (d_448_v7_).length(0))]
        d_524_v51_: _dafny.Array
        nw93_ = _dafny.Array(_dafny.Set({}), 19)
        d_524_v51_ = nw93_
        d_525_v52_: _dafny.Seq
        d_525_v52_ = _dafny.SeqWithoutIsStrInference([d_524_v51_, d_524_v51_, d_524_v51_])
        r2 = (d_525_v52_)[default__.safeIndex(p3, len(d_525_v52_))]
        r3 = p3
        return r0, r1, r2, r3


class C2(T0):
    def  __init__(self):
        self._f13: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f13):
        (self)._f13 = f13

    def fm4(self, p0, p1, p2, globalState):
        return (self).f13

    def fm18(self, p0, p1, p2, globalState):
        return 65

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        r3: bool = False
        d_526_v0_: _dafny.Seq
        d_526_v0_ = _dafny.SeqWithoutIsStrInference([(self).f13, (self).f13])
        r2 = not(((d_526_v0_) + (d_526_v0_)) == ((d_526_v0_) + (d_526_v0_)))
        d_527_v1_: str
        d_527_v1_ = _dafny.CodePoint('j')
        d_528_v2_: _dafny.Map
        d_528_v2_ = _dafny.Map({p3: d_527_v1_})
        d_529_v3_: _dafny.Seq
        d_529_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnmkw"))
        d_530_v4_: D3
        d_530_v4_ = D3_DC8(d_529_v3_)
        d_531_v5_: D3
        d_531_v5_ = D3_DC10(d_530_v4_)
        d_532_v7_: _dafny.Map
        def iife44_():
            coll32_ = _dafny.Map()
            compr_32_: int
            for compr_32_ in _dafny.IntegerRange(225, 316):
                d_533_v6_: int = compr_32_
                if ((225) <= (d_533_v6_)) and ((d_533_v6_) < (316)):
                    coll32_[(d_533_v6_) * (p2)] = d_527_v1_
            return _dafny.Map(coll32_)
        d_532_v7_ = _dafny.Map({d_531_v5_: iife44_()
        })
        d_534_v9_: _dafny.Set
        d_534_v9_ = _dafny.Set({p0, p2})
        d_535_v10_: _dafny.Map
        def iife45_():
            coll33_ = _dafny.Map()
            compr_33_: int
            for compr_33_ in (d_534_v9_).Elements:
                d_536_v8_: int = compr_33_
                if (d_536_v8_) in (d_534_v9_):
                    coll33_[default__.safeDivisionInt(d_536_v8_, p2)] = d_527_v1_
            return _dafny.Map(coll33_)
        d_535_v10_ = _dafny.Map({(self).f13: iife45_()
        })
        d_537_v11_: _dafny.Set
        d_537_v11_ = _dafny.Set({(self).f13, (self).f13})
        d_538_v12_: _dafny.Map
        d_538_v12_ = _dafny.Map({(self).f13: len(d_537_v11_)})
        d_539_v14_: _dafny.Seq
        d_539_v14_ = _dafny.SeqWithoutIsStrInference([d_528_v2_])
        d_540_v15_: _dafny.Seq
        d_540_v15_ = _dafny.SeqWithoutIsStrInference([p0, (0) - (len(d_529_v3_))])
        d_541_v17_: _dafny.Array
        nw94_ = _dafny.Array(None, 28)
        nw94_[int(0)] = (d_528_v2_).set(p2, d_527_v1_)
        nw94_[int(1)] = (d_528_v2_) | (_dafny.Map({952: d_527_v1_}))
        nw94_[int(2)] = ((d_528_v2_).set(p0, d_527_v1_)) | (d_528_v2_)
        nw94_[int(3)] = (d_528_v2_) | (d_528_v2_)
        nw94_[int(4)] = (((d_532_v7_)[d_531_v5_] if (d_531_v5_) in (d_532_v7_) else d_528_v2_)) | (d_528_v2_)
        nw94_[int(5)] = default__.fm19(globalState)
        nw94_[int(6)] = d_528_v2_
        nw94_[int(7)] = (d_528_v2_) | (d_528_v2_)
        nw94_[int(8)] = d_528_v2_
        nw94_[int(9)] = (d_528_v2_).set(p2, _dafny.CodePoint('d'))
        nw94_[int(10)] = (_dafny.Map({(0) - (p2): d_527_v1_}) if (self).f13 else d_528_v2_)
        nw94_[int(11)] = d_528_v2_
        nw94_[int(12)] = _dafny.Map({p2: (d_529_v3_)[default__.safeIndex(p0, len(d_529_v3_))]})
        nw94_[int(13)] = (_dafny.Map({len(_dafny.Map({(self).f13: (self).f13})): d_527_v1_})).set(len(_dafny.SeqWithoutIsStrInference([d_527_v1_ for d_542_i1_ in range(default__.abs(617))])), d_527_v1_)
        nw94_[int(14)] = _dafny.Map({p2: d_527_v1_})
        nw94_[int(15)] = d_528_v2_
        nw94_[int(16)] = (d_528_v2_).set(p3, d_527_v1_)
        nw94_[int(17)] = d_528_v2_
        nw94_[int(18)] = d_528_v2_
        nw94_[int(19)] = d_528_v2_
        nw94_[int(20)] = ((d_535_v10_)[not((self).f13)] if (not((self).f13)) in (d_535_v10_) else d_528_v2_)
        nw94_[int(21)] = d_528_v2_
        nw94_[int(22)] = d_528_v2_
        nw94_[int(23)] = d_528_v2_
        nw94_[int(24)] = _dafny.Map({p3: default__.fm20((0) - (len(d_538_v12_)), (self).f13, p3, globalState)})
        def iife46_():
            coll34_ = _dafny.Map()
            compr_34_: int
            for compr_34_ in (d_534_v9_).Elements:
                d_543_v13_: int = compr_34_
                if (d_543_v13_) in (d_534_v9_):
                    coll34_[default__.safeDivisionInt(d_543_v13_, p3)] = d_527_v1_
            return _dafny.Map(coll34_)
        nw94_[int(25)] = iife46_()
        
        nw94_[int(26)] = (d_539_v14_)[default__.safeIndex((d_540_v15_)[default__.safeIndex(len(d_529_v3_), len(d_540_v15_))], len(d_539_v14_))]
        def iife47_():
            coll35_ = _dafny.Map()
            compr_35_: int
            for compr_35_ in _dafny.IntegerRange(921, -898):
                d_544_v16_: int = compr_35_
                if ((921) <= (d_544_v16_)) and ((d_544_v16_) < (-898)):
                    coll35_[(d_544_v16_) * (p2)] = _dafny.CodePoint('s')
            return _dafny.Map(coll35_)
        nw94_[int(27)] = (iife47_()
        ).set(p2, d_527_v1_)
        d_541_v17_ = nw94_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_541_v17_).length(0)):
            d_545_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_545_i0_)) and ((d_545_i0_) < ((d_541_v17_).length(0)))):
                (d_541_v17_)[(d_545_i0_)] = (d_528_v2_) | ((d_528_v2_).set(74, d_527_v1_))
        d_546_i2_: int
        d_546_i2_ = 0
        with _dafny.label("3"):
            while (self).f13:
                with _dafny.c_label("3"):
                    if (d_546_i2_) >= (100):
                        raise _dafny.Break("3")
                    d_546_i2_ = (d_546_i2_) + (1)
                    (globalState).f5 = (0) - ((65) * (p0))
                    r1 = (default__.safeModuloInt(len(_dafny.Map({p3: (self).f13})), p2)) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mow")))) + ((0) - (p3)))
                    if (self).f13:
                        d_547_v18_: _dafny.Array
                        def lambda35_(d_548_p2_):
                            def lambda36_(d_549_i3_):
                                return (d_549_i3_) + (d_548_p2_)

                            return lambda36_

                        init17_ = lambda35_(p2)
                        nw95_ = _dafny.Array(None, 21)
                        for i0_17_ in range(nw95_.length(0)):
                            nw95_[i0_17_] = init17_(i0_17_)
                        d_547_v18_ = nw95_
                        index67_ = default__.safeIndex(306, (d_547_v18_).length(0))
                        (d_547_v18_)[index67_] = (p0) + (p3)
                        index68_ = default__.safeIndex(306, (d_547_v18_).length(0))
                        (d_547_v18_)[index68_] = p0
                        d_550_v19_: _dafny.Array
                        nw96_ = _dafny.Array(None, 7)
                        nw96_[int(0)] = (self).f13
                        nw96_[int(1)] = False
                        nw96_[int(2)] = ((d_526_v0_)[default__.safeIndex(len(d_529_v3_), len(d_526_v0_))] if (self).f13 else (self).f13)
                        nw96_[int(3)] = (self).f13
                        nw96_[int(4)] = False
                        nw96_[int(5)] = (self).f13
                        nw96_[int(6)] = (self).f13
                        d_550_v19_ = nw96_
                        d_550_v19_ = d_550_v19_
                        d_551_v20_: T1
                        nw97_ = C0()
                        nw97_.ctor__(True, d_550_v19_)
                        d_551_v20_ = nw97_
                        d_552_v21_: _dafny.Set
                        d_552_v21_ = _dafny.Set({D2_DC5(d_551_v20_)})
                        d_553_v22_: D2
                        d_553_v22_ = D2_DC5(d_551_v20_)
                        pat_let_tv14_ = d_551_v20_
                        def iife48_(_pat_let6_0):
                            def iife49_(d_554_dt__update__tmp_h0_):
                                def iife50_(_pat_let7_0):
                                    def iife51_(d_555_dt__update_hcf5_h0_):
                                        return D2_DC5(d_555_dt__update_hcf5_h0_)
                                    return iife51_(_pat_let7_0)
                                return iife50_(pat_let_tv14_)
                            return iife49_(_pat_let6_0)
                        d_552_v21_ = (_dafny.Set({iife48_(D2_DC5(d_551_v20_)), d_553_v22_}) if (self).f13 else _dafny.Set({D2_DC5(d_551_v20_), d_553_v22_, d_553_v22_, d_553_v22_, d_553_v22_}))
                        d_538_v12_ = (d_538_v12_).set(True, default__.safeDivisionInt((d_547_v18_)[default__.safeIndex(306, (d_547_v18_).length(0))], p2))
                        d_556_v23_: D2
                        d_556_v23_ = D2_DC6((self).f13)
                        index69_ = default__.safeIndex(306, (d_547_v18_).length(0))
                        rhs83_ = (p3) * ((self).fm18(d_556_v23_, (0) - (len(default__.fm21(globalState))), p3, globalState))
                        rhs84_ = (self).f13
                        rhs85_ = d_547_v18_
                        lhs58_ = d_547_v18_
                        lhs59_ = default__.safeIndex(306, (d_547_v18_).length(0))
                        lhs60_ = globalState
                        lhs58_[lhs59_] = rhs83_
                        lhs60_.f2 = rhs84_
                        d_547_v18_ = rhs85_
                    elif True:
                        d_557_v24_: D2
                        d_557_v24_ = D2_DC6((self).f13)
                        d_558_v25_: _dafny.Array
                        nw98_ = _dafny.Array(None, 13)
                        nw98_[int(0)] = (self).f13
                        nw98_[int(1)] = (self).f13
                        nw98_[int(2)] = (default__.fm1((self).f13, globalState)) or ((self).f13)
                        nw98_[int(3)] = ((self).f13 if (self).f13 else (self).f13)
                        nw98_[int(4)] = False
                        nw98_[int(5)] = (self).f13
                        nw98_[int(6)] = (self).f13
                        nw98_[int(7)] = (self).f13
                        nw98_[int(8)] = (self).f13
                        nw98_[int(9)] = (self).f13
                        nw98_[int(10)] = (d_557_v24_).cf6
                        nw98_[int(11)] = (self).f13
                        nw98_[int(12)] = (self).f13
                        d_558_v25_ = nw98_
                        index70_ = default__.safeIndex(295, (d_558_v25_).length(0))
                        (d_558_v25_)[index70_] = (self).f13
                        index71_ = default__.safeIndex(295, (d_558_v25_).length(0))
                        (d_558_v25_)[index71_] = not((p3) != (p3))
                        d_559_v26_: _dafny.Array
                        def lambda37_(d_560_v1_):
                            def lambda38_(d_561_i4_):
                                return d_560_v1_

                            return lambda38_

                        init18_ = lambda37_(d_527_v1_)
                        nw99_ = _dafny.Array(None, 7)
                        for i0_18_ in range(nw99_.length(0)):
                            nw99_[i0_18_] = init18_(i0_18_)
                        d_559_v26_ = nw99_
                        index72_ = default__.safeIndex(167, (d_559_v26_).length(0))
                        (d_559_v26_)[index72_] = (d_527_v1_ if (self).f13 else d_527_v1_)
                        index73_ = default__.safeIndex(295, (d_558_v25_).length(0))
                        index74_ = default__.safeIndex(167, (d_559_v26_).length(0))
                        rhs86_ = (d_558_v25_)[default__.safeIndex(295, (d_558_v25_).length(0))]
                        rhs87_ = d_527_v1_
                        lhs61_ = d_558_v25_
                        lhs62_ = default__.safeIndex(295, (d_558_v25_).length(0))
                        lhs63_ = d_559_v26_
                        lhs64_ = default__.safeIndex(167, (d_559_v26_).length(0))
                        lhs61_[lhs62_] = rhs86_
                        lhs63_[lhs64_] = rhs87_
                        (globalState).f2 = (self).f13
                        (globalState).f5 = p3
                        d_562_v27_: D4
                        d_562_v27_ = D4_DC12(p0, p3, d_529_v3_, p3, len(d_526_v0_))
                        d_563_v28_: _dafny.Map
                        d_563_v28_ = _dafny.Map({p0: d_562_v27_})
                        d_563_v28_ = (d_563_v28_).set(p2, d_562_v27_)
                    d_564_v29_: T0
                    nw100_ = C1()
                    nw100_.ctor__()
                    d_564_v29_ = nw100_
                    d_565_v30_: _dafny.Array
                    nw101_ = _dafny.Array(None, 7)
                    nw101_[int(0)] = d_564_v29_
                    nw101_[int(1)] = d_564_v29_
                    nw101_[int(2)] = d_564_v29_
                    nw101_[int(3)] = d_564_v29_
                    nw101_[int(4)] = d_564_v29_
                    nw101_[int(5)] = d_564_v29_
                    nw101_[int(6)] = d_564_v29_
                    d_565_v30_ = nw101_
                    d_566_v31_: D11
                    d_566_v31_ = D11_DC29(d_564_v29_)
                    index75_ = default__.safeIndex(441, (d_565_v30_).length(0))
                    (d_565_v30_)[index75_] = (d_566_v31_).cf43
                    index76_ = default__.safeIndex(441, (d_565_v30_).length(0))
                    (d_565_v30_)[index76_] = d_564_v29_
                    pass
            pass
        d_567_v32_: _dafny.Array
        nw102_ = _dafny.Array(int(0), 10)
        d_567_v32_ = nw102_
        d_567_v32_ = d_567_v32_
        d_568_v33_: T0
        nw103_ = C1()
        nw103_.ctor__()
        d_568_v33_ = nw103_
        d_568_v33_ = d_568_v33_
        d_569_v34_: _dafny.Map
        d_569_v34_ = _dafny.Map({p3: (self).f13})
        d_570_v36_: D2
        d_570_v36_ = D2_DC6((self).f13)
        d_571_v38_: _dafny.Array
        nw104_ = _dafny.Array(None, 17)
        nw104_[int(0)] = d_569_v34_
        def iife52_():
            coll36_ = _dafny.Map()
            compr_36_: int
            for compr_36_ in _dafny.IntegerRange(-111, 985):
                d_572_v35_: int = compr_36_
                if ((-111) <= (d_572_v35_)) and ((d_572_v35_) < (985)):
                    coll36_[(d_572_v35_) - (p0)] = False
            return _dafny.Map(coll36_)
        nw104_[int(1)] = (iife52_()
        ) | (d_569_v34_)
        nw104_[int(2)] = (d_569_v34_) | (d_569_v34_)
        nw104_[int(3)] = _dafny.Map({(self).fm18(d_570_v36_, 510, p2, globalState): (self).f13})
        nw104_[int(4)] = (d_569_v34_).set((0) - (p0), (self).f13)
        def iife53_():
            coll37_ = _dafny.Map()
            compr_37_: int
            for compr_37_ in _dafny.IntegerRange(278, 751):
                d_573_v37_: int = compr_37_
                if ((278) <= (d_573_v37_)) and ((d_573_v37_) < (751)):
                    coll37_[(d_573_v37_) * (p0)] = (self).f13
            return _dafny.Map(coll37_)
        nw104_[int(5)] = (d_569_v34_) | (iife53_()
        )
        nw104_[int(6)] = d_569_v34_
        nw104_[int(7)] = d_569_v34_
        nw104_[int(8)] = d_569_v34_
        nw104_[int(9)] = (d_569_v34_).set(p0, (self).f13)
        nw104_[int(10)] = d_569_v34_
        nw104_[int(11)] = d_569_v34_
        nw104_[int(12)] = d_569_v34_
        nw104_[int(13)] = d_569_v34_
        nw104_[int(14)] = d_569_v34_
        nw104_[int(15)] = (d_569_v34_) | (d_569_v34_)
        nw104_[int(16)] = (d_569_v34_) | (d_569_v34_)
        d_571_v38_ = nw104_
        d_574_v39_: _dafny.Set
        d_574_v39_ = _dafny.Set({(d_529_v3_)[default__.safeIndex(237, len(d_529_v3_))], _dafny.CodePoint('j')})
        d_575_v40_: _dafny.Map
        d_575_v40_ = _dafny.Map({(d_574_v39_) - (d_574_v39_): d_571_v38_})
        rhs88_ = (self).f13
        rhs89_ = d_529_v3_
        rhs90_ = p3
        rhs91_ = ((d_575_v40_)[d_574_v39_] if (d_574_v39_) in (d_575_v40_) else d_571_v38_)
        lhs65_ = globalState
        r3 = rhs88_
        d_529_v3_ = rhs89_
        lhs65_.f5 = rhs90_
        d_571_v38_ = rhs91_
        d_576_v41_: _dafny.Map
        d_576_v41_ = _dafny.Map({p2: d_526_v0_})
        r0 = (757) < (len(d_576_v41_))
        r1 = default__.safeDivisionInt(p0, p0)
        r2 = not(not(False))
        d_577_v42_: D0
        d_577_v42_ = D0_DC0(d_569_v34_)
        d_578_v43_: _dafny.MultiSet
        d_578_v43_ = _dafny.MultiSet([(self).f13])
        d_579_v44_: D8
        d_579_v44_ = D8_DC21((self).f13, (self).f13, d_526_v0_, (d_568_v33_).fm4(d_577_v42_, d_577_v42_, d_578_v43_, globalState), p0)
        r3 = not((d_579_v44_).cf29)
        return r0, r1, r2, r3

    def m4(self, globalState):
        r0: bool = False
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: _dafny.Map = _dafny.Map({})
        d_580_v0_: int
        d_580_v0_ = 945
        hi2_ = d_580_v0_
        for d_581_i0_ in range(d_580_v0_, hi2_):
            d_580_v0_ = d_581_i0_
            d_582_v1_: C1
            nw105_ = C1()
            nw105_.ctor__()
            d_582_v1_ = nw105_
            d_583_v2_: _dafny.Array
            def lambda39_(d_584_i1_):
                return (self).f13

            init19_ = lambda39_
            nw106_ = _dafny.Array(None, 14)
            for i0_19_ in range(nw106_.length(0)):
                nw106_[i0_19_] = init19_(i0_19_)
            d_583_v2_ = nw106_
            d_585_v3_: T1
            nw107_ = C0()
            nw107_.ctor__((self).f13, d_583_v2_)
            d_585_v3_ = nw107_
            d_586_v4_: D2
            d_586_v4_ = D2_DC5(d_585_v3_)
            d_587_v5_: _dafny.Set
            d_587_v5_ = _dafny.Set({d_586_v4_, d_586_v4_})
            if (d_587_v5_).ispropersubset((d_587_v5_) | (d_587_v5_)):
                d_588_v6_: _dafny.Map
                d_588_v6_ = _dafny.Map({(self).f13: d_581_i0_})
                d_589_v7_: D0
                d_589_v7_ = D0_DC0(_dafny.Map({len(d_588_v6_): not((self).f13)}))
                (globalState).f2 = ((self).f13 if (self).f13 else (self).fm4(d_589_v7_, d_589_v7_, _dafny.MultiSet([False]), globalState))
                d_590_v8_: _dafny.Set
                d_590_v8_ = _dafny.Set({d_581_i0_, d_580_v0_, d_581_i0_, d_580_v0_})
                d_591_v9_: _dafny.Seq
                d_591_v9_ = _dafny.SeqWithoutIsStrInference([((0) - (default__.fm0((_dafny.MultiSet([d_581_i0_, d_580_v0_])).cardinality, d_580_v0_, d_580_v0_, (self).f13, globalState))) >= (len(d_590_v8_)), (self).f13, True, (self).f13, (self).f13])
                d_591_v9_ = (d_591_v9_) + (d_591_v9_)
                index77_ = default__.safeIndex(634, (d_583_v2_).length(0))
                (d_583_v2_)[index77_] = (self).f13
                index78_ = default__.safeIndex(634, (d_583_v2_).length(0))
                (d_583_v2_)[index78_] = (self).f13
                def iife54_():
                    coll38_ = _dafny.Set()
                    compr_38_: int
                    for compr_38_ in _dafny.IntegerRange(676, 277):
                        d_592_v10_: int = compr_38_
                        if ((676) <= (d_592_v10_)) and ((d_592_v10_) < (277)):
                            coll38_ = coll38_.union(_dafny.Set([(d_592_v10_) * ((D2_DC7(d_580_v0_, _dafny.Map({_dafny.Map({not(True): (d_583_v2_)[default__.safeIndex(634, (d_583_v2_).length(0))]}): d_581_i0_}), d_580_v0_)).cf7)]))
                    return _dafny.Set(coll38_)
                d_590_v8_ = iife54_()
                
                d_593_v11_: _dafny.Array
                nw108_ = _dafny.Array(int(0), 26)
                d_593_v11_ = nw108_
                index79_ = default__.safeIndex(525, (d_593_v11_).length(0))
                (d_593_v11_)[index79_] = d_580_v0_
                d_594_v12_: _dafny.Map
                d_594_v12_ = _dafny.Map({(d_583_v2_)[default__.safeIndex(634, (d_583_v2_).length(0))]: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_595_i2_ in range(default__.abs(702))])})
                d_596_v13_: _dafny.Seq
                d_596_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
                d_597_v14_: _dafny.Seq
                d_597_v14_ = _dafny.SeqWithoutIsStrInference([d_596_v13_])
                index80_ = default__.safeIndex(634, (d_583_v2_).length(0))
                index81_ = default__.safeIndex(525, (d_593_v11_).length(0))
                rhs92_ = ((d_580_v0_) * (d_580_v0_) if (self).f13 else default__.safeModuloInt(d_580_v0_, len(((d_594_v12_)[False] if (False) in (d_594_v12_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smmjmydnm"))))))
                rhs93_ = (d_597_v14_) != (d_597_v14_)
                rhs94_ = d_581_i0_
                lhs66_ = globalState
                lhs67_ = d_583_v2_
                lhs68_ = default__.safeIndex(634, (d_583_v2_).length(0))
                lhs69_ = d_593_v11_
                lhs70_ = default__.safeIndex(525, (d_593_v11_).length(0))
                lhs66_.f5 = rhs92_
                lhs67_[lhs68_] = rhs93_
                lhs69_[lhs70_] = rhs94_
            elif True:
                d_598_v15_: _dafny.Array
                def lambda40_(d_599_i3_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))

                init20_ = lambda40_
                nw109_ = _dafny.Array(None, 25)
                for i0_20_ in range(nw109_.length(0)):
                    nw109_[i0_20_] = init20_(i0_20_)
                d_598_v15_ = nw109_
                index82_ = default__.safeIndex(815, (d_598_v15_).length(0))
                (d_598_v15_)[index82_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyghxaonv"))
                d_600_v16_: _dafny.Seq
                d_600_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trl"))
                index83_ = default__.safeIndex(815, (d_598_v15_).length(0))
                (d_598_v15_)[index83_] = d_600_v16_
                d_601_v17_: _dafny.Seq
                d_601_v17_ = _dafny.SeqWithoutIsStrInference([(self).f13, (self).f13, (self).f13, False, (self).f13])
                d_602_v18_: _dafny.MultiSet
                d_602_v18_ = _dafny.MultiSet([(self).f13, True])
                r0 = ((d_602_v18_ if (self).f13 else d_602_v18_)).ispropersubset(_dafny.MultiSet(d_601_v17_))
                d_603_v19_: str
                d_603_v19_ = _dafny.CodePoint('w')
                d_604_v20_: _dafny.Map
                d_604_v20_ = _dafny.Map({(self).f13: d_603_v19_})
                d_605_v21_: _dafny.Set
                d_605_v21_ = _dafny.Set({(self).f13})
                d_606_v22_: _dafny.Map
                d_606_v22_ = _dafny.Map({(len(d_604_v20_)) == ((0) - (d_581_i0_)): d_605_v21_})
                d_606_v22_ = d_606_v22_
                d_607_v23_: _dafny.Array
                nw110_ = _dafny.Array(int(0), 28)
                d_607_v23_ = nw110_
                d_608_v24_: D10
                d_608_v24_ = D10_DC26(_dafny.SeqWithoutIsStrInference([d_607_v23_, d_607_v23_]))
                d_608_v24_ = d_608_v24_
                d_580_v0_ = d_580_v0_
            d_609_v25_: _dafny.Array
            nw111_ = _dafny.Array(_dafny.MultiSet({}), 9)
            d_609_v25_ = nw111_
            d_610_v26_: _dafny.MultiSet
            d_610_v26_ = _dafny.MultiSet([(self).f13])
            d_611_v27_: _dafny.Map
            d_611_v27_ = _dafny.Map({(self).f13: d_581_i0_})
            d_612_v28_: str
            d_612_v28_ = _dafny.CodePoint('s')
            index84_ = default__.safeIndex(337, (d_609_v25_).length(0))
            (d_609_v25_)[index84_] = (d_610_v26_).intersection((D10_DC27(True, ((d_611_v27_)[(self).f13] if ((self).f13) in (d_611_v27_) else d_581_i0_), d_612_v28_, d_610_v26_)).cf41)
            d_613_v29_: _dafny.Set
            d_613_v29_ = _dafny.Set({not((self).f13)})
            d_614_v30_: D12
            d_614_v30_ = D12_DC33(d_613_v29_)
            index85_ = default__.safeIndex(337, (d_609_v25_).length(0))
            (d_609_v25_)[index85_] = _dafny.MultiSet([(self).f13, True, True, ((d_614_v30_).cf47).issubset(d_613_v29_)])
        d_615_v31_: _dafny.Array
        nw112_ = _dafny.Array(D2.default()(), 6)
        d_615_v31_ = nw112_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_615_v31_).length(0)):
            d_616_i4_: int = guard_loop_1_
            if (True) and (((0) <= (d_616_i4_)) and ((d_616_i4_) < ((d_615_v31_).length(0)))):
                (d_615_v31_)[(d_616_i4_)] = D2_DC7((0) - (d_580_v0_), _dafny.Map({_dafny.Map({(self).f13: (self).f13}): d_580_v0_}), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wlpnjgxrs"))) if (self).f13 else d_580_v0_))
        rhs95_ = d_580_v0_
        rhs96_ = default__.safeDivisionInt(177, d_580_v0_)
        lhs71_ = globalState
        d_580_v0_ = rhs95_
        lhs71_.f5 = rhs96_
        d_617_v32_: C1
        nw113_ = C1()
        nw113_.ctor__()
        d_617_v32_ = nw113_
        hi3_ = d_580_v0_
        for d_618_i5_ in range(default__.safeDivisionInt(d_580_v0_, d_580_v0_), hi3_):
            d_619_v33_: _dafny.Map
            d_619_v33_ = _dafny.Map({d_580_v0_: (self).f13})
            d_620_v34_: D0
            d_620_v34_ = D0_DC0(d_619_v33_)
            d_621_v35_: _dafny.MultiSet
            d_621_v35_ = _dafny.MultiSet([(self).f13, False])
            d_622_v36_: _dafny.Array
            nw114_ = _dafny.Array(None, 12)
            nw114_[int(0)] = (self).f13
            nw114_[int(1)] = True
            nw114_[int(2)] = (self).f13
            nw114_[int(3)] = (self).f13
            nw114_[int(4)] = (self).f13
            nw114_[int(5)] = (self).f13
            nw114_[int(6)] = (self).f13
            nw114_[int(7)] = (self).f13
            nw114_[int(8)] = (self).f13
            nw114_[int(9)] = (self).f13
            nw114_[int(10)] = (self).f13
            nw114_[int(11)] = False
            d_622_v36_ = nw114_
            d_623_v37_: _dafny.Seq
            d_623_v37_ = _dafny.SeqWithoutIsStrInference([d_622_v36_])
            d_624_v38_: _dafny.Set
            d_624_v38_ = _dafny.Set({d_619_v33_, d_619_v33_})
            d_625_v39_: _dafny.Seq
            d_625_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgk"))
            d_626_v40_: _dafny.Seq
            d_626_v40_ = _dafny.SeqWithoutIsStrInference([(self).f13, (self).f13])
            d_627_v41_: D12
            d_627_v41_ = D12_DC36((self).f13, 75, (self).f13)
            d_628_v42_: C0
            nw115_ = C0()
            nw115_.ctor__((self).f13, d_622_v36_)
            d_628_v42_ = nw115_
            d_629_v43_: _dafny.Seq
            d_629_v43_ = _dafny.SeqWithoutIsStrInference([d_628_v42_])
            d_630_v44_: _dafny.Array
            nw116_ = _dafny.Array(None, 29)
            nw116_[int(0)] = default__.fm1((self).f13, globalState)
            nw116_[int(1)] = False
            nw116_[int(2)] = not((self).f13)
            nw116_[int(3)] = (self).f13
            nw116_[int(4)] = (self).f13
            nw116_[int(5)] = (self).fm4(d_620_v34_, D0_DC0(d_619_v33_), d_621_v35_, globalState)
            nw116_[int(6)] = (d_623_v37_) != (d_623_v37_)
            nw116_[int(7)] = (d_624_v38_).ispropersubset(d_624_v38_)
            nw116_[int(8)] = (self).f13
            nw116_[int(9)] = not((d_625_v39_) <= (d_625_v39_))
            nw116_[int(10)] = (d_621_v35_).ispropersubset(d_621_v35_)
            nw116_[int(11)] = (self).f13
            nw116_[int(12)] = (self).f13
            nw116_[int(13)] = (self).f13
            nw116_[int(14)] = (self).f13
            nw116_[int(15)] = (d_618_i5_) > (len((d_626_v40_).set(default__.safeIndex(d_580_v0_, len(d_626_v40_)), not((self).f13))))
            nw116_[int(16)] = True
            nw116_[int(17)] = not((not(not((self).f13))) or ((d_627_v41_).cf53))
            nw116_[int(18)] = (self).f13
            nw116_[int(19)] = (default__.fm21(globalState)) == (d_625_v39_)
            nw116_[int(20)] = (self).f13
            nw116_[int(21)] = default__.fm1((self).f13, globalState)
            nw116_[int(22)] = not((self).f13)
            nw116_[int(23)] = not((self).f13)
            nw116_[int(24)] = (self).f13
            nw116_[int(25)] = (self).f13
            nw116_[int(26)] = (self).f13
            nw116_[int(27)] = (d_629_v43_) < (d_629_v43_)
            nw116_[int(28)] = ((self).f13 if True else (d_628_v42_).f10)
            d_630_v44_ = nw116_
            d_630_v44_ = (d_628_v42_).f11
            d_631_v45_: _dafny.Array
            nw117_ = _dafny.Array(_dafny.Map({}), 8)
            d_631_v45_ = nw117_
            d_631_v45_ = d_631_v45_
            d_628_v42_ = d_628_v42_
            (globalState).f5 = default__.safeDivisionInt(d_618_i5_, d_618_i5_)
        if (self).f13:
            d_632_v46_: _dafny.Seq
            d_632_v46_ = _dafny.SeqWithoutIsStrInference([d_580_v0_])
            d_633_v47_: _dafny.Seq
            d_633_v47_ = d_632_v46_
            d_633_v47_ = d_633_v47_
            d_634_v48_: _dafny.Array
            nw118_ = _dafny.Array(None, 17)
            nw118_[int(0)] = d_580_v0_
            nw118_[int(1)] = d_580_v0_
            nw118_[int(2)] = 158
            nw118_[int(3)] = (d_580_v0_) * (d_580_v0_)
            nw118_[int(4)] = (d_632_v46_)[default__.safeIndex(d_580_v0_, len(d_632_v46_))]
            nw118_[int(5)] = d_580_v0_
            nw118_[int(6)] = 262
            nw118_[int(7)] = d_580_v0_
            nw118_[int(8)] = d_580_v0_
            nw118_[int(9)] = d_580_v0_
            nw118_[int(10)] = -672
            nw118_[int(11)] = d_580_v0_
            nw118_[int(12)] = (len(d_632_v46_)) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_635_i6_ in range(default__.abs(-606))])))
            nw118_[int(13)] = d_580_v0_
            nw118_[int(14)] = d_580_v0_
            nw118_[int(15)] = d_580_v0_
            nw118_[int(16)] = d_580_v0_
            d_634_v48_ = nw118_
            index86_ = default__.safeIndex(646, (d_634_v48_).length(0))
            (d_634_v48_)[index86_] = (0) - (d_580_v0_)
            index87_ = default__.safeIndex(646, (d_634_v48_).length(0))
            rhs97_ = d_580_v0_
            rhs98_ = (0) - (d_580_v0_)
            lhs72_ = d_634_v48_
            lhs73_ = default__.safeIndex(646, (d_634_v48_).length(0))
            lhs74_ = globalState
            lhs72_[lhs73_] = rhs97_
            lhs74_.f5 = rhs98_
            d_636_v49_: _dafny.Map
            d_636_v49_ = _dafny.Map({d_580_v0_: not(not((self).f13))})
            d_637_v50_: _dafny.Seq
            d_637_v50_ = _dafny.SeqWithoutIsStrInference([(self).f13])
            d_638_v51_: D0
            d_638_v51_ = D0_DC0(d_636_v49_)
            d_639_v52_: _dafny.MultiSet
            d_639_v52_ = _dafny.MultiSet([(self).f13])
            pat_let_tv15_ = d_634_v48_
            pat_let_tv16_ = d_634_v48_
            d_640_v53_: _dafny.Array
            nw119_ = _dafny.Array(None, 22)
            nw119_[int(0)] = (self).f13
            nw119_[int(1)] = (_dafny.SeqWithoutIsStrInference([(self).f13])) < (_dafny.SeqWithoutIsStrInference([default__.fm1((self).f13, globalState)]))
            nw119_[int(2)] = (self).f13
            nw119_[int(3)] = ((d_636_v49_)[d_580_v0_] if (d_580_v0_) in (d_636_v49_) else not((self).f13))
            nw119_[int(4)] = (self).f13
            nw119_[int(5)] = ((d_634_v48_)[default__.safeIndex(646, (d_634_v48_).length(0))]) > (d_580_v0_)
            nw119_[int(6)] = (self).f13
            nw119_[int(7)] = not((self).f13)
            nw119_[int(8)] = not((self).f13)
            nw119_[int(9)] = (self).f13
            nw119_[int(10)] = (self).f13
            nw119_[int(11)] = (self).f13
            nw119_[int(12)] = (self).f13
            nw119_[int(13)] = (self).f13
            nw119_[int(14)] = not(((d_636_v49_)[d_580_v0_] if (d_580_v0_) in (d_636_v49_) else (self).f13))
            nw119_[int(15)] = (d_637_v50_) != (_dafny.SeqWithoutIsStrInference([(self).f13, (self).f13]))
            nw119_[int(16)] = False
            nw119_[int(17)] = (self).f13
            nw119_[int(18)] = (self).f13
            nw119_[int(19)] = (self).f13
            nw119_[int(20)] = not((self).f13)
            def iife55_(_pat_let8_0):
                def iife56_(d_641_dt__update__tmp_h0_):
                    def iife57_(_pat_let9_0):
                        def iife58_(d_642_dt__update_hcf0_h0_):
                            return D0_DC0(d_642_dt__update_hcf0_h0_)
                        return iife58_(_pat_let9_0)
                    return iife57_(_dafny.Map({(pat_let_tv16_)[default__.safeIndex(646, (pat_let_tv15_).length(0))]: (self).f13}))
                return iife56_(_pat_let8_0)
            nw119_[int(21)] = (True if (d_617_v32_).fm4(iife55_(d_638_v51_), d_638_v51_, d_639_v52_, globalState) else (self).f13)
            d_640_v53_ = nw119_
            index88_ = default__.safeIndex(244, (d_640_v53_).length(0))
            (d_640_v53_)[index88_] = (self).f13
            index89_ = default__.safeIndex(244, (d_640_v53_).length(0))
            (d_640_v53_)[index89_] = not((self).f13)
            d_643_v54_: str
            d_643_v54_ = _dafny.CodePoint('p')
            d_643_v54_ = d_643_v54_
            d_644_v55_: _dafny.Array
            nw120_ = _dafny.Array(_dafny.Map({}), 16)
            d_644_v55_ = nw120_
            d_644_v55_ = d_644_v55_
        elif True:
            d_645_v56_: C1
            nw121_ = C1()
            nw121_.ctor__()
            d_645_v56_ = nw121_
            (globalState).f2 = (self).f13
            d_646_v57_: _dafny.Map
            d_646_v57_ = _dafny.Map({d_580_v0_: -311})
            d_647_v58_: _dafny.Set
            d_647_v58_ = _dafny.Set({len(d_646_v57_)})
            if ((default__.fm30(globalState)).intersection(d_647_v58_)).ispropersubset(d_647_v58_):
                d_648_v59_: _dafny.Array
                def lambda41_(d_649_i7_):
                    return False

                init21_ = lambda41_
                nw122_ = _dafny.Array(None, 13)
                for i0_21_ in range(nw122_.length(0)):
                    nw122_[i0_21_] = init21_(i0_21_)
                d_648_v59_ = nw122_
                d_648_v59_ = d_648_v59_
                d_650_v60_: C1
                nw123_ = C1()
                nw123_.ctor__()
                d_650_v60_ = nw123_
                (globalState).f0 = (self).f13
                d_651_v61_: C0
                nw124_ = C0()
                nw124_.ctor__((self).f13, d_648_v59_)
                d_651_v61_ = nw124_
                d_651_v61_ = d_651_v61_
                d_652_v62_: _dafny.Array
                nw125_ = _dafny.Array(int(0), 28)
                d_652_v62_ = nw125_
                d_653_v63_: _dafny.MultiSet
                d_653_v63_ = _dafny.MultiSet([d_652_v62_, d_652_v62_])
                d_654_v64_: _dafny.Map
                d_654_v64_ = _dafny.Map({d_617_v32_: d_653_v63_})
                d_655_v65_: _dafny.Seq
                d_655_v65_ = _dafny.SeqWithoutIsStrInference([d_580_v0_])
                d_656_v66_: _dafny.Array
                d_656_v66_ = d_652_v62_
                d_657_v67_: _dafny.Map
                d_657_v67_ = _dafny.Map({_dafny.Set({(d_651_v61_).f10}): _dafny.CodePoint('v')})
                d_658_v68_: _dafny.Seq
                d_658_v68_ = _dafny.SeqWithoutIsStrInference([d_653_v63_, d_653_v63_, d_653_v63_, ((d_653_v63_).set(d_652_v62_, default__.abs(len(d_657_v67_)))).set(d_652_v62_, default__.abs(d_580_v0_)), d_653_v63_])
                d_659_v69_: D2
                d_659_v69_ = D2_DC6((self).f13)
                d_660_v70_: _dafny.Map
                d_660_v70_ = _dafny.Map({(d_651_v61_).f10: d_580_v0_})
                d_661_v71_: str
                d_661_v71_ = _dafny.CodePoint('q')
                d_662_v72_: _dafny.Map
                d_662_v72_ = _dafny.Map({(self).fm18(d_659_v69_, len(d_647_v58_), len(d_660_v70_), globalState): d_661_v71_})
                d_663_v73_: _dafny.Array
                nw126_ = _dafny.Array(None, 19)
                nw126_[int(0)] = d_653_v63_
                nw126_[int(1)] = _dafny.MultiSet([d_652_v62_])
                nw126_[int(2)] = (((d_654_v64_)[d_645_v56_] if (d_645_v56_) in (d_654_v64_) else _dafny.MultiSet([d_652_v62_]))).set(d_652_v62_, default__.abs(len(d_655_v65_)))
                nw126_[int(3)] = (d_653_v63_).set((d_656_v66_), default__.abs(89))
                nw126_[int(4)] = (d_653_v63_) | (d_653_v63_)
                nw126_[int(5)] = (d_653_v63_) - (d_653_v63_)
                nw126_[int(6)] = _dafny.MultiSet([d_652_v62_, d_652_v62_])
                nw126_[int(7)] = d_653_v63_
                nw126_[int(8)] = (d_658_v68_)[default__.safeIndex(d_580_v0_, len(d_658_v68_))]
                nw126_[int(9)] = (d_653_v63_) - ((d_653_v63_).set(d_652_v62_, default__.abs(len(d_662_v72_))))
                nw126_[int(10)] = (d_653_v63_ if default__.fm1((self).f13, globalState) else d_653_v63_)
                nw126_[int(11)] = d_653_v63_
                nw126_[int(12)] = d_653_v63_
                nw126_[int(13)] = d_653_v63_
                nw126_[int(14)] = d_653_v63_
                nw126_[int(15)] = _dafny.MultiSet([d_652_v62_, d_652_v62_, d_652_v62_])
                nw126_[int(16)] = (d_653_v63_) | (d_653_v63_)
                nw126_[int(17)] = (d_653_v63_).intersection(_dafny.MultiSet([d_652_v62_]))
                nw126_[int(18)] = (d_658_v68_)[default__.safeIndex(d_580_v0_, len(d_658_v68_))]
                d_663_v73_ = nw126_
                index90_ = default__.safeIndex(592, (d_663_v73_).length(0))
                (d_663_v73_)[index90_] = d_653_v63_
                index91_ = default__.safeIndex(592, (d_663_v73_).length(0))
                (d_663_v73_)[index91_] = d_653_v63_
            elif True:
                d_664_v74_: _dafny.Set
                d_664_v74_ = _dafny.Set({d_646_v57_, d_646_v57_})
                d_665_v75_: _dafny.Array
                nw127_ = _dafny.Array(None, 5)
                nw127_[int(0)] = d_580_v0_
                nw127_[int(1)] = (d_580_v0_ if (self).f13 else d_580_v0_)
                nw127_[int(2)] = d_580_v0_
                nw127_[int(3)] = d_580_v0_
                nw127_[int(4)] = len(d_664_v74_)
                d_665_v75_ = nw127_
                index92_ = default__.safeIndex(915, (d_665_v75_).length(0))
                (d_665_v75_)[index92_] = d_580_v0_
                d_666_v76_: _dafny.Seq
                d_666_v76_ = _dafny.SeqWithoutIsStrInference([(self).f13])
                d_667_v77_: _dafny.Seq
                d_667_v77_ = _dafny.SeqWithoutIsStrInference([d_666_v76_])
                index93_ = default__.safeIndex(915, (d_665_v75_).length(0))
                rhs99_ = (0) - (d_580_v0_)
                rhs100_ = d_665_v75_
                rhs101_ = default__.fm30(globalState)
                rhs102_ = (0) - (len((d_667_v77_) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f13]) for d_668_i8_ in range(default__.abs(-310))]))))
                lhs75_ = globalState
                lhs76_ = d_665_v75_
                lhs77_ = default__.safeIndex(915, (d_665_v75_).length(0))
                lhs75_.f5 = rhs99_
                d_665_v75_ = rhs100_
                d_647_v58_ = rhs101_
                lhs76_[lhs77_] = rhs102_
                d_669_v78_: _dafny.Map
                d_669_v78_ = _dafny.Map({(self).f13: (self).f13})
                d_670_v79_: D2
                d_670_v79_ = D2_DC7(d_580_v0_, _dafny.Map({d_669_v78_: d_580_v0_}), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjam")))))
                d_671_v80_: _dafny.MultiSet
                d_671_v80_ = _dafny.MultiSet([(d_670_v79_ if not((self).f13) else d_670_v79_)])
                d_671_v80_ = (d_671_v80_) | (d_671_v80_)
                d_672_v81_: str
                d_672_v81_ = _dafny.CodePoint('p')
                d_672_v81_ = d_672_v81_
                d_673_v82_: C1
                nw128_ = C1()
                nw128_.ctor__()
                d_673_v82_ = nw128_
                (globalState).f0 = (self).f13
            d_674_v83_: _dafny.Array
            nw129_ = _dafny.Array(None, 24)
            nw129_[int(0)] = (self).f13
            nw129_[int(1)] = (self).f13
            nw129_[int(2)] = (self).f13
            nw129_[int(3)] = (self).f13
            nw129_[int(4)] = default__.fm1((self).f13, globalState)
            nw129_[int(5)] = False
            nw129_[int(6)] = (self).f13
            nw129_[int(7)] = not((self).f13)
            nw129_[int(8)] = (self).f13
            nw129_[int(9)] = (self).f13
            nw129_[int(10)] = (self).f13
            nw129_[int(11)] = (self).f13
            nw129_[int(12)] = (self).f13
            nw129_[int(13)] = not((self).f13)
            nw129_[int(14)] = (self).f13
            nw129_[int(15)] = (self).f13
            nw129_[int(16)] = (self).f13
            nw129_[int(17)] = (self).f13
            nw129_[int(18)] = (self).f13
            nw129_[int(19)] = not((self).f13)
            nw129_[int(20)] = (self).f13
            nw129_[int(21)] = (self).f13
            nw129_[int(22)] = (self).f13
            nw129_[int(23)] = (self).f13
            d_674_v83_ = nw129_
            d_675_v84_: C0
            nw130_ = C0()
            nw130_.ctor__((self).f13, d_674_v83_)
            d_675_v84_ = nw130_
            d_676_v85_: _dafny.Seq
            d_676_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pswtnsaun"))
            (globalState).f0 = not(not (False) or ((d_676_v85_) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_677_i9_ in range(default__.abs(-527))]))))
        d_678_v86_: _dafny.Array
        def lambda42_(d_679_i11_):
            return False

        init22_ = lambda42_
        nw131_ = _dafny.Array(None, 11)
        for i0_22_ in range(nw131_.length(0)):
            nw131_[i0_22_] = init22_(i0_22_)
        d_678_v86_ = nw131_
        d_680_v87_: _dafny.Seq
        d_680_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvhxsqh"))
        d_681_v88_: _dafny.Map
        d_681_v88_ = _dafny.Map({d_678_v86_: d_680_v87_})
        d_682_v89_: _dafny.Seq
        d_682_v89_ = _dafny.SeqWithoutIsStrInference([len(d_681_v88_)])
        r0 = (_dafny.SeqWithoutIsStrInference([d_580_v0_ for d_683_i10_ in range(default__.abs(691))])) < (d_682_v89_)
        d_684_v90_: _dafny.MultiSet
        d_684_v90_ = _dafny.MultiSet([507])
        d_685_v91_: _dafny.Map
        d_685_v91_ = _dafny.Map({(self).f13: d_684_v90_})
        r1 = ((d_685_v91_)[(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))) != (len((_dafny.Map({d_580_v0_: not((self).f13)})).set(d_580_v0_, (self).f13)))] if ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))) != (len((_dafny.Map({d_580_v0_: not((self).f13)})).set(d_580_v0_, (self).f13)))) in (d_685_v91_) else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_580_v0_ for d_686_i12_ in range(default__.abs(818))])))
        d_687_v92_: _dafny.Map
        d_687_v92_ = _dafny.Map({d_680_v87_: (self).f13})
        r2 = d_687_v92_
        return r0, r1, r2

    def m5(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        d_688_v0_: _dafny.Array
        def lambda43_(d_689_i0_):
            return (self).f13

        init23_ = lambda43_
        nw132_ = _dafny.Array(None, 10)
        for i0_23_ in range(nw132_.length(0)):
            nw132_[i0_23_] = init23_(i0_23_)
        d_688_v0_ = nw132_
        d_690_v1_: C0
        nw133_ = C0()
        nw133_.ctor__((self).f13, d_688_v0_)
        d_690_v1_ = nw133_
        d_691_v2_: D2
        d_691_v2_ = D2_DC6((d_690_v1_).f10)
        d_692_v3_: _dafny.Map
        d_692_v3_ = _dafny.Map({(self).fm18(d_691_v2_, p1, p1, globalState): not((self).f13)})
        d_693_v4_: _dafny.Set
        d_693_v4_ = _dafny.Set({(self).f13, (d_690_v1_).f10})
        d_692_v3_ = (d_692_v3_).set((len(d_693_v4_)) - (p1), not((self).f13))
        d_694_v5_: str
        d_694_v5_ = _dafny.CodePoint('o')
        d_694_v5_ = default__.fm20(default__.safeDivisionInt(p1, p1), (self).f13, p1, globalState)
        d_695_i1_: int
        d_695_i1_ = 0
        with _dafny.label("4"):
            while False:
                with _dafny.c_label("4"):
                    if (d_695_i1_) >= (100):
                        raise _dafny.Break("4")
                    d_695_i1_ = (d_695_i1_) + (1)
                    d_696_v6_: _dafny.Array
                    def lambda44_(d_697_i2_):
                        return default__.safeModuloInt(d_697_i2_, 349)

                    init24_ = lambda44_
                    nw134_ = _dafny.Array(None, 15)
                    for i0_24_ in range(nw134_.length(0)):
                        nw134_[i0_24_] = init24_(i0_24_)
                    d_696_v6_ = nw134_
                    d_698_v7_: _dafny.Map
                    d_698_v7_ = _dafny.Map({(self).f13: d_696_v6_})
                    d_699_v8_: _dafny.Map
                    d_699_v8_ = _dafny.Map({p1: d_696_v6_})
                    d_700_v9_: _dafny.Array
                    d_700_v9_ = d_696_v6_
                    d_701_v10_: _dafny.Array
                    nw135_ = _dafny.Array(None, 18)
                    nw135_[int(0)] = d_696_v6_
                    nw135_[int(1)] = d_696_v6_
                    nw135_[int(2)] = d_696_v6_
                    nw135_[int(3)] = d_696_v6_
                    nw135_[int(4)] = d_696_v6_
                    nw135_[int(5)] = d_696_v6_
                    nw135_[int(6)] = d_696_v6_
                    nw135_[int(7)] = d_696_v6_
                    nw135_[int(8)] = d_696_v6_
                    nw135_[int(9)] = d_696_v6_
                    nw135_[int(10)] = d_696_v6_
                    nw135_[int(11)] = d_696_v6_
                    nw135_[int(12)] = d_696_v6_
                    nw135_[int(13)] = d_696_v6_
                    nw135_[int(14)] = ((d_698_v7_)[(d_690_v1_).f10] if ((d_690_v1_).f10) in (d_698_v7_) else ((d_699_v8_)[p1] if (p1) in (d_699_v8_) else d_696_v6_))
                    nw135_[int(15)] = d_696_v6_
                    nw135_[int(16)] = (d_700_v9_)
                    nw135_[int(17)] = d_696_v6_
                    d_701_v10_ = nw135_
                    rhs103_ = d_701_v10_
                    rhs104_ = d_688_v0_
                    d_701_v10_ = rhs103_
                    d_688_v0_ = rhs104_
                    d_702_v11_: _dafny.Seq
                    d_702_v11_ = _dafny.SeqWithoutIsStrInference([d_694_v5_, _dafny.CodePoint('y'), d_694_v5_])
                    d_694_v5_ = (d_702_v11_)[default__.safeIndex(940, len(d_702_v11_))]
                    d_703_v12_: _dafny.Seq
                    d_703_v12_ = _dafny.SeqWithoutIsStrInference([(d_690_v1_).f10])
                    (globalState).f0 = ((d_703_v12_)[default__.safeIndex(p1, len(d_703_v12_))] if (self).f13 else ((d_690_v1_).f10) not in (_dafny.MultiSet([(d_690_v1_).f10, (d_690_v1_).f10])))
                    (globalState).f0 = (self).f13
                    pass
            pass
        d_704_v13_: _dafny.Seq
        d_704_v13_ = _dafny.SeqWithoutIsStrInference([(0) - (p1), p1, p1])
        d_705_v14_: _dafny.Map
        d_705_v14_ = _dafny.Map({364: (0) - (p1)})
        d_706_v15_: _dafny.Map
        d_706_v15_ = _dafny.Map({len(d_705_v14_): p1})
        d_707_i3_: int
        d_707_i3_ = 0
        with _dafny.label("5"):
            while not((d_704_v13_) <= ((_dafny.SeqWithoutIsStrInference([599])).set(default__.safeIndex(((d_706_v15_)[820] if (820) in (d_706_v15_) else (d_704_v13_)[default__.safeIndex(p1, len(d_704_v13_))]), len(_dafny.SeqWithoutIsStrInference([599]))), p1))):
                with _dafny.c_label("5"):
                    if (d_707_i3_) >= (100):
                        raise _dafny.Break("5")
                    d_707_i3_ = (d_707_i3_) + (1)
                    d_708_v16_: C1
                    nw136_ = C1()
                    nw136_.ctor__()
                    d_708_v16_ = nw136_
                    d_709_v17_: _dafny.MultiSet
                    d_709_v17_ = _dafny.MultiSet([(d_690_v1_).f10])
                    d_709_v17_ = (d_709_v17_) | (d_709_v17_)
                    (globalState).f5 = (0) - (default__.safeModuloInt(p1, 538))
                    d_710_v18_: _dafny.Map
                    d_710_v18_ = _dafny.Map({not (False) or ((self).f13): (p0)[default__.safeIndex(p1, len(p0))]})
                    d_710_v18_ = d_710_v18_
                    pass
            pass
        d_706_v15_ = (d_706_v15_).set(p1, p1)
        d_711_v19_: _dafny.Map
        d_711_v19_ = _dafny.Map({not((d_690_v1_).f10): 381})
        r0 = default__.safeDivisionInt((p1) - (p1), ((d_711_v19_)[(self).f13] if ((self).f13) in (d_711_v19_) else p1))
        r1 = p1
        d_712_v20_: D12
        d_712_v20_ = D12_DC36((d_690_v1_).f10, p1, False)
        pat_let_tv17_ = d_690_v1_
        pat_let_tv18_ = d_704_v13_
        pat_let_tv19_ = d_704_v13_
        pat_let_tv20_ = d_690_v1_
        pat_let_tv21_ = d_690_v1_
        pat_let_tv22_ = d_690_v1_
        def lambda45_(source10_):
            if source10_.is_DC34:
                d_713___mcc_h0_ = source10_.cf48
                d_714___mcc_h1_ = source10_.cf49
                d_715___mcc_h2_ = source10_.cf50
                d_716_cf50_ = d_715___mcc_h2_
                d_717_cf49_ = d_714___mcc_h1_
                d_718_cf48_ = d_713___mcc_h0_
                return (pat_let_tv17_).f10
            elif source10_.is_DC35:
                return (pat_let_tv18_) <= (pat_let_tv19_)
            elif source10_.is_DC36:
                d_719___mcc_h3_ = source10_.cf51
                d_720___mcc_h4_ = source10_.cf52
                d_721___mcc_h5_ = source10_.cf53
                d_722_cf53_ = d_721___mcc_h5_
                d_723_cf52_ = d_720___mcc_h4_
                d_724_cf51_ = d_719___mcc_h3_
                return (d_723_cf52_) != ((_dafny.MultiSet([D0_DC1((pat_let_tv20_).f10), D0_DC1(d_724_cf51_), D0_DC1(d_724_cf51_)])).cardinality)
            elif source10_.is_DC33:
                d_725___mcc_h6_ = source10_.cf47
                d_726_cf47_ = d_725___mcc_h6_
                return (pat_let_tv21_).f10
            elif True:
                d_727___mcc_h7_ = source10_.cf54
                d_728_cf54_ = d_727___mcc_h7_
                return (pat_let_tv22_).f10

        r2 = lambda45_(d_712_v20_)
        return r0, r1, r2

    @property
    def f13(self):
        return self._f13

class C3(T0, T1):
    def  __init__(self):
        self._f12: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self, f12):
        (self)._f12 = f12

    def fm4(self, p0, p1, p2, globalState):
        return ((self).f12) <= ((self).f12)

    def fm11(self, globalState):
        return ((_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vd")))})) | (_dafny.Map({False: (self).f12}))) | ((_dafny.Map({not(not(not(not(not(not(False)))))): -214})) | (_dafny.Map({True: (self).f12})))

    def fm12(self, p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f12]), _dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: (self).f12})), (self).f12])])) < ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f12, (self).f12, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvi"))), (0) - ((self).f12)]), _dafny.SeqWithoutIsStrInference([(self).f12]), _dafny.SeqWithoutIsStrInference([(self).f12]), _dafny.SeqWithoutIsStrInference([(self).f12, len(_dafny.SeqWithoutIsStrInference([False]))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f12, (self).f12]) for d_729_i0_ in range(default__.abs(-877))])))

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        r3: bool = False
        d_730_v0_: bool
        d_730_v0_ = True
        if not (d_730_v0_) or (d_730_v0_):
            d_731_v1_: _dafny.Array
            def lambda46_(d_732_v0_):
                def lambda47_(d_733_i0_):
                    return (D0_DC1(d_732_v0_)).cf1

                return lambda47_

            init25_ = lambda46_(d_730_v0_)
            nw137_ = _dafny.Array(None, 1)
            for i0_25_ in range(nw137_.length(0)):
                nw137_[i0_25_] = init25_(i0_25_)
            d_731_v1_ = nw137_
            d_731_v1_ = d_731_v1_
            if d_730_v0_:
                d_734_v2_: _dafny.Map
                d_734_v2_ = _dafny.Map({p3: d_730_v0_})
                d_735_v3_: _dafny.Array
                nw138_ = _dafny.Array(None, 6)
                nw138_[int(0)] = p0
                nw138_[int(1)] = (self).f12
                nw138_[int(2)] = len(d_734_v2_)
                nw138_[int(3)] = p2
                nw138_[int(4)] = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqcgspfpy"))))
                nw138_[int(5)] = (self).f12
                d_735_v3_ = nw138_
                d_736_v4_: _dafny.Array
                d_736_v4_ = d_735_v3_
                d_737_v5_: _dafny.Seq
                d_737_v5_ = _dafny.SeqWithoutIsStrInference([d_735_v3_, (d_735_v3_)])
                d_738_v6_: _dafny.Set
                d_738_v6_ = _dafny.Set({d_730_v0_, d_730_v0_})
                d_739_v7_: _dafny.Set
                d_739_v7_ = _dafny.Set({d_735_v3_, d_735_v3_, (d_737_v5_)[default__.safeIndex(len(d_738_v6_), len(d_737_v5_))]})
                d_739_v7_ = d_739_v7_
                d_740_v8_: _dafny.MultiSet
                d_740_v8_ = _dafny.MultiSet([default__.fm1(d_730_v0_, globalState), d_730_v0_, d_730_v0_])
                d_741_v9_: _dafny.Seq
                d_741_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axeend"))
                rhs105_ = ((d_741_v9_) + (d_741_v9_)) < ((d_741_v9_) + (d_741_v9_))
                rhs106_ = d_740_v8_
                r0 = rhs105_
                d_740_v8_ = rhs106_
                d_742_v10_: str
                d_742_v10_ = _dafny.CodePoint('u')
                d_742_v10_ = d_742_v10_
                index94_ = default__.safeIndex(605, (p1).length(0))
                (p1)[index94_] = d_735_v3_
                index95_ = default__.safeIndex(605, (p1).length(0))
                (p1)[index95_] = d_735_v3_
                d_743_v11_: D1
                d_743_v11_ = D1_DC4(p3)
                d_744_v12_: _dafny.Map
                d_744_v12_ = _dafny.Map({d_743_v11_: p3})
                d_745_v13_: _dafny.Map
                d_745_v13_ = _dafny.Map({d_730_v0_: d_744_v12_})
                d_745_v13_ = (_dafny.Map({d_730_v0_: _dafny.Map({d_743_v11_: -438})})) | (_dafny.Map({d_730_v0_: d_744_v12_}))
            elif True:
                d_746_v14_: _dafny.Seq
                d_746_v14_ = _dafny.SeqWithoutIsStrInference([d_730_v0_, not(not(d_730_v0_))])
                d_747_v15_: C0
                nw139_ = C0()
                nw139_.ctor__((d_746_v14_)[default__.safeIndex(p2, len(d_746_v14_))], d_731_v1_)
                d_747_v15_ = nw139_
                d_748_v16_: _dafny.Array
                def lambda48_(d_749_v0_):
                    def lambda49_(d_750_i1_):
                        return (d_750_i1_) * (len(_dafny.Map({d_749_v0_: True})))

                    return lambda49_

                init26_ = lambda48_(d_730_v0_)
                nw140_ = _dafny.Array(None, 22)
                for i0_26_ in range(nw140_.length(0)):
                    nw140_[i0_26_] = init26_(i0_26_)
                d_748_v16_ = nw140_
                index96_ = default__.safeIndex(108, (d_748_v16_).length(0))
                (d_748_v16_)[index96_] = p0
                index97_ = default__.safeIndex(108, (d_748_v16_).length(0))
                (d_748_v16_)[index97_] = (self).f12
                d_751_v17_: _dafny.Set
                d_751_v17_ = _dafny.Set({_dafny.Map({d_730_v0_: (d_747_v15_).f10})})
                d_751_v17_ = default__.fm13(globalState)
                d_752_v18_: D4
                d_752_v18_ = D4_DC14()
                rhs107_ = ((0) - (default__.safeModuloInt(p0, p2)) if (d_747_v15_).f10 else default__.safeDivisionInt((self).f12, p2))
                rhs108_ = d_752_v18_
                lhs78_ = globalState
                lhs78_.f5 = rhs107_
                d_752_v18_ = rhs108_
                d_753_v19_: _dafny.Seq
                d_753_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqxmp"))
                d_753_v19_ = ((d_753_v19_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_754_i2_ in range(default__.abs(-716))])) if d_730_v0_ else (default__.fm14(d_753_v19_, p3, globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fraq"))))
            (globalState).f5 = (0) - ((p0) * (len(_dafny.SeqWithoutIsStrInference([d_730_v0_, d_730_v0_, d_730_v0_]))))
            r1 = p2
            d_755_v20_: D4
            d_755_v20_ = D4_DC11(d_731_v1_)
            rhs109_ = d_755_v20_
            rhs110_ = False
            lhs79_ = globalState
            d_755_v20_ = rhs109_
            lhs79_.f0 = rhs110_
        elif True:
            d_756_v21_: _dafny.Map
            d_756_v21_ = _dafny.Map({d_730_v0_: d_730_v0_})
            d_757_v22_: _dafny.Set
            d_757_v22_ = _dafny.Set({d_730_v0_, d_730_v0_})
            d_756_v21_ = (d_756_v21_).set((d_757_v22_).isdisjoint(d_757_v22_), d_730_v0_)
            d_758_v23_: _dafny.MultiSet
            d_758_v23_ = _dafny.MultiSet([p3, (self).f12])
            d_759_v24_: _dafny.Seq
            d_759_v24_ = _dafny.SeqWithoutIsStrInference([p0])
            d_760_v25_: _dafny.Seq
            d_760_v25_ = _dafny.SeqWithoutIsStrInference([d_730_v0_, d_730_v0_, d_730_v0_, d_730_v0_])
            d_761_v26_: _dafny.Map
            d_761_v26_ = _dafny.Map({p3: d_730_v0_})
            d_762_v27_: D0
            d_762_v27_ = D0_DC0(d_761_v26_)
            d_763_v28_: _dafny.MultiSet
            d_763_v28_ = _dafny.MultiSet([False])
            d_764_v29_: _dafny.Seq
            d_764_v29_ = _dafny.SeqWithoutIsStrInference([d_730_v0_, (d_760_v25_)[default__.safeIndex(p2, len(d_760_v25_))], (self).fm4(d_762_v27_, D0_DC0(d_761_v26_), d_763_v28_, globalState), True])
            rhs111_ = (_dafny.MultiSet(d_759_v24_)).issubset(d_758_v23_)
            rhs112_ = (default__.safeModuloInt(p3, (self).f12)) + (len((d_764_v29_) + (d_764_v29_)))
            lhs80_ = globalState
            lhs81_ = globalState
            lhs80_.f0 = rhs111_
            lhs81_.f5 = rhs112_
            d_765_v30_: D0
            d_765_v30_ = D0_DC1(d_730_v0_)
            source11_ = d_765_v30_
            if source11_.is_DC1:
                d_766___mcc_h0_ = source11_.cf1
                d_767_cf1_ = d_766___mcc_h0_
                (globalState).f0 = d_767_cf1_
                d_768_v31_: _dafny.Map
                d_768_v31_ = _dafny.Map({52: len(d_764_v29_)})
                d_769_v32_: _dafny.Array
                nw141_ = _dafny.Array(None, 19)
                nw141_[int(0)] = not((d_765_v30_).cf1)
                nw141_[int(1)] = d_767_cf1_
                nw141_[int(2)] = d_730_v0_
                nw141_[int(3)] = d_730_v0_
                nw141_[int(4)] = d_767_cf1_
                nw141_[int(5)] = d_730_v0_
                nw141_[int(6)] = True
                nw141_[int(7)] = (self).fm4(d_762_v27_, d_762_v27_, d_763_v28_, globalState)
                nw141_[int(8)] = d_767_cf1_
                nw141_[int(9)] = d_730_v0_
                nw141_[int(10)] = d_730_v0_
                nw141_[int(11)] = d_767_cf1_
                nw141_[int(12)] = d_767_cf1_
                nw141_[int(13)] = d_730_v0_
                nw141_[int(14)] = d_767_cf1_
                nw141_[int(15)] = d_730_v0_
                nw141_[int(16)] = d_767_cf1_
                nw141_[int(17)] = (self).fm4(d_762_v27_, d_762_v27_, d_763_v28_, globalState)
                nw141_[int(18)] = d_767_cf1_
                d_769_v32_ = nw141_
                d_770_v33_: C0
                nw142_ = C0()
                nw142_.ctor__((p2) > (((d_768_v31_)[p2] if (p2) in (d_768_v31_) else -979)), d_769_v32_)
                d_770_v33_ = nw142_
                d_771_v34_: _dafny.Seq
                d_771_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvwujy"))
                d_771_v34_ = d_771_v34_
                d_772_v35_: _dafny.Array
                nw143_ = _dafny.Array(_dafny.CodePoint('D'), 6)
                d_772_v35_ = nw143_
                d_772_v35_ = d_772_v35_
            elif source11_.is_DC0:
                d_773___mcc_h1_ = source11_.cf0
                d_774_cf0_ = d_773___mcc_h1_
                d_775_v36_: _dafny.Array
                nw144_ = _dafny.Array(None, 29)
                nw144_[int(0)] = d_730_v0_
                nw144_[int(1)] = d_730_v0_
                nw144_[int(2)] = d_730_v0_
                nw144_[int(3)] = d_730_v0_
                nw144_[int(4)] = d_730_v0_
                nw144_[int(5)] = d_730_v0_
                nw144_[int(6)] = d_730_v0_
                nw144_[int(7)] = d_730_v0_
                nw144_[int(8)] = d_730_v0_
                nw144_[int(9)] = d_730_v0_
                nw144_[int(10)] = d_730_v0_
                nw144_[int(11)] = d_730_v0_
                nw144_[int(12)] = False
                nw144_[int(13)] = d_730_v0_
                nw144_[int(14)] = d_730_v0_
                nw144_[int(15)] = d_730_v0_
                nw144_[int(16)] = True
                nw144_[int(17)] = d_730_v0_
                nw144_[int(18)] = d_730_v0_
                nw144_[int(19)] = d_730_v0_
                nw144_[int(20)] = not(d_730_v0_)
                nw144_[int(21)] = d_730_v0_
                nw144_[int(22)] = d_730_v0_
                nw144_[int(23)] = d_730_v0_
                nw144_[int(24)] = d_730_v0_
                nw144_[int(25)] = d_730_v0_
                nw144_[int(26)] = (d_760_v25_)[default__.safeIndex(87, len(d_760_v25_))]
                nw144_[int(27)] = d_730_v0_
                nw144_[int(28)] = d_730_v0_
                d_775_v36_ = nw144_
                d_776_v37_: _dafny.Map
                d_776_v37_ = _dafny.Map({d_775_v36_: d_730_v0_})
                d_776_v37_ = (d_776_v37_) | (d_776_v37_)
                d_777_v38_: str
                d_777_v38_ = _dafny.CodePoint('x')
                d_778_v40_: _dafny.Map
                d_778_v40_ = _dafny.Map({d_730_v0_: (self).f12})
                d_779_v41_: _dafny.Seq
                d_779_v41_ = _dafny.SeqWithoutIsStrInference([d_778_v40_])
                def iife59_():
                    coll39_ = _dafny.Set()
                    compr_39_: int
                    for compr_39_ in _dafny.IntegerRange(364, 751):
                        d_780_v39_: int = compr_39_
                        if ((364) <= (d_780_v39_)) and ((d_780_v39_) < (751)):
                            coll39_ = coll39_.union(_dafny.Set([(d_780_v39_) + (165)]))
                    return _dafny.Set(coll39_)
                default__.m0(d_730_v0_, (d_777_v38_ if d_730_v0_ else d_777_v38_), len(iife59_()
                ), (d_778_v40_) | ((d_779_v41_)[default__.safeIndex((self).f12, len(d_779_v41_))]), globalState)
                d_781_v42_: _dafny.Array
                nw145_ = _dafny.Array(_dafny.Map({}), 9)
                d_781_v42_ = nw145_
                d_782_v43_: D1
                d_782_v43_ = D1_DC3(d_781_v42_)
                pat_let_tv23_ = d_781_v42_
                def iife60_(_pat_let10_0):
                    def iife61_(d_783_dt__update__tmp_h2_):
                        def iife62_(_pat_let11_0):
                            def iife63_(d_784_dt__update_hcf3_h0_):
                                return D1_DC3(d_784_dt__update_hcf3_h0_)
                            return iife63_(_pat_let11_0)
                        return iife62_(pat_let_tv23_)
                    return iife61_(_pat_let10_0)
                d_782_v43_ = iife60_(d_782_v43_)
                (globalState).f0 = (d_730_v0_ if d_730_v0_ else d_730_v0_)
            elif True:
                d_785___mcc_h2_ = source11_.cf2
                d_786_cf2_ = d_785___mcc_h2_
                d_787_v44_: _dafny.Array
                nw146_ = _dafny.Array(_dafny.MultiSet({}), 23)
                d_787_v44_ = nw146_
                d_788_v45_: _dafny.MultiSet
                d_788_v45_ = _dafny.MultiSet([d_764_v29_, d_764_v29_, d_760_v25_])
                index98_ = default__.safeIndex(392, (d_787_v44_).length(0))
                (d_787_v44_)[index98_] = (d_788_v45_).intersection((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_764_v29_, _dafny.SeqWithoutIsStrInference([d_730_v0_])]))).set(d_760_v25_, default__.abs(p0)))
                d_789_v46_: _dafny.Seq
                d_789_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofcufvnr"))
                d_790_v47_: _dafny.Map
                d_790_v47_ = _dafny.Map({len(d_789_v46_): p0})
                d_791_v48_: _dafny.Seq
                d_791_v48_ = _dafny.SeqWithoutIsStrInference([d_790_v47_])
                d_792_v49_: _dafny.Array
                nw147_ = _dafny.Array(int(0), 23)
                d_792_v49_ = nw147_
                index99_ = default__.safeIndex(666, (p1).length(0))
                (p1)[index99_] = d_792_v49_
                d_793_v50_: _dafny.Map
                d_793_v50_ = _dafny.Map({d_730_v0_: p2})
                index100_ = default__.safeIndex(392, (d_787_v44_).length(0))
                index101_ = default__.safeIndex(666, (p1).length(0))
                rhs113_ = len(d_793_v50_)
                rhs114_ = d_788_v45_
                rhs115_ = _dafny.SeqWithoutIsStrInference([d_790_v47_])
                rhs116_ = d_792_v49_
                rhs117_ = (self).f12
                lhs82_ = d_787_v44_
                lhs83_ = default__.safeIndex(392, (d_787_v44_).length(0))
                lhs84_ = p1
                lhs85_ = default__.safeIndex(666, (p1).length(0))
                lhs86_ = globalState
                r1 = rhs113_
                lhs82_[lhs83_] = rhs114_
                d_791_v48_ = rhs115_
                lhs84_[lhs85_] = rhs116_
                lhs86_.f5 = rhs117_
                d_794_v51_: _dafny.Array
                nw148_ = _dafny.Array(None, 23)
                nw148_[int(0)] = d_759_v24_
                nw148_[int(1)] = (d_759_v24_) + (d_759_v24_)
                nw148_[int(2)] = (d_759_v24_ if d_730_v0_ else default__.fm15(d_730_v0_, d_730_v0_, globalState))
                nw148_[int(3)] = d_759_v24_
                nw148_[int(4)] = (d_759_v24_) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p0 for d_795_i4_ in range(default__.abs(432))])) for d_796_i3_ in range(default__.abs(-462))]))
                nw148_[int(5)] = _dafny.SeqWithoutIsStrInference([(self).f12])
                nw148_[int(6)] = _dafny.SeqWithoutIsStrInference([p3 for d_797_i5_ in range(default__.abs(56))])
                nw148_[int(7)] = d_759_v24_
                nw148_[int(8)] = (_dafny.SeqWithoutIsStrInference([p0 for d_798_i6_ in range(default__.abs(679))])) + (d_759_v24_)
                nw148_[int(9)] = _dafny.SeqWithoutIsStrInference([(self).f12, p3, default__.fm0(499, len(d_790_v47_), p3, d_730_v0_, globalState)])
                nw148_[int(10)] = d_759_v24_
                nw148_[int(11)] = d_759_v24_
                nw148_[int(12)] = d_759_v24_
                nw148_[int(13)] = d_759_v24_
                nw148_[int(14)] = _dafny.SeqWithoutIsStrInference([p3 for d_799_i7_ in range(default__.abs(985))])
                nw148_[int(15)] = d_759_v24_
                nw148_[int(16)] = d_759_v24_
                nw148_[int(17)] = _dafny.SeqWithoutIsStrInference([(self).f12, p2])
                nw148_[int(18)] = d_759_v24_
                nw148_[int(19)] = (_dafny.SeqWithoutIsStrInference([((d_793_v50_)[False] if (False) in (d_793_v50_) else p2) for d_800_i8_ in range(default__.abs(590))])) + (d_759_v24_)
                nw148_[int(20)] = d_759_v24_
                nw148_[int(21)] = d_759_v24_
                nw148_[int(22)] = d_759_v24_
                d_794_v51_ = nw148_
                index102_ = default__.safeIndex(542, (d_794_v51_).length(0))
                (d_794_v51_)[index102_] = d_759_v24_
                index103_ = default__.safeIndex(542, (d_794_v51_).length(0))
                (d_794_v51_)[index103_] = _dafny.SeqWithoutIsStrInference([p2 for d_801_i9_ in range(default__.abs(939))])
                d_802_v52_: _dafny.Array
                def lambda50_(d_803_v0_):
                    def lambda51_(d_804_i10_):
                        return d_803_v0_

                    return lambda51_

                init27_ = lambda50_(d_730_v0_)
                nw149_ = _dafny.Array(None, 25)
                for i0_27_ in range(nw149_.length(0)):
                    nw149_[i0_27_] = init27_(i0_27_)
                d_802_v52_ = nw149_
                d_802_v52_ = d_802_v52_
                d_805_v53_: _dafny.Map
                d_805_v53_ = _dafny.Map({(p1)[default__.safeIndex(666, (p1).length(0))]: d_730_v0_})
                index104_ = default__.safeIndex(582, (d_792_v49_).length(0))
                (d_792_v49_)[index104_] = default__.fm0(p2, p3, p0, ((d_805_v53_)[d_792_v49_] if (d_792_v49_) in (d_805_v53_) else d_730_v0_), globalState)
                index105_ = default__.safeIndex(582, (d_792_v49_).length(0))
                (d_792_v49_)[index105_] = (self).f12
            d_806_v54_: D0
            d_806_v54_ = D0_DC0(d_761_v26_)
            d_807_v55_: D0
            d_807_v55_ = D0_DC2(d_806_v54_)
            d_808_v56_: _dafny.MultiSet
            d_808_v56_ = _dafny.MultiSet([d_807_v55_])
            d_809_v57_: _dafny.Seq
            d_809_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfcvgsxqa"))
            d_810_v58_: _dafny.Seq
            d_810_v58_ = _dafny.SeqWithoutIsStrInference([d_760_v25_, _dafny.SeqWithoutIsStrInference([d_730_v0_, d_730_v0_]), d_764_v29_, d_764_v29_, d_764_v29_])
            d_811_v59_: D2
            d_811_v59_ = D2_DC6(d_730_v0_)
            d_812_v60_: _dafny.Array
            nw150_ = _dafny.Array(None, 18)
            nw150_[int(0)] = (d_808_v56_).isdisjoint(d_808_v56_)
            nw150_[int(1)] = d_730_v0_
            nw150_[int(2)] = d_730_v0_
            nw150_[int(3)] = d_730_v0_
            nw150_[int(4)] = (d_759_v24_) == (_dafny.SeqWithoutIsStrInference([len(d_764_v29_), p3, len(d_809_v57_)]))
            nw150_[int(5)] = d_730_v0_
            nw150_[int(6)] = not(d_730_v0_)
            nw150_[int(7)] = (d_810_v58_) < (_dafny.SeqWithoutIsStrInference([d_760_v25_, d_764_v29_]))
            nw150_[int(8)] = False
            nw150_[int(9)] = d_730_v0_
            nw150_[int(10)] = d_730_v0_
            nw150_[int(11)] = d_730_v0_
            nw150_[int(12)] = d_730_v0_
            nw150_[int(13)] = d_730_v0_
            nw150_[int(14)] = d_730_v0_
            nw150_[int(15)] = (d_811_v59_).cf6
            nw150_[int(16)] = d_730_v0_
            nw150_[int(17)] = d_730_v0_
            d_812_v60_ = nw150_
            index106_ = default__.safeIndex(498, (d_812_v60_).length(0))
            (d_812_v60_)[index106_] = (d_730_v0_) or (d_730_v0_)
            index107_ = default__.safeIndex(498, (d_812_v60_).length(0))
            rhs118_ = d_730_v0_
            rhs119_ = p3
            rhs120_ = (0) - ((0) - ((p2) - ((p0) * ((self).f12))))
            rhs121_ = (p0 if (d_811_v59_).cf6 else p2)
            lhs87_ = d_812_v60_
            lhs88_ = default__.safeIndex(498, (d_812_v60_).length(0))
            lhs89_ = globalState
            lhs90_ = globalState
            lhs91_ = globalState
            lhs87_[lhs88_] = rhs118_
            lhs89_.f5 = rhs119_
            lhs90_.f5 = rhs120_
            lhs91_.f5 = rhs121_
            d_813_v61_: _dafny.Map
            d_813_v61_ = _dafny.Map({(d_812_v60_)[default__.safeIndex(498, (d_812_v60_).length(0))]: p0})
            d_814_v62_: _dafny.Set
            d_814_v62_ = _dafny.Set({len(d_756_v21_)})
            d_815_v63_: _dafny.Map
            d_815_v63_ = _dafny.Map({(0) - (len(d_814_v62_)): -557})
            d_816_v64_: _dafny.Array
            nw151_ = _dafny.Array(None, 13)
            nw151_[int(0)] = ((self).f12) * (p2)
            nw151_[int(1)] = (self).f12
            nw151_[int(2)] = (d_758_v23_).cardinality
            nw151_[int(3)] = p0
            nw151_[int(4)] = p0
            nw151_[int(5)] = p2
            nw151_[int(6)] = ((d_813_v61_)[d_730_v0_] if (d_730_v0_) in (d_813_v61_) else ((d_815_v63_)[len(d_761_v26_)] if (len(d_761_v26_)) in (d_815_v63_) else p0))
            nw151_[int(7)] = p0
            nw151_[int(8)] = 687
            nw151_[int(9)] = 448
            nw151_[int(10)] = p3
            nw151_[int(11)] = 952
            nw151_[int(12)] = (0) - (p2)
            d_816_v64_ = nw151_
            rhs122_ = default__.safeDivisionInt(p0, p0)
            rhs123_ = (807 if (d_812_v60_)[default__.safeIndex(498, (d_812_v60_).length(0))] else (p3) + ((_dafny.MultiSet([(d_812_v60_)[default__.safeIndex(498, (d_812_v60_).length(0))]])).cardinality))
            rhs124_ = d_816_v64_
            rhs125_ = not(d_730_v0_)
            lhs92_ = globalState
            lhs93_ = globalState
            lhs92_.f5 = rhs122_
            lhs93_.f5 = rhs123_
            d_816_v64_ = rhs124_
            d_730_v0_ = rhs125_
        d_817_v65_: str
        d_817_v65_ = _dafny.CodePoint('u')
        d_817_v65_ = default__.fm16(globalState)
        d_818_v66_: _dafny.MultiSet
        d_818_v66_ = _dafny.MultiSet([True, False])
        r1 = ((self).f12) - ((d_818_v66_).cardinality)
        d_819_i11_: int
        d_819_i11_ = 0
        with _dafny.label("6"):
            while True:
                with _dafny.c_label("6"):
                    if (d_819_i11_) >= (100):
                        raise _dafny.Break("6")
                    d_819_i11_ = (d_819_i11_) + (1)
                    if False:
                        d_820_v67_: _dafny.Map
                        d_820_v67_ = _dafny.Map({(self).f12: d_730_v0_})
                        d_821_v68_: _dafny.Seq
                        d_821_v68_ = _dafny.SeqWithoutIsStrInference([False])
                        d_822_v69_: _dafny.Seq
                        d_822_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdak"))
                        d_823_v70_: _dafny.Map
                        d_823_v70_ = _dafny.Map({_dafny.Map({p2: d_822_v69_}): False})
                        d_824_v72_: _dafny.Array
                        nw152_ = _dafny.Array(None, 28)
                        nw152_[int(0)] = True
                        nw152_[int(1)] = d_730_v0_
                        nw152_[int(2)] = False
                        nw152_[int(3)] = d_730_v0_
                        nw152_[int(4)] = d_730_v0_
                        nw152_[int(5)] = d_730_v0_
                        nw152_[int(6)] = ((d_820_v67_)[len((d_821_v68_).set(default__.safeIndex(p2, len(d_821_v68_)), d_730_v0_))] if (len((d_821_v68_).set(default__.safeIndex(p2, len(d_821_v68_)), d_730_v0_))) in (d_820_v67_) else d_730_v0_)
                        nw152_[int(7)] = d_730_v0_
                        nw152_[int(8)] = default__.fm1(d_730_v0_, globalState)
                        nw152_[int(9)] = d_730_v0_
                        nw152_[int(10)] = d_730_v0_
                        nw152_[int(11)] = d_730_v0_
                        nw152_[int(12)] = d_730_v0_
                        nw152_[int(13)] = d_730_v0_
                        nw152_[int(14)] = True
                        nw152_[int(15)] = d_730_v0_
                        nw152_[int(16)] = d_730_v0_
                        nw152_[int(17)] = d_730_v0_
                        nw152_[int(18)] = d_730_v0_
                        nw152_[int(19)] = d_730_v0_
                        nw152_[int(20)] = d_730_v0_
                        nw152_[int(21)] = not(d_730_v0_)
                        nw152_[int(22)] = d_730_v0_
                        nw152_[int(23)] = d_730_v0_
                        nw152_[int(24)] = d_730_v0_
                        nw152_[int(25)] = default__.fm1(d_730_v0_, globalState)
                        def iife64_():
                            coll40_ = _dafny.Map()
                            compr_40_: int
                            for compr_40_ in (d_820_v67_).keys.Elements:
                                d_825_v71_: int = compr_40_
                                if (d_825_v71_) in (d_820_v67_):
                                    coll40_[(d_825_v71_) * (p0)] = d_822_v69_
                            return _dafny.Map(coll40_)
                        def iife65_():
                            coll41_ = _dafny.Map()
                            compr_41_: int
                            for compr_41_ in (d_820_v67_).keys.Elements:
                                d_826_v71_: int = compr_41_
                                if (d_826_v71_) in (d_820_v67_):
                                    coll41_[(d_826_v71_) * (p0)] = d_822_v69_
                            return _dafny.Map(coll41_)
                        nw152_[int(26)] = ((d_823_v70_)[iife64_()
                        ] if (iife65_()
                        ) in (d_823_v70_) else d_730_v0_)
                        nw152_[int(27)] = d_730_v0_
                        d_824_v72_ = nw152_
                        d_827_v73_: C0
                        nw153_ = C0()
                        nw153_.ctor__(True, d_824_v72_)
                        d_827_v73_ = nw153_
                        d_828_v75_: D0
                        def iife66_():
                            coll42_ = _dafny.Map()
                            compr_42_: int
                            for compr_42_ in _dafny.IntegerRange(725, -252):
                                d_829_v74_: int = compr_42_
                                if ((725) <= (d_829_v74_)) and ((d_829_v74_) < (-252)):
                                    coll42_[default__.safeDivisionInt(d_829_v74_, p3)] = (d_827_v73_).f10
                            return _dafny.Map(coll42_)
                        d_828_v75_ = D0_DC0(iife66_()
)
                        pat_let_tv24_ = d_820_v67_
                        def iife67_(_pat_let12_0):
                            def iife68_(d_830_dt__update__tmp_h3_):
                                def iife69_(_pat_let13_0):
                                    def iife70_(d_831_dt__update_hcf0_h0_):
                                        return D0_DC0(d_831_dt__update_hcf0_h0_)
                                    return iife70_(_pat_let13_0)
                                return iife69_(pat_let_tv24_)
                            return iife68_(_pat_let12_0)
                        d_824_v72_ = ((d_827_v73_).f11 if (self).fm4(iife67_(d_828_v75_), D0_DC0(_dafny.Map({p2: d_730_v0_})), _dafny.MultiSet(d_821_v68_), globalState) else d_824_v72_)
                        d_832_v76_: _dafny.Set
                        d_832_v76_ = _dafny.Set({(d_827_v73_).f10})
                        (globalState).f5 = (len(d_832_v76_)) * (p3)
                        r0 = ((d_832_v76_) - (d_832_v76_)).issubset(d_832_v76_)
                        d_833_v77_: C0
                        nw154_ = C0()
                        nw154_.ctor__(d_730_v0_, (d_827_v73_).f11)
                        d_833_v77_ = nw154_
                    elif True:
                        d_834_v78_: _dafny.Array
                        def lambda52_(d_835_v0_):
                            def lambda53_(d_836_i12_):
                                return d_835_v0_

                            return lambda53_

                        init28_ = lambda52_(d_730_v0_)
                        nw155_ = _dafny.Array(None, 4)
                        for i0_28_ in range(nw155_.length(0)):
                            nw155_[i0_28_] = init28_(i0_28_)
                        d_834_v78_ = nw155_
                        d_837_v79_: C0
                        nw156_ = C0()
                        nw156_.ctor__(d_730_v0_, d_834_v78_)
                        d_837_v79_ = nw156_
                        d_838_v80_: _dafny.Array
                        nw157_ = _dafny.Array(None, 4)
                        nw157_[int(0)] = p2
                        nw157_[int(1)] = p0
                        nw157_[int(2)] = p0
                        nw157_[int(3)] = p0
                        d_838_v80_ = nw157_
                        d_839_v81_: _dafny.Map
                        d_839_v81_ = _dafny.Map({d_730_v0_: d_838_v80_})
                        d_840_v82_: _dafny.Seq
                        d_840_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dypeuvr"))
                        d_841_v83_: _dafny.Set
                        d_841_v83_ = _dafny.Set({(d_840_v82_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etmifa"))), len(d_840_v82_))]})
                        d_842_v84_: _dafny.Map
                        d_842_v84_ = _dafny.Map({d_730_v0_: len((d_840_v82_).set(default__.safeIndex(p3, len(d_840_v82_)), d_817_v65_))})
                        d_843_v85_: _dafny.Array
                        nw158_ = _dafny.Array(None, 23)
                        nw158_[int(0)] = -539
                        nw158_[int(1)] = (len(d_839_v81_)) * ((0) - (p0))
                        nw158_[int(2)] = p2
                        nw158_[int(3)] = p0
                        nw158_[int(4)] = p0
                        nw158_[int(5)] = ((self).f12 if d_730_v0_ else p0)
                        nw158_[int(6)] = default__.fm0(default__.fm0(p0, p0, len(default__.fm17(True, p3, d_730_v0_, p0, globalState)), d_730_v0_, globalState), (self).f12, len(d_841_v83_), d_730_v0_, globalState)
                        nw158_[int(7)] = (len(d_842_v84_) if d_730_v0_ else (0) - ((self).f12))
                        nw158_[int(8)] = (self).f12
                        nw158_[int(9)] = (self).f12
                        nw158_[int(10)] = p3
                        nw158_[int(11)] = p0
                        nw158_[int(12)] = p3
                        nw158_[int(13)] = (p2) * (p2)
                        nw158_[int(14)] = len(((_dafny.SeqWithoutIsStrInference([568])).set(default__.safeIndex(845, len(_dafny.SeqWithoutIsStrInference([568]))), len(d_840_v82_)) if (d_837_v79_).f10 else _dafny.SeqWithoutIsStrInference([p2 for d_844_i13_ in range(default__.abs(225))])))
                        nw158_[int(15)] = p0
                        nw158_[int(16)] = (d_818_v66_).cardinality
                        nw158_[int(17)] = p3
                        nw158_[int(18)] = default__.safeModuloInt(len(d_840_v82_), (self).f12)
                        nw158_[int(19)] = default__.safeDivisionInt(p2, p2)
                        nw158_[int(20)] = (_dafny.MultiSet([(self).f12])).cardinality
                        nw158_[int(21)] = default__.safeModuloInt((self).f12, (self).f12)
                        nw158_[int(22)] = -487
                        d_843_v85_ = nw158_
                        d_843_v85_ = d_843_v85_
                        (globalState).f5 = (self).f12
                        (globalState).f0 = not(not (not((self).fm12(d_840_v82_, (self).f12, p0, globalState))) or ((d_837_v79_).f10))
                        (globalState).f0 = d_730_v0_
                    if d_730_v0_:
                        d_845_v86_: _dafny.Array
                        nw159_ = _dafny.Array(False, 11)
                        d_845_v86_ = nw159_
                        d_845_v86_ = d_845_v86_
                        d_846_v87_: _dafny.Seq
                        d_846_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqxxg"))
                        d_847_v88_: _dafny.Array
                        nw160_ = _dafny.Array(None, 15)
                        d_847_v88_ = nw160_
                        d_848_v89_: int
                        d_849_v90_: _dafny.Array
                        out4_: int
                        out5_: _dafny.Array
                        out4_, out5_ = (self).m3(default__.fm14(d_846_v87_, p2, globalState), d_730_v0_, d_847_v88_, p0, globalState)
                        d_848_v89_ = out4_
                        d_849_v90_ = out5_
                        d_846_v87_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))) + (d_846_v87_)) + (d_846_v87_)
                        d_850_v91_: _dafny.Seq
                        d_850_v91_ = _dafny.SeqWithoutIsStrInference([len(d_846_v87_)])
                        d_851_v92_: _dafny.Map
                        d_851_v92_ = _dafny.Map({d_850_v91_: p2})
                        d_851_v92_ = ((d_851_v92_) | (d_851_v92_)) | (d_851_v92_)
                        index108_ = default__.safeIndex(366, (d_845_v86_).length(0))
                        (d_845_v86_)[index108_] = (d_730_v0_ if d_730_v0_ else True)
                        d_852_v93_: T0
                        nw161_ = C2()
                        nw161_.ctor__(False)
                        d_852_v93_ = nw161_
                        d_853_v94_: _dafny.Map
                        d_853_v94_ = _dafny.Map({(self).f12: d_852_v93_})
                        d_854_v95_: _dafny.MultiSet
                        d_854_v95_ = _dafny.MultiSet([d_848_v89_, -22, 181, len(d_853_v94_), -526])
                        d_855_v96_: _dafny.Map
                        d_855_v96_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([p3 for d_856_i14_ in range(default__.abs(192))])): d_730_v0_})
                        d_857_v97_: D0
                        d_857_v97_ = D0_DC0(d_855_v96_)
                        d_858_v98_: D0
                        d_858_v98_ = D0_DC2(d_857_v97_)
                        d_859_v99_: D0
                        d_859_v99_ = D0_DC0(d_855_v96_)
                        d_860_v100_: D12
                        d_860_v100_ = D12_DC34(d_858_v98_, d_730_v0_, (d_852_v93_).fm4(d_859_v99_, d_859_v99_, d_818_v66_, globalState))
                        d_861_v101_: D12
                        d_861_v101_ = D12_DC34(D0_DC2(d_857_v97_), d_730_v0_, not((d_860_v100_).cf50))
                        d_862_v102_: _dafny.Seq
                        d_862_v102_ = _dafny.SeqWithoutIsStrInference([d_730_v0_, (d_854_v95_).isdisjoint(_dafny.MultiSet(d_850_v91_)), (d_860_v100_).cf49, True])
                        index109_ = default__.safeIndex(366, (d_845_v86_).length(0))
                        (d_845_v86_)[index109_] = (d_862_v102_)[default__.safeIndex(d_848_v89_, len(d_862_v102_))]
                    elif True:
                        d_863_v103_: _dafny.Map
                        d_863_v103_ = _dafny.Map({d_730_v0_: d_730_v0_})
                        d_864_v104_: _dafny.Seq
                        d_864_v104_ = _dafny.SeqWithoutIsStrInference([d_863_v103_])
                        d_865_v105_: _dafny.Map
                        d_865_v105_ = _dafny.Map({(d_864_v104_)[default__.safeIndex(p0, len(d_864_v104_))]: 762})
                        d_866_v106_: D2
                        d_866_v106_ = D2_DC7(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "heyc"))).set(default__.safeIndex((0) - ((self).f12), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "heyc")))), d_817_v65_)), (d_865_v105_).set(d_863_v103_, (self).f12), (self).f12)
                        (globalState).f5 = (d_866_v106_).cf9
                        d_867_v107_: _dafny.Seq
                        d_867_v107_ = _dafny.SeqWithoutIsStrInference([not(d_730_v0_)])
                        d_868_v108_: D9
                        d_868_v108_ = D9_DC24(True, True, (d_867_v107_)[default__.safeIndex((0) - ((self).f12), len(d_867_v107_))])
                        d_869_v109_: _dafny.Map
                        d_869_v109_ = _dafny.Map({p3: d_730_v0_})
                        d_870_v110_: _dafny.Map
                        d_870_v110_ = _dafny.Map({(d_867_v107_)[default__.safeIndex(534, len(d_867_v107_))]: p0})
                        rhs126_ = (D9_DC24(False, d_730_v0_, d_730_v0_) if (((d_869_v109_)[632] if (632) in (d_869_v109_) else ((d_863_v103_)[False] if (False) in (d_863_v103_) else d_730_v0_))) == (d_730_v0_) else d_868_v108_)
                        rhs127_ = default__.safeDivisionInt(len((D13_DC38(d_870_v110_)).cf55), default__.safeDivisionInt(len(d_867_v107_), p3))
                        d_868_v108_ = rhs126_
                        r1 = rhs127_
                        d_871_v111_: _dafny.Map
                        d_871_v111_ = _dafny.Map({p3: (self).f12})
                        d_872_v112_: _dafny.Map
                        d_872_v112_ = _dafny.Map({d_871_v111_: d_730_v0_})
                        (globalState).f0 = not(not(not(((d_872_v112_)[d_871_v111_] if (d_871_v111_) in (d_872_v112_) else d_730_v0_))))
                        d_873_v113_: _dafny.Seq
                        d_873_v113_ = _dafny.SeqWithoutIsStrInference([(self).f12])
                        d_874_v114_: _dafny.Set
                        d_874_v114_ = _dafny.Set({default__.fm0(len(d_873_v113_), p0, p0, d_730_v0_, globalState), p0})
                        d_875_v116_: _dafny.Array
                        nw162_ = _dafny.Array(None, 9)
                        nw162_[int(0)] = d_730_v0_
                        nw162_[int(1)] = True
                        nw162_[int(2)] = (d_873_v113_) == (d_873_v113_)
                        nw162_[int(3)] = d_730_v0_
                        nw162_[int(4)] = not(d_730_v0_)
                        nw162_[int(5)] = d_730_v0_
                        nw162_[int(6)] = d_730_v0_
                        nw162_[int(7)] = d_730_v0_
                        def iife71_():
                            coll43_ = _dafny.Set()
                            compr_43_: int
                            for compr_43_ in _dafny.IntegerRange(750, -18):
                                d_876_v115_: int = compr_43_
                                if ((750) <= (d_876_v115_)) and ((d_876_v115_) < (-18)):
                                    coll43_ = coll43_.union(_dafny.Set([default__.safeModuloInt(d_876_v115_, p0)]))
                            return _dafny.Set(coll43_)
                        nw162_[int(8)] = (d_874_v114_).isdisjoint(iife71_()
                        )
                        d_875_v116_ = nw162_
                        index110_ = default__.safeIndex(933, (d_875_v116_).length(0))
                        (d_875_v116_)[index110_] = d_730_v0_
                        d_877_v117_: _dafny.Array
                        nw163_ = _dafny.Array(_dafny.Array(None, 0), 16)
                        d_877_v117_ = nw163_
                        d_878_v118_: _dafny.Array
                        def lambda54_(d_879_p3_):
                            def lambda55_(d_880_i15_):
                                return default__.safeModuloInt(d_880_i15_, d_879_p3_)

                            return lambda55_

                        init29_ = lambda54_(p3)
                        nw164_ = _dafny.Array(None, 11)
                        for i0_29_ in range(nw164_.length(0)):
                            nw164_[i0_29_] = init29_(i0_29_)
                        d_878_v118_ = nw164_
                        d_881_v119_: _dafny.Array
                        d_881_v119_ = d_878_v118_
                        index111_ = default__.safeIndex(888, (d_877_v117_).length(0))
                        (d_877_v117_)[index111_] = d_881_v119_
                        index112_ = default__.safeIndex(933, (d_875_v116_).length(0))
                        index113_ = default__.safeIndex(888, (d_877_v117_).length(0))
                        rhs128_ = not ((d_730_v0_) == (not(d_730_v0_))) or (d_730_v0_)
                        rhs129_ = False
                        rhs130_ = d_881_v119_
                        lhs94_ = d_875_v116_
                        lhs95_ = default__.safeIndex(933, (d_875_v116_).length(0))
                        lhs96_ = d_877_v117_
                        lhs97_ = default__.safeIndex(888, (d_877_v117_).length(0))
                        r2 = rhs128_
                        lhs94_[lhs95_] = rhs129_
                        lhs96_[lhs97_] = rhs130_
                        d_882_v120_: D3
                        d_882_v120_ = D3_DC9(409)
                        d_883_v121_: D3
                        d_883_v121_ = D3_DC10(d_882_v120_)
                        d_884_v122_: _dafny.Seq
                        d_884_v122_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxm"))
                        pat_let_tv25_ = d_884_v122_
                        d_885_v123_: _dafny.Map
                        def iife72_(_pat_let14_0):
                            def iife73_(d_886_dt__update__tmp_h4_):
                                def iife74_(_pat_let15_0):
                                    def iife75_(d_887_dt__update_hcf12_h0_):
                                        return D3_DC10(d_887_dt__update_hcf12_h0_)
                                    return iife75_(_pat_let15_0)
                                return iife74_(D3_DC8(pat_let_tv25_))
                            return iife73_(_pat_let14_0)
                        d_885_v123_ = _dafny.Map({p3: iife72_(d_883_v121_)})
                        (globalState).f0 = (len(d_885_v123_)) > (p0)
                    d_888_v124_: D11
                    d_888_v124_ = D11_DC30()
                    d_889_v125_: D11
                    d_889_v125_ = D11_DC32(d_888_v124_)
                    d_890_v126_: D11
                    d_890_v126_ = D11_DC32((d_889_v125_).cf46)
                    d_891_v128_: _dafny.Array
                    def lambda56_(d_892_v66_, d_893_p2_):
                        def lambda57_(d_894_i16_):
                            def iife76_():
                                coll44_ = _dafny.Set()
                                compr_44_: int
                                for compr_44_ in _dafny.IntegerRange(349, 538):
                                    d_895_v127_: int = compr_44_
                                    if ((349) <= (d_895_v127_)) and ((d_895_v127_) < (538)):
                                        coll44_ = coll44_.union(_dafny.Set([default__.safeDivisionInt(d_895_v127_, d_893_p2_)]))
                                return _dafny.Set(coll44_)
                            return D4_DC12((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f12, len(iife76_()
)]))).cardinality, len(_dafny.SeqWithoutIsStrInference([d_892_v66_])), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxnwwoae")), (0) - (len(_dafny.Set({True}))), d_893_p2_)

                        return lambda57_

                    init30_ = lambda56_(d_818_v66_, p2)
                    nw165_ = _dafny.Array(None, 12)
                    for i0_30_ in range(nw165_.length(0)):
                        nw165_[i0_30_] = init30_(i0_30_)
                    d_891_v128_ = nw165_
                    d_896_v129_: _dafny.Map
                    d_896_v129_ = _dafny.Map({(d_890_v126_ if d_730_v0_ else d_890_v126_): d_891_v128_})
                    d_896_v129_ = (d_896_v129_).set(d_890_v126_, d_891_v128_)
                    d_897_v130_: D11
                    d_897_v130_ = D11_DC30()
                    d_898_v131_: _dafny.Seq
                    d_898_v131_ = _dafny.SeqWithoutIsStrInference([D11_DC30(), d_897_v130_, d_897_v130_, d_897_v130_])
                    d_899_v132_: _dafny.MultiSet
                    d_899_v132_ = _dafny.MultiSet([(0) - (p0)])
                    d_900_v133_: _dafny.Seq
                    d_900_v133_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, (d_899_v132_).cardinality])])
                    d_901_v134_: _dafny.Seq
                    d_901_v134_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxyqh"))
                    d_902_v135_: _dafny.Array
                    nw166_ = _dafny.Array(None, 19)
                    nw166_[int(0)] = p0
                    nw166_[int(1)] = p2
                    nw166_[int(2)] = ((self).f12) * (p3)
                    nw166_[int(3)] = (0) - (len((_dafny.SeqWithoutIsStrInference([D11_DC30() for d_903_i17_ in range(default__.abs(73))])) + (d_898_v131_)))
                    nw166_[int(4)] = p2
                    nw166_[int(5)] = p3
                    nw166_[int(6)] = p2
                    nw166_[int(7)] = p3
                    nw166_[int(8)] = ((self).f12) + (p3)
                    nw166_[int(9)] = default__.safeDivisionInt(len(d_900_v133_), p3)
                    nw166_[int(10)] = (self).f12
                    nw166_[int(11)] = default__.safeDivisionInt((0) - (p3), len(_dafny.SeqWithoutIsStrInference([d_730_v0_, not(d_730_v0_), not(d_730_v0_), d_730_v0_, d_730_v0_])))
                    nw166_[int(12)] = -586
                    nw166_[int(13)] = (self).f12
                    nw166_[int(14)] = p2
                    nw166_[int(15)] = default__.safeDivisionInt(p0, (self).f12)
                    nw166_[int(16)] = len((d_901_v134_) + (d_901_v134_))
                    nw166_[int(17)] = p0
                    nw166_[int(18)] = p2
                    d_902_v135_ = nw166_
                    index114_ = default__.safeIndex(756, (d_902_v135_).length(0))
                    (d_902_v135_)[index114_] = (p2) + (len(d_901_v134_))
                    index115_ = default__.safeIndex(756, (d_902_v135_).length(0))
                    (d_902_v135_)[index115_] = p2
                    pass
            pass
        d_904_v136_: _dafny.Array
        nw167_ = _dafny.Array(_dafny.Seq({}), 22)
        d_904_v136_ = nw167_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_904_v136_).length(0)):
            d_905_i18_: int = guard_loop_2_
            if (True) and (((0) <= (d_905_i18_)) and ((d_905_i18_) < ((d_904_v136_).length(0)))):
                (d_904_v136_)[(d_905_i18_)] = (_dafny.SeqWithoutIsStrInference([d_730_v0_, d_730_v0_])) + ((_dafny.SeqWithoutIsStrInference([d_730_v0_, d_730_v0_, d_730_v0_])) + (_dafny.SeqWithoutIsStrInference([False])))
        d_906_v137_: _dafny.Seq
        d_906_v137_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
        d_906_v137_ = (d_906_v137_) + (d_906_v137_)
        r0 = d_730_v0_
        r1 = p3
        d_907_v138_: D11
        d_907_v138_ = D11_DC31((self).f12, (0) - (p2))
        pat_let_tv26_ = d_730_v0_
        pat_let_tv27_ = d_730_v0_
        pat_let_tv28_ = d_730_v0_
        pat_let_tv29_ = d_730_v0_
        def lambda58_(source12_):
            if source12_.is_DC30:
                return pat_let_tv26_
            elif source12_.is_DC31:
                d_908___mcc_h3_ = source12_.cf44
                d_909___mcc_h4_ = source12_.cf45
                d_910_cf45_ = d_909___mcc_h4_
                d_911_cf44_ = d_908___mcc_h3_
                return pat_let_tv27_
            elif source12_.is_DC29:
                d_912___mcc_h5_ = source12_.cf43
                d_913_cf43_ = d_912___mcc_h5_
                return pat_let_tv28_
            elif True:
                d_914___mcc_h6_ = source12_.cf46
                d_915_cf46_ = d_914___mcc_h6_
                return pat_let_tv29_

        r2 = lambda58_(d_907_v138_)
        r3 = not(not (d_730_v0_) or (d_730_v0_))
        return r0, r1, r2, r3

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r0 = p3
        (globalState).f2 = not(not(p1))
        d_916_v0_: D0
        d_916_v0_ = D0_DC1(p1)
        d_917_v1_: D0
        d_917_v1_ = D0_DC2(d_916_v0_)
        d_918_v2_: D12
        d_918_v2_ = D12_DC34(d_917_v1_, p1, True)
        d_919_i0_: int
        d_919_i0_ = 0
        with _dafny.label("7"):
            while (d_918_v2_).cf50:
                with _dafny.c_label("7"):
                    if (d_919_i0_) >= (100):
                        raise _dafny.Break("7")
                    d_919_i0_ = (d_919_i0_) + (1)
                    d_920_v3_: C2
                    nw168_ = C2()
                    nw168_.ctor__(p1)
                    d_920_v3_ = nw168_
                    d_921_v4_: _dafny.Map
                    d_921_v4_ = _dafny.Map({p1: d_920_v3_})
                    d_922_v7_: D0
                    def iife77_():
                        def iife79_():
                            coll47_ = _dafny.Map()
                            compr_47_: int
                            for compr_47_ in _dafny.IntegerRange(695, 246):
                                d_923_v6_: int = compr_47_
                                if ((695) <= (d_923_v6_)) and ((d_923_v6_) < (246)):
                                    coll47_[default__.safeModuloInt(d_923_v6_, (0) - (p3))] = (d_920_v3_).f13
                            return _dafny.Map(coll47_)
                        coll45_ = _dafny.Map()
                        def iife78_():
                            coll46_ = _dafny.Map()
                            compr_46_: int
                            for compr_46_ in _dafny.IntegerRange(695, 246):
                                d_923_v6_: int = compr_46_
                                if ((695) <= (d_923_v6_)) and ((d_923_v6_) < (246)):
                                    coll46_[default__.safeModuloInt(d_923_v6_, (0) - (p3))] = (d_920_v3_).f13
                            return _dafny.Map(coll46_)
                        compr_45_: int
                        for compr_45_ in (iife78_()
                        ).keys.Elements:
                            d_924_v5_: int = compr_45_
                            if (d_924_v5_) in (iife79_()
                            ):
                                coll45_[default__.safeDivisionInt(d_924_v5_, p3)] = (d_920_v3_).f13
                        return _dafny.Map(coll45_)
                    d_922_v7_ = D0_DC0(iife77_()
)
                    d_925_v8_: _dafny.MultiSet
                    d_925_v8_ = _dafny.MultiSet([(d_920_v3_).f13])
                    d_926_v9_: _dafny.Set
                    d_926_v9_ = _dafny.Set({p1})
                    d_927_v10_: _dafny.Array
                    nw169_ = _dafny.Array(None, 17)
                    nw169_[int(0)] = (p1) not in (d_921_v4_)
                    nw169_[int(1)] = (d_920_v3_).f13
                    nw169_[int(2)] = True
                    nw169_[int(3)] = p1
                    nw169_[int(4)] = ((d_920_v3_).f13) and (p1)
                    nw169_[int(5)] = p1
                    nw169_[int(6)] = (p3) >= ((self).f12)
                    nw169_[int(7)] = not (p1) or (not((d_920_v3_).f13))
                    nw169_[int(8)] = (d_920_v3_).f13
                    nw169_[int(9)] = ((d_920_v3_).f13) == (p1)
                    nw169_[int(10)] = default__.fm1((d_920_v3_).f13, globalState)
                    nw169_[int(11)] = (d_920_v3_).f13
                    nw169_[int(12)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_928_i1_ in range(default__.abs(-224))])) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_929_i2_ in range(default__.abs(953))]))
                    nw169_[int(13)] = (p1 if (d_920_v3_).f13 else (d_920_v3_).f13)
                    nw169_[int(14)] = (self).fm4(d_922_v7_, d_922_v7_, d_925_v8_, globalState)
                    nw169_[int(15)] = ((self).f12) == (len(d_926_v9_))
                    nw169_[int(16)] = (d_920_v3_).f13
                    d_927_v10_ = nw169_
                    d_930_v11_: _dafny.Seq
                    d_930_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "akjkn"))
                    index116_ = default__.safeIndex(273, (d_927_v10_).length(0))
                    (d_927_v10_)[index116_] = (d_920_v3_).f13
                    d_931_v12_: _dafny.MultiSet
                    d_931_v12_ = _dafny.MultiSet([d_917_v1_])
                    index117_ = default__.safeIndex(273, (d_927_v10_).length(0))
                    rhs131_ = d_927_v10_
                    rhs132_ = _dafny.SeqWithoutIsStrInference([(D10_DC27(p1, (self).f12, _dafny.CodePoint('n'), _dafny.MultiSet([False, (d_920_v3_).f13, False, p1]))).cf40 for d_932_i3_ in range(default__.abs(755))])
                    rhs133_ = (d_931_v12_).isdisjoint((d_931_v12_).intersection(default__.fm31(globalState)))
                    lhs98_ = d_927_v10_
                    lhs99_ = default__.safeIndex(273, (d_927_v10_).length(0))
                    d_927_v10_ = rhs131_
                    d_930_v11_ = rhs132_
                    lhs98_[lhs99_] = rhs133_
                    (globalState).f5 = (self).f12
                    d_933_v13_: str
                    d_933_v13_ = _dafny.CodePoint('s')
                    d_933_v13_ = _dafny.CodePoint('f')
                    d_934_v14_: _dafny.Set
                    d_934_v14_ = _dafny.Set({(self).f12, (self).f12})
                    d_935_v15_: D4
                    d_935_v15_ = D4_DC13(p3, d_927_v10_)
                    d_936_v16_: C0
                    nw170_ = C0()
                    nw170_.ctor__((d_934_v14_) == (d_934_v14_), (d_935_v15_).cf20)
                    d_936_v16_ = nw170_
                    pass
            pass
        d_937_v17_: _dafny.MultiSet
        d_937_v17_ = _dafny.MultiSet([p3, (self).f12, p3])
        d_938_v18_: _dafny.Seq
        d_938_v18_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1, not(p1), default__.fm1(not(p1), globalState)])
        hi4_ = (self).f12
        for d_939_i4_ in range(((d_937_v17_).cardinality) + (len(d_938_v18_)), hi4_):
            d_938_v18_ = d_938_v18_
            d_940_v19_: C2
            nw171_ = C2()
            nw171_.ctor__(p1)
            d_940_v19_ = nw171_
            d_941_v20_: _dafny.Array
            nw172_ = _dafny.Array(False, 21)
            d_941_v20_ = nw172_
            d_942_v21_: C0
            nw173_ = C0()
            nw173_.ctor__((d_938_v18_)[default__.safeIndex(p3, len(d_938_v18_))], d_941_v20_)
            d_942_v21_ = nw173_
            d_943_v22_: _dafny.Map
            d_943_v22_ = _dafny.Map({not(p1): (0) - (p3)})
            d_943_v22_ = (d_943_v22_).set((d_940_v19_).f13, (self).f12)
        hi5_ = ((self).f12) * (p3)
        for d_944_i5_ in range(default__.safeDivisionInt((d_937_v17_).cardinality, p3), hi5_):
            (globalState).f5 = len((p0) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjl"))))
            d_945_v23_: D7
            d_945_v23_ = D7_DC18()
            d_946_v24_: D7
            d_946_v24_ = D7_DC19(d_945_v23_)
            d_947_v25_: D7
            d_947_v25_ = D7_DC19(d_946_v24_)
            d_948_v26_: D7
            d_948_v26_ = D7_DC19((d_947_v25_).cf24)
            d_949_v27_: D7
            d_949_v27_ = D7_DC19(d_945_v23_)
            d_950_v28_: _dafny.Map
            d_950_v28_ = _dafny.Map({d_947_v25_: True})
            d_950_v28_ = (d_950_v28_).set(d_949_v27_, p1)
            d_951_v29_: _dafny.Array
            nw174_ = _dafny.Array(False, 19)
            d_951_v29_ = nw174_
            d_952_v30_: C0
            nw175_ = C0()
            nw175_.ctor__(not (p1) or (p1), d_951_v29_)
            d_952_v30_ = nw175_
            (globalState).f2 = not((len(p0)) >= (len(d_938_v18_)))
        d_953_v31_: _dafny.Array
        nw176_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
        d_953_v31_ = nw176_
        d_953_v31_ = d_953_v31_
        d_954_v32_: _dafny.Map
        d_954_v32_ = _dafny.Map({False: p1})
        d_955_v33_: _dafny.Map
        d_955_v33_ = _dafny.Map({d_954_v32_: (self).f12})
        r0 = (D2_DC7((self).f12, d_955_v33_, p3)).cf9
        d_956_v34_: _dafny.Array
        nw177_ = _dafny.Array(int(0), 2)
        d_956_v34_ = nw177_
        r1 = d_956_v34_
        return r0, r1

    @property
    def f12(self):
        return self._f12

class C4(T1, T0):
    def  __init__(self):
        self._f9: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self, f9):
        (self)._f9 = f9

    def fm4(self, p0, p1, p2, globalState):
        return (default__.safeModuloInt((self).f9, (self).f9)) > (((_dafny.MultiSet([True])).intersection(_dafny.MultiSet([False, True]))).cardinality)

    def fm5(self, p0, p1, p2, p3, globalState):
        return (0) - (len(_dafny.Set({(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urcwm")))) == ((self).f9), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwdsvbe"))) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_957_i0_ in range(default__.abs(-918))])), False})))

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        r3: bool = False
        r1 = default__.safeDivisionInt((0) - ((p0) * (p2)), p3)
        d_958_v0_: _dafny.Array
        nw178_ = _dafny.Array(False, 6)
        d_958_v0_ = nw178_
        d_959_v1_: C0
        nw179_ = C0()
        nw179_.ctor__(True, d_958_v0_)
        d_959_v1_ = nw179_
        d_960_v2_: C0
        nw180_ = C0()
        nw180_.ctor__(((d_959_v1_).f10) == ((d_959_v1_).f10), d_958_v0_)
        d_960_v2_ = nw180_
        index118_ = default__.safeIndex(204, ((d_959_v1_).f11).length(0))
        ((d_959_v1_).f11)[index118_] = (d_959_v1_).f10
        d_961_v3_: str
        d_961_v3_ = _dafny.CodePoint('m')
        d_962_v4_: _dafny.Array
        def lambda59_(d_963_p0_):
            def lambda60_(d_964_i0_):
                return default__.safeDivisionInt(d_964_i0_, d_963_p0_)

            return lambda60_

        init31_ = lambda59_(p0)
        nw181_ = _dafny.Array(None, 8)
        for i0_31_ in range(nw181_.length(0)):
            nw181_[i0_31_] = init31_(i0_31_)
        d_962_v4_ = nw181_
        index119_ = default__.safeIndex(136, (d_962_v4_).length(0))
        (d_962_v4_)[index119_] = p0
        d_965_v5_: _dafny.MultiSet
        d_965_v5_ = _dafny.MultiSet([(d_959_v1_).f10])
        index120_ = default__.safeIndex(204, ((d_959_v1_).f11).length(0))
        index121_ = default__.safeIndex(136, (d_962_v4_).length(0))
        rhs134_ = ((d_965_v5_).intersection(_dafny.MultiSet([(d_959_v1_).f10, (d_959_v1_).f10]))) != ((d_965_v5_) - (d_965_v5_))
        rhs135_ = d_961_v3_
        rhs136_ = default__.safeDivisionInt(p2, (d_965_v5_).cardinality)
        lhs100_ = (d_959_v1_).f11
        lhs101_ = default__.safeIndex(204, ((d_959_v1_).f11).length(0))
        lhs102_ = d_962_v4_
        lhs103_ = default__.safeIndex(136, (d_962_v4_).length(0))
        lhs100_[lhs101_] = rhs134_
        d_961_v3_ = rhs135_
        lhs102_[lhs103_] = rhs136_
        d_966_v6_: _dafny.Seq
        d_966_v6_ = _dafny.SeqWithoutIsStrInference([(d_960_v2_).f10, (d_959_v1_).f10, (d_960_v2_).f10])
        d_967_v7_: _dafny.Map
        d_967_v7_ = _dafny.Map({(d_966_v6_)[default__.safeIndex(p2, len(d_966_v6_))]: (d_962_v4_)[default__.safeIndex(136, (d_962_v4_).length(0))]})
        d_967_v7_ = (d_967_v7_).set(((d_959_v1_).f11)[default__.safeIndex(204, ((d_959_v1_).f11).length(0))], (d_962_v4_)[default__.safeIndex(136, (d_962_v4_).length(0))])
        r2 = ((d_959_v1_).f11)[default__.safeIndex(204, ((d_959_v1_).f11).length(0))]
        r0 = (d_959_v1_).f10
        r1 = p2
        r2 = default__.fm1((d_960_v2_).f10, globalState)
        d_968_v8_: D0
        d_968_v8_ = D0_DC1((d_960_v2_).f10)
        pat_let_tv30_ = d_960_v2_
        pat_let_tv31_ = d_959_v1_
        pat_let_tv32_ = d_959_v1_
        def lambda61_(source13_):
            if source13_.is_DC1:
                d_969___mcc_h0_ = source13_.cf1
                d_970_cf1_ = d_969___mcc_h0_
                return (_dafny.Set({False})).isdisjoint(_dafny.Set({True}))
            elif source13_.is_DC0:
                d_971___mcc_h1_ = source13_.cf0
                d_972_cf0_ = d_971___mcc_h1_
                return not (False) or ((pat_let_tv30_).f10)
            elif True:
                d_973___mcc_h2_ = source13_.cf2
                d_974_cf2_ = d_973___mcc_h2_
                return ((pat_let_tv31_).f10) and ((pat_let_tv32_).f10)

        r3 = lambda61_(d_968_v8_)
        return r0, r1, r2, r3

    def m2(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: int = int(0)
        r2: bool = False
        d_975_v0_: _dafny.Seq
        d_975_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vg"))
        d_976_v1_: bool
        d_976_v1_ = False
        d_977_v2_: D0
        d_977_v2_ = D0_DC0(_dafny.Map({len(d_975_v0_): d_976_v1_}))
        pat_let_tv33_ = d_976_v1_
        pat_let_tv34_ = d_976_v1_
        pat_let_tv35_ = d_976_v1_
        def lambda62_(source15_):
            if source15_.is_DC1:
                d_978___mcc_h3_ = source15_.cf1
                d_979_cf1_ = d_978___mcc_h3_
                if True:
                    return D0_DC1(True)
                elif True:
                    return D0_DC1(pat_let_tv33_)
            elif source15_.is_DC0:
                d_980___mcc_h4_ = source15_.cf0
                d_981_cf0_ = d_980___mcc_h4_
                return D0_DC1(pat_let_tv34_)
            elif True:
                d_982___mcc_h5_ = source15_.cf2
                d_983_cf2_ = d_982___mcc_h5_
                return D0_DC1(not(pat_let_tv35_))

        source14_ = lambda62_(d_977_v2_)
        if source14_.is_DC1:
            d_984___mcc_h0_ = source14_.cf1
            d_985_cf1_ = d_984___mcc_h0_
            r1 = p0
            (globalState).f5 = (0) - ((0) - (p0))
            d_986_v3_: _dafny.Seq
            d_986_v3_ = _dafny.SeqWithoutIsStrInference([False])
            d_987_v4_: _dafny.Seq
            d_987_v4_ = _dafny.SeqWithoutIsStrInference([(len(d_986_v3_)) > (p1)])
            rhs137_ = d_985_cf1_
            rhs138_ = d_987_v4_
            rhs139_ = len(d_987_v4_)
            lhs104_ = globalState
            lhs105_ = globalState
            lhs104_.f0 = rhs137_
            d_986_v3_ = rhs138_
            lhs105_.f5 = rhs139_
            d_988_v5_: str
            d_988_v5_ = _dafny.CodePoint('c')
            d_988_v5_ = d_988_v5_
        elif source14_.is_DC0:
            d_989___mcc_h1_ = source14_.cf0
            d_990_cf0_ = d_989___mcc_h1_
            d_991_v6_: _dafny.Map
            d_991_v6_ = _dafny.Map({d_976_v1_: p1})
            d_992_v7_: _dafny.Array
            nw182_ = _dafny.Array(None, 26)
            nw182_[int(0)] = True
            nw182_[int(1)] = d_976_v1_
            nw182_[int(2)] = d_976_v1_
            nw182_[int(3)] = d_976_v1_
            nw182_[int(4)] = not(d_976_v1_)
            nw182_[int(5)] = d_976_v1_
            nw182_[int(6)] = d_976_v1_
            nw182_[int(7)] = not(d_976_v1_)
            nw182_[int(8)] = d_976_v1_
            nw182_[int(9)] = d_976_v1_
            nw182_[int(10)] = False
            nw182_[int(11)] = d_976_v1_
            nw182_[int(12)] = d_976_v1_
            nw182_[int(13)] = d_976_v1_
            nw182_[int(14)] = d_976_v1_
            nw182_[int(15)] = d_976_v1_
            nw182_[int(16)] = d_976_v1_
            nw182_[int(17)] = d_976_v1_
            nw182_[int(18)] = d_976_v1_
            nw182_[int(19)] = d_976_v1_
            nw182_[int(20)] = d_976_v1_
            nw182_[int(21)] = d_976_v1_
            nw182_[int(22)] = d_976_v1_
            nw182_[int(23)] = not(d_976_v1_)
            nw182_[int(24)] = d_976_v1_
            nw182_[int(25)] = d_976_v1_
            d_992_v7_ = nw182_
            d_993_v8_: T1
            nw183_ = C0()
            nw183_.ctor__(d_976_v1_, d_992_v7_)
            d_993_v8_ = nw183_
            d_994_v9_: _dafny.Map
            d_994_v9_ = _dafny.Map({d_991_v6_: d_993_v8_})
            d_995_v10_: D0
            d_995_v10_ = D0_DC0(d_990_cf0_)
            d_996_v11_: D0
            d_996_v11_ = D0_DC2(d_995_v10_)
            rhs140_ = ((d_994_v9_) | (d_994_v9_)).set((d_991_v6_) | (d_991_v6_), (d_993_v8_ if d_976_v1_ else d_993_v8_))
            rhs141_ = d_996_v11_
            d_994_v9_ = rhs140_
            d_996_v11_ = rhs141_
            if d_976_v1_:
                d_990_cf0_ = (d_990_cf0_).set(p1, d_976_v1_)
                d_997_v12_: _dafny.Set
                d_997_v12_ = _dafny.Set({False})
                d_998_v13_: _dafny.Seq
                d_998_v13_ = _dafny.SeqWithoutIsStrInference([d_997_v12_])
                d_999_v14_: str
                d_999_v14_ = _dafny.CodePoint('f')
                d_1000_v15_: _dafny.MultiSet
                d_1000_v15_ = _dafny.MultiSet([(D2_DC5(d_993_v8_)).cf5, d_993_v8_, d_993_v8_])
                (globalState).f5 = default__.fm0((self).f9, len(d_975_v0_), (p1) + ((self).fm5(d_998_v13_, d_999_v14_, _dafny.SeqWithoutIsStrInference([(d_1000_v15_).cardinality]), _dafny.SeqWithoutIsStrInference([835]), globalState)), d_976_v1_, globalState)
                d_1001_v16_: _dafny.MultiSet
                d_1001_v16_ = _dafny.MultiSet([d_976_v1_])
                d_1002_v17_: _dafny.Map
                d_1002_v17_ = _dafny.Map({d_1001_v16_: d_976_v1_})
                d_999_v14_ = default__.fm9((0) - (default__.safeDivisionInt(len(p2), len(d_1002_v17_))), globalState)
                (globalState).f5 = p1
                d_990_cf0_ = (d_990_cf0_).set((0) - (p1), True)
            elif True:
                d_1003_v18_: _dafny.Array
                def lambda63_(d_1004_i0_):
                    return default__.safeDivisionInt(d_1004_i0_, (self).f9)

                init32_ = lambda63_
                nw184_ = _dafny.Array(None, 21)
                for i0_32_ in range(nw184_.length(0)):
                    nw184_[i0_32_] = init32_(i0_32_)
                d_1003_v18_ = nw184_
                d_1005_v19_: _dafny.Set
                d_1005_v19_ = _dafny.Set({len(d_975_v0_)})
                rhs142_ = d_1003_v18_
                rhs143_ = False
                rhs144_ = d_1005_v19_
                lhs106_ = globalState
                d_1003_v18_ = rhs142_
                lhs106_.f0 = rhs143_
                d_1005_v19_ = rhs144_
                d_1006_v20_: _dafny.MultiSet
                d_1006_v20_ = _dafny.MultiSet([d_976_v1_])
                (globalState).f5 = default__.safeDivisionInt(p1, ((d_1006_v20_)[(d_993_v8_).fm4(d_977_v2_, d_977_v2_, d_1006_v20_, globalState)] if ((d_993_v8_).fm4(d_977_v2_, d_977_v2_, d_1006_v20_, globalState)) in (d_1006_v20_) else -158))
                rhs145_ = (p1) + ((self).f9)
                rhs146_ = (self).f9
                rhs147_ = p1
                lhs107_ = globalState
                r1 = rhs145_
                lhs107_.f5 = rhs146_
                r1 = rhs147_
                d_1007_v21_: str
                d_1007_v21_ = _dafny.CodePoint('i')
                d_1008_v22_: _dafny.Map
                d_1008_v22_ = _dafny.Map({d_976_v1_: d_1007_v21_})
                d_1009_v23_: D3
                d_1009_v23_ = D3_DC8(d_975_v0_)
                d_1010_v24_: _dafny.Array
                nw185_ = _dafny.Array(None, 21)
                nw185_[int(0)] = d_975_v0_
                nw185_[int(1)] = d_975_v0_
                nw185_[int(2)] = d_975_v0_
                nw185_[int(3)] = (d_975_v0_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1011_i1_ in range(default__.abs(667))]))
                nw185_[int(4)] = d_975_v0_
                nw185_[int(5)] = (d_975_v0_) + (d_975_v0_)
                nw185_[int(6)] = d_975_v0_
                nw185_[int(7)] = (d_975_v0_) + (d_975_v0_)
                nw185_[int(8)] = d_975_v0_
                nw185_[int(9)] = (d_975_v0_) + (d_975_v0_)
                nw185_[int(10)] = d_975_v0_
                nw185_[int(11)] = d_975_v0_
                nw185_[int(12)] = d_975_v0_
                nw185_[int(13)] = (d_975_v0_) + (d_975_v0_)
                nw185_[int(14)] = d_975_v0_
                nw185_[int(15)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1012_i2_ in range(default__.abs(-366))])
                nw185_[int(16)] = d_975_v0_
                nw185_[int(17)] = (default__.fm8(d_976_v1_, _dafny.Map({True: d_1007_v21_}), d_976_v1_, globalState)) + (d_975_v0_)
                nw185_[int(18)] = default__.fm8(d_976_v1_, d_1008_v22_, d_976_v1_, globalState)
                nw185_[int(19)] = d_975_v0_
                nw185_[int(20)] = (_dafny.SeqWithoutIsStrInference([d_1007_v21_ for d_1013_i3_ in range(default__.abs(273))])) + ((d_1009_v23_).cf10)
                d_1010_v24_ = nw185_
                index122_ = default__.safeIndex(109, (d_1010_v24_).length(0))
                (d_1010_v24_)[index122_] = (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggh"))) + (d_975_v0_)).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggh"))) + (d_975_v0_))), d_1007_v21_)).set(default__.safeIndex(p1, len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggh"))) + (d_975_v0_)).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggh"))) + (d_975_v0_))), d_1007_v21_))), d_1007_v21_)
                index123_ = default__.safeIndex(109, (d_1010_v24_).length(0))
                (d_1010_v24_)[index123_] = (_dafny.SeqWithoutIsStrInference([d_1007_v21_ for d_1014_i4_ in range(default__.abs(827))])) + (d_975_v0_)
                d_1015_v25_: _dafny.Array
                nw186_ = _dafny.Array(None, 28)
                nw186_[int(0)] = _dafny.CodePoint('g')
                nw186_[int(1)] = d_1007_v21_
                nw186_[int(2)] = d_1007_v21_
                nw186_[int(3)] = _dafny.CodePoint('l')
                nw186_[int(4)] = d_1007_v21_
                nw186_[int(5)] = d_1007_v21_
                nw186_[int(6)] = _dafny.CodePoint('v')
                nw186_[int(7)] = d_1007_v21_
                nw186_[int(8)] = d_1007_v21_
                nw186_[int(9)] = d_1007_v21_
                nw186_[int(10)] = d_1007_v21_
                nw186_[int(11)] = d_1007_v21_
                nw186_[int(12)] = d_1007_v21_
                nw186_[int(13)] = d_1007_v21_
                nw186_[int(14)] = _dafny.CodePoint('l')
                nw186_[int(15)] = d_1007_v21_
                nw186_[int(16)] = d_1007_v21_
                nw186_[int(17)] = d_1007_v21_
                nw186_[int(18)] = d_1007_v21_
                nw186_[int(19)] = ((d_1010_v24_)[default__.safeIndex(109, (d_1010_v24_).length(0))])[default__.safeIndex((0) - ((d_1006_v20_).cardinality), len((d_1010_v24_)[default__.safeIndex(109, (d_1010_v24_).length(0))]))]
                nw186_[int(20)] = d_1007_v21_
                nw186_[int(21)] = d_1007_v21_
                nw186_[int(22)] = d_1007_v21_
                nw186_[int(23)] = d_1007_v21_
                nw186_[int(24)] = d_1007_v21_
                nw186_[int(25)] = d_1007_v21_
                nw186_[int(26)] = d_1007_v21_
                nw186_[int(27)] = d_1007_v21_
                d_1015_v25_ = nw186_
                index124_ = default__.safeIndex(483, (d_1015_v25_).length(0))
                (d_1015_v25_)[index124_] = d_1007_v21_
                index125_ = default__.safeIndex(838, (d_1015_v25_).length(0))
                (d_1015_v25_)[index125_] = d_1007_v21_
                index126_ = default__.safeIndex(483, (d_1015_v25_).length(0))
                index127_ = default__.safeIndex(838, (d_1015_v25_).length(0))
                rhs148_ = d_1003_v18_
                rhs149_ = (p0) - (p0)
                rhs150_ = d_1007_v21_
                rhs151_ = d_1007_v21_
                lhs108_ = globalState
                lhs109_ = d_1015_v25_
                lhs110_ = default__.safeIndex(483, (d_1015_v25_).length(0))
                lhs111_ = d_1015_v25_
                lhs112_ = default__.safeIndex(838, (d_1015_v25_).length(0))
                d_1003_v18_ = rhs148_
                lhs108_.f5 = rhs149_
                lhs109_[lhs110_] = rhs150_
                lhs111_[lhs112_] = rhs151_
            d_1016_v26_: _dafny.Array
            nw187_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_1016_v26_ = nw187_
            d_1017_v27_: _dafny.Array
            def lambda64_(d_1018_cf0_, d_1019_v0_):
                def lambda65_(d_1020_i5_):
                    return (_dafny.SeqWithoutIsStrInference([D3_DC10(D3_DC8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewwnswuai")))), D3_DC10(D3_DC9(len(d_1018_cf0_)))])) + (_dafny.SeqWithoutIsStrInference([D3_DC10((D3_DC10(D3_DC8(d_1019_v0_))).cf12) for d_1021_i6_ in range(default__.abs(724))]))

                return lambda65_

            init33_ = lambda64_(d_990_cf0_, d_975_v0_)
            nw188_ = _dafny.Array(None, 6)
            for i0_33_ in range(nw188_.length(0)):
                nw188_[i0_33_] = init33_(i0_33_)
            d_1017_v27_ = nw188_
            d_1022_v28_: D3
            d_1022_v28_ = D3_DC8(d_975_v0_)
            d_1023_v29_: _dafny.Seq
            d_1023_v29_ = _dafny.SeqWithoutIsStrInference([D3_DC10(d_1022_v28_)])
            index128_ = default__.safeIndex(866, (d_1017_v27_).length(0))
            (d_1017_v27_)[index128_] = (d_1023_v29_) + (d_1023_v29_)
            d_1024_v30_: _dafny.Seq
            d_1024_v30_ = _dafny.SeqWithoutIsStrInference([(True) or (d_976_v1_)])
            index129_ = default__.safeIndex(866, (d_1017_v27_).length(0))
            rhs152_ = (d_1024_v30_)[default__.safeIndex(p1, len(d_1024_v30_))]
            rhs153_ = d_1016_v26_
            rhs154_ = d_1023_v29_
            rhs155_ = (self).f9
            rhs156_ = ((self).f9) < (len((d_1024_v30_) + (_dafny.SeqWithoutIsStrInference([d_976_v1_]))))
            lhs113_ = globalState
            lhs114_ = d_1017_v27_
            lhs115_ = default__.safeIndex(866, (d_1017_v27_).length(0))
            lhs116_ = globalState
            lhs113_.f2 = rhs152_
            d_1016_v26_ = rhs153_
            lhs114_[lhs115_] = rhs154_
            lhs116_.f5 = rhs155_
            r2 = rhs156_
            (globalState).f5 = p0
        elif True:
            d_1025___mcc_h2_ = source14_.cf2
            d_1026_cf2_ = d_1025___mcc_h2_
            rhs157_ = (default__.safeDivisionInt(p0, (self).f9)) - ((self).f9)
            rhs158_ = d_976_v1_
            rhs159_ = (self).f9
            lhs117_ = globalState
            lhs118_ = globalState
            lhs117_.f5 = rhs157_
            d_976_v1_ = rhs158_
            lhs118_.f5 = rhs159_
            d_1027_v31_: _dafny.Array
            nw189_ = _dafny.Array(int(0), 2)
            d_1027_v31_ = nw189_
            index130_ = default__.safeIndex(555, (d_1027_v31_).length(0))
            (d_1027_v31_)[index130_] = 903
            index131_ = default__.safeIndex(555, (d_1027_v31_).length(0))
            (d_1027_v31_)[index131_] = default__.safeDivisionInt(p1, (p0) + (717))
            d_1028_v32_: _dafny.Seq
            d_1028_v32_ = _dafny.SeqWithoutIsStrInference([d_1027_v31_])
            d_1029_v33_: _dafny.Array
            nw190_ = _dafny.Array(None, 2)
            nw190_[int(0)] = (d_1028_v32_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbtwqs"))), len(d_1028_v32_))]
            nw190_[int(1)] = d_1027_v31_
            d_1029_v33_ = nw190_
            d_1030_v34_: _dafny.Array
            nw191_ = _dafny.Array(_dafny.Array(None, 0), 21)
            d_1030_v34_ = nw191_
            d_1031_v35_: _dafny.Array
            nw192_ = _dafny.Array(_dafny.Set({}), 20)
            d_1031_v35_ = nw192_
            index132_ = default__.safeIndex(9, (d_1030_v34_).length(0))
            (d_1030_v34_)[index132_] = d_1031_v35_
            index133_ = default__.safeIndex(9, (d_1030_v34_).length(0))
            rhs160_ = d_1029_v33_
            rhs161_ = d_1031_v35_
            rhs162_ = p1
            rhs163_ = d_975_v0_
            rhs164_ = d_976_v1_
            lhs119_ = d_1030_v34_
            lhs120_ = default__.safeIndex(9, (d_1030_v34_).length(0))
            lhs121_ = globalState
            lhs122_ = globalState
            d_1029_v33_ = rhs160_
            lhs119_[lhs120_] = rhs161_
            lhs121_.f5 = rhs162_
            d_975_v0_ = rhs163_
            lhs122_.f2 = rhs164_
            d_1032_v36_: D3
            d_1032_v36_ = D3_DC9(p0)
            pat_let_tv36_ = d_1032_v36_
            d_1033_v37_: _dafny.MultiSet
            def iife80_(_pat_let16_0):
                def iife81_(d_1034_dt__update__tmp_h0_):
                    def iife82_(_pat_let17_0):
                        def iife83_(d_1035_dt__update_hcf12_h0_):
                            return D3_DC10(d_1035_dt__update_hcf12_h0_)
                        return iife83_(_pat_let17_0)
                    return iife82_(pat_let_tv36_)
                return iife81_(_pat_let16_0)
            d_1033_v37_ = _dafny.MultiSet([D3_DC10(d_1032_v36_), iife80_(D3_DC10(d_1032_v36_))])
            d_1036_v38_: D3
            d_1036_v38_ = D3_DC10(d_1032_v36_)
            d_1037_v39_: D0
            d_1037_v39_ = D0_DC1(d_976_v1_)
            d_1038_v40_: _dafny.Array
            nw193_ = _dafny.Array(None, 15)
            nw193_[int(0)] = d_976_v1_
            nw193_[int(1)] = d_976_v1_
            nw193_[int(2)] = d_976_v1_
            nw193_[int(3)] = d_976_v1_
            nw193_[int(4)] = d_976_v1_
            nw193_[int(5)] = d_976_v1_
            nw193_[int(6)] = True
            nw193_[int(7)] = True
            nw193_[int(8)] = True
            nw193_[int(9)] = d_976_v1_
            nw193_[int(10)] = d_976_v1_
            nw193_[int(11)] = d_976_v1_
            nw193_[int(12)] = d_976_v1_
            nw193_[int(13)] = True
            nw193_[int(14)] = (d_1037_v39_).cf1
            d_1038_v40_ = nw193_
            pat_let_tv37_ = d_1032_v36_
            d_1039_v41_: T1
            nw194_ = C0()
            def iife84_(_pat_let18_0):
                def iife85_(d_1040_dt__update__tmp_h1_):
                    def iife86_(_pat_let19_0):
                        def iife87_(d_1041_dt__update_hcf12_h1_):
                            return D3_DC10(d_1041_dt__update_hcf12_h1_)
                        return iife87_(_pat_let19_0)
                    return iife86_(pat_let_tv37_)
                return iife85_(_pat_let18_0)
            nw194_.ctor__(not(((d_1033_v37_).set(iife84_(d_1036_v38_), default__.abs((d_1027_v31_)[default__.safeIndex(555, (d_1027_v31_).length(0))]))).issubset(d_1033_v37_)), d_1038_v40_)
            d_1039_v41_ = nw194_
            d_1039_v41_ = d_1039_v41_
        if False:
            d_1042_v42_: _dafny.Map
            d_1042_v42_ = _dafny.Map({(self).f9: d_976_v1_})
            d_1043_v43_: _dafny.Seq
            d_1043_v43_ = _dafny.SeqWithoutIsStrInference([d_1042_v42_])
            d_1044_v44_: _dafny.MultiSet
            d_1044_v44_ = _dafny.MultiSet([d_976_v1_])
            d_1045_v45_: _dafny.Set
            d_1045_v45_ = _dafny.Set({d_976_v1_, d_976_v1_})
            d_1046_v46_: str
            d_1046_v46_ = _dafny.CodePoint('u')
            d_1047_v47_: _dafny.Map
            d_1047_v47_ = _dafny.Map({d_976_v1_: d_976_v1_})
            d_1048_v48_: _dafny.Map
            d_1048_v48_ = _dafny.Map({d_1047_v47_: len(d_1045_v45_)})
            d_1049_v49_: D2
            d_1049_v49_ = D2_DC7(p0, d_1048_v48_, len((d_975_v0_).set(default__.safeIndex(p1, len(d_975_v0_)), d_1046_v46_)))
            d_1050_v50_: _dafny.Map
            d_1050_v50_ = _dafny.Map({d_1044_v44_: default__.fm10(d_1045_v45_, d_1046_v46_, d_1049_v49_, p0, globalState)})
            rhs165_ = d_976_v1_
            rhs166_ = (d_1044_v44_) in (d_1050_v50_)
            rhs167_ = ((_dafny.SeqWithoutIsStrInference([(d_1042_v42_).set((self).f9, d_976_v1_), _dafny.Map({p1: d_976_v1_})])) + (_dafny.SeqWithoutIsStrInference([d_1042_v42_ for d_1051_i7_ in range(default__.abs(-788))]))) + (_dafny.SeqWithoutIsStrInference([d_1042_v42_ for d_1052_i8_ in range(default__.abs(971))]))
            d_976_v1_ = rhs165_
            d_976_v1_ = rhs166_
            d_1043_v43_ = rhs167_
            (globalState).f0 = d_976_v1_
            d_1053_v51_: D2
            d_1053_v51_ = D2_DC6(d_976_v1_)
            source16_ = d_1053_v51_
            if source16_.is_DC6:
                d_1054___mcc_h6_ = source16_.cf6
                d_1055_cf6_ = d_1054___mcc_h6_
                d_1056_v52_: _dafny.Array
                def lambda66_(d_1057_cf6_):
                    def lambda67_(d_1058_i9_):
                        return d_1057_cf6_

                    return lambda67_

                init34_ = lambda66_(d_1055_cf6_)
                nw195_ = _dafny.Array(None, 15)
                for i0_34_ in range(nw195_.length(0)):
                    nw195_[i0_34_] = init34_(i0_34_)
                d_1056_v52_ = nw195_
                d_1059_v53_: C0
                nw196_ = C0()
                nw196_.ctor__(d_1055_cf6_, d_1056_v52_)
                d_1059_v53_ = nw196_
                d_1060_v54_: _dafny.Seq
                d_1060_v54_ = _dafny.SeqWithoutIsStrInference([d_1059_v53_, d_1059_v53_])
                d_1061_v55_: _dafny.Array
                nw197_ = _dafny.Array(None, 15)
                nw197_[int(0)] = d_1055_cf6_
                nw197_[int(1)] = d_976_v1_
                nw197_[int(2)] = not (True) or (d_976_v1_)
                nw197_[int(3)] = d_1055_cf6_
                nw197_[int(4)] = (d_1060_v54_) != (_dafny.SeqWithoutIsStrInference([d_1059_v53_]))
                nw197_[int(5)] = d_976_v1_
                nw197_[int(6)] = d_1055_cf6_
                nw197_[int(7)] = (d_1059_v53_).f10
                nw197_[int(8)] = not (d_976_v1_) or (d_1055_cf6_)
                nw197_[int(9)] = (d_975_v0_) != (_dafny.SeqWithoutIsStrInference([d_1046_v46_ for d_1062_i10_ in range(default__.abs(130))]))
                nw197_[int(10)] = d_1055_cf6_
                nw197_[int(11)] = d_976_v1_
                nw197_[int(12)] = d_976_v1_
                nw197_[int(13)] = d_1055_cf6_
                nw197_[int(14)] = d_976_v1_
                d_1061_v55_ = nw197_
                index134_ = default__.safeIndex(638, (d_1061_v55_).length(0))
                (d_1061_v55_)[index134_] = not(d_1055_cf6_)
                index135_ = default__.safeIndex(638, (d_1061_v55_).length(0))
                (d_1061_v55_)[index135_] = ((self).f9) > (p0)
                (globalState).f2 = (d_1059_v53_).f10
                d_1063_v56_: _dafny.Array
                nw198_ = _dafny.Array(int(0), 26)
                d_1063_v56_ = nw198_
                d_1063_v56_ = (d_1063_v56_ if d_1055_cf6_ else d_1063_v56_)
                d_1064_v57_: _dafny.Seq
                d_1064_v57_ = _dafny.SeqWithoutIsStrInference([(d_1061_v55_)[default__.safeIndex(638, (d_1061_v55_).length(0))], (d_1059_v53_).f10])
                d_1065_v58_: _dafny.Map
                d_1065_v58_ = _dafny.Map({d_1046_v46_: (d_1059_v53_).f11})
                d_1066_v59_: D4
                d_1066_v59_ = D4_DC13(len(d_1064_v57_), ((d_1065_v58_)[d_1046_v46_] if (d_1046_v46_) in (d_1065_v58_) else d_1061_v55_))
                d_1067_v60_: _dafny.Map
                d_1067_v60_ = _dafny.Map({(d_1066_v59_).cf20: d_1056_v52_})
                d_1068_v61_: T1
                nw199_ = C0()
                nw199_.ctor__((d_1059_v53_).f10, ((d_1067_v60_)[d_1056_v52_] if (d_1056_v52_) in (d_1067_v60_) else d_1056_v52_))
                d_1068_v61_ = nw199_
                d_1069_v62_: _dafny.Map
                d_1069_v62_ = _dafny.Map({p1: d_1068_v61_})
                d_1070_v63_: _dafny.Seq
                d_1070_v63_ = _dafny.SeqWithoutIsStrInference([d_1045_v45_])
                d_1071_v64_: _dafny.Seq
                d_1071_v64_ = _dafny.SeqWithoutIsStrInference([len(d_1064_v57_), p1])
                d_1072_v65_: _dafny.Set
                d_1072_v65_ = _dafny.Set({len(d_1064_v57_)})
                d_1073_v66_: _dafny.Array
                nw200_ = _dafny.Array(None, 16)
                nw200_[int(0)] = d_1069_v62_
                nw200_[int(1)] = d_1069_v62_
                nw200_[int(2)] = d_1069_v62_
                nw200_[int(3)] = d_1069_v62_
                nw200_[int(4)] = d_1069_v62_
                nw200_[int(5)] = (d_1069_v62_).set((self).f9, d_1068_v61_)
                nw200_[int(6)] = d_1069_v62_
                nw200_[int(7)] = d_1069_v62_
                nw200_[int(8)] = _dafny.Map({p1: d_1068_v61_})
                nw200_[int(9)] = d_1069_v62_
                nw200_[int(10)] = _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False]))).cardinality: d_1068_v61_})
                nw200_[int(11)] = (d_1069_v62_) | (d_1069_v62_)
                nw200_[int(12)] = _dafny.Map({(d_1049_v49_).cf7: d_1068_v61_})
                nw200_[int(13)] = (d_1069_v62_).set((self).fm5(d_1070_v63_, d_1046_v46_, d_1071_v64_, _dafny.SeqWithoutIsStrInference([(0) - (len(d_1072_v65_)), p0]), globalState), d_1068_v61_)
                nw200_[int(14)] = d_1069_v62_
                nw200_[int(15)] = (d_1069_v62_) | (d_1069_v62_)
                d_1073_v66_ = nw200_
                index136_ = default__.safeIndex(787, (d_1073_v66_).length(0))
                (d_1073_v66_)[index136_] = d_1069_v62_
                index137_ = default__.safeIndex(787, (d_1073_v66_).length(0))
                (d_1073_v66_)[index137_] = d_1069_v62_
            elif source16_.is_DC7:
                d_1074___mcc_h7_ = source16_.cf7
                d_1075___mcc_h8_ = source16_.cf8
                d_1076___mcc_h9_ = source16_.cf9
                d_1077_cf9_ = d_1076___mcc_h9_
                d_1078_cf8_ = d_1075___mcc_h8_
                d_1079_cf7_ = d_1074___mcc_h7_
                (globalState).f0 = ((d_976_v1_ if d_976_v1_ else d_976_v1_)) and (d_976_v1_)
                d_1080_v67_: _dafny.Seq
                d_1080_v67_ = _dafny.SeqWithoutIsStrInference([d_1079_cf7_, len(d_1047_v47_)])
                d_1080_v67_ = d_1080_v67_
                d_1049_v49_ = d_1049_v49_
                d_1081_v68_: D0
                d_1081_v68_ = D0_DC1(False)
                d_1082_v69_: _dafny.Seq
                d_1082_v69_ = _dafny.SeqWithoutIsStrInference([d_976_v1_, d_976_v1_, d_976_v1_, d_976_v1_, (d_1081_v68_).cf1])
                d_1083_v70_: _dafny.Array
                nw201_ = _dafny.Array(False, 9)
                d_1083_v70_ = nw201_
                d_1084_v71_: D4
                d_1084_v71_ = D4_DC13(d_1077_cf9_, d_1083_v70_)
                (globalState).f0 = (d_1082_v69_)[default__.safeIndex((d_1084_v71_).cf19, len(d_1082_v69_))]
            elif True:
                d_1085___mcc_h10_ = source16_.cf5
                d_1086_cf5_ = d_1085___mcc_h10_
                r2 = True
                d_1087_v72_: _dafny.Seq
                d_1087_v72_ = _dafny.SeqWithoutIsStrInference([d_976_v1_])
                d_1088_v73_: D0
                d_1088_v73_ = D0_DC1((d_1087_v72_)[default__.safeIndex(p0, len(d_1087_v72_))])
                d_1089_v74_: _dafny.Map
                d_1089_v74_ = _dafny.Map({D3_DC9(610): d_1088_v73_})
                d_1089_v74_ = d_1089_v74_
                d_1090_v75_: _dafny.Array
                def lambda68_(d_1091_v1_):
                    def lambda69_(d_1092_i11_):
                        return d_1091_v1_

                    return lambda69_

                init35_ = lambda68_(d_976_v1_)
                nw202_ = _dafny.Array(None, 10)
                for i0_35_ in range(nw202_.length(0)):
                    nw202_[i0_35_] = init35_(i0_35_)
                d_1090_v75_ = nw202_
                d_1093_v76_: C0
                nw203_ = C0()
                nw203_.ctor__(True, d_1090_v75_)
                d_1093_v76_ = nw203_
                d_1094_v77_: _dafny.Array
                def lambda70_(d_1095_p1_):
                    def lambda71_(d_1096_i12_):
                        return (d_1096_i12_) + (d_1095_p1_)

                    return lambda71_

                init36_ = lambda70_(p1)
                nw204_ = _dafny.Array(None, 10)
                for i0_36_ in range(nw204_.length(0)):
                    nw204_[i0_36_] = init36_(i0_36_)
                d_1094_v77_ = nw204_
                d_1097_v78_: _dafny.Map
                d_1097_v78_ = _dafny.Map({d_1093_v76_: d_1094_v77_})
                d_1098_v79_: _dafny.Map
                d_1098_v79_ = _dafny.Map({(d_1093_v76_).f10: d_1047_v47_})
                d_1099_v80_: _dafny.Map
                d_1099_v80_ = _dafny.Map({default__.fm0(len(d_1098_v79_), p0, p1, (d_1093_v76_).f10, globalState): p1})
                rhs168_ = d_976_v1_
                rhs169_ = (len((d_1097_v78_) | (d_1097_v78_)) if d_976_v1_ else ((self).f9 if (d_1093_v76_).f10 else p0))
                rhs170_ = len((_dafny.Map({(self).f9: p0})) | ((d_1099_v80_) | (d_1099_v80_)))
                lhs123_ = globalState
                lhs124_ = globalState
                lhs123_.f2 = rhs168_
                r1 = rhs169_
                lhs124_.f5 = rhs170_
                d_1100_v81_: _dafny.MultiSet
                d_1100_v81_ = _dafny.MultiSet([d_1042_v42_, (d_1042_v42_) | (d_1042_v42_), d_1042_v42_, d_1042_v42_])
                index138_ = default__.safeIndex(526, ((d_1093_v76_).f11).length(0))
                ((d_1093_v76_).f11)[index138_] = (d_1093_v76_).f10
                d_1101_v82_: _dafny.MultiSet
                d_1101_v82_ = _dafny.MultiSet([d_975_v0_])
                d_1102_v83_: _dafny.Set
                d_1102_v83_ = _dafny.Set({p1})
                index139_ = default__.safeIndex(526, ((d_1093_v76_).f11).length(0))
                rhs171_ = p0
                rhs172_ = ((d_1101_v82_).cardinality) * (((self).f9 if (d_1093_v76_).f10 else default__.fm0(len(d_1102_v83_), p0, p1, (d_1053_v51_).cf6, globalState)))
                rhs173_ = len(((d_1087_v72_) + (_dafny.SeqWithoutIsStrInference([(d_1093_v76_).f10]))) + (d_1087_v72_))
                rhs174_ = d_1100_v81_
                rhs175_ = d_976_v1_
                lhs125_ = globalState
                lhs126_ = globalState
                lhs127_ = globalState
                lhs128_ = (d_1093_v76_).f11
                lhs129_ = default__.safeIndex(526, ((d_1093_v76_).f11).length(0))
                lhs125_.f5 = rhs171_
                lhs126_.f5 = rhs172_
                lhs127_.f5 = rhs173_
                d_1100_v81_ = rhs174_
                lhs128_[lhs129_] = rhs175_
            if (p1) <= (len(d_1045_v45_)):
                d_976_v1_ = d_976_v1_
                d_1103_v84_: D0
                d_1103_v84_ = D0_DC1(d_976_v1_)
                d_1104_v85_: _dafny.Array
                nw205_ = _dafny.Array(None, 3)
                nw205_[int(0)] = False
                nw205_[int(1)] = d_976_v1_
                nw205_[int(2)] = (d_1103_v84_).cf1
                d_1104_v85_ = nw205_
                d_1105_v86_: C0
                nw206_ = C0()
                nw206_.ctor__(not(not(((self).f9) >= ((self).f9))), d_1104_v85_)
                d_1105_v86_ = nw206_
                index140_ = default__.safeIndex(991, (d_1104_v85_).length(0))
                (d_1104_v85_)[index140_] = not (d_976_v1_) or ((d_1105_v86_).f10)
                index141_ = default__.safeIndex(991, (d_1104_v85_).length(0))
                (d_1104_v85_)[index141_] = not(d_976_v1_)
                d_1106_v87_: _dafny.Map
                d_1106_v87_ = _dafny.Map({d_1044_v44_: (d_1104_v85_)[default__.safeIndex(991, (d_1104_v85_).length(0))]})
                r1 = len(((d_1106_v87_) | (d_1106_v87_)) | (d_1106_v87_))
                d_1107_v88_: _dafny.Array
                nw207_ = _dafny.Array(_dafny.Seq({}), 19)
                d_1107_v88_ = nw207_
                d_1108_v89_: T0
                nw208_ = C2()
                nw208_.ctor__(not((d_1105_v86_).f10))
                d_1108_v89_ = nw208_
                d_1109_v90_: _dafny.Seq
                d_1109_v90_ = _dafny.SeqWithoutIsStrInference([(d_1104_v85_)[default__.safeIndex(991, (d_1104_v85_).length(0))], (d_1105_v86_).f10, (d_1104_v85_)[default__.safeIndex(991, (d_1104_v85_).length(0))]])
                d_1110_v91_: D7
                d_1110_v91_ = D7_DC17(d_1109_v90_)
                d_1111_v92_: _dafny.MultiSet
                d_1111_v92_ = _dafny.MultiSet([D7_DC17(_dafny.SeqWithoutIsStrInference([(d_1105_v86_).f10])), d_1110_v91_])
                d_1112_v93_: _dafny.MultiSet
                d_1112_v93_ = _dafny.MultiSet([d_1108_v89_])
                rhs176_ = d_1107_v88_
                rhs177_ = (d_1104_v85_)[default__.safeIndex(991, (d_1104_v85_).length(0))]
                rhs178_ = (d_1105_v86_).f10
                rhs179_ = d_1108_v89_
                rhs180_ = _dafny.Set({d_976_v1_, (d_1111_v92_).isdisjoint(d_1111_v92_), (d_1112_v93_).isdisjoint(d_1112_v93_)})
                lhs130_ = globalState
                lhs131_ = globalState
                d_1107_v88_ = rhs176_
                lhs130_.f2 = rhs177_
                r2 = rhs178_
                d_1108_v89_ = rhs179_
                lhs131_.f8 = rhs180_
            elif True:
                rhs181_ = d_976_v1_
                rhs182_ = p1
                rhs183_ = not(not(d_976_v1_))
                rhs184_ = (self).f9
                lhs132_ = globalState
                lhs133_ = globalState
                lhs134_ = globalState
                lhs135_ = globalState
                lhs132_.f2 = rhs181_
                lhs133_.f5 = rhs182_
                lhs134_.f2 = rhs183_
                lhs135_.f5 = rhs184_
                d_1113_v94_: _dafny.Map
                d_1113_v94_ = _dafny.Map({(p0) - (755): (0) - (p0)})
                d_1114_v96_: _dafny.Seq
                d_1114_v96_ = _dafny.SeqWithoutIsStrInference([d_976_v1_, not(d_976_v1_)])
                def iife88_():
                    coll48_ = _dafny.Map()
                    compr_48_: int
                    for compr_48_ in (_dafny.Map({p0: True})).keys.Elements:
                        d_1115_v95_: int = compr_48_
                        if (d_1115_v95_) in (_dafny.Map({p0: True})):
                            coll48_[(d_1115_v95_) - ((0) - (len(_dafny.Set({p0, p0}))))] = len(d_975_v0_)
                    return _dafny.Map(coll48_)
                d_1113_v94_ = (d_1113_v94_).set(default__.fm0(p0, p1, len(iife88_()
                ), True, globalState), len(d_1114_v96_))
                (globalState).f0 = True
                d_1116_v97_: _dafny.Array
                nw209_ = _dafny.Array(int(0), 19)
                d_1116_v97_ = nw209_
                index142_ = default__.safeIndex(230, (d_1116_v97_).length(0))
                (d_1116_v97_)[index142_] = p0
                d_1117_v98_: _dafny.Array
                nw210_ = _dafny.Array(False, 14)
                d_1117_v98_ = nw210_
                d_1118_v99_: C0
                nw211_ = C0()
                nw211_.ctor__(d_976_v1_, d_1117_v98_)
                d_1118_v99_ = nw211_
                d_1119_v100_: _dafny.MultiSet
                d_1119_v100_ = _dafny.MultiSet([d_1118_v99_, d_1118_v99_, d_1118_v99_, d_1118_v99_, d_1118_v99_])
                d_1120_v101_: _dafny.Seq
                d_1120_v101_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(_dafny.MultiSet([p1])).cardinality: (self).f9}))])
                d_1121_v102_: _dafny.Seq
                d_1121_v102_ = _dafny.SeqWithoutIsStrInference([d_975_v0_, d_975_v0_])
                index143_ = default__.safeIndex(230, (d_1116_v97_).length(0))
                rhs185_ = ((d_1119_v100_)[d_1118_v99_] if (d_1118_v99_) in (d_1119_v100_) else (0) - (((d_1120_v101_)[default__.safeIndex(len(d_1114_v96_), len(d_1120_v101_))] if (d_1118_v99_).f10 else (self).f9)))
                rhs186_ = (0) - (((p1) + (p1)) * (p1))
                rhs187_ = (len(d_1121_v102_) if default__.fm1((d_1118_v99_).f10, globalState) else (p1) * ((self).f9))
                lhs136_ = globalState
                lhs137_ = d_1116_v97_
                lhs138_ = default__.safeIndex(230, (d_1116_v97_).length(0))
                r1 = rhs185_
                lhs136_.f5 = rhs186_
                lhs137_[lhs138_] = rhs187_
                d_1122_v103_: C3
                nw212_ = C3()
                nw212_.ctor__((0) - (p0))
                d_1122_v103_ = nw212_
                d_1123_v104_: _dafny.Array
                nw213_ = _dafny.Array(None, 8)
                nw213_[int(0)] = d_1122_v103_
                nw213_[int(1)] = d_1122_v103_
                nw213_[int(2)] = d_1122_v103_
                nw213_[int(3)] = d_1122_v103_
                nw213_[int(4)] = d_1122_v103_
                nw213_[int(5)] = d_1122_v103_
                nw213_[int(6)] = d_1122_v103_
                nw213_[int(7)] = d_1122_v103_
                d_1123_v104_ = nw213_
                d_1124_v105_: _dafny.Set
                d_1124_v105_ = _dafny.Set({d_1046_v46_})
                index144_ = default__.safeIndex(230, (d_1116_v97_).length(0))
                rhs188_ = (((self).f9 if d_976_v1_ else len(d_1124_v105_))) <= (-348)
                rhs189_ = (d_1114_v96_) != (d_1114_v96_)
                rhs190_ = d_1123_v104_
                rhs191_ = default__.fm27(globalState)
                rhs192_ = (d_1116_v97_)[default__.safeIndex(230, (d_1116_v97_).length(0))]
                lhs139_ = globalState
                lhs140_ = d_1116_v97_
                lhs141_ = default__.safeIndex(230, (d_1116_v97_).length(0))
                lhs139_.f2 = rhs188_
                d_976_v1_ = rhs189_
                d_1123_v104_ = rhs190_
                d_975_v0_ = rhs191_
                lhs140_[lhs141_] = rhs192_
            d_1125_v106_: _dafny.Array
            def lambda72_(d_1126_v1_):
                def lambda73_(d_1127_i13_):
                    return d_1126_v1_

                return lambda73_

            init37_ = lambda72_(d_976_v1_)
            nw214_ = _dafny.Array(None, 1)
            for i0_37_ in range(nw214_.length(0)):
                nw214_[i0_37_] = init37_(i0_37_)
            d_1125_v106_ = nw214_
            index145_ = default__.safeIndex(797, (d_1125_v106_).length(0))
            (d_1125_v106_)[index145_] = (d_976_v1_ if d_976_v1_ else False)
            index146_ = default__.safeIndex(797, (d_1125_v106_).length(0))
            (d_1125_v106_)[index146_] = d_976_v1_
        elif True:
            d_1128_v107_: str
            d_1128_v107_ = _dafny.CodePoint('k')
            d_1129_v108_: _dafny.Map
            d_1129_v108_ = _dafny.Map({p0: d_1128_v107_})
            (globalState).f5 = (len(d_1129_v108_)) * (535)
            r1 = (0) - (p0)
            d_1130_v109_: _dafny.Map
            d_1130_v109_ = _dafny.Map({(self).f9: d_976_v1_})
            d_1131_v110_: D0
            d_1131_v110_ = D0_DC0(d_1130_v109_)
            d_1132_v111_: D0
            d_1132_v111_ = D0_DC2(d_1131_v110_)
            d_1133_v112_: D0
            d_1133_v112_ = D0_DC2(d_1132_v111_)
            source17_ = D0_DC2((d_1132_v111_ if d_976_v1_ else d_1132_v111_))
            if source17_.is_DC1:
                d_1134___mcc_h11_ = source17_.cf1
                d_1135_cf1_ = d_1134___mcc_h11_
                d_1136_v113_: _dafny.Set
                d_1136_v113_ = _dafny.Set({_dafny.CodePoint('l'), _dafny.CodePoint('m'), d_1128_v107_})
                d_1137_v114_: _dafny.Seq
                d_1137_v114_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1138_v115_: _dafny.Seq
                d_1138_v115_ = _dafny.SeqWithoutIsStrInference([d_1135_cf1_, False])
                d_1139_v116_: _dafny.Set
                d_1139_v116_ = _dafny.Set({d_976_v1_, d_1135_cf1_})
                d_1140_v117_: _dafny.Map
                d_1140_v117_ = _dafny.Map({d_1139_v116_: d_1135_cf1_})
                d_1141_v118_: _dafny.Seq
                d_1141_v118_ = _dafny.SeqWithoutIsStrInference([(d_1140_v117_).set(_dafny.Set({d_976_v1_}), True), d_1140_v117_, _dafny.Map({d_1139_v116_: d_1135_cf1_})])
                d_1142_v119_: _dafny.Map
                d_1142_v119_ = _dafny.Map({889: p1})
                d_1143_v120_: _dafny.Seq
                d_1143_v120_ = _dafny.SeqWithoutIsStrInference([d_975_v0_])
                d_1144_v121_: _dafny.Seq
                d_1144_v121_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1128_v107_ for d_1145_i14_ in range(default__.abs(406))]), (d_1143_v120_)[default__.safeIndex(p0, len(d_1143_v120_))], d_975_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nikx")), d_975_v0_])
                d_1146_v122_: _dafny.Array
                nw215_ = _dafny.Array(None, 25)
                nw215_[int(0)] = p0
                nw215_[int(1)] = (default__.fm0((self).f9, (self).f9, p0, d_1135_cf1_, globalState) if ((d_1130_v109_)[len(d_1136_v113_)] if (len(d_1136_v113_)) in (d_1130_v109_) else d_976_v1_) else (self).f9)
                nw215_[int(2)] = (self).f9
                nw215_[int(3)] = (d_1137_v114_)[default__.safeIndex((self).f9, len(d_1137_v114_))]
                nw215_[int(4)] = ((self).f9 if d_976_v1_ else (self).f9)
                nw215_[int(5)] = (self).f9
                nw215_[int(6)] = default__.fm0(len(d_1138_v115_), -482, p1, True, globalState)
                nw215_[int(7)] = (self).f9
                nw215_[int(8)] = p1
                nw215_[int(9)] = p0
                nw215_[int(10)] = (p0) * (len(_dafny.Set({(self).f9})))
                nw215_[int(11)] = p0
                nw215_[int(12)] = len((d_1141_v118_)[default__.safeIndex((0) - (len(d_1137_v114_)), len(d_1141_v118_))])
                nw215_[int(13)] = p0
                nw215_[int(14)] = len((d_1142_v119_) | (d_1142_v119_))
                nw215_[int(15)] = (self).f9
                nw215_[int(16)] = (len(d_1138_v115_)) * (len(d_1139_v116_))
                nw215_[int(17)] = p1
                nw215_[int(18)] = len(((d_1144_v121_)[default__.safeIndex(default__.fm0(len(d_975_v0_), 378, (self).f9, True, globalState), len(d_1144_v121_))]).set(default__.safeIndex(571, len((d_1144_v121_)[default__.safeIndex(default__.fm0(len(d_975_v0_), 378, (self).f9, True, globalState), len(d_1144_v121_))])), d_1128_v107_))
                nw215_[int(19)] = p1
                nw215_[int(20)] = p0
                nw215_[int(21)] = (self).f9
                nw215_[int(22)] = 21
                nw215_[int(23)] = p1
                nw215_[int(24)] = p1
                d_1146_v122_ = nw215_
                d_1146_v122_ = d_1146_v122_
                (globalState).f8 = ((d_1139_v116_) - (d_1139_v116_)) | ((d_1139_v116_) | (d_1139_v116_))
                d_1147_v123_: C3
                nw216_ = C3()
                nw216_.ctor__(p0)
                d_1147_v123_ = nw216_
                d_1148_v124_: _dafny.Array
                nw217_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 12)
                d_1148_v124_ = nw217_
                index147_ = default__.safeIndex(389, (d_1148_v124_).length(0))
                (d_1148_v124_)[index147_] = d_975_v0_
                d_1149_v125_: _dafny.Array
                def lambda74_(d_1150_v1_):
                    def lambda75_(d_1151_i15_):
                        return d_1150_v1_

                    return lambda75_

                init38_ = lambda74_(d_976_v1_)
                nw218_ = _dafny.Array(None, 19)
                for i0_38_ in range(nw218_.length(0)):
                    nw218_[i0_38_] = init38_(i0_38_)
                d_1149_v125_ = nw218_
                index148_ = default__.safeIndex(478, (d_1149_v125_).length(0))
                (d_1149_v125_)[index148_] = d_976_v1_
                d_1152_v126_: _dafny.Set
                d_1152_v126_ = _dafny.Set({(d_1147_v123_).f12})
                index149_ = default__.safeIndex(389, (d_1148_v124_).length(0))
                index150_ = default__.safeIndex(478, (d_1149_v125_).length(0))
                rhs193_ = ((d_975_v0_) + (_dafny.SeqWithoutIsStrInference([d_1128_v107_ for d_1153_i16_ in range(default__.abs(-283))]))).set(default__.safeIndex(p0, len((d_975_v0_) + (_dafny.SeqWithoutIsStrInference([d_1128_v107_ for d_1154_i16_ in range(default__.abs(-283))])))), d_1128_v107_)
                rhs194_ = d_976_v1_
                rhs195_ = d_976_v1_
                rhs196_ = 292
                rhs197_ = ((d_1152_v126_) - (d_1152_v126_)).issubset(d_1152_v126_)
                lhs142_ = d_1148_v124_
                lhs143_ = default__.safeIndex(389, (d_1148_v124_).length(0))
                lhs144_ = d_1149_v125_
                lhs145_ = default__.safeIndex(478, (d_1149_v125_).length(0))
                lhs146_ = globalState
                lhs147_ = globalState
                lhs142_[lhs143_] = rhs193_
                r2 = rhs194_
                lhs144_[lhs145_] = rhs195_
                lhs146_.f5 = rhs196_
                lhs147_.f2 = rhs197_
            elif source17_.is_DC0:
                d_1155___mcc_h12_ = source17_.cf0
                d_1156_cf0_ = d_1155___mcc_h12_
                d_1157_v127_: _dafny.Array
                def lambda76_(d_1158_v1_):
                    def lambda77_(d_1159_i17_):
                        return not (d_1158_v1_) or (False)

                    return lambda77_

                init39_ = lambda76_(d_976_v1_)
                nw219_ = _dafny.Array(None, 3)
                for i0_39_ in range(nw219_.length(0)):
                    nw219_[i0_39_] = init39_(i0_39_)
                d_1157_v127_ = nw219_
                d_1160_v128_: D4
                d_1160_v128_ = D4_DC11(d_1157_v127_)
                pat_let_tv38_ = d_1157_v127_
                def iife89_(_pat_let20_0):
                    def iife90_(d_1161_dt__update__tmp_h2_):
                        def iife91_(_pat_let21_0):
                            def iife92_(d_1162_dt__update_hcf13_h0_):
                                return D4_DC11(d_1162_dt__update_hcf13_h0_)
                            return iife92_(_pat_let21_0)
                        return iife91_(pat_let_tv38_)
                    return iife90_(_pat_let20_0)
                d_1157_v127_ = (iife89_(d_1160_v128_)).cf13
                d_1163_v129_: C1
                nw220_ = C1()
                nw220_.ctor__()
                d_1163_v129_ = nw220_
                d_1164_v130_: C3
                nw221_ = C3()
                nw221_.ctor__((0) - ((self).f9))
                d_1164_v130_ = nw221_
                d_1156_cf0_ = (d_1156_cf0_).set((d_1164_v130_).f12, d_976_v1_)
            elif True:
                d_1165___mcc_h13_ = source17_.cf2
                d_1166_cf2_ = d_1165___mcc_h13_
                (globalState).f5 = (self).f9
                (globalState).f5 = p1
                r1 = (self).f9
                d_1167_v131_: C1
                nw222_ = C1()
                nw222_.ctor__()
                d_1167_v131_ = nw222_
            d_1168_v132_: _dafny.Array
            nw223_ = _dafny.Array(None, 1)
            nw223_[int(0)] = not (True) or (d_976_v1_)
            d_1168_v132_ = nw223_
            d_1169_v133_: _dafny.Map
            d_1169_v133_ = _dafny.Map({d_975_v0_: d_976_v1_})
            d_1170_v134_: _dafny.MultiSet
            d_1170_v134_ = _dafny.MultiSet([d_976_v1_, d_976_v1_])
            index151_ = default__.safeIndex(189, (d_1168_v132_).length(0))
            (d_1168_v132_)[index151_] = ((d_1169_v133_)[(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rngu"))).set(default__.safeIndex((d_1170_v134_).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rngu")))), d_1128_v107_)] if ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rngu"))).set(default__.safeIndex((d_1170_v134_).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rngu")))), d_1128_v107_)) in (d_1169_v133_) else d_976_v1_)
            index152_ = default__.safeIndex(189, (d_1168_v132_).length(0))
            (d_1168_v132_)[index152_] = d_976_v1_
            r1 = (d_1170_v134_).cardinality
        d_1171_v135_: C1
        nw224_ = C1()
        nw224_.ctor__()
        d_1171_v135_ = nw224_
        d_1172_v136_: _dafny.Map
        d_1172_v136_ = _dafny.Map({d_1171_v135_: d_976_v1_})
        d_1173_v137_: _dafny.Map
        d_1173_v137_ = _dafny.Map({d_976_v1_: d_976_v1_})
        d_1174_v138_: str
        d_1174_v138_ = _dafny.CodePoint('o')
        d_1175_v139_: D10
        d_1175_v139_ = D10_DC27(d_976_v1_, p1, d_1174_v138_, _dafny.MultiSet([d_976_v1_]))
        d_1176_v140_: _dafny.Seq
        d_1176_v140_ = _dafny.SeqWithoutIsStrInference([d_975_v0_, d_975_v0_, d_975_v0_])
        d_1177_v141_: _dafny.Set
        d_1177_v141_ = _dafny.Set({(d_1176_v140_)[default__.safeIndex(p0, len(d_1176_v140_))], d_975_v0_})
        d_1178_v142_: _dafny.Array
        nw225_ = _dafny.Array(None, 27)
        nw225_[int(0)] = not(d_976_v1_)
        nw225_[int(1)] = d_976_v1_
        nw225_[int(2)] = (((d_1172_v136_)[d_1171_v135_] if (d_1171_v135_) in (d_1172_v136_) else d_976_v1_) if not(not(d_976_v1_)) else False)
        nw225_[int(3)] = default__.fm1(d_976_v1_, globalState)
        nw225_[int(4)] = not (d_976_v1_) or (d_976_v1_)
        nw225_[int(5)] = (d_976_v1_) == (d_976_v1_)
        nw225_[int(6)] = not(True)
        nw225_[int(7)] = (_dafny.Map({p1: d_976_v1_})) != (_dafny.Map({p1: default__.fm1(d_976_v1_, globalState)}))
        nw225_[int(8)] = not (d_976_v1_) or (((d_1173_v137_)[d_976_v1_] if (d_976_v1_) in (d_1173_v137_) else d_976_v1_))
        nw225_[int(9)] = (D9_DC24(d_976_v1_, d_976_v1_, d_976_v1_)).cf35
        nw225_[int(10)] = d_976_v1_
        nw225_[int(11)] = d_976_v1_
        nw225_[int(12)] = d_976_v1_
        nw225_[int(13)] = d_976_v1_
        nw225_[int(14)] = not(False)
        nw225_[int(15)] = not(False)
        nw225_[int(16)] = d_976_v1_
        nw225_[int(17)] = ((self).f9) > ((self).f9)
        nw225_[int(18)] = (d_1171_v135_).fm4(d_977_v2_, d_977_v2_, (d_1175_v139_).cf41, globalState)
        nw225_[int(19)] = False
        nw225_[int(20)] = (d_1177_v141_).issubset(_dafny.Set({d_975_v0_, d_975_v0_, d_975_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))}))
        nw225_[int(21)] = (p1) == ((self).f9)
        nw225_[int(22)] = d_976_v1_
        nw225_[int(23)] = not (d_976_v1_) or (d_976_v1_)
        nw225_[int(24)] = (d_976_v1_) and (d_976_v1_)
        nw225_[int(25)] = d_976_v1_
        nw225_[int(26)] = d_976_v1_
        d_1178_v142_ = nw225_
        index153_ = default__.safeIndex(902, (d_1178_v142_).length(0))
        (d_1178_v142_)[index153_] = (d_976_v1_) and (d_976_v1_)
        d_1179_v143_: _dafny.Array
        def lambda78_(d_1180_v1_):
            def lambda79_(d_1181_i18_):
                return _dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([607 for d_1182_i19_ in range(default__.abs(912))])): _dafny.Set({d_1180_v1_, d_1180_v1_})})), 241])

            return lambda79_

        init40_ = lambda78_(d_976_v1_)
        nw226_ = _dafny.Array(None, 2)
        for i0_40_ in range(nw226_.length(0)):
            nw226_[i0_40_] = init40_(i0_40_)
        d_1179_v143_ = nw226_
        index154_ = default__.safeIndex(446, (d_1179_v143_).length(0))
        (d_1179_v143_)[index154_] = _dafny.MultiSet([p1])
        d_1183_v144_: _dafny.MultiSet
        d_1183_v144_ = _dafny.MultiSet([(_dafny.MultiSet([d_975_v0_, d_975_v0_, d_975_v0_])).cardinality, (self).f9, (self).f9])
        index155_ = default__.safeIndex(902, (d_1178_v142_).length(0))
        index156_ = default__.safeIndex(446, (d_1179_v143_).length(0))
        rhs198_ = d_976_v1_
        rhs199_ = d_1183_v144_
        lhs148_ = d_1178_v142_
        lhs149_ = default__.safeIndex(902, (d_1178_v142_).length(0))
        lhs150_ = d_1179_v143_
        lhs151_ = default__.safeIndex(446, (d_1179_v143_).length(0))
        lhs148_[lhs149_] = rhs198_
        lhs150_[lhs151_] = rhs199_
        d_1184_v145_: C2
        nw227_ = C2()
        nw227_.ctor__(d_976_v1_)
        d_1184_v145_ = nw227_
        d_1185_v146_: D7
        d_1185_v146_ = D7_DC18()
        d_1185_v146_ = (d_1185_v146_ if (d_1178_v142_)[default__.safeIndex(902, (d_1178_v142_).length(0))] else d_1185_v146_)
        d_1186_v147_: D11
        d_1186_v147_ = D11_DC31(p0, (0) - ((self).f9))
        index157_ = default__.safeIndex(902, (d_1178_v142_).length(0))
        (d_1178_v142_)[index157_] = ((d_1186_v147_).cf45) < (p1)
        d_1187_v148_: _dafny.Seq
        d_1187_v148_ = _dafny.SeqWithoutIsStrInference([(d_1184_v145_).f13])
        d_1188_v149_: _dafny.Seq
        d_1188_v149_ = _dafny.SeqWithoutIsStrInference([d_1187_v148_, d_1187_v148_])
        d_1189_v150_: _dafny.Map
        d_1189_v150_ = _dafny.Map({default__.fm32(d_1188_v149_, globalState): True})
        d_1190_v151_: _dafny.Map
        d_1190_v151_ = _dafny.Map({False: len(d_1189_v150_)})
        r0 = d_1190_v151_
        d_1191_v152_: _dafny.MultiSet
        d_1191_v152_ = _dafny.MultiSet([(d_1178_v142_)[default__.safeIndex(902, (d_1178_v142_).length(0))]])
        d_1192_v153_: D7
        d_1192_v153_ = D7_DC19(D7_DC18())
        d_1193_v154_: _dafny.Map
        d_1193_v154_ = _dafny.Map({d_1192_v153_: (d_1178_v142_)[default__.safeIndex(902, (d_1178_v142_).length(0))]})
        d_1194_v155_: _dafny.Map
        d_1194_v155_ = _dafny.Map({d_1176_v140_: (self).f9})
        r1 = (default__.safeDivisionInt(p0, ((d_1191_v152_)[((d_1193_v154_)[d_1192_v153_] if (d_1192_v153_) in (d_1193_v154_) else (d_1178_v142_)[default__.safeIndex(902, (d_1178_v142_).length(0))])] if (((d_1193_v154_)[d_1192_v153_] if (d_1192_v153_) in (d_1193_v154_) else (d_1178_v142_)[default__.safeIndex(902, (d_1178_v142_).length(0))])) in (d_1191_v152_) else -660))) + (((d_1194_v155_)[_dafny.SeqWithoutIsStrInference([d_975_v0_])] if (_dafny.SeqWithoutIsStrInference([d_975_v0_])) in (d_1194_v155_) else p0))
        d_1195_v156_: _dafny.Set
        d_1195_v156_ = _dafny.Set({d_1178_v142_})
        r2 = (default__.safeModuloInt(len(d_1195_v156_), 326)) == (p1)
        return r0, r1, r2

    @property
    def f9(self):
        return self._f9
